//line :1
package reflect

import (
	"internal/bytealg"
	"unsafe"
)





type j5R2QDEa2K struct {
	d4Ksm3E1Hmzm	unsafe.Pointer
	ailEY5yIjn2y	unsafe.Pointer
	xSHI3FcH78L	YJh989LTZsVX
}




func dZaCGt1(j0rnL3oZWlxW, bctwKfC_ Value, bId9CsuBq map[j5R2QDEa2K]bool) bool {
	if ! /*line :1*/ j0rnL3oZWlxW.IsValid() || ! /*line :1*/ bctwKfC_.IsValid() {
		return  /*line :1*/ j0rnL3oZWlxW.IsValid() ==  /*line :1*/ bctwKfC_.IsValid()
	}
	if  /*line :1*/ j0rnL3oZWlxW.Type() !=  /*line :1*/ bctwKfC_.Type() {
		return false
	}

	yTKr2E5jD := func(j0rnL3oZWlxW, bctwKfC_ Value) bool {
		switch  /*line :1*/ j0rnL3oZWlxW.Kind() {
		case Pointer:
			if  /*line :1*/ j0rnL3oZWlxW.lm5DH3().PtrBytes == 0 {

				return false
			}
			fallthrough
		case Map, Slice, Interface:

			return ! /*line :1*/ j0rnL3oZWlxW.IsNil() && ! /*line :1*/ bctwKfC_.IsNil()
		}
		return false
	}

	if  /*line :1*/ yTKr2E5jD(j0rnL3oZWlxW, bctwKfC_) {

		e1S_Vm4BVSW := func(f8NmcatRx Value) unsafe.Pointer {
			switch  /*line :1*/ f8NmcatRx.Kind() {
			case Pointer, Map:
				return  /*line :1*/ f8NmcatRx.t_VyPi()
			default:
				return f8NmcatRx.ptr
			}
		}
		dpRAaUgamKS :=  /*line :1*/ e1S_Vm4BVSW(j0rnL3oZWlxW)
		eSSRoFva0lG2 :=  /*line :1*/ e1S_Vm4BVSW(bctwKfC_)
		if  /*line :1*/ uintptr(dpRAaUgamKS) >  /*line :1*/ uintptr(eSSRoFva0lG2) {

			dpRAaUgamKS, eSSRoFva0lG2 = eSSRoFva0lG2, dpRAaUgamKS
		}

		lm5DH3 :=  /*line :1*/ j0rnL3oZWlxW.Type()
		f8NmcatRx := j5R2QDEa2K{dpRAaUgamKS, eSSRoFva0lG2, lm5DH3}
		if bId9CsuBq[f8NmcatRx] {
			return true
		}

		bId9CsuBq[f8NmcatRx] = true
	}

	switch  /*line :1*/ j0rnL3oZWlxW.Kind() {
	case Array:
		for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ j0rnL3oZWlxW.Len(); hXZpj1nTZ++ {
			if ! /*line :1*/ dZaCGt1( /*line :1*/ j0rnL3oZWlxW.Index(hXZpj1nTZ),  /*line :1*/ bctwKfC_.Index(hXZpj1nTZ), bId9CsuBq) {
				return false
			}
		}
		return true
	case Slice:
		if  /*line :1*/ j0rnL3oZWlxW.IsNil() !=  /*line :1*/ bctwKfC_.IsNil() {
			return false
		}
		if  /*line :1*/ j0rnL3oZWlxW.Len() !=  /*line :1*/ bctwKfC_.Len() {
			return false
		}
		if  /*line :1*/ j0rnL3oZWlxW.UnsafePointer() ==  /*line :1*/ bctwKfC_.UnsafePointer() {
			return true
		}

		if  /*line :1*/ j0rnL3oZWlxW.Type().Elem().Kind() == Uint8 {
			return  /*line :1*/ bytealg.Equal( /*line :1*/ j0rnL3oZWlxW.Bytes(),  /*line :1*/ bctwKfC_.Bytes())
		}
		for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ j0rnL3oZWlxW.Len(); hXZpj1nTZ++ {
			if ! /*line :1*/ dZaCGt1( /*line :1*/ j0rnL3oZWlxW.Index(hXZpj1nTZ),  /*line :1*/ bctwKfC_.Index(hXZpj1nTZ), bId9CsuBq) {
				return false
			}
		}
		return true
	case Interface:
		if  /*line :1*/ j0rnL3oZWlxW.IsNil() ||  /*line :1*/ bctwKfC_.IsNil() {
			return  /*line :1*/ j0rnL3oZWlxW.IsNil() ==  /*line :1*/ bctwKfC_.IsNil()
		}
		return  /*line :1*/ dZaCGt1( /*line :1*/ j0rnL3oZWlxW.Elem(),  /*line :1*/ bctwKfC_.Elem(), bId9CsuBq)
	case Pointer:
		if  /*line :1*/ j0rnL3oZWlxW.UnsafePointer() ==  /*line :1*/ bctwKfC_.UnsafePointer() {
			return true
		}
		return  /*line :1*/ dZaCGt1( /*line :1*/ j0rnL3oZWlxW.Elem(),  /*line :1*/ bctwKfC_.Elem(), bId9CsuBq)
	case Struct:
		for hXZpj1nTZ, krtR1HwEan := 0,  /*line :1*/ j0rnL3oZWlxW.NumField(); hXZpj1nTZ < krtR1HwEan; hXZpj1nTZ++ {
			if ! /*line :1*/ dZaCGt1( /*line :1*/ j0rnL3oZWlxW.Field(hXZpj1nTZ),  /*line :1*/ bctwKfC_.Field(hXZpj1nTZ), bId9CsuBq) {
				return false
			}
		}
		return true
	case Map:
		if  /*line :1*/ j0rnL3oZWlxW.IsNil() !=  /*line :1*/ bctwKfC_.IsNil() {
			return false
		}
		if  /*line :1*/ j0rnL3oZWlxW.Len() !=  /*line :1*/ bctwKfC_.Len() {
			return false
		}
		if  /*line :1*/ j0rnL3oZWlxW.UnsafePointer() ==  /*line :1*/ bctwKfC_.UnsafePointer() {
			return true
		}
		for _, lYOqUQga := range  /*line :1*/ j0rnL3oZWlxW.MapKeys() {
			gEhdUz2aj :=  /*line :1*/ j0rnL3oZWlxW.MapIndex(lYOqUQga)
			iZjv8mS6P :=  /*line :1*/ bctwKfC_.MapIndex(lYOqUQga)
			if ! /*line :1*/ gEhdUz2aj.IsValid() || ! /*line :1*/ iZjv8mS6P.IsValid() || ! /*line :1*/ dZaCGt1(gEhdUz2aj, iZjv8mS6P, bId9CsuBq) {
				return false
			}
		}
		return true
	case Func:
		if  /*line :1*/ j0rnL3oZWlxW.IsNil() &&  /*line :1*/ bctwKfC_.IsNil() {
			return true
		}

		return false
	case Int, Int8, Int16, Int32, Int64:
		return  /*line :1*/ j0rnL3oZWlxW.Int() ==  /*line :1*/ bctwKfC_.Int()
	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		return  /*line :1*/ j0rnL3oZWlxW.Uint() ==  /*line :1*/ bctwKfC_.Uint()
	case String:
		return  /*line :1*/ j0rnL3oZWlxW.String() ==  /*line :1*/ bctwKfC_.String()
	case Bool:
		return  /*line :1*/ j0rnL3oZWlxW.Bool() ==  /*line :1*/ bctwKfC_.Bool()
	case Float32, Float64:
		return  /*line :1*/ j0rnL3oZWlxW.Float() ==  /*line :1*/ bctwKfC_.Float()
	case Complex64, Complex128:
		return  /*line :1*/ j0rnL3oZWlxW.Complex() ==  /*line :1*/ bctwKfC_.Complex()
	default:

		return  /*line :1*/ mySQReh3uV(j0rnL3oZWlxW, false) ==  /*line :1*/ mySQReh3uV(bctwKfC_, false)
	}
}




















































func Ofi8i_PJan(uiTlN8bdKm, fojrp9VAi1ea any) bool {
	if uiTlN8bdKm == nil || fojrp9VAi1ea == nil {
		return uiTlN8bdKm == fojrp9VAi1ea
	}
	j0rnL3oZWlxW :=  /*line :1*/ SdHoaIQl(uiTlN8bdKm)
	bctwKfC_ :=  /*line :1*/ SdHoaIQl(fojrp9VAi1ea)
	if  /*line :1*/ j0rnL3oZWlxW.Type() !=  /*line :1*/ bctwKfC_.Type() {
		return false
	}
	return  /*line :1*/ dZaCGt1(j0rnL3oZWlxW, bctwKfC_,  /*line :1*/ make(map[j5R2QDEa2K]bool))
}
