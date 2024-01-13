//line :1
package fuA83L

import (
	utf16 "DtJCLKevRp"
	"unsafe"
)






func IPqskAs0K(aG219XWbeGz_ string) (aV1zPhVpati0 *GUID, zsgC7M1 error) {
	aV1zPhVpati0, zsgC7M1 =  /*line :1*/ AtGYsf(aG219XWbeGz_)
	if zsgC7M1 != nil {
		aV1zPhVpati0, zsgC7M1 =  /*line :1*/ QFLzLQWmE(aG219XWbeGz_)
		if zsgC7M1 != nil {
			return
		}
	}
	return
}


func X1aGOJl02G(cUagkN *byte) string {
	hwMamRU := (*[10000] /*line :1*/ uint8)( /*line :1*/ unsafe.Pointer(cUagkN))
	gXX2nE5 := 0
	for hwMamRU[gXX2nE5] != 0 {
		gXX2nE5++
	}
	return  /*line :1*/ string(hwMamRU[:gXX2nE5])
}




func AvM6zX(cUagkN *uint16) string {
	return  /*line :1*/ WMh1Si1a2(cUagkN)
}


func WMh1Si1a2(cUagkN *uint16) string {
	if cUagkN == nil {
		return ""
	}

	nXciaovStia3 :=  /*line :1*/ c4vxxHr70MiM(cUagkN)
	hwMamRU :=  /*line :1*/ make([]uint16, nXciaovStia3)

	jaRy_z :=  /*line :1*/ unsafe.Pointer(cUagkN)

	for gXX2nE5 := 0; gXX2nE5 <  /*line :1*/ int(nXciaovStia3); gXX2nE5++ {
		hwMamRU[gXX2nE5] = *(* /*line :1*/ uint16)(jaRy_z)
		jaRy_z =  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(jaRy_z) + 2)
	}

	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(hwMamRU))
}


func Jp_D7PByTsS(cUagkN *uint16) string {
	if cUagkN == nil {
		return ""
	}
	nXciaovStia3 :=  /*line :1*/ TkFtieTh_k((* /*line :1*/ int16)( /*line :1*/ unsafe.Pointer(cUagkN)))
	hwMamRU :=  /*line :1*/ make([]uint16, nXciaovStia3)

	jaRy_z :=  /*line :1*/ unsafe.Pointer(cUagkN)

	for gXX2nE5 := 0; gXX2nE5 <  /*line :1*/ int(nXciaovStia3); gXX2nE5++ {
		hwMamRU[gXX2nE5] = *(* /*line :1*/ uint16)(jaRy_z)
		jaRy_z =  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(jaRy_z) + 2)
	}
	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(hwMamRU))
}


func c4vxxHr70MiM(cUagkN *uint16) (nXciaovStia3 int64) {
	if cUagkN == nil {
		return 0
	}

	jaRy_z :=  /*line :1*/ unsafe.Pointer(cUagkN)

	for gXX2nE5 := 0; ; gXX2nE5++ {
		if 0 == *(* /*line :1*/ uint16)(jaRy_z) {
			nXciaovStia3 =  /*line :1*/ int64(gXX2nE5)
			break
		}
		jaRy_z =  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(jaRy_z) + 2)
	}
	return
}


func mBJa4FrzpmT(bnJPAzeC uintptr, pWtPB3ETxrK1 uintptr, r76qCBNQkvis error) (zsgC7M1 error) {
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}
