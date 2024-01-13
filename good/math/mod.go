//line :1
package math












func LoDnahP_eoog(_BqkSz0, nNeKvlvK6 float64) float64 {
	if haveArchMod {
		return  /*line :1*/ tygUBzdk93aO(_BqkSz0, nNeKvlvK6)
	}
	return  /*line :1*/ iO_iNU5602a(_BqkSz0, nNeKvlvK6)
}

func iO_iNU5602a(_BqkSz0, nNeKvlvK6 float64) float64 {
	if nNeKvlvK6 == 0 ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0) ||  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ OIdLpDqq(nNeKvlvK6) {
		return  /*line :1*/ WG0hZIT4()
	}
	nNeKvlvK6 =  /*line :1*/ Abs(nNeKvlvK6)

	hIg3nCNnIrp3, wv5MQa :=  /*line :1*/ CuUl2eV(nNeKvlvK6)
	aBK_X3rbPc := _BqkSz0
	if _BqkSz0 < 0 {
		aBK_X3rbPc = -_BqkSz0
	}

	for aBK_X3rbPc >= nNeKvlvK6 {
		cs08nSYf4MtC, oTN6Zd344Kox :=  /*line :1*/ CuUl2eV(aBK_X3rbPc)
		if cs08nSYf4MtC < hIg3nCNnIrp3 {
			oTN6Zd344Kox = oTN6Zd344Kox - 1
		}
		aBK_X3rbPc = aBK_X3rbPc -  /*line :1*/ EU1SFL1cw(nNeKvlvK6, oTN6Zd344Kox-wv5MQa)
	}
	if _BqkSz0 < 0 {
		aBK_X3rbPc = -aBK_X3rbPc
	}
	return aBK_X3rbPc
}
