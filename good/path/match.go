//line :1
package jn9Psh1aa_7H

import (
	errors "iAHoxjmM"
	"internal/bytealg"
	utf8 "CuAc0pPZwf"
)


var VVxvifIdGY =  /*line :1*/ errors.Su6g6hRoi9X("syntax error in pattern")






















func K9FvN0A5kW42(uAJ26yLz, bOXS_Lk1b8 string) (gpQb9rVT bool, to2Ta5aX9rji error) {
Pattern:
	for  /*line :1*/ len(uAJ26yLz) > 0 {
		var iaUurw bool
		var aRg0mrNIGeN string
		iaUurw, aRg0mrNIGeN, uAJ26yLz =  /*line :1*/ yyFd_00(uAJ26yLz)
		if iaUurw && aRg0mrNIGeN == "" {

			return  /*line :1*/ bytealg.IndexByteString(bOXS_Lk1b8, '/') < 0, nil
		}

		tRQSYEGoqUCu, iuRMvseCaT, to2Ta5aX9rji :=  /*line :1*/ v_tx1fAq3J(aRg0mrNIGeN, bOXS_Lk1b8)

		if iuRMvseCaT && ( /*line :1*/ len(tRQSYEGoqUCu) == 0 ||  /*line :1*/ len(uAJ26yLz) > 0) {
			bOXS_Lk1b8 = tRQSYEGoqUCu
			continue
		}
		if to2Ta5aX9rji != nil {
			return false, to2Ta5aX9rji
		}
		if iaUurw {

			for tr99J_O := 0; tr99J_O <  /*line :1*/ len(bOXS_Lk1b8) && bOXS_Lk1b8[tr99J_O] != '/'; tr99J_O++ {
				tRQSYEGoqUCu, iuRMvseCaT, to2Ta5aX9rji :=  /*line :1*/ v_tx1fAq3J(aRg0mrNIGeN, bOXS_Lk1b8[tr99J_O+1:])
				if iuRMvseCaT {

					if  /*line :1*/ len(uAJ26yLz) == 0 &&  /*line :1*/ len(tRQSYEGoqUCu) > 0 {
						continue
					}
					bOXS_Lk1b8 = tRQSYEGoqUCu
					continue Pattern
				}
				if to2Ta5aX9rji != nil {
					return false, to2Ta5aX9rji
				}
			}
		}

		for  /*line :1*/ len(uAJ26yLz) > 0 {
			_, aRg0mrNIGeN, uAJ26yLz =  /*line :1*/ yyFd_00(uAJ26yLz)
			if _, _, to2Ta5aX9rji :=  /*line :1*/ v_tx1fAq3J(aRg0mrNIGeN, ""); to2Ta5aX9rji != nil {
				return false, to2Ta5aX9rji
			}
		}
		return false, nil
	}
	return  /*line :1*/ len(bOXS_Lk1b8) == 0, nil
}



func yyFd_00(uAJ26yLz string) (iaUurw bool, aRg0mrNIGeN, mnHjcjU3 string) {
	for  /*line :1*/ len(uAJ26yLz) > 0 && uAJ26yLz[0] == '*' {
		uAJ26yLz = uAJ26yLz[1:]
		iaUurw = true
	}
	e6FHN932QdHb := false
	var tr99J_O int
Scan:
	for tr99J_O = 0; tr99J_O <  /*line :1*/ len(uAJ26yLz); tr99J_O++ {
		switch uAJ26yLz[tr99J_O] {
		case '\\':

			if tr99J_O+1 <  /*line :1*/ len(uAJ26yLz) {
				tr99J_O++
			}
		case '[':
			e6FHN932QdHb = true
		case ']':
			e6FHN932QdHb = false
		case '*':
			if !e6FHN932QdHb {
				break Scan
			}
		}
	}
	return iaUurw, uAJ26yLz[0:tr99J_O], uAJ26yLz[tr99J_O:]
}




