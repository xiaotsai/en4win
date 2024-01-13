//line :1
package math








func TZVYtdEXoPFm(_BqkSz0 float64) float64 {
	const (
		Huge	= 1e300
		TwoM27	= 1.0 / (1 << 27)		
		TwoM13	= 1.0 / (1 << 13)		
		Two129	= 1 << 129		
		
		R02	= 1.56249999999999947958e-02		
		R03	= -1.89979294238854721751e-04		
		R04	= 1.82954049532700665670e-06		
		R05	= -4.61832688532103189199e-09		
		S01	= 1.56191029464890010492e-02		
		S02	= 1.16926784663337450260e-04		
		S03	= 5.13546550207318111446e-07		
		S04	= 1.16614003333790000205e-09		
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		return _BqkSz0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0):
		return 0
	case _BqkSz0 == 0:
		return 1
	}

	_BqkSz0 =  /*line :1*/ Abs(_BqkSz0)
	if _BqkSz0 >= 2 {
		aK8LO1, cU6qMJer46oH :=  /*line :1*/ BcjJnykwJ(_BqkSz0)
		jfkBeChFBZ := aK8LO1 - cU6qMJer46oH
		rHAPx8 := aK8LO1 + cU6qMJer46oH

		if _BqkSz0 < MaxFloat64/2 {
			gZDpjWvXax_q := - /*line :1*/ Sr_3cKxech(_BqkSz0 + _BqkSz0)
			if aK8LO1*cU6qMJer46oH < 0 {
				rHAPx8 = gZDpjWvXax_q / jfkBeChFBZ
			} else {
				jfkBeChFBZ = gZDpjWvXax_q / rHAPx8
			}
		}

		var gZDpjWvXax_q float64
		if _BqkSz0 > Two129 {
			gZDpjWvXax_q = (1 / SqrtPi) * rHAPx8 /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		} else {
			wkIoLysGG8KY :=  /*line :1*/ lBfEpwKdoHP(_BqkSz0)
			hLZMcJsdS :=  /*line :1*/ caymD5xDljaM(_BqkSz0)
			gZDpjWvXax_q = (1 / SqrtPi) * (wkIoLysGG8KY*rHAPx8 - hLZMcJsdS*jfkBeChFBZ) /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		}
		return gZDpjWvXax_q
	}
	if _BqkSz0 < TwoM13 {
		if _BqkSz0 < TwoM27 {
			return 1
		}
		return 1 - 0.25*_BqkSz0*_BqkSz0
	}
	gZDpjWvXax_q := _BqkSz0 * _BqkSz0
	aBK_X3rbPc := gZDpjWvXax_q * (R02 + gZDpjWvXax_q*(R03+gZDpjWvXax_q*(R04+gZDpjWvXax_q*R05)))
	aK8LO1 := 1 + gZDpjWvXax_q*(S01+gZDpjWvXax_q*(S02+gZDpjWvXax_q*(S03+gZDpjWvXax_q*S04)))
	if _BqkSz0 < 1 {
		return 1 + gZDpjWvXax_q*(-0.25+(aBK_X3rbPc/aK8LO1))
	}
	wkIoLysGG8KY := 0.5 * _BqkSz0
	return (1+wkIoLysGG8KY)*(1-wkIoLysGG8KY) + gZDpjWvXax_q*(aBK_X3rbPc/aK8LO1)
}









