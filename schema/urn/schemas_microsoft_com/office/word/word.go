package word

import (
	_d "encoding/xml"
	_ce "fmt"
	_c "strconv"

	_cg "gitee.com/yu_sheng/gooffice"
)

func (_de *Anchorlock) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0061\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b"
	return _de.CT_AnchorLock.MarshalXML(e, start)
}
func (_dg *Borderright) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "b\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074"
	return _dg.CT_Border.MarshalXML(e, start)
}

type Wrap struct{ CT_Wrap }

func (_egde ST_HorizontalAnchor) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_egde.String(), start)
}
func (_ac *Bordertop) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0062o\u0072\u0064\u0065\u0072\u0074\u006fp"
	return _ac.CT_Border.MarshalXML(e, start)
}
func (_gcf *CT_AnchorLock) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

const (
	ST_BorderTypeUnset                 ST_BorderType = 0
	ST_BorderTypeNone                  ST_BorderType = 1
	ST_BorderTypeSingle                ST_BorderType = 2
	ST_BorderTypeThick                 ST_BorderType = 3
	ST_BorderTypeDouble                ST_BorderType = 4
	ST_BorderTypeHairline              ST_BorderType = 5
	ST_BorderTypeDot                   ST_BorderType = 6
	ST_BorderTypeDash                  ST_BorderType = 7
	ST_BorderTypeDotDash               ST_BorderType = 8
	ST_BorderTypeDashDotDot            ST_BorderType = 9
	ST_BorderTypeTriple                ST_BorderType = 10
	ST_BorderTypeThinThickSmall        ST_BorderType = 11
	ST_BorderTypeThickThinSmall        ST_BorderType = 12
	ST_BorderTypeThickBetweenThinSmall ST_BorderType = 13
	ST_BorderTypeThinThick             ST_BorderType = 14
	ST_BorderTypeThickThin             ST_BorderType = 15
	ST_BorderTypeThickBetweenThin      ST_BorderType = 16
	ST_BorderTypeThinThickLarge        ST_BorderType = 17
	ST_BorderTypeThickThinLarge        ST_BorderType = 18
	ST_BorderTypeThickBetweenThinLarge ST_BorderType = 19
	ST_BorderTypeWave                  ST_BorderType = 20
	ST_BorderTypeDoubleWave            ST_BorderType = 21
	ST_BorderTypeDashedSmall           ST_BorderType = 22
	ST_BorderTypeDashDotStroked        ST_BorderType = 23
	ST_BorderTypeThreeDEmboss          ST_BorderType = 24
	ST_BorderTypeThreeDEngrave         ST_BorderType = 25
	ST_BorderTypeHTMLOutset            ST_BorderType = 26
	ST_BorderTypeHTMLInset             ST_BorderType = 27
)

// ValidateWithPath validates the CT_Wrap and its children, prefixing error messages with path
func (_deff *CT_Wrap) ValidateWithPath(path string) error {
	if _fgc := _deff.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _fgc != nil {
		return _fgc
	}
	if _gff := _deff.SideAttr.ValidateWithPath(path + "\u002fS\u0069\u0064\u0065\u0041\u0074\u0074r"); _gff != nil {
		return _gff
	}
	if _gfc := _deff.AnchorxAttr.ValidateWithPath(path + "\u002f\u0041\u006ec\u0068\u006f\u0072\u0078\u0041\u0074\u0074\u0072"); _gfc != nil {
		return _gfc
	}
	if _dab := _deff.AnchoryAttr.ValidateWithPath(path + "\u002f\u0041\u006ec\u0068\u006f\u0072\u0079\u0041\u0074\u0074\u0072"); _dab != nil {
		return _dab
	}
	return nil
}

type Borderbottom struct{ CT_Border }

func (_ddg *CT_Border) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _acc := range start.Attr {
		if _acc.Name.Local == "\u0074\u0079\u0070\u0065" {
			_ddg.TypeAttr.UnmarshalXMLAttr(_acc)
			continue
		}
		if _acc.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_edba, _gfd := _c.ParseUint(_acc.Value, 10, 32)
			if _gfd != nil {
				return _gfd
			}
			_edbg := uint32(_edba)
			_ddg.WidthAttr = &_edbg
			continue
		}
		if _acc.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_ddg.ShadowAttr.UnmarshalXMLAttr(_acc)
			continue
		}
	}
	for {
		_dfc, _bdf := d.Token()
		if _bdf != nil {
			return _ce.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0042\u006f\u0072d\u0065\u0072\u003a\u0020\u0025\u0073", _bdf)
		}
		if _bbf, _gd := _dfc.(_d.EndElement); _gd && _bbf.Name == start.Name {
			break
		}
	}
	return nil
}
func (_g *Anchorlock) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_g.CT_AnchorLock = *NewCT_AnchorLock()
	for {
		_fe, _bf := d.Token()
		if _bf != nil {
			return _ce.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0041\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b\u003a\u0020%\u0073", _bf)
		}
		if _e, _eb := _fe.(_d.EndElement); _eb && _e.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Borderleft and its children
func (_eec *Borderleft) Validate() error {
	return _eec.ValidateWithPath("\u0042\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074")
}
func (_dgf ST_BorderType) ValidateWithPath(path string) error {
	switch _dgf {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27:
	default:
		return _ce.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dgf))
	}
	return nil
}

type Bordertop struct{ CT_Border }

func (_ddd ST_BorderShadow) String() string {
	switch _ddd {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0074\u0072\u0075\u0065"
	case 3:
		return "\u0066"
	case 4:
		return "\u0066\u0061\u006cs\u0065"
	}
	return ""
}

