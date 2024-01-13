//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package gF9mZs2

import (
	context "fOYpb3F03EG"
	netip "iPdCam_KQOXj"
	syscall "bUKeamF"
)

func cYt7M3(epTaNF0RJ8eN syscall.S4iroLVT4) EqbMXsaU1lE {
	switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
	case *syscall.Hx8lqxSkV:
		return &ZaffanpNx4{ERvaNiiVmytR: epTaNF0RJ8eN.Q3zz2uH[0:], EGf1xSAbsOU5: epTaNF0RJ8eN.AMBsbT8war}
	case *syscall.HKTAy7_BSU2:
		return &ZaffanpNx4{ERvaNiiVmytR: epTaNF0RJ8eN.Y4XxFr[0:], EGf1xSAbsOU5: epTaNF0RJ8eN.NtLA3k, RgTF6a1Xmo:  /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(epTaNF0RJ8eN.AsutdJq2))}
	}
	return nil
}

func (uI7LZDHm6 *ZaffanpNx4) nqsgnnaf() int {
	if uI7LZDHm6 == nil ||  /*line :1*/ len(uI7LZDHm6.ERvaNiiVmytR) <= IPv4len {
		return syscall.AF_INET
	}
	if  /*line :1*/ uI7LZDHm6.ERvaNiiVmytR.To4() != nil {
		return syscall.AF_INET
	}
	return syscall.AF_INET6
}

func (uI7LZDHm6 *ZaffanpNx4) jIv0wLJwloQ(nqsgnnaf int) (syscall.S4iroLVT4, error) {
	if uI7LZDHm6 == nil {
		return nil, nil
	}
	return  /*line :1*/ ggQVAaURvc(nqsgnnaf, uI7LZDHm6.ERvaNiiVmytR, uI7LZDHm6.EGf1xSAbsOU5, uI7LZDHm6.RgTF6a1Xmo)
}

func (uI7LZDHm6 *ZaffanpNx4) kAjqX7j(gF9mZs2 string) jIv0wLJwloQ {
	return &ZaffanpNx4{ /*line :1*/ gygXisQVD(gF9mZs2), uI7LZDHm6.EGf1xSAbsOU5, uI7LZDHm6.RgTF6a1Xmo}
}

func (hl8w2lFX *JdInDLIal) mA1mFe4YP(fIadEXIim6l []byte, qxwkC3VxG0 *ZaffanpNx4) (int, *ZaffanpNx4, error) {
	var doauF8Y int
	var h_ljk48Bm error
	switch hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k {
	case syscall.AF_INET:
		var sdGaORW_DS syscall.Hx8lqxSkV
		doauF8Y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.aqQMdv2(fIadEXIim6l, &sdGaORW_DS)
		if h_ljk48Bm == nil {
			kQ8_UEhxU := sdGaORW_DS.Q3zz2uH
			*qxwkC3VxG0 = ZaffanpNx4{ERvaNiiVmytR: kQ8_UEhxU[:], EGf1xSAbsOU5: sdGaORW_DS.AMBsbT8war}
		}
	case syscall.AF_INET6:
		var sdGaORW_DS syscall.HKTAy7_BSU2
		doauF8Y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.wI_ExTyC(fIadEXIim6l, &sdGaORW_DS)
		if h_ljk48Bm == nil {
			kQ8_UEhxU := sdGaORW_DS.Y4XxFr
			*qxwkC3VxG0 = ZaffanpNx4{ERvaNiiVmytR: kQ8_UEhxU[:], EGf1xSAbsOU5: sdGaORW_DS.NtLA3k, RgTF6a1Xmo:  /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(sdGaORW_DS.AsutdJq2))}
		}
	}
	if h_ljk48Bm != nil {

		qxwkC3VxG0 = nil
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}

