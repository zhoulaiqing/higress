package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	"strings"
)

const PTN_921120 = `[\r\n]\W*?(?:content-(?:type|length)|set-cookie|location):\s*\w`

type Rule921120 struct {
}

func (r *Rule921120) Id() string {
	return "921120"
}

func (r *Rule921120) Phase() int {
	return 2
}

func (r *Rule921120) Evaluate(tx core.Transaction) bool {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, k) {
				return true
			}

			if r.doEvaluate(tx, v) {
				return true
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return true
		}
	}

	return true
}

func (r *Rule921120) doEvaluate(tx core.Transaction, value string) bool {
	m, _ := hyperscan.MatchString(PTN_921120, strings.ToLower(value))

	if m {
		tx.Variables.HttpViolationScore += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return m
}
