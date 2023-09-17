
//line ldap_injection.rl:1
package protocol_attack

func matchLdapInjection(data []byte) bool {

//line ldap_injection.rl:5

//line ldap_injection.go:10
var _ldapInjection_actions []byte = []byte{
	0, 1, 0, 
}

var _ldapInjection_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 1, 1, 1, 1, 
}

const ldapInjection_start int = 1
const ldapInjection_first_final int = 7
const ldapInjection_error int = 0

const ldapInjection_en_main int = 1


//line ldap_injection.rl:6
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof


    
//line ldap_injection.go:33
	{
	cs = ldapInjection_start
	}

//line ldap_injection.go:38
	{
	var _acts int
	var _nacts uint

	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	switch cs {
	case 1:
		switch data[p] {
		case 33:
			goto tr1;
		case 38:
			goto tr1;
		case 40:
			goto tr1;
		case 41:
			goto tr2;
		case 58:
			goto tr1;
		case 60:
			goto tr1;
		case 62:
			goto tr1;
		case 124:
			goto tr1;
		case 126:
			goto tr1;
		}
		goto tr0;
	case 0:
		goto _out
	case 2:
		switch data[p] {
		case 32:
			goto tr3;
		case 33:
			goto tr4;
		case 38:
			goto tr5;
		case 40:
			goto tr6;
		case 41:
			goto tr7;
		case 124:
			goto tr5;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr3;
		}
		goto tr1;
	case 3:
		switch data[p] {
		case 32:
			goto tr3;
		case 41:
			goto tr7;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr3;
		}
		goto tr1;
	case 4:
		switch data[p] {
		case 33:
			goto tr4;
		case 38:
			goto tr5;
		case 40:
			goto tr6;
		case 124:
			goto tr5;
		}
		goto tr1;
	case 5:
		switch data[p] {
		case 32:
			goto tr4;
		case 40:
			goto tr6;
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr4;
		}
		goto tr1;
	case 7:
		switch data[p] {
		case 33:
			goto tr9;
		case 38:
			goto tr10;
		case 44:
			goto tr11;
		case 58:
			goto tr11;
		case 124:
			goto tr10;
		case 126:
			goto tr11;
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr11;
			}
		case data[p] >= 40:
			goto tr11;
		}
		goto tr8;
	case 8:
		switch data[p] {
		case 33:
			goto tr11;
		case 38:
			goto tr11;
		case 44:
			goto tr11;
		case 58:
			goto tr11;
		case 61:
			goto tr13;
		case 124:
			goto tr11;
		case 126:
			goto tr12;
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr12;
			}
		case data[p] >= 40:
			goto tr11;
		}
		goto tr8;
	case 9:
		goto tr11;
	case 10:
		if data[p] == 61 {
			goto tr13;
		}
		goto tr11;
	case 11:
		switch data[p] {
		case 33:
			goto tr11;
		case 38:
			goto tr11;
		case 58:
			goto tr11;
		case 60:
			goto tr11;
		case 62:
			goto tr11;
		case 124:
			goto tr11;
		case 126:
			goto tr11;
		}
		if 40 <= data[p] && data[p] <= 41 {
			goto tr11;
		}
		goto tr13;
	case 12:
		switch data[p] {
		case 32:
			goto tr14;
		case 33:
			goto tr11;
		case 38:
			goto tr11;
		case 44:
			goto tr11;
		case 58:
			goto tr11;
		case 124:
			goto tr11;
		case 126:
			goto tr11;
		}
		switch {
		case data[p] < 40:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr14;
			}
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr11;
			}
		default:
			goto tr11;
		}
		goto tr8;
	case 13:
		switch data[p] {
		case 32:
			goto tr14;
		case 33:
			goto tr11;
		case 38:
			goto tr11;
		case 44:
			goto tr11;
		case 58:
			goto tr11;
		case 61:
			goto tr13;
		case 124:
			goto tr11;
		case 126:
			goto tr12;
		}
		switch {
		case data[p] < 40:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr14;
			}
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr12;
			}
		default:
			goto tr11;
		}
		goto tr8;
	case 14:
		switch data[p] {
		case 33:
			goto tr11;
		case 38:
			goto tr11;
		case 44:
			goto tr11;
		case 58:
			goto tr11;
		case 124:
			goto tr11;
		case 126:
			goto tr11;
		}
		switch {
		case data[p] > 41:
			if 60 <= data[p] && data[p] <= 62 {
				goto tr11;
			}
		case data[p] >= 40:
			goto tr11;
		}
		goto tr8;
	case 6:
		if data[p] == 40 {
			goto tr6;
		}
		goto tr1;
	}

	tr1: cs = 0; goto _again
	tr0: cs = 1; goto _again
	tr2: cs = 2; goto _again
	tr3: cs = 3; goto _again
	tr7: cs = 4; goto _again
	tr4: cs = 5; goto _again
	tr5: cs = 6; goto _again
	tr6: cs = 7; goto _again
	tr8: cs = 8; goto f0
	tr11: cs = 9; goto f0
	tr12: cs = 10; goto f0
	tr13: cs = 11; goto f0
	tr9: cs = 12; goto f0
	tr14: cs = 13; goto f0
	tr10: cs = 14; goto f0

	f0: _acts = 1; goto execFuncs

execFuncs:
	_nacts = uint(_ldapInjection_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _ldapInjection_actions[_acts - 1] {
		case 0:
//line ldap_injection.rl:12

            return true
        
//line ldap_injection.go:327
		}
	}
	goto _again

_again:
	if cs == 0 {
		goto _out
	}
	if p++; p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := int(_ldapInjection_eof_actions[cs])
		__nacts := uint(_ldapInjection_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _ldapInjection_actions[__acts - 1] {
			case 0:
//line ldap_injection.rl:12

            return true
        
//line ldap_injection.go:351
			}
		}
	}

	_out: {}
	}

//line ldap_injection.rl:31

        return false
}

