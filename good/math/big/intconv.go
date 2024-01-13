//line :1
package big

import (
	errors "iAHoxjmM"
	fmt "kFsoCfy5zWG"
	io "xI9ai2djaJ2"
)







func (pmEgen2 *ZMtGuo) Text(weSJS4i4lm int) string {
	if pmEgen2 == nil {
		return "<nil>"
	}
	return  /*line :1*/ string( /*line :1*/ pmEgen2.qho77PBKF.sgVtOtra(pmEgen2.hCTOmp0gckSa, weSJS4i4lm))
}



func (pmEgen2 *ZMtGuo) Append(c3Zu4RY []byte, weSJS4i4lm int) []byte {
	if pmEgen2 == nil {
		return  /*line :1*/ append(c3Zu4RY, "<nil>"...)
	}
	return  /*line :1*/ append(c3Zu4RY,  /*line :1*/ pmEgen2.qho77PBKF.sgVtOtra(pmEgen2.hCTOmp0gckSa, weSJS4i4lm)...)
}



func (pmEgen2 *ZMtGuo) String() string {
	return  /*line :1*/ pmEgen2.Text(10)
}


func nsheC69bwii(nR7KU4mGsdy fmt.Bm1Jqy8i, yK8eFt string, qWfhyOCfC0By int) {
	if  /*line :1*/ len(yK8eFt) > 0 {
		kWa1YrIHZo := [] /*line :1*/ byte(yK8eFt)
		for ; qWfhyOCfC0By > 0; qWfhyOCfC0By-- {
			 /*line :1*/ nR7KU4mGsdy.Write(kWa1YrIHZo)
		}
	}
}

var _ fmt.OpUJYyZOy8 = h2Aj5SV9g2re	












func (pmEgen2 *ZMtGuo) Format(nR7KU4mGsdy fmt.Bm1Jqy8i, qvA6MnWV rune) {
	
	var weSJS4i4lm int
	switch qvA6MnWV {
	case 'b':
		weSJS4i4lm = 2
	case 'o', 'O':
		weSJS4i4lm = 8
	case 'd', 's', 'v':
		weSJS4i4lm = 10
	case 'x', 'X':
		weSJS4i4lm = 16
	default:

		 /*line :1*/ fmt.BYcL2whVEYz(nR7KU4mGsdy, "%%!%c(big.Int=%s)", qvA6MnWV,  /*line :1*/ pmEgen2.String())
		return
	}

	if pmEgen2 == nil {
		 /*line :1*/ fmt.AJRwHmjcphiN(nR7KU4mGsdy, "<nil>")
		return
	}

	e2mU4H := ""
	switch {
	case pmEgen2.hCTOmp0gckSa:
		e2mU4H = "-"
	case  /*line :1*/ nR7KU4mGsdy.Flag('+'):
		e2mU4H = "+"
	case  /*line :1*/ nR7KU4mGsdy.Flag(' '):
		e2mU4H = " "
	}

	vpvKXvnZ0CwR := ""
	if  /*line :1*/ nR7KU4mGsdy.Flag('#') {
		switch qvA6MnWV {
		case 'b':
			vpvKXvnZ0CwR = "0b"
		case 'o':
			vpvKXvnZ0CwR = "0"
		case 'x':
			vpvKXvnZ0CwR = "0x"
		case 'X':
			vpvKXvnZ0CwR = "0X"
		}
	}
	if qvA6MnWV == 'O' {
		vpvKXvnZ0CwR = "0o"
	}

	cmZK_aGSDT :=  /*line :1*/ pmEgen2.qho77PBKF.g3J8Fbn4Ip(weSJS4i4lm)
	if qvA6MnWV == 'X' {

		for aepbxLiWOZ, s31g8J := range cmZK_aGSDT {
			if 'a' <= s31g8J && s31g8J <= 'z' {
				cmZK_aGSDT[aepbxLiWOZ] = 'A' + (s31g8J - 'a')
			}
		}
	}

	
	var imQQ7pmhsOpJ int	
	var mBB2P3 int	
	var riBrn7H5oD int	

	giT3UJmkoxaB, kJt1_wPypBhy :=  /*line :1*/ nR7KU4mGsdy.Precision()
	if kJt1_wPypBhy {
		switch {
		case  /*line :1*/ len(cmZK_aGSDT) < giT3UJmkoxaB:
			mBB2P3 = giT3UJmkoxaB -  /*line :1*/ len(cmZK_aGSDT)
		case  /*line :1*/ len(cmZK_aGSDT) == 1 && cmZK_aGSDT[0] == '0' && giT3UJmkoxaB == 0:
			return
		}
	}

	ia0yKDa8fSgi :=  /*line :1*/ len(e2mU4H) +  /*line :1*/ len(vpvKXvnZ0CwR) + mBB2P3 +  /*line :1*/ len(cmZK_aGSDT)
	if pRxc6yBZ9, yIO9L2XVD :=  /*line :1*/ nR7KU4mGsdy.Width(); yIO9L2XVD && ia0yKDa8fSgi < pRxc6yBZ9 {
		switch s31g8J := pRxc6yBZ9 - ia0yKDa8fSgi; {
		case  /*line :1*/ nR7KU4mGsdy.Flag('-'):

			riBrn7H5oD = s31g8J
		case  /*line :1*/ nR7KU4mGsdy.Flag('0') && !kJt1_wPypBhy:

			mBB2P3 = s31g8J
		default:

			imQQ7pmhsOpJ = s31g8J
		}
	}

	 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, " ", imQQ7pmhsOpJ)
	 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, e2mU4H, 1)
	 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, vpvKXvnZ0CwR, 1)
	 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, "0", mBB2P3)
	 /*line :1*/ nR7KU4mGsdy.Write(cmZK_aGSDT)
	 /*line :1*/ nsheC69bwii(nR7KU4mGsdy, " ", riBrn7H5oD)
}












