package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"golang.org/x/exp/slices"
)

type Rule920450 struct {
}

func (r *Rule920450) Id() string {
	return "920450"
}

func (r *Rule920450) Phase() int {
	return 1
}

func (r *Rule920450) Evaluate(tx core.Transaction) bool {
	for hk, _ := range tx.Variables.RequestHeaders {
		if slices.Contains(core.RESTRICTED_HEADERS, hk) {
			tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
			return true
		}
	}

	return true
}
