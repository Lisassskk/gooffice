package content_types

import (
	_d "encoding/xml"
	_b "fmt"
	_dc "regexp"

	_f "gitee.com/yu_sheng/gooffice"
)

// Validate validates the CT_Default and its children
func (_gbf *CT_Default) Validate() error {
	return _gbf.ValidateWithPath("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074")
}

// Validate validates the CT_Override and its children
func (_ga *CT_Override) Validate() error {
	return _ga.ValidateWithPath("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065")
}
func NewCT_Types() *CT_Types { _ccc := &CT_Types{}; return _ccc }

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_ea *Override) ValidateWithPath(path string) error {
	if _fef := _ea.CT_Override.ValidateWithPath(path); _fef != nil {
		return _fef
	}
	return nil
}
func (_e *CT_Default) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_e.ExtensionAttr = "\u0078\u006d\u006c"
	_e.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	for _, _ad := range start.Attr {
		if _ad.Name.Local == "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn" {
			_bb, _df := _ad.Value, error(nil)
			if _df != nil {
				return _df
			}
			_e.ExtensionAttr = _bb
			continue
		}
		if _ad.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_gb, _gf := _ad.Value, error(nil)
			if _gf != nil {
				return _gf
			}
			_e.ContentTypeAttr = _gb
			continue
		}
	}
	for {
		_eb, _aa := d.Token()
		if _aa != nil {
			return _b.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073", _aa)
		}
		if _ce, _ef := _eb.(_d.EndElement); _ef && _ce.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_Override struct {
	ContentTypeAttr string
	PartNameAttr    string
}

func NewDefault() *Default { _dgc := &Default{}; _dgc.CT_Default = *NewCT_Default(); return _dgc }

type Override struct{ CT_Override }

// Validate validates the Types and its children
func (_cb *Types) Validate() error { return _cb.ValidateWithPath("\u0054\u0079\u0070e\u0073") }

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_fe *CT_Override) ValidateWithPath(path string) error {
	if !ST_ContentTypePatternRe.MatchString(_fe.ContentTypeAttr) {
		return _b.Errorf("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, ST_ContentTypePatternRe, _fe.ContentTypeAttr)
	}
	return nil
}
func (_bc *CT_Override) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"}, Value: _b.Sprintf("\u0025\u0076", _bc.ContentTypeAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"}, Value: _b.Sprintf("\u0025\u0076", _bc.PartNameAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the Default and its children
func (_cag *Default) Validate() error {
	return _cag.ValidateWithPath("\u0044e\u0066\u0061\u0075\u006c\u0074")
}
func (_fcg *Default) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return _fcg.CT_Default.MarshalXML(e, start)
}

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_eg *CT_Default) ValidateWithPath(path string) error {
	if !ST_ExtensionPatternRe.MatchString(_eg.ExtensionAttr) {
		return _b.Errorf("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029", path, ST_ExtensionPatternRe, _eg.ExtensionAttr)
	}
	if !ST_ContentTypePatternRe.MatchString(_eg.ContentTypeAttr) {
		return _b.Errorf("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, ST_ContentTypePatternRe, _eg.ContentTypeAttr)
	}
	return nil
}
func (_ege *CT_Types) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _ege.Default != nil {
		_fdf := _d.StartElement{Name: _d.Name{Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}}
		for _, _ge := range _ege.Default {
			e.EncodeElement(_ge, _fdf)
		}
	}
	if _ege.Override != nil {
		_ba := _d.StartElement{Name: _d.Name{Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}}
		for _, _cf := range _ege.Override {
			e.EncodeElement(_cf, _ba)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_gfg *Override) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return _gfg.CT_Override.MarshalXML(e, start)
}
func (_dd *Override) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dd.CT_Override = *NewCT_Override()
	for _, _bce := range start.Attr {
		if _bce.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_fb, _ecd := _bce.Value, error(nil)
			if _ecd != nil {
				return _ecd
			}
			_dd.ContentTypeAttr = _fb
			continue
		}
		if _bce.Name.Local == "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065" {
			_dee, _fae := _bce.Value, error(nil)
			if _fae != nil {
				return _fae
			}
			_dd.PartNameAttr = _dee
			continue
		}
	}
	for {
		_daf, _bcd := d.Token()
		if _bcd != nil {
			return _b.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073", _bcd)
		}
		if _bda, _dae := _daf.(_d.EndElement); _dae && _bda.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ca *CT_Default) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"}, Value: _b.Sprintf("\u0025\u0076", _ca.ExtensionAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"}, Value: _b.Sprintf("\u0025\u0076", _ca.ContentTypeAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewCT_Default() *CT_Default {
	_c := &CT_Default{}
	_c.ExtensionAttr = "\u0078\u006d\u006c"
	_c.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	return _c
}

type CT_Types struct {
	Default  []*Default
	Override []*Override
}

func (_deeb *Types) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0054\u0079\u0070e\u0073"
	return _deeb.CT_Types.MarshalXML(e, start)
}

