package drawing

import (
	_d "gitee.com/gooffice/gooffice"
	_a "gitee.com/gooffice/gooffice/color"
	_g "gitee.com/gooffice/gooffice/measurement"
	_c "gitee.com/gooffice/gooffice/schema/soo/dml"
)

// Properties returns the run's properties.
func (_bab Run) Properties() RunProperties {
	if _bab._ff.R == nil {
		_bab._ff.R = _c.NewCT_RegularTextRun()
	}
	if _bab._ff.R.RPr == nil {
		_bab._ff.R.RPr = _c.NewCT_TextCharacterProperties()
	}
	return RunProperties{_bab._ff.R.RPr}
}

type LineProperties struct{ _cb *_c.CT_LineProperties }

// LineJoin is the type of line join
type LineJoin byte

func (_ab LineProperties) SetSolidFill(c _a.Color) {
	_ab.clearFill()
	_ab._cb.SolidFill = _c.NewCT_SolidColorFillProperties()
	_ab._cb.SolidFill.SrgbClr = _c.NewCT_SRgbColor()
	_ab._cb.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}
func (_ag LineProperties) SetNoFill() { _ag.clearFill(); _ag._cb.NoFill = _c.NewCT_NoFillProperties() }

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
	LineJoinMiter
)

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct {
	_cc *_c.CT_TextParagraphProperties
}

// Properties returns the paragraph properties.
func (_fd Paragraph) Properties() ParagraphProperties {
	if _fd._da.PPr == nil {
		_fd._da.PPr = _c.NewCT_TextParagraphProperties()
	}
	return MakeParagraphProperties(_fd._da.PPr)
}

// AddBreak adds a new line break to a paragraph.
func (_dfb Paragraph) AddBreak() {
	_dd := _c.NewEG_TextRun()
	_dd.Br = _c.NewCT_TextLineBreak()
	_dfb._da.EG_TextRun = append(_dfb._da.EG_TextRun, _dd)
}

// SetBulletFont controls the font for the bullet character.
func (_geb ParagraphProperties) SetBulletFont(f string) {
	if f == "" {
		_geb._cc.BuFont = nil
	} else {
		_geb._cc.BuFont = _c.NewCT_TextFont()
		_geb._cc.BuFont.TypefaceAttr = f
	}
}

// SetFlipVertical controls if the shape is flipped vertically.
func (_ged ShapeProperties) SetFlipVertical(b bool) {
	_ged.ensureXfrm()
	if !b {
		_ged._ddg.Xfrm.FlipVAttr = nil
	} else {
		_ged._ddg.Xfrm.FlipVAttr = _d.Bool(true)
	}
}

// RunProperties controls the run properties.
type RunProperties struct {
	_fg *_c.CT_TextCharacterProperties
}

// SetFont controls the font of a run.
func (_fbf RunProperties) SetFont(s string) {
	_fbf._fg.Latin = _c.NewCT_TextFont()
	_fbf._fg.Latin.TypefaceAttr = s
}

// X returns the inner wrapped XML type.
func (_ge Paragraph) X() *_c.CT_TextParagraph { return _ge._da }

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_ga LineProperties) SetWidth(w _g.Distance) { _ga._cb.WAttr = _d.Int32(int32(w / _g.EMU)) }
func (_e LineProperties) clearFill() {
	_e._cb.NoFill = nil
	_e._cb.GradFill = nil
	_e._cb.SolidFill = nil
	_e._cb.PattFill = nil
}

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_dbc ShapeProperties) SetFlipHorizontal(b bool) {
	_dbc.ensureXfrm()
	if !b {
		_dbc._ddg.Xfrm.FlipHAttr = nil
	} else {
		_dbc._ddg.Xfrm.FlipHAttr = _d.Bool(true)
	}
}

// SetSize sets the font size of the run text
func (_eb RunProperties) SetSize(sz _g.Distance) {
	_eb._fg.SzAttr = _d.Int32(int32(sz / _g.HundredthPoint))
}
func MakeShapeProperties(x *_c.CT_ShapeProperties) ShapeProperties { return ShapeProperties{x} }

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties(x *_c.CT_TextCharacterProperties) RunProperties { return RunProperties{x} }

// SetWidth sets the width of the shape.
func (_fc ShapeProperties) SetWidth(w _g.Distance) {
	_fc.ensureXfrm()
	if _fc._ddg.Xfrm.Ext == nil {
		_fc._ddg.Xfrm.Ext = _c.NewCT_PositiveSize2D()
	}
	_fc._ddg.Xfrm.Ext.CxAttr = int64(w / _g.EMU)
}

