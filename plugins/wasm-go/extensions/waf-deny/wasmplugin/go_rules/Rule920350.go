package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
)

const PTN_920350 = "(?:^([\\d.]+|\\[[\\da-f:]+\\]|[\\da-f:]+)(:[\\d]+)?$)"

var re920350 *re2.Regexp

type Rule920350 struct {
}

func (r *Rule920350) Id() string {
	return "920350"
}

func (r *Rule920350) Phase() int {
	return 1
}

func (r *Rule920350) Evaluate(tx *core.Transaction) bool {
	matched := re920350.MatchString(tx.Variables.RequestHeaders["Host"])

	if matched {
		tx.Variables.InboundAnomalyScorePl1 += WARNING_ANOMALY_SCORE
	}

	return true
}

func init() {
	re920350, _ = re2.Compile(PTN_920350)
}
