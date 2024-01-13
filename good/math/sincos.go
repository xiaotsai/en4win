//line :1
package math








func BcjJnykwJ(_BqkSz0 float64) (iWDxKcqf, vXNMHuX9eF5y float64) {
	const (
		PI4A	= 7.85398125648498535156e-1		
		PI4B	= 3.77489470793079817668e-8		
		PI4C	= 2.69515142907905952645e-15		
	)

	switch {
	case _BqkSz0 == 0:
		return _BqkSz0, 1
	case  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return  /*line :1*/ WG0hZIT4(),  /*line :1*/ WG0hZIT4()
	}

	tkpaUVI2mANa, flJ3EFee8 := false, false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		tkpaUVI2mANa = true
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
		vnPtu1up3 &= 7
		gZDpjWvXax_q = ((_BqkSz0 - nNeKvlvK6*PI4A) - nNeKvlvK6*PI4B) - nNeKvlvK6*PI4C
	}
	if vnPtu1up3 > 3 {
		vnPtu1up3 -= 4
		tkpaUVI2mANa, flJ3EFee8 = !tkpaUVI2mANa, !flJ3EFee8
	}
	if vnPtu1up3 > 1 {
		flJ3EFee8 = !flJ3EFee8
	}

	zvhewoX := gZDpjWvXax_q * gZDpjWvXax_q
	vXNMHuX9eF5y = 1.0 - 0.5*zvhewoX + zvhewoX*zvhewoX*((((((dQf62tKx[0]*zvhewoX)+dQf62tKx[1])*zvhewoX+dQf62tKx[2])*zvhewoX+dQf62tKx[3])*zvhewoX+dQf62tKx[4])*zvhewoX+dQf62tKx[5])
	iWDxKcqf = gZDpjWvXax_q + gZDpjWvXax_q*zvhewoX*((((((uM0vLKSX7Xt9[0]*zvhewoX)+uM0vLKSX7Xt9[1])*zvhewoX+uM0vLKSX7Xt9[2])*zvhewoX+uM0vLKSX7Xt9[3])*zvhewoX+uM0vLKSX7Xt9[4])*zvhewoX+uM0vLKSX7Xt9[5])
	if vnPtu1up3 == 1 || vnPtu1up3 == 2 {
		iWDxKcqf, vXNMHuX9eF5y = vXNMHuX9eF5y, iWDxKcqf
	}
	if flJ3EFee8 {
		vXNMHuX9eF5y = -vXNMHuX9eF5y
	}
	if tkpaUVI2mANa {
		iWDxKcqf = -iWDxKcqf
	}
	return
}
