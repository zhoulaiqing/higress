package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_931120 = `^(?i:file|ftps?|https?).*?\?+$`

var re931120 *re2.Regexp

type Rule931120 struct {
}

func (r *Rule931120) Id() string {
	return "931120"
}

func (r *Rule931120) Phase() int {
	return 2
}

func (r *Rule931120) Evaluate(tx *core.Transaction) int {
	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if m := re931120.MatchString(k); m {
				tx.Variables.RfiScore += rule_tasks.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}

			if m := re931120.MatchString(v); m {
				tx.Variables.RfiScore += rule_tasks.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}
		}
	}

	return rule_tasks.PASS
}

func init() {
	re931120, _ = re2.Compile(PTN_931120)
}
