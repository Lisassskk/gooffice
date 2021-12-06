
package spreadsheet ;import (_gg "archive/zip";_bf "bytes";_dec "errors";_e "fmt";_b "gitee.com/gooffice/gooffice";_ae "gitee.com/gooffice/gooffice/chart";_fde "gitee.com/gooffice/gooffice/color";_da "gitee.com/gooffice/gooffice/common";_cd "gitee.com/gooffice/gooffice/common/logger";_cdc "gitee.com/gooffice/gooffice/common/tempstorage";_cb "gitee.com/gooffice/gooffice/internal/license";_ec "gitee.com/gooffice/gooffice/measurement";_bgc "gitee.com/gooffice/gooffice/schema/soo/dml";_fc "gitee.com/gooffice/gooffice/schema/soo/dml/chart";_acd "gitee.com/gooffice/gooffice/schema/soo/dml/spreadsheetDrawing";_cef "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships";_ca "gitee.com/gooffice/gooffice/schema/soo/sml";_ga "gitee.com/gooffice/gooffice/spreadsheet/format";_ab "gitee.com/gooffice/gooffice/spreadsheet/formula";_fda "gitee.com/gooffice/gooffice/spreadsheet/reference";_df "gitee.com/gooffice/gooffice/spreadsheet/update";_bg "gitee.com/gooffice/gooffice/vmldrawing";_ce "gitee.com/gooffice/gooffice/zippkg";_bc "image";_dg "image/jpeg";_f "io";_db "math";_ac "math/big";_g "os";_de "path";_a "path/filepath";_bd "regexp";_c "sort";_fb "strconv";_fa "strings";_fd "time";);

// SetFont applies a font to a cell style. The font is referenced by its
// index so modifying the font afterward will affect all styles that reference
// it.
func (_dbec CellStyle )SetFont (f Font ){_dbec ._fbg .FontIdAttr =_b .Uint32 (f .Index ());_dbec ._fbg .ApplyFontAttr =_b .Bool (true );};

// SetRange sets the cell or range of cells that the validation should apply to.
// It can be a single cell (e.g. "A1") or a range of cells (e.g. "A1:B5")
func (_bfg DataValidation )SetRange (cellRange string ){_bfg ._ege .SqrefAttr =_ca .ST_Sqref {cellRange }};func (_ebcb ConditionalFormattingRule )clear (){_ebcb ._afcg .OperatorAttr =_ca .ST_ConditionalFormattingOperatorUnset ;_ebcb ._afcg .ColorScale =nil ;_ebcb ._afcg .IconSet =nil ;_ebcb ._afcg .Formula =nil ;};

// Text returns text from the sheet as one string separated with line breaks.
func (_gcb *SheetText )Text ()string {_aebc :=_bf .NewBuffer ([]byte {});for _ ,_abcf :=range _gcb .Cells {if _abcf .Text !=""{_aebc .WriteString (_abcf .Text );_aebc .WriteString ("\u000a");};};return _aebc .String ();};

// TopLeft returns the CellMaker for the top left corner of the anchor.
func (_fdfe TwoCellAnchor )TopLeft ()CellMarker {return CellMarker {_fdfe ._ffgg .From }};

// SetXSplit sets the column split point
func (_cffae SheetView )SetXSplit (v float64 ){_cffae .ensurePane ();_cffae ._egcb .Pane .XSplitAttr =_b .Float64 (v );};

// SetDate sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel). The cell is not styled via this method, so it will
// display as a number. SetDateWithStyle should normally be used instead.
func (_aad Cell )SetDate (d _fd .Time ){_aad .clearValue ();d =_egd (d );_cdaa :=_aad ._fac .Epoch ();if d .Before (_cdaa ){_cd .Log .Debug ("d\u0061\u0074\u0065\u0073\u0020\u0062e\u0066\u006f\u0072\u0065\u0020\u00319\u0030\u0030\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064");return ;};_ea :=d .Sub (_cdaa );_dcf :=new (_ac .Float );_bae :=new (_ac .Float );_bae .SetPrec (128);_bae .SetUint64 (uint64 (_ea ));_dbe :=new (_ac .Float );_dbe .SetUint64 (24*60*60*1e9);_dcf .Quo (_bae ,_dbe );_cefe ,_ :=_dcf .Uint64 ();_aad ._ecc .V =_b .Stringf ("\u0025\u0064",_cefe );};

// Validate validates the sheet, returning an error if it is found to be invalid.
func (_edcc Sheet )Validate ()error {_cbea :=[]func ()error {_edcc .validateRowCellNumbers ,_edcc .validateMergedCells ,_edcc .validateSheetNames };for _ ,_cdf :=range _cbea {if _fbdd :=_cdf ();_fbdd !=nil {return _fbdd ;};};if _fba :=_edcc ._cffg .Validate ();_fba !=nil {return _fba ;};return _edcc ._cffg .Validate ();};

// HasFormula returns true if the cell has an asoociated formula.
func (_ced Cell )HasFormula ()bool {return _ced ._ecc .F !=nil };

// AddBorder creates a new empty border that can be applied to a cell style.
func (_cfdb StyleSheet )AddBorder ()Border {_eaba :=_ca .NewCT_Border ();_cfdb ._gace .Borders .Border =append (_cfdb ._gace .Borders .Border ,_eaba );_cfdb ._gace .Borders .CountAttr =_b .Uint32 (uint32 (len (_cfdb ._gace .Borders .Border )));return Border {_eaba ,_cfdb ._gace .Borders };};const _bge ="\u00320\u0030\u0036\u002d\u00301\u002d\u0030\u0032\u0054\u00315\u003a0\u0034:\u0030\u0035\u005a\u0030\u0037\u003a\u00300";

// TwoCellAnchor is an anchor that is attached to a top-left cell with a fixed
// width/height in cells.
type TwoCellAnchor struct{_ffgg *_acd .CT_TwoCellAnchor };

// Col returns the column of the cell marker.
func (_cdef CellMarker )Col ()int32 {return _cdef ._gebb .Col };

// Cells returns a slice of cells.  The cells can be manipulated, but appending
// to the slice will have no effect.
func (_ecf Row )Cells ()[]Cell {_afg :=[]Cell {};_dcaa :=-1;_bee :=append ([]*_ca .CT_Cell {},_ecf ._gbgd .C ...);for _ ,_aeebg :=range _bee {if _aeebg .RAttr ==nil {_cd .Log .Debug ("\u0052\u0041\u0074tr\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u0020\u0066o\u0072 \u0061 \u0063e\u006c\u006c\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e");continue ;};_gacf ,_gcfc :=_fda .ParseCellReference (*_aeebg .RAttr );if _gcfc !=nil {_cd .Log .Debug ("\u0052\u0041\u0074t\u0072\u0020\u0069\u0073 \u0069\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u0020\u0063\u0065\u006c\u006c\u003a\u0020"+*_aeebg .RAttr +",\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e");continue ;};_cdea :=int (_gacf .ColumnIdx );if _cdea -_dcaa > 1{for _fbfg :=_dcaa +1;_fbfg < _cdea ;_fbfg ++{_afg =append (_afg ,_ecf .Cell (_fda .IndexToColumn (uint32 (_fbfg ))));};};_dcaa =_cdea ;_afg =append (_afg ,Cell {_ecf ._bdaa ,_ecf ._cgee ,_ecf ._gbgd ,_aeebg });};return _afg ;};var _cedb =false ;

// OneCellAnchor is anchored to a top-left cell with a fixed with/height
// in distance.
type OneCellAnchor struct{_fdd *_acd .CT_OneCellAnchor };

// SetWidth sets the width of the anchored object.
func (_ceaa OneCellAnchor )SetWidth (w _ec .Distance ){_ceaa ._fdd .Ext .CxAttr =int64 (w /_ec .EMU )};

// Cell returns the actual cell behind the merged region
func (_gadb MergedCell )Cell ()Cell {_eccg :=_gadb .Reference ();if _efff :=_fa .Index (_gadb .Reference (),"\u003a");_efff !=-1{_eccg =_eccg [0:_efff ];return _gadb ._dgff .Cell (_eccg );};return Cell {};};func (_bgdc ConditionalFormattingRule )InitializeDefaults (){_bgdc .SetType (_ca .ST_CfTypeCellIs );_bgdc .SetOperator (_ca .ST_ConditionalFormattingOperatorGreaterThan );_bgdc .SetPriority (1);};

// RemoveMergedCell removes merging from a cell range within a sheet.  The cells
// that made up the merged cell remain, but are no lon merged.
func (_ebff *Sheet )RemoveMergedCell (mc MergedCell ){for _ggfc ,_gde :=range _ebff ._cffg .MergeCells .MergeCell {if _gde ==mc .X (){copy (_ebff ._cffg .MergeCells .MergeCell [_ggfc :],_ebff ._cffg .MergeCells .MergeCell [_ggfc +1:]);_ebff ._cffg .MergeCells .MergeCell [len (_ebff ._cffg .MergeCells .MergeCell )-1]=nil ;_ebff ._cffg .MergeCells .MergeCell =_ebff ._cffg .MergeCells .MergeCell [:len (_ebff ._cffg .MergeCells .MergeCell )-1];};};};func (_dfad PatternFill )ClearBgColor (){_dfad ._aedg .BgColor =nil };

// SetCachedFormulaResult sets the cached result of a formula. This is normally
// not needed but is used internally when expanding an array formula.
func (_cae Cell )SetCachedFormulaResult (s string ){_cae ._ecc .V =&s };

// Validate attempts to validate the structure of a workbook.
func (_ceea *Workbook )Validate ()error {if _ceea ==nil ||_ceea ._gdeg ==nil {return _dec .New ("\u0077o\u0072\u006bb\u006f\u006f\u006b\u0020n\u006f\u0074\u0020i\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0065d \u0063\u006f\u0072r\u0065\u0063t\u006c\u0079\u002c\u0020\u006e\u0069l\u0020\u0062a\u0073\u0065");};_acdgf :=uint32 (0);for _ ,_acbg :=range _ceea ._gdeg .Sheets .Sheet {if _acbg .SheetIdAttr > _acdgf {_acdgf =_acbg .SheetIdAttr ;};};if _acdgf !=uint32 (len (_ceea ._bbced )){return _e .Errorf ("\u0066\u006f\u0075\u006e\u0064\u0020%\u0064\u0020\u0077\u006f\u0072\u006b\u0073\u0068\u0065\u0065\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074i\u006f\u006e\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0064\u0020\u0077\u006f\u0072k\u0073h\u0065\u0065\u0074\u0073",_acdgf ,len (_ceea ._bbced ));};_eeec :=map[string ]struct{}{};for _ggff ,_aefd :=range _ceea ._gdeg .Sheets .Sheet {_ebdf :=Sheet {_ceea ,_aefd ,_ceea ._bbced [_ggff ]};if _ ,_dcaaf :=_eeec [_ebdf .Name ()];_dcaaf {return _e .Errorf ("\u0077\u006f\u0072k\u0062\u006f\u006f\u006b\u002f\u0053\u0068\u0065\u0065\u0074\u005b\u0025\u0064\u005d\u0020\u0068\u0061\u0073\u0020\u0064\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065\u0020n\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027",_ggff ,_ebdf .Name ());};_eeec [_ebdf .Name ()]=struct{}{};if _facc :=_ebdf .ValidateWithPath (_e .Sprintf ("\u0077o\u0072k\u0062\u006f\u006f\u006b\u002fS\u0068\u0065e\u0074\u005b\u0025\u0064\u005d",_ggff ));_facc !=nil {return _facc ;};if _gcdb :=_ebdf .Validate ();_gcdb !=nil {return _gcdb ;};};return nil ;};

// SetAllowBlank controls if blank values are accepted.
func (_fef DataValidation )SetAllowBlank (b bool ){if !b {_fef ._ege .AllowBlankAttr =nil ;}else {_fef ._ege .AllowBlankAttr =_b .Bool (true );};};

// AddChart adds an chart to a drawing, returning the chart and an anchor that
// can be used to position the chart within the sheet.
func (_abbf Drawing )AddChart (at AnchorType )(_ae .Chart ,Anchor ){_cdcb :=_fc .NewChartSpace ();_abbf ._fge ._dgeg =append (_abbf ._fge ._dgeg ,_cdcb );_bacc :=_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .ChartContentType ,len (_abbf ._fge ._dgeg ));_abbf ._fge .ContentTypes .AddOverride (_bacc ,_b .ChartContentType );var _gab string ;for _ebcbe ,_cff :=range _abbf ._fge ._cffgf {if _cff ==_abbf ._bda {_aeab :=_b .RelativeFilename (_b .DocTypeSpreadsheet ,_b .DrawingType ,_b .ChartType ,len (_abbf ._fge ._dgeg ));_ede :=_abbf ._fge ._aeec [_ebcbe ].AddRelationship (_aeab ,_b .ChartType );_gab =_ede .ID ();break ;};};var _eefg Anchor ;var _aeabd *_acd .CT_GraphicalObjectFrame ;switch at {case AnchorTypeAbsolute :_bcac :=_cfec ();_abbf ._bda .EG_Anchor =append (_abbf ._bda .EG_Anchor ,&_acd .EG_Anchor {AbsoluteAnchor :_bcac });_bcac .Choice =&_acd .EG_ObjectChoicesChoice {};_bcac .Choice .GraphicFrame =_acd .NewCT_GraphicalObjectFrame ();_aeabd =_bcac .Choice .GraphicFrame ;_eefg =AbsoluteAnchor {_bcac };case AnchorTypeOneCell :_agf :=_eff ();_abbf ._bda .EG_Anchor =append (_abbf ._bda .EG_Anchor ,&_acd .EG_Anchor {OneCellAnchor :_agf });_agf .Choice =&_acd .EG_ObjectChoicesChoice {};_agf .Choice .GraphicFrame =_acd .NewCT_GraphicalObjectFrame ();_aeabd =_agf .Choice .GraphicFrame ;_eefg =OneCellAnchor {_agf };case AnchorTypeTwoCell :_fbba :=_cbdf ();_abbf ._bda .EG_Anchor =append (_abbf ._bda .EG_Anchor ,&_acd .EG_Anchor {TwoCellAnchor :_fbba });_fbba .Choice =&_acd .EG_ObjectChoicesChoice {};_fbba .Choice .GraphicFrame =_acd .NewCT_GraphicalObjectFrame ();_aeabd =_fbba .Choice .GraphicFrame ;_eefg =TwoCellAnchor {_fbba };};_aeabd .NvGraphicFramePr =_acd .NewCT_GraphicalObjectFrameNonVisual ();_aeabd .NvGraphicFramePr .CNvPr .IdAttr =uint32 (len (_abbf ._bda .EG_Anchor ));_aeabd .NvGraphicFramePr .CNvPr .NameAttr ="\u0043\u0068\u0061r\u0074";_aeabd .Graphic =_bgc .NewGraphic ();_aeabd .Graphic .GraphicData .UriAttr ="\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073.\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006dl/\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074";_egc :=_fc .NewChart ();_egc .IdAttr =_gab ;_aeabd .Graphic .GraphicData .Any =[]_b .Any {_egc };_dbef :=_ae .MakeChart (_cdcb );_dbef .Properties ().SetSolidFill (_fde .White );_dbef .SetDisplayBlanksAs (_fc .ST_DispBlanksAsGap );return _dbef ,_eefg ;};

// AddFont adds a new empty font to the stylesheet.
func (_agedc StyleSheet )AddFont ()Font {_daagd :=_ca .NewCT_Font ();_agedc ._gace .Fonts .Font =append (_agedc ._gace .Fonts .Font ,_daagd );_agedc ._gace .Fonts .CountAttr =_b .Uint32 (uint32 (len (_agedc ._gace .Fonts .Font )));return Font {_daagd ,_agedc ._gace };};

// SetStringByID sets the cell type to string, and the value a string in the
// shared strings table.
func (_caa Cell )SetStringByID (id int ){_caa ._fac .ensureSharedStringsRelationships ();_caa .clearValue ();_caa ._ecc .V =_b .String (_fb .Itoa (id ));_caa ._ecc .TAttr =_ca .ST_CellTypeS ;};

// GetCachedFormulaResult returns the cached formula result if it exists. If the
// cell type is not a formula cell, the result will be the cell value if it's a
// string/number/bool cell.
func (_gfe Cell )GetCachedFormulaResult ()string {if _gfe ._ecc .V !=nil {return *_gfe ._ecc .V ;};return "";};

// SetDateWithStyle sets a date with the default date style applied.
func (_gec Cell )SetDateWithStyle (d _fd .Time ){_gec .SetDate (d );for _ ,_fed :=range _gec ._fac .StyleSheet .CellStyles (){if _fed .HasNumberFormat ()&&_fed .NumberFormat ()==uint32 (StandardFormatDate ){_gec .SetStyle (_fed );return ;};};_aaf :=_gec ._fac .StyleSheet .AddCellStyle ();_aaf .SetNumberFormatStandard (StandardFormatDate );_gec .SetStyle (_aaf );};

// LastRow returns the name of last row which contains data in range of context sheet's given columns.
func (_daed *evalContext )LastRow (col string )int {_aade :=_daed ._bcdc ;_bfe :=int (_fda .ColumnToIndex (col ));_eecb :=1;for _ ,_edg :=range _aade ._cffg .SheetData .Row {if _edg .RAttr !=nil {_bbb :=Row {_aade ._fdfd ,_aade ,_edg };_ccf :=len (_bbb .Cells ());if _ccf > _bfe {_eecb =int (_bbb .RowNumber ());};};};return _eecb ;};

// AddCell adds a cell to a spreadsheet.
func (_cegg Row )AddCell ()Cell {_baad :=uint32 (len (_cegg ._gbgd .C ));var _bdbf *string ;if _baad > 0{_cbae :=_b .Stringf ("\u0025\u0073\u0025\u0064",_fda .IndexToColumn (_baad -1),_cegg .RowNumber ());if _cegg ._gbgd .C [_baad -1].RAttr !=nil &&*_cegg ._gbgd .C [_baad -1].RAttr ==*_cbae {_bdbf =_b .Stringf ("\u0025\u0073\u0025\u0064",_fda .IndexToColumn (_baad ),_cegg .RowNumber ());};};_cad :=_ca .NewCT_Cell ();_cegg ._gbgd .C =append (_cegg ._gbgd .C ,_cad );if _bdbf ==nil {_dege :=uint32 (0);for _ ,_decbb :=range _cegg ._gbgd .C {if _decbb .RAttr !=nil {_ddbce ,_ :=_fda .ParseCellReference (*_decbb .RAttr );if _ddbce .ColumnIdx >=_dege {_dege =_ddbce .ColumnIdx +1;};};};_bdbf =_b .Stringf ("\u0025\u0073\u0025\u0064",_fda .IndexToColumn (_dege ),_cegg .RowNumber ());};_cad .RAttr =_bdbf ;return Cell {_cegg ._bdaa ,_cegg ._cgee ,_cegg ._gbgd ,_cad };};

// SetSize sets the text size for a rich text run.
func (_fgea RichTextRun )SetSize (m _ec .Distance ){_fgea .ensureRpr ();_fgea ._dfc .RPr .Sz =_ca .NewCT_FontSize ();_fgea ._dfc .RPr .Sz .ValAttr =float64 (m /_ec .Point );};

// Border is a cell border configuraton.
type Border struct{_fcf *_ca .CT_Border ;_cbd *_ca .CT_Borders ;};const (SortOrderAscending SortOrder =iota ;SortOrderDescending ;);

// Row returns the row of the cell marker.
func (_ada CellMarker )Row ()int32 {return _ada ._gebb .Row };func (_dcbb *Sheet )setList (_ddcf string ,_dabff _ab .Result )error {_gegb ,_acgg :=_fda .ParseCellReference (_ddcf );if _acgg !=nil {return _acgg ;};_fegb :=_dcbb .Row (_gegb .RowIdx );for _eecc ,_egdd :=range _dabff .ValueList {_bced :=_fegb .Cell (_fda .IndexToColumn (_gegb .ColumnIdx +uint32 (_eecc )));if _egdd .Type !=_ab .ResultTypeEmpty {if _egdd .IsBoolean {_bced .SetBool (_egdd .ValueNumber !=0);}else {_bced .SetCachedFormulaResult (_egdd .String ());};};};return nil ;};

// X returns the inner wrapped XML type.
func (_abg Cell )X ()*_ca .CT_Cell {return _abg ._ecc };

// IsStructureLocked returns whether the workbook structure is locked.
func (_fddb WorkbookProtection )IsStructureLocked ()bool {return _fddb ._acaf .LockStructureAttr !=nil &&*_fddb ._acaf .LockStructureAttr ;};func (_cf Border )SetRight (style _ca .ST_BorderStyle ,c _fde .Color ){if _cf ._fcf .Right ==nil {_cf ._fcf .Right =_ca .NewCT_BorderPr ();};_cf ._fcf .Right .Color =_ca .NewCT_Color ();_cf ._fcf .Right .Color .RgbAttr =c .AsRGBAString ();_cf ._fcf .Right .StyleAttr =style ;};

// PasswordHash returns the hash of the workbook password.
func (_ceffc WorkbookProtection )PasswordHash ()string {if _ceffc ._acaf .WorkbookPasswordAttr ==nil {return "";};return *_ceffc ._acaf .WorkbookPasswordAttr ;};

// AddString adds a string to the shared string cache.
func (_bbce SharedStrings )AddString (v string )int {if _aged ,_ddge :=_bbce ._fbfea [v ];_ddge {return _aged ;};_aafg :=_ca .NewCT_Rst ();_aafg .T =_b .String (v );_bbce ._cfca .Si =append (_bbce ._cfca .Si ,_aafg );_fcea :=len (_bbce ._cfca .Si )-1;_bbce ._fbfea [v ]=_fcea ;_bbce ._cfca .CountAttr =_b .Uint32 (uint32 (len (_bbce ._cfca .Si )));_bbce ._cfca .UniqueCountAttr =_bbce ._cfca .CountAttr ;return _fcea ;};

// X returns the inner wrapped XML type.
func (_bgd ColorScale )X ()*_ca .CT_ColorScale {return _bgd ._dfdc };

// SetWidthCells is a no-op.
func (_cccc OneCellAnchor )SetWidthCells (int32 ){};

// Reference returns the region of cells that are merged.
func (_bcff MergedCell )Reference ()string {return _bcff ._fgf .RefAttr };

// Workbook returns sheet's parent workbook.
func (_ebab *Sheet )Workbook ()*Workbook {return _ebab ._fdfd };var _adab =_b .RelativeFilename (_b .DocTypeSpreadsheet ,_b .OfficeDocumentType ,_b .SharedStringsType ,0);

// X returns the inner wrapped XML type.
func (_ddfd NumberFormat )X ()*_ca .CT_NumFmt {return _ddfd ._bgfa };

// LockSheet controls the locking of the sheet.
func (_fefe SheetProtection )LockSheet (b bool ){if !b {_fefe ._ebec .SheetAttr =nil ;}else {_fefe ._ebec .SheetAttr =_b .Bool (true );};};

// ColOffset returns the offset from the row cell.
func (_daf CellMarker )ColOffset ()_ec .Distance {if _daf ._gebb .RowOff .ST_CoordinateUnqualified ==nil {return 0;};return _ec .Distance (float64 (*_daf ._gebb .ColOff .ST_CoordinateUnqualified )*_ec .EMU );};func (_bafa RichTextRun )ensureRpr (){if _bafa ._dfc .RPr ==nil {_bafa ._dfc .RPr =_ca .NewCT_RPrElt ();};};func (_cba Border )SetBottom (style _ca .ST_BorderStyle ,c _fde .Color ){if _cba ._fcf .Bottom ==nil {_cba ._fcf .Bottom =_ca .NewCT_BorderPr ();};_cba ._fcf .Bottom .Color =_ca .NewCT_Color ();_cba ._fcf .Bottom .Color .RgbAttr =c .AsRGBAString ();_cba ._fcf .Bottom .StyleAttr =style ;};

// SetPassword sets the password hash to a hash of the input password.
func (_bbcb SheetProtection )SetPassword (pw string ){_bbcb .SetPasswordHash (PasswordHash (pw ))};

// IsEmpty returns true if the cell is empty.
func (_cbg Cell )IsEmpty ()bool {return _cbg ._ecc .TAttr ==_ca .ST_CellTypeUnset &&_cbg ._ecc .V ==nil &&_cbg ._ecc .F ==nil ;};

// Comments returns the list of comments for this sheet
func (_abcd Comments )Comments ()[]Comment {_agc :=[]Comment {};for _ ,_ffc :=range _abcd ._cfc .CommentList .Comment {_agc =append (_agc ,Comment {_abcd ._gebe ,_ffc ,_abcd ._cfc });};return _agc ;};

// IsBool returns true if the cell boolean value.
func (_efg *evalContext )IsBool (cellRef string )bool {return _efg ._bcdc .Cell (cellRef ).IsBool ()};

// ConditionalFormatting controls the formatting styles and rules for a range of
// cells with the same conditional formatting.
type ConditionalFormatting struct{_gfde *_ca .CT_ConditionalFormatting };func (_ggc Cell )getLabelPrefix ()string {if _ggc ._ecc .SAttr ==nil {return "";};_dgb :=*_ggc ._ecc .SAttr ;_gac :=_ggc ._fac .StyleSheet .GetCellStyle (_dgb );switch _gac ._fbg .Alignment .HorizontalAttr {case _ca .ST_HorizontalAlignmentLeft :return "\u0027";case _ca .ST_HorizontalAlignmentRight :return "\u0022";case _ca .ST_HorizontalAlignmentCenter :return "\u005e";case _ca .ST_HorizontalAlignmentFill :return "\u005c";default:return "";};};

// SetBool sets the cell type to boolean and the value to the given boolean
// value.
func (_cbda Cell )SetBool (v bool ){_cbda .clearValue ();_cbda ._ecc .V =_b .String (_fb .Itoa (_ccab (v )));_cbda ._ecc .TAttr =_ca .ST_CellTypeB ;};

// IsHidden returns whether the row is hidden or not.
func (_bagd Row )IsHidden ()bool {return _bagd ._gbgd .HiddenAttr !=nil &&*_bagd ._gbgd .HiddenAttr };

