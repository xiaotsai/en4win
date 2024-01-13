//line :1
//go:build unix || windows

package gF9mZs2

import (
	context "fOYpb3F03EG"
	poll "MjXXMR"
	os "hPMrClpC"
	syscall "bUKeamF"
)

func _tP9pgLw(yESLyde9LHhT context.MBCyA6, gF9mZs2 string, nqsgnnaf, dr86pU, iaebRs5X3 int, q_vcfH bool, bFNyUpAx, wcA44hhD jIv0wLJwloQ, rhsSAD7C func(context.MBCyA6, string, string, syscall.CVnV1i) error) (vyaiiPXm6k_W *jmJczkl1, h_ljk48Bm error) {
	dRoFccaZ, h_ljk48Bm :=  /*line :1*/ bbSHyj(nqsgnnaf, dr86pU, iaebRs5X3)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	if h_ljk48Bm =  /*line :1*/ faet4q(dRoFccaZ, nqsgnnaf, dr86pU, q_vcfH); h_ljk48Bm != nil {
		 /*line :1*/ poll.KvwyJafTkVwj(dRoFccaZ)
		return nil, h_ljk48Bm
	}
	if vyaiiPXm6k_W, h_ljk48Bm =  /*line :1*/ gmaurSXmXv0u(dRoFccaZ, nqsgnnaf, dr86pU, gF9mZs2); h_ljk48Bm != nil {
		 /*line :1*/ poll.KvwyJafTkVwj(dRoFccaZ)
		return nil, h_ljk48Bm
	}

	if bFNyUpAx != nil && wcA44hhD == nil {
		switch dr86pU {
		case syscall.SOCK_STREAM, syscall.SOCK_SEQPACKET:
			if h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.cb_gZv8(yESLyde9LHhT, bFNyUpAx,  /*line :1*/ ajNsMEWiu(), rhsSAD7C); h_ljk48Bm != nil {
				 /*line :1*/ vyaiiPXm6k_W.Close()
				return nil, h_ljk48Bm
			}
			return vyaiiPXm6k_W, nil
		case syscall.SOCK_DGRAM:
			if h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.wnDBm_znu341(yESLyde9LHhT, bFNyUpAx, rhsSAD7C); h_ljk48Bm != nil {
				 /*line :1*/ vyaiiPXm6k_W.Close()
				return nil, h_ljk48Bm
			}
			return vyaiiPXm6k_W, nil
		}
	}
	if h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.tdqObJ_tTokV(yESLyde9LHhT, bFNyUpAx, wcA44hhD, rhsSAD7C); h_ljk48Bm != nil {
		 /*line :1*/ vyaiiPXm6k_W.Close()
		return nil, h_ljk48Bm
	}
	return vyaiiPXm6k_W, nil
}

func (vyaiiPXm6k_W *jmJczkl1) zmlorv() string {
	switch vyaiiPXm6k_W.fCaRuBA3 {
	case "unix", "unixgram", "unixpacket":
		return vyaiiPXm6k_W.fCaRuBA3
	}
	switch vyaiiPXm6k_W.fCaRuBA3[ /*line :1*/ len(vyaiiPXm6k_W.fCaRuBA3)-1] {
	case '4', '6':
		return vyaiiPXm6k_W.fCaRuBA3
	}
	if vyaiiPXm6k_W.zW6I0G4Xm0k == syscall.AF_INET {
		return vyaiiPXm6k_W.fCaRuBA3 + "4"
	}
	return vyaiiPXm6k_W.fCaRuBA3 + "6"
}

func (vyaiiPXm6k_W *jmJczkl1) xZOQl6heY5() func(syscall.S4iroLVT4) EqbMXsaU1lE {
	switch vyaiiPXm6k_W.zW6I0G4Xm0k {
	case syscall.AF_INET, syscall.AF_INET6:
		switch vyaiiPXm6k_W.zpK4ak_Tpl {
		case syscall.SOCK_STREAM:
			return ya0ZV8JJhjc
		case syscall.SOCK_DGRAM:
			return cYt7M3
		case syscall.SOCK_RAW:
			return llbJwAwVB
		}
	case syscall.AF_UNIX:
		switch vyaiiPXm6k_W.zpK4ak_Tpl {
		case syscall.SOCK_STREAM:
			return fdm4acHSd0vA
		case syscall.SOCK_DGRAM:
			return d1JbegF
		case syscall.SOCK_SEQPACKET:
			return _a2TukK9
		}
	}
	return func(syscall.S4iroLVT4) EqbMXsaU1lE { return nil }
}

