//line :1
package MjXXMR

import (
	errors "iAHoxjmM"
	race "V0p_iv"
	windows "LdptURlN"
	io "xI9ai2djaJ2"
	sync "sync"
	syscall "bUKeamF"
	utf16 "DtJCLKevRp"
	utf8 "CuAc0pPZwf"
	"unsafe"
)

var (
	e8v415Vy9	error
	gacktcmFP	uint64
)

var wkVH8K bool	





func vFYQYkF7S() {
	dtU8tBiUGS4 :=  /*line :1*/ syscall.JTYwZ3()
	if dtU8tBiUGS4 != nil {
		return
	}
	b3NyBNpg := [2]int32{syscall.IPPROTO_TCP, 0}
	var qMghkaiHXh [32]syscall.ZEPlE2KjP71
	ajpah49Ud :=  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(qMghkaiHXh))
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ syscall.AMZXrhL3Tt8(&b3NyBNpg[0], &qMghkaiHXh[0], &ajpah49Ud)
	if dtU8tBiUGS4 != nil {
		return
	}
	for hB6DM2iT :=  /*line :1*/ int32(0); hB6DM2iT < mnkyfx; hB6DM2iT++ {
		if qMghkaiHXh[hB6DM2iT].KRYxJfq&syscall.XP1_IFS_HANDLES == 0 {
			return
		}
	}
	wkVH8K = true
}

func init() {
	var wmR0FCrSvE syscall.RqD0UlyP
	gRlKyDVfhjh :=  /*line :1*/ syscall.Kq4jyT( /*line :1*/ uint32(0x202), &wmR0FCrSvE)
	if gRlKyDVfhjh != nil {
		e8v415Vy9 = gRlKyDVfhjh
	}
	 /*line :1*/ vFYQYkF7S()
}


type lVE75f25P struct {
	
	
	gZWJ1u9D4auz	syscall.ZaNm5QSYC9

	
	c_aCB5q_J8	uintptr
	nay9bs45e0S	int32
	da7rs9c4esa	int32
	sSqPaaxx	uint32

	
	ef3m3FV3JP	*ENfmDmMaozH
	_vTEq5NaQ	syscall.QbWhAp5CqTG
	yW8dpq5f6K	windows.BXk_cV9l_bLI
	rVuSBLM	syscall.S4iroLVT4
	rKvfha6sO	*syscall.IXy5oynSaLQM
	jQwc0Q60T	int32
	azj1aOHqks	syscall.SRlvVjrYa
	oXBY5N	uint32
	awWc3wwV	[]syscall.QbWhAp5CqTG
}

func (aeapDF *lVE75f25P) InitBuf(qMghkaiHXh []byte) {
	aeapDF._vTEq5NaQ.Ix58rFqPW =  /*line :1*/ uint32( /*line :1*/ len(qMghkaiHXh))
	aeapDF._vTEq5NaQ.AcsIulb2fH4 = nil
	if  /*line :1*/ len(qMghkaiHXh) != 0 {
		aeapDF._vTEq5NaQ.AcsIulb2fH4 = &qMghkaiHXh[0]
	}
}

func (aeapDF *lVE75f25P) InitBufs(qMghkaiHXh *[][]byte) {
	if aeapDF.awWc3wwV == nil {
		aeapDF.awWc3wwV =  /*line :1*/ make([]syscall.QbWhAp5CqTG, 0,  /*line :1*/ len(*qMghkaiHXh))
	} else {
		aeapDF.awWc3wwV = aeapDF.awWc3wwV[:0]
	}
	for _, mvTjugW := range *qMghkaiHXh {
		if  /*line :1*/ len(mvTjugW) == 0 {
			aeapDF.awWc3wwV =  /*line :1*/ append(aeapDF.awWc3wwV, syscall.QbWhAp5CqTG{})
			continue
		}
		for  /*line :1*/ len(mvTjugW) > maxRW {
			aeapDF.awWc3wwV =  /*line :1*/ append(aeapDF.awWc3wwV, syscall.QbWhAp5CqTG{Ix58rFqPW: maxRW, AcsIulb2fH4: &mvTjugW[0]})
			mvTjugW = mvTjugW[maxRW:]
		}
		if  /*line :1*/ len(mvTjugW) > 0 {
			aeapDF.awWc3wwV =  /*line :1*/ append(aeapDF.awWc3wwV, syscall.QbWhAp5CqTG{Ix58rFqPW:  /*line :1*/ uint32( /*line :1*/ len(mvTjugW)), AcsIulb2fH4: &mvTjugW[0]})
		}
	}
}



func (aeapDF *lVE75f25P) ClearBufs() {
	for hB6DM2iT := range aeapDF.awWc3wwV {
		aeapDF.awWc3wwV[hB6DM2iT].AcsIulb2fH4 = nil
	}
	aeapDF.awWc3wwV = aeapDF.awWc3wwV[:0]
}

func (aeapDF *lVE75f25P) InitMsg(yZmyVky9Vg []byte, rcxhh3yIRuRD []byte) {
	 /*line :1*/ aeapDF.InitBuf(yZmyVky9Vg)
	aeapDF.yW8dpq5f6K.MD9Fwj = &aeapDF._vTEq5NaQ
	aeapDF.yW8dpq5f6K.Du6szpnO = 1

	aeapDF.yW8dpq5f6K.ClAqGya = nil
	aeapDF.yW8dpq5f6K.MeKqL5 = 0

	aeapDF.yW8dpq5f6K.WDy42Vy = 0
	aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.Ix58rFqPW =  /*line :1*/ uint32( /*line :1*/ len(rcxhh3yIRuRD))
	aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.AcsIulb2fH4 = nil
	if  /*line :1*/ len(rcxhh3yIRuRD) != 0 {
		aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.AcsIulb2fH4 = &rcxhh3yIRuRD[0]
	}
}





func sX12A2hZ(aeapDF *lVE75f25P, nWHIxSmz func(aeapDF *lVE75f25P) error) (int, error) {
	if aeapDF.ef3m3FV3JP.yw4UOLmW.gwF0UKHcgp == 0 {
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("internal error: polling on unsupported descriptor type")
	}

	s_ZR_8 := aeapDF.ef3m3FV3JP

	dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yw4UOLmW.n8O_Sp( /*line :1*/ int(aeapDF.nay9bs45e0S), s_ZR_8.l3fWSyWa)
	if dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}

	dtU8tBiUGS4 =  /*line :1*/ nWHIxSmz(aeapDF)
	switch dtU8tBiUGS4 {
	case nil:

		if aeapDF.ef3m3FV3JP.m5pG8SVivK {

			return  /*line :1*/ int(aeapDF.sSqPaaxx), nil
		}

	case syscall.ERROR_IO_PENDING:

		dtU8tBiUGS4 = nil
	default:
		return 0, dtU8tBiUGS4
	}

	dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.yw4UOLmW.dGMuWJsbCVm( /*line :1*/ int(aeapDF.nay9bs45e0S), s_ZR_8.l3fWSyWa)
	if dtU8tBiUGS4 == nil {

		if aeapDF.da7rs9c4esa != 0 {
			dtU8tBiUGS4 =  /*line :1*/ syscall.O9Mga3(aeapDF.da7rs9c4esa)

			if dtU8tBiUGS4 == syscall.ERROR_MORE_DATA || dtU8tBiUGS4 == windows.WSAEMSGSIZE {
				return  /*line :1*/ int(aeapDF.sSqPaaxx), dtU8tBiUGS4
			}
			return 0, dtU8tBiUGS4
		}
		return  /*line :1*/ int(aeapDF.sSqPaaxx), nil
	}

	m3bQdi0JXmRp := dtU8tBiUGS4
	switch m3bQdi0JXmRp {
	case DA6gHYGx1nT, IQcIMbVlGewk, ZV68d6_RJG:

	default:
		 /*line :1*/ panic("unexpected runtime.netpoll error: " +  /*line :1*/ m3bQdi0JXmRp.Error())
	}

	dtU8tBiUGS4 =  /*line :1*/ syscall.QaAWdiVWBp(s_ZR_8.X8OEsVYSby, &aeapDF.gZWJ1u9D4auz)

	if dtU8tBiUGS4 != nil && dtU8tBiUGS4 != syscall.ERROR_NOT_FOUND {

		 /*line :1*/ panic(dtU8tBiUGS4)
	}

	 /*line :1*/ s_ZR_8.yw4UOLmW.kaTGzw( /*line :1*/ int(aeapDF.nay9bs45e0S))
	if aeapDF.da7rs9c4esa != 0 {
		dtU8tBiUGS4 =  /*line :1*/ syscall.O9Mga3(aeapDF.da7rs9c4esa)
		if dtU8tBiUGS4 == syscall.ERROR_OPERATION_ABORTED {
			dtU8tBiUGS4 = m3bQdi0JXmRp
		}
		return 0, dtU8tBiUGS4
	}

	return  /*line :1*/ int(aeapDF.sSqPaaxx), nil
}



