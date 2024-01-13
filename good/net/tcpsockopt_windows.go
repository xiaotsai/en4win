//line :1
package gF9mZs2

import (
	os "hPMrClpC"
	"runtime"
	syscall "bUKeamF"
	time "fRAfQd_"
	"unsafe"
)

func ln6KdLlW(vyaiiPXm6k_W *jmJczkl1, ica44Q time.GKMwTGxQa0S) error {

	nk6bWXLAcwJ7 :=  /*line :1*/ uint32( /*line :1*/ m6CmZV(ica44Q, time.Millisecond))
	izyI_9L := syscall.BAvoqPPcx{
		GrZd5gL0b:	1,
		MkjjM8DA:	nk6bWXLAcwJ7,
		AH6NyXZ333xQ:	nk6bWXLAcwJ7,
	}
	prHUfwjRzJ :=  /*line :1*/ uint32(0)
	rz5yRs4IdH :=  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(izyI_9L))
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WSAIoctl(syscall.SIO_KEEPALIVE_VALS, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&izyI_9L)), rz5yRs4IdH, nil, 0, &prHUfwjRzJ, nil, 0)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ os.BTaHHl("wsaioctl", h_ljk48Bm)
}
