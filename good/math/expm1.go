//line :1
package math











func U4jtBgdPk(_BqkSz0 float64) float64 {
	if haveArchExpm1 {
		return  /*line :1*/ gVw7eJ(_BqkSz0)
	}
	return  /*line :1*/ dIz8EVGYlAJ(_BqkSz0)
}

func dIz8EVGYlAJ(_BqkSz0 float64) float64 {
	const (
		Othreshold	= 7.09782712893383973096e+02		
		Ln2X56	= 3.88162421113569373274e+01		
		Ln2HalfX3	= 1.03972077083991796413e+00		
		Ln2Half	= 3.46573590279972654709e-01		
		Ln2Hi	= 6.93147180369123816490e-01		
		Ln2Lo	= 1.90821492927058770002e-10		
		InvLn2	= 1.44269504088896338700e+00		
		Tiny	= 1.0 / (1 << 54)		
		
		Q1	= -3.33333333333331316428e-02		
		Q2	= 1.58730158725481460165e-03		
		Q3	= -7.93650757867487942473e-05		
		Q4	= 4.00821782732936239552e-06		
		Q5	= -2.01099218183624371326e-07		
	)

	switch {
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1) ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return _BqkSz0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, -1):
		return -1
	}

	dU7v_ec8cB := _BqkSz0
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		dU7v_ec8cB = -dU7v_ec8cB
		awRlyItEG2gn = true
	}

	if dU7v_ec8cB >= Ln2X56 {
		if awRlyItEG2gn {
			return -1
		}
		if dU7v_ec8cB >= Othreshold {
			return  /*line :1*/ FSug4WHZSBwz(1)
		}
	}

	
	var cU6qMJer46oH float64
	var duDPpte int
	if dU7v_ec8cB > Ln2Half {
		var j2BkoXTMn9D, j0RLYji float64
		if dU7v_ec8cB < Ln2HalfX3 {
			if !awRlyItEG2gn {
				j2BkoXTMn9D = _BqkSz0 - Ln2Hi
				j0RLYji = Ln2Lo
				duDPpte = 1
			} else {
				j2BkoXTMn9D = _BqkSz0 + Ln2Hi
				j0RLYji = -Ln2Lo
				duDPpte = -1
			}
		} else {
			if !awRlyItEG2gn {
				duDPpte =  /*line :1*/ int(InvLn2*_BqkSz0 + 0.5)
			} else {
				duDPpte =  /*line :1*/ int(InvLn2*_BqkSz0 - 0.5)
			}
			aX6xkg_1G :=  /*line :1*/ float64(duDPpte)
			j2BkoXTMn9D = _BqkSz0 - aX6xkg_1G*Ln2Hi
			j0RLYji = aX6xkg_1G * Ln2Lo
		}
		_BqkSz0 = j2BkoXTMn9D - j0RLYji
		cU6qMJer46oH = (j2BkoXTMn9D - _BqkSz0) - j0RLYji
	} else if dU7v_ec8cB < Tiny {
		return _BqkSz0
	} else {
		duDPpte = 0
	}

	xdtQNByfME := 0.5 * _BqkSz0
	f8aIcM0ijCV6 := _BqkSz0 * xdtQNByfME
	yYCutiG6s3 := 1 + f8aIcM0ijCV6*(Q1+f8aIcM0ijCV6*(Q2+f8aIcM0ijCV6*(Q3+f8aIcM0ijCV6*(Q4+f8aIcM0ijCV6*Q5))))
	aX6xkg_1G := 3 - yYCutiG6s3*xdtQNByfME
	gwBaYgYgA := f8aIcM0ijCV6 * ((yYCutiG6s3 - aX6xkg_1G) / (6.0 - _BqkSz0*aX6xkg_1G))
	if duDPpte == 0 {
		return _BqkSz0 - (_BqkSz0*gwBaYgYgA - f8aIcM0ijCV6)
	}
	gwBaYgYgA = (_BqkSz0*(gwBaYgYgA-cU6qMJer46oH) - cU6qMJer46oH)
	gwBaYgYgA -= f8aIcM0ijCV6
	switch {
	case duDPpte == -1:
		return 0.5*(_BqkSz0-gwBaYgYgA) - 0.5
	case duDPpte == 1:
		if _BqkSz0 < -0.25 {
			return -2 * (gwBaYgYgA - (_BqkSz0 + 0.5))
		}
		return 1 + 2*(_BqkSz0-gwBaYgYgA)
	case duDPpte <= -2 || duDPpte > 56:
		nNeKvlvK6 := 1 - (gwBaYgYgA - _BqkSz0)
		nNeKvlvK6 =  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(nNeKvlvK6) +  /*line :1*/ uint64(duDPpte)<<52)
		return nNeKvlvK6 - 1
	}
	if duDPpte < 20 {
		aX6xkg_1G :=  /*line :1*/ QUB2Kf63p(0x3ff0000000000000 - (0x20000000000000 >>  /*line :1*/ uint(duDPpte)))
		nNeKvlvK6 := aX6xkg_1G - (gwBaYgYgA - _BqkSz0)
		nNeKvlvK6 =  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(nNeKvlvK6) +  /*line :1*/ uint64(duDPpte)<<52)
		return nNeKvlvK6
	}
	aX6xkg_1G =  /*line :1*/ QUB2Kf63p( /*line :1*/ uint64(0x3ff-duDPpte) << 52)
	nNeKvlvK6 := _BqkSz0 - (gwBaYgYgA + aX6xkg_1G)
	nNeKvlvK6++
	nNeKvlvK6 =  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(nNeKvlvK6) +  /*line :1*/ uint64(duDPpte)<<52)
	return nNeKvlvK6
}
