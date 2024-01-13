//line :1
package zBESu2SrRjP

import math "math"

var r6YcaCkRh = true	




func hgYVQXC6WLX(a0M8QSnCR, frkCDKOW1Ap string) int {
	wESnmZiAO :=  /*line :1*/ len(frkCDKOW1Ap)
	if wESnmZiAO >  /*line :1*/ len(a0M8QSnCR) {
		wESnmZiAO =  /*line :1*/ len(a0M8QSnCR)
	}
	for cDhV_2D := 0; cDhV_2D < wESnmZiAO; cDhV_2D++ {
		uTN3BXbMd := a0M8QSnCR[cDhV_2D]
		if 'A' <= uTN3BXbMd && uTN3BXbMd <= 'Z' {
			uTN3BXbMd += 'a' - 'A'
		}
		if uTN3BXbMd != frkCDKOW1Ap[cDhV_2D] {
			return cDhV_2D
		}
	}
	return wESnmZiAO
}






func bWm1aFAa3f5E(a0M8QSnCR string) (rhZE59OwqM5U float64, wESnmZiAO int, eva_XlRz4Iz bool) {
	if  /*line :1*/ len(a0M8QSnCR) == 0 {
		return 0, 0, false
	}
	koA2Mi := 1
	oHaAuX4 := 0
	switch a0M8QSnCR[0] {
	case '+', '-':
		if a0M8QSnCR[0] == '-' {
			koA2Mi = -1
		}
		oHaAuX4 = 1
		a0M8QSnCR = a0M8QSnCR[1:]
		fallthrough
	case 'i', 'I':
		wESnmZiAO :=  /*line :1*/ hgYVQXC6WLX(a0M8QSnCR, "infinity")

		if 3 < wESnmZiAO && wESnmZiAO < 8 {
			wESnmZiAO = 3
		}
		if wESnmZiAO == 3 || wESnmZiAO == 8 {
			return  /*line :1*/ math.FSug4WHZSBwz(koA2Mi), oHaAuX4 + wESnmZiAO, true
		}
	case 'n', 'N':
		if  /*line :1*/ hgYVQXC6WLX(a0M8QSnCR, "nan") == 3 {
			return  /*line :1*/ math.WG0hZIT4(), 3, true
		}
	}
	return 0, 0, false
}

