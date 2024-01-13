//line :1
package yAj8QghzkR

import (
	"internal/goarch"
	unsafeheader "EHdBEpZ7y6"
	"unsafe"
)





func TZaSFeAFN1kg(ooBgtTEc_P any) func(dujSIeNg10V, nkLpGFtYnZaZ int) {
	d583PKWfr :=  /*line :1*/ Au73KsHV1lt(ooBgtTEc_P)
	if  /*line :1*/ d583PKWfr.Kind() != Slice {
		 /*line :1*/ panic(&RfP16baj{BH8_yBNG3HjN: "Swapper", Aq5nH3jY:  /*line :1*/ d583PKWfr.Kind()})
	}

	switch  /*line :1*/ d583PKWfr.Len() {
	case 0:
		return func(dujSIeNg10V, nkLpGFtYnZaZ int) {  /*line :1*/ panic("reflect: slice index out of range") }
	case 1:
		return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
			if dujSIeNg10V != 0 || nkLpGFtYnZaZ != 0 {
				 /*line :1*/ panic("reflect: slice index out of range")
			}
		}
	}

	p_IC4wLPVMmq :=  /*line :1*/ d583PKWfr.Type().Elem().gebpy6Q4mW()
	_q1Q5HkJRJC :=  /*line :1*/ p_IC4wLPVMmq.Size()
	a75AFuyR46 := p_IC4wLPVMmq.PtrBytes != 0

	if a75AFuyR46 {
		if _q1Q5HkJRJC == goarch.PtrSize {
			sX9uTNS1gQZ6 := *(*[] /*line :1*/ unsafe.Pointer)(d583PKWfr.k0FQfYmXZ)
			return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
				sX9uTNS1gQZ6[dujSIeNg10V], sX9uTNS1gQZ6[nkLpGFtYnZaZ] = sX9uTNS1gQZ6[nkLpGFtYnZaZ], sX9uTNS1gQZ6[dujSIeNg10V]
			}
		}
		if  /*line :1*/ p_IC4wLPVMmq.Kind() == String {
			aaf9dfo := *(*[] /*line :1*/ string)(d583PKWfr.k0FQfYmXZ)
			return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
				aaf9dfo[dujSIeNg10V], aaf9dfo[nkLpGFtYnZaZ] = aaf9dfo[nkLpGFtYnZaZ], aaf9dfo[dujSIeNg10V]
			}
		}
	} else {
		switch _q1Q5HkJRJC {
		case 8:
			hC9tsOWA := *(*[] /*line :1*/ int64)(d583PKWfr.k0FQfYmXZ)
			return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
				hC9tsOWA[dujSIeNg10V], hC9tsOWA[nkLpGFtYnZaZ] = hC9tsOWA[nkLpGFtYnZaZ], hC9tsOWA[dujSIeNg10V]
			}
		case 4:
			hC9tsOWA := *(*[] /*line :1*/ int32)(d583PKWfr.k0FQfYmXZ)
			return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
				hC9tsOWA[dujSIeNg10V], hC9tsOWA[nkLpGFtYnZaZ] = hC9tsOWA[nkLpGFtYnZaZ], hC9tsOWA[dujSIeNg10V]
			}
		case 2:
			hC9tsOWA := *(*[] /*line :1*/ int16)(d583PKWfr.k0FQfYmXZ)
			return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
				hC9tsOWA[dujSIeNg10V], hC9tsOWA[nkLpGFtYnZaZ] = hC9tsOWA[nkLpGFtYnZaZ], hC9tsOWA[dujSIeNg10V]
			}
		case 1:
			hC9tsOWA := *(*[] /*line :1*/ int8)(d583PKWfr.k0FQfYmXZ)
			return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
				hC9tsOWA[dujSIeNg10V], hC9tsOWA[nkLpGFtYnZaZ] = hC9tsOWA[nkLpGFtYnZaZ], hC9tsOWA[dujSIeNg10V]
			}
		}
	}

	i68MQRL_ := (* /*line :1*/ unsafeheader.AKoOflOZ)(d583PKWfr.k0FQfYmXZ)
	raY6Xg8 :=  /*line :1*/ cUo1caLrG(p_IC4wLPVMmq)

	return func(dujSIeNg10V, nkLpGFtYnZaZ int) {
		if  /*line :1*/ uint(dujSIeNg10V) >=  /*line :1*/ uint(i68MQRL_.AJ2N5B) ||  /*line :1*/ uint(nkLpGFtYnZaZ) >=  /*line :1*/ uint(i68MQRL_.AJ2N5B) {
			 /*line :1*/ panic("reflect: slice index out of range")
		}
		kIXWxBB :=  /*line :1*/ mAjgIGv(i68MQRL_.TanDRHVvuLAL, dujSIeNg10V, _q1Q5HkJRJC, "i < s.Len")
		dIoEi6y :=  /*line :1*/ mAjgIGv(i68MQRL_.TanDRHVvuLAL, nkLpGFtYnZaZ, _q1Q5HkJRJC, "j < s.Len")
		 /*line :1*/ rFMmsMP6H(p_IC4wLPVMmq, raY6Xg8, kIXWxBB)
		 /*line :1*/ rFMmsMP6H(p_IC4wLPVMmq, kIXWxBB, dIoEi6y)
		 /*line :1*/ rFMmsMP6H(p_IC4wLPVMmq, dIoEi6y, raY6Xg8)
	}
}
