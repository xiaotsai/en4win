//line :1
package big

import (
	errors "iAHoxjmM"
	fmt "kFsoCfy5zWG"
	io "xI9ai2djaJ2"
	math "math"
	bits "math/bits"
	sync "sync"
)

const digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


const MaxBase = 10 + ('z' - 'a' + 1) + ('Z' - 'A' + 1)
const maxBaseSmall = 10 + ('z' - 'a' + 1)





func tXBgI4BZgB(kWa1YrIHZo VYe6D0) (djOsNJsLfH5U VYe6D0, h_Wr_yqq int) {
	djOsNJsLfH5U, h_Wr_yqq = kWa1YrIHZo, 1
	for mWLZAmlj := _M / kWa1YrIHZo; djOsNJsLfH5U <= mWLZAmlj; {

		djOsNJsLfH5U *= kWa1YrIHZo
		h_Wr_yqq++
	}

	return
}


func jtC3eCX0vFd(pmEgen2 VYe6D0, h_Wr_yqq int) (djOsNJsLfH5U VYe6D0) {

	djOsNJsLfH5U = 1
	for h_Wr_yqq > 0 {
		if h_Wr_yqq&1 != 0 {
			djOsNJsLfH5U *= pmEgen2
		}
		pmEgen2 *= pmEgen2
		h_Wr_yqq >>= 1
	}
	return
}


var (
	tdXYCwgkg	=  /*line :1*/ errors.Su6g6hRoi9X("number has no digits")
	l9euxO	=  /*line :1*/ errors.Su6g6hRoi9X("'_' must separate successive digits")
)













































func (uW0eIVjy g3X9fa) h9c5N0bF(vFx5p_cm io.HSeaoj16Wq0, weSJS4i4lm int, uzsRp3rMl bool) (iV_tpcVWZqaz g3X9fa, kWa1YrIHZo, qWfhyOCfC0By int, gDonrwI8VP50 error) {

	qjuS_0b8 := weSJS4i4lm == 0 ||
		!uzsRp3rMl && 2 <= weSJS4i4lm && weSJS4i4lm <= MaxBase ||
		uzsRp3rMl && (weSJS4i4lm == 2 || weSJS4i4lm == 8 || weSJS4i4lm == 10 || weSJS4i4lm == 16)
	if !qjuS_0b8 {
		 /*line :1*/ panic( /*line :1*/ fmt.IBsS3Oc("invalid number base %d", weSJS4i4lm))
	}

	iCClg09zMOt := '.'
	j9gIA7FO := false

	qvA6MnWV, gDonrwI8VP50 :=  /*line :1*/ vFx5p_cm.ReadByte()

	kWa1YrIHZo, vpvKXvnZ0CwR := weSJS4i4lm, 0
	if weSJS4i4lm == 0 {

		kWa1YrIHZo = 10
		if gDonrwI8VP50 == nil && qvA6MnWV == '0' {
			iCClg09zMOt = '0'
			qWfhyOCfC0By = 1
			qvA6MnWV, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte()
			if gDonrwI8VP50 == nil {

				switch qvA6MnWV {
				case 'b', 'B':
					kWa1YrIHZo, vpvKXvnZ0CwR = 2, 'b'
				case 'o', 'O':
					kWa1YrIHZo, vpvKXvnZ0CwR = 8, 'o'
				case 'x', 'X':
					kWa1YrIHZo, vpvKXvnZ0CwR = 16, 'x'
				default:
					if !uzsRp3rMl {
						kWa1YrIHZo, vpvKXvnZ0CwR = 8, '0'
					}
				}
				if vpvKXvnZ0CwR != 0 {
					qWfhyOCfC0By = 0
					if vpvKXvnZ0CwR != '0' {
						qvA6MnWV, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte()
					}
				}
			}
		}
	}

	uW0eIVjy = uW0eIVjy[:0]
	opx3YyEz3 :=  /*line :1*/ VYe6D0(kWa1YrIHZo)
	_v8jcap, h_Wr_yqq :=  /*line :1*/ tXBgI4BZgB(opx3YyEz3)
	jSBa5ZlAhUtr :=  /*line :1*/ VYe6D0(0)
	aepbxLiWOZ := 0
	qW0nYQVPj := -1
	for gDonrwI8VP50 == nil {
		if qvA6MnWV == '.' && uzsRp3rMl {
			uzsRp3rMl = false
			if iCClg09zMOt == '_' {
				j9gIA7FO = true
			}
			iCClg09zMOt = '.'
			qW0nYQVPj = qWfhyOCfC0By
		} else if qvA6MnWV == '_' && weSJS4i4lm == 0 {
			if iCClg09zMOt != '0' {
				j9gIA7FO = true
			}
			iCClg09zMOt = '_'
		} else {
			
			var pStBhY VYe6D0
			switch {
			case '0' <= qvA6MnWV && qvA6MnWV <= '9':
				pStBhY =  /*line :1*/ VYe6D0(qvA6MnWV - '0')
			case 'a' <= qvA6MnWV && qvA6MnWV <= 'z':
				pStBhY =  /*line :1*/ VYe6D0(qvA6MnWV - 'a' + 10)
			case 'A' <= qvA6MnWV && qvA6MnWV <= 'Z':
				if kWa1YrIHZo <= maxBaseSmall {
					pStBhY =  /*line :1*/ VYe6D0(qvA6MnWV - 'A' + 10)
				} else {
					pStBhY =  /*line :1*/ VYe6D0(qvA6MnWV - 'A' + maxBaseSmall)
				}
			default:
				pStBhY = MaxBase + 1
			}
			if pStBhY >= opx3YyEz3 {
				 /*line :1*/ vFx5p_cm.UnreadByte()
				break
			}
			iCClg09zMOt = '0'
			qWfhyOCfC0By++

			jSBa5ZlAhUtr = jSBa5ZlAhUtr*opx3YyEz3 + pStBhY
			aepbxLiWOZ++

			if aepbxLiWOZ == h_Wr_yqq {
				uW0eIVjy =  /*line :1*/ uW0eIVjy.aTzUIOqJPczT(uW0eIVjy, _v8jcap, jSBa5ZlAhUtr)
				jSBa5ZlAhUtr = 0
				aepbxLiWOZ = 0
			}
		}

		qvA6MnWV, gDonrwI8VP50 =  /*line :1*/ vFx5p_cm.ReadByte()
	}

	if gDonrwI8VP50 == io.K5Sqco {
		gDonrwI8VP50 = nil
	}

	if gDonrwI8VP50 == nil && (j9gIA7FO || iCClg09zMOt == '_') {
		gDonrwI8VP50 = l9euxO
	}

	if qWfhyOCfC0By == 0 {

		if vpvKXvnZ0CwR == '0' {

			return uW0eIVjy[:0], 10, 1, gDonrwI8VP50
		}
		gDonrwI8VP50 = tdXYCwgkg
	}

	if aepbxLiWOZ > 0 {
		uW0eIVjy =  /*line :1*/ uW0eIVjy.aTzUIOqJPczT(uW0eIVjy,  /*line :1*/ jtC3eCX0vFd(opx3YyEz3, aepbxLiWOZ), jSBa5ZlAhUtr)
	}
	iV_tpcVWZqaz =  /*line :1*/ uW0eIVjy.hitXcy5()

	if qW0nYQVPj >= 0 {

		qWfhyOCfC0By = qW0nYQVPj - qWfhyOCfC0By
	}

	return
}



