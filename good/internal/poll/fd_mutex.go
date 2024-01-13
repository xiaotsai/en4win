//line :1
package MjXXMR

import atomic "sync/atomic"




type hvIBHvdqmStf struct {
	r8KaF5O9ITu	uint64
	p8Gab0	uint32
	iVzLO6uEXF	uint32
}








const (
	mutexClosed	= 1 << 0
	mutexRLock	= 1 << 1
	mutexWLock	= 1 << 2
	mutexRef	= 1 << 3
	mutexRefMask	= (1<<20 - 1) << 3
	mutexRWait	= 1 << 23
	mutexRMask	= (1<<20 - 1) << 23
	mutexWWait	= 1 << 43
	mutexWMask	= (1<<20 - 1) << 43
)

const overflowMsg = "too many concurrent operations on a single file or socket (max 1048575)"



func (litNXbr *hvIBHvdqmStf) yB3OFgBy1() bool {
	for {
		oy_jy9AhMj :=  /*line :1*/ atomic.LoadUint64(&litNXbr.r8KaF5O9ITu)
		if oy_jy9AhMj&mutexClosed != 0 {
			return false
		}
		_rjYU4D_ := oy_jy9AhMj + mutexRef
		if _rjYU4D_&mutexRefMask == 0 {
			 /*line :1*/ panic(overflowMsg)
		}
		if  /*line :1*/ atomic.CompareAndSwapUint64(&litNXbr.r8KaF5O9ITu, oy_jy9AhMj, _rjYU4D_) {
			return true
		}
	}
}



func (litNXbr *hvIBHvdqmStf) wMwcmDUo() bool {
	for {
		oy_jy9AhMj :=  /*line :1*/ atomic.LoadUint64(&litNXbr.r8KaF5O9ITu)
		if oy_jy9AhMj&mutexClosed != 0 {
			return false
		}

		_rjYU4D_ := (oy_jy9AhMj | mutexClosed) + mutexRef
		if _rjYU4D_&mutexRefMask == 0 {
			 /*line :1*/ panic(overflowMsg)
		}

		_rjYU4D_ &^= mutexRMask | mutexWMask
		if  /*line :1*/ atomic.CompareAndSwapUint64(&litNXbr.r8KaF5O9ITu, oy_jy9AhMj, _rjYU4D_) {

			for oy_jy9AhMj&mutexRMask != 0 {
				oy_jy9AhMj -= mutexRWait
				 /*line :1*/ g6gFf1(&litNXbr.p8Gab0)
			}
			for oy_jy9AhMj&mutexWMask != 0 {
				oy_jy9AhMj -= mutexWWait
				 /*line :1*/ g6gFf1(&litNXbr.iVzLO6uEXF)
			}
			return true
		}
	}
}



func (litNXbr *hvIBHvdqmStf) xmYUot4V() bool {
	for {
		oy_jy9AhMj :=  /*line :1*/ atomic.LoadUint64(&litNXbr.r8KaF5O9ITu)
		if oy_jy9AhMj&mutexRefMask == 0 {
			 /*line :1*/ panic("inconsistent poll.fdMutex")
		}
		_rjYU4D_ := oy_jy9AhMj - mutexRef
		if  /*line :1*/ atomic.CompareAndSwapUint64(&litNXbr.r8KaF5O9ITu, oy_jy9AhMj, _rjYU4D_) {
			return _rjYU4D_&(mutexClosed|mutexRefMask) == mutexClosed
		}
	}
}



func (litNXbr *hvIBHvdqmStf) g3no9daDONS(boGBxgkTgp bool) bool {
	var h2zBfwoVqYH, cZYsSeT, pT4ToGosXZ uint64
	var kbEYnTwQu *uint32
	if boGBxgkTgp {
		h2zBfwoVqYH = mutexRLock
		cZYsSeT = mutexRWait
		pT4ToGosXZ = mutexRMask
		kbEYnTwQu = &litNXbr.p8Gab0
	} else {
		h2zBfwoVqYH = mutexWLock
		cZYsSeT = mutexWWait
		pT4ToGosXZ = mutexWMask
		kbEYnTwQu = &litNXbr.iVzLO6uEXF
	}
	for {
		oy_jy9AhMj :=  /*line :1*/ atomic.LoadUint64(&litNXbr.r8KaF5O9ITu)
		if oy_jy9AhMj&mutexClosed != 0 {
			return false
		}
		var _rjYU4D_ uint64
		if oy_jy9AhMj&h2zBfwoVqYH == 0 {

			_rjYU4D_ = (oy_jy9AhMj | h2zBfwoVqYH) + mutexRef
			if _rjYU4D_&mutexRefMask == 0 {
				 /*line :1*/ panic(overflowMsg)
			}
		} else {

			_rjYU4D_ = oy_jy9AhMj + cZYsSeT
			if _rjYU4D_&pT4ToGosXZ == 0 {
				 /*line :1*/ panic(overflowMsg)
			}
		}
		if  /*line :1*/ atomic.CompareAndSwapUint64(&litNXbr.r8KaF5O9ITu, oy_jy9AhMj, _rjYU4D_) {
			if oy_jy9AhMj&h2zBfwoVqYH == 0 {
				return true
			}
			 /*line :1*/ jPXZV7q0q5(kbEYnTwQu)

		}
	}
}



