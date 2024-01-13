//line :1


package wLlfRPv

import (
	"internal/bytealg"
	unicode "vGJV8maRK"
	utf8 "CuAc0pPZwf"
)




func H2kZ75hy(zF4ny_7, ka9IyYc0 []byte) bool {

	return  /*line :1*/ string(zF4ny_7) ==  /*line :1*/ string(ka9IyYc0)
}




func SzbNbaS_(zF4ny_7, ka9IyYc0 []byte) int {
	return  /*line :1*/ bytealg.Compare(zF4ny_7, ka9IyYc0)
}



func bf1afp_V(nhst_x []byte, lQivwVxjomk int) [][]byte {
	if lQivwVxjomk <= 0 || lQivwVxjomk >  /*line :1*/ len(nhst_x) {
		lQivwVxjomk =  /*line :1*/ len(nhst_x)
	}
	zF4ny_7 :=  /*line :1*/ make([][]byte, lQivwVxjomk)
	var kXik_zS1N6 int
	fcOR1KeZ8 := 0
	for  /*line :1*/ len(nhst_x) > 0 {
		if fcOR1KeZ8+1 >= lQivwVxjomk {
			zF4ny_7[fcOR1KeZ8] = nhst_x
			fcOR1KeZ8++
			break
		}
		_, kXik_zS1N6 =  /*line :1*/ utf8.EicfpCPn(nhst_x)
		zF4ny_7[fcOR1KeZ8] = nhst_x[0:kXik_zS1N6:kXik_zS1N6]
		nhst_x = nhst_x[kXik_zS1N6:]
		fcOR1KeZ8++
	}
	return zF4ny_7[0:fcOR1KeZ8]
}



func B_RN62qkDAX(nhst_x, b6Z2P0 []byte) int {

	if  /*line :1*/ len(b6Z2P0) == 0 {
		return  /*line :1*/ utf8.SGiJAHBAT(nhst_x) + 1
	}
	if  /*line :1*/ len(b6Z2P0) == 1 {
		return  /*line :1*/ bytealg.Count(nhst_x, b6Z2P0[0])
	}
	lQivwVxjomk := 0
	for {
		musLbUdLeb :=  /*line :1*/ DxTUlq(nhst_x, b6Z2P0)
		if musLbUdLeb == -1 {
			return lQivwVxjomk
		}
		lQivwVxjomk++
		nhst_x = nhst_x[musLbUdLeb+ /*line :1*/ len(b6Z2P0):]
	}
}


func ObDgeyX5qZ(ka9IyYc0, w_FLshu7 []byte) bool {
	return  /*line :1*/ DxTUlq(ka9IyYc0, w_FLshu7) != -1
}


func RthIO0aHKd(ka9IyYc0 []byte, tjaOlg6yCy string) bool {
	return  /*line :1*/ C9gUkN(ka9IyYc0, tjaOlg6yCy) >= 0
}


func PeLDBV2vxKv(ka9IyYc0 []byte, dIwXIFT rune) bool {
	return  /*line :1*/ K3njXz(ka9IyYc0, dIwXIFT) >= 0
}


func RcUPncmQSI(ka9IyYc0 []byte, m1lAquxM5 func(rune) bool) bool {
	return  /*line :1*/ J9jQPAxN(ka9IyYc0, m1lAquxM5) >= 0
}


func AXZRWQNqa(ka9IyYc0 []byte, wOPKkW byte) int {
	return  /*line :1*/ bytealg.IndexByte(ka9IyYc0, wOPKkW)
}

func hMhMD5W_(nhst_x []byte, wOPKkW byte) int {
	for musLbUdLeb, ka9IyYc0 := range nhst_x {
		if ka9IyYc0 == wOPKkW {
			return musLbUdLeb
		}
	}
	return -1
}


func EWyrU9hZu(nhst_x, b6Z2P0 []byte) int {
	lQivwVxjomk :=  /*line :1*/ len(b6Z2P0)
	switch {
	case lQivwVxjomk == 0:
		return  /*line :1*/ len(nhst_x)
	case lQivwVxjomk == 1:
		return  /*line :1*/ QpBCx3lccz(nhst_x, b6Z2P0[0])
	case lQivwVxjomk ==  /*line :1*/ len(nhst_x):
		if  /*line :1*/ H2kZ75hy(nhst_x, b6Z2P0) {
			return 0
		}
		return -1
	case lQivwVxjomk >  /*line :1*/ len(nhst_x):
		return -1
	}

	aC98Kb1mJZB, xFPCrYGac :=  /*line :1*/ bytealg.HashStrRevBytes(b6Z2P0)
	yMg75blweed :=  /*line :1*/ len(nhst_x) - lQivwVxjomk
	var vkl8ihMYOnJ7 uint32
	for musLbUdLeb :=  /*line :1*/ len(nhst_x) - 1; musLbUdLeb >= yMg75blweed; musLbUdLeb-- {
		vkl8ihMYOnJ7 = vkl8ihMYOnJ7*bytealg.PrimeRK +  /*line :1*/ uint32(nhst_x[musLbUdLeb])
	}
	if vkl8ihMYOnJ7 == aC98Kb1mJZB &&  /*line :1*/ H2kZ75hy(nhst_x[yMg75blweed:], b6Z2P0) {
		return yMg75blweed
	}
	for musLbUdLeb := yMg75blweed - 1; musLbUdLeb >= 0; musLbUdLeb-- {
		vkl8ihMYOnJ7 *= bytealg.PrimeRK
		vkl8ihMYOnJ7 +=  /*line :1*/ uint32(nhst_x[musLbUdLeb])
		vkl8ihMYOnJ7 -= xFPCrYGac *  /*line :1*/ uint32(nhst_x[musLbUdLeb+lQivwVxjomk])
		if vkl8ihMYOnJ7 == aC98Kb1mJZB &&  /*line :1*/ H2kZ75hy(nhst_x[musLbUdLeb:musLbUdLeb+lQivwVxjomk], b6Z2P0) {
			return musLbUdLeb
		}
	}
	return -1
}


func QpBCx3lccz(nhst_x []byte, wOPKkW byte) int {
	for musLbUdLeb :=  /*line :1*/ len(nhst_x) - 1; musLbUdLeb >= 0; musLbUdLeb-- {
		if nhst_x[musLbUdLeb] == wOPKkW {
			return musLbUdLeb
		}
	}
	return -1
}






