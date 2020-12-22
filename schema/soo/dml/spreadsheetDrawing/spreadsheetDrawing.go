package spreadsheetDrawing

import (
	_c "encoding/xml"
	_cf "fmt"
	_ca "strconv"

	_f "gitee.com/yu_sheng/gooffice"
	_b "gitee.com/yu_sheng/gooffice/schema/soo/dml"
)

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (_edbe *CT_Shape) ValidateWithPath(path string) error {
	if _agg := _edbe.NvSpPr.ValidateWithPath(path + "\u002fN\u0076\u0053\u0070\u0050\u0072"); _agg != nil {
		return _agg
	}
	if _ecf := _edbe.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _ecf != nil {
		return _ecf
	}
	if _edbe.Style != nil {
		if _bgdd := _edbe.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _bgdd != nil {
			return _bgdd
		}
	}
	if _edbe.TxBody != nil {
		if _egce := _edbe.TxBody.ValidateWithPath(path + "\u002fT\u0078\u0042\u006f\u0064\u0079"); _egce != nil {
			return _egce
		}
	}
	return nil
}
func NewCT_GroupShapeNonVisual() *CT_GroupShapeNonVisual {
	_afgb := &CT_GroupShapeNonVisual{}
	_afgb.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_afgb.CNvGrpSpPr = _b.NewCT_NonVisualGroupDrawingShapeProps()
	return _afgb
}
func NewFrom() *From { _egf := &From{}; _egf.CT_Marker = *NewCT_Marker(); return _egf }

type CT_ConnectorNonVisual struct {
	CNvPr      *_b.CT_NonVisualDrawingProps
	CNvCxnSpPr *_b.CT_NonVisualConnectorProperties
}

func (_bcdc *CT_ShapeNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_bcdc.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_bcdc.CNvSpPr = _b.NewCT_NonVisualDrawingShapeProps()
_ffb:
	for {
		_ebbbb, _bgfd := d.Token()
		if _bgfd != nil {
			return _bgfd
		}
		switch _deg := _ebbbb.(type) {
		case _c.StartElement:
			switch _deg.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _cfbe := d.DecodeElement(_bcdc.CNvPr, &_deg); _cfbe != nil {
					return _cfbe
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}:
				if _fcgg := d.DecodeElement(_bcdc.CNvSpPr, &_deg); _fcgg != nil {
					return _fcgg
				}
			default:
				_f.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _deg.Name)
				if _adf := d.Skip(); _adf != nil {
					return _adf
				}
			}
		case _c.EndElement:
			break _ffb
		case _c.CharData:
		}
	}
	return nil
}
func (_ecd *CT_GraphicalObjectFrame) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ecd.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	_ecd.Xfrm = _b.NewCT_Transform2D()
	_ecd.Graphic = _b.NewGraphic()
	for _, _cgb := range start.Attr {
		if _cgb.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_dfb, _edg := _cgb.Value, error(nil)
			if _edg != nil {
				return _edg
			}
			_ecd.MacroAttr = &_dfb
			continue
		}
		if _cgb.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_cfc, _ggb := _ca.ParseBool(_cgb.Value)
			if _ggb != nil {
				return _ggb
			}
			_ecd.FPublishedAttr = &_cfc
			continue
		}
	}
