//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	windows "LdptURlN"
	os "hPMrClpC"
	"runtime"
	syscall "bUKeamF"
	time "fRAfQd_"
	"unsafe"
)




const cgoAvailable = true

const (
	_WSAHOST_NOT_FOUND	=  /*line :1*/ syscall.O9Mga3(11001)
	_WSATRY_AGAIN	=  /*line :1*/ syscall.O9Mga3(11002)
)

func z5iHPX1(yppl074 string, h_ljk48Bm error) error {
	switch h_ljk48Bm {
	case _WSAHOST_NOT_FOUND:
		return aamCgdkZikV
	}
	return  /*line :1*/ os.BTaHHl(yppl074, h_ljk48Bm)
}

func h7pec4ADeQWH(tahgc783Bc string) (iaebRs5X3 int, h_ljk48Bm error) {
	fMPVz2iGiyq, h_ljk48Bm :=  /*line :1*/ syscall.IO3AcKmy7(tahgc783Bc)
	if h_ljk48Bm != nil {
		return 0,  /*line :1*/ z5iHPX1("getprotobyname", h_ljk48Bm)
	}
	return  /*line :1*/ int(fMPVz2iGiyq.EoKNay), nil
}


func heQcNLGEPBFn(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) (int, error) {
	
	
	type w1SqzJRu1Q struct {
		xh6tIYGNM	int
		o8DedArH6G	error
	}
	rvEj5N7 :=  /*line :1*/ make(chan w1SqzJRu1Q)
	go func() {
		 /*line :1*/ a0ivyEgPyL()
		defer  /*line :1*/ g4VDk4yZiGdR()
		 /*line :1*/ runtime.LockOSThread()
		defer  /*line :1*/ runtime.UnlockOSThread()
		iaebRs5X3, h_ljk48Bm :=  /*line :1*/ h7pec4ADeQWH(tahgc783Bc)
		select {
		case rvEj5N7 <- w1SqzJRu1Q{xh6tIYGNM: iaebRs5X3, o8DedArH6G: h_ljk48Bm}:
		case <- /*line :1*/ yESLyde9LHhT.Done():
		}
	}()
	select {
	case yxhs4Z := <-rvEj5N7:
		if yxhs4Z.o8DedArH6G != nil {
			if iaebRs5X3, h_ljk48Bm :=  /*line :1*/ lOVQ14ZJh(tahgc783Bc); h_ljk48Bm == nil {
				return iaebRs5X3, nil
			}

			twcnJnC := &SoatTK0{PY3bIIR:  /*line :1*/ yxhs4Z.o8DedArH6G.Error(), FIvHQdTCAg: tahgc783Bc}
			if yxhs4Z.o8DedArH6G == aamCgdkZikV {
				twcnJnC.GM17caVp1uW = true
			}
			yxhs4Z.o8DedArH6G = twcnJnC
		}
		return yxhs4Z.xh6tIYGNM, yxhs4Z.o8DedArH6G
	case <- /*line :1*/ yESLyde9LHhT.Done():
		return 0,  /*line :1*/ gCG7rudfaf( /*line :1*/ yESLyde9LHhT.Err())
	}
}

func (yxhs4Z *W8Q2tfHAD) c4c5ngLuoZKl(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]string, error) {
	ecxaiU5w1ARk, h_ljk48Bm :=  /*line :1*/ yxhs4Z.uPMUqJmf(yESLyde9LHhT, "ip", tahgc783Bc)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	md4QSRkO :=  /*line :1*/ make([]string, 0,  /*line :1*/ len(ecxaiU5w1ARk))
	for _, kQ8_UEhxU := range ecxaiU5w1ARk {
		md4QSRkO =  /*line :1*/ append(md4QSRkO,  /*line :1*/ kQ8_UEhxU.String())
	}
	return md4QSRkO, nil
}




func (yxhs4Z *W8Q2tfHAD) gMUfoR() bool {
	q6Glr4X1u :=  /*line :1*/ yAKpTR()
	fM1jn7loPH, _ :=  /*line :1*/ q6Glr4X1u.mNRpaDPvaD(yxhs4Z, "")
	return fM1jn7loPH != hostLookupCgo
}

