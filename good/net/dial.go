//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	godebug "rayT9AgS"
	nettrace "KwiKvcnR5Bq"
	syscall "bUKeamF"
	time "fRAfQd_"
)

const (
	
	
	defaultTCPKeepAlive	= 15 * time.Second

	
	
	defaultMPTCPEnabled	= false
)

var jMz8EV =  /*line :1*/ godebug.HppHHAII("multipathtcp")


type am59Br2lmBKa uint8

const (
	
	mptcpUseDefault	am59Br2lmBKa	= iota
	mptcpEnabled
	mptcpDisabled
)

func (z_nqmUMeJFH *am59Br2lmBKa) cCaQG2_M1() bool {
	switch *z_nqmUMeJFH {
	case mptcpEnabled:
		return true
	case mptcpDisabled:
		return false
	}

	if  /*line :1*/ jMz8EV.Value() == "1" {
		 /*line :1*/ jMz8EV.IncNonDefault()

		return true
	}

	return defaultMPTCPEnabled
}

func (z_nqmUMeJFH *am59Br2lmBKa) hVlQmKILF2(mFaToFtKuuHJ bool) {
	if mFaToFtKuuHJ {
		*z_nqmUMeJFH = mptcpEnabled
	} else {
		*z_nqmUMeJFH = mptcpDisabled
	}
}








