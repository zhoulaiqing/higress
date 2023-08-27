package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

// Include 920290

type Rule920280 struct {
}

func (r *Rule920280) Id() string {
	return "920280"
}

func (r *Rule920280) Phase() int {
	return 1
}

func (r *Rule920280) Evaluate(tx *core.Transaction) bool {
	if len(tx.Variables.RequestHeaders["Host"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += WARNING_ANOMALY_SCORE
	}

	return true
}
