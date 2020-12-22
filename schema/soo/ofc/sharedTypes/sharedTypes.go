package sharedTypes

import (
	_d "encoding/xml"
	_c "fmt"
	_e "regexp"
)

func (_cc ST_OnOff) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _cc.Bool != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", _dd(*_cc.Bool))))
	}
	if _cc.ST_OnOff1 != ST_OnOff1Unset {
		e.EncodeToken(_d.CharData(_cc.ST_OnOff1.String()))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}

const ST_PercentagePattern = "-\u003f[\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e[\u0030\u002d\u0039\u005d+)\u003f\u0025"

func (_baaa ST_TrueFalse) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_baaa.String(), start)
}
func (_ebd *ST_ConformanceClass) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ebd = 0
	case "\u0073\u0074\u0072\u0069\u0063\u0074":
		*_ebd = 1
	case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":
		*_ebd = 2
	}
	return nil
}
func (_gaac ST_TrueFalseBlank) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_deb := _d.Attr{}
	_deb.Name = name
	switch _gaac {
	case ST_TrueFalseBlankUnset:
		_deb.Value = ""
	case ST_TrueFalseBlankT:
		_deb.Value = "\u0074"
	case ST_TrueFalseBlankF:
		_deb.Value = "\u0066"
	case ST_TrueFalseBlankTrue:
		_deb.Value = "\u0074\u0072\u0075\u0065"
	case ST_TrueFalseBlankFalse:
		_deb.Value = "\u0066\u0061\u006cs\u0065"
	case ST_TrueFalseBlankTrue_:
		_deb.Value = "\u0054\u0072\u0075\u0065"
	case ST_TrueFalseBlankFalse_:
		_deb.Value = "\u0046\u0061\u006cs\u0065"
	}
	return _deb, nil
}
func (_bac *ST_VerticalAlignRun) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bac = 0
	case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":
		*_bac = 1
	case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":
		*_bac = 2
	case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":
		*_bac = 3
	}
	return nil
}
func (_cgc ST_CryptProv) Validate() error { return _cgc.ValidateWithPath("") }
func (_bgce ST_ConformanceClass) ValidateWithPath(path string) error {
	switch _bgce {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bgce))
	}
	return nil
}
func (_g *ST_OnOff) Validate() error { return _g.ValidateWithPath("") }

type ST_TrueFalseBlank byte

func (_add ST_CryptProv) ValidateWithPath(path string) error {
	switch _add {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_add))
	}
	return nil
}
func (_fg ST_AlgClass) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ff := _d.Attr{}
	_ff.Name = name
	switch _fg {
	case ST_AlgClassUnset:
		_ff.Value = ""
	case ST_AlgClassHash:
		_ff.Value = "\u0068\u0061\u0073\u0068"
	case ST_AlgClassCustom:
		_ff.Value = "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return _ff, nil
}
func (_fcc ST_TrueFalseBlank) Validate() error { return _fcc.ValidateWithPath("") }
func (_fd ST_OnOff1) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_fc := _d.Attr{}
	_fc.Name = name
	switch _fd {
	case ST_OnOff1Unset:
		_fc.Value = ""
	case ST_OnOff1On:
		_fc.Value = "\u006f\u006e"
	case ST_OnOff1Off:
		_fc.Value = "\u006f\u0066\u0066"
	}
	return _fc, nil
}
func (_ecd *ST_CryptProv) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gbc, _ab := d.Token()
	if _ab != nil {
		return _ab
	}
	if _cg, _fb := _gbc.(_d.EndElement); _fb && _cg.Name == start.Name {
		*_ecd = 1
		return nil
	}
	if _ca, _ae := _gbc.(_d.CharData); !_ae {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbc)
	} else {
		switch string(_ca) {
		case "":
			*_ecd = 0
		case "\u0072\u0073\u0061\u0041\u0045\u0053":
			*_ecd = 1
		case "\u0072s\u0061\u0046\u0075\u006c\u006c":
			*_ecd = 2
		case "\u0063\u0075\u0073\u0074\u006f\u006d":
			*_ecd = 3
		}
	}
	_gbc, _ab = d.Token()
	if _ab != nil {
		return _ab
	}
	if _bc, _ada := _gbc.(_d.EndElement); _ada && _bc.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbc)
}

const ST_PositivePercentagePattern = "\u005b0\u002d9\u005d\u002b\u0028\u005c\u002e[\u0030\u002d9\u005d\u002b\u0029\u003f\u0025"

func (_acc *ST_AlgClass) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_acc = 0
	case "\u0068\u0061\u0073\u0068":
		*_acc = 1
	case "\u0063\u0075\u0073\u0074\u006f\u006d":
		*_acc = 2
	}
	return nil
}

type ST_TrueFalse byte

const ST_UniversalMeasurePattern = "\u002d\u003f\u005b\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e\u005b\u0030\u002d\u0039\u005d\u002b\u0029\u003f\u0028\u006d\u006d\u007c\u0063m\u007c\u0069\u006e\u007c\u0070t\u007c\u0070c\u007c\u0070\u0069\u0029"

