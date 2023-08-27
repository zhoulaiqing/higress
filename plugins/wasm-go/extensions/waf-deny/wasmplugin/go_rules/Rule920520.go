package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule920520 struct {
}

func (r *Rule920520) Id() string {
	return "920520"
}

func (r *Rule920520) Phase() int {
	return 1
}

func (r *Rule920520) Evaluate(tx *core.Transaction) bool {
	if len(tx.Variables.RequestHeaders["Accept-Encoding"]) > 50 {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return true
}
