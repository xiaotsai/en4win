//line :1
package big

import bits "math/bits"


type VYe6D0 uint

const (
	_S	= _W / 8		

	_W	= bits.UintSize		
	_B	= 1 << _W		
	_M	= _B - 1		
)


func mulWW(pmEgen2, oxFS5wa5 VYe6D0) (mQ0XlC6n_5ot, a174Q6gz VYe6D0) {
	mZ6AbmRX_M, dTexDSO :=  /*line :1*/ bits.Mul( /*line :1*/ uint(pmEgen2),  /*line :1*/ uint(oxFS5wa5))
	return  /*line :1*/ VYe6D0(mZ6AbmRX_M),  /*line :1*/ VYe6D0(dTexDSO)
}


func e5eHrK(pmEgen2, oxFS5wa5, vKNCqP0 VYe6D0) (mQ0XlC6n_5ot, a174Q6gz VYe6D0) {
	mZ6AbmRX_M, dTexDSO :=  /*line :1*/ bits.Mul( /*line :1*/ uint(pmEgen2),  /*line :1*/ uint(oxFS5wa5))
	var gwVXXmzb uint
	dTexDSO, gwVXXmzb =  /*line :1*/ bits.Add(dTexDSO,  /*line :1*/ uint(vKNCqP0), 0)
	return  /*line :1*/ VYe6D0(mZ6AbmRX_M + gwVXXmzb),  /*line :1*/ VYe6D0(dTexDSO)
}



func bUgMSHLa(pmEgen2 VYe6D0) uint {
	return  /*line :1*/ uint( /*line :1*/ bits.M7qYV0PcLYO( /*line :1*/ uint(pmEgen2)))
}


func eSsyLwRrPmuh(uW0eIVjy, pmEgen2, oxFS5wa5 []VYe6D0) (vKNCqP0 VYe6D0) {

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2) && aepbxLiWOZ <  /*line :1*/ len(oxFS5wa5); aepbxLiWOZ++ {
		dfaObZ, gwVXXmzb :=  /*line :1*/ bits.Add( /*line :1*/ uint(pmEgen2[aepbxLiWOZ]),  /*line :1*/ uint(oxFS5wa5[aepbxLiWOZ]),  /*line :1*/ uint(vKNCqP0))
		uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0(dfaObZ)
		vKNCqP0 =  /*line :1*/ VYe6D0(gwVXXmzb)
	}
	return
}


func wXXlHbp7vT(uW0eIVjy, pmEgen2, oxFS5wa5 []VYe6D0) (vKNCqP0 VYe6D0) {

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2) && aepbxLiWOZ <  /*line :1*/ len(oxFS5wa5); aepbxLiWOZ++ {
		dfaObZ, gwVXXmzb :=  /*line :1*/ bits.Sub( /*line :1*/ uint(pmEgen2[aepbxLiWOZ]),  /*line :1*/ uint(oxFS5wa5[aepbxLiWOZ]),  /*line :1*/ uint(vKNCqP0))
		uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0(dfaObZ)
		vKNCqP0 =  /*line :1*/ VYe6D0(gwVXXmzb)
	}
	return
}


func oMoryXyqSn(uW0eIVjy, pmEgen2 []VYe6D0, oxFS5wa5 VYe6D0) (vKNCqP0 VYe6D0) {
	vKNCqP0 = oxFS5wa5

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2); aepbxLiWOZ++ {
		dfaObZ, gwVXXmzb :=  /*line :1*/ bits.Add( /*line :1*/ uint(pmEgen2[aepbxLiWOZ]),  /*line :1*/ uint(vKNCqP0), 0)
		uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0(dfaObZ)
		vKNCqP0 =  /*line :1*/ VYe6D0(gwVXXmzb)
	}
	return
}








func zfga2o(uW0eIVjy, pmEgen2 []VYe6D0, oxFS5wa5 VYe6D0) (vKNCqP0 VYe6D0) {
	vKNCqP0 = oxFS5wa5

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2); aepbxLiWOZ++ {
		if vKNCqP0 == 0 {
			 /*line :1*/ copy(uW0eIVjy[aepbxLiWOZ:], pmEgen2[aepbxLiWOZ:])
			return
		}
		dfaObZ, gwVXXmzb :=  /*line :1*/ bits.Add( /*line :1*/ uint(pmEgen2[aepbxLiWOZ]),  /*line :1*/ uint(vKNCqP0), 0)
		uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0(dfaObZ)
		vKNCqP0 =  /*line :1*/ VYe6D0(gwVXXmzb)
	}
	return
}

