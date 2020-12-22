

package excel

import (
	_ef "encoding/xml"
	_d "fmt"
	_f "gitee.com/yu_sheng/gooffice"
	_b "gitee.com/yu_sheng/gooffice/schema/soo/ofc/sharedTypes"
)

func (_gbg ST_ObjectType) MarshalXMLAttr(name _ef.Name) (_ef.Attr, error) {
	_eage := _ef.Attr{}
	_eage.Name = name
	switch _gbg {
	case ST_ObjectTypeUnset:
		_eage.Value = ""
	case ST_ObjectTypeButton:
		_eage.Value = "\u0042\u0075\u0074\u0074\u006f\u006e"
	case ST_ObjectTypeCheckbox:
		_eage.Value = "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078"
	case ST_ObjectTypeDialog:
		_eage.Value = "\u0044\u0069\u0061\u006c\u006f\u0067"
	case ST_ObjectTypeDrop:
		_eage.Value = "\u0044\u0072\u006f\u0070"
	case ST_ObjectTypeEdit:
		_eage.Value = "\u0045\u0064\u0069\u0074"
	case ST_ObjectTypeGBox:
		_eage.Value = "\u0047\u0042\u006f\u0078"
	case ST_ObjectTypeLabel:
		_eage.Value = "\u004c\u0061\u0062e\u006c"
	case ST_ObjectTypeLineA:
		_eage.Value = "\u004c\u0069\u006ee\u0041"
	case ST_ObjectTypeList:
		_eage.Value = "\u004c\u0069\u0073\u0074"
	case ST_ObjectTypeMovie:
		_eage.Value = "\u004d\u006f\u0076i\u0065"
	case ST_ObjectTypeNote:
		_eage.Value = "\u004e\u006f\u0074\u0065"
	case ST_ObjectTypePict:
		_eage.Value = "\u0050\u0069\u0063\u0074"
	case ST_ObjectTypeRadio:
		_eage.Value = "\u0052\u0061\u0064i\u006f"
	case ST_ObjectTypeRectA:
		_eage.Value = "\u0052\u0065\u0063t\u0041"
	case ST_ObjectTypeScroll:
		_eage.Value = "\u0053\u0063\u0072\u006f\u006c\u006c"
	case ST_ObjectTypeSpin:
		_eage.Value = "\u0053\u0070\u0069\u006e"
	case ST_ObjectTypeShape:
		_eage.Value = "\u0053\u0068\u0061p\u0065"
	case ST_ObjectTypeGroup:
		_eage.Value = "\u0047\u0072\u006fu\u0070"
	case ST_ObjectTypeRect:
		_eage.Value = "\u0052\u0065\u0063\u0074"
	}
	return _eage, nil
}

