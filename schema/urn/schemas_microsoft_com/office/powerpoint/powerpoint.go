
package powerpoint ;import (_e "encoding/xml";_b "fmt";_ef "gitee.com/gooffice/gooffice";);func (_gf *Iscomment )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_gf .CT_Empty =*NewCT_Empty ();for {_bed ,_dd :=d .Token ();if _dd !=nil {return _b .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_dd );};if _ab ,_cg :=_bed .(_e .EndElement );_cg &&_ab .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_aec *CT_Empty )ValidateWithPath (path string )error {return nil };func (_bb *CT_Rel )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for _ ,_ec :=range start .Attr {if _ec .Name .Local =="\u0069\u0064"{_cc ,_g :=_ec .Value ,error (nil );if _g !=nil {return _g ;};_bb .IdAttr =&_cc ;continue ;};};for {_ecg ,_de :=d .Token ();if _de !=nil {return _b .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_de );};if _fe ,_ba :=_ecg .(_e .EndElement );_ba &&_fe .Name ==start .Name {break ;};};return nil ;};func NewCT_Empty ()*CT_Empty {_eg :=&CT_Empty {};return _eg };func NewTextdata ()*Textdata {_gef :=&Textdata {};_gef .CT_Rel =*NewCT_Rel ();return _gef };

// Validate validates the Textdata and its children
func (_ada *Textdata )Validate ()error {return _ada .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};

// Validate validates the Iscomment and its children
func (_ag *Iscomment )Validate ()error {return _ag .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_egd *Iscomment )ValidateWithPath (path string )error {if _egf :=_egd .CT_Empty .ValidateWithPath (path );_egf !=nil {return _egf ;};return nil ;};

// Validate validates the CT_Rel and its children
func (_ge *CT_Rel )Validate ()error {return _ge .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func (_c *CT_Rel )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _c .IdAttr !=nil {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0069\u0064"},Value :_b .Sprintf ("\u0025\u0076",*_c .IdAttr )});};e .EncodeToken (start );e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};type CT_Rel struct{IdAttr *string ;};func NewIscomment ()*Iscomment {_cd :=&Iscomment {};_cd .CT_Empty =*NewCT_Empty ();return _cd };func (_fd *Textdata )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _fd .CT_Rel .MarshalXML (e ,start );};func (_ea *CT_Empty )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};type Textdata struct{CT_Rel };func (_acc *Iscomment )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _acc .CT_Empty .MarshalXML (e ,start );};type CT_Empty struct{};func (_cf *Textdata )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_cf .CT_Rel =*NewCT_Rel ();for _ ,_eag :=range start .Attr {if _eag .Name .Local =="\u0069\u0064"{_ad ,_ce :=_eag .Value ,error (nil );if _ce !=nil {return _ce ;};_cf .IdAttr =&_ad ;continue ;};};for {_fa ,_bd :=d .Token ();if _bd !=nil {return _b .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_bd );};if _gc ,_bf :=_fa .(_e .EndElement );_bf &&_gc .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Empty and its children
func (_ed *CT_Empty )Validate ()error {return _ed .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};type Iscomment struct{CT_Empty };func NewCT_Rel ()*CT_Rel {_be :=&CT_Rel {};return _be };

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_fec *Textdata )ValidateWithPath (path string )error {if _ca :=_fec .CT_Rel .ValidateWithPath (path );_ca !=nil {return _ca ;};return nil ;};

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_ac *CT_Rel )ValidateWithPath (path string )error {return nil };func (_bc *CT_Empty )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for {_a ,_ae :=d .Token ();if _ae !=nil {return _b .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_ae );};if _d ,_eb :=_a .(_e .EndElement );_eb &&_d .Name ==start .Name {break ;};};return nil ;};func init (){_ef .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );_ef .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );_ef .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );_ef .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );};
