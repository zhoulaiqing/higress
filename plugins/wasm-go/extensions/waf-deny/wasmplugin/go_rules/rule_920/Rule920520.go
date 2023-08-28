package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule920520 struct {
}

func (r *Rule920520) Id() string {
	return "920520"
}

func (r *Rule920520) Phase() int {
	return 1
}

func (r *Rule920520) Evaluate(tx *core.Transaction) int {
	if len(tx.Variables.RequestHeaders["accept-encoding"]) > 50 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
