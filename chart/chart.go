
package chart ;import (_ce "fmt";_ca "gitee.com/gooffice/gooffice";_d "gitee.com/gooffice/gooffice/color";_eb "gitee.com/gooffice/gooffice/drawing";_cf "gitee.com/gooffice/gooffice/measurement";_e "gitee.com/gooffice/gooffice/schema/soo/dml";_c "gitee.com/gooffice/gooffice/schema/soo/dml/chart";_fd "math/rand";);func MakeValueAxis (x *_c .CT_ValAx )ValueAxis {return ValueAxis {x }};

// Properties returns the bar chart series shape properties.
func (_fbg BarChartSeries )Properties ()_eb .ShapeProperties {if _fbg ._cff .SpPr ==nil {_fbg ._cff .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_fbg ._cff .SpPr );};func (_bbfe ValueAxis )MajorGridLines ()GridLines {if _bbfe ._efce .MajorGridlines ==nil {_bbfe ._efce .MajorGridlines =_c .NewCT_ChartLines ();};return GridLines {_bbfe ._efce .MajorGridlines };};

// X returns the inner wrapped XML type.
func (_fda Area3DChart )X ()*_c .CT_Area3DChart {return _fda ._b };

// AddSeries adds a default series to an area chart.
func (_ge AreaChart )AddSeries ()AreaChartSeries {_bbc :=_ge .nextColor (len (_ge ._cb .Ser ));_fe :=_c .NewCT_AreaSer ();_ge ._cb .Ser =append (_ge ._cb .Ser ,_fe );_fe .Idx .ValAttr =uint32 (len (_ge ._cb .Ser )-1);_fe .Order .ValAttr =uint32 (len (_ge ._cb .Ser )-1);_fg :=AreaChartSeries {_fe };_fg .InitializeDefaults ();_fg .Properties ().SetSolidFill (_bbc );return _fg ;};func (_abb NumberDataSource )ensureChoice (){if _abb ._aeg .Choice ==nil {_abb ._aeg .Choice =_c .NewCT_NumDataSourceChoice ();};};func (_cef BarChart )AddAxis (axis Axis ){_acfg :=_c .NewCT_UnsignedInt ();_acfg .ValAttr =axis .AxisID ();_cef ._gcd .AxId =append (_cef ._gcd .AxId ,_acfg );};

// AddSeries adds a default series to a Bubble chart.
func (_cbfg BubbleChart )AddSeries ()BubbleChartSeries {_dbb :=_cbfg .nextColor (len (_cbfg ._eed .Ser ));_bf :=_c .NewCT_BubbleSer ();_cbfg ._eed .Ser =append (_cbfg ._eed .Ser ,_bf );_bf .Idx .ValAttr =uint32 (len (_cbfg ._eed .Ser )-1);_bf .Order .ValAttr =uint32 (len (_cbfg ._eed .Ser )-1);_dd :=BubbleChartSeries {_bf };_dd .InitializeDefaults ();_dd .Properties ().SetSolidFill (_dbb );return _dd ;};func (_afe CategoryAxis )SetTickLabelPosition (p _c .ST_TickLblPos ){if p ==_c .ST_TickLblPosUnset {_afe ._bcf .TickLblPos =nil ;}else {_afe ._bcf .TickLblPos =_c .NewCT_TickLblPos ();_afe ._bcf .TickLblPos .ValAttr =p ;};};

// Properties returns the bar chart series shape properties.
func (_afec RadarChartSeries )Properties ()_eb .ShapeProperties {if _afec ._aefed .SpPr ==nil {_afec ._aefed .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_afec ._aefed .SpPr );};

// Values returns the bubble size data source.
func (_fff BubbleChartSeries )BubbleSizes ()NumberDataSource {if _fff ._cba .BubbleSize ==nil {_fff ._cba .BubbleSize =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_fff ._cba .BubbleSize );};func (_fdfc DateAxis )SetPosition (p _c .ST_AxPos ){_fdfc ._dea .AxPos =_c .NewCT_AxPos ();_fdfc ._dea .AxPos .ValAttr =p ;};func (_cfag DateAxis )Properties ()_eb .ShapeProperties {if _cfag ._dea .SpPr ==nil {_cfag ._dea .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_cfag ._dea .SpPr );};

// InitializeDefaults the bar chart to its defaults
func (_ba AreaChart )InitializeDefaults (){};func (_ddbc chartBase )nextColor (_bcb int )_d .Color {return _ebge [_bcb %len (_ebge )]};func MakeChart (x *_c .ChartSpace )Chart {return Chart {x }};

// Marker returns the marker properties.
func (_bgeg ScatterChartSeries )Marker ()Marker {if _bgeg ._bebd .Marker ==nil {_bgeg ._bebd .Marker =_c .NewCT_Marker ();};return MakeMarker (_bgeg ._bebd .Marker );};

// AddPieOfPieChart adds a new pie chart to a chart.
func (_gdg Chart )AddPieOfPieChart ()PieOfPieChart {_ded :=_c .NewCT_PlotAreaChoice ();_gdg ._egd .Chart .PlotArea .Choice =append (_gdg ._egd .Chart .PlotArea .Choice ,_ded );_ded .OfPieChart =_c .NewCT_OfPieChart ();_aged :=PieOfPieChart {_abdb :_ded .OfPieChart };_aged .InitializeDefaults ();return _aged ;};

// AddSeries adds a default series to a bar chart.
func (_eg Bar3DChart )AddSeries ()BarChartSeries {_cbg :=_eg .nextColor (len (_eg ._a .Ser ));_dbe :=_c .NewCT_BarSer ();_eg ._a .Ser =append (_eg ._a .Ser ,_dbe );_dbe .Idx .ValAttr =uint32 (len (_eg ._a .Ser )-1);_dbe .Order .ValAttr =uint32 (len (_eg ._a .Ser )-1);_cgg :=BarChartSeries {_dbe };_cgg .InitializeDefaults ();_cgg .Properties ().SetSolidFill (_cbg );return _cgg ;};

// AreaChartSeries is a series to be used on an area chart.
type AreaChartSeries struct{_be *_c .CT_AreaSer };func (_cagf SurfaceChartSeries )Values ()NumberDataSource {if _cagf ._gaec .Val ==nil {_cagf ._gaec .Val =_c .NewCT_NumDataSource ();};_cgd :=MakeNumberDataSource (_cagf ._gaec .Val );_cgd .CreateEmptyNumberCache ();return _cgd ;};func (_feg BubbleChart )AddAxis (axis Axis ){_cad :=_c .NewCT_UnsignedInt ();_cad .ValAttr =axis .AxisID ();_feg ._eed .AxId =append (_feg ._eed .AxId ,_cad );};

// AddDoughnutChart adds a new doughnut (pie with a hole in the center) chart to a chart.
func (_bacf Chart )AddDoughnutChart ()DoughnutChart {_ffd :=_c .NewCT_PlotAreaChoice ();_bacf ._egd .Chart .PlotArea .Choice =append (_bacf ._egd .Chart .PlotArea .Choice ,_ffd );_ffd .DoughnutChart =_c .NewCT_DoughnutChart ();_cgb :=DoughnutChart {_cac :_ffd .DoughnutChart };_cgb .InitializeDefaults ();return _cgb ;};

// Labels returns the data label properties.
func (_ccag LineChartSeries )Labels ()DataLabels {if _ccag ._adb .DLbls ==nil {_ccag ._adb .DLbls =_c .NewCT_DLbls ();};return MakeDataLabels (_ccag ._adb .DLbls );};

// InitializeDefaults the bar chart to its defaults
func (_fee Bar3DChart )InitializeDefaults (){_fee .SetDirection (_c .ST_BarDirCol )};func (_aeba Marker )SetSize (sz uint8 ){_aeba ._afg .Size =_c .NewCT_MarkerSize ();_aeba ._afg .Size .ValAttr =&sz ;};

// AddBarChart adds a new bar chart to a chart.
func (_gfd Chart )AddBarChart ()BarChart {_eedc :=_c .NewCT_PlotAreaChoice ();_gfd ._egd .Chart .PlotArea .Choice =append (_gfd ._egd .Chart .PlotArea .Choice ,_eedc );_eedc .BarChart =_c .NewCT_BarChart ();_eedc .BarChart .Grouping =_c .NewCT_BarGrouping ();_eedc .BarChart .Grouping .ValAttr =_c .ST_BarGroupingStandard ;_agg :=BarChart {_gcd :_eedc .BarChart };_agg .InitializeDefaults ();return _agg ;};

// SetText sets the series text.
func (_adea PieChartSeries )SetText (s string ){_adea ._faag .Tx =_c .NewCT_SerTx ();_adea ._faag .Tx .Choice .V =&s ;};

// AddArea3DChart adds a new area chart to a chart.
func (_eda Chart )AddArea3DChart ()Area3DChart {_aab (_eda ._egd .Chart );_aff :=_c .NewCT_PlotAreaChoice ();_eda ._egd .Chart .PlotArea .Choice =append (_eda ._egd .Chart .PlotArea .Choice ,_aff );_aff .Area3DChart =_c .NewCT_Area3DChart ();_dbf :=Area3DChart {_b :_aff .Area3DChart };_dbf .InitializeDefaults ();return _dbf ;};func (_agd DataLabels )SetShowLegendKey (b bool ){_agd .ensureChoice ();_agd ._dfga .Choice .ShowLegendKey =_c .NewCT_Boolean ();_agd ._dfga .Choice .ShowLegendKey .ValAttr =_ca .Bool (b );};func (_de CategoryAxis )AxisID ()uint32 {return _de ._bcf .AxId .ValAttr };

// X returns the inner wrapped XML type.
func (_acg BarChart )X ()*_c .CT_BarChart {return _acg ._gcd };

// X returns the inner wrapped XML type.
func (_ad BubbleChart )X ()*_c .CT_BubbleChart {return _ad ._eed };

