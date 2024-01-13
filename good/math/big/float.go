//line :1
package big

import (
	fmt "kFsoCfy5zWG"
	math "math"
	bits "math/bits"
)

const debugFloat = false

type WH8dWN struct {
	iaaUnOsr7_M	uint32
	hNG8iJiHdkxT	WsODiYye5l
	ddq1eQm	IG11Ul7qo
	fNeR0A	cZg9qpR
	g2AL4u0IBtd	bool
	s7DDkIlY	g3X9fa
	khThBIzMZ	int32
}

type At_rvS_hmcW struct {
	tPVEytJidW string
}

func (gDonrwI8VP50 At_rvS_hmcW) Error() string {
	return gDonrwI8VP50.tPVEytJidW
}

func EfoO1hKDhm0(pmEgen2 float64) *WH8dWN {
	if  /*line :1*/ math.OIdLpDqq(pmEgen2) {
		 /*line :1*/ panic(At_rvS_hmcW{"NewFloat(NaN)"})
	}
	return  /*line :1*/ new(WH8dWN).SetFloat64(pmEgen2)
}

const (
	MaxExp	= math.MaxInt32
	MinExp	= math.MinInt32
	MaxPrec	= math.MaxUint32
)

type cZg9qpR byte

const (
	zero	cZg9qpR	= iota
	finite
	inf
)

type WsODiYye5l byte

const (
	ToNearestEven	WsODiYye5l	= iota
	ToNearestAway
	ToZero
	AwayFromZero
	ToNegativeInf
	ToPositiveInf
)

//go:generate stringer -type=RoundingMode

type IG11Ul7qo int8

const (
	Below	IG11Ul7qo	= -1
	Exact	IG11Ul7qo	= 0
	Above	IG11Ul7qo	= +1
)

//go:generate stringer -type=Accuracy

func (uW0eIVjy *WH8dWN) SetPrec(cZUDpnJQn uint) *WH8dWN {
	uW0eIVjy.ddq1eQm = Exact

	if cZUDpnJQn == 0 {
		uW0eIVjy.iaaUnOsr7_M = 0
		if uW0eIVjy.fNeR0A == finite {

			uW0eIVjy.ddq1eQm =  /*line :1*/ dBh3XtAgq(uW0eIVjy.g2AL4u0IBtd)
			uW0eIVjy.fNeR0A = zero
		}
		return uW0eIVjy
	}

	if cZUDpnJQn > MaxPrec {
		cZUDpnJQn = MaxPrec
	}
	eBdyryMZh := uW0eIVjy.iaaUnOsr7_M
	uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ uint32(cZUDpnJQn)
	if uW0eIVjy.iaaUnOsr7_M < eBdyryMZh {
		 /*line :1*/ uW0eIVjy.jrT168o(0)
	}
	return uW0eIVjy
}

func dBh3XtAgq(vUuG8bWwVQC bool) IG11Ul7qo {
	if vUuG8bWwVQC {
		return Above
	}
	return Below
}

func (uW0eIVjy *WH8dWN) SetMode(aca9LekD WsODiYye5l) *WH8dWN {
	uW0eIVjy.hNG8iJiHdkxT = aca9LekD
	uW0eIVjy.ddq1eQm = Exact
	return uW0eIVjy
}

func (pmEgen2 *WH8dWN) Prec() uint {
	return  /*line :1*/ uint(pmEgen2.iaaUnOsr7_M)
}

func (pmEgen2 *WH8dWN) MinPrec() uint {
	if pmEgen2.fNeR0A != finite {
		return 0
	}
	return  /*line :1*/ uint( /*line :1*/ len(pmEgen2.s7DDkIlY))*_W -  /*line :1*/ pmEgen2.s7DDkIlY.kDun6ak()
}

func (pmEgen2 *WH8dWN) Mode() WsODiYye5l {
	return pmEgen2.hNG8iJiHdkxT
}

func (pmEgen2 *WH8dWN) Acc() IG11Ul7qo {
	return pmEgen2.ddq1eQm
}

func (pmEgen2 *WH8dWN) Sign() int {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}
	if pmEgen2.fNeR0A == zero {
		return 0
	}
	if pmEgen2.g2AL4u0IBtd {
		return -1
	}
	return 1
}

func (pmEgen2 *WH8dWN) MantExp(kfG4qg_2AdyH *WH8dWN) (xXsI4WC int) {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}
	if pmEgen2.fNeR0A == finite {
		xXsI4WC =  /*line :1*/ int(pmEgen2.khThBIzMZ)
	}
	if kfG4qg_2AdyH != nil {
		 /*line :1*/ kfG4qg_2AdyH.Copy(pmEgen2)
		if kfG4qg_2AdyH.fNeR0A == finite {
			kfG4qg_2AdyH.khThBIzMZ = 0
		}
	}
	return
}

func (uW0eIVjy *WH8dWN) pKBlNY(xXsI4WC int64, c8rDOxCvz_ZX uint) {
	if xXsI4WC < MinExp {

		uW0eIVjy.ddq1eQm =  /*line :1*/ dBh3XtAgq(uW0eIVjy.g2AL4u0IBtd)
		uW0eIVjy.fNeR0A = zero
		return
	}

	if xXsI4WC > MaxExp {

		uW0eIVjy.ddq1eQm =  /*line :1*/ dBh3XtAgq(!uW0eIVjy.g2AL4u0IBtd)
		uW0eIVjy.fNeR0A = inf
		return
	}

	uW0eIVjy.fNeR0A = finite
	uW0eIVjy.khThBIzMZ =  /*line :1*/ int32(xXsI4WC)
	 /*line :1*/ uW0eIVjy.jrT168o(c8rDOxCvz_ZX)
}

func (uW0eIVjy *WH8dWN) SetMantExp(kfG4qg_2AdyH *WH8dWN, xXsI4WC int) *WH8dWN {
	if debugFloat {
		 /*line :1*/ uW0eIVjy.lxB3waL2v()
		 /*line :1*/ kfG4qg_2AdyH.lxB3waL2v()
	}
	 /*line :1*/ uW0eIVjy.Copy(kfG4qg_2AdyH)

	if uW0eIVjy.fNeR0A == finite {

		 /*line :1*/ uW0eIVjy.pKBlNY( /*line :1*/ int64(uW0eIVjy.khThBIzMZ)+ /*line :1*/ int64(xXsI4WC), 0)
	}
	return uW0eIVjy
}

func (pmEgen2 *WH8dWN) Signbit() bool {
	return pmEgen2.g2AL4u0IBtd
}

func (pmEgen2 *WH8dWN) IsInf() bool {
	return pmEgen2.fNeR0A == inf
}

func (pmEgen2 *WH8dWN) IsInt() bool {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	if pmEgen2.fNeR0A != finite {
		return pmEgen2.fNeR0A == zero
	}

	if pmEgen2.khThBIzMZ <= 0 {
		return false
	}

	return pmEgen2.iaaUnOsr7_M <=  /*line :1*/ uint32(pmEgen2.khThBIzMZ) ||  /*line :1*/ pmEgen2.MinPrec() <=  /*line :1*/ uint(pmEgen2.khThBIzMZ)
}

