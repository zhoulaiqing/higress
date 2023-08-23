package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

type Rule920210 struct {
}

func (r *Rule920210) Id() string {
	return "920210"
}

func (r *Rule920210) Evaluate(tx core.Transaction) bool {

	matched, _ := hyperscan.MatchString(`\b(?:keep-alive|close),\s?(?:keep-alive|close)\b`, tx.Variables.RequestHeaders["Connection"])
	if matched {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}
	return true
}
