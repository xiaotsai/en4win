//line :1
package fuA83L

import (
	bytes "wLlfRPv"
	binary "yLm9uBQG"
	reflect "reflect"
	syscall "bUKeamF"
	"unsafe"
)

func (y0jiLdYHXNnx *IInspectable) GetIids() (qoiGkaHV []*GUID, zsgC7M1 error) {
	var ooG0a0 uint32
	var j7O2yy uintptr
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ y0jiLdYHXNnx.VTable().G2r2FqTaZ,
		3,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&ooG0a0)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&j7O2yy)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
		return
	}
	defer  /*line :1*/ FXrnmIJ5(j7O2yy)

	qoiGkaHV =  /*line :1*/ make([]*GUID, ooG0a0)
	eNsBw6kM := ooG0a0 *  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(GUID{}))
	wrqiXveP76 := reflect.Ei2Iy0{H7OP0hFU: j7O2yy, MCzI2qhUA1:  /*line :1*/ int(eNsBw6kM), A6DWhwcvtSKZ:  /*line :1*/ int(eNsBw6kM)}
	p6aCsz1eL1iC := *(*[] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&wrqiXveP76))
	gnpgzFo :=  /*line :1*/ bytes.FGmhhHol(p6aCsz1eL1iC)
	for gXX2nE5 := range qoiGkaHV {
		i15rO5WD3Nx := GUID{}
		zsgC7M1 =  /*line :1*/ binary.OrEIDcJ0jtm2(gnpgzFo, binary.BE16BND, &i15rO5WD3Nx)
		if zsgC7M1 != nil {
			return
		}
		qoiGkaHV[gXX2nE5] = &i15rO5WD3Nx
	}
	return
}

func (y0jiLdYHXNnx *IInspectable) GetRuntimeClassName() (deH_KI string, zsgC7M1 error) {
	var fF5eVHvvIt4 OA3qlQMuDJWt
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ y0jiLdYHXNnx.VTable().Ub38TPqgdjyu,
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&fF5eVHvvIt4)),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
		return
	}
	deH_KI =  /*line :1*/ fF5eVHvvIt4.String()
	 /*line :1*/ SCq5K_T3a7e(fF5eVHvvIt4)
	return
}

func (y0jiLdYHXNnx *IInspectable) GetTrustLevel() (gCdBa6 uint32, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4(
		 /*line :1*/ y0jiLdYHXNnx.VTable().Ssw_Naq,
		2,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&gCdBa6)),
		0)
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}