func XiHpf21SIVk(_BqkSz0 float64) float64 {
	const (
		TwoM27	= 1.0 / (1 << 27)		
		Two129	= 1 << 129		
		U00	= -7.38042951086872317523e-02		
		U01	= 1.76666452509181115538e-01		
		U02	= -1.38185671945596898896e-02		
		U03	= 3.47453432093683650238e-04		
		U04	= -3.81407053724364161125e-06		
		U05	= 1.95590137035022920206e-08		
		U06	= -3.98205194132103398453e-11		
		V01	= 1.27304834834123699328e-02		
		V02	= 7.60068627350353253702e-05		
		V03	= 2.59150851840457805467e-07		
		V04	= 4.41110311332675467403e-10		
	)

	switch {
	case _BqkSz0 < 0 ||  /*line :1*/ OIdLpDqq(_BqkSz0):
		return  /*line :1*/ WG0hZIT4()
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 1):
		return 0
	case _BqkSz0 == 0:
		return  /*line :1*/ FSug4WHZSBwz(-1)
	}

	if _BqkSz0 >= 2 {

		aK8LO1, cU6qMJer46oH :=  /*line :1*/ BcjJnykwJ(_BqkSz0)
		jfkBeChFBZ := aK8LO1 - cU6qMJer46oH
		rHAPx8 := aK8LO1 + cU6qMJer46oH

		if _BqkSz0 < MaxFloat64/2 {
			gZDpjWvXax_q := - /*line :1*/ Sr_3cKxech(_BqkSz0 + _BqkSz0)
			if aK8LO1*cU6qMJer46oH < 0 {
				rHAPx8 = gZDpjWvXax_q / jfkBeChFBZ
			} else {
				jfkBeChFBZ = gZDpjWvXax_q / rHAPx8
			}
		}
		var gZDpjWvXax_q float64
		if _BqkSz0 > Two129 {
			gZDpjWvXax_q = (1 / SqrtPi) * jfkBeChFBZ /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		} else {
			wkIoLysGG8KY :=  /*line :1*/ lBfEpwKdoHP(_BqkSz0)
			hLZMcJsdS :=  /*line :1*/ caymD5xDljaM(_BqkSz0)
			gZDpjWvXax_q = (1 / SqrtPi) * (wkIoLysGG8KY*jfkBeChFBZ + hLZMcJsdS*rHAPx8) /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		}
		return gZDpjWvXax_q
	}
	if _BqkSz0 <= TwoM27 {
		return U00 + (2/Pi)* /*line :1*/ JrJVPb5WbG(_BqkSz0)
	}
	gZDpjWvXax_q := _BqkSz0 * _BqkSz0
	wkIoLysGG8KY := U00 + gZDpjWvXax_q*(U01+gZDpjWvXax_q*(U02+gZDpjWvXax_q*(U03+gZDpjWvXax_q*(U04+gZDpjWvXax_q*(U05+gZDpjWvXax_q*U06)))))
	hLZMcJsdS := 1 + gZDpjWvXax_q*(V01+gZDpjWvXax_q*(V02+gZDpjWvXax_q*(V03+gZDpjWvXax_q*V04)))
	return wkIoLysGG8KY/hLZMcJsdS + (2/Pi)* /*line :1*/ TZVYtdEXoPFm(_BqkSz0)* /*line :1*/ JrJVPb5WbG(_BqkSz0)
}


var hUToKjoNJfoR = [6]float64{
	0.00000000000000000000e+00,
	-7.03124999999900357484e-02,
	-8.08167041275349795626e+00,
	-2.57063105679704847262e+02,
	-2.48521641009428822144e+03,
	-5.25304380490729545272e+03,
}
var mtIDvPH = [5]float64{
	1.16534364619668181717e+02,
	3.83374475364121826715e+03,
	4.05978572648472545552e+04,
	1.16752972564375915681e+05,
	4.76277284146730962675e+04,
}


var hIiJcWjs9 = [6]float64{
	-1.14125464691894502584e-11,
	-7.03124940873599280078e-02,
	-4.15961064470587782438e+00,
	-6.76747652265167261021e+01,
	-3.31231299649172967747e+02,
	-3.46433388365604912451e+02,
}
var cMFyf2lLxdDy = [5]float64{
	6.07539382692300335975e+01,
	1.05125230595704579173e+03,
	5.97897094333855784498e+03,
	9.62544514357774460223e+03,
	2.40605815922939109441e+03,
}


var rqgIR18GPD = [6]float64{
	-2.54704601771951915620e-09,
	-7.03119616381481654654e-02,
	-2.40903221549529611423e+00,
	-2.19659774734883086467e+01,
	-5.80791704701737572236e+01,
	-3.14479470594888503854e+01,
}
var jw0fDL_qdBMx = [5]float64{
	3.58560338055209726349e+01,
	3.61513983050303863820e+02,
	1.19360783792111533330e+03,
	1.12799679856907414432e+03,
	1.73580930813335754692e+02,
}


var cZC6SFkIAgu1 = [6]float64{
	-8.87534333032526411254e-08,
	-7.03030995483624743247e-02,
	-1.45073846780952986357e+00,
	-7.63569613823527770791e+00,
	-1.11931668860356747786e+01,
	-3.23364579351335335033e+00,
}
var e19xfU0HS_5 = [5]float64{
	2.22202997532088808441e+01,
	1.36206794218215208048e+02,
	2.70470278658083486789e+02,
	1.53875394208320329881e+02,
	1.46576176948256193810e+01,
}

