
package algo ;import _b "strconv";func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_ba :=make ([]byte ,len (s )*cnt );_ea :=[]byte (s );for _c :=0;_c < cnt ;_c ++{copy (_ba [_c :],_ea );};return string (_ba );};func _d (_e byte )bool {return _e >='0'&&_e <='9'};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_dg ,_g :=0,0;for _dg < len (lhs )&&_g < len (rhs ){_gc :=lhs [_dg ];_dc :=rhs [_g ];_gf :=_d (_gc );_gd :=_d (_dc );switch {case _gf &&!_gd :return true ;case !_gf &&_gd :return false ;case !_gf &&!_gd :if _gc !=_dc {return _gc < _dc ;};_dg ++;_g ++;default:_ae :=_dg +1;_ge :=_g +1;for _ae < len (lhs )&&_d (lhs [_ae ]){_ae ++;};for _ge < len (rhs )&&_d (rhs [_ge ]){_ge ++;};_ec ,_ :=_b .ParseUint (lhs [_dg :_ae ],10,64);_aa ,_ :=_b .ParseUint (rhs [_dg :_ge ],10,64);if _ec !=_aa {return _ec < _aa ;};_dg =_ae ;_g =_ge ;};};return len (lhs )< len (rhs );};
