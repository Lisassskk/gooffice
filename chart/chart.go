package chart

import (
	_d "fmt"
	_gd "math/rand"

	_f "gitee.com/yu_sheng/gooffice"
	_ef "gitee.com/yu_sheng/gooffice/color"
	_ge "gitee.com/yu_sheng/gooffice/drawing"
	_e "gitee.com/yu_sheng/gooffice/measurement"
	_ga "gitee.com/yu_sheng/gooffice/schema/soo/dml"
	_g "gitee.com/yu_sheng/gooffice/schema/soo/dml/chart"
)

// AreaChart is an area chart that has a shaded area underneath a curve.
type AreaChart struct {
	chartBase
	_cb *_g.CT_AreaChart
}

// InitializeDefaults the bar chart to its defaults
func (_ebea Pie3DChart) InitializeDefaults() {
	_ebea._ebfb.VaryColors = _g.NewCT_Boolean()
	_ebea._ebfb.VaryColors.ValAttr = _f.Bool(true)
}

type LineChart struct {
	chartBase
	_cbc *_g.CT_LineChart
}

// InitializeDefaults initializes a Bubble chart series to the default values.
func (_cab BubbleChartSeries) InitializeDefaults() {}

// X returns the inner wrapped XML type.
func (_dbc DoughnutChart) X() *_g.CT_DoughnutChart { return _dbc._bae }

// SetLabelReference is used to set the source data to a range of cells
// containing strings.
func (_fda CategoryAxisDataSource) SetLabelReference(s string) {
	_fda._ed.Choice = _g.NewCT_AxDataSourceChoice()
	_fda._ed.Choice.StrRef = _g.NewCT_StrRef()
	_fda._ed.Choice.StrRef.F = s
}

// AddSeries adds a default series to a bar chart.
func (_dfe BarChart) AddSeries() BarChartSeries {
	_bba := _dfe.nextColor(len(_dfe._bd.Ser))
	_ea := _g.NewCT_BarSer()
	_dfe._bd.Ser = append(_dfe._bd.Ser, _ea)
	_ea.Idx.ValAttr = uint32(len(_dfe._bd.Ser) - 1)
	_ea.Order.ValAttr = uint32(len(_dfe._bd.Ser) - 1)
	_fa := BarChartSeries{_ea}
	_fa.InitializeDefaults()
	_fa.Properties().SetSolidFill(_bba)
	return _fa
}
func (_cge LineChartSeries) SetSmooth(b bool) {
	_cge._feg.Smooth = _g.NewCT_Boolean()
	_cge._feg.Smooth.ValAttr = &b
}

// InitializeDefaults the bar chart to its defaults
func (_gefd PieOfPieChart) InitializeDefaults() {
	_gefd._cbff.VaryColors = _g.NewCT_Boolean()
	_gefd._cbff.VaryColors.ValAttr = _f.Bool(true)
	_gefd.SetType(_g.ST_OfPieTypePie)
	_gefd._cbff.SecondPieSize = _g.NewCT_SecondPieSize()
	_gefd._cbff.SecondPieSize.ValAttr = &_g.ST_SecondPieSize{}
	_gefd._cbff.SecondPieSize.ValAttr.ST_SecondPieSizeUShort = _f.Uint16(75)
	_ecg := _g.NewCT_ChartLines()
	_ecg.SpPr = _ga.NewCT_ShapeProperties()
	_gcde := _ge.MakeShapeProperties(_ecg.SpPr)
	_gcde.LineProperties().SetSolidFill(_ef.Auto)
	_gefd._cbff.SerLines = append(_gefd._cbff.SerLines, _ecg)
}
func (_fgfe Chart) AddSeriesAxis() SeriesAxis {
	_aed := _g.NewCT_SerAx()
	if _fgfe._aga.Chart.PlotArea.CChoice == nil {
		_fgfe._aga.Chart.PlotArea.CChoice = _g.NewCT_PlotAreaChoice1()
	}
	_aed.AxId = _g.NewCT_UnsignedInt()
	_aed.AxId.ValAttr = 0x7FFFFFFF & _gd.Uint32()
	_fgfe._aga.Chart.PlotArea.CChoice.SerAx = append(_fgfe._aga.Chart.PlotArea.CChoice.SerAx, _aed)
	_aed.Delete = _g.NewCT_Boolean()
	_aed.Delete.ValAttr = _f.Bool(false)
	_dcf := MakeSeriesAxis(_aed)
	_dcf.InitializeDefaults()
	return _dcf
}

// Values returns the value data source.
func (_bbg BarChartSeries) Values() NumberDataSource {
	if _bbg._fc.Val == nil {
		_bbg._fc.Val = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_bbg._fc.Val)
}

// X returns the inner wrapped XML type.
func (_eecc StockChart) X() *_g.CT_StockChart { return _eecc._bgfg }

// AddLineChart adds a new line chart to a chart.
func (_gab Chart) AddLineChart() LineChart {
	_gb := _g.NewCT_PlotAreaChoice()
	_gab._aga.Chart.PlotArea.Choice = append(_gab._aga.Chart.PlotArea.Choice, _gb)
	_gb.LineChart = _g.NewCT_LineChart()
	_gb.LineChart.Grouping = _g.NewCT_Grouping()
	_gb.LineChart.Grouping.ValAttr = _g.ST_GroupingStandard
	return LineChart{_cbc: _gb.LineChart}
}

// X returns the inner wrapped XML type.
func (_ggc GridLines) X() *_g.CT_ChartLines { return _ggc._gee }

// X returns the inner wrapped XML type.
func (_dab BubbleChartSeries) X() *_g.CT_BubbleSer { return _dab._cf }

// Properties returns the bar chart series shape properties.
func (_ab AreaChartSeries) Properties() _ge.ShapeProperties {
	if _ab._ddc.SpPr == nil {
		_ab._ddc.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_ab._ddc.SpPr)
}
func (_ace SeriesAxis) InitializeDefaults() {}

// Bar3DChart is a 3D bar chart.
type Bar3DChart struct {
	chartBase
	_fgf *_g.CT_Bar3DChart
}

// AddAxis adds an axis to a Scatter chart.
func (_cfcg ScatterChart) AddAxis(axis Axis) {
	_bdg := _g.NewCT_UnsignedInt()
	_bdg.ValAttr = axis.AxisID()
	_cfcg._cfb.AxId = append(_cfcg._cfb.AxId, _bdg)
}

// AddAxis adds an axis to a Surface chart.
func (_begd SurfaceChart) AddAxis(axis Axis) {
	_ccd := _g.NewCT_UnsignedInt()
	_ccd.ValAttr = axis.AxisID()
	_begd._fedb.AxId = append(_begd._fedb.AxId, _ccd)
}

// BarChart is a 2D bar chart.
type BarChart struct {
	chartBase
	_bd *_g.CT_BarChart
}

func (_dgb CategoryAxis) InitializeDefaults() {
	_dgb.SetPosition(_g.ST_AxPosB)
	_dgb.SetMajorTickMark(_g.ST_TickMarkOut)
	_dgb.SetMinorTickMark(_g.ST_TickMarkIn)
	_dgb.SetTickLabelPosition(_g.ST_TickLblPosNextTo)
	_dgb.MajorGridLines().Properties().LineProperties().SetSolidFill(_ef.LightGray)
	_dgb.Properties().LineProperties().SetSolidFill(_ef.Black)
}

// X returns the inner wrapped XML type.
func (_bad Line3DChart) X() *_g.CT_Line3DChart { return _bad._begf }

// Labels returns the data label properties.
func (_fdc ScatterChartSeries) Labels() DataLabels {
	if _fdc._bebg.DLbls == nil {
		_fdc._bebg.DLbls = _g.NewCT_DLbls()
	}
	return MakeDataLabels(_fdc._bebg.DLbls)
}

// SetDirection changes the direction of the bar chart (bar or column).
func (_dbaf BarChart) SetDirection(d _g.ST_BarDir) { _dbaf._bd.BarDir.ValAttr = d }

// SetType sets the type the secone pie to either pie or bar
func (_ddd PieOfPieChart) SetType(t _g.ST_OfPieType) { _ddd._cbff.OfPieType.ValAttr = t }

// Values returns the bubble size data source.
func (_fcd BubbleChartSeries) BubbleSizes() NumberDataSource {
	if _fcd._cf.BubbleSize == nil {
		_fcd._cf.BubbleSize = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_fcd._cf.BubbleSize)
}

// Properties returns the bar chart series shape properties.
func (_ddg BarChartSeries) Properties() _ge.ShapeProperties {
	if _ddg._fc.SpPr == nil {
		_ddg._fc.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_ddg._fc.SpPr)
}

// X returns the inner wrapped XML type.
func (_ccf Title) X() *_g.CT_Title { return _ccf._gga }

// SetText sets the series text.
func (_gc BarChartSeries) SetText(s string) { _gc._fc.Tx = _g.NewCT_SerTx(); _gc._fc.Tx.Choice.V = &s }

// X returns the inner wrapped XML type.
func (_ecd PieOfPieChart) X() *_g.CT_OfPieChart { return _ecd._cbff }

// InitializeDefaults initializes an area series to the default values.
func (_bga AreaChartSeries) InitializeDefaults() {}

type ScatterChart struct {
	chartBase
	_cfb *_g.CT_ScatterChart
}

// AddBarChart adds a new bar chart to a chart.
func (_fce Chart) AddBarChart() BarChart {
	_eg := _g.NewCT_PlotAreaChoice()
	_fce._aga.Chart.PlotArea.Choice = append(_fce._aga.Chart.PlotArea.Choice, _eg)
	_eg.BarChart = _g.NewCT_BarChart()
	_eg.BarChart.Grouping = _g.NewCT_BarGrouping()
	_eg.BarChart.Grouping.ValAttr = _g.ST_BarGroupingStandard
	_fcg := BarChart{_bd: _eg.BarChart}
	_fcg.InitializeDefaults()
	return _fcg
}
func MakeCategoryAxis(x *_g.CT_CatAx) CategoryAxis { return CategoryAxis{x} }