func K3njXz(nhst_x []byte, dIwXIFT rune) int {
	switch {
	case 0 <= dIwXIFT && dIwXIFT < utf8.RuneSelf:
		return  /*line :1*/ AXZRWQNqa(nhst_x,  /*line :1*/ byte(dIwXIFT))
	case dIwXIFT == utf8.RuneError:
		for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); {
			iEnyWtLg0l, lQivwVxjomk :=  /*line :1*/ utf8.EicfpCPn(nhst_x[musLbUdLeb:])
			if iEnyWtLg0l == utf8.RuneError {
				return musLbUdLeb
			}
			musLbUdLeb += lQivwVxjomk
		}
		return -1
	case ! /*line :1*/ utf8.DYq33yNk(dIwXIFT):
		return -1
	default:
		var ka9IyYc0 [utf8.UTFMax]byte
		lQivwVxjomk :=  /*line :1*/ utf8.YoEca6(ka9IyYc0[:], dIwXIFT)
		return  /*line :1*/ DxTUlq(nhst_x, ka9IyYc0[:lQivwVxjomk])
	}
}





func C9gUkN(nhst_x []byte, tjaOlg6yCy string) int {
	if tjaOlg6yCy == "" {

		return -1
	}
	if  /*line :1*/ len(nhst_x) == 1 {
		dIwXIFT :=  /*line :1*/ rune(nhst_x[0])
		if dIwXIFT >= utf8.RuneSelf {

			for _, dIwXIFT = range tjaOlg6yCy {
				if dIwXIFT == utf8.RuneError {
					return 0
				}
			}
			return -1
		}
		if  /*line :1*/ bytealg.IndexByteString(tjaOlg6yCy, nhst_x[0]) >= 0 {
			return 0
		}
		return -1
	}
	if  /*line :1*/ len(tjaOlg6yCy) == 1 {
		dIwXIFT :=  /*line :1*/ rune(tjaOlg6yCy[0])
		if dIwXIFT >= utf8.RuneSelf {
			dIwXIFT = utf8.RuneError
		}
		return  /*line :1*/ K3njXz(nhst_x, dIwXIFT)
	}
	if  /*line :1*/ len(nhst_x) > 8 {
		if gJ_oT8vTNkQ, vlf9xhhY :=  /*line :1*/ vjanqQj(tjaOlg6yCy); vlf9xhhY {
			for musLbUdLeb, wOPKkW := range nhst_x {
				if  /*line :1*/ gJ_oT8vTNkQ.im2o07u9eL(wOPKkW) {
					return musLbUdLeb
				}
			}
			return -1
		}
	}
	var jJGIFmu7 int
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); musLbUdLeb += jJGIFmu7 {
		dIwXIFT :=  /*line :1*/ rune(nhst_x[musLbUdLeb])
		if dIwXIFT < utf8.RuneSelf {
			if  /*line :1*/ bytealg.IndexByteString(tjaOlg6yCy, nhst_x[musLbUdLeb]) >= 0 {
				return musLbUdLeb
			}
			jJGIFmu7 = 1
			continue
		}
		dIwXIFT, jJGIFmu7 =  /*line :1*/ utf8.EicfpCPn(nhst_x[musLbUdLeb:])
		if dIwXIFT != utf8.RuneError {

			if  /*line :1*/ len(tjaOlg6yCy) == jJGIFmu7 {
				if tjaOlg6yCy ==  /*line :1*/ string(dIwXIFT) {
					return musLbUdLeb
				}
				continue
			}

			if bytealg.MaxLen >= jJGIFmu7 {
				if  /*line :1*/ bytealg.IndexString(tjaOlg6yCy,  /*line :1*/ string(dIwXIFT)) >= 0 {
					return musLbUdLeb
				}
				continue
			}
		}
		for _, bKhTMqSi := range tjaOlg6yCy {
			if dIwXIFT == bKhTMqSi {
				return musLbUdLeb
			}
		}
	}
	return -1
}





func S2bothdlq1Zr(nhst_x []byte, tjaOlg6yCy string) int {
	if tjaOlg6yCy == "" {

		return -1
	}
	if  /*line :1*/ len(nhst_x) > 8 {
		if gJ_oT8vTNkQ, vlf9xhhY :=  /*line :1*/ vjanqQj(tjaOlg6yCy); vlf9xhhY {
			for musLbUdLeb :=  /*line :1*/ len(nhst_x) - 1; musLbUdLeb >= 0; musLbUdLeb-- {
				if  /*line :1*/ gJ_oT8vTNkQ.im2o07u9eL(nhst_x[musLbUdLeb]) {
					return musLbUdLeb
				}
			}
			return -1
		}
	}
	if  /*line :1*/ len(nhst_x) == 1 {
		dIwXIFT :=  /*line :1*/ rune(nhst_x[0])
		if dIwXIFT >= utf8.RuneSelf {
			for _, dIwXIFT = range tjaOlg6yCy {
				if dIwXIFT == utf8.RuneError {
					return 0
				}
			}
			return -1
		}
		if  /*line :1*/ bytealg.IndexByteString(tjaOlg6yCy, nhst_x[0]) >= 0 {
			return 0
		}
		return -1
	}
	if  /*line :1*/ len(tjaOlg6yCy) == 1 {
		cBc_cVf3 :=  /*line :1*/ rune(tjaOlg6yCy[0])
		if cBc_cVf3 >= utf8.RuneSelf {
			cBc_cVf3 = utf8.RuneError
		}
		for musLbUdLeb :=  /*line :1*/ len(nhst_x); musLbUdLeb > 0; {
			dIwXIFT, kXik_zS1N6 :=  /*line :1*/ utf8.QG3o80(nhst_x[:musLbUdLeb])
			musLbUdLeb -= kXik_zS1N6
			if dIwXIFT == cBc_cVf3 {
				return musLbUdLeb
			}
		}
		return -1
	}
	for musLbUdLeb :=  /*line :1*/ len(nhst_x); musLbUdLeb > 0; {
		dIwXIFT :=  /*line :1*/ rune(nhst_x[musLbUdLeb-1])
		if dIwXIFT < utf8.RuneSelf {
			if  /*line :1*/ bytealg.IndexByteString(tjaOlg6yCy, nhst_x[musLbUdLeb-1]) >= 0 {
				return musLbUdLeb - 1
			}
			musLbUdLeb--
			continue
		}
		dIwXIFT, kXik_zS1N6 :=  /*line :1*/ utf8.QG3o80(nhst_x[:musLbUdLeb])
		musLbUdLeb -= kXik_zS1N6
		if dIwXIFT != utf8.RuneError {

			if  /*line :1*/ len(tjaOlg6yCy) == kXik_zS1N6 {
				if tjaOlg6yCy ==  /*line :1*/ string(dIwXIFT) {
					return musLbUdLeb
				}
				continue
			}

			if bytealg.MaxLen >= kXik_zS1N6 {
				if  /*line :1*/ bytealg.IndexString(tjaOlg6yCy,  /*line :1*/ string(dIwXIFT)) >= 0 {
					return musLbUdLeb
				}
				continue
			}
		}
		for _, bKhTMqSi := range tjaOlg6yCy {
			if dIwXIFT == bKhTMqSi {
				return musLbUdLeb
			}
		}
	}
	return -1
}



