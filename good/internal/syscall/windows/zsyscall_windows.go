//line :1
package LdptURlN

import (
	sysdll "O8h6BQj"
	syscall "bUKeamF"
	"unsafe"
)

var _ unsafe.Pointer



const (
	errnoERROR_IO_PENDING = 997
)

var (
	izPmTbuvZ	error	=  /*line :1*/ syscall.O9Mga3(errnoERROR_IO_PENDING)
	naDtVSDdWrOa	error	= syscall.EINVAL
)



func pBI0QG(b5udNLYT9oi syscall.O9Mga3) error {
	switch b5udNLYT9oi {
	case 0:
		return naDtVSDdWrOa
	case errnoERROR_IO_PENDING:
		return izPmTbuvZ
	}

	return b5udNLYT9oi
}

var (
	h_U_Yeo5	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("advapi32.dll"))
	kmYawC6aue_	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("iphlpapi.dll"))
	eRyS7zmeUJS	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("kernel32.dll"))
	xwft72E_W	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("netapi32.dll"))
	gM6LpKKCj7sO	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("psapi.dll"))
	goONY0hs40U	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("userenv.dll"))
	zal5rjTSown	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("ws2_32.dll"))

	e7Jz7xHgR	=  /*line :1*/ h_U_Yeo5.NewProc("AdjustTokenPrivileges")
	su9Sm1Jp9	=  /*line :1*/ h_U_Yeo5.NewProc("DuplicateTokenEx")
	b2x8aAOaZe	=  /*line :1*/ h_U_Yeo5.NewProc("ImpersonateSelf")
	p95piFDLtAM	=  /*line :1*/ h_U_Yeo5.NewProc("LookupPrivilegeValueW")
	ndYe5npiJc	=  /*line :1*/ h_U_Yeo5.NewProc("OpenThreadToken")
	v5KdSW6qDgA	=  /*line :1*/ h_U_Yeo5.NewProc("RevertToSelf")
	xaWQhQ	=  /*line :1*/ h_U_Yeo5.NewProc("SetTokenInformation")
	akbQ5tgTelV	=  /*line :1*/ h_U_Yeo5.NewProc("SystemFunction036")
	bBJMhL5GSB	=  /*line :1*/ kmYawC6aue_.NewProc("GetAdaptersAddresses")
	cNQSD189WPX	=  /*line :1*/ eRyS7zmeUJS.NewProc("CreateEventW")
	g0xg_yyF	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetACP")
	q7Gqnbb	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetComputerNameExW")
	i6mWEDC	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetConsoleCP")
	dZZ4zII_2_MK	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetCurrentThread")
	sVgvImVcY	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetFileInformationByHandleEx")
	kJ570MGiPPuX	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetFinalPathNameByHandleW")
	dSxglDxsa	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetModuleFileNameW")
	bGXtbzu	=  /*line :1*/ eRyS7zmeUJS.NewProc("GetTempPath2W")
	bv7UvhJawOi	=  /*line :1*/ eRyS7zmeUJS.NewProc("LockFileEx")
	bS0mpi8dPx	=  /*line :1*/ eRyS7zmeUJS.NewProc("Module32FirstW")
	atjQaAPj	=  /*line :1*/ eRyS7zmeUJS.NewProc("Module32NextW")
	pMsMsb	=  /*line :1*/ eRyS7zmeUJS.NewProc("MoveFileExW")
	fPgvDFSK16H	=  /*line :1*/ eRyS7zmeUJS.NewProc("MultiByteToWideChar")
	d5kRPnZ	=  /*line :1*/ eRyS7zmeUJS.NewProc("RtlLookupFunctionEntry")
	aVCQeZ__	=  /*line :1*/ eRyS7zmeUJS.NewProc("RtlVirtualUnwind")
	zcIJsZ3	=  /*line :1*/ eRyS7zmeUJS.NewProc("SetFileInformationByHandle")
	a0O1CAIc	=  /*line :1*/ eRyS7zmeUJS.NewProc("UnlockFileEx")
	aC6iXDK9lJI	=  /*line :1*/ eRyS7zmeUJS.NewProc("VirtualQuery")
	dNcVGpTifKxA	=  /*line :1*/ xwft72E_W.NewProc("NetShareAdd")
	tL4bg6iiK	=  /*line :1*/ xwft72E_W.NewProc("NetShareDel")
	rxLrHoYEw	=  /*line :1*/ xwft72E_W.NewProc("NetUserGetLocalGroups")
	eOZuXVMWJNL	=  /*line :1*/ gM6LpKKCj7sO.NewProc("GetProcessMemoryInfo")
	zK0ZMtSdQ7e	=  /*line :1*/ goONY0hs40U.NewProc("CreateEnvironmentBlock")
	aS4xPjAKSfSf	=  /*line :1*/ goONY0hs40U.NewProc("DestroyEnvironmentBlock")
	gDUxdtLWMwS	=  /*line :1*/ goONY0hs40U.NewProc("GetProfilesDirectoryW")
	w7rmsXibw	=  /*line :1*/ zal5rjTSown.NewProc("WSASocketW")
)