func (hl8w2lFX *JdInDLIal) bAIcoRDfR0(fIadEXIim6l []byte) (doauF8Y int, qxwkC3VxG0 netip.YTqTf_4VC, h_ljk48Bm error) {
	var kQ8_UEhxU netip.KFLQ1_1E
	var mzUgFPgljgC int
	switch hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k {
	case syscall.AF_INET:
		var sdGaORW_DS syscall.Hx8lqxSkV
		doauF8Y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.aqQMdv2(fIadEXIim6l, &sdGaORW_DS)
		if h_ljk48Bm == nil {
			kQ8_UEhxU =  /*line :1*/ netip.Vh24QUEOnu(sdGaORW_DS.Q3zz2uH)
			mzUgFPgljgC = sdGaORW_DS.AMBsbT8war
		}
	case syscall.AF_INET6:
		var sdGaORW_DS syscall.HKTAy7_BSU2
		doauF8Y, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.wI_ExTyC(fIadEXIim6l, &sdGaORW_DS)
		if h_ljk48Bm == nil {
			kQ8_UEhxU =  /*line :1*/ netip.XtvbgBR(sdGaORW_DS.Y4XxFr).WithZone( /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(sdGaORW_DS.AsutdJq2)))
			mzUgFPgljgC = sdGaORW_DS.NtLA3k
		}
	}
	if h_ljk48Bm == nil {
		qxwkC3VxG0 =  /*line :1*/ netip.Iqr8_1pEDmw0(kQ8_UEhxU,  /*line :1*/ uint16(mzUgFPgljgC))
	}
	return doauF8Y, qxwkC3VxG0, h_ljk48Bm
}

func (hl8w2lFX *JdInDLIal) vTzqIG(fIadEXIim6l, cFiVO46YAy8g []byte) (doauF8Y, ytEzh3580y, dBNjCxLYqs int, qxwkC3VxG0 netip.YTqTf_4VC, h_ljk48Bm error) {
	switch hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k {
	case syscall.AF_INET:
		var epTaNF0RJ8eN syscall.Hx8lqxSkV
		doauF8Y, ytEzh3580y, dBNjCxLYqs, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.gXSpGdZ05(fIadEXIim6l, cFiVO46YAy8g, 0, &epTaNF0RJ8eN)
		kQ8_UEhxU :=  /*line :1*/ netip.Vh24QUEOnu(epTaNF0RJ8eN.Q3zz2uH)
		qxwkC3VxG0 =  /*line :1*/ netip.Iqr8_1pEDmw0(kQ8_UEhxU,  /*line :1*/ uint16(epTaNF0RJ8eN.AMBsbT8war))
	case syscall.AF_INET6:
		var epTaNF0RJ8eN syscall.HKTAy7_BSU2
		doauF8Y, ytEzh3580y, dBNjCxLYqs, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.jQ3N__FhSo(fIadEXIim6l, cFiVO46YAy8g, 0, &epTaNF0RJ8eN)
		kQ8_UEhxU :=  /*line :1*/ netip.XtvbgBR(epTaNF0RJ8eN.Y4XxFr).WithZone( /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(epTaNF0RJ8eN.AsutdJq2)))
		qxwkC3VxG0 =  /*line :1*/ netip.Iqr8_1pEDmw0(kQ8_UEhxU,  /*line :1*/ uint16(epTaNF0RJ8eN.NtLA3k))
	}
	return
}

func (hl8w2lFX *JdInDLIal) omQ684zc(fIadEXIim6l []byte, qxwkC3VxG0 *ZaffanpNx4) (int, error) {
	if hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK {
		return 0, GcnTprLq
	}
	if qxwkC3VxG0 == nil {
		return 0, tccDv0yXpU
	}

	switch hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k {
	case syscall.AF_INET:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ qBOiy4(qxwkC3VxG0.ERvaNiiVmytR, qxwkC3VxG0.EGf1xSAbsOU5)
		if h_ljk48Bm != nil {
			return 0, h_ljk48Bm
		}
		return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.edtDmORC(fIadEXIim6l, &epTaNF0RJ8eN)
	case syscall.AF_INET6:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ eWiFFUFYjR(qxwkC3VxG0.ERvaNiiVmytR, qxwkC3VxG0.EGf1xSAbsOU5, qxwkC3VxG0.RgTF6a1Xmo)
		if h_ljk48Bm != nil {
			return 0, h_ljk48Bm
		}
		return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.wPqovgf6d(fIadEXIim6l, &epTaNF0RJ8eN)
	default:
		return 0, &DAWLIQHc{Z68v0y: "invalid address family", BCCnAgFxG:  /*line :1*/ qxwkC3VxG0.ERvaNiiVmytR.String()}
	}
}

