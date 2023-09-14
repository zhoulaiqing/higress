
//line xss_140.rl:1
package rule_941


func matchXSS140(data []byte) bool {

//line xss_140.rl:6

//line xss_140.go:11
var _xss140_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 5, 1, 6, 
	1, 7, 1, 8, 1, 9, 2, 2, 
	3, 2, 2, 4, 
}

var _xss140_key_offsets []byte = []byte{
	0, 2, 4, 5, 6, 8, 10, 12, 
	14, 16, 18, 20, 22, 24, 26, 28, 
	30, 32, 34, 36, 38, 41, 43, 46, 
}

var _xss140_trans_keys []byte = []byte{
	58, 61, 58, 61, 117, 59, 58, 61, 
	58, 61, 59, 117, 59, 114, 59, 108, 
	40, 59, 59, 106, 59, 97, 59, 118, 
	59, 97, 59, 115, 59, 99, 59, 114, 
	59, 105, 59, 112, 59, 116, 61, 97, 
	122, 97, 122, 61, 97, 122, 59, 
}

var _xss140_single_lengths []byte = []byte{
	2, 2, 1, 1, 2, 2, 2, 2, 
	2, 2, 2, 2, 2, 2, 2, 2, 
	2, 2, 2, 2, 1, 0, 1, 1, 
}

var _xss140_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 1, 1, 0, 
}

var _xss140_index_offsets []byte = []byte{
	0, 3, 6, 8, 10, 13, 16, 19, 
	22, 25, 28, 31, 34, 37, 40, 43, 
	46, 49, 52, 55, 58, 61, 63, 66, 
}

var _xss140_indicies []byte = []byte{
	0, 0, 1, 2, 0, 1, 4, 3, 
	6, 3, 3, 3, 7, 8, 3, 7, 
	6, 4, 3, 6, 9, 3, 6, 10, 
	3, 11, 6, 3, 6, 12, 3, 6, 
	13, 3, 6, 14, 3, 6, 15, 3, 
	6, 16, 3, 6, 17, 3, 6, 18, 
	3, 6, 19, 3, 6, 20, 3, 6, 
	21, 3, 22, 23, 0, 25, 24, 22, 
	23, 26, 6, 3, 
}

var _xss140_trans_targs []byte = []byte{
	21, 1, 2, 3, 7, 21, 4, 5, 
	6, 8, 9, 10, 11, 12, 13, 14, 
	15, 16, 17, 18, 19, 23, 0, 20, 
	21, 22, 21, 21, 
}

var _xss140_trans_actions []byte = []byte{
	11, 0, 0, 0, 0, 13, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 15, 0, 0, 
	5, 18, 9, 7, 
}

var _xss140_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
}

var _xss140_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 3, 0, 0, 
}

var _xss140_eof_trans []byte = []byte{
	1, 1, 1, 6, 6, 6, 6, 6, 
	6, 6, 6, 6, 6, 6, 6, 6, 
	6, 6, 6, 6, 1, 0, 27, 28, 
}

const xss140_start int = 21
const xss140_first_final int = 21
const xss140_error int = -1

const xss140_en_main int = 21


//line xss_140.rl:7
    cs, p, pe, eof := 0, 0, len(data), len(data)
       _ = eof
    var ts, te, act int
    _, _, _ = ts, te, act

    
//line xss_140.go:109
	{
	cs = xss140_start
	ts = 0
	te = 0
	act = 0
	}

//line xss_140.go:117
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
	_acts = int(_xss140_from_state_actions[cs])
	_nacts = uint(_xss140_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _xss140_actions[_acts - 1] {
		case 1:
//line NONE:1
ts = p

//line xss_140.go:137
		}
	}

	_keys = int(_xss140_key_offsets[cs])
	_trans = int(_xss140_index_offsets[cs])

	_klen = int(_xss140_single_lengths[cs])
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
			case data[p] < _xss140_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _xss140_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_xss140_range_lengths[cs])
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
			case data[p] < _xss140_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _xss140_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_xss140_indicies[_trans])
_eof_trans:
	cs = int(_xss140_trans_targs[_trans])

	if _xss140_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_xss140_trans_actions[_trans])
	_nacts = uint(_xss140_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xss140_actions[_acts-1] {
		case 2:
//line NONE:1
te = p+1

		case 3:
//line xss_140.rl:15
act = 1;
		case 4:
//line xss_140.rl:19
act = 2;
		case 5:
//line xss_140.rl:19
te = p+1

		case 6:
//line xss_140.rl:15
te = p
p--
{
                return true
            }
		case 7:
//line xss_140.rl:19
te = p
p--

		case 8:
//line xss_140.rl:19
p = (te) - 1

		case 9:
//line NONE:1
	switch act {
	case 1:
	{p = (te) - 1

                return true
            }
	default:
	{p = (te) - 1
}
	}
	
//line xss_140.go:250
		}
	}

_again:
	_acts = int(_xss140_to_state_actions[cs])
	_nacts = uint(_xss140_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xss140_actions[_acts-1] {
		case 0:
//line NONE:1
ts = 0

//line xss_140.go:264
		}
	}

	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _xss140_eof_trans[cs] > 0 {
			_trans = int(_xss140_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	}

//line xss_140.rl:24


    return false
}
