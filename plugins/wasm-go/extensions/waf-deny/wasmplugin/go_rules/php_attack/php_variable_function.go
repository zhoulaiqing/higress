
//line php_variable_function.rl:1
package php_attack

func matchPhpVariableFunction(data []byte) bool {

//line php_variable_function.rl:5

//line php_variable_function.go:10
const phpVariableFunction_start int = 55
const phpVariableFunction_first_final int = 55
const phpVariableFunction_error int = -1

const phpVariableFunction_en_main int = 55


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
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 57:
		goto st_case_57
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
	case 58:
		goto st_case_58
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
	case 59:
		goto st_case_59
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 60:
		goto st_case_60
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 61:
		goto st_case_61
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
	case 62:
		goto st_case_62
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
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
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 63:
		goto st_case_63
	case 53:
		goto st_case_53
	case 64:
		goto st_case_64
	case 54:
		goto st_case_54
	}
	goto st_out
tr0:
//line php_variable_function.rl:59
p = (te) - 1

	goto st55
tr5:
//line php_variable_function.rl:55
te = p+1
{
                return true
            }
	goto st55
tr16:
//line NONE:1
	switch act {
	case 2:
	{p = (te) - 1

                return true
            }
	default:
	{p = (te) - 1
}
	}
	
	goto st55
tr24:
//line php_variable_function.rl:51
te = p+1
{
                return true
            }
	goto st55
tr63:
//line php_variable_function.rl:59
te = p+1

	goto st55
tr69:
//line php_variable_function.rl:59
te = p
p--

	goto st55
tr70:
//line php_variable_function.rl:55
te = p
p--
{
                return true
            }
	goto st55
	st55:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line NONE:1
ts = p

//line php_variable_function.go:234
		switch data[p] {
		case 34:
			goto tr64
		case 36:
			goto tr65
		case 39:
			goto tr64
		case 40:
			goto tr66
		case 91:
			goto tr67
		case 123:
			goto tr68
		}
		goto tr63
tr64:
//line NONE:1
te = p+1

	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line php_variable_function.go:260
		switch data[p] {
		case 92:
			goto st0
		case 95:
			goto st0
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr69
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 34:
			goto st1
		case 39:
			goto st1
		case 92:
			goto st0
		case 95:
			goto st0
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st0
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 40 {
			goto st2
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 41 {
			goto tr0
		}
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 41 {
			goto tr5
		}
		goto st3
tr65:
//line NONE:1
te = p+1

//line php_variable_function.rl:59
act = 3;
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line php_variable_function.go:347
		switch data[p] {
		case 32:
			goto st5
		case 36:
			goto st38
		case 44:
			goto tr69
		case 55:
			goto st39
		case 59:
			goto tr69
		case 92:
			goto tr69
		case 95:
			goto st39
		case 123:
			goto st6
		}
		switch {
		case data[p] < 46:
			switch {
			case data[p] > 13:
				if 40 <= data[p] && data[p] <= 41 {
					goto tr69
				}
			case data[p] >= 9:
				goto st5
			}
		case data[p] > 47:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto tr69
		}
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 40:
			goto st2
		case 41:
			goto tr0
		case 44:
			goto tr0
		case 59:
			goto tr0
		case 92:
			goto tr0
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto tr0
		}
		goto st4
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 32:
			goto st5
		case 40:
			goto st2
		case 41:
			goto tr0
		case 44:
			goto tr0
		case 59:
			goto tr0
		case 92:
			goto tr0
		case 123:
			goto st6
		}
		switch {
		case data[p] > 13:
			if 46 <= data[p] && data[p] <= 47 {
				goto tr0
			}
		case data[p] >= 9:
			goto st5
		}
		goto st4
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 40:
			goto st8
		case 41:
			goto st10
		case 44:
			goto st10
		case 59:
			goto st10
		case 92:
			goto st10
		case 125:
			goto st4
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto st10
		}
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 40:
			goto st8
		case 41:
			goto st10
		case 44:
			goto st10
		case 59:
			goto st10
		case 92:
			goto st10
		case 125:
			goto st31
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto st10
		}
		goto st7
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 41:
			goto st10
		case 125:
			goto st21
		}
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 41:
			goto tr15
		case 125:
			goto st21
		}
		goto st9
tr15:
//line NONE:1
te = p+1

