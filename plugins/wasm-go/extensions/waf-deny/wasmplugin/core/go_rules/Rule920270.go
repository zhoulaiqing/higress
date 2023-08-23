package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

const BYTE_RANGE_920270 = "1-255"

var validateByteRange920270 *core.ValidateByteRange

type Rule920270 struct {
}

func (r *Rule920270) Id() string {
	return "920270"
}

func (r *Rule920270) Evaluate(tx core.Transaction) bool {
	if !validateByteRange920270.Validate(tx.Variables.RequestUri) {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
		return true
	}

	// todo
	return true
}

func init() {
	validateByteRange920270, _ = core.NewValidateByRange(BYTE_RANGE_920270)
}
