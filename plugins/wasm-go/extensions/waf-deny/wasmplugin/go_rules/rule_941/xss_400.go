
//line xss_400.rl:1
package rule_941

import (
    "fmt"
)

func matchXSS400(data []byte) bool {

//line xss_400.rl:9

//line xss_400.go:14
const xss_start int = 21
const xss_first_final int = 21
const xss_error int = -1

const xss_en_main int = 21


//line xss_400.rl:10
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _, _, _ = ts, te, act

    step := 0

    
//line xss_400.go:32
	{
	cs = xss_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_400.go:40
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 23:
		goto st_case_23
	case 11:
		goto st_case_11
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 28:
		goto st_case_28
	}
	goto st_out
tr0:
//line NONE:1
	switch act {
	case 1:
	{p = (te) - 1

                fmt.Printf("match step1 %s \n", string(data[ts:te]))
                step = 1
            }
	case 2:
	{p = (te) - 1

                if step == 1 {
                    fmt.Printf("match step2 %s \n", string(data[ts:te]))
                    step = 2
                } else {
                    fmt.Println("back when match 2")
                    step = 0
                }
            }
	case 3:
	{p = (te) - 1

                if step == 2 {
                    fmt.Printf("match step3 %s \n", string(data[ts:te]))
                    return true
                } else {
                    fmt.Println("back when match 3")
                    step = 0
                }
            }
	default:
	{p = (te) - 1
}
	}
	
	goto st21
tr15:
//line xss_400.rl:56
p = (te) - 1

	goto st21
tr25:
//line xss_400.rl:56
te = p+1

	goto st21
tr30:
//line xss_400.rl:56
te = p
p--

	goto st21
tr31:
//line xss_400.rl:31
te = p
p--
{
                fmt.Printf("match step1 %s \n", string(data[ts:te]))
                step = 1
            }
	goto st21
tr33:
//line xss_400.rl:46
te = p
p--
{
                if step == 2 {
                    fmt.Printf("match step3 %s \n", string(data[ts:te]))
                    return true
                } else {
                    fmt.Println("back when match 3")
                    step = 0
                }
            }
	goto st21
	st21:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line NONE:1
ts = p

//line xss_400.go:193
		switch data[p] {
		case 46:
			goto tr26
		case 91:
			goto tr27
		case 99:
			goto tr28
		case 114:
			goto tr29
		}
		goto tr25
tr8:
//line NONE:1
te = p+1

//line xss_400.rl:36
act = 2;
	goto st22
tr26:
//line NONE:1
te = p+1

//line xss_400.rl:56
act = 4;
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line xss_400.go:224
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 115:
			goto st5
		}
		goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 115:
			goto st5
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 112:
			goto st8
		case 115:
			goto st5
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 97:
			goto st3
		case 109:
			goto st2
		case 115:
			goto st5
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 112:
			goto st4
		case 115:
			goto st5
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 46 {
			goto tr8
		}
		goto st4
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 111:
			goto st6
		case 115:
			goto st5
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st5
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 115:
			goto st5
		case 116:
			goto st4
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 112:
			goto st9
		case 115:
			goto st5
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 97:
			goto st1
		case 108:
			goto st10
		case 109:
			goto st2
		case 115:
			goto st5
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 97:
			goto st1
		case 109:
			goto st2
		case 115:
			goto st5
		case 121:
			goto st4
		}
		goto st0
tr27:
//line NONE:1
te = p+1

//line xss_400.rl:56
act = 4;
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line xss_400.go:411
		if data[p] == 93 {
			goto tr14
		}
		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 93 {
			goto tr14
		}
		goto st11
tr14:
//line NONE:1
te = p+1

//line xss_400.rl:31
act = 1;
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line xss_400.go:437
		if data[p] == 46 {
			goto st11
		}
		goto tr14
tr28:
//line NONE:1
te = p+1

//line xss_400.rl:56
act = 4;
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line xss_400.go:454
		if data[p] == 97 {
			goto st12
		}
		goto tr30
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 108 {
			goto st13
		}
		goto tr15
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 108 {
			goto st14
		}
		goto tr15
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 96 {
			goto st15
		}
		goto st14
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 96 {
			goto tr19
		}
		goto st15
tr19:
//line NONE:1
te = p+1

//line xss_400.rl:46
act = 3;
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line xss_400.go:507
		if data[p] == 96 {
			goto tr19
		}
		goto st15
tr29:
//line NONE:1
te = p+1

	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line xss_400.go:522
		if data[p] == 101 {
			goto st16
		}
		goto tr30
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 102 {
			goto st17
		}
		goto tr15
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 108 {
			goto st18
		}
		goto tr15
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 101 {
			goto st19
		}
		goto tr15
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 99 {
			goto st20
		}
		goto tr15
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 116 {
			goto st28
		}
		goto tr15
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 46 {
			goto tr31
		}
		goto st28
	st_out:
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 22:
			goto tr0
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		case 4:
			goto tr0
		case 5:
			goto tr0
		case 6:
			goto tr0
		case 7:
			goto tr0
		case 8:
			goto tr0
		case 9:
			goto tr0
		case 10:
			goto tr0
		case 23:
			goto tr30
		case 11:
			goto tr0
		case 24:
			goto tr31
		case 25:
			goto tr30
		case 12:
			goto tr15
		case 13:
			goto tr15
		case 14:
			goto tr15
		case 15:
			goto tr0
		case 26:
			goto tr33
		case 27:
			goto tr30
		case 16:
			goto tr15
		case 17:
			goto tr15
		case 18:
			goto tr15
		case 19:
			goto tr15
		case 20:
			goto tr15
		case 28:
			goto tr31
		}
	}

	}

//line xss_400.rl:61


    return false
}
