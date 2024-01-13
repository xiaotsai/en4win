//line :1
package y1BamVjnXsaa

import (
	path "jn9Psh1aa_7H"
)


type Cfa8Qygf interface {
	XIcFcDgy

	
	
	
	Glob(fBgBQzIsHea string) ([]string, error)
}













func EQtFHFYr0Mac(viJma5vOi XIcFcDgy, fBgBQzIsHea string) (nBjPPaPcYpaA []string, glldLezIW error) {
	return  /*line :1*/ iaKe6Wgftj(viJma5vOi, fBgBQzIsHea, 0)
}

func iaKe6Wgftj(viJma5vOi XIcFcDgy, fBgBQzIsHea string, vaiVKPTjpNZb int) (nBjPPaPcYpaA []string, glldLezIW error) {
	
	
	const pathSeparatorsLimit = 10000
	if vaiVKPTjpNZb > pathSeparatorsLimit {
		return nil, path.VVxvifIdGY
	}
	if viJma5vOi, d5Q0SCgdw := viJma5vOi.(Cfa8Qygf); d5Q0SCgdw {
		return  /*line :1*/ viJma5vOi.Glob(fBgBQzIsHea)
	}

	if _, glldLezIW :=  /*line :1*/ path.K9FvN0A5kW42(fBgBQzIsHea, ""); glldLezIW != nil {
		return nil, glldLezIW
	}
	if ! /*line :1*/ ahnlDp8pViv(fBgBQzIsHea) {
		if _, glldLezIW =  /*line :1*/ PTHkijCMP7c(viJma5vOi, fBgBQzIsHea); glldLezIW != nil {
			return nil, nil
		}
		return []string{fBgBQzIsHea}, nil
	}

	cNaiWwqHDG, kccGEdZzAp :=  /*line :1*/ path.KHYEhOz2l(fBgBQzIsHea)
	cNaiWwqHDG =  /*line :1*/ dfw_69lK(cNaiWwqHDG)

	if ! /*line :1*/ ahnlDp8pViv(cNaiWwqHDG) {
		return  /*line :1*/ vlR3bT0(viJma5vOi, cNaiWwqHDG, kccGEdZzAp, nil)
	}

	if cNaiWwqHDG == fBgBQzIsHea {
		return nil, path.VVxvifIdGY
	}

	var jREaoX_I7TB6 []string
	jREaoX_I7TB6, glldLezIW =  /*line :1*/ iaKe6Wgftj(viJma5vOi, cNaiWwqHDG, vaiVKPTjpNZb+1)
	if glldLezIW != nil {
		return nil, glldLezIW
	}
	for _, kwKOrGt := range jREaoX_I7TB6 {
		nBjPPaPcYpaA, glldLezIW =  /*line :1*/ vlR3bT0(viJma5vOi, kwKOrGt, kccGEdZzAp, nBjPPaPcYpaA)
		if glldLezIW != nil {
			return
		}
	}
	return
}


func dfw_69lK(qkr_B7ssRGvO string) string {
	switch qkr_B7ssRGvO {
	case "":
		return "."
	default:
		return qkr_B7ssRGvO[0 :  /*line :1*/ len(qkr_B7ssRGvO)-1]
	}
}





func vlR3bT0(hH8UsC XIcFcDgy, cNaiWwqHDG, fBgBQzIsHea string, nBjPPaPcYpaA []string) (jREaoX_I7TB6 []string, c8anSmr6wZV error) {
	jREaoX_I7TB6 = nBjPPaPcYpaA
	bUmddsvya0Vg, glldLezIW :=  /*line :1*/ VOXtzDEapl(hH8UsC, cNaiWwqHDG)
	if glldLezIW != nil {
		return
	}

	for _, v9FHH9DB9sM2 := range bUmddsvya0Vg {
		l5arJmbvBL0f :=  /*line :1*/ v9FHH9DB9sM2.Name()
		uj0oTvTAkme, glldLezIW :=  /*line :1*/ path.K9FvN0A5kW42(fBgBQzIsHea, l5arJmbvBL0f)
		if glldLezIW != nil {
			return jREaoX_I7TB6, glldLezIW
		}
		if uj0oTvTAkme {
			jREaoX_I7TB6 =  /*line :1*/ append(jREaoX_I7TB6,  /*line :1*/ path.OT0k1D(cNaiWwqHDG, l5arJmbvBL0f))
		}
	}
	return
}



func ahnlDp8pViv(qkr_B7ssRGvO string) bool {
	for aTOpQutlw := 0; aTOpQutlw <  /*line :1*/ len(qkr_B7ssRGvO); aTOpQutlw++ {
		switch qkr_B7ssRGvO[aTOpQutlw] {
		case '*', '?', '[', '\\':
			return true
		}
	}
	return false
}
