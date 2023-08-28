package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule920610 struct {
}

func (r *Rule920610) Id() string {
	return "920610"
}

func (r *Rule920610) Phase() int {
	return 1
}

func (r *Rule920610) Evaluate(tx *core.Transaction) int {

	if strings.Contains(tx.Variables.RequestUriRaw, "#") {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
