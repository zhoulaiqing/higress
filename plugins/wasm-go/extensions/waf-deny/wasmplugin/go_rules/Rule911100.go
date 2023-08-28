package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"sort"
)

type Rule911100 struct {
}

func (r *Rule911100) Id() string {
	return "911100"
}

func (r *Rule911100) Phase() int {
	return 1
}

func (r *Rule911100) Evaluate(tx *core.Transaction) int {
	found := sort.SearchStrings(rule_tasks.ALLOWED_METHODS, tx.Variables.RequestMethod)
	if found < 0 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
