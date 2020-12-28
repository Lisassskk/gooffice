

package reference ;import (_g "errors";_cf "fmt";_ge "gitee.com/gooffice/gooffice/spreadsheet/update";_c "regexp";_b "strconv";_d "strings";);func _abc (_fg string )(string ,string ,error ){_dae :="";_bg :=_d .LastIndex (_fg ,"\u0021");if _bg > -1{_dae =_fg [:_bg ];_fg =_fg [_bg +1:];if _dae ==""{return "","",_g .New ("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0073h\u0065\u0065t\u0020\u006e\u0061\u006d\u0065");};};return _dae ,_fg ,nil ;};

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference (s string )(_gcg ,_cg CellReference ,_bcd error ){_add ,_fc ,_bcd :=_abc (s );if _bcd !=nil {return CellReference {},CellReference {},_bcd ;};_adc :=_d .Split (_fc ,"\u003a");if len (_adc )!=2{return CellReference {},CellReference {},_g .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");};if _add !=""{_adc [0]=_add +"\u0021"+_adc [0];_adc [1]=_add +"\u0021"+_adc [1];};_dc ,_bcd :=ParseCellReference (_adc [0]);if _bcd !=nil {return CellReference {},CellReference {},_bcd ;};_ea ,_bcd :=ParseCellReference (_adc [1]);if _bcd !=nil {return CellReference {},CellReference {},_bcd ;};return _dc ,_ea ,nil ;};

// ColumnReference is a parsed reference to a column.  Input is of the form 'A',
// '$C', etc.
type ColumnReference struct{ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;SheetName string ;};

// IndexToColumn maps a column number to a column name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn (col uint32 )string {var _cag [64+1]byte ;_ba :=len (_cag );_baf :=col ;const _fdd =26;for _baf >=_fdd {_ba --;_cc :=_baf /_fdd ;_cag [_ba ]=byte ('A'+uint (_baf -_cc *_fdd ));_baf =_cc -1;};_ba --;_cag [_ba ]=byte ('A'+uint (_baf ));return string (_cag [_ba :]);};

// Update updates reference to point one of the neighboring columns with respect to the update type after removing a row/column.
func (_dba *ColumnReference )Update (updateType _ge .UpdateAction )*ColumnReference {switch updateType {case _ge .UpdateActionRemoveColumn :_ag :=_dba ;_ag .ColumnIdx =_dba .ColumnIdx -1;_ag .Column =IndexToColumn (_ag .ColumnIdx );return _ag ;default:return _dba ;};};

// CellReference is a parsed reference to a cell.  Input is of the form 'A1',
// '$C$2', etc.
type CellReference struct{RowIdx uint32 ;ColumnIdx uint32 ;Column string ;AbsoluteColumn bool ;AbsoluteRow bool ;SheetName string ;};

// ParseColumnReference parses a column reference of the form 'Sheet1!A' and splits it
// into sheet name and column segments.
func ParseColumnReference (s string )(ColumnReference ,error ){s =_d .TrimSpace (s );if len (s )< 1{return ColumnReference {},_g .New ("\u0063\u006f\u006c\u0075\u006d\u006e \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0063\u0068a\u0072a\u0063\u0074\u0065\u0072");};_db :=ColumnReference {};_dg ,_egb ,_ab :=_abc (s );if _ab !=nil {return ColumnReference {},_ab ;};if _dg !=""{_db .SheetName =_dg ;};if _egb [0]=='$'{_db .AbsoluteColumn =true ;_egb =_egb [1:];};if !_cb .MatchString (_egb ){return ColumnReference {},_g .New ("\u0063\u006f\u006c\u0075\u006dn\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0041\u0020\u0061\u006e\u0064\u0020\u005a\u005a");};_db .Column =_egb ;_db .ColumnIdx =ColumnToIndex (_db .Column );return _db ,nil ;};

// Update updates reference to point one of the neighboring cells with respect to the update type after removing a row/column.
func (_bf *CellReference )Update (updateType _ge .UpdateAction )*CellReference {switch updateType {case _ge .UpdateActionRemoveColumn :_ad :=_bf ;_ad .ColumnIdx =_bf .ColumnIdx -1;_ad .Column =IndexToColumn (_ad .ColumnIdx );return _ad ;default:return _bf ;};};

// ParseColumnRangeReference splits a range reference of the form "A:B" into its
// components.
func ParseColumnRangeReference (s string )(_be ,_bgg ColumnReference ,_bca error ){_caa :="";_cage :=_d .Split (s ,"\u0021");if len (_cage )==2{_caa =_cage [0];s =_cage [1];};_cd :=_d .Split (s ,"\u003a");if len (_cd )!=2{return ColumnReference {},ColumnReference {},_g .New ("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074");};if _caa !=""{_cd [0]=_caa +"\u0021"+_cd [0];_cd [1]=_caa +"\u0021"+_cd [1];};_dbab ,_bca :=ParseColumnReference (_cd [0]);if _bca !=nil {return ColumnReference {},ColumnReference {},_bca ;};_bb ,_bca :=ParseColumnReference (_cd [1]);if _bca !=nil {return ColumnReference {},ColumnReference {},_bca ;};return _dbab ,_bb ,nil ;};var _cb =_c .MustCompile ("^\u005b\u0061\u002d\u007aA-\u005a]\u0028\u005b\u0061\u002d\u007aA\u002d\u005a\u005d\u003f\u0029\u0024");

// String returns a string representation of ColumnReference.
func (_eb ColumnReference )String ()string {_fa :=make ([]byte ,0,4);if _eb .AbsoluteColumn {_fa =append (_fa ,'$');};_fa =append (_fa ,_eb .Column ...);return string (_fa );};

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference (s string )(CellReference ,error ){s =_d .TrimSpace (s );if len (s )< 2{return CellReference {},_g .New ("\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066e\u0072\u0065\u006ece\u0020\u006d\u0075\u0073\u0074\u0020h\u0061\u0076\u0065\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077o\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0073");};_gg :=CellReference {};_ca ,_e ,_eg :=_abc (s );if _eg !=nil {return CellReference {},_eg ;};if _ca !=""{_gg .SheetName =_ca ;};if s [0]=='$'{_gg .AbsoluteColumn =true ;_e =_e [1:];};_da :=-1;_ec :for _ecf :=0;_ecf < len (_e );_ecf ++{switch {case _e [_ecf ]>='0'&&_e [_ecf ]<='9'||_e [_ecf ]=='$':_da =_ecf ;break _ec ;};};switch _da {case 0:return CellReference {},_cf .Errorf ("\u006e\u006f\u0020\u006cet\u0074\u0065\u0072\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0069\u006e\u0020%\u0073",_e );case -1:return CellReference {},_cf .Errorf ("\u006eo\u0020d\u0069\u0067\u0069\u0074\u0073\u0020\u0069\u006e\u0020\u0025\u0073",_e );};_gg .Column =_e [0:_da ];if _e [_da ]=='$'{_gg .AbsoluteRow =true ;_da ++;};_gg .ColumnIdx =ColumnToIndex (_gg .Column );_cae ,_eg :=_b .ParseUint (_e [_da :],10,32);if _eg !=nil {return CellReference {},_cf .Errorf ("e\u0072\u0072\u006f\u0072 p\u0061r\u0073\u0069\u006e\u0067\u0020r\u006f\u0077\u003a\u0020\u0025\u0073",_eg );};if _cae ==0{return CellReference {},_cf .Errorf ("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0072\u006f\u0077\u003a \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0030");};_gg .RowIdx =uint32 (_cae );return _gg ,nil ;};

// String returns a string representation of CellReference.
func (_f CellReference )String ()string {_fd :=make ([]byte ,0,4);if _f .AbsoluteColumn {_fd =append (_fd ,'$');};_fd =append (_fd ,_f .Column ...);if _f .AbsoluteRow {_fd =append (_fd ,'$');};_fd =_b .AppendInt (_fd ,int64 (_f .RowIdx ),10);return string (_fd );};

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex (col string )uint32 {col =_d .ToUpper (col );_agb :=uint32 (0);for _ ,_bc :=range col {_agb *=26;_agb +=uint32 (_bc -'A'+1);};return _agb -1;};