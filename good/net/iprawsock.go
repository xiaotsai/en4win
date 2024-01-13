//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	syscall "bUKeamF"
)


type FZJphYv9hV struct {
	GdouSNq80bRw	GNraIvYhB
	Cyg5M2vYIby	string	
}


func (uI7LZDHm6 *FZJphYv9hV) Network() string	{ return "ip" }

func (uI7LZDHm6 *FZJphYv9hV) String() string {
	if uI7LZDHm6 == nil {
		return "<nil>"
	}
	kQ8_UEhxU :=  /*line :1*/ pXjCfqAUUC(uI7LZDHm6.GdouSNq80bRw)
	if uI7LZDHm6.Cyg5M2vYIby != "" {
		return kQ8_UEhxU + "%" + uI7LZDHm6.Cyg5M2vYIby
	}
	return kQ8_UEhxU
}

func (uI7LZDHm6 *FZJphYv9hV) tmKV_f7MKX() bool {
	if uI7LZDHm6 == nil || uI7LZDHm6.GdouSNq80bRw == nil {
		return true
	}
	return  /*line :1*/ uI7LZDHm6.GdouSNq80bRw.IsUnspecified()
}

func (uI7LZDHm6 *FZJphYv9hV) gFePPj34XJ0() EqbMXsaU1lE {
	if uI7LZDHm6 == nil {
		return nil
	}
	return uI7LZDHm6
}














func L_TzsKD261w4(vsbiWLn7reX, fwV_ln2dl string) (*FZJphYv9hV, error) {
	if vsbiWLn7reX == "" {
		vsbiWLn7reX = "ip"
	}
	mQBJul9fW, _, h_ljk48Bm :=  /*line :1*/ nihn79f8iY( /*line :1*/ context.GEcgQP5fzA(), vsbiWLn7reX, false)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	switch mQBJul9fW {
	case "ip", "ip4", "ip6":
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
	}
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ Ic4wtIY.hCHWGoQGe2q( /*line :1*/ context.GEcgQP5fzA(), mQBJul9fW, fwV_ln2dl)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ md4QSRkO.dyGWFsnhi4(vsbiWLn7reX, fwV_ln2dl).(*FZJphYv9hV), nil
}



type HTBkPb struct {
	xacP7km5pe
}



func (hl8w2lFX *HTBkPb) SyscallConn() (syscall.CVnV1i, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return nil, syscall.EINVAL
	}
	return  /*line :1*/ zQf8had4tml(hl8w2lFX.sXhoTdSE8wtb)
}


func (hl8w2lFX *HTBkPb) ReadFromIP(fIadEXIim6l []byte) (int, *FZJphYv9hV, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, nil, syscall.EINVAL
	}
	doauF8Y, qxwkC3VxG0, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.mA1mFe4YP(fIadEXIim6l)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}


func (hl8w2lFX *HTBkPb) ReadFrom(fIadEXIim6l []byte) (int, EqbMXsaU1lE, error) {
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








func (hl8w2lFX *HTBkPb) ReadMsgIP(fIadEXIim6l, cFiVO46YAy8g []byte) (doauF8Y, ytEzh3580y, dBNjCxLYqs int, qxwkC3VxG0 *FZJphYv9hV, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, 0, 0, nil, syscall.EINVAL
	}
	doauF8Y, ytEzh3580y, dBNjCxLYqs, qxwkC3VxG0, h_ljk48Bm =  /*line :1*/ hl8w2lFX.vTzqIG(fIadEXIim6l, cFiVO46YAy8g)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}


func (hl8w2lFX *HTBkPb) WriteToIP(fIadEXIim6l []byte, qxwkC3VxG0 *FZJphYv9hV) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.omQ684zc(fIadEXIim6l, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ qxwkC3VxG0.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}


func (hl8w2lFX *HTBkPb) WriteTo(fIadEXIim6l []byte, qxwkC3VxG0 EqbMXsaU1lE) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	uI7LZDHm6, d30HIwxht := qxwkC3VxG0.(*FZJphYv9hV)
	if !d30HIwxht {
		return 0, &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: qxwkC3VxG0, XkWH4CYwNmvD: syscall.EINVAL}
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.omQ684zc(fIadEXIim6l, uI7LZDHm6)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ uI7LZDHm6.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}







func (hl8w2lFX *HTBkPb) WriteMsgIP(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 *FZJphYv9hV) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, 0, syscall.EINVAL
	}
	doauF8Y, ytEzh3580y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.fh4B50S(fIadEXIim6l, cFiVO46YAy8g, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp:  /*line :1*/ qxwkC3VxG0.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}

func c8Eh3kKno(vyaiiPXm6k_W *jmJczkl1) *HTBkPb	{ return &HTBkPb{xacP7km5pe{vyaiiPXm6k_W}} }








func VNkj0DP(vsbiWLn7reX string, bFNyUpAx, wcA44hhD *FZJphYv9hV) (*HTBkPb, error) {
	if wcA44hhD == nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp: nil, XkWH4CYwNmvD: tccDv0yXpU}
	}
	dKLkhP := &vTALwD6Cw{cyZLILkT: vsbiWLn7reX, r0U09j:  /*line :1*/ wcA44hhD.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ dKLkhP.xQAA5av( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx, wcA44hhD)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), JkQ9XFJbp:  /*line :1*/ wcA44hhD.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}








func ENkItDFTG(vsbiWLn7reX string, bFNyUpAx *FZJphYv9hV) (*HTBkPb, error) {
	if bFNyUpAx == nil {
		bFNyUpAx = &FZJphYv9hV{}
	}
	bw2hLt := &tYwm0DwfSTEv{aNAwQMxeTY: vsbiWLn7reX, rxZTzA:  /*line :1*/ bFNyUpAx.String()}
	hl8w2lFX, h_ljk48Bm :=  /*line :1*/ bw2hLt.hq2iTKku3u( /*line :1*/ context.GEcgQP5fzA(), bFNyUpAx)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp:  /*line :1*/ bFNyUpAx.gFePPj34XJ0(), XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}
