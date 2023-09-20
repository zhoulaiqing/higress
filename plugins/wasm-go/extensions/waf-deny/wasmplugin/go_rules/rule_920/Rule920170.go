package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"strings"
)

type Rule920170 struct {
}

func (r *Rule920170) Id() string {
	return "920170"
}

func (r *Rule920170) Phase() int {
	return 1
}

func (r *Rule920170) Evaluate(tx *core.Transaction) int {
	if tx.Variables.RequestMethod != "GET" && tx.Variables.RequestMethod != "HEAD" {
		return rule_tasks.PASS
	}

	contentLength := tx.Variables.RequestHeaders["content-length"]

	if len(contentLength) > 0 && strings.TrimSpace(contentLength) != "0" {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}
	if len(tx.Variables.RequestHeaders["transfer-encoding"]) > 0 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
