//line :1
package fRAfQd_



func H4EBRkSgB(lNAwkbJQVBgU GKMwTGxQa0S)



type lnrkj_BKF5Fp struct {
	us1JkVa	uintptr
	l6EX79SmT	int64
	zoYG3BJUf	int64
	ili5IDhr	func(any, uintptr)	
	e_kIVn	any
	jgF8vwc3o	uintptr
	xa0jxKh	int64
	atqvaG	uint32
}





func m8wvRCD4e(lNAwkbJQVBgU GKMwTGxQa0S) int64 {
	if lNAwkbJQVBgU <= 0 {
		return  /*line :1*/ qQJTJ6m()
	}
	pv6eev :=  /*line :1*/ qQJTJ6m() +  /*line :1*/ int64(lNAwkbJQVBgU)
	if pv6eev < 0 {

		pv6eev = 1<<63 - 1
	}
	return pv6eev
}

func dOaW670z(*lnrkj_BKF5Fp)
func yazRrvLb(*lnrkj_BKF5Fp) bool
func lZXiPYeQ(*lnrkj_BKF5Fp, int64) bool
func l6Unsr(pv6eev *lnrkj_BKF5Fp, m8wvRCD4e, cT78Yd int64, k05jcfKvAvw6 func(any, uintptr), rrnnKxbO any, sXWEj8 uintptr)





type Bkz_9qnV struct {
	N0KrfNaHg	<-chan G4KDkI
	iwIHY_C9NsJs	lnrkj_BKF5Fp
}























func (pv6eev *Bkz_9qnV) Stop() bool {
	if pv6eev.iwIHY_C9NsJs.ili5IDhr == nil {
		 /*line :1*/ panic("time: Stop called on uninitialized Timer")
	}
	return  /*line :1*/ yazRrvLb(&pv6eev.iwIHY_C9NsJs)
}



func Dyu6iYU(lNAwkbJQVBgU GKMwTGxQa0S) *Bkz_9qnV {
	fnpIad3KB :=  /*line :1*/ make(chan G4KDkI, 1)
	pv6eev := &Bkz_9qnV{
		N0KrfNaHg:	fnpIad3KB,
		iwIHY_C9NsJs: lnrkj_BKF5Fp{
			l6EX79SmT:	 /*line :1*/ m8wvRCD4e(lNAwkbJQVBgU),
			ili5IDhr:	ataaWLjK,
			e_kIVn:	fnpIad3KB,
		},
	}
	 /*line :1*/ dOaW670z(&pv6eev.iwIHY_C9NsJs)
	return pv6eev
}



































func (pv6eev *Bkz_9qnV) Reset(lNAwkbJQVBgU GKMwTGxQa0S) bool {
	if pv6eev.iwIHY_C9NsJs.ili5IDhr == nil {
		 /*line :1*/ panic("time: Reset called on uninitialized Timer")
	}
	n2g7_eLcTc :=  /*line :1*/ m8wvRCD4e(lNAwkbJQVBgU)
	return  /*line :1*/ lZXiPYeQ(&pv6eev.iwIHY_C9NsJs, n2g7_eLcTc)
}


func ataaWLjK(fnpIad3KB any, sXWEj8 uintptr) {
	select {
	case fnpIad3KB.(chan G4KDkI) <-  /*line :1*/ Pc_35oTY():
	default:
	}
}







func XflQtb(lNAwkbJQVBgU GKMwTGxQa0S) <-chan G4KDkI {
	return  /*line :1*/ Dyu6iYU(lNAwkbJQVBgU).N0KrfNaHg
}




func RzWe24uxMiv(lNAwkbJQVBgU GKMwTGxQa0S, k05jcfKvAvw6 func()) *Bkz_9qnV {
	pv6eev := &Bkz_9qnV{
		iwIHY_C9NsJs: lnrkj_BKF5Fp{
			l6EX79SmT:	 /*line :1*/ m8wvRCD4e(lNAwkbJQVBgU),
			ili5IDhr:	r3c2UZjI3g,
			e_kIVn:	k05jcfKvAvw6,
		},
	}
	 /*line :1*/ dOaW670z(&pv6eev.iwIHY_C9NsJs)
	return pv6eev
}

func r3c2UZjI3g(rrnnKxbO any, sXWEj8 uintptr) {
	go  /*line :1*/ rrnnKxbO.(func())()
}