// SetWidth controls the width of a column.
func (_dde Column )SetWidth (w _ec .Distance ){_dde ._geg .WidthAttr =_b .Float64 (float64 (w /_ec .Character ));};const (DVCompareTypeWholeNumber =DVCompareType (_ca .ST_DataValidationTypeWhole );DVCompareTypeDecimal =DVCompareType (_ca .ST_DataValidationTypeDecimal );DVCompareTypeDate =DVCompareType (_ca .ST_DataValidationTypeDate );DVCompareTypeTime =DVCompareType (_ca .ST_DataValidationTypeTime );DVompareTypeTextLength =DVCompareType (_ca .ST_DataValidationTypeTextLength ););

// StandardFormat is a standard ECMA 376 number format.
//go:generate stringer -type=StandardFormat
type StandardFormat uint32 ;

// SetColorScale configures the rule as a color scale, removing existing
// configuration.
func (_egag ConditionalFormattingRule )SetColorScale ()ColorScale {_egag .clear ();_egag .SetType (_ca .ST_CfTypeColorScale );_egag ._afcg .ColorScale =_ca .NewCT_ColorScale ();return ColorScale {_egag ._afcg .ColorScale };};

// GetFont gets a Font from a cell style.
func (_abb CellStyle )GetFont ()*_ca .CT_Font {if _ceee :=_abb ._fbg .FontIdAttr ;_ceee !=nil {_ecb :=_abb ._dcb .StyleSheet .Fonts ();if int (*_ceee )< len (_ecb ){return _ecb [int (*_ceee )].X ();};};return nil ;};

// GetVerticalAlignment sets the vertical alignment of a cell style.
func (_dfaa CellStyle )GetVerticalAlignment ()_ca .ST_VerticalAlignment {if _dfaa ._fbg .Alignment ==nil {return _ca .ST_VerticalAlignmentUnset ;};return _dfaa ._fbg .Alignment .VerticalAttr ;};

// SetHeight sets the height of the anchored object.
func (_gea OneCellAnchor )SetHeight (h _ec .Distance ){_gea ._fdd .Ext .CyAttr =int64 (h /_ec .EMU )};

// SetText sets the text to be displayed.
func (_gcac RichTextRun )SetText (s string ){_gcac ._dfc .T =s };

// Clear clears the cell's value and type.
func (_afbg Cell )Clear (){_afbg .clearValue ();_afbg ._ecc .TAttr =_ca .ST_CellTypeUnset };

// SetVerticalAlignment sets the vertical alignment of a cell style.
func (_dcg CellStyle )SetVerticalAlignment (a _ca .ST_VerticalAlignment ){if _dcg ._fbg .Alignment ==nil {_dcg ._fbg .Alignment =_ca .NewCT_CellAlignment ();};_dcg ._fbg .ApplyAlignmentAttr =_b .Bool (true );_dcg ._fbg .Alignment .VerticalAttr =a ;};

// SetNumberFormatStandard sets the format based off of the ECMA 376 standard formats.  These
// formats are standardized and don't need to be defined in the styles.
func (_abd CellStyle )SetNumberFormatStandard (s StandardFormat ){_abd ._fbg .NumFmtIdAttr =_b .Uint32 (uint32 (s ));_abd ._fbg .ApplyNumberFormatAttr =_b .Bool (true );};

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (_adca *Workbook )ClearCachedFormulaResults (){for _ ,_bacac :=range _adca .Sheets (){_bacac .ClearCachedFormulaResults ();};};

// SetConditionValue sets the condition value to be used for style applicaton.
func (_cafc ConditionalFormattingRule )SetConditionValue (v string ){_cafc ._afcg .Formula =[]string {v }};

// Cell is a single cell within a sheet.
type Cell struct{_fac *Workbook ;_cfd *Sheet ;_dfa *_ca .CT_Row ;_ecc *_ca .CT_Cell ;};

// RemoveCalcChain removes the cached caculation chain. This is sometimes needed
// as we don't update it when rows are added/removed.
func (_egce *Workbook )RemoveCalcChain (){var _gafde string ;for _ ,_edcb :=range _egce ._eeead .Relationships (){if _edcb .Type ()=="ht\u0074\u0070\u003a\u002f\u002f\u0073\u0063he\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006et\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068i\u0070s\u002f\u0063\u0061\u006c\u0063\u0043\u0068\u0061\u0069\u006e"{_gafde ="\u0078\u006c\u002f"+_edcb .Target ();_egce ._eeead .Remove (_edcb );break ;};};if _gafde ==""{return ;};_egce .ContentTypes .RemoveOverride (_gafde );for _deeb ,_ddaf :=range _egce .ExtraFiles {if _ddaf .ZipPath ==_gafde {_egce .ExtraFiles [_deeb ]=_egce .ExtraFiles [len (_egce .ExtraFiles )-1];_egce .ExtraFiles =_egce .ExtraFiles [:len (_egce .ExtraFiles )-1];return ;};};};

// SetProtectedAndHidden sets protected and hidden for given cellStyle
func (_deg CellStyle )SetProtection (protected bool ,hidden bool ){_deg ._fbg .Protection =&_ca .CT_CellProtection {LockedAttr :&protected ,HiddenAttr :&hidden };};

// SetFill applies a fill to a cell style. The fill is referenced by its index
// so modifying the fill afterward will affect all styles that reference it.
func (_eec CellStyle )SetFill (f Fill ){_eec ._fbg .FillIdAttr =_b .Uint32 (f .Index ());_eec ._fbg .ApplyFillAttr =_b .Bool (true );};

// SetValue sets the first value to be used in the comparison.  For comparisons
// that need only one value, this is the only value used.  For comparisons like
// 'between' that require two values, SetValue2 must also be used.
func (_ebe DataValidationCompare )SetValue (v string ){_ebe ._ecdc .Formula1 =&v };

// FormulaContext returns a formula evaluation context that can be used to
// evaluate formaulas.
func (_gffca *Sheet )FormulaContext ()_ab .Context {return _efaf (_gffca )};

// AddFormatValue adds a format value to be used in determining which icons to display.
func (_dfaac IconScale )AddFormatValue (t _ca .ST_CfvoType ,val string ){_dac :=_ca .NewCT_Cfvo ();_dac .TypeAttr =t ;_dac .ValAttr =_b .String (val );_dfaac ._dead .Cfvo =append (_dfaac ._dead .Cfvo ,_dac );};func NewFills ()Fills {return Fills {_ca .NewCT_Fills ()}};func (_fcca Border )SetLeft (style _ca .ST_BorderStyle ,c _fde .Color ){if _fcca ._fcf .Left ==nil {_fcca ._fcf .Left =_ca .NewCT_BorderPr ();};_fcca ._fcf .Left .Color =_ca .NewCT_Color ();_fcca ._fcf .Left .Color .RgbAttr =c .AsRGBAString ();_fcca ._fcf .Left .StyleAttr =style ;};

// CopySheetByName copies the existing sheet with the name `name` and puts its copy with the name `copiedSheetName`.
func (_cddb *Workbook )CopySheetByName (name ,copiedSheetName string )(Sheet ,error ){_acdb :=-1;for _abdeb ,_cggb :=range _cddb .Sheets (){if name ==_cggb .Name (){_acdb =_abdeb ;break ;};};if _acdb ==-1{return Sheet {},ErrorNotFound ;};return _cddb .CopySheet (_acdb ,copiedSheetName );};

// Priority returns the rule priority
func (_bbe ConditionalFormattingRule )Priority ()int32 {return _bbe ._afcg .PriorityAttr };

// X returns the inner wrapped XML type.
func (_eb Border )X ()*_ca .CT_Border {return _eb ._fcf };

// ClearFill clears any fill configuration from the cell style.
func (_abag CellStyle )ClearFill (){_abag ._fbg .FillIdAttr =nil ;_abag ._fbg .ApplyFillAttr =nil };

// SetHeightCells is a no-op.
func (_bgb AbsoluteAnchor )SetHeightCells (int32 ){};

// SharedStrings is a shared strings table, where string data can be placed
// outside of the sheet contents and referenced from a sheet.
type SharedStrings struct{_cfca *_ca .Sst ;_fbfea map[string ]int ;};func (_efbf SortOrder )String ()string {if _efbf >=SortOrder (len (_ebdg )-1){return _e .Sprintf ("\u0053\u006f\u0072\u0074\u004f\u0072\u0064\u0065\u0072\u0028\u0025\u0064\u0029",_efbf );};return _agag [_ebdg [_efbf ]:_ebdg [_efbf +1]];};

// GetSheet returns a sheet by name, or an error if a sheet by the given name
// was not found.
func (_aacd *Workbook )GetSheet (name string )(Sheet ,error ){for _ ,_ffddc :=range _aacd .Sheets (){if _ffddc .Name ()==name {return _ffddc ,nil ;};};return Sheet {},ErrorNotFound ;};

// X returns the inner wrapped XML type.
func (_efbg SheetProtection )X ()*_ca .CT_SheetProtection {return _efbg ._ebec };func (_eafbc Sheet )validateMergedCells ()error {_gefg :=map[uint64 ]struct{}{};for _ ,_bagg :=range _eafbc .MergedCells (){_dfdg ,_dfaf ,_gebeg :=_fda .ParseRangeReference (_bagg .Reference ());if _gebeg !=nil {return _e .Errorf ("\u0073\u0068e\u0065\u0074\u0020\u006e\u0061m\u0065\u0020\u0027\u0025\u0073'\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0025\u0073",_eafbc .Name (),_bagg .Reference ());};for _ead :=_dfdg .RowIdx ;_ead <=_dfaf .RowIdx ;_ead ++{for _daaee :=_dfdg .ColumnIdx ;_daaee <=_dfaf .ColumnIdx ;_daaee ++{_egdg :=uint64 (_ead )<<32|uint64 (_daaee );if _ ,_ecee :=_gefg [_egdg ];_ecee {return _e .Errorf ("\u0073\u0068\u0065\u0065\u0074\u0020n\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027\u0020\u0068\u0061\u0073\u0020\u006f\u0076\u0065\u0072\u006c\u0061\u0070p\u0069\u006e\u0067\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006cl\u0020r\u0061\u006e\u0067\u0065",_eafbc .Name ());};_gefg [_egdg ]=struct{}{};};};};return nil ;};

// BottomRight is a no-op.
func (_fdag OneCellAnchor )BottomRight ()CellMarker {return CellMarker {}};

// NumberFormat is a number formatting string that can be applied to a cell
// style.
type NumberFormat struct{_bff *Workbook ;_bgfa *_ca .CT_NumFmt ;};func _ccgd (_face string )bool {_face =_fa .Replace (_face ,"\u0024","",-1);if _fdg :=_ceeg .FindStringSubmatch (_fa .ToLower (_face ));len (_fdg )> 2{_caeb :=_fdg [1];_dggg ,_baa :=_fb .Atoi (_fdg [2]);if _baa !=nil {return false ;};return _dggg <=1048576&&_caeb <="\u007a\u007a";};return false ;};func (_gca DataValidationCompare )SetValue2 (v string ){_gca ._ecdc .Formula2 =&v };

// SetPriority sets the rule priority
func (_cbdg ConditionalFormattingRule )SetPriority (p int32 ){_cbdg ._afcg .PriorityAttr =p };

// DefinedNames returns a slice of all defined names in the workbook.
func (_abdb *Workbook )DefinedNames ()[]DefinedName {if _abdb ._gdeg .DefinedNames ==nil {return nil ;};_cabf :=[]DefinedName {};for _ ,_fbed :=range _abdb ._gdeg .DefinedNames .DefinedName {_cabf =append (_cabf ,DefinedName {_fbed });};return _cabf ;};func (_aeg Cell )GetRawValue ()(string ,error ){switch _aeg ._ecc .TAttr {case _ca .ST_CellTypeInlineStr :if _aeg ._ecc .Is ==nil ||_aeg ._ecc .Is .T ==nil {return "",nil ;};return *_aeg ._ecc .Is .T ,nil ;case _ca .ST_CellTypeS :if _aeg ._ecc .V ==nil {return "",nil ;};_bdc ,_bfc :=_fb .Atoi (*_aeg ._ecc .V );if _bfc !=nil {return "",_bfc ;};return _aeg ._fac .SharedStrings .GetString (_bdc );case _ca .ST_CellTypeStr :if _aeg ._ecc .F !=nil {return _aeg ._ecc .F .Content ,nil ;};};if _aeg ._ecc .V ==nil {return "",nil ;};return *_aeg ._ecc .V ,nil ;};

// PasswordHash returns the password hash for a workbook using the modified
// spreadsheetML password hash that is compatible with Excel.
func PasswordHash (s string )string {_bea :=uint16 (0);if len (s )> 0{for _abcg :=len (s )-1;_abcg >=0;_abcg --{_bfcd :=s [_abcg ];_bea =((_bea >>14)&0x01)|((_bea <<1)&0x7fff);_bea ^=uint16 (_bfcd );};_bea =((_bea >>14)&0x01)|((_bea <<1)&0x7fff);_bea ^=uint16 (len (s ));_bea ^=(0x8000|('N'<<8)|'K');};return _e .Sprintf ("\u0025\u0030\u0034\u0058",uint64 (_bea ));};

// Wrapped returns true if the cell will wrap text.
func (_feg CellStyle )Wrapped ()bool {if _feg ._fbg .Alignment ==nil {return false ;};if _feg ._fbg .Alignment .WrapTextAttr ==nil {return false ;};return *_feg ._fbg .Alignment .WrapTextAttr ;};

// LockStructure controls the locking of the workbook structure.
func (_bcfcd WorkbookProtection )LockStructure (b bool ){if !b {_bcfcd ._acaf .LockStructureAttr =nil ;}else {_bcfcd ._acaf .LockStructureAttr =_b .Bool (true );};};func (_ebf Font )SetSize (size float64 ){_ebf ._ccgc .Sz =[]*_ca .CT_FontSize {{ValAttr :size }}};

// GetValueAsNumber retrieves the cell's value as a number
func (_aef Cell )GetValueAsNumber ()(float64 ,error ){if _aef ._ecc .V ==nil &&_aef ._ecc .Is ==nil {return 0,nil ;};if _aef ._ecc .TAttr ==_ca .ST_CellTypeS ||!_ga .IsNumber (*_aef ._ecc .V ){return _db .NaN (),_dec .New ("\u0063\u0065\u006c\u006c\u0020\u0069\u0073\u0020\u006e\u006f\u0074 \u006f\u0066\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020t\u0079\u0070\u0065");};return _fb .ParseFloat (*_aef ._ecc .V ,64);};const (StandardFormatGeneral StandardFormat =0;StandardFormat0 StandardFormat =0;StandardFormatWholeNumber StandardFormat =1;StandardFormat1 StandardFormat =1;StandardFormat2 StandardFormat =2;StandardFormat3 StandardFormat =3;StandardFormat4 StandardFormat =4;StandardFormatPercent StandardFormat =9;StandardFormat9 StandardFormat =9;StandardFormat10 StandardFormat =10;StandardFormat11 StandardFormat =11;StandardFormat12 StandardFormat =12;StandardFormat13 StandardFormat =13;StandardFormatDate StandardFormat =14;StandardFormat14 StandardFormat =14;StandardFormat15 StandardFormat =15;StandardFormat16 StandardFormat =16;StandardFormat17 StandardFormat =17;StandardFormat18 StandardFormat =18;StandardFormatTime StandardFormat =19;StandardFormat19 StandardFormat =19;StandardFormat20 StandardFormat =20;StandardFormat21 StandardFormat =21;StandardFormatDateTime StandardFormat =22;StandardFormat22 StandardFormat =22;StandardFormat37 StandardFormat =37;StandardFormat38 StandardFormat =38;StandardFormat39 StandardFormat =39;StandardFormat40 StandardFormat =40;StandardFormat45 StandardFormat =45;StandardFormat46 StandardFormat =46;StandardFormat47 StandardFormat =47;StandardFormat48 StandardFormat =48;StandardFormat49 StandardFormat =49;);func (_geb Cell )setLocked (_ccc bool ){_fbeg :=_geb ._ecc .SAttr ;if _fbeg !=nil {_adb :=_geb ._fac .StyleSheet .GetCellStyle (*_fbeg );if _adb ._fbg .Protection ==nil {_adb ._fbg .Protection =_ca .NewCT_CellProtection ();};_adb ._fbg .Protection .LockedAttr =&_ccc ;};};func (_fabe StyleSheet )GetNumberFormat (id uint32 )NumberFormat {if id >=0&&id < 50{return CreateDefaultNumberFormat (StandardFormat (id ));};for _ ,_fgdg :=range _fabe ._gace .NumFmts .NumFmt {if _fgdg .NumFmtIdAttr ==id {return NumberFormat {_fabe ._efga ,_fgdg };};};return NumberFormat {};};

// Tables returns a slice of all defined tables in the workbook.
func (_bbbf *Workbook )Tables ()[]Table {if _bbbf ._beee ==nil {return nil ;};_cbafd :=[]Table {};for _ ,_abae :=range _bbbf ._beee {_cbafd =append (_cbafd ,Table {_abae });};return _cbafd ;};

// SaveToFile writes the workbook out to a file.
func (_bdaf *Workbook )SaveToFile (path string )error {_gbef ,_adad :=_g .Create (path );if _adad !=nil {return _adad ;};defer _gbef .Close ();return _bdaf .Save (_gbef );};

// Close closes the workbook, removing any temporary files that might have been
// created when opening a document.
func (_dfbd *Workbook )Close ()error {if _dfbd .TmpPath !=""{return _cdc .RemoveAll (_dfbd .TmpPath );};return nil ;};

// InsertRow inserts a new row into a spreadsheet at a particular row number.  This
// row will now be the row number specified, and any rows after it will be renumbed.
func (_dbd *Sheet )InsertRow (rowNum int )Row {_dbfe :=uint32 (rowNum );for _ ,_gbbb :=range _dbd .Rows (){if _gbbb ._gbgd .RAttr !=nil &&*_gbbb ._gbgd .RAttr >=_dbfe {*_gbbb ._gbgd .RAttr ++;for _ ,_bbg :=range _gbbb .Cells (){_aefc ,_ebea :=_fda .ParseCellReference (_bbg .Reference ());if _ebea !=nil {continue ;};_aefc .RowIdx ++;_bbg ._ecc .RAttr =_b .String (_aefc .String ());};};};for _ ,_cgbab :=range _dbd .MergedCells (){_gefb ,_edggc ,_ddeg :=_fda .ParseRangeReference (_cgbab .Reference ());if _ddeg !=nil {continue ;};if int (_gefb .RowIdx )>=rowNum {_gefb .RowIdx ++;};if int (_edggc .RowIdx )>=rowNum {_edggc .RowIdx ++;};_ggeg :=_e .Sprintf ("\u0025\u0073\u003a%\u0073",_gefb ,_edggc );_cgbab .SetReference (_ggeg );};return _dbd .AddNumberedRow (_dbfe );};type SheetProtection struct{_ebec *_ca .CT_SheetProtection };

// Font allows editing fonts within a spreadsheet stylesheet.
type Font struct{_ccgc *_ca .CT_Font ;_fbdb *_ca .StyleSheet ;};

// X returns the inner wrapped XML type.
func (_gbe RichTextRun )X ()*_ca .CT_RElt {return _gbe ._dfc };func (_fbc Border )SetDiagonal (style _ca .ST_BorderStyle ,c _fde .Color ,up ,down bool ){if _fbc ._fcf .Diagonal ==nil {_fbc ._fcf .Diagonal =_ca .NewCT_BorderPr ();};_fbc ._fcf .Diagonal .Color =_ca .NewCT_Color ();_fbc ._fcf .Diagonal .Color .RgbAttr =c .AsRGBAString ();_fbc ._fcf .Diagonal .StyleAttr =style ;if up {_fbc ._fcf .DiagonalUpAttr =_b .Bool (true );};if down {_fbc ._fcf .DiagonalDownAttr =_b .Bool (true );};};

// RichText is a container for the rich text within a cell. It's similar to a
// paragaraph for a document, except a cell can only contain one rich text item.
type RichText struct{_gef *_ca .CT_Rst };

// SetHeight is a nop-op.
func (_dcc TwoCellAnchor )SetHeight (h _ec .Distance ){};

// NewStyleSheet constructs a new default stylesheet.
func NewStyleSheet (wb *Workbook )StyleSheet {_gdafb :=_ca .NewStyleSheet ();_gdafb .CellStyleXfs =_ca .NewCT_CellStyleXfs ();_gdafb .CellXfs =_ca .NewCT_CellXfs ();_gdafb .CellStyles =_ca .NewCT_CellStyles ();_edcfb :=_ca .NewCT_CellStyle ();_edcfb .NameAttr =_b .String ("\u004e\u006f\u0072\u006d\u0061\u006c");_edcfb .XfIdAttr =0;_edcfb .BuiltinIdAttr =_b .Uint32 (0);_gdafb .CellStyles .CellStyle =append (_gdafb .CellStyles .CellStyle ,_edcfb );_gdafb .CellStyles .CountAttr =_b .Uint32 (uint32 (len (_gdafb .CellStyles .CellStyle )));_geeb :=_ca .NewCT_Xf ();_geeb .NumFmtIdAttr =_b .Uint32 (0);_geeb .FontIdAttr =_b .Uint32 (0);_geeb .FillIdAttr =_b .Uint32 (0);_geeb .BorderIdAttr =_b .Uint32 (0);_gdafb .CellStyleXfs .Xf =append (_gdafb .CellStyleXfs .Xf ,_geeb );_gdafb .CellStyleXfs .CountAttr =_b .Uint32 (uint32 (len (_gdafb .CellStyleXfs .Xf )));_gfcd :=NewFills ();_gdafb .Fills =_gfcd .X ();_cbb :=_gfcd .AddFill ().SetPatternFill ();_cbb .SetPattern (_ca .ST_PatternTypeNone );_cbb =_gfcd .AddFill ().SetPatternFill ();_cbb .SetPattern (_ca .ST_PatternTypeGray125 );_gdafb .Fonts =_ca .NewCT_Fonts ();_gdafb .Borders =_ca .NewCT_Borders ();_dbgg :=StyleSheet {wb ,_gdafb };_dbgg .AddBorder ().InitializeDefaults ();_dfba :=_dbgg .AddFont ();_dfba .SetName ("\u0043a\u006c\u0069\u0062\u0072\u0069");_dfba .SetSize (11);_cgaag :=_ca .NewCT_Xf ();*_cgaag =*_geeb ;_cgaag .XfIdAttr =_b .Uint32 (0);_gdafb .CellXfs .Xf =append (_gdafb .CellXfs .Xf ,_cgaag );_gdafb .CellXfs .CountAttr =_b .Uint32 (uint32 (len (_gdafb .CellXfs .Xf )));return _dbgg ;};

// Comments returns the comments for a sheet.
func (_cfdf *Sheet )Comments ()Comments {for _gbdb ,_dfed :=range _cfdf ._fdfd ._bbced {if _dfed ==_cfdf ._cffg {if _cfdf ._fdfd ._dfgf [_gbdb ]==nil {_cfdf ._fdfd ._dfgf [_gbdb ]=_ca .NewComments ();_cfdf ._fdfd ._dcdg [_gbdb ].AddAutoRelationship (_b .DocTypeSpreadsheet ,_b .WorksheetType ,_gbdb +1,_b .CommentsType );_cfdf ._fdfd .ContentTypes .AddOverride (_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .CommentsType ,_gbdb +1),_b .CommentsContentType );};if len (_cfdf ._fdfd ._dbff )==0{_cfdf ._fdfd ._dbff =append (_cfdf ._fdfd ._dbff ,_bg .NewCommentDrawing ());_bbcg :=_cfdf ._fdfd ._dcdg [_gbdb ].AddAutoRelationship (_b .DocTypeSpreadsheet ,_b .WorksheetType ,1,_b .VMLDrawingType );if _cfdf ._cffg .LegacyDrawing ==nil {_cfdf ._cffg .LegacyDrawing =_ca .NewCT_LegacyDrawing ();};_cfdf ._cffg .LegacyDrawing .IdAttr =_bbcg .ID ();};return Comments {_cfdf ._fdfd ,_cfdf ._fdfd ._dfgf [_gbdb ]};};};_cd .Log .Debug ("\u0061\u0074\u0074\u0065\u006dp\u0074\u0065\u0064\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073 \u0063\u006f\u006d\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u006e\u006f\u006e\u002d\u0065\u0078\u0069\u0073\u0074\u0065\u006e\u0074\u0020\u0073\u0068\u0065\u0065t");return Comments {};};

// CellStyle is a formatting style for a cell. CellStyles are spreadsheet global
// and can be applied to cells across sheets.
type CellStyle struct{_dcb *Workbook ;_fbg *_ca .CT_Xf ;_ddc *_ca .CT_CellXfs ;};

// SetColor sets teh color of the databar.
func (_ebbe DataBarScale )SetColor (c _fde .Color ){_ebbe ._bdcd .Color =_ca .NewCT_Color ();_ebbe ._bdcd .Color .RgbAttr =c .AsRGBAString ();};

// SetDrawing sets the worksheet drawing.  A worksheet can have a reference to a
// single drawing, but the drawing can have many charts.
func (_caebg *Sheet )SetDrawing (d Drawing ){var _gffcc _da .Relationships ;for _bab ,_dagb :=range _caebg ._fdfd ._bbced {if _dagb ==_caebg ._cffg {_gffcc =_caebg ._fdfd ._dcdg [_bab ];break ;};};var _cfa string ;for _eacf ,_egedf :=range d ._fge ._cffgf {if _egedf ==d ._bda {_gbgff :=_gffcc .AddAutoRelationship (_b .DocTypeSpreadsheet ,_b .WorksheetType ,_eacf +1,_b .DrawingType );_cfa =_gbgff .ID ();break ;};};_caebg ._cffg .Drawing =_ca .NewCT_Drawing ();_caebg ._cffg .Drawing .IdAttr =_cfa ;};

// RangeReference converts a range reference of the form 'A1:A5' to 'Sheet
// 1'!$A$1:$A$5 . Renaming a sheet after calculating a range reference will
// invalidate the reference.
func (_cace Sheet )RangeReference (n string )string {_cccdg :=_fa .Split (n ,"\u003a");_dbfb ,_ :=_fda .ParseCellReference (_cccdg [0]);_gabc :=_e .Sprintf ("\u0024\u0025\u0073\u0024\u0025\u0064",_dbfb .Column ,_dbfb .RowIdx );if len (_cccdg )==1{return _e .Sprintf ("\u0027%\u0073\u0027\u0021\u0025\u0073",_cace .Name (),_gabc );};_edga ,_ :=_fda .ParseCellReference (_cccdg [1]);_bbbce :=_e .Sprintf ("\u0024\u0025\u0073\u0024\u0025\u0064",_edga .Column ,_edga .RowIdx );return _e .Sprintf ("\u0027\u0025\u0073\u0027\u0021\u0025\u0073\u003a\u0025\u0073",_cace .Name (),_gabc ,_bbbce );};

