package rce_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932175 struct {
	*Rule932
}

func (r *Rule932175) Id() string {
	return "932175"
}

func (r *Rule932175) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re932175.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
