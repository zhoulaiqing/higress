package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941370_400 struct {
	*Rule941
}

func (r *Rule941370_400) Id() string {
	return "941370_400"
}

func (r *Rule941370_400) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx, r.doEvaluate, fileNameAddition)
}

func (r *Rule941370_400) doEvaluate(tx *core.Transaction, value *string) bool {
	v, _, _ := core.UrlDecodeUni(*value)
	v, _, _ = core.CompressWhitespace(v)

	m := rule_tasks.Re941370.MatchString(v)
	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	m = rule_tasks.Re941400.MatchString(v)
	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
