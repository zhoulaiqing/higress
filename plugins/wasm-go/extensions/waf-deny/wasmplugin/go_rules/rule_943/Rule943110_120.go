package rule_943

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"strings"
)

type Rule943110_120 struct {
}

func (r *Rule943110_120) Id() string {
	return "943110_120"
}

func (r *Rule943110_120) Evaluate(tx *core.Transaction) int {

	match1 := false
	for _, argMap := range tx.Variables.Args {
		for k, _ := range *argMap {
			m := rule_tasks.Re9431101.MatchString(strings.ToLower(k))
			if m {
				match1 = true
				break
			}
		}
		if match1 {
			break
		}
	}

	if !match1 {
		return rule_tasks.PASS
	}

	// 943120
	referer, ok := tx.Variables.RequestHeaders["referer"]
	if !ok || len(referer) == 0 {
		tx.Variables.SessionFixationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	matches := rule_tasks.Re9431102.FindStringSubmatch(tx.Variables.RequestHeaders["referer"])
	if len(matches) < 2 {
		return rule_tasks.PASS
	}

	if !strings.HasSuffix(matches[1], tx.Variables.RequestHeaders["host"]) {
		tx.Variables.SessionFixationScore += rule_tasks.CRITICAL_ANOMALY_SCORE
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
