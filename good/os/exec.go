//line :1
package hPMrClpC

import (
	errors "iAHoxjmM"
	testlog "X0MraxO81Me"
	"runtime"
	sync "sync"
	atomic "sync/atomic"
	syscall "bUKeamF"
	time "fRAfQd_"
)


var A4466R3 =  /*line :1*/ errors.Su6g6hRoi9X("os: process already finished")


type P2XtkGf3 struct {
	EQoGDb	int
	oNLhPyDvzm9	uintptr	
	g6qmk4H	atomic.Akm7q89_	
	lvkJUzodVTDU	sync.V1sRgjwAglJt	
}

func zvDudDSfa(gAIGP2p int, iVmSWxZA4 uintptr) *P2XtkGf3 {
	baV7tO7i5 := &P2XtkGf3{EQoGDb: gAIGP2p, oNLhPyDvzm9: iVmSWxZA4}
	 /*line :1*/ runtime.SetFinalizer(baV7tO7i5, (*P2XtkGf3).Release)
	return baV7tO7i5
}

func (baV7tO7i5 *P2XtkGf3) qpSVI0bOJG() {
	 /*line :1*/ baV7tO7i5.g6qmk4H.Store(true)
}

func (baV7tO7i5 *P2XtkGf3) abgQEYvfC9rq() bool {
	return  /*line :1*/ baV7tO7i5.g6qmk4H.Load()
}



type A5CIaIZeDjn struct {
	
	
	NX03GuKop6	string
	
	
	
	Ws3LdrwSmPa	[]string
	
	
	
	
	
	
	
	
	TXyhq0	[]*BvGocYcXx

	
	
	
	
	ZMj6ixN	*syscall.YxcIO6yPaPK
}




type J0zkNT interface {
	String() string
	Signal()	
}


func UfFnMrrU9WZ8() int	{ return  /*line :1*/ syscall.HLpaMN9pU4() }


func Z2W18nARKYH() int	{ return  /*line :1*/ syscall.IAWmZ0vBKwfC() }










func TVEypJmLUgVe(gAIGP2p int) (*P2XtkGf3, error) {
	return  /*line :1*/ ya248LtAOt(gAIGP2p)
}














func J893_zaL(jGBs03 string, aXiZTmqaf []string, yEeSjJ *A5CIaIZeDjn) (*P2XtkGf3, error) {
	 /*line :1*/ testlog.NT3Zg6plraKK(jGBs03)
	return  /*line :1*/ rNvAmv3ii(jGBs03, aXiZTmqaf, yEeSjJ)
}




func (baV7tO7i5 *P2XtkGf3) Release() error {
	return  /*line :1*/ baV7tO7i5.qq2HndsNY()
}




func (baV7tO7i5 *P2XtkGf3) Kill() error {
	return  /*line :1*/ baV7tO7i5.xbit3b7ALF()
}






func (baV7tO7i5 *P2XtkGf3) Wait() (*LcYXxKH, error) {
	return  /*line :1*/ baV7tO7i5.xbyZguBq()
}



func (baV7tO7i5 *P2XtkGf3) Signal(twfNTarXO J0zkNT) error {
	return  /*line :1*/ baV7tO7i5.mrvahvDrXb8(twfNTarXO)
}


func (baV7tO7i5 *LcYXxKH) UserTime() time.GKMwTGxQa0S {
	return  /*line :1*/ baV7tO7i5.ivGDNoOUnU()
}


func (baV7tO7i5 *LcYXxKH) SystemTime() time.GKMwTGxQa0S {
	return  /*line :1*/ baV7tO7i5.baMIJzSnuXFy()
}




func (baV7tO7i5 *LcYXxKH) Exited() bool {
	return  /*line :1*/ baV7tO7i5.uOgff09Wu1iI()
}



func (baV7tO7i5 *LcYXxKH) Success() bool {
	return  /*line :1*/ baV7tO7i5.bQvzsO()
}




func (baV7tO7i5 *LcYXxKH) Sys() any {
	return  /*line :1*/ baV7tO7i5.cx130qAaxcYa()
}






func (baV7tO7i5 *LcYXxKH) SysUsage() any {
	return  /*line :1*/ baV7tO7i5.xQWuBvE7A15o()
}