// Validate validates the CT_AnchorLock and its children
func (_daa *CT_AnchorLock) Validate() error {
	return _daa.ValidateWithPath("\u0043\u0054\u005f\u0041\u006e\u0063\u0068\u006f\u0072\u004c\u006f\u0063\u006b")
}
func (_ffee ST_HorizontalAnchor) String() string {
	switch _ffee {
	case 0:
		return ""
	case 1:
		return "\u006d\u0061\u0072\u0067\u0069\u006e"
	case 2:
		return "\u0070\u0061\u0067\u0065"
	case 3:
		return "\u0074\u0065\u0078\u0074"
	case 4:
		return "\u0063\u0068\u0061\u0072"
	}
	return ""
}

// ValidateWithPath validates the Borderright and its children, prefixing error messages with path
func (_gee *Borderright) ValidateWithPath(path string) error {
	if _fed := _gee.CT_Border.ValidateWithPath(path); _fed != nil {
		return _fed
	}
	return nil
}
func (_ffbf ST_HorizontalAnchor) Validate() error { return _ffbf.ValidateWithPath("") }

type ST_HorizontalAnchor byte

func (_fad ST_HorizontalAnchor) ValidateWithPath(path string) error {
	switch _fad {
	case 0, 1, 2, 3, 4:
	default:
		return _ce.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fad))
	}
	return nil
}

type CT_Wrap struct {
	TypeAttr    ST_WrapType
	SideAttr    ST_WrapSide
	AnchorxAttr ST_HorizontalAnchor
	AnchoryAttr ST_VerticalAnchor
}
type Borderleft struct{ CT_Border }

func (_fgce ST_WrapType) String() string {
	switch _fgce {
	case 0:
		return ""
	case 1:
		return "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d"
	case 2:
		return "\u0073\u0071\u0075\u0061\u0072\u0065"
	case 3:
		return "\u006e\u006f\u006e\u0065"
	case 4:
		return "\u0074\u0069\u0067h\u0074"
	case 5:
		return "\u0074h\u0072\u006f\u0075\u0067\u0068"
	}
	return ""
}

// Validate validates the Bordertop and its children
func (_fgg *Bordertop) Validate() error {
	return _fgg.ValidateWithPath("\u0042o\u0072\u0064\u0065\u0072\u0074\u006fp")
}
func NewWrap() *Wrap { _bbfa := &Wrap{}; _bbfa.CT_Wrap = *NewCT_Wrap(); return _bbfa }

// Validate validates the Borderright and its children
func (_be *Borderright) Validate() error {
	return _be.ValidateWithPath("B\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074")
}
func (_cdf ST_BorderShadow) Validate() error { return _cdf.ValidateWithPath("") }
func (_ffdf *ST_WrapType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gbf, _fbb := d.Token()
	if _fbb != nil {
		return _fbb
	}
	if _bbfe, _ffb := _gbf.(_d.EndElement); _ffb && _bbfe.Name == start.Name {
		*_ffdf = 1
		return nil
	}
	if _gcg, _bgc := _gbf.(_d.CharData); !_bgc {
		return _ce.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbf)
	} else {
		switch string(_gcg) {
		case "":
			*_ffdf = 0
		case "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d":
			*_ffdf = 1
		case "\u0073\u0071\u0075\u0061\u0072\u0065":
			*_ffdf = 2
		case "\u006e\u006f\u006e\u0065":
			*_ffdf = 3
		case "\u0074\u0069\u0067h\u0074":
			*_ffdf = 4
		case "\u0074h\u0072\u006f\u0075\u0067\u0068":
			*_ffdf = 5
		}
	}
	_gbf, _fbb = d.Token()
	if _fbb != nil {
		return _fbb
	}
	if _ced, _fcg := _gbf.(_d.EndElement); _fcg && _ced.Name == start.Name {
		return nil
	}
	return _ce.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbf)
}
func (_gfe *ST_WrapSide) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dadg, _eba := d.Token()
	if _eba != nil {
		return _eba
	}
	if _bbbd, _gea := _dadg.(_d.EndElement); _gea && _bbbd.Name == start.Name {
		*_gfe = 1
		return nil
	}
	if _bdg, _abf := _dadg.(_d.CharData); !_abf {
		return _ce.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dadg)
	} else {
		switch string(_bdg) {
		case "":
			*_gfe = 0
		case "\u0062\u006f\u0074\u0068":
			*_gfe = 1
		case "\u006c\u0065\u0066\u0074":
			*_gfe = 2
		case "\u0072\u0069\u0067h\u0074":
			*_gfe = 3
		case "\u006ca\u0072\u0067\u0065\u0073\u0074":
			*_gfe = 4
		}
	}
	_dadg, _eba = d.Token()
	if _eba != nil {
		return _eba
	}
	if _gfbd, _dge := _dadg.(_d.EndElement); _dge && _gfbd.Name == start.Name {
		return nil
	}
	return _ce.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dadg)
}

