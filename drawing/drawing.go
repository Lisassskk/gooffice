
package drawing ;import (_ff "gitee.com/gooffice/gooffice";_c "gitee.com/gooffice/gooffice/color";_g "gitee.com/gooffice/gooffice/measurement";_fc "gitee.com/gooffice/gooffice/schema/soo/dml";);

// Properties returns the paragraph properties.
func (_dd Paragraph )Properties ()ParagraphProperties {if _dd ._fe .PPr ==nil {_dd ._fe .PPr =_fc .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_dd ._fe .PPr );};

// SetSize sets the width and height of the shape.
func (_dab ShapeProperties )SetSize (w ,h _g .Distance ){_dab .SetWidth (w );_dab .SetHeight (h )};

// SetBold controls the bolding of a run.
func (_ga RunProperties )SetBold (b bool ){_ga ._aa .BAttr =_ff .Bool (b )};

// SetNumbered controls if bullets are numbered or not.
func (_ee ParagraphProperties )SetNumbered (scheme _fc .ST_TextAutonumberScheme ){if scheme ==_fc .ST_TextAutonumberSchemeUnset {_ee ._ce .BuAutoNum =nil ;}else {_ee ._ce .BuAutoNum =_fc .NewCT_TextAutonumberBullet ();_ee ._ce .BuAutoNum .TypeAttr =scheme ;};};func (_gg ShapeProperties )SetSolidFill (c _c .Color ){_gg .clearFill ();_gg ._da .SolidFill =_fc .NewCT_SolidColorFillProperties ();_gg ._da .SolidFill .SrgbClr =_fc .NewCT_SRgbColor ();_gg ._da .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_af LineProperties )SetWidth (w _g .Distance ){_af ._b .WAttr =_ff .Int32 (int32 (w /_g .EMU ))};

// SetBulletFont controls the font for the bullet character.
func (_be ParagraphProperties )SetBulletFont (f string ){if f ==""{_be ._ce .BuFont =nil ;}else {_be ._ce .BuFont =_fc .NewCT_TextFont ();_be ._ce .BuFont .TypefaceAttr =f ;};};

// SetFlipVertical controls if the shape is flipped vertically.
func (_ac ShapeProperties )SetFlipVertical (b bool ){_ac .ensureXfrm ();if !b {_ac ._da .Xfrm .FlipVAttr =nil ;}else {_ac ._da .Xfrm .FlipVAttr =_ff .Bool (true );};};

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_fc .EG_TextRun )Run {return Run {x }};func (_fd ShapeProperties )SetNoFill (){_fd .clearFill ();_fd ._da .NoFill =_fc .NewCT_NoFillProperties ();};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_ce *_fc .CT_TextParagraphProperties ;};

// SetJoin sets the line join style.
func (_ef LineProperties )SetJoin (e LineJoin ){_ef ._b .Round =nil ;_ef ._b .Miter =nil ;_ef ._b .Bevel =nil ;switch e {case LineJoinRound :_ef ._b .Round =_fc .NewCT_LineJoinRound ();case LineJoinBevel :_ef ._b .Bevel =_fc .NewCT_LineJoinBevel ();case LineJoinMiter :_ef ._b .Miter =_fc .NewCT_LineJoinMiterProperties ();};};

// X returns the inner wrapped XML type.
func (_ea ShapeProperties )X ()*_fc .CT_ShapeProperties {return _ea ._da };

// Properties returns the run's properties.
func (_dbe Run )Properties ()RunProperties {if _dbe ._eee .R ==nil {_dbe ._eee .R =_fc .NewCT_RegularTextRun ();};if _dbe ._eee .R .RPr ==nil {_dbe ._eee .R .RPr =_fc .NewCT_TextCharacterProperties ();};return RunProperties {_dbe ._eee .R .RPr };};

// AddRun adds a new run to a paragraph.
func (_cg Paragraph )AddRun ()Run {_fg :=MakeRun (_fc .NewEG_TextRun ());_cg ._fe .EG_TextRun =append (_cg ._fe .EG_TextRun ,_fg .X ());return _fg ;};

// SetWidth sets the width of the shape.
func (_cb ShapeProperties )SetWidth (w _g .Distance ){_cb .ensureXfrm ();if _cb ._da .Xfrm .Ext ==nil {_cb ._da .Xfrm .Ext =_fc .NewCT_PositiveSize2D ();};_cb ._da .Xfrm .Ext .CxAttr =int64 (w /_g .EMU );};func MakeShapeProperties (x *_fc .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};

// Paragraph is a paragraph within a document.
type Paragraph struct{_fe *_fc .CT_TextParagraph };

// SetText sets the run's text contents.
func (_eg Run )SetText (s string ){_eg ._eee .Br =nil ;_eg ._eee .Fld =nil ;if _eg ._eee .R ==nil {_eg ._eee .R =_fc .NewCT_RegularTextRun ();};_eg ._eee .R .T =s ;};

// X returns the inner wrapped XML type.
func (_db ParagraphProperties )X ()*_fc .CT_TextParagraphProperties {return _db ._ce };type ShapeProperties struct{_da *_fc .CT_ShapeProperties };

// SetBulletChar sets the bullet character for the paragraph.
func (_bd ParagraphProperties )SetBulletChar (c string ){if c ==""{_bd ._ce .BuChar =nil ;}else {_bd ._ce .BuChar =_fc .NewCT_TextCharBullet ();_bd ._ce .BuChar .CharAttr =c ;};};

