module cc

go 1.20

require (
	github.com/alibaba/higress/plugins/wasm-go v0.0.0-20230708135159-28892cf3ae0f
	github.com/tetratelabs/proxy-wasm-go-sdk v0.22.0
	github.com/tidwall/gjson v1.14.4
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
)

require cc_tools v0.0.0

replace cc_tools => ./cc_tools
