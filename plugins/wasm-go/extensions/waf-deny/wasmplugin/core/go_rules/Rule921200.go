package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	"strings"
)

const PTN_921200 = `^[^:\(\)\&\|\!\<\>\~]*\)\s*(?:\((?:[^,\(\)\=\&\|\!\<\>\~]+[><~]?=|\s*[&!|]\s*(?:\)|\()?\s*)|\)\s*\(\s*[\&\|\!]\s*|[&!|]\s*\([^\(\)\=\&\|\!\<\>\~]+[><~]?=[^:\(\)\&\|\!\<\>\~]*)`

type Rule921200 struct {
}

func (r *Rule921200) Id() string {
	return "921200"
}

func (r *Rule921200) Phase() int {
	return 2
}

func (r *Rule921200) Evaluate(tx core.Transaction) bool {
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

func (r *Rule921200) doEvaluate(tx core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)
	m, _ := hyperscan.MatchString(PTN_921200, tv)

	if m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return m
}