type FySYI4P struct {
	
	
	
	
	
	
	
	
	
	
	
	
	I6mnWza2bU	time.GKMwTGxQa0S

	
	
	
	
	TzkrFSRA	time.G4KDkI

	
	
	
	
	MDCJO9payv	EqbMXsaU1lE

	
	
	
	
	
	
	
	ZVSiBVfg	bool

	
	
	
	
	
	
	
	
	JZKWk_mKDA	time.GKMwTGxQa0S

	
	
	
	
	
	
	
	HQGGkRp	time.GKMwTGxQa0S

	
	X5LCfT	*W8Q2tfHAD

	
	
	
	
	
	DNrYdxGWr	<-chan struct{}

	
	
	
	
	
	
	
	
	YRwug9Xuacz	func(vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error

	
	
	
	
	
	
	
	
	CTcox1A2lWK	func(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error

	
	
	
	bpogz26np	am59Br2lmBKa
}

func (ica44Q *FySYI4P) sMqHctvSisO() bool	{ return ica44Q.JZKWk_mKDA >= 0 }

func jIL275bb(uI7LZDHm6, fIadEXIim6l time.G4KDkI) time.G4KDkI {
	if  /*line :1*/ uI7LZDHm6.IsZero() {
		return fIadEXIim6l
	}
	if  /*line :1*/ fIadEXIim6l.IsZero() ||  /*line :1*/ uI7LZDHm6.Before(fIadEXIim6l) {
		return uI7LZDHm6
	}
	return fIadEXIim6l
}







func (ica44Q *FySYI4P) cSxcyrrgag(yESLyde9LHhT context.MBCyA6, t_5qtVaQY time.G4KDkI) (crIRzT time.G4KDkI) {
	if ica44Q.I6mnWza2bU != 0 {
		crIRzT =  /*line :1*/ t_5qtVaQY.Add(ica44Q.I6mnWza2bU)
	}
	if ica44Q, d30HIwxht :=  /*line :1*/ yESLyde9LHhT.Deadline(); d30HIwxht {
		crIRzT =  /*line :1*/ jIL275bb(crIRzT, ica44Q)
	}
	return  /*line :1*/ jIL275bb(crIRzT, ica44Q.TzkrFSRA)
}

func (ica44Q *FySYI4P) xi_4NEKx() *W8Q2tfHAD {
	if ica44Q.X5LCfT != nil {
		return ica44Q.X5LCfT
	}
	return Ic4wtIY
}



func hTeQHyBCV12i(t_5qtVaQY, cSxcyrrgag time.G4KDkI, geeSyUF8 int) (time.G4KDkI, error) {
	if  /*line :1*/ cSxcyrrgag.IsZero() {
		return cSxcyrrgag, nil
	}
	yySe5EVQ :=  /*line :1*/ cSxcyrrgag.Sub(t_5qtVaQY)
	if yySe5EVQ <= 0 {
		return time.G4KDkI{}, dklm5juf7rzD
	}

	yusr33iR := yySe5EVQ /  /*line :1*/ time.GKMwTGxQa0S(geeSyUF8)
	
	const saneMinimum = 2 * time.Second
	if yusr33iR < saneMinimum {
		if yySe5EVQ < saneMinimum {
			yusr33iR = yySe5EVQ
		} else {
			yusr33iR = saneMinimum
		}
	}
	return  /*line :1*/ t_5qtVaQY.Add(yusr33iR), nil
}

func (ica44Q *FySYI4P) dAexfVouH_6() time.GKMwTGxQa0S {
	if ica44Q.JZKWk_mKDA > 0 {
		return ica44Q.JZKWk_mKDA
	} else {
		return 300 * time.Millisecond
	}
}

func nihn79f8iY(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX string, umEa_6b bool) (mQBJul9fW string, iaebRs5X3 int, h_ljk48Bm error) {
	eaAUaB7Z :=  /*line :1*/ vMwVxD_U(vsbiWLn7reX, ':')
	if eaAUaB7Z < 0 {
		switch vsbiWLn7reX {
		case "tcp", "tcp4", "tcp6":
		case "udp", "udp4", "udp6":
		case "ip", "ip4", "ip6":
			if umEa_6b {
				return "", 0,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
			}
		case "unix", "unixgram", "unixpacket":
		default:
			return "", 0,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
		}
		return vsbiWLn7reX, 0, nil
	}
	mQBJul9fW = vsbiWLn7reX[:eaAUaB7Z]
	switch mQBJul9fW {
	case "ip", "ip4", "ip6":
		x610YmSeRObs := vsbiWLn7reX[eaAUaB7Z+1:]
		iaebRs5X3, eaAUaB7Z, d30HIwxht :=  /*line :1*/ bEihJjrbz(x610YmSeRObs)
		if !d30HIwxht || eaAUaB7Z !=  /*line :1*/ len(x610YmSeRObs) {
			iaebRs5X3, h_ljk48Bm =  /*line :1*/ heQcNLGEPBFn(yESLyde9LHhT, x610YmSeRObs)
			if h_ljk48Bm != nil {
				return "", 0, h_ljk48Bm
			}
		}
		return mQBJul9fW, iaebRs5X3, nil
	}
	return "", 0,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
}




func (yxhs4Z *W8Q2tfHAD) kq9KkFaU1B(yESLyde9LHhT context.MBCyA6, qCPkwqDhs, vsbiWLn7reX, qxwkC3VxG0 string, xfLWnvBRpZ EqbMXsaU1lE) (nTzeZ5mf, error) {
	mQBJul9fW, _, h_ljk48Bm :=  /*line :1*/ nihn79f8iY(yESLyde9LHhT, vsbiWLn7reX, true)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	if qCPkwqDhs == "dial" && qxwkC3VxG0 == "" {
		return nil, tccDv0yXpU
	}
	switch mQBJul9fW {
	case "unix", "unixgram", "unixpacket":
		qxwkC3VxG0, h_ljk48Bm :=  /*line :1*/ E3DL71v(mQBJul9fW, qxwkC3VxG0)
		if h_ljk48Bm != nil {
			return nil, h_ljk48Bm
		}
		if qCPkwqDhs == "dial" && xfLWnvBRpZ != nil &&  /*line :1*/ qxwkC3VxG0.Network() !=  /*line :1*/ xfLWnvBRpZ.Network() {
			return nil, &DAWLIQHc{Z68v0y: "mismatched local address type", BCCnAgFxG:  /*line :1*/ xfLWnvBRpZ.String()}
		}
		return nTzeZ5mf{qxwkC3VxG0}, nil
	}
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ yxhs4Z.hCHWGoQGe2q(yESLyde9LHhT, mQBJul9fW, qxwkC3VxG0)
	if h_ljk48Bm != nil || qCPkwqDhs != "dial" || xfLWnvBRpZ == nil {
		return md4QSRkO, h_ljk48Bm
	}
	var (
		qirroS7pOI0	*HiOzqOVjN1P
		gT6x4Wp	*ZaffanpNx4
		kQ8_UEhxU	*FZJphYv9hV
		luxE92SuuA9	bool
	)
	switch xfLWnvBRpZ := xfLWnvBRpZ.(type) {
	case *HiOzqOVjN1P:
		qirroS7pOI0 = xfLWnvBRpZ
		luxE92SuuA9 =  /*line :1*/ qirroS7pOI0.tmKV_f7MKX()
	case *ZaffanpNx4:
		gT6x4Wp = xfLWnvBRpZ
		luxE92SuuA9 =  /*line :1*/ gT6x4Wp.tmKV_f7MKX()
	case *FZJphYv9hV:
		kQ8_UEhxU = xfLWnvBRpZ
		luxE92SuuA9 =  /*line :1*/ kQ8_UEhxU.tmKV_f7MKX()
	}
	mUM7kah := md4QSRkO[:0]
	for _, qxwkC3VxG0 := range md4QSRkO {
		if  /*line :1*/ qxwkC3VxG0.Network() !=  /*line :1*/ xfLWnvBRpZ.Network() {
			return nil, &DAWLIQHc{Z68v0y: "mismatched local address type", BCCnAgFxG:  /*line :1*/ xfLWnvBRpZ.String()}
		}
		switch qxwkC3VxG0 := qxwkC3VxG0.(type) {
		case *HiOzqOVjN1P:
			if !luxE92SuuA9 && ! /*line :1*/ qxwkC3VxG0.tmKV_f7MKX() && ! /*line :1*/ qxwkC3VxG0.ERvaNiiVmytR.yuO3LyWZOg(qirroS7pOI0.ERvaNiiVmytR) {
				continue
			}
			mUM7kah =  /*line :1*/ append(mUM7kah, qxwkC3VxG0)
		case *ZaffanpNx4:
			if !luxE92SuuA9 && ! /*line :1*/ qxwkC3VxG0.tmKV_f7MKX() && ! /*line :1*/ qxwkC3VxG0.ERvaNiiVmytR.yuO3LyWZOg(gT6x4Wp.ERvaNiiVmytR) {
				continue
			}
			mUM7kah =  /*line :1*/ append(mUM7kah, qxwkC3VxG0)
		case *FZJphYv9hV:
			if !luxE92SuuA9 && ! /*line :1*/ qxwkC3VxG0.tmKV_f7MKX() && ! /*line :1*/ qxwkC3VxG0.GdouSNq80bRw.yuO3LyWZOg(kQ8_UEhxU.GdouSNq80bRw) {
				continue
			}
			mUM7kah =  /*line :1*/ append(mUM7kah, qxwkC3VxG0)
		}
	}
	if  /*line :1*/ len(mUM7kah) == 0 {
		return nil, &DAWLIQHc{Z68v0y:  /*line :1*/ xgpz6w9m.Error(), BCCnAgFxG:  /*line :1*/ xfLWnvBRpZ.String()}
	}
	return mUM7kah, nil
}





func (ica44Q *FySYI4P) MultipathTCP() bool {
	return  /*line :1*/ ica44Q.bpogz26np.cCaQG2_M1()
}







func (ica44Q *FySYI4P) SetMultipathTCP(mFaToFtKuuHJ bool) {
	 /*line :1*/ ica44Q.bpogz26np.hVlQmKILF2(mFaToFtKuuHJ)
}

















































func SkrZwW670VU(vsbiWLn7reX, fwV_ln2dl string) (UJYB3aCQg, error) {
	var ica44Q FySYI4P
	return  /*line :1*/ ica44Q.Dial(vsbiWLn7reX, fwV_ln2dl)
}











func QszW05Sr(vsbiWLn7reX, fwV_ln2dl string, yusr33iR time.GKMwTGxQa0S) (UJYB3aCQg, error) {
	ica44Q := FySYI4P{I6mnWza2bU: yusr33iR}
	return  /*line :1*/ ica44Q.Dial(vsbiWLn7reX, fwV_ln2dl)
}


type vTALwD6Cw struct {
	FySYI4P
	cyZLILkT, r0U09j	string
	rcTu3A	func(yESLyde9LHhT context.MBCyA6, gF9mZs2 string, bFNyUpAx, wcA44hhD *HiOzqOVjN1P) (*Q2rqhDKoOIxM, error)
}








func (ica44Q *FySYI4P) Dial(vsbiWLn7reX, fwV_ln2dl string) (UJYB3aCQg, error) {
	return  /*line :1*/ ica44Q.DialContext( /*line :1*/ context.GEcgQP5fzA(), vsbiWLn7reX, fwV_ln2dl)
}



















func (ica44Q *FySYI4P) DialContext(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string) (UJYB3aCQg, error) {
	if yESLyde9LHhT == nil {
		 /*line :1*/ panic("nil context")
	}
	cSxcyrrgag :=  /*line :1*/ ica44Q.cSxcyrrgag(yESLyde9LHhT,  /*line :1*/ time.Pc_35oTY())
	if ! /*line :1*/ cSxcyrrgag.IsZero() {
		if ica44Q, d30HIwxht :=  /*line :1*/ yESLyde9LHhT.Deadline(); !d30HIwxht ||  /*line :1*/ cSxcyrrgag.Before(ica44Q) {
			x1RIVd, yDIc25 :=  /*line :1*/ context.ZF8aTuau(yESLyde9LHhT, cSxcyrrgag)
			defer  /*line :1*/ yDIc25()
			yESLyde9LHhT = x1RIVd
		}
	}
	if jRDwuaxeP := ica44Q.DNrYdxGWr; jRDwuaxeP != nil {
		x1RIVd, yDIc25 :=  /*line :1*/ context.Cu9yRv(yESLyde9LHhT)
		defer  /*line :1*/ yDIc25()
		go func() {
			select {
			case <- /*line :1*/ jRDwuaxeP:
				 /*line :1*/ yDIc25()
			case <- /*line :1*/ x1RIVd.Done():
			}
		}()
		yESLyde9LHhT = x1RIVd
	}

	okUdpENpo_Y := yESLyde9LHhT
	if jHCsMONOrYFN, _ :=  /*line :1*/ yESLyde9LHhT.Value(nettrace.YRC9V_{}).(*nettrace.UTp90oOoUInw); jHCsMONOrYFN != nil {
		h05uztBPQE_F := *jHCsMONOrYFN
		h05uztBPQE_F.CVZFbB4EJX = nil
		h05uztBPQE_F.EAYlwNxwXa = nil
		okUdpENpo_Y =  /*line :1*/ context.Z2CaDb(okUdpENpo_Y, nettrace.YRC9V_{}, &h05uztBPQE_F)
	}

	md4QSRkO, h_ljk48Bm :=  /*line :1*/ ica44Q.xi_4NEKx().kq9KkFaU1B(okUdpENpo_Y, "dial", vsbiWLn7reX, fwV_ln2dl, ica44Q.MDCJO9payv)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}

	dKLkhP := &vTALwD6Cw{
		FySYI4P:	*ica44Q,
		cyZLILkT:	vsbiWLn7reX,
		r0U09j:	fwV_ln2dl,
	}

	var uR_fXZoDIy, vYb2cNv nTzeZ5mf
	if  /*line :1*/ ica44Q.sMqHctvSisO() && vsbiWLn7reX == "tcp" {
		uR_fXZoDIy, vYb2cNv =  /*line :1*/ md4QSRkO.eK8V6eXO0Mx(rnth8PRZsE2C)
	} else {
		uR_fXZoDIy = md4QSRkO
	}

	return  /*line :1*/ dKLkhP.fShRa1PYQp(yESLyde9LHhT, uR_fXZoDIy, vYb2cNv)
}





func (dKLkhP *vTALwD6Cw) fShRa1PYQp(yESLyde9LHhT context.MBCyA6, uR_fXZoDIy, vYb2cNv nTzeZ5mf) (UJYB3aCQg, error) {
	if  /*line :1*/ len(vYb2cNv) == 0 {
		return  /*line :1*/ dKLkhP.uGDGFRHjpl0(yESLyde9LHhT, uR_fXZoDIy)
	}

	sKajE5FH39 :=  /*line :1*/ make(chan struct{})
	defer  /*line :1*/ close(sKajE5FH39)

	type bgY9_17f8O5x struct {
		UJYB3aCQg
		error
		bNfeI48	bool
		puOrMIlI8e	bool
	}
	iaDi__SK6bka :=  /*line :1*/ make(chan bgY9_17f8O5x)

	jdkVQwN_sasp := func(yESLyde9LHhT context.MBCyA6, efLai2h8_ bool) {
		gPx_2lY2T := uR_fXZoDIy
		if !efLai2h8_ {
			gPx_2lY2T = vYb2cNv
		}
		hl8w2lFX, h_ljk48Bm :=  /*line :1*/ dKLkhP.uGDGFRHjpl0(yESLyde9LHhT, gPx_2lY2T)
		select {
		case iaDi__SK6bka <- bgY9_17f8O5x{UJYB3aCQg: hl8w2lFX, error: h_ljk48Bm, bNfeI48: efLai2h8_, puOrMIlI8e: true}:
		case <-sKajE5FH39:
			if hl8w2lFX != nil {
				 /*line :1*/ hl8w2lFX.Close()
			}
		}
	}

	var efLai2h8_, hmIvSNpIulg bgY9_17f8O5x

	uxSavJI, qncl5Z7wmPAv :=  /*line :1*/ context.Cu9yRv(yESLyde9LHhT)
	defer  /*line :1*/ qncl5Z7wmPAv()
	go  /*line :1*/ jdkVQwN_sasp(uxSavJI, true)

	aNcP6OMexa :=  /*line :1*/ time.Dyu6iYU( /*line :1*/ dKLkhP.dAexfVouH_6())
	defer  /*line :1*/ aNcP6OMexa.Stop()

	for {
		select {
		case <-aNcP6OMexa.N0KrfNaHg:
			uNXNoF, nXrhE7 :=  /*line :1*/ context.Cu9yRv(yESLyde9LHhT)
			defer  /*line :1*/ nXrhE7()
			go  /*line :1*/ jdkVQwN_sasp(uNXNoF, false)

		case pMHEP5J := <-iaDi__SK6bka:
			if pMHEP5J.error == nil {
				return pMHEP5J.UJYB3aCQg, nil
			}
			if pMHEP5J.bNfeI48 {
				efLai2h8_ = pMHEP5J
			} else {
				hmIvSNpIulg = pMHEP5J
			}
			if efLai2h8_.puOrMIlI8e && hmIvSNpIulg.puOrMIlI8e {
				return nil, efLai2h8_.error
			}
			if pMHEP5J.bNfeI48 &&  /*line :1*/ aNcP6OMexa.Stop() {

				 /*line :1*/ aNcP6OMexa.Reset(0)
			}
		}
	}
}



func (dKLkhP *vTALwD6Cw) uGDGFRHjpl0(yESLyde9LHhT context.MBCyA6, gPx_2lY2T nTzeZ5mf) (UJYB3aCQg, error) {
	var qSzdt0s24k error	

	for eaAUaB7Z, ioVu8c := range gPx_2lY2T {
		select {
		case <- /*line :1*/ yESLyde9LHhT.Done():
			return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: dKLkhP.cyZLILkT, F7g5pQacM05: dKLkhP.MDCJO9payv, JkQ9XFJbp: ioVu8c, XkWH4CYwNmvD:  /*line :1*/ gCG7rudfaf( /*line :1*/ yESLyde9LHhT.Err())}
		default:
		}

		xFCuNQ := yESLyde9LHhT
		if cSxcyrrgag, gNHzj0q :=  /*line :1*/ yESLyde9LHhT.Deadline(); gNHzj0q {
			hTeQHyBCV12i, h_ljk48Bm :=  /*line :1*/ hTeQHyBCV12i( /*line :1*/ time.Pc_35oTY(), cSxcyrrgag,  /*line :1*/ len(gPx_2lY2T)-eaAUaB7Z)
			if h_ljk48Bm != nil {

				if qSzdt0s24k == nil {
					qSzdt0s24k = &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: dKLkhP.cyZLILkT, F7g5pQacM05: dKLkhP.MDCJO9payv, JkQ9XFJbp: ioVu8c, XkWH4CYwNmvD: h_ljk48Bm}
				}
				break
			}
			if  /*line :1*/ hTeQHyBCV12i.Before(cSxcyrrgag) {
				var yDIc25 context.Gd7QRmF
				xFCuNQ, yDIc25 =  /*line :1*/ context.ZF8aTuau(yESLyde9LHhT, hTeQHyBCV12i)
				defer  /*line :1*/ yDIc25()
			}
		}

		hl8w2lFX, h_ljk48Bm :=  /*line :1*/ dKLkhP.g_BOsNOjG9N(xFCuNQ, ioVu8c)
		if h_ljk48Bm == nil {
			return hl8w2lFX, nil
		}
		if qSzdt0s24k == nil {
			qSzdt0s24k = h_ljk48Bm
		}
	}

	if qSzdt0s24k == nil {
		qSzdt0s24k = &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: dKLkhP.cyZLILkT, F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: tccDv0yXpU}
	}
	return nil, qSzdt0s24k
}