// SetPosition sets the position of the shape.
func (_gce ShapeProperties) SetPosition(x, y _g.Distance) {
	_gce.ensureXfrm()
	if _gce._ddg.Xfrm.Off == nil {
		_gce._ddg.Xfrm.Off = _c.NewCT_Point2D()
	}
	_gce._ddg.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _d.Int64(int64(x / _g.EMU))
	_gce._ddg.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _d.Int64(int64(y / _g.EMU))
}
func (_ad ShapeProperties) SetSolidFill(c _a.Color) {
	_ad.clearFill()
	_ad._ddg.SolidFill = _c.NewCT_SolidColorFillProperties()
	_ad._ddg.SolidFill.SrgbClr = _c.NewCT_SRgbColor()
	_ad._ddg.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// MakeRun constructs a new Run wrapper.
func MakeRun(x *_c.EG_TextRun) Run { return Run{x} }

// SetGeometry sets the shape type of the shape
func (_cbg ShapeProperties) SetGeometry(g _c.ST_ShapeType) {
	if _cbg._ddg.PrstGeom == nil {
		_cbg._ddg.PrstGeom = _c.NewCT_PresetGeometry2D()
	}
	_cbg._ddg.PrstGeom.PrstAttr = g
}

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph(x *_c.CT_TextParagraph) Paragraph { return Paragraph{x} }

// X returns the inner wrapped XML type.
func (_ee ShapeProperties) X() *_c.CT_ShapeProperties { return _ee._ddg }
func (_ec ShapeProperties) LineProperties() LineProperties {
	if _ec._ddg.Ln == nil {
		_ec._ddg.Ln = _c.NewCT_LineProperties()
	}
	return LineProperties{_ec._ddg.Ln}
}

// SetText sets the run's text contents.
func (_ffb Run) SetText(s string) {
	_ffb._ff.Br = nil
	_ffb._ff.Fld = nil
	if _ffb._ff.R == nil {
		_ffb._ff.R = _c.NewCT_RegularTextRun()
	}
	_ffb._ff.R.T = s
}
func (_db ShapeProperties) clearFill() {
	_db._ddg.NoFill = nil
	_db._ddg.BlipFill = nil
	_db._ddg.GradFill = nil
	_db._ddg.GrpFill = nil
	_db._ddg.SolidFill = nil
	_db._ddg.PattFill = nil
}

// AddRun adds a new run to a paragraph.
func (_f Paragraph) AddRun() Run {
	_gaf := MakeRun(_c.NewEG_TextRun())
	_f._da.EG_TextRun = append(_f._da.EG_TextRun, _gaf.X())
	return _gaf
}

// X returns the inner wrapped XML type.
func (_ba Run) X() *_c.EG_TextRun { return _ba._ff }

// X returns the inner wrapped XML type.
func (_ef ParagraphProperties) X() *_c.CT_TextParagraphProperties { return _ef._cc }

// SetNumbered controls if bullets are numbered or not.
func (_fb ParagraphProperties) SetNumbered(scheme _c.ST_TextAutonumberScheme) {
	if scheme == _c.ST_TextAutonumberSchemeUnset {
		_fb._cc.BuAutoNum = nil
	} else {
		_fb._cc.BuAutoNum = _c.NewCT_TextAutonumberBullet()
		_fb._cc.BuAutoNum.TypeAttr = scheme
	}
}

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties(x *_c.CT_TextParagraphProperties) ParagraphProperties {
	return ParagraphProperties{x}
}

// SetSize sets the width and height of the shape.
func (_ce ShapeProperties) SetSize(w, h _g.Distance) { _ce.SetWidth(w); _ce.SetHeight(h) }

// SetJoin sets the line join style.
func (_bd LineProperties) SetJoin(e LineJoin) {
	_bd._cb.Round = nil
	_bd._cb.Miter = nil
	_bd._cb.Bevel = nil
	switch e {
	case LineJoinRound:
		_bd._cb.Round = _c.NewCT_LineJoinRound()
	case LineJoinBevel:
		_bd._cb.Bevel = _c.NewCT_LineJoinBevel()
	case LineJoinMiter:
		_bd._cb.Miter = _c.NewCT_LineJoinMiterProperties()
	}
}

// SetSolidFill controls the text color of a run.
func (_gca RunProperties) SetSolidFill(c _a.Color) {
	_gca._fg.NoFill = nil
	_gca._fg.BlipFill = nil
	_gca._fg.GradFill = nil
	_gca._fg.GrpFill = nil
	_gca._fg.PattFill = nil
	_gca._fg.SolidFill = _c.NewCT_SolidColorFillProperties()
	_gca._fg.SolidFill.SrgbClr = _c.NewCT_SRgbColor()
	_gca._fg.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// SetBold controls the bolding of a run.
func (_gc RunProperties) SetBold(b bool) { _gc._fg.BAttr = _d.Bool(b) }

// Run is a run within a paragraph.
type Run struct{ _ff *_c.EG_TextRun }
type ShapeProperties struct{ _ddg *_c.CT_ShapeProperties }

// SetLevel sets the level of indentation of a paragraph.
func (_bb ParagraphProperties) SetLevel(idx int32) { _bb._cc.LvlAttr = _d.Int32(idx) }

// SetAlign controls the paragraph alignment
func (_abd ParagraphProperties) SetAlign(a _c.ST_TextAlignType) { _abd._cc.AlgnAttr = a }

// X returns the inner wrapped XML type.
func (_df LineProperties) X() *_c.CT_LineProperties { return _df._cb }
func (_ac ShapeProperties) SetNoFill() {
	_ac.clearFill()
	_ac._ddg.NoFill = _c.NewCT_NoFillProperties()
}
func (_cf ShapeProperties) ensureXfrm() {
	if _cf._ddg.Xfrm == nil {
		_cf._ddg.Xfrm = _c.NewCT_Transform2D()
	}
}

// SetHeight sets the height of the shape.
func (_fbb ShapeProperties) SetHeight(h _g.Distance) {
	_fbb.ensureXfrm()
	if _fbb._ddg.Xfrm.Ext == nil {
		_fbb._ddg.Xfrm.Ext = _c.NewCT_PositiveSize2D()
	}
	_fbb._ddg.Xfrm.Ext.CyAttr = int64(h / _g.EMU)
}

// SetBulletChar sets the bullet character for the paragraph.
func (_gg ParagraphProperties) SetBulletChar(c string) {
	if c == "" {
		_gg._cc.BuChar = nil
	} else {
		_gg._cc.BuChar = _c.NewCT_TextCharBullet()
		_gg._cc.BuChar.CharAttr = c
	}
}

// Paragraph is a paragraph within a document.
type Paragraph struct{ _da *_c.CT_TextParagraph }
