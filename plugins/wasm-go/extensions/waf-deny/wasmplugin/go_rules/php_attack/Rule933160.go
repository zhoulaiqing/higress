package php_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule933160 struct {
	*Rule933
}

func (r *Rule933160) Id() string {
	return "933160"
}

func (r *Rule933160) Evaluate(tx *core.Transaction) int {
	r933 := r.Rule933.Evaluate(tx)
	if r933 == rule_tasks.BLOCK {
		return r933
	}

	if r.doEvaluate(tx, tx.Variables.RequestFileName) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule933160) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re933160.MatchString(value)
	if !m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
