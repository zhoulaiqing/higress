
//line php_variable_function.rl:1
package php_attack

func matchPhpVariableFunction(data []byte) bool {

//line php_variable_function.rl:5

//line php_variable_function.go:10
const phpVariableFunction_start int = 14
const phpVariableFunction_first_final int = 14
const phpVariableFunction_error int = -1

const phpVariableFunction_en_main int = 14


//line php_variable_function.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line php_variable_function.go:25
	{
	cs = phpVariableFunction_start
	ts = 0
	te = 0
	act = 0
	}

//line php_variable_function.go:33
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
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
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	}
	goto st_out
tr0:
//line php_variable_function.rl:36
p = (te) - 1

	goto st14
tr10:
//line php_variable_function.rl:32
te = p+1
{
                return true
            }
	goto st14
tr16:
//line php_variable_function.rl:36
te = p+1

	goto st14
tr18:
//line php_variable_function.rl:36
te = p
p--

	goto st14
	st14:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line NONE:1
ts = p

//line php_variable_function.go:107
		if data[p] == 36 {
			goto tr17
		}
		goto tr16
tr17:
//line NONE:1
te = p+1

	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line php_variable_function.go:122
		switch data[p] {
		case 32:
			goto st0
		case 36:
			goto st12
		case 55:
			goto st13
		case 95:
			goto st13
		case 123:
			goto st1
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr18
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 32:
			goto st0
		case 123:
			goto st1
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st0
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 125 {
			goto tr0
		}
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 125 {
			goto st3
		}
		goto st2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 32:
			goto st3
		case 35:
			goto st4
		case 40:
			goto st5
		case 47:
			goto st7
		case 91:
			goto st10
		case 123:
			goto st1
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st3
		}
		goto tr0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 10 {
			goto st3
		}
		goto st4
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 41 {
			goto tr0
		}
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 41 {
			goto tr10
		}
		goto st6
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 42:
			goto st8
		case 47:
			goto st4
		}
		goto tr0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 42 {
			goto st9
		}
		goto st8
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 42:
			goto st9
		case 47:
			goto st3
		}
		goto st8
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 93 {
			goto tr0
		}
		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 93 {
			goto st3
		}
		goto st11
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 32:
			goto st0
		case 36:
			goto st12
		case 55:
			goto st13
		case 95:
			goto st13
		case 123:
			goto st1
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto tr0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 32:
			goto st3
		case 35:
			goto st4
		case 40:
			goto st5
		case 47:
			goto st7
		case 91:
			goto st10
		case 95:
			goto st13
		case 123:
			goto st1
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st3
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st13
				}
			case data[p] >= 65:
				goto st13
			}
		default:
			goto st13
		}
		goto tr0
	st_out:
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
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
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 15:
			goto tr18
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
		case 11:
			goto tr0
		case 12:
			goto tr0
		case 13:
			goto tr0
		}
	}

	}

//line php_variable_function.rl:41

        return false
}