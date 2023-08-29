package rule_942

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule942151 struct {
	*Rule942
}

func (r *Rule942151) Id() string {
	return "942151"
}

func (r *Rule942151) GetAddition() *Rule942Addition {
	return noAddition942
}

func (r *Rule942151) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re942151.MatchString(strings.ToLower(*value))

	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
