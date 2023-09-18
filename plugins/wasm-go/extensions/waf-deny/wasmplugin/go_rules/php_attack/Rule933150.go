package php_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

// 933120 933130 933150

type Rule933150 struct {
}

func (r *Rule933150) Id() string {
	return "933150"
}

func (r *Rule933150) Phase() int {
	return 2
}

func (r *Rule933150) Evaluate(tx *core.Transaction) int {

	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluateDefault(tx, k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluateDefault(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluateDefault(tx, k) {
				return rule_tasks.BLOCK
			}

			if r.doEvaluateDefault(tx, v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluateDefault(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	if r.doEvaluateFileName(tx, tx.Variables.RequestFileName) {
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func (r *Rule933150) doEvaluateDefault(tx *core.Transaction, value string) bool {
	v, _, _ := core.NormalisePath(value)
	v, _, _ = core.UrlDecodeUni(v)
	v, _, _ = core.RemoveWhitespace(v)

	m, _ := core.PmEvaluate(rule_tasks.Rule933Matcher, v, false)
	if !m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}

func (r *Rule933150) doEvaluateFileName(tx *core.Transaction, value string) bool {
	m, _ := core.PmEvaluate(rule_tasks.Rule933150Matcher, value, false)
	if !m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
