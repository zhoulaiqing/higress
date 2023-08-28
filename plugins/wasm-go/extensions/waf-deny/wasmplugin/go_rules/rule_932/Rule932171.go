package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932171 struct {
}

func (r *Rule932171) Id() string {
	return "932171"
}

func (r *Rule932171) Phase() int {
	return 2
}

func (r *Rule932171) Evaluate(tx *core.Transaction) int {
	for k, v := range tx.Variables.RequestHeaders {
		if r.doEvaluate(tx, k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	if r.doEvaluate(tx, tx.Variables.RequestLine) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule932171) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.UrlDecode(value)
	v, _, _ = core.UrlDecodeUni(v)

	// 复用 Re932170
	m := rule_tasks.Re932170.MatchString(v)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
