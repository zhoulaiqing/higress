package go_rules

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/core"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/generic_attack"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/php_attack"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/protocol_attack"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/rule_932"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/sqli"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/go_rules/xss"
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

var Rule941 = &xss.Rule941{}
var Rule921 = &protocol_attack.Rule921{}
var Rule934 = &generic_attack.Rule934{}

var (
	RULES = []Rule{
		&Rule911100{},
		&Rule913100{}, &Rule913120{},
		//&rule_920.Rule920100{}, &rule_920.Rule920120{}, &rule_920.Rule920160{}, &rule_920.Rule920170{}, &rule_920.Rule920180{},
		//&rule_920.Rule920181{}, &rule_920.Rule920190{}, &rule_920.Rule920210{}, &rule_920.Rule920220{}, &rule_920.Rule920240{},
		//&rule_920.Rule920250{}, &rule_920.Rule920260{}, &rule_920.Rule920270{}, &rule_920.Rule920280{}, &rule_920.Rule920310{},
		//&rule_920.Rule920330{}, &rule_920.Rule920340{}, &rule_920.Rule920350{}, &rule_920.Rule920420{}, &rule_920.Rule920430{},
		//&rule_920.Rule920440{}, &rule_920.Rule920450{}, &rule_920.Rule920470{}, &rule_920.Rule920480{}, &rule_920.Rule920500{},
		//&rule_920.Rule920520{}, &rule_920.Rule920530{}, &rule_920.Rule920540{}, &rule_920.Rule920600{}, &rule_920.Rule920610{},
		Rule921,
		&Rule922100{}, &Rule922110_120{},
		&Rule930100{}, &Rule930110{}, &Rule930120{}, &Rule930130{},
		&Rule931100{}, &Rule931110{}, &Rule931120{},
		&rule_932.Rule932230{}, &rule_932.Rule932115{}, &rule_932.Rule932125{}, &rule_932.Rule932130{},
		&rule_932.Rule932140{}, &rule_932.Rule932170{}, &rule_932.Rule932171{}, &rule_932.Rule932175{},
		&rule_932.Rule932180{}, &rule_932.Rule932230{}, &rule_932.Rule932235{}, &rule_932.Rule932250{}, &rule_932.Rule932260{},
		&rule_932.Rule932330{}, &rule_932.Rule932370{},
		&php_attack.Rule933100{}, &php_attack.Rule933110{}, &php_attack.Rule933140{},
		&php_attack.Rule933150{}, &php_attack.Rule933160{}, &php_attack.Rule933170{}, &php_attack.Rule933180{}, &php_attack.Rule933200{},
		&php_attack.Rule933210{},
		Rule934,
		&xss.Rule941010{}, Rule941,
		&sqli.Rule942100{},
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
		r := Rule921.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule921) {
			return false
		}

		r = Rule934.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule934) {
			return false
		}

		r = Rule941.EvaluatePhase1(tx)
		if !checkBlock(tx, r, Rule941) {
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
