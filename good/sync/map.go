//line :1
package sync

import (
	atomic "sync/atomic"
)

























type Eim1b6bNi9IM struct {
	c7d1a3	DIRWe11KYlYa

	
	
	
	
	
	
	
	
	
	uP8K2Sx	atomic.ToSaNw[a2y5y6NS_]

	
	
	
	
	
	
	
	
	
	
	eYdTJPgrT	map[any]*w_quV9nrf

	
	
	
	
	
	
	tR6O4DZ0Bs	int
}


type a2y5y6NS_ struct {
	hvlU9sP_F7ao	map[any]*w_quV9nrf
	cHXra7Ha4bn5	bool	
}



var xjlp9N =  /*line :1*/ new(any)


type w_quV9nrf struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	d1oekDvz atomic.ToSaNw[any]
}

func gA12Wv0(cmSN7RqmPwHo any) *w_quV9nrf {
	gPfqDBS := &w_quV9nrf{}
	 /*line :1*/ gPfqDBS.d1oekDvz.Store(&cmSN7RqmPwHo)
	return gPfqDBS
}

func (dMWQaSB *Eim1b6bNi9IM) mn5WoH9n2l() a2y5y6NS_ {
	if lcXhtO :=  /*line :1*/ dMWQaSB.uP8K2Sx.Load(); lcXhtO != nil {
		return *lcXhtO
	}
	return a2y5y6NS_{}
}




func (dMWQaSB *Eim1b6bNi9IM) Load(dcQtxL any) (fr4ZEM any, qxRgxXc1hZi bool) {
	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]
	if !qxRgxXc1hZi && oN5icdPwACr.cHXra7Ha4bn5 {
		 /*line :1*/ dMWQaSB.c7d1a3.Lock()

		oN5icdPwACr =  /*line :1*/ dMWQaSB.mn5WoH9n2l()
		gPfqDBS, qxRgxXc1hZi = oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]
		if !qxRgxXc1hZi && oN5icdPwACr.cHXra7Ha4bn5 {
			gPfqDBS, qxRgxXc1hZi = dMWQaSB.eYdTJPgrT[dcQtxL]

			 /*line :1*/ dMWQaSB.abeAXJ4mcx()
		}
		 /*line :1*/ dMWQaSB.c7d1a3.Unlock()
	}
	if !qxRgxXc1hZi {
		return nil, false
	}
	return  /*line :1*/ gPfqDBS.hup3VF7r1dU6()
}

func (gPfqDBS *w_quV9nrf) hup3VF7r1dU6() (fr4ZEM any, qxRgxXc1hZi bool) {
	lcXhtO :=  /*line :1*/ gPfqDBS.d1oekDvz.Load()
	if lcXhtO == nil || lcXhtO == xjlp9N {
		return nil, false
	}
	return *lcXhtO, true
}


func (dMWQaSB *Eim1b6bNi9IM) Store(dcQtxL, fr4ZEM any) {
	_, _ =  /*line :1*/ dMWQaSB.Swap(dcQtxL, fr4ZEM)
}







func (gPfqDBS *w_quV9nrf) w3bmyAivM(fRQ4D1, wlm_tUADFO30 any) bool {
	lcXhtO :=  /*line :1*/ gPfqDBS.d1oekDvz.Load()
	if lcXhtO == nil || lcXhtO == xjlp9N || *lcXhtO != fRQ4D1 {
		return false
	}

	m9ci74gwBFW := wlm_tUADFO30
	for {
		if  /*line :1*/ gPfqDBS.d1oekDvz.CompareAndSwap(lcXhtO, &m9ci74gwBFW) {
			return true
		}
		lcXhtO =  /*line :1*/ gPfqDBS.d1oekDvz.Load()
		if lcXhtO == nil || lcXhtO == xjlp9N || *lcXhtO != fRQ4D1 {
			return false
		}
	}
}





func (gPfqDBS *w_quV9nrf) k3DkWeZ5_gm() (c_crlOFWutTn bool) {
	return  /*line :1*/ gPfqDBS.d1oekDvz.CompareAndSwap(xjlp9N, nil)
}




func (gPfqDBS *w_quV9nrf) zLv2fI3R(cmSN7RqmPwHo *any) *any {
	return  /*line :1*/ gPfqDBS.d1oekDvz.Swap(cmSN7RqmPwHo)
}




