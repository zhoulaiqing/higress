package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
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

func (r *Rule920440) Evaluate(tx core.Transaction) bool {

	requestBaseName := strings.TrimSpace(tx.Variables.RequestBaseName)

	dotIndex := strings.LastIndex(requestBaseName, ".")
	if dotIndex < 0 || dotIndex == len(requestBaseName)-1 {
		return true
	}

	extension := strings.ToLower(requestBaseName[dotIndex:])
	if slices.Contains(core.RESTRICTED_EXTENSIONS, extension) {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
