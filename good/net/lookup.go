//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	errors "iAHoxjmM"
	nettrace "KwiKvcnR5Bq"
	singleflight "EkvK9so"
	netip "iPdCam_KQOXj"
	sync "sync"

	dnsmessage "FeKim_cPgM7z"
)








var vxMdwXut7bf = map[string]int{
	"icmp":	1,
	"igmp":	2,
	"tcp":	6,
	"udp":	17,
	"ipv6-icmp":	58,
}







var mJkMwi1WUav5 = map[string]map[string]int{
	"udp": {
		"domain": 53,
	},
	"tcp": {
		"ftp":	21,
		"ftps":	990,
		"gopher":	70,
		"http":	80,
		"https":	443,
		"imap2":	143,
		"imap3":	220,
		"imaps":	993,
		"pop3":	110,
		"pop3s":	995,
		"smtp":	25,
		"ssh":	22,
		"telnet":	23,
	},
}



var k76lUDr6 sync.NMlO0OA

const maxProtoLength =  /*line :1*/ len("RSVP-E2E-IGNORE") + 10	

func lOVQ14ZJh(tahgc783Bc string) (int, error) {
	var a4XdPaSmQ [maxProtoLength]byte
	doauF8Y :=  /*line :1*/ copy(a4XdPaSmQ[:], tahgc783Bc)
	 /*line :1*/ kuRb7e(a4XdPaSmQ[:doauF8Y])
	iaebRs5X3, wwbLpRWjJdgJ := vxMdwXut7bf[ /*line :1*/ string(a4XdPaSmQ[:doauF8Y])]
	if !wwbLpRWjJdgJ || doauF8Y !=  /*line :1*/ len(tahgc783Bc) {
		return 0, &DAWLIQHc{Z68v0y: "unknown IP protocol specified", BCCnAgFxG: tahgc783Bc}
	}
	return iaebRs5X3, nil
}






const maxPortBufSize =  /*line :1*/ len("mobility-header") + 10

func dTwMnMHVj(vsbiWLn7reX, zvX72x8Ih7C string) (mzUgFPgljgC int, lAieqzJA error) {
	switch vsbiWLn7reX {
	case "tcp4", "tcp6":
		vsbiWLn7reX = "tcp"
	case "udp4", "udp6":
		vsbiWLn7reX = "udp"
	}

	if z_nqmUMeJFH, d30HIwxht := mJkMwi1WUav5[vsbiWLn7reX]; d30HIwxht {
		var cleZab [maxPortBufSize]byte
		doauF8Y :=  /*line :1*/ copy(cleZab[:], zvX72x8Ih7C)
		 /*line :1*/ kuRb7e(cleZab[:doauF8Y])
		if mzUgFPgljgC, d30HIwxht := z_nqmUMeJFH[ /*line :1*/ string(cleZab[:doauF8Y])]; d30HIwxht && doauF8Y ==  /*line :1*/ len(zvX72x8Ih7C) {
			return mzUgFPgljgC, nil
		}
	}
	return 0, &DAWLIQHc{Z68v0y: "unknown port", BCCnAgFxG: vsbiWLn7reX + "/" + zvX72x8Ih7C}
}



func qqb9mdSqc(vsbiWLn7reX string) byte {
	if vsbiWLn7reX == "" {
		return 0
	}
	doauF8Y := vsbiWLn7reX[ /*line :1*/ len(vsbiWLn7reX)-1]
	if doauF8Y != '4' && doauF8Y != '6' {
		doauF8Y = 0
	}
	return doauF8Y
}



var Ic4wtIY = &W8Q2tfHAD{}




type W8Q2tfHAD struct {
	
	
	
	JAh8vU3l	bool

	
	
	
	
	
	
	
	
	PyZBhYz9UpKy	bool

	
	
	
	
	
	
	
	
	
	
	
	B5GyJD	func(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, fwV_ln2dl string) (UJYB3aCQg, error)

	
	
	
	v5sSVkXvOZ	singleflight.JRXHs_
}

func (yxhs4Z *W8Q2tfHAD) uza_Hv() bool	{ return yxhs4Z != nil && yxhs4Z.JAh8vU3l }
func (yxhs4Z *W8Q2tfHAD) jRUqBvW7RAfL() bool	{ return yxhs4Z != nil && yxhs4Z.PyZBhYz9UpKy }

