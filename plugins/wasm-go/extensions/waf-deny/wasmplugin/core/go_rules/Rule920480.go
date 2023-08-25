package core

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
	"golang.org/x/exp/slices"
)

const PTN_920480 = `charset\s*=\s*[\"']?([^;\"'\s]+)`

var re920480 *re2.Regexp

type Rule920480 struct {
}

func (r *Rule920480) Id() string {
	return "920480"
}

func (r *Rule920480) Phase() int {
	return 1
}

func (r *Rule920480) Evaluate(tx core.Transaction) bool {
	match := re920480.FindStringSubmatch(tx.Variables.RequestHeaders["Content-Type"])
	if len(match) < 2 {
		return true
	}

	charset := match[1]
	if !slices.Contains(core.ALLOWED_REQUEST_CONTENT_TYPE_CHARSET, charset) {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}

func init() {
	re920480, _ = re2.Compile(PTN_920480)
}
