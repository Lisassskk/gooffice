package spreadsheet

import (
	_ab "archive/zip"
	_gb "errors"
	_ge "fmt"
	_a "image"
	_f "image/jpeg"
	_fd "io"
	_fb "log"
	_d "math"
	_ag "math/big"
	_g "os"
	_ba "path/filepath"
	_aa "regexp"
	_b "sort"
	_bd "strconv"
	_ad "strings"
	_fbe "time"

	_c "gitee.com/yu_sheng/gooffice"
	_agg "gitee.com/yu_sheng/gooffice/chart"
	_ae "gitee.com/yu_sheng/gooffice/color"
	_ef "gitee.com/yu_sheng/gooffice/common"
	_da "gitee.com/yu_sheng/gooffice/common/license"
	_cf "gitee.com/yu_sheng/gooffice/common/tempstorage"
	_ee "gitee.com/yu_sheng/gooffice/measurement"
	_dc "gitee.com/yu_sheng/gooffice/schema/soo/dml"
	_gg "gitee.com/yu_sheng/gooffice/schema/soo/dml/chart"
	_df "gitee.com/yu_sheng/gooffice/schema/soo/dml/spreadsheetDrawing"
	_fbf "gitee.com/yu_sheng/gooffice/schema/soo/pkg/relationships"
	_ca "gitee.com/yu_sheng/gooffice/schema/soo/sml"
	_ed "gitee.com/yu_sheng/gooffice/spreadsheet/format"
	_gbf "gitee.com/yu_sheng/gooffice/spreadsheet/formula"
	_db "gitee.com/yu_sheng/gooffice/spreadsheet/reference"
	_aac "gitee.com/yu_sheng/gooffice/spreadsheet/update"
	_fe "gitee.com/yu_sheng/gooffice/vmldrawing"
	_cc "gitee.com/yu_sheng/gooffice/zippkg"
)

// AddFormatValue adds a format value to be used in determining which icons to display.
func (_bafg IconScale) AddFormatValue(t _ca.ST_CfvoType, val string) {
	_cfce := _ca.NewCT_Cfvo()
	_cfce.TypeAttr = t
	_cfce.ValAttr = _c.String(val)
	_bafg._efdg.Cfvo = append(_bafg._efdg.Cfvo, _cfce)
}

var _cece = false

// GetString returns the string in a cell if it's an inline or string table
// string. Otherwise it returns an empty string.
func (_cfb Cell) GetString() string {
	switch _cfb._ec.TAttr {
	case _ca.ST_CellTypeInlineStr:
		if _cfb._ec.Is != nil && _cfb._ec.Is.T != nil {
			return *_cfb._ec.Is.T
		}
		if _cfb._ec.V != nil {
			return *_cfb._ec.V
		}
	case _ca.ST_CellTypeS:
		if _cfb._ec.V == nil {
			return ""
		}
		_cbd, _bec := _bd.Atoi(*_cfb._ec.V)
		if _bec != nil {
			return ""
		}
		_efff, _bec := _cfb._adb.SharedStrings.GetString(_cbd)
		if _bec != nil {
			return ""
		}
		return _efff
	}
	if _cfb._ec.V == nil {
		return ""
	}
	return *_cfb._ec.V
}

// SetDateWithStyle sets a date with the default date style applied.
func (_ceca Cell) SetDateWithStyle(d _fbe.Time) {
	_ceca.SetDate(d)
	for _, _cfa := range _ceca._adb.StyleSheet.CellStyles() {
		if _cfa.HasNumberFormat() && _cfa.NumberFormat() == uint32(StandardFormatDate) {
			_ceca.SetStyle(_cfa)
			return
		}
	}
	_aga := _ceca._adb.StyleSheet.AddCellStyle()
	_aga.SetNumberFormatStandard(StandardFormatDate)
	_ceca.SetStyle(_aga)
}
func (_cccd PatternFill) ClearBgColor() { _cccd._fbdg.BgColor = nil }

// LastRow returns the name of last row which contains data in range of context sheet's given columns.
func (_cfga *evalContext) LastRow(col string) int {
	_gca := _cfga._gag
	_dde := int(_db.ColumnToIndex(col))
	_ggf := 1
	for _, _ceafc := range _gca._fcbf.SheetData.Row {
		if _ceafc.RAttr != nil {
			_bagd := Row{_gca._ccae, _gca, _ceafc}
			_bfe := len(_bagd.Cells())
			if _bfe > _dde {
				_ggf = int(_bagd.RowNumber())
			}
		}
	}
	return _ggf
}

// LastColumn returns the name of last column which contains data in range of context sheet's given rows.
func (_fagg *evalContext) LastColumn(rowFrom, rowTo int) string {
	_ddg := _fagg._gag
	_cbfa := 1
	for _dfd := rowFrom; _dfd <= rowTo; _dfd++ {
		_fgff := len(_ddg.Row(uint32(_dfd)).Cells())
		if _fgff > _cbfa {
			_cbfa = _fgff
		}
	}
	return _db.IndexToColumn(uint32(_cbfa - 1))
}

// SetStringByID sets the cell type to string, and the value a string in the
// shared strings table.
func (_cca Cell) SetStringByID(id int) {
	_cca._adb.ensureSharedStringsRelationships()
	_cca.clearValue()
	_cca._ec.V = _c.String(_bd.Itoa(id))
	_cca._ec.TAttr = _ca.ST_CellTypeS
}

// FormulaContext returns a formula evaluation context that can be used to
// evaluate formaulas.
func (_ded *Sheet) FormulaContext() _gbf.Context { return _afge(_ded) }

// SetColor sets the text color.
func (_dbbcc RichTextRun) SetColor(c _ae.Color) {
	_dbbcc.ensureRpr()
	_dbbcc._bead.RPr.Color = _ca.NewCT_Color()
	_cagg := "\u0066\u0066" + *c.AsRGBString()
	_dbbcc._bead.RPr.Color.RgbAttr = &_cagg
}

// InitialView returns the first defined sheet view. If there are no views, one
// is created and returned.
func (_ebde *Sheet) InitialView() SheetView {
	if _ebde._fcbf.SheetViews == nil || len(_ebde._fcbf.SheetViews.SheetView) == 0 {
		return _ebde.AddView()
	}
	return SheetView{_ebde._fcbf.SheetViews.SheetView[0]}
}

// GetValueAsTime retrieves the cell's value as a time.  There is no difference
// in SpreadsheetML between a time/date cell other than formatting, and that
// typically a date cell won't have a fractional component. GetValueAsTime will
// work for date cells as well.
func (_gee Cell) GetValueAsTime() (_fbe.Time, error) {
	if _gee._ec.TAttr != _ca.ST_CellTypeUnset {
		return _fbe.Time{}, _gb.New("\u0063e\u006c\u006c\u0020\u0074y\u0070\u0065\u0020\u0073\u0068o\u0075l\u0064 \u0062\u0065\u0020\u0075\u006e\u0073\u0065t")
	}
	if _gee._ec.V == nil {
		return _fbe.Time{}, _gb.New("\u0063\u0065\u006c\u006c\u0020\u0068\u0061\u0073\u0020\u006e\u006f\u0020v\u0061\u006c\u0075\u0065")
	}
	_cgeb, _, _gge := _ag.ParseFloat(*_gee._ec.V, 10, 128, _ag.ToNearestEven)
	if _gge != nil {
		return _fbe.Time{}, _gge
	}
	_fge := new(_ag.Float)
	_fge.SetUint64(uint64(24 * _fbe.Hour))
	_cgeb.Mul(_cgeb, _fge)
	_aabg, _ := _cgeb.Uint64()
	_ccd := _gee._adb.Epoch().Add(_fbe.Duration(_aabg))
	return _cec(_ccd), nil
}

// SetError sets the cell type to error and the value to the given error message.
func (_gec Cell) SetError(msg string) {
	_gec.clearValue()
	_gec._ec.V = _c.String(msg)
	_gec._ec.TAttr = _ca.ST_CellTypeE
}

// SetWidthCells is a no-op.
func (_aab AbsoluteAnchor) SetWidthCells(int32) {}

const (
	_caggc = "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061tGe\u006e\u0065\u0072\u0061\u006cS\u0074a\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0057\u0068\u006f\u006ce\u004e\u0075\u006d\u0062\u0065\u0072\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0032\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006da\u0074\u0033\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064F\u006f\u0072\u006d\u0061\u0074\u0034"
	_bceb  = "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074P\u0065\u0072\u0063\u0065\u006e\u0074\u0053\u0074\u0061nd\u0061r\u0064F\u006fr\u006d\u0061\u0074\u0031\u0030\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0031\u0031\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064F\u006f\u0072\u006d\u0061\u0074\u0031\u0032\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0033\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0044\u0061\u0074\u0065\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046o\u0072\u006d\u0061\u0074\u00315\u0053\u0074\u0061\u006e\u0064a\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0036\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0037S\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0038\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0054\u0069\u006d\u0065\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u00320\u0053\u0074a\u006e\u0064a\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0032\u0031\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0044\u0061t\u0065\u0054\u0069\u006d\u0065"
	_ccfd  = "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0033\u0037\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006da\u0074\u0033\u0038\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u00339\u0053\u0074\u0061\u006e\u0064\u0061r\u0064\u0046o\u0072\u006da\u00744\u0030"
	_faea  = "\u0053t\u0061\u006e\u0064a\u0072\u0064\u0046o\u0072ma\u0074\u0034\u0035\u0053\u0074\u0061\u006ed\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0034\u0036\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0034\u0037\u0053ta\u006ed\u0061\u0072\u0064\u0046\u006f\u0072m\u0061\u0074\u0034\u0038\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0034\u0039"
)

// X returns the inner wrapped XML type.
func (_dag DefinedName) X() *_ca.CT_DefinedName { return _dag._gagf }
func (_bcecd Row) renumberAs(_dbfa uint32) {
	_bcecd._gfcg.RAttr = _c.Uint32(_dbfa)
	for _, _cfcdg := range _bcecd.Cells() {
		_dgcc, _bdg := _db.ParseCellReference(_cfcdg.Reference())
		if _bdg == nil {
			_ccca := _ge.Sprintf("\u0025\u0073\u0025\u0064", _dgcc.Column, _dbfa)
			_cfcdg._ec.RAttr = _c.String(_ccca)
		}
	}
}

// RemoveSheetByName removes the sheet with the given name from the workbook.
func (_bagbe *Workbook) RemoveSheetByName(name string) error {
	_gefe := -1
	for _edfd, _cfbb := range _bagbe.Sheets() {
		if name == _cfbb.Name() {
			_gefe = _edfd
			break
		}
	}
	if _gefe == -1 {
		return ErrorNotFound
	}
	return _bagbe.RemoveSheet(_gefe)
}
func (_gabg *Sheet) removeColumnFromNamedRanges(_fegg uint32) error {
	for _, _dga := range _gabg._ccae.DefinedNames() {
		_eegfd := _dga.Name()
		_bddg := _dga.Content()
		_cacb := _ad.Split(_bddg, "\u0021")
		if len(_cacb) != 2 {
			return _gb.New("\u0049\u006e\u0063\u006frr\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0064\u0020\u0072\u0061\u006e\u0067e\u003a" + _bddg)
		}
		_bgfe := _cacb[0]
		if _gabg.Name() == _bgfe {
			_gae := _gabg._ccae.RemoveDefinedName(_dga)
			if _gae != nil {
				return _gae
			}
			_fdgaf := _bdca(_cacb[1], _fegg, true)
			if _fdgaf != "" {
				_fbgef := _bgfe + "\u0021" + _fdgaf
				_gabg._ccae.AddDefinedName(_eegfd, _fbgef)
			}
		}
	}
	_baeg := 0
	if _gabg._fcbf.TableParts != nil && _gabg._fcbf.TableParts.TablePart != nil {
		_baeg = len(_gabg._fcbf.TableParts.TablePart)
	}
	if _baeg != 0 {
		_cacbd := 0
		for _, _babc := range _gabg._ccae.Sheets() {
			if _babc.Name() == _gabg.Name() {
				break
			} else {
				if _babc._fcbf.TableParts != nil && _babc._fcbf.TableParts.TablePart != nil {
					_cacbd += len(_babc._fcbf.TableParts.TablePart)
				}
			}
		}
		_baegd := _gabg._ccae._gfgbd[_cacbd : _cacbd+_baeg]
		for _fcag, _gbeg := range _baegd {
			_bbeb := _gbeg
			_bbeb.RefAttr = _bdca(_bbeb.RefAttr, _fegg, false)
			_gabg._ccae._gfgbd[_cacbd+_fcag] = _bbeb
		}
	}
	return nil
}

// X returns the inner wrapped XML type.
func (_agae DataValidation) X() *_ca.CT_DataValidation { return _agae._cecf }
func (_aeec Font) SetColor(c _ae.Color) {
	_aaae := _ca.NewCT_Color()
	_bae := "\u0066\u0066" + *c.AsRGBString()
	_aaae.RgbAttr = &_bae
	_aeec._acaf.Color = []*_ca.CT_Color{_aaae}
}
func (_afc DataValidation) SetList() DataValidationList {
	_afc.clear()
	_afc._cecf.TypeAttr = _ca.ST_DataValidationTypeList
	_afc._cecf.OperatorAttr = _ca.ST_DataValidationOperatorEqual
	return DataValidationList{_afc._cecf}
}

// SetNumberWithStyle sets a number and applies a standard format to the cell.
func (_gebb Cell) SetNumberWithStyle(v float64, f StandardFormat) {
	_gebb.SetNumber(v)
	_gebb.SetStyle(_gebb._adb.StyleSheet.GetOrCreateStandardNumberFormat(f))
}

// SetValue sets the first value to be used in the comparison.  For comparisons
// that need only one value, this is the only value used.  For comparisons like
// 'between' that require two values, SetValue2 must also be used.
func (_gdge DataValidationCompare) SetValue(v string) { _gdge._cdcg.Formula1 = &v }

// RemoveCalcChain removes the cached caculation chain. This is sometimes needed
// as we don't update it when rows are added/removed.
func (_fdcg *Workbook) RemoveCalcChain() {
	var _dcbc string
	for _, _gbed := range _fdcg._eceef.Relationships() {
		if _gbed.Type() == "ht\u0074\u0070\u003a\u002f\u002f\u0073\u0063he\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006et\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068i\u0070s\u002f\u0063\u0061\u006c\u0063\u0043\u0068\u0061\u0069\u006e" {
			_dcbc = "\u0078\u006c\u002f" + _gbed.Target()
			_fdcg._eceef.Remove(_gbed)
			break
		}
	}
	if _dcbc == "" {
		return
	}
	_fdcg.ContentTypes.RemoveOverride(_dcbc)
	for _aege, _gffd := range _fdcg.ExtraFiles {
		if _gffd.ZipPath == _dcbc {
			_fdcg.ExtraFiles[_aege] = _fdcg.ExtraFiles[len(_fdcg.ExtraFiles)-1]
			_fdcg.ExtraFiles = _fdcg.ExtraFiles[:len(_fdcg.ExtraFiles)-1]
			return
		}
	}
}

// X returns the inner wrapped XML type.
func (_ac Border) X() *_ca.CT_Border { return _ac._eac }

// X returns the inner wrapped XML type.
func (_cbea SheetView) X() *_ca.CT_SheetView { return _cbea._bbgg }

// MakeComments constructs a new Comments wrapper.
func MakeComments(w *Workbook, x *_ca.Comments) Comments { return Comments{w, x} }

// Uses1904Dates returns true if the the workbook uses dates relative to
// 1 Jan 1904. This is uncommon.
func (_cae *Workbook) Uses1904Dates() bool {
	if _cae._geca.WorkbookPr == nil || _cae._geca.WorkbookPr.Date1904Attr == nil {
		return false
	}
	return *_cae._geca.WorkbookPr.Date1904Attr
}

// SetNumber sets the cell type to number, and the value to the given number
func (_gab Cell) SetNumber(v float64) {
	_gab.clearValue()
	if _d.IsNaN(v) || _d.IsInf(v, 0) {
		_gab._ec.TAttr = _ca.ST_CellTypeE
		_gab._ec.V = _c.String("\u0023\u004e\u0055M\u0021")
		return
	}
	_gab._ec.TAttr = _ca.ST_CellTypeN
	_gab._ec.V = _c.String(_bd.FormatFloat(v, 'f', -1, 64))
}

type evalContext struct {
	_gag        *Sheet
	_dbc, _gdbc uint32
	_dce        map[string]struct{}
}

// SetHeightAuto sets the row height to be automatically determined.
func (_ffbff Row) SetHeightAuto()                    { _ffbff._gfcg.HtAttr = nil; _ffbff._gfcg.CustomHeightAttr = nil }
func (_fecb *evalContext) SetOffset(col, row uint32) { _fecb._dbc = col; _fecb._gdbc = row }

const (
	StandardFormatGeneral     StandardFormat = 0
	StandardFormat0           StandardFormat = 0
	StandardFormatWholeNumber StandardFormat = 1
	StandardFormat1           StandardFormat = 1
	StandardFormat2           StandardFormat = 2
	StandardFormat3           StandardFormat = 3
	StandardFormat4           StandardFormat = 4
	StandardFormatPercent     StandardFormat = 9
	StandardFormat9           StandardFormat = 9
	StandardFormat10          StandardFormat = 10
	StandardFormat11          StandardFormat = 11
	StandardFormat12          StandardFormat = 12
	StandardFormat13          StandardFormat = 13
	StandardFormatDate        StandardFormat = 14
	StandardFormat14          StandardFormat = 14
	StandardFormat15          StandardFormat = 15
	StandardFormat16          StandardFormat = 16
	StandardFormat17          StandardFormat = 17
	StandardFormat18          StandardFormat = 18
	StandardFormatTime        StandardFormat = 19
	StandardFormat19          StandardFormat = 19
	StandardFormat20          StandardFormat = 20
	StandardFormat21          StandardFormat = 21
	StandardFormatDateTime    StandardFormat = 22
	StandardFormat22          StandardFormat = 22
	StandardFormat37          StandardFormat = 37
	StandardFormat38          StandardFormat = 38
	StandardFormat39          StandardFormat = 39
	StandardFormat40          StandardFormat = 40
	StandardFormat45          StandardFormat = 45
	StandardFormat46          StandardFormat = 46
	StandardFormat47          StandardFormat = 47
	StandardFormat48          StandardFormat = 48
	StandardFormat49          StandardFormat = 49
)

// RichText is a container for the rich text within a cell. It's similar to a
// paragaraph for a document, except a cell can only contain one rich text item.
type RichText struct{ _egacf *_ca.CT_Rst }

// RemoveSheet removes the sheet with the given index from the workbook.
func (_feea *Workbook) RemoveSheet(ind int) error {
	if _feea.SheetCount() <= ind {
		return ErrorNotFound
	}
	for _, _bacgc := range _feea._eceef.Relationships() {
		if _bacgc.ID() == _feea._geca.Sheets.Sheet[ind].IdAttr {
			_feea._eceef.Remove(_bacgc)
			break
		}
	}
	_feea.ContentTypes.RemoveOverride(_c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.WorksheetContentType, ind+1))
	copy(_feea._abcf[ind:], _feea._abcf[ind+1:])
	_feea._abcf = _feea._abcf[:len(_feea._abcf)-1]
	_gadb := _feea._geca.Sheets.Sheet[ind]
	copy(_feea._geca.Sheets.Sheet[ind:], _feea._geca.Sheets.Sheet[ind+1:])
	_feea._geca.Sheets.Sheet = _feea._geca.Sheets.Sheet[:len(_feea._geca.Sheets.Sheet)-1]
	for _eece := range _feea._geca.Sheets.Sheet {
		if _feea._geca.Sheets.Sheet[_eece].SheetIdAttr > _gadb.SheetIdAttr {
			_feea._geca.Sheets.Sheet[_eece].SheetIdAttr--
		}
	}
	copy(_feea._adfgb[ind:], _feea._adfgb[ind+1:])
	_feea._adfgb = _feea._adfgb[:len(_feea._adfgb)-1]
	copy(_feea._ebaf[ind:], _feea._ebaf[ind+1:])
	_feea._ebaf = _feea._ebaf[:len(_feea._ebaf)-1]
	return nil
}
func (_bbb Cell) getLabelPrefix() string {
	if _bbb._ec.SAttr == nil {
		return ""
	}
	_bbfc := *_bbb._ec.SAttr
	_bea := _bbb._adb.StyleSheet.GetCellStyle(_bbfc)
	switch _bea._bee.Alignment.HorizontalAttr {
	case _ca.ST_HorizontalAlignmentLeft:
		return "\u0027"
	case _ca.ST_HorizontalAlignmentRight:
		return "\u0022"
	case _ca.ST_HorizontalAlignmentCenter:
		return "\u005e"
	case _ca.ST_HorizontalAlignmentFill:
		return "\u005c"
	default:
		return ""
	}
}
func (_fbd Cell) setLocked(_dee bool) {
	_cbc := _fbd._ec.SAttr
	if _cbc != nil {
		_fag := _fbd._adb.StyleSheet.GetCellStyle(*_cbc)
		if _fag._bee.Protection == nil {
			_fag._bee.Protection = _ca.NewCT_CellProtection()
		}
		_fag._bee.Protection.LockedAttr = &_dee
	}
}

// GetLabelPrefix returns label prefix which depends on the cell's horizontal alignment.
func (_bad *evalContext) GetLabelPrefix(cellRef string) string {
	return _bad._gag.Cell(cellRef).getLabelPrefix()
}
func (_gc Border) SetTop(style _ca.ST_BorderStyle, c _ae.Color) {
	if _gc._eac.Top == nil {
		_gc._eac.Top = _ca.NewCT_BorderPr()
	}
	_gc._eac.Top.Color = _ca.NewCT_Color()
	_gc._eac.Top.Color.RgbAttr = c.AsRGBAString()
	_gc._eac.Top.StyleAttr = style
}

// AddCommentWithStyle adds a new comment styled in a default way
func (_aeac Comments) AddCommentWithStyle(cellRef string, author string, comment string) error {
	_cac := _aeac.AddComment(cellRef, author)
	_aee := _cac.AddRun()
	_aee.SetBold(true)
	_aee.SetSize(10)
	_aee.SetColor(_ae.Black)
	_aee.SetFont("\u0043a\u006c\u0069\u0062\u0072\u0069")
	_aee.SetText(author + "\u003a")
	_aee = _cac.AddRun()
	_aee.SetSize(10)
	_aee.SetFont("\u0043a\u006c\u0069\u0062\u0072\u0069")
	_aee.SetColor(_ae.Black)
	_aee.SetText("\u000d\u000a" + comment + "\u000d\u000a")
	_egb, _cbb := _db.ParseCellReference(cellRef)
	if _cbb != nil {
		return _cbb
	}
	_aeac._dfe._gedac[0].Shape = append(_aeac._dfe._gedac[0].Shape, _fe.NewCommentShape(int64(_egb.ColumnIdx), int64(_egb.RowIdx-1)))
	return nil
}

// SetBorder applies a border to a cell style.  The border is referenced by its
// index so modifying the border afterward will affect all styles that reference
// it.
func (_bggd CellStyle) SetBorder(b Border) {
	_bggd._bee.BorderIdAttr = _c.Uint32(b.Index())
	_bggd._bee.ApplyBorderAttr = _c.Bool(true)
}

// Col returns the column of the cell marker.
func (_dabb CellMarker) Col() int32 { return _dabb._ggg.Col }

// SetColOffset sets the column offset of the two cell anchor.
func (_bca TwoCellAnchor) SetColOffset(m _ee.Distance) {
	_dbcd := m - _bca.TopLeft().ColOffset()
	_bca.TopLeft().SetColOffset(m)
	_bca.BottomRight().SetColOffset(_bca.BottomRight().ColOffset() + _dbcd)
}

var ErrorNotFound = _gb.New("\u006eo\u0074\u0020\u0066\u006f\u0075\u006ed")

// AddHyperlink creates and sets a hyperlink on a cell.
func (_abg Cell) AddHyperlink(url string) {
	for _afd, _ege := range _abg._adb._abcf {
		if _ege == _abg._ade._fcbf {
			_abg.SetHyperlink(_abg._adb._adfgb[_afd].AddHyperlink(url))
			return
		}
	}
}
func (_fbeg DifferentialStyle) Fill() Fill {
	if _fbeg._dgea.Fill == nil {
		_fbeg._dgea.Fill = _ca.NewCT_Fill()
	}
	return Fill{_fbeg._dgea.Fill, nil}
}

// Column returns or creates a column that with a given index (1-N).  Columns
// can span multiple column indices, this method will return the column that
// applies to a column index if it exists or create a new column that only
// applies to the index passed in otherwise.
func (_debe *Sheet) Column(idx uint32) Column {
	for _, _ccb := range _debe._fcbf.Cols {
		for _, _cfecc := range _ccb.Col {
			if idx >= _cfecc.MinAttr && idx <= _cfecc.MaxAttr {
				return Column{_cfecc}
			}
		}
	}
	var _cgdf *_ca.CT_Cols
	if len(_debe._fcbf.Cols) == 0 {
		_cgdf = _ca.NewCT_Cols()
		_debe._fcbf.Cols = append(_debe._fcbf.Cols, _cgdf)
	} else {
		_cgdf = _debe._fcbf.Cols[0]
	}
	_efbcg := _ca.NewCT_Col()
	_efbcg.MinAttr = idx
	_efbcg.MaxAttr = idx
	_cgdf.Col = append(_cgdf.Col, _efbcg)
	return Column{_efbcg}
}
func (_gfa Comments) getOrCreateAuthor(_cgfb string) uint32 {
	for _ccf, _dbd := range _gfa._eca.Authors.Author {
		if _dbd == _cgfb {
			return uint32(_ccf)
		}
	}
	_bfg := uint32(len(_gfa._eca.Authors.Author))
	_gfa._eca.Authors.Author = append(_gfa._eca.Authors.Author, _cgfb)
	return _bfg
}

