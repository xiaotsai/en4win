//line :1
package y1BamVjnXsaa

import (
	errors "iAHoxjmM"
	sort "gzHaha55n"
)



type Q1kvvSVy interface {
	XIcFcDgy

	
	
	ReadDir(hMWDrBfy string) ([]RbfBMxa, error)
}







func VOXtzDEapl(viJma5vOi XIcFcDgy, hMWDrBfy string) ([]RbfBMxa, error) {
	if viJma5vOi, d5Q0SCgdw := viJma5vOi.(Q1kvvSVy); d5Q0SCgdw {
		return  /*line :1*/ viJma5vOi.ReadDir(hMWDrBfy)
	}

	kccGEdZzAp, glldLezIW :=  /*line :1*/ viJma5vOi.Open(hMWDrBfy)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	defer  /*line :1*/ kccGEdZzAp.Close()

	cNaiWwqHDG, d5Q0SCgdw := kccGEdZzAp.(FhagHI)
	if !d5Q0SCgdw {
		return nil, &CFLL3J{Aeg62j1VQt: "readdir", OmDEtY: hMWDrBfy, OIEifQ0XF:  /*line :1*/ errors.Su6g6hRoi9X("not implemented")}
	}

	gJjGmD, glldLezIW :=  /*line :1*/ cNaiWwqHDG.ReadDir(-1)
	 /*line :1*/ sort.Wo2XqNIPTLLO(gJjGmD, func(aTOpQutlw, opa4xD_Ezoa int) bool { return  /*line :1*/ gJjGmD[aTOpQutlw].Name() <  /*line :1*/ gJjGmD[opa4xD_Ezoa].Name() })
	return gJjGmD, glldLezIW
}


type g0QKytxhM struct {
	lhnAYLU HM4JM2
}

func (ahrLuLh0X g0QKytxhM) IsDir() bool {
	return  /*line :1*/ ahrLuLh0X.lhnAYLU.IsDir()
}

func (ahrLuLh0X g0QKytxhM) Type() PGXMbJjlHBMu {
	return  /*line :1*/ ahrLuLh0X.lhnAYLU.Mode().Type()
}

func (ahrLuLh0X g0QKytxhM) Info() (HM4JM2, error) {
	return ahrLuLh0X.lhnAYLU, nil
}

func (ahrLuLh0X g0QKytxhM) Name() string {
	return  /*line :1*/ ahrLuLh0X.lhnAYLU.Name()
}

func (ahrLuLh0X g0QKytxhM) String() string {
	return  /*line :1*/ EMHCka(ahrLuLh0X)
}



func SbRUvyw(v9FHH9DB9sM2 HM4JM2) RbfBMxa {
	if v9FHH9DB9sM2 == nil {
		return nil
	}
	return g0QKytxhM{lhnAYLU: v9FHH9DB9sM2}
}