type ENfmDmMaozH struct {
	
	eYtKOCiM	hvIBHvdqmStf

	
	X8OEsVYSby	syscall.SRlvVjrYa

	
	x0rW_mPQcb	lVE75f25P
	
	aQ9pa1ac	lVE75f25P

	
	yw4UOLmW	rDHgTN

	
	ezOne7Au	sync.DIRWe11KYlYa

	
	oG7FMdLiK	[]byte	
	oNulFq3a	[]uint16	
	tv0vaX	[]byte	
	dS5gNC	int	

	
	g_7QwmbfWu	uint32

	m5pG8SVivK	bool

	
	
	MDfFexom	bool

	
	
	ZslYVRp1qJ	bool

	
	l3fWSyWa	bool

	
	sTN5PtBj1N	hUwoPh
}


type hUwoPh byte

const (
	kindNet	hUwoPh	= iota
	kindFile
	kindConsole
	kindPipe
)


var k6i8wZg func(pnglWMA string, s_ZR_8 *ENfmDmMaozH, dtU8tBiUGS4 error)






func (s_ZR_8 *ENfmDmMaozH) Init(pnglWMA string, gin2cUpW bool) (string, error) {
	if e8v415Vy9 != nil {
		return "", e8v415Vy9
	}

	switch pnglWMA {
	case "file", "dir":
		s_ZR_8.sTN5PtBj1N = kindFile
	case "console":
		s_ZR_8.sTN5PtBj1N = kindConsole
	case "pipe":
		s_ZR_8.sTN5PtBj1N = kindPipe
	case "tcp", "tcp4", "tcp6",
		"udp", "udp4", "udp6",
		"ip", "ip4", "ip6",
		"unix", "unixgram", "unixpacket":
		s_ZR_8.sTN5PtBj1N = kindNet
	default:
		return "",  /*line :1*/ errors.Su6g6hRoi9X("internal error: unknown network type " + pnglWMA)
	}
	s_ZR_8.l3fWSyWa = s_ZR_8.sTN5PtBj1N != kindNet

	var dtU8tBiUGS4 error
	if gin2cUpW {

		dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.yw4UOLmW.init(s_ZR_8)
	}
	if k6i8wZg != nil {
		 /*line :1*/ k6i8wZg(pnglWMA, s_ZR_8, dtU8tBiUGS4)
	}
	if dtU8tBiUGS4 != nil {
		return "", dtU8tBiUGS4
	}
	if gin2cUpW && wkVH8K {

		a5fb2Y :=  /*line :1*/ uint8(syscall.FILE_SKIP_SET_EVENT_ON_HANDLE)
		switch pnglWMA {
		case "tcp", "tcp4", "tcp6",
			"udp", "udp4", "udp6":
			a5fb2Y |= syscall.FILE_SKIP_COMPLETION_PORT_ON_SUCCESS
		}
		dtU8tBiUGS4 :=  /*line :1*/ syscall.F2ye7PjCo5l(s_ZR_8.X8OEsVYSby, a5fb2Y)
		if dtU8tBiUGS4 == nil && a5fb2Y&syscall.FILE_SKIP_COMPLETION_PORT_ON_SUCCESS != 0 {
			s_ZR_8.m5pG8SVivK = true
		}
	}

	switch pnglWMA {
	case "udp", "udp4", "udp6":
		cioVzGtzZ5 :=  /*line :1*/ uint32(0)
		k3rghaYjK :=  /*line :1*/ uint32(0)
		mxYBfG :=  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(k3rghaYjK))
		dtU8tBiUGS4 :=  /*line :1*/ syscall.YvCdY02(s_ZR_8.X8OEsVYSby, syscall.SIO_UDP_CONNRESET, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&k3rghaYjK)), mxYBfG, nil, 0, &cioVzGtzZ5, nil, 0)
		if dtU8tBiUGS4 != nil {
			return "wsaioctl", dtU8tBiUGS4
		}
	}
	s_ZR_8.x0rW_mPQcb.nay9bs45e0S = 'r'
	s_ZR_8.aQ9pa1ac.nay9bs45e0S = 'w'
	s_ZR_8.x0rW_mPQcb.ef3m3FV3JP = s_ZR_8
	s_ZR_8.aQ9pa1ac.ef3m3FV3JP = s_ZR_8
	s_ZR_8.x0rW_mPQcb.c_aCB5q_J8 = s_ZR_8.yw4UOLmW.gwF0UKHcgp
	s_ZR_8.aQ9pa1ac.c_aCB5q_J8 = s_ZR_8.yw4UOLmW.gwF0UKHcgp
	return "", nil
}

func (s_ZR_8 *ENfmDmMaozH) twRFlT6Xc() error {
	if s_ZR_8.X8OEsVYSby == syscall.InvalidHandle {
		return syscall.EINVAL
	}

	 /*line :1*/ s_ZR_8.yw4UOLmW.zVZzO_nA_()
	var dtU8tBiUGS4 error
	switch s_ZR_8.sTN5PtBj1N {
	case kindNet:

		dtU8tBiUGS4 =  /*line :1*/ KvwyJafTkVwj(s_ZR_8.X8OEsVYSby)
	default:
		dtU8tBiUGS4 =  /*line :1*/ syscall.FhKJLgXjwG(s_ZR_8.X8OEsVYSby)
	}
	s_ZR_8.X8OEsVYSby = syscall.InvalidHandle
	 /*line :1*/ g6gFf1(&s_ZR_8.g_7QwmbfWu)
	return dtU8tBiUGS4
}



func (s_ZR_8 *ENfmDmMaozH) Close() error {
	if ! /*line :1*/ s_ZR_8.eYtKOCiM.wMwcmDUo() {
		return  /*line :1*/ amQJqF(s_ZR_8.l3fWSyWa)
	}
	if s_ZR_8.sTN5PtBj1N == kindPipe {
		 /*line :1*/ syscall.QaAWdiVWBp(s_ZR_8.X8OEsVYSby, nil)
	}

	 /*line :1*/ s_ZR_8.yw4UOLmW.twNLQiAMSc()
	dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.xmYUot4V()

	 /*line :1*/ jPXZV7q0q5(&s_ZR_8.g_7QwmbfWu)
	return dtU8tBiUGS4
}