// GetPosition gets the position of the shape in EMU.
func (_eb ShapeProperties )GetPosition ()(int64 ,int64 ){_eb .ensureXfrm ();if _eb ._da .Xfrm .Off ==nil {_eb ._da .Xfrm .Off =_fc .NewCT_Point2D ();};return *_eb ._da .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_eb ._da .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;};

// SetAlign controls the paragraph alignment
func (_fa ParagraphProperties )SetAlign (a _fc .ST_TextAlignType ){_fa ._ce .AlgnAttr =a };

// SetGeometry sets the shape type of the shape
func (_fgc ShapeProperties )SetGeometry (g _fc .ST_ShapeType ){if _fgc ._da .PrstGeom ==nil {_fgc ._da .PrstGeom =_fc .NewCT_PresetGeometry2D ();};_fgc ._da .PrstGeom .PrstAttr =g ;};func (_ba LineProperties )SetSolidFill (c _c .Color ){_ba .clearFill ();_ba ._b .SolidFill =_fc .NewCT_SolidColorFillProperties ();_ba ._b .SolidFill .SrgbClr =_fc .NewCT_SRgbColor ();_ba ._b .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};func (_e LineProperties )SetNoFill (){_e .clearFill ();_e ._b .NoFill =_fc .NewCT_NoFillProperties ()};

// X returns the inner wrapped XML type.
func (_a LineProperties )X ()*_fc .CT_LineProperties {return _a ._b };

// SetSize sets the font size of the run text
func (_df RunProperties )SetSize (sz _g .Distance ){_df ._aa .SzAttr =_ff .Int32 (int32 (sz /_g .HundredthPoint ));};

// X returns the inner wrapped XML type.
func (_bg Run )X ()*_fc .EG_TextRun {return _bg ._eee };func (_de ShapeProperties )ensureXfrm (){if _de ._da .Xfrm ==nil {_de ._da .Xfrm =_fc .NewCT_Transform2D ();};};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_dea ShapeProperties )SetFlipHorizontal (b bool ){_dea .ensureXfrm ();if !b {_dea ._da .Xfrm .FlipHAttr =nil ;}else {_dea ._da .Xfrm .FlipHAttr =_ff .Bool (true );};};

// SetPosition sets the position of the shape.
func (_dgf ShapeProperties )SetPosition (x ,y _g .Distance ){_dgf .ensureXfrm ();if _dgf ._da .Xfrm .Off ==nil {_dgf ._da .Xfrm .Off =_fc .NewCT_Point2D ();};_dgf ._da .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_ff .Int64 (int64 (x /_g .EMU ));_dgf ._da .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_ff .Int64 (int64 (y /_g .EMU ));};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_fc .CT_TextParagraph )Paragraph {return Paragraph {x }};

// Run is a run within a paragraph.
type Run struct{_eee *_fc .EG_TextRun };func (_ad ShapeProperties )clearFill (){_ad ._da .NoFill =nil ;_ad ._da .BlipFill =nil ;_ad ._da .GradFill =nil ;_ad ._da .GrpFill =nil ;_ad ._da .SolidFill =nil ;_ad ._da .PattFill =nil ;};

// SetSolidFill controls the text color of a run.
func (_eeg RunProperties )SetSolidFill (c _c .Color ){_eeg ._aa .NoFill =nil ;_eeg ._aa .BlipFill =nil ;_eeg ._aa .GradFill =nil ;_eeg ._aa .GrpFill =nil ;_eeg ._aa .PattFill =nil ;_eeg ._aa .SolidFill =_fc .NewCT_SolidColorFillProperties ();_eeg ._aa .SolidFill .SrgbClr =_fc .NewCT_SRgbColor ();_eeg ._aa .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetHeight sets the height of the shape.
func (_bc ShapeProperties )SetHeight (h _g .Distance ){_bc .ensureXfrm ();if _bc ._da .Xfrm .Ext ==nil {_bc ._da .Xfrm .Ext =_fc .NewCT_PositiveSize2D ();};_bc ._da .Xfrm .Ext .CyAttr =int64 (h /_g .EMU );};

// SetFont controls the font of a run.
func (_eeee RunProperties )SetFont (s string ){_eeee ._aa .Latin =_fc .NewCT_TextFont ();_eeee ._aa .Latin .TypefaceAttr =s ;};

// X returns the inner wrapped XML type.
func (_afg Paragraph )X ()*_fc .CT_TextParagraph {return _afg ._fe };

// SetLevel sets the level of indentation of a paragraph.
func (_cc ParagraphProperties )SetLevel (idx int32 ){_cc ._ce .LvlAttr =_ff .Int32 (idx )};type LineProperties struct{_b *_fc .CT_LineProperties };

// LineJoin is the type of line join
type LineJoin byte ;func (_d LineProperties )clearFill (){_d ._b .NoFill =nil ;_d ._b .GradFill =nil ;_d ._b .SolidFill =nil ;_d ._b .PattFill =nil ;};

// AddBreak adds a new line break to a paragraph.
func (_dg Paragraph )AddBreak (){_gc :=_fc .NewEG_TextRun ();_gc .Br =_fc .NewCT_TextLineBreak ();_dg ._fe .EG_TextRun =append (_dg ._fe .EG_TextRun ,_gc );};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_fc .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};

// RunProperties controls the run properties.
type RunProperties struct{_aa *_fc .CT_TextCharacterProperties ;};func (_gb ShapeProperties )LineProperties ()LineProperties {if _gb ._da .Ln ==nil {_gb ._da .Ln =_fc .NewCT_LineProperties ();};return LineProperties {_gb ._da .Ln };};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_fc .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};
