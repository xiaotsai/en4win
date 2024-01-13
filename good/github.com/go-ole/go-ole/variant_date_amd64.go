//line :1
package fuA83L

import (
	errors "iAHoxjmM"
	syscall "bUKeamF"
	time "fRAfQd_"
	"unsafe"
)


func NT_HQpy3w(rORmQfUB uint64) (time.G4KDkI, error) {
	var hZSvWI03t syscall.H6Qxv2jW
	q0G0d09Te, _, _ :=  /*line :1*/ d8DnSC_c01.Call( /*line :1*/ uintptr(rORmQfUB),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&hZSvWI03t)))
	if q0G0d09Te != 0 {
		return  /*line :1*/ time.FaDSoi2w( /*line :1*/ int(hZSvWI03t.Xmnj_bIFJhJ),  /*line :1*/ time.ZyPFXNxLcwpT(hZSvWI03t.EQp3MHA11V),  /*line :1*/ int(hZSvWI03t.HF4FBjj),  /*line :1*/ int(hZSvWI03t.CySfZqD),  /*line :1*/ int(hZSvWI03t.TFSCFS3IfPS0),  /*line :1*/ int(hZSvWI03t.MRSnx59),  /*line :1*/ int(hZSvWI03t.Lo99feSEOsAo/1000), time.PD1NIUjTJ), nil
	}
	return  /*line :1*/ time.Pc_35oTY(),  /*line :1*/ errors.Su6g6hRoi9X("Could not convert to time, passing current time.")
}