func bsW343(nhst_x, b6Z2P0 []byte, xbIvDYg, lQivwVxjomk int) [][]byte {
	if lQivwVxjomk == 0 {
		return nil
	}
	if  /*line :1*/ len(b6Z2P0) == 0 {
		return  /*line :1*/ bf1afp_V(nhst_x, lQivwVxjomk)
	}
	if lQivwVxjomk < 0 {
		lQivwVxjomk =  /*line :1*/ B_RN62qkDAX(nhst_x, b6Z2P0) + 1
	}
	if lQivwVxjomk >  /*line :1*/ len(nhst_x)+1 {
		lQivwVxjomk =  /*line :1*/ len(nhst_x) + 1
	}

	zF4ny_7 :=  /*line :1*/ make([][]byte, lQivwVxjomk)
	lQivwVxjomk--
	musLbUdLeb := 0
	for musLbUdLeb < lQivwVxjomk {
		aJCPDy4w :=  /*line :1*/ DxTUlq(nhst_x, b6Z2P0)
		if aJCPDy4w < 0 {
			break
		}
		zF4ny_7[musLbUdLeb] = nhst_x[: aJCPDy4w+xbIvDYg : aJCPDy4w+xbIvDYg]
		nhst_x = nhst_x[aJCPDy4w+ /*line :1*/ len(b6Z2P0):]
		musLbUdLeb++
	}
	zF4ny_7[musLbUdLeb] = nhst_x
	return zF4ny_7[:musLbUdLeb+1]
}











func ADsXOYUCO9(nhst_x, b6Z2P0 []byte, lQivwVxjomk int) [][]byte {
	return  /*line :1*/ bsW343(nhst_x, b6Z2P0, 0, lQivwVxjomk)
}









func Ix6UMXrTTK(nhst_x, b6Z2P0 []byte, lQivwVxjomk int) [][]byte {
	return  /*line :1*/ bsW343(nhst_x, b6Z2P0,  /*line :1*/ len(b6Z2P0), lQivwVxjomk)
}







func Wp5fb1(nhst_x, b6Z2P0 []byte) [][]byte	{ return  /*line :1*/ bsW343(nhst_x, b6Z2P0, 0, -1) }





func CaaXPMtcg(nhst_x, b6Z2P0 []byte) [][]byte {
	return  /*line :1*/ bsW343(nhst_x, b6Z2P0,  /*line :1*/ len(b6Z2P0), -1)
}

var dU3F3UrRwC = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}





func ZnknwW(nhst_x []byte) [][]byte {

	lQivwVxjomk := 0
	bs8mAadn2sQ := 1

	sRXWJN8mPxsI :=  /*line :1*/ uint8(0)
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); musLbUdLeb++ {
		dIwXIFT := nhst_x[musLbUdLeb]
		sRXWJN8mPxsI |= dIwXIFT
		jStGAo :=  /*line :1*/ int(dU3F3UrRwC[dIwXIFT])
		lQivwVxjomk += bs8mAadn2sQ & ^jStGAo
		bs8mAadn2sQ = jStGAo
	}

	if sRXWJN8mPxsI >= utf8.RuneSelf {

		return  /*line :1*/ OfK2EEbx_K(nhst_x, unicode.Mzbe7SW3gm0k)
	}

	zF4ny_7 :=  /*line :1*/ make([][]byte, lQivwVxjomk)
	fcOR1KeZ8 := 0
	mPekCTMh := 0
	musLbUdLeb := 0

	for musLbUdLeb <  /*line :1*/ len(nhst_x) && dU3F3UrRwC[nhst_x[musLbUdLeb]] != 0 {
		musLbUdLeb++
	}
	mPekCTMh = musLbUdLeb
	for musLbUdLeb <  /*line :1*/ len(nhst_x) {
		if dU3F3UrRwC[nhst_x[musLbUdLeb]] == 0 {
			musLbUdLeb++
			continue
		}
		zF4ny_7[fcOR1KeZ8] = nhst_x[mPekCTMh:musLbUdLeb:musLbUdLeb]
		fcOR1KeZ8++
		musLbUdLeb++

		for musLbUdLeb <  /*line :1*/ len(nhst_x) && dU3F3UrRwC[nhst_x[musLbUdLeb]] != 0 {
			musLbUdLeb++
		}
		mPekCTMh = musLbUdLeb
	}
	if mPekCTMh <  /*line :1*/ len(nhst_x) {
		zF4ny_7[fcOR1KeZ8] = nhst_x[mPekCTMh: /*line :1*/ len(nhst_x): /*line :1*/ len(nhst_x)]
	}
	return zF4ny_7
}








func OfK2EEbx_K(nhst_x []byte, m1lAquxM5 func(rune) bool) [][]byte {
	
	
	type d99hxi2 struct {
		t58suHObXgU9	int
		cvWVr_i6gkp	int
	}
	wkO5kaKx :=  /*line :1*/ make([]d99hxi2, 0, 32)

	z_HY7AHGxq6 := -1
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); {
		kXik_zS1N6 := 1
		dIwXIFT :=  /*line :1*/ rune(nhst_x[musLbUdLeb])
		if dIwXIFT >= utf8.RuneSelf {
			dIwXIFT, kXik_zS1N6 =  /*line :1*/ utf8.EicfpCPn(nhst_x[musLbUdLeb:])
		}
		if  /*line :1*/ m1lAquxM5(dIwXIFT) {
			if z_HY7AHGxq6 >= 0 {
				wkO5kaKx =  /*line :1*/ append(wkO5kaKx, d99hxi2{z_HY7AHGxq6, musLbUdLeb})
				z_HY7AHGxq6 = -1
			}
		} else {
			if z_HY7AHGxq6 < 0 {
				z_HY7AHGxq6 = musLbUdLeb
			}
		}
		musLbUdLeb += kXik_zS1N6
	}

	if z_HY7AHGxq6 >= 0 {
		wkO5kaKx =  /*line :1*/ append(wkO5kaKx, d99hxi2{z_HY7AHGxq6,  /*line :1*/ len(nhst_x)})
	}

	zF4ny_7 :=  /*line :1*/ make([][]byte,  /*line :1*/ len(wkO5kaKx))
	for musLbUdLeb, d99hxi2 := range wkO5kaKx {
		zF4ny_7[musLbUdLeb] = nhst_x[d99hxi2.t58suHObXgU9:d99hxi2.cvWVr_i6gkp:d99hxi2.cvWVr_i6gkp]
	}

	return zF4ny_7
}



