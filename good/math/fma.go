//line :1
package math

import bits "math/bits"

func q7GRsm4g799(_BqkSz0 uint64) uint64 {
	if _BqkSz0 == 0 {
		return 1
	}
	return 0

}

func xkZ1QQX8(_BqkSz0 uint64) uint64 {
	if _BqkSz0 != 0 {
		return 1
	}
	return 0

}

func h8xUEM(haVjyY, xDIQXq_ uint64, d2X2v9 uint) (yYCutiG6s3, uC3FZRnPrx6 uint64) {
	yYCutiG6s3 = haVjyY<<d2X2v9 | xDIQXq_>>(64-d2X2v9) | xDIQXq_<<(d2X2v9-64)
	uC3FZRnPrx6 = xDIQXq_ << d2X2v9
	return
}

func lIEnc7kC(haVjyY, xDIQXq_ uint64, d2X2v9 uint) (yYCutiG6s3, uC3FZRnPrx6 uint64) {
	uC3FZRnPrx6 = xDIQXq_>>d2X2v9 | haVjyY<<(64-d2X2v9) | haVjyY>>(d2X2v9-64)
	yYCutiG6s3 = haVjyY >> d2X2v9
	return
}





func sHuNjkF(haVjyY, xDIQXq_ uint64, d2X2v9 uint) (yYCutiG6s3, uC3FZRnPrx6 uint64) {

	switch {
	case d2X2v9 == 0:
		return haVjyY, xDIQXq_
	case d2X2v9 == 64:
		return 0, haVjyY |  /*line :1*/ xkZ1QQX8(xDIQXq_)
	case d2X2v9 >= 128:
		return 0,  /*line :1*/ xkZ1QQX8(haVjyY | xDIQXq_)
	case d2X2v9 < 64:
		yYCutiG6s3, uC3FZRnPrx6 =  /*line :1*/ lIEnc7kC(haVjyY, xDIQXq_, d2X2v9)
		uC3FZRnPrx6 |=  /*line :1*/ xkZ1QQX8(xDIQXq_ & (1<<d2X2v9 - 1))
	case d2X2v9 < 128:
		yYCutiG6s3, uC3FZRnPrx6 =  /*line :1*/ lIEnc7kC(haVjyY, xDIQXq_, d2X2v9)
		uC3FZRnPrx6 |=  /*line :1*/ xkZ1QQX8(haVjyY&(1<<(d2X2v9-64)-1) | xDIQXq_)
	}
	return
}

func rXbKFJ1(haVjyY, xDIQXq_ uint64) (zCJ2Zc0 int32) {
	zCJ2Zc0 =  /*line :1*/ int32( /*line :1*/ bits.HYxjvQ(haVjyY))
	if zCJ2Zc0 == 64 {
		zCJ2Zc0 +=  /*line :1*/ int32( /*line :1*/ bits.HYxjvQ(xDIQXq_))
	}
	return zCJ2Zc0
}




func gfn5QDx(alfthQvYkJ uint64) (awRlyItEG2gn uint32, nF4u1Vnrit int32, brm3Tg0 uint64) {
	awRlyItEG2gn =  /*line :1*/ uint32(alfthQvYkJ >> 63)
	nF4u1Vnrit =  /*line :1*/ int32(alfthQvYkJ>>52) & mask
	brm3Tg0 = alfthQvYkJ & fracMask

	if nF4u1Vnrit == 0 {

		u26rEEu_t :=  /*line :1*/ uint( /*line :1*/ bits.HYxjvQ(brm3Tg0) - 11)
		brm3Tg0 <<= u26rEEu_t
		nF4u1Vnrit = 1 -  /*line :1*/ int32(u26rEEu_t)
	} else {

		brm3Tg0 |= 1 << 52
	}
	return
}