// GetBorder gets a Border from a cell style.
func (_adbc CellStyle )GetBorder ()*_ca .CT_Border {if _fgb :=_adbc ._fbg .BorderIdAttr ;_fgb !=nil {_ega :=_adbc ._dcb .StyleSheet .Borders ();if int (*_fgb )< len (_ega ){return _ega [int (*_fgb )].X ();};};return nil ;};

// Extents returns the sheet extents in the form "A1:B15". This requires
// scanning the entire sheet.
func (_afga Sheet )Extents ()string {_fffa ,_acgf ,_bbga ,_agca :=_afga .ExtentsIndex ();return _e .Sprintf ("\u0025s\u0025\u0064\u003a\u0025\u0073\u0025d",_fffa ,_acgf ,_bbga ,_agca );};

// SetDataBar configures the rule as a data bar, removing existing
// configuration.
func (_ffg ConditionalFormattingRule )SetDataBar ()DataBarScale {_ffg .clear ();_ffg .SetType (_ca .ST_CfTypeDataBar );_ffg ._afcg .DataBar =_ca .NewCT_DataBar ();_ade :=DataBarScale {_ffg ._afcg .DataBar };_ade .SetShowValue (true );_ade .SetMinLength (10);_ade .SetMaxLength (90);return _ade ;};

// AddCommentWithStyle adds a new comment styled in a default way
func (_ecd Comments )AddCommentWithStyle (cellRef string ,author string ,comment string )error {_eggb :=_ecd .AddComment (cellRef ,author );_cead :=_eggb .AddRun ();_cead .SetBold (true );_cead .SetSize (10);_cead .SetColor (_fde .Black );_cead .SetFont ("\u0043a\u006c\u0069\u0062\u0072\u0069");_cead .SetText (author +"\u003a");_cead =_eggb .AddRun ();_cead .SetSize (10);_cead .SetFont ("\u0043a\u006c\u0069\u0062\u0072\u0069");_cead .SetColor (_fde .Black );_cead .SetText ("\u000d\u000a"+comment +"\u000d\u000a");_ceag ,_bcfd :=_fda .ParseCellReference (cellRef );if _bcfd !=nil {return _bcfd ;};_ecd ._gebe ._dbff [0].Shape =append (_ecd ._gebe ._dbff [0].Shape ,_bg .NewCommentShape (int64 (_ceag .ColumnIdx ),int64 (_ceag .RowIdx -1)));return nil ;};func (_gee SheetView )ensurePane (){if _gee ._egcb .Pane ==nil {_gee ._egcb .Pane =_ca .NewCT_Pane ();_gee ._egcb .Pane .ActivePaneAttr =_ca .ST_PaneBottomLeft ;};};func _eff ()*_acd .CT_OneCellAnchor {_adef :=_acd .NewCT_OneCellAnchor ();return _adef };func (_fdfb Sheet )IsValid ()bool {return _fdfb ._cffg !=nil };

// GetOrCreateStandardNumberFormat gets or creates a cell style with a given
// standard format. This should only be used when you want to perform
// number/date/time formatting only.  Manipulating the style returned will cause
// all cells using style returned from this for a given format to be formatted.
func (_ceac StyleSheet )GetOrCreateStandardNumberFormat (f StandardFormat )CellStyle {for _ ,_gdaa :=range _ceac .CellStyles (){if _gdaa .HasNumberFormat ()&&_gdaa .NumberFormat ()==uint32 (f ){return _gdaa ;};};_deggc :=_ceac .AddCellStyle ();_deggc .SetNumberFormatStandard (f );return _deggc ;};const _agag ="\u0053\u006fr\u0074\u004f\u0072\u0064e\u0072\u0041s\u0063\u0065\u006e\u0064\u0069\u006e\u0067\u0053o\u0072\u0074\u004f\u0072\u0064\u0065\u0072\u0044\u0065\u0073\u0063\u0065n\u0064\u0069\u006e\u0067";

// SetFgColor sets the *fill* foreground color.  As an example, the solid pattern foreground color becomes the
// background color of the cell when applied.
func (_fffbd PatternFill )SetFgColor (c _fde .Color ){_fffbd ._aedg .FgColor =_ca .NewCT_Color ();_fffbd ._aedg .FgColor .RgbAttr =c .AsRGBAString ();};

// DVCompareType is a comparison type for a data validation rule. This restricts
// the input format of the cell.
type DVCompareType byte ;const _gdaf ="_\u0078\u006c\u006e\u006d._\u0046i\u006c\u0074\u0065\u0072\u0044a\u0074\u0061\u0062\u0061\u0073\u0065";const (AnchorTypeAbsolute AnchorType =iota ;AnchorTypeOneCell ;AnchorTypeTwoCell ;);

// IsBool returns true if the cell is a boolean type cell.
func (_cdg Cell )IsBool ()bool {return _cdg ._ecc .TAttr ==_ca .ST_CellTypeB };

// MergedCells returns the merged cell regions within the sheet.
func (_gccf *Sheet )MergedCells ()[]MergedCell {if _gccf ._cffg .MergeCells ==nil {return nil ;};_caag :=[]MergedCell {};for _ ,_gbea :=range _gccf ._cffg .MergeCells .MergeCell {_caag =append (_caag ,MergedCell {_gccf ._fdfd ,_gccf ,_gbea });};return _caag ;};

// SetCellReference sets the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (_fdee Comment )SetCellReference (cellRef string ){_fdee ._daea .RefAttr =cellRef };var _cega []string =[]string {"\u007a\u0068\u002dH\u004b","\u007a\u0068\u002dM\u004f","\u007a\u0068\u002dC\u004e","\u007a\u0068\u002dS\u0047","\u007a\u0068\u002dT\u0057","\u006a\u0061\u002dJ\u0050","\u006b\u006f\u002dK\u0052"};

// Sheets returns the sheets from the workbook.
func (_gfg *Workbook )Sheets ()[]Sheet {_cbfc :=[]Sheet {};for _agea ,_gcdd :=range _gfg ._bbced {_agdf :=_gfg ._gdeg .Sheets .Sheet [_agea ];_geba :=Sheet {_gfg ,_agdf ,_gcdd };_cbfc =append (_cbfc ,_geba );};return _cbfc ;};

// SetShowValue controls if the cell value is displayed.
func (_ddg DataBarScale )SetShowValue (b bool ){_ddg ._bdcd .ShowValueAttr =_b .Bool (b )};func (_fgc Fills )X ()*_ca .CT_Fills {return _fgc ._bcfe };

// AnchorType is the type of anchor.
type AnchorType byte ;

// DataValidationCompare is a view on a data validation rule that is oriented
// towards value comparisons.
type DataValidationCompare struct{_ecdc *_ca .CT_DataValidation };

// Text returns text from the workbook as one string separated with line breaks.
func (_eeab *WorkbookText )Text ()string {_cccd :=_bf .NewBuffer ([]byte {});for _ ,_bfdac :=range _eeab .Sheets {_cccd .WriteString (_bfdac .Text ());};return _cccd .String ();};func (_efbe StyleSheet )GetCellStyle (id uint32 )CellStyle {for _gfaeb ,_dbdc :=range _efbe ._gace .CellXfs .Xf {if uint32 (_gfaeb )==id {return CellStyle {_efbe ._efga ,_dbdc ,_efbe ._gace .CellXfs };};};return CellStyle {};};

// X returns the inner wrapped XML type.
func (_eed ConditionalFormatting )X ()*_ca .CT_ConditionalFormatting {return _eed ._gfde };var _bbfc =_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .SharedStringsType ,0);

// ClearFont clears any font configuration from the cell style.
func (_bebc CellStyle )ClearFont (){_bebc ._fbg .FontIdAttr =nil ;_bebc ._fbg .ApplyFontAttr =nil };

// TopLeft returns the top-left corner of the anchored object.
func (_afeff OneCellAnchor )TopLeft ()CellMarker {return CellMarker {_afeff ._fdd .From }};

// DataBarScale is a colored scale that fills the cell with a background
// gradeint depending on the value.
type DataBarScale struct{_bdcd *_ca .CT_DataBar };

// Column returns the cell column
func (_beg Cell )Column ()(string ,error ){_dea ,_cda :=_fda .ParseCellReference (_beg .Reference ());if _cda !=nil {return "",_cda ;};return _dea .Column ,nil ;};

// SetShowRuler controls the visibility of the ruler
func (_dff SheetView )SetShowRuler (b bool ){if !b {_dff ._egcb .ShowRulerAttr =_b .Bool (false );}else {_dff ._egcb .ShowRulerAttr =nil ;};};

// Rows returns all of the rows in a sheet.
func (_febb *Sheet )Rows ()[]Row {_ecfa :=[]Row {};for _ ,_efc :=range _febb ._cffg .SheetData .Row {_ecfa =append (_ecfa ,Row {_febb ._fdfd ,_febb ,_efc });};return _ecfa ;};

// Operator returns the operator for the rule
func (_bbd ConditionalFormattingRule )Operator ()_ca .ST_ConditionalFormattingOperator {return _bbd ._afcg .OperatorAttr ;};

// SetHidden hides or unhides the row
func (_ddfab Row )SetHidden (hidden bool ){if !hidden {_ddfab ._gbgd .HiddenAttr =nil ;}else {_ddfab ._gbgd .HiddenAttr =_b .Bool (true );};};

// SetFormat sets the number format code.
func (_aaa NumberFormat )SetFormat (f string ){_aaa ._bgfa .FormatCodeAttr =f };const (DVOpGreater =_ca .ST_DataValidationOperatorGreaterThanOrEqual ;);

// X returns the inner wrapped XML type.
func (_cbde Comments )X ()*_ca .Comments {return _cbde ._cfc };func (_abff *Workbook )ensureSharedStringsRelationships (){_agccf :=false ;for _ ,_acgac :=range _abff .ContentTypes .X ().Override {if _acgac .ContentTypeAttr ==_b .SharedStringsContentType {_agccf =true ;break ;};};if !_agccf {_abff .ContentTypes .AddOverride (_bbfc ,_b .SharedStringsContentType );};_cdcde :=false ;for _ ,_ecccd :=range _abff ._eeead .Relationships (){if _ecccd .X ().TargetAttr ==_adab {_cdcde =true ;break ;};};if !_cdcde {_abff ._eeead .AddRelationship (_adab ,_b .SharedStringsType );};};

// X returns the inner wrapped XML type.
func (_fcfg Font )X ()*_ca .CT_Font {return _fcfg ._ccgc };

// X returns the inner wrapped XML type.
func (_aegg DataValidation )X ()*_ca .CT_DataValidation {return _aegg ._ege };func (_cdca Sheet )ExtentsIndex ()(string ,uint32 ,string ,uint32 ){var _gcg ,_facfa ,_befe ,_cbdb uint32 =1,1,0,0;for _ ,_acefb :=range _cdca .Rows (){if _acefb .RowNumber ()< _gcg {_gcg =_acefb .RowNumber ();}else if _acefb .RowNumber ()> _facfa {_facfa =_acefb .RowNumber ();};for _ ,_fdc :=range _acefb .Cells (){_dggd ,_ddfe :=_fda .ParseCellReference (_fdc .Reference ());if _ddfe ==nil {if _dggd .ColumnIdx < _befe {_befe =_dggd .ColumnIdx ;}else if _dggd .ColumnIdx > _cbdb {_cbdb =_dggd .ColumnIdx ;};};};};return _fda .IndexToColumn (_befe ),_gcg ,_fda .IndexToColumn (_cbdb ),_facfa ;};

// Sheet is a single sheet within a workbook.
type Sheet struct{_fdfd *Workbook ;_aacc *_ca .CT_Sheet ;_cffg *_ca .Worksheet ;};func (_cefc Fill )SetPatternFill ()PatternFill {_cefc ._aae .GradientFill =nil ;_cefc ._aae .PatternFill =_ca .NewCT_PatternFill ();_cefc ._aae .PatternFill .PatternTypeAttr =_ca .ST_PatternTypeSolid ;return PatternFill {_cefc ._aae .PatternFill ,_cefc ._aae };};

// X returns the inner wrapped XML type.
func (_ebabf SheetView )X ()*_ca .CT_SheetView {return _ebabf ._egcb };

// SetNumberWithStyle sets a number and applies a standard format to the cell.
func (_fee Cell )SetNumberWithStyle (v float64 ,f StandardFormat ){_fee .SetNumber (v );_fee .SetStyle (_fee ._fac .StyleSheet .GetOrCreateStandardNumberFormat (f ));};

// CellStyles returns the list of defined cell styles
func (_dcbdb StyleSheet )CellStyles ()[]CellStyle {_fdcf :=[]CellStyle {};for _ ,_faeg :=range _dcbdb ._gace .CellXfs .Xf {_fdcf =append (_fdcf ,CellStyle {_dcbdb ._efga ,_faeg ,_dcbdb ._gace .CellXfs });};return _fdcf ;};

// SetInlineString adds a string inline instead of in the shared strings table.
func (_ad Cell )SetInlineString (s string ){_ad .clearValue ();_ad ._ecc .Is =_ca .NewCT_Rst ();_ad ._ecc .Is .T =_b .String (s );_ad ._ecc .TAttr =_ca .ST_CellTypeInlineStr ;};

// Fills returns a Fills object that can be used to add/create/edit fills.
func (_bfcg StyleSheet )Fills ()Fills {return Fills {_bfcg ._gace .Fills }};func (_fg Cell )getRawSortValue ()(string ,bool ){if _fg .HasFormula (){_gcd :=_fg .GetCachedFormulaResult ();return _gcd ,_ga .IsNumber (_gcd );};_beb ,_ :=_fg .GetRawValue ();return _beb ,_ga .IsNumber (_beb );};

// HasFormula returns true if the cell contains formula.
func (_abde *evalContext )HasFormula (cellRef string )bool {return _abde ._bcdc .Cell (cellRef ).HasFormula ();};func (_gfeg *Sheet )getAllCellsInFormulaArrays (_afdd bool )(map[string ]bool ,error ){_ffdg :=_ab .NewEvaluator ();_dfdgg :=_gfeg .FormulaContext ();_cbab :=map[string ]bool {};for _ ,_eeb :=range _gfeg .Rows (){for _ ,_beec :=range _eeb .Cells (){if _beec .X ().F !=nil {_gfdeg :=_beec .X ().F .Content ;if _beec .X ().F .TAttr ==_ca .ST_CellFormulaTypeArray {_fbbaa :=_ffdg .Eval (_dfdgg ,_gfdeg ).AsString ();if _fbbaa .Type ==_ab .ResultTypeError {_cd .Log .Debug ("\u0065\u0072\u0072o\u0072\u0020\u0065\u0076a\u0075\u006c\u0061\u0074\u0069\u006e\u0067 \u0066\u006f\u0072\u006d\u0075\u006c\u0061\u0020\u0025\u0073\u003a\u0020\u0025\u0073",_gfdeg ,_fbbaa .ErrorMessage );_beec .X ().V =nil ;};if _fbbaa .Type ==_ab .ResultTypeArray {_abbfd ,_fbbaf :=_fda .ParseCellReference (_beec .Reference ());if _fbbaf !=nil {return map[string ]bool {},_fbbaf ;};if (_afdd &&len (_fbbaa .ValueArray )==1)||(!_afdd &&len (_fbbaa .ValueArray [0])==1){continue ;};for _fgga ,_dcacb :=range _fbbaa .ValueArray {_bgac :=_abbfd .RowIdx +uint32 (_fgga );for _deae :=range _dcacb {_edbf :=_fda .IndexToColumn (_abbfd .ColumnIdx +uint32 (_deae ));_cbab [_e .Sprintf ("\u0025\u0073\u0025\u0064",_edbf ,_bgac )]=true ;};};}else if _fbbaa .Type ==_ab .ResultTypeList {_bdg ,_beca :=_fda .ParseCellReference (_beec .Reference ());if _beca !=nil {return map[string ]bool {},_beca ;};if _afdd ||len (_fbbaa .ValueList )==1{continue ;};_abce :=_bdg .RowIdx ;for _abf :=range _fbbaa .ValueList {_aagd :=_fda .IndexToColumn (_bdg .ColumnIdx +uint32 (_abf ));_cbab [_e .Sprintf ("\u0025\u0073\u0025\u0064",_aagd ,_abce )]=true ;};};};};};};return _cbab ,nil ;};

// Content returns the content of the defined range (the range in most cases)/
func (_adc DefinedName )Content ()string {return _adc ._dfg .Content };

// AbsoluteAnchor has a fixed top-left corner in distance units as well as a
// fixed height/width.
type AbsoluteAnchor struct{_ag *_acd .CT_AbsoluteAnchor };

// Type returns the type of the rule
func (_bfd ConditionalFormattingRule )Type ()_ca .ST_CfType {return _bfd ._afcg .TypeAttr };

// SetName sets the sheet name.
func (_eecf *Sheet )SetName (name string ){_eecf ._aacc .NameAttr =name };

// SetActiveSheetIndex sets the index of the active sheet (0-n) which will be
// the tab displayed when the spreadsheet is initially opened.
func (_baegg *Workbook )SetActiveSheetIndex (idx uint32 ){if _baegg ._gdeg .BookViews ==nil {_baegg ._gdeg .BookViews =_ca .NewCT_BookViews ();};if len (_baegg ._gdeg .BookViews .WorkbookView )==0{_baegg ._gdeg .BookViews .WorkbookView =append (_baegg ._gdeg .BookViews .WorkbookView ,_ca .NewCT_BookView ());};_baegg ._gdeg .BookViews .WorkbookView [0].ActiveTabAttr =_b .Uint32 (idx );};

// SetAutoFilter creates autofilters on the sheet. These are the automatic
// filters that are common for a header row.  The RangeRef should be of the form
// "A1:C5" and cover the entire range of cells to be filtered, not just the
// header. SetAutoFilter replaces any existing auto filter on the sheet.
func (_cgf *Sheet )SetAutoFilter (rangeRef string ){rangeRef =_fa .Replace (rangeRef ,"\u0024","",-1);_cgf ._cffg .AutoFilter =_ca .NewCT_AutoFilter ();_cgf ._cffg .AutoFilter .RefAttr =_b .String (rangeRef );_dafea :="\u0027"+_cgf .Name ()+"\u0027\u0021";var _eceb DefinedName ;for _ ,_cacf :=range _cgf ._fdfd .DefinedNames (){if _cacf .Name ()==_gdaf {if _fa .HasPrefix (_cacf .Content (),_dafea ){_eceb =_cacf ;_eceb .SetContent (_cgf .RangeReference (rangeRef ));break ;};};};if _eceb .X ()==nil {_eceb =_cgf ._fdfd .AddDefinedName (_gdaf ,_cgf .RangeReference (rangeRef ));};for _ebbd ,_acef :=range _cgf ._fdfd ._bbced {if _acef ==_cgf ._cffg {_eceb .SetLocalSheetID (uint32 (_ebbd ));};};};

// Name returns the name of the defined name.
func (_acde DefinedName )Name ()string {return _acde ._dfg .NameAttr };

// SetRowOffset sets the row offset of the top-left of the image in fixed units.
func (_dae AbsoluteAnchor )SetRowOffset (m _ec .Distance ){_dae ._ag .Pos .YAttr .ST_CoordinateUnqualified =_b .Int64 (int64 (m /_ec .EMU ));};

// SetHeight sets the row height in points.
func (_caeg Row )SetHeight (d _ec .Distance ){_caeg ._gbgd .HtAttr =_b .Float64 (float64 (d ));_caeg ._gbgd .CustomHeightAttr =_b .Bool (true );};func (_eeag DataValidation )SetComparison (t DVCompareType ,op DVCompareOp )DataValidationCompare {_eeag .clear ();_eeag ._ege .TypeAttr =_ca .ST_DataValidationType (t );_eeag ._ege .OperatorAttr =_ca .ST_DataValidationOperator (op );return DataValidationCompare {_eeag ._ege };};

// SheetCount returns the number of sheets in the workbook.
func (_dbcf Workbook )SheetCount ()int {return len (_dbcf ._bbced )};

// Borders returns the list of borders defined in the stylesheet.
func (_ggfg StyleSheet )Borders ()[]Border {_aadcf :=[]Border {};for _ ,_dgdgd :=range _ggfg ._gace .Borders .Border {_aadcf =append (_aadcf ,Border {_fcf :_dgdgd });};return _aadcf ;};

// ExtractText returns text from the workbook as a WorkbookText object.
func (_efab *Workbook )ExtractText ()*WorkbookText {_cacc :=[]*SheetText {};for _ ,_dee :=range _efab .Sheets (){_cacc =append (_cacc ,&SheetText {Cells :_dee .ExtractText ().Cells });};return &WorkbookText {Sheets :_cacc };};

// SetPassword sets the password hash to a hash of the input password.
func (_fec WorkbookProtection )SetPassword (pw string ){_fec .SetPasswordHash (PasswordHash (pw ))};

// SetHidden marks the defined name as hidden.
func (_bcec DefinedName )SetLocalSheetID (id uint32 ){_bcec ._dfg .LocalSheetIdAttr =_b .Uint32 (id )};

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (_dcda *Sheet )ClearCachedFormulaResults (){for _ ,_bade :=range _dcda .Rows (){for _ ,_fafg :=range _bade .Cells (){if _fafg .X ().F !=nil {_fafg .X ().V =nil ;};};};};

// DefinedName is a named range, formula, etc.
type DefinedName struct{_dfg *_ca .CT_DefinedName };

// AddMergedCells merges cells within a sheet.
func (_gbd *Sheet )AddMergedCells (fromRef ,toRef string )MergedCell {if _gbd ._cffg .MergeCells ==nil {_gbd ._cffg .MergeCells =_ca .NewCT_MergeCells ();};_gbdc :=_ca .NewCT_MergeCell ();_gbdc .RefAttr =_e .Sprintf ("\u0025\u0073\u003a%\u0073",fromRef ,toRef );_gbd ._cffg .MergeCells .MergeCell =append (_gbd ._cffg .MergeCells .MergeCell ,_gbdc );_gbd ._cffg .MergeCells .CountAttr =_b .Uint32 (uint32 (len (_gbd ._cffg .MergeCells .MergeCell )));return MergedCell {_gbd ._fdfd ,_gbd ,_gbdc };};

// SetMinLength sets the minimum bar length in percent.
func (_dbada DataBarScale )SetMinLength (l uint32 ){_dbada ._bdcd .MinLengthAttr =_b .Uint32 (l )};

// ValidateWithPath validates the sheet passing path informaton for a better
// error message
func (_cab Sheet )ValidateWithPath (path string )error {return _cab ._cffg .ValidateWithPath (path )};

// IsDBCS returns if a workbook's default language is among DBCS.
func (_aadg *evalContext )IsDBCS ()bool {_gdc :=_aadg ._bcdc ._fdfd .CoreProperties .X ().Language ;if _gdc ==nil {return false ;};_gfee :=string (_gdc .Data );for _ ,_aeag :=range _cega {if _gfee ==_aeag {return true ;};};return false ;};

// SetState sets the sheet view state (frozen/split/frozen-split)
func (_edeb SheetView )SetState (st _ca .ST_PaneState ){_edeb .ensurePane ();_edeb ._egcb .Pane .StateAttr =st ;};

// Uses1904Dates returns true if the the workbook uses dates relative to
// 1 Jan 1904. This is uncommon.
func (_cadf *Workbook )Uses1904Dates ()bool {if _cadf ._gdeg .WorkbookPr ==nil ||_cadf ._gdeg .WorkbookPr .Date1904Attr ==nil {return false ;};return *_cadf ._gdeg .WorkbookPr .Date1904Attr ;};func (_fffd *Sheet )addNumberedRowFast (_eafbf uint32 )Row {_bceg :=_ca .NewCT_Row ();_bceg .RAttr =_b .Uint32 (_eafbf );_fffd ._cffg .SheetData .Row =append (_fffd ._cffg .SheetData .Row ,_bceg );return Row {_fffd ._fdfd ,_fffd ,_bceg };};func (_fbfb *Sheet )setArray (_aeggg string ,_ecfe _ab .Result )error {_bdab ,_faceb :=_fda .ParseCellReference (_aeggg );if _faceb !=nil {return _faceb ;};for _dcbd ,_cfda :=range _ecfe .ValueArray {_fgbd :=_fbfb .Row (_bdab .RowIdx +uint32 (_dcbd ));for _defb ,_cbec :=range _cfda {_deggb :=_fgbd .Cell (_fda .IndexToColumn (_bdab .ColumnIdx +uint32 (_defb )));if _cbec .Type !=_ab .ResultTypeEmpty {if _cbec .IsBoolean {_deggb .SetBool (_cbec .ValueNumber !=0);}else {_deggb .SetCachedFormulaResult (_cbec .String ());};};};};return nil ;};

// X returns the inner wrapped XML type.
func (_efb Column )X ()*_ca .CT_Col {return _efb ._geg };func (_dbcc PatternFill )ClearFgColor (){_dbcc ._aedg .FgColor =nil };

// SetHyperlink sets a hyperlink on a cell.
func (_gcc Cell )SetHyperlink (hl _da .Hyperlink ){_gaf :=_gcc ._cfd ._cffg ;if _gaf .Hyperlinks ==nil {_gaf .Hyperlinks =_ca .NewCT_Hyperlinks ();};_cge :=_da .Relationship (hl );_ecgf :=_ca .NewCT_Hyperlink ();_ecgf .RefAttr =_gcc .Reference ();_ecgf .IdAttr =_b .String (_cge .ID ());_gaf .Hyperlinks .Hyperlink =append (_gaf .Hyperlinks .Hyperlink ,_ecgf );};func _cbdf ()*_acd .CT_TwoCellAnchor {_acfbg :=_acd .NewCT_TwoCellAnchor ();_acfbg .EditAsAttr =_acd .ST_EditAsOneCell ;_acfbg .From .Col =5;_acfbg .From .Row =0;_acfbg .From .ColOff .ST_CoordinateUnqualified =_b .Int64 (0);_acfbg .From .RowOff .ST_CoordinateUnqualified =_b .Int64 (0);_acfbg .To .Col =10;_acfbg .To .Row =20;_acfbg .To .ColOff .ST_CoordinateUnqualified =_b .Int64 (0);_acfbg .To .RowOff .ST_CoordinateUnqualified =_b .Int64 (0);return _acfbg ;};