func (yxhs4Z *W8Q2tfHAD) uPMUqJmf(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, tahgc783Bc string) ([]FZJphYv9hV, error) {
	if  /*line :1*/ yxhs4Z.gMUfoR() {
		return  /*line :1*/ yxhs4Z.yIyaYwBXWU(yESLyde9LHhT, vsbiWLn7reX, tahgc783Bc)
	}

	var nqsgnnaf int32 = syscall.AF_UNSPEC
	switch  /*line :1*/ qqb9mdSqc(vsbiWLn7reX) {
	case '4':
		nqsgnnaf = syscall.AF_INET
	case '6':
		nqsgnnaf = syscall.AF_INET6
	}

	bnmE1OmL0R5 := func() ([]FZJphYv9hV, error) {
		 /*line :1*/ a0ivyEgPyL()
		defer  /*line :1*/ g4VDk4yZiGdR()
		dqd7aAf := syscall.HnpJdSpZM__0{
			WLUX0U:	nqsgnnaf,
			C_dz1S:	syscall.SOCK_STREAM,
			Ax02FCIziFnj:	syscall.IPPROTO_IP,
		}
		var w1SqzJRu1Q *syscall.HnpJdSpZM__0
		ylz4zh, h_ljk48Bm :=  /*line :1*/ syscall.GcOmFfsibES(tahgc783Bc)
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{FIvHQdTCAg: tahgc783Bc, PY3bIIR:  /*line :1*/ h_ljk48Bm.Error()}
		}

		vTBPTe :=  /*line :1*/ mT_rTj1r8a()
		vz3rO6LC5 :=  /*line :1*/ time.Pc_35oTY()

		var wI0aaz error
		for eaAUaB7Z := 0; eaAUaB7Z < vTBPTe.yb6dJ6U; eaAUaB7Z++ {
			wI0aaz =  /*line :1*/ syscall.Nsyrt6(ylz4zh, nil, &dqd7aAf, &w1SqzJRu1Q)
			if wI0aaz == nil || wI0aaz != _WSATRY_AGAIN ||  /*line :1*/ time.Qs3VCKXXED9(vz3rO6LC5) > vTBPTe.c1XqMJ {
				break
			}
		}
		if wI0aaz != nil {
			h_ljk48Bm :=  /*line :1*/ z5iHPX1("getaddrinfow", wI0aaz)
			twcnJnC := &SoatTK0{PY3bIIR:  /*line :1*/ h_ljk48Bm.Error(), FIvHQdTCAg: tahgc783Bc}
			if h_ljk48Bm == aamCgdkZikV {
				twcnJnC.GM17caVp1uW = true
			}
			return nil, twcnJnC
		}
		defer  /*line :1*/ syscall.BmTZjeIGcn(w1SqzJRu1Q)
		md4QSRkO :=  /*line :1*/ make([]FZJphYv9hV, 0, 5)
		for ; w1SqzJRu1Q != nil; w1SqzJRu1Q = w1SqzJRu1Q.Aye498L2D {
			qxwkC3VxG0 :=  /*line :1*/ unsafe.Pointer(w1SqzJRu1Q.JhafcUtas1O)
			switch w1SqzJRu1Q.WLUX0U {
			case syscall.AF_INET:
				uI7LZDHm6 := (* /*line :1*/ syscall.NXUDVMl2)(qxwkC3VxG0).AGDgR2QSLqW
				md4QSRkO =  /*line :1*/ append(md4QSRkO, FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ sNTdcemFI(uI7LZDHm6[:])})
			case syscall.AF_INET6:
				uI7LZDHm6 := (* /*line :1*/ syscall.H7tgUmcLUW)(qxwkC3VxG0).G1aNyX7
				dJhOHOEZlQAA :=  /*line :1*/ x4vxFT.tahgc783Bc( /*line :1*/ int((* /*line :1*/ syscall.H7tgUmcLUW)(qxwkC3VxG0).Jl_oKnHTc))
				md4QSRkO =  /*line :1*/ append(md4QSRkO, FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ sNTdcemFI(uI7LZDHm6[:]), Cyg5M2vYIby: dJhOHOEZlQAA})
			default:
				return nil, &SoatTK0{PY3bIIR:  /*line :1*/ syscall.EWINDOWS.Error(), FIvHQdTCAg: tahgc783Bc}
			}
		}
		return md4QSRkO, nil
	}

	type prHUfwjRzJ struct {
		jUkcbAy3I	[]FZJphYv9hV
		bdn3TOvJQ	error
	}

	var rvEj5N7 chan prHUfwjRzJ
	if  /*line :1*/ yESLyde9LHhT.Err() == nil {
		rvEj5N7 =  /*line :1*/ make(chan prHUfwjRzJ, 1)
		go func() {
			 /*line :1*/ qxwkC3VxG0, h_ljk48Bm :=  /*line :1*/ bnmE1OmL0R5()
			rvEj5N7 <- prHUfwjRzJ{jUkcbAy3I: qxwkC3VxG0, bdn3TOvJQ: h_ljk48Bm}
		}()
	}

	select {
	case yxhs4Z := <-rvEj5N7:
		return yxhs4Z.jUkcbAy3I, yxhs4Z.bdn3TOvJQ
	case <- /*line :1*/ yESLyde9LHhT.Done():

		return nil, &SoatTK0{
			FIvHQdTCAg:	tahgc783Bc,
			PY3bIIR:	 /*line :1*/ yESLyde9LHhT.Err().Error(),
			JYeYrD:	 /*line :1*/ yESLyde9LHhT.Err() == context.Z201U2N,
		}
	}
}

