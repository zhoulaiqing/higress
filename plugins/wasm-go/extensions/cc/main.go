package main

import (
	"cc_tools"
	"fmt"
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
		//wrapper.ProcessTickBy(onTick),
	)
}

type Timestamps map[string]int64

func getTimestamps(t time.Time) *Timestamps {
	ts := Timestamps{}

	ye, mo, da := t.Year(), t.Month(), t.Day()
	ho, mi, se, lo := t.Hour(), t.Minute(), t.Second(), t.Location()

	ts["now"] = t.Unix()
	ts["second"] = time.Date(ye, mo, da, ho, mi, se, 0, lo).Unix()
	ts["minute"] = time.Date(ye, mo, da, ho, mi, 0, 0, lo).Unix()
	ts["day"] = time.Date(ye, mo, da, 0, 0, 0, 0, lo).Unix()

	return &ts
}

type CCConfig struct {
	headerRulesMap map[string]*CCRule
	cookieRulesMap map[string]*CCRule
}

type CCRule struct {
	subRules []CCSubRule
}

type CCSubRule struct {
	limitCount   int64
	period       string
	blockSeconds int64
}

type Usage struct {
	limit     int64
	remaining int64
	usage     int64
	blockUtil int64
	cas       uint32
}

func parseConfig(jsonData gjson.Result, config *CCConfig, log wrapper.Log) error {
	log.Info("Start parsing config......")
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

		var ccSubRules []CCSubRule
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
			ccSubRules = append(ccSubRules, CCSubRule{qps.Int(), "second", blockSeconds})
			//maxPeriod = period
		}
		if qpm.Exists() {
			//period := int64(60 * 1000 * 1000)
			ccSubRules = append(ccSubRules, CCSubRule{qpm.Int(), "minute", blockSeconds})
			//maxPeriod = period
		}
		if qpd.Exists() {
			//period := int64(24 * 60 * 60 * 1000 * 1000)
			ccSubRules = append(ccSubRules, CCSubRule{qpd.Int(), "day", blockSeconds})
			//maxPeriod = period
		}

		rulesMap[key] = &CCRule{ccSubRules}
	}

	return nil
}

func onHttpRequestHeaders(ctx wrapper.HttpContext, config CCConfig, log wrapper.Log) types.Action {

	ts := getTimestamps(time.Now())

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

	for cookieKey, cookieRule := range config.headerRulesMap {
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

type Identifier string

// 返回 periodTime used, blockUntil, cas, err
func localPolicyUsage(id Identifier, period string) (int64, int64, int64, uint32, error) {
	cacheKey := getLocalKey(id, period)

	value, cas, err := proxywasm.GetSharedData(cacheKey)
	if err != nil {
		if err == types.ErrorStatusNotFound {
			return 0, 0, 0, 0, nil
		}
		return 0, 0, 0, 0, nil
	}

	ret := cc_tools.ByteArrayToInt64Array(value)
	if len(ret) != 3 {
		proxywasm.LogError("Shared data format error")
	}

	//proxywasm.LogInfof("Get shared data: %v:  %v, %v, %v", cacheKey, ret[0], ret[1], ret[2])

	return ret[0], ret[1], ret[2], cas, nil
}

func localPolicyIncrement(id Identifier, counters map[string]Usage, ts *Timestamps) {
	for period, usage := range counters {
		cacheKey := getLocalKey(id, period)

		value := usage.usage
		blockUntil := usage.blockUtil
		cas := usage.cas

		saved := false
		var err error
		// cas 保存，重试 10 次
		for i := 0; i < 10; i++ {
			// 保存顺序： time, usage, blockUntil
			buf := cc_tools.Int64ArrayToByteArray([]int64{(*ts)[period], value + 1, blockUntil})

			err = proxywasm.SetSharedData(cacheKey, buf, cas)
			if err == nil {
				saved = true
				break
			} else if err == types.ErrorStatusCasMismatch {
				// cas 不匹配，重试
				buf, cas, err = proxywasm.GetSharedData(cacheKey)
				ret := cc_tools.ByteArrayToInt64Array(buf)
				value = ret[1]
			} else {
				break
			}
		}

		if !saved {
			proxywasm.LogErrorf("Could not increment counter for period '%v': %v", period, err)
		}
	}
}

func getUsage(subRules []CCSubRule, id Identifier, ts *Timestamps) (map[string]Usage, string, error) {
	counters := make(map[string]Usage)
	stop := ""

	for _, subRule := range subRules {
		if subRule.limitCount < 0 {
			continue
		}
		period := subRule.period
		savedTime, curUsage, blockUntil, cas, err := localPolicyUsage(id, period)
		if err != nil {
			return counters, period, err
		}

		if blockUntil >= (*ts)["now"] {
			stop = period
		}

		curTime := (*ts)[period]
		var usage Usage
		if curTime != savedTime {
			// 需要刷新
			usage = Usage{
				limit:     subRule.limitCount,
				remaining: subRule.limitCount,
				usage:     0,
				cas:       cas,
				blockUtil: 0,
			}
		} else {
			remaining := subRule.limitCount - curUsage
			usage = Usage{
				limit:     subRule.limitCount,
				remaining: remaining,
				usage:     curUsage,
				cas:       cas,
				blockUtil: blockUntil,
			}

			if remaining <= 0 {
				stop = period
				if subRule.blockSeconds > 0 {
					usage.blockUtil = (*ts)["now"] + subRule.blockSeconds
				}
			}
		}

		counters[period] = usage
	}

	return counters, stop, nil
}

func getLocalKey(id Identifier, period string) string {
	return fmt.Sprintf("limit:%v:%v", id, period)
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

func validateCC(key string, rule *CCRule, ts *Timestamps, config *CCConfig) bool {

	id := Identifier(key)

	counters, stop, err := getUsage(rule.subRules, id, ts)
	if err != nil {
		proxywasm.LogErrorf("Failed to get usage: %v", err)
	}

	if counters != nil {
		if stop != "" {
			return false
		}

		localPolicyIncrement(id, counters, ts)
	}

	return true
}
