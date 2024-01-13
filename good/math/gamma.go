//line :1
package math

var dhX83oGHz = [...]float64{
	1.60119522476751861407e-04,
	1.19135147006586384913e-03,
	1.04213797561761569935e-02,
	4.76367800457137231464e-02,
	2.07448227648435975150e-01,
	4.94214826801497100753e-01,
	9.99999999999999996796e-01,
}
var j6W4KEo3lE44 = [...]float64{
	-2.31581873324120129819e-05,
	5.39605580493303397842e-04,
	-4.45641913851797240494e-03,
	1.18139785222060435552e-02,
	3.58236398605498653373e-02,
	-2.34591795718243348568e-01,
	7.14304917030273074085e-02,
	1.00000000000000000320e+00,
}
var hRbHUE0mzykn = [...]float64{
	7.87311395793093628397e-04,
	-2.29549961613378126380e-04,
	-2.68132617805781232825e-03,
	3.47222221605458667310e-03,
	8.33333333333482257126e-02,
}








func w2G3Fiab(_BqkSz0 float64) (float64, float64) {
	if _BqkSz0 > 200 {
		return  /*line :1*/ FSug4WHZSBwz(1), 1
	}
	const (
		SqrtTwoPi	= 2.506628274631000502417
		MaxStirling	= 143.01608
	)
	uuSWcjr := 1 / _BqkSz0
	uuSWcjr = 1 + uuSWcjr*((((hRbHUE0mzykn[0]*uuSWcjr+hRbHUE0mzykn[1])*uuSWcjr+hRbHUE0mzykn[2])*uuSWcjr+hRbHUE0mzykn[3])*uuSWcjr+hRbHUE0mzykn[4])
	nzcaAWDqtT3 :=  /*line :1*/ JPvpMNJa(_BqkSz0)
	zgrTNETCmq := 1.0
	if _BqkSz0 > MaxStirling {
		hLZMcJsdS :=  /*line :1*/ CPKyEBb3t3ps(_BqkSz0, 0.5*_BqkSz0-0.25)
		nzcaAWDqtT3, zgrTNETCmq = hLZMcJsdS, hLZMcJsdS/nzcaAWDqtT3
	} else {
		nzcaAWDqtT3 =  /*line :1*/ CPKyEBb3t3ps(_BqkSz0, _BqkSz0-0.5) / nzcaAWDqtT3
	}
	return nzcaAWDqtT3, SqrtTwoPi * uuSWcjr * zgrTNETCmq
}











func IjaFhc8d(_BqkSz0 float64) float64 {
	const Euler = 0.57721566490153286060651209008240243104215933593992	

	switch {
	case  /*line :1*/ nLKbL8A7ERyi(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, -1) ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return  /*line :1*/ FSug4WHZSBwz(1)
	case _BqkSz0 == 0:
		if  /*line :1*/ GLZFJA9L9(_BqkSz0) {
			return  /*line :1*/ FSug4WHZSBwz(-1)
		}
		return  /*line :1*/ FSug4WHZSBwz(1)
	}
	gNDglc3F2QV :=  /*line :1*/ Abs(_BqkSz0)
	fzJsoqu :=  /*line :1*/ Floor(gNDglc3F2QV)
	if gNDglc3F2QV > 33 {
		if _BqkSz0 >= 0 {
			nzcaAWDqtT3, zgrTNETCmq :=  /*line :1*/ w2G3Fiab(_BqkSz0)
			return nzcaAWDqtT3 * zgrTNETCmq
		}

		inNQdqKvjxOy := 1
		if byrSb4 :=  /*line :1*/ int64(fzJsoqu); byrSb4&1 == 0 {
			inNQdqKvjxOy = -1
		}
		gZDpjWvXax_q := gNDglc3F2QV - fzJsoqu
		if gZDpjWvXax_q > 0.5 {
			fzJsoqu = fzJsoqu + 1
			gZDpjWvXax_q = gNDglc3F2QV - fzJsoqu
		}
		gZDpjWvXax_q = gNDglc3F2QV *  /*line :1*/ Cyl8FsLg(Pi*gZDpjWvXax_q)
		if gZDpjWvXax_q == 0 {
			return  /*line :1*/ FSug4WHZSBwz(inNQdqKvjxOy)
		}
		qxnpLxi, gBfFce :=  /*line :1*/ w2G3Fiab(gNDglc3F2QV)
		qRvcahYW :=  /*line :1*/ Abs(gZDpjWvXax_q)
		pjjMpWihfSot := qRvcahYW * qxnpLxi * gBfFce
		if  /*line :1*/ ME1vzpD5wcr(pjjMpWihfSot, 0) {
			gZDpjWvXax_q = Pi / qRvcahYW / qxnpLxi / gBfFce
		} else {
			gZDpjWvXax_q = Pi / pjjMpWihfSot
		}
		return  /*line :1*/ float64(inNQdqKvjxOy) * gZDpjWvXax_q
	}

	gZDpjWvXax_q := 1.0
	for _BqkSz0 >= 3 {
		_BqkSz0 = _BqkSz0 - 1
		gZDpjWvXax_q = gZDpjWvXax_q * _BqkSz0
	}
	for _BqkSz0 < 0 {
		if _BqkSz0 > -1e-09 {
			goto small
		}
		gZDpjWvXax_q = gZDpjWvXax_q / _BqkSz0
		_BqkSz0 = _BqkSz0 + 1
	}
	for _BqkSz0 < 2 {
		if _BqkSz0 < 1e-09 {
			goto small
		}
		gZDpjWvXax_q = gZDpjWvXax_q / _BqkSz0
		_BqkSz0 = _BqkSz0 + 1
	}

	if _BqkSz0 == 2 {
		return gZDpjWvXax_q
	}

	_BqkSz0 = _BqkSz0 - 2
	fzJsoqu = (((((_BqkSz0*dhX83oGHz[0]+dhX83oGHz[1])*_BqkSz0+dhX83oGHz[2])*_BqkSz0+dhX83oGHz[3])*_BqkSz0+dhX83oGHz[4])*_BqkSz0+dhX83oGHz[5])*_BqkSz0 + dhX83oGHz[6]
	gNDglc3F2QV = ((((((_BqkSz0*j6W4KEo3lE44[0]+j6W4KEo3lE44[1])*_BqkSz0+j6W4KEo3lE44[2])*_BqkSz0+j6W4KEo3lE44[3])*_BqkSz0+j6W4KEo3lE44[4])*_BqkSz0+j6W4KEo3lE44[5])*_BqkSz0+j6W4KEo3lE44[6])*_BqkSz0 + j6W4KEo3lE44[7]
	return gZDpjWvXax_q * fzJsoqu / gNDglc3F2QV

small:
	if _BqkSz0 == 0 {
		return  /*line :1*/ FSug4WHZSBwz(1)
	}
	return gZDpjWvXax_q / ((1 + Euler*_BqkSz0) * _BqkSz0)
}

func nLKbL8A7ERyi(_BqkSz0 float64) bool {
	if _BqkSz0 < 0 {
		_, wu_h2B5aU5A :=  /*line :1*/ HJa8famXXqZ(_BqkSz0)
		return wu_h2B5aU5A == 0
	}
	return false
}
