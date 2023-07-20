module waf-deny

go 1.19

require (
	github.com/alibaba/higress/plugins/wasm-go v0.0.0-20230717123545-aa17e9598d1f
	github.com/corazawaf/coraza/v3 v3.0.2
	github.com/tetratelabs/proxy-wasm-go-sdk v0.22.0
	github.com/tidwall/gjson v1.14.4
)

require (
	github.com/corazawaf/libinjection-go v0.1.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/magefile/mage v1.15.0 // indirect
	github.com/petar-dambovaliev/aho-corasick v0.0.0-20211021192214-5ab2d9280aa9 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	golang.org/x/net v0.11.0 // indirect
	rsc.io/binaryregexp v0.2.0 // indirect
)

require coraza_tools v0.0.0

replace (
	coraza_tools => ./coraza_tools
)

