//line :1
package math

func nzjHRDU(_BqkSz0 float64) bool {
	if  /*line :1*/ Abs(_BqkSz0) >= (1 << 53) {

		return false
	}

	u9bKBw, wu_h2B5aU5A :=  /*line :1*/ HJa8famXXqZ(_BqkSz0)
	return wu_h2B5aU5A == 0 &&  /*line :1*/ int64(u9bKBw)&1 == 1
}

























func CPKyEBb3t3ps(_BqkSz0, nNeKvlvK6 float64) float64 {
	if haveArchPow {
		return  /*line :1*/ mvcdT0DQZt(_BqkSz0, nNeKvlvK6)
	}
	return  /*line :1*/ tkYf9b7p5D(_BqkSz0, nNeKvlvK6)
}

func tkYf9b7p5D(_BqkSz0, nNeKvlvK6 float64) float64 {
	switch {
	case nNeKvlvK6 == 0 || _BqkSz0 == 1:
		return 1
	case nNeKvlvK6 == 1:
		return _BqkSz0
	case  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ OIdLpDqq(nNeKvlvK6):
		return  /*line :1*/ WG0hZIT4()
	case _BqkSz0 == 0:
		switch {
		case nNeKvlvK6 < 0:
			if  /*line :1*/ GLZFJA9L9(_BqkSz0) &&  /*line :1*/ nzjHRDU(nNeKvlvK6) {
				return  /*line :1*/ FSug4WHZSBwz(-1)
			}
			return  /*line :1*/ FSug4WHZSBwz(1)
		case nNeKvlvK6 > 0:
			if  /*line :1*/ GLZFJA9L9(_BqkSz0) &&  /*line :1*/ nzjHRDU(nNeKvlvK6) {
				return _BqkSz0
			}
			return 0
		}
	case  /*line :1*/ ME1vzpD5wcr(nNeKvlvK6, 0):
		switch {
		case _BqkSz0 == -1:
			return 1
		case ( /*line :1*/ Abs(_BqkSz0) < 1) ==  /*line :1*/ ME1vzpD5wcr(nNeKvlvK6, 1):
			return 0
		default:
			return  /*line :1*/ FSug4WHZSBwz(1)
		}
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		if  /*line :1*/ ME1vzpD5wcr(_BqkSz0, -1) {
			return  /*line :1*/ CPKyEBb3t3ps(1/_BqkSz0, -nNeKvlvK6)
		}
		switch {
		case nNeKvlvK6 < 0:
			return 0
		case nNeKvlvK6 > 0:
			return  /*line :1*/ FSug4WHZSBwz(1)
		}
	case nNeKvlvK6 == 0.5:
		return  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
	case nNeKvlvK6 == -0.5:
		return 1 /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
	}

	rvil4ZaGVF, n8NSWU :=  /*line :1*/ HJa8famXXqZ( /*line :1*/ Abs(nNeKvlvK6))
	if n8NSWU != 0 && _BqkSz0 < 0 {
		return  /*line :1*/ WG0hZIT4()
	}
	if rvil4ZaGVF >= 1<<63 {

		switch {
		case _BqkSz0 == -1:
			return 1
		case ( /*line :1*/ Abs(_BqkSz0) < 1) == (nNeKvlvK6 > 0):
			return 0
		default:
			return  /*line :1*/ FSug4WHZSBwz(1)
		}
	}

	iC6vamfy4 := 1.0
	xuzp4m := 0

	if n8NSWU != 0 {
		if n8NSWU > 0.5 {
			n8NSWU--
			rvil4ZaGVF++
		}
		iC6vamfy4 =  /*line :1*/ JPvpMNJa(n8NSWU *  /*line :1*/ JrJVPb5WbG(_BqkSz0))
	}

	wqUhnFNgLXTM, lBwiw8EWW :=  /*line :1*/ CuUl2eV(_BqkSz0)
	for jJklNl_1r :=  /*line :1*/ int64(rvil4ZaGVF); jJklNl_1r != 0; jJklNl_1r >>= 1 {
		if lBwiw8EWW < -1<<12 || 1<<12 < lBwiw8EWW {

			xuzp4m += lBwiw8EWW
			break
		}
		if jJklNl_1r&1 == 1 {
			iC6vamfy4 *= wqUhnFNgLXTM
			xuzp4m += lBwiw8EWW
		}
		wqUhnFNgLXTM *= wqUhnFNgLXTM
		lBwiw8EWW <<= 1
		if wqUhnFNgLXTM < .5 {
			wqUhnFNgLXTM += wqUhnFNgLXTM
			lBwiw8EWW--
		}
	}

	if nNeKvlvK6 < 0 {
		iC6vamfy4 = 1 / iC6vamfy4
		xuzp4m = -xuzp4m
	}
	return  /*line :1*/ EU1SFL1cw(iC6vamfy4, xuzp4m)
}