func (o7OjhYc5wQns *fJphQdAo) jEayui(a0M8QSnCR string) (eva_XlRz4Iz bool) {
	cDhV_2D := 0
	o7OjhYc5wQns.hhg70L_qWPbD = false
	o7OjhYc5wQns.lpkYii9C = false

	if cDhV_2D >=  /*line :1*/ len(a0M8QSnCR) {
		return
	}
	switch {
	case a0M8QSnCR[cDhV_2D] == '+':
		cDhV_2D++
	case a0M8QSnCR[cDhV_2D] == '-':
		o7OjhYc5wQns.hhg70L_qWPbD = true
		cDhV_2D++
	}

	akmFAJ6b := false
	zBNC6q25gnGD := false
	for ; cDhV_2D <  /*line :1*/ len(a0M8QSnCR); cDhV_2D++ {
		switch {
		case a0M8QSnCR[cDhV_2D] == '_':

			continue
		case a0M8QSnCR[cDhV_2D] == '.':
			if akmFAJ6b {
				return
			}
			akmFAJ6b = true
			o7OjhYc5wQns.qvuKHF9jDC = o7OjhYc5wQns.ejR6KwzBSk
			continue

		case '0' <= a0M8QSnCR[cDhV_2D] && a0M8QSnCR[cDhV_2D] <= '9':
			zBNC6q25gnGD = true
			if a0M8QSnCR[cDhV_2D] == '0' && o7OjhYc5wQns.ejR6KwzBSk == 0 {
				o7OjhYc5wQns.qvuKHF9jDC--
				continue
			}
			if o7OjhYc5wQns.ejR6KwzBSk <  /*line :1*/ len(o7OjhYc5wQns.dXlUa8ko) {
				o7OjhYc5wQns.dXlUa8ko[o7OjhYc5wQns.ejR6KwzBSk] = a0M8QSnCR[cDhV_2D]
				o7OjhYc5wQns.ejR6KwzBSk++
			} else if a0M8QSnCR[cDhV_2D] != '0' {
				o7OjhYc5wQns.lpkYii9C = true
			}
			continue
		}
		break
	}
	if !zBNC6q25gnGD {
		return
	}
	if !akmFAJ6b {
		o7OjhYc5wQns.qvuKHF9jDC = o7OjhYc5wQns.ejR6KwzBSk
	}

	if cDhV_2D <  /*line :1*/ len(a0M8QSnCR) &&  /*line :1*/ chHT2YtqEP(a0M8QSnCR[cDhV_2D]) == 'e' {
		cDhV_2D++
		if cDhV_2D >=  /*line :1*/ len(a0M8QSnCR) {
			return
		}
		kzbr5fD4 := 1
		if a0M8QSnCR[cDhV_2D] == '+' {
			cDhV_2D++
		} else if a0M8QSnCR[cDhV_2D] == '-' {
			cDhV_2D++
			kzbr5fD4 = -1
		}
		if cDhV_2D >=  /*line :1*/ len(a0M8QSnCR) || a0M8QSnCR[cDhV_2D] < '0' || a0M8QSnCR[cDhV_2D] > '9' {
			return
		}
		afGPrH := 0
		for ; cDhV_2D <  /*line :1*/ len(a0M8QSnCR) && ('0' <= a0M8QSnCR[cDhV_2D] && a0M8QSnCR[cDhV_2D] <= '9' || a0M8QSnCR[cDhV_2D] == '_'); cDhV_2D++ {
			if a0M8QSnCR[cDhV_2D] == '_' {

				continue
			}
			if afGPrH < 10000 {
				afGPrH = afGPrH*10 +  /*line :1*/ int(a0M8QSnCR[cDhV_2D]) - '0'
			}
		}
		o7OjhYc5wQns.qvuKHF9jDC += afGPrH * kzbr5fD4
	}

	if cDhV_2D !=  /*line :1*/ len(a0M8QSnCR) {
		return
	}

	eva_XlRz4Iz = true
	return
}





