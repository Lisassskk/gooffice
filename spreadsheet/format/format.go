/
// As an example, the same value of 1.2345 can be displayed as:
// - "1" with format "0"
// - "1.2" with format "0.0"
// - "1.23" with format "0.00"
// - "1.235" with format "0.000"
// - "123%" with format "0%"
// - "1 23/100" with fornat "0 0/100"
// - "1.23E+00" with format "0.00E+00"
// - "29:37:41s" with format `[h]:mm:ss"s"`
package format

import (
	_g "bytes"
	_b "fmt"
	_d "gitee.com/gooffice/gooffice"
	_c "io"
	_ea "math"
	_ce "strconv"
	_bb "strings"
	_cd "time"
)

func _eb(_cga float64, _bg Format, _ba bool) string {
	if _bg._ca {
		return NumberGeneric(_cga)
	}
	_eca := make([]byte, 0, 20)
	_fe := _ea.Signbit(_cga)
	_gef := _ea.Abs(_cga)
	_fef := int64(0)
	_fg := int64(0)
	if _bg.IsExponential {
		for _gef >= 10 {
			_fg++
			_gef /= 10
		}
		for _gef < 1 {
			_fg--
			_gef *= 10
		}
	} else if _bg._ga {
		_gef *= 100
	} else if _bg._bd {
		if _bg._ee == 0 {
			_fc := _ea.Pow(10, float64(_bg._bdb))
			_fd, _eec := 1.0, 1.0
			_ = _fd
			for _adb := 1.0; _adb < _fc; _adb++ {
				_, _gge := _ea.Modf(_gef * float64(_adb))
				if _gge < _eec {
					_eec = _gge
					_fd = _adb
					if _gge == 0 {
						break
					}
				}
			}
			_bg._ee = int64(_fd)
		}
		_fef = int64(_gef*float64(_bg._ee) + 0.5)
		if len(_bg.Whole) > 0 && _fef > _bg._ee {
			_fef = int64(_gef*float64(_bg._ee)) % _bg._ee
			_gef -= float64(_fef) / float64(_bg._ee)
		} else {
			_gef -= float64(_fef) / float64(_bg._ee)
			if _ea.Abs(_gef) < 1 {
				_cab := true
				for _, _eac := range _bg.Whole {
					if _eac.Type == FmtTypeDigitOpt {
						continue
					}
					if _eac.Type == FmtTypeLiteral && _eac.Literal == ' ' {
						continue
					}
					_cab = false
				}
				if _cab {
					_bg.Whole = nil
				}
			}
		}
	}
	_cb := 1
	for _, _gfd := range _bg.Fractional {
		if _gfd.Type == FmtTypeDigit || _gfd.Type == FmtTypeDigitOpt {
			_cb++
		}
	}
	_gef += 5 * _ea.Pow10(-_cb)
	_ac, _deb := _ea.Modf(_gef)
	_eca = append(_eca, _be(_ac, _cga, _bg)...)
	_eca = append(_eca, _bgf(_deb, _cga, _bg)...)
	_eca = append(_eca, _bce(_fg, _bg)...)
	if _bg._bd {
		_eca = _ce.AppendInt(_eca, _fef, 10)
		_eca = append(_eca, '/')
		_eca = _ce.AppendInt(_eca, _bg._ee, 10)
	}
	if !_ba && _fe {
		return "\u002d" + string(_eca)
	}
	return string(_eca)
}
func _be(_bfd, _eff float64, _gff Format) []byte {
	if len(_gff.Whole) == 0 {
		return nil
	}
	_fgd := _cd.Date(1899, 12, 30, 0, 0, 0, 0, _cd.UTC)
	_cbc := _fgd.Add(_cd.Duration(_eff * float64(24*_cd.Hour)))
	_cbc = _afag(_cbc)
	_ged := _ce.AppendFloat(nil, _bfd, 'f', -1, 64)
	_fdd := make([]byte, 0, len(_ged))
	_ada := 0
	_eab := 1
_cef:
	for _ebg := len(_gff.Whole) - 1; _ebg >= 0; _ebg-- {
		_af := len(_ged) - 1 - _ada
		_aaf := _gff.Whole[_ebg]
		switch _aaf.Type {
		case FmtTypeDigit:
			if _af >= 0 {
				_fdd = append(_fdd, _ged[_af])
				_ada++
				_eab = _ebg
			} else {
				_fdd = append(_fdd, '0')
			}
		case FmtTypeDigitOpt:
			if _af >= 0 {
				_fdd = append(_fdd, _ged[_af])
				_ada++
				_eab = _ebg
			} else {
				for _fddc := _ebg; _fddc >= 0; _fddc-- {
					_bc := _gff.Whole[_fddc]
					if _bc.Type == FmtTypeLiteral {
						_fdd = append(_fdd, _bc.Literal)
					}
				}
				break _cef
			}
		case FmtTypeDollar:
			for _ed := _ada; _ed < len(_ged); _ed++ {
				_fdd = append(_fdd, _ged[len(_ged)-1-_ed])
				_ada++
			}
			_fdd = append(_fdd, '$')
		case FmtTypeComma:
			if !_gff._a {
				_fdd = append(_fdd, ',')
			}
		case FmtTypeLiteral:
			_fdd = append(_fdd, _aaf.Literal)
		case FmtTypeDate:
			_fdd = append(_fdd, _ddf(_cgd(_cbc, _aaf.DateTime))...)
		case FmtTypeTime:
			_fdd = append(_fdd, _ddf(_egc(_cbc, _eff, _aaf.DateTime))...)
		default:
			_d.Log("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0077\u0068\u006f\u006c\u0065 \u0025\u0076", _aaf)
		}
	}
	_aafb := _ddf(_fdd)
	if _ada < len(_ged) && (_ada != 0 || _gff._ge) {
		_ddfe := len(_ged) - _ada
		_acb := make([]byte, len(_aafb)+_ddfe)
		copy(_acb, _aafb[0:_eab])
		copy(_acb[_eab:], _ged[0:])
		copy(_acb[_eab+_ddfe:], _aafb[_eab:])
		_aafb = _acb
	}
	if _gff._a {
		_ddb := _g.Buffer{}
		_cf := 0
		for _bff := len(_aafb) - 1; _bff >= 0; _bff-- {
			if !(_aafb[_bff] >= '0' && _aafb[_bff] <= '9') {
				_cf++
			} else {
				break
			}
		}
		for _daa := 0; _daa < len(_aafb); _daa++ {
			_gad := (len(_aafb) - _daa - _cf)
			if _gad%3 == 0 && _gad != 0 && _daa != 0 {
				_ddb.WriteByte(',')
			}
			_ddb.WriteByte(_aafb[_daa])
		}
		_aafb = _ddb.Bytes()
	}
	return _aafb
}

