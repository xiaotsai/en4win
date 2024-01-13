//line :1
package os208GicmRnc

import math "math"


type Z19UgGXyw struct {
	cc8bQJ	*YAxSls
	lDK27B4PRFS	float64
	t4eTbaVU56K	float64
	jZz_LpeWC5l0	float64
	dQYj8fAt	float64
	dSPY5TK	float64
	aIS_4aqaZ	float64
	r9AgKfP	float64
	abAp5hXrVo	float64
}

func (ruBcOu *Z19UgGXyw) fHqMQhln(u6q_I9fv3S float64) float64 {
	return  /*line :1*/ math.JPvpMNJa(ruBcOu.dSPY5TK* /*line :1*/ math.JrJVPb5WbG(ruBcOu.t4eTbaVU56K+u6q_I9fv3S)) * ruBcOu.aIS_4aqaZ
}

func (ruBcOu *Z19UgGXyw) ayaX8Xv(u6q_I9fv3S float64) float64 {
	return  /*line :1*/ math.JPvpMNJa(ruBcOu.aIS_4aqaZ* /*line :1*/ math.JrJVPb5WbG(ruBcOu.dSPY5TK*u6q_I9fv3S)) - ruBcOu.t4eTbaVU56K
}





func Mhew1ZK8Xr(csIue666O0e *YAxSls, s_FevaW float64, zqelgkLITo9 float64, oFhza8IR9B uint64) *Z19UgGXyw {
	ruBcOu :=  /*line :1*/ new(Z19UgGXyw)
	if s_FevaW <= 1.0 || zqelgkLITo9 < 1 {
		return nil
	}
	ruBcOu.cc8bQJ = csIue666O0e
	ruBcOu.lDK27B4PRFS =  /*line :1*/ float64(oFhza8IR9B)
	ruBcOu.t4eTbaVU56K = zqelgkLITo9
	ruBcOu.jZz_LpeWC5l0 = s_FevaW
	ruBcOu.dSPY5TK = 1.0 - ruBcOu.jZz_LpeWC5l0
	ruBcOu.aIS_4aqaZ = 1.0 / ruBcOu.dSPY5TK
	ruBcOu.r9AgKfP =  /*line :1*/ ruBcOu.fHqMQhln(ruBcOu.lDK27B4PRFS + 0.5)
	ruBcOu.abAp5hXrVo =  /*line :1*/ ruBcOu.fHqMQhln(0.5) -  /*line :1*/ math.JPvpMNJa( /*line :1*/ math.JrJVPb5WbG(ruBcOu.t4eTbaVU56K)*(-ruBcOu.jZz_LpeWC5l0)) - ruBcOu.r9AgKfP
	ruBcOu.dQYj8fAt = 1 -  /*line :1*/ ruBcOu.ayaX8Xv( /*line :1*/ ruBcOu.fHqMQhln(1.5)- /*line :1*/ math.JPvpMNJa(-ruBcOu.jZz_LpeWC5l0* /*line :1*/ math.JrJVPb5WbG(ruBcOu.t4eTbaVU56K+1.0)))
	return ruBcOu
}



func (ruBcOu *Z19UgGXyw) Uint64() uint64 {
	if ruBcOu == nil {
		 /*line :1*/ panic("rand: nil Zipf")
	}
	iDt5dUcDv := 0.0

	for {
		csIue666O0e :=  /*line :1*/ ruBcOu.cc8bQJ.Float64()
		jLeZd0YX := ruBcOu.r9AgKfP + csIue666O0e*ruBcOu.abAp5hXrVo
		u6q_I9fv3S :=  /*line :1*/ ruBcOu.ayaX8Xv(jLeZd0YX)
		iDt5dUcDv =  /*line :1*/ math.Floor(u6q_I9fv3S + 0.5)
		if iDt5dUcDv-u6q_I9fv3S <= ruBcOu.dQYj8fAt {
			break
		}
		if jLeZd0YX >=  /*line :1*/ ruBcOu.fHqMQhln(iDt5dUcDv+0.5)- /*line :1*/ math.JPvpMNJa(- /*line :1*/ math.JrJVPb5WbG(iDt5dUcDv+ruBcOu.t4eTbaVU56K)*ruBcOu.jZz_LpeWC5l0) {
			break
		}
	}
	return  /*line :1*/ uint64(iDt5dUcDv)
}
