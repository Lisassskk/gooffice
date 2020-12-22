package terms

import (
	_b "encoding/xml"
	_g "fmt"

	_ba "gitee.com/gooffice/gooffice"
	_a "gitee.com/gooffice/gooffice/schema/purl.org/dc/elements"
)

func NewPeriod() *Period { _be := &Period{}; return _be }
func (_db *ElementsAndRefinementsGroup) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	if _db.Choice != nil {
		for _, _df := range _db.Choice {
			_df.MarshalXML(e, _b.StartElement{})
		}
	}
	return nil
}

// Validate validates the Point and its children
func (_daf *Point) Validate() error { return _daf.ValidateWithPath("\u0050\u006f\u0069n\u0074") }
func (_bcc *DDC) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_ecf, _dc := d.Token()
		if _dc != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073", _dc)
		}
		if _geb, _bab := _ecf.(_b.EndElement); _bab && _geb.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the UDC and its children
func (_bfb *UDC) Validate() error { return _bfb.ValidateWithPath("\u0055\u0044\u0043") }

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_bd *Box) ValidateWithPath(path string) error { return nil }
func NewLCSH() *LCSH                                { _eff := &LCSH{}; return _eff }
func (_bbg *URI) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_ddd, _ffb := d.Token()
		if _ffb != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073", _ffb)
		}
		if _cbf, _cfd := _ddd.(_b.EndElement); _cfd && _cbf.Name == start.Name {
			break
		}
	}
	return nil
}
func (_gca *ISO3166) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0049S\u004f\u0033\u0031\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// Validate validates the DDC and its children
func (_bag *DDC) Validate() error { return _bag.ValidateWithPath("\u0044\u0044\u0043") }

// Validate validates the ISO639_2 and its children
func (_dbf *ISO639_2) Validate() error {
	return _dbf.ValidateWithPath("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032")
}
func NewElementsAndRefinementsGroupChoice() *ElementsAndRefinementsGroupChoice {
	_ecg := &ElementsAndRefinementsGroupChoice{}
	return _ecg
}
func NewDCMIType() *DCMIType { _ea := &DCMIType{}; return _ea }

type MESH struct{}

func NewMESH() *MESH { _dgd := &MESH{}; return _dgd }
func (_caa *LCC) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_gcf, _gga := d.Token()
		if _gga != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073", _gga)
		}
		if _bcb, _fcce := _gcf.(_b.EndElement); _fcce && _bcb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_cdc *Period) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_dde, _fbe := d.Token()
		if _fbe != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073", _fbe)
		}
		if _gdcb, _eac := _dde.(_b.EndElement); _eac && _gdcb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ede *ElementsAndRefinementsGroup) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
_eed:
	for {
		_fdd, _ae := d.Token()
		if _ae != nil {
			return _ae
		}
		switch _dd := _fdd.(type) {
		case _b.StartElement:
			switch _dd.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_cc := NewElementsAndRefinementsGroupChoice()
				if _efd := d.DecodeElement(&_cc.Any, &_dd); _efd != nil {
					return _efd
				}
				_ede.Choice = append(_ede.Choice, _cc)
			default:
				_ba.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076", _dd.Name)
				if _fde := d.Skip(); _fde != nil {
					return _fde
				}
			}
		case _b.EndElement:
			break _eed
		case _b.CharData:
		}
	}
	return nil
}

type DCMIType struct{}

func (_dff *Point) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0050\u006f\u0069n\u0074"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_gegb *LCSH) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_bac *UDC) ValidateWithPath(path string) error { return nil }
func (_cfga *W3CDTF) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_gea, _dfg := d.Token()
		if _dfg != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073", _dfg)
		}
		if _egc, _egd := _gea.(_b.EndElement); _egd && _egc.Name == start.Name {
			break
		}
	}
	return nil
}
func NewElementsAndRefinementsGroup() *ElementsAndRefinementsGroup {
	_gb := &ElementsAndRefinementsGroup{}
	return _gb
}
func NewBox() *Box { _e := &Box{}; return _e }

