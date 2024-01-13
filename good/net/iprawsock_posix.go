//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package gF9mZs2

import (
	context "fOYpb3F03EG"
	syscall "bUKeamF"
)

func llbJwAwVB(epTaNF0RJ8eN syscall.S4iroLVT4) EqbMXsaU1lE {
	switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
	case *syscall.Hx8lqxSkV:
		return &FZJphYv9hV{GdouSNq80bRw: epTaNF0RJ8eN.Q3zz2uH[0:]}
	case *syscall.HKTAy7_BSU2:
		return &FZJphYv9hV{GdouSNq80bRw: epTaNF0RJ8eN.Y4XxFr[0:], Cyg5M2vYIby:  /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(epTaNF0RJ8eN.AsutdJq2))}
	}
	return nil
}

func (uI7LZDHm6 *FZJphYv9hV) nqsgnnaf() int {
	if uI7LZDHm6 == nil ||  /*line :1*/ len(uI7LZDHm6.GdouSNq80bRw) <= IPv4len {
		return syscall.AF_INET
	}
	if  /*line :1*/ uI7LZDHm6.GdouSNq80bRw.To4() != nil {
		return syscall.AF_INET
	}
	return syscall.AF_INET6
}

func (uI7LZDHm6 *FZJphYv9hV) jIv0wLJwloQ(nqsgnnaf int) (syscall.S4iroLVT4, error) {
	if uI7LZDHm6 == nil {
		return nil, nil
	}
	return  /*line :1*/ ggQVAaURvc(nqsgnnaf, uI7LZDHm6.GdouSNq80bRw, 0, uI7LZDHm6.Cyg5M2vYIby)
}

func (uI7LZDHm6 *FZJphYv9hV) kAjqX7j(gF9mZs2 string) jIv0wLJwloQ {
	return &FZJphYv9hV{ /*line :1*/ gygXisQVD(gF9mZs2), uI7LZDHm6.Cyg5M2vYIby}
}

func (hl8w2lFX *HTBkPb) mA1mFe4YP(fIadEXIim6l []byte) (int, *FZJphYv9hV, error) {

	var qxwkC3VxG0 *FZJphYv9hV
	doauF8Y, epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.mA1mFe4YP(fIadEXIim6l)
	switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
	case *syscall.Hx8lqxSkV:
		qxwkC3VxG0 = &FZJphYv9hV{GdouSNq80bRw: epTaNF0RJ8eN.Q3zz2uH[0:]}
		doauF8Y =  /*line :1*/ eBTq9u(doauF8Y, fIadEXIim6l)
	case *syscall.HKTAy7_BSU2:
		qxwkC3VxG0 = &FZJphYv9hV{GdouSNq80bRw: epTaNF0RJ8eN.Y4XxFr[0:], Cyg5M2vYIby:  /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(epTaNF0RJ8eN.AsutdJq2))}
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}

func eBTq9u(doauF8Y int, fIadEXIim6l []byte) int {
	if  /*line :1*/ len(fIadEXIim6l) < 20 {
		return doauF8Y
	}
	v3upVm :=  /*line :1*/ int(fIadEXIim6l[0]&0x0f) << 2
	if 20 > v3upVm || v3upVm >  /*line :1*/ len(fIadEXIim6l) {
		return doauF8Y
	}
	if fIadEXIim6l[0]>>4 != 4 {
		return doauF8Y
	}
	 /*line :1*/ copy(fIadEXIim6l, fIadEXIim6l[v3upVm:])
	return doauF8Y - v3upVm
}

func (hl8w2lFX *HTBkPb) vTzqIG(fIadEXIim6l, cFiVO46YAy8g []byte) (doauF8Y, ytEzh3580y, dBNjCxLYqs int, qxwkC3VxG0 *FZJphYv9hV, h_ljk48Bm error) {
	var epTaNF0RJ8eN syscall.S4iroLVT4
	doauF8Y, ytEzh3580y, dBNjCxLYqs, epTaNF0RJ8eN, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.vTzqIG(fIadEXIim6l, cFiVO46YAy8g, 0)
	switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
	case *syscall.Hx8lqxSkV:
		qxwkC3VxG0 = &FZJphYv9hV{GdouSNq80bRw: epTaNF0RJ8eN.Q3zz2uH[0:]}
	case *syscall.HKTAy7_BSU2:
		qxwkC3VxG0 = &FZJphYv9hV{GdouSNq80bRw: epTaNF0RJ8eN.Y4XxFr[0:], Cyg5M2vYIby:  /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(epTaNF0RJ8eN.AsutdJq2))}
	}
	return
}

func (hl8w2lFX *HTBkPb) omQ684zc(fIadEXIim6l []byte, qxwkC3VxG0 *FZJphYv9hV) (int, error) {
	if hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK {
		return 0, GcnTprLq
	}
	if qxwkC3VxG0 == nil {
		return 0, tccDv0yXpU
	}
	epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ qxwkC3VxG0.jIv0wLJwloQ(hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k)
	if h_ljk48Bm != nil {
		return 0, h_ljk48Bm
	}
	return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.omQ684zc(fIadEXIim6l, epTaNF0RJ8eN)
}

func (hl8w2lFX *HTBkPb) fh4B50S(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 *FZJphYv9hV) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK {
		return 0, 0, GcnTprLq
	}
	if qxwkC3VxG0 == nil {
		return 0, 0, tccDv0yXpU
	}
	epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ qxwkC3VxG0.jIv0wLJwloQ(hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k)
	if h_ljk48Bm != nil {
		return 0, 0, h_ljk48Bm
	}
	return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.fh4B50S(fIadEXIim6l, cFiVO46YAy8g, epTaNF0RJ8eN)
}

func (dKLkhP *vTALwD6Cw) xQAA5av(yESLyde9LHhT context.MBCyA6, bFNyUpAx, wcA44hhD *FZJphYv9hV) (*HTBkPb, error) {
	vsbiWLn7reX, iaebRs5X3, h_ljk48Bm :=  /*line :1*/ nihn79f8iY(yESLyde9LHhT, dKLkhP.cyZLILkT, true)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	switch vsbiWLn7reX {
	case "ip", "ip4", "ip6":
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(dKLkhP.cyZLILkT)
	}
	rhsSAD7C := dKLkhP.FySYI4P.CTcox1A2lWK
	if rhsSAD7C == nil && dKLkhP.FySYI4P.YRwug9Xuacz != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ dKLkhP.FySYI4P.YRwug9Xuacz(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, vsbiWLn7reX, bFNyUpAx, wcA44hhD, syscall.SOCK_RAW, iaebRs5X3, "dial", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ c8Eh3kKno(vyaiiPXm6k_W), nil
}

func (bw2hLt *tYwm0DwfSTEv) hq2iTKku3u(yESLyde9LHhT context.MBCyA6, bFNyUpAx *FZJphYv9hV) (*HTBkPb, error) {
	vsbiWLn7reX, iaebRs5X3, h_ljk48Bm :=  /*line :1*/ nihn79f8iY(yESLyde9LHhT, bw2hLt.aNAwQMxeTY, true)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	switch vsbiWLn7reX {
	case "ip", "ip4", "ip6":
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(bw2hLt.aNAwQMxeTY)
	}
	var rhsSAD7C func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error
	if bw2hLt.MC4oQv.OJGgvEIn7No != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ bw2hLt.MC4oQv.OJGgvEIn7No(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, vsbiWLn7reX, bFNyUpAx, nil, syscall.SOCK_RAW, iaebRs5X3, "listen", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ c8Eh3kKno(vyaiiPXm6k_W), nil
}
