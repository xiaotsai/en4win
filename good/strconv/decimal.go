//line :1
package zBESu2SrRjP

type fJphQdAo struct {
	dXlUa8ko	[800]byte	
	ejR6KwzBSk	int	
	qvuKHF9jDC	int	
	hhg70L_qWPbD	bool	
	lpkYii9C	bool	
}

func (zxZt5L *fJphQdAo) String() string {
	wESnmZiAO := 10 + zxZt5L.ejR6KwzBSk
	if zxZt5L.qvuKHF9jDC > 0 {
		wESnmZiAO += zxZt5L.qvuKHF9jDC
	}
	if zxZt5L.qvuKHF9jDC < 0 {
		wESnmZiAO += -zxZt5L.qvuKHF9jDC
	}

	uPg2dG3l8KRn :=  /*line :1*/ make([]byte, wESnmZiAO)
	iVet75r18A := 0
	switch {
	case zxZt5L.ejR6KwzBSk == 0:
		return "0"

	case zxZt5L.qvuKHF9jDC <= 0:

		uPg2dG3l8KRn[iVet75r18A] = '0'
		iVet75r18A++
		uPg2dG3l8KRn[iVet75r18A] = '.'
		iVet75r18A++
		iVet75r18A +=  /*line :1*/ taI9BYt5pB2(uPg2dG3l8KRn[iVet75r18A : iVet75r18A+-zxZt5L.qvuKHF9jDC])
		iVet75r18A +=  /*line :1*/ copy(uPg2dG3l8KRn[iVet75r18A:], zxZt5L.dXlUa8ko[0:zxZt5L.ejR6KwzBSk])

	case zxZt5L.qvuKHF9jDC < zxZt5L.ejR6KwzBSk:

		iVet75r18A +=  /*line :1*/ copy(uPg2dG3l8KRn[iVet75r18A:], zxZt5L.dXlUa8ko[0:zxZt5L.qvuKHF9jDC])
		uPg2dG3l8KRn[iVet75r18A] = '.'
		iVet75r18A++
		iVet75r18A +=  /*line :1*/ copy(uPg2dG3l8KRn[iVet75r18A:], zxZt5L.dXlUa8ko[zxZt5L.qvuKHF9jDC:zxZt5L.ejR6KwzBSk])

	default:

		iVet75r18A +=  /*line :1*/ copy(uPg2dG3l8KRn[iVet75r18A:], zxZt5L.dXlUa8ko[0:zxZt5L.ejR6KwzBSk])
		iVet75r18A +=  /*line :1*/ taI9BYt5pB2(uPg2dG3l8KRn[iVet75r18A : iVet75r18A+zxZt5L.qvuKHF9jDC-zxZt5L.ejR6KwzBSk])
	}
	return  /*line :1*/ string(uPg2dG3l8KRn[0:iVet75r18A])
}

func taI9BYt5pB2(l_A2Ytb []byte) int {
	for cDhV_2D := range l_A2Ytb {
		l_A2Ytb[cDhV_2D] = '0'
	}
	return  /*line :1*/ len(l_A2Ytb)
}




func q5wLUB(zxZt5L *fJphQdAo) {
	for zxZt5L.ejR6KwzBSk > 0 && zxZt5L.dXlUa8ko[zxZt5L.ejR6KwzBSk-1] == '0' {
		zxZt5L.ejR6KwzBSk--
	}
	if zxZt5L.ejR6KwzBSk == 0 {
		zxZt5L.qvuKHF9jDC = 0
	}
}


func (zxZt5L *fJphQdAo) Assign(ggqcIE uint64) {
	var uPg2dG3l8KRn [24]byte

	wESnmZiAO := 0
	for ggqcIE > 0 {
		zCsD112MLqZf := ggqcIE / 10
		ggqcIE -= 10 * zCsD112MLqZf
		uPg2dG3l8KRn[wESnmZiAO] =  /*line :1*/ byte(ggqcIE + '0')
		wESnmZiAO++
		ggqcIE = zCsD112MLqZf
	}

	zxZt5L.ejR6KwzBSk = 0
	for wESnmZiAO--; wESnmZiAO >= 0; wESnmZiAO-- {
		zxZt5L.dXlUa8ko[zxZt5L.ejR6KwzBSk] = uPg2dG3l8KRn[wESnmZiAO]
		zxZt5L.ejR6KwzBSk++
	}
	zxZt5L.qvuKHF9jDC = zxZt5L.ejR6KwzBSk
	 /*line :1*/ q5wLUB(zxZt5L)
}