func wHkFOK4J5e(a0M8QSnCR string) (kfBeJy9lqcOd uint64, uZdS176 int, an5n8FniWt67, oersOWF, r6avN8QxQ8D bool, cDhV_2D int, eva_XlRz4Iz bool) {
	jKtJBS := false

	if cDhV_2D >=  /*line :1*/ len(a0M8QSnCR) {
		return
	}
	switch {
	case a0M8QSnCR[cDhV_2D] == '+':
		cDhV_2D++
	case a0M8QSnCR[cDhV_2D] == '-':
		an5n8FniWt67 = true
		cDhV_2D++
	}

	xsj80lMaPD :=  /*line :1*/ uint64(10)
	s1_hnw := 19
	m8ubCa :=  /*line :1*/ byte('e')
	if cDhV_2D+2 <  /*line :1*/ len(a0M8QSnCR) && a0M8QSnCR[cDhV_2D] == '0' &&  /*line :1*/ chHT2YtqEP(a0M8QSnCR[cDhV_2D+1]) == 'x' {
		xsj80lMaPD = 16
		s1_hnw = 16
		cDhV_2D += 2
		m8ubCa = 'p'
		r6avN8QxQ8D = true
	}
	akmFAJ6b := false
	zBNC6q25gnGD := false
	b9NNFc_Hmo_7 := 0
	m0wNmk9 := 0
	fB_k418 := 0
loop:
	for ; cDhV_2D <  /*line :1*/ len(a0M8QSnCR); cDhV_2D++ {
		switch uTN3BXbMd := a0M8QSnCR[cDhV_2D]; true {
		case uTN3BXbMd == '_':
			jKtJBS = true
			continue

		case uTN3BXbMd == '.':
			if akmFAJ6b {
				break loop
			}
			akmFAJ6b = true
			fB_k418 = b9NNFc_Hmo_7
			continue

		case '0' <= uTN3BXbMd && uTN3BXbMd <= '9':
			zBNC6q25gnGD = true
			if uTN3BXbMd == '0' && b9NNFc_Hmo_7 == 0 {
				fB_k418--
				continue
			}
			b9NNFc_Hmo_7++
			if m0wNmk9 < s1_hnw {
				kfBeJy9lqcOd *= xsj80lMaPD
				kfBeJy9lqcOd +=  /*line :1*/ uint64(uTN3BXbMd - '0')
				m0wNmk9++
			} else if uTN3BXbMd != '0' {
				oersOWF = true
			}
			continue

		case xsj80lMaPD == 16 && 'a' <=  /*line :1*/ chHT2YtqEP(uTN3BXbMd) &&  /*line :1*/ chHT2YtqEP(uTN3BXbMd) <= 'f':
			zBNC6q25gnGD = true
			b9NNFc_Hmo_7++
			if m0wNmk9 < s1_hnw {
				kfBeJy9lqcOd *= 16
				kfBeJy9lqcOd +=  /*line :1*/ uint64( /*line :1*/ chHT2YtqEP(uTN3BXbMd) - 'a' + 10)
				m0wNmk9++
			} else {
				oersOWF = true
			}
			continue
		}
		break
	}
	if !zBNC6q25gnGD {
		return
	}
	if !akmFAJ6b {
		fB_k418 = b9NNFc_Hmo_7
	}

	if xsj80lMaPD == 16 {
		fB_k418 *= 4
		m0wNmk9 *= 4
	}

	if cDhV_2D <  /*line :1*/ len(a0M8QSnCR) &&  /*line :1*/ chHT2YtqEP(a0M8QSnCR[cDhV_2D]) == m8ubCa {
		cDhV_2D++
		if cDhV_2D >=  /*line :1*/ len(a0M8QSnCR) {
			return
		}
		kzbr5fD4 := 1
		if a0M8QSnCR[cDhV_2D] == '+' {
			cDhV_2D++
		} else if a0M8QSnCR[cDhV_2D] == '-' {
			cDhV_2D++
			kzbr5fD4 = -1
		}
		if cDhV_2D >=  /*line :1*/ len(a0M8QSnCR) || a0M8QSnCR[cDhV_2D] < '0' || a0M8QSnCR[cDhV_2D] > '9' {
			return
		}
		afGPrH := 0
		for ; cDhV_2D <  /*line :1*/ len(a0M8QSnCR) && ('0' <= a0M8QSnCR[cDhV_2D] && a0M8QSnCR[cDhV_2D] <= '9' || a0M8QSnCR[cDhV_2D] == '_'); cDhV_2D++ {
			if a0M8QSnCR[cDhV_2D] == '_' {
				jKtJBS = true
				continue
			}
			if afGPrH < 10000 {
				afGPrH = afGPrH*10 +  /*line :1*/ int(a0M8QSnCR[cDhV_2D]) - '0'
			}
		}
		fB_k418 += afGPrH * kzbr5fD4
	} else if xsj80lMaPD == 16 {

		return
	}

	if kfBeJy9lqcOd != 0 {
		uZdS176 = fB_k418 - m0wNmk9
	}

	if jKtJBS && ! /*line :1*/ kTU2kqFg(a0M8QSnCR[:cDhV_2D]) {
		return
	}

	eva_XlRz4Iz = true
	return
}


var rtz4iLh = []int{1, 3, 6, 9, 13, 16, 19, 23, 26}

