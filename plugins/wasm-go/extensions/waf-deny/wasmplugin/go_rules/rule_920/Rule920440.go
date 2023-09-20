package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
	"golang.org/x/exp/slices"
	"strings"
)

type Rule920440 struct {
}

func (r *Rule920440) Id() string {
	return "920440"
}

func (r *Rule920440) Phase() int {
	return 1
}

func (r *Rule920440) Evaluate(tx *core.Transaction) int {

	requestBaseName := strings.TrimSpace(tx.Variables.RequestBaseName)

	dotIndex := strings.LastIndex(requestBaseName, ".")
	if dotIndex < 0 || dotIndex == len(requestBaseName)-1 {
		return rule_tasks.PASS
	}

	extension := strings.ToLower(requestBaseName[dotIndex:])
	if slices.Contains(rule_tasks.RESTRICTED_EXTENSIONS, extension) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	return rule_tasks.PASS
}