// Values returns the value data source.
func (_dbg BarChartSeries )Values ()NumberDataSource {if _dbg ._cff .Val ==nil {_dbg ._cff .Val =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_dbg ._cff .Val );};func MakeCategoryAxis (x *_c .CT_CatAx )CategoryAxis {return CategoryAxis {x }};func MakeLegend (l *_c .CT_Legend )Legend {return Legend {l }};

// AddBubbleChart adds a new bubble chart.
func (_bd Chart )AddBubbleChart ()BubbleChart {_cfa :=_c .NewCT_PlotAreaChoice ();_bd ._egd .Chart .PlotArea .Choice =append (_bd ._egd .Chart .PlotArea .Choice ,_cfa );_cfa .BubbleChart =_c .NewCT_BubbleChart ();_edg :=BubbleChart {_eed :_cfa .BubbleChart };_edg .InitializeDefaults ();return _edg ;};

// AddBar3DChart adds a new 3D bar chart to a chart.
func (_aaa Chart )AddBar3DChart ()Bar3DChart {_aab (_aaa ._egd .Chart );_gec :=_c .NewCT_PlotAreaChoice ();_aaa ._egd .Chart .PlotArea .Choice =append (_aaa ._egd .Chart .PlotArea .Choice ,_gec );_gec .Bar3DChart =_c .NewCT_Bar3DChart ();_gec .Bar3DChart .Grouping =_c .NewCT_BarGrouping ();_gec .Bar3DChart .Grouping .ValAttr =_c .ST_BarGroupingStandard ;_bcc :=Bar3DChart {_a :_gec .Bar3DChart };_bcc .InitializeDefaults ();return _bcc ;};

// LineChartSeries is the data series for a line chart.
type LineChartSeries struct{_adb *_c .CT_LineSer };func (_bfg CategoryAxis )SetCrosses (axis Axis ){_bfg ._bcf .Choice =_c .NewEG_AxSharedChoice ();_bfg ._bcf .Choice .Crosses =_c .NewCT_Crosses ();_bfg ._bcf .Choice .Crosses .ValAttr =_c .ST_CrossesAutoZero ;_bfg ._bcf .CrossAx .ValAttr =axis .AxisID ();};type DateAxis struct{_dea *_c .CT_DateAx };

// InitializeDefaults the bar chart to its defaults
func (_gab RadarChart )InitializeDefaults (){_gab ._abe .RadarStyle .ValAttr =_c .ST_RadarStyleMarker };func MakeMarker (x *_c .CT_Marker )Marker {return Marker {x }};

// Properties returns the line chart series shape properties.
func (_aed LineChartSeries )Properties ()_eb .ShapeProperties {if _aed ._adb .SpPr ==nil {_aed ._adb .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_aed ._adb .SpPr );};

// X returns the inner wrapped XML type.
func (_dgc Chart )X ()*_c .ChartSpace {return _dgc ._egd };func (_deac Title )ParagraphProperties ()_eb .ParagraphProperties {if _deac ._feb .Tx ==nil {_deac .SetText ("");};if _deac ._feb .Tx .Choice .Rich .P [0].PPr ==nil {_deac ._feb .Tx .Choice .Rich .P [0].PPr =_e .NewCT_TextParagraphProperties ();};return _eb .MakeParagraphProperties (_deac ._feb .Tx .Choice .Rich .P [0].PPr );};

// Properties returns the Bubble chart series shape properties.
func (_bg BubbleChartSeries )Properties ()_eb .ShapeProperties {if _bg ._cba .SpPr ==nil {_bg ._cba .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_bg ._cba .SpPr );};

// StockChart is a 2D Stock chart.
type StockChart struct{chartBase ;_cfe *_c .CT_StockChart ;};func (_cbe ScatterChartSeries )SetSmooth (b bool ){_cbe ._bebd .Smooth =_c .NewCT_Boolean ();_cbe ._bebd .Smooth .ValAttr =&b ;};func (_fcf Legend )Properties ()_eb .ShapeProperties {if _fcf ._eaf .SpPr ==nil {_fcf ._eaf .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_fcf ._eaf .SpPr );};

// Marker returns the marker properties.
func (_bff LineChartSeries )Marker ()Marker {if _bff ._adb .Marker ==nil {_bff ._adb .Marker =_c .NewCT_Marker ();};return MakeMarker (_bff ._adb .Marker );};var _ebge =[]_d .Color {_d .RGB (0x33,0x66,0xcc),_d .RGB (0xDC,0x39,0x12),_d .RGB (0xFF,0x99,0x00),_d .RGB (0x10,0x96,0x18),_d .RGB (0x99,0x00,0x99),_d .RGB (0x3B,0x3E,0xAC),_d .RGB (0x00,0x99,0xC6),_d .RGB (0xDD,0x44,0x77),_d .RGB (0x66,0xAA,0x00),_d .RGB (0xB8,0x2E,0x2E),_d .RGB (0x31,0x63,0x95),_d .RGB (0x99,0x44,0x99),_d .RGB (0x22,0xAA,0x99),_d .RGB (0xAA,0xAA,0x11),_d .RGB (0x66,0x33,0xCC),_d .RGB (0xE6,0x73,0x00),_d .RGB (0x8B,0x07,0x07),_d .RGB (0x32,0x92,0x62),_d .RGB (0x55,0x74,0xA6),_d .RGB (0x3B,0x3E,0xAC)};type SurfaceChartSeries struct{_gaec *_c .CT_SurfaceSer };func (_cg Area3DChart )AddAxis (axis Axis ){_ec :=_c .NewCT_UnsignedInt ();_ec .ValAttr =axis .AxisID ();_cg ._b .AxId =append (_cg ._b .AxId ,_ec );};func (_ae CategoryAxis )SetMajorTickMark (m _c .ST_TickMark ){if m ==_c .ST_TickMarkUnset {_ae ._bcf .MajorTickMark =nil ;}else {_ae ._bcf .MajorTickMark =_c .NewCT_TickMark ();_ae ._bcf .MajorTickMark .ValAttr =m ;};};

// RemoveTitle removes any existing title from the chart.
func (_bbe Chart )RemoveTitle (){_bbe ._egd .Chart .Title =nil ;_bbe ._egd .Chart .AutoTitleDeleted =_c .NewCT_Boolean ();_bbe ._egd .Chart .AutoTitleDeleted .ValAttr =_ca .Bool (true );};

// AddTitle sets a new title on the chart.
func (_ggg Chart )AddTitle ()Title {_ggg ._egd .Chart .Title =_c .NewCT_Title ();_ggg ._egd .Chart .Title .Overlay =_c .NewCT_Boolean ();_ggg ._egd .Chart .Title .Overlay .ValAttr =_ca .Bool (false );_ggg ._egd .Chart .AutoTitleDeleted =_c .NewCT_Boolean ();_ggg ._egd .Chart .AutoTitleDeleted .ValAttr =_ca .Bool (false );_cfb :=MakeTitle (_ggg ._egd .Chart .Title );_cfb .InitializeDefaults ();return _cfb ;};

// CategoryAxisDataSource specifies the data for an axis.  It's commonly used with
// SetReference to set the axis data to a range of cells.
type CategoryAxisDataSource struct{_ffg *_c .CT_AxDataSource };

// DoughnutChart is a Doughnut chart.
type DoughnutChart struct{chartBase ;_cac *_c .CT_DoughnutChart ;};

// AddAxis adds an axis to a Scatter chart.
func (_bgg ScatterChart )AddAxis (axis Axis ){_abbb :=_c .NewCT_UnsignedInt ();_abbb .ValAttr =axis .AxisID ();_bgg ._eeee .AxId =append (_bgg ._eeee .AxId ,_abbb );};

// AddSurface3DChart adds a new 3D surface chart to a chart.
func (_fde Chart )AddSurface3DChart ()Surface3DChart {_faa :=_c .NewCT_PlotAreaChoice ();_fde ._egd .Chart .PlotArea .Choice =append (_fde ._egd .Chart .PlotArea .Choice ,_faa );_faa .Surface3DChart =_c .NewCT_Surface3DChart ();_aab (_fde ._egd .Chart );_gfg :=Surface3DChart {_daeb :_faa .Surface3DChart };_gfg .InitializeDefaults ();return _gfg ;};func (_aggg RadarChart )AddAxis (axis Axis ){_dcb :=_c .NewCT_UnsignedInt ();_dcb .ValAttr =axis .AxisID ();_aggg ._abe .AxId =append (_aggg ._abe .AxId ,_dcb );};

// SetOrder sets the order of the series
func (_aec SurfaceChartSeries )SetOrder (idx uint32 ){_aec ._gaec .Order .ValAttr =idx };type chartBase struct{};

// AddAxis adds an axis to a Surface chart.
func (_bgegf Surface3DChart )AddAxis (axis Axis ){_aaf :=_c .NewCT_UnsignedInt ();_aaf .ValAttr =axis .AxisID ();_bgegf ._daeb .AxId =append (_bgegf ._daeb .AxId ,_aaf );};func (_dfb AreaChart )AddAxis (axis Axis ){_ga :=_c .NewCT_UnsignedInt ();_ga .ValAttr =axis .AxisID ();_dfb ._cb .AxId =append (_dfb ._cb .AxId ,_ga );};

// AddSeries adds a default series to an Doughnut chart.
func (_gaca DoughnutChart )AddSeries ()PieChartSeries {_abd :=_c .NewCT_PieSer ();_gaca ._cac .Ser =append (_gaca ._cac .Ser ,_abd );_abd .Idx .ValAttr =uint32 (len (_gaca ._cac .Ser )-1);_abd .Order .ValAttr =uint32 (len (_gaca ._cac .Ser )-1);_aeb :=PieChartSeries {_abd };_aeb .InitializeDefaults ();return _aeb ;};

// SetOrder sets the order of the series
func (_dbfd LineChartSeries )SetOrder (idx uint32 ){_dbfd ._adb .Order .ValAttr =idx };

