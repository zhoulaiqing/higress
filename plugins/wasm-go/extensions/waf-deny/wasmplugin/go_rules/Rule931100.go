package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_931100 = `^(?i:file|ftps?|https?)://(?:\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`

var re931100 *re2.Regexp

type Rule931100 struct {
}

func (r *Rule931100) Id() string {
	return "931100"
}

func (r *Rule931100) Phase() int {
	return 2
}

func (r *Rule931100) Evaluate(tx *core.Transaction) int {

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if m := re931100.MatchString(k); m {
				tx.Variables.RfiScore += rule_tasks.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}

			if m := re931100.MatchString(v); m {
				tx.Variables.RfiScore += rule_tasks.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}
		}
	}

	return rule_tasks.PASS
}

func init() {
	re931100, _ = re2.Compile(PTN_931100)
}
