package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_920260 = `\%u[fF]{2}[0-9a-fA-F]{2}`

type Rule920260 struct {
}

func (r *Rule920260) Id() string {
	return "920260"
}

func (r *Rule920260) Evaluate(tx core.Transaction) bool {
	m1, _ := hyperscan.MatchString(PTN_920260, tx.Variables.RequestUri)
	if m1 {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	m2, _ := hyperscan.MatchString(PTN_920260, tx.Variables.RequestBody)
	if m2 {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
