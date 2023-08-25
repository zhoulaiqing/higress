package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_930110 = `(?:(?:^|[\x5c/;])\.{2,3}[\x5c/;]|[\x5c/;]\.{2,3}(?:[\x5c/;]|$))`

type Rule930110 struct {
}

func (r *Rule930110) Id() string {
	return "930110"
}

func (r *Rule930110) Phase() int {
	return 2
}

func (r *Rule930110) Evaluate(tx core.Transaction) bool {
	if r.doEvaluate(tx, tx.Variables.RequestUri) {
		return true
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

	for k, v := range tx.Variables.RequestHeaders {
		if k == "Refer" {
			continue
		}

		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	for k, values := range tx.Variables.Files {
		if r.doEvaluate(tx, k) {
			return true
		}
		for _, v := range values {
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

// 注意是 multiMatch
func (r *Rule930110) doEvaluate(tx core.Transaction, value string) bool {
	m, _ := hyperscan.MatchString(PTN_930110, value)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.LfiScore += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ := core.Utf8ToUnicode(value)
	m, _ = hyperscan.MatchString(PTN_930110, v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.LfiScore += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.UrlDecodeUni(v)
	m, _ = hyperscan.MatchString(PTN_930110, v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.LfiScore += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.RemoveNulls(v)
	m, _ = hyperscan.MatchString(PTN_930110, v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.LfiScore += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.CmdLine(v)
	m, _ = hyperscan.MatchString(PTN_930110, v)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		tx.Variables.LfiScore += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	return m
}
