package document

import (
	_fb "archive/zip"
	_ca "bytes"
	_g "errors"
	_af "fmt"
	_ba "image"
	_eg "image/jpeg"
	_cb "io"
	_fe "log"
	_fg "math/rand"
	_d "os"
	_b "path/filepath"
	_c "strings"
	_e "unicode"

	_f "gitee.com/gooffice/gooffice"
	_ed "gitee.com/gooffice/gooffice/color"
	_ec "gitee.com/gooffice/gooffice/common"
	_eb "gitee.com/gooffice/gooffice/common/license"
	_dd "gitee.com/gooffice/gooffice/common/tempstorage"
	_bf "gitee.com/gooffice/gooffice/measurement"
	_ea "gitee.com/gooffice/gooffice/schema/soo/dml"
	_afd "gitee.com/gooffice/gooffice/schema/soo/dml/picture"
	_dc "gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes"
	_ede "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships"
	_ag "gitee.com/gooffice/gooffice/schema/soo/wml"
	_bab "gitee.com/gooffice/gooffice/zippkg"
)

// Append appends a document d0 to a document d1. All settings, headers and footers remain the same as in the document d0 if they exist there, otherwise they are taken from the d1.
func (_edgb *Document) Append(d1orig *Document) error {
	_efcd, _edca := d1orig.Copy()
	if _edca != nil {
		return _edca
	}
	_edgb.DocBase = _edgb.DocBase.Append(_efcd.DocBase)
	if _efcd._bbgg.ConformanceAttr != _dc.ST_ConformanceClassStrict {
		_edgb._bbgg.ConformanceAttr = _efcd._bbgg.ConformanceAttr
	}
	_adcg := _edgb._agcc.X().Relationship
	_bcge := _efcd._agcc.X().Relationship
	_caad := _efcd._bbgg.Body
	_dcba := map[string]string{}
	_cfgaf := map[int64]int64{}
	_cgad := map[int64]int64{}
	for _, _ffdd := range _bcge {
		_dfc := true
		_efac := _ffdd.IdAttr
		_acfe := _ffdd.TargetAttr
		_bbaa := _ffdd.TypeAttr
		_caaf := _bbaa == _f.ImageType
		_gfddc := _bbaa == _f.HyperLinkType
		var _cbd string
		for _, _aega := range _adcg {
			if _aega.TypeAttr == _bbaa && _aega.TargetAttr == _acfe {
				_dfc = false
				_cbd = _aega.IdAttr
				break
			}
		}
		if _caaf {
			_dacg := "\u0077\u006f\u0072d\u002f" + _acfe
			for _, _dfaa := range _efcd.DocBase.Images {
				if _dfaa.Target() == _dacg {
					_ecag, _dbgb := _ec.ImageFromStorage(_dfaa.Path())
					if _dbgb != nil {
						return _dbgb
					}
					_aeb, _dbgb := _edgb.AddImage(_ecag)
					if _dbgb != nil {
						return _dbgb
					}
					_cbd = _aeb.RelID()
					break
				}
			}
		} else if _dfc {
			if _gfddc {
				_cace := _edgb._agcc.AddHyperlink(_acfe)
				_cbd = _ec.Relationship(_cace).ID()
			} else {
				_aef := _edgb._agcc.AddRelationship(_acfe, _bbaa)
				_cbd = _aef.X().IdAttr
			}
		}
		if _efac != _cbd {
			_dcba[_efac] = _cbd
		}
	}
	if _caad.SectPr != nil {
		for _, _cacg := range _caad.SectPr.EG_HdrFtrReferences {
			if _cacg.HeaderReference != nil {
				if _adfb, _fcfg := _dcba[_cacg.HeaderReference.IdAttr]; _fcfg {
					_cacg.HeaderReference.IdAttr = _adfb
					_edgb._gba = append(_edgb._gba, _ec.NewRelationships())
				}
			} else if _cacg.FooterReference != nil {
				if _bfcf, _fccd := _dcba[_cacg.FooterReference.IdAttr]; _fccd {
					_cacg.FooterReference.IdAttr = _bfcf
					_edgb._dbc = append(_edgb._dbc, _ec.NewRelationships())
				}
			}
		}
	}
	_dgc, _cacc := _edgb._fcb, _efcd._fcb
	if _dgc != nil {
		if _cacc != nil {
			if _dgc.Endnote != nil {
				if _cacc.Endnote != nil {
					_bcdd := int64(len(_dgc.Endnote) + 1)
					for _, _cgeba := range _cacc.Endnote {
						_eef := _cgeba.IdAttr
						if _eef > 0 {
							_cgeba.IdAttr = _bcdd
							_dgc.Endnote = append(_dgc.Endnote, _cgeba)
							_cgad[_eef] = _bcdd
							_bcdd++
						}
					}
				}
			} else {
				_dgc.Endnote = _cacc.Endnote
			}
		}
	} else if _cacc != nil {
		_dgc = _cacc
	}
	_edgb._fcb = _dgc
	_afgg, _gdb := _edgb._ab, _efcd._ab
	if _afgg != nil {
		if _gdb != nil {
			if _afgg.Footnote != nil {
				if _gdb.Footnote != nil {
					_agf := int64(len(_afgg.Footnote) + 1)
					for _, _dfdf := range _gdb.Footnote {
						_eggb := _dfdf.IdAttr
						if _eggb > 0 {
							_dfdf.IdAttr = _agf
							_afgg.Footnote = append(_afgg.Footnote, _dfdf)
							_cfgaf[_eggb] = _agf
							_agf++
						}
					}
				}
			} else {
				_afgg.Footnote = _gdb.Footnote
			}
		}
	} else if _gdb != nil {
		_afgg = _gdb
	}
	_edgb._ab = _afgg
	for _, _fee := range _caad.EG_BlockLevelElts {
		for _, _cfbfe := range _fee.EG_ContentBlockContent {
			for _, _cgd := range _cfbfe.P {
				_eddb(_cgd, _dcba)
				_faea(_cgd, _dcba)
				_bccgd(_cgd, _cfgaf, _cgad)
			}
			for _, _cfag := range _cfbfe.Tbl {
				_cefde(_cfag, _dcba)
				_gdgc(_cfag, _dcba)
				_fbd(_cfag, _cfgaf, _cgad)
			}
		}
	}
	_edgb._bbgg.Body.EG_BlockLevelElts = append(_edgb._bbgg.Body.EG_BlockLevelElts, _efcd._bbgg.Body.EG_BlockLevelElts...)
	if _edgb._bbgg.Body.SectPr == nil {
		_edgb._bbgg.Body.SectPr = _efcd._bbgg.Body.SectPr
	} else {
		var _faae, _cfgcd bool
		for _, _febb := range _edgb._bbgg.Body.SectPr.EG_HdrFtrReferences {
			if _febb.HeaderReference != nil {
				_faae = true
			} else if _febb.FooterReference != nil {
				_cfgcd = true
			}
		}
		if !_faae {
			for _, _eccd := range _efcd._bbgg.Body.SectPr.EG_HdrFtrReferences {
				if _eccd.HeaderReference != nil {
					_edgb._bbgg.Body.SectPr.EG_HdrFtrReferences = append(_edgb._bbgg.Body.SectPr.EG_HdrFtrReferences, _eccd)
					break
				}
			}
		}
		if !_cfgcd {
			for _, _debe := range _efcd._bbgg.Body.SectPr.EG_HdrFtrReferences {
				if _debe.FooterReference != nil {
					_edgb._bbgg.Body.SectPr.EG_HdrFtrReferences = append(_edgb._bbgg.Body.SectPr.EG_HdrFtrReferences, _debe)
					break
				}
			}
		}
		if _edgb._bbgg.Body.SectPr.Cols == nil && _efcd._bbgg.Body.SectPr.Cols != nil {
			_edgb._bbgg.Body.SectPr.Cols = _efcd._bbgg.Body.SectPr.Cols
		}
	}
	_ebbf := _edgb.Numbering._eefa
	_cbe := _efcd.Numbering._eefa
	if _ebbf != nil {
		if _cbe != nil {
			_ebbf.NumPicBullet = append(_ebbf.NumPicBullet, _cbe.NumPicBullet...)
			_ebbf.AbstractNum = append(_ebbf.AbstractNum, _cbe.AbstractNum...)
			_ebbf.Num = append(_ebbf.Num, _cbe.Num...)
		}
	} else if _cbe != nil {
		_ebbf = _cbe
	}
	_edgb.Numbering._eefa = _ebbf
	if _edgb.Styles._faabg == nil && _efcd.Styles._faabg != nil {
		_edgb.Styles._faabg = _efcd.Styles._faabg
	}
	_edgb._adae = append(_edgb._adae, _efcd._adae...)
	if len(_edgb._cg) == 0 {
		_edgb._cg = _efcd._cg
	}
	if len(_edgb._bbc) == 0 {
		_edgb._bbc = _efcd._bbc
	}
	_dgad := _edgb._cef
	_bebg := _efcd._cef
	if _dgad != nil {
		if _bebg != nil {
			if _dgad.Divs != nil {
				if _bebg.Divs != nil {
					_dgad.Divs.Div = append(_dgad.Divs.Div, _bebg.Divs.Div...)
				}
			} else {
				_dgad.Divs = _bebg.Divs
			}
		}
		_dgad.Frameset = nil
	} else if _bebg != nil {
		_dgad = _bebg
		_dgad.Frameset = nil
	}
	_edgb._cef = _dgad
	_bged := _edgb._gc
	_eebb := _efcd._gc
	if _bged != nil {
		if _eebb != nil {
			if _bged.Font != nil {
				if _eebb.Font != nil {
					for _, _cegf := range _eebb.Font {
						_cecb := true
						for _, _acfd := range _bged.Font {
							if _acfd.NameAttr == _cegf.NameAttr {
								_cecb = false
								break
							}
						}
						if _cecb {
							_bged.Font = append(_bged.Font, _cegf)
						}
					}
				}
			} else {
				_bged.Font = _eebb.Font
			}
		}
	} else if _eebb != nil {
		_bged = _eebb
	}
	_edgb._gc = _bged
	return nil
}

// AddTable adds a table to the table cell.
func (_cfg Cell) AddTable() Table {
	_bcb := _ag.NewEG_BlockLevelElts()
	_cfg._afg.EG_BlockLevelElts = append(_cfg._afg.EG_BlockLevelElts, _bcb)
	_acg := _ag.NewEG_ContentBlockContent()
	_bcb.EG_ContentBlockContent = append(_bcb.EG_ContentBlockContent, _acg)
	_bd := _ag.NewCT_Tbl()
	_acg.Tbl = append(_acg.Tbl, _bd)
	return Table{_cfg._dce, _bd}
}

// AddDrawingInline adds an inline drawing from an ImageRef.
func (_acda Run) AddDrawingInline(img _ec.ImageRef) (InlineDrawing, error) {
	_cacd := _acda.newIC()
	_cacd.Drawing = _ag.NewCT_Drawing()
	_dacgd := _ag.NewWdInline()
	_babb := InlineDrawing{_acda._ebdd, _dacgd}
	_dacgd.CNvGraphicFramePr = _ea.NewCT_NonVisualGraphicFrameProperties()
	_cacd.Drawing.Inline = append(_cacd.Drawing.Inline, _dacgd)
	_dacgd.Graphic = _ea.NewGraphic()
	_dacgd.Graphic.GraphicData = _ea.NewCT_GraphicalObjectData()
	_dacgd.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"
	_dacgd.DistTAttr = _f.Uint32(0)
	_dacgd.DistLAttr = _f.Uint32(0)
	_dacgd.DistBAttr = _f.Uint32(0)
	_dacgd.DistRAttr = _f.Uint32(0)
	_dacgd.Extent.CxAttr = int64(float64(img.Size().X*_bf.Pixel72) / _bf.EMU)
	_dacgd.Extent.CyAttr = int64(float64(img.Size().Y*_bf.Pixel72) / _bf.EMU)
	_faaf := 0x7FFFFFFF & _fg.Uint32()
	_dacgd.DocPr.IdAttr = _faaf
	_bdaaf := _afd.NewPic()
	_bdaaf.NvPicPr.CNvPr.IdAttr = _faaf
	_cgbd := img.RelID()
	if _cgbd == "" {
		return _babb, _g.New("\u0063\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020\u0074\u006f\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0020\u0072\u0065l\u0061\u0074\u0069o\u006e\u0073")
	}
	_dacgd.Graphic.GraphicData.Any = append(_dacgd.Graphic.GraphicData.Any, _bdaaf)
	_bdaaf.BlipFill = _ea.NewCT_BlipFillProperties()
	_bdaaf.BlipFill.Blip = _ea.NewCT_Blip()
	_bdaaf.BlipFill.Blip.EmbedAttr = &_cgbd
	_bdaaf.BlipFill.Stretch = _ea.NewCT_StretchInfoProperties()
	_bdaaf.BlipFill.Stretch.FillRect = _ea.NewCT_RelativeRect()
	_bdaaf.SpPr = _ea.NewCT_ShapeProperties()
	_bdaaf.SpPr.Xfrm = _ea.NewCT_Transform2D()
	_bdaaf.SpPr.Xfrm.Off = _ea.NewCT_Point2D()
	_bdaaf.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _f.Int64(0)
	_bdaaf.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _f.Int64(0)
	_bdaaf.SpPr.Xfrm.Ext = _ea.NewCT_PositiveSize2D()
	_bdaaf.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * _bf.Point)
	_bdaaf.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * _bf.Point)
	_bdaaf.SpPr.PrstGeom = _ea.NewCT_PresetGeometry2D()
	_bdaaf.SpPr.PrstGeom.PrstAttr = _ea.ST_ShapeTypeRect
	return _babb, nil
}

// Clear clears all content within a footer
func (_gcg Footer) Clear() { _gcg._edec.EG_ContentBlockContent = nil }

// SetRowBandSize sets the number of Rows in the row band
func (_gfaag TableStyleProperties) SetRowBandSize(rows int64) {
	_gfaag._cgcba.TblStyleRowBandSize = _ag.NewCT_DecimalNumber()
	_gfaag._cgcba.TblStyleRowBandSize.ValAttr = rows
}

// AddParagraph adds a new paragraph to the document body.
func (_dba *Document) AddParagraph() Paragraph {
	_eagd := _ag.NewEG_BlockLevelElts()
	_dba._bbgg.Body.EG_BlockLevelElts = append(_dba._bbgg.Body.EG_BlockLevelElts, _eagd)
	_fgba := _ag.NewEG_ContentBlockContent()
	_eagd.EG_ContentBlockContent = append(_eagd.EG_ContentBlockContent, _fgba)
	_gcf := _ag.NewCT_P()
	_fgba.P = append(_fgba.P, _gcf)
	return Paragraph{_dba, _gcf}
}

// SetAlignment sets the alignment of a table within the page.
func (_ggfb TableProperties) SetAlignment(align _ag.ST_JcTable) {
	if align == _ag.ST_JcTableUnset {
		_ggfb._gfaa.Jc = nil
	} else {
		_ggfb._gfaa.Jc = _ag.NewCT_JcTable()
		_ggfb._gfaa.Jc.ValAttr = align
	}
}

// Font returns the name of paragraph font family.
func (_ddgg ParagraphProperties) Font() string {
	if _gegd := _ddgg._cfdd.RPr.RFonts; _gegd != nil {
		if _gegd.AsciiAttr != nil {
			return *_gegd.AsciiAttr
		} else if _gegd.HAnsiAttr != nil {
			return *_gegd.HAnsiAttr
		} else if _gegd.CsAttr != nil {
			return *_gegd.CsAttr
		}
	}
	return ""
}

// RemoveFootnote removes a footnote from both the paragraph and the document
// the requested footnote must be anchored on the paragraph being referenced.
func (_fdcc Paragraph) RemoveFootnote(id int64) {
	_edda := _fdcc._aefd._ab
	var _bdbb int
	for _effd, _ffaa := range _edda.CT_Footnotes.Footnote {
		if _ffaa.IdAttr == id {
			_bdbb = _effd
		}
	}
	_bdbb = 0
	_edda.CT_Footnotes.Footnote[_bdbb] = nil
	_edda.CT_Footnotes.Footnote[_bdbb] = _edda.CT_Footnotes.Footnote[len(_edda.CT_Footnotes.Footnote)-1]
	_edda.CT_Footnotes.Footnote = _edda.CT_Footnotes.Footnote[:len(_edda.CT_Footnotes.Footnote)-1]
	var _dff Run
	for _, _adba := range _fdcc.Runs() {
		if _dbcca, _fgdc := _adba.IsFootnote(); _dbcca {
			if _fgdc == id {
				_dff = _adba
			}
		}
	}
	_fdcc.RemoveRun(_dff)
}

// NewNumbering constructs a new numbering.
func NewNumbering() Numbering { _dcda := _ag.NewNumbering(); return Numbering{_dcda} }

// Type returns the type of the style.
func (_cbff Style) Type() _ag.ST_StyleType { return _cbff._gade.TypeAttr }

// Properties returns the run properties.
func (_gcede Run) Properties() RunProperties {
	if _gcede._fegg.RPr == nil {
		_gcede._fegg.RPr = _ag.NewCT_RPr()
	}
	return RunProperties{_gcede._fegg.RPr}
}

// AddTabStop adds a tab stop to the paragraph.
func (_dfaac ParagraphStyleProperties) AddTabStop(position _bf.Distance, justificaton _ag.ST_TabJc, leader _ag.ST_TabTlc) {
	if _dfaac._ddde.Tabs == nil {
		_dfaac._ddde.Tabs = _ag.NewCT_Tabs()
	}
	_bdg := _ag.NewCT_TabStop()
	_bdg.LeaderAttr = leader
	_bdg.ValAttr = justificaton
	_bdg.PosAttr.Int64 = _f.Int64(int64(position / _bf.Twips))
	_dfaac._ddde.Tabs.Tab = append(_dfaac._ddde.Tabs.Tab, _bdg)
}
func (_dfff Paragraph) addBeginFldChar(_bcdf string) *_ag.CT_FFData {
	_ebgb := _dfff.addFldChar()
	_ebgb.FldCharTypeAttr = _ag.ST_FldCharTypeBegin
	_ebgb.FfData = _ag.NewCT_FFData()
	_eggg := _ag.NewCT_FFName()
	_eggg.ValAttr = &_bcdf
	_ebgb.FfData.Name = []*_ag.CT_FFName{_eggg}
	return _ebgb.FfData
}

// AddTabStop adds a tab stop to the paragraph.  It controls the position of text when using Run.AddTab()
func (_gecc ParagraphProperties) AddTabStop(position _bf.Distance, justificaton _ag.ST_TabJc, leader _ag.ST_TabTlc) {
	if _gecc._cfdd.Tabs == nil {
		_gecc._cfdd.Tabs = _ag.NewCT_Tabs()
	}
	_cedd := _ag.NewCT_TabStop()
	_cedd.LeaderAttr = leader
	_cedd.ValAttr = justificaton
	_cedd.PosAttr.Int64 = _f.Int64(int64(position / _bf.Twips))
	_gecc._cfdd.Tabs.Tab = append(_gecc._cfdd.Tabs.Tab, _cedd)
}

// Footnote returns the footnote based on the ID; this can be used nicely with
// the run.IsFootnote() functionality.
func (_edee *Document) Footnote(id int64) Footnote {
	for _, _cefd := range _edee.Footnotes() {
		if _cefd.id() == id {
			return _cefd
		}
	}
	return Footnote{}
}

// SetStartIndent controls the start indent of the paragraph.
func (_fbgfge ParagraphStyleProperties) SetStartIndent(m _bf.Distance) {
	if _fbgfge._ddde.Ind == nil {
		_fbgfge._ddde.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_fbgfge._ddde.Ind.StartAttr = nil
	} else {
		_fbgfge._ddde.Ind.StartAttr = &_ag.ST_SignedTwipsMeasure{}
		_fbgfge._ddde.Ind.StartAttr.Int64 = _f.Int64(int64(m / _bf.Twips))
	}
}

// SetAll sets all of the borders to a given value.
func (_da CellBorders) SetAll(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_da.SetBottom(t, c, thickness)
	_da.SetLeft(t, c, thickness)
	_da.SetRight(t, c, thickness)
	_da.SetTop(t, c, thickness)
	_da.SetInsideHorizontal(t, c, thickness)
	_da.SetInsideVertical(t, c, thickness)
}

// Table is a table within a document.
type Table struct {
	_aebg  *Document
	_fdeda *_ag.CT_Tbl
}

// Value returns the tring value of a FormFieldTypeText or FormFieldTypeDropDown.
func (_ggf FormField) Value() string {
	if _ggf._bffef.TextInput != nil && _ggf._eeef.T != nil {
		return _ggf._eeef.T.Content
	} else if _ggf._bffef.DdList != nil && _ggf._bffef.DdList.Result != nil {
		_adbb := _ggf.PossibleValues()
		_cbdf := int(_ggf._bffef.DdList.Result.ValAttr)
		if _cbdf < len(_adbb) {
			return _adbb[_cbdf]
		}
	} else if _ggf._bffef.CheckBox != nil {
		if _ggf.IsChecked() {
			return "\u0074\u0072\u0075\u0065"
		}
		return "\u0066\u0061\u006cs\u0065"
	}
	return ""
}

// Paragraph is a paragraph within a document.
type Paragraph struct {
	_aefd *Document
	_dbfe *_ag.CT_P
}

// SetAlignment controls the paragraph alignment
func (_aaec ParagraphProperties) SetAlignment(align _ag.ST_Jc) {
	if align == _ag.ST_JcUnset {
		_aaec._cfdd.Jc = nil
	} else {
		_aaec._cfdd.Jc = _ag.NewCT_Jc()
		_aaec._cfdd.Jc.ValAttr = align
	}
}

// AddEndnote will create a new endnote and attach it to the Paragraph in the
// location at the end of the previous run (endnotes create their own run within
// the paragraph. The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the endnote later.
func (_dcebf Paragraph) AddEndnote(text string) Endnote {
	var _adgda int64
	if _dcebf._aefd.HasEndnotes() {
		for _, _cddb := range _dcebf._aefd.Endnotes() {
			if _cddb.id() > _adgda {
				_adgda = _cddb.id()
			}
		}
		_adgda++
	} else {
		_adgda = 0
		_dcebf._aefd._fcb = &_ag.Endnotes{}
	}
	_dabc := _ag.NewCT_FtnEdn()
	_faeb := _ag.NewCT_FtnEdnRef()
	_faeb.IdAttr = _adgda
	_dcebf._aefd._fcb.CT_Endnotes.Endnote = append(_dcebf._aefd._fcb.CT_Endnotes.Endnote, _dabc)
	_adfa := _dcebf.AddRun()
	_eeddg := _adfa.Properties()
	_eeddg.SetStyle("\u0045\u006e\u0064\u006e\u006f\u0074\u0065\u0041\u006e\u0063\u0068\u006f\u0072")
	_adfa._fegg.EG_RunInnerContent = []*_ag.EG_RunInnerContent{_ag.NewEG_RunInnerContent()}
	_adfa._fegg.EG_RunInnerContent[0].EndnoteReference = _faeb
	_dagc := Endnote{_dcebf._aefd, _dabc}
	_dagc._gae.IdAttr = _adgda
	_dagc._gae.EG_BlockLevelElts = []*_ag.EG_BlockLevelElts{_ag.NewEG_BlockLevelElts()}
	_agcg := _dagc.AddParagraph()
	_agcg.Properties().SetStyle("\u0045n\u0064\u006e\u006f\u0074\u0065")
	_agcg._dbfe.PPr.RPr = _ag.NewCT_ParaRPr()
	_bfec := _agcg.AddRun()
	_bfec.AddTab()
	_bfec.AddText(text)
	return _dagc
}

// SetFooter sets a section footer.
func (_bagg Section) SetFooter(f Footer, t _ag.ST_HdrFtr) {
	_ddgdc := _ag.NewEG_HdrFtrReferences()
	_bagg._cdec.EG_HdrFtrReferences = append(_bagg._cdec.EG_HdrFtrReferences, _ddgdc)
	_ddgdc.FooterReference = _ag.NewCT_HdrFtrRef()
	_ddgdc.FooterReference.TypeAttr = t
	_faag := _bagg._cce._agcc.FindRIDForN(f.Index(), _f.FooterType)
	if _faag == "" {
		_fe.Print("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u0066\u006f\u006f\u0074\u0065r \u0049\u0044")
	}
	_ddgdc.FooterReference.IdAttr = _faag
}

// TableLook is the conditional formatting associated with a table style that
// has been assigned to a table.
type TableLook struct{ _cfbff *_ag.CT_TblLook }

// Close closes the document, removing any temporary files that might have been
// created when opening a document.
func (_eebe *Document) Close() error {
	if _eebe.TmpPath != "" {
		return _dd.RemoveAll(_eebe.TmpPath)
	}
	return nil
}

// Footers returns the footers defined in the document.
func (_efg *Document) Footers() []Footer {
	_aad := []Footer{}
	for _, _cccg := range _efg._bbc {
		_aad = append(_aad, Footer{_efg, _cccg})
	}
	return _aad
}

