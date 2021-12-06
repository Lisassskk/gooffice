
package reference ;import (_fe "errors";_eg "fmt";_fg "gitee.com/gooffice/gooffice/spreadsheet/update";_e "regexp";_d "strconv";_ea "strings";);

// String returns a string representation of CellReference.
func (_b CellReference )String ()string {_g :=make ([]byte ,0,4);if _b .AbsoluteColumn {_g =append (_g ,'$');};_g =append (_g ,_b .Column ...);if _b .AbsoluteRow {_g =append (_g ,'$');};_g =_d .AppendInt (_g ,int64 (_b .RowIdx ),10);return string (_g );};

// ColumnReference is a parsed reference to a column.  Input is of the form 'A',
// '$C', etc.
type ColumnReference struct{ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;SheetName string ;};

// Update updates reference to point one of the neighboring columns with respect to the update type after removing a row/column.
func (_dbf *ColumnReference )Update (updateType _fg .UpdateAction )*ColumnReference {switch updateType {case _fg .UpdateActionRemoveColumn :_gd :=_dbf ;_gd .ColumnIdx =_dbf .ColumnIdx -1;_gd .Column =IndexToColumn (_gd .ColumnIdx );return _gd ;default:return _dbf ;};};

// CellReference is a parsed reference to a cell.  Input is of the form 'A1',
// '$C$2', etc.
type CellReference struct{RowIdx uint32 ;ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;AbsoluteRow bool ;SheetName string ;};

