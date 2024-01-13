//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package gF9mZs2

import (
	context "fOYpb3F03EG"
	poll "MjXXMR"
	netip "iPdCam_KQOXj"
	"runtime"
	syscall "bUKeamF"
)

func (fMPVz2iGiyq *uPtFhw) rAKCx8() {
	dRoFccaZ, h_ljk48Bm :=  /*line :1*/ bbSHyj(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	switch h_ljk48Bm {
	case syscall.EAFNOSUPPORT, syscall.EPROTONOSUPPORT:
	case nil:
		 /*line :1*/ poll.KvwyJafTkVwj(dRoFccaZ)
		fMPVz2iGiyq.dBENHw = true
	}
	var e_JzQ29Jvp = []struct {
		g28MY22awQ	HiOzqOVjN1P
		tXmeiIWc	int
	}{

		{g28MY22awQ: HiOzqOVjN1P{ERvaNiiVmytR:  /*line :1*/ XtVFDxFx("::1")}, tXmeiIWc: 1},

		{g28MY22awQ: HiOzqOVjN1P{ERvaNiiVmytR:  /*line :1*/ Ip1NFCL6(127, 0, 0, 1)}, tXmeiIWc: 0},
	}
	switch runtime.GOOS {
	case "dragonfly", "openbsd":

		e_JzQ29Jvp = e_JzQ29Jvp[:1]
	}
	for eaAUaB7Z := range e_JzQ29Jvp {
		dRoFccaZ, h_ljk48Bm :=  /*line :1*/ bbSHyj(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
		if h_ljk48Bm != nil {
			continue
		}
		defer  /*line :1*/ poll.KvwyJafTkVwj(dRoFccaZ)
		 /*line :1*/ syscall.ABpNrEFh9ZGb(dRoFccaZ, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY, e_JzQ29Jvp[eaAUaB7Z].tXmeiIWc)
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ e_JzQ29Jvp[eaAUaB7Z].g28MY22awQ.jIv0wLJwloQ(syscall.AF_INET6)
		if h_ljk48Bm != nil {
			continue
		}
		if h_ljk48Bm :=  /*line :1*/ syscall.U5lDiJBkm0(dRoFccaZ, epTaNF0RJ8eN); h_ljk48Bm != nil {
			continue
		}
		if eaAUaB7Z == 0 {
			fMPVz2iGiyq.dERABKyXQio = true
		} else {
			fMPVz2iGiyq.bdF0JaBaZ = true
		}
	}
}

func atc3pbW1N(vsbiWLn7reX string, bFNyUpAx, wcA44hhD jIv0wLJwloQ, bcbtTs string) (nqsgnnaf int, q_vcfH bool) {
	switch vsbiWLn7reX[ /*line :1*/ len(vsbiWLn7reX)-1] {
	case '4':
		return syscall.AF_INET, false
	case '6':
		return syscall.AF_INET6, true
	}

	if bcbtTs == "listen" && (bFNyUpAx == nil ||  /*line :1*/ bFNyUpAx.tmKV_f7MKX()) {
		if  /*line :1*/ gyok94ZA() || ! /*line :1*/ kD5IbHeEWxIx() {
			return syscall.AF_INET6, false
		}
		if bFNyUpAx == nil {
			return syscall.AF_INET, false
		}
		return  /*line :1*/ bFNyUpAx.nqsgnnaf(), false
	}

	if (bFNyUpAx == nil ||  /*line :1*/ bFNyUpAx.nqsgnnaf() == syscall.AF_INET) &&
		(wcA44hhD == nil ||  /*line :1*/ wcA44hhD.nqsgnnaf() == syscall.AF_INET) {
		return syscall.AF_INET, false
	}
	return syscall.AF_INET6, false
}

func e6nNXDsTnM(yESLyde9LHhT context.MBCyA6, gF9mZs2 string, bFNyUpAx, wcA44hhD jIv0wLJwloQ, dr86pU, iaebRs5X3 int, bcbtTs string, rhsSAD7C func(context.MBCyA6, string, string, syscall.CVnV1i) error) (vyaiiPXm6k_W *jmJczkl1, h_ljk48Bm error) {
	if (runtime.GOOS == "aix" || runtime.GOOS == "windows" || runtime.GOOS == "openbsd") && bcbtTs == "dial" &&  /*line :1*/ wcA44hhD.tmKV_f7MKX() {
		wcA44hhD =  /*line :1*/ wcA44hhD.kAjqX7j(gF9mZs2)
	}
	nqsgnnaf, q_vcfH :=  /*line :1*/ atc3pbW1N(gF9mZs2, bFNyUpAx, wcA44hhD, bcbtTs)
	return  /*line :1*/ _tP9pgLw(yESLyde9LHhT, gF9mZs2, nqsgnnaf, dr86pU, iaebRs5X3, q_vcfH, bFNyUpAx, wcA44hhD, rhsSAD7C)
}

func qBOiy4(kQ8_UEhxU GNraIvYhB, mzUgFPgljgC int) (syscall.Hx8lqxSkV, error) {
	if  /*line :1*/ len(kQ8_UEhxU) == 0 {
		kQ8_UEhxU = WHSMKoapR
	}
	di2u16 :=  /*line :1*/ kQ8_UEhxU.To4()
	if di2u16 == nil {
		return syscall.Hx8lqxSkV{}, &DAWLIQHc{Z68v0y: "non-IPv4 address", BCCnAgFxG:  /*line :1*/ kQ8_UEhxU.String()}
	}
	epTaNF0RJ8eN := syscall.Hx8lqxSkV{AMBsbT8war: mzUgFPgljgC}
	 /*line :1*/ copy(epTaNF0RJ8eN.Q3zz2uH[:], di2u16)
	return epTaNF0RJ8eN, nil
}

func eWiFFUFYjR(kQ8_UEhxU GNraIvYhB, mzUgFPgljgC int, dJhOHOEZlQAA string) (syscall.HKTAy7_BSU2, error) {

	if  /*line :1*/ len(kQ8_UEhxU) == 0 ||  /*line :1*/ kQ8_UEhxU.Equal(WHSMKoapR) {
		kQ8_UEhxU = Zi20Wn
	}

	d45x3LPA9WTE :=  /*line :1*/ kQ8_UEhxU.To16()
	if d45x3LPA9WTE == nil {
		return syscall.HKTAy7_BSU2{}, &DAWLIQHc{Z68v0y: "non-IPv6 address", BCCnAgFxG:  /*line :1*/ kQ8_UEhxU.String()}
	}
	epTaNF0RJ8eN := syscall.HKTAy7_BSU2{NtLA3k: mzUgFPgljgC, AsutdJq2:  /*line :1*/ uint32( /*line :1*/ x4vxFT.eJTgFgJkWaQ(dJhOHOEZlQAA))}
	 /*line :1*/ copy(epTaNF0RJ8eN.Y4XxFr[:], d45x3LPA9WTE)
	return epTaNF0RJ8eN, nil
}

func ggQVAaURvc(nqsgnnaf int, kQ8_UEhxU GNraIvYhB, mzUgFPgljgC int, dJhOHOEZlQAA string) (syscall.S4iroLVT4, error) {
	switch nqsgnnaf {
	case syscall.AF_INET:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ qBOiy4(kQ8_UEhxU, mzUgFPgljgC)
		if h_ljk48Bm != nil {
			return nil, h_ljk48Bm
		}
		return &epTaNF0RJ8eN, nil
	case syscall.AF_INET6:
		epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ eWiFFUFYjR(kQ8_UEhxU, mzUgFPgljgC, dJhOHOEZlQAA)
		if h_ljk48Bm != nil {
			return nil, h_ljk48Bm
		}
		return &epTaNF0RJ8eN, nil
	}
	return nil, &DAWLIQHc{Z68v0y: "invalid address family", BCCnAgFxG:  /*line :1*/ kQ8_UEhxU.String()}
}

func xifkdya(dqpqJJfc netip.YTqTf_4VC) (syscall.Hx8lqxSkV, error) {

	qxwkC3VxG0 :=  /*line :1*/ dqpqJJfc.Addr()
	if ! /*line :1*/ qxwkC3VxG0.Is4() {
		return syscall.Hx8lqxSkV{}, &DAWLIQHc{Z68v0y: "non-IPv4 address", BCCnAgFxG:  /*line :1*/ qxwkC3VxG0.String()}
	}
	epTaNF0RJ8eN := syscall.Hx8lqxSkV{
		Q3zz2uH:	 /*line :1*/ qxwkC3VxG0.As4(),
		AMBsbT8war:	 /*line :1*/ int( /*line :1*/ dqpqJJfc.Port()),
	}
	return epTaNF0RJ8eN, nil
}

func tk4LBjOZ_EZy(dqpqJJfc netip.YTqTf_4VC) (syscall.HKTAy7_BSU2, error) {

	qxwkC3VxG0 :=  /*line :1*/ dqpqJJfc.Addr()
	if ! /*line :1*/ qxwkC3VxG0.IsValid() {
		return syscall.HKTAy7_BSU2{}, &DAWLIQHc{Z68v0y: "non-IPv6 address", BCCnAgFxG:  /*line :1*/ qxwkC3VxG0.String()}
	}
	epTaNF0RJ8eN := syscall.HKTAy7_BSU2{
		Y4XxFr:	 /*line :1*/ qxwkC3VxG0.As16(),
		NtLA3k:	 /*line :1*/ int( /*line :1*/ dqpqJJfc.Port()),
		AsutdJq2:	 /*line :1*/ uint32( /*line :1*/ x4vxFT.eJTgFgJkWaQ( /*line :1*/ qxwkC3VxG0.Zone())),
	}
	return epTaNF0RJ8eN, nil
}
