package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941130 struct {
	*Rule941
}

func (r *Rule941130) Id() string {
	return "941130"
}

func (r *Rule941130) GetAddition() *Rule941Addition {
	return additionWithoutReferer
}

func (r *Rule941130) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re941130.MatchString(*value)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
