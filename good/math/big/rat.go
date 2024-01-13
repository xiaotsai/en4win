//line :1
package big

import (
	fmt "kFsoCfy5zWG"
	math "math"
)










type FGtP7gcPU struct {
	
	
	
	
	
	u67K4Sv, _WnhHVFvm7 ZMtGuo
}


func DRKQ7Saq1VNc(sza5atF, kWa1YrIHZo int64) *FGtP7gcPU {
	return  /*line :1*/ new(FGtP7gcPU).SetFrac64(sza5atF, kWa1YrIHZo)
}



func (uW0eIVjy *FGtP7gcPU) SetFloat64(sYV3C6v3w1 float64) *FGtP7gcPU {
	const expMask = 1<<11 - 1
	uopo38aVK_H :=  /*line :1*/ math.GKIRmoHE(sYV3C6v3w1)
	yDzCx8 := uopo38aVK_H & (1<<52 - 1)
	xXsI4WC :=  /*line :1*/ int((uopo38aVK_H >> 52) & expMask)
	switch xXsI4WC {
	case expMask:
		return nil
	case 0:
		xXsI4WC -= 1022
	default:
		yDzCx8 |= 1 << 52
		xXsI4WC -= 1023
	}

	qKn7aIMh5 := 52 - xXsI4WC

	for yDzCx8&1 == 0 && qKn7aIMh5 > 0 {
		yDzCx8 >>= 1
		qKn7aIMh5--
	}

	 /*line :1*/ uW0eIVjy.u67K4Sv.SetUint64(yDzCx8)
	uW0eIVjy.u67K4Sv.hCTOmp0gckSa = sYV3C6v3w1 < 0
	 /*line :1*/ uW0eIVjy._WnhHVFvm7.Set(h2Aj5SV9g2re)
	if qKn7aIMh5 > 0 {
		 /*line :1*/ uW0eIVjy._WnhHVFvm7.Lsh(&uW0eIVjy._WnhHVFvm7,  /*line :1*/ uint(qKn7aIMh5))
	} else {
		 /*line :1*/ uW0eIVjy.u67K4Sv.Lsh(&uW0eIVjy.u67K4Sv,  /*line :1*/ uint(-qKn7aIMh5))
	}
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}





