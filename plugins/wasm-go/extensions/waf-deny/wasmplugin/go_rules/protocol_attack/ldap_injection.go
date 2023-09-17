
//line ldap_injection.rl:1
package protocol_attack

func matchLdapInjection(data []byte) bool {

//line ldap_injection.rl:5

//line ldap_injection.go:10
const ldapInjection_start int = 5
const ldapInjection_first_final int = 5
const ldapInjection_error int = -1

const ldapInjection_en_main int = 5


//line ldap_injection.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof


    var ts, te, act int
    _, _, _ = ts, te, act

    startWithRightPart := false

    
//line ldap_injection.go:29
	{
	cs = ldapInjection_start
	ts = 0
	te = 0
	act = 0
	}

//line ldap_injection.go:37
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 0:
		goto st_case_0
	case 7:
		goto st_case_7
	case 1:
		goto st_case_1
	case 8:
		goto st_case_8
	case 2:
		goto st_case_2
	case 9:
		goto st_case_9
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
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
	}
	goto st_out
tr0:
//line ldap_injection.rl:46
p = (te) - 1

	goto st5
tr3:
//line ldap_injection.rl:40
p = (te) - 1
{
                if !startWithRightPart {
                    startWithRightPart = true
                }
            }
	goto st5
tr5:
//line ldap_injection.rl:40
te = p+1
{
                if !startWithRightPart {
                    startWithRightPart = true
                }
            }
	goto st5
tr8:
//line ldap_injection.rl:32
p = (te) - 1
{
                if startWithRightPart {
                    return true
                } else {
                    return false
                }
            }
	goto st5
tr15:
//line ldap_injection.rl:46
te = p+1

	goto st5
tr16:
//line ldap_injection.rl:46
te = p
p--

	goto st5
tr17:
//line ldap_injection.rl:40
te = p
p--
{
                if !startWithRightPart {
                    startWithRightPart = true
                }
            }
	goto st5
tr18:
//line ldap_injection.rl:32
te = p
p--
{
                if startWithRightPart {
                    return true
                } else {
                    return false
                }
            }
	goto st5
	st5:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line NONE:1
ts = p

//line ldap_injection.go:153
		switch data[p] {
		case 33:
			goto tr13
		case 38:
			goto st14
		case 40:
			goto tr7
		case 41:
			goto tr2
		case 58:
			goto tr15
		case 60:
			goto tr15
		case 62:
			goto tr15
		case 124:
			goto st14
		case 126:
			goto tr15
		}
		goto tr12
tr12:
//line NONE:1
te = p+1

	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line ldap_injection.go:185
		switch data[p] {
		case 33:
			goto tr16
		case 38:
			goto tr16
		case 40:
			goto tr16
		case 41:
			goto tr2
		case 58:
			goto tr16
		case 60:
			goto tr16
		case 62:
			goto tr16
		case 124:
			goto tr16
		case 126:
			goto tr16
		}
		goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 33:
			goto tr0
		case 38:
			goto tr0
		case 40:
			goto tr0
		case 41:
			goto tr2
		case 58:
			goto tr0
		case 60:
			goto tr0
		case 62:
			goto tr0
		case 124:
			goto tr0
		case 126:
			goto tr0
		}
		goto st0
tr2:
//line NONE:1
te = p+1

	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line ldap_injection.go:243
		switch data[p] {
		case 32:
			goto st1
		case 41:
			goto tr5
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto tr17
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st1
		case 41:
			goto tr5
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto tr3
tr13:
//line NONE:1
te = p+1

	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line ldap_injection.go:279
		switch data[p] {
		case 32:
			goto st2
		case 40:
			goto tr7
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st2
		}
		goto tr16
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 32:
			goto st2
		case 40:
			goto tr7
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st2
		}
		goto tr0
tr7:
//line NONE:1
te = p+1

	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ldap_injection.go:315
		switch data[p] {
		case 33:
			goto tr19
		case 38:
			goto tr20
		case 44:
			goto tr18
		case 58:
			goto tr18
		case 124:
			goto tr20
		case 126:
			goto tr18
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr18
			}
		case data[p] >= 40:
			goto tr18
		}
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 33:
			goto tr8
		case 38:
			goto tr8
		case 44:
			goto tr8
		case 58:
			goto tr8
		case 61:
			goto st10
		case 124:
			goto tr8
		case 126:
			goto st4
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto st4
			}
		case data[p] >= 40:
			goto tr8
		}
		goto st3
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 61 {
			goto st10
		}
		goto tr8
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 33:
			goto tr18
		case 38:
			goto tr18
		case 58:
			goto tr18
		case 60:
			goto tr18
		case 62:
			goto tr18
		case 124:
			goto tr18
		case 126:
			goto tr18
		}
		if 40 <= data[p] && data[p] <= 41 {
			goto tr18
		}
		goto st10
tr19:
//line NONE:1
te = p+1

	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line ldap_injection.go:413
		switch data[p] {
		case 32:
			goto tr21
		case 33:
			goto tr18
		case 38:
			goto tr18
		case 44:
			goto tr18
		case 58:
			goto tr18
		case 124:
			goto tr18
		case 126:
			goto tr18
		}
		switch {
		case data[p] < 40:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr21
			}
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr18
			}
		default:
			goto tr18
		}
		goto st3
tr21:
//line NONE:1
te = p+1

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line ldap_injection.go:453
		switch data[p] {
		case 32:
			goto tr21
		case 33:
			goto tr18
		case 38:
			goto tr18
		case 44:
			goto tr18
		case 58:
			goto tr18
		case 61:
			goto st10
		case 124:
			goto tr18
		case 126:
			goto st4
		}
		switch {
		case data[p] < 40:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr21
			}
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto st4
			}
		default:
			goto tr18
		}
		goto st3
tr20:
//line NONE:1
te = p+1

	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line ldap_injection.go:495
		switch data[p] {
		case 33:
			goto tr18
		case 38:
			goto tr18
		case 44:
			goto tr18
		case 58:
			goto tr18
		case 124:
			goto tr18
		case 126:
			goto tr18
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr18
			}
		case data[p] >= 40:
			goto tr18
		}
		goto st3
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 40 {
			goto tr7
		}
		goto tr16
	st_out:
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof0: cs = 0; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 6:
			goto tr16
		case 0:
			goto tr0
		case 7:
			goto tr17
		case 1:
			goto tr3
		case 8:
			goto tr16
		case 2:
			goto tr0
		case 9:
			goto tr18
		case 3:
			goto tr8
		case 4:
			goto tr8
		case 10:
			goto tr18
		case 11:
			goto tr18
		case 12:
			goto tr18
		case 13:
			goto tr18
		case 14:
			goto tr16
		}
	}

	}

//line ldap_injection.rl:53

        return false
}

