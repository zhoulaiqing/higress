package go_rules

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
)

/**
临时使用
*/

type RuleFinal struct {
}

func (r *RuleFinal) Id() string {
	return "Final"
}

func (r *RuleFinal) Phase() int {
	return 100
}

// Evaluate 一个极其简单的评估逻辑
func (r *RuleFinal) Evaluate(tx *core.Transaction) int {

	if tx.Variables.InboundAnomalyScorePl1 >= rule_tasks.INBOUND_ANOMALY_SCORE_THRESHOLD {
		return rule_tasks.DENY
	}
	if tx.Variables.OutboundAnomalyScorePl1 >= rule_tasks.OUTBOUND_ANOMALY_SCORE_THRESHOLD {
		return rule_tasks.DENY
	}

	return rule_tasks.PASS
}