func ph35lMCuC(hp67APc637J3 syscall.Ad4bWd, hls6MaIN bool, eqAC0x1d *LFo8HyR, xUtG63NsJD uint32, l7Zkpa1NV6 *LFo8HyR, hkHgRkM7M *uint32) (jxAixT uint32, zc4PmxS error) {
	var mtoTPR uint32
	if hls6MaIN {
		mtoTPR = 1
	}
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ e7Jz7xHgR.Addr(), 6,  /*line :1*/ uintptr(hp67APc637J3),  /*line :1*/ uintptr(mtoTPR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eqAC0x1d)),  /*line :1*/ uintptr(xUtG63NsJD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(l7Zkpa1NV6)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hkHgRkM7M)))
	jxAixT =  /*line :1*/ uint32(mh8op6)
	if true {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func LEGhQXWr(be7fhbO5R9M syscall.Ad4bWd, fTd6bnP6MFa uint32, xTEbGbV5 *syscall.DDwuM6RpYja, aFvrdT uint32, iAPE2jqeDPi Zs2_lZoOk, fn34gBb0Y *syscall.Ad4bWd) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ su9Sm1Jp9.Addr(), 6,  /*line :1*/ uintptr(be7fhbO5R9M),  /*line :1*/ uintptr(fTd6bnP6MFa),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xTEbGbV5)),  /*line :1*/ uintptr(aFvrdT),  /*line :1*/ uintptr(iAPE2jqeDPi),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fn34gBb0Y)))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func UtZtc93lD3X(uUZ2wAE2 uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ b2x8aAOaZe.Addr(), 1,  /*line :1*/ uintptr(uUZ2wAE2), 0, 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func JoHARDS1(uHLgUSEBzD *uint16, hYBfjjgen_ *uint16, rwUfVOtBlm7 *PkWXUB7wm) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ p95piFDLtAM.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uHLgUSEBzD)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hYBfjjgen_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rwUfVOtBlm7)))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func OXrOoPrrkh(_ZwJmIWb5 syscall.SRlvVjrYa, dJrtnINt2e uint32, mhx_9fykjP bool, hp67APc637J3 *syscall.Ad4bWd) (zc4PmxS error) {
	var mtoTPR uint32
	if mhx_9fykjP {
		mtoTPR = 1
	}
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ ndYe5npiJc.Addr(), 4,  /*line :1*/ uintptr(_ZwJmIWb5),  /*line :1*/ uintptr(dJrtnINt2e),  /*line :1*/ uintptr(mtoTPR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hp67APc637J3)), 0, 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func TNAUtfh() (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ v5KdSW6qDgA.Addr(), 0, 0, 0, 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func SrvMLH(hHhXyaemA syscall.Ad4bWd, c6BJAE uint32, bmVN38GKV uintptr, amoBA91xI uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ xaWQhQ.Addr(), 4,  /*line :1*/ uintptr(hHhXyaemA),  /*line :1*/ uintptr(c6BJAE),  /*line :1*/ uintptr(bmVN38GKV),  /*line :1*/ uintptr(amoBA91xI), 0, 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func CCXhqGcd6Kw(yL26gxsrOT80 []byte) (zc4PmxS error) {
	var mtoTPR *byte
	if  /*line :1*/ len(yL26gxsrOT80) > 0 {
		mtoTPR = &yL26gxsrOT80[0]
	}
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ akbQ5tgTelV.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mtoTPR)),  /*line :1*/ uintptr( /*line :1*/ len(yL26gxsrOT80)), 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func ViNAxxzfGBB(jpcbaz uint32, iK5SVBoZG uint32, jDNnYnE uintptr, ma90bO *WXavkAObzxI, lBojP74y *uint32) (uDH5jepOXIbR error) {
	mh8op6, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ bBJMhL5GSB.Addr(), 5,  /*line :1*/ uintptr(jpcbaz),  /*line :1*/ uintptr(iK5SVBoZG),  /*line :1*/ uintptr(jDNnYnE),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ma90bO)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lBojP74y)), 0)
	if mh8op6 != 0 {
		uDH5jepOXIbR =  /*line :1*/ syscall.O9Mga3(mh8op6)
	}
	return
}