func (yxhs4Z *W8Q2tfHAD) t_OMUaXanKx(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, zvX72x8Ih7C string) (int, error) {
	if  /*line :1*/ yxhs4Z.gMUfoR() {
		return  /*line :1*/ dTwMnMHVj(vsbiWLn7reX, zvX72x8Ih7C)
	}

	 /*line :1*/ a0ivyEgPyL()
	defer  /*line :1*/ g4VDk4yZiGdR()
	var sd96BRq int32
	switch vsbiWLn7reX {
	case "tcp4", "tcp6":
		sd96BRq = syscall.SOCK_STREAM
	case "udp4", "udp6":
		sd96BRq = syscall.SOCK_DGRAM
	}
	dqd7aAf := syscall.HnpJdSpZM__0{
		WLUX0U:	syscall.AF_UNSPEC,
		C_dz1S:	sd96BRq,
		Ax02FCIziFnj:	syscall.IPPROTO_IP,
	}
	var w1SqzJRu1Q *syscall.HnpJdSpZM__0
	wI0aaz :=  /*line :1*/ syscall.Nsyrt6(nil,  /*line :1*/ syscall.EtPVNA(zvX72x8Ih7C), &dqd7aAf, &w1SqzJRu1Q)
	if wI0aaz != nil {
		if mzUgFPgljgC, h_ljk48Bm :=  /*line :1*/ dTwMnMHVj(vsbiWLn7reX, zvX72x8Ih7C); h_ljk48Bm == nil {
			return mzUgFPgljgC, nil
		}
		h_ljk48Bm :=  /*line :1*/ z5iHPX1("getaddrinfow", wI0aaz)
		twcnJnC := &SoatTK0{PY3bIIR:  /*line :1*/ h_ljk48Bm.Error(), FIvHQdTCAg: vsbiWLn7reX + "/" + zvX72x8Ih7C}
		if h_ljk48Bm == aamCgdkZikV {
			twcnJnC.GM17caVp1uW = true
		}
		return 0, twcnJnC
	}
	defer  /*line :1*/ syscall.BmTZjeIGcn(w1SqzJRu1Q)
	if w1SqzJRu1Q == nil {
		return 0, &SoatTK0{PY3bIIR:  /*line :1*/ syscall.EINVAL.Error(), FIvHQdTCAg: vsbiWLn7reX + "/" + zvX72x8Ih7C}
	}
	qxwkC3VxG0 :=  /*line :1*/ unsafe.Pointer(w1SqzJRu1Q.JhafcUtas1O)
	switch w1SqzJRu1Q.WLUX0U {
	case syscall.AF_INET:
		uI7LZDHm6 := (* /*line :1*/ syscall.NXUDVMl2)(qxwkC3VxG0)
		return  /*line :1*/ int( /*line :1*/ syscall.EOoAMh(uI7LZDHm6.VqKTvvibG)), nil
	case syscall.AF_INET6:
		uI7LZDHm6 := (* /*line :1*/ syscall.H7tgUmcLUW)(qxwkC3VxG0)
		return  /*line :1*/ int( /*line :1*/ syscall.EOoAMh(uI7LZDHm6.D_H8WxCL7)), nil
	}
	return 0, &SoatTK0{PY3bIIR:  /*line :1*/ syscall.EINVAL.Error(), FIvHQdTCAg: vsbiWLn7reX + "/" + zvX72x8Ih7C}
}

