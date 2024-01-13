//line :1
package zBESu2SrRjP

import errors "iAHoxjmM"





func chHT2YtqEP(uTN3BXbMd byte) byte {
	return uTN3BXbMd | ('x' - 'X')
}


var TD3HnU4 =  /*line :1*/ errors.Su6g6hRoi9X("value out of range")


var ZTbjMv =  /*line :1*/ errors.Su6g6hRoi9X("invalid syntax")


type RJxImsn8bz struct {
	Xl68D9T59	string	
	HY76vGY_	string	
	S08K00Kl9csG	error	
}

func (afGPrH *RJxImsn8bz) Error() string {
	return "strconv." + afGPrH.Xl68D9T59 + ": " + "parsing " +  /*line :1*/ D35LD5nn(afGPrH.HY76vGY_) + ": " +  /*line :1*/ afGPrH.S08K00Kl9csG.Error()
}

func (afGPrH *RJxImsn8bz) Unwrap() error	{ return afGPrH.S08K00Kl9csG }















func sQ4Jw3Ftqji(eswrOOpIaF string) string	{ return  /*line :1*/ string([] /*line :1*/ byte(eswrOOpIaF)) }

func nF25LgNzv(tEXHDL6, jREhFXLW3 string) *RJxImsn8bz {
	return &RJxImsn8bz{tEXHDL6,  /*line :1*/ sQ4Jw3Ftqji(jREhFXLW3), ZTbjMv}
}

func b_MqXW4(tEXHDL6, jREhFXLW3 string) *RJxImsn8bz {
	return &RJxImsn8bz{tEXHDL6,  /*line :1*/ sQ4Jw3Ftqji(jREhFXLW3), TD3HnU4}
}

func djtiNc(tEXHDL6, jREhFXLW3 string, xsj80lMaPD int) *RJxImsn8bz {
	return &RJxImsn8bz{tEXHDL6,  /*line :1*/ sQ4Jw3Ftqji(jREhFXLW3),  /*line :1*/ errors.Su6g6hRoi9X("invalid base " +  /*line :1*/ ZW1_FxJtq(xsj80lMaPD))}
}

func htEZeDC1(tEXHDL6, jREhFXLW3 string, r2moFi1 int) *RJxImsn8bz {
	return &RJxImsn8bz{tEXHDL6,  /*line :1*/ sQ4Jw3Ftqji(jREhFXLW3),  /*line :1*/ errors.Su6g6hRoi9X("invalid bit size " +  /*line :1*/ ZW1_FxJtq(r2moFi1))}
}

const intSize = 32 << (^ /*line :1*/ uint(0) >> 63)


const IntSize = intSize

const maxUint64 = 1<<64 - 1