// Validate validates the Period and its children
func (_ggb *Period) Validate() error {
	return _ggb.ValidateWithPath("\u0050\u0065\u0072\u0069\u006f\u0064")
}

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_gg *ElementsAndRefinementsGroup) ValidateWithPath(path string) error {
	for _edfa, _gdc := range _gg.Choice {
		if _bb := _gdc.ValidateWithPath(_g.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _edfa)); _bb != nil {
			return _bb
		}
	}
	return nil
}

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_de *ElementOrRefinementContainer) ValidateWithPath(path string) error {
	for _ee, _fd := range _de.Choice {
		if _eef := _fd.ValidateWithPath(_g.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _ee)); _eef != nil {
			return _eef
		}
	}
	return nil
}
func NewISO639_2() *ISO639_2 { _gcae := &ISO639_2{}; return _gcae }

type Box struct{}

func NewISO3166() *ISO3166 { _gee := &ISO3166{}; return _gee }
func (_ca *ElementOrRefinementContainer) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
_da:
	for {
		_ef, _edf := d.Token()
		if _edf != nil {
			return _edf
		}
		switch _acg := _ef.(type) {
		case _b.StartElement:
			switch _acg.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_baa := NewElementsAndRefinementsGroupChoice()
				if _cae := d.DecodeElement(&_baa.Any, &_acg); _cae != nil {
					return _cae
				}
				_ca.Choice = append(_ca.Choice, _baa)
			default:
				_ba.Log("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076", _acg.Name)
				if _dcf := d.Skip(); _dcf != nil {
					return _dcf
				}
			}
		case _b.EndElement:
			break _da
		case _b.CharData:
		}
	}
	return nil
}

type UDC struct{}

func (_d *DDC) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0044\u0044\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func NewIMT() *IMT { _bf := &IMT{}; return _bf }

type DDC struct{}
type ISO3166 struct{}

// Validate validates the RFC1766 and its children
func (_ddc *RFC1766) Validate() error {
	return _ddc.ValidateWithPath("\u0052F\u0043\u0031\u0037\u0036\u0036")
}

type IMT struct{}

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_gfe *MESH) ValidateWithPath(path string) error { return nil }
func (_bgb *MESH) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u004d\u0045\u0053\u0048"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_bcg *Period) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0050\u0065\u0072\u0069\u006f\u0064"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func NewRFC3066() *RFC3066 { _gag := &RFC3066{}; return _gag }

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_eee *ISO3166) ValidateWithPath(path string) error { return nil }

// Validate validates the LCC and its children
func (_ce *LCC) Validate() error { return _ce.ValidateWithPath("\u004c\u0043\u0043") }

type TGN struct{}

// Validate validates the W3CDTF and its children
func (_gfb *W3CDTF) Validate() error {
	return _gfb.ValidateWithPath("\u0057\u0033\u0043\u0044\u0054\u0046")
}

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_gdd *LCC) ValidateWithPath(path string) error { return nil }
func (_bda *LCC) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u004c\u0043\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_bc *DCMIType) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_eae *ISO639_2) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_bcf *ElementsAndRefinementsGroupChoice) ValidateWithPath(path string) error {
	for _gec, _ece := range _bcf.Any {
		if _abb := _ece.ValidateWithPath(_g.Sprintf("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d", path, _gec)); _abb != nil {
			return _abb
		}
	}
	return nil
}

type Point struct{}

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_ebf *DCMIType) ValidateWithPath(path string) error { return nil }
func (_efb *ISO639_2) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_fga, _aga := d.Token()
		if _aga != nil {
			return _g.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073", _aga)
		}
		if _bgfg, _edb := _fga.(_b.EndElement); _edb && _bgfg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_bbd *IMT) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_bge, _aee := d.Token()
		if _aee != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073", _aee)
		}
		if _cg, _cd := _bge.(_b.EndElement); _cd && _cg.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the LCSH and its children
func (_cfe *LCSH) Validate() error { return _cfe.ValidateWithPath("\u004c\u0043\u0053\u0048") }
func (_ggg *MESH) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_dbbc, _ccb := d.Token()
		if _ccb != nil {
			return _g.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073", _ccb)
		}
		if _dbc, _feb := _dbbc.(_b.EndElement); _feb && _dbc.Name == start.Name {
			break
		}
	}
	return nil
}

type LCSH struct{}

func (_gd *Box) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_eb, _c := d.Token()
		if _c != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073", _c)
		}
		if _ac, _ag := _eb.(_b.EndElement); _ag && _ac.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_dba *ISO639_2) ValidateWithPath(path string) error { return nil }
