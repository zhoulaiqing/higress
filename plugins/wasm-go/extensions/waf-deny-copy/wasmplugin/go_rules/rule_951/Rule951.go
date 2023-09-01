package rule_951

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule951 struct {
}

func (r *Rule951) Id() string {
	return "951"
}

func (r *Rule951) Phase() int {
	return 4
}

func (r *Rule951) Evaluate(tx *core.Transaction) int {
	return rule_tasks.PASS
}

type doEvaluateFunc func(*core.Transaction) bool

func (r *Rule951) evaluate(tx *core.Transaction, evaluateFunc doEvaluateFunc) bool {
	if tx.Variables.Skip951 {
		return false
	}

	m := evaluateFunc(tx)
	if m {
		tx.Variables.SqlInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
