//line :1
package zBESu2SrRjP

const fnParseComplex = "ParseComplex"



func eib1137gLY(kWNmhkfeXX error, a0M8QSnCR string) (aU1_xB, jmdm6VR1i error) {
	if eswrOOpIaF, eva_XlRz4Iz := kWNmhkfeXX.(*RJxImsn8bz); eva_XlRz4Iz {
		eswrOOpIaF.Xl68D9T59 = fnParseComplex
		eswrOOpIaF.HY76vGY_ =  /*line :1*/ sQ4Jw3Ftqji(a0M8QSnCR)
		if eswrOOpIaF.S08K00Kl9csG == TD3HnU4 {
			return nil, eswrOOpIaF
		}
	}
	return kWNmhkfeXX, nil
}





















func DjOus4h(a0M8QSnCR string, r2moFi1 int) (complex128, error) {
	mugT3utgZ := 64
	if r2moFi1 == 64 {
		mugT3utgZ = 32
	}

	e6nEwdy := a0M8QSnCR

	if  /*line :1*/ len(a0M8QSnCR) >= 2 && a0M8QSnCR[0] == '(' && a0M8QSnCR[ /*line :1*/ len(a0M8QSnCR)-1] == ')' {
		a0M8QSnCR = a0M8QSnCR[1 :  /*line :1*/ len(a0M8QSnCR)-1]
	}

	var htfmeaWZQMk error	

	hE28rzwD, wESnmZiAO, kWNmhkfeXX :=  /*line :1*/ xBeCkrUFdCv(a0M8QSnCR, mugT3utgZ)
	if kWNmhkfeXX != nil {
		kWNmhkfeXX, htfmeaWZQMk =  /*line :1*/ eib1137gLY(kWNmhkfeXX, e6nEwdy)
		if kWNmhkfeXX != nil {
			return 0, kWNmhkfeXX
		}
	}
	a0M8QSnCR = a0M8QSnCR[wESnmZiAO:]

	if  /*line :1*/ len(a0M8QSnCR) == 0 {
		return  /*line :1*/ complex(hE28rzwD, 0), htfmeaWZQMk
	}

	switch a0M8QSnCR[0] {
	case '+':

		if  /*line :1*/ len(a0M8QSnCR) > 1 && a0M8QSnCR[1] != '+' {
			a0M8QSnCR = a0M8QSnCR[1:]
		}
	case '-':

	case 'i':

		if  /*line :1*/ len(a0M8QSnCR) == 1 {
			return  /*line :1*/ complex(0, hE28rzwD), htfmeaWZQMk
		}
		fallthrough
	default:
		return 0,  /*line :1*/ nF25LgNzv(fnParseComplex, e6nEwdy)
	}

	j6TCto, wESnmZiAO, kWNmhkfeXX :=  /*line :1*/ xBeCkrUFdCv(a0M8QSnCR, mugT3utgZ)
	if kWNmhkfeXX != nil {
		kWNmhkfeXX, htfmeaWZQMk =  /*line :1*/ eib1137gLY(kWNmhkfeXX, e6nEwdy)
		if kWNmhkfeXX != nil {
			return 0, kWNmhkfeXX
		}
	}
	a0M8QSnCR = a0M8QSnCR[wESnmZiAO:]
	if a0M8QSnCR != "i" {
		return 0,  /*line :1*/ nF25LgNzv(fnParseComplex, e6nEwdy)
	}
	return  /*line :1*/ complex(hE28rzwD, j6TCto), htfmeaWZQMk
}
