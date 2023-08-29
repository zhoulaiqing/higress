package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942540 struct {
	*Rule942
}

func (r *Rule942540) Id() string {
	return "942540"
}

func (r *Rule942540) doEvaluate(tx *core.Transaction, value *string) bool {
	v, _, _ := core.ReplaceComments(*value)
	m := rule_tasks.Re942540.MatchString(v)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
