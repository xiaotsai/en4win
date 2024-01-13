//line :1
package sync

import (
	atomic "sync/atomic"
	"unsafe"
)

























type SQLdbUZH635 struct {
	kLeb2yw	tzBuysuv

	
	ZugYvFg2s8	OXw5jPmA

	x4hCvDi	tUPeHtahsCRH
	yWSYnJ3IX	ef6uVlT
}


func UxSfxfw9(fwPZAw_ OXw5jPmA) *SQLdbUZH635 {
	return &SQLdbUZH635{ZugYvFg2s8: fwPZAw_}
}
















func (ew6gz7n0m *SQLdbUZH635) Wait() {
	 /*line :1*/ ew6gz7n0m.yWSYnJ3IX._wJwak()
	c29EVbGUJYy1 :=  /*line :1*/ vL42XlKhVlt(&ew6gz7n0m.x4hCvDi)
	 /*line :1*/ ew6gz7n0m.ZugYvFg2s8.Unlock()
	 /*line :1*/ mbk0gzYcQJ(&ew6gz7n0m.x4hCvDi, c29EVbGUJYy1)
	 /*line :1*/ ew6gz7n0m.ZugYvFg2s8.Lock()
}








func (ew6gz7n0m *SQLdbUZH635) Signal() {
	 /*line :1*/ ew6gz7n0m.yWSYnJ3IX._wJwak()
	 /*line :1*/ niBbH7ABy1uh(&ew6gz7n0m.x4hCvDi)
}





func (ew6gz7n0m *SQLdbUZH635) Broadcast() {
	 /*line :1*/ ew6gz7n0m.yWSYnJ3IX._wJwak()
	 /*line :1*/ cr8Chd(&ew6gz7n0m.x4hCvDi)
}


type ef6uVlT uintptr

func (ew6gz7n0m *ef6uVlT) _wJwak() {
	if  /*line :1*/ uintptr(*ew6gz7n0m) !=  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ew6gz7n0m)) &&
		! /*line :1*/ atomic.CompareAndSwapUintptr((* /*line :1*/ uintptr)(ew6gz7n0m), 0,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ew6gz7n0m))) &&
		 /*line :1*/ uintptr(*ew6gz7n0m) !=  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ew6gz7n0m)) {
		 /*line :1*/ panic("sync.Cond is copied")
	}
}








type tzBuysuv struct{}


func (*tzBuysuv) Lock()	{}
func (*tzBuysuv) Unlock()	{}
