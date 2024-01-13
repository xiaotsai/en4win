//line :1
package zBESu2SrRjP

import (
	bits "math/bits"
)


func g3Ug4aphDkYU(bkcwawBry *cN7vNMaJUh, fdyaiZ uint32, uZdS176 int, gsf1Ql int) {
	if gsf1Ql < 0 {
		 /*line :1*/ panic("ryuFtoaFixed32 called with negative prec")
	}
	if gsf1Ql > 9 {
		 /*line :1*/ panic("ryuFtoaFixed32 called with prec > 9")
	}

	if fdyaiZ == 0 {
		bkcwawBry.hc_0uW85, bkcwawBry.czr7NK2VrVdf = 0, 0
		return
	}

	e4FXxUi := uZdS176
	if o7OjhYc5wQns :=  /*line :1*/ bits.Len32(fdyaiZ); o7OjhYc5wQns < 25 {
		fdyaiZ <<=  /*line :1*/ uint(25 - o7OjhYc5wQns)
		e4FXxUi += o7OjhYc5wQns - 25
	}

	oUyevz5E2Vjr := - /*line :1*/ ouR7FiDq(e4FXxUi+24) + gsf1Ql - 1

	iAXjT3 := oUyevz5E2Vjr <= 27 && oUyevz5E2Vjr >= 0

	fstifV2GX, i93aVl_7Pq7e, sfzCfaP :=  /*line :1*/ xXJNqhjxML(fdyaiZ, e4FXxUi, oUyevz5E2Vjr)
	if i93aVl_7Pq7e >= 0 {
		 /*line :1*/ panic("not enough significant bits after mult64bitPow10")
	}

	if oUyevz5E2Vjr < 0 && oUyevz5E2Vjr >= -10 &&  /*line :1*/ rtIFITrI( /*line :1*/ uint64(fdyaiZ), -oUyevz5E2Vjr) {
		iAXjT3 = true
		sfzCfaP = true
	}

	v6ctH0 :=  /*line :1*/ uint(-i93aVl_7Pq7e)
	nEbahhuccs :=  /*line :1*/ uint32(1<<v6ctH0 - 1)

	fstifV2GX, n0FsyNBnET := fstifV2GX>>v6ctH0, fstifV2GX&nEbahhuccs
	jAxSc9 := false
	if iAXjT3 {

		jAxSc9 = n0FsyNBnET > 1<<(v6ctH0-1) ||
			(n0FsyNBnET == 1<<(v6ctH0-1) && !sfzCfaP) ||
			(n0FsyNBnET == 1<<(v6ctH0-1) && sfzCfaP && fstifV2GX&1 == 1)
	} else {

		jAxSc9 = n0FsyNBnET>>(v6ctH0-1) == 1
	}
	if n0FsyNBnET != 0 {
		sfzCfaP = false
	}

	 /*line :1*/ oHoaWxbajvXb(bkcwawBry,  /*line :1*/ uint64(fstifV2GX), !sfzCfaP, jAxSc9, gsf1Ql)

	bkcwawBry.czr7NK2VrVdf -= oUyevz5E2Vjr
}


