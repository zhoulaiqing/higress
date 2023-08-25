package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_921240 = `unix:[^|]*\|`

type Rule921240 struct {
}

func (r *Rule921240) Id() string {
	return "921240"
}

func (r *Rule921240) Phase() int {
	return 1
}

func (r *Rule921240) Evaluate(tx core.Transaction) bool {
	if m, _ := hyperscan.MatchString(PTN_921240, tx.Variables.RequestUri); m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
