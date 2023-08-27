// Copyright The OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package wasmplugin

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza/v3/debuglog"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"math"
	"net"
	"strconv"
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
	metrics               *wafMetrics
	interruptedAt         interruptionPhase
	logger                debuglog.Logger
	metricLabelsKV        []string
}

func (ctx *httpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	ctx.tx = core.NewTransaction()
	authority, err := proxywasm.GetHttpRequestHeader(":authority")
	if err != nil {
		proxywasm.LogWarnf("Failed to resolve WAF for authority %q ", authority)
		return types.ActionContinue
	}

	ctx.tx.AddRequestHeader("Host", authority)
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
		ctx.logger.Error().Err(err).Msg("Failed to get request headers")
		return types.ActionContinue
	}

	for _, h := range hs {
		ctx.tx.AddRequestHeader(h[0], h[1])
	}

	res := go_rules.ProcessRequestHeaderRules(&ctx.tx)
	if !res {
		proxywasm.SendHttpResponse(403, nil, []byte("denied by waf"), -1)
	}

	return types.ActionContinue
}

// retrieveAddressInfo retrieves address properties from the proxy
// Expected targets are "source" or "destination"
// Envoy ref: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes#connection-attributes
func retrieveAddressInfo(logger debuglog.Logger, target string) (string, int) {
	var targetIP, targetPortStr string
	var targetPort int
	targetAddressRaw, err := proxywasm.GetProperty([]string{target, "address"})
	if err != nil {
		logger.Debug().
			Err(err).
			Msg(fmt.Sprintf("Failed to get %s address", target))
	} else {
		targetIP, targetPortStr, err = net.SplitHostPort(string(targetAddressRaw))
		if err != nil {
			logger.Debug().
				Err(err).
				Msg(fmt.Sprintf("Failed to parse %s address", target))
		}
	}
	targetPortRaw, err := proxywasm.GetProperty([]string{target, "port"})
	if err == nil {
		targetPort, err = parsePort(targetPortRaw)
		if err != nil {
			logger.Debug().
				Err(err).
				Msg(fmt.Sprintf("Failed to parse %s port", target))
		}
	} else if targetPortStr != "" {
		// If GetProperty fails we rely on the port inside the Address property
		// Mostly useful for proxies other than Envoy
		targetPort, err = strconv.Atoi(targetPortStr)
		if err != nil {
			logger.Debug().
				Err(err).
				Msg(fmt.Sprintf("Failed to get %s port", target))

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
func parseServerName(logger debuglog.Logger, authority string) string {
	host, _, err := net.SplitHostPort(authority)
	if err != nil {
		// missing port or bad format
		logger.Debug().
			Err(err).
			Msg("Failed to parse server name from authority")
		host = authority
	}
	return host
}