_ded:
	for {
		_bdb, _eed := d.Token()
		if _eed != nil {
			return _eed
		}
		switch _dfg := _bdb.(type) {
		case _c.StartElement:
			switch _dfg.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}:
				if _ffad := d.DecodeElement(_ecd.NvGraphicFramePr, &_dfg); _ffad != nil {
					return _ffad
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0078\u0066\u0072\u006d"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0078\u0066\u0072\u006d"}:
				if _fgc := d.DecodeElement(_ecd.Xfrm, &_dfg); _fgc != nil {
					return _fgc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}:
				if _ebbd := d.DecodeElement(_ecd.Graphic, &_dfg); _ebbd != nil {
					return _ebbd
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn\u0020\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0061\u006cO\u0062\u006a\u0065\u0063\u0074\u0046r\u0061\u006d\u0065 \u0025\u0076", _dfg.Name)
				if _be := d.Skip(); _be != nil {
					return _be
				}
			}
		case _c.EndElement:
			break _ded
		case _c.CharData:
		}
	}
	return nil
}
func (_cbd *CT_PictureNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_cbd.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_cbd.CNvPicPr = _b.NewCT_NonVisualPictureProperties()
_acea:
	for {
		_dgdc, _dcff := d.Token()
		if _dcff != nil {
			return _dcff
		}
		switch _dgdag := _dgdc.(type) {
		case _c.StartElement:
			switch _dgdag.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _acae := d.DecodeElement(_cbd.CNvPr, &_dgdag); _acae != nil {
					return _acae
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:
				if _fba := d.DecodeElement(_cbd.CNvPicPr, &_dgdag); _fba != nil {
					return _fba
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _dgdag.Name)
				if _bae := d.Skip(); _bae != nil {
					return _bae
				}
			}
		case _c.EndElement:
			break _acea
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_GraphicalObjectFrameNonVisual and its children, prefixing error messages with path
func (_ffce *CT_GraphicalObjectFrameNonVisual) ValidateWithPath(path string) error {
	if _efd := _ffce.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _efd != nil {
		return _efd
	}
	if _bgf := _ffce.CNvGraphicFramePr.ValidateWithPath(path + "\u002fC\u004ev\u0047\u0072\u0061\u0070\u0068i\u0063\u0046r\u0061\u006d\u0065\u0050\u0072"); _bgf != nil {
		return _bgf
	}
	return nil
}

// Validate validates the CT_Marker and its children
func (_acf *CT_Marker) Validate() error {
	return _acf.ValidateWithPath("\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r")
}

type CT_Drawing struct{ EG_Anchor []*EG_Anchor }

func (_aab *CT_Drawing) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	if _aab.EG_Anchor != nil {
		for _, _ffa := range _aab.EG_Anchor {
			_ffa.MarshalXML(e, _c.StartElement{})
		}
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type CT_OneCellAnchor struct {
	From       *CT_Marker
	Ext        *_b.CT_PositiveSize2D
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

// ValidateWithPath validates the CT_ConnectorNonVisual and its children, prefixing error messages with path
func (_cca *CT_ConnectorNonVisual) ValidateWithPath(path string) error {
	if _afe := _cca.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _afe != nil {
		return _afe
	}
	if _df := _cca.CNvCxnSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _df != nil {
		return _df
	}
	return nil
}
func NewCT_Picture() *CT_Picture {
	_dbg := &CT_Picture{}
	_dbg.NvPicPr = NewCT_PictureNonVisual()
	_dbg.BlipFill = _b.NewCT_BlipFillProperties()
	_dbg.SpPr = _b.NewCT_ShapeProperties()
	return _dbg
}
func (_gbe *EG_ObjectChoicesChoice) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_deaa:
	for {
		_gagg, _abf := d.Token()
		if _abf != nil {
			return _abf
		}
		switch _fda := _gagg.(type) {
		case _c.StartElement:
			switch _fda.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_gbe.Sp = NewCT_Shape()
				if _acad := d.DecodeElement(_gbe.Sp, &_fda); _acad != nil {
					return _acad
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_gbe.GrpSp = NewCT_GroupShape()
				if _eedg := d.DecodeElement(_gbe.GrpSp, &_fda); _eedg != nil {
					return _eedg
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_gbe.GraphicFrame = NewCT_GraphicalObjectFrame()
				if _ced := d.DecodeElement(_gbe.GraphicFrame, &_fda); _ced != nil {
					return _ced
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_gbe.CxnSp = NewCT_Connector()
				if _cbg := d.DecodeElement(_gbe.CxnSp, &_fda); _cbg != nil {
					return _cbg
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_gbe.Pic = NewCT_Picture()
				if _dcbg := d.DecodeElement(_gbe.Pic, &_fda); _dcbg != nil {
					return _dcbg
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_gbe.ContentPart = NewCT_Rel()
				if _bbag := d.DecodeElement(_gbe.ContentPart, &_fda); _bbag != nil {
					return _bbag
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045G\u005f\u004f\u0062\u006a\u0065c\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _fda.Name)
				if _gagf := d.Skip(); _gagf != nil {
					return _gagf
				}
			}
		case _c.EndElement:
			break _deaa
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OneCellAnchor and its children
func (_fecc *CT_OneCellAnchor) Validate() error {
	return _fecc.ValidateWithPath("\u0043\u0054_\u004f\u006e\u0065C\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072")
}

// ValidateWithPath validates the CT_AbsoluteAnchor and its children, prefixing error messages with path
func (_fe *CT_AbsoluteAnchor) ValidateWithPath(path string) error {
	if _ff := _fe.Pos.ValidateWithPath(path + "\u002f\u0050\u006f\u0073"); _ff != nil {
		return _ff
	}
	if _ce := _fe.Ext.ValidateWithPath(path + "\u002f\u0045\u0078\u0074"); _ce != nil {
		return _ce
	}
	if _fe.Choice != nil {
		if _da := _fe.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _da != nil {
			return _da
		}
	}
	if _ebb := _fe.ClientData.ValidateWithPath(path + "/\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"); _ebb != nil {
		return _ebb
	}
	return nil
}
func NewCT_GraphicalObjectFrame() *CT_GraphicalObjectFrame {
	_gebg := &CT_GraphicalObjectFrame{}
	_gebg.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	_gebg.Xfrm = _b.NewCT_Transform2D()
	_gebg.Graphic = _b.NewGraphic()
	return _gebg
}

type ST_EditAs byte

func (_dae *CT_GroupShape) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_eac := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u006e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_dae.NvGrpSpPr, _eac)
	_bed := _c.StartElement{Name: _c.Name{Local: "x\u0064\u0072\u003a\u0067\u0072\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_dae.GrpSpPr, _bed)
	if _dae.Choice != nil {
		for _, _cec := range _dae.Choice {
			_cec.MarshalXML(e, _c.StartElement{})
		}
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_aad *From) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_aad.CT_Marker = *NewCT_Marker()
_ggff:
	for {
		_dga, _aga := d.Token()
		if _aga != nil {
			return _aga
		}
		switch _egd := _dga.(type) {
		case _c.StartElement:
			switch _egd.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}:
				if _fgeb := d.DecodeElement(&_aad.Col, &_egd); _fgeb != nil {
					return _fgeb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}:
				_eec, _bbad := d.Token()
				if _bbad != nil {
					return _bbad
				}
				switch _eadg := _eec.(type) {
				case _c.CharData:
					_eegbc := string(_eadg)
					_edc, _daf := _b.ParseUnionST_Coordinate(_eegbc)
					if _daf != nil {
						return nil
					}
					_aad.ColOff = _edc
					d.Skip()
				case _c.EndElement:
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}:
				if _aaca := d.DecodeElement(&_aad.Row, &_egd); _aaca != nil {
					return _aaca
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}:
				_gbee, _cdg := d.Token()
				if _cdg != nil {
					return _cdg
				}
				switch _bcde := _gbee.(type) {
				case _c.CharData:
					_cgeb := string(_bcde)
					_ffef, _bac := _b.ParseUnionST_Coordinate(_cgeb)
					if _bac != nil {
						return nil
					}
					_aad.RowOff = _ffef
					d.Skip()
				case _c.EndElement:
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0046\u0072o\u006d\u0020\u0025\u0076", _egd.Name)
				if _gced := d.Skip(); _gced != nil {
					return _gced
				}
			}
		case _c.EndElement:
			break _ggff
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ObjectChoices and its children
func (_bbgg *EG_ObjectChoices) Validate() error {
	return _bbgg.ValidateWithPath("\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073")
}

type CT_PictureNonVisual struct {
	CNvPr    *_b.CT_NonVisualDrawingProps
	CNvPicPr *_b.CT_NonVisualPictureProperties
}

// ValidateWithPath validates the CT_GraphicalObjectFrame and its children, prefixing error messages with path
func (_ddd *CT_GraphicalObjectFrame) ValidateWithPath(path string) error {
	if _cfa := _ddd.NvGraphicFramePr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"); _cfa != nil {
		return _cfa
	}
	if _fcac := _ddd.Xfrm.ValidateWithPath(path + "\u002f\u0058\u0066r\u006d"); _fcac != nil {
		return _fcac
	}
	if _eca := _ddd.Graphic.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063"); _eca != nil {
		return _eca
	}
	return nil
}
func (_ecc *CT_GraphicalObjectFrameNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ecc.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_ecc.CNvGraphicFramePr = _b.NewCT_NonVisualGraphicFrameProperties()
_cd:
	for {
		_dfe, _afga := d.Token()
		if _afga != nil {
			return _afga
		}
		switch _add := _dfe.(type) {
		case _c.StartElement:
			switch _add.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _bdba := d.DecodeElement(_ecc.CNvPr, &_add); _bdba != nil {
					return _bdba
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}:
				if _agc := d.DecodeElement(_ecc.CNvGraphicFramePr, &_add); _agc != nil {
					return _agc
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u0061p\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006a\u0065\u0063\u0074\u0046\u0072\u0061\u006de\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061l\u0020\u0025\u0076", _add.Name)
				if _gba := d.Skip(); _gba != nil {
					return _gba
				}
			}
		case _c.EndElement:
			break _cd
		case _c.CharData:
		}
	}
	return nil
}
func (_cab *CT_GraphicalObjectFrame) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _cab.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _cf.Sprintf("\u0025\u0076", *_cab.MacroAttr)})
	}
	if _cab.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _cf.Sprintf("\u0025\u0064", _fgcd(*_cab.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_baf := _c.StartElement{Name: _c.Name{Local: "x\u0064r\u003a\u006e\u0076\u0047\u0072\u0061\u0070\u0068i\u0063\u0046\u0072\u0061me\u0050\u0072"}}
	e.EncodeElement(_cab.NvGraphicFramePr, _baf)
	_ccb := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0078\u0066\u0072\u006d"}}
	e.EncodeElement(_cab.Xfrm, _ccb)
	_bgc := _c.StartElement{Name: _c.Name{Local: "\u0061:\u0067\u0072\u0061\u0070\u0068\u0069c"}}
	e.EncodeElement(_cab.Graphic, _bgc)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_gb *CT_Connector) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_gb.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_gb.SpPr = _b.NewCT_ShapeProperties()
	for _, _afg := range start.Attr {
		if _afg.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_baa, _dd := _afg.Value, error(nil)
			if _dd != nil {
				return _dd
			}
			_gb.MacroAttr = &_baa
			continue
		}
		if _afg.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_bg, _ed := _ca.ParseBool(_afg.Value)
			if _ed != nil {
				return _ed
			}
			_gb.FPublishedAttr = &_bg
			continue
		}
	}
_eaf:
	for {
		_dg, _fac := d.Token()
		if _fac != nil {
			return _fac
		}
		switch _dc := _dg.(type) {
		case _c.StartElement:
			switch _dc.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}:
				if _dde := d.DecodeElement(_gb.NvCxnSpPr, &_dc); _dde != nil {
					return _dde
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _bf := d.DecodeElement(_gb.SpPr, &_dc); _bf != nil {
					return _bf
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_gb.Style = _b.NewCT_ShapeStyle()
				if _faf := d.DecodeElement(_gb.Style, &_dc); _faf != nil {
					return _faf
				}
			default:
				_f.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_C\u006f\u006en\u0065\u0063\u0074\u006f\u0072\u0020\u0025\u0076", _dc.Name)
				if _afd := d.Skip(); _afd != nil {
					return _afd
				}
			}
		case _c.EndElement:
			break _eaf
		case _c.CharData:
		}
	}
	return nil
}
func (_gdgb *CT_ShapeNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_fdff := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_gdgb.CNvPr, _fdff)
	_dcb := _c.StartElement{Name: _c.Name{Local: "x\u0064\u0072\u003a\u0063\u004e\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_gdgb.CNvSpPr, _dcb)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_fdbf *CT_Rel) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0072\u003a\u0069\u0064"}, Value: _cf.Sprintf("\u0025\u0076", _fdbf.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (_fef *CT_Drawing) ValidateWithPath(path string) error {
	for _bce, _ada := range _fef.EG_Anchor {
		if _ffff := _ada.ValidateWithPath(_cf.Sprintf("\u0025\u0073/\u0045\u0047\u005fA\u006e\u0063\u0068\u006f\u0072\u005b\u0025\u0064\u005d", path, _bce)); _ffff != nil {
			return _ffff
		}
	}
	return nil
}
func NewCT_GroupShapeChoice() *CT_GroupShapeChoice { _ggbf := &CT_GroupShapeChoice{}; return _ggbf }
func (_fdcg *CT_Marker) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_fdcg.Col = 0
	_fdcg.Row = 0
_ddc:
	for {
		_dab, _bad := d.Token()
		if _bad != nil {
			return _bad
		}
		switch _fded := _dab.(type) {
		case _c.StartElement:
			switch _fded.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}:
				if _gab := d.DecodeElement(&_fdcg.Col, &_fded); _gab != nil {
					return _gab
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}:
				_bgd, _ecgd := d.Token()
				if _ecgd != nil {
					return _ecgd
				}
				switch _caga := _bgd.(type) {
				case _c.CharData:
					_aceb := string(_caga)
					_dca, _dgc := _b.ParseUnionST_Coordinate(_aceb)
					if _dgc != nil {
						return nil
					}
					_fdcg.ColOff = _dca
					d.Skip()
				case _c.EndElement:
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}:
				if _gaf := d.DecodeElement(&_fdcg.Row, &_fded); _gaf != nil {
					return _gaf
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}:
				_def, _dgb := d.Token()
				if _dgb != nil {
					return _dgb
				}
				switch _bda := _def.(type) {
				case _c.CharData:
					_cbbg := string(_bda)
					_dag, _bga := _b.ParseUnionST_Coordinate(_cbbg)
					if _bga != nil {
						return nil
					}
					_fdcg.RowOff = _dag
					d.Skip()
				case _c.EndElement:
				}
			default:
				_f.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004d\u0061\u0072k\u0065\u0072 \u0025\u0076", _fded.Name)
				if _adea := d.Skip(); _adea != nil {
					return _adea
				}
			}
		case _c.EndElement:
			break _ddc
		case _c.CharData:
		}
	}
	return nil
}
func (_dfda *ST_EditAs) UnmarshalXMLAttr(attr _c.Attr) error {
	switch attr.Value {
	case "":
		*_dfda = 0
	case "\u0074w\u006f\u0043\u0065\u006c\u006c":
		*_dfda = 1
	case "\u006fn\u0065\u0043\u0065\u006c\u006c":
		*_dfda = 2
	case "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065":
		*_dfda = 3
	}
	return nil
}

// Validate validates the CT_GroupShapeNonVisual and its children
func (_edb *CT_GroupShapeNonVisual) Validate() error {
	return _edb.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c")
}
func NewCT_TwoCellAnchor() *CT_TwoCellAnchor {
	_agdc := &CT_TwoCellAnchor{}
	_agdc.From = NewCT_Marker()
	_agdc.To = NewCT_Marker()
	_agdc.ClientData = NewCT_AnchorClientData()
	return _agdc
}
func NewWsDr() *WsDr { _eeeeb := &WsDr{}; _eeeeb.CT_Drawing = *NewCT_Drawing(); return _eeeeb }
func (_ebab *CT_Picture) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ebab.NvPicPr = NewCT_PictureNonVisual()
	_ebab.BlipFill = _b.NewCT_BlipFillProperties()
	_ebab.SpPr = _b.NewCT_ShapeProperties()
	for _, _eddb := range start.Attr {
		if _eddb.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_gcd, _bbc := _eddb.Value, error(nil)
			if _bbc != nil {
				return _bbc
			}
			_ebab.MacroAttr = &_gcd
			continue
		}
		if _eddb.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_cdc, _dabd := _ca.ParseBool(_eddb.Value)
			if _dabd != nil {
				return _dabd
			}
			_ebab.FPublishedAttr = &_cdc
			continue
		}
	}
_dbe:
	for {
		_bdae, _ede := d.Token()
		if _ede != nil {
			return _ede
		}
		switch _cfg := _bdae.(type) {
		case _c.StartElement:
			switch _cfg.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _afaf := d.DecodeElement(_ebab.NvPicPr, &_cfg); _afaf != nil {
					return _afaf
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _face := d.DecodeElement(_ebab.BlipFill, &_cfg); _face != nil {
					return _face
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _fed := d.DecodeElement(_ebab.SpPr, &_cfg); _fed != nil {
					return _fed
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_ebab.Style = _b.NewCT_ShapeStyle()
				if _dcf := d.DecodeElement(_ebab.Style, &_cfg); _dcf != nil {
					return _dcf
				}
			default:
				_f.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076", _cfg.Name)
				if _fdga := d.Skip(); _fdga != nil {
					return _fdga
				}
			}
		case _c.EndElement:
			break _dbe
		case _c.CharData:
		}
	}
	return nil
}
func (_fbf *CT_Connector) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _fbf.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _cf.Sprintf("\u0025\u0076", *_fbf.MacroAttr)})
	}
	if _fbf.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _cf.Sprintf("\u0025\u0064", _fgcd(*_fbf.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_gd := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u006e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_fbf.NvCxnSpPr, _gd)
	_ffc := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_fbf.SpPr, _ffc)
	if _fbf.Style != nil {
		_gca := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0073\u0074\u0079\u006ce"}}
		e.EncodeElement(_fbf.Style, _gca)
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_cfec *CT_Rel) ValidateWithPath(path string) error { return nil }
func (_dbf *EG_ObjectChoicesChoice) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _dbf.Sp != nil {
		_cbc := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070"}}
		e.EncodeElement(_dbf.Sp, _cbc)
	}
	if _dbf.GrpSp != nil {
		_fee := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0067\u0072\u0070\u0053p"}}
		e.EncodeElement(_dbf.GrpSp, _fee)
	}
	if _dbf.GraphicFrame != nil {
		_abb := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064r\u003a\u0067\u0072a\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		e.EncodeElement(_dbf.GraphicFrame, _abb)
	}
	if _dbf.CxnSp != nil {
		_cba := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u0078\u006e\u0053p"}}
		e.EncodeElement(_dbf.CxnSp, _cba)
	}
	if _dbf.Pic != nil {
		_cbac := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0070\u0069\u0063"}}
		e.EncodeElement(_dbf.Pic, _cbac)
	}
	if _dbf.ContentPart != nil {
		_bfc := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072:\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}}
		e.EncodeElement(_dbf.ContentPart, _bfc)
	}
	return nil
}
func (_geb *CT_ConnectorNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_ece := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_geb.CNvPr, _ece)
	_fcd := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u004e\u0076\u0043\u0078n\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_geb.CNvCxnSpPr, _fcd)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the From and its children, prefixing error messages with path
