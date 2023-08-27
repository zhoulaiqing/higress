package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
	"strings"
)

const PTN_921130 = `(?:\bhttp/\d|<(?:html|meta)\b)`

var re921130 *re2.Regexp

type Rule921130 struct {
}

func (r *Rule921130) Id() string {
	return "921130"
}

func (r *Rule921130) Phase() int {
	return 2
}

func (r *Rule921130) Evaluate(tx *core.Transaction) bool {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, k) {
				return true
			}

			if r.doEvaluate(tx, v) {
				return true
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return true
		}
	}

	return true
}

func (r *Rule921130) doEvaluate(tx *core.Transaction, value string) bool {
tv, _, _ := core.HtmlEntityDecode(value)
m := re921130.MatchString(strings.ToLower(tv))

if m {
tx.Variables.HttpViolationScore += CRITICAL_ANOMALY_SCORE
tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
}

return m
}

func init() {
	re921130, _ = re2.Compile(PTN_921130)
}