// ValidateWithPath validates the Borderbottom and its children, prefixing error messages with path
func (_ef *Borderbottom) ValidateWithPath(path string) error {
	if _efb := _ef.CT_Border.ValidateWithPath(path); _efb != nil {
		return _efb
	}
	return nil
}
func (_cef *ST_VerticalAnchor) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dda, _fac := d.Token()
	if _fac != nil {
		return _fac
	}
	if _bge, _caf := _dda.(_d.EndElement); _caf && _bge.Name == start.Name {
		*_cef = 1
		return nil
	}
	if _fefg, _eee := _dda.(_d.CharData); !_eee {
		return _ce.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dda)
	} else {
		switch string(_fefg) {
		case "":
			*_cef = 0
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			*_cef = 1
		case "\u0070\u0061\u0067\u0065":
			*_cef = 2
		case "\u0074\u0065\u0078\u0074":
			*_cef = 3
		case "\u006c\u0069\u006e\u0065":
			*_cef = 4
		}
	}
	_dda, _fac = d.Token()
	if _fac != nil {
		return _fac
	}
	if _dde, _fcfd := _dda.(_d.EndElement); _fcfd && _dde.Name == start.Name {
		return nil
	}
	return _ce.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dda)
}
func (_feg *CT_Wrap) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _feg.TypeAttr != ST_WrapTypeUnset {
		_ddbc, _ceg := _feg.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _ceg != nil {
			return _ceg
		}
		start.Attr = append(start.Attr, _ddbc)
	}
	if _feg.SideAttr != ST_WrapSideUnset {
		_fcb, _ec := _feg.SideAttr.MarshalXMLAttr(_d.Name{Local: "\u0073\u0069\u0064\u0065"})
		if _ec != nil {
			return _ec
		}
		start.Attr = append(start.Attr, _fcb)
	}
	if _feg.AnchorxAttr != ST_HorizontalAnchorUnset {
		_gfde, _cbc := _feg.AnchorxAttr.MarshalXMLAttr(_d.Name{Local: "\u0061n\u0063\u0068\u006f\u0072\u0078"})
		if _cbc != nil {
			return _cbc
		}
		start.Attr = append(start.Attr, _gfde)
	}
	if _feg.AnchoryAttr != ST_VerticalAnchorUnset {
		_bbg, _cf := _feg.AnchoryAttr.MarshalXMLAttr(_d.Name{Local: "\u0061n\u0063\u0068\u006f\u0072\u0079"})
		if _cf != nil {
			return _cf
		}
		start.Attr = append(start.Attr, _bbg)
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

type Anchorlock struct{ CT_AnchorLock }

// Validate validates the CT_Border and its children
func (_dff *CT_Border) Validate() error {
	return _dff.ValidateWithPath("\u0043T\u005f\u0042\u006f\u0072\u0064\u0065r")
}
func NewCT_Border() *CT_Border { _dee := &CT_Border{}; return _dee }

type CT_Border struct {
	TypeAttr   ST_BorderType
	WidthAttr  *uint32
	ShadowAttr ST_BorderShadow
}

func (_beb ST_VerticalAnchor) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dce := _d.Attr{}
	_dce.Name = name
	switch _beb {
	case ST_VerticalAnchorUnset:
		_dce.Value = ""
	case ST_VerticalAnchorMargin:
		_dce.Value = "\u006d\u0061\u0072\u0067\u0069\u006e"
	case ST_VerticalAnchorPage:
		_dce.Value = "\u0070\u0061\u0067\u0065"
	case ST_VerticalAnchorText:
		_dce.Value = "\u0074\u0065\u0078\u0074"
	case ST_VerticalAnchorLine:
		_dce.Value = "\u006c\u0069\u006e\u0065"
	}
	return _dce, nil
}
func (_dcb *ST_HorizontalAnchor) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cgb, _ba := d.Token()
	if _ba != nil {
		return _ba
	}
	if _afb, _debf := _cgb.(_d.EndElement); _debf && _afb.Name == start.Name {
		*_dcb = 1
		return nil
	}
	if _eeb, _aaf := _cgb.(_d.CharData); !_aaf {
		return _ce.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cgb)
	} else {
		switch string(_eeb) {
		case "":
			*_dcb = 0
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			*_dcb = 1
		case "\u0070\u0061\u0067\u0065":
			*_dcb = 2
		case "\u0074\u0065\u0078\u0074":
			*_dcb = 3
		case "\u0063\u0068\u0061\u0072":
			*_dcb = 4
		}
	}
	_cgb, _ba = d.Token()
	if _ba != nil {
		return _ba
	}
	if _daae, _bcc := _cgb.(_d.EndElement); _bcc && _daae.Name == start.Name {
		return nil
	}
	return _ce.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cgb)
}

const (
	ST_HorizontalAnchorUnset  ST_HorizontalAnchor = 0
	ST_HorizontalAnchorMargin ST_HorizontalAnchor = 1
	ST_HorizontalAnchorPage   ST_HorizontalAnchor = 2
	ST_HorizontalAnchorText   ST_HorizontalAnchor = 3
	ST_HorizontalAnchorChar   ST_HorizontalAnchor = 4
)

