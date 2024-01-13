//line :1
package math








func MI6hKP(_BqkSz0 float64) float64 {
	if haveArchAcosh {
		return  /*line :1*/ r6JQh_EwvY(_BqkSz0)
	}
	return  /*line :1*/ d9lRgsQ2vo(_BqkSz0)
}

func d9lRgsQ2vo(_BqkSz0 float64) float64 {
	const Large = 1 << 28	

	switch {
	case _BqkSz0 < 1 ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case _BqkSz0 == 1:
		return 0
	case _BqkSz0 >= Large:
		return  /*line :1*/ JrJVPb5WbG(_BqkSz0) + Ln2
	case _BqkSz0 > 2:
		return  /*line :1*/ JrJVPb5WbG(2*_BqkSz0 - 1/(_BqkSz0+ /*line :1*/ NeXs7bSyfaCD(_BqkSz0*_BqkSz0-1)))
	}
	aX6xkg_1G := _BqkSz0 - 1
	return  /*line :1*/ V_uk99t0rhQ(aX6xkg_1G +  /*line :1*/ NeXs7bSyfaCD(2*aX6xkg_1G+aX6xkg_1G*aX6xkg_1G))
}
