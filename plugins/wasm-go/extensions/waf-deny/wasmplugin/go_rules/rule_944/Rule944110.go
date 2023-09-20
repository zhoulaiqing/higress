package rule_944

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"strings"
)

type Rule944110 struct {
}

func (r *Rule944110) Id() string {
	return "944110"
}

func (r *Rule944110) Phase() int {
	return 2
}

func (r *Rule944110) Evaluate(tx *core.Transaction) int {
	if r.evaluateStep(tx, r.doEvaluateStep1) == rule_tasks.PASS {
		return rule_tasks.PASS
	}

	return r.evaluateStep(tx, r.doEvaluateStep2)
}

type evaluateFn func(*core.Transaction, *string) bool

func (r *Rule944110) evaluateStep(tx *core.Transaction, fn evaluateFn) int {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if fn(tx, &k) {
			return rule_tasks.BLOCK
		}

		if fn(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if fn(tx, &k) {
				return rule_tasks.BLOCK
			}

			if fn(tx, &v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for k, v := range tx.Variables.RequestHeaders {
		if fn(tx, &k) {
			return rule_tasks.BLOCK
		}

		if fn(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	if fn(tx, &tx.Variables.RequestBody) {
		return rule_tasks.BLOCK
	}

	for _, v := range tx.Variables.XML["/*"] {
		if fn(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	for _, v := range tx.Variables.XML["//@*"] {
		if fn(tx, &v) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule944110) doEvaluateStep1(tx *core.Transaction, value *string) bool {
	return rule_tasks.Re9441101.MatchString(strings.ToLower(*value))
}

func (r *Rule944110) doEvaluateStep2(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re9441102.MatchString(strings.ToLower(*value))
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
