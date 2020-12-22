package diagram

import (
	_d "encoding/xml"
	_c "fmt"
	_ae "strconv"

	_af "gitee.com/gooffice/gooffice"
	_ca "gitee.com/gooffice/gooffice/schema/soo/dml"
)

func (_beb *CT_CTName) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gfg := range start.Attr {
		if _gfg.Name.Local == "\u006c\u0061\u006e\u0067" {
			_fca, _gfde := _gfg.Value, error(nil)
			if _gfde != nil {
				return _gfde
			}
			_beb.LangAttr = &_fca
			continue
		}
		if _gfg.Name.Local == "\u0076\u0061\u006c" {
			_gdd, _ebgc := _gfg.Value, error(nil)
			if _ebgc != nil {
				return _ebgc
			}
			_beb.ValAttr = _gdd
			continue
		}
	}
	for {
		_bedf, _debb := d.Token()
		if _debb != nil {
			return _c.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0043\u0054\u004ea\u006d\u0065\u003a\u0020\u0025\u0073", _debb)
		}
		if _ddd, _ecgf := _bedf.(_d.EndElement); _ecgf && _ddd.Name == start.Name {
			break
		}
	}
	return nil
}
func (_daeeda ST_ConstraintRelationship) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_daeeda.String(), start)
}
func NewAG_ConstraintRefAttributes() *AG_ConstraintRefAttributes {
	_be := &AG_ConstraintRefAttributes{}
	return _be
}
func (_gdgfe ST_BoolOperator) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_gdgfe.String(), start)
}
func NewCT_Description() *CT_Description { _feca := &CT_Description{}; return _feca }
func (_cedb ST_LinearDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_cedb.String(), start)
}
func (_fgdc *CT_ResizeHandles) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _fgdc.ValAttr != ST_ResizeHandlesStrUnset {
		_abe, _bdge := _fgdc.ValAttr.MarshalXMLAttr(_d.Name{Local: "\u0076\u0061\u006c"})
		if _bdge != nil {
			return _bdge
		}
		start.Attr = append(start.Attr, _abe)
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_adba *ST_HierBranchStyle) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_adba = 0
	case "\u006c":
		*_adba = 1
	case "\u0072":
		*_adba = 2
	case "\u0068\u0061\u006e\u0067":
		*_adba = 3
	case "\u0073\u0074\u0064":
		*_adba = 4
	case "\u0069\u006e\u0069\u0074":
		*_adba = 5
	}
	return nil
}
func (_acbe ST_ConnectorRouting) ValidateWithPath(path string) error {
	switch _acbe {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_acbe))
	}
	return nil
}

// Validate validates the CT_ColorTransformHeader and its children
func (_cdf *CT_ColorTransformHeader) Validate() error {
	return _cdf.ValidateWithPath("\u0043\u0054\u005fCo\u006c\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065\u0072")
}
func (_acc *CT_CTCategories) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _acc.Cat != nil {
		_gda := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074"}}
		for _, _ege := range _acc.Cat {
			e.EncodeElement(_ege, _gda)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_egcbd *ST_ConstraintType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_egcbd = 0
	case "\u006e\u006f\u006e\u0065":
		*_egcbd = 1
	case "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066":
		*_egcbd = 2
	case "\u0062e\u0067\u004d\u0061\u0072\u0067":
		*_egcbd = 3
	case "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074":
		*_egcbd = 4
	case "\u0062\u0065\u0067\u0050\u0061\u0064":
		*_egcbd = 5
	case "\u0062":
		*_egcbd = 6
	case "\u0062\u004d\u0061r\u0067":
		*_egcbd = 7
	case "\u0062\u004f\u0066\u0066":
		*_egcbd = 8
	case "\u0063\u0074\u0072\u0058":
		*_egcbd = 9
	case "\u0063t\u0072\u0058\u004f\u0066\u0066":
		*_egcbd = 10
	case "\u0063\u0074\u0072\u0059":
		*_egcbd = 11
	case "\u0063t\u0072\u0059\u004f\u0066\u0066":
		*_egcbd = 12
	case "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074":
		*_egcbd = 13
	case "\u0064\u0069\u0061\u006d":
		*_egcbd = 14
	case "\u0065n\u0064\u004d\u0061\u0072\u0067":
		*_egcbd = 15
	case "\u0065\u006e\u0064\u0050\u0061\u0064":
		*_egcbd = 16
	case "\u0068":
		*_egcbd = 17
	case "\u0068\u0041\u0072\u0048":
		*_egcbd = 18
	case "\u0068\u004f\u0066\u0066":
		*_egcbd = 19
	case "\u006c":
		*_egcbd = 20
	case "\u006c\u004d\u0061r\u0067":
		*_egcbd = 21
	case "\u006c\u004f\u0066\u0066":
		*_egcbd = 22
	case "\u0072":
		*_egcbd = 23
	case "\u0072\u004d\u0061r\u0067":
		*_egcbd = 24
	case "\u0072\u004f\u0066\u0066":
		*_egcbd = 25
	case "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a":
		*_egcbd = 26
	case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f":
		*_egcbd = 27
	case "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z":
		*_egcbd = 28
	case "\u0073\u0069\u0062S\u0070":
		*_egcbd = 29
	case "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070":
		*_egcbd = 30
	case "\u0073\u0070":
		*_egcbd = 31
	case "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k":
		*_egcbd = 32
	case "\u0074":
		*_egcbd = 33
	case "\u0074\u004d\u0061r\u0067":
		*_egcbd = 34
	case "\u0074\u004f\u0066\u0066":
		*_egcbd = 35
	case "\u0075\u0073\u0065r\u0041":
		*_egcbd = 36
	case "\u0075\u0073\u0065r\u0042":
		*_egcbd = 37
	case "\u0075\u0073\u0065r\u0043":
		*_egcbd = 38
	case "\u0075\u0073\u0065r\u0044":
		*_egcbd = 39
	case "\u0075\u0073\u0065r\u0045":
		*_egcbd = 40
	case "\u0075\u0073\u0065r\u0046":
		*_egcbd = 41
	case "\u0075\u0073\u0065r\u0047":
		*_egcbd = 42
	case "\u0075\u0073\u0065r\u0048":
		*_egcbd = 43
	case "\u0075\u0073\u0065r\u0049":
		*_egcbd = 44
	case "\u0075\u0073\u0065r\u004a":
		*_egcbd = 45
	case "\u0075\u0073\u0065r\u004b":
		*_egcbd = 46
	case "\u0075\u0073\u0065r\u004c":
		*_egcbd = 47
	case "\u0075\u0073\u0065r\u004d":
		*_egcbd = 48
	case "\u0075\u0073\u0065r\u004e":
		*_egcbd = 49
	case "\u0075\u0073\u0065r\u004f":
		*_egcbd = 50
	case "\u0075\u0073\u0065r\u0050":
		*_egcbd = 51
	case "\u0075\u0073\u0065r\u0051":
		*_egcbd = 52
	case "\u0075\u0073\u0065r\u0052":
		*_egcbd = 53
	case "\u0075\u0073\u0065r\u0053":
		*_egcbd = 54
	case "\u0075\u0073\u0065r\u0054":
		*_egcbd = 55
	case "\u0075\u0073\u0065r\u0055":
		*_egcbd = 56
	case "\u0075\u0073\u0065r\u0056":
		*_egcbd = 57
	case "\u0075\u0073\u0065r\u0057":
		*_egcbd = 58
	case "\u0075\u0073\u0065r\u0058":
		*_egcbd = 59
	case "\u0075\u0073\u0065r\u0059":
		*_egcbd = 60
	case "\u0075\u0073\u0065r\u005a":
		*_egcbd = 61
	case "\u0077":
		*_egcbd = 62
	case "\u0077\u0041\u0072\u0048":
		*_egcbd = 63
	case "\u0077\u004f\u0066\u0066":
		*_egcbd = 64
	}
	return nil
}
func (_gdddd ST_PyramidAccentTextMargin) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_gdddd.String(), start)
}

type CT_Colors struct {
	MethAttr       ST_ClrAppMethod
	HueDirAttr     ST_HueDir
	EG_ColorChoice []*_ca.EG_ColorChoice
}

const (
	ST_BreakpointUnset  ST_Breakpoint = 0
	ST_BreakpointEndCnv ST_Breakpoint = 1
	ST_BreakpointBal    ST_Breakpoint = 2
	ST_BreakpointFixed  ST_Breakpoint = 3
)

func (_acfb ST_HueDir) Validate() error { return _acfb.ValidateWithPath("") }

// Validate validates the CT_CTStyleLabel and its children
func (_efb *CT_CTStyleLabel) Validate() error {
	return _efb.ValidateWithPath("\u0043T\u005fC\u0054\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c")
}
func (_bfbd *ST_PrSetCustVal) Validate() error { return _bfbd.ValidateWithPath("") }

// ValidateWithPath validates the ColorsDefHdrLst and its children, prefixing error messages with path
func (_febbf *ColorsDefHdrLst) ValidateWithPath(path string) error {
	if _accgc := _febbf.CT_ColorTransformHeaderLst.ValidateWithPath(path); _accgc != nil {
		return _accgc
	}
	return nil
}
func (_bbff *ST_FunctionType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bbff = 0
	case "\u0063\u006e\u0074":
		*_bbff = 1
	case "\u0070\u006f\u0073":
		*_bbff = 2
	case "\u0072\u0065\u0076\u0050\u006f\u0073":
		*_bbff = 3
	case "\u0070o\u0073\u0045\u0076\u0065\u006e":
		*_bbff = 4
	case "\u0070\u006f\u0073\u004f\u0064\u0064":
		*_bbff = 5
	case "\u0076\u0061\u0072":
		*_bbff = 6
	case "\u0064\u0065\u0070t\u0068":
		*_bbff = 7
	case "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068":
		*_bbff = 8
	}
	return nil
}
func (_caba ST_HueDir) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ggba := _d.Attr{}
	_ggba.Name = name
	switch _caba {
	case ST_HueDirUnset:
		_ggba.Value = ""
	case ST_HueDirCw:
		_ggba.Value = "\u0063\u0077"
	case ST_HueDirCcw:
		_ggba.Value = "\u0063\u0063\u0077"
	}
	return _ggba, nil
}
func (_cdde ST_FunctionOperator) Validate() error { return _cdde.ValidateWithPath("") }
func (_dfbf ST_FunctionType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ddbf := _d.Attr{}
	_ddbf.Name = name
	switch _dfbf {
	case ST_FunctionTypeUnset:
		_ddbf.Value = ""
	case ST_FunctionTypeCnt:
		_ddbf.Value = "\u0063\u006e\u0074"
	case ST_FunctionTypePos:
		_ddbf.Value = "\u0070\u006f\u0073"
	case ST_FunctionTypeRevPos:
		_ddbf.Value = "\u0072\u0065\u0076\u0050\u006f\u0073"
	case ST_FunctionTypePosEven:
		_ddbf.Value = "\u0070o\u0073\u0045\u0076\u0065\u006e"
	case ST_FunctionTypePosOdd:
		_ddbf.Value = "\u0070\u006f\u0073\u004f\u0064\u0064"
	case ST_FunctionTypeVar:
		_ddbf.Value = "\u0076\u0061\u0072"
	case ST_FunctionTypeDepth:
		_ddbf.Value = "\u0064\u0065\u0070t\u0068"
	case ST_FunctionTypeMaxDepth:
		_ddbf.Value = "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068"
	}
	return _ddbf, nil
}
func (_cccb *ST_Offset) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_cccb = 0
	case "\u0063\u0074\u0072":
		*_cccb = 1
	case "\u006f\u0066\u0066":
		*_cccb = 2
	}
	return nil
}
func (_bbef *CT_RelIds) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u003a\u0064\u006d"}, Value: _c.Sprintf("\u0025\u0076", _bbef.DmAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u003a\u006c\u006f"}, Value: _c.Sprintf("\u0025\u0076", _bbef.LoAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u003a\u0071\u0073"}, Value: _c.Sprintf("\u0025\u0076", _bbef.QsAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u003a\u0063\u0073"}, Value: _c.Sprintf("\u0025\u0076", _bbef.CsAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_ebaf *ST_Direction) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ebaf = 0
	case "\u006e\u006f\u0072\u006d":
		*_ebaf = 1
	case "\u0072\u0065\u0076":
		*_ebaf = 2
	}
	return nil
}

// ValidateWithPath validates the CT_Direction and its children, prefixing error messages with path
func (_feaee *CT_Direction) ValidateWithPath(path string) error {
	if _gdbe := _feaee.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gdbe != nil {
		return _gdbe
	}
	return nil
}
func (_gaeg *CT_StyleDefinition) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _gaeg.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_gaeg.UniqueIdAttr)})
	}
	if _gaeg.MinVerAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _c.Sprintf("\u0025\u0076", *_gaeg.MinVerAttr)})
	}
	e.EncodeToken(start)
	if _gaeg.Title != nil {
		_gdbcgc := _d.StartElement{Name: _d.Name{Local: "\u0074\u0069\u0074l\u0065"}}
		for _, _ggcadc := range _gaeg.Title {
			e.EncodeElement(_ggcadc, _gdbcgc)
		}
	}
	if _gaeg.Desc != nil {
		_abcb := _d.StartElement{Name: _d.Name{Local: "\u0064\u0065\u0073\u0063"}}
		for _, _dgeb := range _gaeg.Desc {
			e.EncodeElement(_dgeb, _abcb)
		}
	}
	if _gaeg.CatLst != nil {
		_dfcd := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_gaeg.CatLst, _dfcd)
	}
	if _gaeg.Scene3d != nil {
		_acda := _d.StartElement{Name: _d.Name{Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}}
		e.EncodeElement(_gaeg.Scene3d, _acda)
	}
	_efaee := _d.StartElement{Name: _d.Name{Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}}
	for _, _dcefd := range _gaeg.StyleLbl {
		e.EncodeElement(_dcefd, _efaee)
	}
	if _gaeg.ExtLst != nil {
		_bfaf := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_gaeg.ExtLst, _bfaf)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

const (
	ST_AnimOneStrUnset  ST_AnimOneStr = 0
	ST_AnimOneStrNone   ST_AnimOneStr = 1
	ST_AnimOneStrOne    ST_AnimOneStr = 2
	ST_AnimOneStrBranch ST_AnimOneStr = 3
)

func ParseUnionST_ParameterVal(s string) (ST_ParameterVal, error) { return ST_ParameterVal{}, nil }
func (_fada *CT_ColorTransformHeader) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", _fada.UniqueIdAttr)})
	if _fada.MinVerAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _c.Sprintf("\u0025\u0076", *_fada.MinVerAttr)})
	}
	if _fada.ResIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u0065\u0073I\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fada.ResIdAttr)})
	}
	e.EncodeToken(start)
	_acce := _d.StartElement{Name: _d.Name{Local: "\u0074\u0069\u0074l\u0065"}}
	for _, _efbf := range _fada.Title {
		e.EncodeElement(_efbf, _acce)
	}
	_bfe := _d.StartElement{Name: _d.Name{Local: "\u0064\u0065\u0073\u0063"}}
	for _, _ecgg := range _fada.Desc {
		e.EncodeElement(_ecgg, _bfe)
	}
	if _fada.CatLst != nil {
		_agae := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_fada.CatLst, _agae)
	}
	if _fada.ExtLst != nil {
		_daed := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_fada.ExtLst, _daed)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_effdce ST_BoolOperator) String() string {
	switch _effdce {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0065\u0071\u0075"
	case 3:
		return "\u0067\u0074\u0065"
	case 4:
		return "\u006c\u0074\u0065"
	}
	return ""
}
func (_ebcce *ST_VerticalAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ddfcf, _acdfg := d.Token()
	if _acdfg != nil {
		return _acdfg
	}
	if _cddac, _begdb := _ddfcf.(_d.EndElement); _begdb && _cddac.Name == start.Name {
		*_ebcce = 1
		return nil
	}
	if _efdf, _bbab := _ddfcf.(_d.CharData); !_bbab {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ddfcf)
	} else {
		switch string(_efdf) {
		case "":
			*_ebcce = 0
		case "\u0074":
			*_ebcce = 1
		case "\u006d\u0069\u0064":
			*_ebcce = 2
		case "\u0062":
			*_ebcce = 3
		case "\u006e\u006f\u006e\u0065":
			*_ebcce = 4
		}
	}
	_ddfcf, _acdfg = d.Token()
	if _acdfg != nil {
		return _acdfg
	}
	if _bgfaf, _bagbed := _ddfcf.(_d.EndElement); _bagbed && _bgfaf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ddfcf)
}
func (_gbdf ST_TextBlockDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_gbdf.String(), start)
}

const (
	ST_ConnectorPointUnset  ST_ConnectorPoint = 0
	ST_ConnectorPointAuto   ST_ConnectorPoint = 1
	ST_ConnectorPointBCtr   ST_ConnectorPoint = 2
	ST_ConnectorPointCtr    ST_ConnectorPoint = 3
	ST_ConnectorPointMidL   ST_ConnectorPoint = 4
	ST_ConnectorPointMidR   ST_ConnectorPoint = 5
	ST_ConnectorPointTCtr   ST_ConnectorPoint = 6
	ST_ConnectorPointBL     ST_ConnectorPoint = 7
	ST_ConnectorPointBR     ST_ConnectorPoint = 8
	ST_ConnectorPointTL     ST_ConnectorPoint = 9
	ST_ConnectorPointTR     ST_ConnectorPoint = 10
	ST_ConnectorPointRadial ST_ConnectorPoint = 11
)

// Validate validates the CT_Parameter and its children
func (_gbcfe *CT_Parameter) Validate() error {
	return _gbcfe.ValidateWithPath("\u0043\u0054\u005fP\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072")
}

type ST_Booleans []bool

func (_egbd *CT_Category) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", _egbd.TypeAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0069"}, Value: _c.Sprintf("\u0025\u0076", _egbd.PriAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewCT_Cxn() *CT_Cxn { _dbf := &CT_Cxn{}; return _dbf }

// Validate validates the CT_Name and its children
func (_dece *CT_Name) Validate() error {
	return _dece.ValidateWithPath("\u0043T\u005f\u004e\u0061\u006d\u0065")
}
func (_fggbb *CT_Rules) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_gdab:
	for {
		_acba, _effb := d.Token()
		if _effb != nil {
			return _effb
		}
		switch _aagcg := _acba.(type) {
		case _d.StartElement:
			switch _aagcg.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072\u0075\u006c\u0065"}:
				_abcg := NewCT_NumericRule()
				if _adcd := d.DecodeElement(_abcg, &_aagcg); _adcd != nil {
					return _adcd
				}
				_fggbb.Rule = append(_fggbb.Rule, _abcg)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0075\u006c\u0065\u0073\u0020\u0025\u0076", _aagcg.Name)
				if _decf := d.Skip(); _decf != nil {
					return _decf
				}
			}
		case _d.EndElement:
			break _gdab
		case _d.CharData:
		}
	}
	return nil
}
func NewCT_Shape() *CT_Shape { _efbd := &CT_Shape{}; return _efbd }
func (_bbg *CT_AnimOne) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _bbg.ValAttr != ST_AnimOneStrUnset {
		_aec, _eeb := _bbg.ValAttr.MarshalXMLAttr(_d.Name{Local: "\u0076\u0061\u006c"})
		if _eeb != nil {
			return _eeb
		}
		start.Attr = append(start.Attr, _aec)
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_When and its children, prefixing error messages with path
func (_cgfe *CT_When) ValidateWithPath(path string) error {
	if _cgfe.FuncAttr == ST_FunctionTypeUnset {
		return _c.Errorf("\u0025\u0073\u002f\u0046\u0075\u006e\u0063\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020a\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _bbb := _cgfe.FuncAttr.ValidateWithPath(path + "\u002fF\u0075\u006e\u0063\u0041\u0074\u0074r"); _bbb != nil {
		return _bbb
	}
	if _cgfe.ArgAttr != nil {
		if _babfd := _cgfe.ArgAttr.ValidateWithPath(path + "\u002f\u0041\u0072\u0067\u0041\u0074\u0074\u0072"); _babfd != nil {
			return _babfd
		}
	}
	if _cgfe.OpAttr == ST_FunctionOperatorUnset {
		return _c.Errorf("\u0025\u0073\u002f\u004f\u0070\u0041\u0074\u0074\u0072\u0020i\u0073\u0020\u0061\u0020\u006d\u0061\u006ed\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _aebb := _cgfe.OpAttr.ValidateWithPath(path + "\u002fO\u0070\u0041\u0074\u0074\u0072"); _aebb != nil {
		return _aebb
	}
	if _ecga := _cgfe.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _ecga != nil {
		return _ecga
	}
	for _edfc, _cgcdf := range _cgfe.Alg {
		if _baaf := _cgcdf.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _edfc)); _baaf != nil {
			return _baaf
		}
	}
	for _abbad, _dbcgc := range _cgfe.Shape {
		if _cbgbf := _dbcgc.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _abbad)); _cbgbf != nil {
			return _cbgbf
		}
	}
	for _ecgde, _cdaf := range _cgfe.PresOf {
		if _gcca := _cdaf.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _ecgde)); _gcca != nil {
			return _gcca
		}
	}
	for _bfdc, _bcec := range _cgfe.ConstrLst {
		if _affd := _bcec.ValidateWithPath(_c.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _bfdc)); _affd != nil {
			return _affd
		}
	}
	for _befae, _fbdfd := range _cgfe.RuleLst {
		if _ebfg := _fbdfd.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _befae)); _ebfg != nil {
			return _ebfg
		}
	}
	for _ffee, _beec := range _cgfe.ForEach {
		if _gbfge := _beec.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _ffee)); _gbfge != nil {
			return _gbfge
		}
	}
	for _cacb, _afgea := range _cgfe.LayoutNode {
		if _befbf := _afgea.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _cacb)); _befbf != nil {
			return _befbf
		}
	}
	for _bcga, _afde := range _cgfe.Choose {
		if _cffd := _afde.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _bcga)); _cffd != nil {
			return _cffd
		}
	}
	for _gge, _fgfa := range _cgfe.ExtLst {
		if _cfef := _fgfa.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _gge)); _cfef != nil {
			return _cfef
		}
	}
	return nil
}
func NewCT_DiagramDefinition() *CT_DiagramDefinition {
	_gfbe := &CT_DiagramDefinition{}
	_gfbe.LayoutNode = NewCT_LayoutNode()
	return _gfbe
}
func NewCT_ColorTransform() *CT_ColorTransform { _dee := &CT_ColorTransform{}; return _dee }
func (_dggc *CT_DiagramDefinitionHeaderLst) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_dgagd:
	for {
		_dba, _bbaf := d.Token()
		if _bbaf != nil {
			return _bbaf
		}
		switch _cbgb := _dba.(type) {
		case _d.StartElement:
			switch _cbgb.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_bfae := NewCT_DiagramDefinitionHeader()
				if _cdbbe := d.DecodeElement(_bfae, &_cbgb); _cdbbe != nil {
					return _cdbbe
				}
				_dggc.LayoutDefHdr = append(_dggc.LayoutDefHdr, _bfae)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074 \u006f\u006e\u0020\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006ei\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065r\u004c\u0073t\u0020%\u0076", _cbgb.Name)
				if _facb := d.Skip(); _facb != nil {
					return _facb
				}
			}
		case _d.EndElement:
			break _dgagd
		case _d.CharData:
		}
	}
	return nil
}
func (_fdcfe *ST_PtType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gbeg, _cgeb := d.Token()
	if _cgeb != nil {
		return _cgeb
	}
	if _efbgc, _ebcfd := _gbeg.(_d.EndElement); _ebcfd && _efbgc.Name == start.Name {
		*_fdcfe = 1
		return nil
	}
	if _cgde, _fbag := _gbeg.(_d.CharData); !_fbag {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbeg)
	} else {
		switch string(_cgde) {
		case "":
			*_fdcfe = 0
		case "\u006e\u006f\u0064\u0065":
			*_fdcfe = 1
		case "\u0061\u0073\u0073\u0074":
			*_fdcfe = 2
		case "\u0064\u006f\u0063":
			*_fdcfe = 3
		case "\u0070\u0072\u0065\u0073":
			*_fdcfe = 4
		case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
			*_fdcfe = 5
		case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
			*_fdcfe = 6
		}
	}
	_gbeg, _cgeb = d.Token()
	if _cgeb != nil {
		return _cgeb
	}
	if _eafg, _fecf := _gbeg.(_d.EndElement); _fecf && _eafg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbeg)
}

const (
	ST_FunctionOperatorUnset ST_FunctionOperator = 0
	ST_FunctionOperatorEqu   ST_FunctionOperator = 1
	ST_FunctionOperatorNeq   ST_FunctionOperator = 2
	ST_FunctionOperatorGt    ST_FunctionOperator = 3
	ST_FunctionOperatorLt    ST_FunctionOperator = 4
	ST_FunctionOperatorGte   ST_FunctionOperator = 5
	ST_FunctionOperatorLte   ST_FunctionOperator = 6
)

func (_bbae ST_ConnectorRouting) String() string {
	switch _bbae {
	case 0:
		return ""
	case 1:
		return "\u0073\u0074\u0072\u0061"
	case 2:
		return "\u0062\u0065\u006e\u0064"
	case 3:
		return "\u0063\u0075\u0072v\u0065"
	case 4:
		return "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e"
	}
	return ""
}
func NewCT_Categories() *CT_Categories { _gcd := &CT_Categories{}; return _gcd }

// ST_ParameterVal is a union type
type ST_ParameterVal struct {
	ST_DiagramHorizontalAlignment ST_DiagramHorizontalAlignment
	ST_VerticalAlignment          ST_VerticalAlignment
	ST_ChildDirection             ST_ChildDirection
	ST_ChildAlignment             ST_ChildAlignment
	ST_SecondaryChildAlignment    ST_SecondaryChildAlignment
	ST_LinearDirection            ST_LinearDirection
	ST_SecondaryLinearDirection   ST_SecondaryLinearDirection
	ST_StartingElement            ST_StartingElement
	ST_BendPoint                  ST_BendPoint
	ST_ConnectorRouting           ST_ConnectorRouting
	ST_ArrowheadStyle             ST_ArrowheadStyle
	ST_ConnectorDimension         ST_ConnectorDimension
	ST_RotationPath               ST_RotationPath
	ST_CenterShapeMapping         ST_CenterShapeMapping
	ST_NodeHorizontalAlignment    ST_NodeHorizontalAlignment
	ST_NodeVerticalAlignment      ST_NodeVerticalAlignment
	ST_FallbackDimension          ST_FallbackDimension
	ST_TextDirection              ST_TextDirection
	ST_PyramidAccentPosition      ST_PyramidAccentPosition
	ST_PyramidAccentTextMargin    ST_PyramidAccentTextMargin
	ST_TextBlockDirection         ST_TextBlockDirection
	ST_TextAnchorHorizontal       ST_TextAnchorHorizontal
	ST_TextAnchorVertical         ST_TextAnchorVertical
	ST_DiagramTextAlignment       ST_DiagramTextAlignment
	ST_AutoTextRotation           ST_AutoTextRotation
	ST_GrowDirection              ST_GrowDirection
	ST_FlowDirection              ST_FlowDirection
	ST_ContinueDirection          ST_ContinueDirection
	ST_Breakpoint                 ST_Breakpoint
	ST_Offset                     ST_Offset
	ST_HierarchyAlignment         ST_HierarchyAlignment
	Int32                         *int32
	Float64                       *float64
	Bool                          *bool
	StringVal                     *string
	ST_ConnectorPoint             ST_ConnectorPoint
}

func (_fbad *ST_LinearDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dgdb, _dfgea := d.Token()
	if _dfgea != nil {
		return _dfgea
	}
	if _fabdb, _bfea := _dgdb.(_d.EndElement); _bfea && _fabdb.Name == start.Name {
		*_fbad = 1
		return nil
	}
	if _ddfef, _gdggd := _dgdb.(_d.CharData); !_gdggd {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dgdb)
	} else {
		switch string(_ddfef) {
		case "":
			*_fbad = 0
		case "\u0066\u0072\u006fm\u004c":
			*_fbad = 1
		case "\u0066\u0072\u006fm\u0052":
			*_fbad = 2
		case "\u0066\u0072\u006fm\u0054":
			*_fbad = 3
		case "\u0066\u0072\u006fm\u0042":
			*_fbad = 4
		}
	}
	_dgdb, _dfgea = d.Token()
	if _dfgea != nil {
		return _dfgea
	}
	if _dgbbc, _bgead := _dgdb.(_d.EndElement); _bgead && _dgbbc.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dgdb)
}
func (_beaf ST_CxnType) Validate() error { return _beaf.ValidateWithPath("") }
func (_bdcf ST_TextDirection) String() string {
	switch _bdcf {
	case 0:
		return ""
	case 1:
		return "\u0066\u0072\u006fm\u0054"
	case 2:
		return "\u0066\u0072\u006fm\u0042"
	}
	return ""
}
func (_bdbg *ST_FunctionType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fffcb, _cfcde := d.Token()
	if _cfcde != nil {
		return _cfcde
	}
	if _aabc, _ebfbe := _fffcb.(_d.EndElement); _ebfbe && _aabc.Name == start.Name {
		*_bdbg = 1
		return nil
	}
	if _ebfee, _cfbe := _fffcb.(_d.CharData); !_cfbe {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fffcb)
	} else {
		switch string(_ebfee) {
		case "":
			*_bdbg = 0
		case "\u0063\u006e\u0074":
			*_bdbg = 1
		case "\u0070\u006f\u0073":
			*_bdbg = 2
		case "\u0072\u0065\u0076\u0050\u006f\u0073":
			*_bdbg = 3
		case "\u0070o\u0073\u0045\u0076\u0065\u006e":
			*_bdbg = 4
		case "\u0070\u006f\u0073\u004f\u0064\u0064":
			*_bdbg = 5
		case "\u0076\u0061\u0072":
			*_bdbg = 6
		case "\u0064\u0065\u0070t\u0068":
			*_bdbg = 7
		case "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068":
			*_bdbg = 8
		}
	}
	_fffcb, _cfcde = d.Token()
	if _cfcde != nil {
		return _cfcde
	}
	if _dfcdd, _acge := _fffcb.(_d.EndElement); _acge && _dfcdd.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fffcb)
}

const (
	ST_ElementTypeUnset    ST_ElementType = 0
	ST_ElementTypeAll      ST_ElementType = 1
	ST_ElementTypeDoc      ST_ElementType = 2
	ST_ElementTypeNode     ST_ElementType = 3
	ST_ElementTypeNorm     ST_ElementType = 4
	ST_ElementTypeNonNorm  ST_ElementType = 5
	ST_ElementTypeAsst     ST_ElementType = 6
	ST_ElementTypeNonAsst  ST_ElementType = 7
	ST_ElementTypeParTrans ST_ElementType = 8
	ST_ElementTypePres     ST_ElementType = 9
	ST_ElementTypeSibTrans ST_ElementType = 10
)

func (_efbfb ST_AxisType) Validate() error { return _efbfb.ValidateWithPath("") }

// Validate validates the CT_ChildMax and its children
func (_feg *CT_ChildMax) Validate() error {
	return _feg.ValidateWithPath("C\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u004d\u0061\u0078")
}

// Validate validates the CT_Shape and its children
func (_abac *CT_Shape) Validate() error {
	return _abac.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065")
}

type CT_Constraint struct {
	OpAttr         ST_BoolOperator
	ValAttr        *float64
	FactAttr       *float64
	ExtLst         *_ca.CT_OfficeArtExtensionList
	TypeAttr       ST_ConstraintType
	ForAttr        ST_ConstraintRelationship
	ForNameAttr    *string
	PtTypeAttr     ST_ElementType
	RefTypeAttr    ST_ConstraintType
	RefForAttr     ST_ConstraintRelationship
	RefForNameAttr *string
	RefPtTypeAttr  ST_ElementType
}

func (_bgd *CT_CTStyleLabel) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gba := range start.Attr {
		if _gba.Name.Local == "\u006e\u0061\u006d\u0065" {
			_dag, _bgc := _gba.Value, error(nil)
			if _bgc != nil {
				return _bgc
			}
			_bgd.NameAttr = _dag
			continue
		}
	}
_bga:
	for {
		_ddb, _ceb := d.Token()
		if _ceb != nil {
			return _ceb
		}
		switch _bdb := _ddb.(type) {
		case _d.StartElement:
			switch _bdb.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bgd.FillClrLst = NewCT_Colors()
				if _aac := d.DecodeElement(_bgd.FillClrLst, &_bdb); _aac != nil {
					return _aac
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006ci\u006e\u0043\u006c\u0072\u004c\u0073t"}:
				_bgd.LinClrLst = NewCT_Colors()
				if _gfcd := d.DecodeElement(_bgd.LinClrLst, &_bdb); _gfcd != nil {
					return _gfcd
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0066\u0066e\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bgd.EffectClrLst = NewCT_Colors()
				if _bdde := d.DecodeElement(_bgd.EffectClrLst, &_bdb); _bdde != nil {
					return _bdde
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "t\u0078\u004c\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bgd.TxLinClrLst = NewCT_Colors()
				if _adgg := d.DecodeElement(_bgd.TxLinClrLst, &_bdb); _adgg != nil {
					return _adgg
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0078\u0046i\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bgd.TxFillClrLst = NewCT_Colors()
				if _agc := d.DecodeElement(_bgd.TxFillClrLst, &_bdb); _agc != nil {
					return _agc
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0078\u0045\u0066\u0066\u0065\u0063\u0074\u0043l\u0072\u004c\u0073\u0074"}:
				_bgd.TxEffectClrLst = NewCT_Colors()
				if _dac := d.DecodeElement(_bgd.TxEffectClrLst, &_bdb); _dac != nil {
					return _dac
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bgd.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _feb := d.DecodeElement(_bgd.ExtLst, &_bdb); _feb != nil {
					return _feb
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u0054\u0053t\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c\u0020\u0025\u0076", _bdb.Name)
				if _ccd := d.Skip(); _ccd != nil {
					return _ccd
				}
			}
		case _d.EndElement:
			break _bga
		case _d.CharData:
		}
	}
	return nil
}
func (_gcgbg *ST_TextAnchorHorizontal) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cbddg, _bfcg := d.Token()
	if _bfcg != nil {
		return _bfcg
	}
	if _gcaeg, _aaafb := _cbddg.(_d.EndElement); _aaafb && _gcaeg.Name == start.Name {
		*_gcgbg = 1
		return nil
	}
	if _cbcbg, _ebca := _cbddg.(_d.CharData); !_ebca {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cbddg)
	} else {
		switch string(_cbcbg) {
		case "":
			*_gcgbg = 0
		case "\u006e\u006f\u006e\u0065":
			*_gcgbg = 1
		case "\u0063\u0074\u0072":
			*_gcgbg = 2
		}
	}
	_cbddg, _bfcg = d.Token()
	if _bfcg != nil {
		return _bfcg
	}
	if _caeee, _bfeeb := _cbddg.(_d.EndElement); _bfeeb && _caeee.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cbddg)
}

type ST_FunctionType byte
type ST_TextAnchorVertical byte

func (_agbgfc ST_ResizeHandlesStr) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_agbgfc.String(), start)
}
func (_ggbb *ST_HierarchyAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cgeg, _ddffg := d.Token()
	if _ddffg != nil {
		return _ddffg
	}
	if _bgdfe, _egac := _cgeg.(_d.EndElement); _egac && _bgdfe.Name == start.Name {
		*_ggbb = 1
		return nil
	}
	if _aefgf, _bdgfa := _cgeg.(_d.CharData); !_bdgfa {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cgeg)
	} else {
		switch string(_aefgf) {
		case "":
			*_ggbb = 0
		case "\u0074\u004c":
			*_ggbb = 1
		case "\u0074\u0052":
			*_ggbb = 2
		case "\u0074\u0043\u0074\u0072\u0043\u0068":
			*_ggbb = 3
		case "\u0074C\u0074\u0072\u0044\u0065\u0073":
			*_ggbb = 4
		case "\u0062\u004c":
			*_ggbb = 5
		case "\u0062\u0052":
			*_ggbb = 6
		case "\u0062\u0043\u0074\u0072\u0043\u0068":
			*_ggbb = 7
		case "\u0062C\u0074\u0072\u0044\u0065\u0073":
			*_ggbb = 8
		case "\u006c\u0054":
			*_ggbb = 9
		case "\u006c\u0042":
			*_ggbb = 10
		case "\u006c\u0043\u0074\u0072\u0043\u0068":
			*_ggbb = 11
		case "\u006cC\u0074\u0072\u0044\u0065\u0073":
			*_ggbb = 12
		case "\u0072\u0054":
			*_ggbb = 13
		case "\u0072\u0042":
			*_ggbb = 14
		case "\u0072\u0043\u0074\u0072\u0043\u0068":
			*_ggbb = 15
		case "\u0072C\u0074\u0072\u0044\u0065\u0073":
			*_ggbb = 16
		}
	}
	_cgeg, _ddffg = d.Token()
	if _ddffg != nil {
		return _ddffg
	}
	if _affa, _ggfa := _cgeg.(_d.EndElement); _ggfa && _affa.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cgeg)
}
func (_bfegd *ST_TextBlockDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bfegd = 0
	case "\u0068\u006f\u0072\u007a":
		*_bfegd = 1
	case "\u0076\u0065\u0072\u0074":
		*_bfegd = 2
	}
	return nil
}
func (_dbeb *CT_SDCategory) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _acec := range start.Attr {
		if _acec.Name.Local == "\u0074\u0079\u0070\u0065" {
			_fggge, _gcfee := _acec.Value, error(nil)
			if _gcfee != nil {
				return _gcfee
			}
			_dbeb.TypeAttr = _fggge
			continue
		}
		if _acec.Name.Local == "\u0070\u0072\u0069" {
			_afdbg, _eaaf := _ae.ParseUint(_acec.Value, 10, 32)
			if _eaaf != nil {
				return _eaaf
			}
			_dbeb.PriAttr = uint32(_afdbg)
			continue
		}
	}
	for {
		_eafc, _cdbae := d.Token()
		if _cdbae != nil {
			return _c.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0053D\u0043a\u0074e\u0067\u006f\u0072\u0079\u003a\u0020\u0025s", _cdbae)
		}
		if _adgde, _fecd := _eafc.(_d.EndElement); _fecd && _adgde.Name == start.Name {
			break
		}
	}
	return nil
}
func (_afgef ST_SecondaryChildAlignment) String() string {
	switch _afgef {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0074"
	case 3:
		return "\u0062"
	case 4:
		return "\u006c"
	case 5:
		return "\u0072"
	}
	return ""
}
func (_ga *CT_Adj) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ga.IdxAttr = 1
	for _, _cdc := range start.Attr {
		if _cdc.Name.Local == "\u0069\u0064\u0078" {
			_ebd, _aaf := _ae.ParseUint(_cdc.Value, 10, 32)
			if _aaf != nil {
				return _aaf
			}
			_ga.IdxAttr = uint32(_ebd)
			continue
		}
		if _cdc.Name.Local == "\u0076\u0061\u006c" {
			_bfgd, _ad := _ae.ParseFloat(_cdc.Value, 64)
			if _ad != nil {
				return _ad
			}
			_ga.ValAttr = _bfgd
			continue
		}
	}
	for {
		_fee, _bag := d.Token()
		if _bag != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0041d\u006a\u003a\u0020\u0025\u0073", _bag)
		}
		if _cbb, _afba := _fee.(_d.EndElement); _afba && _cbb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_aggg ST_PyramidAccentTextMargin) Validate() error { return _aggg.ValidateWithPath("") }

const (
	ST_AnimLvlStrUnset ST_AnimLvlStr = 0
	ST_AnimLvlStrNone  ST_AnimLvlStr = 1
	ST_AnimLvlStrLvl   ST_AnimLvlStr = 2
	ST_AnimLvlStrCtr   ST_AnimLvlStr = 3
)

func (_cfc *CT_ColorTransformHeaderLst) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_fcd:
	for {
		_ebcd, _ccfg := d.Token()
		if _ccfg != nil {
			return _ccfg
		}
		switch _gdcc := _ebcd.(type) {
		case _d.StartElement:
			switch _gdcc.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_effe := NewCT_ColorTransformHeader()
				if _fbbf := d.DecodeElement(_effe, &_gdcc); _fbbf != nil {
					return _fbbf
				}
				_cfc.ColorsDefHdr = append(_cfc.ColorsDefHdr, _effe)
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020o\u006e\u0020\u0043\u0054_\u0043\u006fl\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _gdcc.Name)
				if _bcge := d.Skip(); _bcge != nil {
					return _bcge
				}
			}
		case _d.EndElement:
			break _fcd
		case _d.CharData:
		}
	}
	return nil
}
func (_bdafgf ST_CenterShapeMapping) ValidateWithPath(path string) error {
	switch _bdafgf {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bdafgf))
	}
	return nil
}

// Validate validates the CT_Colors and its children
func (_dgcg *CT_Colors) Validate() error {
	return _dgcg.ValidateWithPath("\u0043T\u005f\u0043\u006f\u006c\u006f\u0072s")
}
func (_ccdd *CT_ChildMax) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _ccdd.ValAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", *_ccdd.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_cfcee ST_StartingElement) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_cfcee.String(), start)
}

// Validate validates the StyleDef and its children
func (_gcec *StyleDef) Validate() error {
	return _gcec.ValidateWithPath("\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066")
}
func (_acgb *StyleDefHdrLst) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_acgb.CT_StyleDefinitionHeaderLst = *NewCT_StyleDefinitionHeaderLst()
_dbbd:
	for {
		_fedeg, _ddec := d.Token()
		if _ddec != nil {
			return _ddec
		}
		switch _ecdd := _fedeg.(type) {
		case _d.StartElement:
			switch _ecdd.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_fabae := NewCT_StyleDefinitionHeader()
				if _fffff := d.DecodeElement(_fabae, &_ecdd); _fffff != nil {
					return _fffff
				}
				_acgb.StyleDefHdr = append(_acgb.StyleDefHdr, _fabae)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064r\u004cs\u0074\u0020\u0025\u0076", _ecdd.Name)
				if _fcefc := d.Skip(); _fcefc != nil {
					return _fcefc
				}
			}
		case _d.EndElement:
			break _dbbd
		case _d.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DataModel and its children
func (_ebfd *CT_DataModel) Validate() error {
	return _ebfd.ValidateWithPath("\u0043\u0054\u005fD\u0061\u0074\u0061\u004d\u006f\u0064\u0065\u006c")
}

type CT_Rules struct{ Rule []*CT_NumericRule }

func (_edeb ST_ClrAppMethod) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_edeb.String(), start)
}

type ST_ClrAppMethod byte

// Validate validates the CT_ForEach and its children
func (_afcg *CT_ForEach) Validate() error {
	return _afcg.ValidateWithPath("\u0043\u0054\u005f\u0046\u006f\u0072\u0045\u0061\u0063\u0068")
}
func (_g *AG_ConstraintAttributes) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _g.TypeAttr != ST_ConstraintTypeUnset {
		_cb, _e := _g.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _e != nil {
			return _e
		}
		start.Attr = append(start.Attr, _cb)
	}
	if _g.ForAttr != ST_ConstraintRelationshipUnset {
		_b, _eg := _g.ForAttr.MarshalXMLAttr(_d.Name{Local: "\u0066\u006f\u0072"})
		if _eg != nil {
			return _eg
		}
		start.Attr = append(start.Attr, _b)
	}
	if _g.ForNameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0066o\u0072\u004e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_g.ForNameAttr)})
	}
	if _g.PtTypeAttr != ST_ElementTypeUnset {
		_f, _bd := _g.PtTypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"})
		if _bd != nil {
			return _bd
		}
		start.Attr = append(start.Attr, _f)
	}
	return nil
}

type CT_CTCategories struct{ Cat []*CT_CTCategory }
type CT_Otherwise struct {
	NameAttr   *string
	Alg        []*CT_Algorithm
	Shape      []*CT_Shape
	PresOf     []*CT_PresentationOf
	ConstrLst  []*CT_Constraints
	RuleLst    []*CT_Rules
	ForEach    []*CT_ForEach
	LayoutNode []*CT_LayoutNode
	Choose     []*CT_Choose
	ExtLst     []*_ca.CT_OfficeArtExtensionList
}

func (_agcb ST_FunctionOperator) ValidateWithPath(path string) error {
	switch _agcb {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_agcb))
	}
	return nil
}
func (_aba *CT_Direction) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _ccccb := range start.Attr {
		if _ccccb.Name.Local == "\u0076\u0061\u006c" {
			_aba.ValAttr.UnmarshalXMLAttr(_ccccb)
			continue
		}
	}
	for {
		_cacg, _bffc := d.Token()
		if _bffc != nil {
			return _c.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005f\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073", _bffc)
		}
		if _gcgd, _eceb := _cacg.(_d.EndElement); _eceb && _gcgd.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_Direction() *CT_Direction { _feaf := &CT_Direction{}; return _feaf }
func (_gfcca *ST_PyramidAccentPosition) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ecae, _ebgf := d.Token()
	if _ebgf != nil {
		return _ebgf
	}
	if _caaac, _bbdcf := _ecae.(_d.EndElement); _bbdcf && _caaac.Name == start.Name {
		*_gfcca = 1
		return nil
	}
	if _fgfdg, _ffgad := _ecae.(_d.CharData); !_ffgad {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ecae)
	} else {
		switch string(_fgfdg) {
		case "":
			*_gfcca = 0
		case "\u0062\u0065\u0066":
			*_gfcca = 1
		case "\u0061\u0066\u0074":
			*_gfcca = 2
		}
	}
	_ecae, _ebgf = d.Token()
	if _ebgf != nil {
		return _ebgf
	}
	if _gggf, _cffb := _ecae.(_d.EndElement); _cffb && _gggf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ecae)
}

type ST_ConnectorDimension byte

func NewStyleDefHdr() *StyleDefHdr {
	_adbc := &StyleDefHdr{}
	_adbc.CT_StyleDefinitionHeader = *NewCT_StyleDefinitionHeader()
	return _adbc
}
func (_gab *CT_DataModel) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	_bffb := _d.StartElement{Name: _d.Name{Local: "\u0070\u0074\u004cs\u0074"}}
	e.EncodeElement(_gab.PtLst, _bffb)
	if _gab.CxnLst != nil {
		_adce := _d.StartElement{Name: _d.Name{Local: "\u0063\u0078\u006e\u004c\u0073\u0074"}}
		e.EncodeElement(_gab.CxnLst, _adce)
	}
	if _gab.Bg != nil {
		_dafda := _d.StartElement{Name: _d.Name{Local: "\u0062\u0067"}}
		e.EncodeElement(_gab.Bg, _dafda)
	}
	if _gab.Whole != nil {
		_bebb := _d.StartElement{Name: _d.Name{Local: "\u0077\u0068\u006fl\u0065"}}
		e.EncodeElement(_gab.Whole, _bebb)
	}
	if _gab.ExtLst != nil {
		_beaa := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_gab.ExtLst, _beaa)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_gabag *ST_ParameterVal) ValidateWithPath(path string) error {
	_dagd := []string{}
	if _gabag.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		_dagd = append(_dagd, "\u0053\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0048\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0041\u006c\u0069\u0067\u006eme\u006e\u0074")
	}
	if _gabag.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		_dagd = append(_dagd, "S\u0054_\u0056\u0065\u0072\u0074\u0069\u0063\u0061\u006cA\u006c\u0069\u0067\u006eme\u006e\u0074")
	}
	if _gabag.ST_ChildDirection != ST_ChildDirectionUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u0044\u0069\u0072\u0065c\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_ChildAlignment != ST_ChildAlignmentUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u0041\u006c\u0069\u0067n\u006d\u0065\u006e\u0074")
	}
	if _gabag.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0053\u0065\u0063\u006f\u006e\u0064\u0061\u0072y\u0043\u0068\u0069\u006c\u0064\u0041\u006c\u0069\u0067\u006em\u0065\u006e\u0074")
	}
	if _gabag.ST_LinearDirection != ST_LinearDirectionUnset {
		_dagd = append(_dagd, "\u0053T\u005fL\u0069\u006e\u0065\u0061\u0072D\u0069\u0072e\u0063\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		_dagd = append(_dagd, "S\u0054\u005f\u0053\u0065\u0063\u006fn\u0064\u0061\u0072\u0079\u004c\u0069\u006e\u0065\u0061r\u0044\u0069\u0072e\u0063t\u0069\u006f\u006e")
	}
	if _gabag.ST_StartingElement != ST_StartingElementUnset {
		_dagd = append(_dagd, "\u0053T\u005fS\u0074\u0061\u0072\u0074\u0069n\u0067\u0045l\u0065\u006d\u0065\u006e\u0074")
	}
	if _gabag.ST_BendPoint != ST_BendPointUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005fB\u0065\u006e\u0064\u0050\u006f\u0069\u006e\u0074")
	}
	if _gabag.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0043on\u006e\u0065\u0063\u0074\u006f\u0072\u0052\u006f\u0075\u0074\u0069\u006e\u0067")
	}
	if _gabag.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0041\u0072\u0072\u006f\u0077\u0068\u0065\u0061\u0064S\u0074\u0079\u006c\u0065")
	}
	if _gabag.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		_dagd = append(_dagd, "S\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u0044\u0069m\u0065\u006e\u0073\u0069\u006f\u006e")
	}
	if _gabag.ST_RotationPath != ST_RotationPathUnset {
		_dagd = append(_dagd, "\u0053T\u005fR\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0050\u0061\u0074\u0068")
	}
	if _gabag.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		_dagd = append(_dagd, "S\u0054\u005f\u0043\u0065nt\u0065r\u0053\u0068\u0061\u0070\u0065M\u0061\u0070\u0070\u0069\u006e\u0067")
	}
	if _gabag.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u004e\u006f\u0064\u0065\u0048\u006f\u0072\u0069z\u006f\u006e\u0074\u0061\u006c\u0041\u006c\u0069\u0067\u006em\u0065\u006e\u0074")
	}
	if _gabag.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		_dagd = append(_dagd, "\u0053T\u005f\u004e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0069\u0063a\u006c\u0041\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074")
	}
	if _gabag.ST_FallbackDimension != ST_FallbackDimensionUnset {
		_dagd = append(_dagd, "S\u0054_\u0046\u0061\u006c\u006c\u0062\u0061\u0063\u006bD\u0069\u006d\u0065\u006esi\u006f\u006e")
	}
	if _gabag.ST_TextDirection != ST_TextDirectionUnset {
		_dagd = append(_dagd, "\u0053\u0054_\u0054\u0065\u0078t\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		_dagd = append(_dagd, "\u0053T\u005f\u0050\u0079\u0072\u0061\u006d\u0069\u0064\u0041\u0063\u0063e\u006e\u0074\u0050\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0050\u0079\u0072\u0061\u006d\u0069\u0064\u0041c\u0063\u0065\u006e\u0074\u0054\u0065\u0078\u0074\u004d\u0061r\u0067\u0069\u006e")
	}
	if _gabag.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		_dagd = append(_dagd, "S\u0054\u005f\u0054\u0065xt\u0042l\u006f\u0063\u006b\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005fTe\u0078\u0074\u0041\u006e\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c")
	}
	if _gabag.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		_dagd = append(_dagd, "S\u0054\u005f\u0054\u0065xt\u0041n\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0069\u0063\u0061\u006c")
	}
	if _gabag.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005fDi\u0061\u0067\u0072\u0061\u006d\u0054\u0065\u0078\u0074\u0041\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074")
	}
	if _gabag.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0041ut\u006f\u0054\u0065\u0078\u0074\u0052\u006f\u0074\u0061\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_GrowDirection != ST_GrowDirectionUnset {
		_dagd = append(_dagd, "\u0053\u0054_\u0047\u0072\u006fw\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_FlowDirection != ST_FlowDirectionUnset {
		_dagd = append(_dagd, "\u0053\u0054_\u0046\u006c\u006fw\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _gabag.ST_ContinueDirection != ST_ContinueDirectionUnset {
		_dagd = append(_dagd, "S\u0054_\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065D\u0069\u0072\u0065\u0063ti\u006f\u006e")
	}
	if _gabag.ST_Breakpoint != ST_BreakpointUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0070\u006f\u0069\u006e\u0074")
	}
	if _gabag.ST_Offset != ST_OffsetUnset {
		_dagd = append(_dagd, "\u0053T\u005f\u004f\u0066\u0066\u0073\u0065t")
	}
	if _gabag.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		_dagd = append(_dagd, "S\u0054\u005f\u0048\u0069er\u0061r\u0063\u0068\u0079\u0041\u006ci\u0067\u006e\u006d\u0065\u006e\u0074")
	}
	if _gabag.Int32 != nil {
		_dagd = append(_dagd, "\u0049\u006e\u00743\u0032")
	}
	if _gabag.Float64 != nil {
		_dagd = append(_dagd, "\u0046l\u006f\u0061\u0074\u0036\u0034")
	}
	if _gabag.Bool != nil {
		_dagd = append(_dagd, "\u0042\u006f\u006f\u006c")
	}
	if _gabag.StringVal != nil {
		_dagd = append(_dagd, "\u0053t\u0072\u0069\u006e\u0067\u0056\u0061l")
	}
	if _gabag.ST_ConnectorPoint != ST_ConnectorPointUnset {
		_dagd = append(_dagd, "\u0053\u0054\u005f\u0043\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072P\u006f\u0069\u006e\u0074")
	}
	if len(_dagd) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _dagd)
	}
	return nil
}
func (_bddf *ST_BendPoint) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bddf = 0
	case "\u0062\u0065\u0067":
		*_bddf = 1
	case "\u0064\u0065\u0066":
		*_bddf = 2
	case "\u0065\u006e\u0064":
		*_bddf = 3
	}
	return nil
}
func (_gdgea ST_TextBlockDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dbdg := _d.Attr{}
	_dbdg.Name = name
	switch _gdgea {
	case ST_TextBlockDirectionUnset:
		_dbdg.Value = ""
	case ST_TextBlockDirectionHorz:
		_dbdg.Value = "\u0068\u006f\u0072\u007a"
	case ST_TextBlockDirectionVert:
		_dbdg.Value = "\u0076\u0065\u0072\u0074"
	}
	return _dbdg, nil
}
func (_gdgfd *ST_StartingElement) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gdgfd = 0
	case "\u006e\u006f\u0064\u0065":
		*_gdgfd = 1
	case "\u0074\u0072\u0061n\u0073":
		*_gdgfd = 2
	}
	return nil
}

// ValidateWithPath validates the CT_CTCategories and its children, prefixing error messages with path
func (_ddf *CT_CTCategories) ValidateWithPath(path string) error {
	for _dbc, _bba := range _ddf.Cat {
		if _aegf := _bba.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0061\u0074\u005b\u0025\u0064\u005d", path, _dbc)); _aegf != nil {
			return _aegf
		}
	}
	return nil
}

// ValidateWithPath validates the CT_SampleData and its children, prefixing error messages with path
func (_beca *CT_SampleData) ValidateWithPath(path string) error {
	if _beca.DataModel != nil {
		if _fbfag := _beca.DataModel.ValidateWithPath(path + "\u002f\u0044\u0061\u0074\u0061\u004d\u006f\u0064\u0065\u006c"); _fbfag != nil {
			return _fbfag
		}
	}
	return nil
}
func (_edfg *ST_ModelId) ValidateWithPath(path string) error {
	_gggee := []string{}
	if _edfg.Int32 != nil {
		_gggee = append(_gggee, "\u0049\u006e\u00743\u0032")
	}
	if _edfg.ST_Guid != nil {
		_gggee = append(_gggee, "\u0053T\u005f\u0047\u0075\u0069\u0064")
	}
	if len(_gggee) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _gggee)
	}
	return nil
}

const (
	ST_AxisTypeUnset       ST_AxisType = 0
	ST_AxisTypeSelf        ST_AxisType = 1
	ST_AxisTypeCh          ST_AxisType = 2
	ST_AxisTypeDes         ST_AxisType = 3
	ST_AxisTypeDesOrSelf   ST_AxisType = 4
	ST_AxisTypePar         ST_AxisType = 5
	ST_AxisTypeAncst       ST_AxisType = 6
	ST_AxisTypeAncstOrSelf ST_AxisType = 7
	ST_AxisTypeFollowSib   ST_AxisType = 8
	ST_AxisTypePrecedSib   ST_AxisType = 9
	ST_AxisTypeFollow      ST_AxisType = 10
	ST_AxisTypePreced      ST_AxisType = 11
	ST_AxisTypeRoot        ST_AxisType = 12
	ST_AxisTypeNone        ST_AxisType = 13
)

func (_aafcc ST_TextAnchorVertical) ValidateWithPath(path string) error {
	switch _aafcc {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aafcc))
	}
	return nil
}
func (_gcebgg ST_ElementType) Validate() error { return _gcebgg.ValidateWithPath("") }
func (_fabbc *ST_ConstraintRelationship) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_abeg, _fgaag := d.Token()
	if _fgaag != nil {
		return _fgaag
	}
	if _bdfb, _fbcfd := _abeg.(_d.EndElement); _fbcfd && _bdfb.Name == start.Name {
		*_fabbc = 1
		return nil
	}
	if _gfca, _dgga := _abeg.(_d.CharData); !_dgga {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _abeg)
	} else {
		switch string(_gfca) {
		case "":
			*_fabbc = 0
		case "\u0073\u0065\u006c\u0066":
			*_fabbc = 1
		case "\u0063\u0068":
			*_fabbc = 2
		case "\u0064\u0065\u0073":
			*_fabbc = 3
		}
	}
	_abeg, _fgaag = d.Token()
	if _fgaag != nil {
		return _fgaag
	}
	if _ebgb, _dffdfe := _abeg.(_d.EndElement); _dffdfe && _ebgb.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _abeg)
}
func (_dege ST_ConstraintRelationship) String() string {
	switch _dege {
	case 0:
		return ""
	case 1:
		return "\u0073\u0065\u006c\u0066"
	case 2:
		return "\u0063\u0068"
	case 3:
		return "\u0064\u0065\u0073"
	}
	return ""
}
func (_fbdad ST_HierBranchStyle) Validate() error { return _fbdad.ValidateWithPath("") }
func (_ddde *CT_ForEach) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _fgeg := range start.Attr {
		if _fgeg.Name.Local == "\u0072\u0065\u0066" {
			_fgbae, _fagf := _fgeg.Value, error(nil)
			if _fagf != nil {
				return _fagf
			}
			_ddde.RefAttr = &_fgbae
			continue
		}
		if _fgeg.Name.Local == "\u006e\u0061\u006d\u0065" {
			_feafd, _gfcdc := _fgeg.Value, error(nil)
			if _gfcdc != nil {
				return _gfcdc
			}
			_ddde.NameAttr = &_feafd
			continue
		}
		if _fgeg.Name.Local == "\u0061\u0078\u0069\u0073" {
			_eaffa, _dcff := ParseSliceST_AxisTypes(_fgeg.Value)
			if _dcff != nil {
				return _dcff
			}
			_ddde.AxisAttr = &_eaffa
			continue
		}
		if _fgeg.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_aaee, _bcca := ParseSliceST_ElementTypes(_fgeg.Value)
			if _bcca != nil {
				return _bcca
			}
			_ddde.PtTypeAttr = &_aaee
			continue
		}
		if _fgeg.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_fabe, _dbcg := ParseSliceST_Booleans(_fgeg.Value)
			if _dbcg != nil {
				return _dbcg
			}
			_ddde.HideLastTransAttr = &_fabe
			continue
		}
		if _fgeg.Name.Local == "\u0073\u0074" {
			_gggc, _defb := ParseSliceST_Ints(_fgeg.Value)
			if _defb != nil {
				return _defb
			}
			_ddde.StAttr = &_gggc
			continue
		}
		if _fgeg.Name.Local == "\u0063\u006e\u0074" {
			_bccd, _cfga := ParseSliceST_UnsignedInts(_fgeg.Value)
			if _cfga != nil {
				return _cfga
			}
			_ddde.CntAttr = &_bccd
			continue
		}
		if _fgeg.Name.Local == "\u0073\u0074\u0065\u0070" {
			_fddc, _bgdf := ParseSliceST_Ints(_fgeg.Value)
			if _bgdf != nil {
				return _bgdf
			}
			_ddde.StepAttr = &_fddc
			continue
		}
	}
_ebad:
	for {
		_afcc, _eaee := d.Token()
		if _eaee != nil {
			return _eaee
		}
		switch _fcbe := _afcc.(type) {
		case _d.StartElement:
			switch _fcbe.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_cded := NewCT_Algorithm()
				if _bbe := d.DecodeElement(_cded, &_fcbe); _bbe != nil {
					return _bbe
				}
				_ddde.Alg = append(_ddde.Alg, _cded)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_bagb := NewCT_Shape()
				if _cgb := d.DecodeElement(_bagb, &_fcbe); _cgb != nil {
					return _cgb
				}
				_ddde.Shape = append(_ddde.Shape, _bagb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_fefc := NewCT_PresentationOf()
				if _afab := d.DecodeElement(_fefc, &_fcbe); _afab != nil {
					return _afab
				}
				_ddde.PresOf = append(_ddde.PresOf, _fefc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_dged := NewCT_Constraints()
				if _facc := d.DecodeElement(_dged, &_fcbe); _facc != nil {
					return _facc
				}
				_ddde.ConstrLst = append(_ddde.ConstrLst, _dged)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_dcace := NewCT_Rules()
				if _bdeg := d.DecodeElement(_dcace, &_fcbe); _bdeg != nil {
					return _bdeg
				}
				_ddde.RuleLst = append(_ddde.RuleLst, _dcace)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_cgba := NewCT_ForEach()
				if _fggdd := d.DecodeElement(_cgba, &_fcbe); _fggdd != nil {
					return _fggdd
				}
				_ddde.ForEach = append(_ddde.ForEach, _cgba)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_affb := NewCT_LayoutNode()
				if _eggb := d.DecodeElement(_affb, &_fcbe); _eggb != nil {
					return _eggb
				}
				_ddde.LayoutNode = append(_ddde.LayoutNode, _affb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_fcgf := NewCT_Choose()
				if _dfed := d.DecodeElement(_fcgf, &_fcbe); _dfed != nil {
					return _dfed
				}
				_ddde.Choose = append(_ddde.Choose, _fcgf)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_fcad := _ca.NewCT_OfficeArtExtensionList()
				if _dcdb := d.DecodeElement(_fcad, &_fcbe); _dcdb != nil {
					return _dcdb
				}
				_ddde.ExtLst = append(_ddde.ExtLst, _fcad)
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fF\u006f\u0072\u0045\u0061\u0063\u0068\u0020\u0025\u0076", _fcbe.Name)
				if _ccbg := d.Skip(); _ccbg != nil {
					return _ccbg
				}
			}
		case _d.EndElement:
			break _ebad
		case _d.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CTCategory and its children
func (_bge *CT_CTCategory) Validate() error {
	return _bge.ValidateWithPath("\u0043\u0054\u005f\u0043\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079")
}
func (_aggbf ST_SecondaryLinearDirection) ValidateWithPath(path string) error {
	switch _aggbf {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aggbf))
	}
	return nil
}

const (
	ST_CenterShapeMappingUnset ST_CenterShapeMapping = 0
	ST_CenterShapeMappingNone  ST_CenterShapeMapping = 1
	ST_CenterShapeMappingFNode ST_CenterShapeMapping = 2
)

func (_addb *ST_ClrAppMethod) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_addb = 0
	case "\u0073\u0070\u0061\u006e":
		*_addb = 1
	case "\u0063\u0079\u0063l\u0065":
		*_addb = 2
	case "\u0072\u0065\u0070\u0065\u0061\u0074":
		*_addb = 3
	}
	return nil
}
func (_cedc *ST_ModelId) Validate() error { return _cedc.ValidateWithPath("") }
func NewCT_SDName() *CT_SDName            { _gaef := &CT_SDName{}; return _gaef }

type LayoutDef struct{ CT_DiagramDefinition }

func (_fcdb *CT_StyleDefinitionHeaderLst) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_eecb:
	for {
		_bgea, _afada := d.Token()
		if _afada != nil {
			return _afada
		}
		switch _ecabb := _bgea.(type) {
		case _d.StartElement:
			switch _ecabb.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_eecff := NewCT_StyleDefinitionHeader()
				if _gaee := d.DecodeElement(_eecff, &_ecabb); _gaee != nil {
					return _gaee
				}
				_fcdb.StyleDefHdr = append(_fcdb.StyleDefHdr, _eecff)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _ecabb.Name)
				if _fecbd := d.Skip(); _fecbd != nil {
					return _fecbd
				}
			}
		case _d.EndElement:
			break _eecb
		case _d.CharData:
		}
	}
	return nil
}
func (_dggb *CT_LayoutNode) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _dggb.NameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_dggb.NameAttr)})
	}
	if _dggb.StyleLblAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}, Value: _c.Sprintf("\u0025\u0076", *_dggb.StyleLblAttr)})
	}
	if _dggb.ChOrderAttr != ST_ChildOrderTypeUnset {
		_bcaab, _gbec := _dggb.ChOrderAttr.MarshalXMLAttr(_d.Name{Local: "\u0063h\u004f\u0072\u0064\u0065\u0072"})
		if _gbec != nil {
			return _gbec
		}
		start.Attr = append(start.Attr, _bcaab)
	}
	if _dggb.MoveWithAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u006f\u0076\u0065\u0057\u0069\u0074\u0068"}, Value: _c.Sprintf("\u0025\u0076", *_dggb.MoveWithAttr)})
	}
	e.EncodeToken(start)
	if _dggb.Alg != nil {
		_fda := _d.StartElement{Name: _d.Name{Local: "\u0061\u006c\u0067"}}
		for _, _ebed := range _dggb.Alg {
			e.EncodeElement(_ebed, _fda)
		}
	}
	if _dggb.Shape != nil {
		_dfdc := _d.StartElement{Name: _d.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _daae := range _dggb.Shape {
			e.EncodeElement(_daae, _dfdc)
		}
	}
	if _dggb.PresOf != nil {
		_dfcb := _d.StartElement{Name: _d.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _dacd := range _dggb.PresOf {
			e.EncodeElement(_dacd, _dfcb)
		}
	}
	if _dggb.ConstrLst != nil {
		_gecg := _d.StartElement{Name: _d.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _fbfa := range _dggb.ConstrLst {
			e.EncodeElement(_fbfa, _gecg)
		}
	}
	if _dggb.RuleLst != nil {
		_bcgc := _d.StartElement{Name: _d.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _aadg := range _dggb.RuleLst {
			e.EncodeElement(_aadg, _bcgc)
		}
	}
	if _dggb.VarLst != nil {
		_ffffa := _d.StartElement{Name: _d.Name{Local: "\u0076\u0061\u0072\u004c\u0073\u0074"}}
		for _, _fbff := range _dggb.VarLst {
			e.EncodeElement(_fbff, _ffffa)
		}
	}
	if _dggb.ForEach != nil {
		_cdgg := _d.StartElement{Name: _d.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _eeee := range _dggb.ForEach {
			e.EncodeElement(_eeee, _cdgg)
		}
	}
	if _dggb.LayoutNode != nil {
		_eceg := _d.StartElement{Name: _d.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _cdec := range _dggb.LayoutNode {
			e.EncodeElement(_cdec, _eceg)
		}
	}
	if _dggb.Choose != nil {
		_fcce := _d.StartElement{Name: _d.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _gfda := range _dggb.Choose {
			e.EncodeElement(_gfda, _fcce)
		}
	}
	if _dggb.ExtLst != nil {
		_gdeg := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _aagf := range _dggb.ExtLst {
			e.EncodeElement(_aagf, _gdeg)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

type CT_CTName struct {
	LangAttr *string
	ValAttr  string
}

// ValidateWithPath validates the StyleDefHdrLst and its children, prefixing error messages with path
func (_gcaeb *StyleDefHdrLst) ValidateWithPath(path string) error {
	if _edbd := _gcaeb.CT_StyleDefinitionHeaderLst.ValidateWithPath(path); _edbd != nil {
		return _edbd
	}
	return nil
}
func (_dfacg ST_ConnectorPoint) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_aedfd := _d.Attr{}
	_aedfd.Name = name
	switch _dfacg {
	case ST_ConnectorPointUnset:
		_aedfd.Value = ""
	case ST_ConnectorPointAuto:
		_aedfd.Value = "\u0061\u0075\u0074\u006f"
	case ST_ConnectorPointBCtr:
		_aedfd.Value = "\u0062\u0043\u0074\u0072"
	case ST_ConnectorPointCtr:
		_aedfd.Value = "\u0063\u0074\u0072"
	case ST_ConnectorPointMidL:
		_aedfd.Value = "\u006d\u0069\u0064\u004c"
	case ST_ConnectorPointMidR:
		_aedfd.Value = "\u006d\u0069\u0064\u0052"
	case ST_ConnectorPointTCtr:
		_aedfd.Value = "\u0074\u0043\u0074\u0072"
	case ST_ConnectorPointBL:
		_aedfd.Value = "\u0062\u004c"
	case ST_ConnectorPointBR:
		_aedfd.Value = "\u0062\u0052"
	case ST_ConnectorPointTL:
		_aedfd.Value = "\u0074\u004c"
	case ST_ConnectorPointTR:
		_aedfd.Value = "\u0074\u0052"
	case ST_ConnectorPointRadial:
		_aedfd.Value = "\u0072\u0061\u0064\u0069\u0061\u006c"
	}
	return _aedfd, nil
}

type CT_ElemPropSet struct {
	PresAssocIDAttr          *ST_ModelId
	PresNameAttr             *string
	PresStyleLblAttr         *string
	PresStyleIdxAttr         *int32
	PresStyleCntAttr         *int32
	LoTypeIdAttr             *string
	LoCatIdAttr              *string
	QsTypeIdAttr             *string
	QsCatIdAttr              *string
	CsTypeIdAttr             *string
	CsCatIdAttr              *string
	Coherent3DOffAttr        *bool
	PhldrTAttr               *string
	PhldrAttr                *bool
	CustAngAttr              *int32
	CustFlipVertAttr         *bool
	CustFlipHorAttr          *bool
	CustSzXAttr              *int32
	CustSzYAttr              *int32
	CustScaleXAttr           *ST_PrSetCustVal
	CustScaleYAttr           *ST_PrSetCustVal
	CustTAttr                *bool
	CustLinFactXAttr         *ST_PrSetCustVal
	CustLinFactYAttr         *ST_PrSetCustVal
	CustLinFactNeighborXAttr *ST_PrSetCustVal
	CustLinFactNeighborYAttr *ST_PrSetCustVal
	CustRadScaleRadAttr      *ST_PrSetCustVal
	CustRadScaleIncAttr      *ST_PrSetCustVal
	PresLayoutVars           *CT_LayoutVariablePropertySet
	Style                    *_ca.CT_ShapeStyle
}

func (_gcfec ST_Direction) ValidateWithPath(path string) error {
	switch _gcfec {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gcfec))
	}
	return nil
}
func (_gacag ST_ClrAppMethod) String() string {
	switch _gacag {
	case 0:
		return ""
	case 1:
		return "\u0073\u0070\u0061\u006e"
	case 2:
		return "\u0063\u0079\u0063l\u0065"
	case 3:
		return "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return ""
}
func (_fdggf ST_FallbackDimension) String() string {
	switch _fdggf {
	case 0:
		return ""
	case 1:
		return "\u0031\u0044"
	case 2:
		return "\u0032\u0044"
	}
	return ""
}
func (_badc ST_CenterShapeMapping) String() string {
	switch _badc {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0066\u004e\u006fd\u0065"
	}
	return ""
}

// ValidateWithPath validates the CT_Colors and its children, prefixing error messages with path
func (_efc *CT_Colors) ValidateWithPath(path string) error {
	if _dfbc := _efc.MethAttr.ValidateWithPath(path + "\u002fM\u0065\u0074\u0068\u0041\u0074\u0074r"); _dfbc != nil {
		return _dfbc
	}
	if _cga := _efc.HueDirAttr.ValidateWithPath(path + "/\u0048\u0075\u0065\u0044\u0069\u0072\u0041\u0074\u0074\u0072"); _cga != nil {
		return _cga
	}
	for _fgba, _caca := range _efc.EG_ColorChoice {
		if _afdc := _caca.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0045\u0047_C\u006fl\u006f\u0072\u0043\u0068\u006fi\u0063\u0065\u005b\u0025\u0064\u005d", path, _fgba)); _afdc != nil {
			return _afdc
		}
	}
	return nil
}
func (_ebede *ST_ChildOrderType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ebede = 0
	case "\u0062":
		*_ebede = 1
	case "\u0074":
		*_ebede = 2
	}
	return nil
}
func NewCT_Constraint() *CT_Constraint { _ffbe := &CT_Constraint{}; return _ffbe }
func (_ce *CT_AnimOne) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _dgc := range start.Attr {
		if _dgc.Name.Local == "\u0076\u0061\u006c" {
			_ce.ValAttr.UnmarshalXMLAttr(_dgc)
			continue
		}
	}
	for {
		_bfgc, _ebe := d.Token()
		if _ebe != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065\u003a\u0020%\u0073", _ebe)
		}
		if _cbg, _eef := _bfgc.(_d.EndElement); _eef && _cbg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_dacg ST_Offset) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_adacc := _d.Attr{}
	_adacc.Name = name
	switch _dacg {
	case ST_OffsetUnset:
		_adacc.Value = ""
	case ST_OffsetCtr:
		_adacc.Value = "\u0063\u0074\u0072"
	case ST_OffsetOff:
		_adacc.Value = "\u006f\u0066\u0066"
	}
	return _adacc, nil
}
func (_afaa *ST_StartingElement) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cacdd, _daab := d.Token()
	if _daab != nil {
		return _daab
	}
	if _gacg, _afaac := _cacdd.(_d.EndElement); _afaac && _gacg.Name == start.Name {
		*_afaa = 1
		return nil
	}
	if _ffab, _gcaade := _cacdd.(_d.CharData); !_gcaade {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cacdd)
	} else {
		switch string(_ffab) {
		case "":
			*_afaa = 0
		case "\u006e\u006f\u0064\u0065":
			*_afaa = 1
		case "\u0074\u0072\u0061n\u0073":
			*_afaa = 2
		}
	}
	_cacdd, _daab = d.Token()
	if _daab != nil {
		return _daab
	}
	if _aaec, _bbbg := _cacdd.(_d.EndElement); _bbbg && _aaec.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cacdd)
}
func (_aaddc ST_ParameterVal) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _aaddc.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_DiagramHorizontalAlignment.String()))
	}
	if _aaddc.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_VerticalAlignment.String()))
	}
	if _aaddc.ST_ChildDirection != ST_ChildDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_ChildDirection.String()))
	}
	if _aaddc.ST_ChildAlignment != ST_ChildAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_ChildAlignment.String()))
	}
	if _aaddc.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_SecondaryChildAlignment.String()))
	}
	if _aaddc.ST_LinearDirection != ST_LinearDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_LinearDirection.String()))
	}
	if _aaddc.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_SecondaryLinearDirection.String()))
	}
	if _aaddc.ST_StartingElement != ST_StartingElementUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_StartingElement.String()))
	}
	if _aaddc.ST_BendPoint != ST_BendPointUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_BendPoint.String()))
	}
	if _aaddc.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_ConnectorRouting.String()))
	}
	if _aaddc.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_ArrowheadStyle.String()))
	}
	if _aaddc.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_ConnectorDimension.String()))
	}
	if _aaddc.ST_RotationPath != ST_RotationPathUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_RotationPath.String()))
	}
	if _aaddc.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_CenterShapeMapping.String()))
	}
	if _aaddc.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_NodeHorizontalAlignment.String()))
	}
	if _aaddc.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_NodeVerticalAlignment.String()))
	}
	if _aaddc.ST_FallbackDimension != ST_FallbackDimensionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_FallbackDimension.String()))
	}
	if _aaddc.ST_TextDirection != ST_TextDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_TextDirection.String()))
	}
	if _aaddc.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_PyramidAccentPosition.String()))
	}
	if _aaddc.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_PyramidAccentTextMargin.String()))
	}
	if _aaddc.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_TextBlockDirection.String()))
	}
	if _aaddc.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_TextAnchorHorizontal.String()))
	}
	if _aaddc.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_TextAnchorVertical.String()))
	}
	if _aaddc.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_DiagramTextAlignment.String()))
	}
	if _aaddc.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_AutoTextRotation.String()))
	}
	if _aaddc.ST_GrowDirection != ST_GrowDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_GrowDirection.String()))
	}
	if _aaddc.ST_FlowDirection != ST_FlowDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_FlowDirection.String()))
	}
	if _aaddc.ST_ContinueDirection != ST_ContinueDirectionUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_ContinueDirection.String()))
	}
	if _aaddc.ST_Breakpoint != ST_BreakpointUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_Breakpoint.String()))
	}
	if _aaddc.ST_Offset != ST_OffsetUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_Offset.String()))
	}
	if _aaddc.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_HierarchyAlignment.String()))
	}
	if _aaddc.Int32 != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", *_aaddc.Int32)))
	}
	if _aaddc.Float64 != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0066", *_aaddc.Float64)))
	}
	if _aaddc.Bool != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", _efffd(*_aaddc.Bool))))
	}
	if _aaddc.StringVal != nil {
		e.EncodeToken(_d.CharData(*_aaddc.StringVal))
	}
	if _aaddc.ST_ConnectorPoint != ST_ConnectorPointUnset {
		e.EncodeToken(_d.CharData(_aaddc.ST_ConnectorPoint.String()))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}
func (_gbgee ST_ArrowheadStyle) Validate() error { return _gbgee.ValidateWithPath("") }
func (_bfdf *ST_ParameterId) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dffbb, _gaccg := d.Token()
	if _gaccg != nil {
		return _gaccg
	}
	if _bfgba, _bgebe := _dffbb.(_d.EndElement); _bgebe && _bfgba.Name == start.Name {
		*_bfdf = 1
		return nil
	}
	if _abed, _gabg := _dffbb.(_d.CharData); !_gabg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dffbb)
	} else {
		switch string(_abed) {
		case "":
			*_bfdf = 0
		case "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n":
			*_bfdf = 1
		case "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n":
			*_bfdf = 2
		case "\u0063\u0068\u0044i\u0072":
			*_bfdf = 3
		case "\u0063h\u0041\u006c\u0069\u0067\u006e":
			*_bfdf = 4
		case "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e":
			*_bfdf = 5
		case "\u006c\u0069\u006e\u0044\u0069\u0072":
			*_bfdf = 6
		case "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r":
			*_bfdf = 7
		case "\u0073\u0074\u0045\u006c\u0065\u006d":
			*_bfdf = 8
		case "\u0062\u0065\u006e\u0064\u0050\u0074":
			*_bfdf = 9
		case "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074":
			*_bfdf = 10
		case "\u0062\u0065\u0067\u0053\u0074\u0079":
			*_bfdf = 11
		case "\u0065\u006e\u0064\u0053\u0074\u0079":
			*_bfdf = 12
		case "\u0064\u0069\u006d":
			*_bfdf = 13
		case "\u0072o\u0074\u0050\u0061\u0074\u0068":
			*_bfdf = 14
		case "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p":
			*_bfdf = 15
		case "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e":
			*_bfdf = 16
		case "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e":
			*_bfdf = 17
		case "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b":
			*_bfdf = 18
		case "\u0074\u0078\u0044i\u0072":
			*_bfdf = 19
		case "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073":
			*_bfdf = 20
		case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072":
			*_bfdf = 21
		case "\u0074x\u0042\u006c\u0044\u0069\u0072":
			*_bfdf = 22
		case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a":
			*_bfdf = 23
		case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074":
			*_bfdf = 24
		case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068":
			*_bfdf = 25
		case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068":
			*_bfdf = 26
		case "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e":
			*_bfdf = 27
		case "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e":
			*_bfdf = 28
		case "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
			*_bfdf = 29
		case "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
			*_bfdf = 30
		case "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft":
			*_bfdf = 31
		case "\u0067\u0072\u0044i\u0072":
			*_bfdf = 32
		case "\u0066l\u006f\u0077\u0044\u0069\u0072":
			*_bfdf = 33
		case "\u0063o\u006e\u0074\u0044\u0069\u0072":
			*_bfdf = 34
		case "\u0062\u006b\u0070\u0074":
			*_bfdf = 35
		case "\u006f\u0066\u0066":
			*_bfdf = 36
		case "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n":
			*_bfdf = 37
		case "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c":
			*_bfdf = 38
		case "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c":
			*_bfdf = 39
		case "\u0073\u0074\u0041n\u0067":
			*_bfdf = 40
		case "\u0073p\u0061\u006e\u0041\u006e\u0067":
			*_bfdf = 41
		case "\u0061\u0072":
			*_bfdf = 42
		case "\u006cn\u0053\u0070\u0050\u0061\u0072":
			*_bfdf = 43
		case "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050":
			*_bfdf = 44
		case "\u006c\u006e\u0053\u0070\u0043\u0068":
			*_bfdf = 45
		case "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P":
			*_bfdf = 46
		case "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074":
			*_bfdf = 47
		case "\u0061l\u0069\u0067\u006e\u0054\u0078":
			*_bfdf = 48
		case "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065":
			*_bfdf = 49
		case "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065":
			*_bfdf = 50
		case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065":
			*_bfdf = 51
		case "\u0073r\u0063\u004e\u006f\u0064\u0065":
			*_bfdf = 52
		case "\u0064s\u0074\u004e\u006f\u0064\u0065":
			*_bfdf = 53
		case "\u0062\u0065\u0067\u0050\u0074\u0073":
			*_bfdf = 54
		case "\u0065\u006e\u0064\u0050\u0074\u0073":
			*_bfdf = 55
		}
	}
	_dffbb, _gaccg = d.Token()
	if _gaccg != nil {
		return _gaccg
	}
	if _bgfga, _efcbg := _dffbb.(_d.EndElement); _efcbg && _bgfga.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dffbb)
}

type CT_ColorTransformHeaderLst struct{ ColorsDefHdr []*CT_ColorTransformHeader }

func (_fad *AG_ConstraintRefAttributes) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _ab := range start.Attr {
		if _ab.Name.Local == "\u0072e\u0066\u0054\u0079\u0070\u0065" {
			_fad.RefTypeAttr.UnmarshalXMLAttr(_ab)
			continue
		}
		if _ab.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072" {
			_fad.RefForAttr.UnmarshalXMLAttr(_ab)
			continue
		}
		if _ab.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065" {
			_dc, _bf := _ab.Value, error(nil)
			if _bf != nil {
				return _bf
			}
			_fad.RefForNameAttr = &_dc
			continue
		}
		if _ab.Name.Local == "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e" {
			_fad.RefPtTypeAttr.UnmarshalXMLAttr(_ab)
			continue
		}
	}
	for {
		_ee, _eab := d.Token()
		if _eab != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0041\u0047\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074\u0052\u0065\u0066A\u0074\u0074\u0072\u0069\u0062u\u0074\u0065s\u003a\u0020\u0025\u0073", _eab)
		}
		if _aa, _gf := _ee.(_d.EndElement); _gf && _aa.Name == start.Name {
			break
		}
	}
	return nil
}
func (_efgea ST_SecondaryLinearDirection) Validate() error { return _efgea.ValidateWithPath("") }
func (_geea *CT_Description) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _geea.LangAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_geea.LangAttr)})
	}
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _geea.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_bddef ST_RotationPath) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ebfbb := _d.Attr{}
	_ebfbb.Name = name
	switch _bddef {
	case ST_RotationPathUnset:
		_ebfbb.Value = ""
	case ST_RotationPathNone:
		_ebfbb.Value = "\u006e\u006f\u006e\u0065"
	case ST_RotationPathAlongPath:
		_ebfbb.Value = "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h"
	}
	return _ebfbb, nil
}
func (_dbdb ST_ConstraintRelationship) Validate() error { return _dbdb.ValidateWithPath("") }
func (_adge *CT_ChildPref) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _dgb := range start.Attr {
		if _dgb.Name.Local == "\u0076\u0061\u006c" {
			_aea, _dcgg := _ae.ParseInt(_dgb.Value, 10, 32)
			if _dcgg != nil {
				return _dcgg
			}
			_ebec := int32(_aea)
			_adge.ValAttr = &_ebec
			continue
		}
	}
	for {
		_cbaf, _aegff := d.Token()
		if _aegff != nil {
			return _c.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0068i\u006c\u0064\u0050\u0072\u0065\u0066\u003a\u0020\u0025\u0073", _aegff)
		}
		if _fdda, _ebcf := _cbaf.(_d.EndElement); _ebcf && _fdda.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DiagramDefinitionHeader and its children
func (_geda *CT_DiagramDefinitionHeader) Validate() error {
	return _geda.ValidateWithPath("\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065\u0072")
}
func (_bdbe *ST_ParameterVal) Validate() error { return _bdbe.ValidateWithPath("") }
func (_degb ST_FlowDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_egegfe := _d.Attr{}
	_egegfe.Name = name
	switch _degb {
	case ST_FlowDirectionUnset:
		_egegfe.Value = ""
	case ST_FlowDirectionRow:
		_egegfe.Value = "\u0072\u006f\u0077"
	case ST_FlowDirectionCol:
		_egegfe.Value = "\u0063\u006f\u006c"
	}
	return _egegfe, nil
}
func (_bfbbdf *ST_FallbackDimension) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ecddg, _geagfb := d.Token()
	if _geagfb != nil {
		return _geagfb
	}
	if _eabbd, _cgfg := _ecddg.(_d.EndElement); _cgfg && _eabbd.Name == start.Name {
		*_bfbbdf = 1
		return nil
	}
	if _ggadg, _gbdgg := _ecddg.(_d.CharData); !_gbdgg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ecddg)
	} else {
		switch string(_ggadg) {
		case "":
			*_bfbbdf = 0
		case "\u0031\u0044":
			*_bfbbdf = 1
		case "\u0032\u0044":
			*_bfbbdf = 2
		}
	}
	_ecddg, _geagfb = d.Token()
	if _geagfb != nil {
		return _geagfb
	}
	if _ccab, _gbgg := _ecddg.(_d.EndElement); _gbgg && _ccab.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ecddg)
}

// ST_FunctionValue is a union type
type ST_FunctionValue struct {
	Int32               *int32
	Bool                *bool
	ST_Direction        ST_Direction
	ST_HierBranchStyle  ST_HierBranchStyle
	ST_AnimOneStr       ST_AnimOneStr
	ST_AnimLvlStr       ST_AnimLvlStr
	ST_ResizeHandlesStr ST_ResizeHandlesStr
}

func (_ggagb *ST_CenterShapeMapping) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_beef, _becea := d.Token()
	if _becea != nil {
		return _becea
	}
	if _daefa, _dgce := _beef.(_d.EndElement); _dgce && _daefa.Name == start.Name {
		*_ggagb = 1
		return nil
	}
	if _agbfc, _fafd := _beef.(_d.CharData); !_fafd {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _beef)
	} else {
		switch string(_agbfc) {
		case "":
			*_ggagb = 0
		case "\u006e\u006f\u006e\u0065":
			*_ggagb = 1
		case "\u0066\u004e\u006fd\u0065":
			*_ggagb = 2
		}
	}
	_beef, _becea = d.Token()
	if _becea != nil {
		return _becea
	}
	if _eaea, _adfc := _beef.(_d.EndElement); _adfc && _eaea.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _beef)
}

// ValidateWithPath validates the CT_HierBranchStyle and its children, prefixing error messages with path
func (_effa *CT_HierBranchStyle) ValidateWithPath(path string) error {
	if _dddd := _effa.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _dddd != nil {
		return _dddd
	}
	return nil
}
func NewCT_CTName() *CT_CTName { _gcf := &CT_CTName{}; return _gcf }
func (_fbbd *CT_Cxn) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006do\u0064\u0065\u006c\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", _fbbd.ModelIdAttr)})
	if _fbbd.TypeAttr != ST_CxnTypeUnset {
		_dddb, _bbf := _fbbd.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _bbf != nil {
			return _bbf
		}
		start.Attr = append(start.Attr, _dddb)
	}
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0072\u0063I\u0064"}, Value: _c.Sprintf("\u0025\u0076", _fbbd.SrcIdAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0064\u0065\u0073\u0074\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", _fbbd.DestIdAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0072\u0063\u004f\u0072\u0064"}, Value: _c.Sprintf("\u0025\u0076", _fbbd.SrcOrdAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0064e\u0073\u0074\u004f\u0072\u0064"}, Value: _c.Sprintf("\u0025\u0076", _fbbd.DestOrdAttr)})
	if _fbbd.ParTransIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fbbd.ParTransIdAttr)})
	}
	if _fbbd.SibTransIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fbbd.SibTransIdAttr)})
	}
	if _fbbd.PresIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0065\u0073\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fbbd.PresIdAttr)})
	}
	e.EncodeToken(start)
	if _fbbd.ExtLst != nil {
		_ecbg := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_fbbd.ExtLst, _ecbg)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_LayoutNode and its children, prefixing error messages with path
func (_bgad *CT_LayoutNode) ValidateWithPath(path string) error {
	if _cgdbf := _bgad.ChOrderAttr.ValidateWithPath(path + "\u002f\u0043\u0068O\u0072\u0064\u0065\u0072\u0041\u0074\u0074\u0072"); _cgdbf != nil {
		return _cgdbf
	}
	for _edcb, _bfadd := range _bgad.Alg {
		if _gfbd := _bfadd.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _edcb)); _gfbd != nil {
			return _gfbd
		}
	}
	for _ecbe, _bebd := range _bgad.Shape {
		if _eecgea := _bebd.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _ecbe)); _eecgea != nil {
			return _eecgea
		}
	}
	for _feaa, _aggff := range _bgad.PresOf {
		if _agfag := _aggff.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _feaa)); _agfag != nil {
			return _agfag
		}
	}
	for _edgda, _fecca := range _bgad.ConstrLst {
		if _becf := _fecca.ValidateWithPath(_c.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _edgda)); _becf != nil {
			return _becf
		}
	}
	for _ecbf, _cbeb := range _bgad.RuleLst {
		if _bfda := _cbeb.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _ecbf)); _bfda != nil {
			return _bfda
		}
	}
	for _fgfg, _bebe := range _bgad.VarLst {
		if _cbd := _bebe.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0056\u0061\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _fgfg)); _cbd != nil {
			return _cbd
		}
	}
	for _cfdc, _bfec := range _bgad.ForEach {
		if _gfdd := _bfec.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _cfdc)); _gfdd != nil {
			return _gfdd
		}
	}
	for _cge, _eeggc := range _bgad.LayoutNode {
		if _afgg := _eeggc.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _cge)); _afgg != nil {
			return _afgg
		}
	}
	for _feag, _abdc := range _bgad.Choose {
		if _affbe := _abdc.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _feag)); _affbe != nil {
			return _affbe
		}
	}
	for _ebgdb, _deef := range _bgad.ExtLst {
		if _dfgfa := _deef.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _ebgdb)); _dfgfa != nil {
			return _dfgfa
		}
	}
	return nil
}
func (_gacc ST_PrSetCustVal) String() string {
	if _gacc.ST_Percentage != nil {
		return _c.Sprintf("\u0025\u0076", *_gacc.ST_Percentage)
	}
	if _gacc.Int32 != nil {
		return _c.Sprintf("\u0025\u0076", *_gacc.Int32)
	}
	return ""
}

const (
	ST_ConnectorRoutingUnset     ST_ConnectorRouting = 0
	ST_ConnectorRoutingStra      ST_ConnectorRouting = 1
	ST_ConnectorRoutingBend      ST_ConnectorRouting = 2
	ST_ConnectorRoutingCurve     ST_ConnectorRouting = 3
	ST_ConnectorRoutingLongCurve ST_ConnectorRouting = 4
)
const (
	ST_NodeVerticalAlignmentUnset ST_NodeVerticalAlignment = 0
	ST_NodeVerticalAlignmentT     ST_NodeVerticalAlignment = 1
	ST_NodeVerticalAlignmentMid   ST_NodeVerticalAlignment = 2
	ST_NodeVerticalAlignmentB     ST_NodeVerticalAlignment = 3
)

func (_cgge *ST_ChildAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fceg, _acbc := d.Token()
	if _acbc != nil {
		return _acbc
	}
	if _bdea, _aecfe := _fceg.(_d.EndElement); _aecfe && _bdea.Name == start.Name {
		*_cgge = 1
		return nil
	}
	if _bfbbd, _gdbdc := _fceg.(_d.CharData); !_gdbdc {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fceg)
	} else {
		switch string(_bfbbd) {
		case "":
			*_cgge = 0
		case "\u0074":
			*_cgge = 1
		case "\u0062":
			*_cgge = 2
		case "\u006c":
			*_cgge = 3
		case "\u0072":
			*_cgge = 4
		}
	}
	_fceg, _acbc = d.Token()
	if _acbc != nil {
		return _acbc
	}
	if _egegf, _ddgbd := _fceg.(_d.EndElement); _ddgbd && _egegf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fceg)
}
func NewCT_Pt() *CT_Pt { _gcagg := &CT_Pt{}; return _gcagg }
func (_bfa *CT_Algorithm) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bfa.TypeAttr = ST_AlgorithmType(1)
	for _, _ade := range start.Attr {
		if _ade.Name.Local == "\u0074\u0079\u0070\u0065" {
			_bfa.TypeAttr.UnmarshalXMLAttr(_ade)
			continue
		}
		if _ade.Name.Local == "\u0072\u0065\u0076" {
			_eaeg, _faf := _ae.ParseUint(_ade.Value, 10, 32)
			if _faf != nil {
				return _faf
			}
			_cfa := uint32(_eaeg)
			_bfa.RevAttr = &_cfa
			continue
		}
	}
_adf:
	for {
		_bfab, _bgf := d.Token()
		if _bgf != nil {
			return _bgf
		}
		switch _dfe := _bfab.(type) {
		case _d.StartElement:
			switch _dfe.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0061\u0072a\u006d"}:
				_fc := NewCT_Parameter()
				if _acd := d.DecodeElement(_fc, &_dfe); _acd != nil {
					return _acd
				}
				_bfa.Param = append(_bfa.Param, _fc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bfa.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _dfa := d.DecodeElement(_bfa.ExtLst, &_dfe); _dfa != nil {
					return _dfa
				}
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_A\u006c\u0067o\u0072\u0069\u0074\u0068\u006d\u0020\u0025\u0076", _dfe.Name)
				if _eed := d.Skip(); _eed != nil {
					return _eed
				}
			}
		case _d.EndElement:
			break _adf
		case _d.CharData:
		}
	}
	return nil
}

const (
	ST_FunctionTypeUnset    ST_FunctionType = 0
	ST_FunctionTypeCnt      ST_FunctionType = 1
	ST_FunctionTypePos      ST_FunctionType = 2
	ST_FunctionTypeRevPos   ST_FunctionType = 3
	ST_FunctionTypePosEven  ST_FunctionType = 4
	ST_FunctionTypePosOdd   ST_FunctionType = 5
	ST_FunctionTypeVar      ST_FunctionType = 6
	ST_FunctionTypeDepth    ST_FunctionType = 7
	ST_FunctionTypeMaxDepth ST_FunctionType = 8
)

// Validate validates the CT_Choose and its children
func (_gdb *CT_Choose) Validate() error {
	return _gdb.ValidateWithPath("\u0043T\u005f\u0043\u0068\u006f\u006f\u0073e")
}
func (_aacac *ST_HueDir) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_aacac = 0
	case "\u0063\u0077":
		*_aacac = 1
	case "\u0063\u0063\u0077":
		*_aacac = 2
	}
	return nil
}
func (_aggd *CT_PtList) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _aggd.Pt != nil {
		_ecgd := _d.StartElement{Name: _d.Name{Local: "\u0070\u0074"}}
		for _, _ebbf := range _aggd.Pt {
			e.EncodeElement(_ebbf, _ecgd)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_ddfd ST_RotationPath) String() string {
	switch _ddfd {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h"
	}
	return ""
}
func ParseSliceST_Ints(s string) (ST_Ints, error) { return ST_Ints{}, nil }

type ST_FunctionOperator byte

func (_dbcef ST_AxisType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_dbcef.String(), start)
}
func (_fgag *CT_SampleData) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _fgag.UseDefAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0075\u0073\u0065\u0044\u0065\u0066"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_fgag.UseDefAttr))})
	}
	e.EncodeToken(start)
	if _fgag.DataModel != nil {
		_aecca := _d.StartElement{Name: _d.Name{Local: "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l"}}
		e.EncodeElement(_fgag.DataModel, _aecca)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_gd *CT_AnimLvl) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _dfaa := range start.Attr {
		if _dfaa.Name.Local == "\u0076\u0061\u006c" {
			_gd.ValAttr.UnmarshalXMLAttr(_dfaa)
			continue
		}
	}
	for {
		_dae, _caf := d.Token()
		if _caf != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c\u003a\u0020%\u0073", _caf)
		}
		if _adb, _bdd := _dae.(_d.EndElement); _bdd && _adb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_cacc ST_FunctionArgument) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _cacc.ST_VariableType != ST_VariableTypeUnset {
		e.EncodeToken(_d.CharData(_cacc.ST_VariableType.String()))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}

type ST_LinearDirection byte

func (_eba *CT_DataModel) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_eba.PtLst = NewCT_PtList()
_cbfc:
	for {
		_dgbd, _fgcd := d.Token()
		if _fgcd != nil {
			return _fgcd
		}
		switch _gee := _dgbd.(type) {
		case _d.StartElement:
			switch _gee.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0074\u004cs\u0074"}:
				if _cecc := d.DecodeElement(_eba.PtLst, &_gee); _cecc != nil {
					return _cecc
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0078\u006e\u004c\u0073\u0074"}:
				_eba.CxnLst = NewCT_CxnList()
				if _bdbf := d.DecodeElement(_eba.CxnLst, &_gee); _bdbf != nil {
					return _bdbf
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0062\u0067"}:
				_eba.Bg = _ca.NewCT_BackgroundFormatting()
				if _dcb := d.DecodeElement(_eba.Bg, &_gee); _dcb != nil {
					return _dcb
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0077\u0068\u006fl\u0065"}:
				_eba.Whole = _ca.NewCT_WholeE2oFormatting()
				if _fcde := d.DecodeElement(_eba.Whole, &_gee); _fcde != nil {
					return _fcde
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_eba.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _dbe := d.DecodeElement(_eba.ExtLst, &_gee); _dbe != nil {
					return _dbe
				}
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_D\u0061\u0074a\u004d\u006f\u0064\u0065\u006c\u0020\u0025\u0076", _gee.Name)
				if _ceg := d.Skip(); _ceg != nil {
					return _ceg
				}
			}
		case _d.EndElement:
			break _cbfc
		case _d.CharData:
		}
	}
	return nil
}
func (_abde *LayoutDef) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006ca\u0079\u006f\u0075\u0074\u0044\u0065f"
	return _abde.CT_DiagramDefinition.MarshalXML(e, start)
}

// Validate validates the ColorsDef and its children
func (_edff *ColorsDef) Validate() error {
	return _edff.ValidateWithPath("\u0043o\u006c\u006f\u0072\u0073\u0044\u0065f")
}
func (_abdf ST_FunctionValue) String() string {
	if _abdf.Int32 != nil {
		return _c.Sprintf("\u0025\u0076", *_abdf.Int32)
	}
	if _abdf.Bool != nil {
		return _c.Sprintf("\u0025\u0076", *_abdf.Bool)
	}
	if _abdf.ST_Direction != ST_DirectionUnset {
		return _abdf.ST_Direction.String()
	}
	if _abdf.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		return _abdf.ST_HierBranchStyle.String()
	}
	if _abdf.ST_AnimOneStr != ST_AnimOneStrUnset {
		return _abdf.ST_AnimOneStr.String()
	}
	if _abdf.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		return _abdf.ST_AnimLvlStr.String()
	}
	if _abdf.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		return _abdf.ST_ResizeHandlesStr.String()
	}
	return ""
}
func (_bcce *ST_ConnectorDimension) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bcce = 0
	case "\u0031\u0044":
		*_bcce = 1
	case "\u0032\u0044":
		*_bcce = 2
	case "\u0063\u0075\u0073\u0074":
		*_bcce = 3
	}
	return nil
}
func (_dcbb *ST_NodeHorizontalAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fbadd, _ebeec := d.Token()
	if _ebeec != nil {
		return _ebeec
	}
	if _acfcb, _bfaec := _fbadd.(_d.EndElement); _bfaec && _acfcb.Name == start.Name {
		*_dcbb = 1
		return nil
	}
	if _bage, _ggdbg := _fbadd.(_d.CharData); !_ggdbg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fbadd)
	} else {
		switch string(_bage) {
		case "":
			*_dcbb = 0
		case "\u006c":
			*_dcbb = 1
		case "\u0063\u0074\u0072":
			*_dcbb = 2
		case "\u0072":
			*_dcbb = 3
		}
	}
	_fbadd, _ebeec = d.Token()
	if _ebeec != nil {
		return _ebeec
	}
	if _cbcb, _bffe := _fbadd.(_d.EndElement); _bffe && _cbcb.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fbadd)
}

type CT_ColorTransform struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	Title        []*CT_CTName
	Desc         []*CT_CTDescription
	CatLst       *CT_CTCategories
	StyleLbl     []*CT_CTStyleLabel
	ExtLst       *_ca.CT_OfficeArtExtensionList
}

func (_edaf ST_NodeVerticalAlignment) String() string {
	switch _edaf {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u006d\u0069\u0064"
	case 3:
		return "\u0062"
	}
	return ""
}
func NewColorsDefHdr() *ColorsDefHdr {
	_cdae := &ColorsDefHdr{}
	_cdae.CT_ColorTransformHeader = *NewCT_ColorTransformHeader()
	return _cdae
}
func (_gfea ST_CxnType) String() string {
	switch _gfea {
	case 0:
		return ""
	case 1:
		return "\u0070\u0061\u0072O\u0066"
	case 2:
		return "\u0070\u0072\u0065\u0073\u004f\u0066"
	case 3:
		return "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff"
	case 4:
		return "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"
	}
	return ""
}

type ColorsDefHdrLst struct{ CT_ColorTransformHeaderLst }

func (_efded ST_LayoutShapeType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _efded.ST_ShapeType != _ca.ST_ShapeTypeUnset {
		e.EncodeToken(_d.CharData(_efded.ST_ShapeType.String()))
	}
	if _efded.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		e.EncodeToken(_d.CharData(_efded.ST_OutputShapeType.String()))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}
func (_bcef ST_OutputShapeType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_agbcf := _d.Attr{}
	_agbcf.Name = name
	switch _bcef {
	case ST_OutputShapeTypeUnset:
		_agbcf.Value = ""
	case ST_OutputShapeTypeNone:
		_agbcf.Value = "\u006e\u006f\u006e\u0065"
	case ST_OutputShapeTypeConn:
		_agbcf.Value = "\u0063\u006f\u006e\u006e"
	}
	return _agbcf, nil
}
func (_dfc *AG_IteratorAttributes) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _ged := range start.Attr {
		if _ged.Name.Local == "\u0061\u0078\u0069\u0073" {
			_fb, _cc := ParseSliceST_AxisTypes(_ged.Value)
			if _cc != nil {
				return _cc
			}
			_dfc.AxisAttr = &_fb
			continue
		}
		if _ged.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_db, _ba := ParseSliceST_ElementTypes(_ged.Value)
			if _ba != nil {
				return _ba
			}
			_dfc.PtTypeAttr = &_db
			continue
		}
		if _ged.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_ag, _cac := ParseSliceST_Booleans(_ged.Value)
			if _cac != nil {
				return _cac
			}
			_dfc.HideLastTransAttr = &_ag
			continue
		}
		if _ged.Name.Local == "\u0073\u0074" {
			_abf, _cd := ParseSliceST_Ints(_ged.Value)
			if _cd != nil {
				return _cd
			}
			_dfc.StAttr = &_abf
			continue
		}
		if _ged.Name.Local == "\u0063\u006e\u0074" {
			_gfc, _efe := ParseSliceST_UnsignedInts(_ged.Value)
			if _efe != nil {
				return _efe
			}
			_dfc.CntAttr = &_gfc
			continue
		}
		if _ged.Name.Local == "\u0073\u0074\u0065\u0070" {
			_fgg, _cce := ParseSliceST_Ints(_ged.Value)
			if _cce != nil {
				return _cce
			}
			_dfc.StepAttr = &_fgg
			continue
		}
	}
	for {
		_fgd, _dcg := d.Token()
		if _dcg != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0041\u0047\u005f\u0049\u0074\u0065\u0072\u0061\u0074\u006f\u0072\u0041t\u0074\u0072\u0069\u0062\u0075\u0074\u0065s\u003a\u0020\u0025\u0073", _dcg)
		}
		if _gb, _dcef := _fgd.(_d.EndElement); _dcef && _gb.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_CTDescription and its children, prefixing error messages with path
func (_cec *CT_CTDescription) ValidateWithPath(path string) error { return nil }
func (_ddfg *ST_SecondaryChildAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ddbb, _gdegg := d.Token()
	if _gdegg != nil {
		return _gdegg
	}
	if _gece, _gfeae := _ddbb.(_d.EndElement); _gfeae && _gece.Name == start.Name {
		*_ddfg = 1
		return nil
	}
	if _acab, _ceccc := _ddbb.(_d.CharData); !_ceccc {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ddbb)
	} else {
		switch string(_acab) {
		case "":
			*_ddfg = 0
		case "\u006e\u006f\u006e\u0065":
			*_ddfg = 1
		case "\u0074":
			*_ddfg = 2
		case "\u0062":
			*_ddfg = 3
		case "\u006c":
			*_ddfg = 4
		case "\u0072":
			*_ddfg = 5
		}
	}
	_ddbb, _gdegg = d.Token()
	if _gdegg != nil {
		return _gdegg
	}
	if _daeg, _cbfa := _ddbb.(_d.EndElement); _cbfa && _daeg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ddbb)
}
func (_aagfc *ST_AnimLvlStr) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_efbc, _fgbb := d.Token()
	if _fgbb != nil {
		return _fgbb
	}
	if _becda, _cdad := _efbc.(_d.EndElement); _cdad && _becda.Name == start.Name {
		*_aagfc = 1
		return nil
	}
	if _badb, _cede := _efbc.(_d.CharData); !_cede {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _efbc)
	} else {
		switch string(_badb) {
		case "":
			*_aagfc = 0
		case "\u006e\u006f\u006e\u0065":
			*_aagfc = 1
		case "\u006c\u0076\u006c":
			*_aagfc = 2
		case "\u0063\u0074\u0072":
			*_aagfc = 3
		}
	}
	_efbc, _fgbb = d.Token()
	if _fgbb != nil {
		return _fgbb
	}
	if _acfae, _cfea := _efbc.(_d.EndElement); _cfea && _acfae.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _efbc)
}
func NewCT_ResizeHandles() *CT_ResizeHandles { _fdfc := &CT_ResizeHandles{}; return _fdfc }
func (_adabb *LayoutDefHdrLst) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_adabb.CT_DiagramDefinitionHeaderLst = *NewCT_DiagramDefinitionHeaderLst()
_eagfd:
	for {
		_dbea, _cdge := d.Token()
		if _cdge != nil {
			return _cdge
		}
		switch _caee := _dbea.(type) {
		case _d.StartElement:
			switch _caee.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_fbefb := NewCT_DiagramDefinitionHeader()
				if _faea := d.DecodeElement(_fbefb, &_caee); _faea != nil {
					return _faea
				}
				_adabb.LayoutDefHdr = append(_adabb.LayoutDefHdr, _fbefb)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u004c\u0061\u0079\u006f\u0075\u0074D\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _caee.Name)
				if _gaec := d.Skip(); _gaec != nil {
					return _gaec
				}
			}
		case _d.EndElement:
			break _eagfd
		case _d.CharData:
		}
	}
	return nil
}
func (_feac ST_Breakpoint) ValidateWithPath(path string) error {
	switch _feac {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_feac))
	}
	return nil
}

// Validate validates the CT_HierBranchStyle and its children
func (_gaff *CT_HierBranchStyle) Validate() error {
	return _gaff.ValidateWithPath("\u0043T\u005fH\u0069\u0065\u0072\u0042\u0072a\u006e\u0063h\u0053\u0074\u0079\u006c\u0065")
}
func (_adbee ST_AutoTextRotation) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_adbee.String(), start)
}
func (_cegea ST_ModelId) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _cegea.Int32 != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", *_cegea.Int32)))
	}
	if _cegea.ST_Guid != nil {
		e.EncodeToken(_d.CharData(*_cegea.ST_Guid))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}
func (_bbfga *ColorsDef) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063o\u006c\u006f\u0072\u0073\u0044\u0065f"
	return _bbfga.CT_ColorTransform.MarshalXML(e, start)
}

type CT_Name struct {
	LangAttr *string
	ValAttr  string
}
type ColorsDef struct{ CT_ColorTransform }

func (_gfad ST_PyramidAccentTextMargin) ValidateWithPath(path string) error {
	switch _gfad {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gfad))
	}
	return nil
}

type ST_PyramidAccentPosition byte

const (
	ST_RotationPathUnset     ST_RotationPath = 0
	ST_RotationPathNone      ST_RotationPath = 1
	ST_RotationPathAlongPath ST_RotationPath = 2
)

func (_dfec *CT_Shape) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _dfec.RotAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u006f\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_dfec.RotAttr)})
	}
	if _dfec.TypeAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_dfec.TypeAttr)})
	}
	if _dfec.BlipAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u003a\u0062\u006c\u0069\u0070"}, Value: _c.Sprintf("\u0025\u0076", *_dfec.BlipAttr)})
	}
	if _dfec.ZOrderOffAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u007aO\u0072\u0064\u0065\u0072\u004f\u0066f"}, Value: _c.Sprintf("\u0025\u0076", *_dfec.ZOrderOffAttr)})
	}
	if _dfec.HideGeomAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0068\u0069\u0064\u0065\u0047\u0065\u006f\u006d"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_dfec.HideGeomAttr))})
	}
	if _dfec.LkTxEntryAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006ck\u0054\u0078\u0045\u006e\u0074\u0072y"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_dfec.LkTxEntryAttr))})
	}
	if _dfec.BlipPhldrAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0062l\u0069\u0070\u0050\u0068\u006c\u0064r"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_dfec.BlipPhldrAttr))})
	}
	e.EncodeToken(start)
	if _dfec.AdjLst != nil {
		_cbbcc := _d.StartElement{Name: _d.Name{Local: "\u0061\u0064\u006a\u004c\u0073\u0074"}}
		e.EncodeElement(_dfec.AdjLst, _cbbcc)
	}
	if _dfec.ExtLst != nil {
		_bdgbd := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_dfec.ExtLst, _bdgbd)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_beede *CT_Otherwise) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _fcffe := range start.Attr {
		if _fcffe.Name.Local == "\u006e\u0061\u006d\u0065" {
			_bagdc, _dbdc := _fcffe.Value, error(nil)
			if _dbdc != nil {
				return _dbdc
			}
			_beede.NameAttr = &_bagdc
			continue
		}
	}
_agd:
	for {
		_cbccd, _dgdg := d.Token()
		if _dgdg != nil {
			return _dgdg
		}
		switch _ddbe := _cbccd.(type) {
		case _d.StartElement:
			switch _ddbe.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_bggc := NewCT_Algorithm()
				if _becc := d.DecodeElement(_bggc, &_ddbe); _becc != nil {
					return _becc
				}
				_beede.Alg = append(_beede.Alg, _bggc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_abafb := NewCT_Shape()
				if _agaf := d.DecodeElement(_abafb, &_ddbe); _agaf != nil {
					return _agaf
				}
				_beede.Shape = append(_beede.Shape, _abafb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_fdgb := NewCT_PresentationOf()
				if _daaea := d.DecodeElement(_fdgb, &_ddbe); _daaea != nil {
					return _daaea
				}
				_beede.PresOf = append(_beede.PresOf, _fdgb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_bafag := NewCT_Constraints()
				if _ccfa := d.DecodeElement(_bafag, &_ddbe); _ccfa != nil {
					return _ccfa
				}
				_beede.ConstrLst = append(_beede.ConstrLst, _bafag)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_eaed := NewCT_Rules()
				if _gaed := d.DecodeElement(_eaed, &_ddbe); _gaed != nil {
					return _gaed
				}
				_beede.RuleLst = append(_beede.RuleLst, _eaed)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_dcbgg := NewCT_ForEach()
				if _fgbc := d.DecodeElement(_dcbgg, &_ddbe); _fgbc != nil {
					return _fgbc
				}
				_beede.ForEach = append(_beede.ForEach, _dcbgg)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_fcef := NewCT_LayoutNode()
				if _agface := d.DecodeElement(_fcef, &_ddbe); _agface != nil {
					return _agface
				}
				_beede.LayoutNode = append(_beede.LayoutNode, _fcef)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_gbdc := NewCT_Choose()
				if _cdgfc := d.DecodeElement(_gbdc, &_ddbe); _cdgfc != nil {
					return _cdgfc
				}
				_beede.Choose = append(_beede.Choose, _gbdc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_eda := _ca.NewCT_OfficeArtExtensionList()
				if _aeed := d.DecodeElement(_eda, &_ddbe); _aeed != nil {
					return _aeed
				}
				_beede.ExtLst = append(_beede.ExtLst, _eda)
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_O\u0074\u0068e\u0072\u0077\u0069\u0073\u0065\u0020\u0025\u0076", _ddbe.Name)
				if _fdfd := d.Skip(); _fdfd != nil {
					return _fdfd
				}
			}
		case _d.EndElement:
			break _agd
		case _d.CharData:
		}
	}
	return nil
}
func (_bgbf ST_DiagramHorizontalAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dgabe := _d.Attr{}
	_dgabe.Name = name
	switch _bgbf {
	case ST_DiagramHorizontalAlignmentUnset:
		_dgabe.Value = ""
	case ST_DiagramHorizontalAlignmentL:
		_dgabe.Value = "\u006c"
	case ST_DiagramHorizontalAlignmentCtr:
		_dgabe.Value = "\u0063\u0074\u0072"
	case ST_DiagramHorizontalAlignmentR:
		_dgabe.Value = "\u0072"
	case ST_DiagramHorizontalAlignmentNone:
		_dgabe.Value = "\u006e\u006f\u006e\u0065"
	}
	return _dgabe, nil
}
func (_addg ST_Breakpoint) Validate() error { return _addg.ValidateWithPath("") }

type CT_Algorithm struct {
	TypeAttr ST_AlgorithmType
	RevAttr  *uint32
	Param    []*CT_Parameter
	ExtLst   *_ca.CT_OfficeArtExtensionList
}

// Validate validates the CT_Constraint and its children
func (_agb *CT_Constraint) Validate() error {
	return _agb.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074")
}
func (_efeba ST_NodeHorizontalAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_cdafd := _d.Attr{}
	_cdafd.Name = name
	switch _efeba {
	case ST_NodeHorizontalAlignmentUnset:
		_cdafd.Value = ""
	case ST_NodeHorizontalAlignmentL:
		_cdafd.Value = "\u006c"
	case ST_NodeHorizontalAlignmentCtr:
		_cdafd.Value = "\u0063\u0074\u0072"
	case ST_NodeHorizontalAlignmentR:
		_cdafd.Value = "\u0072"
	}
	return _cdafd, nil
}
func (_eagd ST_ChildAlignment) String() string {
	switch _eagd {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0062"
	case 3:
		return "\u006c"
	case 4:
		return "\u0072"
	}
	return ""
}

type ST_ConnectorPoint byte

func (_fbeff *ST_ArrowheadStyle) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cgac, _cbegb := d.Token()
	if _cbegb != nil {
		return _cbegb
	}
	if _bafae, _accgg := _cgac.(_d.EndElement); _accgg && _bafae.Name == start.Name {
		*_fbeff = 1
		return nil
	}
	if _beefe, _fbee := _cgac.(_d.CharData); !_fbee {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cgac)
	} else {
		switch string(_beefe) {
		case "":
			*_fbeff = 0
		case "\u0061\u0075\u0074\u006f":
			*_fbeff = 1
		case "\u0061\u0072\u0072":
			*_fbeff = 2
		case "\u006e\u006f\u0041r\u0072":
			*_fbeff = 3
		}
	}
	_cgac, _cbegb = d.Token()
	if _cbegb != nil {
		return _cbegb
	}
	if _debe, _gffa := _cgac.(_d.EndElement); _gffa && _debe.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cgac)
}
func NewCT_Category() *CT_Category { _dbb := &CT_Category{}; return _dbb }

// Validate validates the StyleDefHdr and its children
func (_edfgc *StyleDefHdr) Validate() error {
	return _edfgc.ValidateWithPath("S\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072")
}
func (_gadd *CT_Direction) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _gadd.ValAttr != ST_DirectionUnset {
		_efa, _dcdg := _gadd.ValAttr.MarshalXMLAttr(_d.Name{Local: "\u0076\u0061\u006c"})
		if _dcdg != nil {
			return _dcdg
		}
		start.Attr = append(start.Attr, _efa)
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

type CT_AnimLvl struct{ ValAttr ST_AnimLvlStr }

func (_cggc ST_TextAnchorHorizontal) ValidateWithPath(path string) error {
	switch _cggc {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cggc))
	}
	return nil
}

// Validate validates the CT_ColorTransformHeaderLst and its children
func (_eabe *CT_ColorTransformHeaderLst) Validate() error {
	return _eabe.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061n\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065r\u004c\u0073\u0074")
}
func (_gdae *CT_NumericRule) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _caec := range start.Attr {
		if _caec.Name.Local == "\u0076\u0061\u006c" {
			_cgaf, _ffce := _ae.ParseFloat(_caec.Value, 64)
			if _ffce != nil {
				return _ffce
			}
			_gdae.ValAttr = &_cgaf
			continue
		}
		if _caec.Name.Local == "\u0066\u0061\u0063\u0074" {
			_fgcc, _bgadg := _ae.ParseFloat(_caec.Value, 64)
			if _bgadg != nil {
				return _bgadg
			}
			_gdae.FactAttr = &_fgcc
			continue
		}
		if _caec.Name.Local == "\u006d\u0061\u0078" {
			_dfac, _agfc := _ae.ParseFloat(_caec.Value, 64)
			if _agfc != nil {
				return _agfc
			}
			_gdae.MaxAttr = &_dfac
			continue
		}
		if _caec.Name.Local == "\u0074\u0079\u0070\u0065" {
			_gdae.TypeAttr.UnmarshalXMLAttr(_caec)
			continue
		}
		if _caec.Name.Local == "\u0066\u006f\u0072" {
			_gdae.ForAttr.UnmarshalXMLAttr(_caec)
			continue
		}
		if _caec.Name.Local == "\u0066o\u0072\u004e\u0061\u006d\u0065" {
			_edce, _fbdf := _caec.Value, error(nil)
			if _fbdf != nil {
				return _fbdf
			}
			_gdae.ForNameAttr = &_edce
			continue
		}
		if _caec.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_gdae.PtTypeAttr.UnmarshalXMLAttr(_caec)
			continue
		}
	}
_eaef:
	for {
		_dbca, _fgcdf := d.Token()
		if _fgcdf != nil {
			return _fgcdf
		}
		switch _ecbcc := _dbca.(type) {
		case _d.StartElement:
			switch _ecbcc.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_gdae.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _gfbdg := d.DecodeElement(_gdae.ExtLst, &_ecbcc); _gfbdg != nil {
					return _gfbdg
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004e\u0075\u006d\u0065\u0072\u0069\u0063R\u0075l\u0065\u0020\u0025\u0076", _ecbcc.Name)
				if _bfaea := d.Skip(); _bfaea != nil {
					return _bfaea
				}
			}
		case _d.EndElement:
			break _eaef
		case _d.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AdjLst and its children
func (_fba *CT_AdjLst) Validate() error {
	return _fba.ValidateWithPath("\u0043T\u005f\u0041\u0064\u006a\u004c\u0073t")
}

type ST_DiagramTextAlignment byte

func (_aeee *CT_Choose) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _aeee.NameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_aeee.NameAttr)})
	}
	e.EncodeToken(start)
	_eedc := _d.StartElement{Name: _d.Name{Local: "\u0069\u0066"}}
	for _, _cgdc := range _aeee.If {
		e.EncodeElement(_cgdc, _eedc)
	}
	if _aeee.Else != nil {
		_dab := _d.StartElement{Name: _d.Name{Local: "\u0065\u006c\u0073\u0065"}}
		e.EncodeElement(_aeee.Else, _dab)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_fcfg *ST_ConstraintRelationship) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fcfg = 0
	case "\u0073\u0065\u006c\u0066":
		*_fcfg = 1
	case "\u0063\u0068":
		*_fcfg = 2
	case "\u0064\u0065\u0073":
		*_fcfg = 3
	}
	return nil
}
func NewLayoutDefHdr() *LayoutDefHdr {
	_gegd := &LayoutDefHdr{}
	_gegd.CT_DiagramDefinitionHeader = *NewCT_DiagramDefinitionHeader()
	return _gegd
}
func (_afcf ST_DiagramTextAlignment) ValidateWithPath(path string) error {
	switch _afcf {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_afcf))
	}
	return nil
}
func (_bggdc ST_LinearDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gdbeb := _d.Attr{}
	_gdbeb.Name = name
	switch _bggdc {
	case ST_LinearDirectionUnset:
		_gdbeb.Value = ""
	case ST_LinearDirectionFromL:
		_gdbeb.Value = "\u0066\u0072\u006fm\u004c"
	case ST_LinearDirectionFromR:
		_gdbeb.Value = "\u0066\u0072\u006fm\u0052"
	case ST_LinearDirectionFromT:
		_gdbeb.Value = "\u0066\u0072\u006fm\u0054"
	case ST_LinearDirectionFromB:
		_gdbeb.Value = "\u0066\u0072\u006fm\u0042"
	}
	return _gdbeb, nil
}
func (_ccdg ST_HueDir) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ccdg.String(), start)
}

const (
	ST_CxnTypeUnset               ST_CxnType = 0
	ST_CxnTypeParOf               ST_CxnType = 1
	ST_CxnTypePresOf              ST_CxnType = 2
	ST_CxnTypePresParOf           ST_CxnType = 3
	ST_CxnTypeUnknownRelationship ST_CxnType = 4
)

func (_dbg *CT_Colors) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _dbg.MethAttr != ST_ClrAppMethodUnset {
		_fedgc, _dda := _dbg.MethAttr.MarshalXMLAttr(_d.Name{Local: "\u006d\u0065\u0074\u0068"})
		if _dda != nil {
			return _dda
		}
		start.Attr = append(start.Attr, _fedgc)
	}
	if _dbg.HueDirAttr != ST_HueDirUnset {
		_gbaa, _dffb := _dbg.HueDirAttr.MarshalXMLAttr(_d.Name{Local: "\u0068\u0075\u0065\u0044\u0069\u0072"})
		if _dffb != nil {
			return _dffb
		}
		start.Attr = append(start.Attr, _gbaa)
	}
	e.EncodeToken(start)
	if _dbg.EG_ColorChoice != nil {
		for _, _adcc := range _dbg.EG_ColorChoice {
			_adcc.MarshalXML(e, _d.StartElement{})
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_DiagramDefinitionHeaderLst and its children
func (_edee *CT_DiagramDefinitionHeaderLst) Validate() error {
	return _edee.ValidateWithPath("\u0043\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065rL\u0073\u0074")
}
func (_cdcf ST_VerticalAlignment) ValidateWithPath(path string) error {
	switch _cdcf {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cdcf))
	}
	return nil
}
func (_dcga *CT_Otherwise) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _dcga.NameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_dcga.NameAttr)})
	}
	e.EncodeToken(start)
	if _dcga.Alg != nil {
		_fcfa := _d.StartElement{Name: _d.Name{Local: "\u0061\u006c\u0067"}}
		for _, _cdbac := range _dcga.Alg {
			e.EncodeElement(_cdbac, _fcfa)
		}
	}
	if _dcga.Shape != nil {
		_acga := _d.StartElement{Name: _d.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _deed := range _dcga.Shape {
			e.EncodeElement(_deed, _acga)
		}
	}
	if _dcga.PresOf != nil {
		_bedfe := _d.StartElement{Name: _d.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _becfa := range _dcga.PresOf {
			e.EncodeElement(_becfa, _bedfe)
		}
	}
	if _dcga.ConstrLst != nil {
		_dgab := _d.StartElement{Name: _d.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _abaf := range _dcga.ConstrLst {
			e.EncodeElement(_abaf, _dgab)
		}
	}
	if _dcga.RuleLst != nil {
		_cffg := _d.StartElement{Name: _d.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _ffgcd := range _dcga.RuleLst {
			e.EncodeElement(_ffgcd, _cffg)
		}
	}
	if _dcga.ForEach != nil {
		_fgff := _d.StartElement{Name: _d.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _gcdf := range _dcga.ForEach {
			e.EncodeElement(_gcdf, _fgff)
		}
	}
	if _dcga.LayoutNode != nil {
		_adfg := _d.StartElement{Name: _d.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _egce := range _dcga.LayoutNode {
			e.EncodeElement(_egce, _adfg)
		}
	}
	if _dcga.Choose != nil {
		_ccbb := _d.StartElement{Name: _d.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _ffa := range _dcga.Choose {
			e.EncodeElement(_ffa, _ccbb)
		}
	}
	if _dcga.ExtLst != nil {
		_ebfe := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _egfd := range _dcga.ExtLst {
			e.EncodeElement(_egfd, _ebfe)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_dbdcg ST_DiagramTextAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gbfba := _d.Attr{}
	_gbfba.Name = name
	switch _dbdcg {
	case ST_DiagramTextAlignmentUnset:
		_gbfba.Value = ""
	case ST_DiagramTextAlignmentL:
		_gbfba.Value = "\u006c"
	case ST_DiagramTextAlignmentCtr:
		_gbfba.Value = "\u0063\u0074\u0072"
	case ST_DiagramTextAlignmentR:
		_gbfba.Value = "\u0072"
	}
	return _gbfba, nil
}
func (_dcebg ST_ContinueDirection) String() string {
	switch _dcebg {
	case 0:
		return ""
	case 1:
		return "\u0072\u0065\u0076\u0044\u0069\u0072"
	case 2:
		return "\u0073a\u006d\u0065\u0044\u0069\u0072"
	}
	return ""
}

// ValidateWithPath validates the CT_Constraints and its children, prefixing error messages with path
func (_gaf *CT_Constraints) ValidateWithPath(path string) error {
	for _cgaa, _edbg := range _gaf.Constr {
		if _ddef := _edbg.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u006f\u006e\u0073\u0074\u0072\u005b\u0025\u0064\u005d", path, _cgaa)); _ddef != nil {
			return _ddef
		}
	}
	return nil
}

// ST_ModelId is a union type
type ST_ModelId struct {
	Int32   *int32
	ST_Guid *string
}

func NewCT_Adj() *CT_Adj { _eae := &CT_Adj{}; _eae.IdxAttr = 1; return _eae }
func (_gaeeae ST_StartingElement) ValidateWithPath(path string) error {
	switch _gaeeae {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gaeeae))
	}
	return nil
}
func (_cacbd *ST_DiagramHorizontalAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_cacbd = 0
	case "\u006c":
		*_cacbd = 1
	case "\u0063\u0074\u0072":
		*_cacbd = 2
	case "\u0072":
		*_cacbd = 3
	case "\u006e\u006f\u006e\u0065":
		*_cacbd = 4
	}
	return nil
}
func (_bgcg *ST_ResizeHandlesStr) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bgcg = 0
	case "\u0065\u0078\u0061c\u0074":
		*_bgcg = 1
	case "\u0072\u0065\u006c":
		*_bgcg = 2
	}
	return nil
}

// ValidateWithPath validates the CT_TextProps and its children, prefixing error messages with path
func (_cbgce *CT_TextProps) ValidateWithPath(path string) error {
	if _cbgce.Sp3d != nil {
		if _fgfd := _cbgce.Sp3d.ValidateWithPath(path + "\u002f\u0053\u00703\u0064"); _fgfd != nil {
			return _fgfd
		}
	}
	if _cbgce.FlatTx != nil {
		if _egff := _cbgce.FlatTx.ValidateWithPath(path + "\u002fF\u006c\u0061\u0074\u0054\u0078"); _egff != nil {
			return _egff
		}
	}
	return nil
}

type ST_RotationPath byte

func NewCT_StyleDefinition() *CT_StyleDefinition { _ffaac := &CT_StyleDefinition{}; return _ffaac }
func (_faga ST_TextDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_faga.String(), start)
}
func (_face *CT_ColorTransform) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _face.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_face.UniqueIdAttr)})
	}
	if _face.MinVerAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _c.Sprintf("\u0025\u0076", *_face.MinVerAttr)})
	}
	e.EncodeToken(start)
	if _face.Title != nil {
		_gbge := _d.StartElement{Name: _d.Name{Local: "\u0074\u0069\u0074l\u0065"}}
		for _, _cdb := range _face.Title {
			e.EncodeElement(_cdb, _gbge)
		}
	}
	if _face.Desc != nil {
		_gcac := _d.StartElement{Name: _d.Name{Local: "\u0064\u0065\u0073\u0063"}}
		for _, _ddcc := range _face.Desc {
			e.EncodeElement(_ddcc, _gcac)
		}
	}
	if _face.CatLst != nil {
		_abce := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_face.CatLst, _abce)
	}
	if _face.StyleLbl != nil {
		_fcab := _d.StartElement{Name: _d.Name{Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}}
		for _, _bbge := range _face.StyleLbl {
			e.EncodeElement(_bbge, _fcab)
		}
	}
	if _face.ExtLst != nil {
		_cggd := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_face.ExtLst, _cggd)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_ebfbg ST_AxisType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_fdcfc := _d.Attr{}
	_fdcfc.Name = name
	switch _ebfbg {
	case ST_AxisTypeUnset:
		_fdcfc.Value = ""
	case ST_AxisTypeSelf:
		_fdcfc.Value = "\u0073\u0065\u006c\u0066"
	case ST_AxisTypeCh:
		_fdcfc.Value = "\u0063\u0068"
	case ST_AxisTypeDes:
		_fdcfc.Value = "\u0064\u0065\u0073"
	case ST_AxisTypeDesOrSelf:
		_fdcfc.Value = "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf"
	case ST_AxisTypePar:
		_fdcfc.Value = "\u0070\u0061\u0072"
	case ST_AxisTypeAncst:
		_fdcfc.Value = "\u0061\u006e\u0063s\u0074"
	case ST_AxisTypeAncstOrSelf:
		_fdcfc.Value = "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066"
	case ST_AxisTypeFollowSib:
		_fdcfc.Value = "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b"
	case ST_AxisTypePrecedSib:
		_fdcfc.Value = "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b"
	case ST_AxisTypeFollow:
		_fdcfc.Value = "\u0066\u006f\u006c\u006c\u006f\u0077"
	case ST_AxisTypePreced:
		_fdcfc.Value = "\u0070\u0072\u0065\u0063\u0065\u0064"
	case ST_AxisTypeRoot:
		_fdcfc.Value = "\u0072\u006f\u006f\u0074"
	case ST_AxisTypeNone:
		_fdcfc.Value = "\u006e\u006f\u006e\u0065"
	}
	return _fdcfc, nil
}

type ST_Direction byte

// ValidateWithPath validates the CT_ColorTransformHeader and its children, prefixing error messages with path
func (_gbe *CT_ColorTransformHeader) ValidateWithPath(path string) error {
	for _befb, _fedg := range _gbe.Title {
		if _bgb := _fedg.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _befb)); _bgb != nil {
			return _bgb
		}
	}
	for _bcbg, _dfbg := range _gbe.Desc {
		if _babc := _dfbg.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _bcbg)); _babc != nil {
			return _babc
		}
	}
	if _gbe.CatLst != nil {
		if _dbbc := _gbe.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _dbbc != nil {
			return _dbbc
		}
	}
	if _gbe.ExtLst != nil {
		if _ffb := _gbe.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _ffb != nil {
			return _ffb
		}
	}
	return nil
}
func NewCT_CTStyleLabel() *CT_CTStyleLabel { _bee := &CT_CTStyleLabel{}; return _bee }

const (
	ST_OutputShapeTypeUnset ST_OutputShapeType = 0
	ST_OutputShapeTypeNone  ST_OutputShapeType = 1
	ST_OutputShapeTypeConn  ST_OutputShapeType = 2
)

// ValidateWithPath validates the CT_ColorTransformHeaderLst and its children, prefixing error messages with path
func (_gceg *CT_ColorTransformHeaderLst) ValidateWithPath(path string) error {
	for _cae, _gceb := range _gceg.ColorsDefHdr {
		if _aeaa := _gceb.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043ol\u006f\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u005b\u0025\u0064\u005d", path, _cae)); _aeaa != nil {
			return _aeaa
		}
	}
	return nil
}
func (_fadd *CT_SDCategories) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _fadd.Cat != nil {
		_fddfg := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074"}}
		for _, _bcgea := range _fadd.Cat {
			e.EncodeElement(_bcgea, _fddfg)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_bdafg *ST_FunctionOperator) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bdafg = 0
	case "\u0065\u0071\u0075":
		*_bdafg = 1
	case "\u006e\u0065\u0071":
		*_bdafg = 2
	case "\u0067\u0074":
		*_bdafg = 3
	case "\u006c\u0074":
		*_bdafg = 4
	case "\u0067\u0074\u0065":
		*_bdafg = 5
	case "\u006c\u0074\u0065":
		*_bdafg = 6
	}
	return nil
}
func (_daea ST_ChildAlignment) Validate() error { return _daea.ValidateWithPath("") }
func (_eeed *CT_DiagramDefinitionHeaderLst) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _eeed.LayoutDefHdr != nil {
		_afbe := _d.StartElement{Name: _d.Name{Local: "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"}}
		for _, _fgcg := range _eeed.LayoutDefHdr {
			e.EncodeElement(_fgcg, _afbe)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_ccfe ST_AxisType) String() string {
	switch _ccfe {
	case 0:
		return ""
	case 1:
		return "\u0073\u0065\u006c\u0066"
	case 2:
		return "\u0063\u0068"
	case 3:
		return "\u0064\u0065\u0073"
	case 4:
		return "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf"
	case 5:
		return "\u0070\u0061\u0072"
	case 6:
		return "\u0061\u006e\u0063s\u0074"
	case 7:
		return "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066"
	case 8:
		return "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b"
	case 9:
		return "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b"
	case 10:
		return "\u0066\u006f\u006c\u006c\u006f\u0077"
	case 11:
		return "\u0070\u0072\u0065\u0063\u0065\u0064"
	case 12:
		return "\u0072\u006f\u006f\u0074"
	case 13:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}

type ST_DiagramHorizontalAlignment byte

func (_edddf ST_DiagramHorizontalAlignment) String() string {
	switch _edddf {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0063\u0074\u0072"
	case 3:
		return "\u0072"
	case 4:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}
func (_ebbc *CT_SDDescription) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _dabc := range start.Attr {
		if _dabc.Name.Local == "\u006c\u0061\u006e\u0067" {
			_aefga, _eadaa := _dabc.Value, error(nil)
			if _eadaa != nil {
				return _eadaa
			}
			_ebbc.LangAttr = &_aefga
			continue
		}
		if _dabc.Name.Local == "\u0076\u0061\u006c" {
			_eeeeb, _edag := _dabc.Value, error(nil)
			if _edag != nil {
				return _edag
			}
			_ebbc.ValAttr = _eeeeb
			continue
		}
	}
	for {
		_egbb, _bddad := d.Token()
		if _bddad != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0053\u0044\u0044\u0065s\u0063r\u0069\u0070\u0074\u0069\u006f\u006e\u003a \u0025\u0073", _bddad)
		}
		if _efeb, _aceb := _egbb.(_d.EndElement); _aceb && _efeb.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_LayoutVariablePropertySet and its children, prefixing error messages with path
func (_aced *CT_LayoutVariablePropertySet) ValidateWithPath(path string) error {
	if _aced.OrgChart != nil {
		if _ggdf := _aced.OrgChart.ValidateWithPath(path + "\u002fO\u0072\u0067\u0043\u0068\u0061\u0072t"); _ggdf != nil {
			return _ggdf
		}
	}
	if _aced.ChMax != nil {
		if _fegb := _aced.ChMax.ValidateWithPath(path + "\u002f\u0043\u0068\u004d\u0061\u0078"); _fegb != nil {
			return _fegb
		}
	}
	if _aced.ChPref != nil {
		if _gdge := _aced.ChPref.ValidateWithPath(path + "\u002fC\u0068\u0050\u0072\u0065\u0066"); _gdge != nil {
			return _gdge
		}
	}
	if _aced.BulletEnabled != nil {
		if _gccb := _aced.BulletEnabled.ValidateWithPath(path + "\u002f\u0042\u0075\u006c\u006c\u0065\u0074\u0045\u006ea\u0062\u006c\u0065\u0064"); _gccb != nil {
			return _gccb
		}
	}
	if _aced.Dir != nil {
		if _adbb := _aced.Dir.ValidateWithPath(path + "\u002f\u0044\u0069\u0072"); _adbb != nil {
			return _adbb
		}
	}
	if _aced.HierBranch != nil {
		if _efcd := _aced.HierBranch.ValidateWithPath(path + "/\u0048\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"); _efcd != nil {
			return _efcd
		}
	}
	if _aced.AnimOne != nil {
		if _gegb := _aced.AnimOne.ValidateWithPath(path + "\u002f\u0041\u006e\u0069\u006d\u004f\u006e\u0065"); _gegb != nil {
			return _gegb
		}
	}
	if _aced.AnimLvl != nil {
		if _cdca := _aced.AnimLvl.ValidateWithPath(path + "\u002f\u0041\u006e\u0069\u006d\u004c\u0076\u006c"); _cdca != nil {
			return _cdca
		}
	}
	if _aced.ResizeHandles != nil {
		if _ccde := _aced.ResizeHandles.ValidateWithPath(path + "\u002f\u0052\u0065\u0073\u0069\u007a\u0065\u0048\u0061n\u0064\u006c\u0065\u0073"); _ccde != nil {
			return _ccde
		}
	}
	return nil
}

type ST_AnimOneStr byte

func (_bbac *ColorsDefHdr) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"
	return _bbac.CT_ColorTransformHeader.MarshalXML(e, start)
}
func (_dfbcf *ST_FunctionArgument) Validate() error { return _dfbcf.ValidateWithPath("") }
func (_edfb ST_Breakpoint) String() string {
	switch _edfb {
	case 0:
		return ""
	case 1:
		return "\u0065\u006e\u0064\u0043\u006e\u0076"
	case 2:
		return "\u0062\u0061\u006c"
	case 3:
		return "\u0066\u0069\u0078e\u0064"
	}
	return ""
}

// Validate validates the CT_PresentationOf and its children
func (_dbcb *CT_PresentationOf) Validate() error {
	return _dbcb.ValidateWithPath("\u0043\u0054\u005f\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u004f\u0066")
}
func (_bbeag ST_BoolOperator) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dcbac := _d.Attr{}
	_dcbac.Name = name
	switch _bbeag {
	case ST_BoolOperatorUnset:
		_dcbac.Value = ""
	case ST_BoolOperatorNone:
		_dcbac.Value = "\u006e\u006f\u006e\u0065"
	case ST_BoolOperatorEqu:
		_dcbac.Value = "\u0065\u0071\u0075"
	case ST_BoolOperatorGte:
		_dcbac.Value = "\u0067\u0074\u0065"
	case ST_BoolOperatorLte:
		_dcbac.Value = "\u006c\u0074\u0065"
	}
	return _dcbac, nil
}
func (_ddge *ST_FunctionOperator) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_facff, _cgea := d.Token()
	if _cgea != nil {
		return _cgea
	}
	if _gbaf, _adfff := _facff.(_d.EndElement); _adfff && _gbaf.Name == start.Name {
		*_ddge = 1
		return nil
	}
	if _bdfe, _eaaba := _facff.(_d.CharData); !_eaaba {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _facff)
	} else {
		switch string(_bdfe) {
		case "":
			*_ddge = 0
		case "\u0065\u0071\u0075":
			*_ddge = 1
		case "\u006e\u0065\u0071":
			*_ddge = 2
		case "\u0067\u0074":
			*_ddge = 3
		case "\u006c\u0074":
			*_ddge = 4
		case "\u0067\u0074\u0065":
			*_ddge = 5
		case "\u006c\u0074\u0065":
			*_ddge = 6
		}
	}
	_facff, _cgea = d.Token()
	if _cgea != nil {
		return _cgea
	}
	if _gfcc, _bcdgf := _facff.(_d.EndElement); _bcdgf && _gfcc.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _facff)
}
func (_egbbd ST_ArrowheadStyle) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_egbbd.String(), start)
}
func (_gecc ST_FlowDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_gecc.String(), start)
}
func (_adac *CT_Description) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gabf := range start.Attr {
		if _gabf.Name.Local == "\u006c\u0061\u006e\u0067" {
			_egfg, _dedg := _gabf.Value, error(nil)
			if _dedg != nil {
				return _dedg
			}
			_adac.LangAttr = &_egfg
			continue
		}
		if _gabf.Name.Local == "\u0076\u0061\u006c" {
			_bdfa, _abff := _gabf.Value, error(nil)
			if _abff != nil {
				return _abff
			}
			_adac.ValAttr = _bdfa
			continue
		}
	}
	for {
		_fcca, _ggg := d.Token()
		if _ggg != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fD\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e:\u0020\u0025\u0073", _ggg)
		}
		if _bcgb, _ccce := _fcca.(_d.EndElement); _ccce && _bcgb.Name == start.Name {
			break
		}
	}
	return nil
}

type ST_Breakpoint byte

// ValidateWithPath validates the CT_Category and its children, prefixing error messages with path
func (_ddg *CT_Category) ValidateWithPath(path string) error { return nil }
func (_ggcadf ST_Offset) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ggcadf.String(), start)
}
func (_de *CT_Algorithm) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	_bb, _fdd := _de.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
	if _fdd != nil {
		return _fdd
	}
	start.Attr = append(start.Attr, _bb)
	if _de.RevAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u0065\u0076"}, Value: _c.Sprintf("\u0025\u0076", *_de.RevAttr)})
	}
	e.EncodeToken(start)
	if _de.Param != nil {
		_eee := _d.StartElement{Name: _d.Name{Local: "\u0070\u0061\u0072a\u006d"}}
		for _, _cfe := range _de.Param {
			e.EncodeElement(_cfe, _eee)
		}
	}
	if _de.ExtLst != nil {
		_bbd := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_de.ExtLst, _bbd)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

const (
	ST_ArrowheadStyleUnset ST_ArrowheadStyle = 0
	ST_ArrowheadStyleAuto  ST_ArrowheadStyle = 1
	ST_ArrowheadStyleArr   ST_ArrowheadStyle = 2
	ST_ArrowheadStyleNoArr ST_ArrowheadStyle = 3
)

func (_bdbb ST_NodeHorizontalAlignment) Validate() error { return _bdbb.ValidateWithPath("") }

type CT_TextProps struct {
	Sp3d   *_ca.CT_Shape3D
	FlatTx *_ca.CT_FlatText
}

func (_cdag ST_ConstraintRelationship) ValidateWithPath(path string) error {
	switch _cdag {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cdag))
	}
	return nil
}
func (_cbec *ST_NodeVerticalAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cbdab, _fbgc := d.Token()
	if _fbgc != nil {
		return _fbgc
	}
	if _ddgbc, _fccd := _cbdab.(_d.EndElement); _fccd && _ddgbc.Name == start.Name {
		*_cbec = 1
		return nil
	}
	if _acad, _bfdac := _cbdab.(_d.CharData); !_bfdac {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cbdab)
	} else {
		switch string(_acad) {
		case "":
			*_cbec = 0
		case "\u0074":
			*_cbec = 1
		case "\u006d\u0069\u0064":
			*_cbec = 2
		case "\u0062":
			*_cbec = 3
		}
	}
	_cbdab, _fbgc = d.Token()
	if _fbgc != nil {
		return _fbgc
	}
	if _bfce, _bgaag := _cbdab.(_d.EndElement); _bgaag && _bfce.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cbdab)
}
func NewLayoutDefHdrLst() *LayoutDefHdrLst {
	_bdcb := &LayoutDefHdrLst{}
	_bdcb.CT_DiagramDefinitionHeaderLst = *NewCT_DiagramDefinitionHeaderLst()
	return _bdcb
}
func (_fgbf ST_HierarchyAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fgbf.String(), start)
}

const (
	ST_ChildDirectionUnset ST_ChildDirection = 0
	ST_ChildDirectionHorz  ST_ChildDirection = 1
	ST_ChildDirectionVert  ST_ChildDirection = 2
)

// ValidateWithPath validates the CT_Algorithm and its children, prefixing error messages with path
func (_dd *CT_Algorithm) ValidateWithPath(path string) error {
	if _dd.TypeAttr == ST_AlgorithmTypeUnset {
		return _c.Errorf("\u0025\u0073\u002f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020a\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _bbc := _dd.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _bbc != nil {
		return _bbc
	}
	for _cag, _aad := range _dd.Param {
		if _ddc := _aad.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fP\u0061\u0072\u0061\u006d\u005b\u0025\u0064\u005d", path, _cag)); _ddc != nil {
			return _ddc
		}
	}
	if _dd.ExtLst != nil {
		if _abd := _dd.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _abd != nil {
			return _abd
		}
	}
	return nil
}
func (_gabd *ST_ResizeHandlesStr) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ccbc, _agdc := d.Token()
	if _agdc != nil {
		return _agdc
	}
	if _ecgag, _fdcc := _ccbc.(_d.EndElement); _fdcc && _ecgag.Name == start.Name {
		*_gabd = 1
		return nil
	}
	if _adbfg, _befd := _ccbc.(_d.CharData); !_befd {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ccbc)
	} else {
		switch string(_adbfg) {
		case "":
			*_gabd = 0
		case "\u0065\u0078\u0061c\u0074":
			*_gabd = 1
		case "\u0072\u0065\u006c":
			*_gabd = 2
		}
	}
	_ccbc, _agdc = d.Token()
	if _agdc != nil {
		return _agdc
	}
	if _agade, _edeg := _ccbc.(_d.EndElement); _edeg && _agade.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ccbc)
}
func (_adbbd ST_GrowDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_adbbd.String(), start)
}
func ParseUnionST_FunctionArgument(s string) (ST_FunctionArgument, error) {
	return ST_FunctionArgument{}, nil
}
func (_fcdf ST_FunctionType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fcdf.String(), start)
}
func (_eabbe *ST_SecondaryLinearDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_eabbe = 0
	case "\u006e\u006f\u006e\u0065":
		*_eabbe = 1
	case "\u0066\u0072\u006fm\u004c":
		*_eabbe = 2
	case "\u0066\u0072\u006fm\u0052":
		*_eabbe = 3
	case "\u0066\u0072\u006fm\u0054":
		*_eabbe = 4
	case "\u0066\u0072\u006fm\u0042":
		*_eabbe = 5
	}
	return nil
}

type CT_SDCategory struct {
	TypeAttr string
	PriAttr  uint32
}

func (_adef *CT_SDCategory) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", _adef.TypeAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0069"}, Value: _c.Sprintf("\u0025\u0076", _adef.PriAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_gaag ST_ConnectorDimension) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ebbfc := _d.Attr{}
	_ebbfc.Name = name
	switch _gaag {
	case ST_ConnectorDimensionUnset:
		_ebbfc.Value = ""
	case ST_ConnectorDimension1D:
		_ebbfc.Value = "\u0031\u0044"
	case ST_ConnectorDimension2D:
		_ebbfc.Value = "\u0032\u0044"
	case ST_ConnectorDimensionCust:
		_ebbfc.Value = "\u0063\u0075\u0073\u0074"
	}
	return _ebbfc, nil
}

const (
	ST_TextBlockDirectionUnset ST_TextBlockDirection = 0
	ST_TextBlockDirectionHorz  ST_TextBlockDirection = 1
	ST_TextBlockDirectionVert  ST_TextBlockDirection = 2
)

// ValidateWithPath validates the RelIds and its children, prefixing error messages with path
func (_cddf *RelIds) ValidateWithPath(path string) error {
	if _fbfe := _cddf.CT_RelIds.ValidateWithPath(path); _fbfe != nil {
		return _fbfe
	}
	return nil
}
func NewCT_Otherwise() *CT_Otherwise { _cebd := &CT_Otherwise{}; return _cebd }

// Validate validates the CT_StyleDefinitionHeader and its children
func (_egaf *CT_StyleDefinitionHeader) Validate() error {
	return _egaf.ValidateWithPath("\u0043T\u005f\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006ei\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065\u0072")
}

type ST_OutputShapeType byte
type ST_CenterShapeMapping byte
type CT_SDName struct {
	LangAttr *string
	ValAttr  string
}
type CT_CTDescription struct {
	LangAttr *string
	ValAttr  string
}
type CT_StyleLabel struct {
	NameAttr string
	Scene3d  *_ca.CT_Scene3D
	Sp3d     *_ca.CT_Shape3D
	TxPr     *CT_TextProps
	Style    *_ca.CT_ShapeStyle
	ExtLst   *_ca.CT_OfficeArtExtensionList
}

// Validate validates the CT_StyleDefinition and its children
func (_gbdg *CT_StyleDefinition) Validate() error {
	return _gbdg.ValidateWithPath("\u0043T\u005fS\u0074\u0079\u006c\u0065\u0044e\u0066\u0069n\u0069\u0074\u0069\u006f\u006e")
}
func (_cgbg *StyleDefHdr) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cgbg.CT_StyleDefinitionHeader = *NewCT_StyleDefinitionHeader()
	for _, _fdab := range start.Attr {
		if _fdab.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_dbec, _bccgc := _fdab.Value, error(nil)
			if _bccgc != nil {
				return _bccgc
			}
			_cgbg.UniqueIdAttr = _dbec
			continue
		}
		if _fdab.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_ceaf, _gdaee := _fdab.Value, error(nil)
			if _gdaee != nil {
				return _gdaee
			}
			_cgbg.MinVerAttr = &_ceaf
			continue
		}
		if _fdab.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_daba, _fcabf := _ae.ParseInt(_fdab.Value, 10, 32)
			if _fcabf != nil {
				return _fcabf
			}
			_bccde := int32(_daba)
			_cgbg.ResIdAttr = &_bccde
			continue
		}
	}
_dcea:
	for {
		_gedb, _gcecb := d.Token()
		if _gcecb != nil {
			return _gcecb
		}
		switch _deaee := _gedb.(type) {
		case _d.StartElement:
			switch _deaee.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_dgdgf := NewCT_SDName()
				if _cdcec := d.DecodeElement(_dgdgf, &_deaee); _cdcec != nil {
					return _cdcec
				}
				_cgbg.Title = append(_cgbg.Title, _dgdgf)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_cdgga := NewCT_SDDescription()
				if _deab := d.DecodeElement(_cdgga, &_deaee); _deab != nil {
					return _deab
				}
				_cgbg.Desc = append(_cgbg.Desc, _cdgga)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_cgbg.CatLst = NewCT_SDCategories()
				if _eggeg := d.DecodeElement(_cgbg.CatLst, &_deaee); _eggeg != nil {
					return _eggeg
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_cgbg.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _beccb := d.DecodeElement(_cgbg.ExtLst, &_deaee); _beccb != nil {
					return _beccb
				}
			default:
				_af.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072\u0020\u0025\u0076", _deaee.Name)
				if _gaab := d.Skip(); _gaab != nil {
					return _gaab
				}
			}
		case _d.EndElement:
			break _dcea
		case _d.CharData:
		}
	}
	return nil
}

// Validate validates the LayoutDef and its children
func (_fdcb *LayoutDef) Validate() error {
	return _fdcb.ValidateWithPath("\u004ca\u0079\u006f\u0075\u0074\u0044\u0065f")
}
func NewCT_AnimLvl() *CT_AnimLvl { _afg := &CT_AnimLvl{}; return _afg }

// ValidateWithPath validates the CT_Adj and its children, prefixing error messages with path
func (_dg *CT_Adj) ValidateWithPath(path string) error {
	if _dg.IdxAttr < 1 {
		return _c.Errorf("%\u0073\u002f\u006d\u002e\u0049\u0064x\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0031\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _dg.IdxAttr)
	}
	return nil
}

// Validate validates the CT_Cxn and its children
func (_gege *CT_Cxn) Validate() error {
	return _gege.ValidateWithPath("\u0043\u0054\u005f\u0043\u0078\u006e")
}
func (_eabag ST_ResizeHandlesStr) String() string {
	switch _eabag {
	case 0:
		return ""
	case 1:
		return "\u0065\u0078\u0061c\u0074"
	case 2:
		return "\u0072\u0065\u006c"
	}
	return ""
}
func (_feccf *CT_TextProps) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _feccf.Sp3d != nil {
		_gaeea := _d.StartElement{Name: _d.Name{Local: "\u0073\u0070\u0033\u0064"}}
		e.EncodeElement(_feccf.Sp3d, _gaeea)
	}
	if _feccf.FlatTx != nil {
		_becd := _d.StartElement{Name: _d.Name{Local: "\u0066\u006c\u0061\u0074\u0054\u0078"}}
		e.EncodeElement(_feccf.FlatTx, _becd)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_cgdcg ST_FlowDirection) String() string {
	switch _cgdcg {
	case 0:
		return ""
	case 1:
		return "\u0072\u006f\u0077"
	case 2:
		return "\u0063\u006f\u006c"
	}
	return ""
}
func (_eddfe ST_FallbackDimension) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_eddfe.String(), start)
}
func (_cedfa ST_LinearDirection) Validate() error { return _cedfa.ValidateWithPath("") }
func (_gefg ST_AlgorithmType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_becce := _d.Attr{}
	_becce.Name = name
	switch _gefg {
	case ST_AlgorithmTypeUnset:
		_becce.Value = ""
	case ST_AlgorithmTypeComposite:
		_becce.Value = "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e"
	case ST_AlgorithmTypeConn:
		_becce.Value = "\u0063\u006f\u006e\u006e"
	case ST_AlgorithmTypeCycle:
		_becce.Value = "\u0063\u0079\u0063l\u0065"
	case ST_AlgorithmTypeHierChild:
		_becce.Value = "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd"
	case ST_AlgorithmTypeHierRoot:
		_becce.Value = "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074"
	case ST_AlgorithmTypePyra:
		_becce.Value = "\u0070\u0079\u0072\u0061"
	case ST_AlgorithmTypeLin:
		_becce.Value = "\u006c\u0069\u006e"
	case ST_AlgorithmTypeSp:
		_becce.Value = "\u0073\u0070"
	case ST_AlgorithmTypeTx:
		_becce.Value = "\u0074\u0078"
	case ST_AlgorithmTypeSnake:
		_becce.Value = "\u0073\u006e\u0061k\u0065"
	}
	return _becce, nil
}
func (_efcf *ST_ConstraintType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_daad, _gcfg := d.Token()
	if _gcfg != nil {
		return _gcfg
	}
	if _cdgaf, _dccce := _daad.(_d.EndElement); _dccce && _cdgaf.Name == start.Name {
		*_efcf = 1
		return nil
	}
	if _cggf, _fgbe := _daad.(_d.CharData); !_fgbe {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _daad)
	} else {
		switch string(_cggf) {
		case "":
			*_efcf = 0
		case "\u006e\u006f\u006e\u0065":
			*_efcf = 1
		case "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066":
			*_efcf = 2
		case "\u0062e\u0067\u004d\u0061\u0072\u0067":
			*_efcf = 3
		case "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074":
			*_efcf = 4
		case "\u0062\u0065\u0067\u0050\u0061\u0064":
			*_efcf = 5
		case "\u0062":
			*_efcf = 6
		case "\u0062\u004d\u0061r\u0067":
			*_efcf = 7
		case "\u0062\u004f\u0066\u0066":
			*_efcf = 8
		case "\u0063\u0074\u0072\u0058":
			*_efcf = 9
		case "\u0063t\u0072\u0058\u004f\u0066\u0066":
			*_efcf = 10
		case "\u0063\u0074\u0072\u0059":
			*_efcf = 11
		case "\u0063t\u0072\u0059\u004f\u0066\u0066":
			*_efcf = 12
		case "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074":
			*_efcf = 13
		case "\u0064\u0069\u0061\u006d":
			*_efcf = 14
		case "\u0065n\u0064\u004d\u0061\u0072\u0067":
			*_efcf = 15
		case "\u0065\u006e\u0064\u0050\u0061\u0064":
			*_efcf = 16
		case "\u0068":
			*_efcf = 17
		case "\u0068\u0041\u0072\u0048":
			*_efcf = 18
		case "\u0068\u004f\u0066\u0066":
			*_efcf = 19
		case "\u006c":
			*_efcf = 20
		case "\u006c\u004d\u0061r\u0067":
			*_efcf = 21
		case "\u006c\u004f\u0066\u0066":
			*_efcf = 22
		case "\u0072":
			*_efcf = 23
		case "\u0072\u004d\u0061r\u0067":
			*_efcf = 24
		case "\u0072\u004f\u0066\u0066":
			*_efcf = 25
		case "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a":
			*_efcf = 26
		case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f":
			*_efcf = 27
		case "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z":
			*_efcf = 28
		case "\u0073\u0069\u0062S\u0070":
			*_efcf = 29
		case "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070":
			*_efcf = 30
		case "\u0073\u0070":
			*_efcf = 31
		case "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k":
			*_efcf = 32
		case "\u0074":
			*_efcf = 33
		case "\u0074\u004d\u0061r\u0067":
			*_efcf = 34
		case "\u0074\u004f\u0066\u0066":
			*_efcf = 35
		case "\u0075\u0073\u0065r\u0041":
			*_efcf = 36
		case "\u0075\u0073\u0065r\u0042":
			*_efcf = 37
		case "\u0075\u0073\u0065r\u0043":
			*_efcf = 38
		case "\u0075\u0073\u0065r\u0044":
			*_efcf = 39
		case "\u0075\u0073\u0065r\u0045":
			*_efcf = 40
		case "\u0075\u0073\u0065r\u0046":
			*_efcf = 41
		case "\u0075\u0073\u0065r\u0047":
			*_efcf = 42
		case "\u0075\u0073\u0065r\u0048":
			*_efcf = 43
		case "\u0075\u0073\u0065r\u0049":
			*_efcf = 44
		case "\u0075\u0073\u0065r\u004a":
			*_efcf = 45
		case "\u0075\u0073\u0065r\u004b":
			*_efcf = 46
		case "\u0075\u0073\u0065r\u004c":
			*_efcf = 47
		case "\u0075\u0073\u0065r\u004d":
			*_efcf = 48
		case "\u0075\u0073\u0065r\u004e":
			*_efcf = 49
		case "\u0075\u0073\u0065r\u004f":
			*_efcf = 50
		case "\u0075\u0073\u0065r\u0050":
			*_efcf = 51
		case "\u0075\u0073\u0065r\u0051":
			*_efcf = 52
		case "\u0075\u0073\u0065r\u0052":
			*_efcf = 53
		case "\u0075\u0073\u0065r\u0053":
			*_efcf = 54
		case "\u0075\u0073\u0065r\u0054":
			*_efcf = 55
		case "\u0075\u0073\u0065r\u0055":
			*_efcf = 56
		case "\u0075\u0073\u0065r\u0056":
			*_efcf = 57
		case "\u0075\u0073\u0065r\u0057":
			*_efcf = 58
		case "\u0075\u0073\u0065r\u0058":
			*_efcf = 59
		case "\u0075\u0073\u0065r\u0059":
			*_efcf = 60
		case "\u0075\u0073\u0065r\u005a":
			*_efcf = 61
		case "\u0077":
			*_efcf = 62
		case "\u0077\u0041\u0072\u0048":
			*_efcf = 63
		case "\u0077\u004f\u0066\u0066":
			*_efcf = 64
		}
	}
	_daad, _gcfg = d.Token()
	if _gcfg != nil {
		return _gcfg
	}
	if _fdbf, _fgfe := _daad.(_d.EndElement); _fgfe && _fdbf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _daad)
}

type ST_ConstraintType byte

// Validate validates the CT_Categories and its children
func (_dgfb *CT_Categories) Validate() error {
	return _dgfb.ValidateWithPath("\u0043\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073")
}
func (_bgdg *ST_FallbackDimension) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bgdg = 0
	case "\u0031\u0044":
		*_bgdg = 1
	case "\u0032\u0044":
		*_bgdg = 2
	}
	return nil
}

type ST_ElementType byte

func (_eaba *RelIds) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0072\u0065\u006c\u0049\u0064\u0073"
	return _eaba.CT_RelIds.MarshalXML(e, start)
}
func (_bbba ST_ConnectorPoint) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_bbba.String(), start)
}
func (_eaec *ST_LayoutShapeType) Validate() error { return _eaec.ValidateWithPath("") }

const (
	ST_ChildOrderTypeUnset ST_ChildOrderType = 0
	ST_ChildOrderTypeB     ST_ChildOrderType = 1
	ST_ChildOrderTypeT     ST_ChildOrderType = 2
)

func (_daff *CT_LayoutVariablePropertySet) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _daff.OrgChart != nil {
		_bebcf := _d.StartElement{Name: _d.Name{Local: "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"}}
		e.EncodeElement(_daff.OrgChart, _bebcf)
	}
	if _daff.ChMax != nil {
		_aafa := _d.StartElement{Name: _d.Name{Local: "\u0063\u0068\u004da\u0078"}}
		e.EncodeElement(_daff.ChMax, _aafa)
	}
	if _daff.ChPref != nil {
		_eabb := _d.StartElement{Name: _d.Name{Local: "\u0063\u0068\u0050\u0072\u0065\u0066"}}
		e.EncodeElement(_daff.ChPref, _eabb)
	}
	if _daff.BulletEnabled != nil {
		_bgbe := _d.StartElement{Name: _d.Name{Local: "\u0062\u0075\u006c\u006c\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064"}}
		e.EncodeElement(_daff.BulletEnabled, _bgbe)
	}
	if _daff.Dir != nil {
		_bdba := _d.StartElement{Name: _d.Name{Local: "\u0064\u0069\u0072"}}
		e.EncodeElement(_daff.Dir, _bdba)
	}
	if _daff.HierBranch != nil {
		_cfdce := _d.StartElement{Name: _d.Name{Local: "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"}}
		e.EncodeElement(_daff.HierBranch, _cfdce)
	}
	if _daff.AnimOne != nil {
		_aege := _d.StartElement{Name: _d.Name{Local: "\u0061n\u0069\u006d\u004f\u006e\u0065"}}
		e.EncodeElement(_daff.AnimOne, _aege)
	}
	if _daff.AnimLvl != nil {
		_gaae := _d.StartElement{Name: _d.Name{Local: "\u0061n\u0069\u006d\u004c\u0076\u006c"}}
		e.EncodeElement(_daff.AnimLvl, _gaae)
	}
	if _daff.ResizeHandles != nil {
		_deg := _d.StartElement{Name: _d.Name{Local: "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"}}
		e.EncodeElement(_daff.ResizeHandles, _deg)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_cfdcb ST_ElementType) String() string {
	switch _cfdcb {
	case 0:
		return ""
	case 1:
		return "\u0061\u006c\u006c"
	case 2:
		return "\u0064\u006f\u0063"
	case 3:
		return "\u006e\u006f\u0064\u0065"
	case 4:
		return "\u006e\u006f\u0072\u006d"
	case 5:
		return "\u006eo\u006e\u004e\u006f\u0072\u006d"
	case 6:
		return "\u0061\u0073\u0073\u0074"
	case 7:
		return "\u006eo\u006e\u0041\u0073\u0073\u0074"
	case 8:
		return "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case 9:
		return "\u0070\u0072\u0065\u0073"
	case 10:
		return "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return ""
}
func NewCT_When() *CT_When {
	_gfga := &CT_When{}
	_gfga.FuncAttr = ST_FunctionType(1)
	_gfga.OpAttr = ST_FunctionOperator(1)
	return _gfga
}
func (_cgaba *ST_TextDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_eccaa, _efgee := d.Token()
	if _efgee != nil {
		return _efgee
	}
	if _decdb, _fdag := _eccaa.(_d.EndElement); _fdag && _decdb.Name == start.Name {
		*_cgaba = 1
		return nil
	}
	if _aaegc, _ecdge := _eccaa.(_d.CharData); !_ecdge {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eccaa)
	} else {
		switch string(_aaegc) {
		case "":
			*_cgaba = 0
		case "\u0066\u0072\u006fm\u0054":
			*_cgaba = 1
		case "\u0066\u0072\u006fm\u0042":
			*_cgaba = 2
		}
	}
	_eccaa, _efgee = d.Token()
	if _efgee != nil {
		return _efgee
	}
	if _cebg, _becfd := _eccaa.(_d.EndElement); _becfd && _cebg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eccaa)
}

// Validate validates the ColorsDefHdrLst and its children
func (_decd *ColorsDefHdrLst) Validate() error {
	return _decd.ValidateWithPath("\u0043o\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074")
}
func (_dedb ST_TextAnchorVertical) Validate() error { return _dedb.ValidateWithPath("") }
func NewCT_ChildPref() *CT_ChildPref                { _ggcb := &CT_ChildPref{}; return _ggcb }
func NewCT_StyleDefinitionHeaderLst() *CT_StyleDefinitionHeaderLst {
	_fgcee := &CT_StyleDefinitionHeaderLst{}
	return _fgcee
}

type ST_SecondaryLinearDirection byte

func (_cffdf *StyleDefHdr) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"
	return _cffdf.CT_StyleDefinitionHeader.MarshalXML(e, start)
}

type ST_AxisTypes []ST_AxisType

func NewCT_StyleLabel() *CT_StyleLabel { _afef := &CT_StyleLabel{}; return _afef }

type CT_Constraints struct{ Constr []*CT_Constraint }

// ValidateWithPath validates the CT_RelIds and its children, prefixing error messages with path
func (_dbbb *CT_RelIds) ValidateWithPath(path string) error { return nil }

type ST_HierBranchStyle byte

func (_gdcd ST_ArrowheadStyle) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_bdaff := _d.Attr{}
	_bdaff.Name = name
	switch _gdcd {
	case ST_ArrowheadStyleUnset:
		_bdaff.Value = ""
	case ST_ArrowheadStyleAuto:
		_bdaff.Value = "\u0061\u0075\u0074\u006f"
	case ST_ArrowheadStyleArr:
		_bdaff.Value = "\u0061\u0072\u0072"
	case ST_ArrowheadStyleNoArr:
		_bdaff.Value = "\u006e\u006f\u0041r\u0072"
	}
	return _bdaff, nil
}
func NewCT_RelIds() *CT_RelIds { _bgcf := &CT_RelIds{}; return _bgcf }

type ST_AutoTextRotation byte

func (_dcbg *CT_DiagramDefinitionHeader) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _egde := range start.Attr {
		if _egde.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_eagc, _beea := _egde.Value, error(nil)
			if _beea != nil {
				return _beea
			}
			_dcbg.UniqueIdAttr = _eagc
			continue
		}
		if _egde.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_bfbb, _bbfd := _egde.Value, error(nil)
			if _bbfd != nil {
				return _bbfd
			}
			_dcbg.MinVerAttr = &_bfbb
			continue
		}
		if _egde.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_eceag, _fedc := _egde.Value, error(nil)
			if _fedc != nil {
				return _fedc
			}
			_dcbg.DefStyleAttr = &_eceag
			continue
		}
		if _egde.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_cbbf, _befc := _ae.ParseInt(_egde.Value, 10, 32)
			if _befc != nil {
				return _befc
			}
			_agfac := int32(_cbbf)
			_dcbg.ResIdAttr = &_agfac
			continue
		}
	}
_faeg:
	for {
		_bgbde, _fecc := d.Token()
		if _fecc != nil {
			return _fecc
		}
		switch _gbfeb := _bgbde.(type) {
		case _d.StartElement:
			switch _gbfeb.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_cafe := NewCT_Name()
				if _ggga := d.DecodeElement(_cafe, &_gbfeb); _ggga != nil {
					return _ggga
				}
				_dcbg.Title = append(_dcbg.Title, _cafe)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_fbc := NewCT_Description()
				if _gae := d.DecodeElement(_fbc, &_gbfeb); _gae != nil {
					return _gae
				}
				_dcbg.Desc = append(_dcbg.Desc, _fbc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_dcbg.CatLst = NewCT_Categories()
				if _gcea := d.DecodeElement(_dcbg.CatLst, &_gbfeb); _gcea != nil {
					return _gcea
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_dcbg.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _aaaa := d.DecodeElement(_dcbg.ExtLst, &_gbfeb); _aaaa != nil {
					return _aaaa
				}
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020o\u006e\u0020\u0043\u0054_\u0044\u0069a\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065\u0072\u0020\u0025\u0076", _gbfeb.Name)
				if _ecfd := d.Skip(); _ecfd != nil {
					return _ecfd
				}
			}
		case _d.EndElement:
			break _faeg
		case _d.CharData:
		}
	}
	return nil
}
func (_fgdcb *ST_ElementType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_efgb, _abfe := d.Token()
	if _abfe != nil {
		return _abfe
	}
	if _gcff, _agada := _efgb.(_d.EndElement); _agada && _gcff.Name == start.Name {
		*_fgdcb = 1
		return nil
	}
	if _baeeg, _ggcac := _efgb.(_d.CharData); !_ggcac {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _efgb)
	} else {
		switch string(_baeeg) {
		case "":
			*_fgdcb = 0
		case "\u0061\u006c\u006c":
			*_fgdcb = 1
		case "\u0064\u006f\u0063":
			*_fgdcb = 2
		case "\u006e\u006f\u0064\u0065":
			*_fgdcb = 3
		case "\u006e\u006f\u0072\u006d":
			*_fgdcb = 4
		case "\u006eo\u006e\u004e\u006f\u0072\u006d":
			*_fgdcb = 5
		case "\u0061\u0073\u0073\u0074":
			*_fgdcb = 6
		case "\u006eo\u006e\u0041\u0073\u0073\u0074":
			*_fgdcb = 7
		case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
			*_fgdcb = 8
		case "\u0070\u0072\u0065\u0073":
			*_fgdcb = 9
		case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
			*_fgdcb = 10
		}
	}
	_efgb, _abfe = d.Token()
	if _abfe != nil {
		return _abfe
	}
	if _effbf, _gffc := _efgb.(_d.EndElement); _gffc && _effbf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _efgb)
}
func (_fdbe *ST_SecondaryChildAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fdbe = 0
	case "\u006e\u006f\u006e\u0065":
		*_fdbe = 1
	case "\u0074":
		*_fdbe = 2
	case "\u0062":
		*_fdbe = 3
	case "\u006c":
		*_fdbe = 4
	case "\u0072":
		*_fdbe = 5
	}
	return nil
}

// ValidateWithPath validates the CT_Categories and its children, prefixing error messages with path
func (_bafa *CT_Categories) ValidateWithPath(path string) error {
	for _fac, _bff := range _bafa.Cat {
		if _egec := _bff.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0061\u0074\u005b\u0025\u0064\u005d", path, _fac)); _egec != nil {
			return _egec
		}
	}
	return nil
}
func (_gfdcg ST_Direction) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gdac := _d.Attr{}
	_gdac.Name = name
	switch _gfdcg {
	case ST_DirectionUnset:
		_gdac.Value = ""
	case ST_DirectionNorm:
		_gdac.Value = "\u006e\u006f\u0072\u006d"
	case ST_DirectionRev:
		_gdac.Value = "\u0072\u0065\u0076"
	}
	return _gdac, nil
}
func (_abadg ST_ChildOrderType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_abadg.String(), start)
}
func (_dgg *CT_Category) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _cba := range start.Attr {
		if _cba.Name.Local == "\u0074\u0079\u0070\u0065" {
			_gdgc, _edb := _cba.Value, error(nil)
			if _edb != nil {
				return _edb
			}
			_dgg.TypeAttr = _gdgc
			continue
		}
		if _cba.Name.Local == "\u0070\u0072\u0069" {
			_dggd, _acdd := _ae.ParseUint(_cba.Value, 10, 32)
			if _acdd != nil {
				return _acdd
			}
			_dgg.PriAttr = uint32(_dggd)
			continue
		}
	}
	for {
		_gdga, _gca := d.Token()
		if _gca != nil {
			return _c.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079\u003a\u0020\u0025\u0073", _gca)
		}
		if _afdg, _gbg := _gdga.(_d.EndElement); _gbg && _afdg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_caeea ST_ConnectorRouting) Validate() error { return _caeea.ValidateWithPath("") }
func (_bddb ST_TextDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_egae := _d.Attr{}
	_egae.Name = name
	switch _bddb {
	case ST_TextDirectionUnset:
		_egae.Value = ""
	case ST_TextDirectionFromT:
		_egae.Value = "\u0066\u0072\u006fm\u0054"
	case ST_TextDirectionFromB:
		_egae.Value = "\u0066\u0072\u006fm\u0042"
	}
	return _egae, nil
}

// Validate validates the CT_Otherwise and its children
func (_acbd *CT_Otherwise) Validate() error {
	return _acbd.ValidateWithPath("\u0043\u0054\u005fO\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065")
}

// Validate validates the DataModel and its children
func (_fffd *DataModel) Validate() error {
	return _fffd.ValidateWithPath("\u0044a\u0074\u0061\u004d\u006f\u0064\u0065l")
}

// ValidateWithPath validates the CT_ResizeHandles and its children, prefixing error messages with path
func (_cdce *CT_ResizeHandles) ValidateWithPath(path string) error {
	if _gdege := _cdce.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gdege != nil {
		return _gdege
	}
	return nil
}
func _efffd(_dccf bool) uint8 {
	if _dccf {
		return 1
	}
	return 0
}
func (_eaecb ST_ArrowheadStyle) ValidateWithPath(path string) error {
	switch _eaecb {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eaecb))
	}
	return nil
}
func (_abcdd ST_PtType) Validate() error { return _abcdd.ValidateWithPath("") }

// ValidateWithPath validates the CT_AdjLst and its children, prefixing error messages with path
func (_da *CT_AdjLst) ValidateWithPath(path string) error {
	for _cdg, _fab := range _da.Adj {
		if _aff := _fab.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0041\u0064\u006a\u005b\u0025\u0064\u005d", path, _cdg)); _aff != nil {
			return _aff
		}
	}
	return nil
}
func NewColorsDefHdrLst() *ColorsDefHdrLst {
	_fbda := &ColorsDefHdrLst{}
	_fbda.CT_ColorTransformHeaderLst = *NewCT_ColorTransformHeaderLst()
	return _fbda
}

const (
	ST_ParameterIdUnset            ST_ParameterId = 0
	ST_ParameterIdHorzAlign        ST_ParameterId = 1
	ST_ParameterIdVertAlign        ST_ParameterId = 2
	ST_ParameterIdChDir            ST_ParameterId = 3
	ST_ParameterIdChAlign          ST_ParameterId = 4
	ST_ParameterIdSecChAlign       ST_ParameterId = 5
	ST_ParameterIdLinDir           ST_ParameterId = 6
	ST_ParameterIdSecLinDir        ST_ParameterId = 7
	ST_ParameterIdStElem           ST_ParameterId = 8
	ST_ParameterIdBendPt           ST_ParameterId = 9
	ST_ParameterIdConnRout         ST_ParameterId = 10
	ST_ParameterIdBegSty           ST_ParameterId = 11
	ST_ParameterIdEndSty           ST_ParameterId = 12
	ST_ParameterIdDim              ST_ParameterId = 13
	ST_ParameterIdRotPath          ST_ParameterId = 14
	ST_ParameterIdCtrShpMap        ST_ParameterId = 15
	ST_ParameterIdNodeHorzAlign    ST_ParameterId = 16
	ST_ParameterIdNodeVertAlign    ST_ParameterId = 17
	ST_ParameterIdFallback         ST_ParameterId = 18
	ST_ParameterIdTxDir            ST_ParameterId = 19
	ST_ParameterIdPyraAcctPos      ST_ParameterId = 20
	ST_ParameterIdPyraAcctTxMar    ST_ParameterId = 21
	ST_ParameterIdTxBlDir          ST_ParameterId = 22
	ST_ParameterIdTxAnchorHorz     ST_ParameterId = 23
	ST_ParameterIdTxAnchorVert     ST_ParameterId = 24
	ST_ParameterIdTxAnchorHorzCh   ST_ParameterId = 25
	ST_ParameterIdTxAnchorVertCh   ST_ParameterId = 26
	ST_ParameterIdParTxLTRAlign    ST_ParameterId = 27
	ST_ParameterIdParTxRTLAlign    ST_ParameterId = 28
	ST_ParameterIdShpTxLTRAlignCh  ST_ParameterId = 29
	ST_ParameterIdShpTxRTLAlignCh  ST_ParameterId = 30
	ST_ParameterIdAutoTxRot        ST_ParameterId = 31
	ST_ParameterIdGrDir            ST_ParameterId = 32
	ST_ParameterIdFlowDir          ST_ParameterId = 33
	ST_ParameterIdContDir          ST_ParameterId = 34
	ST_ParameterIdBkpt             ST_ParameterId = 35
	ST_ParameterIdOff              ST_ParameterId = 36
	ST_ParameterIdHierAlign        ST_ParameterId = 37
	ST_ParameterIdBkPtFixedVal     ST_ParameterId = 38
	ST_ParameterIdStBulletLvl      ST_ParameterId = 39
	ST_ParameterIdStAng            ST_ParameterId = 40
	ST_ParameterIdSpanAng          ST_ParameterId = 41
	ST_ParameterIdAr               ST_ParameterId = 42
	ST_ParameterIdLnSpPar          ST_ParameterId = 43
	ST_ParameterIdLnSpAfParP       ST_ParameterId = 44
	ST_ParameterIdLnSpCh           ST_ParameterId = 45
	ST_ParameterIdLnSpAfChP        ST_ParameterId = 46
	ST_ParameterIdRtShortDist      ST_ParameterId = 47
	ST_ParameterIdAlignTx          ST_ParameterId = 48
	ST_ParameterIdPyraLvlNode      ST_ParameterId = 49
	ST_ParameterIdPyraAcctBkgdNode ST_ParameterId = 50
	ST_ParameterIdPyraAcctTxNode   ST_ParameterId = 51
	ST_ParameterIdSrcNode          ST_ParameterId = 52
	ST_ParameterIdDstNode          ST_ParameterId = 53
	ST_ParameterIdBegPts           ST_ParameterId = 54
	ST_ParameterIdEndPts           ST_ParameterId = 55
)

func ParseUnionST_FunctionValue(s string) (ST_FunctionValue, error) { return ST_FunctionValue{}, nil }
func (_eedd *CT_Parameter) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	_gfeec, _fffc := _eedd.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
	if _fffc != nil {
		return _fffc
	}
	start.Attr = append(start.Attr, _gfeec)
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _eedd.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_cegd *ST_FlowDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_cegd = 0
	case "\u0072\u006f\u0077":
		*_cegd = 1
	case "\u0063\u006f\u006c":
		*_cegd = 2
	}
	return nil
}

// Validate validates the AG_ConstraintAttributes and its children
func (_ge *AG_ConstraintAttributes) Validate() error {
	return _ge.ValidateWithPath("\u0041\u0047\u005fCo\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074\u0041\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0073")
}

type CT_Category struct {
	TypeAttr string
	PriAttr  uint32
}

func NewCT_PtList() *CT_PtList { _bfeb := &CT_PtList{}; return _bfeb }
func (_fgbba ST_Offset) String() string {
	switch _fgbba {
	case 0:
		return ""
	case 1:
		return "\u0063\u0074\u0072"
	case 2:
		return "\u006f\u0066\u0066"
	}
	return ""
}

type ST_StartingElement byte

func (_ceba ST_HueDir) ValidateWithPath(path string) error {
	switch _ceba {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ceba))
	}
	return nil
}
func (_gffd ST_Direction) Validate() error { return _gffd.ValidateWithPath("") }
func (_ggbd ST_SecondaryChildAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ggbd.String(), start)
}

const (
	ST_FlowDirectionUnset ST_FlowDirection = 0
	ST_FlowDirectionRow   ST_FlowDirection = 1
	ST_FlowDirectionCol   ST_FlowDirection = 2
)

// Validate validates the CT_Direction and its children
func (_fagd *CT_Direction) Validate() error {
	return _fagd.ValidateWithPath("\u0043\u0054\u005fD\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
}
func (_ebdg ST_AutoTextRotation) String() string {
	switch _ebdg {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0075\u0070\u0072"
	case 3:
		return "\u0067\u0072\u0061\u0076"
	}
	return ""
}
func ParseSliceST_ElementTypes(s string) (ST_ElementTypes, error) { return ST_ElementTypes{}, nil }

type CT_StyleDefinitionHeaderLst struct{ StyleDefHdr []*CT_StyleDefinitionHeader }

// ValidateWithPath validates the CT_ColorTransform and its children, prefixing error messages with path
func (_fdc *CT_ColorTransform) ValidateWithPath(path string) error {
	for _bcf, _ddgb := range _fdc.Title {
		if _dgbg := _ddgb.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _bcf)); _dgbg != nil {
			return _dgbg
		}
	}
	for _gbcfg, _gfeg := range _fdc.Desc {
		if _eaf := _gfeg.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _gbcfg)); _eaf != nil {
			return _eaf
		}
	}
	if _fdc.CatLst != nil {
		if _daee := _fdc.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _daee != nil {
			return _daee
		}
	}
	for _egea, _eddd := range _fdc.StyleLbl {
		if _adgf := _eddd.ValidateWithPath(_c.Sprintf("\u0025s\u002fS\u0074\u0079\u006c\u0065\u004c\u0062\u006c\u005b\u0025\u0064\u005d", path, _egea)); _adgf != nil {
			return _adgf
		}
	}
	if _fdc.ExtLst != nil {
		if _dga := _fdc.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _dga != nil {
			return _dga
		}
	}
	return nil
}
func (_aggfe ST_ParameterId) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_fgcce := _d.Attr{}
	_fgcce.Name = name
	switch _aggfe {
	case ST_ParameterIdUnset:
		_fgcce.Value = ""
	case ST_ParameterIdHorzAlign:
		_fgcce.Value = "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n"
	case ST_ParameterIdVertAlign:
		_fgcce.Value = "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n"
	case ST_ParameterIdChDir:
		_fgcce.Value = "\u0063\u0068\u0044i\u0072"
	case ST_ParameterIdChAlign:
		_fgcce.Value = "\u0063h\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdSecChAlign:
		_fgcce.Value = "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdLinDir:
		_fgcce.Value = "\u006c\u0069\u006e\u0044\u0069\u0072"
	case ST_ParameterIdSecLinDir:
		_fgcce.Value = "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r"
	case ST_ParameterIdStElem:
		_fgcce.Value = "\u0073\u0074\u0045\u006c\u0065\u006d"
	case ST_ParameterIdBendPt:
		_fgcce.Value = "\u0062\u0065\u006e\u0064\u0050\u0074"
	case ST_ParameterIdConnRout:
		_fgcce.Value = "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074"
	case ST_ParameterIdBegSty:
		_fgcce.Value = "\u0062\u0065\u0067\u0053\u0074\u0079"
	case ST_ParameterIdEndSty:
		_fgcce.Value = "\u0065\u006e\u0064\u0053\u0074\u0079"
	case ST_ParameterIdDim:
		_fgcce.Value = "\u0064\u0069\u006d"
	case ST_ParameterIdRotPath:
		_fgcce.Value = "\u0072o\u0074\u0050\u0061\u0074\u0068"
	case ST_ParameterIdCtrShpMap:
		_fgcce.Value = "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p"
	case ST_ParameterIdNodeHorzAlign:
		_fgcce.Value = "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdNodeVertAlign:
		_fgcce.Value = "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdFallback:
		_fgcce.Value = "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b"
	case ST_ParameterIdTxDir:
		_fgcce.Value = "\u0074\u0078\u0044i\u0072"
	case ST_ParameterIdPyraAcctPos:
		_fgcce.Value = "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073"
	case ST_ParameterIdPyraAcctTxMar:
		_fgcce.Value = "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072"
	case ST_ParameterIdTxBlDir:
		_fgcce.Value = "\u0074x\u0042\u006c\u0044\u0069\u0072"
	case ST_ParameterIdTxAnchorHorz:
		_fgcce.Value = "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a"
	case ST_ParameterIdTxAnchorVert:
		_fgcce.Value = "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074"
	case ST_ParameterIdTxAnchorHorzCh:
		_fgcce.Value = "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068"
	case ST_ParameterIdTxAnchorVertCh:
		_fgcce.Value = "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068"
	case ST_ParameterIdParTxLTRAlign:
		_fgcce.Value = "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdParTxRTLAlign:
		_fgcce.Value = "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdShpTxLTRAlignCh:
		_fgcce.Value = "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case ST_ParameterIdShpTxRTLAlignCh:
		_fgcce.Value = "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case ST_ParameterIdAutoTxRot:
		_fgcce.Value = "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft"
	case ST_ParameterIdGrDir:
		_fgcce.Value = "\u0067\u0072\u0044i\u0072"
	case ST_ParameterIdFlowDir:
		_fgcce.Value = "\u0066l\u006f\u0077\u0044\u0069\u0072"
	case ST_ParameterIdContDir:
		_fgcce.Value = "\u0063o\u006e\u0074\u0044\u0069\u0072"
	case ST_ParameterIdBkpt:
		_fgcce.Value = "\u0062\u006b\u0070\u0074"
	case ST_ParameterIdOff:
		_fgcce.Value = "\u006f\u0066\u0066"
	case ST_ParameterIdHierAlign:
		_fgcce.Value = "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n"
	case ST_ParameterIdBkPtFixedVal:
		_fgcce.Value = "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c"
	case ST_ParameterIdStBulletLvl:
		_fgcce.Value = "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c"
	case ST_ParameterIdStAng:
		_fgcce.Value = "\u0073\u0074\u0041n\u0067"
	case ST_ParameterIdSpanAng:
		_fgcce.Value = "\u0073p\u0061\u006e\u0041\u006e\u0067"
	case ST_ParameterIdAr:
		_fgcce.Value = "\u0061\u0072"
	case ST_ParameterIdLnSpPar:
		_fgcce.Value = "\u006cn\u0053\u0070\u0050\u0061\u0072"
	case ST_ParameterIdLnSpAfParP:
		_fgcce.Value = "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050"
	case ST_ParameterIdLnSpCh:
		_fgcce.Value = "\u006c\u006e\u0053\u0070\u0043\u0068"
	case ST_ParameterIdLnSpAfChP:
		_fgcce.Value = "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P"
	case ST_ParameterIdRtShortDist:
		_fgcce.Value = "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074"
	case ST_ParameterIdAlignTx:
		_fgcce.Value = "\u0061l\u0069\u0067\u006e\u0054\u0078"
	case ST_ParameterIdPyraLvlNode:
		_fgcce.Value = "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065"
	case ST_ParameterIdPyraAcctBkgdNode:
		_fgcce.Value = "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065"
	case ST_ParameterIdPyraAcctTxNode:
		_fgcce.Value = "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065"
	case ST_ParameterIdSrcNode:
		_fgcce.Value = "\u0073r\u0063\u004e\u006f\u0064\u0065"
	case ST_ParameterIdDstNode:
		_fgcce.Value = "\u0064s\u0074\u004e\u006f\u0064\u0065"
	case ST_ParameterIdBegPts:
		_fgcce.Value = "\u0062\u0065\u0067\u0050\u0074\u0073"
	case ST_ParameterIdEndPts:
		_fgcce.Value = "\u0065\u006e\u0064\u0050\u0074\u0073"
	}
	return _fgcce, nil
}

type ST_Offset byte

// Validate validates the LayoutDefHdrLst and its children
func (_eabec *LayoutDefHdrLst) Validate() error {
	return _eabec.ValidateWithPath("\u004ca\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074")
}
func NewCT_Parameter() *CT_Parameter {
	_eced := &CT_Parameter{}
	_eced.TypeAttr = ST_ParameterId(1)
	return _eced
}
func NewLayoutDef() *LayoutDef {
	_dggf := &LayoutDef{}
	_dggf.CT_DiagramDefinition = *NewCT_DiagramDefinition()
	return _dggf
}
func (_dcefdf *ST_AnimLvlStr) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_dcefdf = 0
	case "\u006e\u006f\u006e\u0065":
		*_dcefdf = 1
	case "\u006c\u0076\u006c":
		*_dcefdf = 2
	case "\u0063\u0074\u0072":
		*_dcefdf = 3
	}
	return nil
}
func (_bgfb *LayoutDef) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bgfb.CT_DiagramDefinition = *NewCT_DiagramDefinition()
	for _, _egag := range start.Attr {
		if _egag.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_fcge, _adefd := _egag.Value, error(nil)
			if _adefd != nil {
				return _adefd
			}
			_bgfb.UniqueIdAttr = &_fcge
			continue
		}
		if _egag.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_dcdca, _eebbd := _egag.Value, error(nil)
			if _eebbd != nil {
				return _eebbd
			}
			_bgfb.MinVerAttr = &_dcdca
			continue
		}
		if _egag.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_eafd, _eacba := _egag.Value, error(nil)
			if _eacba != nil {
				return _eacba
			}
			_bgfb.DefStyleAttr = &_eafd
			continue
		}
	}
_aecff:
	for {
		_bbfda, _bbdbd := d.Token()
		if _bbdbd != nil {
			return _bbdbd
		}
		switch _aaca := _bbfda.(type) {
		case _d.StartElement:
			switch _aaca.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_bfee := NewCT_Name()
				if _dffba := d.DecodeElement(_bfee, &_aaca); _dffba != nil {
					return _dffba
				}
				_bgfb.Title = append(_bgfb.Title, _bfee)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_bccgg := NewCT_Description()
				if _cegg := d.DecodeElement(_bccgg, &_aaca); _cegg != nil {
					return _cegg
				}
				_bgfb.Desc = append(_bgfb.Desc, _bccgg)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_bgfb.CatLst = NewCT_Categories()
				if _fcdae := d.DecodeElement(_bgfb.CatLst, &_aaca); _fcdae != nil {
					return _fcdae
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0061\u006d\u0070\u0044\u0061\u0074\u0061"}:
				_bgfb.SampData = NewCT_SampleData()
				if _abae := d.DecodeElement(_bgfb.SampData, &_aaca); _abae != nil {
					return _abae
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073t\u0079\u006c\u0065\u0044\u0061\u0074a"}:
				_bgfb.StyleData = NewCT_SampleData()
				if _ffbc := d.DecodeElement(_bgfb.StyleData, &_aaca); _ffbc != nil {
					return _ffbc
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063l\u0072\u0044\u0061\u0074\u0061"}:
				_bgfb.ClrData = NewCT_SampleData()
				if _eeea := d.DecodeElement(_bgfb.ClrData, &_aaca); _eeea != nil {
					return _eeea
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				if _bfabe := d.DecodeElement(_bgfb.LayoutNode, &_aaca); _bfabe != nil {
					return _bfabe
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bgfb.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _ddcg := d.DecodeElement(_bgfb.ExtLst, &_aaca); _ddcg != nil {
					return _ddcg
				}
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u004c\u0061\u0079\u006f\u0075\u0074D\u0065\u0066 \u0025\u0076", _aaca.Name)
				if _bbeae := d.Skip(); _bbeae != nil {
					return _bbeae
				}
			}
		case _d.EndElement:
			break _aecff
		case _d.CharData:
		}
	}
	return nil
}
func NewCT_OrgChart() *CT_OrgChart { _eebb := &CT_OrgChart{}; return _eebb }

type CT_StyleDefinition struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	Title        []*CT_SDName
	Desc         []*CT_SDDescription
	CatLst       *CT_SDCategories
	Scene3d      *_ca.CT_Scene3D
	StyleLbl     []*CT_StyleLabel
	ExtLst       *_ca.CT_OfficeArtExtensionList
}

func (_egd *AG_ConstraintAttributes) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gc := range start.Attr {
		if _gc.Name.Local == "\u0074\u0079\u0070\u0065" {
			_egd.TypeAttr.UnmarshalXMLAttr(_gc)
			continue
		}
		if _gc.Name.Local == "\u0066\u006f\u0072" {
			_egd.ForAttr.UnmarshalXMLAttr(_gc)
			continue
		}
		if _gc.Name.Local == "\u0066o\u0072\u004e\u0061\u006d\u0065" {
			_ec, _ac := _gc.Value, error(nil)
			if _ac != nil {
				return _ac
			}
			_egd.ForNameAttr = &_ec
			continue
		}
		if _gc.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_egd.PtTypeAttr.UnmarshalXMLAttr(_gc)
			continue
		}
	}
	for {
		_ecf, _cbf := d.Token()
		if _cbf != nil {
			return _c.Errorf("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0041\u0047\u005f\u0043\u006f\u006es\u0074\u0072\u0061\u0069\u006e\u0074\u0041t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0073\u003a\u0020%\u0073", _cbf)
		}
		if _df, _eb := _ecf.(_d.EndElement); _eb && _df.Name == start.Name {
			break
		}
	}
	return nil
}

type ST_FlowDirection byte
type CT_AnimOne struct{ ValAttr ST_AnimOneStr }

func NewCT_NumericRule() *CT_NumericRule { _afdgb := &CT_NumericRule{}; return _afdgb }
func (_dccd *CT_RelIds) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _bgfd := range start.Attr {
		if _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _bgfd.Name.Local == "\u0064\u006d" || _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _bgfd.Name.Local == "\u0064\u006d" {
			_baeeb, _bdbc := _bgfd.Value, error(nil)
			if _bdbc != nil {
				return _bdbc
			}
			_dccd.DmAttr = _baeeb
			continue
		}
		if _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _bgfd.Name.Local == "\u006c\u006f" || _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _bgfd.Name.Local == "\u006c\u006f" {
			_ccgdd, _cfdg := _bgfd.Value, error(nil)
			if _cfdg != nil {
				return _cfdg
			}
			_dccd.LoAttr = _ccgdd
			continue
		}
		if _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _bgfd.Name.Local == "\u0071\u0073" || _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _bgfd.Name.Local == "\u0071\u0073" {
			_aaab, _abga := _bgfd.Value, error(nil)
			if _abga != nil {
				return _abga
			}
			_dccd.QsAttr = _aaab
			continue
		}
		if _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _bgfd.Name.Local == "\u0063\u0073" || _bgfd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _bgfd.Name.Local == "\u0063\u0073" {
			_bcddc, _cceb := _bgfd.Value, error(nil)
			if _cceb != nil {
				return _cceb
			}
			_dccd.CsAttr = _bcddc
			continue
		}
	}
	for {
		_adeda, _bfaa := d.Token()
		if _bfaa != nil {
			return _c.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0052\u0065\u006cI\u0064\u0073\u003a\u0020\u0025\u0073", _bfaa)
		}
		if _ecggc, _acdf := _adeda.(_d.EndElement); _acdf && _ecggc.Name == start.Name {
			break
		}
	}
	return nil
}
func ParseSliceST_UnsignedInts(s string) (ST_UnsignedInts, error) { return ST_UnsignedInts{}, nil }
func NewCT_TextProps() *CT_TextProps                              { _gbcaf := &CT_TextProps{}; return _gbcaf }
func (_adecaf ST_RotationPath) Validate() error                   { return _adecaf.ValidateWithPath("") }
func (_ebgag *ST_Direction) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gada, _gdgb := d.Token()
	if _gdgb != nil {
		return _gdgb
	}
	if _egfe, _cgebf := _gada.(_d.EndElement); _cgebf && _egfe.Name == start.Name {
		*_ebgag = 1
		return nil
	}
	if _cgdccb, _dfga := _gada.(_d.CharData); !_dfga {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gada)
	} else {
		switch string(_cgdccb) {
		case "":
			*_ebgag = 0
		case "\u006e\u006f\u0072\u006d":
			*_ebgag = 1
		case "\u0072\u0065\u0076":
			*_ebgag = 2
		}
	}
	_gada, _gdgb = d.Token()
	if _gdgb != nil {
		return _gdgb
	}
	if _dcgcc, _aece := _gada.(_d.EndElement); _aece && _dcgcc.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gada)
}
func (_bafb ST_VariableType) Validate() error { return _bafb.ValidateWithPath("") }

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (_gbeb *CT_Shape) ValidateWithPath(path string) error {
	if _gbeb.TypeAttr != nil {
		if _faeb := _gbeb.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _faeb != nil {
			return _faeb
		}
	}
	if _gbeb.AdjLst != nil {
		if _dgfc := _gbeb.AdjLst.ValidateWithPath(path + "\u002fA\u0064\u006a\u004c\u0073\u0074"); _dgfc != nil {
			return _dgfc
		}
	}
	if _gbeb.ExtLst != nil {
		if _fdfa := _gbeb.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _fdfa != nil {
			return _fdfa
		}
	}
	return nil
}

type CT_DiagramDefinitionHeader struct {
	UniqueIdAttr string
	MinVerAttr   *string
	DefStyleAttr *string
	ResIdAttr    *int32
	Title        []*CT_Name
	Desc         []*CT_Description
	CatLst       *CT_Categories
	ExtLst       *_ca.CT_OfficeArtExtensionList
}

func (_gccf *CT_OrgChart) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _gccf.ValAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_gccf.ValAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_egfcb ST_ConstraintType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_egfcb.String(), start)
}
func (_deefe ST_CxnType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dfbe := _d.Attr{}
	_dfbe.Name = name
	switch _deefe {
	case ST_CxnTypeUnset:
		_dfbe.Value = ""
	case ST_CxnTypeParOf:
		_dfbe.Value = "\u0070\u0061\u0072O\u0066"
	case ST_CxnTypePresOf:
		_dfbe.Value = "\u0070\u0072\u0065\u0073\u004f\u0066"
	case ST_CxnTypePresParOf:
		_dfbe.Value = "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff"
	case ST_CxnTypeUnknownRelationship:
		_dfbe.Value = "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"
	}
	return _dfbe, nil
}
func (_cdaea *ST_AnimOneStr) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_cdaea = 0
	case "\u006e\u006f\u006e\u0065":
		*_cdaea = 1
	case "\u006f\u006e\u0065":
		*_cdaea = 2
	case "\u0062\u0072\u0061\u006e\u0063\u0068":
		*_cdaea = 3
	}
	return nil
}

// ValidateWithPath validates the LayoutDef and its children, prefixing error messages with path
func (_cfcf *LayoutDef) ValidateWithPath(path string) error {
	if _fffdb := _cfcf.CT_DiagramDefinition.ValidateWithPath(path); _fffdb != nil {
		return _fffdb
	}
	return nil
}
func (_gdbedg ST_ResizeHandlesStr) Validate() error { return _gdbedg.ValidateWithPath("") }

type CT_Cxn struct {
	ModelIdAttr    ST_ModelId
	TypeAttr       ST_CxnType
	SrcIdAttr      ST_ModelId
	DestIdAttr     ST_ModelId
	SrcOrdAttr     uint32
	DestOrdAttr    uint32
	ParTransIdAttr *ST_ModelId
	SibTransIdAttr *ST_ModelId
	PresIdAttr     *string
	ExtLst         *_ca.CT_OfficeArtExtensionList
}

func (_cbae *ST_ParameterId) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_cbae = 0
	case "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n":
		*_cbae = 1
	case "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n":
		*_cbae = 2
	case "\u0063\u0068\u0044i\u0072":
		*_cbae = 3
	case "\u0063h\u0041\u006c\u0069\u0067\u006e":
		*_cbae = 4
	case "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e":
		*_cbae = 5
	case "\u006c\u0069\u006e\u0044\u0069\u0072":
		*_cbae = 6
	case "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r":
		*_cbae = 7
	case "\u0073\u0074\u0045\u006c\u0065\u006d":
		*_cbae = 8
	case "\u0062\u0065\u006e\u0064\u0050\u0074":
		*_cbae = 9
	case "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074":
		*_cbae = 10
	case "\u0062\u0065\u0067\u0053\u0074\u0079":
		*_cbae = 11
	case "\u0065\u006e\u0064\u0053\u0074\u0079":
		*_cbae = 12
	case "\u0064\u0069\u006d":
		*_cbae = 13
	case "\u0072o\u0074\u0050\u0061\u0074\u0068":
		*_cbae = 14
	case "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p":
		*_cbae = 15
	case "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e":
		*_cbae = 16
	case "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e":
		*_cbae = 17
	case "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b":
		*_cbae = 18
	case "\u0074\u0078\u0044i\u0072":
		*_cbae = 19
	case "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073":
		*_cbae = 20
	case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072":
		*_cbae = 21
	case "\u0074x\u0042\u006c\u0044\u0069\u0072":
		*_cbae = 22
	case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a":
		*_cbae = 23
	case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074":
		*_cbae = 24
	case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068":
		*_cbae = 25
	case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068":
		*_cbae = 26
	case "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e":
		*_cbae = 27
	case "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e":
		*_cbae = 28
	case "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
		*_cbae = 29
	case "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
		*_cbae = 30
	case "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft":
		*_cbae = 31
	case "\u0067\u0072\u0044i\u0072":
		*_cbae = 32
	case "\u0066l\u006f\u0077\u0044\u0069\u0072":
		*_cbae = 33
	case "\u0063o\u006e\u0074\u0044\u0069\u0072":
		*_cbae = 34
	case "\u0062\u006b\u0070\u0074":
		*_cbae = 35
	case "\u006f\u0066\u0066":
		*_cbae = 36
	case "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n":
		*_cbae = 37
	case "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c":
		*_cbae = 38
	case "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c":
		*_cbae = 39
	case "\u0073\u0074\u0041n\u0067":
		*_cbae = 40
	case "\u0073p\u0061\u006e\u0041\u006e\u0067":
		*_cbae = 41
	case "\u0061\u0072":
		*_cbae = 42
	case "\u006cn\u0053\u0070\u0050\u0061\u0072":
		*_cbae = 43
	case "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050":
		*_cbae = 44
	case "\u006c\u006e\u0053\u0070\u0043\u0068":
		*_cbae = 45
	case "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P":
		*_cbae = 46
	case "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074":
		*_cbae = 47
	case "\u0061l\u0069\u0067\u006e\u0054\u0078":
		*_cbae = 48
	case "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065":
		*_cbae = 49
	case "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065":
		*_cbae = 50
	case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065":
		*_cbae = 51
	case "\u0073r\u0063\u004e\u006f\u0064\u0065":
		*_cbae = 52
	case "\u0064s\u0074\u004e\u006f\u0064\u0065":
		*_cbae = 53
	case "\u0062\u0065\u0067\u0050\u0074\u0073":
		*_cbae = 54
	case "\u0065\u006e\u0064\u0050\u0074\u0073":
		*_cbae = 55
	}
	return nil
}

// Validate validates the CT_Category and its children
func (_ega *CT_Category) Validate() error {
	return _ega.ValidateWithPath("C\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079")
}
func (_acfff ST_ConnectorDimension) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_acfff.String(), start)
}
func NewAG_IteratorAttributes() *AG_IteratorAttributes { _dce := &AG_IteratorAttributes{}; return _dce }
func (_bgce ST_CxnType) ValidateWithPath(path string) error {
	switch _bgce {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bgce))
	}
	return nil
}
func (_cabf *CT_CxnList) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _cabf.Cxn != nil {
		_aefc := _d.StartElement{Name: _d.Name{Local: "\u0063\u0078\u006e"}}
		for _, _ecfb := range _cabf.Cxn {
			e.EncodeElement(_ecfb, _aefc)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_decg *ST_ClrAppMethod) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gbgf, _cbda := d.Token()
	if _cbda != nil {
		return _cbda
	}
	if _aeccb, _gbgc := _gbgf.(_d.EndElement); _gbgc && _aeccb.Name == start.Name {
		*_decg = 1
		return nil
	}
	if _bfgfe, _ebabc := _gbgf.(_d.CharData); !_ebabc {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbgf)
	} else {
		switch string(_bfgfe) {
		case "":
			*_decg = 0
		case "\u0073\u0070\u0061\u006e":
			*_decg = 1
		case "\u0063\u0079\u0063l\u0065":
			*_decg = 2
		case "\u0072\u0065\u0070\u0065\u0061\u0074":
			*_decg = 3
		}
	}
	_gbgf, _cbda = d.Token()
	if _cbda != nil {
		return _cbda
	}
	if _fgefb, _bceb := _gbgf.(_d.EndElement); _bceb && _fgefb.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbgf)
}
func (_aecgf ST_BoolOperator) ValidateWithPath(path string) error {
	switch _aecgf {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aecgf))
	}
	return nil
}

// ValidateWithPath validates the CT_Otherwise and its children, prefixing error messages with path
func (_ecfc *CT_Otherwise) ValidateWithPath(path string) error {
	for _gaaf, _egbdf := range _ecfc.Alg {
		if _fgdgd := _egbdf.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _gaaf)); _fgdgd != nil {
			return _fgdgd
		}
	}
	for _facbd, _bfga := range _ecfc.Shape {
		if _caag := _bfga.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _facbd)); _caag != nil {
			return _caag
		}
	}
	for _gdeb, _gcda := range _ecfc.PresOf {
		if _fefe := _gcda.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _gdeb)); _fefe != nil {
			return _fefe
		}
	}
	for _bgef, _ffdf := range _ecfc.ConstrLst {
		if _fdcf := _ffdf.ValidateWithPath(_c.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _bgef)); _fdcf != nil {
			return _fdcf
		}
	}
	for _cfcg, _efcdb := range _ecfc.RuleLst {
		if _acde := _efcdb.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _cfcg)); _acde != nil {
			return _acde
		}
	}
	for _babf, _aage := range _ecfc.ForEach {
		if _dabf := _aage.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _babf)); _dabf != nil {
			return _dabf
		}
	}
	for _cdcc, _ccgd := range _ecfc.LayoutNode {
		if _deec := _ccgd.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _cdcc)); _deec != nil {
			return _deec
		}
	}
	for _bac, _gadb := range _ecfc.Choose {
		if _caaga := _gadb.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _bac)); _caaga != nil {
			return _caaga
		}
	}
	for _facbf, _eeeg := range _ecfc.ExtLst {
		if _fecag := _eeeg.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _facbf)); _fecag != nil {
			return _fecag
		}
	}
	return nil
}
func (_gbgbcc *ST_FunctionValue) Validate() error { return _gbgbcc.ValidateWithPath("") }
func (_fea *CT_CTCategory) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", _fea.TypeAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0069"}, Value: _c.Sprintf("\u0025\u0076", _fea.PriAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_ResizeHandles and its children
func (_daffg *CT_ResizeHandles) Validate() error {
	return _daffg.ValidateWithPath("\u0043\u0054_\u0052\u0065\u0073i\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073")
}
func (_beafd ST_DiagramHorizontalAlignment) Validate() error { return _beafd.ValidateWithPath("") }

// ValidateWithPath validates the CT_StyleDefinitionHeaderLst and its children, prefixing error messages with path
func (_eaada *CT_StyleDefinitionHeaderLst) ValidateWithPath(path string) error {
	for _dcfcg, _ecde := range _eaada.StyleDefHdr {
		if _bgcff := _ecde.ValidateWithPath(_c.Sprintf("\u0025s\u002fS\u0074\u0079\u006c\u0065\u0044e\u0066\u0048d\u0072\u005b\u0025\u0064\u005d", path, _dcfcg)); _bgcff != nil {
			return _bgcff
		}
	}
	return nil
}
func (_fbae *CT_Categories) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_edg:
	for {
		_gcb, _bbdf := d.Token()
		if _bbdf != nil {
			return _bbdf
		}
		switch _eaa := _gcb.(type) {
		case _d.StartElement:
			switch _eaa.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074"}:
				_aga := NewCT_Category()
				if _aabb := d.DecodeElement(_aga, &_eaa); _aabb != nil {
					return _aabb
				}
				_fbae.Cat = append(_fbae.Cat, _aga)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043a\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073 \u0025\u0076", _eaa.Name)
				if _ecac := d.Skip(); _ecac != nil {
					return _ecac
				}
			}
		case _d.EndElement:
			break _edg
		case _d.CharData:
		}
	}
	return nil
}
func (_gdbccb *ST_TextAnchorHorizontal) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gdbccb = 0
	case "\u006e\u006f\u006e\u0065":
		*_gdbccb = 1
	case "\u0063\u0074\u0072":
		*_gdbccb = 2
	}
	return nil
}
func (_bbcf ST_ParameterId) ValidateWithPath(path string) error {
	switch _bbcf {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bbcf))
	}
	return nil
}
func (_aafg ST_HierBranchStyle) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ebaa := _d.Attr{}
	_ebaa.Name = name
	switch _aafg {
	case ST_HierBranchStyleUnset:
		_ebaa.Value = ""
	case ST_HierBranchStyleL:
		_ebaa.Value = "\u006c"
	case ST_HierBranchStyleR:
		_ebaa.Value = "\u0072"
	case ST_HierBranchStyleHang:
		_ebaa.Value = "\u0068\u0061\u006e\u0067"
	case ST_HierBranchStyleStd:
		_ebaa.Value = "\u0073\u0074\u0064"
	case ST_HierBranchStyleInit:
		_ebaa.Value = "\u0069\u006e\u0069\u0074"
	}
	return _ebaa, nil
}
func (_abbf ST_ConnectorDimension) String() string {
	switch _abbf {
	case 0:
		return ""
	case 1:
		return "\u0031\u0044"
	case 2:
		return "\u0032\u0044"
	case 3:
		return "\u0063\u0075\u0073\u0074"
	}
	return ""
}

type ST_TextBlockDirection byte

func (_baggg *LayoutDefHdrLst) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006ca\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074"
	return _baggg.CT_DiagramDefinitionHeaderLst.MarshalXML(e, start)
}
func (_bdede *ST_ConnectorPoint) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_eecgb, _ebef := d.Token()
	if _ebef != nil {
		return _ebef
	}
	if _ggfc, _bbfc := _eecgb.(_d.EndElement); _bbfc && _ggfc.Name == start.Name {
		*_bdede = 1
		return nil
	}
	if _dbeg, _ggagg := _eecgb.(_d.CharData); !_ggagg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eecgb)
	} else {
		switch string(_dbeg) {
		case "":
			*_bdede = 0
		case "\u0061\u0075\u0074\u006f":
			*_bdede = 1
		case "\u0062\u0043\u0074\u0072":
			*_bdede = 2
		case "\u0063\u0074\u0072":
			*_bdede = 3
		case "\u006d\u0069\u0064\u004c":
			*_bdede = 4
		case "\u006d\u0069\u0064\u0052":
			*_bdede = 5
		case "\u0074\u0043\u0074\u0072":
			*_bdede = 6
		case "\u0062\u004c":
			*_bdede = 7
		case "\u0062\u0052":
			*_bdede = 8
		case "\u0074\u004c":
			*_bdede = 9
		case "\u0074\u0052":
			*_bdede = 10
		case "\u0072\u0061\u0064\u0069\u0061\u006c":
			*_bdede = 11
		}
	}
	_eecgb, _ebef = d.Token()
	if _ebef != nil {
		return _ebef
	}
	if _eadec, _babg := _eecgb.(_d.EndElement); _babg && _eadec.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eecgb)
}
func (_beebe ST_Breakpoint) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_bccad := _d.Attr{}
	_bccad.Name = name
	switch _beebe {
	case ST_BreakpointUnset:
		_bccad.Value = ""
	case ST_BreakpointEndCnv:
		_bccad.Value = "\u0065\u006e\u0064\u0043\u006e\u0076"
	case ST_BreakpointBal:
		_bccad.Value = "\u0062\u0061\u006c"
	case ST_BreakpointFixed:
		_bccad.Value = "\u0066\u0069\u0078e\u0064"
	}
	return _bccad, nil
}
func (_bebfd ST_ContinueDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_fdbb := _d.Attr{}
	_fdbb.Name = name
	switch _bebfd {
	case ST_ContinueDirectionUnset:
		_fdbb.Value = ""
	case ST_ContinueDirectionRevDir:
		_fdbb.Value = "\u0072\u0065\u0076\u0044\u0069\u0072"
	case ST_ContinueDirectionSameDir:
		_fdbb.Value = "\u0073a\u006d\u0065\u0044\u0069\u0072"
	}
	return _fdbb, nil
}

const (
	ST_PyramidAccentTextMarginUnset ST_PyramidAccentTextMargin = 0
	ST_PyramidAccentTextMarginStep  ST_PyramidAccentTextMargin = 1
	ST_PyramidAccentTextMarginStack ST_PyramidAccentTextMargin = 2
)

func (_dcebf *CT_StyleDefinitionHeader) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gcgf := range start.Attr {
		if _gcgf.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_efaf, _ddca := _gcgf.Value, error(nil)
			if _ddca != nil {
				return _ddca
			}
			_dcebf.UniqueIdAttr = _efaf
			continue
		}
		if _gcgf.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_dfeda, _bcdda := _gcgf.Value, error(nil)
			if _bcdda != nil {
				return _bcdda
			}
			_dcebf.MinVerAttr = &_dfeda
			continue
		}
		if _gcgf.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_bddg, _ffdag := _ae.ParseInt(_gcgf.Value, 10, 32)
			if _ffdag != nil {
				return _ffdag
			}
			_dfaae := int32(_bddg)
			_dcebf.ResIdAttr = &_dfaae
			continue
		}
	}
_bgge:
	for {
		_agab, _agfbb := d.Token()
		if _agfbb != nil {
			return _agfbb
		}
		switch _cbfg := _agab.(type) {
		case _d.StartElement:
			switch _cbfg.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_afed := NewCT_SDName()
				if _egecf := d.DecodeElement(_afed, &_cbfg); _egecf != nil {
					return _egecf
				}
				_dcebf.Title = append(_dcebf.Title, _afed)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_ebga := NewCT_SDDescription()
				if _gbce := d.DecodeElement(_ebga, &_cbfg); _gbce != nil {
					return _gbce
				}
				_dcebf.Desc = append(_dcebf.Desc, _ebga)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_dcebf.CatLst = NewCT_SDCategories()
				if _ebab := d.DecodeElement(_dcebf.CatLst, &_cbfg); _ebab != nil {
					return _ebab
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_dcebf.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _bbfg := d.DecodeElement(_dcebf.ExtLst, &_cbfg); _bbfg != nil {
					return _bbfg
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048e\u0061\u0064\u0065\u0072\u0020%\u0076", _cbfg.Name)
				if _ddee := d.Skip(); _ddee != nil {
					return _ddee
				}
			}
		case _d.EndElement:
			break _bgge
		case _d.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_AnimOne and its children, prefixing error messages with path
func (_dfea *CT_AnimOne) ValidateWithPath(path string) error {
	if _fcgg := _dfea.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fcgg != nil {
		return _fcgg
	}
	return nil
}

// Validate validates the CT_Adj and its children
func (_egg *CT_Adj) Validate() error {
	return _egg.ValidateWithPath("\u0043\u0054\u005f\u0041\u0064\u006a")
}
func ParseSliceST_Booleans(s string) (ST_Booleans, error) { return ST_Booleans{}, nil }
func (_eegg *CT_ForEach) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _eegg.NameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.NameAttr)})
	}
	if _eegg.RefAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u0065\u0066"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.RefAttr)})
	}
	if _eegg.AxisAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.AxisAttr)})
	}
	if _eegg.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.PtTypeAttr)})
	}
	if _eegg.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.HideLastTransAttr)})
	}
	if _eegg.StAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.StAttr)})
	}
	if _eegg.CntAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u006e\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.CntAttr)})
	}
	if _eegg.StepAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _c.Sprintf("\u0025\u0076", *_eegg.StepAttr)})
	}
	e.EncodeToken(start)
	if _eegg.Alg != nil {
		_ddgd := _d.StartElement{Name: _d.Name{Local: "\u0061\u006c\u0067"}}
		for _, _gfef := range _eegg.Alg {
			e.EncodeElement(_gfef, _ddgd)
		}
	}
	if _eegg.Shape != nil {
		_gcfe := _d.StartElement{Name: _d.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _bgbdg := range _eegg.Shape {
			e.EncodeElement(_bgbdg, _gcfe)
		}
	}
	if _eegg.PresOf != nil {
		_bedeb := _d.StartElement{Name: _d.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _geeea := range _eegg.PresOf {
			e.EncodeElement(_geeea, _bedeb)
		}
	}
	if _eegg.ConstrLst != nil {
		_ggad := _d.StartElement{Name: _d.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _gff := range _eegg.ConstrLst {
			e.EncodeElement(_gff, _ggad)
		}
	}
	if _eegg.RuleLst != nil {
		_edde := _d.StartElement{Name: _d.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _fdggg := range _eegg.RuleLst {
			e.EncodeElement(_fdggg, _edde)
		}
	}
	if _eegg.ForEach != nil {
		_gdcg := _d.StartElement{Name: _d.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _ddgg := range _eegg.ForEach {
			e.EncodeElement(_ddgg, _gdcg)
		}
	}
	if _eegg.LayoutNode != nil {
		_ccda := _d.StartElement{Name: _d.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _abcee := range _eegg.LayoutNode {
			e.EncodeElement(_abcee, _ccda)
		}
	}
	if _eegg.Choose != nil {
		_cbgc := _d.StartElement{Name: _d.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _afad := range _eegg.Choose {
			e.EncodeElement(_afad, _cbgc)
		}
	}
	if _eegg.ExtLst != nil {
		_ebdf := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _faff := range _eegg.ExtLst {
			e.EncodeElement(_faff, _ebdf)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_PtList and its children
func (_cefa *CT_PtList) Validate() error {
	return _cefa.ValidateWithPath("\u0043T\u005f\u0050\u0074\u004c\u0069\u0073t")
}
func (_egegg ST_PtType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_egegg.String(), start)
}

const (
	ST_ChildAlignmentUnset ST_ChildAlignment = 0
	ST_ChildAlignmentT     ST_ChildAlignment = 1
	ST_ChildAlignmentB     ST_ChildAlignment = 2
	ST_ChildAlignmentL     ST_ChildAlignment = 3
	ST_ChildAlignmentR     ST_ChildAlignment = 4
)

type ST_TextDirection byte

func (_fbga ST_ParameterId) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fbga.String(), start)
}

// ValidateWithPath validates the CT_Name and its children, prefixing error messages with path
func (_aaafg *CT_Name) ValidateWithPath(path string) error { return nil }

// Validate validates the AG_IteratorAttributes and its children
func (_cab *AG_IteratorAttributes) Validate() error {
	return _cab.ValidateWithPath("A\u0047\u005f\u0049\u0074er\u0061t\u006f\u0072\u0041\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u0073")
}
func (_bggde ST_PtType) ValidateWithPath(path string) error {
	switch _bggde {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bggde))
	}
	return nil
}
func NewCT_AdjLst() *CT_AdjLst { _dbd := &CT_AdjLst{}; return _dbd }

type LayoutDefHdrLst struct{ CT_DiagramDefinitionHeaderLst }

func (_eefae ST_LinearDirection) ValidateWithPath(path string) error {
	switch _eefae {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eefae))
	}
	return nil
}

type CT_SampleData struct {
	UseDefAttr *bool
	DataModel  *CT_DataModel
}

func (_gcbgb *ST_ChildAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gcbgb = 0
	case "\u0074":
		*_gcbgb = 1
	case "\u0062":
		*_gcbgb = 2
	case "\u006c":
		*_gcbgb = 3
	case "\u0072":
		*_gcbgb = 4
	}
	return nil
}
func (_aagg ST_GrowDirection) Validate() error { return _aagg.ValidateWithPath("") }
func (_dbfae *LayoutDefHdr) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dbfae.CT_DiagramDefinitionHeader = *NewCT_DiagramDefinitionHeader()
	for _, _dgfab := range start.Attr {
		if _dgfab.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_dcfg, _aaccc := _dgfab.Value, error(nil)
			if _aaccc != nil {
				return _aaccc
			}
			_dbfae.UniqueIdAttr = _dcfg
			continue
		}
		if _dgfab.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_febea, _cccea := _dgfab.Value, error(nil)
			if _cccea != nil {
				return _cccea
			}
			_dbfae.MinVerAttr = &_febea
			continue
		}
		if _dgfab.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_dgaf, _fcbdf := _dgfab.Value, error(nil)
			if _fcbdf != nil {
				return _fcbdf
			}
			_dbfae.DefStyleAttr = &_dgaf
			continue
		}
		if _dgfab.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_beacb, _gccc := _ae.ParseInt(_dgfab.Value, 10, 32)
			if _gccc != nil {
				return _gccc
			}
			_fffea := int32(_beacb)
			_dbfae.ResIdAttr = &_fffea
			continue
		}
	}
_fcaf:
	for {
		_fgceg, _ggdc := d.Token()
		if _ggdc != nil {
			return _ggdc
		}
		switch _fbbgf := _fgceg.(type) {
		case _d.StartElement:
			switch _fbbgf.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_gcaad := NewCT_Name()
				if _bcac := d.DecodeElement(_gcaad, &_fbbgf); _bcac != nil {
					return _bcac
				}
				_dbfae.Title = append(_dbfae.Title, _gcaad)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_gdfcd := NewCT_Description()
				if _fggc := d.DecodeElement(_gdfcd, &_fbbgf); _fggc != nil {
					return _fggc
				}
				_dbfae.Desc = append(_dbfae.Desc, _gdfcd)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_dbfae.CatLst = NewCT_Categories()
				if _dgfaf := d.DecodeElement(_dbfae.CatLst, &_fbbgf); _dgfaf != nil {
					return _dgfaf
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_dbfae.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _aebg := d.DecodeElement(_dbfae.ExtLst, &_fbbgf); _aebg != nil {
					return _aebg
				}
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u004c\u0061yo\u0075\u0074D\u0065\u0066\u0048\u0064\u0072\u0020\u0025\u0076", _fbbgf.Name)
				if _bbdef := d.Skip(); _bbdef != nil {
					return _bbdef
				}
			}
		case _d.EndElement:
			break _fcaf
		case _d.CharData:
		}
	}
	return nil
}
func (_ceaab *ST_CxnType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ceaab = 0
	case "\u0070\u0061\u0072O\u0066":
		*_ceaab = 1
	case "\u0070\u0072\u0065\u0073\u004f\u0066":
		*_ceaab = 2
	case "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff":
		*_ceaab = 3
	case "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070":
		*_ceaab = 4
	}
	return nil
}
func (_baee *CT_HierBranchStyle) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _baee.ValAttr != ST_HierBranchStyleUnset {
		_ccg, _bdda := _baee.ValAttr.MarshalXMLAttr(_d.Name{Local: "\u0076\u0061\u006c"})
		if _bdda != nil {
			return _bdda
		}
		start.Attr = append(start.Attr, _ccg)
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_gega ST_ConnectorPoint) String() string {
	switch _gega {
	case 0:
		return ""
	case 1:
		return "\u0061\u0075\u0074\u006f"
	case 2:
		return "\u0062\u0043\u0074\u0072"
	case 3:
		return "\u0063\u0074\u0072"
	case 4:
		return "\u006d\u0069\u0064\u004c"
	case 5:
		return "\u006d\u0069\u0064\u0052"
	case 6:
		return "\u0074\u0043\u0074\u0072"
	case 7:
		return "\u0062\u004c"
	case 8:
		return "\u0062\u0052"
	case 9:
		return "\u0074\u004c"
	case 10:
		return "\u0074\u0052"
	case 11:
		return "\u0072\u0061\u0064\u0069\u0061\u006c"
	}
	return ""
}
func NewCT_HierBranchStyle() *CT_HierBranchStyle { _efge := &CT_HierBranchStyle{}; return _efge }
func NewCT_SDCategories() *CT_SDCategories       { _deaf := &CT_SDCategories{}; return _deaf }
func (_cabgf ST_ConnectorRouting) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ffdae := _d.Attr{}
	_ffdae.Name = name
	switch _cabgf {
	case ST_ConnectorRoutingUnset:
		_ffdae.Value = ""
	case ST_ConnectorRoutingStra:
		_ffdae.Value = "\u0073\u0074\u0072\u0061"
	case ST_ConnectorRoutingBend:
		_ffdae.Value = "\u0062\u0065\u006e\u0064"
	case ST_ConnectorRoutingCurve:
		_ffdae.Value = "\u0063\u0075\u0072v\u0065"
	case ST_ConnectorRoutingLongCurve:
		_ffdae.Value = "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e"
	}
	return _ffdae, nil
}

// ValidateWithPath validates the CT_PresentationOf and its children, prefixing error messages with path
func (_cfce *CT_PresentationOf) ValidateWithPath(path string) error {
	if _cfce.ExtLst != nil {
		if _ebff := _cfce.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _ebff != nil {
			return _ebff
		}
	}
	return nil
}
func (_aacbe *CT_DiagramDefinition) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _aacbe.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_aacbe.UniqueIdAttr)})
	}
	if _aacbe.MinVerAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _c.Sprintf("\u0025\u0076", *_aacbe.MinVerAttr)})
	}
	if _aacbe.DefStyleAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_aacbe.DefStyleAttr)})
	}
	e.EncodeToken(start)
	if _aacbe.Title != nil {
		_fage := _d.StartElement{Name: _d.Name{Local: "\u0074\u0069\u0074l\u0065"}}
		for _, _ggdd := range _aacbe.Title {
			e.EncodeElement(_ggdd, _fage)
		}
	}
	if _aacbe.Desc != nil {
		_cdcd := _d.StartElement{Name: _d.Name{Local: "\u0064\u0065\u0073\u0063"}}
		for _, _cfee := range _aacbe.Desc {
			e.EncodeElement(_cfee, _cdcd)
		}
	}
	if _aacbe.CatLst != nil {
		_bfef := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_aacbe.CatLst, _bfef)
	}
	if _aacbe.SampData != nil {
		_fccad := _d.StartElement{Name: _d.Name{Local: "\u0073\u0061\u006d\u0070\u0044\u0061\u0074\u0061"}}
		e.EncodeElement(_aacbe.SampData, _fccad)
	}
	if _aacbe.StyleData != nil {
		_agaeg := _d.StartElement{Name: _d.Name{Local: "\u0073t\u0079\u006c\u0065\u0044\u0061\u0074a"}}
		e.EncodeElement(_aacbe.StyleData, _agaeg)
	}
	if _aacbe.ClrData != nil {
		_aeec := _d.StartElement{Name: _d.Name{Local: "\u0063l\u0072\u0044\u0061\u0074\u0061"}}
		e.EncodeElement(_aacbe.ClrData, _aeec)
	}
	_geed := _d.StartElement{Name: _d.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
	e.EncodeElement(_aacbe.LayoutNode, _geed)
	if _aacbe.ExtLst != nil {
		_gdaa := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_aacbe.ExtLst, _gdaa)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_fgbcb *ST_ElementType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fgbcb = 0
	case "\u0061\u006c\u006c":
		*_fgbcb = 1
	case "\u0064\u006f\u0063":
		*_fgbcb = 2
	case "\u006e\u006f\u0064\u0065":
		*_fgbcb = 3
	case "\u006e\u006f\u0072\u006d":
		*_fgbcb = 4
	case "\u006eo\u006e\u004e\u006f\u0072\u006d":
		*_fgbcb = 5
	case "\u0061\u0073\u0073\u0074":
		*_fgbcb = 6
	case "\u006eo\u006e\u0041\u0073\u0073\u0074":
		*_fgbcb = 7
	case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
		*_fgbcb = 8
	case "\u0070\u0072\u0065\u0073":
		*_fgbcb = 9
	case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
		*_fgbcb = 10
	}
	return nil
}

type ST_HierarchyAlignment byte

func (_fcbg *CT_ColorTransform) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gbfc := range start.Attr {
		if _gbfc.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_gac, _fdeb := _gbfc.Value, error(nil)
			if _fdeb != nil {
				return _fdeb
			}
			_fcbg.UniqueIdAttr = &_gac
			continue
		}
		if _gbfc.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_dgd, _dfcg := _gbfc.Value, error(nil)
			if _dfcg != nil {
				return _dfcg
			}
			_fcbg.MinVerAttr = &_dgd
			continue
		}
	}
_bcg:
	for {
		_aag, _ece := d.Token()
		if _ece != nil {
			return _ece
		}
		switch _gcaa := _aag.(type) {
		case _d.StartElement:
			switch _gcaa.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_ffeb := NewCT_CTName()
				if _aegfg := d.DecodeElement(_ffeb, &_gcaa); _aegfg != nil {
					return _aegfg
				}
				_fcbg.Title = append(_fcbg.Title, _ffeb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_ffc := NewCT_CTDescription()
				if _gdaf := d.DecodeElement(_ffc, &_gcaa); _gdaf != nil {
					return _gdaf
				}
				_fcbg.Desc = append(_fcbg.Desc, _ffc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_fcbg.CatLst = NewCT_CTCategories()
				if _acgd := d.DecodeElement(_fcbg.CatLst, &_gcaa); _acgd != nil {
					return _acgd
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_bcgg := NewCT_CTStyleLabel()
				if _gaa := d.DecodeElement(_bcgg, &_gcaa); _gaa != nil {
					return _gaa
				}
				_fcbg.StyleLbl = append(_fcbg.StyleLbl, _bcgg)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_fcbg.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _fgce := d.DecodeElement(_fcbg.ExtLst, &_gcaa); _fgce != nil {
					return _fgce
				}
			default:
				_af.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0020\u0025\u0076", _gcaa.Name)
				if _gdc := d.Skip(); _gdc != nil {
					return _gdc
				}
			}
		case _d.EndElement:
			break _bcg
		case _d.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Cxn and its children, prefixing error messages with path
func (_aaeg *CT_Cxn) ValidateWithPath(path string) error {
	if _gdfc := _aaeg.ModelIdAttr.ValidateWithPath(path + "\u002f\u004d\u006fd\u0065\u006c\u0049\u0064\u0041\u0074\u0074\u0072"); _gdfc != nil {
		return _gdfc
	}
	if _bagf := _aaeg.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _bagf != nil {
		return _bagf
	}
	if _ace := _aaeg.SrcIdAttr.ValidateWithPath(path + "\u002f\u0053\u0072\u0063\u0049\u0064\u0041\u0074\u0074\u0072"); _ace != nil {
		return _ace
	}
	if _gaac := _aaeg.DestIdAttr.ValidateWithPath(path + "/\u0044\u0065\u0073\u0074\u0049\u0064\u0041\u0074\u0074\u0072"); _gaac != nil {
		return _gaac
	}
	if _aaeg.ParTransIdAttr != nil {
		if _fgaa := _aaeg.ParTransIdAttr.ValidateWithPath(path + "\u002fP\u0061r\u0054\u0072\u0061\u006e\u0073\u0049\u0064\u0041\u0074\u0074\u0072"); _fgaa != nil {
			return _fgaa
		}
	}
	if _aaeg.SibTransIdAttr != nil {
		if _dfg := _aaeg.SibTransIdAttr.ValidateWithPath(path + "\u002fS\u0069b\u0054\u0072\u0061\u006e\u0073\u0049\u0064\u0041\u0074\u0074\u0072"); _dfg != nil {
			return _dfg
		}
	}
	if _aaeg.ExtLst != nil {
		if _cde := _aaeg.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _cde != nil {
			return _cde
		}
	}
	return nil
}
func (_aefbg *ST_Breakpoint) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fcee, _cecf := d.Token()
	if _cecf != nil {
		return _cecf
	}
	if _ffde, _affbed := _fcee.(_d.EndElement); _affbed && _ffde.Name == start.Name {
		*_aefbg = 1
		return nil
	}
	if _dafag, _egfac := _fcee.(_d.CharData); !_egfac {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fcee)
	} else {
		switch string(_dafag) {
		case "":
			*_aefbg = 0
		case "\u0065\u006e\u0064\u0043\u006e\u0076":
			*_aefbg = 1
		case "\u0062\u0061\u006c":
			*_aefbg = 2
		case "\u0066\u0069\u0078e\u0064":
			*_aefbg = 3
		}
	}
	_fcee, _cecf = d.Token()
	if _cecf != nil {
		return _cecf
	}
	if _ccga, _aecad := _fcee.(_d.EndElement); _aecad && _ccga.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fcee)
}
func (_dcbag ST_LayoutShapeType) String() string {
	if _dcbag.ST_ShapeType != _ca.ST_ShapeTypeUnset {
		return _dcbag.ST_ShapeType.String()
	}
	if _dcbag.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		return _dcbag.ST_OutputShapeType.String()
	}
	return ""
}
func (_cbadf *ST_PrSetCustVal) ValidateWithPath(path string) error {
	_gbbcg := []string{}
	if _cbadf.ST_Percentage != nil {
		_gbbcg = append(_gbbcg, "\u0053\u0054\u005f\u0050\u0065\u0072\u0063\u0065\u006e\u0074\u0061\u0067\u0065")
	}
	if _cbadf.Int32 != nil {
		_gbbcg = append(_gbbcg, "\u0049\u006e\u00743\u0032")
	}
	if len(_gbbcg) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _gbbcg)
	}
	return nil
}
func (_eeae ST_ChildDirection) Validate() error { return _eeae.ValidateWithPath("") }
func (_baaef ST_OutputShapeType) String() string {
	switch _baaef {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0063\u006f\u006e\u006e"
	}
	return ""
}
func (_dacc ST_Direction) String() string {
	switch _dacc {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u0072\u006d"
	case 2:
		return "\u0072\u0065\u0076"
	}
	return ""
}
func (_ccefe ST_ClrAppMethod) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_aggbg := _d.Attr{}
	_aggbg.Name = name
	switch _ccefe {
	case ST_ClrAppMethodUnset:
		_aggbg.Value = ""
	case ST_ClrAppMethodSpan:
		_aggbg.Value = "\u0073\u0070\u0061\u006e"
	case ST_ClrAppMethodCycle:
		_aggbg.Value = "\u0063\u0079\u0063l\u0065"
	case ST_ClrAppMethodRepeat:
		_aggbg.Value = "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return _aggbg, nil
}
func NewCT_Algorithm() *CT_Algorithm {
	_dff := &CT_Algorithm{}
	_dff.TypeAttr = ST_AlgorithmType(1)
	return _dff
}
func (_egafg ST_ConstraintRelationship) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_agbef := _d.Attr{}
	_agbef.Name = name
	switch _egafg {
	case ST_ConstraintRelationshipUnset:
		_agbef.Value = ""
	case ST_ConstraintRelationshipSelf:
		_agbef.Value = "\u0073\u0065\u006c\u0066"
	case ST_ConstraintRelationshipCh:
		_agbef.Value = "\u0063\u0068"
	case ST_ConstraintRelationshipDes:
		_agbef.Value = "\u0064\u0065\u0073"
	}
	return _agbef, nil
}
func (_acfa *ColorsDefHdrLst) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063o\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074"
	return _acfa.CT_ColorTransformHeaderLst.MarshalXML(e, start)
}

type ST_NodeHorizontalAlignment byte

func (_aeab *CT_Constraint) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _dedc := range start.Attr {
		if _dedc.Name.Local == "\u006f\u0070" {
			_aeab.OpAttr.UnmarshalXMLAttr(_dedc)
			continue
		}
		if _dedc.Name.Local == "\u0076\u0061\u006c" {
			_bcaa, _gef := _ae.ParseFloat(_dedc.Value, 64)
			if _gef != nil {
				return _gef
			}
			_aeab.ValAttr = &_bcaa
			continue
		}
		if _dedc.Name.Local == "\u0066\u0061\u0063\u0074" {
			_ddbc, _eeef := _ae.ParseFloat(_dedc.Value, 64)
			if _eeef != nil {
				return _eeef
			}
			_aeab.FactAttr = &_ddbc
			continue
		}
		if _dedc.Name.Local == "\u0074\u0079\u0070\u0065" {
			_aeab.TypeAttr.UnmarshalXMLAttr(_dedc)
			continue
		}
		if _dedc.Name.Local == "\u0066\u006f\u0072" {
			_aeab.ForAttr.UnmarshalXMLAttr(_dedc)
			continue
		}
		if _dedc.Name.Local == "\u0066o\u0072\u004e\u0061\u006d\u0065" {
			_fbba, _bcaf := _dedc.Value, error(nil)
			if _bcaf != nil {
				return _bcaf
			}
			_aeab.ForNameAttr = &_fbba
			continue
		}
		if _dedc.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_aeab.PtTypeAttr.UnmarshalXMLAttr(_dedc)
			continue
		}
		if _dedc.Name.Local == "\u0072e\u0066\u0054\u0079\u0070\u0065" {
			_aeab.RefTypeAttr.UnmarshalXMLAttr(_dedc)
			continue
		}
		if _dedc.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072" {
			_aeab.RefForAttr.UnmarshalXMLAttr(_dedc)
			continue
		}
		if _dedc.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065" {
			_dcgga, _feae := _dedc.Value, error(nil)
			if _feae != nil {
				return _feae
			}
			_aeab.RefForNameAttr = &_dcgga
			continue
		}
		if _dedc.Name.Local == "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e" {
			_aeab.RefPtTypeAttr.UnmarshalXMLAttr(_dedc)
			continue
		}
	}
_ceed:
	for {
		_cdfe, _gfgd := d.Token()
		if _gfgd != nil {
			return _gfgd
		}
		switch _bgfa := _cdfe.(type) {
		case _d.StartElement:
			switch _bgfa.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_aeab.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _gbbc := d.DecodeElement(_aeab.ExtLst, &_bgfa); _gbbc != nil {
					return _gbbc
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043o\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074 \u0025\u0076", _bgfa.Name)
				if _aeb := d.Skip(); _aeb != nil {
					return _aeb
				}
			}
		case _d.EndElement:
			break _ceed
		case _d.CharData:
		}
	}
	return nil
}
func (_ecgcf ST_VerticalAlignment) String() string {
	switch _ecgcf {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u006d\u0069\u0064"
	case 3:
		return "\u0062"
	case 4:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}
func (_cdacg *CT_When) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _cdacg.NameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.NameAttr)})
	}
	_fgad, _eeeb := _cdacg.FuncAttr.MarshalXMLAttr(_d.Name{Local: "\u0066\u0075\u006e\u0063"})
	if _eeeb != nil {
		return _eeeb
	}
	start.Attr = append(start.Attr, _fgad)
	if _cdacg.ArgAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0061\u0072\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.ArgAttr)})
	}
	_fgad, _eeeb = _cdacg.OpAttr.MarshalXMLAttr(_d.Name{Local: "\u006f\u0070"})
	if _eeeb != nil {
		return _eeeb
	}
	start.Attr = append(start.Attr, _fgad)
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _cdacg.ValAttr)})
	if _cdacg.AxisAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.AxisAttr)})
	}
	if _cdacg.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.PtTypeAttr)})
	}
	if _cdacg.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.HideLastTransAttr)})
	}
	if _cdacg.StAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.StAttr)})
	}
	if _cdacg.CntAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u006e\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.CntAttr)})
	}
	if _cdacg.StepAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _c.Sprintf("\u0025\u0076", *_cdacg.StepAttr)})
	}
	e.EncodeToken(start)
	if _cdacg.Alg != nil {
		_acfcc := _d.StartElement{Name: _d.Name{Local: "\u0061\u006c\u0067"}}
		for _, _cgdf := range _cdacg.Alg {
			e.EncodeElement(_cgdf, _acfcc)
		}
	}
	if _cdacg.Shape != nil {
		_cgfd := _d.StartElement{Name: _d.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _cadga := range _cdacg.Shape {
			e.EncodeElement(_cadga, _cgfd)
		}
	}
	if _cdacg.PresOf != nil {
		_bafgc := _d.StartElement{Name: _d.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _dgbde := range _cdacg.PresOf {
			e.EncodeElement(_dgbde, _bafgc)
		}
	}
	if _cdacg.ConstrLst != nil {
		_agbb := _d.StartElement{Name: _d.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _aadd := range _cdacg.ConstrLst {
			e.EncodeElement(_aadd, _agbb)
		}
	}
	if _cdacg.RuleLst != nil {
		_gbgbc := _d.StartElement{Name: _d.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _faba := range _cdacg.RuleLst {
			e.EncodeElement(_faba, _gbgbc)
		}
	}
	if _cdacg.ForEach != nil {
		_fbcg := _d.StartElement{Name: _d.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _eefbd := range _cdacg.ForEach {
			e.EncodeElement(_eefbd, _fbcg)
		}
	}
	if _cdacg.LayoutNode != nil {
		_gdbec := _d.StartElement{Name: _d.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _fbcc := range _cdacg.LayoutNode {
			e.EncodeElement(_fbcc, _gdbec)
		}
	}
	if _cdacg.Choose != nil {
		_agcc := _d.StartElement{Name: _d.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _ggge := range _cdacg.Choose {
			e.EncodeElement(_ggge, _agcc)
		}
	}
	if _cdacg.ExtLst != nil {
		_abbae := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _bdafe := range _cdacg.ExtLst {
			e.EncodeElement(_bdafe, _abbae)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Rules and its children, prefixing error messages with path
func (_ggae *CT_Rules) ValidateWithPath(path string) error {
	for _baba, _egbfg := range _ggae.Rule {
		if _fbcf := _egbfg.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0052\u0075\u006c\u0065\u005b\u0025\u0064\u005d", path, _baba)); _fbcf != nil {
			return _fbcf
		}
	}
	return nil
}

type ST_ArrowheadStyle byte

func (_dffcf *ST_BoolOperator) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bgeae, _fafe := d.Token()
	if _fafe != nil {
		return _fafe
	}
	if _abccd, _eggea := _bgeae.(_d.EndElement); _eggea && _abccd.Name == start.Name {
		*_dffcf = 1
		return nil
	}
	if _gbed, _cbee := _bgeae.(_d.CharData); !_cbee {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bgeae)
	} else {
		switch string(_gbed) {
		case "":
			*_dffcf = 0
		case "\u006e\u006f\u006e\u0065":
			*_dffcf = 1
		case "\u0065\u0071\u0075":
			*_dffcf = 2
		case "\u0067\u0074\u0065":
			*_dffcf = 3
		case "\u006c\u0074\u0065":
			*_dffcf = 4
		}
	}
	_bgeae, _fafe = d.Token()
	if _fafe != nil {
		return _fafe
	}
	if _bfcdg, _agga := _bgeae.(_d.EndElement); _agga && _bfcdg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bgeae)
}

type ST_ResizeHandlesStr byte

func (_bebef ST_FunctionType) ValidateWithPath(path string) error {
	switch _bebef {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bebef))
	}
	return nil
}

const (
	ST_TextAnchorVerticalUnset ST_TextAnchorVertical = 0
	ST_TextAnchorVerticalT     ST_TextAnchorVertical = 1
	ST_TextAnchorVerticalMid   ST_TextAnchorVertical = 2
	ST_TextAnchorVerticalB     ST_TextAnchorVertical = 3
)

func (_dgfd ST_TextAnchorHorizontal) String() string {
	switch _dgfd {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0063\u0074\u0072"
	}
	return ""
}
func (_dcdd ST_StartingElement) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ggcg := _d.Attr{}
	_ggcg.Name = name
	switch _dcdd {
	case ST_StartingElementUnset:
		_ggcg.Value = ""
	case ST_StartingElementNode:
		_ggcg.Value = "\u006e\u006f\u0064\u0065"
	case ST_StartingElementTrans:
		_ggcg.Value = "\u0074\u0072\u0061n\u0073"
	}
	return _ggcg, nil
}

type CT_Categories struct{ Cat []*CT_Category }

func (_dbadd ST_DiagramHorizontalAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_dbadd.String(), start)
}

const (
	ST_NodeHorizontalAlignmentUnset ST_NodeHorizontalAlignment = 0
	ST_NodeHorizontalAlignmentL     ST_NodeHorizontalAlignment = 1
	ST_NodeHorizontalAlignmentCtr   ST_NodeHorizontalAlignment = 2
	ST_NodeHorizontalAlignmentR     ST_NodeHorizontalAlignment = 3
)

// Validate validates the StyleDefHdrLst and its children
func (_decfa *StyleDefHdrLst) Validate() error {
	return _decfa.ValidateWithPath("\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048d\u0072\u004c\u0073\u0074")
}
func (_aabab ST_HierBranchStyle) String() string {
	switch _aabab {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0072"
	case 3:
		return "\u0068\u0061\u006e\u0067"
	case 4:
		return "\u0073\u0074\u0064"
	case 5:
		return "\u0069\u006e\u0069\u0074"
	}
	return ""
}

type RelIds struct{ CT_RelIds }

// ValidateWithPath validates the CT_Pt and its children, prefixing error messages with path
func (_eefa *CT_Pt) ValidateWithPath(path string) error {
	if _cdebb := _eefa.ModelIdAttr.ValidateWithPath(path + "\u002f\u004d\u006fd\u0065\u006c\u0049\u0064\u0041\u0074\u0074\u0072"); _cdebb != nil {
		return _cdebb
	}
	if _edded := _eefa.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _edded != nil {
		return _edded
	}
	if _eefa.CxnIdAttr != nil {
		if _gbgd := _eefa.CxnIdAttr.ValidateWithPath(path + "\u002f\u0043\u0078\u006e\u0049\u0064\u0041\u0074\u0074\u0072"); _gbgd != nil {
			return _gbgd
		}
	}
	if _eefa.PrSet != nil {
		if _cfcga := _eefa.PrSet.ValidateWithPath(path + "\u002f\u0050\u0072\u0053\u0065\u0074"); _cfcga != nil {
			return _cfcga
		}
	}
	if _eefa.SpPr != nil {
		if _dcaca := _eefa.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _dcaca != nil {
			return _dcaca
		}
	}
	if _eefa.T != nil {
		if _eaaee := _eefa.T.ValidateWithPath(path + "\u002f\u0054"); _eaaee != nil {
			return _eaaee
		}
	}
	if _eefa.ExtLst != nil {
		if _aacc := _eefa.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _aacc != nil {
			return _aacc
		}
	}
	return nil
}
func (_cfdd ST_VariableType) ValidateWithPath(path string) error {
	switch _cfdd {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cfdd))
	}
	return nil
}
func (_bfage ST_AutoTextRotation) Validate() error { return _bfage.ValidateWithPath("") }
func (_fgb *CT_ColorTransformHeaderLst) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _fgb.ColorsDefHdr != nil {
		_fgdb := _d.StartElement{Name: _d.Name{Local: "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"}}
		for _, _aade := range _fgb.ColorsDefHdr {
			e.EncodeElement(_aade, _fgdb)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_SDName and its children
func (_befbd *CT_SDName) Validate() error {
	return _befbd.ValidateWithPath("\u0043T\u005f\u0053\u0044\u004e\u0061\u006de")
}
func (_edgee *ST_CxnType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_eagg, _acdde := d.Token()
	if _acdde != nil {
		return _acdde
	}
	if _acgad, _gdfdc := _eagg.(_d.EndElement); _gdfdc && _acgad.Name == start.Name {
		*_edgee = 1
		return nil
	}
	if _ggdg, _faaeg := _eagg.(_d.CharData); !_faaeg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eagg)
	} else {
		switch string(_ggdg) {
		case "":
			*_edgee = 0
		case "\u0070\u0061\u0072O\u0066":
			*_edgee = 1
		case "\u0070\u0072\u0065\u0073\u004f\u0066":
			*_edgee = 2
		case "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff":
			*_edgee = 3
		case "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070":
			*_edgee = 4
		}
	}
	_eagg, _acdde = d.Token()
	if _acdde != nil {
		return _acdde
	}
	if _gbea, _cdda := _eagg.(_d.EndElement); _cdda && _gbea.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eagg)
}
func (_eccb ST_AnimOneStr) String() string {
	switch _eccb {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u006f\u006e\u0065"
	case 3:
		return "\u0062\u0072\u0061\u006e\u0063\u0068"
	}
	return ""
}
func (_aefe *CT_Pt) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006do\u0064\u0065\u006c\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", _aefe.ModelIdAttr)})
	if _aefe.TypeAttr != ST_PtTypeUnset {
		_geeef, _bcbd := _aefe.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _bcbd != nil {
			return _bcbd
		}
		start.Attr = append(start.Attr, _geeef)
	}
	if _aefe.CxnIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0078\u006eI\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_aefe.CxnIdAttr)})
	}
	e.EncodeToken(start)
	if _aefe.PrSet != nil {
		_aecc := _d.StartElement{Name: _d.Name{Local: "\u0070\u0072\u0053e\u0074"}}
		e.EncodeElement(_aefe.PrSet, _aecc)
	}
	if _aefe.SpPr != nil {
		_aagb := _d.StartElement{Name: _d.Name{Local: "\u0073\u0070\u0050\u0072"}}
		e.EncodeElement(_aefe.SpPr, _aagb)
	}
	if _aefe.T != nil {
		_bfc := _d.StartElement{Name: _d.Name{Local: "\u0074"}}
		e.EncodeElement(_aefe.T, _bfc)
	}
	if _aefe.ExtLst != nil {
		_debd := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_aefe.ExtLst, _debd)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

type CT_BulletEnabled struct{ ValAttr *bool }

func (_gdgac ST_PyramidAccentTextMargin) String() string {
	switch _gdgac {
	case 0:
		return ""
	case 1:
		return "\u0073\u0074\u0065\u0070"
	case 2:
		return "\u0073\u0074\u0061c\u006b"
	}
	return ""
}

type StyleDef struct{ CT_StyleDefinition }

func (_gbbe ST_AnimLvlStr) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dgbe := _d.Attr{}
	_dgbe.Name = name
	switch _gbbe {
	case ST_AnimLvlStrUnset:
		_dgbe.Value = ""
	case ST_AnimLvlStrNone:
		_dgbe.Value = "\u006e\u006f\u006e\u0065"
	case ST_AnimLvlStrLvl:
		_dgbe.Value = "\u006c\u0076\u006c"
	case ST_AnimLvlStrCtr:
		_dgbe.Value = "\u0063\u0074\u0072"
	}
	return _dgbe, nil
}
func (_dcede ST_FlowDirection) Validate() error { return _dcede.ValidateWithPath("") }
func (_adga ST_ChildAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_adga.String(), start)
}

// ValidateWithPath validates the CT_ChildMax and its children, prefixing error messages with path
func (_edd *CT_ChildMax) ValidateWithPath(path string) error {
	if _edd.ValAttr != nil {
		if *_edd.ValAttr < -1 {
			return _c.Errorf("\u0025\u0073/m\u002e\u0056\u0061l\u0041\u0074\u0074\u0072 mu\u0073t \u0062\u0065\u0020\u003e\u003d\u0020\u002d1 \u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, *_edd.ValAttr)
		}
	}
	return nil
}
func (_ffegd *ST_RotationPath) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ffegd = 0
	case "\u006e\u006f\u006e\u0065":
		*_ffegd = 1
	case "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h":
		*_ffegd = 2
	}
	return nil
}
func (_baga ST_ClrAppMethod) ValidateWithPath(path string) error {
	switch _baga {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_baga))
	}
	return nil
}

// ST_PrSetCustVal is a union type
type ST_PrSetCustVal struct {
	ST_Percentage *string
	Int32         *int32
}

func (_fdfb ST_ConnectorRouting) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fdfb.String(), start)
}
func (_bbbe *ST_BoolOperator) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bbbe = 0
	case "\u006e\u006f\u006e\u0065":
		*_bbbe = 1
	case "\u0065\u0071\u0075":
		*_bbbe = 2
	case "\u0067\u0074\u0065":
		*_bbbe = 3
	case "\u006c\u0074\u0065":
		*_bbbe = 4
	}
	return nil
}

// ValidateWithPath validates the CT_Parameter and its children, prefixing error messages with path
func (_cfbcg *CT_Parameter) ValidateWithPath(path string) error {
	if _cfbcg.TypeAttr == ST_ParameterIdUnset {
		return _c.Errorf("\u0025\u0073\u002f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020a\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _fdfdg := _cfbcg.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _fdfdg != nil {
		return _fdfdg
	}
	if _acafe := _cfbcg.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _acafe != nil {
		return _acafe
	}
	return nil
}

type ST_VerticalAlignment byte

func NewCT_ColorTransformHeader() *CT_ColorTransformHeader {
	_fed := &CT_ColorTransformHeader{}
	return _fed
}
func (_agfa *CT_Colors) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _fccg := range start.Attr {
		if _fccg.Name.Local == "\u006d\u0065\u0074\u0068" {
			_agfa.MethAttr.UnmarshalXMLAttr(_fccg)
			continue
		}
		if _fccg.Name.Local == "\u0068\u0075\u0065\u0044\u0069\u0072" {
			_agfa.HueDirAttr.UnmarshalXMLAttr(_fccg)
			continue
		}
	}
_ffcd:
	for {
		_edga, _cca := d.Token()
		if _cca != nil {
			return _cca
		}
		switch _eabf := _edga.(type) {
		case _d.StartElement:
			switch _eabf.Name {
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0043\u006f\u006co\u0072\u0073 \u0025\u0076", _eabf.Name)
				if _ggab := d.Skip(); _ggab != nil {
					return _ggab
				}
			}
		case _d.EndElement:
			break _ffcd
		case _d.CharData:
		}
	}
	return nil
}

const (
	ST_ClrAppMethodUnset  ST_ClrAppMethod = 0
	ST_ClrAppMethodSpan   ST_ClrAppMethod = 1
	ST_ClrAppMethodCycle  ST_ClrAppMethod = 2
	ST_ClrAppMethodRepeat ST_ClrAppMethod = 3
)

func (_degf ST_VariableType) String() string {
	switch _degf {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"
	case 3:
		return "\u0063\u0068\u004da\u0078"
	case 4:
		return "\u0063\u0068\u0050\u0072\u0065\u0066"
	case 5:
		return "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064"
	case 6:
		return "\u0064\u0069\u0072"
	case 7:
		return "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"
	case 8:
		return "\u0061n\u0069\u006d\u004f\u006e\u0065"
	case 9:
		return "\u0061n\u0069\u006d\u004c\u0076\u006c"
	case 10:
		return "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"
	}
	return ""
}
func (_aggab ST_ConnectorDimension) Validate() error { return _aggab.ValidateWithPath("") }
func (_acbad ST_ChildDirection) ValidateWithPath(path string) error {
	switch _acbad {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_acbad))
	}
	return nil
}

type ST_BoolOperator byte
type CT_PresentationOf struct {
	ExtLst            *_ca.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

func (_egdc ST_ClrAppMethod) Validate() error { return _egdc.ValidateWithPath("") }
func (_agde *CT_StyleDefinitionHeaderLst) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _agde.StyleDefHdr != nil {
		_dddf := _d.StartElement{Name: _d.Name{Local: "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"}}
		for _, _fcgd := range _agde.StyleDefHdr {
			e.EncodeElement(_fcgd, _dddf)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_gegaa ST_OutputShapeType) ValidateWithPath(path string) error {
	switch _gegaa {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gegaa))
	}
	return nil
}

// ValidateWithPath validates the CT_Description and its children, prefixing error messages with path
func (_bcc *CT_Description) ValidateWithPath(path string) error { return nil }
func (_baege *CT_When) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_baege.FuncAttr = ST_FunctionType(1)
	_baege.OpAttr = ST_FunctionOperator(1)
	for _, _eabd := range start.Attr {
		if _eabd.Name.Local == "\u006e\u0061\u006d\u0065" {
			_bfbg, _cabgd := _eabd.Value, error(nil)
			if _cabgd != nil {
				return _cabgd
			}
			_baege.NameAttr = &_bfbg
			continue
		}
		if _eabd.Name.Local == "\u0061\u0072\u0067" {
			_gdca, _faffd := ParseUnionST_FunctionArgument(_eabd.Value)
			if _faffd != nil {
				return _faffd
			}
			_baege.ArgAttr = &_gdca
			continue
		}
		if _eabd.Name.Local == "\u0076\u0061\u006c" {
			_efde, _bggeb := ParseUnionST_FunctionValue(_eabd.Value)
			if _bggeb != nil {
				return _bggeb
			}
			_baege.ValAttr = _efde
			continue
		}
		if _eabd.Name.Local == "\u0066\u0075\u006e\u0063" {
			_baege.FuncAttr.UnmarshalXMLAttr(_eabd)
			continue
		}
		if _eabd.Name.Local == "\u006f\u0070" {
			_baege.OpAttr.UnmarshalXMLAttr(_eabd)
			continue
		}
		if _eabd.Name.Local == "\u0061\u0078\u0069\u0073" {
			_efef, _ccaa := ParseSliceST_AxisTypes(_eabd.Value)
			if _ccaa != nil {
				return _ccaa
			}
			_baege.AxisAttr = &_efef
			continue
		}
		if _eabd.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_caagf, _cgab := ParseSliceST_ElementTypes(_eabd.Value)
			if _cgab != nil {
				return _cgab
			}
			_baege.PtTypeAttr = &_caagf
			continue
		}
		if _eabd.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_dgece, _eeac := ParseSliceST_Booleans(_eabd.Value)
			if _eeac != nil {
				return _eeac
			}
			_baege.HideLastTransAttr = &_dgece
			continue
		}
		if _eabd.Name.Local == "\u0073\u0074" {
			_gddd, _bgbc := ParseSliceST_Ints(_eabd.Value)
			if _bgbc != nil {
				return _bgbc
			}
			_baege.StAttr = &_gddd
			continue
		}
		if _eabd.Name.Local == "\u0063\u006e\u0074" {
			_bdae, _egcf := ParseSliceST_UnsignedInts(_eabd.Value)
			if _egcf != nil {
				return _egcf
			}
			_baege.CntAttr = &_bdae
			continue
		}
		if _eabd.Name.Local == "\u0073\u0074\u0065\u0070" {
			_age, _gdbed := ParseSliceST_Ints(_eabd.Value)
			if _gdbed != nil {
				return _gdbed
			}
			_baege.StepAttr = &_age
			continue
		}
	}
_egfcd:
	for {
		_ffca, _fcgcd := d.Token()
		if _fcgcd != nil {
			return _fcgcd
		}
		switch _gdgf := _ffca.(type) {
		case _d.StartElement:
			switch _gdgf.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_bgae := NewCT_Algorithm()
				if _aggc := d.DecodeElement(_bgae, &_gdgf); _aggc != nil {
					return _aggc
				}
				_baege.Alg = append(_baege.Alg, _bgae)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_dcdbb := NewCT_Shape()
				if _gaefa := d.DecodeElement(_dcdbb, &_gdgf); _gaefa != nil {
					return _gaefa
				}
				_baege.Shape = append(_baege.Shape, _dcdbb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_abeb := NewCT_PresentationOf()
				if _cgad := d.DecodeElement(_abeb, &_gdgf); _cgad != nil {
					return _cgad
				}
				_baege.PresOf = append(_baege.PresOf, _abeb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_bfebe := NewCT_Constraints()
				if _cegba := d.DecodeElement(_bfebe, &_gdgf); _cegba != nil {
					return _cegba
				}
				_baege.ConstrLst = append(_baege.ConstrLst, _bfebe)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_bfgb := NewCT_Rules()
				if _fcac := d.DecodeElement(_bfgb, &_gdgf); _fcac != nil {
					return _fcac
				}
				_baege.RuleLst = append(_baege.RuleLst, _bfgb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_acdeb := NewCT_ForEach()
				if _eabc := d.DecodeElement(_acdeb, &_gdgf); _eabc != nil {
					return _eabc
				}
				_baege.ForEach = append(_baege.ForEach, _acdeb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_cgcd := NewCT_LayoutNode()
				if _abad := d.DecodeElement(_cgcd, &_gdgf); _abad != nil {
					return _abad
				}
				_baege.LayoutNode = append(_baege.LayoutNode, _cgcd)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_ddfc := NewCT_Choose()
				if _dgcde := d.DecodeElement(_ddfc, &_gdgf); _dgcde != nil {
					return _dgcde
				}
				_baege.Choose = append(_baege.Choose, _ddfc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_accb := _ca.NewCT_OfficeArtExtensionList()
				if _febdf := d.DecodeElement(_accb, &_gdgf); _febdf != nil {
					return _febdf
				}
				_baege.ExtLst = append(_baege.ExtLst, _accb)
			default:
				_af.Log("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0057\u0068\u0065\u006e\u0020\u0025\u0076", _gdgf.Name)
				if _afga := d.Skip(); _afga != nil {
					return _afga
				}
			}
		case _d.EndElement:
			break _egfcd
		case _d.CharData:
		}
	}
	return nil
}
func (_acgbc *ST_TextDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_acgbc = 0
	case "\u0066\u0072\u006fm\u0054":
		*_acgbc = 1
	case "\u0066\u0072\u006fm\u0042":
		*_acgbc = 2
	}
	return nil
}
func (_bcaaf ST_BoolOperator) Validate() error { return _bcaaf.ValidateWithPath("") }

// Validate validates the CT_TextProps and its children
func (_dagg *CT_TextProps) Validate() error {
	return _dagg.ValidateWithPath("\u0043\u0054\u005fT\u0065\u0078\u0074\u0050\u0072\u006f\u0070\u0073")
}
func (_deb *CT_AnimLvl) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _deb.ValAttr != ST_AnimLvlStrUnset {
		_bca, _fcg := _deb.ValAttr.MarshalXMLAttr(_d.Name{Local: "\u0076\u0061\u006c"})
		if _fcg != nil {
			return _fcg
		}
		start.Attr = append(start.Attr, _bca)
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_fdfcc ST_VerticalAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_fagb := _d.Attr{}
	_fagb.Name = name
	switch _fdfcc {
	case ST_VerticalAlignmentUnset:
		_fagb.Value = ""
	case ST_VerticalAlignmentT:
		_fagb.Value = "\u0074"
	case ST_VerticalAlignmentMid:
		_fagb.Value = "\u006d\u0069\u0064"
	case ST_VerticalAlignmentB:
		_fagb.Value = "\u0062"
	case ST_VerticalAlignmentNone:
		_fagb.Value = "\u006e\u006f\u006e\u0065"
	}
	return _fagb, nil
}
func (_efgd ST_FunctionValue) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _efgd.Int32 != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", *_efgd.Int32)))
	}
	if _efgd.Bool != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", _efffd(*_efgd.Bool))))
	}
	if _efgd.ST_Direction != ST_DirectionUnset {
		e.EncodeToken(_d.CharData(_efgd.ST_Direction.String()))
	}
	if _efgd.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		e.EncodeToken(_d.CharData(_efgd.ST_HierBranchStyle.String()))
	}
	if _efgd.ST_AnimOneStr != ST_AnimOneStrUnset {
		e.EncodeToken(_d.CharData(_efgd.ST_AnimOneStr.String()))
	}
	if _efgd.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		e.EncodeToken(_d.CharData(_efgd.ST_AnimLvlStr.String()))
	}
	if _efgd.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		e.EncodeToken(_d.CharData(_efgd.ST_ResizeHandlesStr.String()))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}
func (_cbacf ST_FunctionType) Validate() error { return _cbacf.ValidateWithPath("") }

type ST_CxnType byte

// ValidateWithPath validates the CT_SDCategories and its children, prefixing error messages with path
func (_agbeg *CT_SDCategories) ValidateWithPath(path string) error {
	for _adedab, _adec := range _agbeg.Cat {
		if _bfcd := _adec.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0061\u0074\u005b\u0025\u0064\u005d", path, _adedab)); _bfcd != nil {
			return _bfcd
		}
	}
	return nil
}

type ST_ElementTypes []ST_ElementType

func (_fbac ST_HueDir) String() string {
	switch _fbac {
	case 0:
		return ""
	case 1:
		return "\u0063\u0077"
	case 2:
		return "\u0063\u0063\u0077"
	}
	return ""
}

const (
	ST_AutoTextRotationUnset ST_AutoTextRotation = 0
	ST_AutoTextRotationNone  ST_AutoTextRotation = 1
	ST_AutoTextRotationUpr   ST_AutoTextRotation = 2
	ST_AutoTextRotationGrav  ST_AutoTextRotation = 3
)

// ValidateWithPath validates the CT_CTCategory and its children, prefixing error messages with path
func (_acg *CT_CTCategory) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the CT_CTName and its children, prefixing error messages with path
func (_gcfc *CT_CTName) ValidateWithPath(path string) error { return nil }
func NewCT_DataModel() *CT_DataModel {
	_bgbd := &CT_DataModel{}
	_bgbd.PtLst = NewCT_PtList()
	return _bgbd
}
func (_acacg ST_ElementType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_afeb := _d.Attr{}
	_afeb.Name = name
	switch _acacg {
	case ST_ElementTypeUnset:
		_afeb.Value = ""
	case ST_ElementTypeAll:
		_afeb.Value = "\u0061\u006c\u006c"
	case ST_ElementTypeDoc:
		_afeb.Value = "\u0064\u006f\u0063"
	case ST_ElementTypeNode:
		_afeb.Value = "\u006e\u006f\u0064\u0065"
	case ST_ElementTypeNorm:
		_afeb.Value = "\u006e\u006f\u0072\u006d"
	case ST_ElementTypeNonNorm:
		_afeb.Value = "\u006eo\u006e\u004e\u006f\u0072\u006d"
	case ST_ElementTypeAsst:
		_afeb.Value = "\u0061\u0073\u0073\u0074"
	case ST_ElementTypeNonAsst:
		_afeb.Value = "\u006eo\u006e\u0041\u0073\u0073\u0074"
	case ST_ElementTypeParTrans:
		_afeb.Value = "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case ST_ElementTypePres:
		_afeb.Value = "\u0070\u0072\u0065\u0073"
	case ST_ElementTypeSibTrans:
		_afeb.Value = "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return _afeb, nil
}
func (_fbbgb *ST_ConnectorPoint) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fbbgb = 0
	case "\u0061\u0075\u0074\u006f":
		*_fbbgb = 1
	case "\u0062\u0043\u0074\u0072":
		*_fbbgb = 2
	case "\u0063\u0074\u0072":
		*_fbbgb = 3
	case "\u006d\u0069\u0064\u004c":
		*_fbbgb = 4
	case "\u006d\u0069\u0064\u0052":
		*_fbbgb = 5
	case "\u0074\u0043\u0074\u0072":
		*_fbbgb = 6
	case "\u0062\u004c":
		*_fbbgb = 7
	case "\u0062\u0052":
		*_fbbgb = 8
	case "\u0074\u004c":
		*_fbbgb = 9
	case "\u0074\u0052":
		*_fbbgb = 10
	case "\u0072\u0061\u0064\u0069\u0061\u006c":
		*_fbbgb = 11
	}
	return nil
}
func (_eebag ST_VerticalAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_eebag.String(), start)
}

type CT_StyleDefinitionHeader struct {
	UniqueIdAttr string
	MinVerAttr   *string
	ResIdAttr    *int32
	Title        []*CT_SDName
	Desc         []*CT_SDDescription
	CatLst       *CT_SDCategories
	ExtLst       *_ca.CT_OfficeArtExtensionList
}

func (_bfgfc *RelIds) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bfgfc.CT_RelIds = *NewCT_RelIds()
	for _, _gbfbd := range start.Attr {
		if _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u0064\u006d" || _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u0064\u006d" {
			_fbbc, _gaca := _gbfbd.Value, error(nil)
			if _gaca != nil {
				return _gaca
			}
			_bfgfc.DmAttr = _fbbc
			continue
		}
		if _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u006c\u006f" || _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u006c\u006f" {
			_caed, _gfff := _gbfbd.Value, error(nil)
			if _gfff != nil {
				return _gfff
			}
			_bfgfc.LoAttr = _caed
			continue
		}
		if _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u0071\u0073" || _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u0071\u0073" {
			_gffe, _agef := _gbfbd.Value, error(nil)
			if _agef != nil {
				return _agef
			}
			_bfgfc.QsAttr = _gffe
			continue
		}
		if _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u0063\u0073" || _gbfbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _gbfbd.Name.Local == "\u0063\u0073" {
			_eaaaf, _geagf := _gbfbd.Value, error(nil)
			if _geagf != nil {
				return _geagf
			}
			_bfgfc.CsAttr = _eaaaf
			continue
		}
	}
	for {
		_aacf, _efcdbg := d.Token()
		if _efcdbg != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0052e\u006c\u0049d\u0073\u003a\u0020\u0025\u0073", _efcdbg)
		}
		if _dcfe, _dacb := _aacf.(_d.EndElement); _dacb && _dcfe.Name == start.Name {
			break
		}
	}
	return nil
}
func (_acdae *DataModel) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_acdae.CT_DataModel = *NewCT_DataModel()
_adecg:
	for {
		_cfgad, _acfccb := d.Token()
		if _acfccb != nil {
			return _acfccb
		}
		switch _ccaf := _cfgad.(type) {
		case _d.StartElement:
			switch _ccaf.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0074\u004cs\u0074"}:
				if _cbdb := d.DecodeElement(_acdae.PtLst, &_ccaf); _cbdb != nil {
					return _cbdb
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0078\u006e\u004c\u0073\u0074"}:
				_acdae.CxnLst = NewCT_CxnList()
				if _eacad := d.DecodeElement(_acdae.CxnLst, &_ccaf); _eacad != nil {
					return _eacad
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0062\u0067"}:
				_acdae.Bg = _ca.NewCT_BackgroundFormatting()
				if _afcgf := d.DecodeElement(_acdae.Bg, &_ccaf); _afcgf != nil {
					return _afcgf
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0077\u0068\u006fl\u0065"}:
				_acdae.Whole = _ca.NewCT_WholeE2oFormatting()
				if _acbf := d.DecodeElement(_acdae.Whole, &_ccaf); _acbf != nil {
					return _acbf
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_acdae.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _affec := d.DecodeElement(_acdae.ExtLst, &_ccaf); _affec != nil {
					return _affec
				}
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0044\u0061\u0074\u0061\u004d\u006fd\u0065\u006c \u0025\u0076", _ccaf.Name)
				if _bbafd := d.Skip(); _bbafd != nil {
					return _bbafd
				}
			}
		case _d.EndElement:
			break _adecg
		case _d.CharData:
		}
	}
	return nil
}
func (_ecbdb ST_NodeVerticalAlignment) ValidateWithPath(path string) error {
	switch _ecbdb {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ecbdb))
	}
	return nil
}
func (_cbabc ST_FallbackDimension) ValidateWithPath(path string) error {
	switch _cbabc {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cbabc))
	}
	return nil
}

type CT_Direction struct{ ValAttr ST_Direction }

func (_bbbec ST_ChildOrderType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gfaa := _d.Attr{}
	_gfaa.Name = name
	switch _bbbec {
	case ST_ChildOrderTypeUnset:
		_gfaa.Value = ""
	case ST_ChildOrderTypeB:
		_gfaa.Value = "\u0062"
	case ST_ChildOrderTypeT:
		_gfaa.Value = "\u0074"
	}
	return _gfaa, nil
}
func (_bcdc ST_ArrowheadStyle) String() string {
	switch _bcdc {
	case 0:
		return ""
	case 1:
		return "\u0061\u0075\u0074\u006f"
	case 2:
		return "\u0061\u0072\u0072"
	case 3:
		return "\u006e\u006f\u0041r\u0072"
	}
	return ""
}
func (_dbgf ST_TextDirection) ValidateWithPath(path string) error {
	switch _dbgf {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dbgf))
	}
	return nil
}
func (_gbfcd *ST_ConnectorDimension) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fefa, _fade := d.Token()
	if _fade != nil {
		return _fade
	}
	if _afce, _eedb := _fefa.(_d.EndElement); _eedb && _afce.Name == start.Name {
		*_gbfcd = 1
		return nil
	}
	if _dfab, _dfeg := _fefa.(_d.CharData); !_dfeg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fefa)
	} else {
		switch string(_dfab) {
		case "":
			*_gbfcd = 0
		case "\u0031\u0044":
			*_gbfcd = 1
		case "\u0032\u0044":
			*_gbfcd = 2
		case "\u0063\u0075\u0073\u0074":
			*_gbfcd = 3
		}
	}
	_fefa, _fade = d.Token()
	if _fade != nil {
		return _fade
	}
	if _ffefe, _ecfa := _fefa.(_d.EndElement); _ecfa && _ffefe.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fefa)
}
func (_acbfd *StyleDef) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066"
	return _acbfd.CT_StyleDefinition.MarshalXML(e, start)
}
func (_abfeg ST_Breakpoint) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_abfeg.String(), start)
}
func (_ccbcf ST_NodeVerticalAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_fgcf := _d.Attr{}
	_fgcf.Name = name
	switch _ccbcf {
	case ST_NodeVerticalAlignmentUnset:
		_fgcf.Value = ""
	case ST_NodeVerticalAlignmentT:
		_fgcf.Value = "\u0074"
	case ST_NodeVerticalAlignmentMid:
		_fgcf.Value = "\u006d\u0069\u0064"
	case ST_NodeVerticalAlignmentB:
		_fgcf.Value = "\u0062"
	}
	return _fgcf, nil
}

type CT_ResizeHandles struct{ ValAttr ST_ResizeHandlesStr }

// ValidateWithPath validates the CT_StyleDefinition and its children, prefixing error messages with path
func (_fbcd *CT_StyleDefinition) ValidateWithPath(path string) error {
	for _ccca, _afbc := range _fbcd.Title {
		if _cacae := _afbc.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _ccca)); _cacae != nil {
			return _cacae
		}
	}
	for _dcfc, _faged := range _fbcd.Desc {
		if _ddgf := _faged.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _dcfc)); _ddgf != nil {
			return _ddgf
		}
	}
	if _fbcd.CatLst != nil {
		if _ggcf := _fbcd.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _ggcf != nil {
			return _ggcf
		}
	}
	if _fbcd.Scene3d != nil {
		if _ebccc := _fbcd.Scene3d.ValidateWithPath(path + "\u002f\u0053\u0063\u0065\u006e\u0065\u0033\u0064"); _ebccc != nil {
			return _ebccc
		}
	}
	for _bfbbg, _fcabd := range _fbcd.StyleLbl {
		if _cdga := _fcabd.ValidateWithPath(_c.Sprintf("\u0025s\u002fS\u0074\u0079\u006c\u0065\u004c\u0062\u006c\u005b\u0025\u0064\u005d", path, _bfbbg)); _cdga != nil {
			return _cdga
		}
	}
	if _fbcd.ExtLst != nil {
		if _dfeaf := _fbcd.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _dfeaf != nil {
			return _dfeaf
		}
	}
	return nil
}

// Validate validates the CT_AnimOne and its children
func (_adg *CT_AnimOne) Validate() error {
	return _adg.ValidateWithPath("\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065")
}
func (_gad *CT_Categories) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _gad.Cat != nil {
		_fgc := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074"}}
		for _, _ead := range _gad.Cat {
			e.EncodeElement(_ead, _fgc)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_eefb *CT_NumericRule) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _eefb.ValAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", *_eefb.ValAttr)})
	}
	if _eefb.FactAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0066\u0061\u0063\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_eefb.FactAttr)})
	}
	if _eefb.MaxAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0061\u0078"}, Value: _c.Sprintf("\u0025\u0076", *_eefb.MaxAttr)})
	}
	if _eefb.TypeAttr != ST_ConstraintTypeUnset {
		_bfbba, _egfc := _eefb.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _egfc != nil {
			return _egfc
		}
		start.Attr = append(start.Attr, _bfbba)
	}
	if _eefb.ForAttr != ST_ConstraintRelationshipUnset {
		_cadgc, _gcge := _eefb.ForAttr.MarshalXMLAttr(_d.Name{Local: "\u0066\u006f\u0072"})
		if _gcge != nil {
			return _gcge
		}
		start.Attr = append(start.Attr, _cadgc)
	}
	if _eefb.ForNameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0066o\u0072\u004e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_eefb.ForNameAttr)})
	}
	if _eefb.PtTypeAttr != ST_ElementTypeUnset {
		_fafg, _fcfcd := _eefb.PtTypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"})
		if _fcfcd != nil {
			return _fcfcd
		}
		start.Attr = append(start.Attr, _fafg)
	}
	e.EncodeToken(start)
	if _eefb.ExtLst != nil {
		_affc := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_eefb.ExtLst, _affc)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_gde *CT_BulletEnabled) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _gde.ValAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_gde.ValAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_cdagg *ST_VariableType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_cdagg = 0
	case "\u006e\u006f\u006e\u0065":
		*_cdagg = 1
	case "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074":
		*_cdagg = 2
	case "\u0063\u0068\u004da\u0078":
		*_cdagg = 3
	case "\u0063\u0068\u0050\u0072\u0065\u0066":
		*_cdagg = 4
	case "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064":
		*_cdagg = 5
	case "\u0064\u0069\u0072":
		*_cdagg = 6
	case "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068":
		*_cdagg = 7
	case "\u0061n\u0069\u006d\u004f\u006e\u0065":
		*_cdagg = 8
	case "\u0061n\u0069\u006d\u004c\u0076\u006c":
		*_cdagg = 9
	case "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073":
		*_cdagg = 10
	}
	return nil
}

// Validate validates the CT_SampleData and its children
func (_beac *CT_SampleData) Validate() error {
	return _beac.ValidateWithPath("\u0043\u0054\u005f\u0053\u0061\u006d\u0070\u006c\u0065\u0044\u0061\u0074\u0061")
}

type CT_CTCategory struct {
	TypeAttr string
	PriAttr  uint32
}

const (
	ST_TextDirectionUnset ST_TextDirection = 0
	ST_TextDirectionFromT ST_TextDirection = 1
	ST_TextDirectionFromB ST_TextDirection = 2
)

// ValidateWithPath validates the CT_DataModel and its children, prefixing error messages with path
func (_fbbac *CT_DataModel) ValidateWithPath(path string) error {
	if _eec := _fbbac.PtLst.ValidateWithPath(path + "\u002f\u0050\u0074\u004c\u0073\u0074"); _eec != nil {
		return _eec
	}
	if _fbbac.CxnLst != nil {
		if _acb := _fbbac.CxnLst.ValidateWithPath(path + "\u002fC\u0078\u006e\u004c\u0073\u0074"); _acb != nil {
			return _acb
		}
	}
	if _fbbac.Bg != nil {
		if _fabd := _fbbac.Bg.ValidateWithPath(path + "\u002f\u0042\u0067"); _fabd != nil {
			return _fabd
		}
	}
	if _fbbac.Whole != nil {
		if _acaf := _fbbac.Whole.ValidateWithPath(path + "\u002f\u0057\u0068\u006f\u006c\u0065"); _acaf != nil {
			return _acaf
		}
	}
	if _fbbac.ExtLst != nil {
		if _effc := _fbbac.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _effc != nil {
			return _effc
		}
	}
	return nil
}
func (_gcba *ST_ArrowheadStyle) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gcba = 0
	case "\u0061\u0075\u0074\u006f":
		*_gcba = 1
	case "\u0061\u0072\u0072":
		*_gcba = 2
	case "\u006e\u006f\u0041r\u0072":
		*_gcba = 3
	}
	return nil
}
func (_dfdg ST_FallbackDimension) Validate() error { return _dfdg.ValidateWithPath("") }
func (_dbada ST_AlgorithmType) Validate() error    { return _dbada.ValidateWithPath("") }
func (_gadeda *ST_PtType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gadeda = 0
	case "\u006e\u006f\u0064\u0065":
		*_gadeda = 1
	case "\u0061\u0073\u0073\u0074":
		*_gadeda = 2
	case "\u0064\u006f\u0063":
		*_gadeda = 3
	case "\u0070\u0072\u0065\u0073":
		*_gadeda = 4
	case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
		*_gadeda = 5
	case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
		*_gadeda = 6
	}
	return nil
}
func (_abfg *CT_OrgChart) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gdfa := range start.Attr {
		if _gdfa.Name.Local == "\u0076\u0061\u006c" {
			_efbg, _efae := _ae.ParseBool(_gdfa.Value)
			if _efae != nil {
				return _efae
			}
			_abfg.ValAttr = &_efbg
			continue
		}
	}
	for {
		_egbfc, _fgbd := d.Token()
		if _fgbd != nil {
			return _c.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0072\u0067\u0043\u0068\u0061\u0072\u0074\u003a\u0020\u0025\u0073", _fgbd)
		}
		if _cdgf, _dbge := _egbfc.(_d.EndElement); _dbge && _cdgf.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ccbcfg ST_Offset) ValidateWithPath(path string) error {
	switch _ccbcfg {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ccbcfg))
	}
	return nil
}

const (
	ST_ConstraintRelationshipUnset ST_ConstraintRelationship = 0
	ST_ConstraintRelationshipSelf  ST_ConstraintRelationship = 1
	ST_ConstraintRelationshipCh    ST_ConstraintRelationship = 2
	ST_ConstraintRelationshipDes   ST_ConstraintRelationship = 3
)

// ValidateWithPath validates the StyleDefHdr and its children, prefixing error messages with path
func (_adaef *StyleDefHdr) ValidateWithPath(path string) error {
	if _dedd := _adaef.CT_StyleDefinitionHeader.ValidateWithPath(path); _dedd != nil {
		return _dedd
	}
	return nil
}
func NewCT_LayoutNode() *CT_LayoutNode { _gcgb := &CT_LayoutNode{}; return _gcgb }
func (_afgc *ST_HierBranchStyle) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fddd, _cfedb := d.Token()
	if _cfedb != nil {
		return _cfedb
	}
	if _dcgd, _ecef := _fddd.(_d.EndElement); _ecef && _dcgd.Name == start.Name {
		*_afgc = 1
		return nil
	}
	if _cdgfb, _acbfda := _fddd.(_d.CharData); !_acbfda {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fddd)
	} else {
		switch string(_cdgfb) {
		case "":
			*_afgc = 0
		case "\u006c":
			*_afgc = 1
		case "\u0072":
			*_afgc = 2
		case "\u0068\u0061\u006e\u0067":
			*_afgc = 3
		case "\u0073\u0074\u0064":
			*_afgc = 4
		case "\u0069\u006e\u0069\u0074":
			*_afgc = 5
		}
	}
	_fddd, _cfedb = d.Token()
	if _cfedb != nil {
		return _cfedb
	}
	if _decdf, _gcad := _fddd.(_d.EndElement); _gcad && _decdf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fddd)
}

type CT_ChildMax struct{ ValAttr *int32 }

func (_dffa *ST_AnimOneStr) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ebbb, _dgfac := d.Token()
	if _dgfac != nil {
		return _dgfac
	}
	if _gaecb, _dfbed := _ebbb.(_d.EndElement); _dfbed && _gaecb.Name == start.Name {
		*_dffa = 1
		return nil
	}
	if _accc, _fbea := _ebbb.(_d.CharData); !_fbea {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ebbb)
	} else {
		switch string(_accc) {
		case "":
			*_dffa = 0
		case "\u006e\u006f\u006e\u0065":
			*_dffa = 1
		case "\u006f\u006e\u0065":
			*_dffa = 2
		case "\u0062\u0072\u0061\u006e\u0063\u0068":
			*_dffa = 3
		}
	}
	_ebbb, _dgfac = d.Token()
	if _dgfac != nil {
		return _dgfac
	}
	if _eeeae, _bgebf := _ebbb.(_d.EndElement); _bgebf && _eeeae.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ebbb)
}

type CT_CxnList struct{ Cxn []*CT_Cxn }

func (_edgdf ST_AnimOneStr) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gced := _d.Attr{}
	_gced.Name = name
	switch _edgdf {
	case ST_AnimOneStrUnset:
		_gced.Value = ""
	case ST_AnimOneStrNone:
		_gced.Value = "\u006e\u006f\u006e\u0065"
	case ST_AnimOneStrOne:
		_gced.Value = "\u006f\u006e\u0065"
	case ST_AnimOneStrBranch:
		_gced.Value = "\u0062\u0072\u0061\u006e\u0063\u0068"
	}
	return _gced, nil
}

type CT_When struct {
	NameAttr          *string
	FuncAttr          ST_FunctionType
	ArgAttr           *ST_FunctionArgument
	OpAttr            ST_FunctionOperator
	ValAttr           ST_FunctionValue
	Alg               []*CT_Algorithm
	Shape             []*CT_Shape
	PresOf            []*CT_PresentationOf
	ConstrLst         []*CT_Constraints
	RuleLst           []*CT_Rules
	ForEach           []*CT_ForEach
	LayoutNode        []*CT_LayoutNode
	Choose            []*CT_Choose
	ExtLst            []*_ca.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

const (
	ST_SecondaryChildAlignmentUnset ST_SecondaryChildAlignment = 0
	ST_SecondaryChildAlignmentNone  ST_SecondaryChildAlignment = 1
	ST_SecondaryChildAlignmentT     ST_SecondaryChildAlignment = 2
	ST_SecondaryChildAlignmentB     ST_SecondaryChildAlignment = 3
	ST_SecondaryChildAlignmentL     ST_SecondaryChildAlignment = 4
	ST_SecondaryChildAlignmentR     ST_SecondaryChildAlignment = 5
)

func (_dabcf *CT_SampleData) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _ecgb := range start.Attr {
		if _ecgb.Name.Local == "\u0075\u0073\u0065\u0044\u0065\u0066" {
			_egeg, _dbda := _ae.ParseBool(_ecgb.Value)
			if _dbda != nil {
				return _dbda
			}
			_dabcf.UseDefAttr = &_egeg
			continue
		}
	}
_cagb:
	for {
		_cebc, _bccg := d.Token()
		if _bccg != nil {
			return _bccg
		}
		switch _gdfe := _cebc.(type) {
		case _d.StartElement:
			switch _gdfe.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l"}:
				_dabcf.DataModel = NewCT_DataModel()
				if _aadfg := d.DecodeElement(_dabcf.DataModel, &_gdfe); _aadfg != nil {
					return _aadfg
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053a\u006d\u0070\u006c\u0065\u0044\u0061\u0074\u0061 \u0025\u0076", _gdfe.Name)
				if _degg := d.Skip(); _degg != nil {
					return _degg
				}
			}
		case _d.EndElement:
			break _cagb
		case _d.CharData:
		}
	}
	return nil
}

type CT_OrgChart struct{ ValAttr *bool }

func (_dggg ST_AxisType) ValidateWithPath(path string) error {
	switch _dggg {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dggg))
	}
	return nil
}
func (_ageg ST_ConstraintType) Validate() error { return _ageg.ValidateWithPath("") }
func (_dgge ST_ElementType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_dgge.String(), start)
}

const (
	ST_ConnectorDimensionUnset ST_ConnectorDimension = 0
	ST_ConnectorDimension1D    ST_ConnectorDimension = 1
	ST_ConnectorDimension2D    ST_ConnectorDimension = 2
	ST_ConnectorDimensionCust  ST_ConnectorDimension = 3
)

func (_ecgc *CT_Constraint) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _ecgc.OpAttr != ST_BoolOperatorUnset {
		_bfag, _ccad := _ecgc.OpAttr.MarshalXMLAttr(_d.Name{Local: "\u006f\u0070"})
		if _ccad != nil {
			return _ccad
		}
		start.Attr = append(start.Attr, _bfag)
	}
	if _ecgc.ValAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", *_ecgc.ValAttr)})
	}
	if _ecgc.FactAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0066\u0061\u0063\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_ecgc.FactAttr)})
	}
	if _ecgc.TypeAttr != ST_ConstraintTypeUnset {
		_adab, _dfd := _ecgc.TypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _dfd != nil {
			return _dfd
		}
		start.Attr = append(start.Attr, _adab)
	}
	if _ecgc.ForAttr != ST_ConstraintRelationshipUnset {
		_fbf, _aef := _ecgc.ForAttr.MarshalXMLAttr(_d.Name{Local: "\u0066\u006f\u0072"})
		if _aef != nil {
			return _aef
		}
		start.Attr = append(start.Attr, _fbf)
	}
	if _ecgc.ForNameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0066o\u0072\u004e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_ecgc.ForNameAttr)})
	}
	if _ecgc.PtTypeAttr != ST_ElementTypeUnset {
		_ggd, _ecaa := _ecgc.PtTypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"})
		if _ecaa != nil {
			return _ecaa
		}
		start.Attr = append(start.Attr, _ggd)
	}
	if _ecgc.RefTypeAttr != ST_ConstraintTypeUnset {
		_acfc, _beeb := _ecgc.RefTypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0072e\u0066\u0054\u0079\u0070\u0065"})
		if _beeb != nil {
			return _beeb
		}
		start.Attr = append(start.Attr, _acfc)
	}
	if _ecgc.RefForAttr != ST_ConstraintRelationshipUnset {
		_bda, _dagb := _ecgc.RefForAttr.MarshalXMLAttr(_d.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072"})
		if _dagb != nil {
			return _dagb
		}
		start.Attr = append(start.Attr, _bda)
	}
	if _ecgc.RefForNameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_ecgc.RefForNameAttr)})
	}
	if _ecgc.RefPtTypeAttr != ST_ElementTypeUnset {
		_daa, _bedg := _ecgc.RefPtTypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e"})
		if _bedg != nil {
			return _bedg
		}
		start.Attr = append(start.Attr, _daa)
	}
	e.EncodeToken(start)
	if _ecgc.ExtLst != nil {
		_fff := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_ecgc.ExtLst, _fff)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_ecggg ST_PrSetCustVal) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _ecggg.ST_Percentage != nil {
		e.EncodeToken(_d.CharData(*_ecggg.ST_Percentage))
	}
	if _ecggg.Int32 != nil {
		e.EncodeToken(_d.CharData(_c.Sprintf("\u0025\u0064", *_ecggg.Int32)))
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}
func (_fbaa *ST_DiagramHorizontalAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ecca, _cbacb := d.Token()
	if _cbacb != nil {
		return _cbacb
	}
	if _fabfg, _gcdd := _ecca.(_d.EndElement); _gcdd && _fabfg.Name == start.Name {
		*_fbaa = 1
		return nil
	}
	if _adda, _adffd := _ecca.(_d.CharData); !_adffd {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ecca)
	} else {
		switch string(_adda) {
		case "":
			*_fbaa = 0
		case "\u006c":
			*_fbaa = 1
		case "\u0063\u0074\u0072":
			*_fbaa = 2
		case "\u0072":
			*_fbaa = 3
		case "\u006e\u006f\u006e\u0065":
			*_fbaa = 4
		}
	}
	_ecca, _cbacb = d.Token()
	if _cbacb != nil {
		return _cbacb
	}
	if _efbe, _edfcd := _ecca.(_d.EndElement); _edfcd && _efbe.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ecca)
}
func (_fcabb ST_ModelId) String() string {
	if _fcabb.Int32 != nil {
		return _c.Sprintf("\u0025\u0076", *_fcabb.Int32)
	}
	if _fcabb.ST_Guid != nil {
		return _c.Sprintf("\u0025\u0076", *_fcabb.ST_Guid)
	}
	return ""
}
func (_ebgdc ST_FunctionOperator) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ebgdc.String(), start)
}

// ValidateWithPath validates the CT_StyleDefinitionHeader and its children, prefixing error messages with path
func (_fdad *CT_StyleDefinitionHeader) ValidateWithPath(path string) error {
	for _becfb, _cgcg := range _fdad.Title {
		if _feec := _cgcg.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _becfb)); _feec != nil {
			return _feec
		}
	}
	for _dcgcb, _cabb := range _fdad.Desc {
		if _fggbg := _cabb.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _dcgcb)); _fggbg != nil {
			return _fggbg
		}
	}
	if _fdad.CatLst != nil {
		if _dfggg := _fdad.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _dfggg != nil {
			return _dfggg
		}
	}
	if _fdad.ExtLst != nil {
		if _geag := _fdad.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _geag != nil {
			return _geag
		}
	}
	return nil
}
func (_adcdg ST_ChildOrderType) ValidateWithPath(path string) error {
	switch _adcdg {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_adcdg))
	}
	return nil
}

// ValidateWithPath validates the ColorsDef and its children, prefixing error messages with path
func (_eefbb *ColorsDef) ValidateWithPath(path string) error {
	if _dcag := _eefbb.CT_ColorTransform.ValidateWithPath(path); _dcag != nil {
		return _dcag
	}
	return nil
}
func (_fddcd ST_TextAnchorVertical) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_eebf := _d.Attr{}
	_eebf.Name = name
	switch _fddcd {
	case ST_TextAnchorVerticalUnset:
		_eebf.Value = ""
	case ST_TextAnchorVerticalT:
		_eebf.Value = "\u0074"
	case ST_TextAnchorVerticalMid:
		_eebf.Value = "\u006d\u0069\u0064"
	case ST_TextAnchorVerticalB:
		_eebf.Value = "\u0062"
	}
	return _eebf, nil
}
func (_bcbbf ST_TextAnchorHorizontal) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ddff := _d.Attr{}
	_ddff.Name = name
	switch _bcbbf {
	case ST_TextAnchorHorizontalUnset:
		_ddff.Value = ""
	case ST_TextAnchorHorizontalNone:
		_ddff.Value = "\u006e\u006f\u006e\u0065"
	case ST_TextAnchorHorizontalCtr:
		_ddff.Value = "\u0063\u0074\u0072"
	}
	return _ddff, nil
}
func (_abbac ST_FunctionOperator) String() string {
	switch _abbac {
	case 0:
		return ""
	case 1:
		return "\u0065\u0071\u0075"
	case 2:
		return "\u006e\u0065\u0071"
	case 3:
		return "\u0067\u0074"
	case 4:
		return "\u006c\u0074"
	case 5:
		return "\u0067\u0074\u0065"
	case 6:
		return "\u006c\u0074\u0065"
	}
	return ""
}

type CT_ChildPref struct{ ValAttr *int32 }
type CT_Shape struct {
	RotAttr       *float64
	TypeAttr      *ST_LayoutShapeType
	BlipAttr      *string
	ZOrderOffAttr *int32
	HideGeomAttr  *bool
	LkTxEntryAttr *bool
	BlipPhldrAttr *bool
	AdjLst        *CT_AdjLst
	ExtLst        *_ca.CT_OfficeArtExtensionList
}

func NewCT_CxnList() *CT_CxnList { _dcd := &CT_CxnList{}; return _dcd }
func (_ecedc ST_PyramidAccentPosition) ValidateWithPath(path string) error {
	switch _ecedc {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ecedc))
	}
	return nil
}
func (_gcbf ST_CenterShapeMapping) Validate() error { return _gcbf.ValidateWithPath("") }
func (_efbce *ST_AlgorithmType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ccddb, _ecabe := d.Token()
	if _ecabe != nil {
		return _ecabe
	}
	if _adcb, _baeba := _ccddb.(_d.EndElement); _baeba && _adcb.Name == start.Name {
		*_efbce = 1
		return nil
	}
	if _bdcc, _gffef := _ccddb.(_d.CharData); !_gffef {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ccddb)
	} else {
		switch string(_bdcc) {
		case "":
			*_efbce = 0
		case "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e":
			*_efbce = 1
		case "\u0063\u006f\u006e\u006e":
			*_efbce = 2
		case "\u0063\u0079\u0063l\u0065":
			*_efbce = 3
		case "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd":
			*_efbce = 4
		case "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074":
			*_efbce = 5
		case "\u0070\u0079\u0072\u0061":
			*_efbce = 6
		case "\u006c\u0069\u006e":
			*_efbce = 7
		case "\u0073\u0070":
			*_efbce = 8
		case "\u0074\u0078":
			*_efbce = 9
		case "\u0073\u006e\u0061k\u0065":
			*_efbce = 10
		}
	}
	_ccddb, _ecabe = d.Token()
	if _ecabe != nil {
		return _ecabe
	}
	if _bbgf, _baag := _ccddb.(_d.EndElement); _baag && _bbgf.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ccddb)
}
func (_bfcgf *ST_ContinueDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bfcgf = 0
	case "\u0072\u0065\u0076\u0044\u0069\u0072":
		*_bfcgf = 1
	case "\u0073a\u006d\u0065\u0044\u0069\u0072":
		*_bfcgf = 2
	}
	return nil
}
func (_abfb *CT_CTCategories) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_gfe:
	for {
		_eedg, _gfd := d.Token()
		if _gfd != nil {
			return _gfd
		}
		switch _cfd := _eedg.(type) {
		case _d.StartElement:
			switch _cfd.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074"}:
				_gbf := NewCT_CTCategory()
				if _dge := d.DecodeElement(_gbf, &_cfd); _dge != nil {
					return _dge
				}
				_abfb.Cat = append(_abfb.Cat, _gbf)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u0054\u0043a\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073\u0020\u0025\u0076", _cfd.Name)
				if _egb := d.Skip(); _egb != nil {
					return _egb
				}
			}
		case _d.EndElement:
			break _gfe
		case _d.CharData:
		}
	}
	return nil
}

const (
	ST_VerticalAlignmentUnset ST_VerticalAlignment = 0
	ST_VerticalAlignmentT     ST_VerticalAlignment = 1
	ST_VerticalAlignmentMid   ST_VerticalAlignment = 2
	ST_VerticalAlignmentB     ST_VerticalAlignment = 3
	ST_VerticalAlignmentNone  ST_VerticalAlignment = 4
)

func (_ebaccf ST_PyramidAccentPosition) Validate() error { return _ebaccf.ValidateWithPath("") }

// Validate validates the CT_When and its children
func (_dbce *CT_When) Validate() error {
	return _dbce.ValidateWithPath("\u0043T\u005f\u0057\u0068\u0065\u006e")
}
func (_gcfa *CT_SDCategories) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_eaca:
	for {
		_afec, _cegb := d.Token()
		if _cegb != nil {
			return _cegb
		}
		switch _affe := _afec.(type) {
		case _d.StartElement:
			switch _affe.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074"}:
				_gdbcg := NewCT_SDCategory()
				if _aaeeb := d.DecodeElement(_gdbcg, &_affe); _aaeeb != nil {
					return _aaeeb
				}
				_gcfa.Cat = append(_gcfa.Cat, _gdbcg)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0044\u0043a\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073\u0020\u0025\u0076", _affe.Name)
				if _daaeb := d.Skip(); _daaeb != nil {
					return _daaeb
				}
			}
		case _d.EndElement:
			break _eaca
		case _d.CharData:
		}
	}
	return nil
}
func (_fdfg *CT_StyleLabel) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _edbea := range start.Attr {
		if _edbea.Name.Local == "\u006e\u0061\u006d\u0065" {
			_baae, _ecdf := _edbea.Value, error(nil)
			if _ecdf != nil {
				return _ecdf
			}
			_fdfg.NameAttr = _baae
			continue
		}
	}
_dfdec:
	for {
		_cdab, _dbad := d.Token()
		if _dbad != nil {
			return _dbad
		}
		switch _dcec := _cdab.(type) {
		case _d.StartElement:
			switch _dcec.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}:
				_fdfg.Scene3d = _ca.NewCT_Scene3D()
				if _degd := d.DecodeElement(_fdfg.Scene3d, &_dcec); _degd != nil {
					return _degd
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0070\u0033\u0064"}:
				_fdfg.Sp3d = _ca.NewCT_Shape3D()
				if _bdgeb := d.DecodeElement(_fdfg.Sp3d, &_dcec); _bdgeb != nil {
					return _bdgeb
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0078\u0050\u0072"}:
				_fdfg.TxPr = NewCT_TextProps()
				if _gcbba := d.DecodeElement(_fdfg.TxPr, &_dcec); _gcbba != nil {
					return _gcbba
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079l\u0065"}:
				_fdfg.Style = _ca.NewCT_ShapeStyle()
				if _begc := d.DecodeElement(_fdfg.Style, &_dcec); _begc != nil {
					return _begc
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_fdfg.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _aabg := d.DecodeElement(_fdfg.ExtLst, &_dcec); _aabg != nil {
					return _aabg
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053t\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c \u0025\u0076", _dcec.Name)
				if _edgb := d.Skip(); _edgb != nil {
					return _edgb
				}
			}
		case _d.EndElement:
			break _dfdec
		case _d.CharData:
		}
	}
	return nil
}
func (_ddagc ST_NodeVerticalAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ddagc.String(), start)
}

type ST_AxisType byte

func (_cbdac *ST_PyramidAccentTextMargin) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_cbdac = 0
	case "\u0073\u0074\u0065\u0070":
		*_cbdac = 1
	case "\u0073\u0074\u0061c\u006b":
		*_cbdac = 2
	}
	return nil
}

// Validate validates the CT_ElemPropSet and its children
func (_bccc *CT_ElemPropSet) Validate() error {
	return _bccc.ValidateWithPath("\u0043\u0054\u005f\u0045\u006c\u0065\u006d\u0050\u0072o\u0070\u0053\u0065\u0074")
}
func (_gefd *ST_FunctionArgument) ValidateWithPath(path string) error {
	_ffbee := []string{}
	if _gefd.ST_VariableType != ST_VariableTypeUnset {
		_ffbee = append(_ffbee, "\u0053T\u005fV\u0061\u0072\u0069\u0061\u0062\u006c\u0065\u0054\u0079\u0070\u0065")
	}
	if len(_ffbee) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _ffbee)
	}
	return nil
}

type ST_GrowDirection byte

func (_afdaa ST_AutoTextRotation) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_faeaa := _d.Attr{}
	_faeaa.Name = name
	switch _afdaa {
	case ST_AutoTextRotationUnset:
		_faeaa.Value = ""
	case ST_AutoTextRotationNone:
		_faeaa.Value = "\u006e\u006f\u006e\u0065"
	case ST_AutoTextRotationUpr:
		_faeaa.Value = "\u0075\u0070\u0072"
	case ST_AutoTextRotationGrav:
		_faeaa.Value = "\u0067\u0072\u0061\u0076"
	}
	return _faeaa, nil
}

type StyleDefHdrLst struct{ CT_StyleDefinitionHeaderLst }

func (_dgef *CT_Rules) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _dgef.Rule != nil {
		_ffgd := _d.StartElement{Name: _d.Name{Local: "\u0072\u0075\u006c\u0065"}}
		for _, _gaded := range _dgef.Rule {
			e.EncodeElement(_gaded, _ffgd)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_aefd *ST_TextAnchorVertical) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_aefd = 0
	case "\u0074":
		*_aefd = 1
	case "\u006d\u0069\u0064":
		*_aefd = 2
	case "\u0062":
		*_aefd = 3
	}
	return nil
}
func (_ffbaa ST_DiagramTextAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ffbaa.String(), start)
}
func (_eacbg *ST_AxisType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cbeg, _fggbgd := d.Token()
	if _fggbgd != nil {
		return _fggbgd
	}
	if _ecbge, _dacff := _cbeg.(_d.EndElement); _dacff && _ecbge.Name == start.Name {
		*_eacbg = 1
		return nil
	}
	if _aecb, _aefcf := _cbeg.(_d.CharData); !_aefcf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cbeg)
	} else {
		switch string(_aecb) {
		case "":
			*_eacbg = 0
		case "\u0073\u0065\u006c\u0066":
			*_eacbg = 1
		case "\u0063\u0068":
			*_eacbg = 2
		case "\u0064\u0065\u0073":
			*_eacbg = 3
		case "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf":
			*_eacbg = 4
		case "\u0070\u0061\u0072":
			*_eacbg = 5
		case "\u0061\u006e\u0063s\u0074":
			*_eacbg = 6
		case "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066":
			*_eacbg = 7
		case "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b":
			*_eacbg = 8
		case "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b":
			*_eacbg = 9
		case "\u0066\u006f\u006c\u006c\u006f\u0077":
			*_eacbg = 10
		case "\u0070\u0072\u0065\u0063\u0065\u0064":
			*_eacbg = 11
		case "\u0072\u006f\u006f\u0074":
			*_eacbg = 12
		case "\u006e\u006f\u006e\u0065":
			*_eacbg = 13
		}
	}
	_cbeg, _fggbgd = d.Token()
	if _fggbgd != nil {
		return _fggbgd
	}
	if _ecba, _eeedc := _cbeg.(_d.EndElement); _eeedc && _ecba.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cbeg)
}

// Validate validates the CT_AnimLvl and its children
func (_aaa *CT_AnimLvl) Validate() error {
	return _aaa.ValidateWithPath("\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c")
}
func (_gegdb ST_ChildDirection) String() string {
	switch _gegdb {
	case 0:
		return ""
	case 1:
		return "\u0068\u006f\u0072\u007a"
	case 2:
		return "\u0076\u0065\u0072\u0074"
	}
	return ""
}

// ValidateWithPath validates the CT_SDDescription and its children, prefixing error messages with path
func (_agbf *CT_SDDescription) ValidateWithPath(path string) error { return nil }

type AG_ConstraintRefAttributes struct {
	RefTypeAttr    ST_ConstraintType
	RefForAttr     ST_ConstraintRelationship
	RefForNameAttr *string
	RefPtTypeAttr  ST_ElementType
}

func (_gacb ST_AnimOneStr) ValidateWithPath(path string) error {
	switch _gacb {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gacb))
	}
	return nil
}
func NewCT_PresentationOf() *CT_PresentationOf { _dcbfd := &CT_PresentationOf{}; return _dcbfd }

// ValidateWithPath validates the CT_ForEach and its children, prefixing error messages with path
func (_ebgd *CT_ForEach) ValidateWithPath(path string) error {
	for _dagf, _bgec := range _ebgd.Alg {
		if _eedeg := _bgec.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _dagf)); _eedeg != nil {
			return _eedeg
		}
	}
	for _defba, _edbe := range _ebgd.Shape {
		if _ffba := _edbe.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _defba)); _ffba != nil {
			return _ffba
		}
	}
	for _bbdc, _acac := range _ebgd.PresOf {
		if _adaca := _acac.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _bbdc)); _adaca != nil {
			return _adaca
		}
	}
	for _dede, _cfed := range _ebgd.ConstrLst {
		if _eacg := _cfed.ValidateWithPath(_c.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _dede)); _eacg != nil {
			return _eacg
		}
	}
	for _dgec, _bbdb := range _ebgd.RuleLst {
		if _dcab := _bbdb.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _dgec)); _dcab != nil {
			return _dcab
		}
	}
	for _baeg, _caga := range _ebgd.ForEach {
		if _cdba := _caga.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _baeg)); _cdba != nil {
			return _cdba
		}
	}
	for _dcbf, _cff := range _ebgd.LayoutNode {
		if _bcbga := _cff.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _dcbf)); _bcbga != nil {
			return _bcbga
		}
	}
	for _dgfa, _eeab := range _ebgd.Choose {
		if _febe := _eeab.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _dgfa)); _febe != nil {
			return _febe
		}
	}
	for _gdbg, _dfgf := range _ebgd.ExtLst {
		if _eedebb := _dfgf.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _gdbg)); _eedebb != nil {
			return _eedebb
		}
	}
	return nil
}
func (_fdgec ST_TextBlockDirection) Validate() error { return _fdgec.ValidateWithPath("") }

type CT_DiagramDefinition struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	DefStyleAttr *string
	Title        []*CT_Name
	Desc         []*CT_Description
	CatLst       *CT_Categories
	SampData     *CT_SampleData
	StyleData    *CT_SampleData
	ClrData      *CT_SampleData
	LayoutNode   *CT_LayoutNode
	ExtLst       *_ca.CT_OfficeArtExtensionList
}

func (_dgeef ST_ElementType) ValidateWithPath(path string) error {
	switch _dgeef {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dgeef))
	}
	return nil
}
func (_acaac ST_TextAnchorVertical) String() string {
	switch _acaac {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u006d\u0069\u0064"
	case 3:
		return "\u0062"
	}
	return ""
}

// Validate validates the CT_OrgChart and its children
func (_dgcfg *CT_OrgChart) Validate() error {
	return _dgcfg.ValidateWithPath("C\u0054\u005f\u004f\u0072\u0067\u0043\u0068\u0061\u0072\u0074")
}
func (_ggeg ST_TextAnchorHorizontal) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ggeg.String(), start)
}
func NewCT_BulletEnabled() *CT_BulletEnabled { _add := &CT_BulletEnabled{}; return _add }

// ValidateWithPath validates the CT_AnimLvl and its children, prefixing error messages with path
func (_fdee *CT_AnimLvl) ValidateWithPath(path string) error {
	if _bcd := _fdee.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _bcd != nil {
		return _bcd
	}
	return nil
}
func ParseSliceST_AxisTypes(s string) (ST_AxisTypes, error) { return ST_AxisTypes{}, nil }
func (_gbcff ST_NodeHorizontalAlignment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_gbcff.String(), start)
}

type ST_SecondaryChildAlignment byte

func (_cggdc *CT_DiagramDefinitionHeader) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", _cggdc.UniqueIdAttr)})
	if _cggdc.MinVerAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _c.Sprintf("\u0025\u0076", *_cggdc.MinVerAttr)})
	}
	if _cggdc.DefStyleAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_cggdc.DefStyleAttr)})
	}
	if _cggdc.ResIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u0065\u0073I\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_cggdc.ResIdAttr)})
	}
	e.EncodeToken(start)
	_adde := _d.StartElement{Name: _d.Name{Local: "\u0074\u0069\u0074l\u0065"}}
	for _, _eag := range _cggdc.Title {
		e.EncodeElement(_eag, _adde)
	}
	_aadf := _d.StartElement{Name: _d.Name{Local: "\u0064\u0065\u0073\u0063"}}
	for _, _agfb := range _cggdc.Desc {
		e.EncodeElement(_agfb, _aadf)
	}
	if _cggdc.CatLst != nil {
		_deae := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_cggdc.CatLst, _deae)
	}
	if _cggdc.ExtLst != nil {
		_acaa := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_cggdc.ExtLst, _acaa)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_bebf ST_GrowDirection) String() string {
	switch _bebf {
	case 0:
		return ""
	case 1:
		return "\u0074\u004c"
	case 2:
		return "\u0074\u0052"
	case 3:
		return "\u0062\u004c"
	case 4:
		return "\u0062\u0052"
	}
	return ""
}

const (
	ST_DiagramHorizontalAlignmentUnset ST_DiagramHorizontalAlignment = 0
	ST_DiagramHorizontalAlignmentL     ST_DiagramHorizontalAlignment = 1
	ST_DiagramHorizontalAlignmentCtr   ST_DiagramHorizontalAlignment = 2
	ST_DiagramHorizontalAlignmentR     ST_DiagramHorizontalAlignment = 3
	ST_DiagramHorizontalAlignmentNone  ST_DiagramHorizontalAlignment = 4
)

func (_ffegc ST_ResizeHandlesStr) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_accbd := _d.Attr{}
	_accbd.Name = name
	switch _ffegc {
	case ST_ResizeHandlesStrUnset:
		_accbd.Value = ""
	case ST_ResizeHandlesStrExact:
		_accbd.Value = "\u0065\u0078\u0061c\u0074"
	case ST_ResizeHandlesStrRel:
		_accbd.Value = "\u0072\u0065\u006c"
	}
	return _accbd, nil
}
func (_bbec ST_PyramidAccentPosition) String() string {
	switch _bbec {
	case 0:
		return ""
	case 1:
		return "\u0062\u0065\u0066"
	case 2:
		return "\u0061\u0066\u0074"
	}
	return ""
}

// Validate validates the CT_CTCategories and its children
func (_dcgc *CT_CTCategories) Validate() error {
	return _dcgc.ValidateWithPath("\u0043T\u005fC\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073")
}
func (_cg *AG_ConstraintRefAttributes) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _cg.RefTypeAttr != ST_ConstraintTypeUnset {
		_fa, _bg := _cg.RefTypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0072e\u0066\u0054\u0079\u0070\u0065"})
		if _bg != nil {
			return _bg
		}
		start.Attr = append(start.Attr, _fa)
	}
	if _cg.RefForAttr != ST_ConstraintRelationshipUnset {
		_fg, _aed := _cg.RefForAttr.MarshalXMLAttr(_d.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072"})
		if _aed != nil {
			return _aed
		}
		start.Attr = append(start.Attr, _fg)
	}
	if _cg.RefForNameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_cg.RefForNameAttr)})
	}
	if _cg.RefPtTypeAttr != ST_ElementTypeUnset {
		_ea, _fd := _cg.RefPtTypeAttr.MarshalXMLAttr(_d.Name{Local: "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e"})
		if _fd != nil {
			return _fd
		}
		start.Attr = append(start.Attr, _ea)
	}
	return nil
}

type ST_AnimLvlStr byte

// ValidateWithPath validates the CT_CTStyleLabel and its children, prefixing error messages with path
func (_adff *CT_CTStyleLabel) ValidateWithPath(path string) error {
	if _adff.FillClrLst != nil {
		if _aee := _adff.FillClrLst.ValidateWithPath(path + "/\u0046\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"); _aee != nil {
			return _aee
		}
	}
	if _adff.LinClrLst != nil {
		if _cbe := _adff.LinClrLst.ValidateWithPath(path + "\u002f\u004c\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"); _cbe != nil {
			return _cbe
		}
	}
	if _adff.EffectClrLst != nil {
		if _beg := _adff.EffectClrLst.ValidateWithPath(path + "\u002f\u0045\u0066\u0066\u0065\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"); _beg != nil {
			return _beg
		}
	}
	if _adff.TxLinClrLst != nil {
		if _eged := _adff.TxLinClrLst.ValidateWithPath(path + "\u002f\u0054\u0078L\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"); _eged != nil {
			return _eged
		}
	}
	if _adff.TxFillClrLst != nil {
		if _febg := _adff.TxFillClrLst.ValidateWithPath(path + "\u002f\u0054\u0078\u0046\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"); _febg != nil {
			return _febg
		}
	}
	if _adff.TxEffectClrLst != nil {
		if _fecb := _adff.TxEffectClrLst.ValidateWithPath(path + "\u002fT\u0078E\u0066\u0066\u0065\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"); _fecb != nil {
			return _fecb
		}
	}
	if _adff.ExtLst != nil {
		if _cfg := _adff.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _cfg != nil {
			return _cfg
		}
	}
	return nil
}
func (_fgcdc *CT_StyleLabel) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", _fgcdc.NameAttr)})
	e.EncodeToken(start)
	if _fgcdc.Scene3d != nil {
		_defbaf := _d.StartElement{Name: _d.Name{Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}}
		e.EncodeElement(_fgcdc.Scene3d, _defbaf)
	}
	if _fgcdc.Sp3d != nil {
		_dcfff := _d.StartElement{Name: _d.Name{Local: "\u0073\u0070\u0033\u0064"}}
		e.EncodeElement(_fgcdc.Sp3d, _dcfff)
	}
	if _fgcdc.TxPr != nil {
		_ecec := _d.StartElement{Name: _d.Name{Local: "\u0074\u0078\u0050\u0072"}}
		e.EncodeElement(_fgcdc.TxPr, _ecec)
	}
	if _fgcdc.Style != nil {
		_afda := _d.StartElement{Name: _d.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_fgcdc.Style, _afda)
	}
	if _fgcdc.ExtLst != nil {
		_dbed := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_fgcdc.ExtLst, _dbed)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewCT_AnimOne() *CT_AnimOne { _aeg := &CT_AnimOne{}; return _aeg }
func (_bgfge *ST_SecondaryLinearDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dgbc, _abbdc := d.Token()
	if _abbdc != nil {
		return _abbdc
	}
	if _gbbg, _gadee := _dgbc.(_d.EndElement); _gadee && _gbbg.Name == start.Name {
		*_bgfge = 1
		return nil
	}
	if _bccaf, _efab := _dgbc.(_d.CharData); !_efab {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dgbc)
	} else {
		switch string(_bccaf) {
		case "":
			*_bgfge = 0
		case "\u006e\u006f\u006e\u0065":
			*_bgfge = 1
		case "\u0066\u0072\u006fm\u004c":
			*_bgfge = 2
		case "\u0066\u0072\u006fm\u0052":
			*_bgfge = 3
		case "\u0066\u0072\u006fm\u0054":
			*_bgfge = 4
		case "\u0066\u0072\u006fm\u0042":
			*_bgfge = 5
		}
	}
	_dgbc, _abbdc = d.Token()
	if _abbdc != nil {
		return _abbdc
	}
	if _dgff, _beebb := _dgbc.(_d.EndElement); _beebb && _dgff.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dgbc)
}
func (_afac *ST_VerticalAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_afac = 0
	case "\u0074":
		*_afac = 1
	case "\u006d\u0069\u0064":
		*_afac = 2
	case "\u0062":
		*_afac = 3
	case "\u006e\u006f\u006e\u0065":
		*_afac = 4
	}
	return nil
}

const (
	ST_ConstraintTypeUnset         ST_ConstraintType = 0
	ST_ConstraintTypeNone          ST_ConstraintType = 1
	ST_ConstraintTypeAlignOff      ST_ConstraintType = 2
	ST_ConstraintTypeBegMarg       ST_ConstraintType = 3
	ST_ConstraintTypeBendDist      ST_ConstraintType = 4
	ST_ConstraintTypeBegPad        ST_ConstraintType = 5
	ST_ConstraintTypeB             ST_ConstraintType = 6
	ST_ConstraintTypeBMarg         ST_ConstraintType = 7
	ST_ConstraintTypeBOff          ST_ConstraintType = 8
	ST_ConstraintTypeCtrX          ST_ConstraintType = 9
	ST_ConstraintTypeCtrXOff       ST_ConstraintType = 10
	ST_ConstraintTypeCtrY          ST_ConstraintType = 11
	ST_ConstraintTypeCtrYOff       ST_ConstraintType = 12
	ST_ConstraintTypeConnDist      ST_ConstraintType = 13
	ST_ConstraintTypeDiam          ST_ConstraintType = 14
	ST_ConstraintTypeEndMarg       ST_ConstraintType = 15
	ST_ConstraintTypeEndPad        ST_ConstraintType = 16
	ST_ConstraintTypeH             ST_ConstraintType = 17
	ST_ConstraintTypeHArH          ST_ConstraintType = 18
	ST_ConstraintTypeHOff          ST_ConstraintType = 19
	ST_ConstraintTypeL             ST_ConstraintType = 20
	ST_ConstraintTypeLMarg         ST_ConstraintType = 21
	ST_ConstraintTypeLOff          ST_ConstraintType = 22
	ST_ConstraintTypeR             ST_ConstraintType = 23
	ST_ConstraintTypeRMarg         ST_ConstraintType = 24
	ST_ConstraintTypeROff          ST_ConstraintType = 25
	ST_ConstraintTypePrimFontSz    ST_ConstraintType = 26
	ST_ConstraintTypePyraAcctRatio ST_ConstraintType = 27
	ST_ConstraintTypeSecFontSz     ST_ConstraintType = 28
	ST_ConstraintTypeSibSp         ST_ConstraintType = 29
	ST_ConstraintTypeSecSibSp      ST_ConstraintType = 30
	ST_ConstraintTypeSp            ST_ConstraintType = 31
	ST_ConstraintTypeStemThick     ST_ConstraintType = 32
	ST_ConstraintTypeT             ST_ConstraintType = 33
	ST_ConstraintTypeTMarg         ST_ConstraintType = 34
	ST_ConstraintTypeTOff          ST_ConstraintType = 35
	ST_ConstraintTypeUserA         ST_ConstraintType = 36
	ST_ConstraintTypeUserB         ST_ConstraintType = 37
	ST_ConstraintTypeUserC         ST_ConstraintType = 38
	ST_ConstraintTypeUserD         ST_ConstraintType = 39
	ST_ConstraintTypeUserE         ST_ConstraintType = 40
	ST_ConstraintTypeUserF         ST_ConstraintType = 41
	ST_ConstraintTypeUserG         ST_ConstraintType = 42
	ST_ConstraintTypeUserH         ST_ConstraintType = 43
	ST_ConstraintTypeUserI         ST_ConstraintType = 44
	ST_ConstraintTypeUserJ         ST_ConstraintType = 45
	ST_ConstraintTypeUserK         ST_ConstraintType = 46
	ST_ConstraintTypeUserL         ST_ConstraintType = 47
	ST_ConstraintTypeUserM         ST_ConstraintType = 48
	ST_ConstraintTypeUserN         ST_ConstraintType = 49
	ST_ConstraintTypeUserO         ST_ConstraintType = 50
	ST_ConstraintTypeUserP         ST_ConstraintType = 51
	ST_ConstraintTypeUserQ         ST_ConstraintType = 52
	ST_ConstraintTypeUserR         ST_ConstraintType = 53
	ST_ConstraintTypeUserS         ST_ConstraintType = 54
	ST_ConstraintTypeUserT         ST_ConstraintType = 55
	ST_ConstraintTypeUserU         ST_ConstraintType = 56
	ST_ConstraintTypeUserV         ST_ConstraintType = 57
	ST_ConstraintTypeUserW         ST_ConstraintType = 58
	ST_ConstraintTypeUserX         ST_ConstraintType = 59
	ST_ConstraintTypeUserY         ST_ConstraintType = 60
	ST_ConstraintTypeUserZ         ST_ConstraintType = 61
	ST_ConstraintTypeW             ST_ConstraintType = 62
	ST_ConstraintTypeWArH          ST_ConstraintType = 63
	ST_ConstraintTypeWOff          ST_ConstraintType = 64
)

func (_gcecc ST_BendPoint) Validate() error { return _gcecc.ValidateWithPath("") }
func NewStyleDefHdrLst() *StyleDefHdrLst {
	_daaag := &StyleDefHdrLst{}
	_daaag.CT_StyleDefinitionHeaderLst = *NewCT_StyleDefinitionHeaderLst()
	return _daaag
}
func NewCT_ElemPropSet() *CT_ElemPropSet { _eecge := &CT_ElemPropSet{}; return _eecge }
func (_gdfd *CT_Name) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _addd := range start.Attr {
		if _addd.Name.Local == "\u006c\u0061\u006e\u0067" {
			_dfda, _gabad := _addd.Value, error(nil)
			if _gabad != nil {
				return _gabad
			}
			_gdfd.LangAttr = &_dfda
			continue
		}
		if _addd.Name.Local == "\u0076\u0061\u006c" {
			_cfff, _edgf := _addd.Value, error(nil)
			if _edgf != nil {
				return _edgf
			}
			_gdfd.ValAttr = _cfff
			continue
		}
	}
	for {
		_fddcb, _dffd := d.Token()
		if _dffd != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004e\u0061\u006d\u0065\u003a\u0020\u0025\u0073", _dffd)
		}
		if _gbfd, _aegd := _fddcb.(_d.EndElement); _aegd && _gbfd.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the ColorsDefHdr and its children, prefixing error messages with path
func (_ccdc *ColorsDefHdr) ValidateWithPath(path string) error {
	if _cgfc := _ccdc.CT_ColorTransformHeader.ValidateWithPath(path); _cgfc != nil {
		return _cgfc
	}
	return nil
}
func (_bgdc *CT_ChildPref) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _bgdc.ValAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", *_bgdc.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_egfdg ST_NodeHorizontalAlignment) ValidateWithPath(path string) error {
	switch _egfdg {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_egfdg))
	}
	return nil
}
func (_fgdbb ST_TextBlockDirection) ValidateWithPath(path string) error {
	switch _fgdbb {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fgdbb))
	}
	return nil
}
func (_ebce *CT_ColorTransformHeader) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _cda := range start.Attr {
		if _cda.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_gbb, _ccc := _cda.Value, error(nil)
			if _ccc != nil {
				return _ccc
			}
			_ebce.UniqueIdAttr = _gbb
			continue
		}
		if _cda.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_gcfcf, _agca := _cda.Value, error(nil)
			if _agca != nil {
				return _agca
			}
			_ebce.MinVerAttr = &_gcfcf
			continue
		}
		if _cda.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_agaa, _cee := _ae.ParseInt(_cda.Value, 10, 32)
			if _cee != nil {
				return _cee
			}
			_ggcd := int32(_agaa)
			_ebce.ResIdAttr = &_ggcd
			continue
		}
	}
_eaff:
	for {
		_dace, _fggg := d.Token()
		if _fggg != nil {
			return _fggg
		}
		switch _fga := _dace.(type) {
		case _d.StartElement:
			switch _fga.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_cef := NewCT_CTName()
				if _bcdb := d.DecodeElement(_cef, &_fga); _bcdb != nil {
					return _bcdb
				}
				_ebce.Title = append(_ebce.Title, _cef)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_cdbb := NewCT_CTDescription()
				if _gcbg := d.DecodeElement(_cdbb, &_fga); _gcbg != nil {
					return _gcbg
				}
				_ebce.Desc = append(_ebce.Desc, _cdbb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_ebce.CatLst = NewCT_CTCategories()
				if _ecb := d.DecodeElement(_ebce.CatLst, &_fga); _ecb != nil {
					return _ecb
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_ebce.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _fdf := d.DecodeElement(_ebce.ExtLst, &_fga); _fdf != nil {
					return _fdf
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn\u0020\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006es\u0066\u006f\u0072\u006d\u0048\u0065a\u0064\u0065\u0072 \u0025\u0076", _fga.Name)
				if _aggb := d.Skip(); _aggb != nil {
					return _aggb
				}
			}
		case _d.EndElement:
			break _eaff
		case _d.CharData:
		}
	}
	return nil
}
func (_caeaa ST_ChildOrderType) Validate() error { return _caeaa.ValidateWithPath("") }

type ST_NodeVerticalAlignment byte

func NewCT_CTCategory() *CT_CTCategory { _ecg := &CT_CTCategory{}; return _ecg }
func (_agbegf ST_PtType) String() string {
	switch _agbegf {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u0064\u0065"
	case 2:
		return "\u0061\u0073\u0073\u0074"
	case 3:
		return "\u0064\u006f\u0063"
	case 4:
		return "\u0070\u0072\u0065\u0073"
	case 5:
		return "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case 6:
		return "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return ""
}
func (_bggcg ST_OutputShapeType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_bggcg.String(), start)
}

type AG_ConstraintAttributes struct {
	TypeAttr    ST_ConstraintType
	ForAttr     ST_ConstraintRelationship
	ForNameAttr *string
	PtTypeAttr  ST_ElementType
}
type CT_LayoutVariablePropertySet struct {
	OrgChart      *CT_OrgChart
	ChMax         *CT_ChildMax
	ChPref        *CT_ChildPref
	BulletEnabled *CT_BulletEnabled
	Dir           *CT_Direction
	HierBranch    *CT_HierBranchStyle
	AnimOne       *CT_AnimOne
	AnimLvl       *CT_AnimLvl
	ResizeHandles *CT_ResizeHandles
}

// Validate validates the CT_CxnList and its children
func (_bbde *CT_CxnList) Validate() error {
	return _bbde.ValidateWithPath("\u0043\u0054\u005f\u0043\u0078\u006e\u004c\u0069\u0073\u0074")
}
func (_edeee *ST_PyramidAccentPosition) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_edeee = 0
	case "\u0062\u0065\u0066":
		*_edeee = 1
	case "\u0061\u0066\u0074":
		*_edeee = 2
	}
	return nil
}
func NewCT_ChildMax() *CT_ChildMax                     { _bde := &CT_ChildMax{}; return _bde }
func (_fgbee ST_TextAnchorHorizontal) Validate() error { return _fgbee.ValidateWithPath("") }
func (_cdeb *CT_DiagramDefinition) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cdeb.LayoutNode = NewCT_LayoutNode()
	for _, _cfaf := range start.Attr {
		if _cfaf.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_bdgb, _ceaa := _cfaf.Value, error(nil)
			if _ceaa != nil {
				return _ceaa
			}
			_cdeb.UniqueIdAttr = &_bdgb
			continue
		}
		if _cfaf.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_eebc, _gdec := _cfaf.Value, error(nil)
			if _gdec != nil {
				return _gdec
			}
			_cdeb.MinVerAttr = &_eebc
			continue
		}
		if _cfaf.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_dca, _ebdd := _cfaf.Value, error(nil)
			if _ebdd != nil {
				return _ebdd
			}
			_cdeb.DefStyleAttr = &_dca
			continue
		}
	}
_ecdb:
	for {
		_gcbc, _gdbc := d.Token()
		if _gdbc != nil {
			return _gdbc
		}
		switch _fdeff := _gcbc.(type) {
		case _d.StartElement:
			switch _fdeff.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_eeba := NewCT_Name()
				if _ced := d.DecodeElement(_eeba, &_fdeff); _ced != nil {
					return _ced
				}
				_cdeb.Title = append(_cdeb.Title, _eeba)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_bdac := NewCT_Description()
				if _ddda := d.DecodeElement(_bdac, &_fdeff); _ddda != nil {
					return _ddda
				}
				_cdeb.Desc = append(_cdeb.Desc, _bdac)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_cdeb.CatLst = NewCT_Categories()
				if _gea := d.DecodeElement(_cdeb.CatLst, &_fdeff); _gea != nil {
					return _gea
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0061\u006d\u0070\u0044\u0061\u0074\u0061"}:
				_cdeb.SampData = NewCT_SampleData()
				if _gfa := d.DecodeElement(_cdeb.SampData, &_fdeff); _gfa != nil {
					return _gfa
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073t\u0079\u006c\u0065\u0044\u0061\u0074a"}:
				_cdeb.StyleData = NewCT_SampleData()
				if _cfcd := d.DecodeElement(_cdeb.StyleData, &_fdeff); _cfcd != nil {
					return _cfcd
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063l\u0072\u0044\u0061\u0074\u0061"}:
				_cdeb.ClrData = NewCT_SampleData()
				if _dgag := d.DecodeElement(_cdeb.ClrData, &_fdeff); _dgag != nil {
					return _dgag
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				if _bbfe := d.DecodeElement(_cdeb.LayoutNode, &_fdeff); _bbfe != nil {
					return _bbfe
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_cdeb.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _dgbb := d.DecodeElement(_cdeb.ExtLst, &_fdeff); _dgbb != nil {
					return _dgbb
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070o\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006de\u006et \u006f\u006e\u0020\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0020\u0025\u0076", _fdeff.Name)
				if _gebg := d.Skip(); _gebg != nil {
					return _gebg
				}
			}
		case _d.EndElement:
			break _ecdb
		case _d.CharData:
		}
	}
	return nil
}
func (_aedf *CT_TextProps) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_dgaea:
	for {
		_cecg, _daebd := d.Token()
		if _daebd != nil {
			return _daebd
		}
		switch _bbfa := _cecg.(type) {
		case _d.StartElement:
			switch _bbfa.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0070\u0033\u0064"}, _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0070\u0033\u0064"}:
				_aedf.Sp3d = _ca.NewCT_Shape3D()
				if _gfae := d.DecodeElement(_aedf.Sp3d, &_bbfa); _gfae != nil {
					return _gfae
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0066\u006c\u0061\u0074\u0054\u0078"}, _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0066\u006c\u0061\u0074\u0054\u0078"}:
				_aedf.FlatTx = _ca.NewCT_FlatText()
				if _cbbfb := d.DecodeElement(_aedf.FlatTx, &_bbfa); _cbbfb != nil {
					return _cbbfb
				}
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_T\u0065\u0078t\u0050\u0072\u006f\u0070\u0073\u0020\u0025\u0076", _bbfa.Name)
				if _fede := d.Skip(); _fede != nil {
					return _fede
				}
			}
		case _d.EndElement:
			break _dgaea
		case _d.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the LayoutDefHdrLst and its children, prefixing error messages with path
func (_ccgc *LayoutDefHdrLst) ValidateWithPath(path string) error {
	if _dedae := _ccgc.CT_DiagramDefinitionHeaderLst.ValidateWithPath(path); _dedae != nil {
		return _dedae
	}
	return nil
}
func (_ddeb ST_BendPoint) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ddeed := _d.Attr{}
	_ddeed.Name = name
	switch _ddeb {
	case ST_BendPointUnset:
		_ddeed.Value = ""
	case ST_BendPointBeg:
		_ddeed.Value = "\u0062\u0065\u0067"
	case ST_BendPointDef:
		_ddeed.Value = "\u0064\u0065\u0066"
	case ST_BendPointEnd:
		_ddeed.Value = "\u0065\u006e\u0064"
	}
	return _ddeed, nil
}

const (
	ST_TextAnchorHorizontalUnset ST_TextAnchorHorizontal = 0
	ST_TextAnchorHorizontalNone  ST_TextAnchorHorizontal = 1
	ST_TextAnchorHorizontalCtr   ST_TextAnchorHorizontal = 2
)

type CT_PtList struct{ Pt []*CT_Pt }
type CT_DiagramDefinitionHeaderLst struct{ LayoutDefHdr []*CT_DiagramDefinitionHeader }

func (_dfge *CT_Shape) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _cada := range start.Attr {
		if _cada.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _cada.Name.Local == "\u0062\u006c\u0069\u0070" || _cada.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _cada.Name.Local == "\u0062\u006c\u0069\u0070" {
			_bgaa, _feaae := _cada.Value, error(nil)
			if _feaae != nil {
				return _feaae
			}
			_dfge.BlipAttr = &_bgaa
			continue
		}
		if _cada.Name.Local == "\u0072\u006f\u0074" {
			_dcba, _eeded := _ae.ParseFloat(_cada.Value, 64)
			if _eeded != nil {
				return _eeded
			}
			_dfge.RotAttr = &_dcba
			continue
		}
		if _cada.Name.Local == "\u0074\u0079\u0070\u0065" {
			_afged, _cfcb := ParseUnionST_LayoutShapeType(_cada.Value)
			if _cfcb != nil {
				return _cfcb
			}
			_dfge.TypeAttr = &_afged
			continue
		}
		if _cada.Name.Local == "\u007aO\u0072\u0064\u0065\u0072\u004f\u0066f" {
			_decc, _fbca := _ae.ParseInt(_cada.Value, 10, 32)
			if _fbca != nil {
				return _fbca
			}
			_afdbgb := int32(_decc)
			_dfge.ZOrderOffAttr = &_afdbgb
			continue
		}
		if _cada.Name.Local == "\u0068\u0069\u0064\u0065\u0047\u0065\u006f\u006d" {
			_acgfb, _bafg := _ae.ParseBool(_cada.Value)
			if _bafg != nil {
				return _bafg
			}
			_dfge.HideGeomAttr = &_acgfb
			continue
		}
		if _cada.Name.Local == "\u006ck\u0054\u0078\u0045\u006e\u0074\u0072y" {
			_fcffd, _cdac := _ae.ParseBool(_cada.Value)
			if _cdac != nil {
				return _cdac
			}
			_dfge.LkTxEntryAttr = &_fcffd
			continue
		}
		if _cada.Name.Local == "\u0062l\u0069\u0070\u0050\u0068\u006c\u0064r" {
			_afea, _ddcf := _ae.ParseBool(_cada.Value)
			if _ddcf != nil {
				return _ddcf
			}
			_dfge.BlipPhldrAttr = &_afea
			continue
		}
	}
_fdaa:
	for {
		_geeaa, _affea := d.Token()
		if _affea != nil {
			return _affea
		}
		switch _agad := _geeaa.(type) {
		case _d.StartElement:
			switch _agad.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u0064\u006a\u004c\u0073\u0074"}:
				_dfge.AdjLst = NewCT_AdjLst()
				if _dcbe := d.DecodeElement(_dfge.AdjLst, &_agad); _dcbe != nil {
					return _dcbe
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_dfge.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _cbdd := d.DecodeElement(_dfge.ExtLst, &_agad); _cbdd != nil {
					return _cbdd
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u0020\u0025\u0076", _agad.Name)
				if _bgfg := d.Skip(); _bgfg != nil {
					return _bgfg
				}
			}
		case _d.EndElement:
			break _fdaa
		case _d.CharData:
		}
	}
	return nil
}
func (_ggeb ST_PyramidAccentTextMargin) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_adddd := _d.Attr{}
	_adddd.Name = name
	switch _ggeb {
	case ST_PyramidAccentTextMarginUnset:
		_adddd.Value = ""
	case ST_PyramidAccentTextMarginStep:
		_adddd.Value = "\u0073\u0074\u0065\u0070"
	case ST_PyramidAccentTextMarginStack:
		_adddd.Value = "\u0073\u0074\u0061c\u006b"
	}
	return _adddd, nil
}

type DataModel struct{ CT_DataModel }

func (_beeg *StyleDefHdrLst) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048d\u0072\u004c\u0073\u0074"
	return _beeg.CT_StyleDefinitionHeaderLst.MarshalXML(e, start)
}

type CT_Choose struct {
	NameAttr *string
	If       []*CT_When
	Else     *CT_Otherwise
}

const (
	ST_StartingElementUnset ST_StartingElement = 0
	ST_StartingElementNode  ST_StartingElement = 1
	ST_StartingElementTrans ST_StartingElement = 2
)

func (_dafdg ST_NodeVerticalAlignment) Validate() error { return _dafdg.ValidateWithPath("") }
func (_gfeb ST_ResizeHandlesStr) ValidateWithPath(path string) error {
	switch _gfeb {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gfeb))
	}
	return nil
}
func (_ffcg ST_ContinueDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ffcg.String(), start)
}
func (_ceadbf ST_VariableType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_efefc := _d.Attr{}
	_efefc.Name = name
	switch _ceadbf {
	case ST_VariableTypeUnset:
		_efefc.Value = ""
	case ST_VariableTypeNone:
		_efefc.Value = "\u006e\u006f\u006e\u0065"
	case ST_VariableTypeOrgChart:
		_efefc.Value = "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"
	case ST_VariableTypeChMax:
		_efefc.Value = "\u0063\u0068\u004da\u0078"
	case ST_VariableTypeChPref:
		_efefc.Value = "\u0063\u0068\u0050\u0072\u0065\u0066"
	case ST_VariableTypeBulEnabled:
		_efefc.Value = "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064"
	case ST_VariableTypeDir:
		_efefc.Value = "\u0064\u0069\u0072"
	case ST_VariableTypeHierBranch:
		_efefc.Value = "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"
	case ST_VariableTypeAnimOne:
		_efefc.Value = "\u0061n\u0069\u006d\u004f\u006e\u0065"
	case ST_VariableTypeAnimLvl:
		_efefc.Value = "\u0061n\u0069\u006d\u004c\u0076\u006c"
	case ST_VariableTypeResizeHandles:
		_efefc.Value = "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"
	}
	return _efefc, nil
}
func (_ddgc *CT_Cxn) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _eafb := range start.Attr {
		if _eafb.Name.Local == "\u006do\u0064\u0065\u006c\u0049\u0064" {
			_gfcdf, _gdce := ParseUnionST_ModelId(_eafb.Value)
			if _gdce != nil {
				return _gdce
			}
			_ddgc.ModelIdAttr = _gfcdf
			continue
		}
		if _eafb.Name.Local == "\u0074\u0079\u0070\u0065" {
			_ddgc.TypeAttr.UnmarshalXMLAttr(_eafb)
			continue
		}
		if _eafb.Name.Local == "\u0073\u0072\u0063I\u0064" {
			_ffcc, _bbgd := ParseUnionST_ModelId(_eafb.Value)
			if _bbgd != nil {
				return _bbgd
			}
			_ddgc.SrcIdAttr = _ffcc
			continue
		}
		if _eafb.Name.Local == "\u0064\u0065\u0073\u0074\u0049\u0064" {
			_ggcbg, _ffff := ParseUnionST_ModelId(_eafb.Value)
			if _ffff != nil {
				return _ffff
			}
			_ddgc.DestIdAttr = _ggcbg
			continue
		}
		if _eafb.Name.Local == "\u0073\u0072\u0063\u004f\u0072\u0064" {
			_ffgc, _egef := _ae.ParseUint(_eafb.Value, 10, 32)
			if _egef != nil {
				return _egef
			}
			_ddgc.SrcOrdAttr = uint32(_ffgc)
			continue
		}
		if _eafb.Name.Local == "\u0064e\u0073\u0074\u004f\u0072\u0064" {
			_egdf, _eggg := _ae.ParseUint(_eafb.Value, 10, 32)
			if _eggg != nil {
				return _eggg
			}
			_ddgc.DestOrdAttr = uint32(_egdf)
			continue
		}
		if _eafb.Name.Local == "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073\u0049\u0064" {
			_begd, _fege := ParseUnionST_ModelId(_eafb.Value)
			if _fege != nil {
				return _fege
			}
			_ddgc.ParTransIdAttr = &_begd
			continue
		}
		if _eafb.Name.Local == "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073\u0049\u0064" {
			_ecea, _abcf := ParseUnionST_ModelId(_eafb.Value)
			if _abcf != nil {
				return _abcf
			}
			_ddgc.SibTransIdAttr = &_ecea
			continue
		}
		if _eafb.Name.Local == "\u0070\u0072\u0065\u0073\u0049\u0064" {
			_fdge, _dfde := _eafb.Value, error(nil)
			if _dfde != nil {
				return _dfde
			}
			_ddgc.PresIdAttr = &_fdge
			continue
		}
	}
_gfdebf:
	for {
		_egf, _afa := d.Token()
		if _afa != nil {
			return _afa
		}
		switch _bdc := _egf.(type) {
		case _d.StartElement:
			switch _bdc.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_ddgc.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _gddb := d.DecodeElement(_ddgc.ExtLst, &_bdc); _gddb != nil {
					return _gddb
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u0078\u006e\u0020\u0025\u0076", _bdc.Name)
				if _fffe := d.Skip(); _fffe != nil {
					return _fffe
				}
			}
		case _d.EndElement:
			break _gfdebf
		case _d.CharData:
		}
	}
	return nil
}

const (
	ST_PtTypeUnset    ST_PtType = 0
	ST_PtTypeNode     ST_PtType = 1
	ST_PtTypeAsst     ST_PtType = 2
	ST_PtTypeDoc      ST_PtType = 3
	ST_PtTypePres     ST_PtType = 4
	ST_PtTypeParTrans ST_PtType = 5
	ST_PtTypeSibTrans ST_PtType = 6
)
const (
	ST_FallbackDimensionUnset ST_FallbackDimension = 0
	ST_FallbackDimension1D    ST_FallbackDimension = 1
	ST_FallbackDimension2D    ST_FallbackDimension = 2
)

func (_eacaf ST_SecondaryLinearDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_eacaf.String(), start)
}
func (_bege ST_RotationPath) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_bege.String(), start)
}
func NewCT_SampleData() *CT_SampleData { _gcacd := &CT_SampleData{}; return _gcacd }
func (_cgdgb ST_TextBlockDirection) String() string {
	switch _cgdgb {
	case 0:
		return ""
	case 1:
		return "\u0068\u006f\u0072\u007a"
	case 2:
		return "\u0076\u0065\u0072\u0074"
	}
	return ""
}

// ValidateWithPath validates the CT_BulletEnabled and its children, prefixing error messages with path
func (_ccf *CT_BulletEnabled) ValidateWithPath(path string) error { return nil }
func (_fcda *CT_StyleDefinition) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gdde := range start.Attr {
		if _gdde.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_cefc, _gdfeg := _gdde.Value, error(nil)
			if _gdfeg != nil {
				return _gdfeg
			}
			_fcda.UniqueIdAttr = &_cefc
			continue
		}
		if _gdde.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_fgef, _gggg := _gdde.Value, error(nil)
			if _gggg != nil {
				return _gggg
			}
			_fcda.MinVerAttr = &_fgef
			continue
		}
	}
_cagab:
	for {
		_fbbg, _befa := d.Token()
		if _befa != nil {
			return _befa
		}
		switch _egba := _fbbg.(type) {
		case _d.StartElement:
			switch _egba.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_cdcb := NewCT_SDName()
				if _gecb := d.DecodeElement(_cdcb, &_egba); _gecb != nil {
					return _gecb
				}
				_fcda.Title = append(_fcda.Title, _cdcb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_bfbab := NewCT_SDDescription()
				if _fgdbg := d.DecodeElement(_bfbab, &_egba); _fgdbg != nil {
					return _fgdbg
				}
				_fcda.Desc = append(_fcda.Desc, _bfbab)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_fcda.CatLst = NewCT_SDCategories()
				if _gdbcc := d.DecodeElement(_fcda.CatLst, &_egba); _gdbcc != nil {
					return _gdbcc
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}:
				_fcda.Scene3d = _ca.NewCT_Scene3D()
				if _abba := d.DecodeElement(_fcda.Scene3d, &_egba); _abba != nil {
					return _abba
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_dedf := NewCT_StyleLabel()
				if _cddd := d.DecodeElement(_dedf, &_egba); _cddd != nil {
					return _cddd
				}
				_fcda.StyleLbl = append(_fcda.StyleLbl, _dedf)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_fcda.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _gdbee := d.DecodeElement(_fcda.ExtLst, &_egba); _gdbee != nil {
					return _gdbee
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065l\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u0044e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0020\u0025\u0076", _egba.Name)
				if _gfbf := d.Skip(); _gfbf != nil {
					return _gfbf
				}
			}
		case _d.EndElement:
			break _cagab
		case _d.CharData:
		}
	}
	return nil
}
func (_bcde *CT_Constraints) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _bcde.Constr != nil {
		_bcdd := _d.StartElement{Name: _d.Name{Local: "\u0063\u006f\u006e\u0073\u0074\u0072"}}
		for _, _cdbbb := range _bcde.Constr {
			e.EncodeElement(_cdbbb, _bcdd)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_feea *ST_TextBlockDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bcgdb, _bfdde := d.Token()
	if _bfdde != nil {
		return _bfdde
	}
	if _cbecc, _ccdac := _bcgdb.(_d.EndElement); _ccdac && _cbecc.Name == start.Name {
		*_feea = 1
		return nil
	}
	if _abge, _fbefd := _bcgdb.(_d.CharData); !_fbefd {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bcgdb)
	} else {
		switch string(_abge) {
		case "":
			*_feea = 0
		case "\u0068\u006f\u0072\u007a":
			*_feea = 1
		case "\u0076\u0065\u0072\u0074":
			*_feea = 2
		}
	}
	_bcgdb, _bfdde = d.Token()
	if _bfdde != nil {
		return _bfdde
	}
	if _gffdb, _bcag := _bcgdb.(_d.EndElement); _bcag && _gffdb.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bcgdb)
}

// Validate validates the CT_Algorithm and its children
func (_fagg *CT_Algorithm) Validate() error {
	return _fagg.ValidateWithPath("\u0043\u0054\u005fA\u006c\u0067\u006f\u0072\u0069\u0074\u0068\u006d")
}
func (_geaff ST_NodeHorizontalAlignment) String() string {
	switch _geaff {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0063\u0074\u0072"
	case 3:
		return "\u0072"
	}
	return ""
}
func (_caffe ST_ConstraintType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_edec := _d.Attr{}
	_edec.Name = name
	switch _caffe {
	case ST_ConstraintTypeUnset:
		_edec.Value = ""
	case ST_ConstraintTypeNone:
		_edec.Value = "\u006e\u006f\u006e\u0065"
	case ST_ConstraintTypeAlignOff:
		_edec.Value = "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066"
	case ST_ConstraintTypeBegMarg:
		_edec.Value = "\u0062e\u0067\u004d\u0061\u0072\u0067"
	case ST_ConstraintTypeBendDist:
		_edec.Value = "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074"
	case ST_ConstraintTypeBegPad:
		_edec.Value = "\u0062\u0065\u0067\u0050\u0061\u0064"
	case ST_ConstraintTypeB:
		_edec.Value = "\u0062"
	case ST_ConstraintTypeBMarg:
		_edec.Value = "\u0062\u004d\u0061r\u0067"
	case ST_ConstraintTypeBOff:
		_edec.Value = "\u0062\u004f\u0066\u0066"
	case ST_ConstraintTypeCtrX:
		_edec.Value = "\u0063\u0074\u0072\u0058"
	case ST_ConstraintTypeCtrXOff:
		_edec.Value = "\u0063t\u0072\u0058\u004f\u0066\u0066"
	case ST_ConstraintTypeCtrY:
		_edec.Value = "\u0063\u0074\u0072\u0059"
	case ST_ConstraintTypeCtrYOff:
		_edec.Value = "\u0063t\u0072\u0059\u004f\u0066\u0066"
	case ST_ConstraintTypeConnDist:
		_edec.Value = "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074"
	case ST_ConstraintTypeDiam:
		_edec.Value = "\u0064\u0069\u0061\u006d"
	case ST_ConstraintTypeEndMarg:
		_edec.Value = "\u0065n\u0064\u004d\u0061\u0072\u0067"
	case ST_ConstraintTypeEndPad:
		_edec.Value = "\u0065\u006e\u0064\u0050\u0061\u0064"
	case ST_ConstraintTypeH:
		_edec.Value = "\u0068"
	case ST_ConstraintTypeHArH:
		_edec.Value = "\u0068\u0041\u0072\u0048"
	case ST_ConstraintTypeHOff:
		_edec.Value = "\u0068\u004f\u0066\u0066"
	case ST_ConstraintTypeL:
		_edec.Value = "\u006c"
	case ST_ConstraintTypeLMarg:
		_edec.Value = "\u006c\u004d\u0061r\u0067"
	case ST_ConstraintTypeLOff:
		_edec.Value = "\u006c\u004f\u0066\u0066"
	case ST_ConstraintTypeR:
		_edec.Value = "\u0072"
	case ST_ConstraintTypeRMarg:
		_edec.Value = "\u0072\u004d\u0061r\u0067"
	case ST_ConstraintTypeROff:
		_edec.Value = "\u0072\u004f\u0066\u0066"
	case ST_ConstraintTypePrimFontSz:
		_edec.Value = "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a"
	case ST_ConstraintTypePyraAcctRatio:
		_edec.Value = "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f"
	case ST_ConstraintTypeSecFontSz:
		_edec.Value = "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z"
	case ST_ConstraintTypeSibSp:
		_edec.Value = "\u0073\u0069\u0062S\u0070"
	case ST_ConstraintTypeSecSibSp:
		_edec.Value = "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070"
	case ST_ConstraintTypeSp:
		_edec.Value = "\u0073\u0070"
	case ST_ConstraintTypeStemThick:
		_edec.Value = "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k"
	case ST_ConstraintTypeT:
		_edec.Value = "\u0074"
	case ST_ConstraintTypeTMarg:
		_edec.Value = "\u0074\u004d\u0061r\u0067"
	case ST_ConstraintTypeTOff:
		_edec.Value = "\u0074\u004f\u0066\u0066"
	case ST_ConstraintTypeUserA:
		_edec.Value = "\u0075\u0073\u0065r\u0041"
	case ST_ConstraintTypeUserB:
		_edec.Value = "\u0075\u0073\u0065r\u0042"
	case ST_ConstraintTypeUserC:
		_edec.Value = "\u0075\u0073\u0065r\u0043"
	case ST_ConstraintTypeUserD:
		_edec.Value = "\u0075\u0073\u0065r\u0044"
	case ST_ConstraintTypeUserE:
		_edec.Value = "\u0075\u0073\u0065r\u0045"
	case ST_ConstraintTypeUserF:
		_edec.Value = "\u0075\u0073\u0065r\u0046"
	case ST_ConstraintTypeUserG:
		_edec.Value = "\u0075\u0073\u0065r\u0047"
	case ST_ConstraintTypeUserH:
		_edec.Value = "\u0075\u0073\u0065r\u0048"
	case ST_ConstraintTypeUserI:
		_edec.Value = "\u0075\u0073\u0065r\u0049"
	case ST_ConstraintTypeUserJ:
		_edec.Value = "\u0075\u0073\u0065r\u004a"
	case ST_ConstraintTypeUserK:
		_edec.Value = "\u0075\u0073\u0065r\u004b"
	case ST_ConstraintTypeUserL:
		_edec.Value = "\u0075\u0073\u0065r\u004c"
	case ST_ConstraintTypeUserM:
		_edec.Value = "\u0075\u0073\u0065r\u004d"
	case ST_ConstraintTypeUserN:
		_edec.Value = "\u0075\u0073\u0065r\u004e"
	case ST_ConstraintTypeUserO:
		_edec.Value = "\u0075\u0073\u0065r\u004f"
	case ST_ConstraintTypeUserP:
		_edec.Value = "\u0075\u0073\u0065r\u0050"
	case ST_ConstraintTypeUserQ:
		_edec.Value = "\u0075\u0073\u0065r\u0051"
	case ST_ConstraintTypeUserR:
		_edec.Value = "\u0075\u0073\u0065r\u0052"
	case ST_ConstraintTypeUserS:
		_edec.Value = "\u0075\u0073\u0065r\u0053"
	case ST_ConstraintTypeUserT:
		_edec.Value = "\u0075\u0073\u0065r\u0054"
	case ST_ConstraintTypeUserU:
		_edec.Value = "\u0075\u0073\u0065r\u0055"
	case ST_ConstraintTypeUserV:
		_edec.Value = "\u0075\u0073\u0065r\u0056"
	case ST_ConstraintTypeUserW:
		_edec.Value = "\u0075\u0073\u0065r\u0057"
	case ST_ConstraintTypeUserX:
		_edec.Value = "\u0075\u0073\u0065r\u0058"
	case ST_ConstraintTypeUserY:
		_edec.Value = "\u0075\u0073\u0065r\u0059"
	case ST_ConstraintTypeUserZ:
		_edec.Value = "\u0075\u0073\u0065r\u005a"
	case ST_ConstraintTypeW:
		_edec.Value = "\u0077"
	case ST_ConstraintTypeWArH:
		_edec.Value = "\u0077\u0041\u0072\u0048"
	case ST_ConstraintTypeWOff:
		_edec.Value = "\u0077\u004f\u0066\u0066"
	}
	return _edec, nil
}

// Validate validates the CT_CTName and its children
func (_fdef *CT_CTName) Validate() error {
	return _fdef.ValidateWithPath("\u0043T\u005f\u0043\u0054\u004e\u0061\u006de")
}

type CT_ColorTransformHeader struct {
	UniqueIdAttr string
	MinVerAttr   *string
	ResIdAttr    *int32
	Title        []*CT_CTName
	Desc         []*CT_CTDescription
	CatLst       *CT_CTCategories
	ExtLst       *_ca.CT_OfficeArtExtensionList
}

func (_ecegf ST_BendPoint) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ecegf.String(), start)
}
func (_eacaa *ST_HueDir) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ecc, _fdcdd := d.Token()
	if _fdcdd != nil {
		return _fdcdd
	}
	if _abdfd, _fggca := _ecc.(_d.EndElement); _fggca && _abdfd.Name == start.Name {
		*_eacaa = 1
		return nil
	}
	if _aggce, _bbbd := _ecc.(_d.CharData); !_bbbd {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ecc)
	} else {
		switch string(_aggce) {
		case "":
			*_eacaa = 0
		case "\u0063\u0077":
			*_eacaa = 1
		case "\u0063\u0063\u0077":
			*_eacaa = 2
		}
	}
	_ecc, _fdcdd = d.Token()
	if _fdcdd != nil {
		return _fdcdd
	}
	if _bdbeg, _abab := _ecc.(_d.EndElement); _abab && _bdbeg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ecc)
}
func NewStyleDef() *StyleDef {
	_bdag := &StyleDef{}
	_bdag.CT_StyleDefinition = *NewCT_StyleDefinition()
	return _bdag
}

type CT_HierBranchStyle struct{ ValAttr ST_HierBranchStyle }

func (_ebb *CT_Adj) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0069\u0064\u0078"}, Value: _c.Sprintf("\u0025\u0076", _ebb.IdxAttr)})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _ebb.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewAG_ConstraintAttributes() *AG_ConstraintAttributes {
	_cf := &AG_ConstraintAttributes{}
	return _cf
}
func (_gg *CT_CTCategory) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _egc := range start.Attr {
		if _egc.Name.Local == "\u0074\u0079\u0070\u0065" {
			_bfb, _cabd := _egc.Value, error(nil)
			if _cabd != nil {
				return _cabd
			}
			_gg.TypeAttr = _bfb
			continue
		}
		if _egc.Name.Local == "\u0070\u0072\u0069" {
			_ggc, _gce := _ae.ParseUint(_egc.Value, 10, 32)
			if _gce != nil {
				return _gce
			}
			_gg.PriAttr = uint32(_ggc)
			continue
		}
	}
	for {
		_ffd, _dde := d.Token()
		if _dde != nil {
			return _c.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0043T\u0043a\u0074e\u0067\u006f\u0072\u0079\u003a\u0020\u0025s", _dde)
		}
		if _cbga, _daef := _ffd.(_d.EndElement); _daef && _cbga.Name == start.Name {
			break
		}
	}
	return nil
}
func (_eagb *StyleDef) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_eagb.CT_StyleDefinition = *NewCT_StyleDefinition()
	for _, _afggb := range start.Attr {
		if _afggb.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_edea, _cdfa := _afggb.Value, error(nil)
			if _cdfa != nil {
				return _cdfa
			}
			_eagb.UniqueIdAttr = &_edea
			continue
		}
		if _afggb.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_affeg, _bcfd := _afggb.Value, error(nil)
			if _bcfd != nil {
				return _bcfd
			}
			_eagb.MinVerAttr = &_affeg
			continue
		}
	}
_effbc:
	for {
		_adgdf, _baeb := d.Token()
		if _baeb != nil {
			return _baeb
		}
		switch _gebe := _adgdf.(type) {
		case _d.StartElement:
			switch _gebe.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_fagcd := NewCT_SDName()
				if _aadda := d.DecodeElement(_fagcd, &_gebe); _aadda != nil {
					return _aadda
				}
				_eagb.Title = append(_eagb.Title, _fagcd)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_ffbg := NewCT_SDDescription()
				if _dfcgd := d.DecodeElement(_ffbg, &_gebe); _dfcgd != nil {
					return _dfcgd
				}
				_eagb.Desc = append(_eagb.Desc, _ffbg)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_eagb.CatLst = NewCT_SDCategories()
				if _ecad := d.DecodeElement(_eagb.CatLst, &_gebe); _ecad != nil {
					return _ecad
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}:
				_eagb.Scene3d = _ca.NewCT_Scene3D()
				if _gcce := d.DecodeElement(_eagb.Scene3d, &_gebe); _gcce != nil {
					return _gcce
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_gecfc := NewCT_StyleLabel()
				if _baab := d.DecodeElement(_gecfc, &_gebe); _baab != nil {
					return _baab
				}
				_eagb.StyleLbl = append(_eagb.StyleLbl, _gecfc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_eagb.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _ebbcc := d.DecodeElement(_eagb.ExtLst, &_gebe); _ebbcc != nil {
					return _ebbcc
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0020\u0025\u0076", _gebe.Name)
				if _ffgdb := d.Skip(); _ffgdb != nil {
					return _ffgdb
				}
			}
		case _d.EndElement:
			break _effbc
		case _d.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_NumericRule and its children, prefixing error messages with path
func (_bggd *CT_NumericRule) ValidateWithPath(path string) error {
	if _bggd.ExtLst != nil {
		if _febbg := _bggd.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _febbg != nil {
			return _febbg
		}
	}
	if _eebe := _bggd.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _eebe != nil {
		return _eebe
	}
	if _dgee := _bggd.ForAttr.ValidateWithPath(path + "\u002f\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _dgee != nil {
		return _dgee
	}
	if _bagd := _bggd.PtTypeAttr.ValidateWithPath(path + "/\u0050\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _bagd != nil {
		return _bagd
	}
	return nil
}
func NewCT_Constraints() *CT_Constraints { _bae := &CT_Constraints{}; return _bae }
func (_ebfaf *ST_AlgorithmType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ebfaf = 0
	case "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e":
		*_ebfaf = 1
	case "\u0063\u006f\u006e\u006e":
		*_ebfaf = 2
	case "\u0063\u0079\u0063l\u0065":
		*_ebfaf = 3
	case "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd":
		*_ebfaf = 4
	case "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074":
		*_ebfaf = 5
	case "\u0070\u0079\u0072\u0061":
		*_ebfaf = 6
	case "\u006c\u0069\u006e":
		*_ebfaf = 7
	case "\u0073\u0070":
		*_ebfaf = 8
	case "\u0074\u0078":
		*_ebfaf = 9
	case "\u0073\u006e\u0061k\u0065":
		*_ebfaf = 10
	}
	return nil
}

const (
	ST_HierarchyAlignmentUnset   ST_HierarchyAlignment = 0
	ST_HierarchyAlignmentTL      ST_HierarchyAlignment = 1
	ST_HierarchyAlignmentTR      ST_HierarchyAlignment = 2
	ST_HierarchyAlignmentTCtrCh  ST_HierarchyAlignment = 3
	ST_HierarchyAlignmentTCtrDes ST_HierarchyAlignment = 4
	ST_HierarchyAlignmentBL      ST_HierarchyAlignment = 5
	ST_HierarchyAlignmentBR      ST_HierarchyAlignment = 6
	ST_HierarchyAlignmentBCtrCh  ST_HierarchyAlignment = 7
	ST_HierarchyAlignmentBCtrDes ST_HierarchyAlignment = 8
	ST_HierarchyAlignmentLT      ST_HierarchyAlignment = 9
	ST_HierarchyAlignmentLB      ST_HierarchyAlignment = 10
	ST_HierarchyAlignmentLCtrCh  ST_HierarchyAlignment = 11
	ST_HierarchyAlignmentLCtrDes ST_HierarchyAlignment = 12
	ST_HierarchyAlignmentRT      ST_HierarchyAlignment = 13
	ST_HierarchyAlignmentRB      ST_HierarchyAlignment = 14
	ST_HierarchyAlignmentRCtrCh  ST_HierarchyAlignment = 15
	ST_HierarchyAlignmentRCtrDes ST_HierarchyAlignment = 16
)

type CT_RelIds struct {
	DmAttr string
	LoAttr string
	QsAttr string
	CsAttr string
}

func (_gafa *ColorsDef) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gafa.CT_ColorTransform = *NewCT_ColorTransform()
	for _, _fggbbe := range start.Attr {
		if _fggbbe.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_dfgfe, _dcggd := _fggbbe.Value, error(nil)
			if _dcggd != nil {
				return _dcggd
			}
			_gafa.UniqueIdAttr = &_dfgfe
			continue
		}
		if _fggbbe.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_aafc, _ecdbd := _fggbbe.Value, error(nil)
			if _ecdbd != nil {
				return _ecdbd
			}
			_gafa.MinVerAttr = &_aafc
			continue
		}
	}
_gfafe:
	for {
		_eeca, _adbe := d.Token()
		if _adbe != nil {
			return _adbe
		}
		switch _adcg := _eeca.(type) {
		case _d.StartElement:
			switch _adcg.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_fccea := NewCT_CTName()
				if _dccc := d.DecodeElement(_fccea, &_adcg); _dccc != nil {
					return _dccc
				}
				_gafa.Title = append(_gafa.Title, _fccea)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_gaaa := NewCT_CTDescription()
				if _bagbe := d.DecodeElement(_gaaa, &_adcg); _bagbe != nil {
					return _bagbe
				}
				_gafa.Desc = append(_gafa.Desc, _gaaa)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_gafa.CatLst = NewCT_CTCategories()
				if _cgfec := d.DecodeElement(_gafa.CatLst, &_adcg); _cgfec != nil {
					return _cgfec
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_aegeg := NewCT_CTStyleLabel()
				if _fdebe := d.DecodeElement(_aegeg, &_adcg); _fdebe != nil {
					return _fdebe
				}
				_gafa.StyleLbl = append(_gafa.StyleLbl, _aegeg)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_gafa.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _bded := d.DecodeElement(_gafa.ExtLst, &_adcg); _bded != nil {
					return _bded
				}
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u006f\u006c\u006f\u0072\u0073D\u0065\u0066 \u0025\u0076", _adcg.Name)
				if _cdccb := d.Skip(); _cdccb != nil {
					return _cdccb
				}
			}
		case _d.EndElement:
			break _gfafe
		case _d.CharData:
		}
	}
	return nil
}
func (_effda *CT_Parameter) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_effda.TypeAttr = ST_ParameterId(1)
	for _, _eedee := range start.Attr {
		if _eedee.Name.Local == "\u0074\u0079\u0070\u0065" {
			_effda.TypeAttr.UnmarshalXMLAttr(_eedee)
			continue
		}
		if _eedee.Name.Local == "\u0076\u0061\u006c" {
			_fead, _bddd := ParseUnionST_ParameterVal(_eedee.Value)
			if _bddd != nil {
				return _bddd
			}
			_effda.ValAttr = _fead
			continue
		}
	}
	for {
		_ggcad, _fefb := d.Token()
		if _fefb != nil {
			return _c.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005f\u0050\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u003a\u0020\u0025\u0073", _fefb)
		}
		if _bdcg, _caff := _ggcad.(_d.EndElement); _caff && _bdcg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ddfge *ST_PyramidAccentTextMargin) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cgcb, _feefg := d.Token()
	if _feefg != nil {
		return _feefg
	}
	if _ffffe, _bfcc := _cgcb.(_d.EndElement); _bfcc && _ffffe.Name == start.Name {
		*_ddfge = 1
		return nil
	}
	if _cgebb, _edca := _cgcb.(_d.CharData); !_edca {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cgcb)
	} else {
		switch string(_cgebb) {
		case "":
			*_ddfge = 0
		case "\u0073\u0074\u0065\u0070":
			*_ddfge = 1
		case "\u0073\u0074\u0061c\u006b":
			*_ddfge = 2
		}
	}
	_cgcb, _feefg = d.Token()
	if _feefg != nil {
		return _feefg
	}
	if _adeba, _ddbcc := _cgcb.(_d.EndElement); _ddbcc && _adeba.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cgcb)
}
func (_acbcf ST_StartingElement) String() string {
	switch _acbcf {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u0064\u0065"
	case 2:
		return "\u0074\u0072\u0061n\u0073"
	}
	return ""
}
func NewCT_Choose() *CT_Choose           { _bcad := &CT_Choose{}; return _bcad }
func (_gaedg ST_Offset) Validate() error { return _gaedg.ValidateWithPath("") }
func (_cabg *CT_PresentationOf) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _ecbd := range start.Attr {
		if _ecbd.Name.Local == "\u0061\u0078\u0069\u0073" {
			_edgfg, _aebe := ParseSliceST_AxisTypes(_ecbd.Value)
			if _aebe != nil {
				return _aebe
			}
			_cabg.AxisAttr = &_edgfg
			continue
		}
		if _ecbd.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_dbfab, _acaae := ParseSliceST_ElementTypes(_ecbd.Value)
			if _acaae != nil {
				return _acaae
			}
			_cabg.PtTypeAttr = &_dbfab
			continue
		}
		if _ecbd.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_gedd, _dfca := ParseSliceST_Booleans(_ecbd.Value)
			if _dfca != nil {
				return _dfca
			}
			_cabg.HideLastTransAttr = &_gedd
			continue
		}
		if _ecbd.Name.Local == "\u0073\u0074" {
			_cdd, _cabgg := ParseSliceST_Ints(_ecbd.Value)
			if _cabgg != nil {
				return _cabgg
			}
			_cabg.StAttr = &_cdd
			continue
		}
		if _ecbd.Name.Local == "\u0063\u006e\u0074" {
			_bgcd, _cgcc := ParseSliceST_UnsignedInts(_ecbd.Value)
			if _cgcc != nil {
				return _cgcc
			}
			_cabg.CntAttr = &_bgcd
			continue
		}
		if _ecbd.Name.Local == "\u0073\u0074\u0065\u0070" {
			_ceccf, _fabf := ParseSliceST_Ints(_ecbd.Value)
			if _fabf != nil {
				return _fabf
			}
			_cabg.StepAttr = &_ceccf
			continue
		}
	}
_fcga:
	for {
		_ceede, _abfc := d.Token()
		if _abfc != nil {
			return _abfc
		}
		switch _cffa := _ceede.(type) {
		case _d.StartElement:
			switch _cffa.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_cabg.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _fcdd := d.DecodeElement(_cabg.ExtLst, &_cffa); _fcdd != nil {
					return _fcdd
				}
			default:
				_af.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u004f\u0066\u0020\u0025\u0076", _cffa.Name)
				if _geaa := d.Skip(); _geaa != nil {
					return _geaa
				}
			}
		case _d.EndElement:
			break _fcga
		case _d.CharData:
		}
	}
	return nil
}
func (_gddg ST_ChildOrderType) String() string {
	switch _gddg {
	case 0:
		return ""
	case 1:
		return "\u0062"
	case 2:
		return "\u0074"
	}
	return ""
}
func (_dgbeg ST_SecondaryChildAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_gfce := _d.Attr{}
	_gfce.Name = name
	switch _dgbeg {
	case ST_SecondaryChildAlignmentUnset:
		_gfce.Value = ""
	case ST_SecondaryChildAlignmentNone:
		_gfce.Value = "\u006e\u006f\u006e\u0065"
	case ST_SecondaryChildAlignmentT:
		_gfce.Value = "\u0074"
	case ST_SecondaryChildAlignmentB:
		_gfce.Value = "\u0062"
	case ST_SecondaryChildAlignmentL:
		_gfce.Value = "\u006c"
	case ST_SecondaryChildAlignmentR:
		_gfce.Value = "\u0072"
	}
	return _gfce, nil
}

type ST_BendPoint byte

func NewCT_Rules() *CT_Rules { _cegf := &CT_Rules{}; return _cegf }
func (_dggfb *ST_BendPoint) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gecgb, _bdca := d.Token()
	if _bdca != nil {
		return _bdca
	}
	if _eaeag, _bfeg := _gecgb.(_d.EndElement); _bfeg && _eaeag.Name == start.Name {
		*_dggfb = 1
		return nil
	}
	if _daag, _caeaab := _gecgb.(_d.CharData); !_caeaab {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gecgb)
	} else {
		switch string(_daag) {
		case "":
			*_dggfb = 0
		case "\u0062\u0065\u0067":
			*_dggfb = 1
		case "\u0064\u0065\u0066":
			*_dggfb = 2
		case "\u0065\u006e\u0064":
			*_dggfb = 3
		}
	}
	_gecgb, _bdca = d.Token()
	if _bdca != nil {
		return _bdca
	}
	if _ggfe, _acdac := _gecgb.(_d.EndElement); _acdac && _ggfe.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gecgb)
}

// ValidateWithPath validates the CT_StyleLabel and its children, prefixing error messages with path
func (_deaa *CT_StyleLabel) ValidateWithPath(path string) error {
	if _deaa.Scene3d != nil {
		if _cead := _deaa.Scene3d.ValidateWithPath(path + "\u002f\u0053\u0063\u0065\u006e\u0065\u0033\u0064"); _cead != nil {
			return _cead
		}
	}
	if _deaa.Sp3d != nil {
		if _fbbab := _deaa.Sp3d.ValidateWithPath(path + "\u002f\u0053\u00703\u0064"); _fbbab != nil {
			return _fbbab
		}
	}
	if _deaa.TxPr != nil {
		if _fcgc := _deaa.TxPr.ValidateWithPath(path + "\u002f\u0054\u0078P\u0072"); _fcgc != nil {
			return _fcgc
		}
	}
	if _deaa.Style != nil {
		if _cadea := _deaa.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _cadea != nil {
			return _cadea
		}
	}
	if _deaa.ExtLst != nil {
		if _aaba := _deaa.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _aaba != nil {
			return _aaba
		}
	}
	return nil
}

type ST_ConnectorRouting byte

func (_dfb *CT_CTName) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _dfb.LangAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_dfb.LangAttr)})
	}
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _dfb.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_ffcdf ST_AutoTextRotation) ValidateWithPath(path string) error {
	switch _ffcdf {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ffcdf))
	}
	return nil
}
func (_bbgge ST_SecondaryChildAlignment) ValidateWithPath(path string) error {
	switch _bbgge {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bbgge))
	}
	return nil
}
func (_gbcec *ST_ChildOrderType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dffdf, _fgfdd := d.Token()
	if _fgfdd != nil {
		return _fgfdd
	}
	if _bcfc, _adddf := _dffdf.(_d.EndElement); _adddf && _bcfc.Name == start.Name {
		*_gbcec = 1
		return nil
	}
	if _gebgd, _gcaf := _dffdf.(_d.CharData); !_gcaf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dffdf)
	} else {
		switch string(_gebgd) {
		case "":
			*_gbcec = 0
		case "\u0062":
			*_gbcec = 1
		case "\u0074":
			*_gbcec = 2
		}
	}
	_dffdf, _fgfdd = d.Token()
	if _fgfdd != nil {
		return _fgfdd
	}
	if _bbfac, _ffef := _dffdf.(_d.EndElement); _ffef && _bbfac.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dffdf)
}

// Validate validates the CT_Pt and its children
func (_afe *CT_Pt) Validate() error { return _afe.ValidateWithPath("\u0043\u0054\u005fP\u0074") }

type ST_AlgorithmType byte

func (_cacd ST_DiagramHorizontalAlignment) ValidateWithPath(path string) error {
	switch _cacd {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cacd))
	}
	return nil
}
func (_adc *CT_AdjLst) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _adc.Adj != nil {
		_geb := _d.StartElement{Name: _d.Name{Local: "\u0061\u0064\u006a"}}
		for _, _cacf := range _adc.Adj {
			e.EncodeElement(_cacf, _geb)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_fbbgc ST_Direction) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fbbgc.String(), start)
}
func NewCT_SDCategory() *CT_SDCategory { _caffb := &CT_SDCategory{}; return _caffb }

// ValidateWithPath validates the CT_CxnList and its children, prefixing error messages with path
func (_cbbc *CT_CxnList) ValidateWithPath(path string) error {
	for _ddag, _bdg := range _cbbc.Cxn {
		if _agfd := _bdg.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0078\u006e\u005b\u0025\u0064\u005d", path, _ddag)); _agfd != nil {
			return _agfd
		}
	}
	return nil
}
func (_cbbdc ST_AlgorithmType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_cbbdc.String(), start)
}
func (_afcce *CT_SDName) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gadbc := range start.Attr {
		if _gadbc.Name.Local == "\u006c\u0061\u006e\u0067" {
			_fcgfc, _cbca := _gadbc.Value, error(nil)
			if _cbca != nil {
				return _cbca
			}
			_afcce.LangAttr = &_fcgfc
			continue
		}
		if _gadbc.Name.Local == "\u0076\u0061\u006c" {
			_fccge, _cccg := _gadbc.Value, error(nil)
			if _cccg != nil {
				return _cccg
			}
			_afcce.ValAttr = _fccge
			continue
		}
	}
	for {
		_bebg, _ggddc := d.Token()
		if _ggddc != nil {
			return _c.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0044\u004ea\u006d\u0065\u003a\u0020\u0025\u0073", _ggddc)
		}
		if _agbgf, _ecag := _bebg.(_d.EndElement); _ecag && _agbgf.Name == start.Name {
			break
		}
	}
	return nil
}
func (_cfba ST_SecondaryLinearDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ffed := _d.Attr{}
	_ffed.Name = name
	switch _cfba {
	case ST_SecondaryLinearDirectionUnset:
		_ffed.Value = ""
	case ST_SecondaryLinearDirectionNone:
		_ffed.Value = "\u006e\u006f\u006e\u0065"
	case ST_SecondaryLinearDirectionFromL:
		_ffed.Value = "\u0066\u0072\u006fm\u004c"
	case ST_SecondaryLinearDirectionFromR:
		_ffed.Value = "\u0066\u0072\u006fm\u0052"
	case ST_SecondaryLinearDirectionFromT:
		_ffed.Value = "\u0066\u0072\u006fm\u0054"
	case ST_SecondaryLinearDirectionFromB:
		_ffed.Value = "\u0066\u0072\u006fm\u0042"
	}
	return _ffed, nil
}

type ST_VariableType byte

func (_deccc *ST_FunctionValue) ValidateWithPath(path string) error {
	_fbddf := []string{}
	if _deccc.Int32 != nil {
		_fbddf = append(_fbddf, "\u0049\u006e\u00743\u0032")
	}
	if _deccc.Bool != nil {
		_fbddf = append(_fbddf, "\u0042\u006f\u006f\u006c")
	}
	if _deccc.ST_Direction != ST_DirectionUnset {
		_fbddf = append(_fbddf, "\u0053\u0054\u005fD\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _deccc.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		_fbddf = append(_fbddf, "\u0053T\u005fH\u0069\u0065\u0072\u0042\u0072a\u006e\u0063h\u0053\u0074\u0079\u006c\u0065")
	}
	if _deccc.ST_AnimOneStr != ST_AnimOneStrUnset {
		_fbddf = append(_fbddf, "\u0053\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065\u0053\u0074\u0072")
	}
	if _deccc.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		_fbddf = append(_fbddf, "\u0053\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c\u0053\u0074\u0072")
	}
	if _deccc.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		_fbddf = append(_fbddf, "\u0053\u0054\u005f\u0052es\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073\u0053\u0074\u0072")
	}
	if len(_fbddf) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _fbddf)
	}
	return nil
}
func (_ff *CT_AdjLst) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_cfb:
	for {
		_ecd, _gfb := d.Token()
		if _gfb != nil {
			return _gfb
		}
		switch _cace := _ecd.(type) {
		case _d.StartElement:
			switch _cace.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u0064\u006a"}:
				_eff := NewCT_Adj()
				if _fbb := d.DecodeElement(_eff, &_cace); _fbb != nil {
					return _fbb
				}
				_ff.Adj = append(_ff.Adj, _eff)
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0041\u0064\u006aL\u0073\u0074 \u0025\u0076", _cace.Name)
				if _cgg := d.Skip(); _cgg != nil {
					return _cgg
				}
			}
		case _d.EndElement:
			break _cfb
		case _d.CharData:
		}
	}
	return nil
}
func (_effdg *CT_SDDescription) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _effdg.LangAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_effdg.LangAttr)})
	}
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _effdg.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_ebcdb ST_HierarchyAlignment) String() string {
	switch _ebcdb {
	case 0:
		return ""
	case 1:
		return "\u0074\u004c"
	case 2:
		return "\u0074\u0052"
	case 3:
		return "\u0074\u0043\u0074\u0072\u0043\u0068"
	case 4:
		return "\u0074C\u0074\u0072\u0044\u0065\u0073"
	case 5:
		return "\u0062\u004c"
	case 6:
		return "\u0062\u0052"
	case 7:
		return "\u0062\u0043\u0074\u0072\u0043\u0068"
	case 8:
		return "\u0062C\u0074\u0072\u0044\u0065\u0073"
	case 9:
		return "\u006c\u0054"
	case 10:
		return "\u006c\u0042"
	case 11:
		return "\u006c\u0043\u0074\u0072\u0043\u0068"
	case 12:
		return "\u006cC\u0074\u0072\u0044\u0065\u0073"
	case 13:
		return "\u0072\u0054"
	case 14:
		return "\u0072\u0042"
	case 15:
		return "\u0072\u0043\u0074\u0072\u0043\u0068"
	case 16:
		return "\u0072C\u0074\u0072\u0044\u0065\u0073"
	}
	return ""
}

type CT_DataModel struct {
	PtLst  *CT_PtList
	CxnLst *CT_CxnList
	Bg     *_ca.CT_BackgroundFormatting
	Whole  *_ca.CT_WholeE2oFormatting
	ExtLst *_ca.CT_OfficeArtExtensionList
}

func (_dcgag *CT_Pt) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _bbea := range start.Attr {
		if _bbea.Name.Local == "\u006do\u0064\u0065\u006c\u0049\u0064" {
			_gcdfb, _bbfb := ParseUnionST_ModelId(_bbea.Value)
			if _bbfb != nil {
				return _bbfb
			}
			_dcgag.ModelIdAttr = _gcdfb
			continue
		}
		if _bbea.Name.Local == "\u0074\u0079\u0070\u0065" {
			_dcgag.TypeAttr.UnmarshalXMLAttr(_bbea)
			continue
		}
		if _bbea.Name.Local == "\u0063\u0078\u006eI\u0064" {
			_aaae, _eaae := ParseUnionST_ModelId(_bbea.Value)
			if _eaae != nil {
				return _eaae
			}
			_dcgag.CxnIdAttr = &_aaae
			continue
		}
	}
_ffeg:
	for {
		_ccff, _bbgg := d.Token()
		if _bbgg != nil {
			return _bbgg
		}
		switch _cggg := _ccff.(type) {
		case _d.StartElement:
			switch _cggg.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0053e\u0074"}:
				_dcgag.PrSet = NewCT_ElemPropSet()
				if _baa := d.DecodeElement(_dcgag.PrSet, &_cggg); _baa != nil {
					return _baa
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0070\u0050\u0072"}:
				_dcgag.SpPr = _ca.NewCT_ShapeProperties()
				if _ggag := d.DecodeElement(_dcgag.SpPr, &_cggg); _ggag != nil {
					return _ggag
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074"}:
				_dcgag.T = _ca.NewCT_TextBody()
				if _agbe := d.DecodeElement(_dcgag.T, &_cggg); _agbe != nil {
					return _agbe
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_dcgag.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _gbdcd := d.DecodeElement(_dcgag.ExtLst, &_cggg); _gbdcd != nil {
					return _gbdcd
				}
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0043\u0054\u005fP\u0074 \u0025\u0076", _cggg.Name)
				if _feee := d.Skip(); _feee != nil {
					return _feee
				}
			}
		case _d.EndElement:
			break _ffeg
		case _d.CharData:
		}
	}
	return nil
}

type CT_Adj struct {
	IdxAttr uint32
	ValAttr float64
}

func (_ccdge ST_ConstraintType) ValidateWithPath(path string) error {
	switch _ccdge {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ccdge))
	}
	return nil
}

// Validate validates the CT_StyleLabel and its children
func (_abgag *CT_StyleLabel) Validate() error {
	return _abgag.ValidateWithPath("\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c")
}

const (
	ST_HueDirUnset ST_HueDir = 0
	ST_HueDirCw    ST_HueDir = 1
	ST_HueDirCcw   ST_HueDir = 2
)

// Validate validates the CT_StyleDefinitionHeaderLst and its children
func (_dgad *CT_StyleDefinitionHeaderLst) Validate() error {
	return _dgad.ValidateWithPath("C\u0054\u005f\u0053\u0074\u0079\u006ce\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006fn\u0048\u0065\u0061d\u0065r\u004c\u0073\u0074")
}

// ValidateWithPath validates the CT_ElemPropSet and its children, prefixing error messages with path
func (_caae *CT_ElemPropSet) ValidateWithPath(path string) error {
	if _caae.PresAssocIDAttr != nil {
		if _eedeb := _caae.PresAssocIDAttr.ValidateWithPath(path + "\u002f\u0050r\u0065\u0073\u0041s\u0073\u006f\u0063\u0049\u0044\u0041\u0074\u0074\u0072"); _eedeb != nil {
			return _eedeb
		}
	}
	if _caae.CustScaleXAttr != nil {
		if _fcfc := _caae.CustScaleXAttr.ValidateWithPath(path + "\u002fC\u0075s\u0074\u0053\u0063\u0061\u006c\u0065\u0058\u0041\u0074\u0074\u0072"); _fcfc != nil {
			return _fcfc
		}
	}
	if _caae.CustScaleYAttr != nil {
		if _dcacf := _caae.CustScaleYAttr.ValidateWithPath(path + "\u002fC\u0075s\u0074\u0053\u0063\u0061\u006c\u0065\u0059\u0041\u0074\u0074\u0072"); _dcacf != nil {
			return _dcacf
		}
	}
	if _caae.CustLinFactXAttr != nil {
		if _adee := _caae.CustLinFactXAttr.ValidateWithPath(path + "\u002f\u0043\u0075\u0073\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074X\u0041\u0074\u0074\u0072"); _adee != nil {
			return _adee
		}
	}
	if _caae.CustLinFactYAttr != nil {
		if _gbfcf := _caae.CustLinFactYAttr.ValidateWithPath(path + "\u002f\u0043\u0075\u0073\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074Y\u0041\u0074\u0074\u0072"); _gbfcf != nil {
			return _gbfcf
		}
	}
	if _caae.CustLinFactNeighborXAttr != nil {
		if _afge := _caae.CustLinFactNeighborXAttr.ValidateWithPath(path + "\u002fC\u0075\u0073\u0074\u004ci\u006e\u0046\u0061\u0063\u0074N\u0065i\u0067h\u0062\u006f\u0072\u0058\u0041\u0074\u0074r"); _afge != nil {
			return _afge
		}
	}
	if _caae.CustLinFactNeighborYAttr != nil {
		if _dcdc := _caae.CustLinFactNeighborYAttr.ValidateWithPath(path + "\u002fC\u0075\u0073\u0074\u004ci\u006e\u0046\u0061\u0063\u0074N\u0065i\u0067h\u0062\u006f\u0072\u0059\u0041\u0074\u0074r"); _dcdc != nil {
			return _dcdc
		}
	}
	if _caae.CustRadScaleRadAttr != nil {
		if _beed := _caae.CustRadScaleRadAttr.ValidateWithPath(path + "/\u0043u\u0073\u0074\u0052\u0061\u0064\u0053\u0063\u0061l\u0065\u0052\u0061\u0064At\u0074\u0072"); _beed != nil {
			return _beed
		}
	}
	if _caae.CustRadScaleIncAttr != nil {
		if _ecbc := _caae.CustRadScaleIncAttr.ValidateWithPath(path + "/\u0043u\u0073\u0074\u0052\u0061\u0064\u0053\u0063\u0061l\u0065\u0049\u006e\u0063At\u0074\u0072"); _ecbc != nil {
			return _ecbc
		}
	}
	if _caae.PresLayoutVars != nil {
		if _dceb := _caae.PresLayoutVars.ValidateWithPath(path + "\u002fP\u0072e\u0073\u004c\u0061\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0073"); _dceb != nil {
			return _dceb
		}
	}
	if _caae.Style != nil {
		if _facf := _caae.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _facf != nil {
			return _facf
		}
	}
	return nil
}
func (_ffag *ST_Offset) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fbde, _gbga := d.Token()
	if _gbga != nil {
		return _gbga
	}
	if _dbgc, _gaaeg := _fbde.(_d.EndElement); _gaaeg && _dbgc.Name == start.Name {
		*_ffag = 1
		return nil
	}
	if _dafgf, _bead := _fbde.(_d.CharData); !_bead {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fbde)
	} else {
		switch string(_dafgf) {
		case "":
			*_ffag = 0
		case "\u0063\u0074\u0072":
			*_ffag = 1
		case "\u006f\u0066\u0066":
			*_ffag = 2
		}
	}
	_fbde, _gbga = d.Token()
	if _gbga != nil {
		return _gbga
	}
	if _bfcea, _abfgg := _fbde.(_d.EndElement); _abfgg && _bfcea.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fbde)
}
func (_eafa ST_AnimLvlStr) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_eafa.String(), start)
}

type StyleDefHdr struct{ CT_StyleDefinitionHeader }
type CT_LayoutNode struct {
	NameAttr     *string
	StyleLblAttr *string
	ChOrderAttr  ST_ChildOrderType
	MoveWithAttr *string
	Alg          []*CT_Algorithm
	Shape        []*CT_Shape
	PresOf       []*CT_PresentationOf
	ConstrLst    []*CT_Constraints
	RuleLst      []*CT_Rules
	VarLst       []*CT_LayoutVariablePropertySet
	ForEach      []*CT_ForEach
	LayoutNode   []*CT_LayoutNode
	Choose       []*CT_Choose
	ExtLst       []*_ca.CT_OfficeArtExtensionList
}

// Validate validates the CT_Description and its children
func (_fdb *CT_Description) Validate() error {
	return _fdb.ValidateWithPath("\u0043\u0054\u005f\u0044\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e")
}

// Validate validates the CT_LayoutVariablePropertySet and its children
func (_cfgf *CT_LayoutVariablePropertySet) Validate() error {
	return _cfgf.ValidateWithPath("\u0043\u0054\u005f\u004ca\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0069\u0061\u0062l\u0065P\u0072\u006f\u0070\u0065\u0072\u0074\u0079S\u0065\u0074")
}
func (_ddcd *ST_ConnectorRouting) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ddcd = 0
	case "\u0073\u0074\u0072\u0061":
		*_ddcd = 1
	case "\u0062\u0065\u006e\u0064":
		*_ddcd = 2
	case "\u0063\u0075\u0072v\u0065":
		*_ddcd = 3
	case "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e":
		*_ddcd = 4
	}
	return nil
}
func (_dddgg ST_HierarchyAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_agbdf := _d.Attr{}
	_agbdf.Name = name
	switch _dddgg {
	case ST_HierarchyAlignmentUnset:
		_agbdf.Value = ""
	case ST_HierarchyAlignmentTL:
		_agbdf.Value = "\u0074\u004c"
	case ST_HierarchyAlignmentTR:
		_agbdf.Value = "\u0074\u0052"
	case ST_HierarchyAlignmentTCtrCh:
		_agbdf.Value = "\u0074\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentTCtrDes:
		_agbdf.Value = "\u0074C\u0074\u0072\u0044\u0065\u0073"
	case ST_HierarchyAlignmentBL:
		_agbdf.Value = "\u0062\u004c"
	case ST_HierarchyAlignmentBR:
		_agbdf.Value = "\u0062\u0052"
	case ST_HierarchyAlignmentBCtrCh:
		_agbdf.Value = "\u0062\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentBCtrDes:
		_agbdf.Value = "\u0062C\u0074\u0072\u0044\u0065\u0073"
	case ST_HierarchyAlignmentLT:
		_agbdf.Value = "\u006c\u0054"
	case ST_HierarchyAlignmentLB:
		_agbdf.Value = "\u006c\u0042"
	case ST_HierarchyAlignmentLCtrCh:
		_agbdf.Value = "\u006c\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentLCtrDes:
		_agbdf.Value = "\u006cC\u0074\u0072\u0044\u0065\u0073"
	case ST_HierarchyAlignmentRT:
		_agbdf.Value = "\u0072\u0054"
	case ST_HierarchyAlignmentRB:
		_agbdf.Value = "\u0072\u0042"
	case ST_HierarchyAlignmentRCtrCh:
		_agbdf.Value = "\u0072\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentRCtrDes:
		_agbdf.Value = "\u0072C\u0074\u0072\u0044\u0065\u0073"
	}
	return _agbdf, nil
}
func (_gage *ST_ChildDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cedf, _cdcfa := d.Token()
	if _cdcfa != nil {
		return _cdcfa
	}
	if _dbdbe, _agaef := _cedf.(_d.EndElement); _agaef && _dbdbe.Name == start.Name {
		*_gage = 1
		return nil
	}
	if _gaacg, _geaf := _cedf.(_d.CharData); !_geaf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cedf)
	} else {
		switch string(_gaacg) {
		case "":
			*_gage = 0
		case "\u0068\u006f\u0072\u007a":
			*_gage = 1
		case "\u0076\u0065\u0072\u0074":
			*_gage = 2
		}
	}
	_cedf, _cdcfa = d.Token()
	if _cdcfa != nil {
		return _cdcfa
	}
	if _bgefb, _deedd := _cedf.(_d.EndElement); _deedd && _bgefb.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cedf)
}

type CT_AdjLst struct{ Adj []*CT_Adj }

func (_dccca *ST_VariableType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_aceg, _geba := d.Token()
	if _geba != nil {
		return _geba
	}
	if _gggd, _faac := _aceg.(_d.EndElement); _faac && _gggd.Name == start.Name {
		*_dccca = 1
		return nil
	}
	if _eabcf, _gfacg := _aceg.(_d.CharData); !_gfacg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _aceg)
	} else {
		switch string(_eabcf) {
		case "":
			*_dccca = 0
		case "\u006e\u006f\u006e\u0065":
			*_dccca = 1
		case "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074":
			*_dccca = 2
		case "\u0063\u0068\u004da\u0078":
			*_dccca = 3
		case "\u0063\u0068\u0050\u0072\u0065\u0066":
			*_dccca = 4
		case "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064":
			*_dccca = 5
		case "\u0064\u0069\u0072":
			*_dccca = 6
		case "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068":
			*_dccca = 7
		case "\u0061n\u0069\u006d\u004f\u006e\u0065":
			*_dccca = 8
		case "\u0061n\u0069\u006d\u004c\u0076\u006c":
			*_dccca = 9
		case "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073":
			*_dccca = 10
		}
	}
	_aceg, _geba = d.Token()
	if _geba != nil {
		return _geba
	}
	if _acbca, _gcaed := _aceg.(_d.EndElement); _gcaed && _acbca.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _aceg)
}

type ColorsDefHdr struct{ CT_ColorTransformHeader }

func (_bbeaee *ST_NodeVerticalAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_bbeaee = 0
	case "\u0074":
		*_bbeaee = 1
	case "\u006d\u0069\u0064":
		*_bbeaee = 2
	case "\u0062":
		*_bbeaee = 3
	}
	return nil
}
func (_befaa ST_BendPoint) ValidateWithPath(path string) error {
	switch _befaa {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_befaa))
	}
	return nil
}
func (_aeca *CT_LayoutVariablePropertySet) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_aace:
	for {
		_faag, _ecbed := d.Token()
		if _ecbed != nil {
			return _ecbed
		}
		switch _dgfgd := _faag.(type) {
		case _d.StartElement:
			switch _dgfgd.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"}:
				_aeca.OrgChart = NewCT_OrgChart()
				if _cege := d.DecodeElement(_aeca.OrgChart, &_dgfgd); _cege != nil {
					return _cege
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u004da\u0078"}:
				_aeca.ChMax = NewCT_ChildMax()
				if _cadf := d.DecodeElement(_aeca.ChMax, &_dgfgd); _cadf != nil {
					return _cadf
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u0050\u0072\u0065\u0066"}:
				_aeca.ChPref = NewCT_ChildPref()
				if _caac := d.DecodeElement(_aeca.ChPref, &_dgfgd); _caac != nil {
					return _caac
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0062\u0075\u006c\u006c\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064"}:
				_aeca.BulletEnabled = NewCT_BulletEnabled()
				if _bad := d.DecodeElement(_aeca.BulletEnabled, &_dgfgd); _bad != nil {
					return _bad
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0069\u0072"}:
				_aeca.Dir = NewCT_Direction()
				if _afbf := d.DecodeElement(_aeca.Dir, &_dgfgd); _afbf != nil {
					return _afbf
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"}:
				_aeca.HierBranch = NewCT_HierBranchStyle()
				if _bcbgd := d.DecodeElement(_aeca.HierBranch, &_dgfgd); _bcbgd != nil {
					return _bcbgd
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061n\u0069\u006d\u004f\u006e\u0065"}:
				_aeca.AnimOne = NewCT_AnimOne()
				if _feef := d.DecodeElement(_aeca.AnimOne, &_dgfgd); _feef != nil {
					return _feef
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061n\u0069\u006d\u004c\u0076\u006c"}:
				_aeca.AnimLvl = NewCT_AnimLvl()
				if _cfafa := d.DecodeElement(_aeca.AnimLvl, &_dgfgd); _cfafa != nil {
					return _cfafa
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"}:
				_aeca.ResizeHandles = NewCT_ResizeHandles()
				if _dddg := d.DecodeElement(_aeca.ResizeHandles, &_dgfgd); _dddg != nil {
					return _dddg
				}
			default:
				_af.Log("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020C\u0054\u005f\u004c\u0061\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0069\u0061\u0062\u006c\u0065P\u0072\u006fpe\u0072\u0074\u0079S\u0065\u0074\u0020\u0025\u0076", _dgfgd.Name)
				if _abfa := d.Skip(); _abfa != nil {
					return _abfa
				}
			}
		case _d.EndElement:
			break _aace
		case _d.CharData:
		}
	}
	return nil
}
func (_dafb ST_ConnectorPoint) ValidateWithPath(path string) error {
	switch _dafb {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dafb))
	}
	return nil
}
func (_cadec ST_FunctionType) String() string {
	switch _cadec {
	case 0:
		return ""
	case 1:
		return "\u0063\u006e\u0074"
	case 2:
		return "\u0070\u006f\u0073"
	case 3:
		return "\u0072\u0065\u0076\u0050\u006f\u0073"
	case 4:
		return "\u0070o\u0073\u0045\u0076\u0065\u006e"
	case 5:
		return "\u0070\u006f\u0073\u004f\u0064\u0064"
	case 6:
		return "\u0076\u0061\u0072"
	case 7:
		return "\u0064\u0065\u0070t\u0068"
	case 8:
		return "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068"
	}
	return ""
}
func (_ebabe ST_ConstraintType) String() string {
	switch _ebabe {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066"
	case 3:
		return "\u0062e\u0067\u004d\u0061\u0072\u0067"
	case 4:
		return "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074"
	case 5:
		return "\u0062\u0065\u0067\u0050\u0061\u0064"
	case 6:
		return "\u0062"
	case 7:
		return "\u0062\u004d\u0061r\u0067"
	case 8:
		return "\u0062\u004f\u0066\u0066"
	case 9:
		return "\u0063\u0074\u0072\u0058"
	case 10:
		return "\u0063t\u0072\u0058\u004f\u0066\u0066"
	case 11:
		return "\u0063\u0074\u0072\u0059"
	case 12:
		return "\u0063t\u0072\u0059\u004f\u0066\u0066"
	case 13:
		return "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074"
	case 14:
		return "\u0064\u0069\u0061\u006d"
	case 15:
		return "\u0065n\u0064\u004d\u0061\u0072\u0067"
	case 16:
		return "\u0065\u006e\u0064\u0050\u0061\u0064"
	case 17:
		return "\u0068"
	case 18:
		return "\u0068\u0041\u0072\u0048"
	case 19:
		return "\u0068\u004f\u0066\u0066"
	case 20:
		return "\u006c"
	case 21:
		return "\u006c\u004d\u0061r\u0067"
	case 22:
		return "\u006c\u004f\u0066\u0066"
	case 23:
		return "\u0072"
	case 24:
		return "\u0072\u004d\u0061r\u0067"
	case 25:
		return "\u0072\u004f\u0066\u0066"
	case 26:
		return "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a"
	case 27:
		return "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f"
	case 28:
		return "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z"
	case 29:
		return "\u0073\u0069\u0062S\u0070"
	case 30:
		return "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070"
	case 31:
		return "\u0073\u0070"
	case 32:
		return "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k"
	case 33:
		return "\u0074"
	case 34:
		return "\u0074\u004d\u0061r\u0067"
	case 35:
		return "\u0074\u004f\u0066\u0066"
	case 36:
		return "\u0075\u0073\u0065r\u0041"
	case 37:
		return "\u0075\u0073\u0065r\u0042"
	case 38:
		return "\u0075\u0073\u0065r\u0043"
	case 39:
		return "\u0075\u0073\u0065r\u0044"
	case 40:
		return "\u0075\u0073\u0065r\u0045"
	case 41:
		return "\u0075\u0073\u0065r\u0046"
	case 42:
		return "\u0075\u0073\u0065r\u0047"
	case 43:
		return "\u0075\u0073\u0065r\u0048"
	case 44:
		return "\u0075\u0073\u0065r\u0049"
	case 45:
		return "\u0075\u0073\u0065r\u004a"
	case 46:
		return "\u0075\u0073\u0065r\u004b"
	case 47:
		return "\u0075\u0073\u0065r\u004c"
	case 48:
		return "\u0075\u0073\u0065r\u004d"
	case 49:
		return "\u0075\u0073\u0065r\u004e"
	case 50:
		return "\u0075\u0073\u0065r\u004f"
	case 51:
		return "\u0075\u0073\u0065r\u0050"
	case 52:
		return "\u0075\u0073\u0065r\u0051"
	case 53:
		return "\u0075\u0073\u0065r\u0052"
	case 54:
		return "\u0075\u0073\u0065r\u0053"
	case 55:
		return "\u0075\u0073\u0065r\u0054"
	case 56:
		return "\u0075\u0073\u0065r\u0055"
	case 57:
		return "\u0075\u0073\u0065r\u0056"
	case 58:
		return "\u0075\u0073\u0065r\u0057"
	case 59:
		return "\u0075\u0073\u0065r\u0058"
	case 60:
		return "\u0075\u0073\u0065r\u0059"
	case 61:
		return "\u0075\u0073\u0065r\u005a"
	case 62:
		return "\u0077"
	case 63:
		return "\u0077\u0041\u0072\u0048"
	case 64:
		return "\u0077\u004f\u0066\u0066"
	}
	return ""
}
func (_deaae ST_LinearDirection) String() string {
	switch _deaae {
	case 0:
		return ""
	case 1:
		return "\u0066\u0072\u006fm\u004c"
	case 2:
		return "\u0066\u0072\u006fm\u0052"
	case 3:
		return "\u0066\u0072\u006fm\u0054"
	case 4:
		return "\u0066\u0072\u006fm\u0042"
	}
	return ""
}

// ValidateWithPath validates the CT_OrgChart and its children, prefixing error messages with path
func (_cadd *CT_OrgChart) ValidateWithPath(path string) error { return nil }
func (_ebdfb ST_ParameterId) Validate() error                 { return _ebdfb.ValidateWithPath("") }
func (_fdddb *ST_GrowDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_eddg, _babgb := d.Token()
	if _babgb != nil {
		return _babgb
	}
	if _ecfae, _ecabg := _eddg.(_d.EndElement); _ecabg && _ecfae.Name == start.Name {
		*_fdddb = 1
		return nil
	}
	if _cddb, _bfbbgb := _eddg.(_d.CharData); !_bfbbgb {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eddg)
	} else {
		switch string(_cddb) {
		case "":
			*_fdddb = 0
		case "\u0074\u004c":
			*_fdddb = 1
		case "\u0074\u0052":
			*_fdddb = 2
		case "\u0062\u004c":
			*_fdddb = 3
		case "\u0062\u0052":
			*_fdddb = 4
		}
	}
	_eddg, _babgb = d.Token()
	if _babgb != nil {
		return _babgb
	}
	if _ccac, _abcbd := _eddg.(_d.EndElement); _abcbd && _ccac.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eddg)
}
func (_bgcgg ST_ConnectorPoint) Validate() error { return _bgcgg.ValidateWithPath("") }
func (_adbbc ST_FunctionArgument) String() string {
	if _adbbc.ST_VariableType != ST_VariableTypeUnset {
		return _adbbc.ST_VariableType.String()
	}
	return ""
}

// ValidateWithPath validates the CT_SDName and its children, prefixing error messages with path
func (_bdbff *CT_SDName) ValidateWithPath(path string) error { return nil }
func (_abcd *CT_Constraints) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_fge:
	for {
		_ecff, _gcbb := d.Token()
		if _gcbb != nil {
			return _gcbb
		}
		switch _bea := _ecff.(type) {
		case _d.StartElement:
			switch _bea.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u006f\u006e\u0073\u0074\u0072"}:
				_bdaf := NewCT_Constraint()
				if _gdbf := d.DecodeElement(_bdaf, &_bea); _gdbf != nil {
					return _gdbf
				}
				_abcd.Constr = append(_abcd.Constr, _bdaf)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061i\u006et\u0073\u0020\u0025\u0076", _bea.Name)
				if _gcag := d.Skip(); _gcag != nil {
					return _gcag
				}
			}
		case _d.EndElement:
			break _fge
		case _d.CharData:
		}
	}
	return nil
}
func (_ebee *ST_CenterShapeMapping) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ebee = 0
	case "\u006e\u006f\u006e\u0065":
		*_ebee = 1
	case "\u0066\u004e\u006fd\u0065":
		*_ebee = 2
	}
	return nil
}

type CT_Parameter struct {
	TypeAttr ST_ParameterId
	ValAttr  ST_ParameterVal
}

// Validate validates the CT_BulletEnabled and its children
func (_fec *CT_BulletEnabled) Validate() error {
	return _fec.ValidateWithPath("\u0043\u0054_\u0042\u0075\u006cl\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064")
}
func (_gbaaa *ST_RotationPath) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gcceg, _cecb := d.Token()
	if _cecb != nil {
		return _cecb
	}
	if _facag, _aafe := _gcceg.(_d.EndElement); _aafe && _facag.Name == start.Name {
		*_gbaaa = 1
		return nil
	}
	if _eddf, _ccgf := _gcceg.(_d.CharData); !_ccgf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gcceg)
	} else {
		switch string(_eddf) {
		case "":
			*_gbaaa = 0
		case "\u006e\u006f\u006e\u0065":
			*_gbaaa = 1
		case "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h":
			*_gbaaa = 2
		}
	}
	_gcceg, _cecb = d.Token()
	if _cecb != nil {
		return _cecb
	}
	if _fdfgc, _agddg := _gcceg.(_d.EndElement); _agddg && _fdfgc.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gcceg)
}
func (_bggg ST_ChildDirection) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_bggg.String(), start)
}

// Validate validates the CT_SDCategory and its children
func (_ggb *CT_SDCategory) Validate() error {
	return _ggb.ValidateWithPath("\u0043\u0054\u005f\u0053\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079")
}
func (_gdfb ST_FallbackDimension) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_dccgc := _d.Attr{}
	_dccgc.Name = name
	switch _gdfb {
	case ST_FallbackDimensionUnset:
		_dccgc.Value = ""
	case ST_FallbackDimension1D:
		_dccgc.Value = "\u0031\u0044"
	case ST_FallbackDimension2D:
		_dccgc.Value = "\u0032\u0044"
	}
	return _dccgc, nil
}
func NewRelIds() *RelIds { _bebba := &RelIds{}; _bebba.CT_RelIds = *NewCT_RelIds(); return _bebba }
func (_ebf *CT_CTDescription) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _ebf.LangAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_ebf.LangAttr)})
	}
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _ebf.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

type ST_PyramidAccentTextMargin byte

func (_acef *CT_PtList) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_egfa:
	for {
		_egcb, _bcbf := d.Token()
		if _bcbf != nil {
			return _bcbf
		}
		switch _ffffg := _egcb.(type) {
		case _d.StartElement:
			switch _ffffg.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0074"}:
				_feafe := NewCT_Pt()
				if _cbbd := d.DecodeElement(_feafe, &_ffffg); _cbbd != nil {
					return _cbbd
				}
				_acef.Pt = append(_acef.Pt, _feafe)
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0050\u0074\u004ci\u0073\u0074 \u0025\u0076", _ffffg.Name)
				if _agfdf := d.Skip(); _agfdf != nil {
					return _agfdf
				}
			}
		case _d.EndElement:
			break _egfa
		case _d.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SDCategories and its children
func (_ffaa *CT_SDCategories) Validate() error {
	return _ffaa.ValidateWithPath("\u0043T\u005fS\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073")
}
func (_abc *CT_Choose) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _ccefa := range start.Attr {
		if _ccefa.Name.Local == "\u006e\u0061\u006d\u0065" {
			_bdec, _gfdeb := _ccefa.Value, error(nil)
			if _gfdeb != nil {
				return _gfdeb
			}
			_abc.NameAttr = &_bdec
			continue
		}
	}
_ffe:
	for {
		_fdde, _cacfd := d.Token()
		if _cacfd != nil {
			return _cacfd
		}
		switch _gbcf := _fdde.(type) {
		case _d.StartElement:
			switch _gbcf.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0069\u0066"}:
				_eede := NewCT_When()
				if _eaag := d.DecodeElement(_eede, &_gbcf); _eaag != nil {
					return _eaag
				}
				_abc.If = append(_abc.If, _eede)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u006c\u0073\u0065"}:
				_abc.Else = NewCT_Otherwise()
				if _ded := d.DecodeElement(_abc.Else, &_gbcf); _ded != nil {
					return _ded
				}
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0043\u0068\u006fo\u0073\u0065 \u0025\u0076", _gbcf.Name)
				if _bce := d.Skip(); _bce != nil {
					return _bce
				}
			}
		case _d.EndElement:
			break _ffe
		case _d.CharData:
		}
	}
	return nil
}
func NewCT_DiagramDefinitionHeaderLst() *CT_DiagramDefinitionHeaderLst {
	_cccc := &CT_DiagramDefinitionHeaderLst{}
	return _cccc
}
func (_aeece ST_TextDirection) Validate() error { return _aeece.ValidateWithPath("") }

// Validate validates the CT_RelIds and its children
func (_ecfg *CT_RelIds) Validate() error {
	return _ecfg.ValidateWithPath("\u0043T\u005f\u0052\u0065\u006c\u0049\u0064s")
}
func (_ffaf ST_AnimLvlStr) ValidateWithPath(path string) error {
	switch _ffaf {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ffaf))
	}
	return nil
}
func (_gbdge ST_HierBranchStyle) ValidateWithPath(path string) error {
	switch _gbdge {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gbdge))
	}
	return nil
}
func (_geae ST_ContinueDirection) Validate() error { return _geae.ValidateWithPath("") }
func (_eedce *ST_ConnectorRouting) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_dbac, _eagbe := d.Token()
	if _eagbe != nil {
		return _eagbe
	}
	if _ecgcfc, _ccfea := _dbac.(_d.EndElement); _ccfea && _ecgcfc.Name == start.Name {
		*_eedce = 1
		return nil
	}
	if _cfeee, _deca := _dbac.(_d.CharData); !_deca {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dbac)
	} else {
		switch string(_cfeee) {
		case "":
			*_eedce = 0
		case "\u0073\u0074\u0072\u0061":
			*_eedce = 1
		case "\u0062\u0065\u006e\u0064":
			*_eedce = 2
		case "\u0063\u0075\u0072v\u0065":
			*_eedce = 3
		case "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e":
			*_eedce = 4
		}
	}
	_dbac, _eagbe = d.Token()
	if _eagbe != nil {
		return _eagbe
	}
	if _bgfe, _dcccg := _dbac.(_d.EndElement); _dcccg && _bgfe.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dbac)
}
func (_fccdg ST_HierarchyAlignment) ValidateWithPath(path string) error {
	switch _fccdg {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fccdg))
	}
	return nil
}

const (
	ST_PyramidAccentPositionUnset ST_PyramidAccentPosition = 0
	ST_PyramidAccentPositionBef   ST_PyramidAccentPosition = 1
	ST_PyramidAccentPositionAft   ST_PyramidAccentPosition = 2
)

func NewCT_CTDescription() *CT_CTDescription { _afbd := &CT_CTDescription{}; return _afbd }
func NewColorsDef() *ColorsDef {
	_gecf := &ColorsDef{}
	_gecf.CT_ColorTransform = *NewCT_ColorTransform()
	return _gecf
}
func (_acabd *ST_ContinueDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_fdfad, _beacbd := d.Token()
	if _beacbd != nil {
		return _beacbd
	}
	if _bdee, _fdfga := _fdfad.(_d.EndElement); _fdfga && _bdee.Name == start.Name {
		*_acabd = 1
		return nil
	}
	if _deeg, _beebc := _fdfad.(_d.CharData); !_beebc {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fdfad)
	} else {
		switch string(_deeg) {
		case "":
			*_acabd = 0
		case "\u0072\u0065\u0076\u0044\u0069\u0072":
			*_acabd = 1
		case "\u0073a\u006d\u0065\u0044\u0069\u0072":
			*_acabd = 2
		}
	}
	_fdfad, _beacbd = d.Token()
	if _beacbd != nil {
		return _beacbd
	}
	if _bbafa, _afedd := _fdfad.(_d.EndElement); _afedd && _bbafa.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fdfad)
}

const (
	ST_ContinueDirectionUnset   ST_ContinueDirection = 0
	ST_ContinueDirectionRevDir  ST_ContinueDirection = 1
	ST_ContinueDirectionSameDir ST_ContinueDirection = 2
)

// Validate validates the CT_SDDescription and its children
func (_bfdd *CT_SDDescription) Validate() error {
	return _bfdd.ValidateWithPath("\u0043\u0054_\u0053\u0044\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e")
}
func (_ceef ST_AnimOneStr) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_ceef.String(), start)
}
func (_fgbac ST_SecondaryChildAlignment) Validate() error { return _fgbac.ValidateWithPath("") }

// Validate validates the CT_ColorTransform and its children
func (_aacb *CT_ColorTransform) Validate() error {
	return _aacb.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006es\u0066\u006f\u0072\u006d")
}
func (_eggbg ST_GrowDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_bgag := _d.Attr{}
	_bgag.Name = name
	switch _eggbg {
	case ST_GrowDirectionUnset:
		_bgag.Value = ""
	case ST_GrowDirectionTL:
		_bgag.Value = "\u0074\u004c"
	case ST_GrowDirectionTR:
		_bgag.Value = "\u0074\u0052"
	case ST_GrowDirectionBL:
		_bgag.Value = "\u0062\u004c"
	case ST_GrowDirectionBR:
		_bgag.Value = "\u0062\u0052"
	}
	return _bgag, nil
}
func NewCT_DiagramDefinitionHeader() *CT_DiagramDefinitionHeader {
	_ffga := &CT_DiagramDefinitionHeader{}
	return _ffga
}
func (_defc *CT_LayoutNode) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _egge := range start.Attr {
		if _egge.Name.Local == "\u006e\u0061\u006d\u0065" {
			_ecgfd, _eeaa := _egge.Value, error(nil)
			if _eeaa != nil {
				return _eeaa
			}
			_defc.NameAttr = &_ecgfd
			continue
		}
		if _egge.Name.Local == "\u0063h\u004f\u0072\u0064\u0065\u0072" {
			_defc.ChOrderAttr.UnmarshalXMLAttr(_egge)
			continue
		}
		if _egge.Name.Local == "\u006d\u006f\u0076\u0065\u0057\u0069\u0074\u0068" {
			_fgdgg, _gfgde := _egge.Value, error(nil)
			if _gfgde != nil {
				return _gfgde
			}
			_defc.MoveWithAttr = &_fgdgg
			continue
		}
		if _egge.Name.Local == "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c" {
			_daaa, _edge := _egge.Value, error(nil)
			if _edge != nil {
				return _edge
			}
			_defc.StyleLblAttr = &_daaa
			continue
		}
	}
_cadg:
	for {
		_gafc, _eaac := d.Token()
		if _eaac != nil {
			return _eaac
		}
		switch _daeed := _gafc.(type) {
		case _d.StartElement:
			switch _daeed.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_adbf := NewCT_Algorithm()
				if _faae := d.DecodeElement(_adbf, &_daeed); _faae != nil {
					return _faae
				}
				_defc.Alg = append(_defc.Alg, _adbf)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_fcff := NewCT_Shape()
				if _cdede := d.DecodeElement(_fcff, &_daeed); _cdede != nil {
					return _cdede
				}
				_defc.Shape = append(_defc.Shape, _fcff)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_bebc := NewCT_PresentationOf()
				if _fabb := d.DecodeElement(_bebc, &_daeed); _fabb != nil {
					return _fabb
				}
				_defc.PresOf = append(_defc.PresOf, _bebc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_fbdd := NewCT_Constraints()
				if _fdga := d.DecodeElement(_fbdd, &_daeed); _fdga != nil {
					return _fdga
				}
				_defc.ConstrLst = append(_defc.ConstrLst, _fbdd)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_cdcdc := NewCT_Rules()
				if _ccddf := d.DecodeElement(_cdcdc, &_daeed); _ccddf != nil {
					return _ccddf
				}
				_defc.RuleLst = append(_defc.RuleLst, _cdcdc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0076\u0061\u0072\u004c\u0073\u0074"}:
				_bcff := NewCT_LayoutVariablePropertySet()
				if _bgdd := d.DecodeElement(_bcff, &_daeed); _bgdd != nil {
					return _bgdd
				}
				_defc.VarLst = append(_defc.VarLst, _bcff)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_gafe := NewCT_ForEach()
				if _bfaed := d.DecodeElement(_gafe, &_daeed); _bfaed != nil {
					return _bfaed
				}
				_defc.ForEach = append(_defc.ForEach, _gafe)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_eefdc := NewCT_LayoutNode()
				if _gbca := d.DecodeElement(_eefdc, &_daeed); _gbca != nil {
					return _gbca
				}
				_defc.LayoutNode = append(_defc.LayoutNode, _eefdc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_cgdb := NewCT_Choose()
				if _dbfb := d.DecodeElement(_cgdb, &_daeed); _dbfb != nil {
					return _dbfb
				}
				_defc.Choose = append(_defc.Choose, _cgdb)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bfbe := _ca.NewCT_OfficeArtExtensionList()
				if _aeabe := d.DecodeElement(_bfbe, &_daeed); _aeabe != nil {
					return _aeabe
				}
				_defc.ExtLst = append(_defc.ExtLst, _bfbe)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004ca\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065 \u0025\u0076", _daeed.Name)
				if _cagd := d.Skip(); _cagd != nil {
					return _cagd
				}
			}
		case _d.EndElement:
			break _cadg
		case _d.CharData:
		}
	}
	return nil
}

type ST_ChildDirection byte

// Validate validates the RelIds and its children
func (_dedcc *RelIds) Validate() error {
	return _dedcc.ValidateWithPath("\u0052\u0065\u006c\u0049\u0064\u0073")
}
func (_abdef ST_AnimLvlStr) Validate() error { return _abdef.ValidateWithPath("") }
func (_ccfb ST_ConnectorDimension) ValidateWithPath(path string) error {
	switch _ccfb {
	case 0, 1, 2, 3:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ccfb))
	}
	return nil
}
func (_adbbb *ST_OutputShapeType) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bgggc, _bfgfg := d.Token()
	if _bfgfg != nil {
		return _bfgfg
	}
	if _gcbfd, _ffbge := _bgggc.(_d.EndElement); _ffbge && _gcbfd.Name == start.Name {
		*_adbbb = 1
		return nil
	}
	if _bcgaa, _gdfaf := _bgggc.(_d.CharData); !_gdfaf {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bgggc)
	} else {
		switch string(_bcgaa) {
		case "":
			*_adbbb = 0
		case "\u006e\u006f\u006e\u0065":
			*_adbbb = 1
		case "\u0063\u006f\u006e\u006e":
			*_adbbb = 2
		}
	}
	_bgggc, _bfgfg = d.Token()
	if _bfgfg != nil {
		return _bfgfg
	}
	if _defcg, _aaad := _bgggc.(_d.EndElement); _aaad && _defcg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bgggc)
}

// Validate validates the CT_Constraints and its children
func (_bdf *CT_Constraints) Validate() error {
	return _bdf.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072a\u0069\u006e\u0074\u0073")
}
func (_dacee *ST_ChildDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_dacee = 0
	case "\u0068\u006f\u0072\u007a":
		*_dacee = 1
	case "\u0076\u0065\u0072\u0074":
		*_dacee = 2
	}
	return nil
}

type ST_UnsignedInts []uint32

func (_dafg *ST_TextAnchorVertical) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_agbba, _gcbe := d.Token()
	if _gcbe != nil {
		return _gcbe
	}
	if _bfca, _bdege := _agbba.(_d.EndElement); _bdege && _bfca.Name == start.Name {
		*_dafg = 1
		return nil
	}
	if _ceeff, _fbgcb := _agbba.(_d.CharData); !_fbgcb {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _agbba)
	} else {
		switch string(_ceeff) {
		case "":
			*_dafg = 0
		case "\u0074":
			*_dafg = 1
		case "\u006d\u0069\u0064":
			*_dafg = 2
		case "\u0062":
			*_dafg = 3
		}
	}
	_agbba, _gcbe = d.Token()
	if _gcbe != nil {
		return _gcbe
	}
	if _ececc, _ccabg := _agbba.(_d.EndElement); _ccabg && _ececc.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _agbba)
}
func (_fce *CT_ChildMax) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _fcb := range start.Attr {
		if _fcb.Name.Local == "\u0076\u0061\u006c" {
			_acf, _cgd := _ae.ParseInt(_fcb.Value, 10, 32)
			if _cgd != nil {
				return _cgd
			}
			_gbfe := int32(_acf)
			_fce.ValAttr = &_gbfe
			continue
		}
	}
	for {
		_dcc, _cbfe := d.Token()
		if _cbfe != nil {
			return _c.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u004d\u0061\u0078\u003a\u0020\u0025\u0073", _cbfe)
		}
		if _cbad, _ecab := _dcc.(_d.EndElement); _ecab && _cbad.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_LayoutVariablePropertySet() *CT_LayoutVariablePropertySet {
	_ffecf := &CT_LayoutVariablePropertySet{}
	return _ffecf
}
func (_eccg ST_ChildAlignment) ValidateWithPath(path string) error {
	switch _eccg {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eccg))
	}
	return nil
}

const (
	ST_BendPointUnset ST_BendPoint = 0
	ST_BendPointBeg   ST_BendPoint = 1
	ST_BendPointDef   ST_BendPoint = 2
	ST_BendPointEnd   ST_BendPoint = 3
)

func (_dbef ST_ContinueDirection) ValidateWithPath(path string) error {
	switch _dbef {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dbef))
	}
	return nil
}
func (_gcc *CT_CTDescription) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _aecd := range start.Attr {
		if _aecd.Name.Local == "\u006c\u0061\u006e\u0067" {
			_fdg, _agg := _aecd.Value, error(nil)
			if _agg != nil {
				return _agg
			}
			_gcc.LangAttr = &_fdg
			continue
		}
		if _aecd.Name.Local == "\u0076\u0061\u006c" {
			_eca, _bfad := _aecd.Value, error(nil)
			if _bfad != nil {
				return _bfad
			}
			_gcc.ValAttr = _eca
			continue
		}
	}
	for {
		_daf, _aaaf := d.Token()
		if _aaaf != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0054\u0044\u0065s\u0063r\u0069\u0070\u0074\u0069\u006f\u006e\u003a \u0025\u0073", _aaaf)
		}
		if _ede, _efec := _daf.(_d.EndElement); _efec && _ede.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the AG_ConstraintAttributes and its children, prefixing error messages with path
func (_gcg *AG_ConstraintAttributes) ValidateWithPath(path string) error {
	if _ebg := _gcg.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _ebg != nil {
		return _ebg
	}
	if _ed := _gcg.ForAttr.ValidateWithPath(path + "\u002f\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _ed != nil {
		return _ed
	}
	if _ef := _gcg.PtTypeAttr.ValidateWithPath(path + "/\u0050\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _ef != nil {
		return _ef
	}
	return nil
}
func (_fcacg *ST_AutoTextRotation) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_edda, _gdcaf := d.Token()
	if _gdcaf != nil {
		return _gdcaf
	}
	if _ecdee, _fgac := _edda.(_d.EndElement); _fgac && _ecdee.Name == start.Name {
		*_fcacg = 1
		return nil
	}
	if _ffaae, _fgdba := _edda.(_d.CharData); !_fgdba {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _edda)
	} else {
		switch string(_ffaae) {
		case "":
			*_fcacg = 0
		case "\u006e\u006f\u006e\u0065":
			*_fcacg = 1
		case "\u0075\u0070\u0072":
			*_fcacg = 2
		case "\u0067\u0072\u0061\u0076":
			*_fcacg = 3
		}
	}
	_edda, _gdcaf = d.Token()
	if _gdcaf != nil {
		return _gdcaf
	}
	if _afdga, _ddbeb := _edda.(_d.EndElement); _ddbeb && _afdga.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _edda)
}

// Validate validates the CT_CTDescription and its children
func (_bcb *CT_CTDescription) Validate() error {
	return _bcb.ValidateWithPath("\u0043\u0054_\u0043\u0054\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e")
}
func (_dabe *CT_CxnList) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_abb:
	for {
		_ggf, _ebbg := d.Token()
		if _ebbg != nil {
			return _ebbg
		}
		switch _fdcd := _ggf.(type) {
		case _d.StartElement:
			switch _fdcd.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0078\u006e"}:
				_cfbc := NewCT_Cxn()
				if _cea := d.DecodeElement(_cfbc, &_fdcd); _cea != nil {
					return _cea
				}
				_dabe.Cxn = append(_dabe.Cxn, _cfbc)
			default:
				_af.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fC\u0078\u006e\u004c\u0069\u0073\u0074\u0020\u0025\u0076", _fdcd.Name)
				if _gedc := d.Skip(); _gedc != nil {
					return _gedc
				}
			}
		case _d.EndElement:
			break _abb
		case _d.CharData:
		}
	}
	return nil
}
func (_dgbbg ST_RotationPath) ValidateWithPath(path string) error {
	switch _dgbbg {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dgbbg))
	}
	return nil
}
func (_egbe *ST_HierarchyAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_egbe = 0
	case "\u0074\u004c":
		*_egbe = 1
	case "\u0074\u0052":
		*_egbe = 2
	case "\u0074\u0043\u0074\u0072\u0043\u0068":
		*_egbe = 3
	case "\u0074C\u0074\u0072\u0044\u0065\u0073":
		*_egbe = 4
	case "\u0062\u004c":
		*_egbe = 5
	case "\u0062\u0052":
		*_egbe = 6
	case "\u0062\u0043\u0074\u0072\u0043\u0068":
		*_egbe = 7
	case "\u0062C\u0074\u0072\u0044\u0065\u0073":
		*_egbe = 8
	case "\u006c\u0054":
		*_egbe = 9
	case "\u006c\u0042":
		*_egbe = 10
	case "\u006c\u0043\u0074\u0072\u0043\u0068":
		*_egbe = 11
	case "\u006cC\u0074\u0072\u0044\u0065\u0073":
		*_egbe = 12
	case "\u0072\u0054":
		*_egbe = 13
	case "\u0072\u0042":
		*_egbe = 14
	case "\u0072\u0043\u0074\u0072\u0043\u0068":
		*_egbe = 15
	case "\u0072C\u0074\u0072\u0044\u0065\u0073":
		*_egbe = 16
	}
	return nil
}
func (_ebgad *DataModel) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l"
	return _ebgad.CT_DataModel.MarshalXML(e, start)
}

const (
	ST_GrowDirectionUnset ST_GrowDirection = 0
	ST_GrowDirectionTL    ST_GrowDirection = 1
	ST_GrowDirectionTR    ST_GrowDirection = 2
	ST_GrowDirectionBL    ST_GrowDirection = 3
	ST_GrowDirectionBR    ST_GrowDirection = 4
)

func ParseUnionST_LayoutShapeType(s string) (ST_LayoutShapeType, error) {
	return ST_LayoutShapeType{}, nil
}

const (
	ST_SecondaryLinearDirectionUnset ST_SecondaryLinearDirection = 0
	ST_SecondaryLinearDirectionNone  ST_SecondaryLinearDirection = 1
	ST_SecondaryLinearDirectionFromL ST_SecondaryLinearDirection = 2
	ST_SecondaryLinearDirectionFromR ST_SecondaryLinearDirection = 3
	ST_SecondaryLinearDirectionFromT ST_SecondaryLinearDirection = 4
	ST_SecondaryLinearDirectionFromB ST_SecondaryLinearDirection = 5
)

type ST_ChildOrderType byte

func (_ffda *CT_ElemPropSet) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _aaac := range start.Attr {
		if _aaac.Name.Local == "\u0063\u0075\u0073t\u0046\u006c\u0069\u0070\u0056\u0065\u0072\u0074" {
			_edgd, _aggf := _ae.ParseBool(_aaac.Value)
			if _aggf != nil {
				return _aggf
			}
			_ffda.CustFlipVertAttr = &_edgd
			continue
		}
		if _aaac.Name.Local == "p\u0072\u0065\u0073\u0041\u0073\u0073\u006f\u0063\u0049\u0044" {
			_dec, _eafe := ParseUnionST_ModelId(_aaac.Value)
			if _eafe != nil {
				return _eafe
			}
			_ffda.PresAssocIDAttr = &_dec
			continue
		}
		if _aaac.Name.Local == "c\u0075\u0073\u0074\u0046\u006c\u0069\u0070\u0048\u006f\u0072" {
			_gdccc, _cfbb := _ae.ParseBool(_aaac.Value)
			if _cfbb != nil {
				return _cfbb
			}
			_ffda.CustFlipHorAttr = &_gdccc
			continue
		}
		if _aaac.Name.Local == "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u004c\u0062\u006c" {
			_fecab, _abbd := _aaac.Value, error(nil)
			if _abbd != nil {
				return _abbd
			}
			_ffda.PresStyleLblAttr = &_fecab
			continue
		}
		if _aaac.Name.Local == "\u0063u\u0073\u0074\u0053\u007a\u0058" {
			_gdgab, _egdg := _ae.ParseInt(_aaac.Value, 10, 32)
			if _egdg != nil {
				return _egdg
			}
			_cefb := int32(_gdgab)
			_ffda.CustSzXAttr = &_cefb
			continue
		}
		if _aaac.Name.Local == "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0043\u006e\u0074" {
			_bcfa, _eacf := _ae.ParseInt(_aaac.Value, 10, 32)
			if _eacf != nil {
				return _eacf
			}
			_dgcd := int32(_bcfa)
			_ffda.PresStyleCntAttr = &_dgcd
			continue
		}
		if _aaac.Name.Local == "\u0063u\u0073\u0074\u0053\u007a\u0059" {
			_fcbd, _accg := _ae.ParseInt(_aaac.Value, 10, 32)
			if _accg != nil {
				return _accg
			}
			_bfgf := int32(_fcbd)
			_ffda.CustSzYAttr = &_bfgf
			continue
		}
		if _aaac.Name.Local == "\u006co\u0043\u0061\u0074\u0049\u0064" {
			_dgfg, _ecebe := _aaac.Value, error(nil)
			if _ecebe != nil {
				return _ecebe
			}
			_ffda.LoCatIdAttr = &_dgfg
			continue
		}
		if _aaac.Name.Local == "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0058" {
			_cbcc, _fagc := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _fagc != nil {
				return _fagc
			}
			_ffda.CustScaleXAttr = &_cbcc
			continue
		}
		if _aaac.Name.Local == "\u0071s\u0043\u0061\u0074\u0049\u0064" {
			_fafb, _gfee := _aaac.Value, error(nil)
			if _gfee != nil {
				return _gfee
			}
			_ffda.QsCatIdAttr = &_fafb
			continue
		}
		if _aaac.Name.Local == "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0059" {
			_dgae, _cdbf := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _cdbf != nil {
				return _cdbf
			}
			_ffda.CustScaleYAttr = &_dgae
			continue
		}
		if _aaac.Name.Local == "\u0063u\u0073\u0074\u0041\u006e\u0067" {
			_cagc, _dbgbf := _ae.ParseInt(_aaac.Value, 10, 32)
			if _dbgbf != nil {
				return _dbgbf
			}
			_eeedb := int32(_cagc)
			_ffda.CustAngAttr = &_eeedb
			continue
		}
		if _aaac.Name.Local == "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0052\u0061\u0064" {
			_cage, _eecf := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _eecf != nil {
				return _eecf
			}
			_ffda.CustRadScaleRadAttr = &_cage
			continue
		}
		if _aaac.Name.Local == "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0058" {
			_cfec, _bbaa := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _bbaa != nil {
				return _bbaa
			}
			_ffda.CustLinFactXAttr = &_cfec
			continue
		}
		if _aaac.Name.Local == "\u0071\u0073\u0054\u0079\u0070\u0065\u0049\u0064" {
			_aefg, _bcbb := _aaac.Value, error(nil)
			if _bcbb != nil {
				return _bcbb
			}
			_ffda.QsTypeIdAttr = &_aefg
			continue
		}
		if _aaac.Name.Local == "\u0063\u006f\u0068\u0065\u0072\u0065\u006e\u0074\u0033\u0044\u004f\u0066\u0066" {
			_bbgdg, _eagf := _ae.ParseBool(_aaac.Value)
			if _eagf != nil {
				return _eagf
			}
			_ffda.Coherent3DOffAttr = &_bbgdg
			continue
		}
		if _aaac.Name.Local == "\u0063\u0075\u0073t\u0054" {
			_dgcf, _bfba := _ae.ParseBool(_aaac.Value)
			if _bfba != nil {
				return _bfba
			}
			_ffda.CustTAttr = &_dgcf
			continue
		}
		if _aaac.Name.Local == "\u0070\u0072\u0065\u0073\u004e\u0061\u006d\u0065" {
			_ddad, _edgg := _aaac.Value, error(nil)
			if _edgg != nil {
				return _edgg
			}
			_ffda.PresNameAttr = &_ddad
			continue
		}
		if _aaac.Name.Local == "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0059" {
			_eeg, _gefb := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _gefb != nil {
				return _gefb
			}
			_ffda.CustLinFactNeighborYAttr = &_eeg
			continue
		}
		if _aaac.Name.Local == "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0059" {
			_caaab, _efd := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _efd != nil {
				return _efd
			}
			_ffda.CustLinFactYAttr = &_caaab
			continue
		}
		if _aaac.Name.Local == "\u006c\u006f\u0054\u0079\u0070\u0065\u0049\u0064" {
			_dbfa, _geee := _aaac.Value, error(nil)
			if _geee != nil {
				return _geee
			}
			_ffda.LoTypeIdAttr = &_dbfa
			continue
		}
		if _aaac.Name.Local == "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0049\u006e\u0063" {
			_bffcc, _abbc := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _abbc != nil {
				return _abbc
			}
			_ffda.CustRadScaleIncAttr = &_bffcc
			continue
		}
		if _aaac.Name.Local == "\u0070\u0068\u006cd\u0072" {
			_abfda, _fbd := _ae.ParseBool(_aaac.Value)
			if _fbd != nil {
				return _fbd
			}
			_ffda.PhldrAttr = &_abfda
			continue
		}
		if _aaac.Name.Local == "\u0063\u0073\u0054\u0079\u0070\u0065\u0049\u0064" {
			_gfed, _bdga := _aaac.Value, error(nil)
			if _bdga != nil {
				return _bdga
			}
			_ffda.CsTypeIdAttr = &_gfed
			continue
		}
		if _aaac.Name.Local == "\u0063s\u0043\u0061\u0074\u0049\u0064" {
			_fggd, _gfaf := _aaac.Value, error(nil)
			if _gfaf != nil {
				return _gfaf
			}
			_ffda.CsCatIdAttr = &_fggd
			continue
		}
		if _aaac.Name.Local == "\u0070\u0068\u006c\u0064\u0072\u0054" {
			_fecg, _gbfg := _aaac.Value, error(nil)
			if _gbfg != nil {
				return _gbfg
			}
			_ffda.PhldrTAttr = &_fecg
			continue
		}
		if _aaac.Name.Local == "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0049\u0064\u0078" {
			_decb, _cdgc := _ae.ParseInt(_aaac.Value, 10, 32)
			if _cdgc != nil {
				return _cdgc
			}
			_bfdg := int32(_decb)
			_ffda.PresStyleIdxAttr = &_bfdg
			continue
		}
		if _aaac.Name.Local == "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0058" {
			_gbfb, _fcfe := ParseUnionST_PrSetCustVal(_aaac.Value)
			if _fcfe != nil {
				return _fcfe
			}
			_ffda.CustLinFactNeighborXAttr = &_gbfb
			continue
		}
	}
_aacbeb:
	for {
		_aebc, _dcac := d.Token()
		if _dcac != nil {
			return _dcac
		}
		switch _cbea := _aebc.(type) {
		case _d.StartElement:
			switch _cbea.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004c\u0061\u0079\u006f\u0075t\u0056\u0061\u0072\u0073"}:
				_ffda.PresLayoutVars = NewCT_LayoutVariablePropertySet()
				if _cade := d.DecodeElement(_ffda.PresLayoutVars, &_cbea); _cade != nil {
					return _cade
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079l\u0065"}:
				_ffda.Style = _ca.NewCT_ShapeStyle()
				if _dfgg := d.DecodeElement(_ffda.Style, &_cbea); _dfgg != nil {
					return _dfgg
				}
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0045\u006c\u0065\u006d\u0050\u0072\u006fp\u0053e\u0074\u0020\u0025\u0076", _cbea.Name)
				if _ecaf := d.Skip(); _ecaf != nil {
					return _ecaf
				}
			}
		case _d.EndElement:
			break _aacbeb
		case _d.CharData:
		}
	}
	return nil
}
func (_ccef *CT_BulletEnabled) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _gbd := range start.Attr {
		if _gbd.Name.Local == "\u0076\u0061\u006c" {
			_ebc, _fcc := _ae.ParseBool(_gbd.Value)
			if _fcc != nil {
				return _fcc
			}
			_ccef.ValAttr = &_ebc
			continue
		}
	}
	for {
		_ebdb, _ecda := d.Token()
		if _ecda != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0042\u0075\u006c\u006ce\u0074E\u006e\u0061\u0062\u006c\u0065\u0064\u003a \u0025\u0073", _ecda)
		}
		if _bgg, _aecf := _ebdb.(_d.EndElement); _aecf && _bgg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_gbecd *ST_NodeHorizontalAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gbecd = 0
	case "\u006c":
		*_gbecd = 1
	case "\u0063\u0074\u0072":
		*_gbecd = 2
	case "\u0072":
		*_gbecd = 3
	}
	return nil
}

type CT_Description struct {
	LangAttr *string
	ValAttr  string
}

func NewCT_SDDescription() *CT_SDDescription { _fgbaf := &CT_SDDescription{}; return _fgbaf }

// Validate validates the LayoutDefHdr and its children
func (_cefg *LayoutDefHdr) Validate() error {
	return _cefg.ValidateWithPath("\u004c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072")
}

// Validate validates the CT_Rules and its children
func (_dfcae *CT_Rules) Validate() error {
	return _dfcae.ValidateWithPath("\u0043\u0054\u005f\u0052\u0075\u006c\u0065\u0073")
}
func (_agbd ST_StartingElement) Validate() error { return _agbd.ValidateWithPath("") }

type ST_ChildAlignment byte

// Validate validates the CT_LayoutNode and its children
func (_aded *CT_LayoutNode) Validate() error {
	return _aded.ValidateWithPath("\u0043\u0054\u005f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065")
}
func (_efcb *CT_HierBranchStyle) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _agbg := range start.Attr {
		if _agbg.Name.Local == "\u0076\u0061\u006c" {
			_efcb.ValAttr.UnmarshalXMLAttr(_agbg)
			continue
		}
	}
	for {
		_ceeda, _fddgg := d.Token()
		if _fddgg != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054_\u0048\u0069\u0065\u0072\u0042\u0072\u0061n\u0063\u0068\u0053\u0074\u0079\u006c\u0065\u003a\u0020\u0025\u0073", _fddgg)
		}
		if _aeba, _fegg := _ceeda.(_d.EndElement); _fegg && _aeba.Name == start.Name {
			break
		}
	}
	return nil
}

type ST_ConstraintRelationship byte

// ValidateWithPath validates the LayoutDefHdr and its children, prefixing error messages with path
func (_eadf *LayoutDefHdr) ValidateWithPath(path string) error {
	if _dbcga := _eadf.CT_DiagramDefinitionHeader.ValidateWithPath(path); _dbcga != nil {
		return _dbcga
	}
	return nil
}
func ParseUnionST_PrSetCustVal(s string) (ST_PrSetCustVal, error) { return ST_PrSetCustVal{}, nil }

// ValidateWithPath validates the DataModel and its children, prefixing error messages with path
func (_caad *DataModel) ValidateWithPath(path string) error {
	if _adeca := _caad.CT_DataModel.ValidateWithPath(path); _adeca != nil {
		return _adeca
	}
	return nil
}
func (_aecda *ST_FlowDirection) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cbge, _adeg := d.Token()
	if _adeg != nil {
		return _adeg
	}
	if _cadc, _cddace := _cbge.(_d.EndElement); _cddace && _cadc.Name == start.Name {
		*_aecda = 1
		return nil
	}
	if _ccbf, _cfbgb := _cbge.(_d.CharData); !_cfbgb {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cbge)
	} else {
		switch string(_ccbf) {
		case "":
			*_aecda = 0
		case "\u0072\u006f\u0077":
			*_aecda = 1
		case "\u0063\u006f\u006c":
			*_aecda = 2
		}
	}
	_cbge, _adeg = d.Token()
	if _adeg != nil {
		return _adeg
	}
	if _bgcb, _gdcf := _cbge.(_d.EndElement); _gdcf && _bgcb.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cbge)
}

type CT_ForEach struct {
	NameAttr          *string
	RefAttr           *string
	Alg               []*CT_Algorithm
	Shape             []*CT_Shape
	PresOf            []*CT_PresentationOf
	ConstrLst         []*CT_Constraints
	RuleLst           []*CT_Rules
	ForEach           []*CT_ForEach
	LayoutNode        []*CT_LayoutNode
	Choose            []*CT_Choose
	ExtLst            []*_ca.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}
type AG_IteratorAttributes struct {
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

const (
	ST_LinearDirectionUnset ST_LinearDirection = 0
	ST_LinearDirectionFromL ST_LinearDirection = 1
	ST_LinearDirectionFromR ST_LinearDirection = 2
	ST_LinearDirectionFromT ST_LinearDirection = 3
	ST_LinearDirectionFromB ST_LinearDirection = 4
)

func (_gdbea *CT_ResizeHandles) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _agbc := range start.Attr {
		if _agbc.Name.Local == "\u0076\u0061\u006c" {
			_gdbea.ValAttr.UnmarshalXMLAttr(_agbc)
			continue
		}
	}
	for {
		_fagfa, _abgae := d.Token()
		if _abgae != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0052\u0065\u0073\u0069z\u0065H\u0061\u006e\u0064\u006c\u0065\u0073\u003a \u0025\u0073", _abgae)
		}
		if _bdce, _aagc := _fagfa.(_d.EndElement); _aagc && _bdce.Name == start.Name {
			break
		}
	}
	return nil
}
func (_fdfaf ST_TextAnchorVertical) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fdfaf.String(), start)
}
func (_ddggb ST_ChildAlignment) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_aebd := _d.Attr{}
	_aebd.Name = name
	switch _ddggb {
	case ST_ChildAlignmentUnset:
		_aebd.Value = ""
	case ST_ChildAlignmentT:
		_aebd.Value = "\u0074"
	case ST_ChildAlignmentB:
		_aebd.Value = "\u0062"
	case ST_ChildAlignmentL:
		_aebd.Value = "\u006c"
	case ST_ChildAlignmentR:
		_aebd.Value = "\u0072"
	}
	return _aebd, nil
}
func (_bgde *CT_PresentationOf) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _bgde.AxisAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_bgde.AxisAttr)})
	}
	if _bgde.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_bgde.PtTypeAttr)})
	}
	if _bgde.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_bgde.HideLastTransAttr)})
	}
	if _bgde.StAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_bgde.StAttr)})
	}
	if _bgde.CntAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u006e\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_bgde.CntAttr)})
	}
	if _bgde.StepAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _c.Sprintf("\u0025\u0076", *_bgde.StepAttr)})
	}
	e.EncodeToken(start)
	if _bgde.ExtLst != nil {
		_egcd := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_bgde.ExtLst, _egcd)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_fgdg *CT_ElemPropSet) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _fgdg.PresAssocIDAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "p\u0072\u0065\u0073\u0041\u0073\u0073\u006f\u0063\u0049\u0044"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.PresAssocIDAttr)})
	}
	if _fgdg.PresNameAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0065\u0073\u004e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.PresNameAttr)})
	}
	if _fgdg.PresStyleLblAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.PresStyleLblAttr)})
	}
	if _fgdg.PresStyleIdxAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0049\u0064\u0078"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.PresStyleIdxAttr)})
	}
	if _fgdg.PresStyleCntAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0043\u006e\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.PresStyleCntAttr)})
	}
	if _fgdg.LoTypeIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006c\u006f\u0054\u0079\u0070\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.LoTypeIdAttr)})
	}
	if _fgdg.LoCatIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006co\u0043\u0061\u0074\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.LoCatIdAttr)})
	}
	if _fgdg.QsTypeIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0071\u0073\u0054\u0079\u0070\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.QsTypeIdAttr)})
	}
	if _fgdg.QsCatIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0071s\u0043\u0061\u0074\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.QsCatIdAttr)})
	}
	if _fgdg.CsTypeIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0073\u0054\u0079\u0070\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CsTypeIdAttr)})
	}
	if _fgdg.CsCatIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063s\u0043\u0061\u0074\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CsCatIdAttr)})
	}
	if _fgdg.Coherent3DOffAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u006f\u0068\u0065\u0072\u0065\u006e\u0074\u0033\u0044\u004f\u0066\u0066"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_fgdg.Coherent3DOffAttr))})
	}
	if _fgdg.PhldrTAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0068\u006c\u0064\u0072\u0054"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.PhldrTAttr)})
	}
	if _fgdg.PhldrAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0068\u006cd\u0072"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_fgdg.PhldrAttr))})
	}
	if _fgdg.CustAngAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063u\u0073\u0074\u0041\u006e\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustAngAttr)})
	}
	if _fgdg.CustFlipVertAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0075\u0073t\u0046\u006c\u0069\u0070\u0056\u0065\u0072\u0074"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_fgdg.CustFlipVertAttr))})
	}
	if _fgdg.CustFlipHorAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "c\u0075\u0073\u0074\u0046\u006c\u0069\u0070\u0048\u006f\u0072"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_fgdg.CustFlipHorAttr))})
	}
	if _fgdg.CustSzXAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063u\u0073\u0074\u0053\u007a\u0058"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustSzXAttr)})
	}
	if _fgdg.CustSzYAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063u\u0073\u0074\u0053\u007a\u0059"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustSzYAttr)})
	}
	if _fgdg.CustScaleXAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0058"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustScaleXAttr)})
	}
	if _fgdg.CustScaleYAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0059"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustScaleYAttr)})
	}
	if _fgdg.CustTAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0075\u0073t\u0054"}, Value: _c.Sprintf("\u0025\u0064", _efffd(*_fgdg.CustTAttr))})
	}
	if _fgdg.CustLinFactXAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0058"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustLinFactXAttr)})
	}
	if _fgdg.CustLinFactYAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0059"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustLinFactYAttr)})
	}
	if _fgdg.CustLinFactNeighborXAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0058"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustLinFactNeighborXAttr)})
	}
	if _fgdg.CustLinFactNeighborYAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0059"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustLinFactNeighborYAttr)})
	}
	if _fgdg.CustRadScaleRadAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0052\u0061\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustRadScaleRadAttr)})
	}
	if _fgdg.CustRadScaleIncAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0049\u006e\u0063"}, Value: _c.Sprintf("\u0025\u0076", *_fgdg.CustRadScaleIncAttr)})
	}
	e.EncodeToken(start)
	if _fgdg.PresLayoutVars != nil {
		_gfac := _d.StartElement{Name: _d.Name{Local: "\u0070\u0072\u0065\u0073\u004c\u0061\u0079\u006f\u0075t\u0056\u0061\u0072\u0073"}}
		e.EncodeElement(_fgdg.PresLayoutVars, _gfac)
	}
	if _fgdg.Style != nil {
		_afdd := _d.StartElement{Name: _d.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_fgdg.Style, _afdd)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the ColorsDefHdr and its children
func (_ggdb *ColorsDefHdr) Validate() error {
	return _ggdb.ValidateWithPath("\u0043\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072")
}

type ST_FallbackDimension byte

func (_agafg ST_DiagramTextAlignment) String() string {
	switch _agafg {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0063\u0074\u0072"
	case 3:
		return "\u0072"
	}
	return ""
}

// Validate validates the CT_ChildPref and its children
func (_cbac *CT_ChildPref) Validate() error {
	return _cbac.ValidateWithPath("\u0043\u0054\u005fC\u0068\u0069\u006c\u0064\u0050\u0072\u0065\u0066")
}

const (
	ST_DirectionUnset ST_Direction = 0
	ST_DirectionNorm  ST_Direction = 1
	ST_DirectionRev   ST_Direction = 2
)

func (_aaacb ST_VariableType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_aaacb.String(), start)
}

const (
	ST_OffsetUnset ST_Offset = 0
	ST_OffsetCtr   ST_Offset = 1
	ST_OffsetOff   ST_Offset = 2
)

// Validate validates the CT_DiagramDefinition and its children
func (_dfbcc *CT_DiagramDefinition) Validate() error {
	return _dfbcc.ValidateWithPath("C\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069ti\u006f\u006e")
}
func (_gadba ST_GrowDirection) ValidateWithPath(path string) error {
	switch _gadba {
	case 0, 1, 2, 3, 4:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gadba))
	}
	return nil
}
func (_gdba *ColorsDefHdrLst) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gdba.CT_ColorTransformHeaderLst = *NewCT_ColorTransformHeaderLst()
_bfdb:
	for {
		_gbef, _deac := d.Token()
		if _deac != nil {
			return _deac
		}
		switch _dabd := _gbef.(type) {
		case _d.StartElement:
			switch _dabd.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_caea := NewCT_ColorTransformHeader()
				if _ecdg := d.DecodeElement(_caea, &_dabd); _ecdg != nil {
					return _ecdg
				}
				_gdba.ColorsDefHdr = append(_gdba.ColorsDefHdr, _caea)
			default:
				_af.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u006f\u006c\u006f\u0072\u0073D\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _dabd.Name)
				if _bbce := d.Skip(); _bbce != nil {
					return _bbce
				}
			}
		case _d.EndElement:
			break _bfdb
		case _d.CharData:
		}
	}
	return nil
}
func (_fedb *ST_OutputShapeType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fedb = 0
	case "\u006e\u006f\u006e\u0065":
		*_fedb = 1
	case "\u0063\u006f\u006e\u006e":
		*_fedb = 2
	}
	return nil
}
func (_gcae *LayoutDefHdr) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"
	return _gcae.CT_DiagramDefinitionHeader.MarshalXML(e, start)
}
func (_eeefb ST_AnimLvlStr) String() string {
	switch _eeefb {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u006c\u0076\u006c"
	case 3:
		return "\u0063\u0074\u0072"
	}
	return ""
}
func (_cgabe *ST_LayoutShapeType) ValidateWithPath(path string) error {
	_gdgg := []string{}
	if _cgabe.ST_ShapeType != _ca.ST_ShapeTypeUnset {
		_gdgg = append(_gdgg, "\u0053\u0054\u005fS\u0068\u0061\u0070\u0065\u0054\u0079\u0070\u0065")
	}
	if _cgabe.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		_gdgg = append(_gdgg, "\u0053T\u005fO\u0075\u0074\u0070\u0075\u0074S\u0068\u0061p\u0065\u0054\u0079\u0070\u0065")
	}
	if len(_gdgg) > 1 {
		return _c.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _gdgg)
	}
	return nil
}

const (
	ST_HierBranchStyleUnset ST_HierBranchStyle = 0
	ST_HierBranchStyleL     ST_HierBranchStyle = 1
	ST_HierBranchStyleR     ST_HierBranchStyle = 2
	ST_HierBranchStyleHang  ST_HierBranchStyle = 3
	ST_HierBranchStyleStd   ST_HierBranchStyle = 4
	ST_HierBranchStyleInit  ST_HierBranchStyle = 5
)

type ST_ContinueDirection byte

func (_agcd ST_SecondaryLinearDirection) String() string {
	switch _agcd {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0066\u0072\u006fm\u004c"
	case 3:
		return "\u0066\u0072\u006fm\u0052"
	case 4:
		return "\u0066\u0072\u006fm\u0054"
	case 5:
		return "\u0066\u0072\u006fm\u0042"
	}
	return ""
}

// ValidateWithPath validates the CT_DiagramDefinitionHeaderLst and its children, prefixing error messages with path
func (_abg *CT_DiagramDefinitionHeaderLst) ValidateWithPath(path string) error {
	for _eecg, _gade := range _abg.LayoutDefHdr {
		if _ggca := _gade.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u004cay\u006f\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u005b\u0025\u0064\u005d", path, _eecg)); _ggca != nil {
			return _ggca
		}
	}
	return nil
}

const (
	ST_DiagramTextAlignmentUnset ST_DiagramTextAlignment = 0
	ST_DiagramTextAlignmentL     ST_DiagramTextAlignment = 1
	ST_DiagramTextAlignmentCtr   ST_DiagramTextAlignment = 2
	ST_DiagramTextAlignmentR     ST_DiagramTextAlignment = 3
)

func (_fcag ST_PyramidAccentPosition) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_fcag.String(), start)
}
func (_agbbd ST_AnimOneStr) Validate() error { return _agbbd.ValidateWithPath("") }
func NewCT_Name() *CT_Name                   { _ccfc := &CT_Name{}; return _ccfc }

// ValidateWithPath validates the CT_ChildPref and its children, prefixing error messages with path
func (_efff *CT_ChildPref) ValidateWithPath(path string) error {
	if _efff.ValAttr != nil {
		if *_efff.ValAttr < -1 {
			return _c.Errorf("\u0025\u0073/m\u002e\u0056\u0061l\u0041\u0074\u0074\u0072 mu\u0073t \u0062\u0065\u0020\u003e\u003d\u0020\u002d1 \u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, *_efff.ValAttr)
		}
	}
	return nil
}

type CT_Pt struct {
	ModelIdAttr ST_ModelId
	TypeAttr    ST_PtType
	CxnIdAttr   *ST_ModelId
	PrSet       *CT_ElemPropSet
	SpPr        *_ca.CT_ShapeProperties
	T           *_ca.CT_TextBody
	ExtLst      *_ca.CT_OfficeArtExtensionList
}

func (_ccdgc *ST_AutoTextRotation) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ccdgc = 0
	case "\u006e\u006f\u006e\u0065":
		*_ccdgc = 1
	case "\u0075\u0070\u0072":
		*_ccdgc = 2
	case "\u0067\u0072\u0061\u0076":
		*_ccdgc = 3
	}
	return nil
}
func (_cfbg ST_ParameterVal) String() string {
	if _cfbg.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		return _cfbg.ST_DiagramHorizontalAlignment.String()
	}
	if _cfbg.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		return _cfbg.ST_VerticalAlignment.String()
	}
	if _cfbg.ST_ChildDirection != ST_ChildDirectionUnset {
		return _cfbg.ST_ChildDirection.String()
	}
	if _cfbg.ST_ChildAlignment != ST_ChildAlignmentUnset {
		return _cfbg.ST_ChildAlignment.String()
	}
	if _cfbg.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		return _cfbg.ST_SecondaryChildAlignment.String()
	}
	if _cfbg.ST_LinearDirection != ST_LinearDirectionUnset {
		return _cfbg.ST_LinearDirection.String()
	}
	if _cfbg.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		return _cfbg.ST_SecondaryLinearDirection.String()
	}
	if _cfbg.ST_StartingElement != ST_StartingElementUnset {
		return _cfbg.ST_StartingElement.String()
	}
	if _cfbg.ST_BendPoint != ST_BendPointUnset {
		return _cfbg.ST_BendPoint.String()
	}
	if _cfbg.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		return _cfbg.ST_ConnectorRouting.String()
	}
	if _cfbg.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		return _cfbg.ST_ArrowheadStyle.String()
	}
	if _cfbg.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		return _cfbg.ST_ConnectorDimension.String()
	}
	if _cfbg.ST_RotationPath != ST_RotationPathUnset {
		return _cfbg.ST_RotationPath.String()
	}
	if _cfbg.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		return _cfbg.ST_CenterShapeMapping.String()
	}
	if _cfbg.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		return _cfbg.ST_NodeHorizontalAlignment.String()
	}
	if _cfbg.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		return _cfbg.ST_NodeVerticalAlignment.String()
	}
	if _cfbg.ST_FallbackDimension != ST_FallbackDimensionUnset {
		return _cfbg.ST_FallbackDimension.String()
	}
	if _cfbg.ST_TextDirection != ST_TextDirectionUnset {
		return _cfbg.ST_TextDirection.String()
	}
	if _cfbg.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		return _cfbg.ST_PyramidAccentPosition.String()
	}
	if _cfbg.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		return _cfbg.ST_PyramidAccentTextMargin.String()
	}
	if _cfbg.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		return _cfbg.ST_TextBlockDirection.String()
	}
	if _cfbg.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		return _cfbg.ST_TextAnchorHorizontal.String()
	}
	if _cfbg.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		return _cfbg.ST_TextAnchorVertical.String()
	}
	if _cfbg.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		return _cfbg.ST_DiagramTextAlignment.String()
	}
	if _cfbg.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		return _cfbg.ST_AutoTextRotation.String()
	}
	if _cfbg.ST_GrowDirection != ST_GrowDirectionUnset {
		return _cfbg.ST_GrowDirection.String()
	}
	if _cfbg.ST_FlowDirection != ST_FlowDirectionUnset {
		return _cfbg.ST_FlowDirection.String()
	}
	if _cfbg.ST_ContinueDirection != ST_ContinueDirectionUnset {
		return _cfbg.ST_ContinueDirection.String()
	}
	if _cfbg.ST_Breakpoint != ST_BreakpointUnset {
		return _cfbg.ST_Breakpoint.String()
	}
	if _cfbg.ST_Offset != ST_OffsetUnset {
		return _cfbg.ST_Offset.String()
	}
	if _cfbg.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		return _cfbg.ST_HierarchyAlignment.String()
	}
	if _cfbg.Int32 != nil {
		return _c.Sprintf("\u0025\u0076", *_cfbg.Int32)
	}
	if _cfbg.Float64 != nil {
		return _c.Sprintf("\u0025\u0076", *_cfbg.Float64)
	}
	if _cfbg.Bool != nil {
		return _c.Sprintf("\u0025\u0076", *_cfbg.Bool)
	}
	if _cfbg.StringVal != nil {
		return _c.Sprintf("\u0025\u0076", *_cfbg.StringVal)
	}
	if _cfbg.ST_ConnectorPoint != ST_ConnectorPointUnset {
		return _cfbg.ST_ConnectorPoint.String()
	}
	return ""
}
func (_bgdfc *CT_StyleDefinitionHeader) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _c.Sprintf("\u0025\u0076", _bgdfc.UniqueIdAttr)})
	if _bgdfc.MinVerAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _c.Sprintf("\u0025\u0076", *_bgdfc.MinVerAttr)})
	}
	if _bgdfc.ResIdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0072\u0065\u0073I\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_bgdfc.ResIdAttr)})
	}
	e.EncodeToken(start)
	_gbgb := _d.StartElement{Name: _d.Name{Local: "\u0074\u0069\u0074l\u0065"}}
	for _, _dbgd := range _bgdfc.Title {
		e.EncodeElement(_dbgd, _gbgb)
	}
	_ebacc := _d.StartElement{Name: _d.Name{Local: "\u0064\u0065\u0073\u0063"}}
	for _, _eaab := range _bgdfc.Desc {
		e.EncodeElement(_eaab, _ebacc)
	}
	if _bgdfc.CatLst != nil {
		_egbda := _d.StartElement{Name: _d.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_bgdfc.CatLst, _egbda)
	}
	if _bgdfc.ExtLst != nil {
		_cfaa := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_bgdfc.ExtLst, _cfaa)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_fadf ST_ParameterId) String() string {
	switch _fadf {
	case 0:
		return ""
	case 1:
		return "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n"
	case 2:
		return "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n"
	case 3:
		return "\u0063\u0068\u0044i\u0072"
	case 4:
		return "\u0063h\u0041\u006c\u0069\u0067\u006e"
	case 5:
		return "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e"
	case 6:
		return "\u006c\u0069\u006e\u0044\u0069\u0072"
	case 7:
		return "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r"
	case 8:
		return "\u0073\u0074\u0045\u006c\u0065\u006d"
	case 9:
		return "\u0062\u0065\u006e\u0064\u0050\u0074"
	case 10:
		return "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074"
	case 11:
		return "\u0062\u0065\u0067\u0053\u0074\u0079"
	case 12:
		return "\u0065\u006e\u0064\u0053\u0074\u0079"
	case 13:
		return "\u0064\u0069\u006d"
	case 14:
		return "\u0072o\u0074\u0050\u0061\u0074\u0068"
	case 15:
		return "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p"
	case 16:
		return "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e"
	case 17:
		return "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e"
	case 18:
		return "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b"
	case 19:
		return "\u0074\u0078\u0044i\u0072"
	case 20:
		return "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073"
	case 21:
		return "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072"
	case 22:
		return "\u0074x\u0042\u006c\u0044\u0069\u0072"
	case 23:
		return "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a"
	case 24:
		return "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074"
	case 25:
		return "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068"
	case 26:
		return "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068"
	case 27:
		return "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e"
	case 28:
		return "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e"
	case 29:
		return "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case 30:
		return "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case 31:
		return "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft"
	case 32:
		return "\u0067\u0072\u0044i\u0072"
	case 33:
		return "\u0066l\u006f\u0077\u0044\u0069\u0072"
	case 34:
		return "\u0063o\u006e\u0074\u0044\u0069\u0072"
	case 35:
		return "\u0062\u006b\u0070\u0074"
	case 36:
		return "\u006f\u0066\u0066"
	case 37:
		return "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n"
	case 38:
		return "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c"
	case 39:
		return "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c"
	case 40:
		return "\u0073\u0074\u0041n\u0067"
	case 41:
		return "\u0073p\u0061\u006e\u0041\u006e\u0067"
	case 42:
		return "\u0061\u0072"
	case 43:
		return "\u006cn\u0053\u0070\u0050\u0061\u0072"
	case 44:
		return "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050"
	case 45:
		return "\u006c\u006e\u0053\u0070\u0043\u0068"
	case 46:
		return "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P"
	case 47:
		return "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074"
	case 48:
		return "\u0061l\u0069\u0067\u006e\u0054\u0078"
	case 49:
		return "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065"
	case 50:
		return "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065"
	case 51:
		return "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065"
	case 52:
		return "\u0073r\u0063\u004e\u006f\u0064\u0065"
	case 53:
		return "\u0064s\u0074\u004e\u006f\u0064\u0065"
	case 54:
		return "\u0062\u0065\u0067\u0050\u0074\u0073"
	case 55:
		return "\u0065\u006e\u0064\u0050\u0074\u0073"
	}
	return ""
}
func (_bcda *ST_DiagramTextAlignment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_badd, _acgc := d.Token()
	if _acgc != nil {
		return _acgc
	}
	if _dcge, _faffg := _badd.(_d.EndElement); _faffg && _dcge.Name == start.Name {
		*_bcda = 1
		return nil
	}
	if _abea, _adcbg := _badd.(_d.CharData); !_adcbg {
		return _c.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _badd)
	} else {
		switch string(_abea) {
		case "":
			*_bcda = 0
		case "\u006c":
			*_bcda = 1
		case "\u0063\u0074\u0072":
			*_bcda = 2
		case "\u0072":
			*_bcda = 3
		}
	}
	_badd, _acgc = d.Token()
	if _acgc != nil {
		return _acgc
	}
	if _cdcg, _bgca := _badd.(_d.EndElement); _bgca && _cdcg.Name == start.Name {
		return nil
	}
	return _c.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _badd)
}

type ST_Ints []int32

func NewCT_CTCategories() *CT_CTCategories       { _fafa := &CT_CTCategories{}; return _fafa }
func (_gcgg ST_OutputShapeType) Validate() error { return _gcgg.ValidateWithPath("") }
func (_adeb *CT_Name) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _adeb.LangAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_adeb.LangAttr)})
	}
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _adeb.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_afdb *CT_CTStyleLabel) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _c.Sprintf("\u0025\u0076", _afdb.NameAttr)})
	e.EncodeToken(start)
	if _afdb.FillClrLst != nil {
		_dded := _d.StartElement{Name: _d.Name{Local: "\u0066\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_afdb.FillClrLst, _dded)
	}
	if _afdb.LinClrLst != nil {
		_cgc := _d.StartElement{Name: _d.Name{Local: "\u006ci\u006e\u0043\u006c\u0072\u004c\u0073t"}}
		e.EncodeElement(_afdb.LinClrLst, _cgc)
	}
	if _afdb.EffectClrLst != nil {
		_dgf := _d.StartElement{Name: _d.Name{Local: "\u0065\u0066\u0066e\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_afdb.EffectClrLst, _dgf)
	}
	if _afdb.TxLinClrLst != nil {
		_gdg := _d.StartElement{Name: _d.Name{Local: "t\u0078\u004c\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_afdb.TxLinClrLst, _gdg)
	}
	if _afdb.TxFillClrLst != nil {
		_fddg := _d.StartElement{Name: _d.Name{Local: "\u0074\u0078\u0046i\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_afdb.TxFillClrLst, _fddg)
	}
	if _afdb.TxEffectClrLst != nil {
		_dea := _d.StartElement{Name: _d.Name{Local: "\u0074\u0078\u0045\u0066\u0066\u0065\u0063\u0074\u0043l\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_afdb.TxEffectClrLst, _dea)
	}
	if _afdb.ExtLst != nil {
		_edf := _d.StartElement{Name: _d.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_afdb.ExtLst, _edf)
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewCT_ForEach() *CT_ForEach { _dacf := &CT_ForEach{}; return _dacf }

// ValidateWithPath validates the CT_DiagramDefinitionHeader and its children, prefixing error messages with path
func (_cgdg *CT_DiagramDefinitionHeader) ValidateWithPath(path string) error {
	for _bgdb, _fgcb := range _cgdg.Title {
		if _bccf := _fgcb.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _bgdb)); _bccf != nil {
			return _bccf
		}
	}
	for _eecc, _debbe := range _cgdg.Desc {
		if _efg := _debbe.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _eecc)); _efg != nil {
			return _efg
		}
	}
	if _cgdg.CatLst != nil {
		if _afc := _cgdg.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _afc != nil {
			return _afc
		}
	}
	if _cgdg.ExtLst != nil {
		if _eace := _cgdg.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _eace != nil {
			return _eace
		}
	}
	return nil
}
func (_bagbf *ColorsDefHdr) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_bagbf.CT_ColorTransformHeader = *NewCT_ColorTransformHeader()
	for _, _ebfa := range start.Attr {
		if _ebfa.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_agdd, _eecac := _ebfa.Value, error(nil)
			if _eecac != nil {
				return _eecac
			}
			_bagbf.UniqueIdAttr = _agdd
			continue
		}
		if _ebfa.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_aecce, _afgeab := _ebfa.Value, error(nil)
			if _afgeab != nil {
				return _afgeab
			}
			_bagbf.MinVerAttr = &_aecce
			continue
		}
		if _ebfa.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_ggee, _bgeff := _ae.ParseInt(_ebfa.Value, 10, 32)
			if _bgeff != nil {
				return _bgeff
			}
			_effdc := int32(_ggee)
			_bagbf.ResIdAttr = &_effdc
			continue
		}
	}
_gcgba:
	for {
		_deda, _gaaag := d.Token()
		if _gaaag != nil {
			return _gaaag
		}
		switch _eacb := _deda.(type) {
		case _d.StartElement:
			switch _eacb.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_gbda := NewCT_CTName()
				if _abcbg := d.DecodeElement(_gbda, &_eacb); _abcbg != nil {
					return _abcbg
				}
				_bagbf.Title = append(_bagbf.Title, _gbda)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_adffc := NewCT_CTDescription()
				if _gcdfe := d.DecodeElement(_adffc, &_eacb); _gcdfe != nil {
					return _gcdfe
				}
				_bagbf.Desc = append(_bagbf.Desc, _adffc)
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_bagbf.CatLst = NewCT_CTCategories()
				if _egbdg := d.DecodeElement(_bagbf.CatLst, &_eacb); _egbdg != nil {
					return _egbdg
				}
			case _d.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bagbf.ExtLst = _ca.NewCT_OfficeArtExtensionList()
				if _ebdfg := d.DecodeElement(_bagbf.ExtLst, &_eacb); _ebdfg != nil {
					return _ebdfg
				}
			default:
				_af.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u006flo\u0072\u0073D\u0065\u0066\u0048\u0064\u0072\u0020\u0025\u0076", _eacb.Name)
				if _egeab := d.Skip(); _egeab != nil {
					return _egeab
				}
			}
		case _d.EndElement:
			break _gcgba
		case _d.CharData:
		}
	}
	return nil
}

type ST_HueDir byte

func (_acbea ST_PyramidAccentPosition) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ceada := _d.Attr{}
	_ceada.Name = name
	switch _acbea {
	case ST_PyramidAccentPositionUnset:
		_ceada.Value = ""
	case ST_PyramidAccentPositionBef:
		_ceada.Value = "\u0062\u0065\u0066"
	case ST_PyramidAccentPositionAft:
		_ceada.Value = "\u0061\u0066\u0074"
	}
	return _ceada, nil
}
func (_dgadf ST_FlowDirection) ValidateWithPath(path string) error {
	switch _dgadf {
	case 0, 1, 2:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dgadf))
	}
	return nil
}
func (_edad ST_BendPoint) String() string {
	switch _edad {
	case 0:
		return ""
	case 1:
		return "\u0062\u0065\u0067"
	case 2:
		return "\u0064\u0065\u0066"
	case 3:
		return "\u0065\u006e\u0064"
	}
	return ""
}
func (_dacfd ST_CenterShapeMapping) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_cfbea := _d.Attr{}
	_cfbea.Name = name
	switch _dacfd {
	case ST_CenterShapeMappingUnset:
		_cfbea.Value = ""
	case ST_CenterShapeMappingNone:
		_cfbea.Value = "\u006e\u006f\u006e\u0065"
	case ST_CenterShapeMappingFNode:
		_cfbea.Value = "\u0066\u004e\u006fd\u0065"
	}
	return _cfbea, nil
}

type ST_TextAnchorHorizontal byte

func (_feda ST_DiagramTextAlignment) Validate() error { return _feda.ValidateWithPath("") }
func (_ebbbc *ST_LinearDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_ebbbc = 0
	case "\u0066\u0072\u006fm\u004c":
		*_ebbbc = 1
	case "\u0066\u0072\u006fm\u0052":
		*_ebbbc = 2
	case "\u0066\u0072\u006fm\u0054":
		*_ebbbc = 3
	case "\u0066\u0072\u006fm\u0042":
		*_ebbbc = 4
	}
	return nil
}
func (_faca ST_FunctionOperator) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_eabg := _d.Attr{}
	_eabg.Name = name
	switch _faca {
	case ST_FunctionOperatorUnset:
		_eabg.Value = ""
	case ST_FunctionOperatorEqu:
		_eabg.Value = "\u0065\u0071\u0075"
	case ST_FunctionOperatorNeq:
		_eabg.Value = "\u006e\u0065\u0071"
	case ST_FunctionOperatorGt:
		_eabg.Value = "\u0067\u0074"
	case ST_FunctionOperatorLt:
		_eabg.Value = "\u006c\u0074"
	case ST_FunctionOperatorGte:
		_eabg.Value = "\u0067\u0074\u0065"
	case ST_FunctionOperatorLte:
		_eabg.Value = "\u006c\u0074\u0065"
	}
	return _eabg, nil
}

type CT_NumericRule struct {
	ValAttr     *float64
	FactAttr    *float64
	MaxAttr     *float64
	ExtLst      *_ca.CT_OfficeArtExtensionList
	TypeAttr    ST_ConstraintType
	ForAttr     ST_ConstraintRelationship
	ForNameAttr *string
	PtTypeAttr  ST_ElementType
}

func (_beee ST_ChildDirection) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_ceadb := _d.Attr{}
	_ceadb.Name = name
	switch _beee {
	case ST_ChildDirectionUnset:
		_ceadb.Value = ""
	case ST_ChildDirectionHorz:
		_ceadb.Value = "\u0068\u006f\u0072\u007a"
	case ST_ChildDirectionVert:
		_ceadb.Value = "\u0076\u0065\u0072\u0074"
	}
	return _ceadb, nil
}

type LayoutDefHdr struct{ CT_DiagramDefinitionHeader }

// ST_FunctionArgument is a union type
type ST_FunctionArgument struct{ ST_VariableType ST_VariableType }

// ST_LayoutShapeType is a union type
type ST_LayoutShapeType struct {
	ST_ShapeType       _ca.ST_ShapeType
	ST_OutputShapeType ST_OutputShapeType
}

func (_aedg ST_CenterShapeMapping) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_aedg.String(), start)
}
func (_faa *AG_IteratorAttributes) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _faa.AxisAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_faa.AxisAttr)})
	}
	if _faa.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _c.Sprintf("\u0025\u0076", *_faa.PtTypeAttr)})
	}
	if _faa.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _c.Sprintf("\u0025\u0076", *_faa.HideLastTransAttr)})
	}
	if _faa.StAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_faa.StAttr)})
	}
	if _faa.CntAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0063\u006e\u0074"}, Value: _c.Sprintf("\u0025\u0076", *_faa.CntAttr)})
	}
	if _faa.StepAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _c.Sprintf("\u0025\u0076", *_faa.StepAttr)})
	}
	return nil
}

// ValidateWithPath validates the CT_SDCategory and its children, prefixing error messages with path
func (_abaa *CT_SDCategory) ValidateWithPath(path string) error { return nil }

// Validate validates the CT_NumericRule and its children
func (_daeb *CT_NumericRule) Validate() error {
	return _daeb.ValidateWithPath("\u0043\u0054\u005f\u004e\u0075\u006d\u0065\u0072\u0069c\u0052\u0075\u006c\u0065")
}
func (_faee ST_CxnType) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_faee.String(), start)
}
func NewCT_StyleDefinitionHeader() *CT_StyleDefinitionHeader {
	_ffdg := &CT_StyleDefinitionHeader{}
	return _ffdg
}

const (
	ST_BoolOperatorUnset ST_BoolOperator = 0
	ST_BoolOperatorNone  ST_BoolOperator = 1
	ST_BoolOperatorEqu   ST_BoolOperator = 2
	ST_BoolOperatorGte   ST_BoolOperator = 3
	ST_BoolOperatorLte   ST_BoolOperator = 4
)

type CT_SDDescription struct {
	LangAttr *string
	ValAttr  string
}

func NewCT_Colors() *CT_Colors { _eefd := &CT_Colors{}; return _eefd }
func (_gaaba *ST_GrowDirection) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_gaaba = 0
	case "\u0074\u004c":
		*_gaaba = 1
	case "\u0074\u0052":
		*_gaaba = 2
	case "\u0062\u004c":
		*_gaaba = 3
	case "\u0062\u0052":
		*_gaaba = 4
	}
	return nil
}

// ValidateWithPath validates the AG_ConstraintRefAttributes and its children, prefixing error messages with path
func (_bec *AG_ConstraintRefAttributes) ValidateWithPath(path string) error {
	if _aab := _bec.RefTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _aab != nil {
		return _aab
	}
	if _bfg := _bec.RefForAttr.ValidateWithPath(path + "/\u0052\u0065\u0066\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _bfg != nil {
		return _bfg
	}
	if _fde := _bec.RefPtTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065\u0066\u0050\u0074\u0054\u0079\u0070e\u0041\u0074\u0074\u0072"); _fde != nil {
		return _fde
	}
	return nil
}
func (_adfgf *CT_SDName) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _adfgf.LangAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _c.Sprintf("\u0025\u0076", *_adfgf.LangAttr)})
	}
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0076\u0061\u006c"}, Value: _c.Sprintf("\u0025\u0076", _adfgf.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

const (
	ST_VariableTypeUnset         ST_VariableType = 0
	ST_VariableTypeNone          ST_VariableType = 1
	ST_VariableTypeOrgChart      ST_VariableType = 2
	ST_VariableTypeChMax         ST_VariableType = 3
	ST_VariableTypeChPref        ST_VariableType = 4
	ST_VariableTypeBulEnabled    ST_VariableType = 5
	ST_VariableTypeDir           ST_VariableType = 6
	ST_VariableTypeHierBranch    ST_VariableType = 7
	ST_VariableTypeAnimOne       ST_VariableType = 8
	ST_VariableTypeAnimLvl       ST_VariableType = 9
	ST_VariableTypeResizeHandles ST_VariableType = 10
)

func (_fbcab *ST_AxisType) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_fbcab = 0
	case "\u0073\u0065\u006c\u0066":
		*_fbcab = 1
	case "\u0063\u0068":
		*_fbcab = 2
	case "\u0064\u0065\u0073":
		*_fbcab = 3
	case "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf":
		*_fbcab = 4
	case "\u0070\u0061\u0072":
		*_fbcab = 5
	case "\u0061\u006e\u0063s\u0074":
		*_fbcab = 6
	case "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066":
		*_fbcab = 7
	case "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b":
		*_fbcab = 8
	case "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b":
		*_fbcab = 9
	case "\u0066\u006f\u006c\u006c\u006f\u0077":
		*_fbcab = 10
	case "\u0070\u0072\u0065\u0063\u0065\u0064":
		*_fbcab = 11
	case "\u0072\u006f\u006f\u0074":
		*_fbcab = 12
	case "\u006e\u006f\u006e\u0065":
		*_fbcab = 13
	}
	return nil
}

type ST_PtType byte

func (_efgc *ST_DiagramTextAlignment) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_efgc = 0
	case "\u006c":
		*_efgc = 1
	case "\u0063\u0074\u0072":
		*_efgc = 2
	case "\u0072":
		*_efgc = 3
	}
	return nil
}

// Validate validates the AG_ConstraintRefAttributes and its children
func (_fag *AG_ConstraintRefAttributes) Validate() error {
	return _fag.ValidateWithPath("\u0041\u0047\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069n\u0074\u0052\u0065\u0066\u0041\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0073")
}
func (_abcc ST_HierBranchStyle) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	return e.EncodeElement(_abcc.String(), start)
}

// ValidateWithPath validates the AG_IteratorAttributes and its children, prefixing error messages with path
func (_fe *AG_IteratorAttributes) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the CT_Choose and its children, prefixing error messages with path
func (_abfd *CT_Choose) ValidateWithPath(path string) error {
	for _ccb, _gdbd := range _abfd.If {
		if _gec := _gdbd.ValidateWithPath(_c.Sprintf("\u0025s\u002f\u0049\u0066\u005b\u0025\u0064]", path, _ccb)); _gec != nil {
			return _gec
		}
	}
	if _abfd.Else != nil {
		if _fcf := _abfd.Else.ValidateWithPath(path + "\u002f\u0045\u006cs\u0065"); _fcf != nil {
			return _fcf
		}
	}
	return nil
}

const (
	ST_ResizeHandlesStrUnset ST_ResizeHandlesStr = 0
	ST_ResizeHandlesStrExact ST_ResizeHandlesStr = 1
	ST_ResizeHandlesStrRel   ST_ResizeHandlesStr = 2
)

type CT_CTStyleLabel struct {
	NameAttr       string
	FillClrLst     *CT_Colors
	LinClrLst      *CT_Colors
	EffectClrLst   *CT_Colors
	TxLinClrLst    *CT_Colors
	TxFillClrLst   *CT_Colors
	TxEffectClrLst *CT_Colors
	ExtLst         *_ca.CT_OfficeArtExtensionList
}

// ValidateWithPath validates the CT_Constraint and its children, prefixing error messages with path
func (_eebg *CT_Constraint) ValidateWithPath(path string) error {
	if _aae := _eebg.OpAttr.ValidateWithPath(path + "\u002fO\u0070\u0041\u0074\u0074\u0072"); _aae != nil {
		return _aae
	}
	if _eebg.ExtLst != nil {
		if _ddbce := _eebg.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _ddbce != nil {
			return _ddbce
		}
	}
	if _eada := _eebg.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _eada != nil {
		return _eada
	}
	if _gbcfd := _eebg.ForAttr.ValidateWithPath(path + "\u002f\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _gbcfd != nil {
		return _gbcfd
	}
	if _fdgg := _eebg.PtTypeAttr.ValidateWithPath(path + "/\u0050\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _fdgg != nil {
		return _fdgg
	}
	if _ffg := _eebg.RefTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _ffg != nil {
		return _ffg
	}
	if _dbgb := _eebg.RefForAttr.ValidateWithPath(path + "/\u0052\u0065\u0066\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _dbgb != nil {
		return _dbgb
	}
	if _bcgd := _eebg.RefPtTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065\u0066\u0050\u0074\u0054\u0079\u0070e\u0041\u0074\u0074\u0072"); _bcgd != nil {
		return _bcgd
	}
	return nil
}

// ValidateWithPath validates the StyleDef and its children, prefixing error messages with path
func (_cbdc *StyleDef) ValidateWithPath(path string) error {
	if _fcade := _cbdc.CT_StyleDefinition.ValidateWithPath(path); _fcade != nil {
		return _fcade
	}
	return nil
}
func (_afdab ST_AlgorithmType) String() string {
	switch _afdab {
	case 0:
		return ""
	case 1:
		return "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e"
	case 2:
		return "\u0063\u006f\u006e\u006e"
	case 3:
		return "\u0063\u0079\u0063l\u0065"
	case 4:
		return "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd"
	case 5:
		return "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074"
	case 6:
		return "\u0070\u0079\u0072\u0061"
	case 7:
		return "\u006c\u0069\u006e"
	case 8:
		return "\u0073\u0070"
	case 9:
		return "\u0074\u0078"
	case 10:
		return "\u0073\u006e\u0061k\u0065"
	}
	return ""
}
func ParseUnionST_ModelId(s string) (ST_ModelId, error) { return ST_ModelId{}, nil }
func (_adgdb *ST_Breakpoint) UnmarshalXMLAttr(attr _d.Attr) error {
	switch attr.Value {
	case "":
		*_adgdb = 0
	case "\u0065\u006e\u0064\u0043\u006e\u0076":
		*_adgdb = 1
	case "\u0062\u0061\u006c":
		*_adgdb = 2
	case "\u0066\u0069\u0078e\u0064":
		*_adgdb = 3
	}
	return nil
}

// ValidateWithPath validates the CT_DiagramDefinition and its children, prefixing error messages with path
func (_febb *CT_DiagramDefinition) ValidateWithPath(path string) error {
	for _bcdg, _fddf := range _febb.Title {
		if _bede := _fddf.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _bcdg)); _bede != nil {
			return _bede
		}
	}
	for _fafc, _caaa := range _febb.Desc {
		if _efcg := _caaa.ValidateWithPath(_c.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _fafc)); _efcg != nil {
			return _efcg
		}
	}
	if _febb.CatLst != nil {
		if _bffd := _febb.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _bffd != nil {
			return _bffd
		}
	}
	if _febb.SampData != nil {
		if _bgeb := _febb.SampData.ValidateWithPath(path + "\u002fS\u0061\u006d\u0070\u0044\u0061\u0074a"); _bgeb != nil {
			return _bgeb
		}
	}
	if _febb.StyleData != nil {
		if _gefe := _febb.StyleData.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065\u0044\u0061\u0074\u0061"); _gefe != nil {
			return _gefe
		}
	}
	if _febb.ClrData != nil {
		if _bcbc := _febb.ClrData.ValidateWithPath(path + "\u002f\u0043\u006c\u0072\u0044\u0061\u0074\u0061"); _bcbc != nil {
			return _bcbc
		}
	}
	if _ebac := _febb.LayoutNode.ValidateWithPath(path + "/\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"); _ebac != nil {
		return _ebac
	}
	if _febb.ExtLst != nil {
		if _cccd := _febb.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _cccd != nil {
			return _cccd
		}
	}
	return nil
}
func (_ddgcf ST_VerticalAlignment) Validate() error { return _ddgcf.ValidateWithPath("") }
func (_ggda ST_PtType) MarshalXMLAttr(name _d.Name) (_d.Attr, error) {
	_cgbd := _d.Attr{}
	_cgbd.Name = name
	switch _ggda {
	case ST_PtTypeUnset:
		_cgbd.Value = ""
	case ST_PtTypeNode:
		_cgbd.Value = "\u006e\u006f\u0064\u0065"
	case ST_PtTypeAsst:
		_cgbd.Value = "\u0061\u0073\u0073\u0074"
	case ST_PtTypeDoc:
		_cgbd.Value = "\u0064\u006f\u0063"
	case ST_PtTypePres:
		_cgbd.Value = "\u0070\u0072\u0065\u0073"
	case ST_PtTypeParTrans:
		_cgbd.Value = "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case ST_PtTypeSibTrans:
		_cgbd.Value = "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return _cgbd, nil
}

// ValidateWithPath validates the CT_PtList and its children, prefixing error messages with path
func (_dafa *CT_PtList) ValidateWithPath(path string) error {
	for _ggabc, _geff := range _dafa.Pt {
		if _bdfab := _geff.ValidateWithPath(_c.Sprintf("\u0025s\u002f\u0050\u0074\u005b\u0025\u0064]", path, _ggabc)); _bdfab != nil {
			return _bdfab
		}
	}
	return nil
}
func NewCT_ColorTransformHeaderLst() *CT_ColorTransformHeaderLst {
	_dcf := &CT_ColorTransformHeaderLst{}
	return _dcf
}

const (
	ST_AlgorithmTypeUnset     ST_AlgorithmType = 0
	ST_AlgorithmTypeComposite ST_AlgorithmType = 1
	ST_AlgorithmTypeConn      ST_AlgorithmType = 2
	ST_AlgorithmTypeCycle     ST_AlgorithmType = 3
	ST_AlgorithmTypeHierChild ST_AlgorithmType = 4
	ST_AlgorithmTypeHierRoot  ST_AlgorithmType = 5
	ST_AlgorithmTypePyra      ST_AlgorithmType = 6
	ST_AlgorithmTypeLin       ST_AlgorithmType = 7
	ST_AlgorithmTypeSp        ST_AlgorithmType = 8
	ST_AlgorithmTypeTx        ST_AlgorithmType = 9
	ST_AlgorithmTypeSnake     ST_AlgorithmType = 10
)

func init() {
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0043\u0054\u004e\u0061\u006de", NewCT_CTName)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0043\u0054\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", NewCT_CTDescription)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079", NewCT_CTCategory)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fC\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073", NewCT_CTCategories)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0043\u006f\u006c\u006f\u0072s", NewCT_Colors)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fC\u0054\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c", NewCT_CTStyleLabel)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006es\u0066\u006f\u0072\u006d", NewCT_ColorTransform)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fCo\u006c\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065\u0072", NewCT_ColorTransformHeader)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061n\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065r\u004c\u0073\u0074", NewCT_ColorTransformHeaderLst)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fP\u0074", NewCT_Pt)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0050\u0074\u004c\u0069\u0073t", NewCT_PtList)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0078\u006e", NewCT_Cxn)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0078\u006e\u004c\u0069\u0073\u0074", NewCT_CxnList)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fD\u0061\u0074\u0061\u004d\u006f\u0064\u0065\u006c", NewCT_DataModel)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074", NewCT_Constraint)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072a\u0069\u006e\u0074\u0073", NewCT_Constraints)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u004e\u0075\u006d\u0065\u0072\u0069c\u0052\u0075\u006c\u0065", NewCT_NumericRule)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0052\u0075\u006c\u0065\u0073", NewCT_Rules)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u004f\u0066", NewCT_PresentationOf)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0041\u0064\u006a", NewCT_Adj)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0041\u0064\u006a\u004c\u0073t", NewCT_AdjLst)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065", NewCT_Shape)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fP\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072", NewCT_Parameter)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fA\u006c\u0067\u006f\u0072\u0069\u0074\u0068\u006d", NewCT_Algorithm)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065", NewCT_LayoutNode)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0046\u006f\u0072\u0045\u0061\u0063\u0068", NewCT_ForEach)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0057\u0068\u0065\u006e", NewCT_When)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fO\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065", NewCT_Otherwise)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0043\u0068\u006f\u006f\u0073e", NewCT_Choose)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0061\u006d\u0070\u006c\u0065\u0044\u0061\u0074\u0061", NewCT_SampleData)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079", NewCT_Category)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073", NewCT_Categories)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u004e\u0061\u006d\u0065", NewCT_Name)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0044\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e", NewCT_Description)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069ti\u006f\u006e", NewCT_DiagramDefinition)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065\u0072", NewCT_DiagramDefinitionHeader)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065rL\u0073\u0074", NewCT_DiagramDefinitionHeaderLst)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0052\u0065\u006c\u0049\u0064s", NewCT_RelIds)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0045\u006c\u0065\u006d\u0050\u0072o\u0070\u0053\u0065\u0074", NewCT_ElemPropSet)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u004f\u0072\u0067\u0043\u0068\u0061\u0072\u0074", NewCT_OrgChart)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u004d\u0061\u0078", NewCT_ChildMax)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fC\u0068\u0069\u006c\u0064\u0050\u0072\u0065\u0066", NewCT_ChildPref)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0042\u0075\u006cl\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064", NewCT_BulletEnabled)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fD\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e", NewCT_Direction)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fH\u0069\u0065\u0072\u0042\u0072a\u006e\u0063h\u0053\u0074\u0079\u006c\u0065", NewCT_HierBranchStyle)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065", NewCT_AnimOne)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c", NewCT_AnimLvl)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0052\u0065\u0073i\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073", NewCT_ResizeHandles)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u004ca\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0069\u0061\u0062l\u0065P\u0072\u006f\u0070\u0065\u0072\u0074\u0079S\u0065\u0074", NewCT_LayoutVariablePropertySet)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0053\u0044\u004e\u0061\u006de", NewCT_SDName)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0053\u0044\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", NewCT_SDDescription)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079", NewCT_SDCategory)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fS\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073", NewCT_SDCategories)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fT\u0065\u0078\u0074\u0050\u0072\u006f\u0070\u0073", NewCT_TextProps)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c", NewCT_StyleLabel)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fS\u0074\u0079\u006c\u0065\u0044e\u0066\u0069n\u0069\u0074\u0069\u006f\u006e", NewCT_StyleDefinition)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006ei\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065\u0072", NewCT_StyleDefinitionHeader)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u0053\u0074\u0079\u006ce\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006fn\u0048\u0065\u0061d\u0065r\u004c\u0073\u0074", NewCT_StyleDefinitionHeaderLst)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0063o\u006c\u006f\u0072\u0073\u0044\u0065f", NewColorsDef)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072", NewColorsDefHdr)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0063o\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074", NewColorsDefHdrLst)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l", NewDataModel)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u006ca\u0079\u006f\u0075\u0074\u0044\u0065f", NewLayoutDef)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072", NewLayoutDefHdr)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u006ca\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074", NewLayoutDefHdrLst)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0072\u0065\u006c\u0049\u0064\u0073", NewRelIds)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066", NewStyleDef)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072", NewStyleDefHdr)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048d\u0072\u004c\u0073\u0074", NewStyleDefHdrLst)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "A\u0047\u005f\u0049\u0074er\u0061t\u006f\u0072\u0041\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u0073", NewAG_IteratorAttributes)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0041\u0047\u005fCo\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074\u0041\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0073", NewAG_ConstraintAttributes)
	_af.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0041\u0047\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069n\u0074\u0052\u0065\u0066\u0041\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0073", NewAG_ConstraintRefAttributes)
}
func (_aaddb ST_HierarchyAlignment) Validate() error { return _aaddb.ValidateWithPath("") }
func (_aebf ST_AlgorithmType) ValidateWithPath(path string) error {
	switch _aebf {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return _c.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aebf))
	}
	return nil
}

type ST_ParameterId byte

func NewDataModel() *DataModel {
	_debg := &DataModel{}
	_debg.CT_DataModel = *NewCT_DataModel()
	return _debg
}

type CT_SDCategories struct{ Cat []*CT_SDCategory }
