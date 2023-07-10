package main

import (
	"fmt"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slices"
)

func main() {
	fmt.Printf("hello world")
	ip := "10.1.1.1"
	minIP, maxIP, err := IPRange(ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Min IP:", intToIP(minIP).String()) // 输出: 10.1.1.1
	fmt.Println("Max IP:", intToIP(maxIP).String()) // 输出: 10.1.1.1

	ipWithMask := "10.1.1.1/16"
	minIP, maxIP, err = IPRange(ipWithMask)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Min IP:", intToIP(minIP).String()) // 输出: 10.1.0.0
	fmt.Println("Max IP:", intToIP(maxIP).String()) // 输出: 10.1.255.255

	//wrapper.SetCtx(
	//	"ip_deny",
	//	wrapper.ParseConfigBy(parseConfig),
	//	wrapper.ProcessRequestHeadersBy(onHttpRequestHeaders),
	//)
}

type IpBlacklistConfig struct {
	ipBlackList []string
	ipBlackBit  int32
}

func parseConfig(json gjson.Result, config *IpBlacklistConfig, log wrapper.Log) error {
	ipBlackArray := json.Get("ip_blacklist").Array()
	ipBlackList := make([]string, len(ipBlackArray))
	for _, blackIp := range ipBlackArray {
		ipBlackList = append(ipBlackList, blackIp.String())
	}

	config.ipBlackList = ipBlackList
	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config IpBlacklistConfig, log wrapper.Log) types.Action {

	ip, err := proxywasm.GetHttpRequestHeader("x-real-ip")

	if err != nil {
		log.Error("Failed to get IP Header")
		return types.ActionContinue
	}

	if slices.Contains(config.ipBlackList, ip) {
		log.Error("Hit ip blacklist rule")
		proxywasm.SendHttpResponse(403, nil, []byte("denied by ip"), -1)

		return types.ActionContinue
	}

	return types.ActionContinue
}