//line php_variable_function.rl:55
act = 2;
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line php_variable_function.go:522
		if data[p] == 125 {
			goto st11
		}
		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 125 {
			goto st11
		}
		goto st10
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 32:
			goto st11
		case 35:
			goto st12
		case 40:
			goto st13
		case 47:
			goto st15
		case 91:
			goto st18
		case 123:
			goto st20
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st11
		}
		goto tr16
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 10 {
			goto st11
		}
		goto st12
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 41 {
			goto tr16
		}
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 41 {
			goto tr24
		}
		goto st14
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 42:
			goto st16
		case 47:
			goto st12
		}
		goto tr16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 42 {
			goto st17
		}
		goto st16
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 42:
			goto st17
		case 47:
			goto st11
		}
		goto st16
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 93 {
			goto tr16
		}
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 93 {
			goto st11
		}
		goto st19
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 125 {
			goto tr16
		}
		goto st10
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 32:
			goto st21
		case 35:
			goto st22
		case 40:
			goto st23
		case 41:
			goto tr5
		case 47:
			goto st25
		case 91:
			goto st28
		case 123:
			goto st30
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st21
		}
		goto st3
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 10:
			goto st21
		case 41:
			goto tr33
		}
		goto st22
tr33:
//line NONE:1
te = p+1

//line php_variable_function.rl:55
act = 2;
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line php_variable_function.go:695
		if data[p] == 10 {
			goto st11
		}
		goto st12
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 41 {
			goto tr5
		}
		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 41 {
			goto tr24
		}
		goto st24
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 41:
			goto tr5
		case 42:
			goto st26
		case 47:
			goto st22
		}
		goto st3
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 41:
			goto tr36
		case 42:
			goto st27
		}
		goto st26
tr36:
//line NONE:1
te = p+1

//line php_variable_function.rl:55
act = 2;
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line php_variable_function.go:756
		if data[p] == 42 {
			goto st17
		}
		goto st16
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 41:
			goto tr36
		case 42:
			goto st27
		case 47:
			goto st21
		}
		goto st26
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 41:
			goto tr39
		case 93:
			goto st3
		}
		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 41:
			goto tr39
		case 93:
			goto st21
		}
		goto st29
tr39:
//line NONE:1
te = p+1

//line php_variable_function.rl:55
act = 2;
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line php_variable_function.go:811
		if data[p] == 93 {
			goto st11
		}
		goto st19
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 41:
			goto tr15
		case 125:
			goto st3
		}
		goto st9
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 32:
			goto st31
		case 35:
			goto st32
		case 40:
			goto st34
		case 41:
			goto tr0
		case 44:
			goto tr0
		case 46:
			goto tr0
		case 47:
			goto st15
		case 59:
			goto tr0
		case 91:
			goto st35
		case 92:
			goto tr0
		case 123:
			goto st6
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st31
		}
		goto st4
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 10:
			goto st31
		case 40:
			goto st33
		case 41:
			goto st12
		case 44:
			goto st12
		case 59:
			goto st12
		case 92:
			goto st12
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto st12
		}
		goto st32
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 10:
			goto st21
		case 41:
			goto st12
		}
		goto st22
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 41 {
			goto tr0
		}
		goto st24
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 40:
			goto st37
		case 41:
			goto st19
		case 44:
			goto st19
		case 59:
			goto st19
		case 92:
			goto st19
		case 93:
			goto st4
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto st19
		}
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 40:
			goto st37
		case 41:
			goto st19
		case 44:
			goto st19
		case 59:
			goto st19
		case 92:
			goto st19
		case 93:
			goto st31
		}
		if 46 <= data[p] && data[p] <= 47 {
			goto st19
		}
		goto st36
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 41:
			goto st19
		case 93:
			goto st21
		}
		goto st29
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 32:
			goto st5
		case 36:
			goto st38
		case 40:
			goto st2
		case 41:
			goto tr0
		case 44:
			goto tr0
		case 55:
			goto st39
		case 59:
			goto tr0
		case 92:
			goto tr0
		case 95:
			goto st39
		case 123:
			goto st6
		}
		switch {
		case data[p] < 46:
			if 9 <= data[p] && data[p] <= 13 {
				goto st5
			}
		case data[p] > 47:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto tr0
		}
		goto st4
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 32:
			goto st31
		case 35:
			goto st32
		case 40:
			goto st34
		case 41:
			goto tr0
		case 44:
			goto tr0
		case 46:
			goto tr0
		case 47:
			goto st15
		case 59:
			goto tr0
		case 91:
			goto st35
		case 92:
			goto tr0
		case 95:
			goto st39
		case 123:
			goto st6
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st31
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st4
tr66:
//line NONE:1
te = p+1

	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line php_variable_function.go:1067
		switch data[p] {
		case 41:
			goto tr69
		case 115:
			goto st44
		}
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 41:
			goto st41
		case 115:
			goto st44
		}
		goto st40
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 34:
			goto st42
		case 39:
			goto st42
		case 40:
			goto st2
		}
		goto tr0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 95 {
			goto st43
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st43
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st43
			}
		default:
			goto st43
		}
		goto tr0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 34:
			goto st1
		case 39:
			goto st1
		case 95:
			goto st43
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st43
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st43
			}
		default:
			goto st43
		}
		goto tr0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 41:
			goto st41
		case 115:
			goto st44
		case 116:
			goto st45
		}
		goto st40
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 41:
			goto st41
		case 114:
			goto st46
		case 115:
			goto st44
		}
		goto st40
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 41:
			goto st41
		case 105:
			goto st47
		case 115:
			goto st44
		}
		goto st40
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 41:
			goto st41
		case 110:
			goto st48
		case 115:
			goto st44
		}
		goto st40
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 41:
			goto st41
		case 103:
			goto st49
		case 115:
			goto st44
		}
		goto st40
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		if data[p] == 41 {
			goto st50
		}
		goto st49
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 32:
			goto st51
		case 34:
			goto st52
		case 39:
			goto st52
		case 40:
			goto st2
		case 46:
			goto st51
		case 93:
			goto st51
		case 125:
			goto st51
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st51
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 123 {
					goto st51
				}
			case data[p] >= 65:
				goto st51
			}
		default:
			goto st51
		}
		goto tr0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 32:
			goto st51
		case 34:
			goto st51
		case 39:
			goto st51
		case 40:
			goto st3
		case 46:
			goto st51
		case 93:
			goto st51
		case 125:
			goto st51
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st51
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 123 {
					goto st51
				}
			case data[p] >= 65:
				goto st51
			}
		default:
			goto st51
		}
		goto tr0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 32:
			goto st51
		case 34:
			goto st51
		case 39:
			goto st51
		case 40:
			goto st3
		case 46:
			goto st51
		case 91:
			goto st51
		case 93:
			goto st51
		case 95:
			goto st43
		case 123:
			goto st51
		case 125:
			goto st51
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto st51
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st52
				}
			case data[p] >= 65:
				goto st52
			}
		default:
			goto st52
		}
		goto tr0
