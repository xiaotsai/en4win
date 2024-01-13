//line :1
package bUKeamF

import (
	sysdll "O8h6BQj"
	"unsafe"
)

var _ unsafe.Pointer



const (
	errnoERROR_IO_PENDING = 997
)

var (
	qcRseBIDA2LF	error	=  /*line :1*/ O9Mga3(errnoERROR_IO_PENDING)
	oaAyEiQbVBRN	error	= EINVAL
)



func wxZh39lL(wAwFLr4AGHg O9Mga3) error {
	switch wAwFLr4AGHg {
	case 0:
		return oaAyEiQbVBRN
	case errnoERROR_IO_PENDING:
		return qcRseBIDA2LF
	}

	return wAwFLr4AGHg
}

var (
	oHlJL89c8XQ9	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("advapi32.dll"))
	vQ_zBn2	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("crypt32.dll"))
	kN9aoocGy	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("dnsapi.dll"))
	gukzgDwbpb8_	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("iphlpapi.dll"))
	jXlerCqn	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("kernel32.dll"))
	jswJvZmXL	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("mswsock.dll"))
	ipPQZPUbem3	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("netapi32.dll"))
	aNeJfzmvKLHk	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("ntdll.dll"))
	hE5iOXT	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("secur32.dll"))
	gPMhG5	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("shell32.dll"))
	dd9sVBjtk	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("userenv.dll"))
	qATSmxg	=  /*line :1*/ A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("ws2_32.dll"))

	_Hm7Nq	=  /*line :1*/ oHlJL89c8XQ9.NewProc("ConvertSidToStringSidW")
	vuHyfvtYYs3a	=  /*line :1*/ oHlJL89c8XQ9.NewProc("ConvertStringSidToSidW")
	nApsrr	=  /*line :1*/ oHlJL89c8XQ9.NewProc("CopySid")
	nveuPs	=  /*line :1*/ oHlJL89c8XQ9.NewProc("CreateProcessAsUserW")
	f0JJBOP	=  /*line :1*/ oHlJL89c8XQ9.NewProc("CryptAcquireContextW")
	g2mFoEfZ	=  /*line :1*/ oHlJL89c8XQ9.NewProc("CryptGenRandom")
	b_5DoPRfYi	=  /*line :1*/ oHlJL89c8XQ9.NewProc("CryptReleaseContext")
	eEpZW6NO7DPk	=  /*line :1*/ oHlJL89c8XQ9.NewProc("GetLengthSid")
	dQh81EKNM2	=  /*line :1*/ oHlJL89c8XQ9.NewProc("GetTokenInformation")
	g78yMSan	=  /*line :1*/ oHlJL89c8XQ9.NewProc("LookupAccountNameW")
	wx5yA4	=  /*line :1*/ oHlJL89c8XQ9.NewProc("LookupAccountSidW")
	kTj2LevYAw	=  /*line :1*/ oHlJL89c8XQ9.NewProc("OpenProcessToken")
	lcJCDEXB5t	=  /*line :1*/ oHlJL89c8XQ9.NewProc("RegCloseKey")
	kq95jPllXFVy	=  /*line :1*/ oHlJL89c8XQ9.NewProc("RegEnumKeyExW")
	dj6I7nK	=  /*line :1*/ oHlJL89c8XQ9.NewProc("RegOpenKeyExW")
	qAnHJWyA	=  /*line :1*/ oHlJL89c8XQ9.NewProc("RegQueryInfoKeyW")
	fO3YkUBFk	=  /*line :1*/ oHlJL89c8XQ9.NewProc("RegQueryValueExW")
	a4VXNvJx7	=  /*line :1*/ vQ_zBn2.NewProc("CertAddCertificateContextToStore")
	b5bCkXVR	=  /*line :1*/ vQ_zBn2.NewProc("CertCloseStore")
	ovI8K34apT8	=  /*line :1*/ vQ_zBn2.NewProc("CertCreateCertificateContext")
	a5uM9aJ	=  /*line :1*/ vQ_zBn2.NewProc("CertEnumCertificatesInStore")
	gzbm2K	=  /*line :1*/ vQ_zBn2.NewProc("CertFreeCertificateChain")
	dRoQJJIIQB	=  /*line :1*/ vQ_zBn2.NewProc("CertFreeCertificateContext")
	cGFKoL6	=  /*line :1*/ vQ_zBn2.NewProc("CertGetCertificateChain")
	qaQASfoPiG0	=  /*line :1*/ vQ_zBn2.NewProc("CertOpenStore")
	rNfpvrMIbs7	=  /*line :1*/ vQ_zBn2.NewProc("CertOpenSystemStoreW")
	x23fWxIPr	=  /*line :1*/ vQ_zBn2.NewProc("CertVerifyCertificateChainPolicy")
	dnzSZjpfq0	=  /*line :1*/ kN9aoocGy.NewProc("DnsNameCompare_W")
	hAo8t3aX	=  /*line :1*/ kN9aoocGy.NewProc("DnsQuery_W")
	wPq8lA	=  /*line :1*/ kN9aoocGy.NewProc("DnsRecordListFree")
	yuG7y66	=  /*line :1*/ gukzgDwbpb8_.NewProc("GetAdaptersInfo")
	bcgbKEiGt	=  /*line :1*/ gukzgDwbpb8_.NewProc("GetIfEntry")
	okr2Sl9wBb	=  /*line :1*/ jXlerCqn.NewProc("CancelIo")
	dC_GmCDgrR	=  /*line :1*/ jXlerCqn.NewProc("CancelIoEx")
	hAjxMUv	=  /*line :1*/ jXlerCqn.NewProc("CloseHandle")
	rLNmZY	=  /*line :1*/ jXlerCqn.NewProc("CreateDirectoryW")
	dH2uHyBd	=  /*line :1*/ jXlerCqn.NewProc("CreateFileMappingW")
	cCvppXKP6W83	=  /*line :1*/ jXlerCqn.NewProc("CreateFileW")
	zus5CKQbyQ	=  /*line :1*/ jXlerCqn.NewProc("CreateHardLinkW")
	w6QQ05j	=  /*line :1*/ jXlerCqn.NewProc("CreateIoCompletionPort")
	hZFdxL	=  /*line :1*/ jXlerCqn.NewProc("CreatePipe")
	uGfRBah	=  /*line :1*/ jXlerCqn.NewProc("CreateProcessW")
	soOK9W	=  /*line :1*/ jXlerCqn.NewProc("CreateSymbolicLinkW")
	cxDdWOFEip	=  /*line :1*/ jXlerCqn.NewProc("CreateToolhelp32Snapshot")
	xf2zr4bdg	=  /*line :1*/ jXlerCqn.NewProc("DeleteFileW")
	gVXfdl_I8ne	=  /*line :1*/ jXlerCqn.NewProc("DeleteProcThreadAttributeList")
	enRZYWOlh4	=  /*line :1*/ jXlerCqn.NewProc("DeviceIoControl")
	epbX1qoTStN	=  /*line :1*/ jXlerCqn.NewProc("DuplicateHandle")
	nNMja9	=  /*line :1*/ jXlerCqn.NewProc("ExitProcess")
	sahN1d8mp	=  /*line :1*/ jXlerCqn.NewProc("FindClose")
	cLocy4DjDO	=  /*line :1*/ jXlerCqn.NewProc("FindFirstFileW")
	c5whBLyge1K6	=  /*line :1*/ jXlerCqn.NewProc("FindNextFileW")
	jYR1ar9Y	=  /*line :1*/ jXlerCqn.NewProc("FlushFileBuffers")
	badahuFho	=  /*line :1*/ jXlerCqn.NewProc("FlushViewOfFile")
	zcnyMlEw4JIX	=  /*line :1*/ jXlerCqn.NewProc("FormatMessageW")
	uO8mgeTm3r	=  /*line :1*/ jXlerCqn.NewProc("FreeEnvironmentStringsW")
	u4kiFTZ	=  /*line :1*/ jXlerCqn.NewProc("FreeLibrary")
	xgzVmmArEtci	=  /*line :1*/ jXlerCqn.NewProc("GetCommandLineW")
	g56k60	=  /*line :1*/ jXlerCqn.NewProc("GetComputerNameW")
	rwK2LeIhP5d	=  /*line :1*/ jXlerCqn.NewProc("GetConsoleMode")
	jJW5EcVc	=  /*line :1*/ jXlerCqn.NewProc("GetCurrentDirectoryW")
	aI65Xv	=  /*line :1*/ jXlerCqn.NewProc("GetCurrentProcess")
	wuPQ4Gcnp57	=  /*line :1*/ jXlerCqn.NewProc("GetCurrentProcessId")
	dyW3aKwd5bn	=  /*line :1*/ jXlerCqn.NewProc("GetEnvironmentStringsW")
	iQgSceA	=  /*line :1*/ jXlerCqn.NewProc("GetEnvironmentVariableW")
	f3tIzkTRW	=  /*line :1*/ jXlerCqn.NewProc("GetExitCodeProcess")
	yVk9wi	=  /*line :1*/ jXlerCqn.NewProc("GetFileAttributesExW")
	atp9nGcx	=  /*line :1*/ jXlerCqn.NewProc("GetFileAttributesW")
	jBHRrXeJ	=  /*line :1*/ jXlerCqn.NewProc("GetFileInformationByHandle")
	pNkaasL	=  /*line :1*/ jXlerCqn.NewProc("GetFileType")
	f0xxVP	=  /*line :1*/ jXlerCqn.NewProc("GetFinalPathNameByHandleW")
	t_DfADTW	=  /*line :1*/ jXlerCqn.NewProc("GetFullPathNameW")
	wgfyKiKTY1Yo	=  /*line :1*/ jXlerCqn.NewProc("GetLastError")
	qzzLr7B	=  /*line :1*/ jXlerCqn.NewProc("GetLongPathNameW")
	bfdaj_wm9rA	=  /*line :1*/ jXlerCqn.NewProc("GetProcAddress")
	d1qXWjaARQ	=  /*line :1*/ jXlerCqn.NewProc("GetProcessTimes")
	jDDTbwhjDY	=  /*line :1*/ jXlerCqn.NewProc("GetQueuedCompletionStatus")
	hyUD4tkB	=  /*line :1*/ jXlerCqn.NewProc("GetShortPathNameW")
	xga4aeva	=  /*line :1*/ jXlerCqn.NewProc("GetStartupInfoW")
	twemfk7Xg	=  /*line :1*/ jXlerCqn.NewProc("GetStdHandle")
	dQ2MstIcfyPZ	=  /*line :1*/ jXlerCqn.NewProc("GetSystemTimeAsFileTime")
	epfs28nOzvtO	=  /*line :1*/ jXlerCqn.NewProc("GetTempPathW")
	rX6vZNDMu5dK	=  /*line :1*/ jXlerCqn.NewProc("GetTimeZoneInformation")
	c0AVg4p8r	=  /*line :1*/ jXlerCqn.NewProc("GetVersion")
	gBsNaY6plc	=  /*line :1*/ jXlerCqn.NewProc("InitializeProcThreadAttributeList")
	w7f9jzy	=  /*line :1*/ jXlerCqn.NewProc("LoadLibraryW")
	iJaCKplx	=  /*line :1*/ jXlerCqn.NewProc("LocalFree")
	e8Ujhp	=  /*line :1*/ jXlerCqn.NewProc("MapViewOfFile")
	vXQQlhZc6pyH	=  /*line :1*/ jXlerCqn.NewProc("MoveFileW")
	hyGFQMbH	=  /*line :1*/ jXlerCqn.NewProc("OpenProcess")
	yFqFoUe	=  /*line :1*/ jXlerCqn.NewProc("PostQueuedCompletionStatus")
	q5daQHwuvcZ	=  /*line :1*/ jXlerCqn.NewProc("Process32FirstW")
	sRLLKO4IwOQ	=  /*line :1*/ jXlerCqn.NewProc("Process32NextW")
	fVbbzya	=  /*line :1*/ jXlerCqn.NewProc("ReadConsoleW")
	qkBhr91ZH2	=  /*line :1*/ jXlerCqn.NewProc("ReadDirectoryChangesW")
	dVi2JIQfx	=  /*line :1*/ jXlerCqn.NewProc("ReadFile")
	a_camEYomia	=  /*line :1*/ jXlerCqn.NewProc("RemoveDirectoryW")
	iDm22rip	=  /*line :1*/ jXlerCqn.NewProc("SetCurrentDirectoryW")
	j6PdoeCJe5	=  /*line :1*/ jXlerCqn.NewProc("SetEndOfFile")
	fJAOac3a2Xi	=  /*line :1*/ jXlerCqn.NewProc("SetEnvironmentVariableW")
	xihm6IzgQA	=  /*line :1*/ jXlerCqn.NewProc("SetFileAttributesW")
	dAUCsfm	=  /*line :1*/ jXlerCqn.NewProc("SetFileCompletionNotificationModes")
	y6R4dSDuFII	=  /*line :1*/ jXlerCqn.NewProc("SetFilePointer")
	dn91DjN1bhX	=  /*line :1*/ jXlerCqn.NewProc("SetFileTime")
	glaMy98Qs5	=  /*line :1*/ jXlerCqn.NewProc("SetHandleInformation")
	jSMiME	=  /*line :1*/ jXlerCqn.NewProc("TerminateProcess")
	yiAURBm17Lus	=  /*line :1*/ jXlerCqn.NewProc("UnmapViewOfFile")
	jihmTDgco	=  /*line :1*/ jXlerCqn.NewProc("UpdateProcThreadAttribute")
	wNT4s8M	=  /*line :1*/ jXlerCqn.NewProc("VirtualLock")
	bWCcw5aOP3	=  /*line :1*/ jXlerCqn.NewProc("VirtualUnlock")
	xQGMZ_3dF	=  /*line :1*/ jXlerCqn.NewProc("WaitForSingleObject")
	tl81Ww_JQ	=  /*line :1*/ jXlerCqn.NewProc("WriteConsoleW")
	h6gtSIpLuhao	=  /*line :1*/ jXlerCqn.NewProc("WriteFile")
	nDAteP7QBXTA	=  /*line :1*/ jswJvZmXL.NewProc("AcceptEx")
	eNwean	=  /*line :1*/ jswJvZmXL.NewProc("GetAcceptExSockaddrs")
	datHkUf_q	=  /*line :1*/ jswJvZmXL.NewProc("TransmitFile")
	lLxcq4xc	=  /*line :1*/ ipPQZPUbem3.NewProc("NetApiBufferFree")
	celUNpxNnT	=  /*line :1*/ ipPQZPUbem3.NewProc("NetGetJoinInformation")
	oi9C38I3oW	=  /*line :1*/ ipPQZPUbem3.NewProc("NetUserGetInfo")
	fEJdp1o	=  /*line :1*/ aNeJfzmvKLHk.NewProc("RtlGetNtVersionNumbers")
	jaJ9KC	=  /*line :1*/ hE5iOXT.NewProc("GetUserNameExW")
	iLxgMTvV	=  /*line :1*/ hE5iOXT.NewProc("TranslateNameW")
	p3khPPCvCgiS	=  /*line :1*/ gPMhG5.NewProc("CommandLineToArgvW")
	iGKVJSlcS	=  /*line :1*/ dd9sVBjtk.NewProc("GetUserProfileDirectoryW")
	zHb9Zuab	=  /*line :1*/ qATSmxg.NewProc("FreeAddrInfoW")
	x25ZA6ZEaO4	=  /*line :1*/ qATSmxg.NewProc("GetAddrInfoW")
	zT5cP8	=  /*line :1*/ qATSmxg.NewProc("WSACleanup")
	ijij9sQGyfD	=  /*line :1*/ qATSmxg.NewProc("WSAEnumProtocolsW")
	aAmayrk2oVz	=  /*line :1*/ qATSmxg.NewProc("WSAIoctl")
	nWTVn6cies	=  /*line :1*/ qATSmxg.NewProc("WSARecv")
	eI4_1asU	=  /*line :1*/ qATSmxg.NewProc("WSARecvFrom")
	eFrrx0Kpf	=  /*line :1*/ qATSmxg.NewProc("WSASend")
	epTJJ70Fc	=  /*line :1*/ qATSmxg.NewProc("WSASendTo")
	dDWkr2Ttm	=  /*line :1*/ qATSmxg.NewProc("WSAStartup")
	f42Wfbh2XM	=  /*line :1*/ qATSmxg.NewProc("bind")
	hTUat1z8gv2	=  /*line :1*/ qATSmxg.NewProc("closesocket")
	jahCtX	=  /*line :1*/ qATSmxg.NewProc("connect")
	lX4IVJBYT	=  /*line :1*/ qATSmxg.NewProc("gethostbyname")
	anyxIG	=  /*line :1*/ qATSmxg.NewProc("getpeername")
	qGXtGudaEjdJ	=  /*line :1*/ qATSmxg.NewProc("getprotobyname")
	oHmw41Wc2g	=  /*line :1*/ qATSmxg.NewProc("getservbyname")
	prLAS7DXZwiX	=  /*line :1*/ qATSmxg.NewProc("getsockname")
	wowigi1g	=  /*line :1*/ qATSmxg.NewProc("getsockopt")
	m_csHJ	=  /*line :1*/ qATSmxg.NewProc("listen")
	j7CoLsioqaPP	=  /*line :1*/ qATSmxg.NewProc("ntohs")
	cwZfNC_Xk	=  /*line :1*/ qATSmxg.NewProc("setsockopt")
	ebH4vF2kh	=  /*line :1*/ qATSmxg.NewProc("shutdown")
	gkkRqSs	=  /*line :1*/ qATSmxg.NewProc("socket")
)

