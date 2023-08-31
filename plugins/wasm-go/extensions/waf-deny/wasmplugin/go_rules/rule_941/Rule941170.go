package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941170 struct {
	*Rule941
}

func (r *Rule941170) Id() string {
	return "941170"
}

func (r *Rule941170) Evaluate(tx *core.Transaction) int {
	return r.evaluateByCache(tx, r.doEvaluate, fullAddition)
}

func (r *Rule941170) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re941170.MatchString(*value)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