func (yxhs4Z *W8Q2tfHAD) cbjvbAUX() *singleflight.JRXHs_ {
	if yxhs4Z == nil {
		return &Ic4wtIY.v5sSVkXvOZ
	}
	return &yxhs4Z.v5sSVkXvOZ
}






func QlrHmD8halB(jKxI_T1CK_p string) (md4QSRkO []string, h_ljk48Bm error) {
	return  /*line :1*/ Ic4wtIY.LookupHost( /*line :1*/ context.GEcgQP5fzA(), jKxI_T1CK_p)
}



func (yxhs4Z *W8Q2tfHAD) LookupHost(yESLyde9LHhT context.MBCyA6, jKxI_T1CK_p string) (md4QSRkO []string, h_ljk48Bm error) {

	if jKxI_T1CK_p == "" {
		return nil, &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: jKxI_T1CK_p, GM17caVp1uW: true}
	}
	if _, h_ljk48Bm :=  /*line :1*/ netip.Y1uXU4Sx(jKxI_T1CK_p); h_ljk48Bm == nil {
		return []string{jKxI_T1CK_p}, nil
	}
	return  /*line :1*/ yxhs4Z.c4c5ngLuoZKl(yESLyde9LHhT, jKxI_T1CK_p)
}



func Uw4LxUQaf2c(jKxI_T1CK_p string) ([]GNraIvYhB, error) {
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ Ic4wtIY.LookupIPAddr( /*line :1*/ context.GEcgQP5fzA(), jKxI_T1CK_p)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	ecxaiU5w1ARk :=  /*line :1*/ make([]GNraIvYhB,  /*line :1*/ len(md4QSRkO))
	for eaAUaB7Z, caaBOzZX := range md4QSRkO {
		ecxaiU5w1ARk[eaAUaB7Z] = caaBOzZX.GdouSNq80bRw
	}
	return ecxaiU5w1ARk, nil
}



func (yxhs4Z *W8Q2tfHAD) LookupIPAddr(yESLyde9LHhT context.MBCyA6, jKxI_T1CK_p string) ([]FZJphYv9hV, error) {
	return  /*line :1*/ yxhs4Z.ffGrcqOlo(yESLyde9LHhT, "ip", jKxI_T1CK_p)
}





func (yxhs4Z *W8Q2tfHAD) LookupIP(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, jKxI_T1CK_p string) ([]GNraIvYhB, error) {
	mQBJul9fW, _, h_ljk48Bm :=  /*line :1*/ nihn79f8iY(yESLyde9LHhT, vsbiWLn7reX, false)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	switch mQBJul9fW {
	case "ip", "ip4", "ip6":
	default:
		return nil,  /*line :1*/ JFaUU0ZaU(vsbiWLn7reX)
	}

	if jKxI_T1CK_p == "" {
		return nil, &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: jKxI_T1CK_p, GM17caVp1uW: true}
	}
	md4QSRkO, h_ljk48Bm :=  /*line :1*/ yxhs4Z.hCHWGoQGe2q(yESLyde9LHhT, mQBJul9fW, jKxI_T1CK_p)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}

	ecxaiU5w1ARk :=  /*line :1*/ make([]GNraIvYhB, 0,  /*line :1*/ len(md4QSRkO))
	for _, qxwkC3VxG0 := range md4QSRkO {
		ecxaiU5w1ARk =  /*line :1*/ append(ecxaiU5w1ARk, qxwkC3VxG0.(*FZJphYv9hV).GdouSNq80bRw)
	}
	return ecxaiU5w1ARk, nil
}





func (yxhs4Z *W8Q2tfHAD) LookupNetIP(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, jKxI_T1CK_p string) ([]netip.KFLQ1_1E, error) {

	ecxaiU5w1ARk, h_ljk48Bm :=  /*line :1*/ yxhs4Z.LookupIP(yESLyde9LHhT, vsbiWLn7reX, jKxI_T1CK_p)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	prHUfwjRzJ :=  /*line :1*/ make([]netip.KFLQ1_1E, 0,  /*line :1*/ len(ecxaiU5w1ARk))
	for _, kQ8_UEhxU := range ecxaiU5w1ARk {
		if uI7LZDHm6, d30HIwxht :=  /*line :1*/ netip.OWcivL5ykj9(kQ8_UEhxU); d30HIwxht {
			prHUfwjRzJ =  /*line :1*/ append(prHUfwjRzJ, uI7LZDHm6)
		}
	}
	return prHUfwjRzJ, nil
}



