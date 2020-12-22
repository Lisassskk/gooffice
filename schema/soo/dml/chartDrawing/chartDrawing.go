package chartDrawing

import (
	_c "encoding/xml"
	_aa "fmt"
	_d "strconv"

	_g "gitee.com/yu_sheng/gooffice"
	_f "gitee.com/yu_sheng/gooffice/schema/soo/dml"
)

func NewEG_ObjectChoices() *EG_ObjectChoices { _gbd := &EG_ObjectChoices{}; return _gbd }

// Validate validates the CT_Shape and its children
func (_ecce *CT_Shape) Validate() error {
	return _ecce.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065")
}
func NewCT_GroupShape() *CT_GroupShape {
	_gcc := &CT_GroupShape{}
	_gcc.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_gcc.GrpSpPr = _f.NewCT_GroupShapeProperties()
	return _gcc
}

type CT_GroupShapeNonVisual struct {
	CNvPr      *_f.CT_NonVisualDrawingProps
	CNvGrpSpPr *_f.CT_NonVisualGroupDrawingShapeProps
}

func (_cba *CT_GroupShape) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_bea := _c.StartElement{Name: _c.Name{Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}}
	e.EncodeElement(_cba.NvGrpSpPr, _bea)
	_gca := _c.StartElement{Name: _c.Name{Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_cba.GrpSpPr, _gca)
	if _cba.Choice != nil {
		for _, _aeg := range _cba.Choice {
			_aeg.MarshalXML(e, _c.StartElement{})
		}
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Marker and its children, prefixing error messages with path
func (_dceg *CT_Marker) ValidateWithPath(path string) error {
	if _dceg.X < 0.0 {
		return _aa.Errorf("\u0025\u0073\u002fm\u002e\u0058\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0030\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _dceg.X)
	}
	if _dceg.X > 1.0 {
		return _aa.Errorf("\u0025\u0073\u002fm\u002e\u0058\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0031\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _dceg.X)
	}
	if _dceg.Y < 0.0 {
		return _aa.Errorf("\u0025\u0073\u002fm\u002e\u0059\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0030\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _dceg.Y)
	}
	if _dceg.Y > 1.0 {
		return _aa.Errorf("\u0025\u0073\u002fm\u002e\u0059\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0031\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _dceg.Y)
	}
	return nil
}
func NewCT_Picture() *CT_Picture {
	_fga := &CT_Picture{}
	_fga.NvPicPr = NewCT_PictureNonVisual()
	_fga.BlipFill = _f.NewCT_BlipFillProperties()
	_fga.SpPr = _f.NewCT_ShapeProperties()
	return _fga
}
func (_bfbe *CT_RelSizeAnchor) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_egc := _c.StartElement{Name: _c.Name{Local: "\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_bfbe.From, _egc)
	_dcff := _c.StartElement{Name: _c.Name{Local: "\u0074\u006f"}}
	e.EncodeElement(_bfbe.To, _dcff)
	if _bfbe.Choice != nil {
		_bfbe.Choice.MarshalXML(e, _c.StartElement{})
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_aagc *CT_Marker) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_aagc.X = 0.0
	_aagc.Y = 0.0
_egd:
	for {
		_gcfg, _gdd := d.Token()
		if _gdd != nil {
			return _gdd
		}
		switch _gac := _gcfg.(type) {
		case _c.StartElement:
			switch _gac.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0078"}:
				if _ggb := d.DecodeElement(&_aagc.X, &_gac); _ggb != nil {
					return _ggb
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0079"}:
				if _cgb := d.DecodeElement(&_aagc.Y, &_gac); _cgb != nil {
					return _cgb
				}
			default:
				_g.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004d\u0061\u0072k\u0065\u0072 \u0025\u0076", _gac.Name)
				if _ccb := d.Skip(); _ccb != nil {
					return _ccb
				}
			}
		case _c.EndElement:
			break _egd
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicFrame and its children
func (_ada *CT_GraphicFrame) Validate() error {
	return _ada.ValidateWithPath("\u0043T\u005fG\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065")
}
func (_ffd *CT_AbsSizeAnchor) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_fd := _c.StartElement{Name: _c.Name{Local: "\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_ffd.From, _fd)
	_ca := _c.StartElement{Name: _c.Name{Local: "\u0065\u0078\u0074"}}
	e.EncodeElement(_ffd.Ext, _ca)
	if _ffd.Choice != nil {
		_ffd.Choice.MarshalXML(e, _c.StartElement{})
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_GroupShapeChoice and its children, prefixing error messages with path
func (_cfde *CT_GroupShapeChoice) ValidateWithPath(path string) error {
	for _cfc, _ggd := range _cfde.Sp {
		if _aagf := _ggd.ValidateWithPath(_aa.Sprintf("\u0025s\u002f\u0053\u0070\u005b\u0025\u0064]", path, _cfc)); _aagf != nil {
			return _aagf
		}
	}
	for _cfg, _bac := range _cfde.GrpSp {
		if _ecbb := _bac.ValidateWithPath(_aa.Sprintf("\u0025\u0073\u002fG\u0072\u0070\u0053\u0070\u005b\u0025\u0064\u005d", path, _cfg)); _ecbb != nil {
			return _ecbb
		}
	}
	for _fa, _dage := range _cfde.GraphicFrame {
		if _aga := _dage.ValidateWithPath(_aa.Sprintf("\u0025\u0073\u002f\u0047ra\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u005b\u0025\u0064\u005d", path, _fa)); _aga != nil {
			return _aga
		}
	}
	for _fecg, _afd := range _cfde.CxnSp {
		if _bcg := _afd.ValidateWithPath(_aa.Sprintf("\u0025\u0073\u002fC\u0078\u006e\u0053\u0070\u005b\u0025\u0064\u005d", path, _fecg)); _bcg != nil {
			return _bcg
		}
	}
	for _cgf, _bdfe := range _cfde.Pic {
		if _cfb := _bdfe.ValidateWithPath(_aa.Sprintf("\u0025\u0073\u002f\u0050\u0069\u0063\u005b\u0025\u0064\u005d", path, _cgf)); _cfb != nil {
			return _cfb
		}
	}
	return nil
}

type CT_Shape struct {
	MacroAttr      *string
	TextlinkAttr   *string
	FLocksTextAttr *bool
	FPublishedAttr *bool
	NvSpPr         *CT_ShapeNonVisual
	SpPr           *_f.CT_ShapeProperties
	Style          *_f.CT_ShapeStyle
	TxBody         *_f.CT_TextBody
}

func (_bfb *CT_Connector) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_bfb.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_bfb.SpPr = _f.NewCT_ShapeProperties()
	for _, _ea := range start.Attr {
		if _ea.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_fce, _daa := _ea.Value, error(nil)
			if _daa != nil {
				return _daa
			}
			_bfb.MacroAttr = &_fce
			continue
		}
		if _ea.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_dd, _de := _d.ParseBool(_ea.Value)
			if _de != nil {
				return _de
			}
			_bfb.FPublishedAttr = &_dd
			continue
		}
	}
_gf:
	for {
		_bd, _gg := d.Token()
		if _gg != nil {
			return _gg
		}
		switch _dea := _bd.(type) {
		case _c.StartElement:
			switch _dea.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}:
				if _gd := d.DecodeElement(_bfb.NvCxnSpPr, &_dea); _gd != nil {
					return _gd
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _gbe := d.DecodeElement(_bfb.SpPr, &_dea); _gbe != nil {
					return _gbe
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_bfb.Style = _f.NewCT_ShapeStyle()
				if _gfd := d.DecodeElement(_bfb.Style, &_dea); _gfd != nil {
					return _gfd
				}
			default:
				_g.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_C\u006f\u006en\u0065\u0063\u0074\u006f\u0072\u0020\u0025\u0076", _dea.Name)
				if _caa := d.Skip(); _caa != nil {
					return _caa
				}
			}
		case _c.EndElement:
			break _gf
		case _c.CharData:
		}
	}
	return nil
}

type EG_Anchor struct {
	RelSizeAnchor *CT_RelSizeAnchor
	AbsSizeAnchor *CT_AbsSizeAnchor
}

func (_dbgg *CT_ShapeNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_faa := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_dbgg.CNvPr, _faa)
	_afeg := _c.StartElement{Name: _c.Name{Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_dbgg.CNvSpPr, _afeg)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func NewCT_ConnectorNonVisual() *CT_ConnectorNonVisual {
	_bda := &CT_ConnectorNonVisual{}
	_bda.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_bda.CNvCxnSpPr = _f.NewCT_NonVisualConnectorProperties()
	return _bda
}

type CT_GraphicFrame struct {
	MacroAttr        *string
	FPublishedAttr   *bool
	NvGraphicFramePr *CT_GraphicFrameNonVisual
	Xfrm             *_f.CT_Transform2D
	Graphic          *_f.Graphic
}

func (_aef *CT_GraphicFrameNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_aef.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_aef.CNvGraphicFramePr = _f.NewCT_NonVisualGraphicFrameProperties()
_ege:
	for {
		_fba, _afg := d.Token()
		if _afg != nil {
			return _afg
		}
		switch _bff := _fba.(type) {
		case _c.StartElement:
			switch _bff.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _bgf := d.DecodeElement(_aef.CNvPr, &_bff); _bgf != nil {
					return _bgf
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}:
				if _fgdc := d.DecodeElement(_aef.CNvGraphicFramePr, &_bff); _fgdc != nil {
					return _fgdc
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c\u0020%\u0076", _bff.Name)
				if _fcee := d.Skip(); _fcee != nil {
					return _fcee
				}
			}
		case _c.EndElement:
			break _ege
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Picture and its children
func (_bca *CT_Picture) Validate() error {
	return _bca.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}
func (_bad *CT_GroupShapeChoice) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _bad.Sp != nil {
		_aad := _c.StartElement{Name: _c.Name{Local: "\u0073\u0070"}}
		for _, _abbb := range _bad.Sp {
			e.EncodeElement(_abbb, _aad)
		}
	}
	if _bad.GrpSp != nil {
		_ega := _c.StartElement{Name: _c.Name{Local: "\u0067\u0072\u0070S\u0070"}}
		for _, _acba := range _bad.GrpSp {
			e.EncodeElement(_acba, _ega)
		}
	}
	if _bad.GraphicFrame != nil {
		_efg := _c.StartElement{Name: _c.Name{Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		for _, _bfgc := range _bad.GraphicFrame {
			e.EncodeElement(_bfgc, _efg)
		}
	}
	if _bad.CxnSp != nil {
		_fec := _c.StartElement{Name: _c.Name{Local: "\u0063\u0078\u006eS\u0070"}}
		for _, _afcg := range _bad.CxnSp {
			e.EncodeElement(_afcg, _fec)
		}
	}
	if _bad.Pic != nil {
		_bag := _c.StartElement{Name: _c.Name{Local: "\u0070\u0069\u0063"}}
		for _, _cga := range _bad.Pic {
			e.EncodeElement(_cga, _bag)
		}
	}
	return nil
}
func (_fac *CT_ShapeNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_fac.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_fac.CNvSpPr = _f.NewCT_NonVisualDrawingShapeProps()
_cfbc:
	for {
		_bgce, _cab := d.Token()
		if _cab != nil {
			return _cab
		}
		switch _afb := _bgce.(type) {
		case _c.StartElement:
			switch _afb.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _fcedd := d.DecodeElement(_fac.CNvPr, &_afb); _fcedd != nil {
					return _fcedd
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}:
				if _becb := d.DecodeElement(_fac.CNvSpPr, &_afb); _becb != nil {
					return _becb
				}
			default:
				_g.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _afb.Name)
				if _adea := d.Skip(); _adea != nil {
					return _adea
				}
			}
		case _c.EndElement:
			break _cfbc
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_RelSizeAnchor and its children, prefixing error messages with path
func (_accd *CT_RelSizeAnchor) ValidateWithPath(path string) error {
	if _bee := _accd.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _bee != nil {
		return _bee
	}
	if _dga := _accd.To.ValidateWithPath(path + "\u002f\u0054\u006f"); _dga != nil {
		return _dga
	}
	if _accd.Choice != nil {
		if _cad := _accd.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _cad != nil {
			return _cad
		}
	}
	return nil
}

// ValidateWithPath validates the CT_ConnectorNonVisual and its children, prefixing error messages with path
func (_gc *CT_ConnectorNonVisual) ValidateWithPath(path string) error {
	if _fe := _gc.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _fe != nil {
		return _fe
	}
	if _bdf := _gc.CNvCxnSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _bdf != nil {
		return _bdf
	}
	return nil
}

// ValidateWithPath validates the EG_ObjectChoices and its children, prefixing error messages with path
func (_bdb *EG_ObjectChoices) ValidateWithPath(path string) error {
	if _bdb.Choice != nil {
		if _befa := _bdb.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _befa != nil {
			return _befa
		}
	}
	return nil
}
func NewCT_Connector() *CT_Connector {
	_acb := &CT_Connector{}
	_acb.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_acb.SpPr = _f.NewCT_ShapeProperties()
	return _acb
}
func (_geeb *CT_GroupShapeNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_geeb.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_geeb.CNvGrpSpPr = _f.NewCT_NonVisualGroupDrawingShapeProps()
_bfgd:
	for {
		_gfa, _bbg := d.Token()
		if _bbg != nil {
			return _bbg
		}
		switch _cdd := _gfa.(type) {
		case _c.StartElement:
			switch _cdd.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _ffce := d.DecodeElement(_geeb.CNvPr, &_cdd); _ffce != nil {
					return _ffce
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}:
				if _bec := d.DecodeElement(_geeb.CNvGrpSpPr, &_cdd); _bec != nil {
					return _bec
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0047\u0072\u006f\u0075p\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _cdd.Name)
				if _gbgf := d.Skip(); _gbgf != nil {
					return _gbgf
				}
			}
		case _c.EndElement:
			break _bfgd
		case _c.CharData:
		}
	}
	return nil
}

type CT_Picture struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvPicPr        *CT_PictureNonVisual
	BlipFill       *_f.CT_BlipFillProperties
	SpPr           *_f.CT_ShapeProperties
	Style          *_f.CT_ShapeStyle
}
type EG_ObjectChoicesChoice struct {
	Sp           *CT_Shape
	GrpSp        *CT_GroupShape
	GraphicFrame *CT_GraphicFrame
	CxnSp        *CT_Connector
	Pic          *CT_Picture
}

// Validate validates the CT_ShapeNonVisual and its children
func (_aaf *CT_ShapeNonVisual) Validate() error {
	return _aaf.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c")
}
func NewCT_Marker() *CT_Marker { _ceg := &CT_Marker{}; _ceg.X = 0.0; _ceg.Y = 0.0; return _ceg }

type CT_Drawing struct{ EG_Anchor []*EG_Anchor }

func (_fced *CT_PictureNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_gfe := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_fced.CNvPr, _gfe)
	_bgc := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_fced.CNvPicPr, _bgc)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_GraphicFrame and its children, prefixing error messages with path
func (_gee *CT_GraphicFrame) ValidateWithPath(path string) error {
	if _fdd := _gee.NvGraphicFramePr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"); _fdd != nil {
		return _fdd
	}
	if _bfa := _gee.Xfrm.ValidateWithPath(path + "\u002f\u0058\u0066r\u006d"); _bfa != nil {
		return _bfa
	}
	if _fbe := _gee.Graphic.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063"); _fbe != nil {
		return _fbe
	}
	return nil
}
func (_fdf *CT_Drawing) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Name.Local = "\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067"
	e.EncodeToken(start)
	if _fdf.EG_Anchor != nil {
		for _, _df := range _fdf.EG_Anchor {
			_df.MarshalXML(e, _c.StartElement{})
		}
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_fgg *CT_GraphicFrame) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _fgg.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _aa.Sprintf("\u0025\u0076", *_fgg.MacroAttr)})
	}
	if _fgg.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _aa.Sprintf("\u0025\u0064", _afcb(*_fgg.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_fed := _c.StartElement{Name: _c.Name{Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}}
	e.EncodeElement(_fgg.NvGraphicFramePr, _fed)
	_fgf := _c.StartElement{Name: _c.Name{Local: "\u0078\u0066\u0072\u006d"}}
	e.EncodeElement(_fgg.Xfrm, _fgf)
	_cb := _c.StartElement{Name: _c.Name{Local: "\u0061:\u0067\u0072\u0061\u0070\u0068\u0069c"}}
	e.EncodeElement(_fgg.Graphic, _cb)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_PictureNonVisual and its children
func (_eccd *CT_PictureNonVisual) Validate() error {
	return _eccd.ValidateWithPath("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

// ValidateWithPath validates the CT_Connector and its children, prefixing error messages with path
func (_acc *CT_Connector) ValidateWithPath(path string) error {
	if _ba := _acc.NvCxnSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _ba != nil {
		return _ba
	}
	if _dg := _acc.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _dg != nil {
		return _dg
	}
	if _acc.Style != nil {
		if _ae := _acc.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _ae != nil {
			return _ae
		}
	}
	return nil
}
func (_fg *CT_AbsSizeAnchor) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_fg.From = NewCT_Marker()
	_fg.Ext = _f.NewCT_PositiveSize2D()
_b:
	for {
		_ge, _fgd := d.Token()
		if _fgd != nil {
			return _fgd
		}
		switch _e := _ge.(type) {
		case _c.StartElement:
			switch _e.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _da := d.DecodeElement(_fg.From, &_e); _da != nil {
					return _da
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}:
				if _gb := d.DecodeElement(_fg.Ext, &_e); _gb != nil {
					return _gb
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_fg.Choice = NewEG_ObjectChoicesChoice()
				if _ee := d.DecodeElement(&_fg.Choice.Sp, &_e); _ee != nil {
					return _ee
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_fg.Choice = NewEG_ObjectChoicesChoice()
				if _bc := d.DecodeElement(&_fg.Choice.GrpSp, &_e); _bc != nil {
					return _bc
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_fg.Choice = NewEG_ObjectChoicesChoice()
				if _fb := d.DecodeElement(&_fg.Choice.GraphicFrame, &_e); _fb != nil {
					return _fb
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_fg.Choice = NewEG_ObjectChoicesChoice()
				if _bf := d.DecodeElement(&_fg.Choice.CxnSp, &_e); _bf != nil {
					return _bf
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_fg.Choice = NewEG_ObjectChoicesChoice()
				if _ga := d.DecodeElement(&_fg.Choice.Pic, &_e); _ga != nil {
					return _ga
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _e.Name)
				if _fc := d.Skip(); _fc != nil {
					return _fc
				}
			}
		case _c.EndElement:
			break _b
		case _c.CharData:
		}
	}
	return nil
}
func (_fggd *CT_RelSizeAnchor) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_fggd.From = NewCT_Marker()
	_fggd.To = NewCT_Marker()
_dbad:
	for {
		_ffcb, _ebg := d.Token()
		if _ebg != nil {
			return _ebg
		}
		switch _afa := _ffcb.(type) {
		case _c.StartElement:
			switch _afa.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _ecaa := d.DecodeElement(_fggd.From, &_afa); _ecaa != nil {
					return _ecaa
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u006f"}:
				if _cegd := d.DecodeElement(_fggd.To, &_afa); _cegd != nil {
					return _cegd
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_fggd.Choice = NewEG_ObjectChoicesChoice()
				if _bcfc := d.DecodeElement(&_fggd.Choice.Sp, &_afa); _bcfc != nil {
					return _bcfc
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_fggd.Choice = NewEG_ObjectChoicesChoice()
				if _edc := d.DecodeElement(&_fggd.Choice.GrpSp, &_afa); _edc != nil {
					return _edc
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_fggd.Choice = NewEG_ObjectChoicesChoice()
				if _dda := d.DecodeElement(&_fggd.Choice.GraphicFrame, &_afa); _dda != nil {
					return _dda
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_fggd.Choice = NewEG_ObjectChoicesChoice()
				if _dbg := d.DecodeElement(&_fggd.Choice.CxnSp, &_afa); _dbg != nil {
					return _dbg
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_fggd.Choice = NewEG_ObjectChoicesChoice()
				if _bgca := d.DecodeElement(&_fggd.Choice.Pic, &_afa); _bgca != nil {
					return _bgca
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _afa.Name)
				if _egcd := d.Skip(); _egcd != nil {
					return _egcd
				}
			}
		case _c.EndElement:
			break _dbad
		case _c.CharData:
		}
	}
	return nil
}
func (_edcb *CT_Shape) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _edcb.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _aa.Sprintf("\u0025\u0076", *_edcb.MacroAttr)})
	}
	if _edcb.TextlinkAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b"}, Value: _aa.Sprintf("\u0025\u0076", *_edcb.TextlinkAttr)})
	}
	if _edcb.FLocksTextAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074"}, Value: _aa.Sprintf("\u0025\u0064", _afcb(*_edcb.FLocksTextAttr))})
	}
	if _edcb.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _aa.Sprintf("\u0025\u0064", _afcb(*_edcb.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_cafe := _c.StartElement{Name: _c.Name{Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_edcb.NvSpPr, _cafe)
	_abdd := _c.StartElement{Name: _c.Name{Local: "\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_edcb.SpPr, _abdd)
	if _edcb.Style != nil {
		_bbdg := _c.StartElement{Name: _c.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_edcb.Style, _bbdg)
	}
	if _edcb.TxBody != nil {
		_ffeb := _c.StartElement{Name: _c.Name{Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}}
		e.EncodeElement(_edcb.TxBody, _ffeb)
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_AbsSizeAnchor and its children
func (_cg *CT_AbsSizeAnchor) Validate() error {
	return _cg.ValidateWithPath("\u0043\u0054_\u0041\u0062\u0073S\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072")
}
func NewCT_GroupShapeNonVisual() *CT_GroupShapeNonVisual {
	_ffe := &CT_GroupShapeNonVisual{}
	_ffe.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_ffe.CNvGrpSpPr = _f.NewCT_NonVisualGroupDrawingShapeProps()
	return _ffe
}

type CT_AbsSizeAnchor struct {
	From   *CT_Marker
	Ext    *_f.CT_PositiveSize2D
	Choice *EG_ObjectChoicesChoice
}

// Validate validates the CT_Marker and its children
func (_gfcc *CT_Marker) Validate() error {
	return _gfcc.ValidateWithPath("\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r")
}

// Validate validates the CT_GroupShapeChoice and its children
func (_eac *CT_GroupShapeChoice) Validate() error {
	return _eac.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u0043\u0068\u006f\u0069\u0063\u0065")
}

// ValidateWithPath validates the CT_GroupShape and its children, prefixing error messages with path
func (_bcf *CT_GroupShape) ValidateWithPath(path string) error {
	if _ffc := _bcf.NvGrpSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _ffc != nil {
		return _ffc
	}
	if _eee := _bcf.GrpSpPr.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _eee != nil {
		return _eee
	}
	for _aec, _dcb := range _bcf.Choice {
		if _abd := _dcb.ValidateWithPath(_aa.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _aec)); _abd != nil {
			return _abd
		}
	}
	return nil
}
func (_bbdge *EG_ObjectChoicesChoice) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _bbdge.Sp != nil {
		_gdf := _c.StartElement{Name: _c.Name{Local: "\u0073\u0070"}}
		e.EncodeElement(_bbdge.Sp, _gdf)
	}
	if _bbdge.GrpSp != nil {
		_aecf := _c.StartElement{Name: _c.Name{Local: "\u0067\u0072\u0070S\u0070"}}
		e.EncodeElement(_bbdge.GrpSp, _aecf)
	}
	if _bbdge.GraphicFrame != nil {
		_cddc := _c.StartElement{Name: _c.Name{Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		e.EncodeElement(_bbdge.GraphicFrame, _cddc)
	}
	if _bbdge.CxnSp != nil {
		_fca := _c.StartElement{Name: _c.Name{Local: "\u0063\u0078\u006eS\u0070"}}
		e.EncodeElement(_bbdge.CxnSp, _fca)
	}
	if _bbdge.Pic != nil {
		_eaf := _c.StartElement{Name: _c.Name{Local: "\u0070\u0069\u0063"}}
		e.EncodeElement(_bbdge.Pic, _eaf)
	}
	return nil
}
func NewCT_GroupShapeChoice() *CT_GroupShapeChoice { _daf := &CT_GroupShapeChoice{}; return _daf }
func (_dcec *CT_Marker) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_fae := _c.StartElement{Name: _c.Name{Local: "\u0078"}}
	e.EncodeElement(_dcec.X, _fae)
	_abge := _c.StartElement{Name: _c.Name{Local: "\u0079"}}
	e.EncodeElement(_dcec.Y, _abge)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_GroupShapeNonVisual and its children
func (_ebb *CT_GroupShapeNonVisual) Validate() error {
	return _ebb.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c")
}
func (_cf *CT_Drawing) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_cec:
	for {
		_gae, _aba := d.Token()
		if _aba != nil {
			return _aba
		}
		switch _gdb := _gae.(type) {
		case _c.StartElement:
			switch _gdb.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_eef := NewEG_Anchor()
				_eef.RelSizeAnchor = NewCT_RelSizeAnchor()
				if _add := d.DecodeElement(_eef.RelSizeAnchor, &_gdb); _add != nil {
					return _add
				}
				_cf.EG_Anchor = append(_cf.EG_Anchor, _eef)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_bdfd := NewEG_Anchor()
				_bdfd.AbsSizeAnchor = NewCT_AbsSizeAnchor()
				if _aag := d.DecodeElement(_bdfd.AbsSizeAnchor, &_gdb); _aag != nil {
					return _aag
				}
				_cf.EG_Anchor = append(_cf.EG_Anchor, _bdfd)
			default:
				_g.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fD\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0025\u0076", _gdb.Name)
				if _dc := d.Skip(); _dc != nil {
					return _dc
				}
			}
		case _c.EndElement:
			break _cec
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_ShapeNonVisual and its children, prefixing error messages with path
func (_bgdc *CT_ShapeNonVisual) ValidateWithPath(path string) error {
	if _gde := _bgdc.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _gde != nil {
		return _gde
	}
	if _afdf := _bgdc.CNvSpPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0053\u0070\u0050\u0072"); _afdf != nil {
		return _afdf
	}
	return nil
}

// Validate validates the CT_ConnectorNonVisual and its children
func (_db *CT_ConnectorNonVisual) Validate() error {
	return _db.ValidateWithPath("C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c")
}
func (_efef *CT_Shape) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_efef.NvSpPr = NewCT_ShapeNonVisual()
	_efef.SpPr = _f.NewCT_ShapeProperties()
	for _, _cgbe := range start.Attr {
		if _cgbe.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_ead, _gbf := _cgbe.Value, error(nil)
			if _gbf != nil {
				return _gbf
			}
			_efef.MacroAttr = &_ead
			continue
		}
		if _cgbe.Name.Local == "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b" {
			_cbfd, _def := _cgbe.Value, error(nil)
			if _def != nil {
				return _def
			}
			_efef.TextlinkAttr = &_cbfd
			continue
		}
		if _cgbe.Name.Local == "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074" {
			_aagfc, _bgfd := _d.ParseBool(_cgbe.Value)
			if _bgfd != nil {
				return _bgfd
			}
			_efef.FLocksTextAttr = &_aagfc
			continue
		}
		if _cgbe.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_aeea, _agf := _d.ParseBool(_cgbe.Value)
			if _agf != nil {
				return _agf
			}
			_efef.FPublishedAttr = &_aeea
			continue
		}
	}
_gfge:
	for {
		_acad, _bdg := d.Token()
		if _bdg != nil {
			return _bdg
		}
		switch _dagg := _acad.(type) {
		case _c.StartElement:
			switch _dagg.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}:
				if _fab := d.DecodeElement(_efef.NvSpPr, &_dagg); _fab != nil {
					return _fab
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _edcd := d.DecodeElement(_efef.SpPr, &_dagg); _edcd != nil {
					return _edcd
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_efef.Style = _f.NewCT_ShapeStyle()
				if _eeeg := d.DecodeElement(_efef.Style, &_dagg); _eeeg != nil {
					return _eeeg
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}:
				_efef.TxBody = _f.NewCT_TextBody()
				if _cgd := d.DecodeElement(_efef.TxBody, &_dagg); _cgd != nil {
					return _cgd
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u0020\u0025\u0076", _dagg.Name)
				if _dee := d.Skip(); _dee != nil {
					return _dee
				}
			}
		case _c.EndElement:
			break _gfge
		case _c.CharData:
		}
	}
	return nil
}

type CT_Connector struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvCxnSpPr      *CT_ConnectorNonVisual
	SpPr           *_f.CT_ShapeProperties
	Style          *_f.CT_ShapeStyle
}

// ValidateWithPath validates the CT_AbsSizeAnchor and its children, prefixing error messages with path
func (_ad *CT_AbsSizeAnchor) ValidateWithPath(path string) error {
	if _cgg := _ad.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _cgg != nil {
		return _cgg
	}
	if _ac := _ad.Ext.ValidateWithPath(path + "\u002f\u0045\u0078\u0074"); _ac != nil {
		return _ac
	}
	if _ad.Choice != nil {
		if _ab := _ad.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _ab != nil {
			return _ab
		}
	}
	return nil
}

type CT_PictureNonVisual struct {
	CNvPr    *_f.CT_NonVisualDrawingProps
	CNvPicPr *_f.CT_NonVisualPictureProperties
}

func (_be *CT_ConnectorNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_ce := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_be.CNvPr, _ce)
	_ddf := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_be.CNvCxnSpPr, _ddf)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_adg *CT_ConnectorNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_adg.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_adg.CNvCxnSpPr = _f.NewCT_NonVisualConnectorProperties()
_abb:
	for {
		_cc, _fdg := d.Token()
		if _fdg != nil {
			return _fdg
		}
		switch _dgg := _cc.(type) {
		case _c.StartElement:
			switch _dgg.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _bae := d.DecodeElement(_adg.CNvPr, &_dgg); _bae != nil {
					return _bae
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}:
				if _fcd := d.DecodeElement(_adg.CNvCxnSpPr, &_dgg); _fcd != nil {
					return _fcd
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075n\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u006f\u006e C\u0054\u005f\u0043\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _dgg.Name)
				if _fcb := d.Skip(); _fcb != nil {
					return _fcb
				}
			}
		case _c.EndElement:
			break _abb
		case _c.CharData:
		}
	}
	return nil
}
func (_dgea *EG_Anchor) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_gade:
	for {
		_cgfc, _dbc := d.Token()
		if _dbc != nil {
			return _dbc
		}
		switch _beaa := _cgfc.(type) {
		case _c.StartElement:
			switch _beaa.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_dgea.RelSizeAnchor = NewCT_RelSizeAnchor()
				if _ebbe := d.DecodeElement(_dgea.RelSizeAnchor, &_beaa); _ebbe != nil {
					return _ebbe
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_dgea.AbsSizeAnchor = NewCT_AbsSizeAnchor()
				if _gddc := d.DecodeElement(_dgea.AbsSizeAnchor, &_beaa); _gddc != nil {
					return _gddc
				}
			default:
				_g.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0045\u0047\u005f\u0041\u006e\u0063h\u006f\u0072 \u0025\u0076", _beaa.Name)
				if _fcedc := d.Skip(); _fcedc != nil {
					return _fcedc
				}
			}
		case _c.EndElement:
			break _gade
		case _c.CharData:
		}
	}
	return nil
}
func NewCT_RelSizeAnchor() *CT_RelSizeAnchor {
	_dge := &CT_RelSizeAnchor{}
	_dge.From = NewCT_Marker()
	_dge.To = NewCT_Marker()
	return _dge
}
func (_bed *EG_Anchor) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _bed.RelSizeAnchor != nil {
		_cef := _c.StartElement{Name: _c.Name{Local: "\u0072\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_bed.RelSizeAnchor, _cef)
	}
	if _bed.AbsSizeAnchor != nil {
		_dggf := _c.StartElement{Name: _c.Name{Local: "\u0061\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_bed.AbsSizeAnchor, _dggf)
	}
	return nil
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (_dfc *CT_Drawing) ValidateWithPath(path string) error {
	for _ccd, _eca := range _dfc.EG_Anchor {
		if _bg := _eca.ValidateWithPath(_aa.Sprintf("\u0025\u0073/\u0045\u0047\u005fA\u006e\u0063\u0068\u006f\u0072\u005b\u0025\u0064\u005d", path, _ccd)); _bg != nil {
			return _bg
		}
	}
	return nil
}
func (_baec *CT_PictureNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_baec.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_baec.CNvPicPr = _f.NewCT_NonVisualPictureProperties()
_gceg:
	for {
		_accb, _deb := d.Token()
		if _deb != nil {
			return _deb
		}
		switch _acgd := _accb.(type) {
		case _c.StartElement:
			switch _acgd.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _bffg := d.DecodeElement(_baec.CNvPr, &_acgd); _bffg != nil {
					return _bffg
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:
				if _gece := d.DecodeElement(_baec.CNvPicPr, &_acgd); _gece != nil {
					return _gece
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _acgd.Name)
				if _cgga := d.Skip(); _cgga != nil {
					return _cgga
				}
			}
		case _c.EndElement:
			break _gceg
		case _c.CharData:
		}
	}
	return nil
}
func (_aca *CT_Picture) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _aca.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _aa.Sprintf("\u0025\u0076", *_aca.MacroAttr)})
	}
	if _aca.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _aa.Sprintf("\u0025\u0064", _afcb(*_aca.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_dfda := _c.StartElement{Name: _c.Name{Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_aca.NvPicPr, _dfda)
	_bef := _c.StartElement{Name: _c.Name{Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}}
	e.EncodeElement(_aca.BlipFill, _bef)
	_aefc := _c.StartElement{Name: _c.Name{Local: "\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_aca.SpPr, _aefc)
	if _aca.Style != nil {
		_ade := _c.StartElement{Name: _c.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_aca.Style, _ade)
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type CT_ShapeNonVisual struct {
	CNvPr   *_f.CT_NonVisualDrawingProps
	CNvSpPr *_f.CT_NonVisualDrawingShapeProps
}

func NewCT_ShapeNonVisual() *CT_ShapeNonVisual {
	_bgd := &CT_ShapeNonVisual{}
	_bgd.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_bgd.CNvSpPr = _f.NewCT_NonVisualDrawingShapeProps()
	return _bgd
}

type CT_ConnectorNonVisual struct {
	CNvPr      *_f.CT_NonVisualDrawingProps
	CNvCxnSpPr *_f.CT_NonVisualConnectorProperties
}

func (_cag *CT_GraphicFrameNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_af := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_cag.CNvPr, _af)
	_fdgf := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}}
	e.EncodeElement(_cag.CNvGraphicFramePr, _fdgf)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type EG_ObjectChoices struct{ Choice *EG_ObjectChoicesChoice }

func (_aae *CT_Connector) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _aae.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _aa.Sprintf("\u0025\u0076", *_aae.MacroAttr)})
	}
	if _aae.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _aa.Sprintf("\u0025\u0064", _afcb(*_aae.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_adb := _c.StartElement{Name: _c.Name{Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}}
	e.EncodeElement(_aae.NvCxnSpPr, _adb)
	_ec := _c.StartElement{Name: _c.Name{Local: "\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_aae.SpPr, _ec)
	if _aae.Style != nil {
		_ecg := _c.StartElement{Name: _c.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_aae.Style, _ecg)
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_fdc *CT_GroupShape) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_fdc.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_fdc.GrpSpPr = _f.NewCT_GroupShapeProperties()
_aegf:
	for {
		_caf, _bfg := d.Token()
		if _bfg != nil {
			return _bfg
		}
		switch _egg := _caf.(type) {
		case _c.StartElement:
			switch _egg.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}:
				if _gdc := d.DecodeElement(_fdc.NvGrpSpPr, &_egg); _gdc != nil {
					return _gdc
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}:
				if _gdg := d.DecodeElement(_fdc.GrpSpPr, &_egg); _gdg != nil {
					return _gdg
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_afc := NewCT_GroupShapeChoice()
				if _gfg := d.DecodeElement(&_afc.Sp, &_egg); _gfg != nil {
					return _gfg
				}
				_fdc.Choice = append(_fdc.Choice, _afc)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_gfc := NewCT_GroupShapeChoice()
				if _aee := d.DecodeElement(&_gfc.GrpSp, &_egg); _aee != nil {
					return _aee
				}
				_fdc.Choice = append(_fdc.Choice, _gfc)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_ace := NewCT_GroupShapeChoice()
				if _ecb := d.DecodeElement(&_ace.GraphicFrame, &_egg); _ecb != nil {
					return _ecb
				}
				_fdc.Choice = append(_fdc.Choice, _ace)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_bbd := NewCT_GroupShapeChoice()
				if _fbf := d.DecodeElement(&_bbd.CxnSp, &_egg); _fbf != nil {
					return _fbf
				}
				_fdc.Choice = append(_fdc.Choice, _bbd)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_fdcf := NewCT_GroupShapeChoice()
				if _geb := d.DecodeElement(&_fdcf.Pic, &_egg); _geb != nil {
					return _geb
				}
				_fdc.Choice = append(_fdc.Choice, _fdcf)
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047r\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065 \u0025\u0076", _egg.Name)
				if _ffg := d.Skip(); _ffg != nil {
					return _ffg
				}
			}
		case _c.EndElement:
			break _aegf
		case _c.CharData:
		}
	}
	return nil
}
func NewCT_GraphicFrameNonVisual() *CT_GraphicFrameNonVisual {
	_dcf := &CT_GraphicFrameNonVisual{}
	_dcf.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_dcf.CNvGraphicFramePr = _f.NewCT_NonVisualGraphicFrameProperties()
	return _dcf
}

type CT_GroupShapeChoice struct {
	Sp           []*CT_Shape
	GrpSp        []*CT_GroupShape
	GraphicFrame []*CT_GraphicFrame
	CxnSp        []*CT_Connector
	Pic          []*CT_Picture
}

// ValidateWithPath validates the EG_ObjectChoicesChoice and its children, prefixing error messages with path
func (_cbc *EG_ObjectChoicesChoice) ValidateWithPath(path string) error {
	if _cbc.Sp != nil {
		if _fee := _cbc.Sp.ValidateWithPath(path + "\u002f\u0053\u0070"); _fee != nil {
			return _fee
		}
	}
	if _cbc.GrpSp != nil {
		if _abaf := _cbc.GrpSp.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070"); _abaf != nil {
			return _abaf
		}
	}
	if _cbc.GraphicFrame != nil {
		if _gdec := _cbc.GraphicFrame.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"); _gdec != nil {
			return _gdec
		}
	}
	if _cbc.CxnSp != nil {
		if _egca := _cbc.CxnSp.ValidateWithPath(path + "\u002f\u0043\u0078\u006e\u0053\u0070"); _egca != nil {
			return _egca
		}
	}
	if _cbc.Pic != nil {
		if _eccea := _cbc.Pic.ValidateWithPath(path + "\u002f\u0050\u0069\u0063"); _eccea != nil {
			return _eccea
		}
	}
	return nil
}
func _afcb(_gfce bool) uint8 {
	if _gfce {
		return 1
	}
	return 0
}
func (_fef *CT_GroupShapeNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_gga := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_fef.CNvPr, _gga)
	_gbge := _c.StartElement{Name: _c.Name{Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_fef.CNvGrpSpPr, _gbge)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the EG_Anchor and its children, prefixing error messages with path
func (_cebcg *EG_Anchor) ValidateWithPath(path string) error {
	if _cebcg.RelSizeAnchor != nil {
		if _bga := _cebcg.RelSizeAnchor.ValidateWithPath(path + "\u002f\u0052\u0065\u006c\u0053\u0069\u007a\u0065\u0041n\u0063\u0068\u006f\u0072"); _bga != nil {
			return _bga
		}
	}
	if _cebcg.AbsSizeAnchor != nil {
		if _baga := _cebcg.AbsSizeAnchor.ValidateWithPath(path + "\u002f\u0041\u0062\u0073\u0053\u0069\u007a\u0065\u0041n\u0063\u0068\u006f\u0072"); _baga != nil {
			return _baga
		}
	}
	return nil
}

// Validate validates the CT_Connector and its children
func (_dec *CT_Connector) Validate() error {
	return _dec.ValidateWithPath("\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072")
}
func (_bab *CT_Picture) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_bab.NvPicPr = NewCT_PictureNonVisual()
	_bab.BlipFill = _f.NewCT_BlipFillProperties()
	_bab.SpPr = _f.NewCT_ShapeProperties()
	for _, _eab := range start.Attr {
		if _eab.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_becc, _acg := _eab.Value, error(nil)
			if _acg != nil {
				return _acg
			}
			_bab.MacroAttr = &_becc
			continue
		}
		if _eab.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_ebbg, _abf := _d.ParseBool(_eab.Value)
			if _abf != nil {
				return _abf
			}
			_bab.FPublishedAttr = &_ebbg
			continue
		}
	}
_afe:
	for {
		_gba, _ecag := d.Token()
		if _ecag != nil {
			return _ecag
		}
		switch _fbb := _gba.(type) {
		case _c.StartElement:
			switch _fbb.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _agb := d.DecodeElement(_bab.NvPicPr, &_fbb); _agb != nil {
					return _agb
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _dae := d.DecodeElement(_bab.BlipFill, &_fbb); _dae != nil {
					return _dae
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _gdbc := d.DecodeElement(_bab.SpPr, &_fbb); _gdbc != nil {
					return _gdbc
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_bab.Style = _f.NewCT_ShapeStyle()
				if _aed := d.DecodeElement(_bab.Style, &_fbb); _aed != nil {
					return _aed
				}
			default:
				_g.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076", _fbb.Name)
				if _cae := d.Skip(); _cae != nil {
					return _cae
				}
			}
		case _c.EndElement:
			break _afe
		case _c.CharData:
		}
	}
	return nil
}
func NewEG_ObjectChoicesChoice() *EG_ObjectChoicesChoice {
	_edcc := &EG_ObjectChoicesChoice{}
	return _edcc
}
func NewCT_AbsSizeAnchor() *CT_AbsSizeAnchor {
	_ff := &CT_AbsSizeAnchor{}
	_ff.From = NewCT_Marker()
	_ff.Ext = _f.NewCT_PositiveSize2D()
	return _ff
}

type CT_RelSizeAnchor struct {
	From   *CT_Marker
	To     *CT_Marker
	Choice *EG_ObjectChoicesChoice
}
type CT_GroupShape struct {
	NvGrpSpPr *CT_GroupShapeNonVisual
	GrpSpPr   *_f.CT_GroupShapeProperties
	Choice    []*CT_GroupShapeChoice
}

// Validate validates the EG_Anchor and its children
func (_aeb *EG_Anchor) Validate() error {
	return _aeb.ValidateWithPath("\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr")
}
func (_feg *CT_GraphicFrame) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_feg.NvGraphicFramePr = NewCT_GraphicFrameNonVisual()
	_feg.Xfrm = _f.NewCT_Transform2D()
	_feg.Graphic = _f.NewGraphic()
	for _, _fcf := range start.Attr {
		if _fcf.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_cbf, _dfd := _fcf.Value, error(nil)
			if _dfd != nil {
				return _dfd
			}
			_feg.MacroAttr = &_cbf
			continue
		}
		if _fcf.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_dba, _cbb := _d.ParseBool(_fcf.Value)
			if _cbb != nil {
				return _cbb
			}
			_feg.FPublishedAttr = &_dba
			continue
		}
	}
_gcb:
	for {
		_cecd, _eg := d.Token()
		if _eg != nil {
			return _eg
		}
		switch _dbd := _cecd.(type) {
		case _c.StartElement:
			switch _dbd.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}:
				if _cdf := d.DecodeElement(_feg.NvGraphicFramePr, &_dbd); _cdf != nil {
					return _cdf
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0078\u0066\u0072\u006d"}:
				if _gce := d.DecodeElement(_feg.Xfrm, &_dbd); _gce != nil {
					return _gce
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}:
				if _abg := d.DecodeElement(_feg.Graphic, &_dbd); _abg != nil {
					return _abg
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0020\u0025\u0076", _dbd.Name)
				if _dfdg := d.Skip(); _dfdg != nil {
					return _dfdg
				}
			}
		case _c.EndElement:
			break _gcb
		case _c.CharData:
		}
	}
	return nil
}
func (_bfbd *EG_ObjectChoices) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_gcba:
	for {
		_gab, _dfcd := d.Token()
		if _dfcd != nil {
			return _dfcd
		}
		switch _fgae := _gab.(type) {
		case _c.StartElement:
			switch _fgae.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_bfbd.Choice = NewEG_ObjectChoicesChoice()
				if _fde := d.DecodeElement(&_bfbd.Choice.Sp, &_fgae); _fde != nil {
					return _fde
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_bfbd.Choice = NewEG_ObjectChoicesChoice()
				if _fbd := d.DecodeElement(&_bfbd.Choice.GrpSp, &_fgae); _fbd != nil {
					return _fbd
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_bfbd.Choice = NewEG_ObjectChoicesChoice()
				if _age := d.DecodeElement(&_bfbd.Choice.GraphicFrame, &_fgae); _age != nil {
					return _age
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_bfbd.Choice = NewEG_ObjectChoicesChoice()
				if _acd := d.DecodeElement(&_bfbd.Choice.CxnSp, &_fgae); _acd != nil {
					return _acd
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_bfbd.Choice = NewEG_ObjectChoicesChoice()
				if _abbg := d.DecodeElement(&_bfbd.Choice.Pic, &_fgae); _abbg != nil {
					return _abbg
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u004f\u0062\u006a\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0020\u0025v", _fgae.Name)
				if _gdga := d.Skip(); _gdga != nil {
					return _gdga
				}
			}
		case _c.EndElement:
			break _gcba
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_adf *CT_Picture) ValidateWithPath(path string) error {
	if _dbe := _adf.NvPicPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072"); _dbe != nil {
		return _dbe
	}
	if _afcc := _adf.BlipFill.ValidateWithPath(path + "\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl"); _afcc != nil {
		return _afcc
	}
	if _ceb := _adf.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _ceb != nil {
		return _ceb
	}
	if _adf.Style != nil {
		if _beg := _adf.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _beg != nil {
			return _beg
		}
	}
	return nil
}
func NewCT_Drawing() *CT_Drawing { _fbc := &CT_Drawing{}; return _fbc }
func NewCT_Shape() *CT_Shape {
	_affc := &CT_Shape{}
	_affc.NvSpPr = NewCT_ShapeNonVisual()
	_affc.SpPr = _f.NewCT_ShapeProperties()
	return _affc
}

// Validate validates the CT_Drawing and its children
func (_dgd *CT_Drawing) Validate() error {
	return _dgd.ValidateWithPath("\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067")
}
func NewEG_Anchor() *EG_Anchor { _edcg := &EG_Anchor{}; return _edcg }

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_cebc *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if _adeb := _cebc.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _adeb != nil {
		return _adeb
	}
	if _cddg := _cebc.CNvPicPr.ValidateWithPath(path + "\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r"); _cddg != nil {
		return _cddg
	}
	return nil
}
func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	_cda := &CT_PictureNonVisual{}
	_cda.CNvPr = _f.NewCT_NonVisualDrawingProps()
	_cda.CNvPicPr = _f.NewCT_NonVisualPictureProperties()
	return _cda
}

type CT_GraphicFrameNonVisual struct {
	CNvPr             *_f.CT_NonVisualDrawingProps
	CNvGraphicFramePr *_f.CT_NonVisualGraphicFrameProperties
}

// ValidateWithPath validates the CT_GraphicFrameNonVisual and its children, prefixing error messages with path
func (_eae *CT_GraphicFrameNonVisual) ValidateWithPath(path string) error {
	if _efd := _eae.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _efd != nil {
		return _efd
	}
	if _daag := _eae.CNvGraphicFramePr.ValidateWithPath(path + "\u002fC\u004ev\u0047\u0072\u0061\u0070\u0068i\u0063\u0046r\u0061\u006d\u0065\u0050\u0072"); _daag != nil {
		return _daag
	}
	return nil
}
func NewCT_GraphicFrame() *CT_GraphicFrame {
	_ece := &CT_GraphicFrame{}
	_ece.NvGraphicFramePr = NewCT_GraphicFrameNonVisual()
	_ece.Xfrm = _f.NewCT_Transform2D()
	_ece.Graphic = _f.NewGraphic()
	return _ece
}
func (_cff *EG_ObjectChoices) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _cff.Choice != nil {
		_cff.Choice.MarshalXML(e, _c.StartElement{})
	}
	return nil
}

// Validate validates the CT_GraphicFrameNonVisual and its children
func (_ef *CT_GraphicFrameNonVisual) Validate() error {
	return _ef.ValidateWithPath("\u0043T\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061m\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}
func (_ffb *EG_ObjectChoicesChoice) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_bfae:
	for {
		_decec, _gbb := d.Token()
		if _gbb != nil {
			return _gbb
		}
		switch _cegf := _decec.(type) {
		case _c.StartElement:
			switch _cegf.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_ffb.Sp = NewCT_Shape()
				if _dcfb := d.DecodeElement(_ffb.Sp, &_cegf); _dcfb != nil {
					return _dcfb
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_ffb.GrpSp = NewCT_GroupShape()
				if _afdb := d.DecodeElement(_ffb.GrpSp, &_cegf); _afdb != nil {
					return _afdb
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_ffb.GraphicFrame = NewCT_GraphicFrame()
				if _babe := d.DecodeElement(_ffb.GraphicFrame, &_cegf); _babe != nil {
					return _babe
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_ffb.CxnSp = NewCT_Connector()
				if _dff := d.DecodeElement(_ffb.CxnSp, &_cegf); _dff != nil {
					return _dff
				}
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_ffb.Pic = NewCT_Picture()
				if _gfcg := d.DecodeElement(_ffb.Pic, &_cegf); _gfcg != nil {
					return _gfcg
				}
			default:
				_g.Log("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045G\u005f\u004f\u0062\u006a\u0065c\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _cegf.Name)
				if _gacg := d.Skip(); _gacg != nil {
					return _gacg
				}
			}
		case _c.EndElement:
			break _bfae
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupShape and its children
func (_bffd *CT_GroupShape) Validate() error {
	return _bffd.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065")
}

// Validate validates the EG_ObjectChoices and its children
func (_dca *EG_ObjectChoices) Validate() error {
	return _dca.ValidateWithPath("\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073")
}
func (_efe *CT_GroupShapeChoice) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_dce:
	for {
		_gfb, _ag := d.Token()
		if _ag != nil {
			return _ag
		}
		switch _aff := _gfb.(type) {
		case _c.StartElement:
			switch _aff.Name {
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_dbaa := NewCT_Shape()
				if _ecc := d.DecodeElement(_dbaa, &_aff); _ecc != nil {
					return _ecc
				}
				_efe.Sp = append(_efe.Sp, _dbaa)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_dece := NewCT_GroupShape()
				if _dcba := d.DecodeElement(_dece, &_aff); _dcba != nil {
					return _dcba
				}
				_efe.GrpSp = append(_efe.GrpSp, _dece)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_dcbd := NewCT_GraphicFrame()
				if _gec := d.DecodeElement(_dcbd, &_aff); _gec != nil {
					return _gec
				}
				_efe.GraphicFrame = append(_efe.GraphicFrame, _dcbd)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_gfba := NewCT_Connector()
				if _gaee := d.DecodeElement(_gfba, &_aff); _gaee != nil {
					return _gaee
				}
				_efe.CxnSp = append(_efe.CxnSp, _gfba)
			case _c.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_gcf := NewCT_Picture()
				if _cfd := d.DecodeElement(_gcf, &_aff); _cfd != nil {
					return _cfd
				}
				_efe.Pic = append(_efe.Pic, _gcf)
			default:
				_g.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068ap\u0065\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _aff.Name)
				if _cgac := d.Skip(); _cgac != nil {
					return _cgac
				}
			}
		case _c.EndElement:
			break _dce
		case _c.CharData:
		}
	}
	return nil
}

type CT_Marker struct {
	X float64
	Y float64
}

// Validate validates the EG_ObjectChoicesChoice and its children
func (_cgc *EG_ObjectChoicesChoice) Validate() error {
	return _cgc.ValidateWithPath("\u0045\u0047\u005f\u004fbj\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069c\u0065")
}

// Validate validates the CT_RelSizeAnchor and its children
func (_caca *CT_RelSizeAnchor) Validate() error {
	return _caca.ValidateWithPath("\u0043\u0054_\u0052\u0065\u006cS\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072")
}

// ValidateWithPath validates the CT_GroupShapeNonVisual and its children, prefixing error messages with path
func (_gad *CT_GroupShapeNonVisual) ValidateWithPath(path string) error {
	if _fbeb := _gad.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _fbeb != nil {
		return _fbeb
	}
	if _fgb := _gad.CNvGrpSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _fgb != nil {
		return _fgb
	}
	return nil
}

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (_cca *CT_Shape) ValidateWithPath(path string) error {
	if _bcgb := _cca.NvSpPr.ValidateWithPath(path + "\u002fN\u0076\u0053\u0070\u0050\u0072"); _bcgb != nil {
		return _bcgb
	}
	if _ccbc := _cca.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _ccbc != nil {
		return _ccbc
	}
	if _cca.Style != nil {
		if _gbgc := _cca.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _gbgc != nil {
			return _gbgc
		}
	}
	if _cca.TxBody != nil {
		if _dfcb := _cca.TxBody.ValidateWithPath(path + "\u002fT\u0078\u0042\u006f\u0064\u0079"); _dfcb != nil {
			return _dfcb
		}
	}
	return nil
}
func init() {
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c", NewCT_ShapeNonVisual)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065", NewCT_Shape)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_ConnectorNonVisual)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072", NewCT_Connector)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_PictureNonVisual)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043T\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061m\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_GraphicFrameNonVisual)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043T\u005fG\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065", NewCT_GraphicFrame)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c", NewCT_GroupShapeNonVisual)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065", NewCT_GroupShape)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r", NewCT_Marker)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054_\u0052\u0065\u006cS\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_RelSizeAnchor)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054_\u0041\u0062\u0073S\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_AbsSizeAnchor)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067", NewCT_Drawing)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073", NewEG_ObjectChoices)
	_g.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr", NewEG_Anchor)
}