func (dMWQaSB *Eim1b6bNi9IM) LoadOrStore(dcQtxL, fr4ZEM any) (aa9ShmSK8oj any, v4Hl9Z bool) {

	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	if gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]; qxRgxXc1hZi {
		aa9ShmSK8oj, v4Hl9Z, qxRgxXc1hZi :=  /*line :1*/ gPfqDBS.asB80XI8dry(fr4ZEM)
		if qxRgxXc1hZi {
			return aa9ShmSK8oj, v4Hl9Z
		}
	}

	 /*line :1*/ dMWQaSB.c7d1a3.Lock()
	oN5icdPwACr =  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	if gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]; qxRgxXc1hZi {
		if  /*line :1*/ gPfqDBS.k3DkWeZ5_gm() {
			dMWQaSB.eYdTJPgrT[dcQtxL] = gPfqDBS
		}
		aa9ShmSK8oj, v4Hl9Z, _ =  /*line :1*/ gPfqDBS.asB80XI8dry(fr4ZEM)
	} else if gPfqDBS, qxRgxXc1hZi := dMWQaSB.eYdTJPgrT[dcQtxL]; qxRgxXc1hZi {
		aa9ShmSK8oj, v4Hl9Z, _ =  /*line :1*/ gPfqDBS.asB80XI8dry(fr4ZEM)
		 /*line :1*/ dMWQaSB.abeAXJ4mcx()
	} else {
		if !oN5icdPwACr.cHXra7Ha4bn5 {

			 /*line :1*/ dMWQaSB.dXVk_kXT()
			 /*line :1*/ dMWQaSB.uP8K2Sx.Store(&a2y5y6NS_{hvlU9sP_F7ao: oN5icdPwACr.hvlU9sP_F7ao, cHXra7Ha4bn5: true})
		}
		dMWQaSB.eYdTJPgrT[dcQtxL] =  /*line :1*/ gA12Wv0(fr4ZEM)
		aa9ShmSK8oj, v4Hl9Z = fr4ZEM, false
	}
	 /*line :1*/ dMWQaSB.c7d1a3.Unlock()

	return aa9ShmSK8oj, v4Hl9Z
}






func (gPfqDBS *w_quV9nrf) asB80XI8dry(cmSN7RqmPwHo any) (aa9ShmSK8oj any, v4Hl9Z, qxRgxXc1hZi bool) {
	lcXhtO :=  /*line :1*/ gPfqDBS.d1oekDvz.Load()
	if lcXhtO == xjlp9N {
		return nil, false, false
	}
	if lcXhtO != nil {
		return *lcXhtO, true, true
	}

	bxAOa_AbrlA := cmSN7RqmPwHo
	for {
		if  /*line :1*/ gPfqDBS.d1oekDvz.CompareAndSwap(nil, &bxAOa_AbrlA) {
			return cmSN7RqmPwHo, false, true
		}
		lcXhtO =  /*line :1*/ gPfqDBS.d1oekDvz.Load()
		if lcXhtO == xjlp9N {
			return nil, false, false
		}
		if lcXhtO != nil {
			return *lcXhtO, true, true
		}
	}
}



func (dMWQaSB *Eim1b6bNi9IM) LoadAndDelete(dcQtxL any) (fr4ZEM any, v4Hl9Z bool) {
	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]
	if !qxRgxXc1hZi && oN5icdPwACr.cHXra7Ha4bn5 {
		 /*line :1*/ dMWQaSB.c7d1a3.Lock()
		oN5icdPwACr =  /*line :1*/ dMWQaSB.mn5WoH9n2l()
		gPfqDBS, qxRgxXc1hZi = oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]
		if !qxRgxXc1hZi && oN5icdPwACr.cHXra7Ha4bn5 {
			gPfqDBS, qxRgxXc1hZi = dMWQaSB.eYdTJPgrT[dcQtxL]
			 /*line :1*/ delete(dMWQaSB.eYdTJPgrT, dcQtxL)

			 /*line :1*/ dMWQaSB.abeAXJ4mcx()
		}
		 /*line :1*/ dMWQaSB.c7d1a3.Unlock()
	}
	if qxRgxXc1hZi {
		return  /*line :1*/ gPfqDBS.yD6yYwahQx()
	}
	return nil, false
}


func (dMWQaSB *Eim1b6bNi9IM) Delete(dcQtxL any) {
	 /*line :1*/ dMWQaSB.LoadAndDelete(dcQtxL)
}

