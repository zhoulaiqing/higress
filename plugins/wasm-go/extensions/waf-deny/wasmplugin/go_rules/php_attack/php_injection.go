
//line php_injection.rl:1
package php_attack

func matchPhpInjection(data []byte) bool {

//line php_injection.rl:5

//line php_injection.go:10
const phpInjection_start int = 34
const phpInjection_first_final int = 34
const phpInjection_error int = -1

const phpInjection_en_main int = 34


//line php_injection.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
    _, _, _ = ts, te, act

    
//line php_injection.go:26
	{
	cs = phpInjection_start
	ts = 0
	te = 0
	act = 0
	}

//line php_injection.go:34
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 0:
		goto st_case_0
	case 40:
		goto st_case_40
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 41:
		goto st_case_41
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
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
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
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	}
	goto st_out
tr0:
//line php_injection.rl:22
p = (te) - 1
{
                return true
            }
	goto st34
tr2:
//line php_injection.rl:30
p = (te) - 1

	goto st34
tr6:
//line php_injection.rl:22
te = p+1
{
                return true
            }
	goto st34
tr17:
//line php_injection.rl:26
te = p+1
{
                return true
            }
	goto st34
tr36:
//line php_injection.rl:30
te = p+1

	goto st34
tr40:
//line php_injection.rl:30
te = p
p--

	goto st34
tr42:
//line php_injection.rl:22
te = p
p--
{
                return true
            }
	goto st34
	st34:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line NONE:1
ts = p

//line php_injection.go:182
		switch data[p] {
		case 60:
			goto st35
		case 91:
			goto tr38
		case 112:
			goto tr39
		}
		goto tr36
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 63 {
			goto st36
		}
		goto tr40
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 120 {
			goto st38
		}
		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		goto st37
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 109 {
			goto tr44
		}
		goto st37
tr44:
//line NONE:1
te = p+1

	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line php_injection.go:235
		if data[p] == 108 {
			goto st0
		}
		goto st37
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		goto st37
tr38:
//line NONE:1
te = p+1

	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line php_injection.go:256
		switch data[p] {
		case 47:
			goto st1
		case 92:
			goto st1
		}
		goto tr40
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 112 {
			goto st2
		}
		goto tr2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 104 {
			goto st3
		}
		goto tr2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 112 {
			goto st4
		}
		goto tr2
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 93 {
			goto tr6
		}
		goto tr2
tr39:
//line NONE:1
te = p+1

	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line php_injection.go:310
		if data[p] == 104 {
			goto st5
		}
		goto tr40
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 112 {
			goto st6
		}
		goto tr2
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 58 {
			goto st7
		}
		goto tr2
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 47 {
			goto st8
		}
		goto tr2
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 47 {
			goto st9
		}
		goto tr2
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 102:
			goto st10
		case 105:
			goto st15
		case 109:
			goto st19
		case 111:
			goto st24
		case 115:
			goto st26
		case 116:
			goto st31
		}
		goto tr2
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 100:
			goto tr17
		case 105:
			goto st11
		}
		goto tr2
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 108 {
			goto st12
		}
		goto tr2
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 116 {
			goto st13
		}
		goto tr2
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 101 {
			goto st14
		}
		goto tr2
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 114 {
			goto tr17
		}
		goto tr2
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 110 {
			goto st16
		}
		goto tr2
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 112 {
			goto st17
		}
		goto tr2
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 117 {
			goto st18
		}
		goto tr2
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 116 {
			goto tr17
		}
		goto tr2
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 101 {
			goto st20
		}
		goto tr2
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 109 {
			goto st21
		}
		goto tr2
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 111 {
			goto st22
		}
		goto tr2
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 114 {
			goto st23
		}
		goto tr2
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 121 {
			goto tr17
		}
		goto tr2
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 117 {
			goto st25
		}
		goto tr2
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 116 {
			goto st16
		}
		goto tr2
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 116 {
			goto st27
		}
		goto tr2
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 100 {
			goto st28
		}
		goto tr2
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 101:
			goto st29
		case 105:
			goto st30
		case 111:
			goto st17
		}
		goto tr2
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 114 {
			goto st14
		}
		goto tr2
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 110 {
			goto tr17
		}
		goto tr2
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 101 {
			goto st32
		}
		goto tr2
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 109 {
			goto st33
		}
		goto tr2
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 112 {
			goto tr17
		}
		goto tr2
	st_out:
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 35:
			goto tr40
		case 36:
			goto tr42
		case 37:
			goto tr42
		case 38:
			goto tr42
		case 39:
			goto tr42
		case 0:
			goto tr0
		case 40:
			goto tr40
		case 1:
			goto tr2
		case 2:
			goto tr2
		case 3:
			goto tr2
		case 4:
			goto tr2
		case 41:
			goto tr40
		case 5:
			goto tr2
		case 6:
			goto tr2
		case 7:
			goto tr2
		case 8:
			goto tr2
		case 9:
			goto tr2
		case 10:
			goto tr2
		case 11:
			goto tr2
		case 12:
			goto tr2
		case 13:
			goto tr2
		case 14:
			goto tr2
		case 15:
			goto tr2
		case 16:
			goto tr2
		case 17:
			goto tr2
		case 18:
			goto tr2
		case 19:
			goto tr2
		case 20:
			goto tr2
		case 21:
			goto tr2
		case 22:
			goto tr2
		case 23:
			goto tr2
		case 24:
			goto tr2
		case 25:
			goto tr2
		case 26:
			goto tr2
		case 27:
			goto tr2
		case 28:
			goto tr2
		case 29:
			goto tr2
		case 30:
			goto tr2
		case 31:
			goto tr2
		case 32:
			goto tr2
		case 33:
			goto tr2
		}
	}

	}

//line php_injection.rl:35

        return false
}