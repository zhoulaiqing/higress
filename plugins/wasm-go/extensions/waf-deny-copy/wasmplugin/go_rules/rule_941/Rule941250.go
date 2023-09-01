package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941250 struct {
	*Rule941
}

func (r *Rule941250) Id() string {
	return "941250"
}

func (r *Rule941250) Evaluate(tx *core.Transaction) int {
	return r.evaluateByCache(tx, r.doEvaluate, fileNameAddition)
}

func (r *Rule941250) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re941250.MatchString(*value)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
