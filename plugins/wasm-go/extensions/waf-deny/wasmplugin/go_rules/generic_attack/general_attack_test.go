package generic_attack

import (
	"reflect"
	"strings"
	"testing"
)

func TestMatchDefault(t *testing.T) {
	type testCase struct {
		value string
		want  bool //期望的计算结果
	}

	testGroup := map[string]testCase{
		"case934100-1":  {"_%24%24ND_FUNC%24%24_", true},
		"case934100-2":  {"__js_function", true},
		"case934100-3":  {"eval%28String.fromCharCode", true},
		"case934100-4":  {"function%28%29+%7B", true},
		"case934100-5":  {"new+Function+%28", true},
		"case934100-6":  {"this.constructor.constructor", true},
		"case934100-7":  {"module.exports%3D", true},
		"case934100-8":  {"XyQkTkRfRlVOQyQkXwo=", true},
		"case934100-10": {"process.env", true},
		"case934100-11": {"console.info(1)", true},
		"case934100-12": {"c%5Cu006fnsole.info(1)", true},
		"case934100-13": {"process[\"env\"]", true},
		"case934100-14": {"console[\"info\"](1)", true},
		"case934100-15": {"console.info.call(this,1)", true},
		"case934100-16": {"process.", false},
		"case934100-17": {"console.", false},
		"case934100-18": {"%23%7Bprocess.binding%28foo%29.spawn%28foo2%29%7D", true},
		"case934100-19": {"%23%7Brequire.main.constructor._load%28foo%29.readdirSync%28foo2%29%7D", true},
		"case934100-20": {"process%5Breq.query.a", true},
		"case934100-21": {"require%5Breq.query.a", true},
		"case934100-22": {"process%5BmainModule", true},
		"case934100-23": {"require(\"child_process\").exec('whoami')", true},
		"case934100-26": {"c\\u006fnsole.info(1)", true},
		"case934100-27": {"_$$ND_FUNC$$_function()", true},
		"case934100-28": {"_$$\\u004e\\u0044_FUNC$$_\\u0066unction()", true},
		"case934100-29": {"XyQkTkRfRlVOQyQkX2Z1bmN0aW9uKCkK", true},
		"case934100-30": {"XyQkXHUwMDRlXHUwMDQ0X0ZVTkMkJF9cdTAwNjZ1bmN0aW9uKCkK", true},
		"case934100-31": {"\\u0058\\u0079QkTkRfRlVOQyQkX2Z1bmN0aW9uKCkK", true},
		//"case934100-1":  {"", true},
		"case934130-4": {"__proto__.test=test", true},
		"case934130-5": {"constructor.prototype.test=test", true},

		"case934160-1":  {"while%20(foo)%20is%20bar.", false},
		"case934160-2":  {"while(!false)", true},
		"case934160-3":  {"while(!-0);", true},
		"case934160-4":  {"while(!%2B0);", true},
		"case934160-5":  {"while(!0);", true},
		"case934160-7":  {"while(!\"\");", true},
		"case934160-8":  {"while(!'');", true},
		"case934160-9":  {"while(!``);", true},
		"case934160-10": {"while(true);", true},
		"case934160-11": {"while(Infinity);", true},
		"case934160-12": {"while(-Infinity);", true},
		"case934160-13": {"while(%2BInfinity);", true},
		"case934160-14": {"while(-1);", true},
		"case934160-15": {"while(%2B1);", true},
		"case934160-16": {"while(new%20Date);", true},
		"case934160-17": {"while(this);", true},
		"case934160-18": {"while(String);", true},
		"case934160-19": {"while(!true);", false},
		"case934160-20": {"while(!!{});", true},
		"case934160-21": {"while(!![]);", true},
		"case934160-22": {`while(!!\"\");`, true},
		"case934160-23": {"while(!!'');", true},
		"case934160-24": {"while(!!``);", true},
		"case934160-25": {"while(!null);", true},
		"case934160-26": {"while(!undefined);", true},
		"case934160-27": {"while(!NaN)", true},
		"case934160-28": {"while((true)))", true},
	}

	for key, v := range testGroup { //遍历
		t.Run(key, func(t *testing.T) {
			tv := transformDefault(v.value)
			lowerTv := strings.ToLower(tv)
			tvd := transformWithBase64Decode(tv)
			got := matchDefault(lowerTv, tvd)
			want := v.want
			if !reflect.DeepEqual(want, got) { //比较
				t.Errorf("excepted:%v, got:%v", want, got)
			}
		})
	}
}
