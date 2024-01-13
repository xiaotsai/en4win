//line :1
package fuA83L

import (
	reflect "reflect"
	syscall "bUKeamF"
	utf8 "CuAc0pPZwf"
	"unsafe"
)

var (
	fM6X76mAP5QK	=  /*line :1*/ _j_fj2.NewProc("RoInitialize")
	yVP0MOKVi	=  /*line :1*/ _j_fj2.NewProc("RoActivateInstance")
	ucgFaDY4q5g	=  /*line :1*/ _j_fj2.NewProc("RoGetActivationFactory")
	inMgg3pvApt	=  /*line :1*/ _j_fj2.NewProc("WindowsCreateString")
	r95eFNT	=  /*line :1*/ _j_fj2.NewProc("WindowsDeleteString")
	oXAwlR	=  /*line :1*/ _j_fj2.NewProc("WindowsGetStringRawBuffer")
)

func Qbx05wg2N9Y(qNfYJgJ uint32) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ fM6X76mAP5QK.Call( /*line :1*/ uintptr(qNfYJgJ))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func RBO6rg6kSDL(lcDXHYIAiC8j string) (hQ2aRc *IInspectable, zsgC7M1 error) {
	iV6Ga0, zsgC7M1 :=  /*line :1*/ DeV0qyUAW8K5(lcDXHYIAiC8j)
	if zsgC7M1 != nil {
		return nil, zsgC7M1
	}
	defer  /*line :1*/ SCq5K_T3a7e(iV6Ga0)

	bnJPAzeC, _, _ :=  /*line :1*/ yVP0MOKVi.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iV6Ga0)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&hQ2aRc)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func Fo6dJeNw(lcDXHYIAiC8j string, c1f0ONah *GUID) (hQ2aRc *IInspectable, zsgC7M1 error) {
	iV6Ga0, zsgC7M1 :=  /*line :1*/ DeV0qyUAW8K5(lcDXHYIAiC8j)
	if zsgC7M1 != nil {
		return nil, zsgC7M1
	}
	defer  /*line :1*/ SCq5K_T3a7e(iV6Ga0)

	bnJPAzeC, _, _ :=  /*line :1*/ ucgFaDY4q5g.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iV6Ga0)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c1f0ONah)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&hQ2aRc)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


type OA3qlQMuDJWt uintptr


func DeV0qyUAW8K5(deH_KI string) (fF5eVHvvIt4 OA3qlQMuDJWt, zsgC7M1 error) {
	sIsDZN9 :=  /*line :1*/ syscall.EtPVNA(deH_KI)
	kgY0hespf :=  /*line :1*/ uint32( /*line :1*/ utf8.IaG2lfIQ1LT(deH_KI))
	bnJPAzeC, _, _ :=  /*line :1*/ inMgg3pvApt.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sIsDZN9)),
		 /*line :1*/ uintptr(kgY0hespf),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&fF5eVHvvIt4)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func SCq5K_T3a7e(fF5eVHvvIt4 OA3qlQMuDJWt) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ r95eFNT.Call( /*line :1*/ uintptr(fF5eVHvvIt4))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func (oMVnkTglX OA3qlQMuDJWt) String() string {
	var vY73XeDyE uintptr
	var nxbAEt uint32
	vY73XeDyE, _, _ =  /*line :1*/ oXAwlR.Call(
		 /*line :1*/ uintptr(oMVnkTglX),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&nxbAEt)))

	i3SoqCk4 := reflect.Ei2Iy0{H7OP0hFU: vY73XeDyE, MCzI2qhUA1:  /*line :1*/ int(nxbAEt), A6DWhwcvtSKZ:  /*line :1*/ int(nxbAEt)}
	sIsDZN9 := *(*[] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&i3SoqCk4))
	return  /*line :1*/ syscall.AODVXp8NM3sd(sIsDZN9)
}
