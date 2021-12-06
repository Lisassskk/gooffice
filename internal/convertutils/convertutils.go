
package convertutils ;import (_b "bytes";_gg "errors";_eb "fmt";_bc "gitee.com/gooffice/gooffice/common/logger";_gd "gitee.com/gooffice/gooffice/measurement";_ab "gitee.com/gooffice/gooffice/schema/soo/dml";_cg "gitee.com/gooffice/gooffice/schema/soo/dml/chart";_fg "gitee.com/gooffice/gooffice/spreadsheet/format";_ag "github.com/unidoc/unipdf/v3/creator";_fd "github.com/unidoc/unipdf/v3/model";_gfa "github.com/unidoc/unipdf/v3/render";_cc "github.com/unidoc/unitype";_c "image";_d "math";_e "os";_cf "sort";_a "strconv";_f "strings";_gf "sync";_df "unicode";);func _geddb (_gec float64 )float64 {if _gec < 0{_gec +=float64 (-int (_gec )+1);}else if _gec > 1{_gec -=float64 (int (_gec ));};return _gec ;};func _cab (_gdde float64 )float64 {return _gdde *_gd .Millimeter };func IsNoSpaceLanguage (symbol string )bool {for _ ,_aeea :=range symbol {if _df .Is (_df .Han ,_aeea ){return true ;};};return false ;};type barSerByOrder []*_cg .CT_BarSer ;type fontsMap struct{_dge *_gf .Mutex ;_ccge map[string ]map[FontStyle ]*_fd .PdfFont ;};var _gedd =_cab (2);func MakeImageFromChartSpace (cs *_cg .ChartSpace ,width ,height float64 ,theme *_ab .Theme )(_c .Image ,error ){_ece ,_effc :=_bcd (cs ,width ,height ,theme ,true );if _effc !=nil {return nil ,_effc ;};_adc ,_effc :=GetPageFromCreator (_ece );if _effc !=nil {return nil ,_effc ;};return _gfa .NewImageDevice ().Render (_adc );};func GetColorStringFromDmlColor (dmlColor *_ab .CT_Color )string {var _dacc string ;if _eef :=dmlColor .SrgbClr ;_eef !=nil {_dacc =_eef .ValAttr ;}else if _bce :=dmlColor .SysClr ;_bce !=nil {return "\u0030\u0030\u0030\u0030\u0030\u0030";};return _dacc ;};func _ecaa (_gafb uint8 ,_dfe float64 )string {_aabd :=float64 (_gafb );return _eb .Sprintf ("\u0025\u0030\u0032\u0078",int (_aabd *_dfe ));};func _bdd (_fabg ,_aeg ,_bcb float64 )(uint8 ,uint8 ,uint8 ){var _bfea float64 ;if _bcb < 0.5{_bfea =_bcb *(1+_aeg );}else {_bfea =_bcb +_aeg -_bcb *_aeg ;};_fbg :=_bcb *2-_bfea ;_fabg /=360.0;_gfbf :=_geddb (_fabg +1.0/3.0);_dgcg :=_geddb (_fabg );_ccb :=_geddb (_fabg -1.0/3.0);_bbdd :=_fge (_gfbf ,_bfea ,_fbg );_eaf :=_fge (_dgcg ,_bfea ,_fbg );_bcbg :=_fge (_ccb ,_bfea ,_fbg );return uint8 (255*_bbdd ),uint8 (255*_eaf ),uint8 (255*_bcbg );};const _ffgb =6.0;const (FontStyle_Regular FontStyle =0;FontStyle_Bold FontStyle =1;FontStyle_Italic FontStyle =2;FontStyle_BoldItalic FontStyle =3;);var _efea =_cab (7.5);var _abd =_cab (1.5);func (_bdb *creatorContext )drawAxes (_gea *_cg .CT_PlotAreaChoice1 ,_fcdb ,_edg ,_efee float64 ,_bec []string ,_ffb *Rectangle ,_cgd bool )error {_aafd :=_bdb ._ecf ;_cca :=_bdb ._bbcf ;if _gea ==nil {return _gg .New ("\u004e\u006f\u0020\u0061xi\u0073\u0020\u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e");};if len (_gea .ValAx )==0||(len (_gea .CatAx )==0&&len (_gea .DateAx )==0&&len (_gea .SerAx )==0){return _gg .New ("\u004e\u006f\u0020\u0078\u0020\u006f\u0072\u0020\u0079 \u0061\u0078\u0069\u0073");};var _deg ,_ace ,_ffgd ,_fff uint32 ;var _afcg ,_bde _cg .ST_AxPos ;var _eegd ,_fec _cg .ST_TickMark ;var _agg ,_fgag *_cg .CT_ChartLines ;var _ecdf ,_gcf _cg .ST_TickLblPos ;var _gac ,_efb *_ab .CT_ShapeProperties ;var _fgagd error ;_bdcc :=_ffb .Right -_ffb .Left ;_acfg :=_ffb .Bottom -_ffb .Top ;if len (_gea .ValAx )> 0{_ace ,_bde ,_fec ,_gcf ,_fgag ,_fff ,_efb ,_fgagd =_gggd (_gea .ValAx [0]);};if _bde !=_cg .ST_AxPosL &&_bde !=_cg .ST_AxPosB {return _gg .New ("\u004f\u006e\u006c\u0079\u0020l\u0065\u0066\u0074\u0020\u006f\u0072\u0020\u0062\u006f\u0074\u0074\u006f\u006d \u0079\u0020\u0061\u0078\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0073\u006f\u0020\u0066\u0061\u0072");};_afbe :=_edg -_fcdb ;_dfg :=int (_afbe /_efee )+1;var _acd ,_gde float64 ;switch _fec {case _cg .ST_TickMarkIn :_acd ,_gde =_bfd ,0;case _cg .ST_TickMarkOut :_acd ,_gde =0,_bfd ;case _cg .ST_TickMarkCross :_acd ,_gde =_bfd ,_bfd ;};_acd =_acd *_cca ;_gde =_gde *_cca ;var _ffga *_ab .CT_ShapeProperties ;if _fgag !=nil {_ffga =_fgag .SpPr ;};_fcc ,_dcde :=_ecdf !=_cg .ST_TickLblPosNone ,_gcf !=_cg .ST_TickLblPosNone ;_afe :=_fcdb ;if len (_gea .CatAx )> 0{_deg ,_afcg ,_eegd ,_ecdf ,_agg ,_ffgd ,_gac ,_fgagd =_dca (_gea .CatAx [0]);}else if len (_gea .DateAx )> 0{_deg ,_afcg ,_eegd ,_ecdf ,_agg ,_ffgd ,_gac ,_fgagd =_fgdb (_gea .DateAx [0]);}else if len (_gea .SerAx )> 0{_deg ,_afcg ,_eegd ,_ecdf ,_agg ,_ffgd ,_gac ,_fgagd =_gdf (_gea .SerAx [0]);};if _fgagd !=nil {return _fgagd ;};if _afcg !=_cg .ST_AxPosL &&_afcg !=_cg .ST_AxPosB {return _gg .New ("\u004f\u006e\u006c\u0079\u0020l\u0065\u0066\u0074\u0020\u006f\u0072\u0020\u0062\u006f\u0074\u0074\u006f\u006d \u0078\u0020\u0061\u0078\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0073\u006f\u0020\u0066\u0061\u0072");};if _deg !=_fff ||_ace !=_ffgd {return _gg .New ("a\u0078i\u0073\u0020\u0069\u0064\u0073\u0020\u0064\u006fn\u0027\u0074\u0020\u006dat\u0063\u0068");};_fgb :=len (_bec )+1;var _dbf ,_gba float64 ;switch _eegd {case _cg .ST_TickMarkIn :_dbf ,_gba =_bfd ,0;case _cg .ST_TickMarkOut :_dbf ,_gba =0,_bfd ;case _cg .ST_TickMarkCross :_dbf ,_gba =_bfd ,_bfd ;};_dbf =_dbf *_cca ;_gba =_gba *_cca ;var _fgfb *_ab .CT_ShapeProperties ;if _agg !=nil {_fgfb =_agg .SpPr ;};if _cgd {_egfe :=_acfg /float64 (len (_bec ));_gaa :=_ffb .Left -_fcdb *_bdcc /_afbe ;_eed :=_gaa -_efea *_cca ;if _fcc {var _eeda float64 ;for _dagd :=0;_dagd < _fgb ;_dagd ++{_cfa :=_ffb .Bottom -float64 (_dagd )*_egfe ;if _dagd < _fgb -1{_dbe :=_aafd .NewParagraph (_bec [_dagd ]);_dbe .SetFontSize (_ffgb *_cca );_dbe .SetPos (_eed ,_cfa -_egfe /2-_acc *_cca );_aafd .Draw (_dbe );_daa :=(_dbe .Width ()/1000-_efea )*_cca ;if _daa > 0&&_daa > _eeda {_eeda =_daa ;};};};if _eeda > 0{_ffb .Left +=_eeda +_gedd ;_gaa =_ffb .Left -_fcdb *_bdcc /_afbe ;_bdcc =_ffb .Right -_ffb .Left ;};};_acg :=_gaa -_gba ;_gcb :=_gaa +_dbf ;_adef :=_ffb .Left ;_bgce :=_ffb .Right ;for _gef :=0;_gef < _fgb ;_gef ++{_cfag :=_ffb .Bottom -float64 (_gef )*_egfe ;_bdb .drawLineWithProps (_gac ,_acg ,_cfag ,_gcb ,_cfag ,true );_bdb .drawLineWithProps (_fgfb ,_adef ,_cfag ,_bgce ,_cfag ,true );};_cgc :=_bdcc /_afbe ;_fde :=_ffb .Bottom -_acd ;_caa :=_ffb .Bottom +_gde ;_ceb :=_ffb .Top ;_bbd :=_ffb .Bottom ;for _aad :=0;_aad < _dfg ;_aad ++{_bed :=_ffb .Left +(_afe -_fcdb )*_cgc ;_bdb .drawLineWithProps (_efb ,_bed ,_fde ,_bed ,_caa ,true );_bdb .drawLineWithProps (_ffga ,_bed ,_ceb ,_bed ,_bbd ,true );if _dcde {_agc :=_aafd .NewParagraph (_a .FormatFloat (_afe ,'g',-1,64));_agc .SetFontSize (_ffgb *_cca );_agc .SetPos (_bed -_baee *_cca ,_bbd +_abd *_cca );_aafd .Draw (_agc );};_afe +=_efee ;};}else {_fdaeb :=_acfg /_afbe ;_fdd :=_ffb .Left ;if _dcde {var _fdaf float64 ;for _cgg :=0;_cgg < _dfg ;_cgg ++{_cdf :=_ffb .Bottom -(_afe -_fcdb )*_fdaeb ;_dfff :=_aafd .NewParagraph (_a .FormatFloat (_afe ,'g',-1,64));_dfff .SetFontSize (_ffgb *_cca );_dfff .SetPos (_fdd -_efea *_cca ,_cdf -_acc *_cca );_aafd .Draw (_dfff );_adfe :=(_dfff .Width ()/1000-_efea )*_cca ;if _adfe > 0&&_adfe > _fdaf {_fdaf =_adfe ;};_afe +=_efee ;};if _fdaf > 0{_ffb .Left +=_fdaf +_gedd ;_bdcc =_ffb .Right -_ffb .Left ;};};_afe =_fcdb ;_aae :=_ffb .Left -_gde ;_fdc :=_ffb .Left +_acd ;_fdd =_ffb .Left ;_beg :=_ffb .Right ;for _bbb :=0;_bbb < _dfg ;_bbb ++{_agde :=_ffb .Bottom -(_afe -_fcdb )*_fdaeb ;_bdb .drawLineWithProps (_efb ,_aae ,_agde ,_fdc ,_agde ,true );_bdb .drawLineWithProps (_ffga ,_fdd ,_agde ,_beg ,_agde ,true );_afe +=_efee ;};_abg :=_bdcc /float64 (len (_bec ));_dbfc :=_ffb .Bottom +_fcdb *_acfg /_afbe ;_ggc :=_dbfc -_dbf ;_dcgc :=_dbfc +_gba ;_bcda :=_ffb .Top ;_gae :=_ffb .Bottom ;_gfg :=_dbfc +_abd *_cca ;for _bgb :=0;_bgb < _fgb ;_bgb ++{_efg :=_ffb .Left +float64 (_bgb )*_abg ;_bdb .drawLineWithProps (_gac ,_efg ,_ggc ,_efg ,_dcgc ,true );_bdb .drawLineWithProps (_fgfb ,_efg ,_bcda ,_efg ,_gae ,true );if _fcc &&_bgb < _fgb -1{_fagf :=_aafd .NewParagraph (_bec [_bgb ]);_fagf .SetFontSize (_ffgb *_cca );_fagf .SetPos (_efg +_bfe *_cca ,_gfg );_aafd .Draw (_fagf );};};};return nil ;};const (ImgPart_whole ImgPart =0;ImgPart_t ImgPart =1;ImgPart_b ImgPart =2;ImgPart_l ImgPart =3;ImgPart_r ImgPart =4;ImgPart_lt ImgPart =5;ImgPart_rt ImgPart =6;ImgPart_lb ImgPart =7;ImgPart_rb ImgPart =8;);func (_fgab *creatorContext )getPdfColorFromSolidFill (_aaccg *_ab .CT_SolidColorFillProperties )_ag .Color {if _aaccg ==nil {return nil ;};_dggg :="";if _caae :=_aaccg .SrgbClr ;_caae !=nil {_dggg =_caae .ValAttr ;}else if _efdf :=_aaccg .SchemeClr ;_efdf !=nil {_dggg =_dafg (_efdf .ValAttr ,_fgab ._cacg );};if _dggg ==""{return nil ;};return _ag .ColorRGBFromHex ("\u0023"+_dggg );};func GetRegisteredFont (name string ,style FontStyle )*_fd .PdfFont {_ebe ._dge .Lock ();defer _ebe ._dge .Unlock ();if _fdfgf ,_dda :=_ebe ._ccge [name ];_dda {if _edga ,_bede :=_fdfgf [style ];_bede {return _edga ;};};return nil ;};func AdjustColorByLumMod (colorStr string ,lum float64 )string {var _ffcf ,_acb ,_deaa uint8 ;_baf ,_ :=_eb .Sscanf (colorStr ,"\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078",&_ffcf ,&_acb ,&_deaa );if _baf !=3{return "";};_fdfg ,_aecf ,_cgcc :=_bdg (_ffcf ,_acb ,_deaa );_cgcc =lum *_cgcc ;_ffcf ,_acb ,_deaa =_bdd (_fdfg ,_aecf ,_cgcc );return _eb .Sprintf ("\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078",_ffcf ,_acb ,_deaa );};func (_fbdf *creatorContext )drawRectangleWithProps (_bef *_ab .CT_ShapeProperties ,_aefe ,_babe ,_babd ,_gaab float64 ,_defa bool ){_dfde :=_fbdf ._ecf .NewRectangle (_aefe ,_babe ,_babd ,_gaab );if _bef ==nil {if _defa {_dfde .SetBorderWidth (_gge );}else {return ;};}else {_acfc :=_fbdf .getPdfColorFromSolidFill (_bef .SolidFill );if _acfc !=nil {_dfde .SetFillColor (_acfc );};if _fgfad :=_bef .Ln ;_fgfad !=nil {if _deca :=_fgfad .WAttr ;_deca !=nil {_gdef :=_gd .FromEMU (int64 (*_deca ));_dfde .SetBorderWidth (_gdef );if _eadb :=_fgfad .SolidFill ;_eadb !=nil {_bad :=_fbdf .getPdfColorFromSolidFill (_eadb );if _bad !=nil {_dfde .SetBorderColor (_bad );};};}else {_dfde .SetBorderWidth (0);};};};_fbdf ._ecf .Draw (_dfde );};func _fgdb (_decb *_cg .CT_DateAx )(uint32 ,_cg .ST_AxPos ,_cg .ST_TickMark ,_cg .ST_TickLblPos ,*_cg .CT_ChartLines ,uint32 ,*_ab .CT_ShapeProperties ,error ){var _efeea ,_feae uint32 ;var _ada _cg .ST_AxPos ;var _aaeb _cg .ST_TickMark ;var _fcb *_cg .CT_ChartLines ;var _gaec _cg .ST_TickLblPos ;if _decb .AxId ==nil {return _efeea ,_ada ,_aaeb ,_gaec ,_fcb ,_feae ,_decb .SpPr ,_gg .New ("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_efeea =_decb .AxId .ValAttr ;};if _decb .AxPos ==nil {return _efeea ,_ada ,_aaeb ,_gaec ,_fcb ,_feae ,_decb .SpPr ,_gg .New ("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e");}else {_ada =_decb .AxPos .ValAttr ;};if _decb .MajorTickMark !=nil {_aaeb =_decb .MajorTickMark .ValAttr ;};if _decb .TickLblPos !=nil {_gaec =_decb .TickLblPos .ValAttr ;};if _decb .CrossAx ==nil {return _efeea ,_ada ,_aaeb ,_gaec ,_fcb ,_feae ,_decb .SpPr ,_gg .New ("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_feae =_decb .CrossAx .ValAttr ;};_fcb =_decb .MajorGridlines ;return _efeea ,_ada ,_aaeb ,_gaec ,_fcb ,_feae ,_decb .SpPr ,nil ;};type creatorContext struct{_ecf *_ag .Creator ;_cacg *_ab .Theme ;_bbcf float64 ;};func (_bgceb *creatorContext )drawLineWithProps (_aafdd *_ab .CT_ShapeProperties ,_ebf ,_cea ,_cgcf ,_dgcf float64 ,_egfee bool ){if _aafdd !=nil {if _acfb :=_aafdd .Ln ;_acfb !=nil {_ccad :=_bgceb .getPdfColorFromSolidFill (_acfb .SolidFill );if _ccad ==nil &&_egfee {_ccad =_ag .ColorBlack ;};if _ccad !=nil {var _gdgb float64 ;if _edfa :=_acfb .WAttr ;_edfa !=nil {_gdgb =_gd .FromEMU (int64 (*_edfa ));}else {_gdgb =_gge ;};DrawLine (_bgceb ._ecf ,_ebf ,_cea ,_cgcf ,_dgcf ,_gdgb ,_ccad );};};};};func AdjustColorByShade (colorStr string ,shade float64 )string {var _adaf ,_gcbf ,_aaff uint8 ;_agf ,_ :=_eb .Sscanf (colorStr ,"\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078",&_adaf ,&_gcbf ,&_aaff );if _agf !=3{return "";};return _ecaa (_adaf ,shade )+_ecaa (_gcbf ,shade )+_ecaa (_aaff ,shade );};var _bfd =_cab (1);func _fge (_fed ,_cee ,_cfd float64 )float64 {if _fed *6< 1{return _cfd +(_cee -_cfd )*6*_fed ;}else if _fed *2< 1{return _cee ;}else if _fed *3< 2{return _cfd +(_cee -_cfd )*(2.0/3.0-_fed )*6;}else {return _cfd ;};};func (_acf barSerByOrder )Swap (i ,j int ){_acf [i ],_acf [j ]=_acf [j ],_acf [i ]};var _acc =_cab (1.5);type serValue struct{_aac string ;_fgfc float64 ;_adf *_ab .CT_ShapeProperties ;};func GetPageFromCreator (c *_ag .Creator )(*_fd .PdfPage ,error ){_ecaf :=_b .NewBuffer ([]byte {});_dcag :=c .Write (_ecaf );if _dcag !=nil {return nil ,_dcag ;};_ffcc :=_b .NewReader (_ecaf .Bytes ());_ebdf ,_dcag :=_fd .NewPdfReader (_ffcc );if _dcag !=nil {return nil ,_dcag ;};return _ebdf .GetPage (1);};func (_dg *creatorContext )drawLegend (_gdd *Rectangle ,_dag []*legendItem ,_gfb bool ){_dgc :=_dg ._bbcf ;_ebda :=_cab (2.5)*_dgc ;_fgc :=_efe *_dgc ;_fdae :=(_ebda -_fgc )/2;_dfd :=float64 (len (_dag ));if _gfb {_dde :=&Rectangle {Top :_gdd .Top +_cab (1)*_dgc ,Bottom :_gdd .Bottom -_cab (1)*_dgc ,Left :_gdd .Left +_cab (2.5)*_dgc ,Right :_gdd .Right -_cab (2.5)*_dgc };var _egf float64 ;if _dfd > 1{_egf =(_dde .Right -_dde .Left )/_dfd ;};_ged :=_dde .Left ;_bdc :=_dde .Top ;for _ ,_dcec :=range _dag {if _aaa :=_dcec ._def ;_aaa !=nil {_dg .drawRectangleWithProps (_aaa ,_ged ,_bdc +_fdae ,_fgc ,_fgc ,false );_fcd :=_ged +_fgc *2;_cef :=_dg ._ecf .NewStyledParagraph ();_cef .SetPos (_fcd ,_bdc );_efa :=_cef .Append (_dcec ._gfe );_cge ,_cga :=_fd .NewStandard14Font (_fd .HelveticaName );if _cga ==nil {_efa .Style =_ag .TextStyle {FontSize :_ebda ,Font :_cge ,TextRise :0.4};_dg ._ecf .Draw (_cef );};};_ged +=_egf ;};}else {_dgcb :=&Rectangle {Top :_gdd .Top +_cab (2.5)*_dgc ,Bottom :_gdd .Bottom -_cab (2.5)*_dgc ,Left :_gdd .Left +_cab (2.5)*_dgc ,Right :_gdd .Right -_cab (2.5)*_dgc };var _bfc float64 ;if _dfd > 1{_bfc =(_dgcb .Bottom -_dgcb .Top -_ebda )/(_dfd -1);};_fgg :=_dgcb .Top ;_ffg :=_dgcb .Left ;_cdg :=_ffg +_fgc *2;for _ ,_adff :=range _dag {if _adfb :=_adff ._def ;_adfb !=nil {_dg .drawRectangleWithProps (_adfb ,_ffg ,_fgg +_fdae ,_fgc ,_fgc ,false );_ceff :=_dg ._ecf .NewStyledParagraph ();_ceff .SetPos (_cdg ,_fgg );_cbe :=_ceff .Append (_adff ._gfe );_aef ,_aee :=_fd .NewStandard14Font (_fd .HelveticaName );if _aee ==nil {_cbe .Style =_ag .TextStyle {FontSize :_ebda ,Font :_aef ,TextRise :0.4};_dg ._ecf .Draw (_ceff );};};_fgg +=_bfc ;};};};func _dca (_cgdg *_cg .CT_CatAx )(uint32 ,_cg .ST_AxPos ,_cg .ST_TickMark ,_cg .ST_TickLblPos ,*_cg .CT_ChartLines ,uint32 ,*_ab .CT_ShapeProperties ,error ){var _efcd ,_aacc uint32 ;var _begf _cg .ST_AxPos ;var _abgg _cg .ST_TickMark ;var _caf *_cg .CT_ChartLines ;var _geff _cg .ST_TickLblPos ;if _cgdg .AxId ==nil {return _efcd ,_begf ,_abgg ,_geff ,_caf ,_aacc ,_cgdg .SpPr ,_gg .New ("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_efcd =_cgdg .AxId .ValAttr ;};if _cgdg .AxPos ==nil {return _efcd ,_begf ,_abgg ,_geff ,_caf ,_aacc ,_cgdg .SpPr ,_gg .New ("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e");}else {_begf =_cgdg .AxPos .ValAttr ;};if _cgdg .MajorTickMark !=nil {_abgg =_cgdg .MajorTickMark .ValAttr ;};if _cgdg .TickLblPos !=nil {_geff =_cgdg .TickLblPos .ValAttr ;};if _cgdg .CrossAx ==nil {return _efcd ,_begf ,_abgg ,_geff ,_caf ,_aacc ,_cgdg .SpPr ,_gg .New ("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_aacc =_cgdg .CrossAx .ValAttr ;};_caf =_cgdg .MajorGridlines ;return _efcd ,_begf ,_abgg ,_geff ,_caf ,_aacc ,_cgdg .SpPr ,nil ;};var _ebe =fontsMap {_dge :&_gf .Mutex {},_ccge :map[string ]map[FontStyle ]*_fd .PdfFont {}};func _bdg (_gafc ,_dggb ,_dfc uint8 )(float64 ,float64 ,float64 ){_ecag ,_eaae ,_dcdf :=float64 (_gafc )/255,float64 (_dggb )/255,float64 (_dfc )/255;_cfb :=_ecag ;if _eaae < _cfb {_cfb =_eaae ;};if _dcdf < _cfb {_cfb =_dcdf ;};var _ddd ,_aeb bool ;_dbee :=_ecag ;if _eaae > _dbee {_dbee =_eaae ;_ddd =true ;};if _dcdf > _dbee {_dbee =_dcdf ;_ddd =false ;_aeb =true ;};_efba :=(_cfb +_dbee )/2;var _cgff float64 ;if _cfb !=_dbee {if _efba <=0.5{_cgff =(_dbee -_cfb )/(_dbee +_cfb );}else {_cgff =(_dbee -_cfb )/(2.0-_dbee -_cfb );};};var _cafg float64 ;if _cfb !=_dbee {if _ddd {_cafg =2.0+(_dcdf -_ecag )/(_dbee -_cfb );}else if _aeb {_cafg =4.0+(_ecag -_eaae )/(_dbee -_cfb );}else {_cafg =(_eaae -_dcdf )/(_dbee -_cfb );};_cafg *=60;if _cafg < 0{_cafg +=360;};};return _cafg ,_cgff ,_efba ;};func _ecc (_caaf uint8 ,_facg float64 )string {_cgae :=float64 (_caaf );var _aag float64 ;if _facg < 0{_aag =_cgae *(1+_facg );}else {_aag =_cgae +(255-_cgae )*_facg ;};return _eb .Sprintf ("\u0025\u0030\u0032\u0078",int (_aag ));};type Rectangle struct{Top float64 ;Bottom float64 ;Left float64 ;Right float64 ;};func FromSTCoordinate32 (st _ab .ST_Coordinate32 )int64 {if _bgg :=st .ST_Coordinate32Unqualified ;_bgg !=nil {return int64 (*_bgg );};return 0;};var _bfe =_cab (5);func (_gbeg barSerByOrder )Less (i ,j int )bool {return _gbeg [i ].Order .ValAttr < _gbeg [j ].Order .ValAttr ;};func (_ebgg *Rectangle )scale (_cfe float64 ){_ebgg .Top *=_cfe ;_ebgg .Bottom *=_cfe ;_ebgg .Left *=_cfe ;_ebgg .Right *=_cfe ;};func _bcd (_facc *_cg .ChartSpace ,_fdf ,_dae float64 ,_agaf *_ab .Theme ,_bga bool )(*_ag .Creator ,error ){_ebcb :=1.0;if _bga {_ebcb =8.0;};_fcg :=&Rectangle {};_dac :=&Rectangle {Top :_fcg .Top ,Bottom :_dae -_fcg .Bottom ,Left :_fcg .Left ,Right :_fdf -_fcg .Right };_aec :=MakeTempCreator (_fdf *_ebcb +1,_dae *_ebcb +1);_dec :=&creatorContext {_ecf :_aec ,_cacg :_agaf ,_bbcf :_ebcb };var _fagb bool ;if _dba :=_facc .Chart ;_dba !=nil {_ggg :=_dba .PlotArea ;if _ggg ==nil {return nil ,_gg .New ("\u004e\u006f\u0020p\u006c\u006f\u0074\u0020\u0061\u0072\u0065\u0061");};_bdf :=&Rectangle {Top :_cab (10),Bottom :_dac .Bottom -_cab (15),Left :_cab (10),Right :_dac .Right -_cab (10)};var _ca *Rectangle ;_ead :=_dba .Legend ;if _ead !=nil {_ffa :=_ead .Overlay !=nil &&_ead .Overlay .ValAttr !=nil &&*_ead .Overlay .ValAttr ;if _deaca :=_ead .LegendPos ;_deaca !=nil {switch _deaca .ValAttr {case _cg .ST_LegendPosTr :if !_ffa {_bdf =&Rectangle {Top :_cab (25),Bottom :_dac .Bottom -_cab (10),Left :_cab (10),Right :_dac .Right -_cab (25)};};_ca =&Rectangle {Top :_cab (2.5),Bottom :_cab (22.5),Left :_dac .Right -_cab (22.5),Right :_dac .Right -_cab (2.5)};case _cg .ST_LegendPosT :_ca =&Rectangle {Top :_cab (2.5),Bottom :_cab (7.5),Left :(_dac .Right -_dac .Left )*0.25,Right :(_dac .Right -_dac .Left )*0.75};if !_ffa {_bdf =&Rectangle {Top :_cab (12.5),Bottom :_dac .Bottom -_cab (15),Left :_cab (10),Right :_dac .Right -_cab (5)};};_fagb =true ;case _cg .ST_LegendPosB :_ca =&Rectangle {Top :_dac .Bottom -_cab (7.5),Bottom :_dac .Bottom -_cab (2.5),Left :(_dac .Right -_dac .Left )*0.25,Right :(_dac .Right -_dac .Left )*0.75};if !_ffa {_bdf =&Rectangle {Top :_cab (5),Bottom :_dac .Bottom -_cab (15),Left :_cab (10),Right :_dac .Right -_cab (5)};};_fagb =true ;case _cg .ST_LegendPosR :_ca =&Rectangle {Top :(_dac .Bottom -_dac .Top )/2-_cab (10),Bottom :(_dac .Bottom -_dac .Top )/2+_cab (10),Left :_dac .Right -_cab (22.5),Right :_dac .Right -_cab (2.5)};if !_ffa {_bdf =&Rectangle {Top :_cab (5),Bottom :_dac .Bottom -_cab (12.5),Left :_cab (10),Right :_dac .Right -_cab (25)};};case _cg .ST_LegendPosL :_ca =&Rectangle {Top :(_dac .Bottom -_dac .Top )/2-_cab (10),Bottom :(_dac .Bottom -_dac .Top )/2+_cab (10),Left :_cab (2.5),Right :_cab (22.5)};if !_ffa {_bdf =&Rectangle {Top :_cab (5),Bottom :_dac .Bottom -_cab (12.5),Left :_cab (30),Right :_dac .Right -_cab (5)};};default:_ca =&Rectangle {Top :(_dac .Bottom -_dac .Top )/2-_cab (10),Bottom :(_dac .Bottom -_dac .Top )/2+_cab (10),Left :_dac .Right -_cab (25),Right :_dac .Right -_cab (5)};if !_ffa {_bdf =&Rectangle {Top :_cab (5),Bottom :_dac .Bottom -_cab (12.5),Left :_cab (100),Right :_dac .Right -_cab (25)};};};};};_bdf .scale (_ebcb );_dec .drawBorderWithProps (_ggg .SpPr ,_bdf ,_gge );_ggb :=[]*legendItem {};var _efc error ;_ccf :=_ggg .CChoice ;for _ ,_egd :=range _ggg .Choice {if _fgd :=_egd .BarChart ;_fgd !=nil {_ggb ,_efc =_dec .drawBarChart (_fgd ,_bdf ,_ccf );if _efc !=nil {return nil ,_efc ;};};};if _ead !=nil {_ca .scale (_ebcb );_dec .drawBorderWithProps (_ead .SpPr ,_ca ,_gge );if len (_ggb )!=0{_dec .drawLegend (_ca ,_ggb ,_fagb );};};};_dac .scale (_ebcb );_dec .drawBorderWithProps (_facc .SpPr ,_dac ,_gge );return _aec ,nil ;};func FromSTCoordinate (st _ab .ST_Coordinate )int64 {if _fffa :=st .ST_CoordinateUnqualified ;_fffa !=nil {return *_fffa ;};return 0;};func _gdf (_abf *_cg .CT_SerAx )(uint32 ,_cg .ST_AxPos ,_cg .ST_TickMark ,_cg .ST_TickLblPos ,*_cg .CT_ChartLines ,uint32 ,*_ab .CT_ShapeProperties ,error ){var _cac ,_ccg uint32 ;var _cdfe _cg .ST_AxPos ;var _fab _cg .ST_TickMark ;var _acga *_cg .CT_ChartLines ;var _dgd _cg .ST_TickLblPos ;if _abf .AxId ==nil {return _cac ,_cdfe ,_fab ,_dgd ,_acga ,_ccg ,_abf .SpPr ,_gg .New ("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_cac =_abf .AxId .ValAttr ;};if _abf .AxPos ==nil {return _cac ,_cdfe ,_fab ,_dgd ,_acga ,_ccg ,_abf .SpPr ,_gg .New ("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e");}else {_cdfe =_abf .AxPos .ValAttr ;};if _abf .MajorTickMark !=nil {_fab =_abf .MajorTickMark .ValAttr ;};if _abf .TickLblPos !=nil {_dgd =_abf .TickLblPos .ValAttr ;};if _abf .CrossAx ==nil {return _cac ,_cdfe ,_fab ,_dgd ,_acga ,_ccg ,_abf .SpPr ,_gg .New ("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_ccg =_abf .CrossAx .ValAttr ;};_acga =_abf .MajorGridlines ;return _cac ,_cdfe ,_fab ,_dgd ,_acga ,_ccg ,_abf .SpPr ,nil ;};func (_caed *creatorContext )drawBorderWithProps (_ceef *_ab .CT_ShapeProperties ,_fgff *Rectangle ,_bdccb float64 ){if _fgff !=nil &&_ceef !=nil &&_ceef .Ln !=nil &&_ceef .Ln .SolidFill !=nil {_eagf :=_caed .getPdfColorFromSolidFill (_ceef .Ln .SolidFill );DrawRectangle (_caed ._ecf ,_fgff ,_bdccb ,_eagf );};};func DrawLine (c *_ag .Creator ,x0 ,y0 ,x1 ,y1 ,width float64 ,color _ag .Color ){if color ==nil {return ;};_gacc :=c .NewLine (x0 ,y0 ,x1 ,y1 );_gacc .SetLineWidth (width );_gacc .SetColor (color );c .Draw (_gacc );};func CropImageByRect (sourceImg _c .Image ,rect _c .Rectangle )_c .Image {_aggd ,_bbcc ,_afa ,_cfee :=rect .Min .X ,rect .Min .Y ,rect .Max .X ,rect .Max .Y ;_fdg :=_c .NewNRGBA (_c .Rect (0,0,_afa -_aggd ,_cfee -_bbcc ));for _afbd :=_aggd ;_afbd < _afa ;_afbd ++{for _bdbd :=_bbcc ;_bdbd < _cfee ;_bdbd ++{_fdg .Set (_afbd -_aggd ,_bdbd -_bbcc ,sourceImg .At (_afbd ,_bdbd ));};};return _fdg ;};func _gggd (_edf *_cg .CT_ValAx )(uint32 ,_cg .ST_AxPos ,_cg .ST_TickMark ,_cg .ST_TickLblPos ,*_cg .CT_ChartLines ,uint32 ,*_ab .CT_ShapeProperties ,error ){var _caag ,_egeg uint32 ;var _cgea _cg .ST_AxPos ;var _fbc _cg .ST_TickMark ;var _dgg *_cg .CT_ChartLines ;var _gbef _cg .ST_TickLblPos ;if _edf .AxId ==nil {return _caag ,_cgea ,_fbc ,_gbef ,_dgg ,_egeg ,_edf .SpPr ,_gg .New ("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_caag =_edf .AxId .ValAttr ;};if _edf .AxPos ==nil {return _caag ,_cgea ,_fbc ,_gbef ,_dgg ,_egeg ,_edf .SpPr ,_gg .New ("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e");}else {_cgea =_edf .AxPos .ValAttr ;};if _edf .MajorTickMark !=nil {_fbc =_edf .MajorTickMark .ValAttr ;};if _edf .TickLblPos !=nil {_gbef =_edf .TickLblPos .ValAttr ;};if _edf .CrossAx ==nil {return _caag ,_cgea ,_fbc ,_gbef ,_dgg ,_egeg ,_edf .SpPr ,_gg .New ("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044");}else {_egeg =_edf .CrossAx .ValAttr ;};_dgg =_edf .MajorGridlines ;return _caag ,_cgea ,_fbc ,_gbef ,_dgg ,_egeg ,_edf .SpPr ,nil ;};func AssignStdFontByName (style _ag .TextStyle ,fontName string )*_fd .PdfFont {_dad :=_fd .StdFontName (fontName );return _fd .NewStandard14FontMustCompile (_dad );};var _baee =_cab (0.5);func GetImage (c *_ag .Creator ,goImg _c .Image ,imgHeight ,imgWidth ,left ,top ,dividerX ,dividerY float64 ,part ImgPart )(*_ag .Image ,error ){if goImg ==nil {return nil ,nil ;};_cgaed :=goImg .Bounds ().Size ();_gad :=_cgaed .X ;_fbe :=_cgaed .Y ;if dividerX !=0{dividerX =dividerX /imgWidth *float64 (_gad );};if dividerY !=0{dividerY =dividerY /imgHeight *float64 (_fbe );};var _bdbb _c .Rectangle ;switch part {case ImgPart_t :_bdbb =_c .Rect (0,0,_gad ,int (dividerY ));case ImgPart_b :_bdbb =_c .Rect (0,int (dividerY ),_gad ,_fbe );case ImgPart_l :_bdbb =_c .Rect (0,0,int (dividerX ),_fbe );case ImgPart_r :_bdbb =_c .Rect (int (dividerX ),0,_gad ,_fbe );case ImgPart_lt :_bdbb =_c .Rect (0,0,int (dividerX ),int (dividerY ));case ImgPart_rt :_bdbb =_c .Rect (int (dividerX ),0,_gad ,int (dividerY ));case ImgPart_lb :_bdbb =_c .Rect (0,int (dividerY ),int (dividerX ),_fbe );case ImgPart_rb :_bdbb =_c .Rect (int (dividerX ),int (dividerY ),_gad ,_fbe );default:_bdbb =_c .Rect (0,0,_gad ,_fbe );};_bcdf :=CropImageByRect (goImg ,_bdbb );_gdae ,_dade :=c .NewImageFromGoImage (_bcdf );if _dade !=nil {return nil ,_dade ;};_gdae .Scale (imgWidth /float64 (_gad ),imgHeight /float64 (_fbe ));_gdae .SetPos (left ,top );return _gdae ,nil ;};func FromSTPercentage (st *_ab .ST_Percentage )float64 {if _dcac :=st .ST_PercentageDecimal ;_dcac !=nil {return float64 (*_dcac )/100000;};return 0;};func _dafg (_fgdd _ab .ST_SchemeColorVal ,_eag *_ab .Theme )string {if _aabg :=_eag .ThemeElements ;_aabg !=nil {if _eecf :=_aabg .ClrScheme ;_eecf !=nil {switch _fgdd {case _ab .ST_SchemeColorValLt1 :return GetColorStringFromDmlColor (_eecf .Lt1 );case _ab .ST_SchemeColorValDk1 ,_ab .ST_SchemeColorValTx1 :return GetColorStringFromDmlColor (_eecf .Dk1 );case _ab .ST_SchemeColorValLt2 :return GetColorStringFromDmlColor (_eecf .Lt2 );case _ab .ST_SchemeColorValDk2 :return GetColorStringFromDmlColor (_eecf .Dk2 );case _ab .ST_SchemeColorValAccent1 :return GetColorStringFromDmlColor (_eecf .Accent1 );case _ab .ST_SchemeColorValAccent2 :return GetColorStringFromDmlColor (_eecf .Accent2 );case _ab .ST_SchemeColorValAccent3 :return GetColorStringFromDmlColor (_eecf .Accent3 );case _ab .ST_SchemeColorValAccent4 :return GetColorStringFromDmlColor (_eecf .Accent4 );case _ab .ST_SchemeColorValAccent5 :return GetColorStringFromDmlColor (_eecf .Accent5 );case _ab .ST_SchemeColorValAccent6 :return GetColorStringFromDmlColor (_eecf .Accent6 );};};};return "";};var _efe =_cab (1.5);func (_gbf FontStyle )String ()string {return []string {"\u0052e\u0067\u0075\u006c\u0061\u0072","\u0042\u006f\u006c\u0064","\u0049\u0074\u0061\u006c\u0069\u0063","\u0042\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063"}[int (_gbf )];};func _gaf (_bfaa _ab .ST_SchemeColorVal ,_afg *_ab .Theme )string {if _bda :=_afg .ThemeElements ;_bda !=nil {if _ecac :=_bda .ClrScheme ;_ecac !=nil {switch _bfaa {case _ab .ST_SchemeColorValLt1 :return GetColorStringFromDmlColor (_ecac .Lt1 );case _ab .ST_SchemeColorValDk1 ,_ab .ST_SchemeColorValTx1 :return GetColorStringFromDmlColor (_ecac .Dk1 );case _ab .ST_SchemeColorValLt2 :return GetColorStringFromDmlColor (_ecac .Lt2 );case _ab .ST_SchemeColorValDk2 :return GetColorStringFromDmlColor (_ecac .Dk2 );case _ab .ST_SchemeColorValAccent1 :return GetColorStringFromDmlColor (_ecac .Accent1 );case _ab .ST_SchemeColorValAccent2 :return GetColorStringFromDmlColor (_ecac .Accent2 );case _ab .ST_SchemeColorValAccent3 :return GetColorStringFromDmlColor (_ecac .Accent3 );case _ab .ST_SchemeColorValAccent4 :return GetColorStringFromDmlColor (_ecac .Accent4 );case _ab .ST_SchemeColorValAccent5 :return GetColorStringFromDmlColor (_ecac .Accent5 );case _ab .ST_SchemeColorValAccent6 :return GetColorStringFromDmlColor (_ecac .Accent6 );};};};return "";};type serCategory struct{_bd string ;_bgc []serValue ;};type ImgPart byte ;var _cggg =map[string ]FontStyle {"\u0052e\u0067\u0075\u006c\u0061\u0072":FontStyle_Regular ,"\u0042\u006f\u006c\u0064":FontStyle_Bold ,"\u0049\u0074\u0061\u006c\u0069\u0063":FontStyle_Italic ,"B\u006f\u006c\u0064\u0020\u0049\u0074\u0061\u006c\u0069\u0063":FontStyle_BoldItalic };type legendItem struct{_gfe string ;_def *_ab .CT_ShapeProperties ;};func (_eec barSerByOrder )Len ()int {return len (_eec )};func AdjustColorByTint (colorStr string ,tint float64 )string {var _cdc ,_afeeb ,_fce uint8 ;_gbb ,_ :=_eb .Sscanf (colorStr ,"\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078",&_cdc ,&_afeeb ,&_fce );if _gbb !=3{return "";};return _ecc (_cdc ,tint )+_ecc (_afeeb ,tint )+_ecc (_fce ,tint );};func AdjustColor (colorStr string ,EG_ColorTransform []*_ab .EG_ColorTransform )string {for _ ,_faf :=range EG_ColorTransform {if _fca :=_faf .Tint ;_fca !=nil {if _gbg :=_fca .ValAttr .ST_PositiveFixedPercentageDecimal ;_gbg !=nil {colorStr =AdjustColorByTint (colorStr ,float64 (*_gbg )/100000);};};if _afee :=_faf .Shade ;_afee !=nil {if _gdc :=_afee .ValAttr .ST_PositiveFixedPercentageDecimal ;_gdc !=nil {colorStr =AdjustColorByShade (colorStr ,float64 (*_gdc )/100000);};};if _age :=_faf .LumMod ;_age !=nil {if _fcdf :=_age .ValAttr .ST_PercentageDecimal ;_fcdf !=nil {colorStr =AdjustColorByLumMod (colorStr ,float64 (*_fcdf )/100000);};};};return colorStr ;};func PointsFromTwips (twips int64 )float64 {return float64 (int64 (float64 (twips )*_gd .Twips *10+0.5))/10;};var StdFontsMap =map[string ][]string {"\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a":[]string {"\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a","\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064","\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065","H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065"},"\u0043o\u0075\u0072\u0069\u0065\u0072":[]string {"\u0043o\u0075\u0072\u0069\u0065\u0072","\u0043\u006f\u0075r\u0069\u0065\u0072\u002d\u0042\u006f\u006c\u0064","\u0043o\u0075r\u0069\u0065\u0072\u002d\u004f\u0062\u006c\u0069\u0071\u0075\u0065","\u0043\u006f\u0075\u0072ie\u0072\u002d\u0042\u006f\u006c\u0064\u004f\u0062\u006c\u0069\u0071\u0075\u0065"},"\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e":[]string {"T\u0069\u006d\u0065\u0073\u002d\u0052\u006f\u006d\u0061\u006e","\u0054\u0069\u006d\u0065\u0073\u002d\u0042\u006f\u006c\u0064","\u0054\u0069\u006de\u0073\u002d\u0049\u0074\u0061\u006c\u0069\u0063","\u0054\u0069m\u0065\u0073\u002dB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063"},"\u0064e\u0066\u0061\u0075\u006c\u0074":[]string {"\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a","\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064","\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065","H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065"}};func GetDataFromXfrm (xfrm *_ab .CT_Transform2D )(float64 ,float64 ,float64 ,float64 ){var _afca ,_aaag ,_fdgg ,_dbac float64 ;if _ebge :=xfrm .Off ;_ebge !=nil {_afca =_gd .FromEMU (FromSTCoordinate (_ebge .XAttr ));_aaag =_gd .FromEMU (FromSTCoordinate (_ebge .YAttr ));};if _agfe :=xfrm .Ext ;_agfe !=nil {_fdgg =_gd .FromEMU (_agfe .CxAttr );_dbac =_gd .FromEMU (_agfe .CyAttr );};return _afca ,_aaag ,_fdgg ,_dbac ;};func DrawRectangle (c *_ag .Creator ,r *Rectangle ,w float64 ,color _ag .Color ){if color ==nil {return ;};DrawLine (c ,r .Left ,r .Top ,r .Right ,r .Top ,w ,color );DrawLine (c ,r .Left ,r .Top ,r .Left ,r .Bottom ,w ,color );DrawLine (c ,r .Left ,r .Bottom ,r .Right ,r .Bottom ,w ,color );DrawLine (c ,r .Right ,r .Top ,r .Right ,r .Bottom ,w ,color );};func MakeTempCreator (width ,height float64 )*_ag .Creator {_fba :=_ag .New ();_fba .SetPageSize (_ag .PageSize {width ,height });_fba .SetPageMargins (0,0,0,0);return _fba ;};func TwipsFromPoints (points float64 )float64 {return points /_gd .Twips };func RegisterFont (name string ,style FontStyle ,font *_fd .PdfFont ){_ebe ._dge .Lock ();if _ebe ._ccge [name ]==nil {_ebe ._ccge [name ]=map[FontStyle ]*_fd .PdfFont {};};_ebe ._ccge [name ][style ]=font ;_ebe ._dge .Unlock ();};type FontStyle byte ;func (_ecab *Rectangle )Translate (x ,y float64 ){_ecab .Left +=x ;_ecab .Right +=x ;_ecab .Top +=y ;_ecab .Bottom +=y ;};const DefaultFontSize =12.0;func RegisterFontsFromDirectory (dirName string )error {_gff ,_abb :=_e .Open (dirName );if _abb !=nil {return _abb ;};defer _gff .Close ();_gda ,_abb :=_gff .Readdirnames (0);if _abb !=nil {return _abb ;};for _ ,_fbf :=range _gda {if _f .HasSuffix (_fbf ,"\u002e\u0074\u0074\u0066"){_aadb :=dirName +"\u002f"+_fbf ;_ebg ,_bgba :=_cc .ParseFile (_aadb );if _bgba !=nil {_bc .Log .Debug ("\u0043a\u006e\u006e\u006f\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020T\u0054\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",_bgba );continue ;};_abee :=_ebg .GetNameByID (1);if _abee ==""{_bc .Log .Debug ("\u004e\u006f\u0020\u0066\u006fn\u0074\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u0069\u006e\u0066\u006fr\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",_aadb );continue ;};_aabdc :=_ebg .GetNameByID (2);if _aabdc ==""{_bc .Log .Debug ("N\u006f\u0020\u0073\u0074\u0079\u006ce\u0020\u0069\u006e\u0066\u006f\u0072m\u0061\u0074\u0069\u006f\u006e\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020%\u0073",_aadb );continue ;};_eebd ,_bgba :=_fd .NewCompositePdfFontFromTTFFile (_aadb );if _bgba !=nil {_bc .Log .Debug ("C\u0061\u006e\u006e\u006f\u0074\u0020m\u0061\u006b\u0065\u0020\u0061\u0020f\u006f\u006e\u0074\u0020\u0066\u0072\u006fm\u0020\u0054\u0054\u0046\u0020\u0066\u0069\u006c\u0065\u0020%\u0073",_bgba );continue ;};RegisterFont (_abee ,_cggg [_aabdc ],_eebd );};};return nil ;};var _gge =_cab (0.125);func MakeBlockFromChartSpace (cs *_cg .ChartSpace ,width ,height float64 ,theme *_ab .Theme )(*_ag .Block ,error ){_eff ,_aaf :=_bcd (cs ,width ,height ,theme ,false );if _aaf !=nil {return nil ,_aaf ;};_bab ,_aaf :=GetPageFromCreator (_eff );if _aaf !=nil {return nil ,_aaf ;};_agd ,_aaf :=_ag .NewBlockFromPage (_bab );if _aaf !=nil {return nil ,_aaf ;};return _agd ,nil ;};func MakeBlockFromCreator (c *_ag .Creator )(*_ag .Block ,error ){_cae ,_gfab :=GetPageFromCreator (c );if _gfab !=nil {return nil ,_gfab ;};_bebd ,_gfab :=_ag .NewBlockFromPage (_cae );if _gfab !=nil {return nil ,_gfab ;};return _bebd ,nil ;};func (_da *creatorContext )drawBarChart (_ad *_cg .CT_BarChart ,_af *Rectangle ,_fa *_cg .CT_PlotAreaChoice1 )([]*legendItem ,error ){var _ff bool ;if _fe :=_ad .BarDir ;_fe !=nil {_ff =_fe .ValAttr ==_cg .ST_BarDirBar ;};_ec :=_ad .Ser ;_cf .Sort (barSerByOrder (_ec ));_de :=map[string ]serCategory {};_fb :=[]string {};_aff :=[]*legendItem {};_ba :=_d .Inf (1);_daf :=_d .Inf (-1);for _ ,_dd :=range _ec {var _cfg string ;if _ee :=_dd .Tx ;_ee !=nil {if _be :=_ee .Choice ;_be !=nil {if _be .V !=nil {_cfg =*_be .V ;}else if _eg :=_be .StrRef ;_eg !=nil {if _gdg :=_eg .StrCache ;_gdg !=nil {for _ ,_ea :=range _gdg .Pt {_cfg =_ea .V ;};};};};};if _db :=_dd .Cat ;_db !=nil {if _afc :=_db .Choice ;_afc !=nil {if _abe :=_afc .StrRef ;_abe !=nil {if _dc :=_abe .StrCache ;_dc !=nil {for _ ,_dcd :=range _dc .Pt {_bag :=_dcd .V ;if _ ,_ed :=_de [_bag ];!_ed {_de [_bag ]=serCategory {_bd :_bag ,_bgc :[]serValue {}};_fb =append (_fb ,_bag );};};};}else if _dea :=_afc .NumRef ;_dea !=nil {if _baa :=_dea .NumCache ;_baa !=nil {var _ga string ;if _baa .FormatCode !=nil {_ga =*_baa .FormatCode ;};for _ ,_afff :=range _baa .Pt {var _eee string ;if _afff .FormatCodeAttr ==nil {_eee =_ga ;}else {_eee =*_afff .FormatCodeAttr ;};var _aa string ;_bf ,_deac :=_a .ParseFloat (_afff .V ,64);if _deac !=nil {_aa =_afff .V ;}else {_aa =_fg .Number (_bf ,_eee );};if _ ,_bee :=_de [_aa ];!_bee {_de [_aa ]=serCategory {_bd :_aa ,_bgc :[]serValue {}};_fb =append (_fb ,_aa );};};};};};};if _bb :=_dd .Val ;_bb !=nil {if _fbd :=_bb .Choice ;_fbd !=nil {if _ae :=_fbd .NumRef ;_ae !=nil {if _eac :=_ae .NumCache ;_eac !=nil {for _gbe ,_aga :=range _eac .Pt {_ccc ,_afb :=_a .ParseFloat (_aga .V ,64);if _afb !=nil {_ccc =0;_bc .Log .Debug ("\u0070a\u0072s\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073",_afb );};if _ccc > _daf {_daf =_ccc ;};if _ccc < _ba {_ba =_ccc ;};_ebc :=_de [_fb [_gbe ]];_ebc ._bgc =append (_ebc ._bgc ,serValue {_aac :_cfg ,_fgfc :_ccc ,_adf :_dd .SpPr });_de [_fb [_gbe ]]=_ebc ;};};};};};_aff =append (_aff ,&legendItem {_gfe :_cfg ,_def :_dd .SpPr });};var _ege float64 ;var _eaa ,_aab float64 ;if _daf ==0&&_ba ==0{_ege =0.2;_aab =0;_eaa =1;}else {var _dbb float64 ;if _cff :=_d .Abs (_ba );_daf < _cff {_dbb =_cff ;}else {_dbb =_daf ;};_fag :=_d .Pow (10,_d .Floor (_d .Log10 (_dbb )));_agb :=_dbb /_fag ;if _agb >=1.715&&_agb < 4.29{_ege =0.5;}else if _agb >=4.29&&_agb < 8.58{_ege =1;}else {_ege =2;};_ege *=_fag ;if _daf <=0{_eaa =0;}else {_eaa =(_d .Ceil (_daf /_ege )+1)*_ege ;};if _ba >=0{_aab =0;}else {_aab =(_d .Floor (_ba /_ege )-1)*_ege ;};};_cd :=_da .drawAxes (_fa ,_aab ,_eaa ,_ege ,_fb ,_af ,_ff );if _cd !=nil {return nil ,_cd ;};_fda :=0.0;if _ad .GapWidth !=nil {if _aea :=_ad .GapWidth .ValAttr ;_aea !=nil {if _aeac :=_aea .ST_GapAmountUShort ;_aeac !=nil {_fda =float64 (*_aeac )/100.0;};};};_cgfe :=_af .Right -_af .Left ;_beb :=_af .Bottom -_af .Top ;_agbg :=float64 (len (_fb ));if _ff {_ef :=_eaa /(_eaa -_aab )*_cgfe ;_efd :=-_aab /(_eaa -_aab )*_cgfe ;_dcg :=_af .Left +_efd ;_dce :=_beb /_agbg ;for _dff ,_bbc :=range _fb {_eeb :=_de [_bbc ];_cbb :=float64 (len (_eeb ._bgc ))+_fda ;_eca :=_dce /_cbb ;_ecdc :=_eca *_fda ;_deb :=_af .Bottom -float64 (_dff )*_dce -_ecdc /2-_eca ;for _ ,_fgf :=range _eeb ._bgc {if _fgf ._fgfc ==0{continue ;};var _aabb ,_afcd float64 ;if _fgf ._fgfc > 0{_afcd =_fgf ._fgfc /_eaa *_ef ;_da .drawRectangleWithProps (_fgf ._adf ,_dcg ,_deb ,_afcd ,_eca ,false );}else {_afcd =_fgf ._fgfc /_aab *_efd ;_aabb =_dcg -_afcd ;_da .drawRectangleWithProps (_fgf ._adf ,_aabb ,_deb ,_afcd ,_eca ,false );};_deb -=_eca ;};};}else {_bae :=_eaa /(_eaa -_aab )*_beb ;_fea :=-_aab /(_eaa -_aab )*_beb ;_ce :=_af .Top +_bae ;_fc :=_cgfe /_agbg ;for _gc ,_eeg :=range _fb {_ebd :=_de [_eeg ];_gbc :=float64 (len (_ebd ._bgc ))+_fda ;_ge :=_fc /_gbc ;_feb :=_ge *_fda ;_bfa :=_af .Left +float64 (_gc )*_fc +_feb /2;for _ ,_egb :=range _ebd ._bgc {var _ade ,_ede float64 ;if _egb ._fgfc > 0{_ede =_egb ._fgfc /_eaa *_bae ;_ade =_ce -_ede ;_da .drawRectangleWithProps (_egb ._adf ,_bfa ,_ade ,_ge ,_ede ,false );}else {_ede =_egb ._fgfc /_aab *_fea ;_da .drawRectangleWithProps (_egb ._adf ,_bfa ,_ce ,_ge ,_ede ,false );};_bfa +=_ge ;};};};return _aff ,nil ;};func Lighten (clr float64 )float64 {return 0.6+0.4*clr };func GetOpacityFromColorTransform (trs []*_ab .EG_ColorTransform )float64 {for _ ,_gaeb :=range trs {if _gaeb !=nil {if _abed :=_gaeb .Alpha ;_abed !=nil {if _cebg :=_abed .ValAttr .ST_PositiveFixedPercentageDecimal ;_cebg !=nil {return float64 (*_cebg )/100000;};};};};return 1.0;};
