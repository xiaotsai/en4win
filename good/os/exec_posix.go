//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package hPMrClpC

import (
	itoa "JkjxVS"
	execenv "pu5p3KX8sW"
	"runtime"
	syscall "bUKeamF"
)

var (
	TkUKDRWq	J0zkNT	= syscall.SIGINT
	RebqCuoID	J0zkNT	= syscall.SIGKILL
)

func rNvAmv3ii(jGBs03 string, aXiZTmqaf []string, yEeSjJ *A5CIaIZeDjn) (baV7tO7i5 *P2XtkGf3, ugqb2IW error) {

	if yEeSjJ != nil && yEeSjJ.ZMj6ixN == nil && yEeSjJ.NX03GuKop6 != "" {
		if _, ugqb2IW :=  /*line :1*/ ZYa3wuv(yEeSjJ.NX03GuKop6); ugqb2IW != nil {
			ol2E6Cw := ugqb2IW.(*StNL1lveD40f)
			ol2E6Cw.Aeg62j1VQt = "chdir"
			return nil, ol2E6Cw
		}
	}

	iURRBbIb := &syscall.Ze0uPJo{
		L4Vaw9K:	yEeSjJ.NX03GuKop6,
		PZzy6B:	yEeSjJ.Ws3LdrwSmPa,
		PaHqSua0vh1:	yEeSjJ.ZMj6ixN,
	}
	if iURRBbIb.PZzy6B == nil {
		iURRBbIb.PZzy6B, ugqb2IW =  /*line :1*/ execenv.CYBtWdq6mE(iURRBbIb.PaHqSua0vh1)
		if ugqb2IW != nil {
			return nil, ugqb2IW
		}
	}
	iURRBbIb.INpA17B =  /*line :1*/ make([]uintptr, 0,  /*line :1*/ len(yEeSjJ.TXyhq0))
	for _, qzeVi328uMg := range yEeSjJ.TXyhq0 {
		iURRBbIb.INpA17B =  /*line :1*/ append(iURRBbIb.INpA17B,  /*line :1*/ qzeVi328uMg.Fd())
	}

	gAIGP2p, dB59YpfJ, w50uhPri :=  /*line :1*/ syscall.Z0IL9N5Otmal(jGBs03, aXiZTmqaf, iURRBbIb)

	 /*line :1*/ runtime.KeepAlive(yEeSjJ)

	if w50uhPri != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "fork/exec", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}

	return  /*line :1*/ zvDudDSfa(gAIGP2p, dB59YpfJ), nil
}

func (baV7tO7i5 *P2XtkGf3) xbit3b7ALF() error {
	return  /*line :1*/ baV7tO7i5.Signal(RebqCuoID)
}

type LcYXxKH struct {
	mMdCybHDxid	int
	mleSse0V	syscall.HV8Mf1
	t9b_DGjwz	*syscall.Pd6f6PxVa
}

func (baV7tO7i5 *LcYXxKH) Pid() int {
	return baV7tO7i5.mMdCybHDxid
}

func (baV7tO7i5 *LcYXxKH) uOgff09Wu1iI() bool {
	return  /*line :1*/ baV7tO7i5.mleSse0V.Exited()
}

func (baV7tO7i5 *LcYXxKH) bQvzsO() bool {
	return  /*line :1*/ baV7tO7i5.mleSse0V.ExitStatus() == 0
}

func (baV7tO7i5 *LcYXxKH) cx130qAaxcYa() any {
	return baV7tO7i5.mleSse0V
}

func (baV7tO7i5 *LcYXxKH) xQWuBvE7A15o() any {
	return baV7tO7i5.t9b_DGjwz
}

func (baV7tO7i5 *LcYXxKH) String() string {
	if baV7tO7i5 == nil {
		return "<nil>"
	}
	ryHP0HXqZA :=  /*line :1*/ baV7tO7i5.Sys().(syscall.HV8Mf1)
	mawaECh := ""
	switch {
	case  /*line :1*/ ryHP0HXqZA.Exited():
		aGXbi7Dvax :=  /*line :1*/ ryHP0HXqZA.ExitStatus()
		if runtime.GOOS == "windows" &&  /*line :1*/ uint(aGXbi7Dvax) >= 1<<16 {
			mawaECh = "exit status " +  /*line :1*/ tsxkt9sR( /*line :1*/ uint(aGXbi7Dvax))
		} else {
			mawaECh = "exit status " +  /*line :1*/ itoa.V25ba9G5(aGXbi7Dvax)
		}
	case  /*line :1*/ ryHP0HXqZA.Signaled():
		mawaECh = "signal: " +  /*line :1*/ ryHP0HXqZA.Signal().String()
	case  /*line :1*/ ryHP0HXqZA.Stopped():
		mawaECh = "stop signal: " +  /*line :1*/ ryHP0HXqZA.StopSignal().String()
		if  /*line :1*/ ryHP0HXqZA.StopSignal() == syscall.SIGTRAP &&  /*line :1*/ ryHP0HXqZA.TrapCause() != 0 {
			mawaECh += " (trap " +  /*line :1*/ itoa.V25ba9G5( /*line :1*/ ryHP0HXqZA.TrapCause()) + ")"
		}
	case  /*line :1*/ ryHP0HXqZA.Continued():
		mawaECh = "continued"
	}
	if  /*line :1*/ ryHP0HXqZA.CoreDump() {
		mawaECh += " (core dumped)"
	}
	return mawaECh
}

func (baV7tO7i5 *LcYXxKH) ExitCode() int {

	if baV7tO7i5 == nil {
		return -1
	}
	return  /*line :1*/ baV7tO7i5.mleSse0V.ExitStatus()
}
