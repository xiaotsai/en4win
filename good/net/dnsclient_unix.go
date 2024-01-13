//line :1
//go:build !js

package gF9mZs2

import (
	context "fOYpb3F03EG"
	errors "iAHoxjmM"
	itoa "JkjxVS"
	io "xI9ai2djaJ2"
	os "hPMrClpC"
	"runtime"
	sync "sync"
	atomic "sync/atomic"
	time "fRAfQd_"

	dnsmessage "FeKim_cPgM7z"
)

const (
	useTCPOnly	= true
	useUDPOrTCP	= false

	maxDNSPacketSize	= 1232
)

var (
	ggfOmjBI	=  /*line :1*/ errors.Su6g6hRoi9X("lame referral")
	gqXCLp	=  /*line :1*/ errors.Su6g6hRoi9X("cannot unmarshal DNS message")
	p6JX85H_v	=  /*line :1*/ errors.Su6g6hRoi9X("cannot marshal DNS message")
	seOIPwE6LNoT	=  /*line :1*/ errors.Su6g6hRoi9X("server misbehaving")
	nyFMjrLYpj	=  /*line :1*/ errors.Su6g6hRoi9X("invalid DNS response")
	dMAIPd9zt	=  /*line :1*/ errors.Su6g6hRoi9X("no answer from DNS server")

	k3srm4	=  /*line :1*/ errors.Su6g6hRoi9X("server misbehaving")
)

func fvXyXf4T7EAl(fsa6C0Qdm dnsmessage.LNQ5aNrCORB, hEyH_at bool) (jujzF938Q uint16, oLkxTrt, xgpqdsrhPd []byte, h_ljk48Bm error) {
	jujzF938Q =  /*line :1*/ uint16( /*line :1*/ gOhSWMo9m7A())
	fIadEXIim6l :=  /*line :1*/ dnsmessage.Dt4jP7nnj6c8( /*line :1*/ make([]byte, 2, 514), dnsmessage.HlVtnUC{Xk_bZKgX: jujzF938Q, HXV_4sM2D: true, Ah7I1_6: hEyH_at})
	if h_ljk48Bm :=  /*line :1*/ fIadEXIim6l.StartQuestions(); h_ljk48Bm != nil {
		return 0, nil, nil, h_ljk48Bm
	}
	if h_ljk48Bm :=  /*line :1*/ fIadEXIim6l.Question(fsa6C0Qdm); h_ljk48Bm != nil {
		return 0, nil, nil, h_ljk48Bm
	}

	if h_ljk48Bm :=  /*line :1*/ fIadEXIim6l.StartAdditionals(); h_ljk48Bm != nil {
		return 0, nil, nil, h_ljk48Bm
	}
	var _42csxAa dnsmessage.RcKjQr_MA
	if h_ljk48Bm :=  /*line :1*/ _42csxAa.SetEDNS0(maxDNSPacketSize, dnsmessage.RCodeSuccess, false); h_ljk48Bm != nil {
		return 0, nil, nil, h_ljk48Bm
	}
	if h_ljk48Bm :=  /*line :1*/ fIadEXIim6l.OPTResource(_42csxAa, dnsmessage.PsSXLq1p{}); h_ljk48Bm != nil {
		return 0, nil, nil, h_ljk48Bm
	}

	xgpqdsrhPd, h_ljk48Bm =  /*line :1*/ fIadEXIim6l.Finish()
	if h_ljk48Bm != nil {
		return 0, nil, nil, h_ljk48Bm
	}
	oLkxTrt = xgpqdsrhPd[2:]
	v3upVm :=  /*line :1*/ len(xgpqdsrhPd) - 2
	xgpqdsrhPd[0] =  /*line :1*/ byte(v3upVm >> 8)
	xgpqdsrhPd[1] =  /*line :1*/ byte(v3upVm)
	return jujzF938Q, oLkxTrt, xgpqdsrhPd, nil
}

func drRUQKb(gdpn_b uint16, p8FiQG8 dnsmessage.LNQ5aNrCORB, sG6OZYPPDPXL dnsmessage.HlVtnUC, tcnrIMc71 dnsmessage.LNQ5aNrCORB) bool {
	if !sG6OZYPPDPXL.FlCGSzy1U7Gj {
		return false
	}
	if gdpn_b != sG6OZYPPDPXL.Xk_bZKgX {
		return false
	}
	if p8FiQG8.XsqCP7Wa65g != tcnrIMc71.XsqCP7Wa65g || p8FiQG8.GBiiAZ64kqf6 != tcnrIMc71.GBiiAZ64kqf6 || ! /*line :1*/ emSzQ2I(p8FiQG8.FLDYrkuX, tcnrIMc71.FLDYrkuX) {
		return false
	}
	return true
}

