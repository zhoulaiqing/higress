package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"golang.org/x/exp/slices"
)

type Rule920450 struct {
}

func (r *Rule920450) Id() string {
	return "920450"
}

func (r *Rule920450) Phase() int {
	return 1
}

func (r *Rule920450) Evaluate(tx *core.Transaction) int {
	for hk, _ := range tx.Variables.RequestHeaders {
		if slices.Contains(rule_tasks.RESTRICTED_HEADERS, hk) {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}