type ST_AlgClass byte

func (_ad ST_CalendarType) String() string {
	switch _ad {
	case 0:
		return ""
	case 1:
		return "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n"
	case 2:
		return "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073"
	case 3:
		return "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068"
	case 4:
		return "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063"
	case 5:
		return "\u0068\u0069\u006ar\u0069"
	case 6:
		return "\u0068\u0065\u0062\u0072\u0065\u0077"
	case 7:
		return "\u0074\u0061\u0069\u0077\u0061\u006e"
	case 8:
		return "\u006a\u0061\u0070a\u006e"
	case 9:
		return "\u0074\u0068\u0061\u0069"
	case 10:
		return "\u006b\u006f\u0072e\u0061"
	case 11:
		return "\u0073\u0061\u006b\u0061"
	case 12:
		return "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068"
	case 13:
		return "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068"
	case 14:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}
func (_bb *ST_OnOff1) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bb = 0
	case "\u006f\u006e":
		*_bb = 1
	case "\u006f\u0066\u0066":
		*_bb = 2
	}
	return nil
}
func (_bae ST_CryptProv) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_bae.String(), start)
}
func (_baeb ST_XAlign) String() string {
	switch _baeb {
	case 0:
		return ""
	case 1:
		return "\u006c\u0065\u0066\u0074"
	case 2:
		return "\u0063\u0065\u006e\u0074\u0065\u0072"
	case 3:
		return "\u0072\u0069\u0067h\u0074"
	case 4:
		return "\u0069\u006e\u0073\u0069\u0064\u0065"
	case 5:
		return "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return ""
}
func (_agaa ST_ConformanceClass) String() string {
	switch _agaa {
	case 0:
		return ""
	case 1:
		return "\u0073\u0074\u0072\u0069\u0063\u0074"
	case 2:
		return "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c"
	}
	return ""
}
func (_gff *ST_AlgType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gff = 0
	case "\u0074y\u0070\u0065\u0041\u006e\u0079":
		*_gff = 1
	case "\u0063\u0075\u0073\u0074\u006f\u006d":
		*_gff = 2
	}
	return nil
}

type ST_VerticalAlignRun byte

func (_ccc *ST_OnOff) ValidateWithPath(path string) error {
	_db := []string{}
	if _ccc.Bool != nil {
		_db = append(_db, "\u0042\u006f\u006f\u006c")
	}
	if _ccc.ST_OnOff1 != ST_OnOff1Unset {
		_db = append(_db, "\u0053T\u005f\u004f\u006e\u004f\u0066\u00661")
	}
	if len(_db) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _db)
	}
	return nil
}

const (
	ST_AlgClassUnset  ST_AlgClass = 0
	ST_AlgClassHash   ST_AlgClass = 1
	ST_AlgClassCustom ST_AlgClass = 2
)

func (_gc *ST_CalendarType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ef, _fa := d.Token()
	if _fa != nil {
		return _fa
	}
	if _aa, _dbf := _ef.(_d.EndElement); _dbf && _aa.Name == start.Name {
		*_gc = 1
		return nil
	}
	if _ag, _gf := _ef.(_d.CharData); !_gf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ef)
	} else {
		switch string(_ag) {
		case "":
			*_gc = 0
		case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":
			*_gc = 1
		case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":
			*_gc = 2
		case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":
			*_gc = 3
		case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":
			*_gc = 4
		case "\u0068\u0069\u006ar\u0069":
			*_gc = 5
		case "\u0068\u0065\u0062\u0072\u0065\u0077":
			*_gc = 6
		case "\u0074\u0061\u0069\u0077\u0061\u006e":
			*_gc = 7
		case "\u006a\u0061\u0070a\u006e":
			*_gc = 8
		case "\u0074\u0068\u0061\u0069":
			*_gc = 9
		case "\u006b\u006f\u0072e\u0061":
			*_gc = 10
		case "\u0073\u0061\u006b\u0061":
			*_gc = 11
		case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":
			*_gc = 12
		case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":
			*_gc = 13
		case "\u006e\u006f\u006e\u0065":
			*_gc = 14
		}
	}
	_ef, _fa = d.Token()
	if _fa != nil {
		return _fa
	}
	if _bd, _dg := _ef.(_d.EndElement); _dg && _bd.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ef)
}
func (_ccd ST_YAlign) String() string {
	switch _ccd {
	case 0:
		return ""
	case 1:
		return "\u0069\u006e\u006c\u0069\u006e\u0065"
	case 2:
		return "\u0074\u006f\u0070"
	case 3:
		return "\u0063\u0065\u006e\u0074\u0065\u0072"
	case 4:
		return "\u0062\u006f\u0074\u0074\u006f\u006d"
	case 5:
		return "\u0069\u006e\u0073\u0069\u0064\u0065"
	case 6:
		return "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return ""
}
func (_aad ST_TrueFalse) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gag := _d.Attr{}
	_gag.Name = name
	switch _aad {
	case ST_TrueFalseUnset:
		_gag.Value = ""
	case ST_TrueFalseT:
		_gag.Value = "\u0074"
	case ST_TrueFalseF:
		_gag.Value = "\u0066"
	case ST_TrueFalseTrue:
		_gag.Value = "\u0074\u0072\u0075\u0065"
	case ST_TrueFalseFalse:
		_gag.Value = "\u0066\u0061\u006cs\u0065"
	}
	return _gag, nil
}

