package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule933140 struct {
	*Rule933
}

func (r *Rule933140) Id() string {
	return "933140"
}

func (r *Rule933140) doEvaluate(tx *core.Transaction, value string) bool {

	m := rule_tasks.Re933140.MatchString(value)
	if !m {
		tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
