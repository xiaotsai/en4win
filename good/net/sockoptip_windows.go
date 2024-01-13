//line :1
package gF9mZs2

import (
	os "hPMrClpC"
	"runtime"
	syscall "bUKeamF"
	"unsafe"
)

func iq2nJr(vyaiiPXm6k_W *jmJczkl1, dvFmDYmQz *U_Uc9la) error {
	kQ8_UEhxU, h_ljk48Bm :=  /*line :1*/ lxWp3e(dvFmDYmQz)
	if h_ljk48Bm != nil {
		return  /*line :1*/ os.BTaHHl("setsockopt", h_ljk48Bm)
	}
	var uI7LZDHm6 [4]byte
	 /*line :1*/ copy(uI7LZDHm6[:],  /*line :1*/ kQ8_UEhxU.To4())
	h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Setsockopt(syscall.IPPROTO_IP, syscall.IP_MULTICAST_IF, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&uI7LZDHm6[0])), 4)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}

func kPOWpW616(vyaiiPXm6k_W *jmJczkl1, ljsCSk bool) error {
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptInt(syscall.IPPROTO_IP, syscall.IP_MULTICAST_LOOP,  /*line :1*/ j7bhnpSMpen(ljsCSk))
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}