func asxn8U6lpW(bkcwawBry *cN7vNMaJUh, fdyaiZ uint64, uZdS176 int, gsf1Ql int) {
	if gsf1Ql > 18 {
		 /*line :1*/ panic("ryuFtoaFixed64 called with prec > 18")
	}

	if fdyaiZ == 0 {
		bkcwawBry.hc_0uW85, bkcwawBry.czr7NK2VrVdf = 0, 0
		return
	}

	e4FXxUi := uZdS176
	if o7OjhYc5wQns :=  /*line :1*/ bits.Len64(fdyaiZ); o7OjhYc5wQns < 55 {
		fdyaiZ = fdyaiZ <<  /*line :1*/ uint(55-o7OjhYc5wQns)
		e4FXxUi += o7OjhYc5wQns - 55
	}

	oUyevz5E2Vjr := - /*line :1*/ ouR7FiDq(e4FXxUi+54) + gsf1Ql - 1

	iAXjT3 := oUyevz5E2Vjr <= 55 && oUyevz5E2Vjr >= 0

	fstifV2GX, i93aVl_7Pq7e, sfzCfaP :=  /*line :1*/ gM37tkftIW(fdyaiZ, e4FXxUi, oUyevz5E2Vjr)
	if i93aVl_7Pq7e >= 0 {
		 /*line :1*/ panic("not enough significant bits after mult128bitPow10")
	}

	if oUyevz5E2Vjr < 0 && oUyevz5E2Vjr >= -22 &&  /*line :1*/ rtIFITrI(fdyaiZ, -oUyevz5E2Vjr) {
		iAXjT3 = true
		sfzCfaP = true
	}

	v6ctH0 :=  /*line :1*/ uint(-i93aVl_7Pq7e)
	nEbahhuccs :=  /*line :1*/ uint64(1<<v6ctH0 - 1)

	fstifV2GX, n0FsyNBnET := fstifV2GX>>v6ctH0, fstifV2GX&nEbahhuccs
	jAxSc9 := false
	if iAXjT3 {

		jAxSc9 = n0FsyNBnET > 1<<(v6ctH0-1) ||
			(n0FsyNBnET == 1<<(v6ctH0-1) && !sfzCfaP) ||
			(n0FsyNBnET == 1<<(v6ctH0-1) && sfzCfaP && fstifV2GX&1 == 1)
	} else {

		jAxSc9 = n0FsyNBnET>>(v6ctH0-1) == 1
	}
	if n0FsyNBnET != 0 {
		sfzCfaP = false
	}

	 /*line :1*/ oHoaWxbajvXb(bkcwawBry, fstifV2GX, !sfzCfaP, jAxSc9, gsf1Ql)

	bkcwawBry.czr7NK2VrVdf -= oUyevz5E2Vjr
}

var m8P53iCGlbNy = [...]uint64{
	1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}




func oHoaWxbajvXb(bkcwawBry *cN7vNMaJUh, csRwMg uint64, oersOWF bool, jAxSc9 bool, gsf1Ql int) {
	f2STSXp := m8P53iCGlbNy[gsf1Ql]
	fkm7C_rY3E := 0
	for csRwMg >= f2STSXp {
		zxZt5L, o7OjhYc5wQns := csRwMg/10, csRwMg%10
		csRwMg = zxZt5L
		fkm7C_rY3E++
		if o7OjhYc5wQns > 5 {
			jAxSc9 = true
		} else if o7OjhYc5wQns < 5 {
			jAxSc9 = false
		} else {

			jAxSc9 = oersOWF || csRwMg&1 == 1
		}
		if o7OjhYc5wQns != 0 {
			oersOWF = true
		}
	}
	if jAxSc9 {
		csRwMg++
	}
	if csRwMg >= f2STSXp {

		csRwMg /= 10
		fkm7C_rY3E++
	}

	wESnmZiAO :=  /*line :1*/ uint(gsf1Ql)
	bkcwawBry.hc_0uW85 = gsf1Ql
	ggqcIE := csRwMg
	for ggqcIE >= 100 {
		var zCsD112MLqZf, mOjrrJ87 uint64
		if ggqcIE>>32 == 0 {
			zCsD112MLqZf, mOjrrJ87 =  /*line :1*/ uint64( /*line :1*/ uint32(ggqcIE)/100),  /*line :1*/ uint64( /*line :1*/ uint32(ggqcIE)%100)
		} else {
			zCsD112MLqZf, mOjrrJ87 = ggqcIE/100, ggqcIE%100
		}
		wESnmZiAO -= 2
		bkcwawBry.dvBTkobK6ZN[wESnmZiAO+1] = smallsString[2*mOjrrJ87+1]
		bkcwawBry.dvBTkobK6ZN[wESnmZiAO+0] = smallsString[2*mOjrrJ87+0]
		ggqcIE = zCsD112MLqZf
	}
	if ggqcIE > 0 {
		wESnmZiAO--
		bkcwawBry.dvBTkobK6ZN[wESnmZiAO] = smallsString[2*ggqcIE+1]
	}
	if ggqcIE >= 10 {
		wESnmZiAO--
		bkcwawBry.dvBTkobK6ZN[wESnmZiAO] = smallsString[2*ggqcIE]
	}
	for bkcwawBry.dvBTkobK6ZN[bkcwawBry.hc_0uW85-1] == '0' {
		bkcwawBry.hc_0uW85--
		fkm7C_rY3E++
	}
	bkcwawBry.czr7NK2VrVdf = bkcwawBry.hc_0uW85 + fkm7C_rY3E
}