func iBMPVCmFlN(sza5atF, kWa1YrIHZo g3X9fa) (sYV3C6v3w1 float32, xLAqKDEG bool) {
	const (
		
		Fsize	= 32

		
		Msize	= 23
		Msize1	= Msize + 1		
		Msize2	= Msize1 + 1

		
		Esize	= Fsize - Msize1
		Ebias	= 1<<(Esize-1) - 1
		Emin	= 1 - Ebias
		Emax	= Ebias
	)

	iFXajOnrt :=  /*line :1*/ sza5atF.wZwJ3Y0()
	if iFXajOnrt == 0 {
		return 0, true
	}
	gVTBLaUS_N2 :=  /*line :1*/ kWa1YrIHZo.wZwJ3Y0()
	if gVTBLaUS_N2 == 0 {
		 /*line :1*/ panic("division by zero")
	}

	xXsI4WC := iFXajOnrt - gVTBLaUS_N2
	var diOe0B, _axWTK8IPk6s g3X9fa
	diOe0B =  /*line :1*/ diOe0B.bj0nVC(sza5atF)
	_axWTK8IPk6s =  /*line :1*/ _axWTK8IPk6s.bj0nVC(kWa1YrIHZo)
	if qKn7aIMh5 := Msize2 - xXsI4WC; qKn7aIMh5 > 0 {
		diOe0B =  /*line :1*/ diOe0B.z839tk6CmDHT(diOe0B,  /*line :1*/ uint(qKn7aIMh5))
	} else if qKn7aIMh5 < 0 {
		_axWTK8IPk6s =  /*line :1*/ _axWTK8IPk6s.z839tk6CmDHT(_axWTK8IPk6s,  /*line :1*/ uint(-qKn7aIMh5))
	}

	
	
	
	var yL0_p0wUnc g3X9fa
	yL0_p0wUnc, vFx5p_cm :=  /*line :1*/ yL0_p0wUnc.xOZxoLyl(diOe0B, diOe0B, _axWTK8IPk6s)
	yDzCx8 :=  /*line :1*/ xsbmp8(yL0_p0wUnc)
	sWalVm :=  /*line :1*/ len(vFx5p_cm) > 0

	if yDzCx8>>Msize2 == 1 {
		if yDzCx8&1 == 1 {
			sWalVm = true
		}
		yDzCx8 >>= 1
		xXsI4WC++
	}
	if yDzCx8>>Msize1 != 1 {
		 /*line :1*/ panic( /*line :1*/ fmt.IBsS3Oc("expected exactly %d bits of result", Msize2))
	}

	if Emin-Msize <= xXsI4WC && xXsI4WC <= Emin {

		qKn7aIMh5 :=  /*line :1*/ uint(Emin - (xXsI4WC - 1))
		lNH8PS_TDc := yDzCx8 & (1<<qKn7aIMh5 - 1)
		sWalVm = sWalVm || lNH8PS_TDc != 0
		yDzCx8 >>= qKn7aIMh5
		xXsI4WC = 2 - Ebias
	}

	xLAqKDEG = !sWalVm
	if yDzCx8&1 != 0 {
		xLAqKDEG = false
		if sWalVm || yDzCx8&2 != 0 {
			if yDzCx8++; yDzCx8 >= 1<<Msize2 {

				yDzCx8 >>= 1
				xXsI4WC++
			}
		}
	}
	yDzCx8 >>= 1

	sYV3C6v3w1 =  /*line :1*/ float32( /*line :1*/ math.EU1SFL1cw( /*line :1*/ float64(yDzCx8), xXsI4WC-Msize1))
	if  /*line :1*/ math.ME1vzpD5wcr( /*line :1*/ float64(sYV3C6v3w1), 0) {
		xLAqKDEG = false
	}
	return
}





func mFMYokd8J(sza5atF, kWa1YrIHZo g3X9fa) (sYV3C6v3w1 float64, xLAqKDEG bool) {
	const (
		
		Fsize	= 64

		
		Msize	= 52
		Msize1	= Msize + 1		
		Msize2	= Msize1 + 1

		
		Esize	= Fsize - Msize1
		Ebias	= 1<<(Esize-1) - 1
		Emin	= 1 - Ebias
		Emax	= Ebias
	)

	iFXajOnrt :=  /*line :1*/ sza5atF.wZwJ3Y0()
	if iFXajOnrt == 0 {
		return 0, true
	}
	gVTBLaUS_N2 :=  /*line :1*/ kWa1YrIHZo.wZwJ3Y0()
	if gVTBLaUS_N2 == 0 {
		 /*line :1*/ panic("division by zero")
	}

	xXsI4WC := iFXajOnrt - gVTBLaUS_N2
	var diOe0B, _axWTK8IPk6s g3X9fa
	diOe0B =  /*line :1*/ diOe0B.bj0nVC(sza5atF)
	_axWTK8IPk6s =  /*line :1*/ _axWTK8IPk6s.bj0nVC(kWa1YrIHZo)
	if qKn7aIMh5 := Msize2 - xXsI4WC; qKn7aIMh5 > 0 {
		diOe0B =  /*line :1*/ diOe0B.z839tk6CmDHT(diOe0B,  /*line :1*/ uint(qKn7aIMh5))
	} else if qKn7aIMh5 < 0 {
		_axWTK8IPk6s =  /*line :1*/ _axWTK8IPk6s.z839tk6CmDHT(_axWTK8IPk6s,  /*line :1*/ uint(-qKn7aIMh5))
	}

	
	
	
	var yL0_p0wUnc g3X9fa
	yL0_p0wUnc, vFx5p_cm :=  /*line :1*/ yL0_p0wUnc.xOZxoLyl(diOe0B, diOe0B, _axWTK8IPk6s)
	yDzCx8 :=  /*line :1*/ zA7LNXzkU(yL0_p0wUnc)
	sWalVm :=  /*line :1*/ len(vFx5p_cm) > 0

	if yDzCx8>>Msize2 == 1 {
		if yDzCx8&1 == 1 {
			sWalVm = true
		}
		yDzCx8 >>= 1
		xXsI4WC++
	}
	if yDzCx8>>Msize1 != 1 {
		 /*line :1*/ panic( /*line :1*/ fmt.IBsS3Oc("expected exactly %d bits of result", Msize2))
	}

	if Emin-Msize <= xXsI4WC && xXsI4WC <= Emin {

		qKn7aIMh5 :=  /*line :1*/ uint(Emin - (xXsI4WC - 1))
		lNH8PS_TDc := yDzCx8 & (1<<qKn7aIMh5 - 1)
		sWalVm = sWalVm || lNH8PS_TDc != 0
		yDzCx8 >>= qKn7aIMh5
		xXsI4WC = 2 - Ebias
	}

	xLAqKDEG = !sWalVm
	if yDzCx8&1 != 0 {
		xLAqKDEG = false
		if sWalVm || yDzCx8&2 != 0 {
			if yDzCx8++; yDzCx8 >= 1<<Msize2 {

				yDzCx8 >>= 1
				xXsI4WC++
			}
		}
	}
	yDzCx8 >>= 1

	sYV3C6v3w1 =  /*line :1*/ math.EU1SFL1cw( /*line :1*/ float64(yDzCx8), xXsI4WC-Msize1)
	if  /*line :1*/ math.ME1vzpD5wcr(sYV3C6v3w1, 0) {
		xLAqKDEG = false
	}
	return
}





