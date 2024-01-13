//line :1
package sync

import (
	race "V0p_iv"
	atomic "sync/atomic"
	"unsafe"
)

type V1sRgjwAglJt struct {
	jxiFzuCg	DIRWe11KYlYa
	iWBICrwBVf_	uint32
	gHg271	uint32
	sz1C0Sh0ibW	atomic.JhtCwKEzC
	eJoJ4xHYR	atomic.JhtCwKEzC
}

const rwmutexMaxReaders = 1 << 30

func (owhYZaeupcM_ *V1sRgjwAglJt) RLock() {
	if race.Enabled {
		_ = owhYZaeupcM_.jxiFzuCg.i58pIQ
		 /*line :1*/ race.F_wVXvO1u_kM()
	}
	if  /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.Add(1) < 0 {

		 /*line :1*/ cNUH1DZ_2v(&owhYZaeupcM_.gHg271, false, 0)
	}
	if race.Enabled {
		 /*line :1*/ race.UStrOKX()
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.gHg271))
	}
}

func (owhYZaeupcM_ *V1sRgjwAglJt) TryRLock() bool {
	if race.Enabled {
		_ = owhYZaeupcM_.jxiFzuCg.i58pIQ
		 /*line :1*/ race.F_wVXvO1u_kM()
	}
	for {
		ew6gz7n0m :=  /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.Load()
		if ew6gz7n0m < 0 {
			if race.Enabled {
				 /*line :1*/ race.UStrOKX()
			}
			return false
		}
		if  /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.CompareAndSwap(ew6gz7n0m, ew6gz7n0m+1) {
			if race.Enabled {
				 /*line :1*/ race.UStrOKX()
				 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.gHg271))
			}
			return true
		}
	}
}

func (owhYZaeupcM_ *V1sRgjwAglJt) RUnlock() {
	if race.Enabled {
		_ = owhYZaeupcM_.jxiFzuCg.i58pIQ
		 /*line :1*/ race.AQrXiLDh3RTu( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.iWBICrwBVf_))
		 /*line :1*/ race.F_wVXvO1u_kM()
	}
	if xy07lY :=  /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.Add(-1); xy07lY < 0 {

		 /*line :1*/ owhYZaeupcM_.tYRm8C(xy07lY)
	}
	if race.Enabled {
		 /*line :1*/ race.UStrOKX()
	}
}

func (owhYZaeupcM_ *V1sRgjwAglJt) tYRm8C(xy07lY int32) {
	if xy07lY+1 == 0 || xy07lY+1 == -rwmutexMaxReaders {
		 /*line :1*/ race.UStrOKX()
		 /*line :1*/ h7pm4f2ZL("sync: RUnlock of unlocked RWMutex")
	}

	if  /*line :1*/ owhYZaeupcM_.eJoJ4xHYR.Add(-1) == 0 {

		 /*line :1*/ xAzRYQiZaT0(&owhYZaeupcM_.iWBICrwBVf_, false, 1)
	}
}

func (owhYZaeupcM_ *V1sRgjwAglJt) Lock() {
	if race.Enabled {
		_ = owhYZaeupcM_.jxiFzuCg.i58pIQ
		 /*line :1*/ race.F_wVXvO1u_kM()
	}

	 /*line :1*/ owhYZaeupcM_.jxiFzuCg.Lock()

	xy07lY :=  /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.Add(-rwmutexMaxReaders) + rwmutexMaxReaders

	if xy07lY != 0 &&  /*line :1*/ owhYZaeupcM_.eJoJ4xHYR.Add(xy07lY) != 0 {
		 /*line :1*/ xmhHf_d0jEJ(&owhYZaeupcM_.iWBICrwBVf_, false, 0)
	}
	if race.Enabled {
		 /*line :1*/ race.UStrOKX()
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.gHg271))
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.iWBICrwBVf_))
	}
}

func (owhYZaeupcM_ *V1sRgjwAglJt) TryLock() bool {
	if race.Enabled {
		_ = owhYZaeupcM_.jxiFzuCg.i58pIQ
		 /*line :1*/ race.F_wVXvO1u_kM()
	}
	if ! /*line :1*/ owhYZaeupcM_.jxiFzuCg.TryLock() {
		if race.Enabled {
			 /*line :1*/ race.UStrOKX()
		}
		return false
	}
	if ! /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.CompareAndSwap(0, -rwmutexMaxReaders) {
		 /*line :1*/ owhYZaeupcM_.jxiFzuCg.Unlock()
		if race.Enabled {
			 /*line :1*/ race.UStrOKX()
		}
		return false
	}
	if race.Enabled {
		 /*line :1*/ race.UStrOKX()
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.gHg271))
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.iWBICrwBVf_))
	}
	return true
}

func (owhYZaeupcM_ *V1sRgjwAglJt) Unlock() {
	if race.Enabled {
		_ = owhYZaeupcM_.jxiFzuCg.i58pIQ
		 /*line :1*/ race.VYdLYvY( /*line :1*/ unsafe.Pointer(&owhYZaeupcM_.gHg271))
		 /*line :1*/ race.F_wVXvO1u_kM()
	}

	xy07lY :=  /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.Add(rwmutexMaxReaders)
	if xy07lY >= rwmutexMaxReaders {
		 /*line :1*/ race.UStrOKX()
		 /*line :1*/ h7pm4f2ZL("sync: Unlock of unlocked RWMutex")
	}

	for cmSN7RqmPwHo := 0; cmSN7RqmPwHo <  /*line :1*/ int(xy07lY); cmSN7RqmPwHo++ {
		 /*line :1*/ xAzRYQiZaT0(&owhYZaeupcM_.gHg271, false, 0)
	}

	 /*line :1*/ owhYZaeupcM_.jxiFzuCg.Unlock()
	if race.Enabled {
		 /*line :1*/ race.UStrOKX()
	}
}

//go:linkname ug2GzYF bUKeamF.eoa4SHku0
func ug2GzYF(owhYZaeupcM_ *V1sRgjwAglJt) bool {
	xy07lY :=  /*line :1*/ owhYZaeupcM_.sz1C0Sh0ibW.Load()
	return xy07lY < 0 && xy07lY+rwmutexMaxReaders > 0
}

func (owhYZaeupcM_ *V1sRgjwAglJt) RLocker() OXw5jPmA {
	return (* /*line :1*/ _vhtmea3a70)(owhYZaeupcM_)
}

type _vhtmea3a70 V1sRgjwAglJt

func (xy07lY *_vhtmea3a70) Lock()	{ (* /*line :1*/ V1sRgjwAglJt)(xy07lY).RLock() }
func (xy07lY *_vhtmea3a70) Unlock()	{ (* /*line :1*/ V1sRgjwAglJt)(xy07lY).RUnlock() }
