//line :1
package big

import bits "math/bits"



func (uW0eIVjy g3X9fa) xf9ev_MB(sOU3zRv, iQKtjmO6 g3X9fa) (vFx5p_cm g3X9fa) {
	if  /*line :1*/ ap52A4djK(uW0eIVjy, sOU3zRv) {
		uW0eIVjy = nil
	}
	trfBtovmu :=  /*line :1*/ aQEWrUmnl5I(0)
	yL0_p0wUnc, vFx5p_cm :=  /*line :1*/ trfBtovmu.xOZxoLyl(uW0eIVjy, sOU3zRv, iQKtjmO6)
	*trfBtovmu = yL0_p0wUnc
	 /*line :1*/ zuavpO7v(trfBtovmu)
	return vFx5p_cm
}



func (uW0eIVjy g3X9fa) xOZxoLyl(aDGsAqP, sOU3zRv, iQKtjmO6 g3X9fa) (yL0_p0wUnc, vFx5p_cm g3X9fa) {
	if  /*line :1*/ len(iQKtjmO6) == 0 {
		 /*line :1*/ panic("division by zero")
	}

	if  /*line :1*/ sOU3zRv.beaurQHiT(iQKtjmO6) < 0 {
		yL0_p0wUnc = uW0eIVjy[:0]
		vFx5p_cm =  /*line :1*/ aDGsAqP.bj0nVC(sOU3zRv)
		return
	}

	if  /*line :1*/ len(iQKtjmO6) == 1 {
		
		
		var jt6wZjdRTx VYe6D0
		yL0_p0wUnc, jt6wZjdRTx =  /*line :1*/ uW0eIVjy.htGqJRPj(sOU3zRv, iQKtjmO6[0])
		vFx5p_cm =  /*line :1*/ aDGsAqP.pSzoQ7PMW(jt6wZjdRTx)
		return
	}

	yL0_p0wUnc, vFx5p_cm =  /*line :1*/ uW0eIVjy.qJEP6zbhN2(aDGsAqP, sOU3zRv, iQKtjmO6)
	return
}




func (uW0eIVjy g3X9fa) htGqJRPj(pmEgen2 g3X9fa, oxFS5wa5 VYe6D0) (yL0_p0wUnc g3X9fa, vFx5p_cm VYe6D0) {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	switch {
	case oxFS5wa5 == 0:
		 /*line :1*/ panic("division by zero")
	case oxFS5wa5 == 1:
		yL0_p0wUnc =  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
		return
	case gS2TXhHpYaX4 == 0:
		yL0_p0wUnc = uW0eIVjy[:0]
		return
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4)
	vFx5p_cm =  /*line :1*/ yR7lmNNg(uW0eIVjy, 0, pmEgen2, oxFS5wa5)
	yL0_p0wUnc =  /*line :1*/ uW0eIVjy.hitXcy5()
	return
}


func (pmEgen2 g3X9fa) tPll2cuYqvU(s31g8J VYe6D0) (vFx5p_cm VYe6D0) {
	
	var yL0_p0wUnc g3X9fa
	yL0_p0wUnc =  /*line :1*/ yL0_p0wUnc.gS9bqHpKAJ0( /*line :1*/ len(pmEgen2))
	return  /*line :1*/ yR7lmNNg(yL0_p0wUnc, 0, pmEgen2, s31g8J)
}



