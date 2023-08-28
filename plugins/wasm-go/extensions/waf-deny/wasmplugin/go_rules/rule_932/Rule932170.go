package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932170 struct {
}

func (r *Rule932170) Id() string {
	return "932170"
}

func (r *Rule932170) Phase() int {
	return 1
}

func (r *Rule932170) Evaluate(tx *core.Transaction) bool {
	for k, v := range tx.Variables.RequestHeaders {
		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	if r.doEvaluate(tx, tx.Variables.RequestLine) {
		return true
	}

	return true
}

func (r *Rule932170) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.UrlDecode(value)

	m := rule_tasks.Re932170.MatchString(v)
	if m {
		tx.Variables.RceScore += go_rules.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return m
}
