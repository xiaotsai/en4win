//line :1
package math







func KRpFivd(_BqkSz0 float64) float64 {
	const (
		TwoM27	= 1.0 / (1 << 27)		
		Two129	= 1 << 129		
		
		R00	= -6.25000000000000000000e-02		
		R01	= 1.40705666955189706048e-03		
		R02	= -1.59955631084035597520e-05		
		R03	= 4.96727999609584448412e-08		
		S01	= 1.91537599538363460805e-02		
		S02	= 1.85946785588630915560e-04		
		S03	= 1.17718464042623683263e-06		
		S04	= 5.04636257076217042715e-09		
		S05	= 1.23542274426137913908e-11		
	)

	switch {
	case  /*line :1*/ OIdLpDqq(_BqkSz0):
		return _BqkSz0
	case  /*line :1*/ ME1vzpD5wcr(_BqkSz0, 0) || _BqkSz0 == 0:
		return 0
	}

	awRlyItEG2gn := false
	if _BqkSz0 < 0 {
		_BqkSz0 = -_BqkSz0
		awRlyItEG2gn = true
	}
	if _BqkSz0 >= 2 {
		aK8LO1, cU6qMJer46oH :=  /*line :1*/ BcjJnykwJ(_BqkSz0)
		jfkBeChFBZ := -aK8LO1 - cU6qMJer46oH
		rHAPx8 := aK8LO1 - cU6qMJer46oH

		if _BqkSz0 < MaxFloat64/2 {
			gZDpjWvXax_q :=  /*line :1*/ Sr_3cKxech(_BqkSz0 + _BqkSz0)
			if aK8LO1*cU6qMJer46oH > 0 {
				rHAPx8 = gZDpjWvXax_q / jfkBeChFBZ
			} else {
				jfkBeChFBZ = gZDpjWvXax_q / rHAPx8
			}
		}

		var gZDpjWvXax_q float64
		if _BqkSz0 > Two129 {
			gZDpjWvXax_q = (1 / SqrtPi) * rHAPx8 /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		} else {
			wkIoLysGG8KY :=  /*line :1*/ fAiqnjA7(_BqkSz0)
			hLZMcJsdS :=  /*line :1*/ s9GAxzhu2DTB(_BqkSz0)
			gZDpjWvXax_q = (1 / SqrtPi) * (wkIoLysGG8KY*rHAPx8 - hLZMcJsdS*jfkBeChFBZ) /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		}
		if awRlyItEG2gn {
			return -gZDpjWvXax_q
		}
		return gZDpjWvXax_q
	}
	if _BqkSz0 < TwoM27 {
		return 0.5 * _BqkSz0
	}
	gZDpjWvXax_q := _BqkSz0 * _BqkSz0
	aBK_X3rbPc := gZDpjWvXax_q * (R00 + gZDpjWvXax_q*(R01+gZDpjWvXax_q*(R02+gZDpjWvXax_q*R03)))
	aK8LO1 := 1.0 + gZDpjWvXax_q*(S01+gZDpjWvXax_q*(S02+gZDpjWvXax_q*(S03+gZDpjWvXax_q*(S04+gZDpjWvXax_q*S05))))
	aBK_X3rbPc *= _BqkSz0
	gZDpjWvXax_q = 0.5*_BqkSz0 + aBK_X3rbPc/aK8LO1
	if awRlyItEG2gn {
		return -gZDpjWvXax_q
	}
	return gZDpjWvXax_q
}