type lvV9Cw struct {
	context.MBCyA6
	gkENm6_84CHN	context.MBCyA6
}

var _ context.MBCyA6 = (* /*line :1*/ lvV9Cw)(nil)


func (aAvHgot_ *lvV9Cw) Value(qFUmBERYgmKz any) any {
	select {
	case <- /*line :1*/ aAvHgot_.gkENm6_84CHN.Done():
		return nil
	default:
		return  /*line :1*/ aAvHgot_.gkENm6_84CHN.Value(qFUmBERYgmKz)
	}
}





func pzMSQwF5(n8jaGIF1 context.MBCyA6) context.MBCyA6 {
	return &lvV9Cw{MBCyA6:  /*line :1*/ context.GEcgQP5fzA(), gkENm6_84CHN: n8jaGIF1}
}



func (yxhs4Z *W8Q2tfHAD) ffGrcqOlo(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, jKxI_T1CK_p string) ([]FZJphYv9hV, error) {

	if jKxI_T1CK_p == "" {
		return nil, &SoatTK0{PY3bIIR:  /*line :1*/ aamCgdkZikV.Error(), FIvHQdTCAg: jKxI_T1CK_p, GM17caVp1uW: true}
	}
	if kQ8_UEhxU, h_ljk48Bm :=  /*line :1*/ netip.Y1uXU4Sx(jKxI_T1CK_p); h_ljk48Bm == nil {
		return []FZJphYv9hV{{GdouSNq80bRw:  /*line :1*/ GNraIvYhB( /*line :1*/ kQ8_UEhxU.AsSlice()).To16(), Cyg5M2vYIby:  /*line :1*/ kQ8_UEhxU.Zone()}}, nil
	}
	jHCsMONOrYFN, _ :=  /*line :1*/ yESLyde9LHhT.Value(nettrace.YRC9V_{}).(*nettrace.UTp90oOoUInw)
	if jHCsMONOrYFN != nil && jHCsMONOrYFN.LiO9PhybsIj != nil {
		 /*line :1*/ jHCsMONOrYFN.LiO9PhybsIj(jKxI_T1CK_p)
	}

	j9BKxR := yxhs4Z.uPMUqJmf
	if sskTxsj, _ :=  /*line :1*/ yESLyde9LHhT.Value(nettrace.FfQPhE{}).(func(context.MBCyA6, string, string) ([]FZJphYv9hV, error)); sskTxsj != nil {
		j9BKxR = sskTxsj
	}

	vbSlkFxVlc, jSYGwe :=  /*line :1*/ context.Cu9yRv( /*line :1*/ pzMSQwF5(yESLyde9LHhT))

	fl5wWf := vsbiWLn7reX + "\000" + jKxI_T1CK_p
	 /*line :1*/ k76lUDr6.Add(1)
	rvEj5N7 :=  /*line :1*/ yxhs4Z.cbjvbAUX().DoChan(fl5wWf, func() (any, error) {
		return  /*line :1*/ eV6agl(vbSlkFxVlc, j9BKxR, vsbiWLn7reX, jKxI_T1CK_p)
	})

	w4VFts37B3r := func(rvEj5N7 <-chan singleflight.RfcQmoNLzjV, xrUd2L context.Gd7QRmF) {
		<-rvEj5N7
		 /*line :1*/ k76lUDr6.Done()
		 /*line :1*/ xrUd2L()
	}
	select {
	case <- /*line :1*/ yESLyde9LHhT.Done():

		if  /*line :1*/ yxhs4Z.cbjvbAUX().ForgetUnshared(fl5wWf) {
			 /*line :1*/ jSYGwe()
			go  /*line :1*/ w4VFts37B3r(rvEj5N7, func() {})
		} else {
			go  /*line :1*/ w4VFts37B3r(rvEj5N7, jSYGwe)
		}
		b96lAKcc :=  /*line :1*/ yESLyde9LHhT.Err()
		h_ljk48Bm := &SoatTK0{
			PY3bIIR:	 /*line :1*/ gCG7rudfaf(b96lAKcc).Error(),
			FIvHQdTCAg:	jKxI_T1CK_p,
			JYeYrD:	b96lAKcc == context.Z201U2N,
		}
		if jHCsMONOrYFN != nil && jHCsMONOrYFN.BZoOYDxA9bpo != nil {
			 /*line :1*/ jHCsMONOrYFN.BZoOYDxA9bpo(nil, false, h_ljk48Bm)
		}
		return nil, h_ljk48Bm
	case yxhs4Z := <-rvEj5N7:
		 /*line :1*/ k76lUDr6.Done()
		 /*line :1*/ jSYGwe()
		h_ljk48Bm := yxhs4Z.ZH7an4
		if h_ljk48Bm != nil {
			if _, d30HIwxht := h_ljk48Bm.(*SoatTK0); !d30HIwxht {
				bJrF0aT := false
				if h_ljk48Bm == context.Z201U2N {
					bJrF0aT = true
				} else if pXUGhLFfB, d30HIwxht := h_ljk48Bm.(yusr33iR); d30HIwxht {
					bJrF0aT =  /*line :1*/ pXUGhLFfB.Timeout()
				}
				h_ljk48Bm = &SoatTK0{
					PY3bIIR:	 /*line :1*/ h_ljk48Bm.Error(),
					FIvHQdTCAg:	jKxI_T1CK_p,
					JYeYrD:	bJrF0aT,
				}
			}
		}
		if jHCsMONOrYFN != nil && jHCsMONOrYFN.BZoOYDxA9bpo != nil {
			md4QSRkO, _ := yxhs4Z.YeKlYE.([]FZJphYv9hV)
			 /*line :1*/ jHCsMONOrYFN.BZoOYDxA9bpo( /*line :1*/ rBUpSfc(md4QSRkO), yxhs4Z.YNGawz, h_ljk48Bm)
		}
		return  /*line :1*/ vqQ2AlEIq(yxhs4Z.YeKlYE, h_ljk48Bm, yxhs4Z.YNGawz)
	}
}