func dY6yYkbo(bkcwawBry *cN7vNMaJUh, fdyaiZ uint64, uZdS176 int, sRaJkQTYfNyc *aL4q0U) {
	if fdyaiZ == 0 {
		bkcwawBry.hc_0uW85, bkcwawBry.czr7NK2VrVdf = 0, 0
		return
	}

	if uZdS176 <= 0 &&  /*line :1*/ bits.TrailingZeros64(fdyaiZ) >= -uZdS176 {
		fdyaiZ >>=  /*line :1*/ uint(-uZdS176)
		 /*line :1*/ bG4mTMNp1NT(bkcwawBry, fdyaiZ, fdyaiZ, fdyaiZ, true, false)
		return
	}
	jeaP5BYBjL, c3Uc2_6AsI, wHUK53OghI, e4FXxUi :=  /*line :1*/ s9hB2_0A(fdyaiZ, uZdS176, sRaJkQTYfNyc)
	if e4FXxUi == 0 {
		 /*line :1*/ bG4mTMNp1NT(bkcwawBry, jeaP5BYBjL, c3Uc2_6AsI, wHUK53OghI, true, false)
		return
	}

	oUyevz5E2Vjr :=  /*line :1*/ ouR7FiDq(-e4FXxUi) + 1

	
	
	var ljUgJLytUS, acgiVA06Q2, rAgI2dKQj0 uint64
	var tbdsmf9aTny, nhLf1v, rOTeKd3Q bool
	if sRaJkQTYfNyc == &cuN5E5Tcg_0 {
		var owJDoTg9Te7p, aV3X38DNcguy, aAOSwzgXw2 uint32
		owJDoTg9Te7p, _, tbdsmf9aTny =  /*line :1*/ xXJNqhjxML( /*line :1*/ uint32(jeaP5BYBjL), e4FXxUi, oUyevz5E2Vjr)
		aV3X38DNcguy, _, nhLf1v =  /*line :1*/ xXJNqhjxML( /*line :1*/ uint32(c3Uc2_6AsI), e4FXxUi, oUyevz5E2Vjr)
		aAOSwzgXw2, e4FXxUi, rOTeKd3Q =  /*line :1*/ xXJNqhjxML( /*line :1*/ uint32(wHUK53OghI), e4FXxUi, oUyevz5E2Vjr)
		ljUgJLytUS, acgiVA06Q2, rAgI2dKQj0 =  /*line :1*/ uint64(owJDoTg9Te7p),  /*line :1*/ uint64(aV3X38DNcguy),  /*line :1*/ uint64(aAOSwzgXw2)
	} else {
		ljUgJLytUS, _, tbdsmf9aTny =  /*line :1*/ gM37tkftIW(jeaP5BYBjL, e4FXxUi, oUyevz5E2Vjr)
		acgiVA06Q2, _, nhLf1v =  /*line :1*/ gM37tkftIW(c3Uc2_6AsI, e4FXxUi, oUyevz5E2Vjr)
		rAgI2dKQj0, e4FXxUi, rOTeKd3Q =  /*line :1*/ gM37tkftIW(wHUK53OghI, e4FXxUi, oUyevz5E2Vjr)
	}
	if e4FXxUi >= 0 {
		 /*line :1*/ panic("not enough significant bits after mult128bitPow10")
	}

	if oUyevz5E2Vjr > 55 {

		tbdsmf9aTny, nhLf1v, rOTeKd3Q = false, false, false
	}
	if oUyevz5E2Vjr < 0 && oUyevz5E2Vjr >= -24 {

		if  /*line :1*/ rtIFITrI(jeaP5BYBjL, -oUyevz5E2Vjr) {
			tbdsmf9aTny = true
		}
		if  /*line :1*/ rtIFITrI(c3Uc2_6AsI, -oUyevz5E2Vjr) {
			nhLf1v = true
		}
		if  /*line :1*/ rtIFITrI(wHUK53OghI, -oUyevz5E2Vjr) {
			rOTeKd3Q = true
		}
	}

	v6ctH0 :=  /*line :1*/ uint(-e4FXxUi)
	nEbahhuccs :=  /*line :1*/ uint64(1<<v6ctH0 - 1)

	ljUgJLytUS, b0rEW4 := ljUgJLytUS>>v6ctH0, ljUgJLytUS&nEbahhuccs
	acgiVA06Q2, el7K1Fh := acgiVA06Q2>>v6ctH0, acgiVA06Q2&nEbahhuccs
	rAgI2dKQj0, vCHA3DlNy := rAgI2dKQj0>>v6ctH0, rAgI2dKQj0&nEbahhuccs

	annmjaZy := !rOTeKd3Q || vCHA3DlNy > 0
	if rOTeKd3Q && vCHA3DlNy == 0 {
		annmjaZy = fdyaiZ&1 == 0
	}
	if !annmjaZy {
		rAgI2dKQj0--
	}

	mPphrb4IUtdE := false
	if nhLf1v {

		mPphrb4IUtdE = el7K1Fh > 1<<(v6ctH0-1) ||
			(el7K1Fh == 1<<(v6ctH0-1) && acgiVA06Q2&1 == 1)
	} else {

		mPphrb4IUtdE = el7K1Fh>>(v6ctH0-1) == 1
	}

	fSMbqQq26Izs := tbdsmf9aTny && b0rEW4 == 0 && (fdyaiZ&1 == 0)
	if !fSMbqQq26Izs {
		ljUgJLytUS++
	}

	euMadIB4IUT := nhLf1v && el7K1Fh == 0

	 /*line :1*/ bG4mTMNp1NT(bkcwawBry, ljUgJLytUS, acgiVA06Q2, rAgI2dKQj0, euMadIB4IUT, mPphrb4IUtdE)
	bkcwawBry.czr7NK2VrVdf -= oUyevz5E2Vjr
}






