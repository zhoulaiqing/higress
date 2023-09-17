package generic_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule934101 struct {
	*Rule934
}

func (r *Rule934101) Id() string {
	return "934101"
}

func (r *Rule934101) validateFileName() bool {
	return true
}

func (r *Rule934101) doEvaluate(tx *core.Transaction, value string) bool {

	v, _, _ := core.UrlDecodeUni(value)
	v, _, _ = core.Base64decode(v)

	m := rule_tasks.Re934101.MatchString(v)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
