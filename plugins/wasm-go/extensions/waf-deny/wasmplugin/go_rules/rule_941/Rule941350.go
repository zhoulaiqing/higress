package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941350 struct {
	*Rule941
}

func (r *Rule941350) Id() string {
	return "941350"
}

func (r *Rule941350) Evaluate(tx *core.Transaction) int {
	return r.Rule941.evaluateRawValue(tx, r.doEvaluate, fileNameAddition)
}

func (r *Rule941350) doEvaluate(tx *core.Transaction, value *string) bool {

	v, _, _ := core.UrlDecode(*value)
	v, _, _ = core.HtmlEntityDecode(v)
	v, _, _ = core.JsDecode(v)

	m := rule_tasks.Re941350.MatchString(v)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}