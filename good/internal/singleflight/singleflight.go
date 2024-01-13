//line :1


package EkvK9so

import sync "sync"


type y44M_9Vc struct {
	wemzaA12g	sync.NMlO0OA

	
	
	dCI_aTzGSKs	any
	rVDxhOmW	error

	
	
	
	biQ5cd57K	int
	kUk6iffrLP	[]chan<- RfcQmoNLzjV
}



type JRXHs_ struct {
	eyJCfKG	sync.DIRWe11KYlYa	
	cwFriu_9	map[string]*y44M_9Vc	
}



type RfcQmoNLzjV struct {
	YeKlYE	any
	ZH7an4	error
	YNGawz	bool
}






func (_B2KE20C *JRXHs_) Do(eaOpTegm string, zN7X9q func() (any, error)) (yc5DWa any, aPstLLNV error, cphTOi30Rrl bool) {
	 /*line :1*/ _B2KE20C.eyJCfKG.Lock()
	if _B2KE20C.cwFriu_9 == nil {
		_B2KE20C.cwFriu_9 =  /*line :1*/ make(map[string]*y44M_9Vc)
	}
	if nxivmmivYNa9, vzxwyH_Q := _B2KE20C.cwFriu_9[eaOpTegm]; vzxwyH_Q {
		nxivmmivYNa9.biQ5cd57K++
		 /*line :1*/ _B2KE20C.eyJCfKG.Unlock()
		 /*line :1*/ nxivmmivYNa9.wemzaA12g.Wait()
		return nxivmmivYNa9.dCI_aTzGSKs, nxivmmivYNa9.rVDxhOmW, true
	}
	nxivmmivYNa9 :=  /*line :1*/ new(y44M_9Vc)
	 /*line :1*/ nxivmmivYNa9.wemzaA12g.Add(1)
	_B2KE20C.cwFriu_9[eaOpTegm] = nxivmmivYNa9
	 /*line :1*/ _B2KE20C.eyJCfKG.Unlock()

	 /*line :1*/ _B2KE20C.s9F0Bxx68Eo(nxivmmivYNa9, eaOpTegm, zN7X9q)
	return nxivmmivYNa9.dCI_aTzGSKs, nxivmmivYNa9.rVDxhOmW, nxivmmivYNa9.biQ5cd57K > 0
}



func (_B2KE20C *JRXHs_) DoChan(eaOpTegm string, zN7X9q func() (any, error)) <-chan RfcQmoNLzjV {
	aNXdVjFL1_IV :=  /*line :1*/ make(chan RfcQmoNLzjV, 1)
	 /*line :1*/ _B2KE20C.eyJCfKG.Lock()
	if _B2KE20C.cwFriu_9 == nil {
		_B2KE20C.cwFriu_9 =  /*line :1*/ make(map[string]*y44M_9Vc)
	}
	if nxivmmivYNa9, vzxwyH_Q := _B2KE20C.cwFriu_9[eaOpTegm]; vzxwyH_Q {
		nxivmmivYNa9.biQ5cd57K++
		nxivmmivYNa9.kUk6iffrLP =  /*line :1*/ append(nxivmmivYNa9.kUk6iffrLP, aNXdVjFL1_IV)
		 /*line :1*/ _B2KE20C.eyJCfKG.Unlock()
		return aNXdVjFL1_IV
	}
	nxivmmivYNa9 := &y44M_9Vc{kUk6iffrLP: []chan<- RfcQmoNLzjV{aNXdVjFL1_IV}}
	 /*line :1*/ nxivmmivYNa9.wemzaA12g.Add(1)
	_B2KE20C.cwFriu_9[eaOpTegm] = nxivmmivYNa9
	 /*line :1*/ _B2KE20C.eyJCfKG.Unlock()

	go  /*line :1*/ _B2KE20C.s9F0Bxx68Eo(nxivmmivYNa9, eaOpTegm, zN7X9q)

	return aNXdVjFL1_IV
}


func (_B2KE20C *JRXHs_) s9F0Bxx68Eo(nxivmmivYNa9 *y44M_9Vc, eaOpTegm string, zN7X9q func() (any, error)) {
	nxivmmivYNa9.dCI_aTzGSKs, nxivmmivYNa9.rVDxhOmW =  /*line :1*/ zN7X9q()

	 /*line :1*/ _B2KE20C.eyJCfKG.Lock()
	 /*line :1*/ nxivmmivYNa9.wemzaA12g.Done()
	if _B2KE20C.cwFriu_9[eaOpTegm] == nxivmmivYNa9 {
		 /*line :1*/ delete(_B2KE20C.cwFriu_9, eaOpTegm)
	}
	for _, aNXdVjFL1_IV := range nxivmmivYNa9.kUk6iffrLP {
		aNXdVjFL1_IV <- RfcQmoNLzjV{nxivmmivYNa9.dCI_aTzGSKs, nxivmmivYNa9.rVDxhOmW, nxivmmivYNa9.biQ5cd57K > 0}
	}
	 /*line :1*/ _B2KE20C.eyJCfKG.Unlock()
}






func (_B2KE20C *JRXHs_) ForgetUnshared(eaOpTegm string) bool {
	 /*line :1*/ _B2KE20C.eyJCfKG.Lock()
	defer  /*line :1*/ _B2KE20C.eyJCfKG.Unlock()
	nxivmmivYNa9, vzxwyH_Q := _B2KE20C.cwFriu_9[eaOpTegm]
	if !vzxwyH_Q {
		return true
	}
	if nxivmmivYNa9.biQ5cd57K == 0 {
		 /*line :1*/ delete(_B2KE20C.cwFriu_9, eaOpTegm)
		return true
	}
	return false
}
