//line :1


package yAj8QghzkR

import (
	"internal/abi"
	"unsafe"
)












type LrcFby interface {

	
	
	Name() string

	
	
	
	
	
	PkgPath() string

	
	
	Size() uintptr

	
	Kind() FZ0rl1z

	
	Implements(tvuza9gJ0 LrcFby) bool

	
	AssignableTo(tvuza9gJ0 LrcFby) bool

	
	Comparable() bool

	
	
	
	
	
	String() string

	
	
	Elem() LrcFby

	gebpy6Q4mW() *abi.Type
	o67BjECFG8() *jOilAg
}



type FZ0rl1z = abi.Kind

const Ptr = abi.Pointer

const (
	
	Interface	= abi.Interface
	Slice	= abi.Slice
	String	= abi.String
	Struct	= abi.Struct
)

type sUmjH5 = abi.NameOff
type cnO61WNA_NK = abi.TypeOff
type bpFvO0G = abi.TextOff

type eHKSMvk3XaJH struct {
	*abi.Type
}





type jOilAg = abi.UncommonType


type jCeH0iHPi9 = abi.ArrayType


type ezyeG42w = abi.ChanType

type jVlfi1v = abi.FuncType

type fh7O1l0V_Rdz = abi.InterfaceType


type aq3eQkndnW struct {
	eHKSMvk3XaJH
	O0dYgRFm	*abi.Type	
	Rts0PS05m	*abi.Type	
	K0XLiOzKl	*abi.Type	
	
	Zs5GvZHkd1i7	func(unsafe.Pointer, uintptr) uintptr
	L5cfMS	uint8	
	XVFstVgf	uint8	
	PTkFDNMk	uint16	
	GZrV_x9dibVq	uint32
}


type naf2N4qWKL9 = abi.PtrType


type otda_NzrB = abi.SliceType


type iIRpWLCsj = abi.StructType
























type ja7ClO7 struct {
	atl063P *byte
}

func (zAGif2gz0 ja7ClO7) nvswhUCr3(iegMpqa int, jo0rnS string) *byte {
	return (* /*line :1*/ byte)( /*line :1*/ pOBZZZ( /*line :1*/ unsafe.Pointer(zAGif2gz0.atl063P),  /*line :1*/ uintptr(iegMpqa), jo0rnS))
}

func (zAGif2gz0 ja7ClO7) rs6U8jYOfa() bool {
	return (*zAGif2gz0.atl063P)&(1<<0) != 0
}

func (zAGif2gz0 ja7ClO7) lO6ZERoP4b() bool {
	return (*zAGif2gz0.atl063P)&(1<<1) != 0
}

func (zAGif2gz0 ja7ClO7) swG7n7L0KY() bool {
	return (*zAGif2gz0.atl063P)&(1<<3) != 0
}



func (zAGif2gz0 ja7ClO7) eNhcNAs(iegMpqa int) (int, int) {
	d583PKWfr := 0
	for dujSIeNg10V := 0; ; dujSIeNg10V++ {
		jJCfpek := * /*line :1*/ zAGif2gz0.nvswhUCr3(iegMpqa+dujSIeNg10V, "read varint")
		d583PKWfr +=  /*line :1*/ int(jJCfpek&0x7f) << (7 * dujSIeNg10V)
		if jJCfpek&0x80 == 0 {
			return dujSIeNg10V + 1, d583PKWfr
		}
	}
}

func (zAGif2gz0 ja7ClO7) ja7ClO7() string {
	if zAGif2gz0.atl063P == nil {
		return ""
	}
	dujSIeNg10V, fO4YO5 :=  /*line :1*/ zAGif2gz0.eNhcNAs(1)
	return  /*line :1*/ unsafe.String( /*line :1*/ zAGif2gz0.nvswhUCr3(1+dujSIeNg10V, "non-empty string"), fO4YO5)
}

