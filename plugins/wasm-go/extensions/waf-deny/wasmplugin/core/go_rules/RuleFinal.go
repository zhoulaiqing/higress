package core

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

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
func (r *RuleFinal) Evaluate(tx core.Transaction) bool {
	if tx.Variables.InboundAnomalyScorePl1 >= core.INBOUND_ANOMALY_SCORE_THRESHOLD {
		return false
	}

	return true
}
