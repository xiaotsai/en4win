//line :1
package uBCMx3OoU

import (
	reflect "reflect"
	"unsafe"

	ole "fuA83L"
)

type ar_mWKlub struct {
	cpNdyUaVENy	*f_OiDUX
	aAQE1w	int32
	drlgzSLc	*ole.GUID
	bCqdbzoT0	interface{}
	g1if3NWq	map[string]int32
}

type f_OiDUX struct {
	wHAJNe	uintptr
	fuVf5d7Yf7	uintptr
	a_B0MzQg7J	uintptr
	a_mYlb	uintptr
	q_KkJi9_o	uintptr
	iqvf392iF	uintptr
	kbDaVyp	uintptr
}

func kMSCaur(zzaNK8o2a *ole.IUnknown, kM7ECagotLCo *ole.GUID, c4R6SnZLM2m **ole.IUnknown) uint32 {
	tVOS6NO := (* /*line :1*/ ar_mWKlub)( /*line :1*/ unsafe.Pointer(zzaNK8o2a))
	*c4R6SnZLM2m = nil
	if  /*line :1*/ ole.Apsz2oVG2lN(kM7ECagotLCo, ole.WySQlWR) ||
		 /*line :1*/ ole.Apsz2oVG2lN(kM7ECagotLCo, ole.Y0sP6F) {
		 /*line :1*/ auCx1_dDF8qu(zzaNK8o2a)
		*c4R6SnZLM2m = zzaNK8o2a
		return ole.S_OK
	}
	if  /*line :1*/ ole.Apsz2oVG2lN(kM7ECagotLCo, tVOS6NO.drlgzSLc) {
		 /*line :1*/ auCx1_dDF8qu(zzaNK8o2a)
		*c4R6SnZLM2m = zzaNK8o2a
		return ole.S_OK
	}
	return ole.E_NOINTERFACE
}

func auCx1_dDF8qu(zzaNK8o2a *ole.IUnknown) int32 {
	tVOS6NO := (* /*line :1*/ ar_mWKlub)( /*line :1*/ unsafe.Pointer(zzaNK8o2a))
	tVOS6NO.aAQE1w++
	return tVOS6NO.aAQE1w
}

func krKwiq5ldnCn(zzaNK8o2a *ole.IUnknown) int32 {
	tVOS6NO := (* /*line :1*/ ar_mWKlub)( /*line :1*/ unsafe.Pointer(zzaNK8o2a))
	tVOS6NO.aAQE1w--
	return tVOS6NO.aAQE1w
}

func nAWjalYrV(zzaNK8o2a *ole.IUnknown, kM7ECagotLCo *ole.GUID, tdkFLnvdz7Ot []*uint16, jJBLcI_Q int, ccxDhjiqV int, ppSaE45YewZ []int32) uintptr {
	tVOS6NO := (* /*line :1*/ ar_mWKlub)( /*line :1*/ unsafe.Pointer(zzaNK8o2a))
	eUnZKz :=  /*line :1*/ make([]string,  /*line :1*/ len(tdkFLnvdz7Ot))
	for ghMHcK4 := 0; ghMHcK4 <  /*line :1*/ len(eUnZKz); ghMHcK4++ {
		eUnZKz[ghMHcK4] =  /*line :1*/ ole.WMh1Si1a2(tdkFLnvdz7Ot[ghMHcK4])
	}
	for tDnLQLj := 0; tDnLQLj < jJBLcI_Q; tDnLQLj++ {
		if qePwQaxmG0, rmDZGevOH := tVOS6NO.g1if3NWq[eUnZKz[tDnLQLj]]; rmDZGevOH {
			ppSaE45YewZ[tDnLQLj] = qePwQaxmG0
		}
	}
	return ole.S_OK
}

func fVkBn0q5GbO(zdFLLYVkh69g *int) uintptr {
	if zdFLLYVkh69g != nil {
		*zdFLLYVkh69g = 0
	}
	return ole.S_OK
}

func ePZAQyooJ2(cqKzWFWNb *uintptr) uintptr {
	return ole.E_NOTIMPL
}

func gMeFKNUYgE(zzaNK8o2a *ole.IDispatch, t4dsVB19SvX9 int32, a_N4Dc *ole.GUID, ccxDhjiqV int, emxPdfmG int16, iMzIvTa1zV *ole.QnAeLz, _tP7S5iJ *ole.KEVetAOpxl0D, aypGhHCR3 *ole.I0iHb3gLxw, dlFvUAB *uint) uintptr {
	tVOS6NO := (* /*line :1*/ ar_mWKlub)( /*line :1*/ unsafe.Pointer(zzaNK8o2a))
	lVvQpAb6e := ""
	for yqT6bcsPst, qePwQaxmG0 := range tVOS6NO.g1if3NWq {
		if qePwQaxmG0 == t4dsVB19SvX9 {
			lVvQpAb6e = yqT6bcsPst
		}
	}
	if lVvQpAb6e != "" {
		cBPFBDh :=  /*line :1*/ reflect.SdHoaIQl(tVOS6NO.bCqdbzoT0).Elem()
		lbNe35uJpK :=  /*line :1*/ cBPFBDh.MethodByName(lVvQpAb6e)
		mJeRbEm :=  /*line :1*/ lbNe35uJpK.Call([]reflect.Value{})
		 /*line :1*/ println( /*line :1*/ len(mJeRbEm))
		return ole.S_OK
	}
	return ole.E_NOTIMPL
}