// RemoveParagraph removes a paragraph from a document.
func (_bca *Document) RemoveParagraph(p Paragraph) {
	if _bca._bbgg.Body == nil {
		return
	}
	for _, _ccca := range _bca._bbgg.Body.EG_BlockLevelElts {
		for _, _eac := range _ccca.EG_ContentBlockContent {
			for _dded, _cdb := range _eac.P {
				if _cdb == p._dbfe {
					copy(_eac.P[_dded:], _eac.P[_dded+1:])
					_eac.P = _eac.P[0 : len(_eac.P)-1]
					return
				}
			}
			if _eac.Sdt != nil && _eac.Sdt.SdtContent != nil && _eac.Sdt.SdtContent.P != nil {
				for _aadg, _gde := range _eac.Sdt.SdtContent.P {
					if _gde == p._dbfe {
						copy(_eac.P[_aadg:], _eac.P[_aadg+1:])
						_eac.P = _eac.P[0 : len(_eac.P)-1]
						return
					}
				}
			}
		}
	}
	for _, _aaea := range _bca.Tables() {
		for _, _fff := range _aaea.Rows() {
			for _, _dfd := range _fff.Cells() {
				for _, _dgae := range _dfd._afg.EG_BlockLevelElts {
					for _, _gee := range _dgae.EG_ContentBlockContent {
						for _bbf, _ffe := range _gee.P {
							if _ffe == p._dbfe {
								copy(_gee.P[_bbf:], _gee.P[_bbf+1:])
								_gee.P = _gee.P[0 : len(_gee.P)-1]
								return
							}
						}
					}
				}
			}
		}
	}
	for _, _aee := range _bca.Headers() {
		_aee.RemoveParagraph(p)
	}
	for _, _cgg := range _bca.Footers() {
		_cgg.RemoveParagraph(p)
	}
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (_ccfg TableBorders) SetInsideVertical(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_ccfg._egbeb.InsideV = _ag.NewCT_Border()
	_ddecc(_ccfg._egbeb.InsideV, t, c, thickness)
}

// Headers returns the headers defined in the document.
func (_ggc *Document) Headers() []Header {
	_fef := []Header{}
	for _, _debc := range _ggc._cg {
		_fef = append(_fef, Header{_ggc, _debc})
	}
	return _fef
}

// SetContextualSpacing controls whether to Ignore Spacing Above and Below When
// Using Identical Styles
func (_fdagc ParagraphStyleProperties) SetContextualSpacing(b bool) {
	if !b {
		_fdagc._ddde.ContextualSpacing = nil
	} else {
		_fdagc._ddde.ContextualSpacing = _ag.NewCT_OnOff()
	}
}
func (_agg Paragraph) addEndBookmark(_fffd int64) *_ag.CT_MarkupRange {
	_dfbc := _ag.NewEG_PContent()
	_agg._dbfe.EG_PContent = append(_agg._dbfe.EG_PContent, _dfbc)
	_egee := _ag.NewEG_ContentRunContent()
	_eacb := _ag.NewEG_RunLevelElts()
	_ebd := _ag.NewEG_RangeMarkupElements()
	_bgga := _ag.NewCT_MarkupRange()
	_bgga.IdAttr = _fffd
	_ebd.BookmarkEnd = _bgga
	_dfbc.EG_ContentRunContent = append(_dfbc.EG_ContentRunContent, _egee)
	_egee.EG_RunLevelElts = append(_egee.EG_RunLevelElts, _eacb)
	_eacb.EG_RangeMarkupElements = append(_eacb.EG_RangeMarkupElements, _ebd)
	return _bgga
}

// SetShadow sets the run to shadowed text.
func (_ccad RunProperties) SetShadow(b bool) {
	if !b {
		_ccad._ddda.Shadow = nil
	} else {
		_ccad._ddda.Shadow = _ag.NewCT_OnOff()
	}
}

// SetAll sets all of the borders to a given value.
func (_aeec TableBorders) SetAll(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_aeec.SetBottom(t, c, thickness)
	_aeec.SetLeft(t, c, thickness)
	_aeec.SetRight(t, c, thickness)
	_aeec.SetTop(t, c, thickness)
	_aeec.SetInsideHorizontal(t, c, thickness)
	_aeec.SetInsideVertical(t, c, thickness)
}

type mergeFieldInfo struct {
	_cabg                 string
	_ebce                 string
	_gbeg                 string
	_ebe                  bool
	_gdad                 bool
	_efad                 bool
	_fcd                  bool
	_cgdg                 Paragraph
	_gaab, _cgfaf, _efged int
	_fcdf                 *_ag.EG_PContent
	_baba                 bool
}

// X returns the inner wrapped XML type.
func (_cad Color) X() *_ag.CT_Color { return _cad._ccc }

// Properties returns the numbering level paragraph properties.
func (_bga NumberingLevel) Properties() ParagraphStyleProperties {
	if _bga._bdce.PPr == nil {
		_bga._bdce.PPr = _ag.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_bga._bdce.PPr}
}

// Styles is the document wide styles contained in styles.xml.
type Styles struct{ _faabg *_ag.Styles }

func _cefde(_ffff *_ag.CT_Tbl, _eccde map[string]string) {
	for _, _acaf := range _ffff.EG_ContentRowContent {
		for _, _gbed := range _acaf.Tr {
			for _, _agcd := range _gbed.EG_ContentCellContent {
				for _, _fefb := range _agcd.Tc {
					for _, _cadg := range _fefb.EG_BlockLevelElts {
						for _, _dadc := range _cadg.EG_ContentBlockContent {
							for _, _dag := range _dadc.P {
								_eddb(_dag, _eccde)
							}
							for _, _abe := range _dadc.Tbl {
								_cefde(_abe, _eccde)
							}
						}
					}
				}
			}
		}
	}
}

// SetColumnSpan sets the number of Grid Columns Spanned by the Cell.  This is used
// to give the appearance of merged cells.
func (_gb CellProperties) SetColumnSpan(cols int) {
	if cols == 0 {
		_gb._dac.GridSpan = nil
	} else {
		_gb._dac.GridSpan = _ag.NewCT_DecimalNumber()
		_gb._dac.GridSpan.ValAttr = int64(cols)
	}
}

// Copy makes a deep copy of the document by saving and reading it back.
// It can be useful to avoid sharing common data between two documents.
func (_bgc *Document) Copy() (*Document, error) {
	_eaeg := _ca.NewBuffer([]byte{})
	_eeg := _bgc.Save(_eaeg)
	if _eeg != nil {
		return nil, _eeg
	}
	_bdd := _eaeg.Bytes()
	_fgbc := _ca.NewReader(_bdd)
	return Read(_fgbc, int64(_fgbc.Len()))
}

// ParagraphSpacing controls the spacing for a paragraph and its lines.
type ParagraphSpacing struct{ _gggda *_ag.CT_Spacing }

// Shadow returns true if paragraph shadow is on.
func (_gbcd ParagraphProperties) Shadow() bool { return _decg(_gbcd._cfdd.RPr.Shadow) }

// TableLook returns the table look, or conditional formatting applied to a table style.
func (_ccage TableProperties) TableLook() TableLook {
	if _ccage._gfaa.TblLook == nil {
		_ccage._gfaa.TblLook = _ag.NewCT_TblLook()
	}
	return TableLook{_ccage._gfaa.TblLook}
}

// Read reads a document from an io.Reader.
func Read(r _cb.ReaderAt, size int64) (*Document, error) {
	_bgbf := New()
	_bgbf.Numbering._eefa = nil
	_dcfa, _eabg := _dd.TempDir("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0064\u006f\u0063\u0078")
	if _eabg != nil {
		return nil, _eabg
	}
	_bgbf.TmpPath = _dcfa
	_bdb, _eabg := _fb.NewReader(r, size)
	if _eabg != nil {
		return nil, _af.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073", _eabg)
	}
	_ddb := []*_fb.File{}
	_ddb = append(_ddb, _bdb.File...)
	_adfda := false
	for _, _ecc := range _ddb {
		if _ecc.FileHeader.Name == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
			_adfda = true
			break
		}
	}
	if _adfda {
		_bgbf.CreateCustomProperties()
	}
	_ffbe := _bgbf._bbgg.ConformanceAttr
	_dae := _bab.DecodeMap{}
	_dae.SetOnNewRelationshipFunc(_bgbf.onNewRelationship)
	_dae.AddTarget(_f.ContentTypesFilename, _bgbf.ContentTypes.X(), "", 0)
	_dae.AddTarget(_f.BaseRelsFilename, _bgbf.Rels.X(), "", 0)
	if _fgg := _dae.Decode(_ddb); _fgg != nil {
		return nil, _fgg
	}
	_bgbf._bbgg.ConformanceAttr = _ffbe
	for _, _edcd := range _ddb {
		if _edcd == nil {
			continue
		}
		if _edcf := _bgbf.AddExtraFileFromZip(_edcd); _edcf != nil {
			return nil, _edcf
		}
	}
	if _adfda {
		_bffe := false
		for _, _adc := range _bgbf.Rels.X().Relationship {
			if _adc.TargetAttr == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
				_bffe = true
				break
			}
		}
		if !_bffe {
			_bgbf.AddCustomRelationships()
		}
	}
	return _bgbf, nil
}

// SetStrikeThrough sets the run to strike-through.
func (_gbecc RunProperties) SetStrikeThrough(b bool) {
	if !b {
		_gbecc._ddda.Strike = nil
	} else {
		_gbecc._ddda.Strike = _ag.NewCT_OnOff()
	}
}

// CellProperties returns the cell properties.
func (_eggcg TableConditionalFormatting) CellProperties() CellProperties {
	if _eggcg._bdceb.TcPr == nil {
		_eggcg._bdceb.TcPr = _ag.NewCT_TcPr()
	}
	return CellProperties{_eggcg._bdceb.TcPr}
}

// Color controls the run or styles color.
type Color struct{ _ccc *_ag.CT_Color }

// SetUnderline controls underline for a run style.
func (_ceage RunProperties) SetUnderline(style _ag.ST_Underline, c _ed.Color) {
	if style == _ag.ST_UnderlineUnset {
		_ceage._ddda.U = nil
	} else {
		_ceage._ddda.U = _ag.NewCT_Underline()
		_ceage._ddda.U.ColorAttr = &_ag.ST_HexColor{}
		_ceage._ddda.U.ColorAttr.ST_HexColorRGB = c.AsRGBString()
		_ceage._ddda.U.ValAttr = style
	}
}

// X returns the inner wrapped XML type.
func (_fgcgg Style) X() *_ag.CT_Style { return _fgcgg._gade }

// TableBorders allows manipulation of borders on a table.
type TableBorders struct{ _egbeb *_ag.CT_TblBorders }

// SetConformance sets conformance attribute of the document
// as one of these values from gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (_agegd Document) SetConformance(conformanceAttr _dc.ST_ConformanceClass) {
	_agegd._bbgg.ConformanceAttr = conformanceAttr
}

// Outline returns true if paragraph outline is on.
func (_bgdg ParagraphProperties) Outline() bool { return _decg(_bgdg._cfdd.RPr.Outline) }

// Definitions returns the defined numbering definitions.
func (_gcag Numbering) Definitions() []NumberingDefinition {
	_ffc := []NumberingDefinition{}
	for _, _gacb := range _gcag._eefa.AbstractNum {
		_ffc = append(_ffc, NumberingDefinition{_gacb})
	}
	return _ffc
}

// EastAsiaFont returns the name of paragraph font family for East Asia.
func (_adaed ParagraphProperties) EastAsiaFont() string {
	if _bfbd := _adaed._cfdd.RPr.RFonts; _bfbd != nil {
		if _bfbd.EastAsiaAttr != nil {
			return *_bfbd.EastAsiaAttr
		}
	}
	return ""
}

// X returns the inner wml.CT_TblBorders
func (_aeed TableBorders) X() *_ag.CT_TblBorders { return _aeed._egbeb }

// AnchoredDrawing is an absolutely positioned image within a document page.
type AnchoredDrawing struct {
	_aa *Document
	_de *_ag.WdAnchor
}

// X returns the inner wrapped XML type.
func (_bbac RunProperties) X() *_ag.CT_RPr { return _bbac._ddda }

// SetFormat sets the numbering format.
func (_daag NumberingLevel) SetFormat(f _ag.ST_NumberFormat) {
	if _daag._bdce.NumFmt == nil {
		_daag._bdce.NumFmt = _ag.NewCT_NumFmt()
	}
	_daag._bdce.NumFmt.ValAttr = f
}

// Borders allows manipulation of the table borders.
func (_baeg TableStyleProperties) Borders() TableBorders {
	if _baeg._cgcba.TblBorders == nil {
		_baeg._cgcba.TblBorders = _ag.NewCT_TblBorders()
	}
	return TableBorders{_baeg._cgcba.TblBorders}
}

// SetAlignment sets the paragraph alignment
func (_eeed NumberingLevel) SetAlignment(j _ag.ST_Jc) {
	if j == _ag.ST_JcUnset {
		_eeed._bdce.LvlJc = nil
	} else {
		_eeed._bdce.LvlJc = _ag.NewCT_Jc()
		_eeed._bdce.LvlJc.ValAttr = j
	}
}

// IsEndnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (_baab Run) IsEndnote() (bool, int64) {
	if _baab._fegg.EG_RunInnerContent != nil {
		if _baab._fegg.EG_RunInnerContent[0].EndnoteReference != nil {
			return true, _baab._fegg.EG_RunInnerContent[0].EndnoteReference.IdAttr
		}
	}
	return false, 0
}

// SetEnabled marks a FormField as enabled or disabled.
func (_gaeg FormField) SetEnabled(enabled bool) {
	_bagc := _ag.NewCT_OnOff()
	_bagc.ValAttr = &_dc.ST_OnOff{Bool: &enabled}
	_gaeg._bffef.Enabled = []*_ag.CT_OnOff{_bagc}
}

// AddDropdownList adds dropdown list form field to the paragraph and returns it.
func (_cddbd Paragraph) AddDropdownList(name string) FormField {
	_bfcd := _cddbd.addFldCharsForField(name, "\u0046\u004f\u0052M\u0044\u0052\u004f\u0050\u0044\u004f\u0057\u004e")
	_bfcd._bffef.DdList = _ag.NewCT_FFDDList()
	return _bfcd
}

// SetDoubleStrikeThrough sets the run to double strike-through.
func (_bgbfa RunProperties) SetDoubleStrikeThrough(b bool) {
	if !b {
		_bgbfa._ddda.Dstrike = nil
	} else {
		_bgbfa._ddda.Dstrike = _ag.NewCT_OnOff()
	}
}

// Index returns the index of the footer within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (_cbce Footer) Index() int {
	for _debg, _gabg := range _cbce._dagd._bbc {
		if _gabg == _cbce._edec {
			return _debg
		}
	}
	return -1
}
func _fbd(_daa *_ag.CT_Tbl, _fgbad, _ddfc map[int64]int64) {
	for _, _ccgf := range _daa.EG_ContentRowContent {
		for _, _aegeb := range _ccgf.Tr {
			for _, _cfed := range _aegeb.EG_ContentCellContent {
				for _, _ddbd := range _cfed.Tc {
					for _, _ccbgf := range _ddbd.EG_BlockLevelElts {
						for _, _gbdee := range _ccbgf.EG_ContentBlockContent {
							for _, _ddce := range _gbdee.P {
								_bccgd(_ddce, _fgbad, _ddfc)
							}
							for _, _gac := range _gbdee.Tbl {
								_fbd(_gac, _fgbad, _ddfc)
							}
						}
					}
				}
			}
		}
	}
}

// SetHeadingLevel sets a heading level and style based on the level to a
// paragraph.  The default styles for a new unioffice document support headings
// from level 1 to 8.
func (_cdfec ParagraphProperties) SetHeadingLevel(idx int) {
	_cdfec.SetStyle(_af.Sprintf("\u0048e\u0061\u0064\u0069\u006e\u0067\u0025d", idx))
	if _cdfec._cfdd.NumPr == nil {
		_cdfec._cfdd.NumPr = _ag.NewCT_NumPr()
	}
	_cdfec._cfdd.NumPr.Ilvl = _ag.NewCT_DecimalNumber()
	_cdfec._cfdd.NumPr.Ilvl.ValAttr = int64(idx)
}

// Tables returns the tables defined in the document.
func (_cefb *Document) Tables() []Table {
	_fad := []Table{}
	if _cefb._bbgg.Body == nil {
		return nil
	}
	for _, _ebgc := range _cefb._bbgg.Body.EG_BlockLevelElts {
		for _, _efca := range _ebgc.EG_ContentBlockContent {
			for _, _agccb := range _cefb.tables(_efca) {
				_fad = append(_fad, _agccb)
			}
		}
	}
	return _fad
}

// SetRight sets the right border to a specified type, color and thickness.
func (_acgb CellBorders) SetRight(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_acgb._eaa.Right = _ag.NewCT_Border()
	_ddecc(_acgb._eaa.Right, t, c, thickness)
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (_adb *Document) GetImageByRelID(relID string) (_ec.ImageRef, bool) {
	for _, _bbad := range _adb.Images {
		if _bbad.RelID() == relID {
			return _bbad, true
		}
	}
	return _ec.ImageRef{}, false
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (_agbf ParagraphProperties) Style() string {
	if _agbf._cfdd.PStyle != nil {
		return _agbf._cfdd.PStyle.ValAttr
	}
	return ""
}

// SetHangingIndent controls the indentation of the non-first lines in a paragraph.
func (_dcdg ParagraphProperties) SetHangingIndent(m _bf.Distance) {
	if _dcdg._cfdd.Ind == nil {
		_dcdg._cfdd.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_dcdg._cfdd.Ind.HangingAttr = nil
	} else {
		_dcdg._cfdd.Ind.HangingAttr = &_dc.ST_TwipsMeasure{}
		_dcdg._cfdd.Ind.HangingAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(m / _bf.Twips))
	}
}

// SetMultiLevelType sets the multilevel type.
func (_eaag NumberingDefinition) SetMultiLevelType(t _ag.ST_MultiLevelType) {
	if t == _ag.ST_MultiLevelTypeUnset {
		_eaag._cgfd.MultiLevelType = nil
	} else {
		_eaag._cgfd.MultiLevelType = _ag.NewCT_MultiLevelType()
		_eaag._cgfd.MultiLevelType.ValAttr = t
	}
}

// Italic returns true if paragraph font is italic.
func (_afdgf ParagraphProperties) Italic() bool {
	_baaeg := _afdgf._cfdd.RPr
	return _decg(_baaeg.I) || _decg(_baaeg.ICs)
}

// SetNumberingDefinition sets the numbering definition ID via a NumberingDefinition
// defined in numbering.xml
func (_gccg Paragraph) SetNumberingDefinition(nd NumberingDefinition) {
	_gccg.ensurePPr()
	if _gccg._dbfe.PPr.NumPr == nil {
		_gccg._dbfe.PPr.NumPr = _ag.NewCT_NumPr()
	}
	_dgb := _ag.NewCT_DecimalNumber()
	_egadd := int64(-1)
	for _, _gbce := range _gccg._aefd.Numbering._eefa.Num {
		if _gbce.AbstractNumId != nil && _gbce.AbstractNumId.ValAttr == nd.AbstractNumberID() {
			_egadd = _gbce.NumIdAttr
		}
	}
	if _egadd == -1 {
		_fbbb := _ag.NewCT_Num()
		_gccg._aefd.Numbering._eefa.Num = append(_gccg._aefd.Numbering._eefa.Num, _fbbb)
		_fbbb.NumIdAttr = int64(len(_gccg._aefd.Numbering._eefa.Num))
		_fbbb.AbstractNumId = _ag.NewCT_DecimalNumber()
		_fbbb.AbstractNumId.ValAttr = nd.AbstractNumberID()
	}
	_dgb.ValAttr = _egadd
	_gccg._dbfe.PPr.NumPr.NumId = _dgb
}
func _decg(_bgg *_ag.CT_OnOff) bool { return _bgg != nil }

// X returns the inner wrapped XML type.
func (_cbfd Settings) X() *_ag.Settings { return _cbfd._dcaa }

// SetEndIndent controls the end indentation.
func (_afga ParagraphProperties) SetEndIndent(m _bf.Distance) {
	if _afga._cfdd.Ind == nil {
		_afga._cfdd.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_afga._cfdd.Ind.EndAttr = nil
	} else {
		_afga._cfdd.Ind.EndAttr = &_ag.ST_SignedTwipsMeasure{}
		_afga._cfdd.Ind.EndAttr.Int64 = _f.Int64(int64(m / _bf.Twips))
	}
}
func (_fgd Paragraph) insertRun(_gaag Run, _fcfgf bool) Run {
	for _, _abad := range _fgd._dbfe.EG_PContent {
		for _gbbf, _febd := range _abad.EG_ContentRunContent {
			if _febd.R == _gaag.X() {
				_beee := _ag.NewCT_R()
				_abad.EG_ContentRunContent = append(_abad.EG_ContentRunContent, nil)
				if _fcfgf {
					copy(_abad.EG_ContentRunContent[_gbbf+1:], _abad.EG_ContentRunContent[_gbbf:])
					_abad.EG_ContentRunContent[_gbbf] = _ag.NewEG_ContentRunContent()
					_abad.EG_ContentRunContent[_gbbf].R = _beee
				} else {
					copy(_abad.EG_ContentRunContent[_gbbf+2:], _abad.EG_ContentRunContent[_gbbf+1:])
					_abad.EG_ContentRunContent[_gbbf+1] = _ag.NewEG_ContentRunContent()
					_abad.EG_ContentRunContent[_gbbf+1].R = _beee
				}
				return Run{_fgd._aefd, _beee}
			}
			if _febd.Sdt != nil && _febd.Sdt.SdtContent != nil {
				for _, _fbgb := range _febd.Sdt.SdtContent.EG_ContentRunContent {
					if _fbgb.R == _gaag.X() {
						_cbcg := _ag.NewCT_R()
						_febd.Sdt.SdtContent.EG_ContentRunContent = append(_febd.Sdt.SdtContent.EG_ContentRunContent, nil)
						if _fcfgf {
							copy(_febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf+1:], _febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf:])
							_febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf] = _ag.NewEG_ContentRunContent()
							_febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf].R = _cbcg
						} else {
							copy(_febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf+2:], _febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf+1:])
							_febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf+1] = _ag.NewEG_ContentRunContent()
							_febd.Sdt.SdtContent.EG_ContentRunContent[_gbbf+1].R = _cbcg
						}
						return Run{_fgd._aefd, _cbcg}
					}
				}
			}
		}
	}
	return _fgd.AddRun()
}

// StructuredDocumentTag are a tagged bit of content in a document.
type StructuredDocumentTag struct {
	_eccdf *Document
	_efgce *_ag.CT_SdtBlock
}

// X returns the inner wrapped XML type.
func (_ffadd Table) X() *_ag.CT_Tbl { return _ffadd._fdeda }