func dTiTYwidx(hl8w2lFX UJYB3aCQg, jujzF938Q uint16, cWXda5xN dnsmessage.LNQ5aNrCORB, fIadEXIim6l []byte) (dnsmessage.GShepB, dnsmessage.HlVtnUC, error) {
	if _, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.Write(fIadEXIim6l); h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, h_ljk48Bm
	}

	fIadEXIim6l =  /*line :1*/ make([]byte, maxDNSPacketSize)
	for {
		doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.Read(fIadEXIim6l)
		if h_ljk48Bm != nil {
			return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, h_ljk48Bm
		}
		var fMPVz2iGiyq dnsmessage.GShepB

		cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.Start(fIadEXIim6l[:doauF8Y])
		if h_ljk48Bm != nil {
			continue
		}
		fsa6C0Qdm, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.Question()
		if h_ljk48Bm != nil || ! /*line :1*/ drRUQKb(jujzF938Q, cWXda5xN, cBprzKT8, fsa6C0Qdm) {
			continue
		}
		return fMPVz2iGiyq, cBprzKT8, nil
	}
}

func hGEtvpfCjVC(hl8w2lFX UJYB3aCQg, jujzF938Q uint16, cWXda5xN dnsmessage.LNQ5aNrCORB, fIadEXIim6l []byte) (dnsmessage.GShepB, dnsmessage.HlVtnUC, error) {
	if _, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.Write(fIadEXIim6l); h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, h_ljk48Bm
	}

	fIadEXIim6l =  /*line :1*/ make([]byte, 1280)
	if _, h_ljk48Bm :=  /*line :1*/ io.G0zBgKg(hl8w2lFX, fIadEXIim6l[:2]); h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, h_ljk48Bm
	}
	v3upVm :=  /*line :1*/ int(fIadEXIim6l[0])<<8 |  /*line :1*/ int(fIadEXIim6l[1])
	if v3upVm >  /*line :1*/ len(fIadEXIim6l) {
		fIadEXIim6l =  /*line :1*/ make([]byte, v3upVm)
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ io.G0zBgKg(hl8w2lFX, fIadEXIim6l[:v3upVm])
	if h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, h_ljk48Bm
	}
	var fMPVz2iGiyq dnsmessage.GShepB
	cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.Start(fIadEXIim6l[:doauF8Y])
	if h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, gqXCLp
	}
	fsa6C0Qdm, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.Question()
	if h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, gqXCLp
	}
	if ! /*line :1*/ drRUQKb(jujzF938Q, cWXda5xN, cBprzKT8, fsa6C0Qdm) {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, nyFMjrLYpj
	}
	return fMPVz2iGiyq, cBprzKT8, nil
}

func (yxhs4Z *W8Q2tfHAD) rMsDNw(yESLyde9LHhT context.MBCyA6, c_eHMUdT string, fsa6C0Qdm dnsmessage.LNQ5aNrCORB, yusr33iR time.GKMwTGxQa0S, jxOC8yQT, hEyH_at bool) (dnsmessage.GShepB, dnsmessage.HlVtnUC, error) {
	fsa6C0Qdm.GBiiAZ64kqf6 = dnsmessage.ClassINET
	jujzF938Q, oLkxTrt, xgpqdsrhPd, h_ljk48Bm :=  /*line :1*/ fvXyXf4T7EAl(fsa6C0Qdm, hEyH_at)
	if h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, p6JX85H_v
	}
	var z6TKtFRGr []string
	if jxOC8yQT {
		z6TKtFRGr = []string{"tcp"}
	} else {
		z6TKtFRGr = []string{"udp", "tcp"}
	}
	for _, vsbiWLn7reX := range z6TKtFRGr {
		yESLyde9LHhT, yDIc25 :=  /*line :1*/ context.ZF8aTuau(yESLyde9LHhT,  /*line :1*/ time.Pc_35oTY().Add(yusr33iR))
		defer  /*line :1*/ yDIc25()

		hl8w2lFX, h_ljk48Bm :=  /*line :1*/ yxhs4Z.tdqObJ_tTokV(yESLyde9LHhT, vsbiWLn7reX, c_eHMUdT)
		if h_ljk48Bm != nil {
			return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, h_ljk48Bm
		}
		if ica44Q, d30HIwxht :=  /*line :1*/ yESLyde9LHhT.Deadline(); d30HIwxht && ! /*line :1*/ ica44Q.IsZero() {
			 /*line :1*/ hl8w2lFX.SetDeadline(ica44Q)
		}
		var fMPVz2iGiyq dnsmessage.GShepB
		var cBprzKT8 dnsmessage.HlVtnUC
		if _, d30HIwxht := hl8w2lFX.(KlDeyiWNu); d30HIwxht {
			fMPVz2iGiyq, cBprzKT8, h_ljk48Bm =  /*line :1*/ dTiTYwidx(hl8w2lFX, jujzF938Q, fsa6C0Qdm, oLkxTrt)
		} else {
			fMPVz2iGiyq, cBprzKT8, h_ljk48Bm =  /*line :1*/ hGEtvpfCjVC(hl8w2lFX, jujzF938Q, fsa6C0Qdm, xgpqdsrhPd)
		}
		 /*line :1*/ hl8w2lFX.Close()
		if h_ljk48Bm != nil {
			return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{},  /*line :1*/ gCG7rudfaf(h_ljk48Bm)
		}
		if h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SkipQuestion(); h_ljk48Bm != dnsmessage.BbQhjXWx {
			return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, nyFMjrLYpj
		}
		if cBprzKT8.LeyP1H {
			continue
		}
		return fMPVz2iGiyq, cBprzKT8, nil
	}
	return dnsmessage.GShepB{}, dnsmessage.HlVtnUC{}, dMAIPd9zt
}

