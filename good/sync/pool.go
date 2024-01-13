//line :1
package sync

import (
	race "V0p_iv"
	"runtime"
	atomic "sync/atomic"
	"unsafe"
)

type OrP5FGPq struct {
	pk1kvRF	tzBuysuv

	jPwJrxtqyH	unsafe.Pointer
	n2iXnbS	uintptr

	pZ3gatQ	unsafe.Pointer
	i0CEygyOW_uR	uintptr

	IYbTy050ek	func() any
}

type maNai84Qeo struct {
	kyP8WyAB	any
	gtgauqq_eW4	vwraZoCcx
}

type nK8Z7MxLv struct {
	maNai84Qeo

	gNQZM30inRY	[128 -  /*line :1*/ unsafe.Sizeof(maNai84Qeo{})%128]byte
}

func q3i2dA(kXg2O9i uint32) uint32

var iIer0dzAYCaF [128]uint64

func aZYX94YG(cLQZ4qp0 any) unsafe.Pointer {
	_caa2MSBFai5 :=  /*line :1*/ uintptr((*[2] /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(&cLQZ4qp0))[1])
	bI2P391 :=  /*line :1*/ uint32(( /*line :1*/ uint64( /*line :1*/ uint32(_caa2MSBFai5)) * 0x85ebca6b) >> 16)
	return  /*line :1*/ unsafe.Pointer(&iIer0dzAYCaF[bI2P391% /*line :1*/ uint32( /*line :1*/ len(iIer0dzAYCaF))])
}

func (lcXhtO *OrP5FGPq) Put(cLQZ4qp0 any) {
	if cLQZ4qp0 == nil {
		return
	}
	if race.Enabled {
		if  /*line :1*/ q3i2dA(4) == 0 {

			return
		}
		 /*line :1*/ race.AQrXiLDh3RTu( /*line :1*/ aZYX94YG(cLQZ4qp0))
		 /*line :1*/ race.F_wVXvO1u_kM()
	}
	fwPZAw_, _ :=  /*line :1*/ lcXhtO.wtZgfdDQ3c()
	if fwPZAw_.kyP8WyAB == nil {
		fwPZAw_.kyP8WyAB = cLQZ4qp0
	} else {
		 /*line :1*/ fwPZAw_.gtgauqq_eW4.lzZ69GP3MxUu(cLQZ4qp0)
	}
	 /*line :1*/ zpmT4w()
	if race.Enabled {
		 /*line :1*/ race.UStrOKX()
	}
}

func (lcXhtO *OrP5FGPq) Get() any {
	if race.Enabled {
		 /*line :1*/ race.F_wVXvO1u_kM()
	}
	fwPZAw_, idl064I3TRg :=  /*line :1*/ lcXhtO.wtZgfdDQ3c()
	cLQZ4qp0 := fwPZAw_.kyP8WyAB
	fwPZAw_.kyP8WyAB = nil
	if cLQZ4qp0 == nil {

		cLQZ4qp0, _ =  /*line :1*/ fwPZAw_.gtgauqq_eW4.eod4ht()
		if cLQZ4qp0 == nil {
			cLQZ4qp0 =  /*line :1*/ lcXhtO.qLaVjO(idl064I3TRg)
		}
	}
	 /*line :1*/ zpmT4w()
	if race.Enabled {
		 /*line :1*/ race.UStrOKX()
		if cLQZ4qp0 != nil {
			 /*line :1*/ race.X5ZIZytZ( /*line :1*/ aZYX94YG(cLQZ4qp0))
		}
	}
	if cLQZ4qp0 == nil && lcXhtO.IYbTy050ek != nil {
		cLQZ4qp0 =  /*line :1*/ lcXhtO.IYbTy050ek()
	}
	return cLQZ4qp0
}

func (lcXhtO *OrP5FGPq) qLaVjO(idl064I3TRg int) any {

	fgx09c_VhIsS :=  /*line :1*/ runtime_LoadAcquintptr(&lcXhtO.n2iXnbS)
	fVC4PecXoU := lcXhtO.jPwJrxtqyH

	for cmSN7RqmPwHo := 0; cmSN7RqmPwHo <  /*line :1*/ int(fgx09c_VhIsS); cmSN7RqmPwHo++ {
		fwPZAw_ :=  /*line :1*/ jwV82TVX4(fVC4PecXoU, (idl064I3TRg+cmSN7RqmPwHo+1)% /*line :1*/ int(fgx09c_VhIsS))
		if cLQZ4qp0, _ :=  /*line :1*/ fwPZAw_.gtgauqq_eW4.ulbgyS0tgMVd(); cLQZ4qp0 != nil {
			return cLQZ4qp0
		}
	}

	fgx09c_VhIsS =  /*line :1*/ atomic.LoadUintptr(&lcXhtO.i0CEygyOW_uR)
	if  /*line :1*/ uintptr(idl064I3TRg) >= fgx09c_VhIsS {
		return nil
	}
	fVC4PecXoU = lcXhtO.pZ3gatQ
	fwPZAw_ :=  /*line :1*/ jwV82TVX4(fVC4PecXoU, idl064I3TRg)
	if cLQZ4qp0 := fwPZAw_.kyP8WyAB; cLQZ4qp0 != nil {
		fwPZAw_.kyP8WyAB = nil
		return cLQZ4qp0
	}
	for cmSN7RqmPwHo := 0; cmSN7RqmPwHo <  /*line :1*/ int(fgx09c_VhIsS); cmSN7RqmPwHo++ {
		fwPZAw_ :=  /*line :1*/ jwV82TVX4(fVC4PecXoU, (idl064I3TRg+cmSN7RqmPwHo)% /*line :1*/ int(fgx09c_VhIsS))
		if cLQZ4qp0, _ :=  /*line :1*/ fwPZAw_.gtgauqq_eW4.ulbgyS0tgMVd(); cLQZ4qp0 != nil {
			return cLQZ4qp0
		}
	}

	 /*line :1*/ atomic.StoreUintptr(&lcXhtO.i0CEygyOW_uR, 0)

	return nil
}