func CEmVHJFPr(nhst_x [][]byte, b6Z2P0 []byte) []byte {
	if  /*line :1*/ len(nhst_x) == 0 {
		return []byte{}
	}
	if  /*line :1*/ len(nhst_x) == 1 {

		return  /*line :1*/ append([] /*line :1*/ byte(nil), nhst_x[0]...)
	}

	var lQivwVxjomk int
	if  /*line :1*/ len(b6Z2P0) > 0 {
		if  /*line :1*/ len(b6Z2P0) >= maxInt/( /*line :1*/ len(nhst_x)-1) {
			 /*line :1*/ panic("bytes: Join output length overflow")
		}
		lQivwVxjomk +=  /*line :1*/ len(b6Z2P0) * ( /*line :1*/ len(nhst_x) - 1)
	}
	for _, bzu4YuWrG := range nhst_x {
		if  /*line :1*/ len(bzu4YuWrG) > maxInt-lQivwVxjomk {
			 /*line :1*/ panic("bytes: Join output length overflow")
		}
		lQivwVxjomk +=  /*line :1*/ len(bzu4YuWrG)
	}

	ka9IyYc0 :=  /*line :1*/ bytealg.MakeNoZero(lQivwVxjomk)
	j8d8cgg_6 :=  /*line :1*/ copy(ka9IyYc0, nhst_x[0])
	for _, bzu4YuWrG := range nhst_x[1:] {
		j8d8cgg_6 +=  /*line :1*/ copy(ka9IyYc0[j8d8cgg_6:], b6Z2P0)
		j8d8cgg_6 +=  /*line :1*/ copy(ka9IyYc0[j8d8cgg_6:], bzu4YuWrG)
	}
	return ka9IyYc0
}


func V2Q3qKsmU(nhst_x, sMXnw0t6SSCa []byte) bool {
	return  /*line :1*/ len(nhst_x) >=  /*line :1*/ len(sMXnw0t6SSCa) &&  /*line :1*/ H2kZ75hy(nhst_x[0: /*line :1*/ len(sMXnw0t6SSCa)], sMXnw0t6SSCa)
}


func Ar8BM7P(nhst_x, dcbttECl46 []byte) bool {
	return  /*line :1*/ len(nhst_x) >=  /*line :1*/ len(dcbttECl46) &&  /*line :1*/ H2kZ75hy(nhst_x[ /*line :1*/ len(nhst_x)- /*line :1*/ len(dcbttECl46):], dcbttECl46)
}





func RDzqwqK(xvX4ia func(dIwXIFT rune) rune, nhst_x []byte) []byte {

	ka9IyYc0 :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(nhst_x))
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); {
		idsjak6 := 1
		dIwXIFT :=  /*line :1*/ rune(nhst_x[musLbUdLeb])
		if dIwXIFT >= utf8.RuneSelf {
			dIwXIFT, idsjak6 =  /*line :1*/ utf8.EicfpCPn(nhst_x[musLbUdLeb:])
		}
		dIwXIFT =  /*line :1*/ xvX4ia(dIwXIFT)
		if dIwXIFT >= 0 {
			ka9IyYc0 =  /*line :1*/ utf8.Ht7oMzd(ka9IyYc0, dIwXIFT)
		}
		musLbUdLeb += idsjak6
	}
	return ka9IyYc0
}





func BW2oAsgc15(ka9IyYc0 []byte, edGlaqVMK int) []byte {
	if edGlaqVMK == 0 {
		return []byte{}
	}

	if edGlaqVMK < 0 {
		 /*line :1*/ panic("bytes: negative Repeat count")
	}
	if  /*line :1*/ len(ka9IyYc0) >= maxInt/edGlaqVMK {
		 /*line :1*/ panic("bytes: Repeat output length overflow")
	}
	lQivwVxjomk :=  /*line :1*/ len(ka9IyYc0) * edGlaqVMK

	if  /*line :1*/ len(ka9IyYc0) == 0 {
		return []byte{}
	}

	
	
	
	
	
	
	
	
	
	
	const chunkLimit = 8 * 1024
	tCTSL3 := lQivwVxjomk
	if tCTSL3 > chunkLimit {
		tCTSL3 = chunkLimit /  /*line :1*/ len(ka9IyYc0) *  /*line :1*/ len(ka9IyYc0)
		if tCTSL3 == 0 {
			tCTSL3 =  /*line :1*/ len(ka9IyYc0)
		}
	}
	y6zZK9TTR :=  /*line :1*/ bytealg.MakeNoZero(lQivwVxjomk)
	j8d8cgg_6 :=  /*line :1*/ copy(y6zZK9TTR, ka9IyYc0)
	for j8d8cgg_6 < lQivwVxjomk {
		vB9gckHu := j8d8cgg_6
		if vB9gckHu > tCTSL3 {
			vB9gckHu = tCTSL3
		}
		j8d8cgg_6 +=  /*line :1*/ copy(y6zZK9TTR[j8d8cgg_6:], y6zZK9TTR[:vB9gckHu])
	}
	return y6zZK9TTR
}



func XY6kX4k(nhst_x []byte) []byte {
	vlf9xhhY, vk3KiKm := true, false
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); musLbUdLeb++ {
		wOPKkW := nhst_x[musLbUdLeb]
		if wOPKkW >= utf8.RuneSelf {
			vlf9xhhY = false
			break
		}
		vk3KiKm = vk3KiKm || ('a' <= wOPKkW && wOPKkW <= 'z')
	}

	if vlf9xhhY {
		if !vk3KiKm {

			return  /*line :1*/ append([] /*line :1*/ byte(""), nhst_x...)
		}
		ka9IyYc0 :=  /*line :1*/ bytealg.MakeNoZero( /*line :1*/ len(nhst_x))
		for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); musLbUdLeb++ {
			wOPKkW := nhst_x[musLbUdLeb]
			if 'a' <= wOPKkW && wOPKkW <= 'z' {
				wOPKkW -= 'a' - 'A'
			}
			ka9IyYc0[musLbUdLeb] = wOPKkW
		}
		return ka9IyYc0
	}
	return  /*line :1*/ RDzqwqK(unicode.HJD8X_JbK, nhst_x)
}



func GSTXMCfBXsV(nhst_x []byte) []byte {
	vlf9xhhY, n7zVpg := true, false
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); musLbUdLeb++ {
		wOPKkW := nhst_x[musLbUdLeb]
		if wOPKkW >= utf8.RuneSelf {
			vlf9xhhY = false
			break
		}
		n7zVpg = n7zVpg || ('A' <= wOPKkW && wOPKkW <= 'Z')
	}

	if vlf9xhhY {
		if !n7zVpg {
			return  /*line :1*/ append([] /*line :1*/ byte(""), nhst_x...)
		}
		ka9IyYc0 :=  /*line :1*/ bytealg.MakeNoZero( /*line :1*/ len(nhst_x))
		for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); musLbUdLeb++ {
			wOPKkW := nhst_x[musLbUdLeb]
			if 'A' <= wOPKkW && wOPKkW <= 'Z' {
				wOPKkW += 'a' - 'A'
			}
			ka9IyYc0[musLbUdLeb] = wOPKkW
		}
		return ka9IyYc0
	}
	return  /*line :1*/ RDzqwqK(unicode.OLZ8W2, nhst_x)
}