func (pmEgen2 *WH8dWN) lxB3waL2v() {
	if !debugFloat {

		 /*line :1*/ panic("validate called but debugFloat is not set")
	}
	if isKdjaXeB :=  /*line :1*/ pmEgen2.asvxnebpcnsA(); isKdjaXeB != "" {
		 /*line :1*/ panic(isKdjaXeB)
	}
}

func (pmEgen2 *WH8dWN) asvxnebpcnsA() string {
	if pmEgen2.fNeR0A != finite {
		return ""
	}
	gS2TXhHpYaX4 :=  /*line :1*/ len(pmEgen2.s7DDkIlY)
	if gS2TXhHpYaX4 == 0 {
		return "nonzero finite number with empty mantissa"
	}
	const msb = 1 << (_W - 1)
	if pmEgen2.s7DDkIlY[gS2TXhHpYaX4-1]&msb == 0 {
		return  /*line :1*/ fmt.IBsS3Oc("msb not set in last word %#x of %s", pmEgen2.s7DDkIlY[gS2TXhHpYaX4-1],  /*line :1*/ pmEgen2.Text('p', 0))
	}
	if pmEgen2.iaaUnOsr7_M == 0 {
		return "zero precision finite number"
	}
	return ""
}

func (uW0eIVjy *WH8dWN) jrT168o(c8rDOxCvz_ZX uint) {
	if debugFloat {
		 /*line :1*/ uW0eIVjy.lxB3waL2v()
	}

	uW0eIVjy.ddq1eQm = Exact
	if uW0eIVjy.fNeR0A != finite {

		return
	}

	gS2TXhHpYaX4 :=  /*line :1*/ uint32( /*line :1*/ len(uW0eIVjy.s7DDkIlY))
	uopo38aVK_H := gS2TXhHpYaX4 * _W
	if uopo38aVK_H <= uW0eIVjy.iaaUnOsr7_M {

		return
	}

	vFx5p_cm :=  /*line :1*/ uint(uopo38aVK_H - uW0eIVjy.iaaUnOsr7_M - 1)
	aBuqbUZZ :=  /*line :1*/ uW0eIVjy.s7DDkIlY.bjRK8rI0NMn(vFx5p_cm) & 1

	if c8rDOxCvz_ZX == 0 && (aBuqbUZZ == 0 || uW0eIVjy.hNG8iJiHdkxT == ToNearestEven) {
		c8rDOxCvz_ZX =  /*line :1*/ uW0eIVjy.s7DDkIlY.cbgX3ksmdgax(vFx5p_cm)
	}
	c8rDOxCvz_ZX &= 1

	h_Wr_yqq := (uW0eIVjy.iaaUnOsr7_M + (_W - 1)) / _W
	if gS2TXhHpYaX4 > h_Wr_yqq {
		 /*line :1*/ copy(uW0eIVjy.s7DDkIlY, uW0eIVjy.s7DDkIlY[gS2TXhHpYaX4-h_Wr_yqq:])
		uW0eIVjy.s7DDkIlY = uW0eIVjy.s7DDkIlY[:h_Wr_yqq]
	}

	b_CDmuq7 := h_Wr_yqq*_W - uW0eIVjy.iaaUnOsr7_M
	d4Lnax :=  /*line :1*/ VYe6D0(1) << b_CDmuq7

	if aBuqbUZZ|c8rDOxCvz_ZX != 0 {

		mo5tjh9Xk := false
		switch uW0eIVjy.hNG8iJiHdkxT {
		case ToNegativeInf:
			mo5tjh9Xk = uW0eIVjy.g2AL4u0IBtd
		case ToZero:

		case ToNearestEven:
			mo5tjh9Xk = aBuqbUZZ != 0 && (c8rDOxCvz_ZX != 0 || uW0eIVjy.s7DDkIlY[0]&d4Lnax != 0)
		case ToNearestAway:
			mo5tjh9Xk = aBuqbUZZ != 0
		case AwayFromZero:
			mo5tjh9Xk = true
		case ToPositiveInf:
			mo5tjh9Xk = !uW0eIVjy.g2AL4u0IBtd
		default:
			 /*line :1*/ panic("unreachable")
		}

		uW0eIVjy.ddq1eQm =  /*line :1*/ dBh3XtAgq(mo5tjh9Xk != uW0eIVjy.g2AL4u0IBtd)

		if mo5tjh9Xk {

			if  /*line :1*/ aJRvTH(uW0eIVjy.s7DDkIlY, uW0eIVjy.s7DDkIlY, d4Lnax) != 0 {

				if uW0eIVjy.khThBIzMZ >= MaxExp {

					uW0eIVjy.fNeR0A = inf
					return
				}
				uW0eIVjy.khThBIzMZ++

				 /*line :1*/ t_ku9iRroNp(uW0eIVjy.s7DDkIlY, uW0eIVjy.s7DDkIlY, 1)

				const msb = 1 << (_W - 1)
				uW0eIVjy.s7DDkIlY[h_Wr_yqq-1] |= msb
			}
		}
	}

	uW0eIVjy.s7DDkIlY[0] &^= d4Lnax - 1

	if debugFloat {
		 /*line :1*/ uW0eIVjy.lxB3waL2v()
	}
}

func (uW0eIVjy *WH8dWN) vTdcul9NY2X6(s_9Ih_ bool, pmEgen2 uint64) *WH8dWN {
	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M = 64
	}
	uW0eIVjy.ddq1eQm = Exact
	uW0eIVjy.g2AL4u0IBtd = s_9Ih_
	if pmEgen2 == 0 {
		uW0eIVjy.fNeR0A = zero
		return uW0eIVjy
	}

	uW0eIVjy.fNeR0A = finite
	nR7KU4mGsdy :=  /*line :1*/ bits.HYxjvQ(pmEgen2)
	uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.glzD00sysLql(pmEgen2 <<  /*line :1*/ uint(nR7KU4mGsdy))
	uW0eIVjy.khThBIzMZ =  /*line :1*/ int32(64 - nR7KU4mGsdy)
	if uW0eIVjy.iaaUnOsr7_M < 64 {
		 /*line :1*/ uW0eIVjy.jrT168o(0)
	}
	return uW0eIVjy
}

func (uW0eIVjy *WH8dWN) SetUint64(pmEgen2 uint64) *WH8dWN {
	return  /*line :1*/ uW0eIVjy.vTdcul9NY2X6(false, pmEgen2)
}

func (uW0eIVjy *WH8dWN) SetInt64(pmEgen2 int64) *WH8dWN {
	sOU3zRv := pmEgen2
	if sOU3zRv < 0 {
		sOU3zRv = -sOU3zRv
	}

	return  /*line :1*/ uW0eIVjy.vTdcul9NY2X6(pmEgen2 < 0,  /*line :1*/ uint64(sOU3zRv))
}

