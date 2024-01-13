//line :1
package reflect

import (
	errors "iAHoxjmM"
	"internal/abi"
	"internal/goarch"
	itoa "JkjxVS"
	unsafeheader "EHdBEpZ7y6"
	math "math"
	"runtime"
	"unsafe"
)

type Value struct {
	typ_	*abi.Type

	ptr	unsafe.Pointer

	flag
}

type flag uintptr

const (
	flagKindWidth		= 5
	flagKindMask	flag	= 1<<flagKindWidth - 1
	flagStickyRO	flag	= 1 << 5
	flagEmbedRO	flag	= 1 << 6
	flagIndir	flag	= 1 << 7
	flagAddr	flag	= 1 << 8
	flagMethod	flag	= 1 << 9
	flagMethodShift		= 10
	flagRO	flag	= flagStickyRO | flagEmbedRO
)

func (jToThzM flag) zRPbQda4dT() WzYBjsQL0 {
	return  /*line :1*/ WzYBjsQL0(jToThzM & flagKindMask)
}

func (jToThzM flag) bQ77Iz() flag {
	if jToThzM&flagRO != 0 {
		return flagStickyRO
	}
	return 0
}

func (f8NmcatRx Value) lm5DH3() *abi.Type {

	return (* /*line :1*/ abi.Type)( /*line :1*/ m6KMjaQj( /*line :1*/ unsafe.Pointer(f8NmcatRx.typ_)))
}

func (f8NmcatRx Value) t_VyPi() unsafe.Pointer {
	if  /*line :1*/ f8NmcatRx.lm5DH3().Size() != goarch.PtrSize || ! /*line :1*/ f8NmcatRx.lm5DH3().Pointers() {
		 /*line :1*/ panic("can't call pointer on a non-pointer Value")
	}
	if f8NmcatRx.flag&flagIndir != 0 {
		return *(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr)
	}
	return f8NmcatRx.ptr
}

func hQWV0Z(f8NmcatRx Value) any {
	sw8_lJ :=  /*line :1*/ f8NmcatRx.lm5DH3()
	var hXZpj1nTZ any
	dVDneyuGq := (* /*line :1*/ cKqj6F0L)( /*line :1*/ unsafe.Pointer(&hXZpj1nTZ))

	switch {
	case  /*line :1*/ sw8_lJ.IfaceIndir():
		if f8NmcatRx.flag&flagIndir == 0 {
			 /*line :1*/ panic("bad indir")
		}

		cf3fCV8ayFq := f8NmcatRx.ptr
		if f8NmcatRx.flag&flagAddr != 0 {

			i5hAnuqALA :=  /*line :1*/ qscv_5(sw8_lJ)
			 /*line :1*/ l26oiKGd(sw8_lJ, i5hAnuqALA, cf3fCV8ayFq)
			cf3fCV8ayFq = i5hAnuqALA
		}
		dVDneyuGq.kMGZ7zYC = cf3fCV8ayFq
	case f8NmcatRx.flag&flagIndir != 0:

		dVDneyuGq.kMGZ7zYC = *(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr)
	default:

		dVDneyuGq.kMGZ7zYC = f8NmcatRx.ptr
	}

	dVDneyuGq.ubPJaJymi = sw8_lJ
	return hXZpj1nTZ
}

func i7YECW6sP91m(hXZpj1nTZ any) Value {
	dVDneyuGq := (* /*line :1*/ cKqj6F0L)( /*line :1*/ unsafe.Pointer(&hXZpj1nTZ))

	sw8_lJ := dVDneyuGq.ubPJaJymi
	if sw8_lJ == nil {
		return Value{}
	}
	jToThzM :=  /*line :1*/ flag( /*line :1*/ sw8_lJ.Kind())
	if  /*line :1*/ sw8_lJ.IfaceIndir() {
		jToThzM |= flagIndir
	}
	return Value{sw8_lJ, dVDneyuGq.kMGZ7zYC, jToThzM}
}

type R_QNJpzecp struct {
	Method	string
	N76ipnT	WzYBjsQL0
}

func (dVDneyuGq *R_QNJpzecp) Error() string {
	if dVDneyuGq.N76ipnT == 0 {
		return "reflect: call of " + dVDneyuGq.Method + " on zero Value"
	}
	return "reflect: call of " + dVDneyuGq.Method + " on " +  /*line :1*/ dVDneyuGq.N76ipnT.String() + " Value"
}

func pEyn8EHOAYa() string {
	var gXCmQp9Bo [5]uintptr
	krtR1HwEan :=  /*line :1*/ runtime.Callers(1, gXCmQp9Bo[:])
	dozaD3UK0Cpv :=  /*line :1*/ runtime.CallersFrames(gXCmQp9Bo[:krtR1HwEan])
	var nnl50dE0 runtime.Frame
	for arPLzdhyLR := true; arPLzdhyLR; {
		const prefix = "reflect.Value."
		nnl50dE0, arPLzdhyLR =  /*line :1*/ dozaD3UK0Cpv.Next()
		lGwAXqN4Hb := nnl50dE0.Function
		if  /*line :1*/ len(lGwAXqN4Hb) >  /*line :1*/ len(prefix) && lGwAXqN4Hb[: /*line :1*/ len(prefix)] == prefix {
			vdIVAGarq0 := lGwAXqN4Hb[ /*line :1*/ len(prefix):]
			if  /*line :1*/ len(vdIVAGarq0) > 0 && 'A' <= vdIVAGarq0[0] && vdIVAGarq0[0] <= 'Z' {
				return lGwAXqN4Hb
			}
		}
	}
	return "unknown method"
}

type cKqj6F0L struct {
	ubPJaJymi	*abi.Type
	kMGZ7zYC	unsafe.Pointer
}

type _aMMuPdMX struct {
	iyn2iDLPLe	*struct {
		lSdawoIpLh	*abi.Type
		zayJuicb	*abi.Type
		qn2mLTDFO	uint32
		_	[4]byte
		q0aSwHRR0A	[100000]unsafe.Pointer
	}
	sd1e98	unsafe.Pointer
}

func (jToThzM flag) zc9yxv(hd4WlqIU WzYBjsQL0) {

	if  /*line :1*/ WzYBjsQL0(jToThzM&flagKindMask) != hd4WlqIU {
		 /*line :1*/ panic(&R_QNJpzecp{ /*line :1*/ pEyn8EHOAYa(),  /*line :1*/ jToThzM.zRPbQda4dT()})
	}
}

func (jToThzM flag) rwcaztMf() {
	if jToThzM == 0 || jToThzM&flagRO != 0 {
		 /*line :1*/ jToThzM.tFA_ga()
	}
}

func (jToThzM flag) tFA_ga() {
	if jToThzM == 0 {
		 /*line :1*/ panic(&R_QNJpzecp{ /*line :1*/ pEyn8EHOAYa(), Invalid})
	}
	if jToThzM&flagRO != 0 {
		 /*line :1*/ panic("reflect: " +  /*line :1*/ pEyn8EHOAYa() + " using value obtained using unexported field")
	}
}

func (jToThzM flag) d5PnN9n() {
	if jToThzM&flagRO != 0 || jToThzM&flagAddr == 0 {
		 /*line :1*/ jToThzM.qR0DLrko()
	}
}

func (jToThzM flag) qR0DLrko() {
	if jToThzM == 0 {
		 /*line :1*/ panic(&R_QNJpzecp{ /*line :1*/ pEyn8EHOAYa(), Invalid})
	}

	if jToThzM&flagRO != 0 {
		 /*line :1*/ panic("reflect: " +  /*line :1*/ pEyn8EHOAYa() + " using value obtained using unexported field")
	}
	if jToThzM&flagAddr == 0 {
		 /*line :1*/ panic("reflect: " +  /*line :1*/ pEyn8EHOAYa() + " using unaddressable value")
	}
}

func (f8NmcatRx Value) Addr() Value {
	if f8NmcatRx.flag&flagAddr == 0 {
		 /*line :1*/ panic("reflect.Value.Addr of unaddressable value")
	}

	iiZKy9PE3 := f8NmcatRx.flag & flagRO
	return Value{ /*line :1*/ v8dZcSqC( /*line :1*/ f8NmcatRx.lm5DH3()), f8NmcatRx.ptr, iiZKy9PE3 |  /*line :1*/ flag(Pointer)}
}

func (f8NmcatRx Value) Bool() bool {

	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() != Bool {
		 /*line :1*/ f8NmcatRx.uRDwVFYjPx()
	}
	return *(* /*line :1*/ bool)(f8NmcatRx.ptr)
}

func (f8NmcatRx Value) uRDwVFYjPx() {
	 /*line :1*/ f8NmcatRx.zc9yxv(Bool)
}

var gscejP =  /*line :1*/ gC11Vi(([] /*line :1*/ byte)(nil))

func (f8NmcatRx Value) Bytes() []byte {

	if f8NmcatRx.typ_ == gscejP {
		return *(*[] /*line :1*/ byte)(f8NmcatRx.ptr)
	}
	return  /*line :1*/ f8NmcatRx.tjKs64_u()
}

func (f8NmcatRx Value) tjKs64_u() []byte {
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Slice:
		if  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() != abi.Uint8 {
			 /*line :1*/ panic("reflect.Value.Bytes of non-byte slice")
		}

		return *(*[] /*line :1*/ byte)(f8NmcatRx.ptr)
	case Array:
		if  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() != abi.Uint8 {
			 /*line :1*/ panic("reflect.Value.Bytes of non-byte array")
		}
		if ! /*line :1*/ f8NmcatRx.CanAddr() {
			 /*line :1*/ panic("reflect.Value.Bytes of unaddressable byte array")
		}
		e3IyyaQSEj := (* /*line :1*/ byte)(f8NmcatRx.ptr)
		krtR1HwEan :=  /*line :1*/ int((* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3())).Len)
		return  /*line :1*/ unsafe.Slice(e3IyyaQSEj, krtR1HwEan)
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Bytes",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) fjMiydk() []rune {
	 /*line :1*/ f8NmcatRx.zc9yxv(Slice)
	if  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() != abi.Int32 {
		 /*line :1*/ panic("reflect.Value.Bytes of non-rune slice")
	}

	return *(*[] /*line :1*/ rune)(f8NmcatRx.ptr)
}

func (f8NmcatRx Value) CanAddr() bool {
	return f8NmcatRx.flag&flagAddr != 0
}

func (f8NmcatRx Value) CanSet() bool {
	return f8NmcatRx.flag&(flagAddr|flagRO) == flagAddr
}

func (f8NmcatRx Value) Call(gJ3XGL_F []Value) []Value {
	 /*line :1*/ f8NmcatRx.zc9yxv(Func)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	return  /*line :1*/ f8NmcatRx.d7AWjWcFH("Call", gJ3XGL_F)
}

func (f8NmcatRx Value) CallSlice(gJ3XGL_F []Value) []Value {
	 /*line :1*/ f8NmcatRx.zc9yxv(Func)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	return  /*line :1*/ f8NmcatRx.d7AWjWcFH("CallSlice", gJ3XGL_F)
}

var c79cDyCri2 bool

const debugReflectCall = false

