//line :1
package big

import (
	binary "yLm9uBQG"
	bits "math/bits"
	rand "os208GicmRnc"
	sync "sync"
)












type g3X9fa []VYe6D0

var (
	lG0t2KfqGwyP	= g3X9fa{1}
	euJqHOlfT	= g3X9fa{2}
	m5xi9O	= g3X9fa{5}
	syj6Abagk	= g3X9fa{10}
)

func (uW0eIVjy g3X9fa) String() string {
	return "0x" +  /*line :1*/ string( /*line :1*/ uW0eIVjy.sgVtOtra(false, 16))
}

func (uW0eIVjy g3X9fa) uKEPQHWEut() {
	for aepbxLiWOZ := range uW0eIVjy {
		uW0eIVjy[aepbxLiWOZ] = 0
	}
}

func (uW0eIVjy g3X9fa) hitXcy5() g3X9fa {
	aepbxLiWOZ :=  /*line :1*/ len(uW0eIVjy)
	for aepbxLiWOZ > 0 && uW0eIVjy[aepbxLiWOZ-1] == 0 {
		aepbxLiWOZ--
	}
	return uW0eIVjy[0:aepbxLiWOZ]
}

func (uW0eIVjy g3X9fa) gS9bqHpKAJ0(h_Wr_yqq int) g3X9fa {
	if h_Wr_yqq <=  /*line :1*/ cap(uW0eIVjy) {
		return uW0eIVjy[:h_Wr_yqq]
	}
	if h_Wr_yqq == 1 {

		return  /*line :1*/ make(g3X9fa, 1)
	}
	
	
	const e = 4	
	return  /*line :1*/ make(g3X9fa, h_Wr_yqq, h_Wr_yqq+e)
}

func (uW0eIVjy g3X9fa) pSzoQ7PMW(pmEgen2 VYe6D0) g3X9fa {
	if pmEgen2 == 0 {
		return uW0eIVjy[:0]
	}
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(1)
	uW0eIVjy[0] = pmEgen2
	return uW0eIVjy
}

func (uW0eIVjy g3X9fa) glzD00sysLql(pmEgen2 uint64) g3X9fa {

	if uKQ8Eox :=  /*line :1*/ VYe6D0(pmEgen2);  /*line :1*/ uint64(uKQ8Eox) == pmEgen2 {
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(uKQ8Eox)
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(2)
	uW0eIVjy[1] =  /*line :1*/ VYe6D0(pmEgen2 >> 32)
	uW0eIVjy[0] =  /*line :1*/ VYe6D0(pmEgen2)
	return uW0eIVjy
}

func (uW0eIVjy g3X9fa) bj0nVC(pmEgen2 g3X9fa) g3X9fa {
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0( /*line :1*/ len(pmEgen2))
	 /*line :1*/ copy(uW0eIVjy, pmEgen2)
	return uW0eIVjy
}

func (uW0eIVjy g3X9fa) jXv7nnhY_(pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)

	switch {
	case gS2TXhHpYaX4 < h_Wr_yqq:
		return  /*line :1*/ uW0eIVjy.jXv7nnhY_(oxFS5wa5, pmEgen2)
	case gS2TXhHpYaX4 == 0:

		return uW0eIVjy[:0]
	case h_Wr_yqq == 0:

		return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4 + 1)
	vKNCqP0 :=  /*line :1*/ vmnD6P1U(uW0eIVjy[0:h_Wr_yqq], pmEgen2, oxFS5wa5)
	if gS2TXhHpYaX4 > h_Wr_yqq {
		vKNCqP0 =  /*line :1*/ aJRvTH(uW0eIVjy[h_Wr_yqq:gS2TXhHpYaX4], pmEgen2[h_Wr_yqq:], vKNCqP0)
	}
	uW0eIVjy[gS2TXhHpYaX4] = vKNCqP0

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}

func (uW0eIVjy g3X9fa) lt4VKILNCo(pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)

	switch {
	case gS2TXhHpYaX4 < h_Wr_yqq:
		 /*line :1*/ panic("underflow")
	case gS2TXhHpYaX4 == 0:

		return uW0eIVjy[:0]
	case h_Wr_yqq == 0:

		return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4)
	vKNCqP0 :=  /*line :1*/ omPlI3q_Zc(uW0eIVjy[0:h_Wr_yqq], pmEgen2, oxFS5wa5)
	if gS2TXhHpYaX4 > h_Wr_yqq {
		vKNCqP0 =  /*line :1*/ sQp2wtu(uW0eIVjy[h_Wr_yqq:], pmEgen2[h_Wr_yqq:], vKNCqP0)
	}
	if vKNCqP0 != 0 {
		 /*line :1*/ panic("underflow")
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}

func (pmEgen2 g3X9fa) beaurQHiT(oxFS5wa5 g3X9fa) (vFx5p_cm int) {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)
	if gS2TXhHpYaX4 != h_Wr_yqq || gS2TXhHpYaX4 == 0 {
		switch {
		case gS2TXhHpYaX4 < h_Wr_yqq:
			vFx5p_cm = -1
		case gS2TXhHpYaX4 > h_Wr_yqq:
			vFx5p_cm = 1
		}
		return
	}

	aepbxLiWOZ := gS2TXhHpYaX4 - 1
	for aepbxLiWOZ > 0 && pmEgen2[aepbxLiWOZ] == oxFS5wa5[aepbxLiWOZ] {
		aepbxLiWOZ--
	}

	switch {
	case pmEgen2[aepbxLiWOZ] < oxFS5wa5[aepbxLiWOZ]:
		vFx5p_cm = -1
	case pmEgen2[aepbxLiWOZ] > oxFS5wa5[aepbxLiWOZ]:
		vFx5p_cm = 1
	}
	return
}

func (uW0eIVjy g3X9fa) aTzUIOqJPczT(pmEgen2 g3X9fa, oxFS5wa5, vFx5p_cm VYe6D0) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	if gS2TXhHpYaX4 == 0 || oxFS5wa5 == 0 {
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(vFx5p_cm)
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4 + 1)
	uW0eIVjy[gS2TXhHpYaX4] =  /*line :1*/ pMUCN3(uW0eIVjy[0:gS2TXhHpYaX4], pmEgen2, oxFS5wa5, vFx5p_cm)

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}



func oGJ1hCL(uW0eIVjy, pmEgen2, oxFS5wa5 g3X9fa) {
	 /*line :1*/ uW0eIVjy[0 :  /*line :1*/ len(pmEgen2)+ /*line :1*/ len(oxFS5wa5)].uKEPQHWEut()
	for aepbxLiWOZ, s31g8J := range oxFS5wa5 {
		if s31g8J != 0 {
			uW0eIVjy[ /*line :1*/ len(pmEgen2)+aepbxLiWOZ] =  /*line :1*/ fklqwGXkLu(uW0eIVjy[aepbxLiWOZ:aepbxLiWOZ+ /*line :1*/ len(pmEgen2)], pmEgen2, s31g8J)
		}
	}
}










