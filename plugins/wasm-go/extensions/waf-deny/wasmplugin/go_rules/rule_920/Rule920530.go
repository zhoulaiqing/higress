package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_920530 = `charset.*?charset`

var re920530 *re2.Regexp

type Rule920530 struct {
}

func (r *Rule920530) Id() string {
	return "920530"
}

func (r *Rule920530) Phase() int {
	return 1
}

func (r *Rule920530) Evaluate(tx *core.Transaction) int {
	if m := re920530.MatchString(tx.Variables.RequestHeaders["content-type"]); m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re920530, _ = re2.Compile(PTN_920530)
}
