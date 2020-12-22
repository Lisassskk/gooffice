package presentation

import (
	_bd "archive/zip"
	_ba "bytes"
	_ce "encoding/xml"
	_e "errors"
	_dg "fmt"
	_fcd "image"
	_fe "image/jpeg"
	_ae "io"
	_bg "log"
	_cd "os"
	_c "path"
	_a "strconv"
	_fc "strings"

	_db "gitee.com/gooffice/gooffice"
	_ff "gitee.com/gooffice/gooffice/common"
	_b "gitee.com/gooffice/gooffice/common/license"
	_bb "gitee.com/gooffice/gooffice/common/tempstorage"
	_g "gitee.com/gooffice/gooffice/drawing"
	_d "gitee.com/gooffice/gooffice/measurement"
	_dbb "gitee.com/gooffice/gooffice/schema/soo/dml"
	_fce "gitee.com/gooffice/gooffice/schema/soo/dml/chart"
	_fca "gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes"
	_ac "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships"
	_gb "gitee.com/gooffice/gooffice/schema/soo/pml"
	_fg "gitee.com/gooffice/gooffice/zippkg"
)

// SorterViewPr returns the SorterViewPr property.
func (_cgf ViewProperties) SorterViewPr() *_gb.CT_SlideSorterViewProperties {
	return _cgf._bbab.SorterViewPr
}

// GetTextBoxes returns a list of all text boxes from a slide.
func (_egeg Slide) GetTextBoxes() []*TextBox {
	_bdae := []*TextBox{}
	_efcc := _egeg._caed.CSld.SpTree.Choice
	for _, _aeaa := range _efcc {
		for _, _ebeg := range _aeaa.Sp {
			if _ebeg.NvSpPr.CNvSpPr.TxBoxAttr != nil && *_ebeg.NvSpPr.CNvSpPr.TxBoxAttr {
				_bdae = append(_bdae, &TextBox{_ebeg})
			}
		}
	}
	return _bdae
}

// Save writes the presentation out to a writer in the Zip package format
func (_adfg *Presentation) Save(w _ae.Writer) error { return _adfg.save(w, false) }

// SlideMasters returns the slide masters defined in the presentation.
func (_efc *Presentation) SlideMasters() []SlideMaster {
	_edb := []SlideMaster{}
	for _gdc, _gde := range _efc._ea {
		_edb = append(_edb, SlideMaster{_efc, _efc._accc[_gdc], _gde})
	}
	return _edb
}

// X returns the inner wrapped XML type.
func (_dce *Presentation) X() *_gb.Presentation { return _dce._cc }

// SlideLayouts returns a slice of all layouts in SlideMaster.
func (_beea SlideMaster) SlideLayouts() []SlideLayout {
	_fada := map[string]int{}
	_egfg := []SlideLayout{}
	for _, _fbf := range _beea._dbg.Relationships() {
		_bba := _fc.Replace(_fbf.Target(), ".\u002e\u002f\u0073\u006c\u0069\u0064e\u004c\u0061\u0079\u006f\u0075\u0074\u0073\u002f\u0073l\u0069\u0064\u0065L\u0061y\u006f\u0075\u0074", "", -1)
		_bba = _fc.Replace(_bba, "\u002e\u0078\u006d\u006c", "", -1)
		if _gfbc, _aegb := _a.ParseInt(_bba, 10, 32); _aegb == nil {
			_fada[_fbf.ID()] = int(_gfbc)
		}
	}
	for _, _ddda := range _beea._bbc.SldLayoutIdLst.SldLayoutId {
		if _deb, _ggbb := _fada[_ddda.RIdAttr]; _ggbb {
			_cdgf := _beea._gcfc._fbd[_deb-1]
			_egfg = append(_egfg, SlideLayout{_cdgf})
		}
	}
	return _egfg
}

// AddImageToRels adds an image relationship to a slide without putting image on the slide.
func (_dfea Slide) AddImageToRels(img _ff.ImageRef) string {
	_gefg := 0
	for _cbfe, _aef := range _dfea._dae.Images {
		if _aef == img {
			_gefg = _cbfe + 1
			break
		}
	}
	var _aeec string
	for _acee, _eca := range _dfea._dae.Slides() {
		if _eca._caed == _dfea._caed {
			_gcf := _dg.Sprintf("\u002e\u002e\u002f\u006ded\u0069\u0061\u002f\u0069\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", _gefg, img.Format())
			_bee := _dfea._dae._fba[_acee].AddRelationship(_gcf, _db.ImageType)
			_aeec = _bee.ID()
		}
	}
	return _aeec
}

// SetOffsetX sets horizontal offset of text box in distance units (see measurement package).
func (_efg TextBox) SetOffsetX(offX float64) {
	_bbbdf := _efg.getOff()
	_abage := _d.ToEMU(offX)
	_bbbdf.XAttr = _dbb.ST_Coordinate{ST_CoordinateUnqualified: &_abage}
}

// AddParagraph adds a new paragraph to a placeholder.
func (_bga PlaceHolder) AddParagraph() _g.Paragraph {
	_fcb := _g.MakeParagraph(_dbb.NewCT_TextParagraph())
	_bga._fb.TxBody.P = append(_bga._fb.TxBody.P, _fcb.X())
	return _fcb
}

// AddTextBox adds an empty textbox to a slide.
func (_bfad Slide) AddTextBox() TextBox {
	_abe := _gb.NewCT_GroupShapeChoice()
	_bfad._caed.CSld.SpTree.Choice = append(_bfad._caed.CSld.SpTree.Choice, _abe)
	_beae := _gb.NewCT_Shape()
	_abe.Sp = append(_abe.Sp, _beae)
	_beae.SpPr = _dbb.NewCT_ShapeProperties()
	_beae.SpPr.Xfrm = _dbb.NewCT_Transform2D()
	_beae.SpPr.PrstGeom = _dbb.NewCT_PresetGeometry2D()
	_beae.SpPr.PrstGeom.PrstAttr = _dbb.ST_ShapeTypeRect
	_beae.NvSpPr = _gb.NewCT_ShapeNonVisual()
	_beae.NvSpPr.CNvSpPr = _dbb.NewCT_NonVisualDrawingShapeProps()
	_gdcg := true
	_beae.NvSpPr.CNvSpPr.TxBoxAttr = &_gdcg
	_beae.TxBody = _dbb.NewCT_TextBody()
	_beae.TxBody.BodyPr = _dbb.NewCT_TextBodyProperties()
	_beae.TxBody.BodyPr.WrapAttr = _dbb.ST_TextWrappingTypeSquare
	_beae.TxBody.BodyPr.SpAutoFit = _dbb.NewCT_TextShapeAutofit()
	_fbg := TextBox{_beae}
	_fbg.Properties().SetWidth(3 * _d.Inch)
	_fbg.Properties().SetHeight(1 * _d.Inch)
	_fbg.Properties().SetPosition(0, 0)
	return _fbg
}
func (_eaa *Presentation) saveToFile(_dbef string, _egc bool) error {
	_cgab, _cfa := _cd.Create(_dbef)
	if _cfa != nil {
		return _cfa
	}
	defer _cgab.Close()
	return _eaa.save(_cgab, _egc)
}

// Presentation is the a presentation base document.
type Presentation struct {
	_ff.DocBase
	_cc   *_gb.Presentation
	_bed  _ff.Relationships
	_ddc  []*_gb.Sld
	_fba  []_ff.Relationships
	_ea   []*_gb.SldMaster
	_accc []_ff.Relationships
	_fbd  []*_gb.SldLayout
	_beb  []_ff.Relationships
	_bcd  []*_dbb.Theme
	_gea  []_ff.Relationships
	_ad   _ff.TableStyles
	_ec   PresentationProperties
	_gbg  ViewProperties
	_cde  []*_dbb.CT_Hyperlink
	_eag  []*_fce.ChartSpace
	_ecc  []*_gb.HandoutMaster
	_gd   []*_gb.NotesMaster
	_cgb  []*_db.XSDAny
	_bbd  map[string]string
}

// ClearAll completely clears a placeholder. To be useable, at least one
// paragraph must be added after ClearAll via AddParagraph.
func (_fga PlaceHolder) ClearAll() {
	_fga._fb.SpPr = _dbb.NewCT_ShapeProperties()
	_fga._fb.TxBody = _dbb.NewCT_TextBody()
	_fga._fb.TxBody.LstStyle = _dbb.NewCT_TextListStyle()
}

