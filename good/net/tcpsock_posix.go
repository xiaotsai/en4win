//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package gF9mZs2

import (
	context "fOYpb3F03EG"
	io "xI9ai2djaJ2"
	os "hPMrClpC"
	syscall "bUKeamF"
)

func ya0ZV8JJhjc(epTaNF0RJ8eN syscall.S4iroLVT4) EqbMXsaU1lE {
	switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
	case *syscall.Hx8lqxSkV:
		return &HiOzqOVjN1P{ERvaNiiVmytR: epTaNF0RJ8eN.Q3zz2uH[0:], EGf1xSAbsOU5: epTaNF0RJ8eN.AMBsbT8war}
	case *syscall.HKTAy7_BSU2:
		return &HiOzqOVjN1P{ERvaNiiVmytR: epTaNF0RJ8eN.Y4XxFr[0:], EGf1xSAbsOU5: epTaNF0RJ8eN.NtLA3k, RgTF6a1Xmo:  /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int(epTaNF0RJ8eN.AsutdJq2))}
	}
	return nil
}

func (uI7LZDHm6 *HiOzqOVjN1P) nqsgnnaf() int {
	if uI7LZDHm6 == nil ||  /*line :1*/ len(uI7LZDHm6.ERvaNiiVmytR) <= IPv4len {
		return syscall.AF_INET
	}
	if  /*line :1*/ uI7LZDHm6.ERvaNiiVmytR.To4() != nil {
		return syscall.AF_INET
	}
	return syscall.AF_INET6
}

func (uI7LZDHm6 *HiOzqOVjN1P) jIv0wLJwloQ(nqsgnnaf int) (syscall.S4iroLVT4, error) {
	if uI7LZDHm6 == nil {
		return nil, nil
	}
	return  /*line :1*/ ggQVAaURvc(nqsgnnaf, uI7LZDHm6.ERvaNiiVmytR, uI7LZDHm6.EGf1xSAbsOU5, uI7LZDHm6.RgTF6a1Xmo)
}

func (uI7LZDHm6 *HiOzqOVjN1P) kAjqX7j(gF9mZs2 string) jIv0wLJwloQ {
	return &HiOzqOVjN1P{ /*line :1*/ gygXisQVD(gF9mZs2), uI7LZDHm6.EGf1xSAbsOU5, uI7LZDHm6.RgTF6a1Xmo}
}

func (hl8w2lFX *Q2rqhDKoOIxM) mA1mFe4YP(yxhs4Z io.YJ04Fau) (int64, error) {
	if doauF8Y, h_ljk48Bm, aQiEMDY6O :=  /*line :1*/ x8MzI3xRT7(hl8w2lFX.sXhoTdSE8wtb, yxhs4Z); aQiEMDY6O {
		return doauF8Y, h_ljk48Bm
	}
	if doauF8Y, h_ljk48Bm, aQiEMDY6O :=  /*line :1*/ hm7hJsDd0lSi(hl8w2lFX.sXhoTdSE8wtb, yxhs4Z); aQiEMDY6O {
		return doauF8Y, h_ljk48Bm
	}
	return  /*line :1*/ gvaVYabrVviG(hl8w2lFX, yxhs4Z)
}

func (dKLkhP *vTALwD6Cw) a9XNPFb(yESLyde9LHhT context.MBCyA6, bFNyUpAx, wcA44hhD *HiOzqOVjN1P) (*Q2rqhDKoOIxM, error) {
	if cBprzKT8 := dKLkhP.rcTu3A; cBprzKT8 != nil {
		return  /*line :1*/ cBprzKT8(yESLyde9LHhT, dKLkhP.cyZLILkT, bFNyUpAx, wcA44hhD)
	}
	if cBprzKT8 := fwPYiSsYito; cBprzKT8 != nil {
		return  /*line :1*/ cBprzKT8(yESLyde9LHhT, dKLkhP.cyZLILkT, bFNyUpAx, wcA44hhD)
	}
	return  /*line :1*/ dKLkhP.kQn5cPTO(yESLyde9LHhT, bFNyUpAx, wcA44hhD)
}

func (dKLkhP *vTALwD6Cw) kQn5cPTO(yESLyde9LHhT context.MBCyA6, bFNyUpAx, wcA44hhD *HiOzqOVjN1P) (*Q2rqhDKoOIxM, error) {
	return  /*line :1*/ dKLkhP.i0vj4vkCuqI(yESLyde9LHhT, bFNyUpAx, wcA44hhD, 0)
}