func siVD3Ya(fMPVz2iGiyq *dnsmessage.GShepB, cBprzKT8 dnsmessage.HlVtnUC) error {
	if cBprzKT8.DAtOvNfx7 == dnsmessage.RCodeNameError {
		return aamCgdkZikV
	}

	_, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.AnswerHeader()
	if h_ljk48Bm != nil && h_ljk48Bm != dnsmessage.BbQhjXWx {
		return gqXCLp
	}

	if cBprzKT8.DAtOvNfx7 == dnsmessage.RCodeSuccess && !cBprzKT8.AgqoiYIElD && !cBprzKT8.Xu3NFYuAq && h_ljk48Bm == dnsmessage.BbQhjXWx {
		return ggfOmjBI
	}

	if cBprzKT8.DAtOvNfx7 != dnsmessage.RCodeSuccess && cBprzKT8.DAtOvNfx7 != dnsmessage.RCodeNameError {

		if cBprzKT8.DAtOvNfx7 == dnsmessage.RCodeServerFailure {
			return k3srm4
		}
		return seOIPwE6LNoT
	}

	return nil
}

func ouo_Rd(fMPVz2iGiyq *dnsmessage.GShepB, vBrGtQW dnsmessage.Ze6KllqJw) error {
	for {
		cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.AnswerHeader()
		if h_ljk48Bm == dnsmessage.BbQhjXWx {
			return aamCgdkZikV
		}
		if h_ljk48Bm != nil {
			return gqXCLp
		}
		if cBprzKT8.BxqvzBsK == vBrGtQW {
			return nil
		}
		if h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SkipAnswer(); h_ljk48Bm != nil {
			return gqXCLp
		}
	}
}

