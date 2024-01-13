//line :1
//go:build windows

package pu5p3KX8sW

import (
	windows "LdptURlN"
	syscall "bUKeamF"
	"unsafe"
)

func CYBtWdq6mE(kOtrKuc6jO *syscall.YxcIO6yPaPK) (ldStFTRqxbq []string, j06VAtWM_ error) {
	if kOtrKuc6jO == nil || kOtrKuc6jO.Z7x_Yh == 0 {
		return  /*line :1*/ syscall.FDhTFlUY(), nil
	}
	var fmlu5oK *uint16
	j06VAtWM_ =  /*line :1*/ windows.QPLQfl3(&fmlu5oK, kOtrKuc6jO.Z7x_Yh, false)
	if j06VAtWM_ != nil {
		return nil, j06VAtWM_
	}
	defer  /*line :1*/ windows.Bnr47zSrA6G(fmlu5oK)

	const size =  /*line :1*/ unsafe.Sizeof(*fmlu5oK)
	for *fmlu5oK != 0 {

		ybRhdS :=  /*line :1*/ unsafe.Add( /*line :1*/ unsafe.Pointer(fmlu5oK), size)
		for *(* /*line :1*/ uint16)(ybRhdS) != 0 {
			ybRhdS =  /*line :1*/ unsafe.Add(ybRhdS, size)
		}

		iaKJSjOA :=  /*line :1*/ unsafe.Slice(fmlu5oK, ( /*line :1*/ uintptr(ybRhdS)- /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fmlu5oK)))/2)
		ldStFTRqxbq =  /*line :1*/ append(ldStFTRqxbq,  /*line :1*/ syscall.AODVXp8NM3sd(iaKJSjOA))
		fmlu5oK = (* /*line :1*/ uint16)( /*line :1*/ unsafe.Add(ybRhdS, size))
	}
	return
}