func (_fdbd *From) ValidateWithPath(path string) error {
	if _fbaf := _fdbd.CT_Marker.ValidateWithPath(path); _fbaf != nil {
		return _fbaf
	}
	return nil
}

type CT_GroupShape struct {
	NvGrpSpPr *CT_GroupShapeNonVisual
	GrpSpPr   *_b.CT_GroupShapeProperties
	Choice    []*CT_GroupShapeChoice
}
type CT_ShapeNonVisual struct {
	CNvPr   *_b.CT_NonVisualDrawingProps
	CNvSpPr *_b.CT_NonVisualDrawingShapeProps
}

func (_fff *CT_AnchorClientData) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for _, _aa := range start.Attr {
		if _aa.Name.Local == "\u0066L\u006fc\u006b\u0073\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074" {
			_acd, _cg := _ca.ParseBool(_aa.Value)
			if _cg != nil {
				return _cg
			}
			_fff.FLocksWithSheetAttr = &_acd
			continue
		}
		if _aa.Name.Local == "\u0066\u0050r\u0069\u006e\u0074s\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074" {
			_de, _ad := _ca.ParseBool(_aa.Value)
			if _ad != nil {
				return _ad
			}
			_fff.FPrintsWithSheetAttr = &_de
			continue
		}
	}
	for {
		_cea, _ebg := d.Token()
		if _ebg != nil {
			return _cf.Errorf("\u0070\u0061\u0072s\u0069\u006e\u0067\u0020C\u0054\u005f\u0041\u006e\u0063\u0068\u006fr\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u003a\u0020\u0025\u0073", _ebg)
		}
		if _af, _bab := _cea.(_c.EndElement); _bab && _af.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_AbsoluteAnchor struct {
	Pos        *_b.CT_Point2D
	Ext        *_b.CT_PositiveSize2D
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

func NewCT_AnchorClientData() *CT_AnchorClientData { _feg := &CT_AnchorClientData{}; return _feg }
func NewCT_OneCellAnchor() *CT_OneCellAnchor {
	_ccf := &CT_OneCellAnchor{}
	_ccf.From = NewCT_Marker()
	_ccf.Ext = _b.NewCT_PositiveSize2D()
	_ccf.ClientData = NewCT_AnchorClientData()
	return _ccf
}

type To struct{ CT_Marker }

func (_dfc *CT_TwoCellAnchor) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _dfc.EditAsAttr != ST_EditAsUnset {
		_bdbd, _bcgb := _dfc.EditAsAttr.MarshalXMLAttr(_c.Name{Local: "\u0065\u0064\u0069\u0074\u0041\u0073"})
		if _bcgb != nil {
			return _bcgb
		}
		start.Attr = append(start.Attr, _bdbd)
	}
	e.EncodeToken(start)
	_cce := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_dfc.From, _cce)
	_bgce := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0074\u006f"}}
	e.EncodeElement(_dfc.To, _bgce)
	if _dfc.Choice != nil {
		_dfc.Choice.MarshalXML(e, _c.StartElement{})
	}
	_aef := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006c\u0069\u0065\u006et\u0044\u0061\u0074\u0061"}}
	e.EncodeElement(_dfc.ClientData, _aef)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Drawing and its children
func (_gcaf *CT_Drawing) Validate() error {
	return _gcaf.ValidateWithPath("\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067")
}

// Validate validates the CT_AbsoluteAnchor and its children
func (_ge *CT_AbsoluteAnchor) Validate() error {
	return _ge.ValidateWithPath("\u0043\u0054\u005f\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072")
}
func (_cfb *CT_Drawing) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_eff:
	for {
		_acdc, _dfd := d.Token()
		if _dfd != nil {
			return _dfd
		}
		switch _bcg := _acdc.(type) {
		case _c.StartElement:
			switch _bcg.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_fea := NewEG_Anchor()
				_fea.TwoCellAnchor = NewCT_TwoCellAnchor()
				if _fgf := d.DecodeElement(_fea.TwoCellAnchor, &_bcg); _fgf != nil {
					return _fgf
				}
				_cfb.EG_Anchor = append(_cfb.EG_Anchor, _fea)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_bcf := NewEG_Anchor()
				_bcf.OneCellAnchor = NewCT_OneCellAnchor()
				if _fde := d.DecodeElement(_bcf.OneCellAnchor, &_bcg); _fde != nil {
					return _fde
				}
				_cfb.EG_Anchor = append(_cfb.EG_Anchor, _bcf)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}:
				_ffe := NewEG_Anchor()
				_ffe.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if _eea := d.DecodeElement(_ffe.AbsoluteAnchor, &_bcg); _eea != nil {
					return _eea
				}
				_cfb.EG_Anchor = append(_cfb.EG_Anchor, _ffe)
			default:
				_f.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fD\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0025\u0076", _bcg.Name)
				if _fdcf := d.Skip(); _fdcf != nil {
					return _fdcf
				}
			}
		case _c.EndElement:
			break _eff
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_bgb *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if _dfgb := _bgb.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _dfgb != nil {
		return _dfgb
	}
	if _ffd := _bgb.CNvPicPr.ValidateWithPath(path + "\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r"); _ffd != nil {
		return _ffd
	}
	return nil
}

type EG_ObjectChoices struct{ Choice *EG_ObjectChoicesChoice }

func NewCT_Connector() *CT_Connector {
	_fcg := &CT_Connector{}
	_fcg.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_fcg.SpPr = _b.NewCT_ShapeProperties()
	return _fcg
}