func (yxhs4Z *W8Q2tfHAD) eLhY_2IAd(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) (string, error) {
	if fM1jn7loPH, q6Glr4X1u :=  /*line :1*/ yAKpTR().mNRpaDPvaD(yxhs4Z, ""); fM1jn7loPH != hostLookupCgo {
		return  /*line :1*/ yxhs4Z.kK7DIGIx(yESLyde9LHhT, tahgc783Bc, fM1jn7loPH, q6Glr4X1u)
	}

	 /*line :1*/ a0ivyEgPyL()
	defer  /*line :1*/ g4VDk4yZiGdR()
	var dyDxHas *syscall.OB0HzAtGUg
	wI0aaz :=  /*line :1*/ syscall.BE314j4hm(tahgc783Bc, syscall.DNS_TYPE_CNAME, 0, nil, &dyDxHas, nil)

	if ay5Y311yrq, d30HIwxht := wI0aaz.(syscall.O9Mga3); d30HIwxht && ay5Y311yrq == syscall.DNS_INFO_NO_RECORDS {

		return  /*line :1*/ ztv3vbC8VS(tahgc783Bc), nil
	}
	if wI0aaz != nil {
		return "", &SoatTK0{PY3bIIR:  /*line :1*/ z5iHPX1("dnsquery", wI0aaz).Error(), FIvHQdTCAg: tahgc783Bc}
	}
	defer  /*line :1*/ syscall.DuNM27(dyDxHas, 1)

	uK4KO2Al7L :=  /*line :1*/ nAIVY_m( /*line :1*/ syscall.EtPVNA(tahgc783Bc), dyDxHas)
	jcluQhsSxoD :=  /*line :1*/ windows.DOR2gxA_7npQ(uK4KO2Al7L)
	return  /*line :1*/ ztv3vbC8VS(jcluQhsSxoD), nil
}