const maxRW = 1 << 30	


func (s_ZR_8 *ENfmDmMaozH) Read(qMghkaiHXh []byte) (int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()

	if  /*line :1*/ len(qMghkaiHXh) > maxRW {
		qMghkaiHXh = qMghkaiHXh[:maxRW]
	}

	var mnkyfx int
	var dtU8tBiUGS4 error
	if s_ZR_8.l3fWSyWa {
		 /*line :1*/ s_ZR_8.ezOne7Au.Lock()
		defer  /*line :1*/ s_ZR_8.ezOne7Au.Unlock()
		switch s_ZR_8.sTN5PtBj1N {
		case kindConsole:
			mnkyfx, dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.aWCNTOFB(qMghkaiHXh)
		default:
			mnkyfx, dtU8tBiUGS4 =  /*line :1*/ syscall.XJ0LdRoyxl_l(s_ZR_8.X8OEsVYSby, qMghkaiHXh)
			if s_ZR_8.sTN5PtBj1N == kindPipe && dtU8tBiUGS4 == syscall.ERROR_OPERATION_ABORTED {

				dtU8tBiUGS4 = IQcIMbVlGewk
			}
		}
		if dtU8tBiUGS4 != nil {
			mnkyfx = 0
		}
	} else {
		aeapDF := &s_ZR_8.x0rW_mPQcb
		 /*line :1*/ aeapDF.InitBuf(qMghkaiHXh)
		mnkyfx, dtU8tBiUGS4 =  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ syscall.J67wPAoai1Kb(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, &aeapDF.oXBY5N, &aeapDF.gZWJ1u9D4auz, nil)
		})
		if race.Enabled {
			 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&gacktcmFP))
		}
	}
	if  /*line :1*/ len(qMghkaiHXh) != 0 {
		dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.cKC2EMAVUS2(mnkyfx, dtU8tBiUGS4)
	}
	return mnkyfx, dtU8tBiUGS4
}

var Vc17aFjlyi = syscall.Jkhz6hQ6zGAR	




func (s_ZR_8 *ENfmDmMaozH) aWCNTOFB(mvTjugW []byte) (int, error) {
	if  /*line :1*/ len(mvTjugW) == 0 {
		return 0, nil
	}

	if s_ZR_8.oNulFq3a == nil {

		s_ZR_8.oNulFq3a =  /*line :1*/ make([]uint16, 0, 10000)
		s_ZR_8.tv0vaX =  /*line :1*/ make([]byte, 0, 4* /*line :1*/ cap(s_ZR_8.oNulFq3a))
	}

	for s_ZR_8.dS5gNC >=  /*line :1*/ len(s_ZR_8.tv0vaX) {
		mnkyfx :=  /*line :1*/ cap(s_ZR_8.oNulFq3a) -  /*line :1*/ len(s_ZR_8.oNulFq3a)
		if mnkyfx >  /*line :1*/ len(mvTjugW) {
			mnkyfx =  /*line :1*/ len(mvTjugW)
		}
		var fuM4gU uint32
		dtU8tBiUGS4 :=  /*line :1*/ Vc17aFjlyi(s_ZR_8.X8OEsVYSby, &s_ZR_8.oNulFq3a[: /*line :1*/ len(s_ZR_8.oNulFq3a)+1][ /*line :1*/ len(s_ZR_8.oNulFq3a)],  /*line :1*/ uint32(mnkyfx), &fuM4gU, nil)
		if dtU8tBiUGS4 != nil {
			return 0, dtU8tBiUGS4
		}
		d0ly0sRaYwoo := s_ZR_8.oNulFq3a[: /*line :1*/ len(s_ZR_8.oNulFq3a)+ /*line :1*/ int(fuM4gU)]
		s_ZR_8.oNulFq3a = s_ZR_8.oNulFq3a[:0]
		qMghkaiHXh := s_ZR_8.tv0vaX[:0]
		for hB6DM2iT := 0; hB6DM2iT <  /*line :1*/ len(d0ly0sRaYwoo); hB6DM2iT++ {
			mP5l3KzGhO7G :=  /*line :1*/ rune(d0ly0sRaYwoo[hB6DM2iT])
			if  /*line :1*/ utf16.I0RIjMaMPDrI(mP5l3KzGhO7G) {
				if hB6DM2iT+1 ==  /*line :1*/ len(d0ly0sRaYwoo) {
					if fuM4gU > 0 {

						s_ZR_8.oNulFq3a = s_ZR_8.oNulFq3a[:1]
						s_ZR_8.oNulFq3a[0] =  /*line :1*/ uint16(mP5l3KzGhO7G)
						break
					}
					mP5l3KzGhO7G = utf8.RuneError
				} else {
					mP5l3KzGhO7G =  /*line :1*/ utf16.Rw3auy7U(mP5l3KzGhO7G,  /*line :1*/ rune(d0ly0sRaYwoo[hB6DM2iT+1]))
					if mP5l3KzGhO7G != utf8.RuneError {
						hB6DM2iT++
					}
				}
			}
			qMghkaiHXh =  /*line :1*/ utf8.Ht7oMzd(qMghkaiHXh, mP5l3KzGhO7G)
		}
		s_ZR_8.tv0vaX = qMghkaiHXh
		s_ZR_8.dS5gNC = 0
		if fuM4gU == 0 {
			break
		}
	}

	iEBoOae8j := s_ZR_8.tv0vaX[s_ZR_8.dS5gNC:]
	var hB6DM2iT int
	for hB6DM2iT = 0; hB6DM2iT <  /*line :1*/ len(iEBoOae8j) && hB6DM2iT <  /*line :1*/ len(mvTjugW); hB6DM2iT++ {
		cqimyW17pD := iEBoOae8j[hB6DM2iT]
		if cqimyW17pD == 0x1A {
			if hB6DM2iT == 0 {
				s_ZR_8.dS5gNC++
			}
			break
		}
		mvTjugW[hB6DM2iT] = cqimyW17pD
	}
	s_ZR_8.dS5gNC += hB6DM2iT
	return hB6DM2iT, nil
}


func (s_ZR_8 *ENfmDmMaozH) Pread(mvTjugW []byte, ql_YwvPFhh int64) (int, error) {
	if s_ZR_8.sTN5PtBj1N == kindPipe {

		return 0, syscall.ESPIPE
	}

	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()

	if  /*line :1*/ len(mvTjugW) > maxRW {
		mvTjugW = mvTjugW[:maxRW]
	}

	 /*line :1*/ s_ZR_8.ezOne7Au.Lock()
	defer  /*line :1*/ s_ZR_8.ezOne7Au.Unlock()
	nM73r8CXTbLl, gRlKyDVfhjh :=  /*line :1*/ syscall.O4FqtrPqbHRY(s_ZR_8.X8OEsVYSby, 0, io.SeekCurrent)
	if gRlKyDVfhjh != nil {
		return 0, gRlKyDVfhjh
	}
	defer  /*line :1*/ syscall.O4FqtrPqbHRY(s_ZR_8.X8OEsVYSby, nM73r8CXTbLl, io.SeekStart)
	aeapDF := syscall.ZaNm5QSYC9{
		BvaSyl:	 /*line :1*/ uint32(ql_YwvPFhh >> 32),
		BEJ63D:	 /*line :1*/ uint32(ql_YwvPFhh),
	}
	var oROiQ7SfWWD6 uint32
	gRlKyDVfhjh =  /*line :1*/ syscall.Sg86ugD(s_ZR_8.X8OEsVYSby, mvTjugW, &oROiQ7SfWWD6, &aeapDF)
	if gRlKyDVfhjh != nil {
		oROiQ7SfWWD6 = 0
		if gRlKyDVfhjh == syscall.ERROR_HANDLE_EOF {
			gRlKyDVfhjh = io.K5Sqco
		}
	}
	if  /*line :1*/ len(mvTjugW) != 0 {
		gRlKyDVfhjh =  /*line :1*/ s_ZR_8.cKC2EMAVUS2( /*line :1*/ int(oROiQ7SfWWD6), gRlKyDVfhjh)
	}
	return  /*line :1*/ int(oROiQ7SfWWD6), gRlKyDVfhjh
}


