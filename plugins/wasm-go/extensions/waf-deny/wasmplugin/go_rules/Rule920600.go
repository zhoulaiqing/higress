package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/wasilibs/go-re2"
)

const PTN_920600 = `^(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*(?:[\s\v]*,[\s\v]*(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*)*$`

var re920600 *re2.Regexp

type Rule920600 struct {
}

func (r *Rule920600) Id() string {
	return "920600"
}

func (r *Rule920600) Phase() int {
	return 1
}

func (r *Rule920600) Evaluate(tx *core.Transaction) bool {
	if matched := re920600.MatchString(tx.Variables.RequestHeaders["accept"]); matched {
		tx.Variables.InboundAnomalyScorePl1 += CRITICAL_ANOMALY_SCORE
	}

	return true
}

func init() {
	re920600, _ = re2.Compile(PTN_920600)
}
