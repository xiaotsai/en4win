//line :1
package big

import (
	bytes "wLlfRPv"
	fmt "kFsoCfy5zWG"
	strconv "zBESu2SrRjP"
)































func (pmEgen2 *WH8dWN) Text(t5iiC4Ssnsci byte, cZUDpnJQn int) string {
	qnOAx3e4em := 10
	if cZUDpnJQn > 0 {
		qnOAx3e4em += cZUDpnJQn
	}
	return  /*line :1*/ string( /*line :1*/ pmEgen2.Append( /*line :1*/ make([]byte, 0, qnOAx3e4em), t5iiC4Ssnsci, cZUDpnJQn))
}



func (pmEgen2 *WH8dWN) String() string {
	return  /*line :1*/ pmEgen2.Text('g', 10)
}



func (pmEgen2 *WH8dWN) Append(c3Zu4RY []byte, akFKInzNpAH byte, cZUDpnJQn int) []byte {

	if pmEgen2.g2AL4u0IBtd {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '-')
	}

	if pmEgen2.fNeR0A == inf {
		if !pmEgen2.g2AL4u0IBtd {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '+')
		}
		return  /*line :1*/ append(c3Zu4RY, "Inf"...)
	}

	switch akFKInzNpAH {
	case 'b':
		return  /*line :1*/ pmEgen2.iXhDCuwCei9(c3Zu4RY)
	case 'p':
		return  /*line :1*/ pmEgen2.kOqtFT(c3Zu4RY)
	case 'x':
		return  /*line :1*/ pmEgen2.nd6Oj64Kx(c3Zu4RY, cZUDpnJQn)
	}

	
	var s31g8J dCrlpEM	
	if pmEgen2.fNeR0A == finite {

		 /*line :1*/ s31g8J.init(pmEgen2.s7DDkIlY,  /*line :1*/ int(pmEgen2.khThBIzMZ)- /*line :1*/ pmEgen2.s7DDkIlY.wZwJ3Y0())
	}

	diTgYsHJNUfK := false
	if cZUDpnJQn < 0 {
		diTgYsHJNUfK = true
		 /*line :1*/ s1R7I1(&s31g8J, pmEgen2)

		switch akFKInzNpAH {
		case 'e', 'E':
			cZUDpnJQn =  /*line :1*/ len(s31g8J.mqdHnyxfD) - 1
		case 'f':
			cZUDpnJQn =  /*line :1*/ mWLZAmlj( /*line :1*/ len(s31g8J.mqdHnyxfD)-s31g8J.pMAiADDZ, 0)
		case 'g', 'G':
			cZUDpnJQn =  /*line :1*/ len(s31g8J.mqdHnyxfD)
		}
	} else {

		switch akFKInzNpAH {
		case 'e', 'E':

			 /*line :1*/ s31g8J.jrT168o(1 + cZUDpnJQn)
		case 'f':

			 /*line :1*/ s31g8J.jrT168o(s31g8J.pMAiADDZ + cZUDpnJQn)
		case 'g', 'G':
			if cZUDpnJQn == 0 {
				cZUDpnJQn = 1
			}
			 /*line :1*/ s31g8J.jrT168o(cZUDpnJQn)
		}
	}

	switch akFKInzNpAH {
	case 'e', 'E':
		return  /*line :1*/ f0YPujuFtx(c3Zu4RY, akFKInzNpAH, cZUDpnJQn, s31g8J)
	case 'f':
		return  /*line :1*/ oiyufO(c3Zu4RY, cZUDpnJQn, s31g8J)
	case 'g', 'G':

		qW3lfCcBMf_ := cZUDpnJQn
		if qW3lfCcBMf_ >  /*line :1*/ len(s31g8J.mqdHnyxfD) &&  /*line :1*/ len(s31g8J.mqdHnyxfD) >= s31g8J.pMAiADDZ {
			qW3lfCcBMf_ =  /*line :1*/ len(s31g8J.mqdHnyxfD)
		}

		if diTgYsHJNUfK {
			qW3lfCcBMf_ = 6
		}
		xXsI4WC := s31g8J.pMAiADDZ - 1
		if xXsI4WC < -4 || xXsI4WC >= qW3lfCcBMf_ {
			if cZUDpnJQn >  /*line :1*/ len(s31g8J.mqdHnyxfD) {
				cZUDpnJQn =  /*line :1*/ len(s31g8J.mqdHnyxfD)
			}
			return  /*line :1*/ f0YPujuFtx(c3Zu4RY, akFKInzNpAH+'e'-'g', cZUDpnJQn-1, s31g8J)
		}
		if cZUDpnJQn > s31g8J.pMAiADDZ {
			cZUDpnJQn =  /*line :1*/ len(s31g8J.mqdHnyxfD)
		}
		return  /*line :1*/ oiyufO(c3Zu4RY,  /*line :1*/ mWLZAmlj(cZUDpnJQn-s31g8J.pMAiADDZ, 0), s31g8J)
	}

	if pmEgen2.g2AL4u0IBtd {
		c3Zu4RY = c3Zu4RY[: /*line :1*/ len(c3Zu4RY)-1]
	}
	return  /*line :1*/ append(c3Zu4RY, '%', akFKInzNpAH)
}

