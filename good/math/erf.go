//line :1
package math

const (
	erx	= 8.45062911510467529297e-01		
	
	efx	= 1.28379167095512586316e-01		
	efx8	= 1.02703333676410069053e+00		
	pp0	= 1.28379167095512558561e-01		
	pp1	= -3.25042107247001499370e-01		
	pp2	= -2.84817495755985104766e-02		
	pp3	= -5.77027029648944159157e-03		
	pp4	= -2.37630166566501626084e-05		
	qq1	= 3.97917223959155352819e-01		
	qq2	= 6.50222499887672944485e-02		
	qq3	= 5.08130628187576562776e-03		
	qq4	= 1.32494738004321644526e-04		
	qq5	= -3.96022827877536812320e-06		
	
	pa0	= -2.36211856075265944077e-03		
	pa1	= 4.14856118683748331666e-01		
	pa2	= -3.72207876035701323847e-01		
	pa3	= 3.18346619901161753674e-01		
	pa4	= -1.10894694282396677476e-01		
	pa5	= 3.54783043256182359371e-02		
	pa6	= -2.16637559486879084300e-03		
	qa1	= 1.06420880400844228286e-01		
	qa2	= 5.40397917702171048937e-01		
	qa3	= 7.18286544141962662868e-02		
	qa4	= 1.26171219808761642112e-01		
	qa5	= 1.36370839120290507362e-02		
	qa6	= 1.19844998467991074170e-02		
	
	ra0	= -9.86494403484714822705e-03		
	ra1	= -6.93858572707181764372e-01		
	ra2	= -1.05586262253232909814e+01		
	ra3	= -6.23753324503260060396e+01		
	ra4	= -1.62396669462573470355e+02		
	ra5	= -1.84605092906711035994e+02		
	ra6	= -8.12874355063065934246e+01		
	ra7	= -9.81432934416914548592e+00		
	sa1	= 1.96512716674392571292e+01		
	sa2	= 1.37657754143519042600e+02		
	sa3	= 4.34565877475229228821e+02		
	sa4	= 6.45387271733267880336e+02		
	sa5	= 4.29008140027567833386e+02		
	sa6	= 1.08635005541779435134e+02		
	sa7	= 6.57024977031928170135e+00		
	sa8	= -6.04244152148580987438e-02		
	
	rb0	= -9.86494292470009928597e-03		
	rb1	= -7.99283237680523006574e-01		
	rb2	= -1.77579549177547519889e+01		
	rb3	= -1.60636384855821916062e+02		
	rb4	= -6.37566443368389627722e+02		
	rb5	= -1.02509513161107724954e+03		
	rb6	= -4.83519191608651397019e+02		
	sb1	= 3.03380607434824582924e+01		
	sb2	= 3.25792512996573918826e+02		
	sb3	= 1.53672958608443695994e+03		
	sb4	= 3.19985821950859553908e+03		
	sb5	= 2.55305040643316442583e+03		
	sb6	= 4.74528541206955367215e+02		
	sb7	= -2.24409524465858183362e+01		
)








func A5zD8lfP3r6(_BqkSz0 float64) float64 {
	if haveArchErf {
		return  /*line :1*/ bph78u5ovei(_BqkSz0)
	}
	return  /*line :1*/ psL_tdL(_BqkSz0)
}

