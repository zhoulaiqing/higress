package generic_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule934110 struct {
	*Rule934
}

func (r *Rule934110) Id() string {
	return "934110"
}

func (r *Rule934110) validateFileName() bool {
	return true
}

func (r *Rule934110) doEvaluate(tx *core.Transaction, value string) bool {

	m, _ := core.PmEvaluate(rule_tasks.Rule934110Matcher, value, false)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
