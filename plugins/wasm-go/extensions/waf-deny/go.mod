module waf_deny

go 1.19

require (
	github.com/alibaba/higress/plugins/wasm-go v0.0.0-20230717123545-aa17e9598d1f // indirect
	github.com/corazawaf/coraza-wasilibs v0.0.0-20230620081031-05a5097dbea3
	github.com/corazawaf/coraza/v3 v3.0.2 // indirect
	github.com/stretchr/testify v1.8.4
	github.com/tetratelabs/proxy-wasm-go-sdk v0.22.0
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/wasilibs/nottinygc v0.4.0
)

require (
	github.com/corazawaf/libinjection-go v0.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kr/text v0.1.0 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/petar-dambovaliev/aho-corasick v0.0.0-20211021192214-5ab2d9280aa9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tetratelabs/wazero v1.2.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/wasilibs/go-aho-corasick v0.5.0 // indirect
	github.com/wasilibs/go-libinjection v0.4.0 // indirect
	github.com/wasilibs/go-re2 v1.2.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/binaryregexp v0.2.0 // indirect
)

require (
	github.com/corazawaf/coraza-proxy-wasm v0.1.1
	wasm_plugin v0.0.0
)

replace wasm_plugin => ./wasmplugin
