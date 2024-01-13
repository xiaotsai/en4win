//line :1
package reflect

import (
	"internal/abi"
	"unsafe"
)

type eKObdgwq3 struct {
	j5qnH4NXz
	qwuYVEhPFQ	*bY2Zo8pis
	obIUuea1H528	func([]Value) []Value
}

func EiBL4fuxoss(lm5DH3 YJh989LTZsVX, lb3zWvJNFBQ func(fqjeuS []Value) (gy_vOn []Value)) Value {
	if  /*line :1*/ lm5DH3.Kind() != Func {
		 /*line :1*/ panic("reflect: call of MakeFunc with non-Func type")
	}

	sw8_lJ :=  /*line :1*/ lm5DH3.z6hGxTABM1()
	nbDrlvJDJ := (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer(sw8_lJ))

	jaLOurn4Fy :=  /*line :1*/ abi.FuncPCABI0(mvQHvM)

	_, _, dt53aB1lcTF :=  /*line :1*/ thrShO4Mk(nbDrlvJDJ, nil)

	gPBkB3yRaUy := &eKObdgwq3{
		j5qnH4NXz: j5qnH4NXz{
			gwQNUjAUPp:	jaLOurn4Fy,
			rJiml_:	dt53aB1lcTF.sCYaW3ZJnC,
			qW13qWfyFCF:	dt53aB1lcTF.yUiJTje95S,
			dKBQhiIv:	dt53aB1lcTF.xm5UM36Yljc,
		},
		qwuYVEhPFQ:	nbDrlvJDJ,
		obIUuea1H528:	lb3zWvJNFBQ,
	}

	return Value{sw8_lJ,  /*line :1*/ unsafe.Pointer(gPBkB3yRaUy),  /*line :1*/ flag(Func)}
}

func mvQHvM()

type jk52Gzqosa struct {
	j5qnH4NXz
	erQMhkxBxXr	int
	pYiECcnq0WnG	Value
}

func evDSV6R(kZVCtGQ string, f8NmcatRx Value) Value {
	if f8NmcatRx.flag&flagMethod == 0 {
		 /*line :1*/ panic("reflect: internal error: invalid use of makeMethodValue")
	}

	iiZKy9PE3 := f8NmcatRx.flag & (flagRO | flagAddr | flagIndir)
	iiZKy9PE3 |=  /*line :1*/ flag( /*line :1*/ f8NmcatRx.lm5DH3().Kind())
	iGkAMNgmOX := Value{ /*line :1*/ f8NmcatRx.lm5DH3(), f8NmcatRx.ptr, iiZKy9PE3}

	nbDrlvJDJ := (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.Type().(*rtype)))

	jaLOurn4Fy :=  /*line :1*/ vTsZZIIlq4MX()

	_, _, dt53aB1lcTF :=  /*line :1*/ thrShO4Mk(nbDrlvJDJ, nil)
	dhafmXfJi := &jk52Gzqosa{
		j5qnH4NXz: j5qnH4NXz{
			gwQNUjAUPp:	jaLOurn4Fy,
			rJiml_:	dt53aB1lcTF.sCYaW3ZJnC,
			qW13qWfyFCF:	dt53aB1lcTF.yUiJTje95S,
			dKBQhiIv:	dt53aB1lcTF.xm5UM36Yljc,
		},
		erQMhkxBxXr:	 /*line :1*/ int(f8NmcatRx.flag) >> flagMethodShift,
		pYiECcnq0WnG:	iGkAMNgmOX,
	}

	 /*line :1*/ eyL2pz(kZVCtGQ, dhafmXfJi.pYiECcnq0WnG, dhafmXfJi.erQMhkxBxXr)

	return Value{ /*line :1*/ nbDrlvJDJ.Common(),  /*line :1*/ unsafe.Pointer(dhafmXfJi), f8NmcatRx.flag&flagRO |  /*line :1*/ flag(Func)}
}

func vTsZZIIlq4MX() uintptr {
	return  /*line :1*/ abi.FuncPCABI0(rcXsXTt)
}

func rcXsXTt()

type j5qnH4NXz struct {
	gwQNUjAUPp	uintptr
	rJiml_	*uh88jBcDEay2
	qW13qWfyFCF	uintptr
	dKBQhiIv	abi.IntArgRegBitmap
}

//go:nosplit
func jLSRk5(m_WiSGQZjkA *j5qnH4NXz, fqjeuS *abi.RegArgs) {
	for hXZpj1nTZ, cUbsbg9j := range fqjeuS.Ints {

		if  /*line :1*/ m_WiSGQZjkA.dKBQhiIv.Get(hXZpj1nTZ) {
			*(* /*line :1*/ uintptr)( /*line :1*/ unsafe.Pointer(&fqjeuS.Ptrs[hXZpj1nTZ])) = cUbsbg9j
		} else {

			*(* /*line :1*/ uintptr)( /*line :1*/ unsafe.Pointer(&fqjeuS.Ptrs[hXZpj1nTZ])) = 0
		}
	}
}