const (
	ST_CalendarTypeUnset                ST_CalendarType = 0
	ST_CalendarTypeGregorian            ST_CalendarType = 1
	ST_CalendarTypeGregorianUs          ST_CalendarType = 2
	ST_CalendarTypeGregorianMeFrench    ST_CalendarType = 3
	ST_CalendarTypeGregorianArabic      ST_CalendarType = 4
	ST_CalendarTypeHijri                ST_CalendarType = 5
	ST_CalendarTypeHebrew               ST_CalendarType = 6
	ST_CalendarTypeTaiwan               ST_CalendarType = 7
	ST_CalendarTypeJapan                ST_CalendarType = 8
	ST_CalendarTypeThai                 ST_CalendarType = 9
	ST_CalendarTypeKorea                ST_CalendarType = 10
	ST_CalendarTypeSaka                 ST_CalendarType = 11
	ST_CalendarTypeGregorianXlitEnglish ST_CalendarType = 12
	ST_CalendarTypeGregorianXlitFrench  ST_CalendarType = 13
	ST_CalendarTypeNone                 ST_CalendarType = 14
)

func (_ebg ST_AlgClass) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ebg.String(), start)
}
func (_eff *ST_ConformanceClass) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_acdg, _debg := d.Token()
	if _debg != nil {
		return _debg
	}
	if _bgd, _aafe := _acdg.(_d.EndElement); _aafe && _bgd.Name == start.Name {
		*_eff = 1
		return nil
	}
	if _fe, _dbg := _acdg.(_d.CharData); !_dbg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _acdg)
	} else {
		switch string(_fe) {
		case "":
			*_eff = 0
		case "\u0073\u0074\u0072\u0069\u0063\u0074":
			*_eff = 1
		case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":
			*_eff = 2
		}
	}
	_acdg, _debg = d.Token()
	if _debg != nil {
		return _debg
	}
	if _gfe, _ebda := _acdg.(_d.EndElement); _ebda && _gfe.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _acdg)
}
func (_ffe *ST_XAlign) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_abb, _gffe := d.Token()
	if _gffe != nil {
		return _gffe
	}
	if _efd, _egb := _abb.(_d.EndElement); _egb && _efd.Name == start.Name {
		*_ffe = 1
		return nil
	}
	if _cac, _aaf := _abb.(_d.CharData); !_aaf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _abb)
	} else {
		switch string(_cac) {
		case "":
			*_ffe = 0
		case "\u006c\u0065\u0066\u0074":
			*_ffe = 1
		case "\u0063\u0065\u006e\u0074\u0065\u0072":
			*_ffe = 2
		case "\u0072\u0069\u0067h\u0074":
			*_ffe = 3
		case "\u0069\u006e\u0073\u0069\u0064\u0065":
			*_ffe = 4
		case "\u006fu\u0074\u0073\u0069\u0064\u0065":
			*_ffe = 5
		}
	}
	_abb, _gffe = d.Token()
	if _gffe != nil {
		return _gffe
	}
	if _fcg, _dga := _abb.(_d.EndElement); _dga && _fcg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _abb)
}
func (_dgaf ST_YAlign) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_dgaf.String(), start)
}
func (_gcd ST_CryptProv) String() string {
	switch _gcd {
	case 0:
		return ""
	case 1:
		return "\u0072\u0073\u0061\u0041\u0045\u0053"
	case 2:
		return "\u0072s\u0061\u0046\u0075\u006c\u006c"
	case 3:
		return "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return ""
}

const (
	ST_TrueFalseBlankUnset  ST_TrueFalseBlank = 0
	ST_TrueFalseBlankT      ST_TrueFalseBlank = 1
	ST_TrueFalseBlankF      ST_TrueFalseBlank = 2
	ST_TrueFalseBlankTrue   ST_TrueFalseBlank = 3
	ST_TrueFalseBlankFalse  ST_TrueFalseBlank = 4
	ST_TrueFalseBlankTrue_  ST_TrueFalseBlank = 6
	ST_TrueFalseBlankFalse_ ST_TrueFalseBlank = 7
)

func (_deg ST_TrueFalse) String() string {
	switch _deg {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0066"
	case 3:
		return "\u0074\u0072\u0075\u0065"
	case 4:
		return "\u0066\u0061\u006cs\u0065"
	}
	return ""
}

const ST_FixedPercentagePattern = "-\u003f\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039]\u003f\u0029\u0029\u0028\u005c\u002e\u005b\u0030\u002d\u0039][\u0030\u002d\u0039]\u003f)\u003f\u0025"

func (_de ST_OnOff) String() string {
	if _de.Bool != nil {
		return _c.Sprintf("\u0025\u0076", *_de.Bool)
	}
	if _de.ST_OnOff1 != ST_OnOff1Unset {
		return _de.ST_OnOff1.String()
	}
	return ""
}

type ST_OnOff1 byte

func (_cea ST_AlgType) ValidateWithPath(path string) error {
	switch _cea {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cea))
	}
	return nil
}
func _dd(_ac bool) uint8 {
	if _ac {
		return 1
	}
	return 0
}

