package rce_attack

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
)

type Rule932125 struct {
	*Rule932
}

func (r *Rule932125) Id() string {
	return "932125"
}

func (r *Rule932125) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re932125.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
