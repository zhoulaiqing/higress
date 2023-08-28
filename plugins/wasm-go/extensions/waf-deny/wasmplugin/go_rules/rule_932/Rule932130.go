package rule_932

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule932130 struct {
	*Rule932
}

func (r *Rule932130) Id() string {
	return "932130"
}

func (r *Rule932130) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.CmdLine(value)
	m := rule_tasks.Re932130.MatchString(v)
	if m {
		tx.Variables.RceScore += go_rules.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return m
}