func VyOoq6VwJ(_BqkSz0 float64) float64 {
	const (
		TwoM54	= 1.0 / (1 << 54)		
		Two129	= 1 << 129		
		U00	= -1.96057090646238940668e-01		
		U01	= 5.04438716639811282616e-02		
		U02	= -1.91256895875763547298e-03		
		U03	= 2.35252600561610495928e-05		
		U04	= -9.19099158039878874504e-08		
		V00	= 1.99167318236649903973e-02		
		V01	= 2.02552581025135171496e-04		
		V02	= 1.35608801097516229404e-06		
		V03	= 6.22741452364621501295e-09		
		V04	= 1.66559246207992079114e-11		
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
		jfkBeChFBZ := -aK8LO1 - cU6qMJer46oH
		rHAPx8 := aK8LO1 - cU6qMJer46oH

		if _BqkSz0 < MaxFloat64/2 {
			gZDpjWvXax_q :=  /*line :1*/ Sr_3cKxech(_BqkSz0 + _BqkSz0)
			if aK8LO1*cU6qMJer46oH > 0 {
				rHAPx8 = gZDpjWvXax_q / jfkBeChFBZ
			} else {
				jfkBeChFBZ = gZDpjWvXax_q / rHAPx8
			}
		}

		var gZDpjWvXax_q float64
		if _BqkSz0 > Two129 {
			gZDpjWvXax_q = (1 / SqrtPi) * jfkBeChFBZ /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		} else {
			wkIoLysGG8KY :=  /*line :1*/ fAiqnjA7(_BqkSz0)
			hLZMcJsdS :=  /*line :1*/ s9GAxzhu2DTB(_BqkSz0)
			gZDpjWvXax_q = (1 / SqrtPi) * (wkIoLysGG8KY*jfkBeChFBZ + hLZMcJsdS*rHAPx8) /  /*line :1*/ NeXs7bSyfaCD(_BqkSz0)
		}
		return gZDpjWvXax_q
	}
	if _BqkSz0 <= TwoM54 {
		return -(2 / Pi) / _BqkSz0
	}
	gZDpjWvXax_q := _BqkSz0 * _BqkSz0
	wkIoLysGG8KY := U00 + gZDpjWvXax_q*(U01+gZDpjWvXax_q*(U02+gZDpjWvXax_q*(U03+gZDpjWvXax_q*U04)))
	hLZMcJsdS := 1 + gZDpjWvXax_q*(V00+gZDpjWvXax_q*(V01+gZDpjWvXax_q*(V02+gZDpjWvXax_q*(V03+gZDpjWvXax_q*V04))))
	return _BqkSz0*(wkIoLysGG8KY/hLZMcJsdS) + (2/Pi)*( /*line :1*/ KRpFivd(_BqkSz0)* /*line :1*/ JrJVPb5WbG(_BqkSz0)-1/_BqkSz0)
}


var l82x4H = [6]float64{
	0.00000000000000000000e+00,
	1.17187499999988647970e-01,
	1.32394806593073575129e+01,
	4.12051854307378562225e+02,
	3.87474538913960532227e+03,
	7.91447954031891731574e+03,
}
var cCDLo5MUHy = [5]float64{
	1.14207370375678408436e+02,
	3.65093083420853463394e+03,
	3.69562060269033463555e+04,
	9.76027935934950801311e+04,
	3.08042720627888811578e+04,
}


var nnaXfu4GF = [6]float64{
	1.31990519556243522749e-11,
	1.17187493190614097638e-01,
	6.80275127868432871736e+00,
	1.08308182990189109773e+02,
	5.17636139533199752805e+02,
	5.28715201363337541807e+02,
}
var npzduK0r = [5]float64{
	5.92805987221131331921e+01,
	9.91401418733614377743e+02,
	5.35326695291487976647e+03,
	7.84469031749551231769e+03,
	1.50404688810361062679e+03,
}


var ayWnaKKbWR = [6]float64{
	3.02503916137373618024e-09,
	1.17186865567253592491e-01,
	3.93297750033315640650e+00,
	3.51194035591636932736e+01,
	9.10550110750781271918e+01,
	4.85590685197364919645e+01,
}
var rI1qVLLViX = [5]float64{
	3.47913095001251519989e+01,
	3.36762458747825746741e+02,
	1.04687139975775130551e+03,
	8.90811346398256432622e+02,
	1.03787932439639277504e+02,
}


var dRFWcbyZfTF = [6]float64{
	1.07710830106873743082e-07,
	1.17176219462683348094e-01,
	2.36851496667608785174e+00,
	1.22426109148261232917e+01,
	1.76939711271687727390e+01,
	5.07352312588818499250e+00,
}
var eWaly8K4jJK = [5]float64{
	2.14364859363821409488e+01,
	1.25290227168402751090e+02,
	2.32276469057162813669e+02,
	1.17679373287147100768e+02,
	8.36463893371618283368e+00,
}

func fAiqnjA7(_BqkSz0 float64) float64 {
	var fzJsoqu *[6]float64
	var gNDglc3F2QV *[5]float64
	if _BqkSz0 >= 8 {
		fzJsoqu = &l82x4H
		gNDglc3F2QV = &cCDLo5MUHy
	} else if _BqkSz0 >= 4.5454 {
		fzJsoqu = &nnaXfu4GF
		gNDglc3F2QV = &npzduK0r
	} else if _BqkSz0 >= 2.8571 {
		fzJsoqu = &ayWnaKKbWR
		gNDglc3F2QV = &rI1qVLLViX
	} else if _BqkSz0 >= 2 {
		fzJsoqu = &dRFWcbyZfTF
		gNDglc3F2QV = &eWaly8K4jJK
	}
	gZDpjWvXax_q := 1 / (_BqkSz0 * _BqkSz0)
	aBK_X3rbPc := fzJsoqu[0] + gZDpjWvXax_q*(fzJsoqu[1]+gZDpjWvXax_q*(fzJsoqu[2]+gZDpjWvXax_q*(fzJsoqu[3]+gZDpjWvXax_q*(fzJsoqu[4]+gZDpjWvXax_q*fzJsoqu[5]))))
	aK8LO1 := 1.0 + gZDpjWvXax_q*(gNDglc3F2QV[0]+gZDpjWvXax_q*(gNDglc3F2QV[1]+gZDpjWvXax_q*(gNDglc3F2QV[2]+gZDpjWvXax_q*(gNDglc3F2QV[3]+gZDpjWvXax_q*gNDglc3F2QV[4]))))
	return 1 + aBK_X3rbPc/aK8LO1
}


