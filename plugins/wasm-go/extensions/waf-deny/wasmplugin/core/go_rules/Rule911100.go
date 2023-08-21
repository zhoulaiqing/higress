package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"sort"
)

type Rule911100 struct {
}

func (r *Rule911100) Id() string {
	return "911100"
}

func (r *Rule911100) Evaluate(tx core.Transaction) bool {
	found := sort.SearchStrings(core.ALLOWED_METHODS, tx.Variables.RequestMethod)
	if found < 0 {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
