//line :1
package gF9mZs2

import (
	windows "LdptURlN"
	os "hPMrClpC"
	syscall "bUKeamF"
)

func kUqZG__7Y() int {

	return syscall.SOMAXCONN
}

func bbSHyj(nqsgnnaf, dr86pU, iaebRs5X3 int) (syscall.SRlvVjrYa, error) {
	dRoFccaZ, h_ljk48Bm :=  /*line :1*/ trDXFPq2T( /*line :1*/ int32(nqsgnnaf),  /*line :1*/ int32(dr86pU),  /*line :1*/ int32(iaebRs5X3),
		nil, 0, windows.WSA_FLAG_OVERLAPPED|windows.WSA_FLAG_NO_HANDLE_INHERIT)
	if h_ljk48Bm == nil {
		return dRoFccaZ, nil
	}

	 /*line :1*/ syscall.OLKdSMhcsd.RLock()
	dRoFccaZ, h_ljk48Bm =  /*line :1*/ mWPYEZGYiKQ(nqsgnnaf, dr86pU, iaebRs5X3)
	if h_ljk48Bm == nil {
		 /*line :1*/ syscall.W6smj_(dRoFccaZ)
	}
	 /*line :1*/ syscall.OLKdSMhcsd.RUnlock()
	if h_ljk48Bm != nil {
		return syscall.InvalidHandle,  /*line :1*/ os.BTaHHl("socket", h_ljk48Bm)
	}
	return dRoFccaZ, nil
}