func (s_ZR_8 *ENfmDmMaozH) ReadFrom(qMghkaiHXh []byte) (int, syscall.S4iroLVT4, error) {
	if  /*line :1*/ len(qMghkaiHXh) == 0 {
		return 0, nil, nil
	}
	if  /*line :1*/ len(qMghkaiHXh) > maxRW {
		qMghkaiHXh = qMghkaiHXh[:maxRW]
	}
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return 0, nil, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()
	aeapDF := &s_ZR_8.x0rW_mPQcb
	 /*line :1*/ aeapDF.InitBuf(qMghkaiHXh)
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		if aeapDF.rKvfha6sO == nil {
			aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
		}
		aeapDF.jQwc0Q60T =  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*aeapDF.rKvfha6sO))
		return  /*line :1*/ syscall.AYS5UpcaP(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, &aeapDF.oXBY5N, aeapDF.rKvfha6sO, &aeapDF.jQwc0Q60T, &aeapDF.gZWJ1u9D4auz, nil)
	})
	dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.cKC2EMAVUS2(mnkyfx, dtU8tBiUGS4)
	if dtU8tBiUGS4 != nil {
		return mnkyfx, nil, dtU8tBiUGS4
	}
	iAptnzr4Hy4m, _ :=  /*line :1*/ aeapDF.rKvfha6sO.Sockaddr()
	return mnkyfx, iAptnzr4Hy4m, nil
}


func (s_ZR_8 *ENfmDmMaozH) ReadFromInet4(qMghkaiHXh []byte, _89sWOk *syscall.Hx8lqxSkV) (int, error) {
	if  /*line :1*/ len(qMghkaiHXh) == 0 {
		return 0, nil
	}
	if  /*line :1*/ len(qMghkaiHXh) > maxRW {
		qMghkaiHXh = qMghkaiHXh[:maxRW]
	}
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()
	aeapDF := &s_ZR_8.x0rW_mPQcb
	 /*line :1*/ aeapDF.InitBuf(qMghkaiHXh)
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		if aeapDF.rKvfha6sO == nil {
			aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
		}
		aeapDF.jQwc0Q60T =  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*aeapDF.rKvfha6sO))
		return  /*line :1*/ syscall.AYS5UpcaP(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, &aeapDF.oXBY5N, aeapDF.rKvfha6sO, &aeapDF.jQwc0Q60T, &aeapDF.gZWJ1u9D4auz, nil)
	})
	dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.cKC2EMAVUS2(mnkyfx, dtU8tBiUGS4)
	if dtU8tBiUGS4 != nil {
		return mnkyfx, dtU8tBiUGS4
	}
	 /*line :1*/ dcDIm2u(aeapDF.rKvfha6sO, _89sWOk)
	return mnkyfx, dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) ReadFromInet6(qMghkaiHXh []byte, uWZEc5bfl2 *syscall.HKTAy7_BSU2) (int, error) {
	if  /*line :1*/ len(qMghkaiHXh) == 0 {
		return 0, nil
	}
	if  /*line :1*/ len(qMghkaiHXh) > maxRW {
		qMghkaiHXh = qMghkaiHXh[:maxRW]
	}
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()
	aeapDF := &s_ZR_8.x0rW_mPQcb
	 /*line :1*/ aeapDF.InitBuf(qMghkaiHXh)
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		if aeapDF.rKvfha6sO == nil {
			aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
		}
		aeapDF.jQwc0Q60T =  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*aeapDF.rKvfha6sO))
		return  /*line :1*/ syscall.AYS5UpcaP(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, &aeapDF.oXBY5N, aeapDF.rKvfha6sO, &aeapDF.jQwc0Q60T, &aeapDF.gZWJ1u9D4auz, nil)
	})
	dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.cKC2EMAVUS2(mnkyfx, dtU8tBiUGS4)
	if dtU8tBiUGS4 != nil {
		return mnkyfx, dtU8tBiUGS4
	}
	 /*line :1*/ hrB5fkV5Ww(aeapDF.rKvfha6sO, uWZEc5bfl2)
	return mnkyfx, dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) Write(qMghkaiHXh []byte) (int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()
	if s_ZR_8.l3fWSyWa {
		 /*line :1*/ s_ZR_8.ezOne7Au.Lock()
		defer  /*line :1*/ s_ZR_8.ezOne7Au.Unlock()
	}

	iuvUgMayz := 0
	for  /*line :1*/ len(qMghkaiHXh) > 0 {
		mvTjugW := qMghkaiHXh
		if  /*line :1*/ len(mvTjugW) > maxRW {
			mvTjugW = mvTjugW[:maxRW]
		}
		var mnkyfx int
		var dtU8tBiUGS4 error
		if s_ZR_8.l3fWSyWa {
			switch s_ZR_8.sTN5PtBj1N {
			case kindConsole:
				mnkyfx, dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.fW8gNj5hGG0(mvTjugW)
			default:
				mnkyfx, dtU8tBiUGS4 =  /*line :1*/ syscall.ZsMbSV2(s_ZR_8.X8OEsVYSby, mvTjugW)
				if s_ZR_8.sTN5PtBj1N == kindPipe && dtU8tBiUGS4 == syscall.ERROR_OPERATION_ABORTED {

					dtU8tBiUGS4 = IQcIMbVlGewk
				}
			}
			if dtU8tBiUGS4 != nil {
				mnkyfx = 0
			}
		} else {
			if race.Enabled {
				 /*line :1*/ race.AQrXiLDh3RTu( /*line :1*/ unsafe.Pointer(&gacktcmFP))
			}
			aeapDF := &s_ZR_8.aQ9pa1ac
			 /*line :1*/ aeapDF.InitBuf(mvTjugW)
			mnkyfx, dtU8tBiUGS4 =  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
				return  /*line :1*/ syscall.Ky9hDuJe(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, 0, &aeapDF.gZWJ1u9D4auz, nil)
			})
		}
		iuvUgMayz += mnkyfx
		if dtU8tBiUGS4 != nil {
			return iuvUgMayz, dtU8tBiUGS4
		}
		qMghkaiHXh = qMghkaiHXh[mnkyfx:]
	}
	return iuvUgMayz, nil
}