// GetPlaceholder returns a placeholder given its type.  If there are multiplace
// placeholders of the same type, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (_geae Slide) GetPlaceholder(t _gb.ST_PlaceholderType) (PlaceHolder, error) {
	for _, _dagd := range _geae._caed.CSld.SpTree.Choice {
		for _, _gaec := range _dagd.Sp {
			if _gaec.NvSpPr != nil && _gaec.NvSpPr.NvPr != nil && _gaec.NvSpPr.NvPr.Ph != nil {
				if _gaec.NvSpPr.NvPr.Ph.TypeAttr == t {
					return PlaceHolder{_gaec, _geae._caed}, nil
				}
			}
		}
	}
	return PlaceHolder{}, _e.New("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066i\u006e\u0064\u0020\u0070\u006c\u0061\u0063\u0065\u0068\u006fl\u0064\u0065\u0072")
}

// ValidateWithPath validates the slide passing path informaton for a better
// error message.
func (_geef Slide) ValidateWithPath(path string) error {
	if _daf := _geef._caed.ValidateWithPath(path); _daf != nil {
		return _daf
	}
	for _, _caec := range _geef._caed.CSld.SpTree.Choice {
		for _, _bbbed := range _caec.Sp {
			if _bbbed.TxBody != nil {
				if len(_bbbed.TxBody.P) == 0 {
					return _e.New(path + "\u0020\u003a \u0073\u006c\u0069\u0064\u0065 \u0073\u0068\u0061\u0070\u0065 \u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0074\u0078\u0062\u006f\u0064\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0073")
				}
			}
		}
	}
	return nil
}

// X returns the inner wrapped XML type.
func (_baca ViewProperties) X() *_gb.ViewPr { return _baca._bbab }
func _ga() *Presentation {
	_bea := &Presentation{_cc: _gb.NewPresentation()}
	_bea._cc.SldIdLst = _gb.NewCT_SlideIdList()
	_bea._cc.ConformanceAttr = _fca.ST_ConformanceClassTransitional
	_bea.AppProperties = _ff.NewAppProperties()
	_bea.CoreProperties = _ff.NewCoreProperties()
	_bea._ad = _ff.NewTableStyles()
	_bea.ContentTypes = _ff.NewContentTypes()
	_bea.Rels = _ff.NewRelationships()
	_bea._bed = _ff.NewRelationships()
	_bea._ec = NewPresentationProperties()
	_bea._gbg = NewViewProperties()
	_bea._bbd = map[string]string{}
	return _bea
}

// AddImage adds an image textbox to a slide.
func (_feda Slide) AddImage(img _ff.ImageRef) Image {
	_dbba := _gb.NewCT_GroupShapeChoice()
	_feda._caed.CSld.SpTree.Choice = append(_feda._caed.CSld.SpTree.Choice, _dbba)
	_aafc := _gb.NewCT_Picture()
	_dbba.Pic = append(_dbba.Pic, _aafc)
	_aafc.NvPicPr.CNvPicPr = _dbb.NewCT_NonVisualPictureProperties()
	_aafc.NvPicPr.CNvPicPr.PicLocks = _dbb.NewCT_PictureLocking()
	_aafc.NvPicPr.CNvPicPr.PicLocks.NoChangeAspectAttr = _db.Bool(true)
	_aafc.BlipFill = _dbb.NewCT_BlipFillProperties()
	_aafc.BlipFill.Blip = _dbb.NewCT_Blip()
	_gebf := _feda.AddImageToRels(img)
	_aafc.BlipFill.Blip.EmbedAttr = _db.String(_gebf)
	_aafc.BlipFill.Stretch = _dbb.NewCT_StretchInfoProperties()
	_aafc.BlipFill.Stretch.FillRect = _dbb.NewCT_RelativeRect()
	_aafc.SpPr = _dbb.NewCT_ShapeProperties()
	_aafc.SpPr.PrstGeom = _dbb.NewCT_PresetGeometry2D()
	_aafc.SpPr.PrstGeom.PrstAttr = _dbb.ST_ShapeTypeRect
	_afa := Image{_aafc}
	_aagf := img.Size()
	_afa.Properties().SetWidth(_d.Distance(_aagf.X) * _d.Pixel72)
	_afa.Properties().SetHeight(_d.Distance(_aagf.Y) * _d.Pixel72)
	_afa.Properties().SetPosition(0, 0)
	return _afa
}

// Read reads a document from an io.Reader.
func Read(r _ae.ReaderAt, size int64) (*Presentation, error) {
	_ccd := _ga()
	_adg, _abag := _bb.TempDir("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0070\u0070\u0074\u0078")
	if _abag != nil {
		return nil, _abag
	}
	_ccd.TmpPath = _adg
	_bfg, _abag := _bd.NewReader(r, size)
	if _abag != nil {
		return nil, _dg.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073", _abag)
	}
	_eae := []*_bd.File{}
	_eae = append(_eae, _bfg.File...)
	_ddgb := false
	for _, _dede := range _eae {
		if _dede.FileHeader.Name == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
			_ddgb = true
			break
		}
	}
	if _ddgb {
		_ccd.CreateCustomProperties()
	}
	_bbbe := _fg.DecodeMap{}
	_bbbe.SetOnNewRelationshipFunc(_ccd.onNewRelationship)
	_bbbe.AddTarget(_db.ContentTypesFilename, _ccd.ContentTypes.X(), "", 0)
	_bbbe.AddTarget(_db.BaseRelsFilename, _ccd.Rels.X(), "", 0)
	if _gee := _bbbe.Decode(_eae); _gee != nil {
		return nil, _gee
	}
	for _, _bab := range _eae {
		if _bab == nil {
			continue
		}
		if _ceed := _ccd.AddExtraFileFromZip(_bab); _ceed != nil {
			return nil, _ceed
		}
	}
	if _ddgb {
		_aad := false
		for _, _gggc := range _ccd.Rels.X().Relationship {
			if _gggc.TargetAttr == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
				_aad = true
				break
			}
		}
		if !_aad {
			_ccd.AddCustomRelationships()
		}
	}
	return _ccd, nil
}

// OutlineViewPr returns the OutlineViewPr property.
func (_abd ViewProperties) OutlineViewPr() *_gb.CT_OutlineViewProperties {
	return _abd._bbab.OutlineViewPr
}

// SaveAsTemplate writes the presentation out to a writer in the Zip package format as a template
func (_ceb *Presentation) SaveAsTemplate(w _ae.Writer) error { return _ceb.save(w, true) }

// AddDefaultSlideWithLayout tries to replicate what PowerPoint does when
// inserting a slide with a new style by clearing placeholder content and removing
// some placeholders.  Use AddSlideWithLayout if you need more control.
func (_fged *Presentation) AddDefaultSlideWithLayout(l SlideLayout) (Slide, error) {
	_bgaa, _bcdd := _fged.AddSlideWithLayout(l)
	for _, _bbg := range _bgaa.PlaceHolders() {
		_bbg.Clear()
		switch _bbg.Type() {
		case _gb.ST_PlaceholderTypeFtr, _gb.ST_PlaceholderTypeDt, _gb.ST_PlaceholderTypeSldNum:
			_bbg.Remove()
		}
	}
	return _bgaa, _bcdd
}
func (_gbd *Presentation) Validate() error {
	if _ggcg := _gbd._cc.Validate(); _ggcg != nil {
		return _ggcg
	}
	for _bbf, _fdc := range _gbd.Slides() {
		if _ade := _fdc.ValidateWithPath(_dg.Sprintf("\u0053l\u0069\u0064\u0065\u005b\u0025\u0064]", _bbf)); _ade != nil {
			return _ade
		}
	}
	for _cgaba, _ece := range _gbd._ea {
		if _egg := _ece.ValidateWithPath(_dg.Sprintf("\u0053l\u0069d\u0065\u004d\u0061\u0073\u0074\u0065\u0072\u005b\u0025\u0064\u005d", _cgaba)); _egg != nil {
			return _egg
		}
	}
	for _gc, _bac := range _gbd._fbd {
		if _caeb := _bac.ValidateWithPath(_dg.Sprintf("\u0053l\u0069d\u0065\u004c\u0061\u0079\u006f\u0075\u0074\u005b\u0025\u0064\u005d", _gc)); _caeb != nil {
			return _caeb
		}
	}
	return nil
}

// ExtLst returns the ExtLst property.
func (_ebd ViewProperties) ExtLst() *_gb.CT_ExtensionList { return _ebd._bbab.ExtLst }

// X returns the inner wrapped XML type.
func (_egf PresentationProperties) X() *_gb.PresentationPr { return _egf._aed }

// PresentationProperties contains document specific properties.
type PresentationProperties struct{ _aed *_gb.PresentationPr }

// GridSpacing returns the GridSpacing property.
func (_gfa ViewProperties) GridSpacing() *_dbb.CT_PositiveSize2D { return _gfa._bbab.GridSpacing }

