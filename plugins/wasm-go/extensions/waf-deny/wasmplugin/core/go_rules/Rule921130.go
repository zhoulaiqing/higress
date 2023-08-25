package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	"strings"
)

const PTN_921130 = `(?:\bhttp/\d|<(?:html|meta)\b)`

type Rule921130 struct {
}

func (r *Rule921130) Id() string {
	return "921130"
}

func (r *Rule921130) Phase() int {
	return 2
}

func (r *Rule921130) Evaluate(tx core.Transaction) bool {
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

func (r *Rule921130) doEvaluate(tx core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)
	m, _ := hyperscan.MatchString(PTN_921130, strings.ToLower(tv))

	if m {
		tx.Variables.HttpViolationScore += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return m
}
