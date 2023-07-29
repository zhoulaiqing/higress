package main

import (
	"cc_tools"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tidwall/gjson"
	"net/http"
	"time"
)

func main() {
	//fmt.Println("Hello world.")

	wrapper.SetCtx(
		"cc_deny",
		wrapper.ParseConfigBy(parseConfig),
		wrapper.ProcessRequestHeadersBy(onHttpRequestHeaders),
	)
}

var RULE_DATA map[string]*cc_tools.ZSet
var BLOCK_DATA map[string]int64

type CCConfig struct {
	headerRulesMap map[string]*CCRule
	cookieRulesMap map[string]*CCRule
}

type CCRule struct {
	maxPeriod int
	subRules  []CCSubRule
}

type CCSubRule struct {
	limitCount   int
	limitPeriod  int
	blockSeconds int
}

func parseConfig(json gjson.Result, config *CCConfig, log wrapper.Log) error {
	config.headerRulesMap = make(map[string]*CCRule)
	config.cookieRulesMap = make(map[string]*CCRule)
	RULE_DATA = make(map[string]*cc_tools.ZSet)

	rulesArray := json.Get("cc_rules").Array()
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
		var maxPeriod int
		blockSeconds := 0
		blockSecondsConf := rule.Get("block_seconds")
		if blockSecondsConf.Exists() {
			blockSeconds = int(blockSecondsConf.Int())
		}

		qps := rule.Get("qps")
		qpm := rule.Get("qpm")
		qpd := rule.Get("qpd")

		if qps.Exists() {
			period := 1000
			ccSubRules = append(ccSubRules, CCSubRule{int(qps.Int()), period, blockSeconds})
			maxPeriod = period
		}
		if qpm.Exists() {
			period := 60 * 1000
			ccSubRules = append(ccSubRules, CCSubRule{int(qpm.Int()), period, blockSeconds})
			maxPeriod = period
		}
		if qpd.Exists() {
			period := 24 * 60 * 60 * 1000
			ccSubRules = append(ccSubRules, CCSubRule{int(qpd.Int()), period, blockSeconds})
			maxPeriod = period
		}

		rulesMap[key] = &CCRule{maxPeriod, ccSubRules}
	}

	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config CCConfig, log wrapper.Log) types.Action {

	timestamp := time.Now().UnixNano()
	for headKey, headRule := range config.headerRulesMap {
		head, err := proxywasm.GetHttpRequestHeader(headKey)
		if err != nil {
			continue
		}

		if !validateCC(head, headRule, timestamp) {
			proxywasm.SendHttpResponse(403, nil, []byte("denied by cc"), -1)
			return types.ActionContinue
		}
	}

	cookieStr, err := proxywasm.GetHttpRequestHeader("cookie")
	if err != nil {
		return types.ActionContinue
	}

	cookies, err := parseCookie(cookieStr)
	if err != nil {
		return types.ActionContinue
	}

	for cookieKey, cookieRule := range config.headerRulesMap {
		cookie, ok := cookies[cookieKey]
		if !ok {
			continue
		}

		if !validateCC(cookie, cookieRule, timestamp) {
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

func validateCC(key string, rule *CCRule, timestamp int64) bool {
	blockHist, blocked := BLOCK_DATA[key]
	if blocked && blockHist >= timestamp {
		return false
	}
	delete(BLOCK_DATA, key)

	hist, ok := RULE_DATA[key]

	if ok {
		// 1. Clear obsolete data
		hist.ZRemByScore(key, 0, timestamp-int64(rule.maxPeriod))
		// 2. Validate rules
		for _, subRule := range rule.subRules {
			histCount := hist.ZCount(key, timestamp-int64(subRule.limitPeriod), timestamp)
			if histCount >= subRule.limitCount {
				if subRule.blockSeconds > 0 {
					BLOCK_DATA[key] = timestamp + int64(subRule.blockSeconds)
				}
				return false
			}
		}
		// 3. Update data
		hist.ZAdd(key, timestamp)
	}
	return true
}
