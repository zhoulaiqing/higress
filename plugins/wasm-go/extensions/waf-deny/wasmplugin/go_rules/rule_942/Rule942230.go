package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942230 struct {
	*Rule942
}

func (r *Rule942230) Id() string {
	return "942230"
}

func (r *Rule942230) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942230.MatchString(*value)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
