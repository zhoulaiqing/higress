package go_rules

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	ahocorasick "github.com/wasilibs/go-aho-corasick"
)

var SCANNERS_HEADERS = []string{
	"acunetix-product",
	"(acunetix web vulnerability scanner",
	"acunetix-scanning-agreement",
	"acunetix-user-agreement",
	"myvar=1234",
	"x-ratproxy-loop",
	"bytes=0-,5-0,5-1,5-2,5-3,5-4,5-5,5-6,5-7,5-8,5-9,5-10,5-11,5-12,5-13,5-14",
	"x-scanner",
}

var rule913110Matcher ahocorasick.AhoCorasick

type Rule913110 struct {
}

func (r *Rule913110) Id() string {
	return "913110"
}

func (r *Rule913110) Phase() int {
	return 1
}

func (r *Rule913110) Evaluate(tx *core.Transaction) int {
	matched := false

	for k, v := range tx.Variables.RequestHeaders {
		matched, _ = core.PmEvaluate(rule913110Matcher, k, false)
		if matched {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
		matched, _ = core.PmEvaluate(rule913110Matcher, v, false)
		if matched {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func init() {
	rule913110Matcher = rule_tasks.AHO_CORASICK_BUILDER.Build(SCANNERS_HEADERS)
}
