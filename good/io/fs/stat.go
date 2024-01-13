//line :1
package y1BamVjnXsaa


type Q0mjhH9y interface {
	XIcFcDgy

	
	
	Stat(hMWDrBfy string) (HM4JM2, error)
}





func PTHkijCMP7c(viJma5vOi XIcFcDgy, hMWDrBfy string) (HM4JM2, error) {
	if viJma5vOi, d5Q0SCgdw := viJma5vOi.(Q0mjhH9y); d5Q0SCgdw {
		return  /*line :1*/ viJma5vOi.Stat(hMWDrBfy)
	}

	kccGEdZzAp, glldLezIW :=  /*line :1*/ viJma5vOi.Open(hMWDrBfy)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	defer  /*line :1*/ kccGEdZzAp.Close()
	return  /*line :1*/ kccGEdZzAp.Stat()
}
