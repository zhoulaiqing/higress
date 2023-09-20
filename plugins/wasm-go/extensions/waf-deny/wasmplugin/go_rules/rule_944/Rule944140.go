package rule_944

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"strings"
)

type Rule944140 struct {
}

func (r *Rule944140) Id() string {
	return "944140"
}

func (r *Rule944140) Phase() int {
	return 2
}

func (r *Rule944140) Evaluate(tx *core.Transaction) int {
	for _, values := range tx.Variables.Files {
		for _, v := range values {
			if r.doEvaluate(tx, &v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for _, fileNameHeader := range rule_tasks.FILE_NAME_HEADERS {
		fn, ok := tx.Variables.RequestHeaders[fileNameHeader]
		if !ok {
			continue
		}

		if r.doEvaluate(tx, &fn) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule944140) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re944140.MatchString(strings.ToLower(*value))
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