func (f8NmcatRx Value) d7AWjWcFH(kZVCtGQ string, gJ3XGL_F []Value) []Value {

	sw8_lJ := (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
	var (
		lb3zWvJNFBQ	unsafe.Pointer
		iGkAMNgmOX	Value
		kjn54azLiX	*abi.Type
	)
	if f8NmcatRx.flag&flagMethod != 0 {
		iGkAMNgmOX = f8NmcatRx
		kjn54azLiX, sw8_lJ, lb3zWvJNFBQ =  /*line :1*/ eyL2pz(kZVCtGQ, f8NmcatRx,  /*line :1*/ int(f8NmcatRx.flag)>>flagMethodShift)
	} else if f8NmcatRx.flag&flagIndir != 0 {
		lb3zWvJNFBQ = *(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr)
	} else {
		lb3zWvJNFBQ = f8NmcatRx.ptr
	}

	if lb3zWvJNFBQ == nil {
		 /*line :1*/ panic("reflect.Value.Call: call of nil function")
	}

	hKSkLfqDR := kZVCtGQ == "CallSlice"
	krtR1HwEan :=  /*line :1*/ sw8_lJ.NumIn()
	tEMEZG :=  /*line :1*/ sw8_lJ.IsVariadic()
	if hKSkLfqDR {
		if !tEMEZG {
			 /*line :1*/ panic("reflect: CallSlice of non-variadic function")
		}
		if  /*line :1*/ len(gJ3XGL_F) < krtR1HwEan {
			 /*line :1*/ panic("reflect: CallSlice with too few input arguments")
		}
		if  /*line :1*/ len(gJ3XGL_F) > krtR1HwEan {
			 /*line :1*/ panic("reflect: CallSlice with too many input arguments")
		}
	} else {
		if tEMEZG {
			krtR1HwEan--
		}
		if  /*line :1*/ len(gJ3XGL_F) < krtR1HwEan {
			 /*line :1*/ panic("reflect: Call with too few input arguments")
		}
		if !tEMEZG &&  /*line :1*/ len(gJ3XGL_F) > krtR1HwEan {
			 /*line :1*/ panic("reflect: Call with too many input arguments")
		}
	}
	for _, uiTlN8bdKm := range gJ3XGL_F {
		if  /*line :1*/ uiTlN8bdKm.Kind() == Invalid {
			 /*line :1*/ panic("reflect: " + kZVCtGQ + " using zero Value argument")
		}
	}
	for hXZpj1nTZ := 0; hXZpj1nTZ < krtR1HwEan; hXZpj1nTZ++ {
		if z5c5ft, y2QgMxcpiCbQ :=  /*line :1*/ gJ3XGL_F[hXZpj1nTZ].Type(),  /*line :1*/ sw8_lJ.In(hXZpj1nTZ); ! /*line :1*/ z5c5ft.AssignableTo( /*line :1*/ j5Z_WRWLudek(y2QgMxcpiCbQ)) {
			 /*line :1*/ panic("reflect: " + kZVCtGQ + " using " +  /*line :1*/ z5c5ft.String() + " as type " +  /*line :1*/ iR_1JLd(y2QgMxcpiCbQ))
		}
	}
	if !hKSkLfqDR && tEMEZG {

		hek9DV :=  /*line :1*/ len(gJ3XGL_F) - krtR1HwEan
		qMTDS7fnbc :=  /*line :1*/ LydH53J9Si( /*line :1*/ j5Z_WRWLudek( /*line :1*/ sw8_lJ.In(krtR1HwEan)), hek9DV, hek9DV)
		dPe0EQn :=  /*line :1*/ j5Z_WRWLudek( /*line :1*/ sw8_lJ.In(krtR1HwEan)).Elem()
		for hXZpj1nTZ := 0; hXZpj1nTZ < hek9DV; hXZpj1nTZ++ {
			uiTlN8bdKm := gJ3XGL_F[krtR1HwEan+hXZpj1nTZ]
			if z5c5ft :=  /*line :1*/ uiTlN8bdKm.Type(); ! /*line :1*/ z5c5ft.AssignableTo(dPe0EQn) {
				 /*line :1*/ panic("reflect: cannot use " +  /*line :1*/ z5c5ft.String() + " as type " +  /*line :1*/ dPe0EQn.String() + " in " + kZVCtGQ)
			}
			 /*line :1*/ qMTDS7fnbc.Index(hXZpj1nTZ).Set(uiTlN8bdKm)
		}
		afPJKW := gJ3XGL_F
		gJ3XGL_F =  /*line :1*/ make([]Value, krtR1HwEan+1)
		 /*line :1*/ copy(gJ3XGL_F[:krtR1HwEan], afPJKW)
		gJ3XGL_F[krtR1HwEan] = qMTDS7fnbc
	}

	oFcK6X1M4 :=  /*line :1*/ len(gJ3XGL_F)
	if oFcK6X1M4 !=  /*line :1*/ sw8_lJ.NumIn() {
		 /*line :1*/ panic("reflect.Value.Call: wrong argument count")
	}
	eFYn5eNab :=  /*line :1*/ sw8_lJ.NumOut()

	var sfPBRit abi.RegArgs

	nuCjUv_xpB0p, a2eCYr, dt53aB1lcTF :=  /*line :1*/ thrShO4Mk(sw8_lJ, kjn54azLiX)

	var mkehBv9 unsafe.Pointer
	if  /*line :1*/ nuCjUv_xpB0p.Size() != 0 {
		if eFYn5eNab == 0 {
			mkehBv9 =  /*line :1*/ a2eCYr.Get().(unsafe.Pointer)
		} else {

			mkehBv9 =  /*line :1*/ qscv_5(nuCjUv_xpB0p)
		}
	}
	ayUQnG0gkSPC :=  /*line :1*/ nuCjUv_xpB0p.Size()

	if debugReflectCall {
		 /*line :1*/ println("reflect.call",  /*line :1*/ iR_1JLd(&sw8_lJ.Type))
		 /*line :1*/ dt53aB1lcTF.a6LuuxPayDWk()
	}

	zYksvDWt := 0
	if kjn54azLiX != nil {

		switch pjngE6n4a := dt53aB1lcTF.pO2s2NwMlS.bbEexQdWR[0]; pjngE6n4a._BtGRMo2 {
		case abiStepStack:
			 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX, mkehBv9)
		case abiStepPointer:
			 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX,  /*line :1*/ unsafe.Pointer(&sfPBRit.Ptrs[pjngE6n4a.pCQRHPfHe]))
			fallthrough
		case abiStepIntReg:
			 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX,  /*line :1*/ unsafe.Pointer(&sfPBRit.Ints[pjngE6n4a.pCQRHPfHe]))
		case abiStepFloatReg:
			 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX,  /*line :1*/ unsafe.Pointer(&sfPBRit.Floats[pjngE6n4a.kT87Jzba6Sl]))
		default:
			 /*line :1*/ panic("unknown ABI parameter kind")
		}
		zYksvDWt = 1
	}

	for hXZpj1nTZ, f8NmcatRx := range gJ3XGL_F {
		 /*line :1*/ f8NmcatRx.rwcaztMf()
		y2QgMxcpiCbQ :=  /*line :1*/ j5Z_WRWLudek( /*line :1*/ sw8_lJ.In(hXZpj1nTZ))

		f8NmcatRx =  /*line :1*/ f8NmcatRx.hfrHMyTs1aE("reflect.Value.Call", &y2QgMxcpiCbQ.t, nil)
	stepsLoop:
		for _, pjngE6n4a := range  /*line :1*/ dt53aB1lcTF.pO2s2NwMlS.xAplKg6Om(hXZpj1nTZ + zYksvDWt) {
			switch pjngE6n4a._BtGRMo2 {
			case abiStepStack:

				sitLxDQ9Ye :=  /*line :1*/ lxkP7oV7sm(mkehBv9, pjngE6n4a.ciJuo5C, "precomputed stack arg offset")
				if f8NmcatRx.flag&flagIndir != 0 {
					 /*line :1*/ l26oiKGd(&y2QgMxcpiCbQ.t, sitLxDQ9Ye, f8NmcatRx.ptr)
				} else {
					*(* /*line :1*/ unsafe.Pointer)(sitLxDQ9Ye) = f8NmcatRx.ptr
				}

				break stepsLoop
			case abiStepIntReg, abiStepPointer:

				if f8NmcatRx.flag&flagIndir != 0 {
					dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, pjngE6n4a.v9du88xTkB, "precomputed value offset")
					if pjngE6n4a._BtGRMo2 == abiStepPointer {

						sfPBRit.Ptrs[pjngE6n4a.pCQRHPfHe] = *(* /*line :1*/ unsafe.Pointer)(dcDpXllPZ3e)
					}
					 /*line :1*/ bGUOvK(&sfPBRit, pjngE6n4a.pCQRHPfHe, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
				} else {
					if pjngE6n4a._BtGRMo2 == abiStepPointer {

						sfPBRit.Ptrs[pjngE6n4a.pCQRHPfHe] = f8NmcatRx.ptr
					}
					sfPBRit.Ints[pjngE6n4a.pCQRHPfHe] =  /*line :1*/ uintptr(f8NmcatRx.ptr)
				}
			case abiStepFloatReg:

				if f8NmcatRx.flag&flagIndir == 0 {
					 /*line :1*/ panic("attempted to copy pointer to FP register")
				}
				dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, pjngE6n4a.v9du88xTkB, "precomputed value offset")
				 /*line :1*/ nhGam6sL_Y(&sfPBRit, pjngE6n4a.kT87Jzba6Sl, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
			default:
				 /*line :1*/ panic("unknown ABI part kind")
			}
		}
	}

	ayUQnG0gkSPC =  /*line :1*/ rlY5YEji1B4(ayUQnG0gkSPC, goarch.PtrSize)
	ayUQnG0gkSPC += dt53aB1lcTF.hq8R8Dk

	sfPBRit.ReturnIsPtr = dt53aB1lcTF.aHHpIUai1ke

	if debugReflectCall {
		 /*line :1*/ sfPBRit.Dump()
	}

	if c79cDyCri2 {
		 /*line :1*/ runtime.GC()
	}

	 /*line :1*/ d7AWjWcFH(nuCjUv_xpB0p, lb3zWvJNFBQ, mkehBv9,  /*line :1*/ uint32( /*line :1*/ nuCjUv_xpB0p.Size()),  /*line :1*/ uint32(dt53aB1lcTF.eBcFPCE),  /*line :1*/ uint32(ayUQnG0gkSPC), &sfPBRit)

	if c79cDyCri2 {
		 /*line :1*/ runtime.GC()
	}

	var aLsdcpyP1 []Value
	if eFYn5eNab == 0 {
		if mkehBv9 != nil {
			 /*line :1*/ q2gsmEF(nuCjUv_xpB0p, mkehBv9)
			 /*line :1*/ a2eCYr.Put(mkehBv9)
		}
	} else {
		if mkehBv9 != nil {

			 /*line :1*/ u5gG9lm_(nuCjUv_xpB0p, mkehBv9, 0, dt53aB1lcTF.eBcFPCE)
		}

		aLsdcpyP1 =  /*line :1*/ make([]Value, eFYn5eNab)
		for hXZpj1nTZ := 0; hXZpj1nTZ < eFYn5eNab; hXZpj1nTZ++ {
			sH_6eEqK_wj :=  /*line :1*/ sw8_lJ.Out(hXZpj1nTZ)
			if  /*line :1*/ sH_6eEqK_wj.Size() == 0 {

				aLsdcpyP1[hXZpj1nTZ] =  /*line :1*/ H_B9rU1ADy( /*line :1*/ j5Z_WRWLudek(sH_6eEqK_wj))
				continue
			}
			btXCoaRp0U0 :=  /*line :1*/ dt53aB1lcTF.gSH9DY2KdKea.xAplKg6Om(hXZpj1nTZ)
			if pjngE6n4a := btXCoaRp0U0[0]; pjngE6n4a._BtGRMo2 == abiStepStack {

				iiZKy9PE3 := flagIndir |  /*line :1*/ flag( /*line :1*/ sH_6eEqK_wj.Kind())
				aLsdcpyP1[hXZpj1nTZ] = Value{sH_6eEqK_wj,  /*line :1*/ lxkP7oV7sm(mkehBv9, pjngE6n4a.ciJuo5C, "tv.Size() != 0"), iiZKy9PE3}

				continue
			}

			if ! /*line :1*/ gGK1Oc(sH_6eEqK_wj) {

				if btXCoaRp0U0[0]._BtGRMo2 != abiStepPointer {
					 /*line :1*/ print("kind=", btXCoaRp0U0[0]._BtGRMo2, ", type=",  /*line :1*/ iR_1JLd(sH_6eEqK_wj), "\n")
					 /*line :1*/ panic("mismatch between ABI description and types")
				}
				aLsdcpyP1[hXZpj1nTZ] = Value{sH_6eEqK_wj, sfPBRit.Ptrs[btXCoaRp0U0[0].pCQRHPfHe],  /*line :1*/ flag( /*line :1*/ sH_6eEqK_wj.Kind())}
				continue
			}

			kXTvhUI4Zgg5 :=  /*line :1*/ qscv_5(sH_6eEqK_wj)
			for _, pjngE6n4a := range btXCoaRp0U0 {
				switch pjngE6n4a._BtGRMo2 {
				case abiStepIntReg:
					dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(kXTvhUI4Zgg5, pjngE6n4a.v9du88xTkB, "precomputed value offset")
					 /*line :1*/ hchALWv(&sfPBRit, pjngE6n4a.pCQRHPfHe, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
				case abiStepPointer:
					kXTvhUI4Zgg5 :=  /*line :1*/ lxkP7oV7sm(kXTvhUI4Zgg5, pjngE6n4a.v9du88xTkB, "precomputed value offset")
					*((* /*line :1*/ unsafe.Pointer)(kXTvhUI4Zgg5)) = sfPBRit.Ptrs[pjngE6n4a.pCQRHPfHe]
				case abiStepFloatReg:
					dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(kXTvhUI4Zgg5, pjngE6n4a.v9du88xTkB, "precomputed value offset")
					 /*line :1*/ kh2Bx1R9(&sfPBRit, pjngE6n4a.kT87Jzba6Sl, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
				case abiStepStack:
					 /*line :1*/ panic("register-based return value has stack component")
				default:
					 /*line :1*/ panic("unknown ABI part kind")
				}
			}
			aLsdcpyP1[hXZpj1nTZ] = Value{sH_6eEqK_wj, kXTvhUI4Zgg5, flagIndir |  /*line :1*/ flag( /*line :1*/ sH_6eEqK_wj.Kind())}
		}
	}

	return aLsdcpyP1
}

func iaLgzrMhw(m_WiSGQZjkA *eKObdgwq3, nnl50dE0 unsafe.Pointer, lATw9_ *bool, nwQHFtyw *abi.RegArgs) {
	if c79cDyCri2 {

		 /*line :1*/ runtime.GC()
	}
	nbDrlvJDJ := m_WiSGQZjkA.qwuYVEhPFQ
	jToThzM := m_WiSGQZjkA.obIUuea1H528

	_, _, dt53aB1lcTF :=  /*line :1*/ thrShO4Mk(nbDrlvJDJ, nil)

	cf3fCV8ayFq := nnl50dE0
	gJ3XGL_F :=  /*line :1*/ make([]Value, 0,  /*line :1*/ int(nbDrlvJDJ.InCount))
	for hXZpj1nTZ, lm5DH3 := range  /*line :1*/ nbDrlvJDJ.InSlice() {
		if  /*line :1*/ lm5DH3.Size() == 0 {
			gJ3XGL_F =  /*line :1*/ append(gJ3XGL_F,  /*line :1*/ H_B9rU1ADy( /*line :1*/ j5Z_WRWLudek(lm5DH3)))
			continue
		}
		f8NmcatRx := Value{lm5DH3, nil,  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())}
		btXCoaRp0U0 :=  /*line :1*/ dt53aB1lcTF.pO2s2NwMlS.xAplKg6Om(hXZpj1nTZ)
		if pjngE6n4a := btXCoaRp0U0[0]; pjngE6n4a._BtGRMo2 == abiStepStack {
			if  /*line :1*/ gGK1Oc(lm5DH3) {

				f8NmcatRx.ptr =  /*line :1*/ qscv_5(lm5DH3)
				if  /*line :1*/ lm5DH3.Size() > 0 {
					 /*line :1*/ l26oiKGd(lm5DH3, f8NmcatRx.ptr,  /*line :1*/ lxkP7oV7sm(cf3fCV8ayFq, pjngE6n4a.ciJuo5C, "typ.size > 0"))
				}
				f8NmcatRx.flag |= flagIndir
			} else {
				f8NmcatRx.ptr = *(* /*line :1*/ unsafe.Pointer)( /*line :1*/ lxkP7oV7sm(cf3fCV8ayFq, pjngE6n4a.ciJuo5C, "1-ptr"))
			}
		} else {
			if  /*line :1*/ gGK1Oc(lm5DH3) {

				f8NmcatRx.flag |= flagIndir
				f8NmcatRx.ptr =  /*line :1*/ qscv_5(lm5DH3)
				for _, pjngE6n4a := range btXCoaRp0U0 {
					switch pjngE6n4a._BtGRMo2 {
					case abiStepIntReg:
						dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, pjngE6n4a.v9du88xTkB, "precomputed value offset")
						 /*line :1*/ hchALWv(nwQHFtyw, pjngE6n4a.pCQRHPfHe, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
					case abiStepPointer:
						kXTvhUI4Zgg5 :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, pjngE6n4a.v9du88xTkB, "precomputed value offset")
						*((* /*line :1*/ unsafe.Pointer)(kXTvhUI4Zgg5)) = nwQHFtyw.Ptrs[pjngE6n4a.pCQRHPfHe]
					case abiStepFloatReg:
						dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, pjngE6n4a.v9du88xTkB, "precomputed value offset")
						 /*line :1*/ kh2Bx1R9(nwQHFtyw, pjngE6n4a.kT87Jzba6Sl, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
					case abiStepStack:
						 /*line :1*/ panic("register-based return value has stack component")
					default:
						 /*line :1*/ panic("unknown ABI part kind")
					}
				}
			} else {

				if btXCoaRp0U0[0]._BtGRMo2 != abiStepPointer {
					 /*line :1*/ print("kind=", btXCoaRp0U0[0]._BtGRMo2, ", type=",  /*line :1*/ iR_1JLd(lm5DH3), "\n")
					 /*line :1*/ panic("mismatch between ABI description and types")
				}
				f8NmcatRx.ptr = nwQHFtyw.Ptrs[btXCoaRp0U0[0].pCQRHPfHe]
			}
		}
		gJ3XGL_F =  /*line :1*/ append(gJ3XGL_F, f8NmcatRx)
	}

	hEbpexD :=  /*line :1*/ jToThzM(gJ3XGL_F)
	cUpiIX2x :=  /*line :1*/ nbDrlvJDJ.NumOut()
	if  /*line :1*/ len(hEbpexD) != cUpiIX2x {
		 /*line :1*/ panic("reflect: wrong return count from function created by MakeFunc")
	}

	if cUpiIX2x > 0 {
		for hXZpj1nTZ, lm5DH3 := range  /*line :1*/ nbDrlvJDJ.OutSlice() {
			f8NmcatRx := hEbpexD[hXZpj1nTZ]
			if  /*line :1*/ f8NmcatRx.lm5DH3() == nil {
				 /*line :1*/ panic("reflect: function created by MakeFunc using " +  /*line :1*/ wtBvIFaaYS(jToThzM) +
					" returned zero Value")
			}
			if f8NmcatRx.flag&flagRO != 0 {
				 /*line :1*/ panic("reflect: function created by MakeFunc using " +  /*line :1*/ wtBvIFaaYS(jToThzM) +
					" returned value obtained from unexported field")
			}
			if  /*line :1*/ lm5DH3.Size() == 0 {
				continue
			}

			f8NmcatRx =  /*line :1*/ f8NmcatRx.hfrHMyTs1aE("reflect.MakeFunc", lm5DH3, nil)
		stepsLoop:
			for _, pjngE6n4a := range  /*line :1*/ dt53aB1lcTF.gSH9DY2KdKea.xAplKg6Om(hXZpj1nTZ) {
				switch pjngE6n4a._BtGRMo2 {
				case abiStepStack:

					sitLxDQ9Ye :=  /*line :1*/ lxkP7oV7sm(cf3fCV8ayFq, pjngE6n4a.ciJuo5C, "precomputed stack arg offset")

					if f8NmcatRx.flag&flagIndir != 0 {
						 /*line :1*/ pqAEZr6n(sitLxDQ9Ye, f8NmcatRx.ptr, pjngE6n4a.jLZh5ayKa9mx)
					} else {

						*(* /*line :1*/ uintptr)(sitLxDQ9Ye) =  /*line :1*/ uintptr(f8NmcatRx.ptr)
					}

					break stepsLoop
				case abiStepIntReg, abiStepPointer:

					if f8NmcatRx.flag&flagIndir != 0 {
						dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, pjngE6n4a.v9du88xTkB, "precomputed value offset")
						 /*line :1*/ bGUOvK(nwQHFtyw, pjngE6n4a.pCQRHPfHe, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
					} else {

						nwQHFtyw.Ints[pjngE6n4a.pCQRHPfHe] =  /*line :1*/ uintptr(f8NmcatRx.ptr)
					}
				case abiStepFloatReg:

					if f8NmcatRx.flag&flagIndir == 0 {
						 /*line :1*/ panic("attempted to copy pointer to FP register")
					}
					dcDpXllPZ3e :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, pjngE6n4a.v9du88xTkB, "precomputed value offset")
					 /*line :1*/ nhGam6sL_Y(nwQHFtyw, pjngE6n4a.kT87Jzba6Sl, pjngE6n4a.jLZh5ayKa9mx, dcDpXllPZ3e)
				default:
					 /*line :1*/ panic("unknown ABI part kind")
				}
			}
		}
	}

	*lATw9_ = true

	 /*line :1*/ runtime.KeepAlive(hEbpexD)

	 /*line :1*/ runtime.KeepAlive(m_WiSGQZjkA)
}

func eyL2pz(kZVCtGQ string, f8NmcatRx Value, osP5xU int) (kjn54azLiX *abi.Type, sw8_lJ *bY2Zo8pis, lb3zWvJNFBQ unsafe.Pointer) {
	hXZpj1nTZ := osP5xU
	if  /*line :1*/ f8NmcatRx.lm5DH3().Kind() == abi.Interface {
		mP7QqSVInV := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint( /*line :1*/ len(mP7QqSVInV.Methods)) {
			 /*line :1*/ panic("reflect: internal error: invalid method index")
		}
		hek9DV := &mP7QqSVInV.Methods[hXZpj1nTZ]
		if ! /*line :1*/ mP7QqSVInV.zPya6B(hek9DV.Name).IsExported() {
			 /*line :1*/ panic("reflect: " + kZVCtGQ + " of unexported method")
		}
		_ROffP9x := (* /*line :1*/ _aMMuPdMX)(f8NmcatRx.ptr)
		if _ROffP9x.iyn2iDLPLe == nil {
			 /*line :1*/ panic("reflect: " + kZVCtGQ + " of method on nil interface value")
		}
		kjn54azLiX = _ROffP9x.iyn2iDLPLe.zayJuicb
		lb3zWvJNFBQ =  /*line :1*/ unsafe.Pointer(&_ROffP9x.iyn2iDLPLe.q0aSwHRR0A[hXZpj1nTZ])
		sw8_lJ = (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer( /*line :1*/ mP7QqSVInV.x9CINy8a(hek9DV.Typ)))
	} else {
		kjn54azLiX =  /*line :1*/ f8NmcatRx.lm5DH3()
		vlMabJFJks :=  /*line :1*/ f8NmcatRx.lm5DH3().ExportedMethods()
		if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint( /*line :1*/ len(vlMabJFJks)) {
			 /*line :1*/ panic("reflect: internal error: invalid method index")
		}
		hek9DV := vlMabJFJks[hXZpj1nTZ]
		if ! /*line :1*/ drswiOv( /*line :1*/ f8NmcatRx.lm5DH3(), hek9DV.Name).IsExported() {
			 /*line :1*/ panic("reflect: " + kZVCtGQ + " of unexported method")
		}
		pVePXfqREHol :=  /*line :1*/ zTRltR( /*line :1*/ f8NmcatRx.lm5DH3(), hek9DV.Ifn)
		lb3zWvJNFBQ =  /*line :1*/ unsafe.Pointer(&pVePXfqREHol)
		sw8_lJ = (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer( /*line :1*/ nsOKvPP( /*line :1*/ f8NmcatRx.lm5DH3(), hek9DV.Mtyp)))
	}
	return
}

func d3Ql5e5xiW(f8NmcatRx Value, e3IyyaQSEj unsafe.Pointer) {
	sw8_lJ :=  /*line :1*/ f8NmcatRx.lm5DH3()
	if  /*line :1*/ sw8_lJ.Kind() == abi.Interface {

		_ROffP9x := (* /*line :1*/ _aMMuPdMX)(f8NmcatRx.ptr)
		*(* /*line :1*/ unsafe.Pointer)(e3IyyaQSEj) = _ROffP9x.sd1e98
	} else if f8NmcatRx.flag&flagIndir != 0 && ! /*line :1*/ gGK1Oc(sw8_lJ) {
		*(* /*line :1*/ unsafe.Pointer)(e3IyyaQSEj) = *(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr)
	} else {
		*(* /*line :1*/ unsafe.Pointer)(e3IyyaQSEj) = f8NmcatRx.ptr
	}
}

func rlY5YEji1B4(uiTlN8bdKm, krtR1HwEan uintptr) uintptr {
	return (uiTlN8bdKm + krtR1HwEan - 1) &^ (krtR1HwEan - 1)
}

