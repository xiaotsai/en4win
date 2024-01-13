//line :1
package gF9mZs2

import (
	windows "LdptURlN"
	os "hPMrClpC"
	syscall "bUKeamF"
	"unsafe"
)





func aQhDgDOvXgC() ([]*windows.WXavkAObzxI, error) {
	var fIadEXIim6l []byte
	v3upVm :=  /*line :1*/ uint32(15000)
	for {
		fIadEXIim6l =  /*line :1*/ make([]byte, v3upVm)
		h_ljk48Bm :=  /*line :1*/ windows.ViNAxxzfGBB(syscall.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0, (* /*line :1*/ windows.WXavkAObzxI)( /*line :1*/ unsafe.Pointer(&fIadEXIim6l[0])), &v3upVm)
		if h_ljk48Bm == nil {
			if v3upVm == 0 {
				return nil, nil
			}
			break
		}
		if h_ljk48Bm.(syscall.O9Mga3) != syscall.ERROR_BUFFER_OVERFLOW {
			return nil,  /*line :1*/ os.BTaHHl("getadaptersaddresses", h_ljk48Bm)
		}
		if v3upVm <=  /*line :1*/ uint32( /*line :1*/ len(fIadEXIim6l)) {
			return nil,  /*line :1*/ os.BTaHHl("getadaptersaddresses", h_ljk48Bm)
		}
	}
	var upqMp71No []*windows.WXavkAObzxI
	for oJ5XGLy := (* /*line :1*/ windows.WXavkAObzxI)( /*line :1*/ unsafe.Pointer(&fIadEXIim6l[0])); oJ5XGLy != nil; oJ5XGLy = oJ5XGLy.PS1xLQA {
		upqMp71No =  /*line :1*/ append(upqMp71No, oJ5XGLy)
	}
	return upqMp71No, nil
}




func s0JmDXAcmu(bEl_nrd int) ([]U_Uc9la, error) {
	upqMp71No, h_ljk48Bm :=  /*line :1*/ aQhDgDOvXgC()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	var uotBHWl8c47 []U_Uc9la
	for _, oJ5XGLy := range upqMp71No {
		eJTgFgJkWaQ := oJ5XGLy.ZvNqpW
		if eJTgFgJkWaQ == 0 {
			eJTgFgJkWaQ = oJ5XGLy.KLtlYXkLZOhm
		}
		if bEl_nrd == 0 || bEl_nrd ==  /*line :1*/ int(eJTgFgJkWaQ) {
			dvFmDYmQz := U_Uc9la{
				OJz1cH7Vb:	 /*line :1*/ int(eJTgFgJkWaQ),
				IykpW1qlQjWm:	 /*line :1*/ windows.DOR2gxA_7npQ(oJ5XGLy.MndZKwmw0g_C),
			}
			if oJ5XGLy.AOg4mu == windows.IfOperStatusUp {
				dvFmDYmQz.LFFOqgAI |= FlagUp
				dvFmDYmQz.LFFOqgAI |= FlagRunning
			}

			switch oJ5XGLy.IHFX_Ws {
			case windows.IF_TYPE_ETHERNET_CSMACD, windows.IF_TYPE_ISO88025_TOKENRING, windows.IF_TYPE_IEEE80211, windows.IF_TYPE_IEEE1394:
				dvFmDYmQz.LFFOqgAI |= FlagBroadcast | FlagMulticast
			case windows.IF_TYPE_PPP, windows.IF_TYPE_TUNNEL:
				dvFmDYmQz.LFFOqgAI |= FlagPointToPoint | FlagMulticast
			case windows.IF_TYPE_SOFTWARE_LOOPBACK:
				dvFmDYmQz.LFFOqgAI |= FlagLoopback | FlagMulticast
			case windows.IF_TYPE_ATM:
				dvFmDYmQz.LFFOqgAI |= FlagBroadcast | FlagPointToPoint | FlagMulticast
			}
			if oJ5XGLy.YLKBBLbazwH == 0xffffffff {
				dvFmDYmQz.Vlcyn9WPA2 = -1
			} else {
				dvFmDYmQz.Vlcyn9WPA2 =  /*line :1*/ int(oJ5XGLy.YLKBBLbazwH)
			}
			if oJ5XGLy.DKieIVi0 > 0 {
				dvFmDYmQz.AgZ5lCg =  /*line :1*/ make(CGTWv0eqN, oJ5XGLy.DKieIVi0)
				 /*line :1*/ copy(dvFmDYmQz.AgZ5lCg, oJ5XGLy.XP4Wryx7b2h[:])
			}
			uotBHWl8c47 =  /*line :1*/ append(uotBHWl8c47, dvFmDYmQz)
			if bEl_nrd == dvFmDYmQz.OJz1cH7Vb {
				break
			}
		}
	}
	return uotBHWl8c47, nil
}