func (uW0eIVjy *WH8dWN) SetFloat64(pmEgen2 float64) *WH8dWN {
	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M = 53
	}
	if  /*line :1*/ math.OIdLpDqq(pmEgen2) {
		 /*line :1*/ panic(At_rvS_hmcW{"Float.SetFloat64(NaN)"})
	}
	uW0eIVjy.ddq1eQm = Exact
	uW0eIVjy.g2AL4u0IBtd =  /*line :1*/ math.GLZFJA9L9(pmEgen2)
	if pmEgen2 == 0 {
		uW0eIVjy.fNeR0A = zero
		return uW0eIVjy
	}
	if  /*line :1*/ math.ME1vzpD5wcr(pmEgen2, 0) {
		uW0eIVjy.fNeR0A = inf
		return uW0eIVjy
	}

	uW0eIVjy.fNeR0A = finite
	aP5YWev, xXsI4WC :=  /*line :1*/ math.CuUl2eV(pmEgen2)
	uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.glzD00sysLql(1<<63 |  /*line :1*/ math.GKIRmoHE(aP5YWev)<<11)
	uW0eIVjy.khThBIzMZ =  /*line :1*/ int32(xXsI4WC)
	if uW0eIVjy.iaaUnOsr7_M < 53 {
		 /*line :1*/ uW0eIVjy.jrT168o(0)
	}
	return uW0eIVjy
}

func bqzcoCe14(gS2TXhHpYaX4 g3X9fa) int64 {
	if debugFloat && ( /*line :1*/ len(gS2TXhHpYaX4) == 0 || gS2TXhHpYaX4[ /*line :1*/ len(gS2TXhHpYaX4)-1] == 0) {
		 /*line :1*/ panic("msw of mantissa is 0")
	}
	nR7KU4mGsdy :=  /*line :1*/ bUgMSHLa(gS2TXhHpYaX4[ /*line :1*/ len(gS2TXhHpYaX4)-1])
	if nR7KU4mGsdy > 0 {
		vKNCqP0 :=  /*line :1*/ vdi3DYM4mY(gS2TXhHpYaX4, gS2TXhHpYaX4, nR7KU4mGsdy)
		if debugFloat && vKNCqP0 != 0 {
			 /*line :1*/ panic("nlz or shlVU incorrect")
		}
	}
	return  /*line :1*/ int64(nR7KU4mGsdy)
}

func (uW0eIVjy *WH8dWN) SetInt(pmEgen2 *ZMtGuo) *WH8dWN {

	uopo38aVK_H :=  /*line :1*/ uint32( /*line :1*/ pmEgen2.BitLen())
	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ cLms9Q(uopo38aVK_H, 64)
	}
	uW0eIVjy.ddq1eQm = Exact
	uW0eIVjy.g2AL4u0IBtd = pmEgen2.hCTOmp0gckSa
	if  /*line :1*/ len(pmEgen2.qho77PBKF) == 0 {
		uW0eIVjy.fNeR0A = zero
		return uW0eIVjy
	}

	uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.bj0nVC(pmEgen2.qho77PBKF)
	 /*line :1*/ bqzcoCe14(uW0eIVjy.s7DDkIlY)
	 /*line :1*/ uW0eIVjy.pKBlNY( /*line :1*/ int64(uopo38aVK_H), 0)
	return uW0eIVjy
}

func (uW0eIVjy *WH8dWN) SetRat(pmEgen2 *FGtP7gcPU) *WH8dWN {
	if  /*line :1*/ pmEgen2.IsInt() {
		return  /*line :1*/ uW0eIVjy.SetInt( /*line :1*/ pmEgen2.Num())
	}
	var sza5atF, kWa1YrIHZo WH8dWN
	 /*line :1*/ sza5atF.SetInt( /*line :1*/ pmEgen2.Num())
	 /*line :1*/ kWa1YrIHZo.SetInt( /*line :1*/ pmEgen2.Denom())
	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ cLms9Q(sza5atF.iaaUnOsr7_M, kWa1YrIHZo.iaaUnOsr7_M)
	}
	return  /*line :1*/ uW0eIVjy.Quo(&sza5atF, &kWa1YrIHZo)
}

func (uW0eIVjy *WH8dWN) SetInf(ozqm8AUMbk bool) *WH8dWN {
	uW0eIVjy.ddq1eQm = Exact
	uW0eIVjy.fNeR0A = inf
	uW0eIVjy.g2AL4u0IBtd = ozqm8AUMbk
	return uW0eIVjy
}

func (uW0eIVjy *WH8dWN) Set(pmEgen2 *WH8dWN) *WH8dWN {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}
	uW0eIVjy.ddq1eQm = Exact
	if uW0eIVjy != pmEgen2 {
		uW0eIVjy.fNeR0A = pmEgen2.fNeR0A
		uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd
		if pmEgen2.fNeR0A == finite {
			uW0eIVjy.khThBIzMZ = pmEgen2.khThBIzMZ
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.bj0nVC(pmEgen2.s7DDkIlY)
		}
		if uW0eIVjy.iaaUnOsr7_M == 0 {
			uW0eIVjy.iaaUnOsr7_M = pmEgen2.iaaUnOsr7_M
		} else if uW0eIVjy.iaaUnOsr7_M < pmEgen2.iaaUnOsr7_M {
			 /*line :1*/ uW0eIVjy.jrT168o(0)
		}
	}
	return uW0eIVjy
}

func (uW0eIVjy *WH8dWN) Copy(pmEgen2 *WH8dWN) *WH8dWN {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}
	if uW0eIVjy != pmEgen2 {
		uW0eIVjy.iaaUnOsr7_M = pmEgen2.iaaUnOsr7_M
		uW0eIVjy.hNG8iJiHdkxT = pmEgen2.hNG8iJiHdkxT
		uW0eIVjy.ddq1eQm = pmEgen2.ddq1eQm
		uW0eIVjy.fNeR0A = pmEgen2.fNeR0A
		uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd
		if uW0eIVjy.fNeR0A == finite {
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.bj0nVC(pmEgen2.s7DDkIlY)
			uW0eIVjy.khThBIzMZ = pmEgen2.khThBIzMZ
		}
	}
	return uW0eIVjy
}

func aU8KEbpZl5(pmEgen2 g3X9fa) uint32 {
	aepbxLiWOZ :=  /*line :1*/ len(pmEgen2) - 1
	if aepbxLiWOZ < 0 {
		return 0
	}
	if debugFloat && pmEgen2[aepbxLiWOZ]&(1<<(_W-1)) == 0 {
		 /*line :1*/ panic("x not normalized")
	}
	switch _W {
	case 32:
		return  /*line :1*/ uint32(pmEgen2[aepbxLiWOZ])
	case 64:
		return  /*line :1*/ uint32(pmEgen2[aepbxLiWOZ] >> 32)
	}
	 /*line :1*/ panic("unreachable")
}

