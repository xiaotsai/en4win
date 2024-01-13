//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	itoa "JkjxVS"
	netip "iPdCam_KQOXj"
	syscall "bUKeamF"
)


type ZaffanpNx4 struct {
	ERvaNiiVmytR	GNraIvYhB
	EGf1xSAbsOU5	int
	RgTF6a1Xmo	string	
}






func (uI7LZDHm6 *ZaffanpNx4) AddrPort() netip.YTqTf_4VC {
	if uI7LZDHm6 == nil {
		return netip.YTqTf_4VC{}
	}
	dP8BEL5wfT_, _ :=  /*line :1*/ netip.OWcivL5ykj9(uI7LZDHm6.ERvaNiiVmytR)
	dP8BEL5wfT_ =  /*line :1*/ dP8BEL5wfT_.WithZone(uI7LZDHm6.RgTF6a1Xmo)
	return  /*line :1*/ netip.Iqr8_1pEDmw0(dP8BEL5wfT_,  /*line :1*/ uint16(uI7LZDHm6.EGf1xSAbsOU5))
}


func (uI7LZDHm6 *ZaffanpNx4) Network() string	{ return "udp" }

func (uI7LZDHm6 *ZaffanpNx4) String() string {
	if uI7LZDHm6 == nil {
		return "<nil>"
	}
	kQ8_UEhxU :=  /*line :1*/ pXjCfqAUUC(uI7LZDHm6.ERvaNiiVmytR)
	if uI7LZDHm6.RgTF6a1Xmo != "" {
		return  /*line :1*/ Vd5lcUxG(kQ8_UEhxU+"%"+uI7LZDHm6.RgTF6a1Xmo,  /*line :1*/ itoa.V25ba9G5(uI7LZDHm6.EGf1xSAbsOU5))
	}
	return  /*line :1*/ Vd5lcUxG(kQ8_UEhxU,  /*line :1*/ itoa.V25ba9G5(uI7LZDHm6.EGf1xSAbsOU5))
}

func (uI7LZDHm6 *ZaffanpNx4) tmKV_f7MKX() bool {
	if uI7LZDHm6 == nil || uI7LZDHm6.ERvaNiiVmytR == nil {
		return true
	}
	return  /*line :1*/ uI7LZDHm6.ERvaNiiVmytR.IsUnspecified()
}

func (uI7LZDHm6 *ZaffanpNx4) gFePPj34XJ0() EqbMXsaU1lE {
	if uI7LZDHm6 == nil {
		return nil
	}
	return uI7LZDHm6
}
















func B7QLF2EB20hj(vsbiWLn7reX, fwV_ln2dl string) (*ZaffanpNx4, error) {
	switch vsbiWLn7reX {
	case "udp", "udp4", "udp6":
	case "":
		vsbiWLn7reX = "udp"
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
	}
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ Ic4wtIY.hCHWGoQGe2q( /*line :1*/ context.GEcgQP5fzA(), vsbiWLn7reX, fwV_ln2dl)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ md4QSRkO.dyGWFsnhi4(vsbiWLn7reX, fwV_ln2dl).(*ZaffanpNx4), nil
}




func JjWbxscDw_h(qxwkC3VxG0 netip.YTqTf_4VC) *ZaffanpNx4 {
	return &ZaffanpNx4{
		ERvaNiiVmytR:	 /*line :1*/ qxwkC3VxG0.Addr().AsSlice(),
		RgTF6a1Xmo:	 /*line :1*/ qxwkC3VxG0.Addr().Zone(),
		EGf1xSAbsOU5:	 /*line :1*/ int( /*line :1*/ qxwkC3VxG0.Port()),
	}
}


type jVSK9IB_ struct {
	netip.YTqTf_4VC
}

func (jVSK9IB_) Network() string	{ return "udp" }



type JdInDLIal struct {
	xacP7km5pe
}



func (hl8w2lFX *JdInDLIal) SyscallConn() (syscall.CVnV1i, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	return  /*line :1*/ zQf8had4tml(hl8w2lFX.sXhoTdSE8wtb)
}


func (hl8w2lFX *JdInDLIal) ReadFromUDP(fIadEXIim6l []byte) (doauF8Y int, qxwkC3VxG0 *ZaffanpNx4, h_ljk48Bm error) {

	return  /*line :1*/ hl8w2lFX.gbcJ2VI(fIadEXIim6l, &ZaffanpNx4{})
}


func (hl8w2lFX *JdInDLIal) gbcJ2VI(fIadEXIim6l []byte, qxwkC3VxG0 *ZaffanpNx4) (int, *ZaffanpNx4, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, nil, syscall.EINVAL
	}
	doauF8Y, qxwkC3VxG0, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.mA1mFe4YP(fIadEXIim6l, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}


func (hl8w2lFX *JdInDLIal) ReadFrom(fIadEXIim6l []byte) (int, EqbMXsaU1lE, error) {
	doauF8Y, qxwkC3VxG0, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.gbcJ2VI(fIadEXIim6l, &ZaffanpNx4{})
	if qxwkC3VxG0 == nil {

		return doauF8Y, nil, h_ljk48Bm
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}






func (hl8w2lFX *JdInDLIal) ReadFromUDPAddrPort(fIadEXIim6l []byte) (doauF8Y int, qxwkC3VxG0 netip.YTqTf_4VC, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, netip.YTqTf_4VC{}, syscall.EINVAL
	}
	doauF8Y, qxwkC3VxG0, h_ljk48Bm =  /*line :1*/ hl8w2lFX.bAIcoRDfR0(fIadEXIim6l)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}








