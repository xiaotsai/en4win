//line :1
package big

import (
	errors "iAHoxjmM"
	fmt "kFsoCfy5zWG"
	io "xI9ai2djaJ2"
	strconv "zBESu2SrRjP"
	strings "fQXlzv"
)

func y4FzoGss(qvA6MnWV rune) bool {
	return  /*line :1*/ strings.MLBN4mdkQa3("+-/0123456789.eE", qvA6MnWV)
}

var zE9gB5WY FGtP7gcPU
var _ fmt.MqdnB6 = &zE9gB5WY	



func (uW0eIVjy *FGtP7gcPU) Scan(nR7KU4mGsdy fmt.XDnSs0f3, qvA6MnWV rune) error {
	piXXeOWJa8, gDonrwI8VP50 :=  /*line :1*/ nR7KU4mGsdy.Token(true, y4FzoGss)
	if gDonrwI8VP50 != nil {
		return gDonrwI8VP50
	}
	if ! /*line :1*/ strings.MLBN4mdkQa3("efgEFGv", qvA6MnWV) {
		return  /*line :1*/ errors.Su6g6hRoi9X("Rat.Scan: invalid verb")
	}
	if _, hjqFin :=  /*line :1*/ uW0eIVjy.SetString( /*line :1*/ string(piXXeOWJa8)); !hjqFin {
		return  /*line :1*/ errors.Su6g6hRoi9X("Rat.Scan: invalid syntax")
	}
	return nil
}



















