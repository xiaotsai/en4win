//line :1
package big

import (
	fmt "kFsoCfy5zWG"
	io "xI9ai2djaJ2"
	rand "os208GicmRnc"
	strings "fQXlzv"
)


















type ZMtGuo struct {
	hCTOmp0gckSa	bool	
	qho77PBKF	g3X9fa	
}

var h2Aj5SV9g2re = &ZMtGuo{false, lG0t2KfqGwyP}






func (pmEgen2 *ZMtGuo) Sign() int {

	if  /*line :1*/ len(pmEgen2.qho77PBKF) == 0 {
		return 0
	}
	if pmEgen2.hCTOmp0gckSa {
		return -1
	}
	return 1
}


func (uW0eIVjy *ZMtGuo) SetInt64(pmEgen2 int64) *ZMtGuo {
	s_9Ih_ := false
	if pmEgen2 < 0 {
		s_9Ih_ = true
		pmEgen2 = -pmEgen2
	}
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.glzD00sysLql( /*line :1*/ uint64(pmEgen2))
	uW0eIVjy.hCTOmp0gckSa = s_9Ih_
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) SetUint64(pmEgen2 uint64) *ZMtGuo {
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.glzD00sysLql(pmEgen2)
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}


func Hr58YFtm(pmEgen2 int64) *ZMtGuo {

	sOU3zRv :=  /*line :1*/ uint64(pmEgen2)
	if pmEgen2 < 0 {
		sOU3zRv = -sOU3zRv
	}
	var l2IsegfMGX5Z []VYe6D0
	if pmEgen2 == 0 {
	} else if _W == 32 && sOU3zRv>>32 != 0 {
		l2IsegfMGX5Z = []VYe6D0{ /*line :1*/ VYe6D0(sOU3zRv),  /*line :1*/ VYe6D0(sOU3zRv >> 32)}
	} else {
		l2IsegfMGX5Z = []VYe6D0{ /*line :1*/ VYe6D0(sOU3zRv)}
	}
	return &ZMtGuo{hCTOmp0gckSa: pmEgen2 < 0, qho77PBKF: l2IsegfMGX5Z}
}


func (uW0eIVjy *ZMtGuo) Set(pmEgen2 *ZMtGuo) *ZMtGuo {
	if uW0eIVjy != pmEgen2 {
		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.bj0nVC(pmEgen2.qho77PBKF)
		uW0eIVjy.hCTOmp0gckSa = pmEgen2.hCTOmp0gckSa
	}
	return uW0eIVjy
}






func (pmEgen2 *ZMtGuo) Bits() []VYe6D0 {

	return pmEgen2.qho77PBKF
}






