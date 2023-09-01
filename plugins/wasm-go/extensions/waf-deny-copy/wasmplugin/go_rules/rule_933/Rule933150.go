package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule933150 struct {
	*Rule933
}

func (r *Rule933150) Id() string {
	return "933150"
}

func (r *Rule933150) Evaluate(tx *core.Transaction) int {
	r933 := r.Rule933.Evaluate(tx)
	if r933 == rule_tasks.BLOCK {
		return r933
	}

	if r.doEvaluate(tx, tx.Variables.RequestFileName) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule933150) doEvaluate(tx *core.Transaction, value string) bool {
	m, _ := core.PmEvaluate(rule_tasks.Rule933150Matcher, value, false)
	if !m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