func yR7lmNNg(uW0eIVjy []VYe6D0, aYTwE5hn VYe6D0, pmEgen2 []VYe6D0, oxFS5wa5 VYe6D0) (vFx5p_cm VYe6D0) {
	vFx5p_cm = aYTwE5hn
	if  /*line :1*/ len(pmEgen2) == 1 {
		eqTwRvCQS, mdBq1Cc :=  /*line :1*/ bits.Div( /*line :1*/ uint(vFx5p_cm),  /*line :1*/ uint(pmEgen2[0]),  /*line :1*/ uint(oxFS5wa5))
		uW0eIVjy[0] =  /*line :1*/ VYe6D0(eqTwRvCQS)
		return  /*line :1*/ VYe6D0(mdBq1Cc)
	}
	jaaWqY :=  /*line :1*/ niNOGPaIxagq(oxFS5wa5)
	for aepbxLiWOZ :=  /*line :1*/ len(uW0eIVjy) - 1; aepbxLiWOZ >= 0; aepbxLiWOZ-- {
		uW0eIVjy[aepbxLiWOZ], vFx5p_cm =  /*line :1*/ f6uWqUJ(vFx5p_cm, pmEgen2[aepbxLiWOZ], oxFS5wa5, jaaWqY)
	}
	return vFx5p_cm
}





func (uW0eIVjy g3X9fa) qJEP6zbhN2(sOU3zRv, kGsXFL3Zr5e, aVlEd_Ko g3X9fa) (yL0_p0wUnc, vFx5p_cm g3X9fa) {
	h_Wr_yqq :=  /*line :1*/ len(aVlEd_Ko)
	gS2TXhHpYaX4 :=  /*line :1*/ len(kGsXFL3Zr5e) - h_Wr_yqq

	qKn7aIMh5 :=  /*line :1*/ bUgMSHLa(aVlEd_Ko[h_Wr_yqq-1])
	xFl7eTcfa7wT :=  /*line :1*/ aQEWrUmnl5I(h_Wr_yqq)
	iQKtjmO6 := *xFl7eTcfa7wT
	 /*line :1*/ vdi3DYM4mY(iQKtjmO6, aVlEd_Ko, qKn7aIMh5)
	sOU3zRv =  /*line :1*/ sOU3zRv.gS9bqHpKAJ0( /*line :1*/ len(kGsXFL3Zr5e) + 1)
	sOU3zRv[ /*line :1*/ len(kGsXFL3Zr5e)] =  /*line :1*/ vdi3DYM4mY(sOU3zRv[0: /*line :1*/ len(kGsXFL3Zr5e)], kGsXFL3Zr5e, qKn7aIMh5)

	if  /*line :1*/ ap52A4djK(uW0eIVjy, sOU3zRv) {
		uW0eIVjy = nil
	}
	yL0_p0wUnc =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4 + 1)

	if h_Wr_yqq < divRecursiveThreshold {
		 /*line :1*/ yL0_p0wUnc.dSb9JW6qvq(sOU3zRv, iQKtjmO6)
	} else {
		 /*line :1*/ yL0_p0wUnc.gCGOsVPsJ3(sOU3zRv, iQKtjmO6)
	}
	 /*line :1*/ zuavpO7v(xFl7eTcfa7wT)

	yL0_p0wUnc =  /*line :1*/ yL0_p0wUnc.hitXcy5()

	 /*line :1*/ t_ku9iRroNp(sOU3zRv, sOU3zRv, qKn7aIMh5)
	vFx5p_cm =  /*line :1*/ sOU3zRv.hitXcy5()

	return yL0_p0wUnc, vFx5p_cm
}




