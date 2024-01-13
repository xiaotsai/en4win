//line :1
package gF9mZs2

import (
	context "fOYpb3F03EG"
	poll "MjXXMR"
	windows "LdptURlN"
	os "hPMrClpC"
	"runtime"
	syscall "bUKeamF"
	"unsafe"
)

const (
	readSyscallName	= "wsarecv"
	readFromSyscallName	= "wsarecvfrom"
	readMsgSyscallName	= "wsarecvmsg"
	writeSyscallName	= "wsasend"
	writeToSyscallName	= "wsasendto"
	writeMsgSyscallName	= "wsasendmsg"
)



func bj9LbLuaKRi(gF9mZs2 string) bool {
	switch gF9mZs2 {
	case "tcp", "tcp4", "tcp6":
		return true
	}

	return false
}

func gmaurSXmXv0u(n6TYexZul syscall.SRlvVjrYa, nqsgnnaf, dr86pU int, gF9mZs2 string) (*jmJczkl1, error) {
	prHUfwjRzJ := &jmJczkl1{
		q0t7ga7tBcV5: poll.ENfmDmMaozH{
			X8OEsVYSby:	n6TYexZul,
			MDfFexom:	dr86pU == syscall.SOCK_STREAM,
			ZslYVRp1qJ:	dr86pU != syscall.SOCK_DGRAM && dr86pU != syscall.SOCK_RAW,
		},
		zW6I0G4Xm0k:	nqsgnnaf,
		zpK4ak_Tpl:	dr86pU,
		fCaRuBA3:	gF9mZs2,
	}
	return prHUfwjRzJ, nil
}

func (vyaiiPXm6k_W *jmJczkl1) init() error {
	a9FeDbwq, h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Init(vyaiiPXm6k_W.fCaRuBA3, true)
	if a9FeDbwq != "" {
		h_ljk48Bm =  /*line :1*/ raib0nns(a9FeDbwq, h_ljk48Bm)
	}
	return h_ljk48Bm
}


func (vyaiiPXm6k_W *jmJczkl1) nOGPJIg(yESLyde9LHhT context.MBCyA6, ayKvoOPa, ioVu8c syscall.S4iroLVT4) (syscall.S4iroLVT4, error) {

	if h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.init(); h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	if cSxcyrrgag, d30HIwxht :=  /*line :1*/ yESLyde9LHhT.Deadline(); d30HIwxht && ! /*line :1*/ cSxcyrrgag.IsZero() {
		 /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetWriteDeadline(cSxcyrrgag)
		defer  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetWriteDeadline(wButnm0q8b)
	}
	if ! /*line :1*/ bj9LbLuaKRi(vyaiiPXm6k_W.fCaRuBA3) {
		h_ljk48Bm :=  /*line :1*/ yB6DR7Gt(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby, ioVu8c)
		return nil,  /*line :1*/ os.BTaHHl("connect", h_ljk48Bm)
	}

	if ayKvoOPa == nil {
		switch ioVu8c.(type) {
		case *syscall.Hx8lqxSkV:
			ayKvoOPa = &syscall.Hx8lqxSkV{}
		case *syscall.HKTAy7_BSU2:
			ayKvoOPa = &syscall.HKTAy7_BSU2{}
		default:
			 /*line :1*/ panic("unexpected type in connect")
		}
		if h_ljk48Bm :=  /*line :1*/ syscall.U5lDiJBkm0(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby, ayKvoOPa); h_ljk48Bm != nil {
			return nil,  /*line :1*/ os.BTaHHl("bind", h_ljk48Bm)
		}
	}

	var g6oDBK bool
	switch ioVu8c := ioVu8c.(type) {
	case *syscall.Hx8lqxSkV:
		g6oDBK = ioVu8c.Q3zz2uH[0] == 127
	case *syscall.HKTAy7_BSU2:
		g6oDBK = ioVu8c.Y4XxFr == [16] /*line :1*/ byte(ICyGsSLLvx)
	default:
		 /*line :1*/ panic("unexpected type in connect")
	}
	if g6oDBK {

		pzSLqU := windows.CSH1eB5vRVAy{
			LeR_En3xrj:	windows.TCP_INITIAL_RTO_UNSPECIFIED_RTT,
			H3T1z7dO:	1,
		}
		if  /*line :1*/ windows.N7_CEIm1r() {

			pzSLqU.H3T1z7dO = windows.TCP_INITIAL_RTO_NO_SYN_RETRANSMISSIONS
		}
		var hMkKN3mr uint32

		_ =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WSAIoctl(windows.SIO_TCP_INITIAL_RTO, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&pzSLqU)),  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(pzSLqU)), nil, 0, &hMkKN3mr, nil, 0)
	}

	wdWRIxX1IoF :=  /*line :1*/ make(chan bool)
	defer func() {  /*line :1*/ wdWRIxX1IoF <- true }()
	go func() {
		select {
		case <- /*line :1*/ yESLyde9LHhT.Done():

			 /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetWriteDeadline(eqHMQVuidt)
			<-wdWRIxX1IoF
		case <-wdWRIxX1IoF:
		}
	}()

	if h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.ConnectEx(ioVu8c); h_ljk48Bm != nil {
		select {
		case <- /*line :1*/ yESLyde9LHhT.Done():
			return nil,  /*line :1*/ gCG7rudfaf( /*line :1*/ yESLyde9LHhT.Err())
		default:
			if _, d30HIwxht := h_ljk48Bm.(syscall.O9Mga3); d30HIwxht {
				h_ljk48Bm =  /*line :1*/ os.BTaHHl("connectex", h_ljk48Bm)
			}
			return nil, h_ljk48Bm
		}
	}

	return nil,  /*line :1*/ os.BTaHHl("setsockopt",  /*line :1*/ syscall.JwtaZG7GaLGD(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby, syscall.SOL_SOCKET, syscall.SO_UPDATE_CONNECT_CONTEXT, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby))))
}

