package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_921140 = `[\n\r]`

var re921140 *re2.Regexp

type Rule921140 struct {
}

func (r *Rule921140) Id() string {
	return "921140"
}

func (r *Rule921140) Phase() int {
	return 1
}

func (r *Rule921140) Evaluate(tx *core.Transaction) int {
	for k, v := range tx.Variables.RequestHeaders {
		if r.doEvaluate(tx, k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule921140) doEvaluate(tx *core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)

	m := re921140.MatchString(tv)

	if m {
		tx.Variables.HttpViolationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}

func init() {
	re921140, _ = re2.Compile(PTN_921140)
}
