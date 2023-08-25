package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	"strings"
)

const PTN_921160 = `[\n\r]+(?:\s|location|refresh|(?:set-)?cookie|(?:x-)?(?:forwarded-(?:for|host|server)|host|via|remote-ip|remote-addr|originating-IP))\s*:`

type Rule921160 struct {
}

func (r *Rule921160) Id() string {
	return "921160"
}

func (r *Rule921160) Phase() int {
	return 1
}

func (r *Rule921160) Evaluate(tx core.Transaction) bool {
	for k, v := range tx.Variables.ArgsGet {
		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	return true
}

func (r *Rule921160) doEvaluate(tx core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)
	m, _ := hyperscan.MatchString(PTN_921160, strings.ToLower(tv))

	if m {
		tx.Variables.HttpViolationScore += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return m
}