// SetActiveSheet sets the active sheet which will be the tab displayed when the
// spreadsheet is initially opened.
func (_bfcb *Workbook )SetActiveSheet (s Sheet ){for _bfbd ,_cfcc :=range _bfcb ._bbced {if s ._cffg ==_cfcc {_bfcb .SetActiveSheetIndex (uint32 (_bfbd ));};};};

// SheetText is an array of extracted text items which has some methods for representing extracted text from a sheet.
type SheetText struct{Cells []CellText ;};func (_gadc Fills )AddFill ()Fill {_fagc :=_ca .NewCT_Fill ();_gadc ._bcfe .Fill =append (_gadc ._bcfe .Fill ,_fagc );_gadc ._bcfe .CountAttr =_b .Uint32 (uint32 (len (_gadc ._bcfe .Fill )));return Fill {_fagc ,_gadc ._bcfe };};type PatternFill struct{_aedg *_ca .CT_PatternFill ;_gfcf *_ca .CT_Fill ;};

// RemoveColumn removes column from the sheet and moves all columns to the right of the removed column one step left.
func (_edd *Sheet )RemoveColumn (column string )error {_edbg ,_facfb :=_edd .getAllCellsInFormulaArraysForColumn ();if _facfb !=nil {return _facfb ;};_gfeef :=_fda .ColumnToIndex (column );for _ ,_ggec :=range _edd .Rows (){_eedg :=_e .Sprintf ("\u0025\u0073\u0025\u0064",column ,*_ggec .X ().RAttr );if _ ,_ffdag :=_edbg [_eedg ];_ffdag {return nil ;};};for _ ,_bega :=range _edd .Rows (){_bcfeg :=_bega ._gbgd .C ;for _ded ,_bbag :=range _bcfeg {_dcdf ,_facd :=_fda .ParseCellReference (*_bbag .RAttr );if _facd !=nil {return _facd ;};if _dcdf .ColumnIdx ==_gfeef {_bega ._gbgd .C =append (_bcfeg [:_ded ],_edd .slideCellsLeft (_bcfeg [_ded +1:])...);break ;}else if _dcdf .ColumnIdx > _gfeef {_bega ._gbgd .C =append (_bcfeg [:_ded ],_edd .slideCellsLeft (_bcfeg [_ded :])...);break ;};};};_facfb =_edd .updateAfterRemove (_gfeef ,_df .UpdateActionRemoveColumn );if _facfb !=nil {return _facfb ;};_facfb =_edd .removeColumnFromNamedRanges (_gfeef );if _facfb !=nil {return _facfb ;};_facfb =_edd .removeColumnFromMergedCells (_gfeef );if _facfb !=nil {return _facfb ;};for _ ,_deb :=range _edd ._fdfd .Sheets (){_deb .RecalculateFormulas ();};return nil ;};

// Comment is a single comment within a sheet.
type Comment struct{_gbf *Workbook ;_daea *_ca .CT_Comment ;_gad *_ca .Comments ;};

// ClearSheetViews clears the list of sheet views.  This will clear the results
// of AddView() or SetFrozen.
func (_aedc *Sheet )ClearSheetViews (){_aedc ._cffg .SheetViews =nil };

// WorkbookText is an array of extracted text items which has some methods for representing extracted text from a workbook.
type WorkbookText struct{Sheets []*SheetText ;};

// X returns the inner wrapped XML type.
func (_ace DefinedName )X ()*_ca .CT_DefinedName {return _ace ._dfg };

// X returns the inner wrapped XML type.
func (_aeed SharedStrings )X ()*_ca .Sst {return _aeed ._cfca };const (_bdbd ="\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061tGe\u006e\u0065\u0072\u0061\u006cS\u0074a\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0057\u0068\u006f\u006ce\u004e\u0075\u006d\u0062\u0065\u0072\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0032\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006da\u0074\u0033\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064F\u006f\u0072\u006d\u0061\u0074\u0034";_bcfc ="\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074P\u0065\u0072\u0063\u0065\u006e\u0074\u0053\u0074\u0061nd\u0061r\u0064F\u006fr\u006d\u0061\u0074\u0031\u0030\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0031\u0031\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064F\u006f\u0072\u006d\u0061\u0074\u0031\u0032\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0033\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0044\u0061\u0074\u0065\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046o\u0072\u006d\u0061\u0074\u00315\u0053\u0074\u0061\u006e\u0064a\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0036\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0037S\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0038\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0054\u0069\u006d\u0065\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u00320\u0053\u0074a\u006e\u0064a\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0032\u0031\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0044\u0061t\u0065\u0054\u0069\u006d\u0065";_ffbd ="\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0033\u0037\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006da\u0074\u0033\u0038\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u00339\u0053\u0074\u0061\u006e\u0064\u0061r\u0064\u0046o\u0072\u006da\u00744\u0030";_cffgc ="\u0053t\u0061\u006e\u0064a\u0072\u0064\u0046o\u0072ma\u0074\u0034\u0035\u0053\u0074\u0061\u006ed\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0034\u0036\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0034\u0037\u0053ta\u006ed\u0061\u0072\u0064\u0046\u006f\u0072m\u0061\u0074\u0034\u0038\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0034\u0039";);func (_ebcc *evalContext )Cell (ref string ,ev _ab .Evaluator )_ab .Result {if !_ccgd (ref ){return _ab .MakeErrorResultType (_ab .ErrorTypeName ,"");};_bgfe :=_ebcc ._bcdc .Name ()+"\u0021"+ref ;if _edc ,_ddd :=ev .GetFromCache (_bgfe );_ddd {return _edc ;};_afebf ,_gfa :=_fda .ParseCellReference (ref );if _gfa !=nil {return _ab .MakeErrorResult (_e .Sprintf ("e\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",ref ,_gfa ));};if _ebcc ._geca !=0&&!_afebf .AbsoluteColumn {_afebf .ColumnIdx +=_ebcc ._geca ;_afebf .Column =_fda .IndexToColumn (_afebf .ColumnIdx );};if _ebcc ._bbea !=0&&!_afebf .AbsoluteRow {_afebf .RowIdx +=_ebcc ._bbea ;};_fedb :=_ebcc ._bcdc .Cell (_afebf .String ());if _fedb .HasFormula (){if _ ,_decg :=_ebcc ._dbad [ref ];_decg {return _ab .MakeErrorResult ("r\u0065\u0063\u0075\u0072\u0073\u0069\u006f\u006e\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0065\u0064\u0020d\u0075\u0072\u0069\u006e\u0067\u0020\u0065\u0076\u0061\u006cua\u0074\u0069\u006fn\u0020o\u0066\u0020"+ref );};_ebcc ._dbad [ref ]=struct{}{};_ddbc :=ev .Eval (_ebcc ,_fedb .GetFormula ());delete (_ebcc ._dbad ,ref );ev .SetCache (_bgfe ,_ddbc );return _ddbc ;};if _fedb .IsEmpty (){_decb :=_ab .MakeEmptyResult ();ev .SetCache (_bgfe ,_decb );return _decb ;}else if _fedb .IsNumber (){_dgac ,_ :=_fedb .GetValueAsNumber ();_ggb :=_ab .MakeNumberResult (_dgac );ev .SetCache (_bgfe ,_ggb );return _ggb ;}else if _fedb .IsBool (){_edf ,_ :=_fedb .GetValueAsBool ();_dgf :=_ab .MakeBoolResult (_edf );ev .SetCache (_bgfe ,_dgf );return _dgf ;};_gebed ,_ :=_fedb .GetRawValue ();if _fedb .IsError (){_gafb :=_ab .MakeErrorResult ("");_gafb .ValueString =_gebed ;ev .SetCache (_bgfe ,_gafb );return _gafb ;};_bgg :=_ab .MakeStringResult (_gebed );ev .SetCache (_bgfe ,_bgg );return _bgg ;};

// SetIcons sets the icon set to use for display.
func (_adbb IconScale )SetIcons (t _ca .ST_IconSetType ){_adbb ._dead .IconSetAttr =t };

// GetString returns the string in a cell if it's an inline or string table
// string. Otherwise it returns an empty string.
func (_gga Cell )GetString ()string {switch _gga ._ecc .TAttr {case _ca .ST_CellTypeInlineStr :if _gga ._ecc .Is !=nil &&_gga ._ecc .Is .T !=nil {return *_gga ._ecc .Is .T ;};if _gga ._ecc .V !=nil {return *_gga ._ecc .V ;};case _ca .ST_CellTypeS :if _gga ._ecc .V ==nil {return "";};_ef ,_aeb :=_fb .Atoi (*_gga ._ecc .V );if _aeb !=nil {return "";};_cca ,_aeb :=_gga ._fac .SharedStrings .GetString (_ef );if _aeb !=nil {return "";};return _cca ;};if _gga ._ecc .V ==nil {return "";};return *_gga ._ecc .V ;};

// SetPasswordHash sets the password hash to the input.
func (_affb WorkbookProtection )SetPasswordHash (pwHash string ){_affb ._acaf .WorkbookPasswordAttr =_b .String (pwHash );};type evalContext struct{_bcdc *Sheet ;_geca ,_bbea uint32 ;_dbad map[string ]struct{};};

// SetRotation configures the cell to be rotated.
func (_fcab CellStyle )SetRotation (deg uint8 ){if _fcab ._fbg .Alignment ==nil {_fcab ._fbg .Alignment =_ca .NewCT_CellAlignment ();};_fcab ._fbg .ApplyAlignmentAttr =_b .Bool (true );_fcab ._fbg .Alignment .TextRotationAttr =_b .Uint8 (deg );};

// Anchor is the interface implemented by anchors. It's modeled after the most
// common anchor (Two cell variant with a from/to position), but will also be
// used for one-cell anchors.  In that case the only non-noop methods are
// TopLeft/MoveTo/SetColOffset/SetRowOffset.
type Anchor interface{

// BottomRight returns the CellMaker for the bottom right corner of the
// anchor.
BottomRight ()CellMarker ;

// TopLeft returns the CellMaker for the top left corner of the anchor.
TopLeft ()CellMarker ;

// MoveTo repositions the anchor without changing the objects size.
MoveTo (_dba ,_dbb int32 );

// SetWidth sets the width of the anchored object. It is not compatible with
// SetWidthCells.
SetWidth (_ee _ec .Distance );

// SetWidthCells sets the height the anchored object by moving the right
// hand side. It is not compatible with SetWidth.
SetWidthCells (_bec int32 );

// SetHeight sets the height of the anchored object. It is not compatible
// with SetHeightCells.
SetHeight (_aa _ec .Distance );

// SetHeightCells sets the height the anchored object by moving the bottom.
// It is not compatible with SetHeight.
SetHeightCells (_fcc int32 );

// SetColOffset sets the column offset of the top-left anchor.
SetColOffset (_ceb _ec .Distance );

// SetRowOffset sets the row offset of the top-left anchor.
SetRowOffset (_afb _ec .Distance );

// Type returns the type of anchor
Type ()AnchorType ;};func (_afde PatternFill )X ()*_ca .CT_PatternFill {return _afde ._aedg };func _dgd (_afd _fd .Time )_fd .Time {_afd =_afd .UTC ();return _fd .Date (_afd .Year (),_afd .Month (),_afd .Day (),_afd .Hour (),_afd .Minute (),_afd .Second (),_afd .Nanosecond (),_fd .Local );};func _efaf (_cgb *Sheet )*evalContext {return &evalContext {_bcdc :_cgb ,_dbad :make (map[string ]struct{})};};

// SetHeightAuto sets the row height to be automatically determined.
func (_eged Row )SetHeightAuto (){_eged ._gbgd .HtAttr =nil ;_eged ._gbgd .CustomHeightAttr =nil };

// Cell retrieves or adds a new cell to a row. Col is the column (e.g. 'A', 'B')
func (_def Row )Cell (col string )Cell {_dadg :=_e .Sprintf ("\u0025\u0073\u0025\u0064",col ,_def .RowNumber ());for _ ,_bdeg :=range _def ._gbgd .C {if _bdeg .RAttr !=nil &&*_bdeg .RAttr ==_dadg {return Cell {_def ._bdaa ,_def ._cgee ,_def ._gbgd ,_bdeg };};};return _def .AddNamedCell (col );};

// GetHorizontalAlignment sets the horizontal alignment of a cell style.
func (_gbb CellStyle )GetHorizontalAlignment ()_ca .ST_HorizontalAlignment {if _gbb ._fbg .Alignment ==nil {return _ca .ST_HorizontalAlignmentUnset ;};return _gbb ._fbg .Alignment .HorizontalAttr ;};

// SetColOffset sets the column offset of the top-left of the image in fixed units.
func (_ba AbsoluteAnchor )SetColOffset (m _ec .Distance ){_ba ._ag .Pos .XAttr .ST_CoordinateUnqualified =_b .Int64 (int64 (m /_ec .EMU ));};

// RecalculateFormulas re-computes any computed formula values that are stored
// in the sheet. As unioffice formula support is still new and not all functins are
// supported, if formula execution fails either due to a parse error or missing
// function, or erorr in the result (even if expected) the cached value will be
// left empty allowing Excel to recompute it on load.
func (_gbeag *Workbook )RecalculateFormulas (){for _ ,_adf :=range _gbeag .Sheets (){_adf .RecalculateFormulas ();};};

// MaxColumnIdx returns the max used column of the sheet.
func (_bbfe Sheet )MaxColumnIdx ()uint32 {_adbcf :=uint32 (0);for _ ,_cga :=range _bbfe .Rows (){_defc :=_cga ._gbgd .C ;if len (_defc )> 0{_eabb :=_defc [len (_defc )-1];_aaef ,_ :=_fda .ParseCellReference (*_eabb .RAttr );if _adbcf < _aaef .ColumnIdx {_adbcf =_aaef .ColumnIdx ;};};};return _adbcf ;};

// X returns the inner wrapped XML type.
func (_ggee Table )X ()*_ca .Table {return _ggee ._gcag };

// X returns the inner wrapped XML type.
func (_baga Drawing )X ()*_acd .WsDr {return _baga ._bda };func (_decc Sheet )validateRowCellNumbers ()error {_cdga :=map[uint32 ]struct{}{};for _ ,_ggcf :=range _decc ._cffg .SheetData .Row {if _ggcf .RAttr !=nil {if _ ,_efd :=_cdga [*_ggcf .RAttr ];_efd {return _e .Errorf ("\u0027%\u0073'\u0020\u0072\u0065\u0075\u0073e\u0064\u0020r\u006f\u0077\u0020\u0025\u0064",_decc .Name (),*_ggcf .RAttr );};_cdga [*_ggcf .RAttr ]=struct{}{};};_cgae :=map[string ]struct{}{};for _ ,_aaga :=range _ggcf .C {if _aaga .RAttr ==nil {continue ;};if _ ,_dddd :=_cgae [*_aaga .RAttr ];_dddd {return _e .Errorf ("\u0027\u0025\u0073\u0027 r\u0065\u0075\u0073\u0065\u0064\u0020\u0063\u0065\u006c\u006c\u0020\u0025\u0073",_decc .Name (),*_aaga .RAttr );};_cgae [*_aaga .RAttr ]=struct{}{};};};return nil ;};

// IsSheetLocked returns whether the sheet is locked.
func (_gcgc SheetProtection )IsSheetLocked ()bool {return _gcgc ._ebec .SheetAttr !=nil &&*_gcgc ._ebec .SheetAttr ;};

// SetIcons configures the rule as an icon scale, removing existing
// configuration.
func (_cdaaf ConditionalFormattingRule )SetIcons ()IconScale {_cdaaf .clear ();_cdaaf .SetType (_ca .ST_CfTypeIconSet );_cdaaf ._afcg .IconSet =_ca .NewCT_IconSet ();_gfc :=IconScale {_cdaaf ._afcg .IconSet };_gfc .SetIcons (_ca .ST_IconSetType3TrafficLights1 );return _gfc ;};

// AddRule adds and returns a new rule that can be configured.
func (_afeb ConditionalFormatting )AddRule ()ConditionalFormattingRule {_ebc :=_ca .NewCT_CfRule ();_afeb ._gfde .CfRule =append (_afeb ._gfde .CfRule ,_ebc );_cbc :=ConditionalFormattingRule {_ebc };_cbc .InitializeDefaults ();_cbc .SetPriority (int32 (len (_afeb ._gfde .CfRule )+1));return _cbc ;};func (_bcag DataValidation )clear (){_bcag ._ege .Formula1 =_b .String ("\u0030");_bcag ._ege .Formula2 =_b .String ("\u0030");};

// SetFrozen removes any existing sheet views and creates a new single view with
// either the first row, first column or both frozen.
func (_bfa *Sheet )SetFrozen (firstRow ,firstCol bool ){_bfa ._cffg .SheetViews =nil ;_aff :=_bfa .AddView ();_aff .SetState (_ca .ST_PaneStateFrozen );switch {case firstRow &&firstCol :_aff .SetYSplit (1);_aff .SetXSplit (1);_aff .SetTopLeft ("\u0042\u0032");case firstRow :_aff .SetYSplit (1);_aff .SetTopLeft ("\u0041\u0032");case firstCol :_aff .SetXSplit (1);_aff .SetTopLeft ("\u0042\u0031");};};func (_aded *Sheet )getAllCellsInFormulaArraysForColumn ()(map[string ]bool ,error ){return _aded .getAllCellsInFormulaArrays (false );};

// GetLocked returns true if the cell is locked.
func (_fcgb *evalContext )GetLocked (cellRef string )bool {return _fcgb ._bcdc .Cell (cellRef ).getLocked ();};

// SetRichTextString sets the cell to rich string mode and returns a struct that
// can be used to add formatted text to the cell.
func (_feb Cell )SetRichTextString ()RichText {_feb .clearValue ();_feb ._ecc .Is =_ca .NewCT_Rst ();_feb ._ecc .TAttr =_ca .ST_CellTypeInlineStr ;return RichText {_feb ._ecc .Is };};

// GetValueAsTime retrieves the cell's value as a time.  There is no difference
// in SpreadsheetML between a time/date cell other than formatting, and that
// typically a date cell won't have a fractional component. GetValueAsTime will
// work for date cells as well.
func (_egf Cell )GetValueAsTime ()(_fd .Time ,error ){if _egf ._ecc .TAttr !=_ca .ST_CellTypeUnset {return _fd .Time {},_dec .New ("\u0063e\u006c\u006c\u0020\u0074y\u0070\u0065\u0020\u0073\u0068o\u0075l\u0064 \u0062\u0065\u0020\u0075\u006e\u0073\u0065t");};if _egf ._ecc .V ==nil {return _fd .Time {},_dec .New ("\u0063\u0065\u006c\u006c\u0020\u0068\u0061\u0073\u0020\u006e\u006f\u0020v\u0061\u006c\u0075\u0065");};_deaa ,_ ,_ecg :=_ac .ParseFloat (*_egf ._ecc .V ,10,128,_ac .ToNearestEven );if _ecg !=nil {return _fd .Time {},_ecg ;};_aec :=new (_ac .Float );_aec .SetUint64 (uint64 (24*_fd .Hour ));_deaa .Mul (_deaa ,_aec );_ece ,_ :=_deaa .Uint64 ();_aba :=_egf ._fac .Epoch ().Add (_fd .Duration (_ece ));return _dgd (_aba ),nil ;};

// CopySheet copies the existing sheet at index `ind` and puts its copy with the name `copiedSheetName`.
func (_dffe *Workbook )CopySheet (ind int ,copiedSheetName string )(Sheet ,error ){if _dffe .SheetCount ()<=ind {return Sheet {},ErrorNotFound ;};var _gbfe _da .Relationship ;for _ ,_fcfe :=range _dffe ._eeead .Relationships (){if _fcfe .ID ()==_dffe ._gdeg .Sheets .Sheet [ind ].IdAttr {var _fede bool ;if _gbfe ,_fede =_dffe ._eeead .CopyRelationship (_fcfe .ID ());!_fede {return Sheet {},ErrorNotFound ;};break ;};};_dffe .ContentTypes .CopyOverride (_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .WorksheetContentType ,ind +1),_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .WorksheetContentType ,len (_dffe .ContentTypes .X ().Override )));_dfgfd :=*_dffe ._bbced [ind ];_dffe ._bbced =append (_dffe ._bbced ,&_dfgfd );var _daca uint32 =0;for _ ,_acgae :=range _dffe ._gdeg .Sheets .Sheet {if _acgae .SheetIdAttr > _daca {_daca =_acgae .SheetIdAttr ;};};_daca ++;_ebce :=*_dffe ._gdeg .Sheets .Sheet [ind ];_ebce .IdAttr =_gbfe .ID ();_ebce .NameAttr =copiedSheetName ;_ebce .SheetIdAttr =_daca ;_dffe ._gdeg .Sheets .Sheet =append (_dffe ._gdeg .Sheets .Sheet ,&_ebce );_agd :=_da .NewRelationshipsCopy (_dffe ._dcdg [ind ]);_dffe ._dcdg =append (_dffe ._dcdg ,_agd );_ddba :=_dffe ._dfgf [ind ];if _ddba ==nil {_dffe ._dfgf =append (_dffe ._dfgf ,nil );}else {_aedab :=*_ddba ;_dffe ._dfgf =append (_dffe ._dfgf ,&_aedab );};_dcga :=Sheet {_dffe ,&_ebce ,&_dfgfd };return _dcga ,nil ;};

// IsError returns true if the cell is an error type cell.
func (_aac Cell )IsError ()bool {return _aac ._ecc .TAttr ==_ca .ST_CellTypeE };

// AddRun adds a new run of text to the cell.
func (_cgba RichText )AddRun ()RichTextRun {_fdef :=_ca .NewCT_RElt ();_cgba ._gef .R =append (_cgba ._gef .R ,_fdef );return RichTextRun {_fdef };};

// X returns the inner wrapped XML type.
func (_faaa WorkbookProtection )X ()*_ca .CT_WorkbookProtection {return _faaa ._acaf };

// BottomRight is a no-op.
func (_aee AbsoluteAnchor )BottomRight ()CellMarker {return CellMarker {}};

// X returns the inner XML entity for a stylesheet.
func (_cdgc StyleSheet )X ()*_ca .StyleSheet {return _cdgc ._gace };

// IsEmpty checks if the cell style contains nothing.
func (_acg CellStyle )IsEmpty ()bool {return _acg ._dcb ==nil ||_acg ._fbg ==nil ||_acg ._ddc ==nil ||_acg ._ddc .Xf ==nil ;};

// SetColOffset sets the column offset of the top-left anchor.
func (_bggf OneCellAnchor )SetColOffset (m _ec .Distance ){_bggf .TopLeft ().SetColOffset (m )};var _ceeg *_bd .Regexp =_bd .MustCompile ("\u005e(\u005ba\u002d\u007a\u005d\u002b\u0029(\u005b\u0030-\u0039\u005d\u002b\u0029\u0024");func (_bcad CellStyle )SetShrinkToFit (b bool ){if _bcad ._fbg .Alignment ==nil {_bcad ._fbg .Alignment =_ca .NewCT_CellAlignment ();};_bcad ._fbg .ApplyAlignmentAttr =_b .Bool (true );if !b {_bcad ._fbg .Alignment .ShrinkToFitAttr =nil ;}else {_bcad ._fbg .Alignment .ShrinkToFitAttr =_b .Bool (b );};};

// LessRows compares two rows based off of a column. If the column doesn't exist
// in one row, that row is 'less'.
func (_aag Comparer )LessRows (column string ,lhs ,rhs Row )bool {var _dad ,_dafe Cell ;for _ ,_agcc :=range lhs .Cells (){_ddf ,_ :=_fda .ParseCellReference (_agcc .Reference ());if _ddf .Column ==column {_dad =_agcc ;break ;};};for _ ,_dafb :=range rhs .Cells (){_bed ,_ :=_fda .ParseCellReference (_dafb .Reference ());if _bed .Column ==column {_dafe =_dafb ;break ;};};return _aag .LessCells (_dad ,_dafe );};