func (_cf *DCMIType) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_bad, _ge := d.Token()
		if _ge != nil {
			return _g.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073", _ge)
		}
		if _ec, _ged := _bad.(_b.EndElement); _ged && _ec.Name == start.Name {
			break
		}
	}
	return nil
}

type ElementsAndRefinementsGroupChoice struct{ Any []*_a.Any }

func (_aec *TGN) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0054\u0047\u004e"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_bbdb *LCSH) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u004c\u0043\u0053\u0048"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func NewRFC1766() *RFC1766 { _eaee := &RFC1766{}; return _eaee }
func NewUDC() *UDC         { _ecb := &UDC{}; return _ecb }
func (_gc *IMT) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0049\u004d\u0054"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_bbdbd *RFC1766) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0052F\u0043\u0031\u0037\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

type ElementOrRefinementContainer struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

func (_dab *UDC) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0055\u0044\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// Validate validates the ISO3166 and its children
func (_edc *ISO3166) Validate() error {
	return _edc.ValidateWithPath("\u0049S\u004f\u0033\u0031\u0036\u0036")
}

type RFC1766 struct{}

func (_geg *ElementsAndRefinementsGroupChoice) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
_ebfd:
	for {
		_ab, _gbd := d.Token()
		if _gbd != nil {
			return _gbd
		}
		switch _aed := _ab.(type) {
		case _b.StartElement:
			switch _aed.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_aca := _a.NewAny()
				if _cce := d.DecodeElement(_aca, &_aed); _cce != nil {
					return _cce
				}
				_geg.Any = append(_geg.Any, _aca)
			default:
				_ba.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076", _aed.Name)
				if _dfd := d.Skip(); _dfd != nil {
					return _dfd
				}
			}
		case _b.EndElement:
			break _ebfd
		case _b.CharData:
		}
	}
	return nil
}
func (_gf *ElementsAndRefinementsGroupChoice) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	if _gf.Any != nil {
		_fe := _b.StartElement{Name: _b.Name{Local: "\u0064\u0063\u003a\u0061\u006e\u0079"}}
		for _, _dg := range _gf.Any {
			e.EncodeElement(_dg, _fe)
		}
	}
	return nil
}
func (_bgea *URI) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0055\u0052\u0049"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_bef *Point) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_ebfg, _gef := d.Token()
		if _gef != nil {
			return _g.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073", _gef)
		}
		if _bfg, _eaga := _ebfg.(_b.EndElement); _eaga && _bfg.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_cfg *Point) ValidateWithPath(path string) error { return nil }

// Validate validates the MESH and its children
func (_gae *MESH) Validate() error { return _gae.ValidateWithPath("\u004d\u0045\u0053\u0048") }

type LCC struct{}
type URI struct{}

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_ga *IMT) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_eca *URI) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_fcd *RFC3066) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_bg *DDC) ValidateWithPath(path string) error { return nil }
func NewElementOrRefinementContainer() *ElementOrRefinementContainer {
	_fc := &ElementOrRefinementContainer{}
	return _fc
}

// Validate validates the ElementOrRefinementContainer and its children
func (_fcc *ElementOrRefinementContainer) Validate() error {
	return _fcc.ValidateWithPath("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072")
}

// Validate validates the ElementsAndRefinementsGroup and its children
func (_dce *ElementsAndRefinementsGroup) Validate() error {
	return _dce.ValidateWithPath("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070")
}

// Validate validates the Box and its children
func (_fg *Box) Validate() error { return _fg.ValidateWithPath("\u0042\u006f\u0078") }
func NewTGN() *TGN               { _gdae := &TGN{}; return _gdae }

// Validate validates the RFC3066 and its children
func (_gda *RFC3066) Validate() error {
	return _gda.ValidateWithPath("\u0052F\u0043\u0033\u0030\u0036\u0036")
}

type RFC3066 struct{}

// Validate validates the IMT and its children
func (_caee *IMT) Validate() error { return _caee.ValidateWithPath("\u0049\u004d\u0054") }

// Validate validates the TGN and its children
func (_gbg *TGN) Validate() error { return _gbg.ValidateWithPath("\u0054\u0047\u004e") }

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_agc *RFC1766) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_gegbf *Period) ValidateWithPath(path string) error { return nil }
func (_ed *Box) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0042\u006f\u0078"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

