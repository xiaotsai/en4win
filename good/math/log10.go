//line :1
package math



func PRCRFnW4_8(_BqkSz0 float64) float64 {
	if haveArchLog10 {
		return  /*line :1*/ mSXtrRfbQdmT(_BqkSz0)
	}
	return  /*line :1*/ lzlgMC8Hz(_BqkSz0)
}

func lzlgMC8Hz(_BqkSz0 float64) float64 {
	return  /*line :1*/ JrJVPb5WbG(_BqkSz0) * (1 / Ln10)
}



func JwAjQi(_BqkSz0 float64) float64 {
	if haveArchLog2 {
		return  /*line :1*/ fbNAFcvGO(_BqkSz0)
	}
	return  /*line :1*/ g8AGhVhiq(_BqkSz0)
}

func g8AGhVhiq(_BqkSz0 float64) float64 {
	cmLS3KfTv, nF4u1Vnrit :=  /*line :1*/ CuUl2eV(_BqkSz0)

	if cmLS3KfTv == 0.5 {
		return  /*line :1*/ float64(nF4u1Vnrit - 1)
	}
	return  /*line :1*/ JrJVPb5WbG(cmLS3KfTv)*(1/Ln2) +  /*line :1*/ float64(nF4u1Vnrit)
}
