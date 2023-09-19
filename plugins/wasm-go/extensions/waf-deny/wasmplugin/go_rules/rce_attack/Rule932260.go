package rce_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932260 struct {
	*Rule932
}

func (r *Rule932260) Id() string {
	return "932260"
}

func (r *Rule932260) doEvaluate(tx *core.Transaction, value string) bool {
	matches := rule_tasks.Re9322601.FindStringSubmatch(value)
	if len(matches) == 0 {
		return false
	}

	m := rule_tasks.Re9322602.MatchString(matches[0])
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
