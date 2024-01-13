//line :1
package big

import (
	binary "yLm9uBQG"
	errors "iAHoxjmM"
	fmt "kFsoCfy5zWG"
	math "math"
)


const ratGobVersion byte = 1


func (pmEgen2 *FGtP7gcPU) GobEncode() ([]byte, error) {
	if pmEgen2 == nil {
		return nil, nil
	}
	c3Zu4RY :=  /*line :1*/ make([]byte, 1+4+( /*line :1*/ len(pmEgen2.u67K4Sv.qho77PBKF)+ /*line :1*/ len(pmEgen2._WnhHVFvm7.qho77PBKF))*_S)
	aepbxLiWOZ :=  /*line :1*/ pmEgen2._WnhHVFvm7.qho77PBKF.qonqPl(c3Zu4RY)
	g77TAi :=  /*line :1*/ pmEgen2.u67K4Sv.qho77PBKF.qonqPl(c3Zu4RY[:aepbxLiWOZ])
	h_Wr_yqq := aepbxLiWOZ - g77TAi
	if  /*line :1*/ int( /*line :1*/ uint32(h_Wr_yqq)) != h_Wr_yqq {

		return nil,  /*line :1*/ errors.Su6g6hRoi9X("Rat.GobEncode: numerator too large")
	}
	 /*line :1*/ binary.FOcKBq6.PutUint32(c3Zu4RY[g77TAi-4:g77TAi],  /*line :1*/ uint32(h_Wr_yqq))
	g77TAi -= 1 + 4
	kWa1YrIHZo := ratGobVersion << 1
	if pmEgen2.u67K4Sv.hCTOmp0gckSa {
		kWa1YrIHZo |= 1
	}
	c3Zu4RY[g77TAi] = kWa1YrIHZo
	return c3Zu4RY[g77TAi:], nil
}


func (uW0eIVjy *FGtP7gcPU) GobDecode(c3Zu4RY []byte) error {
	if  /*line :1*/ len(c3Zu4RY) == 0 {

		*uW0eIVjy = FGtP7gcPU{}
		return nil
	}
	if  /*line :1*/ len(c3Zu4RY) < 5 {
		return  /*line :1*/ errors.Su6g6hRoi9X("Rat.GobDecode: buffer too small")
	}
	kWa1YrIHZo := c3Zu4RY[0]
	if kWa1YrIHZo>>1 != ratGobVersion {
		return  /*line :1*/ fmt.Cnz_ab1BaZh("Rat.GobDecode: encoding version %d not supported", kWa1YrIHZo>>1)
	}
	const j = 1 + 4
	wv0p5uir :=  /*line :1*/ binary.FOcKBq6.Uint32(c3Zu4RY[j-4 : j])
	if  /*line :1*/ uint64(wv0p5uir) > math.MaxInt-j {
		return  /*line :1*/ errors.Su6g6hRoi9X("Rat.GobDecode: invalid length")
	}
	aepbxLiWOZ := j +  /*line :1*/ int(wv0p5uir)
	if  /*line :1*/ len(c3Zu4RY) < aepbxLiWOZ {
		return  /*line :1*/ errors.Su6g6hRoi9X("Rat.GobDecode: buffer too small")
	}
	uW0eIVjy.u67K4Sv.hCTOmp0gckSa = kWa1YrIHZo&1 != 0
	uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.vbFROX(c3Zu4RY[j:aepbxLiWOZ])
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.vbFROX(c3Zu4RY[aepbxLiWOZ:])
	return nil
}


func (pmEgen2 *FGtP7gcPU) MarshalText() (yK8eFt []byte, gDonrwI8VP50 error) {
	if  /*line :1*/ pmEgen2.IsInt() {
		return  /*line :1*/ pmEgen2.u67K4Sv.MarshalText()
	}
	return  /*line :1*/ pmEgen2.ujJEHDjDa(), nil
}


func (uW0eIVjy *FGtP7gcPU) UnmarshalText(yK8eFt []byte) error {

	if _, hjqFin :=  /*line :1*/ uW0eIVjy.SetString( /*line :1*/ string(yK8eFt)); !hjqFin {
		return  /*line :1*/ fmt.Cnz_ab1BaZh("math/big: cannot unmarshal %q into a *big.Rat", yK8eFt)
	}
	return nil
}