// Validate validates the CT_GroupShape and its children
func (_fag *CT_GroupShape) Validate() error {
	return _fag.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065")
}
func (_efb *CT_Shape) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _efb.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _cf.Sprintf("\u0025\u0076", *_efb.MacroAttr)})
	}
	if _efb.TextlinkAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b"}, Value: _cf.Sprintf("\u0025\u0076", *_efb.TextlinkAttr)})
	}
	if _efb.FLocksTextAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074"}, Value: _cf.Sprintf("\u0025\u0064", _fgcd(*_efb.FLocksTextAttr))})
	}
	if _efb.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _cf.Sprintf("\u0025\u0064", _fgcd(*_efb.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_gebd := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u006e\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_efb.NvSpPr, _gebd)
	_gae := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_efb.SpPr, _gae)
	if _efb.Style != nil {
		_aeg := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0073\u0074\u0079\u006ce"}}
		e.EncodeElement(_efb.Style, _aeg)
	}
	if _efb.TxBody != nil {
		_adbb := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0074\u0078\u0042\u006f\u0064\u0079"}}
		e.EncodeElement(_efb.TxBody, _adbb)
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type CT_GraphicalObjectFrame struct {
	MacroAttr        *string
	FPublishedAttr   *bool
	NvGraphicFramePr *CT_GraphicalObjectFrameNonVisual
	Xfrm             *_b.CT_Transform2D
	Graphic          *_b.Graphic
}

func NewEG_ObjectChoicesChoice() *EG_ObjectChoicesChoice {
	_bfae := &EG_ObjectChoicesChoice{}
	return _bfae
}

// ValidateWithPath validates the CT_ShapeNonVisual and its children, prefixing error messages with path
func (_ccfg *CT_ShapeNonVisual) ValidateWithPath(path string) error {
	if _agb := _ccfg.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _agb != nil {
		return _agb
	}
	if _ffgf := _ccfg.CNvSpPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0053\u0070\u0050\u0072"); _ffgf != nil {
		return _ffgf
	}
	return nil
}
func NewCT_GroupShape() *CT_GroupShape {
	_bfd := &CT_GroupShape{}
	_bfd.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_bfd.GrpSpPr = _b.NewCT_GroupShapeProperties()
	return _bfd
}

type CT_Connector struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvCxnSpPr      *CT_ConnectorNonVisual
	SpPr           *_b.CT_ShapeProperties
	Style          *_b.CT_ShapeStyle
}
type CT_Picture struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvPicPr        *CT_PictureNonVisual
	BlipFill       *_b.CT_BlipFillProperties
	SpPr           *_b.CT_ShapeProperties
	Style          *_b.CT_ShapeStyle
}

func (_ggbfb *CT_GroupShapeChoice) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_ebc:
	for {
		_ebba, _efe := d.Token()
		if _efe != nil {
			return _efe
		}
		switch _bddc := _ebba.(type) {
		case _c.StartElement:
			switch _bddc.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_bec := NewCT_Shape()
				if _acb := d.DecodeElement(_bec, &_bddc); _acb != nil {
					return _acb
				}
				_ggbfb.Sp = append(_ggbfb.Sp, _bec)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_fec := NewCT_GroupShape()
				if _bfed := d.DecodeElement(_fec, &_bddc); _bfed != nil {
					return _bfed
				}
				_ggbfb.GrpSp = append(_ggbfb.GrpSp, _fec)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_cee := NewCT_GraphicalObjectFrame()
				if _babe := d.DecodeElement(_cee, &_bddc); _babe != nil {
					return _babe
				}
				_ggbfb.GraphicFrame = append(_ggbfb.GraphicFrame, _cee)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_gfb := NewCT_Connector()
				if _eeb := d.DecodeElement(_gfb, &_bddc); _eeb != nil {
					return _eeb
				}
				_ggbfb.CxnSp = append(_ggbfb.CxnSp, _gfb)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_eebe := NewCT_Picture()
				if _gfg := d.DecodeElement(_eebe, &_bddc); _gfg != nil {
					return _gfg
				}
				_ggbfb.Pic = append(_ggbfb.Pic, _eebe)
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068ap\u0065\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _bddc.Name)
				if _ade := d.Skip(); _ade != nil {
					return _ade
				}
			}
		case _c.EndElement:
			break _ebc
		case _c.CharData:
		}
	}
	return nil
}
func (_ggg *EG_Anchor) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _ggg.TwoCellAnchor != nil {
		_effc := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_ggg.TwoCellAnchor, _effc)
	}
	if _ggg.OneCellAnchor != nil {
		_ceag := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_ggg.OneCellAnchor, _ceag)
	}
	if _ggg.AbsoluteAnchor != nil {
		_bddcb := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072:\u0061\u0062\u0073\u006f\u006cu\u0074\u0065A\u006e\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_ggg.AbsoluteAnchor, _bddcb)
	}
	return nil
}
func (_fgcb *CT_GroupShapeNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_fgcb.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_fgcb.CNvGrpSpPr = _b.NewCT_NonVisualGroupDrawingShapeProps()
_gbc:
	for {
		_aba, _aee := d.Token()
		if _aee != nil {
			return _aee
		}
		switch _fegee := _aba.(type) {
		case _c.StartElement:
			switch _fegee.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _db := d.DecodeElement(_fgcb.CNvPr, &_fegee); _db != nil {
					return _db
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}:
				if _ebge := d.DecodeElement(_fgcb.CNvGrpSpPr, &_fegee); _ebge != nil {
					return _ebge
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0047\u0072\u006f\u0075p\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _fegee.Name)
				if _acg := d.Skip(); _acg != nil {
					return _acg
				}
			}
		case _c.EndElement:
			break _gbc
		case _c.CharData:
		}
	}
	return nil
}
func NewCT_ShapeNonVisual() *CT_ShapeNonVisual {
	_dec := &CT_ShapeNonVisual{}
	_dec.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_dec.CNvSpPr = _b.NewCT_NonVisualDrawingShapeProps()
	return _dec
}
func (_ebcd *WsDr) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ebcd.CT_Drawing = *NewCT_Drawing()
_ddb:
	for {
		_badg, _gdf := d.Token()
		if _gdf != nil {
			return _gdf
		}
		switch _cabd := _badg.(type) {
		case _c.StartElement:
			switch _cabd.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_ebga := NewEG_Anchor()
				_ebga.TwoCellAnchor = NewCT_TwoCellAnchor()
				if _gaa := d.DecodeElement(_ebga.TwoCellAnchor, &_cabd); _gaa != nil {
					return _gaa
				}
				_ebcd.EG_Anchor = append(_ebcd.EG_Anchor, _ebga)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_gfdf := NewEG_Anchor()
				_gfdf.OneCellAnchor = NewCT_OneCellAnchor()
				if _gbeb := d.DecodeElement(_gfdf.OneCellAnchor, &_cabd); _gbeb != nil {
					return _gbeb
				}
				_ebcd.EG_Anchor = append(_ebcd.EG_Anchor, _gfdf)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}:
				_efeb := NewEG_Anchor()
				_efeb.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if _cgbe := d.DecodeElement(_efeb.AbsoluteAnchor, &_cabd); _cgbe != nil {
					return _cgbe
				}
				_ebcd.EG_Anchor = append(_ebcd.EG_Anchor, _efeb)
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0057\u0073D\u0072\u0020\u0025\u0076", _cabd.Name)
				if _edgg := d.Skip(); _edgg != nil {
					return _edgg
				}
			}
		case _c.EndElement:
			break _ddb
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the WsDr and its children
func (_adba *WsDr) Validate() error { return _adba.ValidateWithPath("\u0057\u0073\u0044\u0072") }

const (
	ST_EditAsUnset    ST_EditAs = 0
	ST_EditAsTwoCell  ST_EditAs = 1
	ST_EditAsOneCell  ST_EditAs = 2
	ST_EditAsAbsolute ST_EditAs = 3
)

// Validate validates the CT_PictureNonVisual and its children
func (_abg *CT_PictureNonVisual) Validate() error {
	return _abg.ValidateWithPath("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

// Validate validates the EG_Anchor and its children
func (_bde *EG_Anchor) Validate() error {
	return _bde.ValidateWithPath("\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr")
}

// Validate validates the CT_GraphicalObjectFrame and its children
func (_bdg *CT_GraphicalObjectFrame) Validate() error {
	return _bdg.ValidateWithPath("\u0043\u0054\u005fGr\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006a\u0065\u0063\u0074\u0046\u0072\u0061\u006d\u0065")
}
func (_bbd ST_EditAs) Validate() error { return _bbd.ValidateWithPath("") }
func (_dcfc *CT_Rel) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	for _, _fbd := range start.Attr {
		if _fbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _fbd.Name.Local == "\u0069\u0064" || _fbd.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _fbd.Name.Local == "\u0069\u0064" {
			_fdgad, _bfbe := _fbd.Value, error(nil)
			if _bfbe != nil {
				return _bfbe
			}
			_dcfc.IdAttr = _fdgad
			continue
		}
	}
	for {
		_ddcb, _fffc := d.Token()
		if _fffc != nil {
			return _cf.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073", _fffc)
		}
		if _egcg, _gdd := _ddcb.(_c.EndElement); _gdd && _egcg.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_GraphicalObjectFrameNonVisual struct {
	CNvPr             *_b.CT_NonVisualDrawingProps
	CNvGraphicFramePr *_b.CT_NonVisualGraphicFrameProperties
}
type CT_AnchorClientData struct {
	FLocksWithSheetAttr  *bool
	FPrintsWithSheetAttr *bool
}

// Validate validates the CT_Rel and its children
func (_fce *CT_Rel) Validate() error {
	return _fce.ValidateWithPath("\u0043\u0054\u005f\u0052\u0065\u006c")
}
func NewEG_Anchor() *EG_Anchor { _bbg := &EG_Anchor{}; return _bbg }

type EG_ObjectChoicesChoice struct {
	Sp           *CT_Shape
	GrpSp        *CT_GroupShape
	GraphicFrame *CT_GraphicalObjectFrame
	CxnSp        *CT_Connector
	Pic          *CT_Picture
	ContentPart  *CT_Rel
}

// Validate validates the To and its children
func (_bded *To) Validate() error { return _bded.ValidateWithPath("\u0054\u006f") }
func (_bbcfg ST_EditAs) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	return e.EncodeElement(_bbcfg.String(), start)
}

// Validate validates the CT_GroupShapeChoice and its children
func (_ae *CT_GroupShapeChoice) Validate() error {
	return _ae.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u0043\u0068\u006f\u0069\u0063\u0065")
}

// ValidateWithPath validates the CT_AnchorClientData and its children, prefixing error messages with path
func (_fege *CT_AnchorClientData) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the CT_GroupShape and its children, prefixing error messages with path
func (_feaf *CT_GroupShape) ValidateWithPath(path string) error {
	if _eee := _feaf.NvGrpSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _eee != nil {
		return _eee
	}
	if _gf := _feaf.GrpSpPr.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _gf != nil {
		return _gf
	}
	for _gebb, _eeg := range _feaf.Choice {
		if _ecdg := _eeg.ValidateWithPath(_cf.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _gebb)); _ecdg != nil {
			return _ecdg
		}
	}
	return nil
}

type CT_GroupShapeChoice struct {
	Sp           []*CT_Shape
	GrpSp        []*CT_GroupShape
	GraphicFrame []*CT_GraphicalObjectFrame
	CxnSp        []*CT_Connector
	Pic          []*CT_Picture
}

// Validate validates the CT_ConnectorNonVisual and its children
func (_ga *CT_ConnectorNonVisual) Validate() error {
	return _ga.ValidateWithPath("C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c")
}
func NewCT_Drawing() *CT_Drawing { _dea := &CT_Drawing{}; return _dea }
func (_cda *CT_GroupShapeChoice) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _cda.Sp != nil {
		_cag := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070"}}
		for _, _edgb := range _cda.Sp {
			e.EncodeElement(_edgb, _cag)
		}
	}
	if _cda.GrpSp != nil {
		_adc := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0067\u0072\u0070\u0053p"}}
		for _, _cgbg := range _cda.GrpSp {
			e.EncodeElement(_cgbg, _adc)
		}
	}
	if _cda.GraphicFrame != nil {
		_aabg := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064r\u003a\u0067\u0072a\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		for _, _eege := range _cda.GraphicFrame {
			e.EncodeElement(_eege, _aabg)
		}
	}
	if _cda.CxnSp != nil {
		_baafc := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u0078\u006e\u0053p"}}
		for _, _fdb := range _cda.CxnSp {
			e.EncodeElement(_fdb, _baafc)
		}
	}
	if _cda.Pic != nil {
		_cbb := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0070\u0069\u0063"}}
		for _, _cde := range _cda.Pic {
			e.EncodeElement(_cde, _cbb)
		}
	}
	return nil
}

