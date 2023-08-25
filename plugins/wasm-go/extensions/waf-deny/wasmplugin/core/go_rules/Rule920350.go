package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_920350 = "(?:^([\\d.]+|\\[[\\da-f:]+\\]|[\\da-f:]+)(:[\\d]+)?$)"

type Rule920350 struct {
}

func (r *Rule920350) Id() string {
	return "920350"
}

func (r *Rule920350) Phase() int {
	return 1
}

func (r *Rule920350) Evaluate(tx core.Transaction) bool {
	matched, _ := hyperscan.MatchString(PTN_920350, tx.Variables.RequestHeaders["Host"])

	if matched {
		tx.Variables.InboundAnomalyScorePl1 += core.WARNING_ANOMALY_SCORE
	}

	return true
}
