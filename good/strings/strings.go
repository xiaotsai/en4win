//line :1



package fQXlzv

import (
	"internal/bytealg"
	unicode "vGJV8maRK"
	utf8 "CuAc0pPZwf"
)

const maxInt =  /*line :1*/ int(^ /*line :1*/ uint(0) >> 1)




func lejpdo(w4GNKhKtw string, dUrxzl int) []string {
	pW_9Cvdd :=  /*line :1*/ utf8.IaG2lfIQ1LT(w4GNKhKtw)
	if dUrxzl < 0 || dUrxzl > pW_9Cvdd {
		dUrxzl = pW_9Cvdd
	}
	jZJLDz0ImV :=  /*line :1*/ make([]string, dUrxzl)
	for gUiH7N := 0; gUiH7N < dUrxzl-1; gUiH7N++ {
		_, aub89wE2 :=  /*line :1*/ utf8.HVAV7aTqxzg(w4GNKhKtw)
		jZJLDz0ImV[gUiH7N] = w4GNKhKtw[:aub89wE2]
		w4GNKhKtw = w4GNKhKtw[aub89wE2:]
	}
	if dUrxzl > 0 {
		jZJLDz0ImV[dUrxzl-1] = w4GNKhKtw
	}
	return jZJLDz0ImV
}



func L9XssxF(w4GNKhKtw, k2061Z3 string) int {

	if  /*line :1*/ len(k2061Z3) == 0 {
		return  /*line :1*/ utf8.IaG2lfIQ1LT(w4GNKhKtw) + 1
	}
	if  /*line :1*/ len(k2061Z3) == 1 {
		return  /*line :1*/ bytealg.CountString(w4GNKhKtw, k2061Z3[0])
	}
	dUrxzl := 0
	for {
		gUiH7N :=  /*line :1*/ FCsEk_zxUIUY(w4GNKhKtw, k2061Z3)
		if gUiH7N == -1 {
			return dUrxzl
		}
		dUrxzl++
		w4GNKhKtw = w4GNKhKtw[gUiH7N+ /*line :1*/ len(k2061Z3):]
	}
}


func OHUqdLTpfyt(w4GNKhKtw, k2061Z3 string) bool {
	return  /*line :1*/ FCsEk_zxUIUY(w4GNKhKtw, k2061Z3) >= 0
}


func EogZRPaCpV(w4GNKhKtw, tniaLGCuTB string) bool {
	return  /*line :1*/ DD5Se9K(w4GNKhKtw, tniaLGCuTB) >= 0
}


func MLBN4mdkQa3(w4GNKhKtw string, unWb50A rune) bool {
	return  /*line :1*/ JsspDP(w4GNKhKtw, unWb50A) >= 0
}


func RXOhqd(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool) bool {
	return  /*line :1*/ HcW0iiv052D(w4GNKhKtw, vtYnd8LwzVN) >= 0
}


func UgXddvlVDd(w4GNKhKtw, k2061Z3 string) int {
	dUrxzl :=  /*line :1*/ len(k2061Z3)
	switch {
	case dUrxzl == 0:
		return  /*line :1*/ len(w4GNKhKtw)
	case dUrxzl == 1:
		return  /*line :1*/ Owr8r4yL(w4GNKhKtw, k2061Z3[0])
	case dUrxzl ==  /*line :1*/ len(w4GNKhKtw):
		if k2061Z3 == w4GNKhKtw {
			return 0
		}
		return -1
	case dUrxzl >  /*line :1*/ len(w4GNKhKtw):
		return -1
	}

	qpcHxH5TwJus, r5BiemLh :=  /*line :1*/ bytealg.HashStrRev(k2061Z3)
	z6P4_U4 :=  /*line :1*/ len(w4GNKhKtw) - dUrxzl
	var e7P9c3mvU uint32
	for gUiH7N :=  /*line :1*/ len(w4GNKhKtw) - 1; gUiH7N >= z6P4_U4; gUiH7N-- {
		e7P9c3mvU = e7P9c3mvU*bytealg.PrimeRK +  /*line :1*/ uint32(w4GNKhKtw[gUiH7N])
	}
	if e7P9c3mvU == qpcHxH5TwJus && w4GNKhKtw[z6P4_U4:] == k2061Z3 {
		return z6P4_U4
	}
	for gUiH7N := z6P4_U4 - 1; gUiH7N >= 0; gUiH7N-- {
		e7P9c3mvU *= bytealg.PrimeRK
		e7P9c3mvU +=  /*line :1*/ uint32(w4GNKhKtw[gUiH7N])
		e7P9c3mvU -= r5BiemLh *  /*line :1*/ uint32(w4GNKhKtw[gUiH7N+dUrxzl])
		if e7P9c3mvU == qpcHxH5TwJus && w4GNKhKtw[gUiH7N:gUiH7N+dUrxzl] == k2061Z3 {
			return gUiH7N
		}
	}
	return -1
}


func Po8D3DCRZq(w4GNKhKtw string, f4rCFxMKmpo byte) int {
	return  /*line :1*/ bytealg.IndexByteString(w4GNKhKtw, f4rCFxMKmpo)
}





func JsspDP(w4GNKhKtw string, unWb50A rune) int {
	switch {
	case 0 <= unWb50A && unWb50A < utf8.RuneSelf:
		return  /*line :1*/ Po8D3DCRZq(w4GNKhKtw,  /*line :1*/ byte(unWb50A))
	case unWb50A == utf8.RuneError:
		for gUiH7N, unWb50A := range w4GNKhKtw {
			if unWb50A == utf8.RuneError {
				return gUiH7N
			}
		}
		return -1
	case ! /*line :1*/ utf8.DYq33yNk(unWb50A):
		return -1
	default:
		return  /*line :1*/ FCsEk_zxUIUY(w4GNKhKtw,  /*line :1*/ string(unWb50A))
	}
}



