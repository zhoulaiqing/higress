
//line xss_390.rl:1
package rule_941

import (
    "fmt"
    "strings"
    "golang.org/x/exp/slices"
)

var builder390 strings.Builder
var keys390 = []string{"eval", "settimeout", "setinterval", "newfunction", "alert", "atob", "btoa"}

func appendWord390(s string) {
    builder390.WriteString(s)
}

func reset390 () {
    builder390.Reset()
    matchedWord = false
}

func checkSpace390() {
    if matchedWord {
        return
    }

    if checkWord390() >= 0 {
        return
    }

    reset390()
}

func checkWord390() int {
    s := builder390.String()
    fmt.Println(s)

    if slices.Contains(keys390, s) {
        matchedWord = true
        return 1
    } else if s == "new" {
        return 0
    }

    return -1
}

func checkFinish390() bool {
    if matchedWord {
        return true
    }

    if checkWord390() > 0 {
        return true
    }

    return false
}

var matchedWord = false

func matchXSS390(data []byte) bool {

//line xss_390.rl:63

//line xss_390.go:68
const xss_start int = 0
const xss_first_final int = 0
const xss_error int = -1

const xss_en_main int = 0


//line xss_390.rl:64
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _ = act

    
//line xss_390.go:84
	{
	cs = xss_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_390.go:92
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	}
	goto st_out
tr0:
//line xss_390.rl:90
te = p+1
{
                reset390()
            }
	goto st0
tr1:
//line xss_390.rl:86
te = p+1
{
                checkSpace390()
            }
	goto st0
tr2:
//line xss_390.rl:80
te = p+1
{
                if checkFinish390() {
                    return true
                }
            }
	goto st0
tr4:
//line xss_390.rl:76
te = p
p--
{
                appendWord390(string(data[ts:te]))
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

//line xss_390.go:146
		switch data[p] {
		case 32:
			goto tr1
		case 40:
			goto tr2
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		default:
			goto st1
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st1
			}
		case data[p] >= 65:
			goto st1
		}
		goto tr4
	st_out:
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1:
			goto tr4
		}
	}

	}

//line xss_390.rl:97


    return false
}