var ST_ContentTypePatternRe = _dc.MustCompile(ST_ContentTypePattern)

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_cee *Types) ValidateWithPath(path string) error {
	if _bbc := _cee.CT_Types.ValidateWithPath(path); _bbc != nil {
		return _bbc
	}
	return nil
}

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_ff *CT_Types) ValidateWithPath(path string) error {
	for _bef, _fab := range _ff.Default {
		if _bdc := _fab.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d", path, _bef)); _bdc != nil {
			return _bdc
		}
	}
	for _ada, _eba := range _ff.Override {
		if _bbb := _eba.ValidateWithPath(_b.Sprintf("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d", path, _ada)); _bbb != nil {
			return _bbb
		}
	}
	return nil
}

const ST_ExtensionPattern = "\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b"

func (_da *CT_Types) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_gag:
	for {
		_ebd, _dge := d.Token()
		if _dge != nil {
			return _dge
		}
		switch _dgac := _ebd.(type) {
		case _d.StartElement:
			switch _dgac.Name {
			case _d.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_fge := NewDefault()
				if _fa := d.DecodeElement(_fge, &_dgac); _fa != nil {
					return _fa
				}
				_da.Default = append(_da.Default, _fge)
			case _d.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:
				_cg := NewOverride()
				if _fc := d.DecodeElement(_cg, &_dgac); _fc != nil {
					return _fc
				}
				_da.Override = append(_da.Override, _cg)
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076", _dgac.Name)
				if _db := d.Skip(); _db != nil {
					return _db
				}
			}
		case _d.EndElement:
			break _gag
		case _d.CharData:
		}
	}
	return nil
}

type Types struct{ CT_Types }

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_bdf *Default) ValidateWithPath(path string) error {
	if _baa := _bdf.CT_Default.ValidateWithPath(path); _baa != nil {
		return _baa
	}
	return nil
}

// Validate validates the CT_Types and its children
func (_cec *CT_Types) Validate() error {
	return _cec.ValidateWithPath("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073")
}

// Validate validates the Override and its children
func (_bcgf *Override) Validate() error {
	return _bcgf.ValidateWithPath("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065")
}

type CT_Default struct {
	ExtensionAttr   string
	ContentTypeAttr string
}

func NewCT_Override() *CT_Override {
	_be := &CT_Override{}
	_be.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	return _be
}
func (_dg *CT_Override) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dg.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	for _, _fd := range start.Attr {
		if _fd.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_ceg, _cc := _fd.Value, error(nil)
			if _cc != nil {
				return _cc
			}
			_dg.ContentTypeAttr = _ceg
			continue
		}
		if _fd.Name.Local == "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065" {
			_fg, _ab := _fd.Value, error(nil)
			if _ab != nil {
				return _ab
			}
			_dg.PartNameAttr = _fg
			continue
		}
	}
	for {
		_fda, _ec := d.Token()
		if _ec != nil {
			return _b.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073", _ec)
		}
		if _ee, _gbd := _fda.(_d.EndElement); _gbd && _ee.Name == start.Name {
			break
		}
	}
	return nil
}

const ST_ContentTypePattern = "\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024"

func (_cfc *Types) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cfc.CT_Types = *NewCT_Types()
_gd:
	for {
		_efb, _ac := d.Token()
		if _ac != nil {
			return _ac
		}
		switch _dgaf := _efb.(type) {
		case _d.StartElement:
			switch _dgaf.Name {
			case _d.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_abf := NewDefault()
				if _cfe := d.DecodeElement(_abf, &_dgaf); _cfe != nil {
					return _cfe
				}
				_cfc.Default = append(_cfc.Default, _abf)
			case _d.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:
				_aba := NewOverride()
				if _dega := d.DecodeElement(_aba, &_dgaf); _dega != nil {
					return _dega
				}
				_cfc.Override = append(_cfc.Override, _aba)
			default:
				_f.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076", _dgaf.Name)
				if _bcf := d.Skip(); _bcf != nil {
					return _bcf
				}
			}
		case _d.EndElement:
			break _gd
		case _d.CharData:
		}
	}
	return nil
}
func NewTypes() *Types { _fbg := &Types{}; _fbg.CT_Types = *NewCT_Types(); return _fbg }

var ST_ExtensionPatternRe = _dc.MustCompile(ST_ExtensionPattern)

func (_cccc *Default) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cccc.CT_Default = *NewCT_Default()
	for _, _de := range start.Attr {
		if _de.Name.Local == "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn" {
			_gfc, _cd := _de.Value, error(nil)
			if _cd != nil {
				return _cd
			}
			_cccc.ExtensionAttr = _gfc
			continue
		}
		if _de.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_dgef, _cecc := _de.Value, error(nil)
			if _cecc != nil {
				return _cecc
			}
			_cccc.ContentTypeAttr = _dgef
			continue
		}
	}
	for {
		_ae, _gac := d.Token()
		if _gac != nil {
			return _b.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073", _gac)
		}
		if _aab, _gae := _ae.(_d.EndElement); _gae && _aab.Name == start.Name {
			break
		}
	}
	return nil
}
func NewOverride() *Override { _bcg := &Override{}; _bcg.CT_Override = *NewCT_Override(); return _bcg }

type Default struct{ CT_Default }

func init() {
	_f.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073", NewCT_Types)
	_f.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074", NewCT_Default)
	_f.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065", NewCT_Override)
	_f.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0054\u0079\u0070e\u0073", NewTypes)
	_f.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0044e\u0066\u0061\u0075\u006c\u0074", NewDefault)
	_f.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065", NewOverride)
}
