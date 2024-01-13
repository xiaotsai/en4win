//line :1
package reflect

import (
	"internal/abi"
	"internal/goarch"
	"unsafe"
)



















var (
	afQUELHem	= abi.IntArgRegs
	fkCHk3QV5g	= abi.FloatArgRegs
	v_AYCQ3NT4	=  /*line :1*/ uintptr(abi.EffectiveFloatRegSize)
)




type jqNzYrsB struct {
	_BtGRMo2	eWZhCLQa

	
	
	v9du88xTkB	uintptr
	jLZh5ayKa9mx	uintptr	

	
	ciJuo5C	uintptr	
	pCQRHPfHe	int	
	kT87Jzba6Sl	int	
}


type eWZhCLQa int

const (
	abiStepBad	eWZhCLQa	= iota
	abiStepStack			
	abiStepIntReg			
	abiStepPointer			
	abiStepFloatReg			
)






type fUyYynmj4e struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	bbEexQdWR	[]jqNzYrsB
	la4rtfm0X_A	[]int

	rHUhKYMAh	uintptr	
	axGsqV, azUnhJ	int	
}

func (h6TmbCZ *fUyYynmj4e) a6LuuxPayDWk() {
	for hXZpj1nTZ, e3IyyaQSEj := range h6TmbCZ.bbEexQdWR {
		 /*line :1*/ println("part", hXZpj1nTZ, e3IyyaQSEj._BtGRMo2, e3IyyaQSEj.v9du88xTkB, e3IyyaQSEj.jLZh5ayKa9mx, e3IyyaQSEj.ciJuo5C, e3IyyaQSEj.pCQRHPfHe, e3IyyaQSEj.kT87Jzba6Sl)
	}
	 /*line :1*/ print("values ")
	for _, hXZpj1nTZ := range h6TmbCZ.la4rtfm0X_A {
		 /*line :1*/ print(hXZpj1nTZ, " ")
	}
	 /*line :1*/ println()
	 /*line :1*/ println("stack", h6TmbCZ.rHUhKYMAh)
	 /*line :1*/ println("iregs", h6TmbCZ.axGsqV)
	 /*line :1*/ println("fregs", h6TmbCZ.azUnhJ)
}




func (h6TmbCZ *fUyYynmj4e) xAplKg6Om(hXZpj1nTZ int) []jqNzYrsB {
	kXTvhUI4Zgg5 := h6TmbCZ.la4rtfm0X_A[hXZpj1nTZ]
	var dVDneyuGq int
	if hXZpj1nTZ ==  /*line :1*/ len(h6TmbCZ.la4rtfm0X_A)-1 {
		dVDneyuGq =  /*line :1*/ len(h6TmbCZ.bbEexQdWR)
	} else {
		dVDneyuGq = h6TmbCZ.la4rtfm0X_A[hXZpj1nTZ+1]
	}
	return h6TmbCZ.bbEexQdWR[kXTvhUI4Zgg5:dVDneyuGq]
}





func (h6TmbCZ *fUyYynmj4e) q9vGrK81jS(sw8_lJ *abi.Type) *jqNzYrsB {

	l30WE5tu :=  /*line :1*/ len(h6TmbCZ.bbEexQdWR)
	h6TmbCZ.la4rtfm0X_A =  /*line :1*/ append(h6TmbCZ.la4rtfm0X_A, l30WE5tu)
	if  /*line :1*/ sw8_lJ.Size() == 0 {

		h6TmbCZ.rHUhKYMAh =  /*line :1*/ rlY5YEji1B4(h6TmbCZ.rHUhKYMAh,  /*line :1*/ uintptr( /*line :1*/ sw8_lJ.Align()))
		return nil
	}

	kBgwIM := *h6TmbCZ
	if ! /*line :1*/ h6TmbCZ.aS_PETYuGVZ(sw8_lJ, 0) {

		*h6TmbCZ = kBgwIM
		 /*line :1*/ h6TmbCZ.iWMOckuq_i( /*line :1*/ sw8_lJ.Size(),  /*line :1*/ uintptr( /*line :1*/ sw8_lJ.Align()))
		return &h6TmbCZ.bbEexQdWR[ /*line :1*/ len(h6TmbCZ.bbEexQdWR)-1]
	}
	return nil
}







