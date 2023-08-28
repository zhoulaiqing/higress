package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932160 struct {
	*Rule932
}

func (r *Rule932160) Id() string {
	return "932160"
}

func (r *Rule932160) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.CmdLine(value)
	v, _, _ = core.NormalisePath(v)
	m, _ := core.PmEvaluate(rule_tasks.Rule932160Matcher, v, false)
	if m {
		tx.Variables.RceScore += go_rules.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return m
}