func DD5Se9K(w4GNKhKtw, tniaLGCuTB string) int {
	if tniaLGCuTB == "" {

		return -1
	}
	if  /*line :1*/ len(tniaLGCuTB) == 1 {

		unWb50A :=  /*line :1*/ rune(tniaLGCuTB[0])
		if unWb50A >= utf8.RuneSelf {
			unWb50A = utf8.RuneError
		}
		return  /*line :1*/ JsspDP(w4GNKhKtw, unWb50A)
	}
	if  /*line :1*/ len(w4GNKhKtw) > 8 {
		if xK9AxdhlbHn, uMmgxk7Vk :=  /*line :1*/ lTTp_SFS6(tniaLGCuTB); uMmgxk7Vk {
			for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
				if  /*line :1*/ xK9AxdhlbHn.y3qjcJj(w4GNKhKtw[gUiH7N]) {
					return gUiH7N
				}
			}
			return -1
		}
	}
	for gUiH7N, f4rCFxMKmpo := range w4GNKhKtw {
		if  /*line :1*/ JsspDP(tniaLGCuTB, f4rCFxMKmpo) >= 0 {
			return gUiH7N
		}
	}
	return -1
}




func CbawhSupGNAH(w4GNKhKtw, tniaLGCuTB string) int {
	if tniaLGCuTB == "" {

		return -1
	}
	if  /*line :1*/ len(w4GNKhKtw) == 1 {
		iySPOw2h :=  /*line :1*/ rune(w4GNKhKtw[0])
		if iySPOw2h >= utf8.RuneSelf {
			iySPOw2h = utf8.RuneError
		}
		if  /*line :1*/ JsspDP(tniaLGCuTB, iySPOw2h) >= 0 {
			return 0
		}
		return -1
	}
	if  /*line :1*/ len(w4GNKhKtw) > 8 {
		if xK9AxdhlbHn, uMmgxk7Vk :=  /*line :1*/ lTTp_SFS6(tniaLGCuTB); uMmgxk7Vk {
			for gUiH7N :=  /*line :1*/ len(w4GNKhKtw) - 1; gUiH7N >= 0; gUiH7N-- {
				if  /*line :1*/ xK9AxdhlbHn.y3qjcJj(w4GNKhKtw[gUiH7N]) {
					return gUiH7N
				}
			}
			return -1
		}
	}
	if  /*line :1*/ len(tniaLGCuTB) == 1 {
		iySPOw2h :=  /*line :1*/ rune(tniaLGCuTB[0])
		if iySPOw2h >= utf8.RuneSelf {
			iySPOw2h = utf8.RuneError
		}
		for gUiH7N :=  /*line :1*/ len(w4GNKhKtw); gUiH7N > 0; {
			unWb50A, aub89wE2 :=  /*line :1*/ utf8.G3GHxYa4Lm(w4GNKhKtw[:gUiH7N])
			gUiH7N -= aub89wE2
			if iySPOw2h == unWb50A {
				return gUiH7N
			}
		}
		return -1
	}
	for gUiH7N :=  /*line :1*/ len(w4GNKhKtw); gUiH7N > 0; {
		unWb50A, aub89wE2 :=  /*line :1*/ utf8.G3GHxYa4Lm(w4GNKhKtw[:gUiH7N])
		gUiH7N -= aub89wE2
		if  /*line :1*/ JsspDP(tniaLGCuTB, unWb50A) >= 0 {
			return gUiH7N
		}
	}
	return -1
}


func Owr8r4yL(w4GNKhKtw string, f4rCFxMKmpo byte) int {
	for gUiH7N :=  /*line :1*/ len(w4GNKhKtw) - 1; gUiH7N >= 0; gUiH7N-- {
		if w4GNKhKtw[gUiH7N] == f4rCFxMKmpo {
			return gUiH7N
		}
	}
	return -1
}



func iBxAe4Ef(w4GNKhKtw, mDyLGdn7 string, jehsdXXy_tSp, dUrxzl int) []string {
	if dUrxzl == 0 {
		return nil
	}
	if mDyLGdn7 == "" {
		return  /*line :1*/ lejpdo(w4GNKhKtw, dUrxzl)
	}
	if dUrxzl < 0 {
		dUrxzl =  /*line :1*/ L9XssxF(w4GNKhKtw, mDyLGdn7) + 1
	}

	if dUrxzl >  /*line :1*/ len(w4GNKhKtw)+1 {
		dUrxzl =  /*line :1*/ len(w4GNKhKtw) + 1
	}
	jZJLDz0ImV :=  /*line :1*/ make([]string, dUrxzl)
	dUrxzl--
	gUiH7N := 0
	for gUiH7N < dUrxzl {
		ta2_ET :=  /*line :1*/ FCsEk_zxUIUY(w4GNKhKtw, mDyLGdn7)
		if ta2_ET < 0 {
			break
		}
		jZJLDz0ImV[gUiH7N] = w4GNKhKtw[:ta2_ET+jehsdXXy_tSp]
		w4GNKhKtw = w4GNKhKtw[ta2_ET+ /*line :1*/ len(mDyLGdn7):]
		gUiH7N++
	}
	jZJLDz0ImV[gUiH7N] = w4GNKhKtw
	return jZJLDz0ImV[:gUiH7N+1]
}














func LX3sIwU6(w4GNKhKtw, mDyLGdn7 string, dUrxzl int) []string {
	return  /*line :1*/ iBxAe4Ef(w4GNKhKtw, mDyLGdn7, 0, dUrxzl)
}












func NyEFl_GtnJT(w4GNKhKtw, mDyLGdn7 string, dUrxzl int) []string {
	return  /*line :1*/ iBxAe4Ef(w4GNKhKtw, mDyLGdn7,  /*line :1*/ len(mDyLGdn7), dUrxzl)
}













func BmJ7uIgX1gL(w4GNKhKtw, mDyLGdn7 string) []string	{ return  /*line :1*/ iBxAe4Ef(w4GNKhKtw, mDyLGdn7, 0, -1) }











func Iaw6U6(w4GNKhKtw, mDyLGdn7 string) []string {
	return  /*line :1*/ iBxAe4Ef(w4GNKhKtw, mDyLGdn7,  /*line :1*/ len(mDyLGdn7), -1)
}

var m7ckqgSsIBmT = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}




