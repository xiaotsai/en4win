//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	os "hPMrClpC"
	sync "sync"
	syscall "bUKeamF"
	time "fRAfQd_"
)


type E5nPmTZQM struct {
	HIY_L1X2	string
	GrF5D_J	string
}



func (uI7LZDHm6 *E5nPmTZQM) Network() string {
	return uI7LZDHm6.GrF5D_J
}

func (uI7LZDHm6 *E5nPmTZQM) String() string {
	if uI7LZDHm6 == nil {
		return "<nil>"
	}
	return uI7LZDHm6.HIY_L1X2
}

func (uI7LZDHm6 *E5nPmTZQM) tmKV_f7MKX() bool {
	return uI7LZDHm6 == nil || uI7LZDHm6.HIY_L1X2 == ""
}

func (uI7LZDHm6 *E5nPmTZQM) gFePPj34XJ0() EqbMXsaU1lE {
	if uI7LZDHm6 == nil {
		return nil
	}
	return uI7LZDHm6
}







func E3DL71v(vsbiWLn7reX, fwV_ln2dl string) (*E5nPmTZQM, error) {
	switch vsbiWLn7reX {
	case "unix", "unixgram", "unixpacket":
		return &E5nPmTZQM{HIY_L1X2: fwV_ln2dl, GrF5D_J: vsbiWLn7reX}, nil
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
	}
}



type Xud0Zk struct {
	xacP7km5pe
}



func (hl8w2lFX *Xud0Zk) SyscallConn() (syscall.CVnV1i, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	return  /*line :1*/ zQf8had4tml(hl8w2lFX.sXhoTdSE8wtb)
}



func (hl8w2lFX *Xud0Zk) CloseRead() error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.ea_e1ZnpFCQF(); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "close", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}



func (hl8w2lFX *Xud0Zk) CloseWrite() error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.oCRE0sQeNq(); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "close", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}


func (hl8w2lFX *Xud0Zk) ReadFromUnix(fIadEXIim6l []byte) (int, *E5nPmTZQM, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, nil, syscall.EINVAL
	}
	doauF8Y, qxwkC3VxG0, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.mA1mFe4YP(fIadEXIim6l)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}


func (hl8w2lFX *Xud0Zk) ReadFrom(fIadEXIim6l []byte) (int, EqbMXsaU1lE, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, nil, syscall.EINVAL
	}
	doauF8Y, qxwkC3VxG0, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.mA1mFe4YP(fIadEXIim6l)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	if qxwkC3VxG0 == nil {
		return doauF8Y, nil, h_ljk48Bm
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}








func (hl8w2lFX *Xud0Zk) ReadMsgUnix(fIadEXIim6l, cFiVO46YAy8g []byte) (doauF8Y, ytEzh3580y, dBNjCxLYqs int, qxwkC3VxG0 *E5nPmTZQM, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, 0, 0, nil, syscall.EINVAL
	}
	doauF8Y, ytEzh3580y, dBNjCxLYqs, qxwkC3VxG0, h_ljk48Bm =  /*line :1*/ hl8w2lFX.vTzqIG(fIadEXIim6l, cFiVO46YAy8g)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}


func (hl8w2lFX *Xud0Zk) WriteToUnix(fIadEXIim6l []byte, qxwkC3VxG0 *E5nPmTZQM) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.omQ684zc(fIadEXIim6l, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ qxwkC3VxG0.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}


func (hl8w2lFX *Xud0Zk) WriteTo(fIadEXIim6l []byte, qxwkC3VxG0 EqbMXsaU1lE) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	uI7LZDHm6, d30HIwxht := qxwkC3VxG0.(*E5nPmTZQM)
	if !d30HIwxht {
		return 0, &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: qxwkC3VxG0, XkWH4CYwNmvD: syscall.EINVAL}
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.omQ684zc(fIadEXIim6l, uI7LZDHm6)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ uI7LZDHm6.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}







func (hl8w2lFX *Xud0Zk) WriteMsgUnix(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 *E5nPmTZQM) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, 0, syscall.EINVAL
	}
	doauF8Y, ytEzh3580y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.fh4B50S(fIadEXIim6l, cFiVO46YAy8g, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ qxwkC3VxG0.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}