// AddBar3DChart adds a new 3D bar chart to a chart.
func (_abb Chart) AddBar3DChart() Bar3DChart {
	_edf(_abb._aga.Chart)
	_gec := _g.NewCT_PlotAreaChoice()
	_abb._aga.Chart.PlotArea.Choice = append(_abb._aga.Chart.PlotArea.Choice, _gec)
	_gec.Bar3DChart = _g.NewCT_Bar3DChart()
	_gec.Bar3DChart.Grouping = _g.NewCT_BarGrouping()
	_gec.Bar3DChart.Grouping.ValAttr = _g.ST_BarGroupingStandard
	_gcdc := Bar3DChart{_fgf: _gec.Bar3DChart}
	_gcdc.InitializeDefaults()
	return _gcdc
}

// BubbleChartSeries is a series to be used on a Bubble chart.
type BubbleChartSeries struct{ _cf *_g.CT_BubbleSer }

// SetIndex sets the index of the series
func (_bgfa LineChartSeries) SetIndex(idx uint32) { _bgfa._feg.Idx.ValAttr = idx }

type Line3DChart struct {
	chartBase
	_begf *_g.CT_Line3DChart
}

// SetText sets the series text.
func (_ff AreaChartSeries) SetText(s string) {
	_ff._ddc.Tx = _g.NewCT_SerTx()
	_ff._ddc.Tx.Choice.V = &s
}
func (_dbb DateAxis) SetCrosses(axis Axis) { _dbb._bfe.CrossAx.ValAttr = axis.AxisID() }
func (_beg DateAxis) AxisID() uint32       { return _beg._bfe.AxId.ValAttr }

type Marker struct{ _eed *_g.CT_Marker }
type chartBase struct{}

func (_bfa Legend) SetOverlay(b bool) {
	_bfa._cag.Overlay = _g.NewCT_Boolean()
	_bfa._cag.Overlay.ValAttr = _f.Bool(b)
}

// AddArea3DChart adds a new area chart to a chart.
func (_fgb Chart) AddArea3DChart() Area3DChart {
	_edf(_fgb._aga.Chart)
	_acc := _g.NewCT_PlotAreaChoice()
	_fgb._aga.Chart.PlotArea.Choice = append(_fgb._aga.Chart.PlotArea.Choice, _acc)
	_acc.Area3DChart = _g.NewCT_Area3DChart()
	_fec := Area3DChart{_b: _acc.Area3DChart}
	_fec.InitializeDefaults()
	return _fec
}
func (_dgbb DateAxis) SetPosition(p _g.ST_AxPos) {
	_dgbb._bfe.AxPos = _g.NewCT_AxPos()
	_dgbb._bfe.AxPos.ValAttr = p
}

// AddSeries adds a default series to a Scatter chart.
func (_fdb ScatterChart) AddSeries() ScatterChartSeries {
	_cce := _fdb.nextColor(len(_fdb._cfb.Ser))
	_bfb := _g.NewCT_ScatterSer()
	_fdb._cfb.Ser = append(_fdb._cfb.Ser, _bfb)
	_bfb.Idx.ValAttr = uint32(len(_fdb._cfb.Ser) - 1)
	_bfb.Order.ValAttr = uint32(len(_fdb._cfb.Ser) - 1)
	_fff := ScatterChartSeries{_bfb}
	_fff.InitializeDefaults()
	_fff.Marker().Properties().LineProperties().SetSolidFill(_cce)
	_fff.Marker().Properties().SetSolidFill(_cce)
	return _fff
}
func (_ega SeriesAxis) SetCrosses(axis Axis) { _ega._dgdc.CrossAx.ValAttr = axis.AxisID() }
func MakeValueAxis(x *_g.CT_ValAx) ValueAxis { return ValueAxis{x} }

// AddSurface3DChart adds a new 3D surface chart to a chart.
func (_beb Chart) AddSurface3DChart() Surface3DChart {
	_gcce := _g.NewCT_PlotAreaChoice()
	_beb._aga.Chart.PlotArea.Choice = append(_beb._aga.Chart.PlotArea.Choice, _gcce)
	_gcce.Surface3DChart = _g.NewCT_Surface3DChart()
	_edf(_beb._aga.Chart)
	_fgg := Surface3DChart{_eacb: _gcce.Surface3DChart}
	_fgg.InitializeDefaults()
	return _fgg
}

// X returns the inner wrapped XML type.
func (_edca ScatterChart) X() *_g.CT_ScatterChart { return _edca._cfb }
func (_decf LineChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _decf._feg.Cat == nil {
		_decf._feg.Cat = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_decf._feg.Cat)
}

// SetValues sets values directly on a source.
func (_gcb NumberDataSource) SetValues(v []float64) {
	_gcb.ensureChoice()
	_gcb._ffd.Choice.NumRef = nil
	_gcb._ffd.Choice.NumLit = _g.NewCT_NumData()
	_gcb._ffd.Choice.NumLit.PtCount = _g.NewCT_UnsignedInt()
	_gcb._ffd.Choice.NumLit.PtCount.ValAttr = uint32(len(v))
	for _eafa, _bbab := range v {
		_gcb._ffd.Choice.NumLit.Pt = append(_gcb._ffd.Choice.NumLit.Pt, &_g.CT_NumVal{IdxAttr: uint32(_eafa), V: _d.Sprintf("\u0025\u0067", _bbab)})
	}
}

// SetText sets the series text.
func (_eee BubbleChartSeries) SetText(s string) {
	_eee._cf.Tx = _g.NewCT_SerTx()
	_eee._cf.Tx.Choice.V = &s
}
func (_efff SurfaceChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _efff._dee.Cat == nil {
		_efff._dee.Cat = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_efff._dee.Cat)
}

// InitializeDefaults the bar chart to its defaults
func (_afc BarChart) InitializeDefaults() { _afc.SetDirection(_g.ST_BarDirCol) }

// X returns the inner wrapped XML type.
func (_bee BarChartSeries) X() *_g.CT_BarSer { return _bee._fc }

// InitializeDefaults the Bubble chart to its defaults
func (_dgd BubbleChart) InitializeDefaults() {}

// AddTitle sets a new title on the chart.
func (_cc Chart) AddTitle() Title {
	_cc._aga.Chart.Title = _g.NewCT_Title()
	_cc._aga.Chart.Title.Overlay = _g.NewCT_Boolean()
	_cc._aga.Chart.Title.Overlay.ValAttr = _f.Bool(false)
	_cc._aga.Chart.AutoTitleDeleted = _g.NewCT_Boolean()
	_cc._aga.Chart.AutoTitleDeleted.ValAttr = _f.Bool(false)
	_dcc := MakeTitle(_cc._aga.Chart.Title)
	_dcc.InitializeDefaults()
	return _dcc
}
func MakeChart(x *_g.ChartSpace) Chart { return Chart{x} }

var _faad = []_ef.Color{_ef.RGB(0x33, 0x66, 0xcc), _ef.RGB(0xDC, 0x39, 0x12), _ef.RGB(0xFF, 0x99, 0x00), _ef.RGB(0x10, 0x96, 0x18), _ef.RGB(0x99, 0x00, 0x99), _ef.RGB(0x3B, 0x3E, 0xAC), _ef.RGB(0x00, 0x99, 0xC6), _ef.RGB(0xDD, 0x44, 0x77), _ef.RGB(0x66, 0xAA, 0x00), _ef.RGB(0xB8, 0x2E, 0x2E), _ef.RGB(0x31, 0x63, 0x95), _ef.RGB(0x99, 0x44, 0x99), _ef.RGB(0x22, 0xAA, 0x99), _ef.RGB(0xAA, 0xAA, 0x11), _ef.RGB(0x66, 0x33, 0xCC), _ef.RGB(0xE6, 0x73, 0x00), _ef.RGB(0x8B, 0x07, 0x07), _ef.RGB(0x32, 0x92, 0x62), _ef.RGB(0x55, 0x74, 0xA6), _ef.RGB(0x3B, 0x3E, 0xAC)}

// DoughnutChart is a Doughnut chart.
type DoughnutChart struct {
	chartBase
	_bae *_g.CT_DoughnutChart
}
type NumberDataSource struct{ _ffd *_g.CT_NumDataSource }

func (_dag ScatterChart) InitializeDefaults() {
	_dag._cfb.ScatterStyle.ValAttr = _g.ST_ScatterStyleMarker
}

// X returns the inner wrapped XML type.
func (_ad BarChart) X() *_g.CT_BarChart { return _ad._bd }

// CategoryAxis returns the category data source.
func (_geb AreaChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _geb._ddc.Cat == nil {
		_geb._ddc.Cat = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_geb._ddc.Cat)
}

// X returns the inner wrapped XML type.
func (_fggd RadarChart) X() *_g.CT_RadarChart { return _fggd._gaa }

// SetNumberReference is used to set the source data to a range of cells containing
// numbers.
func (_dff CategoryAxisDataSource) SetNumberReference(s string) {
	_dff._ed.Choice = _g.NewCT_AxDataSourceChoice()
	_dff._ed.Choice.NumRef = _g.NewCT_NumRef()
	_dff._ed.Choice.NumRef.F = s
}
func (_gea DateAxis) SetTickLabelPosition(p _g.ST_TickLblPos) {
	if p == _g.ST_TickLblPosUnset {
		_gea._bfe.TickLblPos = nil
	} else {
		_gea._bfe.TickLblPos = _g.NewCT_TickLblPos()
		_gea._bfe.TickLblPos.ValAttr = p
	}
}
func (_fdeb DataLabels) SetShowValue(b bool) {
	_fdeb.ensureChoice()
	_fdeb._gebb.Choice.ShowVal = _g.NewCT_Boolean()
	_fdeb._gebb.Choice.ShowVal.ValAttr = _f.Bool(b)
}
func (_gdfe SurfaceChartSeries) InitializeDefaults() {
	_gdfe.Properties().LineProperties().SetWidth(1 * _e.Point)
	_gdfe.Properties().LineProperties().SetSolidFill(_ef.Black)
	_gdfe.Properties().LineProperties().SetJoin(_ge.LineJoinRound)
}

