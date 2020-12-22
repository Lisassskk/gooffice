package mergesort

func MergeSort(array []float64) []float64 {
	if len(array) <= 1 {
		_f := make([]float64, len(array))
		copy(_f, array)
		return _f
	}
	_gc := len(array) / 2
	_c := MergeSort(array[:_gc])
	_cc := MergeSort(array[_gc:])
	_d := make([]float64, len(array))
	_dd := 0
	_fg := 0
	_fb := 0
	for _fg < len(_c) && _fb < len(_cc) {
		if _c[_fg] <= _cc[_fb] {
			_d[_dd] = _c[_fg]
			_fg++
		} else {
			_d[_dd] = _cc[_fb]
			_fb++
		}
		_dd++
	}
	for _fg < len(_c) {
		_d[_dd] = _c[_fg]
		_fg++
		_dd++
	}
	for _fb < len(_cc) {
		_d[_dd] = _cc[_fb]
		_fb++
		_dd++
	}
	return _d
}