func (s_ZR_8 *ENfmDmMaozH) fW8gNj5hGG0(mvTjugW []byte) (int, error) {
	mnkyfx :=  /*line :1*/ len(mvTjugW)
	ulNeG0l :=  /*line :1*/ make([]rune, 0, 256)
	if  /*line :1*/ len(s_ZR_8.oG7FMdLiK) > 0 {
		mvTjugW =  /*line :1*/ append(s_ZR_8.oG7FMdLiK, mvTjugW...)
		s_ZR_8.oG7FMdLiK = nil

	}
	for  /*line :1*/ len(mvTjugW) >= utf8.UTFMax ||  /*line :1*/ utf8.TAy5Jt8FB7p4(mvTjugW) {
		mP5l3KzGhO7G, gwLmzX :=  /*line :1*/ utf8.EicfpCPn(mvTjugW)
		ulNeG0l =  /*line :1*/ append(ulNeG0l, mP5l3KzGhO7G)
		mvTjugW = mvTjugW[gwLmzX:]
	}
	if  /*line :1*/ len(mvTjugW) > 0 {
		s_ZR_8.oG7FMdLiK =  /*line :1*/ make([]byte,  /*line :1*/ len(mvTjugW))
		 /*line :1*/ copy(s_ZR_8.oG7FMdLiK, mvTjugW)
	}
	
	
	
	const maxWrite = 16000
	for  /*line :1*/ len(ulNeG0l) > 0 {
		vhCWBBljG :=  /*line :1*/ len(ulNeG0l)
		if vhCWBBljG > maxWrite {
			vhCWBBljG = maxWrite
		}
		cakTHVx6EwC := ulNeG0l[:vhCWBBljG]
		ulNeG0l = ulNeG0l[vhCWBBljG:]
		d0ly0sRaYwoo :=  /*line :1*/ utf16.J8eAwOfFc(cakTHVx6EwC)
		for  /*line :1*/ len(d0ly0sRaYwoo) > 0 {
			var xB7wJxWigaca uint32
			dtU8tBiUGS4 :=  /*line :1*/ syscall.Djt4smg(s_ZR_8.X8OEsVYSby, &d0ly0sRaYwoo[0],  /*line :1*/ uint32( /*line :1*/ len(d0ly0sRaYwoo)), &xB7wJxWigaca, nil)
			if dtU8tBiUGS4 != nil {
				return 0, dtU8tBiUGS4
			}
			d0ly0sRaYwoo = d0ly0sRaYwoo[xB7wJxWigaca:]
		}
	}
	return mnkyfx, nil
}


func (s_ZR_8 *ENfmDmMaozH) Pwrite(qMghkaiHXh []byte, ql_YwvPFhh int64) (int, error) {
	if s_ZR_8.sTN5PtBj1N == kindPipe {

		return 0, syscall.ESPIPE
	}

	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()

	 /*line :1*/ s_ZR_8.ezOne7Au.Lock()
	defer  /*line :1*/ s_ZR_8.ezOne7Au.Unlock()
	nM73r8CXTbLl, gRlKyDVfhjh :=  /*line :1*/ syscall.O4FqtrPqbHRY(s_ZR_8.X8OEsVYSby, 0, io.SeekCurrent)
	if gRlKyDVfhjh != nil {
		return 0, gRlKyDVfhjh
	}
	defer  /*line :1*/ syscall.O4FqtrPqbHRY(s_ZR_8.X8OEsVYSby, nM73r8CXTbLl, io.SeekStart)

	iuvUgMayz := 0
	for  /*line :1*/ len(qMghkaiHXh) > 0 {
		mvTjugW := qMghkaiHXh
		if  /*line :1*/ len(mvTjugW) > maxRW {
			mvTjugW = mvTjugW[:maxRW]
		}
		var mnkyfx uint32
		aeapDF := syscall.ZaNm5QSYC9{
			BvaSyl:	 /*line :1*/ uint32(ql_YwvPFhh >> 32),
			BEJ63D:	 /*line :1*/ uint32(ql_YwvPFhh),
		}
		gRlKyDVfhjh =  /*line :1*/ syscall.DaWZrhl(s_ZR_8.X8OEsVYSby, mvTjugW, &mnkyfx, &aeapDF)
		iuvUgMayz +=  /*line :1*/ int(mnkyfx)
		if gRlKyDVfhjh != nil {
			return iuvUgMayz, gRlKyDVfhjh
		}
		qMghkaiHXh = qMghkaiHXh[mnkyfx:]
		ql_YwvPFhh +=  /*line :1*/ int64(mnkyfx)
	}
	return iuvUgMayz, nil
}


func (s_ZR_8 *ENfmDmMaozH) Writev(qMghkaiHXh *[][]byte) (int64, error) {
	if  /*line :1*/ len(*qMghkaiHXh) == 0 {
		return 0, nil
	}
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()
	if race.Enabled {
		 /*line :1*/ race.AQrXiLDh3RTu( /*line :1*/ unsafe.Pointer(&gacktcmFP))
	}
	aeapDF := &s_ZR_8.aQ9pa1ac
	 /*line :1*/ aeapDF.InitBufs(qMghkaiHXh)
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ syscall.Ky9hDuJe(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF.awWc3wwV[0],  /*line :1*/ uint32( /*line :1*/ len(aeapDF.awWc3wwV)), &aeapDF.sSqPaaxx, 0, &aeapDF.gZWJ1u9D4auz, nil)
	})
	 /*line :1*/ aeapDF.ClearBufs()
	 /*line :1*/ Qah2xxOqw(mnkyfx)
	 /*line :1*/ d3erfj9j(qMghkaiHXh,  /*line :1*/ int64(mnkyfx))
	return  /*line :1*/ int64(mnkyfx), dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) WriteTo(qMghkaiHXh []byte, iAptnzr4Hy4m syscall.S4iroLVT4) (int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	if  /*line :1*/ len(qMghkaiHXh) == 0 {

		aeapDF := &s_ZR_8.aQ9pa1ac
		 /*line :1*/ aeapDF.InitBuf(qMghkaiHXh)
		aeapDF.rVuSBLM = iAptnzr4Hy4m
		mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ syscall.CzWMuENIar(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, 0, aeapDF.rVuSBLM, &aeapDF.gZWJ1u9D4auz, nil)
		})
		return mnkyfx, dtU8tBiUGS4
	}

	iuvUgMayz := 0
	for  /*line :1*/ len(qMghkaiHXh) > 0 {
		mvTjugW := qMghkaiHXh
		if  /*line :1*/ len(mvTjugW) > maxRW {
			mvTjugW = mvTjugW[:maxRW]
		}
		aeapDF := &s_ZR_8.aQ9pa1ac
		 /*line :1*/ aeapDF.InitBuf(mvTjugW)
		aeapDF.rVuSBLM = iAptnzr4Hy4m
		mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ syscall.CzWMuENIar(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, 0, aeapDF.rVuSBLM, &aeapDF.gZWJ1u9D4auz, nil)
		})
		iuvUgMayz +=  /*line :1*/ int(mnkyfx)
		if dtU8tBiUGS4 != nil {
			return iuvUgMayz, dtU8tBiUGS4
		}
		qMghkaiHXh = qMghkaiHXh[mnkyfx:]
	}
	return iuvUgMayz, nil
}


func (s_ZR_8 *ENfmDmMaozH) WriteToInet4(qMghkaiHXh []byte, _89sWOk *syscall.Hx8lqxSkV) (int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	if  /*line :1*/ len(qMghkaiHXh) == 0 {

		aeapDF := &s_ZR_8.aQ9pa1ac
		 /*line :1*/ aeapDF.InitBuf(qMghkaiHXh)
		mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ windows.GLhwGjGTr(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, 0, _89sWOk, &aeapDF.gZWJ1u9D4auz, nil)
		})
		return mnkyfx, dtU8tBiUGS4
	}

	iuvUgMayz := 0
	for  /*line :1*/ len(qMghkaiHXh) > 0 {
		mvTjugW := qMghkaiHXh
		if  /*line :1*/ len(mvTjugW) > maxRW {
			mvTjugW = mvTjugW[:maxRW]
		}
		aeapDF := &s_ZR_8.aQ9pa1ac
		 /*line :1*/ aeapDF.InitBuf(mvTjugW)
		mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ windows.GLhwGjGTr(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, 0, _89sWOk, &aeapDF.gZWJ1u9D4auz, nil)
		})
		iuvUgMayz +=  /*line :1*/ int(mnkyfx)
		if dtU8tBiUGS4 != nil {
			return iuvUgMayz, dtU8tBiUGS4
		}
		qMghkaiHXh = qMghkaiHXh[mnkyfx:]
	}
	return iuvUgMayz, nil
}