func psL_tdL(_BqkSz0 float64) float64 {
	const (
		VeryTiny	= 2.848094538889218e-306		
		Small	= 1.0 / (1 << 28)		
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return 1
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, -1):
		return -1
	}
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	if _BqkSz0 < 0.84375 {
		var dE6bQU3 float64
		if _BqkSz0 < Small {
			if _BqkSz0 < VeryTiny {
				dE6bQU3 = 0.125 * (8.0*_BqkSz0 + efx8*_BqkSz0)
			} else {
				dE6bQU3 = _BqkSz0 + efx*_BqkSz0
			}
		} else {
			gZDpjWvXax_q := _BqkSz0 * _BqkSz0
			aBK_X3rbPc := pp0 + gZDpjWvXax_q*(pp1+gZDpjWvXax_q*(pp2+gZDpjWvXax_q*(pp3+gZDpjWvXax_q*pp4)))
			aK8LO1 := 1 + gZDpjWvXax_q*(qq1+gZDpjWvXax_q*(qq2+gZDpjWvXax_q*(qq3+gZDpjWvXax_q*(qq4+gZDpjWvXax_q*qq5))))
			nNeKvlvK6 := aBK_X3rbPc / aK8LO1
			dE6bQU3 = _BqkSz0 + _BqkSz0*nNeKvlvK6
		}
		if awRlyItEG2gn {
			return -dE6bQU3
		}
		return dE6bQU3
	}
	if _BqkSz0 < 1.25 {
		aK8LO1 := _BqkSz0 - 1
		YKa188RBgV7t := pa0 + aK8LO1*(pa1+aK8LO1*(pa2+aK8LO1*(pa3+aK8LO1*(pa4+aK8LO1*(pa5+aK8LO1*pa6)))))
		ODJt_PG := 1 + aK8LO1*(qa1+aK8LO1*(qa2+aK8LO1*(qa3+aK8LO1*(qa4+aK8LO1*(qa5+aK8LO1*qa6)))))
		if awRlyItEG2gn {
			return -erx - YKa188RBgV7t/ODJt_PG
		}
		return erx + YKa188RBgV7t/ODJt_PG
	}
	if _BqkSz0 >= 6 {
		if awRlyItEG2gn {
			return -1
		}
		return 1
	}
	aK8LO1 := 1 / (_BqkSz0 * _BqkSz0)
	var HEiTyNyOi, DDBrTaUra2_6 float64
	if _BqkSz0 < 1/0.35 {
		HEiTyNyOi = ra0 + aK8LO1*(ra1+aK8LO1*(ra2+aK8LO1*(ra3+aK8LO1*(ra4+aK8LO1*(ra5+aK8LO1*(ra6+aK8LO1*ra7))))))
		DDBrTaUra2_6 = 1 + aK8LO1*(sa1+aK8LO1*(sa2+aK8LO1*(sa3+aK8LO1*(sa4+aK8LO1*(sa5+aK8LO1*(sa6+aK8LO1*(sa7+aK8LO1*sa8)))))))
	} else {
		HEiTyNyOi = rb0 + aK8LO1*(rb1+aK8LO1*(rb2+aK8LO1*(rb3+aK8LO1*(rb4+aK8LO1*(rb5+aK8LO1*rb6)))))
		DDBrTaUra2_6 = 1 + aK8LO1*(sb1+aK8LO1*(sb2+aK8LO1*(sb3+aK8LO1*(sb4+aK8LO1*(sb5+aK8LO1*(sb6+aK8LO1*sb7))))))
	}
	gZDpjWvXax_q :=  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(_BqkSz0) & 0xffffffff00000000)
	aBK_X3rbPc :=  /*line :1*/ JPvpMNJa(-gZDpjWvXax_q*gZDpjWvXax_q-0.5625) *  /*line :1*/ JPvpMNJa((gZDpjWvXax_q-_BqkSz0)*(gZDpjWvXax_q+_BqkSz0)+HEiTyNyOi/DDBrTaUra2_6)
	if awRlyItEG2gn {
		return aBK_X3rbPc/_BqkSz0 - 1
	}
	return 1 - aBK_X3rbPc/_BqkSz0
}








func MCv2jG7SFTVO(_BqkSz0 float64) float64 {
	if haveArchErfc {
		return  /*line :1*/ _wJxow6j8LQB(_BqkSz0)
	}
	return  /*line :1*/ xazlwdHDb(_BqkSz0)
}

