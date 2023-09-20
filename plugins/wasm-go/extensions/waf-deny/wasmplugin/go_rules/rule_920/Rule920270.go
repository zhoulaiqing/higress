package rule_920

import (
	"github.com/tianchi/waf/wasmplugin/core"
	"github.com/tianchi/waf/wasmplugin/rule_tasks"
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

	if !validateByteRange920270.Validate(tx.Variables.RequestUri) {
		tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
		return rule_tasks.BLOCK
	}

	for k, v := range tx.Variables.RequestHeaders {
		if !validateByteRange920270.Validate(k) {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}

		if !validateByteRange920270.Validate(v) {
			tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
			return rule_tasks.BLOCK
		}
	}

	for _, argMap := range tx.Variables.Args {
		for k, v := range *argMap {
			if !validateByteRange920270.Validate(k) {
				tx.Variables.InboundAnomalyScorePl1 += rule_tasks.CRITICAL_ANOMALY_SCORE
				return rule_tasks.BLOCK
			}

			if !validateByteRange920270.Validate(v) {
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
