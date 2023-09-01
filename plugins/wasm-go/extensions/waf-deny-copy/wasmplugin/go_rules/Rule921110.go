package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
	"strings"
)

const PTN_921110 = `(?:get|post|head|options|connect|put|delete|trace|track|patch|propfind|propatch|mkcol|copy|move|lock|unlock)\s+[^\s]+\s+http/\d`

var re921110 *re2.Regexp

type Rule921110 struct {
}

func (r *Rule921110) Id() string {
	return "921110"
}

func (r *Rule921110) Phase() int {
	return 2
}

func (r *Rule921110) Evaluate(tx *core.Transaction) int {
	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, k) {
				return rule_tasks.BLOCK
			}

			if r.doEvaluate(tx, v) {
				return rule_tasks.BLOCK
			}
		}
	}

	if r.doEvaluate(tx, tx.Variables.RequestBody) {
		return rule_tasks.BLOCK
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule921110) doEvaluate(tx *core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)
	m := re921110.MatchString(strings.ToLower(tv))

	if m {
		tx.Variables.HttpViolationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}

func init() {
	re921110, _ = re2.Compile(PTN_921110)
}