type Period struct{}

func (_ff *ElementOrRefinementContainer) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072"
	e.EncodeToken(start)
	if _ff.Choice != nil {
		for _, _bdd := range _ff.Choice {
			_bdd.MarshalXML(e, _b.StartElement{})
		}
	}
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_dcg *W3CDTF) ValidateWithPath(path string) error { return nil }
func (_gcg *LCSH) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_eg, _acc := d.Token()
		if _acc != nil {
			return _g.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073", _acc)
		}
		if _bcfd, _ddb := _eg.(_b.EndElement); _ddb && _bcfd.Name == start.Name {
			break
		}
	}
	return nil
}
func NewURI() *URI { _cdf := &URI{}; return _cdf }
func (_ddf *W3CDTF) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0057\u0033\u0043\u0044\u0054\u0046"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_ddbd *RFC3066) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Name.Local = "\u0052F\u0043\u0033\u0030\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}
func (_ggf *UDC) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_eba, _gbb := d.Token()
		if _gbb != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073", _gbb)
		}
		if _dafd, _cb := _eba.(_b.EndElement); _cb && _dafd.Name == start.Name {
			break
		}
	}
	return nil
}
func NewPoint() *Point { _dgc := &Point{}; return _dgc }

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_bed *TGN) ValidateWithPath(path string) error { return nil }
func (_beg *TGN) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_gegd, _effc := d.Token()
		if _effc != nil {
			return _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073", _effc)
		}
		if _cfc, _ebg := _gegd.(_b.EndElement); _ebg && _cfc.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the DCMIType and its children
func (_ad *DCMIType) Validate() error {
	return _ad.ValidateWithPath("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065")
}
func NewW3CDTF() *W3CDTF { _eeb := &W3CDTF{}; return _eeb }

type ElementsAndRefinementsGroup struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

func (_ffe *ISO3166) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_edee, _dbb := d.Token()
		if _dbb != nil {
			return _g.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073", _dbb)
		}
		if _bgf, _gdb := _edee.(_b.EndElement); _gdb && _bgf.Name == start.Name {
			break
		}
	}
	return nil
}
func NewLCC() *LCC { _ffee := &LCC{}; return _ffee }

type W3CDTF struct{}

func (_fbc *RFC3066) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_gcc, _acag := d.Token()
		if _acag != nil {
			return _g.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073", _acag)
		}
		if _bdg, _bcgf := _gcc.(_b.EndElement); _bcgf && _bdg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_cea *RFC1766) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for {
		_bbc, _dgg := d.Token()
		if _dgg != nil {
			return _g.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073", _dgg)
		}
		if _ffg, _adc := _bbc.(_b.EndElement); _adc && _ffg.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the URI and its children
func (_cda *URI) Validate() error { return _cda.ValidateWithPath("\u0055\u0052\u0049") }

type ISO639_2 struct{}

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_gde *ElementsAndRefinementsGroupChoice) Validate() error {
	return _gde.ValidateWithPath("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065")
}
func NewDDC() *DDC { _eag := &DDC{}; return _eag }
func init() {
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004c\u0043\u0053\u0048", NewLCSH)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004d\u0045\u0053\u0048", NewMESH)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0044\u0044\u0043", NewDDC)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004c\u0043\u0043", NewLCC)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0055\u0044\u0043", NewUDC)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0050\u0065\u0072\u0069\u006f\u0064", NewPeriod)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0057\u0033\u0043\u0044\u0054\u0046", NewW3CDTF)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065", NewDCMIType)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049\u004d\u0054", NewIMT)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0055\u0052\u0049", NewURI)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032", NewISO639_2)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0052F\u0043\u0031\u0037\u0036\u0036", NewRFC1766)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0052F\u0043\u0033\u0030\u0036\u0036", NewRFC3066)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0050\u006f\u0069n\u0074", NewPoint)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049S\u004f\u0033\u0031\u0036\u0036", NewISO3166)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0042\u006f\u0078", NewBox)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0054\u0047\u004e", NewTGN)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072", NewElementOrRefinementContainer)
	_ba.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070", NewElementsAndRefinementsGroup)
}