func WfNPKGlYn5(jeNRxf *CLoH6Pka1ufN, gfUn6j **uint16) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ _Hm7Nq.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jeNRxf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gfUn6j)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func ItKNZ1(gfUn6j *uint16, jeNRxf **CLoH6Pka1ufN) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ vuHyfvtYYs3a.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gfUn6j)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jeNRxf)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func KaBHYfBb(_AmReqD uint32, az9vzgB0db *CLoH6Pka1ufN, dG3RUour *CLoH6Pka1ufN) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ nApsrr.Addr(), 3,  /*line :1*/ uintptr(_AmReqD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(az9vzgB0db)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dG3RUour)))
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func MizSUaBO5Kk(pOLqd6J61Lz Ad4bWd, fmYYRC *uint16, oAZypulaO *uint16, u7efeBj9 *DDwuM6RpYja, cdwLgBk *DDwuM6RpYja, h5AurTaAMtM1 bool, bKz7BnjpP5Aa uint32, dqnOhJ8vse *uint16, qmKEwwXG_ *uint16, tv24Vf0a9XQj *J7Nv3S0G31Q, akZSwzochq *NBPVXpO3Kv) (gOCcQzbcC error) {
	var vJKEmfpV uint32
	if h5AurTaAMtM1 {
		vJKEmfpV = 1
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ IiL5Io0Q5I( /*line :1*/ nveuPs.Addr(), 11,  /*line :1*/ uintptr(pOLqd6J61Lz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fmYYRC)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oAZypulaO)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u7efeBj9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cdwLgBk)),  /*line :1*/ uintptr(vJKEmfpV),  /*line :1*/ uintptr(bKz7BnjpP5Aa),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dqnOhJ8vse)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qmKEwwXG_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tv24Vf0a9XQj)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(akZSwzochq)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func OwcELpW(d5WtP21 *SRlvVjrYa, kWeaS5yslD_ *uint16, enWD0FxL0 *uint16, neI4Xucdy uint32, uADzMh0kZ7hL uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ f0JJBOP.Addr(), 5,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d5WtP21)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kWeaS5yslD_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(enWD0FxL0)),  /*line :1*/ uintptr(neI4Xucdy),  /*line :1*/ uintptr(uADzMh0kZ7hL), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func RU0N6wRgd7mM(d5WtP21 SRlvVjrYa, izePNH7xS3 uint32, eun1Jlud5A *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ g2mFoEfZ.Addr(), 3,  /*line :1*/ uintptr(d5WtP21),  /*line :1*/ uintptr(izePNH7xS3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)))
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func SGHE6Da_(d5WtP21 SRlvVjrYa, uADzMh0kZ7hL uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ b_5DoPRfYi.Addr(), 2,  /*line :1*/ uintptr(d5WtP21),  /*line :1*/ uintptr(uADzMh0kZ7hL), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func BFQLdGf9_(jeNRxf *CLoH6Pka1ufN) (mnIhqmz uint32) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ eEpZW6NO7DPk.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jeNRxf)), 0, 0)
	mnIhqmz =  /*line :1*/ uint32(kDaalUAu24K)
	return
}