// GetString retrieves a string from the shared strings table by index.
func (_ggbb SharedStrings )GetString (id int )(string ,error ){if id < 0{return "",_e .Errorf ("\u0069\u006eva\u006c\u0069\u0064 \u0073\u0074\u0072\u0069ng \u0069nd\u0065\u0078\u0020\u0025\u0064\u002c\u0020mu\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030",id );};if id > len (_ggbb ._cfca .Si ){return "",_e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002c\u0020\u0074\u0061b\u006c\u0065\u0020\u006f\u006e\u006c\u0079\u0020\u0068\u0061\u0073\u0020\u0025\u0064 \u0076a\u006c\u0075\u0065\u0073",id ,len (_ggbb ._cfca .Si ));};_bfeb :=_ggbb ._cfca .Si [id ];if _bfeb .T !=nil {return *_bfeb .T ,nil ;};return "",nil ;};

// AddDrawing adds a drawing to a workbook.  However the drawing is not actually
// displayed or used until it's set on a sheet.
func (_acfe *Workbook )AddDrawing ()Drawing {_dgda :=_acd .NewWsDr ();_acfe ._cffgf =append (_acfe ._cffgf ,_dgda );_dbcd :=_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .DrawingType ,len (_acfe ._cffgf ));_acfe .ContentTypes .AddOverride (_dbcd ,_b .DrawingContentType );_acfe ._aeec =append (_acfe ._aeec ,_da .NewRelationships ());return Drawing {_acfe ,_dgda };};type Table struct{_gcag *_ca .Table };

// New constructs a new workbook.
func New ()*Workbook {_ecge :=&Workbook {};_ecge ._gdeg =_ca .NewWorkbook ();_ecge .AppProperties =_da .NewAppProperties ();_ecge .CoreProperties =_da .NewCoreProperties ();_ecge .StyleSheet =NewStyleSheet (_ecge );_ecge .Rels =_da .NewRelationships ();_ecge ._eeead =_da .NewRelationships ();_ecge .Rels .AddRelationship (_b .RelativeFilename (_b .DocTypeSpreadsheet ,"",_b .ExtendedPropertiesType ,0),_b .ExtendedPropertiesType );_ecge .Rels .AddRelationship (_b .RelativeFilename (_b .DocTypeSpreadsheet ,"",_b .CorePropertiesType ,0),_b .CorePropertiesType );_ecge .Rels .AddRelationship (_b .RelativeFilename (_b .DocTypeSpreadsheet ,"",_b .OfficeDocumentType ,0),_b .OfficeDocumentType );_ecge ._eeead .AddRelationship (_b .RelativeFilename (_b .DocTypeSpreadsheet ,_b .OfficeDocumentType ,_b .StylesType ,0),_b .StylesType );_ecge .ContentTypes =_da .NewContentTypes ();_ecge .ContentTypes .AddDefault ("\u0076\u006d\u006c",_b .VMLDrawingContentType );_ecge .ContentTypes .AddOverride (_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .OfficeDocumentType ,0),"\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066o\u0072\u006d\u0061\u0074s\u002d\u006f\u0066\u0066\u0069\u0063e\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068e\u0065\u0074\u006d\u006c\u002e\u0073\u0068\u0065\u0065\u0074\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c");_ecge .ContentTypes .AddOverride (_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .StylesType ,0),_b .SMLStyleSheetContentType );_ecge .SharedStrings =NewSharedStrings ();_ecge .ContentTypes .AddOverride (_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .SharedStringsType ,0),_b .SharedStringsContentType );_ecge ._eeead .AddRelationship (_b .RelativeFilename (_b .DocTypeSpreadsheet ,_b .OfficeDocumentType ,_b .SharedStringsType ,0),_b .SharedStringsType );_ecge ._gdae =map[string ]string {};return _ecge ;};func CreateDefaultNumberFormat (id StandardFormat )NumberFormat {_cccf :=NumberFormat {_bgfa :_ca .NewCT_NumFmt ()};_cccf ._bgfa .NumFmtIdAttr =uint32 (id );_cccf ._bgfa .FormatCodeAttr ="\u0047e\u006e\u0065\u0072\u0061\u006c";switch id {case StandardFormat0 :_cccf ._bgfa .FormatCodeAttr ="\u0047e\u006e\u0065\u0072\u0061\u006c";case StandardFormat1 :_cccf ._bgfa .FormatCodeAttr ="\u0030";case StandardFormat2 :_cccf ._bgfa .FormatCodeAttr ="\u0030\u002e\u0030\u0030";case StandardFormat3 :_cccf ._bgfa .FormatCodeAttr ="\u0023\u002c\u0023#\u0030";case StandardFormat4 :_cccf ._bgfa .FormatCodeAttr ="\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030";case StandardFormat9 :_cccf ._bgfa .FormatCodeAttr ="\u0030\u0025";case StandardFormat10 :_cccf ._bgfa .FormatCodeAttr ="\u0030\u002e\u00300\u0025";case StandardFormat11 :_cccf ._bgfa .FormatCodeAttr ="\u0030\u002e\u0030\u0030\u0045\u002b\u0030\u0030";case StandardFormat12 :_cccf ._bgfa .FormatCodeAttr ="\u0023\u0020\u003f/\u003f";case StandardFormat13 :_cccf ._bgfa .FormatCodeAttr ="\u0023 \u003f\u003f\u002f\u003f\u003f";case StandardFormat14 :_cccf ._bgfa .FormatCodeAttr ="\u006d\u002f\u0064\u002f\u0079\u0079";case StandardFormat15 :_cccf ._bgfa .FormatCodeAttr ="\u0064\u002d\u006d\u006d\u006d\u002d\u0079\u0079";case StandardFormat16 :_cccf ._bgfa .FormatCodeAttr ="\u0064\u002d\u006dm\u006d";case StandardFormat17 :_cccf ._bgfa .FormatCodeAttr ="\u006d\u006d\u006d\u002d\u0079\u0079";case StandardFormat18 :_cccf ._bgfa .FormatCodeAttr ="\u0068\u003a\u006d\u006d\u0020\u0041\u004d\u002f\u0050\u004d";case StandardFormat19 :_cccf ._bgfa .FormatCodeAttr ="\u0068\u003a\u006d\u006d\u003a\u0073\u0073\u0020\u0041\u004d\u002f\u0050\u004d";case StandardFormat20 :_cccf ._bgfa .FormatCodeAttr ="\u0068\u003a\u006d\u006d";case StandardFormat21 :_cccf ._bgfa .FormatCodeAttr ="\u0068:\u006d\u006d\u003a\u0073\u0073";case StandardFormat22 :_cccf ._bgfa .FormatCodeAttr ="m\u002f\u0064\u002f\u0079\u0079\u0020\u0068\u003a\u006d\u006d";case StandardFormat37 :_cccf ._bgfa .FormatCodeAttr ="\u0023\u002c\u0023\u0023\u0030\u0020\u003b\u0028\u0023,\u0023\u0023\u0030\u0029";case StandardFormat38 :_cccf ._bgfa .FormatCodeAttr ="\u0023\u002c\u0023\u00230 \u003b\u005b\u0052\u0065\u0064\u005d\u0028\u0023\u002c\u0023\u0023\u0030\u0029";case StandardFormat39 :_cccf ._bgfa .FormatCodeAttr ="\u0023\u002c\u0023\u00230.\u0030\u0030\u003b\u0028\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030\u0029";case StandardFormat40 :_cccf ._bgfa .FormatCodeAttr ="\u0023,\u0023\u0023\u0030\u002e\u0030\u0030\u003b\u005b\u0052\u0065\u0064]\u0028\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030\u0029";case StandardFormat45 :_cccf ._bgfa .FormatCodeAttr ="\u006d\u006d\u003as\u0073";case StandardFormat46 :_cccf ._bgfa .FormatCodeAttr ="\u005bh\u005d\u003a\u006d\u006d\u003a\u0073s";case StandardFormat47 :_cccf ._bgfa .FormatCodeAttr ="\u006dm\u003a\u0073\u0073\u002e\u0030";case StandardFormat48 :_cccf ._bgfa .FormatCodeAttr ="\u0023\u0023\u0030\u002e\u0030\u0045\u002b\u0030";case StandardFormat49 :_cccf ._bgfa .FormatCodeAttr ="\u0040";};return _cccf ;};

// Index returns the index of the border for use with a cell style.
func (_bcd Border )Index ()uint32 {for _cg ,_ebb :=range _bcd ._cbd .Border {if _ebb ==_bcd ._fcf {return uint32 (_cg );};};return 0;};

// AddGradientStop adds a color gradient stop.
func (_ddb ColorScale )AddGradientStop (color _fde .Color ){_daad :=_ca .NewCT_Color ();_daad .RgbAttr =color .AsRGBAString ();_ddb ._dfdc .Color =append (_ddb ._dfdc .Color ,_daad );};

// SetBorder is a helper function for creating borders across multiple cells. In
// the OOXML spreadsheet format, a border applies to a single cell.  To draw a
// 'boxed' border around multiple cells, you need to apply different styles to
// the cells on the top,left,right,bottom and four corners.  This function
// breaks apart a single border into its components and applies it to cells as
// needed to give the effect of a border applying to multiple cells.
func (_eae *Sheet )SetBorder (cellRange string ,border Border )error {_gcee ,_fcga ,_ddag :=_fda .ParseRangeReference (cellRange );if _ddag !=nil {return _ddag ;};_fagb :=_eae ._fdfd .StyleSheet .AddCellStyle ();_gdb :=_eae ._fdfd .StyleSheet .AddBorder ();_fagb .SetBorder (_gdb );_gdb ._fcf .Top =border ._fcf .Top ;_gdb ._fcf .Left =border ._fcf .Left ;_eadg :=_eae ._fdfd .StyleSheet .AddCellStyle ();_edcf :=_eae ._fdfd .StyleSheet .AddBorder ();_eadg .SetBorder (_edcf );_edcf ._fcf .Top =border ._fcf .Top ;_edcf ._fcf .Right =border ._fcf .Right ;_bdce :=_eae ._fdfd .StyleSheet .AddCellStyle ();_cfaf :=_eae ._fdfd .StyleSheet .AddBorder ();_bdce .SetBorder (_cfaf );_cfaf ._fcf .Top =border ._fcf .Top ;_fgfa :=_eae ._fdfd .StyleSheet .AddCellStyle ();_dcab :=_eae ._fdfd .StyleSheet .AddBorder ();_fgfa .SetBorder (_dcab );_dcab ._fcf .Left =border ._fcf .Left ;_dbaff :=_eae ._fdfd .StyleSheet .AddCellStyle ();_gddd :=_eae ._fdfd .StyleSheet .AddBorder ();_dbaff .SetBorder (_gddd );_gddd ._fcf .Right =border ._fcf .Right ;_fdecc :=_eae ._fdfd .StyleSheet .AddCellStyle ();_ebaa :=_eae ._fdfd .StyleSheet .AddBorder ();_fdecc .SetBorder (_ebaa );_ebaa ._fcf .Bottom =border ._fcf .Bottom ;_bcgf :=_eae ._fdfd .StyleSheet .AddCellStyle ();_fea :=_eae ._fdfd .StyleSheet .AddBorder ();_bcgf .SetBorder (_fea );_fea ._fcf .Bottom =border ._fcf .Bottom ;_fea ._fcf .Left =border ._fcf .Left ;_dbg :=_eae ._fdfd .StyleSheet .AddCellStyle ();_bcaf :=_eae ._fdfd .StyleSheet .AddBorder ();_dbg .SetBorder (_bcaf );_bcaf ._fcf .Bottom =border ._fcf .Bottom ;_bcaf ._fcf .Right =border ._fcf .Right ;_cgg :=_gcee .RowIdx ;_ffbf :=_gcee .ColumnIdx ;_fcggb :=_fcga .RowIdx ;_dfgc :=_fcga .ColumnIdx ;for _ccac :=_cgg ;_ccac <=_fcggb ;_ccac ++{for _edb :=_ffbf ;_edb <=_dfgc ;_edb ++{_efgg :=_e .Sprintf ("\u0025\u0073\u0025\u0064",_fda .IndexToColumn (_edb ),_ccac );switch {case _ccac ==_cgg &&_edb ==_ffbf :_eae .Cell (_efgg ).SetStyle (_fagb );case _ccac ==_cgg &&_edb ==_dfgc :_eae .Cell (_efgg ).SetStyle (_eadg );case _ccac ==_fcggb &&_edb ==_ffbf :_eae .Cell (_efgg ).SetStyle (_bcgf );case _ccac ==_fcggb &&_edb ==_dfgc :_eae .Cell (_efgg ).SetStyle (_dbg );case _ccac ==_cgg :_eae .Cell (_efgg ).SetStyle (_bdce );case _ccac ==_fcggb :_eae .Cell (_efgg ).SetStyle (_fdecc );case _edb ==_ffbf :_eae .Cell (_efgg ).SetStyle (_fgfa );case _edb ==_dfgc :_eae .Cell (_efgg ).SetStyle (_dbaff );};};};return nil ;};

// GetLabelPrefix returns label prefix which depends on the cell's horizontal alignment.
func (_bddc *evalContext )GetLabelPrefix (cellRef string )string {return _bddc ._bcdc .Cell (cellRef ).getLabelPrefix ();};

// Protection allows control over the workbook protections.
func (_aaffg *Workbook )Protection ()WorkbookProtection {if _aaffg ._gdeg .WorkbookProtection ==nil {_aaffg ._gdeg .WorkbookProtection =_ca .NewCT_WorkbookProtection ();};return WorkbookProtection {_aaffg ._gdeg .WorkbookProtection };};

// AddDataValidation adds a data validation rule to a sheet.
func (_bgda *Sheet )AddDataValidation ()DataValidation {if _bgda ._cffg .DataValidations ==nil {_bgda ._cffg .DataValidations =_ca .NewCT_DataValidations ();};_ccfg :=_ca .NewCT_DataValidation ();_ccfg .ShowErrorMessageAttr =_b .Bool (true );_bgda ._cffg .DataValidations .DataValidation =append (_bgda ._cffg .DataValidations .DataValidation ,_ccfg );_bgda ._cffg .DataValidations .CountAttr =_b .Uint32 (uint32 (len (_bgda ._cffg .DataValidations .DataValidation )));return DataValidation {_ccfg };};

// Index returns the index of the differential style.
func (_fffb DifferentialStyle )Index ()uint32 {for _ceda ,_cfe :=range _fffb ._ddca .Dxf {if _fffb ._ccad ==_cfe {return uint32 (_ceda );};};return 0;};

// SetStyle sets the style to be used for conditional rules
func (_gbbe ConditionalFormattingRule )SetStyle (d DifferentialStyle ){_gbbe ._afcg .DxfIdAttr =_b .Uint32 (d .Index ());};

// ClearNumberFormat removes any number formatting from the style.
func (_bccf CellStyle )ClearNumberFormat (){_bccf ._fbg .NumFmtIdAttr =nil ;_bccf ._fbg .ApplyNumberFormatAttr =nil ;};

// SetTopLeft sets the top left visible cell after the split.
func (_cegaa SheetView )SetTopLeft (cellRef string ){_cegaa .ensurePane ();_cegaa ._egcb .Pane .TopLeftCellAttr =&cellRef ;};func _ccab (_fdf bool )int {if _fdf {return 1;};return 0;};func (_deac *Workbook )onNewRelationship (_ggfa *_ce .DecodeMap ,_geeg ,_dbaa string ,_fgfd []*_gg .File ,_bcdb *_cef .Relationship ,_aaaa _ce .Target )error {_deag :=_b .DocTypeSpreadsheet ;switch _dbaa {case _b .OfficeDocumentType :_deac ._gdeg =_ca .NewWorkbook ();_ggfa .AddTarget (_geeg ,_deac ._gdeg ,_dbaa ,0);_deac ._eeead =_da .NewRelationships ();_ggfa .AddTarget (_ce .RelationsPathFor (_geeg ),_deac ._eeead .X (),_dbaa ,0);_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,0);case _b .CorePropertiesType :_ggfa .AddTarget (_geeg ,_deac .CoreProperties .X (),_dbaa ,0);_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,0);case _b .CustomPropertiesType :_ggfa .AddTarget (_geeg ,_deac .CustomProperties .X (),_dbaa ,0);_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,0);case _b .ExtendedPropertiesType :_ggfa .AddTarget (_geeg ,_deac .AppProperties .X (),_dbaa ,0);_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,0);case _b .WorksheetType :_dfdd :=_ca .NewWorksheet ();_daef :=uint32 (len (_deac ._bbced ));_deac ._bbced =append (_deac ._bbced ,_dfdd );_ggfa .AddTarget (_geeg ,_dfdd ,_dbaa ,_daef );_abfb :=_da .NewRelationships ();_ggfa .AddTarget (_ce .RelationsPathFor (_geeg ),_abfb .X (),_dbaa ,0);_deac ._dcdg =append (_deac ._dcdg ,_abfb );_deac ._dfgf =append (_deac ._dfgf ,nil );_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,len (_deac ._bbced ));case _b .StylesType :_deac .StyleSheet =NewStyleSheet (_deac );_ggfa .AddTarget (_geeg ,_deac .StyleSheet .X (),_dbaa ,0);_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,0);case _b .ThemeType :_daeg :=_bgc .NewTheme ();_deac ._gfcg =append (_deac ._gfcg ,_daeg );_ggfa .AddTarget (_geeg ,_daeg ,_dbaa ,0);_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,len (_deac ._gfcg ));case _b .SharedStringsType :_deac .SharedStrings =NewSharedStrings ();_ggfa .AddTarget (_geeg ,_deac .SharedStrings .X (),_dbaa ,0);_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,0);case _b .ThumbnailType :for _gcfe ,_gdbf :=range _fgfd {if _gdbf ==nil {continue ;};if _gdbf .Name ==_geeg {_bcgc ,_dbga :=_gdbf .Open ();if _dbga !=nil {return _e .Errorf ("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073",_dbga );};_deac .Thumbnail ,_ ,_dbga =_bc .Decode (_bcgc );_bcgc .Close ();if _dbga !=nil {return _e .Errorf ("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073",_dbga );};_fgfd [_gcfe ]=nil ;};};case _b .ImageType :for _cada ,_gecfc :=range _deac ._gdae {_ffbg :=_de .Clean (_geeg );if _ffbg ==_cada {_bcdb .TargetAttr =_gecfc ;return nil ;};};_deaf :=_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,len (_deac .Images )+1);for _gfaf ,_eacb :=range _fgfd {if _eacb ==nil {continue ;};if _eacb .Name ==_de .Clean (_geeg ){_eddd ,_ebffd :=_ce .ExtractToDiskTmp (_eacb ,_deac .TmpPath );if _ebffd !=nil {return _ebffd ;};_ccegb ,_ebffd :=_da .ImageFromStorage (_eddd );if _ebffd !=nil {return _ebffd ;};_bfgf :=_da .MakeImageRef (_ccegb ,&_deac .DocBase ,_deac ._eeead );_bfgf .SetTarget (_deaf );_deac ._gdae [_eacb .Name ]=_deaf ;_deac .Images =append (_deac .Images ,_bfgf );_fgfd [_gfaf ]=nil ;};};_bcdb .TargetAttr =_deaf ;case _b .DrawingType :_fdcd :=_acd .NewWsDr ();_abgc :=uint32 (len (_deac ._cffgf ));_ggfa .AddTarget (_geeg ,_fdcd ,_dbaa ,_abgc );_deac ._cffgf =append (_deac ._cffgf ,_fdcd );_dagf :=_da .NewRelationships ();_ggfa .AddTarget (_ce .RelationsPathFor (_geeg ),_dagf .X (),_dbaa ,_abgc );_deac ._aeec =append (_deac ._aeec ,_dagf );_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,len (_deac ._cffgf ));case _b .VMLDrawingType :_eabg :=_bg .NewContainer ();_egeg :=uint32 (len (_deac ._dbff ));_ggfa .AddTarget (_geeg ,_eabg ,_dbaa ,_egeg );_deac ._dbff =append (_deac ._dbff ,_eabg );case _b .CommentsType :_deac ._dfgf [_aaaa .Index ]=_ca .NewComments ();_ggfa .AddTarget (_geeg ,_deac ._dfgf [_aaaa .Index ],_dbaa ,_aaaa .Index );_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,len (_deac ._dfgf ));case _b .ChartType :_gagee :=_fc .NewChartSpace ();_faccd :=uint32 (len (_deac ._dgeg ));_ggfa .AddTarget (_geeg ,_gagee ,_dbaa ,_faccd );_deac ._dgeg =append (_deac ._dgeg ,_gagee );_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,len (_deac ._dgeg ));_deac ._gcbc [_bcdb .TargetAttr ]=_gagee ;case _b .TableType :_gffbb :=_ca .NewTable ();_abda :=uint32 (len (_deac ._beee ));_ggfa .AddTarget (_geeg ,_gffbb ,_dbaa ,_abda );_deac ._beee =append (_deac ._beee ,_gffbb );_bcdb .TargetAttr =_b .RelativeFilename (_deag ,_aaaa .Typ ,_dbaa ,len (_deac ._beee ));default:_cd .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0072\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068\u0069\u0070\u0020\u0025\u0073\u0020\u0025\u0073",_geeg ,_dbaa );};return nil ;};

// Open opens and reads a workbook from a file (.xlsx).
func Open (filename string )(*Workbook ,error ){_gbbed ,_edge :=_g .Open (filename );if _edge !=nil {return nil ,_e .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",filename ,_edge );};defer _gbbed .Close ();_fffg ,_edge :=_g .Stat (filename );if _edge !=nil {return nil ,_e .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",filename ,_edge );};_gffa ,_edge :=Read (_gbbed ,_fffg .Size ());if _edge !=nil {return nil ,_edge ;};_gdce ,_ :=_a .Abs (_a .Dir (filename ));_gffa ._eeaf =_a .Join (_gdce ,filename );return _gffa ,nil ;};

// ClearBorder clears any border configuration from the cell style.
func (_bbf CellStyle )ClearBorder (){_bbf ._fbg .BorderIdAttr =nil ;_bbf ._fbg .ApplyBorderAttr =nil };

// SetWrapped configures the cell to wrap text.
func (_febd CellStyle )SetWrapped (b bool ){if _febd ._fbg .Alignment ==nil {_febd ._fbg .Alignment =_ca .NewCT_CellAlignment ();};if !b {_febd ._fbg .Alignment .WrapTextAttr =nil ;}else {_febd ._fbg .Alignment .WrapTextAttr =_b .Bool (true );_febd ._fbg .ApplyAlignmentAttr =_b .Bool (true );};};

// SetWidthCells sets the height the anchored object by moving the right hand
// side. It is not compatible with SetWidth.
func (_aaac TwoCellAnchor )SetWidthCells (w int32 ){_aggdf :=_aaac .TopLeft ();_abef :=_aaac .BottomRight ();_abef .SetCol (_aggdf .Col ()+w );};

// GetChartByTargetId returns the array of workbook crt.ChartSpace.
func (_bdga *Workbook )GetChartByTargetId (targetAttr string )*_fc .ChartSpace {return _bdga ._gcbc [targetAttr ];};

// IsSheetLocked returns whether the sheet objects are locked.
func (_begg SheetProtection )IsObjectLocked ()bool {return _begg ._ebec .ObjectsAttr !=nil &&*_begg ._ebec .ObjectsAttr ;};type MergedCell struct{_eba *Workbook ;_dgff *Sheet ;_fgf *_ca .CT_MergeCell ;};

// SetNumber sets the cell type to number, and the value to the given number
func (_fbd Cell )SetNumber (v float64 ){_fbd .clearValue ();if _db .IsNaN (v )||_db .IsInf (v ,0){_fbd ._ecc .TAttr =_ca .ST_CellTypeE ;_fbd ._ecc .V =_b .String ("\u0023\u004e\u0055M\u0021");return ;};_fbd ._ecc .TAttr =_ca .ST_CellTypeN ;_fbd ._ecc .V =_b .String (_fb .FormatFloat (v ,'f',-1,64));};

// AddSheet adds a new sheet to a workbook.
func (_ggbe *Workbook )AddSheet ()Sheet {_aeeg :=_ca .NewCT_Sheet ();_aeeg .SheetIdAttr =1;for _ ,_dbae :=range _ggbe ._gdeg .Sheets .Sheet {if _aeeg .SheetIdAttr <=_dbae .SheetIdAttr {_aeeg .SheetIdAttr =_dbae .SheetIdAttr +1;};};_ggbe ._gdeg .Sheets .Sheet =append (_ggbe ._gdeg .Sheets .Sheet ,_aeeg );_aeeg .NameAttr =_e .Sprintf ("\u0053\u0068\u0065\u0065\u0074\u0020\u0025\u0064",_aeeg .SheetIdAttr );_bgcf :=_ca .NewWorksheet ();_bgcf .Dimension =_ca .NewCT_SheetDimension ();_bgcf .Dimension .RefAttr ="\u0041\u0031";_ggbe ._bbced =append (_ggbe ._bbced ,_bgcf );_cbbe :=_da .NewRelationships ();_ggbe ._dcdg =append (_ggbe ._dcdg ,_cbbe );_bgcf .SheetData =_ca .NewCT_SheetData ();_ggbe ._dfgf =append (_ggbe ._dfgf ,nil );_aagff :=_b .DocTypeSpreadsheet ;_egecc :=_ggbe ._eeead .AddAutoRelationship (_aagff ,_b .OfficeDocumentType ,len (_ggbe ._gdeg .Sheets .Sheet ),_b .WorksheetType );_aeeg .IdAttr =_egecc .ID ();_ggbe .ContentTypes .AddOverride (_b .AbsoluteFilename (_aagff ,_b .WorksheetContentType ,len (_ggbe ._gdeg .Sheets .Sheet )),_b .WorksheetContentType );return Sheet {_ggbe ,_aeeg ,_bgcf };};

// SetMaxLength sets the maximum bar length in percent.
func (_eaf DataBarScale )SetMaxLength (l uint32 ){_eaf ._bdcd .MaxLengthAttr =_b .Uint32 (l )};

// RichTextRun is a segment of text within a cell that is directly formatted.
type RichTextRun struct{_dfc *_ca .CT_RElt };

// SetTime sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel).
func (_fbf Cell )SetTime (d _fd .Time ){_fbf .clearValue ();d =_egd (d );_egg :=_fbf ._fac .Epoch ();if d .Before (_egg ){_cd .Log .Debug ("t\u0069\u006d\u0065\u0073\u0020\u0062e\u0066\u006f\u0072\u0065\u0020\u00319\u0030\u0030\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064");return ;};_fcb :=d .Sub (_egg );_gbc :=new (_ac .Float );_dca :=new (_ac .Float );_dca .SetPrec (128);_dca .SetUint64 (uint64 (_fcb ));_dfda :=new (_ac .Float );_dfda .SetUint64 (24*60*60*1e9);_gbc .Quo (_dca ,_dfda );_fbf ._ecc .V =_b .String (_gbc .Text ('g',20));};func (_cefed Sheet )validateSheetNames ()error {if len (_cefed .Name ())> 31{return _e .Errorf ("\u0073\u0068\u0065\u0065\u0074 \u006e\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027\u0020\u0068\u0061\u0073 \u0025\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u002c\u0020\u006d\u0061\u0078\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0069\u0073\u0020\u00331",_cefed .Name (),len (_cefed .Name ()));};return nil ;};

// SetFont sets the font name for a rich text run.
func (_fdbf RichTextRun )SetFont (s string ){_fdbf .ensureRpr ();_fdbf ._dfc .RPr .RFont =_ca .NewCT_FontName ();_fdbf ._dfc .RPr .RFont .ValAttr =s ;};

// X returns the inner wrapped XML type.
func (_gbfc DataBarScale )X ()*_ca .CT_DataBar {return _gbfc ._bdcd };

// GetFilename returns the filename of the context's workbook.
func (_bad *evalContext )GetFilename ()string {return _bad ._bcdc ._fdfd .GetFilename ()};

// SetRowOffset sets a column offset in absolute distance.
func (_bfce CellMarker )SetRowOffset (m _ec .Distance ){_bfce ._gebb .RowOff .ST_CoordinateUnqualified =_b .Int64 (int64 (m /_ec .EMU ));};func (_bde *evalContext )NamedRange (ref string )_ab .Reference {for _ ,_dgge :=range _bde ._bcdc ._fdfd .DefinedNames (){if _dgge .Name ()==ref {return _ab .MakeRangeReference (_dgge .Content ());};};for _ ,_cdcc :=range _bde ._bcdc ._fdfd .Tables (){if _cdcc .Name ()==ref {return _ab .MakeRangeReference (_e .Sprintf ("\u0025\u0073\u0021%\u0073",_bde ._bcdc .Name (),_cdcc .Reference ()));};};return _ab .ReferenceInvalid ;};

