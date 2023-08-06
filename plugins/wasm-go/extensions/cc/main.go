package main

import (
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tidwall/gjson"
	"net/http"
	"sliding_hist_limiter"
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

type CCConfig struct {
	headerRulesMap map[string]*CCRule
	cookieRulesMap map[string]*CCRule
}

type CCRule struct {
	subRules []sliding_hist_limiter.CCSubRule
}

func parseConfig(jsonData gjson.Result, config *CCConfig, log wrapper.Log) error {
	config.headerRulesMap = make(map[string]*CCRule)
	config.cookieRulesMap = make(map[string]*CCRule)

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

		var ccSubRules []sliding_hist_limiter.CCSubRule
		blockSeconds := int64(0)
		blockSecondsConf := rule.Get("block_seconds")
		if blockSecondsConf.Exists() {
			blockSeconds = blockSecondsConf.Int()
		}

		qps := rule.Get("qps")
		qpm := rule.Get("qpm")
		qpd := rule.Get("qpd")

		// 时间都按微妙级别
		if qps.Exists() {
			//period := int64(1000 * 1000)
			ccSubRules = append(ccSubRules, sliding_hist_limiter.GetSubRule("second", qps.Int(), blockSeconds))
			//maxPeriod = period
		}
		if qpm.Exists() {
			//period := int64(60 * 1000 * 1000)
			ccSubRules = append(ccSubRules, sliding_hist_limiter.GetSubRule("minute", qpm.Int(), blockSeconds))
			//maxPeriod = period
		}
		if qpd.Exists() {
			//period := int64(24 * 60 * 60 * 1000 * 1000)
			ccSubRules = append(ccSubRules, sliding_hist_limiter.GetSubRule("day", qpd.Int(), blockSeconds))
			//maxPeriod = period
		}

		rulesMap[key] = &CCRule{ccSubRules}
	}

	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config CCConfig, log wrapper.Log) types.Action {

	//ts := qps_limiter.GetTimestamps(time.Now())
	ts := time.Now().UnixMilli()

	for headKey, headRule := range config.headerRulesMap {
		head, err := proxywasm.GetHttpRequestHeader(headKey)
		if err != nil {
			//log.Error("Get head Error: " + err.Error())
			continue
		}

		//proxywasm.LogInfof("Headstr: %s", head)
		if !validateCC(head, headRule, ts, &config) {
			proxywasm.SendHttpResponse(403, nil, []byte("denied by cc"), -1)
			return types.ActionContinue
		}
	}

	cookieStr, err := proxywasm.GetHttpRequestHeader("cookie")
	if err != nil {
		//log.Error("Get cookie str Error: " + err.Error())
		return types.ActionContinue
	}

	//proxywasm.LogInfof("Cookiestr: %s", cookieStr)
	cookies, err := parseCookie(cookieStr)
	if err != nil {
		return types.ActionContinue
	}

	for cookieKey, cookieRule := range config.cookieRulesMap {
		cookie, ok := cookies[cookieKey]
		if !ok {
			//log.Error("Get cookie Error: " + err.Error())
			continue
		}

		if !validateCC(cookie, cookieRule, ts, &config) {
			proxywasm.SendHttpResponse(403, nil, []byte("denied by cc"), -1)
			return types.ActionContinue
		}
	}

	return types.ActionContinue
}

//func onTick(config *CCConfig, log wrapper.Log) {
//	log.Info("tik tok")
//}

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

//func validateCC(key string, rule *CCRule, ts *qps_limiter.Timestamps, config *CCConfig) bool {
//
//	id := qps_limiter.Identifier(key)
//
//	counters, stop, err := qps_limiter.GetUsage(rule.subRules, id, ts)
//	if err != nil {
//		proxywasm.LogErrorf("Failed to get usage: %v", err)
//	}
//
//	if counters != nil {
//		if stop != "" {
//			return false
//		}
//
//		qps_limiter.LocalPolicyIncrement(id, counters, ts)
//	}
//
//	return true
//}

func validateCC(key string, rule *CCRule, ts int64, config *CCConfig) bool {
	id := sliding_hist_limiter.Identifier(key)

	rateLimitData, stop, err := sliding_hist_limiter.GetData(rule.subRules, id, ts)
	if err != nil {
		proxywasm.LogErrorf("Failed to get rate limit data: %v", err)
	}

	if stop {
		return false
	}

	sliding_hist_limiter.SaveData(id, rateLimitData, ts)
	return true
}