func ouR7FiDq(eswrOOpIaF int) int {

	return (eswrOOpIaF * 78913) >> 18
}






func rd0Lc8(eswrOOpIaF int) int {

	return (eswrOOpIaF * 108853) >> 15
}




func s9hB2_0A(fdyaiZ uint64, uZdS176 int, sRaJkQTYfNyc *aL4q0U) (chHT2YtqEP, h95VpqV2z, mRp4tM uint64, e4FXxUi int) {
	if fdyaiZ != 1<<sRaJkQTYfNyc.hCC_Vh || uZdS176 == sRaJkQTYfNyc.oXPRffP4+1- /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh) {

		chHT2YtqEP, h95VpqV2z, mRp4tM = 2*fdyaiZ-1, 2*fdyaiZ, 2*fdyaiZ+1
		e4FXxUi = uZdS176 - 1
		return
	} else {

		chHT2YtqEP, h95VpqV2z, mRp4tM = 4*fdyaiZ-1, 4*fdyaiZ, 4*fdyaiZ+2
		e4FXxUi = uZdS176 - 2
		return
	}
}

func bG4mTMNp1NT(bkcwawBry *cN7vNMaJUh, chHT2YtqEP, h95VpqV2z, mRp4tM uint64,
	euMadIB4IUT, mPphrb4IUtdE bool) {
	w0YtGYWk, vWQWpsZzzk :=  /*line :1*/ dW3DwaRbQUyy(chHT2YtqEP)
	dEtduK, affpNpM :=  /*line :1*/ dW3DwaRbQUyy(h95VpqV2z)
	sFZIDBgbgF, n1Blq3w :=  /*line :1*/ dW3DwaRbQUyy(mRp4tM)
	if sFZIDBgbgF == 0 {

		 /*line :1*/ rrRUYs6Zl(bkcwawBry, vWQWpsZzzk, affpNpM, n1Blq3w, euMadIB4IUT, mPphrb4IUtdE, 8)
	} else if w0YtGYWk < sFZIDBgbgF {

		if vWQWpsZzzk != 0 {
			w0YtGYWk++
		}
		euMadIB4IUT = euMadIB4IUT && affpNpM == 0
		mPphrb4IUtdE = (affpNpM > 5e8) || (affpNpM == 5e8 && mPphrb4IUtdE)
		 /*line :1*/ rrRUYs6Zl(bkcwawBry, w0YtGYWk, dEtduK, sFZIDBgbgF, euMadIB4IUT, mPphrb4IUtdE, 8)
		bkcwawBry.czr7NK2VrVdf += 9
	} else {
		bkcwawBry.hc_0uW85 = 0

		wESnmZiAO :=  /*line :1*/ uint(9)
		for ggqcIE := dEtduK; ggqcIE > 0; {
			zCsD112MLqZf, mOjrrJ87 := ggqcIE/10, ggqcIE%10
			ggqcIE = zCsD112MLqZf
			wESnmZiAO--
			bkcwawBry.dvBTkobK6ZN[wESnmZiAO] =  /*line :1*/ byte(mOjrrJ87 + '0')
		}
		bkcwawBry.dvBTkobK6ZN = bkcwawBry.dvBTkobK6ZN[wESnmZiAO:]
		bkcwawBry.hc_0uW85 =  /*line :1*/ int(9 - wESnmZiAO)

		 /*line :1*/ rrRUYs6Zl(bkcwawBry, vWQWpsZzzk, affpNpM, n1Blq3w,
			euMadIB4IUT, mPphrb4IUtdE, bkcwawBry.hc_0uW85+8)
	}

	for bkcwawBry.hc_0uW85 > 0 && bkcwawBry.dvBTkobK6ZN[bkcwawBry.hc_0uW85-1] == '0' {
		bkcwawBry.hc_0uW85--
	}

	for bkcwawBry.hc_0uW85 > 0 && bkcwawBry.dvBTkobK6ZN[0] == '0' {
		bkcwawBry.hc_0uW85--
		bkcwawBry.czr7NK2VrVdf--
		bkcwawBry.dvBTkobK6ZN = bkcwawBry.dvBTkobK6ZN[1:]
	}
}