func (hl8w2lFX *JdInDLIal) ReadMsgUDP(fIadEXIim6l, cFiVO46YAy8g []byte) (doauF8Y, ytEzh3580y, dBNjCxLYqs int, qxwkC3VxG0 *ZaffanpNx4, h_ljk48Bm error) {
	var dqpqJJfc netip.YTqTf_4VC
	doauF8Y, ytEzh3580y, dBNjCxLYqs, dqpqJJfc, h_ljk48Bm =  /*line :1*/ hl8w2lFX.ReadMsgUDPAddrPort(fIadEXIim6l, cFiVO46YAy8g)
	if  /*line :1*/ dqpqJJfc.IsValid() {
		qxwkC3VxG0 =  /*line :1*/ JjWbxscDw_h(dqpqJJfc)
	}
	return
}


func (hl8w2lFX *JdInDLIal) ReadMsgUDPAddrPort(fIadEXIim6l, cFiVO46YAy8g []byte) (doauF8Y, ytEzh3580y, dBNjCxLYqs int, qxwkC3VxG0 netip.YTqTf_4VC, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, 0, 0, netip.YTqTf_4VC{}, syscall.EINVAL
	}
	doauF8Y, ytEzh3580y, dBNjCxLYqs, qxwkC3VxG0, h_ljk48Bm =  /*line :1*/ hl8w2lFX.vTzqIG(fIadEXIim6l, cFiVO46YAy8g)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}


func (hl8w2lFX *JdInDLIal) WriteToUDP(fIadEXIim6l []byte, qxwkC3VxG0 *ZaffanpNx4) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.omQ684zc(fIadEXIim6l, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ qxwkC3VxG0.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}


func (hl8w2lFX *JdInDLIal) WriteToUDPAddrPort(fIadEXIim6l []byte, qxwkC3VxG0 netip.YTqTf_4VC) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.goun3aKzAu(fIadEXIim6l, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: jVSK9IB_{qxwkC3VxG0}, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}


func (hl8w2lFX *JdInDLIal) WriteTo(fIadEXIim6l []byte, qxwkC3VxG0 EqbMXsaU1lE) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	uI7LZDHm6, d30HIwxht := qxwkC3VxG0.(*ZaffanpNx4)
	if !d30HIwxht {
		return 0, &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: qxwkC3VxG0, XkWH4CYwNmvD: syscall.EINVAL}
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.omQ684zc(fIadEXIim6l, uI7LZDHm6)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ uI7LZDHm6.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}









func (hl8w2lFX *JdInDLIal) WriteMsgUDP(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 *ZaffanpNx4) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, 0, syscall.EINVAL
	}
	doauF8Y, ytEzh3580y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.fh4B50S(fIadEXIim6l, cFiVO46YAy8g, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ qxwkC3VxG0.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}


func (hl8w2lFX *JdInDLIal) WriteMsgUDPAddrPort(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 netip.YTqTf_4VC) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, 0, syscall.EINVAL
	}
	doauF8Y, ytEzh3580y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.s3pDutKjt9Qk(fIadEXIim6l, cFiVO46YAy8g, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: jVSK9IB_{qxwkC3VxG0}, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}

func mgV5c9(vyaiiPXm6k_W *jmJczkl1) *JdInDLIal	{ return &JdInDLIal{xacP7km5pe{vyaiiPXm6k_W}} }








func GuS3JRd2o(vsbiWLn7reX string, bFNyUpAx, wcA44hhD *ZaffanpNx4) (*JdInDLIal, error) {
	switch vsbiWLn7reX {
	case "udp", "udp4", "udp6":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp:  /*line :1*/ wcA44hhD.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	if wcA44hhD == nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp: nil, XkWH4CYwNmvD: tccDv0yXpU}
	}
	dKLkhP := &vTALwD6Cw{cyZLILkT: vsbiWLn7reX, r0U09j:  /*line :1*/ wcA44hhD.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ dKLkhP.ubPg4ZjrNYRu( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx, wcA44hhD)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp:  /*line :1*/ wcA44hhD.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}










func N517ymwT(vsbiWLn7reX string, bFNyUpAx *ZaffanpNx4) (*JdInDLIal, error) {
	switch vsbiWLn7reX {
	case "udp", "udp4", "udp6":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	if bFNyUpAx == nil {
		bFNyUpAx = &ZaffanpNx4{}
	}
	bw2hLt := &tYwm0DwfSTEv{aNAwQMxeTY: vsbiWLn7reX, rxZTzA:  /*line :1*/ bFNyUpAx.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ bw2hLt.eg0RUV( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}





















func MUF34odS5(vsbiWLn7reX string, dvFmDYmQz *U_Uc9la, oqJawTas3 *ZaffanpNx4) (*JdInDLIal, error) {
	switch vsbiWLn7reX {
	case "udp", "udp4", "udp6":
	default:
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ oqJawTas3.gFePPj34XJ0(), XkWH4CYwNmvD:  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)}
	}
	if oqJawTas3 == nil || oqJawTas3.ERvaNiiVmytR == nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ oqJawTas3.gFePPj34XJ0(), XkWH4CYwNmvD: tccDv0yXpU}
	}
	bw2hLt := &tYwm0DwfSTEv{aNAwQMxeTY: vsbiWLn7reX, rxZTzA:  /*line :1*/ oqJawTas3.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ bw2hLt.zdApnUghKfl( /*line :1*/ context.GEcgQP5fzA(), dvFmDYmQz, oqJawTas3)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ oqJawTas3.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}
