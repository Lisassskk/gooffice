package zippkg

import (
	_b "archive/zip"
	_fc "bytes"
	_fg "encoding/xml"
	_ff "fmt"
	_a "io"
	_dc "path"
	_d "sort"
	_cf "strings"
	_f "time"

	_ae "gitee.com/gooffice/gooffice"
	_ge "gitee.com/gooffice/gooffice/algo"
	_ag "gitee.com/gooffice/gooffice/common/tempstorage"
	_g "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships"
)

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor(path string) string {
	_bcg := _cf.Split(path, "\u002f")
	_bga := _cf.Join(_bcg[0:len(_bcg)-1], "\u002f")
	_ad := _bcg[len(_bcg)-1]
	_bga += "\u002f_\u0072\u0065\u006c\u0073\u002f"
	_ad += "\u002e\u0072\u0065l\u0073"
	return _bga + _ad
}
func (_ffg *DecodeMap) IndexFor(path string) int { return _ffg._ffa[path] }

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp(f *_b.File, path string) (string, error) {
	_fb, _bd := _ag.TempFile(path, "\u007a\u007a")
	if _bd != nil {
		return "", _bd
	}
	defer _fb.Close()
	_cd, _bd := f.Open()
	if _bd != nil {
		return "", _bd
	}
	defer _cd.Close()
	_, _bd = _a.Copy(_fb, _cd)
	if _bd != nil {
		return "", _bd
	}
	return _fb.Name(), nil
}

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk(z *_b.Writer, zipPath, storagePath string) error {
	_cca, _aa := z.Create(zipPath)
	if _aa != nil {
		return _ff.Errorf("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073", zipPath, _aa)
	}
	_agd, _aa := _ag.Open(storagePath)
	if _aa != nil {
		return _ff.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", storagePath, _aa)
	}
	defer _agd.Close()
	_, _aa = _a.Copy(_cca, _agd)
	return _aa
}

var _eeg = []byte{'/', '>'}

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func(_cfd *DecodeMap, _e, _ed string, _bg []*_b.File, _da *_g.Relationship, _edf Target) error

const XMLHeader = "\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e" + "\u000a"

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_bf *DecodeMap) AddTarget(filePath string, ifc interface{}, sourceFileType string, idx uint32) bool {
	if _bf._cg == nil {
		_bf._cg = make(map[string]Target)
		_bf._ac = make(map[*_g.Relationships]string)
		_bf._df = make(map[string]struct{})
		_bf._ffa = make(map[string]int)
	}
	_bc := _dc.Clean(filePath)
	if _, _gee := _bf._df[_bc]; _gee {
		return false
	}
	_bf._df[_bc] = struct{}{}
	_bf._cg[_bc] = Target{Path: filePath, Typ: sourceFileType, Ifc: ifc, Index: idx}
	return true
}

type Target struct {
	Path  string
	Typ   string
	Ifc   interface{}
	Index uint32
}

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_fcd *DecodeMap) SetOnNewRelationshipFunc(fn OnNewRelationshipFunc) { _fcd._ga = fn }

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes(z *_b.Writer, zipPath string, data []byte) error {
	_gb, _acf := z.Create(zipPath)
	if _acf != nil {
		return _ff.Errorf("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073", zipPath, _acf)
	}
	_, _acf = _a.Copy(_gb, _fc.NewReader(data))
	return _acf
}

