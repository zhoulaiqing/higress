package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_921190 = `[\n\r]`

type Rule921190 struct {
}

func (r *Rule921190) Id() string {
	return "921190"
}

func (r *Rule921190) Evaluate(tx core.Transaction) bool {
	requestFileName := tx.Variables.RequestFileName
	requestFileName, _, _ = core.UrlDecodeUni(requestFileName)

	if m, _ := hyperscan.MatchString(PTN_921190, requestFileName); m {
		tx.Variables.HttpViolationScore += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