// Index returns the index of the series
func (_beaf ScatterChartSeries) Index() uint32 { return _beaf._bebg.Idx.ValAttr }

// InitializeDefaults the bar chart to its defaults
func (_gge PieChart) InitializeDefaults() {
	_gge._dgc.VaryColors = _g.NewCT_Boolean()
	_gge._dgc.VaryColors.ValAttr = _f.Bool(true)
}

// Order returns the order of the series
func (_bfd ScatterChartSeries) Order() uint32 { return _bfd._bebg.Order.ValAttr }

// PieOfPieChart is a Pie chart with an extra Pie chart.
type PieOfPieChart struct {
	chartBase
	_cbff *_g.CT_OfPieChart
}

// X returns the inner wrapped XML type.
func (_gcgc Marker) X() *_g.CT_Marker { return _gcgc._eed }

// InitializeDefaults initializes an Radar series to the default values.
func (_baed RadarChartSeries) InitializeDefaults() {}

// CategoryAxis returns the category data source.
func (_ccb PieChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _ccb._eff.Cat == nil {
		_ccb._eff.Cat = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_ccb._eff.Cat)
}

// SetExplosion sets the value that the segements of the pie are 'exploded' by
func (_gdc PieChartSeries) SetExplosion(v uint32) {
	_gdc._eff.Explosion = _g.NewCT_UnsignedInt()
	_gdc._eff.Explosion.ValAttr = v
}
func (_cfe CategoryAxis) MajorGridLines() GridLines {
	if _cfe._bed.MajorGridlines == nil {
		_cfe._bed.MajorGridlines = _g.NewCT_ChartLines()
	}
	return GridLines{_cfe._bed.MajorGridlines}
}

// Labels returns the data label properties.
func (_ada LineChartSeries) Labels() DataLabels {
	if _ada._feg.DLbls == nil {
		_ada._feg.DLbls = _g.NewCT_DLbls()
	}
	return MakeDataLabels(_ada._feg.DLbls)
}

// Surface3DChart is a 3D view of a surface chart.
type Surface3DChart struct {
	chartBase
	_eacb *_g.CT_Surface3DChart
}
type DateAxis struct{ _bfe *_g.CT_DateAx }

func (_aeg ScatterChartSeries) SetSmooth(b bool) {
	_aeg._bebg.Smooth = _g.NewCT_Boolean()
	_aeg._bebg.Smooth.ValAttr = &b
}

type nullAxis byte

func (_fef chartBase) nextColor(_cdgd int) _ef.Color { return _faad[_cdgd%len(_faad)] }

type DataLabels struct{ _gebb *_g.CT_DLbls }

func (_gcd CategoryAxis) AxisID() uint32 { return _gcd._bed.AxId.ValAttr }

// PieChartSeries is a series to be used on an Pie chart.
type PieChartSeries struct{ _eff *_g.CT_PieSer }

// AddPieChart adds a new pie chart to a chart.
func (_aae Chart) AddPieChart() PieChart {
	_efd := _g.NewCT_PlotAreaChoice()
	_aae._aga.Chart.PlotArea.Choice = append(_aae._aga.Chart.PlotArea.Choice, _efd)
	_efd.PieChart = _g.NewCT_PieChart()
	_agba := PieChart{_dgc: _efd.PieChart}
	_agba.InitializeDefaults()
	return _agba
}

// Index returns the index of the series
func (_gdf LineChartSeries) Index() uint32 { return _gdf._feg.Idx.ValAttr }

// BarChartSeries is a series to be used on a bar chart.
type BarChartSeries struct{ _fc *_g.CT_BarSer }

// InitializeDefaults the bar chart to its defaults
func (_ded RadarChart) InitializeDefaults() { _ded._gaa.RadarStyle.ValAttr = _g.ST_RadarStyleMarker }
func (_fbb Legend) SetPosition(p _g.ST_LegendPos) {
	if p == _g.ST_LegendPosUnset {
		_fbb._cag.LegendPos = nil
	} else {
		_fbb._cag.LegendPos = _g.NewCT_LegendPos()
		_fbb._cag.LegendPos.ValAttr = p
	}
}
func (_dbf Legend) Properties() _ge.ShapeProperties {
	if _dbf._cag.SpPr == nil {
		_dbf._cag.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_dbf._cag.SpPr)
}
func (_fbec ValueAxis) SetPosition(p _g.ST_AxPos) {
	_fbec._gbg.AxPos = _g.NewCT_AxPos()
	_fbec._gbg.AxPos.ValAttr = p
}
func MakeSeriesAxis(x *_g.CT_SerAx) SeriesAxis { return SeriesAxis{x} }
func (_ec CategoryAxis) SetMinorTickMark(m _g.ST_TickMark) {
	if m == _g.ST_TickMarkUnset {
		_ec._bed.MinorTickMark = nil
	} else {
		_ec._bed.MinorTickMark = _g.NewCT_TickMark()
		_ec._bed.MinorTickMark.ValAttr = m
	}
}

// X returns the inner wrapped XML type.
func (_gfd RadarChartSeries) X() *_g.CT_RadarSer { return _gfd._fbg }

// InitializeDefaults the bar chart to its defaults
func (_bgf Bar3DChart) InitializeDefaults() { _bgf.SetDirection(_g.ST_BarDirCol) }

// AddSurfaceChart adds a new surface chart to a chart.
func (_age Chart) AddSurfaceChart() SurfaceChart {
	_bea := _g.NewCT_PlotAreaChoice()
	_age._aga.Chart.PlotArea.Choice = append(_age._aga.Chart.PlotArea.Choice, _bea)
	_bea.SurfaceChart = _g.NewCT_SurfaceChart()
	_edf(_age._aga.Chart)
	_age._aga.Chart.View3D.RotX.ValAttr = _f.Int8(90)
	_age._aga.Chart.View3D.RotY.ValAttr = _f.Uint16(0)
	_age._aga.Chart.View3D.Perspective = _g.NewCT_Perspective()
	_age._aga.Chart.View3D.Perspective.ValAttr = _f.Uint8(0)
	_dfc := SurfaceChart{_fedb: _bea.SurfaceChart}
	_dfc.InitializeDefaults()
	return _dfc
}

// SetIndex sets the index of the series
func (_ggcg SurfaceChartSeries) SetIndex(idx uint32) { _ggcg._dee.Idx.ValAttr = idx }

// AddDateAxis adds a value axis to the chart.
func (_dbg Chart) AddDateAxis() DateAxis {
	_bdfc := _g.NewCT_DateAx()
	if _dbg._aga.Chart.PlotArea.CChoice == nil {
		_dbg._aga.Chart.PlotArea.CChoice = _g.NewCT_PlotAreaChoice1()
	}
	_bdfc.AxId = _g.NewCT_UnsignedInt()
	_bdfc.AxId.ValAttr = 0x7FFFFFFF & _gd.Uint32()
	_dbg._aga.Chart.PlotArea.CChoice.DateAx = append(_dbg._aga.Chart.PlotArea.CChoice.DateAx, _bdfc)
	_bdfc.Delete = _g.NewCT_Boolean()
	_bdfc.Delete.ValAttr = _f.Bool(false)
	_bdfc.Scaling = _g.NewCT_Scaling()
	_bdfc.Scaling.Orientation = _g.NewCT_Orientation()
	_bdfc.Scaling.Orientation.ValAttr = _g.ST_OrientationMinMax
	_bdfc.Choice = &_g.EG_AxSharedChoice{}
	_bdfc.Choice.Crosses = _g.NewCT_Crosses()
	_bdfc.Choice.Crosses.ValAttr = _g.ST_CrossesAutoZero
	_bfg := DateAxis{_bdfc}
	_bfg.MajorGridLines().Properties().LineProperties().SetSolidFill(_ef.LightGray)
	_bfg.SetMajorTickMark(_g.ST_TickMarkOut)
	_bfg.SetMinorTickMark(_g.ST_TickMarkIn)
	_bfg.SetTickLabelPosition(_g.ST_TickLblPosNextTo)
	_bfg.Properties().LineProperties().SetSolidFill(_ef.Black)
	_bfg.SetPosition(_g.ST_AxPosL)
	return _bfg
}

// AddPieOfPieChart adds a new pie chart to a chart.
func (_faec Chart) AddPieOfPieChart() PieOfPieChart {
	_fcb := _g.NewCT_PlotAreaChoice()
	_faec._aga.Chart.PlotArea.Choice = append(_faec._aga.Chart.PlotArea.Choice, _fcb)
	_fcb.OfPieChart = _g.NewCT_OfPieChart()
	_gf := PieOfPieChart{_cbff: _fcb.OfPieChart}
	_gf.InitializeDefaults()
	return _gf
}

// SetText sets the series text.
func (_agc RadarChartSeries) SetText(s string) {
	_agc._fbg.Tx = _g.NewCT_SerTx()
	_agc._fbg.Tx.Choice.V = &s
}

// AddValueAxis adds a value axis to the chart.
func (_bge Chart) AddValueAxis() ValueAxis {
	_bgc := _g.NewCT_ValAx()
	if _bge._aga.Chart.PlotArea.CChoice == nil {
		_bge._aga.Chart.PlotArea.CChoice = _g.NewCT_PlotAreaChoice1()
	}
	_bgc.AxId = _g.NewCT_UnsignedInt()
	_bgc.AxId.ValAttr = 0x7FFFFFFF & _gd.Uint32()
	_bge._aga.Chart.PlotArea.CChoice.ValAx = append(_bge._aga.Chart.PlotArea.CChoice.ValAx, _bgc)
	_bgc.Delete = _g.NewCT_Boolean()
	_bgc.Delete.ValAttr = _f.Bool(false)
	_bgc.Scaling = _g.NewCT_Scaling()
	_bgc.Scaling.Orientation = _g.NewCT_Orientation()
	_bgc.Scaling.Orientation.ValAttr = _g.ST_OrientationMinMax
	_bgc.Choice = &_g.EG_AxSharedChoice{}
	_bgc.Choice.Crosses = _g.NewCT_Crosses()
	_bgc.Choice.Crosses.ValAttr = _g.ST_CrossesAutoZero
	_bgc.CrossBetween = _g.NewCT_CrossBetween()
	_bgc.CrossBetween.ValAttr = _g.ST_CrossBetweenBetween
	_gabb := MakeValueAxis(_bgc)
	_gabb.MajorGridLines().Properties().LineProperties().SetSolidFill(_ef.LightGray)
	_gabb.SetMajorTickMark(_g.ST_TickMarkOut)
	_gabb.SetMinorTickMark(_g.ST_TickMarkIn)
	_gabb.SetTickLabelPosition(_g.ST_TickLblPosNextTo)
	_gabb.Properties().LineProperties().SetSolidFill(_ef.Black)
	_gabb.SetPosition(_g.ST_AxPosL)
	return _gabb
}

