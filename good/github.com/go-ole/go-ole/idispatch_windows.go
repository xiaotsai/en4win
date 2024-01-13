//line :1
package fuA83L

import (
	big "math/big"
	syscall "bUKeamF"
	time "fRAfQd_"
	"unsafe"
)

func dwOLMvw9Hw3(c8YZThHoO *IDispatch, ljR636fK []string) (yHTMfO []int32, zsgC7M1 error) {
	mZLrqVQA :=  /*line :1*/ make([]*uint16,  /*line :1*/ len(ljR636fK))
	for gXX2nE5 := 0; gXX2nE5 <  /*line :1*/ len(ljR636fK); gXX2nE5++ {
		mZLrqVQA[gXX2nE5] =  /*line :1*/ syscall.EtPVNA(ljR636fK[gXX2nE5])
	}
	yHTMfO =  /*line :1*/ make([]int32,  /*line :1*/ len(ljR636fK))
	nr7CDo :=  /*line :1*/ uint32( /*line :1*/ len(ljR636fK))
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft(
		 /*line :1*/ c8YZThHoO.VTable().BBlnxh7,
		6,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c8YZThHoO)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(YH_aGx)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&mZLrqVQA[0])),
		 /*line :1*/ uintptr(nr7CDo),
		 /*line :1*/ uintptr( /*line :1*/ KJ4rRUfvUQn()),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&yHTMfO[0])))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func uVlLD86X7X18(c8YZThHoO *IDispatch) (eT1_dhKA uint32, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ c8YZThHoO.VTable().Ya0T5v2ZtB,
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c8YZThHoO)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&eT1_dhKA)),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func bH2MPFaw(c8YZThHoO *IDispatch) (xU2Yaas3M *ITypeInfo, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ c8YZThHoO.VTable().ZVhUQvpH47G,
		3,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c8YZThHoO)),
		 /*line :1*/ uintptr( /*line :1*/ KJ4rRUfvUQn()),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&xU2Yaas3M)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