func (dKLkhP *vTALwD6Cw) g_BOsNOjG9N(yESLyde9LHhT context.MBCyA6, ioVu8c EqbMXsaU1lE) (hl8w2lFX UJYB3aCQg, h_ljk48Bm error) {
	jHCsMONOrYFN, _ :=  /*line :1*/ yESLyde9LHhT.Value(nettrace.YRC9V_{}).(*nettrace.UTp90oOoUInw)
	if jHCsMONOrYFN != nil {
		jM_f7zogPys :=  /*line :1*/ ioVu8c.String()
		if jHCsMONOrYFN.CVZFbB4EJX != nil {
			 /*line :1*/ jHCsMONOrYFN.CVZFbB4EJX(dKLkhP.cyZLILkT, jM_f7zogPys)
		}
		if jHCsMONOrYFN.EAYlwNxwXa != nil {
			defer func() {  /*line :1*/ jHCsMONOrYFN.EAYlwNxwXa(dKLkhP.cyZLILkT, jM_f7zogPys, h_ljk48Bm) }()
		}
	}
	ayKvoOPa := dKLkhP.MDCJO9payv
	switch ioVu8c := ioVu8c.(type) {
	case *HiOzqOVjN1P:
		ayKvoOPa, _ := ayKvoOPa.(*HiOzqOVjN1P)
		if  /*line :1*/ dKLkhP.MultipathTCP() {
			hl8w2lFX, h_ljk48Bm =  /*line :1*/ dKLkhP.znTcO556(yESLyde9LHhT, ayKvoOPa, ioVu8c)
		} else {
			hl8w2lFX, h_ljk48Bm =  /*line :1*/ dKLkhP.a9XNPFb(yESLyde9LHhT, ayKvoOPa, ioVu8c)
		}
	case *ZaffanpNx4:
		ayKvoOPa, _ := ayKvoOPa.(*ZaffanpNx4)
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ dKLkhP.ubPg4ZjrNYRu(yESLyde9LHhT, ayKvoOPa, ioVu8c)
	case *FZJphYv9hV:
		ayKvoOPa, _ := ayKvoOPa.(*FZJphYv9hV)
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ dKLkhP.xQAA5av(yESLyde9LHhT, ayKvoOPa, ioVu8c)
	case *E5nPmTZQM:
		ayKvoOPa, _ := ayKvoOPa.(*E5nPmTZQM)
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ dKLkhP.g5IHkyH(yESLyde9LHhT, ayKvoOPa, ioVu8c)
	default:
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: dKLkhP.cyZLILkT, F7g5pQacM05: ayKvoOPa, JkQ9XFJbp: ioVu8c, XkWH4CYwNmvD: &DAWLIQHc{Z68v0y: "unexpected address type", BCCnAgFxG: dKLkhP.r0U09j}}
	}
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "dial", CiQ5sBtmrVnf: dKLkhP.cyZLILkT, F7g5pQacM05: ayKvoOPa, JkQ9XFJbp: ioVu8c, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}


