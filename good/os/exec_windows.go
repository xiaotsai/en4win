//line :1
package hPMrClpC

import (
	errors "iAHoxjmM"
	windows "LdptURlN"
	"runtime"
	atomic "sync/atomic"
	syscall "bUKeamF"
	time "fRAfQd_"
)

func (baV7tO7i5 *P2XtkGf3) xbyZguBq() (p1ZeM2HShDOx *LcYXxKH, ugqb2IW error) {
	iVmSWxZA4 :=  /*line :1*/ atomic.LoadUintptr(&baV7tO7i5.oNLhPyDvzm9)
	gkORch7na, w50uhPri :=  /*line :1*/ syscall.N_8burTVa( /*line :1*/ syscall.SRlvVjrYa(iVmSWxZA4), syscall.INFINITE)
	switch gkORch7na {
	case syscall.WAIT_OBJECT_0:
		break
	case syscall.WAIT_FAILED:
		return nil,  /*line :1*/ BTaHHl("WaitForSingleObject", w50uhPri)
	default:
		return nil,  /*line :1*/ errors.Su6g6hRoi9X("os: unexpected result from WaitForSingleObject")
	}
	var m8Itsp2hq uint32
	w50uhPri =  /*line :1*/ syscall.YYxQKvX( /*line :1*/ syscall.SRlvVjrYa(iVmSWxZA4), &m8Itsp2hq)
	if w50uhPri != nil {
		return nil,  /*line :1*/ BTaHHl("GetExitCodeProcess", w50uhPri)
	}
	var vN6gjL7zFAe syscall.Pd6f6PxVa
	w50uhPri =  /*line :1*/ syscall.AznBquFp( /*line :1*/ syscall.SRlvVjrYa(iVmSWxZA4), &vN6gjL7zFAe.DInKorDQLw1, &vN6gjL7zFAe.TPYcuXne_w, &vN6gjL7zFAe.UiqiQqR, &vN6gjL7zFAe.JJK8Dbfduef)
	if w50uhPri != nil {
		return nil,  /*line :1*/ BTaHHl("GetProcessTimes", w50uhPri)
	}
	 /*line :1*/ baV7tO7i5.qpSVI0bOJG()

	defer  /*line :1*/ time.H4EBRkSgB(5 * time.Millisecond)
	defer  /*line :1*/ baV7tO7i5.Release()
	return &LcYXxKH{baV7tO7i5.EQoGDb, syscall.HV8Mf1{ZBjv37: m8Itsp2hq}, &vN6gjL7zFAe}, nil
}

func (baV7tO7i5 *P2XtkGf3) mrvahvDrXb8(twfNTarXO J0zkNT) error {
	iVmSWxZA4 :=  /*line :1*/ atomic.LoadUintptr(&baV7tO7i5.oNLhPyDvzm9)
	if iVmSWxZA4 ==  /*line :1*/ uintptr(syscall.InvalidHandle) {
		return syscall.EINVAL
	}
	if  /*line :1*/ baV7tO7i5.abgQEYvfC9rq() {
		return A4466R3
	}
	if twfNTarXO == RebqCuoID {
		var fAm4JU5MM syscall.SRlvVjrYa
		w50uhPri :=  /*line :1*/ syscall.VXs68_ZF(^ /*line :1*/ syscall.SRlvVjrYa(0),  /*line :1*/ syscall.SRlvVjrYa(iVmSWxZA4), ^ /*line :1*/ syscall.SRlvVjrYa(0), &fAm4JU5MM, syscall.PROCESS_TERMINATE, false, 0)
		if w50uhPri != nil {
			return  /*line :1*/ BTaHHl("DuplicateHandle", w50uhPri)
		}
		 /*line :1*/ runtime.KeepAlive(baV7tO7i5)
		defer  /*line :1*/ syscall.FhKJLgXjwG(fAm4JU5MM)
		w50uhPri =  /*line :1*/ syscall.Vp9MpT5Y( /*line :1*/ syscall.SRlvVjrYa(fAm4JU5MM), 1)
		return  /*line :1*/ BTaHHl("TerminateProcess", w50uhPri)
	}

	return  /*line :1*/ syscall.O9Mga3(syscall.EWINDOWS)
}

func (baV7tO7i5 *P2XtkGf3) qq2HndsNY() error {
	iVmSWxZA4 :=  /*line :1*/ atomic.SwapUintptr(&baV7tO7i5.oNLhPyDvzm9,  /*line :1*/ uintptr(syscall.InvalidHandle))
	if iVmSWxZA4 ==  /*line :1*/ uintptr(syscall.InvalidHandle) {
		return syscall.EINVAL
	}
	w50uhPri :=  /*line :1*/ syscall.FhKJLgXjwG( /*line :1*/ syscall.SRlvVjrYa(iVmSWxZA4))
	if w50uhPri != nil {
		return  /*line :1*/ BTaHHl("CloseHandle", w50uhPri)
	}

	 /*line :1*/ runtime.SetFinalizer(baV7tO7i5, nil)
	return nil
}

