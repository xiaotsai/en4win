//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	itoa "JkjxVS"
	io "xI9ai2djaJ2"
	netip "iPdCam_KQOXj"
	os "hPMrClpC"
	syscall "bUKeamF"
	time "fRAfQd_"
)


type HiOzqOVjN1P struct {
	ERvaNiiVmytR	GNraIvYhB
	EGf1xSAbsOU5	int
	RgTF6a1Xmo	string	
}






func (uI7LZDHm6 *HiOzqOVjN1P) AddrPort() netip.YTqTf_4VC {
	if uI7LZDHm6 == nil {
		return netip.YTqTf_4VC{}
	}
	dP8BEL5wfT_, _ :=  /*line :1*/ netip.OWcivL5ykj9(uI7LZDHm6.ERvaNiiVmytR)
	dP8BEL5wfT_ =  /*line :1*/ dP8BEL5wfT_.WithZone(uI7LZDHm6.RgTF6a1Xmo)
	return  /*line :1*/ netip.Iqr8_1pEDmw0(dP8BEL5wfT_,  /*line :1*/ uint16(uI7LZDHm6.EGf1xSAbsOU5))
}


func (uI7LZDHm6 *HiOzqOVjN1P) Network() string	{ return "tcp" }

func (uI7LZDHm6 *HiOzqOVjN1P) String() string {
	if uI7LZDHm6 == nil {
		return "<nil>"
	}
	kQ8_UEhxU :=  /*line :1*/ pXjCfqAUUC(uI7LZDHm6.ERvaNiiVmytR)
	if uI7LZDHm6.RgTF6a1Xmo != "" {
		return  /*line :1*/ Vd5lcUxG(kQ8_UEhxU+"%"+uI7LZDHm6.RgTF6a1Xmo,  /*line :1*/ itoa.V25ba9G5(uI7LZDHm6.EGf1xSAbsOU5))
	}
	return  /*line :1*/ Vd5lcUxG(kQ8_UEhxU,  /*line :1*/ itoa.V25ba9G5(uI7LZDHm6.EGf1xSAbsOU5))
}

func (uI7LZDHm6 *HiOzqOVjN1P) tmKV_f7MKX() bool {
	if uI7LZDHm6 == nil || uI7LZDHm6.ERvaNiiVmytR == nil {
		return true
	}
	return  /*line :1*/ uI7LZDHm6.ERvaNiiVmytR.IsUnspecified()
}

func (uI7LZDHm6 *HiOzqOVjN1P) gFePPj34XJ0() EqbMXsaU1lE {
	if uI7LZDHm6 == nil {
		return nil
	}
	return uI7LZDHm6
}
















func DUjk6ab0K(vsbiWLn7reX, fwV_ln2dl string) (*HiOzqOVjN1P, error) {
	switch vsbiWLn7reX {
	case "tcp", "tcp4", "tcp6":
	case "":
		vsbiWLn7reX = "tcp"
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
	}
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ Ic4wtIY.hCHWGoQGe2q( /*line :1*/ context.GEcgQP5fzA(), vsbiWLn7reX, fwV_ln2dl)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ md4QSRkO.dyGWFsnhi4(vsbiWLn7reX, fwV_ln2dl).(*HiOzqOVjN1P), nil
}




func EGF0vgrzG(qxwkC3VxG0 netip.YTqTf_4VC) *HiOzqOVjN1P {
	return &HiOzqOVjN1P{
		ERvaNiiVmytR:	 /*line :1*/ qxwkC3VxG0.Addr().AsSlice(),
		RgTF6a1Xmo:	 /*line :1*/ qxwkC3VxG0.Addr().Zone(),
		EGf1xSAbsOU5:	 /*line :1*/ int( /*line :1*/ qxwkC3VxG0.Port()),
	}
}



type Q2rqhDKoOIxM struct {
	xacP7km5pe
}



func (hl8w2lFX *Q2rqhDKoOIxM) SyscallConn() (syscall.CVnV1i, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	return  /*line :1*/ zQf8had4tml(hl8w2lFX.sXhoTdSE8wtb)
}


func (hl8w2lFX *Q2rqhDKoOIxM) ReadFrom(yxhs4Z io.YJ04Fau) (int64, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.mA1mFe4YP(yxhs4Z)
	if h_ljk48Bm != nil && h_ljk48Bm != io.K5Sqco {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "readfrom", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}



func (hl8w2lFX *Q2rqhDKoOIxM) CloseRead() error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.ea_e1ZnpFCQF(); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "close", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}



func (hl8w2lFX *Q2rqhDKoOIxM) CloseWrite() error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.oCRE0sQeNq(); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "close", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}















func (hl8w2lFX *Q2rqhDKoOIxM) SetLinger(kG6j6Iy3kd int) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hI6GYQFrkn(hl8w2lFX.sXhoTdSE8wtb, kG6j6Iy3kd); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}



func (hl8w2lFX *Q2rqhDKoOIxM) SetKeepAlive(hOXZGAT bool) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ umO09W(hl8w2lFX.sXhoTdSE8wtb, hOXZGAT); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}


