
//line php_injection.rl:1
package php_attack

func matchPhpInjection(data []byte) bool {

//line php_injection.rl:5

//line php_injection.go:10
const phpInjection_start int = 44
const phpInjection_first_final int = 44
const phpInjection_error int = -1

const phpInjection_en_main int = 44


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
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 0:
		goto st_case_0
	case 50:
		goto st_case_50
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 51:
		goto st_case_51
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
	case 52:
		goto st_case_52
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
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	}
	goto st_out
tr0:
//line php_injection.rl:24
p = (te) - 1
{
                return true
            }
	goto st44
tr2:
//line php_injection.rl:36
p = (te) - 1

	goto st44
tr6:
//line php_injection.rl:24
te = p+1
{
                return true
            }
	goto st44
tr16:
//line php_injection.rl:32
te = p+1
{
                return true
            }
	goto st44
tr27:
//line php_injection.rl:28
te = p+1
{
                return true
            }
	goto st44
tr46:
//line php_injection.rl:36
te = p+1

	goto st44
tr51:
//line php_injection.rl:36
te = p
p--

	goto st44
tr53:
//line php_injection.rl:24
te = p
p--
{
                return true
            }
	goto st44
	st44:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line NONE:1
ts = p

//line php_injection.go:211
		switch data[p] {
		case 60:
			goto st45
		case 91:
			goto tr48
		case 99:
			goto tr49
		case 111:
			goto tr49
		case 112:
			goto tr50
		}
		goto tr46
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 63 {
			goto st46
		}
		goto tr51
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 120 {
			goto st48
		}
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		goto st47
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 109 {
			goto tr55
		}
		goto st47
tr55:
//line NONE:1
te = p+1

	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line php_injection.go:268
		if data[p] == 108 {
			goto st0
		}
		goto st47
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		goto st47
tr48:
//line NONE:1
te = p+1

	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line php_injection.go:289
		switch data[p] {
		case 47:
			goto st1
		case 92:
			goto st1
		}
		goto tr51
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
tr49:
//line NONE:1
te = p+1

	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line php_injection.go:343
		if data[p] == 58 {
			goto st5
		}
		goto tr51
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= data[p] && data[p] <= 57 {
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
		if 48 <= data[p] && data[p] <= 57 {
			goto st6
		}
		goto tr2
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 34 {
			goto st8
		}
		goto tr2
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 34 {
			goto tr2
		}
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 34 {
			goto st10
		}
		goto st9
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 58 {
			goto st11
		}
		goto tr2
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= data[p] && data[p] <= 57 {
			goto st12
		}
		goto tr2
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 58 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st12
		}
		goto tr2
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 123 {
			goto st14
		}
		goto tr2
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 125 {
			goto tr16
		}
		goto st14
tr50:
//line NONE:1
te = p+1

	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line php_injection.go:454
		if data[p] == 104 {
			goto st15
		}
		goto tr51
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 112 {
			goto st16
		}
		goto tr2
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 58 {
			goto st17
		}
		goto tr2
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 47 {
			goto st18
		}
		goto tr2
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 47 {
			goto st19
		}
		goto tr2
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 102:
			goto st20
		case 105:
			goto st25
		case 109:
			goto st29
		case 111:
			goto st34
		case 115:
			goto st36
		case 116:
			goto st41
		}
		goto tr2
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 100:
			goto tr27
		case 105:
			goto st21
		}
		goto tr2
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 108 {
			goto st22
		}
		goto tr2
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 116 {
			goto st23
		}
		goto tr2
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 101 {
			goto st24
		}
		goto tr2
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 114 {
			goto tr27
		}
		goto tr2
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 110 {
			goto st26
		}
		goto tr2
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 112 {
			goto st27
		}
		goto tr2
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 117 {
			goto st28
		}
		goto tr2
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 116 {
			goto tr27
		}
		goto tr2
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 101 {
			goto st30
		}
		goto tr2
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 109 {
			goto st31
		}
		goto tr2
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 111 {
			goto st32
		}
		goto tr2
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 114 {
			goto st33
		}
		goto tr2
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 121 {
			goto tr27
		}
		goto tr2
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 117 {
			goto st35
		}
		goto tr2
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 116 {
			goto st26
		}
		goto tr2
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 116 {
			goto st37
		}
		goto tr2
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 100 {
			goto st38
		}
		goto tr2
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 101:
			goto st39
		case 105:
			goto st40
		case 111:
			goto st27
		}
		goto tr2
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 114 {
			goto st24
		}
		goto tr2
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 110 {
			goto tr27
		}
		goto tr2
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 101 {
			goto st42
		}
		goto tr2
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 109 {
			goto st43
		}
		goto tr2
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 112 {
			goto tr27
		}
		goto tr2
	st_out:
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
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
	_test_eof52: cs = 52; goto _test_eof
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
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 45:
			goto tr51
		case 46:
			goto tr53
		case 47:
			goto tr53
		case 48:
			goto tr53
		case 49:
			goto tr53
		case 0:
			goto tr0
		case 50:
			goto tr51
		case 1:
			goto tr2
		case 2:
			goto tr2
		case 3:
			goto tr2
		case 4:
			goto tr2
		case 51:
			goto tr51
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
		case 52:
			goto tr51
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
		case 34:
			goto tr2
		case 35:
			goto tr2
		case 36:
			goto tr2
		case 37:
			goto tr2
		case 38:
			goto tr2
		case 39:
			goto tr2
		case 40:
			goto tr2
		case 41:
			goto tr2
		case 42:
			goto tr2
		case 43:
			goto tr2
		}
	}

	}

//line php_injection.rl:41

        return false
}