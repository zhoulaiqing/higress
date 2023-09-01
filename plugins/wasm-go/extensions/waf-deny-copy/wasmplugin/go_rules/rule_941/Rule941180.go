package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941180 struct {
	*Rule941
}

func (r *Rule941180) Id() string {
	return "941180"
}

func (r *Rule941180) Evaluate(tx *core.Transaction) int {
	return r.evaluateByCache(tx, r.doEvaluate, fileNameAddition)
}

func (r *Rule941180) doEvaluate(tx *core.Transaction, value *string) bool {

	m, _ := core.PmEvaluate(rule_tasks.Rule941180Matcher, *value, false)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