// InitializeDefaults the bar chart to its defaults
func (_aef Pie3DChart )InitializeDefaults (){_aef ._eaca .VaryColors =_c .NewCT_Boolean ();_aef ._eaca .VaryColors .ValAttr =_ca .Bool (true );};func (_ggc GridLines )Properties ()_eb .ShapeProperties {if _ggc ._cbb .SpPr ==nil {_ggc ._cbb .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_ggc ._cbb .SpPr );};func (_aggf Marker )Properties ()_eb .ShapeProperties {if _aggf ._afg .SpPr ==nil {_aggf ._afg .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_aggf ._afg .SpPr );};

// Values returns the value data source.
func (_fb AreaChartSeries )Values ()NumberDataSource {if _fb ._be .Val ==nil {_fb ._be .Val =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_fb ._be .Val );};

// Order returns the order of the series
func (_ffc SurfaceChartSeries )Order ()uint32 {return _ffc ._gaec .Order .ValAttr };func (_gcdf DataLabels )SetShowLeaderLines (b bool ){_gcdf .ensureChoice ();_gcdf ._dfga .Choice .ShowLeaderLines =_c .NewCT_Boolean ();_gcdf ._dfga .Choice .ShowLeaderLines .ValAttr =_ca .Bool (b );};

// AddSeries adds a default series to an Pie chart.
func (_fega PieChart )AddSeries ()PieChartSeries {_dcd :=_c .NewCT_PieSer ();_fega ._ggdc .Ser =append (_fega ._ggdc .Ser ,_dcd );_dcd .Idx .ValAttr =uint32 (len (_fega ._ggdc .Ser )-1);_dcd .Order .ValAttr =uint32 (len (_fega ._ggdc .Ser )-1);_ddg :=PieChartSeries {_dcd };_ddg .InitializeDefaults ();return _ddg ;};type SeriesAxis struct{_cga *_c .CT_SerAx };func (_bfd SeriesAxis )AxisID ()uint32 {return _bfd ._cga .AxId .ValAttr };

// AddRadarChart adds a new radar chart to a chart.
func (_baf Chart )AddRadarChart ()RadarChart {_fgc :=_c .NewCT_PlotAreaChoice ();_baf ._egd .Chart .PlotArea .Choice =append (_baf ._egd .Chart .PlotArea .Choice ,_fgc );_fgc .RadarChart =_c .NewCT_RadarChart ();_bbb :=RadarChart {_abe :_fgc .RadarChart };_bbb .InitializeDefaults ();return _bbb ;};

// InitializeDefaults initializes an Radar series to the default values.
func (_eafd RadarChartSeries )InitializeDefaults (){};

// X returns the inner wrapped XML type.
func (_db Bar3DChart )X ()*_c .CT_Bar3DChart {return _db ._a };

// SetLabelReference is used to set the source data to a range of cells
// containing strings.
func (_ddb CategoryAxisDataSource )SetLabelReference (s string ){_ddb ._ffg .Choice =_c .NewCT_AxDataSourceChoice ();_ddb ._ffg .Choice .StrRef =_c .NewCT_StrRef ();_ddb ._ffg .Choice .StrRef .F =s ;};

// Values returns the value data source.
func (_beb RadarChartSeries )Values ()NumberDataSource {if _beb ._aefed .Val ==nil {_beb ._aefed .Val =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_beb ._aefed .Val );};

// AddValueAxis adds a value axis to the chart.
func (_ddba Chart )AddValueAxis ()ValueAxis {_efab :=_c .NewCT_ValAx ();if _ddba ._egd .Chart .PlotArea .CChoice ==nil {_ddba ._egd .Chart .PlotArea .CChoice =_c .NewCT_PlotAreaChoice1 ();};_efab .AxId =_c .NewCT_UnsignedInt ();_efab .AxId .ValAttr =0x7FFFFFFF&_fd .Uint32 ();_ddba ._egd .Chart .PlotArea .CChoice .ValAx =append (_ddba ._egd .Chart .PlotArea .CChoice .ValAx ,_efab );_efab .Delete =_c .NewCT_Boolean ();_efab .Delete .ValAttr =_ca .Bool (false );_efab .Scaling =_c .NewCT_Scaling ();_efab .Scaling .Orientation =_c .NewCT_Orientation ();_efab .Scaling .Orientation .ValAttr =_c .ST_OrientationMinMax ;_efab .Choice =&_c .EG_AxSharedChoice {};_efab .Choice .Crosses =_c .NewCT_Crosses ();_efab .Choice .Crosses .ValAttr =_c .ST_CrossesAutoZero ;_efab .CrossBetween =_c .NewCT_CrossBetween ();_efab .CrossBetween .ValAttr =_c .ST_CrossBetweenBetween ;_bgf :=MakeValueAxis (_efab );_bgf .MajorGridLines ().Properties ().LineProperties ().SetSolidFill (_d .LightGray );_bgf .SetMajorTickMark (_c .ST_TickMarkOut );_bgf .SetMinorTickMark (_c .ST_TickMarkIn );_bgf .SetTickLabelPosition (_c .ST_TickLblPosNextTo );_bgf .Properties ().LineProperties ().SetSolidFill (_d .Black );_bgf .SetPosition (_c .ST_AxPosL );return _bgf ;};

// RadarChartSeries is a series to be used on an Radar chart.
type RadarChartSeries struct{_aefed *_c .CT_RadarSer };

// Area3DChart is an area chart that has a shaded area underneath a curve.
type Area3DChart struct{chartBase ;_b *_c .CT_Area3DChart ;};

// CategoryAxis returns the category data source.
func (_cbf AreaChartSeries )CategoryAxis ()CategoryAxisDataSource {if _cbf ._be .Cat ==nil {_cbf ._be .Cat =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_cbf ._be .Cat );};

// Properties returns the line chart series shape properties.
func (_ead ScatterChartSeries )Properties ()_eb .ShapeProperties {if _ead ._bebd .SpPr ==nil {_ead ._bebd .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_ead ._bebd .SpPr );};

// SetDisplayBlanksAs controls how missing values are displayed.
func (_fdaf Chart )SetDisplayBlanksAs (v _c .ST_DispBlanksAs ){_fdaf ._egd .Chart .DispBlanksAs =_c .NewCT_DispBlanksAs ();_fdaf ._egd .Chart .DispBlanksAs .ValAttr =v ;};

// AddDateAxis adds a value axis to the chart.
func (_ebgb Chart )AddDateAxis ()DateAxis {_cea :=_c .NewCT_DateAx ();if _ebgb ._egd .Chart .PlotArea .CChoice ==nil {_ebgb ._egd .Chart .PlotArea .CChoice =_c .NewCT_PlotAreaChoice1 ();};_cea .AxId =_c .NewCT_UnsignedInt ();_cea .AxId .ValAttr =0x7FFFFFFF&_fd .Uint32 ();_ebgb ._egd .Chart .PlotArea .CChoice .DateAx =append (_ebgb ._egd .Chart .PlotArea .CChoice .DateAx ,_cea );_cea .Delete =_c .NewCT_Boolean ();_cea .Delete .ValAttr =_ca .Bool (false );_cea .Scaling =_c .NewCT_Scaling ();_cea .Scaling .Orientation =_c .NewCT_Orientation ();_cea .Scaling .Orientation .ValAttr =_c .ST_OrientationMinMax ;_cea .Choice =&_c .EG_AxSharedChoice {};_cea .Choice .Crosses =_c .NewCT_Crosses ();_cea .Choice .Crosses .ValAttr =_c .ST_CrossesAutoZero ;_eaa :=DateAxis {_cea };_eaa .MajorGridLines ().Properties ().LineProperties ().SetSolidFill (_d .LightGray );_eaa .SetMajorTickMark (_c .ST_TickMarkOut );_eaa .SetMinorTickMark (_c .ST_TickMarkIn );_eaa .SetTickLabelPosition (_c .ST_TickLblPosNextTo );_eaa .Properties ().LineProperties ().SetSolidFill (_d .Black );_eaa .SetPosition (_c .ST_AxPosL );return _eaa ;};

// X returns the inner wrapped XML type.
func (_bcd Legend )X ()*_c .CT_Legend {return _bcd ._eaf };

// X returns the inner wrapped XML type.
func (_bef AreaChartSeries )X ()*_c .CT_AreaSer {return _bef ._be };

// InitializeDefaults initializes an Pie series to the default values.
func (_cde PieChartSeries )InitializeDefaults (){};func (_egb DateAxis )SetMinorTickMark (m _c .ST_TickMark ){if m ==_c .ST_TickMarkUnset {_egb ._dea .MinorTickMark =nil ;}else {_egb ._dea .MinorTickMark =_c .NewCT_TickMark ();_egb ._dea .MinorTickMark .ValAttr =m ;};};

// InitializeDefaults the Bubble chart to its defaults
func (_ebc BubbleChart )InitializeDefaults (){};type ScatterChart struct{chartBase ;_eeee *_c .CT_ScatterChart ;};

// SetHoleSize controls the hole size in the pie chart and is measured in percent.
func (_fag DoughnutChart )SetHoleSize (pct uint8 ){if _fag ._cac .HoleSize ==nil {_fag ._cac .HoleSize =_c .NewCT_HoleSize ();};if _fag ._cac .HoleSize .ValAttr ==nil {_fag ._cac .HoleSize .ValAttr =&_c .ST_HoleSize {};};_fag ._cac .HoleSize .ValAttr .ST_HoleSizeUByte =&pct ;};

// Chart is a generic chart.
type Chart struct{_egd *_c .ChartSpace };func (_dfg nullAxis )AxisID ()uint32 {return 0};

// AddSeries adds a default series to a bar chart.
func (_ee BarChart )AddSeries ()BarChartSeries {_cc :=_ee .nextColor (len (_ee ._gcd .Ser ));_gdd :=_c .NewCT_BarSer ();_ee ._gcd .Ser =append (_ee ._gcd .Ser ,_gdd );_gdd .Idx .ValAttr =uint32 (len (_ee ._gcd .Ser )-1);_gdd .Order .ValAttr =uint32 (len (_ee ._gcd .Ser )-1);_acf :=BarChartSeries {_gdd };_acf .InitializeDefaults ();_acf .Properties ().SetSolidFill (_cc );return _acf ;};