func usOardn(pmEgen2 g3X9fa) uint64 {
	aepbxLiWOZ :=  /*line :1*/ len(pmEgen2) - 1
	if aepbxLiWOZ < 0 {
		return 0
	}
	if debugFloat && pmEgen2[aepbxLiWOZ]&(1<<(_W-1)) == 0 {
		 /*line :1*/ panic("x not normalized")
	}
	switch _W {
	case 32:
		iQKtjmO6 :=  /*line :1*/ uint64(pmEgen2[aepbxLiWOZ]) << 32
		if aepbxLiWOZ > 0 {
			iQKtjmO6 |=  /*line :1*/ uint64(pmEgen2[aepbxLiWOZ-1])
		}
		return iQKtjmO6
	case 64:
		return  /*line :1*/ uint64(pmEgen2[aepbxLiWOZ])
	}
	 /*line :1*/ panic("unreachable")
}

func (pmEgen2 *WH8dWN) Uint64() (uint64, IG11Ul7qo) {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	switch pmEgen2.fNeR0A {
	case finite:
		if pmEgen2.g2AL4u0IBtd {
			return 0, Above
		}

		if pmEgen2.khThBIzMZ <= 0 {

			return 0, Below
		}

		if pmEgen2.khThBIzMZ <= 64 {

			sOU3zRv :=  /*line :1*/ usOardn(pmEgen2.s7DDkIlY) >> (64 -  /*line :1*/ uint32(pmEgen2.khThBIzMZ))
			if  /*line :1*/ pmEgen2.MinPrec() <= 64 {
				return sOU3zRv, Exact
			}
			return sOU3zRv, Below
		}

		return math.MaxUint64, Below

	case zero:
		return 0, Exact

	case inf:
		if pmEgen2.g2AL4u0IBtd {
			return 0, Above
		}
		return math.MaxUint64, Below
	}

	 /*line :1*/ panic("unreachable")
}

func (pmEgen2 *WH8dWN) Int64() (int64, IG11Ul7qo) {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	switch pmEgen2.fNeR0A {
	case finite:

		fDECsVlO :=  /*line :1*/ dBh3XtAgq(pmEgen2.g2AL4u0IBtd)
		if pmEgen2.khThBIzMZ <= 0 {

			return 0, fDECsVlO
		}

		if pmEgen2.khThBIzMZ <= 63 {

			aepbxLiWOZ :=  /*line :1*/ int64( /*line :1*/ usOardn(pmEgen2.s7DDkIlY) >> (64 -  /*line :1*/ uint32(pmEgen2.khThBIzMZ)))
			if pmEgen2.g2AL4u0IBtd {
				aepbxLiWOZ = -aepbxLiWOZ
			}
			if  /*line :1*/ pmEgen2.MinPrec() <=  /*line :1*/ uint(pmEgen2.khThBIzMZ) {
				return aepbxLiWOZ, Exact
			}
			return aepbxLiWOZ, fDECsVlO
		}
		if pmEgen2.g2AL4u0IBtd {

			if pmEgen2.khThBIzMZ == 64 &&  /*line :1*/ pmEgen2.MinPrec() == 1 {
				fDECsVlO = Exact
			}
			return math.MinInt64, fDECsVlO
		}

		return math.MaxInt64, Below

	case zero:
		return 0, Exact

	case inf:
		if pmEgen2.g2AL4u0IBtd {
			return math.MinInt64, Above
		}
		return math.MaxInt64, Below
	}

	 /*line :1*/ panic("unreachable")
}

func (pmEgen2 *WH8dWN) Float32() (float32, IG11Ul7qo) {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	switch pmEgen2.fNeR0A {
	case finite:

		const (
			fbits	= 32
			mbits	= 23
			ebits	= fbits - mbits - 1
			bias	= 1<<(ebits-1) - 1
			dmin	= 1 - bias - mbits
			emin	= 1 - bias
			emax	= bias
		)

		vRGCE256ba4X := pmEgen2.khThBIzMZ - 1

		djOsNJsLfH5U := mbits + 1
		if vRGCE256ba4X < emin {

			djOsNJsLfH5U = mbits + 1 - emin +  /*line :1*/ int(vRGCE256ba4X)

			if djOsNJsLfH5U < 0 || djOsNJsLfH5U == 0 &&  /*line :1*/ pmEgen2.s7DDkIlY.cbgX3ksmdgax( /*line :1*/ uint( /*line :1*/ len(pmEgen2.s7DDkIlY))*_W-1) == 0 {

				if pmEgen2.g2AL4u0IBtd {
					var uW0eIVjy float32
					return -uW0eIVjy, Above
				}
				return 0.0, Below
			}

			if djOsNJsLfH5U == 0 {
				if pmEgen2.g2AL4u0IBtd {
					return -math.SmallestNonzeroFloat32, Below
				}
				return math.SmallestNonzeroFloat32, Above
			}
		}

		var vFx5p_cm WH8dWN
		vFx5p_cm.iaaUnOsr7_M =  /*line :1*/ uint32(djOsNJsLfH5U)
		 /*line :1*/ vFx5p_cm.Set(pmEgen2)
		vRGCE256ba4X = vFx5p_cm.khThBIzMZ - 1

		if vFx5p_cm.fNeR0A == inf || vRGCE256ba4X > emax {

			if pmEgen2.g2AL4u0IBtd {
				return  /*line :1*/ float32( /*line :1*/ math.FSug4WHZSBwz(-1)), Below
			}
			return  /*line :1*/ float32( /*line :1*/ math.FSug4WHZSBwz(+1)), Above
		}

		var e2mU4H, um1iaA, kfG4qg_2AdyH uint32
		if pmEgen2.g2AL4u0IBtd {
			e2mU4H = 1 << (fbits - 1)
		}

		if vRGCE256ba4X < emin {

			djOsNJsLfH5U = mbits + 1 - emin +  /*line :1*/ int(vRGCE256ba4X)
			kfG4qg_2AdyH =  /*line :1*/ aU8KEbpZl5(vFx5p_cm.s7DDkIlY) >>  /*line :1*/ uint(fbits-djOsNJsLfH5U)
		} else {

			um1iaA =  /*line :1*/ uint32(vRGCE256ba4X+bias) << mbits
			kfG4qg_2AdyH =  /*line :1*/ aU8KEbpZl5(vFx5p_cm.s7DDkIlY) >> ebits & (1<<mbits - 1)
		}

		return  /*line :1*/ math.AviWp5b0(e2mU4H | um1iaA | kfG4qg_2AdyH), vFx5p_cm.ddq1eQm

	case zero:
		if pmEgen2.g2AL4u0IBtd {
			var uW0eIVjy float32
			return -uW0eIVjy, Exact
		}
		return 0.0, Exact

	case inf:
		if pmEgen2.g2AL4u0IBtd {
			return  /*line :1*/ float32( /*line :1*/ math.FSug4WHZSBwz(-1)), Exact
		}
		return  /*line :1*/ float32( /*line :1*/ math.FSug4WHZSBwz(+1)), Exact
	}

	 /*line :1*/ panic("unreachable")
}