// X returns the inner wrapped XML type.
func (_agg SurfaceChart) X() *_g.CT_SurfaceChart { return _agg._fedb }
func (_fbd ValueAxis) Properties() _ge.ShapeProperties {
	if _fbd._gbg.SpPr == nil {
		_fbd._gbg.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_fbd._gbg.SpPr)
}

// Order returns the order of the series
func (_egg LineChartSeries) Order() uint32 { return _egg._feg.Order.ValAttr }
func (_df nullAxis) AxisID() uint32        { return 0 }
func (_fcgb ScatterChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _fcgb._bebg.XVal == nil {
		_fcgb._bebg.XVal = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_fcgb._bebg.XVal)
}

// AddAreaChart adds a new area chart to a chart.
func (_bdbd Chart) AddAreaChart() AreaChart {
	_cgda := _g.NewCT_PlotAreaChoice()
	_bdbd._aga.Chart.PlotArea.Choice = append(_bdbd._aga.Chart.PlotArea.Choice, _cgda)
	_cgda.AreaChart = _g.NewCT_AreaChart()
	_ba := AreaChart{_cb: _cgda.AreaChart}
	_ba.InitializeDefaults()
	return _ba
}

// SetText sets the series text
func (_dgg ScatterChartSeries) SetText(s string) {
	_dgg._bebg.Tx = _g.NewCT_SerTx()
	_dgg._bebg.Tx.Choice.V = &s
}

type GridLines struct{ _gee *_g.CT_ChartLines }

// RemoveLegend removes the legend if the chart has one.
func (_eac Chart) RemoveLegend() { _eac._aga.Chart.Legend = nil }
func (_dcfc DateAxis) MajorGridLines() GridLines {
	if _dcfc._bfe.MajorGridlines == nil {
		_dcfc._bfe.MajorGridlines = _g.NewCT_ChartLines()
	}
	return GridLines{_dcfc._bfe.MajorGridlines}
}
func (_fd Area3DChart) AddAxis(axis Axis) {
	_eb := _g.NewCT_UnsignedInt()
	_eb.ValAttr = axis.AxisID()
	_fd._b.AxId = append(_fd._b.AxId, _eb)
}
func (_fgd ValueAxis) SetCrosses(axis Axis) { _fgd._gbg.CrossAx.ValAttr = axis.AxisID() }
func (_gfa DataLabels) ensureChoice() {
	if _gfa._gebb.Choice == nil {
		_gfa._gebb.Choice = _g.NewCT_DLblsChoice()
	}
}

// Values returns the value data source.
func (_fae BubbleChartSeries) Values() NumberDataSource {
	if _fae._cf.YVal == nil {
		_fae._cf.YVal = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_fae._cf.YVal)
}

// ScatterChartSeries is the data series for a scatter chart.
type ScatterChartSeries struct{ _bebg *_g.CT_ScatterSer }
type Title struct{ _gga *_g.CT_Title }

// AddStockChart adds a new stock chart.
func (_gcc Chart) AddStockChart() StockChart {
	_cff := _g.NewCT_PlotAreaChoice()
	_gcc._aga.Chart.PlotArea.Choice = append(_gcc._aga.Chart.PlotArea.Choice, _cff)
	_cff.StockChart = _g.NewCT_StockChart()
	_edg := StockChart{_bgfg: _cff.StockChart}
	_edg.InitializeDefaults()
	return _edg
}

// AddAxis adds an axis to a line chart.
func (_bcd LineChart) AddAxis(axis Axis) {
	_ebe := _g.NewCT_UnsignedInt()
	_ebe.ValAttr = axis.AxisID()
	_bcd._cbc.AxId = append(_bcd._cbc.AxId, _ebe)
}

// AddSeries adds a default series to an Doughnut chart.
func (_ege DoughnutChart) AddSeries() PieChartSeries {
	_cec := _g.NewCT_PieSer()
	_ege._bae.Ser = append(_ege._bae.Ser, _cec)
	_cec.Idx.ValAttr = uint32(len(_ege._bae.Ser) - 1)
	_cec.Order.ValAttr = uint32(len(_ege._bae.Ser) - 1)
	_gff := PieChartSeries{_cec}
	_gff.InitializeDefaults()
	return _gff
}

// SetText sets the series text.
func (_aeff PieChartSeries) SetText(s string) {
	_aeff._eff.Tx = _g.NewCT_SerTx()
	_aeff._eff.Tx.Choice.V = &s
}

// AddSeries adds a default series to a Surface chart.
func (_acbg SurfaceChart) AddSeries() SurfaceChartSeries {
	_ecf := _acbg.nextColor(len(_acbg._fedb.Ser))
	_fecd := _g.NewCT_SurfaceSer()
	_acbg._fedb.Ser = append(_acbg._fedb.Ser, _fecd)
	_fecd.Idx.ValAttr = uint32(len(_acbg._fedb.Ser) - 1)
	_fecd.Order.ValAttr = uint32(len(_acbg._fedb.Ser) - 1)
	_eccd := SurfaceChartSeries{_fecd}
	_eccd.InitializeDefaults()
	_eccd.Properties().LineProperties().SetSolidFill(_ecf)
	return _eccd
}

// X returns the inner wrapped XML type.
func (_gda DateAxis) X() *_g.CT_DateAx { return _gda._bfe }
func (_dad CategoryAxis) SetMajorTickMark(m _g.ST_TickMark) {
	if m == _g.ST_TickMarkUnset {
		_dad._bed.MajorTickMark = nil
	} else {
		_dad._bed.MajorTickMark = _g.NewCT_TickMark()
		_dad._bed.MajorTickMark.ValAttr = m
	}
}

// InitializeDefaults the bar chart to its defaults
func (_gg Area3DChart) InitializeDefaults() {}

// CategoryAxisDataSource specifies the data for an axis.  It's commonly used with
// SetReference to set the axis data to a range of cells.
type CategoryAxisDataSource struct{ _ed *_g.CT_AxDataSource }

func MakeNumberDataSource(x *_g.CT_NumDataSource) NumberDataSource { return NumberDataSource{x} }

// Pie3DChart is a Pie3D chart.
type Pie3DChart struct {
	chartBase
	_ebfb *_g.CT_Pie3DChart
}

func (_bedg ValueAxis) SetMinorTickMark(m _g.ST_TickMark) {
	if m == _g.ST_TickMarkUnset {
		_bedg._gbg.MinorTickMark = nil
	} else {
		_bedg._gbg.MinorTickMark = _g.NewCT_TickMark()
		_bedg._gbg.MinorTickMark.ValAttr = m
	}
}
func (_cfeb NumberDataSource) SetReference(s string) {
	_cfeb.ensureChoice()
	if _cfeb._ffd.Choice.NumRef == nil {
		_cfeb._ffd.Choice.NumRef = _g.NewCT_NumRef()
	}
	_cfeb._ffd.Choice.NumRef.F = s
}

// SetValues is used to set the source data to a set of values.
func (_aff CategoryAxisDataSource) SetValues(v []string) {
	_aff._ed.Choice = _g.NewCT_AxDataSourceChoice()
	_aff._ed.Choice.StrLit = _g.NewCT_StrData()
	_aff._ed.Choice.StrLit.PtCount = _g.NewCT_UnsignedInt()
	_aff._ed.Choice.StrLit.PtCount.ValAttr = uint32(len(v))
	for _cdb, _cfd := range v {
		_aff._ed.Choice.StrLit.Pt = append(_aff._ed.Choice.StrLit.Pt, &_g.CT_StrVal{IdxAttr: uint32(_cdb), V: _cfd})
	}
}

type CategoryAxis struct{ _bed *_g.CT_CatAx }

// X returns the inner wrapped XML type.
func (_fg Area3DChart) X() *_g.CT_Area3DChart { return _fg._b }

// Marker returns the marker properties.
func (_ffcb ScatterChartSeries) Marker() Marker {
	if _ffcb._bebg.Marker == nil {
		_ffcb._bebg.Marker = _g.NewCT_Marker()
	}
	return MakeMarker(_ffcb._bebg.Marker)
}

// Values returns the value data source.
func (_dga PieChartSeries) Values() NumberDataSource {
	if _dga._eff.Val == nil {
		_dga._eff.Val = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_dga._eff.Val)
}
func (_bcb Marker) SetSymbol(s _g.ST_MarkerStyle) {
	if s == _g.ST_MarkerStyleUnset {
		_bcb._eed.Symbol = nil
	} else {
		_bcb._eed.Symbol = _g.NewCT_MarkerStyle()
		_bcb._eed.Symbol.ValAttr = s
	}
}
func (_abf ValueAxis) AxisID() uint32 { return _abf._gbg.AxId.ValAttr }

