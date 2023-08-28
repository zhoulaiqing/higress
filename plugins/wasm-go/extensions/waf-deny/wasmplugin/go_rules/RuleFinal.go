package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
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

	return rule_tasks.PASS
}