func (pmEgen2 *FGtP7gcPU) Float32() (sYV3C6v3w1 float32, xLAqKDEG bool) {
	kWa1YrIHZo := pmEgen2._WnhHVFvm7.qho77PBKF
	if  /*line :1*/ len(kWa1YrIHZo) == 0 {
		kWa1YrIHZo = lG0t2KfqGwyP
	}
	sYV3C6v3w1, xLAqKDEG =  /*line :1*/ iBMPVCmFlN(pmEgen2.u67K4Sv.qho77PBKF, kWa1YrIHZo)
	if pmEgen2.u67K4Sv.hCTOmp0gckSa {
		sYV3C6v3w1 = -sYV3C6v3w1
	}
	return
}





func (pmEgen2 *FGtP7gcPU) Float64() (sYV3C6v3w1 float64, xLAqKDEG bool) {
	kWa1YrIHZo := pmEgen2._WnhHVFvm7.qho77PBKF
	if  /*line :1*/ len(kWa1YrIHZo) == 0 {
		kWa1YrIHZo = lG0t2KfqGwyP
	}
	sYV3C6v3w1, xLAqKDEG =  /*line :1*/ mFMYokd8J(pmEgen2.u67K4Sv.qho77PBKF, kWa1YrIHZo)
	if pmEgen2.u67K4Sv.hCTOmp0gckSa {
		sYV3C6v3w1 = -sYV3C6v3w1
	}
	return
}



func (uW0eIVjy *FGtP7gcPU) SetFrac(sza5atF, kWa1YrIHZo *ZMtGuo) *FGtP7gcPU {
	uW0eIVjy.u67K4Sv.hCTOmp0gckSa = sza5atF.hCTOmp0gckSa != kWa1YrIHZo.hCTOmp0gckSa
	bw0lYdCKY := kWa1YrIHZo.qho77PBKF
	if  /*line :1*/ len(bw0lYdCKY) == 0 {
		 /*line :1*/ panic("division by zero")
	}
	if &uW0eIVjy.u67K4Sv == kWa1YrIHZo ||  /*line :1*/ ap52A4djK(uW0eIVjy.u67K4Sv.qho77PBKF, bw0lYdCKY) {
		bw0lYdCKY =  /*line :1*/ g3X9fa(nil).bj0nVC(bw0lYdCKY)
	}
	uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.bj0nVC(sza5atF.qho77PBKF)
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.bj0nVC(bw0lYdCKY)
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}



