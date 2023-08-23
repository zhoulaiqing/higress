package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_920530 = `charset.*?charset`

type Rule920530 struct {
}

func (r *Rule920530) Id() string {
	return "920530"
}

func (r *Rule920530) Evaluate(tx core.Transaction) bool {
	if m, _ := hyperscan.MatchString(PTN_920530, tx.Variables.RequestHeaders[""]); m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