func (pmEgen2 *WH8dWN) Float64() (float64, IG11Ul7qo) {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	switch pmEgen2.fNeR0A {
	case finite:

		const (
			fbits	= 64
			mbits	= 52
			ebits	= fbits - mbits - 1
			bias	= 1<<(ebits-1) - 1
			dmin	= 1 - bias - mbits
			emin	= 1 - bias
			emax	= bias
		)

		vRGCE256ba4X := pmEgen2.khThBIzMZ - 1

		djOsNJsLfH5U := mbits + 1
		if vRGCE256ba4X < emin {

			djOsNJsLfH5U = mbits + 1 - emin +  /*line :1*/ int(vRGCE256ba4X)

			if djOsNJsLfH5U < 0 || djOsNJsLfH5U == 0 &&  /*line :1*/ pmEgen2.s7DDkIlY.cbgX3ksmdgax( /*line :1*/ uint( /*line :1*/ len(pmEgen2.s7DDkIlY))*_W-1) == 0 {

				if pmEgen2.g2AL4u0IBtd {
					var uW0eIVjy float64
					return -uW0eIVjy, Above
				}
				return 0.0, Below
			}

			if djOsNJsLfH5U == 0 {
				if pmEgen2.g2AL4u0IBtd {
					return -math.SmallestNonzeroFloat64, Below
				}
				return math.SmallestNonzeroFloat64, Above
			}
		}

		var vFx5p_cm WH8dWN
		vFx5p_cm.iaaUnOsr7_M =  /*line :1*/ uint32(djOsNJsLfH5U)
		 /*line :1*/ vFx5p_cm.Set(pmEgen2)
		vRGCE256ba4X = vFx5p_cm.khThBIzMZ - 1

		if vFx5p_cm.fNeR0A == inf || vRGCE256ba4X > emax {

			if pmEgen2.g2AL4u0IBtd {
				return  /*line :1*/ math.FSug4WHZSBwz(-1), Below
			}
			return  /*line :1*/ math.FSug4WHZSBwz(+1), Above
		}

		var e2mU4H, um1iaA, kfG4qg_2AdyH uint64
		if pmEgen2.g2AL4u0IBtd {
			e2mU4H = 1 << (fbits - 1)
		}

		if vRGCE256ba4X < emin {

			djOsNJsLfH5U = mbits + 1 - emin +  /*line :1*/ int(vRGCE256ba4X)
			kfG4qg_2AdyH =  /*line :1*/ usOardn(vFx5p_cm.s7DDkIlY) >>  /*line :1*/ uint(fbits-djOsNJsLfH5U)
		} else {

			um1iaA =  /*line :1*/ uint64(vRGCE256ba4X+bias) << mbits
			kfG4qg_2AdyH =  /*line :1*/ usOardn(vFx5p_cm.s7DDkIlY) >> ebits & (1<<mbits - 1)
		}

		return  /*line :1*/ math.QUB2Kf63p(e2mU4H | um1iaA | kfG4qg_2AdyH), vFx5p_cm.ddq1eQm

	case zero:
		if pmEgen2.g2AL4u0IBtd {
			var uW0eIVjy float64
			return -uW0eIVjy, Exact
		}
		return 0.0, Exact

	case inf:
		if pmEgen2.g2AL4u0IBtd {
			return  /*line :1*/ math.FSug4WHZSBwz(-1), Exact
		}
		return  /*line :1*/ math.FSug4WHZSBwz(+1), Exact
	}

	 /*line :1*/ panic("unreachable")
}

func (pmEgen2 *WH8dWN) Int(uW0eIVjy *ZMtGuo) (*ZMtGuo, IG11Ul7qo) {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	if uW0eIVjy == nil && pmEgen2.fNeR0A <= finite {
		uW0eIVjy =  /*line :1*/ new(ZMtGuo)
	}

	switch pmEgen2.fNeR0A {
	case finite:

		fDECsVlO :=  /*line :1*/ dBh3XtAgq(pmEgen2.g2AL4u0IBtd)
		if pmEgen2.khThBIzMZ <= 0 {

			return  /*line :1*/ uW0eIVjy.SetInt64(0), fDECsVlO
		}

		m2sPeu9 :=  /*line :1*/ uint( /*line :1*/ len(pmEgen2.s7DDkIlY)) * _W
		xXsI4WC :=  /*line :1*/ uint(pmEgen2.khThBIzMZ)
		if  /*line :1*/ pmEgen2.MinPrec() <= xXsI4WC {
			fDECsVlO = Exact
		}

		if uW0eIVjy == nil {
			uW0eIVjy =  /*line :1*/ new(ZMtGuo)
		}
		uW0eIVjy.hCTOmp0gckSa = pmEgen2.g2AL4u0IBtd
		switch {
		case xXsI4WC > m2sPeu9:
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.z839tk6CmDHT(pmEgen2.s7DDkIlY, xXsI4WC-m2sPeu9)
		default:
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.bj0nVC(pmEgen2.s7DDkIlY)
		case xXsI4WC < m2sPeu9:
			uW0eIVjy.qho77PBKF =  /*line :1*/ uW0eIVjy.qho77PBKF.qzOaHmR(pmEgen2.s7DDkIlY, m2sPeu9-xXsI4WC)
		}
		return uW0eIVjy, fDECsVlO

	case zero:
		return  /*line :1*/ uW0eIVjy.SetInt64(0), Exact

	case inf:
		return nil,  /*line :1*/ dBh3XtAgq(pmEgen2.g2AL4u0IBtd)
	}

	 /*line :1*/ panic("unreachable")
}

func (pmEgen2 *WH8dWN) Rat(uW0eIVjy *FGtP7gcPU) (*FGtP7gcPU, IG11Ul7qo) {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	if uW0eIVjy == nil && pmEgen2.fNeR0A <= finite {
		uW0eIVjy =  /*line :1*/ new(FGtP7gcPU)
	}

	switch pmEgen2.fNeR0A {
	case finite:

		m2sPeu9 :=  /*line :1*/ int32( /*line :1*/ len(pmEgen2.s7DDkIlY)) * _W

		uW0eIVjy.u67K4Sv.hCTOmp0gckSa = pmEgen2.g2AL4u0IBtd
		switch {
		case pmEgen2.khThBIzMZ > m2sPeu9:
			uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.z839tk6CmDHT(pmEgen2.s7DDkIlY,  /*line :1*/ uint(pmEgen2.khThBIzMZ-m2sPeu9))
			uW0eIVjy._WnhHVFvm7.qho77PBKF = uW0eIVjy._WnhHVFvm7.qho77PBKF[:0]

		default:
			uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.bj0nVC(pmEgen2.s7DDkIlY)
			uW0eIVjy._WnhHVFvm7.qho77PBKF = uW0eIVjy._WnhHVFvm7.qho77PBKF[:0]

		case pmEgen2.khThBIzMZ < m2sPeu9:
			uW0eIVjy.u67K4Sv.qho77PBKF =  /*line :1*/ uW0eIVjy.u67K4Sv.qho77PBKF.bj0nVC(pmEgen2.s7DDkIlY)
			amoa4Px :=  /*line :1*/ uW0eIVjy._WnhHVFvm7.qho77PBKF.glzD00sysLql(1)
			uW0eIVjy._WnhHVFvm7.qho77PBKF =  /*line :1*/ amoa4Px.z839tk6CmDHT(amoa4Px,  /*line :1*/ uint(m2sPeu9-pmEgen2.khThBIzMZ))
			 /*line :1*/ uW0eIVjy.hitXcy5()
		}
		return uW0eIVjy, Exact

	case zero:
		return  /*line :1*/ uW0eIVjy.SetInt64(0), Exact

	case inf:
		return nil,  /*line :1*/ dBh3XtAgq(pmEgen2.g2AL4u0IBtd)
	}

	 /*line :1*/ panic("unreachable")
}