func QWFYx9x(eGz3Xf9UEla *BR3dBXMrt1du, yj9NUlq uint32, bWgv145aBTy uint32, hYBfjjgen_ *uint16) (lCclitZOD0C syscall.SRlvVjrYa, zc4PmxS error) {
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ cNQSD189WPX.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eGz3Xf9UEla)),  /*line :1*/ uintptr(yj9NUlq),  /*line :1*/ uintptr(bWgv145aBTy),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hYBfjjgen_)), 0, 0)
	lCclitZOD0C =  /*line :1*/ syscall.SRlvVjrYa(mh8op6)
	if lCclitZOD0C == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func TLIw5zSY() (tI7nU5z uint32) {
	mh8op6, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ g0xg_yyF.Addr(), 0, 0, 0, 0)
	tI7nU5z =  /*line :1*/ uint32(mh8op6)
	return
}

func YyY7M_287(eax7su uint32, yL26gxsrOT80 *uint16, pdOCpIN7 *uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ q7Gqnbb.Addr(), 3,  /*line :1*/ uintptr(eax7su),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yL26gxsrOT80)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pdOCpIN7)))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func EQZX7D() (gxOOSq uint32) {
	mh8op6, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ i6mWEDC.Addr(), 0, 0, 0, 0)
	gxOOSq =  /*line :1*/ uint32(mh8op6)
	return
}

func J96G2l() (lWVlMt4 syscall.SRlvVjrYa, zc4PmxS error) {
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dZZ4zII_2_MK.Addr(), 0, 0, 0, 0)
	lWVlMt4 =  /*line :1*/ syscall.SRlvVjrYa(mh8op6)
	if lWVlMt4 == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func JaAxDHL(lCclitZOD0C syscall.SRlvVjrYa, vikBNbY uint32, sTTW8adeka *byte, fvnaX5sfeIpg uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ sVgvImVcY.Addr(), 4,  /*line :1*/ uintptr(lCclitZOD0C),  /*line :1*/ uintptr(vikBNbY),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sTTW8adeka)),  /*line :1*/ uintptr(fvnaX5sfeIpg), 0, 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func DoAI5N(p1B3vu2Dcw syscall.SRlvVjrYa, hxt3evJ *uint16, k8oHaY6o uint32, iK5SVBoZG uint32) (pdOCpIN7 uint32, zc4PmxS error) {
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ kJ570MGiPPuX.Addr(), 4,  /*line :1*/ uintptr(p1B3vu2Dcw),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hxt3evJ)),  /*line :1*/ uintptr(k8oHaY6o),  /*line :1*/ uintptr(iK5SVBoZG), 0, 0)
	pdOCpIN7 =  /*line :1*/ uint32(mh8op6)
	if pdOCpIN7 == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func GCd4lr(mjugcN4 syscall.SRlvVjrYa, dHb9bvZU *uint16, lh2L5DoJ uint32) (pdOCpIN7 uint32, zc4PmxS error) {
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dSxglDxsa.Addr(), 3,  /*line :1*/ uintptr(mjugcN4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dHb9bvZU)),  /*line :1*/ uintptr(lh2L5DoJ))
	pdOCpIN7 =  /*line :1*/ uint32(mh8op6)
	if pdOCpIN7 == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func RdXWiJTR8(xUtG63NsJD uint32, yL26gxsrOT80 *uint16) (pdOCpIN7 uint32, zc4PmxS error) {
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bGXtbzu.Addr(), 2,  /*line :1*/ uintptr(xUtG63NsJD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yL26gxsrOT80)), 0)
	pdOCpIN7 =  /*line :1*/ uint32(mh8op6)
	if pdOCpIN7 == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func OyxdadcVqt_(p1B3vu2Dcw syscall.SRlvVjrYa, iK5SVBoZG uint32, jDNnYnE uint32, tpudt_poe uint32, fpZs7JyEvfsD uint32, bS4DAcuB0sz *syscall.ZaNm5QSYC9) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ bv7UvhJawOi.Addr(), 6,  /*line :1*/ uintptr(p1B3vu2Dcw),  /*line :1*/ uintptr(iK5SVBoZG),  /*line :1*/ uintptr(jDNnYnE),  /*line :1*/ uintptr(tpudt_poe),  /*line :1*/ uintptr(fpZs7JyEvfsD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bS4DAcuB0sz)))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func KUiK3Ucqw(besUFnE9NGA syscall.SRlvVjrYa, g92W3c95ps *LX2aqkMG_K) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bS0mpi8dPx.Addr(), 2,  /*line :1*/ uintptr(besUFnE9NGA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g92W3c95ps)), 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func YD5AyJMDb(besUFnE9NGA syscall.SRlvVjrYa, g92W3c95ps *LX2aqkMG_K) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ atjQaAPj.Addr(), 2,  /*line :1*/ uintptr(besUFnE9NGA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g92W3c95ps)), 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func S6mU1WXHl(y6SejQ5f *uint16, uxhDNcje *uint16, iK5SVBoZG uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ pMsMsb.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y6SejQ5f)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uxhDNcje)),  /*line :1*/ uintptr(iK5SVBoZG))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func W29UxGjFmE(jv5YIV uint32, _En7W3 uint32, bwa5IvQD *byte, x3izHJUCeU int32, hqz5vdZopBIf *uint16, yISmi4 int32) (mahx1UuHZ int32, zc4PmxS error) {
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ fPgvDFSK16H.Addr(), 6,  /*line :1*/ uintptr(jv5YIV),  /*line :1*/ uintptr(_En7W3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bwa5IvQD)),  /*line :1*/ uintptr(x3izHJUCeU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hqz5vdZopBIf)),  /*line :1*/ uintptr(yISmi4))
	mahx1UuHZ =  /*line :1*/ int32(mh8op6)
	if mahx1UuHZ == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func KbaQqI(l4g0XZBwvR uintptr, fEL3hbqm *uintptr, x_bGBRUub *byte) (jxAixT uintptr) {
	mh8op6, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ d5kRPnZ.Addr(), 3,  /*line :1*/ uintptr(l4g0XZBwvR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fEL3hbqm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x_bGBRUub)))
	jxAixT =  /*line :1*/ uintptr(mh8op6)
	return
}

