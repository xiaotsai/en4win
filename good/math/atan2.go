//line :1
package math
























func EgRuD4zewvO(nNeKvlvK6, _BqkSz0 float64) float64 {
	if haveArchAtan2 {
		return  /*line :1*/ lSyji_oaAQ(nNeKvlvK6, _BqkSz0)
	}
	return  /*line :1*/ o0SxywVqsNC(nNeKvlvK6, _BqkSz0)
}

func o0SxywVqsNC(nNeKvlvK6, _BqkSz0 float64) float64 {

	switch {
	case  /*line :1*/ OIdLpDqq(nNeKvlvK6) ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case nNeKvlvK6 == 0:
		if _BqkSz0 >= 0 && ! /*line :1*/ GLZFJA9L9(_BqkSz0) {
			return  /*line :1*/ Copysign(0, nNeKvlvK6)
		}
		return  /*line :1*/ Copysign(Pi, nNeKvlvK6)
	case _BqkSz0 == 0:
		return  /*line :1*/ Copysign(Pi/2, nNeKvlvK6)
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		if  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1) {
			switch {
			case  /*line :1*/ ME1vzpD5wcr(nNeKvlvK6, 0):
				return  /*line :1*/ Copysign(Pi/4, nNeKvlvK6)
			default:
				return  /*line :1*/ Copysign(0, nNeKvlvK6)
			}
		}
		switch {
		case  /*line :1*/ ME1vzpD5wcr(nNeKvlvK6, 0):
			return  /*line :1*/ Copysign(3*Pi/4, nNeKvlvK6)
		default:
			return  /*line :1*/ Copysign(Pi, nNeKvlvK6)
		}
	case  /*line :1*/ ME1vzpD5wcr(nNeKvlvK6, 0):
		return  /*line :1*/ Copysign(Pi/2, nNeKvlvK6)
	}

	gNDglc3F2QV :=  /*line :1*/ DaNPs9X(nNeKvlvK6 / _BqkSz0)
	if _BqkSz0 < 0 {
		if gNDglc3F2QV <= 0 {
			return gNDglc3F2QV + Pi
		}
		return gNDglc3F2QV - Pi
	}
	return gNDglc3F2QV
}
