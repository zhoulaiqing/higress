package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

// 932130 932160 932180

type Rule932180 struct {
	*Rule932
}

func (r *Rule932180) Id() string {
	return "932180"
}

func (r *Rule932180) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.CmdLine(value)
	v, _, _ = core.NormalisePath(v)
	m, _ := core.PmEvaluate(rule_tasks.Rule932Matcher, value, false)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
