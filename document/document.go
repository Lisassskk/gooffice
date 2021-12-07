package document

import (
	_da "archive/zip"
	_cb "bytes"
	_ffa "errors"
	_eg "fmt"
	_d "gitee.com/gooffice/gooffice"
	_bc "gitee.com/gooffice/gooffice/color"
	_ea "gitee.com/gooffice/gooffice/common"
	_dcd "gitee.com/gooffice/gooffice/common/logger"
	_gb "gitee.com/gooffice/gooffice/common/tempstorage"
	_b "gitee.com/gooffice/gooffice/internal/license"
	_eb "gitee.com/gooffice/gooffice/measurement"
	_gf "gitee.com/gooffice/gooffice/schema/schemas.microsoft.com/office/activeX"
	_bb "gitee.com/gooffice/gooffice/schema/soo/dml"
	_dg "gitee.com/gooffice/gooffice/schema/soo/dml/chart"
	_gd "gitee.com/gooffice/gooffice/schema/soo/dml/picture"
	_cg "gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes"
	_eab "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships"
	_eaba "gitee.com/gooffice/gooffice/schema/soo/wml"
	_aae "gitee.com/gooffice/gooffice/schema/urn/schemas_microsoft_com/vml"
	_ffc "gitee.com/gooffice/gooffice/vmldrawing"
	_fd "gitee.com/gooffice/gooffice/zippkg"
	_ac "image"
	_e "image/jpeg"
	_aa "io"
	_a "math/rand"
	_ff "os"
	_c "path/filepath"
	_dc "strconv"
	_f "strings"
	_ae "unicode"
)

// X returns the inner wrapped XML type.
func (_aacf Header) X() *_eaba.Hdr { return _aacf._cbgbf }

// StructuredDocumentTags returns the structured document tags in the document
// which are commonly used in document templates.
func (_ddb *Document) StructuredDocumentTags() []StructuredDocumentTag {
	_ceeg := []StructuredDocumentTag{}
	for _, _gdfe := range _ddb._eeda.Body.EG_BlockLevelElts {
		for _, _bfbg := range _gdfe.EG_ContentBlockContent {
			if _bfbg.Sdt != nil {
				_ceeg = append(_ceeg, StructuredDocumentTag{_ddb, _bfbg.Sdt})
			}
		}
	}
	return _ceeg
}

// Borders allows manipulation of the table borders.
func (_acbdg TableStyleProperties) Borders() TableBorders {
	if _acbdg._ddafc.TblBorders == nil {
		_acbdg._ddafc.TblBorders = _eaba.NewCT_TblBorders()
	}
	return TableBorders{_acbdg._ddafc.TblBorders}
}

// ExtractFromFooter returns text from the document footer as an array of TextItems.
func ExtractFromFooter(footer *_eaba.Ftr) []TextItem {
	return _bddda(footer.EG_ContentBlockContent, nil)
}

// SetAlignment controls the paragraph alignment
func (_ebac ParagraphProperties) SetAlignment(align _eaba.ST_Jc) {
	if align == _eaba.ST_JcUnset {
		_ebac._efec.Jc = nil
	} else {
		_ebac._efec.Jc = _eaba.NewCT_Jc()
		_ebac._efec.Jc.ValAttr = align
	}
}

// FindNodeByStyleId return slice of node base on style id.
func (_dbaf *Nodes) FindNodeByStyleId(styleId string) []Node {
	_dfbe := []Node{}
	for _, _aaag := range _dbaf._agcab {
		switch _efbc := _aaag._efag.(type) {
		case *Paragraph:
			if _efbc != nil && _efbc.Style() == styleId {
				_dfbe = append(_dfbe, _aaag)
			}
		case *Table:
			if _efbc != nil && _efbc.Style() == styleId {
				_dfbe = append(_dfbe, _aaag)
			}
		}
		_gfcba := Nodes{_agcab: _aaag.Children}
		_dfbe = append(_dfbe, _gfcba.FindNodeByStyleId(styleId)...)
	}
	return _dfbe
}

// RightToLeft returns true if paragraph text goes from right to left.
func (_fgeb ParagraphProperties) RightToLeft() bool { return _beff(_fgeb._efec.RPr.Rtl) }

// AddTabStop adds a tab stop to the paragraph.
func (_agaf ParagraphStyleProperties) AddTabStop(position _eb.Distance, justificaton _eaba.ST_TabJc, leader _eaba.ST_TabTlc) {
	if _agaf._fded.Tabs == nil {
		_agaf._fded.Tabs = _eaba.NewCT_Tabs()
	}
	_beea := _eaba.NewCT_TabStop()
	_beea.LeaderAttr = leader
	_beea.ValAttr = justificaton
	_beea.PosAttr.Int64 = _d.Int64(int64(position / _eb.Twips))
	_agaf._fded.Tabs.Tab = append(_agaf._fded.Tabs.Tab, _beea)
}

// AddCheckBox adds checkbox form field to the paragraph and returns it.
func (_edgg Paragraph) AddCheckBox(name string) FormField {
	_agagd := _edgg.addFldCharsForField(name, "\u0046\u004f\u0052M\u0043\u0048\u0045\u0043\u004b\u0042\u004f\u0058")
	_agagd._cdag.CheckBox = _eaba.NewCT_FFCheckBox()
	return _agagd
}

// FormField is a form within a document. It references the document, so changes
// to the form field wil be reflected in the document if it is saved.
type FormField struct {
	_cdag *_eaba.CT_FFData
	_bebe *_eaba.EG_RunInnerContent
}

// IsEndnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (_fgedc Run) IsEndnote() (bool, int64) {
	if _fgedc._bggd.EG_RunInnerContent != nil {
		if _fgedc._bggd.EG_RunInnerContent[0].EndnoteReference != nil {
			return true, _fgedc._bggd.EG_RunInnerContent[0].EndnoteReference.IdAttr
		}
	}
	return false, 0
}

// Numbering return numbering that being use by paragraph.
func (_deeg Paragraph) Numbering() Numbering {
	_deeg.ensurePPr()
	_aede := NewNumbering()
	if _deeg._dbef.PPr.NumPr != nil {
		_cebbef := int64(-1)
		_eded := int64(-1)
		if _deeg._dbef.PPr.NumPr.NumId != nil {
			_cebbef = _deeg._dbef.PPr.NumPr.NumId.ValAttr
		}
		for _, _edee := range _deeg._bbgc.Numbering._ebbdb.Num {
			if _cebbef < 0 {
				break
			}
			if _edee.NumIdAttr == _cebbef {
				if _edee.AbstractNumId != nil {
					_eded = _edee.AbstractNumId.ValAttr
					_aede._ebbdb.Num = append(_aede._ebbdb.Num, _edee)
					break
				}
			}
		}
		for _, _gaeb := range _deeg._bbgc.Numbering._ebbdb.AbstractNum {
			if _eded < 0 {
				break
			}
			if _gaeb.AbstractNumIdAttr == _eded {
				_aede._ebbdb.AbstractNum = append(_aede._ebbdb.AbstractNum, _gaeb)
				break
			}
		}
	}
	return _aede
}

// RemoveParagraph removes a paragraph from a footer.
func (_caae Footer) RemoveParagraph(p Paragraph) {
	for _, _adgf := range _caae._daaa.EG_ContentBlockContent {
		for _ecda, _aeedd := range _adgf.P {
			if _aeedd == p._dbef {
				copy(_adgf.P[_ecda:], _adgf.P[_ecda+1:])
				_adgf.P = _adgf.P[0 : len(_adgf.P)-1]
				return
			}
		}
	}
}

// SetBottomPct sets the cell bottom margin
func (_fe CellMargins) SetBottomPct(pct float64) {
	_fe._ec.Bottom = _eaba.NewCT_TblWidth()
	_cab(_fe._ec.Bottom, pct)
}

// NewWatermarkPicture generates new WatermarkPicture.
func NewWatermarkPicture() WatermarkPicture {
	_ggbd := _aae.NewShapetype()
	_ccdc := _aae.NewEG_ShapeElements()
	_ccdc.Formulas = _edfg()
	_ccdc.Path = _adaaa()
	_ccdc.Lock = _bbbb()
	_ggbd.EG_ShapeElements = []*_aae.EG_ShapeElements{_ccdc}
	var (
		_gebd = "\u005f\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0037\u0035"
		_dbfa = "2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030"
		_gced = float32(75.0)
		_aefa = "\u006d\u0040\u0034\u00405l\u0040\u0034\u0040\u0031\u0031\u0040\u0039\u0040\u0031\u0031\u0040\u0039\u0040\u0035x\u0065"
	)
	_ggbd.IdAttr = &_gebd
	_ggbd.CoordsizeAttr = &_dbfa
	_ggbd.SptAttr = &_gced
	_ggbd.PreferrelativeAttr = _cg.ST_TrueFalseTrue
	_ggbd.PathAttr = &_aefa
	_ggbd.FilledAttr = _cg.ST_TrueFalseFalse
	_ggbd.StrokedAttr = _cg.ST_TrueFalseFalse
	_bcagd := _aae.NewShape()
	_gfcab := _aae.NewEG_ShapeElements()
	_gfcab.Imagedata = _fcgde()
	_bcagd.EG_ShapeElements = []*_aae.EG_ShapeElements{_gfcab}
	var (
		_fgcf  = "\u0057\u006f\u0072\u0064\u0050\u0069\u0063\u0074\u0075\u0072e\u0057\u0061\u0074\u0065\u0072\u006d\u0061r\u006b\u0031\u0036\u0033\u0032\u0033\u0031\u0036\u0035\u0039\u0035"
		_ecccg = "\u005f\u0078\u00300\u0030\u0030\u005f\u0073\u0032\u0030\u0035\u0031"
		_aacc  = "#\u005f\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0037\u0035"
		_cecab = ""
		_agfae = "\u0070os\u0069t\u0069o\u006e\u003a\u0061\u0062\u0073\u006fl\u0075\u0074\u0065\u003bm\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074\u003a\u0030\u003bma\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070\u003a\u0030\u003b\u0077\u0069\u0064\u0074\u0068\u003a\u0030\u0070\u0074;\u0068e\u0069\u0067\u0068\u0074\u003a\u0030\u0070\u0074\u003b\u007a\u002d\u0069\u006ed\u0065\u0078:\u002d\u0032\u00351\u0036\u0035\u0038\u0032\u0034\u0030\u003b\u006d\u0073o-\u0070\u006f\u0073i\u0074\u0069\u006f\u006e-\u0068\u006f\u0072\u0069\u007a\u006fn\u0074\u0061l\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003bm\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u002drela\u0074\u0069\u0076\u0065\u003a\u006d\u0061\u0072\u0067\u0069\u006e\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076\u0065\u0072t\u0069c\u0061l\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e-\u0076\u0065r\u0074\u0069c\u0061l\u002d\u0072\u0065\u006c\u0061\u0074i\u0076\u0065\u003a\u006d\u0061\u0072\u0067\u0069\u006e"
	)
	_bcagd.IdAttr = &_fgcf
	_bcagd.SpidAttr = &_ecccg
	_bcagd.TypeAttr = &_aacc
	_bcagd.AltAttr = &_cecab
	_bcagd.StyleAttr = &_agfae
	_bcagd.AllowincellAttr = _cg.ST_TrueFalseFalse
	_dcbfg := _eaba.NewCT_Picture()
	_dcbfg.Any = []_d.Any{_ggbd, _bcagd}
	return WatermarkPicture{_dadb: _dcbfg, _ggac: _bcagd, _bfag: _ggbd}
}

// X returns the inner wrapped XML type.
func (_efbb Fonts) X() *_eaba.CT_Fonts { return _efbb._aggeg }

// SetPicture sets the watermark picture.
func (_ecdd *WatermarkPicture) SetPicture(imageRef _ea.ImageRef) {
	_dfefd := imageRef.RelID()
	_dccc := _ecdd.getShape()
	if _ecdd._ggac != nil {
		_cfbag := _ecdd._ggac.EG_ShapeElements
		if len(_cfbag) > 0 && _cfbag[0].Imagedata != nil {
			_cfbag[0].Imagedata.IdAttr = &_dfefd
		}
	} else {
		_bgfgc := _ecdd.findNode(_dccc, "\u0069m\u0061\u0067\u0065\u0064\u0061\u0074a")
		for _gbadb, _gbaa := range _bgfgc.Attrs {
			if _gbaa.Name.Local == "\u0069\u0064" {
				_bgfgc.Attrs[_gbadb].Value = _dfefd
			}
		}
	}
}

// SetAll sets all of the borders to a given value.
func (_bcb CellBorders) SetAll(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_bcb.SetBottom(t, c, thickness)
	_bcb.SetLeft(t, c, thickness)
	_bcb.SetRight(t, c, thickness)
	_bcb.SetTop(t, c, thickness)
	_bcb.SetInsideHorizontal(t, c, thickness)
	_bcb.SetInsideVertical(t, c, thickness)
}

// Cells returns the cells defined in the table.
func (_acbg Row) Cells() []Cell {
	_gddga := []Cell{}
	for _, _ffdbc := range _acbg._cegac.EG_ContentCellContent {
		for _, _acec := range _ffdbc.Tc {
			_gddga = append(_gddga, Cell{_acbg._eegcb, _acec})
		}
		if _ffdbc.Sdt != nil && _ffdbc.Sdt.SdtContent != nil {
			for _, _fafc := range _ffdbc.Sdt.SdtContent.Tc {
				_gddga = append(_gddga, Cell{_acbg._eegcb, _fafc})
			}
		}
	}
	return _gddga
}

const _abgga = "\u0046\u006f\u0072\u006d\u0046\u0069\u0065l\u0064\u0054\u0079\u0070\u0065\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0046\u006fr\u006dF\u0069\u0065\u006c\u0064\u0054\u0079p\u0065\u0054\u0065\u0078\u0074\u0046\u006fr\u006d\u0046\u0069\u0065\u006c\u0064\u0054\u0079\u0070\u0065\u0043\u0068\u0065\u0063\u006b\u0042\u006f\u0078\u0046\u006f\u0072\u006d\u0046i\u0065\u006c\u0064\u0054\u0079\u0070\u0065\u0044\u0072\u006f\u0070\u0044\u006fw\u006e"

// VerticalAlign returns the value of run vertical align.
func (_defe RunProperties) VerticalAlignment() _cg.ST_VerticalAlignRun {
	if _ggaa := _defe._gfbgc.VertAlign; _ggaa != nil {
		return _ggaa.ValAttr
	}
	return 0
}

// SetTargetBookmark sets the bookmark target of the hyperlink.
func (_ffdb HyperLink) SetTargetBookmark(bm Bookmark) {
	_ffdb._dgace.AnchorAttr = _d.String(bm.Name())
	_ffdb._dgace.IdAttr = nil
}

// Footers returns the footers defined in the document.
func (_dgb *Document) Footers() []Footer {
	_bbdb := []Footer{}
	for _, _afc := range _dgb._gfdb {
		_bbdb = append(_bbdb, Footer{_dgb, _afc})
	}
	return _bbdb
}

// // SetBeforeLineSpacing sets spacing above paragraph in line units.
func (_cgabe Paragraph) SetBeforeLineSpacing(d _eb.Distance) {
	_cgabe.ensurePPr()
	if _cgabe._dbef.PPr.Spacing == nil {
		_cgabe._dbef.PPr.Spacing = _eaba.NewCT_Spacing()
	}
	_bffg := _cgabe._dbef.PPr.Spacing
	_bffg.BeforeLinesAttr = _d.Int64(int64(d / _eb.Twips))
}

func _dcadc(_dgbf *_eaba.CT_Tbl, _aabe map[string]string) {
	for _, _cabb := range _dgbf.EG_ContentRowContent {
		for _, _gdbe := range _cabb.Tr {
			for _, _cdbg := range _gdbe.EG_ContentCellContent {
				for _, _agge := range _cdbg.Tc {
					for _, _bcad := range _agge.EG_BlockLevelElts {
						for _, _fdba := range _bcad.EG_ContentBlockContent {
							for _, _caedb := range _fdba.P {
								_fffb(_caedb, _aabe)
							}
							for _, _caaa := range _fdba.Tbl {
								_dcadc(_caaa, _aabe)
							}
						}
					}
				}
			}
		}
	}
}

// GetSize return the size of anchor on the page.
func (_ee AnchoredDrawing) GetSize() (_gbe, _eed int64) {
	return _ee._cc.Extent.CxAttr, _ee._cc.Extent.CyAttr
}

func _bbbef(_ecfgf []*_eaba.CT_P, _deaf *TableInfo, _ecaec *DrawingInfo) []TextItem {
	_egdc := []TextItem{}
	for _, _bdcg := range _ecfgf {
		_egdc = append(_egdc, _dfb(_bdcg, nil, _deaf, _ecaec, _bdcg.EG_PContent)...)
	}
	return _egdc
}

// SetTextWrapInFrontOfText sets the text wrap to in front of text.
func (_de AnchoredDrawing) SetTextWrapInFrontOfText() {
	_de._cc.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_de._cc.Choice.WrapNone = _eaba.NewWdCT_WrapNone()
	_de._cc.BehindDocAttr = false
	_de._cc.LayoutInCellAttr = true
	_de._cc.AllowOverlapAttr = true
}

// Name returns the name of the bookmark whcih is the document unique ID that
// identifies the bookmark.
func (_ccf Bookmark) Name() string { return _ccf._dda.NameAttr }

// ParagraphProperties returns the paragraph properties controlling text formatting within the table.
func (_bcfgef TableConditionalFormatting) ParagraphProperties() ParagraphStyleProperties {
	if _bcfgef._agcba.PPr == nil {
		_bcfgef._agcba.PPr = _eaba.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_bcfgef._agcba.PPr}
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (_aaea ParagraphProperties) SetKeepOnOnePage(b bool) {
	if !b {
		_aaea._efec.KeepLines = nil
	} else {
		_aaea._efec.KeepLines = _eaba.NewCT_OnOff()
	}
}

func (_adgce *WatermarkPicture) getShapeImagedata() *_d.XSDAny {
	return _adgce.getInnerElement("\u0069m\u0061\u0067\u0065\u0064\u0061\u0074a")
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (_bdg CellBorders) SetInsideVertical(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_bdg._ggf.InsideV = _eaba.NewCT_Border()
	_gfdf(_bdg._ggf.InsideV, t, c, thickness)
}

// AddTable adds a new table to the document body.
func (_edc *Document) AddTable() Table {
	_cgcg := _eaba.NewEG_BlockLevelElts()
	_edc._eeda.Body.EG_BlockLevelElts = append(_edc._eeda.Body.EG_BlockLevelElts, _cgcg)
	_gbg := _eaba.NewEG_ContentBlockContent()
	_cgcg.EG_ContentBlockContent = append(_cgcg.EG_ContentBlockContent, _gbg)
	_cfc := _eaba.NewCT_Tbl()
	_gbg.Tbl = append(_gbg.Tbl, _cfc)
	return Table{_edc, _cfc}
}

// IsFootnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (_acefc Run) IsFootnote() (bool, int64) {
	if _acefc._bggd.EG_RunInnerContent != nil {
		if _acefc._bggd.EG_RunInnerContent[0].FootnoteReference != nil {
			return true, _acefc._bggd.EG_RunInnerContent[0].FootnoteReference.IdAttr
		}
	}
	return false, 0
}

// RunProperties returns the run style properties.
func (_cffcf Style) RunProperties() RunProperties {
	if _cffcf._bgfca.RPr == nil {
		_cffcf._bgfca.RPr = _eaba.NewCT_RPr()
	}
	return RunProperties{_cffcf._bgfca.RPr}
}

// X returns the inner wrapped XML type.
func (_bfdd Styles) X() *_eaba.Styles { return _bfdd._abad }

func (_eabgb Paragraph) addFldChar() *_eaba.CT_FldChar {
	_afbbf := _eabgb.AddRun()
	_ceab := _afbbf.X()
	_cffg := _eaba.NewEG_RunInnerContent()
	_gdfg := _eaba.NewCT_FldChar()
	_cffg.FldChar = _gdfg
	_ceab.EG_RunInnerContent = append(_ceab.EG_RunInnerContent, _cffg)
	return _gdfg
}

// RemoveParagraph removes a paragraph from the footnote.
func (_eaaece Footnote) RemoveParagraph(p Paragraph) {
	for _, _gaag := range _eaaece.content() {
		for _deabf, _cebbd := range _gaag.P {
			if _cebbd == p._dbef {
				copy(_gaag.P[_deabf:], _gaag.P[_deabf+1:])
				_gaag.P = _gaag.P[0 : len(_gaag.P)-1]
				return
			}
		}
	}
}

// AddParagraph adds a paragraph to the footer.
func (_egef Footer) AddParagraph() Paragraph {
	_fade := _eaba.NewEG_ContentBlockContent()
	_egef._daaa.EG_ContentBlockContent = append(_egef._daaa.EG_ContentBlockContent, _fade)
	_adfa := _eaba.NewCT_P()
	_fade.P = append(_fade.P, _adfa)
	return Paragraph{_egef._dec, _adfa}
}

// SetSemiHidden controls if the style is hidden in the UI.
func (_eaeee Style) SetSemiHidden(b bool) {
	if b {
		_eaeee._bgfca.SemiHidden = _eaba.NewCT_OnOff()
	} else {
		_eaeee._bgfca.SemiHidden = nil
	}
}

func _egacb() *_aae.Handles {
	_cgbfgc := _aae.NewHandles()
	_gfbd := _aae.NewCT_H()
	_aacg := "\u0023\u0030\u002c\u0062\u006f\u0074\u0074\u006f\u006dR\u0069\u0067\u0068\u0074"
	_gfbd.PositionAttr = &_aacg
	_ecec := "\u0036\u0036\u0032\u0039\u002c\u0031\u0034\u0039\u0037\u0031"
	_gfbd.XrangeAttr = &_ecec
	_cgbfgc.H = []*_aae.CT_H{_gfbd}
	return _cgbfgc
}

// NewWatermarkText generates a new WatermarkText.
func NewWatermarkText() WatermarkText {
	_aabaf := _aae.NewShapetype()
	_ggae := _aae.NewEG_ShapeElements()
	_ggae.Formulas = _adegd()
	_ggae.Path = _dffga()
	_ggae.Textpath = _ffbbg()
	_ggae.Handles = _egacb()
	_ggae.Lock = _gcbef()
	_aabaf.EG_ShapeElements = []*_aae.EG_ShapeElements{_ggae}
	var (
		_cdcc  = "_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0031\u0033\u0036"
		_eaeg  = "2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030"
		_gfcge = float32(136.0)
		_eddb  = "\u0031\u0030\u00380\u0030"
		_dfcc  = "m\u0040\u0037\u002c\u006c\u0040\u0038,\u006d\u0040\u0035\u002c\u0032\u0031\u0036\u0030\u0030l\u0040\u0036\u002c2\u00316\u0030\u0030\u0065"
	)
	_aabaf.IdAttr = &_cdcc
	_aabaf.CoordsizeAttr = &_eaeg
	_aabaf.SptAttr = &_gfcge
	_aabaf.AdjAttr = &_eddb
	_aabaf.PathAttr = &_dfcc
	_cbgbd := _aae.NewShape()
	_fdbdgd := _aae.NewEG_ShapeElements()
	_fdbdgd.Textpath = _dbbgc()
	_cbgbd.EG_ShapeElements = []*_aae.EG_ShapeElements{_fdbdgd}
	var (
		_fgba  = "\u0050\u006f\u0077\u0065\u0072\u0050l\u0075\u0073\u0057\u0061\u0074\u0065\u0072\u004d\u0061\u0072\u006b\u004f\u0062j\u0065\u0063\u0074\u0031\u0033\u0036\u00380\u0030\u0038\u0038\u0036"
		_baggg = "\u005f\u0078\u00300\u0030\u0030\u005f\u0073\u0032\u0030\u0035\u0031"
		_dgdb  = "\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0031\u0033\u0036"
		_fcaf  = ""
		_gbafg = "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065f\u0074:\u0030\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074o\u0070\u003a\u0030\u003b\u0077\u0069\u0064\u0074\u0068\u003a\u0034\u0036\u0038\u0070\u0074;\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0032\u0033\u0034\u0070\u0074\u003b\u007a\u002d\u0069\u006ede\u0078\u003a\u002d\u0032\u0035\u0031\u0036\u0035\u0031\u0030\u0037\u0032\u003b\u006d\u0073\u006f\u002d\u0077\u0072\u0061\u0070\u002d\u0065\u0064\u0069\u0074\u0065\u0064\u003a\u0066\u003b\u006d\u0073\u006f\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0070\u0065\u0072\u0063\u0065\u006e\u0074\u003a\u0030\u003b\u006d\u0073\u006f\u002d\u0068\u0065\u0069\u0067h\u0074-p\u0065\u0072\u0063\u0065\u006et\u003a\u0030\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006fn\u0074\u0061\u006c\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003b\u006d\u0073\u006f\u002d\u0070o\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u002d\u0072\u0065l\u0061\u0074\u0069\u0076\u0065:\u006d\u0061\u0072\u0067\u0069n\u003b\u006d\u0073o\u002d\u0070\u006f\u0073\u0069\u0074\u0069o\u006e-\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u002d\u0072e\u006c\u0061\u0074i\u0076\u0065\u003a\u006d\u0061\u0072\u0067\u0069\u006e\u003b\u006d\u0073\u006f\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0070\u0065\u0072\u0063e\u006e\u0074\u003a\u0030\u003b\u006d\u0073\u006f\u002dh\u0065\u0069\u0067\u0068t\u002dp\u0065\u0072\u0063\u0065\u006et\u003a0"
		_gcdf  = "\u0073\u0069\u006c\u0076\u0065\u0072"
	)
	_cbgbd.IdAttr = &_fgba
	_cbgbd.SpidAttr = &_baggg
	_cbgbd.TypeAttr = &_dgdb
	_cbgbd.AltAttr = &_fcaf
	_cbgbd.StyleAttr = &_gbafg
	_cbgbd.AllowincellAttr = _cg.ST_TrueFalseFalse
	_cbgbd.FillcolorAttr = &_gcdf
	_cbgbd.StrokedAttr = _cg.ST_TrueFalseFalse
	_bfgfc := _eaba.NewCT_Picture()
	_bfgfc.Any = []_d.Any{_aabaf, _cbgbd}
	return WatermarkText{_eefg: _bfgfc, _ggfda: _cbgbd, _cfgg: _aabaf}
}

// Properties returns the paragraph properties.
func (_eegg Paragraph) Properties() ParagraphProperties {
	_eegg.ensurePPr()
	return ParagraphProperties{_eegg._bbgc, _eegg._dbef.PPr}
}

// SetName sets the name of the image, visible in the properties of the image
// within Word.
func (_gdf AnchoredDrawing) SetName(name string) {
	_gdf._cc.DocPr.NameAttr = name
	for _, _db := range _gdf._cc.Graphic.GraphicData.Any {
		if _dce, _cgd := _db.(*_gd.Pic); _cgd {
			_dce.NvPicPr.CNvPr.DescrAttr = _d.String(name)
		}
	}
}

// SetVerticalAlignment sets the vertical alignment of content within a table cell.
func (_cdef CellProperties) SetVerticalAlignment(align _eaba.ST_VerticalJc) {
	if align == _eaba.ST_VerticalJcUnset {
		_cdef._be.VAlign = nil
	} else {
		_cdef._be.VAlign = _eaba.NewCT_VerticalJc()
		_cdef._be.VAlign.ValAttr = align
	}
}

// AddHeader creates a header associated with the document, but doesn't add it
// to the document for display.
func (_cec *Document) AddHeader() Header {
	_bcfg := _eaba.NewHdr()
	_cec._bbac = append(_cec._bbac, _bcfg)
	_dbg := _eg.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", len(_cec._bbac))
	_cec._bdb.AddRelationship(_dbg, _d.HeaderType)
	_cec.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_dbg, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0068\u0065\u0061\u0064e\u0072\u002b\u0078\u006d\u006c")
	_cec._cfg = append(_cec._cfg, _ea.NewRelationships())
	return Header{_cec, _bcfg}
}

// GetFooter gets a section Footer for given type
func (_fegbd Section) GetFooter(t _eaba.ST_HdrFtr) (Footer, bool) {
	for _, _daad := range _fegbd._aggca.EG_HdrFtrReferences {
		if _daad.FooterReference.TypeAttr == t {
			for _, _bddf := range _fegbd._ceec.Footers() {
				_agageb := _fegbd._ceec._bdb.FindRIDForN(_bddf.Index(), _d.FooterType)
				if _agageb == _daad.FooterReference.IdAttr {
					return _bddf, true
				}
			}
		}
	}
	return Footer{}, false
}

// IgnoreSpaceBetweenParagraphOfSameStyle sets contextual spacing.
func (_geagg Paragraph) IgnoreSpaceBetweenParagraphOfSameStyle() {
	_geagg.ensurePPr()
	if _geagg._dbef.PPr.ContextualSpacing == nil {
		_geagg._dbef.PPr.ContextualSpacing = _eaba.NewCT_OnOff()
	}
	_geagg._dbef.PPr.ContextualSpacing.ValAttr = &_cg.ST_OnOff{ST_OnOff1: _cg.ST_OnOff1On}
}

// GetImage returns the ImageRef associated with an AnchoredDrawing.
func (_aeg AnchoredDrawing) GetImage() (_ea.ImageRef, bool) {
	_dd := _aeg._cc.Graphic.GraphicData.Any
	if len(_dd) > 0 {
		_bca, _cd := _dd[0].(*_gd.Pic)
		if _cd {
			if _bca.BlipFill != nil && _bca.BlipFill.Blip != nil && _bca.BlipFill.Blip.EmbedAttr != nil {
				return _aeg._fa.GetImageByRelID(*_bca.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return _ea.ImageRef{}, false
}

// Clear resets the numbering.
func (_cebbc Numbering) Clear() {
	_cebbc._ebbdb.AbstractNum = nil
	_cebbc._ebbdb.Num = nil
	_cebbc._ebbdb.NumIdMacAtCleanup = nil
	_cebbc._ebbdb.NumPicBullet = nil
}

// SetLineSpacing controls the line spacing of the paragraph.
func (_dddba ParagraphStyleProperties) SetLineSpacing(m _eb.Distance, rule _eaba.ST_LineSpacingRule) {
	if _dddba._fded.Spacing == nil {
		_dddba._fded.Spacing = _eaba.NewCT_Spacing()
	}
	if rule == _eaba.ST_LineSpacingRuleUnset {
		_dddba._fded.Spacing.LineRuleAttr = _eaba.ST_LineSpacingRuleUnset
		_dddba._fded.Spacing.LineAttr = nil
	} else {
		_dddba._fded.Spacing.LineRuleAttr = rule
		_dddba._fded.Spacing.LineAttr = &_eaba.ST_SignedTwipsMeasure{}
		_dddba._fded.Spacing.LineAttr.Int64 = _d.Int64(int64(m / _eb.Twips))
	}
}

// SetWidthAuto sets the the table width to automatic.
func (_fbga TableProperties) SetWidthAuto() {
	_fbga._efac.TblW = _eaba.NewCT_TblWidth()
	_fbga._efac.TblW.TypeAttr = _eaba.ST_TblWidthAuto
}

func (_ecd *Document) save(_bdce _aa.Writer, _ecgd string) error {
	const _eee = "\u0064o\u0063u\u006d\u0065\u006e\u0074\u003a\u0064\u002e\u0053\u0061\u0076\u0065"
	if _fgfe := _ecd._eeda.Validate(); _fgfe != nil {
		_dcd.Log.Warning("\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u0020i\u006e\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0073", _fgfe)
	}
	_aaa := _d.DocTypeDocument
	if !_b.GetLicenseKey().IsLicensed() && !_gac {
		_eg.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
		_eg.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _ffa.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	}
	if len(_ecd._ebd) == 0 {
		if len(_ecgd) > 0 {
			_ecd._ebd = _ecgd
		} else {
			_dddg, _fdea := _b.GenRefId("\u0064\u0077")
			if _fdea != nil {
				_dcd.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fdea)
				return _fdea
			}
			_ecd._ebd = _dddg
		}
	}
	if _ebdd := _b.Track(_ecd._ebd, _eee); _ebdd != nil {
		_dcd.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ebdd)
		return _ebdd
	}
	_bcg := _da.NewWriter(_bdce)
	defer _bcg.Close()
	if _def := _fd.MarshalXML(_bcg, _d.BaseRelsFilename, _ecd.Rels.X()); _def != nil {
		return _def
	}
	if _gcg := _fd.MarshalXMLByType(_bcg, _aaa, _d.ExtendedPropertiesType, _ecd.AppProperties.X()); _gcg != nil {
		return _gcg
	}
	if _ded := _fd.MarshalXMLByType(_bcg, _aaa, _d.CorePropertiesType, _ecd.CoreProperties.X()); _ded != nil {
		return _ded
	}
	if _ecd.CustomProperties.X() != nil {
		if _gdb := _fd.MarshalXMLByType(_bcg, _aaa, _d.CustomPropertiesType, _ecd.CustomProperties.X()); _gdb != nil {
			return _gdb
		}
	}
	if _ecd.Thumbnail != nil {
		_dffb, _bcbf := _bcg.Create("\u0064\u006f\u0063Pr\u006f\u0070\u0073\u002f\u0074\u0068\u0075\u006d\u0062\u006e\u0061\u0069\u006c\u002e\u006a\u0070\u0065\u0067")
		if _bcbf != nil {
			return _bcbf
		}
		if _fca := _e.Encode(_dffb, _ecd.Thumbnail, nil); _fca != nil {
			return _fca
		}
	}
	if _ecaa := _fd.MarshalXMLByType(_bcg, _aaa, _d.SettingsType, _ecd.Settings.X()); _ecaa != nil {
		return _ecaa
	}
	_gce := _d.AbsoluteFilename(_aaa, _d.OfficeDocumentType, 0)
	if _daag := _fd.MarshalXML(_bcg, _gce, _ecd._eeda); _daag != nil {
		return _daag
	}
	if _gaf := _fd.MarshalXML(_bcg, _fd.RelationsPathFor(_gce), _ecd._bdb.X()); _gaf != nil {
		return _gaf
	}
	if _ecd.Numbering.X() != nil {
		if _cce := _fd.MarshalXMLByType(_bcg, _aaa, _d.NumberingType, _ecd.Numbering.X()); _cce != nil {
			return _cce
		}
	}
	if _fdda := _fd.MarshalXMLByType(_bcg, _aaa, _d.StylesType, _ecd.Styles.X()); _fdda != nil {
		return _fdda
	}
	if _ecd._feb != nil {
		if _dfa := _fd.MarshalXMLByType(_bcg, _aaa, _d.WebSettingsType, _ecd._feb); _dfa != nil {
			return _dfa
		}
	}
	if _ecd._fdd != nil {
		if _ege := _fd.MarshalXMLByType(_bcg, _aaa, _d.FontTableType, _ecd._fdd); _ege != nil {
			return _ege
		}
	}
	if _ecd._bfd != nil {
		if _bdd := _fd.MarshalXMLByType(_bcg, _aaa, _d.EndNotesType, _ecd._bfd); _bdd != nil {
			return _bdd
		}
	}
	if _ecd._edf != nil {
		if _cae := _fd.MarshalXMLByType(_bcg, _aaa, _d.FootNotesType, _ecd._edf); _cae != nil {
			return _cae
		}
	}
	for _eaaa, _agfa := range _ecd._bag {
		if _gag := _fd.MarshalXMLByTypeIndex(_bcg, _aaa, _d.ThemeType, _eaaa+1, _agfa); _gag != nil {
			return _gag
		}
	}
	for _afa, _ffbb := range _ecd._aac {
		if _dcbg := _fd.MarshalXMLByTypeIndex(_bcg, _aaa, _d.ControlType, _afa+1, _ffbb); _dcbg != nil {
			return _dcbg
		}
	}
	for _bfc, _aff := range _ecd._bbac {
		_fcd := _d.AbsoluteFilename(_aaa, _d.HeaderType, _bfc+1)
		if _ggge := _fd.MarshalXML(_bcg, _fcd, _aff); _ggge != nil {
			return _ggge
		}
		if !_ecd._cfg[_bfc].IsEmpty() {
			_fd.MarshalXML(_bcg, _fd.RelationsPathFor(_fcd), _ecd._cfg[_bfc].X())
		}
	}
	for _cbf, _fbf := range _ecd._gfdb {
		_efee := _d.AbsoluteFilename(_aaa, _d.FooterType, _cbf+1)
		if _ega := _fd.MarshalXMLByTypeIndex(_bcg, _aaa, _d.FooterType, _cbf+1, _fbf); _ega != nil {
			return _ega
		}
		if !_ecd._efc[_cbf].IsEmpty() {
			_fd.MarshalXML(_bcg, _fd.RelationsPathFor(_efee), _ecd._efc[_cbf].X())
		}
	}
	for _bdgf, _ebddd := range _ecd.Images {
		if _ffd := _ea.AddImageToZip(_bcg, _ebddd, _bdgf+1, _d.DocTypeDocument); _ffd != nil {
			return _ffd
		}
	}
	for _abf, _gee := range _ecd._cea {
		_cdee := _d.AbsoluteFilename(_aaa, _d.ChartType, _abf+1)
		_fd.MarshalXML(_bcg, _cdee, _gee._cebb)
	}
	if _bfe := _fd.MarshalXML(_bcg, _d.ContentTypesFilename, _ecd.ContentTypes.X()); _bfe != nil {
		return _bfe
	}
	if _dge := _ecd.WriteExtraFiles(_bcg); _dge != nil {
		return _dge
	}
	return _bcg.Close()
}

// InsertRowBefore inserts a row before another row
func (_acca Table) InsertRowBefore(r Row) Row {
	for _gacca, _geacf := range _acca._fecae.EG_ContentRowContent {
		if len(_geacf.Tr) > 0 && r.X() == _geacf.Tr[0] {
			_dfbb := _eaba.NewEG_ContentRowContent()
			_acca._fecae.EG_ContentRowContent = append(_acca._fecae.EG_ContentRowContent, nil)
			copy(_acca._fecae.EG_ContentRowContent[_gacca+1:], _acca._fecae.EG_ContentRowContent[_gacca:])
			_acca._fecae.EG_ContentRowContent[_gacca] = _dfbb
			_eacd := _eaba.NewCT_Row()
			_dfbb.Tr = append(_dfbb.Tr, _eacd)
			return Row{_acca._bggbd, _eacd}
		}
	}
	return _acca.AddRow()
}

// SetChecked marks a FormFieldTypeCheckBox as checked or unchecked.
func (_aeea FormField) SetChecked(b bool) {
	if _aeea._cdag.CheckBox == nil {
		return
	}
	if !b {
		_aeea._cdag.CheckBox.Checked = nil
	} else {
		_aeea._cdag.CheckBox.Checked = _eaba.NewCT_OnOff()
	}
}

// SetLastRow controls the conditional formatting for the last row in a table.
// This is called the 'Total' row within Word.
func (_ccage TableLook) SetLastRow(on bool) {
	if !on {
		_ccage._ffdgf.LastRowAttr = &_cg.ST_OnOff{}
		_ccage._ffdgf.LastRowAttr.ST_OnOff1 = _cg.ST_OnOff1Off
	} else {
		_ccage._ffdgf.LastRowAttr = &_cg.ST_OnOff{}
		_ccage._ffdgf.LastRowAttr.ST_OnOff1 = _cg.ST_OnOff1On
	}
}

// SetHANSITheme sets the font H ANSI Theme.
func (_gdegb Fonts) SetHANSITheme(t _eaba.ST_Theme) { _gdegb._aggeg.HAnsiThemeAttr = t }

// Footnote returns the footnote based on the ID; this can be used nicely with
// the run.IsFootnote() functionality.
func (_aefd *Document) Footnote(id int64) Footnote {
	for _, _gffb := range _aefd.Footnotes() {
		if _gffb.id() == id {
			return _gffb
		}
	}
	return Footnote{}
}

// AddTable adds a table to the table cell.
func (_bab Cell) AddTable() Table {
	_gfg := _eaba.NewEG_BlockLevelElts()
	_bab._acg.EG_BlockLevelElts = append(_bab._acg.EG_BlockLevelElts, _gfg)
	_bdf := _eaba.NewEG_ContentBlockContent()
	_gfg.EG_ContentBlockContent = append(_gfg.EG_ContentBlockContent, _bdf)
	_cde := _eaba.NewCT_Tbl()
	_bdf.Tbl = append(_bdf.Tbl, _cde)
	return Table{_bab._fc, _cde}
}

// Styles returns all styles.
func (_acafc Styles) Styles() []Style {
	_bbgd := []Style{}
	for _, _cecba := range _acafc._abad.Style {
		_bbgd = append(_bbgd, Style{_cecba})
	}
	return _bbgd
}

// SetLayoutInCell sets the layoutInCell attribute of anchor.
func (_ffcf AnchoredDrawing) SetLayoutInCell(val bool) { _ffcf._cc.LayoutInCellAttr = val }

// SetStyle sets the style of a paragraph.
func (_gecd ParagraphProperties) SetStyle(s string) {
	if s == "" {
		_gecd._efec.PStyle = nil
	} else {
		_gecd._efec.PStyle = _eaba.NewCT_String()
		_gecd._efec.PStyle.ValAttr = s
	}
}

// InsertParagraphAfter adds a new empty paragraph after the relativeTo
// paragraph.
func (_ffe *Document) InsertParagraphAfter(relativeTo Paragraph) Paragraph {
	return _ffe.insertParagraph(relativeTo, false)
}

// X returns the inner wrapped XML type.
func (_bccf NumberingDefinition) X() *_eaba.CT_AbstractNum { return _bccf._fbde }

// SetBold sets the run to bold.
func (_ddea RunProperties) SetBold(b bool) {
	if !b {
		_ddea._gfbgc.B = nil
		_ddea._gfbgc.BCs = nil
	} else {
		_ddea._gfbgc.B = _eaba.NewCT_OnOff()
		_ddea._gfbgc.BCs = _eaba.NewCT_OnOff()
	}
}

// SetAlignment set alignment of paragraph.
func (_ceaad Paragraph) SetAlignment(alignment _eaba.ST_Jc) {
	_ceaad.ensurePPr()
	if _ceaad._dbef.PPr.Jc == nil {
		_ceaad._dbef.PPr.Jc = _eaba.NewCT_Jc()
	}
	_ceaad._dbef.PPr.Jc.ValAttr = alignment
}

// DoubleStrike returns true if run is double striked.
func (_bcggc RunProperties) DoubleStrike() bool { return _beff(_bcggc._gfbgc.Dstrike) }

// Borders allows controlling individual cell borders.
func (_eff CellProperties) Borders() CellBorders {
	if _eff._be.TcBorders == nil {
		_eff._be.TcBorders = _eaba.NewCT_TcBorders()
	}
	return CellBorders{_eff._be.TcBorders}
}

// SetTop sets the top border to a specified type, color and thickness.
func (_efe CellBorders) SetTop(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_efe._ggf.Top = _eaba.NewCT_Border()
	_gfdf(_efe._ggf.Top, t, c, thickness)
}

// SetHangingIndent controls special indent of paragraph.
func (_gcee Paragraph) SetHangingIndent(m _eb.Distance) {
	_gcee.ensurePPr()
	_aggc := _gcee._dbef.PPr
	if _aggc.Ind == nil {
		_aggc.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_aggc.Ind.HangingAttr = nil
	} else {
		_aggc.Ind.HangingAttr = &_cg.ST_TwipsMeasure{}
		_aggc.Ind.HangingAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(m / _eb.Twips))
	}
}

// Tables returns the tables defined in the footer.
func (_dade Footer) Tables() []Table {
	_agad := []Table{}
	if _dade._daaa == nil {
		return nil
	}
	for _, _dbeg := range _dade._daaa.EG_ContentBlockContent {
		for _, _fecg := range _dade._dec.tables(_dbeg) {
			_agad = append(_agad, _fecg)
		}
	}
	return _agad
}

// X returns the inner wrapped XML type.
func (_fcab TableStyleProperties) X() *_eaba.CT_TblPrBase { return _fcab._ddafc }

// SetTextStyleItalic set text style of watermark to italic.
func (_baaac *WatermarkText) SetTextStyleItalic(value bool) {
	if _baaac._ggfda != nil {
		_bedca := _baaac.GetStyle()
		_bedca.SetItalic(value)
		_baaac.SetStyle(_bedca)
	}
}

// SetName sets the name of the bookmark. This is the name that is used to
// reference the bookmark from hyperlinks.
func (_cag Bookmark) SetName(name string) { _cag._dda.NameAttr = name }

// PutNodeBefore put node to position before relativeTo.
func (_bdcaf *Document) PutNodeBefore(relativeTo, node Node) {
	_bdcaf.putNode(relativeTo, node, true)
}

// SetAfter sets the spacing that comes after the paragraph.
func (_dbgd ParagraphSpacing) SetAfter(after _eb.Distance) {
	_dbgd._ebfge.AfterAttr = &_cg.ST_TwipsMeasure{}
	_dbgd._ebfge.AfterAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(after / _eb.Twips))
}

// SetWidth sets the table with to a specified width.
func (_afdg TableProperties) SetWidth(d _eb.Distance) {
	_afdg._efac.TblW = _eaba.NewCT_TblWidth()
	_afdg._efac.TblW.TypeAttr = _eaba.ST_TblWidthDxa
	_afdg._efac.TblW.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_afdg._efac.TblW.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_afdg._efac.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(d / _eb.Twips))
}

// SetTextWrapThrough sets the text wrap to through with a give wrap type.
func (_gdg AnchoredDrawing) SetTextWrapThrough(option *AnchorDrawWrapOptions) {
	_gdg._cc.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_gdg._cc.Choice.WrapThrough = _eaba.NewWdCT_WrapThrough()
	_gdg._cc.Choice.WrapThrough.WrapTextAttr = _eaba.WdST_WrapTextBothSides
	_cda := false
	_gdg._cc.Choice.WrapThrough.WrapPolygon.EditedAttr = &_cda
	if option == nil {
		option = NewAnchorDrawWrapOptions()
	}
	_gdg._cc.Choice.WrapThrough.WrapPolygon.Start = option.GetWrapPathStart()
	_gdg._cc.Choice.WrapThrough.WrapPolygon.LineTo = option.GetWrapPathLineTo()
	_gdg._cc.LayoutInCellAttr = true
	_gdg._cc.AllowOverlapAttr = true
}

// Runs returns all of the runs in a paragraph.
func (_cacg Paragraph) Runs() []Run {
	_ecgfc := []Run{}
	for _, _bbdc := range _cacg._dbef.EG_PContent {
		if _bbdc.Hyperlink != nil && _bbdc.Hyperlink.EG_ContentRunContent != nil {
			for _, _fdca := range _bbdc.Hyperlink.EG_ContentRunContent {
				if _fdca.R != nil {
					_ecgfc = append(_ecgfc, Run{_cacg._bbgc, _fdca.R})
				}
			}
		}
		for _, _ddfe := range _bbdc.EG_ContentRunContent {
			if _ddfe.R != nil {
				_ecgfc = append(_ecgfc, Run{_cacg._bbgc, _ddfe.R})
			}
			if _ddfe.Sdt != nil && _ddfe.Sdt.SdtContent != nil {
				for _, _fagb := range _ddfe.Sdt.SdtContent.EG_ContentRunContent {
					if _fagb.R != nil {
						_ecgfc = append(_ecgfc, Run{_cacg._bbgc, _fagb.R})
					}
				}
			}
		}
	}
	return _ecgfc
}

// SetRight sets the right border to a specified type, color and thickness.
func (_gge CellBorders) SetRight(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_gge._ggf.Right = _eaba.NewCT_Border()
	_gfdf(_gge._ggf.Right, t, c, thickness)
}

// Endnotes returns the endnotes defined in the document.
func (_gceg *Document) Endnotes() []Endnote {
	_effa := []Endnote{}
	for _, _fcf := range _gceg._bfd.CT_Endnotes.Endnote {
		_effa = append(_effa, Endnote{_gceg, _fcf})
	}
	return _effa
}

// GetStyleByID returns Style by it's IdAttr.
func (_adab *Document) GetStyleByID(id string) Style {
	for _, _ccdb := range _adab.Styles._abad.Style {
		if _ccdb.StyleIdAttr != nil && *_ccdb.StyleIdAttr == id {
			return Style{_ccdb}
		}
	}
	return Style{}
}

// AddEndnote will create a new endnote and attach it to the Paragraph in the
// location at the end of the previous run (endnotes create their own run within
// the paragraph. The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the endnote later.
func (_dgca Paragraph) AddEndnote(text string) Endnote {
	var _edaf int64
	if _dgca._bbgc.HasEndnotes() {
		for _, _eeed := range _dgca._bbgc.Endnotes() {
			if _eeed.id() > _edaf {
				_edaf = _eeed.id()
			}
		}
		_edaf++
	} else {
		_edaf = 0
		_dgca._bbgc._bfd = &_eaba.Endnotes{}
	}
	_feab := _eaba.NewCT_FtnEdn()
	_dffe := _eaba.NewCT_FtnEdnRef()
	_dffe.IdAttr = _edaf
	_dgca._bbgc._bfd.CT_Endnotes.Endnote = append(_dgca._bbgc._bfd.CT_Endnotes.Endnote, _feab)
	_fabe := _dgca.AddRun()
	_bfecg := _fabe.Properties()
	_bfecg.SetStyle("\u0045\u006e\u0064\u006e\u006f\u0074\u0065\u0041\u006e\u0063\u0068\u006f\u0072")
	_fabe._bggd.EG_RunInnerContent = []*_eaba.EG_RunInnerContent{_eaba.NewEG_RunInnerContent()}
	_fabe._bggd.EG_RunInnerContent[0].EndnoteReference = _dffe
	_fcaa := Endnote{_dgca._bbgc, _feab}
	_fcaa._ggd.IdAttr = _edaf
	_fcaa._ggd.EG_BlockLevelElts = []*_eaba.EG_BlockLevelElts{_eaba.NewEG_BlockLevelElts()}
	_febfd := _fcaa.AddParagraph()
	_febfd.Properties().SetStyle("\u0045n\u0064\u006e\u006f\u0074\u0065")
	_febfd._dbef.PPr.RPr = _eaba.NewCT_ParaRPr()
	_ffcd := _febfd.AddRun()
	_ffcd.AddTab()
	_ffcd.AddText(text)
	return _fcaa
}

func _fgf(_bbf *_eaba.CT_TblWidth, _eca _eb.Distance) {
	_bbf.TypeAttr = _eaba.ST_TblWidthDxa
	_bbf.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_bbf.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_bbf.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(_eca / _eb.Dxa))
}

// Italic returns true if run font is italic.
func (_ggcda RunProperties) Italic() bool {
	_dbeeb := _ggcda._gfbgc
	return _beff(_dbeeb.I) || _beff(_dbeeb.ICs)
}

// Open opens and reads a document from a file (.docx).
func Open(filename string) (*Document, error) {
	_gacea, _gggag := _ff.Open(filename)
	if _gggag != nil {
		return nil, _eg.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _gggag)
	}
	defer _gacea.Close()
	_cdge, _gggag := _ff.Stat(filename)
	if _gggag != nil {
		return nil, _eg.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _gggag)
	}
	_ = _cdge
	return Read(_gacea, _cdge.Size())
}

// AnchorDrawWrapOptions is options to set
// wrapPolygon for wrap text through and tight.
type AnchorDrawWrapOptions struct {
	_fb  bool
	_bba *_bb.CT_Point2D
	_ga  []*_bb.CT_Point2D
}

// SetSize sets the size of the displayed image on the page.
func (_fgc AnchoredDrawing) SetSize(w, h _eb.Distance) {
	_fgc._cc.Extent.CxAttr = int64(float64(w*_eb.Pixel72) / _eb.EMU)
	_fgc._cc.Extent.CyAttr = int64(float64(h*_eb.Pixel72) / _eb.EMU)
}

// SetSize sets size attribute for a FormFieldTypeCheckBox in pt.
func (_adfbf FormField) SetSize(size uint64) {
	size *= 2
	if _adfbf._cdag.CheckBox != nil {
		_adfbf._cdag.CheckBox.Choice = _eaba.NewCT_FFCheckBoxChoice()
		_adfbf._cdag.CheckBox.Choice.Size = _eaba.NewCT_HpsMeasure()
		_adfbf._cdag.CheckBox.Choice.Size.ValAttr = _eaba.ST_HpsMeasure{ST_UnsignedDecimalNumber: &size}
	}
}

// DocText is an array of extracted text items which has some methods for representing extracted text.
type DocText struct {
	Items  []TextItem
	_dbage []listItemInfo
	_bcgfa map[int64]map[int64]int64
}

// Tables returns the tables defined in the header.
func (_bgdd Header) Tables() []Table {
	_cdgbg := []Table{}
	if _bgdd._cbgbf == nil {
		return nil
	}
	for _, _egefg := range _bgdd._cbgbf.EG_ContentBlockContent {
		for _, _fgbf := range _bgdd._bffa.tables(_egefg) {
			_cdgbg = append(_cdgbg, _fgbf)
		}
	}
	return _cdgbg
}

// AddParagraph adds a new paragraph to the document body.
func (_dcba *Document) AddParagraph() Paragraph {
	_dbfb := _eaba.NewEG_BlockLevelElts()
	_dcba._eeda.Body.EG_BlockLevelElts = append(_dcba._eeda.Body.EG_BlockLevelElts, _dbfb)
	_fbb := _eaba.NewEG_ContentBlockContent()
	_dbfb.EG_ContentBlockContent = append(_dbfb.EG_ContentBlockContent, _fbb)
	_aacd := _eaba.NewCT_P()
	_fbb.P = append(_fbb.P, _aacd)
	return Paragraph{_dcba, _aacd}
}

// TextWithOptions extract text with options.
func (_eaec *DocText) TextWithOptions(options ExtractTextOptions) string {
	_abbc := make(map[int64]map[int64]int64, 0)
	_defab := _cb.NewBuffer([]byte{})
	_ccbg := int64(0)
	_fecd := int64(0)
	_gcgbe := int64(0)
	for _ceed, _agbd := range _eaec.Items {
		_fcdd := false
		if _agbd.Text != "" {
			if options.WithNumbering {
				if _ceed > 0 {
					if _agbd.Paragraph != _eaec.Items[_ceed-1].Paragraph {
						_fcdd = true
					}
				} else {
					_fcdd = true
				}
				if _fcdd {
					for _, _fgfa := range _eaec._dbage {
						if _fgfa.FromParagraph == nil {
							continue
						}
						if _fgfa.FromParagraph.X() == _agbd.Paragraph {
							if _fgfa.AbstractNumId != nil && _eaec._bcgfa[*_fgfa.AbstractNumId][_fgfa.NumberingLevel.X().IlvlAttr] > 0 {
								if _, _fdge := _abbc[*_fgfa.AbstractNumId]; _fdge {
									if _, _agbe := _abbc[*_fgfa.AbstractNumId][_fgfa.NumberingLevel.X().IlvlAttr]; _agbe {
										_abbc[*_fgfa.AbstractNumId][_fgfa.NumberingLevel.X().IlvlAttr]++
									} else {
										_abbc[*_fgfa.AbstractNumId][_fgfa.NumberingLevel.X().IlvlAttr] = 1
									}
								} else {
									_abbc[*_fgfa.AbstractNumId] = map[int64]int64{_fgfa.NumberingLevel.X().IlvlAttr: 1}
								}
								if _ccbg == _fgfa.NumberingLevel.X().IlvlAttr && _fgfa.NumberingLevel.X().IlvlAttr > 0 {
									_fecd++
								} else {
									_fecd = _abbc[*_fgfa.AbstractNumId][_fgfa.NumberingLevel.X().IlvlAttr]
									if _fgfa.NumberingLevel.X().IlvlAttr > _ccbg && _gcgbe == *_fgfa.AbstractNumId {
										_fecd = 1
									}
								}
								_ebgc := ""
								if _fgfa.NumberingLevel.X().LvlText.ValAttr != nil {
									_ebgc = *_fgfa.NumberingLevel.X().LvlText.ValAttr
								}
								_cgbe := _fdgf(_fecd, _fgfa.NumberingLevel.X().IlvlAttr, _ebgc, _fgfa.NumberingLevel.X().NumFmt, _abbc[*_fgfa.AbstractNumId])
								_defab.WriteString(_cgbe)
								_eaec._bcgfa[*_fgfa.AbstractNumId][_fgfa.NumberingLevel.X().IlvlAttr]--
								_ccbg = _fgfa.NumberingLevel.X().IlvlAttr
								_gcgbe = *_fgfa.AbstractNumId
								if options.NumberingIndent != "" {
									_defab.WriteString(options.NumberingIndent)
								}
							}
							break
						}
					}
				}
			}
			_defab.WriteString(_agbd.Text)
			_defab.WriteString("\u000a")
		}
	}
	return _defab.String()
}

// Rows returns the rows defined in the table.
func (_badd Table) Rows() []Row {
	_cgae := []Row{}
	for _, _egbdf := range _badd._fecae.EG_ContentRowContent {
		for _, _agccc := range _egbdf.Tr {
			_cgae = append(_cgae, Row{_badd._bggbd, _agccc})
		}
		if _egbdf.Sdt != nil && _egbdf.Sdt.SdtContent != nil {
			for _, _gbeg := range _egbdf.Sdt.SdtContent.Tr {
				_cgae = append(_cgae, Row{_badd._bggbd, _gbeg})
			}
		}
	}
	return _cgae
}

// GetDocRelTargetByID returns TargetAttr of document relationship given its IdAttr.
func (_egcf *Document) GetDocRelTargetByID(idAttr string) string {
	for _, _dad := range _egcf._bdb.X().Relationship {
		if _dad.IdAttr == idAttr {
			return _dad.TargetAttr
		}
	}
	return ""
}

func (_cad *Document) appendParagraph(_bcgd *Paragraph, _dbgg Paragraph, _fcea bool) Paragraph {
	_agcf := _eaba.NewEG_BlockLevelElts()
	_cad._eeda.Body.EG_BlockLevelElts = append(_cad._eeda.Body.EG_BlockLevelElts, _agcf)
	_ecc := _eaba.NewEG_ContentBlockContent()
	_agcf.EG_ContentBlockContent = append(_agcf.EG_ContentBlockContent, _ecc)
	if _bcgd != nil {
		_aee := _bcgd.X()
		for _, _efb := range _cad._eeda.Body.EG_BlockLevelElts {
			for _, _cdb := range _efb.EG_ContentBlockContent {
				for _gdd, _eegb := range _cdb.P {
					if _eegb == _aee {
						_dgea := _dbgg.X()
						_cdb.P = append(_cdb.P, nil)
						if _fcea {
							copy(_cdb.P[_gdd+1:], _cdb.P[_gdd:])
							_cdb.P[_gdd] = _dgea
						} else {
							copy(_cdb.P[_gdd+2:], _cdb.P[_gdd+1:])
							_cdb.P[_gdd+1] = _dgea
						}
						break
					}
				}
				for _, _dddfe := range _cdb.Tbl {
					for _, _aagb := range _dddfe.EG_ContentRowContent {
						for _, _fae := range _aagb.Tr {
							for _, _agbc := range _fae.EG_ContentCellContent {
								for _, _dca := range _agbc.Tc {
									for _, _cfdg := range _dca.EG_BlockLevelElts {
										for _, _bgfe := range _cfdg.EG_ContentBlockContent {
											for _ddc, _agbg := range _bgfe.P {
												if _agbg == _aee {
													_cadg := _dbgg.X()
													_bgfe.P = append(_bgfe.P, nil)
													if _fcea {
														copy(_bgfe.P[_ddc+1:], _bgfe.P[_ddc:])
														_bgfe.P[_ddc] = _cadg
													} else {
														copy(_bgfe.P[_ddc+2:], _bgfe.P[_ddc+1:])
														_bgfe.P[_ddc+1] = _cadg
													}
													break
												}
											}
										}
									}
								}
							}
						}
					}
				}
				if _cdb.Sdt != nil && _cdb.Sdt.SdtContent != nil && _cdb.Sdt.SdtContent.P != nil {
					for _cabea, _gdfa := range _cdb.Sdt.SdtContent.P {
						if _gdfa == _aee {
							_bce := _dbgg.X()
							_cdb.Sdt.SdtContent.P = append(_cdb.Sdt.SdtContent.P, nil)
							if _fcea {
								copy(_cdb.Sdt.SdtContent.P[_cabea+1:], _cdb.Sdt.SdtContent.P[_cabea:])
								_cdb.Sdt.SdtContent.P[_cabea] = _bce
							} else {
								copy(_cdb.Sdt.SdtContent.P[_cabea+2:], _cdb.Sdt.SdtContent.P[_cabea+1:])
								_cdb.Sdt.SdtContent.P[_cabea+1] = _bce
							}
							break
						}
					}
				}
			}
		}
	} else {
		_ecc.P = append(_ecc.P, _dbgg.X())
	}
	_efa := _dbgg.Properties()
	if _geg, _dbea := _efa.Section(); _dbea {
		var (
			_egce map[string]string
			_gfdd map[string]string
		)
		_adea := _geg.X().EG_HdrFtrReferences
		for _, _ede := range _adea {
			if _ede.HeaderReference != nil {
				_egce = map[string]string{_ede.HeaderReference.IdAttr: _geg._ceec._bdb.GetTargetByRelId(_ede.HeaderReference.IdAttr)}
			}
			if _ede.FooterReference != nil {
				_gfdd = map[string]string{_ede.FooterReference.IdAttr: _geg._ceec._bdb.GetTargetByRelId(_ede.FooterReference.IdAttr)}
			}
		}
		var _afe map[int]_ea.ImageRef
		for _, _cfb := range _geg._ceec.Headers() {
			for _fgge, _fafb := range _egce {
				_egca := _eg.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", (_cfb.Index() + 1))
				if _egca == _fafb {
					_fbe := _eg.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", _cfb.Index())
					_cad._bbac = append(_cad._bbac, _cfb.X())
					_fced := _cad._bdb.AddRelationship(_fbe, _d.HeaderType)
					_fced.SetID(_fgge)
					_cad.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_fbe, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0068\u0065\u0061\u0064e\u0072\u002b\u0078\u006d\u006c")
					_cad._cfg = append(_cad._cfg, _ea.NewRelationships())
					_cef := _cfb.Paragraphs()
					for _, _agff := range _cef {
						for _, _efeb := range _agff.Runs() {
							_dbae := _efeb.DrawingAnchored()
							for _, _ceg := range _dbae {
								if _dfadg, _cbad := _ceg.GetImage(); _cbad {
									_afe = map[int]_ea.ImageRef{_cfb.Index(): _dfadg}
								}
							}
							_gfa := _efeb.DrawingInline()
							for _, _aba := range _gfa {
								if _fbff, _cefa := _aba.GetImage(); _cefa {
									_afe = map[int]_ea.ImageRef{_cfb.Index(): _fbff}
								}
							}
						}
					}
				}
			}
		}
		for _fffe, _gegd := range _afe {
			for _, _cfad := range _cad.Headers() {
				if (_cfad.Index() + 1) == _fffe {
					_cdgf, _dbga := _ea.ImageFromFile(_gegd.Path())
					if _dbga != nil {
						_dcd.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _dbga)
					}
					if _, _dbga = _cfad.AddImage(_cdgf); _dbga != nil {
						_dcd.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _dbga)
					}
				}
				for _, _ebf := range _cfad.Paragraphs() {
					if _fbffc, _gcgd := _geg._ceec.Styles.SearchStyleById(_ebf.Style()); _gcgd {
						if _, _beb := _cad.Styles.SearchStyleById(_ebf.Style()); !_beb {
							_cad.Styles.InsertStyle(_fbffc)
						}
					}
				}
			}
		}
		var _fge map[int]_ea.ImageRef
		for _, _agca := range _geg._ceec.Footers() {
			for _eaad, _bcag := range _gfdd {
				_caf := _eg.Sprintf("\u0066\u006f\u006ft\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", (_agca.Index() + 1))
				if _caf == _bcag {
					_edb := _eg.Sprintf("\u0066\u006f\u006ft\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", _agca.Index())
					_cad._gfdb = append(_cad._gfdb, _agca.X())
					_dbffc := _cad._bdb.AddRelationship(_edb, _d.FooterType)
					_dbffc.SetID(_eaad)
					_cad.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_edb, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0066\u006f\u006f\u0074e\u0072\u002b\u0078\u006d\u006c")
					_cad._efc = append(_cad._efc, _ea.NewRelationships())
					_bafe := _agca.Paragraphs()
					for _, _cegc := range _bafe {
						for _, _ecfd := range _cegc.Runs() {
							_dbag := _ecfd.DrawingAnchored()
							for _, _dcea := range _dbag {
								if _egbd, _gfgce := _dcea.GetImage(); _gfgce {
									_fge = map[int]_ea.ImageRef{_agca.Index(): _egbd}
								}
							}
							_cdff := _ecfd.DrawingInline()
							for _, _dgcb := range _cdff {
								if _agfd, _gccg := _dgcb.GetImage(); _gccg {
									_fge = map[int]_ea.ImageRef{_agca.Index(): _agfd}
								}
							}
						}
					}
				}
			}
		}
		for _fga, _bbab := range _fge {
			for _, _gbd := range _cad.Footers() {
				if (_gbd.Index() + 1) == _fga {
					_aceg, _gabg := _ea.ImageFromFile(_bbab.Path())
					if _gabg != nil {
						_dcd.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _gabg)
					}
					if _, _gabg = _gbd.AddImage(_aceg); _gabg != nil {
						_dcd.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _gabg)
					}
				}
				for _, _bec := range _gbd.Paragraphs() {
					if _cfcg, _efegc := _geg._ceec.Styles.SearchStyleById(_bec.Style()); _efegc {
						if _, _gbab := _cad.Styles.SearchStyleById(_bec.Style()); !_gbab {
							_cad.Styles.InsertStyle(_cfcg)
						}
					}
				}
			}
		}
	}
	_eebb := _dbgg.Numbering()
	_cad.Numbering._ebbdb.AbstractNum = append(_cad.Numbering._ebbdb.AbstractNum, _eebb._ebbdb.AbstractNum...)
	_cad.Numbering._ebbdb.Num = append(_cad.Numbering._ebbdb.Num, _eebb._ebbdb.Num...)
	return Paragraph{_cad, _dbgg.X()}
}

// Themes returns document's themes.
func (_gegdg *Document) Themes() []*_bb.Theme { return _gegdg._bag }

// X returns the internally wrapped *wml.CT_SectPr.
func (_ebdee Section) X() *_eaba.CT_SectPr { return _ebdee._aggca }

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (_afag *Document) GetImageByRelID(relID string) (_ea.ImageRef, bool) {
	_bfef := _afag._bdb.GetTargetByRelId(relID)
	_eec := ""
	for _, _cgf := range _afag._cfg {
		if _eec != "" {
			break
		}
		_eec = _cgf.GetTargetByRelId(relID)
	}
	for _, _egde := range _afag.Images {
		if _egde.RelID() == relID {
			return _egde, true
		}
		if _bfef != "" {
			_bddb := _f.Replace(_egde.Target(), "\u0077\u006f\u0072d\u002f", "", 1)
			if _bddb == _bfef {
				if _egde.RelID() == "" {
					_egde.SetRelID(relID)
				}
				return _egde, true
			}
		}
		if _eec != "" {
			_cdec := _f.Replace(_egde.Target(), "\u0077\u006f\u0072d\u002f", "", 1)
			if _cdec == _eec {
				if _egde.RelID() == "" {
					_egde.SetRelID(relID)
				}
				return _egde, true
			}
		}
	}
	return _ea.ImageRef{}, false
}

// MultiLevelType returns the multilevel type, or ST_MultiLevelTypeUnset if not set.
func (_dcaf NumberingDefinition) MultiLevelType() _eaba.ST_MultiLevelType {
	if _dcaf._fbde.MultiLevelType != nil {
		return _dcaf._fbde.MultiLevelType.ValAttr
	} else {
		return _eaba.ST_MultiLevelTypeUnset
	}
}

var _gac = false

// Paragraph is a paragraph within a document.
type Paragraph struct {
	_bbgc *Document
	_dbef *_eaba.CT_P
}

// SetWidthPercent sets the cell to a width percentage.
func (_ecg CellProperties) SetWidthPercent(pct float64) {
	_ecg._be.TcW = _eaba.NewCT_TblWidth()
	_ecg._be.TcW.TypeAttr = _eaba.ST_TblWidthPct
	_ecg._be.TcW.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_ecg._be.TcW.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_ecg._be.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(pct * 50))
}

func (_edcd Styles) initializeDocDefaults() {
	_edcd._abad.DocDefaults = _eaba.NewCT_DocDefaults()
	_edcd._abad.DocDefaults.RPrDefault = _eaba.NewCT_RPrDefault()
	_edcd._abad.DocDefaults.RPrDefault.RPr = _eaba.NewCT_RPr()
	_aaaab := RunProperties{_edcd._abad.DocDefaults.RPrDefault.RPr}
	_aaaab.SetSize(12 * _eb.Point)
	_aaaab.Fonts().SetASCIITheme(_eaba.ST_ThemeMajorAscii)
	_aaaab.Fonts().SetEastAsiaTheme(_eaba.ST_ThemeMajorEastAsia)
	_aaaab.Fonts().SetHANSITheme(_eaba.ST_ThemeMajorHAnsi)
	_aaaab.Fonts().SetCSTheme(_eaba.ST_ThemeMajorBidi)
	_aaaab.X().Lang = _eaba.NewCT_Language()
	_aaaab.X().Lang.ValAttr = _d.String("\u0065\u006e\u002dU\u0053")
	_aaaab.X().Lang.EastAsiaAttr = _d.String("\u0065\u006e\u002dU\u0053")
	_aaaab.X().Lang.BidiAttr = _d.String("\u0061\u0072\u002dS\u0041")
	_edcd._abad.DocDefaults.PPrDefault = _eaba.NewCT_PPrDefault()
}

// Node is document element node,
// contains Paragraph or Table element.
type Node struct {
	_abcbe           *Document
	_efag            interface{}
	Style            Style
	AnchoredDrawings []AnchoredDrawing
	InlineDrawings   []InlineDrawing
	Children         []Node
}

// SetEmboss sets the run to embossed text.
func (_eage RunProperties) SetEmboss(b bool) {
	if !b {
		_eage._gfbgc.Emboss = nil
	} else {
		_eage._gfbgc.Emboss = _eaba.NewCT_OnOff()
	}
}

func _gfff(_bdad *_eaba.CT_P, _cbac, _eeeg map[int64]int64) {
	for _, _ecgg := range _bdad.EG_PContent {
		for _, _gdbb := range _ecgg.EG_ContentRunContent {
			if _gdbb.R != nil {
				for _, _ecgfg := range _gdbb.R.EG_RunInnerContent {
					_eecc := _ecgfg.EndnoteReference
					if _eecc != nil && _eecc.IdAttr > 0 {
						if _dgcc, _abed := _eeeg[_eecc.IdAttr]; _abed {
							_eecc.IdAttr = _dgcc
						}
					}
					_cfcgg := _ecgfg.FootnoteReference
					if _cfcgg != nil && _cfcgg.IdAttr > 0 {
						if _ddbc, _cfba := _cbac[_cfcgg.IdAttr]; _cfba {
							_cfcgg.IdAttr = _ddbc
						}
					}
				}
			}
		}
	}
}

func _eac(_agfg *_eaba.CT_P, _cggg map[string]string) {
	for _, _fadaa := range _agfg.EG_PContent {
		for _, _gfda := range _fadaa.EG_ContentRunContent {
			if _gfda.R != nil {
				for _, _egaab := range _gfda.R.EG_RunInnerContent {
					_gade := _egaab.Drawing
					if _gade != nil {
						for _, _bgfb := range _gade.Anchor {
							for _, _gcca := range _bgfb.Graphic.GraphicData.Any {
								switch _bgeg := _gcca.(type) {
								case *_gd.Pic:
									if _bgeg.BlipFill != nil && _bgeg.BlipFill.Blip != nil {
										_ecae(_bgeg.BlipFill.Blip, _cggg)
									}
								default:
								}
							}
						}
						for _, _cac := range _gade.Inline {
							for _, _fgbe := range _cac.Graphic.GraphicData.Any {
								switch _edgf := _fgbe.(type) {
								case *_gd.Pic:
									if _edgf.BlipFill != nil && _edgf.BlipFill.Blip != nil {
										_ecae(_edgf.BlipFill.Blip, _cggg)
									}
								default:
								}
							}
						}
					}
				}
			}
		}
	}
}

// Text returns text from the document as one string separated with line breaks.
func (_ebaed *DocText) Text() string {
	_fegb := _cb.NewBuffer([]byte{})
	for _, _cefac := range _ebaed.Items {
		if _cefac.Text != "" {
			_fegb.WriteString(_cefac.Text)
			_fegb.WriteString("\u000a")
		}
	}
	return _fegb.String()
}

// AddTab adds tab to a run and can be used with the the Paragraph's tab stops.
func (_cbddg Run) AddTab() { _eafb := _cbddg.newIC(); _eafb.Tab = _eaba.NewCT_Empty() }

// X return element of Node as interface, can be either *Paragraph, *Table and Run.
func (_aedb *Node) X() interface{} { return _aedb._efag }

// StyleID returns the style ID.
func (_ceac Style) StyleID() string {
	if _ceac._bgfca.StyleIdAttr == nil {
		return ""
	}
	return *_ceac._bgfca.StyleIdAttr
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (_bfg CellBorders) SetBottom(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_bfg._ggf.Bottom = _eaba.NewCT_Border()
	_gfdf(_bfg._ggf.Bottom, t, c, thickness)
}

// AddImageRef add ImageRef to header as relationship, returning ImageRef
// that can be used to be placed as header content.
func (_acae Header) AddImageRef(r _ea.ImageRef) (_ea.ImageRef, error) {
	var _bfadc _ea.Relationships
	for _egcc, _fdcg := range _acae._bffa._bbac {
		if _fdcg == _acae._cbgbf {
			_bfadc = _acae._bffa._cfg[_egcc]
		}
	}
	_fcbd := _bfadc.AddRelationship(r.Target(), _d.ImageType)
	r.SetRelID(_fcbd.X().IdAttr)
	return r, nil
}

// X returns the inner wrapped XML type.
func (_adfb Footer) X() *_eaba.Ftr { return _adfb._daaa }

func (_bbbc *WatermarkPicture) getInnerElement(_ageba string) *_d.XSDAny {
	for _, _gfae := range _bbbc._dadb.Any {
		_fbecc, _fcgec := _gfae.(*_d.XSDAny)
		if _fcgec && (_fbecc.XMLName.Local == _ageba || _fbecc.XMLName.Local == "\u0076\u003a"+_ageba) {
			return _fbecc
		}
	}
	return nil
}

func _gfdf(_bgca *_eaba.CT_Border, _ddbf _eaba.ST_Border, _cecd _bc.Color, _ebacb _eb.Distance) {
	_bgca.ValAttr = _ddbf
	_bgca.ColorAttr = &_eaba.ST_HexColor{}
	if _cecd.IsAuto() {
		_bgca.ColorAttr.ST_HexColorAuto = _eaba.ST_HexColorAutoAuto
	} else {
		_bgca.ColorAttr.ST_HexColorRGB = _cecd.AsRGBString()
	}
	if _ebacb != _eb.Zero {
		_bgca.SzAttr = _d.Uint64(uint64(_ebacb / _eb.Point * 8))
	}
}

// Properties returns the cell properties.
func (_gbc Cell) Properties() CellProperties {
	if _gbc._acg.TcPr == nil {
		_gbc._acg.TcPr = _eaba.NewCT_TcPr()
	}
	return CellProperties{_gbc._acg.TcPr}
}

// SetLeftPct sets the cell left margin
func (_ceb CellMargins) SetLeftPct(pct float64) {
	_ceb._ec.Left = _eaba.NewCT_TblWidth()
	_cab(_ceb._ec.Left, pct)
}

// SetPossibleValues sets possible values for a FormFieldTypeDropDown.
func (_faac FormField) SetPossibleValues(values []string) {
	if _faac._cdag.DdList != nil {
		for _, _eagd := range values {
			_eebg := _eaba.NewCT_String()
			_eebg.ValAttr = _eagd
			_faac._cdag.DdList.ListEntry = append(_faac._cdag.DdList.ListEntry, _eebg)
		}
	}
}

// Append appends a document d0 to a document d1. All settings, headers and footers remain the same as in the document d0 if they exist there, otherwise they are taken from the d1.
func (_dabd *Document) Append(d1orig *Document) error {
	_gdga, _cfca := d1orig.Copy()
	if _cfca != nil {
		return _cfca
	}
	_dabd.DocBase = _dabd.DocBase.Append(_gdga.DocBase)
	if _gdga._eeda.ConformanceAttr != _cg.ST_ConformanceClassStrict {
		_dabd._eeda.ConformanceAttr = _gdga._eeda.ConformanceAttr
	}
	_ffed := _dabd._bdb.X().Relationship
	_fcdfa := _gdga._bdb.X().Relationship
	_gbabg := _gdga._eeda.Body
	_ggef := map[string]string{}
	_edfb := map[int64]int64{}
	_abag := map[int64]int64{}
	for _, _adaga := range _fcdfa {
		_eagf := true
		_egac := _adaga.IdAttr
		_agba := _adaga.TargetAttr
		_bgec := _adaga.TypeAttr
		_ceag := _bgec == _d.ImageType
		_fgaf := _bgec == _d.HyperLinkType
		var _gegg string
		for _, _bdaf := range _ffed {
			if _bdaf.TypeAttr == _bgec && _bdaf.TargetAttr == _agba {
				_eagf = false
				_gegg = _bdaf.IdAttr
				break
			}
		}
		if _ceag {
			_gdeg := "\u0077\u006f\u0072d\u002f" + _agba
			for _, _cfag := range _gdga.DocBase.Images {
				if _cfag.Target() == _gdeg {
					_eece, _gbcg := _ea.ImageFromStorage(_cfag.Path())
					if _gbcg != nil {
						return _gbcg
					}
					_bgbda, _gbcg := _dabd.AddImage(_eece)
					if _gbcg != nil {
						return _gbcg
					}
					_gegg = _bgbda.RelID()
					break
				}
			}
		} else if _eagf {
			if _fgaf {
				_eafc := _dabd._bdb.AddHyperlink(_agba)
				_gegg = _ea.Relationship(_eafc).ID()
			} else {
				_becg := _dabd._bdb.AddRelationship(_agba, _bgec)
				_gegg = _becg.X().IdAttr
			}
		}
		if _egac != _gegg {
			_ggef[_egac] = _gegg
		}
	}
	if _gbabg.SectPr != nil {
		for _, _cdbd := range _gbabg.SectPr.EG_HdrFtrReferences {
			if _cdbd.HeaderReference != nil {
				if _fdcf, _bfcc := _ggef[_cdbd.HeaderReference.IdAttr]; _bfcc {
					_cdbd.HeaderReference.IdAttr = _fdcf
					_dabd._cfg = append(_dabd._cfg, _ea.NewRelationships())
				}
			} else if _cdbd.FooterReference != nil {
				if _cbade, _bbc := _ggef[_cdbd.FooterReference.IdAttr]; _bbc {
					_cdbd.FooterReference.IdAttr = _cbade
					_dabd._efc = append(_dabd._efc, _ea.NewRelationships())
				}
			}
		}
	}
	_bgfcf, _afdf := _dabd._bfd, _gdga._bfd
	if _bgfcf != nil {
		if _afdf != nil {
			if _bgfcf.Endnote != nil {
				if _afdf.Endnote != nil {
					_gga := int64(len(_bgfcf.Endnote) + 1)
					for _, _eccf := range _afdf.Endnote {
						_aebd := _eccf.IdAttr
						if _aebd > 0 {
							_eccf.IdAttr = _gga
							_bgfcf.Endnote = append(_bgfcf.Endnote, _eccf)
							_abag[_aebd] = _gga
							_gga++
						}
					}
				}
			} else {
				_bgfcf.Endnote = _afdf.Endnote
			}
		}
	} else if _afdf != nil {
		_bgfcf = _afdf
	}
	_dabd._bfd = _bgfcf
	_bfgca, _aece := _dabd._edf, _gdga._edf
	if _bfgca != nil {
		if _aece != nil {
			if _bfgca.Footnote != nil {
				if _aece.Footnote != nil {
					_cfdb := int64(len(_bfgca.Footnote) + 1)
					for _, _cdea := range _aece.Footnote {
						_gcefg := _cdea.IdAttr
						if _gcefg > 0 {
							_cdea.IdAttr = _cfdb
							_bfgca.Footnote = append(_bfgca.Footnote, _cdea)
							_edfb[_gcefg] = _cfdb
							_cfdb++
						}
					}
				}
			} else {
				_bfgca.Footnote = _aece.Footnote
			}
		}
	} else if _aece != nil {
		_bfgca = _aece
	}
	_dabd._edf = _bfgca
	for _, _dbaa := range _gbabg.EG_BlockLevelElts {
		for _, _feda := range _dbaa.EG_ContentBlockContent {
			for _, _ddfc := range _feda.P {
				_eac(_ddfc, _ggef)
				_fffb(_ddfc, _ggef)
				_gfff(_ddfc, _edfb, _abag)
			}
			for _, _ffcbe := range _feda.Tbl {
				_eaaec(_ffcbe, _ggef)
				_dcadc(_ffcbe, _ggef)
				_edbc(_ffcbe, _edfb, _abag)
			}
		}
	}
	_dabd._eeda.Body.EG_BlockLevelElts = append(_dabd._eeda.Body.EG_BlockLevelElts, _gdga._eeda.Body.EG_BlockLevelElts...)
	if _dabd._eeda.Body.SectPr == nil {
		_dabd._eeda.Body.SectPr = _gdga._eeda.Body.SectPr
	} else {
		var _aabb, _ffba bool
		for _, _fcad := range _dabd._eeda.Body.SectPr.EG_HdrFtrReferences {
			if _fcad.HeaderReference != nil {
				_aabb = true
			} else if _fcad.FooterReference != nil {
				_ffba = true
			}
		}
		if !_aabb {
			for _, _cegf := range _gdga._eeda.Body.SectPr.EG_HdrFtrReferences {
				if _cegf.HeaderReference != nil {
					_dabd._eeda.Body.SectPr.EG_HdrFtrReferences = append(_dabd._eeda.Body.SectPr.EG_HdrFtrReferences, _cegf)
					break
				}
			}
		}
		if !_ffba {
			for _, _baaa := range _gdga._eeda.Body.SectPr.EG_HdrFtrReferences {
				if _baaa.FooterReference != nil {
					_dabd._eeda.Body.SectPr.EG_HdrFtrReferences = append(_dabd._eeda.Body.SectPr.EG_HdrFtrReferences, _baaa)
					break
				}
			}
		}
		if _dabd._eeda.Body.SectPr.Cols == nil && _gdga._eeda.Body.SectPr.Cols != nil {
			_dabd._eeda.Body.SectPr.Cols = _gdga._eeda.Body.SectPr.Cols
		}
	}
	_ccfd := _dabd.Numbering._ebbdb
	_gbce := _gdga.Numbering._ebbdb
	if _ccfd != nil {
		if _gbce != nil {
			_ccfd.NumPicBullet = append(_ccfd.NumPicBullet, _gbce.NumPicBullet...)
			_ccfd.AbstractNum = append(_ccfd.AbstractNum, _gbce.AbstractNum...)
			_ccfd.Num = append(_ccfd.Num, _gbce.Num...)
		}
	} else if _gbce != nil {
		_ccfd = _gbce
	}
	_dabd.Numbering._ebbdb = _ccfd
	if _dabd.Styles._abad == nil && _gdga.Styles._abad != nil {
		_dabd.Styles._abad = _gdga.Styles._abad
	}
	_dabd._bag = append(_dabd._bag, _gdga._bag...)
	_dabd._aac = append(_dabd._aac, _gdga._aac...)
	if len(_dabd._bbac) == 0 {
		_dabd._bbac = _gdga._bbac
	}
	if len(_dabd._gfdb) == 0 {
		_dabd._gfdb = _gdga._gfdb
	}
	_fgaff := _dabd._feb
	_bdba := _gdga._feb
	if _fgaff != nil {
		if _bdba != nil {
			if _fgaff.Divs != nil {
				if _bdba.Divs != nil {
					_fgaff.Divs.Div = append(_fgaff.Divs.Div, _bdba.Divs.Div...)
				}
			} else {
				_fgaff.Divs = _bdba.Divs
			}
		}
		_fgaff.Frameset = nil
	} else if _bdba != nil {
		_fgaff = _bdba
		_fgaff.Frameset = nil
	}
	_dabd._feb = _fgaff
	_dcgb := _dabd._fdd
	_dggg := _gdga._fdd
	if _dcgb != nil {
		if _dggg != nil {
			if _dcgb.Font != nil {
				if _dggg.Font != nil {
					for _, _cfcaf := range _dggg.Font {
						_bcbb := true
						for _, _cegad := range _dcgb.Font {
							if _cegad.NameAttr == _cfcaf.NameAttr {
								_bcbb = false
								break
							}
						}
						if _bcbb {
							_dcgb.Font = append(_dcgb.Font, _cfcaf)
						}
					}
				}
			} else {
				_dcgb.Font = _dggg.Font
			}
		}
	} else if _dggg != nil {
		_dcgb = _dggg
	}
	_dabd._fdd = _dcgb
	return nil
}

// AddTextInput adds text input form field to the paragraph and returns it.
func (_facg Paragraph) AddTextInput(name string) FormField {
	_dead := _facg.addFldCharsForField(name, "\u0046\u004f\u0052\u004d\u0054\u0045\u0058\u0054")
	_dead._cdag.TextInput = _eaba.NewCT_FFTextInput()
	return _dead
}

// SetTextWrapBehindText sets the text wrap to behind text.
func (_df AnchoredDrawing) SetTextWrapBehindText() {
	_df._cc.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_df._cc.Choice.WrapNone = _eaba.NewWdCT_WrapNone()
	_df._cc.BehindDocAttr = true
	_df._cc.LayoutInCellAttr = true
	_df._cc.AllowOverlapAttr = true
}

func (_gdff Footnote) id() int64 { return _gdff._efba.IdAttr }

// SetPageMargins sets the page margins for a section
func (_edgfe Section) SetPageMargins(top, right, bottom, left, header, footer, gutter _eb.Distance) {
	_dbdde := _eaba.NewCT_PageMar()
	_dbdde.TopAttr.Int64 = _d.Int64(int64(top / _eb.Twips))
	_dbdde.BottomAttr.Int64 = _d.Int64(int64(bottom / _eb.Twips))
	_dbdde.RightAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(right / _eb.Twips))
	_dbdde.LeftAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(left / _eb.Twips))
	_dbdde.HeaderAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(header / _eb.Twips))
	_dbdde.FooterAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(footer / _eb.Twips))
	_dbdde.GutterAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(gutter / _eb.Twips))
	_edgfe._aggca.PgMar = _dbdde
}

// InsertRowAfter inserts a row after another row
func (_degaf Table) InsertRowAfter(r Row) Row {
	for _gffdg, _abfe := range _degaf._fecae.EG_ContentRowContent {
		if len(_abfe.Tr) > 0 && r.X() == _abfe.Tr[0] {
			_cada := _eaba.NewEG_ContentRowContent()
			if len(_degaf._fecae.EG_ContentRowContent) < _gffdg+2 {
				return _degaf.AddRow()
			}
			_degaf._fecae.EG_ContentRowContent = append(_degaf._fecae.EG_ContentRowContent, nil)
			copy(_degaf._fecae.EG_ContentRowContent[_gffdg+2:], _degaf._fecae.EG_ContentRowContent[_gffdg+1:])
			_degaf._fecae.EG_ContentRowContent[_gffdg+1] = _cada
			_faedd := _eaba.NewCT_Row()
			_cada.Tr = append(_cada.Tr, _faedd)
			return Row{_degaf._bggbd, _faedd}
		}
	}
	return _degaf.AddRow()
}

// Settings controls the document settings.
type Settings struct{ _ecefe *_eaba.Settings }

// SearchStyleByName return style by its name.
func (_cfea Styles) SearchStyleByName(name string) (Style, bool) {
	for _, _ddfcd := range _cfea._abad.Style {
		if _ddfcd.Name != nil {
			if _ddfcd.Name.ValAttr == name {
				return Style{_ddfcd}, true
			}
		}
	}
	return Style{}, false
}

// AddRow adds a row to a table.
func (_efda Table) AddRow() Row {
	_cdbdf := _eaba.NewEG_ContentRowContent()
	_efda._fecae.EG_ContentRowContent = append(_efda._fecae.EG_ContentRowContent, _cdbdf)
	_ebbea := _eaba.NewCT_Row()
	_cdbdf.Tr = append(_cdbdf.Tr, _ebbea)
	return Row{_efda._bggbd, _ebbea}
}

// Table is a table within a document.
type Table struct {
	_bggbd *Document
	_fecae *_eaba.CT_Tbl
}

// Shadow returns true if paragraph shadow is on.
func (_dfbgc ParagraphProperties) Shadow() bool { return _beff(_dfbgc._efec.RPr.Shadow) }

func (_cfa *Document) tables(_bdfd *_eaba.EG_ContentBlockContent) []Table {
	_afg := []Table{}
	for _, _dcda := range _bdfd.Tbl {
		_afg = append(_afg, Table{_cfa, _dcda})
		for _, _eeg := range _dcda.EG_ContentRowContent {
			for _, _bga := range _eeg.Tr {
				for _, _faa := range _bga.EG_ContentCellContent {
					for _, _cba := range _faa.Tc {
						for _, _agbb := range _cba.EG_BlockLevelElts {
							for _, _deb := range _agbb.EG_ContentBlockContent {
								for _, _fggc := range _cfa.tables(_deb) {
									_afg = append(_afg, _fggc)
								}
							}
						}
					}
				}
			}
		}
	}
	return _afg
}

func (_ccbeg Paragraph) insertRun(_dacdb Run, _cdbf bool) Run {
	for _, _gdfc := range _ccbeg._dbef.EG_PContent {
		for _fcdfe, _bcfd := range _gdfc.EG_ContentRunContent {
			if _bcfd.R == _dacdb.X() {
				_aeee := _eaba.NewCT_R()
				_gdfc.EG_ContentRunContent = append(_gdfc.EG_ContentRunContent, nil)
				if _cdbf {
					copy(_gdfc.EG_ContentRunContent[_fcdfe+1:], _gdfc.EG_ContentRunContent[_fcdfe:])
					_gdfc.EG_ContentRunContent[_fcdfe] = _eaba.NewEG_ContentRunContent()
					_gdfc.EG_ContentRunContent[_fcdfe].R = _aeee
				} else {
					copy(_gdfc.EG_ContentRunContent[_fcdfe+2:], _gdfc.EG_ContentRunContent[_fcdfe+1:])
					_gdfc.EG_ContentRunContent[_fcdfe+1] = _eaba.NewEG_ContentRunContent()
					_gdfc.EG_ContentRunContent[_fcdfe+1].R = _aeee
				}
				return Run{_ccbeg._bbgc, _aeee}
			}
			if _bcfd.Sdt != nil && _bcfd.Sdt.SdtContent != nil {
				for _, _fegd := range _bcfd.Sdt.SdtContent.EG_ContentRunContent {
					if _fegd.R == _dacdb.X() {
						_ffdbf := _eaba.NewCT_R()
						_bcfd.Sdt.SdtContent.EG_ContentRunContent = append(_bcfd.Sdt.SdtContent.EG_ContentRunContent, nil)
						if _cdbf {
							copy(_bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe+1:], _bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe:])
							_bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe] = _eaba.NewEG_ContentRunContent()
							_bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe].R = _ffdbf
						} else {
							copy(_bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe+2:], _bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe+1:])
							_bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe+1] = _eaba.NewEG_ContentRunContent()
							_bcfd.Sdt.SdtContent.EG_ContentRunContent[_fcdfe+1].R = _ffdbf
						}
						return Run{_ccbeg._bbgc, _ffdbf}
					}
				}
			}
		}
	}
	return _ccbeg.AddRun()
}

// X returns the inner wrapped XML type.
func (_bdc *Document) X() *_eaba.Document { return _bdc._eeda }

// SetOrigin sets the origin of the image.  It defaults to ST_RelFromHPage and
// ST_RelFromVPage
func (_ead AnchoredDrawing) SetOrigin(h _eaba.WdST_RelFromH, v _eaba.WdST_RelFromV) {
	_ead._cc.PositionH.RelativeFromAttr = h
	_ead._cc.PositionV.RelativeFromAttr = v
}

// Index returns the index of the footer within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (_cfbf Footer) Index() int {
	for _abcc, _fcga := range _cfbf._dec._gfdb {
		if _fcga == _cfbf._daaa {
			return _abcc
		}
	}
	return -1
}

// Underline returns the type of run underline.
func (_bacfc RunProperties) Underline() _eaba.ST_Underline {
	if _eafcg := _bacfc._gfbgc.U; _eafcg != nil {
		return _eafcg.ValAttr
	}
	return 0
}

func (_ebcb Paragraph) addStartBookmark(_bgaf int64, _bgfbc string) *_eaba.CT_Bookmark {
	_gfceg := _eaba.NewEG_PContent()
	_ebcb._dbef.EG_PContent = append(_ebcb._dbef.EG_PContent, _gfceg)
	_bfgcc := _eaba.NewEG_ContentRunContent()
	_ebfg := _eaba.NewEG_RunLevelElts()
	_ebbde := _eaba.NewEG_RangeMarkupElements()
	_dfecd := _eaba.NewCT_Bookmark()
	_dfecd.NameAttr = _bgfbc
	_dfecd.IdAttr = _bgaf
	_ebbde.BookmarkStart = _dfecd
	_gfceg.EG_ContentRunContent = append(_gfceg.EG_ContentRunContent, _bfgcc)
	_bfgcc.EG_RunLevelElts = append(_bfgcc.EG_RunLevelElts, _ebfg)
	_ebfg.EG_RangeMarkupElements = append(_ebfg.EG_RangeMarkupElements, _ebbde)
	return _dfecd
}

func (_ffeg Paragraph) ensurePPr() {
	if _ffeg._dbef.PPr == nil {
		_ffeg._dbef.PPr = _eaba.NewCT_PPr()
	}
}

// TableProperties are the properties for a table within a document
type TableProperties struct{ _efac *_eaba.CT_TblPr }

// Name returns the name of the field.
func (_fdccf FormField) Name() string { return *_fdccf._cdag.Name[0].ValAttr }

// Strike returns true if paragraph is striked.
func (_dbce ParagraphProperties) Strike() bool { return _beff(_dbce._efec.RPr.Strike) }

// GetNumberingLevelByIds returns a NumberingLevel by its NumId and LevelId attributes
// or an empty one if not found.
func (_aaga *Document) GetNumberingLevelByIds(numId, levelId int64) NumberingLevel {
	for _, _eeac := range _aaga.Numbering._ebbdb.Num {
		if _eeac != nil && _eeac.NumIdAttr == numId {
			_adde := _eeac.AbstractNumId.ValAttr
			for _, _gecbb := range _aaga.Numbering._ebbdb.AbstractNum {
				if _gecbb.AbstractNumIdAttr == _adde {
					for _, _fggf := range _gecbb.Lvl {
						if _fggf.IlvlAttr == levelId {
							return NumberingLevel{_fggf}
						}
					}
				}
			}
		}
	}
	return NumberingLevel{}
}

// Spacing returns the paragraph spacing settings.
func (_dfab ParagraphProperties) Spacing() ParagraphSpacing {
	if _dfab._efec.Spacing == nil {
		_dfab._efec.Spacing = _eaba.NewCT_Spacing()
	}
	return ParagraphSpacing{_dfab._efec.Spacing}
}

// Fonts allows manipulating a style or run's fonts.
type Fonts struct{ _aggeg *_eaba.CT_Fonts }

// SetDefaultValue sets the default value of a FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (_dacd FormField) SetDefaultValue(v string) {
	if _dacd._cdag.DdList != nil {
		for _abgab, _fggg := range _dacd.PossibleValues() {
			if _fggg == v {
				_dacd._cdag.DdList.Default = _eaba.NewCT_DecimalNumber()
				_dacd._cdag.DdList.Default.ValAttr = int64(_abgab)
				break
			}
		}
	}
}

// ExtractFromHeader returns text from the document header as an array of TextItems.
func ExtractFromHeader(header *_eaba.Hdr) []TextItem {
	return _bddda(header.EG_ContentBlockContent, nil)
}

// NewNumbering constructs a new numbering.
func NewNumbering() Numbering { _aaae := _eaba.NewNumbering(); return Numbering{_aaae} }

// Type returns the type of the field.
func (_edaa FormField) Type() FormFieldType {
	if _edaa._cdag.TextInput != nil {
		return FormFieldTypeText
	} else if _edaa._cdag.CheckBox != nil {
		return FormFieldTypeCheckBox
	} else if _edaa._cdag.DdList != nil {
		return FormFieldTypeDropDown
	}
	return FormFieldTypeUnknown
}

// SetColor sets a specific color or auto.
func (_bgb Color) SetColor(v _bc.Color) {
	if v.IsAuto() {
		_bgb._bcf.ValAttr.ST_HexColorAuto = _eaba.ST_HexColorAutoAuto
		_bgb._bcf.ValAttr.ST_HexColorRGB = nil
	} else {
		_bgb._bcf.ValAttr.ST_HexColorAuto = _eaba.ST_HexColorAutoUnset
		_bgb._bcf.ValAttr.ST_HexColorRGB = v.AsRGBString()
	}
}

func _bbbb() *_aae.OfcLock {
	_ebee := _aae.NewOfcLock()
	_ebee.ExtAttr = _aae.ST_ExtEdit
	_ebee.AspectratioAttr = _cg.ST_TrueFalseTrue
	return _ebee
}

// SetSize sets the size of the displayed image on the page.
func (_bbfe InlineDrawing) SetSize(w, h _eb.Distance) {
	_bbfe._fbbd.Extent.CxAttr = int64(float64(w*_eb.Pixel72) / _eb.EMU)
	_bbfe._fbbd.Extent.CyAttr = int64(float64(h*_eb.Pixel72) / _eb.EMU)
}

// SetAllowOverlapAttr sets the allowOverlap attribute of anchor.
func (_gfd AnchoredDrawing) SetAllowOverlapAttr(val bool) { _gfd._cc.AllowOverlapAttr = val }

// SetOutlineLvl sets outline level of paragraph.
func (_aaabe Paragraph) SetOutlineLvl(lvl int64) {
	_aaabe.ensurePPr()
	if _aaabe._dbef.PPr.OutlineLvl == nil {
		_aaabe._dbef.PPr.OutlineLvl = _eaba.NewCT_DecimalNumber()
	}
	_eggb := lvl - 1
	_aaabe._dbef.PPr.OutlineLvl.ValAttr = _eggb
}

// SetThemeColor sets the color from the theme.
func (_gfe Color) SetThemeColor(t _eaba.ST_ThemeColor) { _gfe._bcf.ThemeColorAttr = t }

// SetOffset sets the offset of the image relative to the origin, which by
// default this is the top-left corner of the page. Offset is incompatible with
// SetAlignment, whichever is called last is applied.
func (_gc AnchoredDrawing) SetOffset(x, y _eb.Distance) { _gc.SetXOffset(x); _gc.SetYOffset(y) }

// SetShadow sets the run to shadowed text.
func (_dfef RunProperties) SetShadow(b bool) {
	if !b {
		_dfef._gfbgc.Shadow = nil
	} else {
		_dfef._gfbgc.Shadow = _eaba.NewCT_OnOff()
	}
}

// RemoveRun removes a child run from a paragraph.
func (_bcfge Paragraph) RemoveRun(r Run) {
	for _, _gccbdb := range _bcfge._dbef.EG_PContent {
		for _edeg, _eeag := range _gccbdb.EG_ContentRunContent {
			if _eeag.R == r._bggd {
				copy(_gccbdb.EG_ContentRunContent[_edeg:], _gccbdb.EG_ContentRunContent[_edeg+1:])
				_gccbdb.EG_ContentRunContent = _gccbdb.EG_ContentRunContent[0 : len(_gccbdb.EG_ContentRunContent)-1]
			}
			if _eeag.Sdt != nil && _eeag.Sdt.SdtContent != nil {
				for _afbg, _aabc := range _eeag.Sdt.SdtContent.EG_ContentRunContent {
					if _aabc.R == r._bggd {
						copy(_eeag.Sdt.SdtContent.EG_ContentRunContent[_afbg:], _eeag.Sdt.SdtContent.EG_ContentRunContent[_afbg+1:])
						_eeag.Sdt.SdtContent.EG_ContentRunContent = _eeag.Sdt.SdtContent.EG_ContentRunContent[0 : len(_eeag.Sdt.SdtContent.EG_ContentRunContent)-1]
					}
				}
			}
		}
	}
}

func _gcbef() *_aae.OfcLock {
	_afbgg := _aae.NewOfcLock()
	_afbgg.ExtAttr = _aae.ST_ExtEdit
	_afbgg.TextAttr = _cg.ST_TrueFalseTrue
	_afbgg.ShapetypeAttr = _cg.ST_TrueFalseTrue
	return _afbgg
}

// SetShading controls the cell shading.
func (_gea CellProperties) SetShading(shd _eaba.ST_Shd, foreground, fill _bc.Color) {
	if shd == _eaba.ST_ShdUnset {
		_gea._be.Shd = nil
	} else {
		_gea._be.Shd = _eaba.NewCT_Shd()
		_gea._be.Shd.ValAttr = shd
		_gea._be.Shd.ColorAttr = &_eaba.ST_HexColor{}
		if foreground.IsAuto() {
			_gea._be.Shd.ColorAttr.ST_HexColorAuto = _eaba.ST_HexColorAutoAuto
		} else {
			_gea._be.Shd.ColorAttr.ST_HexColorRGB = foreground.AsRGBString()
		}
		_gea._be.Shd.FillAttr = &_eaba.ST_HexColor{}
		if fill.IsAuto() {
			_gea._be.Shd.FillAttr.ST_HexColorAuto = _eaba.ST_HexColorAutoAuto
		} else {
			_gea._be.Shd.FillAttr.ST_HexColorRGB = fill.AsRGBString()
		}
	}
}

// ParagraphStyleProperties is the styling information for a paragraph.
type ParagraphStyleProperties struct{ _fded *_eaba.CT_PPrGeneral }

func (_fcfb Paragraph) addEndFldChar() *_eaba.CT_FldChar {
	_dddcd := _fcfb.addFldChar()
	_dddcd.FldCharTypeAttr = _eaba.ST_FldCharTypeEnd
	return _dddcd
}

// Clear clears all content within a footer
func (_fdaf Footer) Clear() { _fdaf._daaa.EG_ContentBlockContent = nil }

// TableWidth controls width values in table settings.
type TableWidth struct{ _ebga *_eaba.CT_TblWidth }

// Style returns the style for a paragraph, or an empty string if it is unset.
func (_dgggg ParagraphProperties) Style() string {
	if _dgggg._efec.PStyle != nil {
		return _dgggg._efec.PStyle.ValAttr
	}
	return ""
}

func _bgdc(_eabdc *Document, _fbed Paragraph) listItemInfo {
	if _eabdc.Numbering.X() == nil {
		return listItemInfo{}
	}
	if len(_eabdc.Numbering.Definitions()) < 1 {
		return listItemInfo{}
	}
	_gcegg := _adff(_fbed)
	if _gcegg == nil {
		return listItemInfo{}
	}
	_afea := _eabdc.GetNumberingLevelByIds(_gcegg.NumId.ValAttr, _gcegg.Ilvl.ValAttr)
	if _afea.X().LvlText.ValAttr == nil {
		return listItemInfo{}
	}
	_cgbbf := int64(0)
	for _, _dedg := range _eabdc.Numbering._ebbdb.Num {
		if _dedg != nil && _dedg.NumIdAttr == _gcegg.NumId.ValAttr {
			_cgbbf = _dedg.AbstractNumId.ValAttr
		}
	}
	return listItemInfo{FromParagraph: &_fbed, AbstractNumId: &_cgbbf, NumberingLevel: &_afea}
}

// SetWidthAuto sets the the cell width to automatic.
func (_ccg CellProperties) SetWidthAuto() {
	_ccg._be.TcW = _eaba.NewCT_TblWidth()
	_ccg._be.TcW.TypeAttr = _eaba.ST_TblWidthAuto
}

func (_gfeba *Document) insertParagraph(_dcca Paragraph, _gcef bool) Paragraph {
	if _gfeba._eeda.Body == nil {
		return _gfeba.AddParagraph()
	}
	_gec := _dcca.X()
	for _, _cfbd := range _gfeba._eeda.Body.EG_BlockLevelElts {
		for _, _gdcb := range _cfbd.EG_ContentBlockContent {
			for _dcfd, _cgee := range _gdcb.P {
				if _cgee == _gec {
					_afggg := _eaba.NewCT_P()
					_gdcb.P = append(_gdcb.P, nil)
					if _gcef {
						copy(_gdcb.P[_dcfd+1:], _gdcb.P[_dcfd:])
						_gdcb.P[_dcfd] = _afggg
					} else {
						copy(_gdcb.P[_dcfd+2:], _gdcb.P[_dcfd+1:])
						_gdcb.P[_dcfd+1] = _afggg
					}
					return Paragraph{_gfeba, _afggg}
				}
			}
			for _, _gfcb := range _gdcb.Tbl {
				for _, _bccb := range _gfcb.EG_ContentRowContent {
					for _, _fdbc := range _bccb.Tr {
						for _, _eefd := range _fdbc.EG_ContentCellContent {
							for _, _efeeg := range _eefd.Tc {
								for _, _cdefa := range _efeeg.EG_BlockLevelElts {
									for _, _cfe := range _cdefa.EG_ContentBlockContent {
										for _fbeb, _dgee := range _cfe.P {
											if _dgee == _gec {
												_dddb := _eaba.NewCT_P()
												_cfe.P = append(_cfe.P, nil)
												if _gcef {
													copy(_cfe.P[_fbeb+1:], _cfe.P[_fbeb:])
													_cfe.P[_fbeb] = _dddb
												} else {
													copy(_cfe.P[_fbeb+2:], _cfe.P[_fbeb+1:])
													_cfe.P[_fbeb+1] = _dddb
												}
												return Paragraph{_gfeba, _dddb}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			if _gdcb.Sdt != nil && _gdcb.Sdt.SdtContent != nil && _gdcb.Sdt.SdtContent.P != nil {
				for _gcbd, _geac := range _gdcb.Sdt.SdtContent.P {
					if _geac == _gec {
						_bfefb := _eaba.NewCT_P()
						_gdcb.Sdt.SdtContent.P = append(_gdcb.Sdt.SdtContent.P, nil)
						if _gcef {
							copy(_gdcb.Sdt.SdtContent.P[_gcbd+1:], _gdcb.Sdt.SdtContent.P[_gcbd:])
							_gdcb.Sdt.SdtContent.P[_gcbd] = _bfefb
						} else {
							copy(_gdcb.Sdt.SdtContent.P[_gcbd+2:], _gdcb.Sdt.SdtContent.P[_gcbd+1:])
							_gdcb.Sdt.SdtContent.P[_gcbd+1] = _bfefb
						}
						return Paragraph{_gfeba, _bfefb}
					}
				}
			}
		}
	}
	return _gfeba.AddParagraph()
}

// SetValue sets the value of a FormFieldTypeText or FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (_fbfc FormField) SetValue(v string) {
	if _fbfc._cdag.DdList != nil {
		for _geaf, _cgfb := range _fbfc.PossibleValues() {
			if _cgfb == v {
				_fbfc._cdag.DdList.Result = _eaba.NewCT_DecimalNumber()
				_fbfc._cdag.DdList.Result.ValAttr = int64(_geaf)
				break
			}
		}
	} else if _fbfc._cdag.TextInput != nil {
		_fbfc._bebe.T = _eaba.NewCT_Text()
		_fbfc._bebe.T.Content = v
	}
}

func (_abb *Document) validateTableCells() error {
	for _, _gabeb := range _abb._eeda.Body.EG_BlockLevelElts {
		for _, _fddb := range _gabeb.EG_ContentBlockContent {
			for _, _egfg := range _fddb.Tbl {
				for _, _afee := range _egfg.EG_ContentRowContent {
					for _, _dceab := range _afee.Tr {
						_eabag := false
						for _, _fcgc := range _dceab.EG_ContentCellContent {
							_dcad := false
							for _, _egbb := range _fcgc.Tc {
								_eabag = true
								for _, _eceg := range _egbb.EG_BlockLevelElts {
									for _, _ffbc := range _eceg.EG_ContentBlockContent {
										if len(_ffbc.P) > 0 {
											_dcad = true
											break
										}
									}
								}
							}
							if !_dcad {
								return _ffa.New("t\u0061\u0062\u006c\u0065\u0020\u0063e\u006c\u006c\u0020\u006d\u0075\u0073t\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061p\u0068")
							}
						}
						if !_eabag {
							return _ffa.New("\u0074\u0061b\u006c\u0065\u0020\u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020ce\u006c\u006c")
						}
					}
				}
			}
		}
	}
	return nil
}

// HyperLink is a link within a document.
type HyperLink struct {
	_bcgg  *Document
	_dgace *_eaba.CT_Hyperlink
}

// SetASCIITheme sets the font ASCII Theme.
func (_bdaa Fonts) SetASCIITheme(t _eaba.ST_Theme) { _bdaa._aggeg.AsciiThemeAttr = t }

// NewStyles constructs a new empty Styles
func NewStyles() Styles { return Styles{_eaba.NewStyles()} }

// Bookmark is a bookmarked location within a document that can be referenced
// with a hyperlink.
type Bookmark struct{ _dda *_eaba.CT_Bookmark }

// SetColumnSpan sets the number of Grid Columns Spanned by the Cell.  This is used
// to give the appearance of merged cells.
func (_bcc CellProperties) SetColumnSpan(cols int) {
	if cols == 0 {
		_bcc._be.GridSpan = nil
	} else {
		_bcc._be.GridSpan = _eaba.NewCT_DecimalNumber()
		_bcc._be.GridSpan.ValAttr = int64(cols)
	}
}

// InsertParagraphBefore adds a new empty paragraph before the relativeTo
// paragraph.
func (_daaf *Document) InsertParagraphBefore(relativeTo Paragraph) Paragraph {
	return _daaf.insertParagraph(relativeTo, true)
}

// MailMerge finds mail merge fields and replaces them with the text provided.  It also removes
// the mail merge source info from the document settings.
func (_aafa *Document) MailMerge(mergeContent map[string]string) {
	_adeed := _aafa.mergeFields()
	_ccgab := map[Paragraph][]Run{}
	for _, _cbbdb := range _adeed {
		_cafcd, _gefe := mergeContent[_cbbdb._ggce]
		if _gefe {
			if _cbbdb._aced {
				_cafcd = _f.ToUpper(_cafcd)
			} else if _cbbdb._gfcc {
				_cafcd = _f.ToLower(_cafcd)
			} else if _cbbdb._cbccd {
				_cafcd = _f.Title(_cafcd)
			} else if _cbbdb._cecf {
				_bgfgb := _cb.Buffer{}
				for _agdag, _ecgac := range _cafcd {
					if _agdag == 0 {
						_bgfgb.WriteRune(_ae.ToUpper(_ecgac))
					} else {
						_bgfgb.WriteRune(_ecgac)
					}
				}
				_cafcd = _bgfgb.String()
			}
			if _cafcd != "" && _cbbdb._ddceg != "" {
				_cafcd = _cbbdb._ddceg + _cafcd
			}
			if _cafcd != "" && _cbbdb._caaee != "" {
				_cafcd = _cafcd + _cbbdb._caaee
			}
		}
		if _cbbdb._fddee {
			if len(_cbbdb._acfca.FldSimple) == 1 && len(_cbbdb._acfca.FldSimple[0].EG_PContent) == 1 && len(_cbbdb._acfca.FldSimple[0].EG_PContent[0].EG_ContentRunContent) == 1 {
				_gaef := &_eaba.EG_ContentRunContent{}
				_gaef.R = _cbbdb._acfca.FldSimple[0].EG_PContent[0].EG_ContentRunContent[0].R
				_cbbdb._acfca.FldSimple = nil
				_ccgac := Run{_aafa, _gaef.R}
				_ccgac.ClearContent()
				_ccgac.AddText(_cafcd)
				_cbbdb._acfca.EG_ContentRunContent = append(_cbbdb._acfca.EG_ContentRunContent, _gaef)
			}
		} else {
			_acgf := _cbbdb._bfbb.Runs()
			for _bfeed := _cbbdb._abacb; _bfeed <= _cbbdb._feec; _bfeed++ {
				if _bfeed == _cbbdb._ddbag+1 {
					_acgf[_bfeed].ClearContent()
					_acgf[_bfeed].AddText(_cafcd)
				} else {
					_ccgab[_cbbdb._bfbb] = append(_ccgab[_cbbdb._bfbb], _acgf[_bfeed])
				}
			}
		}
	}
	for _ebed, _dgef := range _ccgab {
		for _, _acaa := range _dgef {
			_ebed.RemoveRun(_acaa)
		}
	}
	_aafa.Settings.RemoveMailMerge()
}

// RemoveParagraph removes a paragraph from the endnote.
func (_bfac Endnote) RemoveParagraph(p Paragraph) {
	for _, _ebfb := range _bfac.content() {
		for _gceb, _fbfg := range _ebfb.P {
			if _fbfg == p._dbef {
				copy(_ebfb.P[_gceb:], _ebfb.P[_gceb+1:])
				_ebfb.P = _ebfb.P[0 : len(_ebfb.P)-1]
				return
			}
		}
	}
}

// SetBeforeSpacing sets spacing above paragraph.
func (_agage Paragraph) SetBeforeSpacing(d _eb.Distance) {
	_agage.ensurePPr()
	if _agage._dbef.PPr.Spacing == nil {
		_agage._dbef.PPr.Spacing = _eaba.NewCT_Spacing()
	}
	_fdfbe := _agage._dbef.PPr.Spacing
	_fdfbe.BeforeAttr = &_cg.ST_TwipsMeasure{}
	_fdfbe.BeforeAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(d / _eb.Twips))
}

// SetThemeShade sets the shade based off the theme color.
func (_gba Color) SetThemeShade(s uint8) {
	_cbg := _eg.Sprintf("\u0025\u0030\u0032\u0078", s)
	_gba._bcf.ThemeShadeAttr = &_cbg
}

// SetWrapPathLineTo sets wrapPath lineTo value.
func (_ba AnchorDrawWrapOptions) SetWrapPathLineTo(coordinates []*_bb.CT_Point2D) {
	_ba._ga = coordinates
}

// SetTextWrapSquare sets the text wrap to square with a given wrap type.
func (_ge AnchoredDrawing) SetTextWrapSquare(t _eaba.WdST_WrapText) {
	_ge._cc.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_ge._cc.Choice.WrapSquare = _eaba.NewWdCT_WrapSquare()
	_ge._cc.Choice.WrapSquare.WrapTextAttr = t
}

// New constructs an empty document that content can be added to.
func New() *Document {
	_ecf := &Document{_eeda: _eaba.NewDocument()}
	_ecf.ContentTypes = _ea.NewContentTypes()
	_ecf._eeda.Body = _eaba.NewCT_Body()
	_ecf._eeda.ConformanceAttr = _cg.ST_ConformanceClassTransitional
	_ecf._bdb = _ea.NewRelationships()
	_ecf.AppProperties = _ea.NewAppProperties()
	_ecf.CoreProperties = _ea.NewCoreProperties()
	_ecf.ContentTypes.AddOverride("\u002fw\u006fr\u0064\u002f\u0064\u006f\u0063u\u006d\u0065n\u0074\u002e\u0078\u006d\u006c", "\u0061p\u0070\u006c\u0069c\u0061\u0074\u0069o\u006e/v\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072dp\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u006d\u006c\u002e\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u002e\u006d\u0061\u0069\u006e\u002bx\u006d\u006c")
	_ecf.Settings = NewSettings()
	_ecf._bdb.AddRelationship("\u0073\u0065\u0074t\u0069\u006e\u0067\u0073\u002e\u0078\u006d\u006c", _d.SettingsType)
	_ecf.ContentTypes.AddOverride("\u002fw\u006fr\u0064\u002f\u0073\u0065\u0074t\u0069\u006eg\u0073\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069o\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006dl\u0066\u006f\u0072\u006da\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c.\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002b\u0078\u006d\u006c")
	_ecf.Rels = _ea.NewRelationships()
	_ecf.Rels.AddRelationship(_d.RelativeFilename(_d.DocTypeDocument, "", _d.CorePropertiesType, 0), _d.CorePropertiesType)
	_ecf.Rels.AddRelationship("\u0064\u006fc\u0050\u0072\u006fp\u0073\u002f\u0061\u0070\u0070\u002e\u0078\u006d\u006c", _d.ExtendedPropertiesType)
	_ecf.Rels.AddRelationship("\u0077\u006f\u0072\u0064\u002f\u0064\u006f\u0063\u0075\u006d\u0065\u006et\u002e\u0078\u006d\u006c", _d.OfficeDocumentType)
	_ecf.Numbering = NewNumbering()
	_ecf.Numbering.InitializeDefault()
	_ecf.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072d/\u006e\u0075\u006d\u0062\u0065\u0072\u0069\u006e\u0067\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002f\u0076n\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063e\u0073\u0073\u0069\u006e\u0067\u006d\u006c\u002e\u006e\u0075\u006d\u0062e\u0072\u0069\u006e\u0067\u002b\u0078m\u006c")
	_ecf._bdb.AddRelationship("\u006e\u0075\u006d\u0062\u0065\u0072\u0069\u006e\u0067\u002e\u0078\u006d\u006c", _d.NumberingType)
	_ecf.Styles = NewStyles()
	_ecf.Styles.InitializeDefault()
	_ecf.ContentTypes.AddOverride("\u002f\u0077o\u0072\u0064\u002fs\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0073\u0074\u0079\u006ce\u0073\u002b\u0078\u006d\u006c")
	_ecf._bdb.AddRelationship("\u0073\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", _d.StylesType)
	_ecf._eeda.Body = _eaba.NewCT_Body()
	return _ecf
}

// SetFooter sets a section footer.
func (_ddg Section) SetFooter(f Footer, t _eaba.ST_HdrFtr) {
	_ffbd := _eaba.NewEG_HdrFtrReferences()
	_ddg._aggca.EG_HdrFtrReferences = append(_ddg._aggca.EG_HdrFtrReferences, _ffbd)
	_ffbd.FooterReference = _eaba.NewCT_HdrFtrRef()
	_ffbd.FooterReference.TypeAttr = t
	_eaecg := _ddg._ceec._bdb.FindRIDForN(f.Index(), _d.FooterType)
	if _eaecg == "" {
		_dcd.Log.Debug("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u0066\u006f\u006f\u0074\u0065r \u0049\u0044")
	}
	_ffbd.FooterReference.IdAttr = _eaecg
}

// SetCellSpacing sets the cell spacing within a table.
func (_cafd TableProperties) SetCellSpacing(m _eb.Distance) {
	_cafd._efac.TblCellSpacing = _eaba.NewCT_TblWidth()
	_cafd._efac.TblCellSpacing.TypeAttr = _eaba.ST_TblWidthDxa
	_cafd._efac.TblCellSpacing.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_cafd._efac.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_cafd._efac.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(m / _eb.Dxa))
}

// SetFollowImageShape sets wrapPath to follow image shape,
// if nil return wrapPath that follow image size.
func (_fbd AnchorDrawWrapOptions) SetFollowImageShape(val bool) {
	_fbd._fb = val
	if !val {
		_bf, _ed := _bg()
		_fbd._bba = _bf
		_fbd._ga = _ed
	}
}

// Copy makes a deep copy of the document by saving and reading it back.
// It can be useful to avoid sharing common data between two documents.
func (_gad *Document) Copy() (*Document, error) {
	_daba := _cb.NewBuffer([]byte{})
	_cdfc := _gad.save(_daba, _gad._ebd)
	if _cdfc != nil {
		return nil, _cdfc
	}
	_bfad := _daba.Bytes()
	_acbd := _cb.NewReader(_bfad)
	return _bed(_acbd, int64(_acbd.Len()), _gad._ebd)
}

// SetTableIndent sets the Table Indent from the Leading Margin
func (_cbef TableStyleProperties) SetTableIndent(ind _eb.Distance) {
	_cbef._ddafc.TblInd = _eaba.NewCT_TblWidth()
	_cbef._ddafc.TblInd.TypeAttr = _eaba.ST_TblWidthDxa
	_cbef._ddafc.TblInd.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_cbef._ddafc.TblInd.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_cbef._ddafc.TblInd.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(ind / _eb.Dxa))
}

// SetHorizontalBanding controls the conditional formatting for horizontal banding.
func (_bcgde TableLook) SetHorizontalBanding(on bool) {
	if !on {
		_bcgde._ffdgf.NoHBandAttr = &_cg.ST_OnOff{}
		_bcgde._ffdgf.NoHBandAttr.ST_OnOff1 = _cg.ST_OnOff1On
	} else {
		_bcgde._ffdgf.NoHBandAttr = &_cg.ST_OnOff{}
		_bcgde._ffdgf.NoHBandAttr.ST_OnOff1 = _cg.ST_OnOff1Off
	}
}

// GetWrapPathLineTo return wrapPath lineTo value.
func (_agf AnchorDrawWrapOptions) GetWrapPathLineTo() []*_bb.CT_Point2D { return _agf._ga }

// X returns the inner wrapped XML type.
func (_feed Run) X() *_eaba.CT_R { return _feed._bggd }

// AddLevel adds a new numbering level to a NumberingDefinition.
func (_acede NumberingDefinition) AddLevel() NumberingLevel {
	_gedaf := _eaba.NewCT_Lvl()
	_gedaf.Start = &_eaba.CT_DecimalNumber{ValAttr: 1}
	_gedaf.IlvlAttr = int64(len(_acede._fbde.Lvl))
	_acede._fbde.Lvl = append(_acede._fbde.Lvl, _gedaf)
	return NumberingLevel{_gedaf}
}

// SetDoubleStrikeThrough sets the run to double strike-through.
func (_gafea RunProperties) SetDoubleStrikeThrough(b bool) {
	if !b {
		_gafea._gfbgc.Dstrike = nil
	} else {
		_gafea._gfbgc.Dstrike = _eaba.NewCT_OnOff()
	}
}

// CellMargins are the margins for an individual cell.
type CellMargins struct{ _ec *_eaba.CT_TcMar }

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_baggb Header) AddImage(i _ea.Image) (_ea.ImageRef, error) {
	var _dag _ea.Relationships
	for _gbfg, _cgcdf := range _baggb._bffa._bbac {
		if _cgcdf == _baggb._cbgbf {
			_dag = _baggb._bffa._cfg[_gbfg]
		}
	}
	_geff := _ea.MakeImageRef(i, &_baggb._bffa.DocBase, _dag)
	if i.Data == nil && i.Path == "" {
		return _geff, _ffa.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _geff, _ffa.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _geff, _ffa.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	_baggb._bffa.Images = append(_baggb._bffa.Images, _geff)
	_cefc := _eg.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_baggb._bffa.Images), i.Format)
	_effgb := _dag.AddRelationship(_cefc, _d.ImageType)
	_geff.SetRelID(_effgb.X().IdAttr)
	return _geff, nil
}

func _efff(_bcef int64, _daea *_eaba.CT_NumFmt) (_egea string) {
	if _daea == nil {
		return
	}
	_deac := _daea.ValAttr
	switch _deac {
	case _eaba.ST_NumberFormatNone:
		_egea = ""
	case _eaba.ST_NumberFormatDecimal:
		_egea = _dc.Itoa(int(_bcef))
	case _eaba.ST_NumberFormatDecimalZero:
		_egea = _dc.Itoa(int(_bcef))
		if _bcef < 10 {
			_egea = "\u0030" + _egea
		}
	case _eaba.ST_NumberFormatUpperRoman:
		var (
			_cgda  = _bcef % 10
			_eeeea = (_bcef % 100) / 10
			_bacf  = (_bcef % 1000) / 100
			_aded  = _bcef / 1000
		)
		_egea = _fcfa[_aded] + _abaa[_bacf] + _cbaf[_eeeea] + _cggc[_cgda]
	case _eaba.ST_NumberFormatLowerRoman:
		var (
			_dgfcb = _bcef % 10
			_bdae  = (_bcef % 100) / 10
			_bbea  = (_bcef % 1000) / 100
			_befa  = _bcef / 1000
		)
		_egea = _fcfa[_befa] + _abaa[_bbea] + _cbaf[_bdae] + _cggc[_dgfcb]
		_egea = _f.ToLower(_egea)
	case _eaba.ST_NumberFormatUpperLetter:
		_ffae := _bcef % 780
		if _ffae == 0 {
			_ffae = 780
		}
		_dbeb := (_ffae - 1) / 26
		_gffd := (_ffae - 1) % 26
		_bbcf := _dbbc[_dbeb+_gffd]
		_egea = string(_bbcf)
	case _eaba.ST_NumberFormatLowerLetter:
		_caee := _bcef % 780
		if _caee == 0 {
			_caee = 780
		}
		_bcea := (_caee - 1) / 26
		_gbaf := (_caee - 1) % 26
		_eaga := _dbbc[_bcea+_gbaf]
		_egea = _f.ToLower(string(_eaga))
	default:
		_egea = ""
	}
	return
}

// SetName sets the name of the style.
func (_acefe Style) SetName(name string) {
	_acefe._bgfca.Name = _eaba.NewCT_String()
	_acefe._bgfca.Name.ValAttr = name
}

// Clear removes all of the content from within a run.
func (_bebd Run) Clear() { _bebd._bggd.EG_RunInnerContent = nil }

// SetTarget sets the URL target of the hyperlink.
func (_ccae HyperLink) SetTarget(url string) {
	_faafa := _ccae._bcgg.AddHyperlink(url)
	_ccae._dgace.IdAttr = _d.String(_ea.Relationship(_faafa).ID())
	_ccae._dgace.AnchorAttr = nil
}

// Paragraphs returns the paragraphs defined in the cell.
func (_cgc Cell) Paragraphs() []Paragraph {
	_aab := []Paragraph{}
	for _, _fac := range _cgc._acg.EG_BlockLevelElts {
		for _, _baf := range _fac.EG_ContentBlockContent {
			for _, _ccb := range _baf.P {
				_aab = append(_aab, Paragraph{_cgc._fc, _ccb})
			}
		}
	}
	return _aab
}

// AddTabStop adds a tab stop to the paragraph.  It controls the position of text when using Run.AddTab()
func (_fbfce ParagraphProperties) AddTabStop(position _eb.Distance, justificaton _eaba.ST_TabJc, leader _eaba.ST_TabTlc) {
	if _fbfce._efec.Tabs == nil {
		_fbfce._efec.Tabs = _eaba.NewCT_Tabs()
	}
	_ddfa := _eaba.NewCT_TabStop()
	_ddfa.LeaderAttr = leader
	_ddfa.ValAttr = justificaton
	_ddfa.PosAttr.Int64 = _d.Int64(int64(position / _eb.Twips))
	_fbfce._efec.Tabs.Tab = append(_fbfce._efec.Tabs.Tab, _ddfa)
}

func _fffb(_abce *_eaba.CT_P, _abca map[string]string) {
	for _, _ebbg := range _abce.EG_PContent {
		if _ebbg.Hyperlink != nil && _ebbg.Hyperlink.IdAttr != nil {
			if _bcbbc, _bbdbc := _abca[*_ebbg.Hyperlink.IdAttr]; _bbdbc {
				*_ebbg.Hyperlink.IdAttr = _bcbbc
			}
		}
	}
}

// X returns the inner wrapped XML type.
func (_edfea Row) X() *_eaba.CT_Row { return _edfea._cegac }

func (_cbeb *Document) insertStyleFromNode(_gfab Node) {
	if _gfab.Style.X() != nil {
		if _, _ebgd := _cbeb.Styles.SearchStyleById(_gfab.Style.StyleID()); !_ebgd {
			_cbeb.Styles.InsertStyle(_gfab.Style)
			_ebbe := _gfab.Style.ParagraphProperties()
			_cbeb.insertNumberingFromStyleProperties(_gfab._abcbe.Numbering, _ebbe)
		}
	}
}

// Value returns the tring value of a FormFieldTypeText or FormFieldTypeDropDown.
func (_gffbe FormField) Value() string {
	if _gffbe._cdag.TextInput != nil && _gffbe._bebe.T != nil {
		return _gffbe._bebe.T.Content
	} else if _gffbe._cdag.DdList != nil && _gffbe._cdag.DdList.Result != nil {
		_cdgb := _gffbe.PossibleValues()
		_dgfc := int(_gffbe._cdag.DdList.Result.ValAttr)
		if _dgfc < len(_cdgb) {
			return _cdgb[_dgfc]
		}
	} else if _gffbe._cdag.CheckBox != nil {
		if _gffbe.IsChecked() {
			return "\u0074\u0072\u0075\u0065"
		}
		return "\u0066\u0061\u006cs\u0065"
	}
	return ""
}

// AddParagraph adds a paragraph to the footnote.
func (_fegc Footnote) AddParagraph() Paragraph {
	_bfacb := _eaba.NewEG_ContentBlockContent()
	_daf := len(_fegc._efba.EG_BlockLevelElts[0].EG_ContentBlockContent)
	_fegc._efba.EG_BlockLevelElts[0].EG_ContentBlockContent = append(_fegc._efba.EG_BlockLevelElts[0].EG_ContentBlockContent, _bfacb)
	_gbadf := _eaba.NewCT_P()
	var _bbfa *_eaba.CT_String
	if _daf != 0 {
		_bbeb := len(_fegc._efba.EG_BlockLevelElts[0].EG_ContentBlockContent[_daf-1].P)
		_bbfa = _fegc._efba.EG_BlockLevelElts[0].EG_ContentBlockContent[_daf-1].P[_bbeb-1].PPr.PStyle
	} else {
		_bbfa = _eaba.NewCT_String()
		_bbfa.ValAttr = "\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065"
	}
	_bfacb.P = append(_bfacb.P, _gbadf)
	_gfeg := Paragraph{_fegc._ffdf, _gbadf}
	_gfeg._dbef.PPr = _eaba.NewCT_PPr()
	_gfeg._dbef.PPr.PStyle = _bbfa
	_gfeg._dbef.PPr.RPr = _eaba.NewCT_ParaRPr()
	return _gfeg
}

// GetTargetByRelId returns a target path with the associated relation ID in the
// document.
func (_fged *Document) GetTargetByRelId(idAttr string) string {
	return _fged._bdb.GetTargetByRelId(idAttr)
}

// SetEffect sets a text effect on the run.
func (_becgc RunProperties) SetEffect(e _eaba.ST_TextEffect) {
	if e == _eaba.ST_TextEffectUnset {
		_becgc._gfbgc.Effect = nil
	} else {
		_becgc._gfbgc.Effect = _eaba.NewCT_TextEffect()
		_becgc._gfbgc.Effect.ValAttr = _eaba.ST_TextEffectShimmer
	}
}

// RemoveParagraph removes a paragraph from a document.
func (_dfca *Document) RemoveParagraph(p Paragraph) {
	if _dfca._eeda.Body == nil {
		return
	}
	for _, _bgc := range _dfca._eeda.Body.EG_BlockLevelElts {
		for _, _bfb := range _bgc.EG_ContentBlockContent {
			for _bcaf, _dcbf := range _bfb.P {
				if _dcbf == p._dbef {
					copy(_bfb.P[_bcaf:], _bfb.P[_bcaf+1:])
					_bfb.P = _bfb.P[0 : len(_bfb.P)-1]
					return
				}
			}
			if _bfb.Sdt != nil && _bfb.Sdt.SdtContent != nil && _bfb.Sdt.SdtContent.P != nil {
				for _dgbb, _cdefe := range _bfb.Sdt.SdtContent.P {
					if _cdefe == p._dbef {
						copy(_bfb.P[_dgbb:], _bfb.P[_dgbb+1:])
						_bfb.P = _bfb.P[0 : len(_bfb.P)-1]
						return
					}
				}
			}
		}
	}
	for _, _gbef := range _dfca.Tables() {
		for _, _aaad := range _gbef.Rows() {
			for _, _egf := range _aaad.Cells() {
				for _, _cdf := range _egf._acg.EG_BlockLevelElts {
					for _, _fba := range _cdf.EG_ContentBlockContent {
						for _ddaa, _dea := range _fba.P {
							if _dea == p._dbef {
								copy(_fba.P[_ddaa:], _fba.P[_ddaa+1:])
								_fba.P = _fba.P[0 : len(_fba.P)-1]
								return
							}
						}
					}
				}
			}
		}
	}
	for _, _cca := range _dfca.Headers() {
		_cca.RemoveParagraph(p)
	}
	for _, _fag := range _dfca.Footers() {
		_fag.RemoveParagraph(p)
	}
}

func (_cceb *WatermarkText) getShape() *_d.XSDAny {
	return _cceb.getInnerElement("\u0073\u0068\u0061p\u0065")
}

// Properties returns the run properties.
func (_beae Run) Properties() RunProperties {
	if _beae._bggd.RPr == nil {
		_beae._bggd.RPr = _eaba.NewCT_RPr()
	}
	return RunProperties{_beae._bggd.RPr}
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (_ddbcdc ParagraphSpacing) SetLineSpacing(d _eb.Distance, rule _eaba.ST_LineSpacingRule) {
	if rule == _eaba.ST_LineSpacingRuleUnset {
		_ddbcdc._ebfge.LineRuleAttr = _eaba.ST_LineSpacingRuleUnset
		_ddbcdc._ebfge.LineAttr = nil
	} else {
		_ddbcdc._ebfge.LineRuleAttr = rule
		_ddbcdc._ebfge.LineAttr = &_eaba.ST_SignedTwipsMeasure{}
		_ddbcdc._ebfge.LineAttr.Int64 = _d.Int64(int64(d / _eb.Twips))
	}
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (_bgdgf Paragraph) SetFirstLineIndent(m _eb.Distance) {
	_bgdgf.ensurePPr()
	_bcfab := _bgdgf._dbef.PPr
	if _bcfab.Ind == nil {
		_bcfab.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_bcfab.Ind.FirstLineAttr = nil
	} else {
		_bcfab.Ind.FirstLineAttr = &_cg.ST_TwipsMeasure{}
		_bcfab.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(m / _eb.Twips))
	}
}

// CharacterSpacingValue returns the value of characters spacing in twips (1/20 of point).
func (_bffec ParagraphProperties) CharacterSpacingValue() int64 {
	if _gdee := _bffec._efec.RPr.Spacing; _gdee != nil {
		_befaa := _gdee.ValAttr
		if _befaa.Int64 != nil {
			return *_befaa.Int64
		}
	}
	return int64(0)
}

func (_deed *Document) insertImageFromNode(_gcdd Node) {
	for _, _fcggd := range _gcdd.AnchoredDrawings {
		if _eggc, _dccd := _fcggd.GetImage(); _dccd {
			_bcbba, _cebed := _ea.ImageFromFile(_eggc.Path())
			if _cebed != nil {
				_dcd.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _cebed)
			}
			_dace, _cebed := _deed.AddImage(_bcbba)
			if _cebed != nil {
				_dcd.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _cebed)
			}
			_fbebd := _deed._bdb.GetByRelId(_dace.RelID())
			_fbebd.SetID(_eggc.RelID())
		}
	}
	for _, _edbg := range _gcdd.InlineDrawings {
		if _aabd, _eege := _edbg.GetImage(); _eege {
			_dbgaa, _deec := _ea.ImageFromFile(_aabd.Path())
			if _deec != nil {
				_dcd.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _deec)
			}
			_gbad, _deec := _deed.AddImage(_dbgaa)
			if _deec != nil {
				_dcd.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _deec)
			}
			_fcag := _deed._bdb.GetByRelId(_gbad.RelID())
			_fcag.SetID(_aabd.RelID())
		}
	}
}

// SetTextWrapTopAndBottom sets the text wrap to top and bottom.
func (_eaf AnchoredDrawing) SetTextWrapTopAndBottom() {
	_eaf._cc.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_eaf._cc.Choice.WrapTopAndBottom = _eaba.NewWdCT_WrapTopBottom()
	_eaf._cc.LayoutInCellAttr = true
	_eaf._cc.AllowOverlapAttr = true
}

// SetWindowControl controls if the first or last line of the paragraph is
// allowed to dispay on a separate page.
func (_dabdd ParagraphProperties) SetWindowControl(b bool) {
	if !b {
		_dabdd._efec.WidowControl = nil
	} else {
		_dabdd._efec.WidowControl = _eaba.NewCT_OnOff()
	}
}

// ReplaceText replace the text inside node.
func (_gege *Node) ReplaceText(oldText, newText string) {
	switch _agged := _gege.X().(type) {
	case *Paragraph:
		for _, _eacb := range _agged.Runs() {
			for _, _bdfgc := range _eacb._bggd.EG_RunInnerContent {
				if _bdfgc.T != nil {
					_ddee := _bdfgc.T.Content
					_ddee = _f.ReplaceAll(_ddee, oldText, newText)
					_bdfgc.T.Content = _ddee
				}
			}
		}
	}
	for _, _fafad := range _gege.Children {
		_fafad.ReplaceText(oldText, newText)
	}
}

// SetTop sets the top border to a specified type, color and thickness.
func (_dcdg TableBorders) SetTop(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_dcdg._dbbcbf.Top = _eaba.NewCT_Border()
	_gfdf(_dcdg._dbbcbf.Top, t, c, thickness)
}

// X returns the inner wrapped XML type.
func (_dff Cell) X() *_eaba.CT_Tc { return _dff._acg }

// X returns the inner wrapped XML type.
func (_babdc Numbering) X() *_eaba.Numbering { return _babdc._ebbdb }

// Header is a header for a document section.
type Header struct {
	_bffa  *Document
	_cbgbf *_eaba.Hdr
}

// AddSection adds a new document section with an optional section break.  If t
// is ST_SectionMarkUnset, then no break will be inserted.
func (_eagfd ParagraphProperties) AddSection(t _eaba.ST_SectionMark) Section {
	_eagfd._efec.SectPr = _eaba.NewCT_SectPr()
	if t != _eaba.ST_SectionMarkUnset {
		_eagfd._efec.SectPr.Type = _eaba.NewCT_SectType()
		_eagfd._efec.SectPr.Type.ValAttr = t
	}
	return Section{_eagfd._gaeba, _eagfd._efec.SectPr}
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_adgfa TableProperties) SetCellSpacingAuto() {
	_adgfa._efac.TblCellSpacing = _eaba.NewCT_TblWidth()
	_adgfa._efac.TblCellSpacing.TypeAttr = _eaba.ST_TblWidthAuto
}

// SetSmallCaps sets the run to small caps.
func (_bebb RunProperties) SetSmallCaps(b bool) {
	if !b {
		_bebb._gfbgc.SmallCaps = nil
	} else {
		_bebb._gfbgc.SmallCaps = _eaba.NewCT_OnOff()
	}
}

// AddRun adds a run of text to a hyperlink. This is the text that will be linked.
func (_dffg HyperLink) AddRun() Run {
	_dbcf := _eaba.NewEG_ContentRunContent()
	_dffg._dgace.EG_ContentRunContent = append(_dffg._dgace.EG_ContentRunContent, _dbcf)
	_degg := _eaba.NewCT_R()
	_dbcf.R = _degg
	return Run{_dffg._bcgg, _degg}
}

// SizeValue returns the value of paragraph font size in points.
func (_ddbcd ParagraphProperties) SizeValue() float64 {
	if _adga := _ddbcd._efec.RPr.Sz; _adga != nil {
		_abfba := _adga.ValAttr
		if _abfba.ST_UnsignedDecimalNumber != nil {
			return float64(*_abfba.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

func (_fbea Paragraph) addEndBookmark(_gfdaa int64) *_eaba.CT_MarkupRange {
	_agdd := _eaba.NewEG_PContent()
	_fbea._dbef.EG_PContent = append(_fbea._dbef.EG_PContent, _agdd)
	_fegg := _eaba.NewEG_ContentRunContent()
	_gfdagg := _eaba.NewEG_RunLevelElts()
	_efeef := _eaba.NewEG_RangeMarkupElements()
	_ebdac := _eaba.NewCT_MarkupRange()
	_ebdac.IdAttr = _gfdaa
	_efeef.BookmarkEnd = _ebdac
	_agdd.EG_ContentRunContent = append(_agdd.EG_ContentRunContent, _fegg)
	_fegg.EG_RunLevelElts = append(_fegg.EG_RunLevelElts, _gfdagg)
	_gfdagg.EG_RangeMarkupElements = append(_gfdagg.EG_RangeMarkupElements, _efeef)
	return _ebdac
}

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_gdcg ParagraphProperties) SizeMeasure() string {
	if _agbae := _gdcg._efec.RPr.Sz; _agbae != nil {
		_gbca := _agbae.ValAttr
		if _gbca.ST_PositiveUniversalMeasure != nil {
			return *_gbca.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// X returns the inner wrapped type
func (_ffb CellBorders) X() *_eaba.CT_TcBorders { return _ffb._ggf }

func (_dbfc *Document) insertNumberingFromStyleProperties(_eccfg Numbering, _cbgb ParagraphStyleProperties) {
	_cbbd := _cbgb.NumId()
	_ggb := int64(-1)
	if _cbbd > -1 {
		for _, _gabea := range _eccfg._ebbdb.Num {
			if _gabea.NumIdAttr == _cbbd {
				if _gabea.AbstractNumId != nil {
					_ggb = _gabea.AbstractNumId.ValAttr
					_fbbc := false
					for _, _dfcag := range _dbfc.Numbering._ebbdb.Num {
						if _dfcag.NumIdAttr == _cbbd {
							_fbbc = true
							break
						}
					}
					if !_fbbc {
						_dbfc.Numbering._ebbdb.Num = append(_dbfc.Numbering._ebbdb.Num, _gabea)
					}
					break
				}
			}
		}
		for _, _fbbbd := range _eccfg._ebbdb.AbstractNum {
			if _fbbbd.AbstractNumIdAttr == _ggb {
				_gffgb := false
				for _, _bcaff := range _dbfc.Numbering._ebbdb.AbstractNum {
					if _bcaff.AbstractNumIdAttr == _ggb {
						_gffgb = true
						break
					}
				}
				if !_gffgb {
					_dbfc.Numbering._ebbdb.AbstractNum = append(_dbfc.Numbering._ebbdb.AbstractNum, _fbbbd)
				}
				break
			}
		}
	}
}

type listItemInfo struct {
	FromStyle      *Style
	FromParagraph  *Paragraph
	AbstractNumId  *int64
	NumberingLevel *NumberingLevel
}

func _gecg(_ccfb *_eaba.CT_P, _ddbad *_eaba.CT_Hyperlink, _fcfd *TableInfo, _ecaf *DrawingInfo, _fdde []*_eaba.EG_ContentRunContent) []TextItem {
	_febc := []TextItem{}
	for _, _geaea := range _fdde {
		if _ddcef := _geaea.R; _ddcef != nil {
			_adee := _cb.NewBuffer([]byte{})
			for _, _feee := range _ddcef.EG_RunInnerContent {
				if _feee.T != nil && _feee.T.Content != "" {
					_adee.WriteString(_feee.T.Content)
				}
			}
			_febc = append(_febc, TextItem{Text: _adee.String(), DrawingInfo: _ecaf, Paragraph: _ccfb, Hyperlink: _ddbad, Run: _ddcef, TableInfo: _fcfd})
			for _, _fgcag := range _ddcef.Extra {
				if _baaaf, _gcaa := _fgcag.(*_eaba.AlternateContentRun); _gcaa {
					_abec := &DrawingInfo{Drawing: _baaaf.Choice.Drawing}
					for _, _fcgcg := range _abec.Drawing.Anchor {
						for _, _gbgg := range _fcgcg.Graphic.GraphicData.Any {
							if _bced, _dcdf := _gbgg.(*_eaba.WdWsp); _dcdf {
								if _bced.WChoice != nil {
									if _fggd := _bced.SpPr; _fggd != nil {
										if _adaba := _fggd.Xfrm; _adaba != nil {
											if _efad := _adaba.Ext; _efad != nil {
												_abec.Width = _efad.CxAttr
												_abec.Height = _efad.CyAttr
											}
										}
									}
									for _, _cafe := range _bced.WChoice.Txbx.TxbxContent.EG_ContentBlockContent {
										_febc = append(_febc, _bbbef(_cafe.P, _fcfd, _abec)...)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return _febc
}

// Levels returns all of the numbering levels defined in the definition.
func (_dceb NumberingDefinition) Levels() []NumberingLevel {
	_afeef := []NumberingLevel{}
	for _, _edda := range _dceb._fbde.Lvl {
		_afeef = append(_afeef, NumberingLevel{_edda})
	}
	return _afeef
}

// X returns the inner wrapped XML type.
func (_dbf Color) X() *_eaba.CT_Color { return _dbf._bcf }

// Shadow returns true if run shadow is on.
func (_cede RunProperties) Shadow() bool { return _beff(_cede._gfbgc.Shadow) }

// SetLinkedStyle sets the style that this style is linked to.
func (_daee Style) SetLinkedStyle(name string) {
	if name == "" {
		_daee._bgfca.Link = nil
	} else {
		_daee._bgfca.Link = _eaba.NewCT_String()
		_daee._bgfca.Link.ValAttr = name
	}
}

// SetSpacing sets the spacing that comes before and after the paragraph.
// Deprecated: See Spacing() instead which allows finer control.
func (_ffaa ParagraphProperties) SetSpacing(before, after _eb.Distance) {
	if _ffaa._efec.Spacing == nil {
		_ffaa._efec.Spacing = _eaba.NewCT_Spacing()
	}
	_ffaa._efec.Spacing.BeforeAttr = &_cg.ST_TwipsMeasure{}
	_ffaa._efec.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(before / _eb.Twips))
	_ffaa._efec.Spacing.AfterAttr = &_cg.ST_TwipsMeasure{}
	_ffaa._efec.Spacing.AfterAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(after / _eb.Twips))
}

// AddHyperlink adds a hyperlink to a document. Adding the hyperlink to a document
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (_cdgc Document) AddHyperlink(url string) _ea.Hyperlink {
	return _cdgc._bdb.AddHyperlink(url)
}

// SetWrapPathStart sets wrapPath start value.
func (_gbf AnchorDrawWrapOptions) SetWrapPathStart(coordinate *_bb.CT_Point2D) {
	_gbf._bba = coordinate
}

// Cell is a table cell within a document (not a spreadsheet)
type Cell struct {
	_fc  *Document
	_acg *_eaba.CT_Tc
}

// SetAfterAuto controls if spacing after a paragraph is automatically determined.
func (_feaa ParagraphSpacing) SetAfterAuto(b bool) {
	if b {
		_feaa._ebfge.AfterAutospacingAttr = &_cg.ST_OnOff{}
		_feaa._ebfge.AfterAutospacingAttr.Bool = _d.Bool(true)
	} else {
		_feaa._ebfge.AfterAutospacingAttr = nil
	}
}

// TableInfo is used for keep information about a table, a row and a cell where the text is located.
type TableInfo struct {
	Table    *_eaba.CT_Tbl
	Row      *_eaba.CT_Row
	Cell     *_eaba.CT_Tc
	RowIndex int
	ColIndex int
}

// SetFormat sets the numbering format.
func (_fdgg NumberingLevel) SetFormat(f _eaba.ST_NumberFormat) {
	if _fdgg._fgbfa.NumFmt == nil {
		_fdgg._fgbfa.NumFmt = _eaba.NewCT_NumFmt()
	}
	_fdgg._fgbfa.NumFmt.ValAttr = f
}

// SetAfterLineSpacing sets spacing below paragraph in line units.
func (_ccfe Paragraph) SetAfterLineSpacing(d _eb.Distance) {
	_ccfe.ensurePPr()
	if _ccfe._dbef.PPr.Spacing == nil {
		_ccfe._dbef.PPr.Spacing = _eaba.NewCT_Spacing()
	}
	_gbda := _ccfe._dbef.PPr.Spacing
	_gbda.AfterLinesAttr = _d.Int64(int64(d / _eb.Twips))
}

// Nodes return the document's element as nodes.
func (_dgfa *Document) Nodes() Nodes {
	_adbe := []Node{}
	for _, _fcc := range _dgfa._eeda.Body.EG_BlockLevelElts {
		_adbe = append(_adbe, _dabf(_dgfa, _fcc.EG_ContentBlockContent, nil)...)
	}
	if _dgfa._eeda.Body.SectPr != nil {
		_adbe = append(_adbe, Node{_efag: _dgfa._eeda.Body.SectPr})
	}
	_gaac := Nodes{_agcab: _adbe}
	return _gaac
}

func (_fabcg *WatermarkPicture) getShapeType() *_d.XSDAny {
	return _fabcg.getInnerElement("\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e")
}

// GetImage returns the ImageRef associated with an InlineDrawing.
func (_dcbga InlineDrawing) GetImage() (_ea.ImageRef, bool) {
	_bffe := _dcbga._fbbd.Graphic.GraphicData.Any
	if len(_bffe) > 0 {
		_gbag, _cdeag := _bffe[0].(*_gd.Pic)
		if _cdeag {
			if _gbag.BlipFill != nil && _gbag.BlipFill.Blip != nil && _gbag.BlipFill.Blip.EmbedAttr != nil {
				return _dcbga._cgdg.GetImageByRelID(*_gbag.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return _ea.ImageRef{}, false
}

// SetStrict is a shortcut for document.SetConformance,
// as one of these values from gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (_efg Document) SetStrict(strict bool) {
	if strict {
		_efg._eeda.ConformanceAttr = _cg.ST_ConformanceClassStrict
	} else {
		_efg._eeda.ConformanceAttr = _cg.ST_ConformanceClassTransitional
	}
}

// GetColor returns the color.Color object representing the run color.
func (_dcgfb RunProperties) GetColor() _bc.Color {
	if _ebbb := _dcgfb._gfbgc.Color; _ebbb != nil {
		_cgbd := _ebbb.ValAttr
		if _cgbd.ST_HexColorRGB != nil {
			return _bc.FromHex(*_cgbd.ST_HexColorRGB)
		}
	}
	return _bc.Color{}
}

// SetVerticalBanding controls the conditional formatting for vertical banding.
func (_aege TableLook) SetVerticalBanding(on bool) {
	if !on {
		_aege._ffdgf.NoVBandAttr = &_cg.ST_OnOff{}
		_aege._ffdgf.NoVBandAttr.ST_OnOff1 = _cg.ST_OnOff1On
	} else {
		_aege._ffdgf.NoVBandAttr = &_cg.ST_OnOff{}
		_aege._ffdgf.NoVBandAttr.ST_OnOff1 = _cg.ST_OnOff1Off
	}
}

// SetTextWrapTight sets the text wrap to tight with a give wrap type.
func (_aef AnchoredDrawing) SetTextWrapTight(option *AnchorDrawWrapOptions) {
	_aef._cc.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_aef._cc.Choice.WrapTight = _eaba.NewWdCT_WrapTight()
	_aef._cc.Choice.WrapTight.WrapTextAttr = _eaba.WdST_WrapTextBothSides
	_daa := false
	_aef._cc.Choice.WrapTight.WrapPolygon.EditedAttr = &_daa
	if option == nil {
		option = NewAnchorDrawWrapOptions()
	}
	_aef._cc.Choice.WrapTight.WrapPolygon.LineTo = option.GetWrapPathLineTo()
	_aef._cc.Choice.WrapTight.WrapPolygon.Start = option.GetWrapPathStart()
	_aef._cc.LayoutInCellAttr = true
	_aef._cc.AllowOverlapAttr = true
}

// SetRightIndent controls right indent of paragraph.
func (_gfed Paragraph) SetRightIndent(m _eb.Distance) {
	_gfed.ensurePPr()
	_ddfbg := _gfed._dbef.PPr
	if _ddfbg.Ind == nil {
		_ddfbg.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_ddfbg.Ind.RightAttr = nil
	} else {
		_ddfbg.Ind.RightAttr = &_eaba.ST_SignedTwipsMeasure{}
		_ddfbg.Ind.RightAttr.Int64 = _d.Int64(int64(m / _eb.Twips))
	}
}

// SetPageBreakBefore controls if there is a page break before this paragraph.
func (_dacc ParagraphProperties) SetPageBreakBefore(b bool) {
	if !b {
		_dacc._efec.PageBreakBefore = nil
	} else {
		_dacc._efec.PageBreakBefore = _eaba.NewCT_OnOff()
	}
}

// SetNumberingDefinition sets the numbering definition ID via a NumberingDefinition
// defined in numbering.xml
func (_gefad Paragraph) SetNumberingDefinition(nd NumberingDefinition) {
	_gefad.ensurePPr()
	if _gefad._dbef.PPr.NumPr == nil {
		_gefad._dbef.PPr.NumPr = _eaba.NewCT_NumPr()
	}
	_dbca := _eaba.NewCT_DecimalNumber()
	_afef := int64(-1)
	for _, _ccadb := range _gefad._bbgc.Numbering._ebbdb.Num {
		if _ccadb.AbstractNumId != nil && _ccadb.AbstractNumId.ValAttr == nd.AbstractNumberID() {
			_afef = _ccadb.NumIdAttr
		}
	}
	if _afef == -1 {
		_cffa := _eaba.NewCT_Num()
		_gefad._bbgc.Numbering._ebbdb.Num = append(_gefad._bbgc.Numbering._ebbdb.Num, _cffa)
		_cffa.NumIdAttr = int64(len(_gefad._bbgc.Numbering._ebbdb.Num))
		_cffa.AbstractNumId = _eaba.NewCT_DecimalNumber()
		_cffa.AbstractNumId.ValAttr = nd.AbstractNumberID()
	}
	_dbca.ValAttr = _afef
	_gefad._dbef.PPr.NumPr.NumId = _dbca
}

func (_dae *Document) insertTable(_aag Paragraph, _gfea bool) Table {
	_eabb := _dae._eeda.Body
	if _eabb == nil {
		return _dae.AddTable()
	}
	_fdf := _aag.X()
	for _dde, _dfc := range _eabb.EG_BlockLevelElts {
		for _, _cagd := range _dfc.EG_ContentBlockContent {
			for _edfe, _bef := range _cagd.P {
				if _bef == _fdf {
					_cbff := _eaba.NewCT_Tbl()
					_eeb := _eaba.NewEG_BlockLevelElts()
					_gfce := _eaba.NewEG_ContentBlockContent()
					_eeb.EG_ContentBlockContent = append(_eeb.EG_ContentBlockContent, _gfce)
					_gfce.Tbl = append(_gfce.Tbl, _cbff)
					_eabb.EG_BlockLevelElts = append(_eabb.EG_BlockLevelElts, nil)
					if _gfea {
						copy(_eabb.EG_BlockLevelElts[_dde+1:], _eabb.EG_BlockLevelElts[_dde:])
						_eabb.EG_BlockLevelElts[_dde] = _eeb
						if _edfe != 0 {
							_bff := _eaba.NewEG_BlockLevelElts()
							_gbfe := _eaba.NewEG_ContentBlockContent()
							_bff.EG_ContentBlockContent = append(_bff.EG_ContentBlockContent, _gbfe)
							_gbfe.P = _cagd.P[:_edfe]
							_eabb.EG_BlockLevelElts = append(_eabb.EG_BlockLevelElts, nil)
							copy(_eabb.EG_BlockLevelElts[_dde+1:], _eabb.EG_BlockLevelElts[_dde:])
							_eabb.EG_BlockLevelElts[_dde] = _bff
						}
						_cagd.P = _cagd.P[_edfe:]
					} else {
						copy(_eabb.EG_BlockLevelElts[_dde+2:], _eabb.EG_BlockLevelElts[_dde+1:])
						_eabb.EG_BlockLevelElts[_dde+1] = _eeb
						if _edfe != len(_cagd.P)-1 {
							_eaff := _eaba.NewEG_BlockLevelElts()
							_fgcb := _eaba.NewEG_ContentBlockContent()
							_eaff.EG_ContentBlockContent = append(_eaff.EG_ContentBlockContent, _fgcb)
							_fgcb.P = _cagd.P[_edfe+1:]
							_eabb.EG_BlockLevelElts = append(_eabb.EG_BlockLevelElts, nil)
							copy(_eabb.EG_BlockLevelElts[_dde+3:], _eabb.EG_BlockLevelElts[_dde+2:])
							_eabb.EG_BlockLevelElts[_dde+2] = _eaff
						}
						_cagd.P = _cagd.P[:_edfe+1]
					}
					return Table{_dae, _cbff}
				}
			}
			for _, _ccbe := range _cagd.Tbl {
				_efeg := _fgg(_ccbe, _fdf, _gfea)
				if _efeg != nil {
					return Table{_dae, _efeg}
				}
			}
		}
	}
	return _dae.AddTable()
}

// SetText sets the text to be used in bullet mode.
func (_dfafe NumberingLevel) SetText(t string) {
	if t == "" {
		_dfafe._fgbfa.LvlText = nil
	} else {
		_dfafe._fgbfa.LvlText = _eaba.NewCT_LevelText()
		_dfafe._fgbfa.LvlText.ValAttr = _d.String(t)
	}
}

// SetXOffset sets the X offset for an image relative to the origin.
func (_dab AnchoredDrawing) SetXOffset(x _eb.Distance) {
	_dab._cc.PositionH.Choice = &_eaba.WdCT_PosHChoice{}
	_dab._cc.PositionH.Choice.PosOffset = _d.Int32(int32(x / _eb.EMU))
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (_ddfab TableBorders) SetBottom(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_ddfab._dbbcbf.Bottom = _eaba.NewCT_Border()
	_gfdf(_ddfab._dbbcbf.Bottom, t, c, thickness)
}

// CellBorders are the borders for an individual
type CellBorders struct{ _ggf *_eaba.CT_TcBorders }

func (_fgac *Document) getWatermarkHeaderInnerContentPictures() []*_eaba.CT_Picture {
	var _eeba []*_eaba.CT_Picture
	for _, _ebg := range _fgac.Headers() {
		for _, _gdde := range _ebg.X().EG_ContentBlockContent {
			for _, _ebfc := range _gdde.P {
				for _, _fdfb := range _ebfc.EG_PContent {
					for _, _agec := range _fdfb.EG_ContentRunContent {
						if _agec.R == nil {
							continue
						}
						for _, _eegc := range _agec.R.EG_RunInnerContent {
							if _eegc.Pict == nil {
								continue
							}
							_ecef := false
							for _, _baac := range _eegc.Pict.Any {
								_bfaeb, _cfgf := _baac.(*_d.XSDAny)
								if _cfgf && _bfaeb.XMLName.Local == "\u0073\u0068\u0061p\u0065" {
									_ecef = true
								}
							}
							if _ecef {
								_eeba = append(_eeba, _eegc.Pict)
							}
						}
					}
				}
			}
		}
	}
	return _eeba
}

func _dbbgc() *_aae.Textpath {
	_edgab := _aae.NewTextpath()
	_fdeb := "\u0066\u006f\u006e\u0074\u002d\u0066\u0061\u006d\u0069l\u0079\u003a\u0022\u0043\u0061\u006c\u0069\u0062\u0072\u0069\u0022\u003b\u0066\u006f\u006e\u0074\u002d\u0073\u0069\u007a\u0065\u003a\u00366\u0070\u0074;\u0066\u006fn\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074\u003a\u0062\u006f\u006c\u0064;f\u006f\u006e\u0074\u002d\u0073\u0074\u0079\u006c\u0065:\u0069\u0074\u0061\u006c\u0069\u0063"
	_edgab.StyleAttr = &_fdeb
	_cfcgb := "\u0041\u0053\u0041\u0050"
	_edgab.StringAttr = &_cfcgb
	return _edgab
}

// AbstractNumberID returns the ID that is unique within all numbering
// definitions that is used to assign the definition to a paragraph.
func (_aagf NumberingDefinition) AbstractNumberID() int64 { return _aagf._fbde.AbstractNumIdAttr }

// X returns the inner wrapped XML type.
func (_aaed CellProperties) X() *_eaba.CT_TcPr { return _aaed._be }

func (_gcccf Footnote) content() []*_eaba.EG_ContentBlockContent {
	var _adgc []*_eaba.EG_ContentBlockContent
	for _, _abga := range _gcccf._efba.EG_BlockLevelElts {
		_adgc = append(_adgc, _abga.EG_ContentBlockContent...)
	}
	return _adgc
}

// SetEastAsiaTheme sets the font East Asia Theme.
func (_acab Fonts) SetEastAsiaTheme(t _eaba.ST_Theme) { _acab._aggeg.EastAsiaThemeAttr = t }

// SetMultiLevelType sets the multilevel type.
func (_egeed NumberingDefinition) SetMultiLevelType(t _eaba.ST_MultiLevelType) {
	if t == _eaba.ST_MultiLevelTypeUnset {
		_egeed._fbde.MultiLevelType = nil
	} else {
		_egeed._fbde.MultiLevelType = _eaba.NewCT_MultiLevelType()
		_egeed._fbde.MultiLevelType.ValAttr = t
	}
}

// SizeValue returns the value of run font size in points.
func (_abdb RunProperties) SizeValue() float64 {
	if _aggcd := _abdb._gfbgc.Sz; _aggcd != nil {
		_ccbf := _aggcd.ValAttr
		if _ccbf.ST_UnsignedDecimalNumber != nil {
			return float64(*_ccbf.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	_gaeba *Document
	_efec  *_eaba.CT_PPr
}

// NumberingLevel is the definition for numbering for a particular level within
// a NumberingDefinition.
type NumberingLevel struct{ _fgbfa *_eaba.CT_Lvl }

// Clear content of node element.
func (_ecggg *Node) Clear() { _ecggg._efag = nil }

// GetStyle returns string style of the text in watermark and format it to TextpathStyle.
func (_acbc *WatermarkText) GetStyle() _ffc.TextpathStyle {
	_gbff := _acbc.getShape()
	if _acbc._ggfda != nil {
		_ffcc := _acbc._ggfda.EG_ShapeElements
		if len(_ffcc) > 0 && _ffcc[0].Textpath != nil {
			return _ffc.NewTextpathStyle(*_ffcc[0].Textpath.StyleAttr)
		}
	} else {
		_cdad := _acbc.findNode(_gbff, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
		for _, _abbdc := range _cdad.Attrs {
			if _abbdc.Name.Local == "\u0073\u0074\u0079l\u0065" {
				return _ffc.NewTextpathStyle(_abbdc.Value)
			}
		}
	}
	return _ffc.NewTextpathStyle("")
}

// SetColumnBandSize sets the number of Columns in the column band
func (_adaff TableStyleProperties) SetColumnBandSize(cols int64) {
	_adaff._ddafc.TblStyleColBandSize = _eaba.NewCT_DecimalNumber()
	_adaff._ddafc.TblStyleColBandSize.ValAttr = cols
}

// FormFields extracts all of the fields from a document.  They can then be
// manipulated via the methods on the field and the document saved.
func (_eecg *Document) FormFields() []FormField {
	_cgbb := []FormField{}
	for _, _aaab := range _eecg.Paragraphs() {
		_egda := _aaab.Runs()
		for _fgd, _eabf := range _egda {
			for _, _gfcf := range _eabf._bggd.EG_RunInnerContent {
				if _gfcf.FldChar == nil || _gfcf.FldChar.FfData == nil {
					continue
				}
				if _gfcf.FldChar.FldCharTypeAttr == _eaba.ST_FldCharTypeBegin {
					if len(_gfcf.FldChar.FfData.Name) == 0 || _gfcf.FldChar.FfData.Name[0].ValAttr == nil {
						continue
					}
					_add := FormField{_cdag: _gfcf.FldChar.FfData}
					if _gfcf.FldChar.FfData.TextInput != nil {
						for _eadc := _fgd + 1; _eadc < len(_egda)-1; _eadc++ {
							if len(_egda[_eadc]._bggd.EG_RunInnerContent) == 0 {
								continue
							}
							_dgdg := _egda[_eadc]._bggd.EG_RunInnerContent[0]
							if _dgdg.FldChar != nil && _dgdg.FldChar.FldCharTypeAttr == _eaba.ST_FldCharTypeSeparate {
								if len(_egda[_eadc+1]._bggd.EG_RunInnerContent) == 0 {
									continue
								}
								if _egda[_eadc+1]._bggd.EG_RunInnerContent[0].FldChar == nil {
									_add._bebe = _egda[_eadc+1]._bggd.EG_RunInnerContent[0]
									break
								}
							}
						}
					}
					_cgbb = append(_cgbb, _add)
				}
			}
		}
	}
	return _cgbb
}

// RemoveEndnote removes a endnote from both the paragraph and the document
// the requested endnote must be anchored on the paragraph being referenced.
func (_eaebb Paragraph) RemoveEndnote(id int64) {
	_ffff := _eaebb._bbgc._bfd
	var _eacf int
	for _dgfb, _dafd := range _ffff.CT_Endnotes.Endnote {
		if _dafd.IdAttr == id {
			_eacf = _dgfb
		}
	}
	_eacf = 0
	_ffff.CT_Endnotes.Endnote[_eacf] = nil
	_ffff.CT_Endnotes.Endnote[_eacf] = _ffff.CT_Endnotes.Endnote[len(_ffff.CT_Endnotes.Endnote)-1]
	_ffff.CT_Endnotes.Endnote = _ffff.CT_Endnotes.Endnote[:len(_ffff.CT_Endnotes.Endnote)-1]
	var _eagdc Run
	for _, _gbcb := range _eaebb.Runs() {
		if _geacg, _dfcad := _gbcb.IsEndnote(); _geacg {
			if _dfcad == id {
				_eagdc = _gbcb
			}
		}
	}
	_eaebb.RemoveRun(_eagdc)
}

// OpenTemplate opens a document, removing all content so it can be used as a
// template.  Since Word removes unused styles from a document upon save, to
// create a template in Word add a paragraph with every style of interest.  When
// opened with OpenTemplate the document's styles will be available but the
// content will be gone.
func OpenTemplate(filename string) (*Document, error) {
	_bdbb, _fed := Open(filename)
	if _fed != nil {
		return nil, _fed
	}
	_bdbb._eeda.Body = _eaba.NewCT_Body()
	return _bdbb, nil
}

// SetCalcOnExit marks if a FormField should be CalcOnExit or not.
func (_cgdc FormField) SetCalcOnExit(calcOnExit bool) {
	_bggbb := _eaba.NewCT_OnOff()
	_bggbb.ValAttr = &_cg.ST_OnOff{Bool: &calcOnExit}
	_cgdc._cdag.CalcOnExit = []*_eaba.CT_OnOff{_bggbb}
}

// Tables returns the tables defined in the document.
func (_aec *Document) Tables() []Table {
	_ged := []Table{}
	if _aec._eeda.Body == nil {
		return nil
	}
	for _, _aggd := range _aec._eeda.Body.EG_BlockLevelElts {
		for _, _dbff := range _aggd.EG_ContentBlockContent {
			for _, _dfd := range _aec.tables(_dbff) {
				_ged = append(_ged, _dfd)
			}
		}
	}
	return _ged
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (_dbbga ParagraphProperties) SetFirstLineIndent(m _eb.Distance) {
	if _dbbga._efec.Ind == nil {
		_dbbga._efec.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_dbbga._efec.Ind.FirstLineAttr = nil
	} else {
		_dbbga._efec.Ind.FirstLineAttr = &_cg.ST_TwipsMeasure{}
		_dbbga._efec.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(m / _eb.Twips))
	}
}

// SetTargetByRef sets the URL target of the hyperlink and is more efficient if a link
// destination will be used many times.
func (_cafc HyperLink) SetTargetByRef(link _ea.Hyperlink) {
	_cafc._dgace.IdAttr = _d.String(_ea.Relationship(link).ID())
	_cafc._dgace.AnchorAttr = nil
}

// AddParagraph adds a paragraph to the endnote.
func (_bbgf Endnote) AddParagraph() Paragraph {
	_cgdd := _eaba.NewEG_ContentBlockContent()
	_gefa := len(_bbgf._ggd.EG_BlockLevelElts[0].EG_ContentBlockContent)
	_bbgf._ggd.EG_BlockLevelElts[0].EG_ContentBlockContent = append(_bbgf._ggd.EG_BlockLevelElts[0].EG_ContentBlockContent, _cgdd)
	_abgg := _eaba.NewCT_P()
	var _ecdb *_eaba.CT_String
	if _gefa != 0 {
		_gefg := len(_bbgf._ggd.EG_BlockLevelElts[0].EG_ContentBlockContent[_gefa-1].P)
		_ecdb = _bbgf._ggd.EG_BlockLevelElts[0].EG_ContentBlockContent[_gefa-1].P[_gefg-1].PPr.PStyle
	} else {
		_ecdb = _eaba.NewCT_String()
		_ecdb.ValAttr = "\u0045n\u0064\u006e\u006f\u0074\u0065"
	}
	_cgdd.P = append(_cgdd.P, _abgg)
	_bfbf := Paragraph{_bbgf._efegf, _abgg}
	_bfbf._dbef.PPr = _eaba.NewCT_PPr()
	_bfbf._dbef.PPr.PStyle = _ecdb
	_bfbf._dbef.PPr.RPr = _eaba.NewCT_ParaRPr()
	return _bfbf
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (_bcde TableBorders) SetInsideHorizontal(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_bcde._dbbcbf.InsideH = _eaba.NewCT_Border()
	_gfdf(_bcde._dbbcbf.InsideH, t, c, thickness)
}

func _cab(_bfa *_eaba.CT_TblWidth, _aea float64) {
	_bfa.TypeAttr = _eaba.ST_TblWidthPct
	_bfa.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_bfa.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_bfa.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(_aea * 50))
}

// SetLayout controls the table layout. wml.ST_TblLayoutTypeAutofit corresponds
// to "Automatically resize to fit contents" being checked, while
// wml.ST_TblLayoutTypeFixed corresponds to it being unchecked.
func (_ffce TableProperties) SetLayout(l _eaba.ST_TblLayoutType) {
	if l == _eaba.ST_TblLayoutTypeUnset || l == _eaba.ST_TblLayoutTypeAutofit {
		_ffce._efac.TblLayout = nil
	} else {
		_ffce._efac.TblLayout = _eaba.NewCT_TblLayoutType()
		_ffce._efac.TblLayout.TypeAttr = l
	}
}

// X returns the inner wml.CT_TblBorders
func (_bgafa TableBorders) X() *_eaba.CT_TblBorders { return _bgafa._dbbcbf }

func _fdgf(_fdce int64, _ccga int64, _bcca string, _edefe *_eaba.CT_NumFmt, _ecac map[int64]int64) string {
	_ddec := _cgeeg(_bcca)
	_ggfd := _efff(_fdce, _edefe)
	_dgddb := int64(0)
	for _abfb, _eddd := range _ddec {
		_ggdb := _eg.Sprintf("\u0025\u0025\u0025\u0064", _abfb+1)
		if len(_ddec) == 1 {
			_ggdb = _eg.Sprintf("\u0025\u0025\u0025\u0064", _ccga+1)
			_ddec[_abfb] = _f.Replace(_eddd, _ggdb, _ggfd, 1)
			break
		}
		if _ccga > 0 && _ccga > _dgddb && _abfb < (len(_ddec)-1) {
			_ecgdb := _efff(_ecac[_dgddb], _edefe)
			_ddec[_abfb] = _f.Replace(_eddd, _ggdb, _ecgdb, 1)
			_dgddb++
		} else {
			_ddec[_abfb] = _f.Replace(_eddd, _ggdb, _ggfd, 1)
		}
	}
	return _f.Join(_ddec, "")
}

// SetStart sets the cell start margin
func (_fdb CellMargins) SetStart(d _eb.Distance) {
	_fdb._ec.Start = _eaba.NewCT_TblWidth()
	_fgf(_fdb._ec.Start, d)
}

// SetRightPct sets the cell right margin
func (_cgb CellMargins) SetRightPct(pct float64) {
	_cgb._ec.Right = _eaba.NewCT_TblWidth()
	_cab(_cgb._ec.Right, pct)
}

// X returns the inner wrapped XML type.
func (_cgab Paragraph) X() *_eaba.CT_P { return _cgab._dbef }

// SetWidth sets the cell width to a specified width.
func (_gfc CellProperties) SetWidth(d _eb.Distance) {
	_gfc._be.TcW = _eaba.NewCT_TblWidth()
	_gfc._be.TcW.TypeAttr = _eaba.ST_TblWidthDxa
	_gfc._be.TcW.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_gfc._be.TcW.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_gfc._be.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(d / _eb.Twips))
}

func (_afff FormFieldType) String() string {
	if _afff >= FormFieldType(len(_fdeac)-1) {
		return _eg.Sprintf("\u0046\u006f\u0072\u006d\u0046\u0069\u0065\u006c\u0064\u0054\u0079\u0070e\u0028\u0025\u0064\u0029", _afff)
	}
	return _abgga[_fdeac[_afff]:_fdeac[_afff+1]]
}

// Row is a row within a table within a document.
type Row struct {
	_eegcb *Document
	_cegac *_eaba.CT_Row
}

// Clear clears all content within a header
func (_ceebd Header) Clear() { _ceebd._cbgbf.EG_ContentBlockContent = nil }

// TableConditionalFormatting controls the conditional formatting within a table
// style.
type TableConditionalFormatting struct{ _agcba *_eaba.CT_TblStylePr }

// RemoveFootnote removes a footnote from both the paragraph and the document
// the requested footnote must be anchored on the paragraph being referenced.
func (_efcgdd Paragraph) RemoveFootnote(id int64) {
	_gfeed := _efcgdd._bbgc._edf
	var _caggd int
	for _adcb, _befc := range _gfeed.CT_Footnotes.Footnote {
		if _befc.IdAttr == id {
			_caggd = _adcb
		}
	}
	_caggd = 0
	_gfeed.CT_Footnotes.Footnote[_caggd] = nil
	_gfeed.CT_Footnotes.Footnote[_caggd] = _gfeed.CT_Footnotes.Footnote[len(_gfeed.CT_Footnotes.Footnote)-1]
	_gfeed.CT_Footnotes.Footnote = _gfeed.CT_Footnotes.Footnote[:len(_gfeed.CT_Footnotes.Footnote)-1]
	var _gggf Run
	for _, _dagg := range _efcgdd.Runs() {
		if _fabf, _dcag := _dagg.IsFootnote(); _fabf {
			if _dcag == id {
				_gggf = _dagg
			}
		}
	}
	_efcgdd.RemoveRun(_gggf)
}

// X returns the inner wrapped XML type.
func (_gcaab ParagraphStyleProperties) X() *_eaba.CT_PPrGeneral { return _gcaab._fded }

// SetSpacing sets the spacing that comes before and after the paragraph.
func (_eacba ParagraphStyleProperties) SetSpacing(before, after _eb.Distance) {
	if _eacba._fded.Spacing == nil {
		_eacba._fded.Spacing = _eaba.NewCT_Spacing()
	}
	if before == _eb.Zero {
		_eacba._fded.Spacing.BeforeAttr = nil
	} else {
		_eacba._fded.Spacing.BeforeAttr = &_cg.ST_TwipsMeasure{}
		_eacba._fded.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(before / _eb.Twips))
	}
	if after == _eb.Zero {
		_eacba._fded.Spacing.AfterAttr = nil
	} else {
		_eacba._fded.Spacing.AfterAttr = &_cg.ST_TwipsMeasure{}
		_eacba._fded.Spacing.AfterAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(after / _eb.Twips))
	}
}

// AddRun adds a run to a paragraph.
func (_efcd Paragraph) AddRun() Run {
	_eeaf := _eaba.NewEG_PContent()
	_efcd._dbef.EG_PContent = append(_efcd._dbef.EG_PContent, _eeaf)
	_dgeb := _eaba.NewEG_ContentRunContent()
	_eeaf.EG_ContentRunContent = append(_eeaf.EG_ContentRunContent, _dgeb)
	_abab := _eaba.NewCT_R()
	_dgeb.R = _abab
	return Run{_efcd._bbgc, _abab}
}

// TableProperties returns the table style properties.
func (_fdbe Style) TableProperties() TableStyleProperties {
	if _fdbe._bgfca.TblPr == nil {
		_fdbe._bgfca.TblPr = _eaba.NewCT_TblPrBase()
	}
	return TableStyleProperties{_fdbe._bgfca.TblPr}
}

// GetImageObjByRelId returns a common.Image with the associated relation ID in the
// document.
func (_dfg *Document) GetImageObjByRelId(relId string) (_ea.Image, error) {
	_abcg := _dfg._bdb.GetTargetByRelId(relId)
	return _dfg.DocBase.GetImageBytesByTarget(_abcg)
}

// Strike returns true if run is striked.
func (_dgfbe RunProperties) Strike() bool { return _beff(_dgfbe._gfbgc.Strike) }

// Index returns the index of the header within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (_gfcee Header) Index() int {
	for _bad, _gabde := range _gfcee._bffa._bbac {
		if _gabde == _gfcee._cbgbf {
			return _bad
		}
	}
	return -1
}

// GetShapeStyle returns string style of the shape in watermark and format it to ShapeStyle.
func (_agdga *WatermarkPicture) GetShapeStyle() _ffc.ShapeStyle {
	if _agdga._ggac != nil && _agdga._ggac.StyleAttr != nil {
		return _ffc.NewShapeStyle(*_agdga._ggac.StyleAttr)
	}
	return _ffc.NewShapeStyle("")
}

// Emboss returns true if run emboss is on.
func (_edgde RunProperties) Emboss() bool { return _beff(_edgde._gfbgc.Emboss) }

// SetOutline sets the run to outlined text.
func (_becga RunProperties) SetOutline(b bool) {
	if !b {
		_becga._gfbgc.Outline = nil
	} else {
		_becga._gfbgc.Outline = _eaba.NewCT_OnOff()
	}
}

func (_bbba Endnote) id() int64 { return _bbba._ggd.IdAttr }

// BodySection returns the default body section used for all preceding
// paragraphs until the previous Section. If there is no previous sections, the
// body section applies to the entire document.
func (_agg *Document) BodySection() Section {
	if _agg._eeda.Body.SectPr == nil {
		_agg._eeda.Body.SectPr = _eaba.NewCT_SectPr()
	}
	return Section{_agg, _agg._eeda.Body.SectPr}
}

// SetHighlight highlights text in a specified color.
func (_cgbf RunProperties) SetHighlight(c _eaba.ST_HighlightColor) {
	_cgbf._gfbgc.Highlight = _eaba.NewCT_Highlight()
	_cgbf._gfbgc.Highlight.ValAttr = c
}

// SetCSTheme sets the font complex script theme.
func (_cfbg Fonts) SetCSTheme(t _eaba.ST_Theme) { _cfbg._aggeg.CsthemeAttr = t }

// SetBehindDoc sets the behindDoc attribute of anchor.
func (_ab AnchoredDrawing) SetBehindDoc(val bool) { _ab._cc.BehindDocAttr = val }

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_cgcf *Document) AddImage(i _ea.Image) (_ea.ImageRef, error) {
	_eeea := _ea.MakeImageRef(i, &_cgcf.DocBase, _cgcf._bdb)
	if i.Data == nil && i.Path == "" {
		return _eeea, _ffa.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _eeea, _ffa.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _eeea, _ffa.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	if i.Path != "" {
		_ggc := _gb.Add(i.Path)
		if _ggc != nil {
			return _eeea, _ggc
		}
	}
	_cgcf.Images = append(_cgcf.Images, _eeea)
	_bfdg := _eg.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_cgcf.Images), i.Format)
	_aggb := _cgcf._bdb.AddRelationship(_bfdg, _d.ImageType)
	_cgcf.ContentTypes.EnsureDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_cgcf.ContentTypes.EnsureDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_cgcf.ContentTypes.EnsureDefault("\u006a\u0070\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_cgcf.ContentTypes.EnsureDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	_cgcf.ContentTypes.EnsureDefault(i.Format, "\u0069\u006d\u0061\u0067\u0065\u002f"+i.Format)
	_eeea.SetRelID(_aggb.X().IdAttr)
	_eeea.SetTarget(_bfdg)
	return _eeea, nil
}

// Bold returns true if run font is bold.
func (_bbda RunProperties) Bold() bool {
	_fagff := _bbda._gfbgc
	return _beff(_fagff.B) || _beff(_fagff.BCs)
}

// SetImprint sets the run to imprinted text.
func (_dbac RunProperties) SetImprint(b bool) {
	if !b {
		_dbac._gfbgc.Imprint = nil
	} else {
		_dbac._gfbgc.Imprint = _eaba.NewCT_OnOff()
	}
}

// RemoveMailMerge removes any mail merge settings
func (_fgcc Settings) RemoveMailMerge() { _fgcc._ecefe.MailMerge = nil }

// SetAll sets all of the borders to a given value.
func (_ebcc TableBorders) SetAll(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_ebcc.SetBottom(t, c, thickness)
	_ebcc.SetLeft(t, c, thickness)
	_ebcc.SetRight(t, c, thickness)
	_ebcc.SetTop(t, c, thickness)
	_ebcc.SetInsideHorizontal(t, c, thickness)
	_ebcc.SetInsideVertical(t, c, thickness)
}

// Color returns the style's Color.
func (_dabg RunProperties) Color() Color {
	if _dabg._gfbgc.Color == nil {
		_dabg._gfbgc.Color = _eaba.NewCT_Color()
	}
	return Color{_dabg._gfbgc.Color}
}

// AddCell adds a cell to a row and returns it
func (_faga Row) AddCell() Cell {
	_gbggf := _eaba.NewEG_ContentCellContent()
	_faga._cegac.EG_ContentCellContent = append(_faga._cegac.EG_ContentCellContent, _gbggf)
	_ccfc := _eaba.NewCT_Tc()
	_gbggf.Tc = append(_gbggf.Tc, _ccfc)
	return Cell{_faga._eegcb, _ccfc}
}

// SetStyle sets the table style name.
func (_feafc TableProperties) SetStyle(name string) {
	if name == "" {
		_feafc._efac.TblStyle = nil
	} else {
		_feafc._efac.TblStyle = _eaba.NewCT_String()
		_feafc._efac.TblStyle.ValAttr = name
	}
}

// ParagraphProperties returns the paragraph style properties.
func (_ddbe Style) ParagraphProperties() ParagraphStyleProperties {
	if _ddbe._bgfca.PPr == nil {
		_ddbe._bgfca.PPr = _eaba.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_ddbe._bgfca.PPr}
}

// Endnote returns the endnote based on the ID; this can be used nicely with
// the run.IsEndnote() functionality.
func (_cbd *Document) Endnote(id int64) Endnote {
	for _, _cbb := range _cbd.Endnotes() {
		if _cbb.id() == id {
			return _cbb
		}
	}
	return Endnote{}
}

// IsBold returns true if the run has been set to bold.
func (_ddfbb RunProperties) IsBold() bool { return _ddfbb.BoldValue() == OnOffValueOn }

// AddWatermarkText adds new watermark text to the document.
func (_dcc *Document) AddWatermarkText(text string) WatermarkText {
	var _cadb []Header
	if _ecfg, _gfge := _dcc.BodySection().GetHeader(_eaba.ST_HdrFtrDefault); _gfge {
		_cadb = append(_cadb, _ecfg)
	}
	if _cfcf, _fdg := _dcc.BodySection().GetHeader(_eaba.ST_HdrFtrEven); _fdg {
		_cadb = append(_cadb, _cfcf)
	}
	if _aeeb, _abg := _dcc.BodySection().GetHeader(_eaba.ST_HdrFtrFirst); _abg {
		_cadb = append(_cadb, _aeeb)
	}
	if len(_cadb) < 1 {
		_aaf := _dcc.AddHeader()
		_dcc.BodySection().SetHeader(_aaf, _eaba.ST_HdrFtrDefault)
		_cadb = append(_cadb, _aaf)
	}
	_adb := NewWatermarkText()
	for _, _abcb := range _cadb {
		_cff := _abcb.Paragraphs()
		if len(_cff) < 1 {
			_bedd := _abcb.AddParagraph()
			_bedd.AddRun().AddText("")
		}
		for _, _acea := range _abcb.X().EG_ContentBlockContent {
			for _, _cega := range _acea.P {
				for _, _gffg := range _cega.EG_PContent {
					for _, _fdeg := range _gffg.EG_ContentRunContent {
						if _fdeg.R == nil {
							continue
						}
						for _, _bgce := range _fdeg.R.EG_RunInnerContent {
							_bgce.Pict = _adb._eefg
							break
						}
					}
				}
			}
		}
	}
	_adb.SetText(text)
	return _adb
}

// Endnote is an individual endnote reference within the document.
type Endnote struct {
	_efegf *Document
	_ggd   *_eaba.CT_FtnEdn
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (_geec ParagraphStyleProperties) SetKeepOnOnePage(b bool) {
	if !b {
		_geec._fded.KeepLines = nil
	} else {
		_geec._fded.KeepLines = _eaba.NewCT_OnOff()
	}
}

// Paragraphs returns the paragraphs defined in an endnote.
func (_cbge Endnote) Paragraphs() []Paragraph {
	_ceff := []Paragraph{}
	for _, _dgdd := range _cbge.content() {
		for _, _daae := range _dgdd.P {
			_ceff = append(_ceff, Paragraph{_cbge._efegf, _daae})
		}
	}
	return _ceff
}

// ClearColor clears the text color.
func (_fddef RunProperties) ClearColor() { _fddef._gfbgc.Color = nil }

// X returns the inner wrapped XML type.
func (_aecea Table) X() *_eaba.CT_Tbl { return _aecea._fecae }

// Footer is a footer for a document section.
type Footer struct {
	_dec  *Document
	_daaa *_eaba.Ftr
}

// SetStartIndent controls the start indent of the paragraph.
func (_dccea ParagraphStyleProperties) SetStartIndent(m _eb.Distance) {
	if _dccea._fded.Ind == nil {
		_dccea._fded.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_dccea._fded.Ind.StartAttr = nil
	} else {
		_dccea._fded.Ind.StartAttr = &_eaba.ST_SignedTwipsMeasure{}
		_dccea._fded.Ind.StartAttr.Int64 = _d.Int64(int64(m / _eb.Twips))
	}
}

// InitializeDefault constructs a default numbering.
func (_bbfb Numbering) InitializeDefault() {
	_fbbdg := _eaba.NewCT_AbstractNum()
	_fbbdg.MultiLevelType = _eaba.NewCT_MultiLevelType()
	_fbbdg.MultiLevelType.ValAttr = _eaba.ST_MultiLevelTypeHybridMultilevel
	_bbfb._ebbdb.AbstractNum = append(_bbfb._ebbdb.AbstractNum, _fbbdg)
	_fbbdg.AbstractNumIdAttr = 1
	const _egacf = 720
	const _facd = 720
	const _gbdc = 360
	for _gcgg := 0; _gcgg < 9; _gcgg++ {
		_aeda := _eaba.NewCT_Lvl()
		_aeda.IlvlAttr = int64(_gcgg)
		_aeda.Start = _eaba.NewCT_DecimalNumber()
		_aeda.Start.ValAttr = 1
		_aeda.NumFmt = _eaba.NewCT_NumFmt()
		_aeda.NumFmt.ValAttr = _eaba.ST_NumberFormatBullet
		_aeda.Suff = _eaba.NewCT_LevelSuffix()
		_aeda.Suff.ValAttr = _eaba.ST_LevelSuffixNothing
		_aeda.LvlText = _eaba.NewCT_LevelText()
		_aeda.LvlText.ValAttr = _d.String("\uf0b7")
		_aeda.LvlJc = _eaba.NewCT_Jc()
		_aeda.LvlJc.ValAttr = _eaba.ST_JcLeft
		_aeda.RPr = _eaba.NewCT_RPr()
		_aeda.RPr.RFonts = _eaba.NewCT_Fonts()
		_aeda.RPr.RFonts.AsciiAttr = _d.String("\u0053\u0079\u006d\u0062\u006f\u006c")
		_aeda.RPr.RFonts.HAnsiAttr = _d.String("\u0053\u0079\u006d\u0062\u006f\u006c")
		_aeda.RPr.RFonts.HintAttr = _eaba.ST_HintDefault
		_aeda.PPr = _eaba.NewCT_PPrGeneral()
		_adfbfg := int64(_gcgg*_facd + _egacf)
		_aeda.PPr.Ind = _eaba.NewCT_Ind()
		_aeda.PPr.Ind.LeftAttr = &_eaba.ST_SignedTwipsMeasure{}
		_aeda.PPr.Ind.LeftAttr.Int64 = _d.Int64(_adfbfg)
		_aeda.PPr.Ind.HangingAttr = &_cg.ST_TwipsMeasure{}
		_aeda.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(_gbdc))
		_fbbdg.Lvl = append(_fbbdg.Lvl, _aeda)
	}
	_dbbf := _eaba.NewCT_Num()
	_dbbf.NumIdAttr = 1
	_dbbf.AbstractNumId = _eaba.NewCT_DecimalNumber()
	_dbbf.AbstractNumId.ValAttr = 1
	_bbfb._ebbdb.Num = append(_bbfb._ebbdb.Num, _dbbf)
}

// Underline returns the type of paragraph underline.
func (_gbcf ParagraphProperties) Underline() _eaba.ST_Underline {
	if _agcgd := _gbcf._efec.RPr.U; _agcgd != nil {
		return _agcgd.ValAttr
	}
	return 0
}

// GetWrapPathStart return wrapPath start value.
func (_ddd AnchorDrawWrapOptions) GetWrapPathStart() *_bb.CT_Point2D { return _ddd._bba }

// X returns the inner wrapped XML type.
func (_afcc Footnote) X() *_eaba.CT_FtnEdn { return _afcc._efba }

// RunProperties returns the RunProperties controlling numbering level font, etc.
func (_cfaf NumberingLevel) RunProperties() RunProperties {
	if _cfaf._fgbfa.RPr == nil {
		_cfaf._fgbfa.RPr = _eaba.NewCT_RPr()
	}
	return RunProperties{_cfaf._fgbfa.RPr}
}

// TableLook returns the table look, or conditional formatting applied to a table style.
func (_adagb TableProperties) TableLook() TableLook {
	if _adagb._efac.TblLook == nil {
		_adagb._efac.TblLook = _eaba.NewCT_TblLook()
	}
	return TableLook{_adagb._efac.TblLook}
}

// Headers returns the headers defined in the document.
func (_dcb *Document) Headers() []Header {
	_gcd := []Header{}
	for _, _bdca := range _dcb._bbac {
		_gcd = append(_gcd, Header{_dcb, _bdca})
	}
	return _gcd
}

// DrawingAnchored returns a slice of AnchoredDrawings.
func (_egge Run) DrawingAnchored() []AnchoredDrawing {
	_aecd := []AnchoredDrawing{}
	for _, _edaad := range _egge._bggd.EG_RunInnerContent {
		if _edaad.Drawing == nil {
			continue
		}
		for _, _ecfa := range _edaad.Drawing.Anchor {
			_aecd = append(_aecd, AnchoredDrawing{_egge._dddbb, _ecfa})
		}
	}
	return _aecd
}

// AddPageBreak adds a page break to a run.
func (_geef Run) AddPageBreak() {
	_edcb := _geef.newIC()
	_edcb.Br = _eaba.NewCT_Br()
	_edcb.Br.TypeAttr = _eaba.ST_BrTypePage
}

// SetStyle sets style to the text in watermark.
func (_eccbe *WatermarkText) SetStyle(style _ffc.TextpathStyle) {
	_accaa := _eccbe.getShape()
	if _eccbe._ggfda != nil {
		_aedf := _eccbe._ggfda.EG_ShapeElements
		if len(_aedf) > 0 && _aedf[0].Textpath != nil {
			var _aacag = style.String()
			_aedf[0].Textpath.StyleAttr = &_aacag
		}
		return
	}
	_fgcge := _eccbe.findNode(_accaa, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
	for _ebdb, _egdg := range _fgcge.Attrs {
		if _egdg.Name.Local == "\u0073\u0074\u0079l\u0065" {
			_fgcge.Attrs[_ebdb].Value = style.String()
		}
	}
}

// InsertStyle insert style to styles.
func (_cded Styles) InsertStyle(ss Style) {
	_cded._abad.Style = append(_cded._abad.Style, ss.X())
}

// TableLook is the conditional formatting associated with a table style that
// has been assigned to a table.
type TableLook struct{ _ffdgf *_eaba.CT_TblLook }

func (_bea *Document) validateBookmarks() error {
	_acaf := make(map[string]struct{})
	for _, _edfd := range _bea.Bookmarks() {
		if _, _bbdbd := _acaf[_edfd.Name()]; _bbdbd {
			return _eg.Errorf("d\u0075\u0070\u006c\u0069\u0063\u0061t\u0065\u0020\u0062\u006f\u006f\u006b\u006d\u0061\u0072k\u0020\u0025\u0073 \u0066o\u0075\u006e\u0064", _edfd.Name())
		}
		_acaf[_edfd.Name()] = struct{}{}
	}
	return nil
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_befb Footer) AddImage(i _ea.Image) (_ea.ImageRef, error) {
	var _baaf _ea.Relationships
	for _decd, _edfa := range _befb._dec._gfdb {
		if _edfa == _befb._daaa {
			_baaf = _befb._dec._efc[_decd]
		}
	}
	_aade := _ea.MakeImageRef(i, &_befb._dec.DocBase, _baaf)
	if i.Data == nil && i.Path == "" {
		return _aade, _ffa.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _aade, _ffa.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _aade, _ffa.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	_befb._dec.Images = append(_befb._dec.Images, _aade)
	_bgae := _eg.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_befb._dec.Images), i.Format)
	_edbb := _baaf.AddRelationship(_bgae, _d.ImageType)
	_aade.SetRelID(_edbb.X().IdAttr)
	return _aade, nil
}

// VerticalAlign returns the value of paragraph vertical align.
func (_dcfe ParagraphProperties) VerticalAlignment() _cg.ST_VerticalAlignRun {
	if _edea := _dcfe._efec.RPr.VertAlign; _edea != nil {
		return _edea.ValAttr
	}
	return 0
}

// PossibleValues returns the possible values for a FormFieldTypeDropDown.
func (_gdaa FormField) PossibleValues() []string {
	if _gdaa._cdag.DdList == nil {
		return nil
	}
	_gdag := []string{}
	for _, _fbfa := range _gdaa._cdag.DdList.ListEntry {
		if _fbfa == nil {
			continue
		}
		_gdag = append(_gdag, _fbfa.ValAttr)
	}
	return _gdag
}

// SetPictureWashout set washout to watermark picture.
func (_eggef *WatermarkPicture) SetPictureWashout(isWashout bool) {
	if _eggef._ggac != nil {
		_aaff := _eggef._ggac.EG_ShapeElements
		if len(_aaff) > 0 && _aaff[0].Imagedata != nil {
			if isWashout {
				_dcfcd := "\u0031\u0039\u0036\u0036\u0031\u0066"
				_gdcd := "\u0032\u0032\u0039\u0033\u0038\u0066"
				_aaff[0].Imagedata.GainAttr = &_dcfcd
				_aaff[0].Imagedata.BlacklevelAttr = &_gdcd
			}
		}
	}
}

func _adff(_bcdc Paragraph) *_eaba.CT_NumPr {
	_bcdc.ensurePPr()
	if _bcdc._dbef.PPr.NumPr == nil {
		return nil
	}
	return _bcdc._dbef.PPr.NumPr
}

// SetNumberingDefinitionByID sets the numbering definition ID directly, which must
// match an ID defined in numbering.xml
func (_fbge Paragraph) SetNumberingDefinitionByID(abstractNumberID int64) {
	_fbge.ensurePPr()
	if _fbge._dbef.PPr.NumPr == nil {
		_fbge._dbef.PPr.NumPr = _eaba.NewCT_NumPr()
	}
	_dcfab := _eaba.NewCT_DecimalNumber()
	_dcfab.ValAttr = int64(abstractNumberID)
	_fbge._dbef.PPr.NumPr.NumId = _dcfab
}

// SearchStylesById returns style by its id.
func (_edbd Styles) SearchStyleById(id string) (Style, bool) {
	for _, _cgbfg := range _edbd._abad.Style {
		if _cgbfg.StyleIdAttr != nil {
			if *_cgbfg.StyleIdAttr == id {
				return Style{_cgbfg}, true
			}
		}
	}
	return Style{}, false
}

// AddBreak adds a line break to a run.
func (_bbcfa Run) AddBreak() { _bggc := _bbcfa.newIC(); _bggc.Br = _eaba.NewCT_Br() }

// SetLeftIndent controls the left indent of the paragraph.
func (_gcefgd ParagraphStyleProperties) SetLeftIndent(m _eb.Distance) {
	if _gcefgd._fded.Ind == nil {
		_gcefgd._fded.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_gcefgd._fded.Ind.LeftAttr = nil
	} else {
		_gcefgd._fded.Ind.LeftAttr = &_eaba.ST_SignedTwipsMeasure{}
		_gcefgd._fded.Ind.LeftAttr.Int64 = _d.Int64(int64(m / _eb.Twips))
	}
}

// NewSettings constructs a new empty Settings
func NewSettings() Settings {
	_faaa := _eaba.NewSettings()
	_faaa.Compat = _eaba.NewCT_Compat()
	_dcbea := _eaba.NewCT_CompatSetting()
	_dcbea.NameAttr = _d.String("\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074y\u004d\u006f\u0064\u0065")
	_dcbea.UriAttr = _d.String("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006fff\u0069\u0063\u0065/\u0077o\u0072\u0064")
	_dcbea.ValAttr = _d.String("\u0031\u0035")
	_faaa.Compat.CompatSetting = append(_faaa.Compat.CompatSetting, _dcbea)
	return Settings{_faaa}
}

// AddDrawingAnchored adds an anchored (floating) drawing from an ImageRef.
func (_dgfgd Run) AddDrawingAnchored(img _ea.ImageRef) (AnchoredDrawing, error) {
	_cegd := _dgfgd.newIC()
	_cegd.Drawing = _eaba.NewCT_Drawing()
	_cafa := _eaba.NewWdAnchor()
	_ebge := AnchoredDrawing{_dgfgd._dddbb, _cafa}
	_cafa.SimplePosAttr = _d.Bool(false)
	_cafa.AllowOverlapAttr = true
	_cafa.CNvGraphicFramePr = _bb.NewCT_NonVisualGraphicFrameProperties()
	_cegd.Drawing.Anchor = append(_cegd.Drawing.Anchor, _cafa)
	_cafa.Graphic = _bb.NewGraphic()
	_cafa.Graphic.GraphicData = _bb.NewCT_GraphicalObjectData()
	_cafa.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"
	_cafa.SimplePos.XAttr.ST_CoordinateUnqualified = _d.Int64(0)
	_cafa.SimplePos.YAttr.ST_CoordinateUnqualified = _d.Int64(0)
	_cafa.PositionH.RelativeFromAttr = _eaba.WdST_RelFromHPage
	_cafa.PositionH.Choice = &_eaba.WdCT_PosHChoice{}
	_cafa.PositionH.Choice.PosOffset = _d.Int32(0)
	_cafa.PositionV.RelativeFromAttr = _eaba.WdST_RelFromVPage
	_cafa.PositionV.Choice = &_eaba.WdCT_PosVChoice{}
	_cafa.PositionV.Choice.PosOffset = _d.Int32(0)
	_cafa.Extent.CxAttr = int64(float64(img.Size().X*_eb.Pixel72) / _eb.EMU)
	_cafa.Extent.CyAttr = int64(float64(img.Size().Y*_eb.Pixel72) / _eb.EMU)
	_cafa.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_cafa.Choice.WrapSquare = _eaba.NewWdCT_WrapSquare()
	_cafa.Choice.WrapSquare.WrapTextAttr = _eaba.WdST_WrapTextBothSides
	_ebde := 0x7FFFFFFF & _a.Uint32()
	_cafa.DocPr.IdAttr = _ebde
	_cfbda := _gd.NewPic()
	_cfbda.NvPicPr.CNvPr.IdAttr = _ebde
	_cfdf := img.RelID()
	if _cfdf == "" {
		return _ebge, _ffa.New("\u0063\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020\u0074\u006f\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0020\u0072\u0065l\u0061\u0074\u0069o\u006e\u0073")
	}
	_cafa.Graphic.GraphicData.Any = append(_cafa.Graphic.GraphicData.Any, _cfbda)
	_cfbda.BlipFill = _bb.NewCT_BlipFillProperties()
	_cfbda.BlipFill.Blip = _bb.NewCT_Blip()
	_cfbda.BlipFill.Blip.EmbedAttr = &_cfdf
	_cfbda.BlipFill.Stretch = _bb.NewCT_StretchInfoProperties()
	_cfbda.BlipFill.Stretch.FillRect = _bb.NewCT_RelativeRect()
	_cfbda.SpPr = _bb.NewCT_ShapeProperties()
	_cfbda.SpPr.Xfrm = _bb.NewCT_Transform2D()
	_cfbda.SpPr.Xfrm.Off = _bb.NewCT_Point2D()
	_cfbda.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _d.Int64(0)
	_cfbda.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _d.Int64(0)
	_cfbda.SpPr.Xfrm.Ext = _bb.NewCT_PositiveSize2D()
	_cfbda.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * _eb.Point)
	_cfbda.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * _eb.Point)
	_cfbda.SpPr.PrstGeom = _bb.NewCT_PresetGeometry2D()
	_cfbda.SpPr.PrstGeom.PrstAttr = _bb.ST_ShapeTypeRect
	return _ebge, nil
}

// CharacterSpacingValue returns the value of run's characters spacing in twips (1/20 of point).
func (_dgeg RunProperties) CharacterSpacingValue() int64 {
	if _ccbd := _dgeg._gfbgc.Spacing; _ccbd != nil {
		_bbabc := _ccbd.ValAttr
		if _bbabc.Int64 != nil {
			return *_bbabc.Int64
		}
	}
	return int64(0)
}

// Styles is the document wide styles contained in styles.xml.
type Styles struct{ _abad *_eaba.Styles }

// AddHyperLink adds a new hyperlink to a parapgraph.
func (_fdcff Paragraph) AddHyperLink() HyperLink {
	_agdcd := _eaba.NewEG_PContent()
	_fdcff._dbef.EG_PContent = append(_fdcff._dbef.EG_PContent, _agdcd)
	_agdcd.Hyperlink = _eaba.NewCT_Hyperlink()
	return HyperLink{_fdcff._bbgc, _agdcd.Hyperlink}
}

// AddStyle adds a new empty style.
func (_cgff Styles) AddStyle(styleID string, t _eaba.ST_StyleType, isDefault bool) Style {
	_ccbc := _eaba.NewCT_Style()
	_ccbc.TypeAttr = t
	if isDefault {
		_ccbc.DefaultAttr = &_cg.ST_OnOff{}
		_ccbc.DefaultAttr.Bool = _d.Bool(isDefault)
	}
	_ccbc.StyleIdAttr = _d.String(styleID)
	_cgff._abad.Style = append(_cgff._abad.Style, _ccbc)
	return Style{_ccbc}
}

// SetPageSizeAndOrientation sets the page size and orientation for a section.
func (_ebfcg Section) SetPageSizeAndOrientation(w, h _eb.Distance, orientation _eaba.ST_PageOrientation) {
	if _ebfcg._aggca.PgSz == nil {
		_ebfcg._aggca.PgSz = _eaba.NewCT_PageSz()
	}
	_ebfcg._aggca.PgSz.OrientAttr = orientation
	if orientation == _eaba.ST_PageOrientationLandscape {
		_ebfcg._aggca.PgSz.WAttr = &_cg.ST_TwipsMeasure{}
		_ebfcg._aggca.PgSz.WAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(h / _eb.Twips))
		_ebfcg._aggca.PgSz.HAttr = &_cg.ST_TwipsMeasure{}
		_ebfcg._aggca.PgSz.HAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(w / _eb.Twips))
	} else {
		_ebfcg._aggca.PgSz.WAttr = &_cg.ST_TwipsMeasure{}
		_ebfcg._aggca.PgSz.WAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(w / _eb.Twips))
		_ebfcg._aggca.PgSz.HAttr = &_cg.ST_TwipsMeasure{}
		_ebfcg._aggca.PgSz.HAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(h / _eb.Twips))
	}
}

// Document is a text document that can be written out in the OOXML .docx
// format. It can be opened from a file on disk and modified, or created from
// scratch.
type Document struct {
	_ea.DocBase
	_eeda     *_eaba.Document
	Settings  Settings
	Numbering Numbering
	Styles    Styles
	_bbac     []*_eaba.Hdr
	_cfg      []_ea.Relationships
	_gfdb     []*_eaba.Ftr
	_efc      []_ea.Relationships
	_bdb      _ea.Relationships
	_bag      []*_bb.Theme
	_feb      *_eaba.WebSettings
	_fdd      *_eaba.Fonts
	_bfd      *_eaba.Endnotes
	_edf      *_eaba.Footnotes
	_aac      []*_gf.Ocx
	_cea      []*chart
	_ebd      string
}

// AddText adds tet to a run.
func (_fdbdg Run) AddText(s string) {
	_ddbce := _eaba.NewEG_RunInnerContent()
	_fdbdg._bggd.EG_RunInnerContent = append(_fdbdg._bggd.EG_RunInnerContent, _ddbce)
	_ddbce.T = _eaba.NewCT_Text()
	if _d.NeedsSpacePreserve(s) {
		_decg := "\u0070\u0072\u0065\u0073\u0065\u0072\u0076\u0065"
		_ddbce.T.SpaceAttr = &_decg
	}
	_ddbce.T.Content = s
}

// X returns the inner wrapped XML type.
func (_fdfa InlineDrawing) X() *_eaba.WdInline { return _fdfa._fbbd }

// Color controls the run or styles color.
type Color struct{ _bcf *_eaba.CT_Color }

// SetUnderline controls underline for a run style.
func (_adbf RunProperties) SetUnderline(style _eaba.ST_Underline, c _bc.Color) {
	if style == _eaba.ST_UnderlineUnset {
		_adbf._gfbgc.U = nil
	} else {
		_adbf._gfbgc.U = _eaba.NewCT_Underline()
		_adbf._gfbgc.U.ColorAttr = &_eaba.ST_HexColor{}
		_adbf._gfbgc.U.ColorAttr.ST_HexColorRGB = c.AsRGBString()
		_adbf._gfbgc.U.ValAttr = style
	}
}

// X returns the inner wrapped XML type.
func (_bbbag Settings) X() *_eaba.Settings { return _bbbag._ecefe }

// SetTextStyleBold set text style of watermark to bold.
func (_dadeg *WatermarkText) SetTextStyleBold(value bool) {
	if _dadeg._ggfda != nil {
		_ffga := _dadeg.GetStyle()
		_ffga.SetBold(value)
		_dadeg.SetStyle(_ffga)
	}
}

// SetEnabled marks a FormField as enabled or disabled.
func (_ddac FormField) SetEnabled(enabled bool) {
	_dabc := _eaba.NewCT_OnOff()
	_dabc.ValAttr = &_cg.ST_OnOff{Bool: &enabled}
	_ddac._cdag.Enabled = []*_eaba.CT_OnOff{_dabc}
}

// SetBeforeAuto controls if spacing before a paragraph is automatically determined.
func (_ebfab ParagraphSpacing) SetBeforeAuto(b bool) {
	if b {
		_ebfab._ebfge.BeforeAutospacingAttr = &_cg.ST_OnOff{}
		_ebfab._ebfge.BeforeAutospacingAttr.Bool = _d.Bool(true)
	} else {
		_ebfab._ebfge.BeforeAutospacingAttr = nil
	}
}

// SetHangingIndent controls the hanging indent of the paragraph.
func (_befd ParagraphStyleProperties) SetHangingIndent(m _eb.Distance) {
	if _befd._fded.Ind == nil {
		_befd._fded.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_befd._fded.Ind.HangingAttr = nil
	} else {
		_befd._fded.Ind.HangingAttr = &_cg.ST_TwipsMeasure{}
		_befd._fded.Ind.HangingAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(m / _eb.Twips))
	}
}

// SetCharacterSpacing sets the run's Character Spacing Adjustment.
func (_gdceb RunProperties) SetCharacterSpacing(size _eb.Distance) {
	_gdceb._gfbgc.Spacing = _eaba.NewCT_SignedTwipsMeasure()
	_gdceb._gfbgc.Spacing.ValAttr.Int64 = _d.Int64(int64(size / _eb.Twips))
}

// SetHeight allows controlling the height of a row within a table.
func (_cggce RowProperties) SetHeight(ht _eb.Distance, rule _eaba.ST_HeightRule) {
	if rule == _eaba.ST_HeightRuleUnset {
		_cggce._bccec.TrHeight = nil
	} else {
		_aaca := _eaba.NewCT_Height()
		_aaca.HRuleAttr = rule
		_aaca.ValAttr = &_cg.ST_TwipsMeasure{}
		_aaca.ValAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(ht / _eb.Twips))
		_cggce._bccec.TrHeight = []*_eaba.CT_Height{_aaca}
	}
}

func (_edgc *Document) InsertTableAfter(relativeTo Paragraph) Table {
	return _edgc.insertTable(relativeTo, false)
}

// SetKeepWithNext controls if this paragraph should be kept with the next.
func (_dfcg ParagraphProperties) SetKeepWithNext(b bool) {
	if !b {
		_dfcg._efec.KeepNext = nil
	} else {
		_dfcg._efec.KeepNext = _eaba.NewCT_OnOff()
	}
}

// SetBasedOn sets the style that this style is based on.
func (_cedf Style) SetBasedOn(name string) {
	if name == "" {
		_cedf._bgfca.BasedOn = nil
	} else {
		_cedf._bgfca.BasedOn = _eaba.NewCT_String()
		_cedf._bgfca.BasedOn.ValAttr = name
	}
}

// X returns the inner wrapped XML type.
func (_ffdg Endnote) X() *_eaba.CT_FtnEdn { return _ffdg._ggd }

// Font returns the name of run font family.
func (_gdege RunProperties) Font() string {
	if _ebea := _gdege._gfbgc.RFonts; _ebea != nil {
		if _ebea.AsciiAttr != nil {
			return *_ebea.AsciiAttr
		} else if _ebea.HAnsiAttr != nil {
			return *_ebea.HAnsiAttr
		} else if _ebea.CsAttr != nil {
			return *_ebea.CsAttr
		}
	}
	return ""
}

// PutNodeAfter put node to position after relativeTo.
func (_acdfa *Document) PutNodeAfter(relativeTo, node Node) {
	_acdfa.putNode(relativeTo, node, false)
}

// Properties returns the numbering level paragraph properties.
func (_gfbf NumberingLevel) Properties() ParagraphStyleProperties {
	if _gfbf._fgbfa.PPr == nil {
		_gfbf._fgbfa.PPr = _eaba.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_gfbf._fgbfa.PPr}
}

type mergeFieldInfo struct {
	_ggce                 string
	_caaee                string
	_ddceg                string
	_aced                 bool
	_gfcc                 bool
	_cecf                 bool
	_cbccd                bool
	_bfbb                 Paragraph
	_abacb, _ddbag, _feec int
	_acfca                *_eaba.EG_PContent
	_fddee                bool
}

func (_bbd *chart) Target() string { return _bbd._fea }

// Caps returns true if run font is capitalized.
func (_daef RunProperties) Caps() bool { return _beff(_daef._gfbgc.Caps) }

func _efcgda(_ebc *_eaba.CT_OnOff) OnOffValue {
	if _ebc == nil {
		return OnOffValueUnset
	}
	if _ebc.ValAttr != nil && _ebc.ValAttr.Bool != nil && *_ebc.ValAttr.Bool == false {
		return OnOffValueOff
	}
	return OnOffValueOn
}

// X returns the inner wrapped XML type.
func (_addec HyperLink) X() *_eaba.CT_Hyperlink { return _addec._dgace }

func _fgg(_cabd *_eaba.CT_Tbl, _gcge *_eaba.CT_P, _agb bool) *_eaba.CT_Tbl {
	for _, _daec := range _cabd.EG_ContentRowContent {
		for _, _cabe := range _daec.Tr {
			for _, _eaac := range _cabe.EG_ContentCellContent {
				for _, _gff := range _eaac.Tc {
					for _fcda, _gcf := range _gff.EG_BlockLevelElts {
						for _, _dgc := range _gcf.EG_ContentBlockContent {
							for _acc, _ccc := range _dgc.P {
								if _ccc == _gcge {
									_gfec := _eaba.NewEG_BlockLevelElts()
									_aabg := _eaba.NewEG_ContentBlockContent()
									_gfec.EG_ContentBlockContent = append(_gfec.EG_ContentBlockContent, _aabg)
									_ffcb := _eaba.NewCT_Tbl()
									_aabg.Tbl = append(_aabg.Tbl, _ffcb)
									_gff.EG_BlockLevelElts = append(_gff.EG_BlockLevelElts, nil)
									if _agb {
										copy(_gff.EG_BlockLevelElts[_fcda+1:], _gff.EG_BlockLevelElts[_fcda:])
										_gff.EG_BlockLevelElts[_fcda] = _gfec
										if _acc != 0 {
											_egb := _eaba.NewEG_BlockLevelElts()
											_ad := _eaba.NewEG_ContentBlockContent()
											_egb.EG_ContentBlockContent = append(_egb.EG_ContentBlockContent, _ad)
											_ad.P = _dgc.P[:_acc]
											_gff.EG_BlockLevelElts = append(_gff.EG_BlockLevelElts, nil)
											copy(_gff.EG_BlockLevelElts[_fcda+1:], _gff.EG_BlockLevelElts[_fcda:])
											_gff.EG_BlockLevelElts[_fcda] = _egb
										}
										_dgc.P = _dgc.P[_acc:]
									} else {
										copy(_gff.EG_BlockLevelElts[_fcda+2:], _gff.EG_BlockLevelElts[_fcda+1:])
										_gff.EG_BlockLevelElts[_fcda+1] = _gfec
										if _acc != len(_dgc.P)-1 {
											_efcf := _eaba.NewEG_BlockLevelElts()
											_eedd := _eaba.NewEG_ContentBlockContent()
											_efcf.EG_ContentBlockContent = append(_efcf.EG_ContentBlockContent, _eedd)
											_eedd.P = _dgc.P[_acc+1:]
											_gff.EG_BlockLevelElts = append(_gff.EG_BlockLevelElts, nil)
											copy(_gff.EG_BlockLevelElts[_fcda+3:], _gff.EG_BlockLevelElts[_fcda+2:])
											_gff.EG_BlockLevelElts[_fcda+2] = _efcf
										} else {
											_fad := _eaba.NewEG_BlockLevelElts()
											_gfgf := _eaba.NewEG_ContentBlockContent()
											_fad.EG_ContentBlockContent = append(_fad.EG_ContentBlockContent, _gfgf)
											_gfgf.P = []*_eaba.CT_P{_eaba.NewCT_P()}
											_gff.EG_BlockLevelElts = append(_gff.EG_BlockLevelElts, nil)
											copy(_gff.EG_BlockLevelElts[_fcda+3:], _gff.EG_BlockLevelElts[_fcda+2:])
											_gff.EG_BlockLevelElts[_fcda+2] = _fad
										}
										_dgc.P = _dgc.P[:_acc+1]
									}
									return _ffcb
								}
							}
							for _, _gbgc := range _dgc.Tbl {
								_gagc := _fgg(_gbgc, _gcge, _agb)
								if _gagc != nil {
									return _gagc
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// SetYOffset sets the Y offset for an image relative to the origin.
func (_af AnchoredDrawing) SetYOffset(y _eb.Distance) {
	_af._cc.PositionV.Choice = &_eaba.WdCT_PosVChoice{}
	_af._cc.PositionV.Choice.PosOffset = _d.Int32(int32(y / _eb.EMU))
}

// Pict returns the pict object.
func (_fgeba *WatermarkPicture) Pict() *_eaba.CT_Picture { return _fgeba._dadb }

// SetLeftIndent controls left indent of paragraph.
func (_gebba Paragraph) SetLeftIndent(m _eb.Distance) {
	_gebba.ensurePPr()
	_dgddc := _gebba._dbef.PPr
	if _dgddc.Ind == nil {
		_dgddc.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_dgddc.Ind.LeftAttr = nil
	} else {
		_dgddc.Ind.LeftAttr = &_eaba.ST_SignedTwipsMeasure{}
		_dgddc.Ind.LeftAttr.Int64 = _d.Int64(int64(m / _eb.Twips))
	}
}

// SetSize sets the font size for a run.
func (_efgc RunProperties) SetSize(size _eb.Distance) {
	_efgc._gfbgc.Sz = _eaba.NewCT_HpsMeasure()
	_efgc._gfbgc.Sz.ValAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(size / _eb.HalfPoint))
	_efgc._gfbgc.SzCs = _eaba.NewCT_HpsMeasure()
	_efgc._gfbgc.SzCs.ValAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(size / _eb.HalfPoint))
}

// SetHAlignment sets the horizontal alignment for an anchored image.
func (_ag AnchoredDrawing) SetHAlignment(h _eaba.WdST_AlignH) {
	_ag._cc.PositionH.Choice = &_eaba.WdCT_PosHChoice{}
	_ag._cc.PositionH.Choice.Align = h
}

// SetRight sets the cell right margin
func (_bda CellMargins) SetRight(d _eb.Distance) {
	_bda._ec.Right = _eaba.NewCT_TblWidth()
	_fgf(_bda._ec.Right, d)
}

// NewAnchorDrawWrapOptions return anchor drawing options property.
func NewAnchorDrawWrapOptions() *AnchorDrawWrapOptions {
	_cgg := &AnchorDrawWrapOptions{}
	if !_cgg._fb {
		_bbb, _gcb := _bg()
		_cgg._bba = _bbb
		_cgg._ga = _gcb
	}
	return _cgg
}

// AddDrawingInline adds an inline drawing from an ImageRef.
func (_fgbea Run) AddDrawingInline(img _ea.ImageRef) (InlineDrawing, error) {
	_bbae := _fgbea.newIC()
	_bbae.Drawing = _eaba.NewCT_Drawing()
	_acge := _eaba.NewWdInline()
	_bbcfaf := InlineDrawing{_fgbea._dddbb, _acge}
	_acge.CNvGraphicFramePr = _bb.NewCT_NonVisualGraphicFrameProperties()
	_bbae.Drawing.Inline = append(_bbae.Drawing.Inline, _acge)
	_acge.Graphic = _bb.NewGraphic()
	_acge.Graphic.GraphicData = _bb.NewCT_GraphicalObjectData()
	_acge.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"
	_acge.DistTAttr = _d.Uint32(0)
	_acge.DistLAttr = _d.Uint32(0)
	_acge.DistBAttr = _d.Uint32(0)
	_acge.DistRAttr = _d.Uint32(0)
	_acge.Extent.CxAttr = int64(float64(img.Size().X*_eb.Pixel72) / _eb.EMU)
	_acge.Extent.CyAttr = int64(float64(img.Size().Y*_eb.Pixel72) / _eb.EMU)
	_dbad := 0x7FFFFFFF & _a.Uint32()
	_acge.DocPr.IdAttr = _dbad
	_gaebf := _gd.NewPic()
	_gaebf.NvPicPr.CNvPr.IdAttr = _dbad
	_egccc := img.RelID()
	if _egccc == "" {
		return _bbcfaf, _ffa.New("\u0063\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020\u0074\u006f\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0020\u0072\u0065l\u0061\u0074\u0069o\u006e\u0073")
	}
	_acge.Graphic.GraphicData.Any = append(_acge.Graphic.GraphicData.Any, _gaebf)
	_gaebf.BlipFill = _bb.NewCT_BlipFillProperties()
	_gaebf.BlipFill.Blip = _bb.NewCT_Blip()
	_gaebf.BlipFill.Blip.EmbedAttr = &_egccc
	_gaebf.BlipFill.Stretch = _bb.NewCT_StretchInfoProperties()
	_gaebf.BlipFill.Stretch.FillRect = _bb.NewCT_RelativeRect()
	_gaebf.SpPr = _bb.NewCT_ShapeProperties()
	_gaebf.SpPr.Xfrm = _bb.NewCT_Transform2D()
	_gaebf.SpPr.Xfrm.Off = _bb.NewCT_Point2D()
	_gaebf.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _d.Int64(0)
	_gaebf.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _d.Int64(0)
	_gaebf.SpPr.Xfrm.Ext = _bb.NewCT_PositiveSize2D()
	_gaebf.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * _eb.Point)
	_gaebf.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * _eb.Point)
	_gaebf.SpPr.PrstGeom = _bb.NewCT_PresetGeometry2D()
	_gaebf.SpPr.PrstGeom.PrstAttr = _bb.ST_ShapeTypeRect
	return _bbcfaf, nil
}

// TableBorders allows manipulation of borders on a table.
type TableBorders struct{ _dbbcbf *_eaba.CT_TblBorders }

// AppendNode append node to document element.
func (_aebg *Document) AppendNode(node Node) {
	_aebg.insertImageFromNode(node)
	_aebg.insertStyleFromNode(node)
	for _, _bcbfb := range node.Children {
		_aebg.insertImageFromNode(_bcbfb)
		_aebg.insertStyleFromNode(_bcbfb)
	}
	switch _bace := node.X().(type) {
	case *Paragraph:
		_aebg.appendParagraph(nil, *_bace, false)
	case *Table:
		_aebg.appendTable(nil, *_bace, false)
	}
	if node._abcbe != nil {
		if node._abcbe._bag != nil {
			if _caff := _aebg._bdb.FindRIDForN(0, _d.ThemeType); _caff == "" {
				if _cffd := node._abcbe._bdb.FindRIDForN(0, _d.ThemeType); _cffd != "" {
					_aebg._bag = append(_aebg._bag, node._abcbe._bag...)
					_gafg := node._abcbe._bdb.GetTargetByRelId(_cffd)
					_aebg.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_gafg, "\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e/\u0076\u006e\u0064.\u006f\u0070e\u006e\u0078\u006d\u006c\u0066\u006fr\u006dat\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0074\u0068\u0065\u006d\u0065\u002b\u0078\u006d\u006c")
					_aebg._bdb.AddRelationship(_gafg, _d.ThemeType)
				}
			}
		}
		_bgfg := _aebg._fdd
		_gef := node._abcbe._fdd
		if _bgfg != nil {
			if _gef != nil {
				if _bgfg.Font != nil {
					if _gef.Font != nil {
						for _, _abbd := range _gef.Font {
							_abbb := true
							for _, _dcgbc := range _bgfg.Font {
								if _dcgbc.NameAttr == _abbd.NameAttr {
									_abbb = false
									break
								}
							}
							if _abbb {
								_bgfg.Font = append(_bgfg.Font, _abbd)
							}
						}
					}
				} else {
					_bgfg.Font = _gef.Font
				}
			}
		} else if _gef != nil {
			_bgfg = _gef
		}
		_aebg._fdd = _bgfg
		if _fgdg := _aebg._bdb.FindRIDForN(0, _d.FontTableType); _fgdg == "" {
			_aebg.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072d/\u0066\u006f\u006e\u0074\u0054\u0061\u0062\u006c\u0065\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002f\u0076n\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063e\u0073\u0073\u0069\u006e\u0067\u006d\u006c\u002e\u0066\u006f\u006e\u0074T\u0061\u0062\u006c\u0065\u002b\u0078m\u006c")
			_aebg._bdb.AddRelationship("\u0066\u006f\u006e\u0074\u0054\u0061\u0062\u006c\u0065\u002e\u0078\u006d\u006c", _d.FontTableType)
		}
	}
}

// SetUnhideWhenUsed controls if a semi hidden style becomes visible when used.
func (_fffg Style) SetUnhideWhenUsed(b bool) {
	if b {
		_fffg._bgfca.UnhideWhenUsed = _eaba.NewCT_OnOff()
	} else {
		_fffg._bgfca.UnhideWhenUsed = nil
	}
}

// UnderlineColor returns the hex color value of run underline.
func (_bbad RunProperties) UnderlineColor() string {
	if _fccbd := _bbad._gfbgc.U; _fccbd != nil {
		_acaea := _fccbd.ColorAttr
		if _acaea != nil && _acaea.ST_HexColorRGB != nil {
			return *_acaea.ST_HexColorRGB
		}
	}
	return ""
}

// SetHeader sets a section header.
func (_bfacc Section) SetHeader(h Header, t _eaba.ST_HdrFtr) {
	_cbfb := _eaba.NewEG_HdrFtrReferences()
	_bfacc._aggca.EG_HdrFtrReferences = append(_bfacc._aggca.EG_HdrFtrReferences, _cbfb)
	_cbfb.HeaderReference = _eaba.NewCT_HdrFtrRef()
	_cbfb.HeaderReference.TypeAttr = t
	_egeag := _bfacc._ceec._bdb.FindRIDForN(h.Index(), _d.HeaderType)
	if _egeag == "" {
		_dcd.Log.Debug("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u0068\u0065\u0061\u0064\u0065r \u0049\u0044")
	}
	_cbfb.HeaderReference.IdAttr = _egeag
}

// Type returns the type of the style.
func (_cbbe Style) Type() _eaba.ST_StyleType { return _cbbe._bgfca.TypeAttr }

// AnchoredDrawing is an absolutely positioned image within a document page.
type AnchoredDrawing struct {
	_fa *Document
	_cc *_eaba.WdAnchor
}

// SetTop sets the cell top margin
func (_ffg CellMargins) SetTop(d _eb.Distance) {
	_ffg._ec.Top = _eaba.NewCT_TblWidth()
	_fgf(_ffg._ec.Top, d)
}

// Outline returns true if paragraph outline is on.
func (_bcff ParagraphProperties) Outline() bool { return _beff(_bcff._efec.RPr.Outline) }

// Fonts returns the style's Fonts.
func (_bffb RunProperties) Fonts() Fonts {
	if _bffb._gfbgc.RFonts == nil {
		_bffb._gfbgc.RFonts = _eaba.NewCT_Fonts()
	}
	return Fonts{_bffb._gfbgc.RFonts}
}

// SetAllCaps sets the run to all caps.
func (_gbdb RunProperties) SetAllCaps(b bool) {
	if !b {
		_gbdb._gfbgc.Caps = nil
	} else {
		_gbdb._gfbgc.Caps = _eaba.NewCT_OnOff()
	}
}

func (_eaee Document) mergeFields() []mergeFieldInfo {
	_babd := []Paragraph{}
	_efdc := []mergeFieldInfo{}
	for _, _aebcf := range _eaee.Tables() {
		for _, _cffed := range _aebcf.Rows() {
			for _, _cdgd := range _cffed.Cells() {
				_babd = append(_babd, _cdgd.Paragraphs()...)
			}
		}
	}
	_babd = append(_babd, _eaee.Paragraphs()...)
	for _, _afaa := range _babd {
		_bead := _afaa.Runs()
		_edce := -1
		_cedg := -1
		_adffb := -1
		_fcca := mergeFieldInfo{}
		for _, _cgdaf := range _afaa._dbef.EG_PContent {
			for _, _ggeae := range _cgdaf.FldSimple {
				if _f.Contains(_ggeae.InstrAttr, "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044") {
					_eefbg := _gfgd(_ggeae.InstrAttr)
					_eefbg._fddee = true
					_eefbg._bfbb = _afaa
					_eefbg._acfca = _cgdaf
					_efdc = append(_efdc, _eefbg)
				}
			}
		}
		for _edae := 0; _edae < len(_bead); _edae++ {
			_dbbcb := _bead[_edae]
			for _, _egdb := range _dbbcb.X().EG_RunInnerContent {
				if _egdb.FldChar != nil {
					switch _egdb.FldChar.FldCharTypeAttr {
					case _eaba.ST_FldCharTypeBegin:
						_edce = _edae
					case _eaba.ST_FldCharTypeSeparate:
						_cedg = _edae
					case _eaba.ST_FldCharTypeEnd:
						_adffb = _edae
						if _fcca._ggce != "" {
							_fcca._bfbb = _afaa
							_fcca._abacb = _edce
							_fcca._feec = _adffb
							_fcca._ddbag = _cedg
							_efdc = append(_efdc, _fcca)
						}
						_edce = -1
						_cedg = -1
						_adffb = -1
						_fcca = mergeFieldInfo{}
					}
				} else if _egdb.InstrText != nil && _f.Contains(_egdb.InstrText.Content, "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044") {
					if _edce != -1 && _adffb == -1 {
						_fcca = _gfgd(_egdb.InstrText.Content)
					}
				}
			}
		}
	}
	return _efdc
}

// SetHangingIndent controls the indentation of the non-first lines in a paragraph.
func (_cbfgf ParagraphProperties) SetHangingIndent(m _eb.Distance) {
	if _cbfgf._efec.Ind == nil {
		_cbfgf._efec.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_cbfgf._efec.Ind.HangingAttr = nil
	} else {
		_cbfgf._efec.Ind.HangingAttr = &_cg.ST_TwipsMeasure{}
		_cbfgf._efec.Ind.HangingAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(m / _eb.Twips))
	}
}

// Footnotes returns the footnotes defined in the document.
func (_age *Document) Footnotes() []Footnote {
	_gfb := []Footnote{}
	for _, _gabe := range _age._edf.CT_Footnotes.Footnote {
		_gfb = append(_gfb, Footnote{_age, _gabe})
	}
	return _gfb
}

// MergeFields returns the list of all mail merge fields found in the document.
func (_ecbe Document) MergeFields() []string {
	_agaa := map[string]struct{}{}
	for _, _dcbe := range _ecbe.mergeFields() {
		_agaa[_dcbe._ggce] = struct{}{}
	}
	_aedd := []string{}
	for _gdgb := range _agaa {
		_aedd = append(_aedd, _gdgb)
	}
	return _aedd
}

// X returns the inner wrapped XML type.
func (_ccefg TableConditionalFormatting) X() *_eaba.CT_TblStylePr { return _ccefg._agcba }

func (_agdf Paragraph) addBeginFldChar(_aedc string) *_eaba.CT_FFData {
	_gfbg := _agdf.addFldChar()
	_gfbg.FldCharTypeAttr = _eaba.ST_FldCharTypeBegin
	_gfbg.FfData = _eaba.NewCT_FFData()
	_bcede := _eaba.NewCT_FFName()
	_bcede.ValAttr = &_aedc
	_gfbg.FfData.Name = []*_eaba.CT_FFName{_bcede}
	return _gfbg.FfData
}

// AddDropdownList adds dropdown list form field to the paragraph and returns it.
func (_ddbg Paragraph) AddDropdownList(name string) FormField {
	_ebda := _ddbg.addFldCharsForField(name, "\u0046\u004f\u0052M\u0044\u0052\u004f\u0050\u0044\u004f\u0057\u004e")
	_ebda._cdag.DdList = _eaba.NewCT_FFDDList()
	return _ebda
}

func (_ebca Styles) initializeStyleDefaults() {
	_aaba := _ebca.AddStyle("\u004e\u006f\u0072\u006d\u0061\u006c", _eaba.ST_StyleTypeParagraph, true)
	_aaba.SetName("\u004e\u006f\u0072\u006d\u0061\u006c")
	_aaba.SetPrimaryStyle(true)
	_adeg := _ebca.AddStyle("D\u0065f\u0061\u0075\u006c\u0074\u0050\u0061\u0072\u0061g\u0072\u0061\u0070\u0068Fo\u006e\u0074", _eaba.ST_StyleTypeCharacter, true)
	_adeg.SetName("\u0044\u0065\u0066\u0061ul\u0074\u0020\u0050\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u0046\u006fn\u0074")
	_adeg.SetUISortOrder(1)
	_adeg.SetSemiHidden(true)
	_adeg.SetUnhideWhenUsed(true)
	_ccfag := _ebca.AddStyle("\u0054i\u0074\u006c\u0065\u0043\u0068\u0061r", _eaba.ST_StyleTypeCharacter, false)
	_ccfag.SetName("\u0054\u0069\u0074\u006c\u0065\u0020\u0043\u0068\u0061\u0072")
	_ccfag.SetBasedOn(_adeg.StyleID())
	_ccfag.SetLinkedStyle("\u0054\u0069\u0074l\u0065")
	_ccfag.SetUISortOrder(10)
	_ccfag.RunProperties().Fonts().SetASCIITheme(_eaba.ST_ThemeMajorAscii)
	_ccfag.RunProperties().Fonts().SetEastAsiaTheme(_eaba.ST_ThemeMajorEastAsia)
	_ccfag.RunProperties().Fonts().SetHANSITheme(_eaba.ST_ThemeMajorHAnsi)
	_ccfag.RunProperties().Fonts().SetCSTheme(_eaba.ST_ThemeMajorBidi)
	_ccfag.RunProperties().SetSize(28 * _eb.Point)
	_ccfag.RunProperties().SetKerning(14 * _eb.Point)
	_ccfag.RunProperties().SetCharacterSpacing(-10 * _eb.Twips)
	_gbga := _ebca.AddStyle("\u0054\u0069\u0074l\u0065", _eaba.ST_StyleTypeParagraph, false)
	_gbga.SetName("\u0054\u0069\u0074l\u0065")
	_gbga.SetBasedOn(_aaba.StyleID())
	_gbga.SetNextStyle(_aaba.StyleID())
	_gbga.SetLinkedStyle(_ccfag.StyleID())
	_gbga.SetUISortOrder(10)
	_gbga.SetPrimaryStyle(true)
	_gbga.ParagraphProperties().SetContextualSpacing(true)
	_gbga.RunProperties().Fonts().SetASCIITheme(_eaba.ST_ThemeMajorAscii)
	_gbga.RunProperties().Fonts().SetEastAsiaTheme(_eaba.ST_ThemeMajorEastAsia)
	_gbga.RunProperties().Fonts().SetHANSITheme(_eaba.ST_ThemeMajorHAnsi)
	_gbga.RunProperties().Fonts().SetCSTheme(_eaba.ST_ThemeMajorBidi)
	_gbga.RunProperties().SetSize(28 * _eb.Point)
	_gbga.RunProperties().SetKerning(14 * _eb.Point)
	_gbga.RunProperties().SetCharacterSpacing(-10 * _eb.Twips)
	_efegb := _ebca.AddStyle("T\u0061\u0062\u006c\u0065\u004e\u006f\u0072\u006d\u0061\u006c", _eaba.ST_StyleTypeTable, false)
	_efegb.SetName("\u004e\u006f\u0072m\u0061\u006c\u0020\u0054\u0061\u0062\u006c\u0065")
	_efegb.SetUISortOrder(99)
	_efegb.SetSemiHidden(true)
	_efegb.SetUnhideWhenUsed(true)
	_efegb.X().TblPr = _eaba.NewCT_TblPrBase()
	_bggag := NewTableWidth()
	_efegb.X().TblPr.TblInd = _bggag.X()
	_bggag.SetValue(0 * _eb.Dxa)
	_efegb.X().TblPr.TblCellMar = _eaba.NewCT_TblCellMar()
	_bggag = NewTableWidth()
	_efegb.X().TblPr.TblCellMar.Top = _bggag.X()
	_bggag.SetValue(0 * _eb.Dxa)
	_bggag = NewTableWidth()
	_efegb.X().TblPr.TblCellMar.Bottom = _bggag.X()
	_bggag.SetValue(0 * _eb.Dxa)
	_bggag = NewTableWidth()
	_efegb.X().TblPr.TblCellMar.Left = _bggag.X()
	_bggag.SetValue(108 * _eb.Dxa)
	_bggag = NewTableWidth()
	_efegb.X().TblPr.TblCellMar.Right = _bggag.X()
	_bggag.SetValue(108 * _eb.Dxa)
	_aagfc := _ebca.AddStyle("\u004e\u006f\u004c\u0069\u0073\u0074", _eaba.ST_StyleTypeNumbering, false)
	_aagfc.SetName("\u004eo\u0020\u004c\u0069\u0073\u0074")
	_aagfc.SetUISortOrder(1)
	_aagfc.SetSemiHidden(true)
	_aagfc.SetUnhideWhenUsed(true)
	_gdcga := []_eb.Distance{16, 13, 12, 11, 11, 11, 11, 11, 11}
	_cbdee := []_eb.Distance{240, 40, 40, 40, 40, 40, 40, 40, 40}
	for _edde := 0; _edde < 9; _edde++ {
		_fgccb := _eg.Sprintf("\u0048e\u0061\u0064\u0069\u006e\u0067\u0025d", _edde+1)
		_eccgf := _ebca.AddStyle(_fgccb+"\u0043\u0068\u0061\u0072", _eaba.ST_StyleTypeCharacter, false)
		_eccgf.SetName(_eg.Sprintf("\u0048e\u0061d\u0069\u006e\u0067\u0020\u0025\u0064\u0020\u0043\u0068\u0061\u0072", _edde+1))
		_eccgf.SetBasedOn(_adeg.StyleID())
		_eccgf.SetLinkedStyle(_fgccb)
		_eccgf.SetUISortOrder(9 + _edde)
		_eccgf.RunProperties().SetSize(_gdcga[_edde] * _eb.Point)
		_ffdbe := _ebca.AddStyle(_fgccb, _eaba.ST_StyleTypeParagraph, false)
		_ffdbe.SetName(_eg.Sprintf("\u0068\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0064", _edde+1))
		_ffdbe.SetNextStyle(_aaba.StyleID())
		_ffdbe.SetLinkedStyle(_ffdbe.StyleID())
		_ffdbe.SetUISortOrder(9 + _edde)
		_ffdbe.SetPrimaryStyle(true)
		_ffdbe.ParagraphProperties().SetKeepNext(true)
		_ffdbe.ParagraphProperties().SetSpacing(_cbdee[_edde]*_eb.Twips, 0)
		_ffdbe.ParagraphProperties().SetOutlineLevel(_edde)
		_ffdbe.RunProperties().SetSize(_gdcga[_edde] * _eb.Point)
	}
}

// SetKerning sets the run's font kerning.
func (_bdbaf RunProperties) SetKerning(size _eb.Distance) {
	_bdbaf._gfbgc.Kern = _eaba.NewCT_HpsMeasure()
	_bdbaf._gfbgc.Kern.ValAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(size / _eb.HalfPoint))
}

// GetHeader gets a section Header for given type t [ST_HdrFtrDefault, ST_HdrFtrEven, ST_HdrFtrFirst]
func (_afdff Section) GetHeader(t _eaba.ST_HdrFtr) (Header, bool) {
	for _, _gdda := range _afdff._aggca.EG_HdrFtrReferences {
		if _gdda.HeaderReference.TypeAttr == t {
			for _, _gabcb := range _afdff._ceec.Headers() {
				_ffdgd := _afdff._ceec._bdb.FindRIDForN(_gabcb.Index(), _d.HeaderType)
				if _ffdgd == _gdda.HeaderReference.IdAttr {
					return _gabcb, true
				}
			}
		}
	}
	return Header{}, false
}

// AddParagraph adds a paragraph to the table cell.
func (_cf Cell) AddParagraph() Paragraph {
	_ce := _eaba.NewEG_BlockLevelElts()
	_cf._acg.EG_BlockLevelElts = append(_cf._acg.EG_BlockLevelElts, _ce)
	_dee := _eaba.NewEG_ContentBlockContent()
	_ce.EG_ContentBlockContent = append(_ce.EG_ContentBlockContent, _dee)
	_baa := _eaba.NewCT_P()
	_dee.P = append(_dee.P, _baa)
	return Paragraph{_cf._fc, _baa}
}

// read reads a document from an io.Reader.
func Read(r _aa.ReaderAt, size int64) (*Document, error) { return _bed(r, size, "") }

// Paragraphs returns the paragraphs defined in a header.
func (_dbdfe Header) Paragraphs() []Paragraph {
	_eeee := []Paragraph{}
	for _, _fdad := range _dbdfe._cbgbf.EG_ContentBlockContent {
		for _, _cebaa := range _fdad.P {
			_eeee = append(_eeee, Paragraph{_dbdfe._bffa, _cebaa})
		}
	}
	for _, _abac := range _dbdfe.Tables() {
		for _, _cgcd := range _abac.Rows() {
			for _, _cfbfe := range _cgcd.Cells() {
				_eeee = append(_eeee, _cfbfe.Paragraphs()...)
			}
		}
	}
	return _eeee
}

// Bookmarks returns all of the bookmarks defined in the document.
func (_fcbe Document) Bookmarks() []Bookmark {
	if _fcbe._eeda.Body == nil {
		return nil
	}
	_adf := []Bookmark{}
	for _, _fgca := range _fcbe._eeda.Body.EG_BlockLevelElts {
		for _, _dcgd := range _fgca.EG_ContentBlockContent {
			for _, _fceag := range _ebb(_dcgd) {
				_adf = append(_adf, _fceag)
			}
		}
	}
	return _adf
}

// SetContextualSpacing controls whether to Ignore Spacing Above and Below When
// Using Identical Styles
func (_edec ParagraphStyleProperties) SetContextualSpacing(b bool) {
	if !b {
		_edec._fded.ContextualSpacing = nil
	} else {
		_edec._fded.ContextualSpacing = _eaba.NewCT_OnOff()
	}
}

// HasEndnotes returns a bool based on the presence or abscence of endnotes within
// the document.
func (_fdbg *Document) HasEndnotes() bool { return _fdbg._bfd != nil }

// CellProperties returns the cell properties.
func (_bagcb TableConditionalFormatting) CellProperties() CellProperties {
	if _bagcb._agcba.TcPr == nil {
		_bagcb._agcba.TcPr = _eaba.NewCT_TcPr()
	}
	return CellProperties{_bagcb._agcba.TcPr}
}

// Style is a style within the styles.xml file.
type Style struct{ _bgfca *_eaba.CT_Style }

func _edbc(_aedg *_eaba.CT_Tbl, _abdd, _dddd map[int64]int64) {
	for _, _aeed := range _aedg.EG_ContentRowContent {
		for _, _cdbdb := range _aeed.Tr {
			for _, _ddcc := range _cdbdb.EG_ContentCellContent {
				for _, _cagg := range _ddcc.Tc {
					for _, _faba := range _cagg.EG_BlockLevelElts {
						for _, _cgca := range _faba.EG_ContentBlockContent {
							for _, _dbdg := range _cgca.P {
								_gfff(_dbdg, _abdd, _dddd)
							}
							for _, _gcdbb := range _cgca.Tbl {
								_edbc(_gcdbb, _abdd, _dddd)
							}
						}
					}
				}
			}
		}
	}
}

// Properties returns the row properties.
func (_accc Row) Properties() RowProperties {
	if _accc._cegac.TrPr == nil {
		_accc._cegac.TrPr = _eaba.NewCT_TrPr()
	}
	return RowProperties{_accc._cegac.TrPr}
}

// X returns the inner wrapped XML type.
func (_bedc TableLook) X() *_eaba.CT_TblLook { return _bedc._ffdgf }

var (
	_cggc = []string{"", "\u0049", "\u0049\u0049", "\u0049\u0049\u0049", "\u0049\u0056", "\u0056", "\u0056\u0049", "\u0056\u0049\u0049", "\u0056\u0049\u0049\u0049", "\u0049\u0058"}
	_cbaf = []string{"", "\u0058", "\u0058\u0058", "\u0058\u0058\u0058", "\u0058\u004c", "\u004c", "\u004c\u0058", "\u004c\u0058\u0058", "\u004c\u0058\u0058\u0058", "\u0058\u0043"}
	_abaa = []string{"", "\u0043", "\u0043\u0043", "\u0043\u0043\u0043", "\u0043\u0044", "\u0044", "\u0044\u0043", "\u0044\u0043\u0043", "\u0044\u0043\u0043\u0043", "\u0043\u004d", "\u004d"}
	_fcfa = []string{"", "\u004d", "\u004d\u004d", "\u004d\u004d\u004d", "\u004d\u004d\u004d\u004d", "\u004d\u004d\u004dM\u004d", "\u004d\u004d\u004d\u004d\u004d\u004d", "\u004dM\u004d\u004d\u004d\u004d\u004d", "\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d", "\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM", "\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"}
	_aegd = []string{"\u006f\u006e\u0065", "\u0074\u0077\u006f", "\u0074\u0068\u0072e\u0065", "\u0066\u006f\u0075\u0072", "\u0066\u0069\u0076\u0065", "\u0073\u0069\u0078", "\u0073\u0065\u0076e\u006e", "\u0065\u0069\u0067h\u0074", "\u006e\u0069\u006e\u0065", "\u0074\u0065\u006e", "\u0065\u006c\u0065\u0076\u0065\u006e", "\u0074\u0077\u0065\u006c\u0076\u0065", "\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e", "\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e", "\u0066i\u0066\u0074\u0065\u0065\u006e", "\u0073i\u0078\u0074\u0065\u0065\u006e", "\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n", "\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e", "\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"}
	_bfcb = []string{"\u0074\u0065\u006e", "\u0074\u0077\u0065\u006e\u0074\u0079", "\u0074\u0068\u0069\u0072\u0074\u0079", "\u0066\u006f\u0072t\u0079", "\u0066\u0069\u0066t\u0079", "\u0073\u0069\u0078t\u0079", "\u0073e\u0076\u0065\u006e\u0074\u0079", "\u0065\u0069\u0067\u0068\u0074\u0079", "\u006e\u0069\u006e\u0065\u0074\u0079"}
	_fdbb = []string{"\u0066\u0069\u0072s\u0074", "\u0073\u0065\u0063\u006f\u006e\u0064", "\u0074\u0068\u0069r\u0064", "\u0066\u006f\u0075\u0072\u0074\u0068", "\u0066\u0069\u0066t\u0068", "\u0073\u0069\u0078t\u0068", "\u0073e\u0076\u0065\u006e\u0074\u0068", "\u0065\u0069\u0067\u0068\u0074\u0068", "\u006e\u0069\u006et\u0068", "\u0074\u0065\u006et\u0068", "\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068", "\u0074w\u0065\u006c\u0066\u0074\u0068", "\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068", "\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068", "\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h", "\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h", "s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068", "\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068", "\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"}
	_bcfa = []string{"\u0074\u0065\u006et\u0068", "\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h", "\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h", "\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068", "\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068", "\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068", "\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068", "\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h", "\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"}
	_dbbc = "\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a"
)

// Section return paragraph properties section value.
func (_gcbf ParagraphProperties) Section() (Section, bool) {
	if _gcbf._efec.SectPr != nil {
		return Section{_gcbf._gaeba, _gcbf._efec.SectPr}, true
	}
	return Section{}, false
}

// HasFootnotes returns a bool based on the presence or abscence of footnotes within
// the document.
func (_gdaf *Document) HasFootnotes() bool { return _gdaf._edf != nil }

// IsItalic returns true if the run has been set to italics.
func (_bdaeb RunProperties) IsItalic() bool { return _bdaeb.ItalicValue() == OnOffValueOn }

func (_eeaa Paragraph) addInstrText(_eccb string) *_eaba.CT_Text {
	_debd := _eeaa.AddRun()
	_dece := _debd.X()
	_cecb := _eaba.NewEG_RunInnerContent()
	_fdbd := _eaba.NewCT_Text()
	_fbbdc := "\u0070\u0072\u0065\u0073\u0065\u0072\u0076\u0065"
	_fdbd.SpaceAttr = &_fbbdc
	_fdbd.Content = "\u0020" + _eccb + "\u0020"
	_cecb.InstrText = _fdbd
	_dece.EG_RunInnerContent = append(_dece.EG_RunInnerContent, _cecb)
	return _fdbd
}

// AddBookmark adds a bookmark to a document that can then be used from a hyperlink. Name is a document
// unique name that identifies the bookmark so it can be referenced from hyperlinks.
func (_ecgfb Paragraph) AddBookmark(name string) Bookmark {
	_dfeeg := _eaba.NewEG_PContent()
	_fef := _eaba.NewEG_ContentRunContent()
	_dfeeg.EG_ContentRunContent = append(_dfeeg.EG_ContentRunContent, _fef)
	_dfadd := _eaba.NewEG_RunLevelElts()
	_fef.EG_RunLevelElts = append(_fef.EG_RunLevelElts, _dfadd)
	_dgbgc := _eaba.NewEG_RangeMarkupElements()
	_fefg := _eaba.NewCT_Bookmark()
	_dgbgc.BookmarkStart = _fefg
	_dfadd.EG_RangeMarkupElements = append(_dfadd.EG_RangeMarkupElements, _dgbgc)
	_dgbgc = _eaba.NewEG_RangeMarkupElements()
	_dgbgc.BookmarkEnd = _eaba.NewCT_MarkupRange()
	_dfadd.EG_RangeMarkupElements = append(_dfadd.EG_RangeMarkupElements, _dgbgc)
	_ecgfb._dbef.EG_PContent = append(_ecgfb._dbef.EG_PContent, _dfeeg)
	_abfc := Bookmark{_fefg}
	_abfc.SetName(name)
	return _abfc
}

// SetVerticalMerge controls the vertical merging of cells.
func (_eaa CellProperties) SetVerticalMerge(mergeVal _eaba.ST_Merge) {
	if mergeVal == _eaba.ST_MergeUnset {
		_eaa._be.VMerge = nil
	} else {
		_eaa._be.VMerge = _eaba.NewCT_VMerge()
		_eaa._be.VMerge.ValAttr = mergeVal
	}
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_fcff TableBorders) SetLeft(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_fcff._dbbcbf.Left = _eaba.NewCT_Border()
	_gfdf(_fcff._dbbcbf.Left, t, c, thickness)
}

// SetAfterSpacing sets spacing below paragraph.
func (_daga Paragraph) SetAfterSpacing(d _eb.Distance) {
	_daga.ensurePPr()
	if _daga._dbef.PPr.Spacing == nil {
		_daga._dbef.PPr.Spacing = _eaba.NewCT_Spacing()
	}
	_dbgb := _daga._dbef.PPr.Spacing
	_dbgb.AfterAttr = &_cg.ST_TwipsMeasure{}
	_dbgb.AfterAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(d / _eb.Twips))
}

// Paragraphs returns the paragraphs defined in a footnote.
func (_gfddc Footnote) Paragraphs() []Paragraph {
	_gabd := []Paragraph{}
	for _, _dega := range _gfddc.content() {
		for _, _gccc := range _dega.P {
			_gabd = append(_gabd, Paragraph{_gfddc._ffdf, _gccc})
		}
	}
	return _gabd
}

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_afgce RunProperties) SizeMeasure() string {
	if _dfgf := _afgce._gfbgc.Sz; _dfgf != nil {
		_fgabg := _dfgf.ValAttr
		if _fgabg.ST_PositiveUniversalMeasure != nil {
			return *_fgabg.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

func _cgeeg(_egfa string) (_afaf []string) {
	for _fbdd := 0; _fbdd < len(_egfa)-2; _fbdd++ {
		if string(_egfa[_fbdd]) == "\u0025" {
			_afaf = append(_afaf, _eg.Sprintf("\u0025\u0073\u0025\u0073\u0025\u0073", string(_egfa[_fbdd]), string(_egfa[_fbdd+1]), string(_egfa[_fbdd+2])))
		}
	}
	return
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_dceg RunProperties) ComplexSizeMeasure() string {
	if _ebec := _dceg._gfbgc.SzCs; _ebec != nil {
		_cggb := _ebec.ValAttr
		if _cggb.ST_PositiveUniversalMeasure != nil {
			return *_cggb.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

func _ffbbg() *_aae.Textpath {
	_cdeage := _aae.NewTextpath()
	_cdeage.OnAttr = _cg.ST_TrueFalseTrue
	_cdeage.FitshapeAttr = _cg.ST_TrueFalseTrue
	return _cdeage
}

// DoubleStrike returns true if paragraph is double striked.
func (_gdbbg ParagraphProperties) DoubleStrike() bool { return _beff(_gdbbg._efec.RPr.Dstrike) }

func _bed(_eead _aa.ReaderAt, _abc int64, _eda string) (*Document, error) {
	const _dfcac = "\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0052\u0065\u0061\u0064"
	if !_b.GetLicenseKey().IsLicensed() && !_gac {
		_eg.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
		_eg.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return nil, _ffa.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	}
	_abe := New()
	_abe.Numbering._ebbdb = nil
	if len(_eda) > 0 {
		_abe._ebd = _eda
	} else {
		_fcdf, _eaaf := _b.GenRefId("\u0064\u0072")
		if _eaaf != nil {
			_dcd.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _eaaf)
			return nil, _eaaf
		}
		_abe._ebd = _fcdf
	}
	if _ecb := _b.Track(_abe._ebd, _dfcac); _ecb != nil {
		_dcd.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ecb)
		return nil, _ecb
	}
	_fcge, _eef := _gb.TempDir("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0064\u006f\u0063\u0078")
	if _eef != nil {
		return nil, _eef
	}
	_abe.TmpPath = _fcge
	_facb, _eef := _da.NewReader(_eead, _abc)
	if _eef != nil {
		return nil, _eg.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073", _eef)
	}
	_bddd := []*_da.File{}
	_bddd = append(_bddd, _facb.File...)
	_gccb := false
	for _, _cfaba := range _bddd {
		if _cfaba.FileHeader.Name == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
			_gccb = true
			break
		}
	}
	if _gccb {
		_abe.CreateCustomProperties()
	}
	_fee := _abe._eeda.ConformanceAttr
	_ecfb := _fd.DecodeMap{}
	_ecfb.SetOnNewRelationshipFunc(_abe.onNewRelationship)
	_ecfb.AddTarget(_d.ContentTypesFilename, _abe.ContentTypes.X(), "", 0)
	_ecfb.AddTarget(_d.BaseRelsFilename, _abe.Rels.X(), "", 0)
	if _dffbg := _ecfb.Decode(_bddd); _dffbg != nil {
		return nil, _dffbg
	}
	_abe._eeda.ConformanceAttr = _fee
	for _, _bgba := range _bddd {
		if _bgba == nil {
			continue
		}
		if _bfae := _abe.AddExtraFileFromZip(_bgba); _bfae != nil {
			return nil, _bfae
		}
	}
	if _gccb {
		_fcee := false
		for _, _cbbb := range _abe.Rels.X().Relationship {
			if _cbbb.TargetAttr == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
				_fcee = true
				break
			}
		}
		if !_fcee {
			_abe.AddCustomRelationships()
		}
	}
	return _abe, nil
}

// RemoveParagraph removes a paragraph from a footer.
func (_fbebe Header) RemoveParagraph(p Paragraph) {
	for _, _abddb := range _fbebe._cbgbf.EG_ContentBlockContent {
		for _bdag, _ecgfe := range _abddb.P {
			if _ecgfe == p._dbef {
				copy(_abddb.P[_bdag:], _abddb.P[_bdag+1:])
				_abddb.P = _abddb.P[0 : len(_abddb.P)-1]
				return
			}
		}
	}
}

func (_addg *WatermarkPicture) findNode(_gbegg *_d.XSDAny, _geaca string) *_d.XSDAny {
	for _, _bbdcb := range _gbegg.Nodes {
		if _bbdcb.XMLName.Local == _geaca {
			return _bbdcb
		}
	}
	return nil
}

// SetVerticalAlignment controls the vertical alignment of the run, this is used
// to control if text is superscript/subscript.
func (_cffee RunProperties) SetVerticalAlignment(v _cg.ST_VerticalAlignRun) {
	if v == _cg.ST_VerticalAlignRunUnset {
		_cffee._gfbgc.VertAlign = nil
	} else {
		_cffee._gfbgc.VertAlign = _eaba.NewCT_VerticalAlignRun()
		_cffee._gfbgc.VertAlign.ValAttr = v
	}
}

// SetBefore sets the spacing that comes before the paragraph.
func (_cefaa ParagraphSpacing) SetBefore(before _eb.Distance) {
	_cefaa._ebfge.BeforeAttr = &_cg.ST_TwipsMeasure{}
	_cefaa._ebfge.BeforeAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(before / _eb.Twips))
}

// Clear clears the styes.
func (_cabcg Styles) Clear() {
	_cabcg._abad.DocDefaults = nil
	_cabcg._abad.LatentStyles = nil
	_cabcg._abad.Style = nil
}

// SetColor sets the text color.
func (_bfbfe RunProperties) SetColor(c _bc.Color) {
	_bfbfe._gfbgc.Color = _eaba.NewCT_Color()
	_bfbfe._gfbgc.Color.ValAttr.ST_HexColorRGB = c.AsRGBString()
}

// SetAlignment sets the paragraph alignment
func (_aegg NumberingLevel) SetAlignment(j _eaba.ST_Jc) {
	if j == _eaba.ST_JcUnset {
		_aegg._fgbfa.LvlJc = nil
	} else {
		_aegg._fgbfa.LvlJc = _eaba.NewCT_Jc()
		_aegg._fgbfa.LvlJc.ValAttr = j
	}
}

// TableStyleProperties are table properties as defined in a style.
type TableStyleProperties struct{ _ddafc *_eaba.CT_TblPrBase }

// Text returns the underlying tet in the run.
func (_gddb Run) Text() string {
	if len(_gddb._bggd.EG_RunInnerContent) == 0 {
		return ""
	}
	_agcb := _cb.Buffer{}
	for _, _aafd := range _gddb._bggd.EG_RunInnerContent {
		if _aafd.T != nil {
			_agcb.WriteString(_aafd.T.Content)
		}
		if _aafd.Tab != nil {
			_agcb.WriteByte('\t')
		}
	}
	return _agcb.String()
}

// Nodes contains slice of Node element.
type Nodes struct{ _agcab []Node }

func _beff(_eccd *_eaba.CT_OnOff) bool { return _eccd != nil }

func _bddda(_bfadg []*_eaba.EG_ContentBlockContent, _gdgaa *TableInfo) []TextItem {
	_gccbd := []TextItem{}
	for _, _agcc := range _bfadg {
		if _cbdd := _agcc.Sdt; _cbdd != nil {
			if _aabba := _cbdd.SdtContent; _aabba != nil {
				_gccbd = append(_gccbd, _bbbef(_aabba.P, _gdgaa, nil)...)
			}
		}
		_gccbd = append(_gccbd, _bbbef(_agcc.P, _gdgaa, nil)...)
		for _, _gabcf := range _agcc.Tbl {
			for _faed, _edfba := range _gabcf.EG_ContentRowContent {
				for _, _gaddd := range _edfba.Tr {
					for _bdcd, _ebbd := range _gaddd.EG_ContentCellContent {
						for _, _bcada := range _ebbd.Tc {
							_dfaf := &TableInfo{Table: _gabcf, Row: _gaddd, Cell: _bcada, RowIndex: _faed, ColIndex: _bdcd}
							for _, _dgbg := range _bcada.EG_BlockLevelElts {
								_gccbd = append(_gccbd, _bddda(_dgbg.EG_ContentBlockContent, _dfaf)...)
							}
						}
					}
				}
			}
		}
	}
	return _gccbd
}

type chart struct {
	_cebb *_dg.ChartSpace
	_ece  string
	_fea  string
}

// Validate validates the structure and in cases where it't possible, the ranges
// of elements within a document. A validation error dones't mean that the
// document won't work in MS Word or LibreOffice, but it's worth checking into.
func (_bgfc *Document) Validate() error {
	if _bgfc == nil || _bgfc._eeda == nil {
		return _ffa.New("\u0064o\u0063\u0075m\u0065\u006e\u0074\u0020n\u006f\u0074\u0020i\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0065d \u0063\u006f\u0072r\u0065\u0063t\u006c\u0079\u002c\u0020\u006e\u0069l\u0020\u0062a\u0073\u0065")
	}
	for _, _ceca := range []func() error{_bgfc.validateTableCells, _bgfc.validateBookmarks} {
		if _cbde := _ceca(); _cbde != nil {
			return _cbde
		}
	}
	if _caba := _bgfc._eeda.Validate(); _caba != nil {
		return _caba
	}
	return nil
}

// Definitions returns the defined numbering definitions.
func (_efab Numbering) Definitions() []NumberingDefinition {
	_gdad := []NumberingDefinition{}
	for _, _ddafd := range _efab._ebbdb.AbstractNum {
		_gdad = append(_gdad, NumberingDefinition{_ddafd})
	}
	return _gdad
}

func (_gebe *Document) appendTable(_bbacf *Paragraph, _adc Table, _dcf bool) Table {
	_fada := _gebe._eeda.Body
	_gfcd := _eaba.NewEG_BlockLevelElts()
	_gebe._eeda.Body.EG_BlockLevelElts = append(_gebe._eeda.Body.EG_BlockLevelElts, _gfcd)
	_faaf := _eaba.NewEG_ContentBlockContent()
	_gfcd.EG_ContentBlockContent = append(_gfcd.EG_ContentBlockContent, _faaf)
	if _bbacf != nil {
		_cabg := _bbacf.X()
		for _cee, _egg := range _fada.EG_BlockLevelElts {
			for _, _ggga := range _egg.EG_ContentBlockContent {
				for _cbca, _bfee := range _faaf.P {
					if _bfee == _cabg {
						_dbfg := _adc.X()
						_effg := _eaba.NewEG_BlockLevelElts()
						_efd := _eaba.NewEG_ContentBlockContent()
						_effg.EG_ContentBlockContent = append(_effg.EG_ContentBlockContent, _efd)
						_efd.Tbl = append(_efd.Tbl, _dbfg)
						_fada.EG_BlockLevelElts = append(_fada.EG_BlockLevelElts, nil)
						if _dcf {
							copy(_fada.EG_BlockLevelElts[_cee+1:], _fada.EG_BlockLevelElts[_cee:])
							_fada.EG_BlockLevelElts[_cee] = _effg
							if _cbca != 0 {
								_eba := _eaba.NewEG_BlockLevelElts()
								_agc := _eaba.NewEG_ContentBlockContent()
								_eba.EG_ContentBlockContent = append(_eba.EG_ContentBlockContent, _agc)
								_agc.P = _ggga.P[:_cbca]
								_fada.EG_BlockLevelElts = append(_fada.EG_BlockLevelElts, nil)
								copy(_fada.EG_BlockLevelElts[_cee+1:], _fada.EG_BlockLevelElts[_cee:])
								_fada.EG_BlockLevelElts[_cee] = _eba
							}
							_ggga.P = _ggga.P[_cbca:]
						} else {
							copy(_fada.EG_BlockLevelElts[_cee+2:], _fada.EG_BlockLevelElts[_cee+1:])
							_fada.EG_BlockLevelElts[_cee+1] = _effg
							if _cbca != len(_ggga.P)-1 {
								_bgf := _eaba.NewEG_BlockLevelElts()
								_gbea := _eaba.NewEG_ContentBlockContent()
								_bgf.EG_ContentBlockContent = append(_bgf.EG_ContentBlockContent, _gbea)
								_gbea.P = _ggga.P[_cbca+1:]
								_fada.EG_BlockLevelElts = append(_fada.EG_BlockLevelElts, nil)
								copy(_fada.EG_BlockLevelElts[_cee+3:], _fada.EG_BlockLevelElts[_cee+2:])
								_fada.EG_BlockLevelElts[_cee+2] = _bgf
							}
							_ggga.P = _ggga.P[:_cbca+1]
						}
						break
					}
				}
				for _, _cbcb := range _ggga.Tbl {
					_dgba := _fgg(_cbcb, _cabg, _dcf)
					if _dgba != nil {
						break
					}
				}
			}
		}
	} else {
		_faaf.Tbl = append(_faaf.Tbl, _adc.X())
	}
	return Table{_gebe, _adc.X()}
}

func (_gebb *chart) X() *_dg.ChartSpace { return _gebb._cebb }

// SetTextWrapNone unsets text wrapping so the image can float on top of the
// text. When used in conjunction with X/Y Offset relative to the page it can be
// used to place a logo at the top of a page at an absolute position that
// doesn't interfere with text.
func (_gdc AnchoredDrawing) SetTextWrapNone() {
	_gdc._cc.Choice = &_eaba.WdEG_WrapTypeChoice{}
	_gdc._cc.Choice.WrapNone = _eaba.NewWdCT_WrapNone()
}

func (_dfad *Document) InsertTableBefore(relativeTo Paragraph) Table {
	return _dfad.insertTable(relativeTo, true)
}

func (_bbbagd *WatermarkPicture) getShape() *_d.XSDAny {
	return _bbbagd.getInnerElement("\u0073\u0068\u0061p\u0065")
}

// FindNodeByCondition return node based on condition function,
// if wholeElements is true, its will extract childs as next node elements.
func (_dfeg *Nodes) FindNodeByCondition(f func(_adaa *Node) bool, wholeElements bool) []Node {
	_bcce := []Node{}
	for _, _fbdf := range _dfeg._agcab {
		if f(&_fbdf) {
			_bcce = append(_bcce, _fbdf)
		}
		if wholeElements {
			_edcee := Nodes{_agcab: _fbdf.Children}
			_bcce = append(_bcce, _edcee.FindNodeByCondition(f, wholeElements)...)
		}
	}
	return _bcce
}

// AddFootnote will create a new footnote and attach it to the Paragraph in the
// location at the end of the previous run (footnotes create their own run within
// the paragraph). The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the footnote later.
func (_cgccd Paragraph) AddFootnote(text string) Footnote {
	var _fgbfe int64
	if _cgccd._bbgc.HasFootnotes() {
		for _, _beg := range _cgccd._bbgc.Footnotes() {
			if _beg.id() > _fgbfe {
				_fgbfe = _beg.id()
			}
		}
		_fgbfe++
	} else {
		_fgbfe = 0
		_cgccd._bbgc._edf = &_eaba.Footnotes{}
		_cgccd._bbgc._edf.CT_Footnotes = _eaba.CT_Footnotes{}
		_cgccd._bbgc._edf.Footnote = make([]*_eaba.CT_FtnEdn, 0)
	}
	_dfbg := _eaba.NewCT_FtnEdn()
	_dbbg := _eaba.NewCT_FtnEdnRef()
	_dbbg.IdAttr = _fgbfe
	_cgccd._bbgc._edf.CT_Footnotes.Footnote = append(_cgccd._bbgc._edf.CT_Footnotes.Footnote, _dfbg)
	_bgad := _cgccd.AddRun()
	_aggdb := _bgad.Properties()
	_aggdb.SetStyle("\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065\u0041n\u0063\u0068\u006f\u0072")
	_bgad._bggd.EG_RunInnerContent = []*_eaba.EG_RunInnerContent{_eaba.NewEG_RunInnerContent()}
	_bgad._bggd.EG_RunInnerContent[0].FootnoteReference = _dbbg
	_eabg := Footnote{_cgccd._bbgc, _dfbg}
	_eabg._efba.IdAttr = _fgbfe
	_eabg._efba.EG_BlockLevelElts = []*_eaba.EG_BlockLevelElts{_eaba.NewEG_BlockLevelElts()}
	_abeaf := _eabg.AddParagraph()
	_abeaf.Properties().SetStyle("\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065")
	_abeaf._dbef.PPr.RPr = _eaba.NewCT_ParaRPr()
	_agef := _abeaf.AddRun()
	_agef.AddTab()
	_agef.AddText(text)
	return _eabg
}

var _fdeac = [...]uint8{0, 20, 37, 58, 79}

const (
	FormFieldTypeUnknown FormFieldType = iota
	FormFieldTypeText
	FormFieldTypeCheckBox
	FormFieldTypeDropDown
)

const (
	FieldCurrentPage   = "\u0050\u0041\u0047\u0045"
	FieldNumberOfPages = "\u004e\u0055\u004d\u0050\u0041\u0047\u0045\u0053"
	FieldDate          = "\u0044\u0041\u0054\u0045"
	FieldCreateDate    = "\u0043\u0052\u0045\u0041\u0054\u0045\u0044\u0041\u0054\u0045"
	FieldEditTime      = "\u0045\u0044\u0049\u0054\u0054\u0049\u004d\u0045"
	FieldPrintDate     = "\u0050R\u0049\u004e\u0054\u0044\u0041\u0054E"
	FieldSaveDate      = "\u0053\u0041\u0056\u0045\u0044\u0041\u0054\u0045"
	FieldTIme          = "\u0054\u0049\u004d\u0045"
	FieldTOC           = "\u0054\u004f\u0043"
)

// X return slice of node.
func (_fafe *Nodes) X() []Node { return _fafe._agcab }

// RightToLeft returns true if run text goes from right to left.
func (_aadg RunProperties) RightToLeft() bool { return _beff(_aadg._gfbgc.Rtl) }

// SetVAlignment sets the vertical alignment for an anchored image.
func (_ca AnchoredDrawing) SetVAlignment(v _eaba.WdST_AlignV) {
	_ca._cc.PositionV.Choice = &_eaba.WdCT_PosVChoice{}
	_ca._cc.PositionV.Choice.Align = v
}

// SetBottom sets the cell bottom margin
func (_bde CellMargins) SetBottom(d _eb.Distance) {
	_bde._ec.Bottom = _eaba.NewCT_TblWidth()
	_fgf(_bde._ec.Bottom, d)
}

// Paragraphs returns the paragraphs within a structured document tag.
func (_dagb StructuredDocumentTag) Paragraphs() []Paragraph {
	if _dagb._efaa.SdtContent == nil {
		return nil
	}
	_gaca := []Paragraph{}
	for _, _cdbge := range _dagb._efaa.SdtContent.P {
		_gaca = append(_gaca, Paragraph{_dagb._bbebc, _cdbge})
	}
	return _gaca
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (_ffec RunProperties) CharacterSpacingMeasure() string {
	if _bccbg := _ffec._gfbgc.Spacing; _bccbg != nil {
		_cfcc := _bccbg.ValAttr
		if _cfcc.ST_UniversalMeasure != nil {
			return *_cfcc.ST_UniversalMeasure
		}
	}
	return ""
}

// SetToolTip sets the tooltip text for a hyperlink.
func (_fcdab HyperLink) SetToolTip(text string) {
	if text == "" {
		_fcdab._dgace.TooltipAttr = nil
	} else {
		_fcdab._dgace.TooltipAttr = _d.String(text)
	}
}

// SetShapeStyle sets style to the element v:shape in watermark.
func (_abfcc *WatermarkPicture) SetShapeStyle(shapeStyle _ffc.ShapeStyle) {
	if _abfcc._ggac != nil {
		_gdcgb := shapeStyle.String()
		_abfcc._ggac.StyleAttr = &_gdcgb
	}
}

// SetOutlineLevel sets the outline level of this style.
func (_eaaef ParagraphStyleProperties) SetOutlineLevel(lvl int) {
	_eaaef._fded.OutlineLvl = _eaba.NewCT_DecimalNumber()
	_eaaef._fded.OutlineLvl.ValAttr = int64(lvl)
}

// SetFirstLineIndent controls the first line indent of the paragraph.
func (_gebf ParagraphStyleProperties) SetFirstLineIndent(m _eb.Distance) {
	if _gebf._fded.Ind == nil {
		_gebf._fded.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_gebf._fded.Ind.FirstLineAttr = nil
	} else {
		_gebf._fded.Ind.FirstLineAttr = &_cg.ST_TwipsMeasure{}
		_gebf._fded.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _d.Uint64(uint64(m / _eb.Twips))
	}
}

// ExtractText returns text from the document as a DocText object.
func (_edga *Document) ExtractText() *DocText {
	_aage := []TextItem{}
	for _, _gadd := range _edga._eeda.Body.EG_BlockLevelElts {
		_aage = append(_aage, _bddda(_gadd.EG_ContentBlockContent, nil)...)
	}
	var _gfdad []listItemInfo
	_daeb := _edga.Paragraphs()
	for _, _acef := range _daeb {
		_ddce := _bgdc(_edga, _acef)
		_gfdad = append(_gfdad, _ddce)
	}
	_fgae := _deeb(_edga)
	return &DocText{Items: _aage, _dbage: _gfdad, _bcgfa: _fgae}
}

// IsChecked returns true if a FormFieldTypeCheckBox is checked.
func (_acda FormField) IsChecked() bool {
	if _acda._cdag.CheckBox == nil {
		return false
	}
	if _acda._cdag.CheckBox.Checked != nil {
		return true
	}
	return false
}

func (_eefgg *WatermarkText) getShapeType() *_d.XSDAny {
	return _eefgg.getInnerElement("\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e")
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (_fgag TableStyleProperties) SetCellSpacingPercent(pct float64) {
	_fgag._ddafc.TblCellSpacing = _eaba.NewCT_TblWidth()
	_fgag._ddafc.TblCellSpacing.TypeAttr = _eaba.ST_TblWidthPct
	_fgag._ddafc.TblCellSpacing.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_fgag._ddafc.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_fgag._ddafc.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(pct * 50))
}

// Borders allows manipulation of the table borders.
func (_bafac TableProperties) Borders() TableBorders {
	if _bafac._efac.TblBorders == nil {
		_bafac._efac.TblBorders = _eaba.NewCT_TblBorders()
	}
	return TableBorders{_bafac._efac.TblBorders}
}

// SetPictureSize set watermark picture size with given width and height.
func (_aaee *WatermarkPicture) SetPictureSize(width, height int64) {
	if _aaee._ggac != nil {
		_face := _aaee.GetShapeStyle()
		_face.SetWidth(int64(width * _eb.Point))
		_face.SetHeight(int64(height * _eb.Point))
		_aaee.SetShapeStyle(_face)
	}
}

// SetTopPct sets the cell top margin
func (_aed CellMargins) SetTopPct(pct float64) {
	_aed._ec.Top = _eaba.NewCT_TblWidth()
	_cab(_aed._ec.Top, pct)
}

// SetNumberingLevel sets the numbering level of a paragraph.  If used, then the
// NumberingDefinition must also be set via SetNumberingDefinition or
// SetNumberingDefinitionByID.
func (_gfceb Paragraph) SetNumberingLevel(listLevel int) {
	_gfceb.ensurePPr()
	if _gfceb._dbef.PPr.NumPr == nil {
		_gfceb._dbef.PPr.NumPr = _eaba.NewCT_NumPr()
	}
	_eabbe := _eaba.NewCT_DecimalNumber()
	_eabbe.ValAttr = int64(listLevel)
	_gfceb._dbef.PPr.NumPr.Ilvl = _eabbe
}

func _dfb(_agda *_eaba.CT_P, _ggec *_eaba.CT_Hyperlink, _cacb *TableInfo, _gfcg *DrawingInfo, _abaf []*_eaba.EG_PContent) []TextItem {
	if len(_abaf) == 0 {
		return []TextItem{TextItem{Text: "", DrawingInfo: _gfcg, Paragraph: _agda, Hyperlink: _ggec, Run: nil, TableInfo: _cacb}}
	}
	_agea := []TextItem{}
	for _, _ddddf := range _abaf {
		for _, _cdfed := range _ddddf.FldSimple {
			if _cdfed != nil {
				_agea = append(_agea, _dfb(_agda, _ggec, _cacb, _gfcg, _cdfed.EG_PContent)...)
			}
		}
		if _dgggd := _ddddf.Hyperlink; _dgggd != nil {
			_agea = append(_agea, _gecg(_agda, _dgggd, _cacb, _gfcg, _dgggd.EG_ContentRunContent)...)
		}
		_agea = append(_agea, _gecg(_agda, nil, _cacb, _gfcg, _ddddf.EG_ContentRunContent)...)
	}
	return _agea
}

func (_dbcg *WatermarkText) getInnerElement(_ffbbgc string) *_d.XSDAny {
	for _, _ebba := range _dbcg._eefg.Any {
		_dfdg, _daggf := _ebba.(*_d.XSDAny)
		if _daggf && (_dfdg.XMLName.Local == _ffbbgc || _dfdg.XMLName.Local == "\u0076\u003a"+_ffbbgc) {
			return _dfdg
		}
	}
	return nil
}

// Paragraphs returns the paragraphs defined in a footer.
func (_eaeb Footer) Paragraphs() []Paragraph {
	_gacc := []Paragraph{}
	for _, _fbac := range _eaeb._daaa.EG_ContentBlockContent {
		for _, _agbcd := range _fbac.P {
			_gacc = append(_gacc, Paragraph{_eaeb._dec, _agbcd})
		}
	}
	for _, _fbffe := range _eaeb.Tables() {
		for _, _cffc := range _fbffe.Rows() {
			for _, _cabge := range _cffc.Cells() {
				_gacc = append(_gacc, _cabge.Paragraphs()...)
			}
		}
	}
	return _gacc
}

// AddFooter creates a Footer associated with the document, but doesn't add it
// to the document for display.
func (_fdc *Document) AddFooter() Footer {
	_cdg := _eaba.NewFtr()
	_fdc._gfdb = append(_fdc._gfdb, _cdg)
	_acd := _eg.Sprintf("\u0066\u006f\u006ft\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", len(_fdc._gfdb))
	_fdc._bdb.AddRelationship(_acd, _d.FooterType)
	_fdc.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_acd, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0066\u006f\u006f\u0074e\u0072\u002b\u0078\u006d\u006c")
	_fdc._efc = append(_fdc._efc, _ea.NewRelationships())
	return Footer{_fdc, _cdg}
}

// Pict returns the pict object.
func (_dgdf *WatermarkText) Pict() *_eaba.CT_Picture { return _dgdf._eefg }

func _dabf(_gbfa *Document, _adeeb []*_eaba.EG_ContentBlockContent, _bacec *TableInfo) []Node {
	_fdec := []Node{}
	for _, _acgff := range _adeeb {
		if _fadf := _acgff.Sdt; _fadf != nil {
			if _dgad := _fadf.SdtContent; _dgad != nil {
				_fdec = append(_fdec, _gbge(_gbfa, _dgad.P, _bacec, nil)...)
			}
		}
		_fdec = append(_fdec, _gbge(_gbfa, _acgff.P, _bacec, nil)...)
		for _, _bfce := range _acgff.Tbl {
			_decb := Table{_gbfa, _bfce}
			_cdagb, _ := _gbfa.Styles.SearchStyleById(_decb.Style())
			_fdafa := []Node{}
			for _ecfc, _bfefg := range _bfce.EG_ContentRowContent {
				for _, _aeaa := range _bfefg.Tr {
					for _gbgb, _aaaae := range _aeaa.EG_ContentCellContent {
						for _, _cbdc := range _aaaae.Tc {
							_aagg := &TableInfo{Table: _bfce, Row: _aeaa, Cell: _cbdc, RowIndex: _ecfc, ColIndex: _gbgb}
							for _, _faab := range _cbdc.EG_BlockLevelElts {
								_fdafa = append(_fdafa, _dabf(_gbfa, _faab.EG_ContentBlockContent, _aagg)...)
							}
						}
					}
				}
			}
			_fdec = append(_fdec, Node{_abcbe: _gbfa, _efag: &_decb, Style: _cdagb, Children: _fdafa})
		}
	}
	return _fdec
}

// X returns the inner wrapped XML type.
func (_fde Bookmark) X() *_eaba.CT_Bookmark { return _fde._dda }

// AddDefinition adds a new numbering definition.
func (_cfce Numbering) AddDefinition() NumberingDefinition {
	_bdcgd := _eaba.NewCT_Num()
	_addeg := int64(1)
	for _, _ceee := range _cfce.Definitions() {
		if _ceee.AbstractNumberID() >= _addeg {
			_addeg = _ceee.AbstractNumberID() + 1
		}
	}
	_fdfc := int64(1)
	for _, _gdeb := range _cfce.X().Num {
		if _gdeb.NumIdAttr >= _fdfc {
			_fdfc = _gdeb.NumIdAttr + 1
		}
	}
	_bdcgd.NumIdAttr = _fdfc
	_bdcgd.AbstractNumId = _eaba.NewCT_DecimalNumber()
	_bdcgd.AbstractNumId.ValAttr = _addeg
	_edfbc := _eaba.NewCT_AbstractNum()
	_edfbc.AbstractNumIdAttr = _addeg
	_cfce._ebbdb.AbstractNum = append(_cfce._ebbdb.AbstractNum, _edfbc)
	_cfce._ebbdb.Num = append(_cfce._ebbdb.Num, _bdcgd)
	return NumberingDefinition{_edfbc}
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_dbdd ParagraphProperties) ComplexSizeMeasure() string {
	if _cedc := _dbdd._efec.RPr.SzCs; _cedc != nil {
		_ecdbc := _cedc.ValAttr
		if _ecdbc.ST_PositiveUniversalMeasure != nil {
			return *_ecdbc.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// RunProperties returns the run properties controlling text formatting within the table.
func (_cfbde TableConditionalFormatting) RunProperties() RunProperties {
	if _cfbde._agcba.RPr == nil {
		_cfbde._agcba.RPr = _eaba.NewCT_RPr()
	}
	return RunProperties{_cfbde._agcba.RPr}
}

// Properties returns the table properties.
func (_eeca Table) Properties() TableProperties {
	if _eeca._fecae.TblPr == nil {
		_eeca._fecae.TblPr = _eaba.NewCT_TblPr()
	}
	return TableProperties{_eeca._fecae.TblPr}
}

// Outline returns true if run outline is on.
func (_dafe RunProperties) Outline() bool { return _beff(_dafe._gfbgc.Outline) }

// SetValue sets the width value.
func (_fdfg TableWidth) SetValue(m _eb.Distance) {
	_fdfg._ebga.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_fdfg._ebga.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_fdfg._ebga.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(m / _eb.Twips))
	_fdfg._ebga.TypeAttr = _eaba.ST_TblWidthDxa
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_eea CellBorders) SetLeft(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_eea._ggf.Left = _eaba.NewCT_Border()
	_gfdf(_eea._ggf.Left, t, c, thickness)
}

// Caps returns true if paragraph font is capitalized.
func (_ccdd ParagraphProperties) Caps() bool { return _beff(_ccdd._efec.RPr.Caps) }

// Italic returns true if paragraph font is italic.
func (_bcggb ParagraphProperties) Italic() bool {
	_bgac := _bcggb._efec.RPr
	return _beff(_bgac.I) || _beff(_bgac.ICs)
}

// Style return the table style.
func (_cbgc Table) Style() string {
	if _cbgc._fecae.TblPr != nil && _cbgc._fecae.TblPr.TblStyle != nil {
		return _cbgc._fecae.TblPr.TblStyle.ValAttr
	}
	return ""
}

func _eaaec(_ebfe *_eaba.CT_Tbl, _ccef map[string]string) {
	for _, _gbfcd := range _ebfe.EG_ContentRowContent {
		for _, _cdce := range _gbfcd.Tr {
			for _, _fcef := range _cdce.EG_ContentCellContent {
				for _, _gecb := range _fcef.Tc {
					for _, _cggdc := range _gecb.EG_BlockLevelElts {
						for _, _daca := range _cggdc.EG_ContentBlockContent {
							for _, _gdce := range _daca.P {
								_eac(_gdce, _ccef)
							}
							for _, _gfddd := range _daca.Tbl {
								_eaaec(_gfddd, _ccef)
							}
						}
					}
				}
			}
		}
	}
}

// AddParagraph adds a paragraph to the header.
func (_ggbb Header) AddParagraph() Paragraph {
	_cfbc := _eaba.NewEG_ContentBlockContent()
	_ggbb._cbgbf.EG_ContentBlockContent = append(_ggbb._cbgbf.EG_ContentBlockContent, _cfbc)
	_gffa := _eaba.NewCT_P()
	_cfbc.P = append(_cfbc.P, _gffa)
	return Paragraph{_ggbb._bffa, _gffa}
}

// NewTableWidth returns a newly intialized TableWidth
func NewTableWidth() TableWidth { return TableWidth{_eaba.NewCT_TblWidth()} }

// Section is the beginning of a new section.
type Section struct {
	_ceec  *Document
	_aggca *_eaba.CT_SectPr
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (_bfecb RunProperties) RStyle() string {
	if _bfecb._gfbgc.RStyle != nil {
		return _bfecb._gfbgc.RStyle.ValAttr
	}
	return ""
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_cbeff TableStyleProperties) SetCellSpacingAuto() {
	_cbeff._ddafc.TblCellSpacing = _eaba.NewCT_TblWidth()
	_cbeff._ddafc.TblCellSpacing.TypeAttr = _eaba.ST_TblWidthAuto
}

const (
	OnOffValueUnset OnOffValue = iota
	OnOffValueOff
	OnOffValueOn
)

func (_caef *Document) putNode(_ccgb, _bcbd Node, _afb bool) bool {
	_caef.insertImageFromNode(_bcbd)
	_caef.insertStyleFromNode(_bcbd)
	switch _cdgg := _bcbd._efag.(type) {
	case *Paragraph:
		if _bac, _dfge := _ccgb.X().(*Paragraph); _dfge {
			_caef.appendParagraph(_bac, *_cdgg, _afb)
			return true
		} else {
			for _, _eegbg := range _ccgb.Children {
				if _afbb := _caef.putNode(_eegbg, _bcbd, _afb); _afbb {
					break
				}
			}
		}
	case *Table:
		if _edbcb, _gbcgg := _ccgb.X().(*Paragraph); _gbcgg {
			_ccfdb := _caef.appendTable(_edbcb, *_cdgg, _afb)
			_ccfdb._fecae = _cdgg._fecae
			return true
		} else {
			for _, _bfff := range _ccgb.Children {
				if _bfdf := _caef.putNode(_bfff, _bcbd, _afb); _bfdf {
					break
				}
			}
		}
	}
	return false
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (_fdbbc Paragraph) SetLineSpacing(d _eb.Distance, rule _eaba.ST_LineSpacingRule) {
	_fdbbc.ensurePPr()
	if _fdbbc._dbef.PPr.Spacing == nil {
		_fdbbc._dbef.PPr.Spacing = _eaba.NewCT_Spacing()
	}
	_effef := _fdbbc._dbef.PPr.Spacing
	if rule == _eaba.ST_LineSpacingRuleUnset {
		_effef.LineRuleAttr = _eaba.ST_LineSpacingRuleUnset
		_effef.LineAttr = nil
	} else {
		_effef.LineRuleAttr = rule
		_effef.LineAttr = &_eaba.ST_SignedTwipsMeasure{}
		_effef.LineAttr.Int64 = _d.Int64(int64(d / _eb.Twips))
	}
}

func _ebb(_ebff *_eaba.EG_ContentBlockContent) []Bookmark {
	_cgfc := []Bookmark{}
	for _, _adgg := range _ebff.P {
		for _, _gfdbe := range _adgg.EG_PContent {
			for _, _aefc := range _gfdbe.EG_ContentRunContent {
				for _, _gccbg := range _aefc.EG_RunLevelElts {
					for _, _cagc := range _gccbg.EG_RangeMarkupElements {
						if _cagc.BookmarkStart != nil {
							_cgfc = append(_cgfc, Bookmark{_cagc.BookmarkStart})
						}
					}
				}
			}
		}
	}
	for _, _baae := range _ebff.EG_RunLevelElts {
		for _, _eafd := range _baae.EG_RangeMarkupElements {
			if _eafd.BookmarkStart != nil {
				_cgfc = append(_cgfc, Bookmark{_eafd.BookmarkStart})
			}
		}
	}
	for _, _adag := range _ebff.Tbl {
		for _, _gbfc := range _adag.EG_ContentRowContent {
			for _, _fdcc := range _gbfc.Tr {
				for _, _ceda := range _fdcc.EG_ContentCellContent {
					for _, _cdaf := range _ceda.Tc {
						for _, _egdec := range _cdaf.EG_BlockLevelElts {
							for _, _cdae := range _egdec.EG_ContentBlockContent {
								for _, _fbgf := range _ebb(_cdae) {
									_cgfc = append(_cgfc, _fbgf)
								}
							}
						}
					}
				}
			}
		}
	}
	return _cgfc
}

// SetUISortOrder controls the order the style is displayed in the UI.
func (_ggcdd Style) SetUISortOrder(order int) {
	_ggcdd._bgfca.UiPriority = _eaba.NewCT_DecimalNumber()
	_ggcdd._bgfca.UiPriority.ValAttr = int64(order)
}

// DrawingInfo is used for keep information about a drawing wrapping a textbox where the text is located.
type DrawingInfo struct {
	Drawing *_eaba.CT_Drawing
	Width   int64
	Height  int64
}

// Numbering is the document wide numbering styles contained in numbering.xml.
type Numbering struct{ _ebbdb *_eaba.Numbering }

func _adaaa() *_aae.Path {
	_bgea := _aae.NewPath()
	_bgea.ExtrusionokAttr = _cg.ST_TrueFalseTrue
	_bgea.GradientshapeokAttr = _cg.ST_TrueFalseTrue
	_bgea.ConnecttypeAttr = _aae.OfcST_ConnectTypeRect
	return _bgea
}

func (_dbe *chart) RelId() string { return _dbe._ece }

// Margins allows controlling individual cell margins.
func (_ffgd CellProperties) Margins() CellMargins {
	if _ffgd._be.TcMar == nil {
		_ffgd._be.TcMar = _eaba.NewCT_TcMar()
	}
	return CellMargins{_ffgd._be.TcMar}
}

// X returns the inner wrapped XML type.
func (_fdfd RunProperties) X() *_eaba.CT_RPr { return _fdfd._gfbgc }

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (_cffac TableBorders) SetInsideVertical(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_cffac._dbbcbf.InsideV = _eaba.NewCT_Border()
	_gfdf(_cffac._dbbcbf.InsideV, t, c, thickness)
}

// SetText sets the watermark text.
func (_acaee *WatermarkText) SetText(text string) {
	_adgaf := _acaee.getShape()
	if _acaee._ggfda != nil {
		_beddg := _acaee._ggfda.EG_ShapeElements
		if len(_beddg) > 0 && _beddg[0].Textpath != nil {
			_beddg[0].Textpath.StringAttr = &text
		}
	} else {
		_edeaf := _acaee.findNode(_adgaf, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
		for _dadd, _gadg := range _edeaf.Attrs {
			if _gadg.Name.Local == "\u0073\u0074\u0072\u0069\u006e\u0067" {
				_edeaf.Attrs[_dadd].Value = text
			}
		}
	}
}

// InlineDrawing is an inlined image within a run.
type InlineDrawing struct {
	_cgdg *Document
	_fbbd *_eaba.WdInline
}

func _adegd() *_aae.Formulas {
	_fagfc := _aae.NewFormulas()
	_fagfc.F = []*_aae.CT_F{_ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0023\u0030\u0020\u0030\u00201\u0030\u0038\u0030\u0030"), _ffc.CreateFormula("p\u0072\u006f\u0064\u0020\u0023\u0030\u0020\u0032\u0020\u0031"), _ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0032\u0031\u0036\u0030\u0030 \u0030\u0020\u0040\u0031"), _ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0030\u0020\u0030\u0020\u0040\u0032"), _ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0032\u0031\u0036\u0030\u0030 \u0030\u0020\u0040\u0033"), _ffc.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0040\u0033\u0020\u0030"), _ffc.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0032\u0031\u00360\u0030\u0020\u0040\u0031"), _ffc.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0030\u0020\u0040\u0032"), _ffc.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0040\u0034\u00202\u0031\u0036\u0030\u0030"), _ffc.CreateFormula("\u006di\u0064\u0020\u0040\u0035\u0020\u00406"), _ffc.CreateFormula("\u006di\u0064\u0020\u0040\u0038\u0020\u00405"), _ffc.CreateFormula("\u006di\u0064\u0020\u0040\u0037\u0020\u00408"), _ffc.CreateFormula("\u006di\u0064\u0020\u0040\u0036\u0020\u00407"), _ffc.CreateFormula("s\u0075\u006d\u0020\u0040\u0036\u0020\u0030\u0020\u0040\u0035")}
	return _fagfc
}

// SetStyle sets the font size.
func (_ffeb RunProperties) SetStyle(style string) {
	if style == "" {
		_ffeb._gfbgc.RStyle = nil
	} else {
		_ffeb._gfbgc.RStyle = _eaba.NewCT_String()
		_ffeb._gfbgc.RStyle.ValAttr = style
	}
}

// StructuredDocumentTag are a tagged bit of content in a document.
type StructuredDocumentTag struct {
	_bbebc *Document
	_efaa  *_eaba.CT_SdtBlock
}

func _edfg() *_aae.Formulas {
	_gaad := _aae.NewFormulas()
	_gaad.F = []*_aae.CT_F{_ffc.CreateFormula("\u0069\u0066 \u006c\u0069\u006e\u0065\u0044\u0072\u0061\u0077\u006e\u0020\u0070\u0069\u0078\u0065\u006c\u004c\u0069\u006e\u0065\u0057\u0069\u0064th\u0020\u0030"), _ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0040\u0030\u0020\u0031\u0020\u0030"), _ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0030\u0020\u0030\u0020\u0040\u0031"), _ffc.CreateFormula("p\u0072\u006f\u0064\u0020\u0040\u0032\u0020\u0031\u0020\u0032"), _ffc.CreateFormula("\u0070r\u006f\u0064\u0020\u0040\u0033\u0020\u0032\u0031\u0036\u0030\u0030 \u0070\u0069\u0078\u0065\u006c\u0057\u0069\u0064\u0074\u0068"), _ffc.CreateFormula("\u0070r\u006f\u0064\u0020\u00403\u0020\u0032\u0031\u0036\u00300\u0020p\u0069x\u0065\u006c\u0048\u0065\u0069\u0067\u0068t"), _ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0040\u0030\u0020\u0030\u0020\u0031"), _ffc.CreateFormula("p\u0072\u006f\u0064\u0020\u0040\u0036\u0020\u0031\u0020\u0032"), _ffc.CreateFormula("\u0070r\u006f\u0064\u0020\u0040\u0037\u0020\u0032\u0031\u0036\u0030\u0030 \u0070\u0069\u0078\u0065\u006c\u0057\u0069\u0064\u0074\u0068"), _ffc.CreateFormula("\u0073\u0075\u006d\u0020\u0040\u0038\u0020\u0032\u00316\u0030\u0030\u0020\u0030"), _ffc.CreateFormula("\u0070r\u006f\u0064\u0020\u00407\u0020\u0032\u0031\u0036\u00300\u0020p\u0069x\u0065\u006c\u0048\u0065\u0069\u0067\u0068t"), _ffc.CreateFormula("\u0073u\u006d \u0040\u0031\u0030\u0020\u0032\u0031\u0036\u0030\u0030\u0020\u0030")}
	return _gaad
}

// InsertRunBefore inserts a run in the paragraph before the relative run.
func (_gcga Paragraph) InsertRunBefore(relativeTo Run) Run {
	return _gcga.insertRun(relativeTo, true)
}

// EastAsiaFont returns the name of paragraph font family for East Asia.
func (_eecb ParagraphProperties) EastAsiaFont() string {
	if _dabbe := _eecb._efec.RPr.RFonts; _dabbe != nil {
		if _dabbe.EastAsiaAttr != nil {
			return *_dabbe.EastAsiaAttr
		}
	}
	return ""
}

// EastAsiaFont returns the name of run font family for East Asia.
func (_degb RunProperties) EastAsiaFont() string {
	if _bccef := _degb._gfbgc.RFonts; _bccef != nil {
		if _bccef.EastAsiaAttr != nil {
			return *_bccef.EastAsiaAttr
		}
	}
	return ""
}

// InsertRunAfter inserts a run in the paragraph after the relative run.
func (_gbde Paragraph) InsertRunAfter(relativeTo Run) Run {
	return _gbde.insertRun(relativeTo, false)
}

// ReplaceText replace text inside node.
func (_fbfe *Nodes) ReplaceText(oldText, newText string) {
	for _, _fafef := range _fbfe._agcab {
		_fafef.ReplaceText(oldText, newText)
	}
}

// Font returns the name of paragraph font family.
func (_agdg ParagraphProperties) Font() string {
	if _aaaba := _agdg._efec.RPr.RFonts; _aaaba != nil {
		if _aaaba.AsciiAttr != nil {
			return *_aaaba.AsciiAttr
		} else if _aaaba.HAnsiAttr != nil {
			return *_aaaba.HAnsiAttr
		} else if _aaaba.CsAttr != nil {
			return *_aaaba.CsAttr
		}
	}
	return ""
}

// RowProperties are the properties for a row within a table
type RowProperties struct{ _bccec *_eaba.CT_TrPr }

// X returns the inner wrapped XML type.
func (_gaead TableProperties) X() *_eaba.CT_TblPr { return _gaead._efac }

func _deeb(_bgcb *Document) map[int64]map[int64]int64 {
	_ecga := _bgcb.Paragraphs()
	_ebab := make(map[int64]map[int64]int64, 0)
	for _, _bgcd := range _ecga {
		_aeff := _bgdc(_bgcb, _bgcd)
		if _aeff.NumberingLevel != nil && _aeff.AbstractNumId != nil {
			_cced := *_aeff.AbstractNumId
			if _, _ageb := _ebab[_cced]; _ageb {
				if _, _ebfd := _ebab[_cced][_aeff.NumberingLevel.X().IlvlAttr]; _ebfd {
					_ebab[_cced][_aeff.NumberingLevel.X().IlvlAttr]++
				} else {
					_ebab[_cced][_aeff.NumberingLevel.X().IlvlAttr] = 1
				}
			} else {
				_ebab[_cced] = map[int64]int64{_aeff.NumberingLevel.X().IlvlAttr: 1}
			}
		}
	}
	return _ebab
}

func (_egcd *WatermarkText) findNode(_gedf *_d.XSDAny, _ecdg string) *_d.XSDAny {
	for _, _ddeb := range _gedf.Nodes {
		if _ddeb.XMLName.Local == _ecdg {
			return _ddeb
		}
	}
	return nil
}

func (_ffaf Run) newIC() *_eaba.EG_RunInnerContent {
	_gcec := _eaba.NewEG_RunInnerContent()
	_ffaf._bggd.EG_RunInnerContent = append(_ffaf._bggd.EG_RunInnerContent, _gcec)
	return _gcec
}

// SetFirstColumn controls the conditional formatting for the first column in a table.
func (_gbba TableLook) SetFirstColumn(on bool) {
	if !on {
		_gbba._ffdgf.FirstColumnAttr = &_cg.ST_OnOff{}
		_gbba._ffdgf.FirstColumnAttr.ST_OnOff1 = _cg.ST_OnOff1Off
	} else {
		_gbba._ffdgf.FirstColumnAttr = &_cg.ST_OnOff{}
		_gbba._ffdgf.FirstColumnAttr.ST_OnOff1 = _cg.ST_OnOff1On
	}
}

// TableConditionalFormatting returns a conditional formatting object of a given
// type.  Calling this method repeatedly will return the same object.
func (_cbda Style) TableConditionalFormatting(typ _eaba.ST_TblStyleOverrideType) TableConditionalFormatting {
	for _, _fbedc := range _cbda._bgfca.TblStylePr {
		if _fbedc.TypeAttr == typ {
			return TableConditionalFormatting{_fbedc}
		}
	}
	_bfcea := _eaba.NewCT_TblStylePr()
	_bfcea.TypeAttr = typ
	_cbda._bgfca.TblStylePr = append(_cbda._bgfca.TblStylePr, _bfcea)
	return TableConditionalFormatting{_bfcea}
}

// FindNodeByStyleName return slice of node base on style name.
func (_egceb *Nodes) FindNodeByStyleName(styleName string) []Node {
	_efbf := []Node{}
	for _, _gbaga := range _egceb._agcab {
		switch _cabce := _gbaga._efag.(type) {
		case *Paragraph:
			if _cabce != nil {
				if _adce, _ccgg := _gbaga._abcbe.Styles.SearchStyleByName(styleName); _ccgg {
					_bffeb := _cabce.Style()
					if _bffeb == _adce.StyleID() {
						_efbf = append(_efbf, _gbaga)
					}
				}
			}
		case *Table:
			if _cabce != nil {
				if _dafc, _eagbg := _gbaga._abcbe.Styles.SearchStyleByName(styleName); _eagbg {
					_dbbee := _cabce.Style()
					if _dbbee == _dafc.StyleID() {
						_efbf = append(_efbf, _gbaga)
					}
				}
			}
		}
		_afbc := Nodes{_agcab: _gbaga.Children}
		_efbf = append(_efbf, _afbc.FindNodeByStyleName(styleName)...)
	}
	return _efbf
}

// SetKeepNext controls if the paragraph is kept with the next paragraph.
func (_gccab ParagraphStyleProperties) SetKeepNext(b bool) {
	if !b {
		_gccab._fded.KeepNext = nil
	} else {
		_gccab._fded.KeepNext = _eaba.NewCT_OnOff()
	}
}

// Name returns the name of the style if set.
func (_gdbg Style) Name() string {
	if _gdbg._bgfca.Name == nil {
		return ""
	}
	return _gdbg._bgfca.Name.ValAttr
}

// SetAlignment positions an anchored image via alignment.  Offset is
// incompatible with SetOffset, whichever is called last is applied.
func (_dac AnchoredDrawing) SetAlignment(h _eaba.WdST_AlignH, v _eaba.WdST_AlignV) {
	_dac.SetHAlignment(h)
	_dac.SetVAlignment(v)
}

// SetStartIndent controls the start indentation.
func (_fgab ParagraphProperties) SetStartIndent(m _eb.Distance) {
	if _fgab._efec.Ind == nil {
		_fgab._efec.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_fgab._efec.Ind.StartAttr = nil
	} else {
		_fgab._efec.Ind.StartAttr = &_eaba.ST_SignedTwipsMeasure{}
		_fgab._efec.Ind.StartAttr.Int64 = _d.Int64(int64(m / _eb.Twips))
	}
}

// X returns the inner wrapped XML type.
func (_bee NumberingLevel) X() *_eaba.CT_Lvl { return _bee._fgbfa }

// SetConformance sets conformance attribute of the document
// as one of these values from gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (_ceeb Document) SetConformance(conformanceAttr _cg.ST_ConformanceClass) {
	_ceeb._eeda.ConformanceAttr = conformanceAttr
}

func _ecae(_dbbe *_bb.CT_Blip, _gbdd map[string]string) {
	if _dbbe.EmbedAttr != nil {
		if _dccac, _bgbb := _gbdd[*_dbbe.EmbedAttr]; _bgbb {
			*_dbbe.EmbedAttr = _dccac
		}
	}
}

// ItalicValue returns the precise nature of the italic setting (unset, off or on).
func (_fbfec RunProperties) ItalicValue() OnOffValue { return _efcgda(_fbfec._gfbgc.I) }

func _fcgde() *_aae.Imagedata {
	_gcebf := _aae.NewImagedata()
	_agbafb := "\u0072\u0049\u0064\u0031"
	_eabdb := "\u0057A\u0054\u0045\u0052\u004d\u0041\u0052K"
	_gcebf.IdAttr = &_agbafb
	_gcebf.TitleAttr = &_eabdb
	return _gcebf
}

// OnOffValue represents an on/off value that can also be unset
type OnOffValue byte

// SetFontFamily sets the Ascii & HAnsi fonly family for a run.
func (_cfbce RunProperties) SetFontFamily(family string) {
	if _cfbce._gfbgc.RFonts == nil {
		_cfbce._gfbgc.RFonts = _eaba.NewCT_Fonts()
	}
	_cfbce._gfbgc.RFonts.AsciiAttr = _d.String(family)
	_cfbce._gfbgc.RFonts.HAnsiAttr = _d.String(family)
	_cfbce._gfbgc.RFonts.EastAsiaAttr = _d.String(family)
}

// SaveToFile writes the document out to a file.
func (_aaeg *Document) SaveToFile(path string) error {
	_dfdc, _eag := _ff.Create(path)
	if _eag != nil {
		return _eag
	}
	defer _dfdc.Close()
	return _aaeg.Save(_dfdc)
}

// GetColor returns the color.Color object representing the run color.
func (_fabea ParagraphProperties) GetColor() _bc.Color {
	if _adggd := _fabea._efec.RPr.Color; _adggd != nil {
		_fgbd := _adggd.ValAttr
		if _fgbd.ST_HexColorRGB != nil {
			return _bc.FromHex(*_fgbd.ST_HexColorRGB)
		}
	}
	return _bc.Color{}
}

// ParagraphStyles returns only the paragraph styles.
func (_fcdfd Styles) ParagraphStyles() []Style {
	_acbac := []Style{}
	for _, _eadb := range _fcdfd._abad.Style {
		if _eadb.TypeAttr != _eaba.ST_StyleTypeParagraph {
			continue
		}
		_acbac = append(_acbac, Style{_eadb})
	}
	return _acbac
}

// SetStyle sets the style of a paragraph and is identical to setting it on the
// paragraph's Properties()
func (_bebf Paragraph) SetStyle(s string) {
	_bebf.ensurePPr()
	if s == "" {
		_bebf._dbef.PPr.PStyle = nil
	} else {
		_bebf._dbef.PPr.PStyle = _eaba.NewCT_String()
		_bebf._dbef.PPr.PStyle.ValAttr = s
	}
}

// SetName marks sets a name attribute for a FormField.
func (_ageg FormField) SetName(name string) {
	_cgddf := _eaba.NewCT_FFName()
	_cgddf.ValAttr = &name
	_ageg._cdag.Name = []*_eaba.CT_FFName{_cgddf}
}

func (_egab Paragraph) addSeparateFldChar() *_eaba.CT_FldChar {
	_abdde := _egab.addFldChar()
	_abdde.FldCharTypeAttr = _eaba.ST_FldCharTypeSeparate
	return _abdde
}

// Bold returns true if paragraph font is bold.
func (_ffdc ParagraphProperties) Bold() bool {
	_effee := _ffdc._efec.RPr
	return _beff(_effee.B) || _beff(_effee.BCs)
}

// SetPrimaryStyle marks the style as a primary style.
func (_efed Style) SetPrimaryStyle(b bool) {
	if b {
		_efed._bgfca.QFormat = _eaba.NewCT_OnOff()
	} else {
		_efed._bgfca.QFormat = nil
	}
}

// SetWidthPercent sets the table to a width percentage.
func (_accg TableProperties) SetWidthPercent(pct float64) {
	_accg._efac.TblW = _eaba.NewCT_TblWidth()
	_accg._efac.TblW.TypeAttr = _eaba.ST_TblWidthPct
	_accg._efac.TblW.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_accg._efac.TblW.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_accg._efac.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(pct * 50))
}

// ParagraphSpacing controls the spacing for a paragraph and its lines.
type ParagraphSpacing struct{ _ebfge *_eaba.CT_Spacing }

// NumId return numbering numId that being use by style properties.
func (_eccfga ParagraphStyleProperties) NumId() int64 {
	if _eccfga._fded.NumPr != nil {
		if _eccfga._fded.NumPr.NumId != nil {
			return _eccfga._fded.NumPr.NumId.ValAttr
		}
	}
	return -1
}

func _bg() (*_bb.CT_Point2D, []*_bb.CT_Point2D) {
	var (
		_bd int64 = 0
		_gg int64 = 21600
	)
	_bgg := _bb.ST_Coordinate{ST_CoordinateUnqualified: &_bd, ST_UniversalMeasure: nil}
	_edg := _bb.ST_Coordinate{ST_CoordinateUnqualified: &_gg, ST_UniversalMeasure: nil}
	_ef := _bb.NewCT_Point2D()
	_ef.XAttr = _bgg
	_ef.YAttr = _bgg
	_cbc := []*_bb.CT_Point2D{&_bb.CT_Point2D{XAttr: _bgg, YAttr: _edg}, &_bb.CT_Point2D{XAttr: _edg, YAttr: _edg}, &_bb.CT_Point2D{XAttr: _edg, YAttr: _bgg}, _ef}
	return _ef, _cbc
}

func (_fggb *Document) onNewRelationship(_effc *_fd.DecodeMap, _agd, _dacb string, _efbge []*_da.File, _eefb *_eab.Relationship, _bagf _fd.Target) error {
	_aecb := _d.DocTypeDocument
	switch _dacb {
	case _d.OfficeDocumentType, _d.OfficeDocumentTypeStrict:
		_fggb._eeda = _eaba.NewDocument()
		_effc.AddTarget(_agd, _fggb._eeda, _dacb, 0)
		_effc.AddTarget(_fd.RelationsPathFor(_agd), _fggb._bdb.X(), _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.CorePropertiesType:
		_effc.AddTarget(_agd, _fggb.CoreProperties.X(), _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.CustomPropertiesType:
		_effc.AddTarget(_agd, _fggb.CustomProperties.X(), _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.ExtendedPropertiesType, _d.ExtendedPropertiesTypeStrict:
		_effc.AddTarget(_agd, _fggb.AppProperties.X(), _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.ThumbnailType, _d.ThumbnailTypeStrict:
		for _cbe, _dcfc := range _efbge {
			if _dcfc == nil {
				continue
			}
			if _dcfc.Name == _agd {
				_bfbc, _eaab := _dcfc.Open()
				if _eaab != nil {
					return _eg.Errorf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073", _eaab)
				}
				_fggb.Thumbnail, _, _eaab = _ac.Decode(_bfbc)
				_bfbc.Close()
				if _eaab != nil {
					return _eg.Errorf("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073", _eaab)
				}
				_efbge[_cbe] = nil
			}
		}
	case _d.SettingsType, _d.SettingsTypeStrict:
		_effc.AddTarget(_agd, _fggb.Settings.X(), _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.NumberingType, _d.NumberingTypeStrict:
		_fggb.Numbering = NewNumbering()
		_effc.AddTarget(_agd, _fggb.Numbering.X(), _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.StylesType, _d.StylesTypeStrict:
		_fggb.Styles.Clear()
		_effc.AddTarget(_agd, _fggb.Styles.X(), _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.HeaderType, _d.HeaderTypeStrict:
		_bfbe := _eaba.NewHdr()
		_effc.AddTarget(_agd, _bfbe, _dacb, uint32(len(_fggb._bbac)))
		_fggb._bbac = append(_fggb._bbac, _bfbe)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, len(_fggb._bbac))
		_ccgf := _ea.NewRelationships()
		_effc.AddTarget(_fd.RelationsPathFor(_agd), _ccgf.X(), _dacb, 0)
		_fggb._cfg = append(_fggb._cfg, _ccgf)
	case _d.FooterType, _d.FooterTypeStrict:
		_afgg := _eaba.NewFtr()
		_effc.AddTarget(_agd, _afgg, _dacb, uint32(len(_fggb._gfdb)))
		_fggb._gfdb = append(_fggb._gfdb, _afgg)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, len(_fggb._gfdb))
		_cgbc := _ea.NewRelationships()
		_effc.AddTarget(_fd.RelationsPathFor(_agd), _cgbc.X(), _dacb, 0)
		_fggb._efc = append(_fggb._efc, _cgbc)
	case _d.ThemeType, _d.ThemeTypeStrict:
		_aeb := _bb.NewTheme()
		_effc.AddTarget(_agd, _aeb, _dacb, uint32(len(_fggb._bag)))
		_fggb._bag = append(_fggb._bag, _aeb)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, len(_fggb._bag))
	case _d.WebSettingsType, _d.WebSettingsTypeStrict:
		_fggb._feb = _eaba.NewWebSettings()
		_effc.AddTarget(_agd, _fggb._feb, _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.FontTableType, _d.FontTableTypeStrict:
		_fggb._fdd = _eaba.NewFonts()
		_effc.AddTarget(_agd, _fggb._fdd, _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.EndNotesType, _d.EndNotesTypeStrict:
		_fggb._bfd = _eaba.NewEndnotes()
		_effc.AddTarget(_agd, _fggb._bfd, _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.FootNotesType, _d.FootNotesTypeStrict:
		_fggb._edf = _eaba.NewFootnotes()
		_effc.AddTarget(_agd, _fggb._edf, _dacb, 0)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, 0)
	case _d.ImageType, _d.ImageTypeStrict:
		var _bbg _ea.ImageRef
		for _cfgd, _adg := range _efbge {
			if _adg == nil {
				continue
			}
			if _adg.Name == _agd {
				_efea, _agdc := _fd.ExtractToDiskTmp(_adg, _fggb.TmpPath)
				if _agdc != nil {
					return _agdc
				}
				_ffad, _agdc := _ea.ImageFromStorage(_efea)
				if _agdc != nil {
					return _agdc
				}
				_bbg = _ea.MakeImageRef(_ffad, &_fggb.DocBase, _fggb._bdb)
				_efbge[_cfgd] = nil
			}
		}
		if _bbg.Format() != "" {
			_agcg := "\u002e" + _f.ToLower(_bbg.Format())
			_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, len(_fggb.Images)+1)
			if _aga := _c.Ext(_eefb.TargetAttr); _aga != _agcg {
				_eefb.TargetAttr = _eefb.TargetAttr[0:len(_eefb.TargetAttr)-len(_aga)] + _agcg
			}
			_bbg.SetTarget("\u0077\u006f\u0072d\u002f" + _eefb.TargetAttr)
			_fggb.Images = append(_fggb.Images, _bbg)
		}
	case _d.ControlType, _d.ControlTypeStrict:
		_acba := _gf.NewOcx()
		_effc.AddTarget(_agd, _acba, _dacb, uint32(len(_fggb._aac)))
		_fggb._aac = append(_fggb._aac, _acba)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, len(_fggb._aac))
	case _d.ChartType:
		_bcafb := chart{_cebb: _dg.NewChartSpace()}
		_ddba := uint32(len(_fggb._cea))
		_effc.AddTarget(_agd, _bcafb._cebb, _dacb, _ddba)
		_fggb._cea = append(_fggb._cea, &_bcafb)
		_eefb.TargetAttr = _d.RelativeFilename(_aecb, _bagf.Typ, _dacb, len(_fggb._cea))
		_bcafb._fea = _eefb.TargetAttr
	default:
		_dcd.Log.Debug("\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u0074\u0067\u0074\u003a\u0020\u0025\u0073", _dacb, _agd)
	}
	return nil
}

// Run is a run of text within a paragraph that shares the same formatting.
type Run struct {
	_dddbb *Document
	_bggd  *_eaba.CT_R
}

// X returns the inner wrapped XML type.
func (_daceg ParagraphProperties) X() *_eaba.CT_PPr { return _daceg._efec }

// BoldValue returns the precise nature of the bold setting (unset, off or on).
func (_gdgab RunProperties) BoldValue() OnOffValue { return _efcgda(_gdgab._gfbgc.B) }

// SetLeft sets the cell left margin
func (_dddf CellMargins) SetLeft(d _eb.Distance) {
	_dddf._ec.Left = _eaba.NewCT_TblWidth()
	_fgf(_dddf._ec.Left, d)
}

// Footnote is an individual footnote reference within the document.
type Footnote struct {
	_ffdf *Document
	_efba *_eaba.CT_FtnEdn
}

// SetUpdateFieldsOnOpen controls if fields are recalculated upon opening the
// document. This is useful for things like a table of contents as the library
// only adds the field code and relies on Word/LibreOffice to actually compute
// the content.
func (_egfga Settings) SetUpdateFieldsOnOpen(b bool) {
	if !b {
		_egfga._ecefe.UpdateFields = nil
	} else {
		_egfga._ecefe.UpdateFields = _eaba.NewCT_OnOff()
	}
}

func _gbge(_gfaba *Document, _bggaf []*_eaba.CT_P, _beaa *TableInfo, _ffda *DrawingInfo) []Node {
	_fbce := []Node{}
	for _, _bdff := range _bggaf {
		_bagfd := Paragraph{_gfaba, _bdff}
		_efebc := Node{_abcbe: _gfaba, _efag: &_bagfd}
		if _eeae, _fgbec := _gfaba.Styles.SearchStyleById(_bagfd.Style()); _fgbec {
			_efebc.Style = _eeae
		}
		for _, _fbaf := range _bagfd.Runs() {
			_efebc.Children = append(_efebc.Children, Node{_abcbe: _gfaba, _efag: _fbaf, AnchoredDrawings: _fbaf.DrawingAnchored(), InlineDrawings: _fbaf.DrawingInline()})
		}
		_fbce = append(_fbce, _efebc)
	}
	return _fbce
}

// ComplexSizeValue returns the value of paragraph font size for complex fonts in points.
func (_eccgg ParagraphProperties) ComplexSizeValue() float64 {
	if _fagf := _eccgg._efec.RPr.SzCs; _fagf != nil {
		_ccdf := _fagf.ValAttr
		if _ccdf.ST_UnsignedDecimalNumber != nil {
			return float64(*_ccdf.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// X returns the inner wrapped XML type.
func (_fg AnchoredDrawing) X() *_eaba.WdAnchor { return _fg._cc }

// ClearContent clears any content in the run (text, tabs, breaks, etc.)
func (_ddeda Run) ClearContent() { _ddeda._bggd.EG_RunInnerContent = nil }

// WatermarkPicture is watermark picture within document.
type WatermarkPicture struct {
	_dadb  *_eaba.CT_Picture
	_cbfgc *_ffc.ShapeStyle
	_ggac  *_aae.Shape
	_bfag  *_aae.Shapetype
}

func _gfgd(_bfgb string) mergeFieldInfo {
	_gdea := []string{}
	_efcgd := _cb.Buffer{}
	_afeg := -1
	for _ggag, _dfgeb := range _bfgb {
		switch _dfgeb {
		case ' ':
			if _efcgd.Len() != 0 {
				_gdea = append(_gdea, _efcgd.String())
			}
			_efcgd.Reset()
		case '"':
			if _afeg != -1 {
				_gdea = append(_gdea, _bfgb[_afeg+1:_ggag])
				_afeg = -1
			} else {
				_afeg = _ggag
			}
		default:
			_efcgd.WriteRune(_dfgeb)
		}
	}
	if _efcgd.Len() != 0 {
		_gdea = append(_gdea, _efcgd.String())
	}
	_fdadb := mergeFieldInfo{}
	for _dfcd := 0; _dfcd < len(_gdea)-1; _dfcd++ {
		_eefa := _gdea[_dfcd]
		switch _eefa {
		case "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044":
			_fdadb._ggce = _gdea[_dfcd+1]
			_dfcd++
		case "\u005c\u0066":
			_fdadb._caaee = _gdea[_dfcd+1]
			_dfcd++
		case "\u005c\u0062":
			_fdadb._ddceg = _gdea[_dfcd+1]
			_dfcd++
		case "\u005c\u002a":
			switch _gdea[_dfcd+1] {
			case "\u0055\u0070\u0070e\u0072":
				_fdadb._aced = true
			case "\u004c\u006f\u0077e\u0072":
				_fdadb._gfcc = true
			case "\u0043\u0061\u0070\u0073":
				_fdadb._cbccd = true
			case "\u0046\u0069\u0072\u0073\u0074\u0043\u0061\u0070":
				_fdadb._cecf = true
			}
			_dfcd++
		}
	}
	return _fdadb
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (_daced ParagraphProperties) RStyle() string {
	if _daced._efec.RPr.RStyle != nil {
		return _daced._efec.RPr.RStyle.ValAttr
	}
	return ""
}

// Text return node and its child text,
func (_gdgc *Node) Text() string {
	_ebfa := _cb.NewBuffer([]byte{})
	switch _eedeg := _gdgc.X().(type) {
	case *Paragraph:
		for _, _cdafg := range _eedeg.Runs() {
			if _cdafg.Text() != "" {
				_ebfa.WriteString(_cdafg.Text())
				_ebfa.WriteString("\u000a")
			}
		}
	}
	for _, _effb := range _gdgc.Children {
		_ebfa.WriteString(_effb.Text())
	}
	return _ebfa.String()
}

// X returns the inner wrapped XML type.
func (_edad Style) X() *_eaba.CT_Style { return _edad._bgfca }

// ComplexSizeValue returns the value of run font size for complex fonts in points.
func (_dcbd RunProperties) ComplexSizeValue() float64 {
	if _dged := _dcbd._gfbgc.SzCs; _dged != nil {
		_befce := _dged.ValAttr
		if _befce.ST_UnsignedDecimalNumber != nil {
			return float64(*_befce.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SetFirstRow controls the conditional formatting for the first row in a table.
func (_bfbfa TableLook) SetFirstRow(on bool) {
	if !on {
		_bfbfa._ffdgf.FirstRowAttr = &_cg.ST_OnOff{}
		_bfbfa._ffdgf.FirstRowAttr.ST_OnOff1 = _cg.ST_OnOff1Off
	} else {
		_bfbfa._ffdgf.FirstRowAttr = &_cg.ST_OnOff{}
		_bfbfa._ffdgf.FirstRowAttr.ST_OnOff1 = _cg.ST_OnOff1On
	}
}

// Close closes the document, removing any temporary files that might have been
// created when opening a document.
func (_agag *Document) Close() error {
	if _agag.TmpPath != "" {
		return _gb.RemoveAll(_agag.TmpPath)
	}
	return nil
}

// SetRight sets the right border to a specified type, color and thickness.
func (_cfff TableBorders) SetRight(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_cfff._dbbcbf.Right = _eaba.NewCT_Border()
	_gfdf(_cfff._dbbcbf.Right, t, c, thickness)
}

// CellProperties are a table cells properties within a document.
type CellProperties struct{ _be *_eaba.CT_TcPr }

// SetLastColumn controls the conditional formatting for the last column in a table.
func (_fgfed TableLook) SetLastColumn(on bool) {
	if !on {
		_fgfed._ffdgf.LastColumnAttr = &_cg.ST_OnOff{}
		_fgfed._ffdgf.LastColumnAttr.ST_OnOff1 = _cg.ST_OnOff1Off
	} else {
		_fgfed._ffdgf.LastColumnAttr = &_cg.ST_OnOff{}
		_fgfed._ffdgf.LastColumnAttr.ST_OnOff1 = _cg.ST_OnOff1On
	}
}

// InitializeDefault constructs the default styles.
func (_gcggg Styles) InitializeDefault() {
	_gcggg.initializeDocDefaults()
	_gcggg.initializeStyleDefaults()
}

// AddFieldWithFormatting adds a field (automatically computed text) to the
// document with field specifc formatting.
func (_acabf Run) AddFieldWithFormatting(code string, fmt string, isDirty bool) {
	_dabe := _acabf.newIC()
	_dabe.FldChar = _eaba.NewCT_FldChar()
	_dabe.FldChar.FldCharTypeAttr = _eaba.ST_FldCharTypeBegin
	if isDirty {
		_dabe.FldChar.DirtyAttr = &_cg.ST_OnOff{}
		_dabe.FldChar.DirtyAttr.Bool = _d.Bool(true)
	}
	_dabe = _acabf.newIC()
	_dabe.InstrText = _eaba.NewCT_Text()
	if fmt != "" {
		_dabe.InstrText.Content = code + "\u0020" + fmt
	} else {
		_dabe.InstrText.Content = code
	}
	_dabe = _acabf.newIC()
	_dabe.FldChar = _eaba.NewCT_FldChar()
	_dabe.FldChar.FldCharTypeAttr = _eaba.ST_FldCharTypeEnd
}

// SetEndIndent controls the end indentation.
func (_agadd ParagraphProperties) SetEndIndent(m _eb.Distance) {
	if _agadd._efec.Ind == nil {
		_agadd._efec.Ind = _eaba.NewCT_Ind()
	}
	if m == _eb.Zero {
		_agadd._efec.Ind.EndAttr = nil
	} else {
		_agadd._efec.Ind.EndAttr = &_eaba.ST_SignedTwipsMeasure{}
		_agadd._efec.Ind.EndAttr.Int64 = _d.Int64(int64(m / _eb.Twips))
	}
}

// GetChartSpaceByRelId returns a *crt.ChartSpace with the associated relation ID in the
// document.
func (_bgd *Document) GetChartSpaceByRelId(relId string) *_dg.ChartSpace {
	_fffbe := _bgd._bdb.GetTargetByRelId(relId)
	for _, _aebc := range _bgd._cea {
		if _fffbe == _aebc.Target() {
			return _aebc._cebb
		}
	}
	return nil
}

// SetItalic sets the run to italic.
func (_adba RunProperties) SetItalic(b bool) {
	if !b {
		_adba._gfbgc.I = nil
		_adba._gfbgc.ICs = nil
	} else {
		_adba._gfbgc.I = _eaba.NewCT_OnOff()
		_adba._gfbgc.ICs = _eaba.NewCT_OnOff()
	}
}

// DrawingInline return a slice of InlineDrawings.
func (_ddbb Run) DrawingInline() []InlineDrawing {
	_edefd := []InlineDrawing{}
	for _, _bcgc := range _ddbb._bggd.EG_RunInnerContent {
		if _bcgc.Drawing == nil {
			continue
		}
		for _, _bgfaf := range _bcgc.Drawing.Inline {
			_edefd = append(_edefd, InlineDrawing{_ddbb._dddbb, _bgfaf})
		}
	}
	return _edefd
}

// UnderlineColor returns the hex color value of paragraph underline.
func (_ccfa ParagraphProperties) UnderlineColor() string {
	if _cabf := _ccfa._efec.RPr.U; _cabf != nil {
		_cacf := _cabf.ColorAttr
		if _cacf != nil && _cacf.ST_HexColorRGB != nil {
			return *_cacf.ST_HexColorRGB
		}
	}
	return ""
}

func (_bcdcc Paragraph) addFldCharsForField(_fafaa, _gbdag string) FormField {
	_degdc := _bcdcc.addBeginFldChar(_fafaa)
	_bcga := FormField{_cdag: _degdc}
	_fggbg := _bcdcc._bbgc.Bookmarks()
	_bgff := int64(len(_fggbg))
	if _fafaa != "" {
		_bcdcc.addStartBookmark(_bgff, _fafaa)
	}
	_bcdcc.addInstrText(_gbdag)
	_bcdcc.addSeparateFldChar()
	if _gbdag == "\u0046\u004f\u0052\u004d\u0054\u0045\u0058\u0054" {
		_abddbd := _bcdcc.AddRun()
		_acee := _eaba.NewEG_RunInnerContent()
		_abddbd._bggd.EG_RunInnerContent = []*_eaba.EG_RunInnerContent{_acee}
		_bcga._bebe = _acee
	}
	_bcdcc.addEndFldChar()
	if _fafaa != "" {
		_bcdcc.addEndBookmark(_bgff)
	}
	return _bcga
}

// TextItem is used for keeping text with references to a paragraph and run or a table, a row and a cell where it is located.
type TextItem struct {
	Text        string
	DrawingInfo *DrawingInfo
	Paragraph   *_eaba.CT_P
	Hyperlink   *_eaba.CT_Hyperlink
	Run         *_eaba.CT_R
	TableInfo   *TableInfo
}

func (_ddaf Endnote) content() []*_eaba.EG_ContentBlockContent {
	var _caeb []*_eaba.EG_ContentBlockContent
	for _, _eebd := range _ddaf._ggd.EG_BlockLevelElts {
		_caeb = append(_caeb, _eebd.EG_ContentBlockContent...)
	}
	return _caeb
}

// Paragraphs returns all of the paragraphs in the document body including tables.
func (_dddc *Document) Paragraphs() []Paragraph {
	_ccec := []Paragraph{}
	if _dddc._eeda.Body == nil {
		return nil
	}
	for _, _dfff := range _dddc._eeda.Body.EG_BlockLevelElts {
		for _, _gfeb := range _dfff.EG_ContentBlockContent {
			for _, _edca := range _gfeb.P {
				_ccec = append(_ccec, Paragraph{_dddc, _edca})
			}
		}
	}
	for _, _gab := range _dddc.Tables() {
		for _, _gabc := range _gab.Rows() {
			for _, _bdgd := range _gabc.Cells() {
				_ccec = append(_ccec, _bdgd.Paragraphs()...)
			}
		}
	}
	return _ccec
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (_cgce Paragraph) Style() string {
	if _cgce._dbef.PPr != nil && _cgce._dbef.PPr.PStyle != nil {
		return _cgce._dbef.PPr.PStyle.ValAttr
	}
	return ""
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (_fcggb ParagraphProperties) CharacterSpacingMeasure() string {
	if _bafa := _fcggb._efec.RPr.Spacing; _bafa != nil {
		_ccgag := _bafa.ValAttr
		if _ccgag.ST_UniversalMeasure != nil {
			return *_ccgag.ST_UniversalMeasure
		}
	}
	return ""
}

// ExtractTextOptions extraction text options.
type ExtractTextOptions struct {
	WithNumbering   bool
	NumberingIndent string
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (_geb CellBorders) SetInsideHorizontal(t _eaba.ST_Border, c _bc.Color, thickness _eb.Distance) {
	_geb._ggf.InsideH = _eaba.NewCT_Border()
	_gfdf(_geb._ggf.InsideH, t, c, thickness)
}

// AddField adds a field (automatically computed text) to the document.
func (_fabc Run) AddField(code string) { _fabc.AddFieldWithFormatting(code, "", true) }

// NumberingDefinition defines a numbering definition for a list of pragraphs.
type NumberingDefinition struct{ _fbde *_eaba.CT_AbstractNum }

// SetRowBandSize sets the number of Rows in the row band
func (_becgd TableStyleProperties) SetRowBandSize(rows int64) {
	_becgd._ddafc.TblStyleRowBandSize = _eaba.NewCT_DecimalNumber()
	_becgd._ddafc.TblStyleRowBandSize.ValAttr = rows
}

// Emboss returns true if paragraph emboss is on.
func (_begc ParagraphProperties) Emboss() bool { return _beff(_begc._efec.RPr.Emboss) }

// SetAlignment controls the paragraph alignment
func (_fbdea ParagraphStyleProperties) SetAlignment(align _eaba.ST_Jc) {
	if align == _eaba.ST_JcUnset {
		_fbdea._fded.Jc = nil
	} else {
		_fbdea._fded.Jc = _eaba.NewCT_Jc()
		_fbdea._fded.Jc.ValAttr = align
	}
}

// X returns the inner wrapped XML type.
func (_dbge TableWidth) X() *_eaba.CT_TblWidth { return _dbge._ebga }

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (_fcbg TableProperties) SetCellSpacingPercent(pct float64) {
	_fcbg._efac.TblCellSpacing = _eaba.NewCT_TblWidth()
	_fcbg._efac.TblCellSpacing.TypeAttr = _eaba.ST_TblWidthPct
	_fcbg._efac.TblCellSpacing.WAttr = &_eaba.ST_MeasurementOrPercent{}
	_fcbg._efac.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_eaba.ST_DecimalNumberOrPercent{}
	_fcbg._efac.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _d.Int64(int64(pct * 50))
}

// AddWatermarkPicture adds new watermark picture to document.
func (_fcgg *Document) AddWatermarkPicture(imageRef _ea.ImageRef) WatermarkPicture {
	var _fbbb []Header
	if _ceef, _cfcgd := _fcgg.BodySection().GetHeader(_eaba.ST_HdrFtrDefault); _cfcgd {
		_fbbb = append(_fbbb, _ceef)
	}
	if _abff, _gcgeg := _fcgg.BodySection().GetHeader(_eaba.ST_HdrFtrEven); _gcgeg {
		_fbbb = append(_fbbb, _abff)
	}
	if _eccc, _ddf := _fcgg.BodySection().GetHeader(_eaba.ST_HdrFtrFirst); _ddf {
		_fbbb = append(_fbbb, _eccc)
	}
	if len(_fbbb) < 1 {
		_dgd := _fcgg.AddHeader()
		_fcgg.BodySection().SetHeader(_dgd, _eaba.ST_HdrFtrDefault)
		_fbbb = append(_fbbb, _dgd)
	}
	var _facf error
	_geeg := NewWatermarkPicture()
	for _, _bgfa := range _fbbb {
		imageRef, _facf = _bgfa.AddImageRef(imageRef)
		if _facf != nil {
			return WatermarkPicture{}
		}
		_dga := _bgfa.Paragraphs()
		if len(_dga) < 1 {
			_bgbd := _bgfa.AddParagraph()
			_bgbd.AddRun().AddText("")
		}
		for _, _aaaa := range _bgfa.X().EG_ContentBlockContent {
			for _, _dgfg := range _aaaa.P {
				for _, _ggea := range _dgfg.EG_PContent {
					for _, _bbbge := range _ggea.EG_ContentRunContent {
						if _bbbge.R == nil {
							continue
						}
						for _, _dcdb := range _bbbge.R.EG_RunInnerContent {
							_dcdb.Pict = _geeg._dadb
							break
						}
					}
				}
			}
		}
	}
	_geeg.SetPicture(imageRef)
	return _geeg
}

// Save writes the document to an io.Writer in the Zip package format.
func (_dcge *Document) Save(w _aa.Writer) error { return _dcge.save(w, _dcge._ebd) }

// SetNextStyle sets the style that the next paragraph will use.
func (_afgbd Style) SetNextStyle(name string) {
	if name == "" {
		_afgbd._bgfca.Next = nil
	} else {
		_afgbd._bgfca.Next = _eaba.NewCT_String()
		_afgbd._bgfca.Next.ValAttr = name
	}
}

func _dffga() *_aae.Path {
	_gegb := _aae.NewPath()
	_gegb.TextpathokAttr = _cg.ST_TrueFalseTrue
	_gegb.ConnecttypeAttr = _aae.OfcST_ConnectTypeCustom
	_adedc := "\u0040\u0039\u002c0;\u0040\u0031\u0030\u002c\u0031\u0030\u0038\u0030\u0030;\u00401\u0031,\u00321\u0036\u0030\u0030\u003b\u0040\u0031\u0032\u002c\u0031\u0030\u0038\u0030\u0030"
	_gegb.ConnectlocsAttr = &_adedc
	_dffbb := "\u0032\u0037\u0030,\u0031\u0038\u0030\u002c\u0039\u0030\u002c\u0030"
	_gegb.ConnectanglesAttr = &_dffbb
	return _gegb
}

// SetHeadingLevel sets a heading level and style based on the level to a
// paragraph.  The default styles for a new unioffice document support headings
// from level 1 to 8.
func (_afaab ParagraphProperties) SetHeadingLevel(idx int) {
	_afaab.SetStyle(_eg.Sprintf("\u0048e\u0061\u0064\u0069\u006e\u0067\u0025d", idx))
	if _afaab._efec.NumPr == nil {
		_afaab._efec.NumPr = _eaba.NewCT_NumPr()
	}
	_afaab._efec.NumPr.Ilvl = _eaba.NewCT_DecimalNumber()
	_afaab._efec.NumPr.Ilvl.ValAttr = int64(idx)
}

// GetText returns text in the watermark.
func (_bbead *WatermarkText) GetText() string {
	_aega := _bbead.getShape()
	if _bbead._ggfda != nil {
		_bcbg := _bbead._ggfda.EG_ShapeElements
		if len(_bcbg) > 0 && _bcbg[0].Textpath != nil {
			return *_bcbg[0].Textpath.StringAttr
		}
	} else {
		_cgbfga := _bbead.findNode(_aega, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
		for _, _bfcd := range _cgbfga.Attrs {
			if _bfcd.Name.Local == "\u0073\u0074\u0072\u0069\u006e\u0067" {
				return _bfcd.Value
			}
		}
	}
	return ""
}

// WatermarkText is watermark text within the document.
type WatermarkText struct {
	_eefg  *_eaba.CT_Picture
	_ffcef *_ffc.TextpathStyle
	_ggfda *_aae.Shape
	_cfgg  *_aae.Shapetype
}

// SetStartPct sets the cell start margin
func (_acb CellMargins) SetStartPct(pct float64) {
	_acb._ec.Start = _eaba.NewCT_TblWidth()
	_cab(_acb._ec.Start, pct)
}

// SetAlignment sets the alignment of a table within the page.
func (_cegb TableProperties) SetAlignment(align _eaba.ST_JcTable) {
	if align == _eaba.ST_JcTableUnset {
		_cegb._efac.Jc = nil
	} else {
		_cegb._efac.Jc = _eaba.NewCT_JcTable()
		_cegb._efac.Jc.ValAttr = align
	}
}

// SetStrikeThrough sets the run to strike-through.
func (_fcdc RunProperties) SetStrikeThrough(b bool) {
	if !b {
		_fcdc._gfbgc.Strike = nil
	} else {
		_fcdc._gfbgc.Strike = _eaba.NewCT_OnOff()
	}
}

// FormFieldType is the type of the form field.
//go:generate stringer -type=FormFieldType
type FormFieldType byte

// RunProperties controls run styling properties
type RunProperties struct{ _gfbgc *_eaba.CT_RPr }
