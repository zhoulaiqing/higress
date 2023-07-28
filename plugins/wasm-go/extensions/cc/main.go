package main

import (
	"fmt"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println("Hello world.")
}

type CCConfig struct {
	headerRulesMap map[string]*CCRule
	cookieRulesMap map[string]*CCRule
}

type CCRule struct {
	maxPeriod int
	subRules  []CCSubRule
}

type CCSubRule struct {
	limitCount   int
	limitPeriod  int
	blockSeconds int
}

func parseConfig(json gjson.Result, config *CCConfig, log wrapper.Log) error {

	config.headerRulesMap = make(map[string]*CCRule)
	config.cookieRulesMap = make(map[string]*CCRule)

	rulesArray := json.Get("cc_rules").Array()
	for _, rule := range rulesArray {
		header := rule.Get("header")
		cookie := rule.Get("cookie")

		// header 和 cookie 都不存在，无效的配置项
		if !header.Exists() && !cookie.Exists() {
			continue
		}

		var rulesMap map[string]*CCRule
		var key string
		if header.Exists() {
			rulesMap = config.headerRulesMap
			key = header.String()
		} else {
			rulesMap = config.cookieRulesMap
			key = cookie.String()
		}

		var ccSubRules []CCSubRule
		var maxPeriod int
		blockSeconds := 0
		blockSecondsConf := rule.Get("block_seconds")
		if blockSecondsConf.Exists() {
			blockSeconds = int(blockSecondsConf.Int())
		}

		qps := rule.Get("qps")
		qpm := rule.Get("qpm")
		qpd := rule.Get("qpd")

		if qps.Exists() {
			period := 1000
			ccSubRules = append(ccSubRules, CCSubRule{int(qps.Int()), period, blockSeconds})
			maxPeriod = period
		}
		if qpm.Exists() {
			period := 60 * 1000
			ccSubRules = append(ccSubRules, CCSubRule{int(qpm.Int()), period, blockSeconds})
			maxPeriod = period
		}
		if qpd.Exists() {
			period := 24 * 60 * 60 * 1000
			ccSubRules = append(ccSubRules, CCSubRule{int(qpd.Int()), period, blockSeconds})
			maxPeriod = period
		}

		rulesMap[key] = &CCRule{maxPeriod, ccSubRules}
	}

	return nil
}
