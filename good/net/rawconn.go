//line :1
package gF9mZs2

import (
	poll "MjXXMR"
	"runtime"
	syscall "bUKeamF"
)

type cfFPDevTPcYC struct {
	sXhoTdSE8wtb *jmJczkl1
}

func (hl8w2lFX *cfFPDevTPcYC) d30HIwxht() bool {
	return hl8w2lFX != nil && hl8w2lFX.sXhoTdSE8wtb != nil
}

func (hl8w2lFX *cfFPDevTPcYC) Control(t5q9DVOD9 func(uintptr)) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.q0t7ga7tBcV5.RawControl(t5q9DVOD9)
	 /*line :1*/ runtime.KeepAlive(hl8w2lFX.sXhoTdSE8wtb)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "raw-control", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return h_ljk48Bm
}

func (hl8w2lFX *cfFPDevTPcYC) Read(t5q9DVOD9 func(uintptr) bool) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.q0t7ga7tBcV5.RawRead(t5q9DVOD9)
	 /*line :1*/ runtime.KeepAlive(hl8w2lFX.sXhoTdSE8wtb)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "raw-read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return h_ljk48Bm
}

func (hl8w2lFX *cfFPDevTPcYC) Write(t5q9DVOD9 func(uintptr) bool) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.q0t7ga7tBcV5.RawWrite(t5q9DVOD9)
	 /*line :1*/ runtime.KeepAlive(hl8w2lFX.sXhoTdSE8wtb)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "raw-write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return h_ljk48Bm
}








func (hl8w2lFX *cfFPDevTPcYC) PollFD() *poll.ENfmDmMaozH {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return nil
	}
	return &hl8w2lFX.sXhoTdSE8wtb.q0t7ga7tBcV5
}

func zQf8had4tml(vyaiiPXm6k_W *jmJczkl1) (*cfFPDevTPcYC, error) {
	return &cfFPDevTPcYC{sXhoTdSE8wtb: vyaiiPXm6k_W}, nil
}

type pGj8DvP9A6 struct {
	cfFPDevTPcYC
}

func (v3upVm *pGj8DvP9A6) Read(func(uintptr) bool) error {
	return syscall.EINVAL
}

func (v3upVm *pGj8DvP9A6) Write(func(uintptr) bool) error {
	return syscall.EINVAL
}

func aJmXTOyud(vyaiiPXm6k_W *jmJczkl1) (*pGj8DvP9A6, error) {
	return &pGj8DvP9A6{cfFPDevTPcYC{sXhoTdSE8wtb: vyaiiPXm6k_W}}, nil
}
