package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942500 struct {
	*Rule942
}

func (r *Rule942500) Id() string {
	return "942500"
}

func (r *Rule942500) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx, r.doEvaluate, noAddition942)
}

func (r *Rule942500) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942500.MatchString(*value)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
