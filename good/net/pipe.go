//line :1
package gF9mZs2

import (
	io "xI9ai2djaJ2"
	os "hPMrClpC"
	sync "sync"
	time "fRAfQd_"
)


type t303apnMaw struct {
	sipG034QB	sync.DIRWe11KYlYa	
	uTuMaHW8	*time.Bkz_9qnV
	any9NRUMOe	chan struct{}	
}

func oZneez_8() t303apnMaw {
	return t303apnMaw{any9NRUMOe:  /*line :1*/ make(chan struct{})}
}







func (ica44Q *t303apnMaw) hVlQmKILF2(aeaqWzxJu time.G4KDkI) {
	 /*line :1*/ ica44Q.sipG034QB.Lock()
	defer  /*line :1*/ ica44Q.sipG034QB.Unlock()

	if ica44Q.uTuMaHW8 != nil && ! /*line :1*/ ica44Q.uTuMaHW8.Stop() {
		<-ica44Q.any9NRUMOe
	}
	ica44Q.uTuMaHW8 = nil

	aHtIA2 :=  /*line :1*/ ogtFcD7gr0(ica44Q.any9NRUMOe)
	if  /*line :1*/ aeaqWzxJu.IsZero() {
		if aHtIA2 {
			ica44Q.any9NRUMOe =  /*line :1*/ make(chan struct{})
		}
		return
	}

	if fwCVIdfZ :=  /*line :1*/ time.Qr_kn6Dzra(aeaqWzxJu); fwCVIdfZ > 0 {
		if aHtIA2 {
			ica44Q.any9NRUMOe =  /*line :1*/ make(chan struct{})
		}
		ica44Q.uTuMaHW8 =  /*line :1*/ time.RzWe24uxMiv(fwCVIdfZ, func() {
			 /*line :1*/ close(ica44Q.any9NRUMOe)
		})
		return
	}

	if !aHtIA2 {
		 /*line :1*/ close(ica44Q.any9NRUMOe)
	}
}


func (ica44Q *t303apnMaw) oPIPVG() chan struct{} {
	 /*line :1*/ ica44Q.sipG034QB.Lock()
	defer  /*line :1*/ ica44Q.sipG034QB.Unlock()
	return ica44Q.any9NRUMOe
}

func ogtFcD7gr0(hl8w2lFX <-chan struct{}) bool {
	select {
	case <-hl8w2lFX:
		return true
	default:
		return false
	}
}

type oNWk55Wsw_ struct{}

func (oNWk55Wsw_) Network() string	{ return "pipe" }
func (oNWk55Wsw_) String() string	{ return "pipe" }

type z6x4gS struct {
	svaa4qoLfcB	sync.DIRWe11KYlYa	

	
	
	vDsT7m	<-chan []byte
	iu71IFhZei	chan<- int

	
	
	umwI5p	chan<- []byte
	ipV8V5d7GdOg	<-chan int

	ufx4HiN	sync.LhBfwn6wa1x	
	nmmZwG1	chan struct{}
	pQo1k6	<-chan struct{}

	n8yZHUbCL	t303apnMaw
	fcoMjm1KF3aC	t303apnMaw
}






func NXsa52o() (UJYB3aCQg, UJYB3aCQg) {
	yxZ2Lcq :=  /*line :1*/ make(chan []byte)
	mwJg1OVVDkuk :=  /*line :1*/ make(chan []byte)
	e1l7U0pY3A :=  /*line :1*/ make(chan int)
	_jVWNDp8sSj :=  /*line :1*/ make(chan int)
	bElQ1bH3wW5 :=  /*line :1*/ make(chan struct{})
	d0ms8Y :=  /*line :1*/ make(chan struct{})

	loPJgB_e := &z6x4gS{
		vDsT7m:	yxZ2Lcq, iu71IFhZei: e1l7U0pY3A,
		umwI5p:	mwJg1OVVDkuk, ipV8V5d7GdOg: _jVWNDp8sSj,
		nmmZwG1:	bElQ1bH3wW5, pQo1k6: d0ms8Y,
		n8yZHUbCL:	 /*line :1*/ oZneez_8(),
		fcoMjm1KF3aC:	 /*line :1*/ oZneez_8(),
	}
	betRxypiuO := &z6x4gS{
		vDsT7m:	mwJg1OVVDkuk, iu71IFhZei: _jVWNDp8sSj,
		umwI5p:	yxZ2Lcq, ipV8V5d7GdOg: e1l7U0pY3A,
		nmmZwG1:	d0ms8Y, pQo1k6: bElQ1bH3wW5,
		n8yZHUbCL:	 /*line :1*/ oZneez_8(),
		fcoMjm1KF3aC:	 /*line :1*/ oZneez_8(),
	}
	return loPJgB_e, betRxypiuO
}

func (*z6x4gS) LocalAddr() EqbMXsaU1lE	{ return oNWk55Wsw_{} }
func (*z6x4gS) RemoteAddr() EqbMXsaU1lE	{ return oNWk55Wsw_{} }

