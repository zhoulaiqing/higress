package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942160 struct {
	*Rule942
}

func (r *Rule942160) Id() string {
	return "942160"
}

func (r *Rule942160) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx, r.doEvaluate, baseNameAddition)
}
func (r *Rule942160) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942160.MatchString(*value)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
