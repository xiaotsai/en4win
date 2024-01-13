//line :1
package math









func JrJVPb5WbG(_BqkSz0 float64) float64 {
	if haveArchLog {
		return  /*line :1*/ t7Na14GYVS(_BqkSz0)
	}
	return  /*line :1*/ tCZZSiFlFSTH(_BqkSz0)
}

func tCZZSiFlFSTH(_BqkSz0 float64) float64 {
	const (
		Ln2Hi	= 6.93147180369123816490e-01		
		Ln2Lo	= 1.90821492927058770002e-10		
		L1	= 6.666666666666735130e-01		
		L2	= 3.999999999940941908e-01		
		L3	= 2.857142874366239149e-01		
		L4	= 2.222219843214978396e-01		
		L5	= 1.818357216161805012e-01		
		L6	= 1.531383769920937332e-01		
		L7	= 1.479819860511658591e-01		
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return _BqkSz0
	case _BqkSz0 < 0:
		return  /*line :1*/ WG0hZIT4()
	case _BqkSz0 == 0:
		return  /*line :1*/ FSug4WHZSBwz(-1)
	}

	x1ZAMNfRGK3, wmBPHu :=  /*line :1*/ CuUl2eV(_BqkSz0)
	if x1ZAMNfRGK3 < Sqrt2/2 {
		x1ZAMNfRGK3 *= 2
		wmBPHu--
	}
	ygoarCBO4et := x1ZAMNfRGK3 - 1
	duDPpte :=  /*line :1*/ float64(wmBPHu)

	aK8LO1 := ygoarCBO4et / (2 + ygoarCBO4et)
	c0W5O46wS := aK8LO1 * aK8LO1
	aKVyc6p6tLBh := c0W5O46wS * c0W5O46wS
	qGY_LtJZ := c0W5O46wS * (L1 + aKVyc6p6tLBh*(L3+aKVyc6p6tLBh*(L5+aKVyc6p6tLBh*L7)))
	baU5Z449bMlW := aKVyc6p6tLBh * (L2 + aKVyc6p6tLBh*(L4+aKVyc6p6tLBh*L6))
	HEiTyNyOi := qGY_LtJZ + baU5Z449bMlW
	sZYhMItCeB := 0.5 * ygoarCBO4et * ygoarCBO4et
	return duDPpte*Ln2Hi - ((sZYhMItCeB - (aK8LO1*(sZYhMItCeB+HEiTyNyOi) + duDPpte*Ln2Lo)) - ygoarCBO4et)
}
