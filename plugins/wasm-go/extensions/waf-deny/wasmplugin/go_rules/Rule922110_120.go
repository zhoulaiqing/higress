package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
	"strings"
)

const PTN_922110_1 = `^content-type\s*:\s*(.*)$`
const PTN_922110_2 = `^(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*(?:[\s\v]*,[\s\v]*(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*)*$`

var re922110_1 *re2.Regexp
var re922110_2 *re2.Regexp

type Rule922110_120 struct {
}

func (r *Rule922110_120) Id() string {
	return "922110_120"
}

func (r *Rule922110_120) Phase() int {
	return 2
}

func (r *Rule922110_120) Evaluate(tx *core.Transaction) bool {

	for k, values := range tx.Variables.MultipartPartHeaders {
		k = strings.ToLower(k)
		if strings.Contains(k, "content-transfer-encoding:") {
			tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
			return true
		}
		matchKey := re922110_1.FindStringSubmatch(k)
		if len(matchKey) >= 2 {
			if m := re922110_2.MatchString(matchKey[1]); m {
				tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
				return true
			}

		}

		for _, value := range values {
			value = strings.ToLower(value)
			if strings.Contains(value, "content-transfer-encoding:") {
				tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
				return true
			}

			match := re922110_1.FindStringSubmatch(value)
			if len(match) >= 2 {
				if m := re922110_2.MatchString(match[1]); m {
					tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
					return true
				}
			}
		}
	}

	return true
}

func init() {
	re922110_1, _ = re2.Compile(PTN_922110_1)
	re922110_2, _ = re2.Compile(PTN_922110_2)
}