type ST_XAlign byte

func (_fbf ST_OnOff1) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fbf.String(), start)
}
func (_aga ST_AlgType) String() string {
	switch _aga {
	case 0:
		return ""
	case 1:
		return "\u0074y\u0070\u0065\u0041\u006e\u0079"
	case 2:
		return "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return ""
}

type ST_AlgType byte

func (_bcd *ST_XAlign) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bcd = 0
	case "\u006c\u0065\u0066\u0074":
		*_bcd = 1
	case "\u0063\u0065\u006e\u0074\u0065\u0072":
		*_bcd = 2
	case "\u0072\u0069\u0067h\u0074":
		*_bcd = 3
	case "\u0069\u006e\u0073\u0069\u0064\u0065":
		*_bcd = 4
	case "\u006fu\u0074\u0073\u0069\u0064\u0065":
		*_bcd = 5
	}
	return nil
}
func (_agb ST_AlgClass) ValidateWithPath(path string) error {
	switch _agb {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_agb))
	}
	return nil
}
func (_fgf ST_VerticalAlignRun) String() string {
	switch _fgf {
	case 0:
		return ""
	case 1:
		return "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065"
	case 2:
		return "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074"
	case 3:
		return "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t"
	}
	return ""
}

const (
	ST_AlgTypeUnset   ST_AlgType = 0
	ST_AlgTypeTypeAny ST_AlgType = 1
	ST_AlgTypeCustom  ST_AlgType = 2
)

func (_gfa ST_OnOff1) String() string {
	switch _gfa {
	case 0:
		return ""
	case 1:
		return "\u006f\u006e"
	case 2:
		return "\u006f\u0066\u0066"
	}
	return ""
}

type ST_CryptProv byte

func (_aba *ST_AlgType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_eed, _bdd := d.Token()
	if _bdd != nil {
		return _bdd
	}
	if _af, _def := _eed.(_d.EndElement); _def && _af.Name == start.Name {
		*_aba = 1
		return nil
	}
	if _bag, _dcg := _eed.(_d.CharData); !_dcg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eed)
	} else {
		switch string(_bag) {
		case "":
			*_aba = 0
		case "\u0074y\u0070\u0065\u0041\u006e\u0079":
			*_aba = 1
		case "\u0063\u0075\u0073\u0074\u006f\u006d":
			*_aba = 2
		}
	}
	_eed, _bdd = d.Token()
	if _bdd != nil {
		return _bdd
	}
	if _aae, _abe := _eed.(_d.EndElement); _abe && _aae.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eed)
}
func (_bgc ST_TrueFalseBlank) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_bgc.String(), start)
}
func (_baa ST_AlgType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_eeg := _d.Attr{}
	_eeg.Name = name
	switch _baa {
	case ST_AlgTypeUnset:
		_eeg.Value = ""
	case ST_AlgTypeTypeAny:
		_eeg.Value = "\u0074y\u0070\u0065\u0041\u006e\u0079"
	case ST_AlgTypeCustom:
		_eeg.Value = "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return _eeg, nil
}

const (
	ST_VerticalAlignRunUnset       ST_VerticalAlignRun = 0
	ST_VerticalAlignRunBaseline    ST_VerticalAlignRun = 1
	ST_VerticalAlignRunSuperscript ST_VerticalAlignRun = 2
	ST_VerticalAlignRunSubscript   ST_VerticalAlignRun = 3
)

var ST_PositiveFixedPercentagePatternRe = _e.MustCompile(ST_PositiveFixedPercentagePattern)

func (_df ST_XAlign) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gga := _d.Attr{}
	_gga.Name = name
	switch _df {
	case ST_XAlignUnset:
		_gga.Value = ""
	case ST_XAlignLeft:
		_gga.Value = "\u006c\u0065\u0066\u0074"
	case ST_XAlignCenter:
		_gga.Value = "\u0063\u0065\u006e\u0074\u0065\u0072"
	case ST_XAlignRight:
		_gga.Value = "\u0072\u0069\u0067h\u0074"
	case ST_XAlignInside:
		_gga.Value = "\u0069\u006e\u0073\u0069\u0064\u0065"
	case ST_XAlignOutside:
		_gga.Value = "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return _gga, nil
}
func (_agf ST_ConformanceClass) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dfa := _d.Attr{}
	_dfa.Name = name
	switch _agf {
	case ST_ConformanceClassUnset:
		_dfa.Value = ""
	case ST_ConformanceClassStrict:
		_dfa.Value = "\u0073\u0074\u0072\u0069\u0063\u0074"
	case ST_ConformanceClassTransitional:
		_dfa.Value = "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c"
	}
	return _dfa, nil
}

