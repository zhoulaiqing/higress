package go_rules

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"golang.org/x/exp/slices"
	"strings"
)

type Rule922100 struct {
}

func (r *Rule922100) Id() string {
	return "922100"
}

func (r *Rule922100) Phase() int {
	return 2
}

func (r *Rule922100) Evaluate(tx *core.Transaction) int {
	if len(tx.Variables.MultipartPartHeaders["_charset_"]) == 0 {
		return rule_tasks.PASS
	}

	for _, argMap := range tx.Variables.Args {
		value, ok := (*argMap)["_charset_"]
		if ok && !slices.Contains(rule_tasks.ALLOWED_REQUEST_CONTENT_TYPE_CHARSET, strings.ToLower(value)) {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}