const uintSize = 32 << (^ /*line :1*/ uint(0) >> 63)
const maxShift = uintSize - 4


func ugKRN3j7ZgMa(zxZt5L *fJphQdAo, tyBlZruomOL uint) {
	vaeEFNU6wEP := 0
	iVet75r18A := 0

	
	var wESnmZiAO uint
	for ; wESnmZiAO>>tyBlZruomOL == 0; vaeEFNU6wEP++ {
		if vaeEFNU6wEP >= zxZt5L.ejR6KwzBSk {
			if wESnmZiAO == 0 {

				zxZt5L.ejR6KwzBSk = 0
				return
			}
			for wESnmZiAO>>tyBlZruomOL == 0 {
				wESnmZiAO = wESnmZiAO * 10
				vaeEFNU6wEP++
			}
			break
		}
		uTN3BXbMd :=  /*line :1*/ uint(zxZt5L.dXlUa8ko[vaeEFNU6wEP])
		wESnmZiAO = wESnmZiAO*10 + uTN3BXbMd - '0'
	}
	zxZt5L.qvuKHF9jDC -= vaeEFNU6wEP - 1

	var wIXudG uint = (1 << tyBlZruomOL) - 1

	for ; vaeEFNU6wEP < zxZt5L.ejR6KwzBSk; vaeEFNU6wEP++ {
		uTN3BXbMd :=  /*line :1*/ uint(zxZt5L.dXlUa8ko[vaeEFNU6wEP])
		aypDIVd4gp := wESnmZiAO >> tyBlZruomOL
		wESnmZiAO &= wIXudG
		zxZt5L.dXlUa8ko[iVet75r18A] =  /*line :1*/ byte(aypDIVd4gp + '0')
		iVet75r18A++
		wESnmZiAO = wESnmZiAO*10 + uTN3BXbMd - '0'
	}

	for wESnmZiAO > 0 {
		aypDIVd4gp := wESnmZiAO >> tyBlZruomOL
		wESnmZiAO &= wIXudG
		if iVet75r18A <  /*line :1*/ len(zxZt5L.dXlUa8ko) {
			zxZt5L.dXlUa8ko[iVet75r18A] =  /*line :1*/ byte(aypDIVd4gp + '0')
			iVet75r18A++
		} else if aypDIVd4gp > 0 {
			zxZt5L.lpkYii9C = true
		}
		wESnmZiAO = wESnmZiAO * 10
	}

	zxZt5L.ejR6KwzBSk = iVet75r18A
	 /*line :1*/ q5wLUB(zxZt5L)
}

type gOcx7H2DT struct {
	rzb1dV	int	
	wqafLVLW	string	
}

var y5MVeqS = []gOcx7H2DT{

	{0, ""},
	{1, "5"},
	{1, "25"},
	{1, "125"},
	{2, "625"},
	{2, "3125"},
	{2, "15625"},
	{3, "78125"},
	{3, "390625"},
	{3, "1953125"},
	{4, "9765625"},
	{4, "48828125"},
	{4, "244140625"},
	{4, "1220703125"},
	{5, "6103515625"},
	{5, "30517578125"},
	{5, "152587890625"},
	{6, "762939453125"},
	{6, "3814697265625"},
	{6, "19073486328125"},
	{7, "95367431640625"},
	{7, "476837158203125"},
	{7, "2384185791015625"},
	{7, "11920928955078125"},
	{8, "59604644775390625"},
	{8, "298023223876953125"},
	{8, "1490116119384765625"},
	{9, "7450580596923828125"},
	{9, "37252902984619140625"},
	{9, "186264514923095703125"},
	{10, "931322574615478515625"},
	{10, "4656612873077392578125"},
	{10, "23283064365386962890625"},
	{10, "116415321826934814453125"},
	{11, "582076609134674072265625"},
	{11, "2910383045673370361328125"},
	{11, "14551915228366851806640625"},
	{12, "72759576141834259033203125"},
	{12, "363797880709171295166015625"},
	{12, "1818989403545856475830078125"},
	{13, "9094947017729282379150390625"},
	{13, "45474735088646411895751953125"},
	{13, "227373675443232059478759765625"},
	{13, "1136868377216160297393798828125"},
	{14, "5684341886080801486968994140625"},
	{14, "28421709430404007434844970703125"},
	{14, "142108547152020037174224853515625"},
	{15, "710542735760100185871124267578125"},
	{15, "3552713678800500929355621337890625"},
	{15, "17763568394002504646778106689453125"},
	{16, "88817841970012523233890533447265625"},
	{16, "444089209850062616169452667236328125"},
	{16, "2220446049250313080847263336181640625"},
	{16, "11102230246251565404236316680908203125"},
	{17, "55511151231257827021181583404541015625"},
	{17, "277555756156289135105907917022705078125"},
	{17, "1387778780781445675529539585113525390625"},
	{18, "6938893903907228377647697925567626953125"},
	{18, "34694469519536141888238489627838134765625"},
	{18, "173472347597680709441192448139190673828125"},
	{19, "867361737988403547205962240695953369140625"},
}


