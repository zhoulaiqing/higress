//line xss_370.rl:1
package xss

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

type machine370 struct {
	Builder strings.Builder
	// 0. None; 1. Passed word; 2. Passed left;
	step int
}

var keys370 = []string{"self", "document", "this", "top", "window"}

func (m *machine370) appendWord(s string) {
	m.Builder.WriteString(s)
}

func (m *machine370) checkWord() {
	if m.step >= 1 {
		return
	}

	s := m.Builder.String()
	fmt.Println(s)

	if slices.Contains(keys370, s) {
		m.step = 1
	} else {
		m.Builder.Reset()
	}
}

func (m *machine370) checkLeft() {
	if m.step >= 1 {
		m.step = 2
	} else {
		m.checkWord()
		if m.step >= 1 {
			m.step = 2
		} else {
			m.Builder.Reset()
		}
	}
}

func (m *machine370) checkRight() bool {
	if m.step == 2 {
		return true
	} else {
		m.Builder.Reset()
		return false
	}
}

func matchXSS370(data []byte) bool {

	//line xss_370.rl:60

	//line xss_370.go:65
	var _xss370_actions []byte = []byte{
		0, 1, 0, 1, 1, 1, 2, 1, 3,
		1, 4, 1, 5, 1, 6, 1, 7,
	}

	var _xss370_key_offsets []byte = []byte{
		0, 12, 13, 14,
	}

	var _xss370_trans_keys []byte = []byte{
		32, 41, 42, 47, 91, 93, 9, 13,
		65, 90, 97, 122, 47, 42, 65, 90,
		97, 122,
	}

	var _xss370_single_lengths []byte = []byte{
		6, 1, 1, 0,
	}

	var _xss370_range_lengths []byte = []byte{
		3, 0, 0, 2,
	}

	var _xss370_index_offsets []byte = []byte{
		0, 10, 12, 14,
	}

	var _xss370_trans_targs []byte = []byte{
		0, 0, 1, 2, 0, 0, 0, 3,
		3, 0, 0, 0, 0, 0, 3, 3,
		0, 0, 0, 0,
	}

	var _xss370_trans_actions []byte = []byte{
		9, 5, 0, 0, 5, 7, 9, 0,
		0, 11, 7, 15, 5, 15, 0, 0,
		13, 15, 15, 13,
	}

	var _xss370_to_state_actions []byte = []byte{
		1, 0, 0, 0,
	}

	var _xss370_from_state_actions []byte = []byte{
		3, 0, 0, 0,
	}

	var _xss370_eof_trans []byte = []byte{
		0, 19, 19, 20,
	}

	const xss370_start int = 0
	const xss370_first_final int = 0
	const xss370_error int = -1

	const xss370_en_main int = 0

	//line xss_370.rl:61
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	var ts, te, act int
	_ = act

	m := &machine370{}

	//line xss_370.go:134
	{
		cs = xss370_start
		ts = 0
		te = 0
		act = 0
	}

	//line xss_370.go:142
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
		_acts = int(_xss370_from_state_actions[cs])
		_nacts = uint(_xss370_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss370_actions[_acts-1] {
			case 1:
				//line NONE:1
				ts = p

				//line xss_370.go:162
			}
		}

		_keys = int(_xss370_key_offsets[cs])
		_trans = int(_xss370_index_offsets[cs])

		_klen = int(_xss370_single_lengths[cs])
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
				case data[p] < _xss370_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _xss370_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_xss370_range_lengths[cs])
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
				case data[p] < _xss370_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _xss370_trans_keys[_mid+1]:
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
		cs = int(_xss370_trans_targs[_trans])

		if _xss370_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_xss370_trans_actions[_trans])
		_nacts = uint(_xss370_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss370_actions[_acts-1] {
			case 2:
				//line xss_370.rl:76
				te = p + 1
				{
					m.checkLeft()
				}
			case 3:
				//line xss_370.rl:80
				te = p + 1
				{
					if m.checkRight() {
						return true
					}
				}
			case 4:
				//line xss_370.rl:86
				te = p + 1
				{
					m.checkWord()
				}
			case 5:
				//line xss_370.rl:90
				te = p + 1
				{
					if m.step < 2 {
						m.Builder.Reset()
					}

				}
			case 6:
				//line xss_370.rl:72
				te = p
				p--
				{
					m.appendWord(string(data[ts:te]))
				}
			case 7:
				//line xss_370.rl:90
				te = p
				p--
				{
					if m.step < 2 {
						m.Builder.Reset()
					}

				}
				//line xss_370.go:277
			}
		}

	_again:
		_acts = int(_xss370_to_state_actions[cs])
		_nacts = uint(_xss370_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss370_actions[_acts-1] {
			case 0:
				//line NONE:1
				ts = 0

				//line xss_370.go:291
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
			if _xss370_eof_trans[cs] > 0 {
				_trans = int(_xss370_eof_trans[cs] - 1)
				goto _eof_trans
			}
		}

	}

	//line xss_370.rl:100

	return false
}