func s1R7I1(s31g8J *dCrlpEM, pmEgen2 *WH8dWN) {

	if  /*line :1*/ len(s31g8J.mqdHnyxfD) == 0 {
		return
	}

	kfG4qg_2AdyH :=  /*line :1*/ g3X9fa(nil).bj0nVC(pmEgen2.s7DDkIlY)
	xXsI4WC :=  /*line :1*/ int(pmEgen2.khThBIzMZ) -  /*line :1*/ kfG4qg_2AdyH.wZwJ3Y0()
	nR7KU4mGsdy :=  /*line :1*/ kfG4qg_2AdyH.wZwJ3Y0() -  /*line :1*/ int(pmEgen2.iaaUnOsr7_M+1)
	switch {
	case nR7KU4mGsdy < 0:
		kfG4qg_2AdyH =  /*line :1*/ kfG4qg_2AdyH.z839tk6CmDHT(kfG4qg_2AdyH,  /*line :1*/ uint(-nR7KU4mGsdy))
	case nR7KU4mGsdy > 0:
		kfG4qg_2AdyH =  /*line :1*/ kfG4qg_2AdyH.qzOaHmR(kfG4qg_2AdyH,  /*line :1*/ uint(+nR7KU4mGsdy))
	}
	xXsI4WC += nR7KU4mGsdy

	
	var symnaqFHlu4I dCrlpEM
	var jRASocRnuh g3X9fa
	 /*line :1*/ symnaqFHlu4I.init( /*line :1*/ jRASocRnuh.lt4VKILNCo(kfG4qg_2AdyH, lG0t2KfqGwyP), xXsI4WC)

	
	var v0oRq9 dCrlpEM
	 /*line :1*/ v0oRq9.init( /*line :1*/ jRASocRnuh.jXv7nnhY_(kfG4qg_2AdyH, lG0t2KfqGwyP), xXsI4WC)

	ed6DnTVnQGm := kfG4qg_2AdyH[0]&2 == 0

	for aepbxLiWOZ, gS2TXhHpYaX4 := range s31g8J.mqdHnyxfD {
		zhswH4n :=  /*line :1*/ symnaqFHlu4I.dTPXZnAco7(aepbxLiWOZ)
		sOU3zRv :=  /*line :1*/ v0oRq9.dTPXZnAco7(aepbxLiWOZ)

		t0dhDEP6p7 := zhswH4n != gS2TXhHpYaX4 || ed6DnTVnQGm && aepbxLiWOZ+1 ==  /*line :1*/ len(symnaqFHlu4I.mqdHnyxfD)

		aDVoZxR := gS2TXhHpYaX4 != sOU3zRv && (ed6DnTVnQGm || gS2TXhHpYaX4+1 < sOU3zRv || aepbxLiWOZ+1 <  /*line :1*/ len(v0oRq9.mqdHnyxfD))

		switch {
		case t0dhDEP6p7 && aDVoZxR:
			 /*line :1*/ s31g8J.jrT168o(aepbxLiWOZ + 1)
			return
		case t0dhDEP6p7:
			 /*line :1*/ s31g8J.dTXcy3_dM8(aepbxLiWOZ + 1)
			return
		case aDVoZxR:
			 /*line :1*/ s31g8J.i4uBHJlZWXcM(aepbxLiWOZ + 1)
			return
		}
	}
}


