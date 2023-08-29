package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule941360 struct {
	*Rule941
}

func (r *Rule941360) Id() string {
	return "941360"
}

func (r *Rule941360) Evaluate(tx *core.Transaction) int {
	return r.evaluateRawValue(tx)
}

func (r *Rule941360) GetAddition() *Rule941Addition {
	return fileNameAddition
}

func (r *Rule941360) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re941360.MatchString(*value)

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