func (uW0eIVjy *FGtP7gcPU) SetFrac64(sza5atF, kWa1YrIHZo int64) *FGtP7gcPU {
	if kWa1YrIHZo == 0 {
		 /*line :1*/ panic("division by zero")
	}
	 /*line :1*/ uW0eIVjy.u67K4Sv.SetInt64(sza5atF)
	if kWa1YrIHZo < 0 {
		kWa1YrIHZo = -kWa1YrIHZo
		uW0eIVjy.u67K4Sv.hCTOmp0gckSa = !uW0eIVjy.u67K4Sv.hCTOmp0gckSa
	}
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.glzD00sysLql( /*line :1*/ uint64(kWa1YrIHZo))
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}


func (uW0eIVjy *FGtP7gcPU) SetInt(pmEgen2 *ZMtGuo) *FGtP7gcPU {
	 /*line :1*/ uW0eIVjy.u67K4Sv.Set(pmEgen2)
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
	return uW0eIVjy
}


func (uW0eIVjy *FGtP7gcPU) SetInt64(pmEgen2 int64) *FGtP7gcPU {
	 /*line :1*/ uW0eIVjy.u67K4Sv.SetInt64(pmEgen2)
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
	return uW0eIVjy
}


func (uW0eIVjy *FGtP7gcPU) SetUint64(pmEgen2 uint64) *FGtP7gcPU {
	 /*line :1*/ uW0eIVjy.u67K4Sv.SetUint64(pmEgen2)
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
	return uW0eIVjy
}


func (uW0eIVjy *FGtP7gcPU) Set(pmEgen2 *FGtP7gcPU) *FGtP7gcPU {
	if uW0eIVjy != pmEgen2 {
		 /*line :1*/ uW0eIVjy.u67K4Sv.Set(&pmEgen2.u67K4Sv)
		 /*line :1*/ uW0eIVjy._WnhHVFvm7.Set(&pmEgen2._WnhHVFvm7)
	}
	if  /*line :1*/ len(uW0eIVjy._WnhHVFvm7.qho77PBKF) == 0 {
		uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
	}
	return uW0eIVjy
}


func (uW0eIVjy *FGtP7gcPU) Abs(pmEgen2 *FGtP7gcPU) *FGtP7gcPU {
	 /*line :1*/ uW0eIVjy.Set(pmEgen2)
	uW0eIVjy.u67K4Sv.hCTOmp0gckSa = false
	return uW0eIVjy
}


func (uW0eIVjy *FGtP7gcPU) Neg(pmEgen2 *FGtP7gcPU) *FGtP7gcPU {
	 /*line :1*/ uW0eIVjy.Set(pmEgen2)
	uW0eIVjy.u67K4Sv.hCTOmp0gckSa =  /*line :1*/ len(uW0eIVjy.u67K4Sv.qho77PBKF) > 0 && !uW0eIVjy.u67K4Sv.hCTOmp0gckSa
	return uW0eIVjy
}



func (uW0eIVjy *FGtP7gcPU) Inv(pmEgen2 *FGtP7gcPU) *FGtP7gcPU {
	if  /*line :1*/ len(pmEgen2.u67K4Sv.qho77PBKF) == 0 {
		 /*line :1*/ panic("division by zero")
	}
	 /*line :1*/ uW0eIVjy.Set(pmEgen2)
	uW0eIVjy.u67K4Sv.qho77PBKF, uW0eIVjy._WnhHVFvm7.qho77PBKF = uW0eIVjy._WnhHVFvm7.qho77PBKF, uW0eIVjy.u67K4Sv.qho77PBKF
	return uW0eIVjy
}






func (pmEgen2 *FGtP7gcPU) Sign() int {
	return  /*line :1*/ pmEgen2.u67K4Sv.Sign()
}


func (pmEgen2 *FGtP7gcPU) IsInt() bool {
	return  /*line :1*/ len(pmEgen2._WnhHVFvm7.qho77PBKF) == 0 ||  /*line :1*/ pmEgen2._WnhHVFvm7.qho77PBKF.beaurQHiT(lG0t2KfqGwyP) == 0
}





