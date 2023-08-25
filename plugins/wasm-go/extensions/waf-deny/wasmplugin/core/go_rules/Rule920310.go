package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
)

// Include 920311

var UA_920310 = []string{"AppleWebKit", "Android", "Business", "Enterprise", "Entreprise"}

var rule920310Matcher ahocorasick.AhoCorasick

type Rule920310 struct {
}

func (r *Rule920310) Id() string {
	return "920310"
}

func (r *Rule920310) Phase() int {
	return 1
}

func (r *Rule920310) Evaluate(tx core.Transaction) bool {
	accept, ok := tx.Variables.RequestHeaders["Accept"]
	if !ok || len(accept) > 0 {
		return true
	}

	if tx.Variables.RequestMethod == "OPTION" {
		return true
	}

	if m, _ := core.PmEvaluate(rule920310Matcher, tx.Variables.RequestHeaders["User-Agent"], false); !m {
		tx.Variables.InboundAnomalyScorePl1 += core.NOTICE_ANOMALY_SCORE
	}

	return true
}

func init() {
	rule920310Matcher = core.AHO_CORASICK_BUILDER.Build(UA_920310)
}
