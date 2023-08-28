package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule933170 struct {
	*Rule933
}

func (r *Rule933170) Id() string {
	return "933170"
}

func (r *Rule933170) Evaluate(tx *core.Transaction) int {
	r933 := r.Rule933.Evaluate(tx)
	if r933 == rule_tasks.BLOCK {
		return r933
	}

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

func (r *Rule933170) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re933170.MatchString(value)
	if !m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