func (bkcwawBry *fJphQdAo) azaJV_Lh6(sRaJkQTYfNyc *aL4q0U) (o7OjhYc5wQns uint64, lFZaf00ou bool) {
	var uZdS176 int
	var fdyaiZ uint64

	if bkcwawBry.ejR6KwzBSk == 0 {
		fdyaiZ = 0
		uZdS176 = sRaJkQTYfNyc.oXPRffP4
		goto out
	}

	if bkcwawBry.qvuKHF9jDC > 310 {
		goto overflow
	}
	if bkcwawBry.qvuKHF9jDC < -330 {

		fdyaiZ = 0
		uZdS176 = sRaJkQTYfNyc.oXPRffP4
		goto out
	}

	uZdS176 = 0
	for bkcwawBry.qvuKHF9jDC > 0 {
		var wESnmZiAO int
		if bkcwawBry.qvuKHF9jDC >=  /*line :1*/ len(rtz4iLh) {
			wESnmZiAO = 27
		} else {
			wESnmZiAO = rtz4iLh[bkcwawBry.qvuKHF9jDC]
		}
		 /*line :1*/ bkcwawBry.Shift(-wESnmZiAO)
		uZdS176 += wESnmZiAO
	}
	for bkcwawBry.qvuKHF9jDC < 0 || bkcwawBry.qvuKHF9jDC == 0 && bkcwawBry.dXlUa8ko[0] < '5' {
		var wESnmZiAO int
		if -bkcwawBry.qvuKHF9jDC >=  /*line :1*/ len(rtz4iLh) {
			wESnmZiAO = 27
		} else {
			wESnmZiAO = rtz4iLh[-bkcwawBry.qvuKHF9jDC]
		}
		 /*line :1*/ bkcwawBry.Shift(wESnmZiAO)
		uZdS176 -= wESnmZiAO
	}

	uZdS176--

	if uZdS176 < sRaJkQTYfNyc.oXPRffP4+1 {
		wESnmZiAO := sRaJkQTYfNyc.oXPRffP4 + 1 - uZdS176
		 /*line :1*/ bkcwawBry.Shift(-wESnmZiAO)
		uZdS176 += wESnmZiAO
	}

	if uZdS176-sRaJkQTYfNyc.oXPRffP4 >= 1<<sRaJkQTYfNyc.iPKjS5sj-1 {
		goto overflow
	}

	 /*line :1*/ bkcwawBry.Shift( /*line :1*/ int(1 + sRaJkQTYfNyc.hCC_Vh))
	fdyaiZ =  /*line :1*/ bkcwawBry.RoundedInteger()

	if fdyaiZ == 2<<sRaJkQTYfNyc.hCC_Vh {
		fdyaiZ >>= 1
		uZdS176++
		if uZdS176-sRaJkQTYfNyc.oXPRffP4 >= 1<<sRaJkQTYfNyc.iPKjS5sj-1 {
			goto overflow
		}
	}

	if fdyaiZ&(1<<sRaJkQTYfNyc.hCC_Vh) == 0 {
		uZdS176 = sRaJkQTYfNyc.oXPRffP4
	}
	goto out

overflow:

	fdyaiZ = 0
	uZdS176 = 1<<sRaJkQTYfNyc.iPKjS5sj - 1 + sRaJkQTYfNyc.oXPRffP4
	lFZaf00ou = true

out:

	gYquMvlAg := fdyaiZ & ( /*line :1*/ uint64(1)<<sRaJkQTYfNyc.hCC_Vh - 1)
	gYquMvlAg |=  /*line :1*/ uint64((uZdS176-sRaJkQTYfNyc.oXPRffP4)&(1<<sRaJkQTYfNyc.iPKjS5sj-1)) << sRaJkQTYfNyc.hCC_Vh
	if bkcwawBry.hhg70L_qWPbD {
		gYquMvlAg |= 1 << sRaJkQTYfNyc.hCC_Vh << sRaJkQTYfNyc.iPKjS5sj
	}
	return gYquMvlAg, lFZaf00ou
}


var i0IAcs = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
	1e20, 1e21, 1e22,
}
var hd23I5LlLWtG = []float32{1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10}










func b59RVM5vs8w(kfBeJy9lqcOd uint64, uZdS176 int, an5n8FniWt67 bool) (rhZE59OwqM5U float64, eva_XlRz4Iz bool) {
	if kfBeJy9lqcOd>>pn7mkGeP.hCC_Vh != 0 {
		return
	}
	rhZE59OwqM5U =  /*line :1*/ float64(kfBeJy9lqcOd)
	if an5n8FniWt67 {
		rhZE59OwqM5U = -rhZE59OwqM5U
	}
	switch {
	case uZdS176 == 0:

		return rhZE59OwqM5U, true

	case uZdS176 > 0 && uZdS176 <= 15+22:

		if uZdS176 > 22 {
			rhZE59OwqM5U *= i0IAcs[uZdS176-22]
			uZdS176 = 22
		}
		if rhZE59OwqM5U > 1e15 || rhZE59OwqM5U < -1e15 {

			return
		}
		return rhZE59OwqM5U * i0IAcs[uZdS176], true
	case uZdS176 < 0 && uZdS176 >= -22:
		return rhZE59OwqM5U / i0IAcs[-uZdS176], true
	}
	return
}



