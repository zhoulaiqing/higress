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
	ipIntervalRootV4 *ip_tools.IPIntervalNode
	ipIntervalRootV6 *ip_tools.IPIntervalNode
}

func parseConfig(json gjson.Result, config *IpBlacklistConfig, log wrapper.Log) error {
	ipBlackArray := json.Get("ip_blacklist").Array()
	for _, blackIp := range ipBlackArray {
		isV4, minIp, maxIp, _ := ip_tools.GetIPIntRange(blackIp.String())
		if isV4 {
			config.ipIntervalRootV4 = ip_tools.IPIntervalTreeInsert(config.ipIntervalRootV4, ip_tools.IPInterval{Start: minIp, End: maxIp})
		} else {
			config.ipIntervalRootV6 = ip_tools.IPIntervalTreeInsert(config.ipIntervalRootV6, ip_tools.IPInterval{Start: minIp, End: maxIp})
		}

	}

	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config IpBlacklistConfig, log wrapper.Log) types.Action {

	ip, err := proxywasm.GetHttpRequestHeader("x-real-ip")

	if err != nil {
		//log.Error("Failed to get IP Header")
		return types.ActionContinue
	}

	isV4, ipInt, _, _ := ip_tools.GetIPIntRange(ip)

	var intervalRoot *ip_tools.IPIntervalNode
	if isV4 {
		intervalRoot = config.ipIntervalRootV4
	} else {
		intervalRoot = config.ipIntervalRootV6
	}

	if ip_tools.IPSearch(intervalRoot, ipInt) != nil {
		//log.Error("Hit ip blacklist rule")
		proxywasm.SendHttpResponse(403, nil, []byte("denied by ip"), -1)

		return types.ActionContinue
	}

	return types.ActionContinue
}
