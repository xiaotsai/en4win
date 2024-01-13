//line :1
package y1BamVjnXsaa

import io "xI9ai2djaJ2"



type DiFnY2 interface {
	XIcFcDgy

	
	
	
	
	
	
	
	ReadFile(hMWDrBfy string) ([]byte, error)
}









func KcCxODByBerE(viJma5vOi XIcFcDgy, hMWDrBfy string) ([]byte, error) {
	if viJma5vOi, d5Q0SCgdw := viJma5vOi.(DiFnY2); d5Q0SCgdw {
		return  /*line :1*/ viJma5vOi.ReadFile(hMWDrBfy)
	}

	kccGEdZzAp, glldLezIW :=  /*line :1*/ viJma5vOi.Open(hMWDrBfy)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	defer  /*line :1*/ kccGEdZzAp.Close()

	var x9px08Yregqq int
	if v9FHH9DB9sM2, glldLezIW :=  /*line :1*/ kccGEdZzAp.Stat(); glldLezIW == nil {
		b76i7L :=  /*line :1*/ v9FHH9DB9sM2.Size()
		if  /*line :1*/ int64( /*line :1*/ int(b76i7L)) == b76i7L {
			x9px08Yregqq =  /*line :1*/ int(b76i7L)
		}
	}

	zdEMFBnd :=  /*line :1*/ make([]byte, 0, x9px08Yregqq+1)
	for {
		if  /*line :1*/ len(zdEMFBnd) >=  /*line :1*/ cap(zdEMFBnd) {
			kwKOrGt :=  /*line :1*/ append(zdEMFBnd[: /*line :1*/ cap(zdEMFBnd)], 0)
			zdEMFBnd = kwKOrGt[: /*line :1*/ len(zdEMFBnd)]
		}
		l5arJmbvBL0f, glldLezIW :=  /*line :1*/ kccGEdZzAp.Read(zdEMFBnd[ /*line :1*/ len(zdEMFBnd): /*line :1*/ cap(zdEMFBnd)])
		zdEMFBnd = zdEMFBnd[: /*line :1*/ len(zdEMFBnd)+l5arJmbvBL0f]
		if glldLezIW != nil {
			if glldLezIW == io.K5Sqco {
				glldLezIW = nil
			}
			return zdEMFBnd, glldLezIW
		}
	}
}
