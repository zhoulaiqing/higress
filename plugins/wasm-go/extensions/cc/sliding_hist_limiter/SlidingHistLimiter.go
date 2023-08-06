package sliding_hist_limiter

import (
	"cc_tools"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"sort"
)

type CCSubRule struct {
	LimitCount   int64
	LimitPeriod  int64 // 毫秒级别
	BlockSeconds int64
}

type Identifier string
type History []int64

type RateLimitData struct {
	history    History
	cas        uint32
	blockUntil int64
}

func GetSubRule(period string, limit int64, blockSeconds int64) CCSubRule {
	var limitPeriod int64

	if period == "second" {
		limitPeriod = 1000
	} else if period == "minute" {
		limitPeriod = 1000 * 60
	} else if period == "day" {
		limitPeriod = 24 * 60 * 60 * 1000
	}

	return CCSubRule{
		LimitCount:   limit,
		LimitPeriod:  limitPeriod,
		BlockSeconds: blockSeconds,
	}
}

func loadData(id Identifier) (RateLimitData, error) {
	value, cas, err := proxywasm.GetSharedData(string(id))

	var rateLimitData RateLimitData

	if err != nil {
		rateLimitData = RateLimitData{
			History(make([]int64, 0)),
			0,
			0,
		}
	}

	data := cc_tools.ByteArrayToInt64Array(value)
	if len(data) == 0 {
		proxywasm.LogError("Get rate limit data error")

		rateLimitData = RateLimitData{
			History(make([]int64, 0)),
			cas,
			0,
		}
	} else {
		rateLimitData = RateLimitData{
			history:    History(data[0 : len(data)-1]),
			cas:        cas,
			blockUntil: data[len(data)-1],
		}
	}

	return rateLimitData, nil
}

func GetData(subRules []CCSubRule, id Identifier, timestamp int64) (RateLimitData, bool, error) {
	stop := false

	maxPeriod := int64(0)

	rateLimitData, _ := loadData(id)
	//proxywasm.LogInfof("Loaded rate limit data: %v", rateLimitData)

	for _, subRule := range subRules {
		if subRule.LimitCount < 0 {
			continue
		}

		if rateLimitData.blockUntil >= timestamp {
			stop = true
		}

		if maxPeriod < subRule.LimitPeriod {
			maxPeriod = subRule.LimitPeriod
		}

		count := countAfter(rateLimitData.history, timestamp-subRule.LimitPeriod)
		if count >= int(subRule.LimitCount) {
			stop = true
			if subRule.BlockSeconds > 0 {
				rateLimitData.blockUntil = timestamp + subRule.BlockSeconds
			}
			break
		}
	}

	rateLimitData.history = removeBefore(rateLimitData.history, timestamp-maxPeriod)
	rateLimitData.history = insert(rateLimitData.history, timestamp)

	return rateLimitData, stop, nil
}

func SaveData(id Identifier, rateLimitData RateLimitData, timestamp int64) {

	value := append(rateLimitData.history, rateLimitData.blockUntil)
	cas := rateLimitData.cas

	saved := false
	var err error
	// cas 保存，重试 10 次
	for i := 0; i < 10; i++ {
		// 保存顺序： time, usage, blockUntil
		buf := cc_tools.Int64ArrayToByteArray(value)
		err = proxywasm.SetSharedData(string(id), buf, cas)

		if err == nil {
			saved = true
			break
		} else if err == types.ErrorStatusCasMismatch {
			// cas 不匹配，重试
			buf, cas, err = proxywasm.GetSharedData(string(id))

			ret := cc_tools.ByteArrayToInt64Array(buf)
			hist := ret[0 : len(ret)-1]
			blockUtil := ret[len(ret)-1]
			hist = insert(hist, timestamp)
			value = append(hist, blockUtil)
		} else {
			break
		}
	}

	if !saved {
		proxywasm.LogErrorf("Could not save data : %v", err)
	}
}

// 移除 timestamp 以前的 history, timestamp 本身不会被移除
func removeBefore(hist History, timestamp int64) History {
	index := sort.Search(len(hist), func(i int) bool {
		return hist[i] >= timestamp
	})

	if index < len(hist) && hist[index] == timestamp {
		return hist[index:]
	}

	return hist[index:]
}

// 计算 timestamp 之后的 history 数量， timestamp 本身会被包括
func countAfter(hist History, timestamp int64) int {
	index := sort.Search(len(hist), func(i int) bool {
		return hist[i] >= timestamp
	})

	count := len(hist) - index

	return count
}

func insert(hist History, target int64) []int64 {
	// 使用 sort.Search 函数找到元素应该插入的位置
	index := sort.Search(len(hist), func(i int) bool {
		return hist[i] >= target
	})

	// 在该位置插入元素
	hist = append(hist, 0)             // 先扩展切片长度
	copy(hist[index+1:], hist[index:]) // 后移元素
	hist[index] = target               // 插入新元素

	return hist
}