// CreateEmptyNumberCache creates an empty number cache, which is used sometimes
// to increase file format compatibility.  It should actually contain the
// computed cell data, but just creating an empty one is good enough.
func (_ggd NumberDataSource) CreateEmptyNumberCache() {
	_ggd.ensureChoice()
	if _ggd._ffd.Choice.NumRef == nil {
		_ggd._ffd.Choice.NumRef = _g.NewCT_NumRef()
	}
	_ggd._ffd.Choice.NumLit = nil
	_ggd._ffd.Choice.NumRef.NumCache = _g.NewCT_NumData()
	_ggd._ffd.Choice.NumRef.NumCache.PtCount = _g.NewCT_UnsignedInt()
	_ggd._ffd.Choice.NumRef.NumCache.PtCount.ValAttr = 0
}

// X returns the inner wrapped XML type.
func (_ddde SurfaceChartSeries) X() *_g.CT_SurfaceSer { return _ddde._dee }
func (_cedc Title) SetText(s string) {
	if _cedc._gga.Tx == nil {
		_cedc._gga.Tx = _g.NewCT_Tx()
	}
	if _cedc._gga.Tx.Choice.Rich == nil {
		_cedc._gga.Tx.Choice.Rich = _ga.NewCT_TextBody()
	}
	var _fdeg *_ga.CT_TextParagraph
	if len(_cedc._gga.Tx.Choice.Rich.P) == 0 {
		_fdeg = _ga.NewCT_TextParagraph()
		_cedc._gga.Tx.Choice.Rich.P = []*_ga.CT_TextParagraph{_fdeg}
	} else {
		_fdeg = _cedc._gga.Tx.Choice.Rich.P[0]
	}
	var _bcf *_ga.EG_TextRun
	if len(_fdeg.EG_TextRun) == 0 {
		_bcf = _ga.NewEG_TextRun()
		_fdeg.EG_TextRun = []*_ga.EG_TextRun{_bcf}
	} else {
		_bcf = _fdeg.EG_TextRun[0]
	}
	if _bcf.R == nil {
		_bcf.R = _ga.NewCT_RegularTextRun()
	}
	_bcf.R.T = s
}

// PieChart is a Pie chart.
type PieChart struct {
	chartBase
	_dgc *_g.CT_PieChart
}

// AddDoughnutChart adds a new doughnut (pie with a hole in the center) chart to a chart.
func (_gcff Chart) AddDoughnutChart() DoughnutChart {
	_fgga := _g.NewCT_PlotAreaChoice()
	_gcff._aga.Chart.PlotArea.Choice = append(_gcff._aga.Chart.PlotArea.Choice, _fgga)
	_fgga.DoughnutChart = _g.NewCT_DoughnutChart()
	_cfa := DoughnutChart{_bae: _fgga.DoughnutChart}
	_cfa.InitializeDefaults()
	return _cfa
}

// AddSeries adds a default series to an area chart.
func (_a Area3DChart) AddSeries() AreaChartSeries {
	_ee := _a.nextColor(len(_a._b.Ser))
	_db := _g.NewCT_AreaSer()
	_a._b.Ser = append(_a._b.Ser, _db)
	_db.Idx.ValAttr = uint32(len(_a._b.Ser) - 1)
	_db.Order.ValAttr = uint32(len(_a._b.Ser) - 1)
	_cg := AreaChartSeries{_db}
	_cg.InitializeDefaults()
	_cg.Properties().SetSolidFill(_ee)
	return _cg
}

// SetIndex sets the index of the series
func (_cfdgc ScatterChartSeries) SetIndex(idx uint32) { _cfdgc._bebg.Idx.ValAttr = idx }

// AddPie3DChart adds a new pie chart to a chart.
func (_baa Chart) AddPie3DChart() Pie3DChart {
	_edf(_baa._aga.Chart)
	_aeb := _g.NewCT_PlotAreaChoice()
	_baa._aga.Chart.PlotArea.Choice = append(_baa._aga.Chart.PlotArea.Choice, _aeb)
	_aeb.Pie3DChart = _g.NewCT_Pie3DChart()
	_agbg := Pie3DChart{_ebfb: _aeb.Pie3DChart}
	_agbg.InitializeDefaults()
	return _agbg
}

// AddSeries adds a default series to an area chart.
func (_dd AreaChart) AddSeries() AreaChartSeries {
	_cga := _dd.nextColor(len(_dd._cb.Ser))
	_af := _g.NewCT_AreaSer()
	_dd._cb.Ser = append(_dd._cb.Ser, _af)
	_af.Idx.ValAttr = uint32(len(_dd._cb.Ser) - 1)
	_af.Order.ValAttr = uint32(len(_dd._cb.Ser) - 1)
	_bg := AreaChartSeries{_af}
	_bg.InitializeDefaults()
	_bg.Properties().SetSolidFill(_cga)
	return _bg
}

// InitializeDefaults initializes an Pie series to the default values.
func (_dbad PieChartSeries) InitializeDefaults() {}
func (_cac Title) InitializeDefaults() {
	_cac.SetText("\u0054\u0069\u0074l\u0065")
	_cac.RunProperties().SetSize(16 * _e.Point)
	_cac.RunProperties().SetSolidFill(_ef.Black)
	_cac.RunProperties().SetFont("\u0043\u0061\u006c\u0069\u0062\u0020\u0072\u0069")
	_cac.RunProperties().SetBold(false)
}

// X returns the inner wrapped XML type.
func (_ggbf SeriesAxis) X() *_g.CT_SerAx { return _ggbf._dgdc }

// InitializeDefaults the bar chart to its defaults
func (_bbdf DoughnutChart) InitializeDefaults() {
	_bbdf._bae.VaryColors = _g.NewCT_Boolean()
	_bbdf._bae.VaryColors.ValAttr = _f.Bool(true)
	_bbdf._bae.HoleSize = _g.NewCT_HoleSize()
	_bbdf._bae.HoleSize.ValAttr = &_g.ST_HoleSize{}
	_bbdf._bae.HoleSize.ValAttr.ST_HoleSizeUByte = _f.Uint8(50)
}

// SetDirection changes the direction of the bar chart (bar or column).
func (_dc Bar3DChart) SetDirection(d _g.ST_BarDir) { _dc._fgf.BarDir.ValAttr = d }

// SetHoleSize controls the hole size in the pie chart and is measured in percent.
func (_cbf DoughnutChart) SetHoleSize(pct uint8) {
	if _cbf._bae.HoleSize == nil {
		_cbf._bae.HoleSize = _g.NewCT_HoleSize()
	}
	if _cbf._bae.HoleSize.ValAttr == nil {
		_cbf._bae.HoleSize.ValAttr = &_g.ST_HoleSize{}
	}
	_cbf._bae.HoleSize.ValAttr.ST_HoleSizeUByte = &pct
}

// SetOrder sets the order of the series
func (_deg ScatterChartSeries) SetOrder(idx uint32) { _deg._bebg.Order.ValAttr = idx }

// AddSeries adds a default series to a Surface chart.
func (_fggb Surface3DChart) AddSeries() SurfaceChartSeries {
	_gbd := _fggb.nextColor(len(_fggb._eacb.Ser))
	_gecf := _g.NewCT_SurfaceSer()
	_fggb._eacb.Ser = append(_fggb._eacb.Ser, _gecf)
	_gecf.Idx.ValAttr = uint32(len(_fggb._eacb.Ser) - 1)
	_gecf.Order.ValAttr = uint32(len(_fggb._eacb.Ser) - 1)
	_fcdg := SurfaceChartSeries{_gecf}
	_fcdg.InitializeDefaults()
	_fcdg.Properties().LineProperties().SetSolidFill(_gbd)
	return _fcdg
}
func (_ecb DataLabels) SetShowLeaderLines(b bool) {
	_ecb.ensureChoice()
	_ecb._gebb.Choice.ShowLeaderLines = _g.NewCT_Boolean()
	_ecb._gebb.Choice.ShowLeaderLines.ValAttr = _f.Bool(b)
}

// AddBubbleChart adds a new bubble chart.
func (_edc Chart) AddBubbleChart() BubbleChart {
	_de := _g.NewCT_PlotAreaChoice()
	_edc._aga.Chart.PlotArea.Choice = append(_edc._aga.Chart.PlotArea.Choice, _de)
	_de.BubbleChart = _g.NewCT_BubbleChart()
	_ebd := BubbleChart{_da: _de.BubbleChart}
	_ebd.InitializeDefaults()
	return _ebd
}

// Properties returns the bar chart series shape properties.
func (_dbadf PieChartSeries) Properties() _ge.ShapeProperties {
	if _dbadf._eff.SpPr == nil {
		_dbadf._eff.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_dbadf._eff.SpPr)
}

// Properties returns the chart's shape properties.
func (_dfa Chart) Properties() _ge.ShapeProperties {
	if _dfa._aga.SpPr == nil {
		_dfa._aga.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_dfa._aga.SpPr)
}

// LineChartSeries is the data series for a line chart.
type LineChartSeries struct{ _feg *_g.CT_LineSer }

func MakeLegend(l *_g.CT_Legend) Legend { return Legend{l} }
func (_bc GridLines) Properties() _ge.ShapeProperties {
	if _bc._gee.SpPr == nil {
		_bc._gee.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_bc._gee.SpPr)
}
func (_dba AreaChart) AddAxis(axis Axis) {
	_ag := _g.NewCT_UnsignedInt()
	_ag.ValAttr = axis.AxisID()
	_dba._cb.AxId = append(_dba._cb.AxId, _ag)
}
func (_ac CategoryAxis) SetCrosses(axis Axis) {
	_ac._bed.Choice = _g.NewEG_AxSharedChoice()
	_ac._bed.Choice.Crosses = _g.NewCT_Crosses()
	_ac._bed.Choice.Crosses.ValAttr = _g.ST_CrossesAutoZero
	_ac._bed.CrossAx.ValAttr = axis.AxisID()
}

// AddSeries adds a default series to a line chart.
func (_aede LineChart) AddSeries() LineChartSeries {
	_eaf := _aede.nextColor(len(_aede._cbc.Ser))
	_gef := _g.NewCT_LineSer()
	_aede._cbc.Ser = append(_aede._cbc.Ser, _gef)
	_gef.Idx.ValAttr = uint32(len(_aede._cbc.Ser) - 1)
	_gef.Order.ValAttr = uint32(len(_aede._cbc.Ser) - 1)
	_eeg := LineChartSeries{_gef}
	_eeg.InitializeDefaults()
	_eeg.Properties().LineProperties().SetSolidFill(_eaf)
	return _eeg
}