func T85EPJR4(mhNoXZ Ad4bWd, ggEgYvJmQeyK uint32, b65sRubEGAyW *byte, mlekZcqafT4 uint32, j4okr6O35R *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ dQh81EKNM2.Addr(), 5,  /*line :1*/ uintptr(mhNoXZ),  /*line :1*/ uintptr(ggEgYvJmQeyK),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b65sRubEGAyW)),  /*line :1*/ uintptr(mlekZcqafT4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j4okr6O35R)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func MyD9XBHt(gtiQta8 *uint16, kh4J_hE2wV *uint16, jeNRxf *CLoH6Pka1ufN, sAHOgmbx *uint32, b_XdKl *uint16, gt8VM690MrW *uint32, sf4VEva *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ g78yMSan.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gtiQta8)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kh4J_hE2wV)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jeNRxf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sAHOgmbx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b_XdKl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gt8VM690MrW)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sf4VEva)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Dra3S4q(gtiQta8 *uint16, jeNRxf *CLoH6Pka1ufN, hzeU8Tl *uint16, y9xWORA1z *uint32, b_XdKl *uint16, gt8VM690MrW *uint32, sf4VEva *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ wx5yA4.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gtiQta8)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jeNRxf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y9xWORA1z)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b_XdKl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gt8VM690MrW)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sf4VEva)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func MXjpZkrAT(vFaWuUyY SRlvVjrYa, wzWkRJs uint32, pOLqd6J61Lz *Ad4bWd) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ kTj2LevYAw.Addr(), 3,  /*line :1*/ uintptr(vFaWuUyY),  /*line :1*/ uintptr(wzWkRJs),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pOLqd6J61Lz)))
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func BqIBx3KRN4(nIB9uB1 SRlvVjrYa) (zlknHC error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ lcJCDEXB5t.Addr(), 1,  /*line :1*/ uintptr(nIB9uB1), 0, 0)
	if kDaalUAu24K != 0 {
		zlknHC =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func oNLWBp2Sr(nIB9uB1 SRlvVjrYa, bjKxPq uint32, hzeU8Tl *uint16, y9xWORA1z *uint32, r4wfXKbPIsE *uint32, zIaTh7G3 *uint16, hiS864dY *uint32, aDyiZ7NGC *T8WbUqAC3v) (zlknHC error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ Rc0O05UsV( /*line :1*/ kq95jPllXFVy.Addr(), 8,  /*line :1*/ uintptr(nIB9uB1),  /*line :1*/ uintptr(bjKxPq),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y9xWORA1z)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(r4wfXKbPIsE)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zIaTh7G3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hiS864dY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aDyiZ7NGC)), 0)
	if kDaalUAu24K != 0 {
		zlknHC =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func FRt5iWSq2fWU(nIB9uB1 SRlvVjrYa, kXW52Xg9 *uint16, tTWqSvX1 uint32, fDmGnWMla uint32, j2GWlZBnk *SRlvVjrYa) (zlknHC error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ dj6I7nK.Addr(), 5,  /*line :1*/ uintptr(nIB9uB1),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kXW52Xg9)),  /*line :1*/ uintptr(tTWqSvX1),  /*line :1*/ uintptr(fDmGnWMla),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j2GWlZBnk)), 0)
	if kDaalUAu24K != 0 {
		zlknHC =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func N2CMOGAc(nIB9uB1 SRlvVjrYa, zIaTh7G3 *uint16, hiS864dY *uint32, r4wfXKbPIsE *uint32, gavHro *uint32, pfHGpeEdH7 *uint32, ju05t1o *uint32, aQVppE7zi9 *uint32, mxymmgIdWW *uint32, qqlZhfH *uint32, nTPuhcvu *uint32, aDyiZ7NGC *T8WbUqAC3v) (zlknHC error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ IiL5Io0Q5I( /*line :1*/ qAnHJWyA.Addr(), 12,  /*line :1*/ uintptr(nIB9uB1),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zIaTh7G3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hiS864dY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(r4wfXKbPIsE)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gavHro)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pfHGpeEdH7)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ju05t1o)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aQVppE7zi9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mxymmgIdWW)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qqlZhfH)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nTPuhcvu)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aDyiZ7NGC)))
	if kDaalUAu24K != 0 {
		zlknHC =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func BcZh0IGWcpB(nIB9uB1 SRlvVjrYa, hzeU8Tl *uint16, r4wfXKbPIsE *uint32, kAEm4Rssxe *uint32, eun1Jlud5A *byte, izePNH7xS3 *uint32) (zlknHC error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ fO3YkUBFk.Addr(), 6,  /*line :1*/ uintptr(nIB9uB1),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(r4wfXKbPIsE)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kAEm4Rssxe)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(izePNH7xS3)))
	if kDaalUAu24K != 0 {
		zlknHC =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func QBtt7h(_lQfGCXzTo SRlvVjrYa, mfPdeXQ *IFUIVHU, awUv9QudOt uint32, sHGlFEM **IFUIVHU) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ a4VXNvJx7.Addr(), 4,  /*line :1*/ uintptr(_lQfGCXzTo),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mfPdeXQ)),  /*line :1*/ uintptr(awUv9QudOt),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sHGlFEM)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Cn_wDH3(_lQfGCXzTo SRlvVjrYa, uADzMh0kZ7hL uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ b5bCkXVR.Addr(), 2,  /*line :1*/ uintptr(_lQfGCXzTo),  /*line :1*/ uintptr(uADzMh0kZ7hL), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func F7ofZO9(c2BerP2 uint32, glAuQ2P *byte, kIOQyYOUG_ uint32) (sR6NrMFBa *IFUIVHU, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ ovI8K34apT8.Addr(), 3,  /*line :1*/ uintptr(c2BerP2),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(glAuQ2P)),  /*line :1*/ uintptr(kIOQyYOUG_))
	sR6NrMFBa = (* /*line :1*/ IFUIVHU)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	if sR6NrMFBa == nil {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func QFgFHhl(_lQfGCXzTo SRlvVjrYa, o1T6GXNURC *IFUIVHU) (sR6NrMFBa *IFUIVHU, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ a5uM9aJ.Addr(), 2,  /*line :1*/ uintptr(_lQfGCXzTo),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(o1T6GXNURC)), 0)
	sR6NrMFBa = (* /*line :1*/ IFUIVHU)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	if sR6NrMFBa == nil {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func F99ySoxQ7kqT(lDwOnJogV *SoXj8D) {
	 /*line :1*/ OwKnfO5tgvm4( /*line :1*/ gzbm2K.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lDwOnJogV)), 0, 0)
	return
}