// X returns the inner wrapped XML type.
func (_bacb WorkbookProtection) X() *_ca.CT_WorkbookProtection { return _bacb._eacd }

// SetValues sets the possible values. This is incompatible with SetRange.
func (_ebec DataValidationList) SetValues(values []string) {
	_ebec._ggbf.Formula1 = _c.String("\u0022" + _ad.Join(values, "\u002c") + "\u0022")
	_ebec._ggbf.Formula2 = _c.String("\u0030")
}

// AddSheet adds a new sheet to a workbook.
func (_gaed *Workbook) AddSheet() Sheet {
	_ebea := _ca.NewCT_Sheet()
	_ebea.SheetIdAttr = 1
	for _, _agde := range _gaed._geca.Sheets.Sheet {
		if _ebea.SheetIdAttr <= _agde.SheetIdAttr {
			_ebea.SheetIdAttr = _agde.SheetIdAttr + 1
		}
	}
	_gaed._geca.Sheets.Sheet = append(_gaed._geca.Sheets.Sheet, _ebea)
	_ebea.NameAttr = _ge.Sprintf("\u0053\u0068\u0065\u0065\u0074\u0020\u0025\u0064", _ebea.SheetIdAttr)
	_cdee := _ca.NewWorksheet()
	_cdee.Dimension = _ca.NewCT_SheetDimension()
	_cdee.Dimension.RefAttr = "\u0041\u0031"
	_gaed._abcf = append(_gaed._abcf, _cdee)
	_geab := _ef.NewRelationships()
	_gaed._adfgb = append(_gaed._adfgb, _geab)
	_cdee.SheetData = _ca.NewCT_SheetData()
	_gaed._ebaf = append(_gaed._ebaf, nil)
	_bedb := _c.DocTypeSpreadsheet
	_fbcb := _gaed._eceef.AddAutoRelationship(_bedb, _c.OfficeDocumentType, len(_gaed._geca.Sheets.Sheet), _c.WorksheetType)
	_ebea.IdAttr = _fbcb.ID()
	_gaed.ContentTypes.AddOverride(_c.AbsoluteFilename(_bedb, _c.WorksheetContentType, len(_gaed._geca.Sheets.Sheet)), _c.WorksheetContentType)
	return Sheet{_gaed, _ebea, _cdee}
}

// IsError returns true if the cell is an error type cell.
func (_eae Cell) IsError() bool { return _eae._ec.TAttr == _ca.ST_CellTypeE }

// SetAllowBlank controls if blank values are accepted.
func (_cbeb DataValidation) SetAllowBlank(b bool) {
	if !b {
		_cbeb._cecf.AllowBlankAttr = nil
	} else {
		_cbeb._cecf.AllowBlankAttr = _c.Bool(true)
	}
}

// IsHidden returns whether the row is hidden or not.
func (_dabd Row) IsHidden() bool { return _dabd._gfcg.HiddenAttr != nil && *_dabd._gfcg.HiddenAttr }

// SetOperator sets the operator for the rule.
func (_bdb ConditionalFormattingRule) SetOperator(t _ca.ST_ConditionalFormattingOperator) {
	_bdb._bfgc.OperatorAttr = t
}
func (_bdd PatternFill) X() *_ca.CT_PatternFill { return _bdd._fbdg }
func (_beg Sheet) IsValid() bool                { return _beg._fcbf != nil }
func NewPatternFill(fills *_ca.CT_Fills) PatternFill {
	_ddfb := _ca.NewCT_Fill()
	_ddfb.PatternFill = _ca.NewCT_PatternFill()
	return PatternFill{_ddfb.PatternFill, _ddfb}
}

// Index returns the index of the differential style.
func (_adc DifferentialStyle) Index() uint32 {
	for _dbf, _gebc := range _adc._bcdb.Dxf {
		if _adc._dgea == _gebc {
			return uint32(_dbf)
		}
	}
	return 0
}

// NumberFormat returns the number format that the cell style uses, or zero if
// it is not set.
func (_ebe CellStyle) NumberFormat() uint32 {
	if _ebe._bee.NumFmtIdAttr == nil {
		return 0
	}
	return *_ebe._bee.NumFmtIdAttr
}

// AddConditionalFormatting adds conditional formatting to the sheet.
func (_bacda *Sheet) AddConditionalFormatting(cellRanges []string) ConditionalFormatting {
	_edgd := _ca.NewCT_ConditionalFormatting()
	_bacda._fcbf.ConditionalFormatting = append(_bacda._fcbf.ConditionalFormatting, _edgd)
	_dbfb := make(_ca.ST_Sqref, 0, 0)
	_edgd.SqrefAttr = &_dbfb
	for _, _bgee := range cellRanges {
		*_edgd.SqrefAttr = append(*_edgd.SqrefAttr, _bgee)
	}
	return ConditionalFormatting{_edgd}
}

// X returns the inner wrapped XML type.
func (_egd SheetProtection) X() *_ca.CT_SheetProtection { return _egd._aebb }

// SetActiveSheetIndex sets the index of the active sheet (0-n) which will be
// the tab displayed when the spreadsheet is initially opened.
func (_beddf *Workbook) SetActiveSheetIndex(idx uint32) {
	if _beddf._geca.BookViews == nil {
		_beddf._geca.BookViews = _ca.NewCT_BookViews()
	}
	if len(_beddf._geca.BookViews.WorkbookView) == 0 {
		_beddf._geca.BookViews.WorkbookView = append(_beddf._geca.BookViews.WorkbookView, _ca.NewCT_BookView())
	}
	_beddf._geca.BookViews.WorkbookView[0].ActiveTabAttr = _c.Uint32(idx)
}

