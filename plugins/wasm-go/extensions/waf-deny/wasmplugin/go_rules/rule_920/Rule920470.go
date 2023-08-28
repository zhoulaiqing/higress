package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules"
	"github.com/wasilibs/go-re2"
	"strings"
)

const PTN_920470 = `^[\w/.+*-]+(?:\s?;\s?(?:action|boundary|charset|component|start(?:-info)?|type|version)\s?=\s?['\"\w.()+,/:=?<>@#*-]+)*$`

var re920470 *re2.Regexp

type Rule920470 struct {
}

func (r *Rule920470) Id() string {
	return "920470"
}

func (r *Rule920470) Phase() int {
	return 1
}

func (r *Rule920470) Evaluate(tx *core.Transaction) bool {
	contentType, ok := tx.Variables.RequestHeaders["content-type"]
	if !ok {
		return true
	}

	contentType = strings.ToLower(contentType)

	if m := re920470.MatchString(contentType); !m {
		tx.Variables.InboundAnomalyScorePl1 += go_rules.CRITICAL_ANOMALY_SCORE
	}

	return true
}

func init() {
	re920470, _ = re2.Compile(PTN_920470)
}
