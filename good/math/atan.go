//line :1
package math


func h4B8022_W(_BqkSz0 float64) float64 {
	const (
		P0	= -8.750608600031904122785e-01
		P1	= -1.615753718733365076637e+01
		P2	= -7.500855792314704667340e+01
		P3	= -1.228866684490136173410e+02
		P4	= -6.485021904942025371773e+01
		Q0	= +2.485846490142306297962e+01
		Q1	= +1.650270098316988542046e+02
		Q2	= +4.328810604912902668951e+02
		Q3	= +4.853903996359136964868e+02
		Q4	= +1.945506571482613964425e+02
	)
	gZDpjWvXax_q := _BqkSz0 * _BqkSz0
	gZDpjWvXax_q = gZDpjWvXax_q * ((((P0*gZDpjWvXax_q+P1)*gZDpjWvXax_q+P2)*gZDpjWvXax_q+P3)*gZDpjWvXax_q + P4) / (((((gZDpjWvXax_q+Q0)*gZDpjWvXax_q+Q1)*gZDpjWvXax_q+Q2)*gZDpjWvXax_q+Q3)*gZDpjWvXax_q + Q4)
	gZDpjWvXax_q = _BqkSz0*gZDpjWvXax_q + _BqkSz0
	return gZDpjWvXax_q
}



func vvqW1jt(_BqkSz0 float64) float64 {
	const (
		Morebits	= 6.123233995736765886130e-17		
		Tan3pio8	= 2.41421356237309504880		
	)
	if _BqkSz0 <= 0.66 {
		return  /*line :1*/ h4B8022_W(_BqkSz0)
	}
	if _BqkSz0 > Tan3pio8 {
		return Pi/2 -  /*line :1*/ h4B8022_W(1/_BqkSz0) + Morebits
	}
	return Pi/4 +  /*line :1*/ h4B8022_W((_BqkSz0-1)/(_BqkSz0+1)) + 0.5*Morebits
}







func DaNPs9X(_BqkSz0 float64) float64 {
	if haveArchAtan {
		return  /*line :1*/ oyVanXy9(_BqkSz0)
	}
	return  /*line :1*/ bXAWVB(_BqkSz0)
}

func bXAWVB(_BqkSz0 float64) float64 {
	if _BqkSz0 == 0 {
		return _BqkSz0
	}
	if _BqkSz0 > 0 {
		return  /*line :1*/ vvqW1jt(_BqkSz0)
	}
	return - /*line :1*/ vvqW1jt(-_BqkSz0)
}