func (h6TmbCZ *fUyYynmj4e) xwidleMRyJ(iGkAMNgmOX *abi.Type) (*jqNzYrsB, bool) {

	h6TmbCZ.la4rtfm0X_A =  /*line :1*/ append(h6TmbCZ.la4rtfm0X_A,  /*line :1*/ len(h6TmbCZ.bbEexQdWR))
	var bDfzXlCm5Raf, cf3fCV8ayFq bool
	if  /*line :1*/ gGK1Oc(iGkAMNgmOX) ||  /*line :1*/ iGkAMNgmOX.Pointers() {
		bDfzXlCm5Raf =  /*line :1*/ h6TmbCZ.n2hARPBRj(0, goarch.PtrSize, 1, 0b1)
		cf3fCV8ayFq = true
	} else {

		bDfzXlCm5Raf =  /*line :1*/ h6TmbCZ.n2hARPBRj(0, goarch.PtrSize, 1, 0b0)
		cf3fCV8ayFq = false
	}
	if !bDfzXlCm5Raf {
		 /*line :1*/ h6TmbCZ.iWMOckuq_i(goarch.PtrSize, goarch.PtrSize)
		return &h6TmbCZ.bbEexQdWR[ /*line :1*/ len(h6TmbCZ.bbEexQdWR)-1], cf3fCV8ayFq
	}
	return nil, cf3fCV8ayFq
}










func (h6TmbCZ *fUyYynmj4e) aS_PETYuGVZ(sw8_lJ *abi.Type, dcDpXllPZ3e uintptr) bool {
	switch  /*line :1*/ WzYBjsQL0( /*line :1*/ sw8_lJ.Kind()) {
	case UnsafePointer, Pointer, Chan, Map, Func:
		return  /*line :1*/ h6TmbCZ.n2hARPBRj(dcDpXllPZ3e,  /*line :1*/ sw8_lJ.Size(), 1, 0b1)
	case Bool, Int, Uint, Int8, Uint8, Int16, Uint16, Int32, Uint32, Uintptr:
		return  /*line :1*/ h6TmbCZ.n2hARPBRj(dcDpXllPZ3e,  /*line :1*/ sw8_lJ.Size(), 1, 0b0)
	case Int64, Uint64:
		switch goarch.PtrSize {
		case 4:
			return  /*line :1*/ h6TmbCZ.n2hARPBRj(dcDpXllPZ3e, 4, 2, 0b0)
		case 8:
			return  /*line :1*/ h6TmbCZ.n2hARPBRj(dcDpXllPZ3e, 8, 1, 0b0)
		}
	case Float32, Float64:
		return  /*line :1*/ h6TmbCZ.oTnMSRwN2o(dcDpXllPZ3e,  /*line :1*/ sw8_lJ.Size(), 1)
	case Complex64:
		return  /*line :1*/ h6TmbCZ.oTnMSRwN2o(dcDpXllPZ3e, 4, 2)
	case Complex128:
		return  /*line :1*/ h6TmbCZ.oTnMSRwN2o(dcDpXllPZ3e, 8, 2)
	case String:
		return  /*line :1*/ h6TmbCZ.n2hARPBRj(dcDpXllPZ3e, goarch.PtrSize, 2, 0b01)
	case Interface:
		return  /*line :1*/ h6TmbCZ.n2hARPBRj(dcDpXllPZ3e, goarch.PtrSize, 2, 0b10)
	case Slice:
		return  /*line :1*/ h6TmbCZ.n2hARPBRj(dcDpXllPZ3e, goarch.PtrSize, 3, 0b001)
	case Array:
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		switch mP7QqSVInV.Len {
		case 0:

			return true
		case 1:
			return  /*line :1*/ h6TmbCZ.aS_PETYuGVZ(mP7QqSVInV.Elem, dcDpXllPZ3e)
		default:
			return false
		}
	case Struct:
		pjngE6n4a := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		for hXZpj1nTZ := range pjngE6n4a.Fields {
			jToThzM := &pjngE6n4a.Fields[hXZpj1nTZ]
			if ! /*line :1*/ h6TmbCZ.aS_PETYuGVZ(jToThzM.Typ, dcDpXllPZ3e+jToThzM.Offset) {
				return false
			}
		}
		return true
	default:
		 /*line :1*/ print("t.Kind == ",  /*line :1*/ sw8_lJ.Kind(), "\n")
		 /*line :1*/ panic("unknown type kind")
	}
	 /*line :1*/ panic("unhandled register assignment path")
}










