package rule_934

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule934150 struct {
	*Rule934
}

func (r *Rule934150) Id() string {
	return "934150"
}

func (r *Rule934150) doEvaluate(tx *core.Transaction, value string) bool {

	m := rule_tasks.Re934150.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