func xazlwdHDb(_BqkSz0 float64) float64 {
	const Tiny = 1.0 / (1 << 56)	

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return 0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, -1):
		return 2
	}
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	if _BqkSz0 < 0.84375 {
		var dE6bQU3 float64
		if _BqkSz0 < Tiny {
			dE6bQU3 = _BqkSz0
		} else {
			gZDpjWvXax_q := _BqkSz0 * _BqkSz0
			aBK_X3rbPc := pp0 + gZDpjWvXax_q*(pp1+gZDpjWvXax_q*(pp2+gZDpjWvXax_q*(pp3+gZDpjWvXax_q*pp4)))
			aK8LO1 := 1 + gZDpjWvXax_q*(qq1+gZDpjWvXax_q*(qq2+gZDpjWvXax_q*(qq3+gZDpjWvXax_q*(qq4+gZDpjWvXax_q*qq5))))
			nNeKvlvK6 := aBK_X3rbPc / aK8LO1
			if _BqkSz0 < 0.25 {
				dE6bQU3 = _BqkSz0 + _BqkSz0*nNeKvlvK6
			} else {
				dE6bQU3 = 0.5 + (_BqkSz0*nNeKvlvK6 + (_BqkSz0 - 0.5))
			}
		}
		if awRlyItEG2gn {
			return 1 + dE6bQU3
		}
		return 1 - dE6bQU3
	}
	if _BqkSz0 < 1.25 {
		aK8LO1 := _BqkSz0 - 1
		YKa188RBgV7t := pa0 + aK8LO1*(pa1+aK8LO1*(pa2+aK8LO1*(pa3+aK8LO1*(pa4+aK8LO1*(pa5+aK8LO1*pa6)))))
		ODJt_PG := 1 + aK8LO1*(qa1+aK8LO1*(qa2+aK8LO1*(qa3+aK8LO1*(qa4+aK8LO1*(qa5+aK8LO1*qa6)))))
		if awRlyItEG2gn {
			return 1 + erx + YKa188RBgV7t/ODJt_PG
		}
		return 1 - erx - YKa188RBgV7t/ODJt_PG

	}
	if _BqkSz0 < 28 {
		aK8LO1 := 1 / (_BqkSz0 * _BqkSz0)
		var HEiTyNyOi, DDBrTaUra2_6 float64
		if _BqkSz0 < 1/0.35 {
			HEiTyNyOi = ra0 + aK8LO1*(ra1+aK8LO1*(ra2+aK8LO1*(ra3+aK8LO1*(ra4+aK8LO1*(ra5+aK8LO1*(ra6+aK8LO1*ra7))))))
			DDBrTaUra2_6 = 1 + aK8LO1*(sa1+aK8LO1*(sa2+aK8LO1*(sa3+aK8LO1*(sa4+aK8LO1*(sa5+aK8LO1*(sa6+aK8LO1*(sa7+aK8LO1*sa8)))))))
		} else {
			if awRlyItEG2gn && _BqkSz0 > 6 {
				return 2
			}
			HEiTyNyOi = rb0 + aK8LO1*(rb1+aK8LO1*(rb2+aK8LO1*(rb3+aK8LO1*(rb4+aK8LO1*(rb5+aK8LO1*rb6)))))
			DDBrTaUra2_6 = 1 + aK8LO1*(sb1+aK8LO1*(sb2+aK8LO1*(sb3+aK8LO1*(sb4+aK8LO1*(sb5+aK8LO1*(sb6+aK8LO1*sb7))))))
		}
		gZDpjWvXax_q :=  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(_BqkSz0) & 0xffffffff00000000)
		aBK_X3rbPc :=  /*line :1*/ JPvpMNJa(-gZDpjWvXax_q*gZDpjWvXax_q-0.5625) *  /*line :1*/ JPvpMNJa((gZDpjWvXax_q-_BqkSz0)*(gZDpjWvXax_q+_BqkSz0)+HEiTyNyOi/DDBrTaUra2_6)
		if awRlyItEG2gn {
			return 2 - aBK_X3rbPc/_BqkSz0
		}
		return aBK_X3rbPc / _BqkSz0
	}
	if awRlyItEG2gn {
		return 2
	}
	return 0
}
