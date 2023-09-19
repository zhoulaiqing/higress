package rce_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule932 struct {
}

func (r *Rule932) Id() string {
	return "932"
}

func (r *Rule932) Phase() int {
	return 2
}

func (r *Rule932) EvaluatePhase1(tx *core.Transaction) int {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if r.doEvaluate(tx, k) {
			return r.block(tx)
		}

		if r.doEvaluate(tx, v) {
			return r.block(tx)
		}
	}

	for k, v := range tx.Variables.ArgsGet {
		if r.doEvaluate(tx, k) {
			return r.block(tx)
		}

		if r.doEvaluate(tx, v) {
			return r.block(tx)
		}
	}

	return rule_tasks.PASS
}

func (r *Rule932) Evaluate(tx *core.Transaction) int {

	for k, v := range tx.Variables.ArgsPost {
		if r.doEvaluate(tx, k) {
			return r.block(tx)
		}

		if r.doEvaluate(tx, v) {
			return r.block(tx)
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if r.doEvaluate(tx, v) {
			return r.block(tx)
		}
	}

	return rule_tasks.PASS
}

func (r *Rule932) doEvaluate(tx *core.Transaction, value string) bool {
	v, _, _ := core.UrlDecode(value)

	if rule_tasks.Re932230.MatchString(v) {
		return true
	}
	//if rule_tasks.Re932235.MatchString(v) {
	//	return true
	//}
	//if rule_tasks.Re932115.MatchString(v) {
	//	return true
	//}
	//if rule_tasks.Re932125.MatchString(v) {
	//	return true
	//}
	//if rule_tasks.Re932250.MatchString(v) {
	//	return true
	//}
	//matches := rule_tasks.Re9322601.FindStringSubmatch(value)
	//if len(matches) == 0 {
	//	return false
	//}
	//
	//if rule_tasks.Re9322602.MatchString(matches[0]) {
	//	return true
	//}
	//if rule_tasks.Re932230.MatchString(v) {
	//	return true
	//}
	//if rule_tasks.Re932175.MatchString(v) {
	//	return true
	//}
	//if rule_tasks.Re932370.MatchString(v) {
	//	return true
	//}
	//
	//vcmd, _, _ := core.CmdLine(v)
	//if rule_tasks.Re932130.MatchString(vcmd) {
	//	return true
	//}
	//if rule_tasks.Re932140.MatchString(vcmd) {
	//	return true
	//}
	//
	//vcmdnp, _, _ := core.NormalisePath(vcmd)
	//if m, _ := core.PmEvaluate(rule_tasks.Rule932Matcher, vcmdnp, false); m {
	//	return true
	//}

	return false
}

func (r *Rule932) block(tx *core.Transaction) int {
	tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
	tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE

	return rule_tasks.BLOCK
}