func (yxhs4Z *W8Q2tfHAD) cFpiL14zqMrj(yESLyde9LHhT context.MBCyA6, fFmWj818I *ivgt6m77Buqr, tahgc783Bc string, vBrGtQW dnsmessage.Ze6KllqJw) (dnsmessage.GShepB, string, error) {
	var foP_Mw6F0ef error
	rQxDMS :=  /*line :1*/ fFmWj818I.rQxDMS()
	bUs30J :=  /*line :1*/ uint32( /*line :1*/ len(fFmWj818I.aeSIvF7Fa))

	doauF8Y, h_ljk48Bm :=  /*line :1*/ dnsmessage.HfBZcpE4(tahgc783Bc)
	if h_ljk48Bm != nil {
		return dnsmessage.GShepB{}, "", p6JX85H_v
	}
	fsa6C0Qdm := dnsmessage.LNQ5aNrCORB{
		FLDYrkuX:	doauF8Y,
		XsqCP7Wa65g:	vBrGtQW,
		GBiiAZ64kqf6:	dnsmessage.ClassINET,
	}

	for eaAUaB7Z := 0; eaAUaB7Z < fFmWj818I.yb6dJ6U; eaAUaB7Z++ {
		for kVgtfLFZ :=  /*line :1*/ uint32(0); kVgtfLFZ < bUs30J; kVgtfLFZ++ {
			c_eHMUdT := fFmWj818I.aeSIvF7Fa[(rQxDMS+kVgtfLFZ)%bUs30J]

			fMPVz2iGiyq, cBprzKT8, h_ljk48Bm :=  /*line :1*/ yxhs4Z.rMsDNw(yESLyde9LHhT, c_eHMUdT, fsa6C0Qdm, fFmWj818I.c1XqMJ, fFmWj818I.u6Uf2W_GEl, fFmWj818I.gsW5rVBCAq3)
			if h_ljk48Bm != nil {
				qlLRYYI22 := &SoatTK0{
					PY3bIIR:	 /*line :1*/ h_ljk48Bm.Error(),
					FIvHQdTCAg:	tahgc783Bc,
					IoaaWnhxbb:	c_eHMUdT,
				}
				if j_J3nr, d30HIwxht := h_ljk48Bm.(XL_ACKzCw); d30HIwxht &&  /*line :1*/ j_J3nr.Timeout() {
					qlLRYYI22.JYeYrD = true
				}

				if _, d30HIwxht := h_ljk48Bm.(*ZOYGdck); d30HIwxht {
					qlLRYYI22.HzqkaqiQE1P = true
				}
				foP_Mw6F0ef = qlLRYYI22
				continue
			}

			if h_ljk48Bm :=  /*line :1*/ siVD3Ya(&fMPVz2iGiyq, cBprzKT8); h_ljk48Bm != nil {
				qlLRYYI22 := &SoatTK0{
					PY3bIIR:	 /*line :1*/ h_ljk48Bm.Error(),
					FIvHQdTCAg:	tahgc783Bc,
					IoaaWnhxbb:	c_eHMUdT,
				}
				if h_ljk48Bm == k3srm4 {
					qlLRYYI22.HzqkaqiQE1P = true
				}
				if h_ljk48Bm == aamCgdkZikV {

					qlLRYYI22.GM17caVp1uW = true
					return fMPVz2iGiyq, c_eHMUdT, qlLRYYI22
				}
				foP_Mw6F0ef = qlLRYYI22
				continue
			}

			h_ljk48Bm =  /*line :1*/ ouo_Rd(&fMPVz2iGiyq, vBrGtQW)
			if h_ljk48Bm == nil {
				return fMPVz2iGiyq, c_eHMUdT, nil
			}
			foP_Mw6F0ef = &SoatTK0{
				PY3bIIR:	 /*line :1*/ h_ljk48Bm.Error(),
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
			if h_ljk48Bm == aamCgdkZikV {

				foP_Mw6F0ef.(*SoatTK0).GM17caVp1uW = true
				return fMPVz2iGiyq, c_eHMUdT, foP_Mw6F0ef
			}
		}
	}
	return dnsmessage.GShepB{}, "", foP_Mw6F0ef
}

type grtITnqUiIDq struct {
	us66PE9d	sync.LhBfwn6wa1x

	fjteYw	chan struct{}
	eYAmefjE	time.G4KDkI

	hTyOoJSqgD	atomic.ToSaNw[ivgt6m77Buqr]
}

var eBVPJvmpX grtITnqUiIDq

func mT_rTj1r8a() *ivgt6m77Buqr {
	 /*line :1*/ eBVPJvmpX.zRkY41bMBy("/etc/resolv.conf")
	return  /*line :1*/ eBVPJvmpX.hTyOoJSqgD.Load()
}

func (q6Glr4X1u *grtITnqUiIDq) init() {

	 /*line :1*/ q6Glr4X1u.hTyOoJSqgD.Store( /*line :1*/ nP4L7Rk4("/etc/resolv.conf"))
	q6Glr4X1u.eYAmefjE =  /*line :1*/ time.Pc_35oTY()

	q6Glr4X1u.fjteYw =  /*line :1*/ make(chan struct{}, 1)
}

func (q6Glr4X1u *grtITnqUiIDq) zRkY41bMBy(tahgc783Bc string) {
	 /*line :1*/ q6Glr4X1u.us66PE9d.Do(q6Glr4X1u.init)

	if  /*line :1*/ q6Glr4X1u.hTyOoJSqgD.Load().aJd1TsOhC3ZG {
		return
	}

	if ! /*line :1*/ q6Glr4X1u.xfyLaBaWu9e4() {
		return
	}
	defer  /*line :1*/ q6Glr4X1u.wDHVRjOehJ()

	t_5qtVaQY :=  /*line :1*/ time.Pc_35oTY()
	if  /*line :1*/ q6Glr4X1u.eYAmefjE.After( /*line :1*/ t_5qtVaQY.Add(-5 * time.Second)) {
		return
	}
	q6Glr4X1u.eYAmefjE = t_5qtVaQY

	switch runtime.GOOS {
	case "windows":

	default:
		var i81SGZwwa time.G4KDkI
		if a6qU1P1, h_ljk48Bm :=  /*line :1*/ os.ZYa3wuv(tahgc783Bc); h_ljk48Bm == nil {
			i81SGZwwa =  /*line :1*/ a6qU1P1.ModTime()
		}
		if  /*line :1*/ i81SGZwwa.Equal( /*line :1*/ q6Glr4X1u.hTyOoJSqgD.Load().zOI26Q) {
			return
		}
	}

	vTBPTe :=  /*line :1*/ nP4L7Rk4(tahgc783Bc)
	 /*line :1*/ q6Glr4X1u.hTyOoJSqgD.Store(vTBPTe)
}

func (q6Glr4X1u *grtITnqUiIDq) xfyLaBaWu9e4() bool {
	select {
	case q6Glr4X1u.fjteYw <- struct{}{}:
		return true
	default:
		return false
	}
}

func (q6Glr4X1u *grtITnqUiIDq) wDHVRjOehJ() {
	<-q6Glr4X1u.fjteYw
}

func (yxhs4Z *W8Q2tfHAD) bp6HtmNGS7(yESLyde9LHhT context.MBCyA6, tahgc783Bc string, vBrGtQW dnsmessage.Ze6KllqJw, q6Glr4X1u *ivgt6m77Buqr) (dnsmessage.GShepB, string, error) {
	if ! /*line :1*/ tgYtaZX8VvFU(tahgc783Bc) {

		return dnsmessage.GShepB{}, "", &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: tahgc783Bc, GM17caVp1uW: true}
	}

	if q6Glr4X1u == nil {
		q6Glr4X1u =  /*line :1*/ mT_rTj1r8a()
	}

	var (
		fMPVz2iGiyq	dnsmessage.GShepB
		c_eHMUdT	string
		h_ljk48Bm	error
	)
	for _, cJoem9 := range  /*line :1*/ q6Glr4X1u.i9l2n2(tahgc783Bc) {
		fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm =  /*line :1*/ yxhs4Z.cFpiL14zqMrj(yESLyde9LHhT, q6Glr4X1u, cJoem9, vBrGtQW)
		if h_ljk48Bm == nil {
			break
		}
		if j_J3nr, d30HIwxht := h_ljk48Bm.(XL_ACKzCw); d30HIwxht &&  /*line :1*/ j_J3nr.Temporary() &&  /*line :1*/ yxhs4Z.jRUqBvW7RAfL() {

			break
		}
	}
	if h_ljk48Bm == nil {
		return fMPVz2iGiyq, c_eHMUdT, nil
	}
	if h_ljk48Bm, d30HIwxht := h_ljk48Bm.(*SoatTK0); d30HIwxht {

		h_ljk48Bm.FIvHQdTCAg = tahgc783Bc
	}
	return dnsmessage.GShepB{}, "", h_ljk48Bm
}

