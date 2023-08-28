package rule_920

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule920181 struct {
}

func (r *Rule920181) Id() string {
	return "920181"
}

func (r *Rule920181) Phase() int {
	return 1
}

func (r *Rule920181) Evaluate(tx *core.Transaction) bool {
	if len(tx.Variables.RequestHeaders["transfer-encoding"]) > 0 && len(tx.Variables.RequestHeaders["content-length"]) > 0 {
		tx.Variables.InboundAnomalyScorePl1 += go_rules.WARNING_ANOMALY_SCORE
	}

	return true
}
