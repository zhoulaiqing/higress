package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_920500 = `\.[^.~]+~(?:/.*|)$`

type Rule920500 struct {
}

func (r *Rule920500) Id() string {
	return "920500"
}

func (r *Rule920500) Phase() int {
	return 1
}

func (r *Rule920500) Evaluate(tx core.Transaction) bool {

	fileName, _, _ := core.UrlDecodeUni(tx.Variables.RequestFileName)

	if m, _ := hyperscan.MatchString(PTN_920500, fileName); m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