func FMA(_BqkSz0, nNeKvlvK6, gZDpjWvXax_q float64) float64 {
	fWQgJ4Me, t_pM5qrMYuH, xu8QFaAdN :=  /*line :1*/ GKIRmoHE(_BqkSz0),  /*line :1*/ GKIRmoHE(nNeKvlvK6),  /*line :1*/ GKIRmoHE(gZDpjWvXax_q)

	if _BqkSz0 == 0.0 || nNeKvlvK6 == 0.0 || gZDpjWvXax_q == 0.0 || fWQgJ4Me&uvinf == uvinf || t_pM5qrMYuH&uvinf == uvinf {
		return _BqkSz0*nNeKvlvK6 + gZDpjWvXax_q
	}

	if xu8QFaAdN&uvinf == uvinf {
		return gZDpjWvXax_q
	}

	xnW4aP, lBwiw8EWW, tc0cBt :=  /*line :1*/ gfn5QDx(fWQgJ4Me)
	fmbRWz9kye, ygs1Wa, gj62m0l :=  /*line :1*/ gfn5QDx(t_pM5qrMYuH)
	eeVIae4JHv, f3HFUYGi, j54cyhMb95nF :=  /*line :1*/ gfn5QDx(xu8QFaAdN)

	kJFq392 := lBwiw8EWW + ygs1Wa - bias + 1

	feRZ6ClS, g2aUPw :=  /*line :1*/ bits.Mul64(tc0cBt<<10, gj62m0l<<11)
	gt6wG_Gy6e4E, aTWVJQ := j54cyhMb95nF<<10,  /*line :1*/ uint64(0)
	wW3UMzuw0dv := xnW4aP ^ fmbRWz9kye

	dXgjxv6AN :=  /*line :1*/ uint((^feRZ6ClS >> 62) & 1)
	feRZ6ClS, g2aUPw =  /*line :1*/ h8xUEM(feRZ6ClS, g2aUPw, dXgjxv6AN)
	kJFq392 -=  /*line :1*/ int32(dXgjxv6AN)

	if kJFq392 < f3HFUYGi || kJFq392 == f3HFUYGi && feRZ6ClS < gt6wG_Gy6e4E {
		wW3UMzuw0dv, kJFq392, feRZ6ClS, g2aUPw, eeVIae4JHv, f3HFUYGi, gt6wG_Gy6e4E, aTWVJQ = eeVIae4JHv, f3HFUYGi, gt6wG_Gy6e4E, aTWVJQ, wW3UMzuw0dv, kJFq392, feRZ6ClS, g2aUPw
	}

	if wW3UMzuw0dv != eeVIae4JHv && kJFq392 == f3HFUYGi && feRZ6ClS == gt6wG_Gy6e4E && g2aUPw == aTWVJQ {
		return 0
	}

	gt6wG_Gy6e4E, aTWVJQ =  /*line :1*/ sHuNjkF(gt6wG_Gy6e4E, aTWVJQ,  /*line :1*/ uint(kJFq392-f3HFUYGi))

	
	var b8CRqLSt, cU6qMJer46oH uint64
	if wW3UMzuw0dv == eeVIae4JHv {

		g2aUPw, cU6qMJer46oH =  /*line :1*/ bits.Add64(g2aUPw, aTWVJQ, 0)
		feRZ6ClS, _ =  /*line :1*/ bits.Add64(feRZ6ClS, gt6wG_Gy6e4E, cU6qMJer46oH)
		kJFq392 -=  /*line :1*/ int32(^feRZ6ClS >> 63)
		feRZ6ClS, b8CRqLSt =  /*line :1*/ sHuNjkF(feRZ6ClS, g2aUPw,  /*line :1*/ uint(64+feRZ6ClS>>63))
	} else {

		g2aUPw, cU6qMJer46oH =  /*line :1*/ bits.Sub64(g2aUPw, aTWVJQ, 0)
		feRZ6ClS, _ =  /*line :1*/ bits.Sub64(feRZ6ClS, gt6wG_Gy6e4E, cU6qMJer46oH)
		_Io_PHw22dtv :=  /*line :1*/ rXbKFJ1(feRZ6ClS, g2aUPw)
		kJFq392 -= _Io_PHw22dtv
		b8CRqLSt, g2aUPw =  /*line :1*/ h8xUEM(feRZ6ClS, g2aUPw,  /*line :1*/ uint(_Io_PHw22dtv-1))
		b8CRqLSt |=  /*line :1*/ xkZ1QQX8(g2aUPw)
	}

	if kJFq392 > 1022+bias || kJFq392 == 1022+bias && (b8CRqLSt+1<<9)>>63 == 1 {

		return  /*line :1*/ QUB2Kf63p( /*line :1*/ uint64(wW3UMzuw0dv)<<63 | uvinf)
	}
	if kJFq392 < 0 {
		d2X2v9 :=  /*line :1*/ uint(-kJFq392)
		b8CRqLSt = b8CRqLSt>>d2X2v9 |  /*line :1*/ xkZ1QQX8(b8CRqLSt&(1<<d2X2v9-1))
		kJFq392 = 0
	}
	b8CRqLSt = ((b8CRqLSt + 1<<9) >> 10) & ^ /*line :1*/ q7GRsm4g799((b8CRqLSt&(1<<10-1))^1<<9)
	kJFq392 &= - /*line :1*/ int32( /*line :1*/ xkZ1QQX8(b8CRqLSt))
	return  /*line :1*/ QUB2Kf63p( /*line :1*/ uint64(wW3UMzuw0dv)<<63 +  /*line :1*/ uint64(kJFq392)<<52 + b8CRqLSt)
}
