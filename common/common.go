package common

import (
	_da "archive/zip"
	_c "bytes"
	_ge "encoding/xml"
	_e "fmt"
	_fb "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_bf "os"
	_d "reflect"
	_gf "regexp"
	_bfaa "strconv"
	_be "strings"
	_bfa "time"

	_dg "gitee.com/gooffice/gooffice"
	_cd "gitee.com/gooffice/gooffice/common/tempstorage"
	_fba "gitee.com/gooffice/gooffice/common/tempstorage/diskstore"
	_fbd "gitee.com/gooffice/gooffice/measurement"
	_bb "gitee.com/gooffice/gooffice/schema/soo/dml"
	_beb "gitee.com/gooffice/gooffice/schema/soo/ofc/custom_properties"
	_a "gitee.com/gooffice/gooffice/schema/soo/ofc/docPropsVTypes"
	_db "gitee.com/gooffice/gooffice/schema/soo/ofc/extended_properties"
	_ce "gitee.com/gooffice/gooffice/schema/soo/pkg/content_types"
	_gg "gitee.com/gooffice/gooffice/schema/soo/pkg/metadata/core_properties"
	_cb "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships"
	_cde "gitee.com/gooffice/gooffice/zippkg"
)

// Image is a container for image information. It's used as we need format and
// and size information to use images.
// It contains either the filesystem path to the image, or the image itself.
type Image struct {
	Size   _fb.Point
	Format string
	Path   string
	Data   *[]byte
}

// AddImageToZip adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImageToZip(z *_da.Writer, img ImageRef, imageNum int, dt _dg.DocType) error {
	_aad := _dg.AbsoluteImageFilename(dt, imageNum, _be.ToLower(img.Format()))
	if img.Data() != nil && len(*img.Data()) > 0 {
		if _eeba := _cde.AddFileFromBytes(z, _aad, *img.Data()); _eeba != nil {
			return _eeba
		}
	} else if img.Path() != "" {
		if _bead := _cde.AddFileFromDisk(z, _aad, img.Path()); _bead != nil {
			return _bead
		}
	} else {
		return _e.Errorf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u003a %\u002b\u0076", img)
	}
	return nil
}

// RelID returns the relationship ID.
func (_eed ImageRef) RelID() string { return _eed._fef }

// SetOffsetY sets vertical offset of a table in distance units (see measurement package).
func (_gef Table) SetOffsetY(offY float64) {
	if _gef._aac.Off == nil {
		_gef._aac.Off = _bb.NewCT_Point2D()
		_dada := int64(0)
		_gef._aac.Off.XAttr = _bb.ST_Coordinate{ST_CoordinateUnqualified: &_dada}
	}
	_cbga := _fbd.ToEMU(offY)
	_gef._aac.Off.YAttr = _bb.ST_Coordinate{ST_CoordinateUnqualified: &_cbga}
}

// Data returns the data of an image file, if any.
func (_egf ImageRef) Data() *[]byte { return _egf._gfgc.Data }

// Hyperlink is just an appropriately configured relationship.
type Hyperlink Relationship

// Remove removes an existing relationship.
func (_ggbd Relationships) Remove(rel Relationship) bool {
	for _ggga, _cdbd := range _ggbd._bab.Relationship {
		if _cdbd == rel._ccb {
			copy(_ggbd._bab.Relationship[_ggga:], _ggbd._bab.Relationship[_ggga+1:])
			_ggbd._bab.Relationship = _ggbd._bab.Relationship[0 : len(_ggbd._bab.Relationship)-1]
			return true
		}
	}
	return false
}

// ContentTypes is the top level "[Content_Types].xml" in a zip package.
type ContentTypes struct{ _bfaf *_ce.Types }

// CustomProperties contains document specific properties.
type CustomProperties struct{ _ffa *_beb.Properties }

