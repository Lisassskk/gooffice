package relationships

import (
	_e "encoding/xml"
	_b "fmt"

	_g "gitee.com/gooffice/gooffice"
)

// ValidateWithPath validates the CT_Relationship and its children, prefixing error messages with path
func (_cgf *CT_Relationship) ValidateWithPath(path string) error {
	if _bae := _cgf.TargetModeAttr.ValidateWithPath(path + "\u002fT\u0061r\u0067\u0065\u0074\u004d\u006f\u0064\u0065\u0041\u0074\u0074\u0072"); _bae != nil {
		return _bae
	}
	return nil
}

// Validate validates the CT_Relationships and its children
func (_ag *CT_Relationships) Validate() error {
	return _ag.ValidateWithPath("\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073")
}

type CT_Relationship struct {
	TargetModeAttr ST_TargetMode
	TargetAttr     string
	TypeAttr       string
	IdAttr         string
	Content        string
}

// ValidateWithPath validates the CT_Relationships and its children, prefixing error messages with path
func (_da *CT_Relationships) ValidateWithPath(path string) error {
	for _cb, _ca := range _da.Relationship {
		if _db := _ca.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002f\u0052el\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u005b\u0025\u0064\u005d", path, _cb)); _db != nil {
			return _db
		}
	}
	return nil
}
func (_ade ST_TargetMode) Validate() error { return _ade.ValidateWithPath("") }
func (_ae *Relationships) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_ae.CT_Relationships = *NewCT_Relationships()
_bag:
	for {
		_ef, _gcc := d.Token()
		if _gcc != nil {
			return _gcc
		}
		switch _abg := _ef.(type) {
		case _e.StartElement:
			switch _abg.Name {
			case _e.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:
				_cad := NewRelationship()
				if _dade := d.DecodeElement(_cad, &_abg); _dade != nil {
					return _dade
				}
				_ae.Relationship = append(_ae.Relationship, _cad)
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0052\u0065\u006c\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073 \u0025\u0076", _abg.Name)
				if _ed := d.Skip(); _ed != nil {
					return _ed
				}
			}
		case _e.EndElement:
			break _bag
		case _e.CharData:
		}
	}
	return nil
}
func (_fb *CT_Relationships) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
_acg:
	for {
		_fd, _fad := d.Token()
		if _fad != nil {
			return _fad
		}
		switch _gf := _fd.(type) {
		case _e.StartElement:
			switch _gf.Name {
			case _e.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:
				_eac := NewRelationship()
				if _ga := d.DecodeElement(_eac, &_gf); _ga != nil {
					return _ga
				}
				_fb.Relationship = append(_fb.Relationship, _eac)
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u0020\u0025v", _gf.Name)
				if _ad := d.Skip(); _ad != nil {
					return _ad
				}
			}
		case _e.EndElement:
			break _acg
		case _e.CharData:
		}
	}
	return nil
}
func (_fcf *ST_TargetMode) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_aea, _ecf := d.Token()
	if _ecf != nil {
		return _ecf
	}
	if _cdb, _fdb := _aea.(_e.EndElement); _fdb && _cdb.Name == start.Name {
		*_fcf = 1
		return nil
	}
	if _cba, _cc := _aea.(_e.CharData); !_cc {
		return _b.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _aea)
	} else {
		switch string(_cba) {
		case "":
			*_fcf = 0
		case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":
			*_fcf = 1
		case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":
			*_fcf = 2
		}
	}
	_aea, _ecf = d.Token()
	if _ecf != nil {
		return _ecf
	}
	if _bda, _efe := _aea.(_e.EndElement); _efe && _bda.Name == start.Name {
		return nil
	}
	return _b.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _aea)
}

const (
	ST_TargetModeUnset    ST_TargetMode = 0
	ST_TargetModeExternal ST_TargetMode = 1
	ST_TargetModeInternal ST_TargetMode = 2
)

type Relationships struct{ CT_Relationships }

func (_bbd *CT_Relationships) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	e.EncodeToken(start)
	if _bbd.Relationship != nil {
		_dcb := _e.StartElement{Name: _e.Name{Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}}
		for _, _ac := range _bbd.Relationship {
			e.EncodeElement(_ac, _dcb)
		}
	}
	e.EncodeToken(_e.EndElement{Name: start.Name})
	return nil
}
func (_cae ST_TargetMode) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_cae.String(), start)
}

// Validate validates the Relationship and its children
func (_ab *Relationship) Validate() error {
	return _ab.ValidateWithPath("\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070")
}
func (_gfc ST_TargetMode) String() string {
	switch _gfc {
	case 0:
		return ""
	case 1:
		return "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c"
	case 2:
		return "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c"
	}
	return ""
}
func NewRelationship() *Relationship {
	_eb := &Relationship{}
	_eb.CT_Relationship = *NewCT_Relationship()
	return _eb
}
func NewCT_Relationship() *CT_Relationship { _c := &CT_Relationship{}; return _c }
func (_fg *CT_Relationship) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	for _, _gb := range start.Attr {
		if _gb.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065" {
			_fg.TargetModeAttr.UnmarshalXMLAttr(_gb)
			continue
		}
		if _gb.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074" {
			_cg, _d := _gb.Value, error(nil)
			if _d != nil {
				return _d
			}
			_fg.TargetAttr = _cg
			continue
		}
		if _gb.Name.Local == "\u0054\u0079\u0070\u0065" {
			_cf, _egb := _gb.Value, error(nil)
			if _egb != nil {
				return _egb
			}
			_fg.TypeAttr = _cf
			continue
		}
		if _gb.Name.Local == "\u0049\u0064" {
			_ce, _dg := _gb.Value, error(nil)
			if _dg != nil {
				return _dg
			}
			_fg.IdAttr = _ce
			continue
		}
	}
	for {
		_ec, _dc := d.Token()
		if _dc != nil {
			return _b.Errorf("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068i\u0070:\u0020\u0025\u0073", _dc)
		}
		if _a, _bd := _ec.(_e.CharData); _bd {
			_fg.Content = string(_a)
		}
		if _ba, _ea := _ec.(_e.EndElement); _ea && _ba.Name == start.Name {
			break
		}
	}
	return nil
}
func NewRelationships() *Relationships {
	_dcba := &Relationships{}
	_dcba.CT_Relationships = *NewCT_Relationships()
	return _dcba
}