func (yxhs4Z *W8Q2tfHAD) fe_7xlr(yESLyde9LHhT context.MBCyA6, zvX72x8Ih7C, iaebRs5X3, tahgc783Bc string) (string, []*LRAj4ahDtDg, error) {
	if  /*line :1*/ yxhs4Z.gMUfoR() {
		return  /*line :1*/ yxhs4Z.dBGgrS2larpJ(yESLyde9LHhT, zvX72x8Ih7C, iaebRs5X3, tahgc783Bc)
	}

	 /*line :1*/ a0ivyEgPyL()
	defer  /*line :1*/ g4VDk4yZiGdR()
	var qu6sFD7 string
	if zvX72x8Ih7C == "" && iaebRs5X3 == "" {
		qu6sFD7 = tahgc783Bc
	} else {
		qu6sFD7 = "_" + zvX72x8Ih7C + "._" + iaebRs5X3 + "." + tahgc783Bc
	}
	var dyDxHas *syscall.OB0HzAtGUg
	wI0aaz :=  /*line :1*/ syscall.BE314j4hm(qu6sFD7, syscall.DNS_TYPE_SRV, 0, nil, &dyDxHas, nil)
	if wI0aaz != nil {
		return "", nil, &SoatTK0{PY3bIIR:  /*line :1*/ z5iHPX1("dnsquery", wI0aaz).Error(), FIvHQdTCAg: qu6sFD7}
	}
	defer  /*line :1*/ syscall.DuNM27(dyDxHas, 1)

	yIBOJND0k5 :=  /*line :1*/ make([]*LRAj4ahDtDg, 0, 10)
	for _, fMPVz2iGiyq := range  /*line :1*/ l81udsGH(dyDxHas, syscall.DNS_TYPE_SRV, qu6sFD7) {
		ljsCSk := (* /*line :1*/ syscall.NL1jD3M7P)( /*line :1*/ unsafe.Pointer(&fMPVz2iGiyq.TnF1oOW[0]))
		yIBOJND0k5 =  /*line :1*/ append(yIBOJND0k5, &LRAj4ahDtDg{ /*line :1*/ ztv3vbC8VS( /*line :1*/ syscall.AODVXp8NM3sd((*[256] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(ljsCSk.JbYsAIP))[:])), ljsCSk.NUAVLSp5MDik, ljsCSk.FkTrEz, ljsCSk.INJblcAM})
	}
	 /*line :1*/ a_iF1ZK0(yIBOJND0k5).wYlOfm()
	return  /*line :1*/ ztv3vbC8VS(qu6sFD7), yIBOJND0k5, nil
}

func (yxhs4Z *W8Q2tfHAD) zr9BDQtsDI(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]*McZ_wBAueqxz, error) {
	if  /*line :1*/ yxhs4Z.gMUfoR() {
		return  /*line :1*/ yxhs4Z.yhVX7Ebw0(yESLyde9LHhT, tahgc783Bc)
	}

	 /*line :1*/ a0ivyEgPyL()
	defer  /*line :1*/ g4VDk4yZiGdR()
	var dyDxHas *syscall.OB0HzAtGUg
	wI0aaz :=  /*line :1*/ syscall.BE314j4hm(tahgc783Bc, syscall.DNS_TYPE_MX, 0, nil, &dyDxHas, nil)
	if wI0aaz != nil {
		return nil, &SoatTK0{PY3bIIR:  /*line :1*/ z5iHPX1("dnsquery", wI0aaz).Error(), FIvHQdTCAg: tahgc783Bc}
	}
	defer  /*line :1*/ syscall.DuNM27(dyDxHas, 1)

	d8Wpo4t :=  /*line :1*/ make([]*McZ_wBAueqxz, 0, 10)
	for _, fMPVz2iGiyq := range  /*line :1*/ l81udsGH(dyDxHas, syscall.DNS_TYPE_MX, tahgc783Bc) {
		ljsCSk := (* /*line :1*/ syscall.AGNDZEpVSE)( /*line :1*/ unsafe.Pointer(&fMPVz2iGiyq.TnF1oOW[0]))
		d8Wpo4t =  /*line :1*/ append(d8Wpo4t, &McZ_wBAueqxz{ /*line :1*/ ztv3vbC8VS( /*line :1*/ windows.DOR2gxA_7npQ(ljsCSk.YNeBpnBS45M)), ljsCSk.ZT51ejkkMgU})
	}
	 /*line :1*/ rcgRLRupL(d8Wpo4t).wYlOfm()
	return d8Wpo4t, nil
}