func (zAGif2gz0 ja7ClO7) ugGhBEo_() string {
	if ! /*line :1*/ zAGif2gz0.lO6ZERoP4b() {
		return ""
	}
	dujSIeNg10V, fO4YO5 :=  /*line :1*/ zAGif2gz0.eNhcNAs(1)
	tDpwOV6XGAJA, daaW3Mb :=  /*line :1*/ zAGif2gz0.eNhcNAs(1 + dujSIeNg10V + fO4YO5)
	return  /*line :1*/ unsafe.String( /*line :1*/ zAGif2gz0.nvswhUCr3(1+dujSIeNg10V+fO4YO5+tDpwOV6XGAJA, "non-empty string"), daaW3Mb)
}

func ggkGz79(zAGif2gz0 abi.Name) string {
	if zAGif2gz0.Bytes == nil || * /*line :1*/ zAGif2gz0.DataChecked(0, "name flag field")&(1<<2) == 0 {
		return ""
	}
	dujSIeNg10V, fO4YO5 :=  /*line :1*/ zAGif2gz0.ReadVarint(1)
	iegMpqa := 1 + dujSIeNg10V + fO4YO5
	if  /*line :1*/ zAGif2gz0.HasTag() {
		tDpwOV6XGAJA, daaW3Mb :=  /*line :1*/ zAGif2gz0.ReadVarint(iegMpqa)
		iegMpqa += tDpwOV6XGAJA + daaW3Mb
	}
	var sUmjH5 int32

	 /*line :1*/ copy((*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&sUmjH5))[:], (*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer( /*line :1*/ zAGif2gz0.DataChecked(iegMpqa, "name offset field")))[:])
	hvrTMuMd8n3 := ja7ClO7{(* /*line :1*/ byte)( /*line :1*/ odBpGgGE8NS( /*line :1*/ unsafe.Pointer(zAGif2gz0.Bytes), sUmjH5))}
	return  /*line :1*/ hvrTMuMd8n3.ja7ClO7()
}




func aWLnpKj(v0aXUS unsafe.Pointer, iegMpqa int32) unsafe.Pointer




func odBpGgGE8NS(eHKSMvk3XaJH unsafe.Pointer, iegMpqa int32) unsafe.Pointer

func (gc1v5B2w4C eHKSMvk3XaJH) sUmjH5(iegMpqa sUmjH5) abi.Name {
	return abi.Name{Bytes: (* /*line :1*/ byte)( /*line :1*/ aWLnpKj( /*line :1*/ unsafe.Pointer(gc1v5B2w4C.Type),  /*line :1*/ int32(iegMpqa)))}
}

func (gc1v5B2w4C eHKSMvk3XaJH) cnO61WNA_NK(iegMpqa cnO61WNA_NK) *abi.Type {
	return (* /*line :1*/ abi.Type)( /*line :1*/ odBpGgGE8NS( /*line :1*/ unsafe.Pointer(gc1v5B2w4C.Type),  /*line :1*/ int32(iegMpqa)))
}

func (gc1v5B2w4C eHKSMvk3XaJH) o67BjECFG8() *jOilAg {
	return  /*line :1*/ gc1v5B2w4C.Uncommon()
}

func (gc1v5B2w4C eHKSMvk3XaJH) String() string {
	i68MQRL_ :=  /*line :1*/ gc1v5B2w4C.sUmjH5(gc1v5B2w4C.Str).Name()
	if gc1v5B2w4C.TFlag&abi.TFlagExtraStar != 0 {
		return i68MQRL_[1:]
	}
	return i68MQRL_
}

func (gc1v5B2w4C eHKSMvk3XaJH) gebpy6Q4mW() *abi.Type	{ return gc1v5B2w4C.Type }

func (gc1v5B2w4C eHKSMvk3XaJH) lX7O8w8() []abi.Method {
	qjwBeq8h :=  /*line :1*/ gc1v5B2w4C.o67BjECFG8()
	if qjwBeq8h == nil {
		return nil
	}
	return  /*line :1*/ qjwBeq8h.ExportedMethods()
}

