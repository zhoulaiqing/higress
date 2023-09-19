package rce_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932370 struct {
	*Rule932
}

func (r *Rule932370) Id() string {
	return "932370"
}

func (r *Rule932370) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re932370.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
