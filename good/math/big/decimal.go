//line :1
package big






type dCrlpEM struct {
	mqdHnyxfD	[]byte	
	pMAiADDZ	int	
}


func (s31g8J *dCrlpEM) dTPXZnAco7(aepbxLiWOZ int) byte {
	if 0 <= aepbxLiWOZ && aepbxLiWOZ <  /*line :1*/ len(s31g8J.mqdHnyxfD) {
		return s31g8J.mqdHnyxfD[aepbxLiWOZ]
	}
	return '0'
}



const maxShift = _W - 4



func (pmEgen2 *dCrlpEM) init(gS2TXhHpYaX4 g3X9fa, qKn7aIMh5 int) {

	if  /*line :1*/ len(gS2TXhHpYaX4) == 0 {
		pmEgen2.mqdHnyxfD = pmEgen2.mqdHnyxfD[:0]
		pmEgen2.pMAiADDZ = 0
		return
	}

	if qKn7aIMh5 < 0 {
		b_CDmuq7 :=  /*line :1*/ gS2TXhHpYaX4.kDun6ak()
		nR7KU4mGsdy :=  /*line :1*/ uint(-qKn7aIMh5)
		if nR7KU4mGsdy >= b_CDmuq7 {
			nR7KU4mGsdy = b_CDmuq7
		}
		gS2TXhHpYaX4 =  /*line :1*/ g3X9fa(nil).qzOaHmR(gS2TXhHpYaX4, nR7KU4mGsdy)
		qKn7aIMh5 +=  /*line :1*/ int(nR7KU4mGsdy)
	}

	if qKn7aIMh5 > 0 {
		gS2TXhHpYaX4 =  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(gS2TXhHpYaX4,  /*line :1*/ uint(qKn7aIMh5))
		qKn7aIMh5 = 0
	}

	nR7KU4mGsdy :=  /*line :1*/ gS2TXhHpYaX4.g3J8Fbn4Ip(10)
	h_Wr_yqq :=  /*line :1*/ len(nR7KU4mGsdy)
	pmEgen2.pMAiADDZ = h_Wr_yqq

	for h_Wr_yqq > 0 && nR7KU4mGsdy[h_Wr_yqq-1] == '0' {
		h_Wr_yqq--
	}
	pmEgen2.mqdHnyxfD =  /*line :1*/ append(pmEgen2.mqdHnyxfD[:0], nR7KU4mGsdy[:h_Wr_yqq]...)

	if qKn7aIMh5 < 0 {
		for qKn7aIMh5 < -maxShift {
			 /*line :1*/ qzOaHmR(pmEgen2, maxShift)
			qKn7aIMh5 += maxShift
		}
		 /*line :1*/ qzOaHmR(pmEgen2,  /*line :1*/ uint(-qKn7aIMh5))
	}
}


func qzOaHmR(pmEgen2 *dCrlpEM, nR7KU4mGsdy uint) {

	vFx5p_cm := 0
	var h_Wr_yqq VYe6D0
	for h_Wr_yqq>>nR7KU4mGsdy == 0 && vFx5p_cm <  /*line :1*/ len(pmEgen2.mqdHnyxfD) {
		qvA6MnWV :=  /*line :1*/ VYe6D0(pmEgen2.mqdHnyxfD[vFx5p_cm])
		vFx5p_cm++
		h_Wr_yqq = h_Wr_yqq*10 + qvA6MnWV - '0'
	}
	if h_Wr_yqq == 0 {

		pmEgen2.mqdHnyxfD = pmEgen2.mqdHnyxfD[:0]
		return
	}
	for h_Wr_yqq>>nR7KU4mGsdy == 0 {
		vFx5p_cm++
		h_Wr_yqq *= 10
	}
	pmEgen2.pMAiADDZ += 1 - vFx5p_cm

	uKQ8Eox := 0
	yWg_RX :=  /*line :1*/ VYe6D0(1)<<nR7KU4mGsdy - 1
	for vFx5p_cm <  /*line :1*/ len(pmEgen2.mqdHnyxfD) {
		qvA6MnWV :=  /*line :1*/ VYe6D0(pmEgen2.mqdHnyxfD[vFx5p_cm])
		vFx5p_cm++
		s31g8J := h_Wr_yqq >> nR7KU4mGsdy
		h_Wr_yqq &= yWg_RX
		pmEgen2.mqdHnyxfD[uKQ8Eox] =  /*line :1*/ byte(s31g8J + '0')
		uKQ8Eox++
		h_Wr_yqq = h_Wr_yqq*10 + qvA6MnWV - '0'
	}

	for h_Wr_yqq > 0 && uKQ8Eox <  /*line :1*/ len(pmEgen2.mqdHnyxfD) {
		s31g8J := h_Wr_yqq >> nR7KU4mGsdy
		h_Wr_yqq &= yWg_RX
		pmEgen2.mqdHnyxfD[uKQ8Eox] =  /*line :1*/ byte(s31g8J + '0')
		uKQ8Eox++
		h_Wr_yqq = h_Wr_yqq * 10
	}
	pmEgen2.mqdHnyxfD = pmEgen2.mqdHnyxfD[:uKQ8Eox]

	for h_Wr_yqq > 0 {
		s31g8J := h_Wr_yqq >> nR7KU4mGsdy
		h_Wr_yqq &= yWg_RX
		pmEgen2.mqdHnyxfD =  /*line :1*/ append(pmEgen2.mqdHnyxfD,  /*line :1*/ byte(s31g8J+'0'))
		h_Wr_yqq = h_Wr_yqq * 10
	}

	 /*line :1*/ e65t874MPFR(pmEgen2)
}

