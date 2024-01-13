//line :1
package math


var uM0vLKSX7Xt9 = [...]float64{
	1.58962301576546568060e-10,
	-2.50507477628578072866e-8,
	2.75573136213857245213e-6,
	-1.98412698295895385996e-4,
	8.33333333332211858878e-3,
	-1.66666666666666307295e-1,
}


var dQf62tKx = [...]float64{
	-1.13585365213876817300e-11,
	2.08757008419747316778e-9,
	-2.75573141792967388112e-7,
	2.48015872888517045348e-5,
	-1.38888888888730564116e-3,
	4.16666666666665929218e-2,
}







func Sr_3cKxech(_BqkSz0 float64) float64 {
	if haveArchCos {
		return  /*line :1*/ mEN2ABta(_BqkSz0)
	}
	return  /*line :1*/ vXNMHuX9eF5y(_BqkSz0)
}

func vXNMHuX9eF5y(_BqkSz0 float64) float64 {
	const (
		PI4A	= 7.85398125648498535156e-1		
		PI4B	= 3.77489470793079817668e-8		
		PI4C	= 2.69515142907905952645e-15		
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0) ||  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return  /*line :1*/ WG0hZIT4()
	}

	awRlyItEG2gn := false
	_BqkSz0 =  /*line :1*/ Abs(_BqkSz0)

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
		awRlyItEG2gn = !awRlyItEG2gn
	}
	if vnPtu1up3 > 1 {
		awRlyItEG2gn = !awRlyItEG2gn
	}

	zvhewoX := gZDpjWvXax_q * gZDpjWvXax_q
	if vnPtu1up3 == 1 || vnPtu1up3 == 2 {
		nNeKvlvK6 = gZDpjWvXax_q + gZDpjWvXax_q*zvhewoX*((((((uM0vLKSX7Xt9[0]*zvhewoX)+uM0vLKSX7Xt9[1])*zvhewoX+uM0vLKSX7Xt9[2])*zvhewoX+uM0vLKSX7Xt9[3])*zvhewoX+uM0vLKSX7Xt9[4])*zvhewoX+uM0vLKSX7Xt9[5])
	} else {
		nNeKvlvK6 = 1.0 - 0.5*zvhewoX + zvhewoX*zvhewoX*((((((dQf62tKx[0]*zvhewoX)+dQf62tKx[1])*zvhewoX+dQf62tKx[2])*zvhewoX+dQf62tKx[3])*zvhewoX+dQf62tKx[4])*zvhewoX+dQf62tKx[5])
	}
	if awRlyItEG2gn {
		nNeKvlvK6 = -nNeKvlvK6
	}
	return nNeKvlvK6
}








func Cyl8FsLg(_BqkSz0 float64) float64 {
	if haveArchSin {
		return  /*line :1*/ zaRUMG5YU(_BqkSz0)
	}
	return  /*line :1*/ iWDxKcqf(_BqkSz0)
}

func iWDxKcqf(_BqkSz0 float64) float64 {
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
		vnPtu1up3 &= 7
		gZDpjWvXax_q = ((_BqkSz0 - nNeKvlvK6*PI4A) - nNeKvlvK6*PI4B) - nNeKvlvK6*PI4C
	}

	if vnPtu1up3 > 3 {
		awRlyItEG2gn = !awRlyItEG2gn
		vnPtu1up3 -= 4
	}
	zvhewoX := gZDpjWvXax_q * gZDpjWvXax_q
	if vnPtu1up3 == 1 || vnPtu1up3 == 2 {
		nNeKvlvK6 = 1.0 - 0.5*zvhewoX + zvhewoX*zvhewoX*((((((dQf62tKx[0]*zvhewoX)+dQf62tKx[1])*zvhewoX+dQf62tKx[2])*zvhewoX+dQf62tKx[3])*zvhewoX+dQf62tKx[4])*zvhewoX+dQf62tKx[5])
	} else {
		nNeKvlvK6 = gZDpjWvXax_q + gZDpjWvXax_q*zvhewoX*((((((uM0vLKSX7Xt9[0]*zvhewoX)+uM0vLKSX7Xt9[1])*zvhewoX+uM0vLKSX7Xt9[2])*zvhewoX+uM0vLKSX7Xt9[3])*zvhewoX+uM0vLKSX7Xt9[4])*zvhewoX+uM0vLKSX7Xt9[5])
	}
	if awRlyItEG2gn {
		nNeKvlvK6 = -nNeKvlvK6
	}
	return nNeKvlvK6
}