func (gc1v5B2w4C eHKSMvk3XaJH) NumMethod() int {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.InterfaceType()
	if tq_UCmMaf0QU != nil {
		return  /*line :1*/ tq_UCmMaf0QU.NumMethod()
	}
	return  /*line :1*/ len( /*line :1*/ gc1v5B2w4C.lX7O8w8())
}

func (gc1v5B2w4C eHKSMvk3XaJH) PkgPath() string {
	if gc1v5B2w4C.TFlag&abi.TFlagNamed == 0 {
		return ""
	}
	qjwBeq8h :=  /*line :1*/ gc1v5B2w4C.o67BjECFG8()
	if qjwBeq8h == nil {
		return ""
	}
	return  /*line :1*/ gc1v5B2w4C.sUmjH5(qjwBeq8h.PkgPath).Name()
}

func (gc1v5B2w4C eHKSMvk3XaJH) Name() string {
	if ! /*line :1*/ gc1v5B2w4C.HasName() {
		return ""
	}
	i68MQRL_ :=  /*line :1*/ gc1v5B2w4C.String()
	dujSIeNg10V :=  /*line :1*/ len(i68MQRL_) - 1
	xdloqQSLv := 0
	for dujSIeNg10V >= 0 && (i68MQRL_[dujSIeNg10V] != '.' || xdloqQSLv != 0) {
		switch i68MQRL_[dujSIeNg10V] {
		case ']':
			xdloqQSLv++
		case '[':
			xdloqQSLv--
		}
		dujSIeNg10V--
	}
	return i68MQRL_[dujSIeNg10V+1:]
}

func qK4nvEMY(gc1v5B2w4C *abi.Type) eHKSMvk3XaJH {
	return eHKSMvk3XaJH{gc1v5B2w4C}
}

func mqP6CIz2(gc1v5B2w4C *abi.Type) *abi.Type {
	_MqzhL :=  /*line :1*/ gc1v5B2w4C.Elem()
	if _MqzhL != nil {
		return _MqzhL
	}
	 /*line :1*/ panic("reflect: Elem of invalid type " +  /*line :1*/ qK4nvEMY(gc1v5B2w4C).String())
}

func (gc1v5B2w4C eHKSMvk3XaJH) Elem() LrcFby {
	return  /*line :1*/ juqSJNrj( /*line :1*/ mqP6CIz2( /*line :1*/ gc1v5B2w4C.gebpy6Q4mW()))
}

func (gc1v5B2w4C eHKSMvk3XaJH) In(dujSIeNg10V int) LrcFby {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.FuncType()
	if tq_UCmMaf0QU == nil {
		 /*line :1*/ panic("reflect: In of non-func type")
	}
	return  /*line :1*/ juqSJNrj( /*line :1*/ tq_UCmMaf0QU.InSlice()[dujSIeNg10V])
}

func (gc1v5B2w4C eHKSMvk3XaJH) Key() LrcFby {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.MapType()
	if tq_UCmMaf0QU == nil {
		 /*line :1*/ panic("reflect: Key of non-map type")
	}
	return  /*line :1*/ juqSJNrj(tq_UCmMaf0QU.Key)
}

func (gc1v5B2w4C eHKSMvk3XaJH) Len() int {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.ArrayType()
	if tq_UCmMaf0QU == nil {
		 /*line :1*/ panic("reflect: Len of non-array type")
	}
	return  /*line :1*/ int(tq_UCmMaf0QU.Len)
}

func (gc1v5B2w4C eHKSMvk3XaJH) NumField() int {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.StructType()
	if tq_UCmMaf0QU == nil {
		 /*line :1*/ panic("reflect: NumField of non-struct type")
	}
	return  /*line :1*/ len(tq_UCmMaf0QU.Fields)
}