func (uW0eIVjy *WH8dWN) Abs(pmEgen2 *WH8dWN) *WH8dWN {
	 /*line :1*/ uW0eIVjy.Set(pmEgen2)
	uW0eIVjy.g2AL4u0IBtd = false
	return uW0eIVjy
}

func (uW0eIVjy *WH8dWN) Neg(pmEgen2 *WH8dWN) *WH8dWN {
	 /*line :1*/ uW0eIVjy.Set(pmEgen2)
	uW0eIVjy.g2AL4u0IBtd = !uW0eIVjy.g2AL4u0IBtd
	return uW0eIVjy
}

func e7TvCJmNN(pmEgen2, oxFS5wa5 *WH8dWN) {
	if !debugFloat {

		 /*line :1*/ panic("validateBinaryOperands called but debugFloat is not set")
	}
	if  /*line :1*/ len(pmEgen2.s7DDkIlY) == 0 {
		 /*line :1*/ panic("empty mantissa for x")
	}
	if  /*line :1*/ len(oxFS5wa5.s7DDkIlY) == 0 {
		 /*line :1*/ panic("empty mantissa for y")
	}
}

func (uW0eIVjy *WH8dWN) tdlGzH99r(pmEgen2, oxFS5wa5 *WH8dWN) {

	if debugFloat {
		 /*line :1*/ e7TvCJmNN(pmEgen2, oxFS5wa5)
	}

	wD9HOPBI :=  /*line :1*/ int64(pmEgen2.khThBIzMZ) -  /*line :1*/ int64( /*line :1*/ len(pmEgen2.s7DDkIlY))*_W
	mgBvUzvVhEbq :=  /*line :1*/ int64(oxFS5wa5.khThBIzMZ) -  /*line :1*/ int64( /*line :1*/ len(oxFS5wa5.s7DDkIlY))*_W

	nCOfjN0 :=  /*line :1*/ ap52A4djK(uW0eIVjy.s7DDkIlY, pmEgen2.s7DDkIlY) ||  /*line :1*/ ap52A4djK(uW0eIVjy.s7DDkIlY, oxFS5wa5.s7DDkIlY)

	switch {
	case wD9HOPBI < mgBvUzvVhEbq:
		if nCOfjN0 {
			amoa4Px :=  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(oxFS5wa5.s7DDkIlY,  /*line :1*/ uint(mgBvUzvVhEbq-wD9HOPBI))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.jXv7nnhY_(pmEgen2.s7DDkIlY, amoa4Px)
		} else {
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.z839tk6CmDHT(oxFS5wa5.s7DDkIlY,  /*line :1*/ uint(mgBvUzvVhEbq-wD9HOPBI))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.jXv7nnhY_(pmEgen2.s7DDkIlY, uW0eIVjy.s7DDkIlY)
		}
	default:

		uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.jXv7nnhY_(pmEgen2.s7DDkIlY, oxFS5wa5.s7DDkIlY)
	case wD9HOPBI > mgBvUzvVhEbq:
		if nCOfjN0 {
			amoa4Px :=  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(pmEgen2.s7DDkIlY,  /*line :1*/ uint(wD9HOPBI-mgBvUzvVhEbq))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.jXv7nnhY_(amoa4Px, oxFS5wa5.s7DDkIlY)
		} else {
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.z839tk6CmDHT(pmEgen2.s7DDkIlY,  /*line :1*/ uint(wD9HOPBI-mgBvUzvVhEbq))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.jXv7nnhY_(uW0eIVjy.s7DDkIlY, oxFS5wa5.s7DDkIlY)
		}
		wD9HOPBI = mgBvUzvVhEbq
	}

	 /*line :1*/ uW0eIVjy.pKBlNY(wD9HOPBI+ /*line :1*/ int64( /*line :1*/ len(uW0eIVjy.s7DDkIlY))*_W- /*line :1*/ bqzcoCe14(uW0eIVjy.s7DDkIlY), 0)
}

func (uW0eIVjy *WH8dWN) vESNS9z_(pmEgen2, oxFS5wa5 *WH8dWN) {

	if debugFloat {
		 /*line :1*/ e7TvCJmNN(pmEgen2, oxFS5wa5)
	}

	wD9HOPBI :=  /*line :1*/ int64(pmEgen2.khThBIzMZ) -  /*line :1*/ int64( /*line :1*/ len(pmEgen2.s7DDkIlY))*_W
	mgBvUzvVhEbq :=  /*line :1*/ int64(oxFS5wa5.khThBIzMZ) -  /*line :1*/ int64( /*line :1*/ len(oxFS5wa5.s7DDkIlY))*_W

	nCOfjN0 :=  /*line :1*/ ap52A4djK(uW0eIVjy.s7DDkIlY, pmEgen2.s7DDkIlY) ||  /*line :1*/ ap52A4djK(uW0eIVjy.s7DDkIlY, oxFS5wa5.s7DDkIlY)

	switch {
	case wD9HOPBI < mgBvUzvVhEbq:
		if nCOfjN0 {
			amoa4Px :=  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(oxFS5wa5.s7DDkIlY,  /*line :1*/ uint(mgBvUzvVhEbq-wD9HOPBI))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ amoa4Px.lt4VKILNCo(pmEgen2.s7DDkIlY, amoa4Px)
		} else {
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.z839tk6CmDHT(oxFS5wa5.s7DDkIlY,  /*line :1*/ uint(mgBvUzvVhEbq-wD9HOPBI))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.lt4VKILNCo(pmEgen2.s7DDkIlY, uW0eIVjy.s7DDkIlY)
		}
	default:

		uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.lt4VKILNCo(pmEgen2.s7DDkIlY, oxFS5wa5.s7DDkIlY)
	case wD9HOPBI > mgBvUzvVhEbq:
		if nCOfjN0 {
			amoa4Px :=  /*line :1*/ g3X9fa(nil).z839tk6CmDHT(pmEgen2.s7DDkIlY,  /*line :1*/ uint(wD9HOPBI-mgBvUzvVhEbq))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ amoa4Px.lt4VKILNCo(amoa4Px, oxFS5wa5.s7DDkIlY)
		} else {
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.z839tk6CmDHT(pmEgen2.s7DDkIlY,  /*line :1*/ uint(wD9HOPBI-mgBvUzvVhEbq))
			uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.lt4VKILNCo(uW0eIVjy.s7DDkIlY, oxFS5wa5.s7DDkIlY)
		}
		wD9HOPBI = mgBvUzvVhEbq
	}

	if  /*line :1*/ len(uW0eIVjy.s7DDkIlY) == 0 {
		uW0eIVjy.ddq1eQm = Exact
		uW0eIVjy.fNeR0A = zero
		uW0eIVjy.g2AL4u0IBtd = false
		return
	}

	 /*line :1*/ uW0eIVjy.pKBlNY(wD9HOPBI+ /*line :1*/ int64( /*line :1*/ len(uW0eIVjy.s7DDkIlY))*_W- /*line :1*/ bqzcoCe14(uW0eIVjy.s7DDkIlY), 0)
}