func rrRUYs6Zl(bkcwawBry *cN7vNMaJUh, chHT2YtqEP, h95VpqV2z, mRp4tM uint32,
	euMadIB4IUT, mPphrb4IUtdE bool, v6C4hPOD7 int) {
	if mRp4tM == 0 {
		bkcwawBry.czr7NK2VrVdf = v6C4hPOD7 + 1
		return
	}
	fkm7C_rY3E := 0

	gzsmdncA := 0
	for mRp4tM > 0 {

		gtrgUXWjTDv := (chHT2YtqEP + 9) / 10
		uTN3BXbMd, pqyhL9s1Aj := h95VpqV2z/10, h95VpqV2z%10
		dnNosEuwQ := mRp4tM / 10
		if gtrgUXWjTDv > dnNosEuwQ {

			break
		}

		if gtrgUXWjTDv == uTN3BXbMd+1 && uTN3BXbMd < dnNosEuwQ {
			uTN3BXbMd++
			pqyhL9s1Aj = 0
			mPphrb4IUtdE = false
		}
		fkm7C_rY3E++

		euMadIB4IUT = euMadIB4IUT && gzsmdncA == 0
		gzsmdncA =  /*line :1*/ int(pqyhL9s1Aj)
		chHT2YtqEP, h95VpqV2z, mRp4tM = gtrgUXWjTDv, uTN3BXbMd, dnNosEuwQ
	}

	if fkm7C_rY3E > 0 {
		mPphrb4IUtdE = gzsmdncA > 5 ||
			(gzsmdncA == 5 && !euMadIB4IUT) ||
			(gzsmdncA == 5 && euMadIB4IUT && h95VpqV2z&1 == 1)
	}
	if h95VpqV2z < mRp4tM && mPphrb4IUtdE {
		h95VpqV2z++
	}

	v6C4hPOD7 -= fkm7C_rY3E
	ggqcIE := h95VpqV2z
	wESnmZiAO := v6C4hPOD7
	for wESnmZiAO > bkcwawBry.hc_0uW85 {
		zCsD112MLqZf, mOjrrJ87 := ggqcIE/100, ggqcIE%100
		bkcwawBry.dvBTkobK6ZN[wESnmZiAO] = smallsString[2*mOjrrJ87+1]
		bkcwawBry.dvBTkobK6ZN[wESnmZiAO-1] = smallsString[2*mOjrrJ87+0]
		wESnmZiAO -= 2
		ggqcIE = zCsD112MLqZf
	}
	if wESnmZiAO == bkcwawBry.hc_0uW85 {
		bkcwawBry.dvBTkobK6ZN[wESnmZiAO] =  /*line :1*/ byte(ggqcIE + '0')
	}
	bkcwawBry.hc_0uW85 = v6C4hPOD7 + 1
	bkcwawBry.czr7NK2VrVdf = bkcwawBry.hc_0uW85 + fkm7C_rY3E
}











