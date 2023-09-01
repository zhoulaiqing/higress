package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule933180 struct {
	*Rule933
}

func (r *Rule933180) Id() string {
	return "933180"
}

func (r *Rule933180) Evaluate(tx *core.Transaction) int {
	r933 := r.Rule933.Evaluate(tx)
	if r933 == rule_tasks.BLOCK {
		return r933
	}

	if r.doEvaluate(tx, tx.Variables.RequestFileName) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule933180) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re933180.MatchString(value)
	if !m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