func (gc1v5B2w4C eHKSMvk3XaJH) NumIn() int {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.FuncType()
	if tq_UCmMaf0QU == nil {
		 /*line :1*/ panic("reflect: NumIn of non-func type")
	}
	return  /*line :1*/ int(tq_UCmMaf0QU.InCount)
}

func (gc1v5B2w4C eHKSMvk3XaJH) NumOut() int {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.FuncType()
	if tq_UCmMaf0QU == nil {
		 /*line :1*/ panic("reflect: NumOut of non-func type")
	}
	return  /*line :1*/ tq_UCmMaf0QU.NumOut()
}

func (gc1v5B2w4C eHKSMvk3XaJH) Out(dujSIeNg10V int) LrcFby {
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.Type.FuncType()
	if tq_UCmMaf0QU == nil {
		 /*line :1*/ panic("reflect: Out of non-func type")
	}
	return  /*line :1*/ juqSJNrj( /*line :1*/ tq_UCmMaf0QU.OutSlice()[dujSIeNg10V])
}








func pOBZZZ(rE5JcBz3l unsafe.Pointer, jJCfpek uintptr, jo0rnS string) unsafe.Pointer {
	return  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(rE5JcBz3l) + jJCfpek)
}



func C08auGx(dujSIeNg10V any) LrcFby {
	mhQwtWWVa := *(* /*line :1*/ pgNclazaiBJc)( /*line :1*/ unsafe.Pointer(&dujSIeNg10V))
	return  /*line :1*/ juqSJNrj(mhQwtWWVa.ubPJaJymi)
}

func (gc1v5B2w4C eHKSMvk3XaJH) Implements(tvuza9gJ0 LrcFby) bool {
	if tvuza9gJ0 == nil {
		 /*line :1*/ panic("reflect: nil type passed to Type.Implements")
	}
	if  /*line :1*/ tvuza9gJ0.Kind() != Interface {
		 /*line :1*/ panic("reflect: non-interface type passed to Type.Implements")
	}
	return  /*line :1*/ t5oBGiG1gBz_( /*line :1*/ tvuza9gJ0.gebpy6Q4mW(),  /*line :1*/ gc1v5B2w4C.gebpy6Q4mW())
}

func (gc1v5B2w4C eHKSMvk3XaJH) AssignableTo(tvuza9gJ0 LrcFby) bool {
	if tvuza9gJ0 == nil {
		 /*line :1*/ panic("reflect: nil type passed to Type.AssignableTo")
	}
	db60unQQpO :=  /*line :1*/ tvuza9gJ0.gebpy6Q4mW()
	tq_UCmMaf0QU :=  /*line :1*/ gc1v5B2w4C.gebpy6Q4mW()
	return  /*line :1*/ aUlsvhDR(db60unQQpO, tq_UCmMaf0QU) ||  /*line :1*/ t5oBGiG1gBz_(db60unQQpO, tq_UCmMaf0QU)
}

func (gc1v5B2w4C eHKSMvk3XaJH) Comparable() bool {
	return gc1v5B2w4C.Equal != nil
}


