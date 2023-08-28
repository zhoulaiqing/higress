package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule933120 struct {
	*Rule933
}

func (r *Rule933120) Id() string {
	return "933120"
}

func (r *Rule933120) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.NormalisePath(value)
	m, values := core.PmEvaluate(rule_tasks.Rule933120Matcher, v, true)
	if !m {
		return false
	}

	for _, v := range values {
		if strings.Contains(v, "=") {
			tx.Variables.PhpInjectionScore += go_rules.CRITICAL_ANOMALY_SCORE
			tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
			return true
		}
	}

	return false
}