func (uW0eIVjy *FGtP7gcPU) SetString(nR7KU4mGsdy string) (*FGtP7gcPU, bool) {
	if  /*line :1*/ len(nR7KU4mGsdy) == 0 {
		return nil, false
	}

	if iCu98ZLbG9 :=  /*line :1*/ strings.FCsEk_zxUIUY(nR7KU4mGsdy, "/"); iCu98ZLbG9 >= 0 {
		if _, hjqFin :=  /*line :1*/ uW0eIVjy.u67K4Sv.SetString(nR7KU4mGsdy[:iCu98ZLbG9], 0); !hjqFin {
			return nil, false
		}
		vFx5p_cm :=  /*line :1*/ strings.X4yXgtBA(nR7KU4mGsdy[iCu98ZLbG9+1:])
		var gDonrwI8VP50 error
		if uW0eIVjy._WnhHVFvm7.qho77PBKF, _, _, gDonrwI8VP50 =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.h9c5N0bF(vFx5p_cm, 0, false); gDonrwI8VP50 != nil {
			return nil, false
		}

		if _, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte(); gDonrwI8VP50 != io.K5Sqco {
			return nil, false
		}
		if  /*line :1*/ len(uW0eIVjy._WnhHVFvm7.qho77PBKF) == 0 {
			return nil, false
		}
		return  /*line :1*/ uW0eIVjy.hitXcy5(), true
	}

	vFx5p_cm :=  /*line :1*/ strings.X4yXgtBA(nR7KU4mGsdy)

	s_9Ih_, gDonrwI8VP50 :=  /*line :1*/ xOqvCaQwvh(vFx5p_cm)
	if gDonrwI8VP50 != nil {
		return nil, false
	}

	
	var weSJS4i4lm int
	var kVeBoojLX int	
	uW0eIVjy.u67K4Sv.qho77PBKF, weSJS4i4lm, kVeBoojLX, gDonrwI8VP50 =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.h9c5N0bF(vFx5p_cm, 0, true)
	if gDonrwI8VP50 != nil {
		return nil, false
	}

	
	var xXsI4WC int64
	var jlDe5QeBx4 int
	xXsI4WC, jlDe5QeBx4, gDonrwI8VP50 =  /*line :1*/ gCqU9SyEYV(vFx5p_cm, true, true)
	if gDonrwI8VP50 != nil {
		return nil, false
	}

	if _, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte(); gDonrwI8VP50 != io.K5Sqco {
		return nil, false
	}

	if  /*line :1*/ len(uW0eIVjy.u67K4Sv.qho77PBKF) == 0 {
		return  /*line :1*/ uW0eIVjy.hitXcy5(), true
	}

	
	var ovNyTM, eXfVq2GG int64
	if kVeBoojLX < 0 {

		s31g8J :=  /*line :1*/ int64(kVeBoojLX)
		switch weSJS4i4lm {
		case 10:
			eXfVq2GG = s31g8J
			fallthrough
		case 2:
			ovNyTM = s31g8J
		case 8:
			ovNyTM = s31g8J * 3
		case 16:
			ovNyTM = s31g8J * 4
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

	if eXfVq2GG != 0 {
		h_Wr_yqq := eXfVq2GG
		if h_Wr_yqq < 0 {
			h_Wr_yqq = -h_Wr_yqq
			if h_Wr_yqq < 0 {

				return nil, false
			}
		}
		if h_Wr_yqq > 1e6 {
			return nil, false
		}
		wwgZlyx :=  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.iGIIsMONlhq(m5xi9O,  /*line :1*/ g3X9fa(nil).pSzoQ7PMW( /*line :1*/ VYe6D0(h_Wr_yqq)), nil, false)
		if eXfVq2GG > 0 {
			uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.d22x6Ygoc80O(uW0eIVjy.u67K4Sv.qho77PBKF, wwgZlyx)
			uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
		} else {
			uW0eIVjy._WnhHVFvm7.qho77PBKF = wwgZlyx
		}
	} else {
		uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
	}

	if ovNyTM < -1e7 || ovNyTM > 1e7 {
		return nil, false
	}
	if ovNyTM > 0 {
		uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.z839tk6CmDHT(uW0eIVjy.u67K4Sv.qho77PBKF,  /*line :1*/ uint(ovNyTM))
	} else if ovNyTM < 0 {
		uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.z839tk6CmDHT(uW0eIVjy._WnhHVFvm7.qho77PBKF,  /*line :1*/ uint(-ovNyTM))
	}

	uW0eIVjy.u67K4Sv.hCTOmp0gckSa = s_9Ih_ &&  /*line :1*/ len(uW0eIVjy.u67K4Sv.qho77PBKF) > 0

	return  /*line :1*/ uW0eIVjy.hitXcy5(), true
}

















func gCqU9SyEYV(vFx5p_cm io.HSeaoj16Wq0, afB2TWk, wA6KfAX2C04 bool) (xXsI4WC int64, weSJS4i4lm int, gDonrwI8VP50 error) {

	qvA6MnWV, gDonrwI8VP50 :=  /*line :1*/ vFx5p_cm.ReadByte()
	if gDonrwI8VP50 != nil {
		if gDonrwI8VP50 == io.K5Sqco {
			gDonrwI8VP50 = nil
		}
		return 0, 10, gDonrwI8VP50
	}

	switch qvA6MnWV {
	case 'e', 'E':
		weSJS4i4lm = 10
	case 'p', 'P':
		if afB2TWk {
			weSJS4i4lm = 2
			break
		}
		fallthrough
	default:
		 /*line :1*/ vFx5p_cm.UnreadByte()
		return 0, 10, nil
	}

	
	var cmZK_aGSDT []byte
	qvA6MnWV, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte()
	if gDonrwI8VP50 == nil && (qvA6MnWV == '+' || qvA6MnWV == '-') {
		if qvA6MnWV == '-' {
			cmZK_aGSDT =  /*line :1*/ append(cmZK_aGSDT, '-')
		}
		qvA6MnWV, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte()
	}

	iCClg09zMOt := '.'
	j9gIA7FO := false

	d1n19p := false
	for gDonrwI8VP50 == nil {
		if '0' <= qvA6MnWV && qvA6MnWV <= '9' {
			cmZK_aGSDT =  /*line :1*/ append(cmZK_aGSDT, qvA6MnWV)
			iCClg09zMOt = '0'
			d1n19p = true
		} else if qvA6MnWV == '_' && wA6KfAX2C04 {
			if iCClg09zMOt != '0' {
				j9gIA7FO = true
			}
			iCClg09zMOt = '_'
		} else {
			 /*line :1*/ vFx5p_cm.UnreadByte()
			break
		}
		qvA6MnWV, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte()
	}

	if gDonrwI8VP50 == io.K5Sqco {
		gDonrwI8VP50 = nil
	}
	if gDonrwI8VP50 == nil && !d1n19p {
		gDonrwI8VP50 = tdXYCwgkg
	}
	if gDonrwI8VP50 == nil {
		xXsI4WC, gDonrwI8VP50 =  /*line :1*/ strconv.N8Xpap1Vt1ot( /*line :1*/ string(cmZK_aGSDT), 10, 64)
	}

	if gDonrwI8VP50 == nil && (j9gIA7FO || iCClg09zMOt == '_') {
		gDonrwI8VP50 = l9euxO
	}

	return
}


func (pmEgen2 *FGtP7gcPU) String() string {
	return  /*line :1*/ string( /*line :1*/ pmEgen2.ujJEHDjDa())
}


func (pmEgen2 *FGtP7gcPU) ujJEHDjDa() []byte {
	var c3Zu4RY []byte
	c3Zu4RY =  /*line :1*/ pmEgen2.u67K4Sv.Append(c3Zu4RY, 10)
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '/')
	if  /*line :1*/ len(pmEgen2._WnhHVFvm7.qho77PBKF) != 0 {
		c3Zu4RY =  /*line :1*/ pmEgen2._WnhHVFvm7.Append(c3Zu4RY, 10)
	} else {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '1')
	}
	return c3Zu4RY
}



