package rule_933

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule933130 struct {
	*Rule933
}

func (r *Rule933130) Id() string {
	return "933130"
}

func (r *Rule933130) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.NormalisePath(value)
	v, _, _ = core.UrlDecodeUni(v)
	m, _ := core.PmEvaluate(rule_tasks.Rule933130Matcher, v, false)
	if !m {
		tx.Variables.PhpInjectionScore += go_rules.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return m
}
