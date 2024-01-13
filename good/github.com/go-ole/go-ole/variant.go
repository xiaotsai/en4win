//line :1
package fuA83L

import "unsafe"


func WQUt07D(tCRLo1W5m9va JtqZSr3B, s878_aqscvO int64) KEVetAOpxl0D {
	return KEVetAOpxl0D{SUbB7_4: tCRLo1W5m9va, IG3mNr: s878_aqscvO}
}


func (y0jiLdYHXNnx *KEVetAOpxl0D) ToIUnknown() *IUnknown {
	if y0jiLdYHXNnx.SUbB7_4 != VT_UNKNOWN {
		return nil
	}
	return (* /*line :1*/ IUnknown)( /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(y0jiLdYHXNnx.IG3mNr)))
}


func (y0jiLdYHXNnx *KEVetAOpxl0D) ToIDispatch() *IDispatch {
	if y0jiLdYHXNnx.SUbB7_4 != VT_DISPATCH {
		return nil
	}
	return (* /*line :1*/ IDispatch)( /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(y0jiLdYHXNnx.IG3mNr)))
}


func (y0jiLdYHXNnx *KEVetAOpxl0D) ToArray() *JA6nr0c8Wi {
	if y0jiLdYHXNnx.SUbB7_4 != VT_SAFEARRAY {
		if y0jiLdYHXNnx.SUbB7_4&VT_ARRAY == 0 {
			return nil
		}
	}
	var hH1fRdl3f *V8vhLAHnL4 = (* /*line :1*/ V8vhLAHnL4)( /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(y0jiLdYHXNnx.IG3mNr)))
	return &JA6nr0c8Wi{hH1fRdl3f}
}


func (y0jiLdYHXNnx *KEVetAOpxl0D) ToString() string {
	if y0jiLdYHXNnx.SUbB7_4 != VT_BSTR {
		return ""
	}
	return  /*line :1*/ Jp_D7PByTsS(*(** /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx.IG3mNr)))
}


func (y0jiLdYHXNnx *KEVetAOpxl0D) Clear() error {
	return  /*line :1*/ JZOr1WC(y0jiLdYHXNnx)
}








func (y0jiLdYHXNnx *KEVetAOpxl0D) Value() interface{} {
	switch y0jiLdYHXNnx.SUbB7_4 {
	case VT_I1:
		return  /*line :1*/ int8(y0jiLdYHXNnx.IG3mNr)
	case VT_UI1:
		return  /*line :1*/ uint8(y0jiLdYHXNnx.IG3mNr)
	case VT_I2:
		return  /*line :1*/ int16(y0jiLdYHXNnx.IG3mNr)
	case VT_UI2:
		return  /*line :1*/ uint16(y0jiLdYHXNnx.IG3mNr)
	case VT_I4:
		return  /*line :1*/ int32(y0jiLdYHXNnx.IG3mNr)
	case VT_UI4:
		return  /*line :1*/ uint32(y0jiLdYHXNnx.IG3mNr)
	case VT_I8:
		return  /*line :1*/ int64(y0jiLdYHXNnx.IG3mNr)
	case VT_UI8:
		return  /*line :1*/ uint64(y0jiLdYHXNnx.IG3mNr)
	case VT_INT:
		return  /*line :1*/ int(y0jiLdYHXNnx.IG3mNr)
	case VT_UINT:
		return  /*line :1*/ uint(y0jiLdYHXNnx.IG3mNr)
	case VT_INT_PTR:
		return  /*line :1*/ uintptr(y0jiLdYHXNnx.IG3mNr)
	case VT_UINT_PTR:
		return  /*line :1*/ uintptr(y0jiLdYHXNnx.IG3mNr)
	case VT_R4:
		return *(* /*line :1*/ float32)( /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx.IG3mNr))
	case VT_R8:
		return *(* /*line :1*/ float64)( /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx.IG3mNr))
	case VT_BSTR:
		return  /*line :1*/ y0jiLdYHXNnx.ToString()
	case VT_DATE:

		b9Flo0 :=  /*line :1*/ uint64(y0jiLdYHXNnx.IG3mNr)
		bgIHRx6HIm, zsgC7M1 :=  /*line :1*/ NT_HQpy3w(b9Flo0)
		if zsgC7M1 != nil {
			return  /*line :1*/ float64(y0jiLdYHXNnx.IG3mNr)
		}
		return bgIHRx6HIm
	case VT_UNKNOWN:
		return  /*line :1*/ y0jiLdYHXNnx.ToIUnknown()
	case VT_DISPATCH:
		return  /*line :1*/ y0jiLdYHXNnx.ToIDispatch()
	case VT_BOOL:
		return y0jiLdYHXNnx.IG3mNr != 0
	}
	return nil
}
