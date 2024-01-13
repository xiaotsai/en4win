//line :1
package big

import rand "os208GicmRnc"


















func (pmEgen2 *ZMtGuo) ProbablyPrime(h_Wr_yqq int) bool {

	if h_Wr_yqq < 0 {
		 /*line :1*/ panic("negative n for ProbablyPrime")
	}
	if pmEgen2.hCTOmp0gckSa ||  /*line :1*/ len(pmEgen2.qho77PBKF) == 0 {
		return false
	}

	
	const primeBitMask uint64 = 1<<2 | 1<<3 | 1<<5 | 1<<7 |
		1<<11 | 1<<13 | 1<<17 | 1<<19 | 1<<23 | 1<<29 | 1<<31 |
		1<<37 | 1<<41 | 1<<43 | 1<<47 | 1<<53 | 1<<59 | 1<<61

	uKQ8Eox := pmEgen2.qho77PBKF[0]
	if  /*line :1*/ len(pmEgen2.qho77PBKF) == 1 && uKQ8Eox < 64 {
		return primeBitMask&(1<<uKQ8Eox) != 0
	}

	if uKQ8Eox&1 == 0 {
		return false
	}

	const primesA = 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23 * 37
	const primesB = 29 * 31 * 41 * 43 * 47 * 53

	var ibehL2sR, he_pECLA4t uint32
	switch _W {
	case 32:
		ibehL2sR =  /*line :1*/ uint32( /*line :1*/ pmEgen2.qho77PBKF.tPll2cuYqvU(primesA))
		he_pECLA4t =  /*line :1*/ uint32( /*line :1*/ pmEgen2.qho77PBKF.tPll2cuYqvU(primesB))
	case 64:
		vFx5p_cm :=  /*line :1*/ pmEgen2.qho77PBKF.tPll2cuYqvU((primesA * primesB) & _M)
		ibehL2sR =  /*line :1*/ uint32(vFx5p_cm % primesA)
		he_pECLA4t =  /*line :1*/ uint32(vFx5p_cm % primesB)
	default:
		 /*line :1*/ panic("math/big: invalid word size")
	}

	if ibehL2sR%3 == 0 || ibehL2sR%5 == 0 || ibehL2sR%7 == 0 || ibehL2sR%11 == 0 || ibehL2sR%13 == 0 || ibehL2sR%17 == 0 || ibehL2sR%19 == 0 || ibehL2sR%23 == 0 || ibehL2sR%37 == 0 ||
		he_pECLA4t%29 == 0 || he_pECLA4t%31 == 0 || he_pECLA4t%41 == 0 || he_pECLA4t%43 == 0 || he_pECLA4t%47 == 0 || he_pECLA4t%53 == 0 {
		return false
	}

	return  /*line :1*/ pmEgen2.qho77PBKF.kx5J7tV2(h_Wr_yqq+1, true) &&  /*line :1*/ pmEgen2.qho77PBKF.b7rY2z1c()
}






func (h_Wr_yqq g3X9fa) kx5J7tV2(u5uVIf9nH4 int, fF2zwdBL57u bool) bool {
	kjEeryP :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(h_Wr_yqq, lG0t2KfqGwyP)

	_DfINwXQ :=  /*line :1*/ kjEeryP.kDun6ak()
	yL0_p0wUnc :=  /*line :1*/ g3X9fa(nil).qzOaHmR(kjEeryP, _DfINwXQ)

	vYYq5UQw :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(kjEeryP, euJqHOlfT)
	biERyAySo :=  /*line :1*/ rand.PamkngYsA( /*line :1*/ rand.DF2qvnmwCjq6( /*line :1*/ int64(h_Wr_yqq[0])))

	var pmEgen2, oxFS5wa5, n5RmhEFEv g3X9fa
	yXPtCxsA4R :=  /*line :1*/ vYYq5UQw.wZwJ3Y0()

NextRandom:
	for aepbxLiWOZ := 0; aepbxLiWOZ < u5uVIf9nH4; aepbxLiWOZ++ {
		if aepbxLiWOZ == u5uVIf9nH4-1 && fF2zwdBL57u {
			pmEgen2 =  /*line :1*/ pmEgen2.bj0nVC(euJqHOlfT)
		} else {
			pmEgen2 =  /*line :1*/ pmEgen2.calV9Cfo(biERyAySo, vYYq5UQw, yXPtCxsA4R)
			pmEgen2 =  /*line :1*/ pmEgen2.jXv7nnhY_(pmEgen2, euJqHOlfT)
		}
		oxFS5wa5 =  /*line :1*/ oxFS5wa5.iGIIsMONlhq(pmEgen2, yL0_p0wUnc, h_Wr_yqq, false)
		if  /*line :1*/ oxFS5wa5.beaurQHiT(lG0t2KfqGwyP) == 0 ||  /*line :1*/ oxFS5wa5.beaurQHiT(kjEeryP) == 0 {
			continue
		}
		for g77TAi :=  /*line :1*/ uint(1); g77TAi < _DfINwXQ; g77TAi++ {
			oxFS5wa5 =  /*line :1*/ oxFS5wa5.pJ61taToc(oxFS5wa5)
			n5RmhEFEv, oxFS5wa5 =  /*line :1*/ n5RmhEFEv.xOZxoLyl(oxFS5wa5, oxFS5wa5, h_Wr_yqq)
			if  /*line :1*/ oxFS5wa5.beaurQHiT(kjEeryP) == 0 {
				continue NextRandom
			}
			if  /*line :1*/ oxFS5wa5.beaurQHiT(lG0t2KfqGwyP) == 0 {
				return false
			}
		}
		return false
	}

	return true
}

























