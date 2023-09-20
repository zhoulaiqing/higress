package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"golang.org/x/exp/slices"
)

type Rule920430 struct {
}

func (r *Rule920430) Id() string {
	return "920430"
}

func (r *Rule920430) Phase() int {
	return 1
}

func (r *Rule920430) Evaluate(tx *core.Transaction) int {
	if !slices.Contains(rule_tasks.ALLOWED_HTTP_VERSIONS, tx.Variables.RequestProtocol) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