func (uW0eIVjy *ZMtGuo) h9c5N0bF(vFx5p_cm io.HSeaoj16Wq0, weSJS4i4lm int) (*ZMtGuo, int, error) {

	s_9Ih_, gDonrwI8VP50 :=  /*line :1*/ xOqvCaQwvh(vFx5p_cm)
	if gDonrwI8VP50 != nil {
		return nil, 0, gDonrwI8VP50
	}

	uW0eIVjy.qho77PBKF, weSJS4i4lm, _, gDonrwI8VP50 =  /*line :1*/ uW0eIVjy.qho77PBKF.h9c5N0bF(vFx5p_cm, weSJS4i4lm, false)
	if gDonrwI8VP50 != nil {
		return nil, weSJS4i4lm, gDonrwI8VP50
	}
	uW0eIVjy.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.qho77PBKF) > 0 && s_9Ih_

	return uW0eIVjy, weSJS4i4lm, nil
}

func xOqvCaQwvh(vFx5p_cm io.HSeaoj16Wq0) (s_9Ih_ bool, gDonrwI8VP50 error) {
	var qvA6MnWV byte
	if qvA6MnWV, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte(); gDonrwI8VP50 != nil {
		return false, gDonrwI8VP50
	}
	switch qvA6MnWV {
	case '-':
		s_9Ih_ = true
	case '+':

	default:
		 /*line :1*/ vFx5p_cm.UnreadByte()
	}
	return
}



type d8BACp struct {
	fmt.XDnSs0f3
}

func (vFx5p_cm d8BACp) ReadByte() (byte, error) {
	qvA6MnWV, fMLuTI3Zh, gDonrwI8VP50 :=  /*line :1*/ vFx5p_cm.ReadRune()
	if fMLuTI3Zh != 1 && gDonrwI8VP50 == nil {
		gDonrwI8VP50 =  /*line :1*/ fmt.Cnz_ab1BaZh("invalid rune %#U", qvA6MnWV)
	}
	return  /*line :1*/ byte(qvA6MnWV), gDonrwI8VP50
}

func (vFx5p_cm d8BACp) UnreadByte() error {
	return  /*line :1*/ vFx5p_cm.UnreadRune()
}

var _ fmt.MqdnB6 = h2Aj5SV9g2re	




func (uW0eIVjy *ZMtGuo) Scan(nR7KU4mGsdy fmt.XDnSs0f3, qvA6MnWV rune) error {
	 /*line :1*/ nR7KU4mGsdy.SkipSpace()
	weSJS4i4lm := 0
	switch qvA6MnWV {
	case 'b':
		weSJS4i4lm = 2
	case 'o':
		weSJS4i4lm = 8
	case 'd':
		weSJS4i4lm = 10
	case 'x', 'X':
		weSJS4i4lm = 16
	case 's', 'v':

	default:
		return  /*line :1*/ errors.Su6g6hRoi9X("Int.Scan: invalid verb")
	}
	_, _, gDonrwI8VP50 :=  /*line :1*/ uW0eIVjy.h9c5N0bF(d8BACp{nR7KU4mGsdy}, weSJS4i4lm)
	return gDonrwI8VP50
}
