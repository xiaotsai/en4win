//line :1
package fuA83L

import (
	syscall "bUKeamF"
	"unsafe"
)

func (xIxQm4DCSzB *IEnumVARIANT) Clone() (rtcvDuTS9B *IEnumVARIANT, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ xIxQm4DCSzB.VTable().C_QXtUavm,
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xIxQm4DCSzB)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&rtcvDuTS9B)),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func (xIxQm4DCSzB *IEnumVARIANT) Reset() (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ xIxQm4DCSzB.VTable().ZF5hbILirlOE,
		1,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xIxQm4DCSzB)),
		0,
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func (xIxQm4DCSzB *IEnumVARIANT) Skip(n6aP_rhY uint) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ xIxQm4DCSzB.VTable().K2TbXm5HnujN,
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xIxQm4DCSzB)),
		 /*line :1*/ uintptr(n6aP_rhY),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func (xIxQm4DCSzB *IEnumVARIANT) Next(n6aP_rhY uint) (j7O2yy KEVetAOpxl0D, nXciaovStia3 uint, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft(
		 /*line :1*/ xIxQm4DCSzB.VTable().C9KKf3aiwWc,
		4,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xIxQm4DCSzB)),
		 /*line :1*/ uintptr(n6aP_rhY),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&j7O2yy)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&nXciaovStia3)),
		0,
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}