// Type returns the placeholder type
func (_gg PlaceHolder) Type() _gb.ST_PlaceholderType { return _gg._fb.NvSpPr.NvPr.Ph.TypeAttr }

// SaveToFileAsTemplate writes the Presentation out to a file as a template.
func (_fed *Presentation) SaveToFileAsTemplate(path string) error { return _fed.saveToFile(path, true) }

// ViewProperties contains presentation specific properties.
type ViewProperties struct{ _bbab *_gb.ViewPr }

func (_fae TextBox) getOff() *_dbb.CT_Point2D {
	if _fae._cfbc.SpPr == nil {
		_fae._cfbc.SpPr = _dbb.NewCT_ShapeProperties()
	}
	if _fae._cfbc.SpPr.Xfrm == nil {
		_fae._cfbc.SpPr.Xfrm = _dbb.NewCT_Transform2D()
	}
	if _fae._cfbc.SpPr.Xfrm.Off == nil {
		_fae._cfbc.SpPr.Xfrm.Off = _dbb.NewCT_Point2D()
	}
	return _fae._cfbc.SpPr.Xfrm.Off
}

// LastViewAttr returns the LastViewAttr property.
func (_bgd ViewProperties) LastViewAttr() _gb.ST_ViewType { return _bgd._bbab.LastViewAttr }

// Open opens and reads a document from a file (.pptx).
func Open(filename string) (*Presentation, error) {
	_fcdg, _da := _cd.Open(filename)
	if _da != nil {
		return nil, _dg.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _da)
	}
	defer _fcdg.Close()
	_cg, _da := _cd.Stat(filename)
	if _da != nil {
		return nil, _dg.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _da)
	}
	_ = _cg
	return Read(_fcdg, _cg.Size())
}

// ExtLst returns the ExtLst property.
func (_bdb PresentationProperties) ExtLst() *_gb.CT_ExtensionList { return _bdb._aed.ExtLst }

// X returns the inner wrapped XML type.
func (_aabg SlideLayout) X() *_gb.SldLayout { return _aabg._cca }

// PrnPr returns the PrnPr property.
func (_bf PresentationProperties) PrnPr() *_gb.CT_PrintProperties { return _bf._aed.PrnPr }

// GetLayoutByName retrieves a slide layout given a layout name.
func (_bcf *Presentation) GetLayoutByName(name string) (SlideLayout, error) {
	for _, _gged := range _bcf._fbd {
		if _gged.CSld.NameAttr != nil && name == *_gged.CSld.NameAttr {
			return SlideLayout{_gged}, nil
		}
	}
	return SlideLayout{}, _e.New("\u0075\u006eab\u006c\u0065\u0020t\u006f\u0020\u0066\u0069nd \u006cay\u006f\u0075\u0074\u0020\u0077\u0069\u0074h \u0074\u0068\u0061\u0074\u0020\u006e\u0061m\u0065")
}

// Slides returns the slides in the presentation.
func (_cec *Presentation) Slides() []Slide {
	_gfd := []Slide{}
	for _dcf, _adae := range _cec._ddc {
		_gfd = append(_gfd, Slide{_cec._cc.SldIdLst.SldId[_dcf], _adae, _cec})
	}
	return _gfd
}

// AddTable adds an empty table to a slide.
func (_deg Slide) AddTable() *_ff.Table {
	_ccf := _gb.NewCT_GroupShapeChoice()
	_deg._caed.CSld.SpTree.Choice = append(_deg._caed.CSld.SpTree.Choice, _ccf)
	_bde := _gb.NewCT_GraphicalObjectFrame()
	_ccf.GraphicFrame = append(_ccf.GraphicFrame, _bde)
	_bde.Xfrm.Off = _dbb.NewCT_Point2D()
	_fedb := int64(1)
	_bde.Xfrm.Off.XAttr = _dbb.ST_Coordinate{ST_CoordinateUnqualified: &_fedb}
	_bde.Xfrm.Off.YAttr = _dbb.ST_Coordinate{ST_CoordinateUnqualified: &_fedb}
	_bbdc := _bde.Graphic.CT_GraphicalObject.GraphicData
	_bbdc.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073.\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006dl/\u0032\u0030\u0030\u0036\u002f\u0074\u0061\u0062\u006c\u0065"
	_cgd := _ff.NewTableWithXfrm(_bde.Xfrm)
	_bbdc.Any = append(_bbdc.Any, _cgd.X())
	return _cgd
}