func (yxhs4Z *W8Q2tfHAD) zoOPu7z1(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]*AEmYkM1, error) {
	if  /*line :1*/ yxhs4Z.gMUfoR() {
		return  /*line :1*/ yxhs4Z.aAJTpSPuXHCL(yESLyde9LHhT, tahgc783Bc)
	}

	 /*line :1*/ a0ivyEgPyL()
	defer  /*line :1*/ g4VDk4yZiGdR()
	var dyDxHas *syscall.OB0HzAtGUg
	wI0aaz :=  /*line :1*/ syscall.BE314j4hm(tahgc783Bc, syscall.DNS_TYPE_NS, 0, nil, &dyDxHas, nil)
	if wI0aaz != nil {
		return nil, &SoatTK0{PY3bIIR:  /*line :1*/ z5iHPX1("dnsquery", wI0aaz).Error(), FIvHQdTCAg: tahgc783Bc}
	}
	defer  /*line :1*/ syscall.DuNM27(dyDxHas, 1)

	cjFUVvWN :=  /*line :1*/ make([]*AEmYkM1, 0, 10)
	for _, fMPVz2iGiyq := range  /*line :1*/ l81udsGH(dyDxHas, syscall.DNS_TYPE_NS, tahgc783Bc) {
		ljsCSk := (* /*line :1*/ syscall.RILagwbjo)( /*line :1*/ unsafe.Pointer(&fMPVz2iGiyq.TnF1oOW[0]))
		cjFUVvWN =  /*line :1*/ append(cjFUVvWN, &AEmYkM1{ /*line :1*/ ztv3vbC8VS( /*line :1*/ syscall.AODVXp8NM3sd((*[256] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(ljsCSk.CVn0dr1Jky6Z))[:]))})
	}
	return cjFUVvWN, nil
}

func (yxhs4Z *W8Q2tfHAD) hxfm94xv(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]string, error) {
	if  /*line :1*/ yxhs4Z.gMUfoR() {
		return  /*line :1*/ yxhs4Z.qR_lwcoTokk(yESLyde9LHhT, tahgc783Bc)
	}

	 /*line :1*/ a0ivyEgPyL()
	defer  /*line :1*/ g4VDk4yZiGdR()
	var dyDxHas *syscall.OB0HzAtGUg
	wI0aaz :=  /*line :1*/ syscall.BE314j4hm(tahgc783Bc, syscall.DNS_TYPE_TEXT, 0, nil, &dyDxHas, nil)
	if wI0aaz != nil {
		return nil, &SoatTK0{PY3bIIR:  /*line :1*/ z5iHPX1("dnsquery", wI0aaz).Error(), FIvHQdTCAg: tahgc783Bc}
	}
	defer  /*line :1*/ syscall.DuNM27(dyDxHas, 1)

	gl9eQhAnSrZ :=  /*line :1*/ make([]string, 0, 10)
	for _, fMPVz2iGiyq := range  /*line :1*/ l81udsGH(dyDxHas, syscall.DNS_TYPE_TEXT, tahgc783Bc) {
		ica44Q := (* /*line :1*/ syscall.ZaocUo291mr)( /*line :1*/ unsafe.Pointer(&fMPVz2iGiyq.TnF1oOW[0]))
		dRoFccaZ := ""
		for _, ljsCSk := range (*[1 << 10]* /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&(ica44Q.AfGeifce[0])))[:ica44Q.EKVQR85t9paB:ica44Q.EKVQR85t9paB] {
			dRoFccaZ +=  /*line :1*/ windows.DOR2gxA_7npQ(ljsCSk)
		}
		gl9eQhAnSrZ =  /*line :1*/ append(gl9eQhAnSrZ, dRoFccaZ)
	}
	return gl9eQhAnSrZ, nil
}

