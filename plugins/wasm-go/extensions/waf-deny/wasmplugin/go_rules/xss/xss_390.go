//line xss_390.rl:1
package xss

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

type machine390 struct {
	Builder     strings.Builder
	matchedWord bool
}

var keys390 = []string{"eval", "settimeout", "setinterval", "newfunction", "alert", "atob", "btoa"}

func (m *machine390) appendWord(s string) {
	m.Builder.WriteString(s)
}

func (m *machine390) reset() {
	m.Builder.Reset()
	m.matchedWord = false
}

func (m *machine390) checkSpace() {
	if m.matchedWord {
		return
	}

	if m.checkWord() >= 0 {
		return
	}

	m.reset()
}

func (m *machine390) checkWord() int {
	s := m.Builder.String()
	fmt.Println(s)

	if slices.Contains(keys390, s) {
		m.matchedWord = true
		return 1
	} else if s == "new" {
		return 0
	}

	return -1
}

func (m *machine390) checkFinish() bool {
	if m.matchedWord {
		return true
	}

	if m.checkWord() > 0 {
		return true
	}

	return false
}

func matchXSS390(data []byte) bool {

	//line xss_390.rl:65

	//line xss_390.go:70
	var _xss390_actions []byte = []byte{
		0, 1, 0, 1, 1, 1, 2, 1, 3,
		1, 4, 1, 5,
	}

	var _xss390_key_offsets []byte = []byte{
		0, 8,
	}

	var _xss390_trans_keys []byte = []byte{
		32, 40, 9, 13, 65, 90, 97, 122,
		65, 90, 97, 122,
	}

	var _xss390_single_lengths []byte = []byte{
		2, 0,
	}

	var _xss390_range_lengths []byte = []byte{
		3, 2,
	}

	var _xss390_index_offsets []byte = []byte{
		0, 6,
	}

	var _xss390_trans_targs []byte = []byte{
		0, 0, 0, 1, 1, 0, 1, 1,
		0, 0,
	}

	var _xss390_trans_actions []byte = []byte{
		7, 5, 7, 0, 0, 9, 0, 0,
		11, 11,
	}

	var _xss390_to_state_actions []byte = []byte{
		1, 0,
	}

	var _xss390_from_state_actions []byte = []byte{
		3, 0,
	}

	var _xss390_eof_trans []byte = []byte{
		0, 10,
	}

	const xss390_start int = 0
	const xss390_first_final int = 0
	const xss390_error int = -1

	const xss390_en_main int = 0

	//line xss_390.rl:66
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	var ts, te, act int
	_ = act

	m := &machine390{}

	//line xss_390.go:136
	{
		cs = xss390_start
		ts = 0
		te = 0
		act = 0
	}

	//line xss_390.go:144
	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if p == pe {
			goto _test_eof
		}
	_resume:
		_acts = int(_xss390_from_state_actions[cs])
		_nacts = uint(_xss390_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss390_actions[_acts-1] {
			case 1:
				//line NONE:1
				ts = p

				//line xss_390.go:164
			}
		}

		_keys = int(_xss390_key_offsets[cs])
		_trans = int(_xss390_index_offsets[cs])

		_klen = int(_xss390_single_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case data[p] < _xss390_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _xss390_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_xss390_range_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case data[p] < _xss390_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _xss390_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
	_eof_trans:
		cs = int(_xss390_trans_targs[_trans])

		if _xss390_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_xss390_trans_actions[_trans])
		_nacts = uint(_xss390_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss390_actions[_acts-1] {
			case 2:
				//line xss_390.rl:81
				te = p + 1
				{
					if m.checkFinish() {
						return true
					}
				}
			case 3:
				//line xss_390.rl:87
				te = p + 1
				{
					m.checkSpace()
				}
			case 4:
				//line xss_390.rl:91
				te = p + 1
				{
					m.reset()
				}
			case 5:
				//line xss_390.rl:77
				te = p
				p--
				{
					m.appendWord(string(data[ts:te]))
				}
				//line xss_390.go:260
			}
		}

	_again:
		_acts = int(_xss390_to_state_actions[cs])
		_nacts = uint(_xss390_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss390_actions[_acts-1] {
			case 0:
				//line NONE:1
				ts = 0

				//line xss_390.go:274
			}
		}

		p++
		if p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
		if p == eof {
			if _xss390_eof_trans[cs] > 0 {
				_trans = int(_xss390_eof_trans[cs] - 1)
				goto _eof_trans
			}
		}

	}

	//line xss_390.rl:98

	return false
}
