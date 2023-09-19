//line short_command_base.rl:1
package rce_attack

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

var commands = []string{
	"fd", "xz", "ssh", "esh", "irb", "env", "scp", "7zx", "get", "bzz", "csh", "rcp", "rc", "ls", "svn", "gpg", "ksh",
	"zsh", "sh", "7z", "lz4", "udp", "sed", "7zr", "lz", "pxz", "fg", "w3m", "df", "hup", "7za", "php", "pwd",
}

func checkCommand(name string) bool {

	name = strings.ReplaceAll(name, ";", "")

	fmt.Printf("possible command name %s \n", name)
	if slices.Contains(commands, name) {
		return true
	}
	if strings.HasPrefix(name, "gcc") {
		return true
	}

	return false
}

func matchShortCommandBase(data []byte) bool {

	//line short_command_base.rl:31

	//line short_command_base.go:36
	const shortCommandBase_start int = 8
	const shortCommandBase_first_final int = 8
	const shortCommandBase_error int = -1

	const shortCommandBase_en_main int = 8

	//line short_command_base.rl:32
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	pb := 0
	var commandName string

	var ts, te, act int
	// _, _, _ = ts, te, act
	_ = act

	//line short_command_base.go:56
	{
		cs = shortCommandBase_start
		ts = 0
		te = 0
		act = 0
	}

	//line short_command_base.go:64
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
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
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 6:
			goto st_case_6
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 7:
			goto st_case_7
		}
		goto st_out
	tr0:
		//line short_command_base.rl:80
		p = (te) - 1

		goto st8
	tr8:
		//line NONE:1
		switch act {
		case 1:
			{
				p = (te) - 1

				fmt.Printf("Full command: %s \n", string(data[ts:te]))
				if checkCommand(commandName) {
					return true
				}
			}
		default:
			{
				p = (te) - 1
			}
		}

		goto st8
	tr12:
		//line short_command_base.rl:80
		te = p + 1

		goto st8
	tr16:
		//line short_command_base.rl:80
		te = p
		p--

		goto st8
	tr17:
		//line short_command_base.rl:52

		commandName = string(data[pb:p])

		//line short_command_base.rl:72
		te = p
		p--
		{

			fmt.Printf("Full command: %s \n", string(data[ts:te]))
			if checkCommand(commandName) {
				return true
			}
		}
		goto st8
	st8:
		//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		//line NONE:1
		ts = p

		//line short_command_base.go:161
		switch data[p] {
		case 10:
			goto tr13
		case 13:
			goto tr13
		case 36:
			goto tr14
		case 38:
			goto tr13
		case 40:
			goto tr15
		case 59:
			goto tr13
		case 60:
			goto tr14
		case 62:
			goto tr14
		case 96:
			goto tr13
		}
		if 123 <= data[p] && data[p] <= 124 {
			goto tr13
		}
		goto tr12
	tr13:
		//line NONE:1
		te = p + 1

		//line short_command_base.rl:80
		act = 2
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		//line short_command_base.go:198
		switch data[p] {
		case 36:
			goto st1
		case 40:
			goto st0
		case 47:
			goto st4
		case 59:
			goto tr6
		case 63:
			goto st3
		case 91:
			goto st3
		case 93:
			goto st3
		case 95:
			goto tr5
		case 123:
			goto st1
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st0
				}
			case data[p] > 33:
				if 41 <= data[p] && data[p] <= 43 {
					goto st3
				}
			default:
				goto st0
			}
		case data[p] > 46:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr5
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr5
				}
			default:
				goto tr5
			}
		default:
			goto st3
		}
		goto tr16
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 36:
			goto st1
		case 40:
			goto st0
		case 47:
			goto st4
		case 59:
			goto tr6
		case 63:
			goto st3
		case 91:
			goto st3
		case 93:
			goto st3
		case 95:
			goto tr5
		case 123:
			goto st1
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] < 32:
				if 9 <= data[p] && data[p] <= 13 {
					goto st0
				}
			case data[p] > 33:
				if 41 <= data[p] && data[p] <= 43 {
					goto st3
				}
			default:
				goto st0
			}
		case data[p] > 46:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr5
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr5
				}
			default:
				goto tr5
			}
		default:
			goto st3
		}
		goto tr0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st2
		case 33:
			goto st0
		case 36:
			goto st1
		case 40:
			goto st0
		case 47:
			goto st4
		case 59:
			goto tr6
		case 63:
			goto st3
		case 91:
			goto st3
		case 93:
			goto st3
		case 95:
			goto tr5
		case 123:
			goto st1
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 41 <= data[p] && data[p] <= 43 {
					goto st3
				}
			case data[p] >= 9:
				goto st2
			}
		case data[p] > 46:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr5
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr5
				}
			default:
				goto tr5
			}
		default:
			goto st3
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 32:
			goto st2
		case 40:
			goto st0
		case 47:
			goto st4
		case 59:
			goto tr6
		case 63:
			goto st3
		case 91:
			goto st3
		case 93:
			goto st3
		case 95:
			goto tr5
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 41 <= data[p] && data[p] <= 43 {
					goto st3
				}
			case data[p] >= 9:
				goto st2
			}
		case data[p] > 46:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr5
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr5
				}
			default:
				goto tr5
			}
		default:
			goto st3
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 47 {
			goto st4
		}
		goto tr8
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 47:
			goto st4
		case 59:
			goto tr6
		case 63:
			goto st3
		case 91:
			goto st3
		case 93:
			goto st3
		case 95:
			goto tr5
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st3
				}
			case data[p] >= 40:
				goto st3
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr5
				}
			case data[p] >= 65:
				goto tr5
			}
		default:
			goto tr5
		}
		goto tr8
	tr5:
		//line short_command_base.rl:48

		pb = p

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		//line short_command_base.go:474
		switch data[p] {
		case 47:
			goto st4
		case 59:
			goto st11
		case 95:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr8
	tr9:
		//line NONE:1
		te = p + 1

		//line short_command_base.rl:72
		act = 1
		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		//line short_command_base.go:508
		switch data[p] {
		case 47:
			goto st4
		case 59:
			goto st11
		case 95:
			goto tr9
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr9
			}
		default:
			goto tr9
		}
		goto tr17
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 59:
			goto st11
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr17
	tr6:
		//line short_command_base.rl:48

		pb = p

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		//line short_command_base.go:565
		switch data[p] {
		case 59:
			goto st11
		case 95:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		goto tr8
	tr14:
		//line NONE:1
		te = p + 1

		//line short_command_base.rl:80
		act = 2
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		//line short_command_base.go:597
		if data[p] == 40 {
			goto st0
		}
		goto tr16
	tr15:
		//line NONE:1
		te = p + 1

		//line short_command_base.rl:80
		act = 2
		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		//line short_command_base.go:614
		switch data[p] {
		case 32:
			goto st7
		case 41:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st7
		}
		goto tr16
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 32:
			goto st7
		case 41:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st7
		}
		goto tr0
	st_out:
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof0:
		cs = 0
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 9:
				goto tr16
			case 0:
				goto tr0
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr8
			case 4:
				goto tr8
			case 5:
				goto tr8
			case 10:
				goto tr17
			case 11:
				goto tr17
			case 6:
				goto tr8
			case 12:
				goto tr16
			case 13:
				goto tr16
			case 7:
				goto tr0
			}
		}

	}

	//line short_command_base.rl:85

	return false
}
