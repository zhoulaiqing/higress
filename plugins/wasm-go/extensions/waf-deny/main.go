// Copyright The OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	wasilibs "github.com/corazawaf/coraza-wasilibs"
	"wasm_plugin"
)

func main() {
	wasilibs.RegisterRX()
	wasilibs.RegisterPM()
	wasilibs.RegisterSQLi()
	wasilibs.RegisterXSS()
	wasmplugin.PluginStart()
}
