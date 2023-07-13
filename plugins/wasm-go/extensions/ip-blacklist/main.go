package main

import (
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tidwall/gjson"
	"ip-tools"
)

func main() {
	// ===================== For Test ============================
	//TestIpDeny(nil)
	// ===================== For Test ============================

	wrapper.SetCtx(
		"ip_deny",
		wrapper.ParseConfigBy(parseConfig),
		wrapper.ProcessRequestHeadersBy(onHttpRequestHeaders),
	)
}

type IpBlacklistConfig struct {
	ipIntervalRoot *ip_tools.IPIntervalNode
}

func parseConfig(json gjson.Result, config *IpBlacklistConfig, log wrapper.Log) error {
	ipBlackArray := json.Get("ip_blacklist").Array()
	for _, blackIp := range ipBlackArray {
		minIp, maxIp, _ := ip_tools.GetIPIntRange(blackIp.String())
		config.ipIntervalRoot = ip_tools.IPIntervalTreeInsert(config.ipIntervalRoot, ip_tools.IPInterval{minIp, maxIp})
	}

	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config IpBlacklistConfig, log wrapper.Log) types.Action {

	ip, err := proxywasm.GetHttpRequestHeader("x-real-ip")

	if err != nil {
		log.Error("Failed to get IP Header")
		return types.ActionContinue
	}

	ipInt, _, _ := ip_tools.GetIPIntRange(ip)

	if ip_tools.IPSearch(config.ipIntervalRoot, ipInt) != nil {
		log.Error("Hit ip blacklist rule")
		proxywasm.SendHttpResponse(403, nil, []byte("denied by ip"), -1)

		return types.ActionContinue
	}

	return types.ActionContinue
}