const ST_PositiveUniversalMeasurePattern = "\u005b\u0030-9\u005d\u002b\u0028\\\u002e\u005b\u0030\u002d9]+\u0029?(\u006d\u006d\u007c\u0063\u006d\u007c\u0069n|\u0070\u0074\u007c\u0070\u0063\u007c\u0070i\u0029"

func (_dcgf ST_OnOff1) Validate() error { return _dcgf.ValidateWithPath("") }
func (_ba ST_CalendarType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gbg := _d.Attr{}
	_gbg.Name = name
	switch _ba {
	case ST_CalendarTypeUnset:
		_gbg.Value = ""
	case ST_CalendarTypeGregorian:
		_gbg.Value = "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n"
	case ST_CalendarTypeGregorianUs:
		_gbg.Value = "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073"
	case ST_CalendarTypeGregorianMeFrench:
		_gbg.Value = "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068"
	case ST_CalendarTypeGregorianArabic:
		_gbg.Value = "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063"
	case ST_CalendarTypeHijri:
		_gbg.Value = "\u0068\u0069\u006ar\u0069"
	case ST_CalendarTypeHebrew:
		_gbg.Value = "\u0068\u0065\u0062\u0072\u0065\u0077"
	case ST_CalendarTypeTaiwan:
		_gbg.Value = "\u0074\u0061\u0069\u0077\u0061\u006e"
	case ST_CalendarTypeJapan:
		_gbg.Value = "\u006a\u0061\u0070a\u006e"
	case ST_CalendarTypeThai:
		_gbg.Value = "\u0074\u0068\u0061\u0069"
	case ST_CalendarTypeKorea:
		_gbg.Value = "\u006b\u006f\u0072e\u0061"
	case ST_CalendarTypeSaka:
		_gbg.Value = "\u0073\u0061\u006b\u0061"
	case ST_CalendarTypeGregorianXlitEnglish:
		_gbg.Value = "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068"
	case ST_CalendarTypeGregorianXlitFrench:
		_gbg.Value = "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068"
	case ST_CalendarTypeNone:
		_gbg.Value = "\u006e\u006f\u006e\u0065"
	}
	return _gbg, nil
}
func (_aed ST_XAlign) ValidateWithPath(path string) error {
	switch _aed {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aed))
	}
	return nil
}
func (_f ST_CalendarType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_f.String(), start)
}

const (
	ST_CryptProvUnset   ST_CryptProv = 0
	ST_CryptProvRsaAES  ST_CryptProv = 1
	ST_CryptProvRsaFull ST_CryptProv = 2
	ST_CryptProvCustom  ST_CryptProv = 3
)

// ST_TwipsMeasure is a union type
type ST_TwipsMeasure struct {
	ST_UnsignedDecimalNumber    *uint64
	ST_PositiveUniversalMeasure *string
}

func (_ded ST_VerticalAlignRun) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ded.String(), start)
}
func (_agd ST_TrueFalse) ValidateWithPath(path string) error {
	switch _agd {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_agd))
	}
	return nil
}
func (_cce ST_VerticalAlignRun) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_baae := _d.Attr{}
	_baae.Name = name
	switch _cce {
	case ST_VerticalAlignRunUnset:
		_baae.Value = ""
	case ST_VerticalAlignRunBaseline:
		_baae.Value = "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065"
	case ST_VerticalAlignRunSuperscript:
		_baae.Value = "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074"
	case ST_VerticalAlignRunSubscript:
		_baae.Value = "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t"
	}
	return _baae, nil
}
func (_cccg ST_VerticalAlignRun) Validate() error { return _cccg.ValidateWithPath("") }
func (_adg ST_ConformanceClass) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_adg.String(), start)
}
func (_ebfe ST_VerticalAlignRun) ValidateWithPath(path string) error {
	switch _ebfe {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ebfe))
	}
	return nil
}

var ST_PercentagePatternRe = _e.MustCompile(ST_PercentagePattern)