// String returns a string representation of ColumnReference.
func (_de ColumnReference )String ()string {_bg :=make ([]byte ,0,4);if _de .AbsoluteColumn {_bg =append (_bg ,'$');};_bg =append (_bg ,_de .Column ...);return string (_bg );};

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference (s string )(_ee ,_bc CellReference ,_cd error ){_afd ,_ed ,_cd :=_gf (s );if _cd !=nil {return CellReference {},CellReference {},_cd ;};_df :=_ea .Split (_ed ,"\u003a");if len (_df )!=2{return CellReference {},CellReference {},_fe .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");};if _afd !=""{_df [0]=_afd +"\u0021"+_df [0];_df [1]=_afd +"\u0021"+_df [1];};_deg ,_cd :=ParseCellReference (_df [0]);if _cd !=nil {return CellReference {},CellReference {},_cd ;};_gff ,_cd :=ParseCellReference (_df [1]);if _cd !=nil {return CellReference {},CellReference {},_cd ;};return _deg ,_gff ,nil ;};

// ParseColumnReference parses a column reference of the form 'Sheet1!A' and splits it
// into sheet name and column segments.
func ParseColumnReference (s string )(ColumnReference ,error ){s =_ea .TrimSpace (s );if len (s )< 1{return ColumnReference {},_fe .New ("\u0063\u006f\u006c\u0075\u006d\u006e \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0063\u0068a\u0072a\u0063\u0074\u0065\u0072");};_af :=ColumnReference {};_ca ,_bf ,_ff :=_gf (s );if _ff !=nil {return ColumnReference {},_ff ;};if _ca !=""{_af .SheetName =_ca ;};if _bf [0]=='$'{_af .AbsoluteColumn =true ;_bf =_bf [1:];};if !_dg .MatchString (_bf ){return ColumnReference {},_fe .New ("\u0063\u006f\u006c\u0075\u006dn\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0041\u0020\u0061\u006e\u0064\u0020\u005a\u005a");};_af .Column =_bf ;_af .ColumnIdx =ColumnToIndex (_af .Column );return _af ,nil ;};

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex (col string )uint32 {col =_ea .ToUpper (col );_ad :=uint32 (0);for _ ,_gc :=range col {_ad *=26;_ad +=uint32 (_gc -'A'+1);};return _ad -1;};

// IndexToColumn maps a column number to a column name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn (col uint32 )string {var _ec [64+1]byte ;_ge :=len (_ec );_dbd :=col ;const _dee =26;for _dbd >=_dee {_ge --;_cge :=_dbd /_dee ;_ec [_ge ]=byte ('A'+uint (_dbd -_cge *_dee ));_dbd =_cge -1;};_ge --;_ec [_ge ]=byte ('A'+uint (_dbd ));return string (_ec [_ge :]);};func _gf (_ef string )(string ,string ,error ){_ce :="";_gdd :=_ea .LastIndex (_ef ,"\u0021");if _gdd > -1{_ce =_ef [:_gdd ];_ef =_ef [_gdd +1:];if _ce ==""{return "","",_fe .New ("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0073h\u0065\u0065t\u0020\u006e\u0061\u006d\u0065");};};return _ce ,_ef ,nil ;};

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference (s string )(CellReference ,error ){s =_ea .TrimSpace (s );if len (s )< 2{return CellReference {},_fe .New ("\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066e\u0072\u0065\u006ece\u0020\u006d\u0075\u0073\u0074\u0020h\u0061\u0076\u0065\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077o\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0073");};_db :=CellReference {};_ba ,_fa ,_eb :=_gf (s );if _eb !=nil {return CellReference {},_eb ;};if _ba !=""{_db .SheetName =_ba ;};if s [0]=='$'{_db .AbsoluteColumn =true ;_fa =_fa [1:];};_gg :=-1;_c :for _cf :=0;_cf < len (_fa );_cf ++{switch {case _fa [_cf ]>='0'&&_fa [_cf ]<='9'||_fa [_cf ]=='$':_gg =_cf ;break _c ;};};switch _gg {case 0:return CellReference {},_eg .Errorf ("\u006e\u006f\u0020\u006cet\u0074\u0065\u0072\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0069\u006e\u0020%\u0073",_fa );case -1:return CellReference {},_eg .Errorf ("\u006eo\u0020d\u0069\u0067\u0069\u0074\u0073\u0020\u0069\u006e\u0020\u0025\u0073",_fa );};_db .Column =_fa [0:_gg ];if _fa [_gg ]=='$'{_db .AbsoluteRow =true ;_gg ++;};_db .ColumnIdx =ColumnToIndex (_db .Column );_a ,_eb :=_d .ParseUint (_fa [_gg :],10,32);if _eb !=nil {return CellReference {},_eg .Errorf ("e\u0072\u0072\u006f\u0072 p\u0061r\u0073\u0069\u006e\u0067\u0020r\u006f\u0077\u003a\u0020\u0025\u0073",_eb );};if _a ==0{return CellReference {},_eg .Errorf ("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0072\u006f\u0077\u003a \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0030");};_db .RowIdx =uint32 (_a );return _db ,nil ;};

// Update updates reference to point one of the neighboring cells with respect to the update type after removing a row/column.
func (_cg *CellReference )Update (updateType _fg .UpdateAction )*CellReference {switch updateType {case _fg .UpdateActionRemoveColumn :_dc :=_cg ;_dc .ColumnIdx =_cg .ColumnIdx -1;_dc .Column =IndexToColumn (_dc .ColumnIdx );return _dc ;default:return _cg ;};};

// ParseColumnRangeReference splits a range reference of the form "A:B" into its
// components.
func ParseColumnRangeReference (s string )(_ebb ,_fd ColumnReference ,_bcd error ){_cee :="";_gda :=_ea .Split (s ,"\u0021");if len (_gda )==2{_cee =_gda [0];s =_gda [1];};_fc :=_ea .Split (s ,"\u003a");if len (_fc )!=2{return ColumnReference {},ColumnReference {},_fe .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");};if _cee !=""{_fc [0]=_cee +"\u0021"+_fc [0];_fc [1]=_cee +"\u0021"+_fc [1];};_fgd ,_bcd :=ParseColumnReference (_fc [0]);if _bcd !=nil {return ColumnReference {},ColumnReference {},_bcd ;};_eae ,_bcd :=ParseColumnReference (_fc [1]);if _bcd !=nil {return ColumnReference {},ColumnReference {},_bcd ;};return _fgd ,_eae ,nil ;};var _dg =_e .MustCompile ("^\u005b\u0061\u002d\u007aA-\u005a]\u0028\u005b\u0061\u002d\u007aA\u002d\u005a\u005d\u003f\u0029\u0024");
