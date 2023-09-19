package php_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule933 struct {
}

func (r *Rule933) Id() string {
	return "933"
}

func (r *Rule933) Phase() int {
	return 2
}

func (r *Rule933) EvaluatePhase1(tx *core.Transaction) int {

	if matchRequestFileName(tx.Variables.RequestFileName) {
		return r.block(tx)
	}

	for _, v := range tx.Variables.RequestHeaders {
		if matchHeader(v) {
			return r.block(tx)
		}
	}

	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		if matchDefault(k) {
			return r.block(tx)
		}

		if matchDefault(v) {
			return r.block(tx)
		}
	}

	for k, v := range tx.Variables.ArgsGet {

		if matchDefault(k) {
			return r.block(tx)
		}

		if matchDefault(v) {
			return r.block(tx)
		}
	}

	return rule_tasks.PASS
}

func (r *Rule933) Evaluate(tx *core.Transaction) int {

	for k, v := range tx.Variables.ArgsPost {

		if matchDefault(k) {
			return r.block(tx)
		}

		if matchDefault(v) {
			return r.block(tx)
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if matchDefault(v) {
			return r.block(tx)
		}
	}

	return rule_tasks.PASS
}

func (r *Rule933) block(tx *core.Transaction) int {
	tx.Variables.PhpInjectionScore += rule_tasks.CRITICAL_ANOMALY_SCORE
	tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE

	return rule_tasks.BLOCK
}