func (_bg *ST_BorderShadow) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bg = 0
	case "\u0074":
		*_bg = 1
	case "\u0074\u0072\u0075\u0065":
		*_bg = 2
	case "\u0066":
		*_bg = 3
	case "\u0066\u0061\u006cs\u0065":
		*_bg = 4
	}
	return nil
}
func (_gc *Borderleft) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gc.CT_Border = *NewCT_Border()
	for _, _dd := range start.Attr {
		if _dd.Name.Local == "\u0074\u0079\u0070\u0065" {
			_gc.TypeAttr.UnmarshalXMLAttr(_dd)
			continue
		}
		if _dd.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_dc, _bdd := _c.ParseUint(_dd.Value, 10, 32)
			if _bdd != nil {
				return _bdd
			}
			_ff := uint32(_dc)
			_gc.WidthAttr = &_ff
			continue
		}
		if _dd.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_gc.ShadowAttr.UnmarshalXMLAttr(_dd)
			continue
		}
	}
	for {
		_ebc, _ae := d.Token()
		if _ae != nil {
			return _ce.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0042\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074\u003a\u0020%\u0073", _ae)
		}
		if _ca, _ee := _ebc.(_d.EndElement); _ee && _ca.Name == start.Name {
			break
		}
	}
	return nil
}
func (_add *ST_HorizontalAnchor) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_add = 0
	case "\u006d\u0061\u0072\u0067\u0069\u006e":
		*_add = 1
	case "\u0070\u0061\u0067\u0065":
		*_add = 2
	case "\u0074\u0065\u0078\u0074":
		*_add = 3
	case "\u0063\u0068\u0061\u0072":
		*_add = 4
	}
	return nil
}
func (_ede ST_BorderType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_cd := _d.Attr{}
	_cd.Name = name
	switch _ede {
	case ST_BorderTypeUnset:
		_cd.Value = ""
	case ST_BorderTypeNone:
		_cd.Value = "\u006e\u006f\u006e\u0065"
	case ST_BorderTypeSingle:
		_cd.Value = "\u0073\u0069\u006e\u0067\u006c\u0065"
	case ST_BorderTypeThick:
		_cd.Value = "\u0074\u0068\u0069c\u006b"
	case ST_BorderTypeDouble:
		_cd.Value = "\u0064\u006f\u0075\u0062\u006c\u0065"
	case ST_BorderTypeHairline:
		_cd.Value = "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065"
	case ST_BorderTypeDot:
		_cd.Value = "\u0064\u006f\u0074"
	case ST_BorderTypeDash:
		_cd.Value = "\u0064\u0061\u0073\u0068"
	case ST_BorderTypeDotDash:
		_cd.Value = "\u0064o\u0074\u0044\u0061\u0073\u0068"
	case ST_BorderTypeDashDotDot:
		_cd.Value = "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074"
	case ST_BorderTypeTriple:
		_cd.Value = "\u0074\u0072\u0069\u0070\u006c\u0065"
	case ST_BorderTypeThinThickSmall:
		_cd.Value = "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c"
	case ST_BorderTypeThickThinSmall:
		_cd.Value = "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c"
	case ST_BorderTypeThickBetweenThinSmall:
		_cd.Value = "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c"
	case ST_BorderTypeThinThick:
		_cd.Value = "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k"
	case ST_BorderTypeThickThin:
		_cd.Value = "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n"
	case ST_BorderTypeThickBetweenThin:
		_cd.Value = "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e"
	case ST_BorderTypeThinThickLarge:
		_cd.Value = "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065"
	case ST_BorderTypeThickThinLarge:
		_cd.Value = "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065"
	case ST_BorderTypeThickBetweenThinLarge:
		_cd.Value = "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065"
	case ST_BorderTypeWave:
		_cd.Value = "\u0077\u0061\u0076\u0065"
	case ST_BorderTypeDoubleWave:
		_cd.Value = "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065"
	case ST_BorderTypeDashedSmall:
		_cd.Value = "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c"
	case ST_BorderTypeDashDotStroked:
		_cd.Value = "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064"
	case ST_BorderTypeThreeDEmboss:
		_cd.Value = "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073"
	case ST_BorderTypeThreeDEngrave:
		_cd.Value = "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065"
	case ST_BorderTypeHTMLOutset:
		_cd.Value = "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074"
	case ST_BorderTypeHTMLInset:
		_cd.Value = "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t"
	}
	return _cd, nil
}
func (_cda ST_VerticalAnchor) Validate() error { return _cda.ValidateWithPath("") }
func NewCT_AnchorLock() *CT_AnchorLock         { _ebd := &CT_AnchorLock{}; return _ebd }

const (
	ST_BorderShadowUnset ST_BorderShadow = 0
	ST_BorderShadowT     ST_BorderShadow = 1
	ST_BorderShadowTrue  ST_BorderShadow = 2
	ST_BorderShadowF     ST_BorderShadow = 3
	ST_BorderShadowFalse ST_BorderShadow = 4
)

// ValidateWithPath validates the Bordertop and its children, prefixing error messages with path
func (_gg *Bordertop) ValidateWithPath(path string) error {
	if _cab := _gg.CT_Border.ValidateWithPath(path); _cab != nil {
		return _cab
	}
	return nil
}

const (
	ST_WrapTypeUnset        ST_WrapType = 0
	ST_WrapTypeTopAndBottom ST_WrapType = 1
	ST_WrapTypeSquare       ST_WrapType = 2
	ST_WrapTypeNone         ST_WrapType = 3
	ST_WrapTypeTight        ST_WrapType = 4
	ST_WrapTypeThrough      ST_WrapType = 5
)

func (_fgca ST_WrapSide) ValidateWithPath(path string) error {
	switch _fgca {
	case 0, 1, 2, 3, 4:
	default:
		return _ce.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fgca))
	}
	return nil
}