func EocJ2WI3sbAG(lDwOnJogV *IFUIVHU) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ dRoQJJIIQB.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lDwOnJogV)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func G5gtSYgA_(owi56xe SRlvVjrYa, akEz9oY8Rh *IFUIVHU, eTypM8uhmGhb *T8WbUqAC3v, bamTPS0GgqZ SRlvVjrYa, diNFRG8zC63 *NyBWtFnlv, uADzMh0kZ7hL uint32, r4wfXKbPIsE uintptr, rmDzrzri **SoXj8D) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ cGFKoL6.Addr(), 8,  /*line :1*/ uintptr(owi56xe),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(akEz9oY8Rh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eTypM8uhmGhb)),  /*line :1*/ uintptr(bamTPS0GgqZ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(diNFRG8zC63)),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr(r4wfXKbPIsE),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rmDzrzri)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func MAx7YVuDBz1(qSd2AI uintptr, aLwXrP uint32, b73VLu5 uintptr, uADzMh0kZ7hL uint32, diNFRG8zC63 uintptr) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ qaQASfoPiG0.Addr(), 5,  /*line :1*/ uintptr(qSd2AI),  /*line :1*/ uintptr(aLwXrP),  /*line :1*/ uintptr(b73VLu5),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr(diNFRG8zC63), 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func YBqkmFA(fRfDs3kP SRlvVjrYa, hzeU8Tl *uint16) (_lQfGCXzTo SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ rNfpvrMIbs7.Addr(), 2,  /*line :1*/ uintptr(fRfDs3kP),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)), 0)
	_lQfGCXzTo =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if _lQfGCXzTo == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func ReanL6uUQ7y(qsgwOS uintptr, icz4ra *SoXj8D, diNFRG8zC63 *XssyNphTC, bOuFGIOlicGF *UvCDcg6fLG) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ x23fWxIPr.Addr(), 4,  /*line :1*/ uintptr(qsgwOS),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(icz4ra)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(diNFRG8zC63)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bOuFGIOlicGF)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func XpXlD6aa(w4EWFSvFKvr *uint16, tMlaMQYu *uint16) (vvb2ibDimRH bool) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ dnzSZjpfq0.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(w4EWFSvFKvr)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tMlaMQYu)), 0)
	vvb2ibDimRH = kDaalUAu24K != 0
	return
}

func BE314j4hm(hzeU8Tl string, bp5GsMBX5z uint16, tTWqSvX1 uint32, k9pIJA75IO9o *byte, oxanlcO **OB0HzAtGUg, kqZ3BjLR *byte) (bOuFGIOlicGF error) {
	var vJKEmfpV *uint16
	vJKEmfpV, bOuFGIOlicGF =  /*line :1*/ GcOmFfsibES(hzeU8Tl)
	if bOuFGIOlicGF != nil {
		return
	}
	return  /*line :1*/ iU7tJ4cdC(vJKEmfpV, bp5GsMBX5z, tTWqSvX1, k9pIJA75IO9o, oxanlcO, kqZ3BjLR)
}

func iU7tJ4cdC(hzeU8Tl *uint16, bp5GsMBX5z uint16, tTWqSvX1 uint32, k9pIJA75IO9o *byte, oxanlcO **OB0HzAtGUg, kqZ3BjLR *byte) (bOuFGIOlicGF error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ hAo8t3aX.Addr(), 6,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr(bp5GsMBX5z),  /*line :1*/ uintptr(tTWqSvX1),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(k9pIJA75IO9o)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oxanlcO)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kqZ3BjLR)))
	if kDaalUAu24K != 0 {
		bOuFGIOlicGF =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func DuNM27(b9bfd3 *OB0HzAtGUg, raTJj2qkF uint32) {
	 /*line :1*/ OwKnfO5tgvm4( /*line :1*/ wPq8lA.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b9bfd3)),  /*line :1*/ uintptr(raTJj2qkF), 0)
	return
}

func YSaoaO(fiAhVNL *VNW4OcO_CoZ7, cQd3FBZUdo *uint32) (xsa0Bcrh error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ yuG7y66.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fiAhVNL)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cQd3FBZUdo)), 0)
	if kDaalUAu24K != 0 {
		xsa0Bcrh =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func WoIrx1c(kNm7KbMHaBbX *SBqHJrlvgiBJ) (xsa0Bcrh error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ bcgbKEiGt.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kNm7KbMHaBbX)), 0, 0)
	if kDaalUAu24K != 0 {
		xsa0Bcrh =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func U060VrUaSEl(wzPhJhIFoI SRlvVjrYa) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ okr2Sl9wBb.Addr(), 1,  /*line :1*/ uintptr(wzPhJhIFoI), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func QaAWdiVWBp(wzPhJhIFoI SRlvVjrYa, qITQAHdKhR *ZaNm5QSYC9) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ dC_GmCDgrR.Addr(), 2,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qITQAHdKhR)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func FhKJLgXjwG(fYz2KOsTDon SRlvVjrYa) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ hAjxMUv.Addr(), 1,  /*line :1*/ uintptr(fYz2KOsTDon), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func ZM4mayH(ro5NCL9 *uint16, zC0xxjnQnl *DDwuM6RpYja) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ rLNmZY.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ro5NCL9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zC0xxjnQnl)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func CgAw2ZTl(vrCif8JWfr SRlvVjrYa, zC0xxjnQnl *DDwuM6RpYja, v6GJi3B uint32, i6aKfrJOD uint32, kZ0tPv uint32, hzeU8Tl *uint16) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ dH2uHyBd.Addr(), 6,  /*line :1*/ uintptr(vrCif8JWfr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zC0xxjnQnl)),  /*line :1*/ uintptr(v6GJi3B),  /*line :1*/ uintptr(i6aKfrJOD),  /*line :1*/ uintptr(kZ0tPv),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)))
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func EyjYVpa(hzeU8Tl *uint16, wzWkRJs uint32, ediyaFOkNQmk uint32, zC0xxjnQnl *DDwuM6RpYja, ygINFYaFDGa8 uint32, ibvWO6_mblkb uint32, bY_Mn1mwz int32) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ cCvppXKP6W83.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr(wzWkRJs),  /*line :1*/ uintptr(ediyaFOkNQmk),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zC0xxjnQnl)),  /*line :1*/ uintptr(ygINFYaFDGa8),  /*line :1*/ uintptr(ibvWO6_mblkb),  /*line :1*/ uintptr(bY_Mn1mwz), 0, 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == InvalidHandle {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Zg0EjHXxfSC(fPECh7D6 *uint16, sVnnQMP *uint16, r4wfXKbPIsE uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ zus5CKQbyQ.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fPECh7D6)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sVnnQMP)),  /*line :1*/ uintptr(r4wfXKbPIsE))
	if iwAHyJGkWJy&0xff == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func nAuRSVpaHRa3(a8KPcHKEe SRlvVjrYa, j6_sksnZU SRlvVjrYa, nIB9uB1 uintptr, d5Ldozqc uint32) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ w6QQ05j.Addr(), 4,  /*line :1*/ uintptr(a8KPcHKEe),  /*line :1*/ uintptr(j6_sksnZU),  /*line :1*/ uintptr(nIB9uB1),  /*line :1*/ uintptr(d5Ldozqc), 0, 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func JtZhNTR6(kDpfNveH8Gt *SRlvVjrYa, vyce8N *SRlvVjrYa, zC0xxjnQnl *DDwuM6RpYja, jFhAAWnAX uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ hZFdxL.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kDpfNveH8Gt)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vyce8N)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zC0xxjnQnl)),  /*line :1*/ uintptr(jFhAAWnAX), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func ATXyyvx(fmYYRC *uint16, oAZypulaO *uint16, u7efeBj9 *DDwuM6RpYja, cdwLgBk *DDwuM6RpYja, h5AurTaAMtM1 bool, bKz7BnjpP5Aa uint32, dqnOhJ8vse *uint16, qmKEwwXG_ *uint16, tv24Vf0a9XQj *J7Nv3S0G31Q, akZSwzochq *NBPVXpO3Kv) (gOCcQzbcC error) {
	var vJKEmfpV uint32
	if h5AurTaAMtM1 {
		vJKEmfpV = 1
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ IiL5Io0Q5I( /*line :1*/ uGfRBah.Addr(), 10,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fmYYRC)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oAZypulaO)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u7efeBj9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cdwLgBk)),  /*line :1*/ uintptr(vJKEmfpV),  /*line :1*/ uintptr(bKz7BnjpP5Aa),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dqnOhJ8vse)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qmKEwwXG_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tv24Vf0a9XQj)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(akZSwzochq)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func HLmNaQslmU(fBGivJZb3lf *uint16, eKnbYvX *uint16, uADzMh0kZ7hL uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ soOK9W.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fBGivJZb3lf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eKnbYvX)),  /*line :1*/ uintptr(uADzMh0kZ7hL))
	if iwAHyJGkWJy&0xff == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func UsCsCLMhll(uADzMh0kZ7hL uint32, nDGAw2C uint32) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ cxDdWOFEip.Addr(), 2,  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr(nDGAw2C), 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == InvalidHandle {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func N_MYEJ7Bb(ro5NCL9 *uint16) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ xf2zr4bdg.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ro5NCL9)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func jbwtrz4(w8MAs38Uot *vYoIhduc) {
	 /*line :1*/ OwKnfO5tgvm4( /*line :1*/ gVXfdl_I8ne.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(w8MAs38Uot)), 0, 0)
	return
}