// ValidateWithPath validates the EG_ObjectChoicesChoice and its children, prefixing error messages with path
func (_bge *EG_ObjectChoicesChoice) ValidateWithPath(path string) error {
	if _bge.Sp != nil {
		if _aag := _bge.Sp.ValidateWithPath(path + "\u002f\u0053\u0070"); _aag != nil {
			return _aag
		}
	}
	if _bge.GrpSp != nil {
		if _bbcf := _bge.GrpSp.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070"); _bbcf != nil {
			return _bbcf
		}
	}
	if _bge.GraphicFrame != nil {
		if _acgb := _bge.GraphicFrame.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"); _acgb != nil {
			return _acgb
		}
	}
	if _bge.CxnSp != nil {
		if _abbg := _bge.CxnSp.ValidateWithPath(path + "\u002f\u0043\u0078\u006e\u0053\u0070"); _abbg != nil {
			return _abbg
		}
	}
	if _bge.Pic != nil {
		if _gga := _bge.Pic.ValidateWithPath(path + "\u002f\u0050\u0069\u0063"); _gga != nil {
			return _gga
		}
	}
	if _bge.ContentPart != nil {
		if _cbgb := _bge.ContentPart.ValidateWithPath(path + "\u002f\u0043\u006fn\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"); _cbgb != nil {
			return _cbgb
		}
	}
	return nil
}
func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	_geba := &CT_PictureNonVisual{}
	_geba.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_geba.CNvPicPr = _b.NewCT_NonVisualPictureProperties()
	return _geba
}
func (_gbab *EG_Anchor) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_ccbc:
	for {
		_gfeg, _aaa := d.Token()
		if _aaa != nil {
			return _aaa
		}
		switch _fdgf := _gfeg.(type) {
		case _c.StartElement:
			switch _fdgf.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_gbab.TwoCellAnchor = NewCT_TwoCellAnchor()
				if _gggc := d.DecodeElement(_gbab.TwoCellAnchor, &_fdgf); _gggc != nil {
					return _gggc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_gbab.OneCellAnchor = NewCT_OneCellAnchor()
				if _gbaf := d.DecodeElement(_gbab.OneCellAnchor, &_fdgf); _gbaf != nil {
					return _gbaf
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}:
				_gbab.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if _faa := d.DecodeElement(_gbab.AbsoluteAnchor, &_fdgf); _faa != nil {
					return _faa
				}
			default:
				_f.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0045\u0047\u005f\u0041\u006e\u0063h\u006f\u0072 \u0025\u0076", _fdgf.Name)
				if _cfge := d.Skip(); _cfge != nil {
					return _cfge
				}
			}
		case _c.EndElement:
			break _ccbc
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the To and its children, prefixing error messages with path
func (_ffdb *To) ValidateWithPath(path string) error {
	if _aded := _ffdb.CT_Marker.ValidateWithPath(path); _aded != nil {
		return _aded
	}
	return nil
}
func (_ab *CT_ConnectorNonVisual) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ab.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_ab.CNvCxnSpPr = _b.NewCT_NonVisualConnectorProperties()
_adb:
	for {
		_aac, _daa := d.Token()
		if _daa != nil {
			return _daa
		}
		switch _bc := _aac.(type) {
		case _c.StartElement:
			switch _bc.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _aae := d.DecodeElement(_ab.CNvPr, &_bc); _aae != nil {
					return _aae
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}:
				if _fae := d.DecodeElement(_ab.CNvCxnSpPr, &_bc); _fae != nil {
					return _fae
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075n\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u006f\u006e C\u0054\u005f\u0043\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _bc.Name)
				if _bfb := d.Skip(); _bfb != nil {
					return _bfb
				}
			}
		case _c.EndElement:
			break _adb
		case _c.CharData:
		}
	}
	return nil
}
func (_bdbe ST_EditAs) String() string {
	switch _bdbe {
	case 0:
		return ""
	case 1:
		return "\u0074w\u006f\u0043\u0065\u006c\u006c"
	case 2:
		return "\u006fn\u0065\u0043\u0065\u006c\u006c"
	case 3:
		return "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065"
	}
	return ""
}

// ValidateWithPath validates the CT_TwoCellAnchor and its children, prefixing error messages with path
func (_dgg *CT_TwoCellAnchor) ValidateWithPath(path string) error {
	if _gfe := _dgg.EditAsAttr.ValidateWithPath(path + "/\u0045\u0064\u0069\u0074\u0041\u0073\u0041\u0074\u0074\u0072"); _gfe != nil {
		return _gfe
	}
	if _beeg := _dgg.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _beeg != nil {
		return _beeg
	}
	if _fdbg := _dgg.To.ValidateWithPath(path + "\u002f\u0054\u006f"); _fdbg != nil {
		return _fdbg
	}
	if _dgg.Choice != nil {
		if _fgff := _dgg.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _fgff != nil {
			return _fgff
		}
	}
	if _gade := _dgg.ClientData.ValidateWithPath(path + "/\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"); _gade != nil {
		return _gade
	}
	return nil
}
func _fgcd(_fabd bool) uint8 {
	if _fabd {
		return 1
	}
	return 0
}
func (_aaef *CT_TwoCellAnchor) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_aaef.From = NewCT_Marker()
	_aaef.To = NewCT_Marker()
	_aaef.ClientData = NewCT_AnchorClientData()
	for _, _cae := range start.Attr {
		if _cae.Name.Local == "\u0065\u0064\u0069\u0074\u0041\u0073" {
			_aaef.EditAsAttr.UnmarshalXMLAttr(_cae)
			continue
		}
	}