func (yL0_p0wUnc g3X9fa) dSb9JW6qvq(sOU3zRv, iQKtjmO6 g3X9fa) {
	h_Wr_yqq :=  /*line :1*/ len(iQKtjmO6)
	gS2TXhHpYaX4 :=  /*line :1*/ len(sOU3zRv) - h_Wr_yqq

	o9KrWrdIsv :=  /*line :1*/ aQEWrUmnl5I(h_Wr_yqq + 1)
	tu0jQu := *o9KrWrdIsv

	txT77f := iQKtjmO6[h_Wr_yqq-1]
	jaaWqY :=  /*line :1*/ niNOGPaIxagq(txT77f)

	for g77TAi := gS2TXhHpYaX4; g77TAi >= 0; g77TAi-- {

		xPSdytG4S :=  /*line :1*/ VYe6D0(_M)
		var p76FLbX6YHIC VYe6D0
		if g77TAi+h_Wr_yqq <  /*line :1*/ len(sOU3zRv) {
			p76FLbX6YHIC = sOU3zRv[g77TAi+h_Wr_yqq]
		}

		if p76FLbX6YHIC != txT77f {
			var y5RYzdlV VYe6D0
			xPSdytG4S, y5RYzdlV =  /*line :1*/ f6uWqUJ(p76FLbX6YHIC, sOU3zRv[g77TAi+h_Wr_yqq-1], txT77f, jaaWqY)

			pLoJx6e := iQKtjmO6[h_Wr_yqq-2]
			dKR0NUYWi, d9kfzR :=  /*line :1*/ mulWW(xPSdytG4S, pLoJx6e)
			rMBr6sXJZh2 := sOU3zRv[g77TAi+h_Wr_yqq-2]
			for  /*line :1*/ eNxLIIFYKXwf(dKR0NUYWi, d9kfzR, y5RYzdlV, rMBr6sXJZh2) {
				xPSdytG4S--
				jnVyfyoa := y5RYzdlV
				y5RYzdlV += txT77f

				if y5RYzdlV < jnVyfyoa {
					break
				}

				dKR0NUYWi, d9kfzR =  /*line :1*/ mulWW(xPSdytG4S, pLoJx6e)
			}
		}

		tu0jQu[h_Wr_yqq] =  /*line :1*/ pMUCN3(tu0jQu[0:h_Wr_yqq], iQKtjmO6, xPSdytG4S, 0)
		eiOa49PAlrv :=  /*line :1*/ len(tu0jQu)
		if g77TAi+eiOa49PAlrv >  /*line :1*/ len(sOU3zRv) && tu0jQu[h_Wr_yqq] == 0 {
			eiOa49PAlrv--
		}

		vKNCqP0 :=  /*line :1*/ omPlI3q_Zc(sOU3zRv[g77TAi:g77TAi+eiOa49PAlrv], sOU3zRv[g77TAi:], tu0jQu)
		if vKNCqP0 != 0 {
			vKNCqP0 :=  /*line :1*/ vmnD6P1U(sOU3zRv[g77TAi:g77TAi+h_Wr_yqq], sOU3zRv[g77TAi:], iQKtjmO6)

			if h_Wr_yqq < eiOa49PAlrv {
				sOU3zRv[g77TAi+h_Wr_yqq] += vKNCqP0
			}
			xPSdytG4S--
		}

		if g77TAi == gS2TXhHpYaX4 && gS2TXhHpYaX4 ==  /*line :1*/ len(yL0_p0wUnc) && xPSdytG4S == 0 {
			continue
		}
		yL0_p0wUnc[g77TAi] = xPSdytG4S
	}

	 /*line :1*/ zuavpO7v(o9KrWrdIsv)
}




func eNxLIIFYKXwf(dKR0NUYWi, d9kfzR, ihENsIZD, xBcmchOU VYe6D0) bool {
	return dKR0NUYWi > ihENsIZD || dKR0NUYWi == ihENsIZD && d9kfzR > xBcmchOU
}



const divRecursiveThreshold = 100






func (uW0eIVjy g3X9fa) gCGOsVPsJ3(sOU3zRv, iQKtjmO6 g3X9fa) {

	ai0CXVW := 2 *  /*line :1*/ bits.Len( /*line :1*/ uint( /*line :1*/ len(iQKtjmO6)))
	jRASocRnuh :=  /*line :1*/ aQEWrUmnl5I(3 *  /*line :1*/ len(iQKtjmO6))
	eo5JjPnaov4S :=  /*line :1*/ make([]*g3X9fa, ai0CXVW)

	 /*line :1*/ uW0eIVjy.uKEPQHWEut()
	 /*line :1*/ uW0eIVjy.gfulzKD3a8(sOU3zRv, iQKtjmO6, 0, jRASocRnuh, eo5JjPnaov4S)

	for _, h_Wr_yqq := range eo5JjPnaov4S {
		if h_Wr_yqq != nil {
			 /*line :1*/ zuavpO7v(h_Wr_yqq)
		}
	}
	 /*line :1*/ zuavpO7v(jRASocRnuh)
}