func (uW0eIVjy g3X9fa) vuR24NZ(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 g3X9fa, _DfINwXQ VYe6D0, h_Wr_yqq int) g3X9fa {

	if  /*line :1*/ len(pmEgen2) != h_Wr_yqq ||  /*line :1*/ len(oxFS5wa5) != h_Wr_yqq ||  /*line :1*/ len(gS2TXhHpYaX4) != h_Wr_yqq {
		 /*line :1*/ panic("math/big: mismatched montgomery number lengths")
	}
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(h_Wr_yqq * 2)
	 /*line :1*/ uW0eIVjy.uKEPQHWEut()
	var vKNCqP0 VYe6D0
	for aepbxLiWOZ := 0; aepbxLiWOZ < h_Wr_yqq; aepbxLiWOZ++ {
		s31g8J := oxFS5wa5[aepbxLiWOZ]
		j_mxLI :=  /*line :1*/ fklqwGXkLu(uW0eIVjy[aepbxLiWOZ:h_Wr_yqq+aepbxLiWOZ], pmEgen2, s31g8J)
		amoa4Px := uW0eIVjy[aepbxLiWOZ] * _DfINwXQ
		ljKfRQ :=  /*line :1*/ fklqwGXkLu(uW0eIVjy[aepbxLiWOZ:h_Wr_yqq+aepbxLiWOZ], gS2TXhHpYaX4, amoa4Px)
		ifvRy8x := vKNCqP0 + j_mxLI
		j6Rv_Q := ifvRy8x + ljKfRQ
		uW0eIVjy[h_Wr_yqq+aepbxLiWOZ] = j6Rv_Q
		if ifvRy8x < j_mxLI || j6Rv_Q < ljKfRQ {
			vKNCqP0 = 1
		} else {
			vKNCqP0 = 0
		}
	}
	if vKNCqP0 != 0 {
		 /*line :1*/ omPlI3q_Zc(uW0eIVjy[:h_Wr_yqq], uW0eIVjy[h_Wr_yqq:], gS2TXhHpYaX4)
	} else {
		 /*line :1*/ copy(uW0eIVjy[:h_Wr_yqq], uW0eIVjy[h_Wr_yqq:])
	}
	return uW0eIVjy[:h_Wr_yqq]
}



func ugs2XUM(uW0eIVjy, pmEgen2 g3X9fa, h_Wr_yqq int) {
	if vKNCqP0 :=  /*line :1*/ vmnD6P1U(uW0eIVjy[0:h_Wr_yqq], uW0eIVjy, pmEgen2); vKNCqP0 != 0 {
		 /*line :1*/ aJRvTH(uW0eIVjy[h_Wr_yqq:h_Wr_yqq+h_Wr_yqq>>1], uW0eIVjy[h_Wr_yqq:], vKNCqP0)
	}
}


func gZWbMwG81aD7(uW0eIVjy, pmEgen2 g3X9fa, h_Wr_yqq int) {
	if vKNCqP0 :=  /*line :1*/ omPlI3q_Zc(uW0eIVjy[0:h_Wr_yqq], uW0eIVjy, pmEgen2); vKNCqP0 != 0 {
		 /*line :1*/ sQp2wtu(uW0eIVjy[h_Wr_yqq:h_Wr_yqq+h_Wr_yqq>>1], uW0eIVjy[h_Wr_yqq:], vKNCqP0)
	}
}




var gqua9d3M7q34 = 40	





func w1pgDrF6JlxT(uW0eIVjy, pmEgen2, oxFS5wa5 g3X9fa) {
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)

	if h_Wr_yqq&1 != 0 || h_Wr_yqq < gqua9d3M7q34 || h_Wr_yqq < 2 {
		 /*line :1*/ oGJ1hCL(uW0eIVjy, pmEgen2, oxFS5wa5)
		return
	}

	fx60ock3 := h_Wr_yqq >> 1
	dKR0NUYWi, _mZiyQO := pmEgen2[fx60ock3:], pmEgen2[0:fx60ock3]
	ihENsIZD, eYTbpqn := oxFS5wa5[fx60ock3:], oxFS5wa5[0:fx60ock3]

	 /*line :1*/ w1pgDrF6JlxT(uW0eIVjy, _mZiyQO, eYTbpqn)
	 /*line :1*/ w1pgDrF6JlxT(uW0eIVjy[h_Wr_yqq:], dKR0NUYWi, ihENsIZD)

	nR7KU4mGsdy := 1
	esxuiO := uW0eIVjy[2*h_Wr_yqq : 2*h_Wr_yqq+fx60ock3]
	if  /*line :1*/ omPlI3q_Zc(esxuiO, dKR0NUYWi, _mZiyQO) != 0 {
		nR7KU4mGsdy = -nR7KU4mGsdy
		 /*line :1*/ omPlI3q_Zc(esxuiO, _mZiyQO, dKR0NUYWi)
	}

	zUeGolyO := uW0eIVjy[2*h_Wr_yqq+fx60ock3 : 3*h_Wr_yqq]
	if  /*line :1*/ omPlI3q_Zc(zUeGolyO, eYTbpqn, ihENsIZD) != 0 {
		nR7KU4mGsdy = -nR7KU4mGsdy
		 /*line :1*/ omPlI3q_Zc(zUeGolyO, ihENsIZD, eYTbpqn)
	}

	djOsNJsLfH5U := uW0eIVjy[h_Wr_yqq*3:]
	 /*line :1*/ w1pgDrF6JlxT(djOsNJsLfH5U, esxuiO, zUeGolyO)

	vFx5p_cm := uW0eIVjy[h_Wr_yqq*4:]
	 /*line :1*/ copy(vFx5p_cm, uW0eIVjy[:h_Wr_yqq*2])

	 /*line :1*/ ugs2XUM(uW0eIVjy[fx60ock3:], vFx5p_cm, h_Wr_yqq)
	 /*line :1*/ ugs2XUM(uW0eIVjy[fx60ock3:], vFx5p_cm[h_Wr_yqq:], h_Wr_yqq)
	if nR7KU4mGsdy > 0 {
		 /*line :1*/ ugs2XUM(uW0eIVjy[fx60ock3:], djOsNJsLfH5U, h_Wr_yqq)
	} else {
		 /*line :1*/ gZWbMwG81aD7(uW0eIVjy[fx60ock3:], djOsNJsLfH5U, h_Wr_yqq)
	}
}







func ap52A4djK(pmEgen2, oxFS5wa5 g3X9fa) bool {
	return  /*line :1*/ cap(pmEgen2) > 0 &&  /*line :1*/ cap(oxFS5wa5) > 0 && &pmEgen2[0: /*line :1*/ cap(pmEgen2)][ /*line :1*/ cap(pmEgen2)-1] == &oxFS5wa5[0: /*line :1*/ cap(oxFS5wa5)][ /*line :1*/ cap(oxFS5wa5)-1]
}




func huhykmvX3UOx(uW0eIVjy, pmEgen2 g3X9fa, aepbxLiWOZ int) {
	if h_Wr_yqq :=  /*line :1*/ len(pmEgen2); h_Wr_yqq > 0 {
		if vKNCqP0 :=  /*line :1*/ vmnD6P1U(uW0eIVjy[aepbxLiWOZ:aepbxLiWOZ+h_Wr_yqq], uW0eIVjy[aepbxLiWOZ:], pmEgen2); vKNCqP0 != 0 {
			g77TAi := aepbxLiWOZ + h_Wr_yqq
			if g77TAi <  /*line :1*/ len(uW0eIVjy) {
				 /*line :1*/ aJRvTH(uW0eIVjy[g77TAi:], uW0eIVjy[g77TAi:], vKNCqP0)
			}
		}
	}
}

func mWLZAmlj(pmEgen2, oxFS5wa5 int) int {
	if pmEgen2 > oxFS5wa5 {
		return pmEgen2
	}
	return oxFS5wa5
}





func y0HfuCGB(h_Wr_yqq, xuC6BI int) int {
	aepbxLiWOZ :=  /*line :1*/ uint(0)
	for h_Wr_yqq > xuC6BI {
		h_Wr_yqq >>= 1
		aepbxLiWOZ++
	}
	return h_Wr_yqq << aepbxLiWOZ
}

