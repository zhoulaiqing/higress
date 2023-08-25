package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_921421 = `^[^\s\v,;]+[\s\v,;].*?(?:application/(?:.+\+)?json|(?:application/(?:soap\+)?|text/)xml)`

type Rule921421 struct {
}

func (r *Rule921421) Id() string {
	return "921421"
}

func (r *Rule921421) Phase() int {
	return 1
}

func (r *Rule921421) Evaluate(tx core.Transaction) bool {
	if m, _ := hyperscan.MatchString(PTN_921421, tx.Variables.RequestHeaders["Content-Type"]); m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