func JSH_h9JePHEv(w4GNKhKtw string) []string {

	dUrxzl := 0
	jn_I3aJv := 1

	enLxysm9mRUl :=  /*line :1*/ uint8(0)
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
		unWb50A := w4GNKhKtw[gUiH7N]
		enLxysm9mRUl |= unWb50A
		lSU7byBPSQqP :=  /*line :1*/ int(m7ckqgSsIBmT[unWb50A])
		dUrxzl += jn_I3aJv & ^lSU7byBPSQqP
		jn_I3aJv = lSU7byBPSQqP
	}

	if enLxysm9mRUl >= utf8.RuneSelf {

		return  /*line :1*/ W6bwbYTrkD(w4GNKhKtw, unicode.Mzbe7SW3gm0k)
	}

	jZJLDz0ImV :=  /*line :1*/ make([]string, dUrxzl)
	dMw9Q6aEq := 0
	gLLJaKzm := 0
	gUiH7N := 0

	for gUiH7N <  /*line :1*/ len(w4GNKhKtw) && m7ckqgSsIBmT[w4GNKhKtw[gUiH7N]] != 0 {
		gUiH7N++
	}
	gLLJaKzm = gUiH7N
	for gUiH7N <  /*line :1*/ len(w4GNKhKtw) {
		if m7ckqgSsIBmT[w4GNKhKtw[gUiH7N]] == 0 {
			gUiH7N++
			continue
		}
		jZJLDz0ImV[dMw9Q6aEq] = w4GNKhKtw[gLLJaKzm:gUiH7N]
		dMw9Q6aEq++
		gUiH7N++

		for gUiH7N <  /*line :1*/ len(w4GNKhKtw) && m7ckqgSsIBmT[w4GNKhKtw[gUiH7N]] != 0 {
			gUiH7N++
		}
		gLLJaKzm = gUiH7N
	}
	if gLLJaKzm <  /*line :1*/ len(w4GNKhKtw) {
		jZJLDz0ImV[dMw9Q6aEq] = w4GNKhKtw[gLLJaKzm:]
	}
	return jZJLDz0ImV
}







func W6bwbYTrkD(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool) []string {
	
	
	type lPQSgq struct {
		t58suHObXgU9	int
		cvWVr_i6gkp	int
	}
	ha5bT861x :=  /*line :1*/ make([]lPQSgq, 0, 32)

	v_VDLwK := -1
	for gpAwFStAAVQ, z_ZqdFDgS := range w4GNKhKtw {
		if  /*line :1*/ vtYnd8LwzVN(z_ZqdFDgS) {
			if v_VDLwK >= 0 {
				ha5bT861x =  /*line :1*/ append(ha5bT861x, lPQSgq{v_VDLwK, gpAwFStAAVQ})

				v_VDLwK = ^v_VDLwK
			}
		} else {
			if v_VDLwK < 0 {
				v_VDLwK = gpAwFStAAVQ
			}
		}
	}

	if v_VDLwK >= 0 {
		ha5bT861x =  /*line :1*/ append(ha5bT861x, lPQSgq{v_VDLwK,  /*line :1*/ len(w4GNKhKtw)})
	}

	jZJLDz0ImV :=  /*line :1*/ make([]string,  /*line :1*/ len(ha5bT861x))
	for gUiH7N, lPQSgq := range ha5bT861x {
		jZJLDz0ImV[gUiH7N] = w4GNKhKtw[lPQSgq.t58suHObXgU9:lPQSgq.cvWVr_i6gkp]
	}

	return jZJLDz0ImV
}



func ZTisTA(usumrj7otiDz []string, mDyLGdn7 string) string {
	switch  /*line :1*/ len(usumrj7otiDz) {
	case 0:
		return ""
	case 1:
		return usumrj7otiDz[0]
	}

	var dUrxzl int
	if  /*line :1*/ len(mDyLGdn7) > 0 {
		if  /*line :1*/ len(mDyLGdn7) >= maxInt/( /*line :1*/ len(usumrj7otiDz)-1) {
			 /*line :1*/ panic("strings: Join output length overflow")
		}
		dUrxzl +=  /*line :1*/ len(mDyLGdn7) * ( /*line :1*/ len(usumrj7otiDz) - 1)
	}
	for _, nzcQrgl := range usumrj7otiDz {
		if  /*line :1*/ len(nzcQrgl) > maxInt-dUrxzl {
			 /*line :1*/ panic("strings: Join output length overflow")
		}
		dUrxzl +=  /*line :1*/ len(nzcQrgl)
	}

	var naCMgzayVI CKrNlN3otWSF
	 /*line :1*/ naCMgzayVI.Grow(dUrxzl)
	 /*line :1*/ naCMgzayVI.WriteString(usumrj7otiDz[0])
	for _, w4GNKhKtw := range usumrj7otiDz[1:] {
		 /*line :1*/ naCMgzayVI.WriteString(mDyLGdn7)
		 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw)
	}
	return  /*line :1*/ naCMgzayVI.String()
}


func E8ZCzfgw(w4GNKhKtw, aC71QFDwvE0T string) bool {
	return  /*line :1*/ len(w4GNKhKtw) >=  /*line :1*/ len(aC71QFDwvE0T) && w4GNKhKtw[0: /*line :1*/ len(aC71QFDwvE0T)] == aC71QFDwvE0T
}


func RZpWO5X7(w4GNKhKtw, triI9zMn4AM string) bool {
	return  /*line :1*/ len(w4GNKhKtw) >=  /*line :1*/ len(triI9zMn4AM) && w4GNKhKtw[ /*line :1*/ len(w4GNKhKtw)- /*line :1*/ len(triI9zMn4AM):] == triI9zMn4AM
}




func GXgOuM(b2bCzp_awe func(rune) rune, w4GNKhKtw string) string {

	
	
	var naCMgzayVI CKrNlN3otWSF

	for gUiH7N, f4rCFxMKmpo := range w4GNKhKtw {
		unWb50A :=  /*line :1*/ b2bCzp_awe(f4rCFxMKmpo)
		if unWb50A == f4rCFxMKmpo && f4rCFxMKmpo != utf8.RuneError {
			continue
		}

		var sOoCRO int
		if f4rCFxMKmpo == utf8.RuneError {
			f4rCFxMKmpo, sOoCRO =  /*line :1*/ utf8.HVAV7aTqxzg(w4GNKhKtw[gUiH7N:])
			if sOoCRO != 1 && unWb50A == f4rCFxMKmpo {
				continue
			}
		} else {
			sOoCRO =  /*line :1*/ utf8.YrF2oG995j9(f4rCFxMKmpo)
		}

		 /*line :1*/ naCMgzayVI.Grow( /*line :1*/ len(w4GNKhKtw) + utf8.UTFMax)
		 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[:gUiH7N])
		if unWb50A >= 0 {
			 /*line :1*/ naCMgzayVI.WriteRune(unWb50A)
		}

		w4GNKhKtw = w4GNKhKtw[gUiH7N+sOoCRO:]
		break
	}

	if  /*line :1*/ naCMgzayVI.Cap() == 0 {
		return w4GNKhKtw
	}

	for _, f4rCFxMKmpo := range w4GNKhKtw {
		unWb50A :=  /*line :1*/ b2bCzp_awe(f4rCFxMKmpo)

		if unWb50A >= 0 {

			if unWb50A < utf8.RuneSelf {
				 /*line :1*/ naCMgzayVI.WriteByte( /*line :1*/ byte(unWb50A))
			} else {

				 /*line :1*/ naCMgzayVI.WriteRune(unWb50A)
			}
		}
	}

	return  /*line :1*/ naCMgzayVI.String()
}





