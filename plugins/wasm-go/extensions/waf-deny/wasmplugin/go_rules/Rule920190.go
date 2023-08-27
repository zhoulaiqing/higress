package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"strconv"
	"strings"
)

var RangeHeaders = [2]string{"Range", "Request-Range"}

type Rule920190 struct {
}

func (r *Rule920190) Id() string {
	return "920190"
}

func (r *Rule920190) Phase() int {
	return 1
}

func (r *Rule920190) Evaluate(tx *core.Transaction) bool {
	for _, rangeHead := range RangeHeaders {
		requestRange := tx.Variables.RequestHeaders[rangeHead]
		b, num1, num2 := splitNumber(requestRange)
		if b && num2 < num1 {
			tx.Variables.InboundAnomalyScorePl1 += WARNING_ANOMALY_SCORE
			return true
		}
	}

	return true
}

func splitNumber(value string) (bool, int, int) {
	arr := strings.Split(value, "-")
	if len(arr) != 2 {
		return false, 0, 0
	}

	num1, err := strconv.Atoi(arr[0])
	if err != nil {
		return false, 0, 0
	}

	num2, err := strconv.Atoi(arr[2])
	if err != nil {
		return false, 0, 0
	}

	return true, num1, num2
}