func (fMPVz2iGiyq *z6x4gS) Read(fIadEXIim6l []byte) (int, error) {
	doauF8Y, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.pfJLmwQdBx39(fIadEXIim6l)
	if h_ljk48Bm != nil && h_ljk48Bm != io.K5Sqco && h_ljk48Bm != io.DeqsCGTY {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: "pipe", XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}

func (fMPVz2iGiyq *z6x4gS) pfJLmwQdBx39(fIadEXIim6l []byte) (doauF8Y int, h_ljk48Bm error) {
	switch {
	case  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.nmmZwG1):
		return 0, io.DeqsCGTY
	case  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.pQo1k6):
		return 0, io.K5Sqco
	case  /*line :1*/ ogtFcD7gr0( /*line :1*/ fMPVz2iGiyq.n8yZHUbCL.oPIPVG()):
		return 0, os.DZEN6c
	}

	select {
	case w5UdWk5lsyE := <-fMPVz2iGiyq.vDsT7m:
		h3w5VWhxC7U :=  /*line :1*/ copy(fIadEXIim6l, w5UdWk5lsyE)
		fMPVz2iGiyq.iu71IFhZei <- h3w5VWhxC7U
		return h3w5VWhxC7U, nil
	case <-fMPVz2iGiyq.nmmZwG1:
		return 0, io.DeqsCGTY
	case <-fMPVz2iGiyq.pQo1k6:
		return 0, io.K5Sqco
	case <- /*line :1*/ fMPVz2iGiyq.n8yZHUbCL.oPIPVG():
		return 0, os.DZEN6c
	}
}

func (fMPVz2iGiyq *z6x4gS) Write(fIadEXIim6l []byte) (int, error) {
	doauF8Y, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.qba9dIp5a4oE(fIadEXIim6l)
	if h_ljk48Bm != nil && h_ljk48Bm != io.DeqsCGTY {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: "pipe", XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}

func (fMPVz2iGiyq *z6x4gS) qba9dIp5a4oE(fIadEXIim6l []byte) (doauF8Y int, h_ljk48Bm error) {
	switch {
	case  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.nmmZwG1):
		return 0, io.DeqsCGTY
	case  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.pQo1k6):
		return 0, io.DeqsCGTY
	case  /*line :1*/ ogtFcD7gr0( /*line :1*/ fMPVz2iGiyq.fcoMjm1KF3aC.oPIPVG()):
		return 0, os.DZEN6c
	}

	 /*line :1*/ fMPVz2iGiyq.svaa4qoLfcB.Lock()
	defer  /*line :1*/ fMPVz2iGiyq.svaa4qoLfcB.Unlock()
	for haJyZOtk := true; haJyZOtk ||  /*line :1*/ len(fIadEXIim6l) > 0; haJyZOtk = false {
		select {
		case fMPVz2iGiyq.umwI5p <- fIadEXIim6l:
			bFQSbJ := <-fMPVz2iGiyq.ipV8V5d7GdOg
			fIadEXIim6l = fIadEXIim6l[bFQSbJ:]
			doauF8Y += bFQSbJ
		case <-fMPVz2iGiyq.nmmZwG1:
			return doauF8Y, io.DeqsCGTY
		case <-fMPVz2iGiyq.pQo1k6:
			return doauF8Y, io.DeqsCGTY
		case <- /*line :1*/ fMPVz2iGiyq.fcoMjm1KF3aC.oPIPVG():
			return doauF8Y, os.DZEN6c
		}
	}
	return doauF8Y, nil
}

func (fMPVz2iGiyq *z6x4gS) SetDeadline(aeaqWzxJu time.G4KDkI) error {
	if  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.nmmZwG1) ||  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.pQo1k6) {
		return io.DeqsCGTY
	}
	 /*line :1*/ fMPVz2iGiyq.n8yZHUbCL.hVlQmKILF2(aeaqWzxJu)
	 /*line :1*/ fMPVz2iGiyq.fcoMjm1KF3aC.hVlQmKILF2(aeaqWzxJu)
	return nil
}

func (fMPVz2iGiyq *z6x4gS) SetReadDeadline(aeaqWzxJu time.G4KDkI) error {
	if  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.nmmZwG1) ||  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.pQo1k6) {
		return io.DeqsCGTY
	}
	 /*line :1*/ fMPVz2iGiyq.n8yZHUbCL.hVlQmKILF2(aeaqWzxJu)
	return nil
}

func (fMPVz2iGiyq *z6x4gS) SetWriteDeadline(aeaqWzxJu time.G4KDkI) error {
	if  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.nmmZwG1) ||  /*line :1*/ ogtFcD7gr0(fMPVz2iGiyq.pQo1k6) {
		return io.DeqsCGTY
	}
	 /*line :1*/ fMPVz2iGiyq.fcoMjm1KF3aC.hVlQmKILF2(aeaqWzxJu)
	return nil
}

func (fMPVz2iGiyq *z6x4gS) Close() error {
	 /*line :1*/ fMPVz2iGiyq.ufx4HiN.Do(func() {  /*line :1*/ close(fMPVz2iGiyq.nmmZwG1) })
	return nil
}