func xXJNqhjxML(csRwMg uint32, e4FXxUi, oUyevz5E2Vjr int) (l7Z79nBFJ uint32, hAsG0B9lF int, iAXjT3 bool) {
	if oUyevz5E2Vjr == 0 {

		return csRwMg << 6, e4FXxUi - 6, true
	}
	if oUyevz5E2Vjr < detailedPowersOfTenMinExp10 || detailedPowersOfTenMaxExp10 < oUyevz5E2Vjr {

		 /*line :1*/ panic("mult64bitPow10: power of 10 is out of range")
	}
	uglgNqnq6z := aIR2N6U[oUyevz5E2Vjr-detailedPowersOfTenMinExp10][1]
	if oUyevz5E2Vjr < 0 {

		uglgNqnq6z += 1
	}
	teBeLUeC, wLrEGVuYyxK :=  /*line :1*/ bits.Mul64( /*line :1*/ uint64(csRwMg), uglgNqnq6z)
	e4FXxUi +=  /*line :1*/ rd0Lc8(oUyevz5E2Vjr) - 63 + 57
	return  /*line :1*/ uint32(teBeLUeC<<7 | wLrEGVuYyxK>>57), e4FXxUi, wLrEGVuYyxK<<7 == 0
}











func gM37tkftIW(csRwMg uint64, e4FXxUi, oUyevz5E2Vjr int) (l7Z79nBFJ uint64, hAsG0B9lF int, iAXjT3 bool) {
	if oUyevz5E2Vjr == 0 {

		return csRwMg << 8, e4FXxUi - 8, true
	}
	if oUyevz5E2Vjr < detailedPowersOfTenMinExp10 || detailedPowersOfTenMaxExp10 < oUyevz5E2Vjr {

		 /*line :1*/ panic("mult128bitPow10: power of 10 is out of range")
	}
	uglgNqnq6z := aIR2N6U[oUyevz5E2Vjr-detailedPowersOfTenMinExp10]
	if oUyevz5E2Vjr < 0 {

		uglgNqnq6z[0] += 1
	}
	e4FXxUi +=  /*line :1*/ rd0Lc8(oUyevz5E2Vjr) - 127 + 119

	h1fQ4wUZCBa1, eUaLl2mmqDM :=  /*line :1*/ bits.Mul64(csRwMg, uglgNqnq6z[0])
	mP_eC48PtpL, umug7Hrz2 :=  /*line :1*/ bits.Mul64(csRwMg, uglgNqnq6z[1])
	g6dcTsse, tquNunL :=  /*line :1*/ bits.Add64(h1fQ4wUZCBa1, umug7Hrz2, 0)
	mP_eC48PtpL += tquNunL
	return mP_eC48PtpL<<9 | g6dcTsse>>55, e4FXxUi, g6dcTsse<<9 == 0 && eUaLl2mmqDM == 0
}

func rtIFITrI(csRwMg uint64, tyBlZruomOL int) bool {
	if csRwMg == 0 {
		return true
	}
	for cDhV_2D := 0; cDhV_2D < tyBlZruomOL; cDhV_2D++ {
		if csRwMg%5 != 0 {
			return false
		}
		csRwMg /= 5
	}
	return true
}



func dW3DwaRbQUyy(eswrOOpIaF uint64) (uint32, uint32) {
	if !host32bit {
		return  /*line :1*/ uint32(eswrOOpIaF / 1e9),  /*line :1*/ uint32(eswrOOpIaF % 1e9)
	}

	teBeLUeC, _ :=  /*line :1*/ bits.Mul64(eswrOOpIaF>>1, 0x89705f4136b4a598)
	oUyevz5E2Vjr := teBeLUeC >> 28
	return  /*line :1*/ uint32(oUyevz5E2Vjr),  /*line :1*/ uint32(eswrOOpIaF - oUyevz5E2Vjr*1e9)
}
