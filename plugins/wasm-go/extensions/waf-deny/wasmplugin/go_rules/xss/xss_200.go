//line xss_200.rl:1
package xss

func matchXSS200(data []byte) bool {

	//line xss_200.rl:6

	//line xss_200.go:11
	var _xss200_actions []byte = []byte{
		0, 1, 0, 1, 1, 1, 2, 1, 5,
		1, 6, 1, 7, 1, 8, 1, 9,
		2, 2, 3, 2, 2, 4,
	}

	var _xss200_key_offsets []byte = []byte{
		0, 1, 3, 5, 7, 9, 11, 13,
		15, 16, 22, 24, 25,
	}

	var _xss200_trans_keys []byte = []byte{
		118, 109, 118, 108, 118, 102, 118, 114,
		118, 97, 118, 109, 118, 101, 118, 99,
		32, 43, 47, 61, 9, 13, 60, 115,
		118, 114,
	}

	var _xss200_single_lengths []byte = []byte{
		1, 2, 2, 2, 2, 2, 2, 2,
		1, 4, 2, 1, 1,
	}

	var _xss200_range_lengths []byte = []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 0, 0,
	}

	var _xss200_index_offsets []byte = []byte{
		0, 2, 5, 8, 11, 14, 17, 20,
		23, 25, 31, 34, 36,
	}

	var _xss200_indicies []byte = []byte{
		2, 1, 3, 2, 1, 4, 2, 1,
		5, 2, 1, 6, 2, 1, 7, 2,
		1, 8, 2, 1, 9, 2, 1, 11,
		10, 11, 11, 11, 12, 11, 10, 14,
		15, 13, 2, 1, 17, 16,
	}

	var _xss200_trans_targs []byte = []byte{
		10, 0, 1, 2, 3, 4, 5, 6,
		7, 11, 10, 9, 10, 10, 11, 12,
		10, 8,
	}

	var _xss200_trans_actions []byte = []byte{
		15, 0, 0, 0, 0, 0, 0, 0,
		0, 17, 13, 0, 7, 9, 20, 5,
		11, 0,
	}

	var _xss200_to_state_actions []byte = []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 0, 0,
	}

	var _xss200_from_state_actions []byte = []byte{
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 3, 0, 0,
	}

	var _xss200_eof_trans []byte = []byte{
		1, 1, 1, 1, 1, 1, 1, 1,
		11, 11, 0, 1, 17,
	}

	const xss200_start int = 10
	const xss200_first_final int = 10
	const xss200_error int = -1

	const xss200_en_main int = 10

	//line xss_200.rl:7
	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	var ts, te, act int
	_, _, _ = ts, te, act

	step := 0

	//line xss_200.go:97
	{
		cs = xss200_start
		ts = 0
		te = 0
		act = 0
	}

	//line xss_200.go:105
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
		_acts = int(_xss200_from_state_actions[cs])
		_nacts = uint(_xss200_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss200_actions[_acts-1] {
			case 1:
				//line NONE:1
				ts = p

				//line xss_200.go:125
			}
		}

		_keys = int(_xss200_key_offsets[cs])
		_trans = int(_xss200_index_offsets[cs])

		_klen = int(_xss200_single_lengths[cs])
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
				case data[p] < _xss200_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _xss200_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_xss200_range_lengths[cs])
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
				case data[p] < _xss200_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _xss200_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_xss200_indicies[_trans])
	_eof_trans:
		cs = int(_xss200_trans_targs[_trans])

		if _xss200_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_xss200_trans_actions[_trans])
		_nacts = uint(_xss200_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss200_actions[_acts-1] {
			case 2:
				//line NONE:1
				te = p + 1

			case 3:
				//line xss_200.rl:21
				act = 1
			case 4:
				//line xss_200.rl:33
				act = 3
			case 5:
				//line xss_200.rl:25
				te = p + 1
				{
					if step == 1 {
						return true
					} else {
						step = 0
					}
				}
			case 6:
				//line xss_200.rl:33
				te = p + 1

			case 7:
				//line xss_200.rl:33
				te = p
				p--

			case 8:
				//line xss_200.rl:33
				p = (te) - 1

			case 9:
				//line NONE:1
				switch act {
				case 1:
					{
						p = (te) - 1

						step = 1
					}
				default:
					{
						p = (te) - 1
					}
				}

				//line xss_200.go:241
			}
		}

	_again:
		_acts = int(_xss200_to_state_actions[cs])
		_nacts = uint(_xss200_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _xss200_actions[_acts-1] {
			case 0:
				//line NONE:1
				ts = 0

				//line xss_200.go:255
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
			if _xss200_eof_trans[cs] > 0 {
				_trans = int(_xss200_eof_trans[cs] - 1)
				goto _eof_trans
			}
		}

	}

	//line xss_200.rl:38

	return false
}
