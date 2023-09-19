package rce_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932140 struct {
	*Rule932
}

func (r *Rule932140) Id() string {
	return "932140"
}

func (r *Rule932140) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.CmdLine(value)
	m := rule_tasks.Re932140.MatchString(v)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
