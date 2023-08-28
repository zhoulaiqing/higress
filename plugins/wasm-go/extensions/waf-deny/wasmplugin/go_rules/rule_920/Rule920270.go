package rule_920

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
)

const BYTE_RANGE_920270 = "1-255"

var validateByteRange920270 *core.ValidateByteRange

type Rule920270 struct {
}

func (r *Rule920270) Id() string {
	return "920270"
}

func (r *Rule920270) Phase() int {
	return 2
}

func (r *Rule920270) Evaluate(tx *core.Transaction) int {

	decodeUri, _, _ := core.UrlDecodeUni(tx.Variables.RequestUri)
	if !validateByteRange920270.Validate(decodeUri) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	for k, v := range tx.Variables.RequestHeaders {
		decodeK, _, _ := core.UrlDecodeUni(k)
		if !validateByteRange920270.Validate(decodeK) {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}

		decodeV, _, _ := core.UrlDecodeUni(v)
		if !validateByteRange920270.Validate(decodeV) {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			decodeK, _, _ := core.UrlDecodeUni(k)
			if !validateByteRange920270.Validate(decodeK) {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}

			decodeV, _, _ := core.UrlDecodeUni(v)
			if !validateByteRange920270.Validate(decodeV) {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}
		}
	}

	return rule_tasks.PASS
}

func init() {
	validateByteRange920270, _ = core.NewValidateByRange(BYTE_RANGE_920270)
}
