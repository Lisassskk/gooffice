package formula

import (
	_gg "bytes"
	_f "errors"
	_d "fmt"
	_g "io"
	_cc "math"
	_gga "math/big"
	_fa "math/rand"
	_ge "regexp"
	_b "sort"
	_c "strconv"
	_dda "strings"
	_fe "sync"
	_ee "time"
	_dd "unicode"

	_bc "gitee.com/gooffice/gooffice"
	_feb "gitee.com/gooffice/gooffice/internal/mergesort"
	_ec "gitee.com/gooffice/gooffice/internal/wildcard"
	_bg "gitee.com/gooffice/gooffice/spreadsheet/format"
	_bb "gitee.com/gooffice/gooffice/spreadsheet/reference"
	_a "gitee.com/gooffice/gooffice/spreadsheet/update"
)

// String returns a string representation for Bool.
func (_dge Bool) String() string {
	if _dge._fd {
		return "\u0054\u0052\u0055\u0045"
	} else {
		return "\u0046\u0041\u004cS\u0045"
	}
}

// Yieldmat implements the Excel YIELDMAT function.
func Yieldmat(args []Result) Result {
	_daacd := len(args)
	if _daacd != 5 && _daacd != 6 {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0066\u0069v\u0065\u0020\u006f\u0072\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_fdcf, _geba, _fdbac := _gdge(args[0], args[1], "\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054")
	if _fdbac.Type == ResultTypeError {
		return _fdbac
	}
	_cbbc, _fdbac := _badf(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054")
	if _fdbac.Type == ResultTypeError {
		return _fdbac
	}
	if _cbbc >= _fdcf {
		return MakeErrorResult("\u0059\u0049\u0045\u004cD\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u0062\u0065\u0066\u006fr\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049E\u004c\u0044\u004d\u0041T\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_bdea := args[3].ValueNumber
	if _bdea < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0049\u0045\u004c\u0044M\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0079\u0069\u0065\u006c\u0064\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_daee := args[4].ValueNumber
	if _daee <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "Y\u0049\u0045\u004c\u0044\u004d\u0041T\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0070\u0072\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069v\u0065")
	}
	_cgef := 0
	if _daacd == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("\u0059I\u0045\u004cD\u004d\u0041\u0054 \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_cgef = int(args[5].ValueNumber)
		if !_gaff(_cgef) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0059\u0049\u0045L\u0044\u004d\u0041\u0054")
		}
	}
	_afebb, _fdbac := _agef(_cbbc, _geba, _cgef)
	if _fdbac.Type == ResultTypeError {
		return _fdbac
	}
	_ffedc, _fdbac := _agef(_cbbc, _fdcf, _cgef)
	if _fdbac.Type == ResultTypeError {
		return _fdbac
	}
	_fadc, _fdbac := _agef(_fdcf, _geba, _cgef)
	if _fdbac.Type == ResultTypeError {
		return _fdbac
	}
	_dfebc := 1 + _afebb*_bdea
	_dfebc /= _daee/100 + _ffedc*_bdea
	_dfebc--
	_dfebc /= _fadc
	return MakeNumberResult(_dfebc)
}

// Floor is an implementation of the FlOOR function.
func Floor(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0046\u004c\u004f\u004f\u0052\u0028\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_fdgad := args[0].AsNumber()
	if _fdgad.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0046\u004c\u004f\u004f\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	var _ebbfc float64
	_efcdd := args[1].AsNumber()
	if _efcdd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0046\u004c\u004f\u004f\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ebbfc = _efcdd.ValueNumber
	if _ebbfc < 0 && _fdgad.ValueNumber >= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0046L\u004f\u004f\u0052")
	}
	_fada := _fdgad.ValueNumber
	_fada, _ecae := _cc.Modf(_fada / _ebbfc)
	if _ecae != 0 {
		if _fdgad.ValueNumber < 0 && _ecae < 0 {
			_fada--
		}
	}
	return MakeNumberResult(_fada * _ebbfc)
}

var _feef = map[string]Function{}

const _afa = "\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u002f\u0028\u0028\u005b\u0030-\u0039]\u0029\u002b\u0029\u002f\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

func _cfaca(_dcab, _aeab float64, _efe, _agdd int) float64 {
	_bed := _dcge(_dcab)
	_deb := _dcge(_aeab)
	_feee := _dccg(_bed, _deb, _efe)
	return _fagd(_bed, _feee, _agdd)
}

// Cumipmt implements the Excel CUMIPMT function.
func Cumipmt(args []Result) Result {
	_afga, _cfae := _ebca(args, "\u0043U\u004d\u0049\u0050\u004d\u0054")
	if _cfae.Type == ResultTypeError {
		return _cfae
	}
	_bbda := _afga._egab
	_fdec := _afga._cefc
	_aca := _afga._dfaf
	_bebb := _afga._gfdb
	_bbcd := _afga._ggcf
	_bbbc := _afga._fgcf
	_fdfe := _acad(_bbda, _fdec, _aca, 0, _bbbc)
	_eebf := 0.0
	if _bebb == 1 {
		if _bbbc == 0 {
			_eebf = -_aca
			_bebb++
		}
	}
	for _bcba := _bebb; _bcba <= _bbcd; _bcba++ {
		if _bbbc == 1 {
			_eebf += _gfec(_bbda, _bcba-2, _fdfe, _aca, 1) - _fdfe
		} else {
			_eebf += _gfec(_bbda, _bcba-1, _fdfe, _aca, 0)
		}
	}
	_eebf *= _bbda
	return MakeNumberResult(_eebf)
}
func _ca(_abf BinOpType, _dc [][]Result, _dag Result) Result {
	_cd := [][]Result{}
	for _ag := range _dc {
		_eaa := _ga(_abf, _dc[_ag], _dag)
		if _eaa.Type == ResultTypeError {
			return _eaa
		}
		_cd = append(_cd, _eaa.ValueList)
	}
	return MakeArrayResult(_cd)
}

// Pricemat implements the Excel PRICEMAT function.
func Pricemat(args []Result) Result {
	_dagg := len(args)
	if _dagg != 5 && _dagg != 6 {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0066\u0069v\u0065\u0020\u006f\u0072\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_efae, _ccd, _ceef := _gdge(args[0], args[1], "\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054")
	if _ceef.Type == ResultTypeError {
		return _ceef
	}
	_deff, _ceef := _badf(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054")
	if _ceef.Type == ResultTypeError {
		return _ceef
	}
	if _deff >= _efae {
		return MakeErrorResult("\u0050\u0052\u0049\u0043E\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u0062\u0065\u0066\u006fr\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052I\u0043\u0045\u004d\u0041T\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_fbgb := args[3].ValueNumber
	if _fbgb < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045M\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0079\u0069\u0065\u006c\u0064\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_ccae := args[4].ValueNumber
	if _ccae < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049C\u0045\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0079\u0069\u0065\u006c\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	_bdbc := 0
	if _dagg == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050R\u0049\u0043E\u004d\u0041\u0054 \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_bdbc = int(args[5].ValueNumber)
		if !_gaff(_bdbc) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0050\u0052\u0049C\u0045\u004d\u0041\u0054")
		}
	}
	_dccbe, _ceef := _agef(_efae, _ccd, _bdbc)
	if _ceef.Type == ResultTypeError {
		return _ceef
	}
	_fagf, _ceef := _agef(_deff, _ccd, _bdbc)
	if _ceef.Type == ResultTypeError {
		return _ceef
	}
	_aaaa, _ceef := _agef(_deff, _efae, _bdbc)
	if _ceef.Type == ResultTypeError {
		return _ceef
	}
	_afac := 1 + _fagf*_fbgb
	_cfdc := 1 + _dccbe*_ccae
	return MakeNumberResult((_afac/_cfdc - _aaaa*_fbgb) * 100)
}
func _dcgec(_abdd, _bgcb, _adgb, _agag, _cfada float64, _dged int) Result {
	_abge, _bggd := _agef(_abdd, _bgcb, _dged)
	if _bggd.Type == ResultTypeError {
		return _bggd
	}
	_fead, _bga := _fbeb(_abdd, _bgcb, int(_cfada), _dged)
	if _bga.Type == ResultTypeError {
		return _bga
	}
	_fegd := 0.0
	_bcd := 0.0
	_adgb *= 100 / _cfada
	_agag /= _cfada
	_agag++
	_ebead := _abge*_cfada - _fead
	for _agce := 1.0; _agce < _fead; _agce++ {
		_degd := _agce + _ebead
		_fbfd := _adgb / _cc.Pow(_agag, _degd)
		_bcd += _fbfd
		_fegd += _degd * _fbfd
	}
	_ecfd := (_adgb + 100) / _cc.Pow(_agag, _fead+_ebead)
	_bcd += _ecfd
	_fegd += (_fead + _ebead) * _ecfd
	_fegd /= _bcd
	_fegd /= _cfada
	return MakeNumberResult(_fegd)
}
func _fffe(_gcdg, _gddg float64, _beb, _fcbd int) float64 {
	_cebf := _dcge(_gcdg)
	_aag := _dcge(_gddg)
	_bfcg := _efbe(_cebf, _aag, _beb, _fcbd)
	return _fagd(_bfcg, _cebf, _fcbd)
}

// Irr implements the Excel IRR function.
func Irr(args []Result) Result {
	_defd := len(args)
	if _defd == 0 || _defd > 2 {
		return MakeErrorResult("\u0049\u0052\u0052\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0072\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeList && args[0].Type != ResultTypeArray {
		return MakeErrorResult("\u0049\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020v\u0061\u006c\u0075\u0065\u0073\u0020t\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0061\u0072\u0072\u0061\u0079 \u0074\u0079\u0070\u0065")
	}
	_acd := _afgc(args[0])
	_feaa := []float64{}
	for _, _gce := range _acd {
		for _, _eagce := range _gce {
			if _eagce.Type == ResultTypeNumber && !_eagce.IsBoolean {
				_feaa = append(_feaa, _eagce.ValueNumber)
			}
		}
	}
	_egec := len(_feaa)
	if len(_feaa) < 2 {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	_aff := 0.1
	if _defd == 2 && args[1].Type != ResultTypeEmpty {
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("I\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_aff = args[1].ValueNumber
		if _aff <= -1 {
			return MakeErrorResult("\u0049\u0052R\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0067\u0075\u0065\u0073\u0073\u0020t\u006f\u0020\u0062\u0065\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u002d\u0031")
		}
	}
	_cbbgb := []float64{}
	for _dfbf := 0; _dfbf < _egec; _dfbf++ {
		if _dfbf == 0 {
			_cbbgb = append(_cbbgb, 0)
		} else {
			_cbbgb = append(_cbbgb, _cbbgb[_dfbf-1]+365)
		}
	}
	return _fdb(_feaa, _cbbgb, _aff)
}

// Replace is an implementation of the Excel REPLACE().
func Replace(args []Result) Result {
	_dggbb, _dcbff := _ebee("\u0052E\u0050\u004c\u0041\u0043\u0045", args)
	if _dcbff.Type != ResultTypeEmpty {
		return _dcbff
	}
	_cegce := _dggbb._bcae
	_accc := _dggbb._agabb
	_agaae := _dggbb._fffd
	_bcee := _dggbb._cedf
	_ggfe := len(_cegce)
	if _accc > _ggfe {
		_accc = _ggfe
	}
	_efea := _accc + _agaae
	if _efea > _ggfe {
		_efea = _ggfe
	}
	_afbe := _cegce[0:_accc] + _bcee + _cegce[_efea:]
	return MakeStringResult(_afbe)
}

var _agaa = []ri{{1000, "\u004d"}, {999, "\u0049\u004d"}, {995, "\u0056\u004d"}, {990, "\u0058\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {499, "\u0049\u0044"}, {495, "\u0056\u0044"}, {490, "\u0058\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {99, "\u0049\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

// RegisterFunctionComplex registers a standard function.
func RegisterFunctionComplex(name string, fn FunctionComplex) {
	_cfbcf.Lock()
	defer _cfbcf.Unlock()
	if _, _ceacge := _gbbbb[name]; _ceacge {
		_bc.Log("\u0064\u0075p\u006c\u0069\u0063\u0061t\u0065\u0020r\u0065\u0067\u0069\u0073\u0074\u0072\u0061\u0074i\u006f\u006e\u0020\u006f\u0066\u0020\u0066\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0025\u0073", name)
	}
	_gbbbb[name] = fn
}

var _dgef = [...]uint8{0, 17, 33, 49, 63, 78, 93, 108}

// PrefixVerticalRange is a range expression that when evaluated returns a list of Results from references like Sheet1!AA:IJ (all cells from columns AA to IJ of sheet 'Sheet1').
type PrefixVerticalRange struct {
	_cbfc        Expression
	_dcdb, _cbcf string
}

func _bcg(_df Result) bool {
	if _df.Type == ResultTypeString {
		return _df.ValueString == ""
	}
	return _df.ValueNumber == 0
}
func _cae() {
	_ceb["\u006d\u006d\u002f\u0064\u0064\u002f\u0079\u0079"] = _ge.MustCompile("\u005e" + _afa + _ffb)
	_ceb["\u006dm\u0020\u0064\u0064\u002c\u0020\u0079y"] = _ge.MustCompile("\u005e" + _facd + _ffb)
	_ceb["\u0079\u0079\u002d\u006d\u006d\u002d\u0064\u0064"] = _ge.MustCompile("\u005e" + _dgb + _ffb)
	_ceb["y\u0079\u002d\u006d\u006d\u0053\u0074\u0072\u002d\u0064\u0064"] = _ge.MustCompile("\u005e" + _fg + _ffb)
	_cfba["\u0068\u0068"] = _ge.MustCompile(_acea + _fdc + "\u0024")
	_cfba["\u0068\u0068\u003am\u006d"] = _ge.MustCompile(_acea + _faf + "\u0024")
	_cfba["\u006d\u006d\u003as\u0073"] = _ge.MustCompile(_acea + _ebg + "\u0024")
	_cfba["\u0068\u0068\u003a\u006d\u006d\u003a\u0073\u0073"] = _ge.MustCompile(_acea + _fdcg + "\u0024")
	_dfg = []*_ge.Regexp{_ge.MustCompile("\u005e" + _afa + "\u0024"), _ge.MustCompile("\u005e" + _facd + "\u0024"), _ge.MustCompile("\u005e" + _dgb + "\u0024"), _ge.MustCompile("\u005e" + _fg + "\u0024")}
	_bea = []*_ge.Regexp{_ge.MustCompile("\u005e" + _fdc + "\u0024"), _ge.MustCompile("\u005e" + _faf + "\u0024"), _ge.MustCompile("\u005e" + _ebg + "\u0024"), _ge.MustCompile("\u005e" + _fdcg + "\u0024")}
}

// String returns a string representation of String.
func (_ggafg String) String() string { return "\u0022" + _ggafg._adfgc + "\u0022" }

// SumProduct is an implementation of the Excel SUMPRODUCT() function.
func SumProduct(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0053\u0055\u004d\u0050\u0052\u004f\u0044U\u0043\u0054\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ddcfg := args[0].Type
	for _, _deda := range args {
		if _deda.Type != _ddcfg {
			return MakeErrorResult("\u0053\u0055M\u0050\u0052\u004f\u0044\u0055C\u0054\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u006c\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0074\u0079\u0070\u0065")
		}
	}
	switch _ddcfg {
	case ResultTypeNumber:
		return Product(args)
	case ResultTypeList, ResultTypeArray:
		_efcdc := len(args[0].ListValues())
		_acgd := make([]float64, _efcdc)
		for _acbgc := range _acgd {
			_acgd[_acbgc] = 1.0
		}
		for _, _fgegc := range args {
			if len(_fgegc.ListValues()) != _efcdc {
				return MakeErrorResult("\u0053\u0055\u004d\u0050\u0052\u004f\u0044\u0055\u0043\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069re\u0073 \u0061\u006c\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074s\u0020\u0074\u006f\u0020\u0068\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065 \u0064\u0069\u006d\u0065\u006e\u0073\u0069\u006f\u006e")
			}
			for _ddge, _ebff := range _fgegc.ListValues() {
				_ebff = _ebff.AsNumber()
				if _ebff.Type != ResultTypeNumber {
					return MakeErrorResult("\u0053\u0055\u004d\u0050\u0052\u004fD\u0055\u0043\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u006c\u006c\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020n\u0075m\u0065\u0072\u0069\u0063")
				}
				_acgd[_ddge] = _acgd[_ddge] * _ebff.ValueNumber
			}
		}
		_agda := 0.0
		for _, _edggb := range _acgd {
			_agda += _edggb
		}
		return MakeNumberResult(_agda)
	}
	return MakeNumberResult(1.0)
}
func _fgg(_ecf int) int {
	if _ecf < 1900 {
		if _ecf < 30 {
			_ecf += 2000
		} else {
			_ecf += 1900
		}
	}
	return _ecf
}

const _bcafg = 57366

// Tbillyield implements the Excel TBILLYIELD function.
func Tbillyield(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("T\u0042\u0049\u004c\u004c\u0059\u0049E\u004c\u0044\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_dffc, _fabc, _dbbb := _gdge(args[0], args[1], "\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044")
	if _dbbb.Type == ResultTypeError {
		return _dbbb
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0064\u0069\u0073\u0063\u006f\u0075n\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ffcg := _fabc - _dffc
	if _ffcg > 365 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004cY\u0049\u0045\u004c\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020m\u0061\u0074\u0075r\u0069\u0074\u0079\u0020t\u006f\u0020\u0062\u0065\u0020\u006eo\u0074\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065\u0020\u0079e\u0061\u0072\u0020\u0061\u0066\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074")
	}
	_fgfb := args[2].ValueNumber
	if _fgfb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020p\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cfffb := (100 - _fgfb) / _fgfb
	_abcdb := 360 / _ffcg
	return MakeNumberResult(_cfffb * _abcdb)
}

// Columns implements the Excel COLUMNS function.
func Columns(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0043\u004fL\u0055\u004d\u004e\u0053\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075me\u006e\u0074")
	}
	_bcaf := args[0]
	if _bcaf.Type != ResultTypeArray && _bcaf.Type != ResultTypeList {
		return MakeErrorResult("\u0043O\u004c\u0055M\u004e\u0053\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0066\u0069\u0072\u0073\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_bdga := _bcaf.ValueArray
	if len(_bdga) == 0 {
		return MakeErrorResult("\u0043\u004f\u004c\u0055\u004d\u004e\u0053\u0020r\u0065\u0071\u0075ir\u0065\u0073\u0020\u0061\u0072\u0072a\u0079\u0020\u0074\u006f\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020\u0031\u0020r\u006f\u0077")
	}
	return MakeNumberResult(float64(len(_bdga[0])))
}

// Rows implements the Excel ROWS function.
func Rows(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0052\u004f\u0057\u0053\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074")
	}
	_bebf := args[0]
	if _bebf.Type != ResultTypeArray && _bebf.Type != ResultTypeList {
		return MakeErrorResult("\u0052\u004f\u0057S\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074y\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_bfde := _bebf.ValueArray
	if len(_bfde) == 0 {
		return MakeErrorResult("\u0052O\u0057\u0053 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0074\u006f\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074\u0020\u006ce\u0061\u0073\u0074\u0020\u0031\u0020\u0072\u006f\u0077")
	}
	return MakeNumberResult(float64(len(_bfde)))
}

// TimeValue is an implementation of the Excel TIMEVALUE() function.
func TimeValue(args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeString {
		return MakeErrorResult("\u0054I\u004d\u0045V\u0041\u004c\u0055\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069ng\u006c\u0065\u0020s\u0074\u0072i\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_gbdc := _dda.ToLower(args[0].ValueString)
	if !_beca(_gbdc) {
		_, _, _, _ccc, _ggb := _gaf(_gbdc)
		if _ggb.Type == ResultTypeError {
			_ggb.ErrorMessage = "\u0049\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020f\u006f\u0072\u0020\u0054\u0049\u004d\u0045V\u0041\u004c\u0055\u0045"
			return _ggb
		}
		if _ccc {
			return MakeNumberResult(0)
		}
	}
	_agf, _bega, _abbg, _ceae, _, _cdg := _eafd(_gbdc)
	if _cdg.Type == ResultTypeError {
		return _cdg
	}
	_cagb := _aabf(float64(_agf), float64(_bega), _abbg)
	if _ceae {
		_cagb += 0.5
	} else if _cagb >= 1 {
		_cagb -= float64(int(_cagb))
	}
	return MakeNumberResult(_cagb)
}

// NewError constructs a new error expression from a string.
func NewError(v string) Expression { return Error{_gbd: v} }

// Update returns the same object as updating sheet references does not affect EmptyExpr.
func (_gaef EmptyExpr) Update(q *_a.UpdateQuery) Expression { return _gaef }

var _ebcaag, _dbce, _faac, _caffa, _fefd, _fefaa, _cbgd, _cdfb, _ddcb, _bdbfg, _egb, _fefc, _eefed, _cdadb, _bacfb *_ge.Regexp

// Match implements the MATCH function.
func Match(args []Result) Result {
	_ggbd := len(args)
	if _ggbd != 2 && _ggbd != 3 {
		return MakeErrorResult("\u004d\u0041T\u0043\u0048\u0020\u0072e\u0071\u0075i\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020o\u0072\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_fccced := 1
	if _ggbd == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("\u004d\u0041\u0054\u0043\u0048\u0020\u0072\u0065q\u0075\u0069\u0072es\u0020\u0074\u0068\u0065\u0020\u0074h\u0069\u0072\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006db\u0065\u0072")
		}
		_ebdc := args[2].ValueNumber
		if _ebdc == -1 || _ebdc == 0 {
			_fccced = int(_ebdc)
		}
	}
	_gggf := args[1]
	var _gdfea []Result
	switch _gggf.Type {
	case ResultTypeList:
		_gdfea = _gggf.ValueList
	case ResultTypeArray:
		_aafdd := _gggf.ValueArray
		for _, _fedfb := range _aafdd {
			if len(_fedfb) != 1 {
				return MakeErrorResult("\u004d\u0041\u0054\u0043\u0048\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068e\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006f\u006e\u0065\u002dd\u0069\u006d\u0065\u006e\u0073\u0069o\u006e\u0061l\u0020\u0072a\u006eg\u0065")
			}
			_gdfea = append(_gdfea, _fedfb[0])
		}
	default:
		return MakeErrorResult("\u004d\u0041\u0054\u0043\u0048\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068e\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006f\u006e\u0065\u002dd\u0069\u006d\u0065\u006e\u0073\u0069o\u006e\u0061l\u0020\u0072a\u006eg\u0065")
	}
	_ccfaf := _fbea(args[0])
	switch _fccced {
	case 0:
		for _adcca, _ebef := range _gdfea {
			if _bggf(_ebef, _ccfaf) {
				return MakeNumberResult(float64(_adcca + 1))
			}
		}
	case -1:
		for _bfdba := 0; _bfdba < len(_gdfea); _bfdba++ {
			if _bggf(_gdfea[_bfdba], _ccfaf) {
				return MakeNumberResult(float64(_bfdba + 1))
			}
			if _ccfaf._ddba && (_gdfea[_bfdba].ValueNumber < _ccfaf._gdfd) {
				if _bfdba == 0 {
					return MakeErrorResultType(ErrorTypeNA, "")
				}
				return MakeNumberResult(float64(_bfdba))
			}
		}
	case 1:
		for _gdfa := 0; _gdfa < len(_gdfea); _gdfa++ {
			if _bggf(_gdfea[_gdfa], _ccfaf) {
				return MakeNumberResult(float64(_gdfa + 1))
			}
			if _ccfaf._ddba && (_gdfea[_gdfa].ValueNumber > _ccfaf._gdfd) {
				if _gdfa == 0 {
					return MakeErrorResultType(ErrorTypeNA, "")
				}
				return MakeNumberResult(float64(_gdfa))
			}
		}
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}
func _fae(_da BinOpType, _ea, _gd []Result) Result {
	_bdf := []Result{}
	for _ceee := range _ea {
		_cf := _ea[_ceee].AsNumber()
		_cbb := _gd[_ceee].AsNumber()
		if _cf.Type != ResultTypeNumber || _cbb.Type != ResultTypeNumber {
			return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006fp\u0065\u0072a\u0074\u0069\u006f\u006e")
		}
		switch _da {
		case BinOpTypePlus:
			_bdf = append(_bdf, MakeNumberResult(_cf.ValueNumber+_cbb.ValueNumber))
		case BinOpTypeMinus:
			_bdf = append(_bdf, MakeNumberResult(_cf.ValueNumber-_cbb.ValueNumber))
		case BinOpTypeMult:
			_bdf = append(_bdf, MakeNumberResult(_cf.ValueNumber*_cbb.ValueNumber))
		case BinOpTypeDiv:
			if _cbb.ValueNumber == 0 {
				return MakeErrorResultType(ErrorTypeDivideByZero, "")
			}
			_bdf = append(_bdf, MakeNumberResult(_cf.ValueNumber/_cbb.ValueNumber))
		case BinOpTypeExp:
			_bdf = append(_bdf, MakeNumberResult(_cc.Pow(_cf.ValueNumber, _cbb.ValueNumber)))
		case BinOpTypeLT:
			_bdf = append(_bdf, MakeBoolResult(_cf.ValueNumber < _cbb.ValueNumber))
		case BinOpTypeGT:
			_bdf = append(_bdf, MakeBoolResult(_cf.ValueNumber > _cbb.ValueNumber))
		case BinOpTypeEQ:
			_bdf = append(_bdf, MakeBoolResult(_cf.ValueNumber == _cbb.ValueNumber))
		case BinOpTypeLEQ:
			_bdf = append(_bdf, MakeBoolResult(_cf.ValueNumber <= _cbb.ValueNumber))
		case BinOpTypeGEQ:
			_bdf = append(_bdf, MakeBoolResult(_cf.ValueNumber >= _cbb.ValueNumber))
		case BinOpTypeNE:
			_bdf = append(_bdf, MakeBoolResult(_cf.ValueNumber != _cbb.ValueNumber))
		default:
			return MakeErrorResult(_d.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006c\u0069\u0073\u0074\u0020\u0062\u0069\u006e\u0061\u0072\u0079\u0020\u006fp \u0025\u0073", _da))
		}
	}
	return MakeListResult(_bdf)
}
func _gabe(_cadd float64) bool { return _cadd == 1 || _cadd == 2 || _cadd == 4 }

// ISODD is an implementation of the Excel ISODD() function.
func IsOdd(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u004f\u0044\u0044\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u004f\u0044\u0044\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061 \u006eu\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fabg := int(args[0].ValueNumber)
	return MakeBoolResult(_fabg != _fabg/2*2)
}

// Cumprinc implements the Excel CUMPRINC function.
func Cumprinc(args []Result) Result {
	_dede, _bdgf := _ebca(args, "\u0043\u0055\u004d\u0050\u0052\u0049\u004e\u0043")
	if _bdgf.Type == ResultTypeError {
		return _bdgf
	}
	_gecc := _dede._egab
	_befe := _dede._cefc
	_eee := _dede._dfaf
	_gfeg := _dede._gfdb
	_ddbb := _dede._ggcf
	_dfae := _dede._fgcf
	_aeeag := _acad(_gecc, _befe, _eee, 0, _dfae)
	_aede := 0.0
	if _gfeg == 1 {
		if _dfae == 0 {
			_aede = _aeeag + _eee*_gecc
		} else {
			_aede = _aeeag
		}
		_gfeg++
	}
	for _dec := _gfeg; _dec <= _ddbb; _dec++ {
		if _dfae == 1 {
			_aede += _aeeag - (_gfec(_gecc, _dec-2, _aeeag, _eee, 1)-_aeeag)*_gecc
		} else {
			_aede += _aeeag - _gfec(_gecc, _dec-1, _aeeag, _eee, 0)*_gecc
		}
	}
	return MakeNumberResult(_aede)
}

const (
	ResultTypeUnknown ResultType = iota
	ResultTypeNumber
	ResultTypeString
	ResultTypeList
	ResultTypeArray
	ResultTypeError
	ResultTypeEmpty
)

func (_dbad *Lexer) nextRaw() *node {
	for len(_dbad._bbddc) != 0 {
		_eeef := <-_dbad._bbddc[len(_dbad._bbddc)-1]
		if _eeef != nil {
			return _eeef
		}
		_dbad._bbddc = _dbad._bbddc[0 : len(_dbad._bbddc)-1]
	}
	return <-_dbad._acbfe
}

// Update returns the same object as updating sheet references does not affect Error.
func (_fdf Error) Update(q *_a.UpdateQuery) Expression { return _fdf }

const _caddf = 57371
const _gceebf = 16

func (_gefgf Result) AsString() Result {
	switch _gefgf.Type {
	case ResultTypeNumber:
		return MakeStringResult(_gefgf.Value())
	default:
		return _gefgf
	}
}

// PrefixRangeExpr is a range expression that when evaluated returns a list of Results from a given sheet like Sheet1!A1:B4 (all cells from A1 to B4 from a sheet 'Sheet1').
type PrefixRangeExpr struct{ _afgf, _cbce, _dbdbg Expression }

const _efaeb = 187

var _eafdd string = string([]byte{92})

// Pduration implements the Excel PDURATION function.
func Pduration(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0050\u0044\u0055RA\u0054\u0049\u004f\u004e\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u0074h\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050D\u0055\u0052A\u0054\u0049\u004fN\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_beaf := args[0].ValueNumber
	if _beaf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020p\u006f\u0073i\u0074\u0069\u0076\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u0020\u0076\u0061l\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_caee := args[1].ValueNumber
	if _caee <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "P\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 c\u0075\u0072\u0072\u0065n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074o \u0062\u0065 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0044\u0055\u0052\u0041\u0054I\u004f\u004e\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0065\u0064\u0020\u0076\u0061lu\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ccfg := args[2].ValueNumber
	if _ccfg <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0044\u0055\u0052\u0041\u0054I\u004f\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070o\u0073i\u0074\u0069\u0076\u0065")
	}
	return MakeNumberResult((_cc.Log10(_ccfg) - _cc.Log10(_caee)) / _cc.Log10(1+_beaf))
}
func _bbbac(_dfa int) bool { return _dfa == 0 || _dfa == 4 }

// DateValue is an implementation of the Excel DATEVALUE() function.
func DateValue(args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeString {
		return MakeErrorResult("\u0044A\u0054\u0045V\u0041\u004c\u0055\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069ng\u006c\u0065\u0020s\u0074\u0072i\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_fde := _dda.ToLower(args[0].ValueString)
	if !_gedc(_fde) {
		_, _, _, _, _aaeg, _eaeg := _eafd(_fde)
		if _eaeg.Type == ResultTypeError {
			_eaeg.ErrorMessage = "\u0049\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020f\u006f\u0072\u0020\u0044\u0041\u0054\u0045V\u0041\u004c\u0055\u0045"
			return _eaeg
		}
		if _aaeg {
			return MakeNumberResult(0)
		}
	}
	_cfdf, _ecc, _bfg, _, _gbg := _gaf(_fde)
	if _gbg.Type == ResultTypeError {
		return _gbg
	}
	return MakeNumberResult(_bad(_cfdf, _ecc, _bfg))
}

// Amordegrc implements the Excel AMORDEGRC function.
func Amordegrc(args []Result) Result {
	_aeeg, _abbgg := _bce(args, "\u0041M\u004f\u0052\u0044\u0045\u0047\u0052C")
	if _abbgg.Type == ResultTypeError {
		return _abbgg
	}
	_fceeb := _aeeg._adac
	_feaf := _aeeg._abda
	_fccd := _aeeg._cde
	_cfbag := _aeeg._cfg
	_abde := _aeeg._gcde
	_bcgb := _aeeg._bgac
	if _bcgb >= 0.5 {
		return MakeErrorResultType(ErrorTypeNum, "\u0041\u004d\u004f\u0052\u0044\u0045\u0047R\u0043\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u002e\u0035")
	}
	_babg := _aeeg._ccg
	_gfca := 1.0 / _bcgb
	_ded := 2.5
	if _gfca < 3 {
		_ded = 1
	} else if _gfca < 5 {
		_ded = 1.5
	} else if _gfca <= 6 {
		_ded = 2
	}
	_bcgb *= _ded
	_gcba, _ggc := _agef(_feaf, _fccd, _babg)
	if _ggc.Type == ResultTypeError {
		return MakeErrorResult("\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0072\u0020\u0041\u004d\u004f\u0052\u0044EG\u0052\u0043")
	}
	_daag := _ecac(_gcba * _bcgb * _fceeb)
	_fceeb -= _daag
	_bfd := _fceeb - _cfbag
	for _eede := 0; _eede < _abde; _eede++ {
		_daag = _ecac(_bcgb * _fceeb)
		_bfd -= _daag
		if _bfd < 0 {
			switch _abde - _eede {
			case 0:
			case 1:
				return MakeNumberResult(_ecac(_fceeb * 0.5))
			default:
				return MakeNumberResult(0)
			}
		}
		_fceeb -= _daag
	}
	return MakeNumberResult(_daag)
}

var _ceb = map[string]*_ge.Regexp{}

// Days is an implementation of the Excel DAYS() function.
func Days(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("D\u0041\u0059\u0053\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	var _cag, _ed float64
	switch args[0].Type {
	case ResultTypeNumber:
		_ed = args[0].ValueNumber
	case ResultTypeString:
		_ggaa := DateValue([]Result{args[0]})
		if _ggaa.Type == ResultTypeError {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0065\u006e\u0064\u0020\u0064\u0061\u0074e\u0020\u0066\u006fr\u0020D\u0041\u0059\u0053")
		}
		_ed = _ggaa.ValueNumber
	default:
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0066\u006fr\u0020D\u0041\u0059\u0053")
	}
	switch args[1].Type {
	case ResultTypeNumber:
		_cag = args[1].ValueNumber
		if _cag < 62 && _ed >= 62 {
			_cag--
		}
	case ResultTypeString:
		_ddff := DateValue([]Result{args[1]})
		if _ddff.Type == ResultTypeError {
			return MakeErrorResult("\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u0020DA\u0059\u0053")
		}
		_cag = _ddff.ValueNumber
	default:
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0066\u006fr\u0020D\u0041\u0059\u0053")
	}
	_cecc := float64(int(_ed - _cag))
	return MakeNumberResult(_cecc)
}
func _afgc(_fbgc Result) [][]Result {
	switch _fbgc.Type {
	case ResultTypeArray:
		return _fbgc.ValueArray
	case ResultTypeList:
		return [][]Result{_fbgc.ValueList}
	default:
		return [][]Result{}
	}
}

// Duration implements the Excel DURATION function.
func Duration(args []Result) Result {
	_gfba, _aeda := _cedgd(args, "\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e")
	if _aeda.Type == ResultTypeError {
		return _aeda
	}
	_abga := _gfba._bfge
	_aecc := _gfba._ddbc
	_fba := _gfba._fbgfe
	_fcbb := _gfba._gda
	_aafb := _gfba._gfcf
	_cfga := _gfba._abff
	return _dcgec(_abga, _aecc, _fba, _fcbb, _aafb, _cfga)
}

// Exact is an implementation of the Excel EXACT() which compares two strings.
func Exact(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0043\u004f\u004e\u0043\u0041\u0054\u0045N\u0041\u0054\u0045(\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_bagff := args[0].AsString()
	_ebcd := args[1].AsString()
	if _bagff.Type != ResultTypeString || _ebcd.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u004f\u004e\u0043\u0041\u0054\u0045N\u0041\u0054\u0045(\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	return MakeBoolResult(_bagff.ValueString == _ebcd.ValueString)
}

var _bfeba = false

// Update returns the same object as updating sheet references does not affect Bool.
func (_gef Bool) Update(q *_a.UpdateQuery) Expression { return _gef }

const _ffb = "\u0028\u0020\u0028" + _fdc + "\u007c" + _faf + "\u007c" + _ebg + "\u007c" + _fdcg + "\u0029\u0029\u003f\u0024"
const _dagfb = 57347

func _ccfgf(_gbef Context, _aeaae, _gefac string) (string, string) {
	_dbfcd := _aeaae + "\u0031"
	_ceff := _gbef.LastRow(_aeaae)
	_gbga := _gefac + _c.Itoa(_ceff)
	return _dbfcd, _gbga
}

// Round is an implementation of the Excel ROUND function that rounds a number
// to a specified number of digits.
func Round(args []Result) Result { return _ceacf(args, _fceda) }
func _bggea(_ecdg []Result, _bcbaf countMode) float64 {
	_ddbbe := 0.0
	for _, _abad := range _ecdg {
		switch _abad.Type {
		case ResultTypeNumber:
			if _bcbaf == _faeea || (_bcbaf == _fedg && !_abad.IsBoolean) {
				_ddbbe++
			}
		case ResultTypeList, ResultTypeArray:
			_ddbbe += _bggea(_abad.ListValues(), _bcbaf)
		case ResultTypeString:
			if _bcbaf == _faeea {
				_ddbbe++
			}
		case ResultTypeEmpty:
			if _bcbaf == _dcbfd {
				_ddbbe++
			}
		}
	}
	return _ddbbe
}

const (
	_fbagf cmpResult = 0
	_geagf cmpResult = -1
	_fgdd  cmpResult = 1
	_eadc  cmpResult = 2
)

// NewCellRef constructs a new cell reference.
func NewCellRef(v string) Expression { return CellRef{_ebd: v} }
func init() {
	_cae()
	RegisterFunction("\u0044\u0041\u0054\u0045", Date)
	RegisterFunction("\u0044A\u0054\u0045\u0044\u0049\u0046", DateDif)
	RegisterFunction("\u0044A\u0054\u0045\u0056\u0041\u004c\u0055E", DateValue)
	RegisterFunction("\u0044\u0041\u0059", Day)
	RegisterFunction("\u0044\u0041\u0059\u0053", Days)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0044\u0041\u0059\u0053", Days)
	RegisterFunction("\u0045\u0044\u0041T\u0045", Edate)
	RegisterFunction("\u0045O\u004d\u004f\u004e\u0054\u0048", Eomonth)
	RegisterFunction("\u004d\u0049\u004e\u0055\u0054\u0045", Minute)
	RegisterFunction("\u004d\u004f\u004eT\u0048", Month)
	RegisterFunction("\u004e\u004f\u0057", Now)
	RegisterFunction("\u0054\u0049\u004d\u0045", Time)
	RegisterFunction("\u0054I\u004d\u0045\u0056\u0041\u004c\u0055E", TimeValue)
	RegisterFunction("\u0054\u004f\u0044A\u0059", Today)
	RegisterFunctionComplex("\u0059\u0045\u0041\u0052", Year)
	RegisterFunction("\u0059\u0045\u0041\u0052\u0046\u0052\u0041\u0043", YearFrac)
}

// String returns a string representation of a range with prefix.
func (_cdccc PrefixRangeExpr) String() string {
	return _d.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _cdccc._afgf.String(), _cdccc._cbce.String(), _cdccc._dbdbg.String())
}
func _ddfg(_degb Result) Result {
	if _degb.Type == ResultTypeEmpty {
		return _degb
	}
	_bfcgg := _degb.AsString()
	if _bfcgg.Type != ResultTypeString {
		return MakeErrorResult("\u004c\u004f\u0057\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	if _degb.IsBoolean {
		if _bfcgg.ValueString == "\u0031" {
			return MakeStringResult("\u0074\u0072\u0075\u0065")
		} else if _bfcgg.ValueString == "\u0030" {
			return MakeStringResult("\u0066\u0061\u006cs\u0065")
		} else {
			return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u004c\u004fW\u0045\u0052")
		}
	} else {
		return MakeStringResult(_dda.ToLower(_bfcgg.ValueString))
	}
}

// NewBinaryExpr constructs a new binary expression with a given operator.
func NewBinaryExpr(lhs Expression, op BinOpType, rhs Expression) Expression {
	return BinaryExpr{_eb: lhs, _bd: rhs, _cb: op}
}
func (_gced VerticalRange) verticalRangeReference() string {
	return _d.Sprintf("\u0025\u0073\u003a%\u0073", _gced._cacbdb, _gced._cddc)
}

var _bggaf = [...]uint8{0, 20, 37, 60, 78, 96}

// Index implements the Excel INDEX function.
func Index(args []Result) Result {
	_fdgf := len(args)
	if _fdgf < 2 || _fdgf > 3 {
		return MakeErrorResult("\u0049\u004e\u0044E\u0058\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u006f\u006e\u0065\u0020\u0074\u006f\u0020\u0074\u0068\u0072\u0065\u0065\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_bbffe := args[0]
	if _bbffe.Type != ResultTypeArray && _bbffe.Type != ResultTypeList {
		return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0066i\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_fggaf := args[1].AsNumber()
	if _fggaf.Type != ResultTypeNumber {
		return MakeErrorResult("I\u004e\u0044\u0045\u0058\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u006e\u0075\u006d\u0065\u0072\u0069\u0063 \u0072\u006f\u0077\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_acbg := int(_fggaf.ValueNumber) - 1
	_geddd := -1
	if _fdgf == 3 && args[2].Type != ResultTypeEmpty {
		_egga := args[2].AsNumber()
		if _egga.Type != ResultTypeNumber {
			return MakeErrorResult("I\u004e\u0044\u0045\u0058\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u006e\u0075\u006d\u0065\u0072\u0069\u0063 \u0063\u006f\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
		}
		_geddd = int(_egga.ValueNumber) - 1
	}
	if _acbg == -1 && _geddd == -1 {
		return MakeErrorResult("\u0049\u004e\u0044EX\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0072o\u0077 \u006fr\u0020\u0063\u006f\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	var _eebg []Result
	if _bbffe.Type == ResultTypeArray {
		_fgge := _bbffe.ValueArray
		if _acbg < -1 || _acbg >= len(_fgge) {
			return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0072o\u0077\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
		}
		if _acbg == -1 {
			if _geddd >= len(_fgge[0]) {
				return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0063o\u006c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
			}
			_cabd := [][]Result{}
			for _, _gdga := range _fgge {
				_bgdb := _gdga[_geddd]
				if _bgdb.Type == ResultTypeEmpty {
					_bgdb = MakeNumberResult(0)
				}
				_cabd = append(_cabd, []Result{_bgdb})
			}
			return MakeArrayResult(_cabd)
		}
		_eebg = _fgge[_acbg]
	} else {
		_ddbg := _bbffe.ValueList
		if _acbg < -1 || _acbg >= 1 {
			return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0072o\u0077\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
		}
		if _acbg == -1 {
			if _geddd >= len(_ddbg) {
				return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0063o\u006c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
			}
			_ecbc := _ddbg[_geddd]
			if _ecbc.Type == ResultTypeEmpty {
				_ecbc = MakeNumberResult(0)
			}
			return _ecbc
		}
		_eebg = _ddbg
	}
	if _geddd < -1 || _geddd > len(_eebg) {
		return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0063o\u006c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
	}
	if _geddd == -1 {
		_aafa := []Result{}
		for _, _aacg := range _eebg {
			if _aacg.Type == ResultTypeEmpty {
				_aafa = append(_aafa, MakeNumberResult(0))
			} else {
				_aafa = append(_aafa, _aacg)
			}
		}
		return MakeArrayResult([][]Result{_aafa})
	}
	_dcd := _eebg[_geddd]
	if _dcd.Type == ResultTypeEmpty {
		return MakeNumberResult(0)
	}
	return _dcd
}

// Len is an implementation of the Excel LEN function that returns length of a string
func Len(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004c\u0045N\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_aaef := args[0].AsString()
	if _aaef.Type != ResultTypeString {
		return MakeErrorResult("\u004c\u0045N\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(float64(len(_aaef.ValueString)))
}

// BinOpType is the binary operation operator type
//go:generate stringer -type=BinOpType
type BinOpType byte

// Reference returns a string reference value to a sheet.
func (_feggf SheetPrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeSheet, Value: _feggf._daeae}
}
func _bad(_efa, _eaf, _bca int) float64 {
	return float64(_fgc(_efa, _ee.Month(_eaf), _bca)/86400) + _gab
}

// Ceiling is an implementation of the CEILING function which
// returns the ceiling of a number.
func Ceiling(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("C\u0045\u0049\u004c\u0049\u004e\u0047\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006ee \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u0028\u0029\u0020\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_bfbe := args[0].AsNumber()
	if _bfbe.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066i\u0072\u0073t\u0020\u0061\u0072\u0067u\u006d\u0065\u006et\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049NG\u0028\u0029\u0020m\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020n\u0075\u006db\u0065\u0072")
	}
	_gbdda := float64(1)
	if _bfbe.ValueNumber < 0 {
		_gbdda = -1
	}
	if len(args) > 1 {
		_gbgdd := args[1].AsNumber()
		if _gbgdd.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073e\u0063\u006fn\u0064\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020t\u006f\u0020\u0043\u0045\u0049\u004cI\u004e\u0047\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062e\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_gbdda = _gbgdd.ValueNumber
	}
	if _gbdda < 0 && _bfbe.ValueNumber > 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u006e\u0065\u0067\u0061\u0074\u0069v\u0065\u0020\u0073\u0069\u0067\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004cI\u004e\u0047\u0028\u0029\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Ceil(_bfbe.ValueNumber))
	}
	_ddbe := _bfbe.ValueNumber
	_ddbe, _dcaa := _cc.Modf(_ddbe / _gbdda)
	if _dcaa > 0 {
		_ddbe++
	}
	return MakeNumberResult(_ddbe * _gbdda)
}

// String returns a string representation of FunctionCall expression.
func (_gdfg FunctionCall) String() string {
	_efffg := _gg.Buffer{}
	_efffg.WriteString(_gdfg._ffcfde)
	_efffg.WriteString("\u0028")
	_gaaff := len(_gdfg._gace) - 1
	for _bgfg, _bffg := range _gdfg._gace {
		_efffg.WriteString(_bffg.String())
		if _bgfg != _gaaff {
			_efffg.WriteString("\u002c")
		}
	}
	_efffg.WriteString("\u0029")
	return _efffg.String()
}
func (_facb *Lexer) emit(_acff tokenType, _fegga []byte) {
	if _bfeba {
		_d.Println("\u0065\u006d\u0069\u0074", _acff, _dbcb(string(_fegga)))
	}
	_facb._acbfe <- &node{_acff, string(_fegga)}
}
func _gfcae(_fece, _ddffe, _gdc, _efbc, _befb, _ebdg float64) float64 {
	var _dbc, _cccg float64
	_bgge := 0.0
	_bgca := _cc.Ceil(_befb)
	_eace := _fece - _ddffe
	_egabc := false
	_agcc := 0.0
	for _affg := 1.0; _affg <= _bgca; _affg++ {
		if !_egabc {
			_dbc = _afec(_fece, _ddffe, _gdc, _affg, _ebdg)
			_agcc = _eace / (_gdc - _affg + 1)
			if _agcc > _dbc {
				_cccg = _agcc
				_egabc = true
			} else {
				_cccg = _dbc
				_eace -= _dbc
			}
		} else {
			_cccg = _agcc
		}
		if _affg == _bgca {
			_cccg *= _befb + 1 - _bgca
		}
		_bgge += _cccg
	}
	return _bgge
}
func _efbe(_bcgg, _bfba _ee.Time, _baafg, _eafdc int) _ee.Time {
	_bgdg := _bfba
	_bdbd := _bcgg.Year() - _bfba.Year()
	_bgdg = _bgdg.AddDate(_bdbd, 0, 0)
	if _bcgg.After(_bgdg) {
		_bgdg = _bgdg.AddDate(1, 0, 0)
	}
	_ccce := -12 / _baafg
	for _bgdg.After(_bcgg) {
		_bgdg = _bgdg.AddDate(0, _ccce, 0)
	}
	return _bgdg
}

// Reference returns an invalid reference for Negate.
func (_geede Negate) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Oddlprice implements the Excel ODDLPRICE function.
func Oddlprice(args []Result) Result {
	if len(args) != 8 && len(args) != 9 {
		return MakeErrorResult("\u004f\u0044\u0044L\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u006e\u0069\u006e\u0065\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_bedc, _gffg, _agge := _gdge(args[0], args[1], "\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E")
	if _agge.Type == ResultTypeError {
		return _agge
	}
	_cagec, _agge := _badf(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E")
	if _agge.Type == ResultTypeError {
		return _agge
	}
	if _cagec >= _bedc {
		return MakeErrorResultType(ErrorTypeNum, "\u004c\u0061\u0073\u0074\u0020i\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0064\u0061\u0074\u0065\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0062\u0065\u0066\u006f\u0072\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074e")
	}
	_edfa := args[3]
	if _edfa.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_gegga := _edfa.ValueNumber
	if _gegga < 0 {
		return MakeErrorResultType(ErrorTypeNum, "R\u0061\u0074\u0065\u0020\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006fn\u0020\u006e\u0065g\u0061t\u0069\u0076\u0065")
	}
	_cgdg := args[4]
	if _cgdg.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0079i\u0065\u006c\u0064\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_cfff := _cgdg.ValueNumber
	if _cfff < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	_fdaf := args[5]
	if _fdaf.Type != ResultTypeNumber {
		return MakeErrorResult("\u004fD\u0044\u004cP\u0052\u0049\u0043\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	_ebcaa := _fdaf.ValueNumber
	if _ebcaa < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	_gacfe := args[6]
	if _gacfe.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049C\u0045\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0072\u0065\u0071\u0075\u0065\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_bebbg := float64(int(_gacfe.ValueNumber))
	if !_gabe(_bebbg) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_ccff := 0
	if len(args) == 8 && args[7].Type != ResultTypeEmpty {
		_daeb := args[7]
		if _daeb.Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0062a\u0073\u0069\u0073\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
		}
		_ccff = int(_daeb.ValueNumber)
		if !_gaff(_ccff) {
			return MakeErrorResultType(ErrorTypeNum, "I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0062\u0061\u0073\u0069s\u0020\u0076\u0061\u006c\u0075\u0065\u0020f\u006f\u0072\u0020\u004f\u0044\u0044\u004c\u0050\u0052\u0049C\u0045")
		}
	}
	_fede, _agge := _agef(_cagec, _gffg, _ccff)
	if _agge.Type == ResultTypeError {
		return _agge
	}
	_fede *= _bebbg
	_gdac, _agge := _agef(_bedc, _gffg, _ccff)
	if _agge.Type == ResultTypeError {
		return _agge
	}
	_gdac *= _bebbg
	_fbed, _agge := _agef(_cagec, _bedc, _ccff)
	if _agge.Type == ResultTypeError {
		return _agge
	}
	_fbed *= _bebbg
	_efcd := _ebcaa + _fede*100*_gegga/_bebbg
	_efcd /= _gdac*_cfff/_bebbg + 1
	_efcd -= _fbed * 100 * _gegga / _bebbg
	return MakeNumberResult(_efcd)
}
func (_aedd *ivr) Sheet(name string) Context { return _aedd }
func _fgba(_ccagd Reference, _dagd Context) bool {
	return _dagd.Sheet(_ccagd.Value) == InvalidReferenceContext
}

// CeilingPrecise is an implementation of the CEILING.PRECISE function which
// returns the ceiling of a number.
func CeilingPrecise(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002eP\u0052\u0045\u0043IS\u0045\u0028\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020o\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0043\u0045I\u004c\u0049\u004e\u0047\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_acfe := args[0].AsNumber()
	if _acfe.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069r\u0073\u0074\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_bbac := float64(1)
	if _acfe.ValueNumber < 0 {
		_bbac = -1
	}
	if len(args) > 1 {
		_cabg := args[1].AsNumber()
		if _cabg.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0043E\u0049L\u0049\u004e\u0047\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_bbac = _cc.Abs(_cabg.ValueNumber)
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Ceil(_acfe.ValueNumber))
	}
	_daga := _acfe.ValueNumber
	_daga, _adcccf := _cc.Modf(_daga / _bbac)
	if _adcccf != 0 {
		if _acfe.ValueNumber > 0 {
			_daga++
		}
	}
	return MakeNumberResult(_daga * _bbac)
}

// NewEvaluator constructs a new defEval object which is the default formula evaluator.
func NewEvaluator() Evaluator            { _bgb := &defEval{}; _bgb.evCache = _gba(); return _bgb }
func _afddf(_becc, _dcgeag float64) bool { return _cc.Abs(_becc-_dcgeag) < 1.0e-6 }

// Reference returns a string reference value to a horizontal range with prefix.
func (_egead PrefixHorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	_affb := _egead._fdcfe.Reference(ctx, ev)
	return Reference{Type: ReferenceTypeHorizontalRange, Value: _egead.horizontalRangeReference(_affb.Value)}
}

// Transpose implements the TRANSPOSE function that transposes a cell range.
func Transpose(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0054\u0052AN\u0053\u0050\u004fS\u0045\u0020\u0072\u0065qui\u0072es\u0020\u0061\u0020\u0073\u0069\u006e\u0067le\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[0].Type != ResultTypeArray && args[0].Type != ResultTypeList {
		return MakeErrorResult("T\u0052\u0041\u004e\u0053\u0050\u004fS\u0045\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0061\u0020\u0072a\u006e\u0067\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_eecce := args[0]
	if _eecce.Type == ResultTypeList {
		_feba := [][]Result{}
		for _, _ebaag := range _eecce.ValueList {
			_feba = append(_feba, []Result{_ebaag})
		}
		return MakeArrayResult(_feba)
	}
	_fdcdd := make([][]Result, len(_eecce.ValueArray[0]))
	for _, _bbfdd := range _eecce.ValueArray {
		for _acbd, _fafa := range _bbfdd {
			_fdcdd[_acbd] = append(_fdcdd[_acbd], _fafa)
		}
	}
	return MakeArrayResult(_fdcdd)
}
func NewLexer() *Lexer { return &Lexer{_acbfe: make(chan *node)} }

// Coupdaybs implements the Excel COUPDAYBS function.
func Coupdaybs(args []Result) Result {
	_fggg, _aegf := _dccb(args, "\u0043O\u0055\u0050\u0044\u0041\u0059\u0042S")
	if _aegf.Type == ResultTypeError {
		return _aegf
	}
	return MakeNumberResult(_fffe(_fggg._eac, _fggg._aea, _fggg._dfff, _fggg._cfeb))
}

// SetLocked does nothing for the invalid reference context.
func (_decb *ivr) SetLocked(cellRef string, locked bool) {}
func _abab(_gbf, _fcagd []float64, _bfeef float64) float64 {
	_gfee := _bfeef + 1
	_fecc := 0.0
	_ccedb := len(_gbf)
	_caff := _fcagd[0]
	for _fgec := 1; _fgec < _ccedb; _fgec++ {
		_edda := (_fcagd[_fgec] - _caff) / 365
		_fecc -= _edda * _gbf[_fgec] / _cc.Pow(_gfee, _edda+1)
	}
	return _fecc
}

// Pi is an implementation of the Excel Pi() function that just returns the Pi
// constant.
func Pi(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u0050I\u0028\u0029\u0020\u0061c\u0063\u0065\u0070\u0074\u0073 \u006eo\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074s")
	}
	return MakeNumberResult(_cc.Pi)
}

// LCM implements the Excel LCM() function which returns the least common
// multiple of a range of numbers.
func LCM(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004c\u0043M(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0061t \u006c\u0065\u0061\u0073\u0074\u0020\u006fne\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_bdacb := []float64{}
	for _, _eebe := range args {
		switch _eebe.Type {
		case ResultTypeString:
			_gedb := _eebe.AsNumber()
			if _gedb.Type != ResultTypeNumber {
				return MakeErrorResult("\u004c\u0043M(\u0029\u0020\u006fn\u006c\u0079\u0020\u0061cce\u0070ts\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
			}
			_bdacb = append(_bdacb, _gedb.ValueNumber)
		case ResultTypeList:
			_ggede := LCM(_eebe.ValueList)
			if _ggede.Type != ResultTypeNumber {
				return _ggede
			}
			_bdacb = append(_bdacb, _ggede.ValueNumber)
		case ResultTypeNumber:
			_bdacb = append(_bdacb, _eebe.ValueNumber)
		case ResultTypeEmpty:
		case ResultTypeError:
			return _eebe
		}
	}
	if len(_bdacb) == 0 {
		return MakeErrorResult("\u004cC\u004d\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006fne\u0020\u006e\u006fn\u002d\u0065m\u0070\u0074\u0079\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	if _bdacb[0] < 0 {
		return MakeErrorResult("\u004c\u0043M\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(_bdacb) == 1 {
		return MakeNumberResult(_bdacb[0])
	}
	_cfdfd := _bdacb[0]
	for _dbefe := 1; _dbefe < len(_bdacb); _dbefe++ {
		if _bdacb[_dbefe] < 0 {
			return MakeErrorResult("\u004c\u0043M\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
		}
		_cfdfd = _bgfdd(_cfdfd, _bdacb[_dbefe])
	}
	return MakeNumberResult(_cfdfd)
}

// Reference returns an invalid reference for EmptyExpr.
func (_ace EmptyExpr) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Price implements the Excel PRICE function.
func Price(args []Result) Result {
	_befg := len(args)
	if _befg != 6 && _befg != 7 {
		return MakeErrorResult("\u0050\u0052I\u0043\u0045\u0020\u0072e\u0071\u0075i\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020o\u0072\u0020\u0073\u0065\u0076\u0065\u006e\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_dfbfa, _fafd, _fcdd := _gdge(args[0], args[1], "\u0050\u0052\u0049C\u0045")
	if _fcdd.Type == ResultTypeError {
		return _fcdd
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049CE\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0072a\u0074e\u0020o\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_eaee := args[2].ValueNumber
	if _eaee < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u006eo\u0074\u0020\u0062\u0065\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("P\u0052\u0049\u0043\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u0079\u0069\u0065\u006c\u0064\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
	}
	_fcefg := args[3].ValueNumber
	if _fcefg < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0079\u0069\u0065\u006c\u0064 \u0074\u006f\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("P\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0072\u0065\u0064em\u0070\u0074\u0069\u006fn\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gaed := args[4].ValueNumber
	if _gaed <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073 \u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e \u0074\u006f\u0020\u0062\u0065 p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_egf := args[5]
	if _egf.Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0066\u0072\u0065\u0071\u0075e\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_cgcbc := _egf.ValueNumber
	if !_gabe(_cgcbc) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_ddfa := 0
	if _befg == 7 && args[6].Type != ResultTypeEmpty {
		if args[6].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0052\u0049C\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_ddfa = int(args[6].ValueNumber)
		if !_gaff(_ddfa) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069s\u0020a\u0072g\u0075m\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0050\u0052\u0049\u0043\u0045")
		}
	}
	_acgb, _fcdd := _gaa(_dfbfa, _fafd, _eaee, _fcefg, _gaed, _cgcbc, _ddfa)
	if _fcdd.Type == ResultTypeError {
		return _fcdd
	}
	return MakeNumberResult(_acgb)
}

// NewPrefixExpr constructs an expression with prefix.
func NewPrefixExpr(pfx, exp Expression) Expression { return &PrefixExpr{_cacbf: pfx, _bgagg: exp} }

// Disc implements the Excel DISC function.
func Disc(args []Result) Result {
	_babce := len(args)
	if _babce != 4 && _babce != 5 {
		return MakeErrorResult("\u0044\u0049SC\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s f\u006fur\u0020\u006f\u0072\u0020\u0066\u0069\u0076e \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_eaca, _dfgg, _bacf := _gdge(args[0], args[1], "\u0044\u0049\u0053\u0043")
	if _bacf.Type == ResultTypeError {
		return _bacf
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fcag := args[2].ValueNumber
	if _fcag <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0070\u0072\u0020\u0074o \u0062\u0065\u0020\u0070o\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0049S\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_acfb := args[3].ValueNumber
	if _acfb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0049\u0053\u0043\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0072e\u0064\u0065m\u0070\u0074\u0069\u006f\u006e\u0020t\u006f b\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bcdf := 0
	if _babce == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0044\u0049\u0053\u0043\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_bcdf = int(args[4].ValueNumber)
		if !_gaff(_bcdf) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0066\u006f\u0072 \u0044\u0049\u0053\u0043")
		}
	}
	_caea, _bacf := _agef(_eaca, _dfgg, _bcdf)
	if _bacf.Type == ResultTypeError {
		return _bacf
	}
	return MakeNumberResult((_acfb - _fcag) / _acfb / _caea)
}

// Received implements the Excel RECEIVED function.
func Received(args []Result) Result {
	_caeg := len(args)
	if _caeg != 4 && _caeg != 5 {
		return MakeErrorResult("R\u0045\u0043\u0045\u0049\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066o\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065 a\u0072\u0067\u0075m\u0065n\u0074\u0073")
	}
	_eacac, _bcfb, _becd := _gdge(args[0], args[1], "\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044")
	if _becd.Type == ResultTypeError {
		return _becd
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020i\u006e\u0076\u0065\u0073\u0074\u006d\u0065n\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cdaa := args[2].ValueNumber
	if _cdaa <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0069\u006ev\u0065\u0073\u0074\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044 \u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u0064\u0069s\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_bfbb := args[3].ValueNumber
	if _bfbb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052\u0045\u0043\u0045I\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020p\u006f\u0073\u0069\u0074\u0069v\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_eabfb := 0
	if _caeg == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052E\u0043\u0045I\u0056\u0045\u0044 \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_eabfb = int(args[4].ValueNumber)
		if !_gaff(_eabfb) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0052\u0045\u0043E\u0049\u0056\u0045\u0044")
		}
	}
	_ggda, _becd := _agef(_eacac, _bcfb, _eabfb)
	if _becd.Type == ResultTypeError {
		return _becd
	}
	return MakeNumberResult(_cdaa / (1 - _bfbb*_ggda))
}
func _dcaba(_febf float64, _ceec *criteriaRegex) bool {
	_bcbf, _gdgd := _c.ParseFloat(_ceec._cfcd, 64)
	if _gdgd != nil {
		return false
	}
	switch _ceec._eefb {
	case _deba:
		return _febf == _bcbf
	case _cfcb:
		return _febf <= _bcbf
	case _adffd:
		return _febf >= _bcbf
	case _bbfb:
		return _febf < _bcbf
	case _aafag:
		return _febf > _bcbf
	}
	return false
}

// Xor is an implementation of the Excel XOR() function and takes a variable
// number of arguments. It's odd to say the least.  If any argument is numeric,
// it returns true if the number of non-zero numeric arguments is odd and false
// otherwise.  If no argument is numeric, it returns an error.
func Xor(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0058\u004f\u0052 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061t\u0020l\u0065a\u0073t\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gfeb := 0
	_cfgag := false
	for _, _faag := range args {
		switch _faag.Type {
		case ResultTypeList, ResultTypeArray:
			_ggeg := Xor(_faag.ListValues())
			if _ggeg.Type == ResultTypeError {
				return _ggeg
			}
			if _ggeg.ValueNumber != 0 {
				_gfeb++
			}
			_cfgag = true
		case ResultTypeNumber:
			if _faag.ValueNumber != 0 {
				_gfeb++
			}
			_cfgag = true
		case ResultTypeString:
		case ResultTypeError:
			return _faag
		default:
			return MakeErrorResult("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0061\u0072\u0067u\u006de\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0058\u004f\u0052")
		}
	}
	if !_cfgag {
		return MakeErrorResult("\u0058\u004f\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0069n\u0070\u0075\u0074")
	}
	return MakeBoolResult(_gfeb%2 != 0)
}
func (_gbcgb PrefixHorizontalRange) horizontalRangeReference(_cacg string) string {
	return _d.Sprintf("\u0025\u0073\u0021\u0025\u0064\u003a\u0025\u0064", _cacg, _gbcgb._ffba, _gbcgb._cebd)
}
func _cfcef(_cdbcb Context, _ceea Evaluator, _dceaeg, _gffaa string) Result {
	_edbfb, _cgbb := _bb.ParseCellReference(_dceaeg)
	if _cgbb != nil {
		return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073e\u0020r\u0061n\u0067e\u0020\u0025\u0073\u003a\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _dceaeg, _cgbb.Error()))
	}
	_ffage, _fgeda := _edbfb.ColumnIdx, _edbfb.RowIdx
	_agbd, _accd := _bb.ParseCellReference(_gffaa)
	if _accd != nil {
		return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073e\u0020r\u0061n\u0067e\u0020\u0025\u0073\u003a\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gffaa, _accd.Error()))
	}
	_gceg, _fdedf := _agbd.ColumnIdx, _agbd.RowIdx
	_afdgb := [][]Result{}
	for _edffg := _fgeda; _edffg <= _fdedf; _edffg++ {
		_dbgg := []Result{}
		for _fcecb := _ffage; _fcecb <= _gceg; _fcecb++ {
			_cdada := _cdbcb.Cell(_d.Sprintf("\u0025\u0073\u0025\u0064", _bb.IndexToColumn(_fcecb), _edffg), _ceea)
			_dbgg = append(_dbgg, _cdada)
		}
		_afdgb = append(_afdgb, _dbgg)
	}
	if len(_afdgb) == 1 {
		if len(_afdgb[0]) == 1 {
			return _afdgb[0][0]
		}
		return MakeListResult(_afdgb[0])
	}
	return MakeArrayResult(_afdgb)
}

// Cell is an implementation of the Excel CELL function that returns information
// about the formatting, location, or contents of a cell.
func Cell(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 && len(args) != 2 {
		return MakeErrorResult("\u0043\u0045\u004cL \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020o\u006ee\u0020o\u0072 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_gbba := args[0].AsString()
	if _gbba.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u0045\u004c\u004c\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065 \u0073t\u0072\u0069\u006e\u0067")
	}
	_aeff := "\u0041\u0031"
	if len(args) == 2 {
		_daeed := args[1].Ref
		if _daeed.Type != ReferenceTypeCell {
			return MakeErrorResult("\u0043\u0045\u004c\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064 \u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079p\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		_aeff = _daeed.Value
	}
	switch _gbba.ValueString {
	case "\u0061d\u0064\u0072\u0065\u0073\u0073":
		_bfea, _ecfa := _bb.ParseCellReference(_aeff)
		if _ecfa != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _aeff)
		}
		_dcbd := "\u0024" + _bfea.Column + "\u0024" + _c.Itoa(int(_bfea.RowIdx))
		if _bfea.SheetName != "" {
			_dcbd = _bfea.SheetName + "\u0021" + _dcbd
		}
		return MakeStringResult(_dcbd)
	case "\u0063\u006f\u006c":
		_afadg, _daedg := _bb.ParseCellReference(_aeff)
		if _daedg != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _aeff)
		}
		return MakeNumberResult(float64(_afadg.ColumnIdx + 1))
	case "\u0063\u006f\u006co\u0072":
		_fecdb := _dda.Contains(ctx.GetFormat(_aeff), "\u005b\u0052\u0045D\u005d")
		return MakeBoolResult(_fecdb)
	case "\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073":
		return args[1]
	case "\u0066\u0069\u006c\u0065\u006e\u0061\u006d\u0065":
		return MakeStringResult(ctx.GetFilename())
	case "\u0066\u006f\u0072\u006d\u0061\u0074":
		_eccg := "\u0047"
		_degdd := ctx.GetFormat(_aeff)
		if _degdd == "\u0047e\u006e\u0065\u0072\u0061\u006c" || _ebcaag.MatchString(_degdd) {
			_eccg = "\u0046\u0030"
		} else if _degdd == "\u0030\u0025" {
			_eccg = "\u0050\u0030"
		} else if _degdd == "\u004d\u004d\u004d\u0020\u0044\u0044" {
			_eccg = "\u0044\u0032"
		} else if _degdd == "\u004d\u004d\u002fY\u0059" {
			_eccg = "\u0044\u0033"
		} else if _degdd == "\u004d\u004d\u002f\u0044D/\u0059\u0059\u005c\u0020\u0048\u0048\u003a\u004d\u004d\u005c\u0020\u0041\u004d\u002fP\u004d" || _degdd == "M\u004d/\u0044\u0044\u002f\u0059\u0059\u0059\u0059\u005c \u0048\u0048\u003a\u004dM:\u0053\u0053" {
			_eccg = "\u0044\u0034"
		} else if _degdd == "\u004d\u004d\u005c\u002d\u0044\u0044" {
			_eccg = "\u0044\u0035"
		} else if _degdd == "\u0048H\u003aM\u004d\u003a\u0053\u0053\u005c\u0020\u0041\u004d\u002f\u0050\u004d" {
			_eccg = "\u0044\u0036"
		} else if _degdd == "\u0048\u0048\u003aM\u004d\u005c\u0020\u0041\u004d\u002f\u0050\u004d" {
			_eccg = "\u0044\u0037"
		} else if _degdd == "\u0048\u0048\u003a\u004d\u004d\u003a\u0053\u0053" {
			_eccg = "\u0044\u0038"
		} else if _degdd == "\u0048\u0048\u003aM\u004d" {
			_eccg = "\u0044\u0039"
		} else if _faac.MatchString(_degdd) {
			_eccg = "\u002e\u0030"
		} else if _caffa.MatchString(_degdd) {
			_eccg = "\u002e\u0030\u0028\u0029"
		} else if _bdbfg.MatchString(_degdd) {
			_eccg = "\u0043\u0030"
		} else if _eefed.MatchString(_degdd) || _cdadb.MatchString(_degdd) {
			_eccg = "\u0044\u0031"
		} else if _cbbb := _dbce.FindStringSubmatch(_degdd); len(_cbbb) > 1 {
			_eccg = "\u0046" + _c.Itoa(len(_cbbb[1]))
		} else if _bdaa := _fefaa.FindStringSubmatch(_degdd); len(_bdaa) > 1 {
			_eccg = "\u002e" + _c.Itoa(len(_bdaa[2]))
		} else if _faee := _ddcb.FindStringSubmatch(_degdd); len(_faee) > 1 {
			_eccg = "\u0050" + _c.Itoa(len(_faee[2]))
		} else if _dbbe := _egb.FindStringSubmatch(_degdd); len(_dbbe) > 1 {
			_eccg = "\u0043" + _eeab(_dbbe, 1)
		} else if _dgabb := _fefc.FindStringSubmatch(_degdd); len(_dgabb) > 1 {
			_eccg = "\u0043" + _eeab(_dgabb, 1)
		} else if _acbga := _cbgd.FindStringSubmatch(_degdd); len(_acbga) > 1 {
			_eccg = "\u002e" + _eeab(_acbga, 1) + "\u0028\u0029"
		} else if _gaba := _fefd.FindStringSubmatch(_degdd); len(_gaba) > 1 {
			_eccg = "\u002e" + _eeab(_gaba, 1)
		} else if _bbdf := _bacfb.FindStringSubmatch(_degdd); len(_bbdf) > 1 {
			_eccg = "\u0053" + _eeab(_bbdf, 3)
		}
		if _eccg != "\u0047" && _dda.Contains(_degdd, "\u005b\u0052\u0045D\u005d") {
			_eccg += "\u002d"
		}
		return MakeStringResult(_eccg)
	case "p\u0061\u0072\u0065\u006e\u0074\u0068\u0065\u0073\u0065\u0073":
		_fefad := ctx.GetFormat(_aeff)
		if _cdfb.MatchString(_fefad) {
			return MakeNumberResult(1)
		} else {
			return MakeNumberResult(0)
		}
	case "\u0070\u0072\u0065\u0066\u0069\u0078":
		return MakeStringResult(ctx.GetLabelPrefix(_aeff))
	case "\u0070r\u006f\u0074\u0065\u0063\u0074":
		_baea := 0.0
		if ctx.GetLocked(_aeff) {
			_baea = 1.0
		}
		return MakeNumberResult(_baea)
	case "\u0072\u006f\u0077":
		_fbbb, _cbaa := _bb.ParseCellReference(_aeff)
		if _cbaa != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _aeff)
		}
		return MakeNumberResult(float64(_fbbb.RowIdx))
	case "\u0074\u0079\u0070\u0065":
		switch args[1].Type {
		case ResultTypeEmpty:
			return MakeStringResult("\u0062")
		case ResultTypeString:
			return MakeStringResult("\u006c")
		default:
			return MakeStringResult("\u0076")
		}
	case "\u0077\u0069\u0064t\u0068":
		_cceb, _gfaa := _bb.ParseCellReference(_aeff)
		if _gfaa != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _aeff)
		}
		if _cceb.SheetName == "" {
			return MakeNumberResult(ctx.GetWidth(int(_cceb.ColumnIdx)))
		} else {
			return MakeNumberResult(ctx.Sheet(_cceb.SheetName).GetWidth(int(_cceb.ColumnIdx)))
		}
	}
	return MakeErrorResult("\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0066\u0069\u0072\u0073t\u0020a\u0072g\u0075m\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0043\u0045\u004c\u004c\u003a\u0020" + _gbba.ValueString)
}

type Expression interface {
	Eval(_bgde Context, _bec Evaluator) Result
	Reference(_bcf Context, _cge Evaluator) Reference
	String() string
	Update(_cef *_a.UpdateQuery) Expression
}

const _dbgca = 57352

// Count implements the COUNT function.
func Count(args []Result) Result { return MakeNumberResult(_bggea(args, _fedg)) }

func Average(args []Result) Result {
	_abfg, _fdcdg := _gfcfe(args, false)
	if _fdcdg == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0041\u0056\u0045\u0052AG\u0045\u0020\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0062\u0079\u0020\u007a\u0065r\u006f")
	}
	return MakeNumberResult(_abfg / _fdcdg)
}

// Find is an implementation of the Excel FIND().
func Find(args []Result) Result {
	_aade, _bdcfb := _debf("\u0046\u0049\u004e\u0044", args)
	if _bdcfb.Type != ResultTypeEmpty {
		return _bdcfb
	}
	_bgbf := _aade._fcgea
	if _bgbf == "" {
		return MakeNumberResult(1.0)
	}
	_adcbd := _aade._bfgd
	_gefa := _aade._bdgb
	_facc := 1
	for _ccfdd := range _adcbd {
		if _facc < _gefa {
			_facc++
			continue
		}
		_dfbfg := _dda.Index(_adcbd[_ccfdd:], _bgbf)
		if _dfbfg == 0 {
			return MakeNumberResult(float64(_facc))
		}
		_facc++
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}

const _cffac = "\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054\u0079\u0070\u0065\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054\u0079\u0070\u0065\u0043\u0065\u006c\u006c\u0052\u0065\u0066\u0065r\u0065\u006ec\u0065\u0054\u0079\u0070e\u004e\u0061\u006d\u0065\u0064\u0052\u0061\u006e\u0067\u0065R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054y\u0070\u0065\u0052\u0061\u006e\u0067\u0065\u0052\u0065\u0066e\u0072\u0065\u006ec\u0065\u0054\u0079\u0070\u0065\u0053\u0068e\u0065\u0074"

func _bgbc(_bbgb, _cfad, _aabd, _fecd int) int {
	if _cfad > _aabd {
		return 0
	}
	if _bbbac(_fecd) {
		return (_aabd - _cfad + 1) * 30
	}
	_dfbc := 0
	for _bfc := _cfad; _bfc <= _aabd; _bfc++ {
		_dfbc += _cdf(_bbgb, _bfc)
	}
	return _dfbc
}

var ReferenceInvalid = Reference{Type: ReferenceTypeInvalid}

const (
	ReferenceTypeInvalid ReferenceType = iota
	ReferenceTypeCell
	ReferenceTypeHorizontalRange
	ReferenceTypeVerticalRange
	ReferenceTypeNamedRange
	ReferenceTypeRange
	ReferenceTypeSheet
)

func _acfa(_fagb, _aeg int) int {
	switch _aeg {
	case 1:
		if _bac(_fagb) {
			return 366
		} else {
			return 365
		}
	case 3:
		return 365
	default:
		return 360
	}
}
func _fdfd(_eaeef []Result, _badaa []string, _bcgfb bool) []string {
	for _, _fabgd := range _eaeef {
		switch _fabgd.Type {
		case ResultTypeEmpty:
			if !_bcgfb {
				_badaa = append(_badaa, "")
			}
		case ResultTypeString:
			if _fabgd.ValueString != "" || !_bcgfb {
				_badaa = append(_badaa, _fabgd.ValueString)
			}
		case ResultTypeNumber:
			_badaa = append(_badaa, _fabgd.Value())
		case ResultTypeList:
			_badaa = _cbeg(_badaa, _fdfd(_fabgd.ValueList, []string{}, _bcgfb))
		case ResultTypeArray:
			for _, _gfecd := range _fabgd.ValueArray {
				_badaa = _cbeg(_badaa, _fdfd(_gfecd, []string{}, _bcgfb))
			}
		}
	}
	return _badaa
}
func _gfec(_cefe, _ceeb, _dee, _ceeg float64, _caeb int) float64 {
	var _adfe float64
	if _cefe == 0 {
		_adfe = _ceeg + _dee*_ceeb
	} else {
		_dggae := _cc.Pow(1+_cefe, _ceeb)
		if _caeb == 1 {
			_adfe = _ceeg*_dggae + _dee*(1+_cefe)*(_dggae-1)/_cefe
		} else {
			_adfe = _ceeg*_dggae + _dee*(_dggae-1)/_cefe
		}
	}
	return -_adfe
}

// Column implements the Excel COLUMN function.
func Column(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0043\u004f\u004c\u0055M\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_eeeg := args[0].Ref
	if _eeeg.Type != ReferenceTypeCell {
		return MakeErrorResult("\u0043\u004f\u004c\u0055\u004dN\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u006e\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_fcfe, _ageb := _bb.ParseCellReference(_eeeg.Value)
	if _ageb != nil {
		return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _eeeg.Value)
	}
	return MakeNumberResult(float64(_fcfe.ColumnIdx + 1))
}

// Reference returns an invalid reference for Number.
func (_bcfa Number) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Couppcd implements the Excel COUPPCD function.
func Couppcd(args []Result) Result {
	_fgfe, _dde := _dccb(args, "\u0043O\u0055\u0050\u0050\u0043\u0044")
	if _dde.Type == ResultTypeError {
		return _dde
	}
	_adgca := _dcge(_fgfe._eac)
	_eca := _dcge(_fgfe._aea)
	_cdfe := _fgfe._dfff
	_geecb := _fgfe._cfeb
	_fedc := _efbe(_adgca, _eca, _cdfe, _geecb)
	_cedc, _ddaf, _cgc := _fedc.Date()
	return MakeNumberResult(_bad(_cedc, int(_ddaf), _cgc))
}

// EmptyExpr is an empty expression.
type EmptyExpr struct{}

const _gcgg int = 0

// ISREF is an implementation of the Excel ISREF() function.
func IsRef(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0052\u0045\u0046\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(ev.LastEvalIsRef())
}
func _cbeg(_babgc, _cbgec []string) []string {
	for _, _cbea := range _cbgec {
		_babgc = append(_babgc, _cbea)
	}
	return _babgc
}

// GetFormat returns an empty string for the invalid reference context.
func (_deaaf *ivr) GetFormat(cellRef string) string { return "" }

type countMode byte

func init() {
	RegisterFunction("\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d", Accrintm)
	RegisterFunction("\u0041M\u004f\u0052\u0044\u0045\u0047\u0052C", Amordegrc)
	RegisterFunction("\u0041\u004d\u004f\u0052\u004c\u0049\u004e\u0043", Amorlinc)
	RegisterFunction("\u0043O\u0055\u0050\u0044\u0041\u0059\u0042S", Coupdaybs)
	RegisterFunction("\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053", Coupdays)
	RegisterFunction("\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053\u004e\u0043", Coupdaysnc)
	RegisterFunction("\u0043O\u0055\u0050\u004e\u0055\u004d", Coupnum)
	RegisterFunction("\u0043O\u0055\u0050\u004e\u0043\u0044", Coupncd)
	RegisterFunction("\u0043O\u0055\u0050\u0050\u0043\u0044", Couppcd)
	RegisterFunction("\u0043U\u004d\u0049\u0050\u004d\u0054", Cumipmt)
	RegisterFunction("\u0043\u0055\u004d\u0050\u0052\u0049\u004e\u0043", Cumprinc)
	RegisterFunction("\u0044\u0042", Db)
	RegisterFunction("\u0044\u0044\u0042", Ddb)
	RegisterFunction("\u0044\u0049\u0053\u0043", Disc)
	RegisterFunction("\u0044\u004f\u004c\u004c\u0041\u0052\u0044\u0045", Dollarde)
	RegisterFunction("\u0044\u004f\u004c\u004c\u0041\u0052\u0046\u0052", Dollarfr)
	RegisterFunction("\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e", Duration)
	RegisterFunction("\u0045\u0046\u0046\u0045\u0043\u0054", Effect)
	RegisterFunction("\u0046\u0056", Fv)
	RegisterFunction("\u0046\u0056\u0053\u0043\u0048\u0045\u0044\u0055\u004c\u0045", Fvschedule)
	RegisterFunction("\u0049N\u0054\u0052\u0041\u0054\u0045", Intrate)
	RegisterFunction("\u0049\u0050\u004d\u0054", Ipmt)
	RegisterFunction("\u0049\u0052\u0052", Irr)
	RegisterFunction("\u0049\u0053\u0050M\u0054", Ispmt)
	RegisterFunction("\u004dD\u0055\u0052\u0041\u0054\u0049\u004fN", Mduration)
	RegisterFunction("\u004d\u0049\u0052\u0052", Mirr)
	RegisterFunction("\u004eO\u004d\u0049\u004e\u0041\u004c", Nominal)
	RegisterFunction("\u004e\u0050\u0045\u0052", Nper)
	RegisterFunction("\u004e\u0050\u0056", Npv)
	RegisterFunction("\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E", Oddlprice)
	RegisterFunction("\u004fD\u0044\u004c\u0059\u0049\u0045\u004cD", Oddlyield)
	RegisterFunction("\u0050D\u0055\u0052\u0041\u0054\u0049\u004fN", Pduration)
	RegisterFunction("\u005fx\u006cf\u006e\u002e\u0050\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e", Pduration)
	RegisterFunction("\u0050\u004d\u0054", Pmt)
	RegisterFunction("\u0050\u0050\u004d\u0054", Ppmt)
	RegisterFunction("\u0050\u0052\u0049C\u0045", Price)
	RegisterFunction("\u0050R\u0049\u0043\u0045\u0044\u0049\u0053C", Pricedisc)
	RegisterFunction("\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054", Pricemat)
	RegisterFunction("\u0050\u0056", Pv)
	RegisterFunction("\u0052\u0041\u0054\u0045", Rate)
	RegisterFunction("\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044", Received)
	RegisterFunction("\u0052\u0052\u0049", Rri)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0052\u0052I", Rri)
	RegisterFunction("\u0053\u004c\u004e", Sln)
	RegisterFunction("\u0053\u0059\u0044", Syd)
	RegisterFunction("\u0054B\u0049\u004c\u004c\u0045\u0051", Tbilleq)
	RegisterFunction("\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045", Tbillprice)
	RegisterFunction("\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044", Tbillyield)
	RegisterFunction("\u0056\u0044\u0042", Vdb)
	RegisterFunction("\u0058\u0049\u0052\u0052", Xirr)
	RegisterFunction("\u0058\u004e\u0050\u0056", Xnpv)
	RegisterFunction("\u0059\u0049\u0045L\u0044", Yield)
	RegisterFunction("\u0059I\u0045\u004c\u0044\u0044\u0049\u0053C", Yielddisc)
	RegisterFunction("\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054", Yieldmat)
}

// Small implements the Excel SMALL function.
func Small(args []Result) Result { return _cecd(args, false) }

// Yield implements the Excel YIELD function.
func Yield(args []Result) Result {
	_fddc := len(args)
	if _fddc != 6 && _fddc != 7 {
		return MakeErrorResult("\u0059\u0049E\u004c\u0044\u0020\u0072e\u0071\u0075i\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020o\u0072\u0020\u0073\u0065\u0076\u0065\u006e\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_aefe, _baca, _acggg := _gdge(args[0], args[1], "\u0059\u0049\u0045L\u0044")
	if _acggg.Type == ResultTypeError {
		return _acggg
	}
	_abbcc := args[2]
	if _abbcc.Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045LD\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0072a\u0074e\u0020o\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_adce := _abbcc.ValueNumber
	if _adce < 0 {
		return MakeErrorResultType(ErrorTypeNum, "R\u0061\u0074\u0065\u0020\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006fn\u0020\u006e\u0065g\u0061t\u0069\u0076\u0065")
	}
	_aadf := args[3]
	if _aadf.Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020p\u0072 \u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ebcc := _aadf.ValueNumber
	if _ebcc <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "p\u0072\u0020\u0073\u0068ou\u006cd\u0020\u0062\u0065\u0020\u0070o\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	_fccce := args[4]
	if _fccce.Type != ResultTypeNumber {
		return MakeErrorResult("Y\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065m\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0074yp\u0065\u0020\u006eu\u006db\u0065\u0072")
	}
	_dcbf := _fccce.ValueNumber
	if _dcbf < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	_gbcb := args[5]
	if _gbcb.Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0066\u0072\u0065\u0071\u0075e\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_gfdd := float64(int(_gbcb.ValueNumber))
	if !_gabe(_gfdd) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_edee := 0
	if _fddc == 7 && args[6].Type != ResultTypeEmpty {
		_abed := args[6]
		if _abed.Type != ResultTypeNumber {
			return MakeErrorResult("Y\u0049\u0045\u004c\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u0062\u0061\u0073\u0069\u0073\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
		}
		_edee = int(_abed.ValueNumber)
		if !_gaff(_edee) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063o\u0072\u0072\u0065\u0063t\u0020\u0062\u0061\u0073\u0069\u0073\u0020v\u0061\u006c\u0075\u0065\u0020\u0066\u006f\u0072\u0020\u0059\u0049\u0045\u004c\u0044")
		}
	}
	_ccde := 0.0
	_adbg := 0.0
	_ecde := 1.0
	_fgeg, _acggg := _gaa(_aefe, _baca, _adce, _adbg, _dcbf, _gfdd, _edee)
	if _acggg.Type == ResultTypeError {
		return _acggg
	}
	_dfggb, _acggg := _gaa(_aefe, _baca, _adce, _ecde, _dcbf, _gfdd, _edee)
	if _acggg.Type == ResultTypeError {
		return _acggg
	}
	_gcac := (_ecde - _adbg) * 0.5
	for _dbacf := 0; _dbacf < 100 && _ccde != _ebcc; _dbacf++ {
		_ccde, _acggg = _gaa(_aefe, _baca, _adce, _gcac, _dcbf, _gfdd, _edee)
		if _acggg.Type == ResultTypeError {
			return _acggg
		}
		if _ebcc == _fgeg {
			return MakeNumberResult(_adbg)
		} else if _ebcc == _dfggb {
			return MakeNumberResult(_ecde)
		} else if _ebcc == _ccde {
			return MakeNumberResult(_gcac)
		} else if _ebcc < _dfggb {
			_ecde *= 2.0
			_dfggb, _acggg = _gaa(_aefe, _baca, _adce, _ecde, _dcbf, _gfdd, _edee)
			if _acggg.Type == ResultTypeError {
				return _acggg
			}
			_gcac = (_ecde - _adbg) * 0.5
		} else {
			if _ebcc < _ccde {
				_adbg = _gcac
				_fgeg = _ccde
			} else {
				_ecde = _gcac
				_dfggb = _ccde
			}
			_gcac = _ecde - (_ecde-_adbg)*((_ebcc-_dfggb)/(_fgeg-_dfggb))
		}
	}
	return MakeNumberResult(_gcac)
}
func _gdge(_bdc, _dgbe Result, _cccc string) (float64, float64, Result) {
	_dbdc, _geec := _badf(_bdc, "\u0073e\u0074t\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065", _cccc)
	if _geec.Type == ResultTypeError {
		return 0, 0, _geec
	}
	_fgf, _geec := _badf(_dgbe, "\u006d\u0061\u0074\u0075\u0072\u0069\u0074\u0079\u0020\u0064\u0061\u0074\u0065", _cccc)
	if _geec.Type == ResultTypeError {
		return 0, 0, _geec
	}
	if _dbdc >= _fgf {
		return 0, 0, MakeErrorResultType(ErrorTypeNum, _cccc+"\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020m\u0061\u0074\u0075r\u0069\u0074\u0079\u0020\u0064\u0061\u0074\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u006cat\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065")
	}
	return _dbdc, _fgf, _efc
}
func _cfbd(_egfb []Result) (float64, float64, Result) {
	_geff := 0.0
	_ddfaf := 1.0
	for _, _gdfb := range _egfb {
		switch _gdfb.Type {
		case ResultTypeNumber:
			_geff += _gdfb.ValueNumber
			_ddfaf *= _dfbag(_gdfb.ValueNumber)
		case ResultTypeList, ResultTypeArray:
			_aadc, _fbcg, _fggf := _cfbd(_gdfb.ListValues())
			_geff += _aadc
			_ddfaf *= _dfbag(_fbcg)
			if _fggf.Type == ResultTypeError {
				return 0, 0, _fggf
			}
		case ResultTypeString:
			return 0, 0, MakeErrorResult("M\u0055\u004c\u0054\u0049\u004e\u004f\u004d\u0049\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063 a\u0072\u0067\u0075m\u0065n\u0074\u0073")
		case ResultTypeError:
			return 0, 0, _gdfb
		}
	}
	return _geff, _ddfaf, _efc
}

// Reference returns an invalid reference for Bool.
func (_age Bool) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Searchb is an implementation of the Excel SEARCHB().
func Searchb(ctx Context, ev Evaluator, args []Result) Result {
	if !ctx.IsDBCS() {
		return Search(args)
	}
	_aceec, _egabf := _debf("\u0046\u0049\u004e\u0044", args)
	if _egabf.Type != ResultTypeEmpty {
		return _egabf
	}
	_cfbb := _dda.ToLower(_aceec._fcgea)
	_eagff := _dda.ToLower(_aceec._bfgd)
	if _cfbb == "" {
		return MakeNumberResult(1.0)
	}
	_gbce := _aceec._bdgb - 1
	_becad := 1
	_cafed := 0
	for _ggdg := range _eagff {
		if _ggdg != 0 {
			_bddag := 1
			if _ggdg-_cafed > 1 {
				_bddag = 2
			}
			_becad += _bddag
		}
		if _becad > _gbce {
			_abffg := _ec.Index(_cfbb, _eagff[_ggdg:])
			if _abffg == 0 {
				return MakeNumberResult(float64(_becad))
			}
		}
		_cafed = _ggdg
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}

// LastColumn returns empty string for the invalid reference context.
func (_dgec *ivr) LastColumn(rowFrom, rowTo int) string { return "" }
func _bcbdd(_edbf _ee.Time) bool                        { return _ee.Now().Sub(_edbf) >= _cfaee }

// NewSheetPrefixExpr constructs a new prefix expression.
func NewSheetPrefixExpr(s string) Expression { return &SheetPrefixExpr{_daeae: s} }

var _fed = [...]uint8{0, 16, 29, 43, 56, 68, 80, 91, 102, 113, 125, 137, 148, 163}

const _ccge = 57356

// Clean is an implementation of the Excel CLEAN function that removes
// unprintable characters.
func Clean(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0043\u004c\u0045\u0041\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_agaf := args[0].AsString()
	if _agaf.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u0048\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_dbbfd := _gg.Buffer{}
	for _, _bgdgg := range _agaf.ValueString {
		if _dd.IsPrint(_bgdgg) {
			_dbbfd.WriteRune(_bgdgg)
		}
	}
	return MakeStringResult(_dbbfd.String())
}

// Eval evaluates and returns the result of the NamedRangeRef reference.
func (_eagfa NamedRangeRef) Eval(ctx Context, ev Evaluator) Result {
	_cbgfg := ctx.NamedRange(_eagfa._bbfde)
	_bfac := _cbgfg.Value
	if _ggbf, _ddeb := ev.GetFromCache(_bfac); _ddeb {
		return _ggbf
	}
	_ggaff := _dda.Split(_bfac, "\u0021")
	if len(_ggaff) != 2 {
		return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006e\u0061\u006de\u0064 \u0072\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0025\u0073", _bfac))
	}
	_gccf := ctx.Sheet(_ggaff[0])
	_bfeee := _dda.Split(_ggaff[1], "\u003a")
	switch len(_bfeee) {
	case 1:
		_gcdfd := ev.Eval(_gccf, _bfeee[0])
		ev.SetCache(_bfac, _gcdfd)
		return _gcdfd
	case 2:
		_cffcf := _cfcef(_gccf, ev, _bfeee[0], _bfeee[1])
		ev.SetCache(_bfac, _cffcf)
		return _cffcf
	}
	return MakeErrorResult(_d.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070e \u0025\u0073", _cbgfg.Type))
}
func _gcee(_daed Result) []Result {
	_aagbe := _daed.ValueList
	if _daed.Type == ResultTypeArray {
		_aagbe = nil
		for _, _bgae := range _daed.ValueArray {
			if len(_bgae) > 0 {
				_aagbe = append(_aagbe, _bgae[0])
			} else {
				_aagbe = append(_aagbe, _efc)
			}
		}
	}
	return _aagbe
}

// LookupFunction looks up and returns a standard function or nil.
func LookupFunction(name string) Function {
	_cfbcf.Lock()
	defer _cfbcf.Unlock()
	if _aaed, _egca := _feef[name]; _egca {
		return _aaed
	}
	return nil
}

const _bdcc = 57353

func _afdc(_agecc string, _gaeab func(_gcebe float64) float64) Function {
	return func(_adaea []Result) Result {
		if len(_adaea) != 1 {
			return MakeErrorResult(_agecc + "\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
		}
		_bccg := _adaea[0].AsNumber()
		switch _bccg.Type {
		case ResultTypeNumber:
			_cbff := _gaeab(_bccg.ValueNumber)
			if _cc.IsNaN(_cbff) {
				return MakeErrorResult(_agecc + "\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0065\u0064\u0020\u004e\u0061\u004e")
			}
			if _cc.IsInf(_cbff, 0) {
				return MakeErrorResult(_agecc + "\u0020r\u0065t\u0075\u0072\u006e\u0065\u0064 \u0069\u006ef\u0069\u006e\u0069\u0074\u0079")
			}
			if _cbff == 0 {
				return MakeErrorResultType(ErrorTypeDivideByZero, _agecc+"\u0020d\u0069v\u0069\u0064\u0065\u0020\u0062\u0079\u0020\u007a\u0065\u0072\u006f")
			}
			return MakeNumberResult(1 / _cbff)
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult(_agecc + "\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072i\u0063\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		case ResultTypeError:
			return _bccg
		default:
			return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0068a\u006e\u0064\u006c\u0065d\u0020\u0025\u0073\u0028\u0029\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _agecc, _bccg.Type))
		}
	}
}
func _bfbd(_adbgc string, _bead func(_babgg float64) float64) Function {
	return func(_dbaa []Result) Result {
		if len(_dbaa) != 1 {
			return MakeErrorResult(_adbgc + "\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
		}
		_bedb := _dbaa[0].AsNumber()
		switch _bedb.Type {
		case ResultTypeNumber:
			_cdcd := _bead(_bedb.ValueNumber)
			if _cc.IsNaN(_cdcd) {
				return MakeErrorResult(_adbgc + "\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0065\u0064\u0020\u004e\u0061\u004e")
			}
			if _cc.IsInf(_cdcd, 0) {
				return MakeErrorResult(_adbgc + "\u0020r\u0065t\u0075\u0072\u006e\u0065\u0064 \u0069\u006ef\u0069\u006e\u0069\u0074\u0079")
			}
			return MakeNumberResult(_cdcd)
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult(_adbgc + "\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072i\u0063\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		case ResultTypeError:
			return _bedb
		default:
			return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0068a\u006e\u0064\u006c\u0065d\u0020\u0025\u0073\u0028\u0029\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _adbgc, _bedb.Type))
		}
	}
}

const _efcc = 57350

type yyParserImpl struct {
	_gddfa yySymType
	_fbcd  [_gceebf]yySymType
	_accfg int
}

// True is an implementation of the Excel TRUE() function.  It takes no
// arguments.
func True(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u0054\u0052\u0055E \u0074\u0061\u006b\u0065\u0073\u0020\u006e\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	return MakeBoolResult(true)
}

// Intrate implements the Excel INTRATE function.
func Intrate(args []Result) Result {
	_dafa := len(args)
	if _dafa != 4 && _dafa != 5 {
		return MakeErrorResult("\u0049\u004e\u0054\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0066\u006f\u0075r\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_agba, _cced, _fcdc := _gdge(args[0], args[1], "\u0049N\u0054\u0052\u0041\u0054\u0045")
	if _fcdc.Type == ResultTypeError {
		return _fcdc
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u004e\u0054\u0052\u0041\u0054E\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0069\u006e\u0076\u0065\u0073\u0074\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_abae := args[2].ValueNumber
	if _abae <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u004e\u0054\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0069\u006e\u0076e\u0073\u0074\u006d\u0065\u006e\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u0061r\u0067\u0075\u006de\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u004e\u0054\u0052\u0041\u0054E\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_dbdb := args[3].ValueNumber
	if _dbdb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u004e\u0054\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064e\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u0061r\u0067\u0075\u006de\u006e\u0074")
	}
	_ecagb := 0
	if _dafa == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0049N\u0054\u0052A\u0054\u0045\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
		}
		_ecagb = int(args[4].ValueNumber)
		if !_gaff(_ecagb) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006eco\u0072\u0072\u0065c\u0074\u0020\u0062\u0061sis\u0020ar\u0067\u0075\u006d\u0065\u006e\u0074\u0020fo\u0072\u0020\u0049\u004e\u0054\u0052\u0041T\u0045")
		}
	}
	_eecc, _fcdc := _agef(_agba, _cced, _ecagb)
	if _fcdc.Type == ResultTypeError {
		return _fcdc
	}
	return MakeNumberResult((_dbdb - _abae) / _abae / _eecc)
}

// VLookup implements the VLOOKUP function that returns a matching value from a
// column in an array.
func VLookup(args []Result) Result {
	_aaadf := len(args)
	if _aaadf < 3 {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004bU\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _aaadf > 4 {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020m\u006f\u0073\u0074\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_adef := args[0]
	_bfbfb := args[1]
	if _bfbfb.Type != ResultTypeArray {
		return MakeErrorResult("\u0056\u004cO\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_agcb := args[2].AsNumber()
	if _agcb.Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u004cO\u004f\u004b\u0055\u0050 \u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075m\u0065\u0072\u0069\u0063\u0020\u0063\u006f\u006c\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_dgce := false
	if _aaadf == 4 && args[3].Type != ResultTypeEmpty {
		_bgdga := args[3].AsNumber()
		if _bgdga.Type != ResultTypeNumber {
			return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		if _bgdga.ValueNumber == 0 {
			_dgce = true
		}
	}
	_adad := int(_agcb.ValueNumber) - 1
	_afff := -1
	_ebaa := false
_ffcag:
	for _bdd, _dbda := range _bfbfb.ValueArray {
		if len(_dbda) == 0 {
			continue
		}
		_dgedd := _dbda[0]
		switch _dbfga(_dgedd, _adef, false, _dgce) {
		case _geagf:
			_afff = _bdd
		case _fbagf:
			_afff = _bdd
			_ebaa = true
			break _ffcag
		}
	}
	if _afff == -1 {
		return MakeErrorResultType(ErrorTypeNA, "\u0056\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_eccb := _bfbfb.ValueArray[_afff]
	if _adad < 0 || _adad >= len(_eccb) {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0063\u006f\u006c\u0075\u006d\u006e\u0020\u0069\u006e\u0064\u0065\u0078")
	}
	if _ebaa || !_dgce {
		return _eccb[_adad]
	}
	return MakeErrorResultType(ErrorTypeNA, "\u0056\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

// Quotient is an implementation of the Excel QUOTIENT function that returns the
// integer portion of division.
func Quotient(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0051\u0055\u004f\u0054\u0049E\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0074\u0077\u006f\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_aebe := args[0].AsNumber()
	_bcga := args[1].AsNumber()
	if _aebe.Type != ResultTypeNumber || _bcga.Type != ResultTypeNumber {
		return MakeErrorResult("\u0051\u0055\u004f\u0054\u0049E\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0074\u0077\u006f\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _bcga.ValueNumber == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0051U\u004f\u0054\u0049\u0045N\u0054\u0028\u0029\u0020\u0064i\u0076i\u0064e\u0020\u0062\u0079\u0020\u007a\u0065\u0072o")
	}
	return MakeNumberResult(_cc.Trunc(_aebe.ValueNumber / _bcga.ValueNumber))
}

// String returns an empty string for Error.
func (_eaaa Error) String() string { return "" }

// Median implements the MEDIAN function that returns the median of a range of
// values.
func Median(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004d\u0045D\u0049\u0041\u004e\u0020r\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020l\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_deea := _fbdc(args)
	_b.Float64s(_deea)
	var _feeac float64
	if len(_deea)%2 == 0 {
		_feeac = (_deea[len(_deea)/2-1] + _deea[len(_deea)/2]) / 2
	} else {
		_feeac = _deea[len(_deea)/2]
	}
	return MakeNumberResult(_feeac)
}

const _fceaf = 57360

// SumIf implements the SUMIF function.
func SumIf(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("\u0053\u0055\u004d\u0049\u0046\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0074\u0068\u0072e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ddbga := args[0]
	if _ddbga.Type != ResultTypeArray && _ddbga.Type != ResultTypeList {
		return MakeErrorResult("\u0053\u0055\u004d\u0049\u0046\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0066i\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cbdf := _afgc(_ddbga)
	_bceb := args[2]
	if _bceb.Type != ResultTypeArray && _bceb.Type != ResultTypeList {
		return MakeErrorResult("\u0053\u0055\u004dI\u0046\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0061\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074y\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_bedf := _afgc(_bceb)
	_bbeec := _fbea(args[1])
	_cfcga := 0.0
	for _cbgbb, _cffg := range _cbdf {
		for _aceac, _ebccb := range _cffg {
			if _caaf(_ebccb, _bbeec) {
				_cfcga += _bedf[_cbgbb][_aceac].ValueNumber
			}
		}
	}
	return MakeNumberResult(_cfcga)
}

// IfNA is an implementation of the Excel IFNA() function. It takes two arguments.
func IfNA(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("I\u0046\u004e\u0041\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	if args[0].Type == ResultTypeError && args[0].ValueString == "\u0023\u004e\u002f\u0041" {
		return args[1]
	}
	return args[0]
}

// Tbillprice implements the Excel TBILLPRICE function.
func Tbillprice(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("T\u0042\u0049\u004c\u004c\u0050\u0052I\u0043\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_bade, _dfac, _ddafe := _gdge(args[0], args[1], "\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045")
	if _ddafe.Type == ResultTypeError {
		return _ddafe
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0064\u0069\u0073\u0063\u006f\u0075n\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ffec := _dfac - _bade
	if _ffec > 365 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004cP\u0052\u0049\u0043\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020m\u0061\u0074\u0075r\u0069\u0074\u0079\u0020t\u006f\u0020\u0062\u0065\u0020\u006eo\u0074\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065\u0020\u0079e\u0061\u0072\u0020\u0061\u0066\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074")
	}
	_dgde := args[2].ValueNumber
	if _dgde <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020d\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeNumberResult(100 * (1 - _dgde*_ffec/360))
}
func _gba() evCache {
	_bde := evCache{}
	_bde._ada = make(map[string]Result)
	_bde._ddg = &_fe.Mutex{}
	return _bde
}

// Choose implements the Excel CHOOSE function.
func Choose(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0043\u0048O\u004f\u0053\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	_fcegeb := args[0]
	if _fcegeb.Type != ResultTypeNumber {
		return MakeErrorResult("\u0043H\u004f\u004fS\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u0069\u0072\u0073\u0074\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074y\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_bbff := int(_fcegeb.ValueNumber)
	if _bbff < 1 {
		return MakeErrorResult("\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u0061 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	if len(args) <= _bbff {
		return MakeErrorResult("\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0068\u006f\u0075\u006cd\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073 \u006fr\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0076\u0061\u006c\u0075\u0065\u0073")
	}
	return args[_bbff]
}

// Vdb implements the Excel VDB function.
func Vdb(args []Result) Result {
	_cgfd := len(args)
	if _cgfd < 5 || _cgfd > 7 {
		return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0066\u0069\u0076\u0065\u0020a\u006e\u0064\u0020\u0073\u0065v\u0065\u006e")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_bebe := args[0].ValueNumber
	if _bebe < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044B \u0072\u0065\u0071u\u0069\u0072\u0065\u0073 co\u0073t \u0074\u006f\u0020\u0062\u0065\u0020\u006eon\u0020\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u0044\u0042 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_aaea := args[1].ValueNumber
	if _aaea < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0061\u006c\u0076\u0061\u0067\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gefd := args[2].ValueNumber
	if _gefd == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if _gefd < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("V\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0073\u0074\u0061\u0072\u0074 p\u0065\u0072\u0069\u006fd\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dcef := args[3].ValueNumber
	if _dcef < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074o\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u006c\u0065\u0073\u0073\u0020\u0074h\u0061n\u0020\u006f\u006e\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056D\u0042\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0065\u006e\u0064 \u0070\u0065\u0072\u0069\u006f\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bdac := args[4].ValueNumber
	if _dcef > _bdac {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020s\u0074\u0061\u0072\u0074\u0020\u0070\u0065r\u0069\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0056\u0044\u0042")
	}
	if _bdac > _gefd {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0065\u006e\u0064\u0020\u0070e\u0072i\u006f\u0064\u0020\u0066\u006f\u0072\u0020V\u0044\u0042")
	}
	_fbdbc := 2.0
	if _cgfd > 5 {
		if args[5].Type == ResultTypeEmpty {
			_fbdbc = 0.0
		} else {
			if args[5].Type != ResultTypeNumber {
				return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
			}
			_fbdbc = args[5].ValueNumber
			if _fbdbc < 0 {
				return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
			}
		}
	}
	_gdbb := false
	if _cgfd > 6 && args[6].Type != ResultTypeEmpty {
		if args[6].Type != ResultTypeNumber {
			return MakeErrorResult("\u0056D\u0042\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020n\u006f\u005f\u0073\u0077\u0069\u0074\u0063\u0068\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
		}
		_gdbb = args[6].ValueNumber != 0
	}
	_febc := 0.0
	_eagg := _cc.Floor(_dcef)
	_ggad := _cc.Ceil(_bdac)
	if _gdbb {
		for _gcbc := _eagg + 1; _gcbc <= _ggad; _gcbc++ {
			_dbbf := _afec(_bebe, _aaea, _gefd, _gcbc, _fbdbc)
			if _gcbc == _eagg+1 {
				_dbbf *= _cc.Min(_bdac, _eagg+1) - _dcef
			} else if _gcbc == _ggad {
				_dbbf *= _bdac + 1 - _ggad
			}
			_febc += _dbbf
		}
	} else {
		_fbc := _gefd
		var _eggf float64
		if !_afddf(_dcef, _cc.Floor(_dcef)) {
			if _fbdbc == 1 {
				_ffaa := _gefd / 2
				if _dcef > _ffaa || _afddf(_dcef, _ffaa) {
					_eggf = _dcef - _ffaa
					_dcef = _ffaa
					_bdac -= _eggf
					_fbc++
				}
			}
		}
		if _fbdbc != 0 {
			_bebe -= _gfcae(_bebe, _aaea, _gefd, _fbc, _dcef, _fbdbc)
		}
		_febc = _gfcae(_bebe, _aaea, _gefd, _gefd-_dcef, _bdac-_dcef, _fbdbc)
	}
	return MakeNumberResult(_febc)
}

// Now is an implementation of the Excel NOW() function.
func Now(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("\u004e\u004fW\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	_ebgg := _ee.Now()
	_, _bbf := _ebgg.Zone()
	_cfbe := _gab + float64(_ebgg.Unix()+int64(_bbf))/86400
	return MakeNumberResult(_cfbe)
}

// Parse parses an io.Reader to get an Expression. If expression is parsed with an error, nil is returned
func Parse(r _g.Reader) Expression {
	_ggec := &plex{_bede: LexReader(r)}
	_efebc(_ggec)
	if _ggec._dccbf != "" {
		return nil
	}
	return _ggec._gggcf
}
func _fbea(_baecb Result) *criteriaParsed {
	_bgbd := _baecb.Type == ResultTypeNumber
	_dbcd := _baecb.ValueNumber
	_gffa := _dda.ToLower(_baecb.ValueString)
	_ffcfd := _ffbbd(_gffa)
	return &criteriaParsed{_bgbd, _dbcd, _gffa, _ffcfd}
}
func _fcb(_bgfd, _dbac int64) float64 { return float64(int(0.5 + float64((_dbac-_bgfd)/86400))) }

// String returns a string representation of Number.
func (_eedb Number) String() string { return _c.FormatFloat(_eedb._ceacb, 'f', -1, 64) }

// CountIf implements the COUNTIF function.
func CountIf(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0043\u004f\u0055N\u0054\u0049\u0046\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0073")
	}
	_dcac := args[0]
	if _dcac.Type != ResultTypeArray && _dcac.Type != ResultTypeList {
		return MakeErrorResult("\u0043O\u0055\u004eT\u0049\u0046\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0066\u0069\u0072\u0073\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cfca := _fbea(args[1])
	_bgbgc := 0
	for _, _ebdcc := range _afgc(_dcac) {
		for _, _gfddc := range _ebdcc {
			if _caaf(_gfddc, _cfca) {
				_bgbgc++
			}
		}
	}
	return MakeNumberResult(float64(_bgbgc))
}

const _bbce int = 30

// ISERROR is an implementation of the Excel ISERROR() function.
func IsError(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("I\u0053\u0045\u0052\u0052\u004f\u0052(\u0029\u0020\u0061\u0063\u0063\u0065p\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeError)
}
func (_acdg *Lexer) lex(_egdb _g.Reader) {
	_defgc, _ccca, _agfa := 0, 0, 0
	_ecdbg := -1
	_aebd, _cbgdb, _cabb := 0, 0, 0
	_ = _cabb
	_fgfeg := 1
	_ = _fgfeg
	_geceg := make([]byte, 4096)
	_afce := false
	for !_afce {
		_fdabag := 0
		if _aebd > 0 {
			_fdabag = _ccca - _aebd
		}
		_ccca = 0
		_fbcdb, _bgcg := _egdb.Read(_geceg[_fdabag:])
		if _fbcdb == 0 || _bgcg != nil {
			_afce = true
		}
		_agfa = _fbcdb + _fdabag
		if _agfa < len(_geceg) {
			_ecdbg = _agfa
		}
		{
			_defgc = _dbbbb
			_aebd = 0
			_cbgdb = 0
			_cabb = 0
		}
		{
			var _gcff int
			var _accce uint
			if _ccca == _agfa {
				goto _bbcdb
			}
			if _defgc == 0 {
				goto _cgbg
			}
		_beee:
			_gcff = int(_geece[_defgc])
			_accce = uint(_agbgdb[_gcff])
			_gcff++
			for ; _accce > 0; _accce-- {
				_gcff++
				switch _agbgdb[_gcff-1] {
				case 2:
					_aebd = _ccca
				}
			}
			switch _defgc {
			case 30:
				switch _geceg[_ccca] {
				case 34:
					goto _daae
				case 35:
					goto _fegb
				case 36:
					goto _dafg
				case 38:
					goto _fccggf
				case 39:
					goto _egecf
				case 40:
					goto _bfad
				case 41:
					goto _edgdf
				case 42:
					goto _ceccd
				case 43:
					goto _eceb
				case 44:
					goto _bcbeg
				case 45:
					goto _ddgf
				case 47:
					goto _acde
				case 58:
					goto _bfcgd
				case 59:
					goto _eaeb
				case 60:
					goto _efef
				case 61:
					goto _caaab
				case 62:
					goto _ceadd
				case 63:
					goto _afdgg
				case 70:
					goto _ccbdc
				case 84:
					goto _dcfe
				case 92:
					goto _cgga
				case 94:
					goto _beedc
				case 95:
					goto _bdcg
				case 123:
					goto _afecf
				case 125:
					goto _feedb
				}
				switch {
				case _geceg[_ccca] < 65:
					switch {
					case _geceg[_ccca] > 37:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _faaf
						}
					case _geceg[_ccca] >= 33:
						goto _afdgg
					}
				case _geceg[_ccca] > 90:
					switch {
					case _geceg[_ccca] > 93:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _ebffa
						}
					case _geceg[_ccca] >= 91:
						goto _afdgg
					}
				default:
					goto _dfgfa
				}
				goto _cccb
			case 1:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _cace
				case 123:
					goto _cace
				case 125:
					goto _cace
				}
				switch {
				case _geceg[_ccca] < 37:
					if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
						goto _cace
					}
				case _geceg[_ccca] > 45:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _cace
						}
					case _geceg[_ccca] >= 58:
						goto _cace
					}
				default:
					goto _cace
				}
				goto _cccb
			case 0:
				goto _cgbg
			case 2:
				if _geceg[_ccca] == 34 {
					goto _cgafa
				}
				goto _daae
			case 31:
				if _geceg[_ccca] == 34 {
					goto _daae
				}
				goto _eceec
			case 3:
				switch _geceg[_ccca] {
				case 78:
					goto _egbba
				case 82:
					goto _gbda
				}
				goto _afdgg
			case 4:
				switch _geceg[_ccca] {
				case 47:
					goto _dfdcb
				case 85:
					goto _eeaf
				}
				goto _afdgg
			case 5:
				if _geceg[_ccca] == 65 {
					goto _bedd
				}
				goto _afdgg
			case 6:
				switch _geceg[_ccca] {
				case 76:
					goto _gcfa
				case 77:
					goto _bbag
				}
				goto _afdgg
			case 7:
				if _geceg[_ccca] == 76 {
					goto _bbag
				}
				goto _afdgg
			case 8:
				if _geceg[_ccca] == 33 {
					goto _bedd
				}
				goto _afdgg
			case 9:
				if _geceg[_ccca] == 69 {
					goto _fcfg
				}
				goto _afdgg
			case 10:
				if _geceg[_ccca] == 70 {
					goto _acfba
				}
				goto _afdgg
			case 11:
				if _geceg[_ccca] == 33 {
					goto _fgcac
				}
				goto _afdgg
			case 12:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _afdgg
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 45 {
							goto _afdgg
						}
					case _geceg[_ccca] >= 34:
						goto _afdgg
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] < 65:
						if 58 <= _geceg[_ccca] && _geceg[_ccca] <= 63 {
							goto _afdgg
						}
					case _geceg[_ccca] > 90:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _afdgg
						}
					default:
						goto _cabdf
					}
				default:
					goto _ebcce
				}
				goto _cccb
			case 13:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _afdgg
				case 58:
					goto _edgc
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 45 {
							goto _afdgg
						}
					case _geceg[_ccca] >= 34:
						goto _afdgg
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _afdgg
						}
					case _geceg[_ccca] >= 59:
						goto _afdgg
					}
				default:
					goto _ebcce
				}
				goto _cccb
			case 14:
				if _geceg[_ccca] == 36 {
					goto _dgfa
				}
				if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
					goto _bbbad
				}
				goto _cace
			case 15:
				if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
					goto _bbbad
				}
				goto _cace
			case 32:
				if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
					goto _bbbad
				}
				goto _cacde
			case 16:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 47:
					goto _afdgg
				case 58:
					goto _dgaff
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 45:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _deeae
						}
					case _geceg[_ccca] >= 34:
						goto _afdgg
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] > 90:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _afdgg
						}
					case _geceg[_ccca] >= 65:
						goto _cabdf
					}
				default:
					goto _afdgg
				}
				goto _cccb
			case 17:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _cace
				case 123:
					goto _cace
				case 125:
					goto _cace
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 45 {
							goto _cace
						}
					case _geceg[_ccca] >= 34:
						goto _cace
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _cace
						}
					case _geceg[_ccca] >= 58:
						goto _cace
					}
				default:
					goto _deeae
				}
				goto _cccb
			case 33:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _fdcga
				case 123:
					goto _fdcga
				case 125:
					goto _fdcga
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 45 {
							goto _fdcga
						}
					case _geceg[_ccca] >= 34:
						goto _fdcga
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _fdcga
						}
					case _geceg[_ccca] >= 58:
						goto _fdcga
					}
				default:
					goto _deeae
				}
				goto _cccb
			case 18:
				if _geceg[_ccca] == 36 {
					goto _fdde
				}
				if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
					goto _fgeb
				}
				goto _cace
			case 19:
				if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
					goto _fgeb
				}
				goto _cace
			case 34:
				if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
					goto _fgeb
				}
				goto _gacef
			case 20:
				switch _geceg[_ccca] {
				case 39:
					goto _afdgg
				case 42:
					goto _afdgg
				case 47:
					goto _afdgg
				case 58:
					goto _afdgg
				case 63:
					goto _afdgg
				}
				if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 93 {
					goto _afdgg
				}
				goto _cgdfc
			case 21:
				switch _geceg[_ccca] {
				case 39:
					goto _fadfd
				case 42:
					goto _afdgg
				case 47:
					goto _afdgg
				case 58:
					goto _afdgg
				case 63:
					goto _afdgg
				}
				if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 93 {
					goto _afdgg
				}
				goto _cgdfc
			case 22:
				if _geceg[_ccca] == 33 {
					goto _ffgd
				}
				goto _afdgg
			case 35:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _fcdde
				case 58:
					goto _edgc
				case 101:
					goto _eddgg
				case 123:
					goto _fbcgf
				case 125:
					goto _fbcgf
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 47 {
							goto _fbcgf
						}
					case _geceg[_ccca] >= 34:
						goto _fbcgf
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _fbcgf
						}
					case _geceg[_ccca] >= 59:
						goto _fbcgf
					}
				default:
					goto _faaf
				}
				goto _cccb
			case 36:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _fbcgf
				case 101:
					goto _eddgg
				case 123:
					goto _fbcgf
				case 125:
					goto _fbcgf
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 45 {
							goto _fbcgf
						}
					case _geceg[_ccca] >= 34:
						goto _fbcgf
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _fbcgf
						}
					case _geceg[_ccca] >= 58:
						goto _fbcgf
					}
				default:
					goto _fcdde
				}
				goto _cccb
			case 23:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _eadca
				case 123:
					goto _eadca
				case 125:
					goto _eadca
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 45 {
							goto _eadca
						}
					case _geceg[_ccca] >= 34:
						goto _eadca
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _eadca
						}
					case _geceg[_ccca] >= 58:
						goto _eadca
					}
				default:
					goto _faebf
				}
				goto _cccb
			case 37:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 47:
					goto _fbcgf
				case 123:
					goto _fbcgf
				case 125:
					goto _fbcgf
				}
				switch {
				case _geceg[_ccca] < 48:
					switch {
					case _geceg[_ccca] > 35:
						if 37 <= _geceg[_ccca] && _geceg[_ccca] <= 45 {
							goto _fbcgf
						}
					case _geceg[_ccca] >= 34:
						goto _fbcgf
					}
				case _geceg[_ccca] > 57:
					switch {
					case _geceg[_ccca] > 63:
						if 91 <= _geceg[_ccca] && _geceg[_ccca] <= 94 {
							goto _fbcgf
						}
					case _geceg[_ccca] >= 58:
						goto _fbcgf
					}
				default:
					goto _faebf
				}
				goto _cccb
			case 38:
				switch _geceg[_ccca] {
				case 61:
					goto _gdea
				case 62:
					goto _ecce
				}
				goto _fbgcd
			case 39:
				if _geceg[_ccca] == 61 {
					goto _egee
				}
				goto _gdae
			case 24:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _afdgg
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _afdgg
					}
				default:
					goto _afdgg
				}
				goto _cccb
			case 40:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _dbea
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _dbea
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 41:
				switch _geceg[_ccca] {
				case 46:
					goto _acfbe
				case 92:
					goto _acfbe
				case 95:
					goto _acfbe
				}
				switch {
				case _geceg[_ccca] < 65:
					if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
						goto _acfbe
					}
				case _geceg[_ccca] > 90:
					if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
						goto _acfbe
					}
				default:
					goto _acfbe
				}
				goto _aedbc
			case 42:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 43:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _fdcga
				case 125:
					goto _fdcga
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _fdcga
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					default:
						goto _fdcga
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _dbea
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _fdcga
					}
				default:
					goto _fdcga
				}
				goto _cccb
			case 44:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _cace
				case 125:
					goto _cace
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _cace
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _cace
					}
				default:
					goto _cace
				}
				goto _cccb
			case 25:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 65:
					goto _gcdd
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _afdgg
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 66 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _afdgg
					}
				default:
					goto _afdgg
				}
				goto _cccb
			case 45:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 76:
					goto _baafc
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 46:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 83:
					goto _ebggb
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 47:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 69:
					goto _ecade
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 26:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 79:
					goto _aaaec
				case 82:
					goto _bage
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _afdgg
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _afdgg
					}
				default:
					goto _afdgg
				}
				goto _cccb
			case 48:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 68:
					goto _eafeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 49:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 79:
					goto _gfga
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 50:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 36:
					goto _ddbff
				case 40:
					goto _gbae
				case 46:
					goto _dbea
				case 58:
					goto _dgaff
				case 85:
					goto _ebggb
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 59:
					switch {
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _efec
						}
					case _geceg[_ccca] >= 34:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _gbcgd
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 27:
				switch _geceg[_ccca] {
				case 46:
					goto _acfbe
				case 92:
					goto _acfbe
				case 95:
					goto _acfbe
				}
				switch {
				case _geceg[_ccca] < 65:
					if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
						goto _acfbe
					}
				case _geceg[_ccca] > 90:
					if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
						goto _acfbe
					}
				default:
					goto _acfbe
				}
				goto _afdgg
			case 28:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 120:
					goto _ggbe
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _afdgg
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _afdgg
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _afdgg
					}
				default:
					goto _afdgg
				}
				goto _cccb
			case 51:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 108:
					goto _gcdeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 52:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 102:
					goto _afge
				case 110:
					goto _fadbf
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 53:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 110:
					goto _eedd
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 54:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _eggd
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 55:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _eeac
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _eeac
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 56:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 40:
					goto _aggf
				case 46:
					goto _eeac
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _eeac
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _eeac
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 57:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 109:
					goto _dfcg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 58:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _ebgce
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _aedbc
				case 125:
					goto _aedbc
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _aedbc
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _aedbc
					}
				default:
					goto _aedbc
				}
				goto _cccb
			case 59:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _aeac
				case 123:
					goto _cace
				case 125:
					goto _cace
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _cace
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _cace
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _aeac
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _cace
					}
				default:
					goto _cace
				}
				goto _cccb
			case 29:
				switch _geceg[_ccca] {
				case 33:
					goto _gbcga
				case 46:
					goto _egbeg
				case 92:
					goto _acfbe
				case 95:
					goto _egbeg
				case 123:
					goto _afdgg
				case 125:
					goto _afdgg
				}
				switch {
				case _geceg[_ccca] < 58:
					switch {
					case _geceg[_ccca] < 37:
						if 34 <= _geceg[_ccca] && _geceg[_ccca] <= 35 {
							goto _afdgg
						}
					case _geceg[_ccca] > 47:
						if 48 <= _geceg[_ccca] && _geceg[_ccca] <= 57 {
							goto _egbeg
						}
					default:
						goto _afdgg
					}
				case _geceg[_ccca] > 63:
					switch {
					case _geceg[_ccca] < 91:
						if 65 <= _geceg[_ccca] && _geceg[_ccca] <= 90 {
							goto _egbeg
						}
					case _geceg[_ccca] > 94:
						if 97 <= _geceg[_ccca] && _geceg[_ccca] <= 122 {
							goto _egbeg
						}
					default:
						goto _afdgg
					}
				default:
					goto _afdgg
				}
				goto _cccb
			}
		_afdgg:
			_defgc = 0
			goto _dggg
		_cccb:
			_defgc = 1
			goto _dggg
		_daae:
			_defgc = 2
			goto _dggg
		_fegb:
			_defgc = 3
			goto _dggg
		_egbba:
			_defgc = 4
			goto _dggg
		_dfdcb:
			_defgc = 5
			goto _dggg
		_eeaf:
			_defgc = 6
			goto _dggg
		_gcfa:
			_defgc = 7
			goto _dggg
		_bbag:
			_defgc = 8
			goto _dggg
		_gbda:
			_defgc = 9
			goto _dggg
		_fcfg:
			_defgc = 10
			goto _dggg
		_acfba:
			_defgc = 11
			goto _dggg
		_dafg:
			_defgc = 12
			goto _dggg
		_ebcce:
			_defgc = 13
			goto _dggg
		_edgc:
			_defgc = 14
			goto _dggg
		_dgfa:
			_defgc = 15
			goto _dggg
		_cabdf:
			_defgc = 16
			goto _dggg
		_ddbff:
			_defgc = 17
			goto _dggg
		_dgaff:
			_defgc = 18
			goto _dggg
		_fdde:
			_defgc = 19
			goto _dggg
		_egecf:
			_defgc = 20
			goto _dggg
		_cgdfc:
			_defgc = 21
			goto _dggg
		_fadfd:
			_defgc = 22
			goto _dggg
		_eddgg:
			_defgc = 23
			goto _dggg
		_dfgfa:
			_defgc = 24
			goto _dggg
		_ccbdc:
			_defgc = 25
			goto _dggg
		_dcfe:
			_defgc = 26
			goto _dggg
		_cgga:
			_defgc = 27
			goto _dggg
		_bdcg:
			_defgc = 28
			goto _dggg
		_ebffa:
			_defgc = 29
			goto _dggg
		_cace:
			_defgc = 30
			goto _ecbgd
		_gbcga:
			_defgc = 30
			goto _edbg
		_bedd:
			_defgc = 30
			goto _ggfd
		_fgcac:
			_defgc = 30
			goto _bedfe
		_ffgd:
			_defgc = 30
			goto _becbd
		_eadca:
			_defgc = 30
			goto _dfdgg
		_gbae:
			_defgc = 30
			goto _bfadf
		_fccggf:
			_defgc = 30
			goto _ababf
		_bfad:
			_defgc = 30
			goto _addd
		_edgdf:
			_defgc = 30
			goto _eedfe
		_ceccd:
			_defgc = 30
			goto _gfcba
		_eceb:
			_defgc = 30
			goto _cdbc
		_bcbeg:
			_defgc = 30
			goto _gcded
		_ddgf:
			_defgc = 30
			goto _fegf
		_acde:
			_defgc = 30
			goto _ebad
		_bfcgd:
			_defgc = 30
			goto _egdaa
		_eaeb:
			_defgc = 30
			goto _ababgc
		_caaab:
			_defgc = 30
			goto _adbd
		_beedc:
			_defgc = 30
			goto _fafg
		_afecf:
			_defgc = 30
			goto _cdge
		_feedb:
			_defgc = 30
			goto _efebd
		_eceec:
			_defgc = 30
			goto _daff
		_cacde:
			_defgc = 30
			goto _febdb
		_fdcga:
			_defgc = 30
			goto _cedda
		_gacef:
			_defgc = 30
			goto _efdabb
		_fbcgf:
			_defgc = 30
			goto _bbgbc
		_fbgcd:
			_defgc = 30
			goto _bace
		_gdea:
			_defgc = 30
			goto _fbeag
		_ecce:
			_defgc = 30
			goto _edab
		_gdae:
			_defgc = 30
			goto _dcgf
		_egee:
			_defgc = 30
			goto _cffce
		_aedbc:
			_defgc = 30
			goto _dgfd
		_aggf:
			_defgc = 30
			goto _ddda
		_cgafa:
			_defgc = 31
			goto _badff
		_bbbad:
			_defgc = 32
			goto _dggg
		_deeae:
			_defgc = 33
			goto _deec
		_fgeb:
			_defgc = 34
			goto _dggg
		_faaf:
			_defgc = 35
			goto _cbfag
		_fcdde:
			_defgc = 36
			goto _cbfag
		_faebf:
			_defgc = 37
			goto _cbfag
		_efef:
			_defgc = 38
			goto _dggg
		_ceadd:
			_defgc = 39
			goto _dggg
		_dbea:
			_defgc = 40
			goto _abeb
		_acfbe:
			_defgc = 41
			goto _dggg
		_egbeg:
			_defgc = 42
			goto _abeb
		_efec:
			_defgc = 43
			goto _deec
		_gbcgd:
			_defgc = 44
			goto _abeb
		_ecade:
			_defgc = 44
			goto _cbee
		_gfga:
			_defgc = 44
			goto _gfddf
		_gcdd:
			_defgc = 45
			goto _abeb
		_baafc:
			_defgc = 46
			goto _abeb
		_ebggb:
			_defgc = 47
			goto _abeb
		_aaaec:
			_defgc = 48
			goto _abeb
		_eafeg:
			_defgc = 49
			goto _abeb
		_bage:
			_defgc = 50
			goto _abeb
		_ggbe:
			_defgc = 51
			goto _abeb
		_gcdeg:
			_defgc = 52
			goto _abeb
		_afge:
			_defgc = 53
			goto _abeb
		_eedd:
			_defgc = 54
			goto _abeb
		_eggd:
			_defgc = 55
			goto _abeb
		_eeac:
			_defgc = 56
			goto _abeb
		_fadbf:
			_defgc = 57
			goto _abeb
		_dfcg:
			_defgc = 58
			goto _abeb
		_ebgce:
			_defgc = 59
			goto _abeb
		_aeac:
			_defgc = 59
			goto _abccd
		_ggfd:
			_gcff = 3
			goto _ggefd
		_bedfe:
			_gcff = 5
			goto _ggefd
		_edbg:
			_gcff = 7
			goto _ggefd
		_becbd:
			_gcff = 9
			goto _ggefd
		_bfadf:
			_gcff = 11
			goto _ggefd
		_ddda:
			_gcff = 13
			goto _ggefd
		_ababf:
			_gcff = 15
			goto _ggefd
		_cdge:
			_gcff = 17
			goto _ggefd
		_efebd:
			_gcff = 19
			goto _ggefd
		_addd:
			_gcff = 21
			goto _ggefd
		_eedfe:
			_gcff = 23
			goto _ggefd
		_cdbc:
			_gcff = 25
			goto _ggefd
		_fegf:
			_gcff = 27
			goto _ggefd
		_gfcba:
			_gcff = 29
			goto _ggefd
		_ebad:
			_gcff = 31
			goto _ggefd
		_fafg:
			_gcff = 33
			goto _ggefd
		_adbd:
			_gcff = 35
			goto _ggefd
		_fbeag:
			_gcff = 37
			goto _ggefd
		_cffce:
			_gcff = 39
			goto _ggefd
		_edab:
			_gcff = 41
			goto _ggefd
		_egdaa:
			_gcff = 43
			goto _ggefd
		_ababgc:
			_gcff = 45
			goto _ggefd
		_gcded:
			_gcff = 47
			goto _ggefd
		_bbgbc:
			_gcff = 49
			goto _ggefd
		_cedda:
			_gcff = 51
			goto _ggefd
		_febdb:
			_gcff = 53
			goto _ggefd
		_efdabb:
			_gcff = 55
			goto _ggefd
		_dgfd:
			_gcff = 57
			goto _ggefd
		_daff:
			_gcff = 59
			goto _ggefd
		_bace:
			_gcff = 61
			goto _ggefd
		_dcgf:
			_gcff = 63
			goto _ggefd
		_dfdgg:
			_gcff = 65
			goto _ggefd
		_ecbgd:
			_gcff = 67
			goto _ggefd
		_cbee:
			_gcff = 72
			goto _ggefd
		_cbfag:
			_gcff = 75
			goto _ggefd
		_deec:
			_gcff = 78
			goto _ggefd
		_gfddf:
			_gcff = 81
			goto _ggefd
		_abccd:
			_gcff = 84
			goto _ggefd
		_abeb:
			_gcff = 87
			goto _ggefd
		_badff:
			_gcff = 90
			goto _ggefd
		_ggefd:
			_accce = uint(_agbgdb[_gcff])
			_gcff++
			for ; _accce > 0; _accce-- {
				_gcff++
				switch _agbgdb[_gcff-1] {
				case 3:
					_cbgdb = _ccca + 1
				case 4:
					_cabb = 1
				case 5:
					_cabb = 2
				case 6:
					_cabb = 3
				case 7:
					_cabb = 4
				case 8:
					_cabb = 11
				case 9:
					_cabb = 14
				case 10:
					_cabb = 15
				case 11:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_ebace, _geceg[_aebd:_cbgdb])
					}
				case 12:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_ccge, _geceg[_aebd:_cbgdb])
					}
				case 13:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_bdbed, _geceg[_aebd:_cbgdb-1])
					}
				case 14:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_bdbed, _geceg[_aebd+1:_cbgdb-2])
					}
				case 15:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_dbfe, _geceg[_aebd:_cbgdb-1])
					}
				case 16:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_dbfe, _geceg[_aebd:_cbgdb-1])
					}
				case 17:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_cdbb, _geceg[_aebd:_cbgdb])
					}
				case 18:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_fceaf, _geceg[_aebd:_cbgdb])
					}
				case 19:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_befd, _geceg[_aebd:_cbgdb])
					}
				case 20:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_geccf, _geceg[_aebd:_cbgdb])
					}
				case 21:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_gfbd, _geceg[_aebd:_cbgdb])
					}
				case 22:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_daad, _geceg[_aebd:_cbgdb])
					}
				case 23:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_ageaa, _geceg[_aebd:_cbgdb])
					}
				case 24:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_bcafg, _geceg[_aebd:_cbgdb])
					}
				case 25:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_adga, _geceg[_aebd:_cbgdb])
					}
				case 26:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_fcaf, _geceg[_aebd:_cbgdb])
					}
				case 27:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_beea, _geceg[_aebd:_cbgdb])
					}
				case 28:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_ggafe, _geceg[_aebd:_cbgdb])
					}
				case 29:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_dbfc, _geceg[_aebd:_cbgdb])
					}
				case 30:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_feabb, _geceg[_aebd:_cbgdb])
					}
				case 31:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_ceccg, _geceg[_aebd:_cbgdb])
					}
				case 32:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_fbfe, _geceg[_aebd:_cbgdb])
					}
				case 33:
					_cbgdb = _ccca + 1
					{
						_acdg.emit(_cdacg, _geceg[_aebd:_cbgdb])
					}
				case 34:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_bdcc, _geceg[_aebd:_cbgdb])
					}
				case 35:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_adab, _geceg[_aebd:_cbgdb])
					}
				case 36:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_dgdf, _geceg[_aebd:_cbgdb])
					}
				case 37:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_dagfb, _geceg[_aebd:_cbgdb])
					}
				case 38:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_dcfd, _geceg[_aebd:_cbgdb])
					}
				case 39:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_eaea, _geceg[_aebd+1:_cbgdb-1])
					}
				case 40:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_fdca, _geceg[_aebd:_cbgdb])
					}
				case 41:
					_cbgdb = _ccca
					_ccca--
					{
						_acdg.emit(_caddf, _geceg[_aebd:_cbgdb])
					}
				case 42:
					_ccca = (_cbgdb) - 1
					{
						_acdg.emit(_bdcc, _geceg[_aebd:_cbgdb])
					}
				case 43:
					switch _cabb {
					case 0:
						{
							_defgc = 0
							goto _dggg
						}
					case 1:
						{
							_ccca = (_cbgdb) - 1
							_acdg.emit(_dbgca, _geceg[_aebd:_cbgdb])
						}
					case 2:
						{
							_ccca = (_cbgdb) - 1
							_acdg.emit(_bdcc, _geceg[_aebd:_cbgdb])
						}
					case 3:
						{
							_ccca = (_cbgdb) - 1
							_acdg.emit(_adab, _geceg[_aebd:_cbgdb])
						}
					case 4:
						{
							_ccca = (_cbgdb) - 1
							_acdg.emit(_eegb, _geceg[_aebd:_cbgdb])
						}
					case 11:
						{
							_ccca = (_cbgdb) - 1
							_acdg.emit(_edge, _geceg[_aebd:_cbgdb])
						}
					case 14:
						{
							_ccca = (_cbgdb) - 1
							_acdg.emit(_dcfd, _geceg[_aebd:_cbgdb])
						}
					case 15:
						{
							_ccca = (_cbgdb) - 1
							_acdg.emit(_eaea, _geceg[_aebd+1:_cbgdb-1])
						}
					}
				}
			}
			goto _dggg
		_dggg:
			_gcff = int(_egbe[_defgc])
			_accce = uint(_agbgdb[_gcff])
			_gcff++
			for ; _accce > 0; _accce-- {
				_gcff++
				switch _agbgdb[_gcff-1] {
				case 0:
					_aebd = 0
				case 1:
					_cabb = 0
				}
			}
			if _defgc == 0 {
				goto _cgbg
			}
			if _ccca++; _ccca != _agfa {
				goto _beee
			}
		_bbcdb:
			{
			}
			if _ccca == _ecdbg {
				switch _defgc {
				case 1:
					goto _cace
				case 2:
					goto _cace
				case 31:
					goto _eceec
				case 14:
					goto _cace
				case 15:
					goto _cace
				case 32:
					goto _cacde
				case 17:
					goto _cace
				case 33:
					goto _fdcga
				case 18:
					goto _cace
				case 19:
					goto _cace
				case 34:
					goto _gacef
				case 35:
					goto _fbcgf
				case 36:
					goto _fbcgf
				case 23:
					goto _eadca
				case 37:
					goto _fbcgf
				case 38:
					goto _fbgcd
				case 39:
					goto _gdae
				case 40:
					goto _aedbc
				case 41:
					goto _aedbc
				case 42:
					goto _aedbc
				case 43:
					goto _fdcga
				case 44:
					goto _cace
				case 45:
					goto _aedbc
				case 46:
					goto _aedbc
				case 47:
					goto _aedbc
				case 48:
					goto _aedbc
				case 49:
					goto _aedbc
				case 50:
					goto _aedbc
				case 51:
					goto _aedbc
				case 52:
					goto _aedbc
				case 53:
					goto _aedbc
				case 54:
					goto _aedbc
				case 55:
					goto _aedbc
				case 56:
					goto _aedbc
				case 57:
					goto _aedbc
				case 58:
					goto _aedbc
				case 59:
					goto _cace
				}
			}
		_cgbg:
			{
			}
		}
		if _aebd > 0 {
			copy(_geceg[0:], _geceg[_aebd:])
		}
	}
	_ = _ecdbg
	if _defgc == _gcgg {
		_acdg.emit(_efcc, nil)
	}
	close(_acdg._acbfe)
}

// RoundDown is an implementation of the Excel ROUNDDOWN function that rounds a number
// down to a specified number of digits.
func RoundDown(args []Result) Result { return _ceacf(args, _dgacc) }

// Atan2 implements the Excel ATAN2 function.  It accepts two numeric arguments,
// and the arguments are (x,y), reversed from normal to match Excel's behaviour.
func Atan2(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0041\u0054\u0041\u004e2\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077o\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	_cdgcc := args[0].AsNumber()
	_agage := args[1].AsNumber()
	if _cdgcc.Type == ResultTypeNumber && _agage.Type == ResultTypeNumber {
		_gcebb := _cc.Atan2(_agage.ValueNumber, _cdgcc.ValueNumber)
		if _gcebb != _gcebb {
			return MakeErrorResult("\u0041T\u0041N\u0032\u0020\u0072\u0065\u0074u\u0072\u006ee\u0064\u0020\u004e\u0061\u004e")
		}
		return MakeNumberResult(_gcebb)
	}
	for _, _aada := range []ResultType{_cdgcc.Type, _agage.Type} {
		switch _aada {
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult("\u0041\u0054\u0041\u004e\u0032\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		case ResultTypeError:
			return _cdgcc
		default:
			return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0068an\u0064\u006c\u0065\u0064\u0020\u0041\u0054\u0041\u004e2\u0028)\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _aada))
		}
	}
	return MakeErrorResult("u\u006e\u0068\u0061\u006e\u0064\u006ce\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0066o\u0072\u0020\u0041T\u0041N\u0032\u0028\u0029")
}
func _aabf(_fdcgd, _edff, _agea float64) float64 { return (_fdcgd*3600 + _edff*60 + _agea) / 86400 }

const _faf = "\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u003a\u0028\u0028\u005b\u0030-\u0039]\u0029\u002b\u0029\u0028\u0020\u0028\u0061\u006d\u007c\u0070\u006d\u0029\u0029\u003f"

func _ebee(_ebda string, _acgbeg []Result) (*parsedReplaceObject, Result) {
	if len(_acgbeg) != 4 {
		return nil, MakeErrorResult(_ebda + "\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _acgbeg[0].Type != ResultTypeString {
		return nil, MakeErrorResult(_ebda + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u0073t\u0072\u0069\u006e\u0067")
	}
	_adaec := _acgbeg[0].ValueString
	if _acgbeg[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_ebda + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e \u0061\u0020\u006eu\u006db\u0065\u0072")
	}
	_ddecf := int(_acgbeg[1].ValueNumber) - 1
	if _acgbeg[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_ebda + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0069r\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_gfebc := int(_acgbeg[2].ValueNumber)
	if _acgbeg[3].Type != ResultTypeString {
		return nil, MakeErrorResult(_ebda + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0074\u0068\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e \u0061\u0020\u0073t\u0072i\u006e\u0067")
	}
	_dbec := _acgbeg[3].ValueString
	return &parsedReplaceObject{_adaec, _ddecf, _gfebc, _dbec}, _efc
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

// ISERR is an implementation of the Excel ISERR() function.
func IsErr(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0045\u0052\u0052\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeError && args[0].ValueString != "\u0023\u004e\u002f\u0041")
}
func _cdf(_fcef, _bbg int) int {
	if _bbg == 2 && _bac(_fcef) {
		return 29
	} else {
		return _aed[_bbg-1]
	}
}

var _agbgdb []byte = []byte{0, 1, 2, 1, 11, 1, 12, 1, 13, 1, 14, 1, 15, 1, 16, 1, 17, 1, 18, 1, 19, 1, 20, 1, 21, 1, 22, 1, 23, 1, 24, 1, 25, 1, 26, 1, 27, 1, 28, 1, 29, 1, 30, 1, 31, 1, 32, 1, 33, 1, 34, 1, 35, 1, 36, 1, 37, 1, 38, 1, 39, 1, 40, 1, 41, 1, 42, 1, 43, 2, 0, 1, 2, 3, 4, 2, 3, 5, 2, 3, 6, 2, 3, 7, 2, 3, 8, 2, 3, 9, 2, 3, 10}

type couponArgs struct {
	_eac  float64
	_aea  float64
	_dfff int
	_cfeb int
}

func _dcfa(_bdfg []Result, _bfccb bool) Result {
	_gefcf := "\u004d\u0041\u0058"
	if _bfccb {
		_gefcf = "\u004d\u0041\u0058\u0041"
	}
	if len(_bdfg) == 0 {
		return MakeErrorResult(_gefcf + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gbcfg := -_cc.MaxFloat64
	for _, _gadb := range _bdfg {
		switch _gadb.Type {
		case ResultTypeNumber:
			if (_bfccb || !_gadb.IsBoolean) && _gadb.ValueNumber > _gbcfg {
				_gbcfg = _gadb.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			_acfg := _dcfa(_gadb.ListValues(), _bfccb)
			if _acfg.ValueNumber > _gbcfg {
				_gbcfg = _acfg.ValueNumber
			}
		case ResultTypeEmpty:
		case ResultTypeString:
			_ggeea := 0.0
			if _bfccb {
				_ggeea = _gadb.AsNumber().ValueNumber
			}
			if _ggeea > _gbcfg {
				_gbcfg = _ggeea
			}
		default:
			_bc.Log("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020"+_gefcf+"\u0028\u0029\u0020\u0061rg\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _gadb.Type)
		}
	}
	if _gbcfg == -_cc.MaxFloat64 {
		_gbcfg = 0
	}
	return MakeNumberResult(_gbcfg)
}

// Value is an implementation of the Excel VALUE function.
func Value(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0056\u0041\u004c\u0055\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020a\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bggag := args[0]
	if _bggag.Type == ResultTypeNumber {
		return _bggag
	}
	if _bggag.Type == ResultTypeString {
		_bedbd, _efdac := _c.ParseFloat(_bggag.Value(), 64)
		if _efdac == nil {
			return MakeNumberResult(_bedbd)
		}
	}
	return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0056\u0041L\u0055\u0045")
}

// Munit is an implementation of the Excel MUNIT function that returns an
// identity matrix.
func Munit(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004d\u0055\u004eIT\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073 \u006fn\u0065 \u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0069\u006e\u0070\u0075\u0074")
	}
	_egda := args[0].AsNumber()
	if _egda.Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0055\u004eIT\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073 \u006fn\u0065 \u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0069\u006e\u0070\u0075\u0074")
	}
	_befcc := int(_egda.ValueNumber)
	_ebdgb := make([][]Result, 0, _befcc)
	for _ffcgb := 0; _ffcgb < _befcc; _ffcgb++ {
		_dceaeb := make([]Result, _befcc)
		for _aegfc := 0; _aegfc < _befcc; _aegfc++ {
			if _ffcgb == _aegfc {
				_dceaeb[_aegfc] = MakeNumberResult(1.0)
			} else {
				_dceaeb[_aegfc] = MakeNumberResult(0.0)
			}
		}
		_ebdgb = append(_ebdgb, _dceaeb)
	}
	return MakeArrayResult(_ebdgb)
}
func (_gfdec PrefixVerticalRange) verticalRangeReference(_deabf string) string {
	return _d.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _deabf, _gfdec._dcdb, _gfdec._cbcf)
}
func _fbdba(_caegd _ee.Time) _ee.Time {
	_caegd = _caegd.UTC()
	return _ee.Date(_caegd.Year(), _caegd.Month(), _caegd.Day(), _caegd.Hour(), _caegd.Minute(), _caegd.Second(), _caegd.Nanosecond(), _ee.Local)
}

// Eomonth is an implementation of the Excel EOMONTH() function.
func Eomonth(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0045\u004f\u004d\u004f\u004e\u0054\u0048\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
	}
	_fab := args[1].ValueNumber
	_feda := args[0]
	var _dgee float64
	switch _feda.Type {
	case ResultTypeEmpty:
		_dgee = 0
	case ResultTypeNumber:
		_dgee = _feda.ValueNumber
	case ResultTypeString:
		_aafd := DateValue([]Result{args[0]})
		if _aafd.Type == ResultTypeError {
			return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
		}
		_dgee = _aafd.ValueNumber
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
	}
	_aaaf := _dcge(_dgee)
	_gabf := _aaaf.AddDate(0, int(_fab+1), 0)
	_ebae, _acfd, _ := _gabf.Date()
	_bbcg := _bad(_ebae, int(_acfd), 0)
	if _bbcg < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
	}
	if _ebae == 1900 && _acfd == 3 {
		_bbcg--
	}
	return MakeNumberResult(_bbcg)
}
func (_gbbbg *Lexer) Next() *node {
	_gbbbg._fggea.Lock()
	defer _gbbbg._fggea.Unlock()
	if len(_gbbbg._bggb) > 0 {
		_edbcc := _gbbbg._bggb[0]
		_gbbbg._bggb = _gbbbg._bggb[1:]
		return _edbcc
	}
	return _gbbbg.nextRaw()
}

// String returns a string representation for Negate.
func (_efebdg Negate) String() string { return "\u002d" + _efebdg._eefg.String() }
func _ceacf(_cbef []Result, _cedac rmode) Result {
	if len(_cbef) != 2 {
		return MakeErrorResult("\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_dggbg := _cbef[0].AsNumber()
	if _dggbg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_dafd := _cbef[1].AsNumber()
	if _dafd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ebde := _dafd.ValueNumber
	_fdgg := _dggbg.ValueNumber
	_fcbba := 1.0
	if _ebde > 0 {
		_fcbba = _cc.Pow(1/10.0, _ebde)
	} else {
		_fcbba = _cc.Pow(10.0, -_ebde)
	}
	_fdgg, _efdaa := _cc.Modf(_fdgg / _fcbba)
	switch _cedac {
	case _fceda:
		const _fbdd = 0.499999999
		if _efdaa >= _fbdd {
			_fdgg++
		} else if _efdaa <= -_fbdd {
			_fdgg--
		}
	case _dgacc:
	case _ffcc:
		if _efdaa > 0 {
			_fdgg++
		} else if _efdaa < 0 {
			_fdgg--
		}
	}
	return MakeNumberResult(_fdgg * _fcbba)
}

const _daad = 57364

func _afec(_edec, _fbag, _ddfag, _ecfed, _cegc float64) float64 {
	var _afeb float64
	_dbge := _cegc / _ddfag
	if _dbge >= 1 {
		_dbge = 1
		if _ecfed == 1 {
			_afeb = _edec
		} else {
			_afeb = 0
		}
	} else {
		_afeb = _edec * _cc.Pow(1-_dbge, _ecfed-1)
	}
	_egef := _edec * _cc.Pow(1-_dbge, _ecfed)
	var _cebc float64
	if _egef < _fbag {
		_cebc = _afeb - _fbag
	} else {
		_cebc = _afeb - _egef
	}
	if _cebc < 0 {
		_cebc = 0
	}
	return _cebc
}
func _ffbbd(_fdac string) *criteriaRegex {
	_gcdf := &criteriaRegex{}
	if _fdac == "" {
		return _gcdf
	}
	if _fcbea := _dbee.FindStringSubmatch(_fdac); len(_fcbea) > 1 {
		_gcdf._eefb = _deba
		_gcdf._cfcd = _fcbea[1]
	} else if _gcdc := _eebef.FindStringSubmatch(_fdac); len(_gcdc) > 1 {
		_gcdf._eefb = _deba
		_gcdf._cfcd = _gcdc[1]
	} else if _bagf := _ddbcf.FindStringSubmatch(_fdac); len(_bagf) > 1 {
		_gcdf._eefb = _cfcb
		_gcdf._cfcd = _bagf[1]
	} else if _cebee := _adag.FindStringSubmatch(_fdac); len(_cebee) > 1 {
		_gcdf._eefb = _adffd
		_gcdf._cfcd = _cebee[1]
	} else if _edag := _fegdb.FindStringSubmatch(_fdac); len(_edag) > 1 {
		_gcdf._eefb = _bbfb
		_gcdf._cfcd = _edag[1]
	} else if _ecgbc := _ebeaa.FindStringSubmatch(_fdac); len(_ecgbc) > 1 {
		_gcdf._eefb = _aafag
		_gcdf._cfcd = _ecgbc[1]
	}
	return _gcdf
}

// Log implements the Excel LOG function which returns the log of a number. By
// default the result is base 10, however the second argument to the function
// can specify a different base.
func Log(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 o\u006e\u0065\u0020n\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("L\u004f\u0047\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u006d\u0061\u0078i\u006d\u0075\u006d\u0020\u006f\u0066\u0020\u0074\u0077\u006f a\u0072\u0067\u0075m\u0065n\u0074\u0073")
	}
	_dceg := args[0].AsNumber()
	if _dceg.Type != ResultTypeNumber {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 o\u006e\u0065\u0020n\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_ffcd := 10.0
	if len(args) > 1 {
		_bcfga := args[1].AsNumber()
		if _bcfga.Type != ResultTypeNumber {
			return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061rg\u0075\u006d\u0065n\u0074\u0020t\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0065r\u0069\u0063")
		}
		_ffcd = args[1].ValueNumber
	}
	if _dceg.ValueNumber == 0 {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072gu\u006d\u0065\u006et\u0020\u0074o\u0020\u0062\u0065\u0020\u006e\u006fn\u002d\u007ae\u0072\u006f")
	}
	if _ffcd == 0 {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0073e\u0063\u006f\u006e\u0064\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u002d\u007a\u0065\u0072\u006f")
	}
	return MakeNumberResult(_cc.Log(_dceg.ValueNumber) / _cc.Log(_ffcd))
}

// Large implements the Excel LARGE function.
func Large(args []Result) Result { return _cecd(args, true) }

// Npv implements the Excel NPV function.
func Npv(args []Result) Result {
	_aefb := len(args)
	if _aefb < 2 {
		return MakeErrorResult("\u004e\u0050\u0056 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f \u006fr\u0020m\u006f\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020r\u0061\u0074\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fdcbc := args[0].ValueNumber
	if _fdcbc == -1 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}
	_fceff := []float64{}
	for _, _bcgbf := range args[1:] {
		switch _bcgbf.Type {
		case ResultTypeNumber:
			_fceff = append(_fceff, _bcgbf.ValueNumber)
		case ResultTypeArray, ResultTypeList:
			_edcc := _afgc(_bcgbf)
			for _, _cdfc := range _edcc {
				for _, _cafe := range _cdfc {
					if _cafe.Type == ResultTypeNumber && !_cafe.IsBoolean {
						_fceff = append(_fceff, _cafe.ValueNumber)
					}
				}
			}
		}
	}
	_acef := 0.0
	for _cbfb, _eeea := range _fceff {
		_acef += _eeea / _cc.Pow(1+_fdcbc, float64(_cbfb)+1)
	}
	return MakeNumberResult(_acef)
}

// Trim is an implementation of the Excel TRIM function that removes leading,
// trailing and consecutive spaces.
func Trim(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0054\u0052\u0049\u004d\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gdec := args[0].AsString()
	if _gdec.Type != ResultTypeString {
		return MakeErrorResult("\u0054\u0052\u0049\u004d\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ffede := _gg.Buffer{}
	_eaae := false
	_cfgc := false
	_dgfeb := 0
	for _, _cdfcf := range _gdec.ValueString {
		_afbdf := _cdfcf == ' '
		if _afbdf {
			if !_eaae {
				continue
			}
			if !_cfgc {
				_dgfeb++
				_ffede.WriteRune(_cdfcf)
			}
		} else {
			_dgfeb = 0
			_eaae = true
			_ffede.WriteRune(_cdfcf)
		}
		_cfgc = _afbdf
	}
	_ffede.Truncate(_ffede.Len() - _dgfeb)
	return MakeStringResult(_ffede.String())
}

// Tbilleq implements the Excel TBILLEQ function.
func Tbilleq(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0054\u0042\u0049\u004c\u004c\u0045\u0051\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020t\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_aegbb, _accb, _fdbf := _gdge(args[0], args[1], "\u0054B\u0049\u004c\u004c\u0045\u0051")
	if _fdbf.Type == ResultTypeError {
		return _fdbf
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("T\u0042\u0049\u004c\u004c\u0045\u0051\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0064is\u0063\u006f\u0075\u006et\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cdcg := _accb - _aegbb
	if _cdcg > 365 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0045\u0051\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006d\u0061\u0074\u0075\u0072\u0069\u0074\u0079\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eo\u0074\u0020m\u006f\u0072e\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065\u0020\u0079\u0065\u0061r \u0061\u0066\u0074\u0065\u0072\u0020\u0073\u0065\u0074t\u006c\u0065\u006d\u0065\u006e\u0074")
	}
	_afbd := args[2].ValueNumber
	if _afbd <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0045Q\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074 \u0074\u006f\u0020\u0062\u0065 p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult((365 * _afbd) / (360 - _afbd*_cdcg))
}

// LastEvalIsRef returns if last evaluation with the evaluator was a reference.
func (_ead *defEval) LastEvalIsRef() bool { return _ead._efd }

// HLookup implements the HLOOKUP function that returns a matching value from a
// row in an array.
func HLookup(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("\u0048\u004c\u004f\u004f\u004bU\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if len(args) > 4 {
		return MakeErrorResult("\u0048\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020m\u006f\u0073\u0074\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_faec := args[0]
	_gcebd := args[1]
	if _gcebd.Type != ResultTypeArray {
		return MakeErrorResult("\u0048\u004cO\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_fddga := args[2].AsNumber()
	if _fddga.Type != ResultTypeNumber {
		return MakeErrorResult("\u0048\u004cO\u004f\u004b\u0055\u0050 \u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075m\u0065\u0072\u0069\u0063\u0020\u0072\u006f\u0077\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_bdcb := false
	if len(args) == 4 {
		_fffed := args[3].AsNumber()
		if _fffed.Type != ResultTypeNumber {
			return MakeErrorResult("\u0048\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		if _fffed.ValueNumber == 0 {
			_bdcb = true
		}
	}
	_facg := -1
	_fgff := false
	if len(_gcebd.ValueArray) == 0 {
		return MakeErrorResult("\u0048\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020n\u006f\u006e\u002d\u0065\u006d\u0070\u0074\u0079\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_babge := _gcebd.ValueArray[0]
_efaa:
	for _bfcc, _decda := range _babge {
		switch _dbfga(_decda, _faec, false, _bdcb) {
		case _geagf:
			_facg = _bfcc
		case _fbagf:
			_facg = _bfcc
			_fgff = true
			break _efaa
		}
	}
	if _facg == -1 {
		return MakeErrorResultType(ErrorTypeNA, "\u0048\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_cedd := int(_fddga.ValueNumber) - 1
	if _cedd < 0 || _cedd > len(_gcebd.ValueArray) {
		return MakeErrorResult("\u0048L\u004f\u004f\u004b\u0055P\u0020\u0068\u0061\u0064\u0020i\u006ev\u0061l\u0069\u0064\u0020\u0069\u006e\u0064\u0065x")
	}
	_babge = _gcebd.ValueArray[_cedd]
	if _facg < 0 || _facg >= len(_babge) {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0063\u006f\u006c\u0075\u006d\u006e\u0020\u0069\u006e\u0064\u0065\u0078")
	}
	if _fgff || !_bdcb {
		return _babge[_facg]
	}
	return MakeErrorResultType(ErrorTypeNA, "\u0056\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

var _aaf = map[string]int{"\u006aa\u006e\u0075\u0061\u0072\u0079": 1, "\u0066\u0065\u0062\u0072\u0075\u0061\u0072\u0079": 2, "\u006d\u0061\u0072c\u0068": 3, "\u0061\u0070\u0072i\u006c": 4, "\u006d\u0061\u0079": 5, "\u006a\u0075\u006e\u0065": 6, "\u006a\u0075\u006c\u0079": 7, "\u0061\u0075\u0067\u0075\u0073\u0074": 8, "\u0073e\u0070\u0074\u0065\u006d\u0070\u0065r": 9, "\u006fc\u0074\u006f\u0062\u0065\u0072": 10, "\u006e\u006f\u0076\u0065\u006d\u0062\u0065\u0072": 11, "\u0064\u0065\u0063\u0065\u006d\u0062\u0065\u0072": 12, "\u006a\u0061\u006e": 1, "\u0066\u0065\u0062": 2, "\u006d\u0061\u0072": 3, "\u0061\u0070\u0072": 4, "\u006a\u0075\u006e": 6, "\u006a\u0075\u006c": 7, "\u0061\u0075\u0067": 8, "\u0073\u0065\u0070": 9, "\u006f\u0063\u0074": 10, "\u006e\u006f\u0076": 11, "\u0064\u0065\u0063": 12}

// Mduration implements the Excel MDURATION function.
func Mduration(args []Result) Result {
	_efg, _fabf := _cedgd(args, "\u004dD\u0055\u0052\u0041\u0054\u0049\u004fN")
	if _fabf.Type == ResultTypeError {
		return _fabf
	}
	_geca := _efg._bfge
	_gdff := _efg._ddbc
	_cbbd := _efg._fbgfe
	_ecea := _efg._gda
	_gafg := _efg._gfcf
	_gca := _efg._abff
	_gcda := _dcgec(_geca, _gdff, _cbbd, _ecea, _gafg, _gca)
	if _gcda.Type == ResultTypeError {
		return _gcda
	}
	_ceeef := _gcda.ValueNumber / (1.0 + _ecea/_gafg)
	return MakeNumberResult(_ceeef)
}
func _bbcf(_aedb, _eafe _ee.Time) bool {
	_eag := _aedb.Unix()
	_cddd := _eafe.Unix()
	_agac := _aedb.Year()
	_gfad := _fgc(_agac, _ee.March, 1)
	if _bac(_agac) && _eag < _gfad && _cddd >= _gfad {
		return true
	}
	var _acb = _eafe.Year()
	var _adfb = _fgc(_acb, _ee.March, 1)
	return (_bac(_acb) && _cddd >= _adfb && _eag < _adfb)
}

// Decimal is an implementation of the Excel function DECIMAL() that parses a string
// in a given base and returns the numeric result.
func Decimal(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0044\u0045\u0043\u0049\u004d\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069r\u0065s\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_bfeb := args[0].AsString()
	if _bfeb.Type != ResultTypeString {
		return MakeErrorResult("D\u0045\u0043\u0049\u004d\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0066\u0069\u0072\u0073t \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_dddd := args[1].AsNumber()
	if _dddd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0045\u0043\u0049\u004dA\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_edcb := _bfeb.ValueString
	if len(_edcb) > 2 && (_dda.HasPrefix(_edcb, "\u0030\u0078") || _dda.HasPrefix(_edcb, "\u0030\u0058")) {
		_edcb = _edcb[2:]
	}
	_fffcd, _edfd := _c.ParseInt(_edcb, int(_dddd.ValueNumber), 64)
	if _edfd != nil {
		return MakeErrorResult("\u0044\u0045C\u0049\u004d\u0041\u004c\u0028\u0029\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0069\u006e\u0020\u0063\u006f\u006e\u0076\u0065\u0072si\u006f\u006e")
	}
	return MakeNumberResult(float64(_fffcd))
}
func _debf(_ebfd string, _cegfa []Result) (*parsedSearchObject, Result) {
	_eagfb := len(_cegfa)
	if _eagfb != 2 && _eagfb != 3 {
		return nil, MakeErrorResult(_ebfd + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u006fr\u0020t\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_fbcgc := _cegfa[0]
	if _fbcgc.Type == ResultTypeError {
		return nil, _fbcgc
	}
	if _fbcgc.Type != ResultTypeString && _fbcgc.Type != ResultTypeNumber {
		return nil, MakeErrorResult("\u0054\u0068e\u0020\u0066\u0069\u0072s\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u0020\u0073t\u0072\u0069\u006e\u0067")
	}
	_cgdf := _cegfa[1]
	if _cgdf.Type == ResultTypeError {
		return nil, _cgdf
	}
	if _cgdf.Type != ResultTypeString && _cgdf.Type != ResultTypeNumber {
		return nil, MakeErrorResult("\u0054\u0068\u0065\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020a\u0020\u0073t\u0072\u0069\u006e\u0067")
	}
	_eabg := _cgdf.Value()
	_gfff := _fbcgc.Value()
	_befbg := 1
	if _eagfb == 3 && _cegfa[2].Type != ResultTypeEmpty {
		_edfea := _cegfa[2]
		if _edfea.Type != ResultTypeNumber {
			return nil, MakeErrorResult("P\u006f\u0073\u0069\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0061\u0020n\u0075m\u0062\u0065\u0072")
		}
		_befbg = int(_edfea.ValueNumber)
		if _befbg < 1 {
			return nil, MakeErrorResultType(ErrorTypeValue, "\u0050\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0061 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074h\u0061\u006e\u0020\u0030")
		}
		if _befbg > len(_eabg) {
			return nil, MakeErrorResultType(ErrorTypeValue, "\u0050\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0061 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074h\u0061\u006e\u0020\u0030")
		}
	}
	return &parsedSearchObject{_gfff, _eabg, _befbg}, _efc
}

// ISNUMBER is an implementation of the Excel ISNUMBER() function.
func IsNumber(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053NU\u004d\u0042\u0045R\u0028\u0029\u0020\u0061cce\u0070ts\u0020\u0061\u0020\u0073\u0069\u006e\u0067le\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeNumber)
}

type amorArgs struct {
	_adac float64
	_abda float64
	_cde  float64
	_cfg  float64
	_gcde int
	_bgac float64
	_ccg  int
}

func _badf(_gdd Result, _deg, _gcbd string) (float64, Result) {
	var _bdfbf float64
	switch _gdd.Type {
	case ResultTypeNumber:
		_bdfbf = float64(int(_gdd.ValueNumber))
	case ResultTypeString:
		_bge := DateValue([]Result{_gdd})
		if _bge.Type == ResultTypeError {
			return 0, MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020" + _deg + "\u0020\u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u0020" + _gcbd)
		}
		_bdfbf = _bge.ValueNumber
	default:
		return 0, MakeErrorResult("\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020" + _gcbd)
	}
	if _bdfbf < 0 {
		return 0, MakeErrorResultType(ErrorTypeNum, _deg+"\u0020\u0073\u0068ou\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	return _bdfbf, _efc
}

// Fv implements the Excel FV function.
func Fv(args []Result) Result {
	_ggef := len(args)
	if _ggef < 3 || _ggef > 5 {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u00205")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_eebfb := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_fcea := args[1].ValueNumber
	if _fcea != float64(int(_fcea)) {
		return MakeErrorResultType(ErrorTypeNum, "\u0046\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0061\u0079\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gdba := args[2].ValueNumber
	_gacf := 0.0
	if _ggef >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("F\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0070\u0072\u0065\u0073\u0065\u006et \u0076\u0061\u006c\u0075e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_gacf = args[3].ValueNumber
	}
	_gagc := 0
	if _ggef == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_gagc = int(args[4].ValueNumber)
		if _gagc != 0 {
			_gagc = 1
		}
	}
	return MakeNumberResult(_gfec(_eebfb, _fcea, _gdba, _gacf, _gagc))
}

// Base is an implementation of the Excel BASE function that returns a string
// form of an integer in a specified base and of a minimum length with padded
// zeros.
func Base(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0042\u0041\u0053\u0045\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0042\u0041S\u0045\u0028\u0029\u0020a\u006c\u006co\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006fs\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_befc := args[0].AsNumber()
	if _befc.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0042A\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ceebg := args[1].AsNumber()
	if _ceebg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063o\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0042\u0041\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ebed := int(_ceebg.ValueNumber)
	if _ebed < 0 || _ebed > 36 {
		return MakeErrorResult("\u0072\u0061\u0064\u0069\u0078\u0020m\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065 \u0072\u0061\u006e\u0067\u0065\u0020\u005b0\u002c\u0033\u0036\u005d")
	}
	_fcca := 0
	if len(args) > 2 {
		_aeca := args[2].AsNumber()
		if _aeca.Type != ResultTypeNumber {
			return MakeErrorResult("\u0074\u0068\u0069\u0072\u0064 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0042A\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_fcca = int(_aeca.ValueNumber)
	}
	_ecgb := _c.FormatInt(int64(_befc.ValueNumber), _ebed)
	if len(_ecgb) < _fcca {
		_ecgb = _dda.Repeat("\u0030", _fcca-len(_ecgb)) + _ecgb
	}
	return MakeStringResult(_ecgb)
}

// If is an implementation of the Excel IF() function. It takes one, two or
// three arguments.
func If(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0049\u0046\u0020re\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074 \u006ce\u0061s\u0074 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0049\u0046\u0020ac\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0074\u0020m\u006fs\u0074 \u0074h\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_daec := args[0]
	switch _daec.Type {
	case ResultTypeError:
		return _daec
	case ResultTypeNumber:
		if len(args) == 1 {
			return MakeBoolResult(_daec.ValueNumber != 0)
		}
		if _daec.ValueNumber != 0 {
			return args[1]
		}
		if len(args) == 3 {
			return args[2]
		} else {
			return MakeBoolResult(false)
		}
	case ResultTypeList:
		return _geeg(args)
	case ResultTypeArray:
		return _fefcg(args)
	default:
		return MakeErrorResult("\u0049F\u0020\u0069n\u0069\u0074\u0069\u0061l\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u006d\u0075\u0073t \u0062\u0065\u0020n\u0075\u006de\u0072\u0069\u0063\u0020\u006f\u0072 \u0061\u0072r\u0061\u0079")
	}
}

// Or is an implementation of the Excel OR() function and takes a variable
// number of arguments.
func Or(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004f\u0052\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 \u006f\u006e\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ecfeb := false
	for _, _dgdc := range args {
		switch _dgdc.Type {
		case ResultTypeList, ResultTypeArray:
			_cdcgb := Or(_dgdc.ListValues())
			if _cdcgb.Type == ResultTypeError {
				return _cdcgb
			}
			if _cdcgb.ValueNumber != 0 {
				_ecfeb = true
			}
		case ResultTypeNumber:
			if _dgdc.ValueNumber != 0 {
				_ecfeb = true
			}
		case ResultTypeString:
			return MakeErrorResult("\u004f\u0052 \u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u0065\u0020\u006f\u006e\u0020\u0073\u0074\u0072in\u0067\u0073")
		case ResultTypeError:
			return _dgdc
		default:
			return MakeErrorResult("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u004f\u0052")
		}
	}
	return MakeBoolResult(_ecfeb)
}

// Roman is an implementation of the Excel ROMAN function that convers numbers
// to roman numerals in one of 5 formats.
func Roman(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0077\u006f\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_fcda := args[0].AsNumber()
	if _fcda.Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ddbf := 0
	if len(args) > 1 {
		_dedd := args[1]
		if _dedd.Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063")
		}
		_ddbf = int(_dedd.ValueNumber)
		if _ddbf < 0 {
			_ddbf = 0
		} else if _ddbf > 4 {
			_ddbf = 4
		}
	}
	_bgec := _cdadd
	switch _ddbf {
	case 1:
		_bgec = _ggaab
	case 2:
		_bgec = _eagcg
	case 3:
		_bgec = _fcge
	case 4:
		_bgec = _agaa
	}
	_ebdb := _cc.Trunc(_fcda.ValueNumber)
	_faca := _gg.Buffer{}
	for _, _geabd := range _bgec {
		for _ebdb >= _geabd._edbe {
			_faca.WriteString(_geabd._ggaf)
			_ebdb -= _geabd._edbe
		}
	}
	return MakeStringResult(_faca.String())
}
func init() {
	RegisterFunction("\u0043\u0048\u004f\u004f\u0053\u0045", Choose)
	RegisterFunction("\u0043\u004f\u004c\u0055\u004d\u004e", Column)
	RegisterFunction("\u0043O\u004c\u0055\u004d\u004e\u0053", Columns)
	RegisterFunction("\u0049\u004e\u0044E\u0058", Index)
	RegisterFunctionComplex("\u0049\u004e\u0044\u0049\u0052\u0045\u0043\u0054", Indirect)
	RegisterFunctionComplex("\u004f\u0046\u0046\u0053\u0045\u0054", Offset)
	RegisterFunction("\u004d\u0041\u0054C\u0048", Match)
	RegisterFunction("\u0048L\u004f\u004f\u004b\u0055\u0050", HLookup)
	RegisterFunction("\u004c\u0041\u0052G\u0045", Large)
	RegisterFunction("\u004c\u004f\u004f\u004b\u0055\u0050", Lookup)
	RegisterFunction("\u0052\u004f\u0057", Row)
	RegisterFunction("\u0052\u004f\u0057\u0053", Rows)
	RegisterFunction("\u0053\u004d\u0041L\u004c", Small)
	RegisterFunction("\u0056L\u004f\u004f\u004b\u0055\u0050", VLookup)
	RegisterFunction("\u0054R\u0041\u004e\u0053\u0050\u004f\u0053E", Transpose)
}
func _eabc(_cafd string, _cebfd _ee.Time) (_ee.Time, error) {
	_afgag, _, _ddea := _gga.ParseFloat(_cafd, 10, 128, _gga.ToNearestEven)
	if _ddea != nil {
		return _ee.Time{}, _ddea
	}
	_bdgfa := new(_gga.Float)
	_bdgfa.SetUint64(uint64(24 * _ee.Hour))
	_afgag.Mul(_afgag, _bdgfa)
	_deab, _ := _afgag.Uint64()
	_aebf := _cebfd.Add(_ee.Duration(_deab))
	return _fbdba(_aebf), nil
}

const _fdcg = "\u0028\u0028\u005b\u0030\u002d\u0039]\u0029\u002b\u0029:\u0028\u0028\u005b0\u002d\u0039\u005d\u0029\u002b\u0029\u003a\u0028\u0028\u005b0\u002d\u0039\u005d\u0029\u002b(\\\u002e\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u003f\u0029\u0028\u0020\u0028\u0061\u006d\u007c\u0070\u006d\u0029\u0029\u003f"

func _efebc(_fgaa yyLexer) int { return _addab().Parse(_fgaa) }

// Reference returns an invalid reference for FunctionCall.
func (_bbde FunctionCall) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Eval evaluates a horizontal range with prefix returning a list of results or an error.
func (_afggd PrefixHorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	_eagab := _afggd._fdcfe.Reference(ctx, ev)
	switch _eagab.Type {
	case ReferenceTypeSheet:
		if _fgba(_eagab, ctx) {
			return MakeErrorResultType(ErrorTypeName, _d.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _eagab.Value))
		}
		_fage := _afggd.horizontalRangeReference(_eagab.Value)
		if _fagdg, _gaaca := ev.GetFromCache(_fage); _gaaca {
			return _fagdg
		}
		_fbcae := ctx.Sheet(_eagab.Value)
		_gcca, _dcdd := _ccaa(_fbcae, _afggd._ffba, _afggd._cebd)
		_dbeeg := _cfcef(_fbcae, ev, _gcca, _dcdd)
		ev.SetCache(_fage, _dbeeg)
		return _dbeeg
	default:
		return MakeErrorResult(_d.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _eagab.Type))
	}
}

const (
	_ byte = iota
	_deba
	_cfcb
	_adffd
	_bbfb
	_aafag
)

func Trunc(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("T\u0052\u0055\u004e\u0043\u0028\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061t \u006c\u0065\u0061\u0073t\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065ri\u0063\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_fgbb := args[0].AsNumber()
	if _fgbb.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0054\u0052\u0055\u004e\u0043\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_abdff := float64(0)
	if len(args) > 1 {
		_fcefb := args[1].AsNumber()
		if _fcefb.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0054\u0052\u0055\u004e\u0043\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_abdff = _fcefb.ValueNumber
	}
	_ddgb := _fgbb.ValueNumber
	_gcfb := 1.0
	if _abdff >= 0 {
		_gcfb = _cc.Pow(1/10.0, _abdff)
	} else {
		return MakeNumberResult(0)
	}
	_ddgb, _cccgg := _cc.Modf(_ddgb / _gcfb)
	_edgd := 0.99999
	if _cccgg > _edgd {
		_ddgb++
	} else if _cccgg < -_edgd {
		_ddgb--
	}
	_ = _cccgg
	return MakeNumberResult(_ddgb * _gcfb)
}

// Arabic implements the Excel ARABIC function which parses roman numerals.  It
// accepts one numeric argument.
func Arabic(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0041\u0052\u0041\u0042I\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_fbcb := args[0]
	switch _fbcb.Type {
	case ResultTypeNumber, ResultTypeList, ResultTypeEmpty:
		return MakeErrorResult("\u0041\u0052\u0041B\u0049\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	case ResultTypeString:
		_cdcge := 0.0
		_gfdg := 0.0
		for _, _bbfc := range _fbcb.ValueString {
			_agbac := 0.0
			switch _bbfc {
			case 'I':
				_agbac = 1
			case 'V':
				_agbac = 5
			case 'X':
				_agbac = 10
			case 'L':
				_agbac = 50
			case 'C':
				_agbac = 100
			case 'D':
				_agbac = 500
			case 'M':
				_agbac = 1000
			}
			_cdcge += _agbac
			switch {
			case _gfdg == _agbac && (_gfdg == 5 || _gfdg == 50 || _gfdg == 500):
				return MakeErrorResult("i\u006e\u0076\u0061\u006cid\u0020A\u0052\u0041\u0042\u0049\u0043 \u0066\u006f\u0072\u006d\u0061\u0074")
			case 2*_gfdg == _agbac:
				return MakeErrorResult("i\u006e\u0076\u0061\u006cid\u0020A\u0052\u0041\u0042\u0049\u0043 \u0066\u006f\u0072\u006d\u0061\u0074")
			}
			if _gfdg < _agbac {
				_cdcge -= 2 * _gfdg
			}
			_gfdg = _agbac
		}
		return MakeNumberResult(_cdcge)
	case ResultTypeError:
		return _fbcb
	default:
		return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0068an\u0064\u006c\u0065\u0064\u0020\u0041\u0043\u004f\u0053H\u0028)\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _fbcb.Type))
	}
}

// Left implements the Excel LEFT(string,[n]) function which returns the
// leftmost n characters.
func Left(args []Result) Result {
	_acdc := 1
	switch len(args) {
	case 1:
	case 2:
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("\u004c\u0045F\u0054\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075me\u006e\u0074")
		}
		_acdc = int(args[1].ValueNumber)
		if _acdc < 0 {
			return MakeErrorResult("\u004c\u0045\u0046T \u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020n\u0075m\u0062e\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u003e\u003d\u0020\u0030")
		}
		if _acdc == 0 {
			return MakeStringResult("")
		}
	default:
		return MakeErrorResult("\u004c\u0045\u0046T \u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020o\u006ee\u0020o\u0072 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type == ResultTypeList {
		return MakeErrorResult("\u004c\u0045\u0046T\u0020\u0063\u0061\u006e'\u0074\u0020\u0062\u0065\u0020\u0063\u0061l\u006c\u0065\u0064\u0020\u006f\u006e\u0020\u0061\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_bdfeg := args[0].Value()
	if _acdc > len(_bdfeg) {
		return MakeStringResult(_bdfeg)
	}
	return MakeStringResult(_bdfeg[0:_acdc])
}
func _dbfga(_bcbe, _geagc Result, _cegcc, _dea bool) cmpResult {
	_bcbe = _bcbe.AsNumber()
	_geagc = _geagc.AsNumber()
	if _bcbe.Type != _geagc.Type {
		return _eadc
	}
	if _bcbe.Type == ResultTypeNumber {
		if _bcbe.ValueNumber == _geagc.ValueNumber {
			return _fbagf
		}
		if _bcbe.ValueNumber < _geagc.ValueNumber {
			return _geagf
		}
		return _fgdd
	}
	if _bcbe.Type == ResultTypeString {
		_cdec := _bcbe.ValueString
		_cgae := _geagc.ValueString
		if !_cegcc {
			_cdec = _dda.ToLower(_cdec)
			_cgae = _dda.ToLower(_cgae)
		}
		if _dea {
			_aeegf := _ec.Match(_cgae, _cdec)
			if _aeegf {
				return _fbagf
			} else {
				return _fgdd
			}
		}
		return cmpResult(_dda.Compare(_cdec, _cgae))
	}
	if _bcbe.Type == ResultTypeEmpty {
		return _fbagf
	}
	if _bcbe.Type == ResultTypeList {
		if len(_bcbe.ValueList) < len(_geagc.ValueList) {
			return _geagf
		}
		if len(_bcbe.ValueList) > len(_geagc.ValueList) {
			return _fgdd
		}
		for _dccga := range _bcbe.ValueList {
			_geab := _dbfga(_bcbe.ValueList[_dccga], _geagc.ValueList[_dccga], _cegcc, _dea)
			if _geab != _fbagf {
				return _geab
			}
		}
		return _fbagf
	}
	if _bcbe.Type == ResultTypeList {
		if len(_bcbe.ValueArray) < len(_geagc.ValueArray) {
			return _geagf
		}
		if len(_bcbe.ValueArray) > len(_geagc.ValueArray) {
			return _fgdd
		}
		for _dacd := range _bcbe.ValueArray {
			_gbcbc := _bcbe.ValueArray[_dacd]
			_fbde := _bcbe.ValueArray[_dacd]
			if len(_gbcbc) < len(_fbde) {
				return _geagf
			}
			if len(_gbcbc) > len(_fbde) {
				return _fgdd
			}
			for _bgdc := range _gbcbc {
				_fadbg := _dbfga(_gbcbc[_bgdc], _fbde[_bgdc], _cegcc, _dea)
				if _fadbg != _fbagf {
					return _fadbg
				}
			}
		}
		return _fbagf
	}
	return _eadc
}

type durationArgs struct {
	_bfge  float64
	_ddbc  float64
	_fbgfe float64
	_gda   float64
	_gfcf  float64
	_abff  int
}

// Value returns a string version of the result.
func (_ecafg Result) Value() string {
	switch _ecafg.Type {
	case ResultTypeNumber:
		_gebda := _c.FormatFloat(_ecafg.ValueNumber, 'f', -1, 64)
		if len(_gebda) > 12 {
			_egdae := 12
			for _fcaec := _egdae; _fcaec > 0 && _gebda[_fcaec] == '0'; _fcaec-- {
				_egdae--
			}
			_gebda = _gebda[0 : _egdae+1]
		}
		return _gebda
	case ResultTypeError:
		return _ecafg.ValueString
	case ResultTypeString:
		return _ecafg.ValueString
	case ResultTypeList:
		if len(_ecafg.ValueList) == 0 {
			return ""
		}
		return _ecafg.ValueList[0].Value()
	case ResultTypeArray:
		if len(_ecafg.ValueArray) == 0 || len(_ecafg.ValueArray[0]) == 0 {
			return ""
		}
		return _ecafg.ValueArray[0][0].Value()
	case ResultTypeEmpty:
		return ""
	default:
		return "\u0075\u006e\u0068\u0061nd\u006c\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0076\u0061\u006cu\u0065"
	}
}

type ri struct {
	_edbe float64
	_ggaf string
}

// Update updates the horizontal range references after removing a row/column.
func (_bdcef HorizontalRange) Update(q *_a.UpdateQuery) Expression { return _bdcef }
func _ccaa(_dead Context, _fegg, _beac int) (string, string) {
	_afdg := "\u0041" + _c.Itoa(_fegg)
	_fedce := _dead.LastColumn(_fegg, _beac)
	_ebeg := _fedce + _c.Itoa(_beac)
	return _afdg, _ebeg
}

var _cdadd = []ri{{1000, "\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

// String returns a string representation of CellRef.
func (_cac CellRef) String() string { return _cac._ebd }

// Reference returns a string reference value to an expression with prefix.
func (_fccaf PrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	_gcaff := _fccaf._cacbf.Reference(ctx, ev)
	_bffbe := _fccaf._bgagg.Reference(ctx, ev)
	if _gcaff.Type == ReferenceTypeSheet && _bffbe.Type == ReferenceTypeCell {
		return Reference{Type: ReferenceTypeCell, Value: _gcaff.Value + "\u0021" + _bffbe.Value}
	}
	return ReferenceInvalid
}

// Eval evaluates a horizontal range returning a list of results or an error.
func (_gdbaf HorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	_egbb := _gdbaf.horizontalRangeReference()
	if _cdca, _eeca := ev.GetFromCache(_egbb); _eeca {
		return _cdca
	}
	_adcbe, _gbcg := _ccaa(ctx, _gdbaf._cffd, _gdbaf._aaaaa)
	_fabcc := _cfcef(ctx, ev, _adcbe, _gbcg)
	ev.SetCache(_egbb, _fabcc)
	return _fabcc
}

const _dbgb = 2

var _dgge = [...]int{0, -2, 1, 2, 0, 0, 0, 0, 11, 12, 13, 14, 0, 16, 5, 6, 7, 8, 22, 0, 24, 46, 0, 26, 25, 29, 30, 31, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 18, 20, 9, 10, 0, 0, 23, 32, 33, 47, 0, 49, 51, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 0, 17, 0, 0, 15, 27, 0, 48, 53, 4, 19, 21, 28, 50, 52}

func init() {
	_cdac()
	RegisterFunction("\u004e\u0041", NA)
	RegisterFunction("\u0049S\u0042\u004c\u0041\u004e\u004b", IsBlank)
	RegisterFunction("\u0049\u0053\u0045R\u0052", IsErr)
	RegisterFunction("\u0049S\u0045\u0052\u0052\u004f\u0052", IsError)
	RegisterFunction("\u0049\u0053\u0045\u0056\u0045\u004e", IsEven)
	RegisterFunctionComplex("\u005fx\u006cf\u006e\u002e\u0049\u0053\u0046\u004f\u0052\u004d\u0055\u004c\u0041", IsFormula)
	RegisterFunctionComplex("\u004fR\u0047\u002e\u004f\u0050E\u004e\u004f\u0046\u0046\u0049C\u0045.\u0049S\u004c\u0045\u0041\u0050\u0059\u0045\u0041R", IsLeapYear)
	RegisterFunctionComplex("\u0049S\u004c\u004f\u0047\u0049\u0043\u0041L", IsLogical)
	RegisterFunction("\u0049\u0053\u004e\u0041", IsNA)
	RegisterFunction("\u0049S\u004e\u004f\u004e\u0054\u0045\u0058T", IsNonText)
	RegisterFunction("\u0049\u0053\u004e\u0055\u004d\u0042\u0045\u0052", IsNumber)
	RegisterFunction("\u0049\u0053\u004fD\u0044", IsOdd)
	RegisterFunctionComplex("\u0049\u0053\u0052E\u0046", IsRef)
	RegisterFunction("\u0049\u0053\u0054\u0045\u0058\u0054", IsText)
	RegisterFunctionComplex("\u0043\u0045\u004c\u004c", Cell)
}

// ListValues converts an array to a list or returns a lists values. This is used
// for functions that can accept an array, but don't care about ordering to
// reuse the list function logic.
func (_babfd Result) ListValues() []Result {
	if _babfd.Type == ResultTypeArray {
		_gbag := []Result{}
		for _, _eccde := range _babfd.ValueArray {
			for _, _fcab := range _eccde {
				_gbag = append(_gbag, _fcab)
			}
		}
		return _gbag
	}
	if _babfd.Type == ResultTypeList {
		return _babfd.ValueList
	}
	return nil
}

var _dgae int64 = _fgc(1900, _ee.January, 1)

// MinIfs implements the MINIFS function.
func MinIfs(args []Result) Result {
	_fccag := _gadg(args, true, "\u004d\u0049\u004e\u0049\u0046\u0053")
	if _fccag.Type != ResultTypeEmpty {
		return _fccag
	}
	_dgcgb := _bdcae(args[1:])
	_aadaf := _cc.MaxFloat64
	_ffecd := _afgc(args[0])
	for _, _ecagg := range _dgcgb {
		_bdbca := _ffecd[_ecagg._cdag][_ecagg._gcbgb].ValueNumber
		if _aadaf > _bdbca {
			_aadaf = _bdbca
		}
	}
	if _aadaf == _cc.MaxFloat64 {
		_aadaf = 0
	}
	return MakeNumberResult(float64(_aadaf))
}
func (_gece *ivr) SetOffset(col, row uint32) {}

var _eagcg = []ri{{1000, "\u004d"}, {990, "\u0058\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {490, "\u0058\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {99, "\u0049\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}
var (
	_fbee = 0
	_cdaf = false
)

// Eval evaluates and returns an expression with prefix.
func (_bggdf PrefixExpr) Eval(ctx Context, ev Evaluator) Result {
	_fadfg := _bggdf._cacbf.Reference(ctx, ev)
	switch _fadfg.Type {
	case ReferenceTypeSheet:
		if _fgba(_fadfg, ctx) {
			return MakeErrorResultType(ErrorTypeName, _d.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _fadfg.Value))
		}
		_ebga := ctx.Sheet(_fadfg.Value)
		return _bggdf._bgagg.Eval(_ebga, ev)
	default:
		return MakeErrorResult(_d.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _fadfg.Type))
	}
}

// Row implements the Excel ROW function.
func Row(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0052O\u0057\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_cffb := args[0].Ref
	if _cffb.Type != ReferenceTypeCell {
		return MakeErrorResult("\u0052\u004f\u0057\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073 a\u006e\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079p\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065")
	}
	_baec, _gcg := _bb.ParseCellReference(_cffb.Value)
	if _gcg != nil {
		return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _cffb.Value)
	}
	return MakeNumberResult(float64(_baec.RowIdx))
}

// Reference returns a string reference value to a vertical range.
func (_caabc VerticalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeVerticalRange, Value: _caabc.verticalRangeReference()}
}
func _gff(_abcc, _dcgc []float64, _gbgf float64) float64 {
	_ecff := _gbgf + 1
	_egge := _abcc[0]
	_eedc := len(_abcc)
	_caa := _dcgc[0]
	for _fded := 1; _fded < _eedc; _fded++ {
		_egge += _abcc[_fded] / _cc.Pow(_ecff, (_dcgc[_fded]-_caa)/365)
	}
	return _egge
}

// Update updates references in the PrefixRangeExpr after removing a row/column.
func (_gcecc PrefixRangeExpr) Update(q *_a.UpdateQuery) Expression {
	_addaf := _gcecc
	_babggd := _gcecc._afgf.String()
	if _babggd == q.SheetToUpdate {
		_dafbc := *q
		_dafbc.UpdateCurrentSheet = true
		_addaf._cbce = _gcecc._cbce.Update(&_dafbc)
		_addaf._dbdbg = _gcecc._dbdbg.Update(&_dafbc)
	}
	return _addaf
}

// Lower is an implementation of the Excel LOWER function that returns a lower
// case version of a string.
func Lower(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004c\u004f\u0057\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gfgdg := args[0]
	switch _gfgdg.Type {
	case ResultTypeError:
		return _gfgdg
	case ResultTypeNumber, ResultTypeString:
		return _ddfg(args[0])
	case ResultTypeList:
		_ceefe := _gfgdg.ValueList
		_gdfec := []Result{}
		for _, _ffcagg := range _ceefe {
			_fcbfa := _ddfg(_ffcagg)
			if _fcbfa.Type == ResultTypeError {
				return _fcbfa
			}
			_gdfec = append(_gdfec, _fcbfa)
		}
		return MakeListResult(_gdfec)
	case ResultTypeArray:
		_efgdc := _gfgdg.ValueArray
		_edgf := [][]Result{}
		for _, _bcggg := range _efgdc {
			_gebc := []Result{}
			for _, _afcb := range _bcggg {
				_bfca := _ddfg(_afcb)
				if _bfca.Type == ResultTypeError {
					return _bfca
				}
				_gebc = append(_gebc, _bfca)
			}
			_edgf = append(_edgf, _gebc)
		}
		return MakeArrayResult(_edgf)
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u004c\u004fW\u0045\u0052")
	}
}

var _geece []byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func (_caab ResultType) String() string {
	if _caab >= ResultType(len(_dgef)-1) {
		return _d.Sprintf("\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070e\u0028\u0025\u0064\u0029", _caab)
	}
	return _eccca[_dgef[_caab]:_dgef[_caab+1]]
}

// RegisterFunction registers a standard function.
func RegisterFunction(name string, fn Function) {
	_cfbcf.Lock()
	defer _cfbcf.Unlock()
	if _, _fgddd := _feef[name]; _fgddd {
		_bc.Log("\u0064\u0075p\u006c\u0069\u0063\u0061t\u0065\u0020r\u0065\u0067\u0069\u0073\u0074\u0072\u0061\u0074i\u006f\u006e\u0020\u006f\u0066\u0020\u0066\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0025\u0073", name)
	}
	_feef[name] = fn
}

// Averagea implements the AVERAGEA function, AVERAGEA counts cells that contain
// text as a zero where AVERAGE ignores them entirely.
func Averagea(args []Result) Result {
	_fgfg, _daea := _gfcfe(args, true)
	if _daea == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0041\u0056\u0045\u0052AG\u0045\u0020\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0062\u0079\u0020\u007a\u0065r\u006f")
	}
	return MakeNumberResult(_fgfg / _daea)
}

const _facd = _gge + "\u0020\u0028\u0028[0\u002d\u0039\u005d\u0029\u002b\u0029\u002c\u0020\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

func LexReader(r _g.Reader) chan *node { _efeg := NewLexer(); go _efeg.lex(r); return _efeg._acbfe }
func _bce(_bbfd []Result, _cecfa string) (*amorArgs, Result) {
	_dbeb := len(_bbfd)
	if _dbeb != 6 && _dbeb != 7 {
		return nil, MakeErrorResult(_cecfa + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020\u006fr\u0020s\u0065\u0076\u0065\u006e\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _bbfd[0].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_cecfa + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0063\u006fst \u0074o \u0062\u0065\u0020\u006e\u0075\u006d\u0062er\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_cggd := _bbfd[0].ValueNumber
	if _cggd < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _cecfa+"\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0063\u006f\u0073\u0074\u0020\u0074\u006f\u0020\u0062\u0065 \u006e\u006f\u006e\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	_afg, _bgbg := _badf(_bbfd[1], "\u0064\u0061\u0074\u0065\u0020\u0070\u0075\u0072\u0063h\u0061\u0073\u0065\u0064", _cecfa)
	if _bgbg.Type == ResultTypeError {
		return nil, _bgbg
	}
	_bfgfe, _bgbg := _badf(_bbfd[2], "\u0066\u0069\u0072s\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064", _cecfa)
	if _bgbg.Type == ResultTypeError {
		return nil, _bgbg
	}
	if _bfgfe < _afg {
		return nil, MakeErrorResultType(ErrorTypeNum, _cecfa+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074 \u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020l\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061te\u0020\u0070\u0075\u0072\u0063\u0068\u0061\u0073\u0065\u0064")
	}
	if _bbfd[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_cecfa + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006cv\u0061\u0067\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fgbc := _bbfd[3].ValueNumber
	if _fgbc < 0 || _fgbc > _cggd {
		return nil, MakeErrorResultType(ErrorTypeNum, _cecfa+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061g\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065e\u006e\u0020\u0030\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020in\u0069\u0074\u0069\u0061\u006c\u0020\u0063\u006f\u0073\u0074")
	}
	if _bbfd[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_cecfa + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_bgcbc := int(_bbfd[4].ValueNumber)
	if _bgcbc < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _cecfa+" \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0070\u0065\u0072\u0069o\u0064\u0020\u0074\u006f\u0020\u0062\u0065 \u006e\u006f\u006e\u002d\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if _bbfd[5].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_cecfa + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_fgad := _bbfd[5].ValueNumber
	if _fgad < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _cecfa+"\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 d\u0065\u0070\u0072\u0065\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u0020\u0072\u0061t\u0065\u0020t\u006f\u0020\u0062e\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u0020\u0061\u006e\u0064 \u006c\u0065ss\u0020\u0074\u0068a\u006e\u0020\u0030\u002e\u0035")
	}
	_gccg := 0
	if _dbeb == 7 && _bbfd[6].Type != ResultTypeEmpty {
		if _bbfd[6].Type != ResultTypeNumber {
			return nil, MakeErrorResult(_cecfa + "\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020b\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_gccg = int(_bbfd[6].ValueNumber)
		if !_gaff(_gccg) || _gccg == 2 {
			return nil, MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020fo\u0072\u0020"+_cecfa)
		}
	}
	return &amorArgs{_cggd, _afg, _bfgfe, _fgbc, _bgcbc, _fgad, _gccg}, _efc
}

// Db implements the Excel DB function.
func Db(args []Result) Result {
	_decd := len(args)
	if _decd != 4 && _decd != 5 {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0020\u006f\u0072 \u0066\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0063\u006f\u0073\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ggac := args[0].ValueNumber
	if _ggac < 0 {
		return MakeErrorResultType(ErrorTypeNum, "D\u0042\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0063\u006fs\u0074\u0020\u0074\u006f\u0020\u0062\u0065 \u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cddg := args[1].ValueNumber
	if _cddg < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069\u0066\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fccg := args[2].ValueNumber
	if _fccg <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0042\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006c\u0069\u0066\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("D\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_aegb := args[3].ValueNumber
	if _aegb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0042\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u0076\u0065")
	}
	if _aegb-_fccg > 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u0042")
	}
	_afaa := 12.0
	if _decd == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006do\u006e\u0074\u0068\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_afaa = args[4].ValueNumber
		if _afaa < 1 || _afaa > 12 {
			return MakeErrorResultType(ErrorTypeNum, "\u0044B\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u006do\u006e\u0074\u0068\u0020\u0074\u006f\u0020\u0062\u0065 i\u006e\u0020\u0072a\u006e\u0067e\u0020\u006f\u0066\u0020\u0031\u0020a\u006e\u0064 \u0031\u0032")
		}
	}
	if _afaa == 12 && _aegb > _fccg {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u0042")
	}
	if _cddg >= _ggac {
		return MakeNumberResult(0)
	}
	_dbef := 1 - _cc.Pow(_cddg/_ggac, 1/_fccg)
	_dbef = float64(int(_dbef*1000+0.5)) / 1000
	_cdb := _ggac * _dbef * _afaa / 12
	if _aegb == 1 {
		return MakeNumberResult(_cdb)
	}
	_fggd := _cdb
	_dgc := 0.0
	_acba := _fccg
	if _acba > _aegb {
		_acba = _aegb
	}
	for _acbf := 2.0; _acbf <= _acba; _acbf++ {
		_dgc = (_ggac - _fggd) * _dbef
		_fggd += _dgc
	}
	if _aegb > _fccg {
		return MakeNumberResult((_ggac - _fggd) * _dbef * (12 - _afaa) / 12)
	}
	return MakeNumberResult(_dgc)
}

// Accrintm implements the Excel ACCRINTM function.
func Accrintm(args []Result) Result {
	_eeb := len(args)
	if _eeb != 4 && _eeb != 5 {
		return MakeErrorResult("A\u0043\u0043\u0052\u0049\u004e\u0054\u004d\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066o\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065 a\u0072\u0067\u0075m\u0065n\u0074\u0073")
	}
	_bfag, _fccfg := _badf(args[0], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d")
	if _fccfg.Type == ResultTypeError {
		return _fccfg
	}
	_eefa, _fccfg := _badf(args[1], "\u0073e\u0074t\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065", "\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d")
	if _fccfg.Type == ResultTypeError {
		return _fccfg
	}
	if _bfag >= _eefa {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u0073\u0073\u0075\u0065\u0020d\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0065\u0061\u0072\u006c\u0069\u0065r\u0020\u0074\u0068\u0061\u006e\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065n\u0074 \u0064\u0061\u0074\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0041C\u0043\u0052I\u004e\u0054\u004d\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_bgeb := args[2].ValueNumber
	if _bgeb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061t\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0041\u0043\u0043\u0052\u0049\u004e\u0054M\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0070\u0061\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_efcgc := args[3].ValueNumber
	if _efcgc <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0041\u0043C\u0052\u0049\u004e\u0054\u004d \u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0070\u0061\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ecga := 0
	if _eeb == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0041C\u0043\u0052I\u004e\u0054\u004d \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_ecga = int(args[4].ValueNumber)
		if !_gaff(_ecga) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0041\u0043\u0043R\u0049\u004e\u0054\u004d")
		}
	}
	_bdg, _fccfg := _agef(_bfag, _eefa, _ecga)
	if _fccfg.Type == ResultTypeError {
		return _fccfg
	}
	return MakeNumberResult(_efcgc * _bgeb * _bdg)
}
func init() {
	RegisterFunction("\u0041\u004e\u0044", And)
	RegisterFunction("\u0046\u0041\u004cS\u0045", False)
	RegisterFunction("\u0049\u0046", If)
	RegisterFunction("\u0049F\u0045\u0052\u0052\u004f\u0052", IfError)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0049\u0046\u004e\u0041", IfNA)
	RegisterFunction("\u0049\u0046\u0053", Ifs)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0049\u0046S", Ifs)
	RegisterFunction("\u004e\u004f\u0054", Not)
	RegisterFunction("\u004f\u0052", Or)
	RegisterFunction("\u0054\u0052\u0055\u0045", True)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0058\u004fR", Xor)
}
func _fbdc(_aeegc []Result) []float64 {
	_ggbb := make([]float64, 0)
	for _, _fcac := range _aeegc {
		if _fcac.Type == ResultTypeEmpty {
			continue
		}
		_fcac = _fcac.AsNumber()
		switch _fcac.Type {
		case ResultTypeNumber:
			if !_fcac.IsBoolean {
				_ggbb = append(_ggbb, _fcac.ValueNumber)
			}
		case ResultTypeList, ResultTypeArray:
			_ggbb = append(_ggbb, _fbdc(_fcac.ListValues())...)
		case ResultTypeString:
		default:
			_bc.Log("\u0075\u006e\u0068\u0061\u006ed\u006c\u0065\u0064\u0020\u0065\u0078\u0074\u0072\u0061\u0063\u0074\u004e\u0075m\u0062\u0065\u0072\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _fcac.Type)
		}
	}
	return _ggbb
}

var _fcge = []ri{{1000, "\u004d"}, {995, "\u0056\u004d"}, {990, "\u0058\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {495, "\u0056\u0044"}, {490, "\u0058\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {99, "\u0049\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

// Reference returns an invalid reference for Error.
func (_fb Error) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }
func (_fgfbb HorizontalRange) horizontalRangeReference() string {
	return _d.Sprintf("\u0025\u0064\u003a%\u0064", _fgfbb._cffd, _fgfbb._aaaaa)
}

type parsedReplaceObject struct {
	_bcae  string
	_agabb int
	_fffd  int
	_cedf  string
}

// T is an implementation of the Excel T function that returns whether the
// argument is text.
func T(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("T\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_gaac := args[0]
	if _gaac.Type == ResultTypeError || _gaac.Type == ResultTypeString {
		return _gaac
	}
	return _efc
}

// Lookup implements the LOOKUP function that returns a matching value from a
// column, or from the same index in a second column.
func Lookup(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0074\u0020\u006do\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_baga := args[0]
	_degda := args[1]
	if _degda.Type != ResultTypeArray && _degda.Type != ResultTypeList {
		return MakeErrorResult("\u0056\u004cO\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_fbdbcg := _gcee(_degda)
	_efcfg := -1
	for _cebaa, _fcefc := range _fbdbcg {
		if _dbfga(_baga, _fcefc, false, false) == _fbagf {
			_efcfg = _cebaa
		}
	}
	if _efcfg == -1 {
		return MakeErrorResultType(ErrorTypeNA, "\u004c\u004f\u004f\u004bUP\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_ffeda := _fbdbcg
	if len(args) == 3 {
		_ffeda = _gcee(args[2])
	}
	if _efcfg < 0 || _efcfg >= len(_ffeda) {
		return MakeErrorResultType(ErrorTypeNA, "\u004c\u004f\u004f\u004bUP\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	return _ffeda[_efcfg]
}

var _gadgg = [...]int{0, 1, 1, 2, 4, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 3, 1, 3, 1, 3, 1, 3, 1, 2, 1, 1, 1, 3, 4, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 3, 1, 3, 1, 1, 0}

// Minute is an implementation of the Excel MINUTE() function.
func Minute(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004d\u0049\u004e\u0055T\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_dbd := args[0]
	switch _dbd.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(0)
	case ResultTypeNumber:
		_babf := _dcge(_dbd.ValueNumber)
		return MakeNumberResult(float64(_babf.Minute()))
	case ResultTypeString:
		_ffbf := _dda.ToLower(_dbd.ValueString)
		if !_beca(_ffbf) {
			_, _, _, _bfaf, _efag := _gaf(_ffbf)
			if _efag.Type == ResultTypeError {
				_efag.ErrorMessage = "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074s\u0020\u0066\u006f\u0072\u0020\u004d\u0049\u004e\u0055\u0054\u0045"
				return _efag
			}
			if _bfaf {
				return MakeNumberResult(0)
			}
		}
		_, _dbag, _, _, _, _cebb := _eafd(_ffbf)
		if _cebb.Type == ResultTypeError {
			return _cebb
		}
		return MakeNumberResult(float64(_dbag))
	default:
		return MakeErrorResult("\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u004d\u0049NU\u0054\u0045")
	}
}

const _feabb = 57374

// CeilingMath implements _xlfn.CEILING.MATH which rounds numbers to the nearest
// multiple of the second argument, toward or away from zero as specified by the
// third argument.
func CeilingMath(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0043E\u0049\u004cI\u004e\u0047\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006ee\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0043E\u0049\u004cI\u004e\u0047\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u0061l\u006c\u006f\u0077\u0073\u0020\u0061t\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_gcaa := args[0].AsNumber()
	if _gcaa.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002e\u004dA\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061 \u006eu\u006d\u0062\u0065\u0072")
	}
	_abfd := float64(1)
	if _gcaa.ValueNumber < 0 {
		_abfd = -1
	}
	if len(args) > 1 {
		_ggfb := args[1].AsNumber()
		if _ggfb.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f \u0043\u0045\u0049\u004c\u0049\u004e\u0047.\u004d\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_abfd = _ggfb.ValueNumber
	}
	_fega := float64(1)
	if len(args) > 2 {
		_eade := args[2].AsNumber()
		if _eade.Type != ResultTypeNumber {
			return MakeErrorResult("\u0074\u0068\u0069\u0072\u0064\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002e\u004dA\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061 \u006eu\u006d\u0062\u0065\u0072")
		}
		_fega = _eade.ValueNumber
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Ceil(_gcaa.ValueNumber))
	}
	_cfce := _gcaa.ValueNumber
	_cfce, _cefd := _cc.Modf(_cfce / _abfd)
	if _cefd != 0 {
		if _gcaa.ValueNumber > 0 {
			_cfce++
		} else if _fega < 0 {
			_cfce--
		}
	}
	return MakeNumberResult(_cfce * _abfd)
}

type plex struct {
	_bede  chan *node
	_gggcf Expression
	_dccbf string
}

// ISBLANK is an implementation of the Excel ISBLANK() function.
func IsBlank(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("I\u0053\u0042\u004c\u0041\u004e\u004b(\u0029\u0020\u0061\u0063\u0063\u0065p\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeEmpty)
}

// MDeterm is an implementation of the Excel MDETERM which finds the determinant
// of a matrix.
func MDeterm(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004d\u0044\u0045T\u0045\u0052\u004d\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0072\u0061\u0079 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_caebe := args[0]
	if _caebe.Type != ResultTypeArray {
		return MakeErrorResult("\u004d\u0044\u0045T\u0045\u0052\u004d\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0072\u0061\u0079 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_adffb := len(_caebe.ValueArray)
	for _, _acdb := range _caebe.ValueArray {
		if len(_acdb) != _adffb {
			return MakeErrorResult("\u004d\u0044\u0045TE\u0052\u004d\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u0061 \u0073\u0071\u0075\u0061\u0072\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078")
		}
	}
	return MakeNumberResult(_adgcd(_caebe.ValueArray))
}
func _eafd(_gfa string) (int, int, float64, bool, bool, Result) {
	_acec := ""
	_acgf := []string{}
	for _dfeb, _adc := range _cfba {
		_acgf = _adc.FindStringSubmatch(_gfa)
		if len(_acgf) > 1 {
			_acec = _dfeb
			break
		}
	}
	if _acec == "" {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
	}
	_faff := _acgf[1] == ""
	_acgf = _acgf[49:]
	_dbg := len(_acgf)
	_gfg := _acgf[_dbg-1]
	_cbgg := _gfg == "\u0061\u006d"
	_ceac := _gfg == "\u0070\u006d"
	var _cfac, _bfe int
	var _gcce float64
	var _afaf error
	switch _acec {
	case "\u0068\u0068":
		_cfac, _afaf = _c.Atoi(_acgf[0])
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
		_bfe = 0
		_gcce = 0
	case "\u0068\u0068\u003am\u006d":
		_cfac, _afaf = _c.Atoi(_acgf[0])
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
		_bfe, _afaf = _c.Atoi(_acgf[2])
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
		_gcce = 0
	case "\u006d\u006d\u003as\u0073":
		_cfac = 0
		_bfe, _afaf = _c.Atoi(_acgf[0])
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
		_gcce, _afaf = _c.ParseFloat(_acgf[2], 64)
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
	case "\u0068\u0068\u003a\u006d\u006d\u003a\u0073\u0073":
		_cfac, _afaf = _c.Atoi(_acgf[0])
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
		_bfe, _afaf = _c.Atoi(_acgf[2])
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
		_gcce, _afaf = _c.ParseFloat(_acgf[4], 64)
		if _afaf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		}
	}
	if _bfe >= 60 {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
	}
	if _cbgg || _ceac {
		if _cfac > 12 || _gcce >= 60 {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
		} else if _cfac == 12 {
			_cfac = 0
		}
	} else if _cfac >= 24 || _gcce >= 10000 {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _ggf)
	}
	return _cfac, _bfe, _gcce, _ceac, _faff, _efc
}

// Eval evaluates a range returning a list of results or an error.
func (_afcbc Range) Eval(ctx Context, ev Evaluator) Result {
	_gebd := _afcbc._ecca.Reference(ctx, ev)
	_efgfb := _afcbc._ddgga.Reference(ctx, ev)
	_fceec := _fegfd(_gebd, _efgfb)
	if _gebd.Type == ReferenceTypeCell && _efgfb.Type == ReferenceTypeCell {
		if _ccfca, _deac := ev.GetFromCache(_fceec); _deac {
			return _ccfca
		} else {
			_dbbc := _cfcef(ctx, ev, _gebd.Value, _efgfb.Value)
			ev.SetCache(_fceec, _dbbc)
			return _dbbc
		}
	}
	return MakeErrorResult("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072a\u006e\u0067\u0065\u0020" + _fceec)
}
func init() {
	_dcee = _fa.New(_fa.NewSource(_ee.Now().UnixNano()))
	RegisterFunction("\u0041\u0042\u0053", _bfbd("\u0041\u0053\u0049\u004e", _cc.Abs))
	RegisterFunction("\u0041\u0043\u004f\u0053", _bfbd("\u0041\u0053\u0049\u004e", _cc.Acos))
	RegisterFunction("\u0041\u0043\u004fS\u0048", _bfbd("\u0041\u0053\u0049\u004e", _cc.Acosh))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0041\u0043\u004f\u0054", _bfbd("\u0041\u0043\u004f\u0054", func(_fddd float64) float64 { return _cc.Pi/2 - _cc.Atan(_fddd) }))
	RegisterFunction("_\u0078\u006c\u0066\u006e\u002e\u0041\u0043\u004f\u0054\u0048", _bfbd("\u0041\u0043\u004fT\u0048", func(_geagd float64) float64 { return _cc.Atanh(1 / _geagd) }))
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u0041\u0052\u0041\u0042\u0049\u0043", Arabic)
	RegisterFunction("\u0041\u0053\u0049\u004e", _bfbd("\u0041\u0053\u0049\u004e", _cc.Asin))
	RegisterFunction("\u0041\u0053\u0049N\u0048", _bfbd("\u0041\u0053\u0049N\u0048", _cc.Asinh))
	RegisterFunction("\u0041\u0054\u0041\u004e", _bfbd("\u0041\u0054\u0041\u004e", _cc.Atan))
	RegisterFunction("\u0041\u0054\u0041N\u0048", _bfbd("\u0041\u0054\u0041N\u0048", _cc.Atanh))
	RegisterFunction("\u0041\u0054\u0041N\u0032", Atan2)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0042\u0041\u0053\u0045", Base)
	RegisterFunction("\u0043E\u0049\u004c\u0049\u004e\u0047", Ceiling)
	RegisterFunction("\u005fx\u006cf\u006e\u002e\u0043\u0045\u0049L\u0049\u004eG\u002e\u004d\u0041\u0054\u0048", CeilingMath)
	RegisterFunction("_\u0078\u006c\u0066\u006e.C\u0045I\u004c\u0049\u004e\u0047\u002eP\u0052\u0045\u0043\u0049\u0053\u0045", CeilingPrecise)
	RegisterFunction("\u0043\u004f\u004d\u0042\u0049\u004e", Combin)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0043\u004f\u004d\u0042\u0049\u004e\u0041", Combina)
	RegisterFunction("\u0043\u004f\u0053", _bfbd("\u0043\u004f\u0053", _cc.Cos))
	RegisterFunction("\u0043\u004f\u0053\u0048", _bfbd("\u0043\u004f\u0053\u0048", _cc.Cosh))
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0043\u004fT", _afdc("\u0043\u004f\u0054", _cc.Tan))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0043\u004f\u0054\u0048", _afdc("\u0043\u004f\u0054\u0048", _cc.Tanh))
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0043\u0053C", _afdc("\u0043\u0053\u0043", _cc.Sin))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0043\u0053\u0043\u0048", _afdc("\u0043\u0053\u0043", _cc.Sinh))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0044\u0045\u0043\u0049\u004d\u0041\u004c", Decimal)
	RegisterFunction("\u0044E\u0047\u0052\u0045\u0045\u0053", Degrees)
	RegisterFunction("\u0045\u0056\u0045\u004e", Even)
	RegisterFunction("\u0045\u0058\u0050", _bfbd("\u0045\u0058\u0050", _cc.Exp))
	RegisterFunction("\u0046\u0041\u0043\u0054", Fact)
	RegisterFunction("\u0046\u0041\u0043\u0054\u0044\u004f\u0055\u0042\u004c\u0045", FactDouble)
	RegisterFunction("\u0046\u004c\u004fO\u0052", Floor)
	RegisterFunction("\u005f\u0078l\u0066\u006e\u002eF\u004c\u004f\u004f\u0052\u002e\u004d\u0041\u0054\u0048", FloorMath)
	RegisterFunction("\u005f\u0078\u006c\u0066n.\u0046\u004c\u004f\u004f\u0052\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045", FloorPrecise)
	RegisterFunction("\u0047\u0043\u0044", GCD)
	RegisterFunction("\u0049\u004e\u0054", Int)
	RegisterFunction("I\u0053\u004f\u002e\u0043\u0045\u0049\u004c\u0049\u004e\u0047", CeilingPrecise)
	RegisterFunction("\u004c\u0043\u004d", LCM)
	RegisterFunction("\u004c\u004e", _bfbd("\u004c\u004e", _cc.Log))
	RegisterFunction("\u004c\u004f\u0047", Log)
	RegisterFunction("\u004c\u004f\u00471\u0030", _bfbd("\u004c\u004f\u00471\u0030", _cc.Log10))
	RegisterFunction("\u004dD\u0045\u0054\u0045\u0052\u004d", MDeterm)
	RegisterFunction("\u004d\u004f\u0044", Mod)
	RegisterFunction("\u004d\u0052\u004f\u0055\u004e\u0044", Mround)
	RegisterFunction("M\u0055\u004c\u0054\u0049\u004e\u004f\u004d\u0049\u0041\u004c", Multinomial)
	RegisterFunction("_\u0078\u006c\u0066\u006e\u002e\u004d\u0055\u004e\u0049\u0054", Munit)
	RegisterFunction("\u004f\u0044\u0044", Odd)
	RegisterFunction("\u0050\u0049", Pi)
	RegisterFunction("\u0050\u004f\u0057E\u0052", Power)
	RegisterFunction("\u0050R\u004f\u0044\u0055\u0043\u0054", Product)
	RegisterFunction("\u0051\u0055\u004f\u0054\u0049\u0045\u004e\u0054", Quotient)
	RegisterFunction("\u0052A\u0044\u0049\u0041\u004e\u0053", Radians)
	RegisterFunction("\u0052\u0041\u004e\u0044", Rand)
	RegisterFunction("R\u0041\u004e\u0044\u0042\u0045\u0054\u0057\u0045\u0045\u004e", RandBetween)
	RegisterFunction("\u0052\u004f\u004dA\u004e", Roman)
	RegisterFunction("\u0052\u004f\u0055N\u0044", Round)
	RegisterFunction("\u0052O\u0055\u004e\u0044\u0044\u004f\u0057N", RoundDown)
	RegisterFunction("\u0052O\u0055\u004e\u0044\u0055\u0050", RoundUp)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0053\u0045C", _afdc("\u0053\u0045\u0043", _cc.Cos))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0053\u0045\u0043\u0048", _afdc("\u0053\u0045\u0043\u0048", _cc.Cosh))
	RegisterFunction("\u0053E\u0052\u0049\u0045\u0053\u0053\u0055M", SeriesSum)
	RegisterFunction("\u0053\u0049\u0047\u004e", Sign)
	RegisterFunction("\u0053\u0049\u004e", _bfbd("\u0053\u0049\u004e", _cc.Sin))
	RegisterFunction("\u0053\u0049\u004e\u0048", _bfbd("\u0053\u0049\u004e\u0048", _cc.Sinh))
	RegisterFunction("\u0053\u0051\u0052\u0054", _bfbd("\u0053\u0051\u0052\u0054", _cc.Sqrt))
	RegisterFunction("\u0053\u0051\u0052\u0054\u0050\u0049", _bfbd("\u0053\u0051\u0052\u0054\u0050\u0049", func(_gafe float64) float64 { return _cc.Sqrt(_gafe * _cc.Pi) }))
	RegisterFunction("\u0053\u0055\u004d", Sum)
	RegisterFunction("\u0053\u0055\u004dI\u0046", SumIf)
	RegisterFunction("\u0053\u0055\u004d\u0049\u0046\u0053", SumIfs)
	RegisterFunction("\u0053\u0055\u004d\u0050\u0052\u004f\u0044\u0055\u0043\u0054", SumProduct)
	RegisterFunction("\u0053\u0055\u004dS\u0051", SumSquares)
	RegisterFunction("\u0054\u0041\u004e", _bfbd("\u0054\u0041\u004e", _cc.Tan))
	RegisterFunction("\u0054\u0041\u004e\u0048", _bfbd("\u0054\u0041\u004e\u0048", _cc.Tanh))
	RegisterFunction("\u0054\u0052\u0055N\u0043", Trunc)
}

// Eval evaluates and returns the result of a formula.
func (_dfed *defEval) Eval(ctx Context, formula string) Result {
	_dbae := ParseString(formula)
	_gcc := make(chan Result)
	go func() {
		if _dbae == nil {
			_gcc <- MakeErrorResult(_d.Sprintf("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070a\u0072\u0073\u0065\u0020\u0066\u006f\u0072\u006d\u0075\u006ca\u0020\u0025\u0073", formula))
		} else {
			_dfed.checkLastEvalIsRef(ctx, _dbae)
			_gcc <- _dbae.Eval(ctx, _dfed)
		}
	}()
	select {
	case _gcb := <-_gcc:
		return _gcb
	case <-_ee.After(_dffa):
		_bc.Log("\u0055\u006e\u0069\u004ff\u0066\u0069\u0063\u0065\u0020\u0065\u0076\u0061\u006c\u0075a\u0074i\u006f\u006e\u0020\u0074\u0069\u006d\u0065o\u0075\u0074")
		return MakeNumberResult(0)
	}
}

// MakeStringResult constructs a string result.
func MakeStringResult(s string) Result { return Result{Type: ResultTypeString, ValueString: s} }

// Rri implements the Excel RRI function.
func Rri(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0052\u0052\u0049\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065e\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0052I\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_aaafd := args[0].ValueNumber
	if _aaafd <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052R\u0049\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0066\u0020p\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0052\u0049\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073 p\u0072\u0065\u0073\u0065\u006e\u0074 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_begae := args[1].ValueNumber
	if _begae <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052\u0052\u0049\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006et\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("R\u0052\u0049\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0075\u0074\u0075\u0072e \u0076\u0061\u006c\u0075e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_adae := args[2].ValueNumber
	if _adae < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052R\u0049\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020f\u0075\u0074\u0075\u0072\u0065\u0020\u0076\u0061\u006cue\u0020\u0074\u006f \u0062\u0065 \u006e\u006f\u006e\u0020\u006e\u0065g\u0061\u0074i\u0076\u0065")
	}
	return MakeNumberResult(_cc.Pow(_adae/_begae, 1/_aaafd) - 1)
}

// Oddlyield implements the Excel ODDLYIELD function.
func Oddlyield(args []Result) Result {
	if len(args) != 7 && len(args) != 8 {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045L\u0044\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0073\u0065\u0076\u0065\u006e\u0020\u006f\u0072\u0020\u0065\u0069\u0067\u0068\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_aeaa, _fadb, _fdedg := _gdge(args[0], args[1], "\u004fD\u0044\u004c\u0059\u0049\u0045\u004cD")
	if _fdedg.Type == ResultTypeError {
		return _fdedg
	}
	_bafg, _fdedg := _badf(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E")
	if _fdedg.Type == ResultTypeError {
		return _fdedg
	}
	if _bafg >= _aeaa {
		return MakeErrorResultType(ErrorTypeNum, "\u004c\u0061\u0073\u0074\u0020i\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0064\u0061\u0074\u0065\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0062\u0065\u0066\u006f\u0072\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074e")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_bcbb := args[3].ValueNumber
	if _bcbb < 0 {
		return MakeErrorResultType(ErrorTypeNum, "R\u0061\u0074\u0065\u0020\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006fn\u0020\u006e\u0065g\u0061t\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("O\u0044\u0044\u004c\u0059\u0049\u0045\u004c\u0044\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 p\u0072\u0065\u0073\u0065n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff \u0074\u0079p\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_dfdg := args[4].ValueNumber
	if _dfdg <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073h\u006fu\u006c\u0064\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[5].Type != ResultTypeNumber {
		return MakeErrorResult("\u004fD\u0044\u004cY\u0049\u0045\u004c\u0044 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	_bbgf := args[5].ValueNumber
	if _bbgf < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	if args[6].Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045L\u0044\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0072\u0065\u0071\u0075\u0065\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_bgga := float64(int(args[6].ValueNumber))
	if !_gabe(_bgga) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_cgdge := 0
	if len(args) == 8 && args[7].Type != ResultTypeEmpty {
		if args[7].Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0062a\u0073\u0069\u0073\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
		}
		_cgdge = int(args[7].ValueNumber)
		if !_gaff(_cgdge) {
			return MakeErrorResultType(ErrorTypeNum, "I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0062\u0061\u0073\u0069s\u0020\u0076\u0061\u006c\u0075\u0065\u0020f\u006f\u0072\u0020\u004f\u0044\u0044\u004c\u0059\u0049\u0045L\u0044")
		}
	}
	_agbf, _fdedg := _agef(_bafg, _fadb, _cgdge)
	if _fdedg.Type == ResultTypeError {
		return _fdedg
	}
	_agbf *= _bgga
	_begac, _fdedg := _agef(_aeaa, _fadb, _cgdge)
	if _fdedg.Type == ResultTypeError {
		return _fdedg
	}
	_begac *= _bgga
	_adfbb, _fdedg := _agef(_bafg, _aeaa, _cgdge)
	if _fdedg.Type == ResultTypeError {
		return _fdedg
	}
	_adfbb *= _bgga
	_adda := _bbgf + _agbf*100*_bcbb/_bgga
	_adda /= _dfdg + _adfbb*100*_bcbb/_bgga
	_adda--
	_adda *= _bgga / _begac
	return MakeNumberResult(_adda)
}

// Xirr implements the Excel XIRR function.
func Xirr(args []Result) Result {
	_fcedc := len(args)
	if _fcedc != 2 && _fcedc != 3 {
		return MakeErrorResult("\u0058\u0049RR\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s t\u0077o \u006f\u0072\u0020\u0074\u0068\u0072\u0065e \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_eddaf, _bgacd := _fdfc(args[0], args[1], "\u0058\u0049\u0052\u0052")
	if _bgacd.Type == ResultTypeError {
		return _bgacd
	}
	_efbg := _eddaf._cede
	_ddeg := _eddaf._eaff
	_cadb := 0.1
	if _fcedc == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("\u0058\u0049\u0052\u0052\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_cadb = args[2].ValueNumber
		if _cadb <= -1 {
			return MakeErrorResult("\u0058\u0049\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073\u0020\u0074\u006f\u0020\u0062e\u0020\u006d\u006f\u0072\u0065 \u0074\u0068a\u006e\u0020\u002d\u0031")
		}
	}
	return _fdb(_efbg, _ddeg, _cadb)
}
func _cefeb(_adfg Result, _bddc *criteriaParsed) bool {
	_eebee := _dda.ToLower(_adfg.ValueString)
	_daggg := _bddc._cdcgg._eefb
	_bgedc := _bddc._cdcgg._cfcd
	if _daggg == _deba {
		return _eebee == _bgedc || _ec.Match(_bgedc, _eebee)
	}
	if _adfg.Type != ResultTypeEmpty {
		if _eebee == _bddc._dcfga || _ec.Match(_bddc._dcfga, _eebee) {
			return true
		}
		if _, _eccd := _c.ParseFloat(_bgedc, 64); _eccd == nil {
			return false
		}
		switch _daggg {
		case _cfcb:
			return _eebee <= _bgedc
		case _adffd:
			return _eebee >= _bgedc
		case _bbfb:
			return _eebee < _bgedc
		case _aafag:
			return _eebee > _bgedc
		}
	}
	return false
}

const _fg = "(\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u002d" + _gge + "-\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

// False is an implementation of the Excel FALSE() function. It takes no
// arguments.
func False(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u0046A\u004c\u0053\u0045\u0020\u0074\u0061\u006b\u0065\u0073\u0020\u006eo\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	return MakeBoolResult(false)
}

// NewVerticalRange constructs a new full columns range.
func NewVerticalRange(v string) Expression {
	_fcfcg := _dda.Split(v, "\u003a")
	if len(_fcfcg) != 2 {
		return nil
	}
	if _fcfcg[0] > _fcfcg[1] {
		_fcfcg[0], _fcfcg[1] = _fcfcg[1], _fcfcg[0]
	}
	return VerticalRange{_cacbdb: _fcfcg[0], _cddc: _fcfcg[1]}
}

// Date is an implementation of the Excel DATE() function.
func Date(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0041TE\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s t\u0068re\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_aba := int(args[0].ValueNumber)
	if _aba < 0 || _aba >= 10000 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074 \u0064\u0061\u0074\u0065")
	} else if _aba <= 1899 {
		_aba += 1900
	}
	_aaff := _ee.Month(args[1].ValueNumber)
	_aab := int(args[2].ValueNumber)
	_efcg := _fgc(_aba, _aaff, _aab)
	_agc := _fcb(_dgae, _efcg) + 1
	if _agc < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074 \u0064\u0061\u0074\u0065")
	}
	return MakeNumberResult(_agc)
}
func _fdfc(_geed, _effg Result, _ccb string) (*xargs, Result) {
	if _geed.Type != ResultTypeList && _geed.Type != ResultTypeArray {
		return nil, MakeErrorResult(_ccb + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0076\u0061lue\u0073 t\u006f\u0020\u0062\u0065\u0020\u006f\u0066 a\u0072\u0072\u0061\u0079\u0020\u0074\u0079p\u0065")
	}
	_fccef := _afgc(_geed)
	_gdaa := []float64{}
	for _, _ffca := range _fccef {
		for _, _bee := range _ffca {
			if _bee.Type == ResultTypeNumber && !_bee.IsBoolean {
				_gdaa = append(_gdaa, _bee.ValueNumber)
			} else {
				return nil, MakeErrorResult(_ccb + "\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006dbe\u0072\u0073")
			}
		}
	}
	_badc := len(_gdaa)
	if len(_gdaa) < 2 {
		return nil, MakeErrorResultType(ErrorTypeNum, "")
	}
	if _effg.Type != ResultTypeList && _effg.Type != ResultTypeArray {
		return nil, MakeErrorResult(_ccb + " \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0064\u0061\u0074\u0065s\u0020\u0074\u006f\u0020\u0062\u0065\u0020o\u0066\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0074\u0079p\u0065")
	}
	_bgag := _afgc(_effg)
	_eegef := []float64{}
	_bcbab := 0.0
	for _, _fbb := range _bgag {
		for _, _feafa := range _fbb {
			if _feafa.Type == ResultTypeNumber && !_feafa.IsBoolean {
				_dgcg := float64(int(_feafa.ValueNumber))
				if _dgcg < _bcbab {
					return nil, MakeErrorResultType(ErrorTypeNum, _ccb+" \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0061\u0074\u0065\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u0069\u006e\u0020\u0061\u0073\u0063\u0065\u006edi\u006e\u0067\u0020o\u0072d\u0065\u0072")
				}
				_eegef = append(_eegef, _dgcg)
				_bcbab = _dgcg
			} else {
				return nil, MakeErrorResult(_ccb + "\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0064\u0061\u0074\u0065\u0073\u0020t\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0073")
			}
		}
	}
	if len(_eegef) != _badc {
		return nil, MakeErrorResultType(ErrorTypeNum, "")
	}
	return &xargs{_gdaa, _eegef}, MakeEmptyResult()
}

// Ifs is an implementation of the Excel IFS() function.
func Ifs(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("I\u0046\u0053\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061t\u0020\u006c\u0065\u0061\u0073\u0074\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	for _aeae := 0; _aeae < len(args)-1; _aeae += 2 {
		if args[_aeae].ValueNumber == 1 {
			return args[_aeae+1]
		}
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}

var _gbbbb = map[string]FunctionComplex{}

// ResultType is the type of the result
//go:generate stringer -type=ResultType
type ResultType byte

const _ceccg = 57375

func _fbeb(_gaec, _gggc float64, _cbebg, _ebbf int) (float64, Result) {
	_adba, _fedf := _dcge(_gaec), _dcge(_gggc)
	if _fedf.After(_adba) {
		_bdcf := _efbe(_adba, _fedf, _cbebg, _ebbf)
		_eea := (_fedf.Year()-_bdcf.Year())*12 + int(_fedf.Month()) - int(_bdcf.Month())
		return float64(_eea*_cbebg) / 12.0, _efc
	}
	return 0, MakeErrorResultType(ErrorTypeNum, "\u0053\u0065t\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0062\u0065\u0066\u006f\u0072\u0065\u0020\u006d\u0061\u0074\u0075\u0072\u0069\u0074\u0079\u0020\u0064\u0061\u0074\u0065")
}
func (_ccfcg ReferenceType) String() string {
	if _ccfcg >= ReferenceType(len(_bggaf)-1) {
		return _d.Sprintf("\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054\u0079\u0070e\u0028\u0025\u0064\u0029", _ccfcg)
	}
	return _cffac[_bggaf[_ccfcg]:_bggaf[_ccfcg+1]]
}

var _bea = []*_ge.Regexp{}

func _beca(_ebgf string) bool {
	for _, _gcd := range _bea {
		_fcfa := _gcd.FindStringSubmatch(_ebgf)
		if len(_fcfa) > 1 {
			return true
		}
	}
	return false
}

// Day is an implementation of the Excel DAY() function.
func Day(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0044A\u0059\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_geg := args[0]
	switch _geg.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(0)
	case ResultTypeNumber:
		_abcd := _dcge(_geg.ValueNumber)
		return MakeNumberResult(float64(_abcd.Day()))
	case ResultTypeString:
		_fge := _dda.ToLower(_geg.ValueString)
		if !_gedc(_fge) {
			_, _, _, _, _aec, _ddgd := _eafd(_fge)
			if _ddgd.Type == ResultTypeError {
				_ddgd.ErrorMessage = "I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073\u0020\u0066o\u0072 \u0044\u0041\u0059"
				return _ddgd
			}
			if _aec {
				return MakeNumberResult(0)
			}
		}
		_, _, _cbbg, _, _dbe := _gaf(_fge)
		if _dbe.Type == ResultTypeError {
			return _dbe
		}
		return MakeNumberResult(float64(_cbbg))
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072 \u0044\u0041\u0059")
	}
}
func _acad(_cab, _babc, _gccb, _acee float64, _gbbb int) float64 {
	var _efdg float64
	if _cab == 0 {
		_efdg = (_gccb + _acee) / _babc
	} else {
		_eec := _cc.Pow(1+_cab, _babc)
		if _gbbb == 1 {
			_efdg = (_acee*_cab/(_eec-1) + _gccb*_cab/(1-1/_eec)) / (1 + _cab)
		} else {
			_efdg = _acee*_cab/(_eec-1) + _gccb*_cab/(1-1/_eec)
		}
	}
	return -_efdg
}
func _efb(_ebaf, _agfe float64, _eed, _ddgda int) float64 {
	_aagb := _dcge(_ebaf)
	_gdf := _dcge(_agfe)
	if _ddgda == 1 {
		_cdda := _efbe(_aagb, _gdf, _eed, 1)
		_dgac := _cdda.AddDate(0, 12/_eed, 0)
		return _fagd(_cdda, _dgac, _ddgda)
	}
	return float64(_acfa(0, _ddgda)) / float64(_eed)
}

var _dbee, _eebef, _ebeaa, _fegdb, _adag, _ddbcf *_ge.Regexp

func (_ffbd node) String() string {
	return _d.Sprintf("\u007b%\u0073\u0020\u0025\u0073\u007d", _ffbd._eddb, _dbcb(string(_ffbd._aaca)))
}

//go:generate ragel -G2 -Z lexer.rl
//go:generate goimports -w lexer.go
type Lexer struct {
	_acbfe chan *node
	_fggea _fe.Mutex
	_bbddc []chan *node
	_bggb  []*node
}
type rmode byte

// GetFilename returns an empty string for the invalid reference context.
func (_ecec *ivr) GetFilename() string { return "" }
func _bdcae(_fadge []Result) []rangeIndex {
	_aaadfc := []rangeIndex{}
	_fdcda := len(_fadge)
	for _aafbf := 0; _aafbf < _fdcda-1; _aafbf += 2 {
		_gffgg := []rangeIndex{}
		_gddf := _afgc(_fadge[_aafbf])
		_cdecf := _fbea(_fadge[_aafbf+1])
		if _aafbf == 0 {
			for _abbgb, _dfdgf := range _gddf {
				for _dddda, _cbdb := range _dfdgf {
					if _caaf(_cbdb, _cdecf) {
						_gffgg = append(_gffgg, rangeIndex{_abbgb, _dddda})
					}
				}
			}
		} else {
			for _, _fedab := range _aaadfc {
				_cgce := _gddf[_fedab._cdag][_fedab._gcbgb]
				if _caaf(_cgce, _cdecf) {
					_gffgg = append(_gffgg, _fedab)
				}
			}
		}
		if len(_gffgg) == 0 {
			return []rangeIndex{}
		}
		_aaadfc = _gffgg[:]
	}
	return _aaadfc
}
func _ga(_cbc BinOpType, _cg []Result, _ffd Result) Result {
	_dfbe := []Result{}
	switch _ffd.Type {
	case ResultTypeNumber:
		_ceg := _ffd.ValueNumber
		for _cca := range _cg {
			_fc := _cg[_cca].AsNumber()
			if _fc.Type != ResultTypeNumber {
				return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006fp\u0065\u0072a\u0074\u0069\u006f\u006e")
			}
			switch _cbc {
			case BinOpTypePlus:
				_dfbe = append(_dfbe, MakeNumberResult(_fc.ValueNumber+_ceg))
			case BinOpTypeMinus:
				_dfbe = append(_dfbe, MakeNumberResult(_fc.ValueNumber-_ceg))
			case BinOpTypeMult:
				_dfbe = append(_dfbe, MakeNumberResult(_fc.ValueNumber*_ceg))
			case BinOpTypeDiv:
				if _ceg == 0 {
					return MakeErrorResultType(ErrorTypeDivideByZero, "")
				}
				_dfbe = append(_dfbe, MakeNumberResult(_fc.ValueNumber/_ceg))
			case BinOpTypeExp:
				_dfbe = append(_dfbe, MakeNumberResult(_cc.Pow(_fc.ValueNumber, _ceg)))
			case BinOpTypeLT:
				_dfbe = append(_dfbe, MakeBoolResult(_fc.ValueNumber < _ceg))
			case BinOpTypeGT:
				_dfbe = append(_dfbe, MakeBoolResult(_fc.ValueNumber > _ceg))
			case BinOpTypeEQ:
				_dfbe = append(_dfbe, MakeBoolResult(_fc.ValueNumber == _ceg))
			case BinOpTypeLEQ:
				_dfbe = append(_dfbe, MakeBoolResult(_fc.ValueNumber <= _ceg))
			case BinOpTypeGEQ:
				_dfbe = append(_dfbe, MakeBoolResult(_fc.ValueNumber >= _ceg))
			case BinOpTypeNE:
				_dfbe = append(_dfbe, MakeBoolResult(_fc.ValueNumber != _ceg))
			default:
				return MakeErrorResult(_d.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006c\u0069\u0073\u0074\u0020\u0062\u0069\u006e\u0061\u0072\u0079\u0020\u006fp \u0025\u0073", _cbc))
			}
		}
	case ResultTypeString:
		_aa := _ffd.ValueString
		for _gae := range _cg {
			_abd := _cg[_gae].AsString()
			if _abd.Type != ResultTypeString {
				return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006fp\u0065\u0072a\u0074\u0069\u006f\u006e")
			}
			switch _cbc {
			case BinOpTypeLT:
				_dfbe = append(_dfbe, MakeBoolResult(_abd.ValueString < _aa))
			case BinOpTypeGT:
				_dfbe = append(_dfbe, MakeBoolResult(_abd.ValueString > _aa))
			case BinOpTypeEQ:
				_dfbe = append(_dfbe, MakeBoolResult(_abd.ValueString == _aa))
			case BinOpTypeLEQ:
				_dfbe = append(_dfbe, MakeBoolResult(_abd.ValueString <= _aa))
			case BinOpTypeGEQ:
				_dfbe = append(_dfbe, MakeBoolResult(_abd.ValueString >= _aa))
			case BinOpTypeNE:
				_dfbe = append(_dfbe, MakeBoolResult(_abd.ValueString != _aa))
			default:
				return MakeErrorResult(_d.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006c\u0069\u0073\u0074\u0020\u0062\u0069\u006e\u0061\u0072\u0079\u0020\u006fp \u0025\u0073", _cbc))
			}
		}
	default:
		return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069c\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006e-\u0073t\u0072\u0069\u006e\u0067\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u006e\u0061r\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e")
	}
	return MakeListResult(_dfbe)
}

var _gbbd = [...]string{}

func _ab(_bgf BinOpType, _ff, _cbd [][]Result) Result {
	_abb := [][]Result{}
	for _dfe := range _ff {
		_dgd := _fae(_bgf, _ff[_dfe], _cbd[_dfe])
		if _dgd.Type == ResultTypeError {
			return _dgd
		}
		_abb = append(_abb, _dgd.ValueList)
	}
	return MakeArrayResult(_abb)
}
func _bgfdd(_eggb, _bfgad float64) float64 {
	_eggb = _cc.Trunc(_eggb)
	_bfgad = _cc.Trunc(_bfgad)
	if _eggb == 0 && _bfgad == 0 {
		return 0
	}
	return _eggb * _bfgad / _afde(_eggb, _bfgad)
}

// Mod is an implementation of the Excel MOD function which returns the
// remainder after division. It requires two numeric argumnts.
func Mod(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u004d\u004fD(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0074wo\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_efgb := args[0].AsNumber()
	_dccd := args[1].AsNumber()
	if _efgb.Type != ResultTypeNumber || _dccd.Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u004fD(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0074wo\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	if _dccd.ValueNumber == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "M\u004fD\u0028\u0029\u0020\u0064\u0069\u0076\u0069\u0064e\u0020\u0062\u0079\u0020ze\u0072\u006f")
	}
	_dcfb, _agefg := _cc.Modf(_efgb.ValueNumber / _dccd.ValueNumber)
	if _agefg < 0 {
		_dcfb--
	}
	return MakeNumberResult(_efgb.ValueNumber - _dccd.ValueNumber*_dcfb)
}

// Findb is an implementation of the Excel FINDB().
func Findb(ctx Context, ev Evaluator, args []Result) Result {
	if !ctx.IsDBCS() {
		return Find(args)
	}
	_begdf, _afgaf := _debf("\u0046\u0049\u004e\u0044", args)
	if _afgaf.Type != ResultTypeEmpty {
		return _afgaf
	}
	_fbagfc := _begdf._fcgea
	if _fbagfc == "" {
		return MakeNumberResult(1.0)
	}
	_beda := _begdf._bfgd
	_bfdeg := _begdf._bdgb - 1
	_fdaa := 1
	_gceeb := 0
	for _egfbc := range _beda {
		if _egfbc != 0 {
			_abgc := 1
			if _egfbc-_gceeb > 1 {
				_abgc = 2
			}
			_fdaa += _abgc
		}
		if _fdaa > _bfdeg {
			_eaefc := _dda.Index(_beda[_egfbc:], _fbagfc)
			if _eaefc == 0 {
				return MakeNumberResult(float64(_fdaa))
			}
		}
		_gceeb = _egfbc
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}

// ISNONTEXT is an implementation of the Excel ISNONTEXT() function.
func IsNonText(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053N\u004f\u004e\u0054\u0045X\u0054\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073 \u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type != ResultTypeString)
}

// Update updates references in the VerticalRange after removing a row/column.
func (_dcadg VerticalRange) Update(q *_a.UpdateQuery) Expression {
	if q.UpdateType == _a.UpdateActionRemoveColumn {
		_bfafa := _dcadg
		if q.UpdateCurrentSheet {
			_cgbgf := q.ColumnIdx
			_bfafa._cacbdb = _acg(_dcadg._cacbdb, _cgbgf)
			_bfafa._cddc = _acg(_dcadg._cddc, _cgbgf)
		}
		return _bfafa
	}
	return _dcadg
}

// NewEmptyExpr constructs a new empty expression.
func NewEmptyExpr() Expression { return EmptyExpr{} }

// Fvschedule implements the Excel FVSCHEDULE function.
func Fvschedule(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048\u0045D\u0055\u004c\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048E\u0044\u0055\u004c\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0070\u0072\u0069\u006e\u0063\u0069\u0070\u0061\u006c\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_abbgf := args[0].ValueNumber
	switch args[1].Type {
	case ResultTypeNumber:
		return MakeNumberResult(_abbgf * (args[1].ValueNumber + 1))
	case ResultTypeList, ResultTypeArray:
		_eggg := _afgc(args[1])
		for _, _ccfa := range _eggg {
			for _, _eda := range _ccfa {
				if _eda.Type != ResultTypeNumber || _eda.IsBoolean {
					return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048\u0045\u0044\u0055\u004c\u0045\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020r\u0061\u0074\u0065\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0073")
				}
				_abbgf *= 1.0 + _eda.ValueNumber
			}
		}
		return MakeNumberResult(_abbgf)
	default:
		return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048\u0045\u0044\u0055\u004c\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0063\u0068\u0065\u0064\u0075\u006c\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0072\u0020\u0061\u0072\u0072a\u0079\u0020\u0074y\u0070\u0065")
	}
}

// NewNumber constructs a new number expression.
func NewNumber(v string) Expression {
	_bgeg, _dcff := _c.ParseFloat(v, 64)
	if _dcff != nil {
		_bc.Log("e\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020f\u006f\u0072\u006d\u0075\u006c\u0061\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0073\u003a\u0020%\u0073", v, _dcff)
	}
	return Number{_ceacb: _bgeg}
}
func _fdcd(_fggb []Result, _gbcf string) (float64, float64, Result) {
	if len(_fggb) != 2 {
		return 0, 0, MakeErrorResult(_gbcf + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _fggb[0].Type != ResultTypeNumber {
		return 0, 0, MakeErrorResult(_gbcf + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0064\u006f\u006c\u006c\u0061\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_dbed := _fggb[0].ValueNumber
	if _fggb[1].Type != ResultTypeNumber {
		return 0, 0, MakeErrorResult(_gbcf + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_bfaa := float64(int(_fggb[1].ValueNumber))
	if _bfaa < 0 {
		return 0, 0, MakeErrorResultType(ErrorTypeNum, _gbcf+"\u0020r\u0065\u0071u\u0069\u0072\u0065\u0073 \u0066\u0072\u0061c\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062e \u006e\u006f\u006e \u006e\u0065g\u0061\u0074\u0069\u0076\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	return _dbed, _bfaa, _efc
}

// MakeErrorResultType makes an error result of a given type with a specified
// debug message
func MakeErrorResultType(t ErrorType, msg string) Result {
	switch t {
	case ErrorTypeNull:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u0055\u004c\u004c\u0021", ErrorMessage: msg}
	case ErrorTypeValue:
		return Result{Type: ResultTypeError, ValueString: "\u0023V\u0041\u004c\u0055\u0045\u0021", ErrorMessage: msg}
	case ErrorTypeRef:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u0052\u0045F\u0021", ErrorMessage: msg}
	case ErrorTypeName:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u0041\u004d\u0045\u003f", ErrorMessage: msg}
	case ErrorTypeNum:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u0055M\u0021", ErrorMessage: msg}
	case ErrorTypeSpill:
		return Result{Type: ResultTypeError, ValueString: "\u0023S\u0050\u0049\u004c\u004c\u0021", ErrorMessage: msg}
	case ErrorTypeNA:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u002f\u0041", ErrorMessage: msg}
	case ErrorTypeDivideByZero:
		return Result{Type: ResultTypeError, ValueString: "\u0023D\u0049\u0056\u002f\u0030\u0021", ErrorMessage: msg}
	default:
		return Result{Type: ResultTypeError, ValueString: "\u0023V\u0041\u004c\u0055\u0045\u0021", ErrorMessage: msg}
	}
}

// Rand is an implementation of the Excel RAND() function that returns random
// numbers in the range [0,1).
func Rand(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("R\u0041\u004e\u0044\u0028\u0029\u0020a\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u006e\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	return MakeNumberResult(_dcee.Float64())
}

var _gab float64 = 25569.0
var _ecgbe = [...]int{123, -1000, -1000, 74, 163, 103, 163, 163, -1000, -1000, -1000, -1000, 163, -1000, -1000, -1000, -1000, -1000, -12, 106, -1000, -1000, 143, -1000, -1000, -1000, -1000, -1000, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 74, 163, 163, 6, -28, 74, -15, -15, 60, 10, -14, -1000, -1000, -1000, 7, -1000, 74, -15, -15, -23, -23, -1000, -8, -8, -8, -8, -8, -8, -4, 33, -1000, 163, 163, -1000, -1000, 10, -1000, 163, -1000, -28, 74, -1000, -1000, 74}

type evCache struct {
	_ada map[string]Result
	_ddg *_fe.Mutex
}

// GetLabelPrefix returns an empty string for the invalid reference context.
func (_bccga *ivr) GetLabelPrefix(cellRef string) string { return "" }
func _gfde(_eecg Result, _geecdd, _fecaa string) (float64, Result) {
	switch _eecg.Type {
	case ResultTypeEmpty:
		return 0, _efc
	case ResultTypeNumber:
		return _eecg.ValueNumber, _efc
	case ResultTypeString:
		_ffga, _bbbf := _c.ParseFloat(_eecg.ValueString, 64)
		if _bbbf != nil {
			return 0, MakeErrorResult(_fecaa + "\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u006f\u0072\u0020" + _geecdd)
		}
		return _ffga, _efc
	default:
		return 0, MakeErrorResult(_geecdd + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020" + _fecaa + "\u0020t\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0065\u006d\u0070\u0074\u0079")
	}
}

// Eval evaluates and returns a boolean.
func (_acf Bool) Eval(ctx Context, ev Evaluator) Result { return MakeBoolResult(_acf._fd) }

// Number is a nubmer expression.
type Number struct{ _ceacb float64 }

func _gad(_bag, _gbc, _fcgb, _geda int) int {
	if !_bbbac(_geda) {
		return _fcgb
	}
	_ecbf := _fcgb
	_bfgf := _cdf(_bag, _gbc)
	if _ecbf > 30 || _fcgb >= _bfgf || _ecbf >= _bfgf {
		_ecbf = 30
	}
	return _ecbf
}

// MakeEmptyResult is ued when parsing an empty argument.
func MakeEmptyResult() Result { return Result{Type: ResultTypeEmpty} }

// Edate is an implementation of the Excel EDATE() function.
func Edate(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0045\u0044\u0041\u0054E\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077o\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	}
	_gee := args[1].ValueNumber
	_ece := args[0]
	var _fcg float64
	switch _ece.Type {
	case ResultTypeEmpty:
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	case ResultTypeNumber:
		_fcg = _ece.ValueNumber
	case ResultTypeString:
		_cacbd := DateValue([]Result{args[0]})
		if _cacbd.Type == ResultTypeError {
			return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
		}
		_fcg = _cacbd.ValueNumber
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	}
	_cdc := _dcge(_fcg)
	_eff := _cdc.AddDate(0, int(_gee), 0)
	_eeg, _cfag, _geaa := _eff.Date()
	_gac := _bad(_eeg, int(_cfag), _geaa)
	if _gac < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	}
	return MakeNumberResult(_gac)
}

// Rept is an implementation of the Excel REPT function that returns n copies of
// a string.
func Rept(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("R\u0045\u0050\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	_dfdf := args[0].AsString()
	if _dfdf.Type != ResultTypeString {
		return MakeErrorResult("\u0050R\u004f\u0050E\u0052\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u0069\u0072\u0073\u0074\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	_cbdc := args[1].AsNumber()
	if _cbdc.Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052O\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if _cbdc.ValueNumber < 0 {
		return MakeErrorResult("\u0050\u0052\u004fP\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074o\u0020\u0062\u0065\u0020\u003e\u003d\u0020\u0030")
	}
	if _cbdc.ValueNumber == 0 {
		return MakeStringResult("")
	}
	_ebeaf := _gg.Buffer{}
	for _efgc := 0; _efgc < int(_cbdc.ValueNumber); _efgc++ {
		_ebeaf.WriteString(_dfdf.ValueString)
	}
	return MakeStringResult(_ebeaf.String())
}

// ISFORMULA is an implementation of the Excel ISFORMULA() function.
func IsFormula(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053F\u004f\u0052\u004d\u0055L\u0041\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073 \u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_egac := args[0].Ref
	if _egac.Type != ReferenceTypeCell {
		return MakeErrorResult("I\u0053\u0046\u004f\u0052\u004d\u0055\u004c\u0041\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073t\u0020a\u0072\u0067\u0075\u006de\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
	}
	return MakeBoolResult(ctx.HasFormula(_egac.Value))
}

// IfError is an implementation of the Excel IFERROR() function. It takes two arguments.
func IfError(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0049\u0046\u0045\u0052\u0052\u004f\u0052\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeError {
		if args[0].Type == ResultTypeEmpty {
			return MakeNumberResult(0)
		}
		return args[0]
	}
	return args[1]
}

// String returns a string representation of SheetPrefixExpr.
func (_ccgf SheetPrefixExpr) String() string { return _ccgf._daeae }

// Update updates references in the Range after removing a row/column.
func (_edgde Range) Update(q *_a.UpdateQuery) Expression {
	_dgdca := _edgde
	if q.UpdateCurrentSheet {
		_dgdca._ecca = _edgde._ecca.Update(q)
		_dgdca._ddgga = _edgde._ddgga.Update(q)
	}
	return _dgdca
}

// Update returns the same object as updating sheet references does not affect ConstArrayExpr.
func (_fea ConstArrayExpr) Update(q *_a.UpdateQuery) Expression { return _fea }

// ConstArrayExpr is a constant array expression.
type ConstArrayExpr struct{ _cece [][]Expression }

// Indirect is an implementation of the Excel INDIRECT function that returns the
// contents of a cell.
func Indirect(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 && len(args) != 2 {
		return MakeErrorResult("\u0049\u004e\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0072 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_ccfd := args[0].AsString()
	if _ccfd.Type != ResultTypeString {
		return MakeErrorResult("\u0049\u004e\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069r\u0073t\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066 \u0074\u0079\u0070\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	return ctx.Cell(_ccfd.ValueString, ev)
}

const _fbfe = 57378

func _debd(_cgff [][]Result, _cfge int) [][]Result {
	_cbbad := [][]Result{}
	for _dbffa := range _cgff {
		if _dbffa == 0 {
			continue
		}
		_begd := []Result{}
		for _fbge := range _cgff {
			if _fbge == _cfge {
				continue
			}
			_begd = append(_begd, _cgff[_dbffa][_fbge])
		}
		_cbbad = append(_cbbad, _begd)
	}
	return _cbbad
}
func init() {
	_fcbdd()
	RegisterFunction("\u0041V\u0045\u0052\u0041\u0047\u0045", Average)
	RegisterFunction("\u0041\u0056\u0045\u0052\u0041\u0047\u0045\u0041", Averagea)
	RegisterFunction("\u0043\u004f\u0055N\u0054", Count)
	RegisterFunction("\u0043\u004f\u0055\u004e\u0054\u0041", Counta)
	RegisterFunction("\u0043O\u0055\u004e\u0054\u0049\u0046", CountIf)
	RegisterFunction("\u0043\u004f\u0055\u004e\u0054\u0049\u0046\u0053", CountIfs)
	RegisterFunction("\u0043\u004f\u0055\u004e\u0054\u0042\u004c\u0041\u004e\u004b", CountBlank)
	RegisterFunction("\u004d\u0041\u0058", Max)
	RegisterFunction("\u004d\u0041\u0058\u0041", MaxA)
	RegisterFunction("\u004d\u0041\u0058\u0049\u0046\u0053", MaxIfs)
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u004d\u0041\u0058\u0049\u0046\u0053", MaxIfs)
	RegisterFunction("\u004d\u0045\u0044\u0049\u0041\u004e", Median)
	RegisterFunction("\u004d\u0049\u004e", Min)
	RegisterFunction("\u004d\u0049\u004e\u0041", MinA)
	RegisterFunction("\u004d\u0049\u004e\u0049\u0046\u0053", MinIfs)
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u004d\u0049\u004e\u0049\u0046\u0053", MinIfs)
}

// MinA is an implementation of the Excel MINA() function.
func MinA(args []Result) Result { return _cgdb(args, true) }

const _dgdf = 57346

type Reference struct {
	Type  ReferenceType
	Value string
}

func _gaff(_efeb int) bool { return _efeb >= 0 && _efeb <= 4 }
func _geeg(_gdfc []Result) Result {
	_cagef := _gdfc[0].ValueList
	_bfaeb := len(_cagef)
	switch len(_gdfc) {
	case 1:
		_ffcf := []Result{}
		for _, _fgaf := range _cagef {
			_ffcf = append(_ffcf, MakeBoolResult(_fgaf.ValueNumber != 0))
		}
		return MakeListResult(_ffcf)
	case 2:
		_adaa := _gdfc[1]
		switch _adaa.Type {
		case ResultTypeNumber, ResultTypeString, ResultTypeEmpty:
			_defgb := []Result{}
			for _, _eegc := range _cagef {
				var _fgdf Result
				if _eegc.ValueNumber == 0 {
					_fgdf = MakeBoolResult(false)
				} else {
					_fgdf = _adaa
				}
				_defgb = append(_defgb, _fgdf)
			}
			return MakeListResult(_defgb)
		case ResultTypeList:
			_eagea := _fgggg(_adaa, _bfaeb)
			_cgcf := []Result{}
			for _eageg, _cabe := range _cagef {
				var _dfdc Result
				if _cabe.ValueNumber == 0 {
					_dfdc = MakeBoolResult(false)
				} else {
					_dfdc = _eagea[_eageg]
				}
				_cgcf = append(_cgcf, _dfdc)
			}
			return MakeListResult(_cgcf)
		case ResultTypeArray:
			_afba := _ebacd(_adaa, len(_adaa.ValueArray), _bfaeb)
			_fgggc := [][]Result{}
			for _, _fffa := range _afba {
				_cbbca := []Result{}
				for _egae, _deffe := range _cagef {
					var _ceda Result
					if _deffe.ValueNumber == 0 {
						_ceda = MakeBoolResult(false)
					} else {
						_ceda = _fffa[_egae]
					}
					_cbbca = append(_cbbca, _ceda)
				}
				_fgggc = append(_fgggc, _cbbca)
			}
			return MakeArrayResult(_fgggc)
		}
	case 3:
		_begg := _gdfc[1]
		_cdgc := _gdfc[2]
		_dfab := _eggag(_begg)
		_dfbfac := _eggag(_cdgc)
		if _dfab && _dfbfac {
			_cegb := []Result{}
			for _, _fecg := range _cagef {
				var _cdfg Result
				if _fecg.ValueNumber == 0 {
					_cdfg = _cdgc
				} else {
					_cdfg = _begg
				}
				_cegb = append(_cegb, _cdfg)
			}
			return MakeListResult(_cegb)
		}
		if _begg.Type != ResultTypeArray && _cdgc.Type != ResultTypeArray {
			_ccda := _fgggg(_begg, _bfaeb)
			_eggge := _fgggg(_cdgc, _bfaeb)
			_edgb := []Result{}
			for _adccc, _gdefe := range _cagef {
				var _fffcb Result
				if _gdefe.ValueNumber == 0 {
					_fffcb = _eggge[_adccc]
				} else {
					_fffcb = _ccda[_adccc]
				}
				_edgb = append(_edgb, _fffcb)
			}
			return MakeListResult(_edgb)
		}
		_cdcc, _cbfd := len(_begg.ValueArray), len(_cdgc.ValueArray)
		_eacd, _effge := _cdcc, _cbfd
		if _cbfd > _eacd {
			_eacd, _effge = _effge, _eacd
		}
		_eggad := _ebacd(_begg, _eacd, _bfaeb)
		_bdca := _ebacd(_cdgc, _eacd, _bfaeb)
		_dbdd := [][]Result{}
		for _gbdb := 0; _gbdb < _eacd; _gbdb++ {
			_fffg := []Result{}
			for _eddafe, _aabbf := range _cagef {
				var _cffcc Result
				if _aabbf.ValueNumber == 0 {
					if _gbdb < _cbfd {
						_cffcc = _bdca[_gbdb][_eddafe]
					} else {
						_cffcc = MakeErrorResultType(ErrorTypeNA, "")
					}
				} else {
					if _gbdb < _cdcc {
						_cffcc = _eggad[_gbdb][_eddafe]
					} else {
						_cffcc = MakeErrorResultType(ErrorTypeNA, "")
					}
				}
				_fffg = append(_fffg, _cffcc)
			}
			_dbdd = append(_dbdd, _fffg)
		}
		return MakeArrayResult(_dbdd)
	}
	return MakeErrorResult("")
}

const _dcfd = 57351

// NewPrefixHorizontalRange constructs a new full rows range with prefix.
func NewPrefixHorizontalRange(pfx Expression, v string) Expression {
	_bfeec := _dda.Split(v, "\u003a")
	if len(_bfeec) != 2 {
		return nil
	}
	_gdcf, _ := _c.Atoi(_bfeec[0])
	_eccga, _ := _c.Atoi(_bfeec[1])
	if _gdcf > _eccga {
		_gdcf, _eccga = _eccga, _gdcf
	}
	return PrefixHorizontalRange{_fdcfe: pfx, _ffba: _gdcf, _cebd: _eccga}
}

// Nominal implements the Excel NOMINAL function.
func Nominal(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u004e\u004f\u004d\u0049\u004e\u0041\u004c\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("NO\u004d\u0049N\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u006f\u006d\u0069\u006e\u0061\u006c\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072g\u0075m\u0065\u006et")
	}
	_fgca := args[0].ValueNumber
	if _fgca <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u004e\u004fM\u0049\u004e\u0041\u004c\u0020r\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0065\u0066\u0066\u0065\u0063\u0074\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u004f\u004d\u0049\u004e\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006d\u0070\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0070\u0065\u0072i\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fdcb := float64(int(args[1].ValueNumber))
	if _fdcb < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u004e\u004f\u004d\u0049\u004e\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006eum\u0062e\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006d\u0070\u006f\u0075\u006ed\u0069\u006e\u0067\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065 \u0031\u0020\u006f\u0072\u0020\u006d\u006f\u0072\u0065")
	}
	return MakeNumberResult((_cc.Pow(_fgca+1, 1/_fdcb) - 1) * _fdcb)
}
func _eggag(_adcb Result) bool {
	_acbe := _adcb.Type
	return _acbe != ResultTypeArray && _acbe != ResultTypeList
}

// NewNamedRangeRef constructs a new named range reference.
func NewNamedRangeRef(v string) Expression { return NamedRangeRef{_bbfde: v} }

// Coupdaysnc implements the Excel COUPDAYSNC function.
func Coupdaysnc(args []Result) Result {
	_fad, _ebc := _dccb(args, "\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053\u004e\u0043")
	if _ebc.Type == ResultTypeError {
		return _ebc
	}
	return MakeNumberResult(_cfaca(_fad._eac, _fad._aea, _fad._dfff, _fad._cfeb))
}

// Eval evaluates and returns the result of a constant array expression.
func (_ad ConstArrayExpr) Eval(ctx Context, ev Evaluator) Result {
	_bab := [][]Result{}
	for _, _aaa := range _ad._cece {
		_agb := []Result{}
		for _, _cdd := range _aaa {
			_agb = append(_agb, _cdd.Eval(ctx, ev))
		}
		_bab = append(_bab, _agb)
	}
	return MakeArrayResult(_bab)
}
func _dgcc(_aced int) string {
	if _aced >= 0 && _aced < len(_gbbd) {
		if _gbbd[_aced] != "" {
			return _gbbd[_aced]
		}
	}
	return _d.Sprintf("\u0073\u0074\u0061\u0074\u0065\u002d\u0025\u0076", _aced)
}

var _abgf = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}

// NewPrefixVerticalRange constructs a new full columns range with prefix.
func NewPrefixVerticalRange(pfx Expression, v string) Expression {
	_bgee := _dda.Split(v, "\u003a")
	if len(_bgee) != 2 {
		return nil
	}
	if _bgee[0] > _bgee[1] {
		_bgee[0], _bgee[1] = _bgee[1], _bgee[0]
	}
	return PrefixVerticalRange{_cbfc: pfx, _dcdb: _bgee[0], _cbcf: _bgee[1]}
}

// Ppmt implements the Excel PPPMT function.
func Ppmt(args []Result) Result {
	_cfec := len(args)
	if _cfec < 4 || _cfec > 6 {
		return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006ff\u0020\u0066\u006f\u0075\u0072\u0020a\u006e\u0064\u0020s\u0069\u0078")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("P\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_bdad := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0050\u004dT\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_eabe := args[1].ValueNumber
	if _eabe <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "P\u0050\u004d\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020p\u0065\u0072\u0069\u006f\u0064\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069v\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bfed := args[2].ValueNumber
	if _bfed < _eabe {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064s\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u006c\u0065s\u0073\u0020\u0074\u0068\u0061\u006e \u0070\u0065\u0072i\u006f\u0064")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_bgdeg := args[3].ValueNumber
	_cbge := 0.0
	if _cfec >= 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_cbge = args[4].ValueNumber
	}
	_bbdc := 0
	if _cfec == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("P\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_bbdc = int(args[5].ValueNumber)
		if _bbdc != 0 {
			_bbdc = 1
		}
	}
	return MakeNumberResult(_acad(_bdad, _bfed, _bgdeg, _cbge, _bbdc) - _feff(_bdad, _eabe, _bfed, _bgdeg, _cbge, _bbdc))
}

// Even is an implementation of the Excel EVEN() that rounds a number to the
// nearest even integer.
func Even(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0045\u0056\u0045\u004e(\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_gfgdd := args[0].AsNumber()
	if _gfgdd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0045\u0056\u0045N\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_agbgd := _cc.Signbit(_gfgdd.ValueNumber)
	_bdda, _aeaeg := _cc.Modf(_gfgdd.ValueNumber / 2)
	_faddef := _bdda * 2
	if _aeaeg != 0 {
		if !_agbgd {
			_faddef += 2
		} else {
			_faddef -= 2
		}
	}
	return MakeNumberResult(_faddef)
}
func (_ddf BinOpType) String() string {
	if _ddf >= BinOpType(len(_fed)-1) {
		return _d.Sprintf("\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u0028\u0025\u0064\u0029", _ddf)
	}
	return _ggd[_fed[_ddf]:_fed[_ddf+1]]
}

// String is a string expression.
type String struct{ _adfgc string }

func _gaa(_gaefe, _gbdcf, _bdfa, _becf, _cded, _bded float64, _cdgg int) (float64, Result) {
	_afgg := int(_bded)
	_gabfe := _efb(_gaefe, _gbdcf, _afgg, _cdgg)
	_gcdgb := _cfaca(_gaefe, _gbdcf, _afgg, _cdgg) / _gabfe
	_fdga, _ecfdd := _fbeb(_gaefe, _gbdcf, _afgg, _cdgg)
	if _ecfdd.Type == ResultTypeError {
		return 0, _ecfdd
	}
	_accg := _fffe(_gaefe, _gbdcf, _afgg, _cdgg)
	_cagee := _cded / _cc.Pow(1+_becf/_bded, _fdga-1+_gcdgb)
	_cagee -= 100 * _bdfa / _bded * _accg / _gabfe
	_cagea := 100 * _bdfa / _bded
	_cface := 1 + _becf/_bded
	for _ecfc := 0.0; _ecfc < _fdga; _ecfc++ {
		_cagee += _cagea / _cc.Pow(_cface, _ecfc+_gcdgb)
	}
	return _cagee, MakeEmptyResult()
}

// SheetPrefixExpr is a reference to a sheet like Sheet1! (reference to sheet 'Sheet1').
type SheetPrefixExpr struct{ _daeae string }

func _badd(_ggeaf, _aeea, _fddf int) bool {
	if _aeea < 1 || _aeea > 12 {
		return false
	}
	if _fddf < 1 {
		return false
	}
	return _fddf <= _cdf(_ggeaf, _aeea)
}

// Ddb implements the Excel DDB function.
func Ddb(args []Result) Result {
	_cdeg := len(args)
	if _cdeg != 4 && _cdeg != 5 {
		return MakeErrorResult("\u0044\u0044\u0042 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fdeca := args[0].ValueNumber
	if _fdeca < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044B \u0072\u0065\u0071u\u0069\u0072\u0065\u0073 co\u0073t \u0074\u006f\u0020\u0062\u0065\u0020\u006eon\u0020\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gfgb := args[1].ValueNumber
	if _gfgb < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0061\u006c\u0076\u0061\u0067\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gfbc := args[2].ValueNumber
	if _gfbc <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ffc := args[3].ValueNumber
	if _ffc < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0070\u0065\u0072i\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065")
	}
	if _ffc > _gfbc {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0070\u0065\u0072i\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u0044\u0042")
	}
	_cgf := 2.0
	if _cdeg == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_cgf = args[4].ValueNumber
		if _cgf < 0 {
			return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
		}
	}
	return MakeNumberResult(_afec(_fdeca, _gfgb, _gfbc, _ffc, _cgf))
}

// Product is an implementation of the Excel PRODUCT() function.
func Product(args []Result) Result {
	_eegcd := 1.0
	for _, _ebcac := range args {
		_ebcac = _ebcac.AsNumber()
		switch _ebcac.Type {
		case ResultTypeNumber:
			_eegcd *= _ebcac.ValueNumber
		case ResultTypeList, ResultTypeArray:
			_dfcfe := Product(_ebcac.ListValues())
			if _dfcfe.Type != ResultTypeNumber {
				return _dfcfe
			}
			_eegcd *= _dfcfe.ValueNumber
		case ResultTypeString:
		case ResultTypeError:
			return _ebcac
		case ResultTypeEmpty:
		default:
			return MakeErrorResult(_d.Sprintf("\u0075\u006eha\u006e\u0064\u006ce\u0064\u0020\u0050\u0052ODU\u0043T(\u0029\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0020\u0074\u0079\u0070\u0065\u0020%\u0073", _ebcac.Type))
		}
	}
	return MakeNumberResult(_eegcd)
}
func _gadg(_geecg []Result, _acfaa bool, _ccag string) Result {
	var _ffcagc, _aeag string
	if _acfaa {
		_ffcagc = "\u0074\u0068\u0072e\u0065"
		_aeag = "\u006f\u0064\u0064"
	} else {
		_ffcagc = "\u0074\u0077\u006f"
		_aeag = "\u0065\u0076\u0065\u006e"
	}
	_cfcac := len(_geecg)
	if (_acfaa && _cfcac < 3) || (!_acfaa && _cfcac < 2) {
		return MakeErrorResult(_ccag + "\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020" + _ffcagc + " \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0073")
	}
	if (_cfcac/2*2 == _cfcac) == _acfaa {
		return MakeErrorResult(_ccag + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020" + _aeag + " \u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020a\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	_eedf := -1
	_fecef := -1
	for _dafb := 0; _dafb < _cfcac; _dafb += 2 {
		_ccdad := _geecg[_dafb]
		if _ccdad.Type != ResultTypeArray && _ccdad.Type != ResultTypeList {
			return MakeErrorResult(_ccag + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u006e\u0067\u0065\u0073\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065 \u006c\u0069\u0073\u0074\u0020o\u0072\u0020a\u0072\u0072\u0061\u0079")
		}
		_ccaf := _afgc(_ccdad)
		if _fecef == -1 {
			_fecef = len(_ccaf)
			_eedf = len(_ccaf[0])
		} else if len(_ccaf) != _fecef || len(_ccaf[0]) != _eedf {
			return MakeErrorResult(_ccag + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0061l\u006c\u0020\u0072\u0061n\u0067\u0065\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0073\u0069\u007a\u0065")
		}
		if _acfaa && _dafb == 0 {
			_dafb--
		}
	}
	return _efc
}

// HasFormula returns FALSE for the invalid reference context.
func (_dcad *ivr) HasFormula(cellRef string) bool { return false }

// NewConstArrayExpr constructs a new constant array expression with a given data.
func NewConstArrayExpr(data [][]Expression) Expression { return &ConstArrayExpr{_cece: data} }

// Amorlinc implements the Excel AMORLINC function.
func Amorlinc(args []Result) Result {
	_cacd, _cffa := _bce(args, "\u0041\u004d\u004f\u0052\u004c\u0049\u004e\u0043")
	if _cffa.Type == ResultTypeError {
		return _cffa
	}
	_aaae := _cacd._adac
	_fadg := _cacd._abda
	_gefg := _cacd._cde
	_bbfg := _cacd._cfg
	_becb := _cacd._gcde
	_fcbc := _cacd._bgac
	_dadea := _cacd._ccg
	_agefc, _cga := _agef(_fadg, _gefg, _dadea)
	if _cga.Type == ResultTypeError {
		return MakeErrorResult("\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0065\u0073 \u0066o\u0072\u0020\u0041\u004d\u004f\u0052\u004cI\u004e\u0043")
	}
	_gbb := _agefc * _fcbc * _aaae
	if _becb == 0 {
		return MakeNumberResult(_gbb)
	}
	_agfd := _aaae * _fcbc
	_cdad := _aaae - _bbfg
	_aaffd := int((_cdad - _gbb) / _agfd)
	if _becb <= _aaffd {
		return MakeNumberResult(_agfd)
	} else if _becb == _aaffd+1 {
		return MakeNumberResult(_cdad - _agfd*float64(_aaffd) - _gbb)
	} else {
		return MakeNumberResult(0)
	}
}

// Mround is an implementation of the Excel MROUND function.  It is not a
// generic rounding function and has some oddities to match Excel's behavior.
func Mround(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u004d\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0074\u0077o\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_dab := args[0].AsNumber()
	if _dab.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074\u0020\u0061r\u0067\u0075\u006de\u006e\u0074\u0020\u0074o\u0020\u004d\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ddec := float64(1)
	_edbd := args[1].AsNumber()
	if _edbd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073e\u0063\u006fn\u0064\u0020\u0061\u0072g\u0075\u006d\u0065n\u0074\u0020\u0074\u006f\u0020\u004d\u0052\u004f\u0055ND\u0028\u0029\u0020m\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020n\u0075\u006db\u0065\u0072")
	}
	_ddec = _edbd.ValueNumber
	if _ddec < 0 && _dab.ValueNumber > 0 || _ddec > 0 && _dab.ValueNumber < 0 {
		return MakeErrorResult("\u004d\u0052\u004fUN\u0044\u0028\u0029\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020s\u0069g\u006e\u0073\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068")
	}
	_ddeaf := _dab.ValueNumber
	_ddeaf, _fgfcd := _cc.Modf(_ddeaf / _ddec)
	if _cc.Trunc(_fgfcd+0.5) > 0 {
		_ddeaf++
	}
	return MakeNumberResult(_ddeaf * _ddec)
}

// GetLocked returns FALSE for the invalid reference context.
func (_abec *ivr) GetLocked(cellRef string) bool { return false }

var _eafb = [...]struct {
	_agffb int
	_agbag int
	_febb  string
}{}

// String returns a string representation of a vertical range.
func (_dbeed VerticalRange) String() string { return _dbeed.verticalRangeReference() }

// Proper is an implementation of the Excel PROPER function that returns a copy
// of the string with each word capitalized.
func Proper(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("P\u0052\u004f\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006eg \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_bccc := args[0].AsString()
	if _bccc.Type != ResultTypeString {
		return MakeErrorResult("P\u0052\u004f\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006eg \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_adgg := _gg.Buffer{}
	_bdag := false
	for _, _gefeb := range _bccc.ValueString {
		if !_bdag && _dd.IsLetter(_gefeb) {
			_adgg.WriteRune(_dd.ToUpper(_gefeb))
		} else {
			_adgg.WriteRune(_dd.ToLower(_gefeb))
		}
		_bdag = _dd.IsLetter(_gefeb)
	}
	return MakeStringResult(_adgg.String())
}

// Reference returns a string reference value to a range.
func (_edgbe Range) Reference(ctx Context, ev Evaluator) Reference {
	_cfcc := _edgbe._ecca.Reference(ctx, ev)
	_afdb := _edgbe._ddgga.Reference(ctx, ev)
	if _cfcc.Type == ReferenceTypeCell && _afdb.Type == ReferenceTypeCell {
		return MakeRangeReference(_fegfd(_cfcc, _afdb))
	}
	return ReferenceInvalid
}

const _gaacg = 1

func init() {
	RegisterFunction("\u0043\u0048\u0041\u0052", Char)
	RegisterFunction("\u0043\u004c\u0045A\u004e", Clean)
	RegisterFunction("\u0043\u004f\u0044\u0045", Code)
	RegisterFunction("C\u004f\u004e\u0043\u0041\u0054\u0045\u004e\u0041\u0054\u0045", Concat)
	RegisterFunction("\u0043\u004f\u004e\u0043\u0041\u0054", Concat)
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u0043\u004f\u004e\u0043\u0041\u0054", Concat)
	RegisterFunction("\u0045\u0058\u0041C\u0054", Exact)
	RegisterFunction("\u0046\u0049\u004e\u0044", Find)
	RegisterFunctionComplex("\u0046\u0049\u004eD\u0042", Findb)
	RegisterFunction("\u004c\u0045\u0046\u0054", Left)
	RegisterFunction("\u004c\u0045\u0046T\u0042", Left)
	RegisterFunction("\u004c\u0045\u004e", Len)
	RegisterFunction("\u004c\u0045\u004e\u0042", Len)
	RegisterFunction("\u004c\u004f\u0057E\u0052", Lower)
	RegisterFunction("\u004d\u0049\u0044", Mid)
	RegisterFunction("\u0050\u0052\u004f\u0050\u0045\u0052", Proper)
	RegisterFunction("\u0052E\u0050\u004c\u0041\u0043\u0045", Replace)
	RegisterFunction("\u0052\u0045\u0050\u0054", Rept)
	RegisterFunction("\u0052\u0049\u0047H\u0054", Right)
	RegisterFunction("\u0052\u0049\u0047\u0048\u0054\u0042", Right)
	RegisterFunction("\u0053\u0045\u0041\u0052\u0043\u0048", Search)
	RegisterFunctionComplex("\u0053E\u0041\u0052\u0043\u0048\u0042", Searchb)
	RegisterFunction("\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", Substitute)
	RegisterFunction("\u0054", T)
	RegisterFunction("\u0054\u0045\u0058\u0054", Text)
	RegisterFunction("\u0054\u0045\u0058\u0054\u004a\u004f\u0049\u004e", TextJoin)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0054\u0045\u0058T\u004a\u004f\u0049\u004e", TextJoin)
	RegisterFunction("\u0054\u0052\u0049\u004d", Trim)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0055\u004e\u0049\u0043\u0048\u0041\u0052", Char)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0055\u004e\u0049\u0043\u004f\u0044\u0045", Unicode)
	RegisterFunction("\u0055\u0050\u0050E\u0052", Upper)
	RegisterFunction("\u0056\u0041\u004cU\u0045", Value)
}

const _fcaf = 57368

var _efc Result = MakeEmptyResult()

func (_fcfd *yyParserImpl) Parse(yylex yyLexer) int {
	_eddf := _ee.Now()
	var _cebeef int
	var _eaggf yySymType
	var _egacf []yySymType
	_ = _egacf
	_agaba := _fcfd._fbcd[:]
	Nerrs := 0
	Errflag := 0
	_cgcfc := 0
	_fcfd._accfg = -1
	_cgced := -1
	defer func() { _cgcfc = -1; _fcfd._accfg = -1; _cgced = -1 }()
	_fdad := -1
	goto _eabd
_dfgf:
	return 0
_ebbc:
	return 1
_eabd:
	if _bcbdd(_eddf) {
		_bc.Log("\u0050\u0061\u0072\u0073\u0065\u0020\u0074\u0069\u006d\u0065\u006f\u0075\u0074")
		goto _ebbc
	}
	if _fbee >= 4 {
		_d.Printf("\u0063\u0068\u0061\u0072\u0020\u0025\u0076\u0020\u0069n\u0020\u0025\u0076\u000a", _cecbd(_cgced), _dgcc(_cgcfc))
	}
	_fdad++
	if _fdad >= len(_agaba) {
		_ddggg := make([]yySymType, len(_agaba)*2)
		copy(_ddggg, _agaba)
		_agaba = _ddggg
	}
	_agaba[_fdad] = _eaggf
	_agaba[_fdad]._bddb = _cgcfc
_begda:
	if _bcbdd(_eddf) {
		_bc.Log("\u0050\u0061\u0072\u0073\u0065\u0020\u0074\u0069\u006d\u0065\u006f\u0075\u0074")
		goto _ebbc
	}
	_cebeef = _ecgbe[_cgcfc]
	if _cebeef <= _febdc {
		goto _febfa
	}
	if _fcfd._accfg < 0 {
		_fcfd._accfg, _cgced = _dagaa(yylex, &_fcfd._gddfa)
	}
	_cebeef += _cgced
	if _cebeef < 0 || _cebeef >= _efaeb {
		goto _febfa
	}
	_cebeef = _dace[_cebeef]
	if _egbc[_cebeef] == _cgced {
		_fcfd._accfg = -1
		_cgced = -1
		_eaggf = _fcfd._gddfa
		_cgcfc = _cebeef
		if Errflag > 0 {
			Errflag--
		}
		goto _eabd
	}
_febfa:
	if _bcbdd(_eddf) {
		_bc.Log("\u0050\u0061\u0072\u0073\u0065\u0020\u0074\u0069\u006d\u0065\u006f\u0075\u0074")
		goto _ebbc
	}
	_cebeef = _dgge[_cgcfc]
	if _cebeef == -2 {
		if _fcfd._accfg < 0 {
			_fcfd._accfg, _cgced = _dagaa(yylex, &_fcfd._gddfa)
		}
		_cggc := 0
		for {
			if _edde[_cggc+0] == -1 && _edde[_cggc+1] == _cgcfc {
				break
			}
			_cggc += 2
		}
		for _cggc += 2; ; _cggc += 2 {
			_cebeef = _edde[_cggc+0]
			if _cebeef < 0 || _cebeef == _cgced {
				break
			}
		}
		_cebeef = _edde[_cggc+1]
		if _cebeef < 0 {
			goto _dfgf
		}
	}
	if _cebeef == 0 {
		switch Errflag {
		case 0:
			yylex.Error(_cggfe(_cgcfc, _cgced))
			Nerrs++
			if _fbee >= 1 {
				_d.Printf("\u0025\u0073", _dgcc(_cgcfc))
				_d.Printf("\u0020\u0073\u0061\u0077\u0020\u0025\u0073\u000a", _cecbd(_cgced))
			}
			fallthrough
		case 1, 2:
			Errflag = 3
			for _fdad >= 0 {
				_cebeef = _ecgbe[_agaba[_fdad]._bddb] + _dbgb
				if _cebeef >= 0 && _cebeef < _efaeb {
					_cgcfc = _dace[_cebeef]
					if _egbc[_cgcfc] == _dbgb {
						goto _eabd
					}
				}
				if _fbee >= 2 {
					_d.Printf("\u0065\u0072r\u006f\u0072\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u0070\u006f\u0070\u0073\u0020\u0073\u0074\u0061\u0074\u0065 %\u0064\u000a", _agaba[_fdad]._bddb)
				}
				_fdad--
			}
			goto _ebbc
		case 3:
			if _fbee >= 2 {
				_d.Printf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u0064\u0069s\u0063\u0061\u0072d\u0073 \u0025\u0073\u000a", _cecbd(_cgced))
			}
			if _cgced == _gaacg {
				goto _ebbc
			}
			_fcfd._accfg = -1
			_cgced = -1
			goto _begda
		}
	}
	if _fbee >= 2 {
		_d.Printf("\u0072e\u0064u\u0063\u0065\u0020\u0025\u0076 \u0069\u006e:\u000a\u0009\u0025\u0076\u000a", _cebeef, _dgcc(_cgcfc))
	}
	_ffaaf := _cebeef
	_fadf := _fdad
	_ = _fadf
	_fdad -= _gadgg[_cebeef]
	if _fdad+1 >= len(_agaba) {
		_gecf := make([]yySymType, len(_agaba)*2)
		copy(_gecf, _agaba)
		_agaba = _gecf
	}
	_eaggf = _agaba[_fdad+1]
	_cebeef = _accf[_cebeef]
	_fagcd := _aabef[_cebeef]
	_bgef := _fagcd + _agaba[_fdad]._bddb + 1
	if _bgef >= _efaeb {
		_cgcfc = _dace[_fagcd]
	} else {
		_cgcfc = _dace[_bgef]
		if _egbc[_cgcfc] != -_cebeef {
			_cgcfc = _dace[_fagcd]
		}
	}
	switch _ffaaf {
	case 1:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			yylex.(*plex)._gggcf = _eaggf._ddfc
		}
	case 3:
		_egacf = _agaba[_fadf-2 : _fadf+1]
		{
			_eaggf._ddfc = _egacf[2]._ddfc
		}
	case 4:
		_egacf = _agaba[_fadf-4 : _fadf+1]
		{
		}
	case 5:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewBool(_egacf[1]._bdeac._aaca)
		}
	case 6:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewNumber(_egacf[1]._bdeac._aaca)
		}
	case 7:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewString(_egacf[1]._bdeac._aaca)
		}
	case 8:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewError(_egacf[1]._bdeac._aaca)
		}
	case 9:
		_egacf = _agaba[_fadf-2 : _fadf+1]
		{
			_eaggf._ddfc = _egacf[2]._ddfc
		}
	case 10:
		_egacf = _agaba[_fadf-2 : _fadf+1]
		{
			_eaggf._ddfc = NewNegate(_egacf[2]._ddfc)
		}
	case 15:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = _egacf[2]._ddfc
		}
	case 17:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewConstArrayExpr(_egacf[2]._ddga)
		}
	case 18:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddga = append(_eaggf._ddga, _egacf[1]._fedd)
		}
	case 19:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddga = append(_egacf[1]._ddga, _egacf[3]._fedd)
		}
	case 20:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._fedd = append(_eaggf._fedd, _egacf[1]._ddfc)
		}
	case 21:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._fedd = append(_egacf[1]._fedd, _egacf[3]._ddfc)
		}
	case 23:
		_egacf = _agaba[_fadf-2 : _fadf+1]
		{
			_eaggf._ddfc = NewPrefixExpr(_egacf[1]._ddfc, _egacf[2]._ddfc)
		}
	case 25:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewSheetPrefixExpr(_egacf[1]._bdeac._aaca)
		}
	case 26:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewCellRef(_egacf[1]._bdeac._aaca)
		}
	case 27:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewRange(_egacf[1]._ddfc, _egacf[3]._ddfc)
		}
	case 28:
		_egacf = _agaba[_fadf-4 : _fadf+1]
		{
			_eaggf._ddfc = NewPrefixRangeExpr(_egacf[1]._ddfc, _egacf[2]._ddfc, _egacf[4]._ddfc)
		}
	case 29:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewNamedRangeRef(_egacf[1]._bdeac._aaca)
		}
	case 30:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewHorizontalRange(_egacf[1]._bdeac._aaca)
		}
	case 31:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._ddfc = NewVerticalRange(_egacf[1]._bdeac._aaca)
		}
	case 32:
		_egacf = _agaba[_fadf-2 : _fadf+1]
		{
			_eaggf._ddfc = NewPrefixHorizontalRange(_egacf[1]._ddfc, _egacf[2]._bdeac._aaca)
		}
	case 33:
		_egacf = _agaba[_fadf-2 : _fadf+1]
		{
			_eaggf._ddfc = NewPrefixVerticalRange(_egacf[1]._ddfc, _egacf[2]._bdeac._aaca)
		}
	case 34:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypePlus, _egacf[3]._ddfc)
		}
	case 35:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeMinus, _egacf[3]._ddfc)
		}
	case 36:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeMult, _egacf[3]._ddfc)
		}
	case 37:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeDiv, _egacf[3]._ddfc)
		}
	case 38:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeExp, _egacf[3]._ddfc)
		}
	case 39:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeLT, _egacf[3]._ddfc)
		}
	case 40:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeGT, _egacf[3]._ddfc)
		}
	case 41:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeLEQ, _egacf[3]._ddfc)
		}
	case 42:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeGEQ, _egacf[3]._ddfc)
		}
	case 43:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeEQ, _egacf[3]._ddfc)
		}
	case 44:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeNE, _egacf[3]._ddfc)
		}
	case 45:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewBinaryExpr(_egacf[1]._ddfc, BinOpTypeConcat, _egacf[3]._ddfc)
		}
	case 47:
		_egacf = _agaba[_fadf-2 : _fadf+1]
		{
			_eaggf._ddfc = NewFunction(_egacf[1]._bdeac._aaca, nil)
		}
	case 48:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._ddfc = NewFunction(_egacf[1]._bdeac._aaca, _egacf[2]._fedd)
		}
	case 49:
		_egacf = _agaba[_fadf-1 : _fadf+1]
		{
			_eaggf._fedd = append(_eaggf._fedd, _egacf[1]._ddfc)
		}
	case 50:
		_egacf = _agaba[_fadf-3 : _fadf+1]
		{
			_eaggf._fedd = append(_egacf[1]._fedd, _egacf[3]._ddfc)
		}
	case 53:
		_egacf = _agaba[_fadf-0 : _fadf+1]
		{
			_eaggf._ddfc = NewEmptyExpr()
		}
	}
	goto _eabd
}

// FunctionCall is a function call expression.
type FunctionCall struct {
	_ffcfde string
	_gace   []Expression
}

// Concat is an implementation of the Excel CONCAT() and deprecated CONCATENATE() function.
func Concat(args []Result) Result {
	_cgcee := _gg.Buffer{}
	for _, _bbacd := range args {
		switch _bbacd.Type {
		case ResultTypeString:
			_cgcee.WriteString(_bbacd.ValueString)
		case ResultTypeNumber:
			var _fcbcg string
			if _bbacd.IsBoolean {
				if _bbacd.ValueNumber == 0 {
					_fcbcg = "\u0046\u0041\u004cS\u0045"
				} else {
					_fcbcg = "\u0054\u0052\u0055\u0045"
				}
			} else {
				_fcbcg = _bbacd.AsString().ValueString
			}
			_cgcee.WriteString(_fcbcg)
		default:
			return MakeErrorResult("\u0043\u004f\u004e\u0043\u0041T\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073")
		}
	}
	return MakeStringResult(_cgcee.String())
}

// Search is an implementation of the Excel SEARCH().
func Search(args []Result) Result {
	_eccge, _ceab := _debf("\u0046\u0049\u004e\u0044", args)
	if _ceab.Type != ResultTypeEmpty {
		return _ceab
	}
	_eaaf := _dda.ToLower(_eccge._fcgea)
	if _eaaf == "" {
		return MakeNumberResult(1.0)
	}
	_dfedd := _dda.ToLower(_eccge._bfgd)
	_abgd := _eccge._bdgb
	_bbdg := 1
	for _aaga := range _dfedd {
		if _bbdg < _abgd {
			_bbdg++
			continue
		}
		_ebfe := _ec.Index(_eaaf, _dfedd[_aaga:])
		if _ebfe == 0 {
			return MakeNumberResult(float64(_bbdg))
		}
		_bbdg++
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}

var _aabef = [...]int{0, 0, 71, 70, 69, 4, 67, 66, 53, 51, 50, 49, 48, 47, 46, 45, 44, 2}

// Combina is an implementation of the Excel COMBINA function whic returns the
// number of combinations with repetitions.
func Combina(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0043\u004f\u004dB\u0049\u004e\u0041\u0028)\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fafb := args[0].AsNumber()
	_ffg := args[1].AsNumber()
	if _fafb.Type != ResultTypeNumber || _ffg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0043\u004fMB\u0049\u004e\u0041(\u0029\u0020\u0072\u0065qui\u0072es\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_caaa := _cc.Trunc(_fafb.ValueNumber)
	_edegb := _cc.Trunc(_ffg.ValueNumber)
	if _caaa < _edegb {
		return MakeErrorResult("\u0043O\u004d\u0042\u0049\u004e\u0041\u0028\u0029\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u006e\u0020\u003e\u0020\u006b")
	}
	if _caaa == 0 {
		return MakeNumberResult(0)
	}
	args[0] = MakeNumberResult(_caaa + _edegb - 1)
	args[1] = MakeNumberResult(_caaa - 1)
	return Combin(args)
}

// String returns a string representation of a horizontal range.
func (_gege HorizontalRange) String() string { return _gege.horizontalRangeReference() }

// Eval evaluates a vertical range with prefix returning a list of results or an error.
func (_bfce PrefixVerticalRange) Eval(ctx Context, ev Evaluator) Result {
	_adde := _bfce._cbfc.Reference(ctx, ev)
	switch _adde.Type {
	case ReferenceTypeSheet:
		if _fgba(_adde, ctx) {
			return MakeErrorResultType(ErrorTypeName, _d.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _adde.Value))
		}
		_eeddc := _bfce.verticalRangeReference(_adde.Value)
		if _abcec, _cdea := ev.GetFromCache(_eeddc); _cdea {
			return _abcec
		}
		_egdaad := ctx.Sheet(_adde.Value)
		_caeaf, _afca := _ccfgf(_egdaad, _bfce._dcdb, _bfce._cbcf)
		_fbebc := _cfcef(_egdaad, ev, _caeaf, _afca)
		ev.SetCache(_eeddc, _fbebc)
		return _fbebc
	default:
		return MakeErrorResult(_d.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _adde.Type))
	}
}

type yyLexer interface {
	Lex(_baffb *yySymType) int
	Error(_caeed string)
}

const _beea = 57369

// Substitute is an implementation of the Excel SUBSTITUTE function.
func Substitute(args []Result) Result {
	_ffag := len(args)
	if _ffag != 3 && _ffag != 4 {
		return MakeErrorResult("\u0053\u0055\u0042\u0053\u0054\u0049\u0054U\u0054\u0045\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u006f\u0072\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_baff, _eccc := _cggf(args[0], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u0074\u0065\u0078\u0074")
	if _eccc.Type == ResultTypeError {
		return _eccc
	}
	_aega, _eccc := _cggf(args[1], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u006f\u006c\u0064\u0020\u0074\u0065\u0078\u0074")
	if _eccc.Type == ResultTypeError {
		return _eccc
	}
	_gdcc, _eccc := _cggf(args[2], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u006e\u0065\u0077\u0020\u0074\u0065\u0078\u0074")
	if _eccc.Type == ResultTypeError {
		return _eccc
	}
	_aefec := 0
	if _ffag == 3 {
		return MakeStringResult(_dda.Replace(_baff, _aega, _gdcc, -1))
	} else {
		_geedg, _fbce := _gfde(args[3], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u0069\u006e\u0073t\u0061\u006e\u0063\u0065\u005f\u006e\u0075\u006d")
		if _fbce.Type == ResultTypeError {
			return _fbce
		}
		_aefec = int(_geedg)
		if _aefec < 1 {
			return MakeErrorResult("\u0069\u006es\u0074\u0061\u006e\u0063e\u005f\u006eu\u006d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u007a\u0065\u0072\u006f")
		}
		_fgae := _baff
		_befbb := _aefec
		_fdaba := -1
		_adcbb := len(_aega)
		_bddcb := 0
		for {
			_befbb--
			_dfca := _dda.Index(_fgae, _aega)
			if _dfca == -1 {
				_fdaba = -1
				break
			} else {
				_fdaba = _dfca + _bddcb
				if _befbb == 0 {
					break
				}
				_agde := _adcbb + _dfca
				_bddcb += _agde
				_fgae = _fgae[_agde:]
			}
		}
		if _fdaba == -1 {
			return MakeStringResult(_baff)
		} else {
			_adgd := _baff[:_fdaba]
			_cdbf := _baff[_fdaba+_adcbb:]
			return MakeStringResult(_adgd + _gdcc + _cdbf)
		}
	}
}

// Coupdays implements the Excel COUPDAYS function.
func Coupdays(args []Result) Result {
	_eadb, _cadg := _dccb(args, "\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053")
	if _cadg.Type == ResultTypeError {
		return _cadg
	}
	return MakeNumberResult(_efb(_eadb._eac, _eadb._aea, _eadb._dfff, _eadb._cfeb))
}

var _egbe []byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func _cedgd(_egea []Result, _aeeab string) (*durationArgs, Result) {
	_fdag := len(_egea)
	if _fdag != 5 && _fdag != 6 {
		return nil, MakeErrorResult(_aeeab + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0066\u0069\u0076\u0065\u0020\u006fr\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_agg, _facf, _gefc := _gdge(_egea[0], _egea[1], _aeeab)
	if _gefc.Type == ResultTypeError {
		return nil, _gefc
	}
	_def := _egea[2]
	if _def.Type != ResultTypeNumber {
		return nil, MakeErrorResult(_aeeab + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0063\u006fupo\u006e r\u0061\u0074\u0065\u0020\u006f\u0066\u0020ty\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
	}
	_cfcg := _def.ValueNumber
	if _cfcg < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "\u0043\u006f\u0075po\u006e\u0020\u0072\u0061\u0074\u0065\u0020\u0073\u0068o\u0075l\u0064 \u006eo\u0074\u0020\u0062\u0065\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	_ggab := _egea[3]
	if _ggab.Type != ResultTypeNumber {
		return nil, MakeErrorResult(_aeeab + " \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0079\u0069\u0065\u006cd\u0020\u0072\u0061\u0074\u0065\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
	}
	_cbab := _ggab.ValueNumber
	if _cbab < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006c\u0064\u0020r\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006ft\u0020\u0062\u0065\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	_aagc := _egea[4]
	if _aagc.Type != ResultTypeNumber {
		return nil, MakeErrorResult(_aeeab + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0066\u0072\u0065\u0071u\u0065n\u0063y\u0020o\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_dgacf := float64(int(_aagc.ValueNumber))
	if !_gabe(_dgacf) {
		return nil, MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_ddc := 0
	if _fdag == 6 && _egea[5].Type != ResultTypeEmpty {
		_ddgg := _egea[5]
		if _ddgg.Type != ResultTypeNumber {
			return nil, MakeErrorResult(_aeeab + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020b\u0061\u0073\u0069\u0073\u0020\u006f\u0066 \u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_ddc = int(_ddgg.ValueNumber)
		if !_gaff(_ddc) {
			return nil, MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062a\u0073\u0069\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020f\u006f\u0072\u0020"+_aeeab)
		}
	}
	return &durationArgs{_agg, _facf, _cfcg, _cbab, _dgacf, _ddc}, _efc
}

// NA is an implementation of the Excel NA() function that just returns the #N/A! error.
func NA(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u004eA\u0028\u0029\u0020\u0061c\u0063\u0065\u0070\u0074\u0073 \u006eo\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074s")
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}
func _adgcd(_egabe [][]Result) float64 {
	if len(_egabe) == 2 {
		_dcbdd := _egabe[0][0].AsNumber()
		_ggeed := _egabe[0][1].AsNumber()
		_abdg := _egabe[1][0].AsNumber()
		_egad := _egabe[1][1].AsNumber()
		if _dcbdd.Type != ResultTypeNumber || _ggeed.Type != ResultTypeNumber || _abdg.Type != ResultTypeNumber || _egad.Type != ResultTypeNumber {
			return _cc.NaN()
		}
		return _dcbdd.ValueNumber*_egad.ValueNumber - _abdg.ValueNumber*_ggeed.ValueNumber
	}
	_fedaa := float64(0)
	_eaab := float64(1)
	for _ceeba := range _egabe {
		_fedaa += _eaab * _egabe[0][_ceeba].ValueNumber * _adgcd(_debd(_egabe, _ceeba))
		_eaab *= -1
	}
	return _fedaa
}
func _bac(_ebea int) bool {
	if _ebea == _ebea/400*400 {
		return true
	}
	if _ebea == _ebea/100*100 {
		return false
	}
	return _ebea == _ebea/4*4
}

// Mirr implements the Excel MIRR function.
func Mirr(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u004d\u0049R\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeList && args[0].Type != ResultTypeArray {
		return MakeErrorResult("M\u0049\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020ar\u0072\u0061\u0079 \u0074y\u0070\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049\u0052\u0052\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0069\u006e\u0061\u006e\u0063e\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_dbbd := args[1].ValueNumber + 1
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049\u0052\u0052\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0069\u006e\u0076\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_fbff := args[2].ValueNumber + 1
	if _fbff == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}
	_aaee := _afgc(args[0])
	_eab := float64(len(_aaee))
	_gbddf, _edeg := 0.0, 0.0
	_ffed, _addf := 1.0, 1.0
	_bbgd, _gfbcc := false, false
	for _, _fdage := range _aaee {
		for _, _acc := range _fdage {
			if _acc.Type == ResultTypeNumber && !_acc.IsBoolean {
				_fcae := _acc.ValueNumber
				if _fcae == 0 {
					continue
				} else {
					if _fcae > 0 {
						_bbgd = true
						_edeg += _acc.ValueNumber * _addf
					} else {
						_gfbcc = true
						_gbddf += _acc.ValueNumber * _ffed
					}
					_ffed /= _dbbd
					_addf /= _fbff
				}
			}
		}
	}
	if !_bbgd || !_gfbcc {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}
	_dfcb := -_edeg / _gbddf
	_dfcb *= _cc.Pow(_fbff, _eab-1)
	_dfcb = _cc.Pow(_dfcb, 1/(_eab-1))
	return MakeNumberResult(_dfcb - 1)
}
func _feff(_dfd, _bgcd, _adcc, _cgcb, _edfb float64, _ddfd int) float64 {
	_ecee := _acad(_dfd, _adcc, _cgcb, _edfb, _ddfd)
	var _febd float64
	if _bgcd == 1 {
		if _ddfd == 1 {
			_febd = 0
		} else {
			_febd = -_cgcb
		}
	} else {
		if _ddfd == 1 {
			_febd = _gfec(_dfd, _bgcd-2, _ecee, _cgcb, 1) - _ecee
		} else {
			_febd = _gfec(_dfd, _bgcd-1, _ecee, _cgcb, 0)
		}
	}
	return _febd * _dfd
}
func _gedc(_daa string) bool {
	for _, _daf := range _dfg {
		_gbdd := _daf.FindStringSubmatch(_daa)
		if len(_gbdd) > 1 {
			return true
		}
	}
	return false
}
func _gaf(_bcb string) (int, int, int, bool, Result) {
	_dbb := ""
	_fda := []string{}
	for _dbbg, _fga := range _ceb {
		_fda = _fga.FindStringSubmatch(_bcb)
		if len(_fda) > 1 {
			_dbb = _dbbg
			break
		}
	}
	if _dbb == "" {
		return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
	}
	_fcf := false
	var _bba, _fbgf, _bdbf int
	var _abce error
	switch _dbb {
	case "\u006d\u006d\u002f\u0064\u0064\u002f\u0079\u0079":
		_fbgf, _abce = _c.Atoi(_fda[1])
		if _abce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_bdbf, _abce = _c.Atoi(_fda[3])
		if _abce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_bba, _abce = _c.Atoi(_fda[5])
		if _abce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		if _bba < 0 || _bba > 9999 || (_bba > 99 && _bba < 1900) {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_bba = _fgg(_bba)
		_fcf = _fda[8] == ""
	case "\u006dm\u0020\u0064\u0064\u002c\u0020\u0079y":
		_fbgf = _aaf[_fda[1]]
		_bdbf, _abce = _c.Atoi(_fda[14])
		if _abce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_bba, _abce = _c.Atoi(_fda[16])
		if _abce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		if _bba < 0 || _bba > 9999 || (_bba > 99 && _bba < 1900) {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_bba = _fgg(_bba)
		_fcf = _fda[19] == ""
	case "\u0079\u0079\u002d\u006d\u006d\u002d\u0064\u0064":
		_cbeb, _dce := _c.Atoi(_fda[1])
		if _dce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_fefa, _dce := _c.Atoi(_fda[3])
		if _dce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_fcaa, _dce := _c.Atoi(_fda[5])
		if _dce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		if _cbeb >= 1900 && _cbeb < 10000 {
			_bba = _cbeb
			_fbgf = _fefa
			_bdbf = _fcaa
		} else if _cbeb > 0 && _cbeb < 13 {
			_fbgf = _cbeb
			_bdbf = _fefa
			_bba = _fcaa
		} else {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_fcf = _fda[8] == ""
	case "y\u0079\u002d\u006d\u006d\u0053\u0074\u0072\u002d\u0064\u0064":
		_bba, _abce = _c.Atoi(_fda[16])
		if _abce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_fbgf = _aaf[_fda[3]]
		_bdbf, _abce = _c.Atoi(_fda[1])
		if _abce != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
		}
		_fcf = _fda[19] == ""
	}
	if !_badd(_bba, _fbgf, _bdbf) {
		return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _geb)
	}
	return _bba, _fbgf, _bdbf, _fcf, _efc
}

var _egbc = [...]int{-1000, -7, -3, -1, 27, 18, 22, 23, -2, -8, -4, -9, 20, -14, 10, 11, 12, 13, -5, -13, -6, -12, 17, 16, 15, 9, 4, 5, 22, 23, 24, 25, 26, 28, 29, 30, 31, 27, 32, 35, -1, 18, 27, -15, -17, -1, -1, -1, -1, 33, -5, 4, 5, 21, -16, -11, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 19, 36, 34, 21, -5, 33, 21, 34, 19, -17, -1, -5, -10, -1}

const _bdbed = 57357

type cmpResult int8

const _cdacg = 57376

// LastRow returns 0 for the invalid reference context.
func (_aadab *ivr) LastRow(colFrom string) int { return 0 }

// MakeBoolResult constructs a boolean result (internally a number).
func MakeBoolResult(b bool) Result {
	if b {
		return Result{Type: ResultTypeNumber, ValueNumber: 1, IsBoolean: true}
	}
	return Result{Type: ResultTypeNumber, ValueNumber: 0, IsBoolean: true}
}

// CellRef is a reference to a single cell
type CellRef struct{ _ebd string }

func (_bbacf *plex) Lex(lval *yySymType) int {
	_cdaf = true
	_cadce := <-_bbacf._bede
	if _cadce != nil {
		lval._bdeac = _cadce
		return int(lval._bdeac._eddb)
	}
	return 0
}
func _addab() yyParser                      { return &yyParserImpl{} }
func (_defdf *yyParserImpl) Lookahead() int { return _defdf._accfg }

// Eval evaluates and returns the result of a Negate expression.
func (_ceca Negate) Eval(ctx Context, ev Evaluator) Result {
	_ccdg := _ceca._eefg.Eval(ctx, ev)
	if _ccdg.Type == ResultTypeNumber {
		return MakeNumberResult(-_ccdg.ValueNumber)
	}
	return MakeErrorResult("\u004e\u0045\u0047A\u0054\u0045\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
}

// Reference returns a string reference value to a named range.
func (_aadga NamedRangeRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeNamedRange, Value: _aadga._bbfde}
}

const _ggd = "\u0042\u0069\u006e\u004f\u0070\u0054y\u0070\u0065\u0055\u006e\u006bn\u006fw\u006e\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u0050\u006c\u0075\u0073\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u004d\u0069\u006e\u0075\u0073\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065M\u0075lt\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u0044\u0069\u0076\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u0045\u0078\u0070\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u004c\u0054\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065G\u0054B\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u0045\u0051\u0042\u0069nO\u0070\u0054\u0079\u0070\u0065\u004c\u0045\u0051\u0042i\u006eO\u0070\u0054\u0079\u0070\u0065\u0047\u0045\u0051\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065N\u0045\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u0043\u006f\u006e\u0063\u0061\u0074"

// BinaryExpr is a binary expression.
type BinaryExpr struct {
	_eb, _bd Expression
	_cb      BinOpType
}
type tokenType int

// Ispmt implements the Excel ISPMT function.
func Ispmt(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("\u0049\u0053P\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bdgfd := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0070e\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gfgd := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069o\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gfegf := args[2].ValueNumber
	if _gfegf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049S\u0050\u004d\u0054\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0065\u0072i\u006f\u0064\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062er\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065s\u0065\u006e\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_efda := args[3].ValueNumber
	return MakeNumberResult(_efda * _bdgfd * (_gfgd/_gfegf - 1))
}

// PrefixHorizontalRange is a range expression that when evaluated returns a list of Results from references like Sheet1!1:4 (all cells from rows 1 to 4 of sheet 'Sheet1').
type PrefixHorizontalRange struct {
	_fdcfe       Expression
	_ffba, _cebd int
}

// NewFunction constructs a new function call expression.
func NewFunction(name string, args []Expression) Expression {
	return FunctionCall{_ffcfde: name, _gace: args}
}

// Parse parses a string to get an Expression.
func ParseString(s string) Expression {
	if s == "" {
		return NewEmptyExpr()
	}
	return Parse(_dda.NewReader(s))
}

// FloorPrecise is an implementation of the FlOOR.PRECISE function.
func FloorPrecise(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0046\u004cO\u004f\u0052\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0046L\u004f\u004fR\u002e\u0050\u0052\u0045C\u0049\u0053\u0045(\u0029\u0020\u0061\u006c\u006c\u006f\u0077\u0073\u0020at\u0020\u006d\u006fs\u0074\u0020t\u0077\u006f\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_adcbg := args[0].AsNumber()
	if _adcbg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020F\u004c\u004f\u004f\u0052\u002e\u0050\u0052E\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_edfc := float64(1)
	if _adcbg.ValueNumber < 0 {
		_edfc = -1
	}
	if len(args) > 1 {
		_dbgc := args[1].AsNumber()
		if _dbgc.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006ed\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020F\u004c\u004f\u004f\u0052\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065r")
		}
		_edfc = _cc.Abs(_dbgc.ValueNumber)
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Floor(_adcbg.ValueNumber))
	}
	_gdbbc := _adcbg.ValueNumber
	_gdbbc, _afcf := _cc.Modf(_gdbbc / _edfc)
	if _afcf != 0 {
		if _adcbg.ValueNumber < 0 {
			_gdbbc--
		}
	}
	return MakeNumberResult(_gdbbc * _edfc)
}

type criteriaRegex struct {
	_eefb byte
	_cfcd string
}

const _dbfc = 57373

// Mid is an implementation of the Excel MID function that returns a copy
// of the string with each word capitalized.
func Mid(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u004d\u0049\u0044\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065e\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	_eadg := args[0]
	if _eadg.Type == ResultTypeError {
		return _eadg
	}
	if _eadg.Type != ResultTypeString && _eadg.Type != ResultTypeNumber && _eadg.Type != ResultTypeEmpty {
		return MakeErrorResult("\u004d\u0049\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0065x\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	_ddbgc := args[0].Value()
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049D\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0073\u0074\u0061\u0072\u0074\u005fn\u0075\u006d\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_effb := int(args[1].ValueNumber)
	if _effb < 1 {
		return MakeErrorResult("M\u0049\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0074\u0061\u0072\u0074\u005fn\u0075\u006d\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006dor\u0065\u0020\u0074h\u0061n\u0020\u0030")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049D\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u005f\u0063\u0068a\u0072\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_aadeb := int(args[2].ValueNumber)
	if _aadeb < 0 {
		return MakeErrorResult("\u004d\u0049\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u005f\u0063\u0068a\u0072\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	_gafa := len(_ddbgc)
	if _effb > _gafa {
		return MakeStringResult("")
	}
	_effb--
	_aebfe := _effb + _aadeb
	if _aebfe > _gafa {
		return MakeStringResult(_ddbgc[_effb:])
	} else {
		return MakeStringResult(_ddbgc[_effb:_aebfe])
	}
}

// Combin is an implementation of the Excel COMBINA function whic returns the
// number of combinations.
func Combin(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0043\u004f\u004d\u0042\u0049\u004e\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gfgc := args[0].AsNumber()
	_feab := args[1].AsNumber()
	if _gfgc.Type != ResultTypeNumber || _feab.Type != ResultTypeNumber {
		return MakeErrorResult("C\u004f\u004d\u0042\u0049\u004e\u0028)\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_aged := _cc.Trunc(_gfgc.ValueNumber)
	_fgecc := _cc.Trunc(_feab.ValueNumber)
	if _fgecc > _aged {
		return MakeErrorResult("\u0043O\u004d\u0042\u0049\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006b\u0020\u003c\u003d\u0020\u006e")
	}
	if _fgecc == _aged || _fgecc == 0 {
		return MakeNumberResult(1)
	}
	_addc := float64(1)
	for _dbff := float64(1); _dbff <= _fgecc; _dbff++ {
		_addc *= (_aged + 1 - _dbff) / _dbff
	}
	return MakeNumberResult(_addc)
}

var _dcee *_fa.Rand

const _cdbb = 57377

var _ceaeg = map[string]bool{"\u0049F\u0045\u0052\u0052\u004f\u0052": true, "\u0049\u0046\u004e\u0041": true, "\u005f\u0078\u006c\u0066\u006e\u002e\u0049\u0046\u004e\u0041": true, "\u0049\u0053\u0045R\u0052": true, "\u0049S\u0045\u0052\u0052\u004f\u0052": true, "\u0049\u0053\u004e\u0041": true, "\u0049\u0053\u0052E\u0046": true}

const _geb = "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0044\u0041\u0054\u0045\u0056\u0041\u004c\u0055\u0045"

var _fdcdab = [...]string{"\u0024\u0065\u006e\u0064", "\u0065\u0072\u0072o\u0072", "\u0024\u0075\u006e\u006b", "t\u006fk\u0065\u006e\u0048\u006f\u0072\u0069\u007a\u006fn\u0074\u0061\u006c\u0052an\u0067\u0065", "\u0074o\u006be\u006e\u0056\u0065\u0072\u0074i\u0063\u0061l\u0052\u0061\u006e\u0067\u0065", "\u0074\u006f\u006b\u0065\u006e\u0052\u0065\u0073\u0065\u0072\u0076\u0065d\u004e\u0061\u006d\u0065", "\u0074\u006f\u006be\u006e\u0044\u0044\u0045\u0043\u0061\u006c\u006c", "\u0074\u006f\u006b\u0065\u006e\u004c\u0065\u0078\u0045\u0072\u0072\u006f\u0072", "\u0074o\u006be\u006e\u004e\u0061\u006d\u0065\u0064\u0052\u0061\u006e\u0067\u0065", "\u0074o\u006b\u0065\u006e\u0042\u006f\u006fl", "t\u006f\u006b\u0065\u006e\u004e\u0075\u006d\u0062\u0065\u0072", "t\u006f\u006b\u0065\u006e\u0053\u0074\u0072\u0069\u006e\u0067", "\u0074\u006f\u006b\u0065\u006e\u0045\u0072\u0072\u006f\u0072", "\u0074\u006f\u006b\u0065\u006e\u0045\u0072\u0072\u006f\u0072\u0052\u0065\u0066", "\u0074\u006f\u006b\u0065\u006e\u0053\u0068\u0065\u0065\u0074", "\u0074o\u006b\u0065\u006e\u0043\u0065\u006cl", "t\u006fk\u0065\u006e\u0046\u0075\u006e\u0063\u0074\u0069o\u006e\u0042\u0075\u0069lt\u0069\u006e", "t\u006f\u006b\u0065\u006e\u004c\u0042\u0072\u0061\u0063\u0065", "t\u006f\u006b\u0065\u006e\u0052\u0042\u0072\u0061\u0063\u0065", "t\u006f\u006b\u0065\u006e\u004c\u0050\u0061\u0072\u0065\u006e", "t\u006f\u006b\u0065\u006e\u0052\u0050\u0061\u0072\u0065\u006e", "\u0074o\u006b\u0065\u006e\u0050\u006c\u0075s", "\u0074\u006f\u006b\u0065\u006e\u004d\u0069\u006e\u0075\u0073", "\u0074o\u006b\u0065\u006e\u004d\u0075\u006ct", "\u0074\u006f\u006b\u0065\u006e\u0044\u0069\u0076", "\u0074\u006f\u006b\u0065\u006e\u0045\u0078\u0070", "\u0074o\u006b\u0065\u006e\u0045\u0051", "\u0074o\u006b\u0065\u006e\u004c\u0054", "\u0074o\u006b\u0065\u006e\u0047\u0054", "\u0074\u006f\u006b\u0065\u006e\u004c\u0045\u0051", "\u0074\u006f\u006b\u0065\u006e\u0047\u0045\u0051", "\u0074o\u006b\u0065\u006e\u004e\u0045", "\u0074\u006f\u006b\u0065\u006e\u0043\u006f\u006c\u006f\u006e", "\u0074\u006f\u006b\u0065\u006e\u0043\u006f\u006d\u006d\u0061", "\u0074\u006f\u006b\u0065\u006e\u0041\u006d\u0070\u0065r\u0073\u0061\u006e\u0064", "\u0074o\u006b\u0065\u006e\u0053\u0065\u006di"}

// Ipmt implements the Excel IPMT function.
func Ipmt(args []Result) Result {
	_cdcf := len(args)
	if _cdcf < 4 || _cdcf > 6 {
		return MakeErrorResult("\u0049P\u004d\u0054\u0020\u0072\u0065\u0071\u0075ir\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074s\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065n\u0020\u0066ou\u0072\u0020\u0061n\u0064\u0020\u0073\u0069\u0078")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("I\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_egd := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0050\u004dT\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_afad := args[1].ValueNumber
	if _afad <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u0050\u004d\u0054\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006fd\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_edac := args[2].ValueNumber
	if _edac <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062er\u0020o\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_afdd := args[3].ValueNumber
	_cbcb := 0.0
	if _cdcf > 4 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0049\u0050\u004d\u0054\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_cbcb = args[4].ValueNumber
	}
	_dcb := 0
	if _cdcf == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("I\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_dcb = int(args[5].ValueNumber)
		if _dcb != 0 {
			_dcb = 1
		}
	}
	return MakeNumberResult(_feff(_egd, _afad, _edac, _afdd, _cbcb, _dcb))
}
func (_febbc *ivr) NamedRange(ref string) Reference { return ReferenceInvalid }

// Update returns the same object as updating sheet references does not affect SheetPrefixExpr.
func (_fdbc SheetPrefixExpr) Update(q *_a.UpdateQuery) Expression { return _fdbc }

const _gge = "\u0028\u0028\u006a\u0061\u006e|\u006a\u0061\u006e\u0075\u0061\u0072\u0079\u0029\u007c\u0028\u0066\u0065\u0062\u007c\u0066\u0065\u0062\u0072\u0075a\u0072\u0079\u0029\u007c\u0028\u006da\u0072\u007c\u006da\u0072\u0063\u0068\u0029\u007c\u0028\u0061\u0070\u0072\u007c\u0061\u0070\u0072\u0069\u006c\u0029\u007c\u0028\u006d\u0061\u0079\u0029\u007c\u0028j\u0075\u006e\u007cj\u0075\u006e\u0065\u0029\u007c\u0028\u006a\u0075\u006c\u007c\u006a\u0075\u006c\u0079\u0029\u007c\u0028a\u0075\u0067\u007c\u0061\u0075\u0067\u0075\u0073t\u0029\u007c\u0028\u0073\u0065\u0070\u007c\u0073\u0065\u0070\u0074\u0065\u006d\u0062\u0065\u0072\u0029\u007c\u0028o\u0063\u0074\u007c\u006f\u0063\u0074\u006f\u0062\u0065\u0072\u0029\u007c\u0028\u006e\u006f\u0076\u007c\u006e\u006f\u0076\u0065\u006d\u0062e\u0072\u0029\u007c\u0028\u0064\u0065\u0063\u007c\u0064\u0065\u0063\u0065\u006d\u0062\u0065\u0072\u0029\u0029"

// Pmt implements the Excel PMT function.
func Pmt(args []Result) Result {
	_gbfa := len(args)
	if _gbfa < 3 || _gbfa > 5 {
		return MakeErrorResult("\u0050\u004dT\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u0020\u0035")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020r\u0061\u0074\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cbdd := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u004dT\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dfgc := args[1].ValueNumber
	if _dfgc == 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u00200")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073 p\u0072\u0065\u0073\u0065\u006e\u0074 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_edgg := args[2].ValueNumber
	_adff := 0.0
	if _gbfa >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("P\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0075\u0074\u0075\u0072e \u0076\u0061\u006c\u0075e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_adff = args[3].ValueNumber
	}
	_cebaf := 0.0
	if _gbfa == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0079\u0070\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_cebaf = args[4].ValueNumber
		if _cebaf != 0 {
			_cebaf = 1
		}
	}
	var _eagd float64
	if _cbdd == 0 {
		_eagd = (_edgg + _adff) / _dfgc
	} else {
		_cbgb := _cc.Pow(1+_cbdd, _dfgc)
		if _cebaf == 1 {
			_eagd = (_adff*_cbdd/(_cbgb-1) + _edgg*_cbdd/(1-1/_cbgb)) / (1 + _cbdd)
		} else {
			_eagd = _adff*_cbdd/(_cbgb-1) + _edgg*_cbdd/(1-1/_cbgb)
		}
	}
	return MakeNumberResult(-_eagd)
}

// Eval evaluates and returns the result of an error expression.
func (_adf Error) Eval(ctx Context, ev Evaluator) Result { return MakeErrorResult(_adf._gbd) }
func _agef(_dfcf, _cgee float64, _bfee int) (float64, Result) {
	_bgc, _ecg := _dcge(_dfcf), _dcge(_cgee)
	_fff := _bgc.Unix()
	_dgfe := _ecg.Unix()
	if _fff == _dgfe {
		return 0, _efc
	}
	_gedd, _fbf, _bfab := _bgc.Date()
	_gfac, _ced, _cedg := _ecg.Date()
	_fee, _bbdd := int(_fbf), int(_ced)
	var _cgd, _bae float64
	switch _bfee {
	case 0:
		if _bfab == 31 {
			_bfab--
		}
		if _bfab == 30 && _cedg == 31 {
			_cedg--
		} else if _ecb := _bac(_gedd); _fee == 2 && ((_ecb && _bfab == 29) || (!_ecb && _bfab == 28)) {
			_bfab = 30
			if _abbd := _bac(_gfac); _bbdd == 2 && ((_abbd && _cedg == 29) || (!_abbd && _cedg == 28)) {
				_cedg = 30
			}
		}
		_cgd = float64((_gfac-_gedd)*360 + (_bbdd-_fee)*30 + (_cedg - _bfab))
		_bae = 360
	case 1:
		_cgd = _cgee - _dfcf
		_cfc := _gedd != _gfac
		if _cfc && (_gfac != _gedd+1 || _fee < _bbdd || (_fee == _bbdd && _bfab < _cedg)) {
			_bbe := 0
			for _agec := _gedd; _agec <= _gfac; _agec++ {
				_bbe += _acfa(_agec, 1)
			}
			_bae = float64(_bbe) / float64(_gfac-_gedd+1)
		} else {
			if !_cfc && _bac(_gedd) {
				_bae = 366
			} else {
				if _cfc && ((_bac(_gedd) && (_fee < 2 || (_fee == 2 && _bfab <= 29))) || (_bac(_gfac) && (_bbdd > 2 || (_bbdd == 2 && _cedg == 29)))) {
					_bae = 366
				} else {
					_bae = 365
				}
			}
		}
	case 2:
		_cgd = _cgee - _dfcf
		_bae = 360
	case 3:
		_cgd = _cgee - _dfcf
		_bae = 365
	case 4:
		if _bfab == 31 {
			_bfab--
		}
		if _cedg == 31 {
			_cedg--
		}
		_cgd = float64((_gfac-_gedd)*360 + (_bbdd-_fee)*30 + (_cedg - _bfab))
		_bae = 360
	default:
		return 0, MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073 \u0066o\u0072\u0020\u0059\u0065\u0061\u0072\u0046r\u0061\u0063")
	}
	return _cgd / _bae, _efc
}

// ErrorType is a formula evaluation error type.
type ErrorType byte

// Eval evaluates and returns a number.
func (_bagd Number) Eval(ctx Context, ev Evaluator) Result { return MakeNumberResult(_bagd._ceacb) }

var _edde = [...]int{-1, 1, 1, -1, -2, 0}

func _fcbdd() {
	_dbee = _ge.MustCompile("\u005e\u0028\u005b\u0030\u002d\u0039\u005d\u002b\u0029\u0024")
	_eebef = _ge.MustCompile("\u005e=\u0028\u002e\u002a\u0029\u0024")
	_fegdb = _ge.MustCompile("\u005e<\u0028\u002e\u002a\u0029\u0024")
	_ebeaa = _ge.MustCompile("\u005e>\u0028\u002e\u002a\u0029\u0024")
	_ddbcf = _ge.MustCompile("\u005e\u003c\u003d\u0028\u002e\u002a\u0029\u0024")
	_adag = _ge.MustCompile("\u005e\u003e\u003d\u0028\u002e\u002a\u0029\u0024")
}

// Error is an error expression.
type Error struct{ _gbd string }

// Eval evaluates and returns the result of an empty expression.
func (_ccac EmptyExpr) Eval(ctx Context, ev Evaluator) Result { return MakeEmptyResult() }

const _ggafe = 57372

// Update updates references in the Negate after removing a row/column.
func (_eada Negate) Update(q *_a.UpdateQuery) Expression { return Negate{_eefg: _eada._eefg.Update(q)} }

// GetWidth returns 0 for the invalid reference context.
func (_gbeb *ivr) GetWidth(colIdx int) float64 { return float64(0) }
func _dagaa(_feed yyLexer, _gcec *yySymType) (_ebafc, _bdbdc int) {
	_bdbdc = 0
	_ebafc = _feed.Lex(_gcec)
	if _ebafc <= 0 {
		_bdbdc = _aeed[0]
		goto _ebdcf
	}
	if _ebafc < len(_aeed) {
		_bdbdc = _aeed[_ebafc]
		goto _ebdcf
	}
	if _ebafc >= _affd {
		if _ebafc < _affd+len(_abgf) {
			_bdbdc = _abgf[_ebafc-_affd]
			goto _ebdcf
		}
	}
	for _dcdad := 0; _dcdad < len(_eaga); _dcdad += 2 {
		_bdbdc = _eaga[_dcdad+0]
		if _bdbdc == _ebafc {
			_bdbdc = _eaga[_dcdad+1]
			goto _ebdcf
		}
	}
_ebdcf:
	if _bdbdc == 0 {
		_bdbdc = _abgf[1]
	}
	if _fbee >= 3 {
		_d.Printf("l\u0065\u0078\u0020\u0025\u0073\u0028\u0025\u0064\u0029\u000a", _cecbd(_bdbdc), uint(_ebafc))
	}
	return _ebafc, _bdbdc
}
func Sign(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0053\u0049\u0047\u004e(\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_bcfbg := args[0].AsNumber()
	if _bcfbg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0049\u0047N(\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020a\u0020n\u0075m\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _bcfbg.ValueNumber < 0 {
		return MakeNumberResult(-1)
	} else if _bcfbg.ValueNumber > 0 {
		return MakeNumberResult(1)
	}
	return MakeNumberResult(0)
}

// Text is an implementation of the Excel TEXT function.
func Text(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("T\u0045\u0058\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	_bdcd := args[0]
	if _bdcd.Type != ResultTypeNumber && _bdcd.Type != ResultTypeString && _bdcd.Type != ResultTypeEmpty {
		return MakeErrorResult("\u0054\u0045\u0058\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	if args[1].Type != ResultTypeString {
		return MakeErrorResult("\u0054E\u0058\u0054 \u0072\u0065\u0071\u0075i\u0072\u0065\u0073 \u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072gu\u006d\u0065\u006et\u0020\u0074o\u0020\u0062\u0065\u0020\u0061\u0020s\u0074\u0072i\u006e\u0067")
	}
	_fdce := args[1].ValueString
	switch _bdcd.Type {
	case ResultTypeNumber:
		return MakeStringResult(_bg.Number(_bdcd.ValueNumber, _fdce))
	case ResultTypeString:
		return MakeStringResult(_bg.String(_bdcd.ValueString, _fdce))
	case ResultTypeEmpty:
		return MakeStringResult(_bg.Number(0, _fdce))
	case ResultTypeArray, ResultTypeList:
		return MakeErrorResultType(ErrorTypeSpill, "\u0054\u0045X\u0054\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0077\u006f\u0072\u006b\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0072ra\u0079\u0073")
	default:
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0066\u006fr\u0020T\u0045\u0058\u0054")
	}
}

// ISEVEN is an implementation of the Excel ISEVEN() function.
func IsEven(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0045VE\u004e\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070t\u0073 \u0061 \u0073i\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0045\u0056\u0045\u004e \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_babb := int(args[0].ValueNumber)
	return MakeBoolResult(_babb == _babb/2*2)
}

// Function is a standard function whose result only depends on its arguments.
type Function func(_bgff []Result) Result

var _aed = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var _cfba = map[string]*_ge.Regexp{}

const _fdca = 57370

// Update makes a reference to point to one of the neighboring cells after removing a row/column with respect to the update type.
func (_bf CellRef) Update(q *_a.UpdateQuery) Expression {
	if q.UpdateCurrentSheet {
		_bf._ebd = _cgg(_bf._ebd, q)
	}
	return _bf
}
func MakeRangeReference(ref string) Reference { return Reference{Type: ReferenceTypeRange, Value: ref} }
func _ebca(_bfdb []Result, _bcge string) (*cumulArgs, Result) {
	if len(_bfdb) != 6 {
		return nil, MakeErrorResult(_bcge + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _bfdb[0].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bcge + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0072\u0061te \u0074o \u0062\u0065\u0020\u006e\u0075\u006d\u0062er\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_caf := _bfdb[0].ValueNumber
	if _caf <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+"\u0020r\u0065\u0071u\u0069\u0072\u0065s\u0020\u0072\u0061\u0074\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006fs\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _bfdb[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bcge + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_acgg := _bfdb[1].ValueNumber
	if _acgg <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u0061r\u0067\u0075\u006de\u006e\u0074")
	}
	if _bfdb[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bcge + "\u0020r\u0065\u0071u\u0069\u0072\u0065s\u0020\u0070\u0072\u0065\u0073\u0065\u006et\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_feac := _bfdb[2].ValueNumber
	if _feac <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _bfdb[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bcge + "\u0020r\u0065\u0071u\u0069\u0072\u0065\u0073 \u0073\u0074\u0061r\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_cagg := _bfdb[3].ValueNumber
	if _cagg <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073t\u0061\u0072\u0074\u0020\u0070\u0065\u0072\u0069o\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _bfdb[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bcge + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0065\u006e\u0064\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_aabe := _bfdb[4].ValueNumber
	if _aabe <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+"\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0065\u006e\u0064\u0020\u0070\u0065\u0072\u0069\u006fd\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	if _aabe < _cagg {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+"\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0065\u006e\u0064\u0020p\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006c\u0061\u0074\u0065\u0072\u0020o\u0072\u0020\u0065\u0071\u0075a\u006c\u0020\u0074\u006f\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064")
	}
	if _aabe > _acgg {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+" \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074o\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0070\u0065\u0072io\u0064\u0073\u0020r\u0061n\u0067\u0065")
	}
	_bda := int(_bfdb[5].ValueNumber)
	if _bda != 0 && _bda != 1 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bcge+" \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u0074\u006f \u0062\u0065\u00200\u0020o\u0072\u0020\u0031")
	}
	return &cumulArgs{_caf, _acgg, _feac, _cagg, _aabe, _bda}, _efc
}

// String returns a string representation of ConstArrayExpr.
func (_dff ConstArrayExpr) String() string { return "" }
func _fefcg(_edfe []Result) Result {
	_gfcgg := _edfe[0].ValueArray
	if len(_edfe) == 1 {
		_ecacf := [][]Result{}
		for _, _bgab := range _gfcgg {
			_ecacf = append(_ecacf, _geeg([]Result{MakeListResult(_bgab)}).ValueList)
		}
		return MakeArrayResult(_ecacf)
	} else if len(_edfe) == 2 {
		_gacfed := len(_gfcgg)
		_bcaa := len(_gfcgg[0])
		_dceae := _ebacd(_edfe[1], _gacfed, _bcaa)
		_ddgdc := len(_dceae)
		_bbbag := [][]Result{}
		var _dddc []Result
		for _dggb, _fedb := range _gfcgg {
			if _dggb < _ddgdc {
				_dddc = _dceae[_dggb]
			} else {
				_dddc = _fgggg(MakeErrorResultType(ErrorTypeNA, ""), _bcaa)
			}
			_bbbag = append(_bbbag, _geeg([]Result{MakeListResult(_fedb), MakeListResult(_dddc)}).ValueList)
		}
		return MakeArrayResult(_bbbag)
	} else if len(_edfe) == 3 {
		_eadf := len(_gfcgg)
		_gaab := len(_gfcgg[0])
		_abbb := _ebacd(_edfe[1], _eadf, _gaab)
		_bgage := _ebacd(_edfe[2], _eadf, _gaab)
		_fadde := len(_abbb)
		_dgda := len(_bgage)
		_edae := [][]Result{}
		var _accae, _fafe []Result
		for _cdaae, _dbbda := range _gfcgg {
			if _cdaae < _fadde {
				_accae = _abbb[_cdaae]
			} else {
				_accae = _fgggg(MakeErrorResultType(ErrorTypeNA, ""), _gaab)
			}
			if _cdaae < _dgda {
				_fafe = _bgage[_cdaae]
			} else {
				_fafe = _fgggg(MakeErrorResultType(ErrorTypeNA, ""), _gaab)
			}
			_edae = append(_edae, _geeg([]Result{MakeListResult(_dbbda), MakeListResult(_accae), MakeListResult(_fafe)}).ValueList)
		}
		return MakeArrayResult(_edae)
	}
	return MakeErrorResultType(ErrorTypeValue, "")
}

// Int is an implementation of the Excel INT() function that rounds a number
// down to an integer.
func Int(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("I\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_aggb := args[0].AsNumber()
	if _aggb.Type != ResultTypeNumber {
		return MakeErrorResult("I\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_eeegf, _bbdad := _cc.Modf(_aggb.ValueNumber)
	if _bbdad < 0 {
		_eeegf--
	}
	return MakeNumberResult(_eeegf)
}
func (_abbdc *noCache) SetCache(key string, value Result) {}
func _dcge(_gf float64) _ee.Time {
	_gbde := int64((_gf - _gab) * _edg)
	return _ee.Unix(0, _gbde).UTC()
}

// GCD implements the Excel GCD() function which returns the greatest common
// divisor of a range of numbers.
func GCD(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0047\u0043D(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0061t \u006c\u0065\u0061\u0073\u0074\u0020\u006fne\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_fgbg := []float64{}
	for _, _ddgdcb := range args {
		switch _ddgdcb.Type {
		case ResultTypeString:
			_edffd := _ddgdcb.AsNumber()
			if _edffd.Type != ResultTypeNumber {
				return MakeErrorResult("\u0047\u0043D(\u0029\u0020\u006fn\u006c\u0079\u0020\u0061cce\u0070ts\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
			}
			_fgbg = append(_fgbg, _edffd.ValueNumber)
		case ResultTypeList, ResultTypeArray:
			_fggdb := GCD(_ddgdcb.ListValues())
			if _fggdb.Type != ResultTypeNumber {
				return _fggdb
			}
			_fgbg = append(_fgbg, _fggdb.ValueNumber)
		case ResultTypeNumber:
			_fgbg = append(_fgbg, _ddgdcb.ValueNumber)
		case ResultTypeError:
			return _ddgdcb
		default:
			return MakeErrorResult(_d.Sprintf("\u0047\u0043\u0044()\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065d\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _ddgdcb.Type))
		}
	}
	if _fgbg[0] < 0 {
		return MakeErrorResult("\u0047\u0043D\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(_fgbg) == 1 {
		return MakeNumberResult(_fgbg[0])
	}
	_efab := _fgbg[0]
	for _geeb := 1; _geeb < len(_fgbg); _geeb++ {
		if _fgbg[_geeb] < 0 {
			return MakeErrorResult("\u0047\u0043D\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
		}
		_efab = _afde(_efab, _fgbg[_geeb])
	}
	return MakeNumberResult(_efab)
}

// Eval evaluates and returns a string.
func (_cfffg String) Eval(ctx Context, ev Evaluator) Result { return MakeStringResult(_cfffg._adfgc) }

type rangeIndex struct {
	_cdag  int
	_gcbgb int
}

// Eval evaluates and returns the result of a sheet expression.
func (_dffe SheetPrefixExpr) Eval(ctx Context, ev Evaluator) Result {
	return MakeErrorResult("\u0073\u0068\u0065\u0065\u0074\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065r\u0020\u0062\u0065\u0020\u0065v\u0061\u006cu\u0061\u0074\u0065\u0064")
}

// Right implements the Excel RIGHT(string,[n]) function which returns the
// rightmost n characters.
func Right(args []Result) Result {
	_addag := 1
	switch len(args) {
	case 1:
	case 2:
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u0049\u0047\u0048\u0054\u0020\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_addag = int(args[1].ValueNumber)
		if _addag < 0 {
			return MakeErrorResult("R\u0049\u0047\u0048\u0054\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u003e\u003d \u0030")
		}
		if _addag == 0 {
			return MakeStringResult("")
		}
	default:
		return MakeErrorResult("\u0052\u0049\u0047HT\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020o\u006ee\u0020o\u0072 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type == ResultTypeList {
		return MakeErrorResult("\u0052\u0049\u0047\u0048\u0054\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020c\u0061l\u006c\u0065\u0064\u0020\u006f\u006e\u0020\u0061\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_agab := args[0].Value()
	_bbaf := len(_agab)
	if _addag > _bbaf {
		return MakeStringResult(_agab)
	}
	return MakeStringResult(_agab[_bbaf-_addag : _bbaf])
}

var _eaga = [...]int{0}

// Dollarfr implements the Excel DOLLARFR function.
func Dollarfr(args []Result) Result {
	_cbfa, _fccdf, _gdfe := _fdcd(args, "\u0044\u004f\u004c\u004c\u0041\u0052\u0046\u0052")
	if _gdfe.Type == ResultTypeError {
		return _gdfe
	}
	if _fccdf == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0044\u004f\u004c\u004c\u0041R\u0046\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066r\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if _cbfa == 0 {
		return MakeNumberResult(0)
	}
	_fbfg := _cbfa < 0
	if _fbfg {
		_cbfa = -_cbfa
	}
	_dgbf := float64(int(_cbfa))
	_faa := args[0].Value()
	_daac := _dda.Split(_faa, "\u002e")
	_dae := 0.0
	if len(_daac) > 1 {
		var _cega error
		_ecag := _daac[1]
		_dae, _cega = _c.ParseFloat(_ecag, 64)
		if _cega != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006fr \u0044\u004f\u004cL\u0041R\u0046\u0052")
		}
		_eaef := float64(len(_ecag))
		_dae /= _cc.Pow(10, _eaef)
	}
	_cead := _dae*_fccdf/_cc.Pow(10, float64(int(_cc.Log10(_fccdf)))+1) + _dgbf
	if _fbfg {
		_cead = -_cead
	}
	return MakeNumberResult(_cead)
}

// MakeErrorResult constructs a #VALUE! error with a given extra error message.
// The error message is for debugging formula evaluation only and is not stored
// in the sheet.
func MakeErrorResult(msg string) Result { return MakeErrorResultType(ErrorTypeValue, msg) }

const _eccca = "\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065U\u006e\u006b\u006e\u006f\u0077\u006e\u0052\u0065\u0073u\u006c\u0074\u0054y\u0070\u0065\u004e\u0075\u006d\u0062\u0065\u0072\u0052\u0065s\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u0053\u0074\u0072\u0069\u006e\u0067\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u004c\u0069\u0073\u0074\u0052\u0065\u0073\u0075lt\u0054\u0079p\u0065\u0041r\u0072\u0061\u0079\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u0045\u0072\u0072\u006f\u0072\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u0045\u006d\u0070\u0074\u0079"

// Radians is an implementation of the Excel function RADIANS() that converts
// degrees to radians.
func Radians(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0052\u0041\u0044I\u0041\u004e\u0053\u0028)\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gfaad := args[0].AsNumber()
	if _gfaad.Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0044IA\u004e\u0053\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u006eu\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(_cc.Pi / 180.0 * _gfaad.ValueNumber)
}

// FunctionComplex is a function whose result  depends on its arguments and the
// context that it's in.  As an example, INDIRECT is a complex function so that
// INDIRECT("A1") which returns the value of the "A1" cell in a sheet can use
// the context to reach into the sheet and pull out required values.
type FunctionComplex func(_aacc Context, _gcad Evaluator, _fadfdg []Result) Result

// Reference returns a string reference value to a cell.
func (_gea CellRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeCell, Value: _gea._ebd}
}

const _ageaa = 57365

// SupportedFunctions returns a list of supported functions.
func SupportedFunctions() []string {
	_gdgdc := []string{}
	for _dbdf := range _feef {
		_gdgdc = append(_gdgdc, _dbdf)
	}
	for _becfa := range _gbbbb {
		_gdgdc = append(_gdgdc, _becfa)
	}
	_b.Strings(_gdgdc)
	return _gdgdc
}

const _edg = 86400000000000

// HorizontalRange is a range expression that when evaluated returns a list of Results from references like 1:4 (all cells from rows 1 to 4).
type HorizontalRange struct{ _cffd, _aaaaa int }

// Year is an implementation of the Excel YEAR() function.
func Year(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0045\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ecfe := ctx.GetEpoch()
	_afab, _efca := _eabc(args[0].Value(), _ecfe)
	if _efca != nil {
		return MakeErrorResult("\u0059\u0045AR\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s a\u0020si\u006e\u0067\u006c\u0065\u0020\u0064\u0061te\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeNumberResult(float64(_afab.Year()))
}

const _acea = "\u005e\u0028\u0028" + _afa + "\u007c" + _facd + "\u007c" + _dgb + "\u007c" + _fg + "\u0029\u0020\u0029\u003f"

func _dbcb(_cggdd string) string {
	_cggdd = _dda.Replace(_cggdd, "\u000a", "\u005c\u006e", -1)
	_cggdd = _dda.Replace(_cggdd, "\u000d", "\u005c\u0072", -1)
	_cggdd = _dda.Replace(_cggdd, "\u0009", "\u005c\u0074", -1)
	return _cggdd
}

const _ggf = "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0054\u0049\u004d\u0045\u0056\u0041\u004c\u0055\u0045"
const _geccf = 57362
const _adab = 57358

// Update updates references in the BinaryExpr after removing a row/column.
func (_baf BinaryExpr) Update(q *_a.UpdateQuery) Expression {
	_cec := _baf
	_cec._eb = _baf._eb.Update(q)
	_cec._bd = _baf._bd.Update(q)
	return _cec
}

// Sum is an implementation of the Excel SUM() function.
func Sum(args []Result) Result {
	_bcgf := MakeNumberResult(0)
	for _, _edege := range args {
		_edege = _edege.AsNumber()
		switch _edege.Type {
		case ResultTypeNumber:
			_bcgf.ValueNumber += _edege.ValueNumber
		case ResultTypeList, ResultTypeArray:
			_ggacc := Sum(_edege.ListValues())
			if _ggacc.Type != ResultTypeNumber {
				return _ggacc
			}
			_bcgf.ValueNumber += _ggacc.ValueNumber
		case ResultTypeString:
		case ResultTypeError:
			return _edege
		case ResultTypeEmpty:
		default:
			return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0053\u0055\u004d\u0028\u0029 \u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _edege.Type))
		}
	}
	return _bcgf
}

const _dbbbb int = 30

// Update updates references in the PrefixVerticalRange after removing a row/column.
func (_cagf PrefixVerticalRange) Update(q *_a.UpdateQuery) Expression {
	if q.UpdateType == _a.UpdateActionRemoveColumn {
		_aeeb := _cagf
		_cgge := _cagf._cbfc.String()
		if _cgge == q.SheetToUpdate {
			_acdgc := q.ColumnIdx
			_aeeb._dcdb = _acg(_cagf._dcdb, _acdgc)
			_aeeb._cbcf = _acg(_cagf._cbcf, _acdgc)
		}
		return _aeeb
	}
	return _cagf
}

type xargs struct {
	_cede []float64
	_eaff []float64
}

func _ce(_cee, _dg [][]Result) bool {
	if len(_cee) != len(_dg) {
		return false
	}
	for _dfb := range _cee {
		if len(_cee[_dfb]) != len(_dg[_dfb]) {
			return false
		}
	}
	return true
}

// MakeNumberResult constructs a number result.
func MakeNumberResult(v float64) Result {
	if v == _cc.Copysign(0, -1) {
		v = 0
	}
	return Result{Type: ResultTypeNumber, ValueNumber: v}
}

// Eval evaluates the binary expression using the context given.
func (_ae BinaryExpr) String() string {
	_ba := ""
	switch _ae._cb {
	case BinOpTypePlus:
		_ba = "\u002b"
	case BinOpTypeMinus:
		_ba = "\u002d"
	case BinOpTypeMult:
		_ba = "\u002a"
	case BinOpTypeDiv:
		_ba = "\u002f"
	case BinOpTypeExp:
		_ba = "\u005e"
	case BinOpTypeLT:
		_ba = "\u003c"
	case BinOpTypeGT:
		_ba = "\u003e"
	case BinOpTypeEQ:
		_ba = "\u003d"
	case BinOpTypeLEQ:
		_ba = "\u003c\u003d"
	case BinOpTypeGEQ:
		_ba = "\u003e\u003d"
	case BinOpTypeNE:
		_ba = "\u003c\u003e"
	case BinOpTypeConcat:
		_ba = "\u0026"
	}
	return _ae._eb.String() + _ba + _ae._bd.String()
}

// And is an implementation of the Excel AND() function.
func And(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0041\u004e\u0044 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061t\u0020l\u0065a\u0073t\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cefea := true
	for _, _bbcb := range args {
		_bbcb = _bbcb.AsNumber()
		switch _bbcb.Type {
		case ResultTypeList, ResultTypeArray:
			_cbga := And(_bbcb.ListValues())
			if _cbga.Type == ResultTypeError {
				return _cbga
			}
			if _cbga.ValueNumber == 0 {
				_cefea = false
			}
		case ResultTypeNumber:
			if _bbcb.ValueNumber == 0 {
				_cefea = false
			}
		case ResultTypeString:
			return MakeErrorResult("\u0041\u004e\u0044\u0020\u0064\u006f\u0065\u0073\u006e\u0027t\u0020\u006f\u0070\u0065\u0072\u0061\u0074e\u0020\u006f\u006e\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073")
		case ResultTypeError:
			return _bbcb
		default:
			return MakeErrorResult("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0061\u0072\u0067u\u006de\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0041\u004e\u0044")
		}
	}
	return MakeBoolResult(_cefea)
}

// IsLeapYear is an implementation of the Excel ISLEAPYEAR() function.
func IsLeapYear(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049S\u004c\u0045A\u0050\u0059\u0045\u0041R\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073in\u0067\u006c\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_defg := ctx.GetEpoch()
	_ebac, _efgf := _eabc(args[0].Value(), _defg)
	if _efgf != nil {
		return MakeErrorResult("\u0049S\u004c\u0045A\u0050\u0059\u0045\u0041R\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073in\u0067\u006c\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_fcec := _ebac.Year()
	return MakeBoolResult(_bac(_fcec))
}

const _gfbd = 57363

func _cecd(_cbgf []Result, _fcbg bool) Result {
	var _aead string
	if _fcbg {
		_aead = "\u004c\u0041\u0052G\u0045"
	} else {
		_aead = "\u0053\u004d\u0041L\u004c"
	}
	if len(_cbgf) != 2 {
		return MakeErrorResult(_aead + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_fbbe := _cbgf[0]
	var _fbfc [][]Result
	switch _fbbe.Type {
	case ResultTypeArray:
		_fbfc = _fbbe.ValueArray
	case ResultTypeList:
		_fbfc = [][]Result{_fbbe.ValueList}
	default:
		return MakeErrorResult(_aead + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u006f\u0066\u0020\u0074\u0079p\u0065\u0020a\u0072\u0072\u0061\u0079")
	}
	if len(_fbfc) == 0 {
		return MakeErrorResult(_aead + "\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0074\u006f\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0031\u0020\u0072\u006f\u0077")
	}
	if _cbgf[1].Type != ResultTypeNumber {
		return MakeErrorResult(_aead + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074yp\u0065\u0020\u006eu\u006db\u0065\u0072")
	}
	_efdf := _cbgf[1].ValueNumber
	if _efdf < 1 {
		return MakeErrorResultType(ErrorTypeNum, _aead+"\u0020\u0072e\u0071\u0075\u0069\u0072\u0065s\u0020\u0073\u0065\u0063\u006fn\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
	}
	_bggec := int(_efdf)
	if float64(_bggec) != _efdf {
		return MakeErrorResultType(ErrorTypeNum, _aead+"\u0020\u0072e\u0071\u0075\u0069\u0072\u0065s\u0020\u0073\u0065\u0063\u006fn\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
	}
	_deaa := []float64{}
	for _, _bggeg := range _fbfc {
		for _, _fcage := range _bggeg {
			if _fcage.Type == ResultTypeNumber {
				_deaa = append(_deaa, _fcage.ValueNumber)
			}
		}
	}
	if _bggec > len(_deaa) {
		return MakeErrorResultType(ErrorTypeNum, _aead+" \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u006c\u0065s\u0073\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u0068\u0061\u006e\u0020t\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u006e\u0075\u006d\u0062\u0065\u0072s\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cagd := _feb.MergeSort(_deaa)
	if _fcbg {
		return MakeNumberResult(_cagd[len(_cagd)-_bggec])
	} else {
		return MakeNumberResult(_cagd[_bggec-1])
	}
}

// Reference returns an invalid reference for ConstArrayExpr.
func (_af ConstArrayExpr) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Error is called in the case of parsing error and saves an error to a plex.
func (_acfgc *plex) Error(s string) {
	_bc.Log("\u0070a\u0072s\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", s)
	_acfgc._dccbf = s
}

// Max is an implementation of the Excel MAX() function.
func Max(args []Result) Result { return _dcfa(args, false) }

type noCache struct{}

// Offset is an implementation of the Excel OFFSET function.
func Offset(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 3 && len(args) != 5 {
		return MakeErrorResult("\u004f\u0046\u0046\u0053\u0045\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0074\u0068\u0072\u0065e\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_gcbg := args[0].Ref
	for _gcbg.Type == ReferenceTypeNamedRange {
		_gcbg = ctx.NamedRange(_gcbg.Value)
	}
	_edfff := ""
	switch _gcbg.Type {
	case ReferenceTypeCell:
		_edfff = _gcbg.Value
	case ReferenceTypeRange:
		_cfgfa := _dda.Split(_gcbg.Value, "\u003a")
		if len(_cfgfa) == 2 {
			_edfff = _cfgfa[0]
		}
	default:
		return MakeErrorResult(_d.Sprintf("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0069\u006e\u0020\u004f\u0046\u0046\u0053\u0045\u0054\u0028\u0029: \u0025\u0073", _gcbg.Type))
	}
	_egaf, _fcbe := _bb.ParseCellReference(_edfff)
	if _fcbe != nil {
		return MakeErrorResult(_d.Sprintf("\u0070\u0061\u0072s\u0065\u0020\u006f\u0072i\u0067\u0069\u006e\u0020\u0065\u0072\u0072o\u0072\u0020\u004f\u0046\u0046\u0053\u0045\u0054\u0028\u0029\u003a\u0020\u0025\u0073", _fcbe.Error()))
	}
	_cbag, _ffda, _badfd := _egaf.Column, _egaf.RowIdx, _egaf.SheetName
	_ecge := args[1].AsNumber()
	if _ecge.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0046\u0046SE\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020n\u0075m\u0065r\u0069\u0063\u0020\u0072\u006f\u0077\u0020\u006f\u0066\u0066\u0073\u0065\u0074")
	}
	_ccbc := args[2].AsNumber()
	if _ccbc.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0046\u0046SE\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020n\u0075m\u0065r\u0069\u0063\u0020\u0063\u006f\u006c\u0020\u006f\u0066\u0066\u0073\u0065\u0074")
	}
	var _gbdeb, _dfce Result
	if len(args) == 3 {
		_gbdeb = MakeNumberResult(1)
		_dfce = MakeNumberResult(1)
	} else {
		_gbdeb = args[3].AsNumber()
		if _gbdeb.Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0046\u0046\u0053\u0045\u0054\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u006e\u0075m\u0065\u0072\u0069\u0063\u0020\u0068\u0065\u0069\u0067\u0068\u0074")
		}
		if _gbdeb.ValueNumber == 0 {
			return MakeErrorResultType(ErrorTypeRef, "")
		}
		_dfce = args[4].AsNumber()
		if _dfce.Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0046F\u0053\u0045\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0077id\u0074\u0068")
		}
		if _dfce.ValueNumber == 0 {
			return MakeErrorResultType(ErrorTypeRef, "")
		}
	}
	_fdbfg := _bb.ColumnToIndex(_cbag)
	_dbfd := _ffda + uint32(_ecge.ValueNumber)
	_cegfd := _fdbfg + uint32(_ccbc.ValueNumber)
	_acca := _dbfd + uint32(_gbdeb.ValueNumber)
	_fbgbd := _cegfd + uint32(_dfce.ValueNumber)
	if _gbdeb.ValueNumber > 0 {
		_acca--
	} else {
		_acca++
		_dbfd, _acca = _acca, _dbfd
	}
	if _dfce.ValueNumber > 0 {
		_fbgbd--
	} else {
		_fbgbd++
		_cegfd, _fbgbd = _fbgbd, _cegfd
	}
	_ggee := _d.Sprintf("\u0025\u0073\u0025\u0064", _bb.IndexToColumn(_cegfd), _dbfd)
	_acgbe := _d.Sprintf("\u0025\u0073\u0025\u0064", _bb.IndexToColumn(_fbgbd), _acca)
	if _badfd == "" {
		return _cfcef(ctx, ev, _ggee, _acgbe)
	} else {
		return _cfcef(ctx.Sheet(_badfd), ev, _ggee, _acgbe)
	}
}

// NamedRangeRef is a reference to a named range.
type NamedRangeRef struct{ _bbfde string }

func (_fgcc *noCache) GetFromCache(key string) (Result, bool) { return _efc, false }

const _dgb = "\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u002d\u0028\u0028\u005b\u0030-\u0039]\u0029\u002b\u0029\u002d\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

// IsDBCS returns false for the invalid reference context.
func (_cfde *ivr) IsDBCS() bool { return false }
func _afgeb(_bccgf, _dgcb Expression) (Expression, Expression, error) {
	_cfdca, _dgfba := _bccgf.(CellRef)
	if !_dgfba {
		return nil, nil, _f.New(_d.Sprintf("\u0049\u006e\u0063\u006frr\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020%\u0073", _bccgf.String()))
	}
	_adbb, _dgfba := _dgcb.(CellRef)
	if !_dgfba {
		return nil, nil, _f.New(_d.Sprintf("\u0049\u006e\u0063\u006frr\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020%\u0073", _dgcb.String()))
	}
	_egged, _edcdg := _bb.ParseCellReference(_cfdca._ebd)
	if _edcdg != nil {
		return nil, nil, _edcdg
	}
	_afadb, _gdbe := _bb.ParseCellReference(_adbb._ebd)
	if _gdbe != nil {
		return nil, nil, _gdbe
	}
	_geaga := false
	if _egged.RowIdx > _afadb.RowIdx {
		_geaga = true
		_egged.RowIdx, _afadb.RowIdx = _afadb.RowIdx, _egged.RowIdx
	}
	if _egged.ColumnIdx > _afadb.ColumnIdx {
		_geaga = true
		_egged.ColumnIdx, _afadb.ColumnIdx = _afadb.ColumnIdx, _egged.ColumnIdx
		_egged.Column, _afadb.Column = _afadb.Column, _egged.Column
	}
	if _geaga {
		return NewCellRef(_egged.String()), NewCellRef(_afadb.String()), nil
	}
	return _bccgf, _dgcb, nil
}

// Eval evaluates a vertical range returning a list of results or an error.
func (_cgedb VerticalRange) Eval(ctx Context, ev Evaluator) Result {
	_eeegfb := _cgedb.verticalRangeReference()
	if _gegd, _geedf := ev.GetFromCache(_eeegfb); _geedf {
		return _gegd
	}
	_gcbf, _eedee := _ccfgf(ctx, _cgedb._cacbdb, _cgedb._cddc)
	_cgcfg := _cfcef(ctx, ev, _gcbf, _eedee)
	ev.SetCache(_eeegfb, _cgcfg)
	return _cgcfg
}

var _accf = [...]int{0, 7, 3, 3, 3, 8, 8, 8, 8, 1, 1, 1, 2, 2, 2, 2, 2, 14, 15, 15, 17, 17, 4, 4, 4, 13, 5, 6, 6, 6, 6, 6, 6, 6, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 9, 9, 9, 16, 16, 11, 10, 10}

// Reference returns an invalid reference for BinaryExpr.
func (_fag BinaryExpr) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }
func (_dca *defEval) checkLastEvalIsRef(_eae Context, _cce Expression) {
	switch _cce.(type) {
	case FunctionCall:
		switch _cce.(FunctionCall)._ffcfde {
		case "\u0049\u0053\u0052E\u0046":
			for _, _ggg := range _cce.(FunctionCall)._gace {
				switch _ggg.(type) {
				case CellRef, Range, HorizontalRange, VerticalRange, NamedRangeRef, PrefixExpr, PrefixRangeExpr, PrefixHorizontalRange, PrefixVerticalRange:
					_fceg := _ggg.Eval(_eae, _dca)
					_dca._efd = !(_fceg.Type == ResultTypeError && _fceg.ValueString == "\u0023\u004e\u0041\u004d\u0045\u003f")
				default:
					_dca._efd = false
				}
			}
		}
	}
}

// IsNA is an implementation of the Excel ISNA() function.
func IsNA(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u004e\u0041\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeError && args[0].ValueString == "\u0023\u004e\u002f\u0041")
}
func _acg(_gdb string, _gb uint32) string {
	_cfb := _bb.ColumnToIndex(_gdb)
	if _cfb == _gb {
		return "\u0023\u0052\u0045F\u0021"
	} else if _cfb > _gb {
		return _bb.IndexToColumn(_cfb - 1)
	} else {
		return _gdb
	}
}

// String returns a string representation of PrefixExpr.
func (_agdea PrefixExpr) String() string {
	return _d.Sprintf("\u0025\u0073\u0021%\u0073", _agdea._cacbf.String(), _agdea._bgagg.String())
}
func _fdb(_fcaaf, _dagb []float64, _bged float64) Result {
	_cage := false
	_fcege := false
	for _cegfe := 0; _cegfe < len(_fcaaf); _cegfe++ {
		if _fcaaf[_cegfe] > 0 {
			_cage = true
		}
		if _fcaaf[_cegfe] < 0 {
			_fcege = true
		}
	}
	if !_cage || !_fcege {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	_fced := _bged
	_ddafc := 1e-10
	_gbec := 0
	_agceg := 50
	_fgfc := false
	for {
		_cebe := _gff(_fcaaf, _dagb, _fced)
		_cfcgd := _fced - _cebe/_abab(_fcaaf, _dagb, _fced)
		_gfcg := _cc.Abs(_cfcgd - _fced)
		_fced = _cfcgd
		_gbec++
		if _gfcg <= _ddafc || _cc.Abs(_cebe) <= _ddafc {
			break
		}
		if _gbec > _agceg {
			_fgfc = true
			break
		}
	}
	if _fgfc || _cc.IsNaN(_fced) || _cc.IsInf(_fced, 0) {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	return MakeNumberResult(_fced)
}

// Update returns the same object as updating sheet references does not affect Number.
func (_dcbfa Number) Update(q *_a.UpdateQuery) Expression { return _dcbfa }

// String returns a string representation of a named range.
func (_gbcgg NamedRangeRef) String() string { return _gbcgg._bbfde }

// MaxIfs implements the MAXIFS function.
func MaxIfs(args []Result) Result {
	_abee := _gadg(args, true, "\u004d\u0041\u0058\u0049\u0046\u0053")
	if _abee.Type != ResultTypeEmpty {
		return _abee
	}
	_ccbd := _bdcae(args[1:])
	_cfaa := -_cc.MaxFloat64
	_eacb := _afgc(args[0])
	for _, _fefdd := range _ccbd {
		_agff := _eacb[_fefdd._cdag][_fefdd._gcbgb].ValueNumber
		if _cfaa < _agff {
			_cfaa = _agff
		}
	}
	if _cfaa == -_cc.MaxFloat64 {
		_cfaa = 0
	}
	return MakeNumberResult(float64(_cfaa))
}
func _geag(_dfba, _dcf, _add int) int {
	if _dfba > _dcf {
		return 0
	}
	if _bbbac(_add) {
		return (_dcf - _dfba + 1) * 360
	}
	_gbe := 0
	for _aabb := _dfba; _aabb <= _dcf; _aabb++ {
		_ggdb := 365
		if _bac(_aabb) {
			_ggdb = 366
		}
		_gbe += _ggdb
	}
	return _gbe
}
func _cggf(_daece Result, _gabg, _cfbc string) (string, Result) {
	switch _daece.Type {
	case ResultTypeString, ResultTypeNumber, ResultTypeEmpty:
		return _daece.Value(), _efc
	default:
		return "", MakeErrorResult(_gabg + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020" + _cfbc + "\u0020t\u006f\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006db\u0065r\u0020o\u0072\u0020\u0073\u0074\u0072\u0069\u006eg")
	}
}

const _ddeab int = 30

// NewPrefixRangeExpr constructs a new range with prefix.
func NewPrefixRangeExpr(pfx, from, to Expression) Expression {
	_cdddf, _fecga, _eabgg := _afgeb(from, to)
	if _eabgg != nil {
		_bc.Log(_eabgg.Error())
		return NewError(_eabgg.Error())
	}
	return PrefixRangeExpr{_afgf: pfx, _cbce: _cdddf, _dbdbg: _fecga}
}

// Update returns the same object as updating sheet references does not affect String.
func (_adeg String) Update(q *_a.UpdateQuery) Expression { return _adeg }

// Update updates references in the PrefixExpr after removing a row/column.
func (_fbbbg PrefixExpr) Update(q *_a.UpdateQuery) Expression {
	_ddffg := _fbbbg
	_efdge := _fbbbg._cacbf.String()
	if _efdge == q.SheetToUpdate {
		_geegd := *q
		_geegd.UpdateCurrentSheet = true
		_ddffg._bgagg = _fbbbg._bgagg.Update(&_geegd)
	}
	return _ddffg
}

// FloorMath implements _xlfn.FLOOR.MATH which rounds numbers down to the
// nearest multiple of the second argument, toward or away from zero as
// specified by the third argument.
func FloorMath(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0046\u004c\u004f\u004f\u0052\u002e\u004dA\u0054\u0048\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0046\u004c\u004f\u004f\u0052\u002e\u004dA\u0054\u0048\u0028)\u0020\u0061\u006c\u006co\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_eecf := args[0].AsNumber()
	if _eecf.Type != ResultTypeNumber {
		return MakeErrorResult("f\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020FL\u004f\u004f\u0052\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073t \u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_fcgc := float64(1)
	if _eecf.ValueNumber < 0 {
		_fcgc = -1
	}
	if len(args) > 1 {
		_dgbc := args[1].AsNumber()
		if _dgbc.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061r\u0067\u0075\u006den\u0074\u0020\u0074\u006f\u0020\u0046L\u004f\u004f\u0052\u002e\u004d\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006db\u0065\u0072")
		}
		_fcgc = _dgbc.ValueNumber
	}
	_bddd := float64(1)
	if len(args) > 2 {
		_cgca := args[2].AsNumber()
		if _cgca.Type != ResultTypeNumber {
			return MakeErrorResult("t\u0068\u0069\u0072\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020FL\u004f\u004f\u0052\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073t \u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_bddd = _cgca.ValueNumber
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Floor(_eecf.ValueNumber))
	}
	_efdb := _eecf.ValueNumber
	_efdb, _feeb := _cc.Modf(_efdb / _fcgc)
	if _feeb != 0 && _eecf.ValueNumber < 0 && _bddd > 0 {
		_efdb++
	}
	return MakeNumberResult(_efdb * _fcgc)
}

// Context is a formula execution context.  Formula evaluation uses the context
// to retreive information from sheets.
type Context interface {

	// Cell returns the result of evaluating a cell.
	Cell(_cea string, _afb Evaluator) Result

	// Sheet returns an evaluation context for a given sheet name.  This is used
	// when evaluating cells that pull data from other sheets (e.g. ='Sheet 2'!A1).
	Sheet(_abc string) Context

	// GetEpoch returns the time epoch of the context's Workbook.
	GetEpoch() _ee.Time

	// GetFilename returns the full filename of the context's Workbook.
	GetFilename() string

	// GetWidth returns a worksheet's column width.
	GetWidth(_dad int) float64

	// GetFormat returns a cell's format.
	GetFormat(_dcc string) string

	// GetLabelPrefix returns cell's label prefix dependent on cell horizontal alignment.
	GetLabelPrefix(_cba string) string

	// GetFormat returns if cell is protected.
	GetLocked(_ef string) bool

	// HasFormula returns if cell contains formula.
	HasFormula(_db string) bool

	// IsBool returns if cell contains boolean value.
	IsBool(_dfc string) bool

	// IsDBCS returns if workbook default language is among DBCS.
	IsDBCS() bool

	// LastColumn returns the name of last column which contains data in range of context sheet's given rows.
	LastColumn(_dgf, _be int) string

	// LastRow returns the name of last row which contains data in range of context sheet's given columns.
	LastRow(_dba string) int

	// SetLocked returns sets cell's protected attribute.
	SetLocked(_ged string, _bbb bool)

	// NamedRange returns a named range.
	NamedRange(_bdb string) Reference

	// SetOffset is used so that the Context can evaluate cell references
	// differently when they are not absolute (e.g. not like '$A$5').  See the
	// shared formula support in Cell for usage.
	SetOffset(_adg, _ade uint32)
}
type cumulArgs struct {
	_egab float64
	_cefc float64
	_dfaf float64
	_gfdb float64
	_ggcf float64
	_fgcf int
}

// Update updates the FunctionCall references after removing a row/column.
func (_dfggg FunctionCall) Update(q *_a.UpdateQuery) Expression {
	_abddb := []Expression{}
	for _, _aceb := range _dfggg._gace {
		_gccd := _aceb.Update(q)
		_abddb = append(_abddb, _gccd)
	}
	return FunctionCall{_ffcfde: _dfggg._ffcfde, _gace: _abddb}
}

const _adga = 57367

// CountBlank implements the COUNTBLANK function.
func CountBlank(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0043\u004f\u0055N\u0054\u0042\u004c\u0041N\u004b\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u006e\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(_bggea(args, _dcbfd))
}

// String returns an empty string for EmptyExpr.
func (_fdd EmptyExpr) String() string { return "" }
func Unicode(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0055\u004e\u0049\u0043\u004fD\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_edccf := args[0].AsString()
	if _edccf.Type != ResultTypeString {
		return MakeErrorResult("\u0055\u004e\u0049\u0043\u004fD\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(_edccf.ValueString) == 0 {
		return MakeErrorResult("\u0055\u004e\u0049\u0043\u004f\u0044\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073 \u0061\u0020\u006e\u006f\u006e\u002d\u007a\u0065\u0072\u006f\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(float64(_edccf.ValueString[0]))
}
func _caaf(_bgabf Result, _dedba *criteriaParsed) bool {
	if _bgabf.IsBoolean {
		return false
	}
	_fgdcc := _bgabf.Type
	if _dedba._ddba {
		return _fgdcc == ResultTypeNumber && _bgabf.ValueNumber == _dedba._gdfd
	} else if _fgdcc == ResultTypeNumber {
		return _dcaba(_bgabf.ValueNumber, _dedba._cdcgg)
	}
	return _cefeb(_bgabf, _dedba)
}

const _befd = 57361

// Pv implements the Excel PV function.
func Pv(args []Result) Result {
	_bcc := len(args)
	if _bcc < 3 || _bcc > 5 {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u00205")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_dffd := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_eecd := args[1].ValueNumber
	if _eecd != float64(int(_eecd)) {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0061\u0079\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fabe := args[2].ValueNumber
	_edcd := 0.0
	if _bcc >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0056 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0075\u0074\u0075\u0072\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_edcd = args[3].ValueNumber
	}
	_gddc := 0.0
	if _bcc == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_gddc = args[4].ValueNumber
		if _gddc != 0 {
			_gddc = 1
		}
	}
	if _dffd == 0 {
		return MakeNumberResult(-_fabe*_eecd - _edcd)
	} else {
		return MakeNumberResult((((1-_cc.Pow(1+_dffd, _eecd))/_dffd)*_fabe*(1+_dffd*_gddc) - _edcd) / _cc.Pow(1+_dffd, _eecd))
	}
}

// Eval evaluates a range with prefix returning a list of results or an error.
func (_fcfef PrefixRangeExpr) Eval(ctx Context, ev Evaluator) Result {
	_dgecb := _fcfef._afgf.Reference(ctx, ev)
	_egece := _fcfef._cbce.Reference(ctx, ev)
	_bagc := _fcfef._dbdbg.Reference(ctx, ev)
	switch _dgecb.Type {
	case ReferenceTypeSheet:
		if _fgba(_dgecb, ctx) {
			return MakeErrorResultType(ErrorTypeName, _d.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _dgecb.Value))
		}
		_ebcf := _adec(_dgecb, _egece, _bagc)
		if _egece.Type == ReferenceTypeCell && _bagc.Type == ReferenceTypeCell {
			if _fadcb, _bcag := ev.GetFromCache(_ebcf); _bcag {
				return _fadcb
			} else {
				_gcbe := _cfcef(ctx.Sheet(_dgecb.Value), ev, _egece.Value, _bagc.Value)
				ev.SetCache(_ebcf, _gcbe)
				return _gcbe
			}
		}
		return MakeErrorResult("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072a\u006e\u0067\u0065\u0020" + _ebcf)
	default:
		return MakeErrorResult(_d.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _dgecb.Type))
	}
}

// Bool is a boolean expression.
type Bool struct{ _fd bool }

// String returns a string of a range.
func (_bcbag Range) String() string {
	return _d.Sprintf("\u0025\u0073\u003a%\u0073", _bcbag._ecca.String(), _bcbag._ddgga.String())
}
func _edebb(_gdbf []Result) (bool, Result) {
	for _, _cffaa := range _gdbf {
		if _cffaa.Type == ResultTypeError {
			return true, _cffaa
		}
	}
	return false, MakeEmptyResult()
}

// Reference returns a string reference value to a vertical range with prefix.
func (_gbaded PrefixVerticalRange) Reference(ctx Context, ev Evaluator) Reference {
	_dfgfd := _gbaded._cbfc.Reference(ctx, ev)
	return Reference{Type: ReferenceTypeVerticalRange, Value: _gbaded.verticalRangeReference(_dfgfd.Value)}
}
func (_geagg tokenType) String() string { return _cecbd(int(_geagg)) }

// Evaluator is the interface for a formula evaluator.  This is needed so we can
// pass it to the spreadsheet to let it evaluate formula cells before returning
// the results.
// NOTE: in order to implement Evaluator without cache embed noCache in it.
type Evaluator interface {
	Eval(_eef Context, formula string) Result
	SetCache(_gag string, _aee Result)
	GetFromCache(_ffa string) (Result, bool)
	LastEvalIsRef() bool
}

// NewBool constructs a new boolean expression.
func NewBool(v string) Expression {
	_ac, _fcc := _c.ParseBool(v)
	if _fcc != nil {
		_bc.Log("\u0065\u0072\u0072\u006f\u0072\u0020p\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0075\u006ca\u0020\u0062\u006f\u006f\u006c\u0020\u0025s\u003a\u0020\u0025\u0073", v, _fcc)
	}
	return Bool{_fd: _ac}
}

type defEval struct {
	evCache
	_efd bool
}

func _fgggg(_afda Result, _dgbfd int) []Result {
	_aadg := []Result{}
	switch _afda.Type {
	case ResultTypeList:
		_efff := _afda.ValueList
		_gbade := len(_efff)
		for _agbe := 0; _agbe < _dgbfd; _agbe++ {
			if _agbe < _gbade {
				_aadg = append(_aadg, _efff[_agbe])
			} else {
				_aadg = append(_aadg, MakeErrorResultType(ErrorTypeNA, ""))
			}
		}
	case ResultTypeNumber, ResultTypeString, ResultTypeError, ResultTypeEmpty:
		for _cdddc := 0; _cdddc < _dgbfd; _cdddc++ {
			_aadg = append(_aadg, _afda)
		}
	}
	return _aadg
}

// Month is an implementation of the Excel MONTH() function.
func Month(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("M\u004f\u004e\u0054\u0048\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006ee\u0020\u0061\u0072g\u0075m\u0065\u006e\u0074")
	}
	_fbe := args[0]
	switch _fbe.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(1)
	case ResultTypeNumber:
		_afbb := _dcge(_fbe.ValueNumber)
		return MakeNumberResult(float64(_afbb.Month()))
	case ResultTypeString:
		_bdbb := _dda.ToLower(_fbe.ValueString)
		if !_gedc(_bdbb) {
			_, _, _, _, _bgg, _fcga := _eafd(_bdbb)
			if _fcga.Type == ResultTypeError {
				_fcga.ErrorMessage = "\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u004dON\u0054\u0048"
				return _fcga
			}
			if _bgg {
				return MakeNumberResult(1)
			}
		}
		_, _edc, _, _, _edf := _gaf(_bdbb)
		if _edf.Type == ResultTypeError {
			return _edf
		}
		return MakeNumberResult(float64(_edc))
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u004d\u004fN\u0054\u0048")
	}
}

// Pricedisc implements the Excel PRICEDISC function.
func Pricedisc(args []Result) Result {
	_cecb := len(args)
	if _cecb != 4 && _cecb != 5 {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u006f\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_abbf, _ccacc, _afadd := _gdge(args[0], args[1], "\u0050R\u0049\u0043\u0045\u0044\u0049\u0053C")
	if _afadd.Type == ResultTypeError {
		return _afadd
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049C\u0045\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079p\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ceaf := args[2].ValueNumber
	if _ceaf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049C\u0045\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050R\u0049\u0043E\u0044\u0049\u0053\u0043 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	_fdab := args[3].ValueNumber
	if _fdab <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050R\u0049\u0043E\u0044\u0049\u0053\u0043 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006fs\u0069\u0074i\u0076\u0065")
	}
	_gbadd := 0
	if _cecb == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0052I\u0043\u0045\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_gbadd = int(args[4].ValueNumber)
		if !_gaff(_gbadd) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0050R\u0049\u0043E\u0044\u0049\u0053\u0043")
		}
	}
	_ddd, _afadd := _agef(_abbf, _ccacc, _gbadd)
	if _afadd.Type == ResultTypeError {
		return _afadd
	}
	return MakeNumberResult(_fdab * (1 - _ceaf*_ddd))
}
func _dccg(_bada, _aaac _ee.Time, _eage int) _ee.Time {
	_ecda := _ee.Date(_bada.Year(), _aaac.Month(), _aaac.Day(), 0, 0, 0, 0, _ee.UTC)
	if _ecda.After(_bada) {
		_ecda = _ecda.AddDate(-1, 0, 0)
	}
	for !_ecda.After(_bada) {
		_ecda = _ecda.AddDate(0, 12/_eage, 0)
	}
	return _ecda
}
func _adec(_ffdg, _cgag, _bfbaa Reference) string {
	return _d.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _ffdg.Value, _cgag.Value, _bfbaa.Value)
}

// NewNegate constructs a new negate expression.
func NewNegate(e Expression) Expression { return Negate{_eefg: e} }

// Range is a range expression that when evaluated returns a list of Results.
type Range struct{ _ecca, _ddgga Expression }

// Code is an implementation of the Excel CODE function that returns the first
// character of the string as a number.
func Code(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0043\u004f\u0044\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gdaf := args[0].AsString()
	if _gdaf.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u004f\u0044\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	if len(_gdaf.ValueString) == 0 {
		return MakeNumberResult(0)
	}
	return MakeNumberResult(float64(_gdaf.ValueString[0]))
}

const (
	_fedg countMode = iota
	_faeea
	_dcbfd
)

func _dccb(_bdfe []Result, _ffe string) (*couponArgs, Result) {
	_fgb := len(_bdfe)
	if _fgb != 3 && _fgb != 4 {
		return nil, MakeErrorResult(_ffe + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u006f\u0072\u0020\u0066o\u0075\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_ffbb, _dgag, _dgga := _gdge(_bdfe[0], _bdfe[1], _ffe)
	if _dgga.Type == ResultTypeError {
		return nil, _dgga
	}
	if _bdfe[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_ffe + "\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0066\u0072\u0065\u0071\u0075\u0065\u006e\u0063\u0079 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_geaag := _bdfe[2].ValueNumber
	if !_gabe(_geaag) {
		return nil, MakeErrorResult("\u0049n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0066\u0072\u0065q\u0075\u0065\u006e\u0063\u0079\u0020\u0066\u006f\u0072\u0020" + _ffe)
	}
	_feag := 0
	if _fgb == 4 && _bdfe[3].Type != ResultTypeEmpty {
		if _bdfe[3].Type != ResultTypeNumber {
			return nil, MakeErrorResult(_ffe + "\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020b\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_feag = int(_bdfe[3].ValueNumber)
		if !_gaff(_feag) {
			return nil, MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020fo\u0072\u0020"+_ffe)
		}
	}
	return &couponArgs{_ffbb, _dgag, int(_geaag), _feag}, _efc
}

// Power is an implementation of the Excel POWER function that raises a number
// to a power. It requires two numeric arguments.
func Power(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0050\u004f\u0057\u0045\u0052\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_eead := args[0].AsNumber()
	if _eead.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0050\u004f\u0057\u0045\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ggade := args[1].AsNumber()
	if _ggade.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0050\u004f\u0057\u0045\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	return MakeNumberResult(_cc.Pow(_eead.ValueNumber, _ggade.ValueNumber))
}

// Min is an implementation of the Excel MIN() function.
func Min(args []Result) Result { return _cgdb(args, false) }

var _dace = [...]int{45, 3, 44, 32, 18, 40, 72, 46, 47, 30, 31, 32, 39, 48, 28, 29, 30, 31, 32, 75, 39, 49, 32, 56, 50, 70, 23, 39, 76, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 77, 71, 69, 54, 43, 13, 19, 21, 55, 82, 11, 78, 9, 74, 28, 29, 30, 31, 32, 37, 33, 34, 35, 36, 38, 1, 20, 39, 10, 2, 8, 0, 80, 79, 0, 0, 0, 83, 0, 81, 73, 28, 29, 30, 31, 32, 37, 33, 34, 35, 36, 38, 0, 0, 39, 28, 29, 30, 31, 32, 37, 33, 34, 35, 36, 38, 26, 27, 39, 51, 52, 25, 14, 15, 16, 17, 0, 24, 23, 22, 41, 23, 12, 0, 6, 7, 26, 27, 0, 42, 0, 25, 14, 15, 16, 17, 0, 24, 23, 22, 5, 0, 12, 0, 6, 7, 26, 27, 0, 4, 0, 25, 14, 15, 16, 17, 0, 24, 23, 22, 41, 0, 12, 53, 6, 7, 26, 27, 0, 0, 0, 25, 14, 15, 16, 17, 0, 24, 23, 22, 41, 0, 12, 0, 6, 7}

// SeriesSum implements the Excel SERIESSUM function.
func SeriesSum(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("\u0053\u0045\u0052\u0049\u0045\u0053\u0053\u0055\u004d\u0028\u0029\u0020\u0072\u0065\u0071u\u0069r\u0065\u0073\u0020\u0034\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_cgaa := args[0].AsNumber()
	_dedb := args[1].AsNumber()
	_eegd := args[2].AsNumber()
	_bdgc := args[3].ListValues()
	if _cgaa.Type != ResultTypeNumber || _dedb.Type != ResultTypeNumber || _eegd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0045\u0052\u0049\u0045\u0053S\u0055\u004d\u0028)\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073t\u0020\u0074\u0068\u0072\u0065e \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063")
	}
	_fcbf := float64(0)
	for _gccc, _gbfd := range _bdgc {
		_fcbf += _gbfd.ValueNumber * _cc.Pow(_cgaa.ValueNumber, _dedb.ValueNumber+float64(_gccc)*_eegd.ValueNumber)
	}
	return MakeNumberResult(_fcbf)
}

// String returns a string representation of a vertical range with prefix.
func (_fgbgb PrefixVerticalRange) String() string {
	return _d.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _fgbgb._cbfc.String(), _fgbgb._dcdb, _fgbgb._cbcf)
}

const (
	BinOpTypeUnknown BinOpType = iota
	BinOpTypePlus
	BinOpTypeMinus
	BinOpTypeMult
	BinOpTypeDiv
	BinOpTypeExp
	BinOpTypeLT
	BinOpTypeGT
	BinOpTypeEQ
	BinOpTypeLEQ
	BinOpTypeGEQ
	BinOpTypeNE
	BinOpTypeConcat
)

func _ebacd(_gcf Result, _gcebc, _dbedc int) [][]Result {
	_dbgec := [][]Result{}
	switch _gcf.Type {
	case ResultTypeArray:
		for _abffd, _abfb := range _gcf.ValueArray {
			if _abffd < _gcebc {
				_dbgec = append(_dbgec, _fgggg(MakeListResult(_abfb), _dbedc))
			} else {
				_dbgec = append(_dbgec, _fgggg(MakeErrorResultType(ErrorTypeNA, ""), _dbedc))
			}
		}
	case ResultTypeList:
		_fgfd := _fgggg(_gcf, _dbedc)
		for _abgb := 0; _abgb < _gcebc; _abgb++ {
			_dbgec = append(_dbgec, _fgfd)
		}
	case ResultTypeNumber, ResultTypeString, ResultTypeError, ResultTypeEmpty:
		for _eafa := 0; _eafa < _gcebc; _eafa++ {
			_cfbf := _fgggg(_gcf, _dbedc)
			_dbgec = append(_dbgec, _cfbf)
		}
	}
	return _dbgec
}

// MakeListResult constructs a list result.
func MakeListResult(list []Result) Result { return Result{Type: ResultTypeList, ValueList: list} }
func _cgg(_fccf string, _eba *_a.UpdateQuery) string {
	_ebe, _bgd := _bb.ParseCellReference(_fccf)
	if _bgd != nil {
		return "\u0023\u0052\u0045F\u0021"
	}
	if _eba.UpdateType == _a.UpdateActionRemoveColumn {
		_fce := _eba.ColumnIdx
		_ddb := _ebe.ColumnIdx
		if _ddb < _fce {
			return _fccf
		} else if _ddb == _fce {
			return "\u0023\u0052\u0045F\u0021"
		} else {
			return _ebe.Update(_a.UpdateActionRemoveColumn).String()
		}
	}
	return _fccf
}

// RoundUp is an implementation of the Excel ROUNDUP function that rounds a number
// up to a specified number of digits.
func RoundUp(args []Result) Result { return _ceacf(args, _ffcc) }

// Negate is a negate expression like -A1.
type Negate struct{ _eefg Expression }
type node struct {
	_eddb tokenType
	_aaca string
}

// DateDif is an implementation of the Excel DATEDIF() function.
func DateDif(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeString {
		return MakeErrorResult("\u0044\u0041\u0054\u0045\u0044I\u0046\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077o\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u006e\u0064\u0020\u006f\u006e\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_ebba := args[0].ValueNumber
	_fcce := args[1].ValueNumber
	if _fcce < _ebba {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0068\u0065\u0020\u0073\u0074\u0061r\u0074\u0020\u0064a\u0074\u0065\u0020\u0069s\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u0065\u006e\u0064\u0020\u0064\u0061\u0074\u0065")
	}
	if _fcce == _ebba {
		return MakeNumberResult(0)
	}
	_cacb := _dda.ToLower(args[2].ValueString)
	if _cacb == "\u0064" {
		return MakeNumberResult(_fcce - _ebba)
	}
	_ccea := _dcge(_ebba)
	_fdda := _dcge(_fcce)
	_dac, _ffde, _ebf := _ccea.Date()
	_cbg, _abdf, _aga := _fdda.Date()
	_cfa := int(_ffde)
	_ecd := int(_abdf)
	var _cfd float64
	switch _cacb {
	case "\u0079":
		_cfd = float64(_cbg - _dac)
		if _ecd < _cfa || (_ecd == _cfa && _aga < _ebf) {
			_cfd--
		}
	case "\u006d":
		_gfe := _cbg - _dac
		_fccc := _ecd - _cfa
		if _aga < _ebf {
			_fccc--
		}
		if _fccc < 0 {
			_gfe--
			_fccc += 12
		}
		_cfd = float64(_gfe*12 + _fccc)
	case "\u006d\u0064":
		_agd := _ecd
		if _aga < _ebf {
			_agd--
		}
		_cfd = float64(int(_fcce - _bad(_cbg, _agd, _ebf)))
	case "\u0079\u006d":
		_cfd = float64(_ecd - _cfa)
		if _aga < _ebf {
			_cfd--
		}
		if _cfd < 0 {
			_cfd += 12
		}
	case "\u0079\u0064":
		_dbf := _cbg
		if _ecd < _cfa || (_ecd == _cfa && _aga < _ebf) {
			_dbf--
		}
		_cfd = float64(int(_fcce - _bad(_dbf, _cfa, _ebf)))
	default:
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0069\u006e\u0074e\u0072\u0076\u0061\u006c\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	return MakeNumberResult(_cfd)
}
func (_fef *evCache) GetFromCache(key string) (Result, bool) {
	_fef._ddg.Lock()
	_baa, _fca := _fef._ada[key]
	_fef._ddg.Unlock()
	return _baa, _fca
}
func _eeab(_fggc []string, _gffd int) string { return _c.Itoa(len(_fggc[len(_fggc)-1-_gffd])) }

// MaxA is an implementation of the Excel MAXA() function.
func MaxA(args []Result) Result { return _dcfa(args, true) }

// Update returns the same object as updating sheet references does not affect named ranges.
func (_faafe NamedRangeRef) Update(q *_a.UpdateQuery) Expression { return _faafe }

// IsBool returns false for the invalid reference context.
func (_dced *ivr) IsBool(cellRef string) bool { return false }

var _aeed = [...]int{1}

// NewString constructs a new string expression.
func NewString(v string) Expression {
	v = _dda.Replace(v, "\u0022\u0022", "\u0022", -1)
	return String{_adfgc: v}
}

// Today is an implementation of the Excel TODAY() function.
func Today(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("\u0054\u004f\u0044A\u0059\u0020\u0064\u006fe\u0073\u006e\u0027\u0074\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_bgdee := _ee.Now()
	_, _cad := _bgdee.Zone()
	_fagc := _fcb(_dgae, _bgdee.Unix()+int64(_cad)) + 1
	return MakeNumberResult(_fagc)
}

// IsLogical is an implementation of the Excel ISLOGICAL() function.
func IsLogical(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u004c\u004f\u0047\u0049\u0043A\u004c\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gbge := args[0].Ref
	if _gbge.Type != ReferenceTypeCell {
		return MakeErrorResult("I\u0053\u004c\u004f\u0047\u0049\u0043\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073t\u0020a\u0072\u0067\u0075\u006de\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
	}
	return MakeBoolResult(ctx.Cell(_gbge.Value, ev).IsBoolean)
}
func _fegfd(_ebbce, _cdcad Reference) string {
	return _d.Sprintf("\u0025\u0073\u003a%\u0073", _ebbce.Value, _cdcad.Value)
}

// MakeArrayResult constructs an array result (matrix).
func MakeArrayResult(arr [][]Result) Result { return Result{Type: ResultTypeArray, ValueArray: arr} }
func _afde(_febg, _cbfff float64) float64 {
	_febg = _cc.Trunc(_febg)
	_cbfff = _cc.Trunc(_cbfff)
	if _febg == 0 {
		return _cbfff
	}
	if _cbfff == 0 {
		return _febg
	}
	for _febg != _cbfff {
		if _febg > _cbfff {
			_febg = _febg - _cbfff
		} else {
			_cbfff = _cbfff - _febg
		}
	}
	return _febg
}

// LookupFunctionComplex looks up and returns a complex function or nil.
func LookupFunctionComplex(name string) FunctionComplex {
	_cfbcf.Lock()
	defer _cfbcf.Unlock()
	if _gfdc, _bdde := _gbbbb[name]; _bdde {
		return _gfdc
	}
	return nil
}

// Char is an implementation of the Excel CHAR function that takes an integer in
// the range [0,255] and returns the corresponding ASCII character.
func Char(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0043\u0048\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gecg := args[0].AsNumber()
	if _gecg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0043\u0048\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_acfec := int(_gecg.ValueNumber)
	if _acfec < 0 || _acfec > 255 {
		return MakeErrorResult("\u0043H\u0041\u0052 \u0072\u0065\u0071\u0075i\u0072\u0065\u0073 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073 i\u006e\u0020\u0074h\u0065\u0020r\u0061\u006e\u0067\u0065\u0020\u005b0\u002c\u00325\u0035\u005d")
	}
	return MakeStringResult(_d.Sprintf("\u0025\u0063", _acfec))
}

type ivr struct{}

const _affd = 57344

func (_abcee *ivr) Cell(ref string, ev Evaluator) Result {
	return MakeErrorResult("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065")
}

// Dollarde implements the Excel DOLLARDE function.
func Dollarde(args []Result) Result {
	_aeb, _cged, _afbg := _fdcd(args, "\u0044\u004f\u004c\u004c\u0041\u0052\u0044\u0045")
	if _afbg.Type == ResultTypeError {
		return _afbg
	}
	if _cged < 1 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0044\u004f\u004c\u004c\u0041\u0052\u0044\u0045\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0072a\u0063t\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u006f\u0072 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0031")
	}
	if _aeb == 0 {
		return MakeNumberResult(0)
	}
	_eefe := _aeb < 0
	if _eefe {
		_aeb = -_aeb
	}
	_gbgd := args[0].Value()
	_bfdd := _dda.Split(_gbgd, "\u002e")
	_eagc := float64(int(_aeb))
	_aac := _bfdd[1]
	_baae := len(_aac)
	_cfaf := int(_cc.Log10(_cged)) + 1
	_dgab := float64(_cfaf - _baae)
	_bcbd, _ecaf := _c.ParseFloat(_aac, 64)
	if _ecaf != nil {
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006fr \u0044\u004f\u004cL\u0041R\u0044\u0045")
	}
	_bcbd *= _cc.Pow(10, _dgab)
	_bdce := _eagc + _bcbd/_cged
	if _eefe {
		_bdce = -_bdce
	}
	return MakeNumberResult(_bdce)
}

// Update updates references in the PrefixHorizontalRange after removing a row/column.
func (_bdab PrefixHorizontalRange) Update(q *_a.UpdateQuery) Expression { return _bdab }

const _febdc = -1000

func _cdac() {
	_ebcaag = _ge.MustCompile("\u005e\u0030\u002b\u0024")
	_faac = _ge.MustCompile("\u005e\u0028\u0028\u0023|0\u0029\u002b\u002c\u0029\u002b\u0028\u0023\u007c\u0030\u0029\u002b\u0028\u003b\u007c$\u0029")
	_caffa = _ge.MustCompile("\u005e\u0028\u0023\u007c\u0030\u007c\u002c\u0029\u002a\u005f\u005c\u0029\u003b")
	_dbce = _ge.MustCompile("\u005e\u0030\u002b\u005c\u002e\u0028\u0030\u002b\u0029\u0024")
	_fefd = _ge.MustCompile("\u005e\u0028\u0028\u0023\u007c\u0030\u0029\u002b\u002c\u0029+\u0028\u0023\u007c\u0030\u0029\u002b\u005c.\u0028\u0030\u002b\u0029\u002e\u002a\u0028\u003b\u007c\u0024\u0029")
	_fefaa = _ge.MustCompile("^\u0028\u005f\u007c\u002d\u007c\u0020)\u002b\u005c\u002a\u0020\u0023\u002b\u002c\u0023\u002b0\u005c\u002e\u00280\u002b)\u002e\u002a\u003b")
	_cbgd = _ge.MustCompile("\u005e\u0028\u0028\u0023\u007c\u0030)\u002b\u002c\u0029\u002b\u0028\u0023\u007c\u0030\u0029\u002b\u005c\u002e\u0028(\u0023\u007c\u0030\u0029\u002b\u0029\u005f\\\u0029\u002e\u002a\u003b")
	_ddcb = _ge.MustCompile("\u005e\u0028\u0023\u007c0)\u002b\u005c\u002e\u0028\u0028\u0023\u007c\u0030\u0029\u002b\u0029\u0025\u0024")
	_bdbfg = _ge.MustCompile("\u005c\u005b\u005c$\u005c\u0024\u002d\u002e+\u005c\u005d\u0028\u005c\u002a\u0020\u0029?\u0028\u0023\u007c\u0030\u0029\u002b\u002c\u0028\u0023\u007c\u0030\u0029\u002b\u003b")
	_egb = _ge.MustCompile("\u005c[\u005c\u0024\\\u0024\u002d\u002e+\u005c\u005d\u0028\u005c\u002a\u0020\u0029?\u0028\u0023\u007c\u0030\u0029\u002b,\u0028\u0023\u007c\u0030\u0029\u002b\u005c\u002e\u0028\u0028\u0023|\u0030\u007c\u002d\u0029\u002b\u0029\u002e\u002a\u003b")
	_fefc = _ge.MustCompile("\u005e(\u0028\u0023|\u0030\u0029\u002b,\u0029\u002b\u0028\u0023\u007c\u0030\u0029+\u0028\u005c\u002e\u0028\u0028\u0023|\u0030\u007c\u002d\u0029\u002b\u0029\u0029\u003f\u002e\u002b\u005c[\u005c\u0024\u002e\u002b\u005c\u005d\u002e\u002a\u003b")
	_eefed = _ge.MustCompile("\u005e\u004d\u002b(\u002f\u007c\u0020\u007c\u002c\u007c\u0022\u007c" + _eafdd + _eafdd + "\u0029\u002b\u0044\u002b\u0028\u002f\u007c\u0020\u007c\u002c\u007c\u0022\u007c" + _eafdd + _eafdd + "\u0029\u002b\u0059+\u0024")
	_cdadb = _ge.MustCompile("\u005e\u0044\u002b\u0028\u002f\u007c\u0020\u007c\u005c\u002e\u007c\u0022\u007c" + _eafdd + _eafdd + "\u0029\u002b\u004d\u002b\u0028\u002f\u007c\u0020\u007c\\\u002e\u007c\u0022\u007c" + _eafdd + _eafdd + "\u0029\u002b\u0059+\u0024")
	_bacfb = _ge.MustCompile("\u005e\u0028\u0023|\u0030\u0029\u002b\u005c.\u0028\u0028\u0023\u007c\u0030\u0029\u002a)\u0045\u005c\u002b\u0028\u0023\u007c\u0030\u0029\u002b\u0028\u003b\u007c\u0024\u0029")
	_cdfb = _ge.MustCompile("\u005e.\u002a\u005f\u005c\u0029\u002e\u002a;")
}

type criteriaParsed struct {
	_ddba  bool
	_gdfd  float64
	_dcfga string
	_cdcgg *criteriaRegex
}

// Result is the result of a formula or cell evaluation .
type Result struct {
	ValueNumber  float64
	ValueString  string
	ValueList    []Result
	ValueArray   [][]Result
	IsBoolean    bool
	ErrorMessage string
	Type         ResultType
	Ref          Reference
}
type yySymType struct {
	_bddb  int
	_bdeac *node
	_ddfc  Expression
	_fedd  []Expression
	_ddga  [][]Expression
}

// ReferenceType is a type of reference
//go:generate stringer -type=ReferenceType
type ReferenceType byte

// TextJoin is an implementation of the Excel TEXTJOIN function.
func TextJoin(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("\u0054\u0045\u0058\u0054\u004aO\u0049\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074h\u0072\u0065\u0065\u0020\u006f\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeString {
		return MakeErrorResult("\u0054\u0045\u0058T\u004a\u004f\u0049\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0065\u006c\u0069\u006d\u0069\u0074\u0065\u0072\u0020\u0074\u006f\u0020\u0062\u0065 \u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	_gfgdga := args[0].ValueString
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0045\u0058\u0054\u004a\u004f\u0049\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065c\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0072 \u0062\u006f\u006f\u006c\u0065a\u006e")
	}
	_bbcbg := args[1].ValueNumber != 0
	_cedgf := _fdfd(args[2:], []string{}, _bbcbg)
	return MakeStringResult(_dda.Join(_cedgf, _gfgdga))
}

// Effect implements the Excel EFFECT function.
func Effect(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0045\u0046F\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0045\u0046\u0046\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u006f\u006d\u0069n\u0061\u006c\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_bcgbg := args[0].ValueNumber
	if _bcgbg <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0045\u0046\u0046\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u006f\u006d\u0069n\u0061\u006c\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0045\u0046\u0046\u0045\u0043\u0054 \u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0066 \u0063\u006f\u006d\u0070\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020p\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074")
	}
	_afdf := float64(int(args[1].ValueNumber))
	if _afdf < 1 {
		return MakeErrorResultType(ErrorTypeNum, "E\u0046\u0046\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063o\u006dp\u006f\u0075\u006e\u0064i\u006e\u0067 \u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0031\u0020\u006f\u0072\u0020\u006d\u006f\u0072\u0065")
	}
	return MakeNumberResult(_cc.Pow((1+_bcgbg/_afdf), _afdf) - 1)
}
func _cgdb(_geecd []Result, _feffa bool) Result {
	_ebgca := "\u004d\u0049\u004e"
	if _feffa {
		_ebgca = "\u004d\u0049\u004e\u0041"
	}
	if len(_geecd) == 0 {
		return MakeErrorResult(_ebgca + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_abea := _cc.MaxFloat64
	for _, _ccacb := range _geecd {
		switch _ccacb.Type {
		case ResultTypeNumber:
			if (_feffa || !_ccacb.IsBoolean) && _ccacb.ValueNumber < _abea {
				_abea = _ccacb.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			_cbdbd := _cgdb(_ccacb.ListValues(), _feffa)
			if _cbdbd.ValueNumber < _abea {
				_abea = _cbdbd.ValueNumber
			}
		case ResultTypeEmpty:
		case ResultTypeString:
			_caca := 0.0
			if _feffa {
				_caca = _ccacb.AsNumber().ValueNumber
			}
			if _caca < _abea {
				_abea = _caca
			}
		default:
			_bc.Log("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020"+_ebgca+"\u0028\u0029\u0020\u0061rg\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _ccacb.Type)
		}
	}
	if _abea == _cc.MaxFloat64 {
		_abea = 0
	}
	return MakeNumberResult(_abea)
}

const _fdc = "\u0028(\u005b0\u002d\u0039\u005d\u0029\u002b)\u0020\u0028a\u006d\u007c\u0070\u006d\u0029"

// RandBetween is an implementation of the Excel RANDBETWEEN() function that returns a random
// integer in the range specified.
func RandBetween(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0052A\u004e\u0044B\u0045\u0054\u0057\u0045E\u004e\u0028\u0029 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020tw\u006f\u0020\u006eu\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_bdec := args[0].AsNumber()
	_eedec := args[1].AsNumber()
	if _bdec.Type != ResultTypeNumber || _eedec.Type != ResultTypeNumber {
		return MakeErrorResult("\u0052A\u004e\u0044B\u0045\u0054\u0057\u0045E\u004e\u0028\u0029 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020tw\u006f\u0020\u006eu\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	if _eedec.ValueNumber < _bdec.ValueNumber {
		return MakeErrorResult("\u0052\u0041\u004e\u0044\u0042E\u0054\u0057\u0045\u0045\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065r")
	}
	_fbca := int64(_bdec.ValueNumber)
	_cbfe := int64(_eedec.ValueNumber)
	return MakeNumberResult(float64(_dcee.Int63n(_cbfe-_fbca+1) + _fbca))
}
func _dfbag(_fdbe float64) float64 {
	_dgfg := float64(1)
	for _gddd := float64(2); _gddd <= _fdbe; _gddd++ {
		_dgfg *= _gddd
	}
	return _dgfg
}

// Sln implements the Excel SLN function.
func Sln(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0053\u004c\u004e\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065e\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u004c\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_baeb := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u004c\u004e \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dbaea := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u004c\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_dcgcg := args[2].ValueNumber
	if _dcgcg == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0053\u004c\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u007a\u0065\u0072\u006f")
	}
	return MakeNumberResult((_baeb - _dbaea) / _dcgcg)
}

var InvalidReferenceContext = &ivr{}

func _cggfe(_ceacg, _fbgef int) string {
	const TOKSTART = 4
	if !_cdaf {
		return "\u0073\u0079\u006et\u0061\u0078\u0020\u0065\u0072\u0072\u006f\u0072"
	}
	for _, _cafcc := range _eafb {
		if _cafcc._agffb == _ceacg && _cafcc._agbag == _fbgef {
			return "\u0073\u0079\u006e\u0074\u0061\u0078\u0020\u0065\u0072r\u006f\u0072\u003a\u0020" + _cafcc._febb
		}
	}
	_bbffea := "\u0073y\u006e\u0074\u0061\u0078 \u0065\u0072\u0072\u006f\u0072:\u0020u\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064 " + _cecbd(_fbgef)
	_dcda := make([]int, 0, 4)
	_dbaab := _ecgbe[_ceacg]
	for _ddddc := TOKSTART; _ddddc-1 < len(_fdcdab); _ddddc++ {
		if _bfgfec := _dbaab + _ddddc; _bfgfec >= 0 && _bfgfec < _efaeb && _egbc[_dace[_bfgfec]] == _ddddc {
			if len(_dcda) == cap(_dcda) {
				return _bbffea
			}
			_dcda = append(_dcda, _ddddc)
		}
	}
	if _dgge[_ceacg] == -2 {
		_faeb := 0
		for _edde[_faeb] != -1 || _edde[_faeb+1] != _ceacg {
			_faeb += 2
		}
		for _faeb += 2; _edde[_faeb] >= 0; _faeb += 2 {
			_egdc := _edde[_faeb]
			if _egdc < TOKSTART || _edde[_faeb+1] == 0 {
				continue
			}
			if len(_dcda) == cap(_dcda) {
				return _bbffea
			}
			_dcda = append(_dcda, _egdc)
		}
		if _edde[_faeb+1] != 0 {
			return _bbffea
		}
	}
	for _gcgd, _dceec := range _dcda {
		if _gcgd == 0 {
			_bbffea += "\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020"
		} else {
			_bbffea += "\u0020\u006f\u0072\u0020"
		}
		_bbffea += _cecbd(_dceec)
	}
	return _bbffea
}

var _dfg = []*_ge.Regexp{}

// AsNumber attempts to intepret a string cell value as a number. Upon success,
// it returns a new number result, upon  failure it returns the original result.
// This is used as functions return strings that can then act like number (e.g.
// LEFT(1.2345,3) + LEFT(1.2345,3) = 2.4)
func (_efefc Result) AsNumber() Result {
	if _efefc.Type == ResultTypeString {
		_ggeaa, _deae := _c.ParseFloat(_efefc.ValueString, 64)
		if _deae == nil {
			return MakeNumberResult(_ggeaa)
		}
	}
	if _efefc.Type == ResultTypeEmpty {
		return MakeNumberResult(0)
	}
	return _efefc
}

// Fact is an implementation of the excel FACT function which returns the
// factorial of a positive numeric input.
func Fact(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("F\u0041\u0043\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_cdbe := args[0].AsNumber()
	if _cdbe.Type != ResultTypeNumber {
		return MakeErrorResult("F\u0041\u0043\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	if _cdbe.ValueNumber < 0 {
		return MakeErrorResult("\u0046\u0041\u0043\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u006f\u006e\u006c\u0079\u0020\u0070\u006f\u0073\u0069t\u0069\u0076\u0065\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	return MakeNumberResult(_dfbag(_cdbe.ValueNumber))
}

type parsedSearchObject struct {
	_fcgea string
	_bfgd  string
	_bdgb  int
}

// NewHorizontalRange constructs a new full rows range.
func NewHorizontalRange(v string) Expression {
	_cffdd := _dda.Split(v, "\u003a")
	if len(_cffdd) != 2 {
		return nil
	}
	_cdee, _ := _c.Atoi(_cffdd[0])
	_gbab, _ := _c.Atoi(_cffdd[1])
	if _cdee > _gbab {
		_cdee, _gbab = _gbab, _cdee
	}
	return HorizontalRange{_cffd: _cdee, _aaaaa: _gbab}
}

// Syd implements the Excel SYD function.
func Syd(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("S\u0059\u0044\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072 \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cebec := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cffc := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_afcd := args[2].ValueNumber
	if _afcd <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0053\u0059\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gceb := args[3].ValueNumber
	if _gceb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0053\u0059\u0044 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070e\u0072i\u006fd\u0020t\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if _gceb > _afcd {
		return MakeErrorResultType(ErrorTypeNum, "\u0053\u0059\u0044\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0065q\u0075\u0061\u006c\u0020\u006f\u0072\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068a\u006e \u006c\u0069\u0066\u0065")
	}
	_dcae := (_cebec - _cffc) * (_afcd - _gceb + 1) * 2
	_dbebb := _afcd * (_afcd + 1)
	return MakeNumberResult(_dcae / _dbebb)
}

// SumIfs implements the SUMIFS function.
func SumIfs(args []Result) Result {
	_eeffe := _gadg(args, true, "\u0053\u0055\u004d\u0049\u0046\u0053")
	if _eeffe.Type != ResultTypeEmpty {
		return _eeffe
	}
	_beed := _bdcae(args[1:])
	_gagdd := 0.0
	_agedb := _afgc(args[0])
	for _, _cadgg := range _beed {
		_gagdd += _agedb[_cadgg._cdag][_cadgg._gcbgb].ValueNumber
	}
	return MakeNumberResult(float64(_gagdd))
}

// Degrees is an implementation of the Excel function DEGREES() that converts
// radians to degrees.
func Degrees(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0044\u0045\u0047R\u0045\u0045\u0053\u0028)\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fccgg := args[0].AsNumber()
	if _fccgg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0045\u0047RE\u0045\u0053\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u006eu\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(180.0 / _cc.Pi * _fccgg.ValueNumber)
}

// PrefixExpr is an expression containing reference to another sheet like Sheet1!A1 (the value of the cell A1 from sheet 'Sheet1').
type PrefixExpr struct {
	_cacbf Expression
	_bgagg Expression
}

// Rate implements the Excel RATE function.
func Rate(args []Result) Result {
	_ddfe := len(args)
	if _ddfe < 3 || _ddfe > 6 {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0072\u0065\u0065 \u0061\u006e\u0064\u0020\u0073i\u0078")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_egfd := args[0].ValueNumber
	if _egfd != float64(int(_egfd)) {
		return MakeErrorResultType(ErrorTypeNum, "R\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072i\u006fd\u0073\u0020\u0074\u006f \u0062\u0065 \u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0070\u0061y\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_baed := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_ggabf := args[2].ValueNumber
	_bbdb := 0.0
	if _ddfe >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_bbdb = args[3].ValueNumber
	}
	_fgga := 0.0
	if _ddfe >= 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("R\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_fgga = args[4].ValueNumber
		if _fgga != 0 {
			_fgga = 1
		}
	}
	_eabf := 0.1
	if _ddfe >= 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_eabf = args[5].ValueNumber
	}
	_bfbf := 100
	_dfbd := 0
	_gfce := false
	_edb := 1e-6
	_cbba := _eabf
	for _dfbd < _bfbf && !_gfce {
		_fggdd := _cc.Pow(_cbba+1, _egfd)
		_fddg := _cc.Pow(_cbba+1, _egfd-1)
		_dcgea := _cbba*_fgga + 1
		_afc := _baed * (_fggdd - 1)
		_cbbgd := _bbdb + _fggdd*_ggabf + _afc*_dcgea/_cbba
		_agae := _egfd*_fddg*_ggabf - _afc*_dcgea/_cc.Pow(_cbba, 2)
		_dbfg := (_egfd*_baed*_fddg*_dcgea + _afc*_fgga) / _cbba
		_eafc := _cbbgd / (_agae + _dbfg)
		if _cc.Abs(_eafc) < _edb {
			_gfce = true
		}
		_dfbd++
		_cbba -= _eafc
	}
	return MakeNumberResult(_cbba)
}

const _dffa = _ee.Second * 1

func (_ddbbd Result) String() string { return _ddbbd.Value() }

const _cfaee = _ee.Millisecond * 1000

// Reference returns a string reference value to a range with prefix.
func (_egc PrefixRangeExpr) Reference(ctx Context, ev Evaluator) Reference {
	_efebdc := _egc._afgf.Reference(ctx, ev)
	_gaeb := _egc._cbce.Reference(ctx, ev)
	_ffcdc := _egc._dbdbg.Reference(ctx, ev)
	if _efebdc.Type == ReferenceTypeSheet && _gaeb.Type == ReferenceTypeCell && _ffcdc.Type == ReferenceTypeCell {
		return MakeRangeReference(_adec(_efebdc, _gaeb, _ffcdc))
	}
	return ReferenceInvalid
}

const _eegb = 57349

// Odd is an implementation of the Excel ODD() that rounds a number to the
// nearest odd integer.
func Odd(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("O\u0044\u0044\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006ee\u0020\u0061\u0072g\u0075m\u0065\u006e\u0074")
	}
	_dcba := args[0].AsNumber()
	if _dcba.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fgdc := _cc.Signbit(_dcba.ValueNumber)
	_dcfg, _fedfg := _cc.Modf((_dcba.ValueNumber - 1) / 2)
	_bdbe := _dcfg*2 + 1
	if _fedfg != 0 {
		if !_fgdc {
			_bdbe += 2
		} else {
			_bdbe -= 2
		}
	}
	return MakeNumberResult(_bdbe)
}

// FactDouble is an implementation of the excel FACTDOUBLE function which
// returns the double factorial of a positive numeric input.
func FactDouble(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0046\u0041C\u0054\u0044\u004f\u0055\u0042\u004c\u0045\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ggfg := args[0].AsNumber()
	if _ggfg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0041C\u0054\u0044\u004f\u0055\u0042\u004c\u0045\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _ggfg.ValueNumber < 0 {
		return MakeErrorResult("\u0046A\u0043\u0054D\u004f\u0055\u0042\u004cE\u0028\u0029\u0020a\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u006f\u006ely\u0020\u0070\u006fs\u0069\u0074i\u0076\u0065\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_fcfc := float64(1)
	_bbdfd := _cc.Trunc(_ggfg.ValueNumber)
	for _deeg := _bbdfd; _deeg > 1; _deeg -= 2 {
		_fcfc *= _deeg
	}
	return MakeNumberResult(_fcfc)
}

// SumSquares is an implementation of the Excel SUMSQ() function.
func SumSquares(args []Result) Result {
	_aagd := MakeNumberResult(0)
	for _, _ffdd := range args {
		_ffdd = _ffdd.AsNumber()
		switch _ffdd.Type {
		case ResultTypeNumber:
			_aagd.ValueNumber += _ffdd.ValueNumber * _ffdd.ValueNumber
		case ResultTypeList, ResultTypeArray:
			_cbfec := SumSquares(_ffdd.ListValues())
			if _cbfec.Type != ResultTypeNumber {
				return _cbfec
			}
			_aagd.ValueNumber += _cbfec.ValueNumber
		case ResultTypeString:
		case ResultTypeError:
			return _ffdd
		case ResultTypeEmpty:
		default:
			return MakeErrorResult(_d.Sprintf("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0053\u0055\u004dS\u0051\u0055\u0041\u0052\u0045\u0053(\u0029\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079p\u0065\u0020\u0025\u0073", _ffdd.Type))
		}
	}
	return _aagd
}

// VerticalRange is a range expression that when evaluated returns a list of Results from references like AA:IJ (all cells from columns AA to IJ).
type VerticalRange struct{ _cacbdb, _cddc string }

const _edge = 57348

// Eval evaluates the binary expression using the context given.
func (_bbd BinaryExpr) Eval(ctx Context, ev Evaluator) Result {
	_fac := _bbd._eb.Eval(ctx, ev)
	if _fac.Type == ResultTypeError {
		return _fac
	}
	_gc := _bbd._bd.Eval(ctx, ev)
	if _gc.Type == ResultTypeError {
		return _gc
	}
	if _fac.Type == _gc.Type {
		if _fac.Type == ResultTypeArray {
			if !_ce(_fac.ValueArray, _gc.ValueArray) {
				return MakeErrorResult("l\u0068\u0073\u002f\u0072\u0068\u0073 \u0073\u0068\u006f\u0075\u006c\u0064 \u0068\u0061\u0076\u0065\u0020\u0073\u0061m\u0065\u0020\u0064\u0069\u006d\u0065\u006e\u0073\u0069\u006fn\u0073")
			}
			return _ab(_bbd._cb, _fac.ValueArray, _gc.ValueArray)
		} else if _fac.Type == ResultTypeList {
			if len(_fac.ValueList) != len(_gc.ValueList) {
				return MakeErrorResult("l\u0068\u0073\u002f\u0072\u0068\u0073 \u0073\u0068\u006f\u0075\u006c\u0064 \u0068\u0061\u0076\u0065\u0020\u0073\u0061m\u0065\u0020\u0064\u0069\u006d\u0065\u006e\u0073\u0069\u006fn\u0073")
			}
			return _fae(_bbd._cb, _fac.ValueList, _gc.ValueList)
		}
	} else if _fac.Type == ResultTypeArray && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
		return _ca(_bbd._cb, _fac.ValueArray, _gc)
	} else if _fac.Type == ResultTypeList && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
		return _ga(_bbd._cb, _fac.ValueList, _gc)
	}
	switch _bbd._cb {
	case BinOpTypePlus:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeNumberResult(_fac.ValueNumber + _gc.ValueNumber)
			}
		}
	case BinOpTypeMinus:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeNumberResult(_fac.ValueNumber - _gc.ValueNumber)
			}
		}
	case BinOpTypeMult:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeNumberResult(_fac.ValueNumber * _gc.ValueNumber)
			}
		}
	case BinOpTypeDiv:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				if _gc.ValueNumber == 0 {
					return MakeErrorResultType(ErrorTypeDivideByZero, "\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0062\u0079 \u007a\u0065\u0072\u006f")
				}
				return MakeNumberResult(_fac.ValueNumber / _gc.ValueNumber)
			}
		}
	case BinOpTypeExp:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeNumberResult(_cc.Pow(_fac.ValueNumber, _gc.ValueNumber))
			}
		}
	case BinOpTypeLT:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeBoolResult(_fac.ValueNumber < _gc.ValueNumber)
			}
			if _fac.Type == ResultTypeString {
				return MakeBoolResult(_fac.ValueString < _gc.ValueString)
			}
			if _fac.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if _fac.Type == ResultTypeString && _gc.Type == ResultTypeNumber {
			return MakeBoolResult(false)
		} else if _fac.Type == ResultTypeNumber && _gc.Type == ResultTypeString {
			return MakeBoolResult(true)
		} else if _fac.Type == ResultTypeEmpty && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
			return MakeBoolResult(true)
		} else if (_fac.Type == ResultTypeNumber || _fac.Type == ResultTypeString) && _gc.Type == ResultTypeEmpty {
			return MakeBoolResult(false)
		}
	case BinOpTypeGT:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeBoolResult(_fac.ValueNumber > _gc.ValueNumber)
			}
			if _fac.Type == ResultTypeString {
				return MakeBoolResult(_fac.ValueString > _gc.ValueString)
			}
			if _fac.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if _fac.Type == ResultTypeString && _gc.Type == ResultTypeNumber {
			return MakeBoolResult(true)
		} else if _fac.Type == ResultTypeNumber && _gc.Type == ResultTypeString {
			return MakeBoolResult(false)
		} else if _fac.Type == ResultTypeEmpty && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
			return MakeBoolResult(false)
		} else if (_fac.Type == ResultTypeNumber || _fac.Type == ResultTypeString) && _gc.Type == ResultTypeEmpty {
			return MakeBoolResult(true)
		}
	case BinOpTypeEQ:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeBoolResult(_fac.ValueNumber == _gc.ValueNumber)
			}
			if _fac.Type == ResultTypeString {
				return MakeBoolResult(_fac.ValueString == _gc.ValueString)
			}
			if _fac.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if (_fac.Type == ResultTypeString && _gc.Type == ResultTypeNumber) || (_fac.Type == ResultTypeNumber && _gc.Type == ResultTypeString) {
			return MakeBoolResult(false)
		} else if _fac.Type == ResultTypeEmpty && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
			return MakeBoolResult(_bcg(_gc))
		} else if (_fac.Type == ResultTypeNumber || _fac.Type == ResultTypeString) && _gc.Type == ResultTypeEmpty {
			return MakeBoolResult(_bcg(_fac))
		}
	case BinOpTypeNE:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeBoolResult(_fac.ValueNumber != _gc.ValueNumber)
			}
			if _fac.Type == ResultTypeString {
				return MakeBoolResult(_fac.ValueString != _gc.ValueString)
			}
			if _fac.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if (_fac.Type == ResultTypeString && _gc.Type == ResultTypeNumber) || (_fac.Type == ResultTypeNumber && _gc.Type == ResultTypeString) {
			return MakeBoolResult(true)
		} else if _fac.Type == ResultTypeEmpty && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
			return MakeBoolResult(!_bcg(_gc))
		} else if (_fac.Type == ResultTypeNumber || _fac.Type == ResultTypeString) && _gc.Type == ResultTypeEmpty {
			return MakeBoolResult(!_bcg(_fac))
		}
	case BinOpTypeLEQ:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeBoolResult(_fac.ValueNumber <= _gc.ValueNumber)
			}
			if _fac.Type == ResultTypeString {
				return MakeBoolResult(_fac.ValueString <= _gc.ValueString)
			}
			if _fac.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if _fac.Type == ResultTypeString && _gc.Type == ResultTypeNumber {
			return MakeBoolResult(false)
		} else if _fac.Type == ResultTypeNumber && _gc.Type == ResultTypeString {
			return MakeBoolResult(true)
		} else if _fac.Type == ResultTypeEmpty && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
			return MakeBoolResult(_bcg(_gc))
		} else if (_fac.Type == ResultTypeNumber || _fac.Type == ResultTypeString) && _gc.Type == ResultTypeEmpty {
			return MakeBoolResult(_bcg(_fac))
		}
	case BinOpTypeGEQ:
		if _fac.Type == _gc.Type {
			if _fac.Type == ResultTypeNumber {
				return MakeBoolResult(_fac.ValueNumber >= _gc.ValueNumber)
			}
			if _fac.Type == ResultTypeString {
				return MakeBoolResult(_fac.ValueString >= _gc.ValueString)
			}
			if _fac.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if _fac.Type == ResultTypeString && _gc.Type == ResultTypeNumber {
			return MakeBoolResult(true)
		} else if _fac.Type == ResultTypeNumber && _gc.Type == ResultTypeString {
			return MakeBoolResult(false)
		} else if _fac.Type == ResultTypeEmpty && (_gc.Type == ResultTypeNumber || _gc.Type == ResultTypeString) {
			return MakeBoolResult(_bcg(_gc))
		} else if (_fac.Type == ResultTypeNumber || _fac.Type == ResultTypeString) && _gc.Type == ResultTypeEmpty {
			return MakeBoolResult(_bcg(_fac))
		}
	case BinOpTypeConcat:
		return MakeStringResult(_fac.Value() + _gc.Value())
	}
	return MakeErrorResult("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006f\u0070")
}

// ISTEXT is an implementation of the Excel ISTEXT() function.
func IsText(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0054EX\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070t\u0073 \u0061 \u0073i\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeString)
}

// Coupnum implements the Excel COUPNUM function.
func Coupnum(args []Result) Result {
	_gfb, _gafc := _dccb(args, "\u0043O\u0055\u0050\u004e\u0055\u004d")
	if _gafc.Type == ResultTypeError {
		return _gafc
	}
	_ccfcf := _gfb._dfff
	_fdgb := _gfb._cfeb
	_ege, _gafc := _fbeb(_gfb._eac, _gfb._aea, _ccfcf, _fdgb)
	if _gafc.Type == ResultTypeError {
		return _gafc
	}
	return MakeNumberResult(_ege)
}

// YearFrac is an implementation of the Excel YEARFRAC() function.
func YearFrac(args []Result) Result {
	_fdeb := len(args)
	if (_fdeb != 2 && _fdeb != 3) || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber {
		return MakeErrorResult("Y\u0045\u0041\u0052\u0046\u0052\u0041\u0043\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020tw\u006f\u0020\u006f\u0072 \u0074\u0068\u0072\u0065\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_edffc := 0
	if _fdeb == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("Y\u0045\u0041\u0052\u0046\u0052\u0041\u0043\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020ba\u0073\u0069\u0073\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074o \u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_edffc = int(args[2].ValueNumber)
		if !_gaff(_edffc) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0059\u0045\u0041R\u0046\u0052\u0041\u0043")
		}
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0045\u0041\u0052\u0046\u0052\u0041\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0074\u0061\u0072\u0074\u0020\u0064\u0061t\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_afe := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0045\u0041\u0052\u0046\u0052\u0041\u0043 \u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u0065\u006ed\u0020\u0064\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_eddg := args[1].ValueNumber
	_gbad, _adb := _agef(_afe, _eddg, _edffc)
	if _adb.Type == ResultTypeError {
		return _adb
	}
	return MakeNumberResult(_gbad)
}

// CountIfs implements the COUNTIFS function.
func CountIfs(args []Result) Result {
	_dgbce := _gadg(args, false, "\u0043\u004f\u0055\u004e\u0054\u0049\u0046\u0053")
	if _dgbce.Type != ResultTypeEmpty {
		return _dgbce
	}
	_bgagb := _bdcae(args)
	return MakeNumberResult(float64(len(_bgagb)))
}

// NewRange constructs a new range.
func NewRange(from, to Expression) Expression {
	_dfabc, _fbebe, _ffbe := _afgeb(from, to)
	if _ffbe != nil {
		_bc.Log(_ffbe.Error())
		return NewError(_ffbe.Error())
	}
	return Range{_ecca: _dfabc, _ddgga: _fbebe}
}

// Multinomial implements the excel MULTINOMIAL function.
func Multinomial(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u004d\u0055\u004c\u0054\u0049\u004eO\u004d\u0049\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006ce\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c\u0020i\u006e\u0070\u0075\u0074")
	}
	_abbgc, _eafab, _eagf := _cfbd(args)
	if _eagf.Type == ResultTypeError {
		return _eagf
	}
	return MakeNumberResult(_dfbag(_abbgc) / _eafab)
}

// Counta implements the COUNTA function.
func Counta(args []Result) Result { return MakeNumberResult(_bggea(args, _faeea)) }

// GetEpoch returns a null time object for the invalid reference context.
func (_fbef *ivr) GetEpoch() _ee.Time { return _ee.Time{} }

// Not is an implementation of the Excel NOT() function and takes a single
// argument.
func Not(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004eO\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	switch args[0].Type {
	case ResultTypeError:
		return args[0]
	case ResultTypeString, ResultTypeList:
		return MakeErrorResult("\u004e\u004f\u0054\u0020\u0065\u0078\u0070\u0065\u0063\u0074s\u0020\u0061\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	case ResultTypeNumber:
		return MakeBoolResult(!(args[0].ValueNumber != 0))
	default:
		return MakeErrorResult("u\u006e\u0068\u0061\u006e\u0064\u006ce\u0064\u0020\u004e\u004f\u0054\u0020\u0061\u0072\u0067u\u006d\u0065\u006et\u0020t\u0079\u0070\u0065")
	}
}
func _fagd(_dacc, _gec _ee.Time, _aegd int) float64 {
	if _dacc.After(_gec) {
		_dacc, _gec = _gec, _dacc
	}
	_cecee := 0
	_bdbg, _eege, _cbf := _dacc.Date()
	_fffc, _gefe, _dade := _gec.Date()
	_gdg, _fec := int(_eege), int(_gefe)
	_bfae, _cda := _gad(_bdbg, _gdg, _cbf, _aegd), _gad(_fffc, _fec, _dade, _aegd)
	if !_bbbac(_aegd) {
		return _bad(_fffc, _fec, _cda) - _bad(_bdbg, _gdg, _bfae)
	}
	if _aegd == 0 {
		if (_gdg == 2 || _bfae < 30) && _dade == 31 {
			_cda = 31
		} else if _fec == 2 && _cda == _cdf(_fffc, _fec) {
			_cda = _cdf(_fffc, 2)
		}
	} else {
		if _gdg == 2 && _bfae == 30 {
			_bfae = _cdf(_bdbg, 2)
		}
		if _fec == 2 && _cda == 30 {
			_cda = _cdf(_fffc, 2)
		}
	}
	if _bdbg < _fffc || (_bdbg == _fffc && _gdg < _fec) {
		_cecee = 30 - _bfae + 1
		_cbf = 1
		_bfae = 1
		_fabb := _ee.Date(_bdbg, _ee.Month(_gdg), _cbf, 0, 0, 0, 0, _ee.UTC).AddDate(0, 1, 0)
		if _fabb.Year() < _fffc {
			_cecee += _bgbc(_fabb.Year(), int(_fabb.Month()), 12, _aegd)
			_fabb = _fabb.AddDate(0, 13-int(_fabb.Month()), 0)
			_cecee += _geag(_fabb.Year(), _fffc-1, _aegd)
		}
		_cecee += _bgbc(_fffc, int(_fabb.Month()), _fec-1, _aegd)
		_fabb = _fabb.AddDate(0, _fec-int(_fabb.Month()), 0)
		_gdg = _fabb.Day()
	}
	_cecee += _cda - _bfae
	if _cecee > 0 {
		return float64(_cecee)
	} else {
		return 0
	}
}
func _gfcfe(_ddecg []Result, _bafb bool) (float64, float64) {
	_feea := 0.0
	_affa := 0.0
	for _, _gbaa := range _ddecg {
		switch _gbaa.Type {
		case ResultTypeNumber:
			if _bafb || !_gbaa.IsBoolean {
				_affa += _gbaa.ValueNumber
				_feea++
			}
		case ResultTypeList, ResultTypeArray:
			_bffb, _ffgb := _gfcfe(_gbaa.ListValues(), _bafb)
			_affa += _bffb
			_feea += _ffgb
		case ResultTypeString:
			if _bafb {
				_feea++
			}
		case ResultTypeEmpty:
		}
	}
	return _affa, _feea
}

const _ebg = "\u0028\u0028\u005b0\u002d\u0039\u005d\u0029\u002b\u0029\u003a\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u005c\u002e\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u0028\u0020(\u0061\u006d\u007c\u0070\u006d\u0029\u0029\u003f"
const (
	_fceda rmode = iota
	_dgacc
	_ffcc
)

func (_dga *evCache) SetCache(key string, value Result) {
	_dga._ddg.Lock()
	_dga._ada[key] = value
	_dga._ddg.Unlock()
}

// Xnpv implements the Excel XNPV function.
func Xnpv(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0058\u004eP\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("X\u004e\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_cafc := args[0].ValueNumber
	if _cafc <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0058\u004e\u0050\u0056\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u0076\u0065")
	}
	_daab, _dbcf := _fdfc(args[1], args[2], "\u0058\u004e\u0050\u0056")
	if _dbcf.Type == ResultTypeError {
		return _dbcf
	}
	_edeb := _daab._cede
	_gaaf := _daab._eaff
	_fecb := 0.0
	_gagd := _gaaf[0]
	for _dcbc, _efgd := range _edeb {
		_fecb += _efgd / _cc.Pow(1+_cafc, (_gaaf[_dcbc]-_gagd)/365)
	}
	return MakeNumberResult(_fecb)
}

// Nper implements the Excel NPER function.
func Nper(args []Result) Result {
	_fdba := len(args)
	if _fdba < 3 || _fdba > 5 {
		return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067e\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u0020\u0035")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("N\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_dcea := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0070\u0061y\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dgaf := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_gfbg := args[2].ValueNumber
	_cgb := 0.0
	if _fdba >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_cgb = args[3].ValueNumber
	}
	_gegg := 0.0
	if _fdba == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("N\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_gegg = args[4].ValueNumber
		if _gegg != 0 {
			_gegg = 1
		}
	}
	_feca := _dgaf*(1+_dcea*_gegg) - _cgb*_dcea
	_cefg := (_gfbg*_dcea + _dgaf*(1+_dcea*_gegg))
	return MakeNumberResult(_cc.Log(_feca/_cefg) / _cc.Log(1+_dcea))
}

// String returns a string representation of a horizontal range with prefix.
func (_fgdfd PrefixHorizontalRange) String() string {
	return _d.Sprintf("\u0025\u0073\u0021\u0025\u0064\u003a\u0025\u0064", _fgdfd._fdcfe.String(), _fgdfd._ffba, _fgdfd._cebd)
}
func _bggf(_ecbg Result, _cafb *criteriaParsed) bool {
	if _ecbg.Type == ResultTypeEmpty {
		return false
	}
	if _cafb._ddba {
		return _ecbg.ValueNumber == _cafb._gdfd
	} else {
		_abdeb := _dda.ToLower(_ecbg.ValueString)
		return _cafb._dcfga == _abdeb || _ec.Match(_cafb._dcfga, _abdeb)
	}
}

var _cfbcf _fe.Mutex

// Time is an implementation of the Excel TIME() function.
func Time(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0049ME\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s t\u0068re\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_fbd := args[0].ValueNumber
	_ega := args[1].ValueNumber
	_fgd := args[2].ValueNumber
	_bdfb := _aabf(_fbd, _ega, _fgd)
	if _bdfb >= 0 {
		return MakeNumberResult(_bdfb)
	} else {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
}
func _ecac(_cfgf float64) float64 { return float64(int(_cfgf + 0.5)) }

const _dbfe = 57359
const _eaea = 57354

// Yielddisc implements the Excel YIELDDISC function.
func Yielddisc(args []Result) Result {
	_gfbf := len(args)
	if _gfbf != 4 && _gfbf != 5 {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u006f\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_agbg, _cgaf, _aad := _gdge(args[0], args[1], "\u0059I\u0045\u004c\u0044\u0044\u0049\u0053C")
	if _aad.Type == ResultTypeError {
		return _aad
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u0044\u0049S\u0043\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0070\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cgcd := args[2].ValueNumber
	if _cgcd <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0049E\u004c\u0044\u0044\u0049\u0053C\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0070\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020r\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_eeba := args[3].ValueNumber
	if _eeba <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "YI\u0045\u004cD\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076e\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072g\u0075m\u0065\u006et")
	}
	_gdbd := 0
	if _gfbf == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0059\u0049E\u004c\u0044\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_gdbd = int(args[4].ValueNumber)
		if !_gaff(_gdbd) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0059I\u0045\u004cD\u0044\u0049\u0053\u0043")
		}
	}
	_gaea, _aad := _agef(_agbg, _cgaf, _gdbd)
	if _aad.Type == ResultTypeError {
		return _aad
	}
	return MakeNumberResult((_eeba/_cgcd - 1) / _gaea)
}

var _ggaab = []ri{{1000, "\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {95, "\u0056\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

const (
	ErrorTypeValue ErrorType = iota
	ErrorTypeNull
	ErrorTypeRef
	ErrorTypeName
	ErrorTypeNum
	ErrorTypeSpill
	ErrorTypeNA
	ErrorTypeDivideByZero
)

// Coupncd implements the Excel COUPNCD function.
func Coupncd(args []Result) Result {
	_gbdg, _egg := _dccb(args, "\u0043O\u0055\u0050\u004e\u0043\u0044")
	if _egg.Type == ResultTypeError {
		return _egg
	}
	_ede := _dcge(_gbdg._eac)
	_fbdb := _dcge(_gbdg._aea)
	_fcgg := _gbdg._dfff
	_dfcc := _dccg(_ede, _fbdb, _fcgg)
	_bef, _gfd, _aaad := _dfcc.Date()
	return MakeNumberResult(_bad(_bef, int(_gfd), _aaad))
}
func _cecbd(_fceffb int) string {
	if _fceffb >= 1 && _fceffb-1 < len(_fdcdab) {
		if _fdcdab[_fceffb-1] != "" {
			return _fdcdab[_fceffb-1]
		}
	}
	return _d.Sprintf("\u0074\u006f\u006b\u002d\u0025\u0076", _fceffb)
}

// Reference returns an invalid reference for String.
func (_bbbg String) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

const _ebace = 57355

// Eval evaluates and returns the result of the cell reference.
func (_aef CellRef) Eval(ctx Context, ev Evaluator) Result { return ctx.Cell(_aef._ebd, ev) }

// Eval evaluates and returns the result of a function call.
func (_bacad FunctionCall) Eval(ctx Context, ev Evaluator) Result {
	_accaa := LookupFunction(_bacad._ffcfde)
	if _accaa != nil {
		_bedab := make([]Result, len(_bacad._gace))
		for _fgcd, _bebg := range _bacad._gace {
			_bedab[_fgcd] = _bebg.Eval(ctx, ev)
			_bedab[_fgcd].Ref = _bebg.Reference(ctx, ev)
		}
		if _, _gcaf := _ceaeg[_bacad._ffcfde]; !_gcaf {
			if _eecdb, _cebeg := _edebb(_bedab); _eecdb {
				return _cebeg
			}
		}
		return _accaa(_bedab)
	}
	_efdab := LookupFunctionComplex(_bacad._ffcfde)
	if _efdab != nil {
		_ddbaa := make([]Result, len(_bacad._gace))
		for _bffe, _efcgb := range _bacad._gace {
			_ddbaa[_bffe] = _efcgb.Eval(ctx, ev)
			_ddbaa[_bffe].Ref = _efcgb.Reference(ctx, ev)
		}
		if _, _cgbed := _ceaeg[_bacad._ffcfde]; !_cgbed {
			if _bcafd, _ecgeg := _edebb(_ddbaa); _bcafd {
				return _ecgeg
			}
		}
		return _efdab(ctx, ev, _ddbaa)
	}
	return MakeErrorResult("\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020" + _bacad._ffcfde)
}

// Reference returns a string reference value to a horizontal range.
func (_ecad HorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeHorizontalRange, Value: _ecad.horizontalRangeReference()}
}
func _fgc(_gfc int, _bfgc _ee.Month, _fcd int) int64 {
	if _gfc == 1900 && int(_bfgc) <= 2 {
		_fcd--
	}
	_fdg := _ee.Date(_gfc, _bfgc, _fcd, 0, 0, 0, 0, _ee.UTC)
	return _fdg.Unix()
}

// Upper is an implementation of the Excel UPPER function that returns a upper
// case version of a string.
func Upper(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0055\u0050\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ecbb := args[0].AsString()
	if _ecbb.Type != ResultTypeString {
		return MakeErrorResult("\u0055\u0050\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeStringResult(_dda.ToUpper(_ecbb.ValueString))
}