func FAfX4kvYP7Z(w4GNKhKtw string, oxCupp5J int) string {
	switch oxCupp5J {
	case 0:
		return ""
	case 1:
		return w4GNKhKtw
	}

	if oxCupp5J < 0 {
		 /*line :1*/ panic("strings: negative Repeat count")
	}
	if  /*line :1*/ len(w4GNKhKtw) >= maxInt/oxCupp5J {
		 /*line :1*/ panic("strings: Repeat output length overflow")
	}
	dUrxzl :=  /*line :1*/ len(w4GNKhKtw) * oxCupp5J

	if  /*line :1*/ len(w4GNKhKtw) == 0 {
		return ""
	}

	
	
	
	
	
	
	
	
	
	
	const chunkLimit = 8 * 1024
	gh7IVvNLFsM := dUrxzl
	if dUrxzl > chunkLimit {
		gh7IVvNLFsM = chunkLimit /  /*line :1*/ len(w4GNKhKtw) *  /*line :1*/ len(w4GNKhKtw)
		if gh7IVvNLFsM == 0 {
			gh7IVvNLFsM =  /*line :1*/ len(w4GNKhKtw)
		}
	}

	var naCMgzayVI CKrNlN3otWSF
	 /*line :1*/ naCMgzayVI.Grow(dUrxzl)
	 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw)
	for  /*line :1*/ naCMgzayVI.Len() < dUrxzl {
		bK699s5U := dUrxzl -  /*line :1*/ naCMgzayVI.Len()
		if bK699s5U >  /*line :1*/ naCMgzayVI.Len() {
			bK699s5U =  /*line :1*/ naCMgzayVI.Len()
		}
		if bK699s5U > gh7IVvNLFsM {
			bK699s5U = gh7IVvNLFsM
		}
		 /*line :1*/ naCMgzayVI.WriteString( /*line :1*/ naCMgzayVI.String()[:bK699s5U])
	}
	return  /*line :1*/ naCMgzayVI.String()
}


func Z7tZslPsj(w4GNKhKtw string) string {
	uMmgxk7Vk, m1ZaBsk := true, false
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
		f4rCFxMKmpo := w4GNKhKtw[gUiH7N]
		if f4rCFxMKmpo >= utf8.RuneSelf {
			uMmgxk7Vk = false
			break
		}
		m1ZaBsk = m1ZaBsk || ('a' <= f4rCFxMKmpo && f4rCFxMKmpo <= 'z')
	}

	if uMmgxk7Vk {
		if !m1ZaBsk {
			return w4GNKhKtw
		}
		var (
			naCMgzayVI	CKrNlN3otWSF
			vhYnNop	int
		)
		 /*line :1*/ naCMgzayVI.Grow( /*line :1*/ len(w4GNKhKtw))
		for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
			f4rCFxMKmpo := w4GNKhKtw[gUiH7N]
			if 'a' <= f4rCFxMKmpo && f4rCFxMKmpo <= 'z' {
				f4rCFxMKmpo -= 'a' - 'A'
				if vhYnNop < gUiH7N {
					 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[vhYnNop:gUiH7N])
				}
				 /*line :1*/ naCMgzayVI.WriteByte(f4rCFxMKmpo)
				vhYnNop = gUiH7N + 1
			}
		}
		if vhYnNop <  /*line :1*/ len(w4GNKhKtw) {
			 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[vhYnNop:])
		}
		return  /*line :1*/ naCMgzayVI.String()
	}
	return  /*line :1*/ GXgOuM(unicode.HJD8X_JbK, w4GNKhKtw)
}


func WKrzOP(w4GNKhKtw string) string {
	uMmgxk7Vk, cf_Vpxgf := true, false
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
		f4rCFxMKmpo := w4GNKhKtw[gUiH7N]
		if f4rCFxMKmpo >= utf8.RuneSelf {
			uMmgxk7Vk = false
			break
		}
		cf_Vpxgf = cf_Vpxgf || ('A' <= f4rCFxMKmpo && f4rCFxMKmpo <= 'Z')
	}

	if uMmgxk7Vk {
		if !cf_Vpxgf {
			return w4GNKhKtw
		}
		var (
			naCMgzayVI	CKrNlN3otWSF
			vhYnNop	int
		)
		 /*line :1*/ naCMgzayVI.Grow( /*line :1*/ len(w4GNKhKtw))
		for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
			f4rCFxMKmpo := w4GNKhKtw[gUiH7N]
			if 'A' <= f4rCFxMKmpo && f4rCFxMKmpo <= 'Z' {
				f4rCFxMKmpo += 'a' - 'A'
				if vhYnNop < gUiH7N {
					 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[vhYnNop:gUiH7N])
				}
				 /*line :1*/ naCMgzayVI.WriteByte(f4rCFxMKmpo)
				vhYnNop = gUiH7N + 1
			}
		}
		if vhYnNop <  /*line :1*/ len(w4GNKhKtw) {
			 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[vhYnNop:])
		}
		return  /*line :1*/ naCMgzayVI.String()
	}
	return  /*line :1*/ GXgOuM(unicode.OLZ8W2, w4GNKhKtw)
}



func DXJR2pp1A4Cx(w4GNKhKtw string) string	{ return  /*line :1*/ GXgOuM(unicode.QgPuNi8kYiC, w4GNKhKtw) }



