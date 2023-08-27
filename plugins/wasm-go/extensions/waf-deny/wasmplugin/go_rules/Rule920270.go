package go_rules

import "github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"

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

func (r *Rule920270) Evaluate(tx *core.Transaction) bool {

	decodeUri, _, _ := core.UrlDecodeUni(tx.Variables.RequestUri)
	if !validateByteRange920270.Validate(decodeUri) {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
		return true
	}

	for k, v := range tx.Variables.RequestHeaders {
		decodeK, _, _ := core.UrlDecodeUni(k)
		if !validateByteRange920270.Validate(decodeK) {
			tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
			return true
		}

		decodeV, _, _ := core.UrlDecodeUni(v)
		if !validateByteRange920270.Validate(decodeV) {
			tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
			return true
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			decodeK, _, _ := core.UrlDecodeUni(k)
			if !validateByteRange920270.Validate(decodeK) {
				tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
				return true
			}

			decodeV, _, _ := core.UrlDecodeUni(v)
			if !validateByteRange920270.Validate(decodeV) {
				tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
				return true
			}
		}
	}

	return true
}

func init() {
	validateByteRange920270, _ = core.NewValidateByRange(BYTE_RANGE_920270)
}
