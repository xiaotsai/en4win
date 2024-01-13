//line :1
package xI9ai2djaJ2

import (
	errors "iAHoxjmM"
	sync "sync"
)


type cPXFQac4H struct {
	sync.DIRWe11KYlYa		
	bQtZvajsN	error
}

func (lCMbCoRJo1w9 *cPXFQac4H) Store(fjG4tVMGB error) {
	 /*line :1*/ lCMbCoRJo1w9.Lock()
	defer  /*line :1*/ lCMbCoRJo1w9.Unlock()
	if lCMbCoRJo1w9.bQtZvajsN != nil {
		return
	}
	lCMbCoRJo1w9.bQtZvajsN = fjG4tVMGB
}
func (lCMbCoRJo1w9 *cPXFQac4H) Load() error {
	 /*line :1*/ lCMbCoRJo1w9.Lock()
	defer  /*line :1*/ lCMbCoRJo1w9.Unlock()
	return lCMbCoRJo1w9.bQtZvajsN
}


var DeqsCGTY =  /*line :1*/ errors.Su6g6hRoi9X("io: read/write on closed pipe")


type mJ7XWtb struct {
	aFnRWKwCaQ	sync.DIRWe11KYlYa	
	kZPiMe5	chan []byte
	iilp9l0ACgm2	chan int

	lMZles	sync.LhBfwn6wa1x	
	xpJmZ52on	chan struct{}
	qQhrp5WT	cPXFQac4H
	suARUrHc	cPXFQac4H
}

func (x6wLQwPqT1n *mJ7XWtb) euZgo9YPiW(rG6_2l2i []byte) (f3MWP_v int, fjG4tVMGB error) {
	select {
	case <-x6wLQwPqT1n.xpJmZ52on:
		return 0,  /*line :1*/ x6wLQwPqT1n.a9iSjpQ()
	default:
	}

	select {
	case jQH8vb := <-x6wLQwPqT1n.kZPiMe5:
		fBxr6gYSESyl :=  /*line :1*/ copy(rG6_2l2i, jQH8vb)
		x6wLQwPqT1n.iilp9l0ACgm2 <- fBxr6gYSESyl
		return fBxr6gYSESyl, nil
	case <-x6wLQwPqT1n.xpJmZ52on:
		return 0,  /*line :1*/ x6wLQwPqT1n.a9iSjpQ()
	}
}

func (x6wLQwPqT1n *mJ7XWtb) i0yM77OagC0(fjG4tVMGB error) error {
	if fjG4tVMGB == nil {
		fjG4tVMGB = DeqsCGTY
	}
	 /*line :1*/ x6wLQwPqT1n.qQhrp5WT.Store(fjG4tVMGB)
	 /*line :1*/ x6wLQwPqT1n.lMZles.Do(func() {  /*line :1*/ close(x6wLQwPqT1n.xpJmZ52on) })
	return nil
}

func (x6wLQwPqT1n *mJ7XWtb) xsiG_J3LMAOQ(rG6_2l2i []byte) (f3MWP_v int, fjG4tVMGB error) {
	select {
	case <-x6wLQwPqT1n.xpJmZ52on:
		return 0,  /*line :1*/ x6wLQwPqT1n.v0lx_zavG()
	default:
		 /*line :1*/ x6wLQwPqT1n.aFnRWKwCaQ.Lock()
		defer  /*line :1*/ x6wLQwPqT1n.aFnRWKwCaQ.Unlock()
	}

	for eQCXTeNUiQf := true; eQCXTeNUiQf ||  /*line :1*/ len(rG6_2l2i) > 0; eQCXTeNUiQf = false {
		select {
		case x6wLQwPqT1n.kZPiMe5 <- rG6_2l2i:
			vpXSicX9M := <-x6wLQwPqT1n.iilp9l0ACgm2
			rG6_2l2i = rG6_2l2i[vpXSicX9M:]
			f3MWP_v += vpXSicX9M
		case <-x6wLQwPqT1n.xpJmZ52on:
			return f3MWP_v,  /*line :1*/ x6wLQwPqT1n.v0lx_zavG()
		}
	}
	return f3MWP_v, nil
}

func (x6wLQwPqT1n *mJ7XWtb) jFbmoLdvkTG(fjG4tVMGB error) error {
	if fjG4tVMGB == nil {
		fjG4tVMGB = K5Sqco
	}
	 /*line :1*/ x6wLQwPqT1n.suARUrHc.Store(fjG4tVMGB)
	 /*line :1*/ x6wLQwPqT1n.lMZles.Do(func() {  /*line :1*/ close(x6wLQwPqT1n.xpJmZ52on) })
	return nil
}


func (x6wLQwPqT1n *mJ7XWtb) a9iSjpQ() error {
	eVyNxJa4w :=  /*line :1*/ x6wLQwPqT1n.qQhrp5WT.Load()
	if zr2DQGP3 :=  /*line :1*/ x6wLQwPqT1n.suARUrHc.Load(); eVyNxJa4w == nil && zr2DQGP3 != nil {
		return zr2DQGP3
	}
	return DeqsCGTY
}


func (x6wLQwPqT1n *mJ7XWtb) v0lx_zavG() error {
	zr2DQGP3 :=  /*line :1*/ x6wLQwPqT1n.suARUrHc.Load()
	if eVyNxJa4w :=  /*line :1*/ x6wLQwPqT1n.qQhrp5WT.Load(); zr2DQGP3 == nil && eVyNxJa4w != nil {
		return eVyNxJa4w
	}
	return DeqsCGTY
}


type Dto9Cwk struct {
	indNJvt *mJ7XWtb
}






func (wZW51lrVDWS *Dto9Cwk) Read(zDMuUN []byte) (f3MWP_v int, fjG4tVMGB error) {
	return  /*line :1*/ wZW51lrVDWS.indNJvt.euZgo9YPiW(zDMuUN)
}



func (wZW51lrVDWS *Dto9Cwk) Close() error {
	return  /*line :1*/ wZW51lrVDWS.CloseWithError(nil)
}






func (wZW51lrVDWS *Dto9Cwk) CloseWithError(fjG4tVMGB error) error {
	return  /*line :1*/ wZW51lrVDWS.indNJvt.i0yM77OagC0(fjG4tVMGB)
}


type Z4T_0yYFX struct {
	indNJvt *mJ7XWtb
}






func (iO5t8eLj5Pd *Z4T_0yYFX) Write(zDMuUN []byte) (f3MWP_v int, fjG4tVMGB error) {
	return  /*line :1*/ iO5t8eLj5Pd.indNJvt.xsiG_J3LMAOQ(zDMuUN)
}



func (iO5t8eLj5Pd *Z4T_0yYFX) Close() error {
	return  /*line :1*/ iO5t8eLj5Pd.CloseWithError(nil)
}







func (iO5t8eLj5Pd *Z4T_0yYFX) CloseWithError(fjG4tVMGB error) error {
	return  /*line :1*/ iO5t8eLj5Pd.indNJvt.jFbmoLdvkTG(fjG4tVMGB)
}
















func Zxg6lE6XIfVo() (*Dto9Cwk, *Z4T_0yYFX) {
	x6wLQwPqT1n := &mJ7XWtb{
		kZPiMe5:	 /*line :1*/ make(chan []byte),
		iilp9l0ACgm2:	 /*line :1*/ make(chan int),
		xpJmZ52on:	 /*line :1*/ make(chan struct{}),
	}
	return &Dto9Cwk{x6wLQwPqT1n}, &Z4T_0yYFX{x6wLQwPqT1n}
}