func w66QQi8o8b6e(uW0eIVjy, pmEgen2 []VYe6D0, oxFS5wa5 VYe6D0) (vKNCqP0 VYe6D0) {
	vKNCqP0 = oxFS5wa5

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2); aepbxLiWOZ++ {
		dfaObZ, gwVXXmzb :=  /*line :1*/ bits.Sub( /*line :1*/ uint(pmEgen2[aepbxLiWOZ]),  /*line :1*/ uint(vKNCqP0), 0)
		uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0(dfaObZ)
		vKNCqP0 =  /*line :1*/ VYe6D0(gwVXXmzb)
	}
	return
}


func hRp2II(uW0eIVjy, pmEgen2 []VYe6D0, oxFS5wa5 VYe6D0) (vKNCqP0 VYe6D0) {
	vKNCqP0 = oxFS5wa5

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2); aepbxLiWOZ++ {
		if vKNCqP0 == 0 {
			 /*line :1*/ copy(uW0eIVjy[aepbxLiWOZ:], pmEgen2[aepbxLiWOZ:])
			return
		}
		dfaObZ, gwVXXmzb :=  /*line :1*/ bits.Sub( /*line :1*/ uint(pmEgen2[aepbxLiWOZ]),  /*line :1*/ uint(vKNCqP0), 0)
		uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0(dfaObZ)
		vKNCqP0 =  /*line :1*/ VYe6D0(gwVXXmzb)
	}
	return
}

func q430LRW(uW0eIVjy, pmEgen2 []VYe6D0, nR7KU4mGsdy uint) (vKNCqP0 VYe6D0) {
	if nR7KU4mGsdy == 0 {
		 /*line :1*/ copy(uW0eIVjy, pmEgen2)
		return
	}
	if  /*line :1*/ len(uW0eIVjy) == 0 {
		return
	}
	nR7KU4mGsdy &= _W - 1
	vLnzt7X6K := _W - nR7KU4mGsdy
	vLnzt7X6K &= _W - 1
	vKNCqP0 = pmEgen2[ /*line :1*/ len(uW0eIVjy)-1] >> vLnzt7X6K
	for aepbxLiWOZ :=  /*line :1*/ len(uW0eIVjy) - 1; aepbxLiWOZ > 0; aepbxLiWOZ-- {
		uW0eIVjy[aepbxLiWOZ] = pmEgen2[aepbxLiWOZ]<<nR7KU4mGsdy | pmEgen2[aepbxLiWOZ-1]>>vLnzt7X6K
	}
	uW0eIVjy[0] = pmEgen2[0] << nR7KU4mGsdy
	return
}

func mD6ZCtIf1Bd(uW0eIVjy, pmEgen2 []VYe6D0, nR7KU4mGsdy uint) (vKNCqP0 VYe6D0) {
	if nR7KU4mGsdy == 0 {
		 /*line :1*/ copy(uW0eIVjy, pmEgen2)
		return
	}
	if  /*line :1*/ len(uW0eIVjy) == 0 {
		return
	}
	if  /*line :1*/ len(pmEgen2) !=  /*line :1*/ len(uW0eIVjy) {

		 /*line :1*/ panic("len(x) != len(z)")
	}
	nR7KU4mGsdy &= _W - 1
	vLnzt7X6K := _W - nR7KU4mGsdy
	vLnzt7X6K &= _W - 1
	vKNCqP0 = pmEgen2[0] << vLnzt7X6K
	for aepbxLiWOZ := 1; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy); aepbxLiWOZ++ {
		uW0eIVjy[aepbxLiWOZ-1] = pmEgen2[aepbxLiWOZ-1]>>nR7KU4mGsdy | pmEgen2[aepbxLiWOZ]<<vLnzt7X6K
	}
	uW0eIVjy[ /*line :1*/ len(uW0eIVjy)-1] = pmEgen2[ /*line :1*/ len(uW0eIVjy)-1] >> nR7KU4mGsdy
	return
}