// RemoveDefinedName removes an existing defined name.
func (_fadd *Workbook) RemoveDefinedName(dn DefinedName) error {
	if dn.X() == nil {
		return _gb.New("\u0061\u0074\u0074\u0065\u006d\u0070t\u0020\u0074\u006f\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u006e\u0069l\u0020\u0044\u0065\u0066\u0069\u006e\u0065d\u004e\u0061\u006d\u0065")
	}
	for _fagf, _cfdd := range _fadd._geca.DefinedNames.DefinedName {
		if _cfdd == dn.X() {
			copy(_fadd._geca.DefinedNames.DefinedName[_fagf:], _fadd._geca.DefinedNames.DefinedName[_fagf+1:])
			_fadd._geca.DefinedNames.DefinedName[len(_fadd._geca.DefinedNames.DefinedName)-1] = nil
			_fadd._geca.DefinedNames.DefinedName = _fadd._geca.DefinedNames.DefinedName[:len(_fadd._geca.DefinedNames.DefinedName)-1]
			return nil
		}
	}
	return _gb.New("\u0064\u0065\u0066\u0069ne\u0064\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
}

// Comparer is used to compare rows based off a column and cells based off of
// their value.
type Comparer struct{ Order SortOrder }

func _gcd(_egee bool) int {
	if _egee {
		return 1
	}
	return 0
}

// SaveToFile writes the workbook out to a file.
func (_fbee *Workbook) SaveToFile(path string) error {
	_aeeg, _ggga := _g.Create(path)
	if _ggga != nil {
		return _ggga
	}
	defer _aeeg.Close()
	return _fbee.Save(_aeeg)
}

// StandardFormat is a standard ECMA 376 number format.
//go:generate stringer -type=StandardFormat
type StandardFormat uint32

// GetFilename returns the name of file from which workbook was opened with full path to it
func (_cccf *Workbook) GetFilename() string { return _cccf._ffaab }

// X returns the inner wrapped XML type.
func (_egag Comments) X() *_ca.Comments { return _egag._eca }

// BottomRight is a no-op.
func (_dcc AbsoluteAnchor) BottomRight() CellMarker { return CellMarker{} }
func (_ead Border) SetDiagonal(style _ca.ST_BorderStyle, c _ae.Color, up, down bool) {
	if _ead._eac.Diagonal == nil {
		_ead._eac.Diagonal = _ca.NewCT_BorderPr()
	}
	_ead._eac.Diagonal.Color = _ca.NewCT_Color()
	_ead._eac.Diagonal.Color.RgbAttr = c.AsRGBAString()
	_ead._eac.Diagonal.StyleAttr = style
	if up {
		_ead._eac.DiagonalUpAttr = _c.Bool(true)
	}
	if down {
		_ead._eac.DiagonalDownAttr = _c.Bool(true)
	}
}

// Cell retrieves or adds a new cell to a row. Col is the column (e.g. 'A', 'B')
func (_baba Row) Cell(col string) Cell {
	_ffcg := _ge.Sprintf("\u0025\u0073\u0025\u0064", col, _baba.RowNumber())
	for _, _bfgg := range _baba._gfcg.C {
		if _bfgg.RAttr != nil && *_bfgg.RAttr == _ffcg {
			return Cell{_baba._bgb, _baba._dcab, _baba._gfcg, _bfgg}
		}
	}
	return _baba.AddNamedCell(col)
}

// Tables returns a slice of all defined tables in the workbook.
func (_badga *Workbook) Tables() []Table {
	if _badga._gfgbd == nil {
		return nil
	}
	_ggecc := []Table{}
	for _, _acafe := range _badga._gfgbd {
		_ggecc = append(_ggecc, Table{_acafe})
	}
	return _ggecc
}

// SetRange sets the range that contains the possible values. This is incompatible with SetValues.
func (_ecbe DataValidationList) SetRange(cellRange string) {
	_ecbe._ggbf.Formula1 = _c.String(cellRange)
	_ecbe._ggbf.Formula2 = _c.String("\u0030")
}
func (_cee ConditionalFormattingRule) InitializeDefaults() {
	_cee.SetType(_ca.ST_CfTypeCellIs)
	_cee.SetOperator(_ca.ST_ConditionalFormattingOperatorGreaterThan)
	_cee.SetPriority(1)
}

// Epoch returns the point at which the dates/times in the workbook are relative to.
func (_eega *Workbook) Epoch() _fbe.Time {
	if _eega.Uses1904Dates() {
		_fbe.Date(1904, 1, 1, 0, 0, 0, 0, _fbe.UTC)
	}
	return _fbe.Date(1899, 12, 30, 0, 0, 0, 0, _fbe.UTC)
}

// SetRotation configures the cell to be rotated.
func (_bgg CellStyle) SetRotation(deg uint8) {
	if _bgg._bee.Alignment == nil {
		_bgg._bee.Alignment = _ca.NewCT_CellAlignment()
	}
	_bgg._bee.ApplyAlignmentAttr = _c.Bool(true)
	_bgg._bee.Alignment.TextRotationAttr = _c.Uint8(deg)
}

type DifferentialStyle struct {
	_dgea *_ca.CT_Dxf
	_ebdg *Workbook
	_bcdb *_ca.CT_Dxfs
}

// Operator returns the operator for the rule
func (_afg ConditionalFormattingRule) Operator() _ca.ST_ConditionalFormattingOperator {
	return _afg._bfgc.OperatorAttr
}

// SetActiveSheet sets the active sheet which will be the tab displayed when the
// spreadsheet is initially opened.
func (_cdeea *Workbook) SetActiveSheet(s Sheet) {
	for _egfg, _dgfec := range _cdeea._abcf {
		if s._fcbf == _dgfec {
			_cdeea.SetActiveSheetIndex(uint32(_egfg))
		}
	}
}

// ClearAutoFilter removes the autofilters from the sheet.
func (_dade *Sheet) ClearAutoFilter() {
	_dade._fcbf.AutoFilter = nil
	_fdda := "\u0027" + _dade.Name() + "\u0027\u0021"
	for _, _ebdd := range _dade._ccae.DefinedNames() {
		if _ebdd.Name() == _dfdb {
			if _ad.HasPrefix(_ebdd.Content(), _fdda) {
				_dade._ccae.RemoveDefinedName(_ebdd)
				break
			}
		}
	}
}
func NewFills() Fills { return Fills{_ca.NewCT_Fills()} }

// SetDrawing sets the worksheet drawing.  A worksheet can have a reference to a
// single drawing, but the drawing can have many charts.
func (_egfcg *Sheet) SetDrawing(d Drawing) {
	var _fagc _ef.Relationships
	for _gbeb, _aedc := range _egfcg._ccae._abcf {
		if _aedc == _egfcg._fcbf {
			_fagc = _egfcg._ccae._adfgb[_gbeb]
			break
		}
	}
	var _agabg string
	for _cbg, _ffef := range d._ffa._eddg {
		if _ffef == d._acgf {
			_cecg := _fagc.AddAutoRelationship(_c.DocTypeSpreadsheet, _c.WorksheetType, _cbg+1, _c.DrawingType)
			_agabg = _cecg.ID()
			break
		}
	}
	_egfcg._fcbf.Drawing = _ca.NewCT_Drawing()
	_egfcg._fcbf.Drawing.IdAttr = _agabg
}

// SetHeightCells is a no-op.
func (_cfg AbsoluteAnchor) SetHeightCells(int32) {}

// InsertRow inserts a new row into a spreadsheet at a particular row number.  This
// row will now be the row number specified, and any rows after it will be renumbed.
func (_adfg *Sheet) InsertRow(rowNum int) Row {
	_cffdf := uint32(rowNum)
	for _, _efda := range _adfg.Rows() {
		if _efda._gfcg.RAttr != nil && *_efda._gfcg.RAttr >= _cffdf {
			*_efda._gfcg.RAttr++
			for _, _acc := range _efda.Cells() {
				_baca, _bfga := _db.ParseCellReference(_acc.Reference())
				if _bfga != nil {
					continue
				}
				_baca.RowIdx++
				_acc._ec.RAttr = _c.String(_baca.String())
			}
		}
	}
	for _, _dad := range _adfg.MergedCells() {
		_efad, _bga, _afgdd := _db.ParseRangeReference(_dad.Reference())
		if _afgdd != nil {
			continue
		}
		if int(_efad.RowIdx) >= rowNum {
			_efad.RowIdx++
		}
		if int(_bga.RowIdx) >= rowNum {
			_bga.RowIdx++
		}
		_egea := _ge.Sprintf("\u0025\u0073\u003a%\u0073", _efad, _bga)
		_dad.SetReference(_egea)
	}
	return _adfg.AddNumberedRow(_cffdf)
}
func (_ggggb Font) Index() uint32 {
	for _fgcb, _eefb := range _ggggb._aagf.Fonts.Font {
		if _ggggb._acaf == _eefb {
			return uint32(_fgcb)
		}
	}
	return 0
}

// SetHorizontalAlignment sets the horizontal alignment of a cell style.
func (_fcg CellStyle) SetHorizontalAlignment(a _ca.ST_HorizontalAlignment) {
	if _fcg._bee.Alignment == nil {
		_fcg._bee.Alignment = _ca.NewCT_CellAlignment()
	}
	_fcg._bee.Alignment.HorizontalAttr = a
	_fcg._bee.ApplyAlignmentAttr = _c.Bool(true)
}

// SetUnderline controls if the run is underlined.
func (_gfgd RichTextRun) SetUnderline(u _ca.ST_UnderlineValues) {
	_gfgd.ensureRpr()
	_gfgd._bead.RPr.U = _ca.NewCT_UnderlineProperty()
	_gfgd._bead.RPr.U.ValAttr = u
}

// SetFormulaShared sets the cell type to formula shared, and the raw formula to
// the given string. The range is the range of cells that the formula applies
// to, and is used to conserve disk space.
func (_fac Cell) SetFormulaShared(formulaStr string, rows, cols uint32) error {
	_dab := _gbf.ParseString(formulaStr)
	if _dab == nil {
		return _gb.New(_ge.Sprintf("\u0043a\u006en\u006f\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0073", formulaStr))
	}
	_fac.clearValue()
	_fac._ec.TAttr = _ca.ST_CellTypeStr
	_fac._ec.F = _ca.NewCT_CellFormula()
	_fac._ec.F.TAttr = _ca.ST_CellFormulaTypeShared
	_fac._ec.F.Content = formulaStr
	_gba, _aff := _db.ParseCellReference(_fac.Reference())
	if _aff != nil {
		return _aff
	}
	_eg := uint32(0)
	for _, _aad := range _fac._ade.Rows() {
		for _, _bba := range _aad._gfcg.C {
			if _bba.F != nil && _bba.F.SiAttr != nil && *_bba.F.SiAttr >= _eg {
				_eg = *_bba.F.SiAttr
			}
		}
	}
	_eg++
	_efa := _ge.Sprintf("\u0025s\u0025\u0064\u003a\u0025\u0073\u0025d", _gba.Column, _gba.RowIdx, _db.IndexToColumn(_gba.ColumnIdx+cols), _gba.RowIdx+rows)
	_fac._ec.F.RefAttr = _c.String(_efa)
	_fac._ec.F.SiAttr = _c.Uint32(_eg)
	_dcd := Sheet{_fac._adb, _fac._ade._faf, _fac._ade._fcbf}
	for _cd := _gba.RowIdx; _cd <= _gba.RowIdx+rows; _cd++ {
		for _cgf := _gba.ColumnIdx; _cgf <= _gba.ColumnIdx+cols; _cgf++ {
			if _cd == _gba.RowIdx && _cgf == _gba.ColumnIdx {
				continue
			}
			_cfe := _ge.Sprintf("\u0025\u0073\u0025\u0064", _db.IndexToColumn(_cgf), _cd)
			_dcd.Cell(_cfe).Clear()
			_dcd.Cell(_cfe).X().F = _ca.NewCT_CellFormula()
			_dcd.Cell(_cfe).X().F.TAttr = _ca.ST_CellFormulaTypeShared
			_dcd.Cell(_cfe).X().F.SiAttr = _c.Uint32(_eg)
		}
	}
	return nil
}

// SetWidthCells is a no-op.
func (_agab OneCellAnchor) SetWidthCells(int32) {}

// DVCompareOp is a comparison operator for a data validation rule.
type DVCompareOp byte

// SetWidth controls the width of a column.
func (_fbce Column) SetWidth(w _ee.Distance) {
	_fbce._gfd.WidthAttr = _c.Float64(float64(w / _ee.Character))
}

// AddImage adds an image to the workbook package, returning a reference that
// can be used to add the image to a drawing.
func (_ccde *Workbook) AddImage(i _ef.Image) (_ef.ImageRef, error) {
	_eceed := _ef.MakeImageRef(i, &_ccde.DocBase, _ccde._eceef)
	if i.Data == nil && i.Path == "" {
		return _eceed, _gb.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _eceed, _gb.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _eceed, _gb.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	if i.Path != "" {
		_bbgga := _cf.Add(i.Path)
		if _bbgga != nil {
			return _eceed, _bbgga
		}
	}
	_ccde.Images = append(_ccde.Images, _eceed)
	return _eceed, nil
}

// DefinedNames returns a slice of all defined names in the workbook.
func (_cfddb *Workbook) DefinedNames() []DefinedName {
	if _cfddb._geca.DefinedNames == nil {
		return nil
	}
	_afece := []DefinedName{}
	for _, _gfdc := range _cfddb._geca.DefinedNames.DefinedName {
		_afece = append(_afece, DefinedName{_gfdc})
	}
	return _afece
}

// NumberFormat is a number formatting string that can be applied to a cell
// style.
type NumberFormat struct {
	_bbbg *Workbook
	_geeb *_ca.CT_NumFmt
}

func (_efbc CellStyle) Index() uint32 {
	for _bded, _eccf := range _efbc._eda.Xf {
		if _efbc._bee == _eccf {
			return uint32(_bded)
		}
	}
	return 0
}

// SetColOffset sets a column offset in absolute distance.
func (_agd CellMarker) SetColOffset(m _ee.Distance) {
	_agd._ggg.ColOff.ST_CoordinateUnqualified = _c.Int64(int64(m / _ee.EMU))
}

// SetColOffset sets the column offset of the top-left anchor.
func (_egff OneCellAnchor) SetColOffset(m _ee.Distance) { _egff.TopLeft().SetColOffset(m) }

// SetHidden hides or unhides the row
func (_cgae Row) SetHidden(hidden bool) {
	if !hidden {
		_cgae._gfcg.HiddenAttr = nil
	} else {
		_cgae._gfcg.HiddenAttr = _c.Bool(true)
	}
}

// SetItalic causes the text to be displayed in italic.
func (_geeg RichTextRun) SetItalic(b bool) {
	_geeg.ensureRpr()
	_geeg._bead.RPr.I = _ca.NewCT_BooleanProperty()
	_geeg._bead.RPr.I.ValAttr = _c.Bool(b)
}
func (_fdd ConditionalFormattingRule) clear() {
	_fdd._bfgc.OperatorAttr = _ca.ST_ConditionalFormattingOperatorUnset
	_fdd._bfgc.ColorScale = nil
	_fdd._bfgc.IconSet = nil
	_fdd._bfgc.Formula = nil
}

// ConditionalFormatting controls the formatting styles and rules for a range of
// cells with the same conditional formatting.
type ConditionalFormatting struct{ _gde *_ca.CT_ConditionalFormatting }

// SetWidthCells sets the height the anchored object by moving the right hand
// side. It is not compatible with SetWidth.
func (_fffd TwoCellAnchor) SetWidthCells(w int32) {
	_ccgf := _fffd.TopLeft()
	_aegc := _fffd.BottomRight()
	_aegc.SetCol(_ccgf.Col() + w)
}

// SharedStrings is a shared strings table, where string data can be placed
// outside of the sheet contents and referenced from a sheet.
type SharedStrings struct {
	_dfea *_ca.Sst
	_facb map[string]int
}

// ColOffset returns the offset from the row cell.
func (_becc CellMarker) ColOffset() _ee.Distance {
	if _becc._ggg.RowOff.ST_CoordinateUnqualified == nil {
		return 0
	}
	return _ee.Distance(float64(*_becc._ggg.ColOff.ST_CoordinateUnqualified) * _ee.EMU)
}

// GetFormula returns the formula for a cell.
func (_cea Cell) GetFormula() string {
	if _cea._ec.F != nil {
		return _cea._ec.F.Content
	}
	return ""
}

// AbsoluteAnchor has a fixed top-left corner in distance units as well as a
// fixed height/width.
type AbsoluteAnchor struct{ _bdf *_df.CT_AbsoluteAnchor }

// MoveTo repositions the anchor without changing the objects size.
func (_bfff TwoCellAnchor) MoveTo(col, row int32) {
	_dacc := _bfff.TopLeft()
	_bafga := _bfff.BottomRight()
	_degg := _bafga.Col() - _dacc.Col()
	_deaf := _bafga.Row() - _dacc.Row()
	_dacc.SetCol(col)
	_dacc.SetRow(row)
	_bafga.SetCol(col + _degg)
	_bafga.SetRow(row + _deaf)
}

// IsSheetLocked returns whether the sheet objects are locked.
func (_cdf SheetProtection) IsObjectLocked() bool {
	return _cdf._aebb.ObjectsAttr != nil && *_cdf._aebb.ObjectsAttr
}

// SetSize sets the text size for a rich text run.
func (_dfca RichTextRun) SetSize(m _ee.Distance) {
	_dfca.ensureRpr()
	_dfca._bead.RPr.Sz = _ca.NewCT_FontSize()
	_dfca._bead.RPr.Sz.ValAttr = float64(m / _ee.Point)
}

const _ccee = "\u0053\u006fr\u0074\u004f\u0072\u0064e\u0072\u0041s\u0063\u0065\u006e\u0064\u0069\u006e\u0067\u0053o\u0072\u0074\u004f\u0072\u0064\u0065\u0072\u0044\u0065\u0073\u0063\u0065n\u0064\u0069\u006e\u0067"

// SetConditionValue sets the condition value to be used for style applicaton.
func (_bge ConditionalFormattingRule) SetConditionValue(v string) { _bge._bfgc.Formula = []string{v} }
func (_fedc StandardFormat) String() string {
	switch {
	case 0 <= _fedc && _fedc <= 4:
		return _caggc[_fgcda[_fedc]:_fgcda[_fedc+1]]
	case 9 <= _fedc && _fedc <= 22:
		_fedc -= 9
		return _bceb[_bfdd[_fedc]:_bfdd[_fedc+1]]
	case 37 <= _fedc && _fedc <= 40:
		_fedc -= 37
		return _ccfd[_ggbdd[_fedc]:_ggbdd[_fedc+1]]
	case 45 <= _fedc && _fedc <= 49:
		_fedc -= 45
		return _faea[_cgec[_fedc]:_cgec[_fedc+1]]
	default:
		return _ge.Sprintf("\u0053t\u0061n\u0064\u0061\u0072\u0064\u0046o\u0072\u006da\u0074\u0028\u0025\u0064\u0029", _fedc)
	}
}

// GetCachedFormulaResult returns the cached formula result if it exists. If the
// cell type is not a formula cell, the result will be the cell value if it's a
// string/number/bool cell.
func (_cdc Cell) GetCachedFormulaResult() string {
	if _cdc._ec.V != nil {
		return *_cdc._ec.V
	}
	return ""
}

// GetEpoch returns a workbook's time epoch.
func (_gfgf *evalContext) GetEpoch() _fbe.Time { return _gfgf._gag._ccae.Epoch() }
func (_ggbd SheetView) ensurePane() {
	if _ggbd._bbgg.Pane == nil {
		_ggbd._bbgg.Pane = _ca.NewCT_Pane()
		_ggbd._bbgg.Pane.ActivePaneAttr = _ca.ST_PaneBottomLeft
	}
}

// MoveTo is a no-op.
func (_dfa AbsoluteAnchor) MoveTo(x, y int32) {}

// NewSharedStrings constructs a new Shared Strings table.
func NewSharedStrings() SharedStrings {
	return SharedStrings{_dfea: _ca.NewSst(), _facb: make(map[string]int)}
}

// PasswordHash returns the hash of the workbook password.
func (_eeac WorkbookProtection) PasswordHash() string {
	if _eeac._eacd.WorkbookPasswordAttr == nil {
		return ""
	}
	return *_eeac._eacd.WorkbookPasswordAttr
}

type Fills struct{ _gad *_ca.CT_Fills }

// Sheet is a single sheet within a workbook.
type Sheet struct {
	_ccae *Workbook
	_faf  *_ca.CT_Sheet
	_fcbf *_ca.Worksheet
}

// MaxColumnIdx returns the max used column of the sheet.
func (_gafb Sheet) MaxColumnIdx() uint32 {
	_bcb := uint32(0)
	for _, _bfcb := range _gafb.Rows() {
		_cdgc := _bfcb._gfcg.C
		if len(_cdgc) > 0 {
			_bdbcc := _cdgc[len(_cdgc)-1]
			_dfgg, _ := _db.ParseCellReference(*_bdbcc.RAttr)
			if _bcb < _dfgg.ColumnIdx {
				_bcb = _dfgg.ColumnIdx
			}
		}
	}
	return _bcb
}

// ColorScale colors a cell background based off of the cell value.
type ColorScale struct{ _aaa *_ca.CT_ColorScale }

// NewStyleSheet constructs a new default stylesheet.
func NewStyleSheet(wb *Workbook) StyleSheet {
	_beac := _ca.NewStyleSheet()
	_beac.CellStyleXfs = _ca.NewCT_CellStyleXfs()
	_beac.CellXfs = _ca.NewCT_CellXfs()
	_beac.CellStyles = _ca.NewCT_CellStyles()
	_cafda := _ca.NewCT_CellStyle()
	_cafda.NameAttr = _c.String("\u004e\u006f\u0072\u006d\u0061\u006c")
	_cafda.XfIdAttr = 0
	_cafda.BuiltinIdAttr = _c.Uint32(0)
	_beac.CellStyles.CellStyle = append(_beac.CellStyles.CellStyle, _cafda)
	_beac.CellStyles.CountAttr = _c.Uint32(uint32(len(_beac.CellStyles.CellStyle)))
	_cbfe := _ca.NewCT_Xf()
	_cbfe.NumFmtIdAttr = _c.Uint32(0)
	_cbfe.FontIdAttr = _c.Uint32(0)
	_cbfe.FillIdAttr = _c.Uint32(0)
	_cbfe.BorderIdAttr = _c.Uint32(0)
	_beac.CellStyleXfs.Xf = append(_beac.CellStyleXfs.Xf, _cbfe)
	_beac.CellStyleXfs.CountAttr = _c.Uint32(uint32(len(_beac.CellStyleXfs.Xf)))
	_baea := NewFills()
	_beac.Fills = _baea.X()
	_cgcfd := _baea.AddFill().SetPatternFill()
	_cgcfd.SetPattern(_ca.ST_PatternTypeNone)
	_cgcfd = _baea.AddFill().SetPatternFill()
	_cgcfd.SetPattern(_ca.ST_PatternTypeGray125)
	_beac.Fonts = _ca.NewCT_Fonts()
	_beac.Borders = _ca.NewCT_Borders()
	_dbcg := StyleSheet{wb, _beac}
	_dbcg.AddBorder().InitializeDefaults()
	_edcd := _dbcg.AddFont()
	_edcd.SetName("\u0043a\u006c\u0069\u0062\u0072\u0069")
	_edcd.SetSize(11)
	_dcga := _ca.NewCT_Xf()
	*_dcga = *_cbfe
	_dcga.XfIdAttr = _c.Uint32(0)
	_beac.CellXfs.Xf = append(_beac.CellXfs.Xf, _dcga)
	_beac.CellXfs.CountAttr = _c.Uint32(uint32(len(_beac.CellXfs.Xf)))
	return _dbcg
}

// SetColorScale configures the rule as a color scale, removing existing
// configuration.
func (_dccg ConditionalFormattingRule) SetColorScale() ColorScale {
	_dccg.clear()
	_dccg.SetType(_ca.ST_CfTypeColorScale)
	_dccg._bfgc.ColorScale = _ca.NewCT_ColorScale()
	return ColorScale{_dccg._bfgc.ColorScale}
}
func (_bgeg Font) SetName(name string) { _bgeg._acaf.Name = []*_ca.CT_FontName{{ValAttr: name}} }

const (
	SortOrderAscending SortOrder = iota
	SortOrderDescending
)

func _ggae() *_df.CT_OneCellAnchor { _bdcb := _df.NewCT_OneCellAnchor(); return _bdcb }

// Sort sorts all of the rows within a sheet by the contents of a column. As the
// file format doesn't suppot indicating that a column should be sorted by the
// viewing/editing program, we actually need to reorder rows and change cell
// references during a sort. If the sheet contains formulas, you should call
// RecalculateFormulas() prior to sorting.  The column is in the form "C" and
// specifies the column to sort by. The firstRow is a 1-based index and
// specifies the firstRow to include in the sort, allowing skipping over a
// header row.
func (_dfef *Sheet) Sort(column string, firstRow uint32, order SortOrder) {
	_eddb := _dfef._fcbf.SheetData.Row
	_bagb := _dfef.Rows()
	for _aeeca, _gccce := range _bagb {
		if _gccce.RowNumber() == firstRow {
			_eddb = _dfef._fcbf.SheetData.Row[_aeeca:]
			break
		}
	}
	_gabc := Comparer{Order: order}
	_b.Slice(_eddb, func(_eafd, _eede int) bool {
		return _gabc.LessRows(column, Row{_dfef._ccae, _dfef, _eddb[_eafd]}, Row{_dfef._ccae, _dfef, _eddb[_eede]})
	})
	for _abd, _gcge := range _dfef.Rows() {
		_bdfd := uint32(_abd + 1)
		if _gcge.RowNumber() != _bdfd {
			_gcge.renumberAs(_bdfd)
		}
	}
}

// IsDBCS returns if a workbook's default language is among DBCS.
func (_cda *evalContext) IsDBCS() bool {
	_ggba := _cda._gag._ccae.CoreProperties.X().Language
	if _ggba == nil {
		return false
	}
	_acga := string(_ggba.Data)
	for _, _ffb := range _becf {
		if _acga == _ffb {
			return true
		}
	}
	return false
}

// Type returns the type of anchor
func (_ffee TwoCellAnchor) Type() AnchorType { return AnchorTypeTwoCell }
func (_bbccb SortOrder) String() string {
	if _bbccb >= SortOrder(len(_aeebd)-1) {
		return _ge.Sprintf("\u0053\u006f\u0072\u0074\u004f\u0072\u0064\u0065\u0072\u0028\u0025\u0064\u0029", _bbccb)
	}
	return _ccee[_aeebd[_bbccb]:_aeebd[_bbccb+1]]
}

// AddNumberFormat adds a new blank number format to the stylesheet.
func (_cccb StyleSheet) AddNumberFormat() NumberFormat {
	if _cccb._bbbe.NumFmts == nil {
		_cccb._bbbe.NumFmts = _ca.NewCT_NumFmts()
	}
	_egcg := _ca.NewCT_NumFmt()
	_egcg.NumFmtIdAttr = uint32(200 + len(_cccb._bbbe.NumFmts.NumFmt))
	_cccb._bbbe.NumFmts.NumFmt = append(_cccb._bbbe.NumFmts.NumFmt, _egcg)
	_cccb._bbbe.NumFmts.CountAttr = _c.Uint32(uint32(len(_cccb._bbbe.NumFmts.NumFmt)))
	return NumberFormat{_cccb._aba, _egcg}
}

// Column returns the cell column
func (_cdd Cell) Column() (string, error) {
	_ecf, _bda := _db.ParseCellReference(_cdd.Reference())
	if _bda != nil {
		return "", _bda
	}
	return _ecf.Column, nil
}

// SetReference sets the regin of cells that the merged cell applies to.
func (_fbbf MergedCell) SetReference(ref string) { _fbbf._gbb.RefAttr = ref }

// SetFrozen removes any existing sheet views and creates a new single view with
// either the first row, first column or both frozen.
func (_beaa *Sheet) SetFrozen(firstRow, firstCol bool) {
	_beaa._fcbf.SheetViews = nil
	_gedc := _beaa.AddView()
	_gedc.SetState(_ca.ST_PaneStateFrozen)
	switch {
	case firstRow && firstCol:
		_gedc.SetYSplit(1)
		_gedc.SetXSplit(1)
		_gedc.SetTopLeft("\u0042\u0032")
	case firstRow:
		_gedc.SetYSplit(1)
		_gedc.SetTopLeft("\u0041\u0032")
	case firstCol:
		_gedc.SetXSplit(1)
		_gedc.SetTopLeft("\u0042\u0031")
	}
}
func (_dgb *evalContext) NamedRange(ref string) _gbf.Reference {
	for _, _bfb := range _dgb._gag._ccae.DefinedNames() {
		if _bfb.Name() == ref {
			return _gbf.MakeRangeReference(_bfb.Content())
		}
	}
	for _, _eaa := range _dgb._gag._ccae.Tables() {
		if _eaa.Name() == ref {
			return _gbf.MakeRangeReference(_ge.Sprintf("\u0025\u0073\u0021%\u0073", _dgb._gag.Name(), _eaa.Reference()))
		}
	}
	return _gbf.ReferenceInvalid
}

// SetInlineString adds a string inline instead of in the shared strings table.
func (_edg Cell) SetInlineString(s string) {
	_edg.clearValue()
	_edg._ec.Is = _ca.NewCT_Rst()
	_edg._ec.Is.T = _c.String(s)
	_edg._ec.TAttr = _ca.ST_CellTypeInlineStr
}

// X returns the inner wrapped XML type.
func (_fff ConditionalFormattingRule) X() *_ca.CT_CfRule { return _fff._bfgc }
func (_ebda *Sheet) addNumberedRowFast(_fgeg uint32) Row {
	_gccc := _ca.NewCT_Row()
	_gccc.RAttr = _c.Uint32(_fgeg)
	_ebda._fcbf.SheetData.Row = append(_ebda._fcbf.SheetData.Row, _gccc)
	return Row{_ebda._ccae, _ebda, _gccc}
}

// Row will return a row with a given row number, creating a new row if
// necessary.
func (_ggdg *Sheet) Row(rowNum uint32) Row {
	for _, _ddfe := range _ggdg._fcbf.SheetData.Row {
		if _ddfe.RAttr != nil && *_ddfe.RAttr == rowNum {
			return Row{_ggdg._ccae, _ggdg, _ddfe}
		}
	}
	return _ggdg.AddNumberedRow(rowNum)
}
func _dfcg() *_df.CT_AbsoluteAnchor { _gfcd := _df.NewCT_AbsoluteAnchor(); return _gfcd }

// X returns the inner wrapped XML type.
func (_dddb SharedStrings) X() *_ca.Sst { return _dddb._dfea }

// SetAuthor sets the author of the comment. If the comment body contains the
// author's name (as is the case with Excel and Comments.AddCommentWithStyle, it
// will not be changed).  This method only changes the metadata author of the
// comment.
func (_abcd Comment) SetAuthor(author string) {
	_abcd._ggb.AuthorIdAttr = Comments{_abcd._dgcb, _abcd._fca}.getOrCreateAuthor(author)
}

// Reference returns the table reference (the cells within the table)
func (_eaag Table) Reference() string { return _eaag._agbc.RefAttr }

// Name returns the name of the table
func (_cbfec Table) Name() string {
	if _cbfec._agbc.NameAttr != nil {
		return *_cbfec._agbc.NameAttr
	}
	return ""
}

// Anchor is the interface implemented by anchors. It's modeled after the most
// common anchor (Two cell variant with a from/to position), but will also be
// used for one-cell anchors.  In that case the only non-noop methods are
// TopLeft/MoveTo/SetColOffset/SetRowOffset.
type Anchor interface {

	// BottomRight returns the CellMaker for the bottom right corner of the
	// anchor.
	BottomRight() CellMarker

	// TopLeft returns the CellMaker for the top left corner of the anchor.
	TopLeft() CellMarker

	// MoveTo repositions the anchor without changing the objects size.
	MoveTo(_fc, _bdce int32)

	// SetWidth sets the width of the anchored object. It is not compatible with
	// SetWidthCells.
	SetWidth(_bc _ee.Distance)

	// SetWidthCells sets the height the anchored object by moving the right
	// hand side. It is not compatible with SetWidth.
	SetWidthCells(_fcc int32)

	// SetHeight sets the height of the anchored object. It is not compatible
	// with SetHeightCells.
	SetHeight(_eeg _ee.Distance)

	// SetHeightCells sets the height the anchored object by moving the bottom.
	// It is not compatible with SetHeight.
	SetHeightCells(_ff int32)

	// SetColOffset sets the column offset of the top-left anchor.
	SetColOffset(_fa _ee.Distance)

	// SetRowOffset sets the row offset of the top-left anchor.
	SetRowOffset(_bdff _ee.Distance)

	// Type returns the type of anchor
	Type() AnchorType
}

// DataBarScale is a colored scale that fills the cell with a background
// gradeint depending on the value.
type DataBarScale struct{ _egfb *_ca.CT_DataBar }

// LockStructure controls the locking of the workbook structure.
func (_cbfeg WorkbookProtection) LockStructure(b bool) {
	if !b {
		_cbfeg._eacd.LockStructureAttr = nil
	} else {
		_cbfeg._eacd.LockStructureAttr = _c.Bool(true)
	}
}

// SheetView is a view of a sheet. There is typically one per sheet, though more
// are supported.
type SheetView struct{ _bbgg *_ca.CT_SheetView }

// LessRows compares two rows based off of a column. If the column doesn't exist
// in one row, that row is 'less'.
func (_ccc Comparer) LessRows(column string, lhs, rhs Row) bool {
	var _eccc, _gfad Cell
	for _, _aggf := range lhs.Cells() {
		_dfgbc, _ := _db.ParseCellReference(_aggf.Reference())
		if _dfgbc.Column == column {
			_eccc = _aggf
			break
		}
	}
	for _, _aed := range rhs.Cells() {
		_ccdb, _ := _db.ParseCellReference(_aed.Reference())
		if _ccdb.Column == column {
			_gfad = _aed
			break
		}
	}
	return _ccc.LessCells(_eccc, _gfad)
}

// SetWrapped configures the cell to wrap text.
func (_cbe CellStyle) SetWrapped(b bool) {
	if _cbe._bee.Alignment == nil {
		_cbe._bee.Alignment = _ca.NewCT_CellAlignment()
	}
	if !b {
		_cbe._bee.Alignment.WrapTextAttr = nil
	} else {
		_cbe._bee.Alignment.WrapTextAttr = _c.Bool(true)
		_cbe._bee.ApplyAlignmentAttr = _c.Bool(true)
	}
}

// SetPriority sets the rule priority
func (_cddc ConditionalFormattingRule) SetPriority(p int32) { _cddc._bfgc.PriorityAttr = p }

var _aeebd = [...]uint8{0, 18, 37}

// SetLocked sets cell locked or not.
func (_cdg *evalContext) SetLocked(cellRef string, locked bool) {
	_cdg._gag.Cell(cellRef).setLocked(locked)
}

// SetRowOffset sets the row offset of the top-left of the image in fixed units.
func (_ea AbsoluteAnchor) SetRowOffset(m _ee.Distance) {
	_ea._bdf.Pos.YAttr.ST_CoordinateUnqualified = _c.Int64(int64(m / _ee.EMU))
}
func (_ebed Fill) Index() uint32 {
	if _ebed._ada == nil {
		return 0
	}
	for _bdege, _cdde := range _ebed._ada.Fill {
		if _ebed._ffad == _cdde {
			return uint32(_bdege)
		}
	}
	return 0
}

// SetStyle applies a style to the cell.  This style is referenced in the
// generated XML via CellStyle.Index().
func (_fbg Cell) SetStyle(cs CellStyle) { _fbg.SetStyleIndex(cs.Index()) }
func (_ga Cell) clearValue() {
	_ga._ec.F = nil
	_ga._ec.Is = nil
	_ga._ec.V = nil
	_ga._ec.TAttr = _ca.ST_CellTypeUnset
}

// X returns the inner wrapped XML type.
func (_gggg Comment) X() *_ca.CT_Comment { return _gggg._ggb }

// RangeReference converts a range reference of the form 'A1:A5' to 'Sheet
// 1'!$A$1:$A$5 . Renaming a sheet after calculating a range reference will
// invalidate the reference.
func (_gbcc Sheet) RangeReference(n string) string {
	_dbfd := _ad.Split(n, "\u003a")
	_bfge, _ := _db.ParseCellReference(_dbfd[0])
	_cagd := _ge.Sprintf("\u0024\u0025\u0073\u0024\u0025\u0064", _bfge.Column, _bfge.RowIdx)
	if len(_dbfd) == 1 {
		return _ge.Sprintf("\u0027%\u0073\u0027\u0021\u0025\u0073", _gbcc.Name(), _cagd)
	}
	_abeb, _ := _db.ParseCellReference(_dbfd[1])
	_efcf := _ge.Sprintf("\u0024\u0025\u0073\u0024\u0025\u0064", _abeb.Column, _abeb.RowIdx)
	return _ge.Sprintf("\u0027\u0025\u0073\u0027\u0021\u0025\u0073\u003a\u0025\u0073", _gbcc.Name(), _cagd, _efcf)
}

// SetXSplit sets the column split point
func (_aggg SheetView) SetXSplit(v float64) {
	_aggg.ensurePane()
	_aggg._bbgg.Pane.XSplitAttr = _c.Float64(v)
}

// AddFormatValue adds a format value (databars require two).
func (_efbg DataBarScale) AddFormatValue(t _ca.ST_CfvoType, val string) {
	_bdfe := _ca.NewCT_Cfvo()
	_bdfe.TypeAttr = t
	_bdfe.ValAttr = _c.String(val)
	_efbg._egfb.Cfvo = append(_efbg._egfb.Cfvo, _bdfe)
}

// SetFont sets the font name for a rich text run.
func (_bgdc RichTextRun) SetFont(s string) {
	_bgdc.ensureRpr()
	_bgdc._bead.RPr.RFont = _ca.NewCT_FontName()
	_bgdc._bead.RPr.RFont.ValAttr = s
}
func (_fbfe DataValidationCompare) SetValue2(v string) { _fbfe._cdcg.Formula2 = &v }

// PasswordHash returns the hash of the workbook password.
func (_cacg SheetProtection) PasswordHash() string {
	if _cacg._aebb.PasswordAttr == nil {
		return ""
	}
	return *_cacg._aebb.PasswordAttr
}

// GetFilename returns the filename of the context's workbook.
func (_bbga *evalContext) GetFilename() string { return _bbga._gag._ccae.GetFilename() }
func (_cgcf Font) SetBold(b bool) {
	if b {
		_cgcf._acaf.B = []*_ca.CT_BooleanProperty{{}}
	} else {
		_cgcf._acaf.B = nil
	}
}

// CellsWithEmpty returns a slice of cells including empty ones from the first column to the last one used in the sheet.
// The cells can be manipulated, but appending to the slice will have no effect.
func (_ecdf Row) CellsWithEmpty(lastColIdx uint32) []Cell {
	_ggda := []Cell{}
	for _dfcab := uint32(0); _dfcab <= lastColIdx; _dfcab++ {
		_cfec := _ecdf.Cell(_db.IndexToColumn(_dfcab))
		_ggda = append(_ggda, _cfec)
	}
	return _ggda
}

// GetFormat returns a cell data format.
func (_gcg *evalContext) GetFormat(cellRef string) string { return _gcg._gag.Cell(cellRef).getFormat() }

// SetFormat sets the number format code.
func (_dbe NumberFormat) SetFormat(f string) { _dbe._geeb.FormatCodeAttr = f }

// AddDefinedName adds a name for a cell or range reference that can be used in
// formulas and charts.
func (_gfbcf *Workbook) AddDefinedName(name, ref string) DefinedName {
	if _gfbcf._geca.DefinedNames == nil {
		_gfbcf._geca.DefinedNames = _ca.NewCT_DefinedNames()
	}
	_ddga := _ca.NewCT_DefinedName()
	_ddga.Content = ref
	_ddga.NameAttr = name
	_gfbcf._geca.DefinedNames.DefinedName = append(_gfbcf._geca.DefinedNames.DefinedName, _ddga)
	return DefinedName{_ddga}
}
func (_cab CellStyle) SetShrinkToFit(b bool) {
	if _cab._bee.Alignment == nil {
		_cab._bee.Alignment = _ca.NewCT_CellAlignment()
	}
	_cab._bee.ApplyAlignmentAttr = _c.Bool(true)
	if !b {
		_cab._bee.Alignment.ShrinkToFitAttr = nil
	} else {
		_cab._bee.Alignment.ShrinkToFitAttr = _c.Bool(b)
	}
}

// Font allows editing fonts within a spreadsheet stylesheet.
type Font struct {
	_acaf *_ca.CT_Font
	_aagf *_ca.StyleSheet
}

// SetTime sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel).
func (_eb Cell) SetTime(d _fbe.Time) {
	_eb.clearValue()
	d = _cgc(d)
	_bbd := _eb._adb.Epoch()
	if d.Before(_bbd) {
		_c.Log("t\u0069\u006d\u0065\u0073\u0020\u0062e\u0066\u006f\u0072\u0065\u0020\u00319\u0030\u0030\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064")
		return
	}
	_cbf := d.Sub(_bbd)
	_egc := new(_ag.Float)
	_fgf := new(_ag.Float)
	_fgf.SetPrec(128)
	_fgf.SetUint64(uint64(_cbf))
	_dfgb := new(_ag.Float)
	_dfgb.SetUint64(24 * 60 * 60 * 1e9)
	_egc.Quo(_fgf, _dfgb)
	_eb._ec.V = _c.String(_egc.Text('g', 20))
}

// DVCompareType is a comparison type for a data validation rule. This restricts
// the input format of the cell.
type DVCompareType byte

// X returns the inner wrapped XML type.
func (_efd CellMarker) X() *_df.CT_Marker { return _efd._ggg }

// SetHyperlink sets a hyperlink on a cell.
func (_affd Cell) SetHyperlink(hl _ef.Hyperlink) {
	_ecc := _affd._ade._fcbf
	if _ecc.Hyperlinks == nil {
		_ecc.Hyperlinks = _ca.NewCT_Hyperlinks()
	}
	_fgca := _ef.Relationship(hl)
	_bage := _ca.NewCT_Hyperlink()
	_bage.RefAttr = _affd.Reference()
	_bage.IdAttr = _c.String(_fgca.ID())
	_ecc.Hyperlinks.Hyperlink = append(_ecc.Hyperlinks.Hyperlink, _bage)
}

// Protection allows control over the workbook protections.
func (_addbf *Workbook) Protection() WorkbookProtection {
	if _addbf._geca.WorkbookProtection == nil {
		_addbf._geca.WorkbookProtection = _ca.NewCT_WorkbookProtection()
	}
	return WorkbookProtection{_addbf._geca.WorkbookProtection}
}

// Validate attempts to validate the structure of a workbook.
func (_fdce *Workbook) Validate() error {
	if _fdce == nil || _fdce._geca == nil {
		return _gb.New("\u0077o\u0072\u006bb\u006f\u006f\u006b\u0020n\u006f\u0074\u0020i\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0065d \u0063\u006f\u0072r\u0065\u0063t\u006c\u0079\u002c\u0020\u006e\u0069l\u0020\u0062a\u0073\u0065")
	}
	_fcge := uint32(0)
	for _, _gbbe := range _fdce._geca.Sheets.Sheet {
		if _gbbe.SheetIdAttr > _fcge {
			_fcge = _gbbe.SheetIdAttr
		}
	}
	if _fcge != uint32(len(_fdce._abcf)) {
		return _ge.Errorf("\u0066\u006f\u0075\u006e\u0064\u0020%\u0064\u0020\u0077\u006f\u0072\u006b\u0073\u0068\u0065\u0065\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074i\u006f\u006e\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0064\u0020\u0077\u006f\u0072k\u0073h\u0065\u0065\u0074\u0073", _fcge, len(_fdce._abcf))
	}
	_edgb := map[string]struct{}{}
	for _aedbg, _gdac := range _fdce._geca.Sheets.Sheet {
		_bced := Sheet{_fdce, _gdac, _fdce._abcf[_aedbg]}
		if _, _bcgc := _edgb[_bced.Name()]; _bcgc {
			return _ge.Errorf("\u0077\u006f\u0072k\u0062\u006f\u006f\u006b\u002f\u0053\u0068\u0065\u0065\u0074\u005b\u0025\u0064\u005d\u0020\u0068\u0061\u0073\u0020\u0064\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065\u0020n\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027", _aedbg, _bced.Name())
		}
		_edgb[_bced.Name()] = struct{}{}
		if _gacf := _bced.ValidateWithPath(_ge.Sprintf("\u0077o\u0072k\u0062\u006f\u006f\u006b\u002fS\u0068\u0065e\u0074\u005b\u0025\u0064\u005d", _aedbg)); _gacf != nil {
			return _gacf
		}
		if _acgab := _bced.Validate(); _acgab != nil {
			return _acgab
		}
	}
	return nil
}

// DataValidationCompare is a view on a data validation rule that is oriented
// towards value comparisons.
type DataValidationCompare struct{ _cdcg *_ca.CT_DataValidation }
type Fill struct {
	_ffad *_ca.CT_Fill
	_ada  *_ca.CT_Fills
}

// Border is a cell border configuraton.
type Border struct {
	_eac *_ca.CT_Border
	_gd  *_ca.CT_Borders
}

// CellReference returns the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (_cffd Comment) CellReference() string { return _cffd._ggb.RefAttr }

// SetHeight sets the height of the anchored object.
func (_fbdc OneCellAnchor) SetHeight(h _ee.Distance) { _fbdc._bbbd.Ext.CyAttr = int64(h / _ee.EMU) }

// PasswordHash returns the password hash for a workbook using the modified
// spreadsheetML password hash that is compatible with Excel.
func PasswordHash(s string) string {
	_ceed := uint16(0)
	if len(s) > 0 {
		for _fdaa := len(s) - 1; _fdaa >= 0; _fdaa-- {
			_gecce := s[_fdaa]
			_ceed = ((_ceed >> 14) & 0x01) | ((_ceed << 1) & 0x7fff)
			_ceed ^= uint16(_gecce)
		}
		_ceed = ((_ceed >> 14) & 0x01) | ((_ceed << 1) & 0x7fff)
		_ceed ^= uint16(len(s))
		_ceed ^= (0x8000 | ('N' << 8) | 'K')
	}
	return _ge.Sprintf("\u0025\u0030\u0034\u0058", uint64(_ceed))
}

// SetShowValue controls if the cell value is displayed.
func (_badg DataBarScale) SetShowValue(b bool) { _badg._egfb.ShowValueAttr = _c.Bool(b) }

// X returns the inner wrapped XML type.
func (_agcc ConditionalFormatting) X() *_ca.CT_ConditionalFormatting { return _agcc._gde }

// DataValidationList is just a view on a DataValidation configured as a list.
// It presents a drop-down combo box for spreadsheet users to select values. The
// contents of the dropdown can either pull from a rang eof cells (SetRange) or
// specified directly (SetValues).
type DataValidationList struct{ _ggbf *_ca.CT_DataValidation }

// Name returns the sheet name
func (_bfed Sheet) Name() string { return _bfed._faf.NameAttr }

// SetRowOffset sets the row offset of the top-left anchor.
func (_ecca OneCellAnchor) SetRowOffset(m _ee.Distance) { _ecca.TopLeft().SetRowOffset(m) }
func (_fdga PatternFill) ClearFgColor()                 { _fdga._fbdg.FgColor = nil }

// SortOrder is a column sort order.
//go:generate stringer -type=SortOrder
type SortOrder byte

func (_eec Cell) Reference() string {
	if _eec._ec.RAttr != nil {
		return *_eec._ec.RAttr
	}
	return ""
}

// SetHidden controls the visibility of a column.
func (_dfb Column) SetHidden(b bool) {
	if !b {
		_dfb._gfd.HiddenAttr = nil
	} else {
		_dfb._gfd.HiddenAttr = _c.Bool(true)
	}
}

type PatternFill struct {
	_fbdg *_ca.CT_PatternFill
	_gbec *_ca.CT_Fill
}

// AddRow adds a new row to a sheet.  You can mix this with numbered rows,
// however it will get confusing. You should prefer to use either automatically
// numbered rows with AddRow or manually numbered rows with Row/AddNumberedRow
func (_cbfb *Sheet) AddRow() Row {
	_addbe := uint32(0)
	_dead := uint32(len(_cbfb._fcbf.SheetData.Row))
	if _dead > 0 && _cbfb._fcbf.SheetData.Row[_dead-1].RAttr != nil && *_cbfb._fcbf.SheetData.Row[_dead-1].RAttr == _dead {
		return _cbfb.addNumberedRowFast(_dead + 1)
	}
	for _, _bbdg := range _cbfb._fcbf.SheetData.Row {
		if _bbdg.RAttr != nil && *_bbdg.RAttr > _addbe {
			_addbe = *_bbdg.RAttr
		}
	}
	return _cbfb.AddNumberedRow(_addbe + 1)
}

// SetHeight is a nop-op.
func (_geee TwoCellAnchor) SetHeight(h _ee.Distance) {}

// SetColOffset sets the column offset of the top-left of the image in fixed units.
func (_daa AbsoluteAnchor) SetColOffset(m _ee.Distance) {
	_daa._bdf.Pos.XAttr.ST_CoordinateUnqualified = _c.Int64(int64(m / _ee.EMU))
}

// SetWidth sets the width of the anchored object.
func (_bf AbsoluteAnchor) SetWidth(w _ee.Distance) { _bf._bdf.Ext.CxAttr = int64(w / _ee.EMU) }

// Drawing is a drawing overlay on a sheet.  Only a single drawing is allowed
// per sheet, so to display multiple charts and images on a single sheet, they
// must be added to the same drawing.
type Drawing struct {
	_ffa  *Workbook
	_acgf *_df.WsDr
}

// ClearProtection clears all workbook protections.
func (_aafef *Workbook) ClearProtection() { _aafef._geca.WorkbookProtection = nil }

// ClearProtection removes any protections applied to teh sheet.
func (_fgcd *Sheet) ClearProtection() { _fgcd._fcbf.SheetProtection = nil }

// IsStructureLocked returns whether the workbook structure is locked.
func (_fgffc WorkbookProtection) IsStructureLocked() bool {
	return _fgffc._eacd.LockStructureAttr != nil && *_fgffc._eacd.LockStructureAttr
}
func (_cdcga *Sheet) removeColumnFromMergedCells(_cbabe uint32) error {
	if _cdcga._fcbf.MergeCells == nil || _cdcga._fcbf.MergeCells.MergeCell == nil {
		return nil
	}
	_dfcc := []*_ca.CT_MergeCell{}
	for _, _abda := range _cdcga.MergedCells() {
		_eegb := _bdca(_abda.Reference(), _cbabe, true)
		if _eegb != "" {
			_abda.SetReference(_eegb)
			_dfcc = append(_dfcc, _abda.X())
		}
	}
	_cdcga._fcbf.MergeCells.MergeCell = _dfcc
	return nil
}

// SetWidth sets the width of the anchored object.
func (_cce OneCellAnchor) SetWidth(w _ee.Distance) { _cce._bbbd.Ext.CxAttr = int64(w / _ee.EMU) }

// SheetCount returns the number of sheets in the workbook.
func (_beee Workbook) SheetCount() int { return len(_beee._abcf) }

// HasFormula returns true if the cell contains formula.
func (_egaga *evalContext) HasFormula(cellRef string) bool {
	return _egaga._gag.Cell(cellRef).HasFormula()
}

// X returns the inner wrapped XML type.
func (_cecb MergedCell) X() *_ca.CT_MergeCell { return _cecb._gbb }

// Protection controls the protection on an individual sheet.
func (_bbdd *Sheet) Protection() SheetProtection {
	if _bbdd._fcbf.SheetProtection == nil {
		_bbdd._fcbf.SheetProtection = _ca.NewCT_SheetProtection()
	}
	return SheetProtection{_bbdd._fcbf.SheetProtection}
}

// IsEmpty checks if the cell style contains nothing.
func (_cgdc CellStyle) IsEmpty() bool {
	return _cgdc._gce == nil || _cgdc._bee == nil || _cgdc._eda == nil || _cgdc._eda.Xf == nil
}

// DataValidation controls cell validation
type DataValidation struct{ _cecf *_ca.CT_DataValidation }

// MergedCells returns the merged cell regions within the sheet.
func (_fcaf *Sheet) MergedCells() []MergedCell {
	if _fcaf._fcbf.MergeCells == nil {
		return nil
	}
	_gcaf := []MergedCell{}
	for _, _aead := range _fcaf._fcbf.MergeCells.MergeCell {
		_gcaf = append(_gcaf, MergedCell{_fcaf._ccae, _fcaf, _aead})
	}
	return _gcaf
}

// AddRun adds a new run of text to the cell.
func (_bdbc RichText) AddRun() RichTextRun {
	_aabf := _ca.NewCT_RElt()
	_bdbc._egacf.R = append(_bdbc._egacf.R, _aabf)
	return RichTextRun{_aabf}
}

// AddDrawing adds a drawing to a workbook.  However the drawing is not actually
// displayed or used until it's set on a sheet.
func (_gbebf *Workbook) AddDrawing() Drawing {
	_bbda := _df.NewWsDr()
	_gbebf._eddg = append(_gbebf._eddg, _bbda)
	_cfed := _c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.DrawingType, len(_gbebf._eddg))
	_gbebf.ContentTypes.AddOverride(_cfed, _c.DrawingContentType)
	_gbebf._fgd = append(_gbebf._fgd, _ef.NewRelationships())
	return Drawing{_gbebf, _bbda}
}

// SetZoom controls the zoom level of the sheet and is measured in percent. The
// default value is 100.
func (_bcfb SheetView) SetZoom(pct uint32) { _bcfb._bbgg.ZoomScaleAttr = &pct }

// GetFormat sets the number format code.
func (_aegd NumberFormat) GetFormat() string { return _aegd._geeb.FormatCodeAttr }

// BottomRight is a no-op.
func (_bfc OneCellAnchor) BottomRight() CellMarker { return CellMarker{} }

// SetMaxLength sets the maximum bar length in percent.
func (_acgaa DataBarScale) SetMaxLength(l uint32) { _acgaa._egfb.MaxLengthAttr = _c.Uint32(l) }

// Author returns the author of the comment
func (_ecb Comment) Author() string {
	if _ecb._ggb.AuthorIdAttr < uint32(len(_ecb._fca.Authors.Author)) {
		return _ecb._fca.Authors.Author[_ecb._ggb.AuthorIdAttr]
	}
	return ""
}
func (_bbbf Cell) GetRawValue() (string, error) {
	switch _bbbf._ec.TAttr {
	case _ca.ST_CellTypeInlineStr:
		if _bbbf._ec.Is == nil || _bbbf._ec.Is.T == nil {
			return "", nil
		}
		return *_bbbf._ec.Is.T, nil
	case _ca.ST_CellTypeS:
		if _bbbf._ec.V == nil {
			return "", nil
		}
		_fgc, _gga := _bd.Atoi(*_bbbf._ec.V)
		if _gga != nil {
			return "", _gga
		}
		return _bbbf._adb.SharedStrings.GetString(_fgc)
	case _ca.ST_CellTypeStr:
		if _bbbf._ec.F != nil {
			return _bbbf._ec.F.Content, nil
		}
	}
	if _bbbf._ec.V == nil {
		return "", nil
	}
	return *_bbbf._ec.V, nil
}

// AddNumberedRow adds a row with a given row number.  If you reuse a row number
// the resulting file will fail validation and fail to open in Office programs. Use
// Row instead which creates a new row or returns an existing row.
func (_adg *Sheet) AddNumberedRow(rowNum uint32) Row {
	_deb := _ca.NewCT_Row()
	_deb.RAttr = _c.Uint32(rowNum)
	_adg._fcbf.SheetData.Row = append(_adg._fcbf.SheetData.Row, _deb)
	_b.Slice(_adg._fcbf.SheetData.Row, func(_bbcc, _cgga int) bool {
		_bcg := _adg._fcbf.SheetData.Row[_bbcc].RAttr
		_bfa := _adg._fcbf.SheetData.Row[_cgga].RAttr
		if _bcg == nil {
			return true
		}
		if _bfa == nil {
			return true
		}
		return *_bcg < *_bfa
	})
	return Row{_adg._ccae, _adg, _deb}
}

// AddString adds a string to the shared string cache.
func (_dea SharedStrings) AddString(v string) int {
	if _egaf, _bbe := _dea._facb[v]; _bbe {
		return _egaf
	}
	_gbef := _ca.NewCT_Rst()
	_gbef.T = _c.String(v)
	_dea._dfea.Si = append(_dea._dfea.Si, _gbef)
	_cfbcc := len(_dea._dfea.Si) - 1
	_dea._facb[v] = _cfbcc
	_dea._dfea.CountAttr = _c.Uint32(uint32(len(_dea._dfea.Si)))
	_dea._dfea.UniqueCountAttr = _dea._dfea.CountAttr
	return _cfbcc
}

// CellStyles returns the list of defined cell styles
func (_bfbe StyleSheet) CellStyles() []CellStyle {
	_cfbdg := []CellStyle{}
	for _, _eacea := range _bfbe._bbbe.CellXfs.Xf {
		_cfbdg = append(_cfbdg, CellStyle{_bfbe._aba, _eacea, _bfbe._bbbe.CellXfs})
	}
	return _cfbdg
}

// AddRule adds and returns a new rule that can be configured.
func (_baa ConditionalFormatting) AddRule() ConditionalFormattingRule {
	_gdc := _ca.NewCT_CfRule()
	_baa._gde.CfRule = append(_baa._gde.CfRule, _gdc)
	_bce := ConditionalFormattingRule{_gdc}
	_bce.InitializeDefaults()
	_bce.SetPriority(int32(len(_baa._gde.CfRule) + 1))
	return _bce
}
func (_fbgd *Sheet) slideCellsLeft(_fdae []*_ca.CT_Cell) []*_ca.CT_Cell {
	for _, _gacg := range _fdae {
		_ecef, _fedf := _db.ParseCellReference(*_gacg.RAttr)
		if _fedf != nil {
			return _fdae
		}
		_gedd := _ecef.ColumnIdx - 1
		_baad := _db.IndexToColumn(_gedd) + _ge.Sprintf("\u0025\u0064", _ecef.RowIdx)
		_gacg.RAttr = &_baad
	}
	return _fdae
}

type WorkbookProtection struct{ _eacd *_ca.CT_WorkbookProtection }

// SetBold causes the text to be displayed in bold.
func (_bbc RichTextRun) SetBold(b bool) {
	_bbc.ensureRpr()
	_bbc._bead.RPr.B = _ca.NewCT_BooleanProperty()
	_bbc._bead.RPr.B.ValAttr = _c.Bool(b)
}

// X returns the inner wrapped XML type.
func (_fbc Column) X() *_ca.CT_Col { return _fbc._gfd }
func (_fefg Fill) SetPatternFill() PatternFill {
	_fefg._ffad.GradientFill = nil
	_fefg._ffad.PatternFill = _ca.NewCT_PatternFill()
	_fefg._ffad.PatternFill.PatternTypeAttr = _ca.ST_PatternTypeSolid
	return PatternFill{_fefg._ffad.PatternFill, _fefg._ffad}
}

// SetHeightCells is a no-op.
func (_def OneCellAnchor) SetHeightCells(int32) {}

// Rows returns all of the rows in a sheet.
func (_faggf *Sheet) Rows() []Row {
	_fae := []Row{}
	for _, _ece := range _faggf._fcbf.SheetData.Row {
		_fae = append(_fae, Row{_faggf._ccae, _faggf, _ece})
	}
	return _fae
}

// BottomRight returns the CellMaker for the bottom right corner of the anchor.
func (_eegd TwoCellAnchor) BottomRight() CellMarker { return CellMarker{_eegd._aeea.To} }

// SetType sets the type of the rule.
func (_gfdb ConditionalFormattingRule) SetType(t _ca.ST_CfType) { _gfdb._bfgc.TypeAttr = t }
func _dda(_fab string) bool {
	_fab = _ad.Replace(_fab, "\u0024", "", -1)
	if _egcfd := _adf.FindStringSubmatch(_ad.ToLower(_fab)); len(_egcfd) > 2 {
		_bab := _egcfd[1]
		_gdff, _ceb := _bd.Atoi(_egcfd[2])
		if _ceb != nil {
			return false
		}
		return _gdff <= 1048576 && _bab <= "\u007a\u007a"
	}
	return false
}

// New constructs a new workbook.
func New() *Workbook {
	_ffaa := &Workbook{}
	_ffaa._geca = _ca.NewWorkbook()
	_ffaa.AppProperties = _ef.NewAppProperties()
	_ffaa.CoreProperties = _ef.NewCoreProperties()
	_ffaa.StyleSheet = NewStyleSheet(_ffaa)
	_ffaa.Rels = _ef.NewRelationships()
	_ffaa._eceef = _ef.NewRelationships()
	_ffaa.Rels.AddRelationship(_c.RelativeFilename(_c.DocTypeSpreadsheet, "", _c.ExtendedPropertiesType, 0), _c.ExtendedPropertiesType)
	_ffaa.Rels.AddRelationship(_c.RelativeFilename(_c.DocTypeSpreadsheet, "", _c.CorePropertiesType, 0), _c.CorePropertiesType)
	_ffaa.Rels.AddRelationship(_c.RelativeFilename(_c.DocTypeSpreadsheet, "", _c.OfficeDocumentType, 0), _c.OfficeDocumentType)
	_ffaa._eceef.AddRelationship(_c.RelativeFilename(_c.DocTypeSpreadsheet, _c.OfficeDocumentType, _c.StylesType, 0), _c.StylesType)
	_ffaa.ContentTypes = _ef.NewContentTypes()
	_ffaa.ContentTypes.AddDefault("\u0076\u006d\u006c", _c.VMLDrawingContentType)
	_ffaa.ContentTypes.AddOverride(_c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.OfficeDocumentType, 0), "\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066o\u0072\u006d\u0061\u0074s\u002d\u006f\u0066\u0066\u0069\u0063e\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068e\u0065\u0074\u006d\u006c\u002e\u0073\u0068\u0065\u0065\u0074\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	_ffaa.ContentTypes.AddOverride(_c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.StylesType, 0), _c.SMLStyleSheetContentType)
	_ffaa.SharedStrings = NewSharedStrings()
	_ffaa.ContentTypes.AddOverride(_c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.SharedStringsType, 0), _c.SharedStringsContentType)
	_ffaa._eceef.AddRelationship(_c.RelativeFilename(_c.DocTypeSpreadsheet, _c.OfficeDocumentType, _c.SharedStringsType, 0), _c.SharedStringsType)
	return _ffaa
}

