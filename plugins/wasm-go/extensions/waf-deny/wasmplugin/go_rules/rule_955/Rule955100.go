package rule_955

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule955100 struct {
}

func (r *Rule955100) Id() string {
	return "955100"
}

func (r *Rule955100) Phase() int {
	return 4
}

func (r *Rule955100) Evaluate(tx *core.Transaction) int {
	if m, _ := core.PmEvaluate(rule_tasks.Rule955100Matcher, tx.Variables.ResponseBody, false); m {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