func (pmEgen2 *FGtP7gcPU) Num() *ZMtGuo {
	return &pmEgen2.u67K4Sv
}








func (pmEgen2 *FGtP7gcPU) Denom() *ZMtGuo {

	if  /*line :1*/ len(pmEgen2._WnhHVFvm7.qho77PBKF) == 0 {

		return &ZMtGuo{qho77PBKF: g3X9fa{1}}
	}
	return &pmEgen2._WnhHVFvm7
}

func (uW0eIVjy *FGtP7gcPU) hitXcy5() *FGtP7gcPU {
	switch {
	case  /*line :1*/ len(uW0eIVjy.u67K4Sv.qho77PBKF) == 0:

		uW0eIVjy.u67K4Sv.hCTOmp0gckSa = false
		fallthrough
	case  /*line :1*/ len(uW0eIVjy._WnhHVFvm7.qho77PBKF) == 0:

		uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
	default:

		s_9Ih_ := uW0eIVjy.u67K4Sv.hCTOmp0gckSa
		uW0eIVjy.u67K4Sv.hCTOmp0gckSa = false
		uW0eIVjy._WnhHVFvm7.hCTOmp0gckSa = false
		if sYV3C6v3w1 :=  /*line :1*/ Hr58YFtm(0).pz5vr0wtn(nil, nil, &uW0eIVjy.u67K4Sv, &uW0eIVjy._WnhHVFvm7);  /*line :1*/ sYV3C6v3w1.Cmp(h2Aj5SV9g2re) != 0 {
			uW0eIVjy.u67K4Sv.qho77PBKF, _ =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.xOZxoLyl(nil, uW0eIVjy.u67K4Sv.qho77PBKF, sYV3C6v3w1.qho77PBKF)
			uW0eIVjy._WnhHVFvm7.qho77PBKF, _ =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.xOZxoLyl(nil, uW0eIVjy._WnhHVFvm7.qho77PBKF, sYV3C6v3w1.qho77PBKF)
		}
		uW0eIVjy.u67K4Sv.hCTOmp0gckSa = s_9Ih_
	}
	return uW0eIVjy
}




func l1mpsCc(uW0eIVjy, pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	switch {
	case  /*line :1*/ len(pmEgen2) == 0 &&  /*line :1*/ len(oxFS5wa5) == 0:
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(1)
	case  /*line :1*/ len(pmEgen2) == 0:
		return  /*line :1*/ uW0eIVjy.bj0nVC(oxFS5wa5)
	case  /*line :1*/ len(oxFS5wa5) == 0:
		return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
	}
	return  /*line :1*/ uW0eIVjy.d22x6Ygoc80O(pmEgen2, oxFS5wa5)
}



func (uW0eIVjy *ZMtGuo) coC9wg(pmEgen2 *ZMtGuo, sYV3C6v3w1 g3X9fa) {
	if  /*line :1*/ len(sYV3C6v3w1) == 0 {
		 /*line :1*/ uW0eIVjy.Set(pmEgen2)
		return
	}
	uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.d22x6Ygoc80O(pmEgen2.qho77PBKF, sYV3C6v3w1)
	uW0eIVjy.hCTOmp0gckSa = pmEgen2.hCTOmp0gckSa
}






func (pmEgen2 *FGtP7gcPU) Cmp(oxFS5wa5 *FGtP7gcPU) int {
	var sza5atF, kWa1YrIHZo ZMtGuo
	 /*line :1*/ sza5atF.coC9wg(&pmEgen2.u67K4Sv, oxFS5wa5._WnhHVFvm7.qho77PBKF)
	 /*line :1*/ kWa1YrIHZo.coC9wg(&oxFS5wa5.u67K4Sv, pmEgen2._WnhHVFvm7.qho77PBKF)
	return  /*line :1*/ sza5atF.Cmp(&kWa1YrIHZo)
}