func o3qCFd0(m_WiSGQZjkA *jk52Gzqosa, nnl50dE0 unsafe.Pointer, lATw9_ *bool, nwQHFtyw *abi.RegArgs) {
	iGkAMNgmOX := m_WiSGQZjkA.pYiECcnq0WnG
	aDX_pcbm8VL, zR_zVa, eKIiSoi0ulS :=  /*line :1*/ eyL2pz("call", iGkAMNgmOX, m_WiSGQZjkA.erQMhkxBxXr)

	_, _, aI3R_aim5Ht :=  /*line :1*/ thrShO4Mk(zR_zVa, nil)
	o4PQFZaVsd, xDAZAoprC2MC := nnl50dE0, nwQHFtyw
	kasvEYpX3A, lpyVCliC5UbT, mtlKwzeEO4 :=  /*line :1*/ thrShO4Mk(zR_zVa, aDX_pcbm8VL)

	a2T1BCJOcmaY :=  /*line :1*/ lpyVCliC5UbT.Get().(unsafe.Pointer)
	var qtyCUf abi.RegArgs

	switch pjngE6n4a := mtlKwzeEO4.pO2s2NwMlS.bbEexQdWR[0]; pjngE6n4a._BtGRMo2 {
	case abiStepStack:

		 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX, a2T1BCJOcmaY)
	case abiStepPointer:

		 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX,  /*line :1*/ unsafe.Pointer(&qtyCUf.Ptrs[pjngE6n4a.pCQRHPfHe]))
		fallthrough
	case abiStepIntReg:
		 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX,  /*line :1*/ unsafe.Pointer(&qtyCUf.Ints[pjngE6n4a.pCQRHPfHe]))
	case abiStepFloatReg:
		 /*line :1*/ d3Ql5e5xiW(iGkAMNgmOX,  /*line :1*/ unsafe.Pointer(&qtyCUf.Floats[pjngE6n4a.kT87Jzba6Sl]))
	default:
		 /*line :1*/ panic("unknown ABI parameter kind")
	}

	for hXZpj1nTZ, sw8_lJ := range  /*line :1*/ zR_zVa.InSlice() {
		fWunZ5qrKW :=  /*line :1*/ aI3R_aim5Ht.pO2s2NwMlS.xAplKg6Om(hXZpj1nTZ)
		an0_Eja :=  /*line :1*/ mtlKwzeEO4.pO2s2NwMlS.xAplKg6Om(hXZpj1nTZ + 1)

		if  /*line :1*/ len(fWunZ5qrKW) == 0 {
			if  /*line :1*/ len(an0_Eja) != 0 {
				 /*line :1*/ panic("method ABI and value ABI do not align")
			}
			continue
		}

		if fNFezZHudDS := fWunZ5qrKW[0]; fNFezZHudDS._BtGRMo2 == abiStepStack {
			mp7hZaO := an0_Eja[0]

			if mp7hZaO._BtGRMo2 == abiStepStack {
				if fNFezZHudDS.jLZh5ayKa9mx != mp7hZaO.jLZh5ayKa9mx {
					 /*line :1*/ panic("method ABI and value ABI do not align")
				}
				 /*line :1*/ l26oiKGd(sw8_lJ,
					 /*line :1*/ lxkP7oV7sm(a2T1BCJOcmaY, mp7hZaO.ciJuo5C, "precomputed stack offset"),
					 /*line :1*/ lxkP7oV7sm(o4PQFZaVsd, fNFezZHudDS.ciJuo5C, "precomputed stack offset"))
				continue
			}

			for _, mp7hZaO := range an0_Eja {
				xFFtzJR :=  /*line :1*/ lxkP7oV7sm(o4PQFZaVsd, fNFezZHudDS.ciJuo5C+mp7hZaO.v9du88xTkB, "precomputed stack offset")
				switch mp7hZaO._BtGRMo2 {
				case abiStepPointer:

					qtyCUf.Ptrs[mp7hZaO.pCQRHPfHe] = *(* /*line :1*/ unsafe.Pointer)(xFFtzJR)
					fallthrough
				case abiStepIntReg:
					 /*line :1*/ bGUOvK(&qtyCUf, mp7hZaO.pCQRHPfHe, mp7hZaO.jLZh5ayKa9mx, xFFtzJR)
				case abiStepFloatReg:
					 /*line :1*/ nhGam6sL_Y(&qtyCUf, mp7hZaO.kT87Jzba6Sl, mp7hZaO.jLZh5ayKa9mx, xFFtzJR)
				default:
					 /*line :1*/ panic("unexpected method step")
				}
			}
			continue
		}

		if mp7hZaO := an0_Eja[0]; mp7hZaO._BtGRMo2 == abiStepStack {
			for _, fNFezZHudDS := range fWunZ5qrKW {
				m4wOIq :=  /*line :1*/ lxkP7oV7sm(a2T1BCJOcmaY, mp7hZaO.ciJuo5C+fNFezZHudDS.v9du88xTkB, "precomputed stack offset")
				switch fNFezZHudDS._BtGRMo2 {
				case abiStepPointer:

					*(* /*line :1*/ unsafe.Pointer)(m4wOIq) = xDAZAoprC2MC.Ptrs[fNFezZHudDS.pCQRHPfHe]
				case abiStepIntReg:
					 /*line :1*/ hchALWv(xDAZAoprC2MC, fNFezZHudDS.pCQRHPfHe, fNFezZHudDS.jLZh5ayKa9mx, m4wOIq)
				case abiStepFloatReg:
					 /*line :1*/ kh2Bx1R9(xDAZAoprC2MC, fNFezZHudDS.kT87Jzba6Sl, fNFezZHudDS.jLZh5ayKa9mx, m4wOIq)
				default:
					 /*line :1*/ panic("unexpected value step")
				}
			}
			continue
		}

		if  /*line :1*/ len(fWunZ5qrKW) !=  /*line :1*/ len(an0_Eja) {

			 /*line :1*/ panic("method ABI and value ABI don't align")
		}
		for hXZpj1nTZ, fNFezZHudDS := range fWunZ5qrKW {
			mp7hZaO := an0_Eja[hXZpj1nTZ]
			if mp7hZaO._BtGRMo2 != fNFezZHudDS._BtGRMo2 {
				 /*line :1*/ panic("method ABI and value ABI don't align")
			}
			switch fNFezZHudDS._BtGRMo2 {
			case abiStepPointer:

				qtyCUf.Ptrs[mp7hZaO.pCQRHPfHe] = xDAZAoprC2MC.Ptrs[fNFezZHudDS.pCQRHPfHe]
				fallthrough
			case abiStepIntReg:
				qtyCUf.Ints[mp7hZaO.pCQRHPfHe] = xDAZAoprC2MC.Ints[fNFezZHudDS.pCQRHPfHe]
			case abiStepFloatReg:
				qtyCUf.Floats[mp7hZaO.kT87Jzba6Sl] = xDAZAoprC2MC.Floats[fNFezZHudDS.kT87Jzba6Sl]
			default:
				 /*line :1*/ panic("unexpected value step")
			}
		}
	}

	tNa8in0Hxfc :=  /*line :1*/ kasvEYpX3A.Size()

	tNa8in0Hxfc =  /*line :1*/ rlY5YEji1B4(tNa8in0Hxfc, goarch.PtrSize)
	tNa8in0Hxfc += mtlKwzeEO4.hq8R8Dk

	qtyCUf.ReturnIsPtr = mtlKwzeEO4.aHHpIUai1ke

	 /*line :1*/ d7AWjWcFH(kasvEYpX3A, eKIiSoi0ulS, a2T1BCJOcmaY,  /*line :1*/ uint32( /*line :1*/ kasvEYpX3A.Size()),  /*line :1*/ uint32(mtlKwzeEO4.eBcFPCE),  /*line :1*/ uint32(tNa8in0Hxfc), &qtyCUf)

	if xDAZAoprC2MC != nil {
		*xDAZAoprC2MC = qtyCUf
	}
	if mptzN13jL_a :=  /*line :1*/ kasvEYpX3A.Size() - mtlKwzeEO4.eBcFPCE; mptzN13jL_a > 0 {
		zzB8dO3p4N :=  /*line :1*/ lxkP7oV7sm(o4PQFZaVsd, aI3R_aim5Ht.eBcFPCE, "valueFrame's size > retOffset")
		zMNkclt :=  /*line :1*/ lxkP7oV7sm(a2T1BCJOcmaY, mtlKwzeEO4.eBcFPCE, "methodFrame's size > retOffset")

		 /*line :1*/ pqAEZr6n(zzB8dO3p4N, zMNkclt, mptzN13jL_a)
	}

	*lATw9_ = true

	 /*line :1*/ q2gsmEF(kasvEYpX3A, a2T1BCJOcmaY)
	 /*line :1*/ lpyVCliC5UbT.Put(a2T1BCJOcmaY)

	 /*line :1*/ runtime.KeepAlive(m_WiSGQZjkA)

	 /*line :1*/ runtime.KeepAlive(xDAZAoprC2MC)
}

func wtBvIFaaYS(jToThzM func([]Value) []Value) string {
	gXCmQp9Bo := *(* /*line :1*/ uintptr)( /*line :1*/ unsafe.Pointer(&jToThzM))
	k82JHyzAdsC :=  /*line :1*/ runtime.FuncForPC(gXCmQp9Bo)
	if k82JHyzAdsC != nil {
		return  /*line :1*/ k82JHyzAdsC.Name()
	}
	return "closure"
}

func (f8NmcatRx Value) Cap() int {

	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == Slice {
		return (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr).Rsh30O4HaDX
	}
	return  /*line :1*/ f8NmcatRx.vIPweXwTN76T()
}

func (f8NmcatRx Value) vIPweXwTN76T() int {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Array:
		return  /*line :1*/ f8NmcatRx.lm5DH3().Len()
	case Chan:
		return  /*line :1*/ qUeiwmjb( /*line :1*/ f8NmcatRx.t_VyPi())
	case Ptr:
		if  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() == abi.Array {
			return  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Len()
		}
		 /*line :1*/ panic("reflect: call of reflect.Value.Cap on ptr to non-array Value")
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Cap",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) Close() {
	 /*line :1*/ f8NmcatRx.zc9yxv(Chan)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	 /*line :1*/ yaFXieDVoP( /*line :1*/ f8NmcatRx.t_VyPi())
}

func (f8NmcatRx Value) CanComplex() bool {
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Complex64, Complex128:
		return true
	default:
		return false
	}
}

func (f8NmcatRx Value) Complex() complex128 {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Complex64:
		return  /*line :1*/ complex128(*(* /*line :1*/ complex64)(f8NmcatRx.ptr))
	case Complex128:
		return *(* /*line :1*/ complex128)(f8NmcatRx.ptr)
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Complex",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) Elem() Value {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Interface:
		var ique858 any
		if  /*line :1*/ f8NmcatRx.lm5DH3().NumMethod() == 0 {
			ique858 = *(* /*line :1*/ any)(f8NmcatRx.ptr)
		} else {
			ique858 = ( /*line :1*/ any)(*(*interface {
				 /*line :1*/ M()
			})(f8NmcatRx.ptr))
		}
		uiTlN8bdKm :=  /*line :1*/ i7YECW6sP91m(ique858)
		if uiTlN8bdKm.flag != 0 {
			uiTlN8bdKm.flag |=  /*line :1*/ f8NmcatRx.flag.bQ77Iz()
		}
		return uiTlN8bdKm
	case Pointer:
		cf3fCV8ayFq := f8NmcatRx.ptr
		if f8NmcatRx.flag&flagIndir != 0 {
			if  /*line :1*/ gGK1Oc( /*line :1*/ f8NmcatRx.lm5DH3()) {

				if ! /*line :1*/ dvvw1_IpyIwu(*(* /*line :1*/ uintptr)(cf3fCV8ayFq)) {
					 /*line :1*/ panic("reflect: reflect.Value.Elem on an invalid notinheap pointer")
				}
			}
			cf3fCV8ayFq = *(* /*line :1*/ unsafe.Pointer)(cf3fCV8ayFq)
		}

		if cf3fCV8ayFq == nil {
			return Value{}
		}
		mP7QqSVInV := (* /*line :1*/ uSkasB)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		lm5DH3 := mP7QqSVInV.Elem
		iiZKy9PE3 := f8NmcatRx.flag&flagRO | flagIndir | flagAddr
		iiZKy9PE3 |=  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())
		return Value{lm5DH3, cf3fCV8ayFq, iiZKy9PE3}
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Elem",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) Field(hXZpj1nTZ int) Value {
	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() != Struct {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Field",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
	}
	mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
	if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint( /*line :1*/ len(mP7QqSVInV.Fields)) {
		 /*line :1*/ panic("reflect: Field index out of range")
	}
	fBabaKImKM2F := &mP7QqSVInV.Fields[hXZpj1nTZ]
	lm5DH3 := fBabaKImKM2F.Typ

	iiZKy9PE3 := f8NmcatRx.flag&(flagStickyRO|flagIndir|flagAddr) |  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())

	if ! /*line :1*/ fBabaKImKM2F.Name.IsExported() {
		if  /*line :1*/ fBabaKImKM2F.Embedded() {
			iiZKy9PE3 |= flagEmbedRO
		} else {
			iiZKy9PE3 |= flagStickyRO
		}
	}

	cf3fCV8ayFq :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, fBabaKImKM2F.Offset, "same as non-reflect &v.field")
	return Value{lm5DH3, cf3fCV8ayFq, iiZKy9PE3}
}

func (f8NmcatRx Value) FieldByIndex(hEqkR0O []int) Value {
	if  /*line :1*/ len(hEqkR0O) == 1 {
		return  /*line :1*/ f8NmcatRx.Field(hEqkR0O[0])
	}
	 /*line :1*/ f8NmcatRx.zc9yxv(Struct)
	for hXZpj1nTZ, uiTlN8bdKm := range hEqkR0O {
		if hXZpj1nTZ > 0 {
			if  /*line :1*/ f8NmcatRx.Kind() == Pointer &&  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() == abi.Struct {
				if  /*line :1*/ f8NmcatRx.IsNil() {
					 /*line :1*/ panic("reflect: indirection through nil pointer to embedded struct")
				}
				f8NmcatRx =  /*line :1*/ f8NmcatRx.Elem()
			}
		}
		f8NmcatRx =  /*line :1*/ f8NmcatRx.Field(uiTlN8bdKm)
	}
	return f8NmcatRx
}

func (f8NmcatRx Value) FieldByIndexErr(hEqkR0O []int) (Value, error) {
	if  /*line :1*/ len(hEqkR0O) == 1 {
		return  /*line :1*/ f8NmcatRx.Field(hEqkR0O[0]), nil
	}
	 /*line :1*/ f8NmcatRx.zc9yxv(Struct)
	for hXZpj1nTZ, uiTlN8bdKm := range hEqkR0O {
		if hXZpj1nTZ > 0 {
			if  /*line :1*/ f8NmcatRx.Kind() == Ptr &&  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() == abi.Struct {
				if  /*line :1*/ f8NmcatRx.IsNil() {
					return Value{},  /*line :1*/ errors.Su6g6hRoi9X("reflect: indirection through nil pointer to embedded struct field " +  /*line :1*/ jv91ZkF( /*line :1*/ f8NmcatRx.lm5DH3().Elem()))
				}
				f8NmcatRx =  /*line :1*/ f8NmcatRx.Elem()
			}
		}
		f8NmcatRx =  /*line :1*/ f8NmcatRx.Field(uiTlN8bdKm)
	}
	return f8NmcatRx, nil
}

func (f8NmcatRx Value) FieldByName(lGwAXqN4Hb string) Value {
	 /*line :1*/ f8NmcatRx.zc9yxv(Struct)
	if jToThzM, bDfzXlCm5Raf :=  /*line :1*/ j5Z_WRWLudek( /*line :1*/ f8NmcatRx.lm5DH3()).FieldByName(lGwAXqN4Hb); bDfzXlCm5Raf {
		return  /*line :1*/ f8NmcatRx.FieldByIndex(jToThzM.KOMPhI)
	}
	return Value{}
}

func (f8NmcatRx Value) FieldByNameFunc(xH9tRvQF4 func(string) bool) Value {
	if jToThzM, bDfzXlCm5Raf :=  /*line :1*/ j5Z_WRWLudek( /*line :1*/ f8NmcatRx.lm5DH3()).FieldByNameFunc(xH9tRvQF4); bDfzXlCm5Raf {
		return  /*line :1*/ f8NmcatRx.FieldByIndex(jToThzM.KOMPhI)
	}
	return Value{}
}

func (f8NmcatRx Value) CanFloat() bool {
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Float32, Float64:
		return true
	default:
		return false
	}
}

func (f8NmcatRx Value) Float() float64 {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Float32:
		return  /*line :1*/ float64(*(* /*line :1*/ float32)(f8NmcatRx.ptr))
	case Float64:
		return *(* /*line :1*/ float64)(f8NmcatRx.ptr)
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Float",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

var ahNMtQ =  /*line :1*/ gC11Vi( /*line :1*/ uint8(0))

func (f8NmcatRx Value) Index(hXZpj1nTZ int) Value {
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Array:
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint(mP7QqSVInV.Len) {
			 /*line :1*/ panic("reflect: array index out of range")
		}
		lm5DH3 := mP7QqSVInV.Elem
		dcDpXllPZ3e :=  /*line :1*/ uintptr(hXZpj1nTZ) *  /*line :1*/ lm5DH3.Size()

		w7xi3Bp :=  /*line :1*/ lxkP7oV7sm(f8NmcatRx.ptr, dcDpXllPZ3e, "same as &v[i], i < tt.len")
		iiZKy9PE3 := f8NmcatRx.flag&(flagIndir|flagAddr) |  /*line :1*/ f8NmcatRx.flag.bQ77Iz() |  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())
		return Value{lm5DH3, w7xi3Bp, iiZKy9PE3}

	case Slice:

		kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
		if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint(kXTvhUI4Zgg5.AJ2N5B) {
			 /*line :1*/ panic("reflect: slice index out of range")
		}
		mP7QqSVInV := (* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		lm5DH3 := mP7QqSVInV.Elem
		w7xi3Bp :=  /*line :1*/ rfJB4Et(kXTvhUI4Zgg5.TanDRHVvuLAL, hXZpj1nTZ,  /*line :1*/ lm5DH3.Size(), "i < s.Len")
		iiZKy9PE3 := flagAddr | flagIndir |  /*line :1*/ f8NmcatRx.flag.bQ77Iz() |  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())
		return Value{lm5DH3, w7xi3Bp, iiZKy9PE3}

	case String:
		kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.HrYeZlWHaf)(f8NmcatRx.ptr)
		if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint(kXTvhUI4Zgg5.SkAFPemec) {
			 /*line :1*/ panic("reflect: string index out of range")
		}
		e3IyyaQSEj :=  /*line :1*/ rfJB4Et(kXTvhUI4Zgg5.QOvpyywy, hXZpj1nTZ, 1, "i < s.Len")
		iiZKy9PE3 :=  /*line :1*/ f8NmcatRx.flag.bQ77Iz() |  /*line :1*/ flag(Uint8) | flagIndir
		return Value{ahNMtQ, e3IyyaQSEj, iiZKy9PE3}
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Index",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) CanInt() bool {
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Int, Int8, Int16, Int32, Int64:
		return true
	default:
		return false
	}
}

func (f8NmcatRx Value) Int() int64 {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	e3IyyaQSEj := f8NmcatRx.ptr
	switch lYOqUQga {
	case Int:
		return  /*line :1*/ int64(*(* /*line :1*/ int)(e3IyyaQSEj))
	case Int8:
		return  /*line :1*/ int64(*(* /*line :1*/ int8)(e3IyyaQSEj))
	case Int16:
		return  /*line :1*/ int64(*(* /*line :1*/ int16)(e3IyyaQSEj))
	case Int32:
		return  /*line :1*/ int64(*(* /*line :1*/ int32)(e3IyyaQSEj))
	case Int64:
		return *(* /*line :1*/ int64)(e3IyyaQSEj)
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Int",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) CanInterface() bool {
	if f8NmcatRx.flag == 0 {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.CanInterface", Invalid})
	}
	return f8NmcatRx.flag&flagRO == 0
}

func (f8NmcatRx Value) Interface() (hXZpj1nTZ any) {
	return  /*line :1*/ mySQReh3uV(f8NmcatRx, true)
}

func mySQReh3uV(f8NmcatRx Value, pnadaDfobanx bool) any {
	if f8NmcatRx.flag == 0 {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Interface", Invalid})
	}
	if pnadaDfobanx && f8NmcatRx.flag&flagRO != 0 {

		 /*line :1*/ panic("reflect.Value.Interface: cannot return value obtained from unexported field or method")
	}
	if f8NmcatRx.flag&flagMethod != 0 {
		f8NmcatRx =  /*line :1*/ evDSV6R("Interface", f8NmcatRx)
	}

	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == Interface {

		if  /*line :1*/ f8NmcatRx.NumMethod() == 0 {
			return *(* /*line :1*/ any)(f8NmcatRx.ptr)
		}
		return *(*interface {
			 /*line :1*/ M()
		})(f8NmcatRx.ptr)
	}

	return  /*line :1*/ hQWV0Z(f8NmcatRx)
}

func (f8NmcatRx Value) InterfaceData() [2]uintptr {
	 /*line :1*/ f8NmcatRx.zc9yxv(Interface)

	 /*line :1*/ vkf0Okel7G(f8NmcatRx.ptr)

	return *(*[2] /*line :1*/ uintptr)(f8NmcatRx.ptr)
}

func (f8NmcatRx Value) IsNil() bool {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Chan, Func, Map, Pointer, UnsafePointer:
		if f8NmcatRx.flag&flagMethod != 0 {
			return false
		}
		cf3fCV8ayFq := f8NmcatRx.ptr
		if f8NmcatRx.flag&flagIndir != 0 {
			cf3fCV8ayFq = *(* /*line :1*/ unsafe.Pointer)(cf3fCV8ayFq)
		}
		return cf3fCV8ayFq == nil
	case Interface, Slice:

		return *(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr) == nil
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.IsNil",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) IsValid() bool {
	return f8NmcatRx.flag != 0
}

func (f8NmcatRx Value) IsZero() bool {
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Bool:
		return ! /*line :1*/ f8NmcatRx.Bool()
	case Int, Int8, Int16, Int32, Int64:
		return  /*line :1*/ f8NmcatRx.Int() == 0
	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		return  /*line :1*/ f8NmcatRx.Uint() == 0
	case Float32, Float64:
		return  /*line :1*/ math.GKIRmoHE( /*line :1*/ f8NmcatRx.Float()) == 0
	case Complex64, Complex128:
		i5hAnuqALA :=  /*line :1*/ f8NmcatRx.Complex()
		return  /*line :1*/ math.GKIRmoHE( /*line :1*/ real(i5hAnuqALA)) == 0 &&  /*line :1*/ math.GKIRmoHE( /*line :1*/ imag(i5hAnuqALA)) == 0
	case Array:

		if  /*line :1*/ f8NmcatRx.lm5DH3().Equal != nil &&  /*line :1*/ f8NmcatRx.lm5DH3().Size() <= maxZero {
			if f8NmcatRx.flag&flagIndir == 0 {
				return f8NmcatRx.ptr == nil
			}

			return  /*line :1*/ f8NmcatRx.lm5DH3().Equal( /*line :1*/ m6KMjaQj(f8NmcatRx.ptr),  /*line :1*/ unsafe.Pointer(&mDOofGs5[0]))
		}

		krtR1HwEan :=  /*line :1*/ f8NmcatRx.Len()
		for hXZpj1nTZ := 0; hXZpj1nTZ < krtR1HwEan; hXZpj1nTZ++ {
			if ! /*line :1*/ f8NmcatRx.Index(hXZpj1nTZ).IsZero() {
				return false
			}
		}
		return true
	case Chan, Func, Interface, Map, Pointer, Slice, UnsafePointer:
		return  /*line :1*/ f8NmcatRx.IsNil()
	case String:
		return  /*line :1*/ f8NmcatRx.Len() == 0
	case Struct:

		if  /*line :1*/ f8NmcatRx.lm5DH3().Equal != nil &&  /*line :1*/ f8NmcatRx.lm5DH3().Size() <= maxZero {
			if f8NmcatRx.flag&flagIndir == 0 {
				return f8NmcatRx.ptr == nil
			}

			return  /*line :1*/ f8NmcatRx.lm5DH3().Equal( /*line :1*/ m6KMjaQj(f8NmcatRx.ptr),  /*line :1*/ unsafe.Pointer(&mDOofGs5[0]))
		}

		krtR1HwEan :=  /*line :1*/ f8NmcatRx.NumField()
		for hXZpj1nTZ := 0; hXZpj1nTZ < krtR1HwEan; hXZpj1nTZ++ {
			if ! /*line :1*/ f8NmcatRx.Field(hXZpj1nTZ).IsZero() {
				return false
			}
		}
		return true
	default:

		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.IsZero",  /*line :1*/ f8NmcatRx.Kind()})
	}
}

