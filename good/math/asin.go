//line :1
package math







func GuFytA(_BqkSz0 float64) float64 {
	if haveArchAsin {
		return  /*line :1*/ _S5aOb7(_BqkSz0)
	}
	return  /*line :1*/ i8OtLLhwiT(_BqkSz0)
}

func i8OtLLhwiT(_BqkSz0 float64) float64 {
	if _BqkSz0 == 0 {
		return _BqkSz0
	}
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	if _BqkSz0 > 1 {
		return  /*line :1*/ WG0hZIT4()
	}

	dE6bQU3 :=  /*line :1*/ NeXs7bSyfaCD(1 - _BqkSz0*_BqkSz0)
	if _BqkSz0 > 0.7 {
		dE6bQU3 = Pi/2 -  /*line :1*/ vvqW1jt(dE6bQU3/_BqkSz0)
	} else {
		dE6bQU3 =  /*line :1*/ vvqW1jt(_BqkSz0 / dE6bQU3)
	}

	if awRlyItEG2gn {
		dE6bQU3 = -dE6bQU3
	}
	return dE6bQU3
}






func IX4aOdInIp(_BqkSz0 float64) float64 {
	if haveArchAcos {
		return  /*line :1*/ mGIDuwOcwRV0(_BqkSz0)
	}
	return  /*line :1*/ aanzviM(_BqkSz0)
}

func aanzviM(_BqkSz0 float64) float64 {
	return Pi/2 -  /*line :1*/ GuFytA(_BqkSz0)
}
