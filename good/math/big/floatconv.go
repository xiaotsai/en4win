//line :1
package big

import (
	fmt "kFsoCfy5zWG"
	io "xI9ai2djaJ2"
	strings "fQXlzv"
)

var qiwusFPu WH8dWN






func (uW0eIVjy *WH8dWN) SetString(nR7KU4mGsdy string) (*WH8dWN, bool) {
	if sYV3C6v3w1, _, gDonrwI8VP50 :=  /*line :1*/ uW0eIVjy.Parse(nR7KU4mGsdy, 0); gDonrwI8VP50 == nil {
		return sYV3C6v3w1, true
	}
	return nil, false
}





func (uW0eIVjy *WH8dWN) h9c5N0bF(vFx5p_cm io.HSeaoj16Wq0, weSJS4i4lm int) (sYV3C6v3w1 *WH8dWN, kWa1YrIHZo int, gDonrwI8VP50 error) {
	cZUDpnJQn := uW0eIVjy.iaaUnOsr7_M
	if cZUDpnJQn == 0 {
		cZUDpnJQn = 64
	}

	uW0eIVjy.fNeR0A = zero

	uW0eIVjy.g2AL4u0IBtd, gDonrwI8VP50 =  /*line :1*/ xOqvCaQwvh(vFx5p_cm)
	if gDonrwI8VP50 != nil {
		return
	}

	
	var kVeBoojLX int	
	uW0eIVjy.s7DDkIlY, kWa1YrIHZo, kVeBoojLX, gDonrwI8VP50 =  /*line :1*/ uW0eIVjy.s7DDkIlY.h9c5N0bF(vFx5p_cm, weSJS4i4lm, true)
	if gDonrwI8VP50 != nil {
		return
	}

	
	var xXsI4WC int64
	var jlDe5QeBx4 int
	xXsI4WC, jlDe5QeBx4, gDonrwI8VP50 =  /*line :1*/ gCqU9SyEYV(vFx5p_cm, true, weSJS4i4lm == 0)
	if gDonrwI8VP50 != nil {
		return
	}

	if  /*line :1*/ len(uW0eIVjy.s7DDkIlY) == 0 {
		uW0eIVjy.iaaUnOsr7_M = cZUDpnJQn
		uW0eIVjy.ddq1eQm = Exact
		uW0eIVjy.fNeR0A = zero
		sYV3C6v3w1 = uW0eIVjy
		return
	}

	ovNyTM :=  /*line :1*/ int64( /*line :1*/ len(uW0eIVjy.s7DDkIlY))*_W -  /*line :1*/ bqzcoCe14(uW0eIVjy.s7DDkIlY)
	eXfVq2GG :=  /*line :1*/ int64(0)

	if kVeBoojLX < 0 {

		s31g8J :=  /*line :1*/ int64(kVeBoojLX)
		switch kWa1YrIHZo {
		case 10:
			eXfVq2GG = s31g8J
			fallthrough
		case 2:
			ovNyTM += s31g8J
		case 8:
			ovNyTM += s31g8J * 3
		case 16:
			ovNyTM += s31g8J * 4
		default:
			 /*line :1*/ panic("unexpected mantissa base")
		}

	}

	switch jlDe5QeBx4 {
	case 10:
		eXfVq2GG += xXsI4WC
		fallthrough
	case 2:
		ovNyTM += xXsI4WC
	default:
		 /*line :1*/ panic("unexpected exponent base")
	}

	if MinExp <= ovNyTM && ovNyTM <= MaxExp {
		uW0eIVjy.iaaUnOsr7_M = cZUDpnJQn
		uW0eIVjy.fNeR0A = finite
		uW0eIVjy.khThBIzMZ =  /*line :1*/ int32(ovNyTM)
		sYV3C6v3w1 = uW0eIVjy
	} else {
		gDonrwI8VP50 =  /*line :1*/ fmt.Cnz_ab1BaZh("exponent overflow")
		return
	}

	if eXfVq2GG == 0 {

		 /*line :1*/ uW0eIVjy.jrT168o(0)
		return
	}

	djOsNJsLfH5U :=  /*line :1*/ new(WH8dWN).SetPrec( /*line :1*/ uW0eIVjy.Prec() + 64)
	if eXfVq2GG < 0 {
		 /*line :1*/ uW0eIVjy.Quo(uW0eIVjy,  /*line :1*/ djOsNJsLfH5U.wwgZlyx( /*line :1*/ uint64(-eXfVq2GG)))
	} else {
		 /*line :1*/ uW0eIVjy.Mul(uW0eIVjy,  /*line :1*/ djOsNJsLfH5U.wwgZlyx( /*line :1*/ uint64(eXfVq2GG)))
	}

	return
}