func v_tx1fAq3J(aRg0mrNIGeN, kSAda01ZGvqp string) (mnHjcjU3 string, iuRMvseCaT bool, to2Ta5aX9rji error) {

	qEZYTKES_K9a := false
	for  /*line :1*/ len(aRg0mrNIGeN) > 0 {
		if !qEZYTKES_K9a &&  /*line :1*/ len(kSAda01ZGvqp) == 0 {
			qEZYTKES_K9a = true
		}
		switch aRg0mrNIGeN[0] {
		case '[':
			
			var aC96awU rune
			if !qEZYTKES_K9a {
				var sfulV7P7QO44 int
				aC96awU, sfulV7P7QO44 =  /*line :1*/ utf8.HVAV7aTqxzg(kSAda01ZGvqp)
				kSAda01ZGvqp = kSAda01ZGvqp[sfulV7P7QO44:]
			}
			aRg0mrNIGeN = aRg0mrNIGeN[1:]

			jVgXW7oAfRu := false
			if  /*line :1*/ len(aRg0mrNIGeN) > 0 && aRg0mrNIGeN[0] == '^' {
				jVgXW7oAfRu = true
				aRg0mrNIGeN = aRg0mrNIGeN[1:]
			}

			ff4LmrYi7nI := false
			tucbw7VT := 0
			for {
				if  /*line :1*/ len(aRg0mrNIGeN) > 0 && aRg0mrNIGeN[0] == ']' && tucbw7VT > 0 {
					aRg0mrNIGeN = aRg0mrNIGeN[1:]
					break
				}
				var uZQuOI, qKwaUq rune
				if uZQuOI, aRg0mrNIGeN, to2Ta5aX9rji =  /*line :1*/ e7rBtHRtU3C(aRg0mrNIGeN); to2Ta5aX9rji != nil {
					return "", false, to2Ta5aX9rji
				}
				qKwaUq = uZQuOI
				if aRg0mrNIGeN[0] == '-' {
					if qKwaUq, aRg0mrNIGeN, to2Ta5aX9rji =  /*line :1*/ e7rBtHRtU3C(aRg0mrNIGeN[1:]); to2Ta5aX9rji != nil {
						return "", false, to2Ta5aX9rji
					}
				}
				if uZQuOI <= aC96awU && aC96awU <= qKwaUq {
					ff4LmrYi7nI = true
				}
				tucbw7VT++
			}
			if ff4LmrYi7nI == jVgXW7oAfRu {
				qEZYTKES_K9a = true
			}

		case '?':
			if !qEZYTKES_K9a {
				if kSAda01ZGvqp[0] == '/' {
					qEZYTKES_K9a = true
				}
				_, sfulV7P7QO44 :=  /*line :1*/ utf8.HVAV7aTqxzg(kSAda01ZGvqp)
				kSAda01ZGvqp = kSAda01ZGvqp[sfulV7P7QO44:]
			}
			aRg0mrNIGeN = aRg0mrNIGeN[1:]

		case '\\':
			aRg0mrNIGeN = aRg0mrNIGeN[1:]
			if  /*line :1*/ len(aRg0mrNIGeN) == 0 {
				return "", false, VVxvifIdGY
			}
			fallthrough

		default:
			if !qEZYTKES_K9a {
				if aRg0mrNIGeN[0] != kSAda01ZGvqp[0] {
					qEZYTKES_K9a = true
				}
				kSAda01ZGvqp = kSAda01ZGvqp[1:]
			}
			aRg0mrNIGeN = aRg0mrNIGeN[1:]
		}
	}
	if qEZYTKES_K9a {
		return "", false, nil
	}
	return kSAda01ZGvqp, true, nil
}


func e7rBtHRtU3C(aRg0mrNIGeN string) (aC96awU rune, fIS37LZ4Taa string, to2Ta5aX9rji error) {
	if  /*line :1*/ len(aRg0mrNIGeN) == 0 || aRg0mrNIGeN[0] == '-' || aRg0mrNIGeN[0] == ']' {
		to2Ta5aX9rji = VVxvifIdGY
		return
	}
	if aRg0mrNIGeN[0] == '\\' {
		aRg0mrNIGeN = aRg0mrNIGeN[1:]
		if  /*line :1*/ len(aRg0mrNIGeN) == 0 {
			to2Ta5aX9rji = VVxvifIdGY
			return
		}
	}
	aC96awU, sfulV7P7QO44 :=  /*line :1*/ utf8.HVAV7aTqxzg(aRg0mrNIGeN)
	if aC96awU == utf8.RuneError && sfulV7P7QO44 == 1 {
		to2Ta5aX9rji = VVxvifIdGY
	}
	fIS37LZ4Taa = aRg0mrNIGeN[sfulV7P7QO44:]
	if  /*line :1*/ len(fIS37LZ4Taa) == 0 {
		to2Ta5aX9rji = VVxvifIdGY
	}
	return
}
