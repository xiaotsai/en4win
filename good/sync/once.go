//line :1
package sync

import (
	atomic "sync/atomic"
)








type LhBfwn6wa1x struct {
	
	
	
	
	
	wPaWAF4RoiGK	uint32
	uR502Yi9haN	DIRWe11KYlYa
}





















func (iutN0D *LhBfwn6wa1x) Do(pUYwrq func()) {

	if  /*line :1*/ atomic.LoadUint32(&iutN0D.wPaWAF4RoiGK) == 0 {

		 /*line :1*/ iutN0D.q0v25n8yf(pUYwrq)
	}
}

func (iutN0D *LhBfwn6wa1x) q0v25n8yf(pUYwrq func()) {
	 /*line :1*/ iutN0D.uR502Yi9haN.Lock()
	defer  /*line :1*/ iutN0D.uR502Yi9haN.Unlock()
	if iutN0D.wPaWAF4RoiGK == 0 {
		defer  /*line :1*/ atomic.StoreUint32(&iutN0D.wPaWAF4RoiGK, 1)
		 /*line :1*/ pUYwrq()
	}
}
