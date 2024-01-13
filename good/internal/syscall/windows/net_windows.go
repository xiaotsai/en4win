//line :1
package LdptURlN

import (
	sync "sync"
	syscall "bUKeamF"
	_ "unsafe"
)

//go:linkname GLhwGjGTr bUKeamF.m2VgBhm
//go:noescape
func GLhwGjGTr(gp3a5i1F syscall.SRlvVjrYa, ilHvMaki1s *syscall.QbWhAp5CqTG, ahjgumbUm uint32, vB40s8jgR *uint32, iK5SVBoZG uint32, uxhDNcje *syscall.Hx8lqxSkV, bS4DAcuB0sz *syscall.ZaNm5QSYC9, bT89xwB *byte) (zc4PmxS error)

//go:linkname BQSWeDJBro18 bUKeamF.pmhM9rv
//go:noescape
func BQSWeDJBro18(gp3a5i1F syscall.SRlvVjrYa, ilHvMaki1s *syscall.QbWhAp5CqTG, ahjgumbUm uint32, vB40s8jgR *uint32, iK5SVBoZG uint32, uxhDNcje *syscall.HKTAy7_BSU2, bS4DAcuB0sz *syscall.ZaNm5QSYC9, bT89xwB *byte) (zc4PmxS error)

const (
	SIO_TCP_INITIAL_RTO	= syscall.IOC_IN | syscall.IOC_VENDOR | 17
	TCP_INITIAL_RTO_UNSPECIFIED_RTT	= ^ /*line :1*/ uint16(0)
	TCP_INITIAL_RTO_NO_SYN_RETRANSMISSIONS	= ^ /*line :1*/ uint8(1)
)

type CSH1eB5vRVAy struct {
	LeR_En3xrj	uint16
	H3T1z7dO	uint8
}

var N7_CEIm1r =  /*line :1*/ sync.DkaH4ad(func() bool {
	var wxDm_c, oB0IVVFuxFR5, fr5X6q1BcGO uint32
	 /*line :1*/ dx4sLYR(&wxDm_c, &oB0IVVFuxFR5, &fr5X6q1BcGO)
	return wxDm_c >= 10 && fr5X6q1BcGO&0xffff >= 16299
})

//go:linkname dx4sLYR bUKeamF.kjlGhacz
//go:noescape
func dx4sLYR(gXKUy11Yd *uint32, axr5JVA3y *uint32, jedBzWsxI *uint32)
