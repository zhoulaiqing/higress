package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_931100 = `^(?i:file|ftps?|https?)://(?:\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`

type Rule931100 struct {
}

func (r *Rule931100) Id() string {
	return "931100"
}

func (r *Rule931100) Phase() int {
	return 2
}

func (r *Rule931100) Evaluate(tx core.Transaction) bool {

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if m, _ := hyperscan.MatchString(PTN_931100, k); m {
				tx.Variables.RfiScore += core.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}

			if m, _ := hyperscan.MatchString(PTN_931100, v); m {
				tx.Variables.RfiScore += core.CRITICAL_ANOMALY_SCORE
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}
		}
	}

	return true
}
