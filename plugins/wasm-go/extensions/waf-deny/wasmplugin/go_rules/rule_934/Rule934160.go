package rule_934

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

type Rule934160 struct {
	*Rule934
}

func (r *Rule934160) Id() string {
	return "934160"
}

func (r *Rule934160) validateFileName() bool {
	return true
}

// This is multimatch
func (r *Rule934160) doEvaluate(tx *core.Transaction, value string) bool {
	m := rule_tasks.Re934160.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ := core.UrlDecodeUni(value)
	m = rule_tasks.Re934160.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.Base64decode(v)
	m = rule_tasks.Re934160.MatchString(value)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return true
	}

	v, _, _ = core.ReplaceComments(v)
	m = rule_tasks.Re934160.MatchString(v)
	if m {
		tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