func mEUqf35hV(uW0eIVjy, pmEgen2 []VYe6D0, oxFS5wa5, vFx5p_cm VYe6D0) (vKNCqP0 VYe6D0) {
	vKNCqP0 = vFx5p_cm

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2); aepbxLiWOZ++ {
		vKNCqP0, uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ e5eHrK(pmEgen2[aepbxLiWOZ], oxFS5wa5, vKNCqP0)
	}
	return
}

func soJqcXF(uW0eIVjy, pmEgen2 []VYe6D0, oxFS5wa5 VYe6D0) (vKNCqP0 VYe6D0) {

	for aepbxLiWOZ := 0; aepbxLiWOZ <  /*line :1*/ len(uW0eIVjy) && aepbxLiWOZ <  /*line :1*/ len(pmEgen2); aepbxLiWOZ++ {
		mQ0XlC6n_5ot, a174Q6gz :=  /*line :1*/ e5eHrK(pmEgen2[aepbxLiWOZ], oxFS5wa5, uW0eIVjy[aepbxLiWOZ])
		dTexDSO, gwVXXmzb :=  /*line :1*/ bits.Add( /*line :1*/ uint(a174Q6gz),  /*line :1*/ uint(vKNCqP0), 0)
		vKNCqP0, uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0(gwVXXmzb),  /*line :1*/ VYe6D0(dTexDSO)
		vKNCqP0 += mQ0XlC6n_5ot
	}
	return
}




func f6uWqUJ(dKR0NUYWi, _mZiyQO, oxFS5wa5, gS2TXhHpYaX4 VYe6D0) (yL0_p0wUnc, vFx5p_cm VYe6D0) {
	nR7KU4mGsdy :=  /*line :1*/ bUgMSHLa(oxFS5wa5)
	if nR7KU4mGsdy != 0 {
		dKR0NUYWi = dKR0NUYWi<<nR7KU4mGsdy | _mZiyQO>>(_W-nR7KU4mGsdy)
		_mZiyQO <<= nR7KU4mGsdy
		oxFS5wa5 <<= nR7KU4mGsdy
	}
	s31g8J :=  /*line :1*/ uint(oxFS5wa5)

	dms1sqFQT, i6DQQOlkb :=  /*line :1*/ bits.Mul( /*line :1*/ uint(gS2TXhHpYaX4),  /*line :1*/ uint(dKR0NUYWi))
	_, vKNCqP0 :=  /*line :1*/ bits.Add(i6DQQOlkb,  /*line :1*/ uint(_mZiyQO), 0)
	dms1sqFQT, _ =  /*line :1*/ bits.Add(dms1sqFQT,  /*line :1*/ uint(dKR0NUYWi), vKNCqP0)

	eqTwRvCQS := dms1sqFQT

	oAkS_mN, alz0gE3 :=  /*line :1*/ bits.Mul(s31g8J, eqTwRvCQS)
	axD9ObuovR, kWa1YrIHZo :=  /*line :1*/ bits.Sub( /*line :1*/ uint(_mZiyQO), alz0gE3, 0)
	s47LHZlQFa, _ :=  /*line :1*/ bits.Sub( /*line :1*/ uint(dKR0NUYWi), oAkS_mN, kWa1YrIHZo)

	if s47LHZlQFa != 0 {
		eqTwRvCQS++
		axD9ObuovR -= s31g8J
	}

	if axD9ObuovR >= s31g8J {
		eqTwRvCQS++
		axD9ObuovR -= s31g8J
	}
	return  /*line :1*/ VYe6D0(eqTwRvCQS),  /*line :1*/ VYe6D0(axD9ObuovR >> nR7KU4mGsdy)
}


func niNOGPaIxagq(pStBhY VYe6D0) VYe6D0 {
	sOU3zRv :=  /*line :1*/ uint(pStBhY <<  /*line :1*/ bUgMSHLa(pStBhY))
	dKR0NUYWi := ^sOU3zRv
	_mZiyQO :=  /*line :1*/ uint(_M)
	jaaWqY, _ :=  /*line :1*/ bits.Div(dKR0NUYWi, _mZiyQO, sOU3zRv)
	return  /*line :1*/ VYe6D0(jaaWqY)
}