func vqQ2AlEIq(d9AV08pDR any, h_ljk48Bm error, anobGUbhVhF bool) ([]FZJphYv9hV, error) {
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	md4QSRkO := d9AV08pDR.([]FZJphYv9hV)
	if anobGUbhVhF {
		t7IqoZoRNB :=  /*line :1*/ make([]FZJphYv9hV,  /*line :1*/ len(md4QSRkO))
		 /*line :1*/ copy(t7IqoZoRNB, md4QSRkO)
		md4QSRkO = t7IqoZoRNB
	}
	return md4QSRkO, nil
}


func rBUpSfc(md4QSRkO []FZJphYv9hV) []any {
	dRoFccaZ :=  /*line :1*/ make([]any,  /*line :1*/ len(md4QSRkO))
	for eaAUaB7Z, ljsCSk := range md4QSRkO {
		dRoFccaZ[eaAUaB7Z] = ljsCSk
	}
	return dRoFccaZ
}





func S0pLV3A(vsbiWLn7reX, zvX72x8Ih7C string) (mzUgFPgljgC int, h_ljk48Bm error) {
	return  /*line :1*/ Ic4wtIY.LookupPort( /*line :1*/ context.GEcgQP5fzA(), vsbiWLn7reX, zvX72x8Ih7C)
}


func (yxhs4Z *W8Q2tfHAD) LookupPort(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, zvX72x8Ih7C string) (mzUgFPgljgC int, h_ljk48Bm error) {
	mzUgFPgljgC, hsgM8Z :=  /*line :1*/ fJXHJ3xj(zvX72x8Ih7C)
	if hsgM8Z {
		switch vsbiWLn7reX {
		case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
		case "":
			vsbiWLn7reX = "ip"
		default:
			return 0, &DAWLIQHc{Z68v0y: "unknown network", BCCnAgFxG: vsbiWLn7reX}
		}
		mzUgFPgljgC, h_ljk48Bm =  /*line :1*/ yxhs4Z.t_OMUaXanKx(yESLyde9LHhT, vsbiWLn7reX, zvX72x8Ih7C)
		if h_ljk48Bm != nil {
			return 0, h_ljk48Bm
		}
	}
	if 0 > mzUgFPgljgC || mzUgFPgljgC > 65535 {
		return 0, &DAWLIQHc{Z68v0y: "invalid port", BCCnAgFxG: zvX72x8Ih7C}
	}
	return mzUgFPgljgC, nil
}

















