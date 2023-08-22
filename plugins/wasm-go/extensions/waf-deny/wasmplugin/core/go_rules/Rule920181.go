package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule920181 struct {
}

func (r *Rule920181) Id() string {
	return "920181"
}

func (r *Rule920181) Evaluate(tx core.Transaction) bool {
	if len(tx.Variables.RequestHeaders["Transfer-Encoding"]) > 0 && len(tx.Variables.RequestHeaders["Content-Length"]) > 0 {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
