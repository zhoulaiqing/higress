
//line xss_210.rl:1
package rule_941

import (
    "fmt"
    "strings"
)

type machine210 struct {
    Builder strings.Builder
}

func (m *machine210) checkColon() bool {
    s := m.Builder.String()
    fmt.Println(s)

    if s == "javascript" || s == "vbscript" {
        m.Builder.Reset()
        return true
    }
    m.Builder.Reset()
    return false
}

func (m *machine210) checkHtmlSpace(word string) bool {
    fmt.Println(word)

    if word == "&newline;" || word == "&tab;" {
        return false
    } else if word == "&colon;" {
        return m.checkColon()
    }

    m.Builder.Reset()
    return false
}

func (m *machine210) appendWord(s string) {
    m.Builder.WriteString(s)
}

func matchXSS210(data []byte) bool {

//line xss_210.rl:43

//line xss_210.go:48
var _xss210_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 
}

var _xss210_key_offsets []byte = []byte{
	0, 5, 15, 19, 
}

var _xss210_trans_keys []byte = []byte{
	59, 65, 90, 97, 122, 0, 32, 38, 
	58, 9, 13, 65, 90, 97, 122, 65, 
	90, 97, 122, 65, 90, 97, 122, 
}

var _xss210_single_lengths []byte = []byte{
	1, 4, 0, 0, 
}

var _xss210_range_lengths []byte = []byte{
	2, 3, 2, 2, 
}

var _xss210_index_offsets []byte = []byte{
	0, 4, 12, 15, 
}

var _xss210_trans_targs []byte = []byte{
	1, 0, 0, 1, 1, 1, 2, 1, 
	1, 3, 3, 1, 0, 0, 1, 3, 
	3, 1, 1, 1, 1, 
}

var _xss210_trans_actions []byte = []byte{
	7, 0, 0, 21, 13, 11, 5, 9, 
	11, 0, 0, 15, 0, 0, 19, 0, 
	0, 17, 21, 19, 17, 
}

var _xss210_to_state_actions []byte = []byte{
	0, 1, 0, 0, 
}

var _xss210_from_state_actions []byte = []byte{
	0, 3, 0, 0, 
}

var _xss210_eof_trans []byte = []byte{
	19, 0, 20, 21, 
}

const xss210_start int = 1
const xss210_first_final int = 1
const xss210_error int = -1

const xss210_en_main int = 1


//line xss_210.rl:44
    cs, p, pe, eof := 0, 0, len(data), len(data)
        _ = eof

    var ts, te, act int
        _ = act

    m := &machine210{}

    
//line xss_210.go:118
	{
	cs = xss210_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_210.go:126
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
	_acts = int(_xss210_from_state_actions[cs])
	_nacts = uint(_xss210_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _xss210_actions[_acts - 1] {
		case 1:
//line NONE:1
ts = p

//line xss_210.go:146
		}
	}

	_keys = int(_xss210_key_offsets[cs])
	_trans = int(_xss210_index_offsets[cs])

	_klen = int(_xss210_single_lengths[cs])
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
			case data[p] < _xss210_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _xss210_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_xss210_range_lengths[cs])
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
			case data[p] < _xss210_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _xss210_trans_keys[_mid + 1]:
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
	cs = int(_xss210_trans_targs[_trans])

	if _xss210_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_xss210_trans_actions[_trans])
	_nacts = uint(_xss210_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xss210_actions[_acts-1] {
		case 2:
//line NONE:1
te = p+1

		case 3:
//line xss_210.rl:59
te = p+1
{
                if m.checkHtmlSpace(string(data[ts:te])) {
                    return true
                }
            }
		case 4:
//line xss_210.rl:69
te = p+1
{
                if m.checkColon() {
                    return true
                }
            }
		case 5:
//line xss_210.rl:75
te = p+1
{
            }
		case 6:
//line xss_210.rl:78
te = p+1
{
                m.Builder.Reset()
            }
		case 7:
//line xss_210.rl:82
te = p+1
{
                m.Builder.Reset()
            }
		case 8:
//line xss_210.rl:65
te = p
p--
{
                m.appendWord(string(data[ts:te]))
            }
		case 9:
//line xss_210.rl:82
te = p
p--
{
                m.Builder.Reset()
            }
		case 10:
//line xss_210.rl:82
p = (te) - 1
{
                m.Builder.Reset()
            }
//line xss_210.go:272
		}
	}

_again:
	_acts = int(_xss210_to_state_actions[cs])
	_nacts = uint(_xss210_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xss210_actions[_acts-1] {
		case 0:
//line NONE:1
ts = 0

//line xss_210.go:286
		}
	}

	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _xss210_eof_trans[cs] > 0 {
			_trans = int(_xss210_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	}

//line xss_210.rl:91


    return false
}