func (pmEgen2 *dCrlpEM) String() string {
	if  /*line :1*/ len(pmEgen2.mqdHnyxfD) == 0 {
		return "0"
	}

	var c3Zu4RY []byte
	switch {
	case pmEgen2.pMAiADDZ <= 0:

		c3Zu4RY =  /*line :1*/ make([]byte, 0, 2+(-pmEgen2.pMAiADDZ)+ /*line :1*/ len(pmEgen2.mqdHnyxfD))
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, "0."...)
		c3Zu4RY =  /*line :1*/ gwGqBHlLAz(c3Zu4RY, -pmEgen2.pMAiADDZ)
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, pmEgen2.mqdHnyxfD...)

	case pmEgen2.pMAiADDZ <  /*line :1*/ len(pmEgen2.mqdHnyxfD):

		c3Zu4RY =  /*line :1*/ make([]byte, 0, 1+ /*line :1*/ len(pmEgen2.mqdHnyxfD))
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, pmEgen2.mqdHnyxfD[:pmEgen2.pMAiADDZ]...)
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '.')
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, pmEgen2.mqdHnyxfD[pmEgen2.pMAiADDZ:]...)

	default:

		c3Zu4RY =  /*line :1*/ make([]byte, 0, pmEgen2.pMAiADDZ)
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, pmEgen2.mqdHnyxfD...)
		c3Zu4RY =  /*line :1*/ gwGqBHlLAz(c3Zu4RY, pmEgen2.pMAiADDZ- /*line :1*/ len(pmEgen2.mqdHnyxfD))
	}

	return  /*line :1*/ string(c3Zu4RY)
}


func gwGqBHlLAz(c3Zu4RY []byte, h_Wr_yqq int) []byte {
	for ; h_Wr_yqq > 0; h_Wr_yqq-- {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
	}
	return c3Zu4RY
}




func mjHDlBNjP5(pmEgen2 *dCrlpEM, h_Wr_yqq int) bool {
	if pmEgen2.mqdHnyxfD[h_Wr_yqq] == '5' && h_Wr_yqq+1 ==  /*line :1*/ len(pmEgen2.mqdHnyxfD) {

		return h_Wr_yqq > 0 && (pmEgen2.mqdHnyxfD[h_Wr_yqq-1]-'0')&1 != 0
	}

	return pmEgen2.mqdHnyxfD[h_Wr_yqq] >= '5'
}




func (pmEgen2 *dCrlpEM) jrT168o(h_Wr_yqq int) {
	if h_Wr_yqq < 0 || h_Wr_yqq >=  /*line :1*/ len(pmEgen2.mqdHnyxfD) {
		return
	}

	if  /*line :1*/ mjHDlBNjP5(pmEgen2, h_Wr_yqq) {
		 /*line :1*/ pmEgen2.i4uBHJlZWXcM(h_Wr_yqq)
	} else {
		 /*line :1*/ pmEgen2.dTXcy3_dM8(h_Wr_yqq)
	}
}

func (pmEgen2 *dCrlpEM) i4uBHJlZWXcM(h_Wr_yqq int) {
	if h_Wr_yqq < 0 || h_Wr_yqq >=  /*line :1*/ len(pmEgen2.mqdHnyxfD) {
		return
	}

	for h_Wr_yqq > 0 && pmEgen2.mqdHnyxfD[h_Wr_yqq-1] >= '9' {
		h_Wr_yqq--
	}

	if h_Wr_yqq == 0 {

		pmEgen2.mqdHnyxfD[0] = '1'
		pmEgen2.mqdHnyxfD = pmEgen2.mqdHnyxfD[:1]
		pmEgen2.pMAiADDZ++
		return
	}

	pmEgen2.mqdHnyxfD[h_Wr_yqq-1]++
	pmEgen2.mqdHnyxfD = pmEgen2.mqdHnyxfD[:h_Wr_yqq]

}

func (pmEgen2 *dCrlpEM) dTXcy3_dM8(h_Wr_yqq int) {
	if h_Wr_yqq < 0 || h_Wr_yqq >=  /*line :1*/ len(pmEgen2.mqdHnyxfD) {
		return
	}
	pmEgen2.mqdHnyxfD = pmEgen2.mqdHnyxfD[:h_Wr_yqq]
	 /*line :1*/ e65t874MPFR(pmEgen2)
}



func e65t874MPFR(pmEgen2 *dCrlpEM) {
	aepbxLiWOZ :=  /*line :1*/ len(pmEgen2.mqdHnyxfD)
	for aepbxLiWOZ > 0 && pmEgen2.mqdHnyxfD[aepbxLiWOZ-1] == '0' {
		aepbxLiWOZ--
	}
	pmEgen2.mqdHnyxfD = pmEgen2.mqdHnyxfD[:aepbxLiWOZ]
	if aepbxLiWOZ == 0 {
		pmEgen2.pMAiADDZ = 0
	}
}