var _cqNnO = [6]float64{
	0.00000000000000000000e+00,
	-1.02539062499992714161e-01,
	-1.62717534544589987888e+01,
	-7.59601722513950107896e+02,
	-1.18498066702429587167e+04,
	-4.84385124285750353010e+04,
}
var yJ2AxuGpFZDn = [6]float64{
	1.61395369700722909556e+02,
	7.82538599923348465381e+03,
	1.33875336287249578163e+05,
	7.19657723683240939863e+05,
	6.66601232617776375264e+05,
	-2.94490264303834643215e+05,
}


var dBy2abI40 = [6]float64{
	-2.08979931141764104297e-11,
	-1.02539050241375426231e-01,
	-8.05644828123936029840e+00,
	-1.83669607474888380239e+02,
	-1.37319376065508163265e+03,
	-2.61244440453215656817e+03,
}
var obiRXcHO = [6]float64{
	8.12765501384335777857e+01,
	1.99179873460485964642e+03,
	1.74684851924908907677e+04,
	4.98514270910352279316e+04,
	2.79480751638918118260e+04,
	-4.71918354795128470869e+03,
}


var glaiNx35apP = [6]float64{
	-5.07831226461766561369e-09,
	-1.02537829820837089745e-01,
	-4.61011581139473403113e+00,
	-5.78472216562783643212e+01,
	-2.28244540737631695038e+02,
	-2.19210128478909325622e+02,
}
var uNn99abdUa = [6]float64{
	4.76651550323729509273e+01,
	6.73865112676699709482e+02,
	3.38015286679526343505e+03,
	5.54772909720722782367e+03,
	1.90311919338810798763e+03,
	-1.35201191444307340817e+02,
}


var cnIUJPloTz6 = [6]float64{
	-1.78381727510958865572e-07,
	-1.02517042607985553460e-01,
	-2.75220568278187460720e+00,
	-1.96636162643703720221e+01,
	-4.23253133372830490089e+01,
	-2.13719211703704061733e+01,
}
var nmqFuwN5h = [6]float64{
	2.95333629060523854548e+01,
	2.52981549982190529136e+02,
	7.57502834868645436472e+02,
	7.39393205320467245656e+02,
	1.55949003336666123687e+02,
	-4.95949898822628210127e+00,
}

func s9GAxzhu2DTB(_BqkSz0 float64) float64 {
	var fzJsoqu, gNDglc3F2QV *[6]float64
	if _BqkSz0 >= 8 {
		fzJsoqu = &_cqNnO
		gNDglc3F2QV = &yJ2AxuGpFZDn
	} else if _BqkSz0 >= 4.5454 {
		fzJsoqu = &dBy2abI40
		gNDglc3F2QV = &obiRXcHO
	} else if _BqkSz0 >= 2.8571 {
		fzJsoqu = &glaiNx35apP
		gNDglc3F2QV = &uNn99abdUa
	} else if _BqkSz0 >= 2 {
		fzJsoqu = &cnIUJPloTz6
		gNDglc3F2QV = &nmqFuwN5h
	}
	gZDpjWvXax_q := 1 / (_BqkSz0 * _BqkSz0)
	aBK_X3rbPc := fzJsoqu[0] + gZDpjWvXax_q*(fzJsoqu[1]+gZDpjWvXax_q*(fzJsoqu[2]+gZDpjWvXax_q*(fzJsoqu[3]+gZDpjWvXax_q*(fzJsoqu[4]+gZDpjWvXax_q*fzJsoqu[5]))))
	aK8LO1 := 1 + gZDpjWvXax_q*(gNDglc3F2QV[0]+gZDpjWvXax_q*(gNDglc3F2QV[1]+gZDpjWvXax_q*(gNDglc3F2QV[2]+gZDpjWvXax_q*(gNDglc3F2QV[3]+gZDpjWvXax_q*(gNDglc3F2QV[4]+gZDpjWvXax_q*gNDglc3F2QV[5])))))
	return (0.375 + aBK_X3rbPc/aK8LO1) / _BqkSz0
}