var htLHwDyiL = [...]uint64{
	1,
	5,
	25,
	125,
	625,
	3125,
	15625,
	78125,
	390625,
	1953125,
	9765625,
	48828125,
	244140625,
	1220703125,
	6103515625,
	30517578125,
	152587890625,
	762939453125,
	3814697265625,
	19073486328125,
	95367431640625,
	476837158203125,
	2384185791015625,
	11920928955078125,
	59604644775390625,
	298023223876953125,
	1490116119384765625,
	7450580596923828125,
}



func (uW0eIVjy *WH8dWN) wwgZlyx(h_Wr_yqq uint64) *WH8dWN {
	const m =  /*line :1*/ uint64( /*line :1*/ len(htLHwDyiL) - 1)
	if h_Wr_yqq <= m {
		return  /*line :1*/ uW0eIVjy.SetUint64(htLHwDyiL[h_Wr_yqq])
	}

	 /*line :1*/ uW0eIVjy.SetUint64(htLHwDyiL[m])
	h_Wr_yqq -= m

	sYV3C6v3w1 :=  /*line :1*/ new(WH8dWN).SetPrec( /*line :1*/ uW0eIVjy.Prec() + 64).SetUint64(5)

	for h_Wr_yqq > 0 {
		if h_Wr_yqq&1 != 0 {
			 /*line :1*/ uW0eIVjy.Mul(uW0eIVjy, sYV3C6v3w1)
		}
		 /*line :1*/ sYV3C6v3w1.Mul(sYV3C6v3w1, sYV3C6v3w1)
		h_Wr_yqq >>= 1
	}

	return uW0eIVjy
}














































func (uW0eIVjy *WH8dWN) Parse(nR7KU4mGsdy string, weSJS4i4lm int) (sYV3C6v3w1 *WH8dWN, kWa1YrIHZo int, gDonrwI8VP50 error) {

	if  /*line :1*/ len(nR7KU4mGsdy) == 3 && (nR7KU4mGsdy == "Inf" || nR7KU4mGsdy == "inf") {
		sYV3C6v3w1 =  /*line :1*/ uW0eIVjy.SetInf(false)
		return
	}
	if  /*line :1*/ len(nR7KU4mGsdy) == 4 && (nR7KU4mGsdy[0] == '+' || nR7KU4mGsdy[0] == '-') && (nR7KU4mGsdy[1:] == "Inf" || nR7KU4mGsdy[1:] == "inf") {
		sYV3C6v3w1 =  /*line :1*/ uW0eIVjy.SetInf(nR7KU4mGsdy[0] == '-')
		return
	}

	vFx5p_cm :=  /*line :1*/ strings.X4yXgtBA(nR7KU4mGsdy)
	if sYV3C6v3w1, kWa1YrIHZo, gDonrwI8VP50 =  /*line :1*/ uW0eIVjy.h9c5N0bF(vFx5p_cm, weSJS4i4lm); gDonrwI8VP50 != nil {
		return
	}

	if qvA6MnWV, law_rimZe :=  /*line :1*/ vFx5p_cm.ReadByte(); law_rimZe == nil {
		gDonrwI8VP50 =  /*line :1*/ fmt.Cnz_ab1BaZh("expected end of string, found %q", qvA6MnWV)
	} else if law_rimZe != io.K5Sqco {
		gDonrwI8VP50 = law_rimZe
	}

	return
}



func VH9qcD(nR7KU4mGsdy string, weSJS4i4lm int, cZUDpnJQn uint, aca9LekD WsODiYye5l) (sYV3C6v3w1 *WH8dWN, kWa1YrIHZo int, gDonrwI8VP50 error) {
	return  /*line :1*/ new(WH8dWN).SetPrec(cZUDpnJQn).SetMode(aca9LekD).Parse(nR7KU4mGsdy, weSJS4i4lm)
}

var _ fmt.MqdnB6 = (* /*line :1*/ WH8dWN)(nil)	






func (uW0eIVjy *WH8dWN) Scan(nR7KU4mGsdy fmt.XDnSs0f3, qvA6MnWV rune) error {
	 /*line :1*/ nR7KU4mGsdy.SkipSpace()
	_, _, gDonrwI8VP50 :=  /*line :1*/ uW0eIVjy.h9c5N0bF(d8BACp{nR7KU4mGsdy}, 0)
	return gDonrwI8VP50
}