// Type returns the type of anchor
func (_eadc TwoCellAnchor )Type ()AnchorType {return AnchorTypeTwoCell };func NewPatternFill (fills *_ca .CT_Fills )PatternFill {_dbce :=_ca .NewCT_Fill ();_dbce .PatternFill =_ca .NewCT_PatternFill ();return PatternFill {_dbce .PatternFill ,_dbce };};

// SetOperator sets the operator for the rule.
func (_cdcd ConditionalFormattingRule )SetOperator (t _ca .ST_ConditionalFormattingOperator ){_cdcd ._afcg .OperatorAttr =t ;};

// SortOrder is a column sort order.
//go:generate stringer -type=SortOrder
type SortOrder byte ;

// InitialView returns the first defined sheet view. If there are no views, one
// is created and returned.
func (_bacf *Sheet )InitialView ()SheetView {if _bacf ._cffg .SheetViews ==nil ||len (_bacf ._cffg .SheetViews .SheetView )==0{return _bacf .AddView ();};return SheetView {_bacf ._cffg .SheetViews .SheetView [0]};};func (_gcba StandardFormat )String ()string {switch {case 0<=_gcba &&_gcba <=4:return _bdbd [_ebffb [_gcba ]:_ebffb [_gcba +1]];case 9<=_gcba &&_gcba <=22:_gcba -=9;return _bcfc [_cfag [_gcba ]:_cfag [_gcba +1]];case 37<=_gcba &&_gcba <=40:_gcba -=37;return _ffbd [_gafd [_gcba ]:_gafd [_gcba +1]];case 45<=_gcba &&_gcba <=49:_gcba -=45;return _cffgc [_bfeg [_gcba ]:_bfeg [_gcba +1]];default:return _e .Sprintf ("\u0053t\u0061n\u0064\u0061\u0072\u0064\u0046o\u0072\u006da\u0074\u0028\u0025\u0064\u0029",_gcba );};};

// Comparer is used to compare rows based off a column and cells based off of
// their value.
type Comparer struct{Order SortOrder ;};func (_ddbgg Font )SetName (name string ){_ddbgg ._ccgc .Name =[]*_ca .CT_FontName {{ValAttr :name }}};

// IsNumber returns true if the cell is a number type cell.
func (_bce Cell )IsNumber ()bool {switch _bce ._ecc .TAttr {case _ca .ST_CellTypeN :return true ;case _ca .ST_CellTypeS ,_ca .ST_CellTypeB :return false ;};return _bce ._ecc .V !=nil &&_ga .IsNumber (*_bce ._ecc .V );};

// ClearProtection clears all workbook protections.
func (_bbdg *Workbook )ClearProtection (){_bbdg ._gdeg .WorkbookProtection =nil };func (_eafb Font )SetBold (b bool ){if b {_eafb ._ccgc .B =[]*_ca .CT_BooleanProperty {{}};}else {_eafb ._ccgc .B =nil ;};};

// SetColOffset sets a column offset in absolute distance.
func (_ed CellMarker )SetColOffset (m _ec .Distance ){_ed ._gebb .ColOff .ST_CoordinateUnqualified =_b .Int64 (int64 (m /_ec .EMU ));};

// AddImage adds an image to the workbook package, returning a reference that
// can be used to add the image to a drawing.
func (_fcfgb *Workbook )AddImage (i _da .Image )(_da .ImageRef ,error ){_geebf :=_da .MakeImageRef (i ,&_fcfgb .DocBase ,_fcfgb ._eeead );if i .Data ==nil &&i .Path ==""{return _geebf ,_dec .New ("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068");};if i .Format ==""{return _geebf ,_dec .New ("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074");};if i .Size .X ==0||i .Size .Y ==0{return _geebf ,_dec .New ("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065");};if i .Path !=""{_fgcf :=_cdc .Add (i .Path );if _fgcf !=nil {return _geebf ,_fgcf ;};};_fcfgb .Images =append (_fcfgb .Images ,_geebf );return _geebf ,nil ;};

// SetStyle applies a style to the cell.  This style is referenced in the
// generated XML via CellStyle.Index().
func (_bb Cell )SetStyle (cs CellStyle ){_bb .SetStyleIndex (cs .Index ())};

// SetItalic causes the text to be displayed in italic.
func (_aace RichTextRun )SetItalic (b bool ){_aace .ensureRpr ();_aace ._dfc .RPr .I =_ca .NewCT_BooleanProperty ();_aace ._dfc .RPr .I .ValAttr =_b .Bool (b );};

// SetFormulaRaw sets the cell type to formula, and the raw formula to the given string
func (_ccb Cell )SetFormulaRaw (s string ){_abe :=_ab .ParseString (s );if _abe ==nil {return ;};_ccb .clearValue ();_ccb ._ecc .TAttr =_ca .ST_CellTypeStr ;_ccb ._ecc .F =_ca .NewCT_CellFormula ();_ccb ._ecc .F .Content =s ;};

// GetEpoch returns a workbook's time epoch.
func (_gege *evalContext )GetEpoch ()_fd .Time {return _gege ._bcdc ._fdfd .Epoch ()};

// DataValidation controls cell validation
type DataValidation struct{_ege *_ca .CT_DataValidation };

// X returns the inner wrapped XML type.
func (_ccba Sheet )X ()*_ca .Worksheet {return _ccba ._cffg };

// GetFormattedValue returns the formatted cell value as it would appear in
// Excel. This involves determining the format string to apply, parsing it, and
// then formatting the value according to the format string.  This should only
// be used if you care about replicating what Excel would show, otherwise
// GetValueAsNumber()/GetValueAsTime
func (_bca Cell )GetFormattedValue ()string {_caf :=_bca .getFormat ();switch _bca ._ecc .TAttr {case _ca .ST_CellTypeB :_fcg ,_ :=_bca .GetValueAsBool ();if _fcg {return "\u0054\u0052\u0055\u0045";};return "\u0046\u0041\u004cS\u0045";case _ca .ST_CellTypeN :_acfb ,_ :=_bca .GetValueAsNumber ();return _ga .Number (_acfb ,_caf );case _ca .ST_CellTypeE :if _bca ._ecc .V !=nil {return *_bca ._ecc .V ;};return "";case _ca .ST_CellTypeS ,_ca .ST_CellTypeInlineStr :return _ga .String (_bca .GetString (),_caf );case _ca .ST_CellTypeStr :_gfd :=_bca .GetString ();if _ga .IsNumber (_gfd ){_bcc ,_ :=_fb .ParseFloat (_gfd ,64);return _ga .Number (_bcc ,_caf );};return _ga .String (_gfd ,_caf );case _ca .ST_CellTypeUnset :fallthrough;default:_cdag ,_ :=_bca .GetRawValue ();if len (_cdag )==0{return "";};_aeeb ,_abc :=_bca .GetValueAsNumber ();if _abc ==nil {return _ga .Number (_aeeb ,_caf );};return _ga .String (_cdag ,_caf );};};

// X returns the inner wrapped XML type.
func (_bafg MergedCell )X ()*_ca .CT_MergeCell {return _bafg ._fgf };

// SetStyleIndex directly sets a style index to the cell.  This should only be
// called with an index retrieved from CellStyle.Index()
func (_bag Cell )SetStyleIndex (idx uint32 ){_bag ._ecc .SAttr =_b .Uint32 (idx )};

// LessCells returns true if the lhs value is less than the rhs value. If the
// cells contain numeric values, their value interpreted as a floating point is
// compared. Otherwise their string contents are compared.
func (_eefa Comparer )LessCells (lhs ,rhs Cell )bool {if _eefa .Order ==SortOrderDescending {lhs ,rhs =rhs ,lhs ;};if lhs .X ()==nil {if rhs .X ()==nil {return false ;};return true ;};if rhs .X ()==nil {return false ;};_ffde ,_dcad :=lhs .getRawSortValue ();_dga ,_eac :=rhs .getRawSortValue ();switch {case _dcad &&_eac :_bba ,_ :=_fb .ParseFloat (_ffde ,64);_ddbg ,_ :=_fb .ParseFloat (_dga ,64);return _bba < _ddbg ;case _dcad :return true ;case _eac :return false ;};_ffde =lhs .GetFormattedValue ();_dga =rhs .GetFormattedValue ();return _ffde < _dga ;};

// GetValueAsBool retrieves the cell's value as a boolean
func (_fdec Cell )GetValueAsBool ()(bool ,error ){if _fdec ._ecc .TAttr !=_ca .ST_CellTypeB {return false ,_dec .New ("\u0063e\u006c\u006c\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006ff\u0020\u0062\u006f\u006f\u006c\u0020\u0074\u0079\u0070\u0065");};if _fdec ._ecc .V ==nil {return false ,_dec .New ("\u0063\u0065\u006c\u006c\u0020\u0068\u0061\u0073\u0020\u006e\u006f\u0020v\u0061\u006c\u0075\u0065");};return _fb .ParseBool (*_fdec ._ecc .V );};

// SetLocked sets cell locked or not.
func (_ceg *evalContext )SetLocked (cellRef string ,locked bool ){_ceg ._bcdc .Cell (cellRef ).setLocked (locked );};var _ebdg =[...]uint8 {0,18,37};

// InitializeDefaults initializes a border to its defaulte empty values.
func (_cea Border )InitializeDefaults (){_cea ._fcf .Left =_ca .NewCT_BorderPr ();_cea ._fcf .Bottom =_ca .NewCT_BorderPr ();_cea ._fcf .Right =_ca .NewCT_BorderPr ();_cea ._fcf .Top =_ca .NewCT_BorderPr ();_cea ._fcf .Diagonal =_ca .NewCT_BorderPr ();};

// X returns the inner wrapped XML type.
func (_acfa *Workbook )X ()*_ca .Workbook {return _acfa ._gdeg };

// Name returns the name of the table
func (_gaed Table )Name ()string {if _gaed ._gcag .NameAttr !=nil {return *_gaed ._gcag .NameAttr ;};return "";};

// Workbook is the top level container item for a set of spreadsheets.
type Workbook struct{_da .DocBase ;_gdeg *_ca .Workbook ;StyleSheet StyleSheet ;SharedStrings SharedStrings ;_dfgf []*_ca .Comments ;_bbced []*_ca .Worksheet ;_dcdg []_da .Relationships ;_eeead _da .Relationships ;_gfcg []*_bgc .Theme ;_cffgf []*_acd .WsDr ;_aeec []_da .Relationships ;_dbff []*_bg .Container ;_dgeg []*_fc .ChartSpace ;_beee []*_ca .Table ;_eeaf string ;_gdae map[string ]string ;_gcbc map[string ]*_fc .ChartSpace ;_fcgbe string ;};func (_ebbb Cell )getFormat ()string {if _ebbb ._ecc .SAttr ==nil {return "\u0047e\u006e\u0065\u0072\u0061\u006c";};_dfb :=*_ebbb ._ecc .SAttr ;_ccg :=_ebbb ._fac .StyleSheet .GetCellStyle (_dfb );_ge :=_ebbb ._fac .StyleSheet .GetNumberFormat (_ccg .NumberFormat ());return _ge .GetFormat ();};

// Themes returns the array of workbook dml.Theme.
func (_gecc *Workbook )Themes ()[]*_bgc .Theme {return _gecc ._gfcg };func (_add Cell )getLocked ()bool {if _add ._ecc .SAttr ==nil {return false ;};_cfg :=*_add ._ecc .SAttr ;_bga :=_add ._fac .StyleSheet .GetCellStyle (_cfg );return *_bga ._fbg .Protection .LockedAttr ;};

// Name returns the sheet name
func (_ebd Sheet )Name ()string {return _ebd ._aacc .NameAttr };type ConditionalFormattingRule struct{_afcg *_ca .CT_CfRule };

// BottomRight returns the CellMaker for the bottom right corner of the anchor.
func (_dcfb TwoCellAnchor )BottomRight ()CellMarker {return CellMarker {_dcfb ._ffgg .To }};

// X returns the inner wrapped XML type.
func (_cce Comment )X ()*_ca .CT_Comment {return _cce ._daea };

// AddDifferentialStyle adds a new empty differential cell style to the stylesheet.
func (_gdbe StyleSheet )AddDifferentialStyle ()DifferentialStyle {if _gdbe ._gace .Dxfs ==nil {_gdbe ._gace .Dxfs =_ca .NewCT_Dxfs ();};_bcca :=_ca .NewCT_Dxf ();_gdbe ._gace .Dxfs .Dxf =append (_gdbe ._gace .Dxfs .Dxf ,_bcca );_gdbe ._gace .Dxfs .CountAttr =_b .Uint32 (uint32 (len (_gdbe ._gace .Dxfs .Dxf )));return DifferentialStyle {_bcca ,_gdbe ._efga ,_gdbe ._gace .Dxfs };};

// Cell creates or returns a cell given a cell reference of the form 'A10'
func (_ebbbg *Sheet )Cell (cellRef string )Cell {_adbbe ,_adee :=_fda .ParseCellReference (cellRef );if _adee !=nil {_cd .Log .Debug ("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063e\u006cl\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u003a\u0020\u0025\u0073",_adee );return _ebbbg .AddRow ().AddCell ();};return _ebbbg .Row (_adbbe .RowIdx ).Cell (_adbbe .Column );};

// Comments is the container for comments for a single sheet.
type Comments struct{_gebe *Workbook ;_cfc *_ca .Comments ;};

// AddImage adds an image with a paricular anchor type, returning an anchor to
// allow adusting the image size/position.
func (_afdc Drawing )AddImage (img _da .ImageRef ,at AnchorType )Anchor {_gecb :=0;for _aadf ,_ffda :=range _afdc ._fge .Images {if _ffda ==img {_gecb =_aadf +1;break ;};};var _ggad string ;for _gbba ,_fdbd :=range _afdc ._fge ._cffgf {if _fdbd ==_afdc ._bda {_eafe :=_e .Sprintf ("\u002e\u002e\u002f\u006ded\u0069\u0061\u002f\u0069\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073",_gecb ,img .Format ());_bacd :=_afdc ._fge ._aeec [_gbba ].AddRelationship (_eafe ,_b .ImageType );_ggad =_bacd .ID ();break ;};};var _aed Anchor ;var _gda *_acd .CT_Picture ;switch at {case AnchorTypeAbsolute :_eefb :=_cfec ();_afdc ._bda .EG_Anchor =append (_afdc ._bda .EG_Anchor ,&_acd .EG_Anchor {AbsoluteAnchor :_eefb });_eefb .Choice =&_acd .EG_ObjectChoicesChoice {};_eefb .Choice .Pic =_acd .NewCT_Picture ();_eefb .Pos .XAttr .ST_CoordinateUnqualified =_b .Int64 (0);_eefb .Pos .YAttr .ST_CoordinateUnqualified =_b .Int64 (0);_gda =_eefb .Choice .Pic ;_aed =AbsoluteAnchor {_eefb };case AnchorTypeOneCell :_bef :=_eff ();_afdc ._bda .EG_Anchor =append (_afdc ._bda .EG_Anchor ,&_acd .EG_Anchor {OneCellAnchor :_bef });_bef .Choice =&_acd .EG_ObjectChoicesChoice {};_bef .Choice .Pic =_acd .NewCT_Picture ();_gda =_bef .Choice .Pic ;_aed =OneCellAnchor {_bef };case AnchorTypeTwoCell :_gdd :=_cbdf ();_afdc ._bda .EG_Anchor =append (_afdc ._bda .EG_Anchor ,&_acd .EG_Anchor {TwoCellAnchor :_gdd });_gdd .Choice =&_acd .EG_ObjectChoicesChoice {};_gdd .Choice .Pic =_acd .NewCT_Picture ();_gda =_gdd .Choice .Pic ;_aed =TwoCellAnchor {_gdd };};_gda .NvPicPr .CNvPr .IdAttr =uint32 (len (_afdc ._bda .EG_Anchor ));_gda .NvPicPr .CNvPr .NameAttr ="\u0049\u006d\u0061g\u0065";_gda .BlipFill .Blip =_bgc .NewCT_Blip ();_gda .BlipFill .Blip .EmbedAttr =_b .String (_ggad );_gda .BlipFill .Stretch =_bgc .NewCT_StretchInfoProperties ();_gda .SpPr =_bgc .NewCT_ShapeProperties ();_gda .SpPr .Xfrm =_bgc .NewCT_Transform2D ();_gda .SpPr .Xfrm .Off =_bgc .NewCT_Point2D ();_gda .SpPr .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_b .Int64 (0);_gda .SpPr .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_b .Int64 (0);_gda .SpPr .Xfrm .Ext =_bgc .NewCT_PositiveSize2D ();_gda .SpPr .Xfrm .Ext .CxAttr =int64 (float64 (img .Size ().X *_ec .Pixel72 )/_ec .EMU );_gda .SpPr .Xfrm .Ext .CyAttr =int64 (float64 (img .Size ().Y *_ec .Pixel72 )/_ec .EMU );_gda .SpPr .PrstGeom =_bgc .NewCT_PresetGeometry2D ();_gda .SpPr .PrstGeom .PrstAttr =_bgc .ST_ShapeTypeRect ;_gda .SpPr .Ln =_bgc .NewCT_LineProperties ();_gda .SpPr .Ln .NoFill =_bgc .NewCT_NoFillProperties ();return _aed ;};

// HasNumberFormat returns true if the cell style has a number format applied.
func (_afe CellStyle )HasNumberFormat ()bool {return _afe ._fbg .NumFmtIdAttr !=nil &&_afe ._fbg .ApplyNumberFormatAttr !=nil &&*_afe ._fbg .ApplyNumberFormatAttr ;};

// GetFormat returns a cell data format.
func (_adde *evalContext )GetFormat (cellRef string )string {return _adde ._bcdc .Cell (cellRef ).getFormat ();};

// DVCompareOp is a comparison operator for a data validation rule.
type DVCompareOp byte ;

// SetHorizontalAlignment sets the horizontal alignment of a cell style.
func (_ccd CellStyle )SetHorizontalAlignment (a _ca .ST_HorizontalAlignment ){if _ccd ._fbg .Alignment ==nil {_ccd ._fbg .Alignment =_ca .NewCT_CellAlignment ();};_ccd ._fbg .Alignment .HorizontalAttr =a ;_ccd ._fbg .ApplyAlignmentAttr =_b .Bool (true );};func (_gf Cell )clearValue (){_gf ._ecc .F =nil ;_gf ._ecc .Is =nil ;_gf ._ecc .V =nil ;_gf ._ecc .TAttr =_ca .ST_CellTypeUnset ;};

// SetColor sets the text color.
func (_egfa RichTextRun )SetColor (c _fde .Color ){_egfa .ensureRpr ();_egfa ._dfc .RPr .Color =_ca .NewCT_Color ();_aaad :="\u0066\u0066"+*c .AsRGBString ();_egfa ._dfc .RPr .Color .RgbAttr =&_aaad ;};

// AddHyperlink adds a hyperlink to a sheet. Adding the hyperlink to the sheet
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (_dgae *Sheet )AddHyperlink (url string )_da .Hyperlink {for _bgfc ,_dfgg :=range _dgae ._fdfd ._bbced {if _dfgg ==_dgae ._cffg {return _dgae ._fdfd ._dcdg [_bgfc ].AddHyperlink (url );};};return _da .Hyperlink {};};

// ID returns the number format ID.  This is not an index as there are some
// predefined number formats which can be used in cell styles and don't need a
// corresponding NumberFormat.
func (_ged NumberFormat )ID ()uint32 {return _ged ._bgfa .NumFmtIdAttr };const (DVCompareOpEqual =DVCompareOp (_ca .ST_DataValidationOperatorEqual );DVCompareOpBetween =DVCompareOp (_ca .ST_DataValidationOperatorBetween );DVCompareOpNotBetween =DVCompareOp (_ca .ST_DataValidationOperatorNotBetween );DVCompareOpNotEqual =DVCompareOp (_ca .ST_DataValidationOperatorNotEqual );DVCompareOpGreater =DVCompareOp (_ca .ST_DataValidationOperatorGreaterThan );DVCompareOpGreaterEqual =DVCompareOp (_ca .ST_DataValidationOperatorGreaterThanOrEqual );DVCompareOpLess =DVCompareOp (_ca .ST_DataValidationOperatorLessThan );DVCompareOpLessEqual =DVCompareOp (_ca .ST_DataValidationOperatorLessThanOrEqual ););func (_ebcf *evalContext )Sheet (name string )_ab .Context {for _ ,_fab :=range _ebcf ._bcdc ._fdfd .Sheets (){if _fab .Name ()==name {return _fab .FormulaContext ();};};return _ab .InvalidReferenceContext ;};type Fills struct{_bcfe *_ca .CT_Fills };

// CellReference returns the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (_dfbe Comment )CellReference ()string {return _dfbe ._daea .RefAttr };

// SetHeightCells is a no-op.
func (_fgd OneCellAnchor )SetHeightCells (int32 ){};

// ColorScale colors a cell background based off of the cell value.
type ColorScale struct{_dfdc *_ca .CT_ColorScale };

// RecalculateFormulas re-computes any computed formula values that are stored
// in the sheet. As unioffice formula support is still new and not all functins are
// supported,  if formula execution fails either due to a parse error or missing
// function, or erorr in the result (even if expected) the cached value will be
// left empty allowing Excel to recompute it on load.
func (_cgaa *Sheet )RecalculateFormulas (){_ddfc :=_ab .NewEvaluator ();_degg :=_cgaa .FormulaContext ();for _ ,_afgf :=range _cgaa .Rows (){for _ ,_gbbg :=range _afgf .Cells (){if _gbbg .X ().F !=nil {_ccfgb :=_gbbg .X ().F .Content ;if _gbbg .X ().F .TAttr ==_ca .ST_CellFormulaTypeShared &&len (_ccfgb )==0{continue ;};_bfbc :=_ddfc .Eval (_degg ,_ccfgb ).AsString ();if _bfbc .Type ==_ab .ResultTypeError {_cd .Log .Debug ("\u0065\u0072\u0072o\u0072\u0020\u0065\u0076a\u0075\u006c\u0061\u0074\u0069\u006e\u0067 \u0066\u006f\u0072\u006d\u0075\u006c\u0061\u0020\u0025\u0073\u003a\u0020\u0025\u0073",_ccfgb ,_bfbc .ErrorMessage );_gbbg .X ().V =nil ;}else {if _bfbc .Type ==_ab .ResultTypeNumber {_gbbg .X ().TAttr =_ca .ST_CellTypeN ;}else {_gbbg .X ().TAttr =_ca .ST_CellTypeInlineStr ;};_gbbg .X ().V =_b .String (_bfbc .Value ());if _gbbg .X ().F .TAttr ==_ca .ST_CellFormulaTypeArray {if _bfbc .Type ==_ab .ResultTypeArray {_cgaa .setArray (_gbbg .Reference (),_bfbc );}else if _bfbc .Type ==_ab .ResultTypeList {_cgaa .setList (_gbbg .Reference (),_bfbc );};}else if _gbbg .X ().F .TAttr ==_ca .ST_CellFormulaTypeShared &&_gbbg .X ().F .RefAttr !=nil {_fcbd ,_dfbf ,_defcc :=_fda .ParseRangeReference (*_gbbg .X ().F .RefAttr );if _defcc !=nil {_cd .Log .Debug ("\u0065\u0072r\u006f\u0072\u0020\u0069n\u0020\u0073h\u0061\u0072\u0065\u0064\u0020\u0066\u006f\u0072m\u0075\u006c\u0061\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u003a\u0020\u0025\u0073",_defcc );continue ;};_cgaa .setShared (_gbbg .Reference (),_fcbd ,_dfbf ,_ccfgb );};};};};};};

// X returns the inner wrapped XML type.
func (_fbfe DifferentialStyle )X ()*_ca .CT_Dxf {return _fbfe ._ccad };

// AddNamedCell adds a new named cell to a row and returns it. You should
// normally prefer Cell() as it will return the existing cell if the cell
// already exists, while AddNamedCell will duplicate the cell creating an
// invaild spreadsheet.
func (_caee Row )AddNamedCell (col string )Cell {_gbbef :=_ca .NewCT_Cell ();_gbbef .RAttr =_b .Stringf ("\u0025\u0073\u0025\u0064",col ,_caee .RowNumber ());_gfdc :=-1;_abdc :=_fda .ColumnToIndex (col );for _dbaf ,_ddgd :=range _caee ._gbgd .C {_baff ,_acc :=_fda .ParseCellReference (*_ddgd .RAttr );if _acc !=nil {return Cell {};};if _abdc < _baff .ColumnIdx {_gfdc =_dbaf ;break ;};};if _gfdc ==-1{_caee ._gbgd .C =append (_caee ._gbgd .C ,_gbbef );}else {_caee ._gbgd .C =append (_caee ._gbgd .C [:_gfdc ],append ([]*_ca .CT_Cell {_gbbef },_caee ._gbgd .C [_gfdc :]...)...);};return Cell {_caee ._bdaa ,_caee ._cgee ,_caee ._gbgd ,_gbbef };};

// MoveTo repositions the anchor without changing the objects size.
func (_fdefd TwoCellAnchor )MoveTo (col ,row int32 ){_efgfg :=_fdefd .TopLeft ();_fbbb :=_fdefd .BottomRight ();_gdg :=_fbbb .Col ()-_efgfg .Col ();_dbafc :=_fbbb .Row ()-_efgfg .Row ();_efgfg .SetCol (col );_efgfg .SetRow (row );_fbbb .SetCol (col +_gdg );_fbbb .SetRow (row +_dbafc );};

// CellText is used for keeping text with references to a cell where it is located.
type CellText struct{Text string ;Cell Cell ;};

// X returns the inner wrapped XML type.
func (_bgfg IconScale )X ()*_ca .CT_IconSet {return _bgfg ._dead };

// SetWidth sets the width of the anchored object.
func (_be AbsoluteAnchor )SetWidth (w _ec .Distance ){_be ._ag .Ext .CxAttr =int64 (w /_ec .EMU )};func (_fgcg Font )SetColor (c _fde .Color ){_gced :=_ca .NewCT_Color ();_egb :="\u0066\u0066"+*c .AsRGBString ();_gced .RgbAttr =&_egb ;_fgcg ._ccgc .Color =[]*_ca .CT_Color {_gced };};

// Protection controls the protection on an individual sheet.
func (_gafg *Sheet )Protection ()SheetProtection {if _gafg ._cffg .SheetProtection ==nil {_gafg ._cffg .SheetProtection =_ca .NewCT_SheetProtection ();};return SheetProtection {_gafg ._cffg .SheetProtection };};