func jAvL_3izoS_(tahgc783Bc string) bool {
	if tahgc783Bc == "" {
		return true
	}
	if tahgc783Bc[ /*line :1*/ len(tahgc783Bc)-1] == '.' {
		tahgc783Bc = tahgc783Bc[: /*line :1*/ len(tahgc783Bc)-1]
	}
	return  /*line :1*/ jRYeKNE(tahgc783Bc, ".onion")
}

func (q6Glr4X1u *ivgt6m77Buqr) i9l2n2(tahgc783Bc string) []string {
	if  /*line :1*/ jAvL_3izoS_(tahgc783Bc) {
		return nil
	}

	v3upVm :=  /*line :1*/ len(tahgc783Bc)
	_j1OBpWGKUdh := v3upVm > 0 && tahgc783Bc[v3upVm-1] == '.'
	if v3upVm > 254 || v3upVm == 254 && !_j1OBpWGKUdh {
		return nil
	}

	if _j1OBpWGKUdh {
		return []string{tahgc783Bc}
	}

	clR6nVya5int :=  /*line :1*/ mB0PZSLg9(tahgc783Bc, '.') >= q6Glr4X1u.kzjIfgwIhJ_p
	tahgc783Bc += "."
	v3upVm++

	z_5pYyI7e :=  /*line :1*/ make([]string, 0, 1+ /*line :1*/ len(q6Glr4X1u.fjA5TZQl7uU))

	if clR6nVya5int {
		z_5pYyI7e =  /*line :1*/ append(z_5pYyI7e, tahgc783Bc)
	}

	for _, oWGiddK := range q6Glr4X1u.fjA5TZQl7uU {
		if v3upVm+ /*line :1*/ len(oWGiddK) <= 254 {
			z_5pYyI7e =  /*line :1*/ append(z_5pYyI7e, tahgc783Bc+oWGiddK)
		}
	}

	if !clR6nVya5int {
		z_5pYyI7e =  /*line :1*/ append(z_5pYyI7e, tahgc783Bc)
	}
	return z_5pYyI7e
}

type mNRpaDPvaD int

const (
	hostLookupCgo	mNRpaDPvaD	= iota
	hostLookupFilesDNS
	hostLookupDNSFiles
	hostLookupFiles
	hostLookupDNS
)

var cb3qsb = map[mNRpaDPvaD]string{
	hostLookupCgo:	"cgo",
	hostLookupFilesDNS:	"files,dns",
	hostLookupDNSFiles:	"dns,files",
	hostLookupFiles:	"files",
	hostLookupDNS:	"dns",
}

func (be01m8m6Z mNRpaDPvaD) String() string {
	if dRoFccaZ, d30HIwxht := cb3qsb[be01m8m6Z]; d30HIwxht {
		return dRoFccaZ
	}
	return "hostLookupOrder=" +  /*line :1*/ itoa.V25ba9G5( /*line :1*/ int(be01m8m6Z)) + "??"
}

func (yxhs4Z *W8Q2tfHAD) fnlAS_QK_(yESLyde9LHhT context.MBCyA6, tahgc783Bc string, fM1jn7loPH mNRpaDPvaD, q6Glr4X1u *ivgt6m77Buqr) (md4QSRkO []string, h_ljk48Bm error) {
	if fM1jn7loPH == hostLookupFilesDNS || fM1jn7loPH == hostLookupFiles {

		md4QSRkO, _ =  /*line :1*/ y_FPhQLXqT(tahgc783Bc)
		if  /*line :1*/ len(md4QSRkO) > 0 {
			return
		}

		if fM1jn7loPH == hostLookupFiles {
			return nil, &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: tahgc783Bc, GM17caVp1uW: true}
		}
	}
	ecxaiU5w1ARk, _, h_ljk48Bm :=  /*line :1*/ yxhs4Z.qq3X26Z6gS(yESLyde9LHhT, "ip", tahgc783Bc, fM1jn7loPH, q6Glr4X1u)
	if h_ljk48Bm != nil {
		return
	}
	md4QSRkO =  /*line :1*/ make([]string, 0,  /*line :1*/ len(ecxaiU5w1ARk))
	for _, kQ8_UEhxU := range ecxaiU5w1ARk {
		md4QSRkO =  /*line :1*/ append(md4QSRkO,  /*line :1*/ kQ8_UEhxU.String())
	}
	return
}

