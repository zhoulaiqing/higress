package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_920"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_932"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_933"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_934"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_941"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_942"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_943"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_944"
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
		&rule_932.Rule932230{}, &rule_932.Rule932115{}, &rule_932.Rule932120{}, &rule_932.Rule932125{}, &rule_932.Rule932130{},
		&rule_932.Rule932140{}, &rule_932.Rule932160{}, &rule_932.Rule932170{}, &rule_932.Rule932171{}, &rule_932.Rule932175{},
		&rule_932.Rule932180{}, &rule_932.Rule932230{}, &rule_932.Rule932235{}, &rule_932.Rule932250{}, &rule_932.Rule932260{},
		&rule_932.Rule932330{}, &rule_932.Rule932370{},
		&rule_933.Rule933100{}, &rule_933.Rule933110{}, &rule_933.Rule933120{}, &rule_933.Rule933130{}, &rule_933.Rule933140{},
		&rule_933.Rule933150{}, &rule_933.Rule933160{}, &rule_933.Rule933170{}, &rule_933.Rule933180{}, &rule_933.Rule933200{},
		&rule_933.Rule933210{},
		&rule_934.Rule934100{}, &rule_934.Rule934101{}, &rule_934.Rule934110{}, &rule_934.Rule934130{}, &rule_934.Rule934150{},
		&rule_934.Rule934160{}, &rule_934.Rule934170{},
		&rule_941.Rule941010{}, &rule_941.Rule941100{}, &rule_941.Rule941110{}, &rule_941.Rule941130{}, &rule_941.Rule941140{},
		&rule_941.Rule941160{}, &rule_941.Rule941170{}, &rule_941.Rule941180{}, &rule_941.Rule941190{}, &rule_941.Rule941200{},
		&rule_941.Rule941210{}, &rule_941.Rule941220{}, &rule_941.Rule941230{}, &rule_941.Rule941240{}, &rule_941.Rule941250{},
		&rule_941.Rule941260{}, &rule_941.Rule941270{}, &rule_941.Rule941280{}, &rule_941.Rule941290{}, &rule_941.Rule941300{},
		&rule_941.Rule941310{}, &rule_941.Rule941350{}, &rule_941.Rule941360{}, &rule_941.Rule941370_400{}, &rule_941.Rule941390{},
		&rule_942.Rule942100{}, &rule_942.Rule942140{}, &rule_942.Rule942151{}, &rule_942.Rule942160{}, &rule_942.Rule942170{},
		&rule_942.Rule942190{}, &rule_942.Rule942220{}, &rule_942.Rule942230{}, &rule_942.Rule942240{}, &rule_942.Rule942250{},
		&rule_942.Rule942270{}, &rule_942.Rule942280{}, &rule_942.Rule942290{}, &rule_942.Rule942320{}, &rule_942.Rule942350{},
		&rule_942.Rule942360{}, &rule_942.Rule942500{}, &rule_942.Rule942540{}, &rule_942.Rule942550{},
		&rule_943.Rule943100{}, &rule_943.Rule943110_120{},
		&rule_944.Rule944100{}, &rule_944.Rule944110{}, &rule_944.Rule944120{}, &rule_944.Rule944130{}, &rule_944.Rule944140{},
		&rule_944.Rule944150{},
		&RuleFinal{},
	}
)

type RuleEngine struct {
}

func ProcessRequestHeaderRules(tx *core.Transaction) bool {
	return ProcessRulesByPhase(tx, 1)
}

func ProcessRequestBodyRules(tx *core.Transaction) bool {
	return ProcessRulesByPhase(tx, 2)
}

func ProcessRulesByPhase(tx *core.Transaction, phase int) bool {

	finalRule := RULES[len(RULES)-1]

	for _, rule := range RULES {
		if rule.Phase() == phase {
			r := rule.Evaluate(tx)
			if r == rule_tasks.BLOCK {
				proxywasm.LogInfof("Blocked at Rule %q ", rule.Id())
				r = finalRule.Evaluate(tx)
			}
			if r == rule_tasks.DENY {
				tx.Variables.Interrupted = true
				return false
			}
		}
	}

	return true
}
