//line :1
package big

import (
	binary "yLm9uBQG"
	errors "iAHoxjmM"
	fmt "kFsoCfy5zWG"
)


const floatGobVersion byte = 1




func (pmEgen2 *WH8dWN) GobEncode() ([]byte, error) {
	if pmEgen2 == nil {
		return nil, nil
	}

	lOdaFO := 1 + 1 + 4
	h_Wr_yqq := 0
	if pmEgen2.fNeR0A == finite {

		h_Wr_yqq =  /*line :1*/ int((pmEgen2.iaaUnOsr7_M + (_W - 1)) / _W)

		if  /*line :1*/ len(pmEgen2.s7DDkIlY) < h_Wr_yqq {
			h_Wr_yqq =  /*line :1*/ len(pmEgen2.s7DDkIlY)
		}

		lOdaFO += 4 + h_Wr_yqq*_S
	}
	c3Zu4RY :=  /*line :1*/ make([]byte, lOdaFO)

	c3Zu4RY[0] = floatGobVersion
	kWa1YrIHZo :=  /*line :1*/ byte(pmEgen2.hNG8iJiHdkxT&7)<<5 |  /*line :1*/ byte((pmEgen2.ddq1eQm+1)&3)<<3 |  /*line :1*/ byte(pmEgen2.fNeR0A&3)<<1
	if pmEgen2.g2AL4u0IBtd {
		kWa1YrIHZo |= 1
	}
	c3Zu4RY[1] = kWa1YrIHZo
	 /*line :1*/ binary.FOcKBq6.PutUint32(c3Zu4RY[2:], pmEgen2.iaaUnOsr7_M)

	if pmEgen2.fNeR0A == finite {
		 /*line :1*/ binary.FOcKBq6.PutUint32(c3Zu4RY[6:],  /*line :1*/ uint32(pmEgen2.khThBIzMZ))
		 /*line :1*/ pmEgen2.s7DDkIlY[ /*line :1*/ len(pmEgen2.s7DDkIlY)-h_Wr_yqq:].qonqPl(c3Zu4RY[10:])
	}

	return c3Zu4RY, nil
}





func (uW0eIVjy *WH8dWN) GobDecode(c3Zu4RY []byte) error {
	if  /*line :1*/ len(c3Zu4RY) == 0 {

		*uW0eIVjy = WH8dWN{}
		return nil
	}
	if  /*line :1*/ len(c3Zu4RY) < 6 {
		return  /*line :1*/ errors.Su6g6hRoi9X("Float.GobDecode: buffer too small")
	}

	if c3Zu4RY[0] != floatGobVersion {
		return  /*line :1*/ fmt.Cnz_ab1BaZh("Float.GobDecode: encoding version %d not supported", c3Zu4RY[0])
	}

	aovRVQZv := uW0eIVjy.iaaUnOsr7_M
	bRv0XJH5h := uW0eIVjy.hNG8iJiHdkxT

	kWa1YrIHZo := c3Zu4RY[1]
	uW0eIVjy.hNG8iJiHdkxT =  /*line :1*/ WsODiYye5l((kWa1YrIHZo >> 5) & 7)
	uW0eIVjy.ddq1eQm =  /*line :1*/ IG11Ul7qo((kWa1YrIHZo>>3)&3) - 1
	uW0eIVjy.fNeR0A =  /*line :1*/ cZg9qpR((kWa1YrIHZo >> 1) & 3)
	uW0eIVjy.g2AL4u0IBtd = kWa1YrIHZo&1 != 0
	uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ binary.FOcKBq6.Uint32(c3Zu4RY[2:])

	if uW0eIVjy.fNeR0A == finite {
		if  /*line :1*/ len(c3Zu4RY) < 10 {
			return  /*line :1*/ errors.Su6g6hRoi9X("Float.GobDecode: buffer too small for finite form float")
		}
		uW0eIVjy.khThBIzMZ =  /*line :1*/ int32( /*line :1*/ binary.FOcKBq6.Uint32(c3Zu4RY[6:]))
		uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.vbFROX(c3Zu4RY[10:])
	}

	if aovRVQZv != 0 {
		uW0eIVjy.hNG8iJiHdkxT = bRv0XJH5h
		 /*line :1*/ uW0eIVjy.SetPrec( /*line :1*/ uint(aovRVQZv))
	}

	if isKdjaXeB :=  /*line :1*/ uW0eIVjy.asvxnebpcnsA(); isKdjaXeB != "" {
		return  /*line :1*/ errors.Su6g6hRoi9X("Float.GobDecode: " + isKdjaXeB)
	}

	return nil
}




func (pmEgen2 *WH8dWN) MarshalText() (yK8eFt []byte, gDonrwI8VP50 error) {
	if pmEgen2 == nil {
		return [] /*line :1*/ byte("<nil>"), nil
	}
	var c3Zu4RY []byte
	return  /*line :1*/ pmEgen2.Append(c3Zu4RY, 'g', -1), nil
}





func (uW0eIVjy *WH8dWN) UnmarshalText(yK8eFt []byte) error {

	_, _, gDonrwI8VP50 :=  /*line :1*/ uW0eIVjy.Parse( /*line :1*/ string(yK8eFt), 0)
	if gDonrwI8VP50 != nil {
		gDonrwI8VP50 =  /*line :1*/ fmt.Cnz_ab1BaZh("math/big: cannot unmarshal %q into a *big.Float (%v)", yK8eFt, gDonrwI8VP50)
	}
	return gDonrwI8VP50
}
