package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
	"strings"
)

const PTN_921120 = `[\r\n]\W*?(?:content-(?:type|length)|set-cookie|location):\s*\w`

var re921120 *re2.Regexp

type Rule921120 struct {
}

func (r *Rule921120) Id() string {
	return "921120"
}

func (r *Rule921120) Phase() int {
	return 2
}

func (r *Rule921120) Evaluate(tx *core.Transaction) bool {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, k) {
			return true
		}

		if r.doEvaluate(tx, v) {
			return true
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if r.doEvaluate(tx, k) {
				return true
			}

			if r.doEvaluate(tx, v) {
				return true
			}
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return true
		}
	}

	return true
}

func (r *Rule921120) doEvaluate(tx *core.Transaction, value string) bool {
	m := re921120.MatchString(strings.ToLower(value))

	if m {
		tx.Variables.HttpViolationScore += CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return m
}

func init() {
	re921120, _ = re2.Compile(PTN_921120)
}