_cagd:
	for {
		_cebd, _fefg := d.Token()
		if _fefg != nil {
			return _fefg
		}
		switch _bcee := _cebd.(type) {
		case _c.StartElement:
			switch _bcee.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _dddb := d.DecodeElement(_aaef.From, &_bcee); _dddb != nil {
					return _dddb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u006f"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u006f"}:
				if _dfge := d.DecodeElement(_aaef.To, &_bcee); _dfge != nil {
					return _dfge
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_aaef.Choice = NewEG_ObjectChoicesChoice()
				if _fab := d.DecodeElement(&_aaef.Choice.Sp, &_bcee); _fab != nil {
					return _fab
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_aaef.Choice = NewEG_ObjectChoicesChoice()
				if _bcc := d.DecodeElement(&_aaef.Choice.GrpSp, &_bcee); _bcc != nil {
					return _bcc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_aaef.Choice = NewEG_ObjectChoicesChoice()
				if _ddega := d.DecodeElement(&_aaef.Choice.GraphicFrame, &_bcee); _ddega != nil {
					return _ddega
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_aaef.Choice = NewEG_ObjectChoicesChoice()
				if _cgcd := d.DecodeElement(&_aaef.Choice.CxnSp, &_bcee); _cgcd != nil {
					return _cgcd
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_aaef.Choice = NewEG_ObjectChoicesChoice()
				if _baef := d.DecodeElement(&_aaef.Choice.Pic, &_bcee); _baef != nil {
					return _baef
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_aaef.Choice = NewEG_ObjectChoicesChoice()
				if _dgdac := d.DecodeElement(&_aaef.Choice.ContentPart, &_bcee); _dgdac != nil {
					return _dgdac
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}:
				if _fdcgb := d.DecodeElement(_aaef.ClientData, &_bcee); _fdcgb != nil {
					return _fdcgb
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _bcee.Name)
				if _fbff := d.Skip(); _fbff != nil {
					return _fbff
				}
			}
		case _c.EndElement:
			break _cagd
		case _c.CharData:
		}
	}
	return nil
}
func (_fbb *ST_EditAs) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_defa, _aggf := d.Token()
	if _aggf != nil {
		return _aggf
	}
	if _aggc, _ecb := _defa.(_c.EndElement); _ecb && _aggc.Name == start.Name {
		*_fbb = 1
		return nil
	}
	if _ged, _ffgg := _defa.(_c.CharData); !_ffgg {
		return _cf.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _defa)
	} else {
		switch string(_ged) {
		case "":
			*_fbb = 0
		case "\u0074w\u006f\u0043\u0065\u006c\u006c":
			*_fbb = 1
		case "\u006fn\u0065\u0043\u0065\u006c\u006c":
			*_fbb = 2
		case "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065":
			*_fbb = 3
		}
	}
	_defa, _aggf = d.Token()
	if _aggf != nil {
		return _aggf
	}
	if _aegf, _fbga := _defa.(_c.EndElement); _fbga && _aegf.Name == start.Name {
		return nil
	}
	return _cf.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _defa)
}

type CT_Marker struct {
	Col    int32
	ColOff _b.ST_Coordinate
	Row    int32
	RowOff _b.ST_Coordinate
}

func (_cff *To) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_cff.CT_Marker = *NewCT_Marker()
_aabd:
	for {
		_baed, _fcc := d.Token()
		if _fcc != nil {
			return _fcc
		}
		switch _daba := _baed.(type) {
		case _c.StartElement:
			switch _daba.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}:
				if _ebcb := d.DecodeElement(&_cff.Col, &_daba); _ebcb != nil {
					return _ebcb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}:
				_dff, _aeeb := d.Token()
				if _aeeb != nil {
					return _aeeb
				}
				switch _cdeb := _dff.(type) {
				case _c.CharData:
					_fgffa := string(_cdeb)
					_dabc, _fafa := _b.ParseUnionST_Coordinate(_fgffa)
					if _fafa != nil {
						return nil
					}
					_cff.ColOff = _dabc
					d.Skip()
				case _c.EndElement:
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}:
				if _afc := d.DecodeElement(&_cff.Row, &_daba); _afc != nil {
					return _afc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}:
				_fbc, _cgd := d.Token()
				if _cgd != nil {
					return _cgd
				}
				switch _aceaf := _fbc.(type) {
				case _c.CharData:
					_bcdge := string(_aceaf)
					_edce, _bbb := _b.ParseUnionST_Coordinate(_bcdge)
					if _bbb != nil {
						return nil
					}
					_cff.RowOff = _edce
					d.Skip()
				case _c.EndElement:
				}
			default:
				_f.Log("\u0073\u006bi\u0070\u0070\u0069\u006eg\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020T\u006f\u0020\u0025\u0076", _daba.Name)
				if _ggcf := d.Skip(); _ggcf != nil {
					return _ggcf
				}
			}
		case _c.EndElement:
			break _aabd
		case _c.CharData:
		}
	}
	return nil
}
func (_cdb *CT_GroupShapeNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_eegb := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_cdb.CNvPr, _eegb)
	_ceea := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u004e\u0076\u0047\u0072p\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_cdb.CNvGrpSpPr, _ceea)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func NewCT_Rel() *CT_Rel { _fbfg := &CT_Rel{}; return _fbfg }

// ValidateWithPath validates the CT_Connector and its children, prefixing error messages with path
func (_ee *CT_Connector) ValidateWithPath(path string) error {
	if _ag := _ee.NvCxnSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _ag != nil {
		return _ag
	}
	if _fdcd := _ee.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _fdcd != nil {
		return _fdcd
	}
	if _ee.Style != nil {
		if _agd := _ee.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _agd != nil {
			return _agd
		}
	}
	return nil
}

type From struct{ CT_Marker }

func NewCT_Marker() *CT_Marker { _dgf := &CT_Marker{}; _dgf.Col = 0; _dgf.Row = 0; return _dgf }
func (_edd *CT_GroupShape) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_edd.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_edd.GrpSpPr = _b.NewCT_GroupShapeProperties()
_cecd:
	for {
		_efae, _feb := d.Token()
		if _feb != nil {
			return _feb
		}
		switch _bafg := _efae.(type) {
		case _c.StartElement:
			switch _bafg.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}:
				if _gcf := d.DecodeElement(_edd.NvGrpSpPr, &_bafg); _gcf != nil {
					return _gcf
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}:
				if _age := d.DecodeElement(_edd.GrpSpPr, &_bafg); _age != nil {
					return _age
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_afa := NewCT_GroupShapeChoice()
				if _baaf := d.DecodeElement(&_afa.Sp, &_bafg); _baaf != nil {
					return _baaf
				}
				_edd.Choice = append(_edd.Choice, _afa)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_bbf := NewCT_GroupShapeChoice()
				if _egg := d.DecodeElement(&_bbf.GrpSp, &_bafg); _egg != nil {
					return _egg
				}
				_edd.Choice = append(_edd.Choice, _bbf)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_bcgf := NewCT_GroupShapeChoice()
				if _ccd := d.DecodeElement(&_bcgf.GraphicFrame, &_bafg); _ccd != nil {
					return _ccd
				}
				_edd.Choice = append(_edd.Choice, _bcgf)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_bfe := NewCT_GroupShapeChoice()
				if _feaa := d.DecodeElement(&_bfe.CxnSp, &_bafg); _feaa != nil {
					return _feaa
				}
				_edd.Choice = append(_edd.Choice, _bfe)
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_gee := NewCT_GroupShapeChoice()
				if _adg := d.DecodeElement(&_gee.Pic, &_bafg); _adg != nil {
					return _adg
				}
				_edd.Choice = append(_edd.Choice, _gee)
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047r\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065 \u0025\u0076", _bafg.Name)
				if _gad := d.Skip(); _gad != nil {
					return _gad
				}
			}
		case _c.EndElement:
			break _cecd
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_GroupShapeChoice and its children, prefixing error messages with path
func (_fga *CT_GroupShapeChoice) ValidateWithPath(path string) error {
	for _efad, _afae := range _fga.Sp {
		if _eba := _afae.ValidateWithPath(_cf.Sprintf("\u0025s\u002f\u0053\u0070\u005b\u0025\u0064]", path, _efad)); _eba != nil {
			return _eba
		}
	}
	for _cafc, _cfe := range _fga.GrpSp {
		if _bcdg := _cfe.ValidateWithPath(_cf.Sprintf("\u0025\u0073\u002fG\u0072\u0070\u0053\u0070\u005b\u0025\u0064\u005d", path, _cafc)); _bcdg != nil {
			return _bcdg
		}
	}
	for _gcff, _ebf := range _fga.GraphicFrame {
		if _cef := _ebf.ValidateWithPath(_cf.Sprintf("\u0025\u0073\u002f\u0047ra\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u005b\u0025\u0064\u005d", path, _gcff)); _cef != nil {
			return _cef
		}
	}
	for _ecag, _geed := range _fga.CxnSp {
		if _ffee := _geed.ValidateWithPath(_cf.Sprintf("\u0025\u0073\u002fC\u0078\u006e\u0053\u0070\u005b\u0025\u0064\u005d", path, _ecag)); _ffee != nil {
			return _ffee
		}
	}
	for _cbbd, _gdc := range _fga.Pic {
		if _gadd := _gdc.ValidateWithPath(_cf.Sprintf("\u0025\u0073\u002f\u0050\u0069\u0063\u005b\u0025\u0064\u005d", path, _cbbd)); _gadd != nil {
			return _gadd
		}
	}
	return nil
}
func NewEG_ObjectChoices() *EG_ObjectChoices { _dcga := &EG_ObjectChoices{}; return _dcga }
func NewCT_ConnectorNonVisual() *CT_ConnectorNonVisual {
	_gg := &CT_ConnectorNonVisual{}
	_gg.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_gg.CNvCxnSpPr = _b.NewCT_NonVisualConnectorProperties()
	return _gg
}
func (_fbffe *WsDr) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0064r"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u0064\u0072\u003a\u0077\u0073\u0044\u0072"
	return _fbffe.CT_Drawing.MarshalXML(e, start)
}
func (_fge *EG_ObjectChoices) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _fge.Choice != nil {
		_fge.Choice.MarshalXML(e, _c.StartElement{})
	}
	return nil
}

// Validate validates the From and its children
func (_ecca *From) Validate() error { return _ecca.ValidateWithPath("\u0046\u0072\u006f\u006d") }

// Validate validates the CT_Connector and its children
func (_ec *CT_Connector) Validate() error {
	return _ec.ValidateWithPath("\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072")
}
func NewCT_GraphicalObjectFrameNonVisual() *CT_GraphicalObjectFrameNonVisual {
	_bgcc := &CT_GraphicalObjectFrameNonVisual{}
	_bgcc.CNvPr = _b.NewCT_NonVisualDrawingProps()
	_bgcc.CNvGraphicFramePr = _b.NewCT_NonVisualGraphicFrameProperties()
	return _bgcc
}
func (_aacb *To) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0064r"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u0064\u0072\u003a\u0074\u006f"
	return _aacb.CT_Marker.MarshalXML(e, start)
}

// Validate validates the CT_Picture and its children
func (_gfa *CT_Picture) Validate() error {
	return _gfa.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}
func NewCT_AbsoluteAnchor() *CT_AbsoluteAnchor {
	_fd := &CT_AbsoluteAnchor{}
	_fd.Pos = _b.NewCT_Point2D()
	_fd.Ext = _b.NewCT_PositiveSize2D()
	_fd.ClientData = NewCT_AnchorClientData()
	return _fd
}

type EG_Anchor struct {
	TwoCellAnchor  *CT_TwoCellAnchor
	OneCellAnchor  *CT_OneCellAnchor
	AbsoluteAnchor *CT_AbsoluteAnchor
}

func (_cb *CT_AbsoluteAnchor) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_g := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0070\u006f\u0073"}}
	e.EncodeElement(_cb.Pos, _g)
	_eb := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0065\u0078\u0074"}}
	e.EncodeElement(_cb.Ext, _eb)
	if _cb.Choice != nil {
		_cb.Choice.MarshalXML(e, _c.StartElement{})
	}
	_bd := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006c\u0069\u0065\u006et\u0044\u0061\u0074\u0061"}}
	e.EncodeElement(_cb.ClientData, _bd)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