func (hl8w2lFX *JdInDLIal) goun3aKzAu(fIadEXIim6l []byte, qxwkC3VxG0 netip.YTqTf_4VC) (int, error) {
	if hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK {
		return 0, GcnTprLq
	}
	if ! /*line :1*/ qxwkC3VxG0.IsValid() {
		return 0, tccDv0yXpU
	}

	switch hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k {
	case syscall.AF_INET:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ xifkdya(qxwkC3VxG0)
		if h_ljk48Bm != nil {
			return 0, h_ljk48Bm
		}
		return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.edtDmORC(fIadEXIim6l, &epTaNF0RJ8eN)
	case syscall.AF_INET6:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ tk4LBjOZ_EZy(qxwkC3VxG0)
		if h_ljk48Bm != nil {
			return 0, h_ljk48Bm
		}
		return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.wPqovgf6d(fIadEXIim6l, &epTaNF0RJ8eN)
	default:
		return 0, &DAWLIQHc{Z68v0y: "invalid address family", BCCnAgFxG:  /*line :1*/ qxwkC3VxG0.Addr().String()}
	}
}

func (hl8w2lFX *JdInDLIal) fh4B50S(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 *ZaffanpNx4) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK && qxwkC3VxG0 != nil {
		return 0, 0, GcnTprLq
	}
	if !hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK && qxwkC3VxG0 == nil {
		return 0, 0, tccDv0yXpU
	}
	epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ qxwkC3VxG0.jIv0wLJwloQ(hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k)
	if h_ljk48Bm != nil {
		return 0, 0, h_ljk48Bm
	}
	return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.fh4B50S(fIadEXIim6l, cFiVO46YAy8g, epTaNF0RJ8eN)
}

func (hl8w2lFX *JdInDLIal) s3pDutKjt9Qk(fIadEXIim6l, cFiVO46YAy8g []byte, qxwkC3VxG0 netip.YTqTf_4VC) (doauF8Y, ytEzh3580y int, h_ljk48Bm error) {
	if hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK &&  /*line :1*/ qxwkC3VxG0.IsValid() {
		return 0, 0, GcnTprLq
	}
	if !hl8w2lFX.sXhoTdSE8wtb.kBOwaEVMqK && ! /*line :1*/ qxwkC3VxG0.IsValid() {
		return 0, 0, tccDv0yXpU
	}

	switch hl8w2lFX.sXhoTdSE8wtb.zW6I0G4Xm0k {
	case syscall.AF_INET:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ xifkdya(qxwkC3VxG0)
		if h_ljk48Bm != nil {
			return 0, 0, h_ljk48Bm
		}
		return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.h92kZU2ngw(fIadEXIim6l, cFiVO46YAy8g, &epTaNF0RJ8eN)
	case syscall.AF_INET6:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ tk4LBjOZ_EZy(qxwkC3VxG0)
		if h_ljk48Bm != nil {
			return 0, 0, h_ljk48Bm
		}
		return  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.i9YFSO8(fIadEXIim6l, cFiVO46YAy8g, &epTaNF0RJ8eN)
	default:
		return 0, 0, &DAWLIQHc{Z68v0y: "invalid address family", BCCnAgFxG:  /*line :1*/ qxwkC3VxG0.Addr().String()}
	}
}