func (f8NmcatRx Value) SetZero() {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Bool:
		*(* /*line :1*/ bool)(f8NmcatRx.ptr) = false
	case Int:
		*(* /*line :1*/ int)(f8NmcatRx.ptr) = 0
	case Int8:
		*(* /*line :1*/ int8)(f8NmcatRx.ptr) = 0
	case Int16:
		*(* /*line :1*/ int16)(f8NmcatRx.ptr) = 0
	case Int32:
		*(* /*line :1*/ int32)(f8NmcatRx.ptr) = 0
	case Int64:
		*(* /*line :1*/ int64)(f8NmcatRx.ptr) = 0
	case Uint:
		*(* /*line :1*/ uint)(f8NmcatRx.ptr) = 0
	case Uint8:
		*(* /*line :1*/ uint8)(f8NmcatRx.ptr) = 0
	case Uint16:
		*(* /*line :1*/ uint16)(f8NmcatRx.ptr) = 0
	case Uint32:
		*(* /*line :1*/ uint32)(f8NmcatRx.ptr) = 0
	case Uint64:
		*(* /*line :1*/ uint64)(f8NmcatRx.ptr) = 0
	case Uintptr:
		*(* /*line :1*/ uintptr)(f8NmcatRx.ptr) = 0
	case Float32:
		*(* /*line :1*/ float32)(f8NmcatRx.ptr) = 0
	case Float64:
		*(* /*line :1*/ float64)(f8NmcatRx.ptr) = 0
	case Complex64:
		*(* /*line :1*/ complex64)(f8NmcatRx.ptr) = 0
	case Complex128:
		*(* /*line :1*/ complex128)(f8NmcatRx.ptr) = 0
	case String:
		*(* /*line :1*/ string)(f8NmcatRx.ptr) = ""
	case Slice:
		*(* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr) = unsafeheader.AKoOflOZ{}
	case Interface:
		*(*[2] /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr) = [2]unsafe.Pointer{}
	case Chan, Func, Map, Pointer, UnsafePointer:
		*(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr) = nil
	case Array, Struct:
		 /*line :1*/ q2gsmEF( /*line :1*/ f8NmcatRx.lm5DH3(), f8NmcatRx.ptr)
	default:

		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.SetZero",  /*line :1*/ f8NmcatRx.Kind()})
	}
}

func (f8NmcatRx Value) Kind() WzYBjsQL0 {
	return  /*line :1*/ f8NmcatRx.zRPbQda4dT()
}

func (f8NmcatRx Value) Len() int {

	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == Slice {
		return (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr).AJ2N5B
	}
	return  /*line :1*/ f8NmcatRx.isNaM9()
}

func (f8NmcatRx Value) isNaM9() int {
	switch lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT(); lYOqUQga {
	case Array:
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		return  /*line :1*/ int(mP7QqSVInV.Len)
	case Chan:
		return  /*line :1*/ joS0Abp( /*line :1*/ f8NmcatRx.t_VyPi())
	case Map:
		return  /*line :1*/ n6Tt0Kac( /*line :1*/ f8NmcatRx.t_VyPi())
	case String:

		return (* /*line :1*/ unsafeheader.HrYeZlWHaf)(f8NmcatRx.ptr).SkAFPemec
	case Ptr:
		if  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() == abi.Array {
			return  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Len()
		}
		 /*line :1*/ panic("reflect: call of reflect.Value.Len on ptr to non-array Value")
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Len",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

var huWmH9B2jl =  /*line :1*/ gC11Vi("")

func (f8NmcatRx Value) MapIndex(ccj6oONwQVN Value) Value {
	 /*line :1*/ f8NmcatRx.zc9yxv(Map)
	mP7QqSVInV := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))

	var dVDneyuGq unsafe.Pointer
	if (mP7QqSVInV.Key == huWmH9B2jl ||  /*line :1*/ ccj6oONwQVN.zRPbQda4dT() == String) && mP7QqSVInV.Key ==  /*line :1*/ ccj6oONwQVN.lm5DH3() &&  /*line :1*/ mP7QqSVInV.Elem.Size() <= maxValSize {
		lYOqUQga := *(* /*line :1*/ string)(ccj6oONwQVN.ptr)
		dVDneyuGq =  /*line :1*/ bjogv5( /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ f8NmcatRx.t_VyPi(), lYOqUQga)
	} else {
		ccj6oONwQVN =  /*line :1*/ ccj6oONwQVN.hfrHMyTs1aE("reflect.Value.MapIndex", mP7QqSVInV.Key, nil)
		var lYOqUQga unsafe.Pointer
		if ccj6oONwQVN.flag&flagIndir != 0 {
			lYOqUQga = ccj6oONwQVN.ptr
		} else {
			lYOqUQga =  /*line :1*/ unsafe.Pointer(&ccj6oONwQVN.ptr)
		}
		dVDneyuGq =  /*line :1*/ unynMPYi( /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ f8NmcatRx.t_VyPi(), lYOqUQga)
	}
	if dVDneyuGq == nil {
		return Value{}
	}
	lm5DH3 := mP7QqSVInV.Elem
	iiZKy9PE3 := ( /*line :1*/ f8NmcatRx.flag | ccj6oONwQVN.flag).bQ77Iz()
	iiZKy9PE3 |=  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())
	return  /*line :1*/ z9iqp94SQ(lm5DH3, iiZKy9PE3, dVDneyuGq)
}

func (f8NmcatRx Value) MapKeys() []Value {
	 /*line :1*/ f8NmcatRx.zc9yxv(Map)
	mP7QqSVInV := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
	xnmYrZliL := mP7QqSVInV.Key

	iiZKy9PE3 :=  /*line :1*/ f8NmcatRx.flag.bQ77Iz() |  /*line :1*/ flag( /*line :1*/ xnmYrZliL.Kind())

	hek9DV :=  /*line :1*/ f8NmcatRx.t_VyPi()
	gxTZ2uHyNSo :=  /*line :1*/ int(0)
	if hek9DV != nil {
		gxTZ2uHyNSo =  /*line :1*/ n6Tt0Kac(hek9DV)
	}
	var gPDgVVB cLVKdW_vuouD
	 /*line :1*/ gRAVyD( /*line :1*/ f8NmcatRx.lm5DH3(), hek9DV, &gPDgVVB)
	h6TmbCZ :=  /*line :1*/ make([]Value, gxTZ2uHyNSo)
	var hXZpj1nTZ int
	for hXZpj1nTZ = 0; hXZpj1nTZ <  /*line :1*/ len(h6TmbCZ); hXZpj1nTZ++ {
		ccj6oONwQVN :=  /*line :1*/ cgg9ghkrEjoF(&gPDgVVB)
		if ccj6oONwQVN == nil {

			break
		}
		h6TmbCZ[hXZpj1nTZ] =  /*line :1*/ z9iqp94SQ(xnmYrZliL, iiZKy9PE3, ccj6oONwQVN)
		 /*line :1*/ i5Fv9sOiLwR(&gPDgVVB)
	}
	return h6TmbCZ[:hXZpj1nTZ]
}

type cLVKdW_vuouD struct {
	oK9WFgSTaC	unsafe.Pointer
	mk1kEF	unsafe.Pointer
	bAWQcqaakJz	unsafe.Pointer
	uh7LgMvqA_a9	unsafe.Pointer
	asoi5j5	unsafe.Pointer
	__uqlm	unsafe.Pointer
	krzILo	*[]unsafe.Pointer
	jAHxO0A	*[]unsafe.Pointer
	pSovMxXar	uintptr
	hnsBwovAqr	uint8
	bvpEOd	bool
	JOhVNB	uint8
	kwjYahwc5XH	uint8
	aFGLYSda	uintptr
	sC7Oh7tj2XyT	uintptr
}

func (cJS_wcOmyXDz *cLVKdW_vuouD) _qw1ykQaw() bool {
	return cJS_wcOmyXDz.bAWQcqaakJz != nil
}

type A7aXtMx struct {
	lABuBUlkSTR9	Value
	rh4qq_PDlbO	cLVKdW_vuouD
}

func (x8vxBoVa *A7aXtMx) Key() Value {
	if ! /*line :1*/ x8vxBoVa.rh4qq_PDlbO._qw1ykQaw() {
		 /*line :1*/ panic("MapIter.Key called before Next")
	}
	crKfxQc4LC :=  /*line :1*/ cgg9ghkrEjoF(&x8vxBoVa.rh4qq_PDlbO)
	if crKfxQc4LC == nil {
		 /*line :1*/ panic("MapIter.Key called on exhausted iterator")
	}

	sw8_lJ := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer( /*line :1*/ x8vxBoVa.lABuBUlkSTR9.lm5DH3()))
	gCXRRw := sw8_lJ.Key
	return  /*line :1*/ z9iqp94SQ(gCXRRw,  /*line :1*/ x8vxBoVa.lABuBUlkSTR9.flag.bQ77Iz()| /*line :1*/ flag( /*line :1*/ gCXRRw.Kind()), crKfxQc4LC)
}

func (f8NmcatRx Value) SetIterKey(x8vxBoVa *A7aXtMx) {
	if ! /*line :1*/ x8vxBoVa.rh4qq_PDlbO._qw1ykQaw() {
		 /*line :1*/ panic("reflect: Value.SetIterKey called before Next")
	}
	crKfxQc4LC :=  /*line :1*/ cgg9ghkrEjoF(&x8vxBoVa.rh4qq_PDlbO)
	if crKfxQc4LC == nil {
		 /*line :1*/ panic("reflect: Value.SetIterKey called on exhausted iterator")
	}

	 /*line :1*/ f8NmcatRx.d5PnN9n()
	var htiV4jZ2U unsafe.Pointer
	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == Interface {
		htiV4jZ2U = f8NmcatRx.ptr
	}

	sw8_lJ := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer( /*line :1*/ x8vxBoVa.lABuBUlkSTR9.lm5DH3()))
	gCXRRw := sw8_lJ.Key

	 /*line :1*/ x8vxBoVa.lABuBUlkSTR9.rwcaztMf()
	ccj6oONwQVN := Value{gCXRRw, crKfxQc4LC, x8vxBoVa.lABuBUlkSTR9.flag |  /*line :1*/ flag( /*line :1*/ gCXRRw.Kind()) | flagIndir}
	ccj6oONwQVN =  /*line :1*/ ccj6oONwQVN.hfrHMyTs1aE("reflect.MapIter.SetKey",  /*line :1*/ f8NmcatRx.lm5DH3(), htiV4jZ2U)
	 /*line :1*/ l26oiKGd( /*line :1*/ f8NmcatRx.lm5DH3(), f8NmcatRx.ptr, ccj6oONwQVN.ptr)
}

func (x8vxBoVa *A7aXtMx) Value() Value {
	if ! /*line :1*/ x8vxBoVa.rh4qq_PDlbO._qw1ykQaw() {
		 /*line :1*/ panic("MapIter.Value called before Next")
	}
	aS05ZBzCg :=  /*line :1*/ tW5jaa7(&x8vxBoVa.rh4qq_PDlbO)
	if aS05ZBzCg == nil {
		 /*line :1*/ panic("MapIter.Value called on exhausted iterator")
	}

	sw8_lJ := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer( /*line :1*/ x8vxBoVa.lABuBUlkSTR9.lm5DH3()))
	pi36nyr31 := sw8_lJ.Elem
	return  /*line :1*/ z9iqp94SQ(pi36nyr31,  /*line :1*/ x8vxBoVa.lABuBUlkSTR9.flag.bQ77Iz()| /*line :1*/ flag( /*line :1*/ pi36nyr31.Kind()), aS05ZBzCg)
}

func (f8NmcatRx Value) SetIterValue(x8vxBoVa *A7aXtMx) {
	if ! /*line :1*/ x8vxBoVa.rh4qq_PDlbO._qw1ykQaw() {
		 /*line :1*/ panic("reflect: Value.SetIterValue called before Next")
	}
	aS05ZBzCg :=  /*line :1*/ tW5jaa7(&x8vxBoVa.rh4qq_PDlbO)
	if aS05ZBzCg == nil {
		 /*line :1*/ panic("reflect: Value.SetIterValue called on exhausted iterator")
	}

	 /*line :1*/ f8NmcatRx.d5PnN9n()
	var htiV4jZ2U unsafe.Pointer
	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == Interface {
		htiV4jZ2U = f8NmcatRx.ptr
	}

	sw8_lJ := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer( /*line :1*/ x8vxBoVa.lABuBUlkSTR9.lm5DH3()))
	pi36nyr31 := sw8_lJ.Elem

	 /*line :1*/ x8vxBoVa.lABuBUlkSTR9.rwcaztMf()
	dPe0EQn := Value{pi36nyr31, aS05ZBzCg, x8vxBoVa.lABuBUlkSTR9.flag |  /*line :1*/ flag( /*line :1*/ pi36nyr31.Kind()) | flagIndir}
	dPe0EQn =  /*line :1*/ dPe0EQn.hfrHMyTs1aE("reflect.MapIter.SetValue",  /*line :1*/ f8NmcatRx.lm5DH3(), htiV4jZ2U)
	 /*line :1*/ l26oiKGd( /*line :1*/ f8NmcatRx.lm5DH3(), f8NmcatRx.ptr, dPe0EQn.ptr)
}

func (x8vxBoVa *A7aXtMx) Next() bool {
	if ! /*line :1*/ x8vxBoVa.lABuBUlkSTR9.IsValid() {
		 /*line :1*/ panic("MapIter.Next called on an iterator that does not have an associated map Value")
	}
	if ! /*line :1*/ x8vxBoVa.rh4qq_PDlbO._qw1ykQaw() {
		 /*line :1*/ gRAVyD( /*line :1*/ x8vxBoVa.lABuBUlkSTR9.lm5DH3(),  /*line :1*/ x8vxBoVa.lABuBUlkSTR9.t_VyPi(), &x8vxBoVa.rh4qq_PDlbO)
	} else {
		if  /*line :1*/ cgg9ghkrEjoF(&x8vxBoVa.rh4qq_PDlbO) == nil {
			 /*line :1*/ panic("MapIter.Next called on exhausted iterator")
		}
		 /*line :1*/ i5Fv9sOiLwR(&x8vxBoVa.rh4qq_PDlbO)
	}
	return  /*line :1*/ cgg9ghkrEjoF(&x8vxBoVa.rh4qq_PDlbO) != nil
}

func (x8vxBoVa *A7aXtMx) Reset(f8NmcatRx Value) {
	if  /*line :1*/ f8NmcatRx.IsValid() {
		 /*line :1*/ f8NmcatRx.zc9yxv(Map)
	}
	x8vxBoVa.lABuBUlkSTR9 = f8NmcatRx
	x8vxBoVa.rh4qq_PDlbO = cLVKdW_vuouD{}
}

func (f8NmcatRx Value) MapRange() *A7aXtMx {

	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() != Map {
		 /*line :1*/ f8NmcatRx.oUqGoSmRpyu()
	}
	return &A7aXtMx{lABuBUlkSTR9: f8NmcatRx}
}

//go:noinline
func (jToThzM flag) oUqGoSmRpyu() {
	 /*line :1*/ jToThzM.zc9yxv(Map)
}

func z9iqp94SQ(lm5DH3 *abi.Type, iiZKy9PE3 flag, cf3fCV8ayFq unsafe.Pointer) Value {
	if  /*line :1*/ lm5DH3.IfaceIndir() {

		i5hAnuqALA :=  /*line :1*/ qscv_5(lm5DH3)
		 /*line :1*/ l26oiKGd(lm5DH3, i5hAnuqALA, cf3fCV8ayFq)
		return Value{lm5DH3, i5hAnuqALA, iiZKy9PE3 | flagIndir}
	}
	return Value{lm5DH3, *(* /*line :1*/ unsafe.Pointer)(cf3fCV8ayFq), iiZKy9PE3}
}

func (f8NmcatRx Value) Method(hXZpj1nTZ int) Value {
	if  /*line :1*/ f8NmcatRx.lm5DH3() == nil {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Method", Invalid})
	}
	if f8NmcatRx.flag&flagMethod != 0 ||  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint( /*line :1*/ j5Z_WRWLudek( /*line :1*/ f8NmcatRx.lm5DH3()).NumMethod()) {
		 /*line :1*/ panic("reflect: Method index out of range")
	}
	if  /*line :1*/ f8NmcatRx.lm5DH3().Kind() == abi.Interface &&  /*line :1*/ f8NmcatRx.IsNil() {
		 /*line :1*/ panic("reflect: Method on nil interface value")
	}
	iiZKy9PE3 :=  /*line :1*/ f8NmcatRx.flag.bQ77Iz() | (f8NmcatRx.flag & flagIndir)
	iiZKy9PE3 |=  /*line :1*/ flag(Func)
	iiZKy9PE3 |=  /*line :1*/ flag(hXZpj1nTZ)<<flagMethodShift | flagMethod
	return Value{ /*line :1*/ f8NmcatRx.lm5DH3(), f8NmcatRx.ptr, iiZKy9PE3}
}

func (f8NmcatRx Value) NumMethod() int {
	if  /*line :1*/ f8NmcatRx.lm5DH3() == nil {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.NumMethod", Invalid})
	}
	if f8NmcatRx.flag&flagMethod != 0 {
		return 0
	}
	return  /*line :1*/ j5Z_WRWLudek( /*line :1*/ f8NmcatRx.lm5DH3()).NumMethod()
}

func (f8NmcatRx Value) MethodByName(lGwAXqN4Hb string) Value {
	if  /*line :1*/ f8NmcatRx.lm5DH3() == nil {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.MethodByName", Invalid})
	}
	if f8NmcatRx.flag&flagMethod != 0 {
		return Value{}
	}
	hek9DV, bDfzXlCm5Raf :=  /*line :1*/ j5Z_WRWLudek( /*line :1*/ f8NmcatRx.lm5DH3()).MethodByName(lGwAXqN4Hb)
	if !bDfzXlCm5Raf {
		return Value{}
	}
	return  /*line :1*/ f8NmcatRx.Method(hek9DV.JhrNuhrH)
}

func (f8NmcatRx Value) NumField() int {
	 /*line :1*/ f8NmcatRx.zc9yxv(Struct)
	mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
	return  /*line :1*/ len(mP7QqSVInV.Fields)
}

func (f8NmcatRx Value) OverflowComplex(uiTlN8bdKm complex128) bool {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Complex64:
		return  /*line :1*/ wkzKpK5j( /*line :1*/ real(uiTlN8bdKm)) ||  /*line :1*/ wkzKpK5j( /*line :1*/ imag(uiTlN8bdKm))
	case Complex128:
		return false
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.OverflowComplex",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) OverflowFloat(uiTlN8bdKm float64) bool {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Float32:
		return  /*line :1*/ wkzKpK5j(uiTlN8bdKm)
	case Float64:
		return false
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.OverflowFloat",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func wkzKpK5j(uiTlN8bdKm float64) bool {
	if uiTlN8bdKm < 0 {
		uiTlN8bdKm = -uiTlN8bdKm
	}
	return math.MaxFloat32 < uiTlN8bdKm && uiTlN8bdKm <= math.MaxFloat64
}

func (f8NmcatRx Value) OverflowInt(uiTlN8bdKm int64) bool {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Int, Int8, Int16, Int32, Int64:
		yybPd_Drwb :=  /*line :1*/ f8NmcatRx.lm5DH3().Size() * 8
		gahjFGH8aB := (uiTlN8bdKm << (64 - yybPd_Drwb)) >> (64 - yybPd_Drwb)
		return uiTlN8bdKm != gahjFGH8aB
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.OverflowInt",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) OverflowUint(uiTlN8bdKm uint64) bool {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Uint, Uintptr, Uint8, Uint16, Uint32, Uint64:
		yybPd_Drwb :=  /*line :1*/ f8NmcatRx.typ_.Size() * 8
		gahjFGH8aB := (uiTlN8bdKm << (64 - yybPd_Drwb)) >> (64 - yybPd_Drwb)
		return uiTlN8bdKm != gahjFGH8aB
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.OverflowUint",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

//go:nocheckptr

func (f8NmcatRx Value) Pointer() uintptr {

	 /*line :1*/ vkf0Okel7G(f8NmcatRx.ptr)

	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Pointer:
		if  /*line :1*/ f8NmcatRx.lm5DH3().PtrBytes == 0 {
			w7xi3Bp := *(* /*line :1*/ uintptr)(f8NmcatRx.ptr)

			if ! /*line :1*/ dvvw1_IpyIwu(w7xi3Bp) {
				 /*line :1*/ panic("reflect: reflect.Value.Pointer on an invalid notinheap pointer")
			}
			return w7xi3Bp
		}
		fallthrough
	case Chan, Map, UnsafePointer:
		return  /*line :1*/ uintptr( /*line :1*/ f8NmcatRx.t_VyPi())
	case Func:
		if f8NmcatRx.flag&flagMethod != 0 {

			return  /*line :1*/ vTsZZIIlq4MX()
		}
		e3IyyaQSEj :=  /*line :1*/ f8NmcatRx.t_VyPi()

		if e3IyyaQSEj != nil {
			e3IyyaQSEj = *(* /*line :1*/ unsafe.Pointer)(e3IyyaQSEj)
		}
		return  /*line :1*/ uintptr(e3IyyaQSEj)

	case Slice:
		return  /*line :1*/ uintptr((* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr).TanDRHVvuLAL)
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Pointer",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

func (f8NmcatRx Value) Recv() (uiTlN8bdKm Value, bDfzXlCm5Raf bool) {
	 /*line :1*/ f8NmcatRx.zc9yxv(Chan)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	return  /*line :1*/ f8NmcatRx.wM9XCjNlp(false)
}

func (f8NmcatRx Value) wM9XCjNlp(j2eisdvd_at bool) (w7xi3Bp Value, bDfzXlCm5Raf bool) {
	mP7QqSVInV := (* /*line :1*/ aK0ovNmjQBD)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
	if  /*line :1*/ Svj3YDOP3fa(mP7QqSVInV.Dir)&RecvDir == 0 {
		 /*line :1*/ panic("reflect: recv on send-only channel")
	}
	sw8_lJ := mP7QqSVInV.Elem
	w7xi3Bp = Value{sw8_lJ, nil,  /*line :1*/ flag( /*line :1*/ sw8_lJ.Kind())}
	var e3IyyaQSEj unsafe.Pointer
	if  /*line :1*/ gGK1Oc(sw8_lJ) {
		e3IyyaQSEj =  /*line :1*/ qscv_5(sw8_lJ)
		w7xi3Bp.ptr = e3IyyaQSEj
		w7xi3Bp.flag |= flagIndir
	} else {
		e3IyyaQSEj =  /*line :1*/ unsafe.Pointer(&w7xi3Bp.ptr)
	}
	a_LOCfaJEYY1, bDfzXlCm5Raf :=  /*line :1*/ aXQaEP2( /*line :1*/ f8NmcatRx.t_VyPi(), j2eisdvd_at, e3IyyaQSEj)
	if !a_LOCfaJEYY1 {
		w7xi3Bp = Value{}
	}
	return
}

func (f8NmcatRx Value) Send(uiTlN8bdKm Value) {
	 /*line :1*/ f8NmcatRx.zc9yxv(Chan)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	 /*line :1*/ f8NmcatRx._eTbCaDQq(uiTlN8bdKm, false)
}

func (f8NmcatRx Value) _eTbCaDQq(uiTlN8bdKm Value, j2eisdvd_at bool) (a_LOCfaJEYY1 bool) {
	mP7QqSVInV := (* /*line :1*/ aK0ovNmjQBD)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
	if  /*line :1*/ Svj3YDOP3fa(mP7QqSVInV.Dir)&SendDir == 0 {
		 /*line :1*/ panic("reflect: send on recv-only channel")
	}
	 /*line :1*/ uiTlN8bdKm.rwcaztMf()
	uiTlN8bdKm =  /*line :1*/ uiTlN8bdKm.hfrHMyTs1aE("reflect.Value.Send", mP7QqSVInV.Elem, nil)
	var e3IyyaQSEj unsafe.Pointer
	if uiTlN8bdKm.flag&flagIndir != 0 {
		e3IyyaQSEj = uiTlN8bdKm.ptr
	} else {
		e3IyyaQSEj =  /*line :1*/ unsafe.Pointer(&uiTlN8bdKm.ptr)
	}
	return  /*line :1*/ nSsxqozFB( /*line :1*/ f8NmcatRx.t_VyPi(), e3IyyaQSEj, j2eisdvd_at)
}

func (f8NmcatRx Value) Set(uiTlN8bdKm Value) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ uiTlN8bdKm.rwcaztMf()
	var htiV4jZ2U unsafe.Pointer
	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == Interface {
		htiV4jZ2U = f8NmcatRx.ptr
	}
	uiTlN8bdKm =  /*line :1*/ uiTlN8bdKm.hfrHMyTs1aE("reflect.Set",  /*line :1*/ f8NmcatRx.lm5DH3(), htiV4jZ2U)
	if uiTlN8bdKm.flag&flagIndir != 0 {
		if uiTlN8bdKm.ptr ==  /*line :1*/ unsafe.Pointer(&mDOofGs5[0]) {
			 /*line :1*/ q2gsmEF( /*line :1*/ f8NmcatRx.lm5DH3(), f8NmcatRx.ptr)
		} else {
			 /*line :1*/ l26oiKGd( /*line :1*/ f8NmcatRx.lm5DH3(), f8NmcatRx.ptr, uiTlN8bdKm.ptr)
		}
	} else {
		*(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr) = uiTlN8bdKm.ptr
	}
}

func (f8NmcatRx Value) SetBool(uiTlN8bdKm bool) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(Bool)
	*(* /*line :1*/ bool)(f8NmcatRx.ptr) = uiTlN8bdKm
}

func (f8NmcatRx Value) SetBytes(uiTlN8bdKm []byte) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(Slice)
	if  /*line :1*/ j5Z_WRWLudek( /*line :1*/ f8NmcatRx.lm5DH3()).Elem().Kind() != Uint8 {
		 /*line :1*/ panic("reflect.Value.SetBytes of non-byte slice")
	}
	*(*[] /*line :1*/ byte)(f8NmcatRx.ptr) = uiTlN8bdKm
}

func (f8NmcatRx Value) oqaG08DN(uiTlN8bdKm []rune) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(Slice)
	if  /*line :1*/ f8NmcatRx.lm5DH3().Elem().Kind() != abi.Int32 {
		 /*line :1*/ panic("reflect.Value.setRunes of non-rune slice")
	}
	*(*[] /*line :1*/ rune)(f8NmcatRx.ptr) = uiTlN8bdKm
}

func (f8NmcatRx Value) SetComplex(uiTlN8bdKm complex128) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	switch lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT(); lYOqUQga {
	default:
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.SetComplex",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
	case Complex64:
		*(* /*line :1*/ complex64)(f8NmcatRx.ptr) =  /*line :1*/ complex64(uiTlN8bdKm)
	case Complex128:
		*(* /*line :1*/ complex128)(f8NmcatRx.ptr) = uiTlN8bdKm
	}
}

func (f8NmcatRx Value) SetFloat(uiTlN8bdKm float64) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	switch lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT(); lYOqUQga {
	default:
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.SetFloat",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
	case Float32:
		*(* /*line :1*/ float32)(f8NmcatRx.ptr) =  /*line :1*/ float32(uiTlN8bdKm)
	case Float64:
		*(* /*line :1*/ float64)(f8NmcatRx.ptr) = uiTlN8bdKm
	}
}

func (f8NmcatRx Value) SetInt(uiTlN8bdKm int64) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	switch lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT(); lYOqUQga {
	default:
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.SetInt",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
	case Int:
		*(* /*line :1*/ int)(f8NmcatRx.ptr) =  /*line :1*/ int(uiTlN8bdKm)
	case Int8:
		*(* /*line :1*/ int8)(f8NmcatRx.ptr) =  /*line :1*/ int8(uiTlN8bdKm)
	case Int16:
		*(* /*line :1*/ int16)(f8NmcatRx.ptr) =  /*line :1*/ int16(uiTlN8bdKm)
	case Int32:
		*(* /*line :1*/ int32)(f8NmcatRx.ptr) =  /*line :1*/ int32(uiTlN8bdKm)
	case Int64:
		*(* /*line :1*/ int64)(f8NmcatRx.ptr) = uiTlN8bdKm
	}
}

func (f8NmcatRx Value) SetLen(krtR1HwEan int) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(Slice)
	kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
	if  /*line :1*/ uint(krtR1HwEan) >  /*line :1*/ uint(kXTvhUI4Zgg5.Rsh30O4HaDX) {
		 /*line :1*/ panic("reflect: slice length out of range in SetLen")
	}
	kXTvhUI4Zgg5.AJ2N5B = krtR1HwEan
}

func (f8NmcatRx Value) SetCap(krtR1HwEan int) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(Slice)
	kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
	if krtR1HwEan < kXTvhUI4Zgg5.AJ2N5B || krtR1HwEan > kXTvhUI4Zgg5.Rsh30O4HaDX {
		 /*line :1*/ panic("reflect: slice capacity out of range in SetCap")
	}
	kXTvhUI4Zgg5.Rsh30O4HaDX = krtR1HwEan
}

func (f8NmcatRx Value) SetMapIndex(ccj6oONwQVN, dPe0EQn Value) {
	 /*line :1*/ f8NmcatRx.zc9yxv(Map)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	 /*line :1*/ ccj6oONwQVN.rwcaztMf()
	mP7QqSVInV := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))

	if (mP7QqSVInV.Key == huWmH9B2jl ||  /*line :1*/ ccj6oONwQVN.zRPbQda4dT() == String) && mP7QqSVInV.Key ==  /*line :1*/ ccj6oONwQVN.lm5DH3() &&  /*line :1*/ mP7QqSVInV.Elem.Size() <= maxValSize {
		lYOqUQga := *(* /*line :1*/ string)(ccj6oONwQVN.ptr)
		if  /*line :1*/ dPe0EQn.lm5DH3() == nil {
			 /*line :1*/ um8hBA( /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ f8NmcatRx.t_VyPi(), lYOqUQga)
			return
		}
		 /*line :1*/ dPe0EQn.rwcaztMf()
		dPe0EQn =  /*line :1*/ dPe0EQn.hfrHMyTs1aE("reflect.Value.SetMapIndex", mP7QqSVInV.Elem, nil)
		var dVDneyuGq unsafe.Pointer
		if dPe0EQn.flag&flagIndir != 0 {
			dVDneyuGq = dPe0EQn.ptr
		} else {
			dVDneyuGq =  /*line :1*/ unsafe.Pointer(&dPe0EQn.ptr)
		}
		 /*line :1*/ i6bZUn( /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ f8NmcatRx.t_VyPi(), lYOqUQga, dVDneyuGq)
		return
	}

	ccj6oONwQVN =  /*line :1*/ ccj6oONwQVN.hfrHMyTs1aE("reflect.Value.SetMapIndex", mP7QqSVInV.Key, nil)
	var lYOqUQga unsafe.Pointer
	if ccj6oONwQVN.flag&flagIndir != 0 {
		lYOqUQga = ccj6oONwQVN.ptr
	} else {
		lYOqUQga =  /*line :1*/ unsafe.Pointer(&ccj6oONwQVN.ptr)
	}
	if  /*line :1*/ dPe0EQn.lm5DH3() == nil {
		 /*line :1*/ a0tirBasjs( /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ f8NmcatRx.t_VyPi(), lYOqUQga)
		return
	}
	 /*line :1*/ dPe0EQn.rwcaztMf()
	dPe0EQn =  /*line :1*/ dPe0EQn.hfrHMyTs1aE("reflect.Value.SetMapIndex", mP7QqSVInV.Elem, nil)
	var dVDneyuGq unsafe.Pointer
	if dPe0EQn.flag&flagIndir != 0 {
		dVDneyuGq = dPe0EQn.ptr
	} else {
		dVDneyuGq =  /*line :1*/ unsafe.Pointer(&dPe0EQn.ptr)
	}
	 /*line :1*/ vJbWE1laDY( /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ f8NmcatRx.t_VyPi(), lYOqUQga, dVDneyuGq)
}

func (f8NmcatRx Value) SetUint(uiTlN8bdKm uint64) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	switch lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT(); lYOqUQga {
	default:
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.SetUint",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
	case Uint:
		*(* /*line :1*/ uint)(f8NmcatRx.ptr) =  /*line :1*/ uint(uiTlN8bdKm)
	case Uint8:
		*(* /*line :1*/ uint8)(f8NmcatRx.ptr) =  /*line :1*/ uint8(uiTlN8bdKm)
	case Uint16:
		*(* /*line :1*/ uint16)(f8NmcatRx.ptr) =  /*line :1*/ uint16(uiTlN8bdKm)
	case Uint32:
		*(* /*line :1*/ uint32)(f8NmcatRx.ptr) =  /*line :1*/ uint32(uiTlN8bdKm)
	case Uint64:
		*(* /*line :1*/ uint64)(f8NmcatRx.ptr) = uiTlN8bdKm
	case Uintptr:
		*(* /*line :1*/ uintptr)(f8NmcatRx.ptr) =  /*line :1*/ uintptr(uiTlN8bdKm)
	}
}

func (f8NmcatRx Value) SetPointer(uiTlN8bdKm unsafe.Pointer) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(UnsafePointer)
	*(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr) = uiTlN8bdKm
}

func (f8NmcatRx Value) SetString(uiTlN8bdKm string) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(String)
	*(* /*line :1*/ string)(f8NmcatRx.ptr) = uiTlN8bdKm
}

func (f8NmcatRx Value) Slice(hXZpj1nTZ, kNVeJaGxs int) Value {
	var (
		xUg7Px	int
		lm5DH3	*vwvVHq
		tDufzwDs061	unsafe.Pointer
	)
	switch zRPbQda4dT :=  /*line :1*/ f8NmcatRx.zRPbQda4dT(); zRPbQda4dT {
	default:
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Slice",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})

	case Array:
		if f8NmcatRx.flag&flagAddr == 0 {
			 /*line :1*/ panic("reflect.Value.Slice: slice of unaddressable array")
		}
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		xUg7Px =  /*line :1*/ int(mP7QqSVInV.Len)
		lm5DH3 = (* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer(mP7QqSVInV.Slice))
		tDufzwDs061 = f8NmcatRx.ptr

	case Slice:
		lm5DH3 = (* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
		tDufzwDs061 = kXTvhUI4Zgg5.TanDRHVvuLAL
		xUg7Px = kXTvhUI4Zgg5.Rsh30O4HaDX

	case String:
		kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.HrYeZlWHaf)(f8NmcatRx.ptr)
		if hXZpj1nTZ < 0 || kNVeJaGxs < hXZpj1nTZ || kNVeJaGxs > kXTvhUI4Zgg5.SkAFPemec {
			 /*line :1*/ panic("reflect.Value.Slice: string slice index out of bounds")
		}
		var sw8_lJ unsafeheader.HrYeZlWHaf
		if hXZpj1nTZ < kXTvhUI4Zgg5.SkAFPemec {
			sw8_lJ = unsafeheader.HrYeZlWHaf{QOvpyywy:  /*line :1*/ rfJB4Et(kXTvhUI4Zgg5.QOvpyywy, hXZpj1nTZ, 1, "i < s.Len"), SkAFPemec: kNVeJaGxs - hXZpj1nTZ}
		}
		return Value{ /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ unsafe.Pointer(&sw8_lJ), f8NmcatRx.flag}
	}

	if hXZpj1nTZ < 0 || kNVeJaGxs < hXZpj1nTZ || kNVeJaGxs > xUg7Px {
		 /*line :1*/ panic("reflect.Value.Slice: slice index out of bounds")
	}

	var uiTlN8bdKm []unsafe.Pointer

	kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)( /*line :1*/ unsafe.Pointer(&uiTlN8bdKm))
	kXTvhUI4Zgg5.AJ2N5B = kNVeJaGxs - hXZpj1nTZ
	kXTvhUI4Zgg5.Rsh30O4HaDX = xUg7Px - hXZpj1nTZ
	if xUg7Px-hXZpj1nTZ > 0 {
		kXTvhUI4Zgg5.TanDRHVvuLAL =  /*line :1*/ rfJB4Et(tDufzwDs061, hXZpj1nTZ,  /*line :1*/ lm5DH3.Elem.Size(), "i < cap")
	} else {

		kXTvhUI4Zgg5.TanDRHVvuLAL = tDufzwDs061
	}

	iiZKy9PE3 :=  /*line :1*/ f8NmcatRx.flag.bQ77Iz() | flagIndir |  /*line :1*/ flag(Slice)
	return Value{ /*line :1*/ lm5DH3.Common(),  /*line :1*/ unsafe.Pointer(&uiTlN8bdKm), iiZKy9PE3}
}

func (f8NmcatRx Value) Slice3(hXZpj1nTZ, kNVeJaGxs, lYOqUQga int) Value {
	var (
		xUg7Px	int
		lm5DH3	*vwvVHq
		tDufzwDs061	unsafe.Pointer
	)
	switch zRPbQda4dT :=  /*line :1*/ f8NmcatRx.zRPbQda4dT(); zRPbQda4dT {
	default:
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Slice3",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})

	case Array:
		if f8NmcatRx.flag&flagAddr == 0 {
			 /*line :1*/ panic("reflect.Value.Slice3: slice of unaddressable array")
		}
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		xUg7Px =  /*line :1*/ int(mP7QqSVInV.Len)
		lm5DH3 = (* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer(mP7QqSVInV.Slice))
		tDufzwDs061 = f8NmcatRx.ptr

	case Slice:
		lm5DH3 = (* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
		tDufzwDs061 = kXTvhUI4Zgg5.TanDRHVvuLAL
		xUg7Px = kXTvhUI4Zgg5.Rsh30O4HaDX
	}

	if hXZpj1nTZ < 0 || kNVeJaGxs < hXZpj1nTZ || lYOqUQga < kNVeJaGxs || lYOqUQga > xUg7Px {
		 /*line :1*/ panic("reflect.Value.Slice3: slice index out of bounds")
	}

	var uiTlN8bdKm []unsafe.Pointer

	kXTvhUI4Zgg5 := (* /*line :1*/ unsafeheader.AKoOflOZ)( /*line :1*/ unsafe.Pointer(&uiTlN8bdKm))
	kXTvhUI4Zgg5.AJ2N5B = kNVeJaGxs - hXZpj1nTZ
	kXTvhUI4Zgg5.Rsh30O4HaDX = lYOqUQga - hXZpj1nTZ
	if lYOqUQga-hXZpj1nTZ > 0 {
		kXTvhUI4Zgg5.TanDRHVvuLAL =  /*line :1*/ rfJB4Et(tDufzwDs061, hXZpj1nTZ,  /*line :1*/ lm5DH3.Elem.Size(), "i < k <= cap")
	} else {

		kXTvhUI4Zgg5.TanDRHVvuLAL = tDufzwDs061
	}

	iiZKy9PE3 :=  /*line :1*/ f8NmcatRx.flag.bQ77Iz() | flagIndir |  /*line :1*/ flag(Slice)
	return Value{ /*line :1*/ lm5DH3.Common(),  /*line :1*/ unsafe.Pointer(&uiTlN8bdKm), iiZKy9PE3}
}

func (f8NmcatRx Value) String() string {

	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == String {
		return *(* /*line :1*/ string)(f8NmcatRx.ptr)
	}
	return  /*line :1*/ f8NmcatRx.nhnfyBcIVBa()
}

func (f8NmcatRx Value) nhnfyBcIVBa() string {
	if  /*line :1*/ f8NmcatRx.zRPbQda4dT() == Invalid {
		return "<invalid Value>"
	}

	return "<" +  /*line :1*/ f8NmcatRx.Type().String() + " Value>"
}

func (f8NmcatRx Value) TryRecv() (uiTlN8bdKm Value, bDfzXlCm5Raf bool) {
	 /*line :1*/ f8NmcatRx.zc9yxv(Chan)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	return  /*line :1*/ f8NmcatRx.wM9XCjNlp(true)
}

func (f8NmcatRx Value) TrySend(uiTlN8bdKm Value) bool {
	 /*line :1*/ f8NmcatRx.zc9yxv(Chan)
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	return  /*line :1*/ f8NmcatRx._eTbCaDQq(uiTlN8bdKm, true)
}

func (f8NmcatRx Value) Type() YJh989LTZsVX {
	if f8NmcatRx.flag != 0 && f8NmcatRx.flag&flagMethod == 0 {
		return (* /*line :1*/ rtype)( /*line :1*/ m6KMjaQj( /*line :1*/ unsafe.Pointer(f8NmcatRx.typ_)))
	}
	return  /*line :1*/ f8NmcatRx.qvaedZPPTM()
}

func (f8NmcatRx Value) qvaedZPPTM() YJh989LTZsVX {
	if f8NmcatRx.flag == 0 {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Type", Invalid})
	}

	lm5DH3 :=  /*line :1*/ f8NmcatRx.lm5DH3()
	if f8NmcatRx.flag&flagMethod == 0 {
		return  /*line :1*/ j5Z_WRWLudek( /*line :1*/ f8NmcatRx.lm5DH3())
	}

	hXZpj1nTZ :=  /*line :1*/ int(f8NmcatRx.flag) >> flagMethodShift
	if  /*line :1*/ f8NmcatRx.lm5DH3().Kind() == abi.Interface {

		mP7QqSVInV := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(lm5DH3))
		if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint( /*line :1*/ len(mP7QqSVInV.Methods)) {
			 /*line :1*/ panic("reflect: internal error: invalid method index")
		}
		hek9DV := &mP7QqSVInV.Methods[hXZpj1nTZ]
		return  /*line :1*/ j5Z_WRWLudek( /*line :1*/ nsOKvPP(lm5DH3, hek9DV.Typ))
	}

	vlMabJFJks :=  /*line :1*/ lm5DH3.ExportedMethods()
	if  /*line :1*/ uint(hXZpj1nTZ) >=  /*line :1*/ uint( /*line :1*/ len(vlMabJFJks)) {
		 /*line :1*/ panic("reflect: internal error: invalid method index")
	}
	hek9DV := vlMabJFJks[hXZpj1nTZ]
	return  /*line :1*/ j5Z_WRWLudek( /*line :1*/ nsOKvPP(lm5DH3, hek9DV.Mtyp))
}

func (f8NmcatRx Value) CanUint() bool {
	switch  /*line :1*/ f8NmcatRx.zRPbQda4dT() {
	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		return true
	default:
		return false
	}
}