// RemoveLegend removes the legend if the chart has one.
func (_ada Chart )RemoveLegend (){_ada ._egd .Chart .Legend =nil };

// BarChart is a 2D bar chart.
type BarChart struct{chartBase ;_gcd *_c .CT_BarChart ;};

// Properties returns the line chart series shape properties.
func (_gabe SurfaceChartSeries )Properties ()_eb .ShapeProperties {if _gabe ._gaec .SpPr ==nil {_gabe ._gaec .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_gabe ._gaec .SpPr );};type NumberDataSource struct{_aeg *_c .CT_NumDataSource };func (_bcaa SeriesAxis )InitializeDefaults (){};

// Index returns the index of the series
func (_afff LineChartSeries )Index ()uint32 {return _afff ._adb .Idx .ValAttr };

// AddScatterChart adds a scatter (X/Y) chart.
func (_ebca Chart )AddScatterChart ()ScatterChart {_bbg :=_c .NewCT_PlotAreaChoice ();_ebca ._egd .Chart .PlotArea .Choice =append (_ebca ._egd .Chart .PlotArea .Choice ,_bbg );_bbg .ScatterChart =_c .NewCT_ScatterChart ();_faf :=ScatterChart {_eeee :_bbg .ScatterChart };_faf .InitializeDefaults ();return _faf ;};type Title struct{_feb *_c .CT_Title };

// SetExplosion sets the value that the segements of the pie are 'exploded' by
func (_cdf PieChartSeries )SetExplosion (v uint32 ){_cdf ._faag .Explosion =_c .NewCT_UnsignedInt ();_cdf ._faag .Explosion .ValAttr =v ;};func MakeDataLabels (x *_c .CT_DLbls )DataLabels {return DataLabels {x }};

// SetText sets the series text.
func (_cge BubbleChartSeries )SetText (s string ){_cge ._cba .Tx =_c .NewCT_SerTx ();_cge ._cba .Tx .Choice .V =&s ;};func (_gfc CategoryAxis )InitializeDefaults (){_gfc .SetPosition (_c .ST_AxPosB );_gfc .SetMajorTickMark (_c .ST_TickMarkOut );_gfc .SetMinorTickMark (_c .ST_TickMarkIn );_gfc .SetTickLabelPosition (_c .ST_TickLblPosNextTo );_gfc .MajorGridLines ().Properties ().LineProperties ().SetSolidFill (_d .LightGray );_gfc .Properties ().LineProperties ().SetSolidFill (_d .Black );};func (_dae DateAxis )AxisID ()uint32 {return _dae ._dea .AxId .ValAttr };

// Order returns the order of the series
func (_geba ScatterChartSeries )Order ()uint32 {return _geba ._bebd .Order .ValAttr };

// X returns the inner wrapped XML type.
func (_fa AreaChart )X ()*_c .CT_AreaChart {return _fa ._cb };

// AddAxis adds an axis to a line chart.
func (_efe LineChart )AddAxis (axis Axis ){_bca :=_c .NewCT_UnsignedInt ();_bca .ValAttr =axis .AxisID ();_efe ._bed .AxId =append (_efe ._bed .AxId ,_bca );};

// InitializeDefaults the bar chart to its defaults
func (_dbd PieOfPieChart )InitializeDefaults (){_dbd ._abdb .VaryColors =_c .NewCT_Boolean ();_dbd ._abdb .VaryColors .ValAttr =_ca .Bool (true );_dbd .SetType (_c .ST_OfPieTypePie );_dbd ._abdb .SecondPieSize =_c .NewCT_SecondPieSize ();_dbd ._abdb .SecondPieSize .ValAttr =&_c .ST_SecondPieSize {};_dbd ._abdb .SecondPieSize .ValAttr .ST_SecondPieSizeUShort =_ca .Uint16 (75);_gdfb :=_c .NewCT_ChartLines ();_gdfb .SpPr =_e .NewCT_ShapeProperties ();_aefe :=_eb .MakeShapeProperties (_gdfb .SpPr );_aefe .LineProperties ().SetSolidFill (_d .Auto );_dbd ._abdb .SerLines =append (_dbd ._abdb .SerLines ,_gdfb );};

// InitializeDefaults initializes a bar chart series to the default values.
func (_ea BarChartSeries )InitializeDefaults (){};

// X returns the inner wrapped XML type.
func (_ffgg PieChartSeries )X ()*_c .CT_PieSer {return _ffgg ._faag };func (_ffb Legend )InitializeDefaults (){_ffb .SetPosition (_c .ST_LegendPosR );_ffb .SetOverlay (false );_ffb .Properties ().SetNoFill ();_ffb .Properties ().LineProperties ().SetNoFill ();};

// X returns the inner wrapped XML type.
func (_eee Pie3DChart )X ()*_c .CT_Pie3DChart {return _eee ._eaca };

// AddLine3DChart adds a new 3D line chart to a chart.
func (_ag Chart )AddLine3DChart ()Line3DChart {_aab (_ag ._egd .Chart );_ebd :=_c .NewCT_PlotAreaChoice ();_ag ._egd .Chart .PlotArea .Choice =append (_ag ._egd .Chart .PlotArea .Choice ,_ebd );_ebd .Line3DChart =_c .NewCT_Line3DChart ();_ebd .Line3DChart .Grouping =_c .NewCT_Grouping ();_ebd .Line3DChart .Grouping .ValAttr =_c .ST_GroupingStandard ;return Line3DChart {_cadd :_ebd .Line3DChart };};

// InitializeDefaults the bar chart to its defaults
func (_bgfc DoughnutChart )InitializeDefaults (){_bgfc ._cac .VaryColors =_c .NewCT_Boolean ();_bgfc ._cac .VaryColors .ValAttr =_ca .Bool (true );_bgfc ._cac .HoleSize =_c .NewCT_HoleSize ();_bgfc ._cac .HoleSize .ValAttr =&_c .ST_HoleSize {};_bgfc ._cac .HoleSize .ValAttr .ST_HoleSizeUByte =_ca .Uint8 (50);};

// InitializeDefaults the Stock chart to its defaults
func (_dce StockChart )InitializeDefaults (){_dce ._cfe .HiLowLines =_c .NewCT_ChartLines ();_dce ._cfe .UpDownBars =_c .NewCT_UpDownBars ();_dce ._cfe .UpDownBars .GapWidth =_c .NewCT_GapAmount ();_dce ._cfe .UpDownBars .GapWidth .ValAttr =&_c .ST_GapAmount {};_dce ._cfe .UpDownBars .GapWidth .ValAttr .ST_GapAmountUShort =_ca .Uint16 (150);_dce ._cfe .UpDownBars .UpBars =_c .NewCT_UpDownBar ();_dce ._cfe .UpDownBars .DownBars =_c .NewCT_UpDownBar ();};func (_agf ValueAxis )AxisID ()uint32 {return _agf ._efce .AxId .ValAttr };func (_bac CategoryAxis )SetMinorTickMark (m _c .ST_TickMark ){if m ==_c .ST_TickMarkUnset {_bac ._bcf .MinorTickMark =nil ;}else {_bac ._bcf .MinorTickMark =_c .NewCT_TickMark ();_bac ._bcf .MinorTickMark .ValAttr =m ;};};

// SetIndex sets the index of the series
func (_efac SurfaceChartSeries )SetIndex (idx uint32 ){_efac ._gaec .Idx .ValAttr =idx };func (_bccd Legend )SetPosition (p _c .ST_LegendPos ){if p ==_c .ST_LegendPosUnset {_bccd ._eaf .LegendPos =nil ;}else {_bccd ._eaf .LegendPos =_c .NewCT_LegendPos ();_bccd ._eaf .LegendPos .ValAttr =p ;};};

// AreaChart is an area chart that has a shaded area underneath a curve.
type AreaChart struct{chartBase ;_cb *_c .CT_AreaChart ;};

// X returns the inner wrapped XML type.
func (_fafg DateAxis )X ()*_c .CT_DateAx {return _fafg ._dea };

// AddPie3DChart adds a new pie chart to a chart.
func (_fdf Chart )AddPie3DChart ()Pie3DChart {_aab (_fdf ._egd .Chart );_bge :=_c .NewCT_PlotAreaChoice ();_fdf ._egd .Chart .PlotArea .Choice =append (_fdf ._egd .Chart .PlotArea .Choice ,_bge );_bge .Pie3DChart =_c .NewCT_Pie3DChart ();_ef :=Pie3DChart {_eaca :_bge .Pie3DChart };_ef .InitializeDefaults ();return _ef ;};

// SetText sets the series text
func (_daec ScatterChartSeries )SetText (s string ){_daec ._bebd .Tx =_c .NewCT_SerTx ();_daec ._bebd .Tx .Choice .V =&s ;};type GridLines struct{_cbb *_c .CT_ChartLines };func (_bbff ValueAxis )Properties ()_eb .ShapeProperties {if _bbff ._efce .SpPr ==nil {_bbff ._efce .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_bbff ._efce .SpPr );};

// AddAreaChart adds a new area chart to a chart.
func (_ggeb Chart )AddAreaChart ()AreaChart {_age :=_c .NewCT_PlotAreaChoice ();_ggeb ._egd .Chart .PlotArea .Choice =append (_ggeb ._egd .Chart .PlotArea .Choice ,_age );_age .AreaChart =_c .NewCT_AreaChart ();_dbc :=AreaChart {_cb :_age .AreaChart };_dbc .InitializeDefaults ();return _dbc ;};

