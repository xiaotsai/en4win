//line :1
package math







func QuAAZFEbIw(d2X2v9 int, _BqkSz0 float64) float64 {
	const (
		TwoM29	= 1.0 / (1 << 29)		
		Two302	= 1 << 302		
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		return _BqkSz0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return 0
	}

	if d2X2v9 == 0 {
		return  /*line :1*/ TZVYtdEXoPFm(_BqkSz0)
	}
	if _BqkSz0 == 0 {
		return 0
	}
	if d2X2v9 < 0 {
		d2X2v9, _BqkSz0 = -d2X2v9, -_BqkSz0
	}
	if d2X2v9 == 1 {
		return  /*line :1*/ KRpFivd(_BqkSz0)
	}
	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		if d2X2v9&1 == 1 {
			awRlyItEG2gn = true
		}
	}
	var alfthQvYkJ float64
	if  /*line :1*/ float64(d2X2v9) <= _BqkSz0 {

		if _BqkSz0 >= Two302 {

			var dE6bQU3 float64
			switch aK8LO1, cU6qMJer46oH :=  /*line :1*/ BcjJnykwJ(_BqkSz0); d2X2v9 & 3 {
			case 0:
				dE6bQU3 = cU6qMJer46oH + aK8LO1
			case 1:
				dE6bQU3 = -cU6qMJer46oH + aK8LO1
			case 2:
				dE6bQU3 = -cU6qMJer46oH - aK8LO1
			case 3:
				dE6bQU3 = cU6qMJer46oH - aK8LO1
			}
			alfthQvYkJ = (1 / SqrtPi) * dE6bQU3 /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		} else {
			alfthQvYkJ =  /*line :1*/ KRpFivd(_BqkSz0)
			for jJklNl_1r, a020cdSBAIO := 1,  /*line :1*/ TZVYtdEXoPFm(_BqkSz0); jJklNl_1r < d2X2v9; jJklNl_1r++ {
				a020cdSBAIO, alfthQvYkJ = alfthQvYkJ, alfthQvYkJ*( /*line :1*/ float64(jJklNl_1r+jJklNl_1r)/_BqkSz0)-a020cdSBAIO
			}
		}
	} else {
		if _BqkSz0 < TwoM29 {

			if d2X2v9 > 33 {
				alfthQvYkJ = 0
			} else {
				dE6bQU3 := _BqkSz0 * 0.5
				alfthQvYkJ = dE6bQU3
				a020cdSBAIO := 1.0
				for jJklNl_1r := 2; jJklNl_1r <= d2X2v9; jJklNl_1r++ {
					a020cdSBAIO *=  /*line :1*/ float64(jJklNl_1r)
					alfthQvYkJ *= dE6bQU3
				}
				alfthQvYkJ /= a020cdSBAIO
			}
		} else {

			uuSWcjr :=  /*line :1*/ float64(d2X2v9+d2X2v9) / _BqkSz0
			az8aTv8dcu := 2 / _BqkSz0
			zQ_5oM22Q := uuSWcjr
			gZDpjWvXax_q := uuSWcjr + az8aTv8dcu
			gouGLnAk := uuSWcjr*gZDpjWvXax_q - 1
			duDPpte := 1
			for gouGLnAk < 1e9 {
				duDPpte++
				gZDpjWvXax_q += az8aTv8dcu
				zQ_5oM22Q, gouGLnAk = gouGLnAk, gZDpjWvXax_q*gouGLnAk-zQ_5oM22Q
			}
			b8CRqLSt := d2X2v9 + d2X2v9
			aX6xkg_1G := 0.0
			for jJklNl_1r := 2 * (d2X2v9 + duDPpte); jJklNl_1r >= b8CRqLSt; jJklNl_1r -= 2 {
				aX6xkg_1G = 1 / ( /*line :1*/ float64(jJklNl_1r)/_BqkSz0 - aX6xkg_1G)
			}
			a020cdSBAIO := aX6xkg_1G
			alfthQvYkJ = 1

			rzdroM5dZ6jU :=  /*line :1*/ float64(d2X2v9)
			hLZMcJsdS := 2 / _BqkSz0
			rzdroM5dZ6jU = rzdroM5dZ6jU *  /*line :1*/ JrJVPb5WbG( /*line :1*/ Abs(hLZMcJsdS*rzdroM5dZ6jU))
			if rzdroM5dZ6jU < 7.09782712893383973096e+02 {
				for jJklNl_1r := d2X2v9 - 1; jJklNl_1r > 0; jJklNl_1r-- {
					h9LPXHf :=  /*line :1*/ float64(jJklNl_1r + jJklNl_1r)
					a020cdSBAIO, alfthQvYkJ = alfthQvYkJ, alfthQvYkJ*h9LPXHf/_BqkSz0-a020cdSBAIO
				}
			} else {
				for jJklNl_1r := d2X2v9 - 1; jJklNl_1r > 0; jJklNl_1r-- {
					h9LPXHf :=  /*line :1*/ float64(jJklNl_1r + jJklNl_1r)
					a020cdSBAIO, alfthQvYkJ = alfthQvYkJ, alfthQvYkJ*h9LPXHf/_BqkSz0-a020cdSBAIO

					if alfthQvYkJ > 1e100 {
						a020cdSBAIO /= alfthQvYkJ
						aX6xkg_1G /= alfthQvYkJ
						alfthQvYkJ = 1
					}
				}
			}
			alfthQvYkJ = aX6xkg_1G *  /*line :1*/ TZVYtdEXoPFm(_BqkSz0) / alfthQvYkJ
		}
	}
	if awRlyItEG2gn {
		return -alfthQvYkJ
	}
	return alfthQvYkJ
}










