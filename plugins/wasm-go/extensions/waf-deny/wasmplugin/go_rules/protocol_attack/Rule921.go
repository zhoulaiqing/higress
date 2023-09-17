package protocol_attack

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/wasilibs/go-re2"
	"strings"
)

type Rule921 struct {
}

func (r *Rule921) Id() string {
	return "921"
}

func (r *Rule921) Phase() int {
	return 2
}

func (r *Rule921) Evaluate(tx *core.Transaction) int {
	for k, v := range tx.Variables.ArgsPost {

		if matchDefaultRisk(k) {
			return r.block(tx)
		}

		if matchDefaultRisk(v) {
			return r.block(tx)
		}
	}

	for _, v := range tx.Variables.XML["/*"] {
		if matchDefaultRisk(v) {
			return r.block(tx)
		}
	}

	return rule_tasks.PASS
}

var re921240 = re2.MustCompile(`unix:[^|]*\|`)

func (r *Rule921) EvaluatePhase1(tx *core.Transaction) int {
	if r.checkContentType(strings.ToLower(tx.Variables.RequestHeaders["content-type"])) {
		return rule_tasks.BLOCK
	}

	if re921240.MatchString(strings.ToLower(tx.Variables.RequestUri)) {
		return rule_tasks.BLOCK
	}

	_, ok := tx.Variables.RequestHeaders["cookie"]
	if ok {
		for k, v := range tx.Variables.RequestCookies {
			if strings.Contains(k, "__utm") {
				continue
			}

			if matchDefaultRisk(transformDefault(k)) {
				return r.block(tx)
			}

			if matchDefaultRisk(transformDefault(v)) {
				return r.block(tx)
			}
		}
	}

	for _, argMap := range tx.Variables.Args {
		for _, v := range *argMap {

			proxywasm.LogInfof("arg value: %s", v)
			if matchArgGet(transformDefault(v)) {
				return r.block(tx)
			}
		}
	}

	return rule_tasks.PASS
}

func (r *Rule921) block(tx *core.Transaction) int {
	tx.Variables.HttpViolationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
	tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE

	return rule_tasks.BLOCK
}

func (r *Rule921) checkContentType(contentType string) bool {
	index := strings.Index(contentType, "application/")
	if index <= 0 {
		return true
	}

	if strings.ContainsAny(contentType[:index], "\\t\\r\\n\\f,;") {
		return false
	}

	return true
}