func (dKLkhP *vTALwD6Cw) ubPg4ZjrNYRu(yESLyde9LHhT context.MBCyA6, bFNyUpAx, wcA44hhD *ZaffanpNx4) (*JdInDLIal, error) {
	rhsSAD7C := dKLkhP.FySYI4P.CTcox1A2lWK
	if rhsSAD7C == nil && dKLkhP.FySYI4P.YRwug9Xuacz != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ dKLkhP.FySYI4P.YRwug9Xuacz(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, dKLkhP.cyZLILkT, bFNyUpAx, wcA44hhD, syscall.SOCK_DGRAM, 0, "dial", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ mgV5c9(vyaiiPXm6k_W), nil
}

func (bw2hLt *tYwm0DwfSTEv) eg0RUV(yESLyde9LHhT context.MBCyA6, bFNyUpAx *ZaffanpNx4) (*JdInDLIal, error) {
	var rhsSAD7C func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error
	if bw2hLt.MC4oQv.OJGgvEIn7No != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ bw2hLt.MC4oQv.OJGgvEIn7No(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, bw2hLt.aNAwQMxeTY, bFNyUpAx, nil, syscall.SOCK_DGRAM, 0, "listen", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ mgV5c9(vyaiiPXm6k_W), nil
}

func (bw2hLt *tYwm0DwfSTEv) zdApnUghKfl(yESLyde9LHhT context.MBCyA6, dvFmDYmQz *U_Uc9la, oqJawTas3 *ZaffanpNx4) (*JdInDLIal, error) {
	var rhsSAD7C func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error
	if bw2hLt.MC4oQv.OJGgvEIn7No != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ bw2hLt.MC4oQv.OJGgvEIn7No(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, bw2hLt.aNAwQMxeTY, oqJawTas3, nil, syscall.SOCK_DGRAM, 0, "listen", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	hl8w2lFX :=  /*line :1*/ mgV5c9(vyaiiPXm6k_W)
	if di2u16 :=  /*line :1*/ oqJawTas3.ERvaNiiVmytR.To4(); di2u16 != nil {
		if h_ljk48Bm :=  /*line :1*/ m5f5rgsoKeM7(hl8w2lFX, dvFmDYmQz, di2u16); h_ljk48Bm != nil {
			 /*line :1*/ hl8w2lFX.Close()
			return nil, h_ljk48Bm
		}
	} else {
		if h_ljk48Bm :=  /*line :1*/ bnRZpqSUmZqu(hl8w2lFX, dvFmDYmQz, oqJawTas3.ERvaNiiVmytR); h_ljk48Bm != nil {
			 /*line :1*/ hl8w2lFX.Close()
			return nil, h_ljk48Bm
		}
	}
	return hl8w2lFX, nil
}

func m5f5rgsoKeM7(hl8w2lFX *JdInDLIal, dvFmDYmQz *U_Uc9la, kQ8_UEhxU GNraIvYhB) error {
	if dvFmDYmQz != nil {
		if h_ljk48Bm :=  /*line :1*/ iq2nJr(hl8w2lFX.sXhoTdSE8wtb, dvFmDYmQz); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
	}
	if h_ljk48Bm :=  /*line :1*/ kPOWpW616(hl8w2lFX.sXhoTdSE8wtb, false); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	if h_ljk48Bm :=  /*line :1*/ visCDO9kf(hl8w2lFX.sXhoTdSE8wtb, dvFmDYmQz, kQ8_UEhxU); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	return nil
}

func bnRZpqSUmZqu(hl8w2lFX *JdInDLIal, dvFmDYmQz *U_Uc9la, kQ8_UEhxU GNraIvYhB) error {
	if dvFmDYmQz != nil {
		if h_ljk48Bm :=  /*line :1*/ tayHWdA(hl8w2lFX.sXhoTdSE8wtb, dvFmDYmQz); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
	}
	if h_ljk48Bm :=  /*line :1*/ rVEEX2Rxk(hl8w2lFX.sXhoTdSE8wtb, false); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	if h_ljk48Bm :=  /*line :1*/ bIw1jOVnFhg(hl8w2lFX.sXhoTdSE8wtb, dvFmDYmQz, kQ8_UEhxU); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	return nil
}
