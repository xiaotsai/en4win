//line :1
package math

var u3SEufEJaHb = [...]float64{
	7.72156649015328655494e-02,
	3.22467033424113591611e-01,
	6.73523010531292681824e-02,
	2.05808084325167332806e-02,
	7.38555086081402883957e-03,
	2.89051383673415629091e-03,
	1.19270763183362067845e-03,
	5.10069792153511336608e-04,
	2.20862790713908385557e-04,
	1.08011567247583939954e-04,
	2.52144565451257326939e-05,
	4.48640949618915160150e-05,
}
var deS_lsNNME = [...]float64{
	1.0,
	1.39200533467621045958e+00,
	7.21935547567138069525e-01,
	1.71933865632803078993e-01,
	1.86459191715652901344e-02,
	7.77942496381893596434e-04,
	7.32668430744625636189e-06,
}
var cZ4jQYz33 = [...]float64{
	-7.72156649015328655494e-02,
	2.14982415960608852501e-01,
	3.25778796408930981787e-01,
	1.46350472652464452805e-01,
	2.66422703033638609560e-02,
	1.84028451407337715652e-03,
	3.19475326584100867617e-05,
}
var wQfXpT6 = [...]float64{
	4.83836122723810047042e-01,
	-1.47587722994593911752e-01,
	6.46249402391333854778e-02,
	-3.27885410759859649565e-02,
	1.79706750811820387126e-02,
	-1.03142241298341437450e-02,
	6.10053870246291332635e-03,
	-3.68452016781138256760e-03,
	2.25964780900612472250e-03,
	-1.40346469989232843813e-03,
	8.81081882437654011382e-04,
	-5.38595305356740546715e-04,
	3.15632070903625950361e-04,
	-3.12754168375120860518e-04,
	3.35529192635519073543e-04,
}
var uNsMp3 = [...]float64{
	-7.72156649015328655494e-02,
	6.32827064025093366517e-01,
	1.45492250137234768737e+00,
	9.77717527963372745603e-01,
	2.28963728064692451092e-01,
	1.33810918536787660377e-02,
}
var gbBT77DQI5t = [...]float64{
	1.0,
	2.45597793713041134822e+00,
	2.12848976379893395361e+00,
	7.69285150456672783825e-01,
	1.04222645593369134254e-01,
	3.21709242282423911810e-03,
}
var vlJaq8Rj6A = [...]float64{
	4.18938533204672725052e-01,
	8.33333333333329678849e-02,
	-2.77777777728775536470e-03,
	7.93650558643019558500e-04,
	-5.95187557450339963135e-04,
	8.36339918996282139126e-04,
	-1.63092934096575273989e-03,
}