// AddLine3DChart adds a new 3D line chart to a chart.
func (_ae Chart) AddLine3DChart() Line3DChart {
	_edf(_ae._aga.Chart)
	_fgfg := _g.NewCT_PlotAreaChoice()
	_ae._aga.Chart.PlotArea.Choice = append(_ae._aga.Chart.PlotArea.Choice, _fgfg)
	_fgfg.Line3DChart = _g.NewCT_Line3DChart()
	_fgfg.Line3DChart.Grouping = _g.NewCT_Grouping()
	_fgfg.Line3DChart.Grouping.ValAttr = _g.ST_GroupingStandard
	return Line3DChart{_begf: _fgfg.Line3DChart}
}

// CategoryAxis returns the category data source.
func (_abe RadarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _abe._fbg.Cat == nil {
		_abe._fbg.Cat = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_abe._fbg.Cat)
}
func (_dec DataLabels) SetShowLegendKey(b bool) {
	_dec.ensureChoice()
	_dec._gebb.Choice.ShowLegendKey = _g.NewCT_Boolean()
	_dec._gebb.Choice.ShowLegendKey.ValAttr = _f.Bool(b)
}

// AddSeries adds a default series to a line chart.
func (_cfff Line3DChart) AddSeries() LineChartSeries {
	_gecc := _cfff.nextColor(len(_cfff._begf.Ser))
	_fcda := _g.NewCT_LineSer()
	_cfff._begf.Ser = append(_cfff._begf.Ser, _fcda)
	_fcda.Idx.ValAttr = uint32(len(_cfff._begf.Ser) - 1)
	_fcda.Order.ValAttr = uint32(len(_cfff._begf.Ser) - 1)
	_fee := LineChartSeries{_fcda}
	_fee.InitializeDefaults()
	_fee.Properties().LineProperties().SetSolidFill(_gecc)
	_fee.Properties().SetSolidFill(_gecc)
	return _fee
}

// InitializeDefaults the Stock chart to its defaults
func (_ccee StockChart) InitializeDefaults() {
	_ccee._bgfg.HiLowLines = _g.NewCT_ChartLines()
	_ccee._bgfg.UpDownBars = _g.NewCT_UpDownBars()
	_ccee._bgfg.UpDownBars.GapWidth = _g.NewCT_GapAmount()
	_ccee._bgfg.UpDownBars.GapWidth.ValAttr = &_g.ST_GapAmount{}
	_ccee._bgfg.UpDownBars.GapWidth.ValAttr.ST_GapAmountUShort = _f.Uint16(150)
	_ccee._bgfg.UpDownBars.UpBars = _g.NewCT_UpDownBar()
	_ccee._bgfg.UpDownBars.DownBars = _g.NewCT_UpDownBar()
}

// SurfaceChart is a 3D surface chart, viewed from the top-down.
type SurfaceChart struct {
	chartBase
	_fedb *_g.CT_SurfaceChart
}

// SetOrder sets the order of the series
func (_dcfg SurfaceChartSeries) SetOrder(idx uint32) { _dcfg._dee.Order.ValAttr = idx }

// AddCategoryAxis adds a category axis.
func (_ead Chart) AddCategoryAxis() CategoryAxis {
	_cbac := _g.NewCT_CatAx()
	if _ead._aga.Chart.PlotArea.CChoice == nil {
		_ead._aga.Chart.PlotArea.CChoice = _g.NewCT_PlotAreaChoice1()
	}
	_cbac.AxId = _g.NewCT_UnsignedInt()
	_cbac.AxId.ValAttr = 0x7FFFFFFF & _gd.Uint32()
	_ead._aga.Chart.PlotArea.CChoice.CatAx = append(_ead._aga.Chart.PlotArea.CChoice.CatAx, _cbac)
	_cbac.Auto = _g.NewCT_Boolean()
	_cbac.Auto.ValAttr = _f.Bool(true)
	_cbac.Delete = _g.NewCT_Boolean()
	_cbac.Delete.ValAttr = _f.Bool(false)
	_baad := MakeCategoryAxis(_cbac)
	_baad.InitializeDefaults()
	return _baad
}

type ValueAxis struct{ _gbg *_g.CT_ValAx }

func MakeTitle(x *_g.CT_Title) Title    { return Title{x} }
func (_abef SeriesAxis) AxisID() uint32 { return _abef._dgdc.AxId.ValAttr }
func (_ddcb CategoryAxis) SetPosition(p _g.ST_AxPos) {
	_ddcb._bed.AxPos = _g.NewCT_AxPos()
	_ddcb._bed.AxPos.ValAttr = p
}

// InitializeDefaults initializes a bar chart series to the default values.
func (_ffa BarChartSeries) InitializeDefaults() {}
func (_aa BubbleChart) AddAxis(axis Axis) {
	_dfb := _g.NewCT_UnsignedInt()
	_dfb.ValAttr = axis.AxisID()
	_aa._da.AxId = append(_aa._da.AxId, _dfb)
}

// X returns the inner wrapped XML type.
func (_ffdc Surface3DChart) X() *_g.CT_Surface3DChart { return _ffdc._eacb }

// MakeAxisDataSource constructs an AxisDataSource wrapper.
func MakeAxisDataSource(x *_g.CT_AxDataSource) CategoryAxisDataSource {
	return CategoryAxisDataSource{x}
}

// X returns the inner wrapped XML type.
func (_bda PieChart) X() *_g.CT_PieChart { return _bda._dgc }
func (_fdd DateAxis) SetMinorTickMark(m _g.ST_TickMark) {
	if m == _g.ST_TickMarkUnset {
		_fdd._bfe.MinorTickMark = nil
	} else {
		_fdd._bfe.MinorTickMark = _g.NewCT_TickMark()
		_fdd._bfe.MinorTickMark.ValAttr = m
	}
}
func (_dca Marker) SetSize(sz uint8) {
	_dca._eed.Size = _g.NewCT_MarkerSize()
	_dca._eed.Size.ValAttr = &sz
}

// X returns the inner wrapped XML type.
func (_dg AreaChartSeries) X() *_g.CT_AreaSer { return _dg._ddc }

// SetText sets the series text
func (_eab LineChartSeries) SetText(s string) {
	_eab._feg.Tx = _g.NewCT_SerTx()
	_eab._feg.Tx.Choice.V = &s
}
func (_cba CategoryAxis) SetTickLabelPosition(p _g.ST_TickLblPos) {
	if p == _g.ST_TickLblPosUnset {
		_cba._bed.TickLblPos = nil
	} else {
		_cba._bed.TickLblPos = _g.NewCT_TickLblPos()
		_cba._bed.TickLblPos.ValAttr = p
	}
}
func (_gcg DataLabels) SetShowCategoryName(b bool) {
	_gcg.ensureChoice()
	_gcg._gebb.Choice.ShowCatName = _g.NewCT_Boolean()
	_gcg._gebb.Choice.ShowCatName.ValAttr = _f.Bool(b)
}

// AddAxis adds an axis to a line chart.
func (_adg Line3DChart) AddAxis(axis Axis) {
	_cdbf := _g.NewCT_UnsignedInt()
	_cdbf.ValAttr = axis.AxisID()
	_adg._begf.AxId = append(_adg._begf.AxId, _cdbf)
}

// Properties returns the line chart series shape properties.
func (_ggf SurfaceChartSeries) Properties() _ge.ShapeProperties {
	if _ggf._dee.SpPr == nil {
		_ggf._dee.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_ggf._dee.SpPr)
}
func (_baada LineChartSeries) InitializeDefaults() {
	_baada.Properties().LineProperties().SetWidth(1 * _e.Point)
	_baada.Properties().LineProperties().SetSolidFill(_ef.Black)
	_baada.Properties().LineProperties().SetJoin(_ge.LineJoinRound)
	_baada.Marker().SetSymbol(_g.ST_MarkerStyleNone)
	_baada.Labels().SetShowLegendKey(false)
	_baada.Labels().SetShowValue(false)
	_baada.Labels().SetShowPercent(false)
	_baada.Labels().SetShowCategoryName(false)
	_baada.Labels().SetShowSeriesName(false)
	_baada.Labels().SetShowLeaderLines(false)
}
func (_eaa RadarChart) AddAxis(axis Axis) {
	_bbf := _g.NewCT_UnsignedInt()
	_bbf.ValAttr = axis.AxisID()
	_eaa._gaa.AxId = append(_eaa._gaa.AxId, _bbf)
}

// X returns the inner wrapped XML type.
func (_fb Bar3DChart) X() *_g.CT_Bar3DChart { return _fb._fgf }
func (_abc ValueAxis) SetMajorTickMark(m _g.ST_TickMark) {
	if m == _g.ST_TickMarkUnset {
		_abc._gbg.MajorTickMark = nil
	} else {
		_abc._gbg.MajorTickMark = _g.NewCT_TickMark()
		_abc._gbg.MajorTickMark.ValAttr = m
	}
}

// X returns the inner wrapped XML type.
func (_cd AreaChart) X() *_g.CT_AreaChart { return _cd._cb }
func (_cgc ScatterChartSeries) InitializeDefaults() {
	_cgc.Properties().LineProperties().SetNoFill()
	_cgc.Marker().SetSymbol(_g.ST_MarkerStyleAuto)
	_cgc.Labels().SetShowLegendKey(false)
	_cgc.Labels().SetShowValue(true)
	_cgc.Labels().SetShowPercent(false)
	_cgc.Labels().SetShowCategoryName(false)
	_cgc.Labels().SetShowSeriesName(false)
	_cgc.Labels().SetShowLeaderLines(false)
}

type SurfaceChartSeries struct{ _dee *_g.CT_SurfaceSer }
type Legend struct{ _cag *_g.CT_Legend }

// AreaChartSeries is a series to be used on an area chart.
type AreaChartSeries struct{ _ddc *_g.CT_AreaSer }

