package rule_934

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule934170 struct {
	*Rule934
}

func (r *Rule934170) Id() string {
	return "934170"
}

func (r *Rule934170) validateFileName() bool {
	return true
}

func (r *Rule934170) doEvaluate(tx *core.Transaction, value string) bool {

	v, _, _ := core.UrlDecodeUni(value)
	m := rule_tasks.Re934170.MatchString(v)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
