//line :1
//go:build unix || windows

package gF9mZs2

import (
	"runtime"
	syscall "bUKeamF"
)

func visCDO9kf(vyaiiPXm6k_W *jmJczkl1, dvFmDYmQz *U_Uc9la, kQ8_UEhxU GNraIvYhB) error {
	vP562Gd := &syscall.W1r5YRtipLa{P8O2NF0WA: [4]byte{kQ8_UEhxU[0], kQ8_UEhxU[1], kQ8_UEhxU[2], kQ8_UEhxU[3]}}
	if h_ljk48Bm :=  /*line :1*/ bJqj_8(vP562Gd, dvFmDYmQz); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptIPMreq(syscall.IPPROTO_IP, syscall.IP_ADD_MEMBERSHIP, vP562Gd)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}

func tayHWdA(vyaiiPXm6k_W *jmJczkl1, dvFmDYmQz *U_Uc9la) error {
	var ljsCSk int
	if dvFmDYmQz != nil {
		ljsCSk = dvFmDYmQz.OJz1cH7Vb
	}
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptInt(syscall.IPPROTO_IPV6, syscall.IPV6_MULTICAST_IF, ljsCSk)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}

func rVEEX2Rxk(vyaiiPXm6k_W *jmJczkl1, ljsCSk bool) error {
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptInt(syscall.IPPROTO_IPV6, syscall.IPV6_MULTICAST_LOOP,  /*line :1*/ j7bhnpSMpen(ljsCSk))
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}

func bIw1jOVnFhg(vyaiiPXm6k_W *jmJczkl1, dvFmDYmQz *U_Uc9la, kQ8_UEhxU GNraIvYhB) error {
	vP562Gd := &syscall.C5ybc0{}
	 /*line :1*/ copy(vP562Gd.RmQSpqMVN[:], kQ8_UEhxU)
	if dvFmDYmQz != nil {
		vP562Gd.GOMCVrEx1k =  /*line :1*/ uint32(dvFmDYmQz.OJz1cH7Vb)
	}
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptIPv6Mreq(syscall.IPPROTO_IPV6, syscall.IPV6_JOIN_GROUP, vP562Gd)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}
