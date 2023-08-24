package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_921150 = `[\n\r]`

type Rule921150 struct {
}

func (r *Rule921150) Id() string {
	return "921140"
}

func (r *Rule921150) Evaluate(tx core.Transaction) bool {

	for _, argMap := range tx.Variables.Args {
		for k, _ := range *argMap {
			if r.doEvaluate(tx, k) {
				return true
			}
		}
	}

	return true
}

func (r *Rule921150) doEvaluate(tx core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)

	m, _ := hyperscan.MatchString(PTN_921150, tv)

	if m {
		tx.Variables.HttpViolationScore += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return m
}