// X returns the inner wrapped XML type.
func (_aee Line3DChart )X ()*_c .CT_Line3DChart {return _aee ._cadd };type Legend struct{_eaf *_c .CT_Legend };func (_bae DataLabels )ensureChoice (){if _bae ._dfga .Choice ==nil {_bae ._dfga .Choice =_c .NewCT_DLblsChoice ();};};func (_dgce StockChart )AddAxis (axis Axis ){_bcbc :=_c .NewCT_UnsignedInt ();_bcbc .ValAttr =axis .AxisID ();_dgce ._cfe .AxId =append (_dgce ._cfe .AxId ,_bcbc );};func _aab (_dbbe *_c .CT_Chart ){_dbbe .View3D =_c .NewCT_View3D ();_dbbe .View3D .RotX =_c .NewCT_RotX ();_dbbe .View3D .RotX .ValAttr =_ca .Int8 (15);_dbbe .View3D .RotY =_c .NewCT_RotY ();_dbbe .View3D .RotY .ValAttr =_ca .Uint16 (20);_dbbe .View3D .RAngAx =_c .NewCT_Boolean ();_dbbe .View3D .RAngAx .ValAttr =_ca .Bool (false );_dbbe .Floor =_c .NewCT_Surface ();_dbbe .Floor .Thickness =_c .NewCT_Thickness ();_dbbe .Floor .Thickness .ValAttr .Uint32 =_ca .Uint32 (0);_dbbe .SideWall =_c .NewCT_Surface ();_dbbe .SideWall .Thickness =_c .NewCT_Thickness ();_dbbe .SideWall .Thickness .ValAttr .Uint32 =_ca .Uint32 (0);_dbbe .BackWall =_c .NewCT_Surface ();_dbbe .BackWall .Thickness =_c .NewCT_Thickness ();_dbbe .BackWall .Thickness .ValAttr .Uint32 =_ca .Uint32 (0);};

// X returns the inner wrapped XML type.
func (_eec DoughnutChart )X ()*_c .CT_DoughnutChart {return _eec ._cac };

// X returns the inner wrapped XML type.
func (_efg PieOfPieChart )X ()*_c .CT_OfPieChart {return _efg ._abdb };

// AddSeries adds a default series to an area chart.
func (_df Area3DChart )AddSeries ()AreaChartSeries {_bb :=_df .nextColor (len (_df ._b .Ser ));_da :=_c .NewCT_AreaSer ();_df ._b .Ser =append (_df ._b .Ser ,_da );_da .Idx .ValAttr =uint32 (len (_df ._b .Ser )-1);_da .Order .ValAttr =uint32 (len (_df ._b .Ser )-1);_g :=AreaChartSeries {_da };_g .InitializeDefaults ();_g .Properties ().SetSolidFill (_bb );return _g ;};

// SetText sets the series text.
func (_fdfg RadarChartSeries )SetText (s string ){_fdfg ._aefed .Tx =_c .NewCT_SerTx ();_fdfg ._aefed .Tx .Choice .V =&s ;};func (_gfb DataLabels )SetShowSeriesName (b bool ){_gfb .ensureChoice ();_gfb ._dfga .Choice .ShowSerName =_c .NewCT_Boolean ();_gfb ._dfga .Choice .ShowSerName .ValAttr =_ca .Bool (b );};

// X returns the inner wrapped XML type.
func (_bfa ScatterChartSeries )X ()*_c .CT_ScatterSer {return _bfa ._bebd };func (_cgga DateAxis )SetTickLabelPosition (p _c .ST_TickLblPos ){if p ==_c .ST_TickLblPosUnset {_cgga ._dea .TickLblPos =nil ;}else {_cgga ._dea .TickLblPos =_c .NewCT_TickLblPos ();_cgga ._dea .TickLblPos .ValAttr =p ;};};

// Labels returns the data label properties.
func (_dgfa ScatterChartSeries )Labels ()DataLabels {if _dgfa ._bebd .DLbls ==nil {_dgfa ._bebd .DLbls =_c .NewCT_DLbls ();};return MakeDataLabels (_dgfa ._bebd .DLbls );};

// CategoryAxis returns the category data source.
func (_cfc BubbleChartSeries )CategoryAxis ()CategoryAxisDataSource {if _cfc ._cba .XVal ==nil {_cfc ._cba .XVal =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_cfc ._cba .XVal );};

// X returns the inner wrapped XML type.
func (_fgce SeriesAxis )X ()*_c .CT_SerAx {return _fgce ._cga };type nullAxis byte ;

// SetValues is used to set the source data to a set of values.
func (_bgb CategoryAxisDataSource )SetValues (v []string ){_bgb ._ffg .Choice =_c .NewCT_AxDataSourceChoice ();_bgb ._ffg .Choice .StrLit =_c .NewCT_StrData ();_bgb ._ffg .Choice .StrLit .PtCount =_c .NewCT_UnsignedInt ();_bgb ._ffg .Choice .StrLit .PtCount .ValAttr =uint32 (len (v ));for _dg ,_dgf :=range v {_bgb ._ffg .Choice .StrLit .Pt =append (_bgb ._ffg .Choice .StrLit .Pt ,&_c .CT_StrVal {IdxAttr :uint32 (_dg ),V :_dgf });};};

// Properties returns the chart's shape properties.
func (_efa Chart )Properties ()_eb .ShapeProperties {if _efa ._egd .SpPr ==nil {_efa ._egd .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_efa ._egd .SpPr );};

// X returns the inner wrapped XML type.
func (_dbga StockChart )X ()*_c .CT_StockChart {return _dbga ._cfe };func (_fef DataLabels )SetPosition (p _c .ST_DLblPos ){_fef .ensureChoice ();_fef ._dfga .Choice .DLblPos =_c .NewCT_DLblPos ();_fef ._dfga .Choice .DLblPos .ValAttr =p ;};

// ScatterChartSeries is the data series for a scatter chart.
type ScatterChartSeries struct{_bebd *_c .CT_ScatterSer };

// InitializeDefaults initializes an area series to the default values.
func (_gc AreaChartSeries )InitializeDefaults (){};

// SurfaceChart is a 3D surface chart, viewed from the top-down.
type SurfaceChart struct{chartBase ;_fbcf *_c .CT_SurfaceChart ;};

// AddSeries adds a default series to a Stock chart.
func (_fbc StockChart )AddSeries ()LineChartSeries {_fdbf :=_c .NewCT_LineSer ();_fbc ._cfe .Ser =append (_fbc ._cfe .Ser ,_fdbf );_fdbf .Idx .ValAttr =uint32 (len (_fbc ._cfe .Ser )-1);_fdbf .Order .ValAttr =uint32 (len (_fbc ._cfe .Ser )-1);_dde :=LineChartSeries {_fdbf };_dde .Values ().CreateEmptyNumberCache ();_dde .Properties ().LineProperties ().SetNoFill ();return _dde ;};func (_gg CategoryAxis )Properties ()_eb .ShapeProperties {if _gg ._bcf .SpPr ==nil {_gg ._bcf .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_gg ._bcf .SpPr );};

// Values returns the value data source.
func (_bbf BubbleChartSeries )Values ()NumberDataSource {if _bbf ._cba .YVal ==nil {_bbf ._cba .YVal =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_bbf ._cba .YVal );};

// AddCategoryAxis adds a category axis.
func (_bbge Chart )AddCategoryAxis ()CategoryAxis {_cbgf :=_c .NewCT_CatAx ();if _bbge ._egd .Chart .PlotArea .CChoice ==nil {_bbge ._egd .Chart .PlotArea .CChoice =_c .NewCT_PlotAreaChoice1 ();};_cbgf .AxId =_c .NewCT_UnsignedInt ();_cbgf .AxId .ValAttr =0x7FFFFFFF&_fd .Uint32 ();_bbge ._egd .Chart .PlotArea .CChoice .CatAx =append (_bbge ._egd .Chart .PlotArea .CChoice .CatAx ,_cbgf );_cbgf .Auto =_c .NewCT_Boolean ();_cbgf .Auto .ValAttr =_ca .Bool (true );_cbgf .Delete =_c .NewCT_Boolean ();_cbgf .Delete .ValAttr =_ca .Bool (false );_add :=MakeCategoryAxis (_cbgf );_add .InitializeDefaults ();return _add ;};func (_acbd ScatterChart )InitializeDefaults (){_acbd ._eeee .ScatterStyle .ValAttr =_c .ST_ScatterStyleMarker ;};type DataLabels struct{_dfga *_c .CT_DLbls };

// AddSeries adds a default series to an Pie chart.
func (_edgg PieOfPieChart )AddSeries ()PieChartSeries {_bcad :=_c .NewCT_PieSer ();_edgg ._abdb .Ser =append (_edgg ._abdb .Ser ,_bcad );_bcad .Idx .ValAttr =uint32 (len (_edgg ._abdb .Ser )-1);_bcad .Order .ValAttr =uint32 (len (_edgg ._abdb .Ser )-1);_ffe :=PieChartSeries {_bcad };_ffe .InitializeDefaults ();return _ffe ;};var NullAxis Axis =nullAxis (0);

// X returns the inner wrapped XML type.
func (_dgdb ValueAxis )X ()*_c .CT_ValAx {return _dgdb ._efce };

// X returns the inner wrapped XML type.
func (_bab Surface3DChart )X ()*_c .CT_Surface3DChart {return _bab ._daeb };

// AddLineChart adds a new line chart to a chart.
func (_ebe Chart )AddLineChart ()LineChart {_caf :=_c .NewCT_PlotAreaChoice ();_ebe ._egd .Chart .PlotArea .Choice =append (_ebe ._egd .Chart .PlotArea .Choice ,_caf );_caf .LineChart =_c .NewCT_LineChart ();_caf .LineChart .Grouping =_c .NewCT_Grouping ();_caf .LineChart .Grouping .ValAttr =_c .ST_GroupingStandard ;return LineChart {_bed :_caf .LineChart };};

// SetOrder sets the order of the series
func (_bea ScatterChartSeries )SetOrder (idx uint32 ){_bea ._bebd .Order .ValAttr =idx };

// PieChartSeries is a series to be used on an Pie chart.
type PieChartSeries struct{_faag *_c .CT_PieSer };

