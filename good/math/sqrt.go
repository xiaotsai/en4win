//line :1
package math









func NeXs7bSyfaCD(_BqkSz0 float64) float64 {
	return  /*line :1*/ sqrt(_BqkSz0)
}

func sqrt(_BqkSz0 float64) float64 {

	switch {
	case _BqkSz0 == 0 ||  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return _BqkSz0
	case _BqkSz0 < 0:
		return  /*line :1*/ WG0hZIT4()
	}
	cCmjKOxt :=  /*line :1*/ GKIRmoHE(_BqkSz0)

	nF4u1Vnrit :=  /*line :1*/ int((cCmjKOxt >> shift) & mask)
	if nF4u1Vnrit == 0 {
		for cCmjKOxt&(1<<shift) == 0 {
			cCmjKOxt <<= 1
			nF4u1Vnrit--
		}
		nF4u1Vnrit++
	}
	nF4u1Vnrit -= bias
	cCmjKOxt &^= mask << shift
	cCmjKOxt |= 1 << shift
	if nF4u1Vnrit&1 == 1 {
		cCmjKOxt <<= 1
	}
	nF4u1Vnrit >>= 1

	cCmjKOxt <<= 1
	var gNDglc3F2QV, aK8LO1 uint64	
	aBK_X3rbPc :=  /*line :1*/ uint64(1 << (shift + 1))
	for aBK_X3rbPc != 0 {
		aX6xkg_1G := aK8LO1 + aBK_X3rbPc
		if aX6xkg_1G <= cCmjKOxt {
			aK8LO1 = aX6xkg_1G + aBK_X3rbPc
			cCmjKOxt -= aX6xkg_1G
			gNDglc3F2QV += aBK_X3rbPc
		}
		cCmjKOxt <<= 1
		aBK_X3rbPc >>= 1
	}

	if cCmjKOxt != 0 {
		gNDglc3F2QV += gNDglc3F2QV & 1
	}
	cCmjKOxt = gNDglc3F2QV>>1 +  /*line :1*/ uint64(nF4u1Vnrit-1+bias)<<shift
	return  /*line :1*/ QUB2Kf63p(cCmjKOxt)
}
