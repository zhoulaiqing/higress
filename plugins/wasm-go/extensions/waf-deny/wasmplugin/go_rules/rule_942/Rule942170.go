package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942170 struct {
	*Rule942
}

func (r *Rule942170) Id() string {
	return "942170"
}

func (r *Rule942170) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942170.MatchString(*value)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