const _gadb int = 34

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric(v float64) string {
	if _ea.Abs(v) >= _eg || _ea.Abs(v) <= _dg && v != 0 {
		return _cda(v)
	}
	_cfd := make([]byte, 0, 15)
	_cfd = _ce.AppendFloat(_cfd, v, 'f', -1, 64)
	if len(_cfd) > 11 {
		_debe := _cfd[11] - '0'
		if _debe >= 5 && _debe <= 9 {
			_cfd[10]++
			_cfd = _cfd[0:11]
			_cfd = _abf(_cfd)
		}
		_cfd = _cfd[0:11]
	} else if len(_cfd) == 11 {
		if _cfd[len(_cfd)-1] == '9' {
			_cfd[len(_cfd)-1]++
			_cfd = _abf(_cfd)
		}
	}
	_cfd = _ggf(_cfd)
	return string(_cfd)
}
func _cda(_fad float64) string {
	_gd := _ce.FormatFloat(_fad, 'E', -1, 64)
	_dbc := _ce.FormatFloat(_fad, 'E', 5, 64)
	if len(_gd) < len(_dbc) {
		return _ce.FormatFloat(_fad, 'E', 2, 64)
	}
	return _dbc
}
func _bgf(_adc, _acbg float64, _dbe Format) []byte {
	if len(_dbe.Fractional) == 0 {
		return nil
	}
	_bbd := _ce.AppendFloat(nil, _adc, 'f', -1, 64)
	if len(_bbd) > 2 {
		_bbd = _bbd[2:]
	} else {
		_bbd = nil
	}
	_baa := make([]byte, 0, len(_bbd))
	_baa = append(_baa, '.')
	_efb := 0
_ggb:
	for _cag := 0; _cag < len(_dbe.Fractional); _cag++ {
		_gc := _cag
		_cec := _dbe.Fractional[_cag]
		switch _cec.Type {
		case FmtTypeDigit:
			if _gc < len(_bbd) {
				_baa = append(_baa, _bbd[_gc])
				_efb++
			} else {
				_baa = append(_baa, '0')
			}
		case FmtTypeDigitOpt:
			if _gc >= 0 {
				_baa = append(_baa, _bbd[_gc])
				_efb++
			} else {
				break _ggb
			}
		case FmtTypeLiteral:
			_baa = append(_baa, _cec.Literal)
		default:
			_d.Log("\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020f\u0072\u0061\u0063\u0074\u0069\u006f\u006ea\u006c\u0020\u0025\u0076", _cec)
		}
	}
	return _baa
}

// Token is a format token in the Excel format string.
type Token struct {
	Type     FmtType
	Literal  byte
	DateTime string
}

func _egc(_fb _cd.Time, _ccd float64, _aed string) []byte {
	_aegb := []byte{}
	_fge := 0
	for _adcc := 0; _adcc < len(_aed); _adcc++ {
		var _degc string
		if _aed[_adcc] == ':' {
			_degc = string(_aed[_fge:_adcc])
			_fge = _adcc + 1
		} else if _adcc == len(_aed)-1 {
			_degc = string(_aed[_fge : _adcc+1])
		} else {
			continue
		}
		switch _degc {
		case "\u0064":
			_aegb = _fb.AppendFormat(_aegb, "\u0032")
		case "\u0068":
			_aegb = _fb.AppendFormat(_aegb, "\u0033")
		case "\u0068\u0068":
			_aegb = _fb.AppendFormat(_aegb, "\u0031\u0035")
		case "\u006d":
			_aegb = _fb.AppendFormat(_aegb, "\u0034")
		case "\u006d\u006d":
			_aegb = _fb.AppendFormat(_aegb, "\u0030\u0034")
		case "\u0073":
			_aegb = _fb.Round(_cd.Second).AppendFormat(_aegb, "\u0035")
		case "\u0073\u002e\u0030":
			_aegb = _fb.Round(_cd.Second/10).AppendFormat(_aegb, "\u0035\u002e\u0030")
		case "\u0073\u002e\u0030\u0030":
			_aegb = _fb.Round(_cd.Second/100).AppendFormat(_aegb, "\u0035\u002e\u0030\u0030")
		case "\u0073\u002e\u00300\u0030":
			_aegb = _fb.Round(_cd.Second/1000).AppendFormat(_aegb, "\u0035\u002e\u00300\u0030")
		case "\u0073\u0073":
			_aegb = _fb.Round(_cd.Second).AppendFormat(_aegb, "\u0030\u0035")
		case "\u0073\u0073\u002e\u0030":
			_aegb = _fb.Round(_cd.Second/10).AppendFormat(_aegb, "\u0030\u0035\u002e\u0030")
		case "\u0073\u0073\u002e0\u0030":
			_aegb = _fb.Round(_cd.Second/100).AppendFormat(_aegb, "\u0030\u0035\u002e0\u0030")
		case "\u0073\u0073\u002e\u0030\u0030\u0030":
			_aegb = _fb.Round(_cd.Second/1000).AppendFormat(_aegb, "\u0030\u0035\u002e\u0030\u0030\u0030")
		case "\u0041\u004d\u002fP\u004d":
			_aegb = _fb.AppendFormat(_aegb, "\u0050\u004d")
		case "\u005b\u0068\u005d":
			_aegb = _ce.AppendInt(_aegb, int64(_ccd*24), 10)
		case "\u005b\u006d\u005d":
			_aegb = _ce.AppendInt(_aegb, int64(_ccd*24*60), 10)
		case "\u005b\u0073\u005d":
			_aegb = _ce.AppendInt(_aegb, int64(_ccd*24*60*60), 10)
		case "":
		default:
			_d.Log("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0069\u006d\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073", _degc)
		}
		if _aed[_adcc] == ':' {
			_aegb = append(_aegb, ':')
		}
	}
	return _aegb
}

const _aef int = -1

// String returns the string formatted according to the type.  In format strings
// this is the fourth item, where '@' is used as a placeholder for text.
func String(v string, f string) string {
	_gb := Parse(f)
	var _gab Format
	if len(_gb) == 1 {
		_gab = _gb[0]
	} else if len(_gb) == 4 {
		_gab = _gb[3]
	}
	_gf := false
	for _, _aa := range _gab.Whole {
		if _aa.Type == FmtTypeText {
			_gf = true
		}
	}
	if !_gf {
		return v
	}
	_cg := _g.Buffer{}
	for _, _ad := range _gab.Whole {
		switch _ad.Type {
		case FmtTypeLiteral:
			_cg.WriteByte(_ad.Literal)
		case FmtTypeText:
			_cg.WriteString(v)
		}
	}
	return _cg.String()
}

const _df = "\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u004c\u0069\u0074\u0065\u0072a\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0069\u0067\u0069\u0074\u0046\u006d\u0074\u0054y\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0043o\u006d\u006d\u0061\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0065\u0063\u0069\u006da\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065Pe\u0072\u0063e\u006e\u0074\u0046\u006d\u0074\u0054\u0079\u0070e\u0044\u006f\u006c\u006c\u0061\u0072\u0046\u006d\u0074Ty\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0054\u0068\u006f\u0075\u0073\u0061n\u0064\u0073\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0055n\u0064\u0065\u0072\u0073c\u006f\u0072\u0065\u0046\u006d\u0074T\u0079\u0070\u0065\u0044\u0061\u0074\u0065\u0046\u006d\u0074\u0054y\u0070e\u0054\u0069\u006d\u0065\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0046\u0072\u0061\u0063t\u0069\u006f\u006e\u0046\u006dt\u0054\u0079\u0070\u0065\u0054e\u0078\u0074"
const (
	FmtTypeLiteral FmtType = iota
	FmtTypeDigit
	FmtTypeDigitOpt
	FmtTypeComma
	FmtTypeDecimal
	FmtTypePercent
	FmtTypeDollar
	FmtTypeDigitOptThousands
	FmtTypeUnderscore
	FmtTypeDate
	FmtTypeTime
	FmtTypeFraction
	FmtTypeText
)

func _bce(_cbf int64, _ggg Format) []byte {
	if !_ggg.IsExponential || len(_ggg.Exponent) == 0 {
		return nil
	}
	_aeg := _ce.AppendInt(nil, _gag(_cbf), 10)
	_afe := make([]byte, 0, len(_aeg)+2)
	_afe = append(_afe, 'E')
	if _cbf >= 0 {
		_afe = append(_afe, '+')
	} else {
		_afe = append(_afe, '-')
		_cbf *= -1
	}
	_fgf := 0
_dc:
	for _ecc := len(_ggg.Exponent) - 1; _ecc >= 0; _ecc-- {
		_gbc := len(_aeg) - 1 - _fgf
		_cae := _ggg.Exponent[_ecc]
		switch _cae.Type {
		case FmtTypeDigit:
			if _gbc >= 0 {
				_afe = append(_afe, _aeg[_gbc])
				_fgf++
			} else {
				_afe = append(_afe, '0')
			}
		case FmtTypeDigitOpt:
			if _gbc >= 0 {
				_afe = append(_afe, _aeg[_gbc])
				_fgf++
			} else {
				for _eeg := _ecc; _eeg >= 0; _eeg-- {
					_dfg := _ggg.Exponent[_eeg]
					if _dfg.Type == FmtTypeLiteral {
						_afe = append(_afe, _dfg.Literal)
					}
				}
				break _dc
			}
		case FmtTypeLiteral:
			_afe = append(_afe, _cae.Literal)
		default:
			_d.Log("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0065\u0078p\u0020\u0025\u0076", _cae)
		}
	}
	if _fgf < len(_aeg) {
		_afe = append(_afe, _aeg[len(_aeg)-_fgf-1:_fgf-1]...)
	}
	_ddf(_afe[2:])
	return _afe
}

