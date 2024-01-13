//line :1
package math











func CuUl2eV(ygoarCBO4et float64) (cmLS3KfTv float64, nF4u1Vnrit int) {
	if haveArchFrexp {
		return  /*line :1*/ hCM1Np(ygoarCBO4et)
	}
	return  /*line :1*/ rmWUA1W(ygoarCBO4et)
}

func rmWUA1W(ygoarCBO4et float64) (cmLS3KfTv float64, nF4u1Vnrit int) {

	switch {
	case ygoarCBO4et == 0:
		return ygoarCBO4et, 0
	case  /*line :1*/ ME1vzpD5wcr(ygoarCBO4et, 0) ||  /*line :1*/ OIdLpDqq(ygoarCBO4et):
		return ygoarCBO4et, 0
	}
	ygoarCBO4et, nF4u1Vnrit =  /*line :1*/ ckXskOBp(ygoarCBO4et)
	_BqkSz0 :=  /*line :1*/ GKIRmoHE(ygoarCBO4et)
	nF4u1Vnrit +=  /*line :1*/ int((_BqkSz0>>shift)&mask) - bias + 1
	_BqkSz0 &^= mask << shift
	_BqkSz0 |= (-1 + bias) << shift
	cmLS3KfTv =  /*line :1*/ QUB2Kf63p(_BqkSz0)
	return
}
