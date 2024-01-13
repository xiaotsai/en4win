//line :1
package math








func F9Do_Ox(_BqkSz0 float64) float64 {
	if haveArchCbrt {
		return  /*line :1*/ bpUInQ2(_BqkSz0)
	}
	return  /*line :1*/ skdwD4wV(_BqkSz0)
}

func skdwD4wV(_BqkSz0 float64) float64 {
	const (
		B1	= 715094163		
		B2	= 696219795		
		C	= 5.42857142857142815906e-01		
		D	= -7.05306122448979611050e-01		
		E	= 1.41428571428571436819e+00		
		F	= 1.60714285714285720630e+00		
		G	= 3.57142857142857150787e-01		
		SmallestNormal	= 2.22507385850720138309e-308		
	)

	switch {
	case _BqkSz0 == 0 ||  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return _BqkSz0
	}

	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}

	aX6xkg_1G :=  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(_BqkSz0)/3 + B1<<32)
	if _BqkSz0 < SmallestNormal {

		aX6xkg_1G =  /*line :1*/ float64(1 << 54)
		aX6xkg_1G *= _BqkSz0
		aX6xkg_1G =  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(aX6xkg_1G)/3 + B2<<32)
	}

	aBK_X3rbPc := aX6xkg_1G * aX6xkg_1G / _BqkSz0
	aK8LO1 := C + aBK_X3rbPc*aX6xkg_1G
	aX6xkg_1G *= G + F/(aK8LO1+E+D/aK8LO1)

	aX6xkg_1G =  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(aX6xkg_1G)&(0xFFFFFFFFC<<28) + 1<<30)

	aK8LO1 = aX6xkg_1G * aX6xkg_1G
	aBK_X3rbPc = _BqkSz0 / aK8LO1
	uuSWcjr := aX6xkg_1G + aX6xkg_1G
	aBK_X3rbPc = (aBK_X3rbPc - aX6xkg_1G) / (uuSWcjr + aBK_X3rbPc)
	aX6xkg_1G = aX6xkg_1G + aX6xkg_1G*aBK_X3rbPc

	if awRlyItEG2gn {
		aX6xkg_1G = -aX6xkg_1G
	}
	return aX6xkg_1G
}
