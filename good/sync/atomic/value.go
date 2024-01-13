//line :1
package atomic

import (
	"unsafe"
)






type ANnCMU struct {
	tlU5RHpD1Z any
}


type iCwIOEixbMf struct {
	c9F4RBeB	unsafe.Pointer
	faR1g3aL	unsafe.Pointer
}



func (cJrwcXSs6FNC *ANnCMU) Load() (fSC10Rr9Vh any) {
	nczAIhG7ay := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(cJrwcXSs6FNC))
	zj3B2F8B8o :=  /*line :1*/ LoadPointer(&nczAIhG7ay.c9F4RBeB)
	if zj3B2F8B8o == nil || zj3B2F8B8o ==  /*line :1*/ unsafe.Pointer(&adMRcyQNZ) {

		return nil
	}
	xi5ElsoQ802 :=  /*line :1*/ LoadPointer(&nczAIhG7ay.faR1g3aL)
	wbTp3eC_FI5 := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&fSC10Rr9Vh))
	wbTp3eC_FI5.c9F4RBeB = zj3B2F8B8o
	wbTp3eC_FI5.faR1g3aL = xi5ElsoQ802
	return
}

var adMRcyQNZ byte




func (cJrwcXSs6FNC *ANnCMU) Store(fSC10Rr9Vh any) {
	if fSC10Rr9Vh == nil {
		 /*line :1*/ panic("sync/atomic: store of nil value into Value")
	}
	nczAIhG7ay := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(cJrwcXSs6FNC))
	wbTp3eC_FI5 := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&fSC10Rr9Vh))
	for {
		zj3B2F8B8o :=  /*line :1*/ LoadPointer(&nczAIhG7ay.c9F4RBeB)
		if zj3B2F8B8o == nil {

			 /*line :1*/ evPptsk()
			if ! /*line :1*/ B2CfK2wiOXKp(&nczAIhG7ay.c9F4RBeB, nil,  /*line :1*/ unsafe.Pointer(&adMRcyQNZ)) {
				 /*line :1*/ d4ySaOaXhB()
				continue
			}

			 /*line :1*/ Fa2hd0oXquKk(&nczAIhG7ay.faR1g3aL, wbTp3eC_FI5.faR1g3aL)
			 /*line :1*/ Fa2hd0oXquKk(&nczAIhG7ay.c9F4RBeB, wbTp3eC_FI5.c9F4RBeB)
			 /*line :1*/ d4ySaOaXhB()
			return
		}
		if zj3B2F8B8o ==  /*line :1*/ unsafe.Pointer(&adMRcyQNZ) {

			continue
		}

		if zj3B2F8B8o != wbTp3eC_FI5.c9F4RBeB {
			 /*line :1*/ panic("sync/atomic: store of inconsistently typed value into Value")
		}
		 /*line :1*/ Fa2hd0oXquKk(&nczAIhG7ay.faR1g3aL, wbTp3eC_FI5.faR1g3aL)
		return
	}
}






func (cJrwcXSs6FNC *ANnCMU) Swap(vxcCAFH4Bs any) (qVwQ2ijo0 any) {
	if vxcCAFH4Bs == nil {
		 /*line :1*/ panic("sync/atomic: swap of nil value into Value")
	}
	nczAIhG7ay := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(cJrwcXSs6FNC))
	rxN2uoF := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&vxcCAFH4Bs))
	for {
		zj3B2F8B8o :=  /*line :1*/ LoadPointer(&nczAIhG7ay.c9F4RBeB)
		if zj3B2F8B8o == nil {

			 /*line :1*/ evPptsk()
			if ! /*line :1*/ B2CfK2wiOXKp(&nczAIhG7ay.c9F4RBeB, nil,  /*line :1*/ unsafe.Pointer(&adMRcyQNZ)) {
				 /*line :1*/ d4ySaOaXhB()
				continue
			}

			 /*line :1*/ Fa2hd0oXquKk(&nczAIhG7ay.faR1g3aL, rxN2uoF.faR1g3aL)
			 /*line :1*/ Fa2hd0oXquKk(&nczAIhG7ay.c9F4RBeB, rxN2uoF.c9F4RBeB)
			 /*line :1*/ d4ySaOaXhB()
			return nil
		}
		if zj3B2F8B8o ==  /*line :1*/ unsafe.Pointer(&adMRcyQNZ) {

			continue
		}

		if zj3B2F8B8o != rxN2uoF.c9F4RBeB {
			 /*line :1*/ panic("sync/atomic: swap of inconsistently typed value into Value")
		}
		t8Jw31y := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&qVwQ2ijo0))
		t8Jw31y.c9F4RBeB, t8Jw31y.faR1g3aL = rxN2uoF.c9F4RBeB,  /*line :1*/ X8DfWh(&nczAIhG7ay.faR1g3aL, rxN2uoF.faR1g3aL)
		return qVwQ2ijo0
	}
}






func (cJrwcXSs6FNC *ANnCMU) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs any) (fHza20D bool) {
	if vxcCAFH4Bs == nil {
		 /*line :1*/ panic("sync/atomic: compare and swap of nil value into Value")
	}
	nczAIhG7ay := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(cJrwcXSs6FNC))
	rxN2uoF := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&vxcCAFH4Bs))
	t8Jw31y := (* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&qVwQ2ijo0))
	if t8Jw31y.c9F4RBeB != nil && rxN2uoF.c9F4RBeB != t8Jw31y.c9F4RBeB {
		 /*line :1*/ panic("sync/atomic: compare and swap of inconsistently typed values")
	}
	for {
		zj3B2F8B8o :=  /*line :1*/ LoadPointer(&nczAIhG7ay.c9F4RBeB)
		if zj3B2F8B8o == nil {
			if qVwQ2ijo0 != nil {
				return false
			}

			 /*line :1*/ evPptsk()
			if ! /*line :1*/ B2CfK2wiOXKp(&nczAIhG7ay.c9F4RBeB, nil,  /*line :1*/ unsafe.Pointer(&adMRcyQNZ)) {
				 /*line :1*/ d4ySaOaXhB()
				continue
			}

			 /*line :1*/ Fa2hd0oXquKk(&nczAIhG7ay.faR1g3aL, rxN2uoF.faR1g3aL)
			 /*line :1*/ Fa2hd0oXquKk(&nczAIhG7ay.c9F4RBeB, rxN2uoF.c9F4RBeB)
			 /*line :1*/ d4ySaOaXhB()
			return true
		}
		if zj3B2F8B8o ==  /*line :1*/ unsafe.Pointer(&adMRcyQNZ) {

			continue
		}

		if zj3B2F8B8o != rxN2uoF.c9F4RBeB {
			 /*line :1*/ panic("sync/atomic: compare and swap of inconsistently typed value into Value")
		}

		xi5ElsoQ802 :=  /*line :1*/ LoadPointer(&nczAIhG7ay.faR1g3aL)
		var dRzo4Adn any
		(* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&dRzo4Adn)).c9F4RBeB = zj3B2F8B8o
		(* /*line :1*/ iCwIOEixbMf)( /*line :1*/ unsafe.Pointer(&dRzo4Adn)).faR1g3aL = xi5ElsoQ802
		if dRzo4Adn != qVwQ2ijo0 {
			return false
		}
		return  /*line :1*/ B2CfK2wiOXKp(&nczAIhG7ay.faR1g3aL, xi5ElsoQ802, rxN2uoF.faR1g3aL)
	}
}


func evPptsk() int
func d4ySaOaXhB()