func f0YPujuFtx(c3Zu4RY []byte, akFKInzNpAH byte, cZUDpnJQn int, s31g8J dCrlpEM) []byte {

	qvA6MnWV :=  /*line :1*/ byte('0')
	if  /*line :1*/ len(s31g8J.mqdHnyxfD) > 0 {
		qvA6MnWV = s31g8J.mqdHnyxfD[0]
	}
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, qvA6MnWV)

	if cZUDpnJQn > 0 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '.')
		aepbxLiWOZ := 1
		gS2TXhHpYaX4 :=  /*line :1*/ rjFa9Vt( /*line :1*/ len(s31g8J.mqdHnyxfD), cZUDpnJQn+1)
		if aepbxLiWOZ < gS2TXhHpYaX4 {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY, s31g8J.mqdHnyxfD[aepbxLiWOZ:gS2TXhHpYaX4]...)
			aepbxLiWOZ = gS2TXhHpYaX4
		}
		for ; aepbxLiWOZ <= cZUDpnJQn; aepbxLiWOZ++ {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
		}
	}

	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, akFKInzNpAH)
	var xXsI4WC int64
	if  /*line :1*/ len(s31g8J.mqdHnyxfD) > 0 {
		xXsI4WC =  /*line :1*/ int64(s31g8J.pMAiADDZ) - 1
	}
	if xXsI4WC < 0 {
		qvA6MnWV = '-'
		xXsI4WC = -xXsI4WC
	} else {
		qvA6MnWV = '+'
	}
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, qvA6MnWV)

	if xXsI4WC < 10 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
	}
	return  /*line :1*/ strconv.YOFr9xxiI(c3Zu4RY, xXsI4WC, 10)
}


func oiyufO(c3Zu4RY []byte, cZUDpnJQn int, s31g8J dCrlpEM) []byte {

	if s31g8J.pMAiADDZ > 0 {
		gS2TXhHpYaX4 :=  /*line :1*/ rjFa9Vt( /*line :1*/ len(s31g8J.mqdHnyxfD), s31g8J.pMAiADDZ)
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, s31g8J.mqdHnyxfD[:gS2TXhHpYaX4]...)
		for ; gS2TXhHpYaX4 < s31g8J.pMAiADDZ; gS2TXhHpYaX4++ {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
		}
	} else {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
	}

	if cZUDpnJQn > 0 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '.')
		for aepbxLiWOZ := 0; aepbxLiWOZ < cZUDpnJQn; aepbxLiWOZ++ {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY,  /*line :1*/ s31g8J.dTPXZnAco7(s31g8J.pMAiADDZ+aepbxLiWOZ))
		}
	}

	return c3Zu4RY
}








func (pmEgen2 *WH8dWN) iXhDCuwCei9(c3Zu4RY []byte) []byte {
	if pmEgen2.fNeR0A == zero {
		return  /*line :1*/ append(c3Zu4RY, '0')
	}

	if debugFloat && pmEgen2.fNeR0A != finite {
		 /*line :1*/ panic("non-finite float")
	}

	gS2TXhHpYaX4 := pmEgen2.s7DDkIlY
	switch uKQ8Eox :=  /*line :1*/ uint32( /*line :1*/ len(pmEgen2.s7DDkIlY)) * _W; {
	case uKQ8Eox < pmEgen2.iaaUnOsr7_M:
		gS2TXhHpYaX4 =  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(gS2TXhHpYaX4,  /*line :1*/ uint(pmEgen2.iaaUnOsr7_M-uKQ8Eox))
	case uKQ8Eox > pmEgen2.iaaUnOsr7_M:
		gS2TXhHpYaX4 =  /*line :1*/ g3X9fa(nil).qzOaHmR(gS2TXhHpYaX4,  /*line :1*/ uint(uKQ8Eox-pmEgen2.iaaUnOsr7_M))
	}

	c3Zu4RY =  /*line :1*/ append(c3Zu4RY,  /*line :1*/ gS2TXhHpYaX4.g3J8Fbn4Ip(10)...)
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, 'p')
	vRGCE256ba4X :=  /*line :1*/ int64(pmEgen2.khThBIzMZ) -  /*line :1*/ int64(pmEgen2.iaaUnOsr7_M)
	if vRGCE256ba4X >= 0 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '+')
	}
	return  /*line :1*/ strconv.YOFr9xxiI(c3Zu4RY, vRGCE256ba4X, 10)
}