func HaZtBWM8_gs(nhst_x []byte) []byte	{ return  /*line :1*/ RDzqwqK(unicode.QgPuNi8kYiC, nhst_x) }



func Vm46gGKS8C(wOPKkW unicode.KmthvGp3gyLu, nhst_x []byte) []byte {
	return  /*line :1*/ RDzqwqK(wOPKkW.ToUpper, nhst_x)
}



func FQwhZBwy(wOPKkW unicode.KmthvGp3gyLu, nhst_x []byte) []byte {
	return  /*line :1*/ RDzqwqK(wOPKkW.ToLower, nhst_x)
}



func QNC8vMhrv(wOPKkW unicode.KmthvGp3gyLu, nhst_x []byte) []byte {
	return  /*line :1*/ RDzqwqK(wOPKkW.ToTitle, nhst_x)
}



func TwHOHlB9R(nhst_x, stGseuanZ []byte) []byte {
	ka9IyYc0 :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(nhst_x)+ /*line :1*/ len(stGseuanZ))
	aweDSMzW605 := false
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(nhst_x); {
		wOPKkW := nhst_x[musLbUdLeb]
		if wOPKkW < utf8.RuneSelf {
			musLbUdLeb++
			aweDSMzW605 = false
			ka9IyYc0 =  /*line :1*/ append(ka9IyYc0, wOPKkW)
			continue
		}
		_, idsjak6 :=  /*line :1*/ utf8.EicfpCPn(nhst_x[musLbUdLeb:])
		if idsjak6 == 1 {
			musLbUdLeb++
			if !aweDSMzW605 {
				aweDSMzW605 = true
				ka9IyYc0 =  /*line :1*/ append(ka9IyYc0, stGseuanZ...)
			}
			continue
		}
		aweDSMzW605 = false
		ka9IyYc0 =  /*line :1*/ append(ka9IyYc0, nhst_x[musLbUdLeb:musLbUdLeb+idsjak6]...)
		musLbUdLeb += idsjak6
	}
	return ka9IyYc0
}



func wNTFc6U2Rcck(dIwXIFT rune) bool {

	if dIwXIFT <= 0x7F {
		switch {
		case '0' <= dIwXIFT && dIwXIFT <= '9':
			return false
		case 'a' <= dIwXIFT && dIwXIFT <= 'z':
			return false
		case 'A' <= dIwXIFT && dIwXIFT <= 'Z':
			return false
		case dIwXIFT == '_':
			return false
		}
		return true
	}

	if  /*line :1*/ unicode.Nyu9Lcgu(dIwXIFT) ||  /*line :1*/ unicode.FquwNEzWV(dIwXIFT) {
		return false
	}

	return  /*line :1*/ unicode.Mzbe7SW3gm0k(dIwXIFT)
}






func EIZS2KIGYDF(nhst_x []byte) []byte {

	yrxtnZyh := ' '
	return  /*line :1*/ RDzqwqK(
		func(dIwXIFT rune) rune {
			if  /*line :1*/ wNTFc6U2Rcck(yrxtnZyh) {
				yrxtnZyh = dIwXIFT
				return  /*line :1*/ unicode.QgPuNi8kYiC(dIwXIFT)
			}
			yrxtnZyh = dIwXIFT
			return dIwXIFT
		},
		nhst_x)
}



func BatzdA(nhst_x []byte, m1lAquxM5 func(dIwXIFT rune) bool) []byte {
	musLbUdLeb :=  /*line :1*/ bxKIwnQn9yy(nhst_x, m1lAquxM5, false)
	if musLbUdLeb == -1 {
		return nil
	}
	return nhst_x[musLbUdLeb:]
}



func BBF5eel6b8(nhst_x []byte, m1lAquxM5 func(dIwXIFT rune) bool) []byte {
	musLbUdLeb :=  /*line :1*/ kXOsRR73(nhst_x, m1lAquxM5, false)
	if musLbUdLeb >= 0 && nhst_x[musLbUdLeb] >= utf8.RuneSelf {
		_, idsjak6 :=  /*line :1*/ utf8.EicfpCPn(nhst_x[musLbUdLeb:])
		musLbUdLeb += idsjak6
	} else {
		musLbUdLeb++
	}
	return nhst_x[0:musLbUdLeb]
}



func JxNiNlj(nhst_x []byte, m1lAquxM5 func(dIwXIFT rune) bool) []byte {
	return  /*line :1*/ BBF5eel6b8( /*line :1*/ BatzdA(nhst_x, m1lAquxM5), m1lAquxM5)
}



func PDGDlKzE(nhst_x, sMXnw0t6SSCa []byte) []byte {
	if  /*line :1*/ V2Q3qKsmU(nhst_x, sMXnw0t6SSCa) {
		return nhst_x[ /*line :1*/ len(sMXnw0t6SSCa):]
	}
	return nhst_x
}



func W2aoqK(nhst_x, dcbttECl46 []byte) []byte {
	if  /*line :1*/ Ar8BM7P(nhst_x, dcbttECl46) {
		return nhst_x[: /*line :1*/ len(nhst_x)- /*line :1*/ len(dcbttECl46)]
	}
	return nhst_x
}




func J9jQPAxN(nhst_x []byte, m1lAquxM5 func(dIwXIFT rune) bool) int {
	return  /*line :1*/ bxKIwnQn9yy(nhst_x, m1lAquxM5, true)
}




func A3VhagB(nhst_x []byte, m1lAquxM5 func(dIwXIFT rune) bool) int {
	return  /*line :1*/ kXOsRR73(nhst_x, m1lAquxM5, true)
}




func bxKIwnQn9yy(nhst_x []byte, m1lAquxM5 func(dIwXIFT rune) bool, iotNa5nNIJN bool) int {
	z_HY7AHGxq6 := 0
	for z_HY7AHGxq6 <  /*line :1*/ len(nhst_x) {
		idsjak6 := 1
		dIwXIFT :=  /*line :1*/ rune(nhst_x[z_HY7AHGxq6])
		if dIwXIFT >= utf8.RuneSelf {
			dIwXIFT, idsjak6 =  /*line :1*/ utf8.EicfpCPn(nhst_x[z_HY7AHGxq6:])
		}
		if  /*line :1*/ m1lAquxM5(dIwXIFT) == iotNa5nNIJN {
			return z_HY7AHGxq6
		}
		z_HY7AHGxq6 += idsjak6
	}
	return -1
}