func ZaNqwF(s1a38Np uint32, fEL3hbqm uintptr, l4g0XZBwvR uintptr, dQcfV_x uintptr, cd79CQK uintptr, aHGeYsQchK *uintptr, p6U7Xc *uintptr, fai1jW *byte) (jxAixT uintptr) {
	mh8op6, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ aVCQeZ__.Addr(), 8,  /*line :1*/ uintptr(s1a38Np),  /*line :1*/ uintptr(fEL3hbqm),  /*line :1*/ uintptr(l4g0XZBwvR),  /*line :1*/ uintptr(dQcfV_x),  /*line :1*/ uintptr(cd79CQK),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aHGeYsQchK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(p6U7Xc)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fai1jW)), 0)
	jxAixT =  /*line :1*/ uintptr(mh8op6)
	return
}

func ApJRwykg1p(lCclitZOD0C syscall.SRlvVjrYa, ujW_nz9_8 uint32, yL26gxsrOT80 uintptr, fvnaX5sfeIpg uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ zcIJsZ3.Addr(), 4,  /*line :1*/ uintptr(lCclitZOD0C),  /*line :1*/ uintptr(ujW_nz9_8),  /*line :1*/ uintptr(yL26gxsrOT80),  /*line :1*/ uintptr(fvnaX5sfeIpg), 0, 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func FStmxj(p1B3vu2Dcw syscall.SRlvVjrYa, jDNnYnE uint32, tpudt_poe uint32, fpZs7JyEvfsD uint32, bS4DAcuB0sz *syscall.ZaNm5QSYC9) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ a0O1CAIc.Addr(), 5,  /*line :1*/ uintptr(p1B3vu2Dcw),  /*line :1*/ uintptr(jDNnYnE),  /*line :1*/ uintptr(tpudt_poe),  /*line :1*/ uintptr(fpZs7JyEvfsD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bS4DAcuB0sz)), 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func YxnRUvQSGQzs(eoHilJhuyCV3 uintptr, d4xEBaN0dams *Xyjgm3eVd8f, uD4foLjE1B uintptr) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aC6iXDK9lJI.Addr(), 3,  /*line :1*/ uintptr(eoHilJhuyCV3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d4xEBaN0dams)),  /*line :1*/ uintptr(uD4foLjE1B))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func OppZebAT3Pv(nWaGKlR *uint16, miklVEjp uint32, yL26gxsrOT80 *byte, dZVAtKuy *uint16) (xCbhJgGczf8 error) {
	mh8op6, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dNcVGpTifKxA.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nWaGKlR)),  /*line :1*/ uintptr(miklVEjp),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yL26gxsrOT80)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dZVAtKuy)), 0, 0)
	if mh8op6 != 0 {
		xCbhJgGczf8 =  /*line :1*/ syscall.O9Mga3(mh8op6)
	}
	return
}

