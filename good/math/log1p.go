//line :1
package math











func V_uk99t0rhQ(_BqkSz0 float64) float64 {
	if haveArchLog1p {
		return  /*line :1*/ dqLFyFj(_BqkSz0)
	}
	return  /*line :1*/ dwbAYih(_BqkSz0)
}

func dwbAYih(_BqkSz0 float64) float64 {
	const (
		Sqrt2M1	= 4.142135623730950488017e-01		
		Sqrt2HalfM1	= -2.928932188134524755992e-01		
		Small	= 1.0 / (1 << 29)		
		Tiny	= 1.0 / (1 << 54)		
		Two53	= 1 << 53		
		Ln2Hi	= 6.93147180369123816490e-01		
		Ln2Lo	= 1.90821492927058770002e-10		
		Lp1	= 6.666666666666735130e-01		
		Lp2	= 3.999999999940941908e-01		
		Lp3	= 2.857142874366239149e-01		
		Lp4	= 2.222219843214978396e-01		
		Lp5	= 1.818357216161805012e-01		
		Lp6	= 1.531383769920937332e-01		
		Lp7	= 1.479819860511658591e-01		
	)

	switch {
	case _BqkSz0 < -1 ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case _BqkSz0 == -1:
		return  /*line :1*/ FSug4WHZSBwz(-1)
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return  /*line :1*/ FSug4WHZSBwz(1)
	}

	dU7v_ec8cB :=  /*line :1*/ Abs(_BqkSz0)

	var ygoarCBO4et float64
	var u46ljcm uint64
	duDPpte := 1
	if dU7v_ec8cB < Sqrt2M1 {
		if dU7v_ec8cB < Small {
			if dU7v_ec8cB < Tiny {
				return _BqkSz0
			}
			return _BqkSz0 - _BqkSz0*_BqkSz0*0.5
		}
		if _BqkSz0 > Sqrt2HalfM1 {

			duDPpte = 0
			ygoarCBO4et = _BqkSz0
			u46ljcm = 1
		}
	}
	var cU6qMJer46oH float64
	if duDPpte != 0 {
		var wkIoLysGG8KY float64
		if dU7v_ec8cB < Two53 {
			wkIoLysGG8KY = 1.0 + _BqkSz0
			u46ljcm =  /*line :1*/ GKIRmoHE(wkIoLysGG8KY)
			duDPpte =  /*line :1*/ int((u46ljcm >> 52) - 1023)

			if duDPpte > 0 {
				cU6qMJer46oH = 1.0 - (wkIoLysGG8KY - _BqkSz0)
			} else {
				cU6qMJer46oH = _BqkSz0 - (wkIoLysGG8KY - 1.0)
			}
			cU6qMJer46oH /= wkIoLysGG8KY
		} else {
			wkIoLysGG8KY = _BqkSz0
			u46ljcm =  /*line :1*/ GKIRmoHE(wkIoLysGG8KY)
			duDPpte =  /*line :1*/ int((u46ljcm >> 52) - 1023)
			cU6qMJer46oH = 0
		}
		u46ljcm &= 0x000fffffffffffff
		if u46ljcm < 0x0006a09e667f3bcd {
			wkIoLysGG8KY =  /*line :1*/ QUB2Kf63p(u46ljcm | 0x3ff0000000000000)
		} else {
			duDPpte++
			wkIoLysGG8KY =  /*line :1*/ QUB2Kf63p(u46ljcm | 0x3fe0000000000000)
			u46ljcm = (0x0010000000000000 - u46ljcm) >> 2
		}
		ygoarCBO4et = wkIoLysGG8KY - 1.0
	}
	sZYhMItCeB := 0.5 * ygoarCBO4et * ygoarCBO4et
	var aK8LO1, HEiTyNyOi, gZDpjWvXax_q float64
	if u46ljcm == 0 {
		if ygoarCBO4et == 0 {
			if duDPpte == 0 {
				return 0
			}
			cU6qMJer46oH +=  /*line :1*/ float64(duDPpte) * Ln2Lo
			return  /*line :1*/ float64(duDPpte)*Ln2Hi + cU6qMJer46oH
		}
		HEiTyNyOi = sZYhMItCeB * (1.0 - 0.66666666666666666*ygoarCBO4et)
		if duDPpte == 0 {
			return ygoarCBO4et - HEiTyNyOi
		}
		return  /*line :1*/ float64(duDPpte)*Ln2Hi - ((HEiTyNyOi - ( /*line :1*/ float64(duDPpte)*Ln2Lo + cU6qMJer46oH)) - ygoarCBO4et)
	}
	aK8LO1 = ygoarCBO4et / (2.0 + ygoarCBO4et)
	gZDpjWvXax_q = aK8LO1 * aK8LO1
	HEiTyNyOi = gZDpjWvXax_q * (Lp1 + gZDpjWvXax_q*(Lp2+gZDpjWvXax_q*(Lp3+gZDpjWvXax_q*(Lp4+gZDpjWvXax_q*(Lp5+gZDpjWvXax_q*(Lp6+gZDpjWvXax_q*Lp7))))))
	if duDPpte == 0 {
		return ygoarCBO4et - (sZYhMItCeB - aK8LO1*(sZYhMItCeB+HEiTyNyOi))
	}
	return  /*line :1*/ float64(duDPpte)*Ln2Hi - ((sZYhMItCeB - (aK8LO1*(sZYhMItCeB+HEiTyNyOi) + ( /*line :1*/ float64(duDPpte)*Ln2Lo + cU6qMJer46oH))) - ygoarCBO4et)
}