func kXOsRR73(nhst_x []byte, m1lAquxM5 func(dIwXIFT rune) bool, iotNa5nNIJN bool) int {
	for musLbUdLeb :=  /*line :1*/ len(nhst_x); musLbUdLeb > 0; {
		dIwXIFT, kXik_zS1N6 :=  /*line :1*/ rune(nhst_x[musLbUdLeb-1]), 1
		if dIwXIFT >= utf8.RuneSelf {
			dIwXIFT, kXik_zS1N6 =  /*line :1*/ utf8.QG3o80(nhst_x[0:musLbUdLeb])
		}
		musLbUdLeb -= kXik_zS1N6
		if  /*line :1*/ m1lAquxM5(dIwXIFT) == iotNa5nNIJN {
			return musLbUdLeb
		}
	}
	return -1
}









type xYXH9a [8]uint32



func vjanqQj(tjaOlg6yCy string) (gJ_oT8vTNkQ xYXH9a, lxjoXcF64YF bool) {
	for musLbUdLeb := 0; musLbUdLeb <  /*line :1*/ len(tjaOlg6yCy); musLbUdLeb++ {
		wOPKkW := tjaOlg6yCy[musLbUdLeb]
		if wOPKkW >= utf8.RuneSelf {
			return gJ_oT8vTNkQ, false
		}
		gJ_oT8vTNkQ[wOPKkW/32] |= 1 << (wOPKkW % 32)
	}
	return gJ_oT8vTNkQ, true
}


func (gJ_oT8vTNkQ *xYXH9a) im2o07u9eL(wOPKkW byte) bool {
	return (gJ_oT8vTNkQ[wOPKkW/32] & (1 << (wOPKkW % 32))) != 0
}




func gBptza(nhst_x string, dIwXIFT rune) bool {
	for _, wOPKkW := range nhst_x {
		if wOPKkW == dIwXIFT {
			return true
		}
	}
	return false
}



func RSqI0O3XG6C(nhst_x []byte, lQi3IN string) []byte {
	if  /*line :1*/ len(nhst_x) == 0 {

		return nil
	}
	if lQi3IN == "" {
		return nhst_x
	}
	if  /*line :1*/ len(lQi3IN) == 1 && lQi3IN[0] < utf8.RuneSelf {
		return  /*line :1*/ ooxkrva( /*line :1*/ e2kNZ1e(nhst_x, lQi3IN[0]), lQi3IN[0])
	}
	if gJ_oT8vTNkQ, lxjoXcF64YF :=  /*line :1*/ vjanqQj(lQi3IN); lxjoXcF64YF {
		return  /*line :1*/ gPp7Mr( /*line :1*/ fBaRMTyj(nhst_x, &gJ_oT8vTNkQ), &gJ_oT8vTNkQ)
	}
	return  /*line :1*/ tBGOEh5( /*line :1*/ f1fpKXLeUgYZ(nhst_x, lQi3IN), lQi3IN)
}



func JuFs7BQqB(nhst_x []byte, lQi3IN string) []byte {
	if  /*line :1*/ len(nhst_x) == 0 {

		return nil
	}
	if lQi3IN == "" {
		return nhst_x
	}
	if  /*line :1*/ len(lQi3IN) == 1 && lQi3IN[0] < utf8.RuneSelf {
		return  /*line :1*/ ooxkrva(nhst_x, lQi3IN[0])
	}
	if gJ_oT8vTNkQ, lxjoXcF64YF :=  /*line :1*/ vjanqQj(lQi3IN); lxjoXcF64YF {
		return  /*line :1*/ gPp7Mr(nhst_x, &gJ_oT8vTNkQ)
	}
	return  /*line :1*/ tBGOEh5(nhst_x, lQi3IN)
}

func ooxkrva(nhst_x []byte, wOPKkW byte) []byte {
	for  /*line :1*/ len(nhst_x) > 0 && nhst_x[0] == wOPKkW {
		nhst_x = nhst_x[1:]
	}
	if  /*line :1*/ len(nhst_x) == 0 {

		return nil
	}
	return nhst_x
}

func gPp7Mr(nhst_x []byte, gJ_oT8vTNkQ *xYXH9a) []byte {
	for  /*line :1*/ len(nhst_x) > 0 {
		if ! /*line :1*/ gJ_oT8vTNkQ.im2o07u9eL(nhst_x[0]) {
			break
		}
		nhst_x = nhst_x[1:]
	}
	if  /*line :1*/ len(nhst_x) == 0 {

		return nil
	}
	return nhst_x
}

func tBGOEh5(nhst_x []byte, lQi3IN string) []byte {
	for  /*line :1*/ len(nhst_x) > 0 {
		dIwXIFT, lQivwVxjomk :=  /*line :1*/ rune(nhst_x[0]), 1
		if dIwXIFT >= utf8.RuneSelf {
			dIwXIFT, lQivwVxjomk =  /*line :1*/ utf8.EicfpCPn(nhst_x)
		}
		if ! /*line :1*/ gBptza(lQi3IN, dIwXIFT) {
			break
		}
		nhst_x = nhst_x[lQivwVxjomk:]
	}
	if  /*line :1*/ len(nhst_x) == 0 {

		return nil
	}
	return nhst_x
}



func TBypK87Hkc(nhst_x []byte, lQi3IN string) []byte {
	if  /*line :1*/ len(nhst_x) == 0 || lQi3IN == "" {
		return nhst_x
	}
	if  /*line :1*/ len(lQi3IN) == 1 && lQi3IN[0] < utf8.RuneSelf {
		return  /*line :1*/ e2kNZ1e(nhst_x, lQi3IN[0])
	}
	if gJ_oT8vTNkQ, lxjoXcF64YF :=  /*line :1*/ vjanqQj(lQi3IN); lxjoXcF64YF {
		return  /*line :1*/ fBaRMTyj(nhst_x, &gJ_oT8vTNkQ)
	}
	return  /*line :1*/ f1fpKXLeUgYZ(nhst_x, lQi3IN)
}

func e2kNZ1e(nhst_x []byte, wOPKkW byte) []byte {
	for  /*line :1*/ len(nhst_x) > 0 && nhst_x[ /*line :1*/ len(nhst_x)-1] == wOPKkW {
		nhst_x = nhst_x[: /*line :1*/ len(nhst_x)-1]
	}
	return nhst_x
}

func fBaRMTyj(nhst_x []byte, gJ_oT8vTNkQ *xYXH9a) []byte {
	for  /*line :1*/ len(nhst_x) > 0 {
		if ! /*line :1*/ gJ_oT8vTNkQ.im2o07u9eL(nhst_x[ /*line :1*/ len(nhst_x)-1]) {
			break
		}
		nhst_x = nhst_x[: /*line :1*/ len(nhst_x)-1]
	}
	return nhst_x
}

