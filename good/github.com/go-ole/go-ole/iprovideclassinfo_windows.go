//line :1
package fuA83L

import (
	syscall "bUKeamF"
	"unsafe"
)

func cl6GDk7T6(c8YZThHoO *IProvideClassInfo) (xU2Yaas3M *ITypeInfo, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ c8YZThHoO.VTable().BmlCvaPHRDhB,
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c8YZThHoO)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&xU2Yaas3M)),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}