// New initializes and reurns a new presentation
func New() *Presentation {
	_bfd := _ga()
	_bfd.ContentTypes.AddOverride("/\u0070\u0070\u0074\u002fpr\u0065s\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006ff\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u006d\u006c\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	_bfd.Rels.AddRelationship("\u0064\u006f\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0063\u006f\u0072e\u002e\u0078\u006d\u006c", "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006ba\u0067\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002f\u006d\u0065\u0074\u0061\u0064\u0061\u0074\u0061/\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006f\u0070e\u0072\u0074i\u0065\u0073")
	_bfd.Rels.AddRelationship("\u0064\u006fc\u0050\u0072\u006fp\u0073\u002f\u0061\u0070\u0070\u002e\u0078\u006d\u006c", "\u0068t\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073.\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069p\u0073\u002f\u0065x\u0074\u0065\u006e\u0064\u0065d\u002d\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	_bfd.Rels.AddRelationship("p\u0070t\u002f\u0070\u0072\u0065\u0073\u0065\u006e\u0074a\u0074\u0069\u006f\u006e.x\u006d\u006c", "\u0068\u0074\u0074\u0070\u003a\u002f\u002fs\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072g\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006fc\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	_bfd.Rels.AddRelationship("\u0070\u0070\u0074\u002f\u0070\u0072\u0065\u0073\u0050\u0072\u006f\u0070s\u002e\u0078\u006d\u006c", "ht\u0074\u0070\u003a\u002f\u002f\u0073\u0063he\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006et\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068i\u0070s\u002f\u0070\u0072\u0065\u0073\u0050\u0072\u006f\u0070\u0073")
	_bfd.Rels.AddRelationship("\u0070\u0070\u0074\u002f\u0076\u0069\u0065\u0077\u0050\u0072\u006f\u0070s\u002e\u0078\u006d\u006c", "ht\u0074\u0070\u003a\u002f\u002f\u0073\u0063he\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006et\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068i\u0070s\u002f\u0076\u0069\u0065\u0077\u0050\u0072\u006f\u0070\u0073")
	_bfd.Rels.AddRelationship("\u0070\u0070\u0074\u002fta\u0062\u006c\u0065\u0053\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", "\u0068\u0074\u0074\u0070\u003a\u002f\u002fs\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006cf\u006fr\u006d\u0061\u0074\u0073\u002e\u006fr\u0067\u002f\u006f\u0066\u0066\u0069\u0063e\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073\u002f\u0074\u0061\u0062\u006c\u0065\u0053\u0074\u0079\u006ce\u0073")
	_bfd._cc.SldMasterIdLst = _gb.NewCT_SlideMasterIdList()
	_dgd := _gb.NewSldMaster()
	_dgd.ClrMap.Bg1Attr = _dbb.ST_ColorSchemeIndexLt1
	_dgd.ClrMap.Bg2Attr = _dbb.ST_ColorSchemeIndexLt2
	_dgd.ClrMap.Tx1Attr = _dbb.ST_ColorSchemeIndexDk1
	_dgd.ClrMap.Tx2Attr = _dbb.ST_ColorSchemeIndexDk2
	_dgd.ClrMap.Accent1Attr = _dbb.ST_ColorSchemeIndexAccent1
	_dgd.ClrMap.Accent2Attr = _dbb.ST_ColorSchemeIndexAccent2
	_dgd.ClrMap.Accent3Attr = _dbb.ST_ColorSchemeIndexAccent3
	_dgd.ClrMap.Accent4Attr = _dbb.ST_ColorSchemeIndexAccent4
	_dgd.ClrMap.Accent5Attr = _dbb.ST_ColorSchemeIndexAccent5
	_dgd.ClrMap.Accent6Attr = _dbb.ST_ColorSchemeIndexAccent6
	_dgd.ClrMap.HlinkAttr = _dbb.ST_ColorSchemeIndexHlink
	_dgd.ClrMap.FolHlinkAttr = _dbb.ST_ColorSchemeIndexFolHlink
	_bfd._ea = append(_bfd._ea, _dgd)
	_cbd := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideMasterType, 1)
	_bfd.ContentTypes.AddOverride(_cbd, _db.SlideMasterContentType)
	_ccc := _bfd._bed.AddAutoRelationship(_db.DocTypePresentation, _db.OfficeDocumentType, 1, _db.SlideMasterType)
	_dcb := _gb.NewCT_SlideMasterIdListEntry()
	_dcb.IdAttr = _db.Uint32(2147483648)
	_dcb.RIdAttr = _ccc.ID()
	_bfd._cc.SldMasterIdLst.SldMasterId = append(_bfd._cc.SldMasterIdLst.SldMasterId, _dcb)
	_dbd := _ff.NewRelationships()
	_bfd._accc = append(_bfd._accc, _dbd)
	_agf := _gb.NewSldLayout()
	_geb := _dbd.AddAutoRelationship(_db.DocTypePresentation, _db.SlideMasterType, 1, _db.SlideLayoutType)
	_cdg := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideLayoutType, 1)
	_bfd.ContentTypes.AddOverride(_cdg, _db.SlideLayoutContentType)
	_dbd.AddAutoRelationship(_db.DocTypePresentation, _db.SlideMasterType, 1, _db.ThemeType)
	_bfd._fbd = append(_bfd._fbd, _agf)
	_dgd.SldLayoutIdLst = _gb.NewCT_SlideLayoutIdList()
	_adf := _gb.NewCT_SlideLayoutIdListEntry()
	_adf.IdAttr = _db.Uint32(2147483649)
	_adf.RIdAttr = _geb.ID()
	_dgd.SldLayoutIdLst.SldLayoutId = append(_dgd.SldLayoutIdLst.SldLayoutId, _adf)
	_egfc := _ff.NewRelationships()
	_bfd._beb = append(_bfd._beb, _egfc)
	_egfc.AddAutoRelationship(_db.DocTypePresentation, _db.SlideType, 1, _db.SlideMasterType)
	_bfd._cc.NotesSz.CxAttr = 6858000
	_bfd._cc.NotesSz.CyAttr = 9144000
	_aab := _dbb.NewTheme()
	_aab.NameAttr = _db.String("\u0075n\u0069o\u0066\u0066\u0069\u0063\u0065\u0020\u0054\u0068\u0065\u006d\u0065")
	_aab.ThemeElements.ClrScheme.NameAttr = "\u004f\u0066\u0066\u0069\u0063\u0065"
	_aab.ThemeElements.ClrScheme.Dk1.SysClr = _dbb.NewCT_SystemColor()
	_aab.ThemeElements.ClrScheme.Dk1.SysClr.LastClrAttr = _db.String("\u0030\u0030\u0030\u0030\u0030\u0030")
	_aab.ThemeElements.ClrScheme.Dk1.SysClr.ValAttr = _dbb.ST_SystemColorValWindowText
	_aab.ThemeElements.ClrScheme.Lt1.SysClr = _dbb.NewCT_SystemColor()
	_aab.ThemeElements.ClrScheme.Lt1.SysClr.LastClrAttr = _db.String("\u0066\u0066\u0066\u0066\u0066\u0066")
	_aab.ThemeElements.ClrScheme.Lt1.SysClr.ValAttr = _dbb.ST_SystemColorValWindow
	_aab.ThemeElements.ClrScheme.Dk2.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Dk2.SrgbClr.ValAttr = "\u0034\u0034\u0035\u0034\u0036\u0061"
	_aab.ThemeElements.ClrScheme.Lt2.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Lt2.SrgbClr.ValAttr = "\u0065\u0037\u0065\u0037\u0065\u0036"
	_aab.ThemeElements.ClrScheme.Accent1.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Accent1.SrgbClr.ValAttr = "\u0034\u0034\u0037\u0032\u0063\u0034"
	_aab.ThemeElements.ClrScheme.Accent2.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Accent2.SrgbClr.ValAttr = "\u0065\u0064\u0037\u0064\u0033\u0031"
	_aab.ThemeElements.ClrScheme.Accent3.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Accent3.SrgbClr.ValAttr = "\u0061\u0035\u0061\u0035\u0061\u0035"
	_aab.ThemeElements.ClrScheme.Accent4.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Accent4.SrgbClr.ValAttr = "\u0066\u0066\u0063\u0030\u0030\u0030"
	_aab.ThemeElements.ClrScheme.Accent5.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Accent5.SrgbClr.ValAttr = "\u0035\u0062\u0039\u0062\u0064\u0035"
	_aab.ThemeElements.ClrScheme.Accent6.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Accent6.SrgbClr.ValAttr = "\u0037\u0030\u0061\u0064\u0034\u0037"
	_aab.ThemeElements.ClrScheme.Hlink.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.Hlink.SrgbClr.ValAttr = "\u0030\u0035\u0036\u0033\u0063\u0031"
	_aab.ThemeElements.ClrScheme.FolHlink.SrgbClr = _dbb.NewCT_SRgbColor()
	_aab.ThemeElements.ClrScheme.FolHlink.SrgbClr.ValAttr = "\u0039\u0035\u0034\u0066\u0037\u0032"
	_aab.ThemeElements.FontScheme.NameAttr = "\u004f\u0066\u0066\u0069\u0063\u0065"
	_aab.ThemeElements.FontScheme.MajorFont.Latin.TypefaceAttr = "\u0043\u0061\u006c\u0069\u0062\u0072\u0069\u0020\u004c\u0069\u0067\u0068\u0074"
	_aab.ThemeElements.FontScheme.MinorFont.Latin.TypefaceAttr = "\u0043a\u006c\u0069\u0062\u0072\u0069"
	_aab.ThemeElements.FmtScheme.NameAttr = _db.String("\u004f\u0066\u0066\u0069\u0063\u0065")
	_ef := _dbb.NewEG_FillProperties()
	_aab.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(_aab.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, _ef)
	_ef.SolidFill = &_dbb.CT_SolidColorFillProperties{SchemeClr: &_dbb.CT_SchemeColor{ValAttr: _dbb.ST_SchemeColorValPhClr}}
	_ef = _dbb.NewEG_FillProperties()
	_aab.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(_aab.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, _ef)
	_aab.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(_aab.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, _ef)
	_ef.GradFill = &_dbb.CT_GradientFillProperties{RotWithShapeAttr: _db.Bool(true), GsLst: &_dbb.CT_GradientStopList{}, Lin: &_dbb.CT_LinearShadeProperties{}}
	_ef.GradFill.Lin.AngAttr = _db.Int32(5400000)
	_ef.GradFill.Lin.ScaledAttr = _db.Bool(false)
	_eb := _dbb.NewCT_GradientStop()
	_eb.PosAttr.ST_PositiveFixedPercentageDecimal = _db.Int32(0)
	_eb.SchemeClr = &_dbb.CT_SchemeColor{ValAttr: _dbb.ST_SchemeColorValPhClr}
	_ef.GradFill.GsLst.Gs = append(_ef.GradFill.GsLst.Gs, _eb)
	_eb = _dbb.NewCT_GradientStop()
	_eb.PosAttr.ST_PositiveFixedPercentageDecimal = _db.Int32(50000)
	_eb.SchemeClr = &_dbb.CT_SchemeColor{ValAttr: _dbb.ST_SchemeColorValPhClr}
	_ef.GradFill.GsLst.Gs = append(_ef.GradFill.GsLst.Gs, _eb)
	_aab.ThemeElements.FmtScheme.LnStyleLst = _dbb.NewCT_LineStyleList()
	for _dda := 0; _dda < 3; _dda++ {
		_gf := _dbb.NewCT_LineProperties()
		_gf.WAttr = _db.Int32(int32(6350 * (_dda + 1)))
		_gf.CapAttr = _dbb.ST_LineCapFlat
		_gf.CmpdAttr = _dbb.ST_CompoundLineSng
		_gf.AlgnAttr = _dbb.ST_PenAlignmentCtr
		_aab.ThemeElements.FmtScheme.LnStyleLst.Ln = append(_aab.ThemeElements.FmtScheme.LnStyleLst.Ln, _gf)
	}
	_aab.ThemeElements.FmtScheme.EffectStyleLst = _dbb.NewCT_EffectStyleList()
	for _bag := 0; _bag < 3; _bag++ {
		_de := _dbb.NewCT_EffectStyleItem()
		_de.EffectLst = _dbb.NewCT_EffectList()
		_aab.ThemeElements.FmtScheme.EffectStyleLst.EffectStyle = append(_aab.ThemeElements.FmtScheme.EffectStyleLst.EffectStyle, _de)
	}
	_aedb := _dbb.NewEG_FillProperties()
	_aedb.SolidFill = &_dbb.CT_SolidColorFillProperties{SchemeClr: &_dbb.CT_SchemeColor{ValAttr: _dbb.ST_SchemeColorValPhClr}}
	_aab.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(_aab.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties, _aedb)
	_aab.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(_aab.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties, _aedb)
	_aab.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(_aab.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties, _ef)
	_bfd._bcd = append(_bfd._bcd, _aab)
	_eac := _db.AbsoluteFilename(_db.DocTypePresentation, _db.ThemeType, 1)
	_bfd.ContentTypes.AddOverride(_eac, _db.ThemeContentType)
	_bfd._bed.AddAutoRelationship(_db.DocTypePresentation, _db.OfficeDocumentType, 1, _db.ThemeType)
	_eff := _ff.NewRelationships()
	_bfd._gea = append(_bfd._gea, _eff)
	return _bfd
}
func (_abaa *Presentation) onNewRelationship(_fdcg *_fg.DecodeMap, _bdag, _ee string, _ada []*_bd.File, _gaeb *_ac.Relationship, _aece _fg.Target) error {
	_adc := _db.DocTypePresentation
	switch _ee {
	case _db.OfficeDocumentType:
		_abaa._cc = _gb.NewPresentation()
		_fdcg.AddTarget(_bdag, _abaa._cc, _ee, 0)
		_fdcg.AddTarget(_fg.RelationsPathFor(_bdag), _abaa._bed.X(), _ee, 0)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, 0)
	case _db.CorePropertiesType:
		_fdcg.AddTarget(_bdag, _abaa.CoreProperties.X(), _ee, 0)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, 0)
	case _db.CustomPropertiesType:
		_fdcg.AddTarget(_bdag, _abaa.CustomProperties.X(), _ee, 0)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, 0)
	case _db.PresentationPropertiesType:
		_fdcg.AddTarget(_bdag, _abaa._ec.X(), _ee, 0)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, 0)
	case _db.ViewPropertiesType:
		_fdcg.AddTarget(_bdag, _abaa._gbg.X(), _ee, 0)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, 0)
	case _db.TableStylesType:
		_fdcg.AddTarget(_bdag, _abaa._ad.X(), _ee, 0)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, 0)
	case _db.HyperLinkType:
		_bdbb := _dbb.NewCT_Hyperlink()
		_adea := uint32(len(_abaa._cde))
		_fdcg.AddTarget(_bdag, _bdbb, _ee, _adea)
		_abaa._cde = append(_abaa._cde, _bdbb)
	case _db.CustomXMLType:
		_cccc := &_db.XSDAny{}
		_bfa := uint32(len(_abaa._cgb))
		_fdcg.AddTarget(_bdag, _cccc, _ee, _bfa)
		_abaa._cgb = append(_abaa._cgb, _cccc)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._cgb))
	case _db.ChartType:
		_gfb := _fce.NewChartSpace()
		_egb := uint32(len(_abaa._eag))
		_fdcg.AddTarget(_bdag, _gfb, _ee, _egb)
		_abaa._eag = append(_abaa._eag, _gfb)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._eag))
	case _db.HandoutMasterType:
		_dceb := _gb.NewHandoutMaster()
		_gbe := uint32(len(_abaa._ecc))
		_fdcg.AddTarget(_bdag, _dceb, _ee, _gbe)
		_abaa._ecc = append(_abaa._ecc, _dceb)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._ecc))
	case _db.NotesMasterType:
		_eaba := _gb.NewNotesMaster()
		_egfa := uint32(len(_abaa._gd))
		_fdcg.AddTarget(_bdag, _eaba, _ee, _egfa)
		_abaa._gd = append(_abaa._gd, _eaba)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._gd))
	case _db.ExtendedPropertiesType:
		_fdcg.AddTarget(_bdag, _abaa.AppProperties.X(), _ee, 0)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, 0)
	case _db.SlideType:
		_aca := _gb.NewSld()
		_abaa._ddc = append(_abaa._ddc, _aca)
		_fdcg.AddTarget(_bdag, _aca, _ee, uint32(len(_abaa._ddc)))
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._ddc))
		_eea := _ff.NewRelationships()
		_fdcg.AddTarget(_fg.RelationsPathFor(_bdag), _eea.X(), _ee, 0)
		_abaa._fba = append(_abaa._fba, _eea)
	case _db.SlideMasterType:
		_gef := _gb.NewSldMaster()
		if !_fdcg.AddTarget(_bdag, _gef, _ee, uint32(len(_abaa._ea)+1)) {
			return nil
		}
		_abaa._ea = append(_abaa._ea, _gef)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._ea))
		_efb := _ff.NewRelationships()
		_fdcg.AddTarget(_fg.RelationsPathFor(_bdag), _efb.X(), _ee, 0)
		_abaa._accc = append(_abaa._accc, _efb)
	case _db.SlideLayoutType:
		_ggba := _gb.NewSldLayout()
		if !_fdcg.AddTarget(_bdag, _ggba, _ee, uint32(len(_abaa._fbd)+1)) {
			return nil
		}
		_abaa._fbd = append(_abaa._fbd, _ggba)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._fbd))
		_fab := _ff.NewRelationships()
		_fdcg.AddTarget(_fg.RelationsPathFor(_bdag), _fab.X(), _ee, 0)
		_abaa._beb = append(_abaa._beb, _fab)
	case _db.ThumbnailType:
		for _dag, _fde := range _ada {
			if _fde == nil {
				continue
			}
			if _fde.Name == _bdag {
				_ebee, _ffd := _fde.Open()
				if _ffd != nil {
					return _dg.Errorf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073", _ffd)
				}
				_abaa.Thumbnail, _, _ffd = _fcd.Decode(_ebee)
				_ebee.Close()
				if _ffd != nil {
					return _dg.Errorf("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073", _ffd)
				}
				_ada[_dag] = nil
			}
		}
	case _db.ThemeType:
		_dcge := _dbb.NewTheme()
		if !_fdcg.AddTarget(_bdag, _dcge, _ee, uint32(len(_abaa._bcd)+1)) {
			return nil
		}
		_abaa._bcd = append(_abaa._bcd, _dcge)
		_gaeb.TargetAttr = _db.RelativeFilename(_adc, _aece.Typ, _ee, len(_abaa._bcd))
		_dee := _ff.NewRelationships()
		_fdcg.AddTarget(_fg.RelationsPathFor(_bdag), _dee.X(), _ee, 0)
		_abaa._gea = append(_abaa._gea, _dee)
	case _db.ImageType:
		_bdag = _c.Clean(_bdag)
		if _gce, _ecg := _abaa._bbd[_bdag]; _ecg {
			_gaeb.TargetAttr = _gce
			return nil
		}
		_eeb := ""
		for _dfeb, _dbdc := range _ada {
			if _dbdc == nil {
				continue
			}
			if _dbdc.Name == _bdag {
				_ege, _caf := _fg.ExtractToDiskTmp(_dbdc, _abaa.TmpPath)
				if _caf != nil {
					return _caf
				}
				_ddg, _caf := _ff.ImageFromStorage(_ege)
				if _caf != nil {
					return _caf
				}
				_eeb = _ddg.Format
				_efe := _ff.MakeImageRef(_ddg, &_abaa.DocBase, _abaa._bed)
				_abaa.Images = append(_abaa.Images, _efe)
				_ada[_dfeb] = nil
				_fdcg.RecordIndex(_bdag, len(_abaa.Images))
				break
			}
		}
		_faf := _fdcg.IndexFor(_bdag)
		_gaeb.TargetAttr = _db.RelativeImageFilename(_adc, _aece.Typ, _ee, _faf, _eeb)
		_abaa._bbd[_bdag] = _gaeb.TargetAttr
	default:
		_db.Log("\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u0074\u0067\u0074\u003a\u0020\u0025\u0073", _ee, _bdag)
	}
	return nil
}

