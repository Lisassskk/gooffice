package core_properties

import (
	_gg "encoding/xml"
	_b "fmt"
	_e "time"

	_ec "gitee.com/gooffice/gooffice"
)

func (_ad *CoreProperties) UnmarshalXML(d *_gg.Decoder, start _gg.StartElement) error {
	_ad.CT_CoreProperties = *NewCT_CoreProperties()
_fg:
	for {
		_ecg, _fgb := d.Token()
		if _fgb != nil {
			return _fgb
		}
		switch _acfb := _ecg.(type) {
		case _gg.StartElement:
			switch _acfb.Name {
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:
				_ad.Category = new(string)
				if _aga := d.DecodeElement(_ad.Category, &_acfb); _aga != nil {
					return _aga
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:
				_ad.ContentStatus = new(string)
				if _acc := d.DecodeElement(_ad.ContentStatus, &_acfb); _acc != nil {
					return _acc
				}
			case _gg.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u0063r\u0065\u0061\u0074\u0065\u0064"}:
				_ad.Created = new(_ec.XSDAny)
				if _ege := d.DecodeElement(_ad.Created, &_acfb); _ege != nil {
					return _ege
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0063r\u0065\u0061\u0074\u006f\u0072"}:
				_ad.Creator = new(_ec.XSDAny)
				if _gbg := d.DecodeElement(_ad.Creator, &_acfb); _gbg != nil {
					return _gbg
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:
				_ad.Description = new(_ec.XSDAny)
				if _fc := d.DecodeElement(_ad.Description, &_acfb); _fc != nil {
					return _fc
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:
				_ad.Identifier = new(_ec.XSDAny)
				if _bdb := d.DecodeElement(_ad.Identifier, &_acfb); _bdb != nil {
					return _bdb
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:
				_ad.Keywords = NewCT_Keywords()
				if _gad := d.DecodeElement(_ad.Keywords, &_acfb); _gad != nil {
					return _gad
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:
				_ad.Language = new(_ec.XSDAny)
				if _gdf := d.DecodeElement(_ad.Language, &_acfb); _gdf != nil {
					return _gdf
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:
				_ad.LastModifiedBy = new(string)
				if _ccc := d.DecodeElement(_ad.LastModifiedBy, &_acfb); _ccc != nil {
					return _ccc
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:
				_ad.LastPrinted = new(_e.Time)
				if _cbb := d.DecodeElement(_ad.LastPrinted, &_acfb); _cbb != nil {
					return _cbb
				}
			case _gg.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:
				_ad.Modified = new(_ec.XSDAny)
				if _dad := d.DecodeElement(_ad.Modified, &_acfb); _dad != nil {
					return _dad
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:
				_ad.Revision = new(string)
				if _bbg := d.DecodeElement(_ad.Revision, &_acfb); _bbg != nil {
					return _bbg
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0073u\u0062\u006a\u0065\u0063\u0074"}:
				_ad.Subject = new(_ec.XSDAny)
				if _bac := d.DecodeElement(_ad.Subject, &_acfb); _bac != nil {
					return _bac
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0074\u0069\u0074l\u0065"}:
				_ad.Title = new(_ec.XSDAny)
				if _aae := d.DecodeElement(_ad.Title, &_acfb); _aae != nil {
					return _aae
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0076e\u0072\u0073\u0069\u006f\u006e"}:
				_ad.Version = new(string)
				if _gga := d.DecodeElement(_ad.Version, &_acfb); _gga != nil {
					return _gga
				}
			default:
				_ec.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072t\u0069e\u0073\u0020\u0025\u0076", _acfb.Name)
				if _ef := d.Skip(); _ef != nil {
					return _ef
				}
			}
		case _gg.EndElement:
			break _fg
		case _gg.CharData:
		}
	}
	return nil
}
func (_bfb *CT_Keywords) UnmarshalXML(d *_gg.Decoder, start _gg.StartElement) error {
	for _, _gge := range start.Attr {
		if _gge.Name.Space == "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065" && _gge.Name.Local == "\u006c\u0061\u006e\u0067" {
			_bgg, _eacc := _gge.Value, error(nil)
			if _eacc != nil {
				return _eacc
			}
			_bfb.LangAttr = &_bgg
			continue
		}
	}
_age:
	for {
		_bef, _gf := d.Token()
		if _gf != nil {
			return _gf
		}
		switch _agc := _bef.(type) {
		case _gg.StartElement:
			switch _agc.Name {
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0076\u0061\u006cu\u0065"}:
				_da := NewCT_Keyword()
				if _ace := d.DecodeElement(_da, &_agc); _ace != nil {
					return _ace
				}
				_bfb.Value = append(_bfb.Value, _da)
			default:
				_ec.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073\u0020\u0025\u0076", _agc.Name)
				if _fbg := d.Skip(); _fbg != nil {
					return _fbg
				}
			}
		case _gg.EndElement:
			break _age
		case _gg.CharData:
		}
	}
	return nil
}
func NewCT_Keywords() *CT_Keywords             { _bed := &CT_Keywords{}; return _bed }
func NewCT_CoreProperties() *CT_CoreProperties { _f := &CT_CoreProperties{}; return _f }

// Validate validates the CT_CoreProperties and its children
func (_gd *CT_CoreProperties) Validate() error {
	return _gd.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073")
}
func (_gc *CT_Keyword) UnmarshalXML(d *_gg.Decoder, start _gg.StartElement) error {
	for _, _bgd := range start.Attr {
		if _bgd.Name.Space == "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065" && _bgd.Name.Local == "\u006c\u0061\u006e\u0067" {
			_aedc, _acf := _bgd.Value, error(nil)
			if _acf != nil {
				return _acf
			}
			_gc.LangAttr = &_aedc
			continue
		}
	}
	for {
		_fd, _aa := d.Token()
		if _aa != nil {
			return _b.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u003a\u0020%\u0073", _aa)
		}
		if _bgb, _aeda := _fd.(_gg.CharData); _aeda {
			_gc.Content = string(_bgb)
		}
		if _gcc, _cdd := _fd.(_gg.EndElement); _cdd && _gcc.Name == start.Name {
			break
		}
	}
	return nil
}
func (_a *CT_CoreProperties) MarshalXML(e *_gg.Encoder, start _gg.StartElement) error {
	e.EncodeToken(start)
	if _a.Category != nil {
		_d := _gg.StartElement{Name: _gg.Name{Local: "c\u0070\u003a\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}}
		_ec.AddPreserveSpaceAttr(&_d, *_a.Category)
		e.EncodeElement(_a.Category, _d)
	}
	if _a.ContentStatus != nil {
		_dc := _gg.StartElement{Name: _gg.Name{Local: "\u0063\u0070:\u0063\u006f\u006et\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}}
		_ec.AddPreserveSpaceAttr(&_dc, *_a.ContentStatus)
		e.EncodeElement(_a.ContentStatus, _dc)
	}
	if _a.Created != nil {
		_ff := _gg.StartElement{Name: _gg.Name{Local: "\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064"}}
		e.EncodeElement(_a.Created, _ff)
	}
	if _a.Creator != nil {
		_c := _gg.StartElement{Name: _gg.Name{Local: "\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}}
		e.EncodeElement(_a.Creator, _c)
	}
	if _a.Description != nil {
		_de := _gg.StartElement{Name: _gg.Name{Local: "\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}}
		e.EncodeElement(_a.Description, _de)
	}
	if _a.Identifier != nil {
		_fb := _gg.StartElement{Name: _gg.Name{Local: "\u0064\u0063\u003a\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}}
		e.EncodeElement(_a.Identifier, _fb)
	}
	if _a.Keywords != nil {
		_cb := _gg.StartElement{Name: _gg.Name{Local: "c\u0070\u003a\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}}
		e.EncodeElement(_a.Keywords, _cb)
	}
	if _a.Language != nil {
		_cf := _gg.StartElement{Name: _gg.Name{Local: "d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}}
		e.EncodeElement(_a.Language, _cf)
	}
	if _a.LastModifiedBy != nil {
		_eg := _gg.StartElement{Name: _gg.Name{Local: "\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}}
		_ec.AddPreserveSpaceAttr(&_eg, *_a.LastModifiedBy)
		e.EncodeElement(_a.LastModifiedBy, _eg)
	}
	if _a.LastPrinted != nil {
		_bf := _gg.StartElement{Name: _gg.Name{Local: "\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u0050\u0072i\u006e\u0074\u0065\u0064"}}
		e.EncodeElement(_a.LastPrinted, _bf)
	}
	if _a.Modified != nil {
		_fe := _gg.StartElement{Name: _gg.Name{Local: "\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}}
		e.EncodeElement(_a.Modified, _fe)
	}
	if _a.Revision != nil {
		_cc := _gg.StartElement{Name: _gg.Name{Local: "c\u0070\u003a\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}}
		_ec.AddPreserveSpaceAttr(&_cc, *_a.Revision)
		e.EncodeElement(_a.Revision, _cc)
	}
	if _a.Subject != nil {
		_cd := _gg.StartElement{Name: _gg.Name{Local: "\u0064\u0063\u003a\u0073\u0075\u0062\u006a\u0065\u0063\u0074"}}
		e.EncodeElement(_a.Subject, _cd)
	}
	if _a.Title != nil {
		_ea := _gg.StartElement{Name: _gg.Name{Local: "\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}}
		e.EncodeElement(_a.Title, _ea)
	}
	if _a.Version != nil {
		_cg := _gg.StartElement{Name: _gg.Name{Local: "\u0063\u0070\u003a\u0076\u0065\u0072\u0073\u0069\u006f\u006e"}}
		_ec.AddPreserveSpaceAttr(&_cg, *_a.Version)
		e.EncodeElement(_a.Version, _cg)
	}
	e.EncodeToken(_gg.EndElement{Name: start.Name})
	return nil
}

type CoreProperties struct{ CT_CoreProperties }

// Validate validates the CT_Keyword and its children
func (_feg *CT_Keyword) Validate() error {
	return _feg.ValidateWithPath("\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064")
}

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_df *CT_CoreProperties) ValidateWithPath(path string) error {
	if _df.Keywords != nil {
		if _ae := _df.Keywords.ValidateWithPath(path + "\u002fK\u0065\u0079\u0077\u006f\u0072\u0064s"); _ae != nil {
			return _ae
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_abb *CT_Keywords) ValidateWithPath(path string) error {
	for _deg, _ecf := range _abb.Value {
		if _gbe := _ecf.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002fV\u0061\u006c\u0075\u0065\u005b\u0025\u0064\u005d", path, _deg)); _gbe != nil {
			return _gbe
		}
	}
	return nil
}

// Validate validates the CT_Keywords and its children
func (_ba *CT_Keywords) Validate() error {
	return _ba.ValidateWithPath("C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073")
}
func (_bge *CoreProperties) MarshalXML(e *_gg.Encoder, start _gg.StartElement) error {
	start.Attr = append(start.Attr, _gg.Attr{Name: _gg.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"})
	start.Attr = append(start.Attr, _gg.Attr{Name: _gg.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0070"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"})
	start.Attr = append(start.Attr, _gg.Attr{Name: _gg.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f"})
	start.Attr = append(start.Attr, _gg.Attr{Name: _gg.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"}, Value: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"})
	start.Attr = append(start.Attr, _gg.Attr{Name: _gg.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063\u0070\u003a\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073"
	return _bge.CT_CoreProperties.MarshalXML(e, start)
}
func (_ca *CT_CoreProperties) UnmarshalXML(d *_gg.Decoder, start _gg.StartElement) error {
_ag:
	for {
		_bg, _be := d.Token()
		if _be != nil {
			return _be
		}
		switch _bb := _bg.(type) {
		case _gg.StartElement:
			switch _bb.Name {
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:
				_ca.Category = new(string)
				if _ece := d.DecodeElement(_ca.Category, &_bb); _ece != nil {
					return _ece
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:
				_ca.ContentStatus = new(string)
				if _eac := d.DecodeElement(_ca.ContentStatus, &_bb); _eac != nil {
					return _eac
				}
			case _gg.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u0063r\u0065\u0061\u0074\u0065\u0064"}:
				_ca.Created = new(_ec.XSDAny)
				if _ce := d.DecodeElement(_ca.Created, &_bb); _ce != nil {
					return _ce
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0063r\u0065\u0061\u0074\u006f\u0072"}:
				_ca.Creator = new(_ec.XSDAny)
				if _ab := d.DecodeElement(_ca.Creator, &_bb); _ab != nil {
					return _ab
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:
				_ca.Description = new(_ec.XSDAny)
				if _ga := d.DecodeElement(_ca.Description, &_bb); _ga != nil {
					return _ga
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:
				_ca.Identifier = new(_ec.XSDAny)
				if _deb := d.DecodeElement(_ca.Identifier, &_bb); _deb != nil {
					return _deb
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:
				_ca.Keywords = NewCT_Keywords()
				if _af := d.DecodeElement(_ca.Keywords, &_bb); _af != nil {
					return _af
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:
				_ca.Language = new(_ec.XSDAny)
				if _gaa := d.DecodeElement(_ca.Language, &_bb); _gaa != nil {
					return _gaa
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:
				_ca.LastModifiedBy = new(string)
				if _cfc := d.DecodeElement(_ca.LastModifiedBy, &_bb); _cfc != nil {
					return _cfc
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:
				_ca.LastPrinted = new(_e.Time)
				if _ded := d.DecodeElement(_ca.LastPrinted, &_bb); _ded != nil {
					return _ded
				}
			case _gg.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:
				_ca.Modified = new(_ec.XSDAny)
				if _ac := d.DecodeElement(_ca.Modified, &_bb); _ac != nil {
					return _ac
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:
				_ca.Revision = new(string)
				if _gb := d.DecodeElement(_ca.Revision, &_bb); _gb != nil {
					return _gb
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0073u\u0062\u006a\u0065\u0063\u0074"}:
				_ca.Subject = new(_ec.XSDAny)
				if _fa := d.DecodeElement(_ca.Subject, &_bb); _fa != nil {
					return _fa
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0074\u0069\u0074l\u0065"}:
				_ca.Title = new(_ec.XSDAny)
				if _deda := d.DecodeElement(_ca.Title, &_bb); _deda != nil {
					return _deda
				}
			case _gg.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0076e\u0072\u0073\u0069\u006f\u006e"}:
				_ca.Version = new(string)
				if _bfg := d.DecodeElement(_ca.Version, &_bb); _bfg != nil {
					return _bfg
				}
			default:
				_ec.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u0020\u0025\u0076", _bb.Name)
				if _ceg := d.Skip(); _ceg != nil {
					return _ceg
				}
			}
		case _gg.EndElement:
			break _ag
		case _gg.CharData:
		}
	}
	return nil
}
func (_cfd *CT_Keywords) MarshalXML(e *_gg.Encoder, start _gg.StartElement) error {
	if _cfd.LangAttr != nil {
		start.Attr = append(start.Attr, _gg.Attr{Name: _gg.Name{Local: "\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"}, Value: _b.Sprintf("\u0025\u0076", *_cfd.LangAttr)})
	}
	e.EncodeToken(start)
	if _cfd.Value != nil {
		_bfgc := _gg.StartElement{Name: _gg.Name{Local: "\u0063\u0070\u003a\u0076\u0061\u006c\u0075\u0065"}}
		for _, _faa := range _cfd.Value {
			e.EncodeElement(_faa, _bfgc)
		}
	}
	e.EncodeToken(_gg.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CoreProperties and its children
func (_gdfb *CoreProperties) Validate() error {
	return _gdfb.ValidateWithPath("\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073")
}
func (_aed *CT_Keyword) MarshalXML(e *_gg.Encoder, start _gg.StartElement) error {
	if _aed.LangAttr != nil {
		start.Attr = append(start.Attr, _gg.Attr{Name: _gg.Name{Local: "\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"}, Value: _b.Sprintf("\u0025\u0076", *_aed.LangAttr)})
	}
	e.EncodeElement(_aed.Content, start)
	e.EncodeToken(_gg.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_ed *CT_Keyword) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_bgef *CoreProperties) ValidateWithPath(path string) error {
	if _fcg := _bgef.CT_CoreProperties.ValidateWithPath(path); _fcg != nil {
		return _fcg
	}
	return nil
}

type CT_Keywords struct {
	LangAttr *string
	Value    []*CT_Keyword
}
type CT_Keyword struct {
	LangAttr *string
	Content  string
}

func NewCoreProperties() *CoreProperties {
	_dcf := &CoreProperties{}
	_dcf.CT_CoreProperties = *NewCT_CoreProperties()
	return _dcf
}
func NewCT_Keyword() *CT_Keyword { _afg := &CT_Keyword{}; return _afg }

type CT_CoreProperties struct {
	Category       *string
	ContentStatus  *string
	Created        *_ec.XSDAny
	Creator        *_ec.XSDAny
	Description    *_ec.XSDAny
	Identifier     *_ec.XSDAny
	Keywords       *CT_Keywords
	Language       *_ec.XSDAny
	LastModifiedBy *string
	LastPrinted    *_e.Time
	Modified       *_ec.XSDAny
	Revision       *string
	Subject        *_ec.XSDAny
	Title          *_ec.XSDAny
	Version        *string
}

func init() {
	_ec.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073", NewCT_CoreProperties)
	_ec.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", NewCT_Keywords)
	_ec.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064", NewCT_Keyword)
	_ec.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073", NewCoreProperties)
}