// Cells returns a slice of cells.  The cells can be manipulated, but appending
// to the slice will have no effect.
func (_acda Row) Cells() []Cell {
	_effg := []Cell{}
	_cecc := -1
	for _, _dfde := range _acda._gfcg.C {
		if _dfde.RAttr == nil {
			_c.Log("\u0052\u0041\u0074tr\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u0020\u0066o\u0072 \u0061 \u0063e\u006c\u006c\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			continue
		}
		_gbfg, _bccec := _db.ParseCellReference(*_dfde.RAttr)
		if _bccec != nil {
			_c.Log("\u0052\u0041\u0074t\u0072\u0020\u0069\u0073 \u0069\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u0020\u0063\u0065\u006c\u006c\u003a\u0020" + *_dfde.RAttr + ",\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			continue
		}
		_ggee := int(_gbfg.ColumnIdx)
		if _ggee-_cecc > 1 {
			for _bdcd := _cecc + 1; _bdcd < _ggee; _bdcd++ {
				_effg = append(_effg, _acda.Cell(_db.IndexToColumn(uint32(_bdcd))))
			}
		}
		_cecc = _ggee
		_effg = append(_effg, Cell{_acda._bgb, _acda._dcab, _acda._gfcg, _dfde})
	}
	return _effg
}

// Cell returns the actual cell behind the merged region
func (_gfga MergedCell) Cell() Cell {
	_gbfc := _gfga.Reference()
	if _bgfa := _ad.Index(_gfga.Reference(), "\u003a"); _bgfa != -1 {
		_gbfc = _gbfc[0:_bgfa]
		return _gfga._fcba.Cell(_gbfc)
	}
	return Cell{}
}

// X returns the inner wrapped XML type.
func (_ega ColorScale) X() *_ca.CT_ColorScale { return _ega._aaa }

// X returns the inner wrapped XML type.
func (_bdcfg Table) X() *_ca.Table { return _bdcfg._agbc }

// IsBool returns true if the cell is a boolean type cell.
func (_dcda Cell) IsBool() bool { return _dcda._ec.TAttr == _ca.ST_CellTypeB }

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (_cebg *Workbook) ClearCachedFormulaResults() {
	for _, _eabe := range _cebg.Sheets() {
		_eabe.ClearCachedFormulaResults()
	}
}

// ClearBorder clears any border configuration from the cell style.
func (_bgd CellStyle) ClearBorder() { _bgd._bee.BorderIdAttr = nil; _bgd._bee.ApplyBorderAttr = nil }

// SetYSplit sets the row split point
func (_ecdb SheetView) SetYSplit(v float64) {
	_ecdb.ensurePane()
	_ecdb._bbgg.Pane.YSplitAttr = _c.Float64(v)
}

// AddBorder creates a new empty border that can be applied to a cell style.
func (_gaeb StyleSheet) AddBorder() Border {
	_bggc := _ca.NewCT_Border()
	_gaeb._bbbe.Borders.Border = append(_gaeb._bbbe.Borders.Border, _bggc)
	_gaeb._bbbe.Borders.CountAttr = _c.Uint32(uint32(len(_gaeb._bbbe.Borders.Border)))
	return Border{_bggc, _gaeb._bbbe.Borders}
}
func (_eace Fills) X() *_ca.CT_Fills { return _eace._gad }

// Save writes the workbook out to a writer in the zipped xlsx format.
func (_cdfb *Workbook) Save(w _fd.Writer) error {
	if !_da.GetLicenseKey().IsLicensed() && !_cece {
		_ge.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
		_ge.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
		return _gb.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	}
	_abae := _ab.NewWriter(w)
	defer _abae.Close()
	_bdaf := _c.DocTypeSpreadsheet
	if _gaga := _cc.MarshalXML(_abae, _c.BaseRelsFilename, _cdfb.Rels.X()); _gaga != nil {
		return _gaga
	}
	if _gcfg := _cc.MarshalXMLByType(_abae, _bdaf, _c.ExtendedPropertiesType, _cdfb.AppProperties.X()); _gcfg != nil {
		return _gcfg
	}
	if _agad := _cc.MarshalXMLByType(_abae, _bdaf, _c.CorePropertiesType, _cdfb.CoreProperties.X()); _agad != nil {
		return _agad
	}
	_bagdf := _c.AbsoluteFilename(_bdaf, _c.OfficeDocumentType, 0)
	if _fadf := _cc.MarshalXML(_abae, _bagdf, _cdfb._geca); _fadf != nil {
		return _fadf
	}
	if _gefa := _cc.MarshalXML(_abae, _cc.RelationsPathFor(_bagdf), _cdfb._eceef.X()); _gefa != nil {
		return _gefa
	}
	if _bgcc := _cc.MarshalXMLByType(_abae, _bdaf, _c.StylesType, _cdfb.StyleSheet.X()); _bgcc != nil {
		return _bgcc
	}
	for _cgdb, _bbbc := range _cdfb._fee {
		if _bcda := _cc.MarshalXMLByTypeIndex(_abae, _bdaf, _c.ThemeType, _cgdb+1, _bbbc); _bcda != nil {
			return _bcda
		}
	}
	for _efca, _ecbf := range _cdfb._abcf {
		_ecbf.Dimension.RefAttr = Sheet{_cdfb, nil, _ecbf}.Extents()
		_ffed := _c.AbsoluteFilename(_bdaf, _c.WorksheetType, _efca+1)
		_cc.MarshalXML(_abae, _ffed, _ecbf)
		_cc.MarshalXML(_abae, _cc.RelationsPathFor(_ffed), _cdfb._adfgb[_efca].X())
	}
	if _cfd := _cc.MarshalXMLByType(_abae, _bdaf, _c.SharedStringsType, _cdfb.SharedStrings.X()); _cfd != nil {
		return _cfd
	}
	if _cdfb.CustomProperties.X() != nil {
		if _fgfg := _cc.MarshalXMLByType(_abae, _bdaf, _c.CustomPropertiesType, _cdfb.CustomProperties.X()); _fgfg != nil {
			return _fgfg
		}
	}
	if _cdfb.Thumbnail != nil {
		_deedb := _c.AbsoluteFilename(_bdaf, _c.ThumbnailType, 0)
		_faba, _cfba := _abae.Create(_deedb)
		if _cfba != nil {
			return _cfba
		}
		if _cfgaf := _f.Encode(_faba, _cdfb.Thumbnail, nil); _cfgaf != nil {
			return _cfgaf
		}
	}
	for _adgbe, _abfe := range _cdfb._agcf {
		_fgg := _c.AbsoluteFilename(_bdaf, _c.ChartType, _adgbe+1)
		_cc.MarshalXML(_abae, _fgg, _abfe)
	}
	for _eefg, _bbgea := range _cdfb._gfgbd {
		_gbaa := _c.AbsoluteFilename(_bdaf, _c.TableType, _eefg+1)
		_cc.MarshalXML(_abae, _gbaa, _bbgea)
	}
	for _cdb, _ccagb := range _cdfb._eddg {
		_ceeb := _c.AbsoluteFilename(_bdaf, _c.DrawingType, _cdb+1)
		_cc.MarshalXML(_abae, _ceeb, _ccagb)
		if !_cdfb._fgd[_cdb].IsEmpty() {
			_cc.MarshalXML(_abae, _cc.RelationsPathFor(_ceeb), _cdfb._fgd[_cdb].X())
		}
	}
	for _gefg, _gcdb := range _cdfb._gedac {
		_cc.MarshalXML(_abae, _c.AbsoluteFilename(_bdaf, _c.VMLDrawingType, _gefg+1), _gcdb)
	}
	for _cfda, _dfefb := range _cdfb.Images {
		if _afgac := _ef.AddImageToZip(_abae, _dfefb, _cfda+1, _c.DocTypeSpreadsheet); _afgac != nil {
			return _afgac
		}
	}
	if _eaae := _cc.MarshalXML(_abae, _c.ContentTypesFilename, _cdfb.ContentTypes.X()); _eaae != nil {
		return _eaae
	}
	for _bdcag, _ebab := range _cdfb._ebaf {
		if _ebab == nil {
			continue
		}
		_cc.MarshalXML(_abae, _c.AbsoluteFilename(_bdaf, _c.CommentsType, _bdcag+1), _ebab)
	}
	if _fdec := _cdfb.WriteExtraFiles(_abae); _fdec != nil {
		return _fdec
	}
	return _abae.Close()
}
func (_ceda *Workbook) onNewRelationship(_ebag *_cc.DecodeMap, _fceb, _afef string, _addg []*_ab.File, _aeed *_fbf.Relationship, _bbgf _cc.Target) error {
	_cdgf := _c.DocTypeSpreadsheet
	switch _afef {
	case _c.OfficeDocumentType:
		_ceda._geca = _ca.NewWorkbook()
		_ebag.AddTarget(_fceb, _ceda._geca, _afef, 0)
		_ceda._eceef = _ef.NewRelationships()
		_ebag.AddTarget(_cc.RelationsPathFor(_fceb), _ceda._eceef.X(), _afef, 0)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, 0)
	case _c.CorePropertiesType:
		_ebag.AddTarget(_fceb, _ceda.CoreProperties.X(), _afef, 0)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, 0)
	case _c.CustomPropertiesType:
		_ebag.AddTarget(_fceb, _ceda.CustomProperties.X(), _afef, 0)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, 0)
	case _c.ExtendedPropertiesType:
		_ebag.AddTarget(_fceb, _ceda.AppProperties.X(), _afef, 0)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, 0)
	case _c.WorksheetType:
		_fgge := _ca.NewWorksheet()
		_fgcfa := uint32(len(_ceda._abcf))
		_ceda._abcf = append(_ceda._abcf, _fgge)
		_ebag.AddTarget(_fceb, _fgge, _afef, _fgcfa)
		_degb := _ef.NewRelationships()
		_ebag.AddTarget(_cc.RelationsPathFor(_fceb), _degb.X(), _afef, 0)
		_ceda._adfgb = append(_ceda._adfgb, _degb)
		_ceda._ebaf = append(_ceda._ebaf, nil)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, len(_ceda._abcf))
	case _c.StylesType:
		_ceda.StyleSheet = NewStyleSheet(_ceda)
		_ebag.AddTarget(_fceb, _ceda.StyleSheet.X(), _afef, 0)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, 0)
	case _c.ThemeType:
		_gegd := _dc.NewTheme()
		_ceda._fee = append(_ceda._fee, _gegd)
		_ebag.AddTarget(_fceb, _gegd, _afef, 0)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, len(_ceda._fee))
	case _c.SharedStringsType:
		_ceda.SharedStrings = NewSharedStrings()
		_ebag.AddTarget(_fceb, _ceda.SharedStrings.X(), _afef, 0)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, 0)
	case _c.ThumbnailType:
		for _fced, _cgfdf := range _addg {
			if _cgfdf == nil {
				continue
			}
			if _cgfdf.Name == _fceb {
				_aeab, _cebf := _cgfdf.Open()
				if _cebf != nil {
					return _ge.Errorf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073", _cebf)
				}
				_ceda.Thumbnail, _, _cebf = _a.Decode(_aeab)
				_aeab.Close()
				if _cebf != nil {
					return _ge.Errorf("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073", _cebf)
				}
				_addg[_fced] = nil
			}
		}
	case _c.ImageType:
		for _bfag, _gccbb := range _addg {
			if _gccbb == nil {
				continue
			}
			if _gccbb.Name == _fceb {
				_ccbe, _deedg := _cc.ExtractToDiskTmp(_gccbb, _ceda.TmpPath)
				if _deedg != nil {
					return _deedg
				}
				_bfeg, _deedg := _ef.ImageFromStorage(_ccbe)
				if _deedg != nil {
					return _deedg
				}
				_dadg := _ef.MakeImageRef(_bfeg, &_ceda.DocBase, _ceda._eceef)
				_ceda.Images = append(_ceda.Images, _dadg)
				_addg[_bfag] = nil
			}
		}
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, len(_ceda.Images)+1)
	case _c.DrawingType:
		_acea := _df.NewWsDr()
		_bbfdg := uint32(len(_ceda._eddg))
		_ebag.AddTarget(_fceb, _acea, _afef, _bbfdg)
		_ceda._eddg = append(_ceda._eddg, _acea)
		_cdgfd := _ef.NewRelationships()
		_ebag.AddTarget(_cc.RelationsPathFor(_fceb), _cdgfd.X(), _afef, _bbfdg)
		_ceda._fgd = append(_ceda._fgd, _cdgfd)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, len(_ceda._eddg))
	case _c.VMLDrawingType:
		_ddbd := _fe.NewContainer()
		_ddeg := uint32(len(_ceda._gedac))
		_ebag.AddTarget(_fceb, _ddbd, _afef, _ddeg)
		_ceda._gedac = append(_ceda._gedac, _ddbd)
	case _c.CommentsType:
		_ceda._ebaf[_bbgf.Index] = _ca.NewComments()
		_ebag.AddTarget(_fceb, _ceda._ebaf[_bbgf.Index], _afef, _bbgf.Index)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, len(_ceda._ebaf))
	case _c.ChartType:
		_fbga := _gg.NewChartSpace()
		_deda := uint32(len(_ceda._agcf))
		_ebag.AddTarget(_fceb, _fbga, _afef, _deda)
		_ceda._agcf = append(_ceda._agcf, _fbga)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, len(_ceda._agcf))
	case _c.TableType:
		_bbbb := _ca.NewTable()
		_ddbb := uint32(len(_ceda._gfgbd))
		_ebag.AddTarget(_fceb, _bbbb, _afef, _ddbb)
		_ceda._gfgbd = append(_ceda._gfgbd, _bbbb)
		_aeed.TargetAttr = _c.RelativeFilename(_cdgf, _bbgf.Typ, _afef, len(_ceda._gfgbd))
	default:
		_c.Log("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0072\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068\u0069\u0070\u0020\u0025\u0073\u0020\u0025\u0073", _fceb, _afef)
	}
	return nil
}