func (litNXbr *hvIBHvdqmStf) b2cl2FJZo(boGBxgkTgp bool) bool {
	var h2zBfwoVqYH, cZYsSeT, pT4ToGosXZ uint64
	var kbEYnTwQu *uint32
	if boGBxgkTgp {
		h2zBfwoVqYH = mutexRLock
		cZYsSeT = mutexRWait
		pT4ToGosXZ = mutexRMask
		kbEYnTwQu = &litNXbr.p8Gab0
	} else {
		h2zBfwoVqYH = mutexWLock
		cZYsSeT = mutexWWait
		pT4ToGosXZ = mutexWMask
		kbEYnTwQu = &litNXbr.iVzLO6uEXF
	}
	for {
		oy_jy9AhMj :=  /*line :1*/ atomic.LoadUint64(&litNXbr.r8KaF5O9ITu)
		if oy_jy9AhMj&h2zBfwoVqYH == 0 || oy_jy9AhMj&mutexRefMask == 0 {
			 /*line :1*/ panic("inconsistent poll.fdMutex")
		}

		_rjYU4D_ := (oy_jy9AhMj &^ h2zBfwoVqYH) - mutexRef
		if oy_jy9AhMj&pT4ToGosXZ != 0 {
			_rjYU4D_ -= cZYsSeT
		}
		if  /*line :1*/ atomic.CompareAndSwapUint64(&litNXbr.r8KaF5O9ITu, oy_jy9AhMj, _rjYU4D_) {
			if oy_jy9AhMj&pT4ToGosXZ != 0 {
				 /*line :1*/ g6gFf1(kbEYnTwQu)
			}
			return _rjYU4D_&(mutexClosed|mutexRefMask) == mutexClosed
		}
	}
}


func jPXZV7q0q5(tjMkgG *uint32)
func g6gFf1(tjMkgG *uint32)



func (s_ZR_8 *ENfmDmMaozH) yB3OFgBy1() error {
	if ! /*line :1*/ s_ZR_8.eYtKOCiM.yB3OFgBy1() {
		return  /*line :1*/ amQJqF(s_ZR_8.l3fWSyWa)
	}
	return nil
}




func (s_ZR_8 *ENfmDmMaozH) xmYUot4V() error {
	if  /*line :1*/ s_ZR_8.eYtKOCiM.xmYUot4V() {
		return  /*line :1*/ s_ZR_8.twRFlT6Xc()
	}
	return nil
}



func (s_ZR_8 *ENfmDmMaozH) kMJTtXNo() error {
	if ! /*line :1*/ s_ZR_8.eYtKOCiM.g3no9daDONS(true) {
		return  /*line :1*/ amQJqF(s_ZR_8.l3fWSyWa)
	}
	return nil
}




func (s_ZR_8 *ENfmDmMaozH) ldjlnUNEdCg6() {
	if  /*line :1*/ s_ZR_8.eYtKOCiM.b2cl2FJZo(true) {
		 /*line :1*/ s_ZR_8.twRFlT6Xc()
	}
}



func (s_ZR_8 *ENfmDmMaozH) aAav5o7pLI() error {
	if ! /*line :1*/ s_ZR_8.eYtKOCiM.g3no9daDONS(false) {
		return  /*line :1*/ amQJqF(s_ZR_8.l3fWSyWa)
	}
	return nil
}




func (s_ZR_8 *ENfmDmMaozH) gb0d2X() {
	if  /*line :1*/ s_ZR_8.eYtKOCiM.b2cl2FJZo(false) {
		 /*line :1*/ s_ZR_8.twRFlT6Xc()
	}
}
