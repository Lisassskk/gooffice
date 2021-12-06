
package common ;import (_ef "archive/zip";_ee "bytes";_fb "encoding/xml";_ddc "errors";_gef "fmt";_eca "gitee.com/gooffice/gooffice";_bdg "gitee.com/gooffice/gooffice/common/logger";_eef "gitee.com/gooffice/gooffice/common/tempstorage";_bag "gitee.com/gooffice/gooffice/common/tempstorage/diskstore";_ga "gitee.com/gooffice/gooffice/measurement";_eb "gitee.com/gooffice/gooffice/schema/soo/dml";_da "gitee.com/gooffice/gooffice/schema/soo/ofc/custom_properties";_ed "gitee.com/gooffice/gooffice/schema/soo/ofc/docPropsVTypes";_ebe "gitee.com/gooffice/gooffice/schema/soo/ofc/extended_properties";_ec "gitee.com/gooffice/gooffice/schema/soo/pkg/content_types";_ba "gitee.com/gooffice/gooffice/schema/soo/pkg/metadata/core_properties";_gb "gitee.com/gooffice/gooffice/schema/soo/pkg/relationships";_fc "gitee.com/gooffice/gooffice/zippkg";_bd "image";_ "image/gif";_ "image/jpeg";_ "image/png";_dd "os";_b "reflect";_ge "regexp";_bef "strconv";_db "strings";_be "time";);func _bcf (_ead _be .Time ,_egf string )*_eca .XSDAny {_aea :=&_eca .XSDAny {XMLName :_fb .Name {Local :_egf }};_aea .Attrs =append (_aea .Attrs ,_fb .Attr {Name :_fb .Name {Local :"\u0078\u0073\u0069\u003a\u0074\u0079\u0070\u0065"},Value :"\u0064\u0063\u0074\u0065\u0072\u006d\u0073\u003a\u00573\u0043\u0044\u0054\u0046"});_aea .Attrs =append (_aea .Attrs ,_fb .Attr {Name :_fb .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u0073i"},Value :"\u0068\u0074\u0074\u0070\u003a/\u002f\u0077\u0077\u0077\u002e\u0077\u0033\u002e\u006f\u0072\u0067\u002f\u00320\u0030\u0031\u002f\u0058\u004d\u004c\u0053\u0063\u0068\u0065\u006d\u0061\u002d\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065"});_aea .Attrs =append (_aea .Attrs ,_fb .Attr {Name :_fb .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"},Value :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"});_aea .Data =[]byte (_ead .Format (_cca ));return _aea ;};

// Title returns the Title of the document
func (_addb CoreProperties )Title ()string {if _addb ._dfe .Title !=nil {return string (_addb ._dfe .Title .Data );};return "";};

// CoreProperties contains document specific properties.
type CoreProperties struct{_dfe *_ba .CoreProperties };

// Rows returns all table rows.
func (_agbd Table )Rows ()[]*TableRow {_fbc :=_agbd ._bff .Tr ;_dbbc :=[]*TableRow {};for _ ,_fga :=range _fbc {_dbbc =append (_dbbc ,&TableRow {_edab :_fga });};return _dbbc ;};func (_gbf CustomProperties )SetPropertyAsLpstr (name string ,lpstr string ){_gfce :=_gbf .getNewProperty (name );_gfce .Lpstr =&lpstr ;_gbf .setOrReplaceProperty (_gfce );};

// Append appends DocBase part of an office document to another DocBase.
func (_ea DocBase )Append (docBase1 DocBase )DocBase {_c :=docBase1 .ContentTypes .X ();for _ ,_bb :=range _c .Default {_ea .ContentTypes .AddDefault (_bb .ExtensionAttr ,_bb .ContentTypeAttr );};for _ ,_ab :=range _c .Override {_ea .ContentTypes .AddOverride (_ab .PartNameAttr ,_ab .ContentTypeAttr );};_fd :=_ea .AppProperties .X ();_efg :=docBase1 .AppProperties .X ();if _fd .Pages !=nil {if _efg .Pages !=nil {*_fd .Pages +=*_efg .Pages ;};}else if _efg .Pages !=nil {_fd .Pages =_efg .Pages ;};if _fd .Words !=nil {if _efg .Words !=nil {*_fd .Words +=*_efg .Words ;};}else if _efg .Words !=nil {_fd .Words =_efg .Words ;};if _fd .Characters !=nil {if _efg .Characters !=nil {*_fd .Characters +=*_efg .Characters ;};}else if _efg .Characters !=nil {_fd .Characters =_efg .Characters ;};if _fd .Lines !=nil {if _efg .Lines !=nil {*_fd .Lines +=*_efg .Lines ;};}else if _efg .Lines !=nil {_fd .Lines =_efg .Lines ;};if _fd .Paragraphs !=nil {if _efg .Paragraphs !=nil {*_fd .Paragraphs +=*_efg .Paragraphs ;};}else if _efg .Paragraphs !=nil {_fd .Paragraphs =_efg .Paragraphs ;};if _fd .Notes !=nil {if _efg .Notes !=nil {*_fd .Notes +=*_efg .Notes ;};}else if _efg .Notes !=nil {_fd .Notes =_efg .Notes ;};if _fd .HiddenSlides !=nil {if _efg .HiddenSlides !=nil {*_fd .HiddenSlides +=*_efg .HiddenSlides ;};}else if _efg .HiddenSlides !=nil {_fd .HiddenSlides =_efg .HiddenSlides ;};if _fd .MMClips !=nil {if _efg .MMClips !=nil {*_fd .MMClips +=*_efg .MMClips ;};}else if _efg .MMClips !=nil {_fd .MMClips =_efg .MMClips ;};if _fd .LinksUpToDate !=nil {if _efg .LinksUpToDate !=nil {*_fd .LinksUpToDate =*_fd .LinksUpToDate &&*_efg .LinksUpToDate ;};}else if _efg .LinksUpToDate !=nil {_fd .LinksUpToDate =_efg .LinksUpToDate ;};if _fd .CharactersWithSpaces !=nil {if _efg .CharactersWithSpaces !=nil {*_fd .CharactersWithSpaces +=*_efg .CharactersWithSpaces ;};}else if _efg .CharactersWithSpaces !=nil {_fd .CharactersWithSpaces =_efg .CharactersWithSpaces ;};if _fd .SharedDoc !=nil {if _efg .SharedDoc !=nil {*_fd .SharedDoc =*_fd .SharedDoc ||*_efg .SharedDoc ;};}else if _efg .SharedDoc !=nil {_fd .SharedDoc =_efg .SharedDoc ;};if _fd .HyperlinksChanged !=nil {if _efg .HyperlinksChanged !=nil {*_fd .HyperlinksChanged =*_fd .HyperlinksChanged ||*_efg .HyperlinksChanged ;};}else if _efg .HyperlinksChanged !=nil {_fd .HyperlinksChanged =_efg .HyperlinksChanged ;};_fd .DigSig =nil ;if _fd .TitlesOfParts ==nil &&_efg .TitlesOfParts !=nil {_fd .TitlesOfParts =_efg .TitlesOfParts ;};if _fd .HeadingPairs !=nil {if _efg .HeadingPairs !=nil {_fg :=_fd .HeadingPairs .Vector ;_eaf :=_efg .HeadingPairs .Vector ;_bf :=_fg .Variant ;_dg :=_eaf .Variant ;_ad :=[]*_ed .Variant {};for _ce :=0;_ce < len (_dg );_ce +=2{_ca :=_dg [_ce ].Lpstr ;_cf :=false ;for _dbg :=0;_dbg < len (_bf );_dbg +=2{_add :=_bf [_dbg ].Lpstr ;if _add !=nil &&_ca !=nil &&*_add ==*_ca {*_bf [_dbg +1].I4 =*_bf [_dbg +1].I4 +*_dg [_ce +1].I4 ;_cf =true ;break ;};};if !_cf {_ad =append (_ad ,&_ed .Variant {CT_Variant :_ed .CT_Variant {Lpstr :_dg [_ce ].Lpstr }});_ad =append (_ad ,&_ed .Variant {CT_Variant :_ed .CT_Variant {I4 :_dg [_ce ].I4 }});};};_bf =append (_bf ,_ad ...);_fg .SizeAttr =uint32 (len (_bf ));};}else if _efg .HeadingPairs !=nil {_fd .HeadingPairs =_efg .HeadingPairs ;};if _fd .HLinks !=nil {if _efg .HLinks !=nil {_df :=_fd .HLinks .Vector ;_bg :=_efg .HLinks .Vector ;_cg :=_df .Variant ;_fda :=_bg .Variant ;for _ ,_gf :=range _fda {_gad :=true ;for _ ,_eeb :=range _cg {if _b .DeepEqual (_eeb ,_gf ){_gad =false ;break ;};};if _gad {_cg =append (_cg ,_gf );_df .SizeAttr ++;};};};}else if _efg .HLinks !=nil {_fd .HLinks =_efg .HLinks ;};_cb :=_ea .GetOrCreateCustomProperties ();_ff :=docBase1 .GetOrCreateCustomProperties ();for _ ,_bgd :=range _ff .PropertiesList (){_cb .setProperty (_bgd );};_ea .CustomProperties =_cb ;_dgg :=_ea .Rels .X ().Relationship ;for _ ,_gae :=range docBase1 .Rels .X ().Relationship {_ffe :=true ;for _ ,_eg :=range _dgg {if _eg .TargetAttr ==_gae .TargetAttr &&_eg .TypeAttr ==_gae .TypeAttr {_ffe =false ;break ;};};if _ffe {_ea .Rels .AddRelationship (_gae .TargetAttr ,_gae .TypeAttr );};};for _ ,_gd :=range docBase1 .ExtraFiles {_fe :=_gd .ZipPath ;_dgf :=true ;for _ ,_aa :=range _ea .ExtraFiles {if _aa .ZipPath ==_fe {_dgf =false ;break ;};};if _dgf {_ea .ExtraFiles =append (_ea .ExtraFiles ,_gd );};};return _ea ;};func (_ceg CustomProperties )setProperty (_eeg *_da .CT_Property ){_ceg .setPropertyHelper (_eeg ,false )};

// SetLanguage records the language of the document.
func (_gba CoreProperties )SetLanguage (s string ){_gba ._dfe .Language =&_eca .XSDAny {XMLName :_fb .Name {Local :"d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}};_gba ._dfe .Language .Data =[]byte (s );};

// RemoveOverrideByIndex removes an override given a path and override index.
func (_ae ContentTypes )RemoveOverrideByIndex (path string ,indexToFind int )error {_cge :=path [0:len (path )-5];if !_db .HasPrefix (_cge ,"\u002f"){_cge ="\u002f"+_cge ;};_gcb ,_cea :=_ge .Compile (_cge +"\u0028\u005b\u0030-\u0039\u005d\u002b\u0029\u002e\u0078\u006d\u006c");if _cea !=nil {return _cea ;};_bac :=0;_fee :=-1;for _bae ,_edae :=range _ae ._beg .Override {if _eebc :=_gcb .FindStringSubmatch (_edae .PartNameAttr );len (_eebc )> 1{if _bac ==indexToFind {_fee =_bae ;}else if _bac > indexToFind {_bbc ,_ :=_bef .Atoi (_eebc [1]);_bbc --;_edae .PartNameAttr =_gef .Sprintf ("\u0025\u0073\u0025\u0064\u002e\u0078\u006d\u006c",_cge ,_bbc );};_bac ++;};};if _fee > -1{copy (_ae ._beg .Override [_fee :],_ae ._beg .Override [_fee +1:]);_ae ._beg .Override =_ae ._beg .Override [0:len (_ae ._beg .Override )-1];};return nil ;};

// LastModifiedBy returns the name of the last person to modify the document
func (_ecf CoreProperties )LastModifiedBy ()string {if _ecf ._dfe .LastModifiedBy !=nil {return *_ecf ._dfe .LastModifiedBy ;};return "";};

// SetDocSecurity sets the document security flag.
func (_gefb AppProperties )SetDocSecurity (v int32 ){_gefb ._ade .DocSecurity =_eca .Int32 (v )};

// ContentTypes is the top level "[Content_Types].xml" in a zip package.
type ContentTypes struct{_beg *_ec .Types };

// Properties returns table properties.
func (_cga Table )Properties ()*_eb .CT_TableProperties {return _cga ._bff .TblPr };

// PropertiesList returns the list of all custom properties of the document.
func (_gee CustomProperties )PropertiesList ()[]*_da .CT_Property {return _gee ._dag .Property };

// SetCategory records the category of the document.
func (_eaa CoreProperties )SetCategory (s string ){_eaa ._dfe .Category =&s };

// Author returns the author of the document
func (_edg CoreProperties )Author ()string {if _edg ._dfe .Creator !=nil {return string (_edg ._dfe .Creator .Data );};return "";};

// SetTitle records the title of the document.
func (_dca CoreProperties )SetTitle (s string ){if _dca ._dfe .Title ==nil {_dca ._dfe .Title =&_eca .XSDAny {XMLName :_fb .Name {Local :"\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}};};_dca ._dfe .Title .Data =[]byte (s );};

// AddCustomRelationships adds relationships related to custom properties to the document.
func (_ddbg *DocBase )AddCustomRelationships (){_ddbg .ContentTypes .AddOverride ("/\u0064o\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0063u\u0073\u0074\u006f\u006d.x\u006d\u006c","\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065n\u0078\u006d\u006c\u0066\u006fr\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0063\u0075\u0073\u0074\u006f\u006d\u002d\u0070r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073+\u0078\u006d\u006c");_ddbg .Rels .AddRelationship ("\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c",_eca .CustomPropertiesType );};

// NewCustomProperties constructs a new CustomProperties.
func NewCustomProperties ()CustomProperties {return CustomProperties {_dag :_da .NewProperties ()}};

// SetHeight sets row height, see measurement package.
func (_cdce TableRow )SetHeight (m _ga .Distance ){_aced :=_ga .ToEMU (float64 (m ));_cdce ._edab .HAttr =_eb .ST_Coordinate {ST_CoordinateUnqualified :&_aced };};

// AddExtraFileFromZip is used when reading an unsupported file from an OOXML
// file. This ensures that unsupported file content will at least round-trip
// correctly.
func (_gbeb *DocBase )AddExtraFileFromZip (f *_ef .File )error {_efc ,_bfg :=_fc .ExtractToDiskTmp (f ,_gbeb .TmpPath );if _bfg !=nil {return _gef .Errorf ("\u0065\u0072r\u006f\u0072\u0020\u0065x\u0074\u0072a\u0063\u0074\u0069\u006e\u0067\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073",_bfg );};_gbeb .ExtraFiles =append (_gbeb .ExtraFiles ,ExtraFile {ZipPath :f .Name ,DiskPath :_efc });return nil ;};

// Path returns the path to an image file, if any.
func (_efe ImageRef )Path ()string {return _efe ._cgg .Path };

// MakeImageRef constructs an image reference which is a reference to a
// particular image file inside a document.  The same image can be used multiple
// times in a document by re-use the ImageRef.
func MakeImageRef (img Image ,d *DocBase ,rels Relationships )ImageRef {return ImageRef {_cgg :img ,_bfbf :d ,_gbfe :rels };};func (_cff CustomProperties )SetPropertyAsUint (name string ,ui uint ){_ddg :=_cff .getNewProperty (name );_gfc :=uint32 (ui );_ddg .Uint =&_gfc ;_cff .setOrReplaceProperty (_ddg );};

// GetPropertyByName returns a custom property selected by it's name.
func (_bfb CustomProperties )GetPropertyByName (name string )CustomProperty {_fgg :=_bfb ._dag .Property ;for _ ,_fccd :=range _fgg {if *_fccd .NameAttr ==name {return CustomProperty {_eag :_fccd };};};return CustomProperty {};};

// Relationship is a relationship within a .rels file.
type Relationship struct{_gaac *_gb .Relationship };

// SetCompany sets the name of the company that created the document.
func (_cgb AppProperties )SetCompany (s string ){_cgb ._ade .Company =&s };

// NewTableWithXfrm makes a new table with a pointer to its parent Xfrm for changing its offset and size.
func NewTableWithXfrm (xfrm *_eb .CT_Transform2D )*Table {_bece :=_eb .NewTbl ();_bece .TblPr =_eb .NewCT_TableProperties ();return &Table {_bff :_bece ,_ebd :xfrm };};func (_bgg CustomProperties )setOrReplaceProperty (_cda *_da .CT_Property ){_bgg .setPropertyHelper (_cda ,true );};func (_cac CustomProperties )SetPropertyAsInt (name string ,i int ){_gfg :=_cac .getNewProperty (name );_ceb :=int32 (i );_gfg .Int =&_ceb ;_cac .setOrReplaceProperty (_gfg );};

// RelID returns the relationship ID.
func (_befb ImageRef )RelID ()string {return _befb ._edcg };

// SetAuthor records the author of the document.
func (_gca CoreProperties )SetAuthor (s string ){if _gca ._dfe .Creator ==nil {_gca ._dfe .Creator =&_eca .XSDAny {XMLName :_fb .Name {Local :"\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}};};_gca ._dfe .Creator .Data =[]byte (s );};const _ddbgg =2021;func (_edaeg CustomProperties )getNewProperty (_gec string )*_da .CT_Property {_gfa :=_edaeg ._dag .Property ;_cag :=int32 (1);for _ ,_gebe :=range _gfa {if _gebe .PidAttr > _cag {_cag =_gebe .PidAttr ;};};_de :=_da .NewCT_Property ();_de .NameAttr =&_gec ;_de .PidAttr =_cag +1;_de .FmtidAttr ="\u007b\u0044\u0035\u0043\u0044\u0044\u0035\u0030\u0035\u002d\u0032\u0045\u0039\u0043\u002d\u0031\u0030\u0031\u0042\u002d\u0039\u0033\u0039\u0037-\u0030\u0038\u0030\u0030\u0032B\u0032\u0043F\u0039\u0041\u0045\u007d";return _de ;};func (_fbd CustomProperties )SetPropertyAsArray (name string ,array *_ed .Array ){_dcg :=_fbd .getNewProperty (name );_dcg .Array =array ;_fbd .setOrReplaceProperty (_dcg );};

// X returns the inner wrapped XML type.
func (_baa CustomProperties )X ()*_da .Properties {return _baa ._dag };

// Table represents a table in the document.
type Table struct{_bff *_eb .Tbl ;_ebd *_eb .CT_Transform2D ;};

// AddDefault registers a default content type for a given file extension.
func (_ecgg ContentTypes )AddDefault (fileExtension string ,contentType string ){fileExtension =_db .ToLower (fileExtension );for _ ,_geb :=range _ecgg ._beg .Default {if _geb .ExtensionAttr ==fileExtension &&_geb .ContentTypeAttr ==contentType {return ;};};_ag :=_ec .NewDefault ();_ag .ExtensionAttr =fileExtension ;_ag .ContentTypeAttr =contentType ;_ecgg ._beg .Default =append (_ecgg ._beg .Default ,_ag );};

// SetApplication sets the name of the application that created the document.
func (_ege AppProperties )SetApplication (s string ){_ege ._ade .Application =&s };

// NewContentTypes returns a wrapper around a newly constructed content-types.
func NewContentTypes ()ContentTypes {_gdg :=ContentTypes {_beg :_ec .NewTypes ()};_gdg .AddDefault ("\u0078\u006d\u006c","\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c");_gdg .AddDefault ("\u0072\u0065\u006c\u0073","\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006fn\u002f\u0076\u006ed\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006fr\u006d\u0061\u0074\u0073\u002dpa\u0063\u006b\u0061\u0067\u0065\u002e\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002b\u0078\u006d\u006c");_gdg .AddDefault ("\u0070\u006e\u0067","\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg");_gdg .AddDefault ("\u006a\u0070\u0065\u0067","\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067");_gdg .AddDefault ("\u006a\u0070\u0067","\u0069m\u0061\u0067\u0065\u002f\u006a\u0070g");_gdg .AddDefault ("\u0077\u006d\u0066","i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066");_gdg .AddOverride ("\u002fd\u006fc\u0050\u0072\u006f\u0070\u0073/\u0063\u006fr\u0065\u002e\u0078\u006d\u006c","\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073-\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002e\u0063\u006f\u0072\u0065\u002dp\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073\u002bx\u006d\u006c");_gdg .AddOverride ("\u002f\u0064\u006f\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0061\u0070p\u002e\u0078\u006d\u006c","a\u0070\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u002e\u0065\u0078\u0074\u0065\u006e\u0064\u0065\u0064\u002dp\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002b\u0078m\u006c");return _gdg ;};

// NewRelationship constructs a new relationship.
func NewRelationship ()Relationship {return Relationship {_gaac :_gb .NewRelationship ()}};

// GetByRelId returns a relationship with the associated relation ID.
func (_cdcc Relationships )GetByRelId (idAttr string )Relationship {for _ ,_dcce :=range _cdcc ._cdba .Relationship {if _dcce .IdAttr ==idAttr {return Relationship {_gaac :_dcce };};};return Relationship {};};func (_afe CustomProperties )SetPropertyAsBool (name string ,b bool ){_gdb :=_afe .getNewProperty (name );_gdb .Bool =&b ;_afe .setOrReplaceProperty (_gdb );};

// AddAutoRelationship adds a relationship with an automatically generated
// filename based off of the type. It should be preferred over AddRelationship
// to ensure consistent filenames are maintained.
func (_fcbga Relationships )AddAutoRelationship (dt _eca .DocType ,src string ,idx int ,ctype string )Relationship {return _fcbga .AddRelationship (_eca .RelativeFilename (dt ,src ,ctype ,idx ),ctype );};func (_aceb CustomProperties )SetPropertyAsStorage (name string ,storage string ){_cdf :=_aceb .getNewProperty (name );_cdf .Storage =&storage ;_aceb .setOrReplaceProperty (_cdf );};func (_ggg CustomProperties )setPropertyHelper (_cdcd *_da .CT_Property ,_ccf bool ){_gbe :=_ggg .GetPropertyByName (*_cdcd .NameAttr );if (_gbe ==CustomProperty {}){_ggg ._dag .Property =append (_ggg ._dag .Property ,_cdcd );}else if _ccf {_cdcd .FmtidAttr =_gbe ._eag .FmtidAttr ;if _gbe ._eag .PidAttr ==0{_cdcd .PidAttr =_gbe ._eag .PidAttr ;};_cdcd .LinkTargetAttr =_gbe ._eag .LinkTargetAttr ;*_gbe ._eag =*_cdcd ;};};

// GetTargetByRelId returns a target path with the associated relation ID.
func (_ceaf Relationships )GetTargetByRelId (idAttr string )string {for _ ,_gdgg :=range _ceaf ._cdba .Relationship {if _gdgg .IdAttr ==idAttr {return _gdgg .TargetAttr ;};};return "";};func (_dbb CustomProperties )SetPropertyAsI1 (name string ,i1 int8 ){_afa :=_dbb .getNewProperty (name );_afa .I1 =&i1 ;_dbb .setOrReplaceProperty (_afa );};

// Size returns the size of an image
func (_bbf ImageRef )Size ()_bd .Point {return _bbf ._cgg .Size };const _ccg ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";

// AddCol adds a column to a table.
func (_afca Table )AddCol ()*TableCol {_dce :=_eb .NewCT_TableCol ();_afca ._bff .TblGrid .GridCol =append (_afca ._bff .TblGrid .GridCol ,_dce );for _ ,_gcf :=range _afca ._bff .Tr {_edaf :=_eb .NewCT_TableCell ();_gcf .Tc =append (_gcf .Tc ,_edaf );};return &TableCol {_bggc :_dce };};func (_ebee CustomProperties )SetPropertyAsDecimal (name string ,decimal float64 ){_gaba :=_ebee .getNewProperty (name );_gaba .Decimal =&decimal ;_ebee .setOrReplaceProperty (_gaba );};

// EnsureDefault esnures that an extension and default content type exist,
// adding it if necessary.
func (_bdb ContentTypes )EnsureDefault (ext ,contentType string ){ext =_db .ToLower (ext );for _ ,_dgb :=range _bdb ._beg .Default {if _dgb .ExtensionAttr ==ext {_dgb .ContentTypeAttr =contentType ;return ;};};_gag :=&_ec .Default {};_gag .ContentTypeAttr =contentType ;_gag .ExtensionAttr =ext ;_bdb ._beg .Default =append (_bdb ._beg .Default ,_gag );};

// NewAppProperties constructs a new AppProperties.
func NewAppProperties ()AppProperties {_caf :=AppProperties {_ade :_ebe .NewProperties ()};_caf .SetCompany ("\u0046\u006f\u0078\u0079\u0055\u0074\u0069\u006c\u0073\u0020\u0065\u0068\u0066");_caf .SetApplication ("g\u0069\u0074\u0068\u0075\u0062\u002ec\u006f\u006d\u002f\u0075\u006e\u0069\u0064\u006f\u0063/\u0075\u006e\u0069o\u0066f\u0069\u0063\u0065");_caf .SetDocSecurity (0);_caf .SetLinksUpToDate (false );var _cc ,_aab ,_fge int64 ;_gef .Sscanf (Version ,"\u0025\u0064\u002e\u0025\u0064\u002e\u0025\u0064",&_cc ,&_aab ,&_fge );_gea :=float64 (_cc )+float64 (_aab )/10000.0;_caf .SetApplicationVersion (_gef .Sprintf ("\u0025\u0030\u0037\u002e\u0034\u0066",_gea ));return _caf ;};

// X returns the inner raw content types.
func (_bab ContentTypes )X ()*_ec .Types {return _bab ._beg };

// DocBase is the type embedded in in the Document/Workbook/Presentation types
// that contains members common to all.
type DocBase struct{ContentTypes ContentTypes ;AppProperties AppProperties ;Rels Relationships ;CoreProperties CoreProperties ;CustomProperties CustomProperties ;Thumbnail _bd .Image ;Images []ImageRef ;ExtraFiles []ExtraFile ;TmpPath string ;};func (_ece CustomProperties )SetPropertyAsOblob (name ,oblob string ){_cdb :=_ece .getNewProperty (name );_cdb .Oblob =&oblob ;_ece .setOrReplaceProperty (_cdb );};

// NewRelationships creates a new relationship wrapper.
func NewRelationships ()Relationships {return Relationships {_cdba :_gb .NewRelationships ()}};func (_cee CustomProperties )SetPropertyAsError (name string ,error string ){_cae :=_cee .getNewProperty (name );_cae .Error =&error ;_cee .setOrReplaceProperty (_cae );};func (_aeac CustomProperties )SetPropertyAsClsid (name string ,clsid string ){_befg :=_aeac .getNewProperty (name );_befg .Clsid =&clsid ;_aeac .setOrReplaceProperty (_befg );};func (_dcf CustomProperties )SetPropertyAsFiletime (name string ,filetime _be .Time ){_dfg :=_dcf .getNewProperty (name );_dfg .Filetime =&filetime ;_dcf .setOrReplaceProperty (_dfg );};

// Category returns the category of the document
func (_ede CoreProperties )Category ()string {if _ede ._dfe .Category !=nil {return *_ede ._dfe .Category ;};return "";};

// CreateCustomProperties creates the custom properties of the document.
func (_edda *DocBase )CreateCustomProperties (){_edda .CustomProperties =NewCustomProperties ();_edda .AddCustomRelationships ();};

// SetWidth sets column width, see measurement package.
func (_becd TableCol )SetWidth (m _ga .Distance ){_edada :=_ga .ToEMU (float64 (m ));_becd ._bggc .WAttr =_eb .ST_Coordinate {ST_CoordinateUnqualified :&_edada };};

// X returns the inner wrapped XML type.
func (_cfg Theme )X ()*_eb .Theme {return _cfg ._bdbf };

// AddOverride adds an override content type for a given path name.
func (_fag ContentTypes )AddOverride (path ,contentType string ){if !_db .HasPrefix (path ,"\u002f"){path ="\u002f"+path ;};if _db .HasPrefix (contentType ,"\u0068\u0074\u0074\u0070"){_bdg .Log .Debug ("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070",contentType );};for _ ,_cgbf :=range _fag ._beg .Override {if _cgbf .PartNameAttr ==path &&_cgbf .ContentTypeAttr ==contentType {return ;};};_fae :=_ec .NewOverride ();_fae .PartNameAttr =path ;_fae .ContentTypeAttr =contentType ;_fag ._beg .Override =append (_fag ._beg .Override ,_fae );};func (_dfd CustomProperties )SetPropertyAsCy (name string ,cy string ){_faf :=_dfd .getNewProperty (name );_faf .Cy =&cy ;_dfd .setOrReplaceProperty (_faf );};

// SetApplicationVersion sets the version of the application that created the
// document.  Per MS, the verison string mut be in the form 'XX.YYYY'.
func (_cef AppProperties )SetApplicationVersion (s string ){_cef ._ade .AppVersion =&s };

// ApplicationVersion returns the version of the application that created the
// document.
func (_fcb AppProperties )ApplicationVersion ()string {if _fcb ._ade .AppVersion !=nil {return *_fcb ._ade .AppVersion ;};return "";};const _cca ="2\u00300\u0036\u002d\u0030\u0031\u002d\u0030\u0032\u00541\u0035\u003a\u0030\u0034:0\u0035\u005a";

// SetStyle assigns TableStyle to a table.
func (_afg Table )SetStyle (style *_eb .CT_TableStyle ){if _afg ._bff .TblPr ==nil {_afg ._bff .TblPr =_eb .NewCT_TableProperties ();};if _afg ._bff .TblPr .Choice ==nil {_afg ._bff .TblPr .Choice =_eb .NewCT_TablePropertiesChoice ();};_afg ._bff .TblPr .Choice .TableStyle =style ;};

// Target returns the target attrubute of the image reference (a path where the image file is located in the document structure).
func (_dfgb *ImageRef )Target ()string {return _dfgb ._fcbg };func (_ddf CustomProperties )SetPropertyAsOstream (name string ,ostream string ){_gdf :=_ddf .getNewProperty (name );_gdf .Ostream =&ostream ;_ddf .setOrReplaceProperty (_gdf );};

// SetID set the ID of a relationship.
func (_eff Relationship )SetID (ID string ){_eff ._gaac .IdAttr =ID ;};func (_aeg CustomProperties )SetPropertyAsI2 (name string ,i2 int16 ){_acc :=_aeg .getNewProperty (name );_acc .I2 =&i2 ;_aeg .setOrReplaceProperty (_acc );};func (_dbd CustomProperties )SetPropertyAsUi8 (name string ,ui8 uint64 ){_fca :=_dbd .getNewProperty (name );_fca .Ui8 =&ui8 ;_dbd .setOrReplaceProperty (_fca );};func (_ggca *ImageRef )SetRelID (id string ){_ggca ._edcg =id };func (_afc CustomProperties )SetPropertyAsI8 (name string ,i8 int64 ){_abb :=_afc .getNewProperty (name );_abb .I8 =&i8 ;_afc .setOrReplaceProperty (_abb );};

// ImageFromFile reads an image from a file on disk. It doesn't keep the image
// in memory and only reads it to determine the format and size. You can also
// construct an Image directly if the file and size are known.
// NOTE: See also ImageFromStorage.
func ImageFromFile (path string )(Image ,error ){_edad ,_fde :=_dd .Open (path );_edce :=Image {};if _fde !=nil {return _edce ,_gef .Errorf ("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_fde );};defer _edad .Close ();_baec ,_cfe ,_fde :=_bd .Decode (_edad );if _fde !=nil {return _edce ,_gef .Errorf ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s",_fde );};_edce .Path =path ;_edce .Format =_cfe ;_edce .Size =_baec .Bounds ().Size ();return _edce ,nil ;};func UtcTimeFormat (t _be .Time )string {return t .Format (_ccg )+"\u0020\u0055\u0054\u0043"};

// SetTarget set the target (path) of a relationship.
func (_bge Relationship )SetTarget (s string ){_bge ._gaac .TargetAttr =s };func (_cbd CustomProperties )SetPropertyAsUi2 (name string ,ui2 uint16 ){_egg :=_cbd .getNewProperty (name );_egg .Ui2 =&ui2 ;_cbd .setOrReplaceProperty (_egg );};

// RelativeWidth returns the relative width of an image given a fixed height.
// This is used when setting image to a fixed height to calculate the width
// required to keep the same image aspect ratio.
func (_fce ImageRef )RelativeWidth (h _ga .Distance )_ga .Distance {_gbab :=float64 (_fce .Size ().X )/float64 (_fce .Size ().Y );return h *_ga .Distance (_gbab );};

// AddImageToZip adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImageToZip (z *_ef .Writer ,img ImageRef ,imageNum int ,dt _eca .DocType )error {_gadg :=_eca .AbsoluteImageFilename (dt ,imageNum ,_db .ToLower (img .Format ()));if img .Data ()!=nil &&len (*img .Data ())> 0{if _dcaf :=_fc .AddFileFromBytes (z ,_gadg ,*img .Data ());_dcaf !=nil {return _dcaf ;};}else if img .Path ()!=""{if _efdd :=_fc .AddFileFromDisk (z ,_gadg ,img .Path ());_efdd !=nil {return _efdd ;};}else {return _gef .Errorf ("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u003a %\u002b\u0076",img );};return nil ;};var _gfgd =_ddc .New ("\u0069\u006d\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0074\u006fr\u0061\u0067\u0065");

// NewTable makes a new table.
func NewTable ()*Table {_dbf :=_eb .NewTbl ();_dbf .TblPr =_eb .NewCT_TableProperties ();return &Table {_bff :_dbf };};

// ImageFromStorage reads an image using the currently set
// temporary storage mechanism (see tempstorage). You can also
// construct an Image directly if the file and size are known.
func ImageFromStorage (path string )(Image ,error ){_dccg :=Image {};_agd ,_egc :=_eef .Open (path );if _egc !=nil {return _dccg ,_gef .Errorf ("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_egc );};defer _agd .Close ();_bec ,_fgb ,_egc :=_bd .Decode (_agd );if _egc !=nil {return _dccg ,_gef .Errorf ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s",_egc );};_dccg .Path =path ;_dccg .Format =_fgb ;_dccg .Size =_bec .Bounds ().Size ();return _dccg ,nil ;};

// SetTarget changes the target attribute of the image reference (e.g. in the case of the creation of the reference or if the image which the reference is related to was moved from one location to another).
func (_fcea *ImageRef )SetTarget (target string ){_fcea ._fcbg =target };

// SetLinksUpToDate sets the links up to date flag.
func (_ecg AppProperties )SetLinksUpToDate (v bool ){_ecg ._ade .LinksUpToDate =_eca .Bool (v )};

// Description returns the description of the document
func (_cgef CoreProperties )Description ()string {if _cgef ._dfe .Description !=nil {return string (_cgef ._dfe .Description .Data );};return "";};

// SetContentStatus records the content status of the document.
func (_ecae CoreProperties )SetContentStatus (s string ){_ecae ._dfe .ContentStatus =&s };func (_bdc CustomProperties )SetPropertyAsI4 (name string ,i4 int32 ){_bfc :=_bdc .getNewProperty (name );_bfc .I4 =&i4 ;_bdc .setOrReplaceProperty (_bfc );};var ReleasedAt =_be .Date (_ddbgg ,_dbda ,_ccgf ,_fged ,_aeaa ,0,0,_be .UTC );

// Relationships returns a slice of all of the relationships.
func (_cfb Relationships )Relationships ()[]Relationship {_bdde :=[]Relationship {};for _ ,_dfgc :=range _cfb ._cdba .Relationship {_bdde =append (_bdde ,Relationship {_gaac :_dfgc });};return _bdde ;};const Version ="\u0031\u002e\u0031\u0036\u002e\u0030";

// Image is a container for image information. It's used as we need format and
// and size information to use images.
// It contains either the filesystem path to the image, or the image itself.
type Image struct{Size _bd .Point ;Format string ;Path string ;Data *[]byte ;};

// Relationships represents a .rels file.
type Relationships struct{_cdba *_gb .Relationships };func (_dccd CustomProperties )SetPropertyAsNull (name string ){_gde :=_dccd .getNewProperty (name );_gde .Null =_ed .NewNull ();_dccd .setOrReplaceProperty (_gde );};

// CustomProperty contains document specific property.
// Using of this type is deprecated.
type CustomProperty struct{_eag *_da .CT_Property };func (_fcf CustomProperties )SetPropertyAsOstorage (name string ,ostorage string ){_cdd :=_fcf .getNewProperty (name );_cdd .Ostorage =&ostorage ;_fcf .setOrReplaceProperty (_cdd );};

// RelativeHeight returns the relative height of an image given a fixed width.
// This is used when setting image to a fixed width to calculate the height
// required to keep the same image aspect ratio.
func (_eba ImageRef )RelativeHeight (w _ga .Distance )_ga .Distance {_fdc :=float64 (_eba .Size ().Y )/float64 (_eba .Size ().X );return w *_ga .Distance (_fdc );};

// TableCol represents a column in a table.
type TableCol struct{_bggc *_eb .CT_TableCol };

// CopyRelationship copies the relationship.
func (_aebe Relationships )CopyRelationship (idAttr string )(Relationship ,bool ){for _adee :=range _aebe ._cdba .Relationship {if _aebe ._cdba .Relationship [_adee ].IdAttr ==idAttr {_fagg :=*_aebe ._cdba .Relationship [_adee ];_fdfe :=len (_aebe ._cdba .Relationship )+1;_gbef :=map[string ]struct{}{};for _ ,_fggd :=range _aebe ._cdba .Relationship {_gbef [_fggd .IdAttr ]=struct{}{};};for _ ,_gcg :=_gbef [_gef .Sprintf ("\u0072\u0049\u0064%\u0064",_fdfe )];_gcg ;_ ,_gcg =_gbef [_gef .Sprintf ("\u0072\u0049\u0064%\u0064",_fdfe )]{_fdfe ++;};_fagg .IdAttr =_gef .Sprintf ("\u0072\u0049\u0064%\u0064",_fdfe );_aebe ._cdba .Relationship =append (_aebe ._cdba .Relationship ,&_fagg );return Relationship {_gaac :&_fagg },true ;};};return Relationship {},false ;};const _dbda =11;

// NewRelationshipsCopy creates a new relationships wrapper as a copy of passed in instance.
func NewRelationshipsCopy (rels Relationships )Relationships {_bbb :=*rels ._cdba ;return Relationships {_cdba :&_bbb };};func (_ddb CustomProperties )SetPropertyAsUi1 (name string ,ui1 uint8 ){_ffc :=_ddb .getNewProperty (name );_ffc .Ui1 =&ui1 ;_ddb .setOrReplaceProperty (_ffc );};func (_cec CustomProperties )SetPropertyAsStream (name string ,stream string ){_aff :=_cec .getNewProperty (name );_aff .Stream =&stream ;_cec .setOrReplaceProperty (_aff );};func (_cbe CustomProperties )SetPropertyAsBstr (name string ,bstr string ){_geea :=_cbe .getNewProperty (name );_geea .Bstr =&bstr ;_cbe .setOrReplaceProperty (_geea );};

// Cells returns an array of row cells.
func (_fff TableRow )Cells ()[]*_eb .CT_TableCell {return _fff ._edab .Tc };

// ExtraFile is an unsupported file type extracted from, or to be written to a
// zip package
type ExtraFile struct{ZipPath string ;DiskPath string ;};

// SetDescription records the description of the document.
func (_bce CoreProperties )SetDescription (s string ){if _bce ._dfe .Description ==nil {_bce ._dfe .Description =&_eca .XSDAny {XMLName :_fb .Name {Local :"\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}};};_bce ._dfe .Description .Data =[]byte (s );};

// X returns the inner wrapped XML type.
func (_gebf Relationship )X ()*_gb .Relationship {return _gebf ._gaac };

// X returns the underlying raw XML data.
func (_cdfb Relationships )X ()*_gb .Relationships {return _cdfb ._cdba };

// SetOffsetY sets vertical offset of a table in distance units (see measurement package).
func (_ffed Table )SetOffsetY (offY float64 ){if _ffed ._ebd .Off ==nil {_ffed ._ebd .Off =_eb .NewCT_Point2D ();_bdbg :=int64 (0);_ffed ._ebd .Off .XAttr =_eb .ST_Coordinate {ST_CoordinateUnqualified :&_bdbg };};_aca :=_ga .ToEMU (offY );_ffed ._ebd .Off .YAttr =_eb .ST_Coordinate {ST_CoordinateUnqualified :&_aca };};

// AddRow adds a row to a table.
func (_dcbb Table )AddRow ()*TableRow {_adeg :=_eb .NewCT_TableRow ();for _gfd :=0;_gfd < len (_dcbb ._bff .TblGrid .GridCol );_gfd ++{_adeg .Tc =append (_adeg .Tc ,_eb .NewCT_TableCell ());};_dcbb ._bff .Tr =append (_dcbb ._bff .Tr ,_adeg );return &TableRow {_edab :_adeg };};

// TableStyles contains document specific properties.
type TableStyles struct{_fbf *_eb .TblStyleLst };func (_edb CustomProperties )SetPropertyAsR4 (name string ,r4 float32 ){_eaaa :=_edb .getNewProperty (name );_eaaa .R4 =&r4 ;_edb .setOrReplaceProperty (_eaaa );};

// CopyOverride copies override content type for a given `path` and puts it with a path `newPath`.
func (_fdb ContentTypes )CopyOverride (path ,newPath string ){if !_db .HasPrefix (path ,"\u002f"){path ="\u002f"+path ;};if !_db .HasPrefix (newPath ,"\u002f"){newPath ="\u002f"+newPath ;};for _cgf :=range _fdb ._beg .Override {if _fdb ._beg .Override [_cgf ].PartNameAttr ==path {_dc :=*_fdb ._beg .Override [_cgf ];_dc .PartNameAttr =newPath ;_fdb ._beg .Override =append (_fdb ._beg .Override ,&_dc );};};};func (_gaa CustomProperties )SetPropertyAsEmpty (name string ){_geab :=_gaa .getNewProperty (name );_geab .Empty =_ed .NewEmpty ();_gaa .setOrReplaceProperty (_geab );};

// SetCreated sets the time that the document was created.
func (_aaf CoreProperties )SetCreated (t _be .Time ){_aaf ._dfe .Created =_bcf (t ,"\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064");};

// SetLastModifiedBy records the last person to modify the document.
func (_edc CoreProperties )SetLastModifiedBy (s string ){_edc ._dfe .LastModifiedBy =&s };

// CustomProperties contains document specific properties.
type CustomProperties struct{_dag *_da .Properties };

// ImageRef is a reference to an image within a document.
type ImageRef struct{_bfbf *DocBase ;_gbfe Relationships ;_cgg Image ;_edcg string ;_fcbg string ;};

// FindRIDForN returns the relationship ID for the i'th relationship of type t.
func (_cdfa Relationships )FindRIDForN (i int ,t string )string {for _ ,_deg :=range _cdfa ._cdba .CT_Relationships .Relationship {if _deg .TypeAttr ==t {if i ==0{return _deg .IdAttr ;};i --;};};return "";};func (_cafe CustomProperties )SetPropertyAsR8 (name string ,r8 float64 ){_bgfa :=_cafe .getNewProperty (name );_bgfa .R8 =&r8 ;_cafe .setOrReplaceProperty (_bgfa );};func (_cfc TableRow )addCell ()*_eb .CT_TableCell {_afgf :=_eb .NewCT_TableCell ();_cfc ._edab .Tc =append (_cfc ._edab .Tc ,_afgf );return _afgf ;};

// X returns the inner wrapped XML type.
func (_gecg TableStyles )X ()*_eb .TblStyleLst {return _gecg ._fbf };

// Hyperlink is just an appropriately configured relationship.
type Hyperlink Relationship ;

// Target returns the target (path) of a relationship.
func (_fbg Relationship )Target ()string {return _fbg ._gaac .TargetAttr };

// ImageFromBytes returns an Image struct for an in-memory image. You can also
// construct an Image directly if the file and size are known.
func ImageFromBytes (data []byte )(Image ,error ){_gcbe :=Image {};_bdba ,_efgg ,_fec :=_bd .Decode (_ee .NewReader (data ));if _fec !=nil {return _gcbe ,_gef .Errorf ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s",_fec );};_gcbe .Data =&data ;_gcbe .Format =_efgg ;_gcbe .Size =_bdba .Bounds ().Size ();return _gcbe ,nil ;};

// NewCoreProperties constructs a new CoreProperties.
func NewCoreProperties ()CoreProperties {return CoreProperties {_dfe :_ba .NewCoreProperties ()}};

// Application returns the name of the application that created the document.
// For unioffice created documents, it defaults to gitee.com/gooffice/gooffice
func (_gc AppProperties )Application ()string {if _gc ._ade .Application !=nil {return *_gc ._ade .Application ;};return "";};

// Type returns the type of a relationship.
func (_gge Relationship )Type ()string {return _gge ._gaac .TypeAttr };

// Format returns the format of the underlying image
func (_ffg ImageRef )Format ()string {return _ffg ._cgg .Format };

// TblStyle returns the TblStyle property.
func (_cgae TableStyles )TblStyle ()[]*_eb .CT_TableStyle {return _cgae ._fbf .TblStyle };

// SetOffsetX sets horizontal offset of a table in distance units (see measurement package).
func (_ddgg Table )SetOffsetX (offX float64 ){if _ddgg ._ebd .Off ==nil {_ddgg ._ebd .Off =_eb .NewCT_Point2D ();_adg :=int64 (0);_ddgg ._ebd .Off .YAttr =_eb .ST_Coordinate {ST_CoordinateUnqualified :&_adg };};_aafd :=_ga .ToEMU (offX );_ddgg ._ebd .Off .XAttr =_eb .ST_Coordinate {ST_CoordinateUnqualified :&_aafd };};

// GetImageBytesByTarget returns Image object with Data bytes read from its target.
func (_aega *DocBase )GetImageBytesByTarget (target string )(Image ,error ){if target !=""{target ="\u0077\u006f\u0072d\u002f"+target ;for _ ,_cgcd :=range _aega .Images {if _cgcd .Target ()==target {return ImageFromStorage (_cgcd .Path ());};};};return Image {},_gfgd ;};

// Created returns the time that the document was created.
func (_cdcg CoreProperties )Created ()_be .Time {return _bbcb (_cdcg ._dfe .Created )};

// X returns the inner wrapped XML type of CustomProperty.
func (_dbeg CustomProperty )X ()*_da .CT_Property {return _dbeg ._eag };

// Pages returns total number of pages which are saved by the text editor which produced the document.
// For unioffice created documents, it is 0.
func (_eafb AppProperties )Pages ()int32 {if _eafb ._ade .Pages !=nil {return *_eafb ._ade .Pages ;};return 0;};

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct{_ade *_ebe .Properties };func init (){_bag .SetAsStorage ()};

// WriteExtraFiles writes the extra files to the zip package.
func (_dac *DocBase )WriteExtraFiles (z *_ef .Writer )error {for _ ,_dgfe :=range _dac .ExtraFiles {if _bde :=_fc .AddFileFromDisk (z ,_dgfe .ZipPath ,_dgfe .DiskPath );_bde !=nil {return _bde ;};};return nil ;};func (_edf CustomProperties )SetPropertyAsUi4 (name string ,ui4 uint32 ){_eced :=_edf .getNewProperty (name );_eced .Ui4 =&ui4 ;_edf .setOrReplaceProperty (_eced );};

// EnsureOverride ensures that an override for the given path exists, adding it if necessary
func (_baf ContentTypes )EnsureOverride (path ,contentType string ){for _ ,_dbe :=range _baf ._beg .Override {if _dbe .PartNameAttr ==path {if _db .HasPrefix (contentType ,"\u0068\u0074\u0074\u0070"){_bdg .Log .Debug ("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070",contentType );};_dbe .ContentTypeAttr =contentType ;return ;};};_baf .AddOverride (path ,contentType );};

// SetModified sets the time that the document was modified.
func (_aeb CoreProperties )SetModified (t _be .Time ){_aeb ._dfe .Modified =_bcf (t ,"\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064");};func (_ffb CustomProperties )SetPropertyAsBlob (name ,blob string ){_ecbd :=_ffb .getNewProperty (name );_ecbd .Blob =&blob ;_ffb .setOrReplaceProperty (_ecbd );};func (_cde CustomProperties )SetPropertyAsVstream (name string ,vstream *_ed .Vstream ){_edd :=_cde .getNewProperty (name );_edd .Vstream =vstream ;_cde .setOrReplaceProperty (_edd );};func (_aead Relationship )String ()string {return _gef .Sprintf ("\u007b\u0049\u0044\u003a \u0025\u0073\u0020\u0054\u0061\u0072\u0067\u0065\u0074\u003a \u0025s\u0020\u0054\u0079\u0070\u0065\u003a\u0020%\u0073\u007d",_aead .ID (),_aead .Target (),_aead .Type ());};func (_gab CustomProperties )SetPropertyAsVector (name string ,vector *_ed .Vector ){_bbcc :=_gab .getNewProperty (name );_bbcc .Vector =vector ;_gab .setOrReplaceProperty (_bbcc );};

// X returns the inner wrapped XML type.
func (_gda CoreProperties )X ()*_ba .CoreProperties {return _gda ._dfe };func (_cbf CustomProperties )SetPropertyAsDate (name string ,date _be .Time ){date =date .UTC ();_dcb ,_bfe ,_gdd :=date .Date ();_bfca ,_gagc ,_cbfg :=date .Clock ();_gfga :=_be .Date (_dcb ,_bfe ,_gdd ,_bfca ,_gagc ,_cbfg ,0,_be .UTC );_bega :=_cbf .getNewProperty (name );_bega .Filetime =&_gfga ;_cbf .setOrReplaceProperty (_bega );};

// AddHyperlink adds an external hyperlink relationship.
func (_dagc Relationships )AddHyperlink (target string )Hyperlink {_feag :=_dagc .AddRelationship (target ,_eca .HyperLinkType );_feag ._gaac .TargetModeAttr =_gb .ST_TargetModeExternal ;return Hyperlink (_feag );};const _fged =15;

// Remove removes an existing relationship.
func (_gcc Relationships )Remove (rel Relationship )bool {for _gadgb ,_ggf :=range _gcc ._cdba .Relationship {if _ggf ==rel ._gaac {copy (_gcc ._cdba .Relationship [_gadgb :],_gcc ._cdba .Relationship [_gadgb +1:]);_gcc ._cdba .Relationship =_gcc ._cdba .Relationship [0:len (_gcc ._cdba .Relationship )-1];return true ;};};return false ;};

// NewTableStyles constructs a new TableStyles.
func NewTableStyles ()TableStyles {return TableStyles {_fbf :_eb .NewTblStyleLst ()}};

// NewTheme constructs a new theme.
func NewTheme ()Theme {return Theme {_eb .NewTheme ()}};

// ContentStatus returns the content status of the document (e.g. "Final", "Draft")
func (_dcc CoreProperties )ContentStatus ()string {if _dcc ._dfe .ContentStatus !=nil {return *_dcc ._dfe .ContentStatus ;};return "";};

// RemoveOverride removes an override given a path.
func (_efd ContentTypes )RemoveOverride (path string ){if !_db .HasPrefix (path ,"\u002f"){path ="\u002f"+path ;};for _cdc ,_ace :=range _efd ._beg .Override {if _ace .PartNameAttr ==path {copy (_efd ._beg .Override [_cdc :],_efd ._beg .Override [_cdc +1:]);_efd ._beg .Override =_efd ._beg .Override [0:len (_efd ._beg .Override )-1];};};};

// X returns the inner wrapped XML type.
func (_egec AppProperties )X ()*_ebe .Properties {return _egec ._ade };

// X returns the inner wrapped XML type.
func (_ccfd Table )X ()*_eb .Tbl {return _ccfd ._bff };

// Company returns the name of the company that created the document.
// For unioffice created documents, it defaults to gitee.com/gooffice/gooffice
func (_af AppProperties )Company ()string {if _af ._ade .Company !=nil {return *_af ._ade .Company ;};return "";};

// ID returns the ID of a relationship.
func (_fdd Relationship )ID ()string {return _fdd ._gaac .IdAttr };

// Properties returns table properties.
func (_eee Table )Grid ()*_eb .CT_TableGrid {return _eee ._bff .TblGrid };

// IsEmpty returns true if there are no relationships.
func (_eec Relationships )IsEmpty ()bool {return _eec ._cdba ==nil ||len (_eec ._cdba .Relationship )==0;};

// Theme is a drawingml theme.
type Theme struct{_bdbf *_eb .Theme };const _aeaa =30;

// GetOrCreateCustomProperties returns the custom properties of the document (and if they not exist yet, creating them first).
func (_ced *DocBase )GetOrCreateCustomProperties ()CustomProperties {if _ced .CustomProperties .X ()==nil {_ced .CreateCustomProperties ();};return _ced .CustomProperties ;};

// AddRelationship adds a relationship.
func (_aed Relationships )AddRelationship (target ,ctype string )Relationship {if !_db .HasPrefix (ctype ,"\u0068t\u0074\u0070\u003a\u002f\u002f"){_bdg .Log .Debug ("\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006es\u0068\u0069\u0070 t\u0079\u0070\u0065\u0020\u0025\u0073 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069t\u0068\u0020\u0027\u0068\u0074\u0074\u0070\u003a/\u002f\u0027",ctype );};_bad :=_gb .NewRelationship ();_affe :=len (_aed ._cdba .Relationship )+1;_affg :=map[string ]struct{}{};for _ ,_cgge :=range _aed ._cdba .Relationship {_affg [_cgge .IdAttr ]=struct{}{};};for _ ,_cded :=_affg [_gef .Sprintf ("\u0072\u0049\u0064%\u0064",_affe )];_cded ;_ ,_cded =_affg [_gef .Sprintf ("\u0072\u0049\u0064%\u0064",_affe )]{_affe ++;};_bad .IdAttr =_gef .Sprintf ("\u0072\u0049\u0064%\u0064",_affe );_bad .TargetAttr =target ;_bad .TypeAttr =ctype ;_aed ._cdba .Relationship =append (_aed ._cdba .Relationship ,_bad );return Relationship {_gaac :_bad };};

// DefAttr returns the DefAttr property.
func (_bea TableStyles )DefAttr ()string {return _bea ._fbf .DefAttr };

// Modified returns the time that the document was modified.
func (_fdf CoreProperties )Modified ()_be .Time {return _bbcb (_fdf ._dfe .Modified )};

// TableRow represents a row in a table.
type TableRow struct{_edab *_eb .CT_TableRow };func (_fgec CustomProperties )SetPropertyAsLpwstr (name string ,lpwstr string ){_abf :=_fgec .getNewProperty (name );_abf .Lpwstr =&lpwstr ;_fgec .setOrReplaceProperty (_abf );};const _ccgf =10;

// Data returns the data of an image file, if any.
func (_afef ImageRef )Data ()*[]byte {return _afef ._cgg .Data };

// Clear removes any existing relationships.
func (_gadb Relationships )Clear (){_gadb ._cdba .Relationship =nil };func _bbcb (_acb *_eca .XSDAny )_be .Time {if _acb ==nil {return _be .Time {};};_ecb ,_bafe :=_be .Parse (_cca ,string (_acb .Data ));if _bafe !=nil {_bdg .Log .Debug ("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0074\u0069\u006d\u0065 \u0066\u0072\u006f\u006d\u0020\u0025\u0073\u003a\u0020\u0025\u0073",string (_acb .Data ),_bafe );};return _ecb ;};