// SetCol set the column of the cell marker.
func (_efb CellMarker) SetCol(col int32) { _efb._ggg.Col = col }
func _cec(_de _fbe.Time) _fbe.Time {
	_de = _de.UTC()
	return _fbe.Date(_de.Year(), _de.Month(), _de.Day(), _de.Hour(), _de.Minute(), _de.Second(), _de.Nanosecond(), _fbe.Local)
}

// SetBool sets the cell type to boolean and the value to the given boolean
// value.
func (_gda Cell) SetBool(v bool) {
	_gda.clearValue()
	_gda._ec.V = _c.String(_bd.Itoa(_gcd(v)))
	_gda._ec.TAttr = _ca.ST_CellTypeB
}

// Reference returns the region of cells that are merged.
func (_gdec MergedCell) Reference() string { return _gdec._gbb.RefAttr }

// ClearFill clears any fill configuration from the cell style.
func (_egcf CellStyle) ClearFill() { _egcf._bee.FillIdAttr = nil; _egcf._bee.ApplyFillAttr = nil }

// MoveTo moves the top-left of the anchored object.
func (_effd OneCellAnchor) MoveTo(col, row int32) {
	_effd.TopLeft().SetCol(col)
	_effd.TopLeft().SetRow(row)
}

// SetName sets the sheet name.
func (_cfbd *Sheet) SetName(name string) { _cfbd._faf.NameAttr = name }

type SheetProtection struct{ _aebb *_ca.CT_SheetProtection }

// OneCellAnchor is anchored to a top-left cell with a fixed with/height
// in distance.
type OneCellAnchor struct{ _bbbd *_df.CT_OneCellAnchor }

const _af = "\u00320\u0030\u0036\u002d\u00301\u002d\u0030\u0032\u0054\u00315\u003a0\u0034:\u0030\u0035\u005a\u0030\u0037\u003a\u00300"

var _ffd = _c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.SharedStringsType, 0)

// Sheets returns the sheets from the workbook.
func (_ffdd *Workbook) Sheets() []Sheet {
	_cbdc := []Sheet{}
	for _abeg, _fabab := range _ffdd._abcf {
		_ccac := _ffdd._geca.Sheets.Sheet[_abeg]
		_dggcf := Sheet{_ffdd, _ccac, _fabab}
		_cbdc = append(_cbdc, _dggcf)
	}
	return _cbdc
}

// X returns the inner wrapped XML type.
func (_deed IconScale) X() *_ca.CT_IconSet { return _deed._efdg }

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (_bgce *Sheet) ClearCachedFormulaResults() {
	for _, _gagfd := range _bgce.Rows() {
		for _, _cddg := range _gagfd.Cells() {
			if _cddg.X().F != nil {
				_cddg.X().V = nil
			}
		}
	}
}

// IconScale maps values to icons.
type IconScale struct{ _efdg *_ca.CT_IconSet }

func (_fedbe StyleSheet) GetNumberFormat(id uint32) NumberFormat {
	if id >= 0 && id < 50 {
		return CreateDefaultNumberFormat(StandardFormat(id))
	}
	for _, _ecaa := range _fedbe._bbbe.NumFmts.NumFmt {
		if _ecaa.NumFmtIdAttr == id {
			return NumberFormat{_fedbe._aba, _ecaa}
		}
	}
	return NumberFormat{}
}

const (
	DVCompareTypeWholeNumber = DVCompareType(_ca.ST_DataValidationTypeWhole)
	DVCompareTypeDecimal     = DVCompareType(_ca.ST_DataValidationTypeDecimal)
	DVCompareTypeDate        = DVCompareType(_ca.ST_DataValidationTypeDate)
	DVCompareTypeTime        = DVCompareType(_ca.ST_DataValidationTypeTime)
	DVompareTypeTextLength   = DVCompareType(_ca.ST_DataValidationTypeTextLength)
)

// RowNumber returns the row number (1-N), or zero if it is unset.
func (_bdfba Row) RowNumber() uint32 {
	if _bdfba._gfcg.RAttr != nil {
		return *_bdfba._gfcg.RAttr
	}
	return 0
}

// SetFgColor sets the *fill* foreground color.  As an example, the solid pattern foreground color becomes the
// background color of the cell when applied.
func (_abf PatternFill) SetFgColor(c _ae.Color) {
	_abf._fbdg.FgColor = _ca.NewCT_Color()
	_abf._fbdg.FgColor.RgbAttr = c.AsRGBAString()
}

// AddGradientStop adds a color gradient stop.
func (_cgce ColorScale) AddGradientStop(color _ae.Color) {
	_aaca := _ca.NewCT_Color()
	_aaca.RgbAttr = color.AsRGBAString()
	_cgce._aaa.Color = append(_cgce._aaa.Color, _aaca)
}

// GetValueAsNumber retrieves the cell's value as a number
func (_beb Cell) GetValueAsNumber() (float64, error) {
	if _beb._ec.V == nil && _beb._ec.Is == nil {
		return 0, nil
	}
	if _beb._ec.TAttr == _ca.ST_CellTypeS || !_ed.IsNumber(*_beb._ec.V) {
		return _d.NaN(), _gb.New("\u0063\u0065\u006c\u006c\u0020\u0069\u0073\u0020\u006e\u006f\u0074 \u006f\u0066\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020t\u0079\u0070\u0065")
	}
	return _bd.ParseFloat(*_beb._ec.V, 64)
}

// AddHyperlink adds a hyperlink to a sheet. Adding the hyperlink to the sheet
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (_eacf *Sheet) AddHyperlink(url string) _ef.Hyperlink {
	for _afec, _gfde := range _eacf._ccae._abcf {
		if _gfde == _eacf._fcbf {
			return _eacf._ccae._adfgb[_afec].AddHyperlink(url)
		}
	}
	return _ef.Hyperlink{}
}

// SetPasswordHash sets the password hash to the input.
func (_cgde SheetProtection) SetPasswordHash(pwHash string) {
	_cgde._aebb.PasswordAttr = _c.String(pwHash)
}
func (_bedc *Workbook) ensureSharedStringsRelationships() {
	_bfdf := false
	for _, _bdab := range _bedc.ContentTypes.X().Override {
		if _bdab.ContentTypeAttr == _c.SharedStringsContentType {
			_bfdf = true
			break
		}
	}
	if !_bfdf {
		_bedc.ContentTypes.AddOverride(_ffd, _c.SharedStringsContentType)
	}
	_eada := false
	for _, _deag := range _bedc._eceef.Relationships() {
		if _deag.X().TargetAttr == _dbab {
			_eada = true
			break
		}
	}
	if !_eada {
		_bedc._eceef.AddRelationship(_dbab, _c.SharedStringsType)
	}
}

// X returns the inner wrapped XML type.
func (_egfc NumberFormat) X() *_ca.CT_NumFmt { return _egfc._geeb }

// X returns the inner XML entity for a stylesheet.
func (_cbddf StyleSheet) X() *_ca.StyleSheet { return _cbddf._bbbe }

// RowOffset returns the offset from the row cell.
func (_agc CellMarker) RowOffset() _ee.Distance {
	if _agc._ggg.RowOff.ST_CoordinateUnqualified == nil {
		return 0
	}
	return _ee.Distance(float64(*_agc._ggg.RowOff.ST_CoordinateUnqualified) * _ee.EMU)
}

// RichTextRun is a segment of text within a cell that is directly formatted.
type RichTextRun struct{ _bead *_ca.CT_RElt }

// IsEmpty returns true if the cell is empty.
func (_bac Cell) IsEmpty() bool {
	return _bac._ec.TAttr == _ca.ST_CellTypeUnset && _bac._ec.V == nil && _bac._ec.F == nil
}

// X returns the inner wrapped XML type.
func (_afcf RichTextRun) X() *_ca.CT_RElt { return _afcf._bead }

// SetState sets the sheet view state (frozen/split/frozen-split)
func (_ggbbc SheetView) SetState(st _ca.ST_PaneState) {
	_ggbbc.ensurePane()
	_ggbbc._bbgg.Pane.StateAttr = st
}

// Row returns the row of the cell marker.
func (_ddc CellMarker) Row() int32      { return _ddc._ggg.Row }
func (_aeaa Font) SetSize(size float64) { _aeaa._acaf.Sz = []*_ca.CT_FontSize{{ValAttr: size}} }

// CellStyle is a formatting style for a cell.  CellStyles are spreadsheet global
// and can be applied to cells across sheets.
type CellStyle struct {
	_gce *Workbook
	_bee *_ca.CT_Xf
	_eda *_ca.CT_CellXfs
}

// SetRowOffset sets a column offset in absolute distance.
func (_egg CellMarker) SetRowOffset(m _ee.Distance) {
	_egg._ggg.RowOff.ST_CoordinateUnqualified = _c.Int64(int64(m / _ee.EMU))
}

// SetString sets the cell type to string, and the value to the given string,
// returning an ID from the shared strings table. To reuse a string, call
// SetStringByID with the ID returned.
func (_fec Cell) SetString(s string) int {
	_fec._adb.ensureSharedStringsRelationships()
	_fec.clearValue()
	_agf := _fec._adb.SharedStrings.AddString(s)
	_fec._ec.V = _c.String(_bd.Itoa(_agf))
	_fec._ec.TAttr = _ca.ST_CellTypeS
	return _agf
}

// SetWidth is a no-op.
func (_aeada TwoCellAnchor) SetWidth(w _ee.Distance) {}

// SetStyle sets the cell style for an entire column.
func (_gdf Column) SetStyle(cs CellStyle) { _gdf._gfd.StyleAttr = _c.Uint32(cs.Index()) }

// AnchorType is the type of anchor.
type AnchorType byte

const _dfdb = "_\u0078\u006c\u006e\u006d._\u0046i\u006c\u0074\u0065\u0072\u0044a\u0074\u0061\u0062\u0061\u0073\u0065"

func (_eaeb DataValidation) clear() {
	_eaeb._cecf.Formula1 = _c.String("\u0030")
	_eaeb._cecf.Formula2 = _c.String("\u0030")
}

// SetFormulaArray sets the cell type to formula array, and the raw formula to
// the given string. This is equivlent to entering a formula and pressing
// Ctrl+Shift+Enter in Excel.
func (_be Cell) SetFormulaArray(s string) {
	_baf := _gbf.ParseString(s)
	if _baf == nil {
		return
	}
	_be.clearValue()
	_be._ec.TAttr = _ca.ST_CellTypeStr
	_be._ec.F = _ca.NewCT_CellFormula()
	_be._ec.F.TAttr = _ca.ST_CellFormulaTypeArray
	_be._ec.F.Content = s
}

// SetColor sets teh color of the databar.
func (_edd DataBarScale) SetColor(c _ae.Color) {
	_edd._egfb.Color = _ca.NewCT_Color()
	_edd._egfb.Color.RgbAttr = c.AsRGBAString()
}

// Type returns the type of anchor
func (_abc AbsoluteAnchor) Type() AnchorType { return AnchorTypeAbsolute }

// SetIcons sets the icon set to use for display.
func (_cag IconScale) SetIcons(t _ca.ST_IconSetType) { _cag._efdg.IconSetAttr = t }

const (
	DVOpGreater = _ca.ST_DataValidationOperatorGreaterThanOrEqual
)

// GetLocked returns true if the cell is locked.
func (_efcc *evalContext) GetLocked(cellRef string) bool { return _efcc._gag.Cell(cellRef).getLocked() }
func (_cde Fills) AddFill() Fill {
	_afgd := _ca.NewCT_Fill()
	_cde._gad.Fill = append(_cde._gad.Fill, _afgd)
	_cde._gad.CountAttr = _c.Uint32(uint32(len(_cde._gad.Fill)))
	return Fill{_afgd, _cde._gad}
}

// Validate validates the sheet, returning an error if it is found to be invalid.
func (_beeb Sheet) Validate() error {
	_ffcd := []func() error{_beeb.validateRowCellNumbers, _beeb.validateMergedCells, _beeb.validateSheetNames}
	for _, _bdffg := range _ffcd {
		if _aaeb := _bdffg(); _aaeb != nil {
			return _aaeb
		}
	}
	if _aede := _beeb._fcbf.Validate(); _aede != nil {
		return _aede
	}
	return _beeb._fcbf.Validate()
}

// Wrapped returns true if the cell will wrap text.
func (_ecd CellStyle) Wrapped() bool {
	if _ecd._bee.Alignment == nil {
		return false
	}
	if _ecd._bee.Alignment.WrapTextAttr == nil {
		return false
	}
	return *_ecd._bee.Alignment.WrapTextAttr
}

// AddFormatValue adds a format value to be used to determine the cell background.
func (_bed ColorScale) AddFormatValue(t _ca.ST_CfvoType, val string) {
	_dggc := _ca.NewCT_Cfvo()
	_dggc.TypeAttr = t
	_dggc.ValAttr = _c.String(val)
	_bed._aaa.Cfvo = append(_bed._aaa.Cfvo, _dggc)
}

// SetCachedFormulaResult sets the cached result of a formula. This is normally
// not needed but is used internally when expanding an array formula.
func (_aded Cell) SetCachedFormulaResult(s string) { _aded._ec.V = &s }

// X returns the inner wrapped XML type.
func (_ace Font) X() *_ca.CT_Font { return _ace._acaf }

// SetHeight sets the row height in points.
func (_dbeb Row) SetHeight(d _ee.Distance) {
	_dbeb._gfcg.HtAttr = _c.Float64(float64(d))
	_dbeb._gfcg.CustomHeightAttr = _c.Bool(true)
}

// TopLeft is a no-op.
func (_cb AbsoluteAnchor) TopLeft() CellMarker { return CellMarker{} }

// AddMergedCells merges cells within a sheet.
func (_eadd *Sheet) AddMergedCells(fromRef, toRef string) MergedCell {
	if _eadd._fcbf.MergeCells == nil {
		_eadd._fcbf.MergeCells = _ca.NewCT_MergeCells()
	}
	_acb := _ca.NewCT_MergeCell()
	_acb.RefAttr = _ge.Sprintf("\u0025\u0073\u003a%\u0073", fromRef, toRef)
	_eadd._fcbf.MergeCells.MergeCell = append(_eadd._fcbf.MergeCells.MergeCell, _acb)
	_eadd._fcbf.MergeCells.CountAttr = _c.Uint32(uint32(len(_eadd._fcbf.MergeCells.MergeCell)))
	return MergedCell{_eadd._ccae, _eadd, _acb}
}

// Open opens and reads a workbook from a file (.xlsx).
func Open(filename string) (*Workbook, error) {
	_facd, _cagf := _g.Open(filename)
	if _cagf != nil {
		return nil, _ge.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _cagf)
	}
	defer _facd.Close()
	_gbgd, _cagf := _g.Stat(filename)
	if _cagf != nil {
		return nil, _ge.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _cagf)
	}
	_aabgc, _cagf := Read(_facd, _gbgd.Size())
	if _cagf != nil {
		return nil, _cagf
	}
	_ffae, _ := _ba.Abs(_ba.Dir(filename))
	_aabgc._ffaab = _ba.Join(_ffae, filename)
	return _aabgc, nil
}

// Extents returns the sheet extents in the form "A1:B15". This requires
// scanning the entire sheet.
func (_dgfa Sheet) Extents() string {
	_adbd, _gdcde, _gffc, _acgb := _dgfa.ExtentsIndex()
	return _ge.Sprintf("\u0025s\u0025\u0064\u003a\u0025\u0073\u0025d", _adbd, _gdcde, _gffc, _acgb)
}

// SetPasswordHash sets the password hash to the input.
func (_gbbeb WorkbookProtection) SetPasswordHash(pwHash string) {
	_gbbeb._eacd.WorkbookPasswordAttr = _c.String(pwHash)
}

// TwoCellAnchor is an anchor that is attached to a top-left cell with a fixed
// width/height in cells.
type TwoCellAnchor struct{ _aeea *_df.CT_TwoCellAnchor }

// SetCellReference sets the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (_fdf Comment) SetCellReference(cellRef string) { _fdf._ggb.RefAttr = cellRef }

// X returns the inner wrapped XML type.
func (_abb Drawing) X() *_df.WsDr { return _abb._acgf }
func (_fccd *Sheet) setShared(_aggb string, _gecf, _gaca _db.CellReference, _gddg string) {
	_bggb := _fccd.FormulaContext()
	_cad := _gbf.NewEvaluator()
	for _fga := _gecf.RowIdx; _fga <= _gaca.RowIdx; _fga++ {
		for _dcfgb := _gecf.ColumnIdx; _dcfgb <= _gaca.ColumnIdx; _dcfgb++ {
			_dcfa := _fga - _gecf.RowIdx
			_bbge := _dcfgb - _gecf.ColumnIdx
			_bggb.SetOffset(_bbge, _dcfa)
			_ggcc := _cad.Eval(_bggb, _gddg)
			_daga := _ge.Sprintf("\u0025\u0073\u0025\u0064", _db.IndexToColumn(_dcfgb), _fga)
			_cfef := _fccd.Cell(_daga)
			if _ggcc.Type == _gbf.ResultTypeNumber {
				_cfef.X().TAttr = _ca.ST_CellTypeN
			} else {
				_cfef.X().TAttr = _ca.ST_CellTypeInlineStr
			}
			_cfef.X().V = _c.String(_ggcc.Value())
		}
	}
	_ = _cad
	_ = _bggb
}
func (_daf *Sheet) setList(_eafc string, _cfcc _gbf.Result) error {
	_bbef, _fged := _db.ParseCellReference(_eafc)
	if _fged != nil {
		return _fged
	}
	_cabf := _daf.Row(_bbef.RowIdx)
	for _gabb, _bdcge := range _cfcc.ValueList {
		_babf := _cabf.Cell(_db.IndexToColumn(_bbef.ColumnIdx + uint32(_gabb)))
		if _bdcge.Type != _gbf.ResultTypeEmpty {
			if _bdcge.IsBoolean {
				_babf.SetBool(_bdcge.ValueNumber != 0)
			} else {
				_babf.SetCachedFormulaResult(_bdcge.String())
			}
		}
	}
	return nil
}

// HasFormula returns true if the cell has an asoociated formula.
func (_dd Cell) HasFormula() bool { return _dd._ec.F != nil }

// GetSheet returns a sheet by name, or an error if a sheet by the given name
// was not found.
func (_aeege *Workbook) GetSheet(name string) (Sheet, error) {
	for _, _cafe := range _aeege.Sheets() {
		if _cafe.Name() == name {
			return _cafe, nil
		}
	}
	return Sheet{}, ErrorNotFound
}
func (_ebg Sheet) validateRowCellNumbers() error {
	_fcac := map[uint32]struct{}{}
	for _, _gggd := range _ebg._fcbf.SheetData.Row {
		if _gggd.RAttr != nil {
			if _, _ddcc := _fcac[*_gggd.RAttr]; _ddcc {
				return _ge.Errorf("\u0027%\u0073'\u0020\u0072\u0065\u0075\u0073e\u0064\u0020r\u006f\u0077\u0020\u0025\u0064", _ebg.Name(), *_gggd.RAttr)
			}
			_fcac[*_gggd.RAttr] = struct{}{}
		}
		_geccc := map[string]struct{}{}
		for _, _dcb := range _gggd.C {
			if _dcb.RAttr == nil {
				continue
			}
			if _, _edag := _geccc[*_dcb.RAttr]; _edag {
				return _ge.Errorf("\u0027\u0025\u0073\u0027 r\u0065\u0075\u0073\u0065\u0064\u0020\u0063\u0065\u006c\u006c\u0020\u0025\u0073", _ebg.Name(), *_dcb.RAttr)
			}
			_geccc[*_dcb.RAttr] = struct{}{}
		}
	}
	return nil
}

const (
	DVCompareOpEqual        = DVCompareOp(_ca.ST_DataValidationOperatorEqual)
	DVCompareOpBetween      = DVCompareOp(_ca.ST_DataValidationOperatorBetween)
	DVCompareOpNotBetween   = DVCompareOp(_ca.ST_DataValidationOperatorNotBetween)
	DVCompareOpNotEqual     = DVCompareOp(_ca.ST_DataValidationOperatorNotEqual)
	DVCompareOpGreater      = DVCompareOp(_ca.ST_DataValidationOperatorGreaterThan)
	DVCompareOpGreaterEqual = DVCompareOp(_ca.ST_DataValidationOperatorGreaterThanOrEqual)
	DVCompareOpLess         = DVCompareOp(_ca.ST_DataValidationOperatorLessThan)
	DVCompareOpLessEqual    = DVCompareOp(_ca.ST_DataValidationOperatorLessThanOrEqual)
)

