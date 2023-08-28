package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932180 struct {
	*Rule932
}

func (r *Rule932180) Id() string {
	return "932180"
}

func (r *Rule932180) doEvaluate(tx *core.Transaction, value string) bool {
	m, _ := core.PmEvaluate(rule_tasks.Rule932180Matcher, value, false)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
