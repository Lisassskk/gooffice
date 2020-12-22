package picture

import (
	_c "encoding/xml"

	_g "gitee.com/yu_sheng/gooffice"
	_f "gitee.com/yu_sheng/gooffice/schema/soo/dml"
)

func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	_ad := &CT_PictureNonVisual{}
	_ad.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_ad.CNvPicPr = _f.NewCT_NonVisualPictureProperties()
	return _ad
}
func (_ggb *Pic) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0070\u0069c"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0070i\u0063\u003a\u0070\u0069\u0063"
	return _ggb.CT_Picture.MarshalXML(e, start)
}
func NewCT_Picture() *CT_Picture {
	_e := &CT_Picture{}
	_e.NvPicPr = NewCT_PictureNonVisual()
	_e.BlipFill = _f.NewCT_BlipFillProperties()
	_e.SpPr = _f.NewCT_ShapeProperties()
	return _e
}

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_da *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if _bc := _da.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _bc != nil {
		return _bc
	}
	if _ff := _da.CNvPicPr.ValidateWithPath(path + "\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r"); _ff != nil {
		return _ff
	}
	return nil
}
func (_efc *CT_PictureNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_efe := _c.StartElement{Name: _c.Name{Local: "\u0070i\u0063\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_efc.CNvPr, _efe)
	_ade := _c.StartElement{Name: _c.Name{Local: "\u0070\u0069\u0063:\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_efc.CNvPicPr, _ade)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type CT_Picture struct {
	NvPicPr  *CT_PictureNonVisual
	BlipFill *_f.CT_BlipFillProperties
	SpPr     *_f.CT_ShapeProperties
}
type Pic struct{ CT_Picture }

// Validate validates the Pic and its children
func (_ag *Pic) Validate() error { return _ag.ValidateWithPath("\u0050\u0069\u0063") }
func (_ee *CT_Picture) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ee.NvPicPr = NewCT_PictureNonVisual()
	_ee.BlipFill = _f.NewCT_BlipFillProperties()
	_ee.SpPr = _f.NewCT_ShapeProperties()
_fe:
	for {
		_bf, _ca := d.Token()
		if _ca != nil {
			return _ca
		}
		switch _dc := _bf.(type) {
		case _c.StartElement:
			switch _dc.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _be := d.DecodeElement(_ee.NvPicPr, &_dc); _be != nil {
					return _be
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _db := d.DecodeElement(_ee.BlipFill, &_dc); _db != nil {
					return _db
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}:
				if _fee := d.DecodeElement(_ee.SpPr, &_dc); _fee != nil {
					return _fee
				}
			default:
				_g.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076", _dc.Name)
				if _bg := d.Skip(); _bg != nil {
					return _bg
				}
			}
		case _c.EndElement:
			break _fe
		case _c.CharData:
		}
	}
	return nil
}
func (_d *CT_Picture) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_b := _c.StartElement{Name: _c.Name{Local: "p\u0069\u0063\u003a\u006e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_d.NvPicPr, _b)
	_ea := _c.StartElement{Name: _c.Name{Local: "\u0070\u0069\u0063:\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}}
	e.EncodeElement(_d.BlipFill, _ea)
	_ef := _c.StartElement{Name: _c.Name{Local: "\u0070\u0069\u0063\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_d.SpPr, _ef)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func NewPic() *Pic { _beg := &Pic{}; _beg.CT_Picture = *NewCT_Picture(); return _beg }

type CT_PictureNonVisual struct {
	CNvPr    *_f.CT_NonVisualDrawingProps
	CNvPicPr *_f.CT_NonVisualPictureProperties
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_fc *CT_Picture) ValidateWithPath(path string) error {
	if _cd := _fc.NvPicPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072"); _cd != nil {
		return _cd
	}
	if _gg := _fc.BlipFill.ValidateWithPath(path + "\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl"); _gg != nil {
		return _gg
	}
	if _dcd := _fc.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _dcd != nil {
		return _dcd
	}
	return nil
}
func (_fb *Pic) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_fb.CT_Picture = *NewCT_Picture()
_de:
	for {
		_ae, _gf := d.Token()
		if _gf != nil {
			return _gf
		}
		switch _gga := _ae.(type) {
		case _c.StartElement:
			switch _gga.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _fa := d.DecodeElement(_fb.NvPicPr, &_gga); _fa != nil {
					return _fa
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _bfb := d.DecodeElement(_fb.BlipFill, &_gga); _bfb != nil {
					return _bfb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}:
				if _cdg := d.DecodeElement(_fb.SpPr, &_gga); _cdg != nil {
					return _cdg
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u0050i\u0063\u0020\u0025\u0076", _gga.Name)
				if _fec := d.Skip(); _fec != nil {
					return _fec
				}
			}
		case _c.EndElement:
			break _de
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the Pic and its children, prefixing error messages with path
func (_deb *Pic) ValidateWithPath(path string) error {
	if _fcf := _deb.CT_Picture.ValidateWithPath(path); _fcf != nil {
		return _fcf
	}
	return nil
}

// Validate validates the CT_Picture and its children
func (_ab *CT_Picture) Validate() error {
	return _ab.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}
func (_cg *CT_PictureNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_cg.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_cg.CNvPicPr = _f.NewCT_NonVisualPictureProperties()
_fd:
	for {
		_fcc, _adc := d.Token()
		if _adc != nil {
			return _adc
		}
		switch _fg := _fcc.(type) {
		case _c.StartElement:
			switch _fg.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0063\u004e\u0076P\u0072"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0063\u004e\u0076P\u0072"}:
				if _ba := d.DecodeElement(_cg.CNvPr, &_fg); _ba != nil {
					return _ba
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}, _c.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:
				if _dg := d.DecodeElement(_cg.CNvPicPr, &_fg); _dg != nil {
					return _dg
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _fg.Name)
				if _ed := d.Skip(); _ed != nil {
					return _ed
				}
			}
		case _c.EndElement:
			break _fd
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PictureNonVisual and its children
func (_adcg *CT_PictureNonVisual) Validate() error {
	return _adcg.ValidateWithPath("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}
func init() {
	_g.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_PictureNonVisual)
	_g.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_g.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0070\u0069\u0063", NewPic)
}