type CT_Rel struct{ IdAttr string }

func NewTo() *To { _baec := &To{}; _baec.CT_Marker = *NewCT_Marker(); return _baec }
func (_cfae *CT_Marker) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_edbd := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u006f\u006c"}}
	e.EncodeElement(_cfae.Col, _edbd)
	_bba := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006f\u006c\u004f\u0066\u0066"}}
	e.EncodeElement(_cfae.ColOff, _bba)
	_effd := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0072\u006f\u0077"}}
	e.EncodeElement(_cfae.Row, _effd)
	_ace := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0072\u006f\u0077\u004f\u0066\u0066"}}
	e.EncodeElement(_cfae.RowOff, _ace)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_bfbb ST_EditAs) MarshalXMLAttr(name _c.Name) (_c.Attr, error) {
	_bfdd := _c.Attr{}
	_bfdd.Name = name
	switch _bfbb {
	case ST_EditAsUnset:
		_bfdd.Value = ""
	case ST_EditAsTwoCell:
		_bfdd.Value = "\u0074w\u006f\u0043\u0065\u006c\u006c"
	case ST_EditAsOneCell:
		_bfdd.Value = "\u006fn\u0065\u0043\u0065\u006c\u006c"
	case ST_EditAsAbsolute:
		_bfdd.Value = "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065"
	}
	return _bfdd, nil
}

// ValidateWithPath validates the CT_OneCellAnchor and its children, prefixing error messages with path
func (_afb *CT_OneCellAnchor) ValidateWithPath(path string) error {
	if _fdg := _afb.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _fdg != nil {
		return _fdg
	}
	if _ddec := _afb.Ext.ValidateWithPath(path + "\u002f\u0045\u0078\u0074"); _ddec != nil {
		return _ddec
	}
	if _afb.Choice != nil {
		if _faeb := _afb.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _faeb != nil {
			return _faeb
		}
	}
	if _gcc := _afb.ClientData.ValidateWithPath(path + "/\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"); _gcc != nil {
		return _gcc
	}
	return nil
}
func (_fca *CT_AnchorClientData) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _fca.FLocksWithSheetAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066L\u006fc\u006b\u0073\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074"}, Value: _cf.Sprintf("\u0025\u0064", _fgcd(*_fca.FLocksWithSheetAttr))})
	}
	if _fca.FPrintsWithSheetAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050r\u0069\u006e\u0074s\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074"}, Value: _cf.Sprintf("\u0025\u0064", _fgcd(*_fca.FPrintsWithSheetAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the EG_ObjectChoices and its children, prefixing error messages with path
func (_bgg *EG_ObjectChoices) ValidateWithPath(path string) error {
	if _bgg.Choice != nil {
		if _eced := _bgg.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _eced != nil {
			return _eced
		}
	}
	return nil
}

// Validate validates the CT_Shape and its children
func (_bbe *CT_Shape) Validate() error {
	return _bbe.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065")
}

type CT_GroupShapeNonVisual struct {
	CNvPr      *_b.CT_NonVisualDrawingProps
	CNvGrpSpPr *_b.CT_NonVisualGroupDrawingShapeProps
}

func (_ceead *CT_OneCellAnchor) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ceead.From = NewCT_Marker()
	_ceead.Ext = _b.NewCT_PositiveSize2D()
	_ceead.ClientData = NewCT_AnchorClientData()
_dgd:
	for {
		_ebd, _gbcf := d.Token()
		if _gbcf != nil {
			return _gbcf
		}
		switch _fead := _ebd.(type) {
		case _c.StartElement:
			switch _fead.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _acfc := d.DecodeElement(_ceead.From, &_fead); _acfc != nil {
					return _acfc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}:
				if _ceae := d.DecodeElement(_ceead.Ext, &_fead); _ceae != nil {
					return _ceae
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_ceead.Choice = NewEG_ObjectChoicesChoice()
				if _fdeb := d.DecodeElement(&_ceead.Choice.Sp, &_fead); _fdeb != nil {
					return _fdeb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_ceead.Choice = NewEG_ObjectChoicesChoice()
				if _dba := d.DecodeElement(&_ceead.Choice.GrpSp, &_fead); _dba != nil {
					return _dba
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_ceead.Choice = NewEG_ObjectChoicesChoice()
				if _gfd := d.DecodeElement(&_ceead.Choice.GraphicFrame, &_fead); _gfd != nil {
					return _gfd
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_ceead.Choice = NewEG_ObjectChoicesChoice()
				if _agf := d.DecodeElement(&_ceead.Choice.CxnSp, &_fead); _agf != nil {
					return _agf
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_ceead.Choice = NewEG_ObjectChoicesChoice()
				if _dgda := d.DecodeElement(&_ceead.Choice.Pic, &_fead); _dgda != nil {
					return _dgda
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_ceead.Choice = NewEG_ObjectChoicesChoice()
				if _gce := d.DecodeElement(&_ceead.Choice.ContentPart, &_fead); _gce != nil {
					return _gce
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}:
				if _bff := d.DecodeElement(_ceead.ClientData, &_fead); _bff != nil {
					return _bff
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _fead.Name)
				if _fafb := d.Skip(); _fafb != nil {
					return _fafb
				}
			}
		case _c.EndElement:
			break _dgd
		case _c.CharData:
		}
	}
	return nil
}

type WsDr struct{ CT_Drawing }

func (_dfeeg *CT_Picture) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	if _dfeeg.MacroAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _cf.Sprintf("\u0025\u0076", *_dfeeg.MacroAttr)})
	}
	if _dfeeg.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _cf.Sprintf("\u0025\u0064", _fgcd(*_dfeeg.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_ffg := _c.StartElement{Name: _c.Name{Local: "x\u0064\u0072\u003a\u006e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_dfeeg.NvPicPr, _ffg)
	_egc := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072:\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}}
	e.EncodeElement(_dfeeg.BlipFill, _egc)
	_cac := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_dfeeg.SpPr, _cac)
	if _dfeeg.Style != nil {
		_efc := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0073\u0074\u0079\u006ce"}}
		e.EncodeElement(_dfeeg.Style, _efc)
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_ecab *CT_OneCellAnchor) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_gbca := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_ecab.From, _gbca)
	_bbab := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0065\u0078\u0074"}}
	e.EncodeElement(_ecab.Ext, _bbab)
	if _ecab.Choice != nil {
		_ecab.Choice.MarshalXML(e, _c.StartElement{})
	}
	_gea := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006c\u0069\u0065\u006et\u0044\u0061\u0074\u0061"}}
	e.EncodeElement(_ecab.ClientData, _gea)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_ea *CT_AbsoluteAnchor) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ea.Pos = _b.NewCT_Point2D()
	_ea.Ext = _b.NewCT_PositiveSize2D()
	_ea.ClientData = NewCT_AnchorClientData()
