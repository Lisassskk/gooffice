package elements

import (
	_e "encoding/xml"
	_f "fmt"

	_ff "gitee.com/gooffice/gooffice"
)

func (_c *Any) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return _c.SimpleLiteral.MarshalXML(e, start)
}

type ElementsGroupChoice struct{ Any []*Any }

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_gd *ElementsGroup) ValidateWithPath(path string) error {
	for _dbg, _edf := range _gd.Choice {
		if _afg := _edf.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _dbg)); _afg != nil {
			return _afg
		}
	}
	return nil
}

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_ec *SimpleLiteral) ValidateWithPath(path string) error { return nil }
func (_adg *ElementsGroupChoice) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	if _adg.Any != nil {
		_dac := _e.StartElement{Name: _e.Name{Local: "\u0064\u0063\u003a\u0061\u006e\u0079"}}
		for _, _dc := range _adg.Any {
			e.EncodeElement(_dc, _dac)
		}
	}
	return nil
}

type ElementContainer struct{ Choice []*ElementsGroupChoice }

func (_ef *ElementsGroup) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	if _ef.Choice != nil {
		for _, _cea := range _ef.Choice {
			_cea.MarshalXML(e, _e.StartElement{})
		}
	}
	return nil
}
func NewElementsGroupChoice() *ElementsGroupChoice { _ga := &ElementsGroupChoice{}; return _ga }

type ElementsGroup struct{ Choice []*ElementsGroupChoice }

func (_afgb *SimpleLiteral) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_e.EndElement{Name: start.Name})
	return nil
}

// Validate validates the ElementsGroup and its children
func (_ccf *ElementsGroup) Validate() error {
	return _ccf.ValidateWithPath("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070")
}
func NewAny() *Any { _g := &Any{}; _g.SimpleLiteral = *NewSimpleLiteral(); return _g }

type SimpleLiteral struct{}

// Validate validates the ElementsGroupChoice and its children
func (_gf *ElementsGroupChoice) Validate() error {
	return _gf.ValidateWithPath("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065")
}

type Any struct{ SimpleLiteral }

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_eg *Any) ValidateWithPath(path string) error {
	if _ag := _eg.SimpleLiteral.ValidateWithPath(path); _ag != nil {
		return _ag
	}
	return nil
}

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_ada *ElementsGroupChoice) ValidateWithPath(path string) error {
	for _gdg, _cef := range _ada.Any {
		if _gcd := _cef.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d", path, _gdg)); _gcd != nil {
			return _gcd
		}
	}
	return nil
}
func (_a *Any) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_a.SimpleLiteral = *NewSimpleLiteral()
	for {
		_be, _ce := d.Token()
		if _ce != nil {
			return _f.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073", _ce)
		}
		if _ca, _cd := _be.(_e.EndElement); _cd && _ca.Name == start.Name {
			break
		}
	}
	return nil
}
func (_dab *ElementsGroupChoice) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
_daa:
	for {
		_gac, _cdc := d.Token()
		if _cdc != nil {
			return _cdc
		}
		switch _cag := _gac.(type) {
		case _e.StartElement:
			switch _cag.Name {
			case _e.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_fg := NewAny()
				if _dcd := d.DecodeElement(_fg, &_cag); _dcd != nil {
					return _dcd
				}
				_dab.Any = append(_dab.Any, _fg)
			default:
				_ff.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _cag.Name)
				if _ee := d.Skip(); _ee != nil {
					return _ee
				}
			}
		case _e.EndElement:
			break _daa
		case _e.CharData:
		}
	}
	return nil
}
func (_fb *ElementContainer) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	start.Name.Local = "\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072"
	e.EncodeToken(start)
	if _fb.Choice != nil {
		for _, _bec := range _fb.Choice {
			_bec.MarshalXML(e, _e.StartElement{})
		}
	}
	e.EncodeToken(_e.EndElement{Name: start.Name})
	return nil
}
func (_cee *SimpleLiteral) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	for {
		_ae, _bbc := d.Token()
		if _bbc != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s", _bbc)
		}
		if _bef, _bg := _ae.(_e.EndElement); _bg && _bef.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the SimpleLiteral and its children
func (_bbd *SimpleLiteral) Validate() error {
	return _bbd.ValidateWithPath("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c")
}
func NewElementContainer() *ElementContainer { _d := &ElementContainer{}; return _d }

// Validate validates the ElementContainer and its children
func (_gc *ElementContainer) Validate() error {
	return _gc.ValidateWithPath("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072")
}
func NewSimpleLiteral() *SimpleLiteral { _eb := &SimpleLiteral{}; return _eb }
func (_fa *ElementsGroup) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
_da:
	for {
		_gca, _daf := d.Token()
		if _daf != nil {
			return _daf
		}
		switch _fc := _gca.(type) {
		case _e.StartElement:
			switch _fc.Name {
			case _e.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_fcb := NewElementsGroupChoice()
				if _dd := d.DecodeElement(&_fcb.Any, &_fc); _dd != nil {
					return _dd
				}
				_fa.Choice = append(_fa.Choice, _fcb)
			default:
				_ff.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076", _fc.Name)
				if _af := d.Skip(); _af != nil {
					return _af
				}
			}
		case _e.EndElement:
			break _da
		case _e.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_dbe *ElementContainer) ValidateWithPath(path string) error {
	for _fee, _edg := range _dbe.Choice {
		if _agd := _edg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _fee)); _agd != nil {
			return _agd
		}
	}
	return nil
}
func (_ffb *ElementContainer) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
_db:
	for {
		_cb, _ad := d.Token()
		if _ad != nil {
			return _ad
		}
		switch _de := _cb.(type) {
		case _e.StartElement:
			switch _de.Name {
			case _e.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_fe := NewElementsGroupChoice()
				if _ed := d.DecodeElement(&_fe.Any, &_de); _ed != nil {
					return _ed
				}
				_ffb.Choice = append(_ffb.Choice, _fe)
			default:
				_ff.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v", _de.Name)
				if _ab := d.Skip(); _ab != nil {
					return _ab
				}
			}
		case _e.EndElement:
			break _db
		case _e.CharData:
		}
	}
	return nil
}
func NewElementsGroup() *ElementsGroup { _cc := &ElementsGroup{}; return _cc }

// Validate validates the Any and its children
func (_ea *Any) Validate() error { return _ea.ValidateWithPath("\u0041\u006e\u0079") }
func init() {
	_ff.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c", NewSimpleLiteral)
	_ff.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072", NewElementContainer)
	_ff.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0061\u006e\u0079", NewAny)
	_ff.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070", NewElementsGroup)
}