func (vyaiiPXm6k_W *jmJczkl1) tdqObJ_tTokV(yESLyde9LHhT context.MBCyA6, bFNyUpAx, wcA44hhD jIv0wLJwloQ, rhsSAD7C func(context.MBCyA6, string, string, syscall.CVnV1i) error) error {
	var hl8w2lFX *cfFPDevTPcYC
	var h_ljk48Bm error
	if rhsSAD7C != nil {
		hl8w2lFX, h_ljk48Bm =  /*line :1*/ zQf8had4tml(vyaiiPXm6k_W)
		if h_ljk48Bm != nil {
			return h_ljk48Bm
		}
		var s0RsoE string
		if wcA44hhD != nil {
			s0RsoE =  /*line :1*/ wcA44hhD.String()
		} else if bFNyUpAx != nil {
			s0RsoE =  /*line :1*/ bFNyUpAx.String()
		}
		if h_ljk48Bm :=  /*line :1*/ rhsSAD7C(yESLyde9LHhT,  /*line :1*/ vyaiiPXm6k_W.zmlorv(), s0RsoE, hl8w2lFX); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
	}

	var agaUaBpv4 syscall.S4iroLVT4
	if bFNyUpAx != nil {
		if agaUaBpv4, h_ljk48Bm =  /*line :1*/ bFNyUpAx.jIv0wLJwloQ(vyaiiPXm6k_W.zW6I0G4Xm0k); h_ljk48Bm != nil {
			return h_ljk48Bm
		} else if agaUaBpv4 != nil {
			if h_ljk48Bm =  /*line :1*/ syscall.U5lDiJBkm0(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby, agaUaBpv4); h_ljk48Bm != nil {
				return  /*line :1*/ os.BTaHHl("bind", h_ljk48Bm)
			}
		}
	}
	var hG_buG syscall.S4iroLVT4
	var bu81iGF05z_A syscall.S4iroLVT4
	if wcA44hhD != nil {
		if hG_buG, h_ljk48Bm =  /*line :1*/ wcA44hhD.jIv0wLJwloQ(vyaiiPXm6k_W.zW6I0G4Xm0k); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
		if bu81iGF05z_A, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.nOGPJIg(yESLyde9LHhT, agaUaBpv4, hG_buG); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
		vyaiiPXm6k_W.kBOwaEVMqK = true
	} else {
		if h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.init(); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
	}

	agaUaBpv4, _ =  /*line :1*/ syscall.DaM3_rD4F(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby)
	if bu81iGF05z_A != nil {
		 /*line :1*/ vyaiiPXm6k_W.xaamQ3HuZ( /*line :1*/ vyaiiPXm6k_W.xZOQl6heY5()(agaUaBpv4),  /*line :1*/ vyaiiPXm6k_W.xZOQl6heY5()(bu81iGF05z_A))
	} else if hG_buG, _ =  /*line :1*/ syscall.BWaUaPS4sKD9(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby); hG_buG != nil {
		 /*line :1*/ vyaiiPXm6k_W.xaamQ3HuZ( /*line :1*/ vyaiiPXm6k_W.xZOQl6heY5()(agaUaBpv4),  /*line :1*/ vyaiiPXm6k_W.xZOQl6heY5()(hG_buG))
	} else {
		 /*line :1*/ vyaiiPXm6k_W.xaamQ3HuZ( /*line :1*/ vyaiiPXm6k_W.xZOQl6heY5()(agaUaBpv4), wcA44hhD)
	}
	return nil
}

