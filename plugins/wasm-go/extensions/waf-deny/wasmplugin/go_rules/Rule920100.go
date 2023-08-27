package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
)

const PTN_920100 = `(?i)^(?:get /[^#\?]*(?:\?[^\s\v#]*)?(?:#[^\s\v]*)?|(?:connect (?:(?:[0-9]{1,3}\.){3}[0-9]{1,3}\.?(?::[0-9]+)?|[\--9A-Z_a-z]+:[0-9]+)|options \*|[a-z]{3,10}[\s\v]+(?:[0-9A-Z_a-z]{3,7}?://[\--9A-Z_a-z]*(?::[0-9]+)?)?/[^#\?]*(?:\?[^\s\v#]*)?(?:#[^\s\v]*)?)[\s\v]+[\.-9A-Z_a-z]+)$`

var re920100 *re2.Regexp

type Rule920100 struct {
}

func (r *Rule920100) Id() string {
	return "920100"
}

func (r *Rule920100) Phase() int {
	return 1
}

func (r *Rule920100) Evaluate(tx *core.Transaction) bool {
	matched := re920100.MatchString(tx.Variables.RequestLine)

	if matched {
		tx.Variables.InboundAnomalyScorePl1 += WARNING_ANOMALY_SCORE
	}

	return true
}

func init() {
	re920100, _ = re2.Compile(PTN_920100)
}