// CategoryAxis returns the category data source.
func (_gagc BubbleChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _gagc._cf.XVal == nil {
		_gagc._cf.XVal = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_gagc._cf.XVal)
}
func (_gcdf ScatterChartSeries) Values() NumberDataSource {
	if _gcdf._bebg.YVal == nil {
		_gcdf._bebg.YVal = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_gcdf._bebg.YVal)
}

// Order returns the order of the series
func (_faaa SurfaceChartSeries) Order() uint32 { return _faaa._dee.Order.ValAttr }

// AddAxis adds an axis to a Surface chart.
func (_bfbf Surface3DChart) AddAxis(axis Axis) {
	_gcfa := _g.NewCT_UnsignedInt()
	_gcfa.ValAttr = axis.AxisID()
	_bfbf._eacb.AxId = append(_bfbf._eacb.AxId, _gcfa)
}

// Values returns the value data source.
func (_bfad RadarChartSeries) Values() NumberDataSource {
	if _bfad._fbg.Val == nil {
		_bfad._fbg.Val = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_bfad._fbg.Val)
}
func (_dabg NumberDataSource) ensureChoice() {
	if _dabg._ffd.Choice == nil {
		_dabg._ffd.Choice = _g.NewCT_NumDataSourceChoice()
	}
}

// Properties returns the line chart series shape properties.
func (_gfg ScatterChartSeries) Properties() _ge.ShapeProperties {
	if _gfg._bebg.SpPr == nil {
		_gfg._bebg.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_gfg._bebg.SpPr)
}
func (_gacd Title) RunProperties() _ge.RunProperties {
	if _gacd._gga.Tx == nil {
		_gacd.SetText("")
	}
	if _gacd._gga.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr == nil {
		_gacd._gga.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr = _ga.NewCT_TextCharacterProperties()
	}
	return _ge.MakeRunProperties(_gacd._gga.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr)
}

// AddSeries adds a default series to an Pie3D chart.
func (_bbb Pie3DChart) AddSeries() PieChartSeries {
	_deb := _g.NewCT_PieSer()
	_bbb._ebfb.Ser = append(_bbb._ebfb.Ser, _deb)
	_deb.Idx.ValAttr = uint32(len(_bbb._ebfb.Ser) - 1)
	_deb.Order.ValAttr = uint32(len(_bbb._ebfb.Ser) - 1)
	_bcc := PieChartSeries{_deb}
	_bcc.InitializeDefaults()
	return _bcc
}

// AddRadarChart adds a new radar chart to a chart.
func (_dac Chart) AddRadarChart() RadarChart {
	_bdf := _g.NewCT_PlotAreaChoice()
	_dac._aga.Chart.PlotArea.Choice = append(_dac._aga.Chart.PlotArea.Choice, _bdf)
	_bdf.RadarChart = _g.NewCT_RadarChart()
	_affd := RadarChart{_gaa: _bdf.RadarChart}
	_affd.InitializeDefaults()
	return _affd
}

// Axis is the interface implemented by different axes when assigning to a
// chart.
type Axis interface{ AxisID() uint32 }

func (_acge ValueAxis) SetTickLabelPosition(p _g.ST_TickLblPos) {
	if p == _g.ST_TickLblPosUnset {
		_acge._gbg.TickLblPos = nil
	} else {
		_acge._gbg.TickLblPos = _g.NewCT_TickLblPos()
		_acge._gbg.TickLblPos.ValAttr = p
	}
}

// X returns the inner wrapped XML type.
func (_acb Chart) X() *_g.ChartSpace { return _acb._aga }

// X returns the inner wrapped XML type.
func (_gdg ValueAxis) X() *_g.CT_ValAx  { return _gdg._gbg }
func MakeMarker(x *_g.CT_Marker) Marker { return Marker{x} }

// AddSeries adds a default series to an Radar chart.
func (_bdc RadarChart) AddSeries() RadarChartSeries {
	_gccf := _bdc.nextColor(len(_bdc._gaa.Ser))
	_gdfb := _g.NewCT_RadarSer()
	_bdc._gaa.Ser = append(_bdc._gaa.Ser, _gdfb)
	_gdfb.Idx.ValAttr = uint32(len(_bdc._gaa.Ser) - 1)
	_gdfb.Order.ValAttr = uint32(len(_bdc._gaa.Ser) - 1)
	_edfa := RadarChartSeries{_gdfb}
	_edfa.InitializeDefaults()
	_edfa.Properties().SetSolidFill(_gccf)
	return _edfa
}
func (_gca ValueAxis) MajorGridLines() GridLines {
	if _gca._gbg.MajorGridlines == nil {
		_gca._gbg.MajorGridlines = _g.NewCT_ChartLines()
	}
	return GridLines{_gca._gbg.MajorGridlines}
}
func (_be Bar3DChart) AddAxis(axis Axis) {
	_agb := _g.NewCT_UnsignedInt()
	_agb.ValAttr = axis.AxisID()
	_be._fgf.AxId = append(_be._fgf.AxId, _agb)
}

// Area3DChart is an area chart that has a shaded area underneath a curve.
type Area3DChart struct {
	chartBase
	_b *_g.CT_Area3DChart
}

// Index returns the index of the series
func (_acbb SurfaceChartSeries) Index() uint32 { return _acbb._dee.Idx.ValAttr }

// Properties returns the bar chart series shape properties.
func (_bcg RadarChartSeries) Properties() _ge.ShapeProperties {
	if _bcg._fbg.SpPr == nil {
		_bcg._fbg.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_bcg._fbg.SpPr)
}

// X returns the inner wrapped XML type.
func (_cfaa LineChartSeries) X() *_g.CT_LineSer { return _cfaa._feg }
func (_dfad SurfaceChart) InitializeDefaults() {
	_dfad._fedb.Wireframe = _g.NewCT_Boolean()
	_dfad._fedb.Wireframe.ValAttr = _f.Bool(false)
	_dfad._fedb.BandFmts = _g.NewCT_BandFmts()
	for _afb := 0; _afb < 15; _afb++ {
		_eeb := _g.NewCT_BandFmt()
		_eeb.Idx.ValAttr = uint32(_afb)
		_eeb.SpPr = _ga.NewCT_ShapeProperties()
		_dfd := _ge.MakeShapeProperties(_eeb.SpPr)
		_dfd.SetSolidFill(_dfad.nextColor(_afb))
		_dfad._fedb.BandFmts.BandFmt = append(_dfad._fedb.BandFmts.BandFmt, _eeb)
	}
}

// AddSeries adds a default series to an Pie chart.
func (_gccc PieOfPieChart) AddSeries() PieChartSeries {
	_eag := _g.NewCT_PieSer()
	_gccc._cbff.Ser = append(_gccc._cbff.Ser, _eag)
	_eag.Idx.ValAttr = uint32(len(_gccc._cbff.Ser) - 1)
	_eag.Order.ValAttr = uint32(len(_gccc._cbff.Ser) - 1)
	_cfdg := PieChartSeries{_eag}
	_cfdg.InitializeDefaults()
	return _cfdg
}

// AddSeries adds a default series to a Bubble chart.
func (_bdb BubbleChart) AddSeries() BubbleChartSeries {
	_bbd := _bdb.nextColor(len(_bdb._da.Ser))
	_gcf := _g.NewCT_BubbleSer()
	_bdb._da.Ser = append(_bdb._da.Ser, _gcf)
	_gcf.Idx.ValAttr = uint32(len(_bdb._da.Ser) - 1)
	_gcf.Order.ValAttr = uint32(len(_bdb._da.Ser) - 1)
	_gag := BubbleChartSeries{_gcf}
	_gag.InitializeDefaults()
	_gag.Properties().SetSolidFill(_bbd)
	return _gag
}

// Values returns the value data source.
func (_ca AreaChartSeries) Values() NumberDataSource {
	if _ca._ddc.Val == nil {
		_ca._ddc.Val = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_ca._ddc.Val)
}
func (_ecc StockChart) AddAxis(axis Axis) {
	_dce := _g.NewCT_UnsignedInt()
	_dce.ValAttr = axis.AxisID()
	_ecc._bgfg.AxId = append(_ecc._bgfg.AxId, _dce)
}
func (_ce DataLabels) SetPosition(p _g.ST_DLblPos) {
	_ce.ensureChoice()
	_ce._gebb.Choice.DLblPos = _g.NewCT_DLblPos()
	_ce._gebb.Choice.DLblPos.ValAttr = p
}
func (_cfc Legend) InitializeDefaults() {
	_cfc.SetPosition(_g.ST_LegendPosR)
	_cfc.SetOverlay(false)
	_cfc.Properties().SetNoFill()
	_cfc.Properties().LineProperties().SetNoFill()
}

// Properties returns the Bubble chart series shape properties.
func (_ggb BubbleChartSeries) Properties() _ge.ShapeProperties {
	if _ggb._cf.SpPr == nil {
		_ggb._cf.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_ggb._cf.SpPr)
}
func (_gac BarChart) AddAxis(axis Axis) {
	_bf := _g.NewCT_UnsignedInt()
	_bf.ValAttr = axis.AxisID()
	_gac._bd.AxId = append(_gac._bd.AxId, _bf)
}

// X returns the inner wrapped XML type.
func (_acg Legend) X() *_g.CT_Legend { return _acg._cag }

// AddSeries adds a default series to a bar chart.
func (_cgd Bar3DChart) AddSeries() BarChartSeries {
	_efe := _cgd.nextColor(len(_cgd._fgf.Ser))
	_fde := _g.NewCT_BarSer()
	_cgd._fgf.Ser = append(_cgd._fgf.Ser, _fde)
	_fde.Idx.ValAttr = uint32(len(_cgd._fgf.Ser) - 1)
	_fde.Order.ValAttr = uint32(len(_cgd._fgf.Ser) - 1)
	_eec := BarChartSeries{_fde}
	_eec.InitializeDefaults()
	_eec.Properties().SetSolidFill(_efe)
	return _eec
}