func w37CzVTyb(kfBeJy9lqcOd uint64, uZdS176 int, an5n8FniWt67 bool) (rhZE59OwqM5U float32, eva_XlRz4Iz bool) {
	if kfBeJy9lqcOd>>cuN5E5Tcg_0.hCC_Vh != 0 {
		return
	}
	rhZE59OwqM5U =  /*line :1*/ float32(kfBeJy9lqcOd)
	if an5n8FniWt67 {
		rhZE59OwqM5U = -rhZE59OwqM5U
	}
	switch {
	case uZdS176 == 0:
		return rhZE59OwqM5U, true

	case uZdS176 > 0 && uZdS176 <= 7+10:

		if uZdS176 > 10 {
			rhZE59OwqM5U *= hd23I5LlLWtG[uZdS176-10]
			uZdS176 = 10
		}
		if rhZE59OwqM5U > 1e7 || rhZE59OwqM5U < -1e7 {

			return
		}
		return rhZE59OwqM5U * hd23I5LlLWtG[uZdS176], true
	case uZdS176 < 0 && uZdS176 >= -10:
		return rhZE59OwqM5U / hd23I5LlLWtG[-uZdS176], true
	}
	return
}






func bkR_S9ftH5(a0M8QSnCR string, sRaJkQTYfNyc *aL4q0U, kfBeJy9lqcOd uint64, uZdS176 int, an5n8FniWt67, oersOWF bool) (float64, error) {
	urdTZptdwxqe := 1<<sRaJkQTYfNyc.iPKjS5sj + sRaJkQTYfNyc.oXPRffP4 - 2
	gQrC6QEiau := sRaJkQTYfNyc.oXPRffP4 + 1
	uZdS176 +=  /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh)

	for kfBeJy9lqcOd != 0 && kfBeJy9lqcOd>>(sRaJkQTYfNyc.hCC_Vh+2) == 0 {
		kfBeJy9lqcOd <<= 1
		uZdS176--
	}
	if oersOWF {
		kfBeJy9lqcOd |= 1
	}
	for kfBeJy9lqcOd>>(1+sRaJkQTYfNyc.hCC_Vh+2) != 0 {
		kfBeJy9lqcOd = kfBeJy9lqcOd>>1 | kfBeJy9lqcOd&1
		uZdS176++
	}

	for kfBeJy9lqcOd > 1 && uZdS176 < gQrC6QEiau-2 {
		kfBeJy9lqcOd = kfBeJy9lqcOd>>1 | kfBeJy9lqcOd&1
		uZdS176++
	}

	rg5fg1Xg := kfBeJy9lqcOd & 3
	kfBeJy9lqcOd >>= 2
	rg5fg1Xg |= kfBeJy9lqcOd & 1
	uZdS176 += 2
	if rg5fg1Xg == 3 {
		kfBeJy9lqcOd++
		if kfBeJy9lqcOd == 1<<(1+sRaJkQTYfNyc.hCC_Vh) {
			kfBeJy9lqcOd >>= 1
			uZdS176++
		}
	}

	if kfBeJy9lqcOd>>sRaJkQTYfNyc.hCC_Vh == 0 {
		uZdS176 = sRaJkQTYfNyc.oXPRffP4
	}
	var kWNmhkfeXX error
	if uZdS176 > urdTZptdwxqe {
		kfBeJy9lqcOd = 1 << sRaJkQTYfNyc.hCC_Vh
		uZdS176 = urdTZptdwxqe + 1
		kWNmhkfeXX =  /*line :1*/ b_MqXW4(fnParseFloat, a0M8QSnCR)
	}

	gYquMvlAg := kfBeJy9lqcOd & (1<<sRaJkQTYfNyc.hCC_Vh - 1)
	gYquMvlAg |=  /*line :1*/ uint64((uZdS176-sRaJkQTYfNyc.oXPRffP4)&(1<<sRaJkQTYfNyc.iPKjS5sj-1)) << sRaJkQTYfNyc.hCC_Vh
	if an5n8FniWt67 {
		gYquMvlAg |= 1 << sRaJkQTYfNyc.hCC_Vh << sRaJkQTYfNyc.iPKjS5sj
	}
	if sRaJkQTYfNyc == &cuN5E5Tcg_0 {
		return  /*line :1*/ float64( /*line :1*/ math.AviWp5b0( /*line :1*/ uint32(gYquMvlAg))), kWNmhkfeXX
	}
	return  /*line :1*/ math.QUB2Kf63p(gYquMvlAg), kWNmhkfeXX
}