// FmtType is the type of a format token.
//go:generate stringer -type=FmtType
type FmtType byte

func _gag(_ebd int64) int64 {
	if _ebd < 0 {
		return -_ebd
	}
	return _ebd
}

const _agg int = 0
const _dg = 1e-10
const _dae int = 0

func _ggf(_bad []byte) []byte {
	_abd := len(_bad)
	_acc := false
	_gga := false
	for _acd := len(_bad) - 1; _acd >= 0; _acd-- {
		if _bad[_acd] == '0' && !_gga && !_acc {
			_abd = _acd
		} else if _bad[_acd] == '.' {
			_acc = true
		} else {
			_gga = true
		}
	}
	if _acc && _gga {
		if _bad[_abd-1] == '.' {
			_abd--
		}
		return _bad[0:_abd]
	}
	return _bad
}
func _abf(_ege []byte) []byte {
	for _gdb := len(_ege) - 1; _gdb > 0; _gdb-- {
		if _ege[_gdb] == '9'+1 {
			_ege[_gdb] = '0'
			if _ege[_gdb-1] == '.' {
				_gdb--
			}
			_ege[_gdb-1]++
		}
	}
	if _ege[0] == '9'+1 {
		_ege[0] = '0'
		copy(_ege[1:], _ege[0:])
		_ege[0] = '1'
	}
	return _ege
}
func IsNumber(data string) (_ecb bool) {
	_aab, _gbcf, _aac := 0, 0, len(data)
	_cfe := len(data)
	_debg, _cba, _cgdc := 0, 0, 0
	_ = _cba
	_ = _cgdc
	_ = _debg
	{
		_aab = _dae
		_debg = 0
		_cba = 0
		_cgdc = 0
	}
	{
		if _gbcf == _aac {
			goto _bea
		}
		switch _aab {
		case 0:
			goto _gdc
		case 1:
			goto _cecb
		case 2:
			goto _ebdca
		case 3:
			goto _bbg
		case 4:
			goto _aaa
		case 5:
			goto _aabg
		case 6:
			goto _edb
		case 7:
			goto _dgd
		}
		goto _ccdc
	_bdg:
		_cba = _gbcf
		_gbcf--
		{
			_ecb = false
		}
		goto _gcb
	_ade:
		_cba = _gbcf
		_gbcf--
		{
			_ecb = _cba == len(data)
		}
		goto _gcb
	_fbc:
		_cba = _gbcf
		_gbcf--
		{
			_ecb = _cba == len(data)
		}
		goto _gcb
	_cbg:
		switch _cgdc {
		case 2:
			{
				_gbcf = (_cba) - 1
				_ecb = _cba == len(data)
			}
		case 3:
			{
				_gbcf = (_cba) - 1
				_ecb = false
			}
		}
		goto _gcb
	_gcb:
		_debg = 0
		if _gbcf++; _gbcf == _aac {
			goto _bdf
		}
	_gdc:
		_debg = _gbcf
		switch data[_gbcf] {
		case 43:
			goto _gaa
		case 45:
			goto _gaa
		}
		if 48 <= data[_gbcf] && data[_gbcf] <= 57 {
			goto _fae
		}
		goto _acf
	_acf:
		if _gbcf++; _gbcf == _aac {
			goto _dff
		}
	_cecb:
		goto _acf
	_gaa:
		if _gbcf++; _gbcf == _aac {
			goto _add
		}
	_ebdca:
		if 48 <= data[_gbcf] && data[_gbcf] <= 57 {
			goto _fae
		}
		goto _acf
	_fae:
		if _gbcf++; _gbcf == _aac {
			goto _geb
		}
	_bbg:
		if data[_gbcf] == 46 {
			goto _gaf
		}
		if 48 <= data[_gbcf] && data[_gbcf] <= 57 {
			goto _fae
		}
		goto _acf
	_gaf:
		if _gbcf++; _gbcf == _aac {
			goto _dag
		}
	_aaa:
		if 48 <= data[_gbcf] && data[_gbcf] <= 57 {
			goto _fea
		}
		goto _acf
	_fea:
		if _gbcf++; _gbcf == _aac {
			goto _dgc
		}
	_aabg:
		if data[_gbcf] == 69 {
			goto _faf
		}
		if 48 <= data[_gbcf] && data[_gbcf] <= 57 {
			goto _fea
		}
		goto _acf
	_faf:
		if _gbcf++; _gbcf == _aac {
			goto _dage
		}
	_edb:
		switch data[_gbcf] {
		case 43:
			goto _eecd
		case 45:
			goto _eecd
		}
		goto _acf
	_eecd:
		_cba = _gbcf + 1
		_cgdc = 3
		goto _edc
	_fafc:
		_cba = _gbcf + 1
		_cgdc = 2
		goto _edc
	_edc:
		if _gbcf++; _gbcf == _aac {
			goto _dgf
		}
	_dgd:
		if 48 <= data[_gbcf] && data[_gbcf] <= 57 {
			goto _fafc
		}
		goto _acf
	_ccdc:
	_bdf:
		_aab = 0
		goto _bea
	_dff:
		_aab = 1
		goto _bea
	_add:
		_aab = 2
		goto _bea
	_geb:
		_aab = 3
		goto _bea
	_dag:
		_aab = 4
		goto _bea
	_dgc:
		_aab = 5
		goto _bea
	_dage:
		_aab = 6
		goto _bea
	_dgf:
		_aab = 7
		goto _bea
	_bea:
		{
		}
		if _gbcf == _cfe {
			switch _aab {
			case 1:
				goto _bdg
			case 2:
				goto _bdg
			case 3:
				goto _ade
			case 4:
				goto _bdg
			case 5:
				goto _fbc
			case 6:
				goto _bdg
			case 7:
				goto _cbg
			}
		}
	}
	if _aab == _aef {
		return false
	}
	return
}

const _eg = 1e11

func (_ef FmtType) String() string {
	if _ef >= FmtType(len(_bf)-1) {
		return _b.Sprintf("F\u006d\u0074\u0054\u0079\u0070\u0065\u0028\u0025\u0064\u0029", _ef)
	}
	return _df[_bf[_ef]:_bf[_ef+1]]
}
func Parse(s string) []Format {
	_aagd := Lexer{}
	_aagd.Lex(_bb.NewReader(s))
	_aagd._aff = append(_aagd._aff, _aagd._bca)
	return _aagd._aff
}

// Format is a parsed number format.
type Format struct {
	Whole         []Token
	Fractional    []Token
	Exponent      []Token
	IsExponential bool
	_bd           bool
	_ga           bool
	_ca           bool
	_a            bool
	_dd           bool
	_ge           bool
	_ee           int64
	_bdb          int
}

var _bf = [...]uint8{0, 14, 26, 41, 53, 67, 81, 94, 118, 135, 146, 157, 172, 183}

// Number is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number(v float64, f string) string {
	if f == "" || f == "\u0047e\u006e\u0065\u0072\u0061\u006c" || f == "\u0040" {
		return NumberGeneric(v)
	}
	_dgg := Parse(f)
	if len(_dgg) == 1 {
		return _eb(v, _dgg[0], false)
	} else if len(_dgg) > 1 && v < 0 {
		return _eb(v, _dgg[1], true)
	} else if len(_dgg) > 2 && v == 0 {
		return _eb(v, _dgg[2], false)
	}
	return _eb(v, _dgg[0], false)
}
func _afag(_dfdc _cd.Time) _cd.Time {
	_dfdc = _dfdc.UTC()
	return _cd.Date(_dfdc.Year(), _dfdc.Month(), _dfdc.Day(), _dfdc.Hour(), _dfdc.Minute(), _dfdc.Second(), _dfdc.Nanosecond(), _cd.Local)
}