func (hl8w2lFX *xacP7km5pe) tT7qFj3HZ(ljsCSk *CqKFIxFUB0B) (int64, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.tT7qFj3HZ(ljsCSk)
	if h_ljk48Bm != nil {
		return doauF8Y, &ZOYGdck{SMNyZk_q: "wsasend", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, nil
}

func (vyaiiPXm6k_W *jmJczkl1) tT7qFj3HZ(zynKBqWHCa *CqKFIxFUB0B) (int64, error) {
	doauF8Y, h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Writev((*[][] /*line :1*/ byte)(zynKBqWHCa))
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y,  /*line :1*/ raib0nns("wsasend", h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) zjXrFvdT() (*jmJczkl1, error) {
	dRoFccaZ, b_C5Ey, jdB9Ifx1v5n, a9FeDbwq, h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Accept(func() (syscall.SRlvVjrYa, error) {
		return  /*line :1*/ bbSHyj(vyaiiPXm6k_W.zW6I0G4Xm0k, vyaiiPXm6k_W.zpK4ak_Tpl, 0)
	})

	if h_ljk48Bm != nil {
		if a9FeDbwq != "" {
			h_ljk48Bm =  /*line :1*/ raib0nns(a9FeDbwq, h_ljk48Bm)
		}
		return nil, h_ljk48Bm
	}

	caG0aoPuTH9, h_ljk48Bm :=  /*line :1*/ gmaurSXmXv0u(dRoFccaZ, vyaiiPXm6k_W.zW6I0G4Xm0k, vyaiiPXm6k_W.zpK4ak_Tpl, vyaiiPXm6k_W.fCaRuBA3)
	if h_ljk48Bm != nil {
		 /*line :1*/ poll.KvwyJafTkVwj(dRoFccaZ)
		return nil, h_ljk48Bm
	}
	if h_ljk48Bm :=  /*line :1*/ caG0aoPuTH9.init(); h_ljk48Bm != nil {
		 /*line :1*/ vyaiiPXm6k_W.Close()
		return nil, h_ljk48Bm
	}

	
	var eTy7LNRDP, dHBrHDwaUj *syscall.IXy5oynSaLQM
	var o7NR3z1apN3, fLzvmD int32
	 /*line :1*/ syscall.NRaqTZ((* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&b_C5Ey[0])),
		0, jdB9Ifx1v5n, jdB9Ifx1v5n, &eTy7LNRDP, &o7NR3z1apN3, &dHBrHDwaUj, &fLzvmD)
	agaUaBpv4, _ :=  /*line :1*/ eTy7LNRDP.Sockaddr()
	hG_buG, _ :=  /*line :1*/ dHBrHDwaUj.Sockaddr()

	 /*line :1*/ caG0aoPuTH9.xaamQ3HuZ( /*line :1*/ caG0aoPuTH9.xZOQl6heY5()(agaUaBpv4),  /*line :1*/ caG0aoPuTH9.xZOQl6heY5()(hG_buG))
	return caG0aoPuTH9, nil
}

func (vyaiiPXm6k_W *jmJczkl1) uNYdhMYF() (*os.BvGocYcXx, error) {

	return nil, syscall.EWINDOWS
}
