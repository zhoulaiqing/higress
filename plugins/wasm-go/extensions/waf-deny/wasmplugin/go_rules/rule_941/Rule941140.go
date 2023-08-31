package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941140 struct {
	*Rule941
}

func (r *Rule941140) Id() string {
	return "941140"
}

func (r *Rule941140) Evaluate(tx *core.Transaction) int {
	return r.evaluateByCache(tx, r.doEvaluate, fullAddition)
}

func (r *Rule941140) doEvaluate(tx *core.Transaction, value *string) bool {
	v, _, _ := core.RemoveWhitespace(*value)
	m := rule_tasks.Re941140.MatchString(v)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
