
//line xss_370.rl:1
package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)
var builder370 strings.Builder

var keys370 = []string{"self", "document", "this", "top", "window"}

// 0. None; 1. Passed word; 2. Passed left;
var step370 = 0

func appendWord370(s string) {
    builder370.WriteString(s)
}

func checkWord370() {
    if step370 >= 1 {
        return
    }

    s := builder370.String()
    fmt.Println(s)

    if slices.Contains(keys370, s) {
        step370 = 1
    } else {
        builder370.Reset()
    }
}

func checkLeft370() {
    if step370 >= 1 {
        step370 = 2
    } else {
        checkWord370()
        if step370 >= 1 {
            step370 = 2
        } else {
            builder370.Reset()
        }
    }
}

func checkRight370() bool {
    if step370 == 2 {
        return true
    } else {
        builder370.Reset()
        return false
    }
}

func matchXSS370(data []byte) bool {

//line xss_370.rl:58

//line xss_370.go:63
const xss_start int = 0
const xss_first_final int = 0
const xss_error int = -1

const xss_en_main int = 0


//line xss_370.rl:59
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
            _ = act

    
//line xss_370.go:79
	{
	cs = xss_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_370.go:87
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	}
	goto st_out
tr0:
//line xss_370.rl:86
te = p+1
{
                if step370 < 2 {
                    builder370.Reset()
                }

            }
	goto st0
tr1:
//line xss_370.rl:82
te = p+1
{
                checkWord370()
            }
	goto st0
tr2:
//line xss_370.rl:72
te = p+1
{
                checkLeft370()
            }
	goto st0
tr6:
//line xss_370.rl:76
te = p+1
{
                if checkRight370() {
                    return true
                }
            }
	goto st0
tr7:
//line xss_370.rl:86
te = p
p--
{
                if step370 < 2 {
                    builder370.Reset()
                }

            }
	goto st0
tr8:
//line xss_370.rl:68
te = p
p--
{
                appendWord370(string(data[ts:te]))
            }
	goto st0
	st0:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
//line NONE:1
ts = p

//line xss_370.go:166
		switch data[p] {
		case 32:
			goto tr1
		case 41:
			goto tr2
		case 42:
			goto st1
		case 47:
			goto st2
		case 91:
			goto tr2
		case 93:
			goto tr6
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st3
			}
		default:
			goto st3
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 47 {
			goto tr6
		}
		goto tr7
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 42 {
			goto tr2
		}
		goto tr7
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st3
			}
		case data[p] >= 65:
			goto st3
		}
		goto tr8
	st_out:
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1:
			goto tr7
		case 2:
			goto tr7
		case 3:
			goto tr8
		}
	}

	}

//line xss_370.rl:96


    return false
}