func (gPfqDBS *w_quV9nrf) yD6yYwahQx() (fr4ZEM any, qxRgxXc1hZi bool) {
	for {
		lcXhtO :=  /*line :1*/ gPfqDBS.d1oekDvz.Load()
		if lcXhtO == nil || lcXhtO == xjlp9N {
			return nil, false
		}
		if  /*line :1*/ gPfqDBS.d1oekDvz.CompareAndSwap(lcXhtO, nil) {
			return *lcXhtO, true
		}
	}
}





func (gPfqDBS *w_quV9nrf) sbJX5UnLDm(cmSN7RqmPwHo *any) (*any, bool) {
	for {
		lcXhtO :=  /*line :1*/ gPfqDBS.d1oekDvz.Load()
		if lcXhtO == xjlp9N {
			return nil, false
		}
		if  /*line :1*/ gPfqDBS.d1oekDvz.CompareAndSwap(lcXhtO, cmSN7RqmPwHo) {
			return lcXhtO, true
		}
	}
}



func (dMWQaSB *Eim1b6bNi9IM) Swap(dcQtxL, fr4ZEM any) (dMa1yMqwbXax any, v4Hl9Z bool) {
	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	if gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]; qxRgxXc1hZi {
		if nBR5jJPd6nc3, qxRgxXc1hZi :=  /*line :1*/ gPfqDBS.sbJX5UnLDm(&fr4ZEM); qxRgxXc1hZi {
			if nBR5jJPd6nc3 == nil {
				return nil, false
			}
			return *nBR5jJPd6nc3, true
		}
	}

	 /*line :1*/ dMWQaSB.c7d1a3.Lock()
	oN5icdPwACr =  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	if gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]; qxRgxXc1hZi {
		if  /*line :1*/ gPfqDBS.k3DkWeZ5_gm() {

			dMWQaSB.eYdTJPgrT[dcQtxL] = gPfqDBS
		}
		if nBR5jJPd6nc3 :=  /*line :1*/ gPfqDBS.zLv2fI3R(&fr4ZEM); nBR5jJPd6nc3 != nil {
			v4Hl9Z = true
			dMa1yMqwbXax = *nBR5jJPd6nc3
		}
	} else if gPfqDBS, qxRgxXc1hZi := dMWQaSB.eYdTJPgrT[dcQtxL]; qxRgxXc1hZi {
		if nBR5jJPd6nc3 :=  /*line :1*/ gPfqDBS.zLv2fI3R(&fr4ZEM); nBR5jJPd6nc3 != nil {
			v4Hl9Z = true
			dMa1yMqwbXax = *nBR5jJPd6nc3
		}
	} else {
		if !oN5icdPwACr.cHXra7Ha4bn5 {

			 /*line :1*/ dMWQaSB.dXVk_kXT()
			 /*line :1*/ dMWQaSB.uP8K2Sx.Store(&a2y5y6NS_{hvlU9sP_F7ao: oN5icdPwACr.hvlU9sP_F7ao, cHXra7Ha4bn5: true})
		}
		dMWQaSB.eYdTJPgrT[dcQtxL] =  /*line :1*/ gA12Wv0(fr4ZEM)
	}
	 /*line :1*/ dMWQaSB.c7d1a3.Unlock()
	return dMa1yMqwbXax, v4Hl9Z
}




func (dMWQaSB *Eim1b6bNi9IM) CompareAndSwap(dcQtxL, fRQ4D1, wlm_tUADFO30 any) bool {
	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	if gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]; qxRgxXc1hZi {
		return  /*line :1*/ gPfqDBS.w3bmyAivM(fRQ4D1, wlm_tUADFO30)
	} else if !oN5icdPwACr.cHXra7Ha4bn5 {
		return false
	}

	 /*line :1*/ dMWQaSB.c7d1a3.Lock()
	defer  /*line :1*/ dMWQaSB.c7d1a3.Unlock()
	oN5icdPwACr =  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	h0W1qm := false
	if gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]; qxRgxXc1hZi {
		h0W1qm =  /*line :1*/ gPfqDBS.w3bmyAivM(fRQ4D1, wlm_tUADFO30)
	} else if gPfqDBS, qxRgxXc1hZi := dMWQaSB.eYdTJPgrT[dcQtxL]; qxRgxXc1hZi {
		h0W1qm =  /*line :1*/ gPfqDBS.w3bmyAivM(fRQ4D1, wlm_tUADFO30)

		 /*line :1*/ dMWQaSB.abeAXJ4mcx()
	}
	return h0W1qm
}