func f1fpKXLeUgYZ(nhst_x []byte, lQi3IN string) []byte {
	for  /*line :1*/ len(nhst_x) > 0 {
		dIwXIFT, lQivwVxjomk :=  /*line :1*/ rune(nhst_x[ /*line :1*/ len(nhst_x)-1]), 1
		if dIwXIFT >= utf8.RuneSelf {
			dIwXIFT, lQivwVxjomk =  /*line :1*/ utf8.QG3o80(nhst_x)
		}
		if ! /*line :1*/ gBptza(lQi3IN, dIwXIFT) {
			break
		}
		nhst_x = nhst_x[: /*line :1*/ len(nhst_x)-lQivwVxjomk]
	}
	return nhst_x
}



func J4XEf7(nhst_x []byte) []byte {

	z_HY7AHGxq6 := 0
	for ; z_HY7AHGxq6 <  /*line :1*/ len(nhst_x); z_HY7AHGxq6++ {
		wOPKkW := nhst_x[z_HY7AHGxq6]
		if wOPKkW >= utf8.RuneSelf {

			return  /*line :1*/ JxNiNlj(nhst_x[z_HY7AHGxq6:], unicode.Mzbe7SW3gm0k)
		}
		if dU3F3UrRwC[wOPKkW] == 0 {
			break
		}
	}

	jZfaFnh :=  /*line :1*/ len(nhst_x)
	for ; jZfaFnh > z_HY7AHGxq6; jZfaFnh-- {
		wOPKkW := nhst_x[jZfaFnh-1]
		if wOPKkW >= utf8.RuneSelf {
			return  /*line :1*/ JxNiNlj(nhst_x[z_HY7AHGxq6:jZfaFnh], unicode.Mzbe7SW3gm0k)
		}
		if dU3F3UrRwC[wOPKkW] == 0 {
			break
		}
	}

	if z_HY7AHGxq6 == jZfaFnh {

		return nil
	}
	return nhst_x[z_HY7AHGxq6:jZfaFnh]
}



func BL0KFYkpINR(nhst_x []byte) []rune {
	qjUVLkM536S3 :=  /*line :1*/ make([]rune,  /*line :1*/ utf8.SGiJAHBAT(nhst_x))
	musLbUdLeb := 0
	for  /*line :1*/ len(nhst_x) > 0 {
		dIwXIFT, eGXSPnra9w :=  /*line :1*/ utf8.EicfpCPn(nhst_x)
		qjUVLkM536S3[musLbUdLeb] = dIwXIFT
		musLbUdLeb++
		nhst_x = nhst_x[eGXSPnra9w:]
	}
	return qjUVLkM536S3
}







func Zuy0tgi(nhst_x, iaSwmW, aCRiMxw8xmvH []byte, lQivwVxjomk int) []byte {
	aJCPDy4w := 0
	if lQivwVxjomk != 0 {

		aJCPDy4w =  /*line :1*/ B_RN62qkDAX(nhst_x, iaSwmW)
	}
	if aJCPDy4w == 0 {

		return  /*line :1*/ append([] /*line :1*/ byte(nil), nhst_x...)
	}
	if lQivwVxjomk < 0 || aJCPDy4w < lQivwVxjomk {
		lQivwVxjomk = aJCPDy4w
	}

	qjUVLkM536S3 :=  /*line :1*/ make([]byte,  /*line :1*/ len(nhst_x)+lQivwVxjomk*( /*line :1*/ len(aCRiMxw8xmvH)- /*line :1*/ len(iaSwmW)))
	yqP0sy := 0
	z_HY7AHGxq6 := 0
	for musLbUdLeb := 0; musLbUdLeb < lQivwVxjomk; musLbUdLeb++ {
		xDiLck8 := z_HY7AHGxq6
		if  /*line :1*/ len(iaSwmW) == 0 {
			if musLbUdLeb > 0 {
				_, idsjak6 :=  /*line :1*/ utf8.EicfpCPn(nhst_x[z_HY7AHGxq6:])
				xDiLck8 += idsjak6
			}
		} else {
			xDiLck8 +=  /*line :1*/ DxTUlq(nhst_x[z_HY7AHGxq6:], iaSwmW)
		}
		yqP0sy +=  /*line :1*/ copy(qjUVLkM536S3[yqP0sy:], nhst_x[z_HY7AHGxq6:xDiLck8])
		yqP0sy +=  /*line :1*/ copy(qjUVLkM536S3[yqP0sy:], aCRiMxw8xmvH)
		z_HY7AHGxq6 = xDiLck8 +  /*line :1*/ len(iaSwmW)
	}
	yqP0sy +=  /*line :1*/ copy(qjUVLkM536S3[yqP0sy:], nhst_x[z_HY7AHGxq6:])
	return qjUVLkM536S3[0:yqP0sy]
}






func Hg7Wjhf8NvB(nhst_x, iaSwmW, aCRiMxw8xmvH []byte) []byte {
	return  /*line :1*/ Zuy0tgi(nhst_x, iaSwmW, aCRiMxw8xmvH, -1)
}




func C6QO2xK4jumi(nhst_x, qjUVLkM536S3 []byte) bool {

	musLbUdLeb := 0
	for ; musLbUdLeb <  /*line :1*/ len(nhst_x) && musLbUdLeb <  /*line :1*/ len(qjUVLkM536S3); musLbUdLeb++ {
		e_9Sa0Od := nhst_x[musLbUdLeb]
		wvM7V9eZ := qjUVLkM536S3[musLbUdLeb]
		if e_9Sa0Od|wvM7V9eZ >= utf8.RuneSelf {
			goto hasUnicode
		}

		if wvM7V9eZ == e_9Sa0Od {
			continue
		}

		if wvM7V9eZ < e_9Sa0Od {
			wvM7V9eZ, e_9Sa0Od = e_9Sa0Od, wvM7V9eZ
		}

		if 'A' <= e_9Sa0Od && e_9Sa0Od <= 'Z' && wvM7V9eZ == e_9Sa0Od+'a'-'A' {
			continue
		}
		return false
	}

	return  /*line :1*/ len(nhst_x) ==  /*line :1*/ len(qjUVLkM536S3)

hasUnicode:
	nhst_x = nhst_x[musLbUdLeb:]
	qjUVLkM536S3 = qjUVLkM536S3[musLbUdLeb:]
	for  /*line :1*/ len(nhst_x) != 0 &&  /*line :1*/ len(qjUVLkM536S3) != 0 {
		
		var e_9Sa0Od, wvM7V9eZ rune
		if nhst_x[0] < utf8.RuneSelf {
			e_9Sa0Od, nhst_x =  /*line :1*/ rune(nhst_x[0]), nhst_x[1:]
		} else {
			dIwXIFT, kXik_zS1N6 :=  /*line :1*/ utf8.EicfpCPn(nhst_x)
			e_9Sa0Od, nhst_x = dIwXIFT, nhst_x[kXik_zS1N6:]
		}
		if qjUVLkM536S3[0] < utf8.RuneSelf {
			wvM7V9eZ, qjUVLkM536S3 =  /*line :1*/ rune(qjUVLkM536S3[0]), qjUVLkM536S3[1:]
		} else {
			dIwXIFT, kXik_zS1N6 :=  /*line :1*/ utf8.EicfpCPn(qjUVLkM536S3)
			wvM7V9eZ, qjUVLkM536S3 = dIwXIFT, qjUVLkM536S3[kXik_zS1N6:]
		}

		if wvM7V9eZ == e_9Sa0Od {
			continue
		}

		if wvM7V9eZ < e_9Sa0Od {
			wvM7V9eZ, e_9Sa0Od = e_9Sa0Od, wvM7V9eZ
		}

		if wvM7V9eZ < utf8.RuneSelf {

			if 'A' <= e_9Sa0Od && e_9Sa0Od <= 'Z' && wvM7V9eZ == e_9Sa0Od+'a'-'A' {
				continue
			}
			return false
		}

		dIwXIFT :=  /*line :1*/ unicode.NGnxX0NFZA4b(e_9Sa0Od)
		for dIwXIFT != e_9Sa0Od && dIwXIFT < wvM7V9eZ {
			dIwXIFT =  /*line :1*/ unicode.NGnxX0NFZA4b(dIwXIFT)
		}
		if dIwXIFT == wvM7V9eZ {
			continue
		}
		return false
	}

	return  /*line :1*/ len(nhst_x) ==  /*line :1*/ len(qjUVLkM536S3)
}