func ZqtA1AI3F3(jKxI_T1CK_p string) (jcluQhsSxoD string, h_ljk48Bm error) {
	return  /*line :1*/ Ic4wtIY.LookupCNAME( /*line :1*/ context.GEcgQP5fzA(), jKxI_T1CK_p)
}














func (yxhs4Z *W8Q2tfHAD) LookupCNAME(yESLyde9LHhT context.MBCyA6, jKxI_T1CK_p string) (string, error) {
	jcluQhsSxoD, h_ljk48Bm :=  /*line :1*/ yxhs4Z.eLhY_2IAd(yESLyde9LHhT, jKxI_T1CK_p)
	if h_ljk48Bm != nil {
		return "", h_ljk48Bm
	}
	if ! /*line :1*/ tgYtaZX8VvFU(jcluQhsSxoD) {
		return "", &SoatTK0{PY3bIIR: zMpsF4BoadE, FIvHQdTCAg: jKxI_T1CK_p}
	}
	return jcluQhsSxoD, nil
}















func OD01HjsW(zvX72x8Ih7C, iaebRs5X3, tahgc783Bc string) (jcluQhsSxoD string, md4QSRkO []*LRAj4ahDtDg, h_ljk48Bm error) {
	return  /*line :1*/ Ic4wtIY.LookupSRV( /*line :1*/ context.GEcgQP5fzA(), zvX72x8Ih7C, iaebRs5X3, tahgc783Bc)
}















func (yxhs4Z *W8Q2tfHAD) LookupSRV(yESLyde9LHhT context.MBCyA6, zvX72x8Ih7C, iaebRs5X3, tahgc783Bc string) (string, []*LRAj4ahDtDg, error) {
	jcluQhsSxoD, md4QSRkO, h_ljk48Bm :=  /*line :1*/ yxhs4Z.fe_7xlr(yESLyde9LHhT, zvX72x8Ih7C, iaebRs5X3, tahgc783Bc)
	if h_ljk48Bm != nil {
		return "", nil, h_ljk48Bm
	}
	if jcluQhsSxoD != "" && ! /*line :1*/ tgYtaZX8VvFU(jcluQhsSxoD) {
		return "", nil, &SoatTK0{PY3bIIR: "SRV header name is invalid", FIvHQdTCAg: tahgc783Bc}
	}
	gOssh_YV :=  /*line :1*/ make([]*LRAj4ahDtDg, 0,  /*line :1*/ len(md4QSRkO))
	for _, qxwkC3VxG0 := range md4QSRkO {
		if qxwkC3VxG0 == nil {
			continue
		}
		if ! /*line :1*/ tgYtaZX8VvFU(qxwkC3VxG0.SCFx8T2sDj) {
			continue
		}
		gOssh_YV =  /*line :1*/ append(gOssh_YV, qxwkC3VxG0)
	}
	if  /*line :1*/ len(md4QSRkO) !=  /*line :1*/ len(gOssh_YV) {
		return jcluQhsSxoD, gOssh_YV, &SoatTK0{PY3bIIR: zMpsF4BoadE, FIvHQdTCAg: tahgc783Bc}
	}
	return jcluQhsSxoD, gOssh_YV, nil
}










func Bz9NPuUp4(tahgc783Bc string) ([]*McZ_wBAueqxz, error) {
	return  /*line :1*/ Ic4wtIY.LookupMX( /*line :1*/ context.GEcgQP5fzA(), tahgc783Bc)
}







func (yxhs4Z *W8Q2tfHAD) LookupMX(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]*McZ_wBAueqxz, error) {
	akahx4H, h_ljk48Bm :=  /*line :1*/ yxhs4Z.zr9BDQtsDI(yESLyde9LHhT, tahgc783Bc)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	sU3Bs5KuyHuB :=  /*line :1*/ make([]*McZ_wBAueqxz, 0,  /*line :1*/ len(akahx4H))
	for _, fn_UjuJ := range akahx4H {
		if fn_UjuJ == nil {
			continue
		}
		if ! /*line :1*/ tgYtaZX8VvFU(fn_UjuJ.BgGVZWj8) {
			continue
		}
		sU3Bs5KuyHuB =  /*line :1*/ append(sU3Bs5KuyHuB, fn_UjuJ)
	}
	if  /*line :1*/ len(akahx4H) !=  /*line :1*/ len(sU3Bs5KuyHuB) {
		return sU3Bs5KuyHuB, &SoatTK0{PY3bIIR: zMpsF4BoadE, FIvHQdTCAg: tahgc783Bc}
	}
	return sU3Bs5KuyHuB, nil
}