func (uW0eIVjy g3X9fa) d22x6Ygoc80O(pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)

	switch {
	case gS2TXhHpYaX4 < h_Wr_yqq:
		return  /*line :1*/ uW0eIVjy.d22x6Ygoc80O(oxFS5wa5, pmEgen2)
	case gS2TXhHpYaX4 == 0 || h_Wr_yqq == 0:
		return uW0eIVjy[:0]
	case h_Wr_yqq == 1:
		return  /*line :1*/ uW0eIVjy.aTzUIOqJPczT(pmEgen2, oxFS5wa5[0], 0)
	}

	if  /*line :1*/ ap52A4djK(uW0eIVjy, pmEgen2) ||  /*line :1*/ ap52A4djK(uW0eIVjy, oxFS5wa5) {
		uW0eIVjy = nil
	}

	if h_Wr_yqq < gqua9d3M7q34 {
		uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4 + h_Wr_yqq)
		 /*line :1*/ oGJ1hCL(uW0eIVjy, pmEgen2, oxFS5wa5)
		return  /*line :1*/ uW0eIVjy.hitXcy5()
	}

	_DfINwXQ :=  /*line :1*/ y0HfuCGB(h_Wr_yqq, gqua9d3M7q34)

	_mZiyQO := pmEgen2[0:_DfINwXQ]
	eYTbpqn := oxFS5wa5[0:_DfINwXQ]
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0( /*line :1*/ mWLZAmlj(6*_DfINwXQ, gS2TXhHpYaX4+h_Wr_yqq))
	 /*line :1*/ w1pgDrF6JlxT(uW0eIVjy, _mZiyQO, eYTbpqn)
	uW0eIVjy = uW0eIVjy[0 : gS2TXhHpYaX4+h_Wr_yqq]
	 /*line :1*/ uW0eIVjy[2*_DfINwXQ:].uKEPQHWEut()

	if _DfINwXQ < h_Wr_yqq || gS2TXhHpYaX4 != h_Wr_yqq {
		cTKgk0 :=  /*line :1*/ aQEWrUmnl5I(3 * _DfINwXQ)
		amoa4Px := *cTKgk0

		_mZiyQO :=  /*line :1*/ _mZiyQO.hitXcy5()
		ihENsIZD := oxFS5wa5[_DfINwXQ:]
		amoa4Px =  /*line :1*/ amoa4Px.d22x6Ygoc80O(_mZiyQO, ihENsIZD)
		 /*line :1*/ huhykmvX3UOx(uW0eIVjy, amoa4Px, _DfINwXQ)

		eYTbpqn :=  /*line :1*/ eYTbpqn.hitXcy5()
		for aepbxLiWOZ := _DfINwXQ; aepbxLiWOZ <  /*line :1*/ len(pmEgen2); aepbxLiWOZ += _DfINwXQ {
			cNh2YNKJO := pmEgen2[aepbxLiWOZ:]
			if  /*line :1*/ len(cNh2YNKJO) > _DfINwXQ {
				cNh2YNKJO = cNh2YNKJO[:_DfINwXQ]
			}
			cNh2YNKJO =  /*line :1*/ cNh2YNKJO.hitXcy5()
			amoa4Px =  /*line :1*/ amoa4Px.d22x6Ygoc80O(cNh2YNKJO, eYTbpqn)
			 /*line :1*/ huhykmvX3UOx(uW0eIVjy, amoa4Px, aepbxLiWOZ)
			amoa4Px =  /*line :1*/ amoa4Px.d22x6Ygoc80O(cNh2YNKJO, ihENsIZD)
			 /*line :1*/ huhykmvX3UOx(uW0eIVjy, amoa4Px, aepbxLiWOZ+_DfINwXQ)
		}

		 /*line :1*/ zuavpO7v(cTKgk0)
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}





func yCpHraIVO0(uW0eIVjy, pmEgen2 g3X9fa) {
	h_Wr_yqq :=  /*line :1*/ len(pmEgen2)
	cTKgk0 :=  /*line :1*/ aQEWrUmnl5I(2 * h_Wr_yqq)
	amoa4Px := *cTKgk0
	 /*line :1*/ amoa4Px.uKEPQHWEut()
	uW0eIVjy[1], uW0eIVjy[0] =  /*line :1*/ mulWW(pmEgen2[0], pmEgen2[0])
	for aepbxLiWOZ := 1; aepbxLiWOZ < h_Wr_yqq; aepbxLiWOZ++ {
		s31g8J := pmEgen2[aepbxLiWOZ]

		uW0eIVjy[2*aepbxLiWOZ+1], uW0eIVjy[2*aepbxLiWOZ] =  /*line :1*/ mulWW(s31g8J, s31g8J)

		amoa4Px[2*aepbxLiWOZ] =  /*line :1*/ fklqwGXkLu(amoa4Px[aepbxLiWOZ:2*aepbxLiWOZ], pmEgen2[0:aepbxLiWOZ], s31g8J)
	}
	amoa4Px[2*h_Wr_yqq-1] =  /*line :1*/ vdi3DYM4mY(amoa4Px[1:2*h_Wr_yqq-1], amoa4Px[1:2*h_Wr_yqq-1], 1)
	 /*line :1*/ vmnD6P1U(uW0eIVjy, uW0eIVjy, amoa4Px)
	 /*line :1*/ zuavpO7v(cTKgk0)
}






func rVMokWMvAuTi(uW0eIVjy, pmEgen2 g3X9fa) {
	h_Wr_yqq :=  /*line :1*/ len(pmEgen2)

	if h_Wr_yqq&1 != 0 || h_Wr_yqq < mbonw69T || h_Wr_yqq < 2 {
		 /*line :1*/ yCpHraIVO0(uW0eIVjy[:2*h_Wr_yqq], pmEgen2)
		return
	}

	fx60ock3 := h_Wr_yqq >> 1
	dKR0NUYWi, _mZiyQO := pmEgen2[fx60ock3:], pmEgen2[0:fx60ock3]

	 /*line :1*/ rVMokWMvAuTi(uW0eIVjy, _mZiyQO)
	 /*line :1*/ rVMokWMvAuTi(uW0eIVjy[h_Wr_yqq:], dKR0NUYWi)

	esxuiO := uW0eIVjy[2*h_Wr_yqq : 2*h_Wr_yqq+fx60ock3]
	if  /*line :1*/ omPlI3q_Zc(esxuiO, dKR0NUYWi, _mZiyQO) != 0 {
		 /*line :1*/ omPlI3q_Zc(esxuiO, _mZiyQO, dKR0NUYWi)
	}

	djOsNJsLfH5U := uW0eIVjy[h_Wr_yqq*3:]
	 /*line :1*/ rVMokWMvAuTi(djOsNJsLfH5U, esxuiO)

	vFx5p_cm := uW0eIVjy[h_Wr_yqq*4:]
	 /*line :1*/ copy(vFx5p_cm, uW0eIVjy[:h_Wr_yqq*2])

	 /*line :1*/ ugs2XUM(uW0eIVjy[fx60ock3:], vFx5p_cm, h_Wr_yqq)
	 /*line :1*/ ugs2XUM(uW0eIVjy[fx60ock3:], vFx5p_cm[h_Wr_yqq:], h_Wr_yqq)
	 /*line :1*/ gZWbMwG81aD7(uW0eIVjy[fx60ock3:], djOsNJsLfH5U, h_Wr_yqq)
}




var _DL3ffHkkveP = 20	
var mbonw69T = 260	