func WOxcnaOAzTeq(a0M8QSnCR string, xsj80lMaPD int, r2moFi1 int) (uint64, error) {
	const fnParseUint = "ParseUint"

	if a0M8QSnCR == "" {
		return 0,  /*line :1*/ nF25LgNzv(fnParseUint, a0M8QSnCR)
	}

	rC7Y3pzBA := xsj80lMaPD == 0

	qq6EItVuMe1d := a0M8QSnCR
	switch {
	case 2 <= xsj80lMaPD && xsj80lMaPD <= 36:

	case xsj80lMaPD == 0:

		xsj80lMaPD = 10
		if a0M8QSnCR[0] == '0' {
			switch {
			case  /*line :1*/ len(a0M8QSnCR) >= 3 &&  /*line :1*/ chHT2YtqEP(a0M8QSnCR[1]) == 'b':
				xsj80lMaPD = 2
				a0M8QSnCR = a0M8QSnCR[2:]
			case  /*line :1*/ len(a0M8QSnCR) >= 3 &&  /*line :1*/ chHT2YtqEP(a0M8QSnCR[1]) == 'o':
				xsj80lMaPD = 8
				a0M8QSnCR = a0M8QSnCR[2:]
			case  /*line :1*/ len(a0M8QSnCR) >= 3 &&  /*line :1*/ chHT2YtqEP(a0M8QSnCR[1]) == 'x':
				xsj80lMaPD = 16
				a0M8QSnCR = a0M8QSnCR[2:]
			default:
				xsj80lMaPD = 8
				a0M8QSnCR = a0M8QSnCR[1:]
			}
		}

	default:
		return 0,  /*line :1*/ djtiNc(fnParseUint, qq6EItVuMe1d, xsj80lMaPD)
	}

	if r2moFi1 == 0 {
		r2moFi1 = IntSize
	} else if r2moFi1 < 0 || r2moFi1 > 64 {
		return 0,  /*line :1*/ htEZeDC1(fnParseUint, qq6EItVuMe1d, r2moFi1)
	}

	
	
	var hmu04Zza9Am uint64
	switch xsj80lMaPD {
	case 10:
		hmu04Zza9Am = maxUint64/10 + 1
	case 16:
		hmu04Zza9Am = maxUint64/16 + 1
	default:
		hmu04Zza9Am = maxUint64/ /*line :1*/ uint64(xsj80lMaPD) + 1
	}

	rc1ZIPH9V8 :=  /*line :1*/ uint64(1)<< /*line :1*/ uint(r2moFi1) - 1

	jKtJBS := false
	var wESnmZiAO uint64
	for _, uTN3BXbMd := range [] /*line :1*/ byte(a0M8QSnCR) {
		var bkcwawBry byte
		switch {
		case uTN3BXbMd == '_' && rC7Y3pzBA:
			jKtJBS = true
			continue
		case '0' <= uTN3BXbMd && uTN3BXbMd <= '9':
			bkcwawBry = uTN3BXbMd - '0'
		case 'a' <=  /*line :1*/ chHT2YtqEP(uTN3BXbMd) &&  /*line :1*/ chHT2YtqEP(uTN3BXbMd) <= 'z':
			bkcwawBry =  /*line :1*/ chHT2YtqEP(uTN3BXbMd) - 'a' + 10
		default:
			return 0,  /*line :1*/ nF25LgNzv(fnParseUint, qq6EItVuMe1d)
		}

		if bkcwawBry >=  /*line :1*/ byte(xsj80lMaPD) {
			return 0,  /*line :1*/ nF25LgNzv(fnParseUint, qq6EItVuMe1d)
		}

		if wESnmZiAO >= hmu04Zza9Am {

			return rc1ZIPH9V8,  /*line :1*/ b_MqXW4(fnParseUint, qq6EItVuMe1d)
		}
		wESnmZiAO *=  /*line :1*/ uint64(xsj80lMaPD)

		gRwuanISn := wESnmZiAO +  /*line :1*/ uint64(bkcwawBry)
		if gRwuanISn < wESnmZiAO || gRwuanISn > rc1ZIPH9V8 {

			return rc1ZIPH9V8,  /*line :1*/ b_MqXW4(fnParseUint, qq6EItVuMe1d)
		}
		wESnmZiAO = gRwuanISn
	}

	if jKtJBS && ! /*line :1*/ kTU2kqFg(qq6EItVuMe1d) {
		return 0,  /*line :1*/ nF25LgNzv(fnParseUint, qq6EItVuMe1d)
	}

	return wESnmZiAO, nil
}


























func N8Xpap1Vt1ot(a0M8QSnCR string, xsj80lMaPD int, r2moFi1 int) (cDhV_2D int64, kWNmhkfeXX error) {
	const fnParseInt = "ParseInt"

	if a0M8QSnCR == "" {
		return 0,  /*line :1*/ nF25LgNzv(fnParseInt, a0M8QSnCR)
	}

	qq6EItVuMe1d := a0M8QSnCR
	an5n8FniWt67 := false
	if a0M8QSnCR[0] == '+' {
		a0M8QSnCR = a0M8QSnCR[1:]
	} else if a0M8QSnCR[0] == '-' {
		an5n8FniWt67 = true
		a0M8QSnCR = a0M8QSnCR[1:]
	}

	
	var hr0Kt5LgWajT uint64
	hr0Kt5LgWajT, kWNmhkfeXX =  /*line :1*/ WOxcnaOAzTeq(a0M8QSnCR, xsj80lMaPD, r2moFi1)
	if kWNmhkfeXX != nil && kWNmhkfeXX.(*RJxImsn8bz).S08K00Kl9csG != TD3HnU4 {
		kWNmhkfeXX.(*RJxImsn8bz).Xl68D9T59 = fnParseInt
		kWNmhkfeXX.(*RJxImsn8bz).HY76vGY_ =  /*line :1*/ sQ4Jw3Ftqji(qq6EItVuMe1d)
		return 0, kWNmhkfeXX
	}

	if r2moFi1 == 0 {
		r2moFi1 = IntSize
	}

	hmu04Zza9Am :=  /*line :1*/ uint64(1 <<  /*line :1*/ uint(r2moFi1-1))
	if !an5n8FniWt67 && hr0Kt5LgWajT >= hmu04Zza9Am {
		return  /*line :1*/ int64(hmu04Zza9Am - 1),  /*line :1*/ b_MqXW4(fnParseInt, qq6EItVuMe1d)
	}
	if an5n8FniWt67 && hr0Kt5LgWajT > hmu04Zza9Am {
		return - /*line :1*/ int64(hmu04Zza9Am),  /*line :1*/ b_MqXW4(fnParseInt, qq6EItVuMe1d)
	}
	wESnmZiAO :=  /*line :1*/ int64(hr0Kt5LgWajT)
	if an5n8FniWt67 {
		wESnmZiAO = -wESnmZiAO
	}
	return wESnmZiAO, nil
}