func BwQx3gKkD(tahgc783Bc string) ([]*AEmYkM1, error) {
	return  /*line :1*/ Ic4wtIY.LookupNS( /*line :1*/ context.GEcgQP5fzA(), tahgc783Bc)
}







func (yxhs4Z *W8Q2tfHAD) LookupNS(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]*AEmYkM1, error) {
	akahx4H, h_ljk48Bm :=  /*line :1*/ yxhs4Z.zoOPu7z1(yESLyde9LHhT, tahgc783Bc)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	hf9YcemgZ7d3 :=  /*line :1*/ make([]*AEmYkM1, 0,  /*line :1*/ len(akahx4H))
	for _, r9NHufqBf4 := range akahx4H {
		if r9NHufqBf4 == nil {
			continue
		}
		if ! /*line :1*/ tgYtaZX8VvFU(r9NHufqBf4.C1_6l1HrdN) {
			continue
		}
		hf9YcemgZ7d3 =  /*line :1*/ append(hf9YcemgZ7d3, r9NHufqBf4)
	}
	if  /*line :1*/ len(akahx4H) !=  /*line :1*/ len(hf9YcemgZ7d3) {
		return hf9YcemgZ7d3, &SoatTK0{PY3bIIR: zMpsF4BoadE, FIvHQdTCAg: tahgc783Bc}
	}
	return hf9YcemgZ7d3, nil
}





func W9keC7n(tahgc783Bc string) ([]string, error) {
	return  /*line :1*/ Ic4wtIY.hxfm94xv( /*line :1*/ context.GEcgQP5fzA(), tahgc783Bc)
}


func (yxhs4Z *W8Q2tfHAD) LookupTXT(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]string, error) {
	return  /*line :1*/ yxhs4Z.hxfm94xv(yESLyde9LHhT, tahgc783Bc)
}













func RNbQpDUCU(qxwkC3VxG0 string) (z_5pYyI7e []string, h_ljk48Bm error) {
	return  /*line :1*/ Ic4wtIY.LookupAddr( /*line :1*/ context.GEcgQP5fzA(), qxwkC3VxG0)
}







func (yxhs4Z *W8Q2tfHAD) LookupAddr(yESLyde9LHhT context.MBCyA6, qxwkC3VxG0 string) ([]string, error) {
	z_5pYyI7e, h_ljk48Bm :=  /*line :1*/ yxhs4Z.tDIayKw(yESLyde9LHhT, qxwkC3VxG0)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	urrBCkFHA :=  /*line :1*/ make([]string, 0,  /*line :1*/ len(z_5pYyI7e))
	for _, tahgc783Bc := range z_5pYyI7e {
		if  /*line :1*/ tgYtaZX8VvFU(tahgc783Bc) {
			urrBCkFHA =  /*line :1*/ append(urrBCkFHA, tahgc783Bc)
		}
	}
	if  /*line :1*/ len(z_5pYyI7e) !=  /*line :1*/ len(urrBCkFHA) {
		return urrBCkFHA, &SoatTK0{PY3bIIR: zMpsF4BoadE, FIvHQdTCAg: qxwkC3VxG0}
	}
	return urrBCkFHA, nil
}




var zMpsF4BoadE = "DNS response contained records which contain invalid names"




func (yxhs4Z *W8Q2tfHAD) tdqObJ_tTokV(yESLyde9LHhT context.MBCyA6, vsbiWLn7reX, c_eHMUdT string) (UJYB3aCQg, error) {
	
	
	
	
	
	var hl8w2lFX UJYB3aCQg
	var h_ljk48Bm error
	if yxhs4Z != nil && yxhs4Z.B5GyJD != nil {
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ yxhs4Z.B5GyJD(yESLyde9LHhT, vsbiWLn7reX, c_eHMUdT)
	} else {
		var ica44Q FySYI4P
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ ica44Q.DialContext(yESLyde9LHhT, vsbiWLn7reX, c_eHMUdT)
	}
	if h_ljk48Bm != nil {
		return nil,  /*line :1*/ gCG7rudfaf(h_ljk48Bm)
	}
	return hl8w2lFX, nil
}