func (uW0eIVjy *ZMtGuo) SetBits(l2IsegfMGX5Z []VYe6D0) *ZMtGuo {
	uW0eIVjy.qho77PBKF =  /*line :1*/ g3X9fa(l2IsegfMGX5Z).hitXcy5()
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Abs(pmEgen2 *ZMtGuo) *ZMtGuo {
	 /*line :1*/ uW0eIVjy.Set(pmEgen2)
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Neg(pmEgen2 *ZMtGuo) *ZMtGuo {
	 /*line :1*/ uW0eIVjy.Set(pmEgen2)
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && !uW0eIVjy.hCTOmp0gckSa
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Add(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	s_9Ih_ := pmEgen2.hCTOmp0gckSa
	if pmEgen2.hCTOmp0gckSa == oxFS5wa5.hCTOmp0gckSa {

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
	} else {

		if  /*line :1*/ pmEgen2.qho77PBKF.beaurQHiT(oxFS5wa5.qho77PBKF) >= 0 {
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
		} else {
			s_9Ih_ = !s_9Ih_
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(oxFS5wa5.qho77PBKF, pmEgen2.qho77PBKF)
		}
	}
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && s_9Ih_
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Sub(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	s_9Ih_ := pmEgen2.hCTOmp0gckSa
	if pmEgen2.hCTOmp0gckSa != oxFS5wa5.hCTOmp0gckSa {

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
	} else {

		if  /*line :1*/ pmEgen2.qho77PBKF.beaurQHiT(oxFS5wa5.qho77PBKF) >= 0 {
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
		} else {
			s_9Ih_ = !s_9Ih_
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(oxFS5wa5.qho77PBKF, pmEgen2.qho77PBKF)
		}
	}
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && s_9Ih_
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Mul(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {

	if pmEgen2 == oxFS5wa5 {
		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.pJ61taToc(pmEgen2.qho77PBKF)
		uW0eIVjy.hCTOmp0gckSa = false
		return uW0eIVjy
	}
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.d22x6Ygoc80O(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && pmEgen2.hCTOmp0gckSa != oxFS5wa5.hCTOmp0gckSa
	return uW0eIVjy
}




func (uW0eIVjy *ZMtGuo) MulRange(sza5atF, kWa1YrIHZo int64) *ZMtGuo {
	switch {
	case sza5atF > kWa1YrIHZo:
		return  /*line :1*/ uW0eIVjy.SetInt64(1)
	case sza5atF <= 0 && kWa1YrIHZo >= 0:
		return  /*line :1*/ uW0eIVjy.SetInt64(0)
	}

	s_9Ih_ := false
	if sza5atF < 0 {
		s_9Ih_ = (kWa1YrIHZo-sza5atF)&1 == 0
		sza5atF, kWa1YrIHZo = -kWa1YrIHZo, -sza5atF
	}

	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.x648hVuJhff( /*line :1*/ uint64(sza5atF),  /*line :1*/ uint64(kWa1YrIHZo))
	uW0eIVjy.hCTOmp0gckSa = s_9Ih_
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Binomial(h_Wr_yqq, _DfINwXQ int64) *ZMtGuo {
	if _DfINwXQ > h_Wr_yqq {
		return  /*line :1*/ uW0eIVjy.SetInt64(0)
	}

	if _DfINwXQ > h_Wr_yqq-_DfINwXQ {
		_DfINwXQ = h_Wr_yqq - _DfINwXQ
	}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	var DeZVlQzgEfm, Bcrt4YD, aepbxLiWOZ, amoa4Px ZMtGuo
	 /*line :1*/ DeZVlQzgEfm.SetInt64(h_Wr_yqq)
	 /*line :1*/ Bcrt4YD.SetInt64(_DfINwXQ)
	 /*line :1*/ uW0eIVjy.Set(h2Aj5SV9g2re)
	for  /*line :1*/ aepbxLiWOZ.Cmp(&Bcrt4YD) < 0 {
		 /*line :1*/ uW0eIVjy.Mul(uW0eIVjy,  /*line :1*/ amoa4Px.Sub(&DeZVlQzgEfm, &aepbxLiWOZ))
		 /*line :1*/ aepbxLiWOZ.Add(&aepbxLiWOZ, h2Aj5SV9g2re)
		 /*line :1*/ uW0eIVjy.Quo(uW0eIVjy, &aepbxLiWOZ)
	}
	return uW0eIVjy
}




func (uW0eIVjy *ZMtGuo) Quo(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	uW0eIVjy.qho77PBKF, _ =  /*line :1*/ uW0eIVjy.qho77PBKF.xOZxoLyl(nil, pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && pmEgen2.hCTOmp0gckSa != oxFS5wa5.hCTOmp0gckSa
	return uW0eIVjy
}




func (uW0eIVjy *ZMtGuo) Rem(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	_, uW0eIVjy.qho77PBKF =  /*line :1*/ g3X9fa(nil).xOZxoLyl(uW0eIVjy.qho77PBKF, pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && pmEgen2.hCTOmp0gckSa
	return uW0eIVjy
}












func (uW0eIVjy *ZMtGuo) QuoRem(pmEgen2, oxFS5wa5, vFx5p_cm *ZMtGuo) (*ZMtGuo, *ZMtGuo) {
	uW0eIVjy.qho77PBKF, vFx5p_cm.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.xOZxoLyl(vFx5p_cm.qho77PBKF, pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
	uW0eIVjy.hCTOmp0gckSa, vFx5p_cm.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && pmEgen2.hCTOmp0gckSa != oxFS5wa5.hCTOmp0gckSa,  /*line :1*/ len(vFx5p_cm.qho77PBKF) > 0 && pmEgen2.hCTOmp0gckSa
	return uW0eIVjy, vFx5p_cm
}




func (uW0eIVjy *ZMtGuo) Div(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	bb38HsBMh := oxFS5wa5.hCTOmp0gckSa
	var vFx5p_cm ZMtGuo
	 /*line :1*/ uW0eIVjy.QuoRem(pmEgen2, oxFS5wa5, &vFx5p_cm)
	if vFx5p_cm.hCTOmp0gckSa {
		if bb38HsBMh {
			 /*line :1*/ uW0eIVjy.Add(uW0eIVjy, h2Aj5SV9g2re)
		} else {
			 /*line :1*/ uW0eIVjy.Sub(uW0eIVjy, h2Aj5SV9g2re)
		}
	}
	return uW0eIVjy
}




func (uW0eIVjy *ZMtGuo) Mod(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	eYTbpqn := oxFS5wa5
	if uW0eIVjy == oxFS5wa5 ||  /*line :1*/ ap52A4djK(uW0eIVjy.qho77PBKF, oxFS5wa5.qho77PBKF) {
		eYTbpqn =  /*line :1*/ new(ZMtGuo).Set(oxFS5wa5)
	}
	var yL0_p0wUnc ZMtGuo
	 /*line :1*/ yL0_p0wUnc.QuoRem(pmEgen2, oxFS5wa5, uW0eIVjy)
	if uW0eIVjy.hCTOmp0gckSa {
		if eYTbpqn.hCTOmp0gckSa {
			 /*line :1*/ uW0eIVjy.Sub(uW0eIVjy, eYTbpqn)
		} else {
			 /*line :1*/ uW0eIVjy.Add(uW0eIVjy, eYTbpqn)
		}
	}
	return uW0eIVjy
}















func (uW0eIVjy *ZMtGuo) DivMod(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 *ZMtGuo) (*ZMtGuo, *ZMtGuo) {
	eYTbpqn := oxFS5wa5
	if uW0eIVjy == oxFS5wa5 ||  /*line :1*/ ap52A4djK(uW0eIVjy.qho77PBKF, oxFS5wa5.qho77PBKF) {
		eYTbpqn =  /*line :1*/ new(ZMtGuo).Set(oxFS5wa5)
	}
	 /*line :1*/ uW0eIVjy.QuoRem(pmEgen2, oxFS5wa5, gS2TXhHpYaX4)
	if gS2TXhHpYaX4.hCTOmp0gckSa {
		if eYTbpqn.hCTOmp0gckSa {
			 /*line :1*/ uW0eIVjy.Add(uW0eIVjy, h2Aj5SV9g2re)
			 /*line :1*/ gS2TXhHpYaX4.Sub(gS2TXhHpYaX4, eYTbpqn)
		} else {
			 /*line :1*/ uW0eIVjy.Sub(uW0eIVjy, h2Aj5SV9g2re)
			 /*line :1*/ gS2TXhHpYaX4.Add(gS2TXhHpYaX4, eYTbpqn)
		}
	}
	return uW0eIVjy, gS2TXhHpYaX4
}






func (pmEgen2 *ZMtGuo) Cmp(oxFS5wa5 *ZMtGuo) (vFx5p_cm int) {

	switch {
	case pmEgen2 == oxFS5wa5:

	case pmEgen2.hCTOmp0gckSa == oxFS5wa5.hCTOmp0gckSa:
		vFx5p_cm =  /*line :1*/ pmEgen2.qho77PBKF.beaurQHiT(oxFS5wa5.qho77PBKF)
		if pmEgen2.hCTOmp0gckSa {
			vFx5p_cm = -vFx5p_cm
		}
	case pmEgen2.hCTOmp0gckSa:
		vFx5p_cm = -1
	default:
		vFx5p_cm = 1
	}
	return
}






func (pmEgen2 *ZMtGuo) CmpAbs(oxFS5wa5 *ZMtGuo) int {
	return  /*line :1*/ pmEgen2.qho77PBKF.beaurQHiT(oxFS5wa5.qho77PBKF)
}


func xsbmp8(pmEgen2 g3X9fa) uint32 {
	if  /*line :1*/ len(pmEgen2) == 0 {
		return 0
	}
	return  /*line :1*/ uint32(pmEgen2[0])
}


func zA7LNXzkU(pmEgen2 g3X9fa) uint64 {
	if  /*line :1*/ len(pmEgen2) == 0 {
		return 0
	}
	iQKtjmO6 :=  /*line :1*/ uint64(pmEgen2[0])
	if _W == 32 &&  /*line :1*/ len(pmEgen2) > 1 {
		return  /*line :1*/ uint64(pmEgen2[1])<<32 | iQKtjmO6
	}
	return iQKtjmO6
}



func (pmEgen2 *ZMtGuo) Int64() int64 {
	iQKtjmO6 :=  /*line :1*/ int64( /*line :1*/ zA7LNXzkU(pmEgen2.qho77PBKF))
	if pmEgen2.hCTOmp0gckSa {
		iQKtjmO6 = -iQKtjmO6
	}
	return iQKtjmO6
}



func (pmEgen2 *ZMtGuo) Uint64() uint64 {
	return  /*line :1*/ zA7LNXzkU(pmEgen2.qho77PBKF)
}


func (pmEgen2 *ZMtGuo) IsInt64() bool {
	if  /*line :1*/ len(pmEgen2.qho77PBKF) <= 64/_W {
		uKQ8Eox :=  /*line :1*/ int64( /*line :1*/ zA7LNXzkU(pmEgen2.qho77PBKF))
		return uKQ8Eox >= 0 || pmEgen2.hCTOmp0gckSa && uKQ8Eox == -uKQ8Eox
	}
	return false
}


func (pmEgen2 *ZMtGuo) IsUint64() bool {
	return !pmEgen2.hCTOmp0gckSa &&  /*line :1*/ len(pmEgen2.qho77PBKF) <= 64/_W
}



func (pmEgen2 *ZMtGuo) Float64() (float64, IG11Ul7qo) {
	h_Wr_yqq :=  /*line :1*/ pmEgen2.qho77PBKF.wZwJ3Y0()
	if h_Wr_yqq == 0 {
		return 0.0, Exact
	}

	if h_Wr_yqq <= 53 || h_Wr_yqq < 64 && h_Wr_yqq- /*line :1*/ int( /*line :1*/ pmEgen2.qho77PBKF.kDun6ak()) <= 53 {
		sYV3C6v3w1 :=  /*line :1*/ float64( /*line :1*/ zA7LNXzkU(pmEgen2.qho77PBKF))
		if pmEgen2.hCTOmp0gckSa {
			sYV3C6v3w1 = -sYV3C6v3w1
		}
		return sYV3C6v3w1, Exact
	}

	return  /*line :1*/ new(WH8dWN).SetInt(pmEgen2).Float64()
}























func (uW0eIVjy *ZMtGuo) SetString(nR7KU4mGsdy string, weSJS4i4lm int) (*ZMtGuo, bool) {
	return  /*line :1*/ uW0eIVjy.dXT8LebZa3( /*line :1*/ strings.X4yXgtBA(nR7KU4mGsdy), weSJS4i4lm)
}



func (uW0eIVjy *ZMtGuo) dXT8LebZa3(vFx5p_cm io.HSeaoj16Wq0, weSJS4i4lm int) (*ZMtGuo, bool) {
	if _, _, gDonrwI8VP50 :=  /*line :1*/ uW0eIVjy.h9c5N0bF(vFx5p_cm, weSJS4i4lm); gDonrwI8VP50 != nil {
		return nil, false
	}

	if _, gDonrwI8VP50 :=  /*line :1*/ vFx5p_cm.ReadByte(); gDonrwI8VP50 != io.K5Sqco {
		return nil, false
	}
	return uW0eIVjy, true
}



func (uW0eIVjy *ZMtGuo) SetBytes(c3Zu4RY []byte) *ZMtGuo {
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.vbFROX(c3Zu4RY)
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}




func (pmEgen2 *ZMtGuo) Bytes() []byte {

	c3Zu4RY :=  /*line :1*/ make([]byte,  /*line :1*/ len(pmEgen2.qho77PBKF)*_S)
	return c3Zu4RY[ /*line :1*/ pmEgen2.qho77PBKF.qonqPl(c3Zu4RY):]
}





func (pmEgen2 *ZMtGuo) FillBytes(c3Zu4RY []byte) []byte {

	for aepbxLiWOZ := range c3Zu4RY {
		c3Zu4RY[aepbxLiWOZ] = 0
	}
	 /*line :1*/ pmEgen2.qho77PBKF.qonqPl(c3Zu4RY)
	return c3Zu4RY
}



func (pmEgen2 *ZMtGuo) BitLen() int {

	return  /*line :1*/ pmEgen2.qho77PBKF.wZwJ3Y0()
}



func (pmEgen2 *ZMtGuo) TrailingZeroBits() uint {
	return  /*line :1*/ pmEgen2.qho77PBKF.kDun6ak()
}







func (uW0eIVjy *ZMtGuo) Exp(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 *ZMtGuo) *ZMtGuo {
	return  /*line :1*/ uW0eIVjy.xXsI4WC(pmEgen2, oxFS5wa5, gS2TXhHpYaX4, false)
}

func (uW0eIVjy *ZMtGuo) yRs1YhdgFpv9(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 *ZMtGuo) *ZMtGuo {
	return  /*line :1*/ uW0eIVjy.xXsI4WC(pmEgen2, oxFS5wa5, gS2TXhHpYaX4, true)
}

func (uW0eIVjy *ZMtGuo) xXsI4WC(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 *ZMtGuo, asP5KKvg bool) *ZMtGuo {

	oy7c28Q := pmEgen2.qho77PBKF
	if oxFS5wa5.hCTOmp0gckSa {
		if gS2TXhHpYaX4 == nil ||  /*line :1*/ len(gS2TXhHpYaX4.qho77PBKF) == 0 {
			return  /*line :1*/ uW0eIVjy.SetInt64(1)
		}

		ofj80mD_ZYv4 :=  /*line :1*/ new(ZMtGuo).ModInverse(pmEgen2, gS2TXhHpYaX4)
		if ofj80mD_ZYv4 == nil {
			return nil
		}
		oy7c28Q = ofj80mD_ZYv4.qho77PBKF
	}
	jELgrWvPfo := oxFS5wa5.qho77PBKF

	var kvnXqSF g3X9fa
	if gS2TXhHpYaX4 != nil {
		if uW0eIVjy == gS2TXhHpYaX4 ||  /*line :1*/ ap52A4djK(uW0eIVjy.qho77PBKF, gS2TXhHpYaX4.qho77PBKF) {
			gS2TXhHpYaX4 =  /*line :1*/ new(ZMtGuo).Set(gS2TXhHpYaX4)
		}
		kvnXqSF = gS2TXhHpYaX4.qho77PBKF
	}

	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.iGIIsMONlhq(oy7c28Q, jELgrWvPfo, kvnXqSF, asP5KKvg)
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && pmEgen2.hCTOmp0gckSa &&  /*line :1*/ len(jELgrWvPfo) > 0 && jELgrWvPfo[0]&1 == 1
	if uW0eIVjy.hCTOmp0gckSa &&  /*line :1*/ len(kvnXqSF) > 0 {

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(kvnXqSF, uW0eIVjy.qho77PBKF)
		uW0eIVjy.hCTOmp0gckSa = false
	}

	return uW0eIVjy
}












func (uW0eIVjy *ZMtGuo) GCD(pmEgen2, oxFS5wa5, sza5atF, kWa1YrIHZo *ZMtGuo) *ZMtGuo {
	if  /*line :1*/ len(sza5atF.qho77PBKF) == 0 ||  /*line :1*/ len(kWa1YrIHZo.qho77PBKF) == 0 {
		dPkF9KQ, aJnn62dP3L, jKPYMNge2bUB, b3MsdcVUNO :=  /*line :1*/ len(sza5atF.qho77PBKF),  /*line :1*/ len(kWa1YrIHZo.qho77PBKF), sza5atF.hCTOmp0gckSa, kWa1YrIHZo.hCTOmp0gckSa
		if dPkF9KQ == 0 {
			 /*line :1*/ uW0eIVjy.Set(kWa1YrIHZo)
		} else {
			 /*line :1*/ uW0eIVjy.Set(sza5atF)
		}
		uW0eIVjy.hCTOmp0gckSa = false
		if pmEgen2 != nil {
			if dPkF9KQ == 0 {
				 /*line :1*/ pmEgen2.SetUint64(0)
			} else {
				 /*line :1*/ pmEgen2.SetUint64(1)
				pmEgen2.hCTOmp0gckSa = jKPYMNge2bUB
			}
		}
		if oxFS5wa5 != nil {
			if aJnn62dP3L == 0 {
				 /*line :1*/ oxFS5wa5.SetUint64(0)
			} else {
				 /*line :1*/ oxFS5wa5.SetUint64(1)
				oxFS5wa5.hCTOmp0gckSa = b3MsdcVUNO
			}
		}
		return uW0eIVjy
	}

	return  /*line :1*/ uW0eIVjy.pz5vr0wtn(pmEgen2, oxFS5wa5, sza5atF, kWa1YrIHZo)
}













func w2Qys0AUc2(WrhRtgvq, FyXHFes *ZMtGuo) (o2Gpdshq, jn5kUHc, gxC1Eu97dL3, hEO_uDvo VYe6D0, gY4NtDHU2FRO bool) {
	
	var oBH2mFO9, diOe0B, aUnwYdna_OLk, nuTQ64i VYe6D0

	gS2TXhHpYaX4 :=  /*line :1*/ len(FyXHFes.qho77PBKF)
	h_Wr_yqq :=  /*line :1*/ len(WrhRtgvq.qho77PBKF)

	dwiMOUz3iD5U :=  /*line :1*/ bUgMSHLa(WrhRtgvq.qho77PBKF[h_Wr_yqq-1])
	oBH2mFO9 = WrhRtgvq.qho77PBKF[h_Wr_yqq-1]<<dwiMOUz3iD5U | WrhRtgvq.qho77PBKF[h_Wr_yqq-2]>>(_W-dwiMOUz3iD5U)

	switch {
	case h_Wr_yqq == gS2TXhHpYaX4:
		diOe0B = FyXHFes.qho77PBKF[h_Wr_yqq-1]<<dwiMOUz3iD5U | FyXHFes.qho77PBKF[h_Wr_yqq-2]>>(_W-dwiMOUz3iD5U)
	case h_Wr_yqq == gS2TXhHpYaX4+1:
		diOe0B = FyXHFes.qho77PBKF[h_Wr_yqq-2] >> (_W - dwiMOUz3iD5U)
	default:
		diOe0B = 0
	}

	gY4NtDHU2FRO = false

	o2Gpdshq, jn5kUHc, aUnwYdna_OLk = 0, 1, 0
	gxC1Eu97dL3, hEO_uDvo, nuTQ64i = 0, 0, 1

	for diOe0B >= nuTQ64i && oBH2mFO9-diOe0B >= hEO_uDvo+nuTQ64i {
		yL0_p0wUnc, vFx5p_cm := oBH2mFO9/diOe0B, oBH2mFO9%diOe0B
		oBH2mFO9, diOe0B = diOe0B, vFx5p_cm
		o2Gpdshq, jn5kUHc, aUnwYdna_OLk = jn5kUHc, aUnwYdna_OLk, jn5kUHc+yL0_p0wUnc*aUnwYdna_OLk
		gxC1Eu97dL3, hEO_uDvo, nuTQ64i = hEO_uDvo, nuTQ64i, hEO_uDvo+yL0_p0wUnc*nuTQ64i
		gY4NtDHU2FRO = !gY4NtDHU2FRO
	}
	return
}










func fXq3aQas(WrhRtgvq, FyXHFes, yL0_p0wUnc, vFx5p_cm, nR7KU4mGsdy, amoa4Px *ZMtGuo, o2Gpdshq, jn5kUHc, gxC1Eu97dL3, hEO_uDvo VYe6D0, gY4NtDHU2FRO bool) {

	amoa4Px.qho77PBKF =  /*line :1*/ amoa4Px.qho77PBKF.pSzoQ7PMW(o2Gpdshq)
	nR7KU4mGsdy.qho77PBKF =  /*line :1*/ nR7KU4mGsdy.qho77PBKF.pSzoQ7PMW(gxC1Eu97dL3)
	amoa4Px.hCTOmp0gckSa = !gY4NtDHU2FRO
	nR7KU4mGsdy.hCTOmp0gckSa = gY4NtDHU2FRO

	 /*line :1*/ amoa4Px.Mul(WrhRtgvq, amoa4Px)
	 /*line :1*/ nR7KU4mGsdy.Mul(FyXHFes, nR7KU4mGsdy)

	vFx5p_cm.qho77PBKF =  /*line :1*/ vFx5p_cm.qho77PBKF.pSzoQ7PMW(jn5kUHc)
	yL0_p0wUnc.qho77PBKF =  /*line :1*/ yL0_p0wUnc.qho77PBKF.pSzoQ7PMW(hEO_uDvo)
	vFx5p_cm.hCTOmp0gckSa = gY4NtDHU2FRO
	yL0_p0wUnc.hCTOmp0gckSa = !gY4NtDHU2FRO

	 /*line :1*/ vFx5p_cm.Mul(WrhRtgvq, vFx5p_cm)
	 /*line :1*/ yL0_p0wUnc.Mul(FyXHFes, yL0_p0wUnc)

	 /*line :1*/ WrhRtgvq.Add(amoa4Px, nR7KU4mGsdy)
	 /*line :1*/ FyXHFes.Add(vFx5p_cm, yL0_p0wUnc)
}



func mE3PwyI(WrhRtgvq, FyXHFes, EjBupc, DbxTjXzv2a, yL0_p0wUnc, vFx5p_cm, nR7KU4mGsdy, amoa4Px *ZMtGuo, spVrnVVc bool) {
	yL0_p0wUnc, vFx5p_cm =  /*line :1*/ yL0_p0wUnc.QuoRem(WrhRtgvq, FyXHFes, vFx5p_cm)

	*WrhRtgvq, *FyXHFes, *vFx5p_cm = *FyXHFes, *vFx5p_cm, *WrhRtgvq

	if spVrnVVc {

		 /*line :1*/ amoa4Px.Set(DbxTjXzv2a)
		 /*line :1*/ nR7KU4mGsdy.Mul(DbxTjXzv2a, yL0_p0wUnc)
		 /*line :1*/ DbxTjXzv2a.Sub(EjBupc, nR7KU4mGsdy)
		 /*line :1*/ EjBupc.Set(amoa4Px)
	}
}











func (uW0eIVjy *ZMtGuo) pz5vr0wtn(pmEgen2, oxFS5wa5, sza5atF, kWa1YrIHZo *ZMtGuo) *ZMtGuo {
	var WrhRtgvq, FyXHFes, EjBupc, DbxTjXzv2a *ZMtGuo

	WrhRtgvq =  /*line :1*/ new(ZMtGuo).Abs(sza5atF)
	FyXHFes =  /*line :1*/ new(ZMtGuo).Abs(kWa1YrIHZo)

	spVrnVVc := pmEgen2 != nil || oxFS5wa5 != nil

	if spVrnVVc {

		EjBupc =  /*line :1*/ new(ZMtGuo).SetInt64(1)
		DbxTjXzv2a =  /*line :1*/ new(ZMtGuo)
	}

	yL0_p0wUnc :=  /*line :1*/ new(ZMtGuo)
	vFx5p_cm :=  /*line :1*/ new(ZMtGuo)
	nR7KU4mGsdy :=  /*line :1*/ new(ZMtGuo)
	amoa4Px :=  /*line :1*/ new(ZMtGuo)

	if  /*line :1*/ WrhRtgvq.qho77PBKF.beaurQHiT(FyXHFes.qho77PBKF) < 0 {
		WrhRtgvq, FyXHFes = FyXHFes, WrhRtgvq
		DbxTjXzv2a, EjBupc = EjBupc, DbxTjXzv2a
	}

	for  /*line :1*/ len(FyXHFes.qho77PBKF) > 1 {

		o2Gpdshq, jn5kUHc, gxC1Eu97dL3, hEO_uDvo, gY4NtDHU2FRO :=  /*line :1*/ w2Qys0AUc2(WrhRtgvq, FyXHFes)

		if gxC1Eu97dL3 != 0 {

			 /*line :1*/ fXq3aQas(WrhRtgvq, FyXHFes, yL0_p0wUnc, vFx5p_cm, nR7KU4mGsdy, amoa4Px, o2Gpdshq, jn5kUHc, gxC1Eu97dL3, hEO_uDvo, gY4NtDHU2FRO)

			if spVrnVVc {

				 /*line :1*/ fXq3aQas(EjBupc, DbxTjXzv2a, yL0_p0wUnc, vFx5p_cm, nR7KU4mGsdy, amoa4Px, o2Gpdshq, jn5kUHc, gxC1Eu97dL3, hEO_uDvo, gY4NtDHU2FRO)
			}

		} else {

			 /*line :1*/ mE3PwyI(WrhRtgvq, FyXHFes, EjBupc, DbxTjXzv2a, yL0_p0wUnc, vFx5p_cm, nR7KU4mGsdy, amoa4Px, spVrnVVc)
		}
	}

	if  /*line :1*/ len(FyXHFes.qho77PBKF) > 0 {

		if  /*line :1*/ len(WrhRtgvq.qho77PBKF) > 1 {

			 /*line :1*/ mE3PwyI(WrhRtgvq, FyXHFes, EjBupc, DbxTjXzv2a, yL0_p0wUnc, vFx5p_cm, nR7KU4mGsdy, amoa4Px, spVrnVVc)
		}
		if  /*line :1*/ len(FyXHFes.qho77PBKF) > 0 {

			ffPqaw, yFKOqg := WrhRtgvq.qho77PBKF[0], FyXHFes.qho77PBKF[0]
			if spVrnVVc {
				var nlazV3mhQ9DK, t2QKYnlh8sv, eF2p08kXTqsr, hHoD6U1aO VYe6D0
				nlazV3mhQ9DK, t2QKYnlh8sv = 1, 0
				eF2p08kXTqsr, hHoD6U1aO = 0, 1
				gY4NtDHU2FRO := true
				for yFKOqg != 0 {
					yL0_p0wUnc, vFx5p_cm := ffPqaw/yFKOqg, ffPqaw%yFKOqg
					ffPqaw, yFKOqg = yFKOqg, vFx5p_cm
					nlazV3mhQ9DK, t2QKYnlh8sv = t2QKYnlh8sv, nlazV3mhQ9DK+yL0_p0wUnc*t2QKYnlh8sv
					eF2p08kXTqsr, hHoD6U1aO = hHoD6U1aO, eF2p08kXTqsr+yL0_p0wUnc*hHoD6U1aO
					gY4NtDHU2FRO = !gY4NtDHU2FRO
				}

				amoa4Px.qho77PBKF =  /*line :1*/ amoa4Px.qho77PBKF.pSzoQ7PMW(nlazV3mhQ9DK)
				nR7KU4mGsdy.qho77PBKF =  /*line :1*/ nR7KU4mGsdy.qho77PBKF.pSzoQ7PMW(eF2p08kXTqsr)
				amoa4Px.hCTOmp0gckSa = !gY4NtDHU2FRO
				nR7KU4mGsdy.hCTOmp0gckSa = gY4NtDHU2FRO

				 /*line :1*/ amoa4Px.Mul(EjBupc, amoa4Px)
				 /*line :1*/ nR7KU4mGsdy.Mul(DbxTjXzv2a, nR7KU4mGsdy)

				 /*line :1*/ EjBupc.Add(amoa4Px, nR7KU4mGsdy)
			} else {
				for yFKOqg != 0 {
					ffPqaw, yFKOqg = yFKOqg, ffPqaw%yFKOqg
				}
			}
			WrhRtgvq.qho77PBKF[0] = ffPqaw
		}
	}
	jKPYMNge2bUB := sza5atF.hCTOmp0gckSa
	if oxFS5wa5 != nil {

		if oxFS5wa5 == kWa1YrIHZo {
			 /*line :1*/ FyXHFes.Set(kWa1YrIHZo)
		} else {
			FyXHFes = kWa1YrIHZo
		}

		 /*line :1*/ oxFS5wa5.Mul(sza5atF, EjBupc)
		if jKPYMNge2bUB {
			oxFS5wa5.hCTOmp0gckSa = !oxFS5wa5.hCTOmp0gckSa
		}
		 /*line :1*/ oxFS5wa5.Sub(WrhRtgvq, oxFS5wa5)
		 /*line :1*/ oxFS5wa5.Div(oxFS5wa5, FyXHFes)
	}

	if pmEgen2 != nil {
		*pmEgen2 = *EjBupc
		if jKPYMNge2bUB {
			pmEgen2.hCTOmp0gckSa = !pmEgen2.hCTOmp0gckSa
		}
	}

	*uW0eIVjy = *WrhRtgvq

	return uW0eIVjy
}





func (uW0eIVjy *ZMtGuo) Rand(go1ceYeWa *rand.YAxSls, h_Wr_yqq *ZMtGuo) *ZMtGuo {

	if h_Wr_yqq.hCTOmp0gckSa ||  /*line :1*/ len(h_Wr_yqq.qho77PBKF) == 0 {
		uW0eIVjy.hCTOmp0gckSa = false
		uW0eIVjy.qho77PBKF = nil
		return uW0eIVjy
	}
	uW0eIVjy.hCTOmp0gckSa = false
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.calV9Cfo(go1ceYeWa, h_Wr_yqq.qho77PBKF,  /*line :1*/ h_Wr_yqq.qho77PBKF.wZwJ3Y0())
	return uW0eIVjy
}





func (uW0eIVjy *ZMtGuo) ModInverse(fzZBPAu, h_Wr_yqq *ZMtGuo) *ZMtGuo {

	if h_Wr_yqq.hCTOmp0gckSa {
		var fx60ock3 ZMtGuo
		h_Wr_yqq =  /*line :1*/ fx60ock3.Neg(h_Wr_yqq)
	}
	if fzZBPAu.hCTOmp0gckSa {
		var qtUS4uqt1T8 ZMtGuo
		fzZBPAu =  /*line :1*/ qtUS4uqt1T8.Mod(fzZBPAu, h_Wr_yqq)
	}
	var s31g8J, pmEgen2 ZMtGuo
	 /*line :1*/ s31g8J.GCD(&pmEgen2, nil, fzZBPAu, h_Wr_yqq)

	if  /*line :1*/ s31g8J.Cmp(h2Aj5SV9g2re) != 0 {
		return nil
	}

	if pmEgen2.hCTOmp0gckSa {
		 /*line :1*/ uW0eIVjy.Add(&pmEgen2, h_Wr_yqq)
	} else {
		 /*line :1*/ uW0eIVjy.Set(&pmEgen2)
	}
	return uW0eIVjy
}

func (uW0eIVjy g3X9fa) ipjmVbr2qtL(fzZBPAu, h_Wr_yqq g3X9fa) g3X9fa {

	return (& /*line :1*/ ZMtGuo{qho77PBKF: uW0eIVjy}).ModInverse(&ZMtGuo{qho77PBKF: fzZBPAu}, &ZMtGuo{qho77PBKF: h_Wr_yqq}).qho77PBKF
}



func SUxhyDUe(pmEgen2, oxFS5wa5 *ZMtGuo) int {
	if  /*line :1*/ len(oxFS5wa5.qho77PBKF) == 0 || oxFS5wa5.qho77PBKF[0]&1 == 0 {
		 /*line :1*/ panic( /*line :1*/ fmt.IBsS3Oc("big: invalid 2nd argument to Int.Jacobi: need odd integer but got %s",  /*line :1*/ oxFS5wa5.String()))
	}

	var sza5atF, kWa1YrIHZo, vKNCqP0 ZMtGuo
	 /*line :1*/ sza5atF.Set(pmEgen2)
	 /*line :1*/ kWa1YrIHZo.Set(oxFS5wa5)
	g77TAi := 1

	if kWa1YrIHZo.hCTOmp0gckSa {
		if sza5atF.hCTOmp0gckSa {
			g77TAi = -1
		}
		kWa1YrIHZo.hCTOmp0gckSa = false
	}

	for {
		if  /*line :1*/ kWa1YrIHZo.Cmp(h2Aj5SV9g2re) == 0 {
			return g77TAi
		}
		if  /*line :1*/ len(sza5atF.qho77PBKF) == 0 {
			return 0
		}
		 /*line :1*/ sza5atF.Mod(&sza5atF, &kWa1YrIHZo)
		if  /*line :1*/ len(sza5atF.qho77PBKF) == 0 {
			return 0
		}

		nR7KU4mGsdy :=  /*line :1*/ sza5atF.qho77PBKF.kDun6ak()
		if nR7KU4mGsdy&1 != 0 {
			dQ4MtZ0eU := kWa1YrIHZo.qho77PBKF[0] & 7
			if dQ4MtZ0eU == 3 || dQ4MtZ0eU == 5 {
				g77TAi = -g77TAi
			}
		}
		 /*line :1*/ vKNCqP0.Rsh(&sza5atF, nR7KU4mGsdy)

		if kWa1YrIHZo.qho77PBKF[0]&3 == 3 && vKNCqP0.qho77PBKF[0]&3 == 3 {
			g77TAi = -g77TAi
		}
		 /*line :1*/ sza5atF.Set(&kWa1YrIHZo)
		 /*line :1*/ kWa1YrIHZo.Set(&vKNCqP0)
	}
}









func (uW0eIVjy *ZMtGuo) cTuKaOO6(pmEgen2, djOsNJsLfH5U *ZMtGuo) *ZMtGuo {
	vRGCE256ba4X :=  /*line :1*/ new(ZMtGuo).Add(djOsNJsLfH5U, h2Aj5SV9g2re)
	 /*line :1*/ vRGCE256ba4X.Rsh(vRGCE256ba4X, 2)
	 /*line :1*/ uW0eIVjy.Exp(pmEgen2, vRGCE256ba4X, djOsNJsLfH5U)
	return uW0eIVjy
}









func (uW0eIVjy *ZMtGuo) zAl0ykkpT2sm(pmEgen2, djOsNJsLfH5U *ZMtGuo) *ZMtGuo {

	vRGCE256ba4X :=  /*line :1*/ new(ZMtGuo).Rsh(djOsNJsLfH5U, 3)
	cA74uCap_Vd :=  /*line :1*/ new(ZMtGuo).Lsh(pmEgen2, 1)
	vX4qPGnGtz :=  /*line :1*/ new(ZMtGuo).Exp(cA74uCap_Vd, vRGCE256ba4X, djOsNJsLfH5U)
	jp8w6aKV6ti :=  /*line :1*/ new(ZMtGuo).Mul(vX4qPGnGtz, vX4qPGnGtz)
	 /*line :1*/ jp8w6aKV6ti.Mod(jp8w6aKV6ti, djOsNJsLfH5U)
	 /*line :1*/ jp8w6aKV6ti.Mul(jp8w6aKV6ti, cA74uCap_Vd)
	 /*line :1*/ jp8w6aKV6ti.Mod(jp8w6aKV6ti, djOsNJsLfH5U)
	 /*line :1*/ jp8w6aKV6ti.Sub(jp8w6aKV6ti, h2Aj5SV9g2re)
	 /*line :1*/ jp8w6aKV6ti.Mul(jp8w6aKV6ti, pmEgen2)
	 /*line :1*/ jp8w6aKV6ti.Mod(jp8w6aKV6ti, djOsNJsLfH5U)
	 /*line :1*/ jp8w6aKV6ti.Mul(jp8w6aKV6ti, vX4qPGnGtz)
	 /*line :1*/ uW0eIVjy.Mod(jp8w6aKV6ti, djOsNJsLfH5U)
	return uW0eIVjy
}



func (uW0eIVjy *ZMtGuo) hYykOg2OYiP(pmEgen2, djOsNJsLfH5U *ZMtGuo) *ZMtGuo {
	
	var nR7KU4mGsdy ZMtGuo
	 /*line :1*/ nR7KU4mGsdy.Sub(djOsNJsLfH5U, h2Aj5SV9g2re)
	vRGCE256ba4X :=  /*line :1*/ nR7KU4mGsdy.qho77PBKF.kDun6ak()
	 /*line :1*/ nR7KU4mGsdy.Rsh(&nR7KU4mGsdy, vRGCE256ba4X)

	
	var h_Wr_yqq ZMtGuo
	 /*line :1*/ h_Wr_yqq.SetInt64(2)
	for  /*line :1*/ SUxhyDUe(&h_Wr_yqq, djOsNJsLfH5U) != -1 {
		 /*line :1*/ h_Wr_yqq.Add(&h_Wr_yqq, h2Aj5SV9g2re)
	}

	
	
	
	
	var oxFS5wa5, kWa1YrIHZo, fzZBPAu, amoa4Px ZMtGuo
	 /*line :1*/ oxFS5wa5.Add(&nR7KU4mGsdy, h2Aj5SV9g2re)
	 /*line :1*/ oxFS5wa5.Rsh(&oxFS5wa5, 1)
	 /*line :1*/ oxFS5wa5.Exp(pmEgen2, &oxFS5wa5, djOsNJsLfH5U)
	 /*line :1*/ kWa1YrIHZo.Exp(pmEgen2, &nR7KU4mGsdy, djOsNJsLfH5U)
	 /*line :1*/ fzZBPAu.Exp(&h_Wr_yqq, &nR7KU4mGsdy, djOsNJsLfH5U)
	vFx5p_cm := vRGCE256ba4X
	for {
		
		var gS2TXhHpYaX4 uint
		 /*line :1*/ amoa4Px.Set(&kWa1YrIHZo)
		for  /*line :1*/ amoa4Px.Cmp(h2Aj5SV9g2re) != 0 {
			 /*line :1*/ amoa4Px.Mul(&amoa4Px, &amoa4Px).Mod(&amoa4Px, djOsNJsLfH5U)
			gS2TXhHpYaX4++
		}

		if gS2TXhHpYaX4 == 0 {
			return  /*line :1*/ uW0eIVjy.Set(&oxFS5wa5)
		}

		 /*line :1*/ amoa4Px.SetInt64(0).SetBit(&amoa4Px,  /*line :1*/ int(vFx5p_cm-gS2TXhHpYaX4-1), 1).Exp(&fzZBPAu, &amoa4Px, djOsNJsLfH5U)

		 /*line :1*/ fzZBPAu.Mul(&amoa4Px, &amoa4Px).Mod(&fzZBPAu, djOsNJsLfH5U)
		 /*line :1*/ oxFS5wa5.Mul(&oxFS5wa5, &amoa4Px).Mod(&oxFS5wa5, djOsNJsLfH5U)
		 /*line :1*/ kWa1YrIHZo.Mul(&kWa1YrIHZo, &fzZBPAu).Mod(&kWa1YrIHZo, djOsNJsLfH5U)
		vFx5p_cm = gS2TXhHpYaX4
	}
}





func (uW0eIVjy *ZMtGuo) ModSqrt(pmEgen2, djOsNJsLfH5U *ZMtGuo) *ZMtGuo {
	switch  /*line :1*/ SUxhyDUe(pmEgen2, djOsNJsLfH5U) {
	case -1:
		return nil
	case 0:
		return  /*line :1*/ uW0eIVjy.SetInt64(0)
	case 1:
		break
	}
	if pmEgen2.hCTOmp0gckSa ||  /*line :1*/ pmEgen2.Cmp(djOsNJsLfH5U) >= 0 {
		pmEgen2 =  /*line :1*/ new(ZMtGuo).Mod(pmEgen2, djOsNJsLfH5U)
	}

	switch {
	case djOsNJsLfH5U.qho77PBKF[0]%4 == 3:

		return  /*line :1*/ uW0eIVjy.cTuKaOO6(pmEgen2, djOsNJsLfH5U)
	case djOsNJsLfH5U.qho77PBKF[0]%8 == 5:

		return  /*line :1*/ uW0eIVjy.zAl0ykkpT2sm(pmEgen2, djOsNJsLfH5U)
	default:

		return  /*line :1*/ uW0eIVjy.hYykOg2OYiP(pmEgen2, djOsNJsLfH5U)
	}
}


func (uW0eIVjy *ZMtGuo) Lsh(pmEgen2 *ZMtGuo, h_Wr_yqq uint) *ZMtGuo {
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.z839tk6CmDHT(pmEgen2.qho77PBKF, h_Wr_yqq)
	uW0eIVjy.hCTOmp0gckSa = pmEgen2.hCTOmp0gckSa
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Rsh(pmEgen2 *ZMtGuo, h_Wr_yqq uint) *ZMtGuo {
	if pmEgen2.hCTOmp0gckSa {

		amoa4Px :=  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
		amoa4Px =  /*line :1*/ amoa4Px.qzOaHmR(amoa4Px, h_Wr_yqq)
		uW0eIVjy.qho77PBKF =  /*line :1*/ amoa4Px.jXv7nnhY_(amoa4Px, lG0t2KfqGwyP)
		uW0eIVjy.hCTOmp0gckSa = true
		return uW0eIVjy
	}

	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.qzOaHmR(pmEgen2.qho77PBKF, h_Wr_yqq)
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}



func (pmEgen2 *ZMtGuo) Bit(aepbxLiWOZ int) uint {
	if aepbxLiWOZ == 0 {

		if  /*line :1*/ len(pmEgen2.qho77PBKF) > 0 {
			return  /*line :1*/ uint(pmEgen2.qho77PBKF[0] & 1)
		}
		return 0
	}
	if aepbxLiWOZ < 0 {
		 /*line :1*/ panic("negative bit index")
	}
	if pmEgen2.hCTOmp0gckSa {
		amoa4Px :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
		return  /*line :1*/ amoa4Px.bjRK8rI0NMn( /*line :1*/ uint(aepbxLiWOZ)) ^ 1
	}

	return  /*line :1*/ pmEgen2.qho77PBKF.bjRK8rI0NMn( /*line :1*/ uint(aepbxLiWOZ))
}





func (uW0eIVjy *ZMtGuo) SetBit(pmEgen2 *ZMtGuo, aepbxLiWOZ int, kWa1YrIHZo uint) *ZMtGuo {
	if aepbxLiWOZ < 0 {
		 /*line :1*/ panic("negative bit index")
	}
	if pmEgen2.hCTOmp0gckSa {
		amoa4Px :=  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
		amoa4Px =  /*line :1*/ amoa4Px.hUdaA0sgaAF(amoa4Px,  /*line :1*/ uint(aepbxLiWOZ), kWa1YrIHZo^1)
		uW0eIVjy.qho77PBKF =  /*line :1*/ amoa4Px.jXv7nnhY_(amoa4Px, lG0t2KfqGwyP)
		uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0
		return uW0eIVjy
	}
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.hUdaA0sgaAF(pmEgen2.qho77PBKF,  /*line :1*/ uint(aepbxLiWOZ), kWa1YrIHZo)
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) And(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	if pmEgen2.hCTOmp0gckSa == oxFS5wa5.hCTOmp0gckSa {
		if pmEgen2.hCTOmp0gckSa {

			dKR0NUYWi :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
			ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_( /*line :1*/ uW0eIVjy.qho77PBKF.ycPkQaVsS(dKR0NUYWi, ihENsIZD), lG0t2KfqGwyP)
			uW0eIVjy.hCTOmp0gckSa = true
			return uW0eIVjy
		}

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.aqWAA8J(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
		uW0eIVjy.hCTOmp0gckSa = false
		return uW0eIVjy
	}

	if pmEgen2.hCTOmp0gckSa {
		pmEgen2, oxFS5wa5 = oxFS5wa5, pmEgen2
	}

	ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.csmID38W(pmEgen2.qho77PBKF, ihENsIZD)
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) AndNot(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	if pmEgen2.hCTOmp0gckSa == oxFS5wa5.hCTOmp0gckSa {
		if pmEgen2.hCTOmp0gckSa {

			dKR0NUYWi :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
			ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.csmID38W(ihENsIZD, dKR0NUYWi)
			uW0eIVjy.hCTOmp0gckSa = false
			return uW0eIVjy
		}

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.csmID38W(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
		uW0eIVjy.hCTOmp0gckSa = false
		return uW0eIVjy
	}

	if pmEgen2.hCTOmp0gckSa {

		dKR0NUYWi :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_( /*line :1*/ uW0eIVjy.qho77PBKF.ycPkQaVsS(dKR0NUYWi, oxFS5wa5.qho77PBKF), lG0t2KfqGwyP)
		uW0eIVjy.hCTOmp0gckSa = true
		return uW0eIVjy
	}

	ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.aqWAA8J(pmEgen2.qho77PBKF, ihENsIZD)
	uW0eIVjy.hCTOmp0gckSa = false
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Or(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	if pmEgen2.hCTOmp0gckSa == oxFS5wa5.hCTOmp0gckSa {
		if pmEgen2.hCTOmp0gckSa {

			dKR0NUYWi :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
			ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_( /*line :1*/ uW0eIVjy.qho77PBKF.aqWAA8J(dKR0NUYWi, ihENsIZD), lG0t2KfqGwyP)
			uW0eIVjy.hCTOmp0gckSa = true
			return uW0eIVjy
		}

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.ycPkQaVsS(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
		uW0eIVjy.hCTOmp0gckSa = false
		return uW0eIVjy
	}

	if pmEgen2.hCTOmp0gckSa {
		pmEgen2, oxFS5wa5 = oxFS5wa5, pmEgen2
	}

	ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_( /*line :1*/ uW0eIVjy.qho77PBKF.csmID38W(ihENsIZD, pmEgen2.qho77PBKF), lG0t2KfqGwyP)
	uW0eIVjy.hCTOmp0gckSa = true
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Xor(pmEgen2, oxFS5wa5 *ZMtGuo) *ZMtGuo {
	if pmEgen2.hCTOmp0gckSa == oxFS5wa5.hCTOmp0gckSa {
		if pmEgen2.hCTOmp0gckSa {

			dKR0NUYWi :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
			ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.ubcVkBmt(dKR0NUYWi, ihENsIZD)
			uW0eIVjy.hCTOmp0gckSa = false
			return uW0eIVjy
		}

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.ubcVkBmt(pmEgen2.qho77PBKF, oxFS5wa5.qho77PBKF)
		uW0eIVjy.hCTOmp0gckSa = false
		return uW0eIVjy
	}

	if pmEgen2.hCTOmp0gckSa {
		pmEgen2, oxFS5wa5 = oxFS5wa5, pmEgen2
	}

	ihENsIZD :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(oxFS5wa5.qho77PBKF, lG0t2KfqGwyP)
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_( /*line :1*/ uW0eIVjy.qho77PBKF.ubcVkBmt(pmEgen2.qho77PBKF, ihENsIZD), lG0t2KfqGwyP)
	uW0eIVjy.hCTOmp0gckSa = true
	return uW0eIVjy
}


func (uW0eIVjy *ZMtGuo) Not(pmEgen2 *ZMtGuo) *ZMtGuo {
	if pmEgen2.hCTOmp0gckSa {

		uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.lt4VKILNCo(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
		uW0eIVjy.hCTOmp0gckSa = false
		return uW0eIVjy
	}

	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.jXv7nnhY_(pmEgen2.qho77PBKF, lG0t2KfqGwyP)
	uW0eIVjy.hCTOmp0gckSa = true
	return uW0eIVjy
}



func (uW0eIVjy *ZMtGuo) Sqrt(pmEgen2 *ZMtGuo) *ZMtGuo {
	if pmEgen2.hCTOmp0gckSa {
		 /*line :1*/ panic("square root of negative number")
	}
	uW0eIVjy.hCTOmp0gckSa = false
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.vf9TNsbnLN(pmEgen2.qho77PBKF)
	return uW0eIVjy
}