func (uW0eIVjy *WH8dWN) vST1Fx6GnP(pmEgen2, oxFS5wa5 *WH8dWN) {
	if debugFloat {
		 /*line :1*/ e7TvCJmNN(pmEgen2, oxFS5wa5)
	}

	vRGCE256ba4X :=  /*line :1*/ int64(pmEgen2.khThBIzMZ) +  /*line :1*/ int64(oxFS5wa5.khThBIzMZ)
	if pmEgen2 == oxFS5wa5 {
		uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.pJ61taToc(pmEgen2.s7DDkIlY)
	} else {
		uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.d22x6Ygoc80O(pmEgen2.s7DDkIlY, oxFS5wa5.s7DDkIlY)
	}
	 /*line :1*/ uW0eIVjy.pKBlNY(vRGCE256ba4X- /*line :1*/ bqzcoCe14(uW0eIVjy.s7DDkIlY), 0)
}

func (uW0eIVjy *WH8dWN) xX0MItaD(pmEgen2, oxFS5wa5 *WH8dWN) {
	if debugFloat {
		 /*line :1*/ e7TvCJmNN(pmEgen2, oxFS5wa5)
	}

	h_Wr_yqq :=  /*line :1*/ int(uW0eIVjy.iaaUnOsr7_M/_W) + 1

	vPeJudD6 := pmEgen2.s7DDkIlY
	if s31g8J := h_Wr_yqq -  /*line :1*/ len(pmEgen2.s7DDkIlY) +  /*line :1*/ len(oxFS5wa5.s7DDkIlY); s31g8J > 0 {

		vPeJudD6 =  /*line :1*/ make(g3X9fa,  /*line :1*/ len(pmEgen2.s7DDkIlY)+s31g8J)
		 /*line :1*/ copy(vPeJudD6[s31g8J:], pmEgen2.s7DDkIlY)
	}

	s31g8J :=  /*line :1*/ len(vPeJudD6) -  /*line :1*/ len(oxFS5wa5.s7DDkIlY)

	var vFx5p_cm g3X9fa
	uW0eIVjy.s7DDkIlY, vFx5p_cm =  /*line :1*/ uW0eIVjy.s7DDkIlY.xOZxoLyl(nil, vPeJudD6, oxFS5wa5.s7DDkIlY)
	vRGCE256ba4X :=  /*line :1*/ int64(pmEgen2.khThBIzMZ) -  /*line :1*/ int64(oxFS5wa5.khThBIzMZ) -  /*line :1*/ int64(s31g8J- /*line :1*/ len(uW0eIVjy.s7DDkIlY))*_W

	var c8rDOxCvz_ZX uint
	if  /*line :1*/ len(vFx5p_cm) > 0 {
		c8rDOxCvz_ZX = 1
	}

	 /*line :1*/ uW0eIVjy.pKBlNY(vRGCE256ba4X- /*line :1*/ bqzcoCe14(uW0eIVjy.s7DDkIlY), c8rDOxCvz_ZX)
}

func (pmEgen2 *WH8dWN) fBXpyjpFe(oxFS5wa5 *WH8dWN) int {
	if debugFloat {
		 /*line :1*/ e7TvCJmNN(pmEgen2, oxFS5wa5)
	}

	switch {
	case pmEgen2.khThBIzMZ < oxFS5wa5.khThBIzMZ:
		return -1
	case pmEgen2.khThBIzMZ > oxFS5wa5.khThBIzMZ:
		return +1
	}

	aepbxLiWOZ :=  /*line :1*/ len(pmEgen2.s7DDkIlY)
	g77TAi :=  /*line :1*/ len(oxFS5wa5.s7DDkIlY)
	for aepbxLiWOZ > 0 || g77TAi > 0 {
		var hQlIMC14oKX, jDEhl_rqJazr VYe6D0
		if aepbxLiWOZ > 0 {
			aepbxLiWOZ--
			hQlIMC14oKX = pmEgen2.s7DDkIlY[aepbxLiWOZ]
		}
		if g77TAi > 0 {
			g77TAi--
			jDEhl_rqJazr = oxFS5wa5.s7DDkIlY[g77TAi]
		}
		switch {
		case hQlIMC14oKX < jDEhl_rqJazr:
			return -1
		case hQlIMC14oKX > jDEhl_rqJazr:
			return +1
		}
	}

	return 0
}

func (uW0eIVjy *WH8dWN) Add(pmEgen2, oxFS5wa5 *WH8dWN) *WH8dWN {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
		 /*line :1*/ oxFS5wa5.lxB3waL2v()
	}

	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ cLms9Q(pmEgen2.iaaUnOsr7_M, oxFS5wa5.iaaUnOsr7_M)
	}

	if pmEgen2.fNeR0A == finite && oxFS5wa5.fNeR0A == finite {

		zi9dcFUDVO := oxFS5wa5.g2AL4u0IBtd

		uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd
		if pmEgen2.g2AL4u0IBtd == zi9dcFUDVO {

			 /*line :1*/ uW0eIVjy.tdlGzH99r(pmEgen2, oxFS5wa5)
		} else {

			if  /*line :1*/ pmEgen2.fBXpyjpFe(oxFS5wa5) > 0 {
				 /*line :1*/ uW0eIVjy.vESNS9z_(pmEgen2, oxFS5wa5)
			} else {
				uW0eIVjy.g2AL4u0IBtd = !uW0eIVjy.g2AL4u0IBtd
				 /*line :1*/ uW0eIVjy.vESNS9z_(oxFS5wa5, pmEgen2)
			}
		}
		if uW0eIVjy.fNeR0A == zero && uW0eIVjy.hNG8iJiHdkxT == ToNegativeInf && uW0eIVjy.ddq1eQm == Exact {
			uW0eIVjy.g2AL4u0IBtd = true
		}
		return uW0eIVjy
	}

	if pmEgen2.fNeR0A == inf && oxFS5wa5.fNeR0A == inf && pmEgen2.g2AL4u0IBtd != oxFS5wa5.g2AL4u0IBtd {

		uW0eIVjy.ddq1eQm = Exact
		uW0eIVjy.fNeR0A = zero
		uW0eIVjy.g2AL4u0IBtd = false
		 /*line :1*/ panic(At_rvS_hmcW{"addition of infinities with opposite signs"})
	}

	if pmEgen2.fNeR0A == zero && oxFS5wa5.fNeR0A == zero {

		uW0eIVjy.ddq1eQm = Exact
		uW0eIVjy.fNeR0A = zero
		uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd && oxFS5wa5.g2AL4u0IBtd
		return uW0eIVjy
	}

	if pmEgen2.fNeR0A == inf || oxFS5wa5.fNeR0A == zero {

		return  /*line :1*/ uW0eIVjy.Set(pmEgen2)
	}

	return  /*line :1*/ uW0eIVjy.Set(oxFS5wa5)
}