func (lcXhtO *OrP5FGPq) wtZgfdDQ3c() (*nK8Z7MxLv, int) {
	idl064I3TRg :=  /*line :1*/ k1sHsm7()

	g1gaB58j_QjX :=  /*line :1*/ runtime_LoadAcquintptr(&lcXhtO.n2iXnbS)
	fwPZAw_ := lcXhtO.jPwJrxtqyH
	if  /*line :1*/ uintptr(idl064I3TRg) < g1gaB58j_QjX {
		return  /*line :1*/ jwV82TVX4(fwPZAw_, idl064I3TRg), idl064I3TRg
	}
	return  /*line :1*/ lcXhtO.e8KJ5i()
}

func (lcXhtO *OrP5FGPq) e8KJ5i() (*nK8Z7MxLv, int) {

	 /*line :1*/ zpmT4w()
	 /*line :1*/ uLBexla4xKrg.Lock()
	defer  /*line :1*/ uLBexla4xKrg.Unlock()
	idl064I3TRg :=  /*line :1*/ k1sHsm7()

	g1gaB58j_QjX := lcXhtO.n2iXnbS
	fwPZAw_ := lcXhtO.jPwJrxtqyH
	if  /*line :1*/ uintptr(idl064I3TRg) < g1gaB58j_QjX {
		return  /*line :1*/ jwV82TVX4(fwPZAw_, idl064I3TRg), idl064I3TRg
	}
	if lcXhtO.jPwJrxtqyH == nil {
		h6BmjCZrSt =  /*line :1*/ append(h6BmjCZrSt, lcXhtO)
	}

	fgx09c_VhIsS :=  /*line :1*/ runtime.GOMAXPROCS(0)
	qA9RC6rc9I6 :=  /*line :1*/ make([]nK8Z7MxLv, fgx09c_VhIsS)
	 /*line :1*/ atomic.Fa2hd0oXquKk(&lcXhtO.jPwJrxtqyH,  /*line :1*/ unsafe.Pointer(&qA9RC6rc9I6[0]))
	 /*line :1*/ runtime_StoreReluintptr(&lcXhtO.n2iXnbS,  /*line :1*/ uintptr(fgx09c_VhIsS))
	return &qA9RC6rc9I6[idl064I3TRg], idl064I3TRg
}

func wvHpG7DxN() {

	for _, lcXhtO := range jqfosL {
		lcXhtO.pZ3gatQ = nil
		lcXhtO.i0CEygyOW_uR = 0
	}

	for _, lcXhtO := range h6BmjCZrSt {
		lcXhtO.pZ3gatQ = lcXhtO.jPwJrxtqyH
		lcXhtO.i0CEygyOW_uR = lcXhtO.n2iXnbS
		lcXhtO.jPwJrxtqyH = nil
		lcXhtO.n2iXnbS = 0
	}

	jqfosL, h6BmjCZrSt = h6BmjCZrSt, nil
}

var (
	uLBexla4xKrg	DIRWe11KYlYa

	h6BmjCZrSt	[]*OrP5FGPq

	jqfosL	[]*OrP5FGPq
)

func init() {
	 /*line :1*/ ew66MMK(wvHpG7DxN)
}

func jwV82TVX4(fwPZAw_ unsafe.Pointer, cmSN7RqmPwHo int) *nK8Z7MxLv {
	fxGFWUmzw5 :=  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(fwPZAw_) +  /*line :1*/ uintptr(cmSN7RqmPwHo)* /*line :1*/ unsafe.Sizeof(nK8Z7MxLv{}))
	return (* /*line :1*/ nK8Z7MxLv)(fxGFWUmzw5)
}

func ew66MMK(axSMAwRRLzT func())
func k1sHsm7() int
func zpmT4w()

//go:linkname runtime_LoadAcquintptr runtime/internal/atomic.LoadAcquintptr
func runtime_LoadAcquintptr(_caa2MSBFai5 *uintptr) uintptr

//go:linkname runtime_StoreReluintptr runtime/internal/atomic.StoreReluintptr
func runtime_StoreReluintptr(_caa2MSBFai5 *uintptr, gk1cOEZT5i1 uintptr) uintptr