func RZ0gawYU9F4(nWaGKlR *uint16, prx_Uxe4D *uint16, jDNnYnE uint32) (xCbhJgGczf8 error) {
	mh8op6, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ tL4bg6iiK.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nWaGKlR)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(prx_Uxe4D)),  /*line :1*/ uintptr(jDNnYnE))
	if mh8op6 != 0 {
		xCbhJgGczf8 =  /*line :1*/ syscall.O9Mga3(mh8op6)
	}
	return
}

func PkaxDnZ(nWaGKlR *uint16, xviWct *uint16, miklVEjp uint32, iK5SVBoZG uint32, yL26gxsrOT80 **byte, dROkHgrP2XM uint32, cq51vNm8Ljk *uint32, mh7sZvBawa *uint32) (xCbhJgGczf8 error) {
	mh8op6, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ rxLrHoYEw.Addr(), 8,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nWaGKlR)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xviWct)),  /*line :1*/ uintptr(miklVEjp),  /*line :1*/ uintptr(iK5SVBoZG),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yL26gxsrOT80)),  /*line :1*/ uintptr(dROkHgrP2XM),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cq51vNm8Ljk)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mh7sZvBawa)), 0)
	if mh8op6 != 0 {
		xCbhJgGczf8 =  /*line :1*/ syscall.O9Mga3(mh8op6)
	}
	return
}

func Z1WBX6xiLx(lCclitZOD0C syscall.SRlvVjrYa, bUNopvBY7 *UQEeR4Vpvpy, uqPKSXKfx1w uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eOZuXVMWJNL.Addr(), 3,  /*line :1*/ uintptr(lCclitZOD0C),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bUNopvBY7)),  /*line :1*/ uintptr(uqPKSXKfx1w))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func QPLQfl3(hHxTZQNWZe **uint16, hp67APc637J3 syscall.Ad4bWd, jUumWyCv bool) (zc4PmxS error) {
	var mtoTPR uint32
	if jUumWyCv {
		mtoTPR = 1
	}
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zK0ZMtSdQ7e.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hHxTZQNWZe)),  /*line :1*/ uintptr(hp67APc637J3),  /*line :1*/ uintptr(mtoTPR))
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func Bnr47zSrA6G(hHxTZQNWZe *uint16) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aS4xPjAKSfSf.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hHxTZQNWZe)), 0, 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func DZrOSkRbLCsS(fZH71iL0Mj5L *uint16, mucWlH *uint32) (zc4PmxS error) {
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gDUxdtLWMwS.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fZH71iL0Mj5L)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mucWlH)), 0)
	if cMhXXOuwZkJB == 0 {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}

func QK7PBK(olYD2keElN int32, qsIvDOY_WD0 int32, uS_9r52BZfU int32, kMcK4Alew0m *syscall.ZEPlE2KjP71, pEvBcHv464wa uint32, iK5SVBoZG uint32) (lCclitZOD0C syscall.SRlvVjrYa, zc4PmxS error) {
	mh8op6, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ w7rmsXibw.Addr(), 6,  /*line :1*/ uintptr(olYD2keElN),  /*line :1*/ uintptr(qsIvDOY_WD0),  /*line :1*/ uintptr(uS_9r52BZfU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kMcK4Alew0m)),  /*line :1*/ uintptr(pEvBcHv464wa),  /*line :1*/ uintptr(iK5SVBoZG))
	lCclitZOD0C =  /*line :1*/ syscall.SRlvVjrYa(mh8op6)
	if lCclitZOD0C == syscall.InvalidHandle {
		zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
	}
	return
}
