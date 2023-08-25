package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_931120 = `^(?i:file|ftps?|https?).*?\?+$`

type Rule931120 struct {
}

func (r *Rule931120) Id() string {
	return "931120"
}

func (r *Rule931120) Phase() int {
	return 2
}

func (r *Rule931120) Evaluate(tx core.Transaction) bool {
	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if m, _ := hyperscan.MatchString(PTN_931120, k); m {
				tx.Variables.RfiScore += core.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}

			if m, _ := hyperscan.MatchString(PTN_931120, v); m {
				tx.Variables.RfiScore += core.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}
		}
	}

	return true
}