func (uW0eIVjy *FGtP7gcPU) Add(pmEgen2, oxFS5wa5 *FGtP7gcPU) *FGtP7gcPU {
	var oBH2mFO9, diOe0B ZMtGuo
	 /*line :1*/ oBH2mFO9.coC9wg(&pmEgen2.u67K4Sv, oxFS5wa5._WnhHVFvm7.qho77PBKF)
	 /*line :1*/ diOe0B.coC9wg(&oxFS5wa5.u67K4Sv, pmEgen2._WnhHVFvm7.qho77PBKF)
	 /*line :1*/ uW0eIVjy.u67K4Sv.Add(&oBH2mFO9, &diOe0B)
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ l1mpsCc(uW0eIVjy._WnhHVFvm7.qho77PBKF, pmEgen2._WnhHVFvm7.qho77PBKF, oxFS5wa5._WnhHVFvm7.qho77PBKF)
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}


func (uW0eIVjy *FGtP7gcPU) Sub(pmEgen2, oxFS5wa5 *FGtP7gcPU) *FGtP7gcPU {
	var oBH2mFO9, diOe0B ZMtGuo
	 /*line :1*/ oBH2mFO9.coC9wg(&pmEgen2.u67K4Sv, oxFS5wa5._WnhHVFvm7.qho77PBKF)
	 /*line :1*/ diOe0B.coC9wg(&oxFS5wa5.u67K4Sv, pmEgen2._WnhHVFvm7.qho77PBKF)
	 /*line :1*/ uW0eIVjy.u67K4Sv.Sub(&oBH2mFO9, &diOe0B)
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ l1mpsCc(uW0eIVjy._WnhHVFvm7.qho77PBKF, pmEgen2._WnhHVFvm7.qho77PBKF, oxFS5wa5._WnhHVFvm7.qho77PBKF)
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}


func (uW0eIVjy *FGtP7gcPU) Mul(pmEgen2, oxFS5wa5 *FGtP7gcPU) *FGtP7gcPU {
	if pmEgen2 == oxFS5wa5 {

		uW0eIVjy.u67K4Sv.hCTOmp0gckSa = false
		uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.pJ61taToc(pmEgen2.u67K4Sv.qho77PBKF)
		if  /*line :1*/ len(pmEgen2._WnhHVFvm7.qho77PBKF) == 0 {
			uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pSzoQ7PMW(1)
		} else {
			uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.pJ61taToc(pmEgen2._WnhHVFvm7.qho77PBKF)
		}
		return uW0eIVjy
	}
	 /*line :1*/ uW0eIVjy.u67K4Sv.Mul(&pmEgen2.u67K4Sv, &oxFS5wa5.u67K4Sv)
	uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ l1mpsCc(uW0eIVjy._WnhHVFvm7.qho77PBKF, pmEgen2._WnhHVFvm7.qho77PBKF, oxFS5wa5._WnhHVFvm7.qho77PBKF)
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}



func (uW0eIVjy *FGtP7gcPU) Quo(pmEgen2, oxFS5wa5 *FGtP7gcPU) *FGtP7gcPU {
	if  /*line :1*/ len(oxFS5wa5.u67K4Sv.qho77PBKF) == 0 {
		 /*line :1*/ panic("division by zero")
	}
	var sza5atF, kWa1YrIHZo ZMtGuo
	 /*line :1*/ sza5atF.coC9wg(&pmEgen2.u67K4Sv, oxFS5wa5._WnhHVFvm7.qho77PBKF)
	 /*line :1*/ kWa1YrIHZo.coC9wg(&oxFS5wa5.u67K4Sv, pmEgen2._WnhHVFvm7.qho77PBKF)
	uW0eIVjy.u67K4Sv.qho77PBKF = sza5atF.qho77PBKF
	uW0eIVjy._WnhHVFvm7.qho77PBKF = kWa1YrIHZo.qho77PBKF
	uW0eIVjy.u67K4Sv.hCTOmp0gckSa = sza5atF.hCTOmp0gckSa != kWa1YrIHZo.hCTOmp0gckSa
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}