// X returns the inner wrapped XML type.
func (_cdd LineChart) X() *_g.CT_LineChart { return _cdd._cbc }

// StockChart is a 2D Stock chart.
type StockChart struct {
	chartBase
	_bgfg *_g.CT_StockChart
}

// RemoveTitle removes any existing title from the chart.
func (_dadd Chart) RemoveTitle() {
	_dadd._aga.Chart.Title = nil
	_dadd._aga.Chart.AutoTitleDeleted = _g.NewCT_Boolean()
	_dadd._aga.Chart.AutoTitleDeleted.ValAttr = _f.Bool(true)
}
func (_ebf DateAxis) Properties() _ge.ShapeProperties {
	if _ebf._bfe.SpPr == nil {
		_ebf._bfe.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_ebf._bfe.SpPr)
}
func (_cbg CategoryAxis) Properties() _ge.ShapeProperties {
	if _cbg._bed.SpPr == nil {
		_cbg._bed.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_cbg._bed.SpPr)
}

// AddSeries adds a default series to an Pie chart.
func (_efb PieChart) AddSeries() PieChartSeries {
	_ffbe := _g.NewCT_PieSer()
	_efb._dgc.Ser = append(_efb._dgc.Ser, _ffbe)
	_ffbe.Idx.ValAttr = uint32(len(_efb._dgc.Ser) - 1)
	_ffbe.Order.ValAttr = uint32(len(_efb._dgc.Ser) - 1)
	_aef := PieChartSeries{_ffbe}
	_aef.InitializeDefaults()
	return _aef
}

// SetDisplayBlanksAs controls how missing values are displayed.
func (_ffc Chart) SetDisplayBlanksAs(v _g.ST_DispBlanksAs) {
	_ffc._aga.Chart.DispBlanksAs = _g.NewCT_DispBlanksAs()
	_ffc._aga.Chart.DispBlanksAs.ValAttr = v
}

// RadarChart is an Radar chart that has a shaded Radar underneath a curve.
type RadarChart struct {
	chartBase
	_gaa *_g.CT_RadarChart
}

// SetOrder sets the order of the series
func (_gece LineChartSeries) SetOrder(idx uint32) { _gece._feg.Order.ValAttr = idx }

// Properties returns the line chart series shape properties.
func (_ffb LineChartSeries) Properties() _ge.ShapeProperties {
	if _ffb._feg.SpPr == nil {
		_ffb._feg.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_ffb._feg.SpPr)
}
func (_fed Marker) Properties() _ge.ShapeProperties {
	if _fed._eed.SpPr == nil {
		_fed._eed.SpPr = _ga.NewCT_ShapeProperties()
	}
	return _ge.MakeShapeProperties(_fed._eed.SpPr)
}

// BubbleChart is a 2D Bubble chart.
type BubbleChart struct {
	chartBase
	_da *_g.CT_BubbleChart
}

func (_bfgb Title) ParagraphProperties() _ge.ParagraphProperties {
	if _bfgb._gga.Tx == nil {
		_bfgb.SetText("")
	}
	if _bfgb._gga.Tx.Choice.Rich.P[0].PPr == nil {
		_bfgb._gga.Tx.Choice.Rich.P[0].PPr = _ga.NewCT_TextParagraphProperties()
	}
	return _ge.MakeParagraphProperties(_bfgb._gga.Tx.Choice.Rich.P[0].PPr)
}
func (_dde DataLabels) SetShowSeriesName(b bool) {
	_dde.ensureChoice()
	_dde._gebb.Choice.ShowSerName = _g.NewCT_Boolean()
	_dde._gebb.Choice.ShowSerName.ValAttr = _f.Bool(b)
}
func _edf(_eea *_g.CT_Chart) {
	_eea.View3D = _g.NewCT_View3D()
	_eea.View3D.RotX = _g.NewCT_RotX()
	_eea.View3D.RotX.ValAttr = _f.Int8(15)
	_eea.View3D.RotY = _g.NewCT_RotY()
	_eea.View3D.RotY.ValAttr = _f.Uint16(20)
	_eea.View3D.RAngAx = _g.NewCT_Boolean()
	_eea.View3D.RAngAx.ValAttr = _f.Bool(false)
	_eea.Floor = _g.NewCT_Surface()
	_eea.Floor.Thickness = _g.NewCT_Thickness()
	_eea.Floor.Thickness.ValAttr.Uint32 = _f.Uint32(0)
	_eea.SideWall = _g.NewCT_Surface()
	_eea.SideWall.Thickness = _g.NewCT_Thickness()
	_eea.SideWall.Thickness.ValAttr.Uint32 = _f.Uint32(0)
	_eea.BackWall = _g.NewCT_Surface()
	_eea.BackWall.Thickness = _g.NewCT_Thickness()
	_eea.BackWall.Thickness.ValAttr.Uint32 = _f.Uint32(0)
}
func (_dcb LineChartSeries) Values() NumberDataSource {
	if _dcb._feg.Val == nil {
		_dcb._feg.Val = _g.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(_dcb._feg.Val)
}
func (_eba Surface3DChart) InitializeDefaults() {
	_eba._eacb.Wireframe = _g.NewCT_Boolean()
	_eba._eacb.Wireframe.ValAttr = _f.Bool(false)
	_eba._eacb.BandFmts = _g.NewCT_BandFmts()
	for _bfeb := 0; _bfeb < 15; _bfeb++ {
		_cdgf := _g.NewCT_BandFmt()
		_cdgf.Idx.ValAttr = uint32(_bfeb)
		_cdgf.SpPr = _ga.NewCT_ShapeProperties()
		_gdfd := _ge.MakeShapeProperties(_cdgf.SpPr)
		_gdfd.SetSolidFill(_eba.nextColor(_bfeb))
		_eba._eacb.BandFmts.BandFmt = append(_eba._eacb.BandFmts.BandFmt, _cdgf)
	}
}
func (_ced DataLabels) SetShowPercent(b bool) {
	_ced.ensureChoice()
	_ced._gebb.Choice.ShowPercent = _g.NewCT_Boolean()
	_ced._gebb.Choice.ShowPercent.ValAttr = _f.Bool(b)
}

var NullAxis Axis = nullAxis(0)

// X returns the inner wrapped XML type.
func (_bgcg Pie3DChart) X() *_g.CT_Pie3DChart { return _bgcg._ebfb }

// InitializeDefaults the bar chart to its defaults
func (_bb AreaChart) InitializeDefaults() {}

// AddScatterChart adds a scatter (X/Y) chart.
func (_edb Chart) AddScatterChart() ScatterChart {
	_cdg := _g.NewCT_PlotAreaChoice()
	_edb._aga.Chart.PlotArea.Choice = append(_edb._aga.Chart.PlotArea.Choice, _cdg)
	_cdg.ScatterChart = _g.NewCT_ScatterChart()
	_faa := ScatterChart{_cfb: _cdg.ScatterChart}
	_faa.InitializeDefaults()
	return _faa
}
func (_fggae SurfaceChartSeries) Values() NumberDataSource {
	if _fggae._dee.Val == nil {
		_fggae._dee.Val = _g.NewCT_NumDataSource()
	}
	_gad := MakeNumberDataSource(_fggae._dee.Val)
	_gad.CreateEmptyNumberCache()
	return _gad
}

// AddLegend adds a legend to a chart, replacing any existing legend.
func (_ffg Chart) AddLegend() Legend {
	_ffg._aga.Chart.Legend = _g.NewCT_Legend()
	_agf := MakeLegend(_ffg._aga.Chart.Legend)
	_agf.InitializeDefaults()
	return _agf
}

// Marker returns the marker properties.
func (_bfgg LineChartSeries) Marker() Marker {
	if _bfgg._feg.Marker == nil {
		_bfgg._feg.Marker = _g.NewCT_Marker()
	}
	return MakeMarker(_bfgg._feg.Marker)
}
func (_aee DateAxis) SetMajorTickMark(m _g.ST_TickMark) {
	if m == _g.ST_TickMarkUnset {
		_aee._bfe.MajorTickMark = nil
	} else {
		_aee._bfe.MajorTickMark = _g.NewCT_TickMark()
		_aee._bfe.MajorTickMark.ValAttr = m
	}
}

// X returns the inner wrapped XML type.
func (_dbaa BubbleChart) X() *_g.CT_BubbleChart { return _dbaa._da }

// CategoryAxis returns the category data source.
func (_fe BarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if _fe._fc.Cat == nil {
		_fe._fc.Cat = _g.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(_fe._fc.Cat)
}

// X returns the inner wrapped XML type.
func (_dcca ScatterChartSeries) X() *_g.CT_ScatterSer { return _dcca._bebg }

// SetText sets the series text
func (_fbe SurfaceChartSeries) SetText(s string) {
	_fbe._dee.Tx = _g.NewCT_SerTx()
	_fbe._dee.Tx.Choice.V = &s
}
func MakeDataLabels(x *_g.CT_DLbls) DataLabels { return DataLabels{x} }

type SeriesAxis struct{ _dgdc *_g.CT_SerAx }

// X returns the inner wrapped XML type.
func (_cfab PieChartSeries) X() *_g.CT_PieSer { return _cfab._eff }

// RadarChartSeries is a series to be used on an Radar chart.
type RadarChartSeries struct{ _fbg *_g.CT_RadarSer }

// AddSeries adds a default series to a Stock chart.
func (_badd StockChart) AddSeries() LineChartSeries {
	_fgbc := _g.NewCT_LineSer()
	_badd._bgfg.Ser = append(_badd._bgfg.Ser, _fgbc)
	_fgbc.Idx.ValAttr = uint32(len(_badd._bgfg.Ser) - 1)
	_fgbc.Order.ValAttr = uint32(len(_badd._bgfg.Ser) - 1)
	_dgca := LineChartSeries{_fgbc}
	_dgca.Values().CreateEmptyNumberCache()
	_dgca.Properties().LineProperties().SetNoFill()
	return _dgca
}

// Chart is a generic chart.
type Chart struct{ _aga *_g.ChartSpace }