func t5oBGiG1gBz_(A1q3iFIrWcfV, B_IOVmHQD *abi.Type) bool {
	gc1v5B2w4C :=  /*line :1*/ A1q3iFIrWcfV.InterfaceType()
	if gc1v5B2w4C == nil {
		return false
	}
	if  /*line :1*/ len(gc1v5B2w4C.Methods) == 0 {
		return true
	}
	cy1Lyoa9a :=  /*line :1*/ qK4nvEMY(A1q3iFIrWcfV)
	s_Lzlr :=  /*line :1*/ qK4nvEMY(B_IOVmHQD)

	if  /*line :1*/ B_IOVmHQD.Kind() == Interface {
		d583PKWfr := (* /*line :1*/ fh7O1l0V_Rdz)( /*line :1*/ unsafe.Pointer(B_IOVmHQD))
		dujSIeNg10V := 0
		for nkLpGFtYnZaZ := 0; nkLpGFtYnZaZ <  /*line :1*/ len(d583PKWfr.Methods); nkLpGFtYnZaZ++ {
			yhrWhzjq6fs := &gc1v5B2w4C.Methods[dujSIeNg10V]
			efAhPexFf :=  /*line :1*/ cy1Lyoa9a.sUmjH5(yhrWhzjq6fs.Name)
			wx_cBPo4P := &d583PKWfr.Methods[nkLpGFtYnZaZ]
			bdwzc8x :=  /*line :1*/ s_Lzlr.sUmjH5(wx_cBPo4P.Name)
			if  /*line :1*/ bdwzc8x.Name() ==  /*line :1*/ efAhPexFf.Name() &&  /*line :1*/ s_Lzlr.cnO61WNA_NK(wx_cBPo4P.Typ) ==  /*line :1*/ cy1Lyoa9a.cnO61WNA_NK(yhrWhzjq6fs.Typ) {
				if ! /*line :1*/ efAhPexFf.IsExported() {
					c0GVOx_v :=  /*line :1*/ ggkGz79(efAhPexFf)
					if c0GVOx_v == "" {
						c0GVOx_v =  /*line :1*/ gc1v5B2w4C.PkgPath.Name()
					}
					a6QflQ :=  /*line :1*/ ggkGz79(bdwzc8x)
					if a6QflQ == "" {
						a6QflQ =  /*line :1*/ d583PKWfr.PkgPath.Name()
					}
					if c0GVOx_v != a6QflQ {
						continue
					}
				}
				if dujSIeNg10V++; dujSIeNg10V >=  /*line :1*/ len(gc1v5B2w4C.Methods) {
					return true
				}
			}
		}
		return false
	}

	d583PKWfr :=  /*line :1*/ B_IOVmHQD.Uncommon()
	if d583PKWfr == nil {
		return false
	}
	dujSIeNg10V := 0
	ycL4Y_ :=  /*line :1*/ d583PKWfr.Methods()
	for nkLpGFtYnZaZ := 0; nkLpGFtYnZaZ <  /*line :1*/ int(d583PKWfr.Mcount); nkLpGFtYnZaZ++ {
		yhrWhzjq6fs := &gc1v5B2w4C.Methods[dujSIeNg10V]
		efAhPexFf :=  /*line :1*/ cy1Lyoa9a.sUmjH5(yhrWhzjq6fs.Name)
		wx_cBPo4P := ycL4Y_[nkLpGFtYnZaZ]
		bdwzc8x :=  /*line :1*/ s_Lzlr.sUmjH5(wx_cBPo4P.Name)
		if  /*line :1*/ bdwzc8x.Name() ==  /*line :1*/ efAhPexFf.Name() &&  /*line :1*/ s_Lzlr.cnO61WNA_NK(wx_cBPo4P.Mtyp) ==  /*line :1*/ cy1Lyoa9a.cnO61WNA_NK(yhrWhzjq6fs.Typ) {
			if ! /*line :1*/ efAhPexFf.IsExported() {
				c0GVOx_v :=  /*line :1*/ ggkGz79(efAhPexFf)
				if c0GVOx_v == "" {
					c0GVOx_v =  /*line :1*/ gc1v5B2w4C.PkgPath.Name()
				}
				a6QflQ :=  /*line :1*/ ggkGz79(bdwzc8x)
				if a6QflQ == "" {
					a6QflQ =  /*line :1*/ s_Lzlr.sUmjH5(d583PKWfr.PkgPath).Name()
				}
				if c0GVOx_v != a6QflQ {
					continue
				}
			}
			if dujSIeNg10V++; dujSIeNg10V >=  /*line :1*/ len(gc1v5B2w4C.Methods) {
				return true
			}
		}
	}
	return false
}