func UDD855c(_BqkSz0 float64) (jnlna1p float64, awRlyItEG2gn int) {
	const (
		Ymin	= 1.461632144968362245
		Two52	= 1 << 52		
		Two53	= 1 << 53		
		Two58	= 1 << 58		
		Tiny	= 1.0 / (1 << 70)		
		Tc	= 1.46163214496836224576e+00		
		Tf	= -1.21486290535849611461e-01		
		
		Tt	= -3.63867699703950536541e-18		
	)

	awRlyItEG2gn = 1
	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		jnlna1p = _BqkSz0
		return
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		jnlna1p = _BqkSz0
		return
	case _BqkSz0 == 0:
		jnlna1p =  /*line :1*/ FSug4WHZSBwz(1)
		return
	}

	dp_kaI299 := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		dp_kaI299 = true
	}

	if _BqkSz0 < Tiny {
		if dp_kaI299 {
			awRlyItEG2gn = -1
		}
		jnlna1p = - /*line :1*/ JrJVPb5WbG(_BqkSz0)
		return
	}
	var u85pFpY0 float64
	if dp_kaI299 {
		if _BqkSz0 >= Two52 {
			jnlna1p =  /*line :1*/ FSug4WHZSBwz(1)
			return
		}
		aX6xkg_1G :=  /*line :1*/ kb1KxYaW(_BqkSz0)
		if aX6xkg_1G == 0 {
			jnlna1p =  /*line :1*/ FSug4WHZSBwz(1)
			return
		}
		u85pFpY0 =  /*line :1*/ JrJVPb5WbG(Pi /  /*line :1*/ Abs(aX6xkg_1G*_BqkSz0))
		if aX6xkg_1G < 0 {
			awRlyItEG2gn = -1
		}
	}

	switch {
	case _BqkSz0 == 1 || _BqkSz0 == 2:
		jnlna1p = 0
		return
	case _BqkSz0 < 2:
		var nNeKvlvK6 float64
		var jJklNl_1r int
		if _BqkSz0 <= 0.9 {
			jnlna1p = - /*line :1*/ JrJVPb5WbG(_BqkSz0)
			switch {
			case _BqkSz0 >= (Ymin - 1 + 0.27):
				nNeKvlvK6 = 1 - _BqkSz0
				jJklNl_1r = 0
			case _BqkSz0 >= (Ymin - 1 - 0.27):
				nNeKvlvK6 = _BqkSz0 - (Tc - 1)
				jJklNl_1r = 1
			default:
				nNeKvlvK6 = _BqkSz0
				jJklNl_1r = 2
			}
		} else {
			jnlna1p = 0
			switch {
			case _BqkSz0 >= (Ymin + 0.27):
				nNeKvlvK6 = 2 - _BqkSz0
				jJklNl_1r = 0
			case _BqkSz0 >= (Ymin - 0.27):
				nNeKvlvK6 = _BqkSz0 - Tc
				jJklNl_1r = 1
			default:
				nNeKvlvK6 = _BqkSz0 - 1
				jJklNl_1r = 2
			}
		}
		switch jJklNl_1r {
		case 0:
			gZDpjWvXax_q := nNeKvlvK6 * nNeKvlvK6
			mFEhRN := u3SEufEJaHb[0] + gZDpjWvXax_q*(u3SEufEJaHb[2]+gZDpjWvXax_q*(u3SEufEJaHb[4]+gZDpjWvXax_q*(u3SEufEJaHb[6]+gZDpjWvXax_q*(u3SEufEJaHb[8]+gZDpjWvXax_q*u3SEufEJaHb[10]))))
			iQwkj1f := gZDpjWvXax_q * (u3SEufEJaHb[1] + gZDpjWvXax_q*(+u3SEufEJaHb[3]+gZDpjWvXax_q*(u3SEufEJaHb[5]+gZDpjWvXax_q*(u3SEufEJaHb[7]+gZDpjWvXax_q*(u3SEufEJaHb[9]+gZDpjWvXax_q*u3SEufEJaHb[11])))))
			fzJsoqu := nNeKvlvK6*mFEhRN + iQwkj1f
			jnlna1p += (fzJsoqu - 0.5*nNeKvlvK6)
		case 1:
			gZDpjWvXax_q := nNeKvlvK6 * nNeKvlvK6
			uuSWcjr := gZDpjWvXax_q * nNeKvlvK6
			mFEhRN := wQfXpT6[0] + uuSWcjr*(wQfXpT6[3]+uuSWcjr*(wQfXpT6[6]+uuSWcjr*(wQfXpT6[9]+uuSWcjr*wQfXpT6[12])))
			iQwkj1f := wQfXpT6[1] + uuSWcjr*(wQfXpT6[4]+uuSWcjr*(wQfXpT6[7]+uuSWcjr*(wQfXpT6[10]+uuSWcjr*wQfXpT6[13])))
			jhgZ0uruoSf := wQfXpT6[2] + uuSWcjr*(wQfXpT6[5]+uuSWcjr*(wQfXpT6[8]+uuSWcjr*(wQfXpT6[11]+uuSWcjr*wQfXpT6[14])))
			fzJsoqu := gZDpjWvXax_q*mFEhRN - (Tt - uuSWcjr*(iQwkj1f+nNeKvlvK6*jhgZ0uruoSf))
			jnlna1p += (Tf + fzJsoqu)
		case 2:
			mFEhRN := nNeKvlvK6 * (uNsMp3[0] + nNeKvlvK6*(uNsMp3[1]+nNeKvlvK6*(uNsMp3[2]+nNeKvlvK6*(uNsMp3[3]+nNeKvlvK6*(uNsMp3[4]+nNeKvlvK6*uNsMp3[5])))))
			iQwkj1f := 1 + nNeKvlvK6*(gbBT77DQI5t[1]+nNeKvlvK6*(gbBT77DQI5t[2]+nNeKvlvK6*(gbBT77DQI5t[3]+nNeKvlvK6*(gbBT77DQI5t[4]+nNeKvlvK6*gbBT77DQI5t[5]))))
			jnlna1p += (-0.5*nNeKvlvK6 + mFEhRN/iQwkj1f)
		}
	case _BqkSz0 < 8:
		jJklNl_1r :=  /*line :1*/ int(_BqkSz0)
		nNeKvlvK6 := _BqkSz0 -  /*line :1*/ float64(jJklNl_1r)
		fzJsoqu := nNeKvlvK6 * (cZ4jQYz33[0] + nNeKvlvK6*(cZ4jQYz33[1]+nNeKvlvK6*(cZ4jQYz33[2]+nNeKvlvK6*(cZ4jQYz33[3]+nNeKvlvK6*(cZ4jQYz33[4]+nNeKvlvK6*(cZ4jQYz33[5]+nNeKvlvK6*cZ4jQYz33[6]))))))
		gNDglc3F2QV := 1 + nNeKvlvK6*(deS_lsNNME[1]+nNeKvlvK6*(deS_lsNNME[2]+nNeKvlvK6*(deS_lsNNME[3]+nNeKvlvK6*(deS_lsNNME[4]+nNeKvlvK6*(deS_lsNNME[5]+nNeKvlvK6*deS_lsNNME[6])))))
		jnlna1p = 0.5*nNeKvlvK6 + fzJsoqu/gNDglc3F2QV
		gZDpjWvXax_q := 1.0
		switch jJklNl_1r {
		case 7:
			gZDpjWvXax_q *= (nNeKvlvK6 + 6)
			fallthrough
		case 6:
			gZDpjWvXax_q *= (nNeKvlvK6 + 5)
			fallthrough
		case 5:
			gZDpjWvXax_q *= (nNeKvlvK6 + 4)
			fallthrough
		case 4:
			gZDpjWvXax_q *= (nNeKvlvK6 + 3)
			fallthrough
		case 3:
			gZDpjWvXax_q *= (nNeKvlvK6 + 2)
			jnlna1p +=  /*line :1*/ JrJVPb5WbG(gZDpjWvXax_q)
		}
	case _BqkSz0 < Two58:
		aX6xkg_1G :=  /*line :1*/ JrJVPb5WbG(_BqkSz0)
		gZDpjWvXax_q := 1 / _BqkSz0
		nNeKvlvK6 := gZDpjWvXax_q * gZDpjWvXax_q
		uuSWcjr := vlJaq8Rj6A[0] + gZDpjWvXax_q*(vlJaq8Rj6A[1]+nNeKvlvK6*(vlJaq8Rj6A[2]+nNeKvlvK6*(vlJaq8Rj6A[3]+nNeKvlvK6*(vlJaq8Rj6A[4]+nNeKvlvK6*(vlJaq8Rj6A[5]+nNeKvlvK6*vlJaq8Rj6A[6])))))
		jnlna1p = (_BqkSz0-0.5)*(aX6xkg_1G-1) + uuSWcjr
	default:
		jnlna1p = _BqkSz0 * ( /*line :1*/ JrJVPb5WbG(_BqkSz0) - 1)
	}
	if dp_kaI299 {
		jnlna1p = u85pFpY0 - jnlna1p
	}
	return
}


