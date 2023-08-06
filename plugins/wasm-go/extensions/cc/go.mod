module cc

go 1.19

require (
	github.com/alibaba/higress/plugins/wasm-go v0.0.0
	github.com/tetratelabs/proxy-wasm-go-sdk v0.22.0
	github.com/tidwall/gjson v1.14.4
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
)

require cc_tools v0.0.0
require qps_limiter v0.0.0
require sliding_hist_limiter v0.0.0

replace cc_tools => ./cc_tools
replace qps_limiter => ./qps_limiter
replace sliding_hist_limiter => ./sliding_hist_limiter

replace (
	github.com/alibaba/higress/plugins/wasm-go => ../..
)