func DxTUlq(nhst_x, b6Z2P0 []byte) int {
	lQivwVxjomk :=  /*line :1*/ len(b6Z2P0)
	switch {
	case lQivwVxjomk == 0:
		return 0
	case lQivwVxjomk == 1:
		return  /*line :1*/ AXZRWQNqa(nhst_x, b6Z2P0[0])
	case lQivwVxjomk ==  /*line :1*/ len(nhst_x):
		if  /*line :1*/ H2kZ75hy(b6Z2P0, nhst_x) {
			return 0
		}
		return -1
	case lQivwVxjomk >  /*line :1*/ len(nhst_x):
		return -1
	case lQivwVxjomk <= bytealg.MaxLen:

		if  /*line :1*/ len(nhst_x) <= bytealg.MaxBruteForce {
			return  /*line :1*/ bytealg.Index(nhst_x, b6Z2P0)
		}
		wazBcxb := b6Z2P0[0]
		cehwfSI2Ahd := b6Z2P0[1]
		musLbUdLeb := 0
		qjUVLkM536S3 :=  /*line :1*/ len(nhst_x) - lQivwVxjomk + 1
		a8sfz7G1E := 0
		for musLbUdLeb < qjUVLkM536S3 {
			if nhst_x[musLbUdLeb] != wazBcxb {

				h9nlufqu :=  /*line :1*/ AXZRWQNqa(nhst_x[musLbUdLeb+1:qjUVLkM536S3], wazBcxb)
				if h9nlufqu < 0 {
					return -1
				}
				musLbUdLeb += h9nlufqu + 1
			}
			if nhst_x[musLbUdLeb+1] == cehwfSI2Ahd &&  /*line :1*/ H2kZ75hy(nhst_x[musLbUdLeb:musLbUdLeb+lQivwVxjomk], b6Z2P0) {
				return musLbUdLeb
			}
			a8sfz7G1E++
			musLbUdLeb++

			if a8sfz7G1E >  /*line :1*/ bytealg.Cutover(musLbUdLeb) {
				dIwXIFT :=  /*line :1*/ bytealg.Index(nhst_x[musLbUdLeb:], b6Z2P0)
				if dIwXIFT >= 0 {
					return dIwXIFT + musLbUdLeb
				}
				return -1
			}
		}
		return -1
	}
	wazBcxb := b6Z2P0[0]
	cehwfSI2Ahd := b6Z2P0[1]
	musLbUdLeb := 0
	a8sfz7G1E := 0
	qjUVLkM536S3 :=  /*line :1*/ len(nhst_x) - lQivwVxjomk + 1
	for musLbUdLeb < qjUVLkM536S3 {
		if nhst_x[musLbUdLeb] != wazBcxb {
			h9nlufqu :=  /*line :1*/ AXZRWQNqa(nhst_x[musLbUdLeb+1:qjUVLkM536S3], wazBcxb)
			if h9nlufqu < 0 {
				break
			}
			musLbUdLeb += h9nlufqu + 1
		}
		if nhst_x[musLbUdLeb+1] == cehwfSI2Ahd &&  /*line :1*/ H2kZ75hy(nhst_x[musLbUdLeb:musLbUdLeb+lQivwVxjomk], b6Z2P0) {
			return musLbUdLeb
		}
		musLbUdLeb++
		a8sfz7G1E++
		if a8sfz7G1E >= 4+musLbUdLeb>>4 && musLbUdLeb < qjUVLkM536S3 {

			xDiLck8 :=  /*line :1*/ bytealg.IndexRabinKarpBytes(nhst_x[musLbUdLeb:], b6Z2P0)
			if xDiLck8 < 0 {
				return -1
			}
			return musLbUdLeb + xDiLck8
		}
	}
	return -1
}







func FidBwOAXl2T(nhst_x, b6Z2P0 []byte) (j4Tqq77, igZA8ZG []byte, b6m87hNuH bool) {
	if musLbUdLeb :=  /*line :1*/ DxTUlq(nhst_x, b6Z2P0); musLbUdLeb >= 0 {
		return nhst_x[:musLbUdLeb], nhst_x[musLbUdLeb+ /*line :1*/ len(b6Z2P0):], true
	}
	return nhst_x, nil, false
}




func ELZHlQt(ka9IyYc0 []byte) []byte {
	if ka9IyYc0 == nil {
		return nil
	}
	return  /*line :1*/ append([]byte{}, ka9IyYc0...)
}







func UD17S4J(nhst_x, sMXnw0t6SSCa []byte) (igZA8ZG []byte, b6m87hNuH bool) {
	if ! /*line :1*/ V2Q3qKsmU(nhst_x, sMXnw0t6SSCa) {
		return nhst_x, false
	}
	return nhst_x[ /*line :1*/ len(sMXnw0t6SSCa):], true
}







func TcKiT1VHWbRK(nhst_x, dcbttECl46 []byte) (j4Tqq77 []byte, b6m87hNuH bool) {
	if ! /*line :1*/ Ar8BM7P(nhst_x, dcbttECl46) {
		return nhst_x, false
	}
	return nhst_x[: /*line :1*/ len(nhst_x)- /*line :1*/ len(dcbttECl46)], true
}