func (_gaf *ST_TrueFalse) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gaf = 0
	case "\u0074":
		*_gaf = 1
	case "\u0066":
		*_gaf = 2
	case "\u0074\u0072\u0075\u0065":
		*_gaf = 3
	case "\u0066\u0061\u006cs\u0065":
		*_gaf = 4
	}
	return nil
}
func (_cb *ST_TrueFalse) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bf, _fgd := d.Token()
	if _fgd != nil {
		return _fgd
	}
	if _gafd, _bfe := _bf.(_d.EndElement); _bfe && _gafd.Name == start.Name {
		*_cb = 1
		return nil
	}
	if _ggb, _gaa := _bf.(_d.CharData); !_gaa {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bf)
	} else {
		switch string(_ggb) {
		case "":
			*_cb = 0
		case "\u0074":
			*_cb = 1
		case "\u0066":
			*_cb = 2
		case "\u0074\u0072\u0075\u0065":
			*_cb = 3
		case "\u0066\u0061\u006cs\u0065":
			*_cb = 4
		}
	}
	_bf, _fgd = d.Token()
	if _fgd != nil {
		return _fgd
	}
	if _gbf, _adc := _bf.(_d.EndElement); _adc && _gbf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bf)
}
func (_eba *ST_YAlign) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cee, _gdd := d.Token()
	if _gdd != nil {
		return _gdd
	}
	if _da, _fgc := _cee.(_d.EndElement); _fgc && _da.Name == start.Name {
		*_eba = 1
		return nil
	}
	if _ed, _accb := _cee.(_d.CharData); !_accb {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cee)
	} else {
		switch string(_ed) {
		case "":
			*_eba = 0
		case "\u0069\u006e\u006c\u0069\u006e\u0065":
			*_eba = 1
		case "\u0074\u006f\u0070":
			*_eba = 2
		case "\u0063\u0065\u006e\u0074\u0065\u0072":
			*_eba = 3
		case "\u0062\u006f\u0074\u0074\u006f\u006d":
			*_eba = 4
		case "\u0069\u006e\u0073\u0069\u0064\u0065":
			*_eba = 5
		case "\u006fu\u0074\u0073\u0069\u0064\u0065":
			*_eba = 6
		}
	}
	_cee, _gdd = d.Token()
	if _gdd != nil {
		return _gdd
	}
	if _dee, _dae := _cee.(_d.EndElement); _dae && _dee.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cee)
}
func (_cd *ST_TwipsMeasure) Validate() error { return _cd.ValidateWithPath("") }
func (_baad ST_OnOff1) ValidateWithPath(path string) error {
	switch _baad {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_baad))
	}
	return nil
}
func (_ddb ST_XAlign) Validate() error { return _ddb.ValidateWithPath("") }

var ST_PositiveUniversalMeasurePatternRe = _e.MustCompile(ST_PositiveUniversalMeasurePattern)

// ST_OnOff is a union type
type ST_OnOff struct {
	Bool      *bool
	ST_OnOff1 ST_OnOff1
}

func (_b ST_TwipsMeasure) String() string {
	if _b.ST_UnsignedDecimalNumber != nil {
		return _c.Sprintf("\u0025\u0076", *_b.ST_UnsignedDecimalNumber)
	}
	if _b.ST_PositiveUniversalMeasure != nil {
		return _c.Sprintf("\u0025\u0076", *_b.ST_PositiveUniversalMeasure)
	}
	return ""
}
func (_ga ST_AlgType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ga.String(), start)
}
func (_gfg ST_ConformanceClass) Validate() error { return _gfg.ValidateWithPath("") }
func (_be *ST_TrueFalseBlank) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dgdc, _efb := d.Token()
	if _efb != nil {
		return _efb
	}
	if _dda, _gab := _dgdc.(_d.EndElement); _gab && _dda.Name == start.Name {
		*_be = 1
		return nil
	}
	if _gef, _gfc := _dgdc.(_d.CharData); !_gfc {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dgdc)
	} else {
		switch string(_gef) {
		case "":
			*_be = 0
		case "\u0074":
			*_be = 1
		case "\u0066":
			*_be = 2
		case "\u0074\u0072\u0075\u0065":
			*_be = 3
		case "\u0066\u0061\u006cs\u0065":
			*_be = 4
		case "\u0054\u0072\u0075\u0065":
			*_be = 6
		case "\u0046\u0061\u006cs\u0065":
			*_be = 7
		}
	}
	_dgdc, _efb = d.Token()
	if _efb != nil {
		return _efb
	}
	if _aade, _ddd := _dgdc.(_d.EndElement); _ddd && _aade.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dgdc)
}

var ST_GuidPatternRe = _e.MustCompile(ST_GuidPattern)

const ST_PositiveFixedPercentagePattern = "\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039\u005d\u003f\u0029\u0029\u0028\u005c\u002e[\u0030\u002d\u0039\u005d\u005b0\u002d\u0039]\u003f\u0029\u003f\u0025"

func (_fbdd *ST_YAlign) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fbdd = 0
	case "\u0069\u006e\u006c\u0069\u006e\u0065":
		*_fbdd = 1
	case "\u0074\u006f\u0070":
		*_fbdd = 2
	case "\u0063\u0065\u006e\u0074\u0065\u0072":
		*_fbdd = 3
	case "\u0062\u006f\u0074\u0074\u006f\u006d":
		*_fbdd = 4
	case "\u0069\u006e\u0073\u0069\u0064\u0065":
		*_fbdd = 5
	case "\u006fu\u0074\u0073\u0069\u0064\u0065":
		*_fbdd = 6
	}
	return nil
}
func (_bab ST_TrueFalseBlank) ValidateWithPath(path string) error {
	switch _bab {
	case 0, 1, 2, 3, 4, 6, 7:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bab))
	}
	return nil
}
func (_bce ST_YAlign) ValidateWithPath(path string) error {
	switch _bce {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bce))
	}
	return nil
}

var ST_FixedPercentagePatternRe = _e.MustCompile(ST_FixedPercentagePattern)

const (
	ST_XAlignUnset   ST_XAlign = 0
	ST_XAlignLeft    ST_XAlign = 1
	ST_XAlignCenter  ST_XAlign = 2
	ST_XAlignRight   ST_XAlign = 3
	ST_XAlignInside  ST_XAlign = 4
	ST_XAlignOutside ST_XAlign = 5
)
const (
	ST_ConformanceClassUnset        ST_ConformanceClass = 0
	ST_ConformanceClassStrict       ST_ConformanceClass = 1
	ST_ConformanceClassTransitional ST_ConformanceClass = 2
)