// ValidateWithPath validates the CT_ClientData and its children, prefixing error messages with path
func (_aaf *CT_ClientData) ValidateWithPath(path string) error {
	if _aaf.ObjectTypeAttr == ST_ObjectTypeUnset {
		return _d.Errorf("\u0025\u0073\u002f\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020\u0061\u0020\u006da\u006e\u0064\u0061\u0074\u006fr\u0079\u0020f\u0069\u0065\u006c\u0064", path)
	}
	if _fga := _aaf.ObjectTypeAttr.ValidateWithPath(path + "\u002fO\u0062j\u0065\u0063\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _fga != nil {
		return _fga
	}
	if _fee := _aaf.MoveWithCells.ValidateWithPath(path + "\u002f\u004d\u006f\u0076\u0065\u0057\u0069\u0074\u0068C\u0065\u006c\u006c\u0073"); _fee != nil {
		return _fee
	}
	if _adf := _aaf.SizeWithCells.ValidateWithPath(path + "\u002f\u0053\u0069\u007a\u0065\u0057\u0069\u0074\u0068C\u0065\u006c\u006c\u0073"); _adf != nil {
		return _adf
	}
	if _gag := _aaf.Locked.ValidateWithPath(path + "\u002fL\u006f\u0063\u006b\u0065\u0064"); _gag != nil {
		return _gag
	}
	if _ceb := _aaf.DefaultSize.ValidateWithPath(path + "\u002f\u0044\u0065f\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"); _ceb != nil {
		return _ceb
	}
	if _geg := _aaf.PrintObject.ValidateWithPath(path + "\u002f\u0050\u0072i\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"); _geg != nil {
		return _geg
	}
	if _gff := _aaf.Disabled.ValidateWithPath(path + "\u002fD\u0069\u0073\u0061\u0062\u006c\u0065d"); _gff != nil {
		return _gff
	}
	if _gae := _aaf.AutoFill.ValidateWithPath(path + "\u002fA\u0075\u0074\u006f\u0046\u0069\u006cl"); _gae != nil {
		return _gae
	}
	if _aac := _aaf.AutoLine.ValidateWithPath(path + "\u002fA\u0075\u0074\u006f\u004c\u0069\u006ee"); _aac != nil {
		return _aac
	}
	if _bag := _aaf.AutoPict.ValidateWithPath(path + "\u002fA\u0075\u0074\u006f\u0050\u0069\u0063t"); _bag != nil {
		return _bag
	}
	if _abc := _aaf.LockText.ValidateWithPath(path + "\u002fL\u006f\u0063\u006b\u0054\u0065\u0078t"); _abc != nil {
		return _abc
	}
	if _dcf := _aaf.JustLastX.ValidateWithPath(path + "\u002f\u004a\u0075\u0073\u0074\u004c\u0061\u0073\u0074\u0058"); _dcf != nil {
		return _dcf
	}
	if _gcga := _aaf.SecretEdit.ValidateWithPath(path + "/\u0053\u0065\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"); _gcga != nil {
		return _gcga
	}
	if _cdag := _aaf.Default.ValidateWithPath(path + "\u002f\u0044\u0065\u0066\u0061\u0075\u006c\u0074"); _cdag != nil {
		return _cdag
	}
	if _fbg := _aaf.Help.ValidateWithPath(path + "\u002f\u0048\u0065l\u0070"); _fbg != nil {
		return _fbg
	}
	if _fgae := _aaf.Cancel.ValidateWithPath(path + "\u002fC\u0061\u006e\u0063\u0065\u006c"); _fgae != nil {
		return _fgae
	}
	if _badde := _aaf.Dismiss.ValidateWithPath(path + "\u002f\u0044\u0069\u0073\u006d\u0069\u0073\u0073"); _badde != nil {
		return _badde
	}
	if _bca := _aaf.Visible.ValidateWithPath(path + "\u002f\u0056\u0069\u0073\u0069\u0062\u006c\u0065"); _bca != nil {
		return _bca
	}
	if _eda := _aaf.RowHidden.ValidateWithPath(path + "\u002f\u0052\u006f\u0077\u0048\u0069\u0064\u0064\u0065\u006e"); _eda != nil {
		return _eda
	}
	if _fec := _aaf.ColHidden.ValidateWithPath(path + "\u002f\u0043\u006f\u006c\u0048\u0069\u0064\u0064\u0065\u006e"); _fec != nil {
		return _fec
	}
	if _dcfe := _aaf.MultiLine.ValidateWithPath(path + "\u002f\u004d\u0075\u006c\u0074\u0069\u004c\u0069\u006e\u0065"); _dcfe != nil {
		return _dcfe
	}
	if _ggc := _aaf.VScroll.ValidateWithPath(path + "\u002f\u0056\u0053\u0063\u0072\u006f\u006c\u006c"); _ggc != nil {
		return _ggc
	}
	if _dcd := _aaf.ValidIds.ValidateWithPath(path + "\u002fV\u0061\u006c\u0069\u0064\u0049\u0064s"); _dcd != nil {
		return _dcd
	}
	if _aced := _aaf.NoThreeD2.ValidateWithPath(path + "\u002f\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044\u0032"); _aced != nil {
		return _aced
	}
	if _cbe := _aaf.Colored.ValidateWithPath(path + "\u002f\u0043\u006f\u006c\u006f\u0072\u0065\u0064"); _cbe != nil {
		return _cbe
	}
	if _fbgf := _aaf.NoThreeD.ValidateWithPath(path + "\u002fN\u006f\u0054\u0068\u0072\u0065\u0065D"); _fbgf != nil {
		return _fbgf
	}
	if _fcda := _aaf.FirstButton.ValidateWithPath(path + "\u002f\u0046\u0069r\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"); _fcda != nil {
		return _fcda
	}
	if _dcfg := _aaf.Horiz.ValidateWithPath(path + "\u002f\u0048\u006f\u0072\u0069\u007a"); _dcfg != nil {
		return _dcfg
	}
	if _fdg := _aaf.MapOCX.ValidateWithPath(path + "\u002fM\u0061\u0070\u004f\u0043\u0058"); _fdg != nil {
		return _fdg
	}
	if _fcc := _aaf.Camera.ValidateWithPath(path + "\u002fC\u0061\u006d\u0065\u0072\u0061"); _fcc != nil {
		return _fcc
	}
	if _af := _aaf.RecalcAlways.ValidateWithPath(path + "\u002f\u0052\u0065\u0063\u0061\u006c\u0063\u0041\u006c\u0077\u0061\u0079\u0073"); _af != nil {
		return _af
	}
	if _bgaf := _aaf.AutoScale.ValidateWithPath(path + "\u002f\u0041\u0075\u0074\u006f\u0053\u0063\u0061\u006c\u0065"); _bgaf != nil {
		return _bgaf
	}
	if _ccgf := _aaf.DDE.ValidateWithPath(path + "\u002f\u0044\u0044\u0045"); _ccgf != nil {
		return _ccgf
	}
	if _gegc := _aaf.UIObj.ValidateWithPath(path + "\u002f\u0055\u0049\u004f\u0062\u006a"); _gegc != nil {
		return _gegc
	}
	return nil
}
func (_bea *ClientData) MarshalXML(e *_ef.Encoder, start _ef.StartElement) error {
	start.Attr = append(start.Attr, _ef.Attr{Name: _ef.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Attr = append(start.Attr, _ef.Attr{Name: _ef.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Attr = append(start.Attr, _ef.Attr{Name: _ef.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u003a\u0043l\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"
	return _bea.CT_ClientData.MarshalXML(e, start)
}

type ST_ObjectType byte

// Validate validates the ClientData and its children
func (_ggeb *ClientData) Validate() error {
	return _ggeb.ValidateWithPath("\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061")
}
func NewCT_ClientData() *CT_ClientData {
	_c := &CT_ClientData{}
	_c.ObjectTypeAttr = ST_ObjectType(1)
	return _c
}
func (_ggdac *ST_ObjectType) UnmarshalXML(d *_ef.Decoder, start _ef.StartElement) error {
	_eae, _dgc := d.Token()
	if _dgc != nil {
		return _dgc
	}
	if _ada, _gdc := _eae.(_ef.EndElement); _gdc && _ada.Name == start.Name {
		*_ggdac = 1
		return nil
	}
	if _fbcg, _abgg := _eae.(_ef.CharData); !_abgg {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eae)
	} else {
		switch string(_fbcg) {
		case "":
			*_ggdac = 0
		case "\u0042\u0075\u0074\u0074\u006f\u006e":
			*_ggdac = 1
		case "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078":
			*_ggdac = 2
		case "\u0044\u0069\u0061\u006c\u006f\u0067":
			*_ggdac = 3
		case "\u0044\u0072\u006f\u0070":
			*_ggdac = 4
		case "\u0045\u0064\u0069\u0074":
			*_ggdac = 5
		case "\u0047\u0042\u006f\u0078":
			*_ggdac = 6
		case "\u004c\u0061\u0062e\u006c":
			*_ggdac = 7
		case "\u004c\u0069\u006ee\u0041":
			*_ggdac = 8
		case "\u004c\u0069\u0073\u0074":
			*_ggdac = 9
		case "\u004d\u006f\u0076i\u0065":
			*_ggdac = 10
		case "\u004e\u006f\u0074\u0065":
			*_ggdac = 11
		case "\u0050\u0069\u0063\u0074":
			*_ggdac = 12
		case "\u0052\u0061\u0064i\u006f":
			*_ggdac = 13
		case "\u0052\u0065\u0063t\u0041":
			*_ggdac = 14
		case "\u0053\u0063\u0072\u006f\u006c\u006c":
			*_ggdac = 15
		case "\u0053\u0070\u0069\u006e":
			*_ggdac = 16
		case "\u0053\u0068\u0061p\u0065":
			*_ggdac = 17
		case "\u0047\u0072\u006fu\u0070":
			*_ggdac = 18
		case "\u0052\u0065\u0063\u0074":
			*_ggdac = 19
		}
	}
	_eae, _dgc = d.Token()
	if _dgc != nil {
		return _dgc
	}
	if _eaf, _dgce := _eae.(_ef.EndElement); _dgce && _eaf.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eae)
}

// Validate validates the CT_ClientData and its children
func (_bbd *CT_ClientData) Validate() error {
	return _bbd.ValidateWithPath("\u0043\u0054\u005f\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061")
}
func (_cbd ST_ObjectType) String() string {
	switch _cbd {
	case 0:
		return ""
	case 1:
		return "\u0042\u0075\u0074\u0074\u006f\u006e"
	case 2:
		return "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078"
	case 3:
		return "\u0044\u0069\u0061\u006c\u006f\u0067"
	case 4:
		return "\u0044\u0072\u006f\u0070"
	case 5:
		return "\u0045\u0064\u0069\u0074"
	case 6:
		return "\u0047\u0042\u006f\u0078"
	case 7:
		return "\u004c\u0061\u0062e\u006c"
	case 8:
		return "\u004c\u0069\u006ee\u0041"
	case 9:
		return "\u004c\u0069\u0073\u0074"
	case 10:
		return "\u004d\u006f\u0076i\u0065"
	case 11:
		return "\u004e\u006f\u0074\u0065"
	case 12:
		return "\u0050\u0069\u0063\u0074"
	case 13:
		return "\u0052\u0061\u0064i\u006f"
	case 14:
		return "\u0052\u0065\u0063t\u0041"
	case 15:
		return "\u0053\u0063\u0072\u006f\u006c\u006c"
	case 16:
		return "\u0053\u0070\u0069\u006e"
	case 17:
		return "\u0053\u0068\u0061p\u0065"
	case 18:
		return "\u0047\u0072\u006fu\u0070"
	case 19:
		return "\u0052\u0065\u0063\u0074"
	}
	return ""
}

// ValidateWithPath validates the ClientData and its children, prefixing error messages with path
func (_cdeg *ClientData) ValidateWithPath(path string) error {
	if _gccc := _cdeg.CT_ClientData.ValidateWithPath(path); _gccc != nil {
		return _gccc
	}
	return nil
}
func (_gcdg *ST_ObjectType) UnmarshalXMLAttr(attr _ef.Attr) error {
	switch attr.Value {
	case "":
		*_gcdg = 0
	case "\u0042\u0075\u0074\u0074\u006f\u006e":
		*_gcdg = 1
	case "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078":
		*_gcdg = 2
	case "\u0044\u0069\u0061\u006c\u006f\u0067":
		*_gcdg = 3
	case "\u0044\u0072\u006f\u0070":
		*_gcdg = 4
	case "\u0045\u0064\u0069\u0074":
		*_gcdg = 5
	case "\u0047\u0042\u006f\u0078":
		*_gcdg = 6
	case "\u004c\u0061\u0062e\u006c":
		*_gcdg = 7
	case "\u004c\u0069\u006ee\u0041":
		*_gcdg = 8
	case "\u004c\u0069\u0073\u0074":
		*_gcdg = 9
	case "\u004d\u006f\u0076i\u0065":
		*_gcdg = 10
	case "\u004e\u006f\u0074\u0065":
		*_gcdg = 11
	case "\u0050\u0069\u0063\u0074":
		*_gcdg = 12
	case "\u0052\u0061\u0064i\u006f":
		*_gcdg = 13
	case "\u0052\u0065\u0063t\u0041":
		*_gcdg = 14
	case "\u0053\u0063\u0072\u006f\u006c\u006c":
		*_gcdg = 15
	case "\u0053\u0070\u0069\u006e":
		*_gcdg = 16
	case "\u0053\u0068\u0061p\u0065":
		*_gcdg = 17
	case "\u0047\u0072\u006fu\u0070":
		*_gcdg = 18
	case "\u0052\u0065\u0063\u0074":
		*_gcdg = 19
	}
	return nil
}

type CT_ClientData struct {
	ObjectTypeAttr ST_ObjectType
	MoveWithCells  _b.ST_TrueFalseBlank
	SizeWithCells  _b.ST_TrueFalseBlank
	Anchor         *string
	Locked         _b.ST_TrueFalseBlank
	DefaultSize    _b.ST_TrueFalseBlank
	PrintObject    _b.ST_TrueFalseBlank
	Disabled       _b.ST_TrueFalseBlank
	AutoFill       _b.ST_TrueFalseBlank
	AutoLine       _b.ST_TrueFalseBlank
	AutoPict       _b.ST_TrueFalseBlank
	FmlaMacro      *string
	TextHAlign     *string
	TextVAlign     *string
	LockText       _b.ST_TrueFalseBlank
	JustLastX      _b.ST_TrueFalseBlank
	SecretEdit     _b.ST_TrueFalseBlank
	Default        _b.ST_TrueFalseBlank
	Help           _b.ST_TrueFalseBlank
	Cancel         _b.ST_TrueFalseBlank
	Dismiss        _b.ST_TrueFalseBlank
	Accel          *int64
	Accel2         *int64
	Row            *int64
	Column         *int64
	Visible        _b.ST_TrueFalseBlank
	RowHidden      _b.ST_TrueFalseBlank
	ColHidden      _b.ST_TrueFalseBlank
	VTEdit         *int64
	MultiLine      _b.ST_TrueFalseBlank
	VScroll        _b.ST_TrueFalseBlank
	ValidIds       _b.ST_TrueFalseBlank
	FmlaRange      *string
	WidthMin       *int64
	Sel            *int64
	NoThreeD2      _b.ST_TrueFalseBlank
	SelType        *string
	MultiSel       *string
	LCT            *string
	ListItem       *string
	DropStyle      *string
	Colored        _b.ST_TrueFalseBlank
	DropLines      *int64
	Checked        *int64
	FmlaLink       *string
	FmlaPict       *string
	NoThreeD       _b.ST_TrueFalseBlank
	FirstButton    _b.ST_TrueFalseBlank
	FmlaGroup      *string
	Val            *int64
	Min            *int64
	Max            *int64
	Inc            *int64
	Page           *int64
	Horiz          _b.ST_TrueFalseBlank
	Dx             *int64
	MapOCX         _b.ST_TrueFalseBlank
	CF             []string
	Camera         _b.ST_TrueFalseBlank
	RecalcAlways   _b.ST_TrueFalseBlank
	AutoScale      _b.ST_TrueFalseBlank
	DDE            _b.ST_TrueFalseBlank
	UIObj          _b.ST_TrueFalseBlank
	ScriptText     *string
	ScriptExtended *string
	ScriptLanguage *uint32
	ScriptLocation *uint32
	FmlaTxbx       *string
}

func (_fbe ST_ObjectType) Validate() error { return _fbe.ValidateWithPath("") }

type ClientData struct{ CT_ClientData }

func (_bece ST_ObjectType) MarshalXML(e *_ef.Encoder, start _ef.StartElement) error {
	return e.EncodeElement(_bece.String(), start)
}
func (_df *CT_ClientData) UnmarshalXML(d *_ef.Decoder, start _ef.StartElement) error {
	_df.ObjectTypeAttr = ST_ObjectType(1)
	for _, _gceb := range start.Attr {
		if _gceb.Name.Local == "\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065" {
			_df.ObjectTypeAttr.UnmarshalXMLAttr(_gceb)
			continue
		}
	}
_cbb:
	for {
		_dae, _ba := d.Token()
		if _ba != nil {
			return _ba
		}
		switch _bac := _dae.(type) {
		case _ef.StartElement:
			switch _bac.Name {
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u006f\u0076\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_df.MoveWithCells = _b.ST_TrueFalseBlankUnset
				if _eag := d.DecodeElement(&_df.MoveWithCells, &_bac); _eag != nil {
					return _eag
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0069\u007a\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_df.SizeWithCells = _b.ST_TrueFalseBlankUnset
				if _ab := d.DecodeElement(&_df.SizeWithCells, &_bac); _ab != nil {
					return _ab
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_df.Anchor = new(string)
				if _acd := d.DecodeElement(_df.Anchor, &_bac); _acd != nil {
					return _acd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0065\u0064"}:
				_df.Locked = _b.ST_TrueFalseBlankUnset
				if _fdbg := d.DecodeElement(&_df.Locked, &_bac); _fdbg != nil {
					return _fdbg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "D\u0065\u0066\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"}:
				_df.DefaultSize = _b.ST_TrueFalseBlankUnset
				if _ddd := d.DecodeElement(&_df.DefaultSize, &_bac); _ddd != nil {
					return _ddd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "P\u0072\u0069\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"}:
				_df.PrintObject = _b.ST_TrueFalseBlankUnset
				if _bdb := d.DecodeElement(&_df.PrintObject, &_bac); _bdb != nil {
					return _bdb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0069\u0073\u0061\u0062\u006c\u0065\u0064"}:
				_df.Disabled = _b.ST_TrueFalseBlankUnset
				if _bcc := d.DecodeElement(&_df.Disabled, &_bac); _bcc != nil {
					return _bcc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0046\u0069\u006c\u006c"}:
				_df.AutoFill = _b.ST_TrueFalseBlankUnset
				if _bcgg := d.DecodeElement(&_df.AutoFill, &_bac); _bcgg != nil {
					return _bcgg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u004c\u0069\u006e\u0065"}:
				_df.AutoLine = _b.ST_TrueFalseBlankUnset
				if _gad := d.DecodeElement(&_df.AutoLine, &_bac); _gad != nil {
					return _gad
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0050\u0069\u0063\u0074"}:
				_df.AutoPict = _b.ST_TrueFalseBlankUnset
				if _ggee := d.DecodeElement(&_df.AutoPict, &_bac); _ggee != nil {
					return _ggee
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u004d\u0061\u0063\u0072o"}:
				_df.FmlaMacro = new(string)
				if _cca := d.DecodeElement(_df.FmlaMacro, &_bac); _cca != nil {
					return _cca
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0048\u0041\u006c\u0069\u0067\u006e"}:
				_df.TextHAlign = new(string)
				if _ece := d.DecodeElement(_df.TextHAlign, &_bac); _ece != nil {
					return _ece
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0056\u0041\u006c\u0069\u0067\u006e"}:
				_df.TextVAlign = new(string)
				if _daec := d.DecodeElement(_df.TextVAlign, &_bac); _daec != nil {
					return _daec
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0054\u0065\u0078\u0074"}:
				_df.LockText = _b.ST_TrueFalseBlankUnset
				if _gee := d.DecodeElement(&_df.LockText, &_bac); _gee != nil {
					return _gee
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004au\u0073\u0074\u004c\u0061\u0073\u0074X"}:
				_df.JustLastX = _b.ST_TrueFalseBlankUnset
				if _bef := d.DecodeElement(&_df.JustLastX, &_bac); _bef != nil {
					return _bef
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"}:
				_df.SecretEdit = _b.ST_TrueFalseBlankUnset
				if _acef := d.DecodeElement(&_df.SecretEdit, &_bac); _acef != nil {
					return _acef
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_df.Default = _b.ST_TrueFalseBlankUnset
				if _bcgd := d.DecodeElement(&_df.Default, &_bac); _bcgd != nil {
					return _bcgd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u0065\u006c\u0070"}:
				_df.Help = _b.ST_TrueFalseBlankUnset
				if _gfe := d.DecodeElement(&_df.Help, &_bac); _gfe != nil {
					return _gfe
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006e\u0063\u0065\u006c"}:
				_df.Cancel = _b.ST_TrueFalseBlankUnset
				if _cgg := d.DecodeElement(&_df.Cancel, &_bac); _cgg != nil {
					return _cgg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044i\u0073\u006d\u0069\u0073\u0073"}:
				_df.Dismiss = _b.ST_TrueFalseBlankUnset
				if _ggda := d.DecodeElement(&_df.Dismiss, &_bac); _ggda != nil {
					return _ggda
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063e\u006c"}:
				_df.Accel = new(int64)
				if _bad := d.DecodeElement(_df.Accel, &_bac); _bad != nil {
					return _bad
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063\u0065\u006c\u0032"}:
				_df.Accel2 = new(int64)
				if _fdf := d.DecodeElement(_df.Accel2, &_bac); _fdf != nil {
					return _fdf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u006f\u0077"}:
				_df.Row = new(int64)
				if _gdg := d.DecodeElement(_df.Row, &_bac); _gdg != nil {
					return _gdg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u006f\u006c\u0075\u006d\u006e"}:
				_df.Column = new(int64)
				if _aed := d.DecodeElement(_df.Column, &_bac); _aed != nil {
					return _aed
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056i\u0073\u0069\u0062\u006c\u0065"}:
				_df.Visible = _b.ST_TrueFalseBlankUnset
				if _dfg := d.DecodeElement(&_df.Visible, &_bac); _dfg != nil {
					return _dfg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052o\u0077\u0048\u0069\u0064\u0064\u0065n"}:
				_df.RowHidden = _b.ST_TrueFalseBlankUnset
				if _ff := d.DecodeElement(&_df.RowHidden, &_bac); _ff != nil {
					return _ff
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u0048\u0069\u0064\u0064\u0065n"}:
				_df.ColHidden = _b.ST_TrueFalseBlankUnset
				if _db := d.DecodeElement(&_df.ColHidden, &_bac); _db != nil {
					return _db
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0054\u0045\u0064\u0069\u0074"}:
				_df.VTEdit = new(int64)
				if _deg := d.DecodeElement(_df.VTEdit, &_bac); _deg != nil {
					return _deg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004du\u006c\u0074\u0069\u004c\u0069\u006ee"}:
				_df.MultiLine = _b.ST_TrueFalseBlankUnset
				if _cd := d.DecodeElement(&_df.MultiLine, &_bac); _cd != nil {
					return _cd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056S\u0063\u0072\u006f\u006c\u006c"}:
				_df.VScroll = _b.ST_TrueFalseBlankUnset
				if _egf := d.DecodeElement(&_df.VScroll, &_bac); _egf != nil {
					return _egf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c\u0069\u0064\u0049\u0064\u0073"}:
				_df.ValidIds = _b.ST_TrueFalseBlankUnset
				if _edc := d.DecodeElement(&_df.ValidIds, &_bac); _edc != nil {
					return _edc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0052\u0061\u006e\u0067e"}:
				_df.FmlaRange = new(string)
				if _ccd := d.DecodeElement(_df.FmlaRange, &_bac); _ccd != nil {
					return _ccd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0057\u0069\u0064\u0074\u0068\u004d\u0069\u006e"}:
				_df.WidthMin = new(int64)
				if _cda := d.DecodeElement(_df.WidthMin, &_bac); _cda != nil {
					return _cda
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u006c"}:
				_df.Sel = new(int64)
				if _bdca := d.DecodeElement(_df.Sel, &_bac); _bdca != nil {
					return _bdca
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004eo\u0054\u0068\u0072\u0065\u0065\u00442"}:
				_df.NoThreeD2 = _b.ST_TrueFalseBlankUnset
				if _gab := d.DecodeElement(&_df.NoThreeD2, &_bac); _gab != nil {
					return _gab
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053e\u006c\u0054\u0079\u0070\u0065"}:
				_df.SelType = new(string)
				if _bfd := d.DecodeElement(_df.SelType, &_bac); _bfd != nil {
					return _bfd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0075\u006c\u0074\u0069\u0053\u0065\u006c"}:
				_df.MultiSel = new(string)
				if _dfgf := d.DecodeElement(_df.MultiSel, &_bac); _dfgf != nil {
					return _dfgf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0043\u0054"}:
				_df.LCT = new(string)
				if _eef := d.DecodeElement(_df.LCT, &_bac); _eef != nil {
					return _eef
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0069\u0073\u0074\u0049\u0074\u0065\u006d"}:
				_df.ListItem = new(string)
				if _cdc := d.DecodeElement(_df.ListItem, &_bac); _cdc != nil {
					return _cdc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u0053\u0074\u0079\u006ce"}:
				_df.DropStyle = new(string)
				if _ffa := d.DecodeElement(_df.DropStyle, &_bac); _ffa != nil {
					return _ffa
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u006f\u0072\u0065\u0064"}:
				_df.Colored = _b.ST_TrueFalseBlankUnset
				if _cdd := d.DecodeElement(&_df.Colored, &_bac); _cdd != nil {
					return _cdd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u004c\u0069\u006e\u0065s"}:
				_df.DropLines = new(int64)
				if _gac := d.DecodeElement(_df.DropLines, &_bac); _gac != nil {
					return _gac
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043h\u0065\u0063\u006b\u0065\u0064"}:
				_df.Checked = new(int64)
				if _ege := d.DecodeElement(_df.Checked, &_bac); _ege != nil {
					return _ege
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u004c\u0069\u006e\u006b"}:
				_df.FmlaLink = new(string)
				if _cge := d.DecodeElement(_df.FmlaLink, &_bac); _cge != nil {
					return _cge
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0050\u0069\u0063\u0074"}:
				_df.FmlaPict = new(string)
				if _aa := d.DecodeElement(_df.FmlaPict, &_bac); _aa != nil {
					return _aa
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044"}:
				_df.NoThreeD = _b.ST_TrueFalseBlankUnset
				if _bfde := d.DecodeElement(&_df.NoThreeD, &_bac); _bfde != nil {
					return _bfde
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "F\u0069\u0072\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"}:
				_df.FirstButton = _b.ST_TrueFalseBlankUnset
				if _cde := d.DecodeElement(&_df.FirstButton, &_bac); _cde != nil {
					return _cde
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0047\u0072\u006f\u0075p"}:
				_df.FmlaGroup = new(string)
				if _fg := d.DecodeElement(_df.FmlaGroup, &_bac); _fg != nil {
					return _fg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c"}:
				_df.Val = new(int64)
				if _accb := d.DecodeElement(_df.Val, &_bac); _accb != nil {
					return _accb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0069\u006e"}:
				_df.Min = new(int64)
				if _ffg := d.DecodeElement(_df.Min, &_bac); _ffg != nil {
					return _ffg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0078"}:
				_df.Max = new(int64)
				if _acb := d.DecodeElement(_df.Max, &_bac); _acb != nil {
					return _acb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0049\u006e\u0063"}:
				_df.Inc = new(int64)
				if _fdfe := d.DecodeElement(_df.Inc, &_bac); _fdfe != nil {
					return _fdfe
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0050\u0061\u0067\u0065"}:
				_df.Page = new(int64)
				if _bb := d.DecodeElement(_df.Page, &_bac); _bb != nil {
					return _bb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u006f\u0072i\u007a"}:
				_df.Horiz = _b.ST_TrueFalseBlankUnset
				if _caab := d.DecodeElement(&_df.Horiz, &_bac); _caab != nil {
					return _caab
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0078"}:
				_df.Dx = new(int64)
				if _dde := d.DecodeElement(_df.Dx, &_bac); _dde != nil {
					return _dde
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0070\u004f\u0043\u0058"}:
				_df.MapOCX = _b.ST_TrueFalseBlankUnset
				if _abg := d.DecodeElement(&_df.MapOCX, &_bac); _abg != nil {
					return _abg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0046"}:
				var _fe string
				if _abd := d.DecodeElement(&_fe, &_bac); _abd != nil {
					return _abd
				}
				_df.CF = append(_df.CF, _fe)
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006d\u0065\u0072\u0061"}:
				_df.Camera = _b.ST_TrueFalseBlankUnset
				if _ffc := d.DecodeElement(&_df.Camera, &_bac); _ffc != nil {
					return _ffc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u0065\u0063a\u006c\u0063\u0041\u006c\u0077\u0061\u0079\u0073"}:
				_df.RecalcAlways = _b.ST_TrueFalseBlankUnset
				if _badd := d.DecodeElement(&_df.RecalcAlways, &_bac); _badd != nil {
					return _badd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041u\u0074\u006f\u0053\u0063\u0061\u006ce"}:
				_df.AutoScale = _b.ST_TrueFalseBlankUnset
				if _gfc := d.DecodeElement(&_df.AutoScale, &_bac); _gfc != nil {
					return _gfc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0044\u0045"}:
				_df.DDE = _b.ST_TrueFalseBlankUnset
				if _bda := d.DecodeElement(&_df.DDE, &_bac); _bda != nil {
					return _bda
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0055\u0049\u004fb\u006a"}:
				_df.UIObj = _b.ST_TrueFalseBlankUnset
				if _ccg := d.DecodeElement(&_df.UIObj, &_bac); _ccg != nil {
					return _ccg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0054\u0065\u0078\u0074"}:
				_df.ScriptText = new(string)
				if _ggb := d.DecodeElement(_df.ScriptText, &_bac); _ggb != nil {
					return _ggb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0045\u0078\u0074e\u006e\u0064\u0065\u0064"}:
				_df.ScriptExtended = new(string)
				if _bgae := d.DecodeElement(_df.ScriptExtended, &_bac); _bgae != nil {
					return _bgae
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"}:
				_df.ScriptLanguage = new(uint32)
				if _fgg := d.DecodeElement(_df.ScriptLanguage, &_bac); _fgg != nil {
					return _fgg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"}:
				_df.ScriptLocation = new(uint32)
				if _eagg := d.DecodeElement(_df.ScriptLocation, &_bac); _eagg != nil {
					return _eagg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0054\u0078\u0062\u0078"}:
				_df.FmlaTxbx = new(string)
				if _dfc := d.DecodeElement(_df.FmlaTxbx, &_bac); _dfc != nil {
					return _dfc
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043l\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061 \u0025\u0076", _bac.Name)
				if _gcgb := d.Skip(); _gcgb != nil {
					return _gcgb
				}
			}
		case _ef.EndElement:
			break _cbb
		case _ef.CharData:
		}
	}
	return nil
}
func NewClientData() *ClientData {
	_fcf := &ClientData{}
	_fcf.CT_ClientData = *NewCT_ClientData()
	return _fcf
}
func (_befb ST_ObjectType) ValidateWithPath(path string) error {
	switch _befb {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_befb))
	}
	return nil
}
func (_gfg *ClientData) UnmarshalXML(d *_ef.Decoder, start _ef.StartElement) error {
	_gfg.CT_ClientData = *NewCT_ClientData()
	for _, _ffd := range start.Attr {
		if _ffd.Name.Local == "\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065" {
			_gfg.ObjectTypeAttr.UnmarshalXMLAttr(_ffd)
			continue
		}
	}
_eac:
	for {
		_egd, _fag := d.Token()
		if _fag != nil {
			return _fag
		}
		switch _bbdc := _egd.(type) {
		case _ef.StartElement:
			switch _bbdc.Name {
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u006f\u0076\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_gfg.MoveWithCells = _b.ST_TrueFalseBlankUnset
				if _afe := d.DecodeElement(&_gfg.MoveWithCells, &_bbdc); _afe != nil {
					return _afe
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0069\u007a\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_gfg.SizeWithCells = _b.ST_TrueFalseBlankUnset
				if _bba := d.DecodeElement(&_gfg.SizeWithCells, &_bbdc); _bba != nil {
					return _bba
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_gfg.Anchor = new(string)
				if _fge := d.DecodeElement(_gfg.Anchor, &_bbdc); _fge != nil {
					return _fge
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0065\u0064"}:
				_gfg.Locked = _b.ST_TrueFalseBlankUnset
				if _gacf := d.DecodeElement(&_gfg.Locked, &_bbdc); _gacf != nil {
					return _gacf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "D\u0065\u0066\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"}:
				_gfg.DefaultSize = _b.ST_TrueFalseBlankUnset
				if _fgee := d.DecodeElement(&_gfg.DefaultSize, &_bbdc); _fgee != nil {
					return _fgee
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "P\u0072\u0069\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"}:
				_gfg.PrintObject = _b.ST_TrueFalseBlankUnset
				if _ead := d.DecodeElement(&_gfg.PrintObject, &_bbdc); _ead != nil {
					return _ead
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0069\u0073\u0061\u0062\u006c\u0065\u0064"}:
				_gfg.Disabled = _b.ST_TrueFalseBlankUnset
				if _bcfc := d.DecodeElement(&_gfg.Disabled, &_bbdc); _bcfc != nil {
					return _bcfc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0046\u0069\u006c\u006c"}:
				_gfg.AutoFill = _b.ST_TrueFalseBlankUnset
				if _cged := d.DecodeElement(&_gfg.AutoFill, &_bbdc); _cged != nil {
					return _cged
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u004c\u0069\u006e\u0065"}:
				_gfg.AutoLine = _b.ST_TrueFalseBlankUnset
				if _eaab := d.DecodeElement(&_gfg.AutoLine, &_bbdc); _eaab != nil {
					return _eaab
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0050\u0069\u0063\u0074"}:
				_gfg.AutoPict = _b.ST_TrueFalseBlankUnset
				if _gda := d.DecodeElement(&_gfg.AutoPict, &_bbdc); _gda != nil {
					return _gda
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u004d\u0061\u0063\u0072o"}:
				_gfg.FmlaMacro = new(string)
				if _agc := d.DecodeElement(_gfg.FmlaMacro, &_bbdc); _agc != nil {
					return _agc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0048\u0041\u006c\u0069\u0067\u006e"}:
				_gfg.TextHAlign = new(string)
				if _efb := d.DecodeElement(_gfg.TextHAlign, &_bbdc); _efb != nil {
					return _efb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0056\u0041\u006c\u0069\u0067\u006e"}:
				_gfg.TextVAlign = new(string)
				if _egdb := d.DecodeElement(_gfg.TextVAlign, &_bbdc); _egdb != nil {
					return _egdb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0054\u0065\u0078\u0074"}:
				_gfg.LockText = _b.ST_TrueFalseBlankUnset
				if _dff := d.DecodeElement(&_gfg.LockText, &_bbdc); _dff != nil {
					return _dff
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004au\u0073\u0074\u004c\u0061\u0073\u0074X"}:
				_gfg.JustLastX = _b.ST_TrueFalseBlankUnset
				if _bgaed := d.DecodeElement(&_gfg.JustLastX, &_bbdc); _bgaed != nil {
					return _bgaed
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"}:
				_gfg.SecretEdit = _b.ST_TrueFalseBlankUnset
				if _daf := d.DecodeElement(&_gfg.SecretEdit, &_bbdc); _daf != nil {
					return _daf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_gfg.Default = _b.ST_TrueFalseBlankUnset
				if _gcd := d.DecodeElement(&_gfg.Default, &_bbdc); _gcd != nil {
					return _gcd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u0065\u006c\u0070"}:
				_gfg.Help = _b.ST_TrueFalseBlankUnset
				if _bbf := d.DecodeElement(&_gfg.Help, &_bbdc); _bbf != nil {
					return _bbf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006e\u0063\u0065\u006c"}:
				_gfg.Cancel = _b.ST_TrueFalseBlankUnset
				if _eca := d.DecodeElement(&_gfg.Cancel, &_bbdc); _eca != nil {
					return _eca
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044i\u0073\u006d\u0069\u0073\u0073"}:
				_gfg.Dismiss = _b.ST_TrueFalseBlankUnset
				if _aacf := d.DecodeElement(&_gfg.Dismiss, &_bbdc); _aacf != nil {
					return _aacf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063e\u006c"}:
				_gfg.Accel = new(int64)
				if _cacb := d.DecodeElement(_gfg.Accel, &_bbdc); _cacb != nil {
					return _cacb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063\u0065\u006c\u0032"}:
				_gfg.Accel2 = new(int64)
				if _ffce := d.DecodeElement(_gfg.Accel2, &_bbdc); _ffce != nil {
					return _ffce
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u006f\u0077"}:
				_gfg.Row = new(int64)
				if _dgg := d.DecodeElement(_gfg.Row, &_bbdc); _dgg != nil {
					return _dgg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u006f\u006c\u0075\u006d\u006e"}:
				_gfg.Column = new(int64)
				if _dec := d.DecodeElement(_gfg.Column, &_bbdc); _dec != nil {
					return _dec
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056i\u0073\u0069\u0062\u006c\u0065"}:
				_gfg.Visible = _b.ST_TrueFalseBlankUnset
				if _dddg := d.DecodeElement(&_gfg.Visible, &_bbdc); _dddg != nil {
					return _dddg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052o\u0077\u0048\u0069\u0064\u0064\u0065n"}:
				_gfg.RowHidden = _b.ST_TrueFalseBlankUnset
				if _ddec := d.DecodeElement(&_gfg.RowHidden, &_bbdc); _ddec != nil {
					return _ddec
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u0048\u0069\u0064\u0064\u0065n"}:
				_gfg.ColHidden = _b.ST_TrueFalseBlankUnset
				if _bbfb := d.DecodeElement(&_gfg.ColHidden, &_bbdc); _bbfb != nil {
					return _bbfb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0054\u0045\u0064\u0069\u0074"}:
				_gfg.VTEdit = new(int64)
				if _cgbb := d.DecodeElement(_gfg.VTEdit, &_bbdc); _cgbb != nil {
					return _cgbb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004du\u006c\u0074\u0069\u004c\u0069\u006ee"}:
				_gfg.MultiLine = _b.ST_TrueFalseBlankUnset
				if _bdg := d.DecodeElement(&_gfg.MultiLine, &_bbdc); _bdg != nil {
					return _bdg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056S\u0063\u0072\u006f\u006c\u006c"}:
				_gfg.VScroll = _b.ST_TrueFalseBlankUnset
				if _bfg := d.DecodeElement(&_gfg.VScroll, &_bbdc); _bfg != nil {
					return _bfg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c\u0069\u0064\u0049\u0064\u0073"}:
				_gfg.ValidIds = _b.ST_TrueFalseBlankUnset
				if _ccde := d.DecodeElement(&_gfg.ValidIds, &_bbdc); _ccde != nil {
					return _ccde
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0052\u0061\u006e\u0067e"}:
				_gfg.FmlaRange = new(string)
				if _gga := d.DecodeElement(_gfg.FmlaRange, &_bbdc); _gga != nil {
					return _gga
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0057\u0069\u0064\u0074\u0068\u004d\u0069\u006e"}:
				_gfg.WidthMin = new(int64)
				if _eed := d.DecodeElement(_gfg.WidthMin, &_bbdc); _eed != nil {
					return _eed
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u006c"}:
				_gfg.Sel = new(int64)
				if _fecc := d.DecodeElement(_gfg.Sel, &_bbdc); _fecc != nil {
					return _fecc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004eo\u0054\u0068\u0072\u0065\u0065\u00442"}:
				_gfg.NoThreeD2 = _b.ST_TrueFalseBlankUnset
				if _gaf := d.DecodeElement(&_gfg.NoThreeD2, &_bbdc); _gaf != nil {
					return _gaf
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053e\u006c\u0054\u0079\u0070\u0065"}:
				_gfg.SelType = new(string)
				if _eeb := d.DecodeElement(_gfg.SelType, &_bbdc); _eeb != nil {
					return _eeb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0075\u006c\u0074\u0069\u0053\u0065\u006c"}:
				_gfg.MultiSel = new(string)
				if _cfd := d.DecodeElement(_gfg.MultiSel, &_bbdc); _cfd != nil {
					return _cfd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0043\u0054"}:
				_gfg.LCT = new(string)
				if _gb := d.DecodeElement(_gfg.LCT, &_bbdc); _gb != nil {
					return _gb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0069\u0073\u0074\u0049\u0074\u0065\u006d"}:
				_gfg.ListItem = new(string)
				if _ccgd := d.DecodeElement(_gfg.ListItem, &_bbdc); _ccgd != nil {
					return _ccgd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u0053\u0074\u0079\u006ce"}:
				_gfg.DropStyle = new(string)
				if _fdge := d.DecodeElement(_gfg.DropStyle, &_bbdc); _fdge != nil {
					return _fdge
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u006f\u0072\u0065\u0064"}:
				_gfg.Colored = _b.ST_TrueFalseBlankUnset
				if _bge := d.DecodeElement(&_gfg.Colored, &_bbdc); _bge != nil {
					return _bge
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u004c\u0069\u006e\u0065s"}:
				_gfg.DropLines = new(int64)
				if _bbac := d.DecodeElement(_gfg.DropLines, &_bbdc); _bbac != nil {
					return _bbac
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043h\u0065\u0063\u006b\u0065\u0064"}:
				_gfg.Checked = new(int64)
				if _bcbb := d.DecodeElement(_gfg.Checked, &_bbdc); _bcbb != nil {
					return _bcbb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u004c\u0069\u006e\u006b"}:
				_gfg.FmlaLink = new(string)
				if _bdaa := d.DecodeElement(_gfg.FmlaLink, &_bbdc); _bdaa != nil {
					return _bdaa
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0050\u0069\u0063\u0074"}:
				_gfg.FmlaPict = new(string)
				if _gcbe := d.DecodeElement(_gfg.FmlaPict, &_bbdc); _gcbe != nil {
					return _gcbe
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044"}:
				_gfg.NoThreeD = _b.ST_TrueFalseBlankUnset
				if _fbc := d.DecodeElement(&_gfg.NoThreeD, &_bbdc); _fbc != nil {
					return _fbc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "F\u0069\u0072\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"}:
				_gfg.FirstButton = _b.ST_TrueFalseBlankUnset
				if _fab := d.DecodeElement(&_gfg.FirstButton, &_bbdc); _fab != nil {
					return _fab
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0047\u0072\u006f\u0075p"}:
				_gfg.FmlaGroup = new(string)
				if _cab := d.DecodeElement(_gfg.FmlaGroup, &_bbdc); _cab != nil {
					return _cab
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c"}:
				_gfg.Val = new(int64)
				if _ecee := d.DecodeElement(_gfg.Val, &_bbdc); _ecee != nil {
					return _ecee
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0069\u006e"}:
				_gfg.Min = new(int64)
				if _efc := d.DecodeElement(_gfg.Min, &_bbdc); _efc != nil {
					return _efc
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0078"}:
				_gfg.Max = new(int64)
				if _aaa := d.DecodeElement(_gfg.Max, &_bbdc); _aaa != nil {
					return _aaa
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0049\u006e\u0063"}:
				_gfg.Inc = new(int64)
				if _bacb := d.DecodeElement(_gfg.Inc, &_bbdc); _bacb != nil {
					return _bacb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0050\u0061\u0067\u0065"}:
				_gfg.Page = new(int64)
				if _bdae := d.DecodeElement(_gfg.Page, &_bbdc); _bdae != nil {
					return _bdae
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u006f\u0072i\u007a"}:
				_gfg.Horiz = _b.ST_TrueFalseBlankUnset
				if _fde := d.DecodeElement(&_gfg.Horiz, &_bbdc); _fde != nil {
					return _fde
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0078"}:
				_gfg.Dx = new(int64)
				if _feg := d.DecodeElement(_gfg.Dx, &_bbdc); _feg != nil {
					return _feg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0070\u004f\u0043\u0058"}:
				_gfg.MapOCX = _b.ST_TrueFalseBlankUnset
				if _dcdd := d.DecodeElement(&_gfg.MapOCX, &_bbdc); _dcdd != nil {
					return _dcdd
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0046"}:
				var _fece string
				if _bgaeb := d.DecodeElement(&_fece, &_bbdc); _bgaeb != nil {
					return _bgaeb
				}
				_gfg.CF = append(_gfg.CF, _fece)
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006d\u0065\u0072\u0061"}:
				_gfg.Camera = _b.ST_TrueFalseBlankUnset
				if _gef := d.DecodeElement(&_gfg.Camera, &_bbdc); _gef != nil {
					return _gef
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u0065\u0063a\u006c\u0063\u0041\u006c\u0077\u0061\u0079\u0073"}:
				_gfg.RecalcAlways = _b.ST_TrueFalseBlankUnset
				if _fdee := d.DecodeElement(&_gfg.RecalcAlways, &_bbdc); _fdee != nil {
					return _fdee
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041u\u0074\u006f\u0053\u0063\u0061\u006ce"}:
				_gfg.AutoScale = _b.ST_TrueFalseBlankUnset
				if _ggbe := d.DecodeElement(&_gfg.AutoScale, &_bbdc); _ggbe != nil {
					return _ggbe
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0044\u0045"}:
				_gfg.DDE = _b.ST_TrueFalseBlankUnset
				if _dac := d.DecodeElement(&_gfg.DDE, &_bbdc); _dac != nil {
					return _dac
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0055\u0049\u004fb\u006a"}:
				_gfg.UIObj = _b.ST_TrueFalseBlankUnset
				if _gcbb := d.DecodeElement(&_gfg.UIObj, &_bbdc); _gcbb != nil {
					return _gcbb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0054\u0065\u0078\u0074"}:
				_gfg.ScriptText = new(string)
				if _fce := d.DecodeElement(_gfg.ScriptText, &_bbdc); _fce != nil {
					return _fce
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0045\u0078\u0074e\u006e\u0064\u0065\u0064"}:
				_gfg.ScriptExtended = new(string)
				if _ccae := d.DecodeElement(_gfg.ScriptExtended, &_bbdc); _ccae != nil {
					return _ccae
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"}:
				_gfg.ScriptLanguage = new(uint32)
				if _cdg := d.DecodeElement(_gfg.ScriptLanguage, &_bbdc); _cdg != nil {
					return _cdg
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"}:
				_gfg.ScriptLocation = new(uint32)
				if _gfcb := d.DecodeElement(_gfg.ScriptLocation, &_bbdc); _gfcb != nil {
					return _gfcb
				}
			case _ef.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0054\u0078\u0062\u0078"}:
				_gfg.FmlaTxbx = new(string)
				if _ede := d.DecodeElement(_gfg.FmlaTxbx, &_bbdc); _ede != nil {
					return _ede
				}
			default:
				_f.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u006c\u0069e\u006e\u0074\u0044\u0061\u0074\u0061\u0020\u0025\u0076", _bbdc.Name)
				if _eeg := d.Skip(); _eeg != nil {
					return _eeg
				}
			}
		case _ef.EndElement:
			break _eac
		case _ef.CharData:
		}
	}
	return nil
}
func (_bg *CT_ClientData) MarshalXML(e *_ef.Encoder, start _ef.StartElement) error {
	_g, _gc := _bg.ObjectTypeAttr.MarshalXMLAttr(_ef.Name{Local: "\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065"})
	if _gc != nil {
		return _gc
	}
	start.Attr = append(start.Attr, _g)
	e.EncodeToken(start)
	if _bg.MoveWithCells != _b.ST_TrueFalseBlankUnset {
		_ce := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u004do\u0076\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}}
		e.EncodeElement(_bg.MoveWithCells, _ce)
	}
	if _bg.SizeWithCells != _b.ST_TrueFalseBlankUnset {
		_a := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0053i\u007a\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}}
		e.EncodeElement(_bg.SizeWithCells, _a)
	}
	if _bg.Anchor != nil {
		_ca := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0041\u006e\u0063\u0068\u006f\u0072"}}
		_f.AddPreserveSpaceAttr(&_ca, *_bg.Anchor)
		e.EncodeElement(_bg.Anchor, _ca)
	}
	if _bg.Locked != _b.ST_TrueFalseBlankUnset {
		_ec := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004c\u006f\u0063\u006b\u0065\u0064"}}
		e.EncodeElement(_bg.Locked, _ec)
	}
	if _bg.DefaultSize != _b.ST_TrueFalseBlankUnset {
		_gcb := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"}}
		e.EncodeElement(_bg.DefaultSize, _gcb)
	}
	if _bg.PrintObject != _b.ST_TrueFalseBlankUnset {
		_gca := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0050\u0072\u0069\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"}}
		e.EncodeElement(_bg.PrintObject, _gca)
	}
	if _bg.Disabled != _b.ST_TrueFalseBlankUnset {
		_fc := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0044\u0069\u0073\u0061\u0062\u006c\u0065\u0064"}}
		e.EncodeElement(_bg.Disabled, _fc)
	}
	if _bg.AutoFill != _b.ST_TrueFalseBlankUnset {
		_bf := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0041\u0075\u0074\u006f\u0046\u0069\u006c\u006c"}}
		e.EncodeElement(_bg.AutoFill, _bf)
	}
	if _bg.AutoLine != _b.ST_TrueFalseBlankUnset {
		_ad := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0041\u0075\u0074\u006f\u004c\u0069\u006e\u0065"}}
		e.EncodeElement(_bg.AutoLine, _ad)
	}
	if _bg.AutoPict != _b.ST_TrueFalseBlankUnset {
		_bd := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0041\u0075\u0074\u006f\u0050\u0069\u0063\u0074"}}
		e.EncodeElement(_bg.AutoPict, _bd)
	}
	if _bg.FmlaMacro != nil {
		_fb := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0046\u006d\u006c\u0061\u004d\u0061\u0063\u0072\u006f"}}
		_f.AddPreserveSpaceAttr(&_fb, *_bg.FmlaMacro)
		e.EncodeElement(_bg.FmlaMacro, _fb)
	}
	if _bg.TextHAlign != nil {
		_bc := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0054e\u0078\u0074\u0048\u0041\u006c\u0069\u0067\u006e"}}
		_f.AddPreserveSpaceAttr(&_bc, *_bg.TextHAlign)
		e.EncodeElement(_bg.TextHAlign, _bc)
	}
	if _bg.TextVAlign != nil {
		_gcc := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0054e\u0078\u0074\u0056\u0041\u006c\u0069\u0067\u006e"}}
		_f.AddPreserveSpaceAttr(&_gcc, *_bg.TextVAlign)
		e.EncodeElement(_bg.TextVAlign, _gcc)
	}
	if _bg.LockText != _b.ST_TrueFalseBlankUnset {
		_bdc := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004c\u006f\u0063\u006b\u0054\u0065\u0078\u0074"}}
		e.EncodeElement(_bg.LockText, _bdc)
	}
	if _bg.JustLastX != _b.ST_TrueFalseBlankUnset {
		_fcd := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u004a\u0075\u0073\u0074\u004c\u0061\u0073\u0074\u0058"}}
		e.EncodeElement(_bg.JustLastX, _fcd)
	}
	if _bg.SecretEdit != _b.ST_TrueFalseBlankUnset {
		_gf := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0053e\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"}}
		e.EncodeElement(_bg.SecretEdit, _gf)
	}
	if _bg.Default != _b.ST_TrueFalseBlankUnset {
		_eg := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0044\u0065\u0066\u0061\u0075\u006ct"}}
		e.EncodeElement(_bg.Default, _eg)
	}
	if _bg.Help != _b.ST_TrueFalseBlankUnset {
		_cag := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0048\u0065\u006c\u0070"}}
		e.EncodeElement(_bg.Help, _cag)
	}
	if _bg.Cancel != _b.ST_TrueFalseBlankUnset {
		_cac := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0043\u0061\u006e\u0063\u0065\u006c"}}
		e.EncodeElement(_bg.Cancel, _cac)
	}
	if _bg.Dismiss != _b.ST_TrueFalseBlankUnset {
		_cf := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0044\u0069\u0073\u006d\u0069\u0073s"}}
		e.EncodeElement(_bg.Dismiss, _cf)
	}
	if _bg.Accel != nil {
		_ea := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0041\u0063\u0063\u0065\u006c"}}
		e.EncodeElement(_bg.Accel, _ea)
	}
	if _bg.Accel2 != nil {
		_ac := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0041\u0063\u0063\u0065\u006c\u0032"}}
		e.EncodeElement(_bg.Accel2, _ac)
	}
	if _bg.Row != nil {
		_gce := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0052o\u0077"}}
		e.EncodeElement(_bg.Row, _gce)
	}
	if _bg.Column != nil {
		_gg := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0043\u006f\u006c\u0075\u006d\u006e"}}
		e.EncodeElement(_bg.Column, _gg)
	}
	if _bg.Visible != _b.ST_TrueFalseBlankUnset {
		_bcb := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0056\u0069\u0073\u0069\u0062\u006ce"}}
		e.EncodeElement(_bg.Visible, _bcb)
	}
	if _bg.RowHidden != _b.ST_TrueFalseBlankUnset {
		_cagd := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0052\u006f\u0077\u0048\u0069\u0064\u0064\u0065\u006e"}}
		e.EncodeElement(_bg.RowHidden, _cagd)
	}
	if _bg.ColHidden != _b.ST_TrueFalseBlankUnset {
		_cfa := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0043\u006f\u006c\u0048\u0069\u0064\u0064\u0065\u006e"}}
		e.EncodeElement(_bg.ColHidden, _cfa)
	}
	if _bg.VTEdit != nil {
		_ge := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0056\u0054\u0045\u0064\u0069\u0074"}}
		e.EncodeElement(_bg.VTEdit, _ge)
	}
	if _bg.MultiLine != _b.ST_TrueFalseBlankUnset {
		_ee := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u004d\u0075\u006c\u0074\u0069\u004c\u0069\u006e\u0065"}}
		e.EncodeElement(_bg.MultiLine, _ee)
	}
	if _bg.VScroll != _b.ST_TrueFalseBlankUnset {
		_fba := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0056\u0053\u0063\u0072\u006f\u006cl"}}
		e.EncodeElement(_bg.VScroll, _fba)
	}
	if _bg.ValidIds != _b.ST_TrueFalseBlankUnset {
		_gd := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0056\u0061\u006c\u0069\u0064\u0049\u0064\u0073"}}
		e.EncodeElement(_bg.ValidIds, _gd)
	}
	if _bg.FmlaRange != nil {
		_egb := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0046\u006d\u006c\u0061\u0052\u0061\u006e\u0067\u0065"}}
		_f.AddPreserveSpaceAttr(&_egb, *_bg.FmlaRange)
		e.EncodeElement(_bg.FmlaRange, _egb)
	}
	if _bg.WidthMin != nil {
		_gcg := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0057\u0069\u0064\u0074\u0068\u004d\u0069\u006e"}}
		e.EncodeElement(_bg.WidthMin, _gcg)
	}
	if _bg.Sel != nil {
		_de := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0053e\u006c"}}
		e.EncodeElement(_bg.Sel, _de)
	}
	if _bg.NoThreeD2 != _b.ST_TrueFalseBlankUnset {
		_dc := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044\u0032"}}
		e.EncodeElement(_bg.NoThreeD2, _dc)
	}
	if _bg.SelType != nil {
		_eaa := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0053\u0065\u006c\u0054\u0079\u0070e"}}
		_f.AddPreserveSpaceAttr(&_eaa, *_bg.SelType)
		e.EncodeElement(_bg.SelType, _eaa)
	}
	if _bg.MultiSel != nil {
		_cg := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004d\u0075\u006c\u0074\u0069\u0053\u0065\u006c"}}
		_f.AddPreserveSpaceAttr(&_cg, *_bg.MultiSel)
		e.EncodeElement(_bg.MultiSel, _cg)
	}
	if _bg.LCT != nil {
		_fd := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004cC\u0054"}}
		_f.AddPreserveSpaceAttr(&_fd, *_bg.LCT)
		e.EncodeElement(_bg.LCT, _fd)
	}
	if _bg.ListItem != nil {
		_fdb := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004c\u0069\u0073\u0074\u0049\u0074\u0065\u006d"}}
		_f.AddPreserveSpaceAttr(&_fdb, *_bg.ListItem)
		e.EncodeElement(_bg.ListItem, _fdb)
	}
	if _bg.DropStyle != nil {
		_ae := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0044\u0072\u006f\u0070\u0053\u0074\u0079\u006c\u0065"}}
		_f.AddPreserveSpaceAttr(&_ae, *_bg.DropStyle)
		e.EncodeElement(_bg.DropStyle, _ae)
	}
	if _bg.Colored != _b.ST_TrueFalseBlankUnset {
		_adg := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0043\u006f\u006c\u006f\u0072\u0065d"}}
		e.EncodeElement(_bg.Colored, _adg)
	}
	if _bg.DropLines != nil {
		_cea := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0044\u0072\u006f\u0070\u004c\u0069\u006e\u0065\u0073"}}
		e.EncodeElement(_bg.DropLines, _cea)
	}
	if _bg.Checked != nil {
		_dd := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0043\u0068\u0065\u0063\u006b\u0065d"}}
		e.EncodeElement(_bg.Checked, _dd)
	}
	if _bg.FmlaLink != nil {
		_ag := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0046\u006d\u006c\u0061\u004c\u0069\u006e\u006b"}}
		_f.AddPreserveSpaceAttr(&_ag, *_bg.FmlaLink)
		e.EncodeElement(_bg.FmlaLink, _ag)
	}
	if _bg.FmlaPict != nil {
		_ded := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0046\u006d\u006c\u0061\u0050\u0069\u0063\u0074"}}
		_f.AddPreserveSpaceAttr(&_ded, *_bg.FmlaPict)
		e.EncodeElement(_bg.FmlaPict, _ded)
	}
	if _bg.NoThreeD != _b.ST_TrueFalseBlankUnset {
		_be := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044"}}
		e.EncodeElement(_bg.NoThreeD, _be)
	}
	if _bg.FirstButton != _b.ST_TrueFalseBlankUnset {
		_acc := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0046\u0069\u0072\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"}}
		e.EncodeElement(_bg.FirstButton, _acc)
	}
	if _bg.FmlaGroup != nil {
		_bga := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0046\u006d\u006c\u0061\u0047\u0072\u006f\u0075\u0070"}}
		_f.AddPreserveSpaceAttr(&_bga, *_bg.FmlaGroup)
		e.EncodeElement(_bg.FmlaGroup, _bga)
	}
	if _bg.Val != nil {
		_gdd := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0056a\u006c"}}
		e.EncodeElement(_bg.Val, _gdd)
	}
	if _bg.Min != nil {
		_gdda := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004di\u006e"}}
		e.EncodeElement(_bg.Min, _gdda)
	}
	if _bg.Max != nil {
		_ace := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004da\u0078"}}
		e.EncodeElement(_bg.Max, _ace)
	}
	if _bg.Inc != nil {
		_bcg := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0049n\u0063"}}
		e.EncodeElement(_bg.Inc, _bcg)
	}
	if _bg.Page != nil {
		_cb := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0050\u0061\u0067\u0065"}}
		e.EncodeElement(_bg.Page, _cb)
	}
	if _bg.Horiz != _b.ST_TrueFalseBlankUnset {
		_ggd := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0048\u006f\u0072\u0069\u007a"}}
		e.EncodeElement(_bg.Horiz, _ggd)
	}
	if _bg.Dx != nil {
		_cgb := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0044\u0078"}}
		e.EncodeElement(_bg.Dx, _cgb)
	}
	if _bg.MapOCX != _b.ST_TrueFalseBlankUnset {
		_ed := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u004d\u0061\u0070\u004f\u0043\u0058"}}
		e.EncodeElement(_bg.MapOCX, _ed)
	}
	if _bg.CF != nil {
		_fa := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0043\u0046"}}
		for _, _dcg := range _bg.CF {
			e.EncodeElement(_dcg, _fa)
		}
	}
	if _bg.Camera != _b.ST_TrueFalseBlankUnset {
		_edb := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0043\u0061\u006d\u0065\u0072\u0061"}}
		e.EncodeElement(_bg.Camera, _edb)
	}
	if _bg.RecalcAlways != _b.ST_TrueFalseBlankUnset {
		_aga := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0052\u0065\u0063\u0061\u006c\u0063\u0041l\u0077\u0061\u0079\u0073"}}
		e.EncodeElement(_bg.RecalcAlways, _aga)
	}
	if _bg.AutoScale != _b.ST_TrueFalseBlankUnset {
		_cc := _ef.StartElement{Name: _ef.Name{Local: "x\u003a\u0041\u0075\u0074\u006f\u0053\u0063\u0061\u006c\u0065"}}
		e.EncodeElement(_bg.AutoScale, _cc)
	}
	if _bg.DDE != _b.ST_TrueFalseBlankUnset {
		_da := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0044D\u0045"}}
		e.EncodeElement(_bg.DDE, _da)
	}
	if _bg.UIObj != _b.ST_TrueFalseBlankUnset {
		_dg := _ef.StartElement{Name: _ef.Name{Local: "\u0078:\u0055\u0049\u004f\u0062\u006a"}}
		e.EncodeElement(_bg.UIObj, _dg)
	}
	if _bg.ScriptText != nil {
		_bgf := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0053c\u0072\u0069\u0070\u0074\u0054\u0065\u0078\u0074"}}
		_f.AddPreserveSpaceAttr(&_bgf, *_bg.ScriptText)
		e.EncodeElement(_bg.ScriptText, _bgf)
	}
	if _bg.ScriptExtended != nil {
		_gge := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003aS\u0063\u0072\u0069p\u0074\u0045\u0078\u0074\u0065\u006e\u0064\u0065\u0064"}}
		_f.AddPreserveSpaceAttr(&_gge, *_bg.ScriptExtended)
		e.EncodeElement(_bg.ScriptExtended, _gge)
	}
	if _bg.ScriptLanguage != nil {
		_dga := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003aS\u0063\u0072\u0069p\u0074\u004c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}}
		e.EncodeElement(_bg.ScriptLanguage, _dga)
	}
	if _bg.ScriptLocation != nil {
		_caa := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003aS\u0063\u0072\u0069p\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"}}
		e.EncodeElement(_bg.ScriptLocation, _caa)
	}
	if _bg.FmlaTxbx != nil {
		_bcf := _ef.StartElement{Name: _ef.Name{Local: "\u0078\u003a\u0046\u006d\u006c\u0061\u0054\u0078\u0062\u0078"}}
		_f.AddPreserveSpaceAttr(&_bcf, *_bg.FmlaTxbx)
		e.EncodeElement(_bg.FmlaTxbx, _bcf)
	}
	e.EncodeToken(_ef.EndElement{Name: start.Name})
	return nil
}

const (
	ST_ObjectTypeUnset    ST_ObjectType = 0
	ST_ObjectTypeButton   ST_ObjectType = 1
	ST_ObjectTypeCheckbox ST_ObjectType = 2
	ST_ObjectTypeDialog   ST_ObjectType = 3
	ST_ObjectTypeDrop     ST_ObjectType = 4
	ST_ObjectTypeEdit     ST_ObjectType = 5
	ST_ObjectTypeGBox     ST_ObjectType = 6
	ST_ObjectTypeLabel    ST_ObjectType = 7
	ST_ObjectTypeLineA    ST_ObjectType = 8
	ST_ObjectTypeList     ST_ObjectType = 9
	ST_ObjectTypeMovie    ST_ObjectType = 10
	ST_ObjectTypeNote     ST_ObjectType = 11
	ST_ObjectTypePict     ST_ObjectType = 12
	ST_ObjectTypeRadio    ST_ObjectType = 13
	ST_ObjectTypeRectA    ST_ObjectType = 14
	ST_ObjectTypeScroll   ST_ObjectType = 15
	ST_ObjectTypeSpin     ST_ObjectType = 16
	ST_ObjectTypeShape    ST_ObjectType = 17
	ST_ObjectTypeGroup    ST_ObjectType = 18
	ST_ObjectTypeRect     ST_ObjectType = 19
)

func init() {
	_f.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", "\u0043\u0054\u005f\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061", NewCT_ClientData)
	_f.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", "\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061", NewClientData)
}
