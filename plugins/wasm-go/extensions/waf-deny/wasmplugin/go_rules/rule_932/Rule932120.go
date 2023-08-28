package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932120 struct {
	*Rule932
}

func (r *Rule932120) Id() string {
	return "932120"
}

func (r *Rule932120) doEvaluate(tx *core.Transaction, value string) bool {
	m, _ := core.PmEvaluate(rule_tasks.Rule932160Matcher, value, false)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