func (yxhs4Z *W8Q2tfHAD) dBGgrS2larpJ(yESLyde9LHhT context.MBCyA6, zvX72x8Ih7C, iaebRs5X3, tahgc783Bc string) (qu6sFD7 string, yIBOJND0k5 []*LRAj4ahDtDg, h_ljk48Bm error) {
	if zvX72x8Ih7C == "" && iaebRs5X3 == "" {
		qu6sFD7 = tahgc783Bc
	} else {
		qu6sFD7 = "_" + zvX72x8Ih7C + "._" + iaebRs5X3 + "." + tahgc783Bc
	}
	fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm :=  /*line :1*/ yxhs4Z.bp6HtmNGS7(yESLyde9LHhT, qu6sFD7, dnsmessage.TypeSRV, nil)
	if h_ljk48Bm != nil {
		return "", nil, h_ljk48Bm
	}
	var jcluQhsSxoD dnsmessage.Toccq2dE3
	for {
		cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.AnswerHeader()
		if h_ljk48Bm == dnsmessage.BbQhjXWx {
			break
		}
		if h_ljk48Bm != nil {
			return "", nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		if cBprzKT8.BxqvzBsK != dnsmessage.TypeSRV {
			if h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SkipAnswer(); h_ljk48Bm != nil {
				return "", nil, &SoatTK0{
					PY3bIIR:	"cannot unmarshal DNS message",
					FIvHQdTCAg:	tahgc783Bc,
					IoaaWnhxbb:	c_eHMUdT,
				}
			}
			continue
		}
		if jcluQhsSxoD.W7uI8JuUD == 0 && cBprzKT8.Pz3XQBChF6.W7uI8JuUD != 0 {
			jcluQhsSxoD = cBprzKT8.Pz3XQBChF6
		}
		ieeN4IUbRFI, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SRVResource()
		if h_ljk48Bm != nil {
			return "", nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		yIBOJND0k5 =  /*line :1*/ append(yIBOJND0k5, &LRAj4ahDtDg{SCFx8T2sDj:  /*line :1*/ ieeN4IUbRFI.HvwgUZxi0.String(), Vb3Y6YRdZxkq: ieeN4IUbRFI.WhxUlP4uIVBk, H5wnRh: ieeN4IUbRFI.AJkmRs98D, BGuDNISZ9: ieeN4IUbRFI.JpLvcVWu})
	}
	 /*line :1*/ a_iF1ZK0(yIBOJND0k5).wYlOfm()
	return  /*line :1*/ jcluQhsSxoD.String(), yIBOJND0k5, nil
}


func (yxhs4Z *W8Q2tfHAD) yhVX7Ebw0(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]*McZ_wBAueqxz, error) {
	fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm :=  /*line :1*/ yxhs4Z.bp6HtmNGS7(yESLyde9LHhT, tahgc783Bc, dnsmessage.TypeMX, nil)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	var d8Wpo4t []*McZ_wBAueqxz
	for {
		cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.AnswerHeader()
		if h_ljk48Bm == dnsmessage.BbQhjXWx {
			break
		}
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		if cBprzKT8.BxqvzBsK != dnsmessage.TypeMX {
			if h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SkipAnswer(); h_ljk48Bm != nil {
				return nil, &SoatTK0{
					PY3bIIR:	"cannot unmarshal DNS message",
					FIvHQdTCAg:	tahgc783Bc,
					IoaaWnhxbb:	c_eHMUdT,
				}
			}
			continue
		}
		fn_UjuJ, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.MXResource()
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		d8Wpo4t =  /*line :1*/ append(d8Wpo4t, &McZ_wBAueqxz{BgGVZWj8:  /*line :1*/ fn_UjuJ.USylcaqC.String(), VG6KZMLF88: fn_UjuJ.DdHoYyYFMHhM})

	}
	 /*line :1*/ rcgRLRupL(d8Wpo4t).wYlOfm()
	return d8Wpo4t, nil
}


func (yxhs4Z *W8Q2tfHAD) aAJTpSPuXHCL(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]*AEmYkM1, error) {
	fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm :=  /*line :1*/ yxhs4Z.bp6HtmNGS7(yESLyde9LHhT, tahgc783Bc, dnsmessage.TypeNS, nil)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	var cjFUVvWN []*AEmYkM1
	for {
		cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.AnswerHeader()
		if h_ljk48Bm == dnsmessage.BbQhjXWx {
			break
		}
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		if cBprzKT8.BxqvzBsK != dnsmessage.TypeNS {
			if h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SkipAnswer(); h_ljk48Bm != nil {
				return nil, &SoatTK0{
					PY3bIIR:	"cannot unmarshal DNS message",
					FIvHQdTCAg:	tahgc783Bc,
					IoaaWnhxbb:	c_eHMUdT,
				}
			}
			continue
		}
		r9NHufqBf4, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.NSResource()
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		cjFUVvWN =  /*line :1*/ append(cjFUVvWN, &AEmYkM1{C1_6l1HrdN:  /*line :1*/ r9NHufqBf4.JccbZj.String()})
	}
	return cjFUVvWN, nil
}