func GmbOy60GCccC(a0M8QSnCR string) (int, error) {
	const fnAtoi = "Atoi"

	zbui8CqMH :=  /*line :1*/ len(a0M8QSnCR)
	if intSize == 32 && (0 < zbui8CqMH && zbui8CqMH < 10) ||
		intSize == 64 && (0 < zbui8CqMH && zbui8CqMH < 19) {

		qq6EItVuMe1d := a0M8QSnCR
		if a0M8QSnCR[0] == '-' || a0M8QSnCR[0] == '+' {
			a0M8QSnCR = a0M8QSnCR[1:]
			if  /*line :1*/ len(a0M8QSnCR) < 1 {
				return 0,  /*line :1*/ nF25LgNzv(fnAtoi, qq6EItVuMe1d)
			}
		}

		wESnmZiAO := 0
		for _, tqWO0Nj := range [] /*line :1*/ byte(a0M8QSnCR) {
			tqWO0Nj -= '0'
			if tqWO0Nj > 9 {
				return 0,  /*line :1*/ nF25LgNzv(fnAtoi, qq6EItVuMe1d)
			}
			wESnmZiAO = wESnmZiAO*10 +  /*line :1*/ int(tqWO0Nj)
		}
		if qq6EItVuMe1d[0] == '-' {
			wESnmZiAO = -wESnmZiAO
		}
		return wESnmZiAO, nil
	}

	vXcx8c0S_, kWNmhkfeXX :=  /*line :1*/ N8Xpap1Vt1ot(a0M8QSnCR, 10, 0)
	if h9X10OsPUB, eva_XlRz4Iz := kWNmhkfeXX.(*RJxImsn8bz); eva_XlRz4Iz {
		h9X10OsPUB.Xl68D9T59 = fnAtoi
	}
	return  /*line :1*/ int(vXcx8c0S_), kWNmhkfeXX
}




func kTU2kqFg(a0M8QSnCR string) bool {

	r3ptyc4arA := '^'
	cDhV_2D := 0

	if  /*line :1*/ len(a0M8QSnCR) >= 1 && (a0M8QSnCR[0] == '-' || a0M8QSnCR[0] == '+') {
		a0M8QSnCR = a0M8QSnCR[1:]
	}

	r6avN8QxQ8D := false
	if  /*line :1*/ len(a0M8QSnCR) >= 2 && a0M8QSnCR[0] == '0' && ( /*line :1*/ chHT2YtqEP(a0M8QSnCR[1]) == 'b' ||  /*line :1*/ chHT2YtqEP(a0M8QSnCR[1]) == 'o' ||  /*line :1*/ chHT2YtqEP(a0M8QSnCR[1]) == 'x') {
		cDhV_2D = 2
		r3ptyc4arA = '0'
		r6avN8QxQ8D =  /*line :1*/ chHT2YtqEP(a0M8QSnCR[1]) == 'x'
	}

	for ; cDhV_2D <  /*line :1*/ len(a0M8QSnCR); cDhV_2D++ {

		if '0' <= a0M8QSnCR[cDhV_2D] && a0M8QSnCR[cDhV_2D] <= '9' || r6avN8QxQ8D && 'a' <=  /*line :1*/ chHT2YtqEP(a0M8QSnCR[cDhV_2D]) &&  /*line :1*/ chHT2YtqEP(a0M8QSnCR[cDhV_2D]) <= 'f' {
			r3ptyc4arA = '0'
			continue
		}

		if a0M8QSnCR[cDhV_2D] == '_' {
			if r3ptyc4arA != '0' {
				return false
			}
			r3ptyc4arA = '_'
			continue
		}

		if r3ptyc4arA == '_' {
			return false
		}

		r3ptyc4arA = '!'
	}
	return r3ptyc4arA != '_'
}
