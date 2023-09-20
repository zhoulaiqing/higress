package rule_943

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
)

type Rule943100 struct {
}

func (r *Rule943100) Id() string {
	return "943100"
}

func (r *Rule943100) Phase() int {
	return 2
}

func (r *Rule943100) Evaluate(tx *core.Transaction) int {
	// todo 由于这里的缓存和 942 规则中一致，故在这里复用 942 的缓存，后续可以考虑更换缓存名称

	cookieKeyCache := tx.Variables.TransMap["cookie_key_942"]
	cookieValueCache := tx.Variables.TransMap["cookie_value_942"]
	for i, _ := range cookieKeyCache {
		if r.doEvaluate(tx, &cookieKeyCache[i]) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, &cookieValueCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	argsKeyCache := tx.Variables.TransMap["arg_key_942"]
	argsValueCache := tx.Variables.TransMap["arg_value_942"]
	for i, _ := range argsKeyCache {
		if r.doEvaluate(tx, &argsKeyCache[i]) {
			return rule_tasks.BLOCK
		}

		if r.doEvaluate(tx, &argsValueCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	xmlCache := tx.Variables.TransMap["xml_942"]
	for i, _ := range xmlCache {
		if r.doEvaluate(tx, &xmlCache[i]) {
			return rule_tasks.BLOCK
		}
	}

	return rule_tasks.PASS
}

func (r *Rule943100) doEvaluate(tx *core.Transaction, value *string) bool {
	m := rule_tasks.Re943100.MatchString(*value)

	if m {
		tx.Variables.SessionFixationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
	}

	return m
}