func (pmEgen2 g3X9fa) g3J8Fbn4Ip(weSJS4i4lm int) []byte {
	return  /*line :1*/ pmEgen2.sgVtOtra(false, weSJS4i4lm)
}


func (pmEgen2 g3X9fa) sgVtOtra(s_9Ih_ bool, weSJS4i4lm int) []byte {
	if weSJS4i4lm < 2 || weSJS4i4lm > MaxBase {
		 /*line :1*/ panic("invalid base")
	}

	if  /*line :1*/ len(pmEgen2) == 0 {
		return [] /*line :1*/ byte("0")
	}

	aepbxLiWOZ :=  /*line :1*/ int( /*line :1*/ float64( /*line :1*/ pmEgen2.wZwJ3Y0())/ /*line :1*/ math.JwAjQi( /*line :1*/ float64(weSJS4i4lm))) + 1
	if s_9Ih_ {
		aepbxLiWOZ++
	}
	nR7KU4mGsdy :=  /*line :1*/ make([]byte, aepbxLiWOZ)

	if kWa1YrIHZo :=  /*line :1*/ VYe6D0(weSJS4i4lm); kWa1YrIHZo == kWa1YrIHZo&-kWa1YrIHZo {

		qKn7aIMh5 :=  /*line :1*/ uint( /*line :1*/ bits.AjW10JsK( /*line :1*/ uint(kWa1YrIHZo)))
		yWg_RX :=  /*line :1*/ VYe6D0(1<<qKn7aIMh5 - 1)
		uKQ8Eox := pmEgen2[0]
		tYy15uELw1 :=  /*line :1*/ uint(_W)

		for _DfINwXQ := 1; _DfINwXQ <  /*line :1*/ len(pmEgen2); _DfINwXQ++ {

			for tYy15uELw1 >= qKn7aIMh5 {
				aepbxLiWOZ--
				nR7KU4mGsdy[aepbxLiWOZ] = digits[uKQ8Eox&yWg_RX]
				uKQ8Eox >>= qKn7aIMh5
				tYy15uELw1 -= qKn7aIMh5
			}

			if tYy15uELw1 == 0 {

				uKQ8Eox = pmEgen2[_DfINwXQ]
				tYy15uELw1 = _W
			} else {

				uKQ8Eox |= pmEgen2[_DfINwXQ] << tYy15uELw1
				aepbxLiWOZ--
				nR7KU4mGsdy[aepbxLiWOZ] = digits[uKQ8Eox&yWg_RX]

				uKQ8Eox = pmEgen2[_DfINwXQ] >> (qKn7aIMh5 - tYy15uELw1)
				tYy15uELw1 = _W - (qKn7aIMh5 - tYy15uELw1)
			}
		}

		for uKQ8Eox != 0 {
			aepbxLiWOZ--
			nR7KU4mGsdy[aepbxLiWOZ] = digits[uKQ8Eox&yWg_RX]
			uKQ8Eox >>= qKn7aIMh5
		}

	} else {
		iIGcv9, vEkdraJ :=  /*line :1*/ tXBgI4BZgB(kWa1YrIHZo)

		djVSllzO5 :=  /*line :1*/ xuyVV9r9G( /*line :1*/ len(pmEgen2), kWa1YrIHZo, vEkdraJ, iIGcv9)

		yL0_p0wUnc :=  /*line :1*/ g3X9fa(nil).bj0nVC(pmEgen2)

		 /*line :1*/ yL0_p0wUnc.lBTFsSwRHg(nR7KU4mGsdy, kWa1YrIHZo, vEkdraJ, iIGcv9, djVSllzO5)

		aepbxLiWOZ = 0
		for nR7KU4mGsdy[aepbxLiWOZ] == '0' {
			aepbxLiWOZ++
		}
	}

	if s_9Ih_ {
		aepbxLiWOZ--
		nR7KU4mGsdy[aepbxLiWOZ] = '-'
	}

	return nR7KU4mGsdy[aepbxLiWOZ:]
}
