func uovWlYao2O(dvFmDYmQz *U_Uc9la) ([]EqbMXsaU1lE, error) {
	upqMp71No, h_ljk48Bm :=  /*line :1*/ aQhDgDOvXgC()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	var zDd8fqQGg []EqbMXsaU1lE
	for _, oJ5XGLy := range upqMp71No {
		eJTgFgJkWaQ := oJ5XGLy.ZvNqpW
		if eJTgFgJkWaQ == 0 {
			eJTgFgJkWaQ = oJ5XGLy.KLtlYXkLZOhm
		}
		if dvFmDYmQz == nil || dvFmDYmQz.OJz1cH7Vb ==  /*line :1*/ int(eJTgFgJkWaQ) {
			for mIrjKzx := oJ5XGLy.N5iUveDkX; mIrjKzx != nil; mIrjKzx = mIrjKzx.QL7SVZyWpMDz {
				epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ mIrjKzx.Zwf7iZEa.NgzCg7T0_8yT.Sockaddr()
				if h_ljk48Bm != nil {
					return nil,  /*line :1*/ os.BTaHHl("sockaddr", h_ljk48Bm)
				}
				switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
				case *syscall.Hx8lqxSkV:
					zDd8fqQGg =  /*line :1*/ append(zDd8fqQGg, &M6TN4rba7{RJ6BaPgcU:  /*line :1*/ Ip1NFCL6(epTaNF0RJ8eN.Q3zz2uH[0], epTaNF0RJ8eN.Q3zz2uH[1], epTaNF0RJ8eN.Q3zz2uH[2], epTaNF0RJ8eN.Q3zz2uH[3]), Yj5CZGxK1oow:  /*line :1*/ RYiA8Y9b( /*line :1*/ int(mIrjKzx.X2zs0IMo), 8*IPv4len)})
				case *syscall.HKTAy7_BSU2:
					bg70ff := &M6TN4rba7{RJ6BaPgcU:  /*line :1*/ make(GNraIvYhB, IPv6len), Yj5CZGxK1oow:  /*line :1*/ RYiA8Y9b( /*line :1*/ int(mIrjKzx.X2zs0IMo), 8*IPv6len)}
					 /*line :1*/ copy(bg70ff.RJ6BaPgcU, epTaNF0RJ8eN.Y4XxFr[:])
					zDd8fqQGg =  /*line :1*/ append(zDd8fqQGg, bg70ff)
				}
			}
			for wfe3sh := oJ5XGLy.BaxbOi; wfe3sh != nil; wfe3sh = wfe3sh.Xp8705 {
				epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ wfe3sh.ZTtP0dAzzo6F.NgzCg7T0_8yT.Sockaddr()
				if h_ljk48Bm != nil {
					return nil,  /*line :1*/ os.BTaHHl("sockaddr", h_ljk48Bm)
				}
				switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
				case *syscall.Hx8lqxSkV:
					zDd8fqQGg =  /*line :1*/ append(zDd8fqQGg, &FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ Ip1NFCL6(epTaNF0RJ8eN.Q3zz2uH[0], epTaNF0RJ8eN.Q3zz2uH[1], epTaNF0RJ8eN.Q3zz2uH[2], epTaNF0RJ8eN.Q3zz2uH[3])})
				case *syscall.HKTAy7_BSU2:
					bg70ff := &FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ make(GNraIvYhB, IPv6len)}
					 /*line :1*/ copy(bg70ff.GdouSNq80bRw, epTaNF0RJ8eN.Y4XxFr[:])
					zDd8fqQGg =  /*line :1*/ append(zDd8fqQGg, bg70ff)
				}
			}
		}
	}
	return zDd8fqQGg, nil
}



func gsdWgpJ3dT(dvFmDYmQz *U_Uc9la) ([]EqbMXsaU1lE, error) {
	upqMp71No, h_ljk48Bm :=  /*line :1*/ aQhDgDOvXgC()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	var zDd8fqQGg []EqbMXsaU1lE
	for _, oJ5XGLy := range upqMp71No {
		eJTgFgJkWaQ := oJ5XGLy.ZvNqpW
		if eJTgFgJkWaQ == 0 {
			eJTgFgJkWaQ = oJ5XGLy.KLtlYXkLZOhm
		}
		if dvFmDYmQz == nil || dvFmDYmQz.OJz1cH7Vb ==  /*line :1*/ int(eJTgFgJkWaQ) {
			for mc6vT_Qz7UMR := oJ5XGLy.BX8u3C; mc6vT_Qz7UMR != nil; mc6vT_Qz7UMR = mc6vT_Qz7UMR.Ltx_FiDtXD {
				epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ mc6vT_Qz7UMR.F07g7TUjQ.NgzCg7T0_8yT.Sockaddr()
				if h_ljk48Bm != nil {
					return nil,  /*line :1*/ os.BTaHHl("sockaddr", h_ljk48Bm)
				}
				switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
				case *syscall.Hx8lqxSkV:
					zDd8fqQGg =  /*line :1*/ append(zDd8fqQGg, &FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ Ip1NFCL6(epTaNF0RJ8eN.Q3zz2uH[0], epTaNF0RJ8eN.Q3zz2uH[1], epTaNF0RJ8eN.Q3zz2uH[2], epTaNF0RJ8eN.Q3zz2uH[3])})
				case *syscall.HKTAy7_BSU2:
					bg70ff := &FZJphYv9hV{GdouSNq80bRw:  /*line :1*/ make(GNraIvYhB, IPv6len)}
					 /*line :1*/ copy(bg70ff.GdouSNq80bRw, epTaNF0RJ8eN.Y4XxFr[:])
					zDd8fqQGg =  /*line :1*/ append(zDd8fqQGg, bg70ff)
				}
			}
		}
	}
	return zDd8fqQGg, nil
}