func HvRvAKalkmNP(f4rCFxMKmpo unicode.KmthvGp3gyLu, w4GNKhKtw string) string {
	return  /*line :1*/ GXgOuM(f4rCFxMKmpo.ToUpper, w4GNKhKtw)
}



func D5qvOM7SNqj(f4rCFxMKmpo unicode.KmthvGp3gyLu, w4GNKhKtw string) string {
	return  /*line :1*/ GXgOuM(f4rCFxMKmpo.ToLower, w4GNKhKtw)
}



func BapDkWnyqLjo(f4rCFxMKmpo unicode.KmthvGp3gyLu, w4GNKhKtw string) string {
	return  /*line :1*/ GXgOuM(f4rCFxMKmpo.ToTitle, w4GNKhKtw)
}



func ODIKjuRRh0jT(w4GNKhKtw, oIyrsS1i1 string) string {
	var naCMgzayVI CKrNlN3otWSF

	for gUiH7N, f4rCFxMKmpo := range w4GNKhKtw {
		if f4rCFxMKmpo != utf8.RuneError {
			continue
		}

		_, aEjXSbFlLK :=  /*line :1*/ utf8.HVAV7aTqxzg(w4GNKhKtw[gUiH7N:])
		if aEjXSbFlLK == 1 {
			 /*line :1*/ naCMgzayVI.Grow( /*line :1*/ len(w4GNKhKtw) +  /*line :1*/ len(oIyrsS1i1))
			 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[:gUiH7N])
			w4GNKhKtw = w4GNKhKtw[gUiH7N:]
			break
		}
	}

	if  /*line :1*/ naCMgzayVI.Cap() == 0 {
		return w4GNKhKtw
	}

	jOQt3FAMXW := false
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); {
		f4rCFxMKmpo := w4GNKhKtw[gUiH7N]
		if f4rCFxMKmpo < utf8.RuneSelf {
			gUiH7N++
			jOQt3FAMXW = false
			 /*line :1*/ naCMgzayVI.WriteByte(f4rCFxMKmpo)
			continue
		}
		_, aEjXSbFlLK :=  /*line :1*/ utf8.HVAV7aTqxzg(w4GNKhKtw[gUiH7N:])
		if aEjXSbFlLK == 1 {
			gUiH7N++
			if !jOQt3FAMXW {
				jOQt3FAMXW = true
				 /*line :1*/ naCMgzayVI.WriteString(oIyrsS1i1)
			}
			continue
		}
		jOQt3FAMXW = false
		 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[gUiH7N : gUiH7N+aEjXSbFlLK])
		gUiH7N += aEjXSbFlLK
	}

	return  /*line :1*/ naCMgzayVI.String()
}



func siU1wIOA(unWb50A rune) bool {

	if unWb50A <= 0x7F {
		switch {
		case '0' <= unWb50A && unWb50A <= '9':
			return false
		case 'a' <= unWb50A && unWb50A <= 'z':
			return false
		case 'A' <= unWb50A && unWb50A <= 'Z':
			return false
		case unWb50A == '_':
			return false
		}
		return true
	}

	if  /*line :1*/ unicode.Nyu9Lcgu(unWb50A) ||  /*line :1*/ unicode.FquwNEzWV(unWb50A) {
		return false
	}

	return  /*line :1*/ unicode.Mzbe7SW3gm0k(unWb50A)
}






func Vp5vI9ZaIL(w4GNKhKtw string) string {

	vnGWPrR7A := ' '
	return  /*line :1*/ GXgOuM(
		func(unWb50A rune) rune {
			if  /*line :1*/ siU1wIOA(vnGWPrR7A) {
				vnGWPrR7A = unWb50A
				return  /*line :1*/ unicode.QgPuNi8kYiC(unWb50A)
			}
			vnGWPrR7A = unWb50A
			return unWb50A
		},
		w4GNKhKtw)
}



func BZoJ_ybs(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool) string {
	gUiH7N :=  /*line :1*/ rN8saFToG(w4GNKhKtw, vtYnd8LwzVN, false)
	if gUiH7N == -1 {
		return ""
	}
	return w4GNKhKtw[gUiH7N:]
}



func GS6vnza_3(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool) string {
	gUiH7N :=  /*line :1*/ lKxrgb_Fv(w4GNKhKtw, vtYnd8LwzVN, false)
	if gUiH7N >= 0 && w4GNKhKtw[gUiH7N] >= utf8.RuneSelf {
		_, aEjXSbFlLK :=  /*line :1*/ utf8.HVAV7aTqxzg(w4GNKhKtw[gUiH7N:])
		gUiH7N += aEjXSbFlLK
	} else {
		gUiH7N++
	}
	return w4GNKhKtw[0:gUiH7N]
}



func Ebkbg2SD(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool) string {
	return  /*line :1*/ GS6vnza_3( /*line :1*/ BZoJ_ybs(w4GNKhKtw, vtYnd8LwzVN), vtYnd8LwzVN)
}



func HcW0iiv052D(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool) int {
	return  /*line :1*/ rN8saFToG(w4GNKhKtw, vtYnd8LwzVN, true)
}



func EZwqZZ4A(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool) int {
	return  /*line :1*/ lKxrgb_Fv(w4GNKhKtw, vtYnd8LwzVN, true)
}




func rN8saFToG(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool, fudH36it0 bool) int {
	for gUiH7N, unWb50A := range w4GNKhKtw {
		if  /*line :1*/ vtYnd8LwzVN(unWb50A) == fudH36it0 {
			return gUiH7N
		}
	}
	return -1
}




func lKxrgb_Fv(w4GNKhKtw string, vtYnd8LwzVN func(rune) bool, fudH36it0 bool) int {
	for gUiH7N :=  /*line :1*/ len(w4GNKhKtw); gUiH7N > 0; {
		unWb50A, aub89wE2 :=  /*line :1*/ utf8.G3GHxYa4Lm(w4GNKhKtw[0:gUiH7N])
		gUiH7N -= aub89wE2
		if  /*line :1*/ vtYnd8LwzVN(unWb50A) == fudH36it0 {
			return gUiH7N
		}
	}
	return -1
}









type hI4PbqjY6T [8]uint32



