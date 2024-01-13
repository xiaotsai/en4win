//line :1
package math

import (
	bits "math/bits"
)












const reduceThreshold = 1 << 29








func noJmE7EjH3k(_BqkSz0 float64) (vnPtu1up3 uint64, gZDpjWvXax_q float64) {
	const PI4 = Pi / 4
	if _BqkSz0 < PI4 {
		return 0, _BqkSz0
	}

	cCmjKOxt :=  /*line :1*/ GKIRmoHE(_BqkSz0)
	nF4u1Vnrit :=  /*line :1*/ int(cCmjKOxt>>shift&mask) - bias - shift
	cCmjKOxt &^= mask << shift
	cCmjKOxt |= 1 << shift

	p2pgND, wgqLr1 :=  /*line :1*/ uint(nF4u1Vnrit+61)/64,  /*line :1*/ uint(nF4u1Vnrit+61)%64
	v2MiDZjtyYiK := (kr3IN8G0iKax[p2pgND] << wgqLr1) | (kr3IN8G0iKax[p2pgND+1] >> (64 - wgqLr1))
	b2xtU6qKJa := (kr3IN8G0iKax[p2pgND+1] << wgqLr1) | (kr3IN8G0iKax[p2pgND+2] >> (64 - wgqLr1))
	otS2p8 := (kr3IN8G0iKax[p2pgND+2] << wgqLr1) | (kr3IN8G0iKax[p2pgND+3] >> (64 - wgqLr1))

	fH7Qdtmg4W, _ :=  /*line :1*/ bits.Mul64(otS2p8, cCmjKOxt)
	eSaqaP, dFWg0kW9zU :=  /*line :1*/ bits.Mul64(b2xtU6qKJa, cCmjKOxt)
	j8bydil := v2MiDZjtyYiK * cCmjKOxt
	j0RLYji, cU6qMJer46oH :=  /*line :1*/ bits.Add64(dFWg0kW9zU, fH7Qdtmg4W, 0)
	j2BkoXTMn9D, _ :=  /*line :1*/ bits.Add64(j8bydil, eSaqaP, cU6qMJer46oH)

	vnPtu1up3 = j2BkoXTMn9D >> 61

	j2BkoXTMn9D = j2BkoXTMn9D<<3 | j0RLYji>>61
	rXbKFJ1 :=  /*line :1*/ uint( /*line :1*/ bits.HYxjvQ(j2BkoXTMn9D))
	gwBaYgYgA :=  /*line :1*/ uint64(bias - (rXbKFJ1 + 1))

	j2BkoXTMn9D = (j2BkoXTMn9D << (rXbKFJ1 + 1)) | (j0RLYji >> (64 - (rXbKFJ1 + 1)))
	j2BkoXTMn9D >>= 64 - shift

	j2BkoXTMn9D |= gwBaYgYgA << shift
	gZDpjWvXax_q =  /*line :1*/ QUB2Kf63p(j2BkoXTMn9D)

	if vnPtu1up3&1 == 1 {
		vnPtu1up3++
		vnPtu1up3 &= 7
		gZDpjWvXax_q--
	}

	return vnPtu1up3, gZDpjWvXax_q * PI4
}





var kr3IN8G0iKax = [...]uint64{
	0x0000000000000001,
	0x45f306dc9c882a53,
	0xf84eafa3ea69bb81,
	0xb6c52b3278872083,
	0xfca2c757bd778ac3,
	0x6e48dc74849ba5c0,
	0x0c925dd413a32439,
	0xfc3bd63962534e7d,
	0xd1046bea5d768909,
	0xd338e04d68befc82,
	0x7323ac7306a673e9,
	0x3908bf177bf25076,
	0x3ff12fffbc0b301f,
	0xde5e2316b414da3e,
	0xda6cfd9e4f96136e,
	0x9e8c7ecd3cbfd45a,
	0xea4f758fd7cbe2f6,
	0x7a0e73ef14a525d4,
	0xd7f6bf623f1aba10,
	0xac06608df8f6d757,
}
