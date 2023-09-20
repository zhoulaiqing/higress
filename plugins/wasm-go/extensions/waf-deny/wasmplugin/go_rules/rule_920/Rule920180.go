package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
)

type Rule920180 struct {
}

func (r *Rule920180) Id() string {
	return "920180"
}

func (r *Rule920180) Phase() int {
	return 1
}

func (r *Rule920180) Evaluate(tx *core.Transaction) int {
	if tx.Variables.RequestProtocol == "HTTP/2" || tx.Variables.RequestProtocol == "HTTP/2.0" {
		return rule_tasks.PASS
	}

	if tx.Variables.RequestMethod == "POST" && len(tx.Variables.RequestHeaders["content-length"]) == 0 &&
		len(tx.Variables.RequestHeaders["transfer-encoding"]) == 0 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