func (_cec *ST_AlgClass) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fff, _dc := d.Token()
	if _dc != nil {
		return _dc
	}
	if _cecd, _ccg := _fff.(_d.EndElement); _ccg && _cecd.Name == start.Name {
		*_cec = 1
		return nil
	}
	if _ea, _ebf := _fff.(_d.CharData); !_ebf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fff)
	} else {
		switch string(_ea) {
		case "":
			*_cec = 0
		case "\u0068\u0061\u0073\u0068":
			*_cec = 1
		case "\u0063\u0075\u0073\u0074\u006f\u006d":
			*_cec = 2
		}
	}
	_fff, _dc = d.Token()
	if _dc != nil {
		return _dc
	}
	if _ee, _cff := _fff.(_d.EndElement); _cff && _ee.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fff)
}
func ParseUnionST_OnOff(s string) (ST_OnOff, error) {
	_gd := ST_OnOff{}
	switch s {
	case "\u0074\u0072\u0075\u0065", "\u0031", "\u006f\u006e":
		_cdb := true
		_gd.Bool = &_cdb
	default:
		_eb := false
		_gd.Bool = &_eb
	}
	return _gd, nil
}
func (_gca ST_CalendarType) ValidateWithPath(path string) error {
	switch _gca {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gca))
	}
	return nil
}
func (_dgd ST_CalendarType) Validate() error { return _dgd.ValidateWithPath("") }
func (_ec ST_AlgClass) Validate() error      { return _ec.ValidateWithPath("") }

const ST_GuidPattern = "\u005c\u007b\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b\u0038\u007d\u002d\u005b\u0030\u002d9\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030-\u0039\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b4\u007d\u002d\u005b\u0030\u002d\u0039A\u002d\u0046]\u007b\u00312\u007d\\\u007d"

func (_dea *ST_TwipsMeasure) ValidateWithPath(path string) error {
	_gb := []string{}
	if _dea.ST_UnsignedDecimalNumber != nil {
		_gb = append(_gb, "\u0053T\u005f\u0055\u006e\u0073\u0069\u0067\u006e\u0065\u0064\u0044\u0065c\u0069\u006d\u0061\u006c\u004e\u0075\u006d\u0062\u0065\u0072")
	}
	if _dea.ST_PositiveUniversalMeasure != nil {
		_gb = append(_gb, "S\u0054\u005f\u0050\u006f\u0073\u0069t\u0069\u0076\u0065\u0055\u006e\u0069\u0076\u0065\u0072s\u0061\u006c\u004de\u0061s\u0075\u0072\u0065")
	}
	if len(_gb) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _gb)
	}
	return nil
}

type ST_ConformanceClass byte

const (
	ST_OnOff1Unset ST_OnOff1 = 0
	ST_OnOff1On    ST_OnOff1 = 1
	ST_OnOff1Off   ST_OnOff1 = 2
)

func (_bddc *ST_OnOff1) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bg, _gg := d.Token()
	if _gg != nil {
		return _gg
	}
	if _bdg, _eg := _bg.(_d.EndElement); _eg && _bdg.Name == start.Name {
		*_bddc = 1
		return nil
	}
	if _ddg, _aea := _bg.(_d.CharData); !_aea {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bg)
	} else {
		switch string(_ddg) {
		case "":
			*_bddc = 0
		case "\u006f\u006e":
			*_bddc = 1
		case "\u006f\u0066\u0066":
			*_bddc = 2
		}
	}
	_bg, _gg = d.Token()
	if _gg != nil {
		return _gg
	}
	if _cfd, _fbd := _bg.(_d.EndElement); _fbd && _cfd.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bg)
}

const (
	ST_TrueFalseUnset ST_TrueFalse = 0
	ST_TrueFalseT     ST_TrueFalse = 1
	ST_TrueFalseF     ST_TrueFalse = 2
	ST_TrueFalseTrue  ST_TrueFalse = 3
	ST_TrueFalseFalse ST_TrueFalse = 4
)

var ST_UniversalMeasurePatternRe = _e.MustCompile(ST_UniversalMeasurePattern)

func (_cbd ST_YAlign) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ccca := _d.Attr{}
	_ccca.Name = name
	switch _cbd {
	case ST_YAlignUnset:
		_ccca.Value = ""
	case ST_YAlignInline:
		_ccca.Value = "\u0069\u006e\u006c\u0069\u006e\u0065"
	case ST_YAlignTop:
		_ccca.Value = "\u0074\u006f\u0070"
	case ST_YAlignCenter:
		_ccca.Value = "\u0063\u0065\u006e\u0074\u0065\u0072"
	case ST_YAlignBottom:
		_ccca.Value = "\u0062\u006f\u0074\u0074\u006f\u006d"
	case ST_YAlignInside:
		_ccca.Value = "\u0069\u006e\u0073\u0069\u0064\u0065"
	case ST_YAlignOutside:
		_ccca.Value = "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return _ccca, nil
}

