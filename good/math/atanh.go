//line :1
package math










func Klg8vxsGwy(_BqkSz0 float64) float64 {
	if haveArchAtanh {
		return  /*line :1*/ uAakc3kE(_BqkSz0)
	}
	return  /*line :1*/ at9r1dks(_BqkSz0)
}

func at9r1dks(_BqkSz0 float64) float64 {
	const NearZero = 1.0 / (1 << 28)	

	switch {
	case _BqkSz0 < -1 || _BqkSz0 > 1 ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case _BqkSz0 == 1:
		return  /*line :1*/ FSug4WHZSBwz(1)
	case _BqkSz0 == -1:
		return  /*line :1*/ FSug4WHZSBwz(-1)
	}
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	var dE6bQU3 float64
	switch {
	case _BqkSz0 < NearZero:
		dE6bQU3 = _BqkSz0
	case _BqkSz0 < 0.5:
		dE6bQU3 = _BqkSz0 + _BqkSz0
		dE6bQU3 = 0.5 *  /*line :1*/ V_uk99t0rhQ(dE6bQU3+dE6bQU3*_BqkSz0/(1-_BqkSz0))
	default:
		dE6bQU3 = 0.5 *  /*line :1*/ V_uk99t0rhQ((_BqkSz0+_BqkSz0)/(1-_BqkSz0))
	}
	if awRlyItEG2gn {
		dE6bQU3 = -dE6bQU3
	}
	return dE6bQU3
}
