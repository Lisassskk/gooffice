package algo

import _a "strconv"

func _e(_c byte) bool { return _c >= '0' && _c <= '9' }
func RepeatString(s string, cnt int) string {
	if cnt <= 0 {
		return ""
	}
	_fg := make([]byte, len(s)*cnt)
	_bb := []byte(s)
	for _df := 0; _df < cnt; _df++ {
		copy(_fg[_df:], _bb)
	}
	return string(_fg)
}

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess(lhs, rhs string) bool {
	_b, _ab := 0, 0
	for _b < len(lhs) && _ab < len(rhs) {
		_abf := lhs[_b]
		_cc := rhs[_ab]
		_d := _e(_abf)
		_dd := _e(_cc)
		switch {
		case _d && !_dd:
			return true
		case !_d && _dd:
			return false
		case !_d && !_dd:
			if _abf != _cc {
				return _abf < _cc
			}
			_b++
			_ab++
		default:
			_be := _b + 1
			_ag := _ab + 1
			for _be < len(lhs) && _e(lhs[_be]) {
				_be++
			}
			for _ag < len(rhs) && _e(rhs[_ag]) {
				_ag++
			}
			_ae, _ := _a.ParseUint(lhs[_b:_be], 10, 64)
			_cb, _ := _a.ParseUint(rhs[_b:_ag], 10, 64)
			if _ae != _cb {
				return _ae < _cb
			}
			_b = _be
			_ab = _ag
		}
	}
	return len(lhs) < len(rhs)
}