// SetRowOffset sets the row offset of the two cell anchor
func (_edbb TwoCellAnchor )SetRowOffset (m _ec .Distance ){_dega :=m -_edbb .TopLeft ().RowOffset ();_edbb .TopLeft ().SetRowOffset (m );_edbb .BottomRight ().SetRowOffset (_edbb .BottomRight ().RowOffset ()+_dega );};var ErrorNotFound =_dec .New ("\u006eo\u0074\u0020\u0066\u006f\u0075\u006ed");

// AddComment adds a new comment and returns a RichText which will contain the
// styled comment text.
func (_dge Comments )AddComment (cellRef string ,author string )RichText {_gegc :=_ca .NewCT_Comment ();_dge ._cfc .CommentList .Comment =append (_dge ._cfc .CommentList .Comment ,_gegc );_gegc .RefAttr =cellRef ;_gegc .AuthorIdAttr =_dge .getOrCreateAuthor (author );_gegc .Text =_ca .NewCT_Rst ();return RichText {_gegc .Text };};

// Reference returns the table reference (the cells within the table)
func (_ceffa Table )Reference ()string {return _ceffa ._gcag .RefAttr };

// Row will return a row with a given row number, creating a new row if
// necessary.
func (_ccabe *Sheet )Row (rowNum uint32 )Row {for _ ,_gbgf :=range _ccabe ._cffg .SheetData .Row {if _gbgf .RAttr !=nil &&*_gbgf .RAttr ==rowNum {return Row {_ccabe ._fdfd ,_ccabe ,_gbgf };};};return _ccabe .AddNumberedRow (rowNum );};

// SetZoom controls the zoom level of the sheet and is measured in percent. The
// default value is 100.
func (_dcacbd SheetView )SetZoom (pct uint32 ){_dcacbd ._egcb .ZoomScaleAttr =&pct };

// StyleSheet is a document style sheet.
type StyleSheet struct{_efga *Workbook ;_gace *_ca .StyleSheet ;};func (_agg *Sheet )setShared (_bedg string ,_fffbda ,_aage _fda .CellReference ,_daff string ){_ddgdf :=_agg .FormulaContext ();_geag :=_ab .NewEvaluator ();for _gbage :=_fffbda .RowIdx ;_gbage <=_aage .RowIdx ;_gbage ++{for _addbb :=_fffbda .ColumnIdx ;_addbb <=_aage .ColumnIdx ;_addbb ++{_dgfb :=_gbage -_fffbda .RowIdx ;_gcaae :=_addbb -_fffbda .ColumnIdx ;_ddgdf .SetOffset (_gcaae ,_dgfb );_ddcd :=_geag .Eval (_ddgdf ,_daff );_agbf :=_e .Sprintf ("\u0025\u0073\u0025\u0064",_fda .IndexToColumn (_addbb ),_gbage );_aaff :=_agg .Cell (_agbf );if _ddcd .Type ==_ab .ResultTypeNumber {_aaff .X ().TAttr =_ca .ST_CellTypeN ;}else {_aaff .X ().TAttr =_ca .ST_CellTypeInlineStr ;};_aaff .X ().V =_b .String (_ddcd .Value ());};};_ =_geag ;_ =_ddgdf ;};

// ClearProtection removes any protections applied to teh sheet.
func (_aggd *Sheet )ClearProtection (){_aggd ._cffg .SheetProtection =nil };func (_afebd Fill )Index ()uint32 {if _afebd ._dcfe ==nil {return 0;};for _ggd ,_ecbb :=range _afebd ._dcfe .Fill {if _afebd ._aae ==_ecbb {return uint32 (_ggd );};};return 0;};

// Fonts returns the list of fonts defined in the stylesheet.
func (_cbge StyleSheet )Fonts ()[]Font {_fgbe :=[]Font {};for _ ,_bggdb :=range _cbge ._gace .Fonts .Font {_fgbe =append (_fgbe ,Font {_bggdb ,_cbge ._gace });};return _fgbe ;};

// SetFormulaArray sets the cell type to formula array, and the raw formula to
// the given string. This is equivlent to entering a formula and pressing
// Ctrl+Shift+Enter in Excel.
func (_dag Cell )SetFormulaArray (s string ){_gd :=_ab .ParseString (s );if _gd ==nil {return ;};_dag .clearValue ();_dag ._ecc .TAttr =_ca .ST_CellTypeStr ;_dag ._ecc .F =_ca .NewCT_CellFormula ();_dag ._ecc .F .TAttr =_ca .ST_CellFormulaTypeArray ;_dag ._ecc .F .Content =s ;};

// AddHyperlink creates and sets a hyperlink on a cell.
func (_cde Cell )AddHyperlink (url string ){for _fff ,_bac :=range _cde ._fac ._bbced {if _bac ==_cde ._cfd ._cffg {_cde .SetHyperlink (_cde ._fac ._dcdg [_fff ].AddHyperlink (url ));return ;};};};

// RemoveDefinedName removes an existing defined name.
func (_fdde *Workbook )RemoveDefinedName (dn DefinedName )error {if dn .X ()==nil {return _dec .New ("\u0061\u0074\u0074\u0065\u006d\u0070t\u0020\u0074\u006f\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u006e\u0069l\u0020\u0044\u0065\u0066\u0069\u006e\u0065d\u004e\u0061\u006d\u0065");};for _gcefe ,_gebc :=range _fdde ._gdeg .DefinedNames .DefinedName {if _gebc ==dn .X (){copy (_fdde ._gdeg .DefinedNames .DefinedName [_gcefe :],_fdde ._gdeg .DefinedNames .DefinedName [_gcefe +1:]);_fdde ._gdeg .DefinedNames .DefinedName [len (_fdde ._gdeg .DefinedNames .DefinedName )-1]=nil ;_fdde ._gdeg .DefinedNames .DefinedName =_fdde ._gdeg .DefinedNames .DefinedName [:len (_fdde ._gdeg .DefinedNames .DefinedName )-1];return nil ;};};return _dec .New ("\u0064\u0065\u0066\u0069ne\u0064\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");};func (_baec *Sheet )removeColumnFromMergedCells (_ccbc uint32 )error {if _baec ._cffg .MergeCells ==nil ||_baec ._cffg .MergeCells .MergeCell ==nil {return nil ;};_gaca :=[]*_ca .CT_MergeCell {};for _ ,_cecaa :=range _baec .MergedCells (){_ggge :=_gcec (_cecaa .Reference (),_ccbc ,true );if _ggge !=""{_cecaa .SetReference (_ggge );_gaca =append (_gaca ,_cecaa .X ());};};_baec ._cffg .MergeCells .MergeCell =_gaca ;return nil ;};

// ClearAutoFilter removes the autofilters from the sheet.
func (_gfeb *Sheet )ClearAutoFilter (){_gfeb ._cffg .AutoFilter =nil ;_dgdg :="\u0027"+_gfeb .Name ()+"\u0027\u0021";for _ ,_ecbg :=range _gfeb ._fdfd .DefinedNames (){if _ecbg .Name ()==_gdaf {if _fa .HasPrefix (_ecbg .Content (),_dgdg ){_gfeb ._fdfd .RemoveDefinedName (_ecbg );break ;};};};};

// SetYSplit sets the row split point
func (_ggce SheetView )SetYSplit (v float64 ){_ggce .ensurePane ();_ggce ._egcb .Pane .YSplitAttr =_b .Float64 (v );};type Fill struct{_aae *_ca .CT_Fill ;_dcfe *_ca .CT_Fills ;};func (_dfd Border )SetTop (style _ca .ST_BorderStyle ,c _fde .Color ){if _dfd ._fcf .Top ==nil {_dfd ._fcf .Top =_ca .NewCT_BorderPr ();};_dfd ._fcf .Top .Color =_ca .NewCT_Color ();_dfd ._fcf .Top .Color .RgbAttr =c .AsRGBAString ();_dfd ._fcf .Top .StyleAttr =style ;};func (_faga *Sheet )removeColumnFromNamedRanges (_afgfg uint32 )error {for _ ,_cbaa :=range _faga ._fdfd .DefinedNames (){_ffdd :=_cbaa .Name ();_defcb :=_cbaa .Content ();_cbca :=_fa .Split (_defcb ,"\u0021");if len (_cbca )!=2{return _dec .New ("\u0049\u006e\u0063\u006frr\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0064\u0020\u0072\u0061\u006e\u0067e\u003a"+_defcb );};_agce :=_cbca [0];if _faga .Name ()==_agce {_cafg :=_faga ._fdfd .RemoveDefinedName (_cbaa );if _cafg !=nil {return _cafg ;};_gabe :=_gcec (_cbca [1],_afgfg ,true );if _gabe !=""{_fbde :=_agce +"\u0021"+_gabe ;_faga ._fdfd .AddDefinedName (_ffdd ,_fbde );};};};_bgae :=0;if _faga ._cffg .TableParts !=nil &&_faga ._cffg .TableParts .TablePart !=nil {_bgae =len (_faga ._cffg .TableParts .TablePart );};if _bgae !=0{_dbde :=0;for _ ,_fffdg :=range _faga ._fdfd .Sheets (){if _fffdg .Name ()==_faga .Name (){break ;}else {if _fffdg ._cffg .TableParts !=nil &&_fffdg ._cffg .TableParts .TablePart !=nil {_dbde +=len (_fffdg ._cffg .TableParts .TablePart );};};};_gcaf :=_faga ._fdfd ._beee [_dbde :_dbde +_bgae ];for _defa ,_cfdac :=range _gcaf {_fged :=_cfdac ;_fged .RefAttr =_gcec (_fged .RefAttr ,_afgfg ,false );_faga ._fdfd ._beee [_dbde +_defa ]=_fged ;};};return nil ;};func (_eee Font )Index ()uint32 {for _bddg ,_eeg :=range _eee ._fbdb .Fonts .Font {if _eee ._ccgc ==_eeg {return uint32 (_bddg );};};return 0;};

// SetRange sets the range that contains the possible values. This is incompatible with SetValues.
func (_dda DataValidationList )SetRange (cellRange string ){_dda ._dce .Formula1 =_b .String (cellRange );_dda ._dce .Formula2 =_b .String ("\u0030");};

// MoveTo moves the top-left of the anchored object.
func (_dfae OneCellAnchor )MoveTo (col ,row int32 ){_dfae .TopLeft ().SetCol (col );_dfae .TopLeft ().SetRow (row );};

// SetBorder applies a border to a cell style. The border is referenced by its
// index so modifying the border afterward will affect all styles that reference
// it.
func (_efa CellStyle )SetBorder (b Border ){_efa ._fbg .BorderIdAttr =_b .Uint32 (b .Index ());_efa ._fbg .ApplyBorderAttr =_b .Bool (true );};

// SetRowOffset sets the row offset of the top-left anchor.
func (_gecf OneCellAnchor )SetRowOffset (m _ec .Distance ){_gecf .TopLeft ().SetRowOffset (m )};func (_acga *evalContext )SetOffset (col ,row uint32 ){_acga ._geca =col ;_acga ._bbea =row };

// X returns the inner wrapped XML type.
func (_gffc Row )X ()*_ca .CT_Row {return _gffc ._gbgd };

// SetCol set the column of the cell marker.
func (_baf CellMarker )SetCol (col int32 ){_baf ._gebb .Col =col };

// RowNumber returns the row number (1-N), or zero if it is unset.
func (_eded Row )RowNumber ()uint32 {if _eded ._gbgd .RAttr !=nil {return *_eded ._gbgd .RAttr ;};return 0;};

// GetFill gets a Fill from a cell style.
func (_gffb CellStyle )GetFill ()*_ca .CT_Fill {if _beba :=_gffb ._fbg .FillIdAttr ;_beba !=nil {_eab :=_gffb ._dcb .StyleSheet .Fills ().X ().Fill ;if int (*_beba )< len (_eab ){return _eab [int (*_beba )];};};return nil ;};func _egd (_eef _fd .Time )_fd .Time {_eef =_eef .Local ();return _fd .Date (_eef .Year (),_eef .Month (),_eef .Day (),_eef .Hour (),_eef .Minute (),_eef .Second (),_eef .Nanosecond (),_fd .UTC );};func (_cbdd *Sheet )slideCellsLeft (_bbba []*_ca .CT_Cell )[]*_ca .CT_Cell {for _ ,_ecfd :=range _bbba {_efgd ,_dgad :=_fda .ParseCellReference (*_ecfd .RAttr );if _dgad !=nil {return _bbba ;};_cdcad :=_efgd .ColumnIdx -1;_daeef :=_fda .IndexToColumn (_cdcad )+_e .Sprintf ("\u0025\u0064",_efgd .RowIdx );_ecfd .RAttr =&_daeef ;};return _bbba ;};

// SetWidth is a no-op.
func (_fdga TwoCellAnchor )SetWidth (w _ec .Distance ){};func (_adeb Font )SetItalic (b bool ){if b {_adeb ._ccgc .I =[]*_ca .CT_BooleanProperty {{}};}else {_adeb ._ccgc .I =nil ;};};

// Drawing is a drawing overlay on a sheet.  Only a single drawing is allowed
// per sheet, so to display multiple charts and images on a single sheet, they
// must be added to the same drawing.
type Drawing struct{_fge *Workbook ;_bda *_acd .WsDr ;};

// SetError sets the cell type to error and the value to the given error message.
func (_bcf Cell )SetError (msg string ){_bcf .clearValue ();_bcf ._ecc .V =_b .String (msg );_bcf ._ecc .TAttr =_ca .ST_CellTypeE ;};type DifferentialStyle struct{_ccad *_ca .CT_Dxf ;_aecc *Workbook ;_ddca *_ca .CT_Dxfs ;};

// SetHeightCells sets the height the anchored object by moving the bottom.  It
// is not compatible with SetHeight.
func (_eggd TwoCellAnchor )SetHeightCells (h int32 ){_eggd .SetHeight (0);_fgbee :=_eggd .TopLeft ();_eedc :=_eggd .BottomRight ();_eedc .SetRow (_fgbee .Row ()+h );};

// Column represents a column within a sheet. It's only used for formatting
// purposes, so it's possible to construct a sheet without configuring columns.
type Column struct{_geg *_ca .CT_Col };func (_badb DifferentialStyle )Fill ()Fill {if _badb ._ccad .Fill ==nil {_badb ._ccad .Fill =_ca .NewCT_Fill ();};return Fill {_badb ._ccad .Fill ,nil };};

// PasswordHash returns the hash of the workbook password.
func (_gfag SheetProtection )PasswordHash ()string {if _gfag ._ebec .PasswordAttr ==nil {return "";};return *_gfag ._ebec .PasswordAttr ;};

// SetRow set the row of the cell marker.
func (_dd CellMarker )SetRow (row int32 ){_dd ._gebb .Row =row };

// CellsWithEmpty returns a slice of cells including empty ones from the first column to the last one used in the sheet.
// The cells can be manipulated, but appending to the slice will have no effect.
func (_aca Row )CellsWithEmpty (lastColIdx uint32 )[]Cell {_dcfg :=[]Cell {};for _gge :=uint32 (0);_gge <=lastColIdx ;_gge ++{_abba :=_aca .Cell (_fda .IndexToColumn (_gge ));_dcfg =append (_dcfg ,_abba );};return _dcfg ;};

// ExtractText returns text from the sheet as a SheetText object.
func (_gcaa *Sheet )ExtractText ()*SheetText {_dfac :=[]CellText {};for _ ,_edgg :=range _gcaa .Rows (){for _ ,_eca :=range _edgg .Cells (){if !_eca .IsEmpty (){if _faa :=_eca .GetFormattedValue ();_faa !=""{_dfac =append (_dfac ,CellText {Text :_faa ,Cell :_eca });};};};};return &SheetText {Cells :_dfac };};

// Epoch returns the point at which the dates/times in the workbook are relative to.
func (_ddcdf *Workbook )Epoch ()_fd .Time {if _ddcdf .Uses1904Dates (){_fd .Date (1904,1,1,0,0,0,0,_fd .UTC );};return _fd .Date (1899,12,30,0,0,0,0,_fd .UTC );};

// GetFormula returns the formula for a cell.
func (_afc Cell )GetFormula ()string {if _afc ._ecc .F !=nil {return _afc ._ecc .F .Content ;};return "";};

// IsWindowLocked returns whether the workbook windows are locked.
func (_acafb WorkbookProtection )IsWindowLocked ()bool {return _acafb ._acaf .LockWindowsAttr !=nil &&*_acafb ._acaf .LockWindowsAttr ;};

// AddFormatValue adds a format value to be used to determine the cell background.
func (_bdd ColorScale )AddFormatValue (t _ca .ST_CfvoType ,val string ){_dgbg :=_ca .NewCT_Cfvo ();_dgbg .TypeAttr =t ;_dgbg .ValAttr =_b .String (val );_bdd ._dfdc .Cfvo =append (_bdd ._dfdc .Cfvo ,_dgbg );};

// SetAuthor sets the author of the comment. If the comment body contains the
// author's name (as is the case with Excel and Comments.AddCommentWithStyle, it
// will not be changed).  This method only changes the metadata author of the
// comment.
func (_ecgb Comment )SetAuthor (author string ){_ecgb ._daea .AuthorIdAttr =Comments {_ecgb ._gbf ,_ecgb ._gad }.getOrCreateAuthor (author );};

// IconScale maps values to icons.
type IconScale struct{_dead *_ca .CT_IconSet };var (_ebffb =[...]uint8 {0,21,46,61,76,91};_cfag =[...]uint8 {0,21,37,53,69,85,103,119,135,151,167,185,201,217,239};_gafd =[...]uint8 {0,16,32,48,64};_bfeg =[...]uint8 {0,16,32,48,64,80};);

// Row is a row within a spreadsheet.
type Row struct{_bdaa *Workbook ;_cgee *Sheet ;_gbgd *_ca .CT_Row ;};

// AddConditionalFormatting adds conditional formatting to the sheet.
func (_bggd *Sheet )AddConditionalFormatting (cellRanges []string )ConditionalFormatting {_fdac :=_ca .NewCT_ConditionalFormatting ();_bggd ._cffg .ConditionalFormatting =append (_bggd ._cffg .ConditionalFormatting ,_fdac );_ffb :=make (_ca .ST_Sqref ,0,0);_fdac .SqrefAttr =&_ffb ;for _ ,_ccce :=range cellRanges {*_fdac .SqrefAttr =append (*_fdac .SqrefAttr ,_ccce );};return ConditionalFormatting {_fdac };};

// RemoveSheetByName removes the sheet with the given name from the workbook.
func (_abeb *Workbook )RemoveSheetByName (name string )error {_gbbf :=-1;for _dcbf ,_acefa :=range _abeb .Sheets (){if name ==_acefa .Name (){_gbbf =_dcbf ;break ;};};if _gbbf ==-1{return ErrorNotFound ;};return _abeb .RemoveSheet (_gbbf );};

// DataValidationList is just a view on a DataValidation configured as a list.
// It presents a drop-down combo box for spreadsheet users to select values. The
// contents of the dropdown can either pull from a rang eof cells (SetRange) or
// specified directly (SetValues).
type DataValidationList struct{_dce *_ca .CT_DataValidation };

// AddCellStyle adds a new empty cell style to the stylesheet.
func (_dfff StyleSheet )AddCellStyle ()CellStyle {_gbfa :=_ca .NewCT_Xf ();_dfff ._gace .CellXfs .Xf =append (_dfff ._gace .CellXfs .Xf ,_gbfa );_dfff ._gace .CellXfs .CountAttr =_b .Uint32 (uint32 (len (_dfff ._gace .CellXfs .Xf )));return CellStyle {_dfff ._efga ,_gbfa ,_dfff ._gace .CellXfs };};

// Sort sorts all of the rows within a sheet by the contents of a column. As the
// file format doesn't suppot indicating that a column should be sorted by the
// viewing/editing program, we actually need to reorder rows and change cell
// references during a sort. If the sheet contains formulas, you should call
// RecalculateFormulas() prior to sorting.  The column is in the form "C" and
// specifies the column to sort by. The firstRow is a 1-based index and
// specifies the firstRow to include in the sort, allowing skipping over a
// header row.
func (_daee *Sheet )Sort (column string ,firstRow uint32 ,order SortOrder ){_bcdce :=_daee ._cffg .SheetData .Row ;_gcfcb :=_daee .Rows ();for _efdd ,_gfdda :=range _gcfcb {if _gfdda .RowNumber ()==firstRow {_bcdce =_daee ._cffg .SheetData .Row [_efdd :];break ;};};_aeda :=Comparer {Order :order };_c .Slice (_bcdce ,func (_dgga ,_fdbdd int )bool {return _aeda .LessRows (column ,Row {_daee ._fdfd ,_daee ,_bcdce [_dgga ]},Row {_daee ._fdfd ,_daee ,_bcdce [_fdbdd ]});});for _baca ,_ddfac :=range _daee .Rows (){_bfaa :=uint32 (_baca +1);if _ddfac .RowNumber ()!=_bfaa {_ddfac .renumberAs (_bfaa );};};};

// AddNumberedRow adds a row with a given row number.  If you reuse a row number
// the resulting file will fail validation and fail to open in Office programs. Use
// Row instead which creates a new row or returns an existing row.
func (_bgeg *Sheet )AddNumberedRow (rowNum uint32 )Row {_gbga :=_ca .NewCT_Row ();_gbga .RAttr =_b .Uint32 (rowNum );_bgeg ._cffg .SheetData .Row =append (_bgeg ._cffg .SheetData .Row ,_gbga );_c .Slice (_bgeg ._cffg .SheetData .Row ,func (_gcde ,_dfag int )bool {_efgf :=_bgeg ._cffg .SheetData .Row [_gcde ].RAttr ;_decbf :=_bgeg ._cffg .SheetData .Row [_dfag ].RAttr ;if _efgf ==nil {return true ;};if _decbf ==nil {return true ;};return *_efgf < *_decbf ;});return Row {_bgeg ._fdfd ,_bgeg ,_gbga };};func (_cbe CellStyle )SetNumberFormat (s string ){_cee :=_cbe ._dcb .StyleSheet .AddNumberFormat ();_cee .SetFormat (s );_cbe ._fbg .ApplyNumberFormatAttr =_b .Bool (true );_cbe ._fbg .NumFmtIdAttr =_b .Uint32 (_cee .ID ());};

// LockWindow controls the locking of the workbook windows.
func (_gdeb WorkbookProtection )LockWindow (b bool ){if !b {_gdeb ._acaf .LockWindowsAttr =nil ;}else {_gdeb ._acaf .LockWindowsAttr =_b .Bool (true );};};

// SetString sets the cell type to string, and the value to the given string,
// returning an ID from the shared strings table. To reuse a string, call
// SetStringByID with the ID returned.
func (_fdeb Cell )SetString (s string )int {_fdeb ._fac .ensureSharedStringsRelationships ();_fdeb .clearValue ();_bgf :=_fdeb ._fac .SharedStrings .AddString (s );_fdeb ._ecc .V =_b .String (_fb .Itoa (_bgf ));_fdeb ._ecc .TAttr =_ca .ST_CellTypeS ;return _bgf ;};func (_fceb Row )renumberAs (_gabg uint32 ){_fceb ._gbgd .RAttr =_b .Uint32 (_gabg );for _ ,_gcfb :=range _fceb .Cells (){_cgbg ,_agbd :=_fda .ParseCellReference (_gcfb .Reference ());if _agbd ==nil {_eeca :=_e .Sprintf ("\u0025\u0073\u0025\u0064",_cgbg .Column ,_gabg );_gcfb ._ecc .RAttr =_b .String (_eeca );};};};

// GetFormat sets the number format code.
func (_cbeb NumberFormat )GetFormat ()string {return _cbeb ._bgfa .FormatCodeAttr };

// X returns the inner wrapped XML type.
func (_dfef CellMarker )X ()*_acd .CT_Marker {return _dfef ._gebb };

