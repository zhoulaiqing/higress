package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941390 struct {
	*Rule941
}

func (r *Rule941390) Id() string {
	return "941390"
}

func (r *Rule941390) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx, r.doEvaluate, fileNameAddition)
}

func (r *Rule941390) doEvaluate(tx *core.Transaction, value *string) bool {
	v, _, _ := core.HtmlEntityDecode(*value)
	v, _, _ = core.JsDecode(v)

	m := rule_tasks.Re941390.MatchString(v)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
