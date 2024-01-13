//line :1
package math








func Floor(_BqkSz0 float64) float64 {
	if haveArchFloor {
		return  /*line :1*/ rKDAZG8(_BqkSz0)
	}
	return  /*line :1*/ ifKKkj(_BqkSz0)
}

func ifKKkj(_BqkSz0 float64) float64 {
	if _BqkSz0 == 0 ||  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0) {
		return _BqkSz0
	}
	if _BqkSz0 < 0 {
		pjjMpWihfSot, aYH3mKqmg :=  /*line :1*/ HJa8famXXqZ(-_BqkSz0)
		if aYH3mKqmg != 0.0 {
			pjjMpWihfSot = pjjMpWihfSot + 1
		}
		return -pjjMpWihfSot
	}
	pjjMpWihfSot, _ :=  /*line :1*/ HJa8famXXqZ(_BqkSz0)
	return pjjMpWihfSot
}








func Ceil(_BqkSz0 float64) float64 {
	if haveArchCeil {
		return  /*line :1*/ mwYxlolYRy(_BqkSz0)
	}
	return  /*line :1*/ fgDuenq(_BqkSz0)
}

func fgDuenq(_BqkSz0 float64) float64 {
	return - /*line :1*/ Floor(-_BqkSz0)
}








func Trunc(_BqkSz0 float64) float64 {
	if haveArchTrunc {
		return  /*line :1*/ y8qAmjse(_BqkSz0)
	}
	return  /*line :1*/ eyevtM(_BqkSz0)
}

func eyevtM(_BqkSz0 float64) float64 {
	if _BqkSz0 == 0 ||  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0) {
		return _BqkSz0
	}
	pjjMpWihfSot, _ :=  /*line :1*/ HJa8famXXqZ(_BqkSz0)
	return pjjMpWihfSot
}








func Round(_BqkSz0 float64) float64 {

	vhY6bOa_vL :=  /*line :1*/ GKIRmoHE(_BqkSz0)
	gwBaYgYgA :=  /*line :1*/ uint(vhY6bOa_vL>>shift) & mask
	if gwBaYgYgA < bias {

		vhY6bOa_vL &= signMask
		if gwBaYgYgA == bias-1 {
			vhY6bOa_vL |= uvone
		}
	} else if gwBaYgYgA < bias+shift {
		
		
		
		
		const half = 1 << (shift - 1)
		gwBaYgYgA -= bias
		vhY6bOa_vL += half >> gwBaYgYgA
		vhY6bOa_vL &^= fracMask >> gwBaYgYgA
	}
	return  /*line :1*/ QUB2Kf63p(vhY6bOa_vL)
}








func RoundToEven(_BqkSz0 float64) float64 {

	vhY6bOa_vL :=  /*line :1*/ GKIRmoHE(_BqkSz0)
	gwBaYgYgA :=  /*line :1*/ uint(vhY6bOa_vL>>shift) & mask
	if gwBaYgYgA >= bias {
		
		
		
		
		const halfMinusULP = (1 << (shift - 1)) - 1
		gwBaYgYgA -= bias
		vhY6bOa_vL += (halfMinusULP + (vhY6bOa_vL>>(shift-gwBaYgYgA))&1) >> gwBaYgYgA
		vhY6bOa_vL &^= fracMask >> gwBaYgYgA
	} else if gwBaYgYgA == bias-1 && vhY6bOa_vL&fracMask != 0 {

		vhY6bOa_vL = vhY6bOa_vL&signMask | uvone
	} else {

		vhY6bOa_vL &= signMask
	}
	return  /*line :1*/ QUB2Kf63p(vhY6bOa_vL)
}
