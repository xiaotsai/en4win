//line :1
package math








func CXTbgGG_Ox(_BqkSz0 float64) float64 {
	if haveArchAsinh {
		return  /*line :1*/ rHEGaC(_BqkSz0)
	}
	return  /*line :1*/ ww9qrm0s(_BqkSz0)
}

func ww9qrm0s(_BqkSz0 float64) float64 {
	const (
		Ln2	= 6.93147180559945286227e-01		
		NearZero	= 1.0 / (1 << 28)		
		Large	= 1 << 28		
	)

	if  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0) {
		return _BqkSz0
	}
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	var dE6bQU3 float64
	switch {
	case _BqkSz0 > Large:
		dE6bQU3 =  /*line :1*/ JrJVPb5WbG(_BqkSz0) + Ln2
	case _BqkSz0 > 2:
		dE6bQU3 =  /*line :1*/ JrJVPb5WbG(2*_BqkSz0 + 1/( /*line :1*/ NeXs7bSyfaCD(_BqkSz0*_BqkSz0+1)+_BqkSz0))
	case _BqkSz0 < NearZero:
		dE6bQU3 = _BqkSz0
	default:
		dE6bQU3 =  /*line :1*/ V_uk99t0rhQ(_BqkSz0 + _BqkSz0*_BqkSz0/(1+ /*line :1*/ NeXs7bSyfaCD(1+_BqkSz0*_BqkSz0)))
	}
	if awRlyItEG2gn {
		dE6bQU3 = -dE6bQU3
	}
	return dE6bQU3
}