func (uW0eIVjy *WH8dWN) Sub(pmEgen2, oxFS5wa5 *WH8dWN) *WH8dWN {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
		 /*line :1*/ oxFS5wa5.lxB3waL2v()
	}

	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ cLms9Q(pmEgen2.iaaUnOsr7_M, oxFS5wa5.iaaUnOsr7_M)
	}

	if pmEgen2.fNeR0A == finite && oxFS5wa5.fNeR0A == finite {

		zi9dcFUDVO := oxFS5wa5.g2AL4u0IBtd
		uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd
		if pmEgen2.g2AL4u0IBtd != zi9dcFUDVO {

			 /*line :1*/ uW0eIVjy.tdlGzH99r(pmEgen2, oxFS5wa5)
		} else {

			if  /*line :1*/ pmEgen2.fBXpyjpFe(oxFS5wa5) > 0 {
				 /*line :1*/ uW0eIVjy.vESNS9z_(pmEgen2, oxFS5wa5)
			} else {
				uW0eIVjy.g2AL4u0IBtd = !uW0eIVjy.g2AL4u0IBtd
				 /*line :1*/ uW0eIVjy.vESNS9z_(oxFS5wa5, pmEgen2)
			}
		}
		if uW0eIVjy.fNeR0A == zero && uW0eIVjy.hNG8iJiHdkxT == ToNegativeInf && uW0eIVjy.ddq1eQm == Exact {
			uW0eIVjy.g2AL4u0IBtd = true
		}
		return uW0eIVjy
	}

	if pmEgen2.fNeR0A == inf && oxFS5wa5.fNeR0A == inf && pmEgen2.g2AL4u0IBtd == oxFS5wa5.g2AL4u0IBtd {

		uW0eIVjy.ddq1eQm = Exact
		uW0eIVjy.fNeR0A = zero
		uW0eIVjy.g2AL4u0IBtd = false
		 /*line :1*/ panic(At_rvS_hmcW{"subtraction of infinities with equal signs"})
	}

	if pmEgen2.fNeR0A == zero && oxFS5wa5.fNeR0A == zero {

		uW0eIVjy.ddq1eQm = Exact
		uW0eIVjy.fNeR0A = zero
		uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd && !oxFS5wa5.g2AL4u0IBtd
		return uW0eIVjy
	}

	if pmEgen2.fNeR0A == inf || oxFS5wa5.fNeR0A == zero {

		return  /*line :1*/ uW0eIVjy.Set(pmEgen2)
	}

	return  /*line :1*/ uW0eIVjy.Neg(oxFS5wa5)
}

func (uW0eIVjy *WH8dWN) Mul(pmEgen2, oxFS5wa5 *WH8dWN) *WH8dWN {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
		 /*line :1*/ oxFS5wa5.lxB3waL2v()
	}

	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ cLms9Q(pmEgen2.iaaUnOsr7_M, oxFS5wa5.iaaUnOsr7_M)
	}

	uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd != oxFS5wa5.g2AL4u0IBtd

	if pmEgen2.fNeR0A == finite && oxFS5wa5.fNeR0A == finite {

		 /*line :1*/ uW0eIVjy.vST1Fx6GnP(pmEgen2, oxFS5wa5)
		return uW0eIVjy
	}

	uW0eIVjy.ddq1eQm = Exact
	if pmEgen2.fNeR0A == zero && oxFS5wa5.fNeR0A == inf || pmEgen2.fNeR0A == inf && oxFS5wa5.fNeR0A == zero {

		uW0eIVjy.fNeR0A = zero
		uW0eIVjy.g2AL4u0IBtd = false
		 /*line :1*/ panic(At_rvS_hmcW{"multiplication of zero with infinity"})
	}

	if pmEgen2.fNeR0A == inf || oxFS5wa5.fNeR0A == inf {

		uW0eIVjy.fNeR0A = inf
		return uW0eIVjy
	}

	uW0eIVjy.fNeR0A = zero
	return uW0eIVjy
}

func (uW0eIVjy *WH8dWN) Quo(pmEgen2, oxFS5wa5 *WH8dWN) *WH8dWN {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
		 /*line :1*/ oxFS5wa5.lxB3waL2v()
	}

	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M =  /*line :1*/ cLms9Q(pmEgen2.iaaUnOsr7_M, oxFS5wa5.iaaUnOsr7_M)
	}

	uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd != oxFS5wa5.g2AL4u0IBtd

	if pmEgen2.fNeR0A == finite && oxFS5wa5.fNeR0A == finite {

		 /*line :1*/ uW0eIVjy.xX0MItaD(pmEgen2, oxFS5wa5)
		return uW0eIVjy
	}

	uW0eIVjy.ddq1eQm = Exact
	if pmEgen2.fNeR0A == zero && oxFS5wa5.fNeR0A == zero || pmEgen2.fNeR0A == inf && oxFS5wa5.fNeR0A == inf {

		uW0eIVjy.fNeR0A = zero
		uW0eIVjy.g2AL4u0IBtd = false
		 /*line :1*/ panic(At_rvS_hmcW{"division of zero by zero or infinity by infinity"})
	}

	if pmEgen2.fNeR0A == zero || oxFS5wa5.fNeR0A == inf {

		uW0eIVjy.fNeR0A = zero
		return uW0eIVjy
	}

	uW0eIVjy.fNeR0A = inf
	return uW0eIVjy
}

func (pmEgen2 *WH8dWN) Cmp(oxFS5wa5 *WH8dWN) int {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
		 /*line :1*/ oxFS5wa5.lxB3waL2v()
	}

	xSaOPd :=  /*line :1*/ pmEgen2.ag3NoaTJ0()
	otj7PW0DHu7n :=  /*line :1*/ oxFS5wa5.ag3NoaTJ0()
	switch {
	case xSaOPd < otj7PW0DHu7n:
		return -1
	case xSaOPd > otj7PW0DHu7n:
		return +1
	}

	switch xSaOPd {
	case -1:
		return  /*line :1*/ oxFS5wa5.fBXpyjpFe(pmEgen2)
	case +1:
		return  /*line :1*/ pmEgen2.fBXpyjpFe(oxFS5wa5)
	}

	return 0
}

func (pmEgen2 *WH8dWN) ag3NoaTJ0() int {
	var gS2TXhHpYaX4 int
	switch pmEgen2.fNeR0A {
	case finite:
		gS2TXhHpYaX4 = 1
	case zero:
		return 0
	case inf:
		gS2TXhHpYaX4 = 2
	}
	if pmEgen2.g2AL4u0IBtd {
		gS2TXhHpYaX4 = -gS2TXhHpYaX4
	}
	return gS2TXhHpYaX4
}

func cLms9Q(pmEgen2, oxFS5wa5 uint32) uint32 {
	if pmEgen2 > oxFS5wa5 {
		return pmEgen2
	}
	return oxFS5wa5
}
