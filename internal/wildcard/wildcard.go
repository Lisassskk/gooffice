package wildcard

func Match(pattern, name string) (_bg bool) {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "\u002a" {
		return true
	}
	_bd := make([]rune, 0, len(name))
	_gf := make([]rune, 0, len(pattern))
	for _, _a := range name {
		_bd = append(_bd, _a)
	}
	for _, _d := range pattern {
		_gf = append(_gf, _d)
	}
	_ed := false
	return _gcb(_bd, _gf, _ed)
}
func MatchSimple(pattern, name string) bool {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "\u002a" {
		return true
	}
	_b := make([]rune, 0, len(name))
	_c := make([]rune, 0, len(pattern))
	for _, _e := range name {
		_b = append(_b, _e)
	}
	for _, _cg := range pattern {
		_c = append(_c, _cg)
	}
	_gc := true
	return _gcb(_b, _c, _gc)
}
func _eb(_ab, _aa []rune, _aeg int) int {
	for len(_aa) > 0 {
		switch _aa[0] {
		default:
			if len(_ab) == 0 {
				return -1
			}
			if _ab[0] != _aa[0] {
				return _eb(_ab[1:], _aa, _aeg+1)
			}
		case '?':
			if len(_ab) == 0 {
				return -1
			}
		case '*':
			if len(_ab) == 0 {
				return -1
			}
			_aab := _eb(_ab, _aa[1:], _aeg)
			if _aab != -1 {
				return _aeg
			} else {
				_aab = _eb(_ab[1:], _aa, _aeg)
				if _aab != -1 {
					return _aeg
				} else {
					return -1
				}
			}
		}
		_ab = _ab[1:]
		_aa = _aa[1:]
	}
	return _aeg
}
func _gcb(_cc, _bb []rune, _ae bool) bool {
	for len(_bb) > 0 {
		switch _bb[0] {
		default:
			if len(_cc) == 0 || _cc[0] != _bb[0] {
				return false
			}
		case '?':
			if len(_cc) == 0 && !_ae {
				return false
			}
		case '*':
			return _gcb(_cc, _bb[1:], _ae) || (len(_cc) > 0 && _gcb(_cc[1:], _bb, _ae))
		}
		_cc = _cc[1:]
		_bb = _bb[1:]
	}
	return len(_cc) == 0 && len(_bb) == 0
}
func Index(pattern, name string) (_aef int) {
	if pattern == "" || pattern == "\u002a" {
		return 0
	}
	_gdf := make([]rune, 0, len(name))
	_ba := make([]rune, 0, len(pattern))
	for _, _ce := range name {
		_gdf = append(_gdf, _ce)
	}
	for _, _db := range pattern {
		_ba = append(_ba, _db)
	}
	return _eb(_gdf, _ba, 0)
}
