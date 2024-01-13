//line :1
package big

import (
	math "math"
	sync "sync"
)

var qUN6W5TRV9 struct {
	sync.LhBfwn6wa1x
	epRH5PFFiX	*WH8dWN
}

func dg8qlb0yH() *WH8dWN {
	 /*line :1*/ qUN6W5TRV9.Do(func() {
		qUN6W5TRV9.epRH5PFFiX =  /*line :1*/ EfoO1hKDhm0(3.0)
	})
	return qUN6W5TRV9.epRH5PFFiX
}










func (uW0eIVjy *WH8dWN) Sqrt(pmEgen2 *WH8dWN) *WH8dWN {
	if debugFloat {
		 /*line :1*/ pmEgen2.lxB3waL2v()
	}

	if uW0eIVjy.iaaUnOsr7_M == 0 {
		uW0eIVjy.iaaUnOsr7_M = pmEgen2.iaaUnOsr7_M
	}

	if  /*line :1*/ pmEgen2.Sign() == -1 {

		 /*line :1*/ panic(At_rvS_hmcW{"square root of negative operand"})
	}

	if pmEgen2.fNeR0A != finite {
		uW0eIVjy.ddq1eQm = Exact
		uW0eIVjy.fNeR0A = pmEgen2.fNeR0A
		uW0eIVjy.g2AL4u0IBtd = pmEgen2.g2AL4u0IBtd
		return uW0eIVjy
	}

	cZUDpnJQn := uW0eIVjy.iaaUnOsr7_M
	kWa1YrIHZo :=  /*line :1*/ pmEgen2.MantExp(uW0eIVjy)
	uW0eIVjy.iaaUnOsr7_M = cZUDpnJQn

	switch kWa1YrIHZo % 2 {
	case 0:

	case 1:
		uW0eIVjy.khThBIzMZ++
	case -1:
		uW0eIVjy.khThBIzMZ--
	}

	 /*line :1*/ uW0eIVjy.lwFPu4JxXN(uW0eIVjy)

	return  /*line :1*/ uW0eIVjy.SetMantExp(uW0eIVjy, kWa1YrIHZo/2)
}






func (uW0eIVjy *WH8dWN) lwFPu4JxXN(pmEgen2 *WH8dWN) {

	sOU3zRv :=  /*line :1*/ gUZEPIC8uIA(uW0eIVjy.iaaUnOsr7_M)
	iQKtjmO6 :=  /*line :1*/ gUZEPIC8uIA(uW0eIVjy.iaaUnOsr7_M)
	dg8qlb0yH :=  /*line :1*/ dg8qlb0yH()
	kXMQLm5mTkv := func(amoa4Px *WH8dWN) *WH8dWN {
		sOU3zRv.iaaUnOsr7_M = amoa4Px.iaaUnOsr7_M
		iQKtjmO6.iaaUnOsr7_M = amoa4Px.iaaUnOsr7_M
		 /*line :1*/ sOU3zRv.Mul(amoa4Px, amoa4Px)
		 /*line :1*/ sOU3zRv.Mul(pmEgen2, sOU3zRv)
		 /*line :1*/ iQKtjmO6.Sub(dg8qlb0yH, sOU3zRv)
		 /*line :1*/ sOU3zRv.Mul(amoa4Px, iQKtjmO6)
		sOU3zRv.khThBIzMZ--
		return  /*line :1*/ amoa4Px.Set(sOU3zRv)
	}

	daCqrmBWa, _ :=  /*line :1*/ pmEgen2.Float64()
	plHJUlaC4H :=  /*line :1*/ gUZEPIC8uIA(uW0eIVjy.iaaUnOsr7_M)
	 /*line :1*/ plHJUlaC4H.SetFloat64(1 /  /*line :1*/ math.NeXs7bSyfaCD(daCqrmBWa))
	for cZUDpnJQn := uW0eIVjy.iaaUnOsr7_M + 32; plHJUlaC4H.iaaUnOsr7_M < cZUDpnJQn; {
		plHJUlaC4H.iaaUnOsr7_M *= 2
		plHJUlaC4H =  /*line :1*/ kXMQLm5mTkv(plHJUlaC4H)
	}

	 /*line :1*/ uW0eIVjy.Mul(pmEgen2, plHJUlaC4H)
}



func gUZEPIC8uIA(ynnmh4ocU0R uint32) *WH8dWN {
	uW0eIVjy :=  /*line :1*/ new(WH8dWN)

	uW0eIVjy.s7DDkIlY =  /*line :1*/ uW0eIVjy.s7DDkIlY.gS9bqHpKAJ0( /*line :1*/ int(ynnmh4ocU0R/_W) * 2)
	return uW0eIVjy
}