func (pmEgen2 *FGtP7gcPU) RatString() string {
	if  /*line :1*/ pmEgen2.IsInt() {
		return  /*line :1*/ pmEgen2.u67K4Sv.String()
	}
	return  /*line :1*/ pmEgen2.String()
}




func (pmEgen2 *FGtP7gcPU) FloatString(cZUDpnJQn int) string {
	var c3Zu4RY []byte

	if  /*line :1*/ pmEgen2.IsInt() {
		c3Zu4RY =  /*line :1*/ pmEgen2.u67K4Sv.Append(c3Zu4RY, 10)
		if cZUDpnJQn > 0 {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '.')
			for aepbxLiWOZ := cZUDpnJQn; aepbxLiWOZ > 0; aepbxLiWOZ-- {
				c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
			}
		}
		return  /*line :1*/ string(c3Zu4RY)
	}

	yL0_p0wUnc, vFx5p_cm :=  /*line :1*/ g3X9fa(nil).xOZxoLyl( /*line :1*/ g3X9fa(nil), pmEgen2.u67K4Sv.qho77PBKF, pmEgen2._WnhHVFvm7.qho77PBKF)

	djOsNJsLfH5U := lG0t2KfqGwyP
	if cZUDpnJQn > 0 {
		djOsNJsLfH5U =  /*line :1*/ g3X9fa(nil).iGIIsMONlhq(syj6Abagk,  /*line :1*/ g3X9fa(nil).glzD00sysLql( /*line :1*/ uint64(cZUDpnJQn)), nil, false)
	}

	vFx5p_cm =  /*line :1*/ vFx5p_cm.d22x6Ygoc80O(vFx5p_cm, djOsNJsLfH5U)
	vFx5p_cm, jt6wZjdRTx :=  /*line :1*/ vFx5p_cm.xOZxoLyl( /*line :1*/ g3X9fa(nil), vFx5p_cm, pmEgen2._WnhHVFvm7.qho77PBKF)

	jt6wZjdRTx =  /*line :1*/ jt6wZjdRTx.jXv7nnhY_(jt6wZjdRTx, jt6wZjdRTx)
	if  /*line :1*/ pmEgen2._WnhHVFvm7.qho77PBKF.beaurQHiT(jt6wZjdRTx) <= 0 {
		vFx5p_cm =  /*line :1*/ vFx5p_cm.jXv7nnhY_(vFx5p_cm, lG0t2KfqGwyP)
		if  /*line :1*/ vFx5p_cm.beaurQHiT(djOsNJsLfH5U) >= 0 {
			yL0_p0wUnc =  /*line :1*/ g3X9fa(nil).jXv7nnhY_(yL0_p0wUnc, lG0t2KfqGwyP)
			vFx5p_cm =  /*line :1*/ g3X9fa(nil).lt4VKILNCo(vFx5p_cm, djOsNJsLfH5U)
		}
	}

	if pmEgen2.u67K4Sv.hCTOmp0gckSa {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '-')
	}
	c3Zu4RY =  /*line :1*/ append(c3Zu4RY,  /*line :1*/ yL0_p0wUnc.g3J8Fbn4Ip(10)...)

	if cZUDpnJQn > 0 {
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '.')
		iNeTscu052a :=  /*line :1*/ vFx5p_cm.g3J8Fbn4Ip(10)
		for aepbxLiWOZ := cZUDpnJQn -  /*line :1*/ len(iNeTscu052a); aepbxLiWOZ > 0; aepbxLiWOZ-- {
			c3Zu4RY =  /*line :1*/ append(c3Zu4RY, '0')
		}
		c3Zu4RY =  /*line :1*/ append(c3Zu4RY, iNeTscu052a...)
	}

	return  /*line :1*/ string(c3Zu4RY)
}