tr67:
//line NONE:1
te = p+1

	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line php_variable_function.go:1360
		if 48 <= data[p] && data[p] <= 57 {
			goto st53
		}
		goto tr69
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 93 {
			goto st1
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st53
		}
		goto tr0
tr68:
//line NONE:1
te = p+1

	goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line php_variable_function.go:1387
		if 48 <= data[p] && data[p] <= 57 {
			goto st54
		}
		goto tr69
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 125 {
			goto st1
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st54
		}
		goto tr0
	st_out:
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
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
	_test_eof59: cs = 59; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
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
	_test_eof62: cs = 62; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 56:
			goto tr69
		case 0:
			goto tr0
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		case 57:
			goto tr69
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
		case 58:
			goto tr70
		case 10:
			goto tr16
		case 11:
			goto tr16
		case 12:
			goto tr16
		case 13:
			goto tr16
		case 14:
			goto tr16
		case 15:
			goto tr16
		case 16:
			goto tr16
		case 17:
			goto tr16
		case 18:
			goto tr16
		case 19:
			goto tr16
		case 20:
			goto tr16
		case 21:
			goto tr0
		case 22:
			goto tr0
		case 59:
			goto tr70
		case 23:
			goto tr0
		case 24:
			goto tr0
		case 25:
			goto tr0
		case 26:
			goto tr0
		case 60:
			goto tr70
		case 27:
			goto tr0
		case 28:
			goto tr0
		case 29:
			goto tr0
		case 61:
			goto tr70
		case 30:
			goto tr0
		case 31:
			goto tr0
		case 32:
			goto tr0
		case 33:
			goto tr0
		case 34:
			goto tr0
		case 35:
			goto tr0
		case 36:
			goto tr0
		case 37:
			goto tr0
		case 38:
			goto tr0
		case 39:
			goto tr0
		case 62:
			goto tr69
		case 40:
			goto tr0
		case 41:
			goto tr0
		case 42:
			goto tr0
		case 43:
			goto tr0
		case 44:
			goto tr0
		case 45:
			goto tr0
		case 46:
			goto tr0
		case 47:
			goto tr0
		case 48:
			goto tr0
		case 49:
			goto tr0
		case 50:
			goto tr0
		case 51:
			goto tr0
		case 52:
			goto tr0
		case 63:
			goto tr69
		case 53:
			goto tr0
		case 64:
			goto tr69
		case 54:
			goto tr0
		}
	}

	}

//line php_variable_function.rl:64

        return false
}