// NotesTextViewPr returns the NotesTextViewPr property.
func (_bdeg ViewProperties) NotesTextViewPr() *_gb.CT_NotesTextViewProperties {
	return _bdeg._bbab.NotesTextViewPr
}

// X returns TextBox's underlying *pml.CT_Shape.
func (_aaeg TextBox) X() *_gb.CT_Shape { return _aaeg._cfbc }

// SlideViewPr returns the SlideViewPr property.
func (_bbbg ViewProperties) SlideViewPr() *_gb.CT_SlideViewProperties { return _bbbg._bbab.SlideViewPr }

// SetOffsetY sets vertical offset of text box in distance units (see measurement package).
func (_adb TextBox) SetOffsetY(offY float64) {
	_agg := _adb.getOff()
	_ecdg := _d.ToEMU(offY)
	_agg.YAttr = _dbb.ST_Coordinate{ST_CoordinateUnqualified: &_ecdg}
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (_acac *Presentation) GetImageByRelID(relID string) (_ff.ImageRef, bool) {
	for _, _bebd := range _acac.Images {
		if _bebd.RelID() == relID {
			return _bebd, true
		}
	}
	return _ff.ImageRef{}, false
}

// Close closes the presentation, removing any temporary files that might have been
// created when opening a document.
func (_agac *Presentation) Close() error {
	if _agac.TmpPath != "" {
		return _bb.RemoveAll(_agac.TmpPath)
	}
	return nil
}

// HtmlPubPr returns the HtmlPubPr property.
func (_cae PresentationProperties) HtmlPubPr() *_gb.CT_HtmlPublishProperties {
	return _cae._aed.HtmlPubPr
}

// Type returns the type of the slide layout.
func (_caa SlideLayout) Type() _gb.ST_SlideLayoutType { return _caa._cca.TypeAttr }

// X returns the inner wrapped XML type.
func (_fge PlaceHolder) X() *_gb.CT_Shape { return _fge._fb }

// Paragraphs returns the paragraphs defined in the placeholder.
func (_aec PlaceHolder) Paragraphs() []_g.Paragraph {
	_eg := []_g.Paragraph{}
	for _, _cb := range _aec._fb.TxBody.P {
		_eg = append(_eg, _g.MakeParagraph(_cb))
	}
	return _eg
}

var _cac = false

// ClrMru returns the ClrMru property.
func (_ge PresentationProperties) ClrMru() *_dbb.CT_ColorMRU { return _ge._aed.ClrMru }

// PlaceHolder is a place holder from a slide.
type PlaceHolder struct {
	_fb *_gb.CT_Shape
	_ag *_gb.Sld
}

func (_ddac *Presentation) nextSlideID() uint32 {
	_gec := uint32(256)
	for _, _ebe := range _ddac._cc.SldIdLst.SldId {
		if _ebe.IdAttr >= _gec {
			_gec = _ebe.IdAttr + 1
		}
	}
	return _gec
}

// NormalViewPr returns the NormalViewPr property.
func (_efee ViewProperties) NormalViewPr() *_gb.CT_NormalViewProperties {
	return _efee._bbab.NormalViewPr
}
func _cad(_agd []*_gb.CT_GroupShapeChoice) []*_gb.CT_GroupShapeChoice {
	var _bec []*_gb.CT_GroupShapeChoice
	for _, _ged := range _agd {
		if len(_ged.Pic) == 0 {
			_bec = append(_bec, _ged)
		}
	}
	return _bec
}

// Clear clears the placeholder contents and adds a single empty paragraph.  The
// empty paragrah is required by PowerPoint or it will report the file as being
// invalid.
func (_cfc PlaceHolder) Clear() {
	_cfc.ClearAll()
	_bgg := _dbb.NewCT_TextParagraph()
	_cfc._fb.TxBody.P = []*_dbb.CT_TextParagraph{_bgg}
	_bgg.EndParaRPr = _dbb.NewCT_TextCharacterProperties()
	_bgg.EndParaRPr.LangAttr = _db.String("\u0065\u006e\u002dU\u0053")
}

// OpenTemplate opens a template file.
func OpenTemplate(fn string) (*Presentation, error) {
	_cf, _fad := Open(fn)
	if _fad != nil {
		return nil, _fad
	}
	return _cf, nil
}

// PlaceHolders returns all of the content place holders within a given slide.
func (_abaf Slide) PlaceHolders() []PlaceHolder {
	_eaab := []PlaceHolder{}
	for _, _ddd := range _abaf._caed.CSld.SpTree.Choice {
		for _, _adeac := range _ddd.Sp {
			if _adeac.NvSpPr != nil && _adeac.NvSpPr.NvPr != nil && _adeac.NvSpPr.NvPr.Ph != nil {
				_eaab = append(_eaab, PlaceHolder{_adeac, _abaf._caed})
			}
		}
	}
	return _eaab
}

// SlideLayout is a layout from which slides can be created.
type SlideLayout struct{ _cca *_gb.SldLayout }

// AddSlideWithLayout adds a new slide with content copied from a layout.  Normally you should
// use AddDefaultSlideWithLayout as it will do some post processing similar to PowerPoint to
// clear place holder text, etc.
func (_fee *Presentation) AddSlideWithLayout(l SlideLayout) (Slide, error) {
	_ded := _gb.NewCT_SlideIdListEntry()
	_ded.IdAttr = 256
	for _, _ffa := range _fee._cc.SldIdLst.SldId {
		if _ffa.IdAttr >= _ded.IdAttr {
			_ded.IdAttr = _ffa.IdAttr + 1
		}
	}
	_fee._cc.SldIdLst.SldId = append(_fee._cc.SldIdLst.SldId, _ded)
	_gbb := _gb.NewSld()
	_ddacf := _ba.Buffer{}
	_dfe := _ce.NewEncoder(&_ddacf)
	_efd := _ce.StartElement{Name: _ce.Name{Local: "\u0073\u006c\u0069d\u0065"}}
	_efd.Attr = append(_efd.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069o\u006e\u006d\u006c\u002f\u0032\u00300\u0036\u002f\u006da\u0069\u006e"})
	_efd.Attr = append(_efd.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	_efd.Attr = append(_efd.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0070"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069o\u006e\u006d\u006c\u002f\u0032\u00300\u0036\u002f\u006da\u0069\u006e"})
	_efd.Attr = append(_efd.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	_efd.Attr = append(_efd.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0073\u0068"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	_efd.Attr = append(_efd.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	if _fd := l._cca.CSld.MarshalXML(_dfe, _efd); _fd != nil {
		return Slide{}, _fd
	}
	_dfe.Flush()
	_cea := _ce.NewDecoder(&_ddacf)
	_gbb.CSld = _gb.NewCT_CommonSlideData()
	if _bbe := _cea.Decode(_gbb.CSld); _bbe != nil {
		return Slide{}, _bbe
	}
	_gbb.CSld.NameAttr = nil
	_gbb.CSld.SpTree.Choice = _cad(_gbb.CSld.SpTree.Choice)
	_fee._ddc = append(_fee._ddc, _gbb)
	_cbdf := _fee._bed.AddAutoRelationship(_db.DocTypePresentation, _db.OfficeDocumentType, len(_fee._ddc), _db.SlideType)
	_ded.RIdAttr = _cbdf.ID()
	_gaa := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideType, len(_fee._ddc))
	_fee.ContentTypes.AddOverride(_gaa, _db.SlideContentType)
	_feb := _ff.NewRelationships()
	_fee._fba = append(_fee._fba, _feb)
	_gaf := len(_fee._fba) - 1
	for _agfa, _dfb := range _fee._fbd {
		if _dfb == l.X() {
			_ceab := _fee._beb[_agfa]
			for _, _fag := range _ceab.X().Relationship {
				if _fag.TypeAttr != _db.SlideMasterType {
					_fee._fba[_gaf].X().Relationship = append(_fee._fba[_gaf].X().Relationship, _fag)
				}
			}
			_feb.AddAutoRelationship(_db.DocTypePresentation, _db.SlideType, _agfa+1, _db.SlideLayoutType)
		}
	}
	_fcbe := Slide{_ded, _gbb, _fee}
	return _fcbe, nil
}

// SlideLayouts returns the slide layouts defined in the presentation.
func (_dfcg *Presentation) SlideLayouts() []SlideLayout {
	_bdab := []SlideLayout{}
	for _, _dbf := range _dfcg._fbd {
		_bdab = append(_bdab, SlideLayout{_dbf})
	}
	return _bdab
}

// Image is an image within a slide.
type Image struct{ _be *_gb.CT_Picture }

func (_dgg *Presentation) save(_bebb _ae.Writer, _gge bool) error {
	if _cbf := _dgg._cc.Validate(); _cbf != nil {
		_bg.Printf("\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u0020i\u006e\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0073", _cbf)
	}
	if !_b.GetLicenseKey().IsLicensed() && !_cac {
		_dg.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
		_dg.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _e.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	}
	if _gge {
		_dgg.ContentTypes.RemoveOverride("\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006ff\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u006d\u006c\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
		_dgg.ContentTypes.EnsureOverride("/\u0070\u0070\u0074\u002fpr\u0065s\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u002e\u0078\u006d\u006c", "\u0061\u0070pl\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074s\u002d\u006f\u0066\u0066ic\u0065\u0064o\u0063u\u006d\u0065\u006e\u0074\u002e\u0070r\u0065\u0073\u0065n\u0074\u0061t\u0069\u006f\u006e\u006d\u006c\u002e\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065.\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	} else {
		_dgg.ContentTypes.RemoveOverride("\u0061\u0070pl\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074s\u002d\u006f\u0066\u0066ic\u0065\u0064o\u0063u\u006d\u0065\u006e\u0074\u002e\u0070r\u0065\u0073\u0065n\u0074\u0061t\u0069\u006f\u006e\u006d\u006c\u002e\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065.\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
		_dgg.ContentTypes.EnsureOverride("/\u0070\u0070\u0074\u002fpr\u0065s\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006ff\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u006d\u006c\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	}
	_fff := _db.DocTypePresentation
	_gfc := _bd.NewWriter(_bebb)
	defer _gfc.Close()
	if _cfb := _fg.MarshalXML(_gfc, _db.BaseRelsFilename, _dgg.Rels.X()); _cfb != nil {
		return _cfb
	}
	if _ecd := _fg.MarshalXMLByType(_gfc, _fff, _db.ExtendedPropertiesType, _dgg.AppProperties.X()); _ecd != nil {
		return _ecd
	}
	if _fcab := _fg.MarshalXMLByType(_gfc, _fff, _db.CorePropertiesType, _dgg.CoreProperties.X()); _fcab != nil {
		return _fcab
	}
	if _ffb := _fg.MarshalXMLByType(_gfc, _fff, _db.PresentationPropertiesType, _dgg._ec.X()); _ffb != nil {
		return _ffb
	}
	if _cdd := _fg.MarshalXMLByType(_gfc, _fff, _db.ViewPropertiesType, _dgg._gbg.X()); _cdd != nil {
		return _cdd
	}
	if _bfb := _fg.MarshalXMLByType(_gfc, _fff, _db.TableStylesType, _dgg._ad.X()); _bfb != nil {
		return _bfb
	}
	if _dgg.CustomProperties.X() != nil {
		if _eab := _fg.MarshalXMLByType(_gfc, _fff, _db.CustomPropertiesType, _dgg.CustomProperties.X()); _eab != nil {
			return _eab
		}
	}
	if _dgg.Thumbnail != nil {
		_geg, _abc := _gfc.Create("\u0064\u006f\u0063Pr\u006f\u0070\u0073\u002f\u0074\u0068\u0075\u006d\u0062\u006e\u0061\u0069\u006c\u002e\u006a\u0070\u0065\u0067")
		if _abc != nil {
			return _abc
		}
		if _abg := _fe.Encode(_geg, _dgg.Thumbnail, nil); _abg != nil {
			return _abg
		}
	}
	_aee := _db.AbsoluteFilename(_fff, _db.OfficeDocumentType, 0)
	if _fcbg := _fg.MarshalXML(_gfc, _aee, _dgg._cc); _fcbg != nil {
		return _fcbg
	}
	if _ggb := _fg.MarshalXML(_gfc, _fg.RelationsPathFor(_aee), _dgg._bed.X()); _ggb != nil {
		return _ggb
	}
	for _ccca, _cee := range _dgg._ddc {
		_add := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideType, _ccca+1)
		_fg.MarshalXML(_gfc, _add, _cee)
		if !_dgg._fba[_ccca].IsEmpty() {
			_bcg := _fg.RelationsPathFor(_add)
			_fg.MarshalXML(_gfc, _bcg, _dgg._fba[_ccca].X())
		}
	}
	for _cab, _dac := range _dgg._ea {
		_fdf := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideMasterType, _cab+1)
		_fg.MarshalXML(_gfc, _fdf, _dac)
		if !_dgg._accc[_cab].IsEmpty() {
			_dbe := _fg.RelationsPathFor(_fdf)
			_fg.MarshalXML(_gfc, _dbe, _dgg._accc[_cab].X())
		}
	}
	for _bdaf, _bfc := range _dgg._fbd {
		_aecc := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideLayoutType, _bdaf+1)
		_fg.MarshalXML(_gfc, _aecc, _bfc)
		if !_dgg._beb[_bdaf].IsEmpty() {
			_dga := _fg.RelationsPathFor(_aecc)
			_fg.MarshalXML(_gfc, _dga, _dgg._beb[_bdaf].X())
		}
	}
	for _geab, _gbf := range _dgg._bcd {
		_ddf := _db.AbsoluteFilename(_db.DocTypePresentation, _db.ThemeType, _geab+1)
		_fg.MarshalXML(_gfc, _ddf, _gbf)
		if !_dgg._gea[_geab].IsEmpty() {
			_efdb := _fg.RelationsPathFor(_ddf)
			_fg.MarshalXML(_gfc, _efdb, _dgg._gea[_geab].X())
		}
	}
	for _gbba, _fced := range _dgg._eag {
		_daa := _db.AbsoluteFilename(_fff, _db.ChartType, _gbba+1)
		_fg.MarshalXML(_gfc, _daa, _fced)
	}
	for _abf, _becg := range _dgg._ecc {
		_gff := _db.AbsoluteFilename(_fff, _db.HandoutMasterType, _abf+1)
		_fg.MarshalXML(_gfc, _gff, _becg)
	}
	for _ace, _abb := range _dgg._gd {
		_gag := _db.AbsoluteFilename(_fff, _db.NotesMasterType, _ace+1)
		_fg.MarshalXML(_gfc, _gag, _abb)
	}
	for _beg, _gae := range _dgg._cgb {
		_ggg := _db.AbsoluteFilename(_fff, _db.CustomXMLType, _beg+1)
		_fg.MarshalXML(_gfc, _ggg, _gae)
	}
	for _eaf, _gagd := range _dgg.Images {
		if _cga := _ff.AddImageToZip(_gfc, _gagd, _eaf+1, _db.DocTypePresentation); _cga != nil {
			return _cga
		}
	}
	_dgg.ContentTypes.EnsureDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_dgg.ContentTypes.EnsureDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_dgg.ContentTypes.EnsureDefault("\u006a\u0070\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_dgg.ContentTypes.EnsureDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	if _bfdf := _fg.MarshalXML(_gfc, _db.ContentTypesFilename, _dgg.ContentTypes.X()); _bfdf != nil {
		return _bfdf
	}
	if _dcc := _dgg.WriteExtraFiles(_gfc); _dcc != nil {
		return _dcc
	}
	return nil
}

// SetTextAnchor controls the text anchoring
func (_ead TextBox) SetTextAnchor(a _dbb.ST_TextAnchoringType) {
	_ead._cfbc.TxBody.BodyPr = _dbb.NewCT_TextBodyProperties()
	_ead._cfbc.TxBody.BodyPr.AnchorAttr = a
}

// Properties returns the properties of the TextBox.
func (_fa Image) Properties() _g.ShapeProperties {
	if _fa._be.SpPr == nil {
		_fa._be.SpPr = _dbb.NewCT_ShapeProperties()
	}
	return _g.MakeShapeProperties(_fa._be.SpPr)
}

// ShowPr returns the ShowPr property.
func (_dcg PresentationProperties) ShowPr() *_gb.CT_ShowProperties { return _dcg._aed.ShowPr }

// Remove removes a placeholder from a presentation.
func (_dc PlaceHolder) Remove() error {
	for _aga, _df := range _dc._ag.CSld.SpTree.Choice {
		for _, _aea := range _df.Sp {
			if _aea == _dc._fb {
				copy(_dc._ag.CSld.SpTree.Choice[_aga:], _dc._ag.CSld.SpTree.Choice[_aga+1:])
				_dc._ag.CSld.SpTree.Choice = _dc._ag.CSld.SpTree.Choice[0 : len(_dc._ag.CSld.SpTree.Choice)-1]
				return nil
			}
		}
	}
	return _e.New("\u0070\u006c\u0061\u0063\u0065\u0068\u006f\u006c\u0064\u0065r\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u006c\u0069\u0064\u0065")
}

// NotesViewPr returns the NotesViewPr property.
func (_bacd ViewProperties) NotesViewPr() *_gb.CT_NotesViewProperties { return _bacd._bbab.NotesViewPr }

// X returns the inner wrapped XML type.
func (_ccdg Slide) X() *_gb.Sld { return _ccdg._caed }

// X returns the inner wrapped XML type.
func (_ebbb SlideMaster) X() *_gb.SldMaster { return _ebbb._bbc }

// SlideMaster is the slide master for a presentation.
type SlideMaster struct {
	_gcfc *Presentation
	_dbg  _ff.Relationships
	_bbc  *_gb.SldMaster
}

// ShowCommentsAttr returns the WebPr property.
func (_agge ViewProperties) ShowCommentsAttr() *bool { return _agge._bbab.ShowCommentsAttr }

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_gfcc *Presentation) AddImage(i _ff.Image) (_ff.ImageRef, error) {
	_bgc := _ff.MakeImageRef(i, &_gfcc.DocBase, _gfcc._bed)
	if i.Data == nil && i.Path == "" {
		return _bgc, _e.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _bgc, _e.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _bgc, _e.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	if i.Path != "" {
		_dfd := _bb.Add(i.Path)
		if _dfd != nil {
			return _bgc, _dfd
		}
	}
	_gfcc.Images = append(_gfcc.Images, _bgc)
	_gfcc.ContentTypes.EnsureDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_gfcc.ContentTypes.EnsureDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_gfcc.ContentTypes.EnsureDefault("\u006a\u0070\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_gfcc.ContentTypes.EnsureDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	_gfcc.ContentTypes.EnsureDefault(i.Format, "\u0069\u006d\u0061\u0067\u0065\u002f"+i.Format)
	return _bgc, nil
}

// Properties returns the properties of the TextBox.
func (_dfeff TextBox) Properties() _g.ShapeProperties {
	if _dfeff._cfbc.SpPr == nil {
		_dfeff._cfbc.SpPr = _dbb.NewCT_ShapeProperties()
	}
	return _g.MakeShapeProperties(_dfeff._cfbc.SpPr)
}

// SetText sets the text of a placeholder for the initial paragraph. This is a
// shortcut method that is useful for things like titles which only contain a
// single paragraph.
func (_dd PlaceHolder) SetText(text string) {
	_dd.Clear()
	_ab := _dbb.NewEG_TextRun()
	_ab.R = _dbb.NewCT_RegularTextRun()
	_ab.R.T = text
	if len(_dd._fb.TxBody.P) == 0 {
		_dd._fb.TxBody.P = append(_dd._fb.TxBody.P, _dbb.NewCT_TextParagraph())
	}
	_dd._fb.TxBody.P[0].EG_TextRun = nil
	_dd._fb.TxBody.P[0].EG_TextRun = append(_dd._fb.TxBody.P[0].EG_TextRun, _ab)
}

// WebPr returns the WebPr property.
func (_aa PresentationProperties) WebPr() *_gb.CT_WebProperties { return _aa._aed.WebPr }

// Index returns the placeholder index
func (_acc PlaceHolder) Index() uint32 {
	if _acc._fb.NvSpPr.NvPr.Ph.IdxAttr == nil {
		return 0
	}
	return *_acc._fb.NvSpPr.NvPr.Ph.IdxAttr
}

// AddParagraph adds a paragraph to the text box
func (_cag TextBox) AddParagraph() _g.Paragraph {
	_fea := _dbb.NewCT_TextParagraph()
	_cag._cfbc.TxBody.P = append(_cag._cfbc.TxBody.P, _fea)
	return _g.MakeParagraph(_fea)
}

// NewViewProperties constructs a new ViewProperties.
func NewViewProperties() ViewProperties { return ViewProperties{_bbab: _gb.NewViewPr()} }

// TextBox is a text box within a slide.
type TextBox struct{ _cfbc *_gb.CT_Shape }

// Slide represents a slide of a presentation.
type Slide struct {
	_gfba *_gb.CT_SlideIdListEntry
	_caed *_gb.Sld
	_dae  *Presentation
}

// Name returns the name of the slide layout.
func (_agaf SlideLayout) Name() string {
	if _agaf._cca.CSld != nil && _agaf._cca.CSld.NameAttr != nil {
		return *_agaf._cca.CSld.NameAttr
	}
	return ""
}

// NewPresentationProperties constructs a new PresentationProperties.
func NewPresentationProperties() PresentationProperties {
	return PresentationProperties{_aed: _gb.NewPresentationPr()}
}

// AddSlide adds a new slide to the presentation.
func (_ed *Presentation) AddSlide() Slide {
	_cbe := _gb.NewCT_SlideIdListEntry()
	_cbe.IdAttr = _ed.nextSlideID()
	_ed._cc.SldIdLst.SldId = append(_ed._cc.SldIdLst.SldId, _cbe)
	_aba := _gb.NewSld()
	_aba.CSld.SpTree.NvGrpSpPr.CNvPr.IdAttr = 1
	_aba.CSld.SpTree.GrpSpPr.Xfrm = _dbb.NewCT_GroupTransform2D()
	_aba.CSld.SpTree.GrpSpPr.Xfrm.Off = _dbb.NewCT_Point2D()
	_aba.CSld.SpTree.GrpSpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _db.Int64(0)
	_aba.CSld.SpTree.GrpSpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _db.Int64(0)
	_aba.CSld.SpTree.GrpSpPr.Xfrm.Ext = _dbb.NewCT_PositiveSize2D()
	_aba.CSld.SpTree.GrpSpPr.Xfrm.Ext.CxAttr = int64(0 * _d.Point)
	_aba.CSld.SpTree.GrpSpPr.Xfrm.Ext.CyAttr = int64(0 * _d.Point)
	_aba.CSld.SpTree.GrpSpPr.Xfrm.ChOff = _aba.CSld.SpTree.GrpSpPr.Xfrm.Off
	_aba.CSld.SpTree.GrpSpPr.Xfrm.ChExt = _aba.CSld.SpTree.GrpSpPr.Xfrm.Ext
	_ed._ddc = append(_ed._ddc, _aba)
	_bbb := _ed._bed.AddAutoRelationship(_db.DocTypePresentation, _db.OfficeDocumentType, len(_ed._ddc), _db.SlideType)
	_cbe.RIdAttr = _bbb.ID()
	_gdf := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideType, len(_ed._ddc))
	_ed.ContentTypes.AddOverride(_gdf, _db.SlideContentType)
	_dfc := _ff.NewRelationships()
	_ed._fba = append(_ed._fba, _dfc)
	_dfc.AddAutoRelationship(_db.DocTypePresentation, _db.SlideType, len(_ed._fbd), _db.SlideLayoutType)
	return Slide{_cbe, _aba, _ed}
}

// SaveToFile writes the Presentation out to a file.
func (_aeg *Presentation) SaveToFile(path string) error { return _aeg.saveToFile(path, false) }

// RemoveSlide removes a slide from a presentation.
func (_gbc *Presentation) RemoveSlide(s Slide) error {
	_fcbec := false
	_bacg := 0
	for _aaf, _dedb := range _gbc._ddc {
		if _dedb == s._caed {
			if _gbc._cc.SldIdLst.SldId[_aaf] != s._gfba {
				return _e.New("i\u006e\u0063\u006f\u006e\u0073\u0069s\u0074\u0065\u006e\u0063\u0079\u0020i\u006e\u0020\u0073\u006c\u0069\u0064\u0065s\u0020\u0061\u006e\u0064\u0020\u0049\u0044\u0020\u006c\u0069s\u0074")
			}
			copy(_gbc._ddc[_aaf:], _gbc._ddc[_aaf+1:])
			_gbc._ddc = _gbc._ddc[0 : len(_gbc._ddc)-1]
			copy(_gbc._fba[_aaf:], _gbc._fba[_aaf+1:])
			_gbc._fba = _gbc._fba[0 : len(_gbc._fba)-1]
			copy(_gbc._cc.SldIdLst.SldId[_aaf:], _gbc._cc.SldIdLst.SldId[_aaf+1:])
			_gbc._cc.SldIdLst.SldId = _gbc._cc.SldIdLst.SldId[0 : len(_gbc._cc.SldIdLst.SldId)-1]
			_fcbec = true
			_bacg = _aaf
		}
	}
	if !_fcbec {
		return _e.New("u\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066i\u006e\u0064\u0020\u0073li\u0064\u0065")
	}
	_fdff := _db.AbsoluteFilename(_db.DocTypePresentation, _db.SlideType, 0)
	return _gbc.ContentTypes.RemoveOverrideByIndex(_fdff, _bacg)
}

// GetPlaceholderByIndex returns a placeholder given its index.  If there are multiplace
// placeholders of the same index, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (_cfae Slide) GetPlaceholderByIndex(idx uint32) (PlaceHolder, error) {
	for _, _ebb := range _cfae._caed.CSld.SpTree.Choice {
		for _, _bgeb := range _ebb.Sp {
			if _bgeb.NvSpPr != nil && _bgeb.NvSpPr.NvPr != nil && _bgeb.NvSpPr.NvPr.Ph != nil {
				if (idx == 0 && _bgeb.NvSpPr.NvPr.Ph.IdxAttr == nil) || (_bgeb.NvSpPr.NvPr.Ph.IdxAttr != nil && *_bgeb.NvSpPr.NvPr.Ph.IdxAttr == idx) {
					return PlaceHolder{_bgeb, _cfae._caed}, nil
				}
			}
		}
	}
	return PlaceHolder{}, _e.New("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066i\u006e\u0064\u0020\u0070\u006c\u0061\u0063\u0065\u0068\u006fl\u0064\u0065\u0072")
}
