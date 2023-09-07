package rule_941

import (
	"github.com/corazawaf/coraza-proxy-wasm/wasmplugin/rule_tasks"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMatchValue(t *testing.T) {
	rule941 := Rule941{}
	rk := rule941.matchValue("xyz", false)
	assert.False(t, rk)

	rv := rule941.matchValue("<script>alert('1');<script>", false)
	assert.True(t, rv)

	r := rule941.matchValue("xlink:href", false)
	assert.True(t, r)

	ra := rule_tasks.Re941130.MatchString(" pattern=_")
	assert.True(t, ra)

	r = rule941.matchValue("xlink:hreflang", false)
	assert.False(t, r)
}

func TestBatchMatchValue(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	rule941 := Rule941{}

	testGroup := map[string]testCase{
		"case1": testCase{"xyz", false},
		"case2": testCase{"<script>alert('1');<script>", true},
		"case3": testCase{"xlink:href", true},
		"case4": testCase{"xlink:hreflang", false},
		"case5": testCase{"formaction ", true},
		"case6": testCase{"pattern:m=abc", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			got := rule941.matchValue(v.value, false)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}