func lTTp_SFS6(tniaLGCuTB string) (xK9AxdhlbHn hI4PbqjY6T, ocyp5bYUYCEP bool) {
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(tniaLGCuTB); gUiH7N++ {
		f4rCFxMKmpo := tniaLGCuTB[gUiH7N]
		if f4rCFxMKmpo >= utf8.RuneSelf {
			return xK9AxdhlbHn, false
		}
		xK9AxdhlbHn[f4rCFxMKmpo/32] |= 1 << (f4rCFxMKmpo % 32)
	}
	return xK9AxdhlbHn, true
}


func (xK9AxdhlbHn *hI4PbqjY6T) y3qjcJj(f4rCFxMKmpo byte) bool {
	return (xK9AxdhlbHn[f4rCFxMKmpo/32] & (1 << (f4rCFxMKmpo % 32))) != 0
}



func TNj1WoSMVoxw(w4GNKhKtw, aOAz9Kx string) string {
	if w4GNKhKtw == "" || aOAz9Kx == "" {
		return w4GNKhKtw
	}
	if  /*line :1*/ len(aOAz9Kx) == 1 && aOAz9Kx[0] < utf8.RuneSelf {
		return  /*line :1*/ xKOw6qoDLdIO( /*line :1*/ wCmRhS(w4GNKhKtw, aOAz9Kx[0]), aOAz9Kx[0])
	}
	if xK9AxdhlbHn, ocyp5bYUYCEP :=  /*line :1*/ lTTp_SFS6(aOAz9Kx); ocyp5bYUYCEP {
		return  /*line :1*/ bfaTANeTdYor( /*line :1*/ xghRN1(w4GNKhKtw, &xK9AxdhlbHn), &xK9AxdhlbHn)
	}
	return  /*line :1*/ lI7xlUx5zaN( /*line :1*/ jdThP2(w4GNKhKtw, aOAz9Kx), aOAz9Kx)
}





func UOOZM1(w4GNKhKtw, aOAz9Kx string) string {
	if w4GNKhKtw == "" || aOAz9Kx == "" {
		return w4GNKhKtw
	}
	if  /*line :1*/ len(aOAz9Kx) == 1 && aOAz9Kx[0] < utf8.RuneSelf {
		return  /*line :1*/ xKOw6qoDLdIO(w4GNKhKtw, aOAz9Kx[0])
	}
	if xK9AxdhlbHn, ocyp5bYUYCEP :=  /*line :1*/ lTTp_SFS6(aOAz9Kx); ocyp5bYUYCEP {
		return  /*line :1*/ bfaTANeTdYor(w4GNKhKtw, &xK9AxdhlbHn)
	}
	return  /*line :1*/ lI7xlUx5zaN(w4GNKhKtw, aOAz9Kx)
}

func xKOw6qoDLdIO(w4GNKhKtw string, f4rCFxMKmpo byte) string {
	for  /*line :1*/ len(w4GNKhKtw) > 0 && w4GNKhKtw[0] == f4rCFxMKmpo {
		w4GNKhKtw = w4GNKhKtw[1:]
	}
	return w4GNKhKtw
}

func bfaTANeTdYor(w4GNKhKtw string, xK9AxdhlbHn *hI4PbqjY6T) string {
	for  /*line :1*/ len(w4GNKhKtw) > 0 {
		if ! /*line :1*/ xK9AxdhlbHn.y3qjcJj(w4GNKhKtw[0]) {
			break
		}
		w4GNKhKtw = w4GNKhKtw[1:]
	}
	return w4GNKhKtw
}

func lI7xlUx5zaN(w4GNKhKtw, aOAz9Kx string) string {
	for  /*line :1*/ len(w4GNKhKtw) > 0 {
		unWb50A, dUrxzl :=  /*line :1*/ rune(w4GNKhKtw[0]), 1
		if unWb50A >= utf8.RuneSelf {
			unWb50A, dUrxzl =  /*line :1*/ utf8.HVAV7aTqxzg(w4GNKhKtw)
		}
		if ! /*line :1*/ MLBN4mdkQa3(aOAz9Kx, unWb50A) {
			break
		}
		w4GNKhKtw = w4GNKhKtw[dUrxzl:]
	}
	return w4GNKhKtw
}





func Nzb734WiVLE(w4GNKhKtw, aOAz9Kx string) string {
	if w4GNKhKtw == "" || aOAz9Kx == "" {
		return w4GNKhKtw
	}
	if  /*line :1*/ len(aOAz9Kx) == 1 && aOAz9Kx[0] < utf8.RuneSelf {
		return  /*line :1*/ wCmRhS(w4GNKhKtw, aOAz9Kx[0])
	}
	if xK9AxdhlbHn, ocyp5bYUYCEP :=  /*line :1*/ lTTp_SFS6(aOAz9Kx); ocyp5bYUYCEP {
		return  /*line :1*/ xghRN1(w4GNKhKtw, &xK9AxdhlbHn)
	}
	return  /*line :1*/ jdThP2(w4GNKhKtw, aOAz9Kx)
}

func wCmRhS(w4GNKhKtw string, f4rCFxMKmpo byte) string {
	for  /*line :1*/ len(w4GNKhKtw) > 0 && w4GNKhKtw[ /*line :1*/ len(w4GNKhKtw)-1] == f4rCFxMKmpo {
		w4GNKhKtw = w4GNKhKtw[: /*line :1*/ len(w4GNKhKtw)-1]
	}
	return w4GNKhKtw
}

func xghRN1(w4GNKhKtw string, xK9AxdhlbHn *hI4PbqjY6T) string {
	for  /*line :1*/ len(w4GNKhKtw) > 0 {
		if ! /*line :1*/ xK9AxdhlbHn.y3qjcJj(w4GNKhKtw[ /*line :1*/ len(w4GNKhKtw)-1]) {
			break
		}
		w4GNKhKtw = w4GNKhKtw[: /*line :1*/ len(w4GNKhKtw)-1]
	}
	return w4GNKhKtw
}

func jdThP2(w4GNKhKtw, aOAz9Kx string) string {
	for  /*line :1*/ len(w4GNKhKtw) > 0 {
		unWb50A, dUrxzl :=  /*line :1*/ rune(w4GNKhKtw[ /*line :1*/ len(w4GNKhKtw)-1]), 1
		if unWb50A >= utf8.RuneSelf {
			unWb50A, dUrxzl =  /*line :1*/ utf8.G3GHxYa4Lm(w4GNKhKtw)
		}
		if ! /*line :1*/ MLBN4mdkQa3(aOAz9Kx, unWb50A) {
			break
		}
		w4GNKhKtw = w4GNKhKtw[: /*line :1*/ len(w4GNKhKtw)-dUrxzl]
	}
	return w4GNKhKtw
}