type MC4oQv struct {
	
	
	
	
	
	
	OJGgvEIn7No	func(vsbiWLn7reX, fwV_ln2dl string, hl8w2lFX syscall.CVnV1i) error

	
	
	
	
	
	
	Wlhapadf	time.GKMwTGxQa0S

	
	
	
	aLSatM4E4	am59Br2lmBKa
}





func (i8Iakd *MC4oQv) MultipathTCP() bool {
	return  /*line :1*/ i8Iakd.aLSatM4E4.cCaQG2_M1()
}







func (i8Iakd *MC4oQv) SetMultipathTCP(mFaToFtKuuHJ bool) {
	 /*line :1*/ i8Iakd.aLSatM4E4.hVlQmKILF2(mFaToFtKuuHJ)
}





func (i8Iakd *MC4oQv) Listen(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string) (Iy6rBQUJ9y7, error) {
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ Ic4wtIY.kq9KkFaU1B(yESLyde9LHhT, "listen", vsbiWLn7reX, fwV_ln2dl, nil)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	bw2hLt := &tYwm0DwfSTEv{
		MC4oQv:	*i8Iakd,
		aNAwQMxeTY:	vsbiWLn7reX,
		rxZTzA:	fwV_ln2dl,
	}
	var v3upVm Iy6rBQUJ9y7
	ayKvoOPa :=  /*line :1*/ md4QSRkO.g_D_9T(rnth8PRZsE2C)
	switch ayKvoOPa := ayKvoOPa.(type) {
	case *HiOzqOVjN1P:
		if  /*line :1*/ bw2hLt.MultipathTCP() {
			v3upVm, h_ljk48Bm =  /*line :1*/ bw2hLt.dscV09iVL(yESLyde9LHhT, ayKvoOPa)
		} else {
			v3upVm, h_ljk48Bm =  /*line :1*/ bw2hLt.bwWMnaxr(yESLyde9LHhT, ayKvoOPa)
		}
	case *E5nPmTZQM:
		v3upVm, h_ljk48Bm =  /*line :1*/ bw2hLt.caN1IhSmUV(yESLyde9LHhT, ayKvoOPa)
	default:
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: bw2hLt.aNAwQMxeTY, F7g5pQacM05: nil, JkQ9XFJbp: ayKvoOPa, XkWH4CYwNmvD: &DAWLIQHc{Z68v0y: "unexpected address type", BCCnAgFxG: fwV_ln2dl}}
	}
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: bw2hLt.aNAwQMxeTY, F7g5pQacM05: nil, JkQ9XFJbp: ayKvoOPa, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return v3upVm, nil
}





