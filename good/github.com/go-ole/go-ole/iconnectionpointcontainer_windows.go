//line :1
package fuA83L

import (
	syscall "bUKeamF"
	"unsafe"
)

func (y0jiLdYHXNnx *IConnectionPointContainer) EnumConnectionPoints(kuauvDC interface{}) error {
	return  /*line :1*/ RLCP7BQDRyuR(E_NOTIMPL)
}

func (y0jiLdYHXNnx *IConnectionPointContainer) FindConnectionPoint(c1f0ONah *GUID, nDul5eT **IConnectionPoint) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ y0jiLdYHXNnx.VTable().Z5g5oB2F8VS,
		3,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c1f0ONah)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nDul5eT)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}