// Validate validates the Anchorlock and its children
func (_a *Anchorlock) Validate() error {
	return _a.ValidateWithPath("\u0041\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b")
}
func (_ead ST_WrapSide) Validate() error { return _ead.ValidateWithPath("") }
func (_fec *Bordertop) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fec.CT_Border = *NewCT_Border()
	for _, _gaa := range start.Attr {
		if _gaa.Name.Local == "\u0074\u0079\u0070\u0065" {
			_fec.TypeAttr.UnmarshalXMLAttr(_gaa)
			continue
		}
		if _gaa.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_dfe, _egc := _c.ParseUint(_gaa.Value, 10, 32)
			if _egc != nil {
				return _egc
			}
			_aae := uint32(_dfe)
			_fec.WidthAttr = &_aae
			continue
		}
		if _gaa.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_fec.ShadowAttr.UnmarshalXMLAttr(_gaa)
			continue
		}
	}
	for {
		_fga, _gcc := d.Token()
		if _gcc != nil {
			return _ce.Errorf("p\u0061\u0072\u0073\u0069ng\u0020B\u006f\u0072\u0064\u0065\u0072t\u006f\u0070\u003a\u0020\u0025\u0073", _gcc)
		}
		if _ffe, _ddb := _fga.(_d.EndElement); _ddb && _ffe.Name == start.Name {
			break
		}
	}
	return nil
}
func (_cfc ST_VerticalAnchor) ValidateWithPath(path string) error {
	switch _cfc {
	case 0, 1, 2, 3, 4:
	default:
		return _ce.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cfc))
	}
	return nil
}
func (_age ST_WrapType) Validate() error { return _age.ValidateWithPath("") }
func (_dcg ST_WrapSide) String() string {
	switch _dcg {
	case 0:
		return ""
	case 1:
		return "\u0062\u006f\u0074\u0068"
	case 2:
		return "\u006c\u0065\u0066\u0074"
	case 3:
		return "\u0072\u0069\u0067h\u0074"
	case 4:
		return "\u006ca\u0072\u0067\u0065\u0073\u0074"
	}
	return ""
}
func NewBorderbottom() *Borderbottom {
	_aa := &Borderbottom{}
	_aa.CT_Border = *NewCT_Border()
	return _aa
}

const (
	ST_WrapSideUnset   ST_WrapSide = 0
	ST_WrapSideBoth    ST_WrapSide = 1
	ST_WrapSideLeft    ST_WrapSide = 2
	ST_WrapSideRight   ST_WrapSide = 3
	ST_WrapSideLargest ST_WrapSide = 4
)

type ST_WrapSide byte

// ValidateWithPath validates the Anchorlock and its children, prefixing error messages with path
func (_fa *Anchorlock) ValidateWithPath(path string) error {
	if _fc := _fa.CT_AnchorLock.ValidateWithPath(path); _fc != nil {
		return _fc
	}
	return nil
}
func (_faa ST_BorderType) Validate() error { return _faa.ValidateWithPath("") }
func (_fd ST_BorderShadow) ValidateWithPath(path string) error {
	switch _fd {
	case 0, 1, 2, 3, 4:
	default:
		return _ce.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fd))
	}
	return nil
}
func (_adge *ST_BorderType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_abb, _bbb := d.Token()
	if _bbb != nil {
		return _bbb
	}
	if _ffd, _gae := _abb.(_d.EndElement); _gae && _ffd.Name == start.Name {
		*_adge = 1
		return nil
	}
	if _gcfa, _gbg := _abb.(_d.CharData); !_gbg {
		return _ce.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _abb)
	} else {
		switch string(_gcfa) {
		case "":
			*_adge = 0
		case "\u006e\u006f\u006e\u0065":
			*_adge = 1
		case "\u0073\u0069\u006e\u0067\u006c\u0065":
			*_adge = 2
		case "\u0074\u0068\u0069c\u006b":
			*_adge = 3
		case "\u0064\u006f\u0075\u0062\u006c\u0065":
			*_adge = 4
		case "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065":
			*_adge = 5
		case "\u0064\u006f\u0074":
			*_adge = 6
		case "\u0064\u0061\u0073\u0068":
			*_adge = 7
		case "\u0064o\u0074\u0044\u0061\u0073\u0068":
			*_adge = 8
		case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074":
			*_adge = 9
		case "\u0074\u0072\u0069\u0070\u006c\u0065":
			*_adge = 10
		case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c":
			*_adge = 11
		case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c":
			*_adge = 12
		case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c":
			*_adge = 13
		case "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k":
			*_adge = 14
		case "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n":
			*_adge = 15
		case "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e":
			*_adge = 16
		case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065":
			*_adge = 17
		case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065":
			*_adge = 18
		case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065":
			*_adge = 19
		case "\u0077\u0061\u0076\u0065":
			*_adge = 20
		case "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065":
			*_adge = 21
		case "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c":
			*_adge = 22
		case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064":
			*_adge = 23
		case "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073":
			*_adge = 24
		case "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065":
			*_adge = 25
		case "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074":
			*_adge = 26
		case "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t":
			*_adge = 27
		}
	}
	_abb, _bbb = d.Token()
	if _bbb != nil {
		return _bbb
	}
	if _aee, _gdeg := _abb.(_d.EndElement); _gdeg && _aee.Name == start.Name {
		return nil
	}
	return _ce.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _abb)
}
func (_fdd *ST_WrapSide) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fdd = 0
	case "\u0062\u006f\u0074\u0068":
		*_fdd = 1
	case "\u006c\u0065\u0066\u0074":
		*_fdd = 2
	case "\u0072\u0069\u0067h\u0074":
		*_fdd = 3
	case "\u006ca\u0072\u0067\u0065\u0073\u0074":
		*_fdd = 4
	}
	return nil
}

type ST_BorderShadow byte

func NewBordertop() *Bordertop { _ege := &Bordertop{}; _ege.CT_Border = *NewCT_Border(); return _ege }

// ValidateWithPath validates the CT_AnchorLock and its children, prefixing error messages with path
func (_ddc *CT_AnchorLock) ValidateWithPath(path string) error { return nil }

