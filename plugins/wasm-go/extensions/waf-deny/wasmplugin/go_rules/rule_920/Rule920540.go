package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_920540 = `(?i)\x5cu[0-9a-f]{4}`

var re920540 *re2.Regexp

type Rule920540 struct {
}

func (r *Rule920540) Id() string {
	return "920540"
}

func (r *Rule920540) Phase() int {
	return 2
}

func (r *Rule920540) Evaluate(tx *core.Transaction) int {
	if tx.Variables.ReqBodyProcessor == "JSON" {
		return rule_tasks.PASS
	}

	if m := re920540.MatchString(tx.Variables.RequestUri); m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	for k, v := range tx.Variables.RequestHeaders {
		if m := re920540.MatchString(k); m {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
		if m := re920540.MatchString(v); m {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if m := re920540.MatchString(k); m {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}
			if m := re920540.MatchString(v); m {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}
		}
	}

	return rule_tasks.PASS
}

func init() {
	re920540, _ = re2.Compile(PTN_920540)
}
