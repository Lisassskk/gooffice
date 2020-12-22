package lockedCanvas

import (
	_b "encoding/xml"
	_af "fmt"

	_d "gitee.com/yu_sheng/gooffice"
	_e "gitee.com/yu_sheng/gooffice/schema/soo/dml"
)

// ValidateWithPath validates the LockedCanvas and its children, prefixing error messages with path
func (_bf *LockedCanvas) ValidateWithPath(path string) error {
	if _ag := _bf.CT_GvmlGroupShape.ValidateWithPath(path); _ag != nil {
		return _ag
	}
	return nil
}

// Validate validates the LockedCanvas and its children
func (_ebf *LockedCanvas) Validate() error {
	return _ebf.ValidateWithPath("\u004c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073")
}

type LockedCanvas struct{ _e.CT_GvmlGroupShape }

func (_eb *LockedCanvas) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073"
	return _eb.CT_GvmlGroupShape.MarshalXML(e, start)
}
func (_f *LockedCanvas) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_f.CT_GvmlGroupShape = *_e.NewCT_GvmlGroupShape()
	for {
		_fa, _ed := d.Token()
		if _ed != nil {
			return _af.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u004c\u006f\u0063\u006b\u0065d\u0043\u0061\u006e\u0076\u0061\u0073\u003a\u0020\u0025\u0073", _ed)
		}
		if _ff, _ba := _fa.(_b.EndElement); _ba && _ff.Name == start.Name {
			break
		}
	}
	return nil
}
func NewLockedCanvas() *LockedCanvas {
	_g := &LockedCanvas{}
	_g.CT_GvmlGroupShape = *_e.NewCT_GvmlGroupShape()
	return _g
}
func init() {
	_d.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073", "\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073", NewLockedCanvas)
}
