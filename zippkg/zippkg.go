
package zippkg ;import (_d "archive/zip";_bc "bytes";_f "encoding/xml";_dff "fmt";_gg "gitee.com/gooffice/gooffice";_e "gitee.com/gooffice/gooffice/algo";_fg "gitee.com/gooffice/gooffice/common/tempstorage";_cb "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships";_bf "io";_c "path";_df "sort";_ge "strings";_g "time";);

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_d .File ,path string )(string ,error ){_aba ,_fcd :=_fg .TempFile (path ,"\u007a\u007a");if _fcd !=nil {return "",_fcd ;};defer _aba .Close ();_gdf ,_fcd :=f .Open ();if _fcd !=nil {return "",_fcd ;};defer _gdf .Close ();_ ,_fcd =_bf .Copy (_aba ,_gdf );if _fcd !=nil {return "",_fcd ;};return _aba .Name (),nil ;};

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_ee *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_ee ._fga =fn };func (_ab *DecodeMap )IndexFor (path string )int {return _ab ._bd [path ]};var _aga =[]byte {'/','>'};

// Decode loops decoding targets registered with AddTarget and calling th
func (_bca *DecodeMap )Decode (files []*_d .File )error {_ca :=1;for _ca > 0{for len (_bca ._ce )> 0{_ae :=_bca ._ce [0];_bca ._ce =_bca ._ce [1:];_ef :=_ae .Ifc .(*_cb .Relationships );for _ ,_dfad :=range _ef .Relationship {_ff ,_ :=_bca ._fc [_ef ];_bca ._fga (_bca ,_ff +_dfad .TargetAttr ,_dfad .TypeAttr ,files ,_dfad ,_ae );};};for _eeb ,_eeg :=range files {if _eeg ==nil {continue ;};if _ad ,_cbc :=_bca ._da [_eeg .Name ];_cbc {delete (_bca ._da ,_eeg .Name );if _bfd :=Decode (_eeg ,_ad .Ifc );_bfd !=nil {return _bfd ;};files [_eeb ]=nil ;if _adf ,_gfc :=_ad .Ifc .(*_cb .Relationships );_gfc {_bca ._ce =append (_bca ._ce ,_ad );_de ,_ :=_c .Split (_c .Clean (_eeg .Name +"\u002f\u002e\u002e\u002f"));_bca ._fc [_adf ]=_de ;_ca ++;};};};_ca --;};return nil ;};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_d .Writer ,filename string ,v interface{})error {_cff :=&_d .FileHeader {};_cff .Method =_d .Deflate ;_cff .Name =filename ;_cff .SetModTime (_g .Now ());_bb ,_ga :=z .CreateHeader (_cff );if _ga !=nil {return _dff .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_ga );};_ ,_ga =_bb .Write ([]byte (XMLHeader ));if _ga !=nil {return _dff .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_ga );};if _ga =_f .NewEncoder (SelfClosingWriter {_bb }).Encode (v );_ga !=nil {return _dff .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_ga );};_ ,_ga =_bb .Write (_ag );return _ga ;};

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_aa :=_ge .Split (path ,"\u002f");_gbb :=_ge .Join (_aa [0:len (_aa )-1],"\u002f");_fgc :=_aa [len (_aa )-1];_gbb +="\u002f_\u0072\u0065\u006c\u0073\u002f";_fgc +="\u002e\u0072\u0065l\u0073";return _gbb +_fgc ;};func MarshalXMLByTypeIndex (z *_d .Writer ,dt _gg .DocType ,typ string ,idx int ,v interface{})error {_gge :=_gg .AbsoluteFilename (dt ,typ ,idx );return MarshalXML (z ,_gge ,v );};

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _bf .Writer ;};var _ag =[]byte {'\r','\n'};func MarshalXMLByType (z *_d .Writer ,dt _gg .DocType ,typ string ,v interface{})error {_eaf :=_gg .AbsoluteFilename (dt ,typ ,0);return MarshalXML (z ,_eaf ,v );};func (_gb *DecodeMap )RecordIndex (path string ,idx int ){_gb ._bd [path ]=idx };

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_d .Writer ,zipPath string ,data []byte )error {_gbaa ,_cad :=z .Create (zipPath );if _cad !=nil {return _dff .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_cad );};_ ,_cad =_bf .Copy (_gbaa ,_bc .NewReader (data ));return _cad ;};

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
type OnNewRelationshipFunc func (_fb *DecodeMap ,_fd ,_cba string ,_dfg []*_d .File ,_ba *_cb .Relationship ,_gf Target )error ;type Target struct{Path string ;Typ string ;Ifc interface{};Index uint32 ;};

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_da map[string ]Target ;_fc map[*_cb .Relationships ]string ;_ce []Target ;_fga OnNewRelationshipFunc ;_be map[string ]struct{};_bd map[string ]int ;};const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";func (_dg SelfClosingWriter )Write (b []byte )(int ,error ){_bg :=0;_fe :=0;for _cbf :=0;_cbf < len (b )-2;_cbf ++{if b [_cbf ]=='>'&&b [_cbf +1]=='<'&&b [_cbf +2]=='/'{_fcf :=[]byte {};_aeg :=_cbf ;for _bbc :=_cbf ;_bbc >=0;_bbc --{if b [_bbc ]==' '{_aeg =_bbc ;}else if b [_bbc ]=='<'{_fcf =b [_bbc +1:_aeg ];break ;};};_cfd :=[]byte {};for _age :=_cbf +3;_age < len (b );_age ++{if b [_age ]=='>'{_cfd =b [_cbf +3:_age ];break ;};};if !_bc .Equal (_fcf ,_cfd ){continue ;};_bdd ,_bcf :=_dg .W .Write (b [_bg :_cbf ]);if _bcf !=nil {return _fe +_bdd ,_bcf ;};_fe +=_bdd ;_ ,_bcf =_dg .W .Write (_aga );if _bcf !=nil {return _fe ,_bcf ;};_fe +=3;for _efb :=_cbf +2;_efb < len (b )&&b [_efb ]!='>';_efb ++{_fe ++;_bg =_efb +2;_cbf =_bg ;};};};_aded ,_dgc :=_dg .W .Write (b [_bg :]);return _aded +_fe ,_dgc ;};

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_d .Writer ,zipPath ,storagePath string )error {_db ,_ac :=z .Create (zipPath );if _ac !=nil {return _dff .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_ac );};_gd ,_ac :=_fg .Open (storagePath );if _ac !=nil {return _dff .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_ac );};defer _gd .Close ();_ ,_ac =_bf .Copy (_db ,_gd );return _ac ;};

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_d .File ,dest interface{})error {_cd ,_ed :=f .Open ();if _ed !=nil {return _dff .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_ed );};defer _cd .Close ();_fgd :=_f .NewDecoder (_cd );if _dc :=_fgd .Decode (dest );_dc !=nil {return _dff .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_dc );};if _aae ,_eb :=dest .(*_cb .Relationships );_eb {for _aec ,_ade :=range _aae .Relationship {switch _ade .TypeAttr {case _gg .OfficeDocumentTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .OfficeDocumentType ;case _gg .StylesTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .StylesType ;case _gg .ThemeTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .ThemeType ;case _gg .ControlTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .ControlType ;case _gg .SettingsTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .SettingsType ;case _gg .ImageTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .ImageType ;case _gg .CommentsTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .CommentsType ;case _gg .ThumbnailTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .ThumbnailType ;case _gg .DrawingTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .DrawingType ;case _gg .ChartTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .ChartType ;case _gg .ExtendedPropertiesTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .ExtendedPropertiesType ;case _gg .CustomXMLTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .CustomXMLType ;case _gg .WorksheetTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .WorksheetType ;case _gg .SharedStringsTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .SharedStringsType ;case _gg .TableTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .TableType ;case _gg .HeaderTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .HeaderType ;case _gg .FooterTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .FooterType ;case _gg .NumberingTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .NumberingType ;case _gg .FontTableTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .FontTableType ;case _gg .WebSettingsTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .WebSettingsType ;case _gg .FootNotesTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .FootNotesType ;case _gg .EndNotesTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .EndNotesType ;case _gg .SlideTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .SlideType ;case _gg .VMLDrawingTypeStrict :_aae .Relationship [_aec ].TypeAttr =_gg .VMLDrawingType ;};};_df .Slice (_aae .Relationship ,func (_aea ,_aece int )bool {_gbab :=_aae .Relationship [_aea ];_ec :=_aae .Relationship [_aece ];return _e .NaturalLess (_gbab .IdAttr ,_ec .IdAttr );});};return nil ;};

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_a *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _a ._da ==nil {_a ._da =make (map[string ]Target );_a ._fc =make (map[*_cb .Relationships ]string );_a ._be =make (map[string ]struct{});_a ._bd =make (map[string ]int );};_cf :=_c .Clean (filePath );if _ ,_dfa :=_a ._be [_cf ];_dfa {return false ;};_a ._be [_cf ]=struct{}{};_a ._da [_cf ]=Target {Path :filePath ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };return true ;};