const _ff int = 34
const _gfdf int = 34

func _cgd(_aag _cd.Time, _dfa string) []byte {
	_aega := []byte{}
	_cdb := 0
	for _dcf := 0; _dcf < len(_dfa); _dcf++ {
		var _aca string
		if _dfa[_dcf] == '/' {
			_aca = string(_dfa[_cdb:_dcf])
			_cdb = _dcf + 1
		} else if _dcf == len(_dfa)-1 {
			_aca = string(_dfa[_cdb : _dcf+1])
		} else {
			continue
		}
		switch _aca {
		case "\u0079\u0079":
			_aega = _aag.AppendFormat(_aega, "\u0030\u0036")
		case "\u0079\u0079\u0079\u0079":
			_aega = _aag.AppendFormat(_aega, "\u0032\u0030\u0030\u0036")
		case "\u006d":
			_aega = _aag.AppendFormat(_aega, "\u0031")
		case "\u006d\u006d":
			_aega = _aag.AppendFormat(_aega, "\u0030\u0031")
		case "\u006d\u006d\u006d":
			_aega = _aag.AppendFormat(_aega, "\u004a\u0061\u006e")
		case "\u006d\u006d\u006d\u006d":
			_aega = _aag.AppendFormat(_aega, "\u004aa\u006e\u0075\u0061\u0072\u0079")
		case "\u006d\u006d\u006dm\u006d":
			switch _aag.Month() {
			case _cd.January, _cd.July, _cd.June:
				_aega = append(_aega, 'J')
			case _cd.February:
				_aega = append(_aega, 'M')
			case _cd.March, _cd.May:
				_aega = append(_aega, 'M')
			case _cd.April, _cd.August:
				_aega = append(_aega, 'A')
			case _cd.September:
				_aega = append(_aega, 'S')
			case _cd.October:
				_aega = append(_aega, 'O')
			case _cd.November:
				_aega = append(_aega, 'N')
			case _cd.December:
				_aega = append(_aega, 'D')
			}
		case "\u0064":
			_aega = _aag.AppendFormat(_aega, "\u0032")
		case "\u0064\u0064":
			_aega = _aag.AppendFormat(_aega, "\u0030\u0032")
		case "\u0064\u0064\u0064":
			_aega = _aag.AppendFormat(_aega, "\u004d\u006f\u006e")
		case "\u0064\u0064\u0064\u0064":
			_aega = _aag.AppendFormat(_aega, "\u004d\u006f\u006e\u0064\u0061\u0079")
		default:
			_d.Log("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073", _aca)
		}
		if _dfa[_dcf] == '/' {
			_aega = append(_aega, '/')
		}
	}
	return _aega
}

const _aga int = 0