var (
	_fgcda = [...]uint8{0, 21, 46, 61, 76, 91}
	_bfdd  = [...]uint8{0, 21, 37, 53, 69, 85, 103, 119, 135, 151, 167, 185, 201, 217, 239}
	_ggbdd = [...]uint8{0, 16, 32, 48, 64}
	_cgec  = [...]uint8{0, 16, 32, 48, 64, 80}
)

// DefinedName is a named range, formula, etc.
type DefinedName struct{ _gagf *_ca.CT_DefinedName }

// X returns the inner wrapped XML type.
func (_bgbg Sheet) X() *_ca.Worksheet { return _bgbg._fcbf }

// LockSheet controls the locking of the sheet.
func (_eba SheetProtection) LockSheet(b bool) {
	if !b {
		_eba._aebb.SheetAttr = nil
	} else {
		_eba._aebb.SheetAttr = _c.Bool(true)
	}
}

// LockWindow controls the locking of the workbook windows.
func (_egde WorkbookProtection) LockWindow(b bool) {
	if !b {
		_egde._eacd.LockWindowsAttr = nil
	} else {
		_egde._eacd.LockWindowsAttr = _c.Bool(true)
	}
}

// Priority returns the rule priority
func (_gdgg ConditionalFormattingRule) Priority() int32 { return _gdgg._bfgc.PriorityAttr }

// CopySheetByName copies the existing sheet with the name `name` and puts its copy with the name `copiedSheetName`.
func (_gbdf *Workbook) CopySheetByName(name, copiedSheetName string) (Sheet, error) {
	_cbae := -1
	for _bege, _adgg := range _gbdf.Sheets() {
		if name == _adgg.Name() {
			_cbae = _bege
			break
		}
	}
	if _cbae == -1 {
		return Sheet{}, ErrorNotFound
	}
	return _gbdf.CopySheet(_cbae, copiedSheetName)
}

// GetString retrieves a string from the shared strings table by index.
func (_cgg SharedStrings) GetString(id int) (string, error) {
	if id < 0 {
		return "", _ge.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0073\u0074\u0072\u0069ng \u0069nd\u0065\u0078\u0020\u0025\u0064\u002c\u0020mu\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030", id)
	}
	if id > len(_cgg._dfea.Si) {
		return "", _ge.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002c\u0020\u0074\u0061b\u006c\u0065\u0020\u006f\u006e\u006c\u0079\u0020\u0068\u0061\u0073\u0020\u0025\u0064 \u0076a\u006c\u0075\u0065\u0073", id, len(_cgg._dfea.Si))
	}
	_aabd := _cgg._dfea.Si[id]
	if _aabd.T != nil {
		return *_aabd.T, nil
	}
	return "", nil
}

// CopySheet copies the existing sheet at index `ind` and puts its copy with the name `copiedSheetName`.
func (_fade *Workbook) CopySheet(ind int, copiedSheetName string) (Sheet, error) {
	if _fade.SheetCount() <= ind {
		return Sheet{}, ErrorNotFound
	}
	var _dfeaa _ef.Relationship
	for _, _ebdb := range _fade._eceef.Relationships() {
		if _ebdb.ID() == _fade._geca.Sheets.Sheet[ind].IdAttr {
			var _fabd bool
			if _dfeaa, _fabd = _fade._eceef.CopyRelationship(_ebdb.ID()); !_fabd {
				return Sheet{}, ErrorNotFound
			}
			break
		}
	}
	_fade.ContentTypes.CopyOverride(_c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.WorksheetContentType, ind+1), _c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.WorksheetContentType, len(_fade.ContentTypes.X().Override)))
	_fgaf := *_fade._abcf[ind]
	_fade._abcf = append(_fade._abcf, &_fgaf)
	var _fbgde uint32 = 0
	for _, _acdd := range _fade._geca.Sheets.Sheet {
		if _acdd.SheetIdAttr > _fbgde {
			_fbgde = _acdd.SheetIdAttr
		}
	}
	_fbgde++
	_ggdga := *_fade._geca.Sheets.Sheet[ind]
	_ggdga.IdAttr = _dfeaa.ID()
	_ggdga.NameAttr = copiedSheetName
	_ggdga.SheetIdAttr = _fbgde
	_fade._geca.Sheets.Sheet = append(_fade._geca.Sheets.Sheet, &_ggdga)
	_fgfd := _ef.NewRelationshipsCopy(_fade._adfgb[ind])
	_fade._adfgb = append(_fade._adfgb, _fgfd)
	_aacb := _fade._ebaf[ind]
	if _aacb == nil {
		_fade._ebaf = append(_fade._ebaf, nil)
	} else {
		_aeadc := *_aacb
		_fade._ebaf = append(_fade._ebaf, &_aeadc)
	}
	_eded := Sheet{_fade, &_ggdga, &_fgaf}
	return _eded, nil
}

// SetAutoFilter creates autofilters on the sheet. These are the automatic
// filters that are common for a header row.  The RangeRef should be of the form
// "A1:C5" and cover the entire range of cells to be filtered, not just the
// header. SetAutoFilter replaces any existing auto filter on the sheet.
func (_ebb *Sheet) SetAutoFilter(rangeRef string) {
	rangeRef = _ad.Replace(rangeRef, "\u0024", "", -1)
	_ebb._fcbf.AutoFilter = _ca.NewCT_AutoFilter()
	_ebb._fcbf.AutoFilter.RefAttr = _c.String(rangeRef)
	_cgfg := "\u0027" + _ebb.Name() + "\u0027\u0021"
	var _bfee DefinedName
	for _, _aafee := range _ebb._ccae.DefinedNames() {
		if _aafee.Name() == _dfdb {
			if _ad.HasPrefix(_aafee.Content(), _cgfg) {
				_bfee = _aafee
				_bfee.SetContent(_ebb.RangeReference(rangeRef))
				break
			}
		}
	}
	if _bfee.X() == nil {
		_bfee = _ebb._ccae.AddDefinedName(_dfdb, _ebb.RangeReference(rangeRef))
	}
	for _cdcgg, _acgd := range _ebb._ccae._abcf {
		if _acgd == _ebb._fcbf {
			_bfee.SetLocalSheetID(uint32(_cdcgg))
		}
	}
}
func (_ce Border) SetRight(style _ca.ST_BorderStyle, c _ae.Color) {
	if _ce._eac.Right == nil {
		_ce._eac.Right = _ca.NewCT_BorderPr()
	}
	_ce._eac.Right.Color = _ca.NewCT_Color()
	_ce._eac.Right.Color.RgbAttr = c.AsRGBAString()
	_ce._eac.Right.StyleAttr = style
}

// SetDate sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel). The cell is not styled via this method, so it will
// display as a number. SetDateWithStyle should normally be used instead.
func (_bcc Cell) SetDate(d _fbe.Time) {
	_bcc.clearValue()
	d = _cgc(d)
	_gfc := _bcc._adb.Epoch()
	if d.Before(_gfc) {
		_c.Log("d\u0061\u0074\u0065\u0073\u0020\u0062e\u0066\u006f\u0072\u0065\u0020\u00319\u0030\u0030\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064")
		return
	}
	_gaf := d.Sub(_gfc)
	_dgf := new(_ag.Float)
	_bdcc := new(_ag.Float)
	_bdcc.SetPrec(128)
	_bdcc.SetUint64(uint64(_gaf))
	_fbb := new(_ag.Float)
	_fbb.SetUint64(24 * 60 * 60 * 1e9)
	_dgf.Quo(_bdcc, _fbb)
	_gbe, _ := _dgf.Uint64()
	_bcc._ec.V = _c.Stringf("\u0025\u0064", _gbe)
}

// SetNumberFormatStandard sets the format based off of the ECMA 376 standard formats.  These
// formats are standardized and don't need to be defined in the styles.
func (_gac CellStyle) SetNumberFormatStandard(s StandardFormat) {
	_gac._bee.NumFmtIdAttr = _c.Uint32(uint32(s))
	_gac._bee.ApplyNumberFormatAttr = _c.Bool(true)
}

// Type returns the type of anchor
func (_fcfb OneCellAnchor) Type() AnchorType { return AnchorTypeOneCell }

// SetStyle sets the style to be used for conditional rules
func (_acd ConditionalFormattingRule) SetStyle(d DifferentialStyle) {
	_acd._bfgc.DxfIdAttr = _c.Uint32(d.Index())
}

// Comment is a single comment within a sheet.
type Comment struct {
	_dgcb *Workbook
	_ggb  *_ca.CT_Comment
	_fca  *_ca.Comments
}

func (_dggb StyleSheet) GetCellStyle(id uint32) CellStyle {
	for _cgfd, _cbdg := range _dggb._bbbe.CellXfs.Xf {
		if uint32(_cgfd) == id {
			return CellStyle{_dggb._aba, _cbdg, _dggb._bbbe.CellXfs}
		}
	}
	return CellStyle{}
}
func (_egf *evalContext) Sheet(name string) _gbf.Context {
	for _, _gfgb := range _egf._gag._ccae.Sheets() {
		if _gfgb.Name() == name {
			return _gfgb.FormulaContext()
		}
	}
	return _gbf.InvalidReferenceContext
}

// InitializeDefaults initializes a border to its defaulte empty values.
func (_cg Border) InitializeDefaults() {
	_cg._eac.Left = _ca.NewCT_BorderPr()
	_cg._eac.Bottom = _ca.NewCT_BorderPr()
	_cg._eac.Right = _ca.NewCT_BorderPr()
	_cg._eac.Top = _ca.NewCT_BorderPr()
	_cg._eac.Diagonal = _ca.NewCT_BorderPr()
}

// SetRow set the row of the cell marker.
func (_deeg CellMarker) SetRow(row int32) { _deeg._ggg.Row = row }

// SetText sets the text to be displayed.
func (_gdgd RichTextRun) SetText(s string) { _gdgd._bead.T = s }

// SetDataBar configures the rule as a data bar, removing existing
// configuration.
func (_eacg ConditionalFormattingRule) SetDataBar() DataBarScale {
	_eacg.clear()
	_eacg.SetType(_ca.ST_CfTypeDataBar)
	_eacg._bfgc.DataBar = _ca.NewCT_DataBar()
	_aaf := DataBarScale{_eacg._bfgc.DataBar}
	_aaf.SetShowValue(true)
	_aaf.SetMinLength(10)
	_aaf.SetMaxLength(90)
	return _aaf
}
func (_ccag Sheet) validateSheetNames() error {
	if len(_ccag.Name()) > 31 {
		return _ge.Errorf("\u0073\u0068\u0065\u0065\u0074 \u006e\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027\u0020\u0068\u0061\u0073 \u0025\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u002c\u0020\u006d\u0061\u0078\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0069\u0073\u0020\u00331", _ccag.Name(), len(_ccag.Name()))
	}
	return nil
}

// Fonts returns the list of fonts defined in the stylesheet.
func (_gfe StyleSheet) Fonts() []Font {
	_efaa := []Font{}
	for _, _gfgc := range _gfe._bbbe.Fonts.Font {
		_efaa = append(_efaa, Font{_gfgc, _gfe._bbbe})
	}
	return _efaa
}

// SetRichTextString sets the cell to rich string mode and returns a struct that
// can be used to add formatted text to the cell.
func (_bbf Cell) SetRichTextString() RichText {
	_bbf.clearValue()
	_bbf._ec.Is = _ca.NewCT_Rst()
	_bbf._ec.TAttr = _ca.ST_CellTypeInlineStr
	return RichText{_bbf._ec.Is}
}

// Index returns the index of the border for use with a cell style.
func (_efc Border) Index() uint32 {
	for _fdg, _acf := range _efc._gd.Border {
		if _acf == _efc._eac {
			return uint32(_fdg)
		}
	}
	return 0
}

// SetMinLength sets the minimum bar length in percent.
func (_efbb DataBarScale) SetMinLength(l uint32) { _efbb._egfb.MinLengthAttr = _c.Uint32(l) }
func (_fg Border) SetLeft(style _ca.ST_BorderStyle, c _ae.Color) {
	if _fg._eac.Left == nil {
		_fg._eac.Left = _ca.NewCT_BorderPr()
	}
	_fg._eac.Left.Color = _ca.NewCT_Color()
	_fg._eac.Left.Color.RgbAttr = c.AsRGBAString()
	_fg._eac.Left.StyleAttr = style
}

// SetShowRuler controls the visibility of the ruler
func (_cabb SheetView) SetShowRuler(b bool) {
	if !b {
		_cabb._bbgg.ShowRulerAttr = _c.Bool(false)
	} else {
		_cabb._bbgg.ShowRulerAttr = nil
	}
}

// Close closes the workbook, removing any temporary files that might have been
// created when opening a document.
func (_efcb *Workbook) Close() error {
	if _efcb.TmpPath != "" {
		return _cf.RemoveAll(_efcb.TmpPath)
	}
	return nil
}

// Comments returns the comments for a sheet.
func (_dcaf *Sheet) Comments() Comments {
	for _aefe, _cbab := range _dcaf._ccae._abcf {
		if _cbab == _dcaf._fcbf {
			if _dcaf._ccae._ebaf[_aefe] == nil {
				_dcaf._ccae._ebaf[_aefe] = _ca.NewComments()
				_dcaf._ccae._adfgb[_aefe].AddAutoRelationship(_c.DocTypeSpreadsheet, _c.WorksheetType, _aefe+1, _c.CommentsType)
				_dcaf._ccae.ContentTypes.AddOverride(_c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.CommentsType, _aefe+1), _c.CommentsContentType)
			}
			if len(_dcaf._ccae._gedac) == 0 {
				_dcaf._ccae._gedac = append(_dcaf._ccae._gedac, _fe.NewCommentDrawing())
				_gcaa := _dcaf._ccae._adfgb[_aefe].AddAutoRelationship(_c.DocTypeSpreadsheet, _c.WorksheetType, 1, _c.VMLDrawingType)
				if _dcaf._fcbf.LegacyDrawing == nil {
					_dcaf._fcbf.LegacyDrawing = _ca.NewCT_LegacyDrawing()
				}
				_dcaf._fcbf.LegacyDrawing.IdAttr = _gcaa.ID()
			}
			return Comments{_dcaf._ccae, _dcaf._ccae._ebaf[_aefe]}
		}
	}
	_c.Log("\u0061\u0074\u0074\u0065\u006dp\u0074\u0065\u0064\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073 \u0063\u006f\u006d\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u006e\u006f\u006e\u002d\u0065\u0078\u0069\u0073\u0074\u0065\u006e\u0074\u0020\u0073\u0068\u0065\u0065t")
	return Comments{}
}

// SetRange sets the cell or range of cells that the validation should apply to.
// It can be a single cell (e.g. "A1") or a range of cells (e.g. "A1:B5")
func (_edc DataValidation) SetRange(cellRange string) { _edc._cecf.SqrefAttr = _ca.ST_Sqref{cellRange} }

// Fills returns a Fills object that can be used to add/create/edit fills.
func (_fdag StyleSheet) Fills() Fills { return Fills{_fdag._bbbe.Fills} }

// Type returns the type of the rule
func (_acg ConditionalFormattingRule) Type() _ca.ST_CfType { return _acg._bfgc.TypeAttr }

// X returns the inner wrapped XML type.
func (_cgad Row) X() *_ca.CT_Row { return _cgad._gfcg }

// ClearSheetViews clears the list of sheet views.  This will clear the results
// of AddView() or SetFrozen.
func (_cfceb *Sheet) ClearSheetViews() { _cfceb._fcbf.SheetViews = nil }
func (_bdcf RichTextRun) ensureRpr() {
	if _bdcf._bead.RPr == nil {
		_bdcf._bead.RPr = _ca.NewCT_RPrElt()
	}
}

// SetTopLeft sets the top left visible cell after the split.
func (_bgfed SheetView) SetTopLeft(cellRef string) {
	_bgfed.ensurePane()
	_bgfed._bbgg.Pane.TopLeftCellAttr = &cellRef
}

var _dbab = _c.RelativeFilename(_c.DocTypeSpreadsheet, _c.OfficeDocumentType, _c.SharedStringsType, 0)

const (
	AnchorTypeAbsolute AnchorType = iota
	AnchorTypeOneCell
	AnchorTypeTwoCell
)

// Comments is the container for comments for a single sheet.
type Comments struct {
	_dfe *Workbook
	_eca *_ca.Comments
}

// SetHeight sets the height of the anchored object.
func (_bdc AbsoluteAnchor) SetHeight(h _ee.Distance) { _bdc._bdf.Ext.CyAttr = int64(h / _ee.EMU) }

// AddView adds a sheet view.
func (_cbef *Sheet) AddView() SheetView {
	if _cbef._fcbf.SheetViews == nil {
		_cbef._fcbf.SheetViews = _ca.NewCT_SheetViews()
	}
	_ced := _ca.NewCT_SheetView()
	_cbef._fcbf.SheetViews.SheetView = append(_cbef._fcbf.SheetViews.SheetView, _ced)
	return SheetView{_ced}
}

// SetIcons configures the rule as an icon scale, removing existing
// configuration.
func (_gbgf ConditionalFormattingRule) SetIcons() IconScale {
	_gbgf.clear()
	_gbgf.SetType(_ca.ST_CfTypeIconSet)
	_gbgf._bfgc.IconSet = _ca.NewCT_IconSet()
	_eagc := IconScale{_gbgf._bfgc.IconSet}
	_eagc.SetIcons(_ca.ST_IconSetType3TrafficLights1)
	return _eagc
}

// Workbook is the top level container item for a set of spreadsheets.
type Workbook struct {
	_ef.DocBase
	_geca         *_ca.Workbook
	StyleSheet    StyleSheet
	SharedStrings SharedStrings
	_ebaf         []*_ca.Comments
	_abcf         []*_ca.Worksheet
	_adfgb        []_ef.Relationships
	_eceef        _ef.Relationships
	_fee          []*_dc.Theme
	_eddg         []*_df.WsDr
	_fgd          []_ef.Relationships
	_gedac        []*_fe.Container
	_agcf         []*_gg.ChartSpace
	_gfgbd        []*_ca.Table
	_ffaab        string
}

// Cell creates or returns a cell given a cell reference of the form 'A10'
func (_gbda *Sheet) Cell(cellRef string) Cell {
	_fba, _edcb := _db.ParseCellReference(cellRef)
	if _edcb != nil {
		_c.Log("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063e\u006cl\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u003a\u0020\u0025\u0073", _edcb)
		return _gbda.AddRow().AddCell()
	}
	return _gbda.Row(_fba.RowIdx).Cell(_fba.Column)
}

// SetHidden marks the defined name as hidden.
func (_ede DefinedName) SetHidden(b bool) { _ede._gagf.HiddenAttr = _c.Bool(b) }

// Name returns the name of the defined name.
func (_bcd DefinedName) Name() string { return _bcd._gagf.NameAttr }

// AddImage adds an image with a paricular anchor type, returning an anchor to
// allow adusting the image size/position.
func (_fcb Drawing) AddImage(img _ef.ImageRef, at AnchorType) Anchor {
	_gecc := 0
	for _deegd, _bdeg := range _fcb._ffa.Images {
		if _bdeg == img {
			_gecc = _deegd + 1
			break
		}
	}
	var _dbae string
	for _acae, _gcc := range _fcb._ffa._eddg {
		if _gcc == _fcb._acgf {
			_aae := _ge.Sprintf("\u002e\u002e\u002f\u006ded\u0069\u0061\u002f\u0069\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", _gecc, img.Format())
			_cgb := _fcb._ffa._fgd[_acae].AddRelationship(_aae, _c.ImageType)
			_dbae = _cgb.ID()
			break
		}
	}
	var _bff Anchor
	var _beed *_df.CT_Picture
	switch at {
	case AnchorTypeAbsolute:
		_addbc := _dfcg()
		_fcb._acgf.EG_Anchor = append(_fcb._acgf.EG_Anchor, &_df.EG_Anchor{AbsoluteAnchor: _addbc})
		_addbc.Choice = &_df.EG_ObjectChoicesChoice{}
		_addbc.Choice.Pic = _df.NewCT_Picture()
		_addbc.Pos.XAttr.ST_CoordinateUnqualified = _c.Int64(0)
		_addbc.Pos.YAttr.ST_CoordinateUnqualified = _c.Int64(0)
		_beed = _addbc.Choice.Pic
		_bff = AbsoluteAnchor{_addbc}
	case AnchorTypeOneCell:
		_bdfb := _ggae()
		_fcb._acgf.EG_Anchor = append(_fcb._acgf.EG_Anchor, &_df.EG_Anchor{OneCellAnchor: _bdfb})
		_bdfb.Choice = &_df.EG_ObjectChoicesChoice{}
		_bdfb.Choice.Pic = _df.NewCT_Picture()
		_beed = _bdfb.Choice.Pic
		_bff = OneCellAnchor{_bdfb}
	case AnchorTypeTwoCell:
		_eef := _cecag()
		_fcb._acgf.EG_Anchor = append(_fcb._acgf.EG_Anchor, &_df.EG_Anchor{TwoCellAnchor: _eef})
		_eef.Choice = &_df.EG_ObjectChoicesChoice{}
		_eef.Choice.Pic = _df.NewCT_Picture()
		_beed = _eef.Choice.Pic
		_bff = TwoCellAnchor{_eef}
	}
	_beed.NvPicPr.CNvPr.IdAttr = uint32(len(_fcb._acgf.EG_Anchor))
	_beed.NvPicPr.CNvPr.NameAttr = "\u0049\u006d\u0061g\u0065"
	_beed.BlipFill.Blip = _dc.NewCT_Blip()
	_beed.BlipFill.Blip.EmbedAttr = _c.String(_dbae)
	_beed.BlipFill.Stretch = _dc.NewCT_StretchInfoProperties()
	_beed.SpPr = _dc.NewCT_ShapeProperties()
	_beed.SpPr.Xfrm = _dc.NewCT_Transform2D()
	_beed.SpPr.Xfrm.Off = _dc.NewCT_Point2D()
	_beed.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _c.Int64(0)
	_beed.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _c.Int64(0)
	_beed.SpPr.Xfrm.Ext = _dc.NewCT_PositiveSize2D()
	_beed.SpPr.Xfrm.Ext.CxAttr = int64(float64(img.Size().X*_ee.Pixel72) / _ee.EMU)
	_beed.SpPr.Xfrm.Ext.CyAttr = int64(float64(img.Size().Y*_ee.Pixel72) / _ee.EMU)
	_beed.SpPr.PrstGeom = _dc.NewCT_PresetGeometry2D()
	_beed.SpPr.PrstGeom.PrstAttr = _dc.ST_ShapeTypeRect
	_beed.SpPr.Ln = _dc.NewCT_LineProperties()
	_beed.SpPr.Ln.NoFill = _dc.NewCT_NoFillProperties()
	return _bff
}
func (_bgeec *Sheet) getAllCellsInFormulaArrays(_egeb bool) (map[string]bool, error) {
	_aafea := _gbf.NewEvaluator()
	_dec := _bgeec.FormulaContext()
	_ggbb := map[string]bool{}
	for _, _afce := range _bgeec.Rows() {
		for _, _abfc := range _afce.Cells() {
			if _abfc.X().F != nil {
				_daef := _abfc.X().F.Content
				if _abfc.X().F.TAttr == _ca.ST_CellFormulaTypeArray {
					_gebbe := _aafea.Eval(_dec, _daef).AsString()
					if _gebbe.Type == _gbf.ResultTypeError {
						_c.Log("\u0065\u0072\u0072o\u0072\u0020\u0065\u0076a\u0075\u006c\u0061\u0074\u0069\u006e\u0067 \u0066\u006f\u0072\u006d\u0075\u006c\u0061\u0020\u0025\u0073\u003a\u0020\u0025\u0073", _daef, _gebbe.ErrorMessage)
						_abfc.X().V = nil
					}
					if _gebbe.Type == _gbf.ResultTypeArray {
						_bbfd, _bfac := _db.ParseCellReference(_abfc.Reference())
						if _bfac != nil {
							return map[string]bool{}, _bfac
						}
						if (_egeb && len(_gebbe.ValueArray) == 1) || (!_egeb && len(_gebbe.ValueArray[0]) == 1) {
							continue
						}
						for _daegb, _caga := range _gebbe.ValueArray {
							_bedd := _bbfd.RowIdx + uint32(_daegb)
							for _eaff := range _caga {
								_ecfb := _db.IndexToColumn(_bbfd.ColumnIdx + uint32(_eaff))
								_ggbb[_ge.Sprintf("\u0025\u0073\u0025\u0064", _ecfb, _bedd)] = true
							}
						}
					} else if _gebbe.Type == _gbf.ResultTypeList {
						_deg, _gdag := _db.ParseCellReference(_abfc.Reference())
						if _gdag != nil {
							return map[string]bool{}, _gdag
						}
						if _egeb || len(_gebbe.ValueList) == 1 {
							continue
						}
						_bdae := _deg.RowIdx
						for _ccg := range _gebbe.ValueList {
							_gead := _db.IndexToColumn(_deg.ColumnIdx + uint32(_ccg))
							_ggbb[_ge.Sprintf("\u0025\u0073\u0025\u0064", _gead, _bdae)] = true
						}
					}
				}
			}
		}
	}
	return _ggbb, nil
}
func (_dcbfa *Sheet) setArray(_dgff string, _ccec _gbf.Result) error {
	_daeg, _ggec := _db.ParseCellReference(_dgff)
	if _ggec != nil {
		return _ggec
	}
	for _gede, _dadb := range _ccec.ValueArray {
		_cgbg := _dcbfa.Row(_daeg.RowIdx + uint32(_gede))
		for _dddd, _cbaba := range _dadb {
			_fbge := _cgbg.Cell(_db.IndexToColumn(_daeg.ColumnIdx + uint32(_dddd)))
			if _cbaba.Type != _gbf.ResultTypeEmpty {
				if _cbaba.IsBoolean {
					_fbge.SetBool(_cbaba.ValueNumber != 0)
				} else {
					_fbge.SetCachedFormulaResult(_cbaba.String())
				}
			}
		}
	}
	return nil
}

