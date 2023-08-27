package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

type Rule920340 struct {
}

func (r *Rule920340) Id() string {
	return "932340"
}

func (r *Rule920340) Phase() int {
	return 1
}

func (r *Rule920340) Evaluate(tx *core.Transaction) bool {
	if tx.Variables.RequestHeaders["content-length"] == "0" {
		return true
	}

	if len(tx.Variables.RequestHeaders["content-type"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += NOTICE_ANOMALY_SCORE
	}

	return true
}