func (s_ZR_8 *ENfmDmMaozH) WriteToInet6(qMghkaiHXh []byte, uWZEc5bfl2 *syscall.HKTAy7_BSU2) (int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	if  /*line :1*/ len(qMghkaiHXh) == 0 {

		aeapDF := &s_ZR_8.aQ9pa1ac
		 /*line :1*/ aeapDF.InitBuf(qMghkaiHXh)
		mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ windows.BQSWeDJBro18(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, 0, uWZEc5bfl2, &aeapDF.gZWJ1u9D4auz, nil)
		})
		return mnkyfx, dtU8tBiUGS4
	}

	iuvUgMayz := 0
	for  /*line :1*/ len(qMghkaiHXh) > 0 {
		mvTjugW := qMghkaiHXh
		if  /*line :1*/ len(mvTjugW) > maxRW {
			mvTjugW = mvTjugW[:maxRW]
		}
		aeapDF := &s_ZR_8.aQ9pa1ac
		 /*line :1*/ aeapDF.InitBuf(mvTjugW)
		mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ windows.BQSWeDJBro18(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, 0, uWZEc5bfl2, &aeapDF.gZWJ1u9D4auz, nil)
		})
		iuvUgMayz +=  /*line :1*/ int(mnkyfx)
		if dtU8tBiUGS4 != nil {
			return iuvUgMayz, dtU8tBiUGS4
		}
		qMghkaiHXh = qMghkaiHXh[mnkyfx:]
	}
	return iuvUgMayz, nil
}




func (s_ZR_8 *ENfmDmMaozH) ConnectEx(zlpc9Z syscall.S4iroLVT4) error {
	aeapDF := &s_ZR_8.aQ9pa1ac
	aeapDF.rVuSBLM = zlpc9Z
	_, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ T6VOMZ(aeapDF.ef3m3FV3JP.X8OEsVYSby, aeapDF.rVuSBLM, nil, 0, nil, &aeapDF.gZWJ1u9D4auz)
	})
	return dtU8tBiUGS4
}

func (s_ZR_8 *ENfmDmMaozH) fjdAw91v(iQK6lNeXaI syscall.SRlvVjrYa, f0TQ94Q5Q []syscall.IXy5oynSaLQM, aeapDF *lVE75f25P) (string, error) {

	aeapDF.azj1aOHqks = iQK6lNeXaI
	aeapDF.jQwc0Q60T =  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(f0TQ94Q5Q[0]))
	_, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ Udl5FaeEpD(aeapDF.ef3m3FV3JP.X8OEsVYSby, aeapDF.azj1aOHqks, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&f0TQ94Q5Q[0])), 0,  /*line :1*/ uint32(aeapDF.jQwc0Q60T),  /*line :1*/ uint32(aeapDF.jQwc0Q60T), &aeapDF.sSqPaaxx, &aeapDF.gZWJ1u9D4auz)
	})
	if dtU8tBiUGS4 != nil {
		 /*line :1*/ KvwyJafTkVwj(iQK6lNeXaI)
		return "acceptex", dtU8tBiUGS4
	}

	dtU8tBiUGS4 =  /*line :1*/ syscall.JwtaZG7GaLGD(iQK6lNeXaI, syscall.SOL_SOCKET, syscall.SO_UPDATE_ACCEPT_CONTEXT, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&s_ZR_8.X8OEsVYSby)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(s_ZR_8.X8OEsVYSby)))
	if dtU8tBiUGS4 != nil {
		 /*line :1*/ KvwyJafTkVwj(iQK6lNeXaI)
		return "setsockopt", dtU8tBiUGS4
	}

	return "", nil
}



func (s_ZR_8 *ENfmDmMaozH) Accept(oW_hUv8TRaHZ func() (syscall.SRlvVjrYa, error)) (syscall.SRlvVjrYa, []syscall.IXy5oynSaLQM, uint32, string, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return syscall.InvalidHandle, nil, 0, "", dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()

	aeapDF := &s_ZR_8.x0rW_mPQcb
	var f0TQ94Q5Q [2]syscall.IXy5oynSaLQM
	for {
		iQK6lNeXaI, dtU8tBiUGS4 :=  /*line :1*/ oW_hUv8TRaHZ()
		if dtU8tBiUGS4 != nil {
			return syscall.InvalidHandle, nil, 0, "", dtU8tBiUGS4
		}

		g3a0c12u51f, dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.fjdAw91v(iQK6lNeXaI, f0TQ94Q5Q[:], aeapDF)
		if dtU8tBiUGS4 == nil {
			return iQK6lNeXaI, f0TQ94Q5Q[:],  /*line :1*/ uint32(aeapDF.jQwc0Q60T), "", nil
		}

		xZ_oGEK, qZcIqN0Ba2CB := dtU8tBiUGS4.(syscall.O9Mga3)
		if !qZcIqN0Ba2CB {
			return syscall.InvalidHandle, nil, 0, g3a0c12u51f, dtU8tBiUGS4
		}
		switch xZ_oGEK {
		case syscall.ERROR_NETNAME_DELETED, syscall.WSAECONNRESET:

		default:
			return syscall.InvalidHandle, nil, 0, g3a0c12u51f, dtU8tBiUGS4
		}
	}
}


func (s_ZR_8 *ENfmDmMaozH) Seek(bxEQLKlt int64, l3vE0vsO int) (int64, error) {
	if s_ZR_8.sTN5PtBj1N == kindPipe {
		return 0, syscall.ESPIPE
	}
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()

	 /*line :1*/ s_ZR_8.ezOne7Au.Lock()
	defer  /*line :1*/ s_ZR_8.ezOne7Au.Unlock()

	return  /*line :1*/ syscall.O4FqtrPqbHRY(s_ZR_8.X8OEsVYSby, bxEQLKlt, l3vE0vsO)
}


func (s_ZR_8 *ENfmDmMaozH) Fchmod(_x0IXPJ uint32) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()

	var wmR0FCrSvE syscall.HCYt4WU_Wmzb
	if dtU8tBiUGS4 :=  /*line :1*/ syscall.D1OiNlWWBa(s_ZR_8.X8OEsVYSby, &wmR0FCrSvE); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	fOhu_vRARXM := wmR0FCrSvE.FrYcZLBau7r
	if _x0IXPJ&syscall.S_IWRITE != 0 {
		fOhu_vRARXM &^= syscall.FILE_ATTRIBUTE_READONLY
	} else {
		fOhu_vRARXM |= syscall.FILE_ATTRIBUTE_READONLY
	}
	if fOhu_vRARXM == wmR0FCrSvE.FrYcZLBau7r {
		return nil
	}

	var giIaOfQk windows.ORAO0wrXtu
	giIaOfQk.FuIvuad = fOhu_vRARXM
	gwLmzX :=  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(wmR0FCrSvE))
	return  /*line :1*/ windows.ApJRwykg1p(s_ZR_8.X8OEsVYSby, windows.FileBasicInfo,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&giIaOfQk)), gwLmzX)
}


func (s_ZR_8 *ENfmDmMaozH) Fchdir() error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.LCrVI4(s_ZR_8.X8OEsVYSby)
}


func (s_ZR_8 *ENfmDmMaozH) GetFileType() (uint32, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.VnXJph2(s_ZR_8.X8OEsVYSby)
}