func (_cde *Lexer) nextFmt() { _cde._aff = append(_cde._aff, _cde._bca); _cde._bca = Format{} }
func (_cdf *Lexer) Lex(r _c.Reader) {
	_ddfee, _agc, _fcb := 0, 0, 0
	_dcfe := -1
	_eee, _cefa, _dggc := 0, 0, 0
	_ = _cefa
	_ = _dggc
	_gde := 1
	_ = _gde
	_ffd := make([]byte, 4096)
	_dega := false
	for !_dega {
		_bec := 0
		if _eee > 0 {
			_bec = _agc - _eee
		}
		_agc = 0
		_fec, _eae := r.Read(_ffd[_bec:])
		if _fec == 0 || _eae != nil {
			_dega = true
		}
		_fcb = _fec + _bec
		if _fcb < len(_ffd) {
			_dcfe = _fcb
		}
		{
			_ddfee = _gadb
			_eee = 0
			_cefa = 0
			_dggc = 0
		}
		{
			if _agc == _fcb {
				goto _aedc
			}
			switch _ddfee {
			case 34:
				goto _cgc
			case 35:
				goto _cbfe
			case 0:
				goto _bfdb
			case 36:
				goto _caf
			case 37:
				goto _cfeb
			case 1:
				goto _eecf
			case 2:
				goto _cgfg
			case 38:
				goto _ebeb
			case 3:
				goto _cbe
			case 4:
				goto _bdbe
			case 39:
				goto _edg
			case 5:
				goto _bcaf
			case 6:
				goto _bbf
			case 7:
				goto _cfdb
			case 8:
				goto _effd
			case 40:
				goto _cbad
			case 9:
				goto _ead
			case 41:
				goto _edcg
			case 10:
				goto _agafe
			case 42:
				goto _beg
			case 11:
				goto _bcc
			case 43:
				goto _edgc
			case 44:
				goto _daf
			case 45:
				goto _ceg
			case 12:
				goto _egg
			case 46:
				goto _eadc
			case 13:
				goto _eag
			case 14:
				goto _dfge
			case 15:
				goto _abg
			case 16:
				goto _fbg
			case 47:
				goto _gadd
			case 17:
				goto _degb
			case 48:
				goto _dac
			case 18:
				goto _egcd
			case 19:
				goto _cbfa
			case 20:
				goto _gagc
			case 49:
				goto _cfdd
			case 50:
				goto _ddd
			case 21:
				goto _abe
			case 22:
				goto _gggc
			case 23:
				goto _degab
			case 24:
				goto _dbg
			case 25:
				goto _gce
			case 51:
				goto _fcg
			case 26:
				goto _degg
			case 52:
				goto _afa
			case 53:
				goto _eadf
			case 54:
				goto _fcbd
			case 55:
				goto _bef
			case 56:
				goto _edea
			case 57:
				goto _cgfa
			case 27:
				goto _bfg
			case 28:
				goto _gbcg
			case 29:
				goto _abbd
			case 30:
				goto _gebd
			case 31:
				goto _becc
			case 58:
				goto _abc
			case 32:
				goto _bgc
			case 59:
				goto _dba
			case 33:
				goto _fabe
			case 60:
				goto _fgea
			case 61:
				goto _fbcg
			case 62:
				goto _abdd
			}
			goto _efg
		_cgf:
			switch _dggc {
			case 2:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeDigit, nil)
				}
			case 3:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeDigitOpt, nil)
				}
			case 5:
				{
					_agc = (_cefa) - 1
				}
			case 8:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypePercent, nil)
				}
			case 13:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeFraction, _ffd[_eee:_cefa])
				}
			case 14:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeDate, _ffd[_eee:_cefa])
				}
			case 15:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeTime, _ffd[_eee:_cefa])
				}
			case 16:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeTime, _ffd[_eee:_cefa])
				}
			case 18:
				{
					_agc = (_cefa) - 1
				}
			case 20:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeLiteral, _ffd[_eee:_cefa])
				}
			case 21:
				{
					_agc = (_cefa) - 1
					_cdf._bca.AddToken(FmtTypeLiteral, _ffd[_eee+1:_cefa-1])
				}
			}
			goto _ebe
		_dbb:
			_agc = (_cefa) - 1
			{
				_cdf._bca.AddToken(FmtTypeFraction, _ffd[_eee:_cefa])
			}
			goto _ebe
		_fbb:
			_agc = (_cefa) - 1
			{
				_cdf._bca.AddToken(FmtTypeDigitOpt, nil)
			}
			goto _ebe
		_dfd:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeDigitOptThousands, nil)
			}
			goto _ebe
		_fbe:
			_agc = (_cefa) - 1
			{
				_cdf._bca.AddToken(FmtTypePercent, nil)
			}
			goto _ebe
		_ede:
			_agc = (_cefa) - 1
			{
				_cdf._bca.AddToken(FmtTypeDate, _ffd[_eee:_cefa])
			}
			goto _ebe
		_bcec:
			_agc = (_cefa) - 1
			{
				_cdf._bca.AddToken(FmtTypeDigit, nil)
			}
			goto _ebe
		_baf:
			_agc = (_cefa) - 1
			{
				_cdf._bca.AddToken(FmtTypeTime, _ffd[_eee:_cefa])
			}
			goto _ebe
		_aacc:
			_agc = (_cefa) - 1
			{
				_cdf._bca.AddToken(FmtTypeLiteral, _ffd[_eee:_cefa])
			}
			goto _ebe
		_ccc:
			_cefa = _agc + 1
			{
				_cdf._bca._ca = true
			}
			goto _ebe
		_geg:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeLiteral, _ffd[_eee:_cefa])
			}
			goto _ebe
		_ggec:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeDollar, nil)
			}
			goto _ebe
		_ffc:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeComma, nil)
			}
			goto _ebe
		_caea:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeDecimal, nil)
			}
			goto _ebe
		_edbe:
			_cefa = _agc + 1
			{
				_cdf.nextFmt()
			}
			goto _ebe
		_acda:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeText, nil)
			}
			goto _ebe
		_gcc:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeUnderscore, nil)
			}
			goto _ebe
		_fac:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypeLiteral, _ffd[_eee:_cefa])
			}
			goto _ebe
		_eda:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypeLiteral, _ffd[_eee+1:_cefa-1])
			}
			goto _ebe
		_fbcf:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypeDigitOpt, nil)
			}
			goto _ebe
		_bdd:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypeFraction, _ffd[_eee:_cefa])
			}
			goto _ebe
		_cfb:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypePercent, nil)
			}
			goto _ebe
		_acdd:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypeDate, _ffd[_eee:_cefa])
			}
			goto _ebe
		_adbc:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypeDigit, nil)
			}
			goto _ebe
		_gedg:
			_cefa = _agc
			_agc--
			{
				_cdf._bca.AddToken(FmtTypeTime, _ffd[_eee:_cefa])
			}
			goto _ebe
		_gae:
			_cefa = _agc
			_agc--
			{
			}
			goto _ebe
		_ebb:
			_cefa = _agc + 1
			{
				_cdf._bca.IsExponential = true
			}
			goto _ebe
		_bfda:
			_cefa = _agc + 1
			{
				_cdf._bca.AddToken(FmtTypeLiteral, _ffd[_eee+1:_cefa])
			}
			goto _ebe
		_ebe:
			_eee = 0
			if _agc++; _agc == _fcb {
				goto _efc
			}
		_cgc:
			_eee = _agc
			switch _ffd[_agc] {
			case 34:
				goto _eecb
			case 35:
				goto _gbg
			case 36:
				goto _ggec
			case 37:
				goto _cff
			case 44:
				goto _ffc
			case 46:
				goto _caea
			case 47:
				goto _adae
			case 48:
				goto _cbd
			case 58:
				goto _ddg
			case 59:
				goto _edbe
			case 63:
				goto _dfged
			case 64:
				goto _acda
			case 65:
				goto _bbb
			case 69:
				goto _baae
			case 71:
				goto _efe
			case 91:
				goto _dbcg
			case 92:
				goto _cbaf
			case 95:
				goto _gcc
			case 100:
				goto _adae
			case 104:
				goto _ddg
			case 109:
				goto _fbcc
			case 115:
				goto _cfbe
			case 121:
				goto _degcg
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _adaf
			}
			goto _geg
		_eecb:
			_cefa = _agc + 1
			_dggc = 20
			goto _ccb
		_ccb:
			if _agc++; _agc == _fcb {
				goto _efdf
			}
		_cbfe:
			if _ffd[_agc] == 34 {
				goto _ebdb
			}
			goto _cbfec
		_cbfec:
			if _agc++; _agc == _fcb {
				goto _degce
			}
		_bfdb:
			if _ffd[_agc] == 34 {
				goto _ebdb
			}
			goto _cbfec
		_ebdb:
			_cefa = _agc + 1
			_dggc = 21
			goto _gbd
		_gbd:
			if _agc++; _agc == _fcb {
				goto _fdgd
			}
		_caf:
			if _ffd[_agc] == 34 {
				goto _cbfec
			}
			goto _eda
		_gbg:
			_cefa = _agc + 1
			_dggc = 3
			goto _eed
		_eed:
			if _agc++; _agc == _fcb {
				goto _fcd
			}
		_cfeb:
			switch _ffd[_agc] {
			case 35:
				goto _ggaf
			case 37:
				goto _ggaf
			case 44:
				goto _ggc
			case 47:
				goto _agaf
			case 48:
				goto _ggaf
			case 63:
				goto _ggaf
			}
			goto _fbcf
		_ggaf:
			if _agc++; _agc == _fcb {
				goto _beba
			}
		_eecf:
			switch _ffd[_agc] {
			case 35:
				goto _ggaf
			case 37:
				goto _ggaf
			case 47:
				goto _agaf
			case 48:
				goto _ggaf
			case 63:
				goto _ggaf
			}
			goto _cgf
		_agaf:
			if _agc++; _agc == _fcb {
				goto _bcdg
			}
		_cgfg:
			switch _ffd[_agc] {
			case 35:
				goto _bcd
			case 37:
				goto _efd
			case 48:
				goto _bcda
			case 63:
				goto _bcd
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cabf
			}
			goto _cgf
		_bcd:
			_cefa = _agc + 1
			goto _agf
		_agf:
			if _agc++; _agc == _fcb {
				goto _edbb
			}
		_ebeb:
			switch _ffd[_agc] {
			case 35:
				goto _bcd
			case 37:
				goto _bcd
			case 44:
				goto _bcd
			case 46:
				goto _bcd
			case 48:
				goto _bcd
			case 63:
				goto _bcd
			case 65:
				goto _bba
			}
			goto _bdd
		_bba:
			if _agc++; _agc == _fcb {
				goto _fbf
			}
		_cbe:
			switch _ffd[_agc] {
			case 47:
				goto _gbgf
			case 77:
				goto _dde
			}
			goto _dbb
		_gbgf:
			if _agc++; _agc == _fcb {
				goto _eaf
			}
		_bdbe:
			if _ffd[_agc] == 80 {
				goto _aaaf
			}
			goto _dbb
		_aaaf:
			_cefa = _agc + 1
			goto _daeb
		_daeb:
			if _agc++; _agc == _fcb {
				goto _fgde
			}
		_edg:
			if _ffd[_agc] == 65 {
				goto _bba
			}
			goto _bdd
		_dde:
			if _agc++; _agc == _fcb {
				goto _debeac
			}
		_bcaf:
			if _ffd[_agc] == 47 {
				goto _gdcd
			}
			goto _dbb
		_gdcd:
			if _agc++; _agc == _fcb {
				goto _cgab
			}
		_bbf:
			if _ffd[_agc] == 80 {
				goto _gage
			}
			goto _dbb
		_gage:
			if _agc++; _agc == _fcb {
				goto _egeba
			}
		_cfdb:
			if _ffd[_agc] == 77 {
				goto _aaaf
			}
			goto _dbb
		_efd:
			if _agc++; _agc == _fcb {
				goto _ccf
			}
		_effd:
			switch _ffd[_agc] {
			case 35:
				goto _ebga
			case 37:
				goto _egeb
			case 63:
				goto _ebga
			}
			if 48 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _dgfg
			}
			goto _cgf
		_ebga:
			_cefa = _agc + 1
			goto _afc
		_afc:
			if _agc++; _agc == _fcb {
				goto _ddbff
			}
		_cbad:
			switch _ffd[_agc] {
			case 35:
				goto _bcd
			case 37:
				goto _fgg
			case 44:
				goto _bcd
			case 46:
				goto _bcd
			case 48:
				goto _bcd
			case 63:
				goto _bcd
			case 65:
				goto _bba
			}
			goto _bdd
		_fgg:
			if _agc++; _agc == _fcb {
				goto _gee
			}
		_ead:
			switch _ffd[_agc] {
			case 35:
				goto _dfe
			case 44:
				goto _dfe
			case 46:
				goto _dfe
			case 48:
				goto _dfe
			case 63:
				goto _dfe
			}
			goto _dbb
		_dfe:
			_cefa = _agc + 1
			goto _dfef
		_dfef:
			if _agc++; _agc == _fcb {
				goto _edegc
			}
		_edcg:
			switch _ffd[_agc] {
			case 35:
				goto _dfe
			case 44:
				goto _dfe
			case 46:
				goto _dfe
			case 48:
				goto _dfe
			case 63:
				goto _dfe
			case 65:
				goto _bba
			}
			goto _bdd
		_egeb:
			if _agc++; _agc == _fcb {
				goto _efgd
			}
		_agafe:
			if _ffd[_agc] == 37 {
				goto _egeb
			}
			if 48 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _dgfg
			}
			goto _cgf
		_dgfg:
			_cefa = _agc + 1
			_dggc = 13
			goto _cad
		_cad:
			if _agc++; _agc == _fcb {
				goto _fcdd
			}
		_beg:
			switch _ffd[_agc] {
			case 35:
				goto _bcd
			case 37:
				goto _ecbe
			case 44:
				goto _bcd
			case 46:
				goto _bcd
			case 48:
				goto _ddbf
			case 63:
				goto _bcd
			case 65:
				goto _bba
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _dgfg
			}
			goto _bdd
		_ecbe:
			if _agc++; _agc == _fcb {
				goto _daed
			}
		_bcc:
			switch _ffd[_agc] {
			case 35:
				goto _dfe
			case 37:
				goto _egeb
			case 44:
				goto _dfe
			case 46:
				goto _dfe
			case 63:
				goto _dfe
			}
			if 48 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _dgfg
			}
			goto _dbb
		_ddbf:
			_cefa = _agc + 1
			goto _bcf
		_bcf:
			if _agc++; _agc == _fcb {
				goto _bcg
			}
		_edgc:
			switch _ffd[_agc] {
			case 35:
				goto _bcd
			case 37:
				goto _ddbf
			case 44:
				goto _bcd
			case 46:
				goto _bcd
			case 48:
				goto _ddbf
			case 63:
				goto _bcd
			case 65:
				goto _bba
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _dgfg
			}
			goto _bdd
		_bcda:
			_cefa = _agc + 1
			goto _edeg
		_edeg:
			if _agc++; _agc == _fcb {
				goto _bde
			}
		_daf:
			switch _ffd[_agc] {
			case 35:
				goto _bcd
			case 37:
				goto _ddbf
			case 44:
				goto _bcd
			case 46:
				goto _bcd
			case 48:
				goto _bcda
			case 63:
				goto _bcd
			case 65:
				goto _bba
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cabf
			}
			goto _bdd
		_cabf:
			_cefa = _agc + 1
			goto _bbgb
		_bbgb:
			if _agc++; _agc == _fcb {
				goto _fbd
			}
		_ceg:
			switch _ffd[_agc] {
			case 35:
				goto _bcd
			case 37:
				goto _dgfg
			case 44:
				goto _bcd
			case 46:
				goto _bcd
			case 48:
				goto _bcda
			case 63:
				goto _bcd
			case 65:
				goto _bba
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cabf
			}
			goto _bdd
		_ggc:
			if _agc++; _agc == _fcb {
				goto _aefe
			}
		_egg:
			if _ffd[_agc] == 35 {
				goto _dfd
			}
			goto _fbb
		_cff:
			_cefa = _agc + 1
			_dggc = 8
			goto _debea
		_debea:
			if _agc++; _agc == _fcb {
				goto _beae
			}
		_eadc:
			switch _ffd[_agc] {
			case 35:
				goto _gdea
			case 37:
				goto _bccc
			case 48:
				goto _dga
			case 63:
				goto _gdea
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cgcd
			}
			goto _cfb
		_gdea:
			if _agc++; _agc == _fcb {
				goto _geba
			}
		_eag:
			switch _ffd[_agc] {
			case 35:
				goto _gdea
			case 47:
				goto _agaf
			case 48:
				goto _gdea
			case 63:
				goto _gdea
			}
			goto _fbe
		_bccc:
			if _agc++; _agc == _fcb {
				goto _aagg
			}
		_dfge:
			if _ffd[_agc] == 37 {
				goto _bccc
			}
			if 48 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cgcd
			}
			goto _cgf
		_cgcd:
			if _agc++; _agc == _fcb {
				goto _adf
			}
		_abg:
			switch _ffd[_agc] {
			case 37:
				goto _bccc
			case 47:
				goto _agaf
			}
			if 48 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cgcd
			}
			goto _cgf
		_dga:
			if _agc++; _agc == _fcb {
				goto _fbdc
			}
		_fbg:
			switch _ffd[_agc] {
			case 35:
				goto _gdea
			case 37:
				goto _bccc
			case 47:
				goto _agaf
			case 48:
				goto _dga
			case 63:
				goto _gdea
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cgcd
			}
			goto _fbe
		_adae:
			_cefa = _agc + 1
			goto _beb
		_beb:
			if _agc++; _agc == _fcb {
				goto _gac
			}
		_gadd:
			switch _ffd[_agc] {
			case 47:
				goto _adae
			case 100:
				goto _adae
			case 109:
				goto _adae
			case 121:
				goto _bgd
			}
			goto _acdd
		_bgd:
			if _agc++; _agc == _fcb {
				goto _gcd
			}
		_degb:
			if _ffd[_agc] == 121 {
				goto _adae
			}
			goto _ede
		_cbd:
			_cefa = _agc + 1
			_dggc = 2
			goto _fadg
		_fadg:
			if _agc++; _agc == _fcb {
				goto _cdae
			}
		_dac:
			switch _ffd[_agc] {
			case 35:
				goto _ggaf
			case 37:
				goto _fafcc
			case 47:
				goto _agaf
			case 48:
				goto _bga
			case 63:
				goto _ggaf
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cdfc
			}
			goto _adbc
		_fafcc:
			if _agc++; _agc == _fcb {
				goto _faca
			}
		_egcd:
			switch _ffd[_agc] {
			case 35:
				goto _ggaf
			case 37:
				goto _fafcc
			case 47:
				goto _agaf
			case 48:
				goto _fafcc
			case 63:
				goto _ggaf
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cgcd
			}
			goto _bcec
		_bga:
			if _agc++; _agc == _fcb {
				goto _age
			}
		_cbfa:
			switch _ffd[_agc] {
			case 35:
				goto _ggaf
			case 37:
				goto _fafcc
			case 47:
				goto _agaf
			case 48:
				goto _bga
			case 63:
				goto _ggaf
			}
			if 49 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cdfc
			}
			goto _bcec
		_cdfc:
			if _agc++; _agc == _fcb {
				goto _aggd
			}
		_gagc:
			switch _ffd[_agc] {
			case 37:
				goto _cgcd
			case 47:
				goto _agaf
			}
			if 48 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cdfc
			}
			goto _cgf
		_adaf:
			_cefa = _agc + 1
			_dggc = 20
			goto _adg
		_adg:
			if _agc++; _agc == _fcb {
				goto _bgde
			}
		_cfdd:
			switch _ffd[_agc] {
			case 37:
				goto _cgcd
			case 47:
				goto _agaf
			}
			if 48 <= _ffd[_agc] && _ffd[_agc] <= 57 {
				goto _cdfc
			}
			goto _fac
		_ddg:
			_cefa = _agc + 1
			_dggc = 15
			goto _gcg
		_gcg:
			if _agc++; _agc == _fcb {
				goto _bgcg
			}
		_ddd:
			switch _ffd[_agc] {
			case 58:
				goto _ddg
			case 65:
				goto _ccbf
			case 104:
				goto _ddg
			case 109:
				goto _ddg
			case 115:
				goto _cfbe
			}
			goto _gedg
		_ccbf:
			if _agc++; _agc == _fcb {
				goto _caa
			}
		_abe:
			switch _ffd[_agc] {
			case 47:
				goto _bdbg
			case 77:
				goto _faef
			}
			goto _cgf
		_bdbg:
			if _agc++; _agc == _fcb {
				goto _gbcfg
			}
		_gggc:
			if _ffd[_agc] == 80 {
				goto _ddg
			}
			goto _cgf
		_faef:
			if _agc++; _agc == _fcb {
				goto _cdfcd
			}
		_degab:
			if _ffd[_agc] == 47 {
				goto _bbff
			}
			goto _cgf
		_bbff:
			if _agc++; _agc == _fcb {
				goto _dda
			}
		_dbg:
			if _ffd[_agc] == 80 {
				goto _cdfd
			}
			goto _cgf
		_cdfd:
			if _agc++; _agc == _fcb {
				goto _dcg
			}
		_gce:
			if _ffd[_agc] == 77 {
				goto _ddg
			}
			goto _cgf
		_cfbe:
			_cefa = _agc + 1
			_dggc = 15
			goto _abfb
		_abfb:
			if _agc++; _agc == _fcb {
				goto _beab
			}
		_fcg:
			switch _ffd[_agc] {
			case 46:
				goto _fafcce
			case 58:
				goto _ddg
			case 65:
				goto _ccbf
			case 104:
				goto _ddg
			case 109:
				goto _ddg
			case 115:
				goto _cfbe
			}
			goto _gedg
		_fafcce:
			if _agc++; _agc == _fcb {
				goto _gbdg
			}
		_degg:
			if _ffd[_agc] == 48 {
				goto _fdg
			}
			goto _baf
		_fdg:
			_cefa = _agc + 1
			_dggc = 15
			goto _abb
		_abb:
			if _agc++; _agc == _fcb {
				goto _efda
			}
		_afa:
			switch _ffd[_agc] {
			case 48:
				goto _fbcfd
			case 58:
				goto _ddg
			case 65:
				goto _ccbf
			case 104:
				goto _ddg
			case 109:
				goto _ddg
			case 115:
				goto _cfbe
			}
			goto _gedg
		_fbcfd:
			_cefa = _agc + 1
			_dggc = 15
			goto _cfde
		_cfde:
			if _agc++; _agc == _fcb {
				goto _cffa
			}
		_eadf:
			switch _ffd[_agc] {
			case 48:
				goto _ddg
			case 58:
				goto _ddg
			case 65:
				goto _ccbf
			case 104:
				goto _ddg
			case 109:
				goto _ddg
			case 115:
				goto _cfbe
			}
			goto _gedg
		_dfged:
			_cefa = _agc + 1
			_dggc = 5
			goto _fab
		_fab:
			if _agc++; _agc == _fcb {
				goto _aee
			}
		_fcbd:
			switch _ffd[_agc] {
			case 35:
				goto _ggaf
			case 37:
				goto _ggaf
			case 47:
				goto _agaf
			case 48:
				goto _ggaf
			case 63:
				goto _ggaf
			}
			goto _gae
		_bbb:
			_cefa = _agc + 1
			_dggc = 20
			goto _ebeg
		_ebeg:
			if _agc++; _agc == _fcb {
				goto _eade
			}
		_bef:
			switch _ffd[_agc] {
			case 47:
				goto _bdbg
			case 77:
				goto _faef
			}
			goto _fac
		_baae:
			if _agc++; _agc == _fcb {
				goto _dfb
			}
		_edea:
			switch _ffd[_agc] {
			case 43:
				goto _ebb
			case 45:
				goto _ebb
			}
			goto _fac
		_efe:
			_cefa = _agc + 1
			goto _fdgf
		_fdgf:
			if _agc++; _agc == _fcb {
				goto _cgcc
			}
		_cgfa:
			if _ffd[_agc] == 101 {
				goto _gdf
			}
			goto _fac
		_gdf:
			if _agc++; _agc == _fcb {
				goto _ggbc
			}
		_bfg:
			if _ffd[_agc] == 110 {
				goto _ebf
			}
			goto _aacc
		_ebf:
			if _agc++; _agc == _fcb {
				goto _ddgb
			}
		_gbcg:
			if _ffd[_agc] == 101 {
				goto _fgfe
			}
			goto _aacc
		_fgfe:
			if _agc++; _agc == _fcb {
				goto _gdd
			}
		_abbd:
			if _ffd[_agc] == 114 {
				goto _acce
			}
			goto _aacc
		_acce:
			if _agc++; _agc == _fcb {
				goto _gbcfd
			}
		_gebd:
			if _ffd[_agc] == 97 {
				goto _ddgg
			}
			goto _aacc
		_ddgg:
			if _agc++; _agc == _fcb {
				goto _gfb
			}
		_becc:
			if _ffd[_agc] == 108 {
				goto _ccc
			}
			goto _aacc
		_dbcg:
			_cefa = _agc + 1
			_dggc = 20
			goto _aec
		_aec:
			if _agc++; _agc == _fcb {
				goto _gbf
			}
		_abc:
			switch _ffd[_agc] {
			case 104:
				goto _fbba
			case 109:
				goto _fbba
			case 115:
				goto _fbba
			}
			goto _bfbd
		_bfbd:
			if _agc++; _agc == _fcb {
				goto _dbaa
			}
		_bgc:
			if _ffd[_agc] == 93 {
				goto _dge
			}
			goto _bfbd
		_dge:
			_cefa = _agc + 1
			_dggc = 18
			goto _befa
		_ggcg:
			_cefa = _agc + 1
			_dggc = 16
			goto _befa
		_befa:
			if _agc++; _agc == _fcb {
				goto _cecf
			}
		_dba:
			if _ffd[_agc] == 93 {
				goto _dge
			}
			goto _bfbd
		_fbba:
			if _agc++; _agc == _fcb {
				goto _ebea
			}
		_fabe:
			if _ffd[_agc] == 93 {
				goto _ggcg
			}
			goto _bfbd
		_cbaf:
			if _agc++; _agc == _fcb {
				goto _gffc
			}
		_fgea:
			goto _bfda
		_fbcc:
			_cefa = _agc + 1
			_dggc = 14
			goto _febe
		_febe:
			if _agc++; _agc == _fcb {
				goto _ebdba
			}
		_fbcg:
			switch _ffd[_agc] {
			case 47:
				goto _adae
			case 58:
				goto _ddg
			case 65:
				goto _ccbf
			case 100:
				goto _adae
			case 104:
				goto _ddg
			case 109:
				goto _fbcc
			case 115:
				goto _cfbe
			case 121:
				goto _bgd
			}
			goto _acdd
		_degcg:
			if _agc++; _agc == _fcb {
				goto _abec
			}
		_abdd:
			if _ffd[_agc] == 121 {
				goto _adae
			}
			goto _fac
		_efg:
		_efc:
			_ddfee = 34
			goto _aedc
		_efdf:
			_ddfee = 35
			goto _aedc
		_degce:
			_ddfee = 0
			goto _aedc
		_fdgd:
			_ddfee = 36
			goto _aedc
		_fcd:
			_ddfee = 37
			goto _aedc
		_beba:
			_ddfee = 1
			goto _aedc
		_bcdg:
			_ddfee = 2
			goto _aedc
		_edbb:
			_ddfee = 38
			goto _aedc
		_fbf:
			_ddfee = 3
			goto _aedc
		_eaf:
			_ddfee = 4
			goto _aedc
		_fgde:
			_ddfee = 39
			goto _aedc
		_debeac:
			_ddfee = 5
			goto _aedc
		_cgab:
			_ddfee = 6
			goto _aedc
		_egeba:
			_ddfee = 7
			goto _aedc
		_ccf:
			_ddfee = 8
			goto _aedc
		_ddbff:
			_ddfee = 40
			goto _aedc
		_gee:
			_ddfee = 9
			goto _aedc
		_edegc:
			_ddfee = 41
			goto _aedc
		_efgd:
			_ddfee = 10
			goto _aedc
		_fcdd:
			_ddfee = 42
			goto _aedc
		_daed:
			_ddfee = 11
			goto _aedc
		_bcg:
			_ddfee = 43
			goto _aedc
		_bde:
			_ddfee = 44
			goto _aedc
		_fbd:
			_ddfee = 45
			goto _aedc
		_aefe:
			_ddfee = 12
			goto _aedc
		_beae:
			_ddfee = 46
			goto _aedc
		_geba:
			_ddfee = 13
			goto _aedc
		_aagg:
			_ddfee = 14
			goto _aedc
		_adf:
			_ddfee = 15
			goto _aedc
		_fbdc:
			_ddfee = 16
			goto _aedc
		_gac:
			_ddfee = 47
			goto _aedc
		_gcd:
			_ddfee = 17
			goto _aedc
		_cdae:
			_ddfee = 48
			goto _aedc
		_faca:
			_ddfee = 18
			goto _aedc
		_age:
			_ddfee = 19
			goto _aedc
		_aggd:
			_ddfee = 20
			goto _aedc
		_bgde:
			_ddfee = 49
			goto _aedc
		_bgcg:
			_ddfee = 50
			goto _aedc
		_caa:
			_ddfee = 21
			goto _aedc
		_gbcfg:
			_ddfee = 22
			goto _aedc
		_cdfcd:
			_ddfee = 23
			goto _aedc
		_dda:
			_ddfee = 24
			goto _aedc
		_dcg:
			_ddfee = 25
			goto _aedc
		_beab:
			_ddfee = 51
			goto _aedc
		_gbdg:
			_ddfee = 26
			goto _aedc
		_efda:
			_ddfee = 52
			goto _aedc
		_cffa:
			_ddfee = 53
			goto _aedc
		_aee:
			_ddfee = 54
			goto _aedc
		_eade:
			_ddfee = 55
			goto _aedc
		_dfb:
			_ddfee = 56
			goto _aedc
		_cgcc:
			_ddfee = 57
			goto _aedc
		_ggbc:
			_ddfee = 27
			goto _aedc
		_ddgb:
			_ddfee = 28
			goto _aedc
		_gdd:
			_ddfee = 29
			goto _aedc
		_gbcfd:
			_ddfee = 30
			goto _aedc
		_gfb:
			_ddfee = 31
			goto _aedc
		_gbf:
			_ddfee = 58
			goto _aedc
		_dbaa:
			_ddfee = 32
			goto _aedc
		_cecf:
			_ddfee = 59
			goto _aedc
		_ebea:
			_ddfee = 33
			goto _aedc
		_gffc:
			_ddfee = 60
			goto _aedc
		_ebdba:
			_ddfee = 61
			goto _aedc
		_abec:
			_ddfee = 62
			goto _aedc
		_aedc:
			{
			}
			if _agc == _dcfe {
				switch _ddfee {
				case 35:
					goto _fac
				case 0:
					goto _cgf
				case 36:
					goto _eda
				case 37:
					goto _fbcf
				case 1:
					goto _cgf
				case 2:
					goto _cgf
				case 38:
					goto _bdd
				case 3:
					goto _dbb
				case 4:
					goto _dbb
				case 39:
					goto _bdd
				case 5:
					goto _dbb
				case 6:
					goto _dbb
				case 7:
					goto _dbb
				case 8:
					goto _cgf
				case 40:
					goto _bdd
				case 9:
					goto _dbb
				case 41:
					goto _bdd
				case 10:
					goto _cgf
				case 42:
					goto _bdd
				case 11:
					goto _dbb
				case 43:
					goto _bdd
				case 44:
					goto _bdd
				case 45:
					goto _bdd
				case 12:
					goto _fbb
				case 46:
					goto _cfb
				case 13:
					goto _fbe
				case 14:
					goto _cgf
				case 15:
					goto _cgf
				case 16:
					goto _fbe
				case 47:
					goto _acdd
				case 17:
					goto _ede
				case 48:
					goto _adbc
				case 18:
					goto _bcec
				case 19:
					goto _bcec
				case 20:
					goto _cgf
				case 49:
					goto _fac
				case 50:
					goto _gedg
				case 21:
					goto _cgf
				case 22:
					goto _cgf
				case 23:
					goto _cgf
				case 24:
					goto _cgf
				case 25:
					goto _cgf
				case 51:
					goto _gedg
				case 26:
					goto _baf
				case 52:
					goto _gedg
				case 53:
					goto _gedg
				case 54:
					goto _gae
				case 55:
					goto _fac
				case 56:
					goto _fac
				case 57:
					goto _fac
				case 27:
					goto _aacc
				case 28:
					goto _aacc
				case 29:
					goto _aacc
				case 30:
					goto _aacc
				case 31:
					goto _aacc
				case 58:
					goto _fac
				case 32:
					goto _cgf
				case 59:
					goto _cgf
				case 33:
					goto _aacc
				case 60:
					goto _fac
				case 61:
					goto _acdd
				case 62:
					goto _fac
				}
			}
		}
		if _eee > 0 {
			copy(_ffd[0:], _ffd[_eee:])
		}
	}
	_ = _dcfe
	if _ddfee == _aef {
		_d.Log("\u0066o\u0072m\u0061\u0074\u0020\u0070\u0061r\u0073\u0065 \u0065\u0072\u0072\u006f\u0072")
	}
}

