//line :1
package y1BamVjnXsaa

import (
	errors "iAHoxjmM"
	path "jn9Psh1aa_7H"
)


type A_8v_RIR interface {
	XIcFcDgy

	
	Sub(cNaiWwqHDG string) (XIcFcDgy, error)
}















func XIUbUvyZc(viJma5vOi XIcFcDgy, cNaiWwqHDG string) (XIcFcDgy, error) {
	if ! /*line :1*/ QuVTpr7RfE(cNaiWwqHDG) {
		return nil, &CFLL3J{Aeg62j1VQt: "sub", OmDEtY: cNaiWwqHDG, OIEifQ0XF:  /*line :1*/ errors.Su6g6hRoi9X("invalid name")}
	}
	if cNaiWwqHDG == "." {
		return viJma5vOi, nil
	}
	if viJma5vOi, d5Q0SCgdw := viJma5vOi.(A_8v_RIR); d5Q0SCgdw {
		return  /*line :1*/ viJma5vOi.Sub(cNaiWwqHDG)
	}
	return &jffr01{viJma5vOi, cNaiWwqHDG}, nil
}

type jffr01 struct {
	dmz5Oq_	XIcFcDgy
	iZZPMHmcWC	string
}


func (iqCsM044 *jffr01) iAQXWD(zc112zGamy8 string, hMWDrBfy string) (string, error) {
	if ! /*line :1*/ QuVTpr7RfE(hMWDrBfy) {
		return "", &CFLL3J{Aeg62j1VQt: zc112zGamy8, OmDEtY: hMWDrBfy, OIEifQ0XF:  /*line :1*/ errors.Su6g6hRoi9X("invalid name")}
	}
	return  /*line :1*/ path.OT0k1D(iqCsM044.iZZPMHmcWC, hMWDrBfy), nil
}


func (iqCsM044 *jffr01) gugba5GQxF(hMWDrBfy string) (lCgjW1E string, d5Q0SCgdw bool) {
	if hMWDrBfy == iqCsM044.iZZPMHmcWC {
		return ".", true
	}
	if  /*line :1*/ len(hMWDrBfy) >=  /*line :1*/ len(iqCsM044.iZZPMHmcWC)+2 && hMWDrBfy[ /*line :1*/ len(iqCsM044.iZZPMHmcWC)] == '/' && hMWDrBfy[: /*line :1*/ len(iqCsM044.iZZPMHmcWC)] == iqCsM044.iZZPMHmcWC {
		return hMWDrBfy[ /*line :1*/ len(iqCsM044.iZZPMHmcWC)+1:], true
	}
	return "", false
}


func (iqCsM044 *jffr01) auFRrM5Ls5(glldLezIW error) error {
	if c8anSmr6wZV, d5Q0SCgdw := glldLezIW.(*CFLL3J); d5Q0SCgdw {
		if cjSqH0k2HM, d5Q0SCgdw :=  /*line :1*/ iqCsM044.gugba5GQxF(c8anSmr6wZV.OmDEtY); d5Q0SCgdw {
			c8anSmr6wZV.OmDEtY = cjSqH0k2HM
		}
	}
	return glldLezIW
}

func (iqCsM044 *jffr01) Open(hMWDrBfy string) (YY1DXRrw, error) {
	yrlte6eXsl4, glldLezIW :=  /*line :1*/ iqCsM044.iAQXWD("open", hMWDrBfy)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	kccGEdZzAp, glldLezIW :=  /*line :1*/ iqCsM044.dmz5Oq_.Open(yrlte6eXsl4)
	return kccGEdZzAp,  /*line :1*/ iqCsM044.auFRrM5Ls5(glldLezIW)
}

func (iqCsM044 *jffr01) ReadDir(hMWDrBfy string) ([]RbfBMxa, error) {
	yrlte6eXsl4, glldLezIW :=  /*line :1*/ iqCsM044.iAQXWD("read", hMWDrBfy)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	cNaiWwqHDG, glldLezIW :=  /*line :1*/ VOXtzDEapl(iqCsM044.dmz5Oq_, yrlte6eXsl4)
	return cNaiWwqHDG,  /*line :1*/ iqCsM044.auFRrM5Ls5(glldLezIW)
}

func (iqCsM044 *jffr01) ReadFile(hMWDrBfy string) ([]byte, error) {
	yrlte6eXsl4, glldLezIW :=  /*line :1*/ iqCsM044.iAQXWD("read", hMWDrBfy)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	zdEMFBnd, glldLezIW :=  /*line :1*/ KcCxODByBerE(iqCsM044.dmz5Oq_, yrlte6eXsl4)
	return zdEMFBnd,  /*line :1*/ iqCsM044.auFRrM5Ls5(glldLezIW)
}

func (iqCsM044 *jffr01) Glob(fBgBQzIsHea string) ([]string, error) {

	if _, glldLezIW :=  /*line :1*/ path.K9FvN0A5kW42(fBgBQzIsHea, ""); glldLezIW != nil {
		return nil, glldLezIW
	}
	if fBgBQzIsHea == "." {
		return []string{"."}, nil
	}

	yrlte6eXsl4 := iqCsM044.iZZPMHmcWC + "/" + fBgBQzIsHea
	gJjGmD, glldLezIW :=  /*line :1*/ EQtFHFYr0Mac(iqCsM044.dmz5Oq_, yrlte6eXsl4)
	for aTOpQutlw, hMWDrBfy := range gJjGmD {
		hMWDrBfy, d5Q0SCgdw :=  /*line :1*/ iqCsM044.gugba5GQxF(hMWDrBfy)
		if !d5Q0SCgdw {
			return nil,  /*line :1*/ errors.Su6g6hRoi9X("invalid result from inner fsys Glob: " + hMWDrBfy + " not in " + iqCsM044.iZZPMHmcWC)
		}
		gJjGmD[aTOpQutlw] = hMWDrBfy
	}
	return gJjGmD,  /*line :1*/ iqCsM044.auFRrM5Ls5(glldLezIW)
}

func (iqCsM044 *jffr01) Sub(cNaiWwqHDG string) (XIcFcDgy, error) {
	if cNaiWwqHDG == "." {
		return iqCsM044, nil
	}
	yrlte6eXsl4, glldLezIW :=  /*line :1*/ iqCsM044.iAQXWD("sub", cNaiWwqHDG)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	return &jffr01{iqCsM044.dmz5Oq_, yrlte6eXsl4}, nil
}