// Decode loops decoding targets registered with AddTarget and calling th
func (_gd *DecodeMap) Decode(files []*_b.File) error {
	_ec := 1
	for _ec > 0 {
		for len(_gd._ee) > 0 {
			_eg := _gd._ee[len(_gd._ee)-1]
			_gd._ee = _gd._ee[0 : len(_gd._ee)-1]
			_gg := _eg.Ifc.(*_g.Relationships)
			for _, _fcf := range _gg.Relationship {
				_gf, _ := _gd._ac[_gg]
				_gd._ga(_gd, _gf+_fcf.TargetAttr, _fcf.TypeAttr, files, _fcf, _eg)
			}
		}
		for _dg, _fee := range files {
			if _fee == nil {
				continue
			}
			if _bca, _gab := _gd._cg[_fee.Name]; _gab {
				delete(_gd._cg, _fee.Name)
				if _be := Decode(_fee, _bca.Ifc); _be != nil {
					return _be
				}
				files[_dg] = nil
				if _dd, _af := _bca.Ifc.(*_g.Relationships); _af {
					_gd._ee = append(_gd._ee, _bca)
					_ffc, _ := _dc.Split(_dc.Clean(_fee.Name + "\u002f\u002e\u002e\u002f"))
					_gd._ac[_dd] = _ffc
					_ec++
				}
			}
		}
		_ec--
	}
	return nil
}

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode(f *_b.File, dest interface{}) error {
	_ba, _ce := f.Open()
	if _ce != nil {
		return _ff.Errorf("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", f.Name, _ce)
	}
	defer _ba.Close()
	_dab := _fg.NewDecoder(_ba)
	if _agb := _dab.Decode(dest); _agb != nil {
		return _ff.Errorf("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073", f.Name, _agb)
	}
	if _ca, _gc := dest.(*_g.Relationships); _gc {
		for _cc, _gea := range _ca.Relationship {
			switch _gea.TypeAttr {
			case _ae.OfficeDocumentTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.OfficeDocumentType
			case _ae.StylesTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.StylesType
			case _ae.ThemeTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.ThemeType
			case _ae.SettingsTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.SettingsType
			case _ae.ImageTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.ImageType
			case _ae.CommentsTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.CommentsType
			case _ae.ThumbnailTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.ThumbnailType
			case _ae.DrawingTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.DrawingType
			case _ae.ChartTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.ChartType
			case _ae.ExtendedPropertiesTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.ExtendedPropertiesType
			case _ae.CustomXMLTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.CustomXMLType
			case _ae.WorksheetTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.WorksheetType
			case _ae.SharedStringsTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.SharedStringsType
			case _ae.TableTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.TableType
			case _ae.HeaderTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.HeaderType
			case _ae.FooterTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.FooterType
			case _ae.NumberingTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.NumberingType
			case _ae.FontTableTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.FontTableType
			case _ae.WebSettingsTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.WebSettingsType
			case _ae.FootNotesTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.FootNotesType
			case _ae.EndNotesTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.EndNotesType
			case _ae.SlideTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.SlideType
			case _ae.VMLDrawingTypeStrict:
				_ca.Relationship[_cc].TypeAttr = _ae.VMLDrawingType
			}
		}
		_d.Slice(_ca.Relationship, func(_bfd, _gcb int) bool {
			_dde := _ca.Relationship[_bfd]
			_ffad := _ca.Relationship[_gcb]
			return _ge.NaturalLess(_dde.IdAttr, _ffad.IdAttr)
		})
	}
	return nil
}

var _dga = []byte{'\r', '\n'}

func (_beea SelfClosingWriter) Write(b []byte) (int, error) {
	_fgd := 0
	_db := 0
	for _ddb := 0; _ddb < len(b)-2; _ddb++ {
		if b[_ddb] == '>' && b[_ddb+1] == '<' && b[_ddb+2] == '/' {
			_ecb := []byte{}
			_gdc := _ddb
			for _fff := _ddb; _fff >= 0; _fff-- {
				if b[_fff] == ' ' {
					_gdc = _fff
				} else if b[_fff] == '<' {
					_ecb = b[_fff+1 : _gdc]
					break
				}
			}
			_fba := []byte{}
			for _adg := _ddb + 3; _adg < len(b); _adg++ {
				if b[_adg] == '>' {
					_fba = b[_ddb+3 : _adg]
					break
				}
			}
			if !_fc.Equal(_ecb, _fba) {
				continue
			}
			_edd, _eb := _beea.W.Write(b[_fgd:_ddb])
			if _eb != nil {
				return _db + _edd, _eb
			}
			_db += _edd
			_, _eb = _beea.W.Write(_eeg)
			if _eb != nil {
				return _db, _eb
			}
			_db += 3
			for _eef := _ddb + 2; _eef < len(b) && b[_eef] != '>'; _eef++ {
				_db++
				_fgd = _eef + 2
				_ddb = _fgd
			}
		}
	}
	_bggg, _ecbb := _beea.W.Write(b[_fgd:])
	return _bggg + _db, _ecbb
}
func MarshalXMLByType(z *_b.Writer, dt _ae.DocType, typ string, v interface{}) error {
	_bee := _ae.AbsoluteFilename(dt, typ, 0)
	return MarshalXML(z, _bee, v)
}
func (_fed *DecodeMap) RecordIndex(path string, idx int) { _fed._ffa[path] = idx }

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{ W _a.Writer }

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML(z *_b.Writer, filename string, v interface{}) error {
	_bab := &_b.FileHeader{}
	_bab.Method = _b.Deflate
	_bab.Name = filename
	_bab.SetModTime(_f.Now())
	_cfg, _ab := z.CreateHeader(_bab)
	if _ab != nil {
		return _ff.Errorf("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073", filename, _ab)
	}
	_, _ab = _cfg.Write([]byte(XMLHeader))
	if _ab != nil {
		return _ff.Errorf("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073", filename, _ab)
	}
	if _ab = _fg.NewEncoder(SelfClosingWriter{_cfg}).Encode(v); _ab != nil {
		return _ff.Errorf("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073", filename, _ab)
	}
	_, _ab = _cfg.Write(_dga)
	return _ab
}
func MarshalXMLByTypeIndex(z *_b.Writer, dt _ae.DocType, typ string, idx int, v interface{}) error {
	_fa := _ae.AbsoluteFilename(dt, typ, idx)
	return MarshalXML(z, _fa, v)
}

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct {
	_cg  map[string]Target
	_ac  map[*_g.Relationships]string
	_ee  []Target
	_ga  OnNewRelationshipFunc
	_df  map[string]struct{}
	_ffa map[string]int
}