// ImageFromFile reads an image from a file on disk. It doesn't keep the image
// in memory and only reads it to determine the format and size. You can also
// construct an Image directly if the file and size are known.
// NOTE: See also ImageFromStorage.
func ImageFromFile(path string) (Image, error) {
	_cae, _cgc := _bf.Open(path)
	_bgg := Image{}
	if _cgc != nil {
		return _bgg, _e.Errorf("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _cgc)
	}
	defer _cae.Close()
	_fea, _dae, _cgc := _fb.Decode(_cae)
	if _cgc != nil {
		return _bgg, _e.Errorf("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s", _cgc)
	}
	_bgg.Path = path
	_bgg.Format = _dae
	_bgg.Size = _fea.Bounds().Size()
	return _bgg, nil
}

// WriteExtraFiles writes the extra files to the zip package.
func (_dcda *DocBase) WriteExtraFiles(z *_da.Writer) error {
	for _, _efcb := range _dcda.ExtraFiles {
		if _baab := _cde.AddFileFromDisk(z, _efcb.ZipPath, _efcb.DiskPath); _baab != nil {
			return _baab
		}
	}
	return nil
}

// ExtraFile is an unsupported file type extracted from, or to be written to a
// zip package
type ExtraFile struct {
	ZipPath  string
	DiskPath string
}

// NewContentTypes returns a wrapper around a newly constructed content-types.
func NewContentTypes() ContentTypes {
	_fa := ContentTypes{_bfaf: _ce.NewTypes()}
	_fa.AddDefault("\u0078\u006d\u006c", "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c")
	_fa.AddDefault("\u0072\u0065\u006c\u0073", "\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006fn\u002f\u0076\u006ed\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006fr\u006d\u0061\u0074\u0073\u002dpa\u0063\u006b\u0061\u0067\u0065\u002e\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002b\u0078\u006d\u006c")
	_fa.AddDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_fa.AddDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_fa.AddDefault("\u006a\u0070\u0067", "\u0069m\u0061\u0067\u0065\u002f\u006a\u0070g")
	_fa.AddDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	_fa.AddOverride("\u002fd\u006fc\u0050\u0072\u006f\u0070\u0073/\u0063\u006fr\u0065\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073-\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002e\u0063\u006f\u0072\u0065\u002dp\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073\u002bx\u006d\u006c")
	_fa.AddOverride("\u002f\u0064\u006f\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0061\u0070p\u002e\u0078\u006d\u006c", "a\u0070\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u002e\u0065\u0078\u0074\u0065\u006e\u0064\u0065\u0064\u002dp\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002b\u0078m\u006c")
	return _fa
}
func (_degd CustomProperties) SetPropertyAsOstream(name string, ostream string) {
	_dgg := _degd.getNewProperty(name)
	_dgg.Ostream = &ostream
	_degd.setOrReplaceProperty(_dgg)
}
func (_fbdb CustomProperties) setOrReplaceProperty(_bad *_beb.CT_Property) {
	_fbdb.setPropertyHelper(_bad, true)
}

// X returns the underlying raw XML data.
func (_acc Relationships) X() *_cb.Relationships { return _acc._bab }
func (_cea CustomProperties) SetPropertyAsError(name string, error string) {
	_dgad := _cea.getNewProperty(name)
	_dgad.Error = &error
	_cea.setOrReplaceProperty(_dgad)
}

// Theme is a drawingml theme.
type Theme struct{ _dag *_bb.Theme }

// SetContentStatus records the content status of the document.
func (_cg CoreProperties) SetContentStatus(s string) { _cg._bea.ContentStatus = &s }

// SetLanguage records the language of the document.
func (_ggd CoreProperties) SetLanguage(s string) {
	_ggd._bea.Language = &_dg.XSDAny{XMLName: _ge.Name{Local: "d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}}
	_ggd._bea.Language.Data = []byte(s)
}

// PropertiesList returns the list of all custom properties of the document.
func (_gdcc CustomProperties) PropertiesList() []*_beb.CT_Property { return _gdcc._ffa.Property }

// MakeImageRef constructs an image reference which is a reference to a
// particular image file inside a document.  The same image can be used multiple
// times in a document by re-use the ImageRef.
func MakeImageRef(img Image, d *DocBase, rels Relationships) ImageRef {
	return ImageRef{_gfgc: img, _aafg: d, _decf: rels}
}
func _eccg(_fbg _bfa.Time, _fff string) *_dg.XSDAny {
	_edd := &_dg.XSDAny{XMLName: _ge.Name{Local: _fff}}
	_edd.Attrs = append(_edd.Attrs, _ge.Attr{Name: _ge.Name{Local: "\u0078\u0073\u0069\u003a\u0074\u0079\u0070\u0065"}, Value: "\u0064\u0063\u0074\u0065\u0072\u006d\u0073\u003a\u00573\u0043\u0044\u0054\u0046"})
	_edd.Attrs = append(_edd.Attrs, _ge.Attr{Name: _ge.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0073i"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0077\u0077\u0077\u002e\u0077\u0033\u002e\u006f\u0072\u0067\u002f\u00320\u0030\u0031\u002f\u0058\u004d\u004c\u0053\u0063\u0068\u0065\u006d\u0061\u002d\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065"})
	_edd.Attrs = append(_edd.Attrs, _ge.Attr{Name: _ge.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"}, Value: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"})
	_edd.Data = []byte(_fbg.Format(_ac))
	return _edd
}

// X returns the inner raw content types.
func (_ea ContentTypes) X() *_ce.Types { return _ea._bfaf }

// NewTheme constructs a new theme.
func NewTheme() Theme { return Theme{_bb.NewTheme()} }
func (_beaf CustomProperties) SetPropertyAsR4(name string, r4 float32) {
	_baee := _beaf.getNewProperty(name)
	_baee.R4 = &r4
	_beaf.setOrReplaceProperty(_baee)
}

// AddOverride adds an override content type for a given path name.
func (_caff ContentTypes) AddOverride(path, contentType string) {
	if !_be.HasPrefix(path, "\u002f") {
		path = "\u002f" + path
	}
	if _be.HasPrefix(contentType, "\u0068\u0074\u0074\u0070") {
		_dg.Log("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070", contentType)
	}
	for _, _ffg := range _caff._bfaf.Override {
		if _ffg.PartNameAttr == path && _ffg.ContentTypeAttr == contentType {
			return
		}
	}
	_fe := _ce.NewOverride()
	_fe.PartNameAttr = path
	_fe.ContentTypeAttr = contentType
	_caff._bfaf.Override = append(_caff._bfaf.Override, _fe)
}
func (_bae CustomProperties) SetPropertyAsArray(name string, array *_a.Array) {
	_add := _bae.getNewProperty(name)
	_add.Array = array
	_bae.setOrReplaceProperty(_add)
}

// CopyOverride copies override content type for a given `path` and puts it with a path `newPath`.
func (_dd ContentTypes) CopyOverride(path, newPath string) {
	if !_be.HasPrefix(path, "\u002f") {
		path = "\u002f" + path
	}
	if !_be.HasPrefix(newPath, "\u002f") {
		newPath = "\u002f" + newPath
	}
	for _dee := range _dd._bfaf.Override {
		if _dd._bfaf.Override[_dee].PartNameAttr == path {
			_efce := *_dd._bfaf.Override[_dee]
			_efce.PartNameAttr = newPath
			_dd._bfaf.Override = append(_dd._bfaf.Override, &_efce)
		}
	}
}

// SetCategory records the category of the document.
func (_daa CoreProperties) SetCategory(s string) { _daa._bea.Category = &s }
func (_bee CustomProperties) SetPropertyAsLpstr(name string, lpstr string) {
	_cad := _bee.getNewProperty(name)
	_cad.Lpstr = &lpstr
	_bee.setOrReplaceProperty(_cad)
}

// SetLastModifiedBy records the last person to modify the document.
func (_gfg CoreProperties) SetLastModifiedBy(s string) { _gfg._bea.LastModifiedBy = &s }

// TableRow represents a row in a table.
type TableRow struct{ _cag *_bb.CT_TableRow }

// FindRIDForN returns the relationship ID for the i'th relationship of type t.
func (_geeb Relationships) FindRIDForN(i int, t string) string {
	for _, _cdcf := range _geeb._bab.CT_Relationships.Relationship {
		if _cdcf.TypeAttr == t {
			if i == 0 {
				return _cdcf.IdAttr
			}
			i--
		}
	}
	return ""
}

// X returns the inner wrapped XML type of CustomProperty.
func (_bcef CustomProperty) X() *_beb.CT_Property { return _bcef._dbg }

// ImageFromStorage reads an image using the currently set
// temporary storage mechanism (see tempstorage). You can also
// construct an Image directly if the file and size are known.
func ImageFromStorage(path string) (Image, error) {
	_eebaf := Image{}
	_daf, _fce := _cd.Open(path)
	if _fce != nil {
		return _eebaf, _e.Errorf("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _fce)
	}
	defer _daf.Close()
	_bbf, _cfd, _fce := _fb.Decode(_daf)
	if _fce != nil {
		return _eebaf, _e.Errorf("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s", _fce)
	}
	_eebaf.Path = path
	_eebaf.Format = _cfd
	_eebaf.Size = _bbf.Bounds().Size()
	return _eebaf, nil
}
func UtcTimeFormat(t _bfa.Time) string { return t.Format(_dacbd) + "\u0020\u0055\u0054\u0043" }
func (_ddca CustomProperties) SetPropertyAsDate(name string, date _bfa.Time) {
	date = date.UTC()
	_egdb, _acb, _egbd := date.Date()
	_ffad, _baf, _gbc := date.Clock()
	_baa := _bfa.Date(_egdb, _acb, _egbd, _ffad, _baf, _gbc, 0, _bfa.UTC)
	_bdf := _ddca.getNewProperty(name)
	_bdf.Filetime = &_baa
	_ddca.setOrReplaceProperty(_bdf)
}

// SetTarget set the target (path) of a relationship.
func (_ddag Relationship) SetTarget(s string) { _ddag._ccb.TargetAttr = s }

// Rows returns all table rows.
func (_aafgf Table) Rows() []*TableRow {
	_ffga := _aafgf._ccfd.Tr
	_fbaa := []*TableRow{}
	for _, _aabg := range _ffga {
		_fbaa = append(_fbaa, &TableRow{_cag: _aabg})
	}
	return _fbaa
}
func (_geec TableRow) addCell() *_bb.CT_TableCell {
	_cddg := _bb.NewCT_TableCell()
	_geec._cag.Tc = append(_geec._cag.Tc, _cddg)
	return _cddg
}

// CustomProperty contains document specific property.
// Using of this type is deprecated.
type CustomProperty struct{ _dbg *_beb.CT_Property }

func (_gfe CustomProperties) SetPropertyAsStream(name string, stream string) {
	_aba := _gfe.getNewProperty(name)
	_aba.Stream = &stream
	_gfe.setOrReplaceProperty(_aba)
}

// IsEmpty returns true if there are no relationships.
func (_agge Relationships) IsEmpty() bool {
	return _agge._bab == nil || len(_agge._bab.Relationship) == 0
}

// Created returns the time that the document was created.
func (_dffg CoreProperties) Created() _bfa.Time { return _edf(_dffg._bea.Created) }
func (_bfafb CustomProperties) SetPropertyAsI1(name string, i1 int8) {
	_bdd := _bfafb.getNewProperty(name)
	_bdd.I1 = &i1
	_bfafb.setOrReplaceProperty(_bdd)
}

// Relationship is a relationship within a .rels file.
type Relationship struct{ _ccb *_cb.Relationship }

func (_dbac *ImageRef) SetRelID(id string) { _dbac._fef = id }

// Modified returns the time that the document was modified.
func (_dda CoreProperties) Modified() _bfa.Time { return _edf(_dda._bea.Modified) }

// CoreProperties contains document specific properties.
type CoreProperties struct{ _bea *_gg.CoreProperties }

// CreateCustomProperties creates the custom properties of the document.
func (_dfac *DocBase) CreateCustomProperties() {
	_dfac.CustomProperties = NewCustomProperties()
	_dfac.AddCustomRelationships()
}
func (_bcg CustomProperties) getNewProperty(_dbff string) *_beb.CT_Property {
	_edb := _bcg._ffa.Property
	_ad := int32(1)
	for _, _bcc := range _edb {
		if _bcc.PidAttr > _ad {
			_ad = _bcc.PidAttr
		}
	}
	_ffe := _beb.NewCT_Property()
	_ffe.NameAttr = &_dbff
	_ffe.PidAttr = _ad + 1
	_ffe.FmtidAttr = "\u007b\u0044\u0035\u0043\u0044\u0044\u0035\u0030\u0035\u002d\u0032\u0045\u0039\u0043\u002d\u0031\u0030\u0031\u0042\u002d\u0039\u0033\u0039\u0037-\u0030\u0038\u0030\u0030\u0032B\u0032\u0043F\u0039\u0041\u0045\u007d"
	return _ffe
}

// X returns the inner wrapped XML type.
func (_cdg Table) X() *_bb.Tbl { return _cdg._ccfd }

// ContentStatus returns the content status of the document (e.g. "Final", "Draft")
func (_ggg CoreProperties) ContentStatus() string {
	if _ggg._bea.ContentStatus != nil {
		return *_ggg._bea.ContentStatus
	}
	return ""
}

// AddRow adds a row to a table.
func (_efa Table) AddRow() *TableRow {
	_bec := _bb.NewCT_TableRow()
	for _bbgf := 0; _bbgf < len(_efa._ccfd.TblGrid.GridCol); _bbgf++ {
		_bec.Tc = append(_bec.Tc, _bb.NewCT_TableCell())
	}
	_efa._ccfd.Tr = append(_efa._ccfd.Tr, _bec)
	return &TableRow{_cag: _bec}
}

// Target returns the target attrubute of the image reference (a path where the image file is located in the document structure).
func (_gcdg *ImageRef) Target() string { return _gcdg._ecdde }

const _fdf = 2020

func (_bfc CustomProperties) SetPropertyAsOblob(name, oblob string) {
	_egc := _bfc.getNewProperty(name)
	_egc.Oblob = &oblob
	_bfc.setOrReplaceProperty(_egc)
}

const _ac = "2\u00300\u0036\u002d\u0030\u0031\u002d\u0030\u0032\u00541\u0035\u003a\u0030\u0034:0\u0035\u005a"

func (_bed Relationship) String() string {
	return _e.Sprintf("\u007b\u0049\u0044\u003a \u0025\u0073\u0020\u0054\u0061\u0072\u0067\u0065\u0074\u003a \u0025s\u0020\u0054\u0079\u0070\u0065\u003a\u0020%\u0073\u007d", _bed.ID(), _bed.Target(), _bed.Type())
}
func (_fee CustomProperties) SetPropertyAsNull(name string) {
	_ggc := _fee.getNewProperty(name)
	_ggc.Null = _a.NewNull()
	_fee.setOrReplaceProperty(_ggc)
}
func (_aff CustomProperties) setPropertyHelper(_cca *_beb.CT_Property, _gbg bool) {
	_ddc := _aff.GetPropertyByName(*_cca.NameAttr)
	if (_ddc == CustomProperty{}) {
		_aff._ffa.Property = append(_aff._ffa.Property, _cca)
	} else if _gbg {
		_cca.FmtidAttr = _ddc._dbg.FmtidAttr
		if _ddc._dbg.PidAttr == 0 {
			_cca.PidAttr = _ddc._dbg.PidAttr
		}
		_cca.LinkTargetAttr = _ddc._dbg.LinkTargetAttr
		*_ddc._dbg = *_cca
	}
}

// RelativeWidth returns the relative width of an image given a fixed height.
// This is used when setting image to a fixed height to calculate the width
// required to keep the same image aspect ratio.
func (_bff ImageRef) RelativeWidth(h _fbd.Distance) _fbd.Distance {
	_fed := float64(_bff.Size().X) / float64(_bff.Size().Y)
	return h * _fbd.Distance(_fed)
}

// NewTableWithXfrm makes a new table with a pointer to its parent Xfrm for changing its offset and size.
func NewTableWithXfrm(xfrm *_bb.CT_Transform2D) *Table {
	_cab := _bb.NewTbl()
	_cab.TblPr = _bb.NewCT_TableProperties()
	return &Table{_ccfd: _cab, _aac: xfrm}
}

// Size returns the size of an image
func (_dbec ImageRef) Size() _fb.Point { return _dbec._gfgc.Size }

// RemoveOverrideByIndex removes an override given a path and override index.
func (_agb ContentTypes) RemoveOverrideByIndex(path string, indexToFind int) error {
	_abe := path[0 : len(path)-5]
	if !_be.HasPrefix(_abe, "\u002f") {
		_abe = "\u002f" + _abe
	}
	_dea, _ebe := _gf.Compile(_abe + "\u0028\u005b\u0030-\u0039\u005d\u002b\u0029\u002e\u0078\u006d\u006c")
	if _ebe != nil {
		return _ebe
	}
	_afa := 0
	_afaa := -1
	for _afae, _cdd := range _agb._bfaf.Override {
		if _gda := _dea.FindStringSubmatch(_cdd.PartNameAttr); len(_gda) > 1 {
			if _afa == indexToFind {
				_afaa = _afae
			} else if _afa > indexToFind {
				_ebg, _ := _bfaa.Atoi(_gda[1])
				_ebg--
				_cdd.PartNameAttr = _e.Sprintf("\u0025\u0073\u0025\u0064\u002e\u0078\u006d\u006c", _abe, _ebg)
			}
			_afa++
		}
	}
	if _afaa > -1 {
		copy(_agb._bfaf.Override[_afaa:], _agb._bfaf.Override[_afaa+1:])
		_agb._bfaf.Override = _agb._bfaf.Override[0 : len(_agb._bfaf.Override)-1]
	}
	return nil
}

// X returns the inner wrapped XML type.
func (_effb Theme) X() *_bb.Theme { return _effb._dag }

// SetHeight sets row height, see measurement package.
func (_ccc TableRow) SetHeight(m _fbd.Distance) {
	_gce := _fbd.ToEMU(float64(m))
	_ccc._cag.HAttr = _bb.ST_Coordinate{ST_CoordinateUnqualified: &_gce}
}

func (_ecf AppProperties) Pages() int32 {
	if _ecf._eff.Pages != nil {
		return *_ecf._eff.Pages
	}
	return 0
}

// SetWidth sets column width, see measurement package.
func (_bfcg TableCol) SetWidth(m _fbd.Distance) {
	_babg := _fbd.ToEMU(float64(m))
	_bfcg._abca.WAttr = _bb.ST_Coordinate{ST_CoordinateUnqualified: &_babg}
}
func (_fg CustomProperties) setProperty(_cbba *_beb.CT_Property) { _fg.setPropertyHelper(_cbba, false) }

func (_ffc AppProperties) Company() string {
	if _ffc._eff.Company != nil {
		return *_ffc._eff.Company
	}
	return ""
}

// Target returns the target (path) of a relationship.
func (_cgf Relationship) Target() string { return _cgf._ccb.TargetAttr }

// DefAttr returns the DefAttr property.
func (_cbbe TableStyles) DefAttr() string { return _cbbe._ada.DefAttr }

// LastModifiedBy returns the name of the last person to modify the document
func (_fag CoreProperties) LastModifiedBy() string {
	if _fag._bea.LastModifiedBy != nil {
		return *_fag._bea.LastModifiedBy
	}
	return ""
}

// Format returns the format of the underlying image
func (_gdab ImageRef) Format() string { return _gdab._gfgc.Format }

// NewTableStyles constructs a new TableStyles.
func NewTableStyles() TableStyles { return TableStyles{_ada: _bb.NewTblStyleLst()} }
func (_dadb CustomProperties) SetPropertyAsUi1(name string, ui1 uint8) {
	_fabe := _dadb.getNewProperty(name)
	_fabe.Ui1 = &ui1
	_dadb.setOrReplaceProperty(_fabe)
}

// ApplicationVersion returns the version of the application that created the
// document.
func (_dac AppProperties) ApplicationVersion() string {
	if _dac._eff.AppVersion != nil {
		return *_dac._eff.AppVersion
	}
	return ""
}

// X returns the inner wrapped XML type.
func (_aab CustomProperties) X() *_beb.Properties { return _aab._ffa }

// X returns the inner wrapped XML type.
func (_bag CoreProperties) X() *_gg.CoreProperties { return _bag._bea }

// X returns the inner wrapped XML type.
func (_bcd Relationship) X() *_cb.Relationship { return _bcd._ccb }

// SetLinksUpToDate sets the links up to date flag.
func (_ee AppProperties) SetLinksUpToDate(v bool) { _ee._eff.LinksUpToDate = _dg.Bool(v) }

// NewCoreProperties constructs a new CoreProperties.
func NewCoreProperties() CoreProperties { return CoreProperties{_bea: _gg.NewCoreProperties()} }
func (_ecd CustomProperties) SetPropertyAsUi8(name string, ui8 uint64) {
	_ace := _ecd.getNewProperty(name)
	_ace.Ui8 = &ui8
	_ecd.setOrReplaceProperty(_ace)
}

// SetTitle records the title of the document.
func (_geg CoreProperties) SetTitle(s string) {
	if _geg._bea.Title == nil {
		_geg._bea.Title = &_dg.XSDAny{XMLName: _ge.Name{Local: "\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}}
	}
	_geg._bea.Title.Data = []byte(s)
}

// NewRelationshipsCopy creates a new relationships wrapper as a copy of passed in instance.
func NewRelationshipsCopy(rels Relationships) Relationships {
	_bga := *rels._bab
	return Relationships{_bab: &_bga}
}

const Version = "\u0031\u002e\u0037.\u0031"

// NewRelationship constructs a new relationship.
func NewRelationship() Relationship { return Relationship{_ccb: _cb.NewRelationship()} }
func (_cff CustomProperties) SetPropertyAsUi2(name string, ui2 uint16) {
	_dba := _cff.getNewProperty(name)
	_dba.Ui2 = &ui2
	_cff.setOrReplaceProperty(_dba)
}

// SetApplication sets the name of the application that created the document.
func (_caf AppProperties) SetApplication(s string) { _caf._eff.Application = &s }
func (_cbf CustomProperties) SetPropertyAsBlob(name, blob string) {
	_dga := _cbf.getNewProperty(name)
	_dga.Blob = &blob
	_cbf.setOrReplaceProperty(_dga)
}
func _edf(_dgf *_dg.XSDAny) _bfa.Time {
	if _dgf == nil {
		return _bfa.Time{}
	}
	_fab, _cbb := _bfa.Parse(_ac, string(_dgf.Data))
	if _cbb != nil {
		_dg.Log("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0074\u0069\u006d\u0065 \u0066\u0072\u006f\u006d\u0020\u0025\u0073\u003a\u0020\u0025\u0073", string(_dgf.Data), _cbb)
	}
	return _fab
}

// X returns the inner wrapped XML type.
func (_fde AppProperties) X() *_db.Properties { return _fde._eff }
func (_ege CustomProperties) SetPropertyAsBool(name string, b bool) {
	_dbcf := _ege.getNewProperty(name)
	_dbcf.Bool = &b
	_ege.setOrReplaceProperty(_dbcf)
}

// Title returns the Title of the document
func (_fcf CoreProperties) Title() string {
	if _fcf._bea.Title != nil {
		return string(_fcf._bea.Title.Data)
	}
	return ""
}

// ID returns the ID of a relationship.
func (_bdfg Relationship) ID() string { return _bdfg._ccb.IdAttr }

// SetOffsetX sets horizontal offset of a table in distance units (see measurement package).
func (_fbgd Table) SetOffsetX(offX float64) {
	if _fbgd._aac.Off == nil {
		_fbgd._aac.Off = _bb.NewCT_Point2D()
		_gde := int64(0)
		_fbgd._aac.Off.YAttr = _bb.ST_Coordinate{ST_CoordinateUnqualified: &_gde}
	}
	_cgd := _fbd.ToEMU(offX)
	_fbgd._aac.Off.XAttr = _bb.ST_Coordinate{ST_CoordinateUnqualified: &_cgd}
}
func (_aegd CustomProperties) SetPropertyAsEmpty(name string) {
	_gcd := _aegd.getNewProperty(name)
	_gcd.Empty = _a.NewEmpty()
	_aegd.setOrReplaceProperty(_gcd)
}

// Properties returns table properties.
func (_cfc Table) Grid() *_bb.CT_TableGrid { return _cfc._ccfd.TblGrid }

// SetDocSecurity sets the document security flag.
func (_ag AppProperties) SetDocSecurity(v int32) { _ag._eff.DocSecurity = _dg.Int32(v) }
func (_abfb CustomProperties) SetPropertyAsCy(name string, cy string) {
	_gbf := _abfb.getNewProperty(name)
	_gbf.Cy = &cy
	_abfb.setOrReplaceProperty(_gbf)
}

const _agd = 17

// SetCompany sets the name of the company that created the document.
func (_cdef AppProperties) SetCompany(s string) { _cdef._eff.Company = &s }

// TableCol represents a column in a table.
type TableCol struct{ _abca *_bb.CT_TableCol }

// SetCreated sets the time that the document was created.
func (_ddg CoreProperties) SetCreated(t _bfa.Time) {
	_ddg._bea.Created = _eccg(t, "\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064")
}

// NewAppProperties constructs a new AppProperties.
func NewAppProperties() AppProperties {
	_fbdc := AppProperties{_eff: _db.NewProperties()}
	_fbdc.SetCompany("\u0046\u006f\u0078\u0079\u0055\u0074\u0069\u006c\u0073\u0020\u0065\u0068\u0066")
	_fbdc.SetApplication("g\u0069\u0074\u0068\u0075\u0062\u002ec\u006f\u006d\u002f\u0075\u006e\u0069\u0064\u006f\u0063/\u0075\u006e\u0069o\u0066f\u0069\u0063\u0065")
	_fbdc.SetDocSecurity(0)
	_fbdc.SetLinksUpToDate(false)
	var _aea, _ba, _cbc int64
	_e.Sscanf(Version, "\u0025\u0064\u002e\u0025\u0064\u002e\u0025\u0064", &_aea, &_ba, &_cbc)
	_gee := float64(_aea) + float64(_ba)/10000.0
	_fbdc.SetApplicationVersion(_e.Sprintf("\u0025\u0030\u0037\u002e\u0034\u0066", _gee))
	return _fbdc
}
func (_gadg CustomProperties) SetPropertyAsLpwstr(name string, lpwstr string) {
	_caa := _gadg.getNewProperty(name)
	_caa.Lpwstr = &lpwstr
	_gadg.setOrReplaceProperty(_caa)
}

// EnsureOverride ensures that an override for the given path exists, adding it if necessary
func (_gdc ContentTypes) EnsureOverride(path, contentType string) {
	for _, _gb := range _gdc._bfaf.Override {
		if _gb.PartNameAttr == path {
			if _be.HasPrefix(contentType, "\u0068\u0074\u0074\u0070") {
				_dg.Log("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070", contentType)
			}
			_gb.ContentTypeAttr = contentType
			return
		}
	}
	_gdc.AddOverride(path, contentType)
}
func (_cbg CustomProperties) SetPropertyAsOstorage(name string, ostorage string) {
	_aeag := _cbg.getNewProperty(name)
	_aeag.Ostorage = &ostorage
	_cbg.setOrReplaceProperty(_aeag)
}

// EnsureDefault esnures that an extension and default content type exist,
// adding it if necessary.
func (_bg ContentTypes) EnsureDefault(ext, contentType string) {
	ext = _be.ToLower(ext)
	for _, _ffge := range _bg._bfaf.Default {
		if _ffge.ExtensionAttr == ext {
			_ffge.ContentTypeAttr = contentType
			return
		}
	}
	_edc := &_ce.Default{}
	_edc.ContentTypeAttr = contentType
	_edc.ExtensionAttr = ext
	_bg._bfaf.Default = append(_bg._bfaf.Default, _edc)
}
func (_aeg CustomProperties) SetPropertyAsVector(name string, vector *_a.Vector) {
	_dfe := _aeg.getNewProperty(name)
	_dfe.Vector = vector
	_aeg.setOrReplaceProperty(_dfe)
}

const _dacbd = "\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034"

// Properties returns table properties.
func (_bedc Table) Properties() *_bb.CT_TableProperties { return _bedc._ccfd.TblPr }

// Table represents a table in the document.
type Table struct {
	_ccfd *_bb.Tbl
	_aac  *_bb.CT_Transform2D
}

const _dadc = 17

// Relationships represents a .rels file.
type Relationships struct{ _bab *_cb.Relationships }

// ImageFromBytes returns an Image struct for an in-memory image. You can also
// construct an Image directly if the file and size are known.
func ImageFromBytes(data []byte) (Image, error) {
	_cddb := Image{}
	_bffg, _aceg, _caec := _fb.Decode(_c.NewReader(data))
	if _caec != nil {
		return _cddb, _e.Errorf("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s", _caec)
	}
	_cddb.Data = &data
	_cddb.Format = _aceg
	_cddb.Size = _bffg.Bounds().Size()
	return _cddb, nil
}

// Category returns the category of the document
func (_dbf CoreProperties) Category() string {
	if _dbf._bea.Category != nil {
		return *_dbf._bea.Category
	}
	return ""
}

// Append appends DocBase part of an office document to another DocBase.
func (_dbe DocBase) Append(docBase1 DocBase) DocBase {
	_af := docBase1.ContentTypes.X()
	for _, _ff := range _af.Default {
		_dbe.ContentTypes.AddDefault(_ff.ExtensionAttr, _ff.ContentTypeAttr)
	}
	for _, _fbf := range _af.Override {
		_dbe.ContentTypes.AddOverride(_fbf.PartNameAttr, _fbf.ContentTypeAttr)
	}
	_eb := _dbe.AppProperties.X()
	_cc := docBase1.AppProperties.X()
	if _eb.Pages != nil {
		if _cc.Pages != nil {
			*_eb.Pages += *_cc.Pages
		}
	} else if _cc.Pages != nil {
		_eb.Pages = _cc.Pages
	}
	if _eb.Words != nil {
		if _cc.Words != nil {
			*_eb.Words += *_cc.Words
		}
	} else if _cc.Words != nil {
		_eb.Words = _cc.Words
	}
	if _eb.Characters != nil {
		if _cc.Characters != nil {
			*_eb.Characters += *_cc.Characters
		}
	} else if _cc.Characters != nil {
		_eb.Characters = _cc.Characters
	}
	if _eb.Lines != nil {
		if _cc.Lines != nil {
			*_eb.Lines += *_cc.Lines
		}
	} else if _cc.Lines != nil {
		_eb.Lines = _cc.Lines
	}
	if _eb.Paragraphs != nil {
		if _cc.Paragraphs != nil {
			*_eb.Paragraphs += *_cc.Paragraphs
		}
	} else if _cc.Paragraphs != nil {
		_eb.Paragraphs = _cc.Paragraphs
	}
	if _eb.Notes != nil {
		if _cc.Notes != nil {
			*_eb.Notes += *_cc.Notes
		}
	} else if _cc.Notes != nil {
		_eb.Notes = _cc.Notes
	}
	if _eb.HiddenSlides != nil {
		if _cc.HiddenSlides != nil {
			*_eb.HiddenSlides += *_cc.HiddenSlides
		}
	} else if _cc.HiddenSlides != nil {
		_eb.HiddenSlides = _cc.HiddenSlides
	}
	if _eb.MMClips != nil {
		if _cc.MMClips != nil {
			*_eb.MMClips += *_cc.MMClips
		}
	} else if _cc.MMClips != nil {
		_eb.MMClips = _cc.MMClips
	}
	if _eb.LinksUpToDate != nil {
		if _cc.LinksUpToDate != nil {
			*_eb.LinksUpToDate = *_eb.LinksUpToDate && *_cc.LinksUpToDate
		}
	} else if _cc.LinksUpToDate != nil {
		_eb.LinksUpToDate = _cc.LinksUpToDate
	}
	if _eb.CharactersWithSpaces != nil {
		if _cc.CharactersWithSpaces != nil {
			*_eb.CharactersWithSpaces += *_cc.CharactersWithSpaces
		}
	} else if _cc.CharactersWithSpaces != nil {
		_eb.CharactersWithSpaces = _cc.CharactersWithSpaces
	}
	if _eb.SharedDoc != nil {
		if _cc.SharedDoc != nil {
			*_eb.SharedDoc = *_eb.SharedDoc || *_cc.SharedDoc
		}
	} else if _cc.SharedDoc != nil {
		_eb.SharedDoc = _cc.SharedDoc
	}
	if _eb.HyperlinksChanged != nil {
		if _cc.HyperlinksChanged != nil {
			*_eb.HyperlinksChanged = *_eb.HyperlinksChanged || *_cc.HyperlinksChanged
		}
	} else if _cc.HyperlinksChanged != nil {
		_eb.HyperlinksChanged = _cc.HyperlinksChanged
	}
	_eb.DigSig = nil
	if _eb.TitlesOfParts == nil && _cc.TitlesOfParts != nil {
		_eb.TitlesOfParts = _cc.TitlesOfParts
	}
	if _eb.HeadingPairs != nil {
		if _cc.HeadingPairs != nil {
			_ed := _eb.HeadingPairs.Vector
			_fc := _cc.HeadingPairs.Vector
			_cdb := _ed.Variant
			_ec := _fc.Variant
			_bba := []*_a.Variant{}
			for _dc := 0; _dc < len(_ec); _dc += 2 {
				_fd := _ec[_dc].Lpstr
				_ca := false
				for _gfa := 0; _gfa < len(_cdb); _gfa += 2 {
					_ebb := _cdb[_gfa].Lpstr
					if _ebb != nil && _fd != nil && *_ebb == *_fd {
						*_cdb[_gfa+1].I4 = *_cdb[_gfa+1].I4 + *_ec[_dc+1].I4
						_ca = true
						break
					}
				}
				if !_ca {
					_bba = append(_bba, &_a.Variant{CT_Variant: _a.CT_Variant{Lpstr: _ec[_dc].Lpstr}})
					_bba = append(_bba, &_a.Variant{CT_Variant: _a.CT_Variant{I4: _ec[_dc].I4}})
				}
			}
			_cdb = append(_cdb, _bba...)
			_ed.SizeAttr = uint32(len(_cdb))
		}
	} else if _cc.HeadingPairs != nil {
		_eb.HeadingPairs = _cc.HeadingPairs
	}
	if _eb.HLinks != nil {
		if _cc.HLinks != nil {
			_aa := _eb.HLinks.Vector
			_afe := _cc.HLinks.Vector
			_ab := _aa.Variant
			_dcd := _afe.Variant
			for _, _ced := range _dcd {
				_ecc := true
				for _, _gc := range _ab {
					if _d.DeepEqual(_gc, _ced) {
						_ecc = false
						break
					}
				}
				if _ecc {
					_ab = append(_ab, _ced)
					_aa.SizeAttr++
				}
			}
		}
	} else if _cc.HLinks != nil {
		_eb.HLinks = _cc.HLinks
	}
	_bd := _dbe.GetOrCreateCustomProperties()
	_abc := docBase1.GetOrCreateCustomProperties()
	for _, _df := range _abc.PropertiesList() {
		_bd.setProperty(_df)
	}
	_dbe.CustomProperties = _bd
	_bef := _dbe.Rels.X().Relationship
	for _, _eg := range docBase1.Rels.X().Relationship {
		_ef := true
		for _, _gge := range _bef {
			if _gge.TargetAttr == _eg.TargetAttr && _gge.TypeAttr == _eg.TypeAttr {
				_ef = false
				break
			}
		}
		if _ef {
			_dbe.Rels.AddRelationship(_eg.TargetAttr, _eg.TypeAttr)
		}
	}
	for _, _afg := range docBase1.ExtraFiles {
		_efc := _afg.ZipPath
		_dfb := true
		for _, _dff := range _dbe.ExtraFiles {
			if _dff.ZipPath == _efc {
				_dfb = false
				break
			}
		}
		if _dfb {
			_dbe.ExtraFiles = append(_dbe.ExtraFiles, _afg)
		}
	}
	return _dbe
}

// SetModified sets the time that the document was modified.
func (_acg CoreProperties) SetModified(t _bfa.Time) {
	_acg._bea.Modified = _eccg(t, "\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064")
}

// AddHyperlink adds an external hyperlink relationship.
func (_bca Relationships) AddHyperlink(target string) Hyperlink {
	_dacb := _bca.AddRelationship(target, _dg.HyperLinkType)
	_dacb._ccb.TargetModeAttr = _cb.ST_TargetModeExternal
	return Hyperlink(_dacb)
}

// SetTarget changes the target attribute of the image reference (e.g. in the case of the creation of the reference or if the image which the reference is related to was moved from one location to another).
func (_agbb *ImageRef) SetTarget(target string) { _agbb._ecdde = target }

// SetDescription records the description of the document.
func (_cdc CoreProperties) SetDescription(s string) {
	if _cdc._bea.Description == nil {
		_cdc._bea.Description = &_dg.XSDAny{XMLName: _ge.Name{Local: "\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}}
	}
	_cdc._bea.Description.Data = []byte(s)
}

const _cbgc = 2

// SetStyle assigns TableStyle to a table.
func (_gbag Table) SetStyle(style *_bb.CT_TableStyle) {
	if _gbag._ccfd.TblPr == nil {
		_gbag._ccfd.TblPr = _bb.NewCT_TableProperties()
	}
	if _gbag._ccfd.TblPr.Choice == nil {
		_gbag._ccfd.TblPr.Choice = _bb.NewCT_TablePropertiesChoice()
	}
	_gbag._ccfd.TblPr.Choice.TableStyle = style
}
func (_adf CustomProperties) SetPropertyAsUint(name string, ui uint) {
	_bgc := _adf.getNewProperty(name)
	_gec := uint32(ui)
	_bgc.Uint = &_gec
	_adf.setOrReplaceProperty(_bgc)
}

// Relationships returns a slice of all of the relationships.
func (_bdg Relationships) Relationships() []Relationship {
	_bda := []Relationship{}
	for _, _fgd := range _bdg._bab.Relationship {
		_bda = append(_bda, Relationship{_ccb: _fgd})
	}
	return _bda
}

var ReleasedAt = _bfa.Date(_fdf, _fcb, _dadc, _agd, _cbgc, 0, 0, _bfa.UTC)

// AddAutoRelationship adds a relationship with an automatically generated
// filename based off of the type. It should be preferred over AddRelationship
// to ensure consistent filenames are maintained.
func (_fbgg Relationships) AddAutoRelationship(dt _dg.DocType, src string, idx int, ctype string) Relationship {
	return _fbgg.AddRelationship(_dg.RelativeFilename(dt, src, ctype, idx), ctype)
}
func (_deae CustomProperties) SetPropertyAsDecimal(name string, decimal float64) {
	_egd := _deae.getNewProperty(name)
	_egd.Decimal = &decimal
	_deae.setOrReplaceProperty(_egd)
}
func (_eeg CustomProperties) SetPropertyAsInt(name string, i int) {
	_cf := _eeg.getNewProperty(name)
	_dfa := int32(i)
	_cf.Int = &_dfa
	_eeg.setOrReplaceProperty(_cf)
}

// NewCustomProperties constructs a new CustomProperties.
func NewCustomProperties() CustomProperties { return CustomProperties{_ffa: _beb.NewProperties()} }

// Cells returns an array of row cells.
func (_cgef TableRow) Cells() []*_bb.CT_TableCell { return _cgef._cag.Tc }
func (_gaa CustomProperties) SetPropertyAsR8(name string, r8 float64) {
	_cfe := _gaa.getNewProperty(name)
	_cfe.R8 = &r8
	_gaa.setOrReplaceProperty(_cfe)
}

// AddExtraFileFromZip is used when reading an unsupported file from an OOXML
// file. This ensures that unsupported file content will at least round-trip
// correctly.
func (_cga *DocBase) AddExtraFileFromZip(f *_da.File) error {
	_acf, _ede := _cde.ExtractToDiskTmp(f, _cga.TmpPath)
	if _ede != nil {
		return _e.Errorf("\u0065\u0072r\u006f\u0072\u0020\u0065x\u0074\u0072a\u0063\u0074\u0069\u006e\u0067\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073", _ede)
	}
	_cga.ExtraFiles = append(_cga.ExtraFiles, ExtraFile{ZipPath: f.Name, DiskPath: _acf})
	return nil
}
func init() { _fba.SetAsStorage() }

// DocBase is the type embedded in in the Document/Workbook/Presentation types
// that contains members common to all.
type DocBase struct {
	ContentTypes     ContentTypes
	AppProperties    AppProperties
	Rels             Relationships
	CoreProperties   CoreProperties
	CustomProperties CustomProperties
	Thumbnail        _fb.Image
	Images           []ImageRef
	ExtraFiles       []ExtraFile
	TmpPath          string
}

// GetOrCreateCustomProperties returns the custom properties of the document (and if they not exist yet, creating them first).
func (_bac *DocBase) GetOrCreateCustomProperties() CustomProperties {
	if _bac.CustomProperties.X() == nil {
		_bac.CreateCustomProperties()
	}
	return _bac.CustomProperties
}
func (_feb CustomProperties) SetPropertyAsI8(name string, i8 int64) {
	_efg := _feb.getNewProperty(name)
	_efg.I8 = &i8
	_feb.setOrReplaceProperty(_efg)
}
func (_agg CustomProperties) SetPropertyAsFiletime(name string, filetime _bfa.Time) {
	_fae := _agg.getNewProperty(name)
	_fae.Filetime = &filetime
	_agg.setOrReplaceProperty(_fae)
}

// AddDefault registers a default content type for a given file extension.
func (_bbd ContentTypes) AddDefault(fileExtension string, contentType string) {
	fileExtension = _be.ToLower(fileExtension)
	for _, _ffcc := range _bbd._bfaf.Default {
		if _ffcc.ExtensionAttr == fileExtension && _ffcc.ContentTypeAttr == contentType {
			return
		}
	}
	_aeb := _ce.NewDefault()
	_aeb.ExtensionAttr = fileExtension
	_aeb.ContentTypeAttr = contentType
	_bbd._bfaf.Default = append(_bbd._bfaf.Default, _aeb)
}

// SetApplicationVersion sets the version of the application that created the
// document.  Per MS, the verison string mut be in the form 'XX.YYYY'.
func (_ffd AppProperties) SetApplicationVersion(s string) { _ffd._eff.AppVersion = &s }

// TblStyle returns the TblStyle property.
func (_fbgc TableStyles) TblStyle() []*_bb.CT_TableStyle { return _fbgc._ada.TblStyle }
func (_gegc CustomProperties) SetPropertyAsI2(name string, i2 int16) {
	_dge := _gegc.getNewProperty(name)
	_dge.I2 = &i2
	_gegc.setOrReplaceProperty(_dge)
}

// SetAuthor records the author of the document.
func (_egb CoreProperties) SetAuthor(s string) {
	if _egb._bea.Creator == nil {
		_egb._bea.Creator = &_dg.XSDAny{XMLName: _ge.Name{Local: "\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}}
	}
	_egb._bea.Creator.Data = []byte(s)
}

// AddCol adds a column to a table.
func (_eegc Table) AddCol() *TableCol {
	_bedb := _bb.NewCT_TableCol()
	_eegc._ccfd.TblGrid.GridCol = append(_eegc._ccfd.TblGrid.GridCol, _bedb)
	for _, _gba := range _eegc._ccfd.Tr {
		_dab := _bb.NewCT_TableCell()
		_gba.Tc = append(_gba.Tc, _dab)
	}
	return &TableCol{_abca: _bedb}
}

// TableStyles contains document specific properties.
type TableStyles struct{ _ada *_bb.TblStyleLst }

func (_acga CustomProperties) SetPropertyAsBstr(name string, bstr string) {
	_daaf := _acga.getNewProperty(name)
	_daaf.Bstr = &bstr
	_acga.setOrReplaceProperty(_daaf)
}

// AddCustomRelationships adds relationships related to custom properties to the document.
func (_dde *DocBase) AddCustomRelationships() {
	_dde.ContentTypes.AddOverride("/\u0064o\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0063u\u0073\u0074\u006f\u006d.x\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065n\u0078\u006d\u006c\u0066\u006fr\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0063\u0075\u0073\u0074\u006f\u006d\u002d\u0070r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073+\u0078\u006d\u006c")
	_dde.Rels.AddRelationship("\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c", _dg.CustomPropertiesType)
}

// NewTable makes a new table.
func NewTable() *Table {
	_ebf := _bb.NewTbl()
	_ebf.TblPr = _bb.NewCT_TableProperties()
	return &Table{_ccfd: _ebf}
}

// RelativeHeight returns the relative height of an image given a fixed width.
// This is used when setting image to a fixed width to calculate the height
// required to keep the same image aspect ratio.
func (_adb ImageRef) RelativeHeight(w _fbd.Distance) _fbd.Distance {
	_ggb := float64(_adb.Size().Y) / float64(_adb.Size().X)
	return w * _fbd.Distance(_ggb)
}
func (_ddd CustomProperties) SetPropertyAsClsid(name string, clsid string) {
	_bgd := _ddd.getNewProperty(name)
	_bgd.Clsid = &clsid
	_ddd.setOrReplaceProperty(_bgd)
}

// Type returns the type of a relationship.
func (_dfd Relationship) Type() string { return _dfd._ccb.TypeAttr }
func (_fabeg CustomProperties) SetPropertyAsStorage(name string, storage string) {
	_eeb := _fabeg.getNewProperty(name)
	_eeb.Storage = &storage
	_fabeg.setOrReplaceProperty(_eeb)
}

// AddRelationship adds a relationship.
func (_eceb Relationships) AddRelationship(target, ctype string) Relationship {
	if !_be.HasPrefix(ctype, "\u0068t\u0074\u0070\u003a\u002f\u002f") {
		_dg.Log("\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006es\u0068\u0069\u0070 t\u0079\u0070\u0065\u0020\u0025\u0073 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069t\u0068\u0020\u0027\u0068\u0074\u0074\u0070\u003a/\u002f\u0027", ctype)
	}
	_ffgd := _cb.NewRelationship()
	_dfbaa := len(_eceb._bab.Relationship) + 1
	_bgb := map[string]struct{}{}
	for _, _cge := range _eceb._bab.Relationship {
		_bgb[_cge.IdAttr] = struct{}{}
	}
	for _, _efcec := _bgb[_e.Sprintf("\u0072\u0049\u0064%\u0064", _dfbaa)]; _efcec; _, _efcec = _bgb[_e.Sprintf("\u0072\u0049\u0064%\u0064", _dfbaa)] {
		_dfbaa++
	}
	_ffgd.IdAttr = _e.Sprintf("\u0072\u0049\u0064%\u0064", _dfbaa)
	_ffgd.TargetAttr = target
	_ffgd.TypeAttr = ctype
	_eceb._bab.Relationship = append(_eceb._bab.Relationship, _ffgd)
	return Relationship{_ccb: _ffgd}
}
func (_dec CustomProperties) SetPropertyAsI4(name string, i4 int32) {
	_ebbb := _dec.getNewProperty(name)
	_ebbb.I4 = &i4
	_dec.setOrReplaceProperty(_ebbb)
}

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct{ _eff *_db.Properties }

// NewRelationships creates a new relationship wrapper.
func NewRelationships() Relationships { return Relationships{_bab: _cb.NewRelationships()} }

// Clear removes any existing relationships.
func (_eca Relationships) Clear() { _eca._bab.Relationship = nil }

// ImageRef is a reference to an image within a document.
type ImageRef struct {
	_aafg  *DocBase
	_decf  Relationships
	_gfgc  Image
	_fef   string
	_ecdde string
}

// X returns the inner wrapped XML type.
func (_bdb TableStyles) X() *_bb.TblStyleLst { return _bdb._ada }
func (_dfed CustomProperties) SetPropertyAsUi4(name string, ui4 uint32) {
	_afb := _dfed.getNewProperty(name)
	_afb.Ui4 = &ui4
	_dfed.setOrReplaceProperty(_afb)
}

// Author returns the author of the document
func (_dad CoreProperties) Author() string {
	if _dad._bea.Creator != nil {
		return string(_dad._bea.Creator.Data)
	}
	return ""
}

// GetPropertyByName returns a custom property selected by it's name.
func (_abf CustomProperties) GetPropertyByName(name string) CustomProperty {
	_ggf := _abf._ffa.Property
	for _, _dfba := range _ggf {
		if *_dfba.NameAttr == name {
			return CustomProperty{_dbg: _dfba}
		}
	}
	return CustomProperty{}
}
func (_ecdd CustomProperties) SetPropertyAsVstream(name string, vstream *_a.Vstream) {
	_eegb := _ecdd.getNewProperty(name)
	_eegb.Vstream = vstream
	_ecdd.setOrReplaceProperty(_eegb)
}

// Description returns the description of the document
func (_ddf CoreProperties) Description() string {
	if _ddf._bea.Description != nil {
		return string(_ddf._bea.Description.Data)
	}
	return ""
}

const _fcb = 12

// Path returns the path to an image file, if any.
func (_aadf ImageRef) Path() string { return _aadf._gfgc.Path }

// CopyRelationship copies the relationship.
func (_afbb Relationships) CopyRelationship(idAttr string) (Relationship, bool) {
	for _beg := range _afbb._bab.Relationship {
		if _afbb._bab.Relationship[_beg].IdAttr == idAttr {
			_ecdb := *_afbb._bab.Relationship[_beg]
			_cfb := len(_afbb._bab.Relationship) + 1
			_egg := map[string]struct{}{}
			for _, _ega := range _afbb._bab.Relationship {
				_egg[_ega.IdAttr] = struct{}{}
			}
			for _, _dadf := _egg[_e.Sprintf("\u0072\u0049\u0064%\u0064", _cfb)]; _dadf; _, _dadf = _egg[_e.Sprintf("\u0072\u0049\u0064%\u0064", _cfb)] {
				_cfb++
			}
			_ecdb.IdAttr = _e.Sprintf("\u0072\u0049\u0064%\u0064", _cfb)
			_afbb._bab.Relationship = append(_afbb._bab.Relationship, &_ecdb)
			return Relationship{_ccb: &_ecdb}, true
		}
	}
	return Relationship{}, false
}

func (_abg AppProperties) Application() string {
	if _abg._eff.Application != nil {
		return *_abg._eff.Application
	}
	return ""
}

// RemoveOverride removes an override given a path.
func (_bcfb ContentTypes) RemoveOverride(path string) {
	if !_be.HasPrefix(path, "\u002f") {
		path = "\u002f" + path
	}
	for _agc, _ece := range _bcfb._bfaf.Override {
		if _ece.PartNameAttr == path {
			copy(_bcfb._bfaf.Override[_agc:], _bcfb._bfaf.Override[_agc+1:])
			_bcfb._bfaf.Override = _bcfb._bfaf.Override[0 : len(_bcfb._bfaf.Override)-1]
		}
	}
}
