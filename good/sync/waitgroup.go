//line :1
package sync

import (
	race "V0p_iv"
	atomic "sync/atomic"
	"unsafe"
)











type NMlO0OA struct {
	gY4vTHKnJHS	tzBuysuv

	axpDliphDdUW	atomic.ZskaiQXr1j	
	tOUps8CMyuS	uint32
}














func (a9ro9qSM *NMlO0OA) Add(yza2GiF int) {
	if race.Enabled {
		if yza2GiF < 0 {

			 /*line :1*/ race.AQrXiLDh3RTu( /*line :1*/ unsafe.Pointer(a9ro9qSM))
		}
		 /*line :1*/ race.F_wVXvO1u_kM()
		defer  /*line :1*/ race.UStrOKX()
	}
	c5eCUapxazHh :=  /*line :1*/ a9ro9qSM.axpDliphDdUW.Add( /*line :1*/ uint64(yza2GiF) << 32)
	nBR5jJPd6nc3 :=  /*line :1*/ int32(c5eCUapxazHh >> 32)
	hG2g0Z :=  /*line :1*/ uint32(c5eCUapxazHh)
	if race.Enabled && yza2GiF > 0 && nBR5jJPd6nc3 ==  /*line :1*/ int32(yza2GiF) {

		 /*line :1*/ race.RyhN0dW( /*line :1*/ unsafe.Pointer(&a9ro9qSM.tOUps8CMyuS))
	}
	if nBR5jJPd6nc3 < 0 {
		 /*line :1*/ panic("sync: negative WaitGroup counter")
	}
	if hG2g0Z != 0 && yza2GiF > 0 && nBR5jJPd6nc3 ==  /*line :1*/ int32(yza2GiF) {
		 /*line :1*/ panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	if nBR5jJPd6nc3 > 0 || hG2g0Z == 0 {
		return
	}

	if  /*line :1*/ a9ro9qSM.axpDliphDdUW.Load() != c5eCUapxazHh {
		 /*line :1*/ panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}

	 /*line :1*/ a9ro9qSM.axpDliphDdUW.Store(0)
	for ; hG2g0Z != 0; hG2g0Z-- {
		 /*line :1*/ xAzRYQiZaT0(&a9ro9qSM.tOUps8CMyuS, false, 0)
	}
}


func (a9ro9qSM *NMlO0OA) Done() {
	 /*line :1*/ a9ro9qSM.Add(-1)
}


func (a9ro9qSM *NMlO0OA) Wait() {
	if race.Enabled {
		 /*line :1*/ race.F_wVXvO1u_kM()
	}
	for {
		c5eCUapxazHh :=  /*line :1*/ a9ro9qSM.axpDliphDdUW.Load()
		nBR5jJPd6nc3 :=  /*line :1*/ int32(c5eCUapxazHh >> 32)
		hG2g0Z :=  /*line :1*/ uint32(c5eCUapxazHh)
		if nBR5jJPd6nc3 == 0 {

			if race.Enabled {
				 /*line :1*/ race.UStrOKX()
				 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(a9ro9qSM))
			}
			return
		}

		if  /*line :1*/ a9ro9qSM.axpDliphDdUW.CompareAndSwap(c5eCUapxazHh, c5eCUapxazHh+1) {
			if race.Enabled && hG2g0Z == 0 {

				 /*line :1*/ race.FxGYURh( /*line :1*/ unsafe.Pointer(&a9ro9qSM.tOUps8CMyuS))
			}
			 /*line :1*/ nHODo9(&a9ro9qSM.tOUps8CMyuS)
			if  /*line :1*/ a9ro9qSM.axpDliphDdUW.Load() != 0 {
				 /*line :1*/ panic("sync: WaitGroup is reused before previous Wait has returned")
			}
			if race.Enabled {
				 /*line :1*/ race.UStrOKX()
				 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(a9ro9qSM))
			}
			return
		}
	}
}
