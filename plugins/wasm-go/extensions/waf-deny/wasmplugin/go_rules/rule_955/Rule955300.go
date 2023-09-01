package rule_955

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule955300 struct {
}

func (r *Rule955300) Id() string {
	return "955300"
}

func (r *Rule955300) Phase() int {
	return 4
}

func (r *Rule955300) Evaluate(tx *core.Transaction) int {
	if strings.Contains(tx.Variables.ResponseBody, "<title>punkholicshell</title>") {
		tx.Variables.OutboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
