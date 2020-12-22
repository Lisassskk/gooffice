package schemaLibrary

import (
	_fa "encoding/xml"
	_d "fmt"

	_a "gitee.com/yu_sheng/gooffice"
)

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_fab *CT_Schema) ValidateWithPath(path string) error { return nil }
func NewCT_Schema() *CT_Schema                             { _fb := &CT_Schema{}; return _fb }

// Validate validates the CT_SchemaLibrary and its children
func (_cdg *CT_SchemaLibrary) Validate() error {
	return _cdg.ValidateWithPath("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079")
}
func (_af *CT_Schema) MarshalXML(e *_fa.Encoder, start _fa.StartElement) error {
	if _af.UriAttr != nil {
		start.Attr = append(start.Attr, _fa.Attr{Name: _fa.Name{Local: "\u006d\u0061\u003a\u0075\u0072\u0069"}, Value: _d.Sprintf("\u0025\u0076", *_af.UriAttr)})
	}
	if _af.ManifestLocationAttr != nil {
		start.Attr = append(start.Attr, _fa.Attr{Name: _fa.Name{Local: "\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"}, Value: _d.Sprintf("\u0025\u0076", *_af.ManifestLocationAttr)})
	}
	if _af.SchemaLocationAttr != nil {
		start.Attr = append(start.Attr, _fa.Attr{Name: _fa.Name{Local: "\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"}, Value: _d.Sprintf("\u0025\u0076", *_af.SchemaLocationAttr)})
	}
	if _af.SchemaLanguageAttr != nil {
		start.Attr = append(start.Attr, _fa.Attr{Name: _fa.Name{Local: "\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"}, Value: _d.Sprintf("\u0025\u0076", *_af.SchemaLanguageAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_fa.EndElement{Name: start.Name})
	return nil
}
func (_ec *CT_SchemaLibrary) MarshalXML(e *_fa.Encoder, start _fa.StartElement) error {
	e.EncodeToken(start)
	if _ec.Schema != nil {
		_ad := _fa.StartElement{Name: _fa.Name{Local: "\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}}
		for _, _cab := range _ec.Schema {
			e.EncodeElement(_cab, _ad)
		}
	}
	e.EncodeToken(_fa.EndElement{Name: start.Name})
	return nil
}
func (_fbd *CT_SchemaLibrary) UnmarshalXML(d *_fa.Decoder, start _fa.StartElement) error {
_bf:
	for {
		_dd, _eb := d.Token()
		if _eb != nil {
			return _eb
		}
		switch _fe := _dd.(type) {
		case _fa.StartElement:
			switch _fe.Name {
			case _fa.Name{Space: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0063\u0068\u0065\u006d\u0061"}:
				_da := NewCT_Schema()
				if _ef := d.DecodeElement(_da, &_fe); _ef != nil {
					return _ef
				}
				_fbd.Schema = append(_fbd.Schema, _da)
			default:
				_a.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v", _fe.Name)
				if _efb := d.Skip(); _efb != nil {
					return _efb
				}
			}
		case _fa.EndElement:
			break _bf
		case _fa.CharData:
		}
	}
	return nil
}
func (_db *SchemaLibrary) MarshalXML(e *_fa.Encoder, start _fa.StartElement) error {
	start.Attr = append(start.Attr, _fa.Attr{Name: _fa.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _fa.Attr{Name: _fa.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"}, Value: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _fa.Attr{Name: _fa.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079"
	return _db.CT_SchemaLibrary.MarshalXML(e, start)
}

// Validate validates the SchemaLibrary and its children
func (_gc *SchemaLibrary) Validate() error {
	return _gc.ValidateWithPath("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079")
}

type CT_SchemaLibrary struct{ Schema []*CT_Schema }

func NewSchemaLibrary() *SchemaLibrary {
	_gb := &SchemaLibrary{}
	_gb.CT_SchemaLibrary = *NewCT_SchemaLibrary()
	return _gb
}

type CT_Schema struct {
	UriAttr              *string
	ManifestLocationAttr *string
	SchemaLocationAttr   *string
	SchemaLanguageAttr   *string
}

// Validate validates the CT_Schema and its children
func (_dg *CT_Schema) Validate() error {
	return _dg.ValidateWithPath("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da")
}
func (_gdd *SchemaLibrary) UnmarshalXML(d *_fa.Decoder, start _fa.StartElement) error {
	_gdd.CT_SchemaLibrary = *NewCT_SchemaLibrary()
_ecb:
	for {
		_cdf, _gga := d.Token()
		if _gga != nil {
			return _gga
		}
		switch _cg := _cdf.(type) {
		case _fa.StartElement:
			switch _cg.Name {
			case _fa.Name{Space: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0063\u0068\u0065\u006d\u0061"}:
				_ggab := NewCT_Schema()
				if _efa := d.DecodeElement(_ggab, &_cg); _efa != nil {
					return _efa
				}
				_gdd.Schema = append(_gdd.Schema, _ggab)
			default:
				_a.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076", _cg.Name)
				if _ba := d.Skip(); _ba != nil {
					return _ba
				}
			}
		case _fa.EndElement:
			break _ecb
		case _fa.CharData:
		}
	}
	return nil
}

type SchemaLibrary struct{ CT_SchemaLibrary }

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_ab *SchemaLibrary) ValidateWithPath(path string) error {
	if _aa := _ab.CT_SchemaLibrary.ValidateWithPath(path); _aa != nil {
		return _aa
	}
	return nil
}
func NewCT_SchemaLibrary() *CT_SchemaLibrary { _ca := &CT_SchemaLibrary{}; return _ca }

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_bfa *CT_SchemaLibrary) ValidateWithPath(path string) error {
	for _ea, _dc := range _bfa.Schema {
		if _fg := _dc.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d", path, _ea)); _fg != nil {
			return _fg
		}
	}
	return nil
}
func (_b *CT_Schema) UnmarshalXML(d *_fa.Decoder, start _fa.StartElement) error {
	for _, _afa := range start.Attr {
		if _afa.Name.Local == "\u0075\u0072\u0069" {
			_c, _g := _afa.Value, error(nil)
			if _g != nil {
				return _g
			}
			_b.UriAttr = &_c
			continue
		}
		if _afa.Name.Local == "\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e" {
			_gg, _e := _afa.Value, error(nil)
			if _e != nil {
				return _e
			}
			_b.ManifestLocationAttr = &_gg
			continue
		}
		if _afa.Name.Local == "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e" {
			_fd, _ggg := _afa.Value, error(nil)
			if _ggg != nil {
				return _ggg
			}
			_b.SchemaLocationAttr = &_fd
			continue
		}
		if _afa.Name.Local == "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065" {
			_fc, _fae := _afa.Value, error(nil)
			if _fae != nil {
				return _fae
			}
			_b.SchemaLanguageAttr = &_fc
			continue
		}
	}
	for {
		_fcc, _ee := d.Token()
		if _ee != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073", _ee)
		}
		if _gd, _de := _fcc.(_fa.EndElement); _de && _gd.Name == start.Name {
			break
		}
	}
	return nil
}
func init() {
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0043T\u005f\u0053\u0063\u0068\u0065\u006da", NewCT_Schema)
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079", NewCT_SchemaLibrary)
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079", NewSchemaLibrary)
}