// X returns the inner wrapped XML type.
func (_fggb RadarChartSeries )X ()*_c .CT_RadarSer {return _fggb ._aefed };func (_bba Legend )SetOverlay (b bool ){_bba ._eaf .Overlay =_c .NewCT_Boolean ();_bba ._eaf .Overlay .ValAttr =_ca .Bool (b );};func (_bfe Marker )SetSymbol (s _c .ST_MarkerStyle ){if s ==_c .ST_MarkerStyleUnset {_bfe ._afg .Symbol =nil ;}else {_bfe ._afg .Symbol =_c .NewCT_MarkerStyle ();_bfe ._afg .Symbol .ValAttr =s ;};};

// AddSurfaceChart adds a new surface chart to a chart.
func (_dc Chart )AddSurfaceChart ()SurfaceChart {_gb :=_c .NewCT_PlotAreaChoice ();_dc ._egd .Chart .PlotArea .Choice =append (_dc ._egd .Chart .PlotArea .Choice ,_gb );_gb .SurfaceChart =_c .NewCT_SurfaceChart ();_aab (_dc ._egd .Chart );_dc ._egd .Chart .View3D .RotX .ValAttr =_ca .Int8 (90);_dc ._egd .Chart .View3D .RotY .ValAttr =_ca .Uint16 (0);_dc ._egd .Chart .View3D .Perspective =_c .NewCT_Perspective ();_dc ._egd .Chart .View3D .Perspective .ValAttr =_ca .Uint8 (0);_eea :=SurfaceChart {_fbcf :_gb .SurfaceChart };_eea .InitializeDefaults ();return _eea ;};

// AddSeries adds a default series to a Scatter chart.
func (_bgbe ScatterChart )AddSeries ()ScatterChartSeries {_gcf :=_bgbe .nextColor (len (_bgbe ._eeee .Ser ));_gbd :=_c .NewCT_ScatterSer ();_bgbe ._eeee .Ser =append (_bgbe ._eeee .Ser ,_gbd );_gbd .Idx .ValAttr =uint32 (len (_bgbe ._eeee .Ser )-1);_gbd .Order .ValAttr =uint32 (len (_bgbe ._eeee .Ser )-1);_beba :=ScatterChartSeries {_gbd };_beba .InitializeDefaults ();_beba .Marker ().Properties ().LineProperties ().SetSolidFill (_gcf );_beba .Marker ().Properties ().SetSolidFill (_gcf );return _beba ;};

// Properties returns the bar chart series shape properties.
func (_cgef PieChartSeries )Properties ()_eb .ShapeProperties {if _cgef ._faag .SpPr ==nil {_cgef ._faag .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_cgef ._faag .SpPr );};func MakeSeriesAxis (x *_c .CT_SerAx )SeriesAxis {return SeriesAxis {x }};

// Pie3DChart is a Pie3D chart.
type Pie3DChart struct{chartBase ;_eaca *_c .CT_Pie3DChart ;};

// CreateEmptyNumberCache creates an empty number cache, which is used sometimes
// to increase file format compatibility.  It should actually contain the
// computed cell data, but just creating an empty one is good enough.
func (_aae NumberDataSource )CreateEmptyNumberCache (){_aae .ensureChoice ();if _aae ._aeg .Choice .NumRef ==nil {_aae ._aeg .Choice .NumRef =_c .NewCT_NumRef ();};_aae ._aeg .Choice .NumLit =nil ;_aae ._aeg .Choice .NumRef .NumCache =_c .NewCT_NumData ();_aae ._aeg .Choice .NumRef .NumCache .PtCount =_c .NewCT_UnsignedInt ();_aae ._aeg .Choice .NumRef .NumCache .PtCount .ValAttr =0;};type ValueAxis struct{_efce *_c .CT_ValAx };

// InitializeDefaults the bar chart to its defaults
func (_gd BarChart )InitializeDefaults (){_gd .SetDirection (_c .ST_BarDirCol )};func (_aca DataLabels )SetShowPercent (b bool ){_aca .ensureChoice ();_aca ._dfga .Choice .ShowPercent =_c .NewCT_Boolean ();_aca ._dfga .Choice .ShowPercent .ValAttr =_ca .Bool (b );};

// AddAxis adds an axis to a line chart.
func (_fgg Line3DChart )AddAxis (axis Axis ){_fbe :=_c .NewCT_UnsignedInt ();_fbe .ValAttr =axis .AxisID ();_fgg ._cadd .AxId =append (_fgg ._cadd .AxId ,_fbe );};func (_ac Bar3DChart )AddAxis (axis Axis ){_bc :=_c .NewCT_UnsignedInt ();_bc .ValAttr =axis .AxisID ();_ac ._a .AxId =append (_ac ._a .AxId ,_bc );};func (_bcfc SurfaceChartSeries )InitializeDefaults (){_bcfc .Properties ().LineProperties ().SetWidth (1*_cf .Point );_bcfc .Properties ().LineProperties ().SetSolidFill (_d .Black );_bcfc .Properties ().LineProperties ().SetJoin (_eb .LineJoinRound );};

// X returns the inner wrapped XML type.
func (_dad LineChartSeries )X ()*_c .CT_LineSer {return _dad ._adb };func (_afb ScatterChartSeries )Values ()NumberDataSource {if _afb ._bebd .YVal ==nil {_afb ._bebd .YVal =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_afb ._bebd .YVal );};func (_fdg LineChartSeries )SetSmooth (b bool ){_fdg ._adb .Smooth =_c .NewCT_Boolean ();_fdg ._adb .Smooth .ValAttr =&b ;};

// AddSeries adds a default series to an Pie3D chart.
func (_faad Pie3DChart )AddSeries ()PieChartSeries {_aea :=_c .NewCT_PieSer ();_faad ._eaca .Ser =append (_faad ._eaca .Ser ,_aea );_aea .Idx .ValAttr =uint32 (len (_faad ._eaca .Ser )-1);_aea .Order .ValAttr =uint32 (len (_faad ._eaca .Ser )-1);_bbae :=PieChartSeries {_aea };_bbae .InitializeDefaults ();return _bbae ;};

// InitializeDefaults the bar chart to its defaults
func (_fgfc PieChart )InitializeDefaults (){_fgfc ._ggdc .VaryColors =_c .NewCT_Boolean ();_fgfc ._ggdc .VaryColors .ValAttr =_ca .Bool (true );};func (_fdgf ValueAxis )SetTickLabelPosition (p _c .ST_TickLblPos ){if p ==_c .ST_TickLblPosUnset {_fdgf ._efce .TickLblPos =nil ;}else {_fdgf ._efce .TickLblPos =_c .NewCT_TickLblPos ();_fdgf ._efce .TickLblPos .ValAttr =p ;};};

// X returns the inner wrapped XML type.
func (_deb SurfaceChart )X ()*_c .CT_SurfaceChart {return _deb ._fbcf };

// Axis is the interface implemented by different axes when assigning to a
// chart.
type Axis interface{AxisID ()uint32 ;};func (_dbcc SurfaceChart )InitializeDefaults (){_dbcc ._fbcf .Wireframe =_c .NewCT_Boolean ();_dbcc ._fbcf .Wireframe .ValAttr =_ca .Bool (false );_dbcc ._fbcf .BandFmts =_c .NewCT_BandFmts ();for _dbdb :=0;_dbdb < 15;_dbdb ++{_feeb :=_c .NewCT_BandFmt ();_feeb .Idx .ValAttr =uint32 (_dbdb );_feeb .SpPr =_e .NewCT_ShapeProperties ();_acfe :=_eb .MakeShapeProperties (_feeb .SpPr );_acfe .SetSolidFill (_dbcc .nextColor (_dbdb ));_dbcc ._fbcf .BandFmts .BandFmt =append (_dbcc ._fbcf .BandFmts .BandFmt ,_feeb );};};

// InitializeDefaults initializes a Bubble chart series to the default values.
func (_cca BubbleChartSeries )InitializeDefaults (){};

// SetText sets the series text.
func (_gad BarChartSeries )SetText (s string ){_gad ._cff .Tx =_c .NewCT_SerTx ();_gad ._cff .Tx .Choice .V =&s ;};

// AddStockChart adds a new stock chart.
func (_ccd Chart )AddStockChart ()StockChart {_acb :=_c .NewCT_PlotAreaChoice ();_ccd ._egd .Chart .PlotArea .Choice =append (_ccd ._egd .Chart .PlotArea .Choice ,_acb );_acb .StockChart =_c .NewCT_StockChart ();_gge :=StockChart {_cfe :_acb .StockChart };_gge .InitializeDefaults ();return _gge ;};

// X returns the inner wrapped XML type.
func (_ecf Title )X ()*_c .CT_Title {return _ecf ._feb };

// AddSeries adds a default series to a line chart.
func (_fgf LineChart )AddSeries ()LineChartSeries {_gbg :=_fgf .nextColor (len (_fgf ._bed .Ser ));_ggda :=_c .NewCT_LineSer ();_fgf ._bed .Ser =append (_fgf ._bed .Ser ,_ggda );_ggda .Idx .ValAttr =uint32 (len (_fgf ._bed .Ser )-1);_ggda .Order .ValAttr =uint32 (len (_fgf ._bed .Ser )-1);_acbf :=LineChartSeries {_ggda };_acbf .InitializeDefaults ();_acbf .Properties ().LineProperties ().SetSolidFill (_gbg );return _acbf ;};

// Index returns the index of the series
func (_fgd ScatterChartSeries )Index ()uint32 {return _fgd ._bebd .Idx .ValAttr };

// Bar3DChart is a 3D bar chart.
type Bar3DChart struct{chartBase ;_a *_c .CT_Bar3DChart ;};

// Properties returns the bar chart series shape properties.
func (_ff AreaChartSeries )Properties ()_eb .ShapeProperties {if _ff ._be .SpPr ==nil {_ff ._be .SpPr =_e .NewCT_ShapeProperties ();};return _eb .MakeShapeProperties (_ff ._be .SpPr );};

