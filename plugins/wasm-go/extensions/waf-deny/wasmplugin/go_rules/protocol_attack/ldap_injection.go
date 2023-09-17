
//line ldap_injection.rl:1
package protocol_attack

func matchLdapInjection(data []byte) bool {

//line ldap_injection.rl:5

//line ldap_injection.go:10
const ldapInjection_start int = 1
const ldapInjection_first_final int = 7
const ldapInjection_error int = 0

const ldapInjection_en_main int = 1


//line ldap_injection.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof


    
//line ldap_injection.go:24
	{
	cs = ldapInjection_start
	}

//line ldap_injection.go:29
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
	case 6:
		goto st_case_6
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 33:
			goto st0
		case 38:
			goto st0
		case 40:
			goto st0
		case 41:
			goto st2
		case 58:
			goto st0
		case 60:
			goto st0
		case 62:
			goto st0
		case 124:
			goto st0
		case 126:
			goto st0
		}
		goto st1
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 32:
			goto st3
		case 33:
			goto st5
		case 38:
			goto st6
		case 40:
			goto st7
		case 41:
			goto st4
		case 124:
			goto st6
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 32:
			goto st3
		case 41:
			goto st4
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st3
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 33:
			goto st5
		case 38:
			goto st6
		case 40:
			goto st7
		case 124:
			goto st6
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 32:
			goto st5
		case 40:
			goto st7
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st5
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 33:
			goto tr9
		case 38:
			goto tr10
		case 44:
			goto tr11
		case 58:
			goto tr11
		case 124:
			goto tr10
		case 126:
			goto tr11
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr11
			}
		case data[p] >= 40:
			goto tr11
		}
		goto tr8
tr8:
//line ldap_injection.rl:12

            return true
        
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line ldap_injection.go:205
		switch data[p] {
		case 33:
			goto tr11
		case 38:
			goto tr11
		case 44:
			goto tr11
		case 58:
			goto tr11
		case 61:
			goto tr13
		case 124:
			goto tr11
		case 126:
			goto tr12
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr12
			}
		case data[p] >= 40:
			goto tr11
		}
		goto tr8
tr11:
//line ldap_injection.rl:12

            return true
        
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line ldap_injection.go:242
		goto tr11
tr12:
//line ldap_injection.rl:12

            return true
        
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line ldap_injection.go:255
		if data[p] == 61 {
			goto tr13
		}
		goto tr11
tr13:
//line ldap_injection.rl:12

            return true
        
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line ldap_injection.go:271
		switch data[p] {
		case 33:
			goto tr11
		case 38:
			goto tr11
		case 58:
			goto tr11
		case 60:
			goto tr11
		case 62:
			goto tr11
		case 124:
			goto tr11
		case 126:
			goto tr11
		}
		if 40 <= data[p] && data[p] <= 41 {
			goto tr11
		}
		goto tr13
tr9:
//line ldap_injection.rl:12

            return true
        
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line ldap_injection.go:303
		switch data[p] {
		case 32:
			goto tr14
		case 33:
			goto tr11
		case 38:
			goto tr11
		case 44:
			goto tr11
		case 58:
			goto tr11
		case 124:
			goto tr11
		case 126:
			goto tr11
		}
		switch {
		case data[p] < 40:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr14
			}
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr8
tr14:
//line ldap_injection.rl:12

            return true
        
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line ldap_injection.go:344
		switch data[p] {
		case 32:
			goto tr14
		case 33:
			goto tr11
		case 38:
			goto tr11
		case 44:
			goto tr11
		case 58:
			goto tr11
		case 61:
			goto tr13
		case 124:
			goto tr11
		case 126:
			goto tr12
		}
		switch {
		case data[p] < 40:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr14
			}
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr12
			}
		default:
			goto tr11
		}
		goto tr8
tr10:
//line ldap_injection.rl:12

            return true
        
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line ldap_injection.go:387
		switch data[p] {
		case 33:
			goto tr11
		case 38:
			goto tr11
		case 44:
			goto tr11
		case 58:
			goto tr11
		case 124:
			goto tr11
		case 126:
			goto tr11
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr11
			}
		case data[p] >= 40:
			goto tr11
		}
		goto tr8
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 40 {
			goto st7
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 7, 11, 12, 13, 14:
//line ldap_injection.rl:12

            return true
        
//line ldap_injection.go:444
		}
	}

	_out: {}
	}

//line ldap_injection.rl:31

        return false
}

