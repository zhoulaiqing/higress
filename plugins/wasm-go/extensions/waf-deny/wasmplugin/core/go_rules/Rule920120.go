package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

type Rule920120 struct {
}

const PTN_920120 = "(?i)^(?:&(?:(?:[acegiln-or-suz]acut|[aeiou]grav|[ain-o]tild)e|[c-elnr-tz]caron|(?:[cgk-lnr-t]cedi|[aeiouy]um)l|[aceg-josuwy]circ|[au]ring|a(?:mp|pos)|nbsp|oslash);|[^\\\"';=])*$"

func (r *Rule920120) Id() string {
	return "920120"
}

func (r *Rule920120) Evaluate(tx core.Transaction) bool {

	for fk, _ := range tx.Variables.Files {
		m, _ := hyperscan.MatchString(PTN_920120, fk)
		if m {
			return true
		}
	}

	for fk, _ := range tx.Variables.FilesNames {
		m, _ := hyperscan.MatchString(PTN_920120, fk)
		if m {
			return true
		}
	}

	tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	return true
}