// SetFill applies a fill to a cell style.  The fill is referenced by its index
// so modifying the fill afterward will affect all styles that reference it.
func (_ecdd CellStyle) SetFill(f Fill) {
	_ecdd._bee.FillIdAttr = _c.Uint32(f.Index())
	_ecdd._bee.ApplyFillAttr = _c.Bool(true)
}

// Cell is a single cell within a sheet.
type Cell struct {
	_adb *Workbook
	_ade *Sheet
	_gbd *_ca.CT_Row
	_ec  *_ca.CT_Cell
}

func (_dbcc *evalContext) Cell(ref string, ev _gbf.Evaluator) _gbf.Result {
	if !_dda(ref) {
		return _gbf.MakeErrorResultType(_gbf.ErrorTypeName, "")
	}
	_efde := _dbcc._gag.Name() + "\u0021" + ref
	if _cfcd, _caa := ev.GetFromCache(_efde); _caa {
		return _cfcd
	}
	_agag, _geg := _db.ParseCellReference(ref)
	if _geg != nil {
		return _gbf.MakeErrorResult(_ge.Sprintf("e\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", ref, _geg))
	}
	if _dbcc._dbc != 0 && !_agag.AbsoluteColumn {
		_agag.ColumnIdx += _dbcc._dbc
		_agag.Column = _db.IndexToColumn(_agag.ColumnIdx)
	}
	if _dbcc._gdbc != 0 && !_agag.AbsoluteRow {
		_agag.RowIdx += _dbcc._gdbc
	}
	_gfg := _dbcc._gag.Cell(_agag.String())
	if _gfg.HasFormula() {
		if _, _eecd := _dbcc._dce[ref]; _eecd {
			return _gbf.MakeErrorResult("r\u0065\u0063\u0075\u0072\u0073\u0069\u006f\u006e\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0065\u0064\u0020d\u0075\u0072\u0069\u006e\u0067\u0020\u0065\u0076\u0061\u006cua\u0074\u0069\u006fn\u0020o\u0066\u0020" + ref)
		}
		_dbcc._dce[ref] = struct{}{}
		_aeg := ev.Eval(_dbcc, _gfg.GetFormula())
		delete(_dbcc._dce, ref)
		ev.SetCache(_efde, _aeg)
		return _aeg
	}
	if _gfg.IsEmpty() {
		_ceaf := _gbf.MakeEmptyResult()
		ev.SetCache(_efde, _ceaf)
		return _ceaf
	} else if _gfg.IsNumber() {
		_dac, _ := _gfg.GetValueAsNumber()
		_acfd := _gbf.MakeNumberResult(_dac)
		ev.SetCache(_efde, _acfd)
		return _acfd
	} else if _gfg.IsBool() {
		_dff, _ := _gfg.GetValueAsBool()
		_gbgb := _gbf.MakeBoolResult(_dff)
		ev.SetCache(_efde, _gbgb)
		return _gbgb
	}
	_dba, _ := _gfg.GetRawValue()
	if _gfg.IsError() {
		_cdcd := _gbf.MakeErrorResult("")
		_cdcd.ValueString = _dba
		ev.SetCache(_efde, _cdcd)
		return _cdcd
	}
	_addb := _gbf.MakeStringResult(_dba)
	ev.SetCache(_efde, _addb)
	return _addb
}

// SetHeightCells sets the height the anchored object by moving the bottom.  It
// is not compatible with SetHeight.
func (_ddgc TwoCellAnchor) SetHeightCells(h int32) {
	_ddgc.SetHeight(0)
	_abfg := _ddgc.TopLeft()
	_bfgf := _ddgc.BottomRight()
	_bfgf.SetRow(_abfg.Row() + h)
}

// Column represents a column within a sheet. It's only used for formatting
// purposes, so it's possible to construct a sheet without configuring columns.
type Column struct{ _gfd *_ca.CT_Col }

// AddFont adds a new empty font to the stylesheet.
func (_ffg StyleSheet) AddFont() Font {
	_cbag := _ca.NewCT_Font()
	_ffg._bbbe.Fonts.Font = append(_ffg._bbbe.Fonts.Font, _cbag)
	_ffg._bbbe.Fonts.CountAttr = _c.Uint32(uint32(len(_ffg._bbbe.Fonts.Font)))
	return Font{_cbag, _ffg._bbbe}
}

// SetHidden marks the defined name as hidden.
func (_aag DefinedName) SetLocalSheetID(id uint32) { _aag._gagf.LocalSheetIdAttr = _c.Uint32(id) }

// IsNumber returns true if the cell is a number type cell.
func (_gdb Cell) IsNumber() bool {
	switch _gdb._ec.TAttr {
	case _ca.ST_CellTypeN:
		return true
	case _ca.ST_CellTypeS, _ca.ST_CellTypeB:
		return false
	}
	return _gdb._ec.V != nil && _ed.IsNumber(*_gdb._ec.V)
}
func _cecag() *_df.CT_TwoCellAnchor {
	_dca := _df.NewCT_TwoCellAnchor()
	_dca.EditAsAttr = _df.ST_EditAsOneCell
	_dca.From.Col = 5
	_dca.From.Row = 0
	_dca.From.ColOff.ST_CoordinateUnqualified = _c.Int64(0)
	_dca.From.RowOff.ST_CoordinateUnqualified = _c.Int64(0)
	_dca.To.Col = 10
	_dca.To.Row = 20
	_dca.To.ColOff.ST_CoordinateUnqualified = _c.Int64(0)
	_dca.To.RowOff.ST_CoordinateUnqualified = _c.Int64(0)
	return _dca
}

// GetFormattedValue returns the formatted cell value as it would appear in
// Excel. This involves determining the format string to apply, parsing it, and
// then formatting the value according to the format string.  This should only
// be used if you care about replicating what Excel would show, otherwise
// GetValueAsNumber()/GetValueAsTime
func (_aef Cell) GetFormattedValue() string {
	_cge := _aef.getFormat()
	switch _aef._ec.TAttr {
	case _ca.ST_CellTypeB:
		_eag, _ := _aef.GetValueAsBool()
		if _eag {
			return "\u0054\u0052\u0055\u0045"
		}
		return "\u0046\u0041\u004cS\u0045"
	case _ca.ST_CellTypeN:
		_dbb, _ := _aef.GetValueAsNumber()
		return _ed.Number(_dbb, _cge)
	case _ca.ST_CellTypeE:
		if _aef._ec.V != nil {
			return *_aef._ec.V
		}
		return ""
	case _ca.ST_CellTypeS, _ca.ST_CellTypeInlineStr:
		return _ed.String(_aef.GetString(), _cge)
	case _ca.ST_CellTypeStr:
		_fbeb := _aef.GetString()
		if _ed.IsNumber(_fbeb) {
			_bag, _ := _bd.ParseFloat(_fbeb, 64)
			return _ed.Number(_bag, _cge)
		}
		return _ed.String(_fbeb, _cge)
	case _ca.ST_CellTypeUnset:
		fallthrough
	default:
		_dfc, _ := _aef.GetRawValue()
		if len(_dfc) == 0 {
			return ""
		}
		_cgd, _abe := _aef.GetValueAsNumber()
		if _abe == nil {
			return _ed.Number(_cgd, _cge)
		}
		return _ed.String(_dfc, _cge)
	}
}
func (_dbbc DataValidation) SetComparison(t DVCompareType, op DVCompareOp) DataValidationCompare {
	_dbbc.clear()
	_dbbc._cecf.TypeAttr = _ca.ST_DataValidationType(t)
	_dbbc._cecf.OperatorAttr = _ca.ST_DataValidationOperator(op)
	return DataValidationCompare{_dbbc._cecf}
}

// AddCellStyle adds a new empty cell style to the stylesheet.
func (_faeab StyleSheet) AddCellStyle() CellStyle {
	_gccb := _ca.NewCT_Xf()
	_faeab._bbbe.CellXfs.Xf = append(_faeab._bbbe.CellXfs.Xf, _gccb)
	_faeab._bbbe.CellXfs.CountAttr = _c.Uint32(uint32(len(_faeab._bbbe.CellXfs.Xf)))
	return CellStyle{_faeab._aba, _gccb, _faeab._bbbe.CellXfs}
}

// AddDifferentialStyle adds a new empty differential cell style to the stylesheet.
func (_fddd StyleSheet) AddDifferentialStyle() DifferentialStyle {
	if _fddd._bbbe.Dxfs == nil {
		_fddd._bbbe.Dxfs = _ca.NewCT_Dxfs()
	}
	_fedb := _ca.NewCT_Dxf()
	_fddd._bbbe.Dxfs.Dxf = append(_fddd._bbbe.Dxfs.Dxf, _fedb)
	_fddd._bbbe.Dxfs.CountAttr = _c.Uint32(uint32(len(_fddd._bbbe.Dxfs.Dxf)))
	return DifferentialStyle{_fedb, _fddd._aba, _fddd._bbbe.Dxfs}
}

// LockObject controls the locking of the sheet objects.
func (_afgeb SheetProtection) LockObject(b bool) {
	if !b {
		_afgeb._aebb.ObjectsAttr = nil
	} else {
		_afgeb._aebb.ObjectsAttr = _c.Bool(true)
	}
}
func (_dbfg *Sheet) updateAfterRemove(_aaec uint32, _aeeb _aac.UpdateAction) error {
	_cdaa := _dbfg.Name()
	_befa := &_aac.UpdateQuery{UpdateType: _aeeb, ColumnIdx: _aaec, SheetToUpdate: _cdaa}
	for _, _bageb := range _dbfg._ccae.Sheets() {
		_befa.UpdateCurrentSheet = _cdaa == _bageb.Name()
		for _, _gacc := range _bageb.Rows() {
			for _, _abfd := range _gacc.Cells() {
				if _abfd.X().F != nil {
					_geff := _abfd.X().F.Content
					_cddd := _gbf.ParseString(_geff)
					if _cddd == nil {
						_abfd.SetError("\u0023\u0052\u0045F\u0021")
					} else {
						_daba := _cddd.Update(_befa)
						_abfd.X().F.Content = _ge.Sprintf("\u003d\u0025\u0073", _daba.String())
					}
				}
			}
		}
	}
	return nil
}

// GetValueAsBool retrieves the cell's value as a boolean
func (_add Cell) GetValueAsBool() (bool, error) {
	if _add._ec.TAttr != _ca.ST_CellTypeB {
		return false, _gb.New("\u0063e\u006c\u006c\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006ff\u0020\u0062\u006f\u006f\u006c\u0020\u0074\u0079\u0070\u0065")
	}
	if _add._ec.V == nil {
		return false, _gb.New("\u0063\u0065\u006c\u006c\u0020\u0068\u0061\u0073\u0020\u006e\u006f\u0020v\u0061\u006c\u0075\u0065")
	}
	return _bd.ParseBool(*_add._ec.V)
}

// CellMarker represents a cell position
type CellMarker struct{ _ggg *_df.CT_Marker }

// IsBool returns true if the cell boolean value.
func (_gdfd *evalContext) IsBool(cellRef string) bool { return _gdfd._gag.Cell(cellRef).IsBool() }

type ConditionalFormattingRule struct{ _bfgc *_ca.CT_CfRule }

// TopLeft returns the top-left corner of the anchored object.
func (_ffc OneCellAnchor) TopLeft() CellMarker { return CellMarker{_ffc._bbbd.From} }

// AddChart adds an chart to a drawing, returning the chart and an anchor that
// can be used to position the chart within the sheet.
func (_cgcg Drawing) AddChart(at AnchorType) (_agg.Chart, Anchor) {
	_dbg := _gg.NewChartSpace()
	_cgcg._ffa._agcf = append(_cgcg._ffa._agcf, _dbg)
	_agga := _c.AbsoluteFilename(_c.DocTypeSpreadsheet, _c.ChartContentType, len(_cgcg._ffa._agcf))
	_cgcg._ffa.ContentTypes.AddOverride(_agga, _c.ChartContentType)
	var _bafe string
	for _ggc, _cefb := range _cgcg._ffa._eddg {
		if _cefb == _cgcg._acgf {
			_ffe := _c.RelativeFilename(_c.DocTypeSpreadsheet, _c.DrawingType, _c.ChartType, len(_cgcg._ffa._agcf))
			_dgd := _cgcg._ffa._fgd[_ggc].AddRelationship(_ffe, _c.ChartType)
			_bafe = _dgd.ID()
			break
		}
	}
	var _fdca Anchor
	var _bfbb *_df.CT_GraphicalObjectFrame
	switch at {
	case AnchorTypeAbsolute:
		_cfgc := _dfcg()
		_cgcg._acgf.EG_Anchor = append(_cgcg._acgf.EG_Anchor, &_df.EG_Anchor{AbsoluteAnchor: _cfgc})
		_cfgc.Choice = &_df.EG_ObjectChoicesChoice{}
		_cfgc.Choice.GraphicFrame = _df.NewCT_GraphicalObjectFrame()
		_bfbb = _cfgc.Choice.GraphicFrame
		_fdca = AbsoluteAnchor{_cfgc}
	case AnchorTypeOneCell:
		_gbab := _ggae()
		_cgcg._acgf.EG_Anchor = append(_cgcg._acgf.EG_Anchor, &_df.EG_Anchor{OneCellAnchor: _gbab})
		_gbab.Choice = &_df.EG_ObjectChoicesChoice{}
		_gbab.Choice.GraphicFrame = _df.NewCT_GraphicalObjectFrame()
		_bfbb = _gbab.Choice.GraphicFrame
		_fdca = OneCellAnchor{_gbab}
	case AnchorTypeTwoCell:
		_eed := _cecag()
		_cgcg._acgf.EG_Anchor = append(_cgcg._acgf.EG_Anchor, &_df.EG_Anchor{TwoCellAnchor: _eed})
		_eed.Choice = &_df.EG_ObjectChoicesChoice{}
		_eed.Choice.GraphicFrame = _df.NewCT_GraphicalObjectFrame()
		_bfbb = _eed.Choice.GraphicFrame
		_fdca = TwoCellAnchor{_eed}
	}
	_bfbb.NvGraphicFramePr = _df.NewCT_GraphicalObjectFrameNonVisual()
	_bfbb.NvGraphicFramePr.CNvPr.IdAttr = uint32(len(_cgcg._acgf.EG_Anchor))
	_bfbb.NvGraphicFramePr.CNvPr.NameAttr = "\u0043\u0068\u0061r\u0074"
	_bfbb.Graphic = _dc.NewGraphic()
	_bfbb.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073.\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006dl/\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074"
	_caae := _gg.NewChart()
	_caae.IdAttr = _bafe
	_bfbb.Graphic.GraphicData.Any = []_c.Any{_caae}
	_eab := _agg.MakeChart(_dbg)
	_eab.Properties().SetSolidFill(_ae.White)
	_eab.SetDisplayBlanksAs(_gg.ST_DispBlanksAsGap)
	return _eab, _fdca
}

// TopLeft returns the CellMaker for the top left corner of the anchor.
func (_bdfee TwoCellAnchor) TopLeft() CellMarker { return CellMarker{_bdfee._aeea.From} }

// AddDataValidation adds a data validation rule to a sheet.
func (_fecc *Sheet) AddDataValidation() DataValidation {
	if _fecc._fcbf.DataValidations == nil {
		_fecc._fcbf.DataValidations = _ca.NewCT_DataValidations()
	}
	_egbc := _ca.NewCT_DataValidation()
	_egbc.ShowErrorMessageAttr = _c.Bool(true)
	_fecc._fcbf.DataValidations.DataValidation = append(_fecc._fcbf.DataValidations.DataValidation, _egbc)
	_fecc._fcbf.DataValidations.CountAttr = _c.Uint32(uint32(len(_fecc._fcbf.DataValidations.DataValidation)))
	return DataValidation{_egbc}
}

// SetBorder is a helper function for creating borders across multiple cells. In
// the OOXML spreadsheet format, a border applies to a single cell.  To draw a
// 'boxed' border around multiple cells, you need to apply different styles to
// the cells on the top,left,right,bottom and four corners.  This function
// breaks apart a single border into its components and applies it to cells as
// needed to give the effect of a border applying to multiple cells.
func (_fcd *Sheet) SetBorder(cellRange string, border Border) error {
	_gbbd, _efg, _dcgd := _db.ParseRangeReference(cellRange)
	if _dcgd != nil {
		return _dcgd
	}
	_cdcb := _fcd._ccae.StyleSheet.AddCellStyle()
	_gbbf := _fcd._ccae.StyleSheet.AddBorder()
	_cdcb.SetBorder(_gbbf)
	_gbbf._eac.Top = border._eac.Top
	_gbbf._eac.Left = border._eac.Left
	_gdd := _fcd._ccae.StyleSheet.AddCellStyle()
	_aaeg := _fcd._ccae.StyleSheet.AddBorder()
	_gdd.SetBorder(_aaeg)
	_aaeg._eac.Top = border._eac.Top
	_aaeg._eac.Right = border._eac.Right
	_bbff := _fcd._ccae.StyleSheet.AddCellStyle()
	_ccaf := _fcd._ccae.StyleSheet.AddBorder()
	_bbff.SetBorder(_ccaf)
	_ccaf._eac.Top = border._eac.Top
	_ecbb := _fcd._ccae.StyleSheet.AddCellStyle()
	_gfdg := _fcd._ccae.StyleSheet.AddBorder()
	_ecbb.SetBorder(_gfdg)
	_gfdg._eac.Left = border._eac.Left
	_dfdea := _fcd._ccae.StyleSheet.AddCellStyle()
	_gbbc := _fcd._ccae.StyleSheet.AddBorder()
	_dfdea.SetBorder(_gbbc)
	_gbbc._eac.Right = border._eac.Right
	_fgcaf := _fcd._ccae.StyleSheet.AddCellStyle()
	_acfg := _fcd._ccae.StyleSheet.AddBorder()
	_fgcaf.SetBorder(_acfg)
	_acfg._eac.Bottom = border._eac.Bottom
	_ged := _fcd._ccae.StyleSheet.AddCellStyle()
	_efffe := _fcd._ccae.StyleSheet.AddBorder()
	_ged.SetBorder(_efffe)
	_efffe._eac.Bottom = border._eac.Bottom
	_efffe._eac.Left = border._eac.Left
	_afdb := _fcd._ccae.StyleSheet.AddCellStyle()
	_caf := _fcd._ccae.StyleSheet.AddBorder()
	_afdb.SetBorder(_caf)
	_caf._eac.Bottom = border._eac.Bottom
	_caf._eac.Right = border._eac.Right
	_cbdd := _gbbd.RowIdx
	_dgfe := _gbbd.ColumnIdx
	_ddfeg := _efg.RowIdx
	_ddab := _efg.ColumnIdx
	for _dcfg := _cbdd; _dcfg <= _ddfeg; _dcfg++ {
		for _bbfgb := _dgfe; _bbfgb <= _ddab; _bbfgb++ {
			_abcg := _ge.Sprintf("\u0025\u0073\u0025\u0064", _db.IndexToColumn(_bbfgb), _dcfg)
			switch {
			case _dcfg == _cbdd && _bbfgb == _dgfe:
				_fcd.Cell(_abcg).SetStyle(_cdcb)
			case _dcfg == _cbdd && _bbfgb == _ddab:
				_fcd.Cell(_abcg).SetStyle(_gdd)
			case _dcfg == _ddfeg && _bbfgb == _dgfe:
				_fcd.Cell(_abcg).SetStyle(_ged)
			case _dcfg == _ddfeg && _bbfgb == _ddab:
				_fcd.Cell(_abcg).SetStyle(_afdb)
			case _dcfg == _cbdd:
				_fcd.Cell(_abcg).SetStyle(_bbff)
			case _dcfg == _ddfeg:
				_fcd.Cell(_abcg).SetStyle(_fgcaf)
			case _bbfgb == _dgfe:
				_fcd.Cell(_abcg).SetStyle(_ecbb)
			case _bbfgb == _ddab:
				_fcd.Cell(_abcg).SetStyle(_dfdea)
			}
		}
	}
	return nil
}

// ClearFont clears any font configuration from the cell style.
func (_bbg CellStyle) ClearFont() { _bbg._bee.FontIdAttr = nil; _bbg._bee.ApplyFontAttr = nil }
func _bdca(_adac string, _debc uint32, _ffag bool) string {
	_eeb, _fdfa, _gedb := _db.ParseRangeReference(_adac)
	if _gedb == nil {
		_gfdef, _bcf := _eeb.ColumnIdx, _fdfa.ColumnIdx
		if _debc >= _gfdef && _debc <= _bcf {
			if _gfdef == _bcf {
				if _ffag {
					return ""
				} else {
					return _adac
				}
			} else {
				_efaf := _fdfa.Update(_aac.UpdateActionRemoveColumn)
				return _ge.Sprintf("\u0025\u0073\u003a%\u0073", _eeb.String(), _efaf.String())
			}
		} else if _debc < _gfdef {
			_fcgg := _eeb.Update(_aac.UpdateActionRemoveColumn)
			_cafd := _fdfa.Update(_aac.UpdateActionRemoveColumn)
			return _ge.Sprintf("\u0025\u0073\u003a%\u0073", _fcgg.String(), _cafd.String())
		}
	} else {
		_fbfeg, _dggca, _dfefg := _db.ParseColumnRangeReference(_adac)
		if _dfefg != nil {
			return ""
		}
		_agda, _bdea := _fbfeg.ColumnIdx, _dggca.ColumnIdx
		if _debc >= _agda && _debc <= _bdea {
			if _agda == _bdea {
				if _ffag {
					return ""
				} else {
					return _adac
				}
			} else {
				_befc := _dggca.Update(_aac.UpdateActionRemoveColumn)
				return _ge.Sprintf("\u0025\u0073\u003a%\u0073", _fbfeg.String(), _befc.String())
			}
		} else if _debc < _agda {
			_cafb := _fbfeg.Update(_aac.UpdateActionRemoveColumn)
			_dbac := _dggca.Update(_aac.UpdateActionRemoveColumn)
			return _ge.Sprintf("\u0025\u0073\u003a%\u0073", _cafb.String(), _dbac.String())
		}
	}
	return ""
}

type Table struct{ _agbc *_ca.Table }


func (_aedb *Sheet) RecalculateFormulas() {
	_afga := _gbf.NewEvaluator()
	_fafa := _aedb.FormulaContext()
	for _, _cdga := range _aedb.Rows() {
		for _, _gfca := range _cdga.Cells() {
			if _gfca.X().F != nil {
				_gef := _gfca.X().F.Content
				if _gfca.X().F.TAttr == _ca.ST_CellFormulaTypeShared && len(_gef) == 0 {
					continue
				}
				_gfae := _afga.Eval(_fafa, _gef).AsString()
				if _gfae.Type == _gbf.ResultTypeError {
					_c.Log("\u0065\u0072\u0072o\u0072\u0020\u0065\u0076a\u0075\u006c\u0061\u0074\u0069\u006e\u0067 \u0066\u006f\u0072\u006d\u0075\u006c\u0061\u0020\u0025\u0073\u003a\u0020\u0025\u0073", _gef, _gfae.ErrorMessage)
					_gfca.X().V = nil
				} else {
					if _gfae.Type == _gbf.ResultTypeNumber {
						_gfca.X().TAttr = _ca.ST_CellTypeN
					} else {
						_gfca.X().TAttr = _ca.ST_CellTypeInlineStr
					}
					_gfca.X().V = _c.String(_gfae.Value())
					if _gfca.X().F.TAttr == _ca.ST_CellFormulaTypeArray {
						if _gfae.Type == _gbf.ResultTypeArray {
							_aedb.setArray(_gfca.Reference(), _gfae)
						} else if _gfae.Type == _gbf.ResultTypeList {
							_aedb.setList(_gfca.Reference(), _gfae)
						}
					} else if _gfca.X().F.TAttr == _ca.ST_CellFormulaTypeShared && _gfca.X().F.RefAttr != nil {
						_cfaf, _bcdbe, _gbgfb := _db.ParseRangeReference(*_gfca.X().F.RefAttr)
						if _gbgfb != nil {
							_fb.Printf("\u0065\u0072r\u006f\u0072\u0020\u0069n\u0020\u0073h\u0061\u0072\u0065\u0064\u0020\u0066\u006f\u0072m\u0075\u006c\u0061\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u003a\u0020\u0025\u0073", _gbgfb)
							continue
						}
						_aedb.setShared(_gfca.Reference(), _cfaf, _bcdbe, _gef)
					}
				}
			}
		}
	}
}

// SetPattern sets the pattern of the fill.
func (_fgcf PatternFill) SetPattern(p _ca.ST_PatternType) { _fgcf._fbdg.PatternTypeAttr = p }

