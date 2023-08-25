package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_921140 = `[\n\r]`

type Rule921140 struct {
}

func (r *Rule921140) Id() string {
	return "921140"
}

func (r *Rule921140) Phase() int {
	return 1
}

func (r *Rule921140) Evaluate(tx core.Transaction) bool {
	for k, v := range tx.Variables.RequestHeaders {
		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	return true
}

func (r *Rule921140) doEvaluate(tx core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)

	m, _ := hyperscan.MatchString(PTN_921140, tv)

	if m {
		tx.Variables.HttpViolationScore += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return m
}
