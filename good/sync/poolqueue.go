//line :1
package sync

import (
	atomic "sync/atomic"
	"unsafe"
)








type bK3pXaKsSha struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	sGvICdeIFpqn	uint64

	
	
	
	
	
	
	
	
	gdi6MGkMSZ	[]nOiyLF5S8V2
}

type nOiyLF5S8V2 struct {
	ckkHCUM, yUhGavCQfE unsafe.Pointer
}

const dequeueBits = 32






const dequeueLimit = (1 << dequeueBits) / 4




type ae43vY *struct{}

func (vPvqTph8Beb *bK3pXaKsSha) e90qTwsuO(jM7gmxX4WHR uint64) (fprJoAlsUFvg, fJW1yq8 uint32) {
	const mask = 1<<dequeueBits - 1
	fprJoAlsUFvg =  /*line :1*/ uint32((jM7gmxX4WHR >> dequeueBits) & mask)
	fJW1yq8 =  /*line :1*/ uint32(jM7gmxX4WHR & mask)
	return
}

func (vPvqTph8Beb *bK3pXaKsSha) qAP4t2(fprJoAlsUFvg, fJW1yq8 uint32) uint64 {
	const mask = 1<<dequeueBits - 1
	return ( /*line :1*/ uint64(fprJoAlsUFvg) << dequeueBits) |
		 /*line :1*/ uint64(fJW1yq8&mask)
}



func (vPvqTph8Beb *bK3pXaKsSha) lzZ69GP3MxUu(gk1cOEZT5i1 any) bool {
	jM7gmxX4WHR :=  /*line :1*/ atomic.LoadUint64(&vPvqTph8Beb.sGvICdeIFpqn)
	fprJoAlsUFvg, fJW1yq8 :=  /*line :1*/ vPvqTph8Beb.e90qTwsuO(jM7gmxX4WHR)
	if (fJW1yq8+ /*line :1*/ uint32( /*line :1*/ len(vPvqTph8Beb.gdi6MGkMSZ)))&(1<<dequeueBits-1) == fprJoAlsUFvg {

		return false
	}
	mgrfWFYvyr := &vPvqTph8Beb.gdi6MGkMSZ[fprJoAlsUFvg& /*line :1*/ uint32( /*line :1*/ len(vPvqTph8Beb.gdi6MGkMSZ)-1)]

	tW51dVgQkXB :=  /*line :1*/ atomic.LoadPointer(&mgrfWFYvyr.ckkHCUM)
	if tW51dVgQkXB != nil {

		return false
	}

	if gk1cOEZT5i1 == nil {
		gk1cOEZT5i1 =  /*line :1*/ ae43vY(nil)
	}
	*(* /*line :1*/ any)( /*line :1*/ unsafe.Pointer(mgrfWFYvyr)) = gk1cOEZT5i1

	 /*line :1*/ atomic.AddUint64(&vPvqTph8Beb.sGvICdeIFpqn, 1<<dequeueBits)
	return true
}




func (vPvqTph8Beb *bK3pXaKsSha) eod4ht() (any, bool) {
	var mgrfWFYvyr *nOiyLF5S8V2
	for {
		jM7gmxX4WHR :=  /*line :1*/ atomic.LoadUint64(&vPvqTph8Beb.sGvICdeIFpqn)
		fprJoAlsUFvg, fJW1yq8 :=  /*line :1*/ vPvqTph8Beb.e90qTwsuO(jM7gmxX4WHR)
		if fJW1yq8 == fprJoAlsUFvg {

			return nil, false
		}

		fprJoAlsUFvg--
		nKzeu_xQY :=  /*line :1*/ vPvqTph8Beb.qAP4t2(fprJoAlsUFvg, fJW1yq8)
		if  /*line :1*/ atomic.CompareAndSwapUint64(&vPvqTph8Beb.sGvICdeIFpqn, jM7gmxX4WHR, nKzeu_xQY) {

			mgrfWFYvyr = &vPvqTph8Beb.gdi6MGkMSZ[fprJoAlsUFvg& /*line :1*/ uint32( /*line :1*/ len(vPvqTph8Beb.gdi6MGkMSZ)-1)]
			break
		}
	}

	gk1cOEZT5i1 := *(* /*line :1*/ any)( /*line :1*/ unsafe.Pointer(mgrfWFYvyr))
	if gk1cOEZT5i1 ==  /*line :1*/ ae43vY(nil) {
		gk1cOEZT5i1 = nil
	}

	*mgrfWFYvyr = nOiyLF5S8V2{}
	return gk1cOEZT5i1, true
}




