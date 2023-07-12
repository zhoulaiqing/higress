package main

import (
	"fmt"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println("hello world")

	// ===================== For Test ============================
	var ipIntervalRoot *Node
	minIp, maxIp, _ := IPRange("1.1.1.1")
	ipIntervalRoot = Insert(ipIntervalRoot, Interval{minIp, maxIp})

	minIp2, maxIp2, _ := IPRange("2.2.2.2/16")
	ipIntervalRoot = Insert(ipIntervalRoot, Interval{minIp2, maxIp2})

	res1 := Search(ipIntervalRoot, minIp) != nil
	fmt.Println("Contains 1.1.1.1", res1)

	otherIp, _, _ := IPRange("10.23.2.2")
	res2 := Search(ipIntervalRoot, otherIp) != nil
	fmt.Println("Contains 10.23.2.2", res2)

	ip3, _, _ := IPRange("2.2.1.1")
	res3 := Search(ipIntervalRoot, ip3) != nil
	fmt.Println("Contains 2.2.1.1", res3)

	// ===================== For Test ============================

	//wrapper.SetCtx(
	//	"ip_deny",
	//	wrapper.ParseConfigBy(parseConfig),
	//	wrapper.ProcessRequestHeadersBy(onHttpRequestHeaders),
	//)
}

type IpBlacklistConfig struct {
	ipIntervalRoot *Node
}

func parseConfig(json gjson.Result, config *IpBlacklistConfig, log wrapper.Log) error {
	ipBlackArray := json.Get("ip_blacklist").Array()
	for _, blackIp := range ipBlackArray {
		minIp, maxIp, _ := IPRange(blackIp.String())
		config.ipIntervalRoot = Insert(config.ipIntervalRoot, Interval{minIp, maxIp})
	}

	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config IpBlacklistConfig, log wrapper.Log) types.Action {

	ip, err := proxywasm.GetHttpRequestHeader("x-real-ip")

	if err != nil {
		log.Error("Failed to get IP Header")
		return types.ActionContinue
	}

	ipInt, _, _ := IPRange(ip)

	if Search(config.ipIntervalRoot, ipInt) != nil {
		log.Error("Hit ip blacklist rule")
		proxywasm.SendHttpResponse(403, nil, []byte("denied by ip"), -1)

		return types.ActionContinue
	}

	return types.ActionContinue
}