func EynTCZg(fYz2KOsTDon SRlvVjrYa, cu_NRaDt7KOw uint32, rNm1TEzdH *byte, lo3FwKNRKQ uint32, aDoSyaf2wEEX *byte, exqGxJEzwB uint32, dKjLqwpa *uint32, kR5OwbP *ZaNm5QSYC9) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ enRZYWOlh4.Addr(), 8,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(cu_NRaDt7KOw),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rNm1TEzdH)),  /*line :1*/ uintptr(lo3FwKNRKQ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aDoSyaf2wEEX)),  /*line :1*/ uintptr(exqGxJEzwB),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dKjLqwpa)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func VXs68_ZF(pDaslf SRlvVjrYa, _6NxLx_VCd SRlvVjrYa, jaLQanj0j SRlvVjrYa, unIBnL8Q *SRlvVjrYa, bTGPe_b uint32, mjVbH1av bool, vnbklHVww7 uint32) (gOCcQzbcC error) {
	var vJKEmfpV uint32
	if mjVbH1av {
		vJKEmfpV = 1
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ epbX1qoTStN.Addr(), 7,  /*line :1*/ uintptr(pDaslf),  /*line :1*/ uintptr(_6NxLx_VCd),  /*line :1*/ uintptr(jaLQanj0j),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(unIBnL8Q)),  /*line :1*/ uintptr(bTGPe_b),  /*line :1*/ uintptr(vJKEmfpV),  /*line :1*/ uintptr(vnbklHVww7), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func HfCDGbxP(dp1JYmM uint32) {
	 /*line :1*/ OwKnfO5tgvm4( /*line :1*/ nNMja9.Addr(), 1,  /*line :1*/ uintptr(dp1JYmM), 0, 0)
	return
}

func JaVhl9Mkgfp(fYz2KOsTDon SRlvVjrYa) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ sahN1d8mp.Addr(), 1,  /*line :1*/ uintptr(fYz2KOsTDon), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func deKyXFXaIT(hzeU8Tl *uint16, gv3kki *w8vwhA) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ cLocy4DjDO.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gv3kki)), 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == InvalidHandle {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func oXHZwN5(fYz2KOsTDon SRlvVjrYa, gv3kki *w8vwhA) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ c5whBLyge1K6.Addr(), 2,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gv3kki)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func IvXGh3(fYz2KOsTDon SRlvVjrYa) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ jYR1ar9Y.Addr(), 1,  /*line :1*/ uintptr(fYz2KOsTDon), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func F9mebY(ndm10JSX uintptr, cne_74 uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ badahuFho.Addr(), 2,  /*line :1*/ uintptr(ndm10JSX),  /*line :1*/ uintptr(cne_74), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func _MtwugS(uADzMh0kZ7hL uint32, mEtGX6Jm uintptr, fpKe4WGIx uint32, gLEjStHSQo uint32, eun1Jlud5A []uint16, dqAabTiua *byte) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	var vJKEmfpV *uint16
	if  /*line :1*/ len(eun1Jlud5A) > 0 {
		vJKEmfpV = &eun1Jlud5A[0]
	}
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ zcnyMlEw4JIX.Addr(), 7,  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr(mEtGX6Jm),  /*line :1*/ uintptr(fpKe4WGIx),  /*line :1*/ uintptr(gLEjStHSQo),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vJKEmfpV)),  /*line :1*/ uintptr( /*line :1*/ len(eun1Jlud5A)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dqAabTiua)), 0, 0)
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func CP4mgHYW5e8D(kMhDYkqWv2mI *uint16) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ uO8mgeTm3r.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kMhDYkqWv2mI)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func L8_GfNWaKJ(fYz2KOsTDon SRlvVjrYa) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ u4kiFTZ.Addr(), 1,  /*line :1*/ uintptr(fYz2KOsTDon), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func IaCOl3xN() (bVlKr7uHVgba *uint16) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ xgzVmmArEtci.Addr(), 0, 0, 0, 0)
	bVlKr7uHVgba = (* /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	return
}

func LwbTmPN(eun1Jlud5A *uint16, m5Tq_PL7 *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ g56k60.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m5Tq_PL7)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func HdjCvUIyL9B(cRLrlncAqN SRlvVjrYa, ediyaFOkNQmk *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ rwK2LeIhP5d.Addr(), 2,  /*line :1*/ uintptr(cRLrlncAqN),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ediyaFOkNQmk)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func WRA3NP30nq(izePNH7xS3 uint32, eun1Jlud5A *uint16) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ jJW5EcVc.Addr(), 2,  /*line :1*/ uintptr(izePNH7xS3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)), 0)
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func ZoL0wDe() (k5i2XPxMadw3 SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ aI65Xv.Addr(), 0, 0, 0, 0)
	k5i2XPxMadw3 =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if k5i2XPxMadw3 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func gEvRXHL77A() (aJF0So uint32) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ wuPQ4Gcnp57.Addr(), 0, 0, 0, 0)
	aJF0So =  /*line :1*/ uint32(kDaalUAu24K)
	return
}

func ZEq5mFM20k() (kMhDYkqWv2mI *uint16, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ dyW3aKwd5bn.Addr(), 0, 0, 0, 0)
	kMhDYkqWv2mI = (* /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	if kMhDYkqWv2mI == nil {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func UVmkZzea(hzeU8Tl *uint16, bUmqniD *uint16, jFhAAWnAX uint32) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ iQgSceA.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bUmqniD)),  /*line :1*/ uintptr(jFhAAWnAX))
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func YYxQKvX(fYz2KOsTDon SRlvVjrYa, dp1JYmM *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ f3tIzkTRW.Addr(), 2,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dp1JYmM)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Z8GRiq0cH(hzeU8Tl *uint16, jAzpOh uint32, b65sRubEGAyW *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ yVk9wi.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr(jAzpOh),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b65sRubEGAyW)))
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func BDVJtIREn(hzeU8Tl *uint16) (ibvWO6_mblkb uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ atp9nGcx.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)), 0, 0)
	ibvWO6_mblkb =  /*line :1*/ uint32(kDaalUAu24K)
	if ibvWO6_mblkb == INVALID_FILE_ATTRIBUTES {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func D1OiNlWWBa(fYz2KOsTDon SRlvVjrYa, gv3kki *HCYt4WU_Wmzb) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ jBHRrXeJ.Addr(), 2,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gv3kki)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func VnXJph2(a8KPcHKEe SRlvVjrYa) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ pNkaasL.Addr(), 1,  /*line :1*/ uintptr(a8KPcHKEe), 0, 0)
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func a7UqCnQe0(kfCaVeXmeDA SRlvVjrYa, ep1m_VHJpz2 *uint16, d7aG1dWLx uint32, uADzMh0kZ7hL uint32) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ f0xxVP.Addr(), 4,  /*line :1*/ uintptr(kfCaVeXmeDA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ep1m_VHJpz2)),  /*line :1*/ uintptr(d7aG1dWLx),  /*line :1*/ uintptr(uADzMh0kZ7hL), 0, 0)
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 || m5Tq_PL7 >= d7aG1dWLx {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func T0pUSJXE6Z(ro5NCL9 *uint16, izePNH7xS3 uint32, eun1Jlud5A *uint16, drMObaaRZ_ **uint16) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ t_DfADTW.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ro5NCL9)),  /*line :1*/ uintptr(izePNH7xS3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(drMObaaRZ_)), 0, 0)
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func GzCiuJ1CJ6() (arKETGLcO error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ wgfyKiKTY1Yo.Addr(), 0, 0, 0, 0)
	if kDaalUAu24K != 0 {
		arKETGLcO =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func B0zwZjcMlAiw(ro5NCL9 *uint16, eun1Jlud5A *uint16, izePNH7xS3 uint32) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ qzzLr7B.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ro5NCL9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr(izePNH7xS3))
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func No2FavKgeK(j92hK8Qc0RLB SRlvVjrYa, iF0pFakwl string) (zguyPht uintptr, gOCcQzbcC error) {
	var vJKEmfpV *byte
	vJKEmfpV, gOCcQzbcC =  /*line :1*/ Oea4iRaRU2r(iF0pFakwl)
	if gOCcQzbcC != nil {
		return
	}
	return  /*line :1*/ vBMNDoVgD_T7(j92hK8Qc0RLB, vJKEmfpV)
}

