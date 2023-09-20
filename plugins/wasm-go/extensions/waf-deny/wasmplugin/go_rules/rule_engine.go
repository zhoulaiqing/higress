package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/generic_attack"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/php_attack"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/protocol_attack"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rce_attack"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_920"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/sqli"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/xss"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

var Rule941 = &xss.Rule941{}
var Rule921 = &protocol_attack.Rule921{}
var Rule933 = &php_attack.Rule933{}
var Rule934 = &generic_attack.Rule934{}
var Rule932 = &rce_attack.Rule932{}

var (
	RULES = []Rule{
		&Rule911100{},
		&Rule913100{}, &Rule913120{},
		&rule_920.Rule920120{},
		Rule921,
		&Rule922100{}, &Rule922110_120{},
		&Rule930100{}, &Rule930110{}, &Rule930120{}, &Rule930130{},
		&Rule931100{}, &Rule931110{}, &Rule931120{},
		Rule932, &rce_attack.Rule932170{}, &rce_attack.Rule932171{},
		Rule933, &php_attack.Rule933110{}, &php_attack.Rule933150{},
		Rule934,
		&xss.Rule941010{}, Rule941,
		&sqli.Rule942100{},
		//&rule_943.Rule943100{}, &rule_943.Rule943110_120{},
		//&rule_944.Rule944100{}, &rule_944.Rule944110{}, &rule_944.Rule944120{}, &rule_944.Rule944130{}, &rule_944.Rule944140{},
		//&rule_944.Rule944150{},
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

func ProcessResponseBodyRules(tx *core.Transaction) bool {
	return ProcessRulesByPhase(tx, 4)
}

func ProcessRulesByPhase(tx *core.Transaction, phase int) bool {

	for _, rule := range RULES {
		if rule.Phase() == phase {
			r := rule.Evaluate(tx)
			if !checkBlock(tx, r, rule) {
				return false
			}
		}
	}

	if phase == 1 {
		r := Rule941.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule941) {
			return false
		}

		r = Rule921.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule921) {
			return false
		}

		r = Rule933.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule933) {
			return false
		}

		r = Rule934.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule934) {
			return false
		}

		r = Rule932.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule934) {
			return false
		}

	}

	return true
}

func checkBlock(tx *core.Transaction, r int, rule Rule) bool {
	if r == rule_tasks.BLOCK {
		proxywasm.LogInfof("Blocked at Rule %s ", rule.Id())
		tx.Variables.Interrupted = true
		return false
	}

	return true
}