func hBkeecmr(o7OjhYc5wQns []byte, a0M8QSnCR string) bool {
	for cDhV_2D := 0; cDhV_2D <  /*line :1*/ len(a0M8QSnCR); cDhV_2D++ {
		if cDhV_2D >=  /*line :1*/ len(o7OjhYc5wQns) {
			return true
		}
		if o7OjhYc5wQns[cDhV_2D] != a0M8QSnCR[cDhV_2D] {
			return o7OjhYc5wQns[cDhV_2D] < a0M8QSnCR[cDhV_2D]
		}
	}
	return false
}


func o_bB2xsb_YSH(zxZt5L *fJphQdAo, tyBlZruomOL uint) {
	g5P4r0d := y5MVeqS[tyBlZruomOL].rzb1dV
	if  /*line :1*/ hBkeecmr(zxZt5L.dXlUa8ko[0:zxZt5L.ejR6KwzBSk], y5MVeqS[tyBlZruomOL].wqafLVLW) {
		g5P4r0d--
	}

	vaeEFNU6wEP := zxZt5L.ejR6KwzBSk
	iVet75r18A := zxZt5L.ejR6KwzBSk + g5P4r0d

	
	var wESnmZiAO uint
	for vaeEFNU6wEP--; vaeEFNU6wEP >= 0; vaeEFNU6wEP-- {
		wESnmZiAO += ( /*line :1*/ uint(zxZt5L.dXlUa8ko[vaeEFNU6wEP]) - '0') << tyBlZruomOL
		rDEF7o5rfT := wESnmZiAO / 10
		pNRglmCkMAx := wESnmZiAO - 10*rDEF7o5rfT
		iVet75r18A--
		if iVet75r18A <  /*line :1*/ len(zxZt5L.dXlUa8ko) {
			zxZt5L.dXlUa8ko[iVet75r18A] =  /*line :1*/ byte(pNRglmCkMAx + '0')
		} else if pNRglmCkMAx != 0 {
			zxZt5L.lpkYii9C = true
		}
		wESnmZiAO = rDEF7o5rfT
	}

	for wESnmZiAO > 0 {
		rDEF7o5rfT := wESnmZiAO / 10
		pNRglmCkMAx := wESnmZiAO - 10*rDEF7o5rfT
		iVet75r18A--
		if iVet75r18A <  /*line :1*/ len(zxZt5L.dXlUa8ko) {
			zxZt5L.dXlUa8ko[iVet75r18A] =  /*line :1*/ byte(pNRglmCkMAx + '0')
		} else if pNRglmCkMAx != 0 {
			zxZt5L.lpkYii9C = true
		}
		wESnmZiAO = rDEF7o5rfT
	}

	zxZt5L.ejR6KwzBSk += g5P4r0d
	if zxZt5L.ejR6KwzBSk >=  /*line :1*/ len(zxZt5L.dXlUa8ko) {
		zxZt5L.ejR6KwzBSk =  /*line :1*/ len(zxZt5L.dXlUa8ko)
	}
	zxZt5L.qvuKHF9jDC += g5P4r0d
	 /*line :1*/ q5wLUB(zxZt5L)
}


