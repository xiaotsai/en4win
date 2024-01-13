//line :1
package math








func ZQZCHD(_BqkSz0 float64) float64 {

	switch {
	case _BqkSz0 == 0:
		return  /*line :1*/ FSug4WHZSBwz(-1)
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return  /*line :1*/ FSug4WHZSBwz(1)
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		return _BqkSz0
	}
	return  /*line :1*/ float64( /*line :1*/ eKZjePZd(_BqkSz0))
}








func YFzHiLgE6aG_(_BqkSz0 float64) int {

	switch {
	case _BqkSz0 == 0:
		return MinInt32
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		return MaxInt32
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return MaxInt32
	}
	return  /*line :1*/ eKZjePZd(_BqkSz0)
}



func eKZjePZd(_BqkSz0 float64) int {
	_BqkSz0, nF4u1Vnrit :=  /*line :1*/ ckXskOBp(_BqkSz0)
	return  /*line :1*/ int(( /*line :1*/ GKIRmoHE(_BqkSz0)>>shift)&mask) - bias + nF4u1Vnrit
}