func (yL0_p0wUnc g3X9fa) lBTFsSwRHg(nR7KU4mGsdy []byte, kWa1YrIHZo VYe6D0, vEkdraJ int, iIGcv9 VYe6D0, djVSllzO5 []bi6eHE2RO) {

	if djVSllzO5 != nil {
		
		var vFx5p_cm g3X9fa
		pc6gn1FfNE :=  /*line :1*/ len(djVSllzO5) - 1
		for  /*line :1*/ len(yL0_p0wUnc) > paETBQw {

			gxI7iH :=  /*line :1*/ yL0_p0wUnc.wZwJ3Y0()
			_Nnv2mkj := gxI7iH >> 1
			for pc6gn1FfNE > 0 && djVSllzO5[pc6gn1FfNE-1].cev83XwLwKz > _Nnv2mkj {
				pc6gn1FfNE--
			}
			if djVSllzO5[pc6gn1FfNE].cev83XwLwKz >= gxI7iH &&  /*line :1*/ djVSllzO5[pc6gn1FfNE].rrnbR4.beaurQHiT(yL0_p0wUnc) >= 0 {
				pc6gn1FfNE--
				if pc6gn1FfNE < 0 {
					 /*line :1*/ panic("internal inconsistency")
				}
			}

			yL0_p0wUnc, vFx5p_cm =  /*line :1*/ yL0_p0wUnc.xOZxoLyl(vFx5p_cm, yL0_p0wUnc, djVSllzO5[pc6gn1FfNE].rrnbR4)

			dwiMOUz3iD5U :=  /*line :1*/ len(nR7KU4mGsdy) - djVSllzO5[pc6gn1FfNE].pMhnIOjkP
			 /*line :1*/ vFx5p_cm.lBTFsSwRHg(nR7KU4mGsdy[dwiMOUz3iD5U:], kWa1YrIHZo, vEkdraJ, iIGcv9, djVSllzO5[0:pc6gn1FfNE])
			nR7KU4mGsdy = nR7KU4mGsdy[:dwiMOUz3iD5U]
		}
	}

	aepbxLiWOZ :=  /*line :1*/ len(nR7KU4mGsdy)
	var vFx5p_cm VYe6D0
	if kWa1YrIHZo == 10 {

		for  /*line :1*/ len(yL0_p0wUnc) > 0 {

			yL0_p0wUnc, vFx5p_cm =  /*line :1*/ yL0_p0wUnc.htGqJRPj(yL0_p0wUnc, iIGcv9)
			for g77TAi := 0; g77TAi < vEkdraJ && aepbxLiWOZ > 0; g77TAi++ {
				aepbxLiWOZ--

				amoa4Px := vFx5p_cm / 10
				nR7KU4mGsdy[aepbxLiWOZ] = '0' +  /*line :1*/ byte(vFx5p_cm-amoa4Px*10)
				vFx5p_cm = amoa4Px
			}
		}
	} else {
		for  /*line :1*/ len(yL0_p0wUnc) > 0 {

			yL0_p0wUnc, vFx5p_cm =  /*line :1*/ yL0_p0wUnc.htGqJRPj(yL0_p0wUnc, iIGcv9)
			for g77TAi := 0; g77TAi < vEkdraJ && aepbxLiWOZ > 0; g77TAi++ {
				aepbxLiWOZ--
				nR7KU4mGsdy[aepbxLiWOZ] = digits[vFx5p_cm%kWa1YrIHZo]
				vFx5p_cm /= kWa1YrIHZo
			}
		}
	}

	for aepbxLiWOZ > 0 {
		aepbxLiWOZ--
		nR7KU4mGsdy[aepbxLiWOZ] = '0'
	}
}