func (zxZt5L *fJphQdAo) Shift(tyBlZruomOL int) {
	switch {
	case zxZt5L.ejR6KwzBSk == 0:

	case tyBlZruomOL > 0:
		for tyBlZruomOL > maxShift {
			 /*line :1*/ o_bB2xsb_YSH(zxZt5L, maxShift)
			tyBlZruomOL -= maxShift
		}
		 /*line :1*/ o_bB2xsb_YSH(zxZt5L,  /*line :1*/ uint(tyBlZruomOL))
	case tyBlZruomOL < 0:
		for tyBlZruomOL < -maxShift {
			 /*line :1*/ ugKRN3j7ZgMa(zxZt5L, maxShift)
			tyBlZruomOL += maxShift
		}
		 /*line :1*/ ugKRN3j7ZgMa(zxZt5L,  /*line :1*/ uint(-tyBlZruomOL))
	}
}


func rPOhYV_7J(zxZt5L *fJphQdAo, b9NNFc_Hmo_7 int) bool {
	if b9NNFc_Hmo_7 < 0 || b9NNFc_Hmo_7 >= zxZt5L.ejR6KwzBSk {
		return false
	}
	if zxZt5L.dXlUa8ko[b9NNFc_Hmo_7] == '5' && b9NNFc_Hmo_7+1 == zxZt5L.ejR6KwzBSk {

		if zxZt5L.lpkYii9C {
			return true
		}
		return b9NNFc_Hmo_7 > 0 && (zxZt5L.dXlUa8ko[b9NNFc_Hmo_7-1]-'0')%2 != 0
	}

	return zxZt5L.dXlUa8ko[b9NNFc_Hmo_7] >= '5'
}





func (zxZt5L *fJphQdAo) Round(b9NNFc_Hmo_7 int) {
	if b9NNFc_Hmo_7 < 0 || b9NNFc_Hmo_7 >= zxZt5L.ejR6KwzBSk {
		return
	}
	if  /*line :1*/ rPOhYV_7J(zxZt5L, b9NNFc_Hmo_7) {
		 /*line :1*/ zxZt5L.RoundUp(b9NNFc_Hmo_7)
	} else {
		 /*line :1*/ zxZt5L.RoundDown(b9NNFc_Hmo_7)
	}
}


func (zxZt5L *fJphQdAo) RoundDown(b9NNFc_Hmo_7 int) {
	if b9NNFc_Hmo_7 < 0 || b9NNFc_Hmo_7 >= zxZt5L.ejR6KwzBSk {
		return
	}
	zxZt5L.ejR6KwzBSk = b9NNFc_Hmo_7
	 /*line :1*/ q5wLUB(zxZt5L)
}


func (zxZt5L *fJphQdAo) RoundUp(b9NNFc_Hmo_7 int) {
	if b9NNFc_Hmo_7 < 0 || b9NNFc_Hmo_7 >= zxZt5L.ejR6KwzBSk {
		return
	}

	for cDhV_2D := b9NNFc_Hmo_7 - 1; cDhV_2D >= 0; cDhV_2D-- {
		uTN3BXbMd := zxZt5L.dXlUa8ko[cDhV_2D]
		if uTN3BXbMd < '9' {
			zxZt5L.dXlUa8ko[cDhV_2D]++
			zxZt5L.ejR6KwzBSk = cDhV_2D + 1
			return
		}
	}

	zxZt5L.dXlUa8ko[0] = '1'
	zxZt5L.ejR6KwzBSk = 1
	zxZt5L.qvuKHF9jDC++
}



func (zxZt5L *fJphQdAo) RoundedInteger() uint64 {
	if zxZt5L.qvuKHF9jDC > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var cDhV_2D int
	wESnmZiAO :=  /*line :1*/ uint64(0)
	for cDhV_2D = 0; cDhV_2D < zxZt5L.qvuKHF9jDC && cDhV_2D < zxZt5L.ejR6KwzBSk; cDhV_2D++ {
		wESnmZiAO = wESnmZiAO*10 +  /*line :1*/ uint64(zxZt5L.dXlUa8ko[cDhV_2D]-'0')
	}
	for ; cDhV_2D < zxZt5L.qvuKHF9jDC; cDhV_2D++ {
		wESnmZiAO *= 10
	}
	if  /*line :1*/ rPOhYV_7J(zxZt5L, zxZt5L.qvuKHF9jDC) {
		wESnmZiAO++
	}
	return wESnmZiAO
}
