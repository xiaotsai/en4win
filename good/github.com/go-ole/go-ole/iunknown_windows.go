//line :1
package fuA83L

import (
	reflect "reflect"
	syscall "bUKeamF"
	"unsafe"
)

func eRuA_UMJfKp(h3TBmo interface{}, r8V4jqlunkmQ uintptr, yxjNU385 *GUID, qVhHBZar4a6 interface{}) (zsgC7M1 error) {
	chpzh1m3 :=  /*line :1*/ reflect.SdHoaIQl(h3TBmo).Elem()
	yUEtpqD :=  /*line :1*/ reflect.SdHoaIQl(qVhHBZar4a6).Elem()

	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		r8V4jqlunkmQ,
		3,
		 /*line :1*/ chpzh1m3.UnsafeAddr(),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yxjNU385)),
		 /*line :1*/ yUEtpqD.Addr().Pointer())
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func sV8DJK(iGHTYs7q8FBk *IUnknown, c1f0ONah *GUID) (c8YZThHoO *IDispatch, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ iGHTYs7q8FBk.VTable().XVGclIGZIiw,
		3,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iGHTYs7q8FBk)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c1f0ONah)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&c8YZThHoO)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func y8tTraIf7(iGHTYs7q8FBk *IUnknown) int32 {
	y7wx667nI, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ iGHTYs7q8FBk.VTable().GkT7ZP,
		1,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iGHTYs7q8FBk)),
		0,
		0)
	return  /*line :1*/ int32(y7wx667nI)
}

func o1qH5JwREvw2(iGHTYs7q8FBk *IUnknown) int32 {
	y7wx667nI, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ iGHTYs7q8FBk.VTable().Aq2pNFiJR,
		1,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iGHTYs7q8FBk)),
		0,
		0)
	return  /*line :1*/ int32(y7wx667nI)
}