func (yxhs4Z *W8Q2tfHAD) qR_lwcoTokk(yESLyde9LHhT context.MBCyA6, tahgc783Bc string) ([]string, error) {
	fMPVz2iGiyq, c_eHMUdT, h_ljk48Bm :=  /*line :1*/ yxhs4Z.bp6HtmNGS7(yESLyde9LHhT, tahgc783Bc, dnsmessage.TypeTXT, nil)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	var gl9eQhAnSrZ []string
	for {
		cBprzKT8, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.AnswerHeader()
		if h_ljk48Bm == dnsmessage.BbQhjXWx {
			break
		}
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}
		if cBprzKT8.BxqvzBsK != dnsmessage.TypeTXT {
			if h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.SkipAnswer(); h_ljk48Bm != nil {
				return nil, &SoatTK0{
					PY3bIIR:	"cannot unmarshal DNS message",
					FIvHQdTCAg:	tahgc783Bc,
					IoaaWnhxbb:	c_eHMUdT,
				}
			}
			continue
		}
		tyFajMMBa, h_ljk48Bm :=  /*line :1*/ fMPVz2iGiyq.TXTResource()
		if h_ljk48Bm != nil {
			return nil, &SoatTK0{
				PY3bIIR:	"cannot unmarshal DNS message",
				FIvHQdTCAg:	tahgc783Bc,
				IoaaWnhxbb:	c_eHMUdT,
			}
		}

		doauF8Y := 0
		for _, dRoFccaZ := range tyFajMMBa.FhHS8hCoPF2m {
			doauF8Y +=  /*line :1*/ len(dRoFccaZ)
		}
		bUN5azO7bk :=  /*line :1*/ make([]byte, 0, doauF8Y)
		for _, dRoFccaZ := range tyFajMMBa.FhHS8hCoPF2m {
			bUN5azO7bk =  /*line :1*/ append(bUN5azO7bk, dRoFccaZ...)
		}
		if  /*line :1*/ len(gl9eQhAnSrZ) == 0 {
			gl9eQhAnSrZ =  /*line :1*/ make([]string, 0, 1)
		}
		gl9eQhAnSrZ =  /*line :1*/ append(gl9eQhAnSrZ,  /*line :1*/ string(bUN5azO7bk))
	}
	return gl9eQhAnSrZ, nil
}

func cCj_s1v63(hgyhkxJ5TRr []dnsmessage.Y21iaz2cH_hd) (string, error) {
	if  /*line :1*/ len(hgyhkxJ5TRr) == 0 {
		return "",  /*line :1*/ errors.Su6g6hRoi9X("no CNAME record received")
	}
	hl8w2lFX, d30HIwxht := hgyhkxJ5TRr[0].GaL3Hix.(*dnsmessage.Wi6AD2AtGlVJ)
	if !d30HIwxht {
		return "",  /*line :1*/ errors.Su6g6hRoi9X("could not parse CNAME record")
	}
	return  /*line :1*/ hl8w2lFX.SiT_7E895B87.String(), nil
}
