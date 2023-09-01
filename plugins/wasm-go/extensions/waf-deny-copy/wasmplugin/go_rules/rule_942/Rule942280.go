package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942280 struct {
	*Rule942
}

func (r *Rule942280) Id() string {
	return "942280"
}

func (r *Rule942280) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx, r.doEvaluate, noAddition942)
}

func (r *Rule942280) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942280.MatchString(*value)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
