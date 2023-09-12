
//line xss_210.rl:1
package rule_941

import (
    "fmt"
    "strings"
)
var builder strings.Builder

func checkColon() bool {
    s := builder.String()
    fmt.Println(s)

    if s == "javascript" || s == "vbscript" {
        fmt.Println("here")
        builder.Reset()
        return true
    }
    builder.Reset()
    return false
}

func checkHtmlSpace(word string) bool {
    fmt.Println(word)

    if word == "&newline;" || word == "&tab;" {
        return false
    } else if word == "&colon" {
        return checkColon()
    }

    builder.Reset()
    return false
}

func appendWord(s string) {
    builder.WriteString(s)
}

func matchXSS210(data []byte) bool {

//line xss_210.rl:41

//line xss_210.go:46
const xss_start int = 1
const xss_first_final int = 1
const xss_error int = -1

const xss_en_main int = 1


//line xss_210.rl:42
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _ = act

    
//line xss_210.go:62
	{
	cs = xss_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_210.go:70
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 0:
		goto st_case_0
	case 3:
		goto st_case_3
	}
	goto st_out
tr0:
//line xss_210.rl:74
p = (te) - 1
{
                builder.Reset()
            }
	goto st1
tr1:
//line xss_210.rl:55
te = p+1
{
                if checkHtmlSpace(string(data[ts:te])) {
                    return true
                }
            }
	goto st1
tr3:
//line xss_210.rl:74
te = p+1
{
                builder.Reset()
            }
	goto st1
tr4:
//line xss_210.rl:71
te = p+1
{
            }
	goto st1
tr6:
//line xss_210.rl:65
te = p+1
{
                if checkColon() {
                    return true
                }
            }
	goto st1
tr8:
//line xss_210.rl:74
te = p
p--
{
                builder.Reset()
            }
	goto st1
tr9:
//line xss_210.rl:61
te = p
p--
{
                appendWord(string(data[ts:te]))
            }
	goto st1
	st1:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line NONE:1
ts = p

//line xss_210.go:151
		switch data[p] {
		case 32:
			goto tr4
		case 38:
			goto tr5
		case 58:
			goto tr6
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st3
			}
		default:
			goto st3
		}
		goto tr3
tr5:
//line NONE:1
te = p+1

	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line xss_210.go:183
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st0
			}
		case data[p] >= 65:
			goto st0
		}
		goto tr8
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if data[p] == 59 {
			goto tr1
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st0
			}
		case data[p] >= 65:
			goto st0
		}
		goto tr0
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
		goto tr9
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 2:
			goto tr8
		case 0:
			goto tr0
		case 3:
			goto tr9
		}
	}

	}

//line xss_210.rl:83


    return false
}
