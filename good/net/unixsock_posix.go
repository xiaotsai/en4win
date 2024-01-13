//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package gF9mZs2

import (
	context "fOYpb3F03EG"
	errors "iAHoxjmM"
	os "hPMrClpC"
	syscall "bUKeamF"
)

func egoWzbMJtMkt(yESLyde9LHhT context.MBCyA6, gF9mZs2 string, bFNyUpAx, wcA44hhD jIv0wLJwloQ, bcbtTs string, b0rQyegaPi func(context.MBCyA6, string, string, syscall.CVnV1i) error) (*jmJczkl1, error) {
	var dr86pU int
	switch gF9mZs2 {
	case "unix":
		dr86pU = syscall.SOCK_STREAM
	case "unixgram":
		dr86pU = syscall.SOCK_DGRAM
	case "unixpacket":
		dr86pU = syscall.SOCK_SEQPACKET
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(gF9mZs2)
	}

	switch bcbtTs {
	case "dial":
		if bFNyUpAx != nil &&  /*line :1*/ bFNyUpAx.tmKV_f7MKX() {
			bFNyUpAx = nil
		}
		if wcA44hhD != nil &&  /*line :1*/ wcA44hhD.tmKV_f7MKX() {
			wcA44hhD = nil
		}
		if wcA44hhD == nil && (dr86pU != syscall.SOCK_DGRAM || bFNyUpAx == nil) {
			return nil, tccDv0yXpU
		}
	case "listen":
	default:
		return nil,  /*line :1*/ errors.Su6g6hRoi9X("unknown mode: " + bcbtTs)
	}

	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ _tP9pgLw(yESLyde9LHhT, gF9mZs2, syscall.AF_UNIX, dr86pU, 0, false, bFNyUpAx, wcA44hhD, b0rQyegaPi)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return vyaiiPXm6k_W, nil
}

func fdm4acHSd0vA(epTaNF0RJ8eN syscall.S4iroLVT4) EqbMXsaU1lE {
	if dRoFccaZ, d30HIwxht := epTaNF0RJ8eN.(*syscall.Vaumry9a); d30HIwxht {
		return &E5nPmTZQM{HIY_L1X2: dRoFccaZ.Ql_3XaCdaHy, GrF5D_J: "unix"}
	}
	return nil
}

func d1JbegF(epTaNF0RJ8eN syscall.S4iroLVT4) EqbMXsaU1lE {
	if dRoFccaZ, d30HIwxht := epTaNF0RJ8eN.(*syscall.Vaumry9a); d30HIwxht {
		return &E5nPmTZQM{HIY_L1X2: dRoFccaZ.Ql_3XaCdaHy, GrF5D_J: "unixgram"}
	}
	return nil
}

func _a2TukK9(epTaNF0RJ8eN syscall.S4iroLVT4) EqbMXsaU1lE {
	if dRoFccaZ, d30HIwxht := epTaNF0RJ8eN.(*syscall.Vaumry9a); d30HIwxht {
		return &E5nPmTZQM{HIY_L1X2: dRoFccaZ.Ql_3XaCdaHy, GrF5D_J: "unixpacket"}
	}
	return nil
}

func hDMjdPt(dr86pU int) string {
	switch dr86pU {
	case syscall.SOCK_STREAM:
		return "unix"
	case syscall.SOCK_DGRAM:
		return "unixgram"
	case syscall.SOCK_SEQPACKET:
		return "unixpacket"
	default:
		 /*line :1*/ panic("sotypeToNet unknown socket type")
	}
}

func (uI7LZDHm6 *E5nPmTZQM) nqsgnnaf() int {
	return syscall.AF_UNIX
}

func (uI7LZDHm6 *E5nPmTZQM) jIv0wLJwloQ(nqsgnnaf int) (syscall.S4iroLVT4, error) {
	if uI7LZDHm6 == nil {
		return nil, nil
	}
	return &syscall.Vaumry9a{Ql_3XaCdaHy: uI7LZDHm6.HIY_L1X2}, nil
}

func (uI7LZDHm6 *E5nPmTZQM) kAjqX7j(gF9mZs2 string) jIv0wLJwloQ {
	return uI7LZDHm6
}

func (hl8w2lFX *Xud0Zk) mA1mFe4YP(fIadEXIim6l []byte) (int, *E5nPmTZQM, error) {
	var qxwkC3VxG0 *E5nPmTZQM
	doauF8Y, epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.mA1mFe4YP(fIadEXIim6l)
	switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
	case *syscall.Vaumry9a:
		if epTaNF0RJ8eN.Ql_3XaCdaHy != "" {
			qxwkC3VxG0 = &E5nPmTZQM{HIY_L1X2: epTaNF0RJ8eN.Ql_3XaCdaHy, GrF5D_J:  /*line :1*/ hDMjdPt(hl8w2lFX.sXhoTdSE8wtb.zpK4ak_Tpl)}
		}
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}

func (hl8w2lFX *Xud0Zk) vTzqIG(fIadEXIim6l, cFiVO46YAy8g []byte) (doauF8Y, ytEzh3580y, dBNjCxLYqs int, qxwkC3VxG0 *E5nPmTZQM, h_ljk48Bm error) {
	var epTaNF0RJ8eN syscall.S4iroLVT4
	doauF8Y, ytEzh3580y, dBNjCxLYqs, epTaNF0RJ8eN, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.vTzqIG(fIadEXIim6l, cFiVO46YAy8g, readMsgFlags)
	if readMsgFlags == 0 && h_ljk48Bm == nil && ytEzh3580y > 0 {
		 /*line :1*/ g7BJvg(cFiVO46YAy8g[:ytEzh3580y])
	}

	switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
	case *syscall.Vaumry9a:
		if epTaNF0RJ8eN.Ql_3XaCdaHy != "" {
			qxwkC3VxG0 = &E5nPmTZQM{HIY_L1X2: epTaNF0RJ8eN.Ql_3XaCdaHy, GrF5D_J:  /*line :1*/ hDMjdPt(hl8w2lFX.sXhoTdSE8wtb.zpK4ak_Tpl)}
		}
	}
	return
}