func aUlsvhDR(A1q3iFIrWcfV, B_IOVmHQD *abi.Type) bool {

	if A1q3iFIrWcfV == B_IOVmHQD {
		return true
	}

	if  /*line :1*/ A1q3iFIrWcfV.HasName() &&  /*line :1*/ B_IOVmHQD.HasName() ||  /*line :1*/ A1q3iFIrWcfV.Kind() !=  /*line :1*/ B_IOVmHQD.Kind() {
		return false
	}

	return  /*line :1*/ _VXkqUYqn(A1q3iFIrWcfV, B_IOVmHQD, true)
}

func upkhuIJrxZ(A1q3iFIrWcfV, B_IOVmHQD *abi.Type, yMi_c7Xu bool) bool {
	if yMi_c7Xu {
		return A1q3iFIrWcfV == B_IOVmHQD
	}

	if  /*line :1*/ qK4nvEMY(A1q3iFIrWcfV).Name() !=  /*line :1*/ qK4nvEMY(B_IOVmHQD).Name() ||  /*line :1*/ A1q3iFIrWcfV.Kind() !=  /*line :1*/ B_IOVmHQD.Kind() {
		return false
	}

	return  /*line :1*/ _VXkqUYqn(A1q3iFIrWcfV, B_IOVmHQD, false)
}