func (s_ZR_8 *ENfmDmMaozH) GetFileInformationByHandle(rgg02EpWaS *syscall.HCYt4WU_Wmzb) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.D1OiNlWWBa(s_ZR_8.X8OEsVYSby, rgg02EpWaS)
}


func (s_ZR_8 *ENfmDmMaozH) RawRead(orn_5D2Pdk func(uintptr) bool) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()
	for {
		if  /*line :1*/ orn_5D2Pdk( /*line :1*/ uintptr(s_ZR_8.X8OEsVYSby)) {
			return nil
		}

		aeapDF := &s_ZR_8.x0rW_mPQcb
		 /*line :1*/ aeapDF.InitBuf(nil)
		if !s_ZR_8.MDfFexom {
			aeapDF.oXBY5N |= windows.MSG_PEEK
		}
		_, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ syscall.J67wPAoai1Kb(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF._vTEq5NaQ, 1, &aeapDF.sSqPaaxx, &aeapDF.oXBY5N, &aeapDF.gZWJ1u9D4auz, nil)
		})
		if dtU8tBiUGS4 == windows.WSAEMSGSIZE {

		} else if dtU8tBiUGS4 != nil {
			return dtU8tBiUGS4
		}
	}
}


func (s_ZR_8 *ENfmDmMaozH) RawWrite(orn_5D2Pdk func(uintptr) bool) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	if  /*line :1*/ orn_5D2Pdk( /*line :1*/ uintptr(s_ZR_8.X8OEsVYSby)) {
		return nil
	}

	return syscall.EWINDOWS
}

func taUhjWDE(vRMcaJ *syscall.IXy5oynSaLQM, iAptnzr4Hy4m *syscall.Hx8lqxSkV) int32 {
	*vRMcaJ = syscall.IXy5oynSaLQM{}
	cIrI2Oif := (* /*line :1*/ syscall.NXUDVMl2)( /*line :1*/ unsafe.Pointer(vRMcaJ))
	cIrI2Oif.JgQng5oWY8Ul = syscall.AF_INET
	yZmyVky9Vg := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&cIrI2Oif.VqKTvvibG))
	yZmyVky9Vg[0] =  /*line :1*/ byte(iAptnzr4Hy4m.AMBsbT8war >> 8)
	yZmyVky9Vg[1] =  /*line :1*/ byte(iAptnzr4Hy4m.AMBsbT8war)
	cIrI2Oif.AGDgR2QSLqW = iAptnzr4Hy4m.Q3zz2uH
	return  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*cIrI2Oif))
}

func vq_4GZ(vRMcaJ *syscall.IXy5oynSaLQM, iAptnzr4Hy4m *syscall.HKTAy7_BSU2) int32 {
	*vRMcaJ = syscall.IXy5oynSaLQM{}
	cIrI2Oif := (* /*line :1*/ syscall.H7tgUmcLUW)( /*line :1*/ unsafe.Pointer(vRMcaJ))
	cIrI2Oif.XYs3jaLEe = syscall.AF_INET6
	yZmyVky9Vg := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&cIrI2Oif.D_H8WxCL7))
	yZmyVky9Vg[0] =  /*line :1*/ byte(iAptnzr4Hy4m.NtLA3k >> 8)
	yZmyVky9Vg[1] =  /*line :1*/ byte(iAptnzr4Hy4m.NtLA3k)
	cIrI2Oif.Jl_oKnHTc = iAptnzr4Hy4m.AsutdJq2
	cIrI2Oif.G1aNyX7 = iAptnzr4Hy4m.Y4XxFr
	return  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*cIrI2Oif))
}

func dcDIm2u(vRMcaJ *syscall.IXy5oynSaLQM, iAptnzr4Hy4m *syscall.Hx8lqxSkV) {
	nimorQuFVzHa := (* /*line :1*/ syscall.NXUDVMl2)( /*line :1*/ unsafe.Pointer(vRMcaJ))
	yZmyVky9Vg := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&nimorQuFVzHa.VqKTvvibG))
	iAptnzr4Hy4m.AMBsbT8war =  /*line :1*/ int(yZmyVky9Vg[0])<<8 +  /*line :1*/ int(yZmyVky9Vg[1])
	iAptnzr4Hy4m.Q3zz2uH = nimorQuFVzHa.AGDgR2QSLqW
}

func hrB5fkV5Ww(vRMcaJ *syscall.IXy5oynSaLQM, iAptnzr4Hy4m *syscall.HKTAy7_BSU2) {
	nimorQuFVzHa := (* /*line :1*/ syscall.H7tgUmcLUW)( /*line :1*/ unsafe.Pointer(vRMcaJ))
	yZmyVky9Vg := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&nimorQuFVzHa.D_H8WxCL7))
	iAptnzr4Hy4m.NtLA3k =  /*line :1*/ int(yZmyVky9Vg[0])<<8 +  /*line :1*/ int(yZmyVky9Vg[1])
	iAptnzr4Hy4m.AsutdJq2 = nimorQuFVzHa.Jl_oKnHTc
	iAptnzr4Hy4m.Y4XxFr = nimorQuFVzHa.G1aNyX7
}

func kzTWkT5(vRMcaJ *syscall.IXy5oynSaLQM, iAptnzr4Hy4m syscall.S4iroLVT4) (int32, error) {
	switch iAptnzr4Hy4m := iAptnzr4Hy4m.(type) {
	case *syscall.Hx8lqxSkV:
		f2zf400_mtPp :=  /*line :1*/ taUhjWDE(vRMcaJ, iAptnzr4Hy4m)
		return f2zf400_mtPp, nil
	case *syscall.HKTAy7_BSU2:
		f2zf400_mtPp :=  /*line :1*/ vq_4GZ(vRMcaJ, iAptnzr4Hy4m)
		return f2zf400_mtPp, nil
	default:
		return 0, syscall.EWINDOWS
	}
}


