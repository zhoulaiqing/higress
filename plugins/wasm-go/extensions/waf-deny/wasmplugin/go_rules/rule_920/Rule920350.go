package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_920350 = "(?:^([\\d.]+|\\[[\\da-f:]+\\]|[\\da-f:]+)(:[\\d]+)?$)"

var re920350 *re2.Regexp

type Rule920350 struct {
}

func (r *Rule920350) Id() string {
	return "920350"
}

func (r *Rule920350) Phase() int {
	return 1
}

func (r *Rule920350) Evaluate(tx *core.Transaction) int {
	matched := re920350.MatchString(tx.Variables.RequestHeaders["host"])

	if matched {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re920350, _ = re2.Compile(PTN_920350)
}