func (pmEgen2 *WH8dWN) nd6Oj64Kx(c3Zu4RY []byte, cZUDpnJQn int) []byte {
	if pmEgen2.fNeR0A == zero {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, "0x0"...)
		if cZUDpnJQn > 0 {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '.')
			for aepbxLiWOZ := 0; aepbxLiWOZ < cZUDpnJQn; aepbxLiWOZ++ {
				c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
			}
		}
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, "p+00"...)
		return c3Zu4RY
	}

	if debugFloat && pmEgen2.fNeR0A != finite {
		 /*line :1*/ panic("non-finite float")
	}

	
	var h_Wr_yqq uint
	if cZUDpnJQn < 0 {
		h_Wr_yqq = 1 + ( /*line :1*/ pmEgen2.MinPrec()-1+3)/4*4
	} else {
		h_Wr_yqq = 1 + 4* /*line :1*/ uint(cZUDpnJQn)
	}

	pmEgen2 =  /*line :1*/ new(WH8dWN).SetPrec(h_Wr_yqq).SetMode(pmEgen2.hNG8iJiHdkxT).Set(pmEgen2)

	gS2TXhHpYaX4 := pmEgen2.s7DDkIlY
	switch uKQ8Eox :=  /*line :1*/ uint( /*line :1*/ len(pmEgen2.s7DDkIlY)) * _W; {
	case uKQ8Eox < h_Wr_yqq:
		gS2TXhHpYaX4 =  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(gS2TXhHpYaX4, h_Wr_yqq-uKQ8Eox)
	case uKQ8Eox > h_Wr_yqq:
		gS2TXhHpYaX4 =  /*line :1*/ g3X9fa(nil).qzOaHmR(gS2TXhHpYaX4, uKQ8Eox-h_Wr_yqq)
	}
	hm7pd0I3 :=  /*line :1*/ int64(pmEgen2.khThBIzMZ) - 1

	ioE6aLT8 :=  /*line :1*/ gS2TXhHpYaX4.g3J8Fbn4Ip(16)
	if debugFloat && ioE6aLT8[0] != '1' {
		 /*line :1*/ panic("incorrect mantissa: " +  /*line :1*/ string(ioE6aLT8))
	}
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, "0x1"...)
	if  /*line :1*/ len(ioE6aLT8) > 1 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '.')
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, ioE6aLT8[1:]...)
	}

	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, 'p')
	if hm7pd0I3 >= 0 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '+')
	} else {
		hm7pd0I3 = -hm7pd0I3
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '-')
	}

	if hm7pd0I3 < 10 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
	}
	return  /*line :1*/ strconv.YOFr9xxiI(c3Zu4RY, hm7pd0I3, 10)
}







func (pmEgen2 *WH8dWN) kOqtFT(c3Zu4RY []byte) []byte {
	if pmEgen2.fNeR0A == zero {
		return  /*line :1*/ append(c3Zu4RY, '0')
	}

	if debugFloat && pmEgen2.fNeR0A != finite {
		 /*line :1*/ panic("non-finite float")
	}

	gS2TXhHpYaX4 := pmEgen2.s7DDkIlY
	aepbxLiWOZ := 0
	for aepbxLiWOZ <  /*line :1*/ len(gS2TXhHpYaX4) && gS2TXhHpYaX4[aepbxLiWOZ] == 0 {
		aepbxLiWOZ++
	}
	gS2TXhHpYaX4 = gS2TXhHpYaX4[aepbxLiWOZ:]

	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, "0x."...)
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY,  /*line :1*/ bytes.TBypK87Hkc( /*line :1*/ gS2TXhHpYaX4.g3J8Fbn4Ip(16), "0")...)
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, 'p')
	if pmEgen2.khThBIzMZ >= 0 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '+')
	}
	return  /*line :1*/ strconv.YOFr9xxiI(c3Zu4RY,  /*line :1*/ int64(pmEgen2.khThBIzMZ), 10)
}

