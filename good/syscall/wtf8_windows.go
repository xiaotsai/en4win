//line :1
package bUKeamF

import (
	utf16 "DtJCLKevRp"
	utf8 "CuAc0pPZwf"
)

const (
	surr1	= 0xd800
	surr2	= 0xdc00
	surr3	= 0xe000

	tx	= 0b10000000
	t3	= 0b11100000
	maskx	= 0b00111111
	mask3	= 0b00001111

	rune1Max	= 1<<7 - 1
	rune2Max	= 1<<11 - 1
)



func jsi_qEcoyU9L(wzPhJhIFoI string, eun1Jlud5A []uint16) []uint16 {
	for hA99q3EOK := 0; hA99q3EOK <  /*line :1*/ len(wzPhJhIFoI); {

		iVSjWu, jFhAAWnAX :=  /*line :1*/ utf8.HVAV7aTqxzg(wzPhJhIFoI[hA99q3EOK:])
		if iVSjWu == utf8.RuneError {

			if n84jqC0Z3wU := wzPhJhIFoI[hA99q3EOK:];  /*line :1*/ len(n84jqC0Z3wU) >= 3 && n84jqC0Z3wU[0] == 0xED && 0xA0 <= n84jqC0Z3wU[1] && n84jqC0Z3wU[1] <= 0xBF && 0x80 <= n84jqC0Z3wU[2] && n84jqC0Z3wU[2] <= 0xBF {
				iVSjWu =  /*line :1*/ rune(n84jqC0Z3wU[0]&mask3)<<12 +  /*line :1*/ rune(n84jqC0Z3wU[1]&maskx)<<6 +  /*line :1*/ rune(n84jqC0Z3wU[2]&maskx)
				eun1Jlud5A =  /*line :1*/ append(eun1Jlud5A,  /*line :1*/ uint16(iVSjWu))
				hA99q3EOK += 3
				continue
			}
		}
		hA99q3EOK += jFhAAWnAX
		eun1Jlud5A =  /*line :1*/ utf16.J23fS69TUl7H(eun1Jlud5A, iVSjWu)
	}
	return eun1Jlud5A
}



func mtdGDMk8v(wzPhJhIFoI []uint16, eun1Jlud5A []byte) []byte {
	for hA99q3EOK := 0; hA99q3EOK <  /*line :1*/ len(wzPhJhIFoI); hA99q3EOK++ {
		var mRWd3pJY_E rune
		switch iVSjWu := wzPhJhIFoI[hA99q3EOK]; {
		case iVSjWu < surr1, surr3 <= iVSjWu:

			mRWd3pJY_E =  /*line :1*/ rune(iVSjWu)
		case surr1 <= iVSjWu && iVSjWu < surr2 && hA99q3EOK+1 <  /*line :1*/ len(wzPhJhIFoI) &&
			surr2 <= wzPhJhIFoI[hA99q3EOK+1] && wzPhJhIFoI[hA99q3EOK+1] < surr3:

			mRWd3pJY_E =  /*line :1*/ utf16.Rw3auy7U( /*line :1*/ rune(iVSjWu),  /*line :1*/ rune(wzPhJhIFoI[hA99q3EOK+1]))
			hA99q3EOK++
		default:

			mRWd3pJY_E =  /*line :1*/ rune(iVSjWu)
			if mRWd3pJY_E > utf8.MaxRune {
				mRWd3pJY_E = utf8.RuneError
			}
			eun1Jlud5A =  /*line :1*/ append(eun1Jlud5A, t3| /*line :1*/ byte(mRWd3pJY_E>>12), tx| /*line :1*/ byte(mRWd3pJY_E>>6)&maskx, tx| /*line :1*/ byte(mRWd3pJY_E)&maskx)
			continue
		}
		eun1Jlud5A =  /*line :1*/ utf8.Ht7oMzd(eun1Jlud5A, mRWd3pJY_E)
	}
	return eun1Jlud5A
}
