package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	ahocorasick "github.com/wasilibs/go-aho-corasick"
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

func (r *Rule920310) Evaluate(tx *core.Transaction) int {
	accept, ok := tx.Variables.RequestHeaders["accept"]
	if !ok || len(accept) > 0 {
		return rule_tasks.PASS
	}

	if tx.Variables.RequestMethod == "OPTION" {
		return rule_tasks.PASS
	}

	if m, _ := core.PmEvaluate(rule920310Matcher, tx.Variables.RequestHeaders["user-agent"], false); !m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.NOTICE_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	rule920310Matcher = rule_tasks.AHO_CORASICK_BUILDER.Build(UA_920310)
}