func H8cdur8LCY(w4GNKhKtw string) string {

	v_VDLwK := 0
	for ; v_VDLwK <  /*line :1*/ len(w4GNKhKtw); v_VDLwK++ {
		f4rCFxMKmpo := w4GNKhKtw[v_VDLwK]
		if f4rCFxMKmpo >= utf8.RuneSelf {

			return  /*line :1*/ Ebkbg2SD(w4GNKhKtw[v_VDLwK:], unicode.Mzbe7SW3gm0k)
		}
		if m7ckqgSsIBmT[f4rCFxMKmpo] == 0 {
			break
		}
	}

	cWqdijMv5A :=  /*line :1*/ len(w4GNKhKtw)
	for ; cWqdijMv5A > v_VDLwK; cWqdijMv5A-- {
		f4rCFxMKmpo := w4GNKhKtw[cWqdijMv5A-1]
		if f4rCFxMKmpo >= utf8.RuneSelf {

			return  /*line :1*/ GS6vnza_3(w4GNKhKtw[v_VDLwK:cWqdijMv5A], unicode.Mzbe7SW3gm0k)
		}
		if m7ckqgSsIBmT[f4rCFxMKmpo] == 0 {
			break
		}
	}

	return w4GNKhKtw[v_VDLwK:cWqdijMv5A]
}



func TqmzigPVrfn(w4GNKhKtw, aC71QFDwvE0T string) string {
	if  /*line :1*/ E8ZCzfgw(w4GNKhKtw, aC71QFDwvE0T) {
		return w4GNKhKtw[ /*line :1*/ len(aC71QFDwvE0T):]
	}
	return w4GNKhKtw
}



func CLQXk6(w4GNKhKtw, triI9zMn4AM string) string {
	if  /*line :1*/ RZpWO5X7(w4GNKhKtw, triI9zMn4AM) {
		return w4GNKhKtw[: /*line :1*/ len(w4GNKhKtw)- /*line :1*/ len(triI9zMn4AM)]
	}
	return w4GNKhKtw
}







func K0w1XKcYee9(w4GNKhKtw, fqJMOHuycS, wKpC79A9cHC string, dUrxzl int) string {
	if fqJMOHuycS == wKpC79A9cHC || dUrxzl == 0 {
		return w4GNKhKtw
	}

	if ta2_ET :=  /*line :1*/ L9XssxF(w4GNKhKtw, fqJMOHuycS); ta2_ET == 0 {
		return w4GNKhKtw
	} else if dUrxzl < 0 || ta2_ET < dUrxzl {
		dUrxzl = ta2_ET
	}

	
	var naCMgzayVI CKrNlN3otWSF
	 /*line :1*/ naCMgzayVI.Grow( /*line :1*/ len(w4GNKhKtw) + dUrxzl*( /*line :1*/ len(wKpC79A9cHC)- /*line :1*/ len(fqJMOHuycS)))
	v_VDLwK := 0
	for gUiH7N := 0; gUiH7N < dUrxzl; gUiH7N++ {
		xXcc3ZyUB3kD := v_VDLwK
		if  /*line :1*/ len(fqJMOHuycS) == 0 {
			if gUiH7N > 0 {
				_, aEjXSbFlLK :=  /*line :1*/ utf8.HVAV7aTqxzg(w4GNKhKtw[v_VDLwK:])
				xXcc3ZyUB3kD += aEjXSbFlLK
			}
		} else {
			xXcc3ZyUB3kD +=  /*line :1*/ FCsEk_zxUIUY(w4GNKhKtw[v_VDLwK:], fqJMOHuycS)
		}
		 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[v_VDLwK:xXcc3ZyUB3kD])
		 /*line :1*/ naCMgzayVI.WriteString(wKpC79A9cHC)
		v_VDLwK = xXcc3ZyUB3kD +  /*line :1*/ len(fqJMOHuycS)
	}
	 /*line :1*/ naCMgzayVI.WriteString(w4GNKhKtw[v_VDLwK:])
	return  /*line :1*/ naCMgzayVI.String()
}






func Tq1MwZX3Za(w4GNKhKtw, fqJMOHuycS, wKpC79A9cHC string) string {
	return  /*line :1*/ K0w1XKcYee9(w4GNKhKtw, fqJMOHuycS, wKpC79A9cHC, -1)
}




func UPM6mljQncm9(w4GNKhKtw, az1oIQ9_Q6N string) bool {

	gUiH7N := 0
	for ; gUiH7N <  /*line :1*/ len(w4GNKhKtw) && gUiH7N <  /*line :1*/ len(az1oIQ9_Q6N); gUiH7N++ {
		cBR9L5k := w4GNKhKtw[gUiH7N]
		hIThvPiOWy := az1oIQ9_Q6N[gUiH7N]
		if cBR9L5k|hIThvPiOWy >= utf8.RuneSelf {
			goto hasUnicode
		}

		if hIThvPiOWy == cBR9L5k {
			continue
		}

		if hIThvPiOWy < cBR9L5k {
			hIThvPiOWy, cBR9L5k = cBR9L5k, hIThvPiOWy
		}

		if 'A' <= cBR9L5k && cBR9L5k <= 'Z' && hIThvPiOWy == cBR9L5k+'a'-'A' {
			continue
		}
		return false
	}

	return  /*line :1*/ len(w4GNKhKtw) ==  /*line :1*/ len(az1oIQ9_Q6N)

hasUnicode:
	w4GNKhKtw = w4GNKhKtw[gUiH7N:]
	az1oIQ9_Q6N = az1oIQ9_Q6N[gUiH7N:]
	for _, cBR9L5k := range w4GNKhKtw {

		if  /*line :1*/ len(az1oIQ9_Q6N) == 0 {
			return false
		}

		
		var hIThvPiOWy rune
		if az1oIQ9_Q6N[0] < utf8.RuneSelf {
			hIThvPiOWy, az1oIQ9_Q6N =  /*line :1*/ rune(az1oIQ9_Q6N[0]), az1oIQ9_Q6N[1:]
		} else {
			unWb50A, aub89wE2 :=  /*line :1*/ utf8.HVAV7aTqxzg(az1oIQ9_Q6N)
			hIThvPiOWy, az1oIQ9_Q6N = unWb50A, az1oIQ9_Q6N[aub89wE2:]
		}

		if hIThvPiOWy == cBR9L5k {
			continue
		}

		if hIThvPiOWy < cBR9L5k {
			hIThvPiOWy, cBR9L5k = cBR9L5k, hIThvPiOWy
		}

		if hIThvPiOWy < utf8.RuneSelf {

			if 'A' <= cBR9L5k && cBR9L5k <= 'Z' && hIThvPiOWy == cBR9L5k+'a'-'A' {
				continue
			}
			return false
		}

		unWb50A :=  /*line :1*/ unicode.NGnxX0NFZA4b(cBR9L5k)
		for unWb50A != cBR9L5k && unWb50A < hIThvPiOWy {
			unWb50A =  /*line :1*/ unicode.NGnxX0NFZA4b(unWb50A)
		}
		if unWb50A == hIThvPiOWy {
			continue
		}
		return false
	}

	return  /*line :1*/ len(az1oIQ9_Q6N) == 0
}


