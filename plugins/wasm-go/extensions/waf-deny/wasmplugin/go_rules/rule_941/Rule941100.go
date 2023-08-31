package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/corazawaf/libinjection-go"
)

type Rule941100 struct {
	*Rule941
}

func (r *Rule941100) Id() string {
	return "941100"
}

func (r *Rule941100) Evaluate(tx *core.Transaction) int {
	return r.evaluateByCache(tx, r.doEvaluate, noAddition941)
}

func (r *Rule941100) doEvaluate(tx *core.Transaction, value *string) bool {
	m := libinjection.IsXSS(*value)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
