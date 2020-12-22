package math

import (
	_g "encoding/xml"
	_f "fmt"
	_e "strconv"

	_gg "gitee.com/gooffice/gooffice"
	_be "gitee.com/gooffice/gooffice/schema/soo/ofc/sharedTypes"
)

// ValidateWithPath validates the CT_MCS and its children, prefixing error messages with path
func (_gdca *CT_MCS) ValidateWithPath(path string) error {
	for _aaf, _cgab := range _gdca.Mc {
		if _dccb := _cgab.ValidateWithPath(_f.Sprintf("\u0025s\u002f\u004d\u0063\u005b\u0025\u0064]", path, _aaf)); _dccb != nil {
			return _dccb
		}
	}
	return nil
}

// ValidateWithPath validates the CT_EqArrPr and its children, prefixing error messages with path
func (_acc *CT_EqArrPr) ValidateWithPath(path string) error {
	if _acc.BaseJc != nil {
		if _edcff := _acc.BaseJc.ValidateWithPath(path + "\u002fB\u0061\u0073\u0065\u004a\u0063"); _edcff != nil {
			return _edcff
		}
	}
	if _acc.MaxDist != nil {
		if _eea := _acc.MaxDist.ValidateWithPath(path + "\u002f\u004d\u0061\u0078\u0044\u0069\u0073\u0074"); _eea != nil {
			return _eea
		}
	}
	if _acc.ObjDist != nil {
		if _abcg := _acc.ObjDist.ValidateWithPath(path + "\u002f\u004f\u0062\u006a\u0044\u0069\u0073\u0074"); _abcg != nil {
			return _abcg
		}
	}
	if _acc.RSpRule != nil {
		if _ebad := _acc.RSpRule.ValidateWithPath(path + "\u002f\u0052\u0053\u0070\u0052\u0075\u006c\u0065"); _ebad != nil {
			return _ebad
		}
	}
	if _acc.RSp != nil {
		if _bef := _acc.RSp.ValidateWithPath(path + "\u002f\u0052\u0053\u0070"); _bef != nil {
			return _bef
		}
	}
	if _acc.CtrlPr != nil {
		if _agd := _acc.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _agd != nil {
			return _agd
		}
	}
	return nil
}
func (_bcff *CT_RadPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gfcf:
	for {
		_gbea, _acab := d.Token()
		if _acab != nil {
			return _acab
		}
		switch _bbedf := _gbea.(type) {
		case _g.StartElement:
			switch _bbedf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064e\u0067\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064e\u0067\u0048\u0069\u0064\u0065"}:
				_bcff.DegHide = NewCT_OnOff()
				if _dcgf := d.DecodeElement(_bcff.DegHide, &_bbedf); _dcgf != nil {
					return _dcgf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_bcff.CtrlPr = NewCT_CtrlPr()
				if _bfef := d.DecodeElement(_bcff.CtrlPr, &_bbedf); _bfef != nil {
					return _bfef
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0061\u0064\u0050\u0072\u0020\u0025\u0076", _bbedf.Name)
				if _dcce := d.Skip(); _dcce != nil {
					return _dcce
				}
			}
		case _g.EndElement:
			break _gfcf
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OMathJc and its children
func (_gafbb *CT_OMathJc) Validate() error {
	return _gafbb.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u004a\u0063")
}

// ValidateWithPath validates the CT_MathPr and its children, prefixing error messages with path
func (_ffgga *CT_MathPr) ValidateWithPath(path string) error {
	if _ffgga.MathFont != nil {
		if _dfbb := _ffgga.MathFont.ValidateWithPath(path + "\u002fM\u0061\u0074\u0068\u0046\u006f\u006et"); _dfbb != nil {
			return _dfbb
		}
	}
	if _ffgga.BrkBin != nil {
		if _cgec := _ffgga.BrkBin.ValidateWithPath(path + "\u002fB\u0072\u006b\u0042\u0069\u006e"); _cgec != nil {
			return _cgec
		}
	}
	if _ffgga.BrkBinSub != nil {
		if _fffb := _ffgga.BrkBinSub.ValidateWithPath(path + "\u002f\u0042\u0072\u006b\u0042\u0069\u006e\u0053\u0075\u0062"); _fffb != nil {
			return _fffb
		}
	}
	if _ffgga.SmallFrac != nil {
		if _edgab := _ffgga.SmallFrac.ValidateWithPath(path + "\u002f\u0053\u006d\u0061\u006c\u006c\u0046\u0072\u0061\u0063"); _edgab != nil {
			return _edgab
		}
	}
	if _ffgga.DispDef != nil {
		if _agba := _ffgga.DispDef.ValidateWithPath(path + "\u002f\u0044\u0069\u0073\u0070\u0044\u0065\u0066"); _agba != nil {
			return _agba
		}
	}
	if _ffgga.LMargin != nil {
		if _befb := _ffgga.LMargin.ValidateWithPath(path + "\u002f\u004c\u004d\u0061\u0072\u0067\u0069\u006e"); _befb != nil {
			return _befb
		}
	}
	if _ffgga.RMargin != nil {
		if _gbdc := _ffgga.RMargin.ValidateWithPath(path + "\u002f\u0052\u004d\u0061\u0072\u0067\u0069\u006e"); _gbdc != nil {
			return _gbdc
		}
	}
	if _ffgga.DefJc != nil {
		if _bgcf := _ffgga.DefJc.ValidateWithPath(path + "\u002f\u0044\u0065\u0066\u004a\u0063"); _bgcf != nil {
			return _bgcf
		}
	}
	if _ffgga.PreSp != nil {
		if _ggba := _ffgga.PreSp.ValidateWithPath(path + "\u002f\u0050\u0072\u0065\u0053\u0070"); _ggba != nil {
			return _ggba
		}
	}
	if _ffgga.PostSp != nil {
		if _dabd := _ffgga.PostSp.ValidateWithPath(path + "\u002fP\u006f\u0073\u0074\u0053\u0070"); _dabd != nil {
			return _dabd
		}
	}
	if _ffgga.InterSp != nil {
		if _cdeb := _ffgga.InterSp.ValidateWithPath(path + "\u002f\u0049\u006e\u0074\u0065\u0072\u0053\u0070"); _cdeb != nil {
			return _cdeb
		}
	}
	if _ffgga.IntraSp != nil {
		if _dcfcb := _ffgga.IntraSp.ValidateWithPath(path + "\u002f\u0049\u006e\u0074\u0072\u0061\u0053\u0070"); _dcfcb != nil {
			return _dcfcb
		}
	}
	if _ffgga.Choice != nil {
		if _gec := _ffgga.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _gec != nil {
			return _gec
		}
	}
	if _ffgga.IntLim != nil {
		if _ebga := _ffgga.IntLim.ValidateWithPath(path + "\u002fI\u006e\u0074\u004c\u0069\u006d"); _ebga != nil {
			return _ebga
		}
	}
	if _ffgga.NaryLim != nil {
		if _bbde := _ffgga.NaryLim.ValidateWithPath(path + "\u002f\u004e\u0061\u0072\u0079\u004c\u0069\u006d"); _bbde != nil {
			return _bbde
		}
	}
	return nil
}

type MathPr struct{ CT_MathPr }

func (_edcdd ST_LimLoc) String() string {
	switch _edcdd {
	case 0:
		return ""
	case 1:
		return "\u0075\u006e\u0064\u004f\u0076\u0072"
	case 2:
		return "\u0073\u0075\u0062\u0053\u0075\u0070"
	}
	return ""
}

// ValidateWithPath validates the CT_XAlign and its children, prefixing error messages with path
func (_ceaa *CT_XAlign) ValidateWithPath(path string) error {
	if _ceaa.ValAttr == _be.ST_XAlignUnset {
		return _f.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _dgdcf := _ceaa.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _dgdcf != nil {
		return _dgdcf
	}
	return nil
}

// ValidateWithPath validates the CT_MPr and its children, prefixing error messages with path
func (_edega *CT_MPr) ValidateWithPath(path string) error {
	if _edega.BaseJc != nil {
		if _gbef := _edega.BaseJc.ValidateWithPath(path + "\u002fB\u0061\u0073\u0065\u004a\u0063"); _gbef != nil {
			return _gbef
		}
	}
	if _edega.PlcHide != nil {
		if _ecga := _edega.PlcHide.ValidateWithPath(path + "\u002f\u0050\u006c\u0063\u0048\u0069\u0064\u0065"); _ecga != nil {
			return _ecga
		}
	}
	if _edega.RSpRule != nil {
		if _ebeg := _edega.RSpRule.ValidateWithPath(path + "\u002f\u0052\u0053\u0070\u0052\u0075\u006c\u0065"); _ebeg != nil {
			return _ebeg
		}
	}
	if _edega.CGpRule != nil {
		if _dfff := _edega.CGpRule.ValidateWithPath(path + "\u002f\u0043\u0047\u0070\u0052\u0075\u006c\u0065"); _dfff != nil {
			return _dfff
		}
	}
	if _edega.RSp != nil {
		if _ddaa := _edega.RSp.ValidateWithPath(path + "\u002f\u0052\u0053\u0070"); _ddaa != nil {
			return _ddaa
		}
	}
	if _edega.CSp != nil {
		if _bgd := _edega.CSp.ValidateWithPath(path + "\u002f\u0043\u0053\u0070"); _bgd != nil {
			return _bgd
		}
	}
	if _edega.CGp != nil {
		if _badb := _edega.CGp.ValidateWithPath(path + "\u002f\u0043\u0047\u0070"); _badb != nil {
			return _badb
		}
	}
	if _edega.Mcs != nil {
		if _cgd := _edega.Mcs.ValidateWithPath(path + "\u002f\u004d\u0063\u0073"); _cgd != nil {
			return _cgd
		}
	}
	if _edega.CtrlPr != nil {
		if _egea := _edega.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _egea != nil {
			return _egea
		}
	}
	return nil
}

// Validate validates the CT_BreakBin and its children
func (_bdd *CT_BreakBin) Validate() error {
	return _bdd.ValidateWithPath("C\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042\u0069\u006e")
}

// Validate validates the CT_Integer2 and its children
func (_afg *CT_Integer2) Validate() error {
	return _afg.ValidateWithPath("C\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032")
}
func NewCT_BreakBinSub() *CT_BreakBinSub { _bde := &CT_BreakBinSub{}; return _bde }

// Validate validates the CT_XAlign and its children
func (_cegb *CT_XAlign) Validate() error {
	return _cegb.ValidateWithPath("\u0043T\u005f\u0058\u0041\u006c\u0069\u0067n")
}

type CT_MPr struct {
	BaseJc  *CT_YAlign
	PlcHide *CT_OnOff
	RSpRule *CT_SpacingRule
	CGpRule *CT_SpacingRule
	RSp     *CT_UnSignedInteger
	CSp     *CT_UnSignedInteger
	CGp     *CT_UnSignedInteger
	Mcs     *CT_MCS
	CtrlPr  *CT_CtrlPr
}

// Validate validates the CT_F and its children
func (_dcg *CT_F) Validate() error { return _dcg.ValidateWithPath("\u0043\u0054\u005f\u0046") }
func (_ecfe *CT_Shp) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_agee, _bccd := _ecfe.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _bccd != nil {
		return _bccd
	}
	start.Attr = append(start.Attr, _agee)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_SPrePr struct{ CtrlPr *CT_CtrlPr }

func (_ace *CT_LimLowPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_egae:
	for {
		_effd, _dcfa := d.Token()
		if _dcfa != nil {
			return _dcfa
		}
		switch _efbf := _effd.(type) {
		case _g.StartElement:
			switch _efbf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ace.CtrlPr = NewCT_CtrlPr()
				if _gce := d.DecodeElement(_ace.CtrlPr, &_efbf); _gce != nil {
					return _gce
				}
			default:
				_gg.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004c\u0069\u006d\u004c\u006f\u0077\u0050\u0072\u0020\u0025\u0076", _efbf.Name)
				if _fca := d.Skip(); _fca != nil {
					return _fca
				}
			}
		case _g.EndElement:
			break _egae
		case _g.CharData:
		}
	}
	return nil
}
func (_cdca *CT_Integer255) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", _cdca.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_dbafc ST_BreakBin) String() string {
	switch _dbafc {
	case 0:
		return ""
	case 1:
		return "\u0062\u0065\u0066\u006f\u0072\u0065"
	case 2:
		return "\u0061\u0066\u0074e\u0072"
	case 3:
		return "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return ""
}

type CT_Acc struct {
	AccPr *CT_AccPr
	E     *CT_OMathArg
}

func (_ecfae ST_TopBot) ValidateWithPath(path string) error {
	switch _ecfae {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ecfae))
	}
	return nil
}
func NewCT_R() *CT_R { _cbd := &CT_R{}; return _cbd }
func (_fggg *CT_M) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fggg.MPr != nil {
		_bcg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006dP\u0072"}}
		e.EncodeElement(_fggg.MPr, _bcg)
	}
	_fcfde := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0072"}}
	for _, _cgcc := range _fggg.Mr {
		e.EncodeElement(_cgcc, _fcfde)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_cga *CT_FPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cga.Type != nil {
		_bff := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0074\u0079\u0070\u0065"}}
		e.EncodeElement(_cga.Type, _bff)
	}
	if _cga.CtrlPr != nil {
		_fbg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_cga.CtrlPr, _fbg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_OMathArg and its children
func (_aacd *CT_OMathArg) Validate() error {
	return _aacd.ValidateWithPath("C\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067")
}
func (_fbfed *CT_SpacingRule) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fbfed.ValAttr = 0
	for _, _adfg := range start.Attr {
		if _adfg.Name.Local == "\u0076\u0061\u006c" {
			_ebea, _ffbc := _e.ParseInt(_adfg.Value, 10, 64)
			if _ffbc != nil {
				return _ffbc
			}
			_fbfed.ValAttr = _ebea
			continue
		}
	}
	for {
		_ceff, _gggg := d.Token()
		if _gggg != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fS\u0070\u0061\u0063\u0069\u006e\u0067\u0052\u0075\u006c\u0065:\u0020\u0025\u0073", _gggg)
		}
		if _eebb, _gbec := _ceff.(_g.EndElement); _gbec && _eebb.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_CtrlPr struct{}

func (_bged *CT_TwipsMeasure) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", _bged.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_LimLowPr() *CT_LimLowPr { _beaa := &CT_LimLowPr{}; return _beaa }
func (_ccbc *CT_XAlign) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ccbc.ValAttr = _be.ST_XAlign(1)
	for _, _bgfa := range start.Attr {
		if _bgfa.Name.Local == "\u0076\u0061\u006c" {
			_ccbc.ValAttr.UnmarshalXMLAttr(_bgfa)
			continue
		}
	}
	for {
		_dcfcg, _bcaa := d.Token()
		if _bcaa != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0058\u0041\u006ci\u0067\u006e\u003a\u0020\u0025\u0073", _bcaa)
		}
		if _cfbd, _edce := _dcfcg.(_g.EndElement); _edce && _cfbd.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_RChoice and its children
func (_gfaa *CT_RChoice) Validate() error {
	return _gfaa.ValidateWithPath("\u0043\u0054\u005f\u0052\u0043\u0068\u006f\u0069\u0063\u0065")
}
func (_ceeg *CT_SSub) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ceeg.E = NewCT_OMathArg()
	_ceeg.Sub = NewCT_OMathArg()
_fdbc:
	for {
		_aeaga, _gbdd := d.Token()
		if _gbdd != nil {
			return _gbdd
		}
		switch _egfe := _aeaga.(type) {
		case _g.StartElement:
			switch _egfe.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062\u0050\u0072"}:
				_ceeg.SSubPr = NewCT_SSubPr()
				if _gaccd := d.DecodeElement(_ceeg.SSubPr, &_egfe); _gaccd != nil {
					return _gaccd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _afba := d.DecodeElement(_ceeg.E, &_egfe); _afba != nil {
					return _afba
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _efed := d.DecodeElement(_ceeg.Sub, &_egfe); _efed != nil {
					return _efed
				}
			default:
				_gg.Log("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0053\u0053\u0075\u0062\u0020\u0025\u0076", _egfe.Name)
				if _cbaa := d.Skip(); _cbaa != nil {
					return _cbaa
				}
			}
		case _g.EndElement:
			break _fdbc
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Script and its children
func (_efbb *CT_Script) Validate() error {
	return _efbb.ValidateWithPath("\u0043T\u005f\u0053\u0063\u0072\u0069\u0070t")
}
func (_eggf *CT_EqArr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gdge:
	for {
		_ffdf, _dae := d.Token()
		if _dae != nil {
			return _dae
		}
		switch _adaf := _ffdf.(type) {
		case _g.StartElement:
			switch _adaf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065q\u0041\u0072\u0072\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065q\u0041\u0072\u0072\u0050\u0072"}:
				_eggf.EqArrPr = NewCT_EqArrPr()
				if _acf := d.DecodeElement(_eggf.EqArrPr, &_adaf); _acf != nil {
					return _acf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				_dfb := NewCT_OMathArg()
				if _bce := d.DecodeElement(_dfb, &_adaf); _bce != nil {
					return _bce
				}
				_eggf.E = append(_eggf.E, _dfb)
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072\u0020\u0025\u0076", _adaf.Name)
				if _egec := d.Skip(); _egec != nil {
					return _egec
				}
			}
		case _g.EndElement:
			break _gdge
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BarPr and its children
func (_bec *CT_BarPr) Validate() error {
	return _bec.ValidateWithPath("\u0043\u0054\u005f\u0042\u0061\u0072\u0050\u0072")
}

// ValidateWithPath validates the CT_BarPr and its children, prefixing error messages with path
func (_de *CT_BarPr) ValidateWithPath(path string) error {
	if _de.Pos != nil {
		if _cec := _de.Pos.ValidateWithPath(path + "\u002f\u0050\u006f\u0073"); _cec != nil {
			return _cec
		}
	}
	if _de.CtrlPr != nil {
		if _dab := _de.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _dab != nil {
			return _dab
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Integer2 and its children, prefixing error messages with path
func (_faa *CT_Integer2) ValidateWithPath(path string) error {
	if _faa.ValAttr < -2 {
		return _f.Errorf("\u0025\u0073/m\u002e\u0056\u0061l\u0041\u0074\u0074\u0072 mu\u0073t \u0062\u0065\u0020\u003e\u003d\u0020\u002d2 \u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _faa.ValAttr)
	}
	if _faa.ValAttr > 2 {
		return _f.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0032\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _faa.ValAttr)
	}
	return nil
}
func (_efba *CT_MR) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	_eggd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	for _, _efa := range _efba.E {
		e.EncodeElement(_efa, _eggd)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_R struct {
	RPr    *CT_RPR
	Choice []*CT_RChoice
}
type CT_BarPr struct {
	Pos    *CT_TopBot
	CtrlPr *CT_CtrlPr
}

func (_beb *CT_Box) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_beb.E = NewCT_OMathArg()
_cfcf:
	for {
		_ddc, _ccg := d.Token()
		if _ccg != nil {
			return _ccg
		}
		switch _daa := _ddc.(type) {
		case _g.StartElement:
			switch _daa.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078P\u0072"}:
				_beb.BoxPr = NewCT_BoxPr()
				if _fe := d.DecodeElement(_beb.BoxPr, &_daa); _fe != nil {
					return _fe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _fce := d.DecodeElement(_beb.E, &_daa); _fce != nil {
					return _fce
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u006f\u0078\u0020\u0025\u0076", _daa.Name)
				if _bbg := d.Skip(); _bbg != nil {
					return _bbg
				}
			}
		case _g.EndElement:
			break _cfcf
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_D() *CT_D { _cbe := &CT_D{}; return _cbe }
func (_bac *CT_EqArr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _bac.EqArrPr != nil {
		_cbg := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0065\u0071\u0041\u0072\u0072\u0050r"}}
		e.EncodeElement(_bac.EqArrPr, _cbg)
	}
	_fgbf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	for _, _bbgf := range _bac.E {
		e.EncodeElement(_bbgf, _fgbf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Shp and its children
func (_bfdg *CT_Shp) Validate() error {
	return _bfdg.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0070")
}

type CT_OMath struct{ EG_OMathMathElements []*EG_OMathMathElements }

func (_ddea ST_LimLoc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_ddea.String(), start)
}

// ValidateWithPath validates the CT_SSubSup and its children, prefixing error messages with path
func (_bgbf *CT_SSubSup) ValidateWithPath(path string) error {
	if _bgbf.SSubSupPr != nil {
		if _bcffb := _bgbf.SSubSupPr.ValidateWithPath(path + "\u002f\u0053\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072"); _bcffb != nil {
			return _bcffb
		}
	}
	if _cbcgc := _bgbf.E.ValidateWithPath(path + "\u002f\u0045"); _cbcgc != nil {
		return _cbcgc
	}
	if _gdgd := _bgbf.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _gdgd != nil {
		return _gdgd
	}
	if _cedf := _bgbf.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _cedf != nil {
		return _cedf
	}
	return nil
}
func NewCT_CtrlPr() *CT_CtrlPr { _ead := &CT_CtrlPr{}; return _ead }

// ValidateWithPath validates the CT_Style and its children, prefixing error messages with path
func (_eefb *CT_Style) ValidateWithPath(path string) error {
	if _ddcde := _eefb.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _ddcde != nil {
		return _ddcde
	}
	return nil
}
func NewCT_Box() *CT_Box { _afa := &CT_Box{}; _afa.E = NewCT_OMathArg(); return _afa }

type ST_Script byte
type CT_FPr struct {
	Type   *CT_FType
	CtrlPr *CT_CtrlPr
}
type CT_F struct {
	FPr *CT_FPr
	Num *CT_OMathArg
	Den *CT_OMathArg
}
type CT_M struct {
	MPr *CT_MPr
	Mr  []*CT_MR
}

// Validate validates the EG_OMathElements and its children
func (_bgcfa *EG_OMathElements) Validate() error {
	return _bgcfa.ValidateWithPath("\u0045\u0047_\u004f\u004d\u0061t\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073")
}
func NewCT_FPr() *CT_FPr { _age := &CT_FPr{}; return _age }
func NewCT_MPr() *CT_MPr { _egaaa := &CT_MPr{}; return _egaaa }
func (_efcba *CT_UnSignedInteger) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", _efcba.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_LimUppPr and its children, prefixing error messages with path
func (_bbbee *CT_LimUppPr) ValidateWithPath(path string) error {
	if _bbbee.CtrlPr != nil {
		if _ecbg := _bbbee.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _ecbg != nil {
			return _ecbg
		}
	}
	return nil
}
func NewCT_SPre() *CT_SPre {
	_bbfc := &CT_SPre{}
	_bbfc.Sub = NewCT_OMathArg()
	_bbfc.Sup = NewCT_OMathArg()
	_bbfc.E = NewCT_OMathArg()
	return _bbfc
}

// ValidateWithPath validates the EG_OMathElements and its children, prefixing error messages with path
func (_eede *EG_OMathElements) ValidateWithPath(path string) error {
	for _dcfab, _gggfa := range _eede.EG_OMathMathElements {
		if _eacc := _gggfa.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0045\u0047\u005f\u004fM\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006ce\u006d\u0065\u006et\u0073[\u0025\u0064\u005d", path, _dcfab)); _eacc != nil {
			return _eacc
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Integer255 and its children, prefixing error messages with path
func (_fgbb *CT_Integer255) ValidateWithPath(path string) error {
	if _fgbb.ValAttr < 1 {
		return _f.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0031\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _fgbb.ValAttr)
	}
	if _fgbb.ValAttr > 255 {
		return _f.Errorf("\u0025\u0073/\u006d\u002e\u0056\u0061l\u0041\u0074t\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u0062e\u0020\u003c\u003d\u0020\u0032\u0035\u0035\u0020\u0028\u0068\u0061\u0076e\u0020\u0025\u0076\u0029", path, _fgbb.ValAttr)
	}
	return nil
}
func (_adg *CT_MCS) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	_gbfe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0063"}}
	for _, _ebag := range _adg.Mc {
		e.EncodeElement(_ebag, _gbfe)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_TopBot() *CT_TopBot { _cgee := &CT_TopBot{}; _cgee.ValAttr = ST_TopBot(1); return _cgee }

// ValidateWithPath validates the CT_Box and its children, prefixing error messages with path
func (_bfee *CT_Box) ValidateWithPath(path string) error {
	if _bfee.BoxPr != nil {
		if _dce := _bfee.BoxPr.ValidateWithPath(path + "\u002f\u0042\u006f\u0078\u0050\u0072"); _dce != nil {
			return _dce
		}
	}
	if _fcd := _bfee.E.ValidateWithPath(path + "\u002f\u0045"); _fcd != nil {
		return _fcd
	}
	return nil
}
func (_ffc *CT_BorderBoxPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_bdf:
	for {
		_ddb, _ade := d.Token()
		if _ade != nil {
			return _ade
		}
		switch _abc := _ddb.(type) {
		case _g.StartElement:
			switch _abc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068i\u0064\u0065\u0054\u006f\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068i\u0064\u0065\u0054\u006f\u0070"}:
				_ffc.HideTop = NewCT_OnOff()
				if _dcc := d.DecodeElement(_ffc.HideTop, &_abc); _dcc != nil {
					return _dcc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068i\u0064\u0065\u0042\u006f\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068i\u0064\u0065\u0042\u006f\u0074"}:
				_ffc.HideBot = NewCT_OnOff()
				if _ee := d.DecodeElement(_ffc.HideBot, &_abc); _ee != nil {
					return _ee
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068\u0069\u0064\u0065\u004c\u0065\u0066\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068\u0069\u0064\u0065\u004c\u0065\u0066\u0074"}:
				_ffc.HideLeft = NewCT_OnOff()
				if _gcbd := d.DecodeElement(_ffc.HideLeft, &_abc); _gcbd != nil {
					return _gcbd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068i\u0064\u0065\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068i\u0064\u0065\u0052\u0069\u0067\u0068t"}:
				_ffc.HideRight = NewCT_OnOff()
				if _dcb := d.DecodeElement(_ffc.HideRight, &_abc); _dcb != nil {
					return _dcb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0048"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0048"}:
				_ffc.StrikeH = NewCT_OnOff()
				if _ffg := d.DecodeElement(_ffc.StrikeH, &_abc); _ffg != nil {
					return _ffg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0056"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0056"}:
				_ffc.StrikeV = NewCT_OnOff()
				if _aaea := d.DecodeElement(_ffc.StrikeV, &_abc); _aaea != nil {
					return _aaea
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"}:
				_ffc.StrikeBLTR = NewCT_OnOff()
				if _ag := d.DecodeElement(_ffc.StrikeBLTR, &_abc); _ag != nil {
					return _ag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"}:
				_ffc.StrikeTLBR = NewCT_OnOff()
				if _bg := d.DecodeElement(_ffc.StrikeTLBR, &_abc); _bg != nil {
					return _bg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ffc.CtrlPr = NewCT_CtrlPr()
				if _bfe := d.DecodeElement(_ffc.CtrlPr, &_abc); _bfe != nil {
					return _bfe
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u006f\u0072\u0064\u0065\u0072\u0042o\u0078P\u0072\u0020\u0025\u0076", _abc.Name)
				if _eda := d.Skip(); _eda != nil {
					return _eda
				}
			}
		case _g.EndElement:
			break _bdf
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_EqArr and its children, prefixing error messages with path
func (_edcf *CT_EqArr) ValidateWithPath(path string) error {
	if _edcf.EqArrPr != nil {
		if _bddd := _edcf.EqArrPr.ValidateWithPath(path + "\u002f\u0045\u0071\u0041\u0072\u0072\u0050\u0072"); _bddd != nil {
			return _bddd
		}
	}
	for _eaf, _cbb := range _edcf.E {
		if _gac := _cbb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0045\u005b\u0025\u0064\u005d", path, _eaf)); _gac != nil {
			return _gac
		}
	}
	return nil
}
func NewCT_OMathArgPr() *CT_OMathArgPr { _bbaae := &CT_OMathArgPr{}; return _bbaae }
func NewCT_EqArr() *CT_EqArr           { _dbe := &CT_EqArr{}; return _dbe }
func (_cada *CT_SSub) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cada.SSubPr != nil {
		_bfeef := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0062\u0050\u0072"}}
		e.EncodeElement(_cada.SSubPr, _bfeef)
	}
	_dace := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_cada.E, _dace)
	_fegdd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_cada.Sub, _fegdd)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_fceg ST_Jc) ValidateWithPath(path string) error {
	switch _fceg {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fceg))
	}
	return nil
}
func (_eeedg *ST_Style) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_eeedg = 0
	case "\u0070":
		*_eeedg = 1
	case "\u0062":
		*_eeedg = 2
	case "\u0069":
		*_eeedg = 3
	case "\u0062\u0069":
		*_eeedg = 4
	}
	return nil
}
func (_bdcf ST_FType) ValidateWithPath(path string) error {
	switch _bdcf {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bdcf))
	}
	return nil
}
func NewCT_TwipsMeasure() *CT_TwipsMeasure { _ebda := &CT_TwipsMeasure{}; return _ebda }

type CT_SpacingRule struct{ ValAttr int64 }
type EG_OMathElements struct{ EG_OMathMathElements []*EG_OMathMathElements }

// ValidateWithPath validates the CT_SSubPr and its children, prefixing error messages with path
func (_bfcd *CT_SSubPr) ValidateWithPath(path string) error {
	if _bfcd.CtrlPr != nil {
		if _ecdec := _bfcd.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _ecdec != nil {
			return _ecdec
		}
	}
	return nil
}

// ValidateWithPath validates the CT_BorderBoxPr and its children, prefixing error messages with path
func (_ffd *CT_BorderBoxPr) ValidateWithPath(path string) error {
	if _ffd.HideTop != nil {
		if _ef := _ffd.HideTop.ValidateWithPath(path + "\u002f\u0048\u0069\u0064\u0065\u0054\u006f\u0070"); _ef != nil {
			return _ef
		}
	}
	if _ffd.HideBot != nil {
		if _adee := _ffd.HideBot.ValidateWithPath(path + "\u002f\u0048\u0069\u0064\u0065\u0042\u006f\u0074"); _adee != nil {
			return _adee
		}
	}
	if _ffd.HideLeft != nil {
		if _dbd := _ffd.HideLeft.ValidateWithPath(path + "\u002fH\u0069\u0064\u0065\u004c\u0065\u0066t"); _dbd != nil {
			return _dbd
		}
	}
	if _ffd.HideRight != nil {
		if _cbc := _ffd.HideRight.ValidateWithPath(path + "\u002f\u0048\u0069\u0064\u0065\u0052\u0069\u0067\u0068\u0074"); _cbc != nil {
			return _cbc
		}
	}
	if _ffd.StrikeH != nil {
		if _eef := _ffd.StrikeH.ValidateWithPath(path + "\u002f\u0053\u0074\u0072\u0069\u006b\u0065\u0048"); _eef != nil {
			return _eef
		}
	}
	if _ffd.StrikeV != nil {
		if _gae := _ffd.StrikeV.ValidateWithPath(path + "\u002f\u0053\u0074\u0072\u0069\u006b\u0065\u0056"); _gae != nil {
			return _gae
		}
	}
	if _ffd.StrikeBLTR != nil {
		if _ccc := _ffd.StrikeBLTR.ValidateWithPath(path + "/\u0053\u0074\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"); _ccc != nil {
			return _ccc
		}
	}
	if _ffd.StrikeTLBR != nil {
		if _egb := _ffd.StrikeTLBR.ValidateWithPath(path + "/\u0053\u0074\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"); _egb != nil {
			return _egb
		}
	}
	if _ffd.CtrlPr != nil {
		if _ec := _ffd.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _ec != nil {
			return _ec
		}
	}
	return nil
}

type CT_Style struct{ ValAttr ST_Style }

func (_ffe *CT_DPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_aeb:
	for {
		_ddcd, _aga := d.Token()
		if _aga != nil {
			return _aga
		}
		switch _beg := _ddcd.(type) {
		case _g.StartElement:
			switch _beg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0065\u0067\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0065\u0067\u0043\u0068\u0072"}:
				_ffe.BegChr = NewCT_Char()
				if _gdf := d.DecodeElement(_ffe.BegChr, &_beg); _gdf != nil {
					return _gdf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0065\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0065\u0070\u0043\u0068\u0072"}:
				_ffe.SepChr = NewCT_Char()
				if _dfcg := d.DecodeElement(_ffe.SepChr, &_beg); _dfcg != nil {
					return _dfcg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u006e\u0064\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u006e\u0064\u0043\u0068\u0072"}:
				_ffe.EndChr = NewCT_Char()
				if _fgf := d.DecodeElement(_ffe.EndChr, &_beg); _fgf != nil {
					return _fgf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}:
				_ffe.Grow = NewCT_OnOff()
				if _gaa := d.DecodeElement(_ffe.Grow, &_beg); _gaa != nil {
					return _gaa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0068\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0068\u0070"}:
				_ffe.Shp = NewCT_Shp()
				if _bdef := d.DecodeElement(_ffe.Shp, &_beg); _bdef != nil {
					return _bdef
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ffe.CtrlPr = NewCT_CtrlPr()
				if _egg := d.DecodeElement(_ffe.CtrlPr, &_beg); _egg != nil {
					return _egg
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0044\u0050\u0072\u0020\u0025\u0076", _beg.Name)
				if _gdga := d.Skip(); _gdga != nil {
					return _gdga
				}
			}
		case _g.EndElement:
			break _aeb
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DPr and its children
func (_efd *CT_DPr) Validate() error {
	return _efd.ValidateWithPath("\u0043\u0054\u005f\u0044\u0050\u0072")
}
func (_da *CT_Bar) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _da.BarPr != nil {
		_cb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0062\u0061\u0072\u0050\u0072"}}
		e.EncodeElement(_da.BarPr, _cb)
	}
	_ad := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_da.E, _ad)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_MCS struct{ Mc []*CT_MC }

func (_bfefg *ST_LimLoc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_afdfd, _begb := d.Token()
	if _begb != nil {
		return _begb
	}
	if _affc, _cgddc := _afdfd.(_g.EndElement); _cgddc && _affc.Name == start.Name {
		*_bfefg = 1
		return nil
	}
	if _fffg, _fdfda := _afdfd.(_g.CharData); !_fdfda {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _afdfd)
	} else {
		switch string(_fffg) {
		case "":
			*_bfefg = 0
		case "\u0075\u006e\u0064\u004f\u0076\u0072":
			*_bfefg = 1
		case "\u0073\u0075\u0062\u0053\u0075\u0070":
			*_bfefg = 2
		}
	}
	_afdfd, _begb = d.Token()
	if _begb != nil {
		return _begb
	}
	if _bdfe, _bgdg := _afdfd.(_g.EndElement); _bgdg && _bdfe.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _afdfd)
}
func (_edg *CT_Box) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _edg.BoxPr != nil {
		_ffa := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0062\u006f\u0078\u0050\u0072"}}
		e.EncodeElement(_edg.BoxPr, _ffa)
	}
	_cfca := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_edg.E, _cfca)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_SPre and its children, prefixing error messages with path
func (_fffba *CT_SPre) ValidateWithPath(path string) error {
	if _fffba.SPrePr != nil {
		if _aced := _fffba.SPrePr.ValidateWithPath(path + "\u002fS\u0050\u0072\u0065\u0050\u0072"); _aced != nil {
			return _aced
		}
	}
	if _beea := _fffba.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _beea != nil {
		return _beea
	}
	if _gffg := _fffba.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _gffg != nil {
		return _gffg
	}
	if _gggdf := _fffba.E.ValidateWithPath(path + "\u002f\u0045"); _gggdf != nil {
		return _gggdf
	}
	return nil
}

type CT_String struct{ ValAttr *string }

// Validate validates the CT_SPrePr and its children
func (_caca *CT_SPrePr) Validate() error {
	return _caca.ValidateWithPath("\u0043T\u005f\u0053\u0050\u0072\u0065\u0050r")
}

// Validate validates the CT_Rad and its children
func (_cfbf *CT_Rad) Validate() error {
	return _cfbf.ValidateWithPath("\u0043\u0054\u005f\u0052\u0061\u0064")
}

type CT_GroupChr struct {
	GroupChrPr *CT_GroupChrPr
	E          *CT_OMathArg
}

func NewCT_LimUpp() *CT_LimUpp {
	_ffca := &CT_LimUpp{}
	_ffca.E = NewCT_OMathArg()
	_ffca.Lim = NewCT_OMathArg()
	return _ffca
}
func (_adafe *CT_SSupPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_efcb:
	for {
		_acffc, _aad := d.Token()
		if _aad != nil {
			return _aad
		}
		switch _dgda := _acffc.(type) {
		case _g.StartElement:
			switch _dgda.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_adafe.CtrlPr = NewCT_CtrlPr()
				if _cedeb := d.DecodeElement(_adafe.CtrlPr, &_dgda); _cedeb != nil {
					return _cedeb
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0053\u0053\u0075p\u0050\u0072 \u0025\u0076", _dgda.Name)
				if _acbc := d.Skip(); _acbc != nil {
					return _acbc
				}
			}
		case _g.EndElement:
			break _efcb
		case _g.CharData:
		}
	}
	return nil
}
func (_gdfa *CT_RChoice) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _gdfa.T != nil {
		_agea := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0074"}}
		for _, _feab := range _gdfa.T {
			e.EncodeElement(_feab, _agea)
		}
	}
	return nil
}

// ValidateWithPath validates the CT_SSupPr and its children, prefixing error messages with path
func (_acdd *CT_SSupPr) ValidateWithPath(path string) error {
	if _acdd.CtrlPr != nil {
		if _ecdf := _acdd.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _ecdf != nil {
			return _ecdf
		}
	}
	return nil
}

// ValidateWithPath validates the CT_NaryPr and its children, prefixing error messages with path
func (_ageb *CT_NaryPr) ValidateWithPath(path string) error {
	if _ageb.Chr != nil {
		if _gfg := _ageb.Chr.ValidateWithPath(path + "\u002f\u0043\u0068\u0072"); _gfg != nil {
			return _gfg
		}
	}
	if _ageb.LimLoc != nil {
		if _eebd := _ageb.LimLoc.ValidateWithPath(path + "\u002fL\u0069\u006d\u004c\u006f\u0063"); _eebd != nil {
			return _eebd
		}
	}
	if _ageb.Grow != nil {
		if _aeg := _ageb.Grow.ValidateWithPath(path + "\u002f\u0047\u0072o\u0077"); _aeg != nil {
			return _aeg
		}
	}
	if _ageb.SubHide != nil {
		if _egcd := _ageb.SubHide.ValidateWithPath(path + "\u002f\u0053\u0075\u0062\u0048\u0069\u0064\u0065"); _egcd != nil {
			return _egcd
		}
	}
	if _ageb.SupHide != nil {
		if _abgc := _ageb.SupHide.ValidateWithPath(path + "\u002f\u0053\u0075\u0070\u0048\u0069\u0064\u0065"); _abgc != nil {
			return _abgc
		}
	}
	if _ageb.CtrlPr != nil {
		if _efdf := _ageb.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _efdf != nil {
			return _efdf
		}
	}
	return nil
}

// Validate validates the CT_BorderBoxPr and its children
func (_bcb *CT_BorderBoxPr) Validate() error {
	return _bcb.ValidateWithPath("\u0043\u0054\u005f\u0042\u006f\u0072\u0064\u0065\u0072B\u006f\u0078\u0050\u0072")
}
func (_fgfd *CT_FPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_acdb:
	for {
		_aeaf, _fdaa := d.Token()
		if _fdaa != nil {
			return _fdaa
		}
		switch _dca := _aeaf.(type) {
		case _g.StartElement:
			switch _dca.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074\u0079\u0070\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074\u0079\u0070\u0065"}:
				_fgfd.Type = NewCT_FType()
				if _bbbe := d.DecodeElement(_fgfd.Type, &_dca); _bbbe != nil {
					return _bbbe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fgfd.CtrlPr = NewCT_CtrlPr()
				if _dad := d.DecodeElement(_fgfd.CtrlPr, &_dca); _dad != nil {
					return _dad
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0046\u0050\u0072\u0020\u0025\u0076", _dca.Name)
				if _abce := d.Skip(); _abce != nil {
					return _abce
				}
			}
		case _g.EndElement:
			break _acdb
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LimUpp and its children
func (_adag *CT_LimUpp) Validate() error {
	return _adag.ValidateWithPath("\u0043T\u005f\u004c\u0069\u006d\u0055\u0070p")
}
func (_bea *CT_Integer2) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", _bea.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_dgeg *CT_Nary) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dgeg.Sub = NewCT_OMathArg()
	_dgeg.Sup = NewCT_OMathArg()
	_dgeg.E = NewCT_OMathArg()
_cgef:
	for {
		_eee, _ebgc := d.Token()
		if _ebgc != nil {
			return _ebgc
		}
		switch _eeff := _eee.(type) {
		case _g.StartElement:
			switch _eeff.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079\u0050\u0072"}:
				_dgeg.NaryPr = NewCT_NaryPr()
				if _bcf := d.DecodeElement(_dgeg.NaryPr, &_eeff); _bcf != nil {
					return _bcf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _fbdc := d.DecodeElement(_dgeg.Sub, &_eeff); _fbdc != nil {
					return _fbdc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _gada := d.DecodeElement(_dgeg.Sup, &_eeff); _gada != nil {
					return _gada
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _eddd := d.DecodeElement(_dgeg.E, &_eeff); _eddd != nil {
					return _eddd
				}
			default:
				_gg.Log("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u004e\u0061\u0072\u0079\u0020\u0025\u0076", _eeff.Name)
				if _dafc := d.Skip(); _dafc != nil {
					return _dafc
				}
			}
		case _g.EndElement:
			break _cgef
		case _g.CharData:
		}
	}
	return nil
}
func (_gafg *CT_R) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gafg.RPr != nil {
		_bcae := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072P\u0072"}}
		e.EncodeElement(_gafg.RPr, _bcae)
	}
	if _gafg.Choice != nil {
		for _, _bbbef := range _gafg.Choice {
			_bbbef.MarshalXML(e, _g.StartElement{})
		}
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_Script() *CT_Script { _dbbf := &CT_Script{}; return _dbbf }

// ValidateWithPath validates the CT_MCPr and its children, prefixing error messages with path
func (_ffcec *CT_MCPr) ValidateWithPath(path string) error {
	if _ffcec.Count != nil {
		if _feed := _ffcec.Count.ValidateWithPath(path + "\u002f\u0043\u006f\u0075\u006e\u0074"); _feed != nil {
			return _feed
		}
	}
	if _ffcec.McJc != nil {
		if _cbef := _ffcec.McJc.ValidateWithPath(path + "\u002f\u004d\u0063J\u0063"); _cbef != nil {
			return _cbef
		}
	}
	return nil
}
func (_abbg *CT_LimUppPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _abbg.CtrlPr != nil {
		_caeg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_abbg.CtrlPr, _caeg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_aeab ST_BreakBinSub) Validate() error { return _aeab.ValidateWithPath("") }

// Validate validates the CT_M and its children
func (_badd *CT_M) Validate() error { return _badd.ValidateWithPath("\u0043\u0054\u005f\u004d") }

// ValidateWithPath validates the CT_BorderBox and its children, prefixing error messages with path
func (_dbc *CT_BorderBox) ValidateWithPath(path string) error {
	if _dbc.BorderBoxPr != nil {
		if _ged := _dbc.BorderBoxPr.ValidateWithPath(path + "\u002f\u0042\u006fr\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"); _ged != nil {
			return _ged
		}
	}
	if _ega := _dbc.E.ValidateWithPath(path + "\u002f\u0045"); _ega != nil {
		return _ega
	}
	return nil
}

// Validate validates the CT_EqArr and its children
func (_dcec *CT_EqArr) Validate() error {
	return _dcec.ValidateWithPath("\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072")
}

type CT_MR struct{ E []*CT_OMathArg }

func (_cccg *CT_OMathArgPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cccg.ArgSz != nil {
		_ddfbf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0061\u0072\u0067\u0053\u007a"}}
		e.EncodeElement(_cccg.ArgSz, _ddfbf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_OMathPara and its children, prefixing error messages with path
func (_bbdb *CT_OMathPara) ValidateWithPath(path string) error {
	if _bbdb.OMathParaPr != nil {
		if _edgae := _bbdb.OMathParaPr.ValidateWithPath(path + "\u002f\u004f\u004da\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"); _edgae != nil {
			return _edgae
		}
	}
	for _gdad, _cdga := range _bbdb.OMath {
		if _eafba := _cdga.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fO\u004d\u0061\u0074\u0068\u005b\u0025\u0064\u005d", path, _gdad)); _eafba != nil {
			return _eafba
		}
	}
	return nil
}

// Validate validates the CT_EqArrPr and its children
func (_egf *CT_EqArrPr) Validate() error {
	return _egf.ValidateWithPath("\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072\u0050\u0072")
}
func (_bcec *CT_Style) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _dcfae := range start.Attr {
		if _dcfae.Name.Local == "\u0076\u0061\u006c" {
			_bcec.ValAttr.UnmarshalXMLAttr(_dcfae)
			continue
		}
	}
	for {
		_dedf, _dbdb := d.Token()
		if _dbdb != nil {
			return _f.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fS\u0074\u0079\u006c\u0065: \u0025\u0073", _dbdb)
		}
		if _feeb, _aadb := _dedf.(_g.EndElement); _aadb && _feeb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_edbc *CT_RadPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _edbc.DegHide != nil {
		_dfee := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0064\u0065\u0067\u0048\u0069\u0064e"}}
		e.EncodeElement(_edbc.DegHide, _dfee)
	}
	if _edbc.CtrlPr != nil {
		_caad := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_edbc.CtrlPr, _caad)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the OMath and its children, prefixing error messages with path
func (_fcaca *OMath) ValidateWithPath(path string) error {
	if _gbca := _fcaca.CT_OMath.ValidateWithPath(path); _gbca != nil {
		return _gbca
	}
	return nil
}
func (_fgb *CT_CtrlPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_fedf *CT_OMathPara) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_agbb:
	for {
		_eagdf, _cecc := d.Token()
		if _cecc != nil {
			return _cecc
		}
		switch _ffabf := _eagdf.(type) {
		case _g.StartElement:
			switch _ffabf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}:
				_fedf.OMathParaPr = NewCT_OMathParaPr()
				if _dac := d.DecodeElement(_fedf.OMathParaPr, &_ffabf); _dac != nil {
					return _dac
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}:
				_bbed := NewCT_OMath()
				if _cbge := d.DecodeElement(_bbed, &_ffabf); _cbge != nil {
					return _cbge
				}
				_fedf.OMath = append(_fedf.OMath, _bbed)
			default:
				_gg.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_O\u004d\u0061t\u0068\u0050\u0061\u0072\u0061\u0020\u0025\u0076", _ffabf.Name)
				if _ffgd := d.Skip(); _ffgd != nil {
					return _ffgd
				}
			}
		case _g.EndElement:
			break _agbb
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_GroupChrPr and its children, prefixing error messages with path
func (_gaca *CT_GroupChrPr) ValidateWithPath(path string) error {
	if _gaca.Chr != nil {
		if _gcfc := _gaca.Chr.ValidateWithPath(path + "\u002f\u0043\u0068\u0072"); _gcfc != nil {
			return _gcfc
		}
	}
	if _gaca.Pos != nil {
		if _abcb := _gaca.Pos.ValidateWithPath(path + "\u002f\u0050\u006f\u0073"); _abcb != nil {
			return _abcb
		}
	}
	if _gaca.VertJc != nil {
		if _gfa := _gaca.VertJc.ValidateWithPath(path + "\u002fV\u0065\u0072\u0074\u004a\u0063"); _gfa != nil {
			return _gfa
		}
	}
	if _gaca.CtrlPr != nil {
		if _dfe := _gaca.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _dfe != nil {
			return _dfe
		}
	}
	return nil
}

type ST_Jc byte

func (_agbc *CT_Rad) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _agbc.RadPr != nil {
		_ecaa := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u0061\u0064\u0050\u0072"}}
		e.EncodeElement(_agbc.RadPr, _ecaa)
	}
	_cede := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064e\u0067"}}
	e.EncodeElement(_agbc.Deg, _cede)
	_bddc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_agbc.E, _bddc)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_beeae ST_FType) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_bebf := _g.Attr{}
	_bebf.Name = name
	switch _beeae {
	case ST_FTypeUnset:
		_bebf.Value = ""
	case ST_FTypeBar:
		_bebf.Value = "\u0062\u0061\u0072"
	case ST_FTypeSkw:
		_bebf.Value = "\u0073\u006b\u0077"
	case ST_FTypeLin:
		_bebf.Value = "\u006c\u0069\u006e"
	case ST_FTypeNoBar:
		_bebf.Value = "\u006e\u006f\u0042a\u0072"
	}
	return _bebf, nil
}
func (_fdbg ST_TopBot) Validate() error { return _fdbg.ValidateWithPath("") }
func (_bdbf ST_Jc) Validate() error     { return _bdbf.ValidateWithPath("") }

// ValidateWithPath validates the CT_OMathArgPr and its children, prefixing error messages with path
func (_cbff *CT_OMathArgPr) ValidateWithPath(path string) error {
	if _cbff.ArgSz != nil {
		if _afff := _cbff.ArgSz.ValidateWithPath(path + "\u002f\u0041\u0072\u0067\u0053\u007a"); _afff != nil {
			return _afff
		}
	}
	return nil
}

type CT_AccPr struct {
	Chr    *CT_Char
	CtrlPr *CT_CtrlPr
}

func NewCT_SpacingRule() *CT_SpacingRule { _agcf := &CT_SpacingRule{}; _agcf.ValAttr = 0; return _agcf }

type CT_SPre struct {
	SPrePr *CT_SPrePr
	Sub    *CT_OMathArg
	Sup    *CT_OMathArg
	E      *CT_OMathArg
}

func NewOMath() *OMath { _acaaa := &OMath{}; _acaaa.CT_OMath = *NewCT_OMath(); return _acaaa }
func (_gggfg ST_BreakBin) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_gggfg.String(), start)
}

// Validate validates the CT_BreakBinSub and its children
func (_acgdf *CT_BreakBinSub) Validate() error {
	return _acgdf.ValidateWithPath("\u0043\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042i\u006e\u0053\u0075\u0062")
}

// ValidateWithPath validates the CT_Phant and its children, prefixing error messages with path
func (_gggd *CT_Phant) ValidateWithPath(path string) error {
	if _gggd.PhantPr != nil {
		if _cbed := _gggd.PhantPr.ValidateWithPath(path + "\u002f\u0050\u0068\u0061\u006e\u0074\u0050\u0072"); _cbed != nil {
			return _cbed
		}
	}
	if _cefa := _gggd.E.ValidateWithPath(path + "\u002f\u0045"); _cefa != nil {
		return _cefa
	}
	return nil
}

type CT_D struct {
	DPr *CT_DPr
	E   []*CT_OMathArg
}

func NewCT_BarPr() *CT_BarPr { _ced := &CT_BarPr{}; return _ced }
func (_efbg *CT_MC) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _efbg.McPr != nil {
		_gea := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0063\u0050\u0072"}}
		e.EncodeElement(_efbg.McPr, _gea)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_bgf *CT_Char) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", _bgf.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Bar and its children, prefixing error messages with path
func (_dgd *CT_Bar) ValidateWithPath(path string) error {
	if _dgd.BarPr != nil {
		if _cd := _dgd.BarPr.ValidateWithPath(path + "\u002f\u0042\u0061\u0072\u0050\u0072"); _cd != nil {
			return _cd
		}
	}
	if _fgg := _dgd.E.ValidateWithPath(path + "\u002f\u0045"); _fgg != nil {
		return _fgg
	}
	return nil
}

type CT_LimLowPr struct{ CtrlPr *CT_CtrlPr }

func (_dcfd *CT_SPrePr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dcfd.CtrlPr != nil {
		_bdcab := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_dcfd.CtrlPr, _bdcab)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func ParseUnionST_OnOff(s string) (_be.ST_OnOff, error) { return _be.ParseUnionST_OnOff(s) }
func (_aaad *CT_YAlign) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_ggec, _gecf := _aaad.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _gecf != nil {
		return _gecf
	}
	start.Attr = append(start.Attr, _ggec)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_BoxPr() *CT_BoxPr { _aab := &CT_BoxPr{}; return _aab }

type CT_Rad struct {
	RadPr *CT_RadPr
	Deg   *CT_OMathArg
	E     *CT_OMathArg
}

func NewCT_PhantPr() *CT_PhantPr { _dccbc := &CT_PhantPr{}; return _dccbc }

type OMath struct{ CT_OMath }

func (_fcde *CT_FuncPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fcde.CtrlPr != nil {
		_fcgg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_fcde.CtrlPr, _fcgg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_GroupChr and its children, prefixing error messages with path
func (_ecde *CT_GroupChr) ValidateWithPath(path string) error {
	if _ecde.GroupChrPr != nil {
		if _ecb := _ecde.GroupChrPr.ValidateWithPath(path + "/\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"); _ecb != nil {
			return _ecb
		}
	}
	if _cea := _ecde.E.ValidateWithPath(path + "\u002f\u0045"); _cea != nil {
		return _cea
	}
	return nil
}
func (_gcd *CT_CtrlPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for {
		_abgb, _fed := d.Token()
		if _fed != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0043\u0074\u0072l\u0050\u0072\u003a\u0020\u0025\u0073", _fed)
		}
		if _fac, _acd := _abgb.(_g.EndElement); _acd && _fac.Name == start.Name {
			break
		}
	}
	return nil
}
func (_bdb *EG_OMathElements) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fceb:
	for {
		_edgd, _ebgae := d.Token()
		if _ebgae != nil {
			return _ebgae
		}
		switch _gbcca := _edgd.(type) {
		case _g.StartElement:
			switch _gbcca.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_feebf := NewEG_OMathMathElements()
				_feebf.Acc = NewCT_Acc()
				if _fddde := d.DecodeElement(_feebf.Acc, &_gbcca); _fddde != nil {
					return _fddde
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _feebf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_eeba := NewEG_OMathMathElements()
				_eeba.Bar = NewCT_Bar()
				if _fbbca := d.DecodeElement(_eeba.Bar, &_gbcca); _fbbca != nil {
					return _fbbca
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _eeba)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_gbgb := NewEG_OMathMathElements()
				_gbgb.Box = NewCT_Box()
				if _ggdgc := d.DecodeElement(_gbgb.Box, &_gbcca); _ggdgc != nil {
					return _ggdgc
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _gbgb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_dfebc := NewEG_OMathMathElements()
				_dfebc.BorderBox = NewCT_BorderBox()
				if _acgda := d.DecodeElement(_dfebc.BorderBox, &_gbcca); _acgda != nil {
					return _acgda
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _dfebc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_ebbe := NewEG_OMathMathElements()
				_ebbe.D = NewCT_D()
				if _ebef := d.DecodeElement(_ebbe.D, &_gbcca); _ebef != nil {
					return _ebef
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _ebbe)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_fcfdg := NewEG_OMathMathElements()
				_fcfdg.EqArr = NewCT_EqArr()
				if _ggbgc := d.DecodeElement(_fcfdg.EqArr, &_gbcca); _ggbgc != nil {
					return _ggbgc
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _fcfdg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_bfgc := NewEG_OMathMathElements()
				_bfgc.F = NewCT_F()
				if _dbace := d.DecodeElement(_bfgc.F, &_gbcca); _dbace != nil {
					return _dbace
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _bfgc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_ddbdb := NewEG_OMathMathElements()
				_ddbdb.Func = NewCT_Func()
				if _cgbc := d.DecodeElement(_ddbdb.Func, &_gbcca); _cgbc != nil {
					return _cgbc
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _ddbdb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_bbbga := NewEG_OMathMathElements()
				_bbbga.GroupChr = NewCT_GroupChr()
				if _bebb := d.DecodeElement(_bbbga.GroupChr, &_gbcca); _bebb != nil {
					return _bebb
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _bbbga)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_gfdd := NewEG_OMathMathElements()
				_gfdd.LimLow = NewCT_LimLow()
				if _ffdgee := d.DecodeElement(_gfdd.LimLow, &_gbcca); _ffdgee != nil {
					return _ffdgee
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _gfdd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_bbdgd := NewEG_OMathMathElements()
				_bbdgd.LimUpp = NewCT_LimUpp()
				if _gadg := d.DecodeElement(_bbdgd.LimUpp, &_gbcca); _gadg != nil {
					return _gadg
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _bbdgd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_deab := NewEG_OMathMathElements()
				_deab.M = NewCT_M()
				if _aecd := d.DecodeElement(_deab.M, &_gbcca); _aecd != nil {
					return _aecd
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _deab)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_cfae := NewEG_OMathMathElements()
				_cfae.Nary = NewCT_Nary()
				if _acafb := d.DecodeElement(_cfae.Nary, &_gbcca); _acafb != nil {
					return _acafb
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _cfae)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_aada := NewEG_OMathMathElements()
				_aada.Phant = NewCT_Phant()
				if _aecc := d.DecodeElement(_aada.Phant, &_gbcca); _aecc != nil {
					return _aecc
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _aada)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_eada := NewEG_OMathMathElements()
				_eada.Rad = NewCT_Rad()
				if _gfcd := d.DecodeElement(_eada.Rad, &_gbcca); _gfcd != nil {
					return _gfcd
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _eada)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_gfed := NewEG_OMathMathElements()
				_gfed.SPre = NewCT_SPre()
				if _faec := d.DecodeElement(_gfed.SPre, &_gbcca); _faec != nil {
					return _faec
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _gfed)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_edde := NewEG_OMathMathElements()
				_edde.SSub = NewCT_SSub()
				if _fdag := d.DecodeElement(_edde.SSub, &_gbcca); _fdag != nil {
					return _fdag
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _edde)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_dafd := NewEG_OMathMathElements()
				_dafd.SSubSup = NewCT_SSubSup()
				if _bcedb := d.DecodeElement(_dafd.SSubSup, &_gbcca); _bcedb != nil {
					return _bcedb
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _dafd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_cfeg := NewEG_OMathMathElements()
				_cfeg.SSup = NewCT_SSup()
				if _dgca := d.DecodeElement(_cfeg.SSup, &_gbcca); _dgca != nil {
					return _dgca
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _cfeg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_bdcg := NewEG_OMathMathElements()
				_bdcg.R = NewCT_R()
				if _gedcg := d.DecodeElement(_bdcg.R, &_gbcca); _gedcg != nil {
					return _gedcg
				}
				_bdb.EG_OMathMathElements = append(_bdb.EG_OMathMathElements, _bdcg)
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u004f\u004d\u0061\u0074\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0020\u0025v", _gbcca.Name)
				if _dcbbc := d.Skip(); _dcbbc != nil {
					return _dcbbc
				}
			}
		case _g.EndElement:
			break _fceb
		case _g.CharData:
		}
	}
	return nil
}
func (_afbc *CT_Script) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _afbc.ValAttr != ST_ScriptUnset {
		_cgga, _cebab := _afbc.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _cebab != nil {
			return _cebab
		}
		start.Attr = append(start.Attr, _cgga)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_bdcfd ST_TopBot) String() string {
	switch _bdcfd {
	case 0:
		return ""
	case 1:
		return "\u0074\u006f\u0070"
	case 2:
		return "\u0062\u006f\u0074"
	}
	return ""
}
func (_cgb *CT_OMathJc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _agc := range start.Attr {
		if _agc.Name.Local == "\u0076\u0061\u006c" {
			_cgb.ValAttr.UnmarshalXMLAttr(_agc)
			continue
		}
	}
	for {
		_gdcf, _cbfag := d.Token()
		if _cbfag != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u004a\u0063\u003a\u0020%\u0073", _cbfag)
		}
		if _dcdg, _egce := _gdcf.(_g.EndElement); _egce && _dcdg.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_Integer255 struct{ ValAttr int64 }

func (_fdfbf ST_BreakBinSub) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_ebdd := _g.Attr{}
	_ebdd.Name = name
	switch _fdfbf {
	case ST_BreakBinSubUnset:
		_ebdd.Value = ""
	case ST_BreakBinSub__:
		_ebdd.Value = "\u002d\u002d"
	case ST_BreakBinSub___:
		_ebdd.Value = "\u002d\u002b"
	case ST_BreakBinSub____:
		_ebdd.Value = "\u002b\u002d"
	}
	return _ebdd, nil
}
func (_ffcf *CT_UnSignedInteger) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _dggb := range start.Attr {
		if _dggb.Name.Local == "\u0076\u0061\u006c" {
			_acge, _gecd := _e.ParseUint(_dggb.Value, 10, 32)
			if _gecd != nil {
				return _gecd
			}
			_ffcf.ValAttr = uint32(_acge)
			continue
		}
	}
	for {
		_gbacc, _ccbf := d.Token()
		if _ccbf != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054_\u0055\u006e\u0053\u0069\u0067\u006e\u0065d\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u003a\u0020\u0025\u0073", _ccbf)
		}
		if _dccaf, _bdec := _gbacc.(_g.EndElement); _bdec && _dccaf.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_FuncPr and its children, prefixing error messages with path
func (_ddba *CT_FuncPr) ValidateWithPath(path string) error {
	if _ddba.CtrlPr != nil {
		if _fea := _ddba.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _fea != nil {
			return _fea
		}
	}
	return nil
}

// Validate validates the CT_MPr and its children
func (_bagf *CT_MPr) Validate() error {
	return _bagf.ValidateWithPath("\u0043\u0054\u005f\u004d\u0050\u0072")
}
func (_fa *CT_AccPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fa.Chr != nil {
		_dc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063h\u0072"}}
		e.EncodeElement(_fa.Chr, _dc)
	}
	if _fa.CtrlPr != nil {
		_gf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_fa.CtrlPr, _gf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

const (
	ST_ShpUnset    ST_Shp = 0
	ST_ShpCentered ST_Shp = 1
	ST_ShpMatch    ST_Shp = 2
)

func NewMathPr() *MathPr { _baba := &MathPr{}; _baba.CT_MathPr = *NewCT_MathPr(); return _baba }

type EG_OMathMathElements struct {
	Acc       *CT_Acc
	Bar       *CT_Bar
	Box       *CT_Box
	BorderBox *CT_BorderBox
	D         *CT_D
	EqArr     *CT_EqArr
	F         *CT_F
	Func      *CT_Func
	GroupChr  *CT_GroupChr
	LimLow    *CT_LimLow
	LimUpp    *CT_LimUpp
	M         *CT_M
	Nary      *CT_Nary
	Phant     *CT_Phant
	Rad       *CT_Rad
	SPre      *CT_SPre
	SSub      *CT_SSub
	SSubSup   *CT_SSubSup
	SSup      *CT_SSup
	R         *CT_R
}

func (_cfaf *CT_PhantPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cfaf.Show != nil {
		_gaga := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0068\u006f\u0077"}}
		e.EncodeElement(_cfaf.Show, _gaga)
	}
	if _cfaf.ZeroWid != nil {
		_eedb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u007a\u0065\u0072\u006f\u0057\u0069d"}}
		e.EncodeElement(_cfaf.ZeroWid, _eedb)
	}
	if _cfaf.ZeroAsc != nil {
		_efaad := _g.StartElement{Name: _g.Name{Local: "\u006d:\u007a\u0065\u0072\u006f\u0041\u0073c"}}
		e.EncodeElement(_cfaf.ZeroAsc, _efaad)
	}
	if _cfaf.ZeroDesc != nil {
		_gffce := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u007a\u0065\u0072\u006f\u0044\u0065\u0073\u0063"}}
		e.EncodeElement(_cfaf.ZeroDesc, _gffce)
	}
	if _cfaf.Transp != nil {
		_ddde := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0074\u0072\u0061\u006e\u0073\u0070"}}
		e.EncodeElement(_cfaf.Transp, _ddde)
	}
	if _cfaf.CtrlPr != nil {
		_ecaf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_cfaf.CtrlPr, _ecaf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_XAlign() *CT_XAlign { _gegg := &CT_XAlign{}; _gegg.ValAttr = _be.ST_XAlign(1); return _gegg }
func NewCT_LimLoc() *CT_LimLoc { _aff := &CT_LimLoc{}; _aff.ValAttr = ST_LimLoc(1); return _aff }
func (_ffdc *CT_NaryPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ffdc.Chr != nil {
		_gabe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063h\u0072"}}
		e.EncodeElement(_ffdc.Chr, _gabe)
	}
	if _ffdc.LimLoc != nil {
		_cgabec := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u004c\u006f\u0063"}}
		e.EncodeElement(_ffdc.LimLoc, _cgabec)
	}
	if _ffdc.Grow != nil {
		_ecge := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067\u0072\u006f\u0077"}}
		e.EncodeElement(_ffdc.Grow, _ecge)
	}
	if _ffdc.SubHide != nil {
		_ggd := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0075\u0062\u0048\u0069\u0064e"}}
		e.EncodeElement(_ffdc.SubHide, _ggd)
	}
	if _ffdc.SupHide != nil {
		_gdgb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0075\u0070\u0048\u0069\u0064e"}}
		e.EncodeElement(_ffdc.SupHide, _gdgb)
	}
	if _ffdc.CtrlPr != nil {
		_cadg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ffdc.CtrlPr, _cadg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_BreakBin struct{ ValAttr ST_BreakBin }

func (_cbac ST_Shp) Validate() error { return _cbac.ValidateWithPath("") }
func (_gfbg *CT_Func) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gfbg.FName = NewCT_OMathArg()
	_gfbg.E = NewCT_OMathArg()
_adc:
	for {
		_dfga, _bae := d.Token()
		if _bae != nil {
			return _bae
		}
		switch _fcg := _dfga.(type) {
		case _g.StartElement:
			switch _fcg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063\u0050\u0072"}:
				_gfbg.FuncPr = NewCT_FuncPr()
				if _dde := d.DecodeElement(_gfbg.FuncPr, &_fcg); _dde != nil {
					return _dde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u004e\u0061m\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u004e\u0061m\u0065"}:
				if _bag := d.DecodeElement(_gfbg.FName, &_fcg); _bag != nil {
					return _bag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _eedf := d.DecodeElement(_gfbg.E, &_fcg); _eedf != nil {
					return _eedf
				}
			default:
				_gg.Log("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0046\u0075\u006e\u0063\u0020\u0025\u0076", _fcg.Name)
				if _fgbfb := d.Skip(); _fgbfb != nil {
					return _fgbfb
				}
			}
		case _g.EndElement:
			break _adc
		case _g.CharData:
		}
	}
	return nil
}
func (_cgfd *ST_TopBot) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_cgfd = 0
	case "\u0074\u006f\u0070":
		*_cgfd = 1
	case "\u0062\u006f\u0074":
		*_cgfd = 2
	}
	return nil
}

const (
	ST_TopBotUnset ST_TopBot = 0
	ST_TopBotTop   ST_TopBot = 1
	ST_TopBotBot   ST_TopBot = 2
)

// ValidateWithPath validates the CT_LimUpp and its children, prefixing error messages with path
func (_dgbb *CT_LimUpp) ValidateWithPath(path string) error {
	if _dgbb.LimUppPr != nil {
		if _egge := _dgbb.LimUppPr.ValidateWithPath(path + "\u002fL\u0069\u006d\u0055\u0070\u0070\u0050r"); _egge != nil {
			return _egge
		}
	}
	if _bbd := _dgbb.E.ValidateWithPath(path + "\u002f\u0045"); _bbd != nil {
		return _bbd
	}
	if _ffb := _dgbb.Lim.ValidateWithPath(path + "\u002f\u004c\u0069\u006d"); _ffb != nil {
		return _ffb
	}
	return nil
}
func NewCT_SSubSup() *CT_SSubSup {
	_cgdf := &CT_SSubSup{}
	_cgdf.E = NewCT_OMathArg()
	_cgdf.Sub = NewCT_OMathArg()
	_cgdf.Sup = NewCT_OMathArg()
	return _cgdf
}

// Validate validates the CT_RPRChoice and its children
func (_cgcd *CT_RPRChoice) Validate() error {
	return _cgcd.ValidateWithPath("\u0043\u0054\u005fR\u0050\u0052\u0043\u0068\u006f\u0069\u0063\u0065")
}
func (_cgcda *ST_Jc) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_cgcda = 0
	case "\u006c\u0065\u0066\u0074":
		*_cgcda = 1
	case "\u0072\u0069\u0067h\u0074":
		*_cgcda = 2
	case "\u0063\u0065\u006e\u0074\u0065\u0072":
		*_cgcda = 3
	case "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070":
		*_cgcda = 4
	}
	return nil
}
func (_degc *CT_OMath) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gbee:
	for {
		_gdgbe, _aeag := d.Token()
		if _aeag != nil {
			return _aeag
		}
		switch _cacf := _gdgbe.(type) {
		case _g.StartElement:
			switch _cacf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_ccca := NewEG_OMathMathElements()
				_ccca.Acc = NewCT_Acc()
				if _dfggc := d.DecodeElement(_ccca.Acc, &_cacf); _dfggc != nil {
					return _dfggc
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _ccca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_bagg := NewEG_OMathMathElements()
				_bagg.Bar = NewCT_Bar()
				if _ffac := d.DecodeElement(_bagg.Bar, &_cacf); _ffac != nil {
					return _ffac
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _bagg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_aeae := NewEG_OMathMathElements()
				_aeae.Box = NewCT_Box()
				if _cff := d.DecodeElement(_aeae.Box, &_cacf); _cff != nil {
					return _cff
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _aeae)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_fcfdb := NewEG_OMathMathElements()
				_fcfdb.BorderBox = NewCT_BorderBox()
				if _fcaa := d.DecodeElement(_fcfdb.BorderBox, &_cacf); _fcaa != nil {
					return _fcaa
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _fcfdb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_ceba := NewEG_OMathMathElements()
				_ceba.D = NewCT_D()
				if _fgd := d.DecodeElement(_ceba.D, &_cacf); _fgd != nil {
					return _fgd
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _ceba)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_cee := NewEG_OMathMathElements()
				_cee.EqArr = NewCT_EqArr()
				if _cef := d.DecodeElement(_cee.EqArr, &_cacf); _cef != nil {
					return _cef
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _cee)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_bgfc := NewEG_OMathMathElements()
				_bgfc.F = NewCT_F()
				if _ebde := d.DecodeElement(_bgfc.F, &_cacf); _ebde != nil {
					return _ebde
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _bgfc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_fddb := NewEG_OMathMathElements()
				_fddb.Func = NewCT_Func()
				if _bbeg := d.DecodeElement(_fddb.Func, &_cacf); _bbeg != nil {
					return _bbeg
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _fddb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_fcea := NewEG_OMathMathElements()
				_fcea.GroupChr = NewCT_GroupChr()
				if _geda := d.DecodeElement(_fcea.GroupChr, &_cacf); _geda != nil {
					return _geda
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _fcea)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_adef := NewEG_OMathMathElements()
				_adef.LimLow = NewCT_LimLow()
				if _gggf := d.DecodeElement(_adef.LimLow, &_cacf); _gggf != nil {
					return _gggf
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _adef)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_gaef := NewEG_OMathMathElements()
				_gaef.LimUpp = NewCT_LimUpp()
				if _baec := d.DecodeElement(_gaef.LimUpp, &_cacf); _baec != nil {
					return _baec
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _gaef)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_gfdc := NewEG_OMathMathElements()
				_gfdc.M = NewCT_M()
				if _fdce := d.DecodeElement(_gfdc.M, &_cacf); _fdce != nil {
					return _fdce
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _gfdc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_dfcd := NewEG_OMathMathElements()
				_dfcd.Nary = NewCT_Nary()
				if _fdcc := d.DecodeElement(_dfcd.Nary, &_cacf); _fdcc != nil {
					return _fdcc
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _dfcd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_ffdgd := NewEG_OMathMathElements()
				_ffdgd.Phant = NewCT_Phant()
				if _gdcb := d.DecodeElement(_ffdgd.Phant, &_cacf); _gdcb != nil {
					return _gdcb
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _ffdgd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_bcfb := NewEG_OMathMathElements()
				_bcfb.Rad = NewCT_Rad()
				if _gfdf := d.DecodeElement(_bcfb.Rad, &_cacf); _gfdf != nil {
					return _gfdf
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _bcfb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_dada := NewEG_OMathMathElements()
				_dada.SPre = NewCT_SPre()
				if _aacb := d.DecodeElement(_dada.SPre, &_cacf); _aacb != nil {
					return _aacb
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _dada)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_bdfab := NewEG_OMathMathElements()
				_bdfab.SSub = NewCT_SSub()
				if _cdec := d.DecodeElement(_bdfab.SSub, &_cacf); _cdec != nil {
					return _cdec
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _bdfab)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_cggg := NewEG_OMathMathElements()
				_cggg.SSubSup = NewCT_SSubSup()
				if _aaef := d.DecodeElement(_cggg.SSubSup, &_cacf); _aaef != nil {
					return _aaef
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _cggg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_eadc := NewEG_OMathMathElements()
				_eadc.SSup = NewCT_SSup()
				if _gefa := d.DecodeElement(_eadc.SSup, &_cacf); _gefa != nil {
					return _gefa
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _eadc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_dcge := NewEG_OMathMathElements()
				_dcge.R = NewCT_R()
				if _ggac := d.DecodeElement(_dcge.R, &_cacf); _ggac != nil {
					return _ggac
				}
				_degc.EG_OMathMathElements = append(_degc.EG_OMathMathElements, _dcge)
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0020\u0025\u0076", _cacf.Name)
				if _gafb := d.Skip(); _gafb != nil {
					return _gafb
				}
			}
		case _g.EndElement:
			break _gbee
		case _g.CharData:
		}
	}
	return nil
}
func (_egfc *CT_OnOff) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _egfc.ValAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_egfc.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_gdbb ST_Script) String() string {
	switch _gdbb {
	case 0:
		return ""
	case 1:
		return "\u0072\u006f\u006da\u006e"
	case 2:
		return "\u0073\u0063\u0072\u0069\u0070\u0074"
	case 3:
		return "\u0066r\u0061\u006b\u0074\u0075\u0072"
	case 4:
		return "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b"
	case 5:
		return "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066"
	case 6:
		return "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e"
	}
	return ""
}

type CT_LimLow struct {
	LimLowPr *CT_LimLowPr
	E        *CT_OMathArg
	Lim      *CT_OMathArg
}

func NewCT_DPr() *CT_DPr { _ece := &CT_DPr{}; return _ece }
func (_ede *CT_Func) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ede.FuncPr != nil {
		_cdeg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066\u0075\u006e\u0063\u0050\u0072"}}
		e.EncodeElement(_ede.FuncPr, _cdeg)
	}
	_fdf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0066\u004e\u0061\u006d\u0065"}}
	e.EncodeElement(_ede.FName, _fdf)
	_ggb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_ede.E, _ggb)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_OMath and its children
func (_bdcc *CT_OMath) Validate() error {
	return _bdcc.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068")
}
func (_egcf ST_Shp) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_aafd := _g.Attr{}
	_aafd.Name = name
	switch _egcf {
	case ST_ShpUnset:
		_aafd.Value = ""
	case ST_ShpCentered:
		_aafd.Value = "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064"
	case ST_ShpMatch:
		_aafd.Value = "\u006d\u0061\u0074c\u0068"
	}
	return _aafd, nil
}
func (_ffbd *CT_String) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _ffbd.ValAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_ffbd.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_aedd *CT_ManualBreak) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _bgb := range start.Attr {
		if _bgb.Name.Local == "\u0061\u006c\u006eA\u0074" {
			_fbcf, _fecf := _e.ParseInt(_bgb.Value, 10, 64)
			if _fecf != nil {
				return _fecf
			}
			_aedd.AlnAtAttr = &_fbcf
			continue
		}
	}
	for {
		_geba, _gcaa := d.Token()
		if _gcaa != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fM\u0061\u006e\u0075\u0061\u006c\u0042\u0072\u0065\u0061\u006b:\u0020\u0025\u0073", _gcaa)
		}
		if _befc, _gfe := _geba.(_g.EndElement); _gfe && _befc.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SpacingRule and its children
func (_dadd *CT_SpacingRule) Validate() error {
	return _dadd.ValidateWithPath("\u0043\u0054\u005f\u0053\u0070\u0061\u0063\u0069\u006eg\u0052\u0075\u006c\u0065")
}

// Validate validates the CT_CtrlPr and its children
func (_ddfa *CT_CtrlPr) Validate() error {
	return _ddfa.ValidateWithPath("\u0043T\u005f\u0043\u0074\u0072\u006c\u0050r")
}
func (_gga *CT_BoxPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_bad:
	for {
		_ggaa, _gaf := d.Token()
		if _gaf != nil {
			return _gaf
		}
		switch _acg := _ggaa.(type) {
		case _g.StartElement:
			switch _acg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006f\u0070\u0045m\u0075"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006f\u0070\u0045m\u0075"}:
				_gga.OpEmu = NewCT_OnOff()
				if _bba := d.DecodeElement(_gga.OpEmu, &_acg); _bba != nil {
					return _bba
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006eo\u0042\u0072\u0065\u0061\u006b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006eo\u0042\u0072\u0065\u0061\u006b"}:
				_gga.NoBreak = NewCT_OnOff()
				if _efg := d.DecodeElement(_gga.NoBreak, &_acg); _efg != nil {
					return _efg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0069\u0066\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0069\u0066\u0066"}:
				_gga.Diff = NewCT_OnOff()
				if _feg := d.DecodeElement(_gga.Diff, &_acg); _feg != nil {
					return _feg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b"}:
				_gga.Brk = NewCT_ManualBreak()
				if _cgc := d.DecodeElement(_gga.Brk, &_acg); _cgc != nil {
					return _cgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u006c\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u006c\u006e"}:
				_gga.Aln = NewCT_OnOff()
				if _gde := d.DecodeElement(_gga.Aln, &_acg); _gde != nil {
					return _gde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_gga.CtrlPr = NewCT_CtrlPr()
				if _acgd := d.DecodeElement(_gga.CtrlPr, &_acg); _acgd != nil {
					return _acgd
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u006f\u0078\u0050\u0072\u0020\u0025\u0076", _acg.Name)
				if _abe := d.Skip(); _abe != nil {
					return _abe
				}
			}
		case _g.EndElement:
			break _bad
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_PhantPr and its children, prefixing error messages with path
func (_dfa *CT_PhantPr) ValidateWithPath(path string) error {
	if _dfa.Show != nil {
		if _gfbb := _dfa.Show.ValidateWithPath(path + "\u002f\u0053\u0068o\u0077"); _gfbb != nil {
			return _gfbb
		}
	}
	if _dfa.ZeroWid != nil {
		if _gefc := _dfa.ZeroWid.ValidateWithPath(path + "\u002f\u005a\u0065\u0072\u006f\u0057\u0069\u0064"); _gefc != nil {
			return _gefc
		}
	}
	if _dfa.ZeroAsc != nil {
		if _egac := _dfa.ZeroAsc.ValidateWithPath(path + "\u002f\u005a\u0065\u0072\u006f\u0041\u0073\u0063"); _egac != nil {
			return _egac
		}
	}
	if _dfa.ZeroDesc != nil {
		if _bdfbd := _dfa.ZeroDesc.ValidateWithPath(path + "\u002fZ\u0065\u0072\u006f\u0044\u0065\u0073c"); _bdfbd != nil {
			return _bdfbd
		}
	}
	if _dfa.Transp != nil {
		if _decg := _dfa.Transp.ValidateWithPath(path + "\u002fT\u0072\u0061\u006e\u0073\u0070"); _decg != nil {
			return _decg
		}
	}
	if _dfa.CtrlPr != nil {
		if _gebg := _dfa.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _gebg != nil {
			return _gebg
		}
	}
	return nil
}
func (_ebded *ST_BreakBinSub) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_eeef, _cege := d.Token()
	if _cege != nil {
		return _cege
	}
	if _beaab, _adge := _eeef.(_g.EndElement); _adge && _beaab.Name == start.Name {
		*_ebded = 1
		return nil
	}
	if _eaea, _gebc := _eeef.(_g.CharData); !_gebc {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eeef)
	} else {
		switch string(_eaea) {
		case "":
			*_ebded = 0
		case "\u002d\u002d":
			*_ebded = 1
		case "\u002d\u002b":
			*_ebded = 2
		case "\u002b\u002d":
			*_ebded = 3
		}
	}
	_eeef, _cege = d.Token()
	if _cege != nil {
		return _cege
	}
	if _ecefe, _abacg := _eeef.(_g.EndElement); _abacg && _ecefe.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eeef)
}
func (_ccec *CT_LimUpp) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ccec.E = NewCT_OMathArg()
	_ccec.Lim = NewCT_OMathArg()
_eeb:
	for {
		_dccfe, _babe := d.Token()
		if _babe != nil {
			return _babe
		}
		switch _faf := _dccfe.(type) {
		case _g.StartElement:
			switch _faf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070\u0050\u0072"}:
				_ccec.LimUppPr = NewCT_LimUppPr()
				if _eedg := d.DecodeElement(_ccec.LimUppPr, &_faf); _eedg != nil {
					return _eedg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _efde := d.DecodeElement(_ccec.E, &_faf); _efde != nil {
					return _efde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d"}:
				if _edb := d.DecodeElement(_ccec.Lim, &_faf); _edb != nil {
					return _edb
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004c\u0069\u006dU\u0070\u0070 \u0025\u0076", _faf.Name)
				if _fbf := d.Skip(); _fbf != nil {
					return _fbf
				}
			}
		case _g.EndElement:
			break _eeb
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_OMathArg() *CT_OMathArg { _ebdg := &CT_OMathArg{}; return _ebdg }

type CT_Shp struct{ ValAttr ST_Shp }

// ValidateWithPath validates the CT_BoxPr and its children, prefixing error messages with path
func (_afe *CT_BoxPr) ValidateWithPath(path string) error {
	if _afe.OpEmu != nil {
		if _df := _afe.OpEmu.ValidateWithPath(path + "\u002f\u004f\u0070\u0045\u006d\u0075"); _df != nil {
			return _df
		}
	}
	if _afe.NoBreak != nil {
		if _dee := _afe.NoBreak.ValidateWithPath(path + "\u002f\u004e\u006f\u0042\u0072\u0065\u0061\u006b"); _dee != nil {
			return _dee
		}
	}
	if _afe.Diff != nil {
		if _cfb := _afe.Diff.ValidateWithPath(path + "\u002f\u0044\u0069f\u0066"); _cfb != nil {
			return _cfb
		}
	}
	if _afe.Brk != nil {
		if _egaf := _afe.Brk.ValidateWithPath(path + "\u002f\u0042\u0072\u006b"); _egaf != nil {
			return _egaf
		}
	}
	if _afe.Aln != nil {
		if _gfc := _afe.Aln.ValidateWithPath(path + "\u002f\u0041\u006c\u006e"); _gfc != nil {
			return _gfc
		}
	}
	if _afe.CtrlPr != nil {
		if _cda := _afe.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _cda != nil {
			return _cda
		}
	}
	return nil
}
func (_ege *CT_BarPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ege.Pos != nil {
		_fdd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0070o\u0073"}}
		e.EncodeElement(_ege.Pos, _fdd)
	}
	if _ege.CtrlPr != nil {
		_ba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ege.CtrlPr, _ba)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_ManualBreak() *CT_ManualBreak { _ebc := &CT_ManualBreak{}; return _ebc }

// ValidateWithPath validates the CT_RPRChoice and its children, prefixing error messages with path
func (_gbfc *CT_RPRChoice) ValidateWithPath(path string) error {
	if _gbfc.Nor != nil {
		if _fdcfc := _gbfc.Nor.ValidateWithPath(path + "\u002f\u004e\u006f\u0072"); _fdcfc != nil {
			return _fdcfc
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Nary and its children, prefixing error messages with path
func (_abgff *CT_Nary) ValidateWithPath(path string) error {
	if _abgff.NaryPr != nil {
		if _fba := _abgff.NaryPr.ValidateWithPath(path + "\u002fN\u0061\u0072\u0079\u0050\u0072"); _fba != nil {
			return _fba
		}
	}
	if _fbad := _abgff.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _fbad != nil {
		return _fbad
	}
	if _gadaa := _abgff.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _gadaa != nil {
		return _gadaa
	}
	if _fgaag := _abgff.E.ValidateWithPath(path + "\u002f\u0045"); _fgaag != nil {
		return _fgaag
	}
	return nil
}
func (_adcc ST_Style) String() string {
	switch _adcc {
	case 0:
		return ""
	case 1:
		return "\u0070"
	case 2:
		return "\u0062"
	case 3:
		return "\u0069"
	case 4:
		return "\u0062\u0069"
	}
	return ""
}

// Validate validates the CT_SSupPr and its children
func (_cgccc *CT_SSupPr) Validate() error {
	return _cgccc.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0070\u0050r")
}
func NewCT_YAlign() *CT_YAlign { _beed := &CT_YAlign{}; _beed.ValAttr = _be.ST_YAlign(1); return _beed }
func (_cdaa *CT_SSup) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cdaa.E = NewCT_OMathArg()
	_cdaa.Sup = NewCT_OMathArg()
_eade:
	for {
		_baef, _gdcc := d.Token()
		if _gdcc != nil {
			return _gdcc
		}
		switch _cccf := _baef.(type) {
		case _g.StartElement:
			switch _cccf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070\u0050\u0072"}:
				_cdaa.SSupPr = NewCT_SSupPr()
				if _bgfg := d.DecodeElement(_cdaa.SSupPr, &_cccf); _bgfg != nil {
					return _bgfg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _fedge := d.DecodeElement(_cdaa.E, &_cccf); _fedge != nil {
					return _fedge
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _fffd := d.DecodeElement(_cdaa.Sup, &_cccf); _fffd != nil {
					return _fffd
				}
			default:
				_gg.Log("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0053\u0053\u0075\u0070\u0020\u0025\u0076", _cccf.Name)
				if _fbfe := d.Skip(); _fbfe != nil {
					return _fbfe
				}
			}
		case _g.EndElement:
			break _eade
		case _g.CharData:
		}
	}
	return nil
}

type CT_RPRChoice struct{ Nor *CT_OnOff }

// Validate validates the EG_OMathMathElements and its children
func (_bcdf *EG_OMathMathElements) Validate() error {
	return _bcdf.ValidateWithPath("E\u0047_\u004f\u004d\u0061\u0074\u0068\u004d\u0061\u0074h\u0045\u006c\u0065\u006den\u0074\u0073")
}

type CT_MC struct{ McPr *CT_MCPr }

func NewCT_FuncPr() *CT_FuncPr { _gacc := &CT_FuncPr{}; return _gacc }

// ValidateWithPath validates the CT_TopBot and its children, prefixing error messages with path
func (_ggeag *CT_TopBot) ValidateWithPath(path string) error {
	if _ggeag.ValAttr == ST_TopBotUnset {
		return _f.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _gfee := _ggeag.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gfee != nil {
		return _gfee
	}
	return nil
}
func NewCT_SPrePr() *CT_SPrePr { _cdbb := &CT_SPrePr{}; return _cdbb }
func (_cdcg *CT_MR) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_dbbd:
	for {
		_bbef, _fbdd := d.Token()
		if _fbdd != nil {
			return _fbdd
		}
		switch _gcg := _bbef.(type) {
		case _g.StartElement:
			switch _gcg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				_eafb := NewCT_OMathArg()
				if _bbc := d.DecodeElement(_eafb, &_gcg); _bbc != nil {
					return _bbc
				}
				_cdcg.E = append(_cdcg.E, _eafb)
			default:
				_gg.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0043\u0054\u005fM\u0052 \u0025\u0076", _gcg.Name)
				if _cfea := d.Skip(); _cfea != nil {
					return _cfea
				}
			}
		case _g.EndElement:
			break _dbbd
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Rad and its children, prefixing error messages with path
func (_eaff *CT_Rad) ValidateWithPath(path string) error {
	if _eaff.RadPr != nil {
		if _gggb := _eaff.RadPr.ValidateWithPath(path + "\u002f\u0052\u0061\u0064\u0050\u0072"); _gggb != nil {
			return _gggb
		}
	}
	if _cbfaf := _eaff.Deg.ValidateWithPath(path + "\u002f\u0044\u0065\u0067"); _cbfaf != nil {
		return _cbfaf
	}
	if _cfcab := _eaff.E.ValidateWithPath(path + "\u002f\u0045"); _cfcab != nil {
		return _cfcab
	}
	return nil
}

type CT_Script struct{ ValAttr ST_Script }
type CT_OMathParaPr struct{ Jc *CT_OMathJc }

// Validate validates the CT_OMathParaPr and its children
func (_ffda *CT_OMathParaPr) Validate() error {
	return _ffda.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0050a\u0072\u0061\u0050\u0072")
}
func (_eedd *CT_LimLoc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_bdc, _acga := _eedd.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _acga != nil {
		return _acga
	}
	start.Attr = append(start.Attr, _bdc)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewEG_OMathMathElements() *EG_OMathMathElements { _gdbc := &EG_OMathMathElements{}; return _gdbc }

// ValidateWithPath validates the CT_OMath and its children, prefixing error messages with path
func (_fbddf *CT_OMath) ValidateWithPath(path string) error {
	for _dbbc, _egaae := range _fbddf.EG_OMathMathElements {
		if _eagd := _egaae.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0045\u0047\u005f\u004fM\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006ce\u006d\u0065\u006et\u0073[\u0025\u0064\u005d", path, _dbbc)); _eagd != nil {
			return _eagd
		}
	}
	return nil
}
func NewCT_RadPr() *CT_RadPr                     { _fedbb := &CT_RadPr{}; return _fedbb }
func NewCT_UnSignedInteger() *CT_UnSignedInteger { _dbba := &CT_UnSignedInteger{}; return _dbba }
func (_babb *CT_F) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_babb.Num = NewCT_OMathArg()
	_babb.Den = NewCT_OMathArg()
_dbaf:
	for {
		_dbcc, _bca := d.Token()
		if _bca != nil {
			return _bca
		}
		switch _aea := _dbcc.(type) {
		case _g.StartElement:
			switch _aea.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0050\u0072"}:
				_babb.FPr = NewCT_FPr()
				if _cae := d.DecodeElement(_babb.FPr, &_aea); _cae != nil {
					return _cae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0075\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0075\u006d"}:
				if _ebe := d.DecodeElement(_babb.Num, &_aea); _ebe != nil {
					return _ebe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u006e"}:
				if _egdf := d.DecodeElement(_babb.Den, &_aea); _egdf != nil {
					return _egdf
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0046\u0020\u0025\u0076", _aea.Name)
				if _fff := d.Skip(); _fff != nil {
					return _fff
				}
			}
		case _g.EndElement:
			break _dbaf
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_Phant() *CT_Phant { _ebaf := &CT_Phant{}; _ebaf.E = NewCT_OMathArg(); return _ebaf }

type CT_XAlign struct{ ValAttr _be.ST_XAlign }

func (_fcgb *EG_ScriptStyle) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Name.Local = "\u006d\u003aE\u0047\u005f\u0053c\u0072\u0069\u0070\u0074\u0053\u0074\u0079\u006c\u0065"
	if _fcgb.Scr != nil {
		_ecba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073c\u0072"}}
		e.EncodeElement(_fcgb.Scr, _ecba)
	}
	if _fcgb.Sty != nil {
		_ffabe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073t\u0079"}}
		e.EncodeElement(_fcgb.Sty, _ffabe)
	}
	return nil
}
func (_dfdf *CT_OMathArg) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dfdf.ArgPr != nil {
		_gdbe := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0061\u0072\u0067\u0050\u0072"}}
		e.EncodeElement(_dfdf.ArgPr, _gdbe)
	}
	if _dfdf.EG_OMathMathElements != nil {
		for _, _efdc := range _dfdf.EG_OMathMathElements {
			_efdc.MarshalXML(e, _g.StartElement{})
		}
	}
	if _dfdf.CtrlPr != nil {
		_abdc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_dfdf.CtrlPr, _abdc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_RPR struct {
	Lit    *CT_OnOff
	Choice *CT_RPRChoice
	Brk    *CT_ManualBreak
	Aln    *CT_OnOff
}

func NewCT_RChoice() *CT_RChoice { _gdfe := &CT_RChoice{}; return _gdfe }
func (_gbc *CT_MathPrChoice) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _gbc.WrapIndent != nil {
		_eag := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0077r\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}}
		e.EncodeElement(_gbc.WrapIndent, _eag)
	}
	if _gbc.WrapRight != nil {
		_egad := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0077\u0072\u0061\u0070\u0052\u0069\u0067\u0068\u0074"}}
		e.EncodeElement(_gbc.WrapRight, _egad)
	}
	return nil
}

// ValidateWithPath validates the CT_SSup and its children, prefixing error messages with path
func (_ccfe *CT_SSup) ValidateWithPath(path string) error {
	if _ccfe.SSupPr != nil {
		if _gfbe := _ccfe.SSupPr.ValidateWithPath(path + "\u002fS\u0053\u0075\u0070\u0050\u0072"); _gfbe != nil {
			return _gfbe
		}
	}
	if _adfe := _ccfe.E.ValidateWithPath(path + "\u002f\u0045"); _adfe != nil {
		return _adfe
	}
	if _aeef := _ccfe.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _aeef != nil {
		return _aeef
	}
	return nil
}
func NewCT_RPR() *CT_RPR { _defcc := &CT_RPR{}; return _defcc }

// Validate validates the CT_AccPr and its children
func (_bd *CT_AccPr) Validate() error {
	return _bd.ValidateWithPath("\u0043\u0054\u005f\u0041\u0063\u0063\u0050\u0072")
}
func (_gdg *CT_D) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gdg.DPr != nil {
		_gba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064P\u0072"}}
		e.EncodeElement(_gdg.DPr, _gba)
	}
	_dfc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	for _, _edc := range _gdg.E {
		e.EncodeElement(_edc, _dfc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_GroupChrPr and its children
func (_baff *CT_GroupChrPr) Validate() error {
	return _baff.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072")
}
func (_fdaf *CT_MC) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gfd:
	for {
		_eadd, _bfd := d.Token()
		if _bfd != nil {
			return _bfd
		}
		switch _afgc := _eadd.(type) {
		case _g.StartElement:
			switch _afgc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063\u0050\u0072"}:
				_fdaf.McPr = NewCT_MCPr()
				if _gbd := d.DecodeElement(_fdaf.McPr, &_afgc); _gbd != nil {
					return _gbd
				}
			default:
				_gg.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0043\u0054\u005fM\u0043 \u0025\u0076", _afgc.Name)
				if _bgcg := d.Skip(); _bgcg != nil {
					return _bgcg
				}
			}
		case _g.EndElement:
			break _gfd
		case _g.CharData:
		}
	}
	return nil
}

type CT_Char struct{ ValAttr string }
type CT_SSubSup struct {
	SSubSupPr *CT_SSubSupPr
	E         *CT_OMathArg
	Sub       *CT_OMathArg
	Sup       *CT_OMathArg
}

func (_ddgf *CT_OMathArg) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gced:
	for {
		_ggcd, _bfcf := d.Token()
		if _bfcf != nil {
			return _bfcf
		}
		switch _beeg := _ggcd.(type) {
		case _g.StartElement:
			switch _beeg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0072\u0067P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0072\u0067P\u0072"}:
				_ddgf.ArgPr = NewCT_OMathArgPr()
				if _gefe := d.DecodeElement(_ddgf.ArgPr, &_beeg); _gefe != nil {
					return _gefe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_ecac := NewEG_OMathMathElements()
				_ecac.Acc = NewCT_Acc()
				if _agdg := d.DecodeElement(_ecac.Acc, &_beeg); _agdg != nil {
					return _agdg
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _ecac)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_bddab := NewEG_OMathMathElements()
				_bddab.Bar = NewCT_Bar()
				if _dcea := d.DecodeElement(_bddab.Bar, &_beeg); _dcea != nil {
					return _dcea
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _bddab)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_fae := NewEG_OMathMathElements()
				_fae.Box = NewCT_Box()
				if _bcafc := d.DecodeElement(_fae.Box, &_beeg); _bcafc != nil {
					return _bcafc
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _fae)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_ddad := NewEG_OMathMathElements()
				_ddad.BorderBox = NewCT_BorderBox()
				if _gedb := d.DecodeElement(_ddad.BorderBox, &_beeg); _gedb != nil {
					return _gedb
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _ddad)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_bacea := NewEG_OMathMathElements()
				_bacea.D = NewCT_D()
				if _ggbag := d.DecodeElement(_bacea.D, &_beeg); _ggbag != nil {
					return _ggbag
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _bacea)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_acb := NewEG_OMathMathElements()
				_acb.EqArr = NewCT_EqArr()
				if _fafe := d.DecodeElement(_acb.EqArr, &_beeg); _fafe != nil {
					return _fafe
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _acb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_edfbgg := NewEG_OMathMathElements()
				_edfbgg.F = NewCT_F()
				if _dcdf := d.DecodeElement(_edfbgg.F, &_beeg); _dcdf != nil {
					return _dcdf
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _edfbgg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_eaad := NewEG_OMathMathElements()
				_eaad.Func = NewCT_Func()
				if _dcff := d.DecodeElement(_eaad.Func, &_beeg); _dcff != nil {
					return _dcff
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _eaad)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_ggfd := NewEG_OMathMathElements()
				_ggfd.GroupChr = NewCT_GroupChr()
				if _aaed := d.DecodeElement(_ggfd.GroupChr, &_beeg); _aaed != nil {
					return _aaed
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _ggfd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_cefb := NewEG_OMathMathElements()
				_cefb.LimLow = NewCT_LimLow()
				if _cefg := d.DecodeElement(_cefb.LimLow, &_beeg); _cefg != nil {
					return _cefg
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _cefb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_ded := NewEG_OMathMathElements()
				_ded.LimUpp = NewCT_LimUpp()
				if _fecd := d.DecodeElement(_ded.LimUpp, &_beeg); _fecd != nil {
					return _fecd
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _ded)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_agdgg := NewEG_OMathMathElements()
				_agdgg.M = NewCT_M()
				if _ggdg := d.DecodeElement(_agdgg.M, &_beeg); _ggdg != nil {
					return _ggdg
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _agdgg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_gdff := NewEG_OMathMathElements()
				_gdff.Nary = NewCT_Nary()
				if _fdfc := d.DecodeElement(_gdff.Nary, &_beeg); _fdfc != nil {
					return _fdfc
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _gdff)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_dgbc := NewEG_OMathMathElements()
				_dgbc.Phant = NewCT_Phant()
				if _eagdb := d.DecodeElement(_dgbc.Phant, &_beeg); _eagdb != nil {
					return _eagdb
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _dgbc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_bfggg := NewEG_OMathMathElements()
				_bfggg.Rad = NewCT_Rad()
				if _babed := d.DecodeElement(_bfggg.Rad, &_beeg); _babed != nil {
					return _babed
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _bfggg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_cegd := NewEG_OMathMathElements()
				_cegd.SPre = NewCT_SPre()
				if _bdca := d.DecodeElement(_cegd.SPre, &_beeg); _bdca != nil {
					return _bdca
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _cegd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_gbcb := NewEG_OMathMathElements()
				_gbcb.SSub = NewCT_SSub()
				if _bdgg := d.DecodeElement(_gbcb.SSub, &_beeg); _bdgg != nil {
					return _bdgg
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _gbcb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_bbf := NewEG_OMathMathElements()
				_bbf.SSubSup = NewCT_SSubSup()
				if _aabea := d.DecodeElement(_bbf.SSubSup, &_beeg); _aabea != nil {
					return _aabea
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _bbf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_gbde := NewEG_OMathMathElements()
				_gbde.SSup = NewCT_SSup()
				if _cdbe := d.DecodeElement(_gbde.SSup, &_beeg); _cdbe != nil {
					return _cdbe
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _gbde)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_ddbf := NewEG_OMathMathElements()
				_ddbf.R = NewCT_R()
				if _gcgf := d.DecodeElement(_ddbf.R, &_beeg); _gcgf != nil {
					return _gcgf
				}
				_ddgf.EG_OMathMathElements = append(_ddgf.EG_OMathMathElements, _ddbf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ddgf.CtrlPr = NewCT_CtrlPr()
				if _gfbd := d.DecodeElement(_ddgf.CtrlPr, &_beeg); _gfbd != nil {
					return _gfbd
				}
			default:
				_gg.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067\u0020\u0025\u0076", _beeg.Name)
				if _edbe := d.Skip(); _edbe != nil {
					return _edbe
				}
			}
		case _g.EndElement:
			break _gced
		case _g.CharData:
		}
	}
	return nil
}
func (_ebegg *ST_Shp) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_ebegg = 0
	case "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064":
		*_ebegg = 1
	case "\u006d\u0061\u0074c\u0068":
		*_ebegg = 2
	}
	return nil
}
func (_adefb ST_Script) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_adefb.String(), start)
}

type CT_OnOff struct{ ValAttr *_be.ST_OnOff }

// ValidateWithPath validates the CT_OMathJc and its children, prefixing error messages with path
func (_edfgf *CT_OMathJc) ValidateWithPath(path string) error {
	if _efaa := _edfgf.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _efaa != nil {
		return _efaa
	}
	return nil
}

// Validate validates the CT_OMathArgPr and its children
func (_bada *CT_OMathArgPr) Validate() error {
	return _bada.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067\u0050\u0072")
}
func (_gdfb *CT_Text) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _gagbf := range start.Attr {
		if _gagbf.Name.Space == "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065" && _gagbf.Name.Local == "\u0073\u0070\u0061c\u0065" {
			_fdddc, _aafc := _gagbf.Value, error(nil)
			if _aafc != nil {
				return _aafc
			}
			_gdfb.SpaceAttr = &_fdddc
			continue
		}
	}
	for {
		_fcga, _fbbd := d.Token()
		if _fbbd != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0054\u0065\u0078\u0074\u003a\u0020\u0025\u0073", _fbbd)
		}
		if _cgac, _fcac := _fcga.(_g.CharData); _fcac {
			_gdfb.Content = string(_cgac)
		}
		if _fagcg, _bcfbe := _fcga.(_g.EndElement); _bcfbe && _fagcg.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_BorderBoxPr() *CT_BorderBoxPr { _def := &CT_BorderBoxPr{}; return _def }

type CT_SSubSupPr struct {
	AlnScr *CT_OnOff
	CtrlPr *CT_CtrlPr
}
type CT_NaryPr struct {
	Chr     *CT_Char
	LimLoc  *CT_LimLoc
	Grow    *CT_OnOff
	SubHide *CT_OnOff
	SupHide *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

// ValidateWithPath validates the CT_M and its children, prefixing error messages with path
func (_gaff *CT_M) ValidateWithPath(path string) error {
	if _gaff.MPr != nil {
		if _adab := _gaff.MPr.ValidateWithPath(path + "\u002f\u004d\u0050\u0072"); _adab != nil {
			return _adab
		}
	}
	for _ggaaa, _bbag := range _gaff.Mr {
		if _aabd := _bbag.ValidateWithPath(_f.Sprintf("\u0025s\u002f\u004d\u0072\u005b\u0025\u0064]", path, _ggaaa)); _aabd != nil {
			return _aabd
		}
	}
	return nil
}
func (_gefdg *EG_OMathElements) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _gefdg.EG_OMathMathElements != nil {
		for _, _edeb := range _gefdg.EG_OMathMathElements {
			_edeb.MarshalXML(e, _g.StartElement{})
		}
	}
	return nil
}
func (_bdfdg *CT_SSubSup) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bdfdg.E = NewCT_OMathArg()
	_bdfdg.Sub = NewCT_OMathArg()
	_bdfdg.Sup = NewCT_OMathArg()
_cegdc:
	for {
		_aabb, _gdebe := d.Token()
		if _gdebe != nil {
			return _gdebe
		}
		switch _ggea := _aabb.(type) {
		case _g.StartElement:
			switch _ggea.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070\u0050r"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070\u0050r"}:
				_bdfdg.SSubSupPr = NewCT_SSubSupPr()
				if _cbaac := d.DecodeElement(_bdfdg.SSubSupPr, &_ggea); _cbaac != nil {
					return _cbaac
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _bfec := d.DecodeElement(_bdfdg.E, &_ggea); _bfec != nil {
					return _bfec
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _ggcf := d.DecodeElement(_bdfdg.Sub, &_ggea); _ggcf != nil {
					return _ggcf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _aege := d.DecodeElement(_bdfdg.Sup, &_ggea); _aege != nil {
					return _aege
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fS\u0053\u0075\u0062\u0053\u0075\u0070\u0020\u0025\u0076", _ggea.Name)
				if _cedd := d.Skip(); _cedd != nil {
					return _cedd
				}
			}
		case _g.EndElement:
			break _cegdc
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RadPr and its children
func (_edgf *CT_RadPr) Validate() error {
	return _edgf.ValidateWithPath("\u0043\u0054\u005f\u0052\u0061\u0064\u0050\u0072")
}
func (_ebec *CT_Nary) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ebec.NaryPr != nil {
		_cafdc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006e\u0061\u0072\u0079\u0050\u0072"}}
		e.EncodeElement(_ebec.NaryPr, _cafdc)
	}
	_gdaa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_ebec.Sub, _gdaa)
	_gadf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_ebec.Sup, _gadf)
	_fggb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_ebec.E, _fggb)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_SSub() *CT_SSub {
	_fdff := &CT_SSub{}
	_fdff.E = NewCT_OMathArg()
	_fdff.Sub = NewCT_OMathArg()
	return _fdff
}

type CT_OMathPara struct {
	OMathParaPr *CT_OMathParaPr
	OMath       []*CT_OMath
}

// ValidateWithPath validates the CT_Acc and its children, prefixing error messages with path
func (_gee *CT_Acc) ValidateWithPath(path string) error {
	if _gee.AccPr != nil {
		if _af := _gee.AccPr.ValidateWithPath(path + "\u002f\u0041\u0063\u0063\u0050\u0072"); _af != nil {
			return _af
		}
	}
	if _cc := _gee.E.ValidateWithPath(path + "\u002f\u0045"); _cc != nil {
		return _cc
	}
	return nil
}

// ValidateWithPath validates the CT_Func and its children, prefixing error messages with path
func (_ggg *CT_Func) ValidateWithPath(path string) error {
	if _ggg.FuncPr != nil {
		if _aagc := _ggg.FuncPr.ValidateWithPath(path + "\u002fF\u0075\u006e\u0063\u0050\u0072"); _aagc != nil {
			return _aagc
		}
	}
	if _ceg := _ggg.FName.ValidateWithPath(path + "\u002f\u0046\u004e\u0061\u006d\u0065"); _ceg != nil {
		return _ceg
	}
	if _cedc := _ggg.E.ValidateWithPath(path + "\u002f\u0045"); _cedc != nil {
		return _cedc
	}
	return nil
}

type CT_SSup struct {
	SSupPr *CT_SSupPr
	E      *CT_OMathArg
	Sup    *CT_OMathArg
}

func NewCT_Rad() *CT_Rad {
	_ebege := &CT_Rad{}
	_ebege.Deg = NewCT_OMathArg()
	_ebege.E = NewCT_OMathArg()
	return _ebege
}
func (_fcfg *CT_FType) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_dgga, _bbe := _fcfg.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _bbe != nil {
		return _bbe
	}
	start.Attr = append(start.Attr, _dgga)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_dbbbd *CT_SPrePr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gdfd:
	for {
		_agcee, _aedb := d.Token()
		if _aedb != nil {
			return _aedb
		}
		switch _gcgd := _agcee.(type) {
		case _g.StartElement:
			switch _gcgd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_dbbbd.CtrlPr = NewCT_CtrlPr()
				if _eddda := d.DecodeElement(_dbbbd.CtrlPr, &_gcgd); _eddda != nil {
					return _eddda
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0053\u0050\u0072e\u0050\u0072 \u0025\u0076", _gcgd.Name)
				if _eacb := d.Skip(); _eacb != nil {
					return _eacb
				}
			}
		case _g.EndElement:
			break _gdfd
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OnOff and its children
func (_aggd *CT_OnOff) Validate() error {
	return _aggd.ValidateWithPath("\u0043\u0054\u005f\u004f\u006e\u004f\u0066\u0066")
}

// ValidateWithPath validates the CT_RadPr and its children, prefixing error messages with path
func (_fgcc *CT_RadPr) ValidateWithPath(path string) error {
	if _fgcc.DegHide != nil {
		if _cfefa := _fgcc.DegHide.ValidateWithPath(path + "\u002f\u0044\u0065\u0067\u0048\u0069\u0064\u0065"); _cfefa != nil {
			return _cfefa
		}
	}
	if _fgcc.CtrlPr != nil {
		if _afge := _fgcc.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _afge != nil {
			return _afge
		}
	}
	return nil
}
func (_cebg *ST_LimLoc) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_cebg = 0
	case "\u0075\u006e\u0064\u004f\u0076\u0072":
		*_cebg = 1
	case "\u0073\u0075\u0062\u0053\u0075\u0070":
		*_cebg = 2
	}
	return nil
}

// Validate validates the CT_FType and its children
func (_dcd *CT_FType) Validate() error {
	return _dcd.ValidateWithPath("\u0043\u0054\u005f\u0046\u0054\u0079\u0070\u0065")
}

// Validate validates the OMathPara and its children
func (_fbffa *OMathPara) Validate() error {
	return _fbffa.ValidateWithPath("\u004fM\u0061\u0074\u0068\u0050\u0061\u0072a")
}
func NewCT_MR() *CT_MR { _abf := &CT_MR{}; return _abf }

// ValidateWithPath validates the CT_D and its children, prefixing error messages with path
func (_edf *CT_D) ValidateWithPath(path string) error {
	if _edf.DPr != nil {
		if _gedc := _edf.DPr.ValidateWithPath(path + "\u002f\u0044\u0050\u0072"); _gedc != nil {
			return _gedc
		}
	}
	for _gefb, _ffdg := range _edf.E {
		if _ddg := _ffdg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0045\u005b\u0025\u0064\u005d", path, _gefb)); _ddg != nil {
			return _ddg
		}
	}
	return nil
}
func (_gcce *CT_MathPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_dbcg:
	for {
		_fag, _bedg := d.Token()
		if _bedg != nil {
			return _bedg
		}
		switch _defc := _fag.(type) {
		case _g.StartElement:
			switch _defc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}:
				_gcce.MathFont = NewCT_String()
				if _afd := d.DecodeElement(_gcce.MathFont, &_defc); _afd != nil {
					return _afd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}:
				_gcce.BrkBin = NewCT_BreakBin()
				if _fbbg := d.DecodeElement(_gcce.BrkBin, &_defc); _fbbg != nil {
					return _fbbg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}:
				_gcce.BrkBinSub = NewCT_BreakBinSub()
				if _cdgg := d.DecodeElement(_gcce.BrkBinSub, &_defc); _cdgg != nil {
					return _cdgg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}:
				_gcce.SmallFrac = NewCT_OnOff()
				if _aedf := d.DecodeElement(_gcce.SmallFrac, &_defc); _aedf != nil {
					return _aedf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}:
				_gcce.DispDef = NewCT_OnOff()
				if _febb := d.DecodeElement(_gcce.DispDef, &_defc); _febb != nil {
					return _febb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}:
				_gcce.LMargin = NewCT_TwipsMeasure()
				if _gebd := d.DecodeElement(_gcce.LMargin, &_defc); _gebd != nil {
					return _gebd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}:
				_gcce.RMargin = NewCT_TwipsMeasure()
				if _bfa := d.DecodeElement(_gcce.RMargin, &_defc); _bfa != nil {
					return _bfa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}:
				_gcce.DefJc = NewCT_OMathJc()
				if _bacb := d.DecodeElement(_gcce.DefJc, &_defc); _bacb != nil {
					return _bacb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}:
				_gcce.PreSp = NewCT_TwipsMeasure()
				if _acdaf := d.DecodeElement(_gcce.PreSp, &_defc); _acdaf != nil {
					return _acdaf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}:
				_gcce.PostSp = NewCT_TwipsMeasure()
				if _gbfa := d.DecodeElement(_gcce.PostSp, &_defc); _gbfa != nil {
					return _gbfa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}:
				_gcce.InterSp = NewCT_TwipsMeasure()
				if _ddceb := d.DecodeElement(_gcce.InterSp, &_defc); _ddceb != nil {
					return _ddceb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}:
				_gcce.IntraSp = NewCT_TwipsMeasure()
				if _cgabe := d.DecodeElement(_gcce.IntraSp, &_defc); _cgabe != nil {
					return _cgabe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}:
				_gcce.Choice = NewCT_MathPrChoice()
				if _deg := d.DecodeElement(&_gcce.Choice.WrapIndent, &_defc); _deg != nil {
					return _deg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}:
				_gcce.Choice = NewCT_MathPrChoice()
				if _bee := d.DecodeElement(&_gcce.Choice.WrapRight, &_defc); _bee != nil {
					return _bee
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}:
				_gcce.IntLim = NewCT_LimLoc()
				if _egag := d.DecodeElement(_gcce.IntLim, &_defc); _egag != nil {
					return _egag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}:
				_gcce.NaryLim = NewCT_LimLoc()
				if _feec := d.DecodeElement(_gcce.NaryLim, &_defc); _feec != nil {
					return _feec
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004d\u0061\u0074h\u0050\u0072 \u0025\u0076", _defc.Name)
				if _egdb := d.Skip(); _egdb != nil {
					return _egdb
				}
			}
		case _g.EndElement:
			break _dbcg
		case _g.CharData:
		}
	}
	return nil
}
func (_caa *CT_MCPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_dfea:
	for {
		_gfad, _gace := d.Token()
		if _gace != nil {
			return _gace
		}
		switch _gcba := _gfad.(type) {
		case _g.StartElement:
			switch _gcba.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u006f\u0075n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u006f\u0075n\u0074"}:
				_caa.Count = NewCT_Integer255()
				if _gfcg := d.DecodeElement(_caa.Count, &_gcba); _gfcg != nil {
					return _gfcg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063\u004a\u0063"}:
				_caa.McJc = NewCT_XAlign()
				if _edga := d.DecodeElement(_caa.McJc, &_gcba); _edga != nil {
					return _edga
				}
			default:
				_gg.Log("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u004d\u0043\u0050\u0072\u0020\u0025\u0076", _gcba.Name)
				if _ebfc := d.Skip(); _ebfc != nil {
					return _ebfc
				}
			}
		case _g.EndElement:
			break _dfea
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_F() *CT_F {
	_cfa := &CT_F{}
	_cfa.Num = NewCT_OMathArg()
	_cfa.Den = NewCT_OMathArg()
	return _cfa
}
func (_fdcf *CT_PhantPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cbbd:
	for {
		_abab, _bbgcc := d.Token()
		if _bbgcc != nil {
			return _bbgcc
		}
		switch _edcd := _abab.(type) {
		case _g.StartElement:
			switch _edcd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0068\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0068\u006f\u0077"}:
				_fdcf.Show = NewCT_OnOff()
				if _fggc := d.DecodeElement(_fdcf.Show, &_edcd); _fggc != nil {
					return _fggc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u007ae\u0072\u006f\u0057\u0069\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u007ae\u0072\u006f\u0057\u0069\u0064"}:
				_fdcf.ZeroWid = NewCT_OnOff()
				if _ebff := d.DecodeElement(_fdcf.ZeroWid, &_edcd); _ebff != nil {
					return _ebff
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u007ae\u0072\u006f\u0041\u0073\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u007ae\u0072\u006f\u0041\u0073\u0063"}:
				_fdcf.ZeroAsc = NewCT_OnOff()
				if _eeaf := d.DecodeElement(_fdcf.ZeroAsc, &_edcd); _eeaf != nil {
					return _eeaf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u007a\u0065\u0072\u006f\u0044\u0065\u0073\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u007a\u0065\u0072\u006f\u0044\u0065\u0073\u0063"}:
				_fdcf.ZeroDesc = NewCT_OnOff()
				if _dcfe := d.DecodeElement(_fdcf.ZeroDesc, &_edcd); _dcfe != nil {
					return _dcfe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074\u0072\u0061\u006e\u0073\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074\u0072\u0061\u006e\u0073\u0070"}:
				_fdcf.Transp = NewCT_OnOff()
				if _fddd := d.DecodeElement(_fdcf.Transp, &_edcd); _fddd != nil {
					return _fddd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fdcf.CtrlPr = NewCT_CtrlPr()
				if _fgag := d.DecodeElement(_fdcf.CtrlPr, &_edcd); _fgag != nil {
					return _fgag
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0068\u0061\u006e\u0074\u0050\u0072\u0020\u0025\u0076", _edcd.Name)
				if _fdbb := d.Skip(); _fdbb != nil {
					return _fdbb
				}
			}
		case _g.EndElement:
			break _cbbd
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_M() *CT_M { _bebd := &CT_M{}; return _bebd }

type OMathPara struct{ CT_OMathPara }

func NewCT_Bar() *CT_Bar             { _bf := &CT_Bar{}; _bf.E = NewCT_OMathArg(); return _bf }
func NewCT_RPRChoice() *CT_RPRChoice { _ddeg := &CT_RPRChoice{}; return _ddeg }
func (_cagea *CT_SPre) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cagea.Sub = NewCT_OMathArg()
	_cagea.Sup = NewCT_OMathArg()
	_cagea.E = NewCT_OMathArg()
_ceec:
	for {
		_adcda, _afb := d.Token()
		if _afb != nil {
			return _afb
		}
		switch _bced := _adcda.(type) {
		case _g.StartElement:
			switch _bced.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065\u0050\u0072"}:
				_cagea.SPrePr = NewCT_SPrePr()
				if _cega := d.DecodeElement(_cagea.SPrePr, &_bced); _cega != nil {
					return _cega
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _febe := d.DecodeElement(_cagea.Sub, &_bced); _febe != nil {
					return _febe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _fbef := d.DecodeElement(_cagea.Sup, &_bced); _fbef != nil {
					return _fbef
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _gfbgg := d.DecodeElement(_cagea.E, &_bced); _gfbgg != nil {
					return _gfbgg
				}
			default:
				_gg.Log("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0053\u0050\u0072\u0065\u0020\u0025\u0076", _bced.Name)
				if _eece := d.Skip(); _eece != nil {
					return _eece
				}
			}
		case _g.EndElement:
			break _ceec
		case _g.CharData:
		}
	}
	return nil
}
func (_ccb *CT_SSubPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ccb.CtrlPr != nil {
		_cdgaf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ccb.CtrlPr, _cdgaf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_NaryPr and its children
func (_fadg *CT_NaryPr) Validate() error {
	return _fadg.ValidateWithPath("\u0043T\u005f\u004e\u0061\u0072\u0079\u0050r")
}

type CT_MathPr struct {
	MathFont  *CT_String
	BrkBin    *CT_BreakBin
	BrkBinSub *CT_BreakBinSub
	SmallFrac *CT_OnOff
	DispDef   *CT_OnOff
	LMargin   *CT_TwipsMeasure
	RMargin   *CT_TwipsMeasure
	DefJc     *CT_OMathJc
	PreSp     *CT_TwipsMeasure
	PostSp    *CT_TwipsMeasure
	InterSp   *CT_TwipsMeasure
	IntraSp   *CT_TwipsMeasure
	Choice    *CT_MathPrChoice
	IntLim    *CT_LimLoc
	NaryLim   *CT_LimLoc
}
type CT_Text struct {
	SpaceAttr *string
	Content   string
}

func (_ga *CT_AccPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_geb:
	for {
		_ca, _gbb := d.Token()
		if _gbb != nil {
			return _gbb
		}
		switch _gcc := _ca.(type) {
		case _g.StartElement:
			switch _gcc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0068\u0072"}:
				_ga.Chr = NewCT_Char()
				if _eg := d.DecodeElement(_ga.Chr, &_gcc); _eg != nil {
					return _eg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ga.CtrlPr = NewCT_CtrlPr()
				if _fd := d.DecodeElement(_ga.CtrlPr, &_gcc); _fd != nil {
					return _fd
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0063\u0063\u0050\u0072\u0020\u0025\u0076", _gcc.Name)
				if _bb := d.Skip(); _bb != nil {
					return _bb
				}
			}
		case _g.EndElement:
			break _geb
		case _g.CharData:
		}
	}
	return nil
}

type CT_OMathJc struct{ ValAttr ST_Jc }

// Validate validates the CT_Bar and its children
func (_gd *CT_Bar) Validate() error {
	return _gd.ValidateWithPath("\u0043\u0054\u005f\u0042\u0061\u0072")
}

type CT_PhantPr struct {
	Show     *CT_OnOff
	ZeroWid  *CT_OnOff
	ZeroAsc  *CT_OnOff
	ZeroDesc *CT_OnOff
	Transp   *CT_OnOff
	CtrlPr   *CT_CtrlPr
}

func (_bcde *CT_ManualBreak) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _bcde.AlnAtAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d:\u0061\u006c\u006e\u0041\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_bcde.AlnAtAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_Acc() *CT_Acc { _ge := &CT_Acc{}; _ge.E = NewCT_OMathArg(); return _ge }
func (_bbce *CT_Rad) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bbce.Deg = NewCT_OMathArg()
	_bbce.E = NewCT_OMathArg()
_cdag:
	for {
		_bdcag, _cedeg := d.Token()
		if _cedeg != nil {
			return _cedeg
		}
		switch _gafd := _bdcag.(type) {
		case _g.StartElement:
			switch _gafd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064P\u0072"}:
				_bbce.RadPr = NewCT_RadPr()
				if _fgga := d.DecodeElement(_bbce.RadPr, &_gafd); _fgga != nil {
					return _fgga
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u0067"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u0067"}:
				if _eebe := d.DecodeElement(_bbce.Deg, &_gafd); _eebe != nil {
					return _eebe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _cage := d.DecodeElement(_bbce.E, &_gafd); _cage != nil {
					return _cage
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0061\u0064\u0020\u0025\u0076", _gafd.Name)
				if _dfcb := d.Skip(); _dfcb != nil {
					return _dfcb
				}
			}
		case _g.EndElement:
			break _cdag
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_Style() *CT_Style { _aeefc := &CT_Style{}; return _aeefc }
func (_abfa *CT_RPR) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_dddg:
	for {
		_cgfb, _bdggb := d.Token()
		if _bdggb != nil {
			return _bdggb
		}
		switch _dbf := _cgfb.(type) {
		case _g.StartElement:
			switch _dbf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u0074"}:
				_abfa.Lit = NewCT_OnOff()
				if _adcd := d.DecodeElement(_abfa.Lit, &_dbf); _adcd != nil {
					return _adcd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u006f\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u006f\u0072"}:
				_abfa.Choice = NewCT_RPRChoice()
				if _beac := d.DecodeElement(&_abfa.Choice.Nor, &_dbf); _beac != nil {
					return _beac
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b"}:
				_abfa.Brk = NewCT_ManualBreak()
				if _gcbag := d.DecodeElement(_abfa.Brk, &_dbf); _gcbag != nil {
					return _gcbag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u006c\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u006c\u006e"}:
				_abfa.Aln = NewCT_OnOff()
				if _fcfb := d.DecodeElement(_abfa.Aln, &_dbf); _fcfb != nil {
					return _fcfb
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0050\u0052\u0020\u0025\u0076", _dbf.Name)
				if _gdbd := d.Skip(); _gdbd != nil {
					return _gdbd
				}
			}
		case _g.EndElement:
			break _dddg
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Box and its children
func (_aef *CT_Box) Validate() error {
	return _aef.ValidateWithPath("\u0043\u0054\u005f\u0042\u006f\u0078")
}

// Validate validates the CT_BoxPr and its children
func (_gccg *CT_BoxPr) Validate() error {
	return _gccg.ValidateWithPath("\u0043\u0054\u005f\u0042\u006f\u0078\u0050\u0072")
}
func NewOMathPara() *OMathPara {
	_dfaa := &OMathPara{}
	_dfaa.CT_OMathPara = *NewCT_OMathPara()
	return _dfaa
}
func (_eac *CT_GroupChr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_eac.E = NewCT_OMathArg()
_cdae:
	for {
		_bbbf, _bfg := d.Token()
		if _bfg != nil {
			return _bfg
		}
		switch _bda := _bbbf.(type) {
		case _g.StartElement:
			switch _bda.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"}:
				_eac.GroupChrPr = NewCT_GroupChrPr()
				if _cag := d.DecodeElement(_eac.GroupChrPr, &_bda); _cag != nil {
					return _cag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _dfbc := d.DecodeElement(_eac.E, &_bda); _dfbc != nil {
					return _dfbc
				}
			default:
				_gg.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0020\u0025\u0076", _bda.Name)
				if _cab := d.Skip(); _cab != nil {
					return _cab
				}
			}
		case _g.EndElement:
			break _cdae
		case _g.CharData:
		}
	}
	return nil
}

type CT_OMathArg struct {
	ArgPr                *CT_OMathArgPr
	EG_OMathMathElements []*EG_OMathMathElements
	CtrlPr               *CT_CtrlPr
}
type ST_Style byte

func (_cffe *CT_OMathPara) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cffe.OMathParaPr != nil {
		_ggfa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006f\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}}
		e.EncodeElement(_cffe.OMathParaPr, _ggfa)
	}
	_cca := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006f\u004d\u0061\u0074\u0068"}}
	for _, _befa := range _cffe.OMath {
		e.EncodeElement(_befa, _cca)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_Integer255() *CT_Integer255 { _cba := &CT_Integer255{}; _cba.ValAttr = 1; return _cba }

// Validate validates the CT_OMathPara and its children
func (_fdad *CT_OMathPara) Validate() error {
	return _fdad.ValidateWithPath("\u0043\u0054\u005fO\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061")
}
func NewCT_OMathParaPr() *CT_OMathParaPr { _dddb := &CT_OMathParaPr{}; return _dddb }
func (_gacag *CT_Style) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _gacag.ValAttr != ST_StyleUnset {
		_egfcc, _ggcag := _gacag.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _ggcag != nil {
			return _ggcag
		}
		start.Attr = append(start.Attr, _egfcc)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_bbaee *ST_Style) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dgcg, _cdcgf := d.Token()
	if _cdcgf != nil {
		return _cdcgf
	}
	if _aabeg, _cabd := _dgcg.(_g.EndElement); _cabd && _aabeg.Name == start.Name {
		*_bbaee = 1
		return nil
	}
	if _cfce, _gbad := _dgcg.(_g.CharData); !_gbad {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dgcg)
	} else {
		switch string(_cfce) {
		case "":
			*_bbaee = 0
		case "\u0070":
			*_bbaee = 1
		case "\u0062":
			*_bbaee = 2
		case "\u0069":
			*_bbaee = 3
		case "\u0062\u0069":
			*_bbaee = 4
		}
	}
	_dgcg, _cdcgf = d.Token()
	if _cdcgf != nil {
		return _cdcgf
	}
	if _acde, _dbfb := _dgcg.(_g.EndElement); _dbfb && _acde.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dgcg)
}

// Validate validates the CT_LimLoc and its children
func (_ffdge *CT_LimLoc) Validate() error {
	return _ffdge.ValidateWithPath("\u0043T\u005f\u004c\u0069\u006d\u004c\u006fc")
}

// Validate validates the CT_YAlign and its children
func (_bdaa *CT_YAlign) Validate() error {
	return _bdaa.ValidateWithPath("\u0043T\u005f\u0059\u0041\u006c\u0069\u0067n")
}
func NewCT_OMathJc() *CT_OMathJc { _cdaeb := &CT_OMathJc{}; return _cdaeb }

// Validate validates the CT_RPR and its children
func (_egcef *CT_RPR) Validate() error {
	return _egcef.ValidateWithPath("\u0043\u0054\u005f\u0052\u0050\u0052")
}

// Validate validates the CT_MathPr and its children
func (_cadf *CT_MathPr) Validate() error {
	return _cadf.ValidateWithPath("\u0043T\u005f\u004d\u0061\u0074\u0068\u0050r")
}
func (_babc *ST_Script) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gbga, _dbfe := d.Token()
	if _dbfe != nil {
		return _dbfe
	}
	if _adefd, _gcge := _gbga.(_g.EndElement); _gcge && _adefd.Name == start.Name {
		*_babc = 1
		return nil
	}
	if _bbdc, _fecfg := _gbga.(_g.CharData); !_fecfg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbga)
	} else {
		switch string(_bbdc) {
		case "":
			*_babc = 0
		case "\u0072\u006f\u006da\u006e":
			*_babc = 1
		case "\u0073\u0063\u0072\u0069\u0070\u0074":
			*_babc = 2
		case "\u0066r\u0061\u006b\u0074\u0075\u0072":
			*_babc = 3
		case "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b":
			*_babc = 4
		case "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066":
			*_babc = 5
		case "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e":
			*_babc = 6
		}
	}
	_gbga, _dbfe = d.Token()
	if _dbfe != nil {
		return _dbfe
	}
	if _cbcf, _abedf := _gbga.(_g.EndElement); _abedf && _cbcf.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbga)
}
func (_fdagg ST_Style) ValidateWithPath(path string) error {
	switch _fdagg {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fdagg))
	}
	return nil
}

// Validate validates the CT_Style and its children
func (_dfeb *CT_Style) Validate() error {
	return _dfeb.ValidateWithPath("\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065")
}
func (_fegg *CT_LimLow) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fegg.LimLowPr != nil {
		_abaa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u004c\u006f\u0077\u0050\u0072"}}
		e.EncodeElement(_fegg.LimLowPr, _abaa)
	}
	_ebge := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_fegg.E, _ebge)
	_ccd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006ci\u006d"}}
	e.EncodeElement(_fegg.Lim, _ccd)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_afgg *ST_TopBot) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_abga, _afeb := d.Token()
	if _afeb != nil {
		return _afeb
	}
	if _aggge, _ebae := _abga.(_g.EndElement); _ebae && _aggge.Name == start.Name {
		*_afgg = 1
		return nil
	}
	if _dfeaf, _gcbe := _abga.(_g.CharData); !_gcbe {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _abga)
	} else {
		switch string(_dfeaf) {
		case "":
			*_afgg = 0
		case "\u0074\u006f\u0070":
			*_afgg = 1
		case "\u0062\u006f\u0074":
			*_afgg = 2
		}
	}
	_abga, _afeb = d.Token()
	if _afeb != nil {
		return _afeb
	}
	if _bgddg, _facff := _abga.(_g.EndElement); _facff && _bgddg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _abga)
}

// Validate validates the CT_Text and its children
func (_aedc *CT_Text) Validate() error {
	return _aedc.ValidateWithPath("\u0043T\u005f\u0054\u0065\u0078\u0074")
}
func (_cf *CT_BarPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ada:
	for {
		_gca, _baa := d.Token()
		if _baa != nil {
			return _baa
		}
		switch _cfc := _gca.(type) {
		case _g.StartElement:
			switch _cfc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073"}:
				_cf.Pos = NewCT_TopBot()
				if _db := d.DecodeElement(_cf.Pos, &_cfc); _db != nil {
					return _db
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_cf.CtrlPr = NewCT_CtrlPr()
				if _aae := d.DecodeElement(_cf.CtrlPr, &_cfc); _aae != nil {
					return _aae
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u0061\u0072\u0050\u0072\u0020\u0025\u0076", _cfc.Name)
				if _dba := d.Skip(); _dba != nil {
					return _dba
				}
			}
		case _g.EndElement:
			break _ada
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_Shp() *CT_Shp                   { _ebfab := &CT_Shp{}; _ebfab.ValAttr = ST_Shp(1); return _ebfab }
func (_ecgef ST_BreakBin) Validate() error { return _ecgef.ValidateWithPath("") }

type CT_LimLoc struct{ ValAttr ST_LimLoc }

func NewCT_SSup() *CT_SSup {
	_fada := &CT_SSup{}
	_fada.E = NewCT_OMathArg()
	_fada.Sup = NewCT_OMathArg()
	return _fada
}
func (_daab *CT_Shp) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_daab.ValAttr = ST_Shp(1)
	for _, _egff := range start.Attr {
		if _egff.Name.Local == "\u0076\u0061\u006c" {
			_daab.ValAttr.UnmarshalXMLAttr(_egff)
			continue
		}
	}
	for {
		_fgfa, _gfca := d.Token()
		if _gfca != nil {
			return _f.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0053h\u0070\u003a\u0020\u0025\u0073", _gfca)
		}
		if _eeag, _gadae := _fgfa.(_g.EndElement); _gadae && _eeag.Name == start.Name {
			break
		}
	}
	return nil
}
func (_bdcd *CT_SpacingRule) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _f.Sprintf("\u0025\u0076", _bdcd.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type ST_BreakBin byte
type CT_LimUppPr struct{ CtrlPr *CT_CtrlPr }

const (
	ST_ScriptUnset         ST_Script = 0
	ST_ScriptRoman         ST_Script = 1
	ST_ScriptScript        ST_Script = 2
	ST_ScriptFraktur       ST_Script = 3
	ST_ScriptDouble_struck ST_Script = 4
	ST_ScriptSans_serif    ST_Script = 5
	ST_ScriptMonospace     ST_Script = 6
)

type CT_EqArrPr struct {
	BaseJc  *CT_YAlign
	MaxDist *CT_OnOff
	ObjDist *CT_OnOff
	RSpRule *CT_SpacingRule
	RSp     *CT_UnSignedInteger
	CtrlPr  *CT_CtrlPr
}

// Validate validates the CT_FuncPr and its children
func (_gdd *CT_FuncPr) Validate() error {
	return _gdd.ValidateWithPath("\u0043T\u005f\u0046\u0075\u006e\u0063\u0050r")
}
func (_cfgd ST_TopBot) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_cfgd.String(), start)
}

// ValidateWithPath validates the CT_CtrlPr and its children, prefixing error messages with path
func (_fgc *CT_CtrlPr) ValidateWithPath(path string) error { return nil }
func (_fecc *CT_EqArrPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ffab:
	for {
		_cad, _ffgg := d.Token()
		if _ffgg != nil {
			return _ffgg
		}
		switch _dfg := _cad.(type) {
		case _g.StartElement:
			switch _dfg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}:
				_fecc.BaseJc = NewCT_YAlign()
				if _cbfa := d.DecodeElement(_fecc.BaseJc, &_dfg); _cbfa != nil {
					return _cbfa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006da\u0078\u0044\u0069\u0073\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006da\u0078\u0044\u0069\u0073\u0074"}:
				_fecc.MaxDist = NewCT_OnOff()
				if _bbaa := d.DecodeElement(_fecc.MaxDist, &_dfg); _bbaa != nil {
					return _bbaa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006fb\u006a\u0044\u0069\u0073\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006fb\u006a\u0044\u0069\u0073\u0074"}:
				_fecc.ObjDist = NewCT_OnOff()
				if _fcc := d.DecodeElement(_fecc.ObjDist, &_dfg); _fcc != nil {
					return _fcc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}:
				_fecc.RSpRule = NewCT_SpacingRule()
				if _cce := d.DecodeElement(_fecc.RSpRule, &_dfg); _cce != nil {
					return _cce
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0053\u0070"}:
				_fecc.RSp = NewCT_UnSignedInteger()
				if _gagc := d.DecodeElement(_fecc.RSp, &_dfg); _gagc != nil {
					return _gagc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fecc.CtrlPr = NewCT_CtrlPr()
				if _ebdf := d.DecodeElement(_fecc.CtrlPr, &_dfg); _ebdf != nil {
					return _ebdf
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fE\u0071\u0041\u0072\u0072\u0050\u0072\u0020\u0025\u0076", _dfg.Name)
				if _gffc := d.Skip(); _gffc != nil {
					return _gffc
				}
			}
		case _g.EndElement:
			break _ffab
		case _g.CharData:
		}
	}
	return nil
}

type CT_Box struct {
	BoxPr *CT_BoxPr
	E     *CT_OMathArg
}
type ST_BreakBinSub byte

// ValidateWithPath validates the CT_RPR and its children, prefixing error messages with path
func (_begc *CT_RPR) ValidateWithPath(path string) error {
	if _begc.Lit != nil {
		if _cafdf := _begc.Lit.ValidateWithPath(path + "\u002f\u004c\u0069\u0074"); _cafdf != nil {
			return _cafdf
		}
	}
	if _begc.Choice != nil {
		if _abfb := _begc.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _abfb != nil {
			return _abfb
		}
	}
	if _begc.Brk != nil {
		if _effe := _begc.Brk.ValidateWithPath(path + "\u002f\u0042\u0072\u006b"); _effe != nil {
			return _effe
		}
	}
	if _begc.Aln != nil {
		if _dbbb := _begc.Aln.ValidateWithPath(path + "\u002f\u0041\u006c\u006e"); _dbbb != nil {
			return _dbbb
		}
	}
	return nil
}
func NewCT_LimLow() *CT_LimLow {
	_cabb := &CT_LimLow{}
	_cabb.E = NewCT_OMathArg()
	_cabb.Lim = NewCT_OMathArg()
	return _cabb
}

// Validate validates the CT_SPre and its children
func (_dgbe *CT_SPre) Validate() error {
	return _dgbe.ValidateWithPath("\u0043T\u005f\u0053\u0050\u0072\u0065")
}
func (_ddcf *CT_MCS) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fcgc:
	for {
		_abgf, _bacc := d.Token()
		if _bacc != nil {
			return _bacc
		}
		switch _geed := _abgf.(type) {
		case _g.StartElement:
			switch _geed.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063"}:
				_cfdc := NewCT_MC()
				if _babd := d.DecodeElement(_cfdc, &_geed); _babd != nil {
					return _babd
				}
				_ddcf.Mc = append(_ddcf.Mc, _cfdc)
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004d\u0043\u0053\u0020\u0025\u0076", _geed.Name)
				if _geeb := d.Skip(); _geeb != nil {
					return _geeb
				}
			}
		case _g.EndElement:
			break _fcgc
		case _g.CharData:
		}
	}
	return nil
}
func (_fbgd *CT_Text) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _fbgd.SpaceAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u003a\u0073\u0070\u0061\u0063e"}, Value: _f.Sprintf("\u0025\u0076", *_fbgd.SpaceAttr)})
	}
	e.EncodeElement(_fbgd.Content, start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_daag ST_FType) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_daag.String(), start)
}

type ST_TopBot byte

func NewEG_OMathElements() *EG_OMathElements { _cfafa := &EG_OMathElements{}; return _cfafa }
func NewCT_SSupPr() *CT_SSupPr               { _gbcc := &CT_SSupPr{}; return _gbcc }
func (_cfef *CT_MPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cfef.BaseJc != nil {
		_bcgf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0061\u0073\u0065\u004a\u0063"}}
		e.EncodeElement(_cfef.BaseJc, _bcgf)
	}
	if _cfef.PlcHide != nil {
		_bbbb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u006c\u0063\u0048\u0069\u0064e"}}
		e.EncodeElement(_cfef.PlcHide, _bbbb)
	}
	if _cfef.RSpRule != nil {
		_aaa := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u0053\u0070\u0052\u0075\u006ce"}}
		e.EncodeElement(_cfef.RSpRule, _aaa)
	}
	if _cfef.CGpRule != nil {
		_gdb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0063\u0047\u0070\u0052\u0075\u006ce"}}
		e.EncodeElement(_cfef.CGpRule, _gdb)
	}
	if _cfef.RSp != nil {
		_affb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072S\u0070"}}
		e.EncodeElement(_cfef.RSp, _affb)
	}
	if _cfef.CSp != nil {
		_cdaf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063S\u0070"}}
		e.EncodeElement(_cfef.CSp, _cdaf)
	}
	if _cfef.CGp != nil {
		_bdab := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063G\u0070"}}
		e.EncodeElement(_cfef.CGp, _bdab)
	}
	if _cfef.Mcs != nil {
		_bgg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006dc\u0073"}}
		e.EncodeElement(_cfef.Mcs, _bgg)
	}
	if _cfef.CtrlPr != nil {
		_bed := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_cfef.CtrlPr, _bed)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type ST_Shp byte

func (_faaff ST_Style) Validate() error { return _faaff.ValidateWithPath("") }

// ValidateWithPath validates the CT_MathPrChoice and its children, prefixing error messages with path
func (_agfc *CT_MathPrChoice) ValidateWithPath(path string) error {
	if _agfc.WrapIndent != nil {
		if _gfcef := _agfc.WrapIndent.ValidateWithPath(path + "/\u0057\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"); _gfcef != nil {
			return _gfcef
		}
	}
	if _agfc.WrapRight != nil {
		if _fcge := _agfc.WrapRight.ValidateWithPath(path + "\u002f\u0057\u0072\u0061\u0070\u0052\u0069\u0067\u0068\u0074"); _fcge != nil {
			return _fcge
		}
	}
	return nil
}

// ValidateWithPath validates the CT_YAlign and its children, prefixing error messages with path
func (_ccecc *CT_YAlign) ValidateWithPath(path string) error {
	if _ccecc.ValAttr == _be.ST_YAlignUnset {
		return _f.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _abgd := _ccecc.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _abgd != nil {
		return _abgd
	}
	return nil
}

// Validate validates the CT_FPr and its children
func (_efbd *CT_FPr) Validate() error {
	return _efbd.ValidateWithPath("\u0043\u0054\u005f\u0046\u0050\u0072")
}

// Validate validates the CT_PhantPr and its children
func (_eceae *CT_PhantPr) Validate() error {
	return _eceae.ValidateWithPath("\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074\u0050\u0072")
}

type CT_TwipsMeasure struct{ ValAttr _be.ST_TwipsMeasure }

func (_adbc *ST_BreakBin) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_aabg, _bdadd := d.Token()
	if _bdadd != nil {
		return _bdadd
	}
	if _bgda, _bdce := _aabg.(_g.EndElement); _bdce && _bgda.Name == start.Name {
		*_adbc = 1
		return nil
	}
	if _gddb, _edae := _aabg.(_g.CharData); !_edae {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _aabg)
	} else {
		switch string(_gddb) {
		case "":
			*_adbc = 0
		case "\u0062\u0065\u0066\u006f\u0072\u0065":
			*_adbc = 1
		case "\u0061\u0066\u0074e\u0072":
			*_adbc = 2
		case "\u0072\u0065\u0070\u0065\u0061\u0074":
			*_adbc = 3
		}
	}
	_aabg, _bdadd = d.Token()
	if _bdadd != nil {
		return _bdadd
	}
	if _cfcaa, _bbfe := _aabg.(_g.EndElement); _bbfe && _cfcaa.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _aabg)
}
func (_efgg *CT_BreakBinSub) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _gccd := range start.Attr {
		if _gccd.Name.Local == "\u0076\u0061\u006c" {
			_efgg.ValAttr.UnmarshalXMLAttr(_gccd)
			continue
		}
	}
	for {
		_abg, _cece := d.Token()
		if _cece != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fB\u0072\u0065\u0061\u006b\u0042\u0069\u006e\u0053\u0075\u0062:\u0020\u0025\u0073", _cece)
		}
		if _fec, _cbcg := _abg.(_g.EndElement); _cbcg && _fec.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_String() *CT_String { _dbbba := &CT_String{}; return _dbbba }
func NewCT_Nary() *CT_Nary {
	_dbac := &CT_Nary{}
	_dbac.Sub = NewCT_OMathArg()
	_dbac.Sup = NewCT_OMathArg()
	_dbac.E = NewCT_OMathArg()
	return _dbac
}

type CT_SSupPr struct{ CtrlPr *CT_CtrlPr }

func (_ecfc *CT_MCPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ecfc.Count != nil {
		_fbda := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0063\u006f\u0075\u006e\u0074"}}
		e.EncodeElement(_ecfc.Count, _fbda)
	}
	if _ecfc.McJc != nil {
		_gagg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0063\u004a\u0063"}}
		e.EncodeElement(_ecfc.McJc, _gagg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_R and its children
func (_fccb *CT_R) Validate() error  { return _fccb.ValidateWithPath("\u0043\u0054\u005f\u0052") }
func NewCT_BorderBox() *CT_BorderBox { _ab := &CT_BorderBox{}; _ab.E = NewCT_OMathArg(); return _ab }

// ValidateWithPath validates the CT_SSub and its children, prefixing error messages with path
func (_bcab *CT_SSub) ValidateWithPath(path string) error {
	if _bcab.SSubPr != nil {
		if _gebea := _bcab.SSubPr.ValidateWithPath(path + "\u002fS\u0053\u0075\u0062\u0050\u0072"); _gebea != nil {
			return _gebea
		}
	}
	if _dagg := _bcab.E.ValidateWithPath(path + "\u002f\u0045"); _dagg != nil {
		return _dagg
	}
	if _badg := _bcab.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _badg != nil {
		return _badg
	}
	return nil
}

type CT_BoxPr struct {
	OpEmu   *CT_OnOff
	NoBreak *CT_OnOff
	Diff    *CT_OnOff
	Brk     *CT_ManualBreak
	Aln     *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

func (_eab *CT_OMath) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _eab.EG_OMathMathElements != nil {
		for _, _bcad := range _eab.EG_OMathMathElements {
			_bcad.MarshalXML(e, _g.StartElement{})
		}
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_fdcg *CT_RChoice) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gdceb:
	for {
		_cbfb, _edfbe := d.Token()
		if _edfbe != nil {
			return _edfbe
		}
		switch _edegg := _cbfb.(type) {
		case _g.StartElement:
			switch _edegg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074"}:
				_dafg := NewCT_Text()
				if _cbfbb := d.DecodeElement(_dafg, &_edegg); _cbfbb != nil {
					return _cbfbb
				}
				_fdcg.T = append(_fdcg.T, _dafg)
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fR\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _edegg.Name)
				if _bfcb := d.Skip(); _bfcb != nil {
					return _bfcb
				}
			}
		case _g.EndElement:
			break _gdceb
		case _g.CharData:
		}
	}
	return nil
}

type CT_FType struct{ ValAttr ST_FType }

func (_eeed *ST_Script) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_eeed = 0
	case "\u0072\u006f\u006da\u006e":
		*_eeed = 1
	case "\u0073\u0063\u0072\u0069\u0070\u0074":
		*_eeed = 2
	case "\u0066r\u0061\u006b\u0074\u0075\u0072":
		*_eeed = 3
	case "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b":
		*_eeed = 4
	case "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066":
		*_eeed = 5
	case "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e":
		*_eeed = 6
	}
	return nil
}

// ValidateWithPath validates the MathPr and its children, prefixing error messages with path
func (_gdffd *MathPr) ValidateWithPath(path string) error {
	if _dbdg := _gdffd.CT_MathPr.ValidateWithPath(path); _dbdg != nil {
		return _dbdg
	}
	return nil
}
func (_cbfc *CT_FType) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cbfc.ValAttr = ST_FType(1)
	for _, _edfe := range start.Attr {
		if _edfe.Name.Local == "\u0076\u0061\u006c" {
			_cbfc.ValAttr.UnmarshalXMLAttr(_edfe)
			continue
		}
	}
	for {
		_gcbc, _gbg := d.Token()
		if _gbg != nil {
			return _f.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fF\u0054\u0079\u0070\u0065: \u0025\u0073", _gbg)
		}
		if _aag, _ceca := _gcbc.(_g.EndElement); _ceca && _aag.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MCPr and its children
func (_fcce *CT_MCPr) Validate() error {
	return _fcce.ValidateWithPath("\u0043T\u005f\u004d\u0043\u0050\u0072")
}
func (_gddd ST_Jc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_gddd.String(), start)
}
func (_gfbba ST_BreakBin) ValidateWithPath(path string) error {
	switch _gfbba {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gfbba))
	}
	return nil
}

// Validate validates the CT_Acc and its children
func (_d *CT_Acc) Validate() error {
	return _d.ValidateWithPath("\u0043\u0054\u005f\u0041\u0063\u0063")
}

const (
	ST_LimLocUnset  ST_LimLoc = 0
	ST_LimLocUndOvr ST_LimLoc = 1
	ST_LimLocSubSup ST_LimLoc = 2
)
const (
	ST_FTypeUnset ST_FType = 0
	ST_FTypeBar   ST_FType = 1
	ST_FTypeSkw   ST_FType = 2
	ST_FTypeLin   ST_FType = 3
	ST_FTypeNoBar ST_FType = 4
)

func (_fcfa *CT_LimUpp) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fcfa.LimUppPr != nil {
		_abac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u0055\u0070\u0070\u0050\u0072"}}
		e.EncodeElement(_fcfa.LimUppPr, _abac)
	}
	_cabbb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_fcfa.E, _cabbb)
	_ecg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006ci\u006d"}}
	e.EncodeElement(_fcfa.Lim, _ecg)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_OnOff() *CT_OnOff { _bafd := &CT_OnOff{}; return _bafd }

// ValidateWithPath validates the CT_DPr and its children, prefixing error messages with path
func (_fdae *CT_DPr) ValidateWithPath(path string) error {
	if _fdae.BegChr != nil {
		if _afea := _fdae.BegChr.ValidateWithPath(path + "\u002fB\u0065\u0067\u0043\u0068\u0072"); _afea != nil {
			return _afea
		}
	}
	if _fdae.SepChr != nil {
		if _edd := _fdae.SepChr.ValidateWithPath(path + "\u002fS\u0065\u0070\u0043\u0068\u0072"); _edd != nil {
			return _edd
		}
	}
	if _fdae.EndChr != nil {
		if _egee := _fdae.EndChr.ValidateWithPath(path + "\u002fE\u006e\u0064\u0043\u0068\u0072"); _egee != nil {
			return _egee
		}
	}
	if _fdae.Grow != nil {
		if _eed := _fdae.Grow.ValidateWithPath(path + "\u002f\u0047\u0072o\u0077"); _eed != nil {
			return _eed
		}
	}
	if _fdae.Shp != nil {
		if _dgdc := _fdae.Shp.ValidateWithPath(path + "\u002f\u0053\u0068\u0070"); _dgdc != nil {
			return _dgdc
		}
	}
	if _fdae.CtrlPr != nil {
		if _aeca := _fdae.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _aeca != nil {
			return _aeca
		}
	}
	return nil
}

// Validate validates the MathPr and its children
func (_bbae *MathPr) Validate() error {
	return _bbae.ValidateWithPath("\u004d\u0061\u0074\u0068\u0050\u0072")
}

// ValidateWithPath validates the EG_OMathMathElements and its children, prefixing error messages with path
func (_ecbb *EG_OMathMathElements) ValidateWithPath(path string) error {
	if _ecbb.Acc != nil {
		if _gedae := _ecbb.Acc.ValidateWithPath(path + "\u002f\u0041\u0063\u0063"); _gedae != nil {
			return _gedae
		}
	}
	if _ecbb.Bar != nil {
		if _aaag := _ecbb.Bar.ValidateWithPath(path + "\u002f\u0042\u0061\u0072"); _aaag != nil {
			return _aaag
		}
	}
	if _ecbb.Box != nil {
		if _dcedd := _ecbb.Box.ValidateWithPath(path + "\u002f\u0042\u006f\u0078"); _dcedd != nil {
			return _dcedd
		}
	}
	if _ecbb.BorderBox != nil {
		if _gbdf := _ecbb.BorderBox.ValidateWithPath(path + "\u002f\u0042\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078"); _gbdf != nil {
			return _gbdf
		}
	}
	if _ecbb.D != nil {
		if _gaegc := _ecbb.D.ValidateWithPath(path + "\u002f\u0044"); _gaegc != nil {
			return _gaegc
		}
	}
	if _ecbb.EqArr != nil {
		if _dfad := _ecbb.EqArr.ValidateWithPath(path + "\u002f\u0045\u0071\u0041\u0072\u0072"); _dfad != nil {
			return _dfad
		}
	}
	if _ecbb.F != nil {
		if _cabbg := _ecbb.F.ValidateWithPath(path + "\u002f\u0046"); _cabbg != nil {
			return _cabbg
		}
	}
	if _ecbb.Func != nil {
		if _fdfb := _ecbb.Func.ValidateWithPath(path + "\u002f\u0046\u0075n\u0063"); _fdfb != nil {
			return _fdfb
		}
	}
	if _ecbb.GroupChr != nil {
		if _aged := _ecbb.GroupChr.ValidateWithPath(path + "\u002fG\u0072\u006f\u0075\u0070\u0043\u0068r"); _aged != nil {
			return _aged
		}
	}
	if _ecbb.LimLow != nil {
		if _geff := _ecbb.LimLow.ValidateWithPath(path + "\u002fL\u0069\u006d\u004c\u006f\u0077"); _geff != nil {
			return _geff
		}
	}
	if _ecbb.LimUpp != nil {
		if _cbeg := _ecbb.LimUpp.ValidateWithPath(path + "\u002fL\u0069\u006d\u0055\u0070\u0070"); _cbeg != nil {
			return _cbeg
		}
	}
	if _ecbb.M != nil {
		if _abgcf := _ecbb.M.ValidateWithPath(path + "\u002f\u004d"); _abgcf != nil {
			return _abgcf
		}
	}
	if _ecbb.Nary != nil {
		if _cgbcb := _ecbb.Nary.ValidateWithPath(path + "\u002f\u004e\u0061r\u0079"); _cgbcb != nil {
			return _cgbcb
		}
	}
	if _ecbb.Phant != nil {
		if _gedf := _ecbb.Phant.ValidateWithPath(path + "\u002f\u0050\u0068\u0061\u006e\u0074"); _gedf != nil {
			return _gedf
		}
	}
	if _ecbb.Rad != nil {
		if _abdf := _ecbb.Rad.ValidateWithPath(path + "\u002f\u0052\u0061\u0064"); _abdf != nil {
			return _abdf
		}
	}
	if _ecbb.SPre != nil {
		if _abcec := _ecbb.SPre.ValidateWithPath(path + "\u002f\u0053\u0050r\u0065"); _abcec != nil {
			return _abcec
		}
	}
	if _ecbb.SSub != nil {
		if _gcfb := _ecbb.SSub.ValidateWithPath(path + "\u002f\u0053\u0053u\u0062"); _gcfb != nil {
			return _gcfb
		}
	}
	if _ecbb.SSubSup != nil {
		if _ebecb := _ecbb.SSubSup.ValidateWithPath(path + "\u002f\u0053\u0053\u0075\u0062\u0053\u0075\u0070"); _ebecb != nil {
			return _ebecb
		}
	}
	if _ecbb.SSup != nil {
		if _cfbfd := _ecbb.SSup.ValidateWithPath(path + "\u002f\u0053\u0053u\u0070"); _cfbfd != nil {
			return _cfbfd
		}
	}
	if _ecbb.R != nil {
		if _feee := _ecbb.R.ValidateWithPath(path + "\u002f\u0052"); _feee != nil {
			return _feee
		}
	}
	return nil
}
func NewCT_EqArrPr() *CT_EqArrPr { _bcdd := &CT_EqArrPr{}; return _bcdd }

// ValidateWithPath validates the CT_UnSignedInteger and its children, prefixing error messages with path
func (_cccac *CT_UnSignedInteger) ValidateWithPath(path string) error { return nil }
func NewCT_Char() *CT_Char                                            { _aed := &CT_Char{}; return _aed }

// Validate validates the CT_SSubPr and its children
func (_bega *CT_SSubPr) Validate() error {
	return _bega.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0062\u0050r")
}
func (_dfggcf ST_LimLoc) Validate() error { return _dfggcf.ValidateWithPath("") }
func (_fbgaa *CT_RPRChoice) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _fbgaa.Nor != nil {
		_bbee := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006eo\u0072"}}
		e.EncodeElement(_fbgaa.Nor, _bbee)
	}
	return nil
}
func (_aggg *CT_TwipsMeasure) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _ecdff := range start.Attr {
		if _ecdff.Name.Local == "\u0076\u0061\u006c" {
			_bgdb, _agde := ParseUnionST_TwipsMeasure(_ecdff.Value)
			if _agde != nil {
				return _agde
			}
			_aggg.ValAttr = _bgdb
			continue
		}
	}
	for {
		_abee, _afdf := d.Token()
		if _afdf != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0054\u0077\u0069\u0070\u0073\u004de\u0061\u0073\u0075r\u0065:\u0020\u0025\u0073", _afdf)
		}
		if _faac, _fgcca := _abee.(_g.EndElement); _fgcca && _faac.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_Text() *CT_Text { _gbda := &CT_Text{}; return _gbda }

// ValidateWithPath validates the CT_Char and its children, prefixing error messages with path
func (_dbg *CT_Char) ValidateWithPath(path string) error { return nil }
func (_dcga *CT_Phant) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dcga.PhantPr != nil {
		_fbdac := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u0068\u0061\u006e\u0074\u0050r"}}
		e.EncodeElement(_dcga.PhantPr, _fbdac)
	}
	_ddcg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_dcga.E, _ddcg)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_Nary struct {
	NaryPr *CT_NaryPr
	Sub    *CT_OMathArg
	Sup    *CT_OMathArg
	E      *CT_OMathArg
}

// ValidateWithPath validates the CT_ManualBreak and its children, prefixing error messages with path
func (_efbge *CT_ManualBreak) ValidateWithPath(path string) error {
	if _efbge.AlnAtAttr != nil {
		if *_efbge.AlnAtAttr < 1 {
			return _f.Errorf("\u0025\u0073/\u006d\u002e\u0041\u006cn\u0041\u0074A\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u003e\u003d\u0020\u0031\u0020\u0028\u0068\u0061\u0076e\u0020\u0025\u0076\u0029", path, *_efbge.AlnAtAttr)
		}
		if *_efbge.AlnAtAttr > 255 {
			return _f.Errorf("\u0025\u0073\u002f\u006d\u002e\u0041\u006c\u006e\u0041\u0074\u0041\u0074\u0074r\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u003c\u003d\u0020\u0032\u0035\u0035\u0020\u0028\u0068\u0061\u0076e\u0020\u0025\u0076\u0029", path, *_efbge.AlnAtAttr)
		}
	}
	return nil
}
func (_geg *CT_RPRChoice) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ggcg:
	for {
		_bceg, _cbaf := d.Token()
		if _cbaf != nil {
			return _cbaf
		}
		switch _ebfa := _bceg.(type) {
		case _g.StartElement:
			switch _ebfa.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u006f\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u006f\u0072"}:
				_geg.Nor = NewCT_OnOff()
				if _degd := d.DecodeElement(_geg.Nor, &_ebfa); _degd != nil {
					return _degd
				}
			default:
				_gg.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_R\u0050\u0052C\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _ebfa.Name)
				if _cfdb := d.Skip(); _cfdb != nil {
					return _cfdb
				}
			}
		case _g.EndElement:
			break _ggcg
		case _g.CharData:
		}
	}
	return nil
}
func (_edec *CT_OMathParaPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fafc:
	for {
		_ceag, _agfe := d.Token()
		if _agfe != nil {
			return _agfe
		}
		switch _dfdfd := _ceag.(type) {
		case _g.StartElement:
			switch _dfdfd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006a\u0063"}:
				_edec.Jc = NewCT_OMathJc()
				if _efbde := d.DecodeElement(_edec.Jc, &_dfdfd); _efbde != nil {
					return _efbde
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0050\u0061r\u0061P\u0072\u0020\u0025\u0076", _dfdfd.Name)
				if _dede := d.Skip(); _dede != nil {
					return _dede
				}
			}
		case _g.EndElement:
			break _fafc
		case _g.CharData:
		}
	}
	return nil
}
func (_dfag *OMath) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006d"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0077"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065s\u0073i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u00306\u002fm\u0061\u0069n"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006d:\u006f\u004d\u0061\u0074\u0068"
	return _dfag.CT_OMath.MarshalXML(e, start)
}

// ValidateWithPath validates the CT_SSubSupPr and its children, prefixing error messages with path
func (_gcbde *CT_SSubSupPr) ValidateWithPath(path string) error {
	if _gcbde.AlnScr != nil {
		if _bfff := _gcbde.AlnScr.ValidateWithPath(path + "\u002fA\u006c\u006e\u0053\u0063\u0072"); _bfff != nil {
			return _bfff
		}
	}
	if _gcbde.CtrlPr != nil {
		if _dffab := _gcbde.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _dffab != nil {
			return _dffab
		}
	}
	return nil
}
func (_ddag *EG_OMathMathElements) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _ddag.Acc != nil {
		_bfegd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061c\u0063"}}
		e.EncodeElement(_ddag.Acc, _bfegd)
	}
	if _ddag.Bar != nil {
		_bdfde := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062a\u0072"}}
		e.EncodeElement(_ddag.Bar, _bdfde)
	}
	if _ddag.Box != nil {
		_bgdd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062o\u0078"}}
		e.EncodeElement(_ddag.Box, _bgdd)
	}
	if _ddag.BorderBox != nil {
		_feecc := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0062\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078"}}
		e.EncodeElement(_ddag.BorderBox, _feecc)
	}
	if _ddag.D != nil {
		_bcdcb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064"}}
		e.EncodeElement(_ddag.D, _bcdcb)
	}
	if _ddag.EqArr != nil {
		_bede := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0065\u0071\u0041\u0072\u0072"}}
		e.EncodeElement(_ddag.EqArr, _bede)
	}
	if _ddag.F != nil {
		_cfac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066"}}
		e.EncodeElement(_ddag.F, _cfac)
	}
	if _ddag.Func != nil {
		_bagd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066\u0075\u006e\u0063"}}
		e.EncodeElement(_ddag.Func, _bagd)
	}
	if _ddag.GroupChr != nil {
		_ffdfe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}}
		e.EncodeElement(_ddag.GroupChr, _ffdfe)
	}
	if _ddag.LimLow != nil {
		_gaae := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u004c\u006f\u0077"}}
		e.EncodeElement(_ddag.LimLow, _gaae)
	}
	if _ddag.LimUpp != nil {
		_dddee := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u0055\u0070\u0070"}}
		e.EncodeElement(_ddag.LimUpp, _dddee)
	}
	if _ddag.M != nil {
		_cfgcc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d"}}
		e.EncodeElement(_ddag.M, _cfgcc)
	}
	if _ddag.Nary != nil {
		_fgda := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006e\u0061\u0072\u0079"}}
		e.EncodeElement(_ddag.Nary, _fgda)
	}
	if _ddag.Phant != nil {
		_ebcb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u0068\u0061\u006e\u0074"}}
		e.EncodeElement(_ddag.Phant, _ebcb)
	}
	if _ddag.Rad != nil {
		_gdgg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072a\u0064"}}
		e.EncodeElement(_ddag.Rad, _gdgg)
	}
	if _ddag.SPre != nil {
		_acdba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0050\u0072\u0065"}}
		e.EncodeElement(_ddag.SPre, _acdba)
	}
	if _ddag.SSub != nil {
		_bdga := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0062"}}
		e.EncodeElement(_ddag.SSub, _bdga)
	}
	if _ddag.SSubSup != nil {
		_cagf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0053\u0075\u0062\u0053\u0075p"}}
		e.EncodeElement(_ddag.SSubSup, _cagf)
	}
	if _ddag.SSup != nil {
		_caab := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0070"}}
		e.EncodeElement(_ddag.SSup, _caab)
	}
	if _ddag.R != nil {
		_ffdcf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072"}}
		e.EncodeElement(_ddag.R, _ffdcf)
	}
	return nil
}
func (_ccgd *CT_LimLowPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ccgd.CtrlPr != nil {
		_fdbf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ccgd.CtrlPr, _fdbf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_bfba *CT_OMathJc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _bfba.ValAttr != ST_JcUnset {
		_eeab, _accg := _bfba.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _accg != nil {
			return _accg
		}
		start.Attr = append(start.Attr, _eeab)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

const (
	ST_StyleUnset ST_Style = 0
	ST_StyleP     ST_Style = 1
	ST_StyleB     ST_Style = 2
	ST_StyleI     ST_Style = 3
	ST_StyleBi    ST_Style = 4
)

// Validate validates the CT_Integer255 and its children
func (_bcc *CT_Integer255) Validate() error {
	return _bcc.ValidateWithPath("\u0043\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032\u0035\u0035")
}
func (_fgbfg ST_Script) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_agdd := _g.Attr{}
	_agdd.Name = name
	switch _fgbfg {
	case ST_ScriptUnset:
		_agdd.Value = ""
	case ST_ScriptRoman:
		_agdd.Value = "\u0072\u006f\u006da\u006e"
	case ST_ScriptScript:
		_agdd.Value = "\u0073\u0063\u0072\u0069\u0070\u0074"
	case ST_ScriptFraktur:
		_agdd.Value = "\u0066r\u0061\u006b\u0074\u0075\u0072"
	case ST_ScriptDouble_struck:
		_agdd.Value = "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b"
	case ST_ScriptSans_serif:
		_agdd.Value = "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066"
	case ST_ScriptMonospace:
		_agdd.Value = "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e"
	}
	return _agdd, nil
}
func (_dege ST_Jc) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_febg := _g.Attr{}
	_febg.Name = name
	switch _dege {
	case ST_JcUnset:
		_febg.Value = ""
	case ST_JcLeft:
		_febg.Value = "\u006c\u0065\u0066\u0074"
	case ST_JcRight:
		_febg.Value = "\u0072\u0069\u0067h\u0074"
	case ST_JcCenter:
		_febg.Value = "\u0063\u0065\u006e\u0074\u0065\u0072"
	case ST_JcCenterGroup:
		_febg.Value = "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070"
	}
	return _febg, nil
}
func (_gc *CT_Acc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gc.AccPr != nil {
		_c := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0061\u0063\u0063\u0050\u0072"}}
		e.EncodeElement(_gc.AccPr, _c)
	}
	_gge := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_gc.E, _gge)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_Func struct {
	FuncPr *CT_FuncPr
	FName  *CT_OMathArg
	E      *CT_OMathArg
}

// ValidateWithPath validates the EG_ScriptStyle and its children, prefixing error messages with path
func (_bece *EG_ScriptStyle) ValidateWithPath(path string) error {
	if _bece.Scr != nil {
		if _dfcf := _bece.Scr.ValidateWithPath(path + "\u002f\u0053\u0063\u0072"); _dfcf != nil {
			return _dfcf
		}
	}
	if _bece.Sty != nil {
		if _ecbbe := _bece.Sty.ValidateWithPath(path + "\u002f\u0053\u0074\u0079"); _ecbbe != nil {
			return _ecbbe
		}
	}
	return nil
}
func NewCT_Integer2() *CT_Integer2 { _fcfc := &CT_Integer2{}; _fcfc.ValAttr = -2; return _fcfc }
func (_gcde *CT_GroupChr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gcde.GroupChrPr != nil {
		_bdfb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067r\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"}}
		e.EncodeElement(_gcde.GroupChrPr, _bdfb)
	}
	_eeda := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_gcde.E, _eeda)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_babad *ST_BreakBinSub) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_babad = 0
	case "\u002d\u002d":
		*_babad = 1
	case "\u002d\u002b":
		*_babad = 2
	case "\u002b\u002d":
		*_babad = 3
	}
	return nil
}
func (_adce *CT_M) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gcdd:
	for {
		_bgab, _gbbd := d.Token()
		if _gbbd != nil {
			return _gbbd
		}
		switch _bace := _bgab.(type) {
		case _g.StartElement:
			switch _bace.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0050\u0072"}:
				_adce.MPr = NewCT_MPr()
				if _dbdd := d.DecodeElement(_adce.MPr, &_bace); _dbdd != nil {
					return _dbdd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0072"}:
				_acff := NewCT_MR()
				if _fbdb := d.DecodeElement(_acff, &_bace); _fbdb != nil {
					return _fbdb
				}
				_adce.Mr = append(_adce.Mr, _acff)
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u004d\u0020\u0025\u0076", _bace.Name)
				if _bbbea := d.Skip(); _bbbea != nil {
					return _bbbea
				}
			}
		case _g.EndElement:
			break _gcdd
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSup and its children
func (_fdde *CT_SSup) Validate() error {
	return _fdde.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0070")
}
func (_gacad *MathPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006d"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0077"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065s\u0073i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u00306\u002fm\u0061\u0069n"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006d\u003a\u006d\u0061\u0074\u0068\u0050\u0072"
	return _gacad.CT_MathPr.MarshalXML(e, start)
}
func (_eaddd *OMath) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_eaddd.CT_OMath = *NewCT_OMath()
_agcfc:
	for {
		_affa, _gacgb := d.Token()
		if _gacgb != nil {
			return _gacgb
		}
		switch _eefc := _affa.(type) {
		case _g.StartElement:
			switch _eefc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_bgdc := NewEG_OMathMathElements()
				_bgdc.Acc = NewCT_Acc()
				if _dccff := d.DecodeElement(_bgdc.Acc, &_eefc); _dccff != nil {
					return _dccff
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _bgdc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_bffe := NewEG_OMathMathElements()
				_bffe.Bar = NewCT_Bar()
				if _dabf := d.DecodeElement(_bffe.Bar, &_eefc); _dabf != nil {
					return _dabf
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _bffe)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_edfec := NewEG_OMathMathElements()
				_edfec.Box = NewCT_Box()
				if _gefbc := d.DecodeElement(_edfec.Box, &_eefc); _gefbc != nil {
					return _gefbc
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _edfec)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_agge := NewEG_OMathMathElements()
				_agge.BorderBox = NewCT_BorderBox()
				if _fabb := d.DecodeElement(_agge.BorderBox, &_eefc); _fabb != nil {
					return _fabb
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _agge)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_faff := NewEG_OMathMathElements()
				_faff.D = NewCT_D()
				if _agda := d.DecodeElement(_faff.D, &_eefc); _agda != nil {
					return _agda
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _faff)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_dagf := NewEG_OMathMathElements()
				_dagf.EqArr = NewCT_EqArr()
				if _gbab := d.DecodeElement(_dagf.EqArr, &_eefc); _gbab != nil {
					return _gbab
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _dagf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_aecg := NewEG_OMathMathElements()
				_aecg.F = NewCT_F()
				if _fgee := d.DecodeElement(_aecg.F, &_eefc); _fgee != nil {
					return _fgee
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _aecg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_dbbg := NewEG_OMathMathElements()
				_dbbg.Func = NewCT_Func()
				if _eccb := d.DecodeElement(_dbbg.Func, &_eefc); _eccb != nil {
					return _eccb
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _dbbg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_dbgcc := NewEG_OMathMathElements()
				_dbgcc.GroupChr = NewCT_GroupChr()
				if _acbb := d.DecodeElement(_dbgcc.GroupChr, &_eefc); _acbb != nil {
					return _acbb
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _dbgcc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_ecdb := NewEG_OMathMathElements()
				_ecdb.LimLow = NewCT_LimLow()
				if _dfeaa := d.DecodeElement(_ecdb.LimLow, &_eefc); _dfeaa != nil {
					return _dfeaa
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _ecdb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_baae := NewEG_OMathMathElements()
				_baae.LimUpp = NewCT_LimUpp()
				if _fbff := d.DecodeElement(_baae.LimUpp, &_eefc); _fbff != nil {
					return _fbff
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _baae)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_gbce := NewEG_OMathMathElements()
				_gbce.M = NewCT_M()
				if _ceafe := d.DecodeElement(_gbce.M, &_eefc); _ceafe != nil {
					return _ceafe
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _gbce)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_acce := NewEG_OMathMathElements()
				_acce.Nary = NewCT_Nary()
				if _fdaad := d.DecodeElement(_acce.Nary, &_eefc); _fdaad != nil {
					return _fdaad
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _acce)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_fgagb := NewEG_OMathMathElements()
				_fgagb.Phant = NewCT_Phant()
				if _gdfeg := d.DecodeElement(_fgagb.Phant, &_eefc); _gdfeg != nil {
					return _gdfeg
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _fgagb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_fgeed := NewEG_OMathMathElements()
				_fgeed.Rad = NewCT_Rad()
				if _gdea := d.DecodeElement(_fgeed.Rad, &_eefc); _gdea != nil {
					return _gdea
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _fgeed)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_aabbd := NewEG_OMathMathElements()
				_aabbd.SPre = NewCT_SPre()
				if _cace := d.DecodeElement(_aabbd.SPre, &_eefc); _cace != nil {
					return _cace
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _aabbd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_fedgg := NewEG_OMathMathElements()
				_fedgg.SSub = NewCT_SSub()
				if _efaf := d.DecodeElement(_fedgg.SSub, &_eefc); _efaf != nil {
					return _efaf
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _fedgg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_ceaaf := NewEG_OMathMathElements()
				_ceaaf.SSubSup = NewCT_SSubSup()
				if _afde := d.DecodeElement(_ceaaf.SSubSup, &_eefc); _afde != nil {
					return _afde
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _ceaaf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_gcga := NewEG_OMathMathElements()
				_gcga.SSup = NewCT_SSup()
				if _eagf := d.DecodeElement(_gcga.SSup, &_eefc); _eagf != nil {
					return _eagf
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _gcga)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_gaffa := NewEG_OMathMathElements()
				_gaffa.R = NewCT_R()
				if _dabed := d.DecodeElement(_gaffa.R, &_eefc); _dabed != nil {
					return _dabed
				}
				_eaddd.EG_OMathMathElements = append(_eaddd.EG_OMathMathElements, _gaffa)
			default:
				_gg.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u004f\u004d\u0061t\u0068 \u0025\u0076", _eefc.Name)
				if _ccdc := d.Skip(); _ccdc != nil {
					return _ccdc
				}
			}
		case _g.EndElement:
			break _agcfc
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_GroupChr() *CT_GroupChr { _edeg := &CT_GroupChr{}; _edeg.E = NewCT_OMathArg(); return _edeg }
func (_dcef *ST_Jc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bfga, _gbeed := d.Token()
	if _gbeed != nil {
		return _gbeed
	}
	if _bdefe, _acbcg := _bfga.(_g.EndElement); _acbcg && _bdefe.Name == start.Name {
		*_dcef = 1
		return nil
	}
	if _cfbfa, _baccb := _bfga.(_g.CharData); !_baccb {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bfga)
	} else {
		switch string(_cfbfa) {
		case "":
			*_dcef = 0
		case "\u006c\u0065\u0066\u0074":
			*_dcef = 1
		case "\u0072\u0069\u0067h\u0074":
			*_dcef = 2
		case "\u0063\u0065\u006e\u0074\u0065\u0072":
			*_dcef = 3
		case "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070":
			*_dcef = 4
		}
	}
	_bfga, _gbeed = d.Token()
	if _gbeed != nil {
		return _gbeed
	}
	if _ccfbc, _gfaba := _bfga.(_g.EndElement); _gfaba && _ccfbc.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bfga)
}

type CT_MathPrChoice struct {
	WrapIndent *CT_TwipsMeasure
	WrapRight  *CT_OnOff
}

func (_afc *CT_EqArrPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _afc.BaseJc != nil {
		_ccf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0061\u0073\u0065\u004a\u0063"}}
		e.EncodeElement(_afc.BaseJc, _ccf)
	}
	if _afc.MaxDist != nil {
		_ggf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006d\u0061\u0078\u0044\u0069\u0073t"}}
		e.EncodeElement(_afc.MaxDist, _ggf)
	}
	if _afc.ObjDist != nil {
		_ggab := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006f\u0062\u006a\u0044\u0069\u0073t"}}
		e.EncodeElement(_afc.ObjDist, _ggab)
	}
	if _afc.RSpRule != nil {
		_feb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u0053\u0070\u0052\u0075\u006ce"}}
		e.EncodeElement(_afc.RSpRule, _feb)
	}
	if _afc.RSp != nil {
		_eba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072S\u0070"}}
		e.EncodeElement(_afc.RSp, _eba)
	}
	if _afc.CtrlPr != nil {
		_gag := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_afc.CtrlPr, _gag)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_gadd *CT_OMathArgPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fafd:
	for {
		_dfffb, _ggca := d.Token()
		if _ggca != nil {
			return _ggca
		}
		switch _fffbc := _dfffb.(type) {
		case _g.StartElement:
			switch _fffbc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0072\u0067S\u007a"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0072\u0067S\u007a"}:
				_gadd.ArgSz = NewCT_Integer2()
				if _bcbb := d.DecodeElement(_gadd.ArgSz, &_fffbc); _bcbb != nil {
					return _bcbb
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004fM\u0061\u0074\u0068\u0041\u0072\u0067\u0050\u0072 \u0025\u0076", _fffbc.Name)
				if _efga := d.Skip(); _efga != nil {
					return _efga
				}
			}
		case _g.EndElement:
			break _fafd
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_MC() *CT_MC { _adb := &CT_MC{}; return _adb }
func (_bfeg *CT_SSubSup) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _bfeg.SSubSupPr != nil {
		_dbcb := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0073\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072"}}
		e.EncodeElement(_bfeg.SSubSupPr, _dbcb)
	}
	_dacc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_bfeg.E, _dacc)
	_aeec := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_bfeg.Sub, _aeec)
	_cegdg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_bfeg.Sup, _cegdg)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_bga *CT_GroupChrPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _bga.Chr != nil {
		_egaa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063h\u0072"}}
		e.EncodeElement(_bga.Chr, _egaa)
	}
	if _bga.Pos != nil {
		_ddec := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0070o\u0073"}}
		e.EncodeElement(_bga.Pos, _ddec)
	}
	if _bga.VertJc != nil {
		_bfc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0076\u0065\u0072\u0074\u004a\u0063"}}
		e.EncodeElement(_bga.VertJc, _bfc)
	}
	if _bga.CtrlPr != nil {
		_gbf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_bga.CtrlPr, _gbf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_decgd *CT_SSup) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _decgd.SSupPr != nil {
		_bfdd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0070\u0050\u0072"}}
		e.EncodeElement(_decgd.SSupPr, _bfdd)
	}
	_dgc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_decgd.E, _dgc)
	_ggbd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_decgd.Sup, _ggbd)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_RChoice and its children, prefixing error messages with path
func (_cecb *CT_RChoice) ValidateWithPath(path string) error {
	for _cccd, _aace := range _cecb.T {
		if _fdef := _aace.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0054\u005b\u0025\u0064\u005d", path, _cccd)); _fdef != nil {
			return _fdef
		}
	}
	return nil
}

// Validate validates the CT_UnSignedInteger and its children
func (_bege *CT_UnSignedInteger) Validate() error {
	return _bege.ValidateWithPath("\u0043T\u005fU\u006e\u0053\u0069\u0067\u006ee\u0064\u0049n\u0074\u0065\u0067\u0065\u0072")
}

// Validate validates the CT_D and its children
func (_edaa *CT_D) Validate() error { return _edaa.ValidateWithPath("\u0043\u0054\u005f\u0044") }
func (_fagc *CT_SSubSupPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fagc.AlnScr != nil {
		_daba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061\u006c\u006e\u0053\u0063\u0072"}}
		e.EncodeElement(_fagc.AlnScr, _daba)
	}
	if _fagc.CtrlPr != nil {
		_cadc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_fagc.CtrlPr, _cadc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_edff *CT_TopBot) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_addf, _gefcd := _edff.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _gefcd != nil {
		return _gefcd
	}
	start.Attr = append(start.Attr, _addf)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_MathPrChoice and its children
func (_eggec *CT_MathPrChoice) Validate() error {
	return _eggec.ValidateWithPath("\u0043T\u005fM\u0061\u0074\u0068\u0050\u0072\u0043\u0068\u006f\u0069\u0063\u0065")
}
func (_effg *ST_BreakBin) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_effg = 0
	case "\u0062\u0065\u0066\u006f\u0072\u0065":
		*_effg = 1
	case "\u0061\u0066\u0074e\u0072":
		*_effg = 2
	case "\u0072\u0065\u0070\u0065\u0061\u0074":
		*_effg = 3
	}
	return nil
}
func (_dadb *CT_MathPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dadb.MathFont != nil {
		_fbe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}}
		e.EncodeElement(_dadb.MathFont, _fbe)
	}
	if _dadb.BrkBin != nil {
		_agfa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0072\u006b\u0042\u0069\u006e"}}
		e.EncodeElement(_dadb.BrkBin, _agfa)
	}
	if _dadb.BrkBinSub != nil {
		_dcaf := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0062\u0072\u006b\u0042\u0069\u006e\u0053\u0075\u0062"}}
		e.EncodeElement(_dadb.BrkBinSub, _dcaf)
	}
	if _dadb.SmallFrac != nil {
		_ffba := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0073\u006d\u0061\u006c\u006c\u0046\u0072\u0061\u0063"}}
		e.EncodeElement(_dadb.SmallFrac, _ffba)
	}
	if _dadb.DispDef != nil {
		_abcc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0064\u0069\u0073\u0070\u0044\u0065f"}}
		e.EncodeElement(_dadb.DispDef, _abcc)
	}
	if _dadb.LMargin != nil {
		_abd := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006c\u004d\u0061\u0072\u0067\u0069n"}}
		e.EncodeElement(_dadb.LMargin, _abd)
	}
	if _dadb.RMargin != nil {
		_dfed := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u004d\u0061\u0072\u0067\u0069n"}}
		e.EncodeElement(_dadb.RMargin, _dfed)
	}
	if _dadb.DefJc != nil {
		_aead := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0064\u0065\u0066\u004a\u0063"}}
		e.EncodeElement(_dadb.DefJc, _aead)
	}
	if _dadb.PreSp != nil {
		_caae := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u0072\u0065\u0053\u0070"}}
		e.EncodeElement(_dadb.PreSp, _caae)
	}
	if _dadb.PostSp != nil {
		_fge := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0070\u006f\u0073\u0074\u0053\u0070"}}
		e.EncodeElement(_dadb.PostSp, _fge)
	}
	if _dadb.InterSp != nil {
		_afgf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0069\u006e\u0074\u0065\u0072\u0053p"}}
		e.EncodeElement(_dadb.InterSp, _afgf)
	}
	if _dadb.IntraSp != nil {
		_edfb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0069\u006e\u0074\u0072\u0061\u0053p"}}
		e.EncodeElement(_dadb.IntraSp, _edfb)
	}
	if _dadb.Choice != nil {
		_dadb.Choice.MarshalXML(e, _g.StartElement{})
	}
	if _dadb.IntLim != nil {
		_bbec := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0069\u006e\u0074\u004c\u0069\u006d"}}
		e.EncodeElement(_dadb.IntLim, _bbec)
	}
	if _dadb.NaryLim != nil {
		_egdfc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006e\u0061\u0072\u0079\u004c\u0069m"}}
		e.EncodeElement(_dadb.NaryLim, _egdfc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_SSubPr struct{ CtrlPr *CT_CtrlPr }

// ValidateWithPath validates the CT_TwipsMeasure and its children, prefixing error messages with path
func (_dccfed *CT_TwipsMeasure) ValidateWithPath(path string) error {
	if _fbgda := _dccfed.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fbgda != nil {
		return _fbgda
	}
	return nil
}

type CT_UnSignedInteger struct{ ValAttr uint32 }

func (_fc *CT_Bar) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fc.E = NewCT_OMathArg()
_dcf:
	for {
		_ff, _fga := d.Token()
		if _fga != nil {
			return _fga
		}
		switch _fb := _ff.(type) {
		case _g.StartElement:
			switch _fb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072P\u0072"}:
				_fc.BarPr = NewCT_BarPr()
				if _gcf := d.DecodeElement(_fc.BarPr, &_fb); _gcf != nil {
					return _gcf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _aa := d.DecodeElement(_fc.E, &_fb); _aa != nil {
					return _aa
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u0061\u0072\u0020\u0025\u0076", _fb.Name)
				if _dg := d.Skip(); _dg != nil {
					return _dg
				}
			}
		case _g.EndElement:
			break _dcf
		case _g.CharData:
		}
	}
	return nil
}

type CT_GroupChrPr struct {
	Chr    *CT_Char
	Pos    *CT_TopBot
	VertJc *CT_TopBot
	CtrlPr *CT_CtrlPr
}

// ValidateWithPath validates the CT_OMathArg and its children, prefixing error messages with path
func (_addb *CT_OMathArg) ValidateWithPath(path string) error {
	if _addb.ArgPr != nil {
		if _abdca := _addb.ArgPr.ValidateWithPath(path + "\u002f\u0041\u0072\u0067\u0050\u0072"); _abdca != nil {
			return _abdca
		}
	}
	for _ceac, _fgdd := range _addb.EG_OMathMathElements {
		if _ecea := _fgdd.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0045\u0047\u005f\u004fM\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006ce\u006d\u0065\u006et\u0073[\u0025\u0064\u005d", path, _ceac)); _ecea != nil {
			return _ecea
		}
	}
	if _addb.CtrlPr != nil {
		if _gacg := _addb.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _gacg != nil {
			return _gacg
		}
	}
	return nil
}
func (_adgc ST_Style) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_cdgd := _g.Attr{}
	_cdgd.Name = name
	switch _adgc {
	case ST_StyleUnset:
		_cdgd.Value = ""
	case ST_StyleP:
		_cdgd.Value = "\u0070"
	case ST_StyleB:
		_cdgd.Value = "\u0062"
	case ST_StyleI:
		_cdgd.Value = "\u0069"
	case ST_StyleBi:
		_cdgd.Value = "\u0062\u0069"
	}
	return _cdgd, nil
}

// ValidateWithPath validates the CT_LimLoc and its children, prefixing error messages with path
func (_cgf *CT_LimLoc) ValidateWithPath(path string) error {
	if _cgf.ValAttr == ST_LimLocUnset {
		return _f.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _bacg := _cgf.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _bacg != nil {
		return _bacg
	}
	return nil
}
func (_ddcdg *CT_RPR) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ddcdg.Lit != nil {
		_degg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006ci\u0074"}}
		e.EncodeElement(_ddcdg.Lit, _degg)
	}
	if _ddcdg.Choice != nil {
		_ddcdg.Choice.MarshalXML(e, _g.StartElement{})
	}
	if _ddcdg.Brk != nil {
		_dfde := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062r\u006b"}}
		e.EncodeElement(_ddcdg.Brk, _dfde)
	}
	if _ddcdg.Aln != nil {
		_gcfg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061l\u006e"}}
		e.EncodeElement(_ddcdg.Aln, _gcfg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_cgeb *CT_LimLoc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cgeb.ValAttr = ST_LimLoc(1)
	for _, _gdc := range start.Attr {
		if _gdc.Name.Local == "\u0076\u0061\u006c" {
			_cgeb.ValAttr.UnmarshalXMLAttr(_gdc)
			continue
		}
	}
	for {
		_gaeg, _feac := d.Token()
		if _feac != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u004c\u0069\u006dL\u006f\u0063\u003a\u0020\u0025\u0073", _feac)
		}
		if _bbaag, _ceab := _gaeg.(_g.EndElement); _ceab && _bbaag.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_RadPr struct {
	DegHide *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

func (_eccg ST_BreakBinSub) String() string {
	switch _eccg {
	case 0:
		return ""
	case 1:
		return "\u002d\u002d"
	case 2:
		return "\u002d\u002b"
	case 3:
		return "\u002b\u002d"
	}
	return ""
}
func (_ggbg *CT_FuncPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cafd:
	for {
		_afee, _bcaf := d.Token()
		if _bcaf != nil {
			return _bcaf
		}
		switch _agb := _afee.(type) {
		case _g.StartElement:
			switch _agb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ggbg.CtrlPr = NewCT_CtrlPr()
				if _fbbc := d.DecodeElement(_ggbg.CtrlPr, &_agb); _fbbc != nil {
					return _fbbc
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0046\u0075\u006ec\u0050\u0072 \u0025\u0076", _agb.Name)
				if _egbd := d.Skip(); _egbd != nil {
					return _egbd
				}
			}
		case _g.EndElement:
			break _cafd
		case _g.CharData:
		}
	}
	return nil
}

type CT_RChoice struct{ T []*CT_Text }

func (_daed ST_Script) ValidateWithPath(path string) error {
	switch _daed {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_daed))
	}
	return nil
}

type ST_FType byte
type CT_Bar struct {
	BarPr *CT_BarPr
	E     *CT_OMathArg
}

func (_dgdg *CT_GroupChrPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_edfg:
	for {
		_befe, _dffb := d.Token()
		if _dffb != nil {
			return _dffb
		}
		switch _dea := _befe.(type) {
		case _g.StartElement:
			switch _dea.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0068\u0072"}:
				_dgdg.Chr = NewCT_Char()
				if _eged := d.DecodeElement(_dgdg.Chr, &_dea); _eged != nil {
					return _eged
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073"}:
				_dgdg.Pos = NewCT_TopBot()
				if _fbga := d.DecodeElement(_dgdg.Pos, &_dea); _fbga != nil {
					return _fbga
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0076\u0065\u0072\u0074\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0076\u0065\u0072\u0074\u004a\u0063"}:
				_dgdg.VertJc = NewCT_TopBot()
				if _fee := d.DecodeElement(_dgdg.VertJc, &_dea); _fee != nil {
					return _fee
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_dgdg.CtrlPr = NewCT_CtrlPr()
				if _dccf := d.DecodeElement(_dgdg.CtrlPr, &_dea); _dccf != nil {
					return _dccf
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047r\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072 \u0025\u0076", _dea.Name)
				if _dgdge := d.Skip(); _dgdge != nil {
					return _dgdge
				}
			}
		case _g.EndElement:
			break _edfg
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TopBot and its children
func (_dcbb *CT_TopBot) Validate() error {
	return _dcbb.ValidateWithPath("\u0043T\u005f\u0054\u006f\u0070\u0042\u006ft")
}
func (_cbdb *CT_R) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ffbf:
	for {
		_gabg, _accc := d.Token()
		if _accc != nil {
			return _accc
		}
		switch _bgga := _gabg.(type) {
		case _g.StartElement:
			switch _bgga.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0050\u0072"}:
				_cbdb.RPr = NewCT_RPR()
				if _ccgc := d.DecodeElement(_cbdb.RPr, &_bgga); _ccgc != nil {
					return _ccgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074"}:
				_ecef := NewCT_RChoice()
				if _gbac := d.DecodeElement(&_ecef.T, &_bgga); _gbac != nil {
					return _gbac
				}
				_cbdb.Choice = append(_cbdb.Choice, _ecef)
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0052\u0020\u0025\u0076", _bgga.Name)
				if _fagg := d.Skip(); _fagg != nil {
					return _fagg
				}
			}
		case _g.EndElement:
			break _ffbf
		case _g.CharData:
		}
	}
	return nil
}
func (_bdfg *CT_D) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fcfd:
	for {
		_acda, _bcbg := d.Token()
		if _bcbg != nil {
			return _bcbg
		}
		switch _dbae := _acda.(type) {
		case _g.StartElement:
			switch _dbae.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0050\u0072"}:
				_bdfg.DPr = NewCT_DPr()
				if _dfd := d.DecodeElement(_bdfg.DPr, &_dbae); _dfd != nil {
					return _dfd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				_add := NewCT_OMathArg()
				if _aba := d.DecodeElement(_add, &_dbae); _aba != nil {
					return _aba
				}
				_bdfg.E = append(_bdfg.E, _add)
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0044\u0020\u0025\u0076", _dbae.Name)
				if _ddce := d.Skip(); _ddce != nil {
					return _ddce
				}
			}
		case _g.EndElement:
			break _fcfd
		case _g.CharData:
		}
	}
	return nil
}

type CT_Phant struct {
	PhantPr *CT_PhantPr
	E       *CT_OMathArg
}

func NewCT_SSubPr() *CT_SSubPr { _bdad := &CT_SSubPr{}; return _bdad }
func NewCT_FType() *CT_FType   { _fde := &CT_FType{}; _fde.ValAttr = ST_FType(1); return _fde }
func (_ceaec *EG_ScriptStyle) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fega:
	for {
		_dgge, _dedb := d.Token()
		if _dedb != nil {
			return _dedb
		}
		switch _ggae := _dgge.(type) {
		case _g.StartElement:
			switch _ggae.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0063\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0063\u0072"}:
				_ceaec.Scr = NewCT_Script()
				if _bgcda := d.DecodeElement(_ceaec.Scr, &_ggae); _bgcda != nil {
					return _bgcda
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0074\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0074\u0079"}:
				_ceaec.Sty = NewCT_Style()
				if _gaed := d.DecodeElement(_ceaec.Sty, &_ggae); _gaed != nil {
					return _gaed
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u0053\u0063\u0072\u0069\u0070\u0074\u0053t\u0079l\u0065\u0020\u0025\u0076", _ggae.Name)
				if _badgg := d.Skip(); _badgg != nil {
					return _badgg
				}
			}
		case _g.EndElement:
			break _fega
		case _g.CharData:
		}
	}
	return nil
}
func NewCT_MCPr() *CT_MCPr { _acaa := &CT_MCPr{}; return _acaa }

// ValidateWithPath validates the CT_OMathParaPr and its children, prefixing error messages with path
func (_ebb *CT_OMathParaPr) ValidateWithPath(path string) error {
	if _ebb.Jc != nil {
		if _edgg := _ebb.Jc.ValidateWithPath(path + "\u002f\u004a\u0063"); _edgg != nil {
			return _edgg
		}
	}
	return nil
}

// Validate validates the CT_MR and its children
func (_bdg *CT_MR) Validate() error { return _bdg.ValidateWithPath("\u0043\u0054\u005fM\u0052") }
func (_ddfc *OMathPara) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ddfc.CT_OMathPara = *NewCT_OMathPara()
_gabb:
	for {
		_cfgcd, _fbbf := d.Token()
		if _fbbf != nil {
			return _fbbf
		}
		switch _egdg := _cfgcd.(type) {
		case _g.StartElement:
			switch _egdg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}:
				_ddfc.OMathParaPr = NewCT_OMathParaPr()
				if _ecdg := d.DecodeElement(_ddfc.OMathParaPr, &_egdg); _ecdg != nil {
					return _ecdg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}:
				_eecg := NewCT_OMath()
				if _bdcaf := d.DecodeElement(_eecg, &_egdg); _bdcaf != nil {
					return _bdcaf
				}
				_ddfc.OMath = append(_ddfc.OMath, _eecg)
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u004f\u004d\u0061\u0074\u0068\u0050a\u0072\u0061 \u0025\u0076", _egdg.Name)
				if _egecc := d.Skip(); _egecc != nil {
					return _egecc
				}
			}
		case _g.EndElement:
			break _gabb
		case _g.CharData:
		}
	}
	return nil
}
func NewEG_ScriptStyle() *EG_ScriptStyle { _bcgfa := &EG_ScriptStyle{}; return _bcgfa }
func (_afacg ST_Style) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_afacg.String(), start)
}
func ParseUnionST_TwipsMeasure(s string) (_be.ST_TwipsMeasure, error) {
	_cdd := _be.ST_TwipsMeasure{}
	if _be.ST_PositiveUniversalMeasurePatternRe.MatchString(s) {
		_cdd.ST_PositiveUniversalMeasure = &s
	} else {
		_cebc, _aega := _e.ParseFloat(s, 64)
		if _aega != nil {
			return _cdd, _f.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0025\u0073\u0020\u0061\u0073\u0020\u0075\u0069\u006e\u0074\u003a\u0020%\u0073", s, _aega)
		}
		_cdd.ST_UnsignedDecimalNumber = _gg.Uint64(uint64(_cebc))
	}
	return _cdd, nil
}

const (
	ST_JcUnset       ST_Jc = 0
	ST_JcLeft        ST_Jc = 1
	ST_JcRight       ST_Jc = 2
	ST_JcCenter      ST_Jc = 3
	ST_JcCenterGroup ST_Jc = 4
)

// ValidateWithPath validates the CT_F and its children, prefixing error messages with path
func (_dff *CT_F) ValidateWithPath(path string) error {
	if _dff.FPr != nil {
		if _fbb := _dff.FPr.ValidateWithPath(path + "\u002f\u0046\u0050\u0072"); _fbb != nil {
			return _fbb
		}
	}
	if _cde := _dff.Num.ValidateWithPath(path + "\u002f\u004e\u0075\u006d"); _cde != nil {
		return _cde
	}
	if _dgb := _dff.Den.ValidateWithPath(path + "\u002f\u0044\u0065\u006e"); _dgb != nil {
		return _dgb
	}
	return nil
}

type CT_Integer2 struct{ ValAttr int64 }

func NewCT_MCS() *CT_MCS { _adeeg := &CT_MCS{}; return _adeeg }
func (_edad *CT_YAlign) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_edad.ValAttr = _be.ST_YAlign(1)
	for _, _cebabd := range start.Attr {
		if _cebabd.Name.Local == "\u0076\u0061\u006c" {
			_edad.ValAttr.UnmarshalXMLAttr(_cebabd)
			continue
		}
	}
	for {
		_cbdg, _bfcg := d.Token()
		if _bfcg != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0059\u0041\u006ci\u0067\u006e\u003a\u0020\u0025\u0073", _bfcg)
		}
		if _acgab, _ceae := _cbdg.(_g.EndElement); _ceae && _acgab.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ce *CT_Acc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ce.E = NewCT_OMathArg()
_eb:
	for {
		_ceb, _gb := d.Token()
		if _gb != nil {
			return _gb
		}
		switch _a := _ceb.(type) {
		case _g.StartElement:
			switch _a.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063P\u0072"}:
				_ce.AccPr = NewCT_AccPr()
				if _gcb := d.DecodeElement(_ce.AccPr, &_a); _gcb != nil {
					return _gcb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _bc := d.DecodeElement(_ce.E, &_a); _bc != nil {
					return _bc
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0063\u0063\u0020\u0025\u0076", _a.Name)
				if _bcd := d.Skip(); _bcd != nil {
					return _bcd
				}
			}
		case _g.EndElement:
			break _eb
		case _g.CharData:
		}
	}
	return nil
}
func (_bcdg *CT_MPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_bcbf:
	for {
		_fgac, _ebadc := d.Token()
		if _ebadc != nil {
			return _ebadc
		}
		switch _eggb := _fgac.(type) {
		case _g.StartElement:
			switch _eggb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}:
				_bcdg.BaseJc = NewCT_YAlign()
				if _fgaca := d.DecodeElement(_bcdg.BaseJc, &_eggb); _fgaca != nil {
					return _fgaca
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070l\u0063\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070l\u0063\u0048\u0069\u0064\u0065"}:
				_bcdg.PlcHide = NewCT_OnOff()
				if _defe := d.DecodeElement(_bcdg.PlcHide, &_eggb); _defe != nil {
					return _defe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}:
				_bcdg.RSpRule = NewCT_SpacingRule()
				if _gefd := d.DecodeElement(_bcdg.RSpRule, &_eggb); _gefd != nil {
					return _gefd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063G\u0070\u0052\u0075\u006c\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063G\u0070\u0052\u0075\u006c\u0065"}:
				_bcdg.CGpRule = NewCT_SpacingRule()
				if _dcfc := d.DecodeElement(_bcdg.CGpRule, &_eggb); _dcfc != nil {
					return _dcfc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0053\u0070"}:
				_bcdg.RSp = NewCT_UnSignedInteger()
				if _dffd := d.DecodeElement(_bcdg.RSp, &_eggb); _dffd != nil {
					return _dffd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0053\u0070"}:
				_bcdg.CSp = NewCT_UnSignedInteger()
				if _fgca := d.DecodeElement(_bcdg.CSp, &_eggb); _fgca != nil {
					return _fgca
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0047\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0047\u0070"}:
				_bcdg.CGp = NewCT_UnSignedInteger()
				if _bfed := d.DecodeElement(_bcdg.CGp, &_eggb); _bfed != nil {
					return _bfed
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063\u0073"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063\u0073"}:
				_bcdg.Mcs = NewCT_MCS()
				if _gdac := d.DecodeElement(_bcdg.Mcs, &_eggb); _gdac != nil {
					return _gdac
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_bcdg.CtrlPr = NewCT_CtrlPr()
				if _cfefc := d.DecodeElement(_bcdg.CtrlPr, &_eggb); _cfefc != nil {
					return _cfefc
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004d\u0050\u0072\u0020\u0025\u0076", _eggb.Name)
				if _bbdg := d.Skip(); _bbdg != nil {
					return _bbdg
				}
			}
		case _g.EndElement:
			break _bcbf
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LimUppPr and its children
func (_eaab *CT_LimUppPr) Validate() error {
	return _eaab.ValidateWithPath("C\u0054\u005f\u004c\u0069\u006d\u0055\u0070\u0070\u0050\u0072")
}

type CT_YAlign struct{ ValAttr _be.ST_YAlign }

func (_fdeff ST_TopBot) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_cdee := _g.Attr{}
	_cdee.Name = name
	switch _fdeff {
	case ST_TopBotUnset:
		_cdee.Value = ""
	case ST_TopBotTop:
		_cdee.Value = "\u0074\u006f\u0070"
	case ST_TopBotBot:
		_cdee.Value = "\u0062\u006f\u0074"
	}
	return _cdee, nil
}

// ValidateWithPath validates the CT_SPrePr and its children, prefixing error messages with path
func (_befbe *CT_SPrePr) ValidateWithPath(path string) error {
	if _befbe.CtrlPr != nil {
		if _cdfg := _befbe.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _cdfg != nil {
			return _cdfg
		}
	}
	return nil
}
func (_ecd *CT_DPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ecd.BegChr != nil {
		_abag := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0065\u0067\u0043\u0068\u0072"}}
		e.EncodeElement(_ecd.BegChr, _abag)
	}
	if _ecd.SepChr != nil {
		_becd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0065\u0070\u0043\u0068\u0072"}}
		e.EncodeElement(_ecd.SepChr, _becd)
	}
	if _ecd.EndChr != nil {
		_ddbd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065\u006e\u0064\u0043\u0068\u0072"}}
		e.EncodeElement(_ecd.EndChr, _ddbd)
	}
	if _ecd.Grow != nil {
		_baf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067\u0072\u006f\u0077"}}
		e.EncodeElement(_ecd.Grow, _baf)
	}
	if _ecd.Shp != nil {
		_eff := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073h\u0070"}}
		e.EncodeElement(_ecd.Shp, _eff)
	}
	if _ecd.CtrlPr != nil {
		_bdee := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ecd.CtrlPr, _bdee)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_ggga *MathPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ggga.CT_MathPr = *NewCT_MathPr()
_ecfa:
	for {
		_ddbc, _beee := d.Token()
		if _beee != nil {
			return _beee
		}
		switch _ccaf := _ddbc.(type) {
		case _g.StartElement:
			switch _ccaf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}:
				_ggga.MathFont = NewCT_String()
				if _gcgb := d.DecodeElement(_ggga.MathFont, &_ccaf); _gcgb != nil {
					return _gcgb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}:
				_ggga.BrkBin = NewCT_BreakBin()
				if _aacee := d.DecodeElement(_ggga.BrkBin, &_ccaf); _aacee != nil {
					return _aacee
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}:
				_ggga.BrkBinSub = NewCT_BreakBinSub()
				if _aceg := d.DecodeElement(_ggga.BrkBinSub, &_ccaf); _aceg != nil {
					return _aceg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}:
				_ggga.SmallFrac = NewCT_OnOff()
				if _bdea := d.DecodeElement(_ggga.SmallFrac, &_ccaf); _bdea != nil {
					return _bdea
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}:
				_ggga.DispDef = NewCT_OnOff()
				if _bcef := d.DecodeElement(_ggga.DispDef, &_ccaf); _bcef != nil {
					return _bcef
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}:
				_ggga.LMargin = NewCT_TwipsMeasure()
				if _bdeaa := d.DecodeElement(_ggga.LMargin, &_ccaf); _bdeaa != nil {
					return _bdeaa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}:
				_ggga.RMargin = NewCT_TwipsMeasure()
				if _affbg := d.DecodeElement(_ggga.RMargin, &_ccaf); _affbg != nil {
					return _affbg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}:
				_ggga.DefJc = NewCT_OMathJc()
				if _aagcb := d.DecodeElement(_ggga.DefJc, &_ccaf); _aagcb != nil {
					return _aagcb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}:
				_ggga.PreSp = NewCT_TwipsMeasure()
				if _eaffc := d.DecodeElement(_ggga.PreSp, &_ccaf); _eaffc != nil {
					return _eaffc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}:
				_ggga.PostSp = NewCT_TwipsMeasure()
				if _adcg := d.DecodeElement(_ggga.PostSp, &_ccaf); _adcg != nil {
					return _adcg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}:
				_ggga.InterSp = NewCT_TwipsMeasure()
				if _cbffd := d.DecodeElement(_ggga.InterSp, &_ccaf); _cbffd != nil {
					return _cbffd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}:
				_ggga.IntraSp = NewCT_TwipsMeasure()
				if _abed := d.DecodeElement(_ggga.IntraSp, &_ccaf); _abed != nil {
					return _abed
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}:
				_ggga.Choice = NewCT_MathPrChoice()
				if _efad := d.DecodeElement(&_ggga.Choice.WrapIndent, &_ccaf); _efad != nil {
					return _efad
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}:
				_ggga.Choice = NewCT_MathPrChoice()
				if _egeb := d.DecodeElement(&_ggga.Choice.WrapRight, &_ccaf); _egeb != nil {
					return _egeb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}:
				_ggga.IntLim = NewCT_LimLoc()
				if _gdda := d.DecodeElement(_ggga.IntLim, &_ccaf); _gdda != nil {
					return _gdda
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}:
				_ggga.NaryLim = NewCT_LimLoc()
				if _cgdfc := d.DecodeElement(_ggga.NaryLim, &_ccaf); _cgdfc != nil {
					return _cgdfc
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u004d\u0061\u0074\u0068\u0050\u0072\u0020\u0025\u0076", _ccaf.Name)
				if _dgeac := d.Skip(); _dgeac != nil {
					return _dgeac
				}
			}
		case _g.EndElement:
			break _ecfa
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_String and its children, prefixing error messages with path
func (_bfeb *CT_String) ValidateWithPath(path string) error { return nil }
func (_dgf ST_FType) String() string {
	switch _dgf {
	case 0:
		return ""
	case 1:
		return "\u0062\u0061\u0072"
	case 2:
		return "\u0073\u006b\u0077"
	case 3:
		return "\u006c\u0069\u006e"
	case 4:
		return "\u006e\u006f\u0042a\u0072"
	}
	return ""
}

// ValidateWithPath validates the CT_SpacingRule and its children, prefixing error messages with path
func (_gfef *CT_SpacingRule) ValidateWithPath(path string) error {
	if _gfef.ValAttr < 0 {
		return _f.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _gfef.ValAttr)
	}
	if _gfef.ValAttr > 4 {
		return _f.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0034\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _gfef.ValAttr)
	}
	return nil
}
func (_ccag ST_Script) Validate() error { return _ccag.ValidateWithPath("") }
func NewCT_BreakBin() *CT_BreakBin      { _fgaa := &CT_BreakBin{}; return _fgaa }

type ST_LimLoc byte

// Validate validates the CT_LimLow and its children
func (_efff *CT_LimLow) Validate() error {
	return _efff.ValidateWithPath("\u0043T\u005f\u004c\u0069\u006d\u004c\u006fw")
}
func (_fdefb ST_Shp) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_fdefb.String(), start)
}

// ValidateWithPath validates the CT_AccPr and its children, prefixing error messages with path
func (_fg *CT_AccPr) ValidateWithPath(path string) error {
	if _fg.Chr != nil {
		if _ed := _fg.Chr.ValidateWithPath(path + "\u002f\u0043\u0068\u0072"); _ed != nil {
			return _ed
		}
	}
	if _fg.CtrlPr != nil {
		if _ebg := _fg.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _ebg != nil {
			return _ebg
		}
	}
	return nil
}
func (_cgdd ST_LimLoc) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_abff := _g.Attr{}
	_abff.Name = name
	switch _cgdd {
	case ST_LimLocUnset:
		_abff.Value = ""
	case ST_LimLocUndOvr:
		_abff.Value = "\u0075\u006e\u0064\u004f\u0076\u0072"
	case ST_LimLocSubSup:
		_abff.Value = "\u0073\u0075\u0062\u0053\u0075\u0070"
	}
	return _abff, nil
}

// Validate validates the CT_BorderBox and its children
func (_adf *CT_BorderBox) Validate() error {
	return _adf.ValidateWithPath("\u0043\u0054\u005fB\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078")
}
func (_aeff ST_Jc) String() string {
	switch _aeff {
	case 0:
		return ""
	case 1:
		return "\u006c\u0065\u0066\u0074"
	case 2:
		return "\u0072\u0069\u0067h\u0074"
	case 3:
		return "\u0063\u0065\u006e\u0074\u0065\u0072"
	case 4:
		return "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070"
	}
	return ""
}
func (_fdc *CT_BorderBoxPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fdc.HideTop != nil {
		_deb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0068\u0069\u0064\u0065\u0054\u006fp"}}
		e.EncodeElement(_fdc.HideTop, _deb)
	}
	if _fdc.HideBot != nil {
		_fbc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0068\u0069\u0064\u0065\u0042\u006ft"}}
		e.EncodeElement(_fdc.HideBot, _fbc)
	}
	if _fdc.HideLeft != nil {
		_gebe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0068\u0069\u0064\u0065\u004c\u0065\u0066\u0074"}}
		e.EncodeElement(_fdc.HideLeft, _gebe)
	}
	if _fdc.HideRight != nil {
		_aec := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0068\u0069\u0064\u0065\u0052\u0069\u0067\u0068\u0074"}}
		e.EncodeElement(_fdc.HideRight, _aec)
	}
	if _fdc.StrikeH != nil {
		_aca := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0074\u0072\u0069\u006b\u0065H"}}
		e.EncodeElement(_fdc.StrikeH, _aca)
	}
	if _fdc.StrikeV != nil {
		_fcf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0074\u0072\u0069\u006b\u0065V"}}
		e.EncodeElement(_fdc.StrikeV, _fcf)
	}
	if _fdc.StrikeBLTR != nil {
		_abb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073t\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"}}
		e.EncodeElement(_fdc.StrikeBLTR, _abb)
	}
	if _fdc.StrikeTLBR != nil {
		_gcfe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073t\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"}}
		e.EncodeElement(_fdc.StrikeTLBR, _gcfe)
	}
	if _fdc.CtrlPr != nil {
		_fcb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_fdc.CtrlPr, _fcb)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_cgca ST_Shp) ValidateWithPath(path string) error {
	switch _cgca {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cgca))
	}
	return nil
}

// Validate validates the CT_SSubSupPr and its children
func (_cccdb *CT_SSubSupPr) Validate() error {
	return _cccdb.ValidateWithPath("\u0043\u0054\u005fS\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072")
}
func (_gab *CT_MathPrChoice) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_bbgc:
	for {
		_fgcac, _edfbg := d.Token()
		if _edfbg != nil {
			return _edfbg
		}
		switch _ddgc := _fgcac.(type) {
		case _g.StartElement:
			switch _ddgc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}:
				_gab.WrapIndent = NewCT_TwipsMeasure()
				if _dfbf := d.DecodeElement(_gab.WrapIndent, &_ddgc); _dfbf != nil {
					return _dfbf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}:
				_gab.WrapRight = NewCT_OnOff()
				if _aaaa := d.DecodeElement(_gab.WrapRight, &_ddgc); _aaaa != nil {
					return _aaaa
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004d\u0061\u0074h\u0050\u0072\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _ddgc.Name)
				if _abgfe := d.Skip(); _abgfe != nil {
					return _abgfe
				}
			}
		case _g.EndElement:
			break _bbgc
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Char and its children
func (_cge *CT_Char) Validate() error {
	return _cge.ValidateWithPath("\u0043T\u005f\u0043\u0068\u0061\u0072")
}

// ValidateWithPath validates the CT_Script and its children, prefixing error messages with path
func (_bfde *CT_Script) ValidateWithPath(path string) error {
	if _fgff := _bfde.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fgff != nil {
		return _fgff
	}
	return nil
}

type CT_TopBot struct{ ValAttr ST_TopBot }

// Validate validates the CT_GroupChr and its children
func (_dgea *CT_GroupChr) Validate() error {
	return _dgea.ValidateWithPath("C\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072")
}
func (_aebf ST_BreakBinSub) ValidateWithPath(path string) error {
	switch _aebf {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aebf))
	}
	return nil
}
func (_ecc *CT_OMathParaPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ecc.Jc != nil {
		_effa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006a\u0063"}}
		e.EncodeElement(_ecc.Jc, _effa)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_fda *CT_BreakBin) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _ecf := range start.Attr {
		if _ecf.Name.Local == "\u0076\u0061\u006c" {
			_fda.ValAttr.UnmarshalXMLAttr(_ecf)
			continue
		}
	}
	for {
		_cafc, _fad := d.Token()
		if _fad != nil {
			return _f.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042\u0069\u006e\u003a\u0020\u0025\u0073", _fad)
		}
		if _ddd, _gff := _cafc.(_g.EndElement); _gff && _ddd.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Shp and its children, prefixing error messages with path
func (_afcb *CT_Shp) ValidateWithPath(path string) error {
	if _afcb.ValAttr == ST_ShpUnset {
		return _f.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _cfgc := _afcb.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _cfgc != nil {
		return _cfgc
	}
	return nil
}

type CT_SSub struct {
	SSubPr *CT_SSubPr
	E      *CT_OMathArg
	Sub    *CT_OMathArg
}

// Validate validates the CT_Nary and its children
func (_fgbce *CT_Nary) Validate() error {
	return _fgbce.ValidateWithPath("\u0043T\u005f\u004e\u0061\u0072\u0079")
}
func NewCT_MathPr() *CT_MathPr { _bffc := &CT_MathPr{}; return _bffc }

type CT_EqArr struct {
	EqArrPr *CT_EqArrPr
	E       []*CT_OMathArg
}

func (_afead ST_Shp) String() string {
	switch _afead {
	case 0:
		return ""
	case 1:
		return "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064"
	case 2:
		return "\u006d\u0061\u0074c\u0068"
	}
	return ""
}

// Validate validates the CT_MCS and its children
func (_aacf *CT_MCS) Validate() error {
	return _aacf.ValidateWithPath("\u0043\u0054\u005f\u004d\u0043\u0053")
}

// Validate validates the CT_LimLowPr and its children
func (_daeg *CT_LimLowPr) Validate() error {
	return _daeg.ValidateWithPath("C\u0054\u005f\u004c\u0069\u006d\u004c\u006f\u0077\u0050\u0072")
}
func (_gbfee ST_BreakBin) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_bagdg := _g.Attr{}
	_bagdg.Name = name
	switch _gbfee {
	case ST_BreakBinUnset:
		_bagdg.Value = ""
	case ST_BreakBinBefore:
		_bagdg.Value = "\u0062\u0065\u0066\u006f\u0072\u0065"
	case ST_BreakBinAfter:
		_bagdg.Value = "\u0061\u0066\u0074e\u0072"
	case ST_BreakBinRepeat:
		_bagdg.Value = "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return _bagdg, nil
}

// Validate validates the CT_TwipsMeasure and its children
func (_fbgf *CT_TwipsMeasure) Validate() error {
	return _fbgf.ValidateWithPath("\u0043T\u005fT\u0077\u0069\u0070\u0073\u004d\u0065\u0061\u0073\u0075\u0072\u0065")
}

type CT_DPr struct {
	BegChr *CT_Char
	SepChr *CT_Char
	EndChr *CT_Char
	Grow   *CT_OnOff
	Shp    *CT_Shp
	CtrlPr *CT_CtrlPr
}

func (_abgfa *CT_OnOff) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _defg := range start.Attr {
		if _defg.Name.Local == "\u0076\u0061\u006c" {
			_efc, _befbg := ParseUnionST_OnOff(_defg.Value)
			if _befbg != nil {
				return _befbg
			}
			_abgfa.ValAttr = &_efc
			continue
		}
	}
	for {
		_fgaf, _cbea := d.Token()
		if _cbea != nil {
			return _f.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fO\u006e\u004f\u0066\u0066: \u0025\u0073", _cbea)
		}
		if _bgbd, _eggga := _fgaf.(_g.EndElement); _eggga && _bgbd.Name == start.Name {
			break
		}
	}
	return nil
}
func (_cbbb *ST_FType) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_cbbb = 0
	case "\u0062\u0061\u0072":
		*_cbbb = 1
	case "\u0073\u006b\u0077":
		*_cbbb = 2
	case "\u006c\u0069\u006e":
		*_cbbb = 3
	case "\u006e\u006f\u0042a\u0072":
		*_cbbb = 4
	}
	return nil
}
func NewCT_OMath() *CT_OMath { _eadg := &CT_OMath{}; return _eadg }

type CT_BorderBox struct {
	BorderBoxPr *CT_BorderBoxPr
	E           *CT_OMathArg
}

func NewCT_Func() *CT_Func {
	_cdb := &CT_Func{}
	_cdb.FName = NewCT_OMathArg()
	_cdb.E = NewCT_OMathArg()
	return _cdb
}

// Validate validates the CT_MC and its children
func (_gad *CT_MC) Validate() error { return _gad.ValidateWithPath("\u0043\u0054\u005fM\u0043") }
func (_cceg *CT_SPre) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cceg.SPrePr != nil {
		_gdeb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0050\u0072\u0065\u0050\u0072"}}
		e.EncodeElement(_cceg.SPrePr, _gdeb)
	}
	_deac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_cceg.Sub, _deac)
	_efgc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_cceg.Sup, _efgc)
	_efgb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_cceg.E, _efgb)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_BorderBoxPr struct {
	HideTop    *CT_OnOff
	HideBot    *CT_OnOff
	HideLeft   *CT_OnOff
	HideRight  *CT_OnOff
	StrikeH    *CT_OnOff
	StrikeV    *CT_OnOff
	StrikeBLTR *CT_OnOff
	StrikeTLBR *CT_OnOff
	CtrlPr     *CT_CtrlPr
}

// ValidateWithPath validates the CT_FType and its children, prefixing error messages with path
func (_abgg *CT_FType) ValidateWithPath(path string) error {
	if _abgg.ValAttr == ST_FTypeUnset {
		return _f.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _fegd := _abgg.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fegd != nil {
		return _fegd
	}
	return nil
}
func (_gbcea *ST_FType) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_acgb, _cegdb := d.Token()
	if _cegdb != nil {
		return _cegdb
	}
	if _gegd, _eebdb := _acgb.(_g.EndElement); _eebdb && _gegd.Name == start.Name {
		*_gbcea = 1
		return nil
	}
	if _ddaad, _fbbe := _acgb.(_g.CharData); !_fbbe {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _acgb)
	} else {
		switch string(_ddaad) {
		case "":
			*_gbcea = 0
		case "\u0062\u0061\u0072":
			*_gbcea = 1
		case "\u0073\u006b\u0077":
			*_gbcea = 2
		case "\u006c\u0069\u006e":
			*_gbcea = 3
		case "\u006e\u006f\u0042a\u0072":
			*_gbcea = 4
		}
	}
	_acgb, _cegdb = d.Token()
	if _cegdb != nil {
		return _cegdb
	}
	if _cgfc, _begea := _acgb.(_g.EndElement); _begea && _cgfc.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _acgb)
}
func (_bcbc *EG_OMathMathElements) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fagce:
	for {
		_dabe, _ffgf := d.Token()
		if _ffgf != nil {
			return _ffgf
		}
		switch _fbed := _dabe.(type) {
		case _g.StartElement:
			switch _fbed.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_bcbc.Acc = NewCT_Acc()
				if _dgdb := d.DecodeElement(_bcbc.Acc, &_fbed); _dgdb != nil {
					return _dgdb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_bcbc.Bar = NewCT_Bar()
				if _ffec := d.DecodeElement(_bcbc.Bar, &_fbed); _ffec != nil {
					return _ffec
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_bcbc.Box = NewCT_Box()
				if _aeda := d.DecodeElement(_bcbc.Box, &_fbed); _aeda != nil {
					return _aeda
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_bcbc.BorderBox = NewCT_BorderBox()
				if _fffbd := d.DecodeElement(_bcbc.BorderBox, &_fbed); _fffbd != nil {
					return _fffbd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_bcbc.D = NewCT_D()
				if _gdfdf := d.DecodeElement(_bcbc.D, &_fbed); _gdfdf != nil {
					return _gdfdf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_bcbc.EqArr = NewCT_EqArr()
				if _aefe := d.DecodeElement(_bcbc.EqArr, &_fbed); _aefe != nil {
					return _aefe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_bcbc.F = NewCT_F()
				if _ccfb := d.DecodeElement(_bcbc.F, &_fbed); _ccfb != nil {
					return _ccfb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_bcbc.Func = NewCT_Func()
				if _edaae := d.DecodeElement(_bcbc.Func, &_fbed); _edaae != nil {
					return _edaae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_bcbc.GroupChr = NewCT_GroupChr()
				if _ebeaa := d.DecodeElement(_bcbc.GroupChr, &_fbed); _ebeaa != nil {
					return _ebeaa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_bcbc.LimLow = NewCT_LimLow()
				if _efbe := d.DecodeElement(_bcbc.LimLow, &_fbed); _efbe != nil {
					return _efbe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_bcbc.LimUpp = NewCT_LimUpp()
				if _fgae := d.DecodeElement(_bcbc.LimUpp, &_fbed); _fgae != nil {
					return _fgae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_bcbc.M = NewCT_M()
				if _beacg := d.DecodeElement(_bcbc.M, &_fbed); _beacg != nil {
					return _beacg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_bcbc.Nary = NewCT_Nary()
				if _feag := d.DecodeElement(_bcbc.Nary, &_fbed); _feag != nil {
					return _feag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_bcbc.Phant = NewCT_Phant()
				if _bdecd := d.DecodeElement(_bcbc.Phant, &_fbed); _bdecd != nil {
					return _bdecd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_bcbc.Rad = NewCT_Rad()
				if _dbddg := d.DecodeElement(_bcbc.Rad, &_fbed); _dbddg != nil {
					return _dbddg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_bcbc.SPre = NewCT_SPre()
				if _bgeb := d.DecodeElement(_bcbc.SPre, &_fbed); _bgeb != nil {
					return _bgeb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_bcbc.SSub = NewCT_SSub()
				if _cbdd := d.DecodeElement(_bcbc.SSub, &_fbed); _cbdd != nil {
					return _cbdd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_bcbc.SSubSup = NewCT_SSubSup()
				if _edfeb := d.DecodeElement(_bcbc.SSubSup, &_fbed); _edfeb != nil {
					return _edfeb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_bcbc.SSup = NewCT_SSup()
				if _deff := d.DecodeElement(_bcbc.SSup, &_fbed); _deff != nil {
					return _deff
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_bcbc.R = NewCT_R()
				if _bfce := d.DecodeElement(_bcbc.R, &_fbed); _bfce != nil {
					return _bfce
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070o\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006de\u006et \u006f\u006e\u0020\u0045\u0047\u005f\u004f\u004d\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0020\u0025\u0076", _fbed.Name)
				if _cgff := d.Skip(); _cgff != nil {
					return _cgff
				}
			}
		case _g.EndElement:
			break _fagce
		case _g.CharData:
		}
	}
	return nil
}
func (_dcca *CT_SSubSupPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gebgb:
	for {
		_fgce, _bgcd := d.Token()
		if _bgcd != nil {
			return _bgcd
		}
		switch _aabeb := _fgce.(type) {
		case _g.StartElement:
			switch _aabeb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u006c\u006e\u0053\u0063\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u006c\u006e\u0053\u0063\u0072"}:
				_dcca.AlnScr = NewCT_OnOff()
				if _bfdb := d.DecodeElement(_dcca.AlnScr, &_aabeb); _bfdb != nil {
					return _bfdb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_dcca.CtrlPr = NewCT_CtrlPr()
				if _aceb := d.DecodeElement(_dcca.CtrlPr, &_aabeb); _aceb != nil {
					return _aceb
				}
			default:
				_gg.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_S\u0053\u0075b\u0053\u0075\u0070\u0050\u0072\u0020\u0025\u0076", _aabeb.Name)
				if _febed := d.Skip(); _febed != nil {
					return _febed
				}
			}
		case _g.EndElement:
			break _gebgb
		case _g.CharData:
		}
	}
	return nil
}
func (_dabfg *OMathPara) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006d"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0077"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065s\u0073i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u00306\u002fm\u0061\u0069n"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "m\u003a\u006f\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061"
	return _dabfg.CT_OMathPara.MarshalXML(e, start)
}

// ValidateWithPath validates the CT_MR and its children, prefixing error messages with path
func (_agf *CT_MR) ValidateWithPath(path string) error {
	for _dcgb, _fdaae := range _agf.E {
		if _cegc := _fdaae.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0045\u005b\u0025\u0064\u005d", path, _dcgb)); _cegc != nil {
			return _cegc
		}
	}
	return nil
}

// ValidateWithPath validates the CT_LimLowPr and its children, prefixing error messages with path
func (_acdaa *CT_LimLowPr) ValidateWithPath(path string) error {
	if _acdaa.CtrlPr != nil {
		if _cfbe := _acdaa.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _cfbe != nil {
			return _cfbe
		}
	}
	return nil
}
func NewCT_OMathPara() *CT_OMathPara { _gaffg := &CT_OMathPara{}; return _gaffg }
func (_fbd *CT_Integer255) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fbd.ValAttr = 1
	for _, _gfab := range start.Attr {
		if _gfab.Name.Local == "\u0076\u0061\u006c" {
			_fdb, _abeb := _e.ParseInt(_gfab.Value, 10, 64)
			if _abeb != nil {
				return _abeb
			}
			_fbd.ValAttr = _fdb
			continue
		}
	}
	for {
		_fgbc, _aeed := d.Token()
		if _aeed != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0049n\u0074e\u0067e\u0072\u0032\u0035\u0035\u003a\u0020\u0025s", _aeed)
		}
		if _gaag, _ddda := _fgbc.(_g.EndElement); _ddda && _gaag.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_SSubSupPr() *CT_SSubSupPr { _feaa := &CT_SSubSupPr{}; return _feaa }
func (_feae *CT_LimUppPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fef:
	for {
		_bded, _ffce := d.Token()
		if _ffce != nil {
			return _ffce
		}
		switch _dcda := _bded.(type) {
		case _g.StartElement:
			switch _dcda.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_feae.CtrlPr = NewCT_CtrlPr()
				if _ceaf := d.DecodeElement(_feae.CtrlPr, &_dcda); _ceaf != nil {
					return _ceaf
				}
			default:
				_gg.Log("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004c\u0069\u006d\u0055\u0070\u0070\u0050\u0072\u0020\u0025\u0076", _dcda.Name)
				if _fgaae := d.Skip(); _fgaae != nil {
					return _fgaae
				}
			}
		case _g.EndElement:
			break _fef
		case _g.CharData:
		}
	}
	return nil
}
func (_afaf *ST_Shp) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_geee, _eceff := d.Token()
	if _eceff != nil {
		return _eceff
	}
	if _cgfg, _eadgg := _geee.(_g.EndElement); _eadgg && _cgfg.Name == start.Name {
		*_afaf = 1
		return nil
	}
	if _fcceb, _fabg := _geee.(_g.CharData); !_fabg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _geee)
	} else {
		switch string(_fcceb) {
		case "":
			*_afaf = 0
		case "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064":
			*_afaf = 1
		case "\u006d\u0061\u0074c\u0068":
			*_afaf = 2
		}
	}
	_geee, _eceff = d.Token()
	if _eceff != nil {
		return _eceff
	}
	if _gedfg, _cead := _geee.(_g.EndElement); _cead && _gedfg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _geee)
}

// Validate validates the CT_Func and its children
func (_aabe *CT_Func) Validate() error {
	return _aabe.ValidateWithPath("\u0043T\u005f\u0046\u0075\u006e\u0063")
}

// ValidateWithPath validates the CT_Text and its children, prefixing error messages with path
func (_bebg *CT_Text) ValidateWithPath(path string) error { return nil }

// Validate validates the CT_String and its children
func (_cefbd *CT_String) Validate() error {
	return _cefbd.ValidateWithPath("\u0043T\u005f\u0053\u0074\u0072\u0069\u006eg")
}

// Validate validates the CT_SSub and its children
func (_bagc *CT_SSub) Validate() error {
	return _bagc.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0062")
}
func (_gbgg *CT_XAlign) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_gdcae, _ggdb := _gbgg.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _ggdb != nil {
		return _ggdb
	}
	start.Attr = append(start.Attr, _gdcae)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_dfead ST_BreakBinSub) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_dfead.String(), start)
}

// Validate validates the CT_ManualBreak and its children
func (_bcce *CT_ManualBreak) Validate() error {
	return _bcce.ValidateWithPath("\u0043\u0054\u005f\u004d\u0061\u006e\u0075\u0061\u006cB\u0072\u0065\u0061\u006b")
}
func (_becf *CT_BorderBox) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_becf.E = NewCT_OMathArg()
_ae:
	for {
		_ea, _cfd := d.Token()
		if _cfd != nil {
			return _cfd
		}
		switch _gdab := _ea.(type) {
		case _g.StartElement:
			switch _gdab.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "b\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "b\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"}:
				_becf.BorderBoxPr = NewCT_BorderBoxPr()
				if _aac := d.DecodeElement(_becf.BorderBoxPr, &_gdab); _aac != nil {
					return _aac
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _caf := d.DecodeElement(_becf.E, &_gdab); _caf != nil {
					return _caf
				}
			default:
				_gg.Log("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_B\u006f\u0072d\u0065\u0072\u0042\u006f\u0078\u0020\u0025\u0076", _gdab.Name)
				if _cfe := d.Skip(); _cfe != nil {
					return _cfe
				}
			}
		case _g.EndElement:
			break _ae
		case _g.CharData:
		}
	}
	return nil
}

const (
	ST_BreakBinUnset  ST_BreakBin = 0
	ST_BreakBinBefore ST_BreakBin = 1
	ST_BreakBinAfter  ST_BreakBin = 2
	ST_BreakBinRepeat ST_BreakBin = 3
)
const (
	ST_BreakBinSubUnset ST_BreakBinSub = 0
	ST_BreakBinSub__    ST_BreakBinSub = 1
	ST_BreakBinSub___   ST_BreakBinSub = 2
	ST_BreakBinSub____  ST_BreakBinSub = 3
)

func (_gdgc *CT_Script) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _dcee := range start.Attr {
		if _dcee.Name.Local == "\u0076\u0061\u006c" {
			_gdgc.ValAttr.UnmarshalXMLAttr(_dcee)
			continue
		}
	}
	for {
		_decb, _dbbcf := d.Token()
		if _dbbcf != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0072i\u0070\u0074\u003a\u0020\u0025\u0073", _dbbcf)
		}
		if _bbab, _ggaag := _decb.(_g.EndElement); _ggaag && _bbab.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the OMathPara and its children, prefixing error messages with path
func (_ffag *OMathPara) ValidateWithPath(path string) error {
	if _dedbf := _ffag.CT_OMathPara.ValidateWithPath(path); _dedbf != nil {
		return _dedbf
	}
	return nil
}

type CT_ManualBreak struct{ AlnAtAttr *int64 }

func NewCT_AccPr() *CT_AccPr { _dd := &CT_AccPr{}; return _dd }

// ValidateWithPath validates the CT_MC and its children, prefixing error messages with path
func (_bdfd *CT_MC) ValidateWithPath(path string) error {
	if _bdfd.McPr != nil {
		if _ebdc := _bdfd.McPr.ValidateWithPath(path + "\u002f\u004d\u0063P\u0072"); _ebdc != nil {
			return _ebdc
		}
	}
	return nil
}
func (_gcae *CT_SSubPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ffbfa:
	for {
		_gcfgf, _fgde := d.Token()
		if _fgde != nil {
			return _fgde
		}
		switch _edecg := _gcfgf.(type) {
		case _g.StartElement:
			switch _edecg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_gcae.CtrlPr = NewCT_CtrlPr()
				if _daaf := d.DecodeElement(_gcae.CtrlPr, &_edecg); _daaf != nil {
					return _daaf
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0053\u0053\u0075b\u0050\u0072 \u0025\u0076", _edecg.Name)
				if _bafa := d.Skip(); _bafa != nil {
					return _bafa
				}
			}
		case _g.EndElement:
			break _ffbfa
		case _g.CharData:
		}
	}
	return nil
}
func (_fffc *CT_NaryPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_dceb:
	for {
		_dfbg, _bedf := d.Token()
		if _bedf != nil {
			return _bedf
		}
		switch _bbecg := _dfbg.(type) {
		case _g.StartElement:
			switch _bbecg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0068\u0072"}:
				_fffc.Chr = NewCT_Char()
				if _fbba := d.DecodeElement(_fffc.Chr, &_bbecg); _fbba != nil {
					return _fbba
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0063"}:
				_fffc.LimLoc = NewCT_LimLoc()
				if _cgg := d.DecodeElement(_fffc.LimLoc, &_bbecg); _cgg != nil {
					return _cgg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}:
				_fffc.Grow = NewCT_OnOff()
				if _gefdc := d.DecodeElement(_fffc.Grow, &_bbecg); _gefdc != nil {
					return _gefdc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073u\u0062\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073u\u0062\u0048\u0069\u0064\u0065"}:
				_fffc.SubHide = NewCT_OnOff()
				if _edfea := d.DecodeElement(_fffc.SubHide, &_bbecg); _edfea != nil {
					return _edfea
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073u\u0070\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073u\u0070\u0048\u0069\u0064\u0065"}:
				_fffc.SupHide = NewCT_OnOff()
				if _faaf := d.DecodeElement(_fffc.SupHide, &_bbecg); _faaf != nil {
					return _faaf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fffc.CtrlPr = NewCT_CtrlPr()
				if _gcfa := d.DecodeElement(_fffc.CtrlPr, &_bbecg); _gcfa != nil {
					return _gcfa
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004e\u0061\u0072y\u0050\u0072 \u0025\u0076", _bbecg.Name)
				if _ebeb := d.Skip(); _ebeb != nil {
					return _ebeb
				}
			}
		case _g.EndElement:
			break _dceb
		case _g.CharData:
		}
	}
	return nil
}
func (_daf *CT_BreakBin) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _daf.ValAttr != ST_BreakBinUnset {
		_bab, _fab := _daf.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _fab != nil {
			return _fab
		}
		start.Attr = append(start.Attr, _bab)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_BreakBinSub struct{ ValAttr ST_BreakBinSub }

// ValidateWithPath validates the CT_R and its children, prefixing error messages with path
func (_gcbb *CT_R) ValidateWithPath(path string) error {
	if _gcbb.RPr != nil {
		if _faed := _gcbb.RPr.ValidateWithPath(path + "\u002f\u0052\u0050\u0072"); _faed != nil {
			return _faed
		}
	}
	for _agce, _fgcd := range _gcbb.Choice {
		if _fedg := _fgcd.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _agce)); _fedg != nil {
			return _fedg
		}
	}
	return nil
}

type CT_MCPr struct {
	Count *CT_Integer255
	McJc  *CT_XAlign
}

// Validate validates the EG_ScriptStyle and its children
func (_eeafg *EG_ScriptStyle) Validate() error {
	return _eeafg.ValidateWithPath("\u0045\u0047\u005f\u0053\u0063\u0072\u0069\u0070\u0074S\u0074\u0079\u006c\u0065")
}
func (_afae *CT_LimLow) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_afae.E = NewCT_OMathArg()
	_afae.Lim = NewCT_OMathArg()
_fdfd:
	for {
		_agga, _fabf := d.Token()
		if _fabf != nil {
			return _fabf
		}
		switch _eggg := _agga.(type) {
		case _g.StartElement:
			switch _eggg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077\u0050\u0072"}:
				_afae.LimLowPr = NewCT_LimLowPr()
				if _dbgc := d.DecodeElement(_afae.LimLowPr, &_eggg); _dbgc != nil {
					return _dbgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _ddeca := d.DecodeElement(_afae.E, &_eggg); _ddeca != nil {
					return _ddeca
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d"}:
				if _gdce := d.DecodeElement(_afae.Lim, &_eggg); _gdce != nil {
					return _gdce
				}
			default:
				_gg.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004c\u0069\u006dL\u006f\u0077 \u0025\u0076", _eggg.Name)
				if _dda := d.Skip(); _dda != nil {
					return _dda
				}
			}
		case _g.EndElement:
			break _fdfd
		case _g.CharData:
		}
	}
	return nil
}
func (_gda *CT_BorderBox) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gda.BorderBoxPr != nil {
		_egc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"}}
		e.EncodeElement(_gda.BorderBoxPr, _egc)
	}
	_ac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_gda.E, _ac)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_FuncPr struct{ CtrlPr *CT_CtrlPr }

// Validate validates the CT_SSubSup and its children
func (_cafda *CT_SSubSup) Validate() error {
	return _cafda.ValidateWithPath("\u0043\u0054\u005f\u0053\u0053\u0075\u0062\u0053\u0075\u0070")
}
func (_fcfbc *CT_SSupPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fcfbc.CtrlPr != nil {
		_eae := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_fcfbc.CtrlPr, _eae)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_cbf *CT_BreakBinSub) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _cbf.ValAttr != ST_BreakBinSubUnset {
		_dbb, _dgg := _cbf.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _dgg != nil {
			return _dgg
		}
		start.Attr = append(start.Attr, _dbb)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_acaf *CT_Char) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _gfce := range start.Attr {
		if _gfce.Name.Local == "\u0076\u0061\u006c" {
			_ebd, _dge := _gfce.Value, error(nil)
			if _dge != nil {
				return _dge
			}
			_acaf.ValAttr = _ebd
			continue
		}
	}
	for {
		_eca, _ecff := d.Token()
		if _ecff != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0043\u0068\u0061\u0072\u003a\u0020\u0025\u0073", _ecff)
		}
		if _fcbf, _gbe := _eca.(_g.EndElement); _gbe && _fcbf.Name == start.Name {
			break
		}
	}
	return nil
}
func (_aegb *CT_String) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _eeaa := range start.Attr {
		if _eeaa.Name.Local == "\u0076\u0061\u006c" {
			_eaeb, _eabb := _eeaa.Value, error(nil)
			if _eabb != nil {
				return _eabb
			}
			_aegb.ValAttr = &_eaeb
			continue
		}
	}
	for {
		_fcgd, _gdba := d.Token()
		if _gdba != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0074\u0072i\u006e\u0067\u003a\u0020\u0025\u0073", _gdba)
		}
		if _edgb, _fdbd := _fcgd.(_g.EndElement); _fdbd && _edgb.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_OnOff and its children, prefixing error messages with path
func (_cabbbc *CT_OnOff) ValidateWithPath(path string) error {
	if _cabbbc.ValAttr != nil {
		if _fdg := _cabbbc.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fdg != nil {
			return _fdg
		}
	}
	return nil
}
func (_gaee *CT_BoxPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gaee.OpEmu != nil {
		_agg := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006f\u0070\u0045\u006d\u0075"}}
		e.EncodeElement(_gaee.OpEmu, _agg)
	}
	if _gaee.NoBreak != nil {
		_cg := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006e\u006f\u0042\u0072\u0065\u0061k"}}
		e.EncodeElement(_gaee.NoBreak, _cg)
	}
	if _gaee.Diff != nil {
		_ddf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064\u0069\u0066\u0066"}}
		e.EncodeElement(_gaee.Diff, _ddf)
	}
	if _gaee.Brk != nil {
		_egd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062r\u006b"}}
		e.EncodeElement(_gaee.Brk, _egd)
	}
	if _gaee.Aln != nil {
		_ddfb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061l\u006e"}}
		e.EncodeElement(_gaee.Aln, _ddfb)
	}
	if _gaee.CtrlPr != nil {
		_eec := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_gaee.CtrlPr, _eec)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_LimUppPr() *CT_LimUppPr { _acfb := &CT_LimUppPr{}; return _acfb }
func (_ddee ST_LimLoc) ValidateWithPath(path string) error {
	switch _ddee {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ddee))
	}
	return nil
}

// ValidateWithPath validates the CT_BreakBin and its children, prefixing error messages with path
func (_efb *CT_BreakBin) ValidateWithPath(path string) error {
	if _gef := _efb.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gef != nil {
		return _gef
	}
	return nil
}

// Validate validates the CT_Phant and its children
func (_ebgd *CT_Phant) Validate() error {
	return _ebgd.ValidateWithPath("\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074")
}
func (_cabe *CT_TopBot) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cabe.ValAttr = ST_TopBot(1)
	for _, _fdgb := range start.Attr {
		if _fdgb.Name.Local == "\u0076\u0061\u006c" {
			_cabe.ValAttr.UnmarshalXMLAttr(_fdgb)
			continue
		}
	}
	for {
		_ggdf, _fccf := d.Token()
		if _fccf != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0054\u006f\u0070B\u006f\u0074\u003a\u0020\u0025\u0073", _fccf)
		}
		if _eggea, _fgcde := _ggdf.(_g.EndElement); _fgcde && _eggea.Name == start.Name {
			break
		}
	}
	return nil
}

type EG_ScriptStyle struct {
	Scr *CT_Script
	Sty *CT_Style
}

// ValidateWithPath validates the CT_BreakBinSub and its children, prefixing error messages with path
func (_cbcgd *CT_BreakBinSub) ValidateWithPath(path string) error {
	if _dag := _cbcgd.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _dag != nil {
		return _dag
	}
	return nil
}

// ValidateWithPath validates the CT_FPr and its children, prefixing error messages with path
func (_bge *CT_FPr) ValidateWithPath(path string) error {
	if _bge.Type != nil {
		if _efe := _bge.Type.ValidateWithPath(path + "\u002f\u0054\u0079p\u0065"); _efe != nil {
			return _efe
		}
	}
	if _bge.CtrlPr != nil {
		if _dffa := _bge.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _dffa != nil {
			return _dffa
		}
	}
	return nil
}

type CT_OMathArgPr struct{ ArgSz *CT_Integer2 }
type CT_LimUpp struct {
	LimUppPr *CT_LimUppPr
	E        *CT_OMathArg
	Lim      *CT_OMathArg
}

func NewCT_MathPrChoice() *CT_MathPrChoice { _bdfa := &CT_MathPrChoice{}; return _bdfa }

// Validate validates the OMath and its children
func (_aecda *OMath) Validate() error { return _aecda.ValidateWithPath("\u004f\u004d\u0061t\u0068") }
func (_aee *CT_Integer2) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_aee.ValAttr = -2
	for _, _afce := range start.Attr {
		if _afce.Name.Local == "\u0076\u0061\u006c" {
			_ggc, _eaa := _e.ParseInt(_afce.Value, 10, 64)
			if _eaa != nil {
				return _eaa
			}
			_aee.ValAttr = _ggc
			continue
		}
	}
	for {
		_bfb, _fggf := d.Token()
		if _fggf != nil {
			return _f.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032\u003a\u0020\u0025\u0073", _fggf)
		}
		if _fedb, _bdda := _bfb.(_g.EndElement); _bdda && _fedb.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_LimLow and its children, prefixing error messages with path
func (_bgfb *CT_LimLow) ValidateWithPath(path string) error {
	if _bgfb.LimLowPr != nil {
		if _dfgg := _bgfb.LimLowPr.ValidateWithPath(path + "\u002fL\u0069\u006d\u004c\u006f\u0077\u0050r"); _dfgg != nil {
			return _dfgg
		}
	}
	if _facf := _bgfb.E.ValidateWithPath(path + "\u002f\u0045"); _facf != nil {
		return _facf
	}
	if _gagb := _bgfb.Lim.ValidateWithPath(path + "\u002f\u004c\u0069\u006d"); _gagb != nil {
		return _gagb
	}
	return nil
}
func NewCT_NaryPr() *CT_NaryPr { _bgfd := &CT_NaryPr{}; return _bgfd }
func (_cagc *CT_Phant) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cagc.E = NewCT_OMathArg()
_cdad:
	for {
		_dcfcc, _ffdfd := d.Token()
		if _ffdfd != nil {
			return _ffdfd
		}
		switch _bbagg := _dcfcc.(type) {
		case _g.StartElement:
			switch _bbagg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070h\u0061\u006e\u0074\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070h\u0061\u006e\u0074\u0050\u0072"}:
				_cagc.PhantPr = NewCT_PhantPr()
				if _cafe := d.DecodeElement(_cagc.PhantPr, &_bbagg); _cafe != nil {
					return _cafe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _dcdff := d.DecodeElement(_cagc.E, &_bbagg); _dcdff != nil {
					return _dcdff
				}
			default:
				_gg.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074\u0020\u0025\u0076", _bbagg.Name)
				if _fbdg := d.Skip(); _fbdg != nil {
					return _fbdg
				}
			}
		case _g.EndElement:
			break _cdad
		case _g.CharData:
		}
	}
	return nil
}
func (_bfad ST_FType) Validate() error { return _bfad.ValidateWithPath("") }
func (_abbf *CT_F) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _abbf.FPr != nil {
		_cac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066P\u0072"}}
		e.EncodeElement(_abbf.FPr, _cac)
	}
	_gfb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006eu\u006d"}}
	e.EncodeElement(_abbf.Num, _gfb)
	_bgc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064e\u006e"}}
	e.EncodeElement(_abbf.Den, _bgc)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_GroupChrPr() *CT_GroupChrPr { _ebf := &CT_GroupChrPr{}; return _ebf }
func init() {
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032\u0035\u0035", NewCT_Integer255)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032", NewCT_Integer2)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0070\u0061\u0063\u0069\u006eg\u0052\u0075\u006c\u0065", NewCT_SpacingRule)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005fU\u006e\u0053\u0069\u0067\u006ee\u0064\u0049n\u0074\u0065\u0067\u0065\u0072", NewCT_UnSignedInteger)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0043\u0068\u0061\u0072", NewCT_Char)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u006e\u004f\u0066\u0066", NewCT_OnOff)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0074\u0072\u0069\u006eg", NewCT_String)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0058\u0041\u006c\u0069\u0067n", NewCT_XAlign)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0059\u0041\u006c\u0069\u0067n", NewCT_YAlign)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0068\u0070", NewCT_Shp)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0046\u0054\u0079\u0070\u0065", NewCT_FType)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004c\u0069\u006d\u004c\u006fc", NewCT_LimLoc)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0054\u006f\u0070\u0042\u006ft", NewCT_TopBot)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0063\u0072\u0069\u0070t", NewCT_Script)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065", NewCT_Style)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d\u0061\u006e\u0075\u0061\u006cB\u0072\u0065\u0061\u006b", NewCT_ManualBreak)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052\u0050\u0052", NewCT_RPR)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0054\u0065\u0078\u0074", NewCT_Text)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052", NewCT_R)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0043\u0074\u0072\u006c\u0050r", NewCT_CtrlPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0041\u0063\u0063\u0050\u0072", NewCT_AccPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0041\u0063\u0063", NewCT_Acc)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u0061\u0072\u0050\u0072", NewCT_BarPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u0061\u0072", NewCT_Bar)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u006f\u0078\u0050\u0072", NewCT_BoxPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u006f\u0078", NewCT_Box)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u006f\u0072\u0064\u0065\u0072B\u006f\u0078\u0050\u0072", NewCT_BorderBoxPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fB\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078", NewCT_BorderBox)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0044\u0050\u0072", NewCT_DPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0044", NewCT_D)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072\u0050\u0072", NewCT_EqArrPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072", NewCT_EqArr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0046\u0050\u0072", NewCT_FPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0046", NewCT_F)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0046\u0075\u006e\u0063\u0050r", NewCT_FuncPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0046\u0075\u006e\u0063", NewCT_Func)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072", NewCT_GroupChrPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072", NewCT_GroupChr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u004c\u0069\u006d\u004c\u006f\u0077\u0050\u0072", NewCT_LimLowPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004c\u0069\u006d\u004c\u006fw", NewCT_LimLow)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u004c\u0069\u006d\u0055\u0070\u0070\u0050\u0072", NewCT_LimUppPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004c\u0069\u006d\u0055\u0070p", NewCT_LimUpp)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004d\u0043\u0050\u0072", NewCT_MCPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fM\u0043", NewCT_MC)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d\u0043\u0053", NewCT_MCS)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d\u0050\u0072", NewCT_MPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fM\u0052", NewCT_MR)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d", NewCT_M)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004e\u0061\u0072\u0079\u0050r", NewCT_NaryPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004e\u0061\u0072\u0079", NewCT_Nary)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074\u0050\u0072", NewCT_PhantPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074", NewCT_Phant)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052\u0061\u0064\u0050\u0072", NewCT_RadPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052\u0061\u0064", NewCT_Rad)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0050\u0072\u0065\u0050r", NewCT_SPrePr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0050\u0072\u0065", NewCT_SPre)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0062\u0050r", NewCT_SSubPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0062", NewCT_SSub)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fS\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072", NewCT_SSubSupPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0053\u0075\u0062\u0053\u0075\u0070", NewCT_SSubSup)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0070\u0050r", NewCT_SSupPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0070", NewCT_SSup)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067\u0050\u0072", NewCT_OMathArgPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067", NewCT_OMathArg)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u004a\u0063", NewCT_OMathJc)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0050a\u0072\u0061\u0050\u0072", NewCT_OMathParaPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005fT\u0077\u0069\u0070\u0073\u004d\u0065\u0061\u0073\u0075\u0072\u0065", NewCT_TwipsMeasure)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042\u0069\u006e", NewCT_BreakBin)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042i\u006e\u0053\u0075\u0062", NewCT_BreakBinSub)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004d\u0061\u0074\u0068\u0050r", NewCT_MathPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fO\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061", NewCT_OMathPara)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068", NewCT_OMath)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u006d\u0061\u0074\u0068\u0050\u0072", NewMathPr)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u006fM\u0061\u0074\u0068\u0050\u0061\u0072a", NewOMathPara)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u006f\u004d\u0061t\u0068", NewOMath)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0045\u0047\u005f\u0053\u0063\u0072\u0069\u0070\u0074S\u0074\u0079\u006c\u0065", NewEG_ScriptStyle)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "E\u0047_\u004f\u004d\u0061\u0074\u0068\u004d\u0061\u0074h\u0045\u006c\u0065\u006den\u0074\u0073", NewEG_OMathMathElements)
	_gg.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0045\u0047_\u004f\u004d\u0061t\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073", NewEG_OMathElements)
}
