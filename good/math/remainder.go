//line :1
package math










func Wg0amH3(_BqkSz0, nNeKvlvK6 float64) float64 {
	if haveArchRemainder {
		return  /*line :1*/ meqVVFJyUlQ9(_BqkSz0, nNeKvlvK6)
	}
	return  /*line :1*/ r3BEvz(_BqkSz0, nNeKvlvK6)
}

func r3BEvz(_BqkSz0, nNeKvlvK6 float64) float64 {
	const (
		Tiny	= 4.45014771701440276618e-308		
		HalfMax	= MaxFloat64 / 2
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ OIdLpDqq(nNeKvlvK6) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0) || nNeKvlvK6 == 0:
		return  /*line :1*/ WG0hZIT4()
	case  /*line :1*/ ME1vzpD5wcr(nNeKvlvK6, 0):
		return _BqkSz0
	}
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	if nNeKvlvK6 < 0 {
		nNeKvlvK6 = -nNeKvlvK6
	}
	if _BqkSz0 == nNeKvlvK6 {
		if awRlyItEG2gn {
			q7GRsm4g799 := 0.0
			return -q7GRsm4g799
		}
		return 0
	}
	if nNeKvlvK6 <= HalfMax {
		_BqkSz0 =  /*line :1*/ LoDnahP_eoog(_BqkSz0, nNeKvlvK6+nNeKvlvK6)
	}
	if nNeKvlvK6 < Tiny {
		if _BqkSz0+_BqkSz0 > nNeKvlvK6 {
			_BqkSz0 -= nNeKvlvK6
			if _BqkSz0+_BqkSz0 >= nNeKvlvK6 {
				_BqkSz0 -= nNeKvlvK6
			}
		}
	} else {
		iBvFE_v := 0.5 * nNeKvlvK6
		if _BqkSz0 > iBvFE_v {
			_BqkSz0 -= nNeKvlvK6
			if _BqkSz0 >= iBvFE_v {
				_BqkSz0 -= nNeKvlvK6
			}
		}
	}
	if awRlyItEG2gn {
		_BqkSz0 = -_BqkSz0
	}
	return _BqkSz0
}