func (yxhs4Z *W8Q2tfHAD) tDIayKw(yESLyde9LHhT context.MBCyA6, qxwkC3VxG0 string) ([]string, error) {
	if fM1jn7loPH, q6Glr4X1u :=  /*line :1*/ yAKpTR().mNRpaDPvaD(yxhs4Z, ""); fM1jn7loPH != hostLookupCgo {
		return  /*line :1*/ yxhs4Z.nHwfF3h(yESLyde9LHhT, qxwkC3VxG0, fM1jn7loPH, q6Glr4X1u)
	}

	 /*line :1*/ a0ivyEgPyL()
	defer  /*line :1*/ g4VDk4yZiGdR()
	akvMHAa84xdt, h_ljk48Bm :=  /*line :1*/ sPSXZcPfQnVj(qxwkC3VxG0)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	var dyDxHas *syscall.OB0HzAtGUg
	wI0aaz :=  /*line :1*/ syscall.BE314j4hm(akvMHAa84xdt, syscall.DNS_TYPE_PTR, 0, nil, &dyDxHas, nil)
	if wI0aaz != nil {
		return nil, &SoatTK0{PY3bIIR:  /*line :1*/ z5iHPX1("dnsquery", wI0aaz).Error(), FIvHQdTCAg: qxwkC3VxG0}
	}
	defer  /*line :1*/ syscall.DuNM27(dyDxHas, 1)

	uvaup8zGu :=  /*line :1*/ make([]string, 0, 10)
	for _, fMPVz2iGiyq := range  /*line :1*/ l81udsGH(dyDxHas, syscall.DNS_TYPE_PTR, akvMHAa84xdt) {
		ljsCSk := (* /*line :1*/ syscall.RILagwbjo)( /*line :1*/ unsafe.Pointer(&fMPVz2iGiyq.TnF1oOW[0]))
		uvaup8zGu =  /*line :1*/ append(uvaup8zGu,  /*line :1*/ ztv3vbC8VS( /*line :1*/ windows.DOR2gxA_7npQ(ljsCSk.CVn0dr1Jky6Z)))
	}
	return uvaup8zGu, nil
}

const dnsSectionMask = 0x0003


func l81udsGH(yxhs4Z *syscall.OB0HzAtGUg, jbVjxZQUk2 uint16, tahgc783Bc string) []*syscall.OB0HzAtGUg {
	jcluQhsSxoD :=  /*line :1*/ syscall.EtPVNA(tahgc783Bc)
	if jbVjxZQUk2 != syscall.DNS_TYPE_CNAME {
		jcluQhsSxoD =  /*line :1*/ nAIVY_m(jcluQhsSxoD, yxhs4Z)
	}
	dyDxHas :=  /*line :1*/ make([]*syscall.OB0HzAtGUg, 0, 10)
	for fMPVz2iGiyq := yxhs4Z; fMPVz2iGiyq != nil; fMPVz2iGiyq = fMPVz2iGiyq.IBwV54lKvxs {

		if fMPVz2iGiyq.Dkc4C3R3KS&dnsSectionMask != syscall.DnsSectionAnswer && fMPVz2iGiyq.Dkc4C3R3KS&dnsSectionMask != syscall.DnsSectionQuestion {
			continue
		}
		if fMPVz2iGiyq.SKG0xmj0Cq != jbVjxZQUk2 {
			continue
		}
		if ! /*line :1*/ syscall.XpXlD6aa(jcluQhsSxoD, fMPVz2iGiyq.NHneOoxbcY) {
			continue
		}
		dyDxHas =  /*line :1*/ append(dyDxHas, fMPVz2iGiyq)
	}
	return dyDxHas
}


func nAIVY_m(tahgc783Bc *uint16, yxhs4Z *syscall.OB0HzAtGUg) *uint16 {

Cname:
	for nOIrjY := 0; nOIrjY < 10; nOIrjY++ {
		for fMPVz2iGiyq := yxhs4Z; fMPVz2iGiyq != nil; fMPVz2iGiyq = fMPVz2iGiyq.IBwV54lKvxs {
			if fMPVz2iGiyq.Dkc4C3R3KS&dnsSectionMask != syscall.DnsSectionAnswer {
				continue
			}
			if fMPVz2iGiyq.SKG0xmj0Cq != syscall.DNS_TYPE_CNAME {
				continue
			}
			if ! /*line :1*/ syscall.XpXlD6aa(tahgc783Bc, fMPVz2iGiyq.NHneOoxbcY) {
				continue
			}
			tahgc783Bc = (* /*line :1*/ syscall.RILagwbjo)( /*line :1*/ unsafe.Pointer(&yxhs4Z.TnF1oOW[0])).CVn0dr1Jky6Z
			continue Cname
		}
		break
	}
	return tahgc783Bc
}



func nxi0S_() int {
	return 500
}
