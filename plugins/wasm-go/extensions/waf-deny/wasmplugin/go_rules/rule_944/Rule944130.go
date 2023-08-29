package rule_944

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule944130 struct {
}

func (r *Rule944130) Id() string {
	return "944130"
}

func (r *Rule944130) Phase() int {
	return 2
}

func (r *Rule944130) Evaluate(tx *core.Transaction) int {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, &k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, &k) {
				return rule_tasks.BLOCK
			}

			if r.doEvaluate(tx, &v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for k, v := range tx.Variables.RequestHeaders {
		if r.doEvaluate(tx, &k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	if r.doEvaluate(tx, &tx.Variables.RequestBody) {
		return rule_tasks.BLOCK
	}

	if r.doEvaluate(tx, &tx.Variables.RequestFileName) {
		return rule_tasks.BLOCK
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	for _, v := range tx.Variables.XML["//@*"] {
		if r.doEvaluate(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule944130) doEvaluate(tx *core.Transaction, value *string) bool {
	m, _ := core.PmEvaluate(rule_tasks.Rule944130Matcher, *value, false)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
