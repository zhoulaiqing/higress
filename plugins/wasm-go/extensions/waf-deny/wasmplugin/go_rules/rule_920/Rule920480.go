package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
	"golang.org/x/exp/slices"
)

const PTN_920480 = `charset\s*=\s*[\"']?([^;\"'\s]+)`

var re920480 *re2.Regexp

type Rule920480 struct {
}

func (r *Rule920480) Id() string {
	return "920480"
}

func (r *Rule920480) Phase() int {
	return 1
}

func (r *Rule920480) Evaluate(tx *core.Transaction) int {
	match := re920480.FindStringSubmatch(tx.Variables.RequestHeaders["content-type"])
	if len(match) < 2 {
		return rule_tasks.PASS
	}

	charset := match[1]
	if !slices.Contains(rule_tasks.ALLOWED_REQUEST_CONTENT_TYPE_CHARSET, charset) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re920480, _ = re2.Compile(PTN_920480)
}