// Read reads a workbook from an io.Reader(.xlsx).
func Read(r _fd.ReaderAt, size int64) (*Workbook, error) {
	_dgdg := New()
	_ffbf, _fgfb := _cf.TempDir("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0078\u006c\u0073\u0078")
	if _fgfb != nil {
		return nil, _fgfb
	}
	_dgdg.TmpPath = _ffbf
	_bgef, _fgfb := _ab.NewReader(r, size)
	if _fgfb != nil {
		return nil, _ge.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073", _fgfb)
	}
	_egffa := []*_ab.File{}
	_egffa = append(_egffa, _bgef.File...)
	_gcb := false
	for _, _dfbg := range _egffa {
		if _dfbg.FileHeader.Name == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
			_gcb = true
			break
		}
	}
	if _gcb {
		_dgdg.CreateCustomProperties()
	}
	_cbcg := _cc.DecodeMap{}
	_cbcg.SetOnNewRelationshipFunc(_dgdg.onNewRelationship)
	_cbcg.AddTarget(_c.ContentTypesFilename, _dgdg.ContentTypes.X(), "", 0)
	_cbcg.AddTarget(_c.BaseRelsFilename, _dgdg.Rels.X(), "", 0)
	if _cgeg := _cbcg.Decode(_egffa); _cgeg != nil {
		return nil, _cgeg
	}
	for _, _daee := range _egffa {
		if _daee == nil {
			continue
		}
		if _bgc := _dgdg.AddExtraFileFromZip(_daee); _bgc != nil {
			return nil, _bgc
		}
	}
	if _gcb {
		_dcdag := false
		for _, _babd := range _dgdg.Rels.X().Relationship {
			if _babd.TargetAttr == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
				_dcdag = true
				break
			}
		}
		if !_dcdag {
			_dgdg.AddCustomRelationships()
		}
	}
	return _dgdg, nil
}

// X returns the inner wrapped XML type.
func (_ddf DifferentialStyle) X() *_ca.CT_Dxf { return _ddf._dgea }

// SetContent sets the defined name content.
func (_gdgb DefinedName) SetContent(s string) { _gdgb._gagf.Content = s }


func (_abefb *Workbook) RecalculateFormulas() {
	for _, _accd := range _abefb.Sheets() {
		_accd.RecalculateFormulas()
	}
}

// SetStyleIndex directly sets a style index to the cell.  This should only be
// called with an index retrieved from CellStyle.Index()
func (_afa Cell) SetStyleIndex(idx uint32) { _afa._ec.SAttr = _c.Uint32(idx) }
func _cgc(_dgc _fbe.Time) _fbe.Time {
	_dgc = _dgc.Local()
	return _fbe.Date(_dgc.Year(), _dgc.Month(), _dgc.Day(), _dgc.Hour(), _dgc.Minute(), _dgc.Second(), _dgc.Nanosecond(), _fbe.UTC)
}

// GetWidth returns a worksheet's column width.
func (_aedd *evalContext) GetWidth(colIdx int) float64 {
	colIdx++
	for _, _bcce := range _aedd._gag.X().Cols[0].Col {
		if int(_bcce.MinAttr) <= colIdx && colIdx <= int(_bcce.MaxAttr) {
			return float64(int(*_bcce.WidthAttr))
		}
	}
	return 0
}

// X returns the inner wrapped XML type.
func (_edf Cell) X() *_ca.CT_Cell { return _edf._ec }

// SetPassword sets the password hash to a hash of the input password.
func (_daae WorkbookProtection) SetPassword(pw string) { _daae.SetPasswordHash(PasswordHash(pw)) }
func (_cfaa Sheet) ExtentsIndex() (string, uint32, string, uint32) {
	var _aafd, _dagc, _aaba, _gfbc uint32 = 1, 1, 0, 0
	for _, _dfga := range _cfaa.Rows() {
		if _dfga.RowNumber() < _aafd {
			_aafd = _dfga.RowNumber()
		} else if _dfga.RowNumber() > _dagc {
			_dagc = _dfga.RowNumber()
		}
		for _, _bfd := range _dfga.Cells() {
			_ecgf, _dcf := _db.ParseCellReference(_bfd.Reference())
			if _dcf == nil {
				if _ecgf.ColumnIdx < _aaba {
					_aaba = _ecgf.ColumnIdx
				} else if _ecgf.ColumnIdx > _gfbc {
					_gfbc = _ecgf.ColumnIdx
				}
			}
		}
	}
	return _db.IndexToColumn(_aaba), _aafd, _db.IndexToColumn(_gfbc), _dagc
}

// SetPassword sets the password hash to a hash of the input password.
func (_bcbd SheetProtection) SetPassword(pw string) { _bcbd.SetPasswordHash(PasswordHash(pw)) }

// SheetViews returns the sheet views defined.  This is where splits and frozen
// rows/cols are configured.  Multiple sheet views are allowed, but I'm not
// aware of there being a use for more than a single sheet view.
func (_cbcf *Sheet) SheetViews() []SheetView {
	if _cbcf._fcbf.SheetViews == nil {
		return nil
	}
	_aadg := []SheetView{}
	for _, _cggb := range _cbcf._fcbf.SheetViews.SheetView {
		_aadg = append(_aadg, SheetView{_cggb})
	}
	return _aadg
}
func CreateDefaultNumberFormat(id StandardFormat) NumberFormat {
	_adbc := NumberFormat{_geeb: _ca.NewCT_NumFmt()}
	_adbc._geeb.NumFmtIdAttr = uint32(id)
	_adbc._geeb.FormatCodeAttr = "\u0047e\u006e\u0065\u0072\u0061\u006c"
	switch id {
	case StandardFormat0:
		_adbc._geeb.FormatCodeAttr = "\u0047e\u006e\u0065\u0072\u0061\u006c"
	case StandardFormat1:
		_adbc._geeb.FormatCodeAttr = "\u0030"
	case StandardFormat2:
		_adbc._geeb.FormatCodeAttr = "\u0030\u002e\u0030\u0030"
	case StandardFormat3:
		_adbc._geeb.FormatCodeAttr = "\u0023\u002c\u0023#\u0030"
	case StandardFormat4:
		_adbc._geeb.FormatCodeAttr = "\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030"
	case StandardFormat9:
		_adbc._geeb.FormatCodeAttr = "\u0030\u0025"
	case StandardFormat10:
		_adbc._geeb.FormatCodeAttr = "\u0030\u002e\u00300\u0025"
	case StandardFormat11:
		_adbc._geeb.FormatCodeAttr = "\u0030\u002e\u0030\u0030\u0045\u002b\u0030\u0030"
	case StandardFormat12:
		_adbc._geeb.FormatCodeAttr = "\u0023\u0020\u003f/\u003f"
	case StandardFormat13:
		_adbc._geeb.FormatCodeAttr = "\u0023 \u003f\u003f\u002f\u003f\u003f"
	case StandardFormat14:
		_adbc._geeb.FormatCodeAttr = "\u006d\u002f\u0064\u002f\u0079\u0079"
	case StandardFormat15:
		_adbc._geeb.FormatCodeAttr = "\u0064\u002d\u006d\u006d\u006d\u002d\u0079\u0079"
	case StandardFormat16:
		_adbc._geeb.FormatCodeAttr = "\u0064\u002d\u006dm\u006d"
	case StandardFormat17:
		_adbc._geeb.FormatCodeAttr = "\u006d\u006d\u006d\u002d\u0079\u0079"
	case StandardFormat18:
		_adbc._geeb.FormatCodeAttr = "\u0068\u003a\u006d\u006d\u0020\u0041\u004d\u002f\u0050\u004d"
	case StandardFormat19:
		_adbc._geeb.FormatCodeAttr = "\u0068\u003a\u006d\u006d\u003a\u0073\u0073\u0020\u0041\u004d\u002f\u0050\u004d"
	case StandardFormat20:
		_adbc._geeb.FormatCodeAttr = "\u0068\u003a\u006d\u006d"
	case StandardFormat21:
		_adbc._geeb.FormatCodeAttr = "\u0068:\u006d\u006d\u003a\u0073\u0073"
	case StandardFormat22:
		_adbc._geeb.FormatCodeAttr = "m\u002f\u0064\u002f\u0079\u0079\u0020\u0068\u003a\u006d\u006d"
	case StandardFormat37:
		_adbc._geeb.FormatCodeAttr = "\u0023\u002c\u0023\u0023\u0030\u0020\u003b\u0028\u0023,\u0023\u0023\u0030\u0029"
	case StandardFormat38:
		_adbc._geeb.FormatCodeAttr = "\u0023\u002c\u0023\u00230 \u003b\u005b\u0052\u0065\u0064\u005d\u0028\u0023\u002c\u0023\u0023\u0030\u0029"
	case StandardFormat39:
		_adbc._geeb.FormatCodeAttr = "\u0023\u002c\u0023\u00230.\u0030\u0030\u003b\u0028\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030\u0029"
	case StandardFormat40:
		_adbc._geeb.FormatCodeAttr = "\u0023,\u0023\u0023\u0030\u002e\u0030\u0030\u003b\u005b\u0052\u0065\u0064]\u0028\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030\u0029"
	case StandardFormat45:
		_adbc._geeb.FormatCodeAttr = "\u006d\u006d\u003as\u0073"
	case StandardFormat46:
		_adbc._geeb.FormatCodeAttr = "\u005bh\u005d\u003a\u006d\u006d\u003a\u0073s"
	case StandardFormat47:
		_adbc._geeb.FormatCodeAttr = "\u006dm\u003a\u0073\u0073\u002e\u0030"
	case StandardFormat48:
		_adbc._geeb.FormatCodeAttr = "\u0023\u0023\u0030\u002e\u0030\u0045\u002b\u0030"
	case StandardFormat49:
		_adbc._geeb.FormatCodeAttr = "\u0040"
	}
	return _adbc
}

// X returns the inner wrapped XML type.
func (_egac DataBarScale) X() *_ca.CT_DataBar { return _egac._egfb }

// Comments returns the list of comments for this sheet
func (_bgf Comments) Comments() []Comment {
	_bafb := []Comment{}
	for _, _bgfd := range _bgf._eca.CommentList.Comment {
		_bafb = append(_bafb, Comment{_bgf._dfe, _bgfd, _bgf._eca})
	}
	return _bafb
}

// ClearNumberFormat removes any number formatting from the style.
func (_dcg CellStyle) ClearNumberFormat() {
	_dcg._bee.NumFmtIdAttr = nil
	_dcg._bee.ApplyNumberFormatAttr = nil
}

// SetFont applies a font to a cell style.  The font is referenced by its
// index so modifying the font afterward will affect all styles that reference
// it.
func (_ebf CellStyle) SetFont(f Font) {
	_ebf._bee.FontIdAttr = _c.Uint32(f.Index())
	_ebf._bee.ApplyFontAttr = _c.Bool(true)
}

// Row is a row within a spreadsheet.
type Row struct {
	_bgb  *Workbook
	_dcab *Sheet
	_gfcg *_ca.CT_Row
}

// IsWindowLocked returns whether the workbook windows are locked.
func (_eddf WorkbookProtection) IsWindowLocked() bool {
	return _eddf._eacd.LockWindowsAttr != nil && *_eddf._eacd.LockWindowsAttr
}

// IsSheetLocked returns whether the sheet is locked.
func (_cfgf SheetProtection) IsSheetLocked() bool {
	return _cfgf._aebb.SheetAttr != nil && *_cfgf._aebb.SheetAttr
}

// Content returns the content of the defined range (the range in most cases)/
func (_cef DefinedName) Content() string { return _cef._gagf.Content }

// AddNamedCell adds a new named cell to a row and returns it. You should
// normally prefer Cell() as it will return the existing cell if the cell
// already exists, while AddNamedCell will duplicate the cell creating an
// invaild spreadsheet.
func (_gcgb Row) AddNamedCell(col string) Cell {
	_bffb := _ca.NewCT_Cell()
	_bffb.RAttr = _c.Stringf("\u0025\u0073\u0025\u0064", col, _gcgb.RowNumber())
	_bccf := -1
	_aeb := _db.ColumnToIndex(col)
	for _gccd, _baff := range _gcgb._gfcg.C {
		_badgc, _bacd := _db.ParseCellReference(*_baff.RAttr)
		if _bacd != nil {
			return Cell{}
		}
		if _aeb < _badgc.ColumnIdx {
			_bccf = _gccd
			break
		}
	}
	if _bccf == -1 {
		_gcgb._gfcg.C = append(_gcgb._gfcg.C, _bffb)
	} else {
		_gcgb._gfcg.C = append(_gcgb._gfcg.C[:_bccf], append([]*_ca.CT_Cell{_bffb}, _gcgb._gfcg.C[_bccf:]...)...)
	}
	return Cell{_gcgb._bgb, _gcgb._dcab, _gcgb._gfcg, _bffb}
}

// RemoveMergedCell removes merging from a cell range within a sheet.  The cells
// that made up the merged cell remain, but are no lon merged.
func (_ffeb *Sheet) RemoveMergedCell(mc MergedCell) {
	for _fabe, _dggg := range _ffeb._fcbf.MergeCells.MergeCell {
		if _dggg == mc.X() {
			copy(_ffeb._fcbf.MergeCells.MergeCell[_fabe:], _ffeb._fcbf.MergeCells.MergeCell[_fabe+1:])
			_ffeb._fcbf.MergeCells.MergeCell[len(_ffeb._fcbf.MergeCells.MergeCell)-1] = nil
			_ffeb._fcbf.MergeCells.MergeCell = _ffeb._fcbf.MergeCells.MergeCell[:len(_ffeb._fcbf.MergeCells.MergeCell)-1]
		}
	}
}

// GetOrCreateStandardNumberFormat gets or creates a cell style with a given
// standard format. This should only be used when you want to perform
// number/date/time formatting only.  Manipulating the style returned will cause
// all cells using style returned from this for a given format to be formatted.
func (_gcbb StyleSheet) GetOrCreateStandardNumberFormat(f StandardFormat) CellStyle {
	for _, _ggbfg := range _gcbb.CellStyles() {
		if _ggbfg.HasNumberFormat() && _ggbfg.NumberFormat() == uint32(f) {
			return _ggbfg
		}
	}
	_ggfg := _gcbb.AddCellStyle()
	_ggfg.SetNumberFormatStandard(f)
	return _ggfg
}

// HasNumberFormat returns true if the cell style has a number format applied.
func (_dge CellStyle) HasNumberFormat() bool {
	return _dge._bee.NumFmtIdAttr != nil && _dge._bee.ApplyNumberFormatAttr != nil && *_dge._bee.ApplyNumberFormatAttr
}

// AddComment adds a new comment and returns a RichText which will contain the
// styled comment text.
func (_fgce Comments) AddComment(cellRef string, author string) RichText {
	_cabe := _ca.NewCT_Comment()
	_fgce._eca.CommentList.Comment = append(_fgce._eca.CommentList.Comment, _cabe)
	_cabe.RefAttr = cellRef
	_cabe.AuthorIdAttr = _fgce.getOrCreateAuthor(author)
	_cabe.Text = _ca.NewCT_Rst()
	return RichText{_cabe.Text}
}

var _becf []string = []string{"\u007a\u0068\u002dH\u004b", "\u007a\u0068\u002dM\u004f", "\u007a\u0068\u002dC\u004e", "\u007a\u0068\u002dS\u0047", "\u007a\u0068\u002dT\u0057", "\u006a\u0061\u002dJ\u0050", "\u006b\u006f\u002dK\u0052"}

// SetRowOffset sets the row offset of the two cell anchor
func (_acfa TwoCellAnchor) SetRowOffset(m _ee.Distance) {
	_bccb := m - _acfa.TopLeft().RowOffset()
	_acfa.TopLeft().SetRowOffset(m)
	_acfa.BottomRight().SetRowOffset(_acfa.BottomRight().RowOffset() + _bccb)
}
func (_geb Cell) getFormat() string {
	if _geb._ec.SAttr == nil {
		return "\u0047e\u006e\u0065\u0072\u0061\u006c"
	}
	_edb := *_geb._ec.SAttr
	_dfg := _geb._adb.StyleSheet.GetCellStyle(_edb)
	_abce := _geb._adb.StyleSheet.GetNumberFormat(_dfg.NumberFormat())
	return _abce.GetFormat()
}
func (_ggd Cell) getRawSortValue() (string, bool) {
	if _ggd.HasFormula() {
		_agb := _ggd.GetCachedFormulaResult()
		return _agb, _ed.IsNumber(_agb)
	}
	_gff, _ := _ggd.GetRawValue()
	return _gff, _ed.IsNumber(_gff)
}

// Clear clears the cell's value and type.
func (_gf Cell) Clear() { _gf.clearValue(); _gf._ec.TAttr = _ca.ST_CellTypeUnset }
func (_bb Border) SetBottom(style _ca.ST_BorderStyle, c _ae.Color) {
	if _bb._eac.Bottom == nil {
		_bb._eac.Bottom = _ca.NewCT_BorderPr()
	}
	_bb._eac.Bottom.Color = _ca.NewCT_Color()
	_bb._eac.Bottom.Color.RgbAttr = c.AsRGBAString()
	_bb._eac.Bottom.StyleAttr = style
}

var _adf *_aa.Regexp = _aa.MustCompile("\u005e(\u005ba\u002d\u007a\u005d\u002b\u0029(\u005b\u0030-\u0039\u005d\u002b\u0029\u0024")

func (_eagd Font) SetItalic(b bool) {
	if b {
		_eagd._acaf.I = []*_ca.CT_BooleanProperty{{}}
	} else {
		_eagd._acaf.I = nil
	}
}

// AddCell adds a cell to a spreadsheet.
func (_gded Row) AddCell() Cell {
	_fcbb := uint32(len(_gded._gfcg.C))
	var _abbf *string
	if _fcbb > 0 {
		_cba := _c.Stringf("\u0025\u0073\u0025\u0064", _db.IndexToColumn(_fcbb-1), _gded.RowNumber())
		if _gded._gfcg.C[_fcbb-1].RAttr != nil && *_gded._gfcg.C[_fcbb-1].RAttr == *_cba {
			_abbf = _c.Stringf("\u0025\u0073\u0025\u0064", _db.IndexToColumn(_fcbb), _gded.RowNumber())
		}
	}
	_gdgc := _ca.NewCT_Cell()
	_gded._gfcg.C = append(_gded._gfcg.C, _gdgc)
	if _abbf == nil {
		_eadc := uint32(0)
		for _, _bcec := range _gded._gfcg.C {
			if _bcec.RAttr != nil {
				_faa, _ := _db.ParseCellReference(*_bcec.RAttr)
				if _faa.ColumnIdx >= _eadc {
					_eadc = _faa.ColumnIdx + 1
				}
			}
		}
		_abbf = _c.Stringf("\u0025\u0073\u0025\u0064", _db.IndexToColumn(_eadc), _gded.RowNumber())
	}
	_gdgc.RAttr = _abbf
	return Cell{_gded._bgb, _gded._dcab, _gded._gfcg, _gdgc}
}

// X returns the inner wrapped XML type.
func (_abge *Workbook) X() *_ca.Workbook { return _abge._geca }

// SetVerticalAlignment sets the vertical alignment of a cell style.
func (_gacd CellStyle) SetVerticalAlignment(a _ca.ST_VerticalAlignment) {
	if _gacd._bee.Alignment == nil {
		_gacd._bee.Alignment = _ca.NewCT_CellAlignment()
	}
	_gacd._bee.ApplyAlignmentAttr = _c.Bool(true)
	_gacd._bee.Alignment.VerticalAttr = a
}

// LessCells returns true if the lhs value is less than the rhs value. If the
// cells contain numeric values, their value interpreted as a floating point is
// compared. Otherwise their string contents are compared.
func (_gdg Comparer) LessCells(lhs, rhs Cell) bool {
	if _gdg.Order == SortOrderDescending {
		lhs, rhs = rhs, lhs
	}
	if lhs.X() == nil {
		if rhs.X() == nil {
			return false
		}
		return true
	}
	if rhs.X() == nil {
		return false
	}
	_fef, _dfge := lhs.getRawSortValue()
	_afac, _dcdf := rhs.getRawSortValue()
	switch {
	case _dfge && _dcdf:
		_cfc, _ := _bd.ParseFloat(_fef, 64)
		_cfbc, _ := _bd.ParseFloat(_afac, 64)
		return _cfc < _cfbc
	case _dfge:
		return true
	case _dcdf:
		return false
	}
	_fef = lhs.GetFormattedValue()
	_afac = rhs.GetFormattedValue()
	return _fef < _afac
}
func (_bbfg CellStyle) SetNumberFormat(s string) {
	_adbf := _bbfg._gce.StyleSheet.AddNumberFormat()
	_adbf.SetFormat(s)
	_bbfg._bee.ApplyNumberFormatAttr = _c.Bool(true)
	_bbfg._bee.NumFmtIdAttr = _c.Uint32(_adbf.ID())
}
func (_ddfbf PatternFill) SetBgColor(c _ae.Color) {
	_ddfbf._fbdg.BgColor = _ca.NewCT_Color()
	_ddfbf._fbdg.BgColor.RgbAttr = c.AsRGBAString()
}

// SetFormulaRaw sets the cell type to formula, and the raw formula to the given string
func (_fdc Cell) SetFormulaRaw(s string) {
	_eff := _gbf.ParseString(s)
	if _eff == nil {
		return
	}
	_fdc.clearValue()
	_fdc._ec.TAttr = _ca.ST_CellTypeStr
	_fdc._ec.F = _ca.NewCT_CellFormula()
	_fdc._ec.F.Content = s
}

// ID returns the number format ID.  This is not an index as there are some
// predefined number formats which can be used in cell styles and don't need a
// corresponding NumberFormat.
func (_cgcd NumberFormat) ID() uint32 { return _cgcd._geeb.NumFmtIdAttr }

// ValidateWithPath validates the sheet passing path informaton for a better
// error message
func (_bgca Sheet) ValidateWithPath(path string) error { return _bgca._fcbf.ValidateWithPath(path) }

// RemoveColumn removes column from the sheet and moves all columns to the right of the removed column one step left.
func (_cggc *Sheet) RemoveColumn(column string) error {
	_bccg, _gceg := _cggc.getAllCellsInFormulaArraysForColumn()
	if _gceg != nil {
		return _gceg
	}
	_cffg := _db.ColumnToIndex(column)
	for _, _gdfdb := range _cggc.Rows() {
		_dfda := _ge.Sprintf("\u0025\u0073\u0025\u0064", column, *_gdfdb.X().RAttr)
		if _, _bdffa := _bccg[_dfda]; _bdffa {
			return nil
		}
	}
	for _, _eea := range _cggc.Rows() {
		_ebga := _eea._gfcg.C
		for _dcafa, _gafa := range _ebga {
			_gege, _dfab := _db.ParseCellReference(*_gafa.RAttr)
			if _dfab != nil {
				return _dfab
			}
			if _gege.ColumnIdx == _cffg {
				_eea._gfcg.C = append(_ebga[:_dcafa], _cggc.slideCellsLeft(_ebga[_dcafa+1:])...)
				break
			} else if _gege.ColumnIdx > _cffg {
				_eea._gfcg.C = append(_ebga[:_dcafa], _cggc.slideCellsLeft(_ebga[_dcafa:])...)
				break
			}
		}
	}
	_gceg = _cggc.updateAfterRemove(_cffg, _aac.UpdateActionRemoveColumn)
	if _gceg != nil {
		return _gceg
	}
	_gceg = _cggc.removeColumnFromNamedRanges(_cffg)
	if _gceg != nil {
		return _gceg
	}
	_gceg = _cggc.removeColumnFromMergedCells(_cffg)
	if _gceg != nil {
		return _gceg
	}
	for _, _feg := range _cggc._ccae.Sheets() {
		_feg.RecalculateFormulas()
	}
	return nil
}

// StyleSheet is a document style sheet.
type StyleSheet struct {
	_aba  *Workbook
	_bbbe *_ca.StyleSheet
}

func _afge(_ccfg *Sheet) *evalContext {
	return &evalContext{_gag: _ccfg, _dce: make(map[string]struct{})}
}

// X returns the inner wrapped XML type.
func (_abca RichText) X() *_ca.CT_Rst { return _abca._egacf }
func (_bef Cell) getLocked() bool {
	if _bef._ec.SAttr == nil {
		return false
	}
	_fda := *_bef._ec.SAttr
	_eacc := _bef._adb.StyleSheet.GetCellStyle(_fda)
	return *_eacc._bee.Protection.LockedAttr
}
func (_fed Sheet) validateMergedCells() error {
	_edea := map[uint64]struct{}{}
	for _, _bfeb := range _fed.MergedCells() {
		_eegf, _egef, _fagd := _db.ParseRangeReference(_bfeb.Reference())
		if _fagd != nil {
			return _ge.Errorf("\u0073\u0068e\u0065\u0074\u0020\u006e\u0061m\u0065\u0020\u0027\u0025\u0073'\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0025\u0073", _fed.Name(), _bfeb.Reference())
		}
		for _gbfgd := _eegf.RowIdx; _gbfgd <= _egef.RowIdx; _gbfgd++ {
			for _fbbg := _eegf.ColumnIdx; _fbbg <= _egef.ColumnIdx; _fbbg++ {
				_fad := uint64(_gbfgd)<<32 | uint64(_fbbg)
				if _, _fbde := _edea[_fad]; _fbde {
					return _ge.Errorf("\u0073\u0068\u0065\u0065\u0074\u0020n\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027\u0020\u0068\u0061\u0073\u0020\u006f\u0076\u0065\u0072\u006c\u0061\u0070p\u0069\u006e\u0067\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006cl\u0020r\u0061\u006e\u0067\u0065", _fed.Name())
				}
				_edea[_fad] = struct{}{}
			}
		}
	}
	return nil
}

type MergedCell struct {
	_bdcg *Workbook
	_fcba *Sheet
	_gbb  *_ca.CT_MergeCell
}

func (_geda *Sheet) getAllCellsInFormulaArraysForColumn() (map[string]bool, error) {
	return _geda.getAllCellsInFormulaArrays(false)
}

// RemoveFont removes a font from the style sheet.  It *does not* update styles that refer
// to this font.
func (_cfge StyleSheet) RemoveFont(f Font) error {
	for _cfefb, _gcf := range _cfge._bbbe.Fonts.Font {
		if _gcf == f.X() {
			_cfge._bbbe.Fonts.Font = append(_cfge._bbbe.Fonts.Font[:_cfefb], _cfge._bbbe.Fonts.Font[_cfefb+1:]...)
			return nil
		}
	}
	return _gb.New("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}