var paETBQw int = 8	

type bi6eHE2RO struct {
	rrnbR4	g3X9fa	
	cev83XwLwKz	int	
	pMhnIOjkP	int	
}

var tFHi4yBEr struct {
	sync.DIRWe11KYlYa
	a5bigni_xD	[64]bi6eHE2RO	
}


func (uW0eIVjy g3X9fa) ea7MDYcwtuzk(pmEgen2, oxFS5wa5 VYe6D0) g3X9fa {
	return  /*line :1*/ uW0eIVjy.iGIIsMONlhq( /*line :1*/ g3X9fa(nil).pSzoQ7PMW(pmEgen2),  /*line :1*/ g3X9fa(nil).pSzoQ7PMW(oxFS5wa5), nil, false)
}


func xuyVV9r9G(gS2TXhHpYaX4 int, kWa1YrIHZo VYe6D0, vEkdraJ int, iIGcv9 VYe6D0) []bi6eHE2RO {

	if paETBQw == 0 || gS2TXhHpYaX4 <= paETBQw {
		return nil
	}

	_DfINwXQ := 1
	for f9mTywV1nCq := paETBQw; f9mTywV1nCq < gS2TXhHpYaX4>>1 && _DfINwXQ <  /*line :1*/ len(tFHi4yBEr.a5bigni_xD); f9mTywV1nCq <<= 1 {
		_DfINwXQ++
	}

	
	var djVSllzO5 []bi6eHE2RO	
	if kWa1YrIHZo == 10 {
		 /*line :1*/ tFHi4yBEr.Lock()
		djVSllzO5 = tFHi4yBEr.a5bigni_xD[0:_DfINwXQ]
	} else {
		djVSllzO5 =  /*line :1*/ make([]bi6eHE2RO, _DfINwXQ)
	}

	if djVSllzO5[_DfINwXQ-1].pMhnIOjkP == 0 {
		
		var sGlZ9QmSSqFV g3X9fa
		for aepbxLiWOZ := 0; aepbxLiWOZ < _DfINwXQ; aepbxLiWOZ++ {
			if djVSllzO5[aepbxLiWOZ].pMhnIOjkP == 0 {
				if aepbxLiWOZ == 0 {
					djVSllzO5[0].rrnbR4 =  /*line :1*/ g3X9fa(nil).ea7MDYcwtuzk(iIGcv9,  /*line :1*/ VYe6D0(paETBQw))
					djVSllzO5[0].pMhnIOjkP = vEkdraJ * paETBQw
				} else {
					djVSllzO5[aepbxLiWOZ].rrnbR4 =  /*line :1*/ g3X9fa(nil).pJ61taToc(djVSllzO5[aepbxLiWOZ-1].rrnbR4)
					djVSllzO5[aepbxLiWOZ].pMhnIOjkP = 2 * djVSllzO5[aepbxLiWOZ-1].pMhnIOjkP
				}

				sGlZ9QmSSqFV =  /*line :1*/ g3X9fa(nil).bj0nVC(djVSllzO5[aepbxLiWOZ].rrnbR4)
				for  /*line :1*/ pMUCN3(sGlZ9QmSSqFV, sGlZ9QmSSqFV, kWa1YrIHZo, 0) == 0 {
					djVSllzO5[aepbxLiWOZ].rrnbR4 =  /*line :1*/ djVSllzO5[aepbxLiWOZ].rrnbR4.bj0nVC(sGlZ9QmSSqFV)
					djVSllzO5[aepbxLiWOZ].pMhnIOjkP++
				}

				djVSllzO5[aepbxLiWOZ].cev83XwLwKz =  /*line :1*/ djVSllzO5[aepbxLiWOZ].rrnbR4.wZwJ3Y0()
			}
		}
	}

	if kWa1YrIHZo == 10 {
		 /*line :1*/ tFHi4yBEr.Unlock()
	}

	return djVSllzO5
}
