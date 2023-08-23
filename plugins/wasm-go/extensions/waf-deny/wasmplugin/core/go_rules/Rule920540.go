package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_920540 = `(?i)\x5cu[0-9a-f]{4}`

type Rule920540 struct {
}

func (r *Rule920540) Id() string {
	return "920540"
}

func (r *Rule920540) Evaluate(tx core.Transaction) bool {
	if tx.Variables.ReqBodyProcessor == "JSON" {
		return true
	}

	if m, _ := hyperscan.MatchString(PTN_920540, tx.Variables.RequestUri); m {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	for k, v := range tx.Variables.RequestHeaders {
		if m, _ := hyperscan.MatchString(PTN_920540, k); m {
			tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
			return true
		}
		if m, _ := hyperscan.MatchString(PTN_920540, v); m {
			tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
			return true
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if m, _ := hyperscan.MatchString(PTN_920540, k); m {
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}
			if m, _ := hyperscan.MatchString(PTN_920540, v); m {
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}
		}
	}

	return true
}
