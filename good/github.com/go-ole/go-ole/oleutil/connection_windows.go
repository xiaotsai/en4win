//line :1
package uBCMx3OoU

import (
	reflect "reflect"
	syscall "bUKeamF"
	"unsafe"

	ole "fuA83L"
)


func U0f01__(hGnmPcU *ole.IDispatch, kM7ECagotLCo *ole.GUID, usacailNCq interface{}) (a4jQt6a6 uint32, ylNwKji6NI error) {
	bwwyhlUjNv, ylNwKji6NI :=  /*line :1*/ hGnmPcU.QueryInterface(ole.D3Vh84Kk)
	if ylNwKji6NI != nil {
		return
	}

	mNpr0k := (* /*line :1*/ ole.IConnectionPointContainer)( /*line :1*/ unsafe.Pointer(bwwyhlUjNv))
	var hssCO6 *ole.IConnectionPoint
	ylNwKji6NI =  /*line :1*/ mNpr0k.FindConnectionPoint(kM7ECagotLCo, &hssCO6)
	if ylNwKji6NI != nil {
		return
	}
	if nfrMYUO, rmDZGevOH := usacailNCq.(*ole.IUnknown); rmDZGevOH {
		a4jQt6a6, ylNwKji6NI =  /*line :1*/ hssCO6.Advise(nfrMYUO)
		 /*line :1*/ mNpr0k.Release()
		if ylNwKji6NI != nil {
			return
		}
	}
	cBPFBDh :=  /*line :1*/ reflect.SdHoaIQl(hGnmPcU).Elem()
	if  /*line :1*/ cBPFBDh.Type().Kind() == reflect.Struct {
		gx7RHEaAHc := &ar_mWKlub{}
		gx7RHEaAHc.cpNdyUaVENy = &f_OiDUX{}
		gx7RHEaAHc.cpNdyUaVENy.wHAJNe =  /*line :1*/ syscall.VSPtNpjXc4(kMSCaur)
		gx7RHEaAHc.cpNdyUaVENy.fuVf5d7Yf7 =  /*line :1*/ syscall.VSPtNpjXc4(auCx1_dDF8qu)
		gx7RHEaAHc.cpNdyUaVENy.a_B0MzQg7J =  /*line :1*/ syscall.VSPtNpjXc4(krKwiq5ldnCn)
		gx7RHEaAHc.cpNdyUaVENy.a_mYlb =  /*line :1*/ syscall.VSPtNpjXc4(fVkBn0q5GbO)
		gx7RHEaAHc.cpNdyUaVENy.q_KkJi9_o =  /*line :1*/ syscall.VSPtNpjXc4(ePZAQyooJ2)
		gx7RHEaAHc.cpNdyUaVENy.iqvf392iF =  /*line :1*/ syscall.VSPtNpjXc4(nAWjalYrV)
		gx7RHEaAHc.cpNdyUaVENy.kbDaVyp =  /*line :1*/ syscall.VSPtNpjXc4(gMeFKNUYgE)
		gx7RHEaAHc.bCqdbzoT0 = hGnmPcU
		gx7RHEaAHc.drlgzSLc = kM7ECagotLCo
		a4jQt6a6, ylNwKji6NI =  /*line :1*/ hssCO6.Advise((* /*line :1*/ ole.IUnknown)( /*line :1*/ unsafe.Pointer(gx7RHEaAHc)))
		 /*line :1*/ mNpr0k.Release()
		if ylNwKji6NI != nil {
			 /*line :1*/ hssCO6.Release()
			return
		}
		return
	}

	 /*line :1*/ mNpr0k.Release()

	return 0,  /*line :1*/ ole.RLCP7BQDRyuR(ole.E_INVALIDARG)
}