const fnParseFloat = "ParseFloat"

func gUMSQN0(a0M8QSnCR string) (rhZE59OwqM5U float32, wESnmZiAO int, kWNmhkfeXX error) {
	if ic2uVKEK, wESnmZiAO, eva_XlRz4Iz :=  /*line :1*/ bWm1aFAa3f5E(a0M8QSnCR); eva_XlRz4Iz {
		return  /*line :1*/ float32(ic2uVKEK), wESnmZiAO, nil
	}

	kfBeJy9lqcOd, uZdS176, an5n8FniWt67, oersOWF, r6avN8QxQ8D, wESnmZiAO, eva_XlRz4Iz :=  /*line :1*/ wHkFOK4J5e(a0M8QSnCR)
	if !eva_XlRz4Iz {
		return 0, wESnmZiAO,  /*line :1*/ nF25LgNzv(fnParseFloat, a0M8QSnCR)
	}

	if r6avN8QxQ8D {
		rhZE59OwqM5U, kWNmhkfeXX :=  /*line :1*/ bkR_S9ftH5(a0M8QSnCR[:wESnmZiAO], &cuN5E5Tcg_0, kfBeJy9lqcOd, uZdS176, an5n8FniWt67, oersOWF)
		return  /*line :1*/ float32(rhZE59OwqM5U), wESnmZiAO, kWNmhkfeXX
	}

	if r6YcaCkRh {

		if !oersOWF {
			if rhZE59OwqM5U, eva_XlRz4Iz :=  /*line :1*/ w37CzVTyb(kfBeJy9lqcOd, uZdS176, an5n8FniWt67); eva_XlRz4Iz {
				return rhZE59OwqM5U, wESnmZiAO, nil
			}
		}
		rhZE59OwqM5U, eva_XlRz4Iz :=  /*line :1*/ letapKfAwewW(kfBeJy9lqcOd, uZdS176, an5n8FniWt67)
		if eva_XlRz4Iz {
			if !oersOWF {
				return rhZE59OwqM5U, wESnmZiAO, nil
			}

			kRS35eW_a, eva_XlRz4Iz :=  /*line :1*/ letapKfAwewW(kfBeJy9lqcOd+1, uZdS176, an5n8FniWt67)
			if eva_XlRz4Iz && rhZE59OwqM5U == kRS35eW_a {
				return rhZE59OwqM5U, wESnmZiAO, nil
			}
		}
	}

	
	var bkcwawBry fJphQdAo
	if ! /*line :1*/ bkcwawBry.jEayui(a0M8QSnCR[:wESnmZiAO]) {
		return 0, wESnmZiAO,  /*line :1*/ nF25LgNzv(fnParseFloat, a0M8QSnCR)
	}
	o7OjhYc5wQns, xePYwh0fBD :=  /*line :1*/ bkcwawBry.azaJV_Lh6(&cuN5E5Tcg_0)
	rhZE59OwqM5U =  /*line :1*/ math.AviWp5b0( /*line :1*/ uint32(o7OjhYc5wQns))
	if xePYwh0fBD {
		kWNmhkfeXX =  /*line :1*/ b_MqXW4(fnParseFloat, a0M8QSnCR)
	}
	return rhZE59OwqM5U, wESnmZiAO, kWNmhkfeXX
}