func vBMNDoVgD_T7(j92hK8Qc0RLB SRlvVjrYa, iF0pFakwl *byte) (zguyPht uintptr, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ bfdaj_wm9rA.Addr(), 2,  /*line :1*/ uintptr(j92hK8Qc0RLB),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iF0pFakwl)), 0)
	zguyPht =  /*line :1*/ uintptr(kDaalUAu24K)
	if zguyPht == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func AznBquFp(fYz2KOsTDon SRlvVjrYa, x1_qdQi *T8WbUqAC3v, aIa6PcmKL *T8WbUqAC3v, lQ9GZFjMXX *T8WbUqAC3v, mdDLmXq8Vs *T8WbUqAC3v) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ d1qXWjaARQ.Addr(), 5,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x1_qdQi)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aIa6PcmKL)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lQ9GZFjMXX)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mdDLmXq8Vs)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func cFzyNYNTK(j6_sksnZU SRlvVjrYa, lh_XEK *uint32, nIB9uB1 *uintptr, kR5OwbP **ZaNm5QSYC9, vYbNuNe uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ jDDTbwhjDY.Addr(), 5,  /*line :1*/ uintptr(j6_sksnZU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lh_XEK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nIB9uB1)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr(vYbNuNe), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func SNTS2pp(qhXeSuLtxhMK *uint16, kAdyP7Cu *uint16, izePNH7xS3 uint32) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ hyUD4tkB.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qhXeSuLtxhMK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kAdyP7Cu)),  /*line :1*/ uintptr(izePNH7xS3))
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func DtiNKM74hL(tv24Vf0a9XQj *J7Nv3S0G31Q) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ xga4aeva.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tv24Vf0a9XQj)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Kgv7ZgME9(zbjgDwolpB int) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ twemfk7Xg.Addr(), 1,  /*line :1*/ uintptr(zbjgDwolpB), 0, 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == InvalidHandle {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func ItrSBaJMcT(eTypM8uhmGhb *T8WbUqAC3v) {
	 /*line :1*/ OwKnfO5tgvm4( /*line :1*/ dQ2MstIcfyPZ.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eTypM8uhmGhb)), 0, 0)
	return
}

func Jr8XEH5r(izePNH7xS3 uint32, eun1Jlud5A *uint16) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ epfs28nOzvtO.Addr(), 2,  /*line :1*/ uintptr(izePNH7xS3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)), 0)
	m5Tq_PL7 =  /*line :1*/ uint32(kDaalUAu24K)
	if m5Tq_PL7 == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func BbXWjqCVp(fqTbBB57 *MAPVxQZiq) (aGlzqq0Y uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ rX6vZNDMu5dK.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fqTbBB57)), 0, 0)
	aGlzqq0Y =  /*line :1*/ uint32(kDaalUAu24K)
	if aGlzqq0Y == 0xffffffff {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func MTvUNch79atA() (bHZIL1KaqLQ uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ c0AVg4p8r.Addr(), 0, 0, 0, 0)
	bHZIL1KaqLQ =  /*line :1*/ uint32(kDaalUAu24K)
	if bHZIL1KaqLQ == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func eavEkDP(w8MAs38Uot *vYoIhduc, cpbK0I uint32, uADzMh0kZ7hL uint32, jFhAAWnAX *uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ gBsNaY6plc.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(w8MAs38Uot)),  /*line :1*/ uintptr(cpbK0I),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jFhAAWnAX)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func K51HyH2_(jA9VhEW347 string) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	var vJKEmfpV *uint16
	vJKEmfpV, gOCcQzbcC =  /*line :1*/ GcOmFfsibES(jA9VhEW347)
	if gOCcQzbcC != nil {
		return
	}
	return  /*line :1*/ t_KUNS(vJKEmfpV)
}

func t_KUNS(jA9VhEW347 *uint16) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ w7f9jzy.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jA9VhEW347)), 0, 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func AFNwp5z(l5PRYkce SRlvVjrYa) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ iJaCKplx.Addr(), 1,  /*line :1*/ uintptr(l5PRYkce), 0, 0)
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon != 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func SQWyxhx(fYz2KOsTDon SRlvVjrYa, wzWkRJs uint32, cJ2xhVctS uint32, fZ0xrP uint32, cne_74 uintptr) (ndm10JSX uintptr, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ e8Ujhp.Addr(), 5,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(wzWkRJs),  /*line :1*/ uintptr(cJ2xhVctS),  /*line :1*/ uintptr(fZ0xrP),  /*line :1*/ uintptr(cne_74), 0)
	ndm10JSX =  /*line :1*/ uintptr(kDaalUAu24K)
	if ndm10JSX == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Ib01EQM(d4xMeWOatV4j *uint16, c7wlCtHEZpy *uint16) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ vXQQlhZc6pyH.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d4xMeWOatV4j)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c7wlCtHEZpy)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func JkYb3Icqg(g4rSzi4z uint32, oncHBlfpJu7 bool, aJF0So uint32) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	var vJKEmfpV uint32
	if oncHBlfpJu7 {
		vJKEmfpV = 1
	}
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ hyGFQMbH.Addr(), 3,  /*line :1*/ uintptr(g4rSzi4z),  /*line :1*/ uintptr(vJKEmfpV),  /*line :1*/ uintptr(aJF0So))
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func jlfCPe(j6_sksnZU SRlvVjrYa, lh_XEK uint32, nIB9uB1 uintptr, kR5OwbP *ZaNm5QSYC9) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ yFqFoUe.Addr(), 4,  /*line :1*/ uintptr(j6_sksnZU),  /*line :1*/ uintptr(lh_XEK),  /*line :1*/ uintptr(nIB9uB1),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func WkbE7c(piKsTT2T SRlvVjrYa, mowY07B *GRU31Ht3FPn) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ q5daQHwuvcZ.Addr(), 2,  /*line :1*/ uintptr(piKsTT2T),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mowY07B)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func IhA67uCoej5(piKsTT2T SRlvVjrYa, mowY07B *GRU31Ht3FPn) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ sRLLKO4IwOQ.Addr(), 2,  /*line :1*/ uintptr(piKsTT2T),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mowY07B)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Jkhz6hQ6zGAR(cRLrlncAqN SRlvVjrYa, eun1Jlud5A *uint16, b0KG9AEBv uint32, hD0llwcR *uint32, dXKY2EFuy6M *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ fVbbzya.Addr(), 5,  /*line :1*/ uintptr(cRLrlncAqN),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr(b0KG9AEBv),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hD0llwcR)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dXKY2EFuy6M)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func FtXEnZ(fYz2KOsTDon SRlvVjrYa, eun1Jlud5A *byte, izePNH7xS3 uint32, ueXuVkOz1 bool, pcbZEtmEyH uint32, gVZNdKoW9Uwr *uint32, kR5OwbP *ZaNm5QSYC9, laEtYUWmcz uintptr) (gOCcQzbcC error) {
	var vJKEmfpV uint32
	if ueXuVkOz1 {
		vJKEmfpV = 1
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ qkBhr91ZH2.Addr(), 8,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr(izePNH7xS3),  /*line :1*/ uintptr(vJKEmfpV),  /*line :1*/ uintptr(pcbZEtmEyH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gVZNdKoW9Uwr)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr(laEtYUWmcz), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func qzDaRXT(fYz2KOsTDon SRlvVjrYa, eun1Jlud5A []byte, ohzXBFIQ *uint32, kR5OwbP *ZaNm5QSYC9) (gOCcQzbcC error) {
	var vJKEmfpV *byte
	if  /*line :1*/ len(eun1Jlud5A) > 0 {
		vJKEmfpV = &eun1Jlud5A[0]
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ dVi2JIQfx.Addr(), 5,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vJKEmfpV)),  /*line :1*/ uintptr( /*line :1*/ len(eun1Jlud5A)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohzXBFIQ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func PQFNcx(ro5NCL9 *uint16) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ a_camEYomia.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ro5NCL9)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func VqvfqI(ro5NCL9 *uint16) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ iDm22rip.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ro5NCL9)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func TYvDTZlPLYZ0(fYz2KOsTDon SRlvVjrYa) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ j6PdoeCJe5.Addr(), 1,  /*line :1*/ uintptr(fYz2KOsTDon), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func EkkQwxRH(hzeU8Tl *uint16, fsp_0a7i *uint16) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ fJAOac3a2Xi.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fsp_0a7i)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func D3KivaD(hzeU8Tl *uint16, ibvWO6_mblkb uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ xihm6IzgQA.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr(ibvWO6_mblkb), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func F2ye7PjCo5l(fYz2KOsTDon SRlvVjrYa, uADzMh0kZ7hL uint8) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ dAUCsfm.Addr(), 2,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(uADzMh0kZ7hL), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Hw81Kz715KV(fYz2KOsTDon SRlvVjrYa, sXxiEnR7K4 int32, hXZ74l1D *int32, ddcX8W uint32) (lt6bHOSlDybB uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ y6R4dSDuFII.Addr(), 4,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(sXxiEnR7K4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hXZ74l1D)),  /*line :1*/ uintptr(ddcX8W), 0, 0)
	lt6bHOSlDybB =  /*line :1*/ uint32(kDaalUAu24K)
	if lt6bHOSlDybB == 0xffffffff {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func CA7BjaTdOLL(fYz2KOsTDon SRlvVjrYa, iO4Cae *T8WbUqAC3v, h0Nh76HWu *T8WbUqAC3v, gNtx9STqM1J *T8WbUqAC3v) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ dn91DjN1bhX.Addr(), 4,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iO4Cae)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(h0Nh76HWu)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gNtx9STqM1J)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func LUuQqH65KsfT(fYz2KOsTDon SRlvVjrYa, pcbZEtmEyH uint32, uADzMh0kZ7hL uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ glaMy98Qs5.Addr(), 3,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(pcbZEtmEyH),  /*line :1*/ uintptr(uADzMh0kZ7hL))
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Vp9MpT5Y(fYz2KOsTDon SRlvVjrYa, dp1JYmM uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ jSMiME.Addr(), 2,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(dp1JYmM), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func KNTqdurWWHPD(ndm10JSX uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ yiAURBm17Lus.Addr(), 1,  /*line :1*/ uintptr(ndm10JSX), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func rq14e8x(w8MAs38Uot *vYoIhduc, uADzMh0kZ7hL uint32, c4UzX7Pz uintptr, fsp_0a7i unsafe.Pointer, jFhAAWnAX uintptr, aDRratjyBr unsafe.Pointer, vdP0F84wQ991 *uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ jihmTDgco.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(w8MAs38Uot)),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr(c4UzX7Pz),  /*line :1*/ uintptr(fsp_0a7i),  /*line :1*/ uintptr(jFhAAWnAX),  /*line :1*/ uintptr(aDRratjyBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vdP0F84wQ991)), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func G1w1ForDw(ndm10JSX uintptr, cne_74 uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ wNT4s8M.Addr(), 2,  /*line :1*/ uintptr(ndm10JSX),  /*line :1*/ uintptr(cne_74), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func IlrO4hSDxT(ndm10JSX uintptr, cne_74 uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ bWCcw5aOP3.Addr(), 2,  /*line :1*/ uintptr(ndm10JSX),  /*line :1*/ uintptr(cne_74), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func N_8burTVa(fYz2KOsTDon SRlvVjrYa, nVFcJK uint32) (i55lUlUIkpuf uint32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ xQGMZ_3dF.Addr(), 2,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(nVFcJK), 0)
	i55lUlUIkpuf =  /*line :1*/ uint32(kDaalUAu24K)
	if i55lUlUIkpuf == 0xffffffff {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Djt4smg(cRLrlncAqN SRlvVjrYa, eun1Jlud5A *uint16, hL730zIIaaTZ uint32, ywelE2J *uint32, r4wfXKbPIsE *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ tl81Ww_JQ.Addr(), 5,  /*line :1*/ uintptr(cRLrlncAqN),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr(hL730zIIaaTZ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ywelE2J)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(r4wfXKbPIsE)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func b5BOiDlX(fYz2KOsTDon SRlvVjrYa, eun1Jlud5A []byte, ohzXBFIQ *uint32, kR5OwbP *ZaNm5QSYC9) (gOCcQzbcC error) {
	var vJKEmfpV *byte
	if  /*line :1*/ len(eun1Jlud5A) > 0 {
		vJKEmfpV = &eun1Jlud5A[0]
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ h6gtSIpLuhao.Addr(), 5,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vJKEmfpV)),  /*line :1*/ uintptr( /*line :1*/ len(eun1Jlud5A)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohzXBFIQ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func VCi2jNl(dKzZXg SRlvVjrYa, jwDa5JpX9emi SRlvVjrYa, eun1Jlud5A *byte, r_qJ6g17g uint32, d39hhv_mJGP uint32, wqozthmjpBU uint32, iG0caX *uint32, kR5OwbP *ZaNm5QSYC9) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ nDAteP7QBXTA.Addr(), 8,  /*line :1*/ uintptr(dKzZXg),  /*line :1*/ uintptr(jwDa5JpX9emi),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr(r_qJ6g17g),  /*line :1*/ uintptr(d39hhv_mJGP),  /*line :1*/ uintptr(wqozthmjpBU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iG0caX)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)), 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func NRaqTZ(eun1Jlud5A *byte, r_qJ6g17g uint32, d39hhv_mJGP uint32, wqozthmjpBU uint32, tDV1QrEsML **IXy5oynSaLQM, yMp0PajTA *int32, gQgIshUa **IXy5oynSaLQM, jp2Q9m *int32) {
	 /*line :1*/ Rc0O05UsV( /*line :1*/ eNwean.Addr(), 8,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)),  /*line :1*/ uintptr(r_qJ6g17g),  /*line :1*/ uintptr(d39hhv_mJGP),  /*line :1*/ uintptr(wqozthmjpBU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tDV1QrEsML)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yMp0PajTA)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gQgIshUa)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jp2Q9m)), 0)
	return
}

