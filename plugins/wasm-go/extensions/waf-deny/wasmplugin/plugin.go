// Copyright The OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package wasmplugin

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza/v3/debuglog"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"math"
	"net"
	"strconv"
	"strings"
)

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

func NewVMContext() types.VMContext {
	return &vmContext{}
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &corazaPlugin{}
}

type corazaPlugin struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

func (ctx *corazaPlugin) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	return types.OnPluginStartStatusOK
}

func (ctx *corazaPlugin) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpContext{
		contextID: contextID,
	}
}

type interruptionPhase int8

func (p interruptionPhase) isInterrupted() bool {
	return p != interruptionPhaseNone
}

func (p interruptionPhase) String() string {
	switch p {
	case interruptionPhaseHttpRequestHeaders:
		return "http_request_headers"
	case interruptionPhaseHttpRequestBody:
		return "http_request_body"
	case interruptionPhaseHttpResponseHeaders:
		return "http_response_headers"
	case interruptionPhaseHttpResponseBody:
		return "http_response_body"
	default:
		return "no interruption yet"
	}
}

const (
	interruptionPhaseNone                = iota
	interruptionPhaseHttpRequestHeaders  = iota
	interruptionPhaseHttpRequestBody     = iota
	interruptionPhaseHttpResponseHeaders = iota
	interruptionPhaseHttpResponseBody    = iota
)

type httpContext struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID             uint32
	tx                    core.Transaction
	httpProtocol          string
	processedRequestBody  bool
	processedResponseBody bool
	bodyReadIndex         int
}

func (ctx *httpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	ctx.tx = core.NewTransaction()
	authority, err := proxywasm.GetHttpRequestHeader(":authority")
	if err != nil {
		proxywasm.LogWarnf("Failed to resolve WAF for authority %q ", authority)
		return types.ActionContinue
	}

	ctx.tx.AddRequestHeader("host", authority)

	srcIP, srcPort := retrieveAddressInfo("source")
	dstIP, dstPort := retrieveAddressInfo("destination")

	ctx.tx.ProcessConnection(srcIP, srcPort, dstIP, dstPort)

	uri, err := proxywasm.GetHttpRequestHeader(":path")
	if err != nil {
		return types.ActionContinue
	}

	method, err := proxywasm.GetHttpRequestHeader(":method")
	if err != nil {
		return types.ActionContinue
	}

	protocol, err := proxywasm.GetProperty([]string{"request", "protocol"})
	if err != nil {
		protocol = []byte("HTTP/2.0")
	}

	ctx.httpProtocol = string(protocol)
	ctx.tx.ProcessURI(uri, method, ctx.httpProtocol)

	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogErrorf("Failed to get request headers: %q ", err)
		return types.ActionContinue
	}

	for _, h := range hs {
		proxywasm.LogInfof("Add Header, header key: %s, value: %s", h[0], h[1])
		ctx.tx.AddRequestHeader(strings.ToLower(h[0]), h[1])
	}

	res := go_rules.ProcessRequestHeaderRules(&ctx.tx)
	if !res {
		_ = proxywasm.SendHttpResponse(403, nil, []byte("denied by waf"), -1)
	}

	return types.ActionContinue
}

func (ctx *httpContext) OnHttpRequestBody(bodySize int, endOfStream bool) types.Action {
	if ctx.tx.Variables.Interrupted {
		return types.ActionContinue
	}

	tx := ctx.tx

	if bodySize > 0 {
		b, err := proxywasm.GetHttpRequestBody(ctx.bodyReadIndex, bodySize)
		if err == nil {
			r, _, _ := tx.WriteRequestBody(b)
			if !r {
				proxywasm.LogError("Failed to write response body")
				return types.ActionContinue
			}
			ctx.bodyReadIndex += bodySize
		}
	}

	if endOfStream {
		ctx.bodyReadIndex = 0
		r, _ := tx.ProcessRequestBody()
		if !r {
			proxywasm.LogError("Failed to write process body")
		}
		res := go_rules.ProcessRequestBodyRules(&ctx.tx)
		if !res {
			_ = proxywasm.SendHttpResponse(403, nil, []byte("denied by waf"), -1)
		}

		return types.ActionContinue
	}

	return types.ActionPause
}

func (ctx *httpContext) OnHttpStreamDone() {
	_ = ctx.tx.Close()
}

// retrieveAddressInfo retrieves address properties from the proxy
// Expected targets are "source" or "destination"
// Envoy ref: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes#connection-attributes
func retrieveAddressInfo(target string) (string, int) {
	var targetIP, targetPortStr string
	var targetPort int
	targetAddressRaw, err := proxywasm.GetProperty([]string{target, "address"})
	if err != nil {
		proxywasm.LogErrorf("Failed to get %s address", target)
	} else {
		targetIP, targetPortStr, err = net.SplitHostPort(string(targetAddressRaw))
		if err != nil {
			proxywasm.LogErrorf("Failed to parse %s address", target)
		}
	}
	targetPortRaw, err := proxywasm.GetProperty([]string{target, "port"})
	if err == nil {
		targetPort, err = parsePort(targetPortRaw)
		if err != nil {
			proxywasm.LogErrorf("Failed to parse %s port", target)
		}
	} else if targetPortStr != "" {
		// If GetProperty fails we rely on the port inside the Address property
		// Mostly useful for proxies other than Envoy
		targetPort, err = strconv.Atoi(targetPortStr)
		if err != nil {
			proxywasm.LogErrorf("Failed to get %s port", target)

		}
	}
	return targetIP, targetPort
}

// parsePort converts port, retrieved as little-endian bytes, into int
func parsePort(b []byte) (int, error) {
	// Port attribute ({"source", "port"}) is populated as uint64 (8 byte)
	// Ref: https://github.com/envoyproxy/envoy/blob/1b3da361279a54956f01abba830fc5d3a5421828/source/common/network/utility.cc#L201
	if len(b) < 8 {
		return 0, errors.New("port bytes not found")
	}
	// 0 < Port number <= 65535, therefore the retrieved value should never exceed 16 bits
	// and correctly fit int (at least 32 bits in size)
	unsignedInt := binary.LittleEndian.Uint64(b)
	if unsignedInt > math.MaxInt32 {
		return 0, errors.New("port conversion error")
	}
	return int(unsignedInt), nil
}

// replaceResponseBodyWhenInterrupted address an interruption raised during phase 4.
// At this phase, response headers are already sent downstream, therefore an interruption
// can not change anymore the status code, but only tweak the response body
func replaceResponseBodyWhenInterrupted(logger debuglog.Logger, bodySize int) types.Action {
	// TODO(M4tteoP): Update response body interruption logic after https://github.com/corazawaf/coraza-proxy-wasm/issues/26
	// Currently returns a body filled with null bytes that replaces the sensitive data potentially leaked
	err := proxywasm.ReplaceHttpResponseBody(bytes.Repeat([]byte("\x00"), bodySize))
	if err != nil {
		logger.Error().Err(err).Msg("Failed to replace response body")
		return types.ActionContinue
	}
	logger.Warn().Msg("Response body intervention occurred: body replaced")
	return types.ActionContinue
}

// parseServerName parses :authority pseudo-header in order to retrieve the
// virtual host.
func parseServerName(authority string) string {
	host, _, err := net.SplitHostPort(authority)
	if err != nil {
		// missing port or bad format
		proxywasm.LogError("Failed to parse server name from authority")
		host = authority
	}
	return host
}
