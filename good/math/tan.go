//line :1
package math


var mn6YyXjuIV70 = [...]float64{
	-1.30936939181383777646e4,
	1.15351664838587416140e6,
	-1.79565251976484877988e7,
}
var s6QSkPPQKB_ = [...]float64{
	1.00000000000000000000e0,
	1.36812963470692954678e4,
	-1.32089234440210967447e6,
	2.50083801823357915839e7,
	-5.38695755929454629881e7,
}








func HBKp9wgTNl7(_BqkSz0 float64) float64 {
	if haveArchTan {
		return  /*line :1*/ l6dM2yVTcj_(_BqkSz0)
	}
	return  /*line :1*/ eAeA_aAKlPa(_BqkSz0)
}

func eAeA_aAKlPa(_BqkSz0 float64) float64 {
	const (
		PI4A	= 7.85398125648498535156e-1		
		PI4B	= 3.77489470793079817668e-8		
		PI4C	= 2.69515142907905952645e-15		
	)

	switch {
	case _BqkSz0 == 0 ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return _BqkSz0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return  /*line :1*/ WG0hZIT4()
	}

	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	var vnPtu1up3 uint64
	var nNeKvlvK6, gZDpjWvXax_q float64
	if _BqkSz0 >= reduceThreshold {
		vnPtu1up3, gZDpjWvXax_q =  /*line :1*/ noJmE7EjH3k(_BqkSz0)
	} else {
		vnPtu1up3 =  /*line :1*/ uint64(_BqkSz0 * (4 / Pi))
		nNeKvlvK6 =  /*line :1*/ float64(vnPtu1up3)

		if vnPtu1up3&1 == 1 {
			vnPtu1up3++
			nNeKvlvK6++
		}

		gZDpjWvXax_q = ((_BqkSz0 - nNeKvlvK6*PI4A) - nNeKvlvK6*PI4B) - nNeKvlvK6*PI4C
	}
	zvhewoX := gZDpjWvXax_q * gZDpjWvXax_q

	if zvhewoX > 1e-14 {
		nNeKvlvK6 = gZDpjWvXax_q + gZDpjWvXax_q*(zvhewoX*(((mn6YyXjuIV70[0]*zvhewoX)+mn6YyXjuIV70[1])*zvhewoX+mn6YyXjuIV70[2])/((((zvhewoX+s6QSkPPQKB_[1])*zvhewoX+s6QSkPPQKB_[2])*zvhewoX+s6QSkPPQKB_[3])*zvhewoX+s6QSkPPQKB_[4]))
	} else {
		nNeKvlvK6 = gZDpjWvXax_q
	}
	if vnPtu1up3&2 == 2 {
		nNeKvlvK6 = -1 / nNeKvlvK6
	}
	if awRlyItEG2gn {
		nNeKvlvK6 = -nNeKvlvK6
	}
	return nNeKvlvK6
}