// Properties returns the table properties.
func (_bgfbg Table) Properties() TableProperties {
	if _bgfbg._fdeda.TblPr == nil {
		_bgfbg._fdeda.TblPr = _ag.NewCT_TblPr()
	}
	return TableProperties{_bgfbg._fdeda.TblPr}
}
func _eddb(_efgc *_ag.CT_P, _ebaf map[string]string) {
	for _, _fgcd := range _efgc.EG_PContent {
		for _, _beea := range _fgcd.EG_ContentRunContent {
			if _beea.R != nil {
				for _, _eff := range _beea.R.EG_RunInnerContent {
					_eacc := _eff.Drawing
					if _eacc != nil {
						for _, _ggdg := range _eacc.Anchor {
							for _, _cdfa := range _ggdg.Graphic.GraphicData.Any {
								switch _eccdd := _cdfa.(type) {
								case *_afd.Pic:
									if _eccdd.BlipFill != nil && _eccdd.BlipFill.Blip != nil {
										_faefd(_eccdd.BlipFill.Blip, _ebaf)
									}
								default:
								}
							}
						}
						for _, _efdb := range _eacc.Inline {
							for _, _gff := range _efdb.Graphic.GraphicData.Any {
								switch _cfgd := _gff.(type) {
								case *_afd.Pic:
									if _cfgd.BlipFill != nil && _cfgd.BlipFill.Blip != nil {
										_faefd(_cfgd.BlipFill.Blip, _ebaf)
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

// Clear clears the styes.
func (_eadb Styles) Clear() {
	_eadb._faabg.DocDefaults = nil
	_eadb._faabg.LatentStyles = nil
	_eadb._faabg.Style = nil
}

// Bookmark is a bookmarked location within a document that can be referenced
// with a hyperlink.
type Bookmark struct{ _ebg *_ag.CT_Bookmark }

// SetCharacterSpacing sets the run's Character Spacing Adjustment.
func (_geeec RunProperties) SetCharacterSpacing(size _bf.Distance) {
	_geeec._ddda.Spacing = _ag.NewCT_SignedTwipsMeasure()
	_geeec._ddda.Spacing.ValAttr.Int64 = _f.Int64(int64(size / _bf.Twips))
}

// SetValue sets the value of a FormFieldTypeText or FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (_egeg FormField) SetValue(v string) {
	if _egeg._bffef.DdList != nil {
		for _dbag, _bcf := range _egeg.PossibleValues() {
			if _bcf == v {
				_egeg._bffef.DdList.Result = _ag.NewCT_DecimalNumber()
				_egeg._bffef.DdList.Result.ValAttr = int64(_dbag)
				break
			}
		}
	} else if _egeg._bffef.TextInput != nil {
		_egeg._eeef.T = _ag.NewCT_Text()
		_egeg._eeef.T.Content = v
	}
}

// Name returns the name of the field.
func (_abd FormField) Name() string { return *_abd._bffef.Name[0].ValAttr }

// SetTargetByRef sets the URL target of the hyperlink and is more efficient if a link
// destination will be used many times.
func (_ggefc HyperLink) SetTargetByRef(link _ec.Hyperlink) {
	_ggefc._dceb.IdAttr = _f.String(_ec.Relationship(link).ID())
	_ggefc._dceb.AnchorAttr = nil
}
func _aabb(_dcgd string) mergeFieldInfo {
	_gdcb := []string{}
	_ceec := _ca.Buffer{}
	_gdge := -1
	for _cfac, _eabd := range _dcgd {
		switch _eabd {
		case ' ':
			if _ceec.Len() != 0 {
				_gdcb = append(_gdcb, _ceec.String())
			}
			_ceec.Reset()
		case '"':
			if _gdge != -1 {
				_gdcb = append(_gdcb, _dcgd[_gdge+1:_cfac])
				_gdge = -1
			} else {
				_gdge = _cfac
			}
		default:
			_ceec.WriteRune(_eabd)
		}
	}
	if _ceec.Len() != 0 {
		_gdcb = append(_gdcb, _ceec.String())
	}
	_gadg := mergeFieldInfo{}
	for _cecbe := 0; _cecbe < len(_gdcb)-1; _cecbe++ {
		_cbfg := _gdcb[_cecbe]
		switch _cbfg {
		case "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044":
			_gadg._cabg = _gdcb[_cecbe+1]
			_cecbe++
		case "\u005c\u0066":
			_gadg._ebce = _gdcb[_cecbe+1]
			_cecbe++
		case "\u005c\u0062":
			_gadg._gbeg = _gdcb[_cecbe+1]
			_cecbe++
		case "\u005c\u002a":
			switch _gdcb[_cecbe+1] {
			case "\u0055\u0070\u0070e\u0072":
				_gadg._ebe = true
			case "\u004c\u006f\u0077e\u0072":
				_gadg._gdad = true
			case "\u0043\u0061\u0070\u0073":
				_gadg._fcd = true
			case "\u0046\u0069\u0072\u0073\u0074\u0043\u0061\u0070":
				_gadg._efad = true
			}
			_cecbe++
		}
	}
	return _gadg
}

// SetStart sets the cell start margin
func (_ega CellMargins) SetStart(d _bf.Distance) {
	_ega._cacb.Start = _ag.NewCT_TblWidth()
	_edf(_ega._cacb.Start, d)
}

// AddCell adds a cell to a row and returns it
func (_dgbg Row) AddCell() Cell {
	_ggea := _ag.NewEG_ContentCellContent()
	_dgbg._dffc.EG_ContentCellContent = append(_dgbg._dffc.EG_ContentCellContent, _ggea)
	_gdca := _ag.NewCT_Tc()
	_ggea.Tc = append(_ggea.Tc, _gdca)
	return Cell{_dgbg._eefac, _gdca}
}

// SetHANSITheme sets the font H ANSI Theme.
func (_ddcb Fonts) SetHANSITheme(t _ag.ST_Theme) { _ddcb._agaa.HAnsiThemeAttr = t }

// SetHighlight highlights text in a specified color.
func (_eeaa RunProperties) SetHighlight(c _ag.ST_HighlightColor) {
	_eeaa._ddda.Highlight = _ag.NewCT_Highlight()
	_eeaa._ddda.Highlight.ValAttr = c
}

// Italic returns true if run font is italic.
func (_ecee RunProperties) Italic() bool {
	_dbgfe := _ecee._ddda
	return _decg(_dbgfe.I) || _decg(_dbgfe.ICs)
}

// SetEffect sets a text effect on the run.
func (_gcef RunProperties) SetEffect(e _ag.ST_TextEffect) {
	if e == _ag.ST_TextEffectUnset {
		_gcef._ddda.Effect = nil
	} else {
		_gcef._ddda.Effect = _ag.NewCT_TextEffect()
		_gcef._ddda.Effect.ValAttr = _ag.ST_TextEffectShimmer
	}
}

// SetAfter sets the spacing that comes after the paragraph.
func (_dgff ParagraphSpacing) SetAfter(after _bf.Distance) {
	_dgff._gggda.AfterAttr = &_dc.ST_TwipsMeasure{}
	_dgff._gggda.AfterAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(after / _bf.Twips))
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (_eecd ParagraphProperties) CharacterSpacingMeasure() string {
	if _agefa := _eecd._cfdd.RPr.Spacing; _agefa != nil {
		_bcbd := _agefa.ValAttr
		if _bcbd.ST_UniversalMeasure != nil {
			return *_bcbd.ST_UniversalMeasure
		}
	}
	return ""
}

// Style is a style within the styles.xml file.
type Style struct{ _gade *_ag.CT_Style }

// AddText adds tet to a run.
func (_ebab Run) AddText(s string) {
	_cgef := _ag.NewEG_RunInnerContent()
	_ebab._fegg.EG_RunInnerContent = append(_ebab._fegg.EG_RunInnerContent, _cgef)
	_cgef.T = _ag.NewCT_Text()
	if _f.NeedsSpacePreserve(s) {
		_ccfa := "\u0070\u0072\u0065\u0073\u0065\u0072\u0076\u0065"
		_cgef.T.SpaceAttr = &_ccfa
	}
	_cgef.T.Content = s
}

// SetFirstRow controls the conditional formatting for the first row in a table.
func (_dbaf TableLook) SetFirstRow(on bool) {
	if !on {
		_dbaf._cfbff.FirstRowAttr = &_dc.ST_OnOff{}
		_dbaf._cfbff.FirstRowAttr.ST_OnOff1 = _dc.ST_OnOff1Off
	} else {
		_dbaf._cfbff.FirstRowAttr = &_dc.ST_OnOff{}
		_dbaf._cfbff.FirstRowAttr.ST_OnOff1 = _dc.ST_OnOff1On
	}
}

// HasFootnotes returns a bool based on the presence or abscence of footnotes within
// the document.
func (_cee *Document) HasFootnotes() bool { return _cee._ab != nil }

// VerticalAlign returns the value of paragraph vertical align.
func (_agca ParagraphProperties) VerticalAlignment() _dc.ST_VerticalAlignRun {
	if _bege := _agca._cfdd.RPr.VertAlign; _bege != nil {
		return _bege.ValAttr
	}
	return 0
}

// HyperLink is a link within a document.
type HyperLink struct {
	_debd *Document
	_dceb *_ag.CT_Hyperlink
}

// ParagraphStyleProperties is the styling information for a paragraph.
type ParagraphStyleProperties struct{ _ddde *_ag.CT_PPrGeneral }

// Clear resets the numbering.
func (_ebgdg Numbering) Clear() {
	_ebgdg._eefa.AbstractNum = nil
	_ebgdg._eefa.Num = nil
	_ebgdg._eefa.NumIdMacAtCleanup = nil
	_ebgdg._eefa.NumPicBullet = nil
}

// SetChecked marks a FormFieldTypeCheckBox as checked or unchecked.
func (_acec FormField) SetChecked(b bool) {
	if _acec._bffef.CheckBox == nil {
		return
	}
	if !b {
		_acec._bffef.CheckBox.Checked = nil
	} else {
		_acec._bffef.CheckBox.Checked = _ag.NewCT_OnOff()
	}
}

// SetStyle sets the table style name.
func (_dafge TableProperties) SetStyle(name string) {
	if name == "" {
		_dafge._gfaa.TblStyle = nil
	} else {
		_dafge._gfaa.TblStyle = _ag.NewCT_String()
		_dafge._gfaa.TblStyle.ValAttr = name
	}
}

// SetSize sets the size of the displayed image on the page.
func (_eag AnchoredDrawing) SetSize(w, h _bf.Distance) {
	_eag._de.Extent.CxAttr = int64(float64(w*_bf.Pixel72) / _bf.EMU)
	_eag._de.Extent.CyAttr = int64(float64(h*_bf.Pixel72) / _bf.EMU)
}

// Emboss returns true if run emboss is on.
func (_fbfd RunProperties) Emboss() bool { return _decg(_fbfd._ddda.Emboss) }

// HasEndnotes returns a bool based on the presence or abscence of endnotes within
// the document.
func (_dfe *Document) HasEndnotes() bool { return _dfe._fcb != nil }
func _bccgd(_gfda *_ag.CT_P, _cgb, _faca map[int64]int64) {
	for _, _bcgg := range _gfda.EG_PContent {
		for _, _gag := range _bcgg.EG_ContentRunContent {
			if _gag.R != nil {
				for _, _eagbb := range _gag.R.EG_RunInnerContent {
					_bgd := _eagbb.EndnoteReference
					if _bgd != nil && _bgd.IdAttr > 0 {
						if _egge, _dfaae := _faca[_bgd.IdAttr]; _dfaae {
							_bgd.IdAttr = _egge
						}
					}
					_defa := _eagbb.FootnoteReference
					if _defa != nil && _defa.IdAttr > 0 {
						if _ffeg, _aege := _cgb[_defa.IdAttr]; _aege {
							_defa.IdAttr = _ffeg
						}
					}
				}
			}
		}
	}
}

// BodySection returns the default body section used for all preceding
// paragraphs until the previous Section. If there is no previous sections, the
// body section applies to the entire document.
func (_cga *Document) BodySection() Section {
	if _cga._bbgg.Body.SectPr == nil {
		_cga._bbgg.Body.SectPr = _ag.NewCT_SectPr()
	}
	return Section{_cga, _cga._bbgg.Body.SectPr}
}

// X returns the inner wrapped XML type.
func (_gaaf ParagraphProperties) X() *_ag.CT_PPr { return _gaaf._cfdd }

// InsertRunBefore inserts a run in the paragraph before the relative run.
func (_dfgb Paragraph) InsertRunBefore(relativeTo Run) Run { return _dfgb.insertRun(relativeTo, true) }

// Endnotes returns the endnotes defined in the document.
func (_dbfb *Document) Endnotes() []Endnote {
	_eda := []Endnote{}
	for _, _eed := range _dbfb._fcb.CT_Endnotes.Endnote {
		_eda = append(_eda, Endnote{_dbfb, _eed})
	}
	return _eda
}

// AddFooter creates a Footer associated with the document, but doesn't add it
// to the document for display.
func (_deab *Document) AddFooter() Footer {
	_fde := _ag.NewFtr()
	_deab._bbc = append(_deab._bbc, _fde)
	_egac := _af.Sprintf("\u0066\u006f\u006ft\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", len(_deab._bbc))
	_deab._agcc.AddRelationship(_egac, _f.FooterType)
	_deab.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_egac, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0066\u006f\u006f\u0074e\u0072\u002b\u0078\u006d\u006c")
	_deab._dbc = append(_deab._dbc, _ec.NewRelationships())
	return Footer{_deab, _fde}
}

// SaveToFile writes the document out to a file.
func (_fdfd *Document) SaveToFile(path string) error {
	_bfa, _ece := _d.Create(path)
	if _ece != nil {
		return _ece
	}
	defer _bfa.Close()
	return _fdfd.Save(_bfa)
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (_gfef ParagraphProperties) RStyle() string {
	if _gfef._cfdd.RPr.RStyle != nil {
		return _gfef._cfdd.RPr.RStyle.ValAttr
	}
	return ""
}

// UnderlineColor returns the hex color value of paragraph underline.
func (_cgcc ParagraphProperties) UnderlineColor() string {
	if _abeg := _cgcc._cfdd.RPr.U; _abeg != nil {
		_cfgf := _abeg.ColorAttr
		if _cfgf != nil && _cfgf.ST_HexColorRGB != nil {
			return *_cfgf.ST_HexColorRGB
		}
	}
	return ""
}
func (_cd *Document) tables(_eaf *_ag.EG_ContentBlockContent) []Table {
	_bfc := []Table{}
	for _, _bac := range _eaf.Tbl {
		_bfc = append(_bfc, Table{_cd, _bac})
		for _, _dbcb := range _bac.EG_ContentRowContent {
			for _, _cff := range _dbcb.Tr {
				for _, _gaf := range _cff.EG_ContentCellContent {
					for _, _cfef := range _gaf.Tc {
						for _, _afbg := range _cfef.EG_BlockLevelElts {
							for _, _ebb := range _afbg.EG_ContentBlockContent {
								for _, _fca := range _cd.tables(_ebb) {
									_bfc = append(_bfc, _fca)
								}
							}
						}
					}
				}
			}
		}
	}
	return _bfc
}

// SetSpacing sets the spacing that comes before and after the paragraph.
// Deprecated: See Spacing() instead which allows finer control.
func (_bcac ParagraphProperties) SetSpacing(before, after _bf.Distance) {
	if _bcac._cfdd.Spacing == nil {
		_bcac._cfdd.Spacing = _ag.NewCT_Spacing()
	}
	_bcac._cfdd.Spacing.BeforeAttr = &_dc.ST_TwipsMeasure{}
	_bcac._cfdd.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(before / _bf.Twips))
	_bcac._cfdd.Spacing.AfterAttr = &_dc.ST_TwipsMeasure{}
	_bcac._cfdd.Spacing.AfterAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(after / _bf.Twips))
}

// X returns the inner wrapped XML type.
func (_fbaf Header) X() *_ag.Hdr { return _fbaf._aeaf }

// SetLeftIndent controls the left indent of the paragraph.
func (_eeee ParagraphStyleProperties) SetLeftIndent(m _bf.Distance) {
	if _eeee._ddde.Ind == nil {
		_eeee._ddde.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_eeee._ddde.Ind.LeftAttr = nil
	} else {
		_eeee._ddde.Ind.LeftAttr = &_ag.ST_SignedTwipsMeasure{}
		_eeee._ddde.Ind.LeftAttr.Int64 = _f.Int64(int64(m / _bf.Twips))
	}
}

// NewStyles constructs a new empty Styles
func NewStyles() Styles { return Styles{_ag.NewStyles()} }

// Paragraphs returns the paragraphs defined in a footnote.
func (_fdag Footnote) Paragraphs() []Paragraph {
	_dcg := []Paragraph{}
	for _, _bbbe := range _fdag.content() {
		for _, _gegc := range _bbbe.P {
			_dcg = append(_dcg, Paragraph{_fdag._gagb, _gegc})
		}
	}
	return _dcg
}

// Cell is a table cell within a document (not a spreadsheet)
type Cell struct {
	_dce *Document
	_afg *_ag.CT_Tc
}

// X returns the inner wrapped XML type.
func (_caed Numbering) X() *_ag.Numbering { return _caed._eefa }

// Validate validates the structure and in cases where it't possible, the ranges
// of elements within a document. A validation error dones't mean that the
// document won't work in MS Word or LibreOffice, but it's worth checking into.
func (_fbcda *Document) Validate() error {
	if _fbcda == nil || _fbcda._bbgg == nil {
		return _g.New("\u0064o\u0063\u0075m\u0065\u006e\u0074\u0020n\u006f\u0074\u0020i\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0065d \u0063\u006f\u0072r\u0065\u0063t\u006c\u0079\u002c\u0020\u006e\u0069l\u0020\u0062a\u0073\u0065")
	}
	for _, _ccaa := range []func() error{_fbcda.validateTableCells, _fbcda.validateBookmarks} {
		if _bba := _ccaa(); _bba != nil {
			return _bba
		}
	}
	if _ecce := _fbcda._bbgg.Validate(); _ecce != nil {
		return _ecce
	}
	return nil
}
func (_dfcb Paragraph) addEndFldChar() *_ag.CT_FldChar {
	_bgaf := _dfcb.addFldChar()
	_bgaf.FldCharTypeAttr = _ag.ST_FldCharTypeEnd
	return _bgaf
}

// RemoveEndnote removes a endnote from both the paragraph and the document
// the requested endnote must be anchored on the paragraph being referenced.
func (_caab Paragraph) RemoveEndnote(id int64) {
	_afe := _caab._aefd._fcb
	var _eebee int
	for _bcda, _cfedg := range _afe.CT_Endnotes.Endnote {
		if _cfedg.IdAttr == id {
			_eebee = _bcda
		}
	}
	_eebee = 0
	_afe.CT_Endnotes.Endnote[_eebee] = nil
	_afe.CT_Endnotes.Endnote[_eebee] = _afe.CT_Endnotes.Endnote[len(_afe.CT_Endnotes.Endnote)-1]
	_afe.CT_Endnotes.Endnote = _afe.CT_Endnotes.Endnote[:len(_afe.CT_Endnotes.Endnote)-1]
	var _fecda Run
	for _, _fcgg := range _caab.Runs() {
		if _fded, _dgde := _fcgg.IsEndnote(); _fded {
			if _dgde == id {
				_fecda = _fcgg
			}
		}
	}
	_caab.RemoveRun(_fecda)
}

// Settings controls the document settings.
type Settings struct{ _dcaa *_ag.Settings }

func (_afcb Paragraph) addStartBookmark(_bbggag int64, _ffgff string) *_ag.CT_Bookmark {
	_aebf := _ag.NewEG_PContent()
	_afcb._dbfe.EG_PContent = append(_afcb._dbfe.EG_PContent, _aebf)
	_dabec := _ag.NewEG_ContentRunContent()
	_bfbff := _ag.NewEG_RunLevelElts()
	_bdff := _ag.NewEG_RangeMarkupElements()
	_eedc := _ag.NewCT_Bookmark()
	_eedc.NameAttr = _ffgff
	_eedc.IdAttr = _bbggag
	_bdff.BookmarkStart = _eedc
	_aebf.EG_ContentRunContent = append(_aebf.EG_ContentRunContent, _dabec)
	_dabec.EG_RunLevelElts = append(_dabec.EG_RunLevelElts, _bfbff)
	_bfbff.EG_RangeMarkupElements = append(_bfbff.EG_RangeMarkupElements, _bdff)
	return _eedc
}

// Underline returns the type of run underline.
func (_ffgaa RunProperties) Underline() _ag.ST_Underline {
	if _fdgbe := _ffgaa._ddda.U; _fdgbe != nil {
		return _fdgbe.ValAttr
	}
	return 0
}

// SetLeft sets the cell left margin
func (_ada CellMargins) SetLeft(d _bf.Distance) {
	_ada._cacb.Left = _ag.NewCT_TblWidth()
	_edf(_ada._cacb.Left, d)
}

// SetHorizontalBanding controls the conditional formatting for horizontal banding.
func (_agdc TableLook) SetHorizontalBanding(on bool) {
	if !on {
		_agdc._cfbff.NoHBandAttr = &_dc.ST_OnOff{}
		_agdc._cfbff.NoHBandAttr.ST_OnOff1 = _dc.ST_OnOff1On
	} else {
		_agdc._cfbff.NoHBandAttr = &_dc.ST_OnOff{}
		_agdc._cfbff.NoHBandAttr.ST_OnOff1 = _dc.ST_OnOff1Off
	}
}

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

// AddTable adds a new table to the document body.
func (_cfe *Document) AddTable() Table {
	_fbcc := _ag.NewEG_BlockLevelElts()
	_cfe._bbgg.Body.EG_BlockLevelElts = append(_cfe._bbgg.Body.EG_BlockLevelElts, _fbcc)
	_efc := _ag.NewEG_ContentBlockContent()
	_fbcc.EG_ContentBlockContent = append(_fbcc.EG_ContentBlockContent, _efc)
	_efb := _ag.NewCT_Tbl()
	_efc.Tbl = append(_efc.Tbl, _efb)
	return Table{_cfe, _efb}
}

// SetAllCaps sets the run to all caps.
func (_acfde RunProperties) SetAllCaps(b bool) {
	if !b {
		_acfde._ddda.Caps = nil
	} else {
		_acfde._ddda.Caps = _ag.NewCT_OnOff()
	}
}

// AddStyle adds a new empty style.
func (_aafc Styles) AddStyle(styleID string, t _ag.ST_StyleType, isDefault bool) Style {
	_befg := _ag.NewCT_Style()
	_befg.TypeAttr = t
	if isDefault {
		_befg.DefaultAttr = &_dc.ST_OnOff{}
		_befg.DefaultAttr.Bool = _f.Bool(isDefault)
	}
	_befg.StyleIdAttr = _f.String(styleID)
	_aafc._faabg.Style = append(_aafc._faabg.Style, _befg)
	return Style{_befg}
}

// Document is a text document that can be written out in the OOXML .docx
// format. It can be opened from a file on disk and modified, or created from
// scratch.
type Document struct {
	_ec.DocBase
	_bbgg     *_ag.Document
	Settings  Settings
	Numbering Numbering
	Styles    Styles
	_cg       []*_ag.Hdr
	_gba      []_ec.Relationships
	_bbc      []*_ag.Ftr
	_dbc      []_ec.Relationships
	_agcc     _ec.Relationships
	_adae     []*_ea.Theme
	_cef      *_ag.WebSettings
	_gc       *_ag.Fonts
	_fcb      *_ag.Endnotes
	_ab       *_ag.Footnotes
}

// AddPageBreak adds a page break to a run.
func (_eace Run) AddPageBreak() {
	_gabb := _eace.newIC()
	_gabb.Br = _ag.NewCT_Br()
	_gabb.Br.TypeAttr = _ag.ST_BrTypePage
}

// SetASCIITheme sets the font ASCII Theme.
func (_cgc Fonts) SetASCIITheme(t _ag.ST_Theme) { _cgc._agaa.AsciiThemeAttr = t }

// Properties returns the row properties.
func (_fbgg Row) Properties() RowProperties {
	if _fbgg._dffc.TrPr == nil {
		_fbgg._dffc.TrPr = _ag.NewCT_TrPr()
	}
	return RowProperties{_fbgg._dffc.TrPr}
}

// AbstractNumberID returns the ID that is unique within all numbering
// definitions that is used to assign the definition to a paragraph.
func (_ceca NumberingDefinition) AbstractNumberID() int64 { return _ceca._cgfd.AbstractNumIdAttr }

// NumberingDefinition defines a numbering definition for a list of pragraphs.
type NumberingDefinition struct{ _cgfd *_ag.CT_AbstractNum }

// SetSemiHidden controls if the style is hidden in the UI.
func (_gfee Style) SetSemiHidden(b bool) {
	if b {
		_gfee._gade.SemiHidden = _ag.NewCT_OnOff()
	} else {
		_gfee._gade.SemiHidden = nil
	}
}

// SetPossibleValues sets possible values for a FormFieldTypeDropDown.
func (_fcfa FormField) SetPossibleValues(values []string) {
	if _fcfa._bffef.DdList != nil {
		for _, _bbcd := range values {
			_fbae := _ag.NewCT_String()
			_fbae.ValAttr = _bbcd
			_fcfa._bffef.DdList.ListEntry = append(_fcfa._bffef.DdList.ListEntry, _fbae)
		}
	}
}

// OnOffValue represents an on/off value that can also be unset
type OnOffValue byte

// ComplexSizeValue returns the value of run font size for complex fonts in points.
func (_daecc RunProperties) ComplexSizeValue() float64 {
	if _edbc := _daecc._ddda.SzCs; _edbc != nil {
		_eeaaf := _edbc.ValAttr
		if _eeaaf.ST_UnsignedDecimalNumber != nil {
			return float64(*_eeaaf.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// Rows returns the rows defined in the table.
func (_ggge Table) Rows() []Row {
	_gacbb := []Row{}
	for _, _aeggb := range _ggge._fdeda.EG_ContentRowContent {
		for _, _bagee := range _aeggb.Tr {
			_gacbb = append(_gacbb, Row{_ggge._aebg, _bagee})
		}
		if _aeggb.Sdt != nil && _aeggb.Sdt.SdtContent != nil {
			for _, _dggbe := range _aeggb.Sdt.SdtContent.Tr {
				_gacbb = append(_gacbb, Row{_ggge._aebg, _dggbe})
			}
		}
	}
	return _gacbb
}

// SetWidthPercent sets the table to a width percentage.
func (_adgdae TableProperties) SetWidthPercent(pct float64) {
	_adgdae._gfaa.TblW = _ag.NewCT_TblWidth()
	_adgdae._gfaa.TblW.TypeAttr = _ag.ST_TblWidthPct
	_adgdae._gfaa.TblW.WAttr = &_ag.ST_MeasurementOrPercent{}
	_adgdae._gfaa.TblW.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_adgdae._gfaa.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(pct * 50))
}

// ItalicValue returns the precise nature of the italic setting (unset, off or on).
func (_acag RunProperties) ItalicValue() OnOffValue { return _eea(_acag._ddda.I) }

// AddParagraph adds a paragraph to the table cell.
func (_faf Cell) AddParagraph() Paragraph {
	_cc := _ag.NewEG_BlockLevelElts()
	_faf._afg.EG_BlockLevelElts = append(_faf._afg.EG_BlockLevelElts, _cc)
	_egg := _ag.NewEG_ContentBlockContent()
	_cc.EG_ContentBlockContent = append(_cc.EG_ContentBlockContent, _egg)
	_fc := _ag.NewCT_P()
	_egg.P = append(_egg.P, _fc)
	return Paragraph{_faf._dce, _fc}
}

// Properties returns the paragraph properties.
func (_efbd Paragraph) Properties() ParagraphProperties {
	_efbd.ensurePPr()
	return ParagraphProperties{_efbd._aefd, _efbd._dbfe.PPr}
}

// Bold returns true if paragraph font is bold.
func (_abgd ParagraphProperties) Bold() bool {
	_gcaf := _abgd._cfdd.RPr
	return _decg(_gcaf.B) || _decg(_gcaf.BCs)
}

// SetStyle sets the font size.
func (_efcdg RunProperties) SetStyle(style string) {
	if style == "" {
		_efcdg._ddda.RStyle = nil
	} else {
		_efcdg._ddda.RStyle = _ag.NewCT_String()
		_efcdg._ddda.RStyle.ValAttr = style
	}
}

// X returns the inner wrapped XML type.
func (_dbce Run) X() *_ag.CT_R { return _dbce._fegg }

// SetLinkedStyle sets the style that this style is linked to.
func (_fbbea Style) SetLinkedStyle(name string) {
	if name == "" {
		_fbbea._gade.Link = nil
	} else {
		_fbbea._gade.Link = _ag.NewCT_String()
		_fbbea._gade.Link.ValAttr = name
	}
}

// X returns the inner wrapped XML type.
func (_faab Footnote) X() *_ag.CT_FtnEdn { return _faab._faaeg }

// InsertParagraphAfter adds a new empty paragraph after the relativeTo
// paragraph.
func (_fgce *Document) InsertParagraphAfter(relativeTo Paragraph) Paragraph {
	return _fgce.insertParagraph(relativeTo, false)
}

// SetThemeShade sets the shade based off the theme color.
func (_daf Color) SetThemeShade(s uint8) {
	_aaa := _af.Sprintf("\u0025\u0030\u0032\u0078", s)
	_daf._ccc.ThemeShadeAttr = &_aaa
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (_fcdb TableStyleProperties) SetCellSpacingPercent(pct float64) {
	_fcdb._cgcba.TblCellSpacing = _ag.NewCT_TblWidth()
	_fcdb._cgcba.TblCellSpacing.TypeAttr = _ag.ST_TblWidthPct
	_fcdb._cgcba.TblCellSpacing.WAttr = &_ag.ST_MeasurementOrPercent{}
	_fcdb._cgcba.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_fcdb._cgcba.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(pct * 50))
}

// ParagraphStyles returns only the paragraph styles.
func (_bebab Styles) ParagraphStyles() []Style {
	_eaaac := []Style{}
	for _, _daff := range _bebab._faabg.Style {
		if _daff.TypeAttr != _ag.ST_StyleTypeParagraph {
			continue
		}
		_eaaac = append(_eaaac, Style{_daff})
	}
	return _eaaac
}

// SetTopPct sets the cell top margin
func (_bff CellMargins) SetTopPct(pct float64) {
	_bff._cacb.Top = _ag.NewCT_TblWidth()
	_adg(_bff._cacb.Top, pct)
}

// SetOrigin sets the origin of the image.  It defaults to ST_RelFromHPage and
// ST_RelFromVPage
func (_aae AnchoredDrawing) SetOrigin(h _ag.WdST_RelFromH, v _ag.WdST_RelFromV) {
	_aae._de.PositionH.RelativeFromAttr = h
	_aae._de.PositionV.RelativeFromAttr = v
}

// X returns the inner wrapped XML type.
func (_bcbe Fonts) X() *_ag.CT_Fonts { return _bcbe._agaa }

// SetStartIndent controls the start indentation.
func (_ggfe ParagraphProperties) SetStartIndent(m _bf.Distance) {
	if _ggfe._cfdd.Ind == nil {
		_ggfe._cfdd.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_ggfe._cfdd.Ind.StartAttr = nil
	} else {
		_ggfe._cfdd.Ind.StartAttr = &_ag.ST_SignedTwipsMeasure{}
		_ggfe._cfdd.Ind.StartAttr.Int64 = _f.Int64(int64(m / _bf.Twips))
	}
}

// UnderlineColor returns the hex color value of run underline.
func (_aabg RunProperties) UnderlineColor() string {
	if _gccfa := _aabg._ddda.U; _gccfa != nil {
		_dcef := _gccfa.ColorAttr
		if _dcef != nil && _dcef.ST_HexColorRGB != nil {
			return *_dcef.ST_HexColorRGB
		}
	}
	return ""
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_cede TableProperties) SetCellSpacingAuto() {
	_cede._gfaa.TblCellSpacing = _ag.NewCT_TblWidth()
	_cede._gfaa.TblCellSpacing.TypeAttr = _ag.ST_TblWidthAuto
}

// SetEastAsiaTheme sets the font East Asia Theme.
func (_fced Fonts) SetEastAsiaTheme(t _ag.ST_Theme) { _fced._agaa.EastAsiaThemeAttr = t }

// X returns the inner wrapped XML type.
func (_dcebc Paragraph) X() *_ag.CT_P { return _dcebc._dbfe }

// ParagraphProperties returns the paragraph style properties.
func (_fdd Style) ParagraphProperties() ParagraphStyleProperties {
	if _fdd._gade.PPr == nil {
		_fdd._gade.PPr = _ag.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_fdd._gade.PPr}
}

// ComplexSizeValue returns the value of paragraph font size for complex fonts in points.
func (_cbcgc ParagraphProperties) ComplexSizeValue() float64 {
	if _bdfg := _cbcgc._cfdd.RPr.SzCs; _bdfg != nil {
		_fcag := _bdfg.ValAttr
		if _fcag.ST_UnsignedDecimalNumber != nil {
			return float64(*_fcag.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// Borders allows controlling individual cell borders.
func (_bad CellProperties) Borders() CellBorders {
	if _bad._dac.TcBorders == nil {
		_bad._dac.TcBorders = _ag.NewCT_TcBorders()
	}
	return CellBorders{_bad._dac.TcBorders}
}

// Runs returns all of the runs in a paragraph.
func (_ddee Paragraph) Runs() []Run {
	_ebff := []Run{}
	for _, _gggd := range _ddee._dbfe.EG_PContent {
		for _, _cbea := range _gggd.EG_ContentRunContent {
			if _cbea.R != nil {
				_ebff = append(_ebff, Run{_ddee._aefd, _cbea.R})
			}
			if _cbea.Sdt != nil && _cbea.Sdt.SdtContent != nil {
				for _, _eaaga := range _cbea.Sdt.SdtContent.EG_ContentRunContent {
					if _eaaga.R != nil {
						_ebff = append(_ebff, Run{_ddee._aefd, _eaaga.R})
					}
				}
			}
		}
	}
	return _ebff
}

// SetVerticalBanding controls the conditional formatting for vertical banding.
func (_ccdc TableLook) SetVerticalBanding(on bool) {
	if !on {
		_ccdc._cfbff.NoVBandAttr = &_dc.ST_OnOff{}
		_ccdc._cfbff.NoVBandAttr.ST_OnOff1 = _dc.ST_OnOff1On
	} else {
		_ccdc._cfbff.NoVBandAttr = &_dc.ST_OnOff{}
		_ccdc._cfbff.NoVBandAttr.ST_OnOff1 = _dc.ST_OnOff1Off
	}
}

// AddSection adds a new document section with an optional section break.  If t
// is ST_SectionMarkUnset, then no break will be inserted.
func (_affa ParagraphProperties) AddSection(t _ag.ST_SectionMark) Section {
	_affa._cfdd.SectPr = _ag.NewCT_SectPr()
	if t != _ag.ST_SectionMarkUnset {
		_affa._cfdd.SectPr.Type = _ag.NewCT_SectType()
		_affa._cfdd.SectPr.Type.ValAttr = t
	}
	return Section{_affa._bgfc, _affa._cfdd.SectPr}
}

// X returns the inner wrapped XML type.
func (_badd HyperLink) X() *_ag.CT_Hyperlink { return _badd._dceb }

// AddRun adds a run to a paragraph.
func (_afaa Paragraph) AddRun() Run {
	_gced := _ag.NewEG_PContent()
	_afaa._dbfe.EG_PContent = append(_afaa._dbfe.EG_PContent, _gced)
	_ecbd := _ag.NewEG_ContentRunContent()
	_gced.EG_ContentRunContent = append(_gced.EG_ContentRunContent, _ecbd)
	_gegcc := _ag.NewCT_R()
	_ecbd.R = _gegcc
	return Run{_afaa._aefd, _gegcc}
}

// FormFields extracts all of the fields from a document.  They can then be
// manipulated via the methods on the field and the document saved.
func (_fcc *Document) FormFields() []FormField {
	_dabf := []FormField{}
	for _, _dgfd := range _fcc.Paragraphs() {
		_cgec := _dgfd.Runs()
		for _fccg, _bcaf := range _cgec {
			for _, _caa := range _bcaf._fegg.EG_RunInnerContent {
				if _caa.FldChar == nil || _caa.FldChar.FfData == nil {
					continue
				}
				if _caa.FldChar.FldCharTypeAttr == _ag.ST_FldCharTypeBegin {
					if len(_caa.FldChar.FfData.Name) == 0 || _caa.FldChar.FfData.Name[0].ValAttr == nil {
						continue
					}
					_debcd := FormField{_bffef: _caa.FldChar.FfData}
					if _caa.FldChar.FfData.TextInput != nil {
						for _ecg := _fccg + 1; _ecg < len(_cgec)-1; _ecg++ {
							if len(_cgec[_ecg]._fegg.EG_RunInnerContent) == 0 {
								continue
							}
							_geb := _cgec[_ecg]._fegg.EG_RunInnerContent[0]
							if _geb.FldChar != nil && _geb.FldChar.FldCharTypeAttr == _ag.ST_FldCharTypeSeparate {
								if len(_cgec[_ecg+1]._fegg.EG_RunInnerContent) == 0 {
									continue
								}
								if _cgec[_ecg+1]._fegg.EG_RunInnerContent[0].FldChar == nil {
									_debcd._eeef = _cgec[_ecg+1]._fegg.EG_RunInnerContent[0]
									break
								}
							}
						}
					}
					_dabf = append(_dabf, _debcd)
				}
			}
		}
	}
	return _dabf
}

// RemoveParagraph removes a paragraph from the footnote.
func (_dbggf Footnote) RemoveParagraph(p Paragraph) {
	for _, _cab := range _dbggf.content() {
		for _edbg, _dbb := range _cab.P {
			if _dbb == p._dbfe {
				copy(_cab.P[_edbg:], _cab.P[_edbg+1:])
				_cab.P = _cab.P[0 : len(_cab.P)-1]
				return
			}
		}
	}
}

// CellMargins are the margins for an individual cell.
type CellMargins struct{ _cacb *_ag.CT_TcMar }

// AddBreak adds a line break to a run.
func (_cgebb Run) AddBreak() { _gabd := _cgebb.newIC(); _gabd.Br = _ag.NewCT_Br() }

// X returns the inner wrapped XML type.
func (_bddd Endnote) X() *_ag.CT_FtnEdn { return _bddd._gae }

// SetColumnBandSize sets the number of Columns in the column band
func (_gbege TableStyleProperties) SetColumnBandSize(cols int64) {
	_gbege._cgcba.TblStyleColBandSize = _ag.NewCT_DecimalNumber()
	_gbege._cgcba.TblStyleColBandSize.ValAttr = cols
}

// IsFootnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (_ecda Run) IsFootnote() (bool, int64) {
	if _ecda._fegg.EG_RunInnerContent != nil {
		if _ecda._fegg.EG_RunInnerContent[0].FootnoteReference != nil {
			return true, _ecda._fegg.EG_RunInnerContent[0].FootnoteReference.IdAttr
		}
	}
	return false, 0
}

// AddTab adds tab to a run and can be used with the the Paragraph's tab stops.
func (_bgdgb Run) AddTab() { _ffga := _bgdgb.newIC(); _ffga.Tab = _ag.NewCT_Empty() }

// Save writes the document to an io.Writer in the Zip package format.
func (_fbc *Document) Save(w _cb.Writer) error {
	if _dcc := _fbc._bbgg.Validate(); _dcc != nil {
		_f.Log("\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u0020i\u006e\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0073", _dcc)
	}
	_bbce := _f.DocTypeDocument
	if !_eb.GetLicenseKey().IsLicensed() && !_bfb {
		_af.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
		_af.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _g.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	}
	_fefc := _fb.NewWriter(w)
	defer _fefc.Close()
	if _ccf := _bab.MarshalXML(_fefc, _f.BaseRelsFilename, _fbc.Rels.X()); _ccf != nil {
		return _ccf
	}
	if _edfb := _bab.MarshalXMLByType(_fefc, _bbce, _f.ExtendedPropertiesType, _fbc.AppProperties.X()); _edfb != nil {
		return _edfb
	}
	if _age := _bab.MarshalXMLByType(_fefc, _bbce, _f.CorePropertiesType, _fbc.CoreProperties.X()); _age != nil {
		return _age
	}
	if _fbc.CustomProperties.X() != nil {
		if _ee := _bab.MarshalXMLByType(_fefc, _bbce, _f.CustomPropertiesType, _fbc.CustomProperties.X()); _ee != nil {
			return _ee
		}
	}
	if _fbc.Thumbnail != nil {
		_edg, _ga := _fefc.Create("\u0064\u006f\u0063Pr\u006f\u0070\u0073\u002f\u0074\u0068\u0075\u006d\u0062\u006e\u0061\u0069\u006c\u002e\u006a\u0070\u0065\u0067")
		if _ga != nil {
			return _ga
		}
		if _ccff := _eg.Encode(_edg, _fbc.Thumbnail, nil); _ccff != nil {
			return _ccff
		}
	}
	if _fgeg := _bab.MarshalXMLByType(_fefc, _bbce, _f.SettingsType, _fbc.Settings.X()); _fgeg != nil {
		return _fgeg
	}
	_ddgd := _f.AbsoluteFilename(_bbce, _f.OfficeDocumentType, 0)
	if _gfe := _bab.MarshalXML(_fefc, _ddgd, _fbc._bbgg); _gfe != nil {
		return _gfe
	}
	if _ced := _bab.MarshalXML(_fefc, _bab.RelationsPathFor(_ddgd), _fbc._agcc.X()); _ced != nil {
		return _ced
	}
	if _fbc.Numbering.X() != nil {
		if _afb := _bab.MarshalXMLByType(_fefc, _bbce, _f.NumberingType, _fbc.Numbering.X()); _afb != nil {
			return _afb
		}
	}
	if _bbgga := _bab.MarshalXMLByType(_fefc, _bbce, _f.StylesType, _fbc.Styles.X()); _bbgga != nil {
		return _bbgga
	}
	if _fbc._cef != nil {
		if _ageg := _bab.MarshalXMLByType(_fefc, _bbce, _f.WebSettingsType, _fbc._cef); _ageg != nil {
			return _ageg
		}
	}
	if _fbc._gc != nil {
		if _aeg := _bab.MarshalXMLByType(_fefc, _bbce, _f.FontTableType, _fbc._gc); _aeg != nil {
			return _aeg
		}
	}
	if _fbc._fcb != nil {
		if _bcba := _bab.MarshalXMLByType(_fefc, _bbce, _f.EndNotesType, _fbc._fcb); _bcba != nil {
			return _bcba
		}
	}
	if _fbc._ab != nil {
		if _ccb := _bab.MarshalXMLByType(_fefc, _bbce, _f.FootNotesType, _fbc._ab); _ccb != nil {
			return _ccb
		}
	}
	for _bfe, _dfa := range _fbc._adae {
		if _fdf := _bab.MarshalXMLByTypeIndex(_fefc, _bbce, _f.ThemeType, _bfe+1, _dfa); _fdf != nil {
			return _fdf
		}
	}
	for _eggf, _fdg := range _fbc._cg {
		_gca := _f.AbsoluteFilename(_bbce, _f.HeaderType, _eggf+1)
		if _dgf := _bab.MarshalXML(_fefc, _gca, _fdg); _dgf != nil {
			return _dgf
		}
		if !_fbc._gba[_eggf].IsEmpty() {
			_bab.MarshalXML(_fefc, _bab.RelationsPathFor(_gca), _fbc._gba[_eggf].X())
		}
	}
	for _dbd, _fce := range _fbc._bbc {
		_fcba := _f.AbsoluteFilename(_bbce, _f.FooterType, _dbd+1)
		if _bffb := _bab.MarshalXMLByTypeIndex(_fefc, _bbce, _f.FooterType, _dbd+1, _fce); _bffb != nil {
			return _bffb
		}
		if !_fbc._dbc[_dbd].IsEmpty() {
			_bab.MarshalXML(_fefc, _bab.RelationsPathFor(_fcba), _fbc._dbc[_dbd].X())
		}
	}
	for _eae, _dcb := range _fbc.Images {
		if _cfb := _ec.AddImageToZip(_fefc, _dcb, _eae+1, _f.DocTypeDocument); _cfb != nil {
			return _cfb
		}
	}
	if _baa := _bab.MarshalXML(_fefc, _f.ContentTypesFilename, _fbc.ContentTypes.X()); _baa != nil {
		return _baa
	}
	if _bcc := _fbc.WriteExtraFiles(_fefc); _bcc != nil {
		return _bcc
	}
	return _fefc.Close()
}

// RemoveRun removes a child run from a paragraph.
func (_dabb Paragraph) RemoveRun(r Run) {
	for _, _dcbc := range _dabb._dbfe.EG_PContent {
		for _ggba, _aeeg := range _dcbc.EG_ContentRunContent {
			if _aeeg.R == r._fegg {
				copy(_dcbc.EG_ContentRunContent[_ggba:], _dcbc.EG_ContentRunContent[_ggba+1:])
				_dcbc.EG_ContentRunContent = _dcbc.EG_ContentRunContent[0 : len(_dcbc.EG_ContentRunContent)-1]
			}
			if _aeeg.Sdt != nil && _aeeg.Sdt.SdtContent != nil {
				for _fcede, _gegb := range _aeeg.Sdt.SdtContent.EG_ContentRunContent {
					if _gegb.R == r._fegg {
						copy(_aeeg.Sdt.SdtContent.EG_ContentRunContent[_fcede:], _aeeg.Sdt.SdtContent.EG_ContentRunContent[_fcede+1:])
						_aeeg.Sdt.SdtContent.EG_ContentRunContent = _aeeg.Sdt.SdtContent.EG_ContentRunContent[0 : len(_aeeg.Sdt.SdtContent.EG_ContentRunContent)-1]
					}
				}
			}
		}
	}
}

// Emboss returns true if paragraph emboss is on.
func (_fgbcb ParagraphProperties) Emboss() bool { return _decg(_fgbcb._cfdd.RPr.Emboss) }
func (_gdee Styles) initializeStyleDefaults() {
	_ecega := _gdee.AddStyle("\u004e\u006f\u0072\u006d\u0061\u006c", _ag.ST_StyleTypeParagraph, true)
	_ecega.SetName("\u004e\u006f\u0072\u006d\u0061\u006c")
	_ecega.SetPrimaryStyle(true)
	_ddeaa := _gdee.AddStyle("D\u0065f\u0061\u0075\u006c\u0074\u0050\u0061\u0072\u0061g\u0072\u0061\u0070\u0068Fo\u006e\u0074", _ag.ST_StyleTypeCharacter, true)
	_ddeaa.SetName("\u0044\u0065\u0066\u0061ul\u0074\u0020\u0050\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u0046\u006fn\u0074")
	_ddeaa.SetUISortOrder(1)
	_ddeaa.SetSemiHidden(true)
	_ddeaa.SetUnhideWhenUsed(true)
	_ggdb := _gdee.AddStyle("\u0054i\u0074\u006c\u0065\u0043\u0068\u0061r", _ag.ST_StyleTypeCharacter, false)
	_ggdb.SetName("\u0054\u0069\u0074\u006c\u0065\u0020\u0043\u0068\u0061\u0072")
	_ggdb.SetBasedOn(_ddeaa.StyleID())
	_ggdb.SetLinkedStyle("\u0054\u0069\u0074l\u0065")
	_ggdb.SetUISortOrder(10)
	_ggdb.RunProperties().Fonts().SetASCIITheme(_ag.ST_ThemeMajorAscii)
	_ggdb.RunProperties().Fonts().SetEastAsiaTheme(_ag.ST_ThemeMajorEastAsia)
	_ggdb.RunProperties().Fonts().SetHANSITheme(_ag.ST_ThemeMajorHAnsi)
	_ggdb.RunProperties().Fonts().SetCSTheme(_ag.ST_ThemeMajorBidi)
	_ggdb.RunProperties().SetSize(28 * _bf.Point)
	_ggdb.RunProperties().SetKerning(14 * _bf.Point)
	_ggdb.RunProperties().SetCharacterSpacing(-10 * _bf.Twips)
	_affb := _gdee.AddStyle("\u0054\u0069\u0074l\u0065", _ag.ST_StyleTypeParagraph, false)
	_affb.SetName("\u0054\u0069\u0074l\u0065")
	_affb.SetBasedOn(_ecega.StyleID())
	_affb.SetNextStyle(_ecega.StyleID())
	_affb.SetLinkedStyle(_ggdb.StyleID())
	_affb.SetUISortOrder(10)
	_affb.SetPrimaryStyle(true)
	_affb.ParagraphProperties().SetContextualSpacing(true)
	_affb.RunProperties().Fonts().SetASCIITheme(_ag.ST_ThemeMajorAscii)
	_affb.RunProperties().Fonts().SetEastAsiaTheme(_ag.ST_ThemeMajorEastAsia)
	_affb.RunProperties().Fonts().SetHANSITheme(_ag.ST_ThemeMajorHAnsi)
	_affb.RunProperties().Fonts().SetCSTheme(_ag.ST_ThemeMajorBidi)
	_affb.RunProperties().SetSize(28 * _bf.Point)
	_affb.RunProperties().SetKerning(14 * _bf.Point)
	_affb.RunProperties().SetCharacterSpacing(-10 * _bf.Twips)
	_edfeb := _gdee.AddStyle("T\u0061\u0062\u006c\u0065\u004e\u006f\u0072\u006d\u0061\u006c", _ag.ST_StyleTypeTable, false)
	_edfeb.SetName("\u004e\u006f\u0072m\u0061\u006c\u0020\u0054\u0061\u0062\u006c\u0065")
	_edfeb.SetUISortOrder(99)
	_edfeb.SetSemiHidden(true)
	_edfeb.SetUnhideWhenUsed(true)
	_edfeb.X().TblPr = _ag.NewCT_TblPrBase()
	_baff := NewTableWidth()
	_edfeb.X().TblPr.TblInd = _baff.X()
	_baff.SetValue(0 * _bf.Dxa)
	_edfeb.X().TblPr.TblCellMar = _ag.NewCT_TblCellMar()
	_baff = NewTableWidth()
	_edfeb.X().TblPr.TblCellMar.Top = _baff.X()
	_baff.SetValue(0 * _bf.Dxa)
	_baff = NewTableWidth()
	_edfeb.X().TblPr.TblCellMar.Bottom = _baff.X()
	_baff.SetValue(0 * _bf.Dxa)
	_baff = NewTableWidth()
	_edfeb.X().TblPr.TblCellMar.Left = _baff.X()
	_baff.SetValue(108 * _bf.Dxa)
	_baff = NewTableWidth()
	_edfeb.X().TblPr.TblCellMar.Right = _baff.X()
	_baff.SetValue(108 * _bf.Dxa)
	_caba := _gdee.AddStyle("\u004e\u006f\u004c\u0069\u0073\u0074", _ag.ST_StyleTypeNumbering, false)
	_caba.SetName("\u004eo\u0020\u004c\u0069\u0073\u0074")
	_caba.SetUISortOrder(1)
	_caba.SetSemiHidden(true)
	_caba.SetUnhideWhenUsed(true)
	_cfdfb := []_bf.Distance{16, 13, 12, 11, 11, 11, 11, 11, 11}
	_aebd := []_bf.Distance{240, 40, 40, 40, 40, 40, 40, 40, 40}
	for _defad := 0; _defad < 9; _defad++ {
		_dfbcf := _af.Sprintf("\u0048e\u0061\u0064\u0069\u006e\u0067\u0025d", _defad+1)
		_fdfc := _gdee.AddStyle(_dfbcf+"\u0043\u0068\u0061\u0072", _ag.ST_StyleTypeCharacter, false)
		_fdfc.SetName(_af.Sprintf("\u0048e\u0061d\u0069\u006e\u0067\u0020\u0025\u0064\u0020\u0043\u0068\u0061\u0072", _defad+1))
		_fdfc.SetBasedOn(_ddeaa.StyleID())
		_fdfc.SetLinkedStyle(_dfbcf)
		_fdfc.SetUISortOrder(9 + _defad)
		_fdfc.RunProperties().SetSize(_cfdfb[_defad] * _bf.Point)
		_afdb := _gdee.AddStyle(_dfbcf, _ag.ST_StyleTypeParagraph, false)
		_afdb.SetName(_af.Sprintf("\u0068\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0064", _defad+1))
		_afdb.SetNextStyle(_ecega.StyleID())
		_afdb.SetLinkedStyle(_afdb.StyleID())
		_afdb.SetUISortOrder(9 + _defad)
		_afdb.SetPrimaryStyle(true)
		_afdb.ParagraphProperties().SetKeepNext(true)
		_afdb.ParagraphProperties().SetSpacing(_aebd[_defad]*_bf.Twips, 0)
		_afdb.ParagraphProperties().SetOutlineLevel(_defad)
		_afdb.RunProperties().SetSize(_cfdfb[_defad] * _bf.Point)
	}
}
func _faea(_cfgb *_ag.CT_P, _aaaca map[string]string) {
	for _, _eefg := range _cfgb.EG_PContent {
		if _eefg.Hyperlink != nil && _eefg.Hyperlink.IdAttr != nil {
			if _aegf, _bbb := _aaaca[*_eefg.Hyperlink.IdAttr]; _bbb {
				*_eefg.Hyperlink.IdAttr = _aegf
			}
		}
	}
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (_bddg ParagraphProperties) SetKeepOnOnePage(b bool) {
	if !b {
		_bddg._cfdd.KeepLines = nil
	} else {
		_bddg._cfdd.KeepLines = _ag.NewCT_OnOff()
	}
}

// TableConditionalFormatting controls the conditional formatting within a table
// style.
type TableConditionalFormatting struct{ _bdceb *_ag.CT_TblStylePr }

// SetTableIndent sets the Table Indent from the Leading Margin
func (_ddada TableStyleProperties) SetTableIndent(ind _bf.Distance) {
	_ddada._cgcba.TblInd = _ag.NewCT_TblWidth()
	_ddada._cgcba.TblInd.TypeAttr = _ag.ST_TblWidthDxa
	_ddada._cgcba.TblInd.WAttr = &_ag.ST_MeasurementOrPercent{}
	_ddada._cgcba.TblInd.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_ddada._cgcba.TblInd.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(ind / _bf.Dxa))
}

// SetOutlineLevel sets the outline level of this style.
func (_fadcf ParagraphStyleProperties) SetOutlineLevel(lvl int) {
	_fadcf._ddde.OutlineLvl = _ag.NewCT_DecimalNumber()
	_fadcf._ddde.OutlineLvl.ValAttr = int64(lvl)
}

// GetHeader gets a section Header for given type t [ST_HdrFtrDefault, ST_HdrFtrEven, ST_HdrFtrFirst]
func (_eeeeg Section) GetHeader(t _ag.ST_HdrFtr) (Header, bool) {
	for _, _cgce := range _eeeeg._cdec.EG_HdrFtrReferences {
		if _cgce.HeaderReference.TypeAttr == t {
			for _, _eebd := range _eeeeg._cce.Headers() {
				_bgcc := _eeeeg._cce._agcc.FindRIDForN(_eebd.Index(), _f.HeaderType)
				if _bgcc == _cgce.HeaderReference.IdAttr {
					return _eebd, true
				}
			}
		}
	}
	return Header{}, false
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (_cddc TableProperties) SetCellSpacingPercent(pct float64) {
	_cddc._gfaa.TblCellSpacing = _ag.NewCT_TblWidth()
	_cddc._gfaa.TblCellSpacing.TypeAttr = _ag.ST_TblWidthPct
	_cddc._gfaa.TblCellSpacing.WAttr = &_ag.ST_MeasurementOrPercent{}
	_cddc._gfaa.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_cddc._gfaa.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(pct * 50))
}

// SetWindowControl controls if the first or last line of the paragraph is
// allowed to dispay on a separate page.
func (_bgfa ParagraphProperties) SetWindowControl(b bool) {
	if !b {
		_bgfa._cfdd.WidowControl = nil
	} else {
		_bgfa._cfdd.WidowControl = _ag.NewCT_OnOff()
	}
}

// SetRight sets the right border to a specified type, color and thickness.
func (_gdcc TableBorders) SetRight(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_gdcc._egbeb.Right = _ag.NewCT_Border()
	_ddecc(_gdcc._egbeb.Right, t, c, thickness)
}

// SetNumberingDefinitionByID sets the numbering definition ID directly, which must
// match an ID defined in numbering.xml
func (_fbga Paragraph) SetNumberingDefinitionByID(abstractNumberID int64) {
	_fbga.ensurePPr()
	if _fbga._dbfe.PPr.NumPr == nil {
		_fbga._dbfe.PPr.NumPr = _ag.NewCT_NumPr()
	}
	_fceg := _ag.NewCT_DecimalNumber()
	_fceg.ValAttr = int64(abstractNumberID)
	_fbga._dbfe.PPr.NumPr.NumId = _fceg
}

// GetImage returns the ImageRef associated with an AnchoredDrawing.
func (_ge AnchoredDrawing) GetImage() (_ec.ImageRef, bool) {
	_be := _ge._de.Graphic.GraphicData.Any
	if len(_be) > 0 {
		_gg, _bg := _be[0].(*_afd.Pic)
		if _bg {
			if _gg.BlipFill != nil && _gg.BlipFill.Blip != nil && _gg.BlipFill.Blip.EmbedAttr != nil {
				return _ge._aa.GetImageByRelID(*_gg.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return _ec.ImageRef{}, false
}

// RemoveMailMerge removes any mail merge settings
func (_cccf Settings) RemoveMailMerge() { _cccf._dcaa.MailMerge = nil }

// SetAfterAuto controls if spacing after a paragraph is automatically determined.
func (_babf ParagraphSpacing) SetAfterAuto(b bool) {
	if b {
		_babf._gggda.AfterAutospacingAttr = &_dc.ST_OnOff{}
		_babf._gggda.AfterAutospacingAttr.Bool = _f.Bool(true)
	} else {
		_babf._gggda.AfterAutospacingAttr = nil
	}
}

// RightToLeft returns true if run text goes from right to left.
func (_cbdd RunProperties) RightToLeft() bool { return _decg(_cbdd._ddda.Rtl) }

// ClearContent clears any content in the run (text, tabs, breaks, etc.)
func (_gffa Run) ClearContent() { _gffa._fegg.EG_RunInnerContent = nil }

// Outline returns true if run outline is on.
func (_ebgdb RunProperties) Outline() bool { return _decg(_ebgdb._ddda.Outline) }

var _bfb = false

func (_afa *Document) validateBookmarks() error {
	_gdd := make(map[string]struct{})
	for _, _dggf := range _afa.Bookmarks() {
		if _, _egd := _gdd[_dggf.Name()]; _egd {
			return _af.Errorf("d\u0075\u0070\u006c\u0069\u0063\u0061t\u0065\u0020\u0062\u006f\u006f\u006b\u006d\u0061\u0072k\u0020\u0025\u0073 \u0066o\u0075\u006e\u0064", _dggf.Name())
		}
		_gdd[_dggf.Name()] = struct{}{}
	}
	return nil
}

// SetValue sets the width value.
func (_fgac TableWidth) SetValue(m _bf.Distance) {
	_fgac._eccef.WAttr = &_ag.ST_MeasurementOrPercent{}
	_fgac._eccef.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_fgac._eccef.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(m / _bf.Twips))
	_fgac._eccef.TypeAttr = _ag.ST_TblWidthDxa
}

// Footnote is an individual footnote reference within the document.
type Footnote struct {
	_gagb  *Document
	_faaeg *_ag.CT_FtnEdn
}

// Bookmarks returns all of the bookmarks defined in the document.
func (_cba Document) Bookmarks() []Bookmark {
	if _cba._bbgg.Body == nil {
		return nil
	}
	_dbgf := []Bookmark{}
	for _, _edff := range _cba._bbgg.Body.EG_BlockLevelElts {
		for _, _cbab := range _edff.EG_ContentBlockContent {
			for _, _dfgf := range _bagf(_cbab) {
				_dbgf = append(_dbgf, _dfgf)
			}
		}
	}
	return _dbgf
}

// MultiLevelType returns the multilevel type, or ST_MultiLevelTypeUnset if not set.
func (_ffbea NumberingDefinition) MultiLevelType() _ag.ST_MultiLevelType {
	if _ffbea._cgfd.MultiLevelType != nil {
		return _ffbea._cgfd.MultiLevelType.ValAttr
	} else {
		return _ag.ST_MultiLevelTypeUnset
	}
}

// AddBookmark adds a bookmark to a document that can then be used from a hyperlink. Name is a document
// unique name that identifies the bookmark so it can be referenced from hyperlinks.
func (_eeaf Paragraph) AddBookmark(name string) Bookmark {
	_baad := _ag.NewEG_PContent()
	_egeb := _ag.NewEG_ContentRunContent()
	_baad.EG_ContentRunContent = append(_baad.EG_ContentRunContent, _egeb)
	_fggg := _ag.NewEG_RunLevelElts()
	_egeb.EG_RunLevelElts = append(_egeb.EG_RunLevelElts, _fggg)
	_baae := _ag.NewEG_RangeMarkupElements()
	_beecc := _ag.NewCT_Bookmark()
	_baae.BookmarkStart = _beecc
	_fggg.EG_RangeMarkupElements = append(_fggg.EG_RangeMarkupElements, _baae)
	_baae = _ag.NewEG_RangeMarkupElements()
	_baae.BookmarkEnd = _ag.NewCT_MarkupRange()
	_fggg.EG_RangeMarkupElements = append(_fggg.EG_RangeMarkupElements, _baae)
	_eeaf._dbfe.EG_PContent = append(_eeaf._dbfe.EG_PContent, _baad)
	_fbgfg := Bookmark{_beecc}
	_fbgfg.SetName(name)
	return _fbgfg
}

// AddFieldWithFormatting adds a field (automatically computed text) to the
// document with field specifc formatting.
func (_gbec Run) AddFieldWithFormatting(code string, fmt string, isDirty bool) {
	_eefc := _gbec.newIC()
	_eefc.FldChar = _ag.NewCT_FldChar()
	_eefc.FldChar.FldCharTypeAttr = _ag.ST_FldCharTypeBegin
	if isDirty {
		_eefc.FldChar.DirtyAttr = &_dc.ST_OnOff{}
		_eefc.FldChar.DirtyAttr.Bool = _f.Bool(true)
	}
	_eefc = _gbec.newIC()
	_eefc.InstrText = _ag.NewCT_Text()
	if fmt != "" {
		_eefc.InstrText.Content = code + "\u0020" + fmt
	} else {
		_eefc.InstrText.Content = code
	}
	_eefc = _gbec.newIC()
	_eefc.FldChar = _ag.NewCT_FldChar()
	_eefc.FldChar.FldCharTypeAttr = _ag.ST_FldCharTypeEnd
}

// RunProperties returns the run style properties.
func (_dgfgca Style) RunProperties() RunProperties {
	if _dgfgca._gade.RPr == nil {
		_dgfgca._gade.RPr = _ag.NewCT_RPr()
	}
	return RunProperties{_dgfgca._gade.RPr}
}

// SetWidthAuto sets the the table width to automatic.
func (_dedc TableProperties) SetWidthAuto() {
	_dedc._gfaa.TblW = _ag.NewCT_TblWidth()
	_dedc._gfaa.TblW.TypeAttr = _ag.ST_TblWidthAuto
}

// Font returns the name of run font family.
func (_fffce RunProperties) Font() string {
	if _abef := _fffce._ddda.RFonts; _abef != nil {
		if _abef.AsciiAttr != nil {
			return *_abef.AsciiAttr
		} else if _abef.HAnsiAttr != nil {
			return *_abef.HAnsiAttr
		} else if _abef.CsAttr != nil {
			return *_abef.CsAttr
		}
	}
	return ""
}
func (_fcgdc Paragraph) addSeparateFldChar() *_ag.CT_FldChar {
	_ffae := _fcgdc.addFldChar()
	_ffae.FldCharTypeAttr = _ag.ST_FldCharTypeSeparate
	return _ffae
}
func (_aaeg Footnote) content() []*_ag.EG_ContentBlockContent {
	var _bgeb []*_ag.EG_ContentBlockContent
	for _, _efge := range _aaeg._faaeg.EG_BlockLevelElts {
		_bgeb = append(_bgeb, _efge.EG_ContentBlockContent...)
	}
	return _bgeb
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (_acfc Paragraph) SetFirstLineIndent(m _bf.Distance) {
	_acfc.ensurePPr()
	_edeg := _acfc._dbfe.PPr
	if _edeg.Ind == nil {
		_edeg.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_edeg.Ind.FirstLineAttr = nil
	} else {
		_edeg.Ind.FirstLineAttr = &_dc.ST_TwipsMeasure{}
		_edeg.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(m / _bf.Twips))
	}
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_gge CellBorders) SetLeft(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_gge._eaa.Left = _ag.NewCT_Border()
	_ddecc(_gge._eaa.Left, t, c, thickness)
}

// DoubleStrike returns true if paragraph is double striked.
func (_acdf ParagraphProperties) DoubleStrike() bool { return _decg(_acdf._cfdd.RPr.Dstrike) }

// SetColor sets a specific color or auto.
func (_aaf Color) SetColor(v _ed.Color) {
	if v.IsAuto() {
		_aaf._ccc.ValAttr.ST_HexColorAuto = _ag.ST_HexColorAutoAuto
		_aaf._ccc.ValAttr.ST_HexColorRGB = nil
	} else {
		_aaf._ccc.ValAttr.ST_HexColorAuto = _ag.ST_HexColorAutoUnset
		_aaf._ccc.ValAttr.ST_HexColorRGB = v.AsRGBString()
	}
}

// GetColor returns the color.Color object representing the run color.
func (_fgcg RunProperties) GetColor() _ed.Color {
	if _cgdc := _fgcg._ddda.Color; _cgdc != nil {
		_bebaa := _cgdc.ValAttr
		if _bebaa.ST_HexColorRGB != nil {
			return _ed.FromHex(*_bebaa.ST_HexColorRGB)
		}
	}
	return _ed.Color{}
}

// SetLastColumn controls the conditional formatting for the last column in a table.
func (_debf TableLook) SetLastColumn(on bool) {
	if !on {
		_debf._cfbff.LastColumnAttr = &_dc.ST_OnOff{}
		_debf._cfbff.LastColumnAttr.ST_OnOff1 = _dc.ST_OnOff1Off
	} else {
		_debf._cfbff.LastColumnAttr = &_dc.ST_OnOff{}
		_debf._cfbff.LastColumnAttr.ST_OnOff1 = _dc.ST_OnOff1On
	}
}
func (_gcac *Document) validateTableCells() error {
	for _, _ead := range _gcac._bbgg.Body.EG_BlockLevelElts {
		for _, _dafg := range _ead.EG_ContentBlockContent {
			for _, _gaaa := range _dafg.Tbl {
				for _, _gcbf := range _gaaa.EG_ContentRowContent {
					for _, _dad := range _gcbf.Tr {
						_fffc := false
						for _, _faef := range _dad.EG_ContentCellContent {
							_bcde := false
							for _, _gfg := range _faef.Tc {
								_fffc = true
								for _, _cdbc := range _gfg.EG_BlockLevelElts {
									for _, _dabe := range _cdbc.EG_ContentBlockContent {
										if len(_dabe.P) > 0 {
											_bcde = true
											break
										}
									}
								}
							}
							if !_bcde {
								return _g.New("t\u0061\u0062\u006c\u0065\u0020\u0063e\u006c\u006c\u0020\u006d\u0075\u0073t\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061p\u0068")
							}
						}
						if !_fffc {
							return _g.New("\u0074\u0061b\u006c\u0065\u0020\u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020ce\u006c\u006c")
						}
					}
				}
			}
		}
	}
	return nil
}

// SetSmallCaps sets the run to small caps.
func (_cafc RunProperties) SetSmallCaps(b bool) {
	if !b {
		_cafc._ddda.SmallCaps = nil
	} else {
		_cafc._ddda.SmallCaps = _ag.NewCT_OnOff()
	}
}

// X returns the inner wrapped XML type.
func (_cbcd Footer) X() *_ag.Ftr { return _cbcd._edec }

// SetEmboss sets the run to embossed text.
func (_fafb RunProperties) SetEmboss(b bool) {
	if !b {
		_fafb._ddda.Emboss = nil
	} else {
		_fafb._ddda.Emboss = _ag.NewCT_OnOff()
	}
}

// SetAlignment controls the paragraph alignment
func (_efddd ParagraphStyleProperties) SetAlignment(align _ag.ST_Jc) {
	if align == _ag.ST_JcUnset {
		_efddd._ddde.Jc = nil
	} else {
		_efddd._ddde.Jc = _ag.NewCT_Jc()
		_efddd._ddde.Jc.ValAttr = align
	}
}

// InlineDrawing is an inlined image within a run.
type InlineDrawing struct {
	_agaeg *Document
	_egff  *_ag.WdInline
}

// FormFieldType is the type of the form field.
//go:generate stringer -type=FormFieldType
type FormFieldType byte

// SetRightPct sets the cell right margin
func (_bbg CellMargins) SetRightPct(pct float64) {
	_bbg._cacb.Right = _ag.NewCT_TblWidth()
	_adg(_bbg._cacb.Right, pct)
}

// NewSettings constructs a new empty Settings
func NewSettings() Settings {
	_bbgd := _ag.NewSettings()
	_bbgd.Compat = _ag.NewCT_Compat()
	_eeedb := _ag.NewCT_CompatSetting()
	_eeedb.NameAttr = _f.String("\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074y\u004d\u006f\u0064\u0065")
	_eeedb.UriAttr = _f.String("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006fff\u0069\u0063\u0065/\u0077o\u0072\u0064")
	_eeedb.ValAttr = _f.String("\u0031\u0035")
	_bbgd.Compat.CompatSetting = append(_bbgd.Compat.CompatSetting, _eeedb)
	return Settings{_bbgd}
}

// ParagraphProperties returns the paragraph properties controlling text formatting within the table.
func (_dfad TableConditionalFormatting) ParagraphProperties() ParagraphStyleProperties {
	if _dfad._bdceb.PPr == nil {
		_dfad._bdceb.PPr = _ag.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_dfad._bdceb.PPr}
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (_ebca ParagraphProperties) SetFirstLineIndent(m _bf.Distance) {
	if _ebca._cfdd.Ind == nil {
		_ebca._cfdd.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_ebca._cfdd.Ind.FirstLineAttr = nil
	} else {
		_ebca._cfdd.Ind.FirstLineAttr = &_dc.ST_TwipsMeasure{}
		_ebca._cfdd.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(m / _bf.Twips))
	}
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (_afdg CellBorders) SetBottom(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_afdg._eaa.Bottom = _ag.NewCT_Border()
	_ddecc(_afdg._eaa.Bottom, t, c, thickness)
}
func _bagf(_dcd *_ag.EG_ContentBlockContent) []Bookmark {
	_fgf := []Bookmark{}
	for _, _acfb := range _dcd.P {
		for _, _baec := range _acfb.EG_PContent {
			for _, _ebbe := range _baec.EG_ContentRunContent {
				for _, _ccag := range _ebbe.EG_RunLevelElts {
					for _, _dfg := range _ccag.EG_RangeMarkupElements {
						if _dfg.BookmarkStart != nil {
							_fgf = append(_fgf, Bookmark{_dfg.BookmarkStart})
						}
					}
				}
			}
		}
	}
	for _, _fdgb := range _dcd.EG_RunLevelElts {
		for _, _cbb := range _fdgb.EG_RangeMarkupElements {
			if _cbb.BookmarkStart != nil {
				_fgf = append(_fgf, Bookmark{_cbb.BookmarkStart})
			}
		}
	}
	for _, _gbg := range _dcd.Tbl {
		for _, _cgeg := range _gbg.EG_ContentRowContent {
			for _, _bfdf := range _cgeg.Tr {
				for _, _egbb := range _bfdf.EG_ContentCellContent {
					for _, _dbdb := range _egbb.Tc {
						for _, _gadf := range _dbdb.EG_BlockLevelElts {
							for _, _egga := range _gadf.EG_ContentBlockContent {
								for _, _fbce := range _bagf(_egga) {
									_fgf = append(_fgf, _fbce)
								}
							}
						}
					}
				}
			}
		}
	}
	return _fgf
}

// AddRow adds a row to a table.
func (_feda Table) AddRow() Row {
	_agcad := _ag.NewEG_ContentRowContent()
	_feda._fdeda.EG_ContentRowContent = append(_feda._fdeda.EG_ContentRowContent, _agcad)
	_bgcf := _ag.NewCT_Row()
	_agcad.Tr = append(_agcad.Tr, _bgcf)
	return Row{_feda._aebg, _bgcf}
}

// SetName sets the name of the image, visible in the properties of the image
// within Word.
func (_ad AnchoredDrawing) SetName(name string) {
	_ad._de.DocPr.NameAttr = name
	for _, _bc := range _ad._de.Graphic.GraphicData.Any {
		if _cf, _fa := _bc.(*_afd.Pic); _fa {
			_cf.NvPicPr.CNvPr.DescrAttr = _f.String(name)
		}
	}
}

// SetLastRow controls the conditional formatting for the last row in a table.
// This is called the 'Total' row within Word.
func (_facab TableLook) SetLastRow(on bool) {
	if !on {
		_facab._cfbff.LastRowAttr = &_dc.ST_OnOff{}
		_facab._cfbff.LastRowAttr.ST_OnOff1 = _dc.ST_OnOff1Off
	} else {
		_facab._cfbff.LastRowAttr = &_dc.ST_OnOff{}
		_facab._cfbff.LastRowAttr.ST_OnOff1 = _dc.ST_OnOff1On
	}
}

// X returns the inner wrapped XML type.
func (_bgbgc TableWidth) X() *_ag.CT_TblWidth { return _bgbgc._eccef }

// SetSize sets the font size for a run.
func (_gbfd RunProperties) SetSize(size _bf.Distance) {
	_gbfd._ddda.Sz = _ag.NewCT_HpsMeasure()
	_gbfd._ddda.Sz.ValAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(size / _bf.HalfPoint))
	_gbfd._ddda.SzCs = _ag.NewCT_HpsMeasure()
	_gbfd._ddda.SzCs.ValAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(size / _bf.HalfPoint))
}

// SizeValue returns the value of paragraph font size in points.
func (_fbgff ParagraphProperties) SizeValue() float64 {
	if _defc := _fbgff._cfdd.RPr.Sz; _defc != nil {
		_aaag := _defc.ValAttr
		if _aaag.ST_UnsignedDecimalNumber != nil {
			return float64(*_aaag.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SetNumberingLevel sets the numbering level of a paragraph.  If used, then the
// NumberingDefinition must also be set via SetNumberingDefinition or
// SetNumberingDefinitionByID.
func (_cggfg Paragraph) SetNumberingLevel(listLevel int) {
	_cggfg.ensurePPr()
	if _cggfg._dbfe.PPr.NumPr == nil {
		_cggfg._dbfe.PPr.NumPr = _ag.NewCT_NumPr()
	}
	_ddbde := _ag.NewCT_DecimalNumber()
	_ddbde.ValAttr = int64(listLevel)
	_cggfg._dbfe.PPr.NumPr.Ilvl = _ddbde
}

// Caps returns true if paragraph font is capitalized.
func (_cgcg ParagraphProperties) Caps() bool { return _decg(_cgcg._cfdd.RPr.Caps) }
func _adg(_adfc *_ag.CT_TblWidth, _dab float64) {
	_adfc.TypeAttr = _ag.ST_TblWidthPct
	_adfc.WAttr = &_ag.ST_MeasurementOrPercent{}
	_adfc.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_adfc.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(_dab * 50))
}

// MailMerge finds mail merge fields and replaces them with the text provided.  It also removes
// the mail merge source info from the document settings.
func (_eceg *Document) MailMerge(mergeContent map[string]string) {
	_gceg := _eceg.mergeFields()
	_daebg := map[Paragraph][]Run{}
	for _, _ccda := range _gceg {
		_baeee, _acc := mergeContent[_ccda._cabg]
		if _acc {
			if _ccda._ebe {
				_baeee = _c.ToUpper(_baeee)
			} else if _ccda._gdad {
				_baeee = _c.ToLower(_baeee)
			} else if _ccda._fcd {
				_baeee = _c.Title(_baeee)
			} else if _ccda._efad {
				_eece := _ca.Buffer{}
				for _fgbb, _fffa := range _baeee {
					if _fgbb == 0 {
						_eece.WriteRune(_e.ToUpper(_fffa))
					} else {
						_eece.WriteRune(_fffa)
					}
				}
				_baeee = _eece.String()
			}
			if _baeee != "" && _ccda._gbeg != "" {
				_baeee = _ccda._gbeg + _baeee
			}
			if _baeee != "" && _ccda._ebce != "" {
				_baeee = _baeee + _ccda._ebce
			}
		}
		if _ccda._baba {
			if len(_ccda._fcdf.FldSimple) == 1 && len(_ccda._fcdf.FldSimple[0].EG_PContent) == 1 && len(_ccda._fcdf.FldSimple[0].EG_PContent[0].EG_ContentRunContent) == 1 {
				_caeb := &_ag.EG_ContentRunContent{}
				_caeb.R = _ccda._fcdf.FldSimple[0].EG_PContent[0].EG_ContentRunContent[0].R
				_ccda._fcdf.FldSimple = nil
				_cgcb := Run{_eceg, _caeb.R}
				_cgcb.ClearContent()
				_cgcb.AddText(_baeee)
				_ccda._fcdf.EG_ContentRunContent = append(_ccda._fcdf.EG_ContentRunContent, _caeb)
			}
		} else {
			_ggaf := _ccda._cgdg.Runs()
			for _bccd := _ccda._gaab; _bccd <= _ccda._efged; _bccd++ {
				if _bccd == _ccda._cgfaf+1 {
					_ggaf[_bccd].ClearContent()
					_ggaf[_bccd].AddText(_baeee)
				} else {
					_daebg[_ccda._cgdg] = append(_daebg[_ccda._cgdg], _ggaf[_bccd])
				}
			}
		}
	}
	for _ggfa, _egbg := range _daebg {
		for _, _bcbb := range _egbg {
			_ggfa.RemoveRun(_bcbb)
		}
	}
	_eceg.Settings.RemoveMailMerge()
}

// NewTableWidth returns a newly intialized TableWidth
func NewTableWidth() TableWidth { return TableWidth{_ag.NewCT_TblWidth()} }

// SetStyle sets the style of a paragraph and is identical to setting it on the
// paragraph's Properties()
func (_fcgcd Paragraph) SetStyle(s string) {
	_fcgcd.ensurePPr()
	if s == "" {
		_fcgcd._dbfe.PPr.PStyle = nil
	} else {
		_fcgcd._dbfe.PPr.PStyle = _ag.NewCT_String()
		_fcgcd._dbfe.PPr.PStyle.ValAttr = s
	}
}

// SetTargetBookmark sets the bookmark target of the hyperlink.
func (_aeafd HyperLink) SetTargetBookmark(bm Bookmark) {
	_aeafd._dceb.AnchorAttr = _f.String(bm.Name())
	_aeafd._dceb.IdAttr = nil
}

// SetCellSpacing sets the cell spacing within a table.
func (_bbcda TableProperties) SetCellSpacing(m _bf.Distance) {
	_bbcda._gfaa.TblCellSpacing = _ag.NewCT_TblWidth()
	_bbcda._gfaa.TblCellSpacing.TypeAttr = _ag.ST_TblWidthDxa
	_bbcda._gfaa.TblCellSpacing.WAttr = &_ag.ST_MeasurementOrPercent{}
	_bbcda._gfaa.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_bbcda._gfaa.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(m / _bf.Dxa))
}

// X returns the inner wrapped XML type.
func (_fdff ParagraphStyleProperties) X() *_ag.CT_PPrGeneral { return _fdff._ddde }

const (
	OnOffValueUnset OnOffValue = iota
	OnOffValueOff
	OnOffValueOn
)

// SetWidthPercent sets the cell to a width percentage.
func (_agc CellProperties) SetWidthPercent(pct float64) {
	_agc._dac.TcW = _ag.NewCT_TblWidth()
	_agc._dac.TcW.TypeAttr = _ag.ST_TblWidthPct
	_agc._dac.TcW.WAttr = &_ag.ST_MeasurementOrPercent{}
	_agc._dac.TcW.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_agc._dac.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(pct * 50))
}
func (_deeg Paragraph) ensurePPr() {
	if _deeg._dbfe.PPr == nil {
		_deeg._dbfe.PPr = _ag.NewCT_PPr()
	}
}

// Shadow returns true if run shadow is on.
func (_dggc RunProperties) Shadow() bool { return _decg(_dggc._ddda.Shadow) }

// TableConditionalFormatting returns a conditional formatting object of a given
// type.  Calling this method repeatedly will return the same object.
func (_dbfc Style) TableConditionalFormatting(typ _ag.ST_TblStyleOverrideType) TableConditionalFormatting {
	for _, _gbddd := range _dbfc._gade.TblStylePr {
		if _gbddd.TypeAttr == typ {
			return TableConditionalFormatting{_gbddd}
		}
	}
	_cggcg := _ag.NewCT_TblStylePr()
	_cggcg.TypeAttr = typ
	_dbfc._gade.TblStylePr = append(_dbfc._gade.TblStylePr, _cggcg)
	return TableConditionalFormatting{_cggcg}
}
func (_eee *Document) insertTable(_fae Paragraph, _egb bool) Table {
	_fec := _eee._bbgg.Body
	if _fec == nil {
		return _eee.AddTable()
	}
	_gcb := _fae.X()
	for _efd, _fba := range _fec.EG_BlockLevelElts {
		for _, _gfa := range _fba.EG_ContentBlockContent {
			for _dcf, _egbe := range _gfa.P {
				if _egbe == _gcb {
					_fcbc := _ag.NewCT_Tbl()
					_eaga := _ag.NewEG_BlockLevelElts()
					_ggda := _ag.NewEG_ContentBlockContent()
					_eaga.EG_ContentBlockContent = append(_eaga.EG_ContentBlockContent, _ggda)
					_ggda.Tbl = append(_ggda.Tbl, _fcbc)
					_fec.EG_BlockLevelElts = append(_fec.EG_BlockLevelElts, nil)
					if _egb {
						copy(_fec.EG_BlockLevelElts[_efd+1:], _fec.EG_BlockLevelElts[_efd:])
						_fec.EG_BlockLevelElts[_efd] = _eaga
						if _dcf != 0 {
							_ffa := _ag.NewEG_BlockLevelElts()
							_baf := _ag.NewEG_ContentBlockContent()
							_ffa.EG_ContentBlockContent = append(_ffa.EG_ContentBlockContent, _baf)
							_baf.P = _gfa.P[:_dcf]
							_fec.EG_BlockLevelElts = append(_fec.EG_BlockLevelElts, nil)
							copy(_fec.EG_BlockLevelElts[_efd+1:], _fec.EG_BlockLevelElts[_efd:])
							_fec.EG_BlockLevelElts[_efd] = _ffa
						}
						_gfa.P = _gfa.P[_dcf:]
					} else {
						copy(_fec.EG_BlockLevelElts[_efd+2:], _fec.EG_BlockLevelElts[_efd+1:])
						_fec.EG_BlockLevelElts[_efd+1] = _eaga
						if _dcf != len(_gfa.P)-1 {
							_eeb := _ag.NewEG_BlockLevelElts()
							_cae := _ag.NewEG_ContentBlockContent()
							_eeb.EG_ContentBlockContent = append(_eeb.EG_ContentBlockContent, _cae)
							_cae.P = _gfa.P[_dcf+1:]
							_fec.EG_BlockLevelElts = append(_fec.EG_BlockLevelElts, nil)
							copy(_fec.EG_BlockLevelElts[_efd+3:], _fec.EG_BlockLevelElts[_efd+2:])
							_fec.EG_BlockLevelElts[_efd+2] = _eeb
						}
						_gfa.P = _gfa.P[:_dcf+1]
					}
					return Table{_eee, _fcbc}
				}
			}
			for _, _fea := range _gfa.Tbl {
				_fdc := _babc(_fea, _gcb, _egb)
				if _fdc != nil {
					return Table{_eee, _fdc}
				}
			}
		}
	}
	return _eee.AddTable()
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_baga Header) AddImage(i _ec.Image) (_ec.ImageRef, error) {
	var _fegf _ec.Relationships
	for _fgegb, _afbb := range _baga._cedg._cg {
		if _afbb == _baga._aeaf {
			_fegf = _baga._cedg._gba[_fgegb]
		}
	}
	_febg := _ec.MakeImageRef(i, &_baga._cedg.DocBase, _fegf)
	if i.Data == nil && i.Path == "" {
		return _febg, _g.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _febg, _g.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _febg, _g.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	_baga._cedg.Images = append(_baga._cedg.Images, _febg)
	_fdcb := _af.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_baga._cedg.Images), i.Format)
	_aafa := _fegf.AddRelationship(_fdcb, _f.ImageType)
	_febg.SetRelID(_aafa.X().IdAttr)
	return _febg, nil
}
func (_cfee Endnote) content() []*_ag.EG_ContentBlockContent {
	var _gfed []*_ag.EG_ContentBlockContent
	for _, _cdde := range _cfee._gae.EG_BlockLevelElts {
		_gfed = append(_gfed, _cdde.EG_ContentBlockContent...)
	}
	return _gfed
}

// SetHeight allows controlling the height of a row within a table.
func (_dgbb RowProperties) SetHeight(ht _bf.Distance, rule _ag.ST_HeightRule) {
	if rule == _ag.ST_HeightRuleUnset {
		_dgbb._ebad.TrHeight = nil
	} else {
		_egfa := _ag.NewCT_Height()
		_egfa.HRuleAttr = rule
		_egfa.ValAttr = &_dc.ST_TwipsMeasure{}
		_egfa.ValAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(ht / _bf.Twips))
		_dgbb._ebad.TrHeight = []*_ag.CT_Height{_egfa}
	}
}

// Clear removes all of the content from within a run.
func (_ggdf Run) Clear() { _ggdf._fegg.EG_RunInnerContent = nil }

// SetText sets the text to be used in bullet mode.
func (_gdfa NumberingLevel) SetText(t string) {
	if t == "" {
		_gdfa._bdce.LvlText = nil
	} else {
		_gdfa._bdce.LvlText = _ag.NewCT_LevelText()
		_gdfa._bdce.LvlText.ValAttr = _f.String(t)
	}
}

// MergeFields returns the list of all mail merge fields found in the document.
func (_ebf Document) MergeFields() []string {
	_dfcg := map[string]struct{}{}
	for _, _ffge := range _ebf.mergeFields() {
		_dfcg[_ffge._cabg] = struct{}{}
	}
	_edfe := []string{}
	for _gbedb := range _dfcg {
		_edfe = append(_edfe, _gbedb)
	}
	return _edfe
}

// SetPageBreakBefore controls if there is a page break before this paragraph.
func (_edbd ParagraphProperties) SetPageBreakBefore(b bool) {
	if !b {
		_edbd._cfdd.PageBreakBefore = nil
	} else {
		_edbd._cfdd.PageBreakBefore = _ag.NewCT_OnOff()
	}
}

// AddHyperLink adds a new hyperlink to a parapgraph.
func (_dbggfe Paragraph) AddHyperLink() HyperLink {
	_deef := _ag.NewEG_PContent()
	_dbggfe._dbfe.EG_PContent = append(_dbggfe._dbfe.EG_PContent, _deef)
	_deef.Hyperlink = _ag.NewCT_Hyperlink()
	return HyperLink{_dbggfe._aefd, _deef.Hyperlink}
}

// AddParagraph adds a paragraph to the endnote.
func (_ggg Endnote) AddParagraph() Paragraph {
	_fdfgg := _ag.NewEG_ContentBlockContent()
	_ggde := len(_ggg._gae.EG_BlockLevelElts[0].EG_ContentBlockContent)
	_ggg._gae.EG_BlockLevelElts[0].EG_ContentBlockContent = append(_ggg._gae.EG_BlockLevelElts[0].EG_ContentBlockContent, _fdfgg)
	_fdcg := _ag.NewCT_P()
	var _gcfa *_ag.CT_String
	if _ggde != 0 {
		_bcafc := len(_ggg._gae.EG_BlockLevelElts[0].EG_ContentBlockContent[_ggde-1].P)
		_gcfa = _ggg._gae.EG_BlockLevelElts[0].EG_ContentBlockContent[_ggde-1].P[_bcafc-1].PPr.PStyle
	} else {
		_gcfa = _ag.NewCT_String()
		_gcfa.ValAttr = "\u0045n\u0064\u006e\u006f\u0074\u0065"
	}
	_fdfgg.P = append(_fdfgg.P, _fdcg)
	_beec := Paragraph{_ggg._ddcec, _fdcg}
	_beec._dbfe.PPr = _ag.NewCT_PPr()
	_beec._dbfe.PPr.PStyle = _gcfa
	_beec._dbfe.PPr.RPr = _ag.NewCT_ParaRPr()
	return _beec
}

// RunProperties controls run styling properties
type RunProperties struct{ _ddda *_ag.CT_RPr }

// AddHeader creates a header associated with the document, but doesn't add it
// to the document for display.
func (_gbf *Document) AddHeader() Header {
	_fd := _ag.NewHdr()
	_gbf._cg = append(_gbf._cg, _fd)
	_dea := _af.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", len(_gbf._cg))
	_gbf._agcc.AddRelationship(_dea, _f.HeaderType)
	_gbf.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_dea, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0068\u0065\u0061\u0064e\u0072\u002b\u0078\u006d\u006c")
	_gbf._gba = append(_gbf._gba, _ec.NewRelationships())
	return Header{_gbf, _fd}
}

// TableStyleProperties are table properties as defined in a style.
type TableStyleProperties struct{ _cgcba *_ag.CT_TblPrBase }

func (_ccbb *Document) insertParagraph(_ffdg Paragraph, _ceef bool) Paragraph {
	if _ccbb._bbgg.Body == nil {
		return _ccbb.AddParagraph()
	}
	_bdfd := _ffdg.X()
	for _, _adac := range _ccbb._bbgg.Body.EG_BlockLevelElts {
		for _, _ffgf := range _adac.EG_ContentBlockContent {
			for _cgf, _ecd := range _ffgf.P {
				if _ecd == _bdfd {
					_gad := _ag.NewCT_P()
					_ffgf.P = append(_ffgf.P, nil)
					if _ceef {
						copy(_ffgf.P[_cgf+1:], _ffgf.P[_cgf:])
						_ffgf.P[_cgf] = _gad
					} else {
						copy(_ffgf.P[_cgf+2:], _ffgf.P[_cgf+1:])
						_ffgf.P[_cgf+1] = _gad
					}
					return Paragraph{_ccbb, _gad}
				}
			}
			for _, _afgd := range _ffgf.Tbl {
				for _, _fdee := range _afgd.EG_ContentRowContent {
					for _, _bce := range _fdee.Tr {
						for _, _egfd := range _bce.EG_ContentCellContent {
							for _, _ecceg := range _egfd.Tc {
								for _, _deaf := range _ecceg.EG_BlockLevelElts {
									for _, _bbga := range _deaf.EG_ContentBlockContent {
										for _feb, _beba := range _bbga.P {
											if _beba == _bdfd {
												_bfbfb := _ag.NewCT_P()
												_bbga.P = append(_bbga.P, nil)
												if _ceef {
													copy(_bbga.P[_feb+1:], _bbga.P[_feb:])
													_bbga.P[_feb] = _bfbfb
												} else {
													copy(_bbga.P[_feb+2:], _bbga.P[_feb+1:])
													_bbga.P[_feb+1] = _bfbfb
												}
												return Paragraph{_ccbb, _bfbfb}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			if _ffgf.Sdt != nil && _ffgf.Sdt.SdtContent != nil && _ffgf.Sdt.SdtContent.P != nil {
				for _fbgf, _fdeeg := range _ffgf.Sdt.SdtContent.P {
					if _fdeeg == _bdfd {
						_bgbfe := _ag.NewCT_P()
						_ffgf.Sdt.SdtContent.P = append(_ffgf.Sdt.SdtContent.P, nil)
						if _ceef {
							copy(_ffgf.Sdt.SdtContent.P[_fbgf+1:], _ffgf.Sdt.SdtContent.P[_fbgf:])
							_ffgf.Sdt.SdtContent.P[_fbgf] = _bgbfe
						} else {
							copy(_ffgf.Sdt.SdtContent.P[_fbgf+2:], _ffgf.Sdt.SdtContent.P[_fbgf+1:])
							_ffgf.Sdt.SdtContent.P[_fbgf+1] = _bgbfe
						}
						return Paragraph{_ccbb, _bgbfe}
					}
				}
			}
		}
	}
	return _ccbb.AddParagraph()
}
func (_ff *Document) InsertTableAfter(relativeTo Paragraph) Table {
	return _ff.insertTable(relativeTo, false)
}

// BoldValue returns the precise nature of the bold setting (unset, off or on).
func (_fagc RunProperties) BoldValue() OnOffValue { return _eea(_fagc._ddda.B) }

// Paragraphs returns all of the paragraphs in the document body including tables.
func (_ffg *Document) Paragraphs() []Paragraph {
	_cfga := []Paragraph{}
	if _ffg._bbgg.Body == nil {
		return nil
	}
	for _, _eba := range _ffg._bbgg.Body.EG_BlockLevelElts {
		for _, _ddfg := range _eba.EG_ContentBlockContent {
			for _, _ecaf := range _ddfg.P {
				_cfga = append(_cfga, Paragraph{_ffg, _ecaf})
			}
		}
	}
	for _, _bfcb := range _ffg.Tables() {
		for _, _fdfg := range _bfcb.Rows() {
			for _, _aga := range _fdfg.Cells() {
				_cfga = append(_cfga, _aga.Paragraphs()...)
			}
		}
	}
	return _cfga
}

// SetWidthAuto sets the the cell width to automatic.
func (_bdf CellProperties) SetWidthAuto() {
	_bdf._dac.TcW = _ag.NewCT_TblWidth()
	_bdf._dac.TcW.TypeAttr = _ag.ST_TblWidthAuto
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (_gccgc Paragraph) SetLineSpacing(d _bf.Distance, rule _ag.ST_LineSpacingRule) {
	_gccgc.ensurePPr()
	if _gccgc._dbfe.PPr.Spacing == nil {
		_gccgc._dbfe.PPr.Spacing = _ag.NewCT_Spacing()
	}
	_cbaba := _gccgc._dbfe.PPr.Spacing
	if rule == _ag.ST_LineSpacingRuleUnset {
		_cbaba.LineRuleAttr = _ag.ST_LineSpacingRuleUnset
		_cbaba.LineAttr = nil
	} else {
		_cbaba.LineRuleAttr = rule
		_cbaba.LineAttr = &_ag.ST_SignedTwipsMeasure{}
		_cbaba.LineAttr.Int64 = _f.Int64(int64(d / _bf.Twips))
	}
}

// X returns the inner wrapped XML type.
func (_bae Cell) X() *_ag.CT_Tc { return _bae._afg }

// X returns the inner wrapped XML type.
func (_eagf Row) X() *_ag.CT_Row { return _eagf._dffc }

// SetHeader sets a section header.
func (_dcae Section) SetHeader(h Header, t _ag.ST_HdrFtr) {
	_gabc := _ag.NewEG_HdrFtrReferences()
	_dcae._cdec.EG_HdrFtrReferences = append(_dcae._cdec.EG_HdrFtrReferences, _gabc)
	_gabc.HeaderReference = _ag.NewCT_HdrFtrRef()
	_gabc.HeaderReference.TypeAttr = t
	_afag := _dcae._cce._agcc.FindRIDForN(h.Index(), _f.HeaderType)
	if _afag == "" {
		_fe.Print("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u0068\u0065\u0061\u0064\u0065r \u0049\u0044")
	}
	_gabc.HeaderReference.IdAttr = _afag
}

// Tables returns the tables defined in the header.
func (_fbbe Header) Tables() []Table {
	_cddee := []Table{}
	if _fbbe._aeaf == nil {
		return nil
	}
	for _, _bef := range _fbbe._aeaf.EG_ContentBlockContent {
		for _, _dbge := range _fbbe._cedg.tables(_bef) {
			_cddee = append(_cddee, _dbge)
		}
	}
	return _cddee
}

// SetBefore sets the spacing that comes before the paragraph.
func (_fcggg ParagraphSpacing) SetBefore(before _bf.Distance) {
	_fcggg._gggda.BeforeAttr = &_dc.ST_TwipsMeasure{}
	_fcggg._gggda.BeforeAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(before / _bf.Twips))
}

// SetUnhideWhenUsed controls if a semi hidden style becomes visible when used.
func (_fdedf Style) SetUnhideWhenUsed(b bool) {
	if b {
		_fdedf._gade.UnhideWhenUsed = _ag.NewCT_OnOff()
	} else {
		_fdedf._gade.UnhideWhenUsed = nil
	}
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_aaac *Document) AddImage(i _ec.Image) (_ec.ImageRef, error) {
	_bbfb := _ec.MakeImageRef(i, &_aaac.DocBase, _aaac._agcc)
	if i.Data == nil && i.Path == "" {
		return _bbfb, _g.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _bbfb, _g.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _bbfb, _g.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	if i.Path != "" {
		_fcbcd := _dd.Add(i.Path)
		if _fcbcd != nil {
			return _bbfb, _fcbcd
		}
	}
	_aaac.Images = append(_aaac.Images, _bbfb)
	_gbc := _af.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_aaac.Images), i.Format)
	_gfeg := _aaac._agcc.AddRelationship(_gbc, _f.ImageType)
	_aaac.ContentTypes.EnsureDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_aaac.ContentTypes.EnsureDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_aaac.ContentTypes.EnsureDefault("\u006a\u0070\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_aaac.ContentTypes.EnsureDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	_aaac.ContentTypes.EnsureDefault(i.Format, "\u0069\u006d\u0061\u0067\u0065\u002f"+i.Format)
	_bbfb.SetRelID(_gfeg.X().IdAttr)
	_bbfb.SetTarget(_gbc)
	return _bbfb, nil
}

// Run is a run of text within a paragraph that shares the same formatting.
type Run struct {
	_ebdd *Document
	_fegg *_ag.CT_R
}

// AddRun adds a run of text to a hyperlink. This is the text that will be linked.
func (_ddeb HyperLink) AddRun() Run {
	_bcee := _ag.NewEG_ContentRunContent()
	_ddeb._dceb.EG_ContentRunContent = append(_ddeb._dceb.EG_ContentRunContent, _bcee)
	_cdee := _ag.NewCT_R()
	_bcee.R = _cdee
	return Run{_ddeb._debd, _cdee}
}

// SetHangingIndent controls the hanging indent of the paragraph.
func (_fefce ParagraphStyleProperties) SetHangingIndent(m _bf.Distance) {
	if _fefce._ddde.Ind == nil {
		_fefce._ddde.Ind = _ag.NewCT_Ind()
	}
	if m == _bf.Zero {
		_fefce._ddde.Ind.HangingAttr = nil
	} else {
		_fefce._ddde.Ind.HangingAttr = &_dc.ST_TwipsMeasure{}
		_fefce._ddde.Ind.HangingAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(m / _bf.Twips))
	}
}

// RowProperties are the properties for a row within a table
type RowProperties struct{ _ebad *_ag.CT_TrPr }

// SetYOffset sets the Y offset for an image relative to the origin.
func (_ddc AnchoredDrawing) SetYOffset(y _bf.Distance) {
	_ddc._de.PositionV.Choice = &_ag.WdCT_PosVChoice{}
	_ddc._de.PositionV.Choice.PosOffset = _f.Int32(int32(y / _bf.EMU))
}

// Fonts allows manipulating a style or run's fonts.
type Fonts struct{ _agaa *_ag.CT_Fonts }

// SetKeepWithNext controls if this paragraph should be kept with the next.
func (_bfcbg ParagraphProperties) SetKeepWithNext(b bool) {
	if !b {
		_bfcbg._cfdd.KeepNext = nil
	} else {
		_bfcbg._cfdd.KeepNext = _ag.NewCT_OnOff()
	}
}

// EastAsiaFont returns the name of run font family for East Asia.
func (_fbfeb RunProperties) EastAsiaFont() string {
	if _acfecb := _fbfeb._ddda.RFonts; _acfecb != nil {
		if _acfecb.EastAsiaAttr != nil {
			return *_acfecb.EastAsiaAttr
		}
	}
	return ""
}

// Paragraphs returns the paragraphs defined in the cell.
func (_fge Cell) Paragraphs() []Paragraph {
	_bb := []Paragraph{}
	for _, _dde := range _fge._afg.EG_BlockLevelElts {
		for _, _dg := range _dde.EG_ContentBlockContent {
			for _, _egf := range _dg.P {
				_bb = append(_bb, Paragraph{_fge._dce, _egf})
			}
		}
	}
	return _bb
}

// X returns the inner wrapped XML type.
func (_dga *Document) X() *_ag.Document { return _dga._bbgg }

// Bold returns true if run font is bold.
func (_fcadf RunProperties) Bold() bool {
	_ffca := _fcadf._ddda
	return _decg(_ffca.B) || _decg(_ffca.BCs)
}

// X returns the inner wrapped XML type.
func (_ecdb TableConditionalFormatting) X() *_ag.CT_TblStylePr { return _ecdb._bdceb }

// SetUpdateFieldsOnOpen controls if fields are recalculated upon opening the
// document. This is useful for things like a table of contents as the library
// only adds the field code and relies on Word/LibreOffice to actually compute
// the content.
func (_cagd Settings) SetUpdateFieldsOnOpen(b bool) {
	if !b {
		_cagd._dcaa.UpdateFields = nil
	} else {
		_cagd._dcaa.UpdateFields = _ag.NewCT_OnOff()
	}
}

// SetOutline sets the run to outlined text.
func (_cbcde RunProperties) SetOutline(b bool) {
	if !b {
		_cbcde._ddda.Outline = nil
	} else {
		_cbcde._ddda.Outline = _ag.NewCT_OnOff()
	}
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (_bbfd TableBorders) SetInsideHorizontal(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_bbfd._egbeb.InsideH = _ag.NewCT_Border()
	_ddecc(_bbfd._egbeb.InsideH, t, c, thickness)
}

// ClearColor clears the text color.
func (_fab RunProperties) ClearColor() { _fab._ddda.Color = nil }
func _babc(_adgd *_ag.CT_Tbl, _gbd *_ag.CT_P, _gfdd bool) *_ag.CT_Tbl {
	for _, _efbg := range _adgd.EG_ContentRowContent {
		for _, _ffb := range _efbg.Tr {
			for _, _fbb := range _ffb.EG_ContentCellContent {
				for _, _bafa := range _fbb.Tc {
					for _efga, _cea := range _bafa.EG_BlockLevelElts {
						for _, _ebc := range _cea.EG_ContentBlockContent {
							for _agb, _bgb := range _ebc.P {
								if _bgb == _gbd {
									_cbf := _ag.NewEG_BlockLevelElts()
									_bfbf := _ag.NewEG_ContentBlockContent()
									_cbf.EG_ContentBlockContent = append(_cbf.EG_ContentBlockContent, _bfbf)
									_eagb := _ag.NewCT_Tbl()
									_bfbf.Tbl = append(_bfbf.Tbl, _eagb)
									_bafa.EG_BlockLevelElts = append(_bafa.EG_BlockLevelElts, nil)
									if _gfdd {
										copy(_bafa.EG_BlockLevelElts[_efga+1:], _bafa.EG_BlockLevelElts[_efga:])
										_bafa.EG_BlockLevelElts[_efga] = _cbf
										if _agb != 0 {
											_bee := _ag.NewEG_BlockLevelElts()
											_bbe := _ag.NewEG_ContentBlockContent()
											_bee.EG_ContentBlockContent = append(_bee.EG_ContentBlockContent, _bbe)
											_bbe.P = _ebc.P[:_agb]
											_bafa.EG_BlockLevelElts = append(_bafa.EG_BlockLevelElts, nil)
											copy(_bafa.EG_BlockLevelElts[_efga+1:], _bafa.EG_BlockLevelElts[_efga:])
											_bafa.EG_BlockLevelElts[_efga] = _bee
										}
										_ebc.P = _ebc.P[_agb:]
									} else {
										copy(_bafa.EG_BlockLevelElts[_efga+2:], _bafa.EG_BlockLevelElts[_efga+1:])
										_bafa.EG_BlockLevelElts[_efga+1] = _cbf
										if _agb != len(_ebc.P)-1 {
											_gdg := _ag.NewEG_BlockLevelElts()
											_acd := _ag.NewEG_ContentBlockContent()
											_gdg.EG_ContentBlockContent = append(_gdg.EG_ContentBlockContent, _acd)
											_acd.P = _ebc.P[_agb+1:]
											_bafa.EG_BlockLevelElts = append(_bafa.EG_BlockLevelElts, nil)
											copy(_bafa.EG_BlockLevelElts[_efga+3:], _bafa.EG_BlockLevelElts[_efga+2:])
											_bafa.EG_BlockLevelElts[_efga+2] = _gdg
										} else {
											_cbfc := _ag.NewEG_BlockLevelElts()
											_egc := _ag.NewEG_ContentBlockContent()
											_cbfc.EG_ContentBlockContent = append(_cbfc.EG_ContentBlockContent, _egc)
											_egc.P = []*_ag.CT_P{_ag.NewCT_P()}
											_bafa.EG_BlockLevelElts = append(_bafa.EG_BlockLevelElts, nil)
											copy(_bafa.EG_BlockLevelElts[_efga+3:], _bafa.EG_BlockLevelElts[_efga+2:])
											_bafa.EG_BlockLevelElts[_efga+2] = _cbfc
										}
										_ebc.P = _ebc.P[:_agb+1]
									}
									return _eagb
								}
							}
							for _, _fbg := range _ebc.Tbl {
								_ccg := _babc(_fbg, _gbd, _gfdd)
								if _ccg != nil {
									return _ccg
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

// Paragraphs returns the paragraphs defined in an endnote.
func (_gec Endnote) Paragraphs() []Paragraph {
	_gfbe := []Paragraph{}
	for _, _cbeb := range _gec.content() {
		for _, _fdbg := range _cbeb.P {
			_gfbe = append(_gfbe, Paragraph{_gec._ddcec, _fdbg})
		}
	}
	return _gfbe
}

// Name returns the name of the bookmark whcih is the document unique ID that
// identifies the bookmark.
func (_dee Bookmark) Name() string { return _dee._ebg.NameAttr }

// Row is a row within a table within a document.
type Row struct {
	_eefac *Document
	_dffc  *_ag.CT_Row
}

// X returns the internally wrapped *wml.CT_SectPr.
func (_ddad Section) X() *_ag.CT_SectPr { return _ddad._cdec }

// VerticalAlign returns the value of run vertical align.
func (_defd RunProperties) VerticalAlignment() _dc.ST_VerticalAlignRun {
	if _fdeae := _defd._ddda.VertAlign; _fdeae != nil {
		return _fdeae.ValAttr
	}
	return 0
}

// TableProperties are the properties for a table within a document
type TableProperties struct{ _gfaa *_ag.CT_TblPr }

// Color returns the style's Color.
func (_fbgaf RunProperties) Color() Color {
	if _fbgaf._ddda.Color == nil {
		_fbgaf._ddda.Color = _ag.NewCT_Color()
	}
	return Color{_fbgaf._ddda.Color}
}

// Text returns the underlying tet in the run.
func (_ecfb Run) Text() string {
	if len(_ecfb._fegg.EG_RunInnerContent) == 0 {
		return ""
	}
	_cbbd := _ca.Buffer{}
	for _, _fdeegf := range _ecfb._fegg.EG_RunInnerContent {
		if _fdeegf.T != nil {
			_cbbd.WriteString(_fdeegf.T.Content)
		}
		if _fdeegf.Tab != nil {
			_cbbd.WriteByte('\t')
		}
	}
	return _cbbd.String()
}

// RightToLeft returns true if paragraph text goes from right to left.
func (_caca ParagraphProperties) RightToLeft() bool { return _decg(_caca._cfdd.RPr.Rtl) }

// GetFooter gets a section Footer for given type
func (_cbfcg Section) GetFooter(t _ag.ST_HdrFtr) (Footer, bool) {
	for _, _egbc := range _cbfcg._cdec.EG_HdrFtrReferences {
		if _egbc.FooterReference.TypeAttr == t {
			for _, _bfbc := range _cbfcg._cce.Footers() {
				_agec := _cbfcg._cce._agcc.FindRIDForN(_bfbc.Index(), _f.FooterType)
				if _agec == _egbc.FooterReference.IdAttr {
					return _bfbc, true
				}
			}
		}
	}
	return Footer{}, false
}

// SetToolTip sets the tooltip text for a hyperlink.
func (_bedb HyperLink) SetToolTip(text string) {
	if text == "" {
		_bedb._dceb.TooltipAttr = nil
	} else {
		_bedb._dceb.TooltipAttr = _f.String(text)
	}
}

// SetShading controls the cell shading.
func (_edfa CellProperties) SetShading(shd _ag.ST_Shd, foreground, fill _ed.Color) {
	if shd == _ag.ST_ShdUnset {
		_edfa._dac.Shd = nil
	} else {
		_edfa._dac.Shd = _ag.NewCT_Shd()
		_edfa._dac.Shd.ValAttr = shd
		_edfa._dac.Shd.ColorAttr = &_ag.ST_HexColor{}
		if foreground.IsAuto() {
			_edfa._dac.Shd.ColorAttr.ST_HexColorAuto = _ag.ST_HexColorAutoAuto
		} else {
			_edfa._dac.Shd.ColorAttr.ST_HexColorRGB = foreground.AsRGBString()
		}
		_edfa._dac.Shd.FillAttr = &_ag.ST_HexColor{}
		if fill.IsAuto() {
			_edfa._dac.Shd.FillAttr.ST_HexColorAuto = _ag.ST_HexColorAutoAuto
		} else {
			_edfa._dac.Shd.FillAttr.ST_HexColorRGB = fill.AsRGBString()
		}
	}
}

// DoubleStrike returns true if run is double striked.
func (_cecg RunProperties) DoubleStrike() bool { return _decg(_cecg._ddda.Dstrike) }

// NumberingLevel is the definition for numbering for a particular level within
// a NumberingDefinition.
type NumberingLevel struct{ _bdce *_ag.CT_Lvl }

// GetImage returns the ImageRef associated with an InlineDrawing.
func (_dcag InlineDrawing) GetImage() (_ec.ImageRef, bool) {
	_cagf := _dcag._egff.Graphic.GraphicData.Any
	if len(_cagf) > 0 {
		_dgd, _fefe := _cagf[0].(*_afd.Pic)
		if _fefe {
			if _dgd.BlipFill != nil && _dgd.BlipFill.Blip != nil && _dgd.BlipFill.Blip.EmbedAttr != nil {
				return _dcag._agaeg.GetImageByRelID(*_dgd.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return _ec.ImageRef{}, false
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (_cfa CellBorders) SetInsideHorizontal(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_cfa._eaa.InsideH = _ag.NewCT_Border()
	_ddecc(_cfa._eaa.InsideH, t, c, thickness)
}

// SetXOffset sets the X offset for an image relative to the origin.
func (_ac AnchoredDrawing) SetXOffset(x _bf.Distance) {
	_ac._de.PositionH.Choice = &_ag.WdCT_PosHChoice{}
	_ac._de.PositionH.Choice.PosOffset = _f.Int32(int32(x / _bf.EMU))
}

// SetLeftPct sets the cell left margin
func (_fbef CellMargins) SetLeftPct(pct float64) {
	_fbef._cacb.Left = _ag.NewCT_TblWidth()
	_adg(_fbef._cacb.Left, pct)
}

// X returns the inner wrapped XML type.
func (_ccfc NumberingDefinition) X() *_ag.CT_AbstractNum { return _ccfc._cgfd }

// Type returns the type of the field.
func (_ffffe FormField) Type() FormFieldType {
	if _ffffe._bffef.TextInput != nil {
		return FormFieldTypeText
	} else if _ffffe._bffef.CheckBox != nil {
		return FormFieldTypeCheckBox
	} else if _ffffe._bffef.DdList != nil {
		return FormFieldTypeDropDown
	}
	return FormFieldTypeUnknown
}

// IsItalic returns true if the run has been set to italics.
func (_eacea RunProperties) IsItalic() bool { return _eacea.ItalicValue() == OnOffValueOn }

// InsertParagraphBefore adds a new empty paragraph before the relativeTo
// paragraph.
func (_gcc *Document) InsertParagraphBefore(relativeTo Paragraph) Paragraph {
	return _gcc.insertParagraph(relativeTo, true)
}

// CellProperties are a table cells properties within a document.
type CellProperties struct{ _dac *_ag.CT_TcPr }

// X returns the inner wrapped XML type.
func (_agdd Styles) X() *_ag.Styles { return _agdd._faabg }

// X returns the inner wrapped XML type.
func (_cdag TableProperties) X() *_ag.CT_TblPr { return _cdag._gfaa }

// Header is a header for a document section.
type Header struct {
	_cedg *Document
	_aeaf *_ag.Hdr
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (_bade RunProperties) CharacterSpacingMeasure() string {
	if _ggca := _bade._ddda.Spacing; _ggca != nil {
		_bbbf := _ggca.ValAttr
		if _bbbf.ST_UniversalMeasure != nil {
			return *_bbbf.ST_UniversalMeasure
		}
	}
	return ""
}

// SetTop sets the top border to a specified type, color and thickness.
func (_beb CellBorders) SetTop(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_beb._eaa.Top = _ag.NewCT_Border()
	_ddecc(_beb._eaa.Top, t, c, thickness)
}

// Strike returns true if paragraph is striked.
func (_eded ParagraphProperties) Strike() bool { return _decg(_eded._cfdd.RPr.Strike) }

// Styles returns all styles.
func (_eaeb Styles) Styles() []Style {
	_gggb := []Style{}
	for _, _cbebb := range _eaeb._faabg.Style {
		_gggb = append(_gggb, Style{_cbebb})
	}
	return _gggb
}

// X returns the inner wrapped XML type.
func (_gd CellProperties) X() *_ag.CT_TcPr { return _gd._dac }

// SetNextStyle sets the style that the next paragraph will use.
func (_bfccg Style) SetNextStyle(name string) {
	if name == "" {
		_bfccg._gade.Next = nil
	} else {
		_bfccg._gade.Next = _ag.NewCT_String()
		_bfccg._gade.Next.ValAttr = name
	}
}

// SetTextWrapSquare sets the text wrap to square with a given wrap type.
func (_df AnchoredDrawing) SetTextWrapSquare(t _ag.WdST_WrapText) {
	_df._de.Choice = &_ag.WdEG_WrapTypeChoice{}
	_df._de.Choice.WrapSquare = _ag.NewWdCT_WrapSquare()
	_df._de.Choice.WrapSquare.WrapTextAttr = t
}

// AddFootnote will create a new footnote and attach it to the Paragraph in the
// location at the end of the previous run (footnotes create their own run within
// the paragraph). The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the footnote later.
func (_agff Paragraph) AddFootnote(text string) Footnote {
	var _deafe int64
	if _agff._aefd.HasFootnotes() {
		for _, _bacce := range _agff._aefd.Footnotes() {
			if _bacce.id() > _deafe {
				_deafe = _bacce.id()
			}
		}
		_deafe++
	} else {
		_deafe = 0
		_agff._aefd._ab = &_ag.Footnotes{}
		_agff._aefd._ab.CT_Footnotes = _ag.CT_Footnotes{}
		_agff._aefd._ab.Footnote = make([]*_ag.CT_FtnEdn, 0)
	}
	_bgbfeg := _ag.NewCT_FtnEdn()
	_dbbd := _ag.NewCT_FtnEdnRef()
	_dbbd.IdAttr = _deafe
	_agff._aefd._ab.CT_Footnotes.Footnote = append(_agff._aefd._ab.CT_Footnotes.Footnote, _bgbfeg)
	_gcdd := _agff.AddRun()
	_gcacf := _gcdd.Properties()
	_gcacf.SetStyle("\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065\u0041n\u0063\u0068\u006f\u0072")
	_gcdd._fegg.EG_RunInnerContent = []*_ag.EG_RunInnerContent{_ag.NewEG_RunInnerContent()}
	_gcdd._fegg.EG_RunInnerContent[0].FootnoteReference = _dbbd
	_fgfc := Footnote{_agff._aefd, _bgbfeg}
	_fgfc._faaeg.IdAttr = _deafe
	_fgfc._faaeg.EG_BlockLevelElts = []*_ag.EG_BlockLevelElts{_ag.NewEG_BlockLevelElts()}
	_ccbf := _fgfc.AddParagraph()
	_ccbf.Properties().SetStyle("\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065")
	_ccbf._dbfe.PPr.RPr = _ag.NewCT_ParaRPr()
	_cfeb := _ccbf.AddRun()
	_cfeb.AddTab()
	_cfeb.AddText(text)
	return _fgfc
}

// SetPageMargins sets the page margins for a section
func (_fcae Section) SetPageMargins(top, right, bottom, left, header, footer, gutter _bf.Distance) {
	_ffad := _ag.NewCT_PageMar()
	_ffad.TopAttr.Int64 = _f.Int64(int64(top / _bf.Twips))
	_ffad.BottomAttr.Int64 = _f.Int64(int64(bottom / _bf.Twips))
	_ffad.RightAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(right / _bf.Twips))
	_ffad.LeftAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(left / _bf.Twips))
	_ffad.HeaderAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(header / _bf.Twips))
	_ffad.FooterAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(footer / _bf.Twips))
	_ffad.GutterAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(gutter / _bf.Twips))
	_fcae._cdec.PgMar = _ffad
}

// Clear clears all content within a header
func (_dbca Header) Clear() { _dbca._aeaf.EG_ContentBlockContent = nil }

// SetSize sets size attribute for a FormFieldTypeCheckBox in pt.
func (_aadgf FormField) SetSize(size uint64) {
	size *= 2
	if _aadgf._bffef.CheckBox != nil {
		_aadgf._bffef.CheckBox.Choice = _ag.NewCT_FFCheckBoxChoice()
		_aadgf._bffef.CheckBox.Choice.Size = _ag.NewCT_HpsMeasure()
		_aadgf._bffef.CheckBox.Choice.Size.ValAttr = _ag.ST_HpsMeasure{ST_UnsignedDecimalNumber: &size}
	}
}

// InitializeDefault constructs a default numbering.
func (_cdfe Numbering) InitializeDefault() {
	_ebbd := _ag.NewCT_AbstractNum()
	_ebbd.MultiLevelType = _ag.NewCT_MultiLevelType()
	_ebbd.MultiLevelType.ValAttr = _ag.ST_MultiLevelTypeHybridMultilevel
	_cdfe._eefa.AbstractNum = append(_cdfe._eefa.AbstractNum, _ebbd)
	_ebbd.AbstractNumIdAttr = 1
	const _egdcc = 720
	const _acgf = 720
	const _cfdc = 360
	for _effb := 0; _effb < 9; _effb++ {
		_bdee := _ag.NewCT_Lvl()
		_bdee.IlvlAttr = int64(_effb)
		_bdee.Start = _ag.NewCT_DecimalNumber()
		_bdee.Start.ValAttr = 1
		_bdee.NumFmt = _ag.NewCT_NumFmt()
		_bdee.NumFmt.ValAttr = _ag.ST_NumberFormatBullet
		_bdee.Suff = _ag.NewCT_LevelSuffix()
		_bdee.Suff.ValAttr = _ag.ST_LevelSuffixNothing
		_bdee.LvlText = _ag.NewCT_LevelText()
		_bdee.LvlText.ValAttr = _f.String("\uf0b7")
		_bdee.LvlJc = _ag.NewCT_Jc()
		_bdee.LvlJc.ValAttr = _ag.ST_JcLeft
		_bdee.RPr = _ag.NewCT_RPr()
		_bdee.RPr.RFonts = _ag.NewCT_Fonts()
		_bdee.RPr.RFonts.AsciiAttr = _f.String("\u0053\u0079\u006d\u0062\u006f\u006c")
		_bdee.RPr.RFonts.HAnsiAttr = _f.String("\u0053\u0079\u006d\u0062\u006f\u006c")
		_bdee.RPr.RFonts.HintAttr = _ag.ST_HintDefault
		_bdee.PPr = _ag.NewCT_PPrGeneral()
		_eadgd := int64(_effb*_acgf + _egdcc)
		_bdee.PPr.Ind = _ag.NewCT_Ind()
		_bdee.PPr.Ind.LeftAttr = &_ag.ST_SignedTwipsMeasure{}
		_bdee.PPr.Ind.LeftAttr.Int64 = _f.Int64(_eadgd)
		_bdee.PPr.Ind.HangingAttr = &_dc.ST_TwipsMeasure{}
		_bdee.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(_cfdc))
		_ebbd.Lvl = append(_ebbd.Lvl, _bdee)
	}
	_ccbbg := _ag.NewCT_Num()
	_ccbbg.NumIdAttr = 1
	_ccbbg.AbstractNumId = _ag.NewCT_DecimalNumber()
	_ccbbg.AbstractNumId.ValAttr = 1
	_cdfe._eefa.Num = append(_cdfe._eefa.Num, _ccbbg)
}

// SetImprint sets the run to imprinted text.
func (_dgaa RunProperties) SetImprint(b bool) {
	if !b {
		_dgaa._ddda.Imprint = nil
	} else {
		_dgaa._ddda.Imprint = _ag.NewCT_OnOff()
	}
}
func (_ebae Endnote) id() int64 { return _ebae._gae.IdAttr }

// X returns the inner wrapped XML type.
func (_agd Bookmark) X() *_ag.CT_Bookmark { return _agd._ebg }

// SetVerticalAlignment controls the vertical alignment of the run, this is used
// to control if text is superscript/subscript.
func (_aegdb RunProperties) SetVerticalAlignment(v _dc.ST_VerticalAlignRun) {
	if v == _dc.ST_VerticalAlignRunUnset {
		_aegdb._ddda.VertAlign = nil
	} else {
		_aegdb._ddda.VertAlign = _ag.NewCT_VerticalAlignRun()
		_aegdb._ddda.VertAlign.ValAttr = v
	}
}

// SetStartPct sets the cell start margin
func (_ggd CellMargins) SetStartPct(pct float64) {
	_ggd._cacb.Start = _ag.NewCT_TblWidth()
	_adg(_ggd._cacb.Start, pct)
}

// AddDrawingAnchored adds an anchored (floating) drawing from an ImageRef.
func (_aceec Run) AddDrawingAnchored(img _ec.ImageRef) (AnchoredDrawing, error) {
	_dggbg := _aceec.newIC()
	_dggbg.Drawing = _ag.NewCT_Drawing()
	_fbfbe := _ag.NewWdAnchor()
	_dcdge := AnchoredDrawing{_aceec._ebdd, _fbfbe}
	_fbfbe.SimplePosAttr = _f.Bool(false)
	_fbfbe.AllowOverlapAttr = true
	_fbfbe.CNvGraphicFramePr = _ea.NewCT_NonVisualGraphicFrameProperties()
	_dggbg.Drawing.Anchor = append(_dggbg.Drawing.Anchor, _fbfbe)
	_fbfbe.Graphic = _ea.NewGraphic()
	_fbfbe.Graphic.GraphicData = _ea.NewCT_GraphicalObjectData()
	_fbfbe.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"
	_fbfbe.SimplePos.XAttr.ST_CoordinateUnqualified = _f.Int64(0)
	_fbfbe.SimplePos.YAttr.ST_CoordinateUnqualified = _f.Int64(0)
	_fbfbe.PositionH.RelativeFromAttr = _ag.WdST_RelFromHPage
	_fbfbe.PositionH.Choice = &_ag.WdCT_PosHChoice{}
	_fbfbe.PositionH.Choice.PosOffset = _f.Int32(0)
	_fbfbe.PositionV.RelativeFromAttr = _ag.WdST_RelFromVPage
	_fbfbe.PositionV.Choice = &_ag.WdCT_PosVChoice{}
	_fbfbe.PositionV.Choice.PosOffset = _f.Int32(0)
	_fbfbe.Extent.CxAttr = int64(float64(img.Size().X*_bf.Pixel72) / _bf.EMU)
	_fbfbe.Extent.CyAttr = int64(float64(img.Size().Y*_bf.Pixel72) / _bf.EMU)
	_fbfbe.Choice = &_ag.WdEG_WrapTypeChoice{}
	_fbfbe.Choice.WrapSquare = _ag.NewWdCT_WrapSquare()
	_fbfbe.Choice.WrapSquare.WrapTextAttr = _ag.WdST_WrapTextBothSides
	_cagfd := 0x7FFFFFFF & _fg.Uint32()
	_fbfbe.DocPr.IdAttr = _cagfd
	_dfee := _afd.NewPic()
	_dfee.NvPicPr.CNvPr.IdAttr = _cagfd
	_bfgc := img.RelID()
	if _bfgc == "" {
		return _dcdge, _g.New("\u0063\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020\u0074\u006f\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0020\u0072\u0065l\u0061\u0074\u0069o\u006e\u0073")
	}
	_fbfbe.Graphic.GraphicData.Any = append(_fbfbe.Graphic.GraphicData.Any, _dfee)
	_dfee.BlipFill = _ea.NewCT_BlipFillProperties()
	_dfee.BlipFill.Blip = _ea.NewCT_Blip()
	_dfee.BlipFill.Blip.EmbedAttr = &_bfgc
	_dfee.BlipFill.Stretch = _ea.NewCT_StretchInfoProperties()
	_dfee.BlipFill.Stretch.FillRect = _ea.NewCT_RelativeRect()
	_dfee.SpPr = _ea.NewCT_ShapeProperties()
	_dfee.SpPr.Xfrm = _ea.NewCT_Transform2D()
	_dfee.SpPr.Xfrm.Off = _ea.NewCT_Point2D()
	_dfee.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _f.Int64(0)
	_dfee.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _f.Int64(0)
	_dfee.SpPr.Xfrm.Ext = _ea.NewCT_PositiveSize2D()
	_dfee.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * _bf.Point)
	_dfee.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * _bf.Point)
	_dfee.SpPr.PrstGeom = _ea.NewCT_PresetGeometry2D()
	_dfee.SpPr.PrstGeom.PrstAttr = _ea.ST_ShapeTypeRect
	return _dcdge, nil
}

// AddTextInput adds text input form field to the paragraph and returns it.
func (_abcbe Paragraph) AddTextInput(name string) FormField {
	_cadgb := _abcbe.addFldCharsForField(name, "\u0046\u004f\u0052\u004d\u0054\u0045\u0058\u0054")
	_cadgb._bffef.TextInput = _ag.NewCT_FFTextInput()
	return _cadgb
}

// SetFirstColumn controls the conditional formatting for the first column in a table.
func (_dgdg TableLook) SetFirstColumn(on bool) {
	if !on {
		_dgdg._cfbff.FirstColumnAttr = &_dc.ST_OnOff{}
		_dgdg._cfbff.FirstColumnAttr.ST_OnOff1 = _dc.ST_OnOff1Off
	} else {
		_dgdg._cfbff.FirstColumnAttr = &_dc.ST_OnOff{}
		_dgdg._cfbff.FirstColumnAttr.ST_OnOff1 = _dc.ST_OnOff1On
	}
}

// New constructs an empty document that content can be added to.
func New() *Document {
	_bcg := &Document{_bbgg: _ag.NewDocument()}
	_bcg.ContentTypes = _ec.NewContentTypes()
	_bcg._bbgg.Body = _ag.NewCT_Body()
	_bcg._bbgg.ConformanceAttr = _dc.ST_ConformanceClassTransitional
	_bcg._agcc = _ec.NewRelationships()
	_bcg.AppProperties = _ec.NewAppProperties()
	_bcg.CoreProperties = _ec.NewCoreProperties()
	_bcg.ContentTypes.AddOverride("\u002fw\u006fr\u0064\u002f\u0064\u006f\u0063u\u006d\u0065n\u0074\u002e\u0078\u006d\u006c", "\u0061p\u0070\u006c\u0069c\u0061\u0074\u0069o\u006e/v\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072dp\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u006d\u006c\u002e\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u002e\u006d\u0061\u0069\u006e\u002bx\u006d\u006c")
	_bcg.Settings = NewSettings()
	_bcg._agcc.AddRelationship("\u0073\u0065\u0074t\u0069\u006e\u0067\u0073\u002e\u0078\u006d\u006c", _f.SettingsType)
	_bcg.ContentTypes.AddOverride("\u002fw\u006fr\u0064\u002f\u0073\u0065\u0074t\u0069\u006eg\u0073\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069o\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006dl\u0066\u006f\u0072\u006da\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c.\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002b\u0078\u006d\u006c")
	_bcg.Rels = _ec.NewRelationships()
	_bcg.Rels.AddRelationship(_f.RelativeFilename(_f.DocTypeDocument, "", _f.CorePropertiesType, 0), _f.CorePropertiesType)
	_bcg.Rels.AddRelationship("\u0064\u006fc\u0050\u0072\u006fp\u0073\u002f\u0061\u0070\u0070\u002e\u0078\u006d\u006c", _f.ExtendedPropertiesType)
	_bcg.Rels.AddRelationship("\u0077\u006f\u0072\u0064\u002f\u0064\u006f\u0063\u0075\u006d\u0065\u006et\u002e\u0078\u006d\u006c", _f.OfficeDocumentType)
	_bcg.Numbering = NewNumbering()
	_bcg.Numbering.InitializeDefault()
	_bcg.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072d/\u006e\u0075\u006d\u0062\u0065\u0072\u0069\u006e\u0067\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002f\u0076n\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063e\u0073\u0073\u0069\u006e\u0067\u006d\u006c\u002e\u006e\u0075\u006d\u0062e\u0072\u0069\u006e\u0067\u002b\u0078m\u006c")
	_bcg._agcc.AddRelationship("\u006e\u0075\u006d\u0062\u0065\u0072\u0069\u006e\u0067\u002e\u0078\u006d\u006c", _f.NumberingType)
	_bcg.Styles = NewStyles()
	_bcg.Styles.InitializeDefault()
	_bcg.ContentTypes.AddOverride("\u002f\u0077o\u0072\u0064\u002fs\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0073\u0074\u0079\u006ce\u0073\u002b\u0078\u006d\u006c")
	_bcg._agcc.AddRelationship("\u0073\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", _f.StylesType)
	_bcg._bbgg.Body = _ag.NewCT_Body()
	return _bcg
}

// Underline returns the type of paragraph underline.
func (_cdfae ParagraphProperties) Underline() _ag.ST_Underline {
	if _feeg := _cdfae._cfdd.RPr.U; _feeg != nil {
		return _feeg.ValAttr
	}
	return 0
}

// SetName sets the name of the bookmark. This is the name that is used to
// reference the bookmark from hyperlinks.
func (_deb Bookmark) SetName(name string) { _deb._ebg.NameAttr = name }

// GetColor returns the color.Color object representing the run color.
func (_gada ParagraphProperties) GetColor() _ed.Color {
	if _ffaac := _gada._cfdd.RPr.Color; _ffaac != nil {
		_bgfb := _ffaac.ValAttr
		if _bgfb.ST_HexColorRGB != nil {
			return _ed.FromHex(*_bgfb.ST_HexColorRGB)
		}
	}
	return _ed.Color{}
}

// StyleID returns the style ID.
func (_gfeff Style) StyleID() string {
	if _gfeff._gade.StyleIdAttr == nil {
		return ""
	}
	return *_gfeff._gade.StyleIdAttr
}

// SetStyle sets the style of a paragraph.
func (_dbbdd ParagraphProperties) SetStyle(s string) {
	if s == "" {
		_dbbdd._cfdd.PStyle = nil
	} else {
		_dbbdd._cfdd.PStyle = _ag.NewCT_String()
		_dbbdd._cfdd.PStyle.ValAttr = s
	}
}

// Cells returns the cells defined in the table.
func (_cfbgg Row) Cells() []Cell {
	_dagdc := []Cell{}
	for _, _aeac := range _cfbgg._dffc.EG_ContentCellContent {
		for _, _fdcd := range _aeac.Tc {
			_dagdc = append(_dagdc, Cell{_cfbgg._eefac, _fdcd})
		}
		if _aeac.Sdt != nil && _aeac.Sdt.SdtContent != nil {
			for _, _fbfb := range _aeac.Sdt.SdtContent.Tc {
				_dagdc = append(_dagdc, Cell{_cfbgg._eefac, _fbfb})
			}
		}
	}
	return _dagdc
}

// SetTop sets the top border to a specified type, color and thickness.
func (_cfeg TableBorders) SetTop(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_cfeg._egbeb.Top = _ag.NewCT_Border()
	_ddecc(_cfeg._egbeb.Top, t, c, thickness)
}

// SetVerticalMerge controls the vertical merging of cells.
func (_gf CellProperties) SetVerticalMerge(mergeVal _ag.ST_Merge) {
	if mergeVal == _ag.ST_MergeUnset {
		_gf._dac.VMerge = nil
	} else {
		_gf._dac.VMerge = _ag.NewCT_VMerge()
		_gf._dac.VMerge.ValAttr = mergeVal
	}
}

// StructuredDocumentTags returns the structured document tags in the document
// which are commonly used in document templates.
func (_efba *Document) StructuredDocumentTags() []StructuredDocumentTag {
	_cec := []StructuredDocumentTag{}
	for _, _cbc := range _efba._bbgg.Body.EG_BlockLevelElts {
		for _, _fdea := range _cbc.EG_ContentBlockContent {
			if _fdea.Sdt != nil {
				_cec = append(_cec, StructuredDocumentTag{_efba, _fdea.Sdt})
			}
		}
	}
	return _cec
}

// RunProperties returns the RunProperties controlling numbering level font, etc.
func (_aceb NumberingLevel) RunProperties() RunProperties {
	if _aceb._bdce.RPr == nil {
		_aceb._bdce.RPr = _ag.NewCT_RPr()
	}
	return RunProperties{_aceb._bdce.RPr}
}

// AddField adds a field (automatically computed text) to the document.
func (_bedg Run) AddField(code string) { _bedg.AddFieldWithFormatting(code, "", true) }

// X returns the inner wrapped XML type.
func (_abfb TableStyleProperties) X() *_ag.CT_TblPrBase { return _abfb._cgcba }

const (
	FormFieldTypeUnknown FormFieldType = iota
	FormFieldTypeText
	FormFieldTypeCheckBox
	FormFieldTypeDropDown
)

func (_eaaa Footnote) id() int64 { return _eaaa._faaeg.IdAttr }

// TableWidth controls width values in table settings.
type TableWidth struct{ _eccef *_ag.CT_TblWidth }

// Footnotes returns the footnotes defined in the document.
func (_cfbg *Document) Footnotes() []Footnote {
	_dbcc := []Footnote{}
	for _, _edef := range _cfbg._ab.CT_Footnotes.Footnote {
		_dbcc = append(_dbcc, Footnote{_cfbg, _edef})
	}
	return _dbcc
}

// AddCheckBox adds checkbox form field to the paragraph and returns it.
func (_caafd Paragraph) AddCheckBox(name string) FormField {
	_bacb := _caafd.addFldCharsForField(name, "\u0046\u004f\u0052M\u0043\u0048\u0045\u0043\u004b\u0042\u004f\u0058")
	_bacb._bffef.CheckBox = _ag.NewCT_FFCheckBox()
	return _bacb
}

// X returns the inner wrapped type
func (_ae CellBorders) X() *_ag.CT_TcBorders { return _ae._eaa }

// Strike returns true if run is striked.
func (_caae RunProperties) Strike() bool { return _decg(_caae._ddda.Strike) }

// SetFontFamily sets the Ascii & HAnsi fonly family for a run.
func (_dccg RunProperties) SetFontFamily(family string) {
	if _dccg._ddda.RFonts == nil {
		_dccg._ddda.RFonts = _ag.NewCT_Fonts()
	}
	_dccg._ddda.RFonts.AsciiAttr = _f.String(family)
	_dccg._ddda.RFonts.HAnsiAttr = _f.String(family)
	_dccg._ddda.RFonts.EastAsiaAttr = _f.String(family)
}

// SetBottom sets the cell bottom margin
func (_eab CellMargins) SetBottom(d _bf.Distance) {
	_eab._cacb.Bottom = _ag.NewCT_TblWidth()
	_edf(_eab._cacb.Bottom, d)
}

// RemoveParagraph removes a paragraph from the endnote.
func (_gfdg Endnote) RemoveParagraph(p Paragraph) {
	for _, _fbea := range _gfdg.content() {
		for _fbbc, _eeff := range _fbea.P {
			if _eeff == p._dbfe {
				copy(_fbea.P[_fbbc:], _fbea.P[_fbbc+1:])
				_fbea.P = _fbea.P[0 : len(_fbea.P)-1]
				return
			}
		}
	}
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_fed Footer) AddImage(i _ec.Image) (_ec.ImageRef, error) {
	var _fcbbe _ec.Relationships
	for _eadf, _bafaf := range _fed._dagd._bbc {
		if _bafaf == _fed._edec {
			_fcbbe = _fed._dagd._dbc[_eadf]
		}
	}
	_bed := _ec.MakeImageRef(i, &_fed._dagd.DocBase, _fcbbe)
	if i.Data == nil && i.Path == "" {
		return _bed, _g.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _bed, _g.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _bed, _g.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	_fed._dagd.Images = append(_fed._dagd.Images, _bed)
	_agfa := _af.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_fed._dagd.Images), i.Format)
	_adfce := _fcbbe.AddRelationship(_agfa, _f.ImageType)
	_bed.SetRelID(_adfce.X().IdAttr)
	return _bed, nil
}
func _ddecc(_fgdf *_ag.CT_Border, _bada _ag.ST_Border, _ffgc _ed.Color, _ggcb _bf.Distance) {
	_fgdf.ValAttr = _bada
	_fgdf.ColorAttr = &_ag.ST_HexColor{}
	if _ffgc.IsAuto() {
		_fgdf.ColorAttr.ST_HexColorAuto = _ag.ST_HexColorAutoAuto
	} else {
		_fgdf.ColorAttr.ST_HexColorRGB = _ffgc.AsRGBString()
	}
	if _ggcb != _bf.Zero {
		_fgdf.SzAttr = _f.Uint64(uint64(_ggcb / _bf.Point * 8))
	}
}

// InsertRowAfter inserts a row after another row
func (_dfebf Table) InsertRowAfter(r Row) Row {
	for _bfbcd, _bbbff := range _dfebf._fdeda.EG_ContentRowContent {
		if len(_bbbff.Tr) > 0 && r.X() == _bbbff.Tr[0] {
			_ecdc := _ag.NewEG_ContentRowContent()
			if len(_dfebf._fdeda.EG_ContentRowContent) < _bfbcd+2 {
				return _dfebf.AddRow()
			}
			_dfebf._fdeda.EG_ContentRowContent = append(_dfebf._fdeda.EG_ContentRowContent, nil)
			copy(_dfebf._fdeda.EG_ContentRowContent[_bfbcd+2:], _dfebf._fdeda.EG_ContentRowContent[_bfbcd+1:])
			_dfebf._fdeda.EG_ContentRowContent[_bfbcd+1] = _ecdc
			_ggaag := _ag.NewCT_Row()
			_ecdc.Tr = append(_ecdc.Tr, _ggaag)
			return Row{_dfebf._aebg, _ggaag}
		}
	}
	return _dfebf.AddRow()
}

// SetHAlignment sets the horizontal alignment for an anchored image.
func (_fbf AnchoredDrawing) SetHAlignment(h _ag.WdST_AlignH) {
	_fbf._de.PositionH.Choice = &_ag.WdCT_PosHChoice{}
	_fbf._de.PositionH.Choice.Align = h
}

// SetVerticalAlignment sets the vertical alignment of content within a table cell.
func (_ddg CellProperties) SetVerticalAlignment(align _ag.ST_VerticalJc) {
	if align == _ag.ST_VerticalJcUnset {
		_ddg._dac.VAlign = nil
	} else {
		_ddg._dac.VAlign = _ag.NewCT_VerticalJc()
		_ddg._dac.VAlign.ValAttr = align
	}
}

// X returns the inner wrapped XML type.
func (_ceag InlineDrawing) X() *_ag.WdInline { return _ceag._egff }

// Paragraphs returns the paragraphs within a structured document tag.
func (_baabf StructuredDocumentTag) Paragraphs() []Paragraph {
	if _baabf._efgce.SdtContent == nil {
		return nil
	}
	_daeg := []Paragraph{}
	for _, _geed := range _baabf._efgce.SdtContent.P {
		_daeg = append(_daeg, Paragraph{_baabf._eccdf, _geed})
	}
	return _daeg
}
func (_efdd Paragraph) addFldChar() *_ag.CT_FldChar {
	_baea := _efdd.AddRun()
	_fggf := _baea.X()
	_edeea := _ag.NewEG_RunInnerContent()
	_ceefa := _ag.NewCT_FldChar()
	_edeea.FldChar = _ceefa
	_fggf.EG_RunInnerContent = append(_fggf.EG_RunInnerContent, _edeea)
	return _ceefa
}

// TableProperties returns the table style properties.
func (_abaa Style) TableProperties() TableStyleProperties {
	if _abaa._gade.TblPr == nil {
		_abaa._gade.TblPr = _ag.NewCT_TblPrBase()
	}
	return TableStyleProperties{_abaa._gade.TblPr}
}

// X returns the inner wrapped XML type.
func (_bdceba TableLook) X() *_ag.CT_TblLook { return _bdceba._cfbff }

// SetTextWrapNone unsets text wrapping so the image can float on top of the
// text. When used in conjunction with X/Y Offset relative to the page it can be
// used to place a logo at the top of a page at an absolute position that
// doesn't interfere with text.
func (_bge AnchoredDrawing) SetTextWrapNone() {
	_bge._de.Choice = &_ag.WdEG_WrapTypeChoice{}
	_bge._de.Choice.WrapNone = _ag.NewWdCT_WrapNone()
}

// SetPrimaryStyle marks the style as a primary style.
func (_beda Style) SetPrimaryStyle(b bool) {
	if b {
		_beda._gade.QFormat = _ag.NewCT_OnOff()
	} else {
		_beda._gade.QFormat = nil
	}
}

// OpenTemplate opens a document, removing all content so it can be used as a
// template.  Since Word removes unused styles from a document upon save, to
// create a template in Word add a paragraph with every style of interest.  When
// opened with OpenTemplate the document's styles will be available but the
// content will be gone.
func OpenTemplate(filename string) (*Document, error) {
	_bde, _abcb := Open(filename)
	if _abcb != nil {
		return nil, _abcb
	}
	_bde._bbgg.Body = _ag.NewCT_Body()
	return _bde, nil
}
func (_fgcb *Document) onNewRelationship(_acf *_bab.DecodeMap, _fda, _dgfg string, _gda []*_fb.File, _gef *_ede.Relationship, _beae _bab.Target) error {
	_fcgc := _f.DocTypeDocument
	switch _dgfg {
	case _f.OfficeDocumentType, _f.OfficeDocumentTypeStrict:
		_fgcb._bbgg = _ag.NewDocument()
		_acf.AddTarget(_fda, _fgcb._bbgg, _dgfg, 0)
		_acf.AddTarget(_bab.RelationsPathFor(_fda), _fgcb._agcc.X(), _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.CorePropertiesType:
		_acf.AddTarget(_fda, _fgcb.CoreProperties.X(), _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.CustomPropertiesType:
		_acf.AddTarget(_fda, _fgcb.CustomProperties.X(), _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.ExtendedPropertiesType, _f.ExtendedPropertiesTypeStrict:
		_acf.AddTarget(_fda, _fgcb.AppProperties.X(), _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.ThumbnailType, _f.ThumbnailTypeStrict:
		for _adbf, _dec := range _gda {
			if _dec == nil {
				continue
			}
			if _dec.Name == _fda {
				_ceab, _dbgg := _dec.Open()
				if _dbgg != nil {
					return _af.Errorf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073", _dbgg)
				}
				_fgcb.Thumbnail, _, _dbgg = _ba.Decode(_ceab)
				_ceab.Close()
				if _dbgg != nil {
					return _af.Errorf("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073", _dbgg)
				}
				_gda[_adbf] = nil
			}
		}
	case _f.SettingsType, _f.SettingsTypeStrict:
		_acf.AddTarget(_fda, _fgcb.Settings.X(), _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.NumberingType, _f.NumberingTypeStrict:
		_fgcb.Numbering = NewNumbering()
		_acf.AddTarget(_fda, _fgcb.Numbering.X(), _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.StylesType, _f.StylesTypeStrict:
		_fgcb.Styles.Clear()
		_acf.AddTarget(_fda, _fgcb.Styles.X(), _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.HeaderType, _f.HeaderTypeStrict:
		_ceg := _ag.NewHdr()
		_acf.AddTarget(_fda, _ceg, _dgfg, uint32(len(_fgcb._cg)))
		_fgcb._cg = append(_fgcb._cg, _ceg)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, len(_fgcb._cg))
		_aba := _ec.NewRelationships()
		_acf.AddTarget(_bab.RelationsPathFor(_fda), _aba.X(), _dgfg, 0)
		_fgcb._gba = append(_fgcb._gba, _aba)
	case _f.FooterType, _f.FooterTypeStrict:
		_fcf := _ag.NewFtr()
		_acf.AddTarget(_fda, _fcf, _dgfg, uint32(len(_fgcb._bbc)))
		_fgcb._bbc = append(_fgcb._bbc, _fcf)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, len(_fgcb._bbc))
		_cadb := _ec.NewRelationships()
		_acf.AddTarget(_bab.RelationsPathFor(_fda), _cadb.X(), _dgfg, 0)
		_fgcb._dbc = append(_fgcb._dbc, _cadb)
	case _f.ThemeType, _f.ThemeTypeStrict:
		_ggce := _ea.NewTheme()
		_acf.AddTarget(_fda, _ggce, _dgfg, uint32(len(_fgcb._adae)))
		_fgcb._adae = append(_fgcb._adae, _ggce)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, len(_fgcb._adae))
	case _f.WebSettingsType, _f.WebSettingsTypeStrict:
		_fgcb._cef = _ag.NewWebSettings()
		_acf.AddTarget(_fda, _fgcb._cef, _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.FontTableType, _f.FontTableTypeStrict:
		_fgcb._gc = _ag.NewFonts()
		_acf.AddTarget(_fda, _fgcb._gc, _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.EndNotesType, _f.EndNotesTypeStrict:
		_fgcb._fcb = _ag.NewEndnotes()
		_acf.AddTarget(_fda, _fgcb._fcb, _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.FootNotesType, _f.FootNotesTypeStrict:
		_fgcb._ab = _ag.NewFootnotes()
		_acf.AddTarget(_fda, _fgcb._ab, _dgfg, 0)
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, 0)
	case _f.ImageType, _f.ImageTypeStrict:
		var _cgga _ec.ImageRef
		for _gefg, _egdc := range _gda {
			if _egdc == nil {
				continue
			}
			if _egdc.Name == _fda {
				_fbbg, _fdga := _bab.ExtractToDiskTmp(_egdc, _fgcb.TmpPath)
				if _fdga != nil {
					return _fdga
				}
				_gga, _fdga := _ec.ImageFromStorage(_fbbg)
				if _fdga != nil {
					return _fdga
				}
				_cgga = _ec.MakeImageRef(_gga, &_fgcb.DocBase, _fgcb._agcc)
				_gda[_gefg] = nil
			}
		}
		_ecbc := "\u002e" + _c.ToLower(_cgga.Format())
		_gef.TargetAttr = _f.RelativeFilename(_fcgc, _beae.Typ, _dgfg, len(_fgcb.Images)+1)
		if _bfaa := _b.Ext(_gef.TargetAttr); _bfaa != _ecbc {
			_gef.TargetAttr = _gef.TargetAttr[0:len(_gef.TargetAttr)-len(_bfaa)] + _ecbc
		}
		_cgga.SetTarget("\u0077\u006f\u0072d\u002f" + _gef.TargetAttr)
		_fgcb.Images = append(_fgcb.Images, _cgga)
	default:
		_f.Log("\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u0074\u0067\u0074\u003a\u0020\u0025\u0073", _dgfg, _fda)
	}
	return nil
}
func (_caag Paragraph) addFldCharsForField(_abga, _ffcf string) FormField {
	_eccdeb := _caag.addBeginFldChar(_abga)
	_eceb := FormField{_bffef: _eccdeb}
	_ddff := _caag._aefd.Bookmarks()
	_gfbb := int64(len(_ddff))
	if _abga != "" {
		_caag.addStartBookmark(_gfbb, _abga)
	}
	_caag.addInstrText(_ffcf)
	_caag.addSeparateFldChar()
	if _ffcf == "\u0046\u004f\u0052\u004d\u0054\u0045\u0058\u0054" {
		_fbfg := _caag.AddRun()
		_cdea := _ag.NewEG_RunInnerContent()
		_fbfg._fegg.EG_RunInnerContent = []*_ag.EG_RunInnerContent{_cdea}
		_eceb._eeef = _cdea
	}
	_caag.addEndFldChar()
	if _abga != "" {
		_caag.addEndBookmark(_gfbb)
	}
	return _eceb
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_afeb TableStyleProperties) SetCellSpacingAuto() {
	_afeb._cgcba.TblCellSpacing = _ag.NewCT_TblWidth()
	_afeb._cgcba.TblCellSpacing.TypeAttr = _ag.ST_TblWidthAuto
}

// Caps returns true if run font is capitalized.
func (_egfe RunProperties) Caps() bool { return _decg(_egfe._ddda.Caps) }

// SetUISortOrder controls the order the style is displayed in the UI.
func (_adcbc Style) SetUISortOrder(order int) {
	_adcbc._gade.UiPriority = _ag.NewCT_DecimalNumber()
	_adcbc._gade.UiPriority.ValAttr = int64(order)
}

// SetSize sets the size of the displayed image on the page.
func (_cbad InlineDrawing) SetSize(w, h _bf.Distance) {
	_cbad._egff.Extent.CxAttr = int64(float64(w*_bf.Pixel72) / _bf.EMU)
	_cbad._egff.Extent.CyAttr = int64(float64(h*_bf.Pixel72) / _bf.EMU)
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (_bede Paragraph) Style() string {
	if _bede._dbfe.PPr != nil && _bede._dbfe.PPr.PStyle != nil {
		return _bede._dbfe.PPr.PStyle.ValAttr
	}
	return ""
}

// Paragraphs returns the paragraphs defined in a footer.
func (_eaea Footer) Paragraphs() []Paragraph {
	_gdf := []Paragraph{}
	for _, _gfaf := range _eaea._edec.EG_ContentBlockContent {
		for _, _dcff := range _gfaf.P {
			_gdf = append(_gdf, Paragraph{_eaea._dagd, _dcff})
		}
	}
	for _, _gfbc := range _eaea.Tables() {
		for _, _bbd := range _gfbc.Rows() {
			for _, _caf := range _bbd.Cells() {
				_gdf = append(_gdf, _caf.Paragraphs()...)
			}
		}
	}
	return _gdf
}

// Margins allows controlling individual cell margins.
func (_geg CellProperties) Margins() CellMargins {
	if _geg._dac.TcMar == nil {
		_geg._dac.TcMar = _ag.NewCT_TcMar()
	}
	return CellMargins{_geg._dac.TcMar}
}

// SetDefaultValue sets the default value of a FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (_bgca FormField) SetDefaultValue(v string) {
	if _bgca._bffef.DdList != nil {
		for _afab, _bbdf := range _bgca.PossibleValues() {
			if _bbdf == v {
				_bgca._bffef.DdList.Default = _ag.NewCT_DecimalNumber()
				_bgca._bffef.DdList.Default.ValAttr = int64(_afab)
				break
			}
		}
	}
}

// SetRight sets the cell right margin
func (_ce CellMargins) SetRight(d _bf.Distance) {
	_ce._cacb.Right = _ag.NewCT_TblWidth()
	_edf(_ce._cacb.Right, d)
}

// IsBold returns true if the run has been set to bold.
func (_ggbe RunProperties) IsBold() bool { return _ggbe.BoldValue() == OnOffValueOn }

// SetWidth sets the cell width to a specified width.
func (_dbf CellProperties) SetWidth(d _bf.Distance) {
	_dbf._dac.TcW = _ag.NewCT_TblWidth()
	_dbf._dac.TcW.TypeAttr = _ag.ST_TblWidthDxa
	_dbf._dac.TcW.WAttr = &_ag.ST_MeasurementOrPercent{}
	_dbf._dac.TcW.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_dbf._dac.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(d / _bf.Twips))
}

// RunProperties returns the run properties controlling text formatting within the table.
func (_agfe TableConditionalFormatting) RunProperties() RunProperties {
	if _agfe._bdceb.RPr == nil {
		_agfe._bdceb.RPr = _ag.NewCT_RPr()
	}
	return RunProperties{_agfe._bdceb.RPr}
}

// InitializeDefault constructs the default styles.
func (_fdgca Styles) InitializeDefault() {
	_fdgca.initializeDocDefaults()
	_fdgca.initializeStyleDefaults()
}

var _gbbb = [...]uint8{0, 20, 37, 58, 79}

// InsertRowBefore inserts a row before another row
func (_bcgbe Table) InsertRowBefore(r Row) Row {
	for _gece, _dcgg := range _bcgbe._fdeda.EG_ContentRowContent {
		if len(_dcgg.Tr) > 0 && r.X() == _dcgg.Tr[0] {
			_dgffa := _ag.NewEG_ContentRowContent()
			_bcgbe._fdeda.EG_ContentRowContent = append(_bcgbe._fdeda.EG_ContentRowContent, nil)
			copy(_bcgbe._fdeda.EG_ContentRowContent[_gece+1:], _bcgbe._fdeda.EG_ContentRowContent[_gece:])
			_bcgbe._fdeda.EG_ContentRowContent[_gece] = _dgffa
			_bbgdd := _ag.NewCT_Row()
			_dgffa.Tr = append(_dgffa.Tr, _bbgdd)
			return Row{_bcgbe._aebg, _bbgdd}
		}
	}
	return _bcgbe.AddRow()
}

// CellBorders are the borders for an individual
type CellBorders struct{ _eaa *_ag.CT_TcBorders }

// RemoveParagraph removes a paragraph from a footer.
func (_ffffc Footer) RemoveParagraph(p Paragraph) {
	for _, _cddg := range _ffffc._edec.EG_ContentBlockContent {
		for _dfdfe, _dagb := range _cddg.P {
			if _dagb == p._dbfe {
				copy(_cddg.P[_dfdfe:], _cddg.P[_dfdfe+1:])
				_cddg.P = _cddg.P[0 : len(_cddg.P)-1]
				return
			}
		}
	}
}
func _faefd(_face *_ea.CT_Blip, _dgge map[string]string) {
	if _face.EmbedAttr != nil {
		if _dggb, _bec := _dgge[*_face.EmbedAttr]; _bec {
			*_face.EmbedAttr = _dggb
		}
	}
}

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_deea RunProperties) SizeMeasure() string {
	if _gcgf := _deea._ddda.Sz; _gcgf != nil {
		_ggced := _gcgf.ValAttr
		if _ggced.ST_PositiveUniversalMeasure != nil {
			return *_ggced.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// AddParagraph adds a paragraph to the footer.
func (_dafd Footer) AddParagraph() Paragraph {
	_bacc := _ag.NewEG_ContentBlockContent()
	_dafd._edec.EG_ContentBlockContent = append(_dafd._edec.EG_ContentBlockContent, _bacc)
	_adbc := _ag.NewCT_P()
	_bacc.P = append(_bacc.P, _adbc)
	return Paragraph{_dafd._dagd, _adbc}
}

// SetBottomPct sets the cell bottom margin
func (_ddd CellMargins) SetBottomPct(pct float64) {
	_ddd._cacb.Bottom = _ag.NewCT_TblWidth()
	_adg(_ddd._cacb.Bottom, pct)
}

// SetBeforeAuto controls if spacing before a paragraph is automatically determined.
func (_fbcf ParagraphSpacing) SetBeforeAuto(b bool) {
	if b {
		_fbcf._gggda.BeforeAutospacingAttr = &_dc.ST_OnOff{}
		_fbcf._gggda.BeforeAutospacingAttr.Bool = _f.Bool(true)
	} else {
		_fbcf._gggda.BeforeAutospacingAttr = nil
	}
}

// RemoveParagraph removes a paragraph from a footer.
func (_bdeb Header) RemoveParagraph(p Paragraph) {
	for _, _fead := range _bdeb._aeaf.EG_ContentBlockContent {
		for _fgfg, _cafg := range _fead.P {
			if _cafg == p._dbfe {
				copy(_fead.P[_fgfg:], _fead.P[_fgfg+1:])
				_fead.P = _fead.P[0 : len(_fead.P)-1]
				return
			}
		}
	}
}

// AddHyperlink adds a hyperlink to a document. Adding the hyperlink to a document
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (_bccg Document) AddHyperlink(url string) _ec.Hyperlink { return _bccg._agcc.AddHyperlink(url) }

// X returns the inner wrapped XML type.
func (_eaec NumberingLevel) X() *_ag.CT_Lvl { return _eaec._bdce }

// SizeValue returns the value of run font size in points.
func (_edae RunProperties) SizeValue() float64 {
	if _gcdb := _edae._ddda.Sz; _gcdb != nil {
		_adbbb := _gcdb.ValAttr
		if _adbbb.ST_UnsignedDecimalNumber != nil {
			return float64(*_adbbb.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_ebbea TableBorders) SetLeft(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_ebbea._egbeb.Left = _ag.NewCT_Border()
	_ddecc(_ebbea._egbeb.Left, t, c, thickness)
}
func _gdgc(_fcaf *_ag.CT_Tbl, _fcfc map[string]string) {
	for _, _bfcc := range _fcaf.EG_ContentRowContent {
		for _, _eabb := range _bfcc.Tr {
			for _, _ddcag := range _eabb.EG_ContentCellContent {
				for _, _dfcf := range _ddcag.Tc {
					for _, _gfc := range _dfcf.EG_BlockLevelElts {
						for _, _ddfgf := range _gfc.EG_ContentBlockContent {
							for _, _gbb := range _ddfgf.P {
								_faea(_gbb, _fcfc)
							}
							for _, _fbff := range _ddfgf.Tbl {
								_gdgc(_fbff, _fcfc)
							}
						}
					}
				}
			}
		}
	}
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_bfeg ParagraphProperties) ComplexSizeMeasure() string {
	if _efcg := _bfeg._cfdd.RPr.SzCs; _efcg != nil {
		_eggc := _efcg.ValAttr
		if _eggc.ST_PositiveUniversalMeasure != nil {
			return *_eggc.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// Tables returns the tables defined in the footer.
func (_egfb Footer) Tables() []Table {
	_cbfca := []Table{}
	if _egfb._edec == nil {
		return nil
	}
	for _, _eaab := range _egfb._edec.EG_ContentBlockContent {
		for _, _gbgg := range _egfb._dagd.tables(_eaab) {
			_cbfca = append(_cbfca, _gbgg)
		}
	}
	return _cbfca
}

// SetTop sets the cell top margin
func (_fafc CellMargins) SetTop(d _bf.Distance) {
	_fafc._cacb.Top = _ag.NewCT_TblWidth()
	_edf(_fafc._cacb.Top, d)
}

// SetColor sets the text color.
func (_ebaa RunProperties) SetColor(c _ed.Color) {
	_ebaa._ddda.Color = _ag.NewCT_Color()
	_ebaa._ddda.Color.ValAttr.ST_HexColorRGB = c.AsRGBString()
}

// Index returns the index of the header within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (_ddedg Header) Index() int {
	for _abf, _aeef := range _ddedg._cedg._cg {
		if _aeef == _ddedg._aeaf {
			return _abf
		}
	}
	return -1
}

// SetLayout controls the table layout. wml.ST_TblLayoutTypeAutofit corresponds
// to "Automatically resize to fit contents" being checked, while
// wml.ST_TblLayoutTypeFixed corresponds to it being unchecked.
func (_bedaf TableProperties) SetLayout(l _ag.ST_TblLayoutType) {
	if l == _ag.ST_TblLayoutTypeUnset || l == _ag.ST_TblLayoutTypeAutofit {
		_bedaf._gfaa.TblLayout = nil
	} else {
		_bedaf._gfaa.TblLayout = _ag.NewCT_TblLayoutType()
		_bedaf._gfaa.TblLayout.TypeAttr = l
	}
}

// SetItalic sets the run to italic.
func (_gdadc RunProperties) SetItalic(b bool) {
	if !b {
		_gdadc._ddda.I = nil
		_gdadc._ddda.ICs = nil
	} else {
		_gdadc._ddda.I = _ag.NewCT_OnOff()
		_gdadc._ddda.ICs = _ag.NewCT_OnOff()
	}
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_cfdf RunProperties) ComplexSizeMeasure() string {
	if _fefcb := _cfdf._ddda.SzCs; _fefcb != nil {
		_ggbae := _fefcb.ValAttr
		if _ggbae.ST_PositiveUniversalMeasure != nil {
			return *_ggbae.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// SetBold sets the run to bold.
func (_cage RunProperties) SetBold(b bool) {
	if !b {
		_cage._ddda.B = nil
		_cage._ddda.BCs = nil
	} else {
		_cage._ddda.B = _ag.NewCT_OnOff()
		_cage._ddda.BCs = _ag.NewCT_OnOff()
	}
}

// CharacterSpacingValue returns the value of characters spacing in twips (1/20 of point).
func (_fbgd ParagraphProperties) CharacterSpacingValue() int64 {
	if _ggaa := _fbgd._cfdd.RPr.Spacing; _ggaa != nil {
		_ecfg := _ggaa.ValAttr
		if _ecfg.Int64 != nil {
			return *_ecfg.Int64
		}
	}
	return int64(0)
}
func (_dcfc Document) mergeFields() []mergeFieldInfo {
	_fecd := []Paragraph{}
	_dfgfg := []mergeFieldInfo{}
	for _, _gega := range _dcfc.Tables() {
		for _, _caaff := range _gega.Rows() {
			for _, _bbfe := range _caaff.Cells() {
				_fecd = append(_fecd, _bbfe.Paragraphs()...)
			}
		}
	}
	_fecd = append(_fecd, _dcfc.Paragraphs()...)
	for _, _ggb := range _fecd {
		_cbda := _ggb.Runs()
		_fccb := -1
		_ceb := -1
		_bcag := -1
		_aegc := mergeFieldInfo{}
		for _, _fdgf := range _ggb._dbfe.EG_PContent {
			for _, _gefa := range _fdgf.FldSimple {
				if _c.Contains(_gefa.InstrAttr, "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044") {
					_ccd := _aabb(_gefa.InstrAttr)
					_ccd._baba = true
					_ccd._cgdg = _ggb
					_ccd._fcdf = _fdgf
					_dfgfg = append(_dfgfg, _ccd)
				}
			}
		}
		for _cddf := 0; _cddf < len(_cbda); _cddf++ {
			_efgf := _cbda[_cddf]
			for _, _bdeg := range _efgf.X().EG_RunInnerContent {
				if _bdeg.FldChar != nil {
					switch _bdeg.FldChar.FldCharTypeAttr {
					case _ag.ST_FldCharTypeBegin:
						_fccb = _cddf
					case _ag.ST_FldCharTypeSeparate:
						_ceb = _cddf
					case _ag.ST_FldCharTypeEnd:
						_bcag = _cddf
						if _aegc._cabg != "" {
							_aegc._cgdg = _ggb
							_aegc._gaab = _fccb
							_aegc._efged = _bcag
							_aegc._cgfaf = _ceb
							_dfgfg = append(_dfgfg, _aegc)
						}
						_fccb = -1
						_ceb = -1
						_bcag = -1
						_aegc = mergeFieldInfo{}
					}
				} else if _bdeg.InstrText != nil && _c.Contains(_bdeg.InstrText.Content, "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044") {
					if _fccb != -1 && _bcag == -1 {
						_aegc = _aabb(_bdeg.InstrText.Content)
					}
				}
			}
		}
	}
	return _dfgfg
}

// Section is the beginning of a new section.
type Section struct {
	_cce  *Document
	_cdec *_ag.CT_SectPr
}

// Endnote is an individual endnote reference within the document.
type Endnote struct {
	_ddcec *Document
	_gae   *_ag.CT_FtnEdn
}

func _eea(_aebc *_ag.CT_OnOff) OnOffValue {
	if _aebc == nil {
		return OnOffValueUnset
	}
	if _aebc.ValAttr != nil && _aebc.ValAttr.Bool != nil && *_aebc.ValAttr.Bool == false {
		return OnOffValueOff
	}
	return OnOffValueOn
}

// PossibleValues returns the possible values for a FormFieldTypeDropDown.
func (_eeeb FormField) PossibleValues() []string {
	if _eeeb._bffef.DdList == nil {
		return nil
	}
	_gccf := []string{}
	for _, _eadg := range _eeeb._bffef.DdList.ListEntry {
		if _eadg == nil {
			continue
		}
		_gccf = append(_gccf, _eadg.ValAttr)
	}
	return _gccf
}

// DrawingAnchored returns a slice of AnchoredDrawings.
func (_fgfge Run) DrawingAnchored() []AnchoredDrawing {
	_fdffg := []AnchoredDrawing{}
	for _, _ffcd := range _fgfge._fegg.EG_RunInnerContent {
		if _ffcd.Drawing == nil {
			continue
		}
		for _, _edede := range _ffcd.Drawing.Anchor {
			_fdffg = append(_fdffg, AnchoredDrawing{_fgfge._ebdd, _edede})
		}
	}
	return _fdffg
}

// SetCSTheme sets the font complex script theme.
func (_gfga Fonts) SetCSTheme(t _ag.ST_Theme) { _gfga._agaa.CsthemeAttr = t }

// SetTarget sets the URL target of the hyperlink.
func (_daec HyperLink) SetTarget(url string) {
	_dcee := _daec._debd.AddHyperlink(url)
	_daec._dceb.IdAttr = _f.String(_ec.Relationship(_dcee).ID())
	_daec._dceb.AnchorAttr = nil
}

// SetVAlignment sets the vertical alignment for an anchored image.
func (_db AnchoredDrawing) SetVAlignment(v _ag.WdST_AlignV) {
	_db._de.PositionV.Choice = &_ag.WdCT_PosVChoice{}
	_db._de.PositionV.Choice.Align = v
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (_gadgd TableBorders) SetBottom(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_gadgd._egbeb.Bottom = _ag.NewCT_Border()
	_ddecc(_gadgd._egbeb.Bottom, t, c, thickness)
}

// SetBasedOn sets the style that this style is based on.
func (_begb Style) SetBasedOn(name string) {
	if name == "" {
		_begb._gade.BasedOn = nil
	} else {
		_begb._gade.BasedOn = _ag.NewCT_String()
		_begb._gade.BasedOn.ValAttr = name
	}
}
func (_eagc Styles) initializeDocDefaults() {
	_eagc._faabg.DocDefaults = _ag.NewCT_DocDefaults()
	_eagc._faabg.DocDefaults.RPrDefault = _ag.NewCT_RPrDefault()
	_eagc._faabg.DocDefaults.RPrDefault.RPr = _ag.NewCT_RPr()
	_gcbb := RunProperties{_eagc._faabg.DocDefaults.RPrDefault.RPr}
	_gcbb.SetSize(12 * _bf.Point)
	_gcbb.Fonts().SetASCIITheme(_ag.ST_ThemeMajorAscii)
	_gcbb.Fonts().SetEastAsiaTheme(_ag.ST_ThemeMajorEastAsia)
	_gcbb.Fonts().SetHANSITheme(_ag.ST_ThemeMajorHAnsi)
	_gcbb.Fonts().SetCSTheme(_ag.ST_ThemeMajorBidi)
	_gcbb.X().Lang = _ag.NewCT_Language()
	_gcbb.X().Lang.ValAttr = _f.String("\u0065\u006e\u002dU\u0053")
	_gcbb.X().Lang.EastAsiaAttr = _f.String("\u0065\u006e\u002dU\u0053")
	_gcbb.X().Lang.BidiAttr = _f.String("\u0061\u0072\u002dS\u0041")
	_eagc._faabg.DocDefaults.PPrDefault = _ag.NewCT_PPrDefault()
}

// Spacing returns the paragraph spacing settings.
func (_fcfe ParagraphProperties) Spacing() ParagraphSpacing {
	if _fcfe._cfdd.Spacing == nil {
		_fcfe._cfdd.Spacing = _ag.NewCT_Spacing()
	}
	return ParagraphSpacing{_fcfe._cfdd.Spacing}
}
func (_feg *Document) InsertTableBefore(relativeTo Paragraph) Table {
	return _feg.insertTable(relativeTo, true)
}
func _edf(_edc *_ag.CT_TblWidth, _bfd _bf.Distance) {
	_edc.TypeAttr = _ag.ST_TblWidthDxa
	_edc.WAttr = &_ag.ST_MeasurementOrPercent{}
	_edc.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_edc.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(_bfd / _bf.Dxa))
}

// AddParagraph adds a paragraph to the header.
func (_ccba Header) AddParagraph() Paragraph {
	_agae := _ag.NewEG_ContentBlockContent()
	_ccba._aeaf.EG_ContentBlockContent = append(_ccba._aeaf.EG_ContentBlockContent, _agae)
	_afge := _ag.NewCT_P()
	_agae.P = append(_agae.P, _afge)
	return Paragraph{_ccba._cedg, _afge}
}

// Properties returns the cell properties.
func (_def Cell) Properties() CellProperties {
	if _def._afg.TcPr == nil {
		_def._afg.TcPr = _ag.NewCT_TcPr()
	}
	return CellProperties{_def._afg.TcPr}
}

// Endnote returns the endnote based on the ID; this can be used nicely with
// the run.IsEndnote() functionality.
func (_bea *Document) Endnote(id int64) Endnote {
	for _, _gfb := range _bea.Endnotes() {
		if _gfb.id() == id {
			return _gfb
		}
	}
	return Endnote{}
}

// Open opens and reads a document from a file (.docx).
func Open(filename string) (*Document, error) {
	_cde, _cge := _d.Open(filename)
	if _cge != nil {
		return nil, _af.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _cge)
	}
	defer _cde.Close()
	_cfab, _cge := _d.Stat(filename)
	if _cge != nil {
		return nil, _af.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _cge)
	}
	_ = _cfab
	return Read(_cde, _cfab.Size())
}

// Paragraphs returns the paragraphs defined in a header.
func (_cgff Header) Paragraphs() []Paragraph {
	_deag := []Paragraph{}
	for _, _gffc := range _cgff._aeaf.EG_ContentBlockContent {
		for _, _dbab := range _gffc.P {
			_deag = append(_deag, Paragraph{_cgff._cedg, _dbab})
		}
	}
	for _, _gbca := range _cgff.Tables() {
		for _, _beg := range _gbca.Rows() {
			for _, _gfce := range _beg.Cells() {
				_deag = append(_deag, _gfce.Paragraphs()...)
			}
		}
	}
	return _deag
}
func (_daga Run) newIC() *_ag.EG_RunInnerContent {
	_bcgb := _ag.NewEG_RunInnerContent()
	_daga._fegg.EG_RunInnerContent = append(_daga._fegg.EG_RunInnerContent, _bcgb)
	return _bcgb
}

// SetKeepNext controls if the paragraph is kept with the next paragraph.
func (_fcff ParagraphStyleProperties) SetKeepNext(b bool) {
	if !b {
		_fcff._ddde.KeepNext = nil
	} else {
		_fcff._ddde.KeepNext = _ag.NewCT_OnOff()
	}
}

// SetThemeColor sets the color from the theme.
func (_gfd Color) SetThemeColor(t _ag.ST_ThemeColor) { _gfd._ccc.ThemeColorAttr = t }

// Levels returns all of the numbering levels defined in the definition.
func (_eadfa NumberingDefinition) Levels() []NumberingLevel {
	_bdaa := []NumberingLevel{}
	for _, _fgee := range _eadfa._cgfd.Lvl {
		_bdaa = append(_bdaa, NumberingLevel{_fgee})
	}
	return _bdaa
}

// SetOffset sets the offset of the image relative to the origin, which by
// default this is the top-left corner of the page. Offset is incompatible with
// SetAlignment, whichever is called last is applied.
func (_adf AnchoredDrawing) SetOffset(x, y _bf.Distance) { _adf.SetXOffset(x); _adf.SetYOffset(y) }

// FormField is a form within a document. It references the document, so changes
// to the form field wil be reflected in the document if it is saved.
type FormField struct {
	_bffef *_ag.CT_FFData
	_eeef  *_ag.EG_RunInnerContent
}

// AddLevel adds a new numbering level to a NumberingDefinition.
func (_ffag NumberingDefinition) AddLevel() NumberingLevel {
	_bbef := _ag.NewCT_Lvl()
	_bbef.Start = &_ag.CT_DecimalNumber{ValAttr: 1}
	_bbef.IlvlAttr = int64(len(_ffag._cgfd.Lvl))
	_ffag._cgfd.Lvl = append(_ffag._cgfd.Lvl, _bbef)
	return NumberingLevel{_bbef}
}

// SetName marks sets a name attribute for a FormField.
func (_dbdf FormField) SetName(name string) {
	_cdbf := _ag.NewCT_FFName()
	_cdbf.ValAttr = &name
	_dbdf._bffef.Name = []*_ag.CT_FFName{_cdbf}
}

// Footer is a footer for a document section.
type Footer struct {
	_dagd *Document
	_edec *_ag.Ftr
}

// SetSpacing sets the spacing that comes before and after the paragraph.
func (_cbbf ParagraphStyleProperties) SetSpacing(before, after _bf.Distance) {
	if _cbbf._ddde.Spacing == nil {
		_cbbf._ddde.Spacing = _ag.NewCT_Spacing()
	}
	if before == _bf.Zero {
		_cbbf._ddde.Spacing.BeforeAttr = nil
	} else {
		_cbbf._ddde.Spacing.BeforeAttr = &_dc.ST_TwipsMeasure{}
		_cbbf._ddde.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(before / _bf.Twips))
	}
	if after == _bf.Zero {
		_cbbf._ddde.Spacing.AfterAttr = nil
	} else {
		_cbbf._ddde.Spacing.AfterAttr = &_dc.ST_TwipsMeasure{}
		_cbbf._ddde.Spacing.AfterAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(after / _bf.Twips))
	}
}
func (_caee FormFieldType) String() string {
	if _caee >= FormFieldType(len(_gbbb)-1) {
		return _af.Sprintf("\u0046\u006f\u0072\u006d\u0046\u0069\u0065\u006c\u0064\u0054\u0079\u0070e\u0028\u0025\u0064\u0029", _caee)
	}
	return _ded[_gbbb[_caee]:_gbbb[_caee+1]]
}

// InsertRunAfter inserts a run in the paragraph after the relative run.
func (_fadc Paragraph) InsertRunAfter(relativeTo Run) Run { return _fadc.insertRun(relativeTo, false) }

// X returns the inner wrapped XML type.
func (_ddf AnchoredDrawing) X() *_ag.WdAnchor { return _ddf._de }

// Numbering is the document wide numbering styles contained in numbering.xml.
type Numbering struct{ _eefa *_ag.Numbering }

// SetName sets the name of the style.
func (_bgef Style) SetName(name string) {
	_bgef._gade.Name = _ag.NewCT_String()
	_bgef._gade.Name.ValAttr = name
}

// Borders allows manipulation of the table borders.
func (_cegb TableProperties) Borders() TableBorders {
	if _cegb._gfaa.TblBorders == nil {
		_cegb._gfaa.TblBorders = _ag.NewCT_TblBorders()
	}
	return TableBorders{_cegb._gfaa.TblBorders}
}

// Fonts returns the style's Fonts.
func (_fga RunProperties) Fonts() Fonts {
	if _fga._ddda.RFonts == nil {
		_fga._ddda.RFonts = _ag.NewCT_Fonts()
	}
	return Fonts{_fga._ddda.RFonts}
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (_edd CellBorders) SetInsideVertical(t _ag.ST_Border, c _ed.Color, thickness _bf.Distance) {
	_edd._eaa.InsideV = _ag.NewCT_Border()
	_ddecc(_edd._eaa.InsideV, t, c, thickness)
}

// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	_bgfc *Document
	_cfdd *_ag.CT_PPr
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (_ffdc RunProperties) RStyle() string {
	if _ffdc._ddda.RStyle != nil {
		return _ffdc._ddda.RStyle.ValAttr
	}
	return ""
}

// CharacterSpacingValue returns the value of run's characters spacing in twips (1/20 of point).
func (_ffgbe RunProperties) CharacterSpacingValue() int64 {
	if _afca := _ffgbe._ddda.Spacing; _afca != nil {
		_bfgf := _afca.ValAttr
		if _bfgf.Int64 != nil {
			return *_bfgf.Int64
		}
	}
	return int64(0)
}

// Name returns the name of the style if set.
func (_deda Style) Name() string {
	if _deda._gade.Name == nil {
		return ""
	}
	return _deda._gade.Name.ValAttr
}

// SetKerning sets the run's font kerning.
func (_aaaf RunProperties) SetKerning(size _bf.Distance) {
	_aaaf._ddda.Kern = _ag.NewCT_HpsMeasure()
	_aaaf._ddda.Kern.ValAttr.ST_UnsignedDecimalNumber = _f.Uint64(uint64(size / _bf.HalfPoint))
}

// SetStrict is a shortcut for document.SetConformance,
// as one of these values from gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (_bfff Document) SetStrict(strict bool) {
	if strict {
		_bfff._bbgg.ConformanceAttr = _dc.ST_ConformanceClassStrict
	} else {
		_bfff._bbgg.ConformanceAttr = _dc.ST_ConformanceClassTransitional
	}
}
func (_gadc Paragraph) addInstrText(_dabdd string) *_ag.CT_Text {
	_decd := _gadc.AddRun()
	_dgfgc := _decd.X()
	_gdcbb := _ag.NewEG_RunInnerContent()
	_gggg := _ag.NewCT_Text()
	_deg := "\u0070\u0072\u0065\u0073\u0065\u0072\u0076\u0065"
	_gggg.SpaceAttr = &_deg
	_gggg.Content = "\u0020" + _dabdd + "\u0020"
	_gdcbb.InstrText = _gggg
	_dgfgc.EG_RunInnerContent = append(_dgfgc.EG_RunInnerContent, _gdcbb)
	return _gggg
}

// SetWidth sets the table with to a specified width.
func (_abbb TableProperties) SetWidth(d _bf.Distance) {
	_abbb._gfaa.TblW = _ag.NewCT_TblWidth()
	_abbb._gfaa.TblW.TypeAttr = _ag.ST_TblWidthDxa
	_abbb._gfaa.TblW.WAttr = &_ag.ST_MeasurementOrPercent{}
	_abbb._gfaa.TblW.WAttr.ST_DecimalNumberOrPercent = &_ag.ST_DecimalNumberOrPercent{}
	_abbb._gfaa.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _f.Int64(int64(d / _bf.Twips))
}

// IsChecked returns true if a FormFieldTypeCheckBox is checked.
func (_egcg FormField) IsChecked() bool {
	if _egcg._bffef.CheckBox == nil {
		return false
	}
	if _egcg._bffef.CheckBox.Checked != nil {
		return true
	}
	return false
}

// SetCalcOnExit marks if a FormField should be CalcOnExit or not.
func (_eggee FormField) SetCalcOnExit(calcOnExit bool) {
	_acab := _ag.NewCT_OnOff()
	_acab.ValAttr = &_dc.ST_OnOff{Bool: &calcOnExit}
	_eggee._bffef.CalcOnExit = []*_ag.CT_OnOff{_acab}
}

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_dbgbf ParagraphProperties) SizeMeasure() string {
	if _abdc := _dbgbf._cfdd.RPr.Sz; _abdc != nil {
		_cdac := _abdc.ValAttr
		if _cdac.ST_PositiveUniversalMeasure != nil {
			return *_cdac.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

const _ded = "\u0046\u006f\u0072\u006d\u0046\u0069\u0065l\u0064\u0054\u0079\u0070\u0065\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0046\u006fr\u006dF\u0069\u0065\u006c\u0064\u0054\u0079p\u0065\u0054\u0065\u0078\u0074\u0046\u006fr\u006d\u0046\u0069\u0065\u006c\u0064\u0054\u0079\u0070\u0065\u0043\u0068\u0065\u0063\u006b\u0042\u006f\u0078\u0046\u006f\u0072\u006d\u0046i\u0065\u006c\u0064\u0054\u0079\u0070\u0065\u0044\u0072\u006f\u0070\u0044\u006fw\u006e"

// AddDefinition adds a new numbering definition.
func (_ddeac Numbering) AddDefinition() NumberingDefinition {
	_ebcb := _ag.NewCT_Num()
	_facae := int64(1)
	for _, _edaa := range _ddeac.Definitions() {
		if _edaa.AbstractNumberID() >= _facae {
			_facae = _edaa.AbstractNumberID() + 1
		}
	}
	_aaef := int64(1)
	for _, _eadc := range _ddeac.X().Num {
		if _eadc.NumIdAttr >= _aaef {
			_aaef = _eadc.NumIdAttr + 1
		}
	}
	_ebcb.NumIdAttr = _aaef
	_ebcb.AbstractNumId = _ag.NewCT_DecimalNumber()
	_ebcb.AbstractNumId.ValAttr = _facae
	_dafgg := _ag.NewCT_AbstractNum()
	_dafgg.AbstractNumIdAttr = _facae
	_ddeac._eefa.AbstractNum = append(_ddeac._eefa.AbstractNum, _dafgg)
	_ddeac._eefa.Num = append(_ddeac._eefa.Num, _ebcb)
	return NumberingDefinition{_dafgg}
}

// SetAlignment positions an anchored image via alignment.  Offset is
// incompatible with SetOffset, whichever is called last is applied.
func (_ade AnchoredDrawing) SetAlignment(h _ag.WdST_AlignH, v _ag.WdST_AlignV) {
	_ade.SetHAlignment(h)
	_ade.SetVAlignment(v)
}

// AddParagraph adds a paragraph to the footnote.
func (_gfcf Footnote) AddParagraph() Paragraph {
	_bbda := _ag.NewEG_ContentBlockContent()
	_gcee := len(_gfcf._faaeg.EG_BlockLevelElts[0].EG_ContentBlockContent)
	_gfcf._faaeg.EG_BlockLevelElts[0].EG_ContentBlockContent = append(_gfcf._faaeg.EG_BlockLevelElts[0].EG_ContentBlockContent, _bbda)
	_dcge := _ag.NewCT_P()
	var _dbcbd *_ag.CT_String
	if _gcee != 0 {
		_bcec := len(_gfcf._faaeg.EG_BlockLevelElts[0].EG_ContentBlockContent[_gcee-1].P)
		_dbcbd = _gfcf._faaeg.EG_BlockLevelElts[0].EG_ContentBlockContent[_gcee-1].P[_bcec-1].PPr.PStyle
	} else {
		_dbcbd = _ag.NewCT_String()
		_dbcbd.ValAttr = "\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065"
	}
	_bbda.P = append(_bbda.P, _dcge)
	_ggae := Paragraph{_gfcf._gagb, _dcge}
	_ggae._dbfe.PPr = _ag.NewCT_PPr()
	_ggae._dbfe.PPr.PStyle = _dbcbd
	_ggae._dbfe.PPr.RPr = _ag.NewCT_ParaRPr()
	return _ggae
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (_fadce ParagraphStyleProperties) SetKeepOnOnePage(b bool) {
	if !b {
		_fadce._ddde.KeepLines = nil
	} else {
		_fadce._ddde.KeepLines = _ag.NewCT_OnOff()
	}
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (_agbd ParagraphSpacing) SetLineSpacing(d _bf.Distance, rule _ag.ST_LineSpacingRule) {
	if rule == _ag.ST_LineSpacingRuleUnset {
		_agbd._gggda.LineRuleAttr = _ag.ST_LineSpacingRuleUnset
		_agbd._gggda.LineAttr = nil
	} else {
		_agbd._gggda.LineRuleAttr = rule
		_agbd._gggda.LineAttr = &_ag.ST_SignedTwipsMeasure{}
		_agbd._gggda.LineAttr.Int64 = _f.Int64(int64(d / _bf.Twips))
	}
}