func qx1kWm(a0M8QSnCR string) (rhZE59OwqM5U float64, wESnmZiAO int, kWNmhkfeXX error) {
	if ic2uVKEK, wESnmZiAO, eva_XlRz4Iz :=  /*line :1*/ bWm1aFAa3f5E(a0M8QSnCR); eva_XlRz4Iz {
		return ic2uVKEK, wESnmZiAO, nil
	}

	kfBeJy9lqcOd, uZdS176, an5n8FniWt67, oersOWF, r6avN8QxQ8D, wESnmZiAO, eva_XlRz4Iz :=  /*line :1*/ wHkFOK4J5e(a0M8QSnCR)
	if !eva_XlRz4Iz {
		return 0, wESnmZiAO,  /*line :1*/ nF25LgNzv(fnParseFloat, a0M8QSnCR)
	}

	if r6avN8QxQ8D {
		rhZE59OwqM5U, kWNmhkfeXX :=  /*line :1*/ bkR_S9ftH5(a0M8QSnCR[:wESnmZiAO], &pn7mkGeP, kfBeJy9lqcOd, uZdS176, an5n8FniWt67, oersOWF)
		return rhZE59OwqM5U, wESnmZiAO, kWNmhkfeXX
	}

	if r6YcaCkRh {

		if !oersOWF {
			if rhZE59OwqM5U, eva_XlRz4Iz :=  /*line :1*/ b59RVM5vs8w(kfBeJy9lqcOd, uZdS176, an5n8FniWt67); eva_XlRz4Iz {
				return rhZE59OwqM5U, wESnmZiAO, nil
			}
		}
		rhZE59OwqM5U, eva_XlRz4Iz :=  /*line :1*/ hJPSKYGb(kfBeJy9lqcOd, uZdS176, an5n8FniWt67)
		if eva_XlRz4Iz {
			if !oersOWF {
				return rhZE59OwqM5U, wESnmZiAO, nil
			}

			kRS35eW_a, eva_XlRz4Iz :=  /*line :1*/ hJPSKYGb(kfBeJy9lqcOd+1, uZdS176, an5n8FniWt67)
			if eva_XlRz4Iz && rhZE59OwqM5U == kRS35eW_a {
				return rhZE59OwqM5U, wESnmZiAO, nil
			}
		}
	}

	
	var bkcwawBry fJphQdAo
	if ! /*line :1*/ bkcwawBry.jEayui(a0M8QSnCR[:wESnmZiAO]) {
		return 0, wESnmZiAO,  /*line :1*/ nF25LgNzv(fnParseFloat, a0M8QSnCR)
	}
	o7OjhYc5wQns, xePYwh0fBD :=  /*line :1*/ bkcwawBry.azaJV_Lh6(&pn7mkGeP)
	rhZE59OwqM5U =  /*line :1*/ math.QUB2Kf63p(o7OjhYc5wQns)
	if xePYwh0fBD {
		kWNmhkfeXX =  /*line :1*/ b_MqXW4(fnParseFloat, a0M8QSnCR)
	}
	return rhZE59OwqM5U, wESnmZiAO, kWNmhkfeXX
}




























func YciY2Zn(a0M8QSnCR string, r2moFi1 int) (float64, error) {
	rhZE59OwqM5U, wESnmZiAO, kWNmhkfeXX :=  /*line :1*/ xBeCkrUFdCv(a0M8QSnCR, r2moFi1)
	if wESnmZiAO !=  /*line :1*/ len(a0M8QSnCR) && (kWNmhkfeXX == nil || kWNmhkfeXX.(*RJxImsn8bz).S08K00Kl9csG != ZTbjMv) {
		return 0,  /*line :1*/ nF25LgNzv(fnParseFloat, a0M8QSnCR)
	}
	return rhZE59OwqM5U, kWNmhkfeXX
}

func xBeCkrUFdCv(a0M8QSnCR string, r2moFi1 int) (float64, int, error) {
	if r2moFi1 == 32 {
		rhZE59OwqM5U, wESnmZiAO, kWNmhkfeXX :=  /*line :1*/ gUMSQN0(a0M8QSnCR)
		return  /*line :1*/ float64(rhZE59OwqM5U), wESnmZiAO, kWNmhkfeXX
	}
	return  /*line :1*/ qx1kWm(a0M8QSnCR)
}