func FCsEk_zxUIUY(w4GNKhKtw, k2061Z3 string) int {
	dUrxzl :=  /*line :1*/ len(k2061Z3)
	switch {
	case dUrxzl == 0:
		return 0
	case dUrxzl == 1:
		return  /*line :1*/ Po8D3DCRZq(w4GNKhKtw, k2061Z3[0])
	case dUrxzl ==  /*line :1*/ len(w4GNKhKtw):
		if k2061Z3 == w4GNKhKtw {
			return 0
		}
		return -1
	case dUrxzl >  /*line :1*/ len(w4GNKhKtw):
		return -1
	case dUrxzl <= bytealg.MaxLen:

		if  /*line :1*/ len(w4GNKhKtw) <= bytealg.MaxBruteForce {
			return  /*line :1*/ bytealg.IndexString(w4GNKhKtw, k2061Z3)
		}
		xTtypjHh := k2061Z3[0]
		tyS7UZE3BJ := k2061Z3[1]
		gUiH7N := 0
		az1oIQ9_Q6N :=  /*line :1*/ len(w4GNKhKtw) - dUrxzl + 1
		tn0fyO0 := 0
		for gUiH7N < az1oIQ9_Q6N {
			if w4GNKhKtw[gUiH7N] != xTtypjHh {

				aan9Pjdhtu :=  /*line :1*/ Po8D3DCRZq(w4GNKhKtw[gUiH7N+1:az1oIQ9_Q6N], xTtypjHh)
				if aan9Pjdhtu < 0 {
					return -1
				}
				gUiH7N += aan9Pjdhtu + 1
			}
			if w4GNKhKtw[gUiH7N+1] == tyS7UZE3BJ && w4GNKhKtw[gUiH7N:gUiH7N+dUrxzl] == k2061Z3 {
				return gUiH7N
			}
			tn0fyO0++
			gUiH7N++

			if tn0fyO0 >  /*line :1*/ bytealg.Cutover(gUiH7N) {
				unWb50A :=  /*line :1*/ bytealg.IndexString(w4GNKhKtw[gUiH7N:], k2061Z3)
				if unWb50A >= 0 {
					return unWb50A + gUiH7N
				}
				return -1
			}
		}
		return -1
	}
	xTtypjHh := k2061Z3[0]
	tyS7UZE3BJ := k2061Z3[1]
	gUiH7N := 0
	az1oIQ9_Q6N :=  /*line :1*/ len(w4GNKhKtw) - dUrxzl + 1
	tn0fyO0 := 0
	for gUiH7N < az1oIQ9_Q6N {
		if w4GNKhKtw[gUiH7N] != xTtypjHh {
			aan9Pjdhtu :=  /*line :1*/ Po8D3DCRZq(w4GNKhKtw[gUiH7N+1:az1oIQ9_Q6N], xTtypjHh)
			if aan9Pjdhtu < 0 {
				return -1
			}
			gUiH7N += aan9Pjdhtu + 1
		}
		if w4GNKhKtw[gUiH7N+1] == tyS7UZE3BJ && w4GNKhKtw[gUiH7N:gUiH7N+dUrxzl] == k2061Z3 {
			return gUiH7N
		}
		gUiH7N++
		tn0fyO0++
		if tn0fyO0 >= 4+gUiH7N>>4 && gUiH7N < az1oIQ9_Q6N {

			xXcc3ZyUB3kD :=  /*line :1*/ bytealg.IndexRabinKarp(w4GNKhKtw[gUiH7N:], k2061Z3)
			if xXcc3ZyUB3kD < 0 {
				return -1
			}
			return gUiH7N + xXcc3ZyUB3kD
		}
	}
	return -1
}





func E093HWXxK(w4GNKhKtw, mDyLGdn7 string) (fBp4ZO, aIA7L7gaXkq string, eCIbzGMFM2p bool) {
	if gUiH7N :=  /*line :1*/ FCsEk_zxUIUY(w4GNKhKtw, mDyLGdn7); gUiH7N >= 0 {
		return w4GNKhKtw[:gUiH7N], w4GNKhKtw[gUiH7N+ /*line :1*/ len(mDyLGdn7):], true
	}
	return w4GNKhKtw, "", false
}





func ZOvqjnAMcI(w4GNKhKtw, aC71QFDwvE0T string) (aIA7L7gaXkq string, eCIbzGMFM2p bool) {
	if ! /*line :1*/ E8ZCzfgw(w4GNKhKtw, aC71QFDwvE0T) {
		return w4GNKhKtw, false
	}
	return w4GNKhKtw[ /*line :1*/ len(aC71QFDwvE0T):], true
}





func JrAaJF(w4GNKhKtw, triI9zMn4AM string) (fBp4ZO string, eCIbzGMFM2p bool) {
	if ! /*line :1*/ RZpWO5X7(w4GNKhKtw, triI9zMn4AM) {
		return w4GNKhKtw, false
	}
	return w4GNKhKtw[: /*line :1*/ len(w4GNKhKtw)- /*line :1*/ len(triI9zMn4AM)], true
}
