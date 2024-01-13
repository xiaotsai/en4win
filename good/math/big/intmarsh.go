//line :1
package big

import (
	bytes "wLlfRPv"
	fmt "kFsoCfy5zWG"
)


const intGobVersion byte = 1


func (pmEgen2 *ZMtGuo) GobEncode() ([]byte, error) {
	if pmEgen2 == nil {
		return nil, nil
	}
	c3Zu4RY :=  /*line :1*/ make([]byte, 1+ /*line :1*/ len(pmEgen2.qho77PBKF)*_S)
	aepbxLiWOZ :=  /*line :1*/ pmEgen2.qho77PBKF.qonqPl(c3Zu4RY) - 1
	kWa1YrIHZo := intGobVersion << 1
	if pmEgen2.hCTOmp0gckSa {
		kWa1YrIHZo |= 1
	}
	c3Zu4RY[aepbxLiWOZ] = kWa1YrIHZo
	return c3Zu4RY[aepbxLiWOZ:], nil
}


func (uW0eIVjy *ZMtGuo) GobDecode(c3Zu4RY []byte) error {
	if  /*line :1*/ len(c3Zu4RY) == 0 {

		*uW0eIVjy = ZMtGuo{}
		return nil
	}
	kWa1YrIHZo := c3Zu4RY[0]
	if kWa1YrIHZo>>1 != intGobVersion {
		return  /*line :1*/ fmt.Cnz_ab1BaZh("Int.GobDecode: encoding version %d not supported", kWa1YrIHZo>>1)
	}
	uW0eIVjy.hCTOmp0gckSa = kWa1YrIHZo&1 != 0
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.vbFROX(c3Zu4RY[1:])
	return nil
}


func (pmEgen2 *ZMtGuo) MarshalText() (yK8eFt []byte, gDonrwI8VP50 error) {
	if pmEgen2 == nil {
		return [] /*line :1*/ byte("<nil>"), nil
	}
	return  /*line :1*/ pmEgen2.qho77PBKF.sgVtOtra(pmEgen2.hCTOmp0gckSa, 10), nil
}


func (uW0eIVjy *ZMtGuo) UnmarshalText(yK8eFt []byte) error {
	if _, hjqFin :=  /*line :1*/ uW0eIVjy.dXT8LebZa3( /*line :1*/ bytes.FGmhhHol(yK8eFt), 0); !hjqFin {
		return  /*line :1*/ fmt.Cnz_ab1BaZh("math/big: cannot unmarshal %q into a *big.Int", yK8eFt)
	}
	return nil
}


func (pmEgen2 *ZMtGuo) MarshalJSON() ([]byte, error) {
	if pmEgen2 == nil {
		return [] /*line :1*/ byte("null"), nil
	}
	return  /*line :1*/ pmEgen2.qho77PBKF.sgVtOtra(pmEgen2.hCTOmp0gckSa, 10), nil
}


func (uW0eIVjy *ZMtGuo) UnmarshalJSON(yK8eFt []byte) error {

	if  /*line :1*/ string(yK8eFt) == "null" {
		return nil
	}
	return  /*line :1*/ uW0eIVjy.UnmarshalText(yK8eFt)
}