func rjFa9Vt(pmEgen2, oxFS5wa5 int) int {
	if pmEgen2 < oxFS5wa5 {
		return pmEgen2
	}
	return oxFS5wa5
}

var _ fmt.OpUJYyZOy8 = &qiwusFPu	










func (pmEgen2 *WH8dWN) Format(nR7KU4mGsdy fmt.Bm1Jqy8i, t5iiC4Ssnsci rune) {
	cZUDpnJQn, bNnpPgTwPsV :=  /*line :1*/ nR7KU4mGsdy.Precision()
	if !bNnpPgTwPsV {
		cZUDpnJQn = 6
	}

	switch t5iiC4Ssnsci {
	case 'e', 'E', 'f', 'b', 'p', 'x':

	case 'F':

		t5iiC4Ssnsci = 'f'
	case 'v':

		t5iiC4Ssnsci = 'g'
		fallthrough
	case 'g', 'G':
		if !bNnpPgTwPsV {
			cZUDpnJQn = -1
		}
	default:
		 /*line :1*/ fmt.BYcL2whVEYz(nR7KU4mGsdy, "%%!%c(*big.Float=%s)", t5iiC4Ssnsci,  /*line :1*/ pmEgen2.String())
		return
	}
	var c3Zu4RY []byte
	c3Zu4RY =  /*line :1*/ pmEgen2.Append(c3Zu4RY,  /*line :1*/ byte(t5iiC4Ssnsci), cZUDpnJQn)
	if  /*line :1*/ len(c3Zu4RY) == 0 {
		c3Zu4RY = [] /*line :1*/ byte("?")
	}

	var e2mU4H string
	switch {
	case c3Zu4RY[0] == '-':
		e2mU4H = "-"
		c3Zu4RY = c3Zu4RY[1:]
	case c3Zu4RY[0] == '+':

		e2mU4H = "+"
		if  /*line :1*/ nR7KU4mGsdy.Flag(' ') {
			e2mU4H = " "
		}
		c3Zu4RY = c3Zu4RY[1:]
	case  /*line :1*/ nR7KU4mGsdy.Flag('+'):
		e2mU4H = "+"
	case  /*line :1*/ nR7KU4mGsdy.Flag(' '):
		e2mU4H = " "
	}

	var c8nKd_1lQn int
	if pRxc6yBZ9, hLmsaN :=  /*line :1*/ nR7KU4mGsdy.Width(); hLmsaN && pRxc6yBZ9 >  /*line :1*/ len(e2mU4H)+ /*line :1*/ len(c3Zu4RY) {
		c8nKd_1lQn = pRxc6yBZ9 -  /*line :1*/ len(e2mU4H) -  /*line :1*/ len(c3Zu4RY)
	}

	switch {
	case  /*line :1*/ nR7KU4mGsdy.Flag('0') && ! /*line :1*/ pmEgen2.IsInf():

		 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, e2mU4H, 1)
		 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, "0", c8nKd_1lQn)
		 /*line :1*/ nR7KU4mGsdy.Write(c3Zu4RY)
	case  /*line :1*/ nR7KU4mGsdy.Flag('-'):

		 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, e2mU4H, 1)
		 /*line :1*/ nR7KU4mGsdy.Write(c3Zu4RY)
		 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, " ", c8nKd_1lQn)
	default:

		 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, " ", c8nKd_1lQn)
		 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, e2mU4H, 1)
		 /*line :1*/ nR7KU4mGsdy.Write(c3Zu4RY)
	}
}