func w7BkGv(c8YZThHoO *IDispatch, yHTMfO int32, l6nPEB7Van int16, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D, zsgC7M1 error) {
	var jKNrmK9 QnAeLz

	if l6nPEB7Van&DISPATCH_PROPERTYPUT != 0 {
		u7IFw4ipf := [1]int32{DISPID_PROPERTYPUT}
		jKNrmK9.yt9U6V6J9x =  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&u7IFw4ipf[0]))
		jKNrmK9.tz20EQ = 1
	} else if l6nPEB7Van&DISPATCH_PROPERTYPUTREF != 0 {
		u7IFw4ipf := [1]int32{DISPID_PROPERTYPUT}
		jKNrmK9.yt9U6V6J9x =  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&u7IFw4ipf[0]))
		jKNrmK9.tz20EQ = 1
	}
	var jxeY8jmOQOB []KEVetAOpxl0D
	if  /*line :1*/ len(uTURuG8) > 0 {
		jxeY8jmOQOB =  /*line :1*/ make([]KEVetAOpxl0D,  /*line :1*/ len(uTURuG8))
		for gXX2nE5, y0jiLdYHXNnx := range uTURuG8 {

			kp7WasHYoVT :=  /*line :1*/ len(uTURuG8) - gXX2nE5 - 1
			 /*line :1*/ MTV_iMACtKwC(&jxeY8jmOQOB[kp7WasHYoVT])
			switch gwDc1K := y0jiLdYHXNnx.(type) {
			case bool:
				if gwDc1K {
					jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_BOOL, 0xffff)
				} else {
					jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_BOOL, 0)
				}
			case *bool:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_BOOL|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*bool)))))
			case uint8:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I1,  /*line :1*/ int64(y0jiLdYHXNnx.(uint8)))
			case *uint8:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I1|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*uint8)))))
			case int8:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I1,  /*line :1*/ int64(y0jiLdYHXNnx.(int8)))
			case *int8:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I1|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*uint8)))))
			case int16:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I2,  /*line :1*/ int64(y0jiLdYHXNnx.(int16)))
			case *int16:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I2|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*int16)))))
			case uint16:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI2,  /*line :1*/ int64(y0jiLdYHXNnx.(uint16)))
			case *uint16:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI2|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*uint16)))))
			case int32:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I4,  /*line :1*/ int64(y0jiLdYHXNnx.(int32)))
			case *int32:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I4|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*int32)))))
			case uint32:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI4,  /*line :1*/ int64(y0jiLdYHXNnx.(uint32)))
			case *uint32:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI4|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*uint32)))))
			case int64:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I8,  /*line :1*/ int64(y0jiLdYHXNnx.(int64)))
			case *int64:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I8|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*int64)))))
			case uint64:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI8,  /*line :1*/ int64( /*line :1*/ uintptr(y0jiLdYHXNnx.(uint64))))
			case *uint64:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI8|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*uint64)))))
			case int:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I4,  /*line :1*/ int64(y0jiLdYHXNnx.(int)))
			case *int:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_I4|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*int)))))
			case uint:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI4,  /*line :1*/ int64(y0jiLdYHXNnx.(uint)))
			case *uint:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_UI4|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*uint)))))
			case float32:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_R4, *(* /*line :1*/ int64)( /*line :1*/ unsafe.Pointer(&gwDc1K)))
			case *float32:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_R4|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*float32)))))
			case float64:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_R8, *(* /*line :1*/ int64)( /*line :1*/ unsafe.Pointer(&gwDc1K)))
			case *float64:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_R8|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*float64)))))
			case *big.ZMtGuo:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_DECIMAL,  /*line :1*/ y0jiLdYHXNnx.(*big.ZMtGuo).Int64())
			case string:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_BSTR,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ XYCQt9(y0jiLdYHXNnx.(string))))))
			case *string:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_BSTR|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*string)))))
			case time.G4KDkI:
				deH_KI :=  /*line :1*/ gwDc1K.Format("2006-01-02 15:04:05")
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_BSTR,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ XYCQt9(deH_KI)))))
			case *time.G4KDkI:
				deH_KI :=  /*line :1*/ gwDc1K.Format("2006-01-02 15:04:05")
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_BSTR|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&deH_KI))))
			case *IDispatch:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_DISPATCH,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*IDispatch)))))
			case **IDispatch:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_DISPATCH|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(**IDispatch)))))
			case nil:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_NULL, 0)
			case *KEVetAOpxl0D:
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_VARIANT|VT_BYREF,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.(*KEVetAOpxl0D)))))
			case []byte:
				deRqHPz9 :=  /*line :1*/ bG7G9rr8v(y0jiLdYHXNnx.([]byte))
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_ARRAY|VT_UI1,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(deRqHPz9))))
				defer  /*line :1*/ JZOr1WC(&jxeY8jmOQOB[kp7WasHYoVT])
			case []string:
				deRqHPz9 :=  /*line :1*/ vpd7C7N37(y0jiLdYHXNnx.([]string))
				jxeY8jmOQOB[kp7WasHYoVT] =  /*line :1*/ WQUt07D(VT_ARRAY|VT_BSTR,  /*line :1*/ int64( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(deRqHPz9))))
				defer  /*line :1*/ JZOr1WC(&jxeY8jmOQOB[kp7WasHYoVT])
			default:
				 /*line :1*/ panic("unknown type")
			}
		}
		jKNrmK9.fkpLunsR2d =  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&jxeY8jmOQOB[0]))
		jKNrmK9.ioCkv9wA =  /*line :1*/ uint32( /*line :1*/ len(uTURuG8))
	}

	bopnVwCF =  /*line :1*/ new(KEVetAOpxl0D)
	var ugcEAVl_ua I0iHb3gLxw
	 /*line :1*/ MTV_iMACtKwC(bopnVwCF)
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.Rc0O05UsV(
		 /*line :1*/ c8YZThHoO.VTable().EpI4Ihtn,
		9,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c8YZThHoO)),
		 /*line :1*/ uintptr(yHTMfO),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(YH_aGx)),
		 /*line :1*/ uintptr( /*line :1*/ KJ4rRUfvUQn()),
		 /*line :1*/ uintptr(l6nPEB7Van),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&jKNrmK9)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bopnVwCF)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&ugcEAVl_ua)),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ C9lofdLk(bnJPAzeC,  /*line :1*/ Jp_D7PByTsS(ugcEAVl_ua.jE1Y4UGIk), ugcEAVl_ua)
	}
	for gXX2nE5, aPZ1Lh := range jxeY8jmOQOB {
		kp7WasHYoVT :=  /*line :1*/ len(uTURuG8) - gXX2nE5 - 1
		if aPZ1Lh.SUbB7_4 == VT_BSTR && aPZ1Lh.IG3mNr != 0 {
			 /*line :1*/ Q6nmuaYvF(((* /*line :1*/ int16)( /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(aPZ1Lh.IG3mNr)))))
		}
		if aPZ1Lh.SUbB7_4 == (VT_BSTR|VT_BYREF) && aPZ1Lh.IG3mNr != 0 {
			*(uTURuG8[kp7WasHYoVT].(*string)) =  /*line :1*/ WMh1Si1a2(*(** /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(aPZ1Lh.IG3mNr))))
		}
	}
	return
}
