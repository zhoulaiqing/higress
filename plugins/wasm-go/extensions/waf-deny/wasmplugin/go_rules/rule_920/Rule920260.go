package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_920260 = `\%u[fF]{2}[0-9a-fA-F]{2}`

var re920260 *re2.Regexp

type Rule920260 struct {
}

func (r *Rule920260) Id() string {
	return "920260"
}

func (r *Rule920260) Phase() int {
	return 2
}

func (r *Rule920260) Evaluate(tx *core.Transaction) int {
	m1 := re920260.MatchString(tx.Variables.RequestUri)
	if m1 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	m2 := re920260.MatchString(tx.Variables.RequestBody)
	if m2 {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.WARNING_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re920260, _ = re2.Compile(PTN_920260)
}