func (dKLkhP *vTALwD6Cw) i0vj4vkCuqI(yESLyde9LHhT context.MBCyA6, bFNyUpAx, wcA44hhD *HiOzqOVjN1P, iaebRs5X3 int) (*Q2rqhDKoOIxM, error) {
	rhsSAD7C := dKLkhP.FySYI4P.CTcox1A2lWK
	if rhsSAD7C == nil && dKLkhP.FySYI4P.YRwug9Xuacz != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ dKLkhP.FySYI4P.YRwug9Xuacz(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, dKLkhP.cyZLILkT, bFNyUpAx, wcA44hhD, syscall.SOCK_STREAM, iaebRs5X3, "dial", rhsSAD7C)

	for eaAUaB7Z := 0; eaAUaB7Z < 2 && (bFNyUpAx == nil || bFNyUpAx.EGf1xSAbsOU5 == 0) && ( /*line :1*/ zrlAMpZ(vyaiiPXm6k_W, h_ljk48Bm) ||  /*line :1*/ ipZwaqZ3(h_ljk48Bm)); eaAUaB7Z++ {
		if h_ljk48Bm == nil {
			 /*line :1*/ vyaiiPXm6k_W.Close()
		}
		vyaiiPXm6k_W, h_ljk48Bm =  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, dKLkhP.cyZLILkT, bFNyUpAx, wcA44hhD, syscall.SOCK_STREAM, iaebRs5X3, "dial", rhsSAD7C)
	}

	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ xMgOZ10goOXP(vyaiiPXm6k_W, dKLkhP.FySYI4P.HQGGkRp, fsDujdBrz), nil
}

func zrlAMpZ(vyaiiPXm6k_W *jmJczkl1, h_ljk48Bm error) bool {

	if h_ljk48Bm != nil {
		return false
	}

	if vyaiiPXm6k_W.y9pjEf == nil || vyaiiPXm6k_W.jeyDlYdkG0nF == nil {
		return true
	}
	v3upVm := vyaiiPXm6k_W.y9pjEf.(*HiOzqOVjN1P)
	yxhs4Z := vyaiiPXm6k_W.jeyDlYdkG0nF.(*HiOzqOVjN1P)
	return v3upVm.EGf1xSAbsOU5 == yxhs4Z.EGf1xSAbsOU5 &&  /*line :1*/ v3upVm.ERvaNiiVmytR.Equal(yxhs4Z.ERvaNiiVmytR)
}

func ipZwaqZ3(h_ljk48Bm error) bool {
	if qCPkwqDhs, d30HIwxht := h_ljk48Bm.(*ZOYGdck); d30HIwxht {
		h_ljk48Bm = qCPkwqDhs.XkWH4CYwNmvD
	}
	if cW6_CaaHVh, d30HIwxht := h_ljk48Bm.(*os.XGyr7J); d30HIwxht {
		h_ljk48Bm = cW6_CaaHVh.V_xDVM
	}
	return h_ljk48Bm == syscall.EADDRNOTAVAIL
}

func (dhvebBa4tmD *S0u1keY) d30HIwxht() bool	{ return dhvebBa4tmD != nil && dhvebBa4tmD._Tj5wN != nil }

func (dhvebBa4tmD *S0u1keY) zjXrFvdT() (*Q2rqhDKoOIxM, error) {
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ dhvebBa4tmD._Tj5wN.zjXrFvdT()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return  /*line :1*/ xMgOZ10goOXP(vyaiiPXm6k_W, dhvebBa4tmD.kuQlluF.Wlhapadf, nil), nil
}

func (dhvebBa4tmD *S0u1keY) tFIv2Ppbx0H() error {
	return  /*line :1*/ dhvebBa4tmD._Tj5wN.Close()
}

func (dhvebBa4tmD *S0u1keY) cZsWMjIJ6() (*os.BvGocYcXx, error) {
	t5q9DVOD9, h_ljk48Bm :=  /*line :1*/ dhvebBa4tmD._Tj5wN.uNYdhMYF()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return t5q9DVOD9, nil
}

func (bw2hLt *tYwm0DwfSTEv) bwWMnaxr(yESLyde9LHhT context.MBCyA6, bFNyUpAx *HiOzqOVjN1P) (*S0u1keY, error) {
	return  /*line :1*/ bw2hLt.hpdSTocw1Jh2(yESLyde9LHhT, bFNyUpAx, 0)
}

func (bw2hLt *tYwm0DwfSTEv) hpdSTocw1Jh2(yESLyde9LHhT context.MBCyA6, bFNyUpAx *HiOzqOVjN1P, iaebRs5X3 int) (*S0u1keY, error) {
	var rhsSAD7C func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error
	if bw2hLt.MC4oQv.OJGgvEIn7No != nil {
		rhsSAD7C = func(bnbaCXzk context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error {
			return  /*line :1*/ bw2hLt.MC4oQv.OJGgvEIn7No(vsbiWLn7reX, fwV_ln2dl, hl8w2lFX)
		}
	}
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ e6nNXDsTnM(yESLyde9LHhT, bw2hLt.aNAwQMxeTY, bFNyUpAx, nil, syscall.SOCK_STREAM, iaebRs5X3, "listen", rhsSAD7C)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return &S0u1keY{_Tj5wN: vyaiiPXm6k_W, kuQlluF: bw2hLt.MC4oQv}, nil
}