// SetIndex sets the index of the series
func (_ceffg ScatterChartSeries )SetIndex (idx uint32 ){_ceffg ._bebd .Idx .ValAttr =idx };func (_bec DataLabels )SetShowCategoryName (b bool ){_bec .ensureChoice ();_bec ._dfga .Choice .ShowCatName =_c .NewCT_Boolean ();_bec ._dfga .Choice .ShowCatName .ValAttr =_ca .Bool (b );};func (_edggg Title )InitializeDefaults (){_edggg .SetText ("\u0054\u0069\u0074l\u0065");_edggg .RunProperties ().SetSize (16*_cf .Point );_edggg .RunProperties ().SetSolidFill (_d .Black );_edggg .RunProperties ().SetFont ("\u0043\u0061\u006c\u0069\u0062\u0020\u0072\u0069");_edggg .RunProperties ().SetBold (false );};

// CategoryAxis returns the category data source.
func (_bgfe RadarChartSeries )CategoryAxis ()CategoryAxisDataSource {if _bgfe ._aefed .Cat ==nil {_bgfe ._aefed .Cat =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_bgfe ._aefed .Cat );};

// SetType sets the type the secone pie to either pie or bar
func (_gfe PieOfPieChart )SetType (t _c .ST_OfPieType ){_gfe ._abdb .OfPieType .ValAttr =t };func (_fcg Title )RunProperties ()_eb .RunProperties {if _fcg ._feb .Tx ==nil {_fcg .SetText ("");};if _fcg ._feb .Tx .Choice .Rich .P [0].EG_TextRun [0].R .RPr ==nil {_fcg ._feb .Tx .Choice .Rich .P [0].EG_TextRun [0].R .RPr =_e .NewCT_TextCharacterProperties ();};return _eb .MakeRunProperties (_fcg ._feb .Tx .Choice .Rich .P [0].EG_TextRun [0].R .RPr );};func (_dgd DateAxis )SetCrosses (axis Axis ){_dgd ._dea .CrossAx .ValAttr =axis .AxisID ()};

// SetText sets the series text
func (_dgcb LineChartSeries )SetText (s string ){_dgcb ._adb .Tx =_c .NewCT_SerTx ();_dgcb ._adb .Tx .Choice .V =&s ;};

// X returns the inner wrapped XML type.
func (_gadd ScatterChart )X ()*_c .CT_ScatterChart {return _gadd ._eeee };

// SetIndex sets the index of the series
func (_bacg LineChartSeries )SetIndex (idx uint32 ){_bacg ._adb .Idx .ValAttr =idx };

// AddAxis adds an axis to a Surface chart.
func (_fgfg SurfaceChart )AddAxis (axis Axis ){_cffb :=_c .NewCT_UnsignedInt ();_cffb .ValAttr =axis .AxisID ();_fgfg ._fbcf .AxId =append (_fgfg ._fbcf .AxId ,_cffb );};

// Order returns the order of the series
func (_dab LineChartSeries )Order ()uint32 {return _dab ._adb .Order .ValAttr };

// SetText sets the series text
func (_afc SurfaceChartSeries )SetText (s string ){_afc ._gaec .Tx =_c .NewCT_SerTx ();_afc ._gaec .Tx .Choice .V =&s ;};func (_ageg LineChartSeries )InitializeDefaults (){_ageg .Properties ().LineProperties ().SetWidth (1*_cf .Point );_ageg .Properties ().LineProperties ().SetSolidFill (_d .Black );_ageg .Properties ().LineProperties ().SetJoin (_eb .LineJoinRound );_ageg .Marker ().SetSymbol (_c .ST_MarkerStyleNone );_ageg .Labels ().SetShowLegendKey (false );_ageg .Labels ().SetShowValue (false );_ageg .Labels ().SetShowPercent (false );_ageg .Labels ().SetShowCategoryName (false );_ageg .Labels ().SetShowSeriesName (false );_ageg .Labels ().SetShowLeaderLines (false );};

// SetDirection changes the direction of the bar chart (bar or column).
func (_aa Bar3DChart )SetDirection (d _c .ST_BarDir ){_aa ._a .BarDir .ValAttr =d };

// BarChartSeries is a series to be used on a bar chart.
type BarChartSeries struct{_cff *_c .CT_BarSer };

// AddPieChart adds a new pie chart to a chart.
func (_adc Chart )AddPieChart ()PieChart {_ceff :=_c .NewCT_PlotAreaChoice ();_adc ._egd .Chart .PlotArea .Choice =append (_adc ._egd .Chart .PlotArea .Choice ,_ceff );_ceff .PieChart =_c .NewCT_PieChart ();_edd :=PieChart {_ggdc :_ceff .PieChart };_edd .InitializeDefaults ();return _edd ;};func (_gbc Chart )AddSeriesAxis ()SeriesAxis {_eeb :=_c .NewCT_SerAx ();if _gbc ._egd .Chart .PlotArea .CChoice ==nil {_gbc ._egd .Chart .PlotArea .CChoice =_c .NewCT_PlotAreaChoice1 ();};_eeb .AxId =_c .NewCT_UnsignedInt ();_eeb .AxId .ValAttr =0x7FFFFFFF&_fd .Uint32 ();_gbc ._egd .Chart .PlotArea .CChoice .SerAx =append (_gbc ._egd .Chart .PlotArea .CChoice .SerAx ,_eeb );_eeb .Delete =_c .NewCT_Boolean ();_eeb .Delete .ValAttr =_ca .Bool (false );_ggd :=MakeSeriesAxis (_eeb );_ggd .InitializeDefaults ();return _ggd ;};func (_gbga ValueAxis )SetMinorTickMark (m _c .ST_TickMark ){if m ==_c .ST_TickMarkUnset {_gbga ._efce .MinorTickMark =nil ;}else {_gbga ._efce .MinorTickMark =_c .NewCT_TickMark ();_gbga ._efce .MinorTickMark .ValAttr =m ;};};

// Values returns the value data source.
func (_cbac PieChartSeries )Values ()NumberDataSource {if _cbac ._faag .Val ==nil {_cbac ._faag .Val =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_cbac ._faag .Val );};

// X returns the inner wrapped XML type.
func (_cfd SurfaceChartSeries )X ()*_c .CT_SurfaceSer {return _cfd ._gaec };

// X returns the inner wrapped XML type.
func (_af BubbleChartSeries )X ()*_c .CT_BubbleSer {return _af ._cba };

// X returns the inner wrapped XML type.
func (_geb RadarChart )X ()*_c .CT_RadarChart {return _geb ._abe };

// CategoryAxis returns the category data source.
func (_gac BarChartSeries )CategoryAxis ()CategoryAxisDataSource {if _gac ._cff .Cat ==nil {_gac ._cff .Cat =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_gac ._cff .Cat );};func (_eac LineChartSeries )Values ()NumberDataSource {if _eac ._adb .Val ==nil {_eac ._adb .Val =_c .NewCT_NumDataSource ();};return MakeNumberDataSource (_eac ._adb .Val );};func (_fdb DataLabels )SetShowValue (b bool ){_fdb .ensureChoice ();_fdb ._dfga .Choice .ShowVal =_c .NewCT_Boolean ();_fdb ._dfga .Choice .ShowVal .ValAttr =_ca .Bool (b );};

// X returns the inner wrapped XML type.
func (_ebg BarChartSeries )X ()*_c .CT_BarSer {return _ebg ._cff };func (_dfe CategoryAxis )MajorGridLines ()GridLines {if _dfe ._bcf .MajorGridlines ==nil {_dfe ._bcf .MajorGridlines =_c .NewCT_ChartLines ();};return GridLines {_dfe ._bcf .MajorGridlines };};

// SetNumberReference is used to set the source data to a range of cells containing
// numbers.
func (_cbae CategoryAxisDataSource )SetNumberReference (s string ){_cbae ._ffg .Choice =_c .NewCT_AxDataSourceChoice ();_cbae ._ffg .Choice .NumRef =_c .NewCT_NumRef ();_cbae ._ffg .Choice .NumRef .F =s ;};

// AddSeries adds a default series to an Radar chart.
func (_efabb RadarChart )AddSeries ()RadarChartSeries {_acge :=_efabb .nextColor (len (_efabb ._abe .Ser ));_fggf :=_c .NewCT_RadarSer ();_efabb ._abe .Ser =append (_efabb ._abe .Ser ,_fggf );_fggf .Idx .ValAttr =uint32 (len (_efabb ._abe .Ser )-1);_fggf .Order .ValAttr =uint32 (len (_efabb ._abe .Ser )-1);_adef :=RadarChartSeries {_fggf };_adef .InitializeDefaults ();_adef .Properties ().SetSolidFill (_acge );return _adef ;};func MakeNumberDataSource (x *_c .CT_NumDataSource )NumberDataSource {return NumberDataSource {x }};

// AddSeries adds a default series to a line chart.
func (_cd Line3DChart )AddSeries ()LineChartSeries {_edgc :=_cd .nextColor (len (_cd ._cadd .Ser ));_ggca :=_c .NewCT_LineSer ();_cd ._cadd .Ser =append (_cd ._cadd .Ser ,_ggca );_ggca .Idx .ValAttr =uint32 (len (_cd ._cadd .Ser )-1);_ggca .Order .ValAttr =uint32 (len (_cd ._cadd .Ser )-1);_fga :=LineChartSeries {_ggca };_fga .InitializeDefaults ();_fga .Properties ().LineProperties ().SetSolidFill (_edgc );_fga .Properties ().SetSolidFill (_edgc );return _fga ;};type CategoryAxis struct{_bcf *_c .CT_CatAx };func (_acfc Surface3DChart )InitializeDefaults (){_acfc ._daeb .Wireframe =_c .NewCT_Boolean ();_acfc ._daeb .Wireframe .ValAttr =_ca .Bool (false );_acfc ._daeb .BandFmts =_c .NewCT_BandFmts ();for _fbd :=0;_fbd < 15;_fbd ++{_fbda :=_c .NewCT_BandFmt ();_fbda .Idx .ValAttr =uint32 (_fbd );_fbda .SpPr =_e .NewCT_ShapeProperties ();_dbdg :=_eb .MakeShapeProperties (_fbda .SpPr );_dbdg .SetSolidFill (_acfc .nextColor (_fbd ));_acfc ._daeb .BandFmts .BandFmt =append (_acfc ._daeb .BandFmts .BandFmt ,_fbda );};};func (_ccf DateAxis )MajorGridLines ()GridLines {if _ccf ._dea .MajorGridlines ==nil {_ccf ._dea .MajorGridlines =_c .NewCT_ChartLines ();};return GridLines {_ccf ._dea .MajorGridlines };};

