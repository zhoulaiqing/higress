package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
	"github.com/wasilibs/go-re2"
	"strings"
)

const PTN_922110_1 = `^content-type\s*:\s*(.*)$`
const PTN_922110_2 = `^(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*(?:[\s\v]*,[\s\v]*(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*)*$`

var re920110_1 *re2.Regexp

type Rule922110_120 struct {
}

func (r *Rule922110_120) Id() string {
	return "922110_120"
}

func (r *Rule922110_120) Evaluate(tx core.Transaction) bool {

	for k, values := range tx.Variables.MultipartPartHeaders {
		k = strings.ToLower(k)
		if strings.Contains(k, "content-transfer-encoding:") {
			tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
			return true
		}
		matchKey := re920110_1.FindStringSubmatch(k)
		if len(matchKey) >= 2 {
			if m, _ := hyperscan.MatchString(PTN_922110_2, matchKey[1]); m {
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}

		}

		for _, value := range values {
			value = strings.ToLower(value)
			if strings.Contains(value, "content-transfer-encoding:") {
				tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
				return true
			}

			match := re920110_1.FindStringSubmatch(value)
			if len(match) >= 2 {
				if m, _ := hyperscan.MatchString(PTN_922110_2, match[1]); m {
					tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
					return true
				}
			}
		}
	}

	return true
}

func init() {
	re920110_1, _ = re2.Compile(PTN_922110_1)
}
