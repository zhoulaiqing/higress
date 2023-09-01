package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_921190 = `[\n\r]`

var re921190 *re2.Regexp

type Rule921190 struct {
}

func (r *Rule921190) Id() string {
	return "921190"
}

func (r *Rule921190) Phase() int {
	return 1
}

func (r *Rule921190) Evaluate(tx *core.Transaction) int {
	requestFileName := tx.Variables.RequestFileName
	requestFileName, _, _ = core.UrlDecodeUni(requestFileName)

	if m := re921190.MatchString(requestFileName); m {
		tx.Variables.HttpViolationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}

func init() {
	re921190, _ = re2.Compile(PTN_921190)
}