func kb1KxYaW(_BqkSz0 float64) float64 {
	const (
		Two52	= 1 << 52		
		Two53	= 1 << 53		
	)
	if _BqkSz0 < 0.25 {
		return - /*line :1*/ Cyl8FsLg(Pi * _BqkSz0)
	}

	gZDpjWvXax_q :=  /*line :1*/ Floor(_BqkSz0)
	var d2X2v9 int
	if gZDpjWvXax_q != _BqkSz0 {
		_BqkSz0 =  /*line :1*/ LoDnahP_eoog(_BqkSz0, 2)
		d2X2v9 =  /*line :1*/ int(_BqkSz0 * 4)
	} else {
		if _BqkSz0 >= Two53 {
			_BqkSz0 = 0
			d2X2v9 = 0
		} else {
			if _BqkSz0 < Two52 {
				gZDpjWvXax_q = _BqkSz0 + Two52
			}
			d2X2v9 =  /*line :1*/ int(1 &  /*line :1*/ GKIRmoHE(gZDpjWvXax_q))
			_BqkSz0 =  /*line :1*/ float64(d2X2v9)
			d2X2v9 <<= 2
		}
	}
	switch d2X2v9 {
	case 0:
		_BqkSz0 =  /*line :1*/ Cyl8FsLg(Pi * _BqkSz0)
	case 1, 2:
		_BqkSz0 =  /*line :1*/ Sr_3cKxech(Pi * (0.5 - _BqkSz0))
	case 3, 4:
		_BqkSz0 =  /*line :1*/ Cyl8FsLg(Pi * (1 - _BqkSz0))
	case 5, 6:
		_BqkSz0 = - /*line :1*/ Sr_3cKxech(Pi * (_BqkSz0 - 1.5))
	default:
		_BqkSz0 =  /*line :1*/ Cyl8FsLg(Pi * (_BqkSz0 - 2))
	}
	return -_BqkSz0
}
