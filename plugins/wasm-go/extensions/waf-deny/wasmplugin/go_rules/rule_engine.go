package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_920"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_932"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

var (
	RULES = []Rule{
		&Rule911100{},
		&Rule913100{}, &Rule913120{},
		&rule_920.Rule920100{}, &rule_920.Rule920120{}, &rule_920.Rule920160{}, &rule_920.Rule920170{}, &rule_920.Rule920180{},
		&rule_920.Rule920181{}, &rule_920.Rule920190{}, &rule_920.Rule920210{}, &rule_920.Rule920220{}, &rule_920.Rule920240{},
		&rule_920.Rule920250{}, &rule_920.Rule920260{}, &rule_920.Rule920270{}, &rule_920.Rule920280{}, &rule_920.Rule920310{},
		&rule_920.Rule920330{}, &rule_920.Rule920340{}, &rule_920.Rule920350{}, &rule_920.Rule920420{}, &rule_920.Rule920430{},
		&rule_920.Rule920440{}, &rule_920.Rule920450{}, &rule_920.Rule920470{}, &rule_920.Rule920480{}, &rule_920.Rule920500{},
		&rule_920.Rule920520{}, &rule_920.Rule920530{}, &rule_920.Rule920540{}, &rule_920.Rule920600{}, &rule_920.Rule920610{},
		&Rule921110{}, &Rule921120{}, &Rule921130{}, &Rule921140{}, &Rule921150{},
		&Rule921160{}, &Rule921190{}, &Rule921200{}, &Rule921240{}, &Rule921421{},
		&Rule922100{}, &Rule922110_120{},
		&Rule930100{}, &Rule930110{}, &Rule930120{}, &Rule930130{},
		&Rule931100{}, &Rule931110{}, &Rule931120{},
		&rule_932.Rule932230{}, &RuleFinal{},
	}
)

type RuleEngine struct {
}

func ProcessRequestHeaderRules(tx *core.Transaction) bool {
	lastInboundScore := 0
	for _, rule := range RULES {
		if rule.Phase() == 1 || rule.Phase() == 100 {
			r := rule.Evaluate(tx)
			if tx.Variables.InboundAnomalyScorePl1 != lastInboundScore {
				proxywasm.LogInfof("Rule %q, InboundScore diff %d .", rule.Id(), tx.Variables.InboundAnomalyScorePl1-lastInboundScore)
				lastInboundScore = tx.Variables.InboundAnomalyScorePl1
			}
			if r == rule_tasks.DENY {
				return false
			}
		}
	}

	return true
}
