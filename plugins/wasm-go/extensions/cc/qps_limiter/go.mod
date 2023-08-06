module qps_limiter

go 1.19

require github.com/tetratelabs/proxy-wasm-go-sdk v0.22.0

require (
	cc_tools v0.0.0
)

replace (
	cc_tools => ../cc_tools
)