func AlpeWvYY4zD(d2X2v9 int, _BqkSz0 float64) float64 {
	const Two302 = 1 << 302	

	switch {
	case _BqkSz0 < 0 ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return 0
	}

	if d2X2v9 == 0 {
		return  /*line :1*/ XiHpf21SIVk(_BqkSz0)
	}
	if _BqkSz0 == 0 {
		if d2X2v9 < 0 && d2X2v9&1 == 1 {
			return  /*line :1*/ FSug4WHZSBwz(1)
		}
		return  /*line :1*/ FSug4WHZSBwz(-1)
	}
	awRlyItEG2gn := false
	if d2X2v9 < 0 {
		d2X2v9 = -d2X2v9
		if d2X2v9&1 == 1 {
			awRlyItEG2gn = true
		}
	}
	if d2X2v9 == 1 {
		if awRlyItEG2gn {
			return - /*line :1*/ VyOoq6VwJ(_BqkSz0)
		}
		return  /*line :1*/ VyOoq6VwJ(_BqkSz0)
	}
	var alfthQvYkJ float64
	if _BqkSz0 >= Two302 {

		var dE6bQU3 float64
		switch aK8LO1, cU6qMJer46oH :=  /*line :1*/ BcjJnykwJ(_BqkSz0); d2X2v9 & 3 {
		case 0:
			dE6bQU3 = aK8LO1 - cU6qMJer46oH
		case 1:
			dE6bQU3 = -aK8LO1 - cU6qMJer46oH
		case 2:
			dE6bQU3 = -aK8LO1 + cU6qMJer46oH
		case 3:
			dE6bQU3 = aK8LO1 + cU6qMJer46oH
		}
		alfthQvYkJ = (1 / SqrtPi) * dE6bQU3 /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
	} else {
		a020cdSBAIO :=  /*line :1*/ XiHpf21SIVk(_BqkSz0)
		alfthQvYkJ =  /*line :1*/ VyOoq6VwJ(_BqkSz0)

		for jJklNl_1r := 1; jJklNl_1r < d2X2v9 && ! /*line :1*/ ME1vzpD5wcr(alfthQvYkJ, -1); jJklNl_1r++ {
			a020cdSBAIO, alfthQvYkJ = alfthQvYkJ, ( /*line :1*/ float64(jJklNl_1r+jJklNl_1r)/_BqkSz0)*alfthQvYkJ-a020cdSBAIO
		}
	}
	if awRlyItEG2gn {
		return -alfthQvYkJ
	}
	return alfthQvYkJ
}
