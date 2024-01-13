//line :1
package math

var gOnNGf = [...]float64{
	-9.64399179425052238628e-1,
	-9.92877231001918586564e1,
	-1.61468768441708447952e3,
}
var iX5w2LUYwJ0 = [...]float64{
	1.12811678491632931402e2,
	2.23548839060100448583e3,
	4.84406305325125486048e3,
}








func Ft5ot3Jg(_BqkSz0 float64) float64 {
	if haveArchTanh {
		return  /*line :1*/ bXuNTLqDNAKR(_BqkSz0)
	}
	return  /*line :1*/ yrOj3f_(_BqkSz0)
}

func yrOj3f_(_BqkSz0 float64) float64 {
	const MAXLOG = 8.8029691931113054295988e+01	
	gZDpjWvXax_q :=  /*line :1*/ Abs(_BqkSz0)
	switch {
	case gZDpjWvXax_q > 0.5*MAXLOG:
		if _BqkSz0 < 0 {
			return -1
		}
		return 1
	case gZDpjWvXax_q >= 0.625:
		aK8LO1 :=  /*line :1*/ JPvpMNJa(2 * gZDpjWvXax_q)
		gZDpjWvXax_q = 1 - 2/(aK8LO1+1)
		if _BqkSz0 < 0 {
			gZDpjWvXax_q = -gZDpjWvXax_q
		}
	default:
		if _BqkSz0 == 0 {
			return _BqkSz0
		}
		aK8LO1 := _BqkSz0 * _BqkSz0
		gZDpjWvXax_q = _BqkSz0 + _BqkSz0*aK8LO1*((gOnNGf[0]*aK8LO1+gOnNGf[1])*aK8LO1+gOnNGf[2])/(((aK8LO1+iX5w2LUYwJ0[0])*aK8LO1+iX5w2LUYwJ0[1])*aK8LO1+iX5w2LUYwJ0[2])
	}
	return gZDpjWvXax_q
}