func (uW0eIVjy g3X9fa) pJ61taToc(pmEgen2 g3X9fa) g3X9fa {
	h_Wr_yqq :=  /*line :1*/ len(pmEgen2)
	switch {
	case h_Wr_yqq == 0:
		return uW0eIVjy[:0]
	case h_Wr_yqq == 1:
		s31g8J := pmEgen2[0]
		uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(2)
		uW0eIVjy[1], uW0eIVjy[0] =  /*line :1*/ mulWW(s31g8J, s31g8J)
		return  /*line :1*/ uW0eIVjy.hitXcy5()
	}

	if  /*line :1*/ ap52A4djK(uW0eIVjy, pmEgen2) {
		uW0eIVjy = nil
	}

	if h_Wr_yqq < _DL3ffHkkveP {
		uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(2 * h_Wr_yqq)
		 /*line :1*/ oGJ1hCL(uW0eIVjy, pmEgen2, pmEgen2)
		return  /*line :1*/ uW0eIVjy.hitXcy5()
	}
	if h_Wr_yqq < mbonw69T {
		uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(2 * h_Wr_yqq)
		 /*line :1*/ yCpHraIVO0(uW0eIVjy, pmEgen2)
		return  /*line :1*/ uW0eIVjy.hitXcy5()
	}

	_DfINwXQ :=  /*line :1*/ y0HfuCGB(h_Wr_yqq, mbonw69T)

	_mZiyQO := pmEgen2[0:_DfINwXQ]
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0( /*line :1*/ mWLZAmlj(6*_DfINwXQ, 2*h_Wr_yqq))
	 /*line :1*/ rVMokWMvAuTi(uW0eIVjy, _mZiyQO)
	uW0eIVjy = uW0eIVjy[0 : 2*h_Wr_yqq]
	 /*line :1*/ uW0eIVjy[2*_DfINwXQ:].uKEPQHWEut()

	if _DfINwXQ < h_Wr_yqq {
		cTKgk0 :=  /*line :1*/ aQEWrUmnl5I(2 * _DfINwXQ)
		amoa4Px := *cTKgk0
		_mZiyQO :=  /*line :1*/ _mZiyQO.hitXcy5()
		dKR0NUYWi := pmEgen2[_DfINwXQ:]
		amoa4Px =  /*line :1*/ amoa4Px.d22x6Ygoc80O(_mZiyQO, dKR0NUYWi)
		 /*line :1*/ huhykmvX3UOx(uW0eIVjy, amoa4Px, _DfINwXQ)
		 /*line :1*/ huhykmvX3UOx(uW0eIVjy, amoa4Px, _DfINwXQ)
		amoa4Px =  /*line :1*/ amoa4Px.pJ61taToc(dKR0NUYWi)
		 /*line :1*/ huhykmvX3UOx(uW0eIVjy, amoa4Px, 2*_DfINwXQ)
		 /*line :1*/ zuavpO7v(cTKgk0)
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}



func (uW0eIVjy g3X9fa) x648hVuJhff(sza5atF, kWa1YrIHZo uint64) g3X9fa {
	switch {
	case sza5atF == 0:

		return  /*line :1*/ uW0eIVjy.glzD00sysLql(0)
	case sza5atF > kWa1YrIHZo:
		return  /*line :1*/ uW0eIVjy.glzD00sysLql(1)
	case sza5atF == kWa1YrIHZo:
		return  /*line :1*/ uW0eIVjy.glzD00sysLql(sza5atF)
	case sza5atF+1 == kWa1YrIHZo:
		return  /*line :1*/ uW0eIVjy.d22x6Ygoc80O( /*line :1*/ g3X9fa(nil).glzD00sysLql(sza5atF),  /*line :1*/ g3X9fa(nil).glzD00sysLql(kWa1YrIHZo))
	}
	gS2TXhHpYaX4 := (sza5atF + kWa1YrIHZo) / 2
	return  /*line :1*/ uW0eIVjy.d22x6Ygoc80O( /*line :1*/ g3X9fa(nil).x648hVuJhff(sza5atF, gS2TXhHpYaX4),  /*line :1*/ g3X9fa(nil).x648hVuJhff(gS2TXhHpYaX4+1, kWa1YrIHZo))
}



func aQEWrUmnl5I(h_Wr_yqq int) *g3X9fa {
	var uW0eIVjy *g3X9fa
	if iQKtjmO6 :=  /*line :1*/ yN__3JuT3A.Get(); iQKtjmO6 != nil {
		uW0eIVjy = iQKtjmO6.(*g3X9fa)
	}
	if uW0eIVjy == nil {
		uW0eIVjy =  /*line :1*/ new(g3X9fa)
	}
	*uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(h_Wr_yqq)
	if h_Wr_yqq > 0 {
		(*uW0eIVjy)[0] = 0xfedcb
	}
	return uW0eIVjy
}

func zuavpO7v(pmEgen2 *g3X9fa) {
	 /*line :1*/ yN__3JuT3A.Put(pmEgen2)
}

var yN__3JuT3A sync.OrP5FGPq



func (pmEgen2 g3X9fa) wZwJ3Y0() int {

	if aepbxLiWOZ :=  /*line :1*/ len(pmEgen2) - 1; aepbxLiWOZ >= 0 {

		hoZOuUT :=  /*line :1*/ uint(pmEgen2[aepbxLiWOZ])
		hoZOuUT |= hoZOuUT >> 1
		hoZOuUT |= hoZOuUT >> 2
		hoZOuUT |= hoZOuUT >> 4
		hoZOuUT |= hoZOuUT >> 8
		hoZOuUT |= hoZOuUT >> 16
		hoZOuUT |= hoZOuUT >> 16 >> 16
		return aepbxLiWOZ*_W +  /*line :1*/ bits.Len(hoZOuUT)
	}
	return 0
}



func (pmEgen2 g3X9fa) kDun6ak() uint {
	if  /*line :1*/ len(pmEgen2) == 0 {
		return 0
	}
	var aepbxLiWOZ uint
	for pmEgen2[aepbxLiWOZ] == 0 {
		aepbxLiWOZ++
	}

	return aepbxLiWOZ*_W +  /*line :1*/ uint( /*line :1*/ bits.AjW10JsK( /*line :1*/ uint(pmEgen2[aepbxLiWOZ])))
}


func (pmEgen2 g3X9fa) m23Kvd4wkaPL() (uint, bool) {
	var aepbxLiWOZ uint
	for pmEgen2[aepbxLiWOZ] == 0 {
		aepbxLiWOZ++
	}
	if aepbxLiWOZ ==  /*line :1*/ uint( /*line :1*/ len(pmEgen2))-1 && pmEgen2[aepbxLiWOZ]&(pmEgen2[aepbxLiWOZ]-1) == 0 {
		return aepbxLiWOZ*_W +  /*line :1*/ uint( /*line :1*/ bits.AjW10JsK( /*line :1*/ uint(pmEgen2[aepbxLiWOZ]))), true
	}
	return 0, false
}

func qeC76AZ94(pmEgen2, oxFS5wa5 g3X9fa) bool {
	return  /*line :1*/ len(pmEgen2) ==  /*line :1*/ len(oxFS5wa5) &&  /*line :1*/ len(pmEgen2) > 0 && &pmEgen2[0] == &oxFS5wa5[0]
}


func (uW0eIVjy g3X9fa) z839tk6CmDHT(pmEgen2 g3X9fa, nR7KU4mGsdy uint) g3X9fa {
	if nR7KU4mGsdy == 0 {
		if  /*line :1*/ qeC76AZ94(uW0eIVjy, pmEgen2) {
			return uW0eIVjy
		}
		if ! /*line :1*/ ap52A4djK(uW0eIVjy, pmEgen2) {
			return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
		}
	}

	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	if gS2TXhHpYaX4 == 0 {
		return uW0eIVjy[:0]
	}

	h_Wr_yqq := gS2TXhHpYaX4 +  /*line :1*/ int(nR7KU4mGsdy/_W)
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(h_Wr_yqq + 1)
	uW0eIVjy[h_Wr_yqq] =  /*line :1*/ vdi3DYM4mY(uW0eIVjy[h_Wr_yqq-gS2TXhHpYaX4:h_Wr_yqq], pmEgen2, nR7KU4mGsdy%_W)
	 /*line :1*/ uW0eIVjy[0 : h_Wr_yqq-gS2TXhHpYaX4].uKEPQHWEut()

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}


func (uW0eIVjy g3X9fa) qzOaHmR(pmEgen2 g3X9fa, nR7KU4mGsdy uint) g3X9fa {
	if nR7KU4mGsdy == 0 {
		if  /*line :1*/ qeC76AZ94(uW0eIVjy, pmEgen2) {
			return uW0eIVjy
		}
		if ! /*line :1*/ ap52A4djK(uW0eIVjy, pmEgen2) {
			return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
		}
	}

	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq := gS2TXhHpYaX4 -  /*line :1*/ int(nR7KU4mGsdy/_W)
	if h_Wr_yqq <= 0 {
		return uW0eIVjy[:0]
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(h_Wr_yqq)
	 /*line :1*/ t_ku9iRroNp(uW0eIVjy, pmEgen2[gS2TXhHpYaX4-h_Wr_yqq:], nR7KU4mGsdy%_W)

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}

func (uW0eIVjy g3X9fa) hUdaA0sgaAF(pmEgen2 g3X9fa, aepbxLiWOZ uint, kWa1YrIHZo uint) g3X9fa {
	g77TAi :=  /*line :1*/ int(aepbxLiWOZ / _W)
	gS2TXhHpYaX4 :=  /*line :1*/ VYe6D0(1) << (aepbxLiWOZ % _W)
	h_Wr_yqq :=  /*line :1*/ len(pmEgen2)
	switch kWa1YrIHZo {
	case 0:
		uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(h_Wr_yqq)
		 /*line :1*/ copy(uW0eIVjy, pmEgen2)
		if g77TAi >= h_Wr_yqq {

			return uW0eIVjy
		}
		uW0eIVjy[g77TAi] &^= gS2TXhHpYaX4
		return  /*line :1*/ uW0eIVjy.hitXcy5()
	case 1:
		if g77TAi >= h_Wr_yqq {
			uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(g77TAi + 1)
			 /*line :1*/ uW0eIVjy[h_Wr_yqq:].uKEPQHWEut()
		} else {
			uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(h_Wr_yqq)
		}
		 /*line :1*/ copy(uW0eIVjy, pmEgen2)
		uW0eIVjy[g77TAi] |= gS2TXhHpYaX4

		return uW0eIVjy
	}
	 /*line :1*/ panic("set bit is not 0 or 1")
}


func (pmEgen2 g3X9fa) bjRK8rI0NMn(aepbxLiWOZ uint) uint {
	g77TAi := aepbxLiWOZ / _W
	if g77TAi >=  /*line :1*/ uint( /*line :1*/ len(pmEgen2)) {
		return 0
	}

	return  /*line :1*/ uint(pmEgen2[g77TAi] >> (aepbxLiWOZ % _W) & 1)
}



func (pmEgen2 g3X9fa) cbgX3ksmdgax(aepbxLiWOZ uint) uint {
	g77TAi := aepbxLiWOZ / _W
	if g77TAi >=  /*line :1*/ uint( /*line :1*/ len(pmEgen2)) {
		if  /*line :1*/ len(pmEgen2) == 0 {
			return 0
		}
		return 1
	}

	for _, pmEgen2 := range pmEgen2[:g77TAi] {
		if pmEgen2 != 0 {
			return 1
		}
	}
	if pmEgen2[g77TAi]<<(_W-aepbxLiWOZ%_W) != 0 {
		return 1
	}
	return 0
}

func (uW0eIVjy g3X9fa) aqWAA8J(pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)
	if gS2TXhHpYaX4 > h_Wr_yqq {
		gS2TXhHpYaX4 = h_Wr_yqq
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4)
	for aepbxLiWOZ := 0; aepbxLiWOZ < gS2TXhHpYaX4; aepbxLiWOZ++ {
		uW0eIVjy[aepbxLiWOZ] = pmEgen2[aepbxLiWOZ] & oxFS5wa5[aepbxLiWOZ]
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}


func (uW0eIVjy g3X9fa) q7_lqtML4Gu(pmEgen2 g3X9fa, h_Wr_yqq uint) g3X9fa {
	uKQ8Eox := (h_Wr_yqq + _W - 1) / _W
	if  /*line :1*/ uint( /*line :1*/ len(pmEgen2)) < uKQ8Eox {
		return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
	}
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0( /*line :1*/ int(uKQ8Eox))
	 /*line :1*/ copy(uW0eIVjy, pmEgen2)
	if h_Wr_yqq%_W != 0 {
		uW0eIVjy[ /*line :1*/ len(uW0eIVjy)-1] &= 1<<(h_Wr_yqq%_W) - 1
	}
	return  /*line :1*/ uW0eIVjy.hitXcy5()
}

func (uW0eIVjy g3X9fa) csmID38W(pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)
	if h_Wr_yqq > gS2TXhHpYaX4 {
		h_Wr_yqq = gS2TXhHpYaX4
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4)
	for aepbxLiWOZ := 0; aepbxLiWOZ < h_Wr_yqq; aepbxLiWOZ++ {
		uW0eIVjy[aepbxLiWOZ] = pmEgen2[aepbxLiWOZ] &^ oxFS5wa5[aepbxLiWOZ]
	}
	 /*line :1*/ copy(uW0eIVjy[h_Wr_yqq:gS2TXhHpYaX4], pmEgen2[h_Wr_yqq:gS2TXhHpYaX4])

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}

func (uW0eIVjy g3X9fa) ycPkQaVsS(pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)
	nR7KU4mGsdy := pmEgen2
	if gS2TXhHpYaX4 < h_Wr_yqq {
		h_Wr_yqq, gS2TXhHpYaX4 = gS2TXhHpYaX4, h_Wr_yqq
		nR7KU4mGsdy = oxFS5wa5
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4)
	for aepbxLiWOZ := 0; aepbxLiWOZ < h_Wr_yqq; aepbxLiWOZ++ {
		uW0eIVjy[aepbxLiWOZ] = pmEgen2[aepbxLiWOZ] | oxFS5wa5[aepbxLiWOZ]
	}
	 /*line :1*/ copy(uW0eIVjy[h_Wr_yqq:gS2TXhHpYaX4], nR7KU4mGsdy[h_Wr_yqq:gS2TXhHpYaX4])

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}