func (vPvqTph8Beb *bK3pXaKsSha) ulbgyS0tgMVd() (any, bool) {
	var mgrfWFYvyr *nOiyLF5S8V2
	for {
		jM7gmxX4WHR :=  /*line :1*/ atomic.LoadUint64(&vPvqTph8Beb.sGvICdeIFpqn)
		fprJoAlsUFvg, fJW1yq8 :=  /*line :1*/ vPvqTph8Beb.e90qTwsuO(jM7gmxX4WHR)
		if fJW1yq8 == fprJoAlsUFvg {

			return nil, false
		}

		nKzeu_xQY :=  /*line :1*/ vPvqTph8Beb.qAP4t2(fprJoAlsUFvg, fJW1yq8+1)
		if  /*line :1*/ atomic.CompareAndSwapUint64(&vPvqTph8Beb.sGvICdeIFpqn, jM7gmxX4WHR, nKzeu_xQY) {

			mgrfWFYvyr = &vPvqTph8Beb.gdi6MGkMSZ[fJW1yq8& /*line :1*/ uint32( /*line :1*/ len(vPvqTph8Beb.gdi6MGkMSZ)-1)]
			break
		}
	}

	gk1cOEZT5i1 := *(* /*line :1*/ any)( /*line :1*/ unsafe.Pointer(mgrfWFYvyr))
	if gk1cOEZT5i1 ==  /*line :1*/ ae43vY(nil) {
		gk1cOEZT5i1 = nil
	}

	mgrfWFYvyr.yUhGavCQfE = nil
	 /*line :1*/ atomic.Fa2hd0oXquKk(&mgrfWFYvyr.ckkHCUM, nil)

	return gk1cOEZT5i1, true
}








type vwraZoCcx struct {
	
	
	o5VSQTv	*kaG3i5hD

	
	
	lYR9D1t7GKo	*kaG3i5hD
}

type kaG3i5hD struct {
	bK3pXaKsSha

	
	
	
	
	
	
	
	
	
	
	aMZGE8VYP, dLKu_5WG934	*kaG3i5hD
}

func sMfW5YdRfVQ_(ihdtLnYe0K **kaG3i5hD, nBR5jJPd6nc3 *kaG3i5hD) {
	 /*line :1*/ atomic.Fa2hd0oXquKk((* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(ihdtLnYe0K)),  /*line :1*/ unsafe.Pointer(nBR5jJPd6nc3))
}

func wR22VZQs(ihdtLnYe0K **kaG3i5hD) *kaG3i5hD {
	return (* /*line :1*/ kaG3i5hD)( /*line :1*/ atomic.LoadPointer((* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(ihdtLnYe0K))))
}

func (ew6gz7n0m *vwraZoCcx) lzZ69GP3MxUu(gk1cOEZT5i1 any) {
	vPvqTph8Beb := ew6gz7n0m.o5VSQTv
	if vPvqTph8Beb == nil {
		
		const initSize = 8	
		vPvqTph8Beb =  /*line :1*/ new(kaG3i5hD)
		vPvqTph8Beb.gdi6MGkMSZ =  /*line :1*/ make([]nOiyLF5S8V2, initSize)
		ew6gz7n0m.o5VSQTv = vPvqTph8Beb
		 /*line :1*/ sMfW5YdRfVQ_(&ew6gz7n0m.lYR9D1t7GKo, vPvqTph8Beb)
	}

	if  /*line :1*/ vPvqTph8Beb.lzZ69GP3MxUu(gk1cOEZT5i1) {
		return
	}

	bcnA_Gqs :=  /*line :1*/ len(vPvqTph8Beb.gdi6MGkMSZ) * 2
	if bcnA_Gqs >= dequeueLimit {

		bcnA_Gqs = dequeueLimit
	}

	ii53LYE := &kaG3i5hD{dLKu_5WG934: vPvqTph8Beb}
	ii53LYE.gdi6MGkMSZ =  /*line :1*/ make([]nOiyLF5S8V2, bcnA_Gqs)
	ew6gz7n0m.o5VSQTv = ii53LYE
	 /*line :1*/ sMfW5YdRfVQ_(&vPvqTph8Beb.aMZGE8VYP, ii53LYE)
	 /*line :1*/ ii53LYE.lzZ69GP3MxUu(gk1cOEZT5i1)
}

func (ew6gz7n0m *vwraZoCcx) eod4ht() (any, bool) {
	vPvqTph8Beb := ew6gz7n0m.o5VSQTv
	for vPvqTph8Beb != nil {
		if gk1cOEZT5i1, qxRgxXc1hZi :=  /*line :1*/ vPvqTph8Beb.eod4ht(); qxRgxXc1hZi {
			return gk1cOEZT5i1, qxRgxXc1hZi
		}

		vPvqTph8Beb =  /*line :1*/ wR22VZQs(&vPvqTph8Beb.dLKu_5WG934)
	}
	return nil, false
}

func (ew6gz7n0m *vwraZoCcx) ulbgyS0tgMVd() (any, bool) {
	vPvqTph8Beb :=  /*line :1*/ wR22VZQs(&ew6gz7n0m.lYR9D1t7GKo)
	if vPvqTph8Beb == nil {
		return nil, false
	}

	for {

		ii53LYE :=  /*line :1*/ wR22VZQs(&vPvqTph8Beb.aMZGE8VYP)

		if gk1cOEZT5i1, qxRgxXc1hZi :=  /*line :1*/ vPvqTph8Beb.ulbgyS0tgMVd(); qxRgxXc1hZi {
			return gk1cOEZT5i1, qxRgxXc1hZi
		}

		if ii53LYE == nil {

			return nil, false
		}

		if  /*line :1*/ atomic.B2CfK2wiOXKp((* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(&ew6gz7n0m.lYR9D1t7GKo)),  /*line :1*/ unsafe.Pointer(vPvqTph8Beb),  /*line :1*/ unsafe.Pointer(ii53LYE)) {

			 /*line :1*/ sMfW5YdRfVQ_(&ii53LYE.dLKu_5WG934, nil)
		}
		vPvqTph8Beb = ii53LYE
	}
}
