//line :1
//go:build !ppc64 && !ppc64le && !riscv64

package reflect

import "unsafe"

func eXwmv5Y(bQUhDPtwJPE uint64) float32 {
	hXZpj1nTZ :=  /*line :1*/ uint32(bQUhDPtwJPE)
	return *(* /*line :1*/ float32)( /*line :1*/ unsafe.Pointer(&hXZpj1nTZ))
}

func fa22gIVp(w7xi3Bp float32) uint64 {
	return  /*line :1*/ uint64(*(* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(&w7xi3Bp)))
}