func (vyaiiPXm6k_W *jmJczkl1) cb_gZv8(yESLyde9LHhT context.MBCyA6, bFNyUpAx jIv0wLJwloQ, jNvgLHkq6v int, rhsSAD7C func(context.MBCyA6, string, string, syscall.CVnV1i) error) error {
	var h_ljk48Bm error
	if h_ljk48Bm =  /*line :1*/ atNQ_vLoo(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	var agaUaBpv4 syscall.S4iroLVT4
	if agaUaBpv4, h_ljk48Bm =  /*line :1*/ bFNyUpAx.jIv0wLJwloQ(vyaiiPXm6k_W.zW6I0G4Xm0k); h_ljk48Bm != nil {
		return h_ljk48Bm
	}

	if rhsSAD7C != nil {
		hl8w2lFX, h_ljk48Bm :=  /*line :1*/ zQf8had4tml(vyaiiPXm6k_W)
		if h_ljk48Bm != nil {
			return h_ljk48Bm
		}
		if h_ljk48Bm :=  /*line :1*/ rhsSAD7C(yESLyde9LHhT,  /*line :1*/ vyaiiPXm6k_W.zmlorv(),  /*line :1*/ bFNyUpAx.String(), hl8w2lFX); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
	}

	if h_ljk48Bm =  /*line :1*/ syscall.U5lDiJBkm0(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby, agaUaBpv4); h_ljk48Bm != nil {
		return  /*line :1*/ os.BTaHHl("bind", h_ljk48Bm)
	}
	if h_ljk48Bm =  /*line :1*/ veHf5d(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby, jNvgLHkq6v); h_ljk48Bm != nil {
		return  /*line :1*/ os.BTaHHl("listen", h_ljk48Bm)
	}
	if h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.init(); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	agaUaBpv4, _ =  /*line :1*/ syscall.DaM3_rD4F(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby)
	 /*line :1*/ vyaiiPXm6k_W.xaamQ3HuZ( /*line :1*/ vyaiiPXm6k_W.xZOQl6heY5()(agaUaBpv4), nil)
	return nil
}

func (vyaiiPXm6k_W *jmJczkl1) wnDBm_znu341(yESLyde9LHhT context.MBCyA6, bFNyUpAx jIv0wLJwloQ, rhsSAD7C func(context.MBCyA6, string, string, syscall.CVnV1i) error) error {
	switch qxwkC3VxG0 := bFNyUpAx.(type) {
	case *ZaffanpNx4:

		if qxwkC3VxG0.ERvaNiiVmytR != nil &&  /*line :1*/ qxwkC3VxG0.ERvaNiiVmytR.IsMulticast() {
			if h_ljk48Bm :=  /*line :1*/ ojx7Y3NFA(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby); h_ljk48Bm != nil {
				return h_ljk48Bm
			}
			qxwkC3VxG0 := *qxwkC3VxG0
			switch vyaiiPXm6k_W.zW6I0G4Xm0k {
			case syscall.AF_INET:
				qxwkC3VxG0.ERvaNiiVmytR = WHSMKoapR
			case syscall.AF_INET6:
				qxwkC3VxG0.ERvaNiiVmytR = VIAfFsvboRvQ
			}
			bFNyUpAx = &qxwkC3VxG0
		}
	}
	var h_ljk48Bm error
	var agaUaBpv4 syscall.S4iroLVT4
	if agaUaBpv4, h_ljk48Bm =  /*line :1*/ bFNyUpAx.jIv0wLJwloQ(vyaiiPXm6k_W.zW6I0G4Xm0k); h_ljk48Bm != nil {
		return h_ljk48Bm
	}

	if rhsSAD7C != nil {
		hl8w2lFX, h_ljk48Bm :=  /*line :1*/ zQf8had4tml(vyaiiPXm6k_W)
		if h_ljk48Bm != nil {
			return h_ljk48Bm
		}
		if h_ljk48Bm :=  /*line :1*/ rhsSAD7C(yESLyde9LHhT,  /*line :1*/ vyaiiPXm6k_W.zmlorv(),  /*line :1*/ bFNyUpAx.String(), hl8w2lFX); h_ljk48Bm != nil {
			return h_ljk48Bm
		}
	}
	if h_ljk48Bm =  /*line :1*/ syscall.U5lDiJBkm0(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby, agaUaBpv4); h_ljk48Bm != nil {
		return  /*line :1*/ os.BTaHHl("bind", h_ljk48Bm)
	}
	if h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.init(); h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	agaUaBpv4, _ =  /*line :1*/ syscall.DaM3_rD4F(vyaiiPXm6k_W.q0t7ga7tBcV5.X8OEsVYSby)
	 /*line :1*/ vyaiiPXm6k_W.xaamQ3HuZ( /*line :1*/ vyaiiPXm6k_W.xZOQl6heY5()(agaUaBpv4), nil)
	return nil
}
