//line :1
//go:build unix || windows

package gF9mZs2

import (
	poll "MjXXMR"
	"runtime"
	syscall "bUKeamF"
	time "fRAfQd_"
)

type jmJczkl1 struct {
	q0t7ga7tBcV5	poll.ENfmDmMaozH

	zW6I0G4Xm0k	int
	zpK4ak_Tpl	int
	kBOwaEVMqK	bool
	fCaRuBA3	string
	y9pjEf	EqbMXsaU1lE
	jeyDlYdkG0nF	EqbMXsaU1lE
}

func (vyaiiPXm6k_W *jmJczkl1) xaamQ3HuZ(bFNyUpAx, wcA44hhD EqbMXsaU1lE) {
	vyaiiPXm6k_W.y9pjEf = bFNyUpAx
	vyaiiPXm6k_W.jeyDlYdkG0nF = wcA44hhD
	 /*line :1*/ runtime.SetFinalizer(vyaiiPXm6k_W, (*jmJczkl1).Close)
}

func (vyaiiPXm6k_W *jmJczkl1) Close() error {
	 /*line :1*/ runtime.SetFinalizer(vyaiiPXm6k_W, nil)
	return  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Close()
}

func (vyaiiPXm6k_W *jmJczkl1) ntyS77AwzG(qldFxXX int) error {
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Shutdown(qldFxXX)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("shutdown", h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) ea_e1ZnpFCQF() error {
	return  /*line :1*/ vyaiiPXm6k_W.ntyS77AwzG(syscall.SHUT_RD)
}

func (vyaiiPXm6k_W *jmJczkl1) oCRE0sQeNq() error {
	return  /*line :1*/ vyaiiPXm6k_W.ntyS77AwzG(syscall.SHUT_WR)
}

func (vyaiiPXm6k_W *jmJczkl1) Read(fMPVz2iGiyq []byte) (doauF8Y int, h_ljk48Bm error) {
	doauF8Y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Read(fMPVz2iGiyq)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y,  /*line :1*/ raib0nns(readSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) mA1mFe4YP(fMPVz2iGiyq []byte) (doauF8Y int, epTaNF0RJ8eN syscall.S4iroLVT4, h_ljk48Bm error) {
	doauF8Y, epTaNF0RJ8eN, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.ReadFrom(fMPVz2iGiyq)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y, epTaNF0RJ8eN,  /*line :1*/ raib0nns(readFromSyscallName, h_ljk48Bm)
}
func (vyaiiPXm6k_W *jmJczkl1) aqQMdv2(fMPVz2iGiyq []byte, sdGaORW_DS *syscall.Hx8lqxSkV) (doauF8Y int, h_ljk48Bm error) {
	doauF8Y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.ReadFromInet4(fMPVz2iGiyq, sdGaORW_DS)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y,  /*line :1*/ raib0nns(readFromSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) wI_ExTyC(fMPVz2iGiyq []byte, sdGaORW_DS *syscall.HKTAy7_BSU2) (doauF8Y int, h_ljk48Bm error) {
	doauF8Y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.ReadFromInet6(fMPVz2iGiyq, sdGaORW_DS)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y,  /*line :1*/ raib0nns(readFromSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) vTzqIG(fMPVz2iGiyq []byte, cFiVO46YAy8g []byte, dBNjCxLYqs int) (doauF8Y, ytEzh3580y, p2dSNx1Ma int, epTaNF0RJ8eN syscall.S4iroLVT4, h_ljk48Bm error) {
	doauF8Y, ytEzh3580y, p2dSNx1Ma, epTaNF0RJ8eN, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.ReadMsg(fMPVz2iGiyq, cFiVO46YAy8g, dBNjCxLYqs)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y, ytEzh3580y, p2dSNx1Ma, epTaNF0RJ8eN,  /*line :1*/ raib0nns(readMsgSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) gXSpGdZ05(fMPVz2iGiyq []byte, cFiVO46YAy8g []byte, dBNjCxLYqs int, epTaNF0RJ8eN *syscall.Hx8lqxSkV) (doauF8Y, ytEzh3580y, p2dSNx1Ma int, h_ljk48Bm error) {
	doauF8Y, ytEzh3580y, p2dSNx1Ma, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.ReadMsgInet4(fMPVz2iGiyq, cFiVO46YAy8g, dBNjCxLYqs, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y, ytEzh3580y, p2dSNx1Ma,  /*line :1*/ raib0nns(readMsgSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) jQ3N__FhSo(fMPVz2iGiyq []byte, cFiVO46YAy8g []byte, dBNjCxLYqs int, epTaNF0RJ8eN *syscall.HKTAy7_BSU2) (doauF8Y, ytEzh3580y, p2dSNx1Ma int, h_ljk48Bm error) {
	doauF8Y, ytEzh3580y, p2dSNx1Ma, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.ReadMsgInet6(fMPVz2iGiyq, cFiVO46YAy8g, dBNjCxLYqs, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y, ytEzh3580y, p2dSNx1Ma,  /*line :1*/ raib0nns(readMsgSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) Write(fMPVz2iGiyq []byte) (jBXrVi int, h_ljk48Bm error) {
	jBXrVi, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.Write(fMPVz2iGiyq)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return jBXrVi,  /*line :1*/ raib0nns(writeSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) omQ684zc(fMPVz2iGiyq []byte, epTaNF0RJ8eN syscall.S4iroLVT4) (doauF8Y int, h_ljk48Bm error) {
	doauF8Y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WriteTo(fMPVz2iGiyq, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y,  /*line :1*/ raib0nns(writeToSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) edtDmORC(fMPVz2iGiyq []byte, epTaNF0RJ8eN *syscall.Hx8lqxSkV) (doauF8Y int, h_ljk48Bm error) {
	doauF8Y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WriteToInet4(fMPVz2iGiyq, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y,  /*line :1*/ raib0nns(writeToSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) wPqovgf6d(fMPVz2iGiyq []byte, epTaNF0RJ8eN *syscall.HKTAy7_BSU2) (doauF8Y int, h_ljk48Bm error) {
	doauF8Y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WriteToInet6(fMPVz2iGiyq, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y,  /*line :1*/ raib0nns(writeToSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) fh4B50S(fMPVz2iGiyq []byte, cFiVO46YAy8g []byte, epTaNF0RJ8eN syscall.S4iroLVT4) (doauF8Y int, ytEzh3580y int, h_ljk48Bm error) {
	doauF8Y, ytEzh3580y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WriteMsg(fMPVz2iGiyq, cFiVO46YAy8g, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y, ytEzh3580y,  /*line :1*/ raib0nns(writeMsgSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) h92kZU2ngw(fMPVz2iGiyq []byte, cFiVO46YAy8g []byte, epTaNF0RJ8eN *syscall.Hx8lqxSkV) (doauF8Y int, ytEzh3580y int, h_ljk48Bm error) {
	doauF8Y, ytEzh3580y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WriteMsgInet4(fMPVz2iGiyq, cFiVO46YAy8g, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y, ytEzh3580y,  /*line :1*/ raib0nns(writeMsgSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) i9YFSO8(fMPVz2iGiyq []byte, cFiVO46YAy8g []byte, epTaNF0RJ8eN *syscall.HKTAy7_BSU2) (doauF8Y int, ytEzh3580y int, h_ljk48Bm error) {
	doauF8Y, ytEzh3580y, h_ljk48Bm =  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.WriteMsgInet6(fMPVz2iGiyq, cFiVO46YAy8g, epTaNF0RJ8eN)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return doauF8Y, ytEzh3580y,  /*line :1*/ raib0nns(writeMsgSyscallName, h_ljk48Bm)
}

func (vyaiiPXm6k_W *jmJczkl1) SetDeadline(aeaqWzxJu time.G4KDkI) error {
	return  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetDeadline(aeaqWzxJu)
}

func (vyaiiPXm6k_W *jmJczkl1) SetReadDeadline(aeaqWzxJu time.G4KDkI) error {
	return  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetReadDeadline(aeaqWzxJu)
}

func (vyaiiPXm6k_W *jmJczkl1) SetWriteDeadline(aeaqWzxJu time.G4KDkI) error {
	return  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetWriteDeadline(aeaqWzxJu)
}