// Value formats a value as a number or string depending on  if it appears to be
// a number or string.
func Value(v string, f string) string {
	if IsNumber(v) {
		_eef, _ := _ce.ParseFloat(v, 64)
		return Number(_eef, f)
	}
	return String(v, f)
}

type Lexer struct {
	_bca Format
	_aff []Format
}

// AddToken adds a format token to the format.
func (_bfb *Format) AddToken(t FmtType, l []byte) {
	if _bfb._dd {
		_bfb._dd = false
		return
	}
	switch t {
	case FmtTypeDecimal:
		_bfb._ge = true
	case FmtTypeUnderscore:
		_bfb._dd = true
	case FmtTypeText:
		_bfb.Whole = append(_bfb.Whole, Token{Type: t})
	case FmtTypeDate, FmtTypeTime:
		_bfb.Whole = append(_bfb.Whole, Token{Type: t, DateTime: string(l)})
	case FmtTypePercent:
		_bfb._ga = true
		t = FmtTypeLiteral
		l = []byte{'%'}
		fallthrough
	case FmtTypeDigitOpt:
		fallthrough
	case FmtTypeLiteral, FmtTypeDigit, FmtTypeDollar, FmtTypeComma:
		if l == nil {
			l = []byte{0}
		}
		for _, _ae := range l {
			if _bfb.IsExponential {
				_bfb.Exponent = append(_bfb.Exponent, Token{Type: t, Literal: _ae})
			} else if !_bfb._ge {
				_bfb.Whole = append(_bfb.Whole, Token{Type: t, Literal: _ae})
			} else {
				_bfb.Fractional = append(_bfb.Fractional, Token{Type: t, Literal: _ae})
			}
		}
	case FmtTypeDigitOptThousands:
		_bfb._a = true
	case FmtTypeFraction:
		_cc := _bb.Split(string(l), "\u002f")
		if len(_cc) == 2 {
			_bfb._bd = true
			_bfb._ee, _ = _ce.ParseInt(_cc[1], 10, 64)
			for _, _ag := range _cc[1] {
				if _ag == '?' || _ag == '0' {
					_bfb._bdb++
				}
			}
		}
	default:
		_d.Log("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0070\u0068\u0020\u0074y\u0070\u0065\u0020\u0069\u006e\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0076", t)
	}
}
func _ddf(_bfc []byte) []byte {
	for _fa := 0; _fa < len(_bfc)/2; _fa++ {
		_db := len(_bfc) - 1 - _fa
		_bfc[_fa], _bfc[_db] = _bfc[_db], _bfc[_fa]
	}
	return _bfc
}

const _ebdc int = -1