func (s_ZR_8 *ENfmDmMaozH) ReadMsg(yZmyVky9Vg []byte, rcxhh3yIRuRD []byte, a5fb2Y int) (int, int, int, syscall.S4iroLVT4, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return 0, 0, 0, nil, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()

	if  /*line :1*/ len(yZmyVky9Vg) > maxRW {
		yZmyVky9Vg = yZmyVky9Vg[:maxRW]
	}

	aeapDF := &s_ZR_8.x0rW_mPQcb
	 /*line :1*/ aeapDF.InitMsg(yZmyVky9Vg, rcxhh3yIRuRD)
	if aeapDF.rKvfha6sO == nil {
		aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
	}
	aeapDF.yW8dpq5f6K.ClAqGya = ( /*line :1*/ syscall.VjfACnL)( /*line :1*/ unsafe.Pointer(aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.MeKqL5 =  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.WDy42Vy =  /*line :1*/ uint32(a5fb2Y)
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ windows.GhYfJv(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF.yW8dpq5f6K, &aeapDF.sSqPaaxx, &aeapDF.gZWJ1u9D4auz, nil)
	})
	dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.cKC2EMAVUS2(mnkyfx, dtU8tBiUGS4)
	var iAptnzr4Hy4m syscall.S4iroLVT4
	if dtU8tBiUGS4 == nil {
		iAptnzr4Hy4m, dtU8tBiUGS4 =  /*line :1*/ aeapDF.rKvfha6sO.Sockaddr()
	}
	return mnkyfx,  /*line :1*/ int(aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.Ix58rFqPW),  /*line :1*/ int(aeapDF.yW8dpq5f6K.WDy42Vy), iAptnzr4Hy4m, dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) ReadMsgInet4(yZmyVky9Vg []byte, rcxhh3yIRuRD []byte, a5fb2Y int, _89sWOk *syscall.Hx8lqxSkV) (int, int, int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return 0, 0, 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()

	if  /*line :1*/ len(yZmyVky9Vg) > maxRW {
		yZmyVky9Vg = yZmyVky9Vg[:maxRW]
	}

	aeapDF := &s_ZR_8.x0rW_mPQcb
	 /*line :1*/ aeapDF.InitMsg(yZmyVky9Vg, rcxhh3yIRuRD)
	if aeapDF.rKvfha6sO == nil {
		aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
	}
	aeapDF.yW8dpq5f6K.ClAqGya = ( /*line :1*/ syscall.VjfACnL)( /*line :1*/ unsafe.Pointer(aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.MeKqL5 =  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.WDy42Vy =  /*line :1*/ uint32(a5fb2Y)
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ windows.GhYfJv(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF.yW8dpq5f6K, &aeapDF.sSqPaaxx, &aeapDF.gZWJ1u9D4auz, nil)
	})
	dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.cKC2EMAVUS2(mnkyfx, dtU8tBiUGS4)
	if dtU8tBiUGS4 == nil {
		 /*line :1*/ dcDIm2u(aeapDF.rKvfha6sO, _89sWOk)
	}
	return mnkyfx,  /*line :1*/ int(aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.Ix58rFqPW),  /*line :1*/ int(aeapDF.yW8dpq5f6K.WDy42Vy), dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) ReadMsgInet6(yZmyVky9Vg []byte, rcxhh3yIRuRD []byte, a5fb2Y int, uWZEc5bfl2 *syscall.HKTAy7_BSU2) (int, int, int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.kMJTtXNo(); dtU8tBiUGS4 != nil {
		return 0, 0, 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.ldjlnUNEdCg6()

	if  /*line :1*/ len(yZmyVky9Vg) > maxRW {
		yZmyVky9Vg = yZmyVky9Vg[:maxRW]
	}

	aeapDF := &s_ZR_8.x0rW_mPQcb
	 /*line :1*/ aeapDF.InitMsg(yZmyVky9Vg, rcxhh3yIRuRD)
	if aeapDF.rKvfha6sO == nil {
		aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
	}
	aeapDF.yW8dpq5f6K.ClAqGya = ( /*line :1*/ syscall.VjfACnL)( /*line :1*/ unsafe.Pointer(aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.MeKqL5 =  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.WDy42Vy =  /*line :1*/ uint32(a5fb2Y)
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ windows.GhYfJv(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF.yW8dpq5f6K, &aeapDF.sSqPaaxx, &aeapDF.gZWJ1u9D4auz, nil)
	})
	dtU8tBiUGS4 =  /*line :1*/ s_ZR_8.cKC2EMAVUS2(mnkyfx, dtU8tBiUGS4)
	if dtU8tBiUGS4 == nil {
		 /*line :1*/ hrB5fkV5Ww(aeapDF.rKvfha6sO, uWZEc5bfl2)
	}
	return mnkyfx,  /*line :1*/ int(aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.Ix58rFqPW),  /*line :1*/ int(aeapDF.yW8dpq5f6K.WDy42Vy), dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) WriteMsg(yZmyVky9Vg []byte, rcxhh3yIRuRD []byte, iAptnzr4Hy4m syscall.S4iroLVT4) (int, int, error) {
	if  /*line :1*/ len(yZmyVky9Vg) > maxRW {
		return 0, 0,  /*line :1*/ errors.Su6g6hRoi9X("packet is too large (only 1GB is allowed)")
	}

	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	aeapDF := &s_ZR_8.aQ9pa1ac
	 /*line :1*/ aeapDF.InitMsg(yZmyVky9Vg, rcxhh3yIRuRD)
	if iAptnzr4Hy4m != nil {
		if aeapDF.rKvfha6sO == nil {
			aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
		}
		ajpah49Ud, dtU8tBiUGS4 :=  /*line :1*/ kzTWkT5(aeapDF.rKvfha6sO, iAptnzr4Hy4m)
		if dtU8tBiUGS4 != nil {
			return 0, 0, dtU8tBiUGS4
		}
		aeapDF.yW8dpq5f6K.ClAqGya = ( /*line :1*/ syscall.VjfACnL)( /*line :1*/ unsafe.Pointer(aeapDF.rKvfha6sO))
		aeapDF.yW8dpq5f6K.MeKqL5 = ajpah49Ud
	}
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ windows.UsGtgcD5(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF.yW8dpq5f6K, 0, &aeapDF.sSqPaaxx, &aeapDF.gZWJ1u9D4auz, nil)
	})
	return mnkyfx,  /*line :1*/ int(aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.Ix58rFqPW), dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) WriteMsgInet4(yZmyVky9Vg []byte, rcxhh3yIRuRD []byte, iAptnzr4Hy4m *syscall.Hx8lqxSkV) (int, int, error) {
	if  /*line :1*/ len(yZmyVky9Vg) > maxRW {
		return 0, 0,  /*line :1*/ errors.Su6g6hRoi9X("packet is too large (only 1GB is allowed)")
	}

	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	aeapDF := &s_ZR_8.aQ9pa1ac
	 /*line :1*/ aeapDF.InitMsg(yZmyVky9Vg, rcxhh3yIRuRD)
	if aeapDF.rKvfha6sO == nil {
		aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
	}
	ajpah49Ud :=  /*line :1*/ taUhjWDE(aeapDF.rKvfha6sO, iAptnzr4Hy4m)
	aeapDF.yW8dpq5f6K.ClAqGya = ( /*line :1*/ syscall.VjfACnL)( /*line :1*/ unsafe.Pointer(aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.MeKqL5 = ajpah49Ud
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ windows.UsGtgcD5(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF.yW8dpq5f6K, 0, &aeapDF.sSqPaaxx, &aeapDF.gZWJ1u9D4auz, nil)
	})
	return mnkyfx,  /*line :1*/ int(aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.Ix58rFqPW), dtU8tBiUGS4
}


func (s_ZR_8 *ENfmDmMaozH) WriteMsgInet6(yZmyVky9Vg []byte, rcxhh3yIRuRD []byte, iAptnzr4Hy4m *syscall.HKTAy7_BSU2) (int, int, error) {
	if  /*line :1*/ len(yZmyVky9Vg) > maxRW {
		return 0, 0,  /*line :1*/ errors.Su6g6hRoi9X("packet is too large (only 1GB is allowed)")
	}

	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	aeapDF := &s_ZR_8.aQ9pa1ac
	 /*line :1*/ aeapDF.InitMsg(yZmyVky9Vg, rcxhh3yIRuRD)
	if aeapDF.rKvfha6sO == nil {
		aeapDF.rKvfha6sO =  /*line :1*/ new(syscall.IXy5oynSaLQM)
	}
	ajpah49Ud :=  /*line :1*/ vq_4GZ(aeapDF.rKvfha6sO, iAptnzr4Hy4m)
	aeapDF.yW8dpq5f6K.ClAqGya = ( /*line :1*/ syscall.VjfACnL)( /*line :1*/ unsafe.Pointer(aeapDF.rKvfha6sO))
	aeapDF.yW8dpq5f6K.MeKqL5 = ajpah49Ud
	mnkyfx, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
		return  /*line :1*/ windows.UsGtgcD5(aeapDF.ef3m3FV3JP.X8OEsVYSby, &aeapDF.yW8dpq5f6K, 0, &aeapDF.sSqPaaxx, &aeapDF.gZWJ1u9D4auz, nil)
	})
	return mnkyfx,  /*line :1*/ int(aeapDF.yW8dpq5f6K.RjCEgxXvL4QC.Ix58rFqPW), dtU8tBiUGS4
}
