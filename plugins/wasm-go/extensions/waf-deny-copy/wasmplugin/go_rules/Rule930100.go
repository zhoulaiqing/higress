package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/wasilibs/go-re2"
)

const PTN_930100 = `(?i)(?:[/\x5c]|%(?:2(?:f|5(?:2f|5c|c(?:1%259c|0%25af))|%46)|5c|c(?:0%(?:[2aq]f|5c|9v)|1%(?:[19p]c|8s|af))|(?:bg%q|(?:e|f(?:8%8)?0%8)0%80%a)f|u(?:221[5-6]|EFC8|F025|002f)|%3(?:2(?:%(?:%6|4)6|F)|5%%63)|1u)|0x(?:2f|5c))(?:\.(?:%0[0-1]|\?)?|\?\.?|%(?:2(?:(?:5(?:2|c0%25a))?e|%45)|c0(?:\.|%[25-6ae-f]e)|u(?:(?:ff0|002)e|2024)|%32(?:%(?:%6|4)5|E)|(?:e|f(?:(?:8|c%80)%8)?0%8)0%80%ae)|0x2e){2,3}(?:[/\x5c]|%(?:2(?:f|5(?:2f|5c|c(?:1%259c|0%25af))|%46)|5c|c(?:0%(?:[2aq]f|5c|9v)|1%(?:[19p]c|8s|af))|(?:bg%q|(?:e|f(?:8%8)?0%8)0%80%a)f|u(?:221[5-6]|EFC8|F025|002f)|%3(?:2(?:%(?:%6|4)6|F)|5%%63)|1u)|0x(?:2f|5c))`

var re930100 *re2.Regexp

type Rule930100 struct {
}

func (r *Rule930100) Id() string {
	return "930100"
}

func (r *Rule930100) Phase() int {
	return 2
}

func (r *Rule930100) Evaluate(tx *core.Transaction) int {
	if r.doEvaluate(tx, tx.Variables.RequestUriRaw) {
		return rule_tasks.BLOCK
	}

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

	for k, v := range tx.Variables.RequestHeaders {
		if k == "Refer" {
			continue
		}

		if r.doEvaluate(tx, k) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	for k, values := range tx.Variables.Files {
		if r.doEvaluate(tx, k) {
			return rule_tasks.BLOCK
		}
		for _, v := range values {
			if r.doEvaluate(tx, v) {
				return rule_tasks.BLOCK
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule930100) doEvaluate(tx *core.Transaction, value string) bool {
	m := re930100.MatchString(value)
	if m {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.LfiScore += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}

func init() {
	re930100, _ = re2.Compile(PTN_930100)
}