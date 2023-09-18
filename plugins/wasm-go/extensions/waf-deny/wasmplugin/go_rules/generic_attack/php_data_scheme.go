
//line php_data_scheme.rl:1
package generic_attack

func matchPhpDataScheme(data []byte) bool {

//line php_data_scheme.rl:5

//line php_data_scheme.go:10
const phpDataScheme_start int = 1
const phpDataScheme_first_final int = 11
const phpDataScheme_error int = 0

const phpDataScheme_en_main int = 1


//line php_data_scheme.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line php_data_scheme.go:26
	{
	cs = phpDataScheme_start
	}

//line php_data_scheme.go:31
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
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
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 10:
		goto st_case_10
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	}
	goto st_out
	st_case_1:
		if data[p] == 100 {
			goto st2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 97 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 116 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 97 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 58 {
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 42:
			goto st10
		case 44:
			goto st0
		case 47:
			goto st0
		case 123:
			goto st0
		case 125:
			goto st0
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto st0
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto st0
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 44:
			goto st0
		case 47:
			goto st8
		case 123:
			goto st0
		case 125:
			goto st0
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto st0
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto st0
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto st7
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 44:
			goto st0
		case 47:
			goto st0
		case 123:
			goto st0
		case 125:
			goto st0
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto st0
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto st0
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 44:
			goto tr11
		case 47:
			goto tr11
		case 123:
			goto tr11
		case 125:
			goto tr11
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto tr11
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto tr11
				}
			case data[p] >= 58:
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr10
tr10:
//line php_data_scheme.rl:13

            return true
        
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line php_data_scheme.go:259
		switch data[p] {
		case 44:
			goto tr11
		case 47:
			goto tr11
		case 123:
			goto tr11
		case 125:
			goto tr11
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto tr11
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto tr11
				}
			case data[p] >= 58:
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr10
tr11:
//line php_data_scheme.rl:13

            return true
        
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line php_data_scheme.go:299
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 44:
			goto tr11
		case 47:
			goto tr13
		case 123:
			goto tr11
		case 125:
			goto tr11
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto tr11
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto tr11
				}
			case data[p] >= 58:
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr12
tr12:
//line php_data_scheme.rl:13

            return true
        
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line php_data_scheme.go:345
		switch data[p] {
		case 44:
			goto st0
		case 47:
			goto st8
		case 123:
			goto st0
		case 125:
			goto st0
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto st0
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto st0
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto st7
tr13:
//line php_data_scheme.rl:13

            return true
        
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line php_data_scheme.go:385
		switch data[p] {
		case 44:
			goto st0
		case 47:
			goto st0
		case 123:
			goto st0
		case 125:
			goto st0
		}
		switch {
		case data[p] < 40:
			if 33 <= data[p] && data[p] <= 34 {
				goto st0
			}
		case data[p] > 41:
			switch {
			case data[p] > 63:
				if 91 <= data[p] && data[p] <= 93 {
					goto st0
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto st9
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 9, 10, 11:
//line php_data_scheme.rl:13

            return true
        
//line php_data_scheme.go:437
		}
	}

	_out: {}
	}

//line php_data_scheme.rl:25

        return false
}

