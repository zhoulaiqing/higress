package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/flier/gohs/hyperscan"
)

const PTN_920600 = `^(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*(?:[\s\v]*,[\s\v]*(?:(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)/(?:\*|[^!-\"\(-\),/:-\?\[-\]\{\}]+)|\*)(?:[\s\v]*;[\s\v]*(?:charset[\s\v]*=[\s\v]*\"?(?:iso-8859-15?|utf-8|windows-1252)\b\"?|(?:[^\s\v -\"\(-\),/:-\?\[-\]c\{\}]|c(?:[^!-\"\(-\),/:-\?\[-\]h\{\}]|h(?:[^!-\"\(-\),/:-\?\[-\]a\{\}]|a(?:[^!-\"\(-\),/:-\?\[-\]r\{\}]|r(?:[^!-\"\(-\),/:-\?\[-\]s\{\}]|s(?:[^!-\"\(-\),/:-\?\[-\]e\{\}]|e[^!-\"\(-\),/:-\?\[-\]t\{\}]))))))[^!-\"\(-\),/:-\?\[-\]\{\}]*[\s\v]*=[\s\v]*[^!\(-\),/:-\?\[-\]\{\}]+);?)*)*$`

type Rule920600 struct {
}

func (r *Rule920600) Id() string {
	return "920600"
}

func (r *Rule920600) Phase() int {
	return 1
}

func (r *Rule920600) Evaluate(tx core.Transaction) bool {
	if matched, _ := hyperscan.MatchString(PTN_920600, tx.Variables.RequestHeaders["Accept"]); matched {
		tx.Variables.InboundAnomalyScorePl1 += core.CRITICAL_ANOMALY_SCORE
	}

	return true
}