// Validate validates the Wrap and its children
func (_bfg *Wrap) Validate() error { return _bfg.ValidateWithPath("\u0057\u0072\u0061\u0070") }
func (_gfdc ST_WrapType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_gfdc.String(), start)
}
func (_deg ST_WrapSide) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_aaa := _d.Attr{}
	_aaa.Name = name
	switch _deg {
	case ST_WrapSideUnset:
		_aaa.Value = ""
	case ST_WrapSideBoth:
		_aaa.Value = "\u0062\u006f\u0074\u0068"
	case ST_WrapSideLeft:
		_aaa.Value = "\u006c\u0065\u0066\u0074"
	case ST_WrapSideRight:
		_aaa.Value = "\u0072\u0069\u0067h\u0074"
	case ST_WrapSideLargest:
		_aaa.Value = "\u006ca\u0072\u0067\u0065\u0073\u0074"
	}
	return _aaa, nil
}
func (_afc *ST_BorderType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_afc = 0
	case "\u006e\u006f\u006e\u0065":
		*_afc = 1
	case "\u0073\u0069\u006e\u0067\u006c\u0065":
		*_afc = 2
	case "\u0074\u0068\u0069c\u006b":
		*_afc = 3
	case "\u0064\u006f\u0075\u0062\u006c\u0065":
		*_afc = 4
	case "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065":
		*_afc = 5
	case "\u0064\u006f\u0074":
		*_afc = 6
	case "\u0064\u0061\u0073\u0068":
		*_afc = 7
	case "\u0064o\u0074\u0044\u0061\u0073\u0068":
		*_afc = 8
	case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074":
		*_afc = 9
	case "\u0074\u0072\u0069\u0070\u006c\u0065":
		*_afc = 10
	case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c":
		*_afc = 11
	case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c":
		*_afc = 12
	case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c":
		*_afc = 13
	case "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k":
		*_afc = 14
	case "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n":
		*_afc = 15
	case "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e":
		*_afc = 16
	case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065":
		*_afc = 17
	case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065":
		*_afc = 18
	case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065":
		*_afc = 19
	case "\u0077\u0061\u0076\u0065":
		*_afc = 20
	case "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065":
		*_afc = 21
	case "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c":
		*_afc = 22
	case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064":
		*_afc = 23
	case "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073":
		*_afc = 24
	case "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065":
		*_afc = 25
	case "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074":
		*_afc = 26
	case "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t":
		*_afc = 27
	}
	return nil
}

type Borderright struct{ CT_Border }

func (_aac ST_WrapType) ValidateWithPath(path string) error {
	switch _aac {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _ce.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aac))
	}
	return nil
}

// ValidateWithPath validates the Wrap and its children, prefixing error messages with path
func (_gac *Wrap) ValidateWithPath(path string) error {
	if _ddga := _gac.CT_Wrap.ValidateWithPath(path); _ddga != nil {
		return _ddga
	}
	return nil
}
func (_ebgd *Borderright) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ebgd.CT_Border = *NewCT_Border()
	for _, _bda := range start.Attr {
		if _bda.Name.Local == "\u0074\u0079\u0070\u0065" {
			_ebgd.TypeAttr.UnmarshalXMLAttr(_bda)
			continue
		}
		if _bda.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_fg, _edb := _c.ParseUint(_bda.Value, 10, 32)
			if _edb != nil {
				return _edb
			}
			_ega := uint32(_fg)
			_ebgd.WidthAttr = &_ega
			continue
		}
		if _bda.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_ebgd.ShadowAttr.UnmarshalXMLAttr(_bda)
			continue
		}
	}
	for {
		_dgg, _ea := d.Token()
		if _ea != nil {
			return _ce.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0042\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074\u003a\u0020\u0025\u0073", _ea)
		}
		if _db, _aeb := _dgg.(_d.EndElement); _aeb && _db.Name == start.Name {
			break
		}
	}
	return nil
}
func (_efa ST_HorizontalAnchor) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_afae := _d.Attr{}
	_afae.Name = name
	switch _efa {
	case ST_HorizontalAnchorUnset:
		_afae.Value = ""
	case ST_HorizontalAnchorMargin:
		_afae.Value = "\u006d\u0061\u0072\u0067\u0069\u006e"
	case ST_HorizontalAnchorPage:
		_afae.Value = "\u0070\u0061\u0067\u0065"
	case ST_HorizontalAnchorText:
		_afae.Value = "\u0074\u0065\u0078\u0074"
	case ST_HorizontalAnchorChar:
		_afae.Value = "\u0063\u0068\u0061\u0072"
	}
	return _afae, nil
}
func (_bdac ST_WrapType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gceg := _d.Attr{}
	_gceg.Name = name
	switch _bdac {
	case ST_WrapTypeUnset:
		_gceg.Value = ""
	case ST_WrapTypeTopAndBottom:
		_gceg.Value = "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d"
	case ST_WrapTypeSquare:
		_gceg.Value = "\u0073\u0071\u0075\u0061\u0072\u0065"
	case ST_WrapTypeNone:
		_gceg.Value = "\u006e\u006f\u006e\u0065"
	case ST_WrapTypeTight:
		_gceg.Value = "\u0074\u0069\u0067h\u0074"
	case ST_WrapTypeThrough:
		_gceg.Value = "\u0074h\u0072\u006f\u0075\u0067\u0068"
	}
	return _gceg, nil
}

// Validate validates the Borderbottom and its children
func (_ebg *Borderbottom) Validate() error {
	return _ebg.ValidateWithPath("\u0042\u006f\u0072d\u0065\u0072\u0062\u006f\u0074\u0074\u006f\u006d")
}
func (_fcac ST_WrapSide) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fcac.String(), start)
}
func NewCT_Wrap() *CT_Wrap { _cb := &CT_Wrap{}; return _cb }

const (
	ST_VerticalAnchorUnset  ST_VerticalAnchor = 0
	ST_VerticalAnchorMargin ST_VerticalAnchor = 1
	ST_VerticalAnchorPage   ST_VerticalAnchor = 2
	ST_VerticalAnchorText   ST_VerticalAnchor = 3
	ST_VerticalAnchorLine   ST_VerticalAnchor = 4
)