type Relationship struct{ CT_Relationship }

func (_efa ST_TargetMode) ValidateWithPath(path string) error {
	switch _efa {
	case 0, 1, 2:
	default:
		return _b.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_efa))
	}
	return nil
}
func (_gga *Relationship) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_gga.CT_Relationship = *NewCT_Relationship()
	for _, _fea := range start.Attr {
		if _fea.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065" {
			_gga.TargetModeAttr.UnmarshalXMLAttr(_fea)
			continue
		}
		if _fea.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074" {
			_dad, _ced := _fea.Value, error(nil)
			if _ced != nil {
				return _ced
			}
			_gga.TargetAttr = _dad
			continue
		}
		if _fea.Name.Local == "\u0054\u0079\u0070\u0065" {
			_gc, _dgb := _fea.Value, error(nil)
			if _dgb != nil {
				return _dgb
			}
			_gga.TypeAttr = _gc
			continue
		}
		if _fea.Name.Local == "\u0049\u0064" {
			_fde, _df := _fea.Value, error(nil)
			if _df != nil {
				return _df
			}
			_gga.IdAttr = _fde
			continue
		}
	}
	for {
		_fdc, _acb := d.Token()
		if _acb != nil {
			return _b.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0052\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069\u0070\u003a\u0020\u0025\u0073", _acb)
		}
		if _fc, _dab := _fdc.(_e.EndElement); _dab && _fc.Name == start.Name {
			break
		}
	}
	return nil
}
func (_gccd *ST_TargetMode) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_gccd = 0
	case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":
		*_gccd = 1
	case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":
		*_gccd = 2
	}
	return nil
}

type CT_Relationships struct{ Relationship []*Relationship }

func NewCT_Relationships() *CT_Relationships { _bde := &CT_Relationships{}; return _bde }
func (_cag ST_TargetMode) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_gaa := _e.Attr{}
	_gaa.Name = name
	switch _cag {
	case ST_TargetModeUnset:
		_gaa.Value = ""
	case ST_TargetModeExternal:
		_gaa.Value = "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c"
	case ST_TargetModeInternal:
		_gaa.Value = "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c"
	}
	return _gaa, nil
}
func (_cgb *Relationships) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	start.Attr = append(start.Attr, _e.Attr{Name: _e.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s"})
	start.Attr = append(start.Attr, _e.Attr{Name: _e.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"
	return _cgb.CT_Relationships.MarshalXML(e, start)
}

// ValidateWithPath validates the Relationships and its children, prefixing error messages with path
func (_fac *Relationships) ValidateWithPath(path string) error {
	if _cec := _fac.CT_Relationships.ValidateWithPath(path); _cec != nil {
		return _cec
	}
	return nil
}

// ValidateWithPath validates the Relationship and its children, prefixing error messages with path
func (_aa *Relationship) ValidateWithPath(path string) error {
	if _ff := _aa.CT_Relationship.ValidateWithPath(path); _ff != nil {
		return _ff
	}
	return nil
}
func (_fa *CT_Relationship) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	if _fa.TargetModeAttr != ST_TargetModeUnset {
		_bg, _eg := _fa.TargetModeAttr.MarshalXMLAttr(_e.Name{Local: "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"})
		if _eg != nil {
			return _eg
		}
		start.Attr = append(start.Attr, _bg)
	}
	start.Attr = append(start.Attr, _e.Attr{Name: _e.Name{Local: "\u0054\u0061\u0072\u0067\u0065\u0074"}, Value: _b.Sprintf("\u0025\u0076", _fa.TargetAttr)})
	start.Attr = append(start.Attr, _e.Attr{Name: _e.Name{Local: "\u0054\u0079\u0070\u0065"}, Value: _b.Sprintf("\u0025\u0076", _fa.TypeAttr)})
	start.Attr = append(start.Attr, _e.Attr{Name: _e.Name{Local: "\u0049\u0064"}, Value: _b.Sprintf("\u0025\u0076", _fa.IdAttr)})
	e.EncodeElement(_fa.Content, start)
	e.EncodeToken(_e.EndElement{Name: start.Name})
	return nil
}

// Validate validates the Relationships and its children
func (_ecd *Relationships) Validate() error {
	return _ecd.ValidateWithPath("\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073")
}

type ST_TargetMode byte

func (_ece *Relationship) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return _ece.CT_Relationship.MarshalXML(e, start)
}

// Validate validates the CT_Relationship and its children
func (_bb *CT_Relationship) Validate() error {
	return _bb.ValidateWithPath("\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070")
}
func init() {
	_g.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073", NewCT_Relationships)
	_g.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070", NewCT_Relationship)
	_g.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073", NewRelationships)
	_g.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070", NewRelationship)
}
