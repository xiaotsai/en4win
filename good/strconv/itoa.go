//line :1
package zBESu2SrRjP

import bits "math/bits"

const fastSmalls = true	




func DXSqN7(cDhV_2D uint64, xsj80lMaPD int) string {
	if fastSmalls && cDhV_2D < nSmalls && xsj80lMaPD == 10 {
		return  /*line :1*/ ddzELfcCHa( /*line :1*/ int(cDhV_2D))
	}
	_, a0M8QSnCR :=  /*line :1*/ lHgNcSQTJz1(nil, cDhV_2D, xsj80lMaPD, false, false)
	return a0M8QSnCR
}




func DV3SPN_4O(cDhV_2D int64, xsj80lMaPD int) string {
	if fastSmalls && 0 <= cDhV_2D && cDhV_2D < nSmalls && xsj80lMaPD == 10 {
		return  /*line :1*/ ddzELfcCHa( /*line :1*/ int(cDhV_2D))
	}
	_, a0M8QSnCR :=  /*line :1*/ lHgNcSQTJz1(nil,  /*line :1*/ uint64(cDhV_2D), xsj80lMaPD, cDhV_2D < 0, false)
	return a0M8QSnCR
}


func ZW1_FxJtq(cDhV_2D int) string {
	return  /*line :1*/ DV3SPN_4O( /*line :1*/ int64(cDhV_2D), 10)
}



func YOFr9xxiI(l_A2Ytb []byte, cDhV_2D int64, xsj80lMaPD int) []byte {
	if fastSmalls && 0 <= cDhV_2D && cDhV_2D < nSmalls && xsj80lMaPD == 10 {
		return  /*line :1*/ append(l_A2Ytb,  /*line :1*/ ddzELfcCHa( /*line :1*/ int(cDhV_2D))...)
	}
	l_A2Ytb, _ =  /*line :1*/ lHgNcSQTJz1(l_A2Ytb,  /*line :1*/ uint64(cDhV_2D), xsj80lMaPD, cDhV_2D < 0, true)
	return l_A2Ytb
}



func DxG2yRkVZO(l_A2Ytb []byte, cDhV_2D uint64, xsj80lMaPD int) []byte {
	if fastSmalls && cDhV_2D < nSmalls && xsj80lMaPD == 10 {
		return  /*line :1*/ append(l_A2Ytb,  /*line :1*/ ddzELfcCHa( /*line :1*/ int(cDhV_2D))...)
	}
	l_A2Ytb, _ =  /*line :1*/ lHgNcSQTJz1(l_A2Ytb, cDhV_2D, xsj80lMaPD, false, true)
	return l_A2Ytb
}


func ddzELfcCHa(cDhV_2D int) string {
	if cDhV_2D < 10 {
		return digits[cDhV_2D : cDhV_2D+1]
	}
	return smallsString[cDhV_2D*2 : cDhV_2D*2+2]
}

const nSmalls = 100

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

const host32bit = ^ /*line :1*/ uint(0)>>32 == 0

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"






func lHgNcSQTJz1(l_A2Ytb []byte, dnNosEuwQ uint64, xsj80lMaPD int, an5n8FniWt67, w6j3ZY bool) (bkcwawBry []byte, a0M8QSnCR string) {
	if xsj80lMaPD < 2 || xsj80lMaPD >  /*line :1*/ len(digits) {
		 /*line :1*/ panic("strconv: illegal AppendInt/FormatInt base")
	}

	var zxZt5L [64 + 1]byte	
	cDhV_2D :=  /*line :1*/ len(zxZt5L)

	if an5n8FniWt67 {
		dnNosEuwQ = -dnNosEuwQ
	}

	if xsj80lMaPD == 10 {

		if host32bit {

			for dnNosEuwQ >= 1e9 {

				oUyevz5E2Vjr := dnNosEuwQ / 1e9
				dhzuIMF3FD :=  /*line :1*/ uint(dnNosEuwQ - oUyevz5E2Vjr*1e9)
				for dTPavgbkmK := 4; dTPavgbkmK > 0; dTPavgbkmK-- {
					pwP7gqLPyt := dhzuIMF3FD % 100 * 2
					dhzuIMF3FD /= 100
					cDhV_2D -= 2
					zxZt5L[cDhV_2D+1] = smallsString[pwP7gqLPyt+1]
					zxZt5L[cDhV_2D+0] = smallsString[pwP7gqLPyt+0]
				}

				cDhV_2D--
				zxZt5L[cDhV_2D] = smallsString[dhzuIMF3FD*2+1]

				dnNosEuwQ = oUyevz5E2Vjr
			}

		}

		dhzuIMF3FD :=  /*line :1*/ uint(dnNosEuwQ)
		for dhzuIMF3FD >= 100 {
			pwP7gqLPyt := dhzuIMF3FD % 100 * 2
			dhzuIMF3FD /= 100
			cDhV_2D -= 2
			zxZt5L[cDhV_2D+1] = smallsString[pwP7gqLPyt+1]
			zxZt5L[cDhV_2D+0] = smallsString[pwP7gqLPyt+0]
		}

		pwP7gqLPyt := dhzuIMF3FD * 2
		cDhV_2D--
		zxZt5L[cDhV_2D] = smallsString[pwP7gqLPyt+1]
		if dhzuIMF3FD >= 10 {
			cDhV_2D--
			zxZt5L[cDhV_2D] = smallsString[pwP7gqLPyt]
		}

	} else if  /*line :1*/ r_I8uF5i4WW(xsj80lMaPD) {

		naILJ1HmxfWn :=  /*line :1*/ uint( /*line :1*/ bits.AjW10JsK( /*line :1*/ uint(xsj80lMaPD))) & 7
		o7OjhYc5wQns :=  /*line :1*/ uint64(xsj80lMaPD)
		csRwMg :=  /*line :1*/ uint(xsj80lMaPD) - 1
		for dnNosEuwQ >= o7OjhYc5wQns {
			cDhV_2D--
			zxZt5L[cDhV_2D] = digits[ /*line :1*/ uint(dnNosEuwQ)&csRwMg]
			dnNosEuwQ >>= naILJ1HmxfWn
		}

		cDhV_2D--
		zxZt5L[cDhV_2D] = digits[ /*line :1*/ uint(dnNosEuwQ)]
	} else {

		o7OjhYc5wQns :=  /*line :1*/ uint64(xsj80lMaPD)
		for dnNosEuwQ >= o7OjhYc5wQns {
			cDhV_2D--

			oUyevz5E2Vjr := dnNosEuwQ / o7OjhYc5wQns
			zxZt5L[cDhV_2D] = digits[ /*line :1*/ uint(dnNosEuwQ-oUyevz5E2Vjr*o7OjhYc5wQns)]
			dnNosEuwQ = oUyevz5E2Vjr
		}

		cDhV_2D--
		zxZt5L[cDhV_2D] = digits[ /*line :1*/ uint(dnNosEuwQ)]
	}

	if an5n8FniWt67 {
		cDhV_2D--
		zxZt5L[cDhV_2D] = '-'
	}

	if w6j3ZY {
		bkcwawBry =  /*line :1*/ append(l_A2Ytb, zxZt5L[cDhV_2D:]...)
		return
	}
	a0M8QSnCR =  /*line :1*/ string(zxZt5L[cDhV_2D:])
	return
}

func r_I8uF5i4WW(eswrOOpIaF int) bool {
	return eswrOOpIaF&(eswrOOpIaF-1) == 0
}
