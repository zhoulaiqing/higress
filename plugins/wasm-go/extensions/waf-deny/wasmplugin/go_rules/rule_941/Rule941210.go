package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941210 struct {
	*Rule941
}

func (r *Rule941210) Id() string {
	return "941210"
}

func (r *Rule941210) GetAddition() *Rule941Addition {
	return fileNameAddition
}

func (r *Rule941210) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re941210.MatchString(*value)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
