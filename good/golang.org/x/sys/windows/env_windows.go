//line :1
package NJ4MerJ

import (
	syscall "bUKeamF"
	"unsafe"
)

func O2SfSC0i(hGZWAKz string) (wvZBcfh string, eBQTw2tMark bool) {
	return  /*line :1*/ syscall.ITYf1B(hGZWAKz)
}

func CZ6yQWKUUf(hGZWAKz, wvZBcfh string) error {
	return  /*line :1*/ syscall.BNcaDu(hGZWAKz, wvZBcfh)
}

func EvaruFMvvGW() {
	 /*line :1*/ syscall.J1XgGX_nmS4b()
}

func WTyokD() []string {
	return  /*line :1*/ syscall.FDhTFlUY()
}




func (lAmUDbCC TaSPPoJMlh) Environ(wSMYeeS bool) (uG96gcyOx []string, jeWMpOaQtWV error) {
	var o4R_INp *uint16
	jeWMpOaQtWV =  /*line :1*/ WdSzERJZLW(&o4R_INp, lAmUDbCC, wSMYeeS)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	defer  /*line :1*/ CG120MaPv8q(o4R_INp)
	a0SAP1YR :=  /*line :1*/ unsafe.Pointer(o4R_INp)
	for {
		gak4tmVgN :=  /*line :1*/ M_Sea9j((* /*line :1*/ uint16)(a0SAP1YR))
		if  /*line :1*/ len(gak4tmVgN) == 0 {
			break
		}
		uG96gcyOx =  /*line :1*/ append(uG96gcyOx, gak4tmVgN)
		a0SAP1YR =  /*line :1*/ unsafe.Add(a0SAP1YR, 2*( /*line :1*/ len(gak4tmVgN)+1))
	}
	return uG96gcyOx, nil
}

func Tt7Ank0SWU(hGZWAKz string) error {
	return  /*line :1*/ syscall.OuD_oLNMM(hGZWAKz)
}