func ya248LtAOt(gAIGP2p int) (baV7tO7i5 *P2XtkGf3, ugqb2IW error) {
	const da = syscall.STANDARD_RIGHTS_READ |
		syscall.PROCESS_QUERY_INFORMATION | syscall.SYNCHRONIZE
	dB59YpfJ, w50uhPri :=  /*line :1*/ syscall.JkYb3Icqg(da, false,  /*line :1*/ uint32(gAIGP2p))
	if w50uhPri != nil {
		return nil,  /*line :1*/ BTaHHl("OpenProcess", w50uhPri)
	}
	return  /*line :1*/ zvDudDSfa(gAIGP2p,  /*line :1*/ uintptr(dB59YpfJ)), nil
}

func init() {
	eSsokQa :=  /*line :1*/ windows.DOR2gxA_7npQ( /*line :1*/ syscall.IaCOl3xN())
	if  /*line :1*/ len(eSsokQa) == 0 {
		zCgNNL3Y, _ :=  /*line :1*/ X7ThLxab()
		USSgbG96n = []string{zCgNNL3Y}
	} else {
		USSgbG96n =  /*line :1*/ rOZMdrSzey0v(eSsokQa)
	}
}


func b3VONUKvqZ(nR2aPlreQFZA []byte, zHDjCTHI int) []byte {
	for ; zHDjCTHI > 0; zHDjCTHI-- {
		nR2aPlreQFZA =  /*line :1*/ append(nR2aPlreQFZA, '\\')
	}
	return nR2aPlreQFZA
}



func pMo3yTJ(eSsokQa string) (chSRwpAFkaUI []byte, gwmDr8O1q string) {
	var nR2aPlreQFZA []byte
	var zJreFHC bool
	var jAZ4aR6 int
	for ;  /*line :1*/ len(eSsokQa) > 0; eSsokQa = eSsokQa[1:] {
		ivfnLGc := eSsokQa[0]
		switch ivfnLGc {
		case ' ', '\t':
			if !zJreFHC {
				return  /*line :1*/ b3VONUKvqZ(nR2aPlreQFZA, jAZ4aR6), eSsokQa[1:]
			}
		case '"':
			nR2aPlreQFZA =  /*line :1*/ b3VONUKvqZ(nR2aPlreQFZA, jAZ4aR6/2)
			if jAZ4aR6%2 == 0 {

				if zJreFHC &&  /*line :1*/ len(eSsokQa) > 1 && eSsokQa[1] == '"' {
					nR2aPlreQFZA =  /*line :1*/ append(nR2aPlreQFZA, ivfnLGc)
					eSsokQa = eSsokQa[1:]
				}
				zJreFHC = !zJreFHC
			} else {
				nR2aPlreQFZA =  /*line :1*/ append(nR2aPlreQFZA, ivfnLGc)
			}
			jAZ4aR6 = 0
			continue
		case '\\':
			jAZ4aR6++
			continue
		}
		nR2aPlreQFZA =  /*line :1*/ b3VONUKvqZ(nR2aPlreQFZA, jAZ4aR6)
		jAZ4aR6 = 0
		nR2aPlreQFZA =  /*line :1*/ append(nR2aPlreQFZA, ivfnLGc)
	}
	return  /*line :1*/ b3VONUKvqZ(nR2aPlreQFZA, jAZ4aR6), ""
}




func rOZMdrSzey0v(eSsokQa string) []string {
	var wVx1awjoBH []string
	for  /*line :1*/ len(eSsokQa) > 0 {
		if eSsokQa[0] == ' ' || eSsokQa[0] == '\t' {
			eSsokQa = eSsokQa[1:]
			continue
		}
		var chSRwpAFkaUI []byte
		chSRwpAFkaUI, eSsokQa =  /*line :1*/ pMo3yTJ(eSsokQa)
		wVx1awjoBH =  /*line :1*/ append(wVx1awjoBH,  /*line :1*/ string(chSRwpAFkaUI))
	}
	return wVx1awjoBH
}

func ajnov9HQCK(ehGW_z *syscall.T8WbUqAC3v) time.GKMwTGxQa0S {
	zHDjCTHI :=  /*line :1*/ int64(ehGW_z.L9KWrjMHS)<<32 +  /*line :1*/ int64(ehGW_z.Cl_x1Rd)
	return  /*line :1*/ time.GKMwTGxQa0S(zHDjCTHI*100) * time.Nanosecond
}

func (baV7tO7i5 *LcYXxKH) ivGDNoOUnU() time.GKMwTGxQa0S {
	return  /*line :1*/ ajnov9HQCK(&baV7tO7i5.t9b_DGjwz.JJK8Dbfduef)
}

func (baV7tO7i5 *LcYXxKH) baMIJzSnuXFy() time.GKMwTGxQa0S {
	return  /*line :1*/ ajnov9HQCK(&baV7tO7i5.t9b_DGjwz.UiqiQqR)
}
