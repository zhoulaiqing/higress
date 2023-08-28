package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

// Include 920290

type Rule920280 struct {
}

func (r *Rule920280) Id() string {
	return "920280"
}

func (r *Rule920280) Phase() int {
	return 1
}

func (r *Rule920280) Evaluate(tx *core.Transaction) int {
	if len(tx.Variables.RequestHeaders["host"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