func (uW0eIVjy g3X9fa) ubcVkBmt(pmEgen2, oxFS5wa5 g3X9fa) g3X9fa {
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2)
	h_Wr_yqq :=  /*line :1*/ len(oxFS5wa5)
	nR7KU4mGsdy := pmEgen2
	if gS2TXhHpYaX4 < h_Wr_yqq {
		h_Wr_yqq, gS2TXhHpYaX4 = gS2TXhHpYaX4, h_Wr_yqq
		nR7KU4mGsdy = oxFS5wa5
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(gS2TXhHpYaX4)
	for aepbxLiWOZ := 0; aepbxLiWOZ < h_Wr_yqq; aepbxLiWOZ++ {
		uW0eIVjy[aepbxLiWOZ] = pmEgen2[aepbxLiWOZ] ^ oxFS5wa5[aepbxLiWOZ]
	}
	 /*line :1*/ copy(uW0eIVjy[h_Wr_yqq:gS2TXhHpYaX4], nR7KU4mGsdy[h_Wr_yqq:gS2TXhHpYaX4])

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}



func (uW0eIVjy g3X9fa) calV9Cfo(biERyAySo *rand.YAxSls, ae_JRbwxg2A g3X9fa, h_Wr_yqq int) g3X9fa {
	if  /*line :1*/ ap52A4djK(uW0eIVjy, ae_JRbwxg2A) {
		uW0eIVjy = nil
	}
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0( /*line :1*/ len(ae_JRbwxg2A))

	jYtl72Qc :=  /*line :1*/ uint(h_Wr_yqq % _W)
	if jYtl72Qc == 0 {
		jYtl72Qc = _W
	}
	yWg_RX :=  /*line :1*/ VYe6D0((1 << jYtl72Qc) - 1)

	for {
		switch _W {
		case 32:
			for aepbxLiWOZ := range uW0eIVjy {
				uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0( /*line :1*/ biERyAySo.Uint32())
			}
		case 64:
			for aepbxLiWOZ := range uW0eIVjy {
				uW0eIVjy[aepbxLiWOZ] =  /*line :1*/ VYe6D0( /*line :1*/ biERyAySo.Uint32()) |  /*line :1*/ VYe6D0( /*line :1*/ biERyAySo.Uint32())<<32
			}
		default:
			 /*line :1*/ panic("unknown word size")
		}
		uW0eIVjy[ /*line :1*/ len(ae_JRbwxg2A)-1] &= yWg_RX
		if  /*line :1*/ uW0eIVjy.beaurQHiT(ae_JRbwxg2A) < 0 {
			break
		}
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}



func (uW0eIVjy g3X9fa) iGIIsMONlhq(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 g3X9fa, asP5KKvg bool) g3X9fa {
	if  /*line :1*/ ap52A4djK(uW0eIVjy, pmEgen2) ||  /*line :1*/ ap52A4djK(uW0eIVjy, oxFS5wa5) {

		uW0eIVjy = nil
	}

	if  /*line :1*/ len(gS2TXhHpYaX4) == 1 && gS2TXhHpYaX4[0] == 1 {
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(0)
	}

	if  /*line :1*/ len(oxFS5wa5) == 0 {
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(1)
	}

	if  /*line :1*/ len(pmEgen2) == 0 {
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(0)
	}

	if  /*line :1*/ len(pmEgen2) == 1 && pmEgen2[0] == 1 {
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(1)
	}

	if  /*line :1*/ len(oxFS5wa5) == 1 && oxFS5wa5[0] == 1 {
		if  /*line :1*/ len(gS2TXhHpYaX4) != 0 {
			return  /*line :1*/ uW0eIVjy.xf9ev_MB(pmEgen2, gS2TXhHpYaX4)
		}
		return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
	}

	if  /*line :1*/ len(gS2TXhHpYaX4) != 0 {

		uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0( /*line :1*/ len(gS2TXhHpYaX4))

		if  /*line :1*/ len(oxFS5wa5) > 1 && !asP5KKvg {
			if gS2TXhHpYaX4[0]&1 == 1 {
				return  /*line :1*/ uW0eIVjy.ac3srszwAxg(pmEgen2, oxFS5wa5, gS2TXhHpYaX4)
			}
			if rrhLpsyBtahy, hjqFin :=  /*line :1*/ gS2TXhHpYaX4.m23Kvd4wkaPL(); hjqFin {
				return  /*line :1*/ uW0eIVjy.bJy49Mi(pmEgen2, oxFS5wa5, rrhLpsyBtahy)
			}
			return  /*line :1*/ uW0eIVjy.pcaSjGyA1(pmEgen2, oxFS5wa5, gS2TXhHpYaX4)
		}
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
	iQKtjmO6 := oxFS5wa5[ /*line :1*/ len(oxFS5wa5)-1]
	qKn7aIMh5 :=  /*line :1*/ bUgMSHLa(iQKtjmO6) + 1
	iQKtjmO6 <<= qKn7aIMh5
	var yL0_p0wUnc g3X9fa

	const mask = 1 << (_W - 1)

	uKQ8Eox := _W -  /*line :1*/ int(qKn7aIMh5)
	
	
	var cMYvRzI, vFx5p_cm g3X9fa
	for g77TAi := 0; g77TAi < uKQ8Eox; g77TAi++ {
		cMYvRzI =  /*line :1*/ cMYvRzI.pJ61taToc(uW0eIVjy)
		cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI

		if iQKtjmO6&mask != 0 {
			cMYvRzI =  /*line :1*/ cMYvRzI.d22x6Ygoc80O(uW0eIVjy, pmEgen2)
			cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI
		}

		if  /*line :1*/ len(gS2TXhHpYaX4) != 0 {
			cMYvRzI, vFx5p_cm =  /*line :1*/ cMYvRzI.xOZxoLyl(vFx5p_cm, uW0eIVjy, gS2TXhHpYaX4)
			cMYvRzI, vFx5p_cm, yL0_p0wUnc, uW0eIVjy = yL0_p0wUnc, uW0eIVjy, cMYvRzI, vFx5p_cm
		}

		iQKtjmO6 <<= 1
	}

	for aepbxLiWOZ :=  /*line :1*/ len(oxFS5wa5) - 2; aepbxLiWOZ >= 0; aepbxLiWOZ-- {
		iQKtjmO6 = oxFS5wa5[aepbxLiWOZ]

		for g77TAi := 0; g77TAi < _W; g77TAi++ {
			cMYvRzI =  /*line :1*/ cMYvRzI.pJ61taToc(uW0eIVjy)
			cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI

			if iQKtjmO6&mask != 0 {
				cMYvRzI =  /*line :1*/ cMYvRzI.d22x6Ygoc80O(uW0eIVjy, pmEgen2)
				cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI
			}

			if  /*line :1*/ len(gS2TXhHpYaX4) != 0 {
				cMYvRzI, vFx5p_cm =  /*line :1*/ cMYvRzI.xOZxoLyl(vFx5p_cm, uW0eIVjy, gS2TXhHpYaX4)
				cMYvRzI, vFx5p_cm, yL0_p0wUnc, uW0eIVjy = yL0_p0wUnc, uW0eIVjy, cMYvRzI, vFx5p_cm
			}

			iQKtjmO6 <<= 1
		}
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}









func (uW0eIVjy g3X9fa) pcaSjGyA1(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 g3X9fa) g3X9fa {

	h_Wr_yqq :=  /*line :1*/ gS2TXhHpYaX4.kDun6ak()
	ggJ4S0UJLL :=  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(lG0t2KfqGwyP, h_Wr_yqq)
	tHbZl68b_ :=  /*line :1*/ g3X9fa(nil).qzOaHmR(gS2TXhHpYaX4, h_Wr_yqq)

	mQ0XlC6n_5ot :=  /*line :1*/ g3X9fa(nil).iGIIsMONlhq(pmEgen2, oxFS5wa5, ggJ4S0UJLL, false)
	aDGsAqP :=  /*line :1*/ g3X9fa(nil).iGIIsMONlhq(pmEgen2, oxFS5wa5, tHbZl68b_, false)

	uW0eIVjy =  /*line :1*/ uW0eIVjy.bj0nVC(aDGsAqP)

	mQ0XlC6n_5ot =  /*line :1*/ mQ0XlC6n_5ot.fnzh1oLA(mQ0XlC6n_5ot, aDGsAqP, h_Wr_yqq)

	oLtPH5eV9h :=  /*line :1*/ g3X9fa(nil).ipjmVbr2qtL(tHbZl68b_, ggJ4S0UJLL)
	aDGsAqP =  /*line :1*/ aDGsAqP.d22x6Ygoc80O(mQ0XlC6n_5ot, oLtPH5eV9h)
	aDGsAqP =  /*line :1*/ aDGsAqP.q7_lqtML4Gu(aDGsAqP, h_Wr_yqq)

	uW0eIVjy =  /*line :1*/ uW0eIVjy.jXv7nnhY_(uW0eIVjy,  /*line :1*/ mQ0XlC6n_5ot.d22x6Ygoc80O(aDGsAqP, tHbZl68b_))

	return uW0eIVjy
}



func (uW0eIVjy g3X9fa) bJy49Mi(pmEgen2, oxFS5wa5 g3X9fa, rrhLpsyBtahy uint) g3X9fa {
	if  /*line :1*/ len(oxFS5wa5) <= 1 {
		 /*line :1*/ panic("big: misuse of expNNWindowed")
	}
	if pmEgen2[0]&1 == 0 {

		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(0)
	}
	if rrhLpsyBtahy == 1 {
		return  /*line :1*/ uW0eIVjy.pSzoQ7PMW(1)
	}

	uKQ8Eox :=  /*line :1*/ int((rrhLpsyBtahy + _W - 1) / _W)
	kWFCWJWp2vR :=  /*line :1*/ aQEWrUmnl5I(uKQ8Eox)
	cMYvRzI := *kWFCWJWp2vR

	const n = 4
	
	var gknpj5L1zszQ [1 << n]*g3X9fa
	for aepbxLiWOZ := range gknpj5L1zszQ {
		gknpj5L1zszQ[aepbxLiWOZ] =  /*line :1*/ aQEWrUmnl5I(uKQ8Eox)
	}
	*gknpj5L1zszQ[0] =  /*line :1*/ gknpj5L1zszQ[0].bj0nVC(lG0t2KfqGwyP)
	*gknpj5L1zszQ[1] =  /*line :1*/ gknpj5L1zszQ[1].q7_lqtML4Gu(pmEgen2, rrhLpsyBtahy)
	for aepbxLiWOZ := 2; aepbxLiWOZ < 1<<n; aepbxLiWOZ += 2 {
		gMTZW90Q, djOsNJsLfH5U, yWhzWpOqK := gknpj5L1zszQ[aepbxLiWOZ/2], gknpj5L1zszQ[aepbxLiWOZ], gknpj5L1zszQ[aepbxLiWOZ+1]
		*djOsNJsLfH5U =  /*line :1*/ djOsNJsLfH5U.pJ61taToc(*gMTZW90Q)
		*djOsNJsLfH5U =  /*line :1*/ djOsNJsLfH5U.q7_lqtML4Gu(*djOsNJsLfH5U, rrhLpsyBtahy)
		*yWhzWpOqK =  /*line :1*/ yWhzWpOqK.d22x6Ygoc80O(*djOsNJsLfH5U, pmEgen2)
		*yWhzWpOqK =  /*line :1*/ yWhzWpOqK.q7_lqtML4Gu(*yWhzWpOqK, rrhLpsyBtahy)
	}

	aepbxLiWOZ :=  /*line :1*/ len(oxFS5wa5) - 1
	xW4D0F :=  /*line :1*/ int((rrhLpsyBtahy - 2) / _W)
	llUDie := ^ /*line :1*/ VYe6D0(0)
	if gCkzmZlhV := (rrhLpsyBtahy - 1) & (_W - 1); gCkzmZlhV != 0 {
		llUDie = (1 << gCkzmZlhV) - 1
	}
	if aepbxLiWOZ > xW4D0F {
		aepbxLiWOZ = xW4D0F
	}
	kuGztq6R := false
	uW0eIVjy =  /*line :1*/ uW0eIVjy.pSzoQ7PMW(1)
	for ; aepbxLiWOZ >= 0; aepbxLiWOZ-- {
		mCtoosf := oxFS5wa5[aepbxLiWOZ]
		if aepbxLiWOZ == xW4D0F {
			mCtoosf &= llUDie
		}
		for g77TAi := 0; g77TAi < _W; g77TAi += n {
			if kuGztq6R {

				cMYvRzI =  /*line :1*/ cMYvRzI.pJ61taToc(uW0eIVjy)
				cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI
				uW0eIVjy =  /*line :1*/ uW0eIVjy.q7_lqtML4Gu(uW0eIVjy, rrhLpsyBtahy)

				cMYvRzI =  /*line :1*/ cMYvRzI.pJ61taToc(uW0eIVjy)
				cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI
				uW0eIVjy =  /*line :1*/ uW0eIVjy.q7_lqtML4Gu(uW0eIVjy, rrhLpsyBtahy)

				cMYvRzI =  /*line :1*/ cMYvRzI.pJ61taToc(uW0eIVjy)
				cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI
				uW0eIVjy =  /*line :1*/ uW0eIVjy.q7_lqtML4Gu(uW0eIVjy, rrhLpsyBtahy)

				cMYvRzI =  /*line :1*/ cMYvRzI.pJ61taToc(uW0eIVjy)
				cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI
				uW0eIVjy =  /*line :1*/ uW0eIVjy.q7_lqtML4Gu(uW0eIVjy, rrhLpsyBtahy)
			}

			cMYvRzI =  /*line :1*/ cMYvRzI.d22x6Ygoc80O(uW0eIVjy, *gknpj5L1zszQ[mCtoosf>>(_W-n)])
			cMYvRzI, uW0eIVjy = uW0eIVjy, cMYvRzI
			uW0eIVjy =  /*line :1*/ uW0eIVjy.q7_lqtML4Gu(uW0eIVjy, rrhLpsyBtahy)

			mCtoosf <<= n
			kuGztq6R = true
		}
	}

	*kWFCWJWp2vR = cMYvRzI
	 /*line :1*/ zuavpO7v(kWFCWJWp2vR)
	for aepbxLiWOZ := range gknpj5L1zszQ {
		 /*line :1*/ zuavpO7v(gknpj5L1zszQ[aepbxLiWOZ])
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}



func (uW0eIVjy g3X9fa) ac3srszwAxg(pmEgen2, oxFS5wa5, gS2TXhHpYaX4 g3X9fa) g3X9fa {
	wrtADGWkbODY :=  /*line :1*/ len(gS2TXhHpYaX4)

	if  /*line :1*/ len(pmEgen2) > wrtADGWkbODY {
		_, pmEgen2 =  /*line :1*/ g3X9fa(nil).xOZxoLyl(nil, pmEgen2, gS2TXhHpYaX4)

	}
	if  /*line :1*/ len(pmEgen2) < wrtADGWkbODY {
		mdBq1Cc :=  /*line :1*/ make(g3X9fa, wrtADGWkbODY)
		 /*line :1*/ copy(mdBq1Cc, pmEgen2)
		pmEgen2 = mdBq1Cc
	}

	h3HhLgkdmQXN := 2 - gS2TXhHpYaX4[0]
	amoa4Px := gS2TXhHpYaX4[0] - 1
	for aepbxLiWOZ := 1; aepbxLiWOZ < _W; aepbxLiWOZ <<= 1 {
		amoa4Px *= amoa4Px
		h3HhLgkdmQXN *= (amoa4Px + 1)
	}
	h3HhLgkdmQXN = -h3HhLgkdmQXN

	W190NFsOc :=  /*line :1*/ g3X9fa(nil).pSzoQ7PMW(1)
	cMYvRzI :=  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(W190NFsOc,  /*line :1*/ uint(2*wrtADGWkbODY*_W))
	_, W190NFsOc =  /*line :1*/ g3X9fa(nil).xOZxoLyl(W190NFsOc, cMYvRzI, gS2TXhHpYaX4)
	if  /*line :1*/ len(W190NFsOc) < wrtADGWkbODY {
		cMYvRzI =  /*line :1*/ cMYvRzI.gS9bqHpKAJ0(wrtADGWkbODY)
		 /*line :1*/ copy(cMYvRzI, W190NFsOc)
		W190NFsOc = cMYvRzI
	}

	tI0S7pUq :=  /*line :1*/ make(g3X9fa, wrtADGWkbODY)
	tI0S7pUq[0] = 1

	const n = 4
	
	var gknpj5L1zszQ [1 << n]g3X9fa
	gknpj5L1zszQ[0] =  /*line :1*/ gknpj5L1zszQ[0].vuR24NZ(tI0S7pUq, W190NFsOc, gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
	gknpj5L1zszQ[1] =  /*line :1*/ gknpj5L1zszQ[1].vuR24NZ(pmEgen2, W190NFsOc, gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
	for aepbxLiWOZ := 2; aepbxLiWOZ < 1<<n; aepbxLiWOZ++ {
		gknpj5L1zszQ[aepbxLiWOZ] =  /*line :1*/ gknpj5L1zszQ[aepbxLiWOZ].vuR24NZ(gknpj5L1zszQ[aepbxLiWOZ-1], gknpj5L1zszQ[1], gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(wrtADGWkbODY)
	 /*line :1*/ copy(uW0eIVjy, gknpj5L1zszQ[0])

	cMYvRzI =  /*line :1*/ cMYvRzI.gS9bqHpKAJ0(wrtADGWkbODY)

	for aepbxLiWOZ :=  /*line :1*/ len(oxFS5wa5) - 1; aepbxLiWOZ >= 0; aepbxLiWOZ-- {
		mCtoosf := oxFS5wa5[aepbxLiWOZ]
		for g77TAi := 0; g77TAi < _W; g77TAi += n {
			if aepbxLiWOZ !=  /*line :1*/ len(oxFS5wa5)-1 || g77TAi != 0 {
				cMYvRzI =  /*line :1*/ cMYvRzI.vuR24NZ(uW0eIVjy, uW0eIVjy, gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
				uW0eIVjy =  /*line :1*/ uW0eIVjy.vuR24NZ(cMYvRzI, cMYvRzI, gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
				cMYvRzI =  /*line :1*/ cMYvRzI.vuR24NZ(uW0eIVjy, uW0eIVjy, gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
				uW0eIVjy =  /*line :1*/ uW0eIVjy.vuR24NZ(cMYvRzI, cMYvRzI, gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
			}
			cMYvRzI =  /*line :1*/ cMYvRzI.vuR24NZ(uW0eIVjy, gknpj5L1zszQ[mCtoosf>>(_W-n)], gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)
			uW0eIVjy, cMYvRzI = cMYvRzI, uW0eIVjy
			mCtoosf <<= n
		}
	}

	cMYvRzI =  /*line :1*/ cMYvRzI.vuR24NZ(uW0eIVjy, tI0S7pUq, gS2TXhHpYaX4, h3HhLgkdmQXN, wrtADGWkbODY)

	if  /*line :1*/ cMYvRzI.beaurQHiT(gS2TXhHpYaX4) >= 0 {

		cMYvRzI =  /*line :1*/ cMYvRzI.lt4VKILNCo(cMYvRzI, gS2TXhHpYaX4)
		if  /*line :1*/ cMYvRzI.beaurQHiT(gS2TXhHpYaX4) >= 0 {
			_, cMYvRzI =  /*line :1*/ g3X9fa(nil).xOZxoLyl(nil, cMYvRzI, gS2TXhHpYaX4)
		}
	}

	return  /*line :1*/ cMYvRzI.hitXcy5()
}





func (uW0eIVjy g3X9fa) qonqPl(c3Zu4RY []byte) (aepbxLiWOZ int) {

	aepbxLiWOZ =  /*line :1*/ len(c3Zu4RY)
	for _, s31g8J := range uW0eIVjy {
		for g77TAi := 0; g77TAi < _S; g77TAi++ {
			aepbxLiWOZ--
			if aepbxLiWOZ >= 0 {
				c3Zu4RY[aepbxLiWOZ] =  /*line :1*/ byte(s31g8J)
			} else if  /*line :1*/ byte(s31g8J) != 0 {
				 /*line :1*/ panic("math/big: buffer too small to fit value")
			}
			s31g8J >>= 8
		}
	}

	if aepbxLiWOZ < 0 {
		aepbxLiWOZ = 0
	}
	for aepbxLiWOZ <  /*line :1*/ len(c3Zu4RY) && c3Zu4RY[aepbxLiWOZ] == 0 {
		aepbxLiWOZ++
	}

	return
}


func zdjEbAwFIjd(c3Zu4RY []byte) VYe6D0 {
	if _W == 64 {
		return  /*line :1*/ VYe6D0( /*line :1*/ binary.FOcKBq6.Uint64(c3Zu4RY))
	}
	return  /*line :1*/ VYe6D0( /*line :1*/ binary.FOcKBq6.Uint32(c3Zu4RY))
}



func (uW0eIVjy g3X9fa) vbFROX(c3Zu4RY []byte) g3X9fa {
	uW0eIVjy =  /*line :1*/ uW0eIVjy.gS9bqHpKAJ0(( /*line :1*/ len(c3Zu4RY) + _S - 1) / _S)

	aepbxLiWOZ :=  /*line :1*/ len(c3Zu4RY)
	for _DfINwXQ := 0; aepbxLiWOZ >= _S; _DfINwXQ++ {
		uW0eIVjy[_DfINwXQ] =  /*line :1*/ zdjEbAwFIjd(c3Zu4RY[aepbxLiWOZ-_S : aepbxLiWOZ])
		aepbxLiWOZ -= _S
	}
	if aepbxLiWOZ > 0 {
		var s31g8J VYe6D0
		for nR7KU4mGsdy :=  /*line :1*/ uint(0); aepbxLiWOZ > 0; nR7KU4mGsdy += 8 {
			s31g8J |=  /*line :1*/ VYe6D0(c3Zu4RY[aepbxLiWOZ-1]) << nR7KU4mGsdy
			aepbxLiWOZ--
		}
		uW0eIVjy[ /*line :1*/ len(uW0eIVjy)-1] = s31g8J
	}

	return  /*line :1*/ uW0eIVjy.hitXcy5()
}


func (uW0eIVjy g3X9fa) vf9TNsbnLN(pmEgen2 g3X9fa) g3X9fa {
	if  /*line :1*/ pmEgen2.beaurQHiT(lG0t2KfqGwyP) <= 0 {
		return  /*line :1*/ uW0eIVjy.bj0nVC(pmEgen2)
	}
	if  /*line :1*/ ap52A4djK(uW0eIVjy, pmEgen2) {
		uW0eIVjy = nil
	}

	
	
	
	
	
	var mQ0XlC6n_5ot, aDGsAqP g3X9fa
	mQ0XlC6n_5ot = uW0eIVjy
	mQ0XlC6n_5ot =  /*line :1*/ mQ0XlC6n_5ot.glzD00sysLql(1)
	mQ0XlC6n_5ot =  /*line :1*/ mQ0XlC6n_5ot.z839tk6CmDHT(mQ0XlC6n_5ot,  /*line :1*/ uint( /*line :1*/ pmEgen2.wZwJ3Y0()+1)/2)
	for h_Wr_yqq := 0; ; h_Wr_yqq++ {
		aDGsAqP, _ =  /*line :1*/ aDGsAqP.xOZxoLyl(nil, pmEgen2, mQ0XlC6n_5ot)
		aDGsAqP =  /*line :1*/ aDGsAqP.jXv7nnhY_(aDGsAqP, mQ0XlC6n_5ot)
		aDGsAqP =  /*line :1*/ aDGsAqP.qzOaHmR(aDGsAqP, 1)
		if  /*line :1*/ aDGsAqP.beaurQHiT(mQ0XlC6n_5ot) >= 0 {

			if h_Wr_yqq&1 == 0 {
				return mQ0XlC6n_5ot
			}
			return  /*line :1*/ uW0eIVjy.bj0nVC(mQ0XlC6n_5ot)
		}
		mQ0XlC6n_5ot, aDGsAqP = aDGsAqP, mQ0XlC6n_5ot
	}
}


func (uW0eIVjy g3X9fa) fnzh1oLA(pmEgen2, oxFS5wa5 g3X9fa, h_Wr_yqq uint) g3X9fa {
	if  /*line :1*/ uint( /*line :1*/ pmEgen2.wZwJ3Y0()) > h_Wr_yqq {
		if  /*line :1*/ ap52A4djK(uW0eIVjy, pmEgen2) {

			pmEgen2 =  /*line :1*/ pmEgen2.q7_lqtML4Gu(pmEgen2, h_Wr_yqq)
		} else {
			pmEgen2 =  /*line :1*/ g3X9fa(nil).q7_lqtML4Gu(pmEgen2, h_Wr_yqq)
		}
	}
	if  /*line :1*/ uint( /*line :1*/ oxFS5wa5.wZwJ3Y0()) > h_Wr_yqq {
		if  /*line :1*/ ap52A4djK(uW0eIVjy, oxFS5wa5) {

			oxFS5wa5 =  /*line :1*/ oxFS5wa5.q7_lqtML4Gu(oxFS5wa5, h_Wr_yqq)
		} else {
			oxFS5wa5 =  /*line :1*/ g3X9fa(nil).q7_lqtML4Gu(oxFS5wa5, h_Wr_yqq)
		}
	}
	if  /*line :1*/ pmEgen2.beaurQHiT(oxFS5wa5) >= 0 {
		return  /*line :1*/ uW0eIVjy.lt4VKILNCo(pmEgen2, oxFS5wa5)
	}

	uW0eIVjy =  /*line :1*/ uW0eIVjy.lt4VKILNCo(oxFS5wa5, pmEgen2)
	for  /*line :1*/ uint( /*line :1*/ len(uW0eIVjy))*_W < h_Wr_yqq {
		uW0eIVjy =  /*line :1*/ append(uW0eIVjy, 0)
	}
	for aepbxLiWOZ := range uW0eIVjy {
		uW0eIVjy[aepbxLiWOZ] = ^uW0eIVjy[aepbxLiWOZ]
	}
	uW0eIVjy =  /*line :1*/ uW0eIVjy.q7_lqtML4Gu(uW0eIVjy, h_Wr_yqq)
	return  /*line :1*/ uW0eIVjy.jXv7nnhY_(uW0eIVjy, lG0t2KfqGwyP)
}
