package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule942220 struct {
	*Rule942
}

func (r *Rule942220) Id() string {
	return "942220"
}

func (r *Rule942220) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942220.MatchString(*value)

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
