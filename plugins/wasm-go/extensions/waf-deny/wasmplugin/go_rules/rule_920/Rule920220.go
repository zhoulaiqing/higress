package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"strings"
)

type Rule920220 struct {
}

func (r *Rule920220) Id() string {
	return "920220"
}

func (r *Rule920220) Phase() int {
	return 1
}

func (r *Rule920220) Evaluate(tx *core.Transaction) int {
	requestUri := tx.Variables.RequestUri

	if strings.Contains(requestUri, "%") && core.ValidateURLEncoding(requestUri) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
