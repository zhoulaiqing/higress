package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
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

func (r *Rule920540) Evaluate(tx *core.Transaction) bool {
	if tx.Variables.ReqBodyProcessor == "JSON" {
		return true
	}

	if m := re920540.MatchString(tx.Variables.RequestUri); m {
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
		return true
	}

	for k, v := range tx.Variables.RequestHeaders {
		if m := re920540.MatchString(k); m {
			tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
			return true
		}
		if m := re920540.MatchString(v); m {
			tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
			return true
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if m := re920540.MatchString(k); m {
				tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
				return true
			}
			if m := re920540.MatchString(v); m {
				tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
				return true
			}
		}
	}

	return true
}

func init() {
	re920540, _ = re2.Compile(PTN_920540)
}