func (h_Wr_yqq g3X9fa) b7rY2z1c() bool {

	if  /*line :1*/ len(h_Wr_yqq) == 0 ||  /*line :1*/ h_Wr_yqq.beaurQHiT(lG0t2KfqGwyP) == 0 {
		return false
	}

	if h_Wr_yqq[0]&1 == 0 {
		return  /*line :1*/ h_Wr_yqq.beaurQHiT(euJqHOlfT) == 0
	}

	djOsNJsLfH5U :=  /*line :1*/ VYe6D0(3)
	s31g8J := g3X9fa{1}
	dms1sqFQT :=  /*line :1*/ g3X9fa(nil)
	seiXH8aEU0 := &ZMtGuo{qho77PBKF: s31g8J}
	v5lS9Y4_ := &ZMtGuo{qho77PBKF: h_Wr_yqq}
	for ; ; djOsNJsLfH5U++ {
		if djOsNJsLfH5U > 10000 {

			 /*line :1*/ panic("math/big: internal error: cannot find (D/n) = -1 for " +  /*line :1*/ v5lS9Y4_.String())
		}
		s31g8J[0] = djOsNJsLfH5U*djOsNJsLfH5U - 4
		g77TAi :=  /*line :1*/ SUxhyDUe(seiXH8aEU0, v5lS9Y4_)
		if g77TAi == -1 {
			break
		}
		if g77TAi == 0 {

			return  /*line :1*/ len(h_Wr_yqq) == 1 && h_Wr_yqq[0] == djOsNJsLfH5U+2
		}
		if djOsNJsLfH5U == 40 {

			dms1sqFQT =  /*line :1*/ dms1sqFQT.vf9TNsbnLN(h_Wr_yqq)
			dms1sqFQT =  /*line :1*/ dms1sqFQT.pJ61taToc(dms1sqFQT)
			if  /*line :1*/ dms1sqFQT.beaurQHiT(h_Wr_yqq) == 0 {
				return false
			}
		}
	}

	nR7KU4mGsdy :=  /*line :1*/ g3X9fa(nil).jXv7nnhY_(h_Wr_yqq, lG0t2KfqGwyP)
	vFx5p_cm :=  /*line :1*/ int( /*line :1*/ nR7KU4mGsdy.kDun6ak())
	nR7KU4mGsdy =  /*line :1*/ nR7KU4mGsdy.qzOaHmR(nR7KU4mGsdy,  /*line :1*/ uint(vFx5p_cm))
	kHfItpQmtc :=  /*line :1*/ g3X9fa(nil).lt4VKILNCo(h_Wr_yqq, euJqHOlfT)

	aCtTeb :=  /*line :1*/ g3X9fa(nil).pSzoQ7PMW(djOsNJsLfH5U)
	tMv2jJ3DQIL1 :=  /*line :1*/ g3X9fa(nil).pSzoQ7PMW(2)
	kDFhFTSMEa :=  /*line :1*/ g3X9fa(nil).pSzoQ7PMW(djOsNJsLfH5U)
	gr_YTl04l5nz :=  /*line :1*/ g3X9fa(nil)
	for aepbxLiWOZ :=  /*line :1*/ int( /*line :1*/ nR7KU4mGsdy.wZwJ3Y0()); aepbxLiWOZ >= 0; aepbxLiWOZ-- {
		if  /*line :1*/ nR7KU4mGsdy.bjRK8rI0NMn( /*line :1*/ uint(aepbxLiWOZ)) != 0 {

			dms1sqFQT =  /*line :1*/ dms1sqFQT.d22x6Ygoc80O(tMv2jJ3DQIL1, kDFhFTSMEa)
			dms1sqFQT =  /*line :1*/ dms1sqFQT.jXv7nnhY_(dms1sqFQT, h_Wr_yqq)
			dms1sqFQT =  /*line :1*/ dms1sqFQT.lt4VKILNCo(dms1sqFQT, aCtTeb)
			gr_YTl04l5nz, tMv2jJ3DQIL1 =  /*line :1*/ gr_YTl04l5nz.xOZxoLyl(tMv2jJ3DQIL1, dms1sqFQT, h_Wr_yqq)

			dms1sqFQT =  /*line :1*/ dms1sqFQT.pJ61taToc(kDFhFTSMEa)
			dms1sqFQT =  /*line :1*/ dms1sqFQT.jXv7nnhY_(dms1sqFQT, kHfItpQmtc)
			gr_YTl04l5nz, kDFhFTSMEa =  /*line :1*/ gr_YTl04l5nz.xOZxoLyl(kDFhFTSMEa, dms1sqFQT, h_Wr_yqq)
		} else {

			dms1sqFQT =  /*line :1*/ dms1sqFQT.d22x6Ygoc80O(tMv2jJ3DQIL1, kDFhFTSMEa)
			dms1sqFQT =  /*line :1*/ dms1sqFQT.jXv7nnhY_(dms1sqFQT, h_Wr_yqq)
			dms1sqFQT =  /*line :1*/ dms1sqFQT.lt4VKILNCo(dms1sqFQT, aCtTeb)
			gr_YTl04l5nz, kDFhFTSMEa =  /*line :1*/ gr_YTl04l5nz.xOZxoLyl(kDFhFTSMEa, dms1sqFQT, h_Wr_yqq)

			dms1sqFQT =  /*line :1*/ dms1sqFQT.pJ61taToc(tMv2jJ3DQIL1)
			dms1sqFQT =  /*line :1*/ dms1sqFQT.jXv7nnhY_(dms1sqFQT, kHfItpQmtc)
			gr_YTl04l5nz, tMv2jJ3DQIL1 =  /*line :1*/ gr_YTl04l5nz.xOZxoLyl(tMv2jJ3DQIL1, dms1sqFQT, h_Wr_yqq)
		}
	}

	if  /*line :1*/ tMv2jJ3DQIL1.beaurQHiT(euJqHOlfT) == 0 ||  /*line :1*/ tMv2jJ3DQIL1.beaurQHiT(kHfItpQmtc) == 0 {

		dms1sqFQT :=  /*line :1*/ dms1sqFQT.d22x6Ygoc80O(tMv2jJ3DQIL1, aCtTeb)
		gr_YTl04l5nz :=  /*line :1*/ gr_YTl04l5nz.z839tk6CmDHT(kDFhFTSMEa, 1)
		if  /*line :1*/ dms1sqFQT.beaurQHiT(gr_YTl04l5nz) < 0 {
			dms1sqFQT, gr_YTl04l5nz = gr_YTl04l5nz, dms1sqFQT
		}
		dms1sqFQT =  /*line :1*/ dms1sqFQT.lt4VKILNCo(dms1sqFQT, gr_YTl04l5nz)
		gz_H7UklVHu := kDFhFTSMEa
		kDFhFTSMEa = nil
		_ = kDFhFTSMEa
		gr_YTl04l5nz, gz_H7UklVHu =  /*line :1*/ gr_YTl04l5nz.xOZxoLyl(gz_H7UklVHu, dms1sqFQT, h_Wr_yqq)
		if  /*line :1*/ len(gz_H7UklVHu) == 0 {
			return true
		}
	}

	for amoa4Px := 0; amoa4Px < vFx5p_cm-1; amoa4Px++ {
		if  /*line :1*/ len(tMv2jJ3DQIL1) == 0 {
			return true
		}

		if  /*line :1*/ len(tMv2jJ3DQIL1) == 1 && tMv2jJ3DQIL1[0] == 2 {
			return false
		}

		dms1sqFQT =  /*line :1*/ dms1sqFQT.pJ61taToc(tMv2jJ3DQIL1)
		dms1sqFQT =  /*line :1*/ dms1sqFQT.lt4VKILNCo(dms1sqFQT, euJqHOlfT)
		gr_YTl04l5nz, tMv2jJ3DQIL1 =  /*line :1*/ gr_YTl04l5nz.xOZxoLyl(tMv2jJ3DQIL1, dms1sqFQT, h_Wr_yqq)
	}
	return false
}
