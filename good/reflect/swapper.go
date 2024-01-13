//line :1
package reflect

import (
	"internal/abi"
	"internal/goarch"
	unsafeheader "EHdBEpZ7y6"
	"unsafe"
)





func QVWFICP(qMTDS7fnbc any) func(hXZpj1nTZ, kNVeJaGxs int) {
	f8NmcatRx :=  /*line :1*/ SdHoaIQl(qMTDS7fnbc)
	if  /*line :1*/ f8NmcatRx.Kind() != Slice {
		 /*line :1*/ panic(&R_QNJpzecp{Method: "Swapper", N76ipnT:  /*line :1*/ f8NmcatRx.Kind()})
	}

	switch  /*line :1*/ f8NmcatRx.Len() {
	case 0:
		return func(hXZpj1nTZ, kNVeJaGxs int) {  /*line :1*/ panic("reflect: slice index out of range") }
	case 1:
		return func(hXZpj1nTZ, kNVeJaGxs int) {
			if hXZpj1nTZ != 0 || kNVeJaGxs != 0 {
				 /*line :1*/ panic("reflect: slice index out of range")
			}
		}
	}

	lm5DH3 :=  /*line :1*/ f8NmcatRx.Type().Elem().z6hGxTABM1()
	wnTnByH :=  /*line :1*/ lm5DH3.Size()
	yatSKKXlq := lm5DH3.PtrBytes != 0

	if yatSKKXlq {
		if wnTnByH == goarch.PtrSize {
			q5rOo0DP := *(*[] /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr)
			return func(hXZpj1nTZ, kNVeJaGxs int) {
				q5rOo0DP[hXZpj1nTZ], q5rOo0DP[kNVeJaGxs] = q5rOo0DP[kNVeJaGxs], q5rOo0DP[hXZpj1nTZ]
			}
		}
		if  /*line :1*/ lm5DH3.Kind() == abi.String {
			sr4mKYvFuU9 := *(*[] /*line :1*/ string)(f8NmcatRx.ptr)
			return func(hXZpj1nTZ, kNVeJaGxs int) {
				sr4mKYvFuU9[hXZpj1nTZ], sr4mKYvFuU9[kNVeJaGxs] = sr4mKYvFuU9[kNVeJaGxs], sr4mKYvFuU9[hXZpj1nTZ]
			}
		}
	} else {
		switch wnTnByH {
		case 8:
			mEKr2xGa := *(*[] /*line :1*/ int64)(f8NmcatRx.ptr)
			return func(hXZpj1nTZ, kNVeJaGxs int) {
				mEKr2xGa[hXZpj1nTZ], mEKr2xGa[kNVeJaGxs] = mEKr2xGa[kNVeJaGxs], mEKr2xGa[hXZpj1nTZ]
			}
		case 4:
			mEKr2xGa := *(*[] /*line :1*/ int32)(f8NmcatRx.ptr)
			return func(hXZpj1nTZ, kNVeJaGxs int) {
				mEKr2xGa[hXZpj1nTZ], mEKr2xGa[kNVeJaGxs] = mEKr2xGa[kNVeJaGxs], mEKr2xGa[hXZpj1nTZ]
			}
		case 2:
			mEKr2xGa := *(*[] /*line :1*/ int16)(f8NmcatRx.ptr)
			return func(hXZpj1nTZ, kNVeJaGxs int) {
				mEKr2xGa[hXZpj1nTZ], mEKr2xGa[kNVeJaGxs] = mEKr2xGa[kNVeJaGxs], mEKr2xGa[hXZpj1nTZ]
			}
		case 1:
			mEKr2xGa := *(*[] /*line :1*/ int8)(f8NmcatRx.ptr)
			return func(hXZpj1nTZ, kNVeJaGxs int) {
				mEKr2xGa[hXZpj1nTZ], mEKr2xGa[kNVeJaGxs] = mEKr2xGa[kNVeJaGxs], mEKr2xGa[hXZpj1nTZ]
			}
		}
	}

	kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
	u1ertr :=  /*line :1*/ qscv_5(lm5DH3)

	return func(hXZpj1nTZ, kNVeJaGxs int) {
		if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint(kXTvhUI4Zgg5.AJ2N5B) ||  /*line :1*/ uint(kNVeJaGxs) >=  /*line :1*/ uint(kXTvhUI4Zgg5.AJ2N5B) {
			 /*line :1*/ panic("reflect: slice index out of range")
		}
		gEhdUz2aj :=  /*line :1*/ rfJB4Et(kXTvhUI4Zgg5.TanDRHVvuLAL, hXZpj1nTZ, wnTnByH, "i < s.Len")
		iZjv8mS6P :=  /*line :1*/ rfJB4Et(kXTvhUI4Zgg5.TanDRHVvuLAL, kNVeJaGxs, wnTnByH, "j < s.Len")
		 /*line :1*/ l26oiKGd(lm5DH3, u1ertr, gEhdUz2aj)
		 /*line :1*/ l26oiKGd(lm5DH3, gEhdUz2aj, iZjv8mS6P)
		 /*line :1*/ l26oiKGd(lm5DH3, iZjv8mS6P, u1ertr)
	}
}