func (hl8w2lFX *Q2rqhDKoOIxM) SetKeepAlivePeriod(ica44Q time.GKMwTGxQa0S) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ ln6KdLlW(hl8w2lFX.sXhoTdSE8wtb, ica44Q); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}





func (hl8w2lFX *Q2rqhDKoOIxM) SetNoDelay(aZNIDsV1NcU bool) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ aN3c7TO6l(hl8w2lFX.sXhoTdSE8wtb, aZNIDsV1NcU); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}










func (hl8w2lFX *Q2rqhDKoOIxM) MultipathTCP() (bool, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return false, syscall.EINVAL
	}
	return  /*line :1*/ r1dRlP3HTSS(hl8w2lFX.sXhoTdSE8wtb), nil
}

func xMgOZ10goOXP(vyaiiPXm6k_W *jmJczkl1, vDKTMW time.GKMwTGxQa0S, bjNSUq0b func(time.GKMwTGxQa0S)) *Q2rqhDKoOIxM {
	 /*line :1*/ aN3c7TO6l(vyaiiPXm6k_W, true)
	if vDKTMW == 0 {
		vDKTMW = defaultTCPKeepAlive
	}
	if vDKTMW > 0 {
		 /*line :1*/ umO09W(vyaiiPXm6k_W, true)
		 /*line :1*/ ln6KdLlW(vyaiiPXm6k_W, vDKTMW)
		if bjNSUq0b != nil {
			 /*line :1*/ bjNSUq0b(vDKTMW)
		}
	}
	return &Q2rqhDKoOIxM{xacP7km5pe{vyaiiPXm6k_W}}
}








func JEM7FVRnlcC(vsbiWLn7reX string, bFNyUpAx, wcA44hhD *HiOzqOVjN1P) (*Q2rqhDKoOIxM, error) {
	switch vsbiWLn7reX {
	case "tcp", "tcp4", "tcp6":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp:  /*line :1*/ wcA44hhD.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	if wcA44hhD == nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp: nil, XkWH4CYwNmvD: tccDv0yXpU}
	}
	dKLkhP := &vTALwD6Cw{cyZLILkT: vsbiWLn7reX, r0U09j:  /*line :1*/ wcA44hhD.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ dKLkhP.a9XNPFb( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx, wcA44hhD)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp:  /*line :1*/ wcA44hhD.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}



type S0u1keY struct {
	_Tj5wN	*jmJczkl1
	kuQlluF	MC4oQv
}






func (v3upVm *S0u1keY) SyscallConn() (syscall.CVnV1i, error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	return  /*line :1*/ aJmXTOyud(v3upVm._Tj5wN)
}



func (v3upVm *S0u1keY) AcceptTCP() (*Q2rqhDKoOIxM, error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ v3upVm.zjXrFvdT()
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "accept", CiQ5sBtmrVnf: v3upVm._Tj5wN.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm._Tj5wN.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}



func (v3upVm *S0u1keY) Accept() (UJYB3aCQg, error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ v3upVm.zjXrFvdT()
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "accept", CiQ5sBtmrVnf: v3upVm._Tj5wN.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm._Tj5wN.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}



func (v3upVm *S0u1keY) Close() error {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ v3upVm.tFIv2Ppbx0H(); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "close", CiQ5sBtmrVnf: v3upVm._Tj5wN.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm._Tj5wN.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}




func (v3upVm *S0u1keY) Addr() EqbMXsaU1lE	{ return v3upVm._Tj5wN.y9pjEf }



func (v3upVm *S0u1keY) SetDeadline(aeaqWzxJu time.G4KDkI) error {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ v3upVm._Tj5wN.q0t7ga7tBcV5.SetDeadline(aeaqWzxJu); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: v3upVm._Tj5wN.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm._Tj5wN.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}








func (v3upVm *S0u1keY) File() (t5q9DVOD9 *os.BvGocYcXx, h_ljk48Bm error) {
	if ! /*line :1*/ v3upVm.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	t5q9DVOD9, h_ljk48Bm =  /*line :1*/ v3upVm.cZsWMjIJ6()
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "file", CiQ5sBtmrVnf: v3upVm._Tj5wN.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: v3upVm._Tj5wN.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}










func JoJDaL(vsbiWLn7reX string, bFNyUpAx *HiOzqOVjN1P) (*S0u1keY, error) {
	switch vsbiWLn7reX {
	case "tcp", "tcp4", "tcp6":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	if bFNyUpAx == nil {
		bFNyUpAx = &HiOzqOVjN1P{}
	}
	bw2hLt := &tYwm0DwfSTEv{aNAwQMxeTY: vsbiWLn7reX, rxZTzA:  /*line :1*/ bFNyUpAx.String()}
	dhvebBa4tmD, h_ljk48Bm :=  /*line :1*/ bw2hLt.bwWMnaxr( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return dhvebBa4tmD, nil
}


func m6CmZV(ica44Q time.GKMwTGxQa0S, sbNWXQna time.GKMwTGxQa0S) time.GKMwTGxQa0S {
	return (ica44Q + sbNWXQna - 1) / sbNWXQna
}
