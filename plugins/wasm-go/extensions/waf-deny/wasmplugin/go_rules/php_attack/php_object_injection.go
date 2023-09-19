
//line php_object_injection.rl:1
package php_attack

func matchPhpObjectInjection(data []byte) bool {

//line php_object_injection.rl:5

//line php_object_injection.go:10
const phpObjectInjection_start int = 10
const phpObjectInjection_first_final int = 10
const phpObjectInjection_error int = -1

const phpObjectInjection_en_main int = 10


//line php_object_injection.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line php_object_injection.go:26
	{
	cs = phpObjectInjection_start
	ts = 0
	te = 0
	act = 0
	}

//line php_object_injection.go:34
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
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
	}
	goto st_out
tr0:
//line php_object_injection.rl:21
p = (te) - 1

	goto st10
tr10:
//line php_object_injection.rl:17
te = p+1
{
                return true
            }
	goto st10
tr11:
//line php_object_injection.rl:21
te = p+1

	goto st10
tr13:
//line php_object_injection.rl:21
te = p
p--

	goto st10
	st10:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line NONE:1
ts = p

//line php_object_injection.go:100
		switch data[p] {
		case 99:
			goto tr12
		case 111:
			goto tr12
		}
		goto tr11
tr12:
//line NONE:1
te = p+1

	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line php_object_injection.go:118
		if data[p] == 58 {
			goto st0
		}
		goto tr13
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if 48 <= data[p] && data[p] <= 57 {
			goto st1
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 58 {
			goto st2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st1
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 34 {
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 34 {
			goto tr0
		}
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 34 {
			goto st5
		}
		goto st4
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 58 {
			goto st6
		}
		goto tr0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if 48 <= data[p] && data[p] <= 57 {
			goto st7
		}
		goto tr0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 58 {
			goto st8
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st7
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 123 {
			goto st9
		}
		goto tr0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 125 {
			goto tr10
		}
		goto st9
	st_out:
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
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

	_test_eof: {}
	if p == eof {
		switch cs {
		case 11:
			goto tr13
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
		}
	}

	}

//line php_object_injection.rl:26

        return false
}