func (h6TmbCZ *fUyYynmj4e) n2hARPBRj(dcDpXllPZ3e, wnTnByH uintptr, krtR1HwEan int, vGn3oAB uint8) bool {
	if krtR1HwEan > 8 || krtR1HwEan < 0 {
		 /*line :1*/ panic("invalid n")
	}
	if vGn3oAB != 0 && wnTnByH != goarch.PtrSize {
		 /*line :1*/ panic("non-empty pointer map passed for non-pointer-size values")
	}
	if h6TmbCZ.axGsqV+krtR1HwEan > afQUELHem {
		return false
	}
	for hXZpj1nTZ := 0; hXZpj1nTZ < krtR1HwEan; hXZpj1nTZ++ {
		zRPbQda4dT := abiStepIntReg
		if vGn3oAB&( /*line :1*/ uint8(1)<<hXZpj1nTZ) != 0 {
			zRPbQda4dT = abiStepPointer
		}
		h6TmbCZ.bbEexQdWR =  /*line :1*/ append(h6TmbCZ.bbEexQdWR, jqNzYrsB{
			_BtGRMo2:	zRPbQda4dT,
			v9du88xTkB:	dcDpXllPZ3e +  /*line :1*/ uintptr(hXZpj1nTZ)*wnTnByH,
			jLZh5ayKa9mx:	wnTnByH,
			pCQRHPfHe:	h6TmbCZ.axGsqV,
		})
		h6TmbCZ.axGsqV++
	}
	return true
}







func (h6TmbCZ *fUyYynmj4e) oTnMSRwN2o(dcDpXllPZ3e, wnTnByH uintptr, krtR1HwEan int) bool {
	if krtR1HwEan < 0 {
		 /*line :1*/ panic("invalid n")
	}
	if h6TmbCZ.azUnhJ+krtR1HwEan > fkCHk3QV5g || v_AYCQ3NT4 < wnTnByH {
		return false
	}
	for hXZpj1nTZ := 0; hXZpj1nTZ < krtR1HwEan; hXZpj1nTZ++ {
		h6TmbCZ.bbEexQdWR =  /*line :1*/ append(h6TmbCZ.bbEexQdWR, jqNzYrsB{
			_BtGRMo2:	abiStepFloatReg,
			v9du88xTkB:	dcDpXllPZ3e +  /*line :1*/ uintptr(hXZpj1nTZ)*wnTnByH,
			jLZh5ayKa9mx:	wnTnByH,
			kT87Jzba6Sl:	h6TmbCZ.azUnhJ,
		})
		h6TmbCZ.azUnhJ++
	}
	return true
}





func (h6TmbCZ *fUyYynmj4e) iWMOckuq_i(wnTnByH, sExT8uzOyaD uintptr) {
	h6TmbCZ.rHUhKYMAh =  /*line :1*/ rlY5YEji1B4(h6TmbCZ.rHUhKYMAh, sExT8uzOyaD)
	h6TmbCZ.bbEexQdWR =  /*line :1*/ append(h6TmbCZ.bbEexQdWR, jqNzYrsB{
		_BtGRMo2:	abiStepStack,
		v9du88xTkB:	0,
		jLZh5ayKa9mx:	wnTnByH,
		ciJuo5C:	h6TmbCZ.rHUhKYMAh,
	})
	h6TmbCZ.rHUhKYMAh += wnTnByH
}


type cxVCV8 struct {
	
	
	pO2s2NwMlS, gSH9DY2KdKea	fUyYynmj4e

	
	
	
	
	
	
	
	yUiJTje95S, eBcFPCE, hq8R8Dk	uintptr

	
	
	
	
	
	sCYaW3ZJnC	*uh88jBcDEay2

	
	
	
	
	
	
	
	
	xm5UM36Yljc, aHHpIUai1ke	abi.IntArgRegBitmap
}

func (h6TmbCZ *cxVCV8) a6LuuxPayDWk() {
	 /*line :1*/ println("ABI")
	 /*line :1*/ println("call")
	 /*line :1*/ h6TmbCZ.pO2s2NwMlS.a6LuuxPayDWk()
	 /*line :1*/ println("ret")
	 /*line :1*/ h6TmbCZ.gSH9DY2KdKea.a6LuuxPayDWk()
	 /*line :1*/ println("stackCallArgsSize", h6TmbCZ.yUiJTje95S)
	 /*line :1*/ println("retOffset", h6TmbCZ.eBcFPCE)
	 /*line :1*/ println("spill", h6TmbCZ.hq8R8Dk)
	 /*line :1*/ print("inRegPtrs:")
	 /*line :1*/ oyGqKWu6cN(h6TmbCZ.xm5UM36Yljc)
	 /*line :1*/ println()
	 /*line :1*/ print("outRegPtrs:")
	 /*line :1*/ oyGqKWu6cN(h6TmbCZ.aHHpIUai1ke)
	 /*line :1*/ println()
}

func oyGqKWu6cN(w1Ndp0ySZgL abi.IntArgRegBitmap) {
	for hXZpj1nTZ := 0; hXZpj1nTZ < afQUELHem; hXZpj1nTZ++ {
		uiTlN8bdKm := 0
		if  /*line :1*/ w1Ndp0ySZgL.Get(hXZpj1nTZ) {
			uiTlN8bdKm = 1
		}
		 /*line :1*/ print(" ", uiTlN8bdKm)
	}
}

func y6pUOfj6v(sw8_lJ *bY2Zo8pis, iGkAMNgmOX *abi.Type) cxVCV8 {

	sSY_YtPAn :=  /*line :1*/ uintptr(0)

	rqXadRRDtZ :=  /*line :1*/ new(uh88jBcDEay2)

	jtWhFa8 := abi.IntArgRegBitmap{}

	
	var gJ3XGL_F fUyYynmj4e
	if iGkAMNgmOX != nil {
		eogO8b, hj8SFI :=  /*line :1*/ gJ3XGL_F.xwidleMRyJ(iGkAMNgmOX)
		if eogO8b != nil {
			if hj8SFI {
				 /*line :1*/ rqXadRRDtZ.cn3HN7Iy(1)
			} else {
				 /*line :1*/ rqXadRRDtZ.cn3HN7Iy(0)
			}
		} else {
			sSY_YtPAn += goarch.PtrSize
		}
	}
	for hXZpj1nTZ, cUbsbg9j := range  /*line :1*/ sw8_lJ.InSlice() {
		eogO8b :=  /*line :1*/ gJ3XGL_F.q9vGrK81jS(cUbsbg9j)
		if eogO8b != nil {
			 /*line :1*/ hK9wu75JAa(rqXadRRDtZ, eogO8b.ciJuo5C, cUbsbg9j)
		} else {
			sSY_YtPAn =  /*line :1*/ rlY5YEji1B4(sSY_YtPAn,  /*line :1*/ uintptr( /*line :1*/ cUbsbg9j.Align()))
			sSY_YtPAn +=  /*line :1*/ cUbsbg9j.Size()
			for _, pjngE6n4a := range  /*line :1*/ gJ3XGL_F.xAplKg6Om(hXZpj1nTZ) {
				if pjngE6n4a._BtGRMo2 == abiStepPointer {
					 /*line :1*/ jtWhFa8.Set(pjngE6n4a.pCQRHPfHe)
				}
			}
		}
	}
	sSY_YtPAn =  /*line :1*/ rlY5YEji1B4(sSY_YtPAn, goarch.PtrSize)

	egJf0B := gJ3XGL_F.rHUhKYMAh
	aeACV2 :=  /*line :1*/ rlY5YEji1B4(gJ3XGL_F.rHUhKYMAh, goarch.PtrSize)

	ck0_wE633vYX := abi.IntArgRegBitmap{}

	
	var hEbpexD fUyYynmj4e

	hEbpexD.rHUhKYMAh = aeACV2
	for hXZpj1nTZ, vaEUmTC := range  /*line :1*/ sw8_lJ.OutSlice() {
		eogO8b :=  /*line :1*/ hEbpexD.q9vGrK81jS(vaEUmTC)
		if eogO8b != nil {
			 /*line :1*/ hK9wu75JAa(rqXadRRDtZ, eogO8b.ciJuo5C, vaEUmTC)
		} else {
			for _, pjngE6n4a := range  /*line :1*/ hEbpexD.xAplKg6Om(hXZpj1nTZ) {
				if pjngE6n4a._BtGRMo2 == abiStepPointer {
					 /*line :1*/ ck0_wE633vYX.Set(pjngE6n4a.pCQRHPfHe)
				}
			}
		}
	}

	hEbpexD.rHUhKYMAh -= aeACV2
	return cxVCV8{gJ3XGL_F, hEbpexD, egJf0B, aeACV2, sSY_YtPAn, rqXadRRDtZ, jtWhFa8, ck0_wE633vYX}
}




func hchALWv(hgw67Aaoc0 *abi.RegArgs, bQUhDPtwJPE int, pqOC6eJCq9i2 uintptr, m4wOIq unsafe.Pointer) {
	 /*line :1*/ pqAEZr6n(m4wOIq,  /*line :1*/ hgw67Aaoc0.IntRegArgAddr(bQUhDPtwJPE, pqOC6eJCq9i2), pqOC6eJCq9i2)
}




func bGUOvK(hgw67Aaoc0 *abi.RegArgs, bQUhDPtwJPE int, pqOC6eJCq9i2 uintptr, xFFtzJR unsafe.Pointer) {
	 /*line :1*/ pqAEZr6n( /*line :1*/ hgw67Aaoc0.IntRegArgAddr(bQUhDPtwJPE, pqOC6eJCq9i2), xFFtzJR, pqOC6eJCq9i2)
}




func kh2Bx1R9(hgw67Aaoc0 *abi.RegArgs, bQUhDPtwJPE int, pqOC6eJCq9i2 uintptr, m4wOIq unsafe.Pointer) {
	switch pqOC6eJCq9i2 {
	case 4:
		*(* /*line :1*/ float32)(m4wOIq) =  /*line :1*/ eXwmv5Y(hgw67Aaoc0.Floats[bQUhDPtwJPE])
	case 8:
		*(* /*line :1*/ float64)(m4wOIq) = *(* /*line :1*/ float64)( /*line :1*/ unsafe.Pointer(&hgw67Aaoc0.Floats[bQUhDPtwJPE]))
	default:
		 /*line :1*/ panic("bad argSize")
	}
}




func nhGam6sL_Y(hgw67Aaoc0 *abi.RegArgs, bQUhDPtwJPE int, pqOC6eJCq9i2 uintptr, xFFtzJR unsafe.Pointer) {
	switch pqOC6eJCq9i2 {
	case 4:
		hgw67Aaoc0.Floats[bQUhDPtwJPE] =  /*line :1*/ fa22gIVp(*(* /*line :1*/ float32)(xFFtzJR))
	case 8:
		hgw67Aaoc0.Floats[bQUhDPtwJPE] = *(* /*line :1*/ uint64)(xFFtzJR)
	default:
		 /*line :1*/ panic("bad argSize")
	}
}
