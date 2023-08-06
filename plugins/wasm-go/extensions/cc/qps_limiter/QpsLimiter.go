package qps_limiter

import (
	"cc_tools"
	"fmt"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"time"
)

type Timestamps map[string]int64

func GetTimestamps(t time.Time) *Timestamps {
	ts := Timestamps{}

	ye, mo, da := t.Year(), t.Month(), t.Day()
	ho, mi, se, lo := t.Hour(), t.Minute(), t.Second(), t.Location()

	ts["now"] = t.Unix()
	ts["second"] = time.Date(ye, mo, da, ho, mi, se, 0, lo).Unix()
	ts["minute"] = time.Date(ye, mo, da, ho, mi, 0, 0, lo).Unix()
	ts["day"] = time.Date(ye, mo, da, 0, 0, 0, 0, lo).Unix()

	return &ts
}

type CCSubRule struct {
	LimitCount   int64
	Period       string
	BlockSeconds int64
}

type Usage struct {
	limit     int64
	remaining int64
	usage     int64
	blockUtil int64
	cas       uint32
}

type Identifier string

func GetSubRule(period string, limit int64, blockSeconds int64) CCSubRule {
	return CCSubRule{
		LimitCount:   limit,
		Period:       period,
		BlockSeconds: blockSeconds,
	}
}

// 返回 periodTime used, blockUntil, cas, err
func localPolicyUsage(id Identifier, period string) (int64, int64, int64, uint32, error) {
	cacheKey := getLocalKey(id, period)

	value, cas, err := proxywasm.GetSharedData(cacheKey)
	if err != nil {
		if err == types.ErrorStatusNotFound {
			return 0, 0, 0, 0, nil
		}
		return 0, 0, 0, 0, nil
	}

	ret := cc_tools.ByteArrayToInt64Array(value)
	if len(ret) != 3 {
		proxywasm.LogError("Shared data format error")
	}

	//proxywasm.LogInfof("Get shared data: %v:  %v, %v, %v", cacheKey, ret[0], ret[1], ret[2])

	return ret[0], ret[1], ret[2], cas, nil
}

func LocalPolicyIncrement(id Identifier, counters map[string]Usage, ts *Timestamps) {
	for period, usage := range counters {
		cacheKey := getLocalKey(id, period)

		value := usage.usage
		blockUntil := usage.blockUtil
		cas := usage.cas

		saved := false
		var err error
		// cas 保存，重试 10 次
		for i := 0; i < 10; i++ {
			// 保存顺序： time, usage, blockUntil
			buf := cc_tools.Int64ArrayToByteArray([]int64{(*ts)[period], value + 1, blockUntil})

			err = proxywasm.SetSharedData(cacheKey, buf, cas)
			if err == nil {
				saved = true
				break
			} else if err == types.ErrorStatusCasMismatch {
				// cas 不匹配，重试
				buf, cas, err = proxywasm.GetSharedData(cacheKey)
				ret := cc_tools.ByteArrayToInt64Array(buf)
				value = ret[1]
			} else {
				break
			}
		}

		if !saved {
			proxywasm.LogErrorf("Could not increment counter for Period '%v': %v", period, err)
		}
	}
}

func GetUsage(subRules []CCSubRule, id Identifier, ts *Timestamps) (map[string]Usage, string, error) {
	counters := make(map[string]Usage)
	stop := ""

	for _, subRule := range subRules {
		if subRule.LimitCount < 0 {
			continue
		}
		period := subRule.Period
		savedTime, curUsage, blockUntil, cas, err := localPolicyUsage(id, period)
		if err != nil {
			return counters, period, err
		}

		if blockUntil >= (*ts)["now"] {
			stop = period
		}

		curTime := (*ts)[period]
		var usage Usage
		if curTime != savedTime {
			// 需要刷新
			usage = Usage{
				limit:     subRule.LimitCount,
				remaining: subRule.LimitCount,
				usage:     0,
				cas:       cas,
				blockUtil: 0,
			}
		} else {
			remaining := subRule.LimitCount - curUsage
			usage = Usage{
				limit:     subRule.LimitCount,
				remaining: remaining,
				usage:     curUsage,
				cas:       cas,
				blockUtil: blockUntil,
			}

			if remaining <= 0 {
				stop = period
				if subRule.BlockSeconds > 0 {
					usage.blockUtil = (*ts)["now"] + subRule.BlockSeconds
				}
			}
		}

		counters[period] = usage
	}

	return counters, stop, nil
}

func getLocalKey(id Identifier, period string) string {
	return fmt.Sprintf("limit:%v:%v", id, period)
}