// Read reads a workbook from an io.Reader(.xlsx).
func Read (r _f .ReaderAt ,size int64 )(*Workbook ,error ){const _ggdb ="\u0073\u0070r\u0065\u0061\u0064s\u0068\u0065\u0065\u0074\u003a\u0052\u0065\u0061\u0064";if !_cb .GetLicenseKey ().IsLicensed ()&&!_cedb {_e .Println ("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065");_e .Println ("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f");return nil ,_dec .New ("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064");};_daae :=New ();_abgg ,_dabf :=_cb .GenRefId ("\u0073\u0072");if _dabf !=nil {_cd .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_dabf );return nil ,_dabf ;};_daae ._fcgbe =_abgg ;if _ddfa :=_cb .Track (_daae ._fcgbe ,_ggdb );_ddfa !=nil {_cd .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ddfa );return nil ,_ddfa ;};_caaa ,_dabf :=_cdc .TempDir ("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0078\u006c\u0073\u0078");if _dabf !=nil {return nil ,_dabf ;};_daae .TmpPath =_caaa ;_fagg ,_dabf :=_gg .NewReader (r ,size );if _dabf !=nil {return nil ,_e .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073",_dabf );};_dcd :=[]*_gg .File {};_dcd =append (_dcd ,_fagg .File ...);_gfdd :=false ;for _ ,_eccc :=range _dcd {if _eccc .FileHeader .Name =="\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c"{_gfdd =true ;break ;};};if _gfdd {_daae .CreateCustomProperties ();};_fbgb :=_ce .DecodeMap {};_fbgb .SetOnNewRelationshipFunc (_daae .onNewRelationship );_fbgb .AddTarget (_b .ContentTypesFilename ,_daae .ContentTypes .X (),"",0);_fbgb .AddTarget (_b .BaseRelsFilename ,_daae .Rels .X (),"",0);if _fbda :=_fbgb .Decode (_dcd );_fbda !=nil {return nil ,_fbda ;};for _ ,_geae :=range _dcd {if _geae ==nil {continue ;};if _egff :=_daae .AddExtraFileFromZip (_geae );_egff !=nil {return nil ,_egff ;};};if _gfdd {_fce :=false ;for _ ,_gfdef :=range _daae .Rels .X ().Relationship {if _gfdef .TargetAttr =="\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c"{_fce =true ;break ;};};if !_fce {_daae .AddCustomRelationships ();};};return _daae ,nil ;};

// SetReference sets the regin of cells that the merged cell applies to.
func (_bcadc MergedCell )SetReference (ref string ){_bcadc ._fgf .RefAttr =ref };

// AddDefinedName adds a name for a cell or range reference that can be used in
// formulas and charts.
func (_gfegb *Workbook )AddDefinedName (name ,ref string )DefinedName {if _gfegb ._gdeg .DefinedNames ==nil {_gfegb ._gdeg .DefinedNames =_ca .NewCT_DefinedNames ();};_fagf :=_ca .NewCT_DefinedName ();_fagf .Content =ref ;_fagf .NameAttr =name ;_gfegb ._gdeg .DefinedNames .DefinedName =append (_gfegb ._gdeg .DefinedNames .DefinedName ,_fagf );return DefinedName {_fagf };};

// NewSharedStrings constructs a new Shared Strings table.
func NewSharedStrings ()SharedStrings {return SharedStrings {_cfca :_ca .NewSst (),_fbfea :make (map[string ]int )};};

// Save writes the workbook out to a writer in the zipped xlsx format.
func (_cdeb *Workbook )Save (w _f .Writer )error {const _cgcc ="\u0073\u0070\u0072\u0065ad\u0073\u0068\u0065\u0065\u0074\u003a\u0077\u0062\u002e\u0053\u0061\u0076\u0065";if !_cb .GetLicenseKey ().IsLicensed ()&&!_cedb {_e .Println ("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065");_e .Println ("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f");return _dec .New ("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064");};if len (_cdeb ._fcgbe )==0{_afgd ,_ceeed :=_cb .GenRefId ("\u0073\u0077");if _ceeed !=nil {_cd .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ceeed );return _ceeed ;};_cdeb ._fcgbe =_afgd ;};if _bece :=_cb .Track (_cdeb ._fcgbe ,_cgcc );_bece !=nil {_cd .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_bece );return _bece ;};_ccbf :=_gg .NewWriter (w );defer _ccbf .Close ();_abbae :=_b .DocTypeSpreadsheet ;if _aggf :=_ce .MarshalXML (_ccbf ,_b .BaseRelsFilename ,_cdeb .Rels .X ());_aggf !=nil {return _aggf ;};if _egffe :=_ce .MarshalXMLByType (_ccbf ,_abbae ,_b .ExtendedPropertiesType ,_cdeb .AppProperties .X ());_egffe !=nil {return _egffe ;};if _ggab :=_ce .MarshalXMLByType (_ccbf ,_abbae ,_b .CorePropertiesType ,_cdeb .CoreProperties .X ());_ggab !=nil {return _ggab ;};_afbb :=_b .AbsoluteFilename (_abbae ,_b .OfficeDocumentType ,0);if _cdcf :=_ce .MarshalXML (_ccbf ,_afbb ,_cdeb ._gdeg );_cdcf !=nil {return _cdcf ;};if _fded :=_ce .MarshalXML (_ccbf ,_ce .RelationsPathFor (_afbb ),_cdeb ._eeead .X ());_fded !=nil {return _fded ;};if _cdbf :=_ce .MarshalXMLByType (_ccbf ,_abbae ,_b .StylesType ,_cdeb .StyleSheet .X ());_cdbf !=nil {return _cdbf ;};for _deaeg ,_aefe :=range _cdeb ._gfcg {if _adcc :=_ce .MarshalXMLByTypeIndex (_ccbf ,_abbae ,_b .ThemeType ,_deaeg +1,_aefe );_adcc !=nil {return _adcc ;};};for _dbgc ,_bbca :=range _cdeb ._bbced {_bbca .Dimension .RefAttr =Sheet {_cdeb ,nil ,_bbca }.Extents ();_gcef :=_b .AbsoluteFilename (_abbae ,_b .WorksheetType ,_dbgc +1);_ce .MarshalXML (_ccbf ,_gcef ,_bbca );_ce .MarshalXML (_ccbf ,_ce .RelationsPathFor (_gcef ),_cdeb ._dcdg [_dbgc ].X ());};if _cfeb :=_ce .MarshalXMLByType (_ccbf ,_abbae ,_b .SharedStringsType ,_cdeb .SharedStrings .X ());_cfeb !=nil {return _cfeb ;};if _cdeb .CustomProperties .X ()!=nil {if _adeea :=_ce .MarshalXMLByType (_ccbf ,_abbae ,_b .CustomPropertiesType ,_cdeb .CustomProperties .X ());_adeea !=nil {return _adeea ;};};if _cdeb .Thumbnail !=nil {_egca :=_b .AbsoluteFilename (_abbae ,_b .ThumbnailType ,0);_bbgc ,_ddbe :=_ccbf .Create (_egca );if _ddbe !=nil {return _ddbe ;};if _ddda :=_dg .Encode (_bbgc ,_cdeb .Thumbnail ,nil );_ddda !=nil {return _ddda ;};};for _efaff ,_gddb :=range _cdeb ._dgeg {_dfgga :=_b .AbsoluteFilename (_abbae ,_b .ChartType ,_efaff +1);_ce .MarshalXML (_ccbf ,_dfgga ,_gddb );};for _cdeag ,_adce :=range _cdeb ._beee {_ecebc :=_b .AbsoluteFilename (_abbae ,_b .TableType ,_cdeag +1);_ce .MarshalXML (_ccbf ,_ecebc ,_adce );};for _ccbe ,_fcaf :=range _cdeb ._cffgf {_cgfe :=_b .AbsoluteFilename (_abbae ,_b .DrawingType ,_ccbe +1);_ce .MarshalXML (_ccbf ,_cgfe ,_fcaf );if !_cdeb ._aeec [_ccbe ].IsEmpty (){_ce .MarshalXML (_ccbf ,_ce .RelationsPathFor (_cgfe ),_cdeb ._aeec [_ccbe ].X ());};};for _baef ,_geec :=range _cdeb ._dbff {_ce .MarshalXML (_ccbf ,_b .AbsoluteFilename (_abbae ,_b .VMLDrawingType ,_baef +1),_geec );};for _egagb ,_eecdf :=range _cdeb .Images {if _bdf :=_da .AddImageToZip (_ccbf ,_eecdf ,_egagb +1,_b .DocTypeSpreadsheet );_bdf !=nil {return _bdf ;};};if _gddg :=_ce .MarshalXML (_ccbf ,_b .ContentTypesFilename ,_cdeb .ContentTypes .X ());_gddg !=nil {return _gddg ;};for _aedgc ,_gdgf :=range _cdeb ._dfgf {if _gdgf ==nil {continue ;};_ce .MarshalXML (_ccbf ,_b .AbsoluteFilename (_abbae ,_b .CommentsType ,_aedgc +1),_gdgf );};if _egee :=_cdeb .WriteExtraFiles (_ccbf );_egee !=nil {return _egee ;};return _ccbf .Close ();};

// SetBold causes the text to be displayed in bold.
func (_ebac RichTextRun )SetBold (b bool ){_ebac .ensureRpr ();_ebac ._dfc .RPr .B =_ca .NewCT_BooleanProperty ();_ebac ._dfc .RPr .B .ValAttr =_b .Bool (b );};

// LockObject controls the locking of the sheet objects.
func (_fgbdc SheetProtection )LockObject (b bool ){if !b {_fgbdc ._ebec .ObjectsAttr =nil ;}else {_fgbdc ._ebec .ObjectsAttr =_b .Bool (true );};};

// AddRow adds a new row to a sheet.  You can mix this with numbered rows,
// however it will get confusing. You should prefer to use either automatically
// numbered rows with AddRow or manually numbered rows with Row/AddNumberedRow
func (_agcb *Sheet )AddRow ()Row {_dcgb :=uint32 (0);_eag :=uint32 (len (_agcb ._cffg .SheetData .Row ));if _eag > 0&&_agcb ._cffg .SheetData .Row [_eag -1].RAttr !=nil &&*_agcb ._cffg .SheetData .Row [_eag -1].RAttr ==_eag {return _agcb .addNumberedRowFast (_eag +1);};for _ ,_ebbf :=range _agcb ._cffg .SheetData .Row {if _ebbf .RAttr !=nil &&*_ebbf .RAttr > _dcgb {_dcgb =*_ebbf .RAttr ;};};return _agcb .AddNumberedRow (_dcgb +1);};func _gcec (_gedd string ,_eggc uint32 ,_bfgg bool )string {_cbddd ,_ffge ,_eced :=_fda .ParseRangeReference (_gedd );if _eced ==nil {_ceeba ,_ebfe :=_cbddd .ColumnIdx ,_ffge .ColumnIdx ;if _eggc >=_ceeba &&_eggc <=_ebfe {if _ceeba ==_ebfe {if _bfgg {return "";}else {return _gedd ;};}else {_fgfc :=_ffge .Update (_df .UpdateActionRemoveColumn );return _e .Sprintf ("\u0025\u0073\u003a%\u0073",_cbddd .String (),_fgfc .String ());};}else if _eggc < _ceeba {_dbed :=_cbddd .Update (_df .UpdateActionRemoveColumn );_egfg :=_ffge .Update (_df .UpdateActionRemoveColumn );return _e .Sprintf ("\u0025\u0073\u003a%\u0073",_dbed .String (),_egfg .String ());};}else {_ffa ,_ceebd ,_aga :=_fda .ParseColumnRangeReference (_gedd );if _aga !=nil {return "";};_gegd ,_bfebf :=_ffa .ColumnIdx ,_ceebd .ColumnIdx ;if _eggc >=_gegd &&_eggc <=_bfebf {if _gegd ==_bfebf {if _bfgg {return "";}else {return _gedd ;};}else {_cadc :=_ceebd .Update (_df .UpdateActionRemoveColumn );return _e .Sprintf ("\u0025\u0073\u003a%\u0073",_ffa .String (),_cadc .String ());};}else if _eggc < _gegd {_dafg :=_ffa .Update (_df .UpdateActionRemoveColumn );_dbbf :=_ceebd .Update (_df .UpdateActionRemoveColumn );return _e .Sprintf ("\u0025\u0073\u003a%\u0073",_dafg .String (),_dbbf .String ());};};return "";};

// CellMarker represents a cell position
type CellMarker struct{_gebb *_acd .CT_Marker };

// MakeComments constructs a new Comments wrapper.
func MakeComments (w *Workbook ,x *_ca .Comments )Comments {return Comments {w ,x }};

// Type returns the type of anchor
func (_gb AbsoluteAnchor )Type ()AnchorType {return AnchorTypeAbsolute };

// RemoveFont removes a font from the style sheet.  It *does not* update styles that refer
// to this font.
func (_aadc StyleSheet )RemoveFont (f Font )error {for _fcac ,_ffec :=range _aadc ._gace .Fonts .Font {if _ffec ==f .X (){_aadc ._gace .Fonts .Font =append (_aadc ._gace .Fonts .Font [:_fcac ],_aadc ._gace .Fonts .Font [_fcac +1:]...);return nil ;};};return _dec .New ("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};func (_gfcfe PatternFill )SetBgColor (c _fde .Color ){_gfcfe ._aedg .BgColor =_ca .NewCT_Color ();_gfcfe ._aedg .BgColor .RgbAttr =c .AsRGBAString ();};

// SetFormulaShared sets the cell type to formula shared, and the raw formula to
// the given string. The range is the range of cells that the formula applies
// to, and is used to conserve disk space.
func (_gc Cell )SetFormulaShared (formulaStr string ,rows ,cols uint32 )error {_acdg :=_ab .ParseString (formulaStr );if _acdg ==nil {return _dec .New (_e .Sprintf ("\u0043a\u006en\u006f\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0073",formulaStr ));};_gc .clearValue ();_gc ._ecc .TAttr =_ca .ST_CellTypeStr ;_gc ._ecc .F =_ca .NewCT_CellFormula ();_gc ._ecc .F .TAttr =_ca .ST_CellFormulaTypeShared ;_gc ._ecc .F .Content =formulaStr ;_eea ,_dc :=_fda .ParseCellReference (_gc .Reference ());if _dc !=nil {return _dc ;};_ceff :=uint32 (0);for _ ,_fag :=range _gc ._cfd .Rows (){for _ ,_fbb :=range _fag ._gbgd .C {if _fbb .F !=nil &&_fbb .F .SiAttr !=nil &&*_fbb .F .SiAttr >=_ceff {_ceff =*_fbb .F .SiAttr ;};};};_ceff ++;_bcg :=_e .Sprintf ("\u0025s\u0025\u0064\u003a\u0025\u0073\u0025d",_eea .Column ,_eea .RowIdx ,_fda .IndexToColumn (_eea .ColumnIdx +cols ),_eea .RowIdx +rows );_gc ._ecc .F .RefAttr =_b .String (_bcg );_gc ._ecc .F .SiAttr =_b .Uint32 (_ceff );_cec :=Sheet {_gc ._fac ,_gc ._cfd ._aacc ,_gc ._cfd ._cffg };for _gcf :=_eea .RowIdx ;_gcf <=_eea .RowIdx +rows ;_gcf ++{for _gff :=_eea .ColumnIdx ;_gff <=_eea .ColumnIdx +cols ;_gff ++{if _gcf ==_eea .RowIdx &&_gff ==_eea .ColumnIdx {continue ;};_ff :=_e .Sprintf ("\u0025\u0073\u0025\u0064",_fda .IndexToColumn (_gff ),_gcf );_cec .Cell (_ff ).Clear ();_cec .Cell (_ff ).X ().F =_ca .NewCT_CellFormula ();_cec .Cell (_ff ).X ().F .TAttr =_ca .ST_CellFormulaTypeShared ;_cec .Cell (_ff ).X ().F .SiAttr =_b .Uint32 (_ceff );};};return nil ;};

// SetHidden marks the defined name as hidden.
func (_cag DefinedName )SetHidden (b bool ){_cag ._dfg .HiddenAttr =_b .Bool (b )};

// SetHeight sets the height of the anchored object.
func (_fe AbsoluteAnchor )SetHeight (h _ec .Distance ){_fe ._ag .Ext .CyAttr =int64 (h /_ec .EMU )};

// RemoveSheet removes the sheet with the given index from the workbook.
func (_fbff *Workbook )RemoveSheet (ind int )error {if _fbff .SheetCount ()<=ind {return ErrorNotFound ;};for _ ,_abfe :=range _fbff ._eeead .Relationships (){if _abfe .ID ()==_fbff ._gdeg .Sheets .Sheet [ind ].IdAttr {_fbff ._eeead .Remove (_abfe );break ;};};_fbff .ContentTypes .RemoveOverride (_b .AbsoluteFilename (_b .DocTypeSpreadsheet ,_b .WorksheetContentType ,ind +1));copy (_fbff ._bbced [ind :],_fbff ._bbced [ind +1:]);_fbff ._bbced =_fbff ._bbced [:len (_fbff ._bbced )-1];_acb :=_fbff ._gdeg .Sheets .Sheet [ind ];copy (_fbff ._gdeg .Sheets .Sheet [ind :],_fbff ._gdeg .Sheets .Sheet [ind +1:]);_fbff ._gdeg .Sheets .Sheet =_fbff ._gdeg .Sheets .Sheet [:len (_fbff ._gdeg .Sheets .Sheet )-1];for _aecb :=range _fbff ._gdeg .Sheets .Sheet {if _fbff ._gdeg .Sheets .Sheet [_aecb ].SheetIdAttr > _acb .SheetIdAttr {_fbff ._gdeg .Sheets .Sheet [_aecb ].SheetIdAttr --;};};copy (_fbff ._dcdg [ind :],_fbff ._dcdg [ind +1:]);_fbff ._dcdg =_fbff ._dcdg [:len (_fbff ._dcdg )-1];copy (_fbff ._dfgf [ind :],_fbff ._dfgf [ind +1:]);_fbff ._dfgf =_fbff ._dfgf [:len (_fbff ._dfgf )-1];return nil ;};

// Reference returns the cell reference (e.g. "A4"). This is not required,
// however both unioffice and Excel will always set it.
func (_cc Cell )Reference ()string {if _cc ._ecc .RAttr !=nil {return *_cc ._ecc .RAttr ;};return "";};

// GetDrawing return the worksheet drawing and its relationships if exists.
func (_cacg *Sheet )GetDrawing ()(*_acd .WsDr ,_da .Relationships ){if _addb :=_cacg ._cffg .Drawing ;_addb !=nil {_cfab :=0;for _ ,_eagg :=range _cacg ._fdfd ._bbced {if _dceg :=_eagg .Drawing ;_dceg !=nil {if _eagg ==_cacg ._cffg {return _cacg ._fdfd ._cffgf [_cfab ],_cacg ._fdfd ._aeec [_cfab ];};_cfab ++;};};};return nil ,_da .Relationships {};};

// SetPattern sets the pattern of the fill.
func (_fbfd PatternFill )SetPattern (p _ca .ST_PatternType ){_fbfd ._aedg .PatternTypeAttr =p };

// RowOffset returns the offset from the row cell.
func (_febc CellMarker )RowOffset ()_ec .Distance {if _febc ._gebb .RowOff .ST_CoordinateUnqualified ==nil {return 0;};return _ec .Distance (float64 (*_febc ._gebb .RowOff .ST_CoordinateUnqualified )*_ec .EMU );};

// X returns the inner wrapped XML type.
func (_fcgg RichText )X ()*_ca .CT_Rst {return _fcgg ._gef };type WorkbookProtection struct{_acaf *_ca .CT_WorkbookProtection };

// GetWidth returns a worksheet's column width.
func (_dbc *evalContext )GetWidth (colIdx int )float64 {colIdx ++;for _ ,_aeeba :=range _dbc ._bcdc .X ().Cols [0].Col {if int (_aeeba .MinAttr )<=colIdx &&colIdx <=int (_aeeba .MaxAttr ){return float64 (int (*_aeeba .WidthAttr ));};};return 0;};

// SheetViews returns the sheet views defined.  This is where splits and frozen
// rows/cols are configured.  Multiple sheet views are allowed, but I'm not
// aware of there being a use for more than a single sheet view.
func (_cbgf *Sheet )SheetViews ()[]SheetView {if _cbgf ._cffg .SheetViews ==nil {return nil ;};_eceeg :=[]SheetView {};for _ ,_bbfg :=range _cbgf ._cffg .SheetViews .SheetView {_eceeg =append (_eceeg ,SheetView {_bbfg });};return _eceeg ;};func (_dfgcb *Sheet )updateAfterRemove (_defbg uint32 ,_dbefg _df .UpdateAction )error {_gada :=_dfgcb .Name ();_gfae :=&_df .UpdateQuery {UpdateType :_dbefg ,ColumnIdx :_defbg ,SheetToUpdate :_gada };for _ ,_eadd :=range _dfgcb ._fdfd .Sheets (){_gfae .UpdateCurrentSheet =_gada ==_eadd .Name ();for _ ,_cbga :=range _eadd .Rows (){for _ ,_efe :=range _cbga .Cells (){if _efe .X ().F !=nil {_fcbg :=_efe .X ().F .Content ;_gcbg :=_ab .ParseString (_fcbg );if _gcbg ==nil {_efe .SetError ("\u0023\u0052\u0045F\u0021");}else {_gage :=_gcbg .Update (_gfae );_efe .X ().F .Content =_e .Sprintf ("\u003d\u0025\u0073",_gage .String ());};};};};};return nil ;};

// TopLeft is a no-op.
func (_eg AbsoluteAnchor )TopLeft ()CellMarker {return CellMarker {}};

// GetFilename returns the name of file from which workbook was opened with full path to it
func (_agbda *Workbook )GetFilename ()string {return _agbda ._eeaf };

// AddView adds a sheet view.
func (_ccfe *Sheet )AddView ()SheetView {if _ccfe ._cffg .SheetViews ==nil {_ccfe ._cffg .SheetViews =_ca .NewCT_SheetViews ();};_aeac :=_ca .NewCT_SheetView ();_ccfe ._cffg .SheetViews .SheetView =append (_ccfe ._cffg .SheetViews .SheetView ,_aeac );return SheetView {_aeac };};

// Author returns the author of the comment
func (_gbfd Comment )Author ()string {if _gbfd ._daea .AuthorIdAttr < uint32 (len (_gbfd ._gad .Authors .Author )){return _gbfd ._gad .Authors .Author [_gbfd ._daea .AuthorIdAttr ];};return "";};

// AddNumberFormat adds a new blank number format to the stylesheet.
func (_fafe StyleSheet )AddNumberFormat ()NumberFormat {if _fafe ._gace .NumFmts ==nil {_fafe ._gace .NumFmts =_ca .NewCT_NumFmts ();};_fabb :=_ca .NewCT_NumFmt ();_fabb .NumFmtIdAttr =uint32 (200+len (_fafe ._gace .NumFmts .NumFmt ));_fafe ._gace .NumFmts .NumFmt =append (_fafe ._gace .NumFmts .NumFmt ,_fabb );_fafe ._gace .NumFmts .CountAttr =_b .Uint32 (uint32 (len (_fafe ._gace .NumFmts .NumFmt )));return NumberFormat {_fafe ._efga ,_fabb };};

// SetStyle sets the cell style for an entire column.
func (_ggcd Column )SetStyle (cs CellStyle ){_ggcd ._geg .StyleAttr =_b .Uint32 (cs .Index ())};

// SetUnderline controls if the run is underlined.
func (_gcda RichTextRun )SetUnderline (u _ca .ST_UnderlineValues ){_gcda .ensureRpr ();_gcda ._dfc .RPr .U =_ca .NewCT_UnderlineProperty ();_gcda ._dfc .RPr .U .ValAttr =u ;};

// LastColumn returns the name of last column which contains data in range of context sheet's given rows.
func (_beda *evalContext )LastColumn (rowFrom ,rowTo int )string {_gdf :=_beda ._bcdc ;_cac :=1;for _fgg :=rowFrom ;_fgg <=rowTo ;_fgg ++{_ffeb :=len (_gdf .Row (uint32 (_fgg )).Cells ());if _ffeb > _cac {_cac =_ffeb ;};};return _fda .IndexToColumn (uint32 (_cac -1));};func _cfec ()*_acd .CT_AbsoluteAnchor {_cgd :=_acd .NewCT_AbsoluteAnchor ();return _cgd };

// SetPasswordHash sets the password hash to the input.
func (_ccbaf SheetProtection )SetPasswordHash (pwHash string ){_ccbaf ._ebec .PasswordAttr =_b .String (pwHash );};

// Type returns the type of anchor
func (_dcfa OneCellAnchor )Type ()AnchorType {return AnchorTypeOneCell };func (_daa CellStyle )Index ()uint32 {for _aadd ,_dgg :=range _daa ._ddc .Xf {if _daa ._fbg ==_dgg {return uint32 (_aadd );};};return 0;};

// SetType sets the type of the rule.
func (_dbf ConditionalFormattingRule )SetType (t _ca .ST_CfType ){_dbf ._afcg .TypeAttr =t };func (_egec DataValidation )SetList ()DataValidationList {_egec .clear ();_egec ._ege .TypeAttr =_ca .ST_DataValidationTypeList ;_egec ._ege .OperatorAttr =_ca .ST_DataValidationOperatorEqual ;return DataValidationList {_egec ._ege };};

// AddFormatValue adds a format value (databars require two).
func (_daag DataBarScale )AddFormatValue (t _ca .ST_CfvoType ,val string ){_bfda :=_ca .NewCT_Cfvo ();_bfda .TypeAttr =t ;_bfda .ValAttr =_b .String (val );_daag ._bdcd .Cfvo =append (_daag ._bdcd .Cfvo ,_bfda );};

// SetWidthCells is a no-op.
func (_af AbsoluteAnchor )SetWidthCells (int32 ){};

// NumberFormat returns the number format that the cell style uses, or zero if
// it is not set.
func (_bege CellStyle )NumberFormat ()uint32 {if _bege ._fbg .NumFmtIdAttr ==nil {return 0;};return *_bege ._fbg .NumFmtIdAttr ;};

// SetContent sets the defined name content.
func (_dcac DefinedName )SetContent (s string ){_dcac ._dfg .Content =s };

// Column returns or creates a column that with a given index (1-N).  Columns
// can span multiple column indices, this method will return the column that
// applies to a column index if it exists or create a new column that only
// applies to the index passed in otherwise.
func (_ceeb *Sheet )Column (idx uint32 )Column {for _ ,_gcgb :=range _ceeb ._cffg .Cols {for _ ,_gbag :=range _gcgb .Col {if idx >=_gbag .MinAttr &&idx <=_gbag .MaxAttr {return Column {_gbag };};};};var _caccg *_ca .CT_Cols ;if len (_ceeb ._cffg .Cols )==0{_caccg =_ca .NewCT_Cols ();_ceeb ._cffg .Cols =append (_ceeb ._cffg .Cols ,_caccg );}else {_caccg =_ceeb ._cffg .Cols [0];};_dfgd :=_ca .NewCT_Col ();_dfgd .MinAttr =idx ;_dfgd .MaxAttr =idx ;_caccg .Col =append (_caccg .Col ,_dfgd );return Column {_dfgd };};

// SetHidden controls the visibility of a column.
func (_cgeg Column )SetHidden (b bool ){if !b {_cgeg ._geg .HiddenAttr =nil ;}else {_cgeg ._geg .HiddenAttr =_b .Bool (true );};};

// SetValues sets the possible values. This is incompatible with SetRange.
func (_bbbc DataValidationList )SetValues (values []string ){_bbbc ._dce .Formula1 =_b .String ("\u0022"+_fa .Join (values ,"\u002c")+"\u0022");_bbbc ._dce .Formula2 =_b .String ("\u0030");};

// SheetView is a view of a sheet. There is typically one per sheet, though more
// are supported.
type SheetView struct{_egcb *_ca .CT_SheetView };

// SetColOffset sets the column offset of the two cell anchor.
func (_eeea TwoCellAnchor )SetColOffset (m _ec .Distance ){_gaef :=m -_eeea .TopLeft ().ColOffset ();_eeea .TopLeft ().SetColOffset (m );_eeea .BottomRight ().SetColOffset (_eeea .BottomRight ().ColOffset ()+_gaef );};

// MoveTo is a no-op.
func (_acf AbsoluteAnchor )MoveTo (x ,y int32 ){};func (_age Comments )getOrCreateAuthor (_febca string )uint32 {for _caae ,_eecd :=range _age ._cfc .Authors .Author {if _eecd ==_febca {return uint32 (_caae );};};_ffe :=uint32 (len (_age ._cfc .Authors .Author ));_age ._cfc .Authors .Author =append (_age ._cfc .Authors .Author ,_febca );return _ffe ;};

// X returns the inner wrapped XML type.
func (_bcb ConditionalFormattingRule )X ()*_ca .CT_CfRule {return _bcb ._afcg };
