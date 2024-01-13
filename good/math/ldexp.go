//line :1
package math









func EU1SFL1cw(cmLS3KfTv float64, nF4u1Vnrit int) float64 {
	if haveArchLdexp {
		return  /*line :1*/ yCrE8ReJQ(cmLS3KfTv, nF4u1Vnrit)
	}
	return  /*line :1*/ yFwapcGBgSvf(cmLS3KfTv, nF4u1Vnrit)
}

func yFwapcGBgSvf(cmLS3KfTv float64, nF4u1Vnrit int) float64 {

	switch {
	case cmLS3KfTv == 0:
		return cmLS3KfTv
	case  /*line :1*/ ME1vzpD5wcr(cmLS3KfTv, 0) ||  /*line :1*/ OIdLpDqq(cmLS3KfTv):
		return cmLS3KfTv
	}
	cmLS3KfTv, gwBaYgYgA :=  /*line :1*/ ckXskOBp(cmLS3KfTv)
	nF4u1Vnrit += gwBaYgYgA
	_BqkSz0 :=  /*line :1*/ GKIRmoHE(cmLS3KfTv)
	nF4u1Vnrit +=  /*line :1*/ int(_BqkSz0>>shift)&mask - bias
	if nF4u1Vnrit < -1075 {
		return  /*line :1*/ Copysign(0, cmLS3KfTv)
	}
	if nF4u1Vnrit > 1023 {
		if cmLS3KfTv < 0 {
			return  /*line :1*/ FSug4WHZSBwz(-1)
		}
		return  /*line :1*/ FSug4WHZSBwz(1)
	}
	var b8CRqLSt float64 = 1
	if nF4u1Vnrit < -1022 {
		nF4u1Vnrit += 53
		b8CRqLSt = 1.0 / (1 << 53)
	}
	_BqkSz0 &^= mask << shift
	_BqkSz0 |=  /*line :1*/ uint64(nF4u1Vnrit+bias) << shift
	return b8CRqLSt *  /*line :1*/ QUB2Kf63p(_BqkSz0)
}