func _VXkqUYqn(A1q3iFIrWcfV, B_IOVmHQD *abi.Type, yMi_c7Xu bool) bool {
	if A1q3iFIrWcfV == B_IOVmHQD {
		return true
	}

	xm3h_2oMKuTJ :=  /*line :1*/ A1q3iFIrWcfV.Kind()
	if xm3h_2oMKuTJ !=  /*line :1*/ B_IOVmHQD.Kind() {
		return false
	}

	if abi.Bool <= xm3h_2oMKuTJ && xm3h_2oMKuTJ <= abi.Complex128 || xm3h_2oMKuTJ == abi.String || xm3h_2oMKuTJ == abi.UnsafePointer {
		return true
	}

	switch xm3h_2oMKuTJ {
	case abi.Array:
		return  /*line :1*/ A1q3iFIrWcfV.Len() ==  /*line :1*/ B_IOVmHQD.Len() &&  /*line :1*/ upkhuIJrxZ( /*line :1*/ A1q3iFIrWcfV.Elem(),  /*line :1*/ B_IOVmHQD.Elem(), yMi_c7Xu)

	case abi.Chan:

		if  /*line :1*/ B_IOVmHQD.ChanDir() == abi.BothDir &&  /*line :1*/ upkhuIJrxZ( /*line :1*/ A1q3iFIrWcfV.Elem(),  /*line :1*/ B_IOVmHQD.Elem(), yMi_c7Xu) {
			return true
		}

		return  /*line :1*/ B_IOVmHQD.ChanDir() ==  /*line :1*/ A1q3iFIrWcfV.ChanDir() &&  /*line :1*/ upkhuIJrxZ( /*line :1*/ A1q3iFIrWcfV.Elem(),  /*line :1*/ B_IOVmHQD.Elem(), yMi_c7Xu)

	case abi.Func:
		gc1v5B2w4C := (* /*line :1*/ jVlfi1v)( /*line :1*/ unsafe.Pointer(A1q3iFIrWcfV))
		d583PKWfr := (* /*line :1*/ jVlfi1v)( /*line :1*/ unsafe.Pointer(B_IOVmHQD))
		if gc1v5B2w4C.OutCount != d583PKWfr.OutCount || gc1v5B2w4C.InCount != d583PKWfr.InCount {
			return false
		}
		for dujSIeNg10V := 0; dujSIeNg10V <  /*line :1*/ gc1v5B2w4C.NumIn(); dujSIeNg10V++ {
			if ! /*line :1*/ upkhuIJrxZ( /*line :1*/ gc1v5B2w4C.In(dujSIeNg10V),  /*line :1*/ d583PKWfr.In(dujSIeNg10V), yMi_c7Xu) {
				return false
			}
		}
		for dujSIeNg10V := 0; dujSIeNg10V <  /*line :1*/ gc1v5B2w4C.NumOut(); dujSIeNg10V++ {
			if ! /*line :1*/ upkhuIJrxZ( /*line :1*/ gc1v5B2w4C.Out(dujSIeNg10V),  /*line :1*/ d583PKWfr.Out(dujSIeNg10V), yMi_c7Xu) {
				return false
			}
		}
		return true

	case Interface:
		gc1v5B2w4C := (* /*line :1*/ fh7O1l0V_Rdz)( /*line :1*/ unsafe.Pointer(A1q3iFIrWcfV))
		d583PKWfr := (* /*line :1*/ fh7O1l0V_Rdz)( /*line :1*/ unsafe.Pointer(B_IOVmHQD))
		if  /*line :1*/ len(gc1v5B2w4C.Methods) == 0 &&  /*line :1*/ len(d583PKWfr.Methods) == 0 {
			return true
		}

		return false

	case abi.Map:
		return  /*line :1*/ upkhuIJrxZ( /*line :1*/ A1q3iFIrWcfV.Key(),  /*line :1*/ B_IOVmHQD.Key(), yMi_c7Xu) &&  /*line :1*/ upkhuIJrxZ( /*line :1*/ A1q3iFIrWcfV.Elem(),  /*line :1*/ B_IOVmHQD.Elem(), yMi_c7Xu)

	case Ptr, abi.Slice:
		return  /*line :1*/ upkhuIJrxZ( /*line :1*/ A1q3iFIrWcfV.Elem(),  /*line :1*/ B_IOVmHQD.Elem(), yMi_c7Xu)

	case abi.Struct:
		gc1v5B2w4C := (* /*line :1*/ iIRpWLCsj)( /*line :1*/ unsafe.Pointer(A1q3iFIrWcfV))
		d583PKWfr := (* /*line :1*/ iIRpWLCsj)( /*line :1*/ unsafe.Pointer(B_IOVmHQD))
		if  /*line :1*/ len(gc1v5B2w4C.Fields) !=  /*line :1*/ len(d583PKWfr.Fields) {
			return false
		}
		if  /*line :1*/ gc1v5B2w4C.PkgPath.Name() !=  /*line :1*/ d583PKWfr.PkgPath.Name() {
			return false
		}
		for dujSIeNg10V := range gc1v5B2w4C.Fields {
			eaT0Rpa := &gc1v5B2w4C.Fields[dujSIeNg10V]
			djY9orHhAmKr := &d583PKWfr.Fields[dujSIeNg10V]
			if  /*line :1*/ eaT0Rpa.Name.Name() !=  /*line :1*/ djY9orHhAmKr.Name.Name() {
				return false
			}
			if ! /*line :1*/ upkhuIJrxZ(eaT0Rpa.Typ, djY9orHhAmKr.Typ, yMi_c7Xu) {
				return false
			}
			if yMi_c7Xu &&  /*line :1*/ eaT0Rpa.Name.Tag() !=  /*line :1*/ djY9orHhAmKr.Name.Tag() {
				return false
			}
			if eaT0Rpa.Offset != djY9orHhAmKr.Offset {
				return false
			}
			if  /*line :1*/ eaT0Rpa.Embedded() !=  /*line :1*/ djY9orHhAmKr.Embedded() {
				return false
			}
		}
		return true
	}

	return false
}






func juqSJNrj(gc1v5B2w4C *abi.Type) LrcFby {
	if gc1v5B2w4C == nil {
		return nil
	}
	return  /*line :1*/ qK4nvEMY(gc1v5B2w4C)
}


func vJEQ9cGoTSlB(gc1v5B2w4C *abi.Type) bool {
	return gc1v5B2w4C.Kind_&abi.KindDirectIface == 0
}
