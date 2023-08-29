package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule941240 struct {
	*Rule941
}

func (r *Rule941240) Id() string {
	return "941240"
}

func (r *Rule941240) GetAddition() *Rule941Addition {
	return fileNameAddition
}

func (r *Rule941240) doEvaluate(tx *core.Transaction, value *string) bool {

	m := rule_tasks.Re941240.MatchString(strings.ToLower(*value))

	if m {
		tx.Variables.XssScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