var ST_PositivePercentagePatternRe = _e.MustCompile(ST_PositivePercentagePattern)

func (_ffd ST_AlgClass) String() string {
	switch _ffd {
	case 0:
		return ""
	case 1:
		return "\u0068\u0061\u0073\u0068"
	case 2:
		return "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return ""
}
func (_cf ST_TwipsMeasure) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _cf.ST_UnsignedDecimalNumber != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", *_cf.ST_UnsignedDecimalNumber)))
	}
	if _cf.ST_PositiveUniversalMeasure != nil {
		e.EncodeToken(_d.CharData(*_cf.ST_PositiveUniversalMeasure))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}
func (_ecda ST_YAlign) Validate() error { return _ecda.ValidateWithPath("") }
func (_ffg ST_AlgType) Validate() error { return _ffg.ValidateWithPath("") }
func (_eec ST_XAlign) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_eec.String(), start)
}
func (_gfb ST_TrueFalseBlank) String() string {
	switch _gfb {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0066"
	case 3:
		return "\u0074\u0072\u0075\u0065"
	case 4:
		return "\u0066\u0061\u006cs\u0065"
	case 6:
		return "\u0054\u0072\u0075\u0065"
	case 7:
		return "\u0046\u0061\u006cs\u0065"
	}
	return ""
}

type ST_CalendarType byte

func (_fga *ST_CryptProv) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fga = 0
	case "\u0072\u0073\u0061\u0041\u0045\u0053":
		*_fga = 1
	case "\u0072s\u0061\u0046\u0075\u006c\u006c":
		*_fga = 2
	case "\u0063\u0075\u0073\u0074\u006f\u006d":
		*_fga = 3
	}
	return nil
}
func (_bdb *ST_VerticalAlignRun) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gbd, _eee := d.Token()
	if _eee != nil {
		return _eee
	}
	if _dde, _bbd := _gbd.(_d.EndElement); _bbd && _dde.Name == start.Name {
		*_bdb = 1
		return nil
	}
	if _cdg, _fde := _gbd.(_d.CharData); !_fde {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbd)
	} else {
		switch string(_cdg) {
		case "":
			*_bdb = 0
		case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":
			*_bdb = 1
		case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":
			*_bdb = 2
		case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":
			*_bdb = 3
		}
	}
	_gbd, _eee = d.Token()
	if _eee != nil {
		return _eee
	}
	if _ggc, _acd := _gbd.(_d.EndElement); _acd && _ggc.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbd)
}
func (_ge ST_CryptProv) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_eea := _d.Attr{}
	_eea.Name = name
	switch _ge {
	case ST_CryptProvUnset:
		_eea.Value = ""
	case ST_CryptProvRsaAES:
		_eea.Value = "\u0072\u0073\u0061\u0041\u0045\u0053"
	case ST_CryptProvRsaFull:
		_eea.Value = "\u0072s\u0061\u0046\u0075\u006c\u006c"
	case ST_CryptProvCustom:
		_eea.Value = "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return _eea, nil
}
func (_eeb ST_TrueFalse) Validate() error { return _eeb.ValidateWithPath("") }
func (_gdb *ST_TrueFalseBlank) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gdb = 0
	case "\u0074":
		*_gdb = 1
	case "\u0066":
		*_gdb = 2
	case "\u0074\u0072\u0075\u0065":
		*_gdb = 3
	case "\u0066\u0061\u006cs\u0065":
		*_gdb = 4
	case "\u0054\u0072\u0075\u0065":
		*_gdb = 6
	case "\u0046\u0061\u006cs\u0065":
		*_gdb = 7
	}
	return nil
}
func (_ce *ST_CalendarType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ce = 0
	case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":
		*_ce = 1
	case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":
		*_ce = 2
	case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":
		*_ce = 3
	case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":
		*_ce = 4
	case "\u0068\u0069\u006ar\u0069":
		*_ce = 5
	case "\u0068\u0065\u0062\u0072\u0065\u0077":
		*_ce = 6
	case "\u0074\u0061\u0069\u0077\u0061\u006e":
		*_ce = 7
	case "\u006a\u0061\u0070a\u006e":
		*_ce = 8
	case "\u0074\u0068\u0061\u0069":
		*_ce = 9
	case "\u006b\u006f\u0072e\u0061":
		*_ce = 10
	case "\u0073\u0061\u006b\u0061":
		*_ce = 11
	case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":
		*_ce = 12
	case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":
		*_ce = 13
	case "\u006e\u006f\u006e\u0065":
		*_ce = 14
	}
	return nil
}

type ST_YAlign byte

const (
	ST_YAlignUnset   ST_YAlign = 0
	ST_YAlignInline  ST_YAlign = 1
	ST_YAlignTop     ST_YAlign = 2
	ST_YAlignCenter  ST_YAlign = 3
	ST_YAlignBottom  ST_YAlign = 4
	ST_YAlignInside  ST_YAlign = 5
	ST_YAlignOutside ST_YAlign = 6
)