func pUWAqmsC4s(vyaiiPXm6k_W *jmJczkl1) *Xud0Zk	{ return &Xud0Zk{xacP7km5pe{vyaiiPXm6k_W}} }







func IUxreizh5ys0(vsbiWLn7reX string, bFNyUpAx, wcA44hhD *E5nPmTZQM) (*Xud0Zk, error) {
	switch vsbiWLn7reX {
	case "unix", "unixgram", "unixpacket":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp:  /*line :1*/ wcA44hhD.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	dKLkhP := &vTALwD6Cw{cyZLILkT: vsbiWLn7reX, r0U09j:  /*line :1*/ wcA44hhD.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ dKLkhP.g5IHkyH( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx, wcA44hhD)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp:  /*line :1*/ wcA44hhD.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}




type X3_EsNk struct {
	wTPbhAE	*jmJczkl1
	ghYAeP	string
	u612R4LZ	bool
	v96dRno	sync.LhBfwn6wa1x
}

func (dhvebBa4tmD *X3_EsNk) d30HIwxht() bool	{ return dhvebBa4tmD != nil && dhvebBa4tmD.wTPbhAE != nil }






func (v3upVm *X3_EsNk) SyscallConn() (syscall.CVnV1i, error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	return  /*line :1*/ aJmXTOyud(v3upVm.wTPbhAE)
}



func (v3upVm *X3_EsNk) AcceptUnix() (*Xud0Zk, error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ v3upVm.zjXrFvdT()
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "accept", CiQ5sBtmrVnf: v3upVm.wTPbhAE.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm.wTPbhAE.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}



func (v3upVm *X3_EsNk) Accept() (UJYB3aCQg, error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ v3upVm.zjXrFvdT()
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "accept", CiQ5sBtmrVnf: v3upVm.wTPbhAE.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm.wTPbhAE.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}



func (v3upVm *X3_EsNk) Close() error {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ v3upVm.tFIv2Ppbx0H(); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "close", CiQ5sBtmrVnf: v3upVm.wTPbhAE.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm.wTPbhAE.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}




func (v3upVm *X3_EsNk) Addr() EqbMXsaU1lE	{ return v3upVm.wTPbhAE.y9pjEf }



func (v3upVm *X3_EsNk) SetDeadline(aeaqWzxJu time.G4KDkI) error {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ v3upVm.wTPbhAE.q0t7ga7tBcV5.SetDeadline(aeaqWzxJu); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: v3upVm.wTPbhAE.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm.wTPbhAE.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}








func (v3upVm *X3_EsNk) File() (t5q9DVOD9 *os.BvGocYcXx, h_ljk48Bm error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	t5q9DVOD9, h_ljk48Bm =  /*line :1*/ v3upVm.cZsWMjIJ6()
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "file", CiQ5sBtmrVnf: v3upVm.wTPbhAE.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm.wTPbhAE.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}




func J09pWjwGr1u(vsbiWLn7reX string, bFNyUpAx *E5nPmTZQM) (*X3_EsNk, error) {
	switch vsbiWLn7reX {
	case "unix", "unixpacket":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	if bFNyUpAx == nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD: tccDv0yXpU}
	}
	bw2hLt := &tYwm0DwfSTEv{aNAwQMxeTY: vsbiWLn7reX, rxZTzA:  /*line :1*/ bFNyUpAx.String()}
	dhvebBa4tmD, h_ljk48Bm :=  /*line :1*/ bw2hLt.caN1IhSmUV( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return dhvebBa4tmD, nil
}




func PkApm7IzK2Xg(vsbiWLn7reX string, bFNyUpAx *E5nPmTZQM) (*Xud0Zk, error) {
	switch vsbiWLn7reX {
	case "unixgram":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	if bFNyUpAx == nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: tccDv0yXpU}
	}
	bw2hLt := &tYwm0DwfSTEv{aNAwQMxeTY: vsbiWLn7reX, rxZTzA:  /*line :1*/ bFNyUpAx.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ bw2hLt.r5LehBZ4( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}