func PFrj34i0s(wzPhJhIFoI SRlvVjrYa, fYz2KOsTDon SRlvVjrYa, b8GZr5rxWbr7 uint32, wpd85vC uint32, kR5OwbP *ZaNm5QSYC9, zHjLpYPHV *GMHaGunA, uADzMh0kZ7hL uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ datHkUf_q.Addr(), 7,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(b8GZr5rxWbr7),  /*line :1*/ uintptr(wpd85vC),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zHjLpYPHV)),  /*line :1*/ uintptr(uADzMh0kZ7hL), 0, 0)
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func IgpeqHzNV1_D(eun1Jlud5A *byte) (g1czeUaV error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ lLxcq4xc.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)), 0, 0)
	if kDaalUAu24K != 0 {
		g1czeUaV =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func HAGAx7(sSdkTzj_and *uint16, hzeU8Tl **uint16, ix8TNU *uint32) (g1czeUaV error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ celUNpxNnT.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sSdkTzj_and)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ix8TNU)))
	if kDaalUAu24K != 0 {
		g1czeUaV =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func BmQbh8h(k5nOMEd *uint16, ns33NRifpD *uint16, jAzpOh uint32, eun1Jlud5A **byte) (g1czeUaV error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ oi9C38I3oW.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(k5nOMEd)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ns33NRifpD)),  /*line :1*/ uintptr(jAzpOh),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eun1Jlud5A)), 0, 0)
	if kDaalUAu24K != 0 {
		g1czeUaV =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func kjlGhacz(pCAJ0g2o *uint32, iFfYmN *uint32, t17Ejt0dg5 *uint32) {
	 /*line :1*/ OwKnfO5tgvm4( /*line :1*/ fEJdp1o.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pCAJ0g2o)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iFfYmN)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t17Ejt0dg5)))
	return
}

func FbYqdfML3d(b7ktF_QK uint32, jgObQxfRb *uint16, acEzQBms *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ jaJ9KC.Addr(), 3,  /*line :1*/ uintptr(b7ktF_QK),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jgObQxfRb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(acEzQBms)))
	if iwAHyJGkWJy&0xff == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Zk7h0UIFNHVK(c5MDVqhhIe *uint16, orMxeBRF9wG uint32, v344ix9EmX uint32, a7pwSTgb *uint16, acEzQBms *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ iLxgMTvV.Addr(), 5,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c5MDVqhhIe)),  /*line :1*/ uintptr(orMxeBRF9wG),  /*line :1*/ uintptr(v344ix9EmX),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7pwSTgb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(acEzQBms)), 0)
	if iwAHyJGkWJy&0xff == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func W9amGeFo(bVlKr7uHVgba *uint16, t2JOScepbE *int32) (eLPqUuZoy *[8192]*[8192]uint16, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ p3khPPCvCgiS.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bVlKr7uHVgba)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t2JOScepbE)), 0)
	eLPqUuZoy = (*[8192]*[8192] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	if eLPqUuZoy == nil {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func AlC59xuX(mhNoXZ Ad4bWd, hY26vae8L *uint16, beHsJqOwdR *uint32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ iGKVJSlcS.Addr(), 3,  /*line :1*/ uintptr(mhNoXZ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hY26vae8L)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(beHsJqOwdR)))
	if iwAHyJGkWJy == 0 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func BmTZjeIGcn(liNEWExz *HnpJdSpZM__0) {
	 /*line :1*/ OwKnfO5tgvm4( /*line :1*/ zHb9Zuab.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(liNEWExz)), 0, 0)
	return
}

func Nsyrt6(dLcHPC86U *uint16, bCStjziXtJw *uint16, pRgnfqPf *HnpJdSpZM__0, j2GWlZBnk **HnpJdSpZM__0) (lND4H_r error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ x25ZA6ZEaO4.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dLcHPC86U)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bCStjziXtJw)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pRgnfqPf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j2GWlZBnk)), 0, 0)
	if kDaalUAu24K != 0 {
		lND4H_r =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func VYCkASwoW() (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ zT5cP8.Addr(), 0, 0, 0, 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func AMZXrhL3Tt8(mqQEcmN1 *int32, nZRZeCs3B *ZEPlE2KjP71, rY8nYlkA *uint32) (m5Tq_PL7 int32, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ ijij9sQGyfD.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mqQEcmN1)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nZRZeCs3B)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rY8nYlkA)))
	m5Tq_PL7 =  /*line :1*/ int32(kDaalUAu24K)
	if m5Tq_PL7 == -1 {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func YvCdY02(wzPhJhIFoI SRlvVjrYa, iaKJsfNh4 uint32, oYMJftb *byte, zAiyjqZ uint32, aMmvVcAqO *byte, aTXGHh3EI uint32, zCN5aoN *uint32, kR5OwbP *ZaNm5QSYC9, laEtYUWmcz uintptr) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ aAmayrk2oVz.Addr(), 9,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(iaKJsfNh4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oYMJftb)),  /*line :1*/ uintptr(zAiyjqZ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aMmvVcAqO)),  /*line :1*/ uintptr(aTXGHh3EI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zCN5aoN)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr(laEtYUWmcz))
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func J67wPAoai1Kb(wzPhJhIFoI SRlvVjrYa, m8dz2Ors *QbWhAp5CqTG, t8xduL uint32, iG0caX *uint32, uADzMh0kZ7hL *uint32, kR5OwbP *ZaNm5QSYC9, bpTsK_7YjdlT *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ nWTVn6cies.Addr(), 7,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m8dz2Ors)),  /*line :1*/ uintptr(t8xduL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iG0caX)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uADzMh0kZ7hL)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bpTsK_7YjdlT)), 0, 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func AYS5UpcaP(wzPhJhIFoI SRlvVjrYa, m8dz2Ors *QbWhAp5CqTG, t8xduL uint32, iG0caX *uint32, uADzMh0kZ7hL *uint32, d4xMeWOatV4j *IXy5oynSaLQM, miPApaWKcl *int32, kR5OwbP *ZaNm5QSYC9, bpTsK_7YjdlT *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ eI4_1asU.Addr(), 9,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m8dz2Ors)),  /*line :1*/ uintptr(t8xduL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iG0caX)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uADzMh0kZ7hL)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d4xMeWOatV4j)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(miPApaWKcl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bpTsK_7YjdlT)))
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Ky9hDuJe(wzPhJhIFoI SRlvVjrYa, m8dz2Ors *QbWhAp5CqTG, t8xduL uint32, lu937Dl_88 *uint32, uADzMh0kZ7hL uint32, kR5OwbP *ZaNm5QSYC9, bpTsK_7YjdlT *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ eFrrx0Kpf.Addr(), 7,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m8dz2Ors)),  /*line :1*/ uintptr(t8xduL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lu937Dl_88)),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bpTsK_7YjdlT)), 0, 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func D_R3GW(wzPhJhIFoI SRlvVjrYa, m8dz2Ors *QbWhAp5CqTG, t8xduL uint32, lu937Dl_88 *uint32, uADzMh0kZ7hL uint32, c7wlCtHEZpy *IXy5oynSaLQM, tba65LWnk1Yt int32, kR5OwbP *ZaNm5QSYC9, bpTsK_7YjdlT *byte) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ epTJJ70Fc.Addr(), 9,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m8dz2Ors)),  /*line :1*/ uintptr(t8xduL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lu937Dl_88)),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c7wlCtHEZpy)),  /*line :1*/ uintptr(tba65LWnk1Yt),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bpTsK_7YjdlT)))
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func Kq4jyT(qjs8wrT7 uint32, gv3kki *RqD0UlyP) (lND4H_r error) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ dDWkr2Ttm.Addr(), 2,  /*line :1*/ uintptr(qjs8wrT7),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gv3kki)), 0)
	if kDaalUAu24K != 0 {
		lND4H_r =  /*line :1*/ O9Mga3(kDaalUAu24K)
	}
	return
}

func tKqRrokO7GR(wzPhJhIFoI SRlvVjrYa, hzeU8Tl unsafe.Pointer, eixaYPaXLSYp int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ f42Wfbh2XM.Addr(), 3,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(hzeU8Tl),  /*line :1*/ uintptr(eixaYPaXLSYp))
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func TBz9CFcaHh0(wzPhJhIFoI SRlvVjrYa) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ hTUat1z8gv2.Addr(), 1,  /*line :1*/ uintptr(wzPhJhIFoI), 0, 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func poAg04RDL(wzPhJhIFoI SRlvVjrYa, hzeU8Tl unsafe.Pointer, eixaYPaXLSYp int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ jahCtX.Addr(), 3,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(hzeU8Tl),  /*line :1*/ uintptr(eixaYPaXLSYp))
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func J0MCA2iRu3OD(hzeU8Tl string) (vFaWuUyY *MpOyP3z, gOCcQzbcC error) {
	var vJKEmfpV *byte
	vJKEmfpV, gOCcQzbcC =  /*line :1*/ Oea4iRaRU2r(hzeU8Tl)
	if gOCcQzbcC != nil {
		return
	}
	return  /*line :1*/ zeHao4WQ(vJKEmfpV)
}

func zeHao4WQ(hzeU8Tl *byte) (vFaWuUyY *MpOyP3z, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ lX4IVJBYT.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)), 0, 0)
	vFaWuUyY = (* /*line :1*/ MpOyP3z)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	if vFaWuUyY == nil {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func r3XXDB(wzPhJhIFoI SRlvVjrYa, nmgZCOdj *IXy5oynSaLQM, o7y8u5a9 *int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ anyxIG.Addr(), 3,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nmgZCOdj)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(o7y8u5a9)))
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func IO3AcKmy7(hzeU8Tl string) (rd6leevODT *DT8kuezeJ, gOCcQzbcC error) {
	var vJKEmfpV *byte
	vJKEmfpV, gOCcQzbcC =  /*line :1*/ Oea4iRaRU2r(hzeU8Tl)
	if gOCcQzbcC != nil {
		return
	}
	return  /*line :1*/ iNN3xn1X(vJKEmfpV)
}

func iNN3xn1X(hzeU8Tl *byte) (rd6leevODT *DT8kuezeJ, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ qGXtGudaEjdJ.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)), 0, 0)
	rd6leevODT = (* /*line :1*/ DT8kuezeJ)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	if rd6leevODT == nil {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func BRZyUKwdZ(hzeU8Tl string, qUFjymZN6q7 string) (wzPhJhIFoI *EEonh5, gOCcQzbcC error) {
	var vJKEmfpV *byte
	vJKEmfpV, gOCcQzbcC =  /*line :1*/ Oea4iRaRU2r(hzeU8Tl)
	if gOCcQzbcC != nil {
		return
	}
	var woWLOGrtTWFD *byte
	woWLOGrtTWFD, gOCcQzbcC =  /*line :1*/ Oea4iRaRU2r(qUFjymZN6q7)
	if gOCcQzbcC != nil {
		return
	}
	return  /*line :1*/ zVK38f(vJKEmfpV, woWLOGrtTWFD)
}

func zVK38f(hzeU8Tl *byte, qUFjymZN6q7 *byte) (wzPhJhIFoI *EEonh5, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ oHmw41Wc2g.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hzeU8Tl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qUFjymZN6q7)), 0)
	wzPhJhIFoI = (* /*line :1*/ EEonh5)( /*line :1*/ unsafe.Pointer(kDaalUAu24K))
	if wzPhJhIFoI == nil {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func aB6zPKR3lN(wzPhJhIFoI SRlvVjrYa, nmgZCOdj *IXy5oynSaLQM, o7y8u5a9 *int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ prLAS7DXZwiX.Addr(), 3,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nmgZCOdj)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(o7y8u5a9)))
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func N6ULYnxNR(wzPhJhIFoI SRlvVjrYa, jAzpOh int32, vCK0JqwoNbp int32, w5lLxQVpgJ9x *byte, a1qDBh2S6sde *int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ wowigi1g.Addr(), 5,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(jAzpOh),  /*line :1*/ uintptr(vCK0JqwoNbp),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(w5lLxQVpgJ9x)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a1qDBh2S6sde)), 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func nLqRlHL2wYJI(wzPhJhIFoI SRlvVjrYa, qrG5zKB int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ m_csHJ.Addr(), 2,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(qrG5zKB), 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func EOoAMh(op8hEkurD4Z uint16) (a5cTwa_MAAU uint16) {
	kDaalUAu24K, _, _ :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ j7CoLsioqaPP.Addr(), 1,  /*line :1*/ uintptr(op8hEkurD4Z), 0, 0)
	a5cTwa_MAAU =  /*line :1*/ uint16(kDaalUAu24K)
	return
}

func JwtaZG7GaLGD(wzPhJhIFoI SRlvVjrYa, jAzpOh int32, vCK0JqwoNbp int32, w5lLxQVpgJ9x *byte, a1qDBh2S6sde int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ VeAKF8sAwft( /*line :1*/ cwZfNC_Xk.Addr(), 5,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(jAzpOh),  /*line :1*/ uintptr(vCK0JqwoNbp),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(w5lLxQVpgJ9x)),  /*line :1*/ uintptr(a1qDBh2S6sde), 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func hIfg4o6v(wzPhJhIFoI SRlvVjrYa, tLIUarFXW4S int32) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ ebH4vF2kh.Addr(), 2,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(tLIUarFXW4S), 0)
	if iwAHyJGkWJy == socket_error {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}

func wA8pjx79d(si3u7lR_C int32, b85up2dBCm int32, t4Q5acX int32) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	kDaalUAu24K, _, mecVf9f6lPG :=  /*line :1*/ OwKnfO5tgvm4( /*line :1*/ gkkRqSs.Addr(), 3,  /*line :1*/ uintptr(si3u7lR_C),  /*line :1*/ uintptr(b85up2dBCm),  /*line :1*/ uintptr(t4Q5acX))
	fYz2KOsTDon =  /*line :1*/ SRlvVjrYa(kDaalUAu24K)
	if fYz2KOsTDon == InvalidHandle {
		gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return
}
