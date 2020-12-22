package powerpoint

import (
	_b "encoding/xml"
	_g "fmt"

	_a "gitee.com/gooffice/gooffice"
)

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_egc *CT_Rel) ValidateWithPath(path string) error { return nil }
func NewCT_Empty() *CT_Empty                            { _gb := &CT_Empty{}; return _gb }
func (_dg *CT_Rel) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	if _dg.IdAttr != nil {
		start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0069\u0064"}, Value: _g.Sprintf("\u0025\u0076", *_dg.IdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

type CT_Rel struct{ IdAttr *string }

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_cf *CT_Empty) ValidateWithPath(path string) error { return nil }

// Validate validates the Textdata and its children
func (_gbdb *Textdata) Validate() error {
	return _gbdb.ValidateWithPath("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061")
}

type Textdata struct{ CT_Rel }

func (_edf *Iscomment) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_edf.CT_Empty = *NewCT_Empty()
	for {
		_gbd, _db := d.Token()
		if _db != nil {
			return _g.Errorf("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073", _db)
		}
		if _da, _cfb := _gbd.(_b.EndElement); _cfb && _da.Name == start.Name {
			break
		}
	}
	return nil
}
func (_bcc *Textdata) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061"
	return _bcc.CT_Rel.MarshalXML(e, start)
}
func (_dad *Textdata) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_dad.CT_Rel = *NewCT_Rel()
	for _, _ba := range start.Attr {
		if _ba.Name.Local == "\u0069\u0064" {
			_fd, _dge := _ba.Value, error(nil)
			if _dge != nil {
				return _dge
			}
			_dad.IdAttr = &_fd
			continue
		}
	}
	for {
		_gc, _gce := d.Token()
		if _gce != nil {
			return _g.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073", _gce)
		}
		if _edd, _bac := _gc.(_b.EndElement); _bac && _edd.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ab *CT_Empty) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_f *CT_Empty) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_e, _abb := d.Token()
		if _abb != nil {
			return _g.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073", _abb)
		}
		if _eb, _gd := _e.(_b.EndElement); _gd && _eb.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Iscomment and its children
func (_ff *Iscomment) Validate() error {
	return _ff.ValidateWithPath("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et")
}

type CT_Empty struct{}

func (_abg *Iscomment) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0069s\u0063\u006f\u006d\u006d\u0065\u006et"
	return _abg.CT_Empty.MarshalXML(e, start)
}

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_be *Textdata) ValidateWithPath(path string) error {
	if _aba := _be.CT_Rel.ValidateWithPath(path); _aba != nil {
		return _aba
	}
	return nil
}
func NewTextdata() *Textdata { _aaa := &Textdata{}; _aaa.CT_Rel = *NewCT_Rel(); return _aaa }

type Iscomment struct{ CT_Empty }

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_fb *Iscomment) ValidateWithPath(path string) error {
	if _ca := _fb.CT_Empty.ValidateWithPath(path); _ca != nil {
		return _ca
	}
	return nil
}
func (_ed *CT_Rel) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for _, _ge := range start.Attr {
		if _ge.Name.Local == "\u0069\u0064" {
			_bc, _aa := _ge.Value, error(nil)
			if _aa != nil {
				return _aa
			}
			_ed.IdAttr = &_bc
			continue
		}
	}
	for {
		_fa, _bcb := d.Token()
		if _bcb != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073", _bcb)
		}
		if _gbb, _eg := _fa.(_b.EndElement); _eg && _gbb.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_Rel() *CT_Rel { _d := &CT_Rel{}; return _d }

// Validate validates the CT_Empty and its children
func (_gg *CT_Empty) Validate() error {
	return _gg.ValidateWithPath("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079")
}
func NewIscomment() *Iscomment { _ggb := &Iscomment{}; _ggb.CT_Empty = *NewCT_Empty(); return _ggb }

// Validate validates the CT_Rel and its children
func (_cc *CT_Rel) Validate() error {
	return _cc.ValidateWithPath("\u0043\u0054\u005f\u0052\u0065\u006c")
}
func init() {
	_a.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079", NewCT_Empty)
	_a.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0043\u0054\u005f\u0052\u0065\u006c", NewCT_Rel)
	_a.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0069s\u0063\u006f\u006d\u006d\u0065\u006et", NewIscomment)
	_a.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061", NewTextdata)
}
