//line :1






package sync

import (
	race "V0p_iv"
	atomic "sync/atomic"
	"unsafe"
)


func f6Asug(string)
func h7pm4f2ZL(string)












type DIRWe11KYlYa struct {
	i58pIQ	int32
	r6mFI6Pfy9q	uint32
}


type OXw5jPmA interface {
	Lock()
	Unlock()
}

const (
	mutexLocked	= 1 << iota		
	mutexWoken
	mutexStarving
	mutexWaiterShift	= iota

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	starvationThresholdNs	= 1e6
)




func (dMWQaSB *DIRWe11KYlYa) Lock() {

	if  /*line :1*/ atomic.CompareAndSwapInt32(&dMWQaSB.i58pIQ, 0, mutexLocked) {
		if race.Enabled {
			 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(dMWQaSB))
		}
		return
	}

	 /*line :1*/ dMWQaSB.gjzcQIz9Psdj()
}






func (dMWQaSB *DIRWe11KYlYa) TryLock() bool {
	fRQ4D1 := dMWQaSB.i58pIQ
	if fRQ4D1&(mutexLocked|mutexStarving) != 0 {
		return false
	}

	if ! /*line :1*/ atomic.CompareAndSwapInt32(&dMWQaSB.i58pIQ, fRQ4D1, fRQ4D1|mutexLocked) {
		return false
	}

	if race.Enabled {
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(dMWQaSB))
	}
	return true
}

func (dMWQaSB *DIRWe11KYlYa) gjzcQIz9Psdj() {
	var rhOw_X8jb int64
	q0LX6Qb := false
	j5kUAanRO4d := false
	pLFWAdQ := 0
	fRQ4D1 := dMWQaSB.i58pIQ
	for {

		if fRQ4D1&(mutexLocked|mutexStarving) == mutexLocked &&  /*line :1*/ azKndVE(pLFWAdQ) {

			if !j5kUAanRO4d && fRQ4D1&mutexWoken == 0 && fRQ4D1>>mutexWaiterShift != 0 &&
				 /*line :1*/ atomic.CompareAndSwapInt32(&dMWQaSB.i58pIQ, fRQ4D1, fRQ4D1|mutexWoken) {
				j5kUAanRO4d = true
			}
			 /*line :1*/ qamXjagBlRdS()
			pLFWAdQ++
			fRQ4D1 = dMWQaSB.i58pIQ
			continue
		}
		wlm_tUADFO30 := fRQ4D1

		if fRQ4D1&mutexStarving == 0 {
			wlm_tUADFO30 |= mutexLocked
		}
		if fRQ4D1&(mutexLocked|mutexStarving) != 0 {
			wlm_tUADFO30 += 1 << mutexWaiterShift
		}

		if q0LX6Qb && fRQ4D1&mutexLocked != 0 {
			wlm_tUADFO30 |= mutexStarving
		}
		if j5kUAanRO4d {

			if wlm_tUADFO30&mutexWoken == 0 {
				 /*line :1*/ f6Asug("sync: inconsistent mutex state")
			}
			wlm_tUADFO30 &^= mutexWoken
		}
		if  /*line :1*/ atomic.CompareAndSwapInt32(&dMWQaSB.i58pIQ, fRQ4D1, wlm_tUADFO30) {
			if fRQ4D1&(mutexLocked|mutexStarving) == 0 {
				break
			}

			u67bOS8 := rhOw_X8jb != 0
			if rhOw_X8jb == 0 {
				rhOw_X8jb =  /*line :1*/ s6iizcJK2iq()
			}
			 /*line :1*/ vL6PE7(&dMWQaSB.r6mFI6Pfy9q, u67bOS8, 1)
			q0LX6Qb = q0LX6Qb ||  /*line :1*/ s6iizcJK2iq()-rhOw_X8jb > starvationThresholdNs
			fRQ4D1 = dMWQaSB.i58pIQ
			if fRQ4D1&mutexStarving != 0 {

				if fRQ4D1&(mutexLocked|mutexWoken) != 0 || fRQ4D1>>mutexWaiterShift == 0 {
					 /*line :1*/ f6Asug("sync: inconsistent mutex state")
				}
				yza2GiF :=  /*line :1*/ int32(mutexLocked - 1<<mutexWaiterShift)
				if !q0LX6Qb || fRQ4D1>>mutexWaiterShift == 1 {

					yza2GiF -= mutexStarving
				}
				 /*line :1*/ atomic.AddInt32(&dMWQaSB.i58pIQ, yza2GiF)
				break
			}
			j5kUAanRO4d = true
			pLFWAdQ = 0
		} else {
			fRQ4D1 = dMWQaSB.i58pIQ
		}
	}

	if race.Enabled {
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(dMWQaSB))
	}
}







func (dMWQaSB *DIRWe11KYlYa) Unlock() {
	if race.Enabled {
		_ = dMWQaSB.i58pIQ
		 /*line :1*/ race.VYdLYvY( /*line :1*/ unsafe.Pointer(dMWQaSB))
	}

	wlm_tUADFO30 :=  /*line :1*/ atomic.AddInt32(&dMWQaSB.i58pIQ, -mutexLocked)
	if wlm_tUADFO30 != 0 {

		 /*line :1*/ dMWQaSB.bKiWKgz(wlm_tUADFO30)
	}
}

func (dMWQaSB *DIRWe11KYlYa) bKiWKgz(wlm_tUADFO30 int32) {
	if (wlm_tUADFO30+mutexLocked)&mutexLocked == 0 {
		 /*line :1*/ h7pm4f2ZL("sync: unlock of unlocked mutex")
	}
	if wlm_tUADFO30&mutexStarving == 0 {
		fRQ4D1 := wlm_tUADFO30
		for {

			if fRQ4D1>>mutexWaiterShift == 0 || fRQ4D1&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}

			wlm_tUADFO30 = (fRQ4D1 - 1<<mutexWaiterShift) | mutexWoken
			if  /*line :1*/ atomic.CompareAndSwapInt32(&dMWQaSB.i58pIQ, fRQ4D1, wlm_tUADFO30) {
				 /*line :1*/ xAzRYQiZaT0(&dMWQaSB.r6mFI6Pfy9q, false, 1)
				return
			}
			fRQ4D1 = dMWQaSB.i58pIQ
		}
	} else {

		 /*line :1*/ xAzRYQiZaT0(&dMWQaSB.r6mFI6Pfy9q, true, 1)
	}
}