func (_eca *ST_WrapType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_eca = 0
	case "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d":
		*_eca = 1
	case "\u0073\u0071\u0075\u0061\u0072\u0065":
		*_eca = 2
	case "\u006e\u006f\u006e\u0065":
		*_eca = 3
	case "\u0074\u0069\u0067h\u0074":
		*_eca = 4
	case "\u0074h\u0072\u006f\u0075\u0067\u0068":
		*_eca = 5
	}
	return nil
}

// Validate validates the CT_Wrap and its children
func (_gccd *CT_Wrap) Validate() error {
	return _gccd.ValidateWithPath("\u0043T\u005f\u0057\u0072\u0061\u0070")
}
func NewBorderleft() *Borderleft {
	_efbe := &Borderleft{}
	_efbe.CT_Border = *NewCT_Border()
	return _efbe
}
func (_ag ST_BorderShadow) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_afa := _d.Attr{}
	_afa.Name = name
	switch _ag {
	case ST_BorderShadowUnset:
		_afa.Value = ""
	case ST_BorderShadowT:
		_afa.Value = "\u0074"
	case ST_BorderShadowTrue:
		_afa.Value = "\u0074\u0072\u0075\u0065"
	case ST_BorderShadowF:
		_afa.Value = "\u0066"
	case ST_BorderShadowFalse:
		_afa.Value = "\u0066\u0061\u006cs\u0065"
	}
	return _afa, nil
}
func (_dbc *CT_AnchorLock) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for {
		_fab, _bb := d.Token()
		if _bb != nil {
			return _ce.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0041n\u0063h\u006fr\u004c\u006f\u0063\u006b\u003a\u0020\u0025s", _bb)
		}
		if _fcag, _fee := _fab.(_d.EndElement); _fee && _fcag.Name == start.Name {
			break
		}
	}
	return nil
}
func (_gaae ST_BorderType) String() string {
	switch _gaae {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0073\u0069\u006e\u0067\u006c\u0065"
	case 3:
		return "\u0074\u0068\u0069c\u006b"
	case 4:
		return "\u0064\u006f\u0075\u0062\u006c\u0065"
	case 5:
		return "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065"
	case 6:
		return "\u0064\u006f\u0074"
	case 7:
		return "\u0064\u0061\u0073\u0068"
	case 8:
		return "\u0064o\u0074\u0044\u0061\u0073\u0068"
	case 9:
		return "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074"
	case 10:
		return "\u0074\u0072\u0069\u0070\u006c\u0065"
	case 11:
		return "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c"
	case 12:
		return "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c"
	case 13:
		return "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c"
	case 14:
		return "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k"
	case 15:
		return "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n"
	case 16:
		return "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e"
	case 17:
		return "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065"
	case 18:
		return "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065"
	case 19:
		return "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065"
	case 20:
		return "\u0077\u0061\u0076\u0065"
	case 21:
		return "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065"
	case 22:
		return "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c"
	case 23:
		return "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064"
	case 24:
		return "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073"
	case 25:
		return "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065"
	case 26:
		return "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074"
	case 27:
		return "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t"
	}
	return ""
}
func (_ga *Borderbottom) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ga.CT_Border = *NewCT_Border()
	for _, _gf := range start.Attr {
		if _gf.Name.Local == "\u0074\u0079\u0070\u0065" {
			_ga.TypeAttr.UnmarshalXMLAttr(_gf)
			continue
		}
		if _gf.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_ge, _fca := _c.ParseUint(_gf.Value, 10, 32)
			if _fca != nil {
				return _fca
			}
			_da := uint32(_ge)
			_ga.WidthAttr = &_da
			continue
		}
		if _gf.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_ga.ShadowAttr.UnmarshalXMLAttr(_gf)
			continue
		}
	}
	for {
		_fb, _aad := d.Token()
		if _aad != nil {
			return _ce.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0042\u006f\u0072\u0064\u0065r\u0062\u006f\u0074\u0074\u006f\u006d\u003a\u0020\u0025\u0073", _aad)
		}
		if _bc, _gec := _fb.(_d.EndElement); _gec && _bc.Name == start.Name {
			break
		}
	}
	return nil
}
func (_def *Borderbottom) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0062\u006f\u0072d\u0065\u0072\u0062\u006f\u0074\u0074\u006f\u006d"
	return _def.CT_Border.MarshalXML(e, start)
}

type ST_VerticalAnchor byte
type ST_BorderType byte

func (_cga *CT_Border) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _cga.TypeAttr != ST_BorderTypeUnset {
		_ggc, _dbcb := _cga.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _dbcb != nil {
			return _dbcb
		}
		start.Attr = append(start.Attr, _ggc)
	}
	if _cga.WidthAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0077\u0069\u0064t\u0068"}, Value: _ce.Sprintf("\u0025\u0076", *_cga.WidthAttr)})
	}
	if _cga.ShadowAttr != ST_BorderShadowUnset {
		_gb, _deb := _cga.ShadowAttr.MarshalXMLAttr(_d.Name{Local: "\u0073\u0068\u0061\u0064\u006f\u0077"})
		if _deb != nil {
			return _deb
		}
		start.Attr = append(start.Attr, _gb)
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewBorderright() *Borderright {
	_caa := &Borderright{}
	_caa.CT_Border = *NewCT_Border()
	return _caa
}
func (_febe ST_VerticalAnchor) String() string {
	switch _febe {
	case 0:
		return ""
	case 1:
		return "\u006d\u0061\u0072\u0067\u0069\u006e"
	case 2:
		return "\u0070\u0061\u0067\u0065"
	case 3:
		return "\u0074\u0065\u0078\u0074"
	case 4:
		return "\u006c\u0069\u006e\u0065"
	}
	return ""
}
func (_bfe *CT_Wrap) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _efc := range start.Attr {
		if _efc.Name.Local == "\u0074\u0079\u0070\u0065" {
			_bfe.TypeAttr.UnmarshalXMLAttr(_efc)
			continue
		}
		if _efc.Name.Local == "\u0073\u0069\u0064\u0065" {
			_bfe.SideAttr.UnmarshalXMLAttr(_efc)
			continue
		}
		if _efc.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0078" {
			_bfe.AnchorxAttr.UnmarshalXMLAttr(_efc)
			continue
		}
		if _efc.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0079" {
			_bfe.AnchoryAttr.UnmarshalXMLAttr(_efc)
			continue
		}
	}
	for {
		_adg, _ece := d.Token()
		if _ece != nil {
			return _ce.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0057\u0072\u0061\u0070\u003a\u0020\u0025\u0073", _ece)
		}
		if _egd, _fgga := _adg.(_d.EndElement); _fgga && _egd.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ccc ST_BorderType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ccc.String(), start)
}
func (_fcf *Wrap) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0077\u0072\u0061\u0070"
	return _fcf.CT_Wrap.MarshalXML(e, start)
}

