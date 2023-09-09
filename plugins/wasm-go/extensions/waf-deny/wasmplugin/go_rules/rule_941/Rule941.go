package rule_941

import (
	"fmt"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/corazawaf/libinjection-go"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/wasilibs/go-re2"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type Rule941 struct {
}

func (r *Rule941) Id() string {
	return "941"
}

func (r *Rule941) Phase() int {
	return 2
}

func (r *Rule941) Evaluate(tx *core.Transaction) int {
	return r.evaluateAll(tx)
}

func (r *Rule941) evaluateAll(tx *core.Transaction) int {

	cookieLiteral, ok := tx.Variables.RequestHeaders["cookie"]

	//proxywasm.LogInfof("cookie literal: %s", cookieLiteral)
	if ok {
		if r.matchValue(cookieLiteral, false) {
			return r.block(tx)
		}

		for k, v := range tx.Variables.RequestCookies {
			if strings.Contains(k, "__utm") {
				continue
			}

			//proxywasm.LogInfof("cookie key: %s", k)

			if r.matchValue(k, false) {
				return r.block(tx)
			}

			//proxywasm.LogInfof("cookie value: %s", v)

			if r.matchValue(v, false) {
				return r.block(tx)
			}
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			//proxywasm.LogInfof("arg key: %s", k)

			if r.matchValue(k, false) {
				return r.block(tx)
			}

			//proxywasm.LogInfof("arg value: %s", v)
			if r.matchValue(v, false) {
				return r.block(tx)
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {

		//proxywasm.LogInfof("xml value: %s", v)

		if r.matchValue(v, false) {
			return r.block(tx)
		}
	}

	if !tx.Variables.Skip941ForFileName {

		//proxywasm.LogInfof("file name: %s", tx.Variables.RequestFileName)

		if r.matchValue(tx.Variables.RequestFileName, false) {
			return r.block(tx)
		}
	}

	referer := tx.Variables.RequestHeaders["referer"]

	//proxywasm.LogInfof("referer: %s", referer)
	if r.matchValue(referer, true) {
		return r.block(tx)
	}

	ua := tx.Variables.RequestHeaders["user-agent"]
	//proxywasm.LogInfof("user agent: %s", ua)
	if r.matchValue(ua, true) {
		return r.block(tx)
	}

	return rule_tasks.PASS
}

var headerRuleIds = []int{941110, 941130, 941140, 941160, 941170}

func (r *Rule941) block(tx *core.Transaction) int {
	tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
	tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE

	return rule_tasks.BLOCK
}

func (r *Rule941) matchValue(value string, isHeader bool) bool {

	if !isHeader && libinjection.IsXSS(value) {
		return true
	}

	// pure value e.g. 941360
	//fmt.Println("941360: " + value)
	if m := rule_tasks.Re941360.MatchString(value); m {
		proxywasm.LogInfof("Match 941360")
		return true
	}

	// default transform 941110 - 941310
	var Re941ForCacheMap = map[int]*re2.Regexp{
		//941110:  rule_tasks.Re941110,
		//941130:  rule_tasks.Re941130,
		//941140:  rule_tasks.Re941140,
		//941160:  rule_tasks.Re941160,
		941170:  rule_tasks.Re941170,
		941190:  rule_tasks.Re941190,
		941200:  rule_tasks.Re941200,
		941210:  rule_tasks.Re941210,
		941220:  rule_tasks.Re941220,
		941230:  rule_tasks.Re941230,
		941240:  rule_tasks.Re941240,
		941250:  rule_tasks.Re941250,
		941260:  rule_tasks.Re941260,
		941270:  rule_tasks.Re941270,
		941280:  rule_tasks.Re941280,
		941290:  rule_tasks.Re941290,
		941300:  rule_tasks.Re941300,
		9413101: rule_tasks.Re9413101,
		9413102: rule_tasks.Re9413102,
	}

	v1 := r.transformDefault(value)

	//if matchXSS([]byte(v1)) {
	//	//proxywasm.LogInfo("Ragel Match")
	//	fmt.Println("Ragel Match")
	//	return true
	//}

	for k, re := range Re941ForCacheMap {
		//fmt.Println(strconv.Itoa(k) + ": " + v1)
		if isHeader && !slices.Contains(headerRuleIds, k) {
			continue
		}

		//fmt.Println(re)
		a := re.FindStringSubmatch(v1)
		m := re.MatchString(v1)
		//if m := re.MatchString(v1); m {
		if len(a) > 0 || m {
			fmt.Printf("Match %s \n", strconv.Itoa(k))
			return true
		}
	}

	if m, _ := core.PmEvaluate(rule_tasks.Rule941180Matcher, v1, false); m {
		return true
	}

	// 941350
	v2 := r.transform350(value)
	//fmt.Println("941350: " + v2)
	if m := rule_tasks.Re941350.MatchString(v2); m {
		fmt.Println("Match 941350")
		return true
	}

	// 941370, 941400
	v3 := r.transform370And400(value)
	//fmt.Println("941370, 941400: " + v3)
	if m := rule_tasks.Re941370.MatchString(v3); m {
		fmt.Println("Match 941370")
		return true
	}
	if m := rule_tasks.Re941400.MatchString(v3); m {
		fmt.Println("Match 941400")
		return true
	}

	// 941390
	v4 := r.transform390(value)
	//fmt.Println("941390: " + v4)
	if m := rule_tasks.Re941390.MatchString(v4); m {
		fmt.Println("Match 941390")
		return true
	}

	return false
}

func (r *Rule941) transformDefault(value string) string {
	v, _, _ := core.Utf8ToUnicode(value)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.HtmlEntityDecode(v)
	v, _, _ = core.JsDecode(v)
	v, _, _ = core.CssDecode(v)
	v, _, _ = core.RemoveNulls(v)

	return v
}

func (r *Rule941) transform350(value string) string {
	v, _, _ := core.UrlDecode(value)
	v, _, _ = core.HtmlEntityDecode(v)
	v, _, _ = core.JsDecode(v)

	return v
}

func (r *Rule941) transform370And400(value string) string {
	v, _, _ := core.UrlDecodeUni(value)
	v, _, _ = core.CompressWhitespace(v)

	return v
}

func (r *Rule941) transform390(value string) string {
	v, _, _ := core.HtmlEntityDecode(value)
	v, _, _ = core.JsDecode(v)

	return v
}
