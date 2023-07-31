package main

import (
	"cc_tools"
	"encoding/json"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tidwall/gjson"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//fmt.Println("Hello world.")

	wrapper.SetCtx(
		"cc_deny",
		wrapper.ParseConfigBy(parseConfig),
		wrapper.ProcessRequestHeadersBy(onHttpRequestHeaders),
		//wrapper.ProcessTickBy(onTick),
	)
}

//var mu sync.Mutex
//var HIST_DATA *cc_tools.ZSet
//var BLOCK_DATA map[string]int64

type CCConfig struct {
	headerRulesMap map[string]*CCRule
	cookieRulesMap map[string]*CCRule
	HIST_DATA      *cc_tools.ZSet
	BLOCK_DATA     map[string]int64
}

type CCRule struct {
	maxPeriod int64
	subRules  []CCSubRule
}

type CCSubRule struct {
	limitCount   int
	limitPeriod  int64
	blockSeconds int
}

type BlockData struct {
	blockUntil int
}

func parseConfig(jsonData gjson.Result, config *CCConfig, log wrapper.Log) error {
	log.Info("Start parsing config......")
	config.headerRulesMap = make(map[string]*CCRule)
	config.cookieRulesMap = make(map[string]*CCRule)
	//blockMap := make(map[string]BlockData)
	//blockBytes, blockErr := json.Marshal(blockMap)
	//if blockErr != nil {
	//	log.Error("Json Marshal Error: " + blockErr.Error())
	//}
	//
	//setErr := proxywasm.SetSharedData("BLOCK_DATA", blockBytes, uint32(time.Now().UnixNano()/1000))
	//if setErr != nil {
	//	log.Error("Set blocked data error: " + setErr.Error())
	//}
	config.HIST_DATA = &cc_tools.ZSet{}
	config.BLOCK_DATA = make(map[string]int64)

	rulesArray := jsonData.Get("cc_rules").Array()
	for _, rule := range rulesArray {
		header := rule.Get("header")
		cookie := rule.Get("cookie")

		// header 和 cookie 都不存在，无效的配置项
		if !header.Exists() && !cookie.Exists() {
			continue
		}

		var rulesMap map[string]*CCRule
		var key string
		if header.Exists() {
			rulesMap = config.headerRulesMap
			key = header.String()
		} else {
			rulesMap = config.cookieRulesMap
			key = cookie.String()
		}

		var ccSubRules []CCSubRule
		var maxPeriod int64
		blockSeconds := 0
		blockSecondsConf := rule.Get("block_seconds")
		if blockSecondsConf.Exists() {
			blockSeconds = int(blockSecondsConf.Int())
		}

		qps := rule.Get("qps")
		qpm := rule.Get("qpm")
		qpd := rule.Get("qpd")

		// 时间都按微妙级别
		if qps.Exists() {
			period := int64(1000 * 1000)
			ccSubRules = append(ccSubRules, CCSubRule{int(qps.Int()), period, blockSeconds})
			maxPeriod = period
		}
		if qpm.Exists() {
			period := int64(60 * 1000 * 1000)
			ccSubRules = append(ccSubRules, CCSubRule{int(qpm.Int()), period, blockSeconds})
			maxPeriod = period
		}
		if qpd.Exists() {
			period := int64(24 * 60 * 60 * 1000 * 1000)
			ccSubRules = append(ccSubRules, CCSubRule{int(qpd.Int()), period, blockSeconds})
			maxPeriod = period
		}

		rulesMap[key] = &CCRule{maxPeriod, ccSubRules}
	}

	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config CCConfig, log wrapper.Log) types.Action {

	timestamp := time.Now().UnixNano() / 1000
	for headKey, headRule := range config.headerRulesMap {
		head, err := proxywasm.GetHttpRequestHeader(headKey)
		if err != nil {
			log.Error("Get head Error: " + err.Error())
			continue
		}

		if !validateCC(head, headRule, timestamp, &config) {
			proxywasm.SendHttpResponse(403, nil, []byte("denied by cc"), -1)
			return types.ActionContinue
		}
	}

	cookieStr, err := proxywasm.GetHttpRequestHeader("cookie")
	if err != nil {
		log.Error("Get cookie str Error: " + err.Error())
		return types.ActionContinue
	}

	cookies, err := parseCookie(cookieStr)
	if err != nil {
		return types.ActionContinue
	}

	for cookieKey, cookieRule := range config.headerRulesMap {
		cookie, ok := cookies[cookieKey]
		if !ok {
			log.Error("Get cookie Error: " + err.Error())
			continue
		}

		if !validateCC(cookie, cookieRule, timestamp, &config) {
			proxywasm.SendHttpResponse(403, nil, []byte("denied by cc"), -1)
			return types.ActionContinue
		}
	}

	return types.ActionContinue
}

func parseCookie(cookieStr string) (map[string]string, error) {
	cookies := make(map[string]string)

	// Parse the cookie string using http.ParseCookie from Go standard library
	cookieHeader := http.Header{}
	cookieHeader.Add("Cookie", cookieStr)
	request := http.Request{Header: cookieHeader}

	parsedCookies := request.Cookies()
	for _, cookie := range parsedCookies {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies, nil
}

func validateCC(key string, rule *CCRule, timestamp int64, config *CCConfig) bool {
	blockValue, _, _ := proxywasm.GetSharedData("BLOCK_DATA")
	var blockData map[string]BlockData
	json.Unmarshal(blockValue, &blockData)
	blockHist, blocked := blockData[key]
	if blocked && int64(blockHist.blockUntil) >= timestamp {
		return false
	}
	//delete(BLOCK_DATA, key)

	// Update hist data
	HIST_DATA := config.HIST_DATA
	HIST_DATA.ZAdd(key, timestamp)
	proxywasm.LogInfo("BEFORE REMOVE: " + HIST_DATA.ToString(key))
	// Clear obsolete data
	proxywasm.LogInfo("MaxPeriod: " + strconv.FormatInt(rule.maxPeriod, 10))
	HIST_DATA.ZRemByScore(key, 0, timestamp-rule.maxPeriod)
	proxywasm.LogInfo("AFTER REMOVE: " + HIST_DATA.ToString(key))
	// Validate rules
	for _, subRule := range rule.subRules {
		histCount := HIST_DATA.ZCount(key, timestamp-subRule.limitPeriod, timestamp)
		proxywasm.LogInfo("Hist count: key " + key + ", count " + strconv.Itoa(histCount))
		if histCount >= subRule.limitCount {
			if subRule.blockSeconds > 0 {
				blockData[key] = BlockData{int(timestamp + int64(subRule.blockSeconds*1000*1000))}
				blockJson, _ := json.Marshal(blockData)
				proxywasm.SetSharedData("BLOCK_DATA", blockJson, uint32(timestamp))
			}

			return false
		}
	}

	return true
}