func qLmu3LrEJFy(tahgc783Bc string) (md4QSRkO []FZJphYv9hV, eLUDzideg string) {
	qxwkC3VxG0, eLUDzideg :=  /*line :1*/ y_FPhQLXqT(tahgc783Bc)
	for _, _WIn4q := range qxwkC3VxG0 {
		_WIn4q, dJhOHOEZlQAA :=  /*line :1*/ jewoNRQXK6(_WIn4q)
		if kQ8_UEhxU :=  /*line :1*/ XtVFDxFx(_WIn4q); kQ8_UEhxU != nil {
			qxwkC3VxG0 := FZJphYv9hV{GdouSNq80bRw: kQ8_UEhxU, Cyg5M2vYIby: dJhOHOEZlQAA}
			md4QSRkO =  /*line :1*/ append(md4QSRkO, qxwkC3VxG0)
		}
	}
	 /*line :1*/ e6vy2evrzCS2(md4QSRkO)
	return md4QSRkO, eLUDzideg
}

func (yxhs4Z *W8Q2tfHAD) yIyaYwBXWU(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, jKxI_T1CK_p string) (md4QSRkO []FZJphYv9hV, h_ljk48Bm error) {
	fM1jn7loPH, q6Glr4X1u :=  /*line :1*/ yAKpTR().mNRpaDPvaD(yxhs4Z, jKxI_T1CK_p)
	md4QSRkO, _, h_ljk48Bm =  /*line :1*/ yxhs4Z.qq3X26Z6gS(yESLyde9LHhT, vsbiWLn7reX, jKxI_T1CK_p, fM1jn7loPH, q6Glr4X1u)
	return
}

