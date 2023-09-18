package generic_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule934 struct {
}

func (r *Rule934) Id() string {
	return "934"
}

func (r *Rule934) Phase() int {
	return 2
}

func (r *Rule934) EvaluatePhase1(tx *core.Transaction) int {
	for k, v := range tx.Variables.RequestCookies {
		if strings.Contains(k, "__utm") {
			continue
		}

		tk := transformDefault(k)
		tkd := transformWithBase64Decode(tk)
		if matchDefault(strings.ToLower(tk), tkd) {
			return r.block(tx)
		}

		tv := transformDefault(v)
		tvd := transformWithBase64Decode(tv)
		if matchDefault(strings.ToLower(tv), tvd) {
			return r.block(tx)
		}
	}

	for k, v := range tx.Variables.ArgsGet {
		tk := transformDefault(k)
		tkd := transformWithBase64Decode(tk)
		if matchDefault(strings.ToLower(tk), tkd) {
			return r.block(tx)
		}

		tv := transformDefault(v)
		tvd := transformWithBase64Decode(tv)
		if matchDefault(strings.ToLower(tv), tvd) {
			return r.block(tx)
		}
	}

	return rule_tasks.PASS
}

func (r *Rule934) Evaluate(tx *core.Transaction) int {
	for k, v := range tx.Variables.ArgsPost {
		tk := transformDefault(k)
		tkd := transformWithBase64Decode(tk)
		if matchDefault(strings.ToLower(tk), tkd) {
			return r.block(tx)
		}

		tv := transformDefault(v)
		tvd := transformWithBase64Decode(tv)
		if matchDefault(strings.ToLower(tv), tvd) {
			return r.block(tx)
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		tv := transformDefault(v)
		tvd := transformWithBase64Decode(tv)
		if matchDefault(strings.ToLower(tv), tvd) {
			return r.block(tx)
		}
	}

	tv := transformDefault(tx.Variables.RequestFileName)
	tvd := transformWithBase64Decode(tv)
	if matchDefault(strings.ToLower(tv), tvd) {
		return r.block(tx)
	}

	return rule_tasks.PASS
}

func (r *Rule934) block(tx *core.Transaction) int {
	tx.Variables.RceScore += rule_tasks.CRITICAL_ANOMALY_SCORE
	tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE

	return rule_tasks.BLOCK
}
