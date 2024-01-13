//line :1
package math










func JPvpMNJa(_BqkSz0 float64) float64 {
	if haveArchExp {
		return  /*line :1*/ hY3LRfw5im(_BqkSz0)
	}
	return  /*line :1*/ nF4u1Vnrit(_BqkSz0)
}

func nF4u1Vnrit(_BqkSz0 float64) float64 {
	const (
		Ln2Hi	= 6.93147180369123816490e-01
		Ln2Lo	= 1.90821492927058770002e-10
		Log2e	= 1.44269504088896338700e+00

		Overflow	= 7.09782712893383973096e+02
		Underflow	= -7.45133219101941108420e+02
		NearZero	= 1.0 / (1 << 28)		
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return _BqkSz0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, -1):
		return 0
	case _BqkSz0 > Overflow:
		return  /*line :1*/ FSug4WHZSBwz(1)
	case _BqkSz0 < Underflow:
		return 0
	case -NearZero < _BqkSz0 && _BqkSz0 < NearZero:
		return 1 + _BqkSz0
	}

	
	var duDPpte int
	switch {
	case _BqkSz0 < 0:
		duDPpte =  /*line :1*/ int(Log2e*_BqkSz0 - 0.5)
	case _BqkSz0 > 0:
		duDPpte =  /*line :1*/ int(Log2e*_BqkSz0 + 0.5)
	}
	j2BkoXTMn9D := _BqkSz0 -  /*line :1*/ float64(duDPpte)*Ln2Hi
	j0RLYji :=  /*line :1*/ float64(duDPpte) * Ln2Lo

	return  /*line :1*/ e6coO41I(j2BkoXTMn9D, j0RLYji, duDPpte)
}




func MvtEseQrDJwy(_BqkSz0 float64) float64 {
	if haveArchExp2 {
		return  /*line :1*/ zaiIIMMakF8(_BqkSz0)
	}
	return  /*line :1*/ qUxfhIu5Fg(_BqkSz0)
}

func qUxfhIu5Fg(_BqkSz0 float64) float64 {
	const (
		Ln2Hi	= 6.93147180369123816490e-01
		Ln2Lo	= 1.90821492927058770002e-10

		Overflow	= 1.0239999999999999e+03
		Underflow	= -1.0740e+03
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return _BqkSz0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, -1):
		return 0
	case _BqkSz0 > Overflow:
		return  /*line :1*/ FSug4WHZSBwz(1)
	case _BqkSz0 < Underflow:
		return 0
	}

	
	
	var duDPpte int
	switch {
	case _BqkSz0 > 0:
		duDPpte =  /*line :1*/ int(_BqkSz0 + 0.5)
	case _BqkSz0 < 0:
		duDPpte =  /*line :1*/ int(_BqkSz0 - 0.5)
	}
	aX6xkg_1G := _BqkSz0 -  /*line :1*/ float64(duDPpte)
	j2BkoXTMn9D := aX6xkg_1G * Ln2Hi
	j0RLYji := -aX6xkg_1G * Ln2Lo

	return  /*line :1*/ e6coO41I(j2BkoXTMn9D, j0RLYji, duDPpte)
}


func e6coO41I(j2BkoXTMn9D, j0RLYji float64, duDPpte int) float64 {
	const (
		P1	= 1.66666666666666657415e-01		
		P2	= -2.77777777770155933842e-03		
		P3	= 6.61375632143793436117e-05		
		P4	= -1.65339022054652515390e-06		
		P5	= 4.13813679705723846039e-08		
	)

	aBK_X3rbPc := j2BkoXTMn9D - j0RLYji
	aX6xkg_1G := aBK_X3rbPc * aBK_X3rbPc
	cU6qMJer46oH := aBK_X3rbPc - aX6xkg_1G*(P1+aX6xkg_1G*(P2+aX6xkg_1G*(P3+aX6xkg_1G*(P4+aX6xkg_1G*P5))))
	nNeKvlvK6 := 1 - ((j0RLYji - (aBK_X3rbPc*cU6qMJer46oH)/(2-cU6qMJer46oH)) - j2BkoXTMn9D)

	return  /*line :1*/ EU1SFL1cw(nNeKvlvK6, duDPpte)
}