func (yxhs4Z *W8Q2tfHAD) qq3X26Z6gS(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, tahgc783Bc string, fM1jn7loPH mNRpaDPvaD, q6Glr4X1u *ivgt6m77Buqr) (md4QSRkO []FZJphYv9hV, jcluQhsSxoD dnsmessage.Toccq2dE3, h_ljk48Bm error) {
	if fM1jn7loPH == hostLookupFilesDNS || fM1jn7loPH == hostLookupFiles {
		var eLUDzideg string
		md4QSRkO, eLUDzideg =  /*line :1*/ qLmu3LrEJFy(tahgc783Bc)

		if  /*line :1*/ len(md4QSRkO) > 0 {
			var h_ljk48Bm error
			jcluQhsSxoD, h_ljk48Bm =  /*line :1*/ dnsmessage.HfBZcpE4(eLUDzideg)
			if h_ljk48Bm != nil {
				return nil, dnsmessage.Toccq2dE3{}, h_ljk48Bm
			}
			return md4QSRkO, jcluQhsSxoD, nil
		}

		if fM1jn7loPH == hostLookupFiles {
			return nil, dnsmessage.Toccq2dE3{}, &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: tahgc783Bc, GM17caVp1uW: true}
		}
	}

	if ! /*line :1*/ tgYtaZX8VvFU(tahgc783Bc) {

		return nil, dnsmessage.Toccq2dE3{}, &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: tahgc783Bc, GM17caVp1uW: true}
	}
	type w1SqzJRu1Q struct {
		sVtIqUn	dnsmessage.GShepB
		jRMUAoa_	string
		error
	}

	if q6Glr4X1u == nil {
		q6Glr4X1u =  /*line :1*/ mT_rTj1r8a()
	}

	gQ27FU0_a7gW :=  /*line :1*/ make(chan w1SqzJRu1Q, 1)
	jmXM_eUMkAj := []dnsmessage.Ze6KllqJw{dnsmessage.TypeA, dnsmessage.TypeAAAA}
	if vsbiWLn7reX == "CNAME" {
		jmXM_eUMkAj =  /*line :1*/ append(jmXM_eUMkAj, dnsmessage.TypeCNAME)
	}
	switch  /*line :1*/ qqb9mdSqc(vsbiWLn7reX) {
	case '4':
		jmXM_eUMkAj = []dnsmessage.Ze6KllqJw{dnsmessage.TypeA}
	case '6':
		jmXM_eUMkAj = []dnsmessage.Ze6KllqJw{dnsmessage.TypeAAAA}
	}
	var bjgISM1 func(cJoem9 string, vBrGtQW dnsmessage.Ze6KllqJw)
	var mFUAm1 func(cJoem9 string, vBrGtQW dnsmessage.Ze6KllqJw) w1SqzJRu1Q
	if q6Glr4X1u.eMyOLk {
		bjgISM1 = func(cJoem9 string, vBrGtQW dnsmessage.Ze6KllqJw) {}
		mFUAm1 = func(cJoem9 string, vBrGtQW dnsmessage.Ze6KllqJw) w1SqzJRu1Q {
			 /*line :1*/ k76lUDr6.Add(1)
			defer  /*line :1*/ k76lUDr6.Done()
			fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm :=  /*line :1*/ yxhs4Z.cFpiL14zqMrj(yESLyde9LHhT, q6Glr4X1u, cJoem9, vBrGtQW)
			return w1SqzJRu1Q{fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm}
		}
	} else {
		bjgISM1 = func(cJoem9 string, vBrGtQW dnsmessage.Ze6KllqJw) {
			 /*line :1*/ k76lUDr6.Add(1)
			go func( /*line :1*/ vBrGtQW dnsmessage.Ze6KllqJw) {
				fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm :=  /*line :1*/ yxhs4Z.cFpiL14zqMrj(yESLyde9LHhT, q6Glr4X1u, cJoem9, vBrGtQW)
				gQ27FU0_a7gW <- w1SqzJRu1Q{fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm}
				 /*line :1*/ k76lUDr6.Done()
			}(vBrGtQW)
		}
		mFUAm1 = func(cJoem9 string, vBrGtQW dnsmessage.Ze6KllqJw) w1SqzJRu1Q {
			return <-gQ27FU0_a7gW
		}
	}
	var foP_Mw6F0ef error
	for _, cJoem9 := range  /*line :1*/ q6Glr4X1u.i9l2n2(tahgc783Bc) {
		for _, vBrGtQW := range jmXM_eUMkAj {
			 /*line :1*/ bjgISM1(cJoem9, vBrGtQW)
		}
		kPOlAHg := false
		for _, vBrGtQW := range jmXM_eUMkAj {
			w1SqzJRu1Q :=  /*line :1*/ mFUAm1(cJoem9, vBrGtQW)
			if w1SqzJRu1Q.error != nil {
				if j_J3nr, d30HIwxht := w1SqzJRu1Q.error.(XL_ACKzCw); d30HIwxht &&  /*line :1*/ j_J3nr.Temporary() &&  /*line :1*/ yxhs4Z.jRUqBvW7RAfL() {

					kPOlAHg = true
					foP_Mw6F0ef = w1SqzJRu1Q.error
				} else if foP_Mw6F0ef == nil || cJoem9 == tahgc783Bc+"." {

					foP_Mw6F0ef = w1SqzJRu1Q.error
				}
				continue
			}

		loop:
			for {
				cBprzKT8, h_ljk48Bm :=  /*line :1*/ w1SqzJRu1Q.sVtIqUn.AnswerHeader()
				if h_ljk48Bm != nil && h_ljk48Bm != dnsmessage.BbQhjXWx {
					foP_Mw6F0ef = &SoatTK0{
						PY3bIIR:	"cannot marshal DNS message",
						FIvHQdTCAg:	tahgc783Bc,
						IoaaWnhxbb:	w1SqzJRu1Q.jRMUAoa_,
					}
				}
				if h_ljk48Bm != nil {
					break
				}
				switch cBprzKT8.BxqvzBsK {
				case dnsmessage.TypeA:
					uI7LZDHm6, h_ljk48Bm :=  /*line :1*/ w1SqzJRu1Q.sVtIqUn.AResource()
					if h_ljk48Bm != nil {
						foP_Mw6F0ef = &SoatTK0{
							PY3bIIR:	"cannot marshal DNS message",
							FIvHQdTCAg:	tahgc783Bc,
							IoaaWnhxbb:	w1SqzJRu1Q.jRMUAoa_,
						}
						break loop
					}
					md4QSRkO =  /*line :1*/ append(md4QSRkO, FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ GNraIvYhB(uI7LZDHm6.M5FpSRlH9[:])})
					if jcluQhsSxoD.W7uI8JuUD == 0 && cBprzKT8.Pz3XQBChF6.W7uI8JuUD != 0 {
						jcluQhsSxoD = cBprzKT8.Pz3XQBChF6
					}

				case dnsmessage.TypeAAAA:
					gI1LqI7Zaw, h_ljk48Bm :=  /*line :1*/ w1SqzJRu1Q.sVtIqUn.AAAAResource()
					if h_ljk48Bm != nil {
						foP_Mw6F0ef = &SoatTK0{
							PY3bIIR:	"cannot marshal DNS message",
							FIvHQdTCAg:	tahgc783Bc,
							IoaaWnhxbb:	w1SqzJRu1Q.jRMUAoa_,
						}
						break loop
					}
					md4QSRkO =  /*line :1*/ append(md4QSRkO, FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ GNraIvYhB(gI1LqI7Zaw.DaUu4HQOB9[:])})
					if jcluQhsSxoD.W7uI8JuUD == 0 && cBprzKT8.Pz3XQBChF6.W7uI8JuUD != 0 {
						jcluQhsSxoD = cBprzKT8.Pz3XQBChF6
					}

				case dnsmessage.TypeCNAME:
					hl8w2lFX, h_ljk48Bm :=  /*line :1*/ w1SqzJRu1Q.sVtIqUn.CNAMEResource()
					if h_ljk48Bm != nil {
						foP_Mw6F0ef = &SoatTK0{
							PY3bIIR:	"cannot marshal DNS message",
							FIvHQdTCAg:	tahgc783Bc,
							IoaaWnhxbb:	w1SqzJRu1Q.jRMUAoa_,
						}
						break loop
					}
					if jcluQhsSxoD.W7uI8JuUD == 0 && hl8w2lFX.SiT_7E895B87.W7uI8JuUD > 0 {
						jcluQhsSxoD = hl8w2lFX.SiT_7E895B87
					}

				default:
					if h_ljk48Bm :=  /*line :1*/ w1SqzJRu1Q.sVtIqUn.SkipAnswer(); h_ljk48Bm != nil {
						foP_Mw6F0ef = &SoatTK0{
							PY3bIIR:	"cannot marshal DNS message",
							FIvHQdTCAg:	tahgc783Bc,
							IoaaWnhxbb:	w1SqzJRu1Q.jRMUAoa_,
						}
						break loop
					}
					continue
				}
			}
		}
		if kPOlAHg {

			md4QSRkO = nil
			break
		}
		if  /*line :1*/ len(md4QSRkO) > 0 || vsbiWLn7reX == "CNAME" && jcluQhsSxoD.W7uI8JuUD > 0 {
			break
		}
	}
	if foP_Mw6F0ef, d30HIwxht := foP_Mw6F0ef.(*SoatTK0); d30HIwxht {

		foP_Mw6F0ef.FIvHQdTCAg = tahgc783Bc
	}
	 /*line :1*/ e6vy2evrzCS2(md4QSRkO)
	if  /*line :1*/ len(md4QSRkO) == 0 && !(vsbiWLn7reX == "CNAME" && jcluQhsSxoD.W7uI8JuUD > 0) {
		if fM1jn7loPH == hostLookupDNSFiles {
			var eLUDzideg string
			md4QSRkO, eLUDzideg =  /*line :1*/ qLmu3LrEJFy(tahgc783Bc)
			if  /*line :1*/ len(md4QSRkO) > 0 {
				var h_ljk48Bm error
				jcluQhsSxoD, h_ljk48Bm =  /*line :1*/ dnsmessage.HfBZcpE4(eLUDzideg)
				if h_ljk48Bm != nil {
					return nil, dnsmessage.Toccq2dE3{}, h_ljk48Bm
				}
				return md4QSRkO, jcluQhsSxoD, nil
			}
		}
		if foP_Mw6F0ef != nil {
			return nil, dnsmessage.Toccq2dE3{}, foP_Mw6F0ef
		}
	}
	return md4QSRkO, jcluQhsSxoD, nil
}