type CT_AnchorLock struct{}

func (_ed *Borderleft) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0062\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074"
	return _ed.CT_Border.MarshalXML(e, start)
}
func (_feca *Wrap) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_feca.CT_Wrap = *NewCT_Wrap()
	for _, _cgad := range start.Attr {
		if _cgad.Name.Local == "\u0074\u0079\u0070\u0065" {
			_feca.TypeAttr.UnmarshalXMLAttr(_cgad)
			continue
		}
		if _cgad.Name.Local == "\u0073\u0069\u0064\u0065" {
			_feca.SideAttr.UnmarshalXMLAttr(_cgad)
			continue
		}
		if _cgad.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0078" {
			_feca.AnchorxAttr.UnmarshalXMLAttr(_cgad)
			continue
		}
		if _cgad.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0079" {
			_feca.AnchoryAttr.UnmarshalXMLAttr(_cgad)
			continue
		}
	}
	for {
		_af, _cc := d.Token()
		if _cc != nil {
			return _ce.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u0057\u0072\u0061\u0070\u003a\u0020\u0025\u0073", _cc)
		}
		if _egf, _feb := _af.(_d.EndElement); _feb && _egf.Name == start.Name {
			break
		}
	}
	return nil
}
func NewAnchorlock() *Anchorlock {
	_b := &Anchorlock{}
	_b.CT_AnchorLock = *NewCT_AnchorLock()
	return _b
}
func (_fge ST_VerticalAnchor) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fge.String(), start)
}

type ST_WrapType byte

// ValidateWithPath validates the CT_Border and its children, prefixing error messages with path
func (_gfb *CT_Border) ValidateWithPath(path string) error {
	if _gde := _gfb.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _gde != nil {
		return _gde
	}
	if _aca := _gfb.ShadowAttr.ValidateWithPath(path + "/\u0053\u0068\u0061\u0064\u006f\u0077\u0041\u0074\u0074\u0072"); _aca != nil {
		return _aca
	}
	return nil
}
func (_fbd *ST_VerticalAnchor) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fbd = 0
	case "\u006d\u0061\u0072\u0067\u0069\u006e":
		*_fbd = 1
	case "\u0070\u0061\u0067\u0065":
		*_fbd = 2
	case "\u0074\u0065\u0078\u0074":
		*_fbd = 3
	case "\u006c\u0069\u006e\u0065":
		*_fbd = 4
	}
	return nil
}
func (_efbg ST_BorderShadow) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_efbg.String(), start)
}
func (_dfb *ST_BorderShadow) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ddf, _dad := d.Token()
	if _dad != nil {
		return _dad
	}
	if _ccb, _aec := _ddf.(_d.EndElement); _aec && _ccb.Name == start.Name {
		*_dfb = 1
		return nil
	}
	if _bgf, _fgd := _ddf.(_d.CharData); !_fgd {
		return _ce.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ddf)
	} else {
		switch string(_bgf) {
		case "":
			*_dfb = 0
		case "\u0074":
			*_dfb = 1
		case "\u0074\u0072\u0075\u0065":
			*_dfb = 2
		case "\u0066":
			*_dfb = 3
		case "\u0066\u0061\u006cs\u0065":
			*_dfb = 4
		}
	}
	_ddf, _dad = d.Token()
	if _dad != nil {
		return _dad
	}
	if _dba, _aff := _ddf.(_d.EndElement); _aff && _dba.Name == start.Name {
		return nil
	}
	return _ce.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ddf)
}

// ValidateWithPath validates the Borderleft and its children, prefixing error messages with path
func (_eg *Borderleft) ValidateWithPath(path string) error {
	if _ab := _eg.CT_Border.ValidateWithPath(path); _ab != nil {
		return _ab
	}
	return nil
}
func init() {
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0043T\u005f\u0042\u006f\u0072\u0064\u0065r", NewCT_Border)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0043T\u005f\u0057\u0072\u0061\u0070", NewCT_Wrap)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0043\u0054\u005f\u0041\u006e\u0063\u0068\u006f\u0072\u004c\u006f\u0063\u006b", NewCT_AnchorLock)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0062o\u0072\u0064\u0065\u0072\u0074\u006fp", NewBordertop)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0062\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074", NewBorderleft)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "b\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074", NewBorderright)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0062\u006f\u0072d\u0065\u0072\u0062\u006f\u0074\u0074\u006f\u006d", NewBorderbottom)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0077\u0072\u0061\u0070", NewWrap)
	_cg.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0061\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b", NewAnchorlock)
}
