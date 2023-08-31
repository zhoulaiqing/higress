package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942290 struct {
	*Rule942
}

func (r *Rule942290) Id() string {
	return "942290"
}

func (r *Rule942290) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx, r.doEvaluate, noAddition942)
}

func (r *Rule942290) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942290.MatchString(*value)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