func lBfEpwKdoHP(_BqkSz0 float64) float64 {
	var fzJsoqu *[6]float64
	var gNDglc3F2QV *[5]float64
	if _BqkSz0 >= 8 {
		fzJsoqu = &hUToKjoNJfoR
		gNDglc3F2QV = &mtIDvPH
	} else if _BqkSz0 >= 4.5454 {
		fzJsoqu = &hIiJcWjs9
		gNDglc3F2QV = &cMFyf2lLxdDy
	} else if _BqkSz0 >= 2.8571 {
		fzJsoqu = &rqgIR18GPD
		gNDglc3F2QV = &jw0fDL_qdBMx
	} else if _BqkSz0 >= 2 {
		fzJsoqu = &cZC6SFkIAgu1
		gNDglc3F2QV = &e19xfU0HS_5
	}
	gZDpjWvXax_q := 1 / (_BqkSz0 * _BqkSz0)
	aBK_X3rbPc := fzJsoqu[0] + gZDpjWvXax_q*(fzJsoqu[1]+gZDpjWvXax_q*(fzJsoqu[2]+gZDpjWvXax_q*(fzJsoqu[3]+gZDpjWvXax_q*(fzJsoqu[4]+gZDpjWvXax_q*fzJsoqu[5]))))
	aK8LO1 := 1 + gZDpjWvXax_q*(gNDglc3F2QV[0]+gZDpjWvXax_q*(gNDglc3F2QV[1]+gZDpjWvXax_q*(gNDglc3F2QV[2]+gZDpjWvXax_q*(gNDglc3F2QV[3]+gZDpjWvXax_q*gNDglc3F2QV[4]))))
	return 1 + aBK_X3rbPc/aK8LO1
}


var mGwG9NG = [6]float64{
	0.00000000000000000000e+00,
	7.32421874999935051953e-02,
	1.17682064682252693899e+01,
	5.57673380256401856059e+02,
	8.85919720756468632317e+03,
	3.70146267776887834771e+04,
}
var vqFilt = [6]float64{
	1.63776026895689824414e+02,
	8.09834494656449805916e+03,
	1.42538291419120476348e+05,
	8.03309257119514397345e+05,
	8.40501579819060512818e+05,
	-3.43899293537866615225e+05,
}


var vHne4ouq = [6]float64{
	1.84085963594515531381e-11,
	7.32421766612684765896e-02,
	5.83563508962056953777e+00,
	1.35111577286449829671e+02,
	1.02724376596164097464e+03,
	1.98997785864605384631e+03,
}
var zFgIZuN = [6]float64{
	8.27766102236537761883e+01,
	2.07781416421392987104e+03,
	1.88472887785718085070e+04,
	5.67511122894947329769e+04,
	3.59767538425114471465e+04,
	-5.35434275601944773371e+03,
}


var jdcgh_LaC1l = [6]float64{
	4.37741014089738620906e-09,
	7.32411180042911447163e-02,
	3.34423137516170720929e+00,
	4.26218440745412650017e+01,
	1.70808091340565596283e+02,
	1.66733948696651168575e+02,
}
var c4F22rqKbe = [6]float64{
	4.87588729724587182091e+01,
	7.09689221056606015736e+02,
	3.70414822620111362994e+03,
	6.46042516752568917582e+03,
	2.51633368920368957333e+03,
	-1.49247451836156386662e+02,
}


var cauBLHf6 = [6]float64{
	1.50444444886983272379e-07,
	7.32234265963079278272e-02,
	1.99819174093815998816e+00,
	1.44956029347885735348e+01,
	3.16662317504781540833e+01,
	1.62527075710929267416e+01,
}
var d9gznujkJit = [6]float64{
	3.03655848355219184498e+01,
	2.69348118608049844624e+02,
	8.44783757595320139444e+02,
	8.82935845112488550512e+02,
	2.12666388511798828631e+02,
	-5.31095493882666946917e+00,
}

func caymD5xDljaM(_BqkSz0 float64) float64 {
	var fzJsoqu, gNDglc3F2QV *[6]float64
	if _BqkSz0 >= 8 {
		fzJsoqu = &mGwG9NG
		gNDglc3F2QV = &vqFilt
	} else if _BqkSz0 >= 4.5454 {
		fzJsoqu = &vHne4ouq
		gNDglc3F2QV = &zFgIZuN
	} else if _BqkSz0 >= 2.8571 {
		fzJsoqu = &jdcgh_LaC1l
		gNDglc3F2QV = &c4F22rqKbe
	} else if _BqkSz0 >= 2 {
		fzJsoqu = &cauBLHf6
		gNDglc3F2QV = &d9gznujkJit
	}
	gZDpjWvXax_q := 1 / (_BqkSz0 * _BqkSz0)
	aBK_X3rbPc := fzJsoqu[0] + gZDpjWvXax_q*(fzJsoqu[1]+gZDpjWvXax_q*(fzJsoqu[2]+gZDpjWvXax_q*(fzJsoqu[3]+gZDpjWvXax_q*(fzJsoqu[4]+gZDpjWvXax_q*fzJsoqu[5]))))
	aK8LO1 := 1 + gZDpjWvXax_q*(gNDglc3F2QV[0]+gZDpjWvXax_q*(gNDglc3F2QV[1]+gZDpjWvXax_q*(gNDglc3F2QV[2]+gZDpjWvXax_q*(gNDglc3F2QV[3]+gZDpjWvXax_q*(gNDglc3F2QV[4]+gZDpjWvXax_q*gNDglc3F2QV[5])))))
	return (-0.125 + aBK_X3rbPc/aK8LO1) / _BqkSz0
}
