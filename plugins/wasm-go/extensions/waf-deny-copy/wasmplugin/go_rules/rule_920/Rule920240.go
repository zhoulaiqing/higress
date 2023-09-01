package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule920240 struct {
}

func (r *Rule920240) Id() string {
	return "920240"
}

func (r *Rule920240) Phase() int {
	return 2
}

func (r *Rule920240) Evaluate(tx *core.Transaction) int {
	if !strings.HasPrefix(tx.Variables.RequestHeaders["content-type"], "application/x-www-form-urlencoded") {
		return rule_tasks.PASS
	}

	requestBody := tx.Variables.RequestBody
	if strings.Contains(requestBody, "%") && core.ValidateURLEncoding(requestBody) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