func (yxhs4Z *W8Q2tfHAD) kK7DIGIx(yESLyde9LHhT context.MBCyA6, jKxI_T1CK_p string, fM1jn7loPH mNRpaDPvaD, q6Glr4X1u *ivgt6m77Buqr) (string, error) {
	_, jcluQhsSxoD, h_ljk48Bm :=  /*line :1*/ yxhs4Z.qq3X26Z6gS(yESLyde9LHhT, "CNAME", jKxI_T1CK_p, fM1jn7loPH, q6Glr4X1u)
	return  /*line :1*/ jcluQhsSxoD.String(), h_ljk48Bm
}

func (yxhs4Z *W8Q2tfHAD) nHwfF3h(yESLyde9LHhT context.MBCyA6, qxwkC3VxG0 string, fM1jn7loPH mNRpaDPvaD, q6Glr4X1u *ivgt6m77Buqr) ([]string, error) {
	if fM1jn7loPH == hostLookupFiles || fM1jn7loPH == hostLookupFilesDNS {
		z_5pYyI7e :=  /*line :1*/ cb0AaP(qxwkC3VxG0)
		if  /*line :1*/ len(z_5pYyI7e) > 0 {
			return z_5pYyI7e, nil
		}

		if fM1jn7loPH == hostLookupFiles {
			return nil, &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: qxwkC3VxG0, GM17caVp1uW: true}
		}
	}

	akvMHAa84xdt, h_ljk48Bm :=  /*line :1*/ sPSXZcPfQnVj(qxwkC3VxG0)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm :=  /*line :1*/ yxhs4Z.bp6HtmNGS7(yESLyde9LHhT, akvMHAa84xdt, dnsmessage.TypePTR, q6Glr4X1u)
	if h_ljk48Bm != nil {
		var qlLRYYI22 *SoatTK0
		if  /*line :1*/ errors.CyKhHznARJK(h_ljk48Bm, &qlLRYYI22) && qlLRYYI22.GM17caVp1uW {
			if fM1jn7loPH == hostLookupDNSFiles {
				z_5pYyI7e :=  /*line :1*/ cb0AaP(qxwkC3VxG0)
				if  /*line :1*/ len(z_5pYyI7e) > 0 {
					return z_5pYyI7e, nil
				}
			}
		}
		return nil, h_ljk48Bm
	}
	var uvaup8zGu []string
	for {
		cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.AnswerHeader()
		if h_ljk48Bm == dnsmessage.BbQhjXWx {
			break
		}
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot marshal DNS message",
				FIvHQdTCAg:	qxwkC3VxG0,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		if cBprzKT8.BxqvzBsK != dnsmessage.TypePTR {
			h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SkipAnswer()
			if h_ljk48Bm != nil {
				return nil, &SoatTK0{
					PY3bIIR:	"cannot marshal DNS message",
					FIvHQdTCAg:	qxwkC3VxG0,
					IoaaWnhxbb:	c_eHMUdT,
				}
			}
			continue
		}
		na3CNRWz, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.PTRResource()
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot marshal DNS message",
				FIvHQdTCAg:	qxwkC3VxG0,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		uvaup8zGu =  /*line :1*/ append(uvaup8zGu,  /*line :1*/ na3CNRWz.ExEq5bU3Lem.String())

	}

	return uvaup8zGu, nil
}
