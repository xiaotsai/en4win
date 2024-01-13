//line :1
package fuA83L

import (
	syscall "bUKeamF"
	"unsafe"
)

func (y0jiLdYHXNnx *IConnectionPoint) GetConnectionInterface(d8qkYJu1w8h **GUID) int32 {

	return  /*line :1*/ o1qH5JwREvw2((* /*line :1*/ IUnknown)( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)))
}

func (y0jiLdYHXNnx *IConnectionPoint) Advise(e9b4y2W *IUnknown) (uihp1XXXr uint32, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ y0jiLdYHXNnx.VTable().NX5cFm,
		3,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(e9b4y2W)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&uihp1XXXr)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func (y0jiLdYHXNnx *IConnectionPoint) Unadvise(uihp1XXXr uint32) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ y0jiLdYHXNnx.VTable().Tqkg5U0aTGbL,
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)),
		 /*line :1*/ uintptr(uihp1XXXr),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func (y0jiLdYHXNnx *IConnectionPoint) EnumConnections(cUagkN *unsafe.Pointer) error {
	return  /*line :1*/ RLCP7BQDRyuR(E_NOTIMPL)
}