func (i8Iakd *MC4oQv) ListenPacket(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string) (KlDeyiWNu, error) {
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ Ic4wtIY.kq9KkFaU1B(yESLyde9LHhT, "listen", vsbiWLn7reX, fwV_ln2dl, nil)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: vsbiWLn7reX, F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	bw2hLt := &tYwm0DwfSTEv{
		MC4oQv:	*i8Iakd,
		aNAwQMxeTY:	vsbiWLn7reX,
		rxZTzA:	fwV_ln2dl,
	}
	var hl8w2lFX KlDeyiWNu
	ayKvoOPa :=  /*line :1*/ md4QSRkO.g_D_9T(rnth8PRZsE2C)
	switch ayKvoOPa := ayKvoOPa.(type) {
	case *ZaffanpNx4:
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ bw2hLt.eg0RUV(yESLyde9LHhT, ayKvoOPa)
	case *FZJphYv9hV:
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ bw2hLt.hq2iTKku3u(yESLyde9LHhT, ayKvoOPa)
	case *E5nPmTZQM:
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ bw2hLt.r5LehBZ4(yESLyde9LHhT, ayKvoOPa)
	default:
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: bw2hLt.aNAwQMxeTY, F7g5pQacM05: nil, JkQ9XFJbp: ayKvoOPa, XkWH4CYwNmvD: &DAWLIQHc{Z68v0y: "unexpected address type", BCCnAgFxG: fwV_ln2dl}}
	}
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "listen", CiQ5sBtmrVnf: bw2hLt.aNAwQMxeTY, F7g5pQacM05: nil, JkQ9XFJbp: ayKvoOPa, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return hl8w2lFX, nil
}


type tYwm0DwfSTEv struct {
	MC4oQv
	aNAwQMxeTY, rxZTzA	string
}






















func C6_kCWO8Km(vsbiWLn7reX, fwV_ln2dl string) (Iy6rBQUJ9y7, error) {
	var i8Iakd MC4oQv
	return  /*line :1*/ i8Iakd.Listen( /*line :1*/ context.GEcgQP5fzA(), vsbiWLn7reX, fwV_ln2dl)
}


























func G0aWAq9gQRj(vsbiWLn7reX, fwV_ln2dl string) (KlDeyiWNu, error) {
	var i8Iakd MC4oQv
	return  /*line :1*/ i8Iakd.ListenPacket( /*line :1*/ context.GEcgQP5fzA(), vsbiWLn7reX, fwV_ln2dl)
}
