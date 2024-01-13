//line :1
package fRAfQd_



type DvKj1iMo struct {
	N0KrfNaHg	<-chan G4KDkI	
	iwIHY_C9NsJs	lnrkj_BKF5Fp
}







func ZrJvk4Xf6iI(lNAwkbJQVBgU GKMwTGxQa0S) *DvKj1iMo {
	if lNAwkbJQVBgU <= 0 {
		 /*line :1*/ panic("non-positive interval for NewTicker")
	}

	fnpIad3KB :=  /*line :1*/ make(chan G4KDkI, 1)
	pv6eev := &DvKj1iMo{
		N0KrfNaHg:	fnpIad3KB,
		iwIHY_C9NsJs: lnrkj_BKF5Fp{
			l6EX79SmT:	 /*line :1*/ m8wvRCD4e(lNAwkbJQVBgU),
			zoYG3BJUf:	 /*line :1*/ int64(lNAwkbJQVBgU),
			ili5IDhr:	ataaWLjK,
			e_kIVn:	fnpIad3KB,
		},
	}
	 /*line :1*/ dOaW670z(&pv6eev.iwIHY_C9NsJs)
	return pv6eev
}




func (pv6eev *DvKj1iMo) Stop() {
	 /*line :1*/ yazRrvLb(&pv6eev.iwIHY_C9NsJs)
}




func (pv6eev *DvKj1iMo) Reset(lNAwkbJQVBgU GKMwTGxQa0S) {
	if lNAwkbJQVBgU <= 0 {
		 /*line :1*/ panic("non-positive interval for Ticker.Reset")
	}
	if pv6eev.iwIHY_C9NsJs.ili5IDhr == nil {
		 /*line :1*/ panic("time: Reset called on uninitialized Ticker")
	}
	 /*line :1*/ l6Unsr(&pv6eev.iwIHY_C9NsJs,  /*line :1*/ m8wvRCD4e(lNAwkbJQVBgU),  /*line :1*/ int64(lNAwkbJQVBgU), pv6eev.iwIHY_C9NsJs.ili5IDhr, pv6eev.iwIHY_C9NsJs.e_kIVn, pv6eev.iwIHY_C9NsJs.jgF8vwc3o)
}






func Fp8gzF(lNAwkbJQVBgU GKMwTGxQa0S) <-chan G4KDkI {
	if lNAwkbJQVBgU <= 0 {
		return nil
	}
	return  /*line :1*/ ZrJvk4Xf6iI(lNAwkbJQVBgU).N0KrfNaHg
}