func (f8NmcatRx Value) Uint() uint64 {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	e3IyyaQSEj := f8NmcatRx.ptr
	switch lYOqUQga {
	case Uint:
		return  /*line :1*/ uint64(*(* /*line :1*/ uint)(e3IyyaQSEj))
	case Uint8:
		return  /*line :1*/ uint64(*(* /*line :1*/ uint8)(e3IyyaQSEj))
	case Uint16:
		return  /*line :1*/ uint64(*(* /*line :1*/ uint16)(e3IyyaQSEj))
	case Uint32:
		return  /*line :1*/ uint64(*(* /*line :1*/ uint32)(e3IyyaQSEj))
	case Uint64:
		return *(* /*line :1*/ uint64)(e3IyyaQSEj)
	case Uintptr:
		return  /*line :1*/ uint64(*(* /*line :1*/ uintptr)(e3IyyaQSEj))
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Uint",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

//go:nocheckptr

func (f8NmcatRx Value) UnsafeAddr() uintptr {
	if  /*line :1*/ f8NmcatRx.lm5DH3() == nil {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.UnsafeAddr", Invalid})
	}
	if f8NmcatRx.flag&flagAddr == 0 {
		 /*line :1*/ panic("reflect.Value.UnsafeAddr of unaddressable value")
	}

	 /*line :1*/ vkf0Okel7G(f8NmcatRx.ptr)
	return  /*line :1*/ uintptr(f8NmcatRx.ptr)
}

func (f8NmcatRx Value) UnsafePointer() unsafe.Pointer {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.zRPbQda4dT()
	switch lYOqUQga {
	case Pointer:
		if  /*line :1*/ f8NmcatRx.lm5DH3().PtrBytes == 0 {

			if ! /*line :1*/ dvvw1_IpyIwu(*(* /*line :1*/ uintptr)(f8NmcatRx.ptr)) {
				 /*line :1*/ panic("reflect: reflect.Value.UnsafePointer on an invalid notinheap pointer")
			}
			return *(* /*line :1*/ unsafe.Pointer)(f8NmcatRx.ptr)
		}
		fallthrough
	case Chan, Map, UnsafePointer:
		return  /*line :1*/ f8NmcatRx.t_VyPi()
	case Func:
		if f8NmcatRx.flag&flagMethod != 0 {

			jaLOurn4Fy :=  /*line :1*/ vTsZZIIlq4MX()
			return *(* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(&jaLOurn4Fy))
		}
		e3IyyaQSEj :=  /*line :1*/ f8NmcatRx.t_VyPi()

		if e3IyyaQSEj != nil {
			e3IyyaQSEj = *(* /*line :1*/ unsafe.Pointer)(e3IyyaQSEj)
		}
		return e3IyyaQSEj

	case Slice:
		return (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr).TanDRHVvuLAL
	}
	 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.UnsafePointer",  /*line :1*/ f8NmcatRx.zRPbQda4dT()})
}

type SkmwA6 struct {
	CXigOex	uintptr
	EqOEBtiY	int
}

type Ei2Iy0 struct {
	H7OP0hFU	uintptr
	MCzI2qhUA1	int
	A6DWhwcvtSKZ	int
}

func ony4ojC67Y(rJmzo9n string, hQav8nZq, z6VQxP YJh989LTZsVX) {
	if hQav8nZq != z6VQxP {
		 /*line :1*/ panic(rJmzo9n + ": " +  /*line :1*/ hQav8nZq.String() + " != " +  /*line :1*/ z6VQxP.String())
	}
}

func rfJB4Et(e3IyyaQSEj unsafe.Pointer, hXZpj1nTZ int, vt31zHk8x35q uintptr, cdghXjtt string) unsafe.Pointer {
	return  /*line :1*/ lxkP7oV7sm(e3IyyaQSEj,  /*line :1*/ uintptr(hXZpj1nTZ)*vt31zHk8x35q, "i < len")
}

func (f8NmcatRx Value) Grow(krtR1HwEan int) {
	 /*line :1*/ f8NmcatRx.d5PnN9n()
	 /*line :1*/ f8NmcatRx.zc9yxv(Slice)
	 /*line :1*/ f8NmcatRx.at_aeFMxD0(krtR1HwEan)
}

func (f8NmcatRx Value) at_aeFMxD0(krtR1HwEan int) {
	e3IyyaQSEj := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
	switch {
	case krtR1HwEan < 0:
		 /*line :1*/ panic("reflect.Value.Grow: negative len")
	case e3IyyaQSEj.AJ2N5B+krtR1HwEan < 0:
		 /*line :1*/ panic("reflect.Value.Grow: slice overflow")
	case e3IyyaQSEj.AJ2N5B+krtR1HwEan > e3IyyaQSEj.Rsh30O4HaDX:
		sw8_lJ :=  /*line :1*/ f8NmcatRx.lm5DH3().Elem()
		*e3IyyaQSEj =  /*line :1*/ j23aSPt(sw8_lJ, *e3IyyaQSEj, krtR1HwEan)
	}
}

func (f8NmcatRx Value) c_Da2HLej3(krtR1HwEan int) Value {
	 /*line :1*/ f8NmcatRx.rwcaztMf()
	 /*line :1*/ f8NmcatRx.zc9yxv(Slice)

	z4pDntaDIY9Q := *(* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
	kXTvhUI4Zgg5 := &z4pDntaDIY9Q
	f8NmcatRx.ptr =  /*line :1*/ unsafe.Pointer(kXTvhUI4Zgg5)
	f8NmcatRx.flag = flagIndir |  /*line :1*/ flag(Slice)

	 /*line :1*/ f8NmcatRx.at_aeFMxD0(krtR1HwEan)
	kXTvhUI4Zgg5.AJ2N5B += krtR1HwEan
	return f8NmcatRx
}

func (f8NmcatRx Value) Clear() {
	switch  /*line :1*/ f8NmcatRx.Kind() {
	case Slice:
		z4pDntaDIY9Q := *(* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
		pjngE6n4a := (* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer( /*line :1*/ f8NmcatRx.lm5DH3()))
		 /*line :1*/ ijszg8v7(pjngE6n4a.Elem, z4pDntaDIY9Q.TanDRHVvuLAL, z4pDntaDIY9Q.AJ2N5B)
	case Map:
		 /*line :1*/ vuwNVEmTRp( /*line :1*/ f8NmcatRx.lm5DH3(),  /*line :1*/ f8NmcatRx.t_VyPi())
	default:
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Value.Clear",  /*line :1*/ f8NmcatRx.Kind()})
	}
}

func JrTptcQhpl0_(kXTvhUI4Zgg5 Value, uiTlN8bdKm ...Value) Value {
	 /*line :1*/ kXTvhUI4Zgg5.zc9yxv(Slice)
	krtR1HwEan :=  /*line :1*/ kXTvhUI4Zgg5.Len()
	kXTvhUI4Zgg5 =  /*line :1*/ kXTvhUI4Zgg5.c_Da2HLej3( /*line :1*/ len(uiTlN8bdKm))
	for hXZpj1nTZ, f8NmcatRx := range uiTlN8bdKm {
		 /*line :1*/ kXTvhUI4Zgg5.Index(krtR1HwEan + hXZpj1nTZ).Set(f8NmcatRx)
	}
	return kXTvhUI4Zgg5
}

func ItkAMZI(kXTvhUI4Zgg5, sw8_lJ Value) Value {
	 /*line :1*/ kXTvhUI4Zgg5.zc9yxv(Slice)
	 /*line :1*/ sw8_lJ.zc9yxv(Slice)
	 /*line :1*/ ony4ojC67Y("reflect.AppendSlice",  /*line :1*/ kXTvhUI4Zgg5.Type().Elem(),  /*line :1*/ sw8_lJ.Type().Elem())
	tAdmMo6P5NjN :=  /*line :1*/ kXTvhUI4Zgg5.Len()
	aa55OE2MTf :=  /*line :1*/ sw8_lJ.Len()
	kXTvhUI4Zgg5 =  /*line :1*/ kXTvhUI4Zgg5.c_Da2HLej3(aa55OE2MTf)
	 /*line :1*/ SlS0O3aFD_( /*line :1*/ kXTvhUI4Zgg5.Slice(tAdmMo6P5NjN, tAdmMo6P5NjN+aa55OE2MTf), sw8_lJ)
	return kXTvhUI4Zgg5
}

func SlS0O3aFD_(w2dYwY0, emdbbyWr Value) int {
	dtoxzmh :=  /*line :1*/ w2dYwY0.zRPbQda4dT()
	if dtoxzmh != Array && dtoxzmh != Slice {
		 /*line :1*/ panic(&R_QNJpzecp{"reflect.Copy", dtoxzmh})
	}
	if dtoxzmh == Array {
		 /*line :1*/ w2dYwY0.d5PnN9n()
	}
	 /*line :1*/ w2dYwY0.rwcaztMf()

	uJiywIxDZb :=  /*line :1*/ emdbbyWr.zRPbQda4dT()
	var saCTb4bezo bool
	if uJiywIxDZb != Array && uJiywIxDZb != Slice {
		saCTb4bezo = uJiywIxDZb == String &&  /*line :1*/ w2dYwY0.lm5DH3().Elem().Kind() == abi.Uint8
		if !saCTb4bezo {
			 /*line :1*/ panic(&R_QNJpzecp{"reflect.Copy", uJiywIxDZb})
		}
	}
	 /*line :1*/ emdbbyWr.rwcaztMf()

	aE1RA05i :=  /*line :1*/ w2dYwY0.lm5DH3().Elem()
	if !saCTb4bezo {
		kDT0Ji :=  /*line :1*/ emdbbyWr.lm5DH3().Elem()
		 /*line :1*/ ony4ojC67Y("reflect.Copy",  /*line :1*/ nadfj4Ka8cQf(aE1RA05i),  /*line :1*/ nadfj4Ka8cQf(kDT0Ji))
	}

	var gXy67My, sr4mKYvFuU9 unsafeheader.AKoOflOZ
	if dtoxzmh == Array {
		gXy67My.TanDRHVvuLAL = w2dYwY0.ptr
		gXy67My.AJ2N5B =  /*line :1*/ w2dYwY0.Len()
		gXy67My.Rsh30O4HaDX = gXy67My.AJ2N5B
	} else {
		gXy67My = *(* /*line :1*/ unsafeheader.AKoOflOZ)(w2dYwY0.ptr)
	}
	if uJiywIxDZb == Array {
		sr4mKYvFuU9.TanDRHVvuLAL = emdbbyWr.ptr
		sr4mKYvFuU9.AJ2N5B =  /*line :1*/ emdbbyWr.Len()
		sr4mKYvFuU9.Rsh30O4HaDX = sr4mKYvFuU9.AJ2N5B
	} else if uJiywIxDZb == Slice {
		sr4mKYvFuU9 = *(* /*line :1*/ unsafeheader.AKoOflOZ)(emdbbyWr.ptr)
	} else {
		z4pDntaDIY9Q := *(* /*line :1*/ unsafeheader.HrYeZlWHaf)(emdbbyWr.ptr)
		sr4mKYvFuU9.TanDRHVvuLAL = z4pDntaDIY9Q.QOvpyywy
		sr4mKYvFuU9.AJ2N5B = z4pDntaDIY9Q.SkAFPemec
		sr4mKYvFuU9.Rsh30O4HaDX = z4pDntaDIY9Q.SkAFPemec
	}

	return  /*line :1*/ raZwrTstH( /*line :1*/ aE1RA05i.Common(), gXy67My, sr4mKYvFuU9)
}

type bEera85TdK struct {
	ohBQVa7	SelectDir
	be4rYLILFd	*rtype
	rZwTK0DnETg	unsafe.Pointer
	fyNQ2j	unsafe.Pointer
}

//go:noescape
func b10lqWHete([]bEera85TdK) (ltSOS8D8nnvF int, hXwN33VNPZe bool)

type SelectDir int

const (
	_	SelectDir	= iota
	SelectSend
	SelectRecv
	SelectDefault
)

type SelectCase struct {
	Dir	SelectDir
	Chan	Value
	Send	Value
}

func We5qUeGn2f(kKX8L69yrdD []SelectCase) (ltSOS8D8nnvF int, wM9XCjNlp Value, hXwN33VNPZe bool) {
	if  /*line :1*/ len(kKX8L69yrdD) > 65536 {
		 /*line :1*/ panic("reflect.Select: too many cases (max 65536)")
	}

	var janThRCKpI []bEera85TdK
	if  /*line :1*/ len(kKX8L69yrdD) > 4 {

		janThRCKpI =  /*line :1*/ make([]bEera85TdK,  /*line :1*/ len(kKX8L69yrdD))
	} else {

		janThRCKpI =  /*line :1*/ make([]bEera85TdK,  /*line :1*/ len(kKX8L69yrdD), 4)
	}

	aL_mWBftP8Bq := false
	for hXZpj1nTZ, i5hAnuqALA := range kKX8L69yrdD {
		ma_YLG := &janThRCKpI[hXZpj1nTZ]
		ma_YLG.ohBQVa7 = i5hAnuqALA.Dir
		switch i5hAnuqALA.Dir {
		default:
			 /*line :1*/ panic("reflect.Select: invalid Dir")

		case SelectDefault:
			if aL_mWBftP8Bq {
				 /*line :1*/ panic("reflect.Select: multiple default cases")
			}
			aL_mWBftP8Bq = true
			if  /*line :1*/ i5hAnuqALA.Chan.IsValid() {
				 /*line :1*/ panic("reflect.Select: default case has Chan value")
			}
			if  /*line :1*/ i5hAnuqALA.Send.IsValid() {
				 /*line :1*/ panic("reflect.Select: default case has Send value")
			}

		case SelectSend:
			qqxx1RLlUhKF := i5hAnuqALA.Chan
			if ! /*line :1*/ qqxx1RLlUhKF.IsValid() {
				break
			}
			 /*line :1*/ qqxx1RLlUhKF.zc9yxv(Chan)
			 /*line :1*/ qqxx1RLlUhKF.rwcaztMf()
			mP7QqSVInV := (* /*line :1*/ aK0ovNmjQBD)( /*line :1*/ unsafe.Pointer( /*line :1*/ qqxx1RLlUhKF.lm5DH3()))
			if  /*line :1*/ Svj3YDOP3fa(mP7QqSVInV.Dir)&SendDir == 0 {
				 /*line :1*/ panic("reflect.Select: SendDir case using recv-only channel")
			}
			ma_YLG.rZwTK0DnETg =  /*line :1*/ qqxx1RLlUhKF.t_VyPi()
			ma_YLG.be4rYLILFd =  /*line :1*/ j5Z_WRWLudek(&mP7QqSVInV.Type)
			f8NmcatRx := i5hAnuqALA.Send
			if ! /*line :1*/ f8NmcatRx.IsValid() {
				 /*line :1*/ panic("reflect.Select: SendDir case missing Send value")
			}
			 /*line :1*/ f8NmcatRx.rwcaztMf()
			f8NmcatRx =  /*line :1*/ f8NmcatRx.hfrHMyTs1aE("reflect.Select", mP7QqSVInV.Elem, nil)
			if f8NmcatRx.flag&flagIndir != 0 {
				ma_YLG.fyNQ2j = f8NmcatRx.ptr
			} else {
				ma_YLG.fyNQ2j =  /*line :1*/ unsafe.Pointer(&f8NmcatRx.ptr)
			}

			 /*line :1*/ vkf0Okel7G(ma_YLG.fyNQ2j)

		case SelectRecv:
			if  /*line :1*/ i5hAnuqALA.Send.IsValid() {
				 /*line :1*/ panic("reflect.Select: RecvDir case has Send value")
			}
			qqxx1RLlUhKF := i5hAnuqALA.Chan
			if ! /*line :1*/ qqxx1RLlUhKF.IsValid() {
				break
			}
			 /*line :1*/ qqxx1RLlUhKF.zc9yxv(Chan)
			 /*line :1*/ qqxx1RLlUhKF.rwcaztMf()
			mP7QqSVInV := (* /*line :1*/ aK0ovNmjQBD)( /*line :1*/ unsafe.Pointer( /*line :1*/ qqxx1RLlUhKF.lm5DH3()))
			if  /*line :1*/ Svj3YDOP3fa(mP7QqSVInV.Dir)&RecvDir == 0 {
				 /*line :1*/ panic("reflect.Select: RecvDir case using send-only channel")
			}
			ma_YLG.rZwTK0DnETg =  /*line :1*/ qqxx1RLlUhKF.t_VyPi()
			ma_YLG.be4rYLILFd =  /*line :1*/ j5Z_WRWLudek(&mP7QqSVInV.Type)
			ma_YLG.fyNQ2j =  /*line :1*/ qscv_5(mP7QqSVInV.Elem)
		}
	}

	ltSOS8D8nnvF, hXwN33VNPZe =  /*line :1*/ b10lqWHete(janThRCKpI)
	if janThRCKpI[ltSOS8D8nnvF].ohBQVa7 == SelectRecv {
		mP7QqSVInV := (* /*line :1*/ aK0ovNmjQBD)( /*line :1*/ unsafe.Pointer(janThRCKpI[ltSOS8D8nnvF].be4rYLILFd))
		sw8_lJ := mP7QqSVInV.Elem
		e3IyyaQSEj := janThRCKpI[ltSOS8D8nnvF].fyNQ2j
		iiZKy9PE3 :=  /*line :1*/ flag( /*line :1*/ sw8_lJ.Kind())
		if  /*line :1*/ sw8_lJ.IfaceIndir() {
			wM9XCjNlp = Value{sw8_lJ, e3IyyaQSEj, iiZKy9PE3 | flagIndir}
		} else {
			wM9XCjNlp = Value{sw8_lJ, *(* /*line :1*/ unsafe.Pointer)(e3IyyaQSEj), iiZKy9PE3}
		}
	}
	return ltSOS8D8nnvF, wM9XCjNlp, hXwN33VNPZe
}

//go:noescape
func qscv_5(*abi.Type) unsafe.Pointer

//go:noescape
func l9WZQcaw(*abi.Type, int) unsafe.Pointer

func LydH53J9Si(lm5DH3 YJh989LTZsVX, gI0GRNfkNT4H, xUg7Px int) Value {
	if  /*line :1*/ lm5DH3.Kind() != Slice {
		 /*line :1*/ panic("reflect.MakeSlice of non-slice type")
	}
	if gI0GRNfkNT4H < 0 {
		 /*line :1*/ panic("reflect.MakeSlice: negative len")
	}
	if xUg7Px < 0 {
		 /*line :1*/ panic("reflect.MakeSlice: negative cap")
	}
	if gI0GRNfkNT4H > xUg7Px {
		 /*line :1*/ panic("reflect.MakeSlice: len > cap")
	}

	kXTvhUI4Zgg5 := unsafeheader.AKoOflOZ{TanDRHVvuLAL:  /*line :1*/ l9WZQcaw(&( /*line :1*/ lm5DH3.Elem().(*rtype).t), xUg7Px), AJ2N5B: gI0GRNfkNT4H, Rsh30O4HaDX: xUg7Px}
	return Value{&lm5DH3.(*rtype).t,  /*line :1*/ unsafe.Pointer(&kXTvhUI4Zgg5), flagIndir |  /*line :1*/ flag(Slice)}
}

func STTz99(lm5DH3 YJh989LTZsVX, ian_TcbM int) Value {
	if  /*line :1*/ lm5DH3.Kind() != Chan {
		 /*line :1*/ panic("reflect.MakeChan of non-chan type")
	}
	if ian_TcbM < 0 {
		 /*line :1*/ panic("reflect.MakeChan: negative buffer size")
	}
	if  /*line :1*/ lm5DH3.ChanDir() != BothDir {
		 /*line :1*/ panic("reflect.MakeChan: unidirectional channel type")
	}
	sw8_lJ :=  /*line :1*/ lm5DH3.z6hGxTABM1()
	qqxx1RLlUhKF :=  /*line :1*/ cMJfVAc6x(sw8_lJ, ian_TcbM)
	return Value{sw8_lJ, qqxx1RLlUhKF,  /*line :1*/ flag(Chan)}
}

func Fe3kdX79Jd(lm5DH3 YJh989LTZsVX) Value {
	return  /*line :1*/ ECnZcEYWRco(lm5DH3, 0)
}

func ECnZcEYWRco(lm5DH3 YJh989LTZsVX, krtR1HwEan int) Value {
	if  /*line :1*/ lm5DH3.Kind() != Map {
		 /*line :1*/ panic("reflect.MakeMapWithSize of non-map type")
	}
	sw8_lJ :=  /*line :1*/ lm5DH3.z6hGxTABM1()
	hek9DV :=  /*line :1*/ hxEiL_K1Bf(sw8_lJ, krtR1HwEan)
	return Value{sw8_lJ, hek9DV,  /*line :1*/ flag(Map)}
}

func XF6ikODQ8h(f8NmcatRx Value) Value {
	if  /*line :1*/ f8NmcatRx.Kind() != Pointer {
		return f8NmcatRx
	}
	return  /*line :1*/ f8NmcatRx.Elem()
}

const go121noForceValueEscape = true

func SdHoaIQl(hXZpj1nTZ any) Value {
	if hXZpj1nTZ == nil {
		return Value{}
	}

	if !go121noForceValueEscape {
		 /*line :1*/ vkf0Okel7G(hXZpj1nTZ)
	}

	return  /*line :1*/ i7YECW6sP91m(hXZpj1nTZ)
}

func H_B9rU1ADy(lm5DH3 YJh989LTZsVX) Value {
	if lm5DH3 == nil {
		 /*line :1*/ panic("reflect: Zero(nil)")
	}
	sw8_lJ := &lm5DH3.(*rtype).t
	iiZKy9PE3 :=  /*line :1*/ flag( /*line :1*/ sw8_lJ.Kind())
	if  /*line :1*/ sw8_lJ.IfaceIndir() {
		var e3IyyaQSEj unsafe.Pointer
		if  /*line :1*/ sw8_lJ.Size() <= maxZero {
			e3IyyaQSEj =  /*line :1*/ unsafe.Pointer(&mDOofGs5[0])
		} else {
			e3IyyaQSEj =  /*line :1*/ qscv_5(sw8_lJ)
		}
		return Value{sw8_lJ, e3IyyaQSEj, iiZKy9PE3 | flagIndir}
	}
	return Value{sw8_lJ, nil, iiZKy9PE3}
}

const maxZero = 1024

//go:linkname mDOofGs5 runtime.zeroVal
var mDOofGs5 [maxZero]byte

func L6EAuajfYe1(lm5DH3 YJh989LTZsVX) Value {
	if lm5DH3 == nil {
		 /*line :1*/ panic("reflect: New(nil)")
	}
	sw8_lJ := &lm5DH3.(*rtype).t
	xNIRhY :=  /*line :1*/ v8dZcSqC(sw8_lJ)
	if  /*line :1*/ gGK1Oc(xNIRhY) {

		 /*line :1*/ panic("reflect: New of type that may not be allocated in heap (possibly undefined cgo C type)")
	}
	cf3fCV8ayFq :=  /*line :1*/ qscv_5(sw8_lJ)
	iiZKy9PE3 :=  /*line :1*/ flag(Pointer)
	return Value{xNIRhY, cf3fCV8ayFq, iiZKy9PE3}
}

func MGIFN3e(lm5DH3 YJh989LTZsVX, e3IyyaQSEj unsafe.Pointer) Value {
	iiZKy9PE3 :=  /*line :1*/ flag(Pointer)
	sw8_lJ := lm5DH3.(*rtype)
	return Value{ /*line :1*/ sw8_lJ.v8dZcSqC(), e3IyyaQSEj, iiZKy9PE3}
}

func (f8NmcatRx Value) hfrHMyTs1aE(iCCtSqfz string, w2dYwY0 *abi.Type, htiV4jZ2U unsafe.Pointer) Value {
	if f8NmcatRx.flag&flagMethod != 0 {
		f8NmcatRx =  /*line :1*/ evDSV6R(iCCtSqfz, f8NmcatRx)
	}

	switch {
	case  /*line :1*/ vE0PYJd(w2dYwY0,  /*line :1*/ f8NmcatRx.lm5DH3()):

		iiZKy9PE3 := f8NmcatRx.flag&(flagAddr|flagIndir) |  /*line :1*/ f8NmcatRx.flag.bQ77Iz()
		iiZKy9PE3 |=  /*line :1*/ flag( /*line :1*/ w2dYwY0.Kind())
		return Value{w2dYwY0, f8NmcatRx.ptr, iiZKy9PE3}

	case  /*line :1*/ jih7qN7qX(w2dYwY0,  /*line :1*/ f8NmcatRx.lm5DH3()):
		if  /*line :1*/ f8NmcatRx.Kind() == Interface &&  /*line :1*/ f8NmcatRx.IsNil() {

			return Value{w2dYwY0, nil,  /*line :1*/ flag(Interface)}
		}
		uiTlN8bdKm :=  /*line :1*/ mySQReh3uV(f8NmcatRx, false)
		if htiV4jZ2U == nil {
			htiV4jZ2U =  /*line :1*/ qscv_5(w2dYwY0)
		}
		if  /*line :1*/ w2dYwY0.NumMethod() == 0 {
			*(* /*line :1*/ any)(htiV4jZ2U) = uiTlN8bdKm
		} else {
			 /*line :1*/ _TpI9kW0(w2dYwY0, uiTlN8bdKm, htiV4jZ2U)
		}
		return Value{w2dYwY0, htiV4jZ2U, flagIndir |  /*line :1*/ flag(Interface)}
	}

	 /*line :1*/ panic(iCCtSqfz + ": value of type " +  /*line :1*/ iR_1JLd( /*line :1*/ f8NmcatRx.lm5DH3()) + " is not assignable to type " +  /*line :1*/ iR_1JLd(w2dYwY0))
}

func (f8NmcatRx Value) Convert(sw8_lJ YJh989LTZsVX) Value {
	if f8NmcatRx.flag&flagMethod != 0 {
		f8NmcatRx =  /*line :1*/ evDSV6R("Convert", f8NmcatRx)
	}
	kZVCtGQ :=  /*line :1*/ kamZNfVk4( /*line :1*/ sw8_lJ.z6hGxTABM1(),  /*line :1*/ f8NmcatRx.lm5DH3())
	if kZVCtGQ == nil {
		 /*line :1*/ panic("reflect.Value.Convert: value of type " +  /*line :1*/ iR_1JLd( /*line :1*/ f8NmcatRx.lm5DH3()) + " cannot be converted to type " +  /*line :1*/ sw8_lJ.String())
	}
	return  /*line :1*/ kZVCtGQ(f8NmcatRx, sw8_lJ)
}

func (f8NmcatRx Value) CanConvert(sw8_lJ YJh989LTZsVX) bool {
	hDkOdBapcl :=  /*line :1*/ f8NmcatRx.Type()
	if ! /*line :1*/ hDkOdBapcl.ConvertibleTo(sw8_lJ) {
		return false
	}

	switch {
	case  /*line :1*/ hDkOdBapcl.Kind() == Slice &&  /*line :1*/ sw8_lJ.Kind() == Array:
		if  /*line :1*/ sw8_lJ.Len() >  /*line :1*/ f8NmcatRx.Len() {
			return false
		}
	case  /*line :1*/ hDkOdBapcl.Kind() == Slice &&  /*line :1*/ sw8_lJ.Kind() == Pointer &&  /*line :1*/ sw8_lJ.Elem().Kind() == Array:
		krtR1HwEan :=  /*line :1*/ sw8_lJ.Elem().Len()
		if krtR1HwEan >  /*line :1*/ f8NmcatRx.Len() {
			return false
		}
	}
	return true
}

func (f8NmcatRx Value) Comparable() bool {
	lYOqUQga :=  /*line :1*/ f8NmcatRx.Kind()
	switch lYOqUQga {
	case Invalid:
		return false

	case Array:
		switch  /*line :1*/ f8NmcatRx.Type().Elem().Kind() {
		case Interface, Array, Struct:
			for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ f8NmcatRx.Type().Len(); hXZpj1nTZ++ {
				if ! /*line :1*/ f8NmcatRx.Index(hXZpj1nTZ).Comparable() {
					return false
				}
			}
			return true
		}
		return  /*line :1*/ f8NmcatRx.Type().Comparable()

	case Interface:
		return  /*line :1*/ f8NmcatRx.Elem().Comparable()

	case Struct:
		for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ f8NmcatRx.NumField(); hXZpj1nTZ++ {
			if ! /*line :1*/ f8NmcatRx.Field(hXZpj1nTZ).Comparable() {
				return false
			}
		}
		return true

	default:
		return  /*line :1*/ f8NmcatRx.Type().Comparable()
	}
}

func (f8NmcatRx Value) Equal(iaoHrRg1 Value) bool {
	if  /*line :1*/ f8NmcatRx.Kind() == Interface {
		f8NmcatRx =  /*line :1*/ f8NmcatRx.Elem()
	}
	if  /*line :1*/ iaoHrRg1.Kind() == Interface {
		iaoHrRg1 =  /*line :1*/ iaoHrRg1.Elem()
	}

	if ! /*line :1*/ f8NmcatRx.IsValid() || ! /*line :1*/ iaoHrRg1.IsValid() {
		return  /*line :1*/ f8NmcatRx.IsValid() ==  /*line :1*/ iaoHrRg1.IsValid()
	}

	if  /*line :1*/ f8NmcatRx.Kind() !=  /*line :1*/ iaoHrRg1.Kind() ||  /*line :1*/ f8NmcatRx.Type() !=  /*line :1*/ iaoHrRg1.Type() {
		return false
	}

	switch  /*line :1*/ f8NmcatRx.Kind() {
	default:
		 /*line :1*/ panic("reflect.Value.Equal: invalid Kind")
	case Bool:
		return  /*line :1*/ f8NmcatRx.Bool() ==  /*line :1*/ iaoHrRg1.Bool()
	case Int, Int8, Int16, Int32, Int64:
		return  /*line :1*/ f8NmcatRx.Int() ==  /*line :1*/ iaoHrRg1.Int()
	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		return  /*line :1*/ f8NmcatRx.Uint() ==  /*line :1*/ iaoHrRg1.Uint()
	case Float32, Float64:
		return  /*line :1*/ f8NmcatRx.Float() ==  /*line :1*/ iaoHrRg1.Float()
	case Complex64, Complex128:
		return  /*line :1*/ f8NmcatRx.Complex() ==  /*line :1*/ iaoHrRg1.Complex()
	case String:
		return  /*line :1*/ f8NmcatRx.String() ==  /*line :1*/ iaoHrRg1.String()
	case Chan, Pointer, UnsafePointer:
		return  /*line :1*/ f8NmcatRx.Pointer() ==  /*line :1*/ iaoHrRg1.Pointer()
	case Array:

		swRRaJ4TX :=  /*line :1*/ f8NmcatRx.Len()
		if swRRaJ4TX == 0 {

			if ! /*line :1*/ f8NmcatRx.Type().Elem().Comparable() {
				break
			}
			return true
		}
		for hXZpj1nTZ := 0; hXZpj1nTZ < swRRaJ4TX; hXZpj1nTZ++ {
			if ! /*line :1*/ f8NmcatRx.Index(hXZpj1nTZ).Equal( /*line :1*/ iaoHrRg1.Index(hXZpj1nTZ)) {
				return false
			}
		}
		return true
	case Struct:

		xo8HeAZf8e :=  /*line :1*/ f8NmcatRx.NumField()
		for hXZpj1nTZ := 0; hXZpj1nTZ < xo8HeAZf8e; hXZpj1nTZ++ {
			if ! /*line :1*/ f8NmcatRx.Field(hXZpj1nTZ).Equal( /*line :1*/ iaoHrRg1.Field(hXZpj1nTZ)) {
				return false
			}
		}
		return true
	case Func, Map, Slice:
		break
	}
	 /*line :1*/ panic("reflect.Value.Equal: values of type " +  /*line :1*/ f8NmcatRx.Type().String() + " are not comparable")
}

func kamZNfVk4(w2dYwY0, emdbbyWr *abi.Type) func(Value, YJh989LTZsVX) Value {
	switch  /*line :1*/ WzYBjsQL0( /*line :1*/ emdbbyWr.Kind()) {
	case Int, Int8, Int16, Int32, Int64:
		switch  /*line :1*/ WzYBjsQL0( /*line :1*/ w2dYwY0.Kind()) {
		case Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
			return _AzZi2BQ4
		case Float32, Float64:
			return kJOiXDGL
		case String:
			return sVnUDMeTE
		}

	case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
		switch  /*line :1*/ WzYBjsQL0( /*line :1*/ w2dYwY0.Kind()) {
		case Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
			return dONeI3yzm_y
		case Float32, Float64:
			return c0gIW9G4
		case String:
			return km94poZZ
		}

	case Float32, Float64:
		switch  /*line :1*/ WzYBjsQL0( /*line :1*/ w2dYwY0.Kind()) {
		case Int, Int8, Int16, Int32, Int64:
			return iEVhtKVSA5
		case Uint, Uint8, Uint16, Uint32, Uint64, Uintptr:
			return hLaA2l9CzA1
		case Float32, Float64:
			return d_S_xqT
		}

	case Complex64, Complex128:
		switch  /*line :1*/ WzYBjsQL0( /*line :1*/ w2dYwY0.Kind()) {
		case Complex64, Complex128:
			return ew3IJq_auO
		}

	case String:
		if  /*line :1*/ w2dYwY0.Kind() == abi.Slice &&  /*line :1*/ hP5IxCk8Hfi( /*line :1*/ w2dYwY0.Elem()) == "" {
			switch  /*line :1*/ WzYBjsQL0( /*line :1*/ w2dYwY0.Elem().Kind()) {
			case Uint8:
				return pyaBZp
			case Int32:
				return uCYS2oSb
			}
		}

	case Slice:
		if  /*line :1*/ w2dYwY0.Kind() == abi.String &&  /*line :1*/ hP5IxCk8Hfi( /*line :1*/ emdbbyWr.Elem()) == "" {
			switch  /*line :1*/ WzYBjsQL0( /*line :1*/ emdbbyWr.Elem().Kind()) {
			case Uint8:
				return pdqUpnEtUxz
			case Int32:
				return d3lXxCYX
			}
		}

		if  /*line :1*/ w2dYwY0.Kind() == abi.Pointer &&  /*line :1*/ w2dYwY0.Elem().Kind() == abi.Array &&  /*line :1*/ emdbbyWr.Elem() ==  /*line :1*/ w2dYwY0.Elem().Elem() {
			return ucrc_2SB
		}

		if  /*line :1*/ w2dYwY0.Kind() == abi.Array &&  /*line :1*/ emdbbyWr.Elem() ==  /*line :1*/ w2dYwY0.Elem() {
			return uYtNd5F
		}

	case Chan:
		if  /*line :1*/ w2dYwY0.Kind() == abi.Chan &&  /*line :1*/ cY0jZX(w2dYwY0, emdbbyWr) {
			return prRSMia
		}
	}

	if  /*line :1*/ sr7apHgO(w2dYwY0, emdbbyWr, false) {
		return prRSMia
	}

	if  /*line :1*/ w2dYwY0.Kind() == abi.Pointer &&  /*line :1*/ jv91ZkF(w2dYwY0) == "" &&
		 /*line :1*/ emdbbyWr.Kind() == abi.Pointer &&  /*line :1*/ jv91ZkF(emdbbyWr) == "" &&
		 /*line :1*/ sr7apHgO( /*line :1*/ dPe0EQn(w2dYwY0),  /*line :1*/ dPe0EQn(emdbbyWr), false) {
		return prRSMia
	}

	if  /*line :1*/ jih7qN7qX(w2dYwY0, emdbbyWr) {
		if  /*line :1*/ emdbbyWr.Kind() == abi.Interface {
			return as2t8VFG1
		}
		return lVf01fyKFr
	}

	return nil
}

func wmmhWxqOg4Z(jToThzM flag, mHqzluwTelo uint64, sw8_lJ YJh989LTZsVX) Value {
	lm5DH3 :=  /*line :1*/ sw8_lJ.z6hGxTABM1()
	cf3fCV8ayFq :=  /*line :1*/ qscv_5(lm5DH3)
	switch  /*line :1*/ lm5DH3.Size() {
	case 1:
		*(* /*line :1*/ uint8)(cf3fCV8ayFq) =  /*line :1*/ uint8(mHqzluwTelo)
	case 2:
		*(* /*line :1*/ uint16)(cf3fCV8ayFq) =  /*line :1*/ uint16(mHqzluwTelo)
	case 4:
		*(* /*line :1*/ uint32)(cf3fCV8ayFq) =  /*line :1*/ uint32(mHqzluwTelo)
	case 8:
		*(* /*line :1*/ uint64)(cf3fCV8ayFq) = mHqzluwTelo
	}
	return Value{lm5DH3, cf3fCV8ayFq, jToThzM | flagIndir |  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())}
}

func _sHbcMn(jToThzM flag, f8NmcatRx float64, sw8_lJ YJh989LTZsVX) Value {
	lm5DH3 :=  /*line :1*/ sw8_lJ.z6hGxTABM1()
	cf3fCV8ayFq :=  /*line :1*/ qscv_5(lm5DH3)
	switch  /*line :1*/ lm5DH3.Size() {
	case 4:
		*(* /*line :1*/ float32)(cf3fCV8ayFq) =  /*line :1*/ float32(f8NmcatRx)
	case 8:
		*(* /*line :1*/ float64)(cf3fCV8ayFq) = f8NmcatRx
	}
	return Value{lm5DH3, cf3fCV8ayFq, jToThzM | flagIndir |  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())}
}

func eBwd98(jToThzM flag, f8NmcatRx float32, sw8_lJ YJh989LTZsVX) Value {
	lm5DH3 :=  /*line :1*/ sw8_lJ.z6hGxTABM1()
	cf3fCV8ayFq :=  /*line :1*/ qscv_5(lm5DH3)
	*(* /*line :1*/ float32)(cf3fCV8ayFq) = f8NmcatRx
	return Value{lm5DH3, cf3fCV8ayFq, jToThzM | flagIndir |  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())}
}

func ufNQZXyXyx(jToThzM flag, f8NmcatRx complex128, sw8_lJ YJh989LTZsVX) Value {
	lm5DH3 :=  /*line :1*/ sw8_lJ.z6hGxTABM1()
	cf3fCV8ayFq :=  /*line :1*/ qscv_5(lm5DH3)
	switch  /*line :1*/ lm5DH3.Size() {
	case 8:
		*(* /*line :1*/ complex64)(cf3fCV8ayFq) =  /*line :1*/ complex64(f8NmcatRx)
	case 16:
		*(* /*line :1*/ complex128)(cf3fCV8ayFq) = f8NmcatRx
	}
	return Value{lm5DH3, cf3fCV8ayFq, jToThzM | flagIndir |  /*line :1*/ flag( /*line :1*/ lm5DH3.Kind())}
}

func odUHkk9Ta9G7(jToThzM flag, f8NmcatRx string, sw8_lJ YJh989LTZsVX) Value {
	aLsdcpyP1 :=  /*line :1*/ L6EAuajfYe1(sw8_lJ).Elem()
	 /*line :1*/ aLsdcpyP1.SetString(f8NmcatRx)
	aLsdcpyP1.flag = aLsdcpyP1.flag&^flagAddr | jToThzM
	return aLsdcpyP1
}

func onbiNhpV(jToThzM flag, f8NmcatRx []byte, sw8_lJ YJh989LTZsVX) Value {
	aLsdcpyP1 :=  /*line :1*/ L6EAuajfYe1(sw8_lJ).Elem()
	 /*line :1*/ aLsdcpyP1.SetBytes(f8NmcatRx)
	aLsdcpyP1.flag = aLsdcpyP1.flag&^flagAddr | jToThzM
	return aLsdcpyP1
}

func zr1Fe9b1ACd(jToThzM flag, f8NmcatRx []rune, sw8_lJ YJh989LTZsVX) Value {
	aLsdcpyP1 :=  /*line :1*/ L6EAuajfYe1(sw8_lJ).Elem()
	 /*line :1*/ aLsdcpyP1.oqaG08DN(f8NmcatRx)
	aLsdcpyP1.flag = aLsdcpyP1.flag&^flagAddr | jToThzM
	return aLsdcpyP1
}

func _AzZi2BQ4(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ wmmhWxqOg4Z( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ uint64( /*line :1*/ f8NmcatRx.Int()), sw8_lJ)
}

func dONeI3yzm_y(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ wmmhWxqOg4Z( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ f8NmcatRx.Uint(), sw8_lJ)
}

func iEVhtKVSA5(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ wmmhWxqOg4Z( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ uint64( /*line :1*/ int64( /*line :1*/ f8NmcatRx.Float())), sw8_lJ)
}

func hLaA2l9CzA1(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ wmmhWxqOg4Z( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ uint64( /*line :1*/ f8NmcatRx.Float()), sw8_lJ)
}

func kJOiXDGL(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ _sHbcMn( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ float64( /*line :1*/ f8NmcatRx.Int()), sw8_lJ)
}

func c0gIW9G4(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ _sHbcMn( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ float64( /*line :1*/ f8NmcatRx.Uint()), sw8_lJ)
}

func d_S_xqT(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	if  /*line :1*/ f8NmcatRx.Type().Kind() == Float32 &&  /*line :1*/ sw8_lJ.Kind() == Float32 {

		return  /*line :1*/ eBwd98( /*line :1*/ f8NmcatRx.flag.bQ77Iz(), *(* /*line :1*/ float32)(f8NmcatRx.ptr), sw8_lJ)
	}
	return  /*line :1*/ _sHbcMn( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ f8NmcatRx.Float(), sw8_lJ)
}

func ew3IJq_auO(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ ufNQZXyXyx( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ f8NmcatRx.Complex(), sw8_lJ)
}

func sVnUDMeTE(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	kXTvhUI4Zgg5 := "\uFFFD"
	if uiTlN8bdKm :=  /*line :1*/ f8NmcatRx.Int();  /*line :1*/ int64( /*line :1*/ rune(uiTlN8bdKm)) == uiTlN8bdKm {
		kXTvhUI4Zgg5 =  /*line :1*/ string( /*line :1*/ rune(uiTlN8bdKm))
	}
	return  /*line :1*/ odUHkk9Ta9G7( /*line :1*/ f8NmcatRx.flag.bQ77Iz(), kXTvhUI4Zgg5, sw8_lJ)
}

func km94poZZ(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	kXTvhUI4Zgg5 := "\uFFFD"
	if uiTlN8bdKm :=  /*line :1*/ f8NmcatRx.Uint();  /*line :1*/ uint64( /*line :1*/ rune(uiTlN8bdKm)) == uiTlN8bdKm {
		kXTvhUI4Zgg5 =  /*line :1*/ string( /*line :1*/ rune(uiTlN8bdKm))
	}
	return  /*line :1*/ odUHkk9Ta9G7( /*line :1*/ f8NmcatRx.flag.bQ77Iz(), kXTvhUI4Zgg5, sw8_lJ)
}

func pdqUpnEtUxz(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ odUHkk9Ta9G7( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ string( /*line :1*/ f8NmcatRx.Bytes()), sw8_lJ)
}

func pyaBZp(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ onbiNhpV( /*line :1*/ f8NmcatRx.flag.bQ77Iz(), [] /*line :1*/ byte( /*line :1*/ f8NmcatRx.String()), sw8_lJ)
}

func d3lXxCYX(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ odUHkk9Ta9G7( /*line :1*/ f8NmcatRx.flag.bQ77Iz(),  /*line :1*/ string( /*line :1*/ f8NmcatRx.fjMiydk()), sw8_lJ)
}

func uCYS2oSb(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	return  /*line :1*/ zr1Fe9b1ACd( /*line :1*/ f8NmcatRx.flag.bQ77Iz(), [] /*line :1*/ rune( /*line :1*/ f8NmcatRx.String()), sw8_lJ)
}

func ucrc_2SB(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	krtR1HwEan :=  /*line :1*/ sw8_lJ.Elem().Len()
	if krtR1HwEan >  /*line :1*/ f8NmcatRx.Len() {
		 /*line :1*/ panic("reflect: cannot convert slice with length " +  /*line :1*/ itoa.V25ba9G5( /*line :1*/ f8NmcatRx.Len()) + " to pointer to array with length " +  /*line :1*/ itoa.V25ba9G5(krtR1HwEan))
	}
	cJS_wcOmyXDz := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
	return Value{ /*line :1*/ sw8_lJ.z6hGxTABM1(), cJS_wcOmyXDz.TanDRHVvuLAL, f8NmcatRx.flag&^(flagIndir|flagAddr|flagKindMask) |  /*line :1*/ flag(Pointer)}
}

func uYtNd5F(f8NmcatRx Value, sw8_lJ YJh989LTZsVX) Value {
	krtR1HwEan :=  /*line :1*/ sw8_lJ.Len()
	if krtR1HwEan >  /*line :1*/ f8NmcatRx.Len() {
		 /*line :1*/ panic("reflect: cannot convert slice with length " +  /*line :1*/ itoa.V25ba9G5( /*line :1*/ f8NmcatRx.Len()) + " to array with length " +  /*line :1*/ itoa.V25ba9G5(krtR1HwEan))
	}
	cJS_wcOmyXDz := (* /*line :1*/ unsafeheader.AKoOflOZ)(f8NmcatRx.ptr)
	lm5DH3 :=  /*line :1*/ sw8_lJ.z6hGxTABM1()
	cf3fCV8ayFq := cJS_wcOmyXDz.TanDRHVvuLAL
	i5hAnuqALA :=  /*line :1*/ qscv_5(lm5DH3)
	 /*line :1*/ l26oiKGd(lm5DH3, i5hAnuqALA, cf3fCV8ayFq)
	cf3fCV8ayFq = i5hAnuqALA

	return Value{lm5DH3, cf3fCV8ayFq, f8NmcatRx.flag&^(flagAddr|flagKindMask) |  /*line :1*/ flag(Array)}
}

func prRSMia(f8NmcatRx Value, lm5DH3 YJh989LTZsVX) Value {
	jToThzM := f8NmcatRx.flag
	sw8_lJ :=  /*line :1*/ lm5DH3.z6hGxTABM1()
	cf3fCV8ayFq := f8NmcatRx.ptr
	if jToThzM&flagAddr != 0 {

		i5hAnuqALA :=  /*line :1*/ qscv_5(sw8_lJ)
		 /*line :1*/ l26oiKGd(sw8_lJ, i5hAnuqALA, cf3fCV8ayFq)
		cf3fCV8ayFq = i5hAnuqALA
		jToThzM &^= flagAddr
	}
	return Value{sw8_lJ, cf3fCV8ayFq,  /*line :1*/ f8NmcatRx.flag.bQ77Iz() | jToThzM}
}

func lVf01fyKFr(f8NmcatRx Value, lm5DH3 YJh989LTZsVX) Value {
	htiV4jZ2U :=  /*line :1*/ qscv_5( /*line :1*/ lm5DH3.z6hGxTABM1())
	uiTlN8bdKm :=  /*line :1*/ mySQReh3uV(f8NmcatRx, false)
	if  /*line :1*/ lm5DH3.NumMethod() == 0 {
		*(* /*line :1*/ any)(htiV4jZ2U) = uiTlN8bdKm
	} else {
		 /*line :1*/ _TpI9kW0( /*line :1*/ lm5DH3.z6hGxTABM1(), uiTlN8bdKm, htiV4jZ2U)
	}
	return Value{ /*line :1*/ lm5DH3.z6hGxTABM1(), htiV4jZ2U,  /*line :1*/ f8NmcatRx.flag.bQ77Iz() | flagIndir |  /*line :1*/ flag(Interface)}
}

func as2t8VFG1(f8NmcatRx Value, lm5DH3 YJh989LTZsVX) Value {
	if  /*line :1*/ f8NmcatRx.IsNil() {
		aLsdcpyP1 :=  /*line :1*/ H_B9rU1ADy(lm5DH3)
		aLsdcpyP1.flag |=  /*line :1*/ f8NmcatRx.flag.bQ77Iz()
		return aLsdcpyP1
	}
	return  /*line :1*/ lVf01fyKFr( /*line :1*/ f8NmcatRx.Elem(), lm5DH3)
}

//go:noescape
func qUeiwmjb(qqxx1RLlUhKF unsafe.Pointer) int

//go:noescape
func yaFXieDVoP(qqxx1RLlUhKF unsafe.Pointer)

//go:noescape
func joS0Abp(qqxx1RLlUhKF unsafe.Pointer) int

//go:noescape
func aXQaEP2(qqxx1RLlUhKF unsafe.Pointer, j2eisdvd_at bool, w7xi3Bp unsafe.Pointer) (a_LOCfaJEYY1, gGhXsK bool)

//go:noescape
func h4YVZtJ0(qqxx1RLlUhKF unsafe.Pointer, w7xi3Bp unsafe.Pointer, j2eisdvd_at bool) bool

func nSsxqozFB(qqxx1RLlUhKF unsafe.Pointer, w7xi3Bp unsafe.Pointer, j2eisdvd_at bool) bool {
	 /*line :1*/ i6JoCnk57U(w7xi3Bp)
	return  /*line :1*/ h4YVZtJ0(qqxx1RLlUhKF, w7xi3Bp, j2eisdvd_at)
}

func cMJfVAc6x(lm5DH3 *abi.Type, wnTnByH int) (qqxx1RLlUhKF unsafe.Pointer)
func hxEiL_K1Bf(sw8_lJ *abi.Type, xUg7Px int) (hek9DV unsafe.Pointer)

//go:noescape
func unynMPYi(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN unsafe.Pointer) (w7xi3Bp unsafe.Pointer)

//go:noescape
func bjogv5(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN string) (w7xi3Bp unsafe.Pointer)

//go:noescape
func ysYaaSlj(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN, w7xi3Bp unsafe.Pointer)

func vJbWE1laDY(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN, w7xi3Bp unsafe.Pointer) {
	 /*line :1*/ i6JoCnk57U(ccj6oONwQVN)
	 /*line :1*/ i6JoCnk57U(w7xi3Bp)
	 /*line :1*/ ysYaaSlj(sw8_lJ, hek9DV, ccj6oONwQVN, w7xi3Bp)
}

//go:noescape
func lJ9j4Ry(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN string, w7xi3Bp unsafe.Pointer)

func i6bZUn(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN string, w7xi3Bp unsafe.Pointer) {
	 /*line :1*/ i6JoCnk57U((* /*line :1*/ unsafeheader.HrYeZlWHaf)( /*line :1*/ unsafe.Pointer(&ccj6oONwQVN)).QOvpyywy)
	 /*line :1*/ i6JoCnk57U(w7xi3Bp)
	 /*line :1*/ lJ9j4Ry(sw8_lJ, hek9DV, ccj6oONwQVN, w7xi3Bp)
}

//go:noescape
func a0tirBasjs(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN unsafe.Pointer)

//go:noescape
func um8hBA(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, ccj6oONwQVN string)

//go:noescape
func gRAVyD(sw8_lJ *abi.Type, hek9DV unsafe.Pointer, gPDgVVB *cLVKdW_vuouD)

//go:noescape
func cgg9ghkrEjoF(gPDgVVB *cLVKdW_vuouD) (ccj6oONwQVN unsafe.Pointer)

//go:noescape
func tW5jaa7(gPDgVVB *cLVKdW_vuouD) (dPe0EQn unsafe.Pointer)

//go:noescape
func i5Fv9sOiLwR(gPDgVVB *cLVKdW_vuouD)

//go:noescape
func n6Tt0Kac(hek9DV unsafe.Pointer) int

func vuwNVEmTRp(sw8_lJ *abi.Type, hek9DV unsafe.Pointer)

//go:noescape
//go:linkname d7AWjWcFH runtime.reflectcall
func d7AWjWcFH(jeT9kEnJt *abi.Type, jToThzM, mkehBv9 unsafe.Pointer, jfQJRc2TNG, rIWHHy, ayUQnG0gkSPC uint32, sfPBRit *abi.RegArgs)

func _TpI9kW0(sw8_lJ *abi.Type, emdbbyWr any, w2dYwY0 unsafe.Pointer)

//go:noescape
func pqAEZr6n(w2dYwY0, emdbbyWr unsafe.Pointer, wnTnByH uintptr)

//go:noescape
func l26oiKGd(sw8_lJ *abi.Type, w2dYwY0, emdbbyWr unsafe.Pointer)

//go:noescape
func q2gsmEF(sw8_lJ *abi.Type, cf3fCV8ayFq unsafe.Pointer)

//go:noescape
func u5gG9lm_(sw8_lJ *abi.Type, cf3fCV8ayFq unsafe.Pointer, e65m4G53plB, wnTnByH uintptr)

//go:noescape
func raZwrTstH(sw8_lJ *abi.Type, w2dYwY0, emdbbyWr unsafeheader.AKoOflOZ) int

//go:noescape
func ijszg8v7(oTa846jJ3kuJ *abi.Type, cf3fCV8ayFq unsafe.Pointer, gI0GRNfkNT4H int)

//go:noescape
func hmksVe18ic(sw8_lJ *abi.Type, e3IyyaQSEj unsafe.Pointer, cJS_wcOmyXDz uintptr) uintptr

func dvvw1_IpyIwu(e3IyyaQSEj uintptr) bool

//go:noescape
func j23aSPt(sw8_lJ *abi.Type, abs0ool unsafeheader.AKoOflOZ, smEbRO1UZkvL int) unsafeheader.AKoOflOZ

func vkf0Okel7G(uiTlN8bdKm any) {
	if eljqQaZxLWA.y56CrSJqBkV {
		eljqQaZxLWA.j6wtjqZor = uiTlN8bdKm
	}
}

var eljqQaZxLWA struct {
	y56CrSJqBkV	bool
	j6wtjqZor	any
}

func i6JoCnk57U(uiTlN8bdKm unsafe.Pointer) {
	if eljqQaZxLWA.y56CrSJqBkV {
		 /*line :1*/ vkf0Okel7G(*(* /*line :1*/ any)(uiTlN8bdKm))
	}
}

//go:nosplit
func m6KMjaQj(e3IyyaQSEj unsafe.Pointer) unsafe.Pointer {
	uiTlN8bdKm :=  /*line :1*/ uintptr(e3IyyaQSEj)
	return  /*line :1*/ unsafe.Pointer(uiTlN8bdKm ^ 0)
}