func (hl8w2lFX *Xud0Zk) omQ684zc(fIadEXIim6l []byte, qxwkC3VxG0 *E5nPmTZQM) (int, error) {
	if hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK {
		return 0, GcnTprLq
	}
	if qxwkC3VxG0 == nil {
		return 0, tccDv0yXpU
	}
	if qxwkC3VxG0.GrF5D_J !=  /*line :1*/ hDMjdPt(hl8w2lFX.sXhoTdSE8wtb.zpK4ak_Tpl) {
		return 0, syscall.EAFNOSUPPORT
	}
	epTaNF0RJ8eN := &syscall.Vaumry9a{Ql_3XaCdaHy: qxwkC3VxG0.HIY_L1X2}
	return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.omQ684zc(fIadEXIim6l, epTaNF0RJ8eN)
}

func (hl8w2lFX *Xud0Zk) fh4B50S(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 *E5nPmTZQM) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if hl8w2lFX.sXhoTdSE8wtb.zpK4ak_Tpl == syscall.SOCK_DGRAM && hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK {
		return 0, 0, GcnTprLq
	}
	var epTaNF0RJ8eN syscall.S4iroLVT4
	if qxwkC3VxG0 != nil {
		if qxwkC3VxG0.GrF5D_J !=  /*line :1*/ hDMjdPt(hl8w2lFX.sXhoTdSE8wtb.zpK4ak_Tpl) {
			return 0, 0, syscall.EAFNOSUPPORT
		}
		epTaNF0RJ8eN = &syscall.Vaumry9a{Ql_3XaCdaHy: qxwkC3VxG0.HIY_L1X2}
	}
	return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.fh4B50S(fIadEXIim6l, cFiVO46YAy8g, epTaNF0RJ8eN)
}

func (dKLkhP *vTALwD6Cw) g5IHkyH(yESLyde9LHhT context.MBCyA6, bFNyUpAx, wcA44hhD *E5nPmTZQM) (*Xud0Zk, error) {
	rhsSAD7C := dKLkhP.FySYI4P.CTcox1A2lWK
	if rhsSAD7C == nil && dKLkhP.FySYI4P.YRwug9Xuacz != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ dKLkhP.FySYI4P.YRwug9Xuacz(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ egoWzbMJtMkt(yESLyde9LHhT, dKLkhP.cyZLILkT, bFNyUpAx, wcA44hhD, "dial", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ pUWAqmsC4s(vyaiiPXm6k_W), nil
}

func (dhvebBa4tmD *X3_EsNk) zjXrFvdT() (*Xud0Zk, error) {
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ dhvebBa4tmD.wTPbhAE.zjXrFvdT()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ pUWAqmsC4s(vyaiiPXm6k_W), nil
}

func (dhvebBa4tmD *X3_EsNk) tFIv2Ppbx0H() error {

	 /*line :1*/ dhvebBa4tmD.v96dRno.Do(func() {
		if dhvebBa4tmD.ghYAeP[0] != '@' && dhvebBa4tmD.u612R4LZ {
			 /*line :1*/ syscall.Q83qUK4q(dhvebBa4tmD.ghYAeP)
		}
	})
	return  /*line :1*/ dhvebBa4tmD.wTPbhAE.Close()
}

func (dhvebBa4tmD *X3_EsNk) cZsWMjIJ6() (*os.BvGocYcXx, error) {
	t5q9DVOD9, h_ljk48Bm :=  /*line :1*/ dhvebBa4tmD.wTPbhAE.uNYdhMYF()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return t5q9DVOD9, nil
}

func (v3upVm *X3_EsNk) SetUnlinkOnClose(x68dMeun bool) {
	v3upVm.u612R4LZ = x68dMeun
}

func (bw2hLt *tYwm0DwfSTEv) caN1IhSmUV(yESLyde9LHhT context.MBCyA6, bFNyUpAx *E5nPmTZQM) (*X3_EsNk, error) {
	var rhsSAD7C func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error
	if bw2hLt.MC4oQv.OJGgvEIn7No != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ bw2hLt.MC4oQv.OJGgvEIn7No(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ egoWzbMJtMkt(yESLyde9LHhT, bw2hLt.aNAwQMxeTY, bFNyUpAx, nil, "listen", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return &X3_EsNk{wTPbhAE: vyaiiPXm6k_W, ghYAeP:  /*line :1*/ vyaiiPXm6k_W.y9pjEf.String(), u612R4LZ: true}, nil
}

func (bw2hLt *tYwm0DwfSTEv) r5LehBZ4(yESLyde9LHhT context.MBCyA6, bFNyUpAx *E5nPmTZQM) (*Xud0Zk, error) {
	var rhsSAD7C func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error
	if bw2hLt.MC4oQv.OJGgvEIn7No != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ bw2hLt.MC4oQv.OJGgvEIn7No(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ egoWzbMJtMkt(yESLyde9LHhT, bw2hLt.aNAwQMxeTY, bFNyUpAx, nil, "listen", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ pUWAqmsC4s(vyaiiPXm6k_W), nil
}
