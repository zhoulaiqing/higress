package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
	"golang.org/x/exp/slices"
)

const PTN_920420 = `^[^;\s]+`

var re920420 *re2.Regexp

type Rule920420 struct {
}

func (r *Rule920420) Id() string {
	return "920420"
}

func (r *Rule920420) Phase() int {
	return 1
}

func (r *Rule920420) Evaluate(tx *core.Transaction) int {

	matches := re920420.FindStringSubmatch(tx.Variables.RequestHeaders["content-type"])
	if len(matches) == 0 {
		return rule_tasks.PASS
	}

	if !slices.Contains(rule_tasks.ALLOWED_REQUEST_CONTENT_TYPE, matches[0]) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re920420, _ = re2.Compile(PTN_920420)
}