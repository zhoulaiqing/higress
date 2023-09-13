package rule_941

import (
	"fmt"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
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

func (r *Rule941) EvaluatePhase1(tx *core.Transaction) int {
	_, ok := tx.Variables.RequestHeaders["cookie"]
	if ok {
		for k, v := range tx.Variables.RequestCookies {
			if strings.Contains(k, "__utm") {
				continue
			}

			if r.matchValue(k, false) {
				return r.block(tx)
			}

			if r.matchValue(v, false) {
				return r.block(tx)
			}
		}
	}
	//fmt.Printf("RequestUriRaw: %s \n", tx.Variables.RequestUriRaw)
	if r.matchValue(tx.Variables.RequestUriRaw, false) {
		return r.block(tx)
	}

	ua := tx.Variables.RequestHeaders["user-agent"]
	//proxywasm.LogInfof("user agent: %s", ua)
	if r.matchValue(ua, true) {
		return r.block(tx)
	}

	return rule_tasks.PASS
}

func (r *Rule941) evaluateAll(tx *core.Transaction) int {

	for k, v := range tx.Variables.ArgsPost {
		//proxywasm.LogInfof("arg key: %s", k)

		if r.matchValue(k, false) {
			return r.block(tx)
		}

		//proxywasm.LogInfof("arg value: %s", v)
		if r.matchValue(v, false) {
			return r.block(tx)
		}
	}

	for _, v := range tx.Variables.XML["/*"] {

		if r.matchValue(v, false) {
			return r.block(tx)
		}
	}

	if !tx.Variables.Skip941ForFileName {
		if r.matchValue(tx.Variables.RequestFileName, false) {
			return r.block(tx)
		}
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
	// default transform 941110 - 941310
	v1 := r.transformDefault(value)

	if matchXss(v1) {
		fmt.Println("Match ragel machines")
		return true
	}

	// pure value e.g. 941360
	//fmt.Println("941360: " + value)
	//if m := rule_tasks.Re941360.MatchString(value); m {
	//	proxywasm.LogInfof("Match 941360")
	//	return true
	//}

	if m, _ := core.PmEvaluate(rule_tasks.Rule941180Matcher, v1, false); m {
		return true
	}

	// 941350
	//v2 := r.transform350(value)
	////fmt.Println("941350: " + v2)
	//if m := rule_tasks.Re941350.MatchString(v2); m {
	//	fmt.Println("Match 941350")
	//	return true
	//}
	//
	//// 941370, 941400
	//v3 := r.transform370And400(value)
	////fmt.Println("941370, 941400: " + v3)
	//if m := rule_tasks.Re941370.MatchString(v3); m {
	//	fmt.Println("Match 941370")
	//	return true
	//}
	//if m := rule_tasks.Re941400.MatchString(v3); m {
	//	fmt.Println("Match 941400")
	//	return true
	//}
	//
	//// 941390
	//v4 := r.transform390(value)
	////fmt.Println("941390: " + v4)
	//if m := rule_tasks.Re941390.MatchString(v4); m {
	//	fmt.Println("Match 941390")
	//	return true
	//}

	return false
}

func (r *Rule941) transformDefault(value string) string {
	v := value

	v, _, _ = core.Utf8ToUnicode(v)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.HtmlEntityDecode(v)
	v, _, _ = core.JsDecode(v)
	v, _, _ = core.CssDecode(v)
	v, _, _ = core.RemoveNulls(v)

	v, _, _ = core.Utf8ToUnicode(v)
	v, _, _ = core.UrlDecodeUni(v)

	return strings.ToLower(v) + " "
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
