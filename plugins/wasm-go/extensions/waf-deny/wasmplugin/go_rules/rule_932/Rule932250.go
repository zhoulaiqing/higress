package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932250 struct {
	*Rule932
}

func (r *Rule932250) Id() string {
	return "932250"
}

func (r *Rule932250) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re932250.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