func (dMWQaSB *Eim1b6bNi9IM) CompareAndDelete(dcQtxL, fRQ4D1 any) (z2G7ihrQ bool) {
	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	gPfqDBS, qxRgxXc1hZi := oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]
	if !qxRgxXc1hZi && oN5icdPwACr.cHXra7Ha4bn5 {
		 /*line :1*/ dMWQaSB.c7d1a3.Lock()
		oN5icdPwACr =  /*line :1*/ dMWQaSB.mn5WoH9n2l()
		gPfqDBS, qxRgxXc1hZi = oN5icdPwACr.hvlU9sP_F7ao[dcQtxL]
		if !qxRgxXc1hZi && oN5icdPwACr.cHXra7Ha4bn5 {
			gPfqDBS, qxRgxXc1hZi = dMWQaSB.eYdTJPgrT[dcQtxL]

			 /*line :1*/ dMWQaSB.abeAXJ4mcx()
		}
		 /*line :1*/ dMWQaSB.c7d1a3.Unlock()
	}
	for qxRgxXc1hZi {
		lcXhtO :=  /*line :1*/ gPfqDBS.d1oekDvz.Load()
		if lcXhtO == nil || lcXhtO == xjlp9N || *lcXhtO != fRQ4D1 {
			return false
		}
		if  /*line :1*/ gPfqDBS.d1oekDvz.CompareAndSwap(lcXhtO, nil) {
			return true
		}
	}
	return false
}












func (dMWQaSB *Eim1b6bNi9IM) Range(pUYwrq func(dcQtxL, fr4ZEM any) bool) {

	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	if oN5icdPwACr.cHXra7Ha4bn5 {

		 /*line :1*/ dMWQaSB.c7d1a3.Lock()
		oN5icdPwACr =  /*line :1*/ dMWQaSB.mn5WoH9n2l()
		if oN5icdPwACr.cHXra7Ha4bn5 {
			oN5icdPwACr = a2y5y6NS_{hvlU9sP_F7ao: dMWQaSB.eYdTJPgrT}
			 /*line :1*/ dMWQaSB.uP8K2Sx.Store(&oN5icdPwACr)
			dMWQaSB.eYdTJPgrT = nil
			dMWQaSB.tR6O4DZ0Bs = 0
		}
		 /*line :1*/ dMWQaSB.c7d1a3.Unlock()
	}

	for vSKa8czf, gPfqDBS := range oN5icdPwACr.hvlU9sP_F7ao {
		nBR5jJPd6nc3, qxRgxXc1hZi :=  /*line :1*/ gPfqDBS.hup3VF7r1dU6()
		if !qxRgxXc1hZi {
			continue
		}
		if ! /*line :1*/ pUYwrq(vSKa8czf, nBR5jJPd6nc3) {
			break
		}
	}
}

func (dMWQaSB *Eim1b6bNi9IM) abeAXJ4mcx() {
	dMWQaSB.tR6O4DZ0Bs++
	if dMWQaSB.tR6O4DZ0Bs <  /*line :1*/ len(dMWQaSB.eYdTJPgrT) {
		return
	}
	 /*line :1*/ dMWQaSB.uP8K2Sx.Store(&a2y5y6NS_{hvlU9sP_F7ao: dMWQaSB.eYdTJPgrT})
	dMWQaSB.eYdTJPgrT = nil
	dMWQaSB.tR6O4DZ0Bs = 0
}

func (dMWQaSB *Eim1b6bNi9IM) dXVk_kXT() {
	if dMWQaSB.eYdTJPgrT != nil {
		return
	}

	oN5icdPwACr :=  /*line :1*/ dMWQaSB.mn5WoH9n2l()
	dMWQaSB.eYdTJPgrT =  /*line :1*/ make(map[any]*w_quV9nrf,  /*line :1*/ len(oN5icdPwACr.hvlU9sP_F7ao))
	for vSKa8czf, gPfqDBS := range oN5icdPwACr.hvlU9sP_F7ao {
		if ! /*line :1*/ gPfqDBS.zLgjbg3() {
			dMWQaSB.eYdTJPgrT[vSKa8czf] = gPfqDBS
		}
	}
}

func (gPfqDBS *w_quV9nrf) zLgjbg3() (cvUvCY bool) {
	lcXhtO :=  /*line :1*/ gPfqDBS.d1oekDvz.Load()
	for lcXhtO == nil {
		if  /*line :1*/ gPfqDBS.d1oekDvz.CompareAndSwap(nil, xjlp9N) {
			return true
		}
		lcXhtO =  /*line :1*/ gPfqDBS.d1oekDvz.Load()
	}
	return lcXhtO == xjlp9N
}
