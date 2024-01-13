//line :1
package math








func DYMqjo(_BqkSz0 float64) float64 {
	if haveArchSinh {
		return  /*line :1*/ lDGwmjo2f0Dn(_BqkSz0)
	}
	return  /*line :1*/ n0b_9K(_BqkSz0)
}

func n0b_9K(_BqkSz0 float64) float64 {
	
	const (
		P0	= -0.6307673640497716991184787251e+6
		P1	= -0.8991272022039509355398013511e+5
		P2	= -0.2894211355989563807284660366e+4
		P3	= -0.2630563213397497062819489e+2
		Q0	= -0.6307673640497716991212077277e+6
		Q1	= 0.1521517378790019070696485176e+5
		Q2	= -0.173678953558233699533450911e+3
	)

	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}

	var dE6bQU3 float64
	switch {
	case _BqkSz0 > 21:
		dE6bQU3 =  /*line :1*/ JPvpMNJa(_BqkSz0) * 0.5

	case _BqkSz0 > 0.5:
		wVMD5ORsN :=  /*line :1*/ JPvpMNJa(_BqkSz0)
		dE6bQU3 = (wVMD5ORsN - 1/wVMD5ORsN) * 0.5

	default:
		w3CcYaPx := _BqkSz0 * _BqkSz0
		dE6bQU3 = (((P3*w3CcYaPx+P2)*w3CcYaPx+P1)*w3CcYaPx + P0) * _BqkSz0
		dE6bQU3 = dE6bQU3 / (((w3CcYaPx+Q2)*w3CcYaPx+Q1)*w3CcYaPx + Q0)
	}

	if awRlyItEG2gn {
		dE6bQU3 = -dE6bQU3
	}
	return dE6bQU3
}








func LJfEkO9jWHp(_BqkSz0 float64) float64 {
	if haveArchCosh {
		return  /*line :1*/ tZxuxQrUn(_BqkSz0)
	}
	return  /*line :1*/ wxurRbXRsWt0(_BqkSz0)
}

func wxurRbXRsWt0(_BqkSz0 float64) float64 {
	_BqkSz0 =  /*line :1*/ Abs(_BqkSz0)
	if _BqkSz0 > 21 {
		return  /*line :1*/ JPvpMNJa(_BqkSz0) * 0.5
	}
	wVMD5ORsN :=  /*line :1*/ JPvpMNJa(_BqkSz0)
	return (wVMD5ORsN + 1/wVMD5ORsN) * 0.5
}
