//line :1
package fuA83L

import (
	syscall "bUKeamF"
	"unsafe"
)

func (y0jiLdYHXNnx *ITypeInfo) GetTypeAttr() (mOgOUi *CFaF1Ke, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ uintptr( /*line :1*/ y0jiLdYHXNnx.VTable().I9s9GMsZd),
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&mOgOUi)),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}
