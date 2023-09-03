package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
	"strings"
)

const PTN_921160 = `[\n\r]+(?:\s|location|refresh|(?:set-)?cookie|(?:x-)?(?:forwarded-(?:for|host|server)|host|via|remote-ip|remote-addr|originating-IP))\s*:`

var re921160 *re2.Regexp

type Rule921160 struct {
}

func (r *Rule921160) Id() string {
	return "921160"
}

func (r *Rule921160) Phase() int {
	return 1
}

func (r *Rule921160) Evaluate(tx *core.Transaction) int {
	for k, v := range tx.Variables.ArgsGet {
		if r.doEvaluate(tx, k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule921160) doEvaluate(tx *core.Transaction, value string) bool {
	tv, _, _ := core.HtmlEntityDecode(value)
	m := re921160.MatchString(strings.ToLower(tv))

	if m {
		tx.Variables.HttpViolationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}

func init() {
	re921160, _ = re2.Compile(PTN_921160)
}