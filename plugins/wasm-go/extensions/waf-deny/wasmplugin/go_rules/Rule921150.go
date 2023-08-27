package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
)

const PTN_921150 = `[\n\r]`

var re921150 *re2.Regexp

type Rule921150 struct {
}

func (r *Rule921150) Id() string {
	return "921140"
}

func (r *Rule921150) Phase() int {
	return 2
}

func (r *Rule921150) Evaluate(tx *core.Transaction) bool {

	for _, argMap := range tx.Variables.Args {
		for k, _ := range *argMap {
			if r.doEvaluate(tx, k) {
				return true
			}
		}
	}

	return true
}

func (r *Rule921150) doEvaluate(tx *core.Transaction, value string) bool {
tv, _, _ := core.HtmlEntityDecode(value)

m := re921150.MatchString(tv)

if m {
tx.Variables.HttpViolationScore += CRITICAL_ANOMALY_SCORE
tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
}

return m
}

func init() {
	re921150, _ = re2.Compile(PTN_921150)
}