_caf:
	for {
		_bdd, _d := d.Token()
		if _d != nil {
			return _d
		}
		switch _fdc := _bdd.(type) {
		case _c.StartElement:
			switch _fdc.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u006f\u0073"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u006f\u0073"}:
				if _fc := d.DecodeElement(_ea.Pos, &_fdc); _fc != nil {
					return _fc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}:
				if _bb := d.DecodeElement(_ea.Ext, &_fdc); _bb != nil {
					return _bb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_ea.Choice = NewEG_ObjectChoicesChoice()
				if _gc := d.DecodeElement(&_ea.Choice.Sp, &_fdc); _gc != nil {
					return _gc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_ea.Choice = NewEG_ObjectChoicesChoice()
				if _cc := d.DecodeElement(&_ea.Choice.GrpSp, &_fdc); _cc != nil {
					return _cc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_ea.Choice = NewEG_ObjectChoicesChoice()
				if _ef := d.DecodeElement(&_ea.Choice.GraphicFrame, &_fdc); _ef != nil {
					return _ef
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_ea.Choice = NewEG_ObjectChoicesChoice()
				if _ba := d.DecodeElement(&_ea.Choice.CxnSp, &_fdc); _ba != nil {
					return _ba
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_ea.Choice = NewEG_ObjectChoicesChoice()
				if _a := d.DecodeElement(&_ea.Choice.Pic, &_fdc); _a != nil {
					return _a
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_ea.Choice = NewEG_ObjectChoicesChoice()
				if _fa := d.DecodeElement(&_ea.Choice.ContentPart, &_fdc); _fa != nil {
					return _fa
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}:
				if _fb := d.DecodeElement(_ea.ClientData, &_fdc); _fb != nil {
					return _fb
				}
			default:
				_f.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025\u0076", _fdc.Name)
				if _eg := d.Skip(); _eg != nil {
					return _eg
				}
			}
		case _c.EndElement:
			break _caf
		case _c.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the EG_Anchor and its children, prefixing error messages with path
func (_cfeca *EG_Anchor) ValidateWithPath(path string) error {
	if _cfeca.TwoCellAnchor != nil {
		if _bfdb := _cfeca.TwoCellAnchor.ValidateWithPath(path + "\u002f\u0054\u0077\u006f\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"); _bfdb != nil {
			return _bfdb
		}
	}
	if _cfeca.OneCellAnchor != nil {
		if _dedg := _cfeca.OneCellAnchor.ValidateWithPath(path + "\u002f\u004f\u006e\u0065\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"); _dedg != nil {
			return _dedg
		}
	}
	if _cfeca.AbsoluteAnchor != nil {
		if _bfff := _cfeca.AbsoluteAnchor.ValidateWithPath(path + "\u002fA\u0062s\u006f\u006c\u0075\u0074\u0065\u0041\u006e\u0063\u0068\u006f\u0072"); _bfff != nil {
			return _bfff
		}
	}
	return nil
}
func (_dcc *CT_Shape) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_dcc.NvSpPr = NewCT_ShapeNonVisual()
	_dcc.SpPr = _b.NewCT_ShapeProperties()
	for _, _eae := range start.Attr {
		if _eae.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_fcec, _gag := _eae.Value, error(nil)
			if _gag != nil {
				return _gag
			}
			_dcc.MacroAttr = &_fcec
			continue
		}
		if _eae.Name.Local == "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b" {
			_eaa, _facb := _eae.Value, error(nil)
			if _facb != nil {
				return _facb
			}
			_dcc.TextlinkAttr = &_eaa
			continue
		}
		if _eae.Name.Local == "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074" {
			_gdg, _fdf := _ca.ParseBool(_eae.Value)
			if _fdf != nil {
				return _fdf
			}
			_dcc.FLocksTextAttr = &_gdg
			continue
		}
		if _eae.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_eaeb, _gdde := _ca.ParseBool(_eae.Value)
			if _gdde != nil {
				return _gdde
			}
			_dcc.FPublishedAttr = &_eaeb
			continue
		}
	}
_edgc:
	for {
		_ffge, _ccfe := d.Token()
		if _ccfe != nil {
			return _ccfe
		}
		switch _eeee := _ffge.(type) {
		case _c.StartElement:
			switch _eeee.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}:
				if _cfaeb := d.DecodeElement(_dcc.NvSpPr, &_eeee); _cfaeb != nil {
					return _cfaeb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _aec := d.DecodeElement(_dcc.SpPr, &_eeee); _aec != nil {
					return _aec
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_dcc.Style = _b.NewCT_ShapeStyle()
				if _ceb := d.DecodeElement(_dcc.Style, &_eeee); _ceb != nil {
					return _ceb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}:
				_dcc.TxBody = _b.NewCT_TextBody()
				if _dabg := d.DecodeElement(_dcc.TxBody, &_eeee); _dabg != nil {
					return _dabg
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u0020\u0025\u0076", _eeee.Name)
				if _dgdb := d.Skip(); _dgdb != nil {
					return _dgdb
				}
			}
		case _c.EndElement:
			break _edgc
		case _c.CharData:
		}
	}
	return nil
}
func (_bgcce ST_EditAs) ValidateWithPath(path string) error {
	switch _bgcce {
	case 0, 1, 2, 3:
	default:
		return _cf.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bgcce))
	}
	return nil
}
func NewCT_Shape() *CT_Shape {
	_ccfc := &CT_Shape{}
	_ccfc.NvSpPr = NewCT_ShapeNonVisual()
	_ccfc.SpPr = _b.NewCT_ShapeProperties()
	return _ccfc
}

// ValidateWithPath validates the CT_Marker and its children, prefixing error messages with path
func (_dfee *CT_Marker) ValidateWithPath(path string) error {
	if _dfee.Col < 0 {
		return _cf.Errorf("\u0025\u0073\u002fm\u002e\u0043\u006f\u006c \u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u003d\u0020\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _dfee.Col)
	}
	if _gcg := _dfee.ColOff.ValidateWithPath(path + "\u002fC\u006f\u006c\u004f\u0066\u0066"); _gcg != nil {
		return _gcg
	}
	if _dfee.Row < 0 {
		return _cf.Errorf("\u0025\u0073\u002fm\u002e\u0052\u006f\u0077 \u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u003d\u0020\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _dfee.Row)
	}
	if _bddf := _dfee.RowOff.ValidateWithPath(path + "\u002fR\u006f\u0077\u004f\u0066\u0066"); _bddf != nil {
		return _bddf
	}
	return nil
}
func (_ebe *CT_GraphicalObjectFrameNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_aca := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_ebe.CNvPr, _aca)
	_dcg := _c.StartElement{Name: _c.Name{Local: "x\u0064\u0072\u003a\u0063Nv\u0047r\u0061\u0070\u0068\u0069\u0063F\u0072\u0061\u006d\u0065\u0050\u0072"}}
	e.EncodeElement(_ebe.CNvGraphicFramePr, _dcg)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}
func (_bfg *From) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0064r"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u0064\u0072\u003a\u0066\u0072\u006f\u006d"
	return _bfg.CT_Marker.MarshalXML(e, start)
}
func (_dabdg *CT_PictureNonVisual) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	_ggbd := _c.StartElement{Name: _c.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_dabdg.CNvPr, _ggbd)
	_gfge := _c.StartElement{Name: _c.Name{Local: "\u0078\u0064\u0072:\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_dabdg.CNvPicPr, _gfge)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_AnchorClientData and its children
func (_fg *CT_AnchorClientData) Validate() error {
	return _fg.ValidateWithPath("\u0043\u0054\u005f\u0041nc\u0068\u006f\u0072\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061")
}

type CT_TwoCellAnchor struct {
	EditAsAttr ST_EditAs
	From       *CT_Marker
	To         *CT_Marker
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

// ValidateWithPath validates the WsDr and its children, prefixing error messages with path
func (_bfede *WsDr) ValidateWithPath(path string) error {
	if _cfeg := _bfede.CT_Drawing.ValidateWithPath(path); _cfeg != nil {
		return _cfeg
	}
	return nil
}

// Validate validates the CT_ShapeNonVisual and its children
func (_ggf *CT_ShapeNonVisual) Validate() error {
	return _ggf.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c")
}

// Validate validates the CT_GraphicalObjectFrameNonVisual and its children
func (_dge *CT_GraphicalObjectFrameNonVisual) Validate() error {
	return _dge.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006ae\u0063t\u0046\u0072\u0061\u006d\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}
func (_ggc *EG_ObjectChoices) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_defe:
	for {
		_dcd, _dbgc := d.Token()
		if _dbgc != nil {
			return _dbgc
		}
		switch _dagg := _dcd.(type) {
		case _c.StartElement:
			switch _dagg.Name {
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_ggc.Choice = NewEG_ObjectChoicesChoice()
				if _bgdc := d.DecodeElement(&_ggc.Choice.Sp, &_dagg); _bgdc != nil {
					return _bgdc
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_ggc.Choice = NewEG_ObjectChoicesChoice()
				if _fcb := d.DecodeElement(&_ggc.Choice.GrpSp, &_dagg); _fcb != nil {
					return _fcb
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_ggc.Choice = NewEG_ObjectChoicesChoice()
				if _cdef := d.DecodeElement(&_ggc.Choice.GraphicFrame, &_dagg); _cdef != nil {
					return _cdef
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_ggc.Choice = NewEG_ObjectChoicesChoice()
				if _eeba := d.DecodeElement(&_ggc.Choice.CxnSp, &_dagg); _eeba != nil {
					return _eeba
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_ggc.Choice = NewEG_ObjectChoicesChoice()
				if _dged := d.DecodeElement(&_ggc.Choice.Pic, &_dagg); _dged != nil {
					return _dged
				}
			case _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _c.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_ggc.Choice = NewEG_ObjectChoicesChoice()
				if _ead := d.DecodeElement(&_ggc.Choice.ContentPart, &_dagg); _ead != nil {
					return _ead
				}
			default:
				_f.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u004f\u0062\u006a\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0020\u0025v", _dagg.Name)
				if _aeb := d.Skip(); _aeb != nil {
					return _aeb
				}
			}
		case _c.EndElement:
			break _defe
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TwoCellAnchor and its children
func (_caa *CT_TwoCellAnchor) Validate() error {
	return _caa.ValidateWithPath("\u0043\u0054_\u0054\u0077\u006fC\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072")
}

// ValidateWithPath validates the CT_GroupShapeNonVisual and its children, prefixing error messages with path
func (_deb *CT_GroupShapeNonVisual) ValidateWithPath(path string) error {
	if _gde := _deb.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _gde != nil {
		return _gde
	}
	if _ddeg := _deb.CNvGrpSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _ddeg != nil {
		return _ddeg
	}
	return nil
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_adeab *CT_Picture) ValidateWithPath(path string) error {
	if _ddg := _adeab.NvPicPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072"); _ddg != nil {
		return _ddg
	}
	if _fbe := _adeab.BlipFill.ValidateWithPath(path + "\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl"); _fbe != nil {
		return _fbe
	}
	if _dfba := _adeab.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _dfba != nil {
		return _dfba
	}
	if _adeab.Style != nil {
		if _bafa := _adeab.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _bafa != nil {
			return _bafa
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
	SpPr           *_b.CT_ShapeProperties
	Style          *_b.CT_ShapeStyle
	TxBody         *_b.CT_TextBody
}

// Validate validates the EG_ObjectChoicesChoice and its children
func (_dcbc *EG_ObjectChoicesChoice) Validate() error {
	return _dcbc.ValidateWithPath("\u0045\u0047\u005f\u004fbj\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069c\u0065")
}
func init() {
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0041nc\u0068\u006f\u0072\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061", NewCT_AnchorClientData)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c", NewCT_ShapeNonVisual)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065", NewCT_Shape)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_ConnectorNonVisual)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072", NewCT_Connector)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_PictureNonVisual)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006ae\u0063t\u0046\u0072\u0061\u006d\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_GraphicalObjectFrameNonVisual)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005fGr\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006a\u0065\u0063\u0074\u0046\u0072\u0061\u006d\u0065", NewCT_GraphicalObjectFrame)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c", NewCT_GroupShapeNonVisual)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065", NewCT_GroupShape)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0052\u0065\u006c", NewCT_Rel)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r", NewCT_Marker)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054_\u0054\u0077\u006fC\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_TwoCellAnchor)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054_\u004f\u006e\u0065C\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_OneCellAnchor)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072", NewCT_AbsoluteAnchor)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067", NewCT_Drawing)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0066\u0072\u006f\u006d", NewFrom)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0074\u006f", NewTo)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0077\u0073\u0044\u0072", NewWsDr)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073", NewEG_ObjectChoices)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr", NewEG_Anchor)
}