// AddSeries adds a default series to a Surface chart.
func (_gabf SurfaceChart )AddSeries ()SurfaceChartSeries {_daba :=_gabf .nextColor (len (_gabf ._fbcf .Ser ));_dbcd :=_c .NewCT_SurfaceSer ();_gabf ._fbcf .Ser =append (_gabf ._fbcf .Ser ,_dbcd );_dbcd .Idx .ValAttr =uint32 (len (_gabf ._fbcf .Ser )-1);_dbcd .Order .ValAttr =uint32 (len (_gabf ._fbcf .Ser )-1);_adf :=SurfaceChartSeries {_dbcd };_adf .InitializeDefaults ();_adf .Properties ().LineProperties ().SetSolidFill (_daba );return _adf ;};

// SetDirection changes the direction of the bar chart (bar or column).
func (_gf BarChart )SetDirection (d _c .ST_BarDir ){_gf ._gcd .BarDir .ValAttr =d };

// SetText sets the series text.
func (_daa AreaChartSeries )SetText (s string ){_daa ._be .Tx =_c .NewCT_SerTx ();_daa ._be .Tx .Choice .V =&s ;};func (_gae CategoryAxis )SetPosition (p _c .ST_AxPos ){_gae ._bcf .AxPos =_c .NewCT_AxPos ();_gae ._bcf .AxPos .ValAttr =p ;};func (_ecd ValueAxis )SetCrosses (axis Axis ){_ecd ._efce .CrossAx .ValAttr =axis .AxisID ()};

// X returns the inner wrapped XML type.
func (_bbcg LineChart )X ()*_c .CT_LineChart {return _bbcg ._bed };func (_ebf LineChartSeries )CategoryAxis ()CategoryAxisDataSource {if _ebf ._adb .Cat ==nil {_ebf ._adb .Cat =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_ebf ._adb .Cat );};

// BubbleChart is a 2D Bubble chart.
type BubbleChart struct{chartBase ;_eed *_c .CT_BubbleChart ;};

// SetValues sets values directly on a source.
func (_fcd NumberDataSource )SetValues (v []float64 ){_fcd .ensureChoice ();_fcd ._aeg .Choice .NumRef =nil ;_fcd ._aeg .Choice .NumLit =_c .NewCT_NumData ();_fcd ._aeg .Choice .NumLit .PtCount =_c .NewCT_UnsignedInt ();_fcd ._aeg .Choice .NumLit .PtCount .ValAttr =uint32 (len (v ));for _agb ,_gdb :=range v {_fcd ._aeg .Choice .NumLit .Pt =append (_fcd ._aeg .Choice .NumLit .Pt ,&_c .CT_NumVal {IdxAttr :uint32 (_agb ),V :_ce .Sprintf ("\u0025\u0067",_gdb )});};};func (_ffea Title )SetText (s string ){if _ffea ._feb .Tx ==nil {_ffea ._feb .Tx =_c .NewCT_Tx ();};if _ffea ._feb .Tx .Choice .Rich ==nil {_ffea ._feb .Tx .Choice .Rich =_e .NewCT_TextBody ();};var _ddgd *_e .CT_TextParagraph ;if len (_ffea ._feb .Tx .Choice .Rich .P )==0{_ddgd =_e .NewCT_TextParagraph ();_ffea ._feb .Tx .Choice .Rich .P =[]*_e .CT_TextParagraph {_ddgd };}else {_ddgd =_ffea ._feb .Tx .Choice .Rich .P [0];};var _bbaeg *_e .EG_TextRun ;if len (_ddgd .EG_TextRun )==0{_bbaeg =_e .NewEG_TextRun ();_ddgd .EG_TextRun =[]*_e .EG_TextRun {_bbaeg };}else {_bbaeg =_ddgd .EG_TextRun [0];};if _bbaeg .R ==nil {_bbaeg .R =_e .NewCT_RegularTextRun ();};_bbaeg .R .T =s ;};

// MakeAxisDataSource constructs an AxisDataSource wrapper.
func MakeAxisDataSource (x *_c .CT_AxDataSource )CategoryAxisDataSource {return CategoryAxisDataSource {x };};func (_gcdb SurfaceChartSeries )CategoryAxis ()CategoryAxisDataSource {if _gcdb ._gaec .Cat ==nil {_gcdb ._gaec .Cat =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_gcdb ._gaec .Cat );};

// X returns the inner wrapped XML type.
func (_gfgc PieChart )X ()*_c .CT_PieChart {return _gfgc ._ggdc };type Line3DChart struct{chartBase ;_cadd *_c .CT_Line3DChart ;};type LineChart struct{chartBase ;_bed *_c .CT_LineChart ;};func MakeTitle (x *_c .CT_Title )Title {return Title {x }};

// PieOfPieChart is a Pie chart with an extra Pie chart.
type PieOfPieChart struct{chartBase ;_abdb *_c .CT_OfPieChart ;};

// RadarChart is an Radar chart that has a shaded Radar underneath a curve.
type RadarChart struct{chartBase ;_abe *_c .CT_RadarChart ;};func (_efc ScatterChartSeries )InitializeDefaults (){_efc .Properties ().LineProperties ().SetNoFill ();_efc .Marker ().SetSymbol (_c .ST_MarkerStyleAuto );_efc .Labels ().SetShowLegendKey (false );_efc .Labels ().SetShowValue (true );_efc .Labels ().SetShowPercent (false );_efc .Labels ().SetShowCategoryName (false );_efc .Labels ().SetShowSeriesName (false );_efc .Labels ().SetShowLeaderLines (false );};func (_edda ScatterChartSeries )CategoryAxis ()CategoryAxisDataSource {if _edda ._bebd .XVal ==nil {_edda ._bebd .XVal =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_edda ._bebd .XVal );};func (_ccff ValueAxis )SetPosition (p _c .ST_AxPos ){_ccff ._efce .AxPos =_c .NewCT_AxPos ();_ccff ._efce .AxPos .ValAttr =p ;};

// X returns the inner wrapped XML type.
func (_adab Marker )X ()*_c .CT_Marker {return _adab ._afg };

// BubbleChartSeries is a series to be used on a Bubble chart.
type BubbleChartSeries struct{_cba *_c .CT_BubbleSer };

// AddLegend adds a legend to a chart, replacing any existing legend.
func (_ab Chart )AddLegend ()Legend {_ab ._egd .Chart .Legend =_c .NewCT_Legend ();_fc :=MakeLegend (_ab ._egd .Chart .Legend );_fc .InitializeDefaults ();return _fc ;};

// InitializeDefaults the bar chart to its defaults
func (_ed Area3DChart )InitializeDefaults (){};

// AddSeries adds a default series to a Surface chart.
func (_cgc Surface3DChart )AddSeries ()SurfaceChartSeries {_eca :=_cgc .nextColor (len (_cgc ._daeb .Ser ));_ggf :=_c .NewCT_SurfaceSer ();_cgc ._daeb .Ser =append (_cgc ._daeb .Ser ,_ggf );_ggf .Idx .ValAttr =uint32 (len (_cgc ._daeb .Ser )-1);_ggf .Order .ValAttr =uint32 (len (_cgc ._daeb .Ser )-1);_cbd :=SurfaceChartSeries {_ggf };_cbd .InitializeDefaults ();_cbd .Properties ().LineProperties ().SetSolidFill (_eca );return _cbd ;};func (_bbgc DateAxis )SetMajorTickMark (m _c .ST_TickMark ){if m ==_c .ST_TickMarkUnset {_bbgc ._dea .MajorTickMark =nil ;}else {_bbgc ._dea .MajorTickMark =_c .NewCT_TickMark ();_bbgc ._dea .MajorTickMark .ValAttr =m ;};};

// Surface3DChart is a 3D view of a surface chart.
type Surface3DChart struct{chartBase ;_daeb *_c .CT_Surface3DChart ;};

// X returns the inner wrapped XML type.
func (_gdf GridLines )X ()*_c .CT_ChartLines {return _gdf ._cbb };

// Index returns the index of the series
func (_bgge SurfaceChartSeries )Index ()uint32 {return _bgge ._gaec .Idx .ValAttr };

// CategoryAxis returns the category data source.
func (_cag PieChartSeries )CategoryAxis ()CategoryAxisDataSource {if _cag ._faag .Cat ==nil {_cag ._faag .Cat =_c .NewCT_AxDataSource ();};return MakeAxisDataSource (_cag ._faag .Cat );};func (_ade NumberDataSource )SetReference (s string ){_ade .ensureChoice ();if _ade ._aeg .Choice .NumRef ==nil {_ade ._aeg .Choice .NumRef =_c .NewCT_NumRef ();};_ade ._aeg .Choice .NumRef .F =s ;};

// PieChart is a Pie chart.
type PieChart struct{chartBase ;_ggdc *_c .CT_PieChart ;};func (_ebdf ValueAxis )SetMajorTickMark (m _c .ST_TickMark ){if m ==_c .ST_TickMarkUnset {_ebdf ._efce .MajorTickMark =nil ;}else {_ebdf ._efce .MajorTickMark =_c .NewCT_TickMark ();_ebdf ._efce .MajorTickMark .ValAttr =m ;};};func (_ffeb SeriesAxis )SetCrosses (axis Axis ){_ffeb ._cga .CrossAx .ValAttr =axis .AxisID ()};type Marker struct{_afg *_c .CT_Marker };
