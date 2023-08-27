package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule920330 struct {
}

func (r *Rule920330) Id() string {
	return "920330"
}

func (r *Rule920330) Phase() int {
	return 1
}

func (r *Rule920330) Evaluate(tx *core.Transaction) bool {
	if len(tx.Variables.RequestHeaders["user-agent"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += NOTICE_ANOMALY_SCORE
	}

	return true
}
