package rule_944

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"strings"
)

type Rule944150 struct {
}

func (r *Rule944150) Id() string {
	return "944150"
}

func (r *Rule944150) Phase() int {
	return 2
}

func (r *Rule944150) Evaluate(tx *core.Transaction) int {
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

	if r.doEvaluate(tx, &tx.Variables.RequestLine) {
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

func (r *Rule944150) doEvaluate(tx *core.Transaction, value *string) bool {
	v, _, _ := core.JsDecode(*value)
	v, _, _ = core.HtmlEntityDecode(v)

	m := rule_tasks.Re944150.MatchString(v)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