func (uW0eIVjy g3X9fa) gfulzKD3a8(sOU3zRv, iQKtjmO6 g3X9fa, c6AGhm int, jRASocRnuh *g3X9fa, eo5JjPnaov4S []*g3X9fa) {

	sOU3zRv =  /*line :1*/ sOU3zRv.hitXcy5()
	iQKtjmO6 =  /*line :1*/ iQKtjmO6.hitXcy5()
	if  /*line :1*/ len(sOU3zRv) == 0 {
		 /*line :1*/ uW0eIVjy.uKEPQHWEut()
		return
	}

	h_Wr_yqq :=  /*line :1*/ len(iQKtjmO6)
	if h_Wr_yqq < divRecursiveThreshold {
		 /*line :1*/ uW0eIVjy.dSb9JW6qvq(sOU3zRv, iQKtjmO6)
		return
	}

	gS2TXhHpYaX4 :=  /*line :1*/ len(sOU3zRv) - h_Wr_yqq
	if gS2TXhHpYaX4 < 0 {
		return
	}

	FyXHFes := h_Wr_yqq / 2

	if eo5JjPnaov4S[c6AGhm] == nil {
		eo5JjPnaov4S[c6AGhm] =  /*line :1*/ aQEWrUmnl5I(h_Wr_yqq)
	} else {
		*eo5JjPnaov4S[c6AGhm] =  /*line :1*/ eo5JjPnaov4S[c6AGhm].gS9bqHpKAJ0(FyXHFes + 1)
	}

	g77TAi := gS2TXhHpYaX4
	for g77TAi > FyXHFes {

		nR7KU4mGsdy := (FyXHFes - 1)

		dkEbuX55Af := sOU3zRv[g77TAi-FyXHFes:]

		xPSdytG4S := *eo5JjPnaov4S[c6AGhm]
		 /*line :1*/ xPSdytG4S.uKEPQHWEut()
		 /*line :1*/ xPSdytG4S.gfulzKD3a8(dkEbuX55Af[nR7KU4mGsdy:FyXHFes+h_Wr_yqq], iQKtjmO6[nR7KU4mGsdy:], c6AGhm+1, jRASocRnuh, eo5JjPnaov4S)
		xPSdytG4S =  /*line :1*/ xPSdytG4S.hitXcy5()

		tu0jQu :=  /*line :1*/ jRASocRnuh.gS9bqHpKAJ0(3 * h_Wr_yqq)
		 /*line :1*/ tu0jQu.uKEPQHWEut()
		tu0jQu =  /*line :1*/ tu0jQu.d22x6Ygoc80O(xPSdytG4S, iQKtjmO6[:nR7KU4mGsdy])
		for aepbxLiWOZ := 0; aepbxLiWOZ < 2; aepbxLiWOZ++ {
			vRGCE256ba4X :=  /*line :1*/ tu0jQu.beaurQHiT( /*line :1*/ dkEbuX55Af.hitXcy5())
			if vRGCE256ba4X <= 0 {
				break
			}
			 /*line :1*/ sQp2wtu(xPSdytG4S, xPSdytG4S, 1)
			vKNCqP0 :=  /*line :1*/ omPlI3q_Zc(tu0jQu[:nR7KU4mGsdy], tu0jQu[:nR7KU4mGsdy], iQKtjmO6[:nR7KU4mGsdy])
			if  /*line :1*/ len(tu0jQu) > nR7KU4mGsdy {
				 /*line :1*/ sQp2wtu(tu0jQu[nR7KU4mGsdy:], tu0jQu[nR7KU4mGsdy:], vKNCqP0)
			}
			 /*line :1*/ huhykmvX3UOx(dkEbuX55Af[nR7KU4mGsdy:], iQKtjmO6[nR7KU4mGsdy:], 0)
		}
		if  /*line :1*/ tu0jQu.beaurQHiT( /*line :1*/ dkEbuX55Af.hitXcy5()) > 0 {
			 /*line :1*/ panic("impossible")
		}
		vKNCqP0 :=  /*line :1*/ omPlI3q_Zc(dkEbuX55Af[: /*line :1*/ len(tu0jQu)], dkEbuX55Af[: /*line :1*/ len(tu0jQu)], tu0jQu)
		if vKNCqP0 > 0 {
			 /*line :1*/ sQp2wtu(dkEbuX55Af[ /*line :1*/ len(tu0jQu):], dkEbuX55Af[ /*line :1*/ len(tu0jQu):], vKNCqP0)
		}
		 /*line :1*/ huhykmvX3UOx(uW0eIVjy, xPSdytG4S, g77TAi-FyXHFes)
		g77TAi -= FyXHFes
	}

	nR7KU4mGsdy := FyXHFes - 1
	xPSdytG4S := *eo5JjPnaov4S[c6AGhm]
	 /*line :1*/ xPSdytG4S.uKEPQHWEut()
	 /*line :1*/ xPSdytG4S.gfulzKD3a8( /*line :1*/ sOU3zRv[nR7KU4mGsdy:].hitXcy5(), iQKtjmO6[nR7KU4mGsdy:], c6AGhm+1, jRASocRnuh, eo5JjPnaov4S)
	xPSdytG4S =  /*line :1*/ xPSdytG4S.hitXcy5()
	tu0jQu :=  /*line :1*/ jRASocRnuh.gS9bqHpKAJ0(3 * h_Wr_yqq)
	 /*line :1*/ tu0jQu.uKEPQHWEut()
	tu0jQu =  /*line :1*/ tu0jQu.d22x6Ygoc80O(xPSdytG4S, iQKtjmO6[:nR7KU4mGsdy])

	for aepbxLiWOZ := 0; aepbxLiWOZ < 2; aepbxLiWOZ++ {
		if vRGCE256ba4X :=  /*line :1*/ tu0jQu.beaurQHiT( /*line :1*/ sOU3zRv.hitXcy5()); vRGCE256ba4X > 0 {
			 /*line :1*/ sQp2wtu(xPSdytG4S, xPSdytG4S, 1)
			vKNCqP0 :=  /*line :1*/ omPlI3q_Zc(tu0jQu[:nR7KU4mGsdy], tu0jQu[:nR7KU4mGsdy], iQKtjmO6[:nR7KU4mGsdy])
			if  /*line :1*/ len(tu0jQu) > nR7KU4mGsdy {
				 /*line :1*/ sQp2wtu(tu0jQu[nR7KU4mGsdy:], tu0jQu[nR7KU4mGsdy:], vKNCqP0)
			}
			 /*line :1*/ huhykmvX3UOx(sOU3zRv[nR7KU4mGsdy:], iQKtjmO6[nR7KU4mGsdy:], 0)
		}
	}
	if  /*line :1*/ tu0jQu.beaurQHiT( /*line :1*/ sOU3zRv.hitXcy5()) > 0 {
		 /*line :1*/ panic("impossible")
	}
	vKNCqP0 :=  /*line :1*/ omPlI3q_Zc(sOU3zRv[0: /*line :1*/ len(tu0jQu)], sOU3zRv[0: /*line :1*/ len(tu0jQu)], tu0jQu)
	if vKNCqP0 > 0 {
		vKNCqP0 =  /*line :1*/ sQp2wtu(sOU3zRv[ /*line :1*/ len(tu0jQu):], sOU3zRv[ /*line :1*/ len(tu0jQu):], vKNCqP0)
	}
	if vKNCqP0 > 0 {
		 /*line :1*/ panic("impossible")
	}

	 /*line :1*/ huhykmvX3UOx(uW0eIVjy,  /*line :1*/ xPSdytG4S.hitXcy5(), 0)
}
