package vmldrawing

import (
	_fd "encoding/xml"
	_b "fmt"

	_bf "gitee.com/yu_sheng/gooffice"
	_fe "gitee.com/yu_sheng/gooffice/schema/soo/ofc/sharedTypes"
	_ba "gitee.com/yu_sheng/gooffice/schema/urn/schemas_microsoft_com/office/excel"
	_g "gitee.com/yu_sheng/gooffice/schema/urn/schemas_microsoft_com/vml"
)

type Container struct {
	Layout    *_g.OfcShapelayout
	ShapeType *_g.Shapetype
	Shape     []*_g.Shape
}

func NewContainer() *Container { return &Container{} }

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape(col, row int64) *_g.Shape {
	_ga := _g.NewShape()
	_ga.IdAttr = _bf.String(_b.Sprintf("\u0063\u0073\u005f\u0025\u0064\u005f\u0025\u0064", col, row))
	_ga.TypeAttr = _bf.String("\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_ga.StyleAttr = _bf.String("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006cu\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074:\u0038\u0030\u0070\u0074;\u006d\u0061\u0072\u0067\u0069n-\u0074o\u0070\u003a\u0032pt\u003b\u0077\u0069\u0064\u0074\u0068\u003a1\u0030\u0034\u0070\u0074\u003b\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0037\u0036\u0070\u0074\u003b\u007a\u002d\u0069\u006e\u0064\u0065x\u003a\u0031\u003bv\u0069\u0073\u0069\u0062\u0069\u006c\u0069t\u0079\u003a\u0068\u0069\u0064\u0064\u0065\u006e")
	_ga.FillcolorAttr = _bf.String("\u0023f\u0062\u0066\u0036\u0064\u0036")
	_ga.StrokecolorAttr = _bf.String("\u0023e\u0064\u0065\u0061\u0061\u0031")
	_e := _g.NewEG_ShapeElements()
	_e.Fill = _g.NewFill()
	_e.Fill.Color2Attr = _bf.String("\u0023f\u0062\u0066\u0065\u0038\u0032")
	_e.Fill.AngleAttr = _bf.Float64(-180)
	_e.Fill.TypeAttr = _g.ST_FillTypeGradient
	_e.Fill.Fill = _g.NewOfcFill()
	_e.Fill.Fill.ExtAttr = _g.ST_ExtView
	_e.Fill.Fill.TypeAttr = _g.OfcST_FillTypeGradientUnscaled
	_ga.EG_ShapeElements = append(_ga.EG_ShapeElements, _e)
	_ge := _g.NewEG_ShapeElements()
	_ge.Shadow = _g.NewShadow()
	_ge.Shadow.OnAttr = _fe.ST_TrueFalseT
	_ge.Shadow.ObscuredAttr = _fe.ST_TrueFalseT
	_ga.EG_ShapeElements = append(_ga.EG_ShapeElements, _ge)
	_c := _g.NewEG_ShapeElements()
	_c.Path = _g.NewPath()
	_c.Path.ConnecttypeAttr = _g.OfcST_ConnectTypeNone
	_ga.EG_ShapeElements = append(_ga.EG_ShapeElements, _c)
	_ef := _g.NewEG_ShapeElements()
	_ef.Textbox = _g.NewTextbox()
	_ef.Textbox.StyleAttr = _bf.String("\u006d\u0073\u006f\u002ddi\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u002d\u0061\u006c\u0074\u003a\u0061\u0075t\u006f")
	_ga.EG_ShapeElements = append(_ga.EG_ShapeElements, _ef)
	_gb := _g.NewEG_ShapeElements()
	_gb.ClientData = _ba.NewClientData()
	_gb.ClientData.ObjectTypeAttr = _ba.ST_ObjectTypeNote
	_gb.ClientData.MoveWithCells = _fe.ST_TrueFalseBlankT
	_gb.ClientData.SizeWithCells = _fe.ST_TrueFalseBlankT
	_gb.ClientData.Anchor = _bf.String("\u0031,\u0020\u0031\u0035\u002c\u0020\u0030\u002c\u0020\u0032\u002c\u00202\u002c\u0020\u0035\u0034\u002c\u0020\u0035\u002c\u0020\u0033")
	_gb.ClientData.AutoFill = _fe.ST_TrueFalseBlankFalse
	_gb.ClientData.Row = _bf.Int64(row)
	_gb.ClientData.Column = _bf.Int64(col)
	_ga.EG_ShapeElements = append(_ga.EG_ShapeElements, _gb)
	return _ga
}
func (_d *Container) MarshalXML(e *_fd.Encoder, start _fd.StartElement) error {
	start.Attr = append(start.Attr, _fd.Attr{Name: _fd.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0076"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d:v\u006d\u006c"})
	start.Attr = append(start.Attr, _fd.Attr{Name: _fd.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006f"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006di\u0063\u0072\u006f\u0073\u006f\u0066t\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u006ff\u0066\u0069\u0063\u0065"})
	start.Attr = append(start.Attr, _fd.Attr{Name: _fd.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Name.Local = "\u0078\u006d\u006c"
	e.EncodeToken(start)
	if _d.Layout != nil {
		_ed := _fd.StartElement{Name: _fd.Name{Local: "\u006f\u003a\u0073\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074"}}
		e.EncodeElement(_d.Layout, _ed)
	}
	if _d.ShapeType != nil {
		_bc := _fd.StartElement{Name: _fd.Name{Local: "v\u003a\u0073\u0068\u0061\u0070\u0065\u0074\u0079\u0070\u0065"}}
		e.EncodeElement(_d.ShapeType, _bc)
	}
	for _, _fda := range _d.Shape {
		_gaa := _fd.StartElement{Name: _fd.Name{Local: "\u0076:\u0073\u0068\u0061\u0070\u0065"}}
		e.EncodeElement(_fda, _gaa)
	}
	return e.EncodeToken(_fd.EndElement{Name: start.Name})
}

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing() *Container {
	_fc := NewContainer()
	_fc.Layout = _g.NewOfcShapelayout()
	_fc.Layout.ExtAttr = _g.ST_ExtEdit
	_fc.Layout.Idmap = _g.NewOfcCT_IdMap()
	_fc.Layout.Idmap.DataAttr = _bf.String("\u0031")
	_fc.Layout.Idmap.ExtAttr = _g.ST_ExtEdit
	_fc.ShapeType = _g.NewShapetype()
	_fc.ShapeType.IdAttr = _bf.String("_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_fc.ShapeType.CoordsizeAttr = _bf.String("2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030")
	_fc.ShapeType.SptAttr = _bf.Float32(202)
	_fc.ShapeType.PathAttr = _bf.String("\u006d\u0030\u002c0l\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u00321\u00360\u0030,\u00321\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u0030\u0078\u0065")
	_a := _g.NewEG_ShapeElements()
	_fc.ShapeType.EG_ShapeElements = append(_fc.ShapeType.EG_ShapeElements, _a)
	_a.Path = _g.NewPath()
	_a.Path.GradientshapeokAttr = _fe.ST_TrueFalseT
	_a.Path.ConnecttypeAttr = _g.OfcST_ConnectTypeRect
	return _fc
}
func (_be *Container) UnmarshalXML(d *_fd.Decoder, start _fd.StartElement) error {
	_be.Shape = nil
_eg:
	for {
		_ede, _edec := d.Token()
		if _edec != nil {
			return _edec
		}
		switch _cc := _ede.(type) {
		case _fd.StartElement:
			switch _cc.Name.Local {
			case "s\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074":
				_be.Layout = _g.NewOfcShapelayout()
				if _cf := d.DecodeElement(_be.Layout, &_cc); _cf != nil {
					return _cf
				}
			case "\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e":
				_be.ShapeType = _g.NewShapetype()
				if _eb := d.DecodeElement(_be.ShapeType, &_cc); _eb != nil {
					return _eb
				}
			case "\u0073\u0068\u0061p\u0065":
				_ce := _g.NewShape()
				if _geg := d.DecodeElement(_ce, &_cc); _geg != nil {
					return _geg
				}
				_be.Shape = append(_be.Shape, _ce)
			}
		case _fd.EndElement:
			break _eg
		}
	}
	return nil
}
