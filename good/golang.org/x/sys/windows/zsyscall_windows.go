//line :1
package NJ4MerJ

import (
	syscall "bUKeamF"
	"unsafe"
)

var _ unsafe.Pointer



const (
	errnoERROR_IO_PENDING = 997
)

var (
	ibFCohQg4b	error	=  /*line :1*/ syscall.O9Mga3(errnoERROR_IO_PENDING)
	fhP2386ivK0	error	= syscall.EINVAL
)



func rxn0etMCbt1(aYtmqRVkqc_8 syscall.O9Mga3) error {
	switch aYtmqRVkqc_8 {
	case 0:
		return fhP2386ivK0
	case errnoERROR_IO_PENDING:
		return ibFCohQg4b
	}

	return aYtmqRVkqc_8
}

var (
	w3DmnWERN	=  /*line :1*/ FDchPxUZ("CfgMgr32.dll")
	m1F5hvJa	=  /*line :1*/ FDchPxUZ("advapi32.dll")
	ztsLbseg	=  /*line :1*/ FDchPxUZ("crypt32.dll")
	eXI3Ca0p	=  /*line :1*/ FDchPxUZ("dnsapi.dll")
	coRoRaDrZR	=  /*line :1*/ FDchPxUZ("dwmapi.dll")
	pkrTSa	=  /*line :1*/ FDchPxUZ("iphlpapi.dll")
	fqaBrBSogv	=  /*line :1*/ FDchPxUZ("kernel32.dll")
	acDn5jcmJa	=  /*line :1*/ FDchPxUZ("mswsock.dll")
	wlbPK3P6wm	=  /*line :1*/ FDchPxUZ("netapi32.dll")
	e6pKCXfI	=  /*line :1*/ FDchPxUZ("ntdll.dll")
	mopZuQmjBn1	=  /*line :1*/ FDchPxUZ("ole32.dll")
	b9oGlF_vq7w	=  /*line :1*/ FDchPxUZ("psapi.dll")
	uq_VNn	=  /*line :1*/ FDchPxUZ("sechost.dll")
	th0X7dYaiGS	=  /*line :1*/ FDchPxUZ("secur32.dll")
	mz5Rgx	=  /*line :1*/ FDchPxUZ("setupapi.dll")
	eHlQWja3W	=  /*line :1*/ FDchPxUZ("shell32.dll")
	tBzL7HWDo	=  /*line :1*/ FDchPxUZ("user32.dll")
	cah2XcMwlQ	=  /*line :1*/ FDchPxUZ("userenv.dll")
	qLIhTl60e	=  /*line :1*/ FDchPxUZ("version.dll")
	z9p6iU	=  /*line :1*/ FDchPxUZ("winmm.dll")
	lZlB2nMZVPQ	=  /*line :1*/ FDchPxUZ("wintrust.dll")
	lHby3h	=  /*line :1*/ FDchPxUZ("ws2_32.dll")
	dpK9M7EuXT	=  /*line :1*/ FDchPxUZ("wtsapi32.dll")

	tFFQbdwDhh	=  /*line :1*/ w3DmnWERN.NewProc("CM_Get_DevNode_Status")
	aVYLuvsm57	=  /*line :1*/ w3DmnWERN.NewProc("CM_Get_Device_Interface_ListW")
	fGT4J_oCW	=  /*line :1*/ w3DmnWERN.NewProc("CM_Get_Device_Interface_List_SizeW")
	hiQsfw2rroCK	=  /*line :1*/ w3DmnWERN.NewProc("CM_MapCrToWin32Err")
	lBxQsVT	=  /*line :1*/ m1F5hvJa.NewProc("AdjustTokenGroups")
	rJZwiiMaGShr	=  /*line :1*/ m1F5hvJa.NewProc("AdjustTokenPrivileges")
	e0cwIqUmW	=  /*line :1*/ m1F5hvJa.NewProc("AllocateAndInitializeSid")
	nAnyEFS	=  /*line :1*/ m1F5hvJa.NewProc("BuildSecurityDescriptorW")
	uDrVRL4YzAV	=  /*line :1*/ m1F5hvJa.NewProc("ChangeServiceConfig2W")
	eDXu9g	=  /*line :1*/ m1F5hvJa.NewProc("ChangeServiceConfigW")
	yuRJn0ct2cO	=  /*line :1*/ m1F5hvJa.NewProc("CheckTokenMembership")
	eEfpnAv	=  /*line :1*/ m1F5hvJa.NewProc("CloseServiceHandle")
	zFSqKoaZXrU	=  /*line :1*/ m1F5hvJa.NewProc("ControlService")
	jOjTENA1PnO	=  /*line :1*/ m1F5hvJa.NewProc("ConvertSecurityDescriptorToStringSecurityDescriptorW")
	xatbdJZ1qGu	=  /*line :1*/ m1F5hvJa.NewProc("ConvertSidToStringSidW")
	v5QSJxw	=  /*line :1*/ m1F5hvJa.NewProc("ConvertStringSecurityDescriptorToSecurityDescriptorW")
	ezdAuq	=  /*line :1*/ m1F5hvJa.NewProc("ConvertStringSidToSidW")
	bx0KI1oy3MH	=  /*line :1*/ m1F5hvJa.NewProc("CopySid")
	ewtTalzq8K	=  /*line :1*/ m1F5hvJa.NewProc("CreateProcessAsUserW")
	fmRXc6Sw	=  /*line :1*/ m1F5hvJa.NewProc("CreateServiceW")
	ikJM5yaow9MS	=  /*line :1*/ m1F5hvJa.NewProc("CreateWellKnownSid")
	pg1hjm2rmuwg	=  /*line :1*/ m1F5hvJa.NewProc("CryptAcquireContextW")
	c8nLOIMJIsO	=  /*line :1*/ m1F5hvJa.NewProc("CryptGenRandom")
	cql2kwk	=  /*line :1*/ m1F5hvJa.NewProc("CryptReleaseContext")
	djEhYE	=  /*line :1*/ m1F5hvJa.NewProc("DeleteService")
	d0ecYKvv4	=  /*line :1*/ m1F5hvJa.NewProc("DeregisterEventSource")
	c_NdFS2aR	=  /*line :1*/ m1F5hvJa.NewProc("DuplicateTokenEx")
	jVle5HujZhu	=  /*line :1*/ m1F5hvJa.NewProc("EnumDependentServicesW")
	p_nAnOD	=  /*line :1*/ m1F5hvJa.NewProc("EnumServicesStatusExW")
	eeezyOsg	=  /*line :1*/ m1F5hvJa.NewProc("EqualSid")
	qjRLR1u_A	=  /*line :1*/ m1F5hvJa.NewProc("FreeSid")
	keIJ6j	=  /*line :1*/ m1F5hvJa.NewProc("GetLengthSid")
	dxLHi1t	=  /*line :1*/ m1F5hvJa.NewProc("GetNamedSecurityInfoW")
	lF7Vo3QbA	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityDescriptorControl")
	oqQtKDKgUS	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityDescriptorDacl")
	hGsGrxrHtn	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityDescriptorGroup")
	n0HbqHRwUwQA	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityDescriptorLength")
	vMRHziXX	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityDescriptorOwner")
	bj_QgkJG	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityDescriptorRMControl")
	_bwYz1	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityDescriptorSacl")
	vW6cF2tJp	=  /*line :1*/ m1F5hvJa.NewProc("GetSecurityInfo")
	avHMQVxXM	=  /*line :1*/ m1F5hvJa.NewProc("GetSidIdentifierAuthority")
	gHpdU2	=  /*line :1*/ m1F5hvJa.NewProc("GetSidSubAuthority")
	oaox2QLIw	=  /*line :1*/ m1F5hvJa.NewProc("GetSidSubAuthorityCount")
	uBqHfd9	=  /*line :1*/ m1F5hvJa.NewProc("GetTokenInformation")
	iYL04eOVwUN	=  /*line :1*/ m1F5hvJa.NewProc("ImpersonateSelf")
	rLUH8hPnmpv	=  /*line :1*/ m1F5hvJa.NewProc("InitializeSecurityDescriptor")
	gAmE9Sf	=  /*line :1*/ m1F5hvJa.NewProc("InitiateSystemShutdownExW")
	vrD_AZf88ah	=  /*line :1*/ m1F5hvJa.NewProc("IsTokenRestricted")
	ugeWaHsC4s	=  /*line :1*/ m1F5hvJa.NewProc("IsValidSecurityDescriptor")
	vAMA2aKZy	=  /*line :1*/ m1F5hvJa.NewProc("IsValidSid")
	aNFHkZYj1a	=  /*line :1*/ m1F5hvJa.NewProc("IsWellKnownSid")
	dm_EfGa	=  /*line :1*/ m1F5hvJa.NewProc("LookupAccountNameW")
	s2XfInzyjf1H	=  /*line :1*/ m1F5hvJa.NewProc("LookupAccountSidW")
	hP4bVgoWV	=  /*line :1*/ m1F5hvJa.NewProc("LookupPrivilegeValueW")
	go70DP	=  /*line :1*/ m1F5hvJa.NewProc("MakeAbsoluteSD")
	xK0citfiso_u	=  /*line :1*/ m1F5hvJa.NewProc("MakeSelfRelativeSD")
	gGbVEF6B	=  /*line :1*/ m1F5hvJa.NewProc("NotifyServiceStatusChangeW")
	tHa02a1GeCW	=  /*line :1*/ m1F5hvJa.NewProc("OpenProcessToken")
	jKB3JzepJ	=  /*line :1*/ m1F5hvJa.NewProc("OpenSCManagerW")
	kv0g00muN	=  /*line :1*/ m1F5hvJa.NewProc("OpenServiceW")
	j41UBMeaVV	=  /*line :1*/ m1F5hvJa.NewProc("OpenThreadToken")
	dDTNWYBO	=  /*line :1*/ m1F5hvJa.NewProc("QueryServiceConfig2W")
	etHz3XK	=  /*line :1*/ m1F5hvJa.NewProc("QueryServiceConfigW")
	zJFCJzhIRSx	=  /*line :1*/ m1F5hvJa.NewProc("QueryServiceDynamicInformation")
	kGmliiXuyeu	=  /*line :1*/ m1F5hvJa.NewProc("QueryServiceLockStatusW")
	pb2XaHC	=  /*line :1*/ m1F5hvJa.NewProc("QueryServiceStatus")
	gFDHk5H	=  /*line :1*/ m1F5hvJa.NewProc("QueryServiceStatusEx")
	e90awnRjK9	=  /*line :1*/ m1F5hvJa.NewProc("RegCloseKey")
	bQ2WMBaH	=  /*line :1*/ m1F5hvJa.NewProc("RegEnumKeyExW")
	cb9WQ0wtxY	=  /*line :1*/ m1F5hvJa.NewProc("RegNotifyChangeKeyValue")
	jlYF3eJWaPK	=  /*line :1*/ m1F5hvJa.NewProc("RegOpenKeyExW")
	vpgTqGRBg	=  /*line :1*/ m1F5hvJa.NewProc("RegQueryInfoKeyW")
	mh2e13t8abh	=  /*line :1*/ m1F5hvJa.NewProc("RegQueryValueExW")
	vowp8C8	=  /*line :1*/ m1F5hvJa.NewProc("RegisterEventSourceW")
	iD7x4md	=  /*line :1*/ m1F5hvJa.NewProc("RegisterServiceCtrlHandlerExW")
	aZaUnkxaI	=  /*line :1*/ m1F5hvJa.NewProc("ReportEventW")
	qUzKKN_	=  /*line :1*/ m1F5hvJa.NewProc("RevertToSelf")
	lHfFxw2FC	=  /*line :1*/ m1F5hvJa.NewProc("SetEntriesInAclW")
	i9on9y	=  /*line :1*/ m1F5hvJa.NewProc("SetKernelObjectSecurity")
	cmJ0sPauhG9q	=  /*line :1*/ m1F5hvJa.NewProc("SetNamedSecurityInfoW")
	orWzIs9FB	=  /*line :1*/ m1F5hvJa.NewProc("SetSecurityDescriptorControl")
	ajPxiSSw0tnI	=  /*line :1*/ m1F5hvJa.NewProc("SetSecurityDescriptorDacl")
	mxY4SkMLjpsU	=  /*line :1*/ m1F5hvJa.NewProc("SetSecurityDescriptorGroup")
	yBEYlnXato35	=  /*line :1*/ m1F5hvJa.NewProc("SetSecurityDescriptorOwner")
	arjuCqav	=  /*line :1*/ m1F5hvJa.NewProc("SetSecurityDescriptorRMControl")
	wdnaaUk	=  /*line :1*/ m1F5hvJa.NewProc("SetSecurityDescriptorSacl")
	zXhtiaLhV	=  /*line :1*/ m1F5hvJa.NewProc("SetSecurityInfo")
	woc2sm4	=  /*line :1*/ m1F5hvJa.NewProc("SetServiceStatus")
	hU0uN4LpB0d	=  /*line :1*/ m1F5hvJa.NewProc("SetThreadToken")
	g0O7MBn	=  /*line :1*/ m1F5hvJa.NewProc("SetTokenInformation")
	uGtcfa2YGM	=  /*line :1*/ m1F5hvJa.NewProc("StartServiceCtrlDispatcherW")
	uqaI0Vx	=  /*line :1*/ m1F5hvJa.NewProc("StartServiceW")
	pVLZEgaaqCvT	=  /*line :1*/ ztsLbseg.NewProc("CertAddCertificateContextToStore")
	f1WG2ds	=  /*line :1*/ ztsLbseg.NewProc("CertCloseStore")
	iZ6BTc2715	=  /*line :1*/ ztsLbseg.NewProc("CertCreateCertificateContext")
	uefQ0653B7Q	=  /*line :1*/ ztsLbseg.NewProc("CertDeleteCertificateFromStore")
	jaSFAnO	=  /*line :1*/ ztsLbseg.NewProc("CertDuplicateCertificateContext")
	elioVh8	=  /*line :1*/ ztsLbseg.NewProc("CertEnumCertificatesInStore")
	cwp59dU_C4DP	=  /*line :1*/ ztsLbseg.NewProc("CertFindCertificateInStore")
	zofSVx1U	=  /*line :1*/ ztsLbseg.NewProc("CertFindChainInStore")
	ofyKNlEiNQ	=  /*line :1*/ ztsLbseg.NewProc("CertFindExtension")
	es9AG6kX5eUb	=  /*line :1*/ ztsLbseg.NewProc("CertFreeCertificateChain")
	hR_xTjndo1Op	=  /*line :1*/ ztsLbseg.NewProc("CertFreeCertificateContext")
	zz4LZpa	=  /*line :1*/ ztsLbseg.NewProc("CertGetCertificateChain")
	dqk1LIahL	=  /*line :1*/ ztsLbseg.NewProc("CertGetNameStringW")
	aCKPLJWbT	=  /*line :1*/ ztsLbseg.NewProc("CertOpenStore")
	xERaMMGOivl	=  /*line :1*/ ztsLbseg.NewProc("CertOpenSystemStoreW")
	r4lpHHGKD	=  /*line :1*/ ztsLbseg.NewProc("CertVerifyCertificateChainPolicy")
	pMCa3SCIoqp	=  /*line :1*/ ztsLbseg.NewProc("CryptAcquireCertificatePrivateKey")
	g29_mRR6	=  /*line :1*/ ztsLbseg.NewProc("CryptDecodeObject")
	gxkyPGnJiO	=  /*line :1*/ ztsLbseg.NewProc("CryptProtectData")
	dupqPJx	=  /*line :1*/ ztsLbseg.NewProc("CryptQueryObject")
	zCYV5nU	=  /*line :1*/ ztsLbseg.NewProc("CryptUnprotectData")
	j6asE6	=  /*line :1*/ ztsLbseg.NewProc("PFXImportCertStore")
	snRiy6niCa6u	=  /*line :1*/ eXI3Ca0p.NewProc("DnsNameCompare_W")
	vOLc6_	=  /*line :1*/ eXI3Ca0p.NewProc("DnsQuery_W")
	j7fpugNwoH	=  /*line :1*/ eXI3Ca0p.NewProc("DnsRecordListFree")
	deA1m_XBF	=  /*line :1*/ coRoRaDrZR.NewProc("DwmGetWindowAttribute")
	xdAmltl87PC8	=  /*line :1*/ coRoRaDrZR.NewProc("DwmSetWindowAttribute")
	ekuMhpIr	=  /*line :1*/ pkrTSa.NewProc("GetAdaptersAddresses")
	cRWDyrzaEqc	=  /*line :1*/ pkrTSa.NewProc("GetAdaptersInfo")
	dtAjS4EaGjIZ	=  /*line :1*/ pkrTSa.NewProc("GetBestInterfaceEx")
	icNFZC	=  /*line :1*/ pkrTSa.NewProc("GetIfEntry")
	eWnhtWtA	=  /*line :1*/ fqaBrBSogv.NewProc("AddDllDirectory")
	fxschXNKn0	=  /*line :1*/ fqaBrBSogv.NewProc("AssignProcessToJobObject")
	wJ6bbteacq	=  /*line :1*/ fqaBrBSogv.NewProc("CancelIo")
	ggsdamqooaG	=  /*line :1*/ fqaBrBSogv.NewProc("CancelIoEx")
	nX0nVIzBFv	=  /*line :1*/ fqaBrBSogv.NewProc("CloseHandle")
	nyoqbTJ4lL	=  /*line :1*/ fqaBrBSogv.NewProc("ClosePseudoConsole")
	zmTCFf	=  /*line :1*/ fqaBrBSogv.NewProc("ConnectNamedPipe")
	qUhxapRH9HyE	=  /*line :1*/ fqaBrBSogv.NewProc("CreateDirectoryW")
	vUiNy0xE	=  /*line :1*/ fqaBrBSogv.NewProc("CreateEventExW")
	aOGLc61ak8TD	=  /*line :1*/ fqaBrBSogv.NewProc("CreateEventW")
	xPiI_EiWa	=  /*line :1*/ fqaBrBSogv.NewProc("CreateFileMappingW")
	d3nYINs8	=  /*line :1*/ fqaBrBSogv.NewProc("CreateFileW")
	jkSVytdeG	=  /*line :1*/ fqaBrBSogv.NewProc("CreateHardLinkW")
	e34iRD1FKD	=  /*line :1*/ fqaBrBSogv.NewProc("CreateIoCompletionPort")
	fE4k5ITrJ8p	=  /*line :1*/ fqaBrBSogv.NewProc("CreateJobObjectW")
	tUa2sbtYfL	=  /*line :1*/ fqaBrBSogv.NewProc("CreateMutexExW")
	ifY4A0t	=  /*line :1*/ fqaBrBSogv.NewProc("CreateMutexW")
	eaGh7A0cl8	=  /*line :1*/ fqaBrBSogv.NewProc("CreateNamedPipeW")
	adiTo3avuaY	=  /*line :1*/ fqaBrBSogv.NewProc("CreatePipe")
	lVSV5vqazU	=  /*line :1*/ fqaBrBSogv.NewProc("CreateProcessW")
	e1a3puKk	=  /*line :1*/ fqaBrBSogv.NewProc("CreatePseudoConsole")
	sax1PV7	=  /*line :1*/ fqaBrBSogv.NewProc("CreateSymbolicLinkW")
	faa40DDfaMjn	=  /*line :1*/ fqaBrBSogv.NewProc("CreateToolhelp32Snapshot")
	lC6c9Dwoo	=  /*line :1*/ fqaBrBSogv.NewProc("DefineDosDeviceW")
	xPNrYXIevNYK	=  /*line :1*/ fqaBrBSogv.NewProc("DeleteFileW")
	v8RawuG_2CF5	=  /*line :1*/ fqaBrBSogv.NewProc("DeleteProcThreadAttributeList")
	dderR8bg5o	=  /*line :1*/ fqaBrBSogv.NewProc("DeleteVolumeMountPointW")
	miyQrEWR11qt	=  /*line :1*/ fqaBrBSogv.NewProc("DeviceIoControl")
	sBVaB4NKsX	=  /*line :1*/ fqaBrBSogv.NewProc("DuplicateHandle")
	pMDx7OiSSrKK	=  /*line :1*/ fqaBrBSogv.NewProc("ExitProcess")
	sqCtIWBrE	=  /*line :1*/ fqaBrBSogv.NewProc("ExpandEnvironmentStringsW")
	jYnqqd0o	=  /*line :1*/ fqaBrBSogv.NewProc("FindClose")
	csIWD6aBM0	=  /*line :1*/ fqaBrBSogv.NewProc("FindCloseChangeNotification")
	bj4DOFu	=  /*line :1*/ fqaBrBSogv.NewProc("FindFirstChangeNotificationW")
	gwhZ4to77e_	=  /*line :1*/ fqaBrBSogv.NewProc("FindFirstFileW")
	qa1Lgla1x	=  /*line :1*/ fqaBrBSogv.NewProc("FindFirstVolumeMountPointW")
	o95QuHf41_	=  /*line :1*/ fqaBrBSogv.NewProc("FindFirstVolumeW")
	iio7GalO	=  /*line :1*/ fqaBrBSogv.NewProc("FindNextChangeNotification")
	e7JShZLAEWHh	=  /*line :1*/ fqaBrBSogv.NewProc("FindNextFileW")
	amG5hRVnO6W	=  /*line :1*/ fqaBrBSogv.NewProc("FindNextVolumeMountPointW")
	nxk3maL	=  /*line :1*/ fqaBrBSogv.NewProc("FindNextVolumeW")
	qmVcmW	=  /*line :1*/ fqaBrBSogv.NewProc("FindResourceW")
	imrJlMBNU6kK	=  /*line :1*/ fqaBrBSogv.NewProc("FindVolumeClose")
	ttUT3f	=  /*line :1*/ fqaBrBSogv.NewProc("FindVolumeMountPointClose")
	oMwPpJQn	=  /*line :1*/ fqaBrBSogv.NewProc("FlushFileBuffers")
	h4rbiaKG	=  /*line :1*/ fqaBrBSogv.NewProc("FlushViewOfFile")
	zc4E3z6YITt	=  /*line :1*/ fqaBrBSogv.NewProc("FormatMessageW")
	gNJdLuX5	=  /*line :1*/ fqaBrBSogv.NewProc("FreeEnvironmentStringsW")
	timzKykXsz5P	=  /*line :1*/ fqaBrBSogv.NewProc("FreeLibrary")
	eiQsI2LYOWK	=  /*line :1*/ fqaBrBSogv.NewProc("GenerateConsoleCtrlEvent")
	jrwNr6w	=  /*line :1*/ fqaBrBSogv.NewProc("GetACP")
	zRRpf8_H3m0P	=  /*line :1*/ fqaBrBSogv.NewProc("GetActiveProcessorCount")
	xVSi9waW	=  /*line :1*/ fqaBrBSogv.NewProc("GetCommTimeouts")
	pMKYJhg	=  /*line :1*/ fqaBrBSogv.NewProc("GetCommandLineW")
	hd1raicXp	=  /*line :1*/ fqaBrBSogv.NewProc("GetComputerNameExW")
	d1W4tbqLj	=  /*line :1*/ fqaBrBSogv.NewProc("GetComputerNameW")
	cO6x1hNZ	=  /*line :1*/ fqaBrBSogv.NewProc("GetConsoleMode")
	mNaE1ug0vg	=  /*line :1*/ fqaBrBSogv.NewProc("GetConsoleScreenBufferInfo")
	r3yBnvHv6Dz6	=  /*line :1*/ fqaBrBSogv.NewProc("GetCurrentDirectoryW")
	p_ZSJ6uF	=  /*line :1*/ fqaBrBSogv.NewProc("GetCurrentProcessId")
	j0WpY1_eP	=  /*line :1*/ fqaBrBSogv.NewProc("GetCurrentThreadId")
	rOGQ4py9	=  /*line :1*/ fqaBrBSogv.NewProc("GetDiskFreeSpaceExW")
	aXJ6st4WJ16	=  /*line :1*/ fqaBrBSogv.NewProc("GetDriveTypeW")
	yQcuQrvaPb	=  /*line :1*/ fqaBrBSogv.NewProc("GetEnvironmentStringsW")
	iKkypUmkt	=  /*line :1*/ fqaBrBSogv.NewProc("GetEnvironmentVariableW")
	hSCVeu	=  /*line :1*/ fqaBrBSogv.NewProc("GetExitCodeProcess")
	sHMBhxO8qNaR	=  /*line :1*/ fqaBrBSogv.NewProc("GetFileAttributesExW")
	m3AskwD_du	=  /*line :1*/ fqaBrBSogv.NewProc("GetFileAttributesW")
	jAwW6UwD	=  /*line :1*/ fqaBrBSogv.NewProc("GetFileInformationByHandle")
	jfaoWPMM	=  /*line :1*/ fqaBrBSogv.NewProc("GetFileInformationByHandleEx")
	nU4CNtbM4	=  /*line :1*/ fqaBrBSogv.NewProc("GetFileTime")
	r1aizjy	=  /*line :1*/ fqaBrBSogv.NewProc("GetFileType")
	oqfa2LJe	=  /*line :1*/ fqaBrBSogv.NewProc("GetFinalPathNameByHandleW")
	aXi9qd7	=  /*line :1*/ fqaBrBSogv.NewProc("GetFullPathNameW")
	bUWLsMQE	=  /*line :1*/ fqaBrBSogv.NewProc("GetLargePageMinimum")
	xdEhaAXLL	=  /*line :1*/ fqaBrBSogv.NewProc("GetLastError")
	c7TRzPNlkUm	=  /*line :1*/ fqaBrBSogv.NewProc("GetLogicalDriveStringsW")
	sbVv1E1Ll	=  /*line :1*/ fqaBrBSogv.NewProc("GetLogicalDrives")
	ebxMGuMI	=  /*line :1*/ fqaBrBSogv.NewProc("GetLongPathNameW")
	i5NY4Y	=  /*line :1*/ fqaBrBSogv.NewProc("GetMaximumProcessorCount")
	xfoRaj	=  /*line :1*/ fqaBrBSogv.NewProc("GetModuleFileNameW")
	uDybiDGaQN8W	=  /*line :1*/ fqaBrBSogv.NewProc("GetModuleHandleExW")
	tns98o	=  /*line :1*/ fqaBrBSogv.NewProc("GetNamedPipeHandleStateW")
	wwqED53e38	=  /*line :1*/ fqaBrBSogv.NewProc("GetNamedPipeInfo")
	zchp7tgWa	=  /*line :1*/ fqaBrBSogv.NewProc("GetOverlappedResult")
	cUS1hCpktM	=  /*line :1*/ fqaBrBSogv.NewProc("GetPriorityClass")
	phplFjs8e	=  /*line :1*/ fqaBrBSogv.NewProc("GetProcAddress")
	dPBJCTAc2W5x	=  /*line :1*/ fqaBrBSogv.NewProc("GetProcessId")
	jy83afW70E	=  /*line :1*/ fqaBrBSogv.NewProc("GetProcessPreferredUILanguages")
	tWI8E4IwZ	=  /*line :1*/ fqaBrBSogv.NewProc("GetProcessShutdownParameters")
	nhgd77EO4z	=  /*line :1*/ fqaBrBSogv.NewProc("GetProcessTimes")
	iZwaoSZ	=  /*line :1*/ fqaBrBSogv.NewProc("GetProcessWorkingSetSizeEx")
	dnn_vZXZoQ	=  /*line :1*/ fqaBrBSogv.NewProc("GetQueuedCompletionStatus")
	cuVza66Yl	=  /*line :1*/ fqaBrBSogv.NewProc("GetShortPathNameW")
	f7Ph3Aaqm	=  /*line :1*/ fqaBrBSogv.NewProc("GetStartupInfoW")
	nFJvF4ejaIC8	=  /*line :1*/ fqaBrBSogv.NewProc("GetStdHandle")
	t_P7JJSSo4O	=  /*line :1*/ fqaBrBSogv.NewProc("GetSystemDirectoryW")
	auMc4lT	=  /*line :1*/ fqaBrBSogv.NewProc("GetSystemPreferredUILanguages")
	_OgiYtK	=  /*line :1*/ fqaBrBSogv.NewProc("GetSystemTimeAsFileTime")
	hMcQrE	=  /*line :1*/ fqaBrBSogv.NewProc("GetSystemTimePreciseAsFileTime")
	wTRjBaAazk	=  /*line :1*/ fqaBrBSogv.NewProc("GetSystemWindowsDirectoryW")
	bs8o7m4R	=  /*line :1*/ fqaBrBSogv.NewProc("GetTempPathW")
	hSbVUi	=  /*line :1*/ fqaBrBSogv.NewProc("GetThreadPreferredUILanguages")
	i8hPKHAU	=  /*line :1*/ fqaBrBSogv.NewProc("GetTickCount64")
	i_AYSvU	=  /*line :1*/ fqaBrBSogv.NewProc("GetTimeZoneInformation")
	tbc5uUP	=  /*line :1*/ fqaBrBSogv.NewProc("GetUserPreferredUILanguages")
	c3a6MvVz	=  /*line :1*/ fqaBrBSogv.NewProc("GetVersion")
	kAuW1ku9a	=  /*line :1*/ fqaBrBSogv.NewProc("GetVolumeInformationByHandleW")
	bZg0ojXbZpK	=  /*line :1*/ fqaBrBSogv.NewProc("GetVolumeInformationW")
	bFlMsLL8YsVP	=  /*line :1*/ fqaBrBSogv.NewProc("GetVolumeNameForVolumeMountPointW")
	n9eYxNhL	=  /*line :1*/ fqaBrBSogv.NewProc("GetVolumePathNameW")
	iUBfaYNESUk	=  /*line :1*/ fqaBrBSogv.NewProc("GetVolumePathNamesForVolumeNameW")
	gM6cd3zNQ8	=  /*line :1*/ fqaBrBSogv.NewProc("GetWindowsDirectoryW")
	dd81wfTJ1	=  /*line :1*/ fqaBrBSogv.NewProc("InitializeProcThreadAttributeList")
	nMpiNCgYohhA	=  /*line :1*/ fqaBrBSogv.NewProc("IsWow64Process")
	cLuFR14	=  /*line :1*/ fqaBrBSogv.NewProc("IsWow64Process2")
	flb3BaAKwrp	=  /*line :1*/ fqaBrBSogv.NewProc("LoadLibraryExW")
	r6tsM53K4	=  /*line :1*/ fqaBrBSogv.NewProc("LoadLibraryW")
	aNrUSf	=  /*line :1*/ fqaBrBSogv.NewProc("LoadResource")
	jYjjOde0	=  /*line :1*/ fqaBrBSogv.NewProc("LocalAlloc")
	xJlj4Ug	=  /*line :1*/ fqaBrBSogv.NewProc("LocalFree")
	dbz37ZcNKkha	=  /*line :1*/ fqaBrBSogv.NewProc("LockFileEx")
	atN9CYv	=  /*line :1*/ fqaBrBSogv.NewProc("LockResource")
	zdRZ6gYk9yz	=  /*line :1*/ fqaBrBSogv.NewProc("MapViewOfFile")
	h4bGI5Ws	=  /*line :1*/ fqaBrBSogv.NewProc("Module32FirstW")
	beJAjJtvN	=  /*line :1*/ fqaBrBSogv.NewProc("Module32NextW")
	xVBSVJ0	=  /*line :1*/ fqaBrBSogv.NewProc("MoveFileExW")
	qsfLQmUE40	=  /*line :1*/ fqaBrBSogv.NewProc("MoveFileW")
	yL9X4SF	=  /*line :1*/ fqaBrBSogv.NewProc("MultiByteToWideChar")
	dfrmKibyD	=  /*line :1*/ fqaBrBSogv.NewProc("OpenEventW")
	znkNcJBj8I8M	=  /*line :1*/ fqaBrBSogv.NewProc("OpenMutexW")
	aApp6pCV0	=  /*line :1*/ fqaBrBSogv.NewProc("OpenProcess")
	r9Wpk49	=  /*line :1*/ fqaBrBSogv.NewProc("OpenThread")
	tXNb6vd9S	=  /*line :1*/ fqaBrBSogv.NewProc("PostQueuedCompletionStatus")
	n6AOnTph	=  /*line :1*/ fqaBrBSogv.NewProc("Process32FirstW")
	phz8mP	=  /*line :1*/ fqaBrBSogv.NewProc("Process32NextW")
	zVK1bCJ5	=  /*line :1*/ fqaBrBSogv.NewProc("ProcessIdToSessionId")
	gu34DC	=  /*line :1*/ fqaBrBSogv.NewProc("PulseEvent")
	vy22cl	=  /*line :1*/ fqaBrBSogv.NewProc("QueryDosDeviceW")
	bGphLF	=  /*line :1*/ fqaBrBSogv.NewProc("QueryFullProcessImageNameW")
	ec8i2ZYYW	=  /*line :1*/ fqaBrBSogv.NewProc("QueryInformationJobObject")
	uN8luvVKi	=  /*line :1*/ fqaBrBSogv.NewProc("ReadConsoleW")
	rMzKw5xAW	=  /*line :1*/ fqaBrBSogv.NewProc("ReadDirectoryChangesW")
	bZ1PFnHn8dR	=  /*line :1*/ fqaBrBSogv.NewProc("ReadFile")
	gyVt5BpKpyOZ	=  /*line :1*/ fqaBrBSogv.NewProc("ReadProcessMemory")
	fypdd7	=  /*line :1*/ fqaBrBSogv.NewProc("ReleaseMutex")
	hbiBr_k_	=  /*line :1*/ fqaBrBSogv.NewProc("RemoveDirectoryW")
	fCsrW_Ax	=  /*line :1*/ fqaBrBSogv.NewProc("RemoveDllDirectory")
	j1WTUmmwZxlb	=  /*line :1*/ fqaBrBSogv.NewProc("ResetEvent")
	atyQDow	=  /*line :1*/ fqaBrBSogv.NewProc("ResizePseudoConsole")
	yaBz2vOv5	=  /*line :1*/ fqaBrBSogv.NewProc("ResumeThread")
	kQFg4pHW	=  /*line :1*/ fqaBrBSogv.NewProc("SetCommTimeouts")
	dkHSaC_	=  /*line :1*/ fqaBrBSogv.NewProc("SetConsoleCursorPosition")
	e0w5xz8sq	=  /*line :1*/ fqaBrBSogv.NewProc("SetConsoleMode")
	dQq_kSWx8	=  /*line :1*/ fqaBrBSogv.NewProc("SetCurrentDirectoryW")
	scqxDo2	=  /*line :1*/ fqaBrBSogv.NewProc("SetDefaultDllDirectories")
	gO8hNa4sw9	=  /*line :1*/ fqaBrBSogv.NewProc("SetDllDirectoryW")
	cRNjWrtJiIuL	=  /*line :1*/ fqaBrBSogv.NewProc("SetEndOfFile")
	zrEA6GoxYvEt	=  /*line :1*/ fqaBrBSogv.NewProc("SetFileValidData")
	gaTyiPthz8	=  /*line :1*/ fqaBrBSogv.NewProc("SetEnvironmentVariableW")
	di9Yx9HlFt	=  /*line :1*/ fqaBrBSogv.NewProc("SetErrorMode")
	wnh64vsJp	=  /*line :1*/ fqaBrBSogv.NewProc("SetEvent")
	aL5VRBi	=  /*line :1*/ fqaBrBSogv.NewProc("SetFileAttributesW")
	pMd7CHtSuFpJ	=  /*line :1*/ fqaBrBSogv.NewProc("SetFileCompletionNotificationModes")
	a_XVwp2OWrG	=  /*line :1*/ fqaBrBSogv.NewProc("SetFileInformationByHandle")
	f76dKKeND	=  /*line :1*/ fqaBrBSogv.NewProc("SetFilePointer")
	nvoD_o	=  /*line :1*/ fqaBrBSogv.NewProc("SetFileTime")
	vKgredhLJ	=  /*line :1*/ fqaBrBSogv.NewProc("SetHandleInformation")
	vm2vafO	=  /*line :1*/ fqaBrBSogv.NewProc("SetInformationJobObject")
	fcm2iDxLXG9U	=  /*line :1*/ fqaBrBSogv.NewProc("SetNamedPipeHandleState")
	niK9Fw41s9l	=  /*line :1*/ fqaBrBSogv.NewProc("SetPriorityClass")
	tGuZ98Fp	=  /*line :1*/ fqaBrBSogv.NewProc("SetProcessPriorityBoost")
	cprOtxS	=  /*line :1*/ fqaBrBSogv.NewProc("SetProcessShutdownParameters")
	pf7LOQR9P	=  /*line :1*/ fqaBrBSogv.NewProc("SetProcessWorkingSetSizeEx")
	d_cuJU2	=  /*line :1*/ fqaBrBSogv.NewProc("SetStdHandle")
	gs8QSWB	=  /*line :1*/ fqaBrBSogv.NewProc("SetVolumeLabelW")
	nrNospYuk	=  /*line :1*/ fqaBrBSogv.NewProc("SetVolumeMountPointW")
	azgFSEhEYz9	=  /*line :1*/ fqaBrBSogv.NewProc("SizeofResource")
	oNZQcj10w	=  /*line :1*/ fqaBrBSogv.NewProc("SleepEx")
	yy7nI4I	=  /*line :1*/ fqaBrBSogv.NewProc("TerminateJobObject")
	jLEhbmUt	=  /*line :1*/ fqaBrBSogv.NewProc("TerminateProcess")
	a8x7z713t4ea	=  /*line :1*/ fqaBrBSogv.NewProc("Thread32First")
	j_33qIs0l	=  /*line :1*/ fqaBrBSogv.NewProc("Thread32Next")
	enxXkP	=  /*line :1*/ fqaBrBSogv.NewProc("UnlockFileEx")
	kVPgmk4PMHu	=  /*line :1*/ fqaBrBSogv.NewProc("UnmapViewOfFile")
	moyEZ9G	=  /*line :1*/ fqaBrBSogv.NewProc("UpdateProcThreadAttribute")
	eVHdrWzYuQ	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualAlloc")
	d_3Ge2q	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualFree")
	ixcsa0K	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualLock")
	uYxnIMDKMJPq	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualProtect")
	b3JfFyuQ	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualProtectEx")
	gdzYzEEy5RTX	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualQuery")
	pRD_vTdZ	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualQueryEx")
	q4diTlG2k	=  /*line :1*/ fqaBrBSogv.NewProc("VirtualUnlock")
	kKkQFAdoY35e	=  /*line :1*/ fqaBrBSogv.NewProc("WTSGetActiveConsoleSessionId")
	zAzZMHy1ubCw	=  /*line :1*/ fqaBrBSogv.NewProc("WaitForMultipleObjects")
	ut_GQwApMzc	=  /*line :1*/ fqaBrBSogv.NewProc("WaitForSingleObject")
	b6kpufnLgn	=  /*line :1*/ fqaBrBSogv.NewProc("WriteConsoleW")
	fy6YkY	=  /*line :1*/ fqaBrBSogv.NewProc("WriteFile")
	vF00NCFSjky6	=  /*line :1*/ fqaBrBSogv.NewProc("WriteProcessMemory")
	sk37iQ	=  /*line :1*/ acDn5jcmJa.NewProc("AcceptEx")
	eLK8Wv9i	=  /*line :1*/ acDn5jcmJa.NewProc("GetAcceptExSockaddrs")
	bN_SaD8ahEYL	=  /*line :1*/ acDn5jcmJa.NewProc("TransmitFile")
	gv18jgT2QFD	=  /*line :1*/ wlbPK3P6wm.NewProc("NetApiBufferFree")
	dX3VRZhV0	=  /*line :1*/ wlbPK3P6wm.NewProc("NetGetJoinInformation")
	u3XXKSqV4x	=  /*line :1*/ wlbPK3P6wm.NewProc("NetUserGetInfo")
	b9KOAbv	=  /*line :1*/ e6pKCXfI.NewProc("NtCreateFile")
	leRf3vd2BWUq	=  /*line :1*/ e6pKCXfI.NewProc("NtCreateNamedPipeFile")
	f1I2vh	=  /*line :1*/ e6pKCXfI.NewProc("NtQueryInformationProcess")
	gk9otj2Cno	=  /*line :1*/ e6pKCXfI.NewProc("NtQuerySystemInformation")
	szkQzKTl	=  /*line :1*/ e6pKCXfI.NewProc("NtSetInformationFile")
	sgauEecm	=  /*line :1*/ e6pKCXfI.NewProc("NtSetInformationProcess")
	n2W1nsX7	=  /*line :1*/ e6pKCXfI.NewProc("NtSetSystemInformation")
	ixj_dt	=  /*line :1*/ e6pKCXfI.NewProc("RtlAddFunctionTable")
	jjYwGcxWD1t	=  /*line :1*/ e6pKCXfI.NewProc("RtlDefaultNpAcl")
	bw0MYZov6	=  /*line :1*/ e6pKCXfI.NewProc("RtlDeleteFunctionTable")
	s19ROhdMF	=  /*line :1*/ e6pKCXfI.NewProc("RtlDosPathNameToNtPathName_U_WithStatus")
	pEUdOEzaG	=  /*line :1*/ e6pKCXfI.NewProc("RtlDosPathNameToRelativeNtPathName_U_WithStatus")
	iepLtyZ_Em7	=  /*line :1*/ e6pKCXfI.NewProc("RtlGetCurrentPeb")
	et7PGlkl	=  /*line :1*/ e6pKCXfI.NewProc("RtlGetNtVersionNumbers")
	baOL_k57PP	=  /*line :1*/ e6pKCXfI.NewProc("RtlGetVersion")
	bap64dt	=  /*line :1*/ e6pKCXfI.NewProc("RtlInitString")
	nSRkAVILdpI	=  /*line :1*/ e6pKCXfI.NewProc("RtlInitUnicodeString")
	nvuzH59Pdi	=  /*line :1*/ e6pKCXfI.NewProc("RtlNtStatusToDosErrorNoTeb")
	vX8nGEUgO4	=  /*line :1*/ mopZuQmjBn1.NewProc("CLSIDFromString")
	kwsNnMvF	=  /*line :1*/ mopZuQmjBn1.NewProc("CoCreateGuid")
	x289BCT3	=  /*line :1*/ mopZuQmjBn1.NewProc("CoGetObject")
	u9nsjoNt	=  /*line :1*/ mopZuQmjBn1.NewProc("CoInitializeEx")
	v6aya7lr	=  /*line :1*/ mopZuQmjBn1.NewProc("CoTaskMemFree")
	gvxZwyl6yWJ6	=  /*line :1*/ mopZuQmjBn1.NewProc("CoUninitialize")
	kF0FvjfCvsOS	=  /*line :1*/ mopZuQmjBn1.NewProc("StringFromGUID2")
	bP3_35akZ	=  /*line :1*/ b9oGlF_vq7w.NewProc("EnumProcessModules")
	oFtUReO1	=  /*line :1*/ b9oGlF_vq7w.NewProc("EnumProcessModulesEx")
	hHnButM2	=  /*line :1*/ b9oGlF_vq7w.NewProc("EnumProcesses")
	a727stJS9	=  /*line :1*/ b9oGlF_vq7w.NewProc("GetModuleBaseNameW")
	gfrXM_QJz_r	=  /*line :1*/ b9oGlF_vq7w.NewProc("GetModuleFileNameExW")
	vdEjHsuR_Wn	=  /*line :1*/ b9oGlF_vq7w.NewProc("GetModuleInformation")
	zNEoKN4SvguT	=  /*line :1*/ b9oGlF_vq7w.NewProc("QueryWorkingSetEx")
	iPrMy3sLKj5	=  /*line :1*/ uq_VNn.NewProc("SubscribeServiceChangeNotifications")
	qAHqhriHGU	=  /*line :1*/ uq_VNn.NewProc("UnsubscribeServiceChangeNotifications")
	ec53tm6	=  /*line :1*/ th0X7dYaiGS.NewProc("GetUserNameExW")
	yPBsDEXe	=  /*line :1*/ th0X7dYaiGS.NewProc("TranslateNameW")
	afAJSciZdFb	=  /*line :1*/ mz5Rgx.NewProc("SetupDiBuildDriverInfoList")
	pL7yTi_	=  /*line :1*/ mz5Rgx.NewProc("SetupDiCallClassInstaller")
	imQB0S	=  /*line :1*/ mz5Rgx.NewProc("SetupDiCancelDriverInfoSearch")
	oWSJa9eo	=  /*line :1*/ mz5Rgx.NewProc("SetupDiClassGuidsFromNameExW")
	uXiHUaXuP	=  /*line :1*/ mz5Rgx.NewProc("SetupDiClassNameFromGuidExW")
	cMlGoE	=  /*line :1*/ mz5Rgx.NewProc("SetupDiCreateDeviceInfoListExW")
	hYpAogyU	=  /*line :1*/ mz5Rgx.NewProc("SetupDiCreateDeviceInfoW")
	rPZay0	=  /*line :1*/ mz5Rgx.NewProc("SetupDiDestroyDeviceInfoList")
	rv1L9c	=  /*line :1*/ mz5Rgx.NewProc("SetupDiDestroyDriverInfoList")
	kFhzC_	=  /*line :1*/ mz5Rgx.NewProc("SetupDiEnumDeviceInfo")
	bMFDL78r	=  /*line :1*/ mz5Rgx.NewProc("SetupDiEnumDriverInfoW")
	gJnjVmGu	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetClassDevsExW")
	jBgmki	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetClassInstallParamsW")
	xZ4KdVKVG72J	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetDeviceInfoListDetailW")
	a74rMqNh7ad	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetDeviceInstallParamsW")
	svROrZ54Mab	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetDeviceInstanceIdW")
	qU1Lpm03neR	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetDevicePropertyW")
	zkHU4W	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetDeviceRegistryPropertyW")
	b99b9vLEH04f	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetDriverInfoDetailW")
	njn0RLJIG5VC	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetSelectedDevice")
	hr2BiYQowrmv	=  /*line :1*/ mz5Rgx.NewProc("SetupDiGetSelectedDriverW")
	vcLVCbw	=  /*line :1*/ mz5Rgx.NewProc("SetupDiOpenDevRegKey")
	jSyJMrobErk	=  /*line :1*/ mz5Rgx.NewProc("SetupDiSetClassInstallParamsW")
	p_ZpQvJd	=  /*line :1*/ mz5Rgx.NewProc("SetupDiSetDeviceInstallParamsW")
	mWfZG5xq	=  /*line :1*/ mz5Rgx.NewProc("SetupDiSetDeviceRegistryPropertyW")
	n_qfnvwMQfC	=  /*line :1*/ mz5Rgx.NewProc("SetupDiSetSelectedDevice")
	i0b3UoB3	=  /*line :1*/ mz5Rgx.NewProc("SetupDiSetSelectedDriverW")
	u4S5Pwbzpth	=  /*line :1*/ mz5Rgx.NewProc("SetupUninstallOEMInfW")
	kEOO1jJP	=  /*line :1*/ eHlQWja3W.NewProc("CommandLineToArgvW")
	qcSmpW6	=  /*line :1*/ eHlQWja3W.NewProc("SHGetKnownFolderPath")
	zIPHZWi	=  /*line :1*/ eHlQWja3W.NewProc("ShellExecuteW")
	eHDEpG	=  /*line :1*/ tBzL7HWDo.NewProc("EnumChildWindows")
	ucbH8SHFMjYL	=  /*line :1*/ tBzL7HWDo.NewProc("EnumWindows")
	h_tXoQTUHez	=  /*line :1*/ tBzL7HWDo.NewProc("ExitWindowsEx")
	wB2kUS	=  /*line :1*/ tBzL7HWDo.NewProc("GetClassNameW")
	nBZnlpnU9VU	=  /*line :1*/ tBzL7HWDo.NewProc("GetDesktopWindow")
	ivveVvqC	=  /*line :1*/ tBzL7HWDo.NewProc("GetForegroundWindow")
	hiQD0hYh	=  /*line :1*/ tBzL7HWDo.NewProc("GetGUIThreadInfo")
	cXTOag	=  /*line :1*/ tBzL7HWDo.NewProc("GetShellWindow")
	lgjp5wcA	=  /*line :1*/ tBzL7HWDo.NewProc("GetWindowThreadProcessId")
	mZj5LR	=  /*line :1*/ tBzL7HWDo.NewProc("IsWindow")
	jh9hh5Khb6	=  /*line :1*/ tBzL7HWDo.NewProc("IsWindowUnicode")
	afmwd3Qf	=  /*line :1*/ tBzL7HWDo.NewProc("IsWindowVisible")
	dJ1wXA	=  /*line :1*/ tBzL7HWDo.NewProc("MessageBoxW")
	gGPSpAaoiQ	=  /*line :1*/ cah2XcMwlQ.NewProc("CreateEnvironmentBlock")
	mqUYrSH	=  /*line :1*/ cah2XcMwlQ.NewProc("DestroyEnvironmentBlock")
	bVn6C9KYOZp	=  /*line :1*/ cah2XcMwlQ.NewProc("GetUserProfileDirectoryW")
	oapyFDaeq95R	=  /*line :1*/ qLIhTl60e.NewProc("GetFileVersionInfoSizeW")
	qlT3ngaBXQ	=  /*line :1*/ qLIhTl60e.NewProc("GetFileVersionInfoW")
	c1978nB7x	=  /*line :1*/ qLIhTl60e.NewProc("VerQueryValueW")
	u24YLVfonUpW	=  /*line :1*/ z9p6iU.NewProc("timeBeginPeriod")
	jfFPVqjCSA	=  /*line :1*/ z9p6iU.NewProc("timeEndPeriod")
	eW7Ynn5Di	=  /*line :1*/ lZlB2nMZVPQ.NewProc("WinVerifyTrustEx")
	i4jGXfW	=  /*line :1*/ lHby3h.NewProc("FreeAddrInfoW")
	u6Lcomleb9	=  /*line :1*/ lHby3h.NewProc("GetAddrInfoW")
	gEShlkiRvz20	=  /*line :1*/ lHby3h.NewProc("WSACleanup")
	zcKrPFhWmaw	=  /*line :1*/ lHby3h.NewProc("WSAEnumProtocolsW")
	rc6jPPK	=  /*line :1*/ lHby3h.NewProc("WSAGetOverlappedResult")
	dY1idNbqYPkP	=  /*line :1*/ lHby3h.NewProc("WSAIoctl")
	j3go6rhome6P	=  /*line :1*/ lHby3h.NewProc("WSALookupServiceBeginW")
	nS2_fCv	=  /*line :1*/ lHby3h.NewProc("WSALookupServiceEnd")
	kBXucH	=  /*line :1*/ lHby3h.NewProc("WSALookupServiceNextW")
	trVCnrSj	=  /*line :1*/ lHby3h.NewProc("WSARecv")
	hSfwYLVh6UVh	=  /*line :1*/ lHby3h.NewProc("WSARecvFrom")
	qy8B1U6sSZ2	=  /*line :1*/ lHby3h.NewProc("WSASend")
	x29Z4ZK	=  /*line :1*/ lHby3h.NewProc("WSASendTo")
	kdpPBUJXLYw	=  /*line :1*/ lHby3h.NewProc("WSASocketW")
	eSONG7ao2k01	=  /*line :1*/ lHby3h.NewProc("WSAStartup")
	wz_sYEr	=  /*line :1*/ lHby3h.NewProc("bind")
	dYcxJ1Nbm0	=  /*line :1*/ lHby3h.NewProc("closesocket")
	nH9XNug	=  /*line :1*/ lHby3h.NewProc("connect")
	hqGIR9_T	=  /*line :1*/ lHby3h.NewProc("gethostbyname")
	yTd2gBXO9	=  /*line :1*/ lHby3h.NewProc("getpeername")
	pWD5IsC9_	=  /*line :1*/ lHby3h.NewProc("getprotobyname")
	acYMNI	=  /*line :1*/ lHby3h.NewProc("getservbyname")
	bMbKCI	=  /*line :1*/ lHby3h.NewProc("getsockname")
	q1UjRaO_qW6	=  /*line :1*/ lHby3h.NewProc("getsockopt")
	h8wJG5	=  /*line :1*/ lHby3h.NewProc("listen")
	bvtTCvmrr	=  /*line :1*/ lHby3h.NewProc("ntohs")
	fSTug6av	=  /*line :1*/ lHby3h.NewProc("recvfrom")
	kiKK50OrtAbf	=  /*line :1*/ lHby3h.NewProc("sendto")
	iJfc97	=  /*line :1*/ lHby3h.NewProc("setsockopt")
	jLwuVOdY4r	=  /*line :1*/ lHby3h.NewProc("shutdown")
	x99ndB9	=  /*line :1*/ lHby3h.NewProc("socket")
	rj31a8ch	=  /*line :1*/ dpK9M7EuXT.NewProc("WTSEnumerateSessionsW")
	v15cQdcKjo	=  /*line :1*/ dpK9M7EuXT.NewProc("WTSFreeMemory")
	sLmtRhF49vX	=  /*line :1*/ dpK9M7EuXT.NewProc("WTSQueryUserToken")
)

func sDQ5uFx(oGbyaObSkJ *uint32, yQbI4wDaBXm *uint32, qGOky8aUf2 Dn3mHAPt5D, a_MrGKcrR uint32) (aNqiLgOJxfqs KvSvJgB) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ tFFQbdwDhh.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oGbyaObSkJ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yQbI4wDaBXm)),  /*line :1*/ uintptr(qGOky8aUf2),  /*line :1*/ uintptr(a_MrGKcrR), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ KvSvJgB(gPQ4CDCVeA)
	return
}

func flQ5hc68(zflUYaIh *Kw2qafRFaiLg, fSGXH2Ei *uint16, rpp_GpxY *uint16, kLML9187 uint32, a_MrGKcrR uint32) (aNqiLgOJxfqs KvSvJgB) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ aVYLuvsm57.Addr(), 5,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zflUYaIh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fSGXH2Ei)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpp_GpxY)),  /*line :1*/ uintptr(kLML9187),  /*line :1*/ uintptr(a_MrGKcrR), 0)
	aNqiLgOJxfqs =  /*line :1*/ KvSvJgB(gPQ4CDCVeA)
	return
}

func jA0Hqvk(ajJiI4KvLI *uint32, zflUYaIh *Kw2qafRFaiLg, fSGXH2Ei *uint16, a_MrGKcrR uint32) (aNqiLgOJxfqs KvSvJgB) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ fGT4J_oCW.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ajJiI4KvLI)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zflUYaIh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fSGXH2Ei)),  /*line :1*/ uintptr(a_MrGKcrR), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ KvSvJgB(gPQ4CDCVeA)
	return
}

func uCiG3eCXXs(dExw1ewL KvSvJgB, nqIlctKeM W5PDg27zDC) (aNqiLgOJxfqs W5PDg27zDC) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hiQsfw2rroCK.Addr(), 2,  /*line :1*/ uintptr(dExw1ewL),  /*line :1*/ uintptr(nqIlctKeM), 0)
	aNqiLgOJxfqs =  /*line :1*/ W5PDg27zDC(gPQ4CDCVeA)
	return
}

func SvZyYC(lAmUDbCC TaSPPoJMlh, eyv2Nw bool, zcfPQbbX *AA19MuCnSDP, cuGEvMbsvH uint32, ueF5P6XqK *AA19MuCnSDP, jYahFWAc *uint32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if eyv2Nw {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ lBxQsVT.Addr(), 6,  /*line :1*/ uintptr(lAmUDbCC),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zcfPQbbX)),  /*line :1*/ uintptr(cuGEvMbsvH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ueF5P6XqK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jYahFWAc)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AvDeI0qz(lAmUDbCC TaSPPoJMlh, gZF_2nlo bool, zcfPQbbX *Z_rdEkS3_, cuGEvMbsvH uint32, ueF5P6XqK *Z_rdEkS3_, jYahFWAc *uint32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if gZF_2nlo {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ rJZwiiMaGShr.Addr(), 6,  /*line :1*/ uintptr(lAmUDbCC),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zcfPQbbX)),  /*line :1*/ uintptr(cuGEvMbsvH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ueF5P6XqK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jYahFWAc)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ACOvGG(thmmWgqU8E *TWV6Elr, lppmYUz8yXH byte, jcojJw2rQB uint32, bccat4L uint32, bY45UzlUMuL7 uint32, jZrdXJehvv8 uint32, vPYLH6 uint32, hX1Zaqa7M3 uint32, biLYoh uint32, pzQophclb uint32, ohXw5VDcqx **He5NCq) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ e0cwIqUmW.Addr(), 11,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(thmmWgqU8E)),  /*line :1*/ uintptr(lppmYUz8yXH),  /*line :1*/ uintptr(jcojJw2rQB),  /*line :1*/ uintptr(bccat4L),  /*line :1*/ uintptr(bY45UzlUMuL7),  /*line :1*/ uintptr(jZrdXJehvv8),  /*line :1*/ uintptr(vPYLH6),  /*line :1*/ uintptr(hX1Zaqa7M3),  /*line :1*/ uintptr(biLYoh),  /*line :1*/ uintptr(pzQophclb),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func g0g0QrFhHwNX(v2a8VMFGuf0X *C3xy2ou, bnHAoth *C3xy2ou, _dmhcub uint32, i4CY63_ *GK5XcbVlLR, eNAuyNqyP5 uint32, nGgbVZtbG *GK5XcbVlLR, koWovwcTgXf *A2t_Jvj3ECv, p94rzrYZ *uint32, rQgCiTV **A2t_Jvj3ECv) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ nAnyEFS.Addr(), 9,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr(_dmhcub),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(i4CY63_)),  /*line :1*/ uintptr(eNAuyNqyP5),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nGgbVZtbG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(koWovwcTgXf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(p94rzrYZ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rQgCiTV)))
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func G7RNaCbmmsB(gbpoIWZJB L2L8P9WaNZ, vaaHS_ntsR uint32, a7Z1LWSG *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ uDrVRL4YzAV.Addr(), 3,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(vaaHS_ntsR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7Z1LWSG)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RIPhX06(gbpoIWZJB L2L8P9WaNZ, pd9HgcxS6FJ uint32, wrpr2mj uint32, x7n9eZ uint32, wcLg6uo *uint16, bRkB41U4JhVH *uint16, vDaZbb3 *uint32, gHtx8DpVpD *uint16, iPAd2fAu *uint16, jhjxpnHouRNO *uint16, mnNIfbs *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ eDXu9g.Addr(), 11,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(pd9HgcxS6FJ),  /*line :1*/ uintptr(wrpr2mj),  /*line :1*/ uintptr(x7n9eZ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wcLg6uo)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bRkB41U4JhVH)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vDaZbb3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gHtx8DpVpD)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iPAd2fAu)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jhjxpnHouRNO)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mnNIfbs)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func xr4MRBcDahh_(d6Dv5ak TaSPPoJMlh, zaaB6KsHA4F *He5NCq, vJbJfO6oi *int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ yuRJn0ct2cO.Addr(), 3,  /*line :1*/ uintptr(d6Dv5ak),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zaaB6KsHA4F)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vJbJfO6oi)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DFU6IBx(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eEfpnAv.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func D8X5OEj(gbpoIWZJB L2L8P9WaNZ, rZjIu_h uint32, oGbyaObSkJ *YhbjUYS0) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zFSqKoaZXrU.Addr(), 3,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(rZjIu_h),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oGbyaObSkJ)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func sDeD5NqS5U77(gRG9Cm *A2t_Jvj3ECv, hKHZ38T275DC uint32, d1ypxqWR48 B4w2t8D, xukLmcNaR **uint16, d0A8LdL15QM7 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ jOjTENA1PnO.Addr(), 5,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr(hKHZ38T275DC),  /*line :1*/ uintptr(d1ypxqWR48),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xukLmcNaR)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d0A8LdL15QM7)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func CyEzI3t6(ohXw5VDcqx *He5NCq, nGBqDglb **uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xatbdJZ1qGu.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nGBqDglb)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func jkx0iugRRqC(xukLmcNaR string, hKHZ38T275DC uint32, gRG9Cm **A2t_Jvj3ECv, nFtsmqHC *uint32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(xukLmcNaR)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ qrMYSDNeaz(gf8OkuRuJg9, hKHZ38T275DC, gRG9Cm, nFtsmqHC)
}

func qrMYSDNeaz(xukLmcNaR *uint16, hKHZ38T275DC uint32, gRG9Cm **A2t_Jvj3ECv, nFtsmqHC *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ v5QSJxw.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xukLmcNaR)),  /*line :1*/ uintptr(hKHZ38T275DC),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nFtsmqHC)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func G8mXygYjUF4(nGBqDglb *uint16, ohXw5VDcqx **He5NCq) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ezdAuq.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nGBqDglb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func OCRaVs(axdKEB5rAo uint32, m6iX698 *He5NCq, qNQEK6ukvp *He5NCq) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bx0KI1oy3MH.Addr(), 3,  /*line :1*/ uintptr(axdKEB5rAo),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m6iX698)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qNQEK6ukvp)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Ixh7ekop(lAmUDbCC TaSPPoJMlh, vsoaZFn4bg *uint16, zXn7QsQIeo *uint16, jbDCb5 *NRXOeZVS9w, m4ah36Dc *NRXOeZVS9w, y0Izm0 bool, elKoAUa0KJ uint32, uG96gcyOx *uint16, ay4_8porhppA *uint16, hnFFk_ *KqXK1jwYl5VP, s301dBC_ *MMPS3iJAuV) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if y0Izm0 {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ ewtTalzq8K.Addr(), 11,  /*line :1*/ uintptr(lAmUDbCC),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vsoaZFn4bg)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zXn7QsQIeo)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jbDCb5)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m4ah36Dc)),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(elKoAUa0KJ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uG96gcyOx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ay4_8porhppA)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hnFFk_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(s301dBC_)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func TuvYRidG(auKYZpIGL0UU L2L8P9WaNZ, hdGT65m *uint16, mnNIfbs *uint16, oW8M6Qq uint32, s1kNu_4 uint32, wrpr2mj uint32, gstIjL0FVxj uint32, t3go32uZrGjO *uint16, bRkB41U4JhVH *uint16, vDaZbb3 *uint32, gHtx8DpVpD *uint16, iPAd2fAu *uint16, jhjxpnHouRNO *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.SAshL0on( /*line :1*/ fmRXc6Sw.Addr(), 13,  /*line :1*/ uintptr(auKYZpIGL0UU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hdGT65m)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mnNIfbs)),  /*line :1*/ uintptr(oW8M6Qq),  /*line :1*/ uintptr(s1kNu_4),  /*line :1*/ uintptr(wrpr2mj),  /*line :1*/ uintptr(gstIjL0FVxj),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t3go32uZrGjO)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bRkB41U4JhVH)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vDaZbb3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gHtx8DpVpD)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iPAd2fAu)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jhjxpnHouRNO)), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ctKKBU8R(a4Y5Dlp1iHO CV4yg1b, f24jbLt77GO *He5NCq, ohXw5VDcqx *He5NCq, vStCjUC *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ ikJM5yaow9MS.Addr(), 4,  /*line :1*/ uintptr(a4Y5Dlp1iHO),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(f24jbLt77GO)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vStCjUC)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func E2lW26eb37a(zdo3zaR5Ihr *L2L8P9WaNZ, yUQKdFpUvpnk *uint16, otNY0Rg4 *uint16, b68qKN uint32, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ pg1hjm2rmuwg.Addr(), 5,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zdo3zaR5Ihr)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yUQKdFpUvpnk)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(otNY0Rg4)),  /*line :1*/ uintptr(b68qKN),  /*line :1*/ uintptr(a_MrGKcrR), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func GqcjMbi3qoH(zdo3zaR5Ihr L2L8P9WaNZ, cuGEvMbsvH uint32, etRYtnVni *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ c8nLOIMJIsO.Addr(), 3,  /*line :1*/ uintptr(zdo3zaR5Ihr),  /*line :1*/ uintptr(cuGEvMbsvH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func GcqexYW(zdo3zaR5Ihr L2L8P9WaNZ, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cql2kwk.Addr(), 2,  /*line :1*/ uintptr(zdo3zaR5Ihr),  /*line :1*/ uintptr(a_MrGKcrR), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HejiznQiv2(gbpoIWZJB L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ djEhYE.Addr(), 1,  /*line :1*/ uintptr(gbpoIWZJB), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Ucld93(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ d0ecYKvv4.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Gk6OcWvv(qEwymps TaSPPoJMlh, nSYdQwq uint32, qcosdu *NRXOeZVS9w, yuippTj94aE uint32, iMmwVan uint32, xmI3SpwUYF54 *TaSPPoJMlh) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ c_NdFS2aR.Addr(), 6,  /*line :1*/ uintptr(qEwymps),  /*line :1*/ uintptr(nSYdQwq),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qcosdu)),  /*line :1*/ uintptr(yuippTj94aE),  /*line :1*/ uintptr(iMmwVan),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xmI3SpwUYF54)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func XAxEKtY(gbpoIWZJB L2L8P9WaNZ, bjku7YYtC7Q uint32, suyWYd8a0Ll *OHY0gnouiRt6, iudr3o149 uint32, xDvDTK7h9Wm *uint32, c5nI8t *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ jVle5HujZhu.Addr(), 6,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(bjku7YYtC7Q),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(suyWYd8a0Ll)),  /*line :1*/ uintptr(iudr3o149),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xDvDTK7h9Wm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c5nI8t)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SlepUHaIO(auKYZpIGL0UU L2L8P9WaNZ, vaaHS_ntsR uint32, pd9HgcxS6FJ uint32, cEX5t6 uint32, suyWYd8a0Ll *byte, kvo8765 uint32, xDvDTK7h9Wm *uint32, c5nI8t *uint32, eaOGsUE3a *uint32, bue32Tui *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ p_nAnOD.Addr(), 10,  /*line :1*/ uintptr(auKYZpIGL0UU),  /*line :1*/ uintptr(vaaHS_ntsR),  /*line :1*/ uintptr(pd9HgcxS6FJ),  /*line :1*/ uintptr(cEX5t6),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(suyWYd8a0Ll)),  /*line :1*/ uintptr(kvo8765),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xDvDTK7h9Wm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c5nI8t)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eaOGsUE3a)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bue32Tui)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SJNWqa9WP(s6WjuyjSu *He5NCq, fW6sSe *He5NCq) (nyLYwmp bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eeezyOsg.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(s6WjuyjSu)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fW6sSe)), 0)
	nyLYwmp = gPQ4CDCVeA != 0
	return
}

func IRJH2b(ohXw5VDcqx *He5NCq) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ qjRLR1u_A.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)), 0, 0)
	if i_EJOiAVI != 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func C6xEOAk(ohXw5VDcqx *He5NCq) (ajJiI4KvLI uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ keIJ6j.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)), 0, 0)
	ajJiI4KvLI =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func brDxO5WbHd3O(x_GAOsRtS480 string, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D, v2a8VMFGuf0X **He5NCq, bnHAoth **He5NCq, qgpBZb **Na2cliTdMq, lsVxy_e **Na2cliTdMq, gRG9Cm **A2t_Jvj3ECv) (aNqiLgOJxfqs error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, aNqiLgOJxfqs =  /*line :1*/ syscall.GcOmFfsibES(x_GAOsRtS480)
	if aNqiLgOJxfqs != nil {
		return
	}
	return  /*line :1*/ wkAhPNJ6AqZI(gf8OkuRuJg9, mf05G8fOkrTC, d1ypxqWR48, v2a8VMFGuf0X, bnHAoth, qgpBZb, lsVxy_e, gRG9Cm)
}

func wkAhPNJ6AqZI(x_GAOsRtS480 *uint16, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D, v2a8VMFGuf0X **He5NCq, bnHAoth **He5NCq, qgpBZb **Na2cliTdMq, lsVxy_e **Na2cliTdMq, gRG9Cm **A2t_Jvj3ECv) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ dxLHi1t.Addr(), 8,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x_GAOsRtS480)),  /*line :1*/ uintptr(mf05G8fOkrTC),  /*line :1*/ uintptr(d1ypxqWR48),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgpBZb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lsVxy_e)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)), 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func lYFwBD(gRG9Cm *A2t_Jvj3ECv, rZjIu_h *OjSqAgjIy, hKHZ38T275DC *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ lF7Vo3QbA.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rZjIu_h)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hKHZ38T275DC)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func pTSboT(gRG9Cm *A2t_Jvj3ECv, qr3E9HOD *bool, qgpBZb **Na2cliTdMq, acoMd0O *bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if *qr3E9HOD {
		gf8OkuRuJg9 = 1
	}
	var otNQQdF4 uint32
	if *acoMd0O {
		otNQQdF4 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ oqQtKDKgUS.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&gf8OkuRuJg9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgpBZb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&otNQQdF4)), 0, 0)
	*qr3E9HOD = gf8OkuRuJg9 != 0
	*acoMd0O = otNQQdF4 != 0
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func dnmkbYUfgau(gRG9Cm *A2t_Jvj3ECv, bnHAoth **He5NCq, oqGZV8H *bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if *oqGZV8H {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hGsGrxrHtn.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&gf8OkuRuJg9)))
	*oqGZV8H = gf8OkuRuJg9 != 0
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func mu616mhz(gRG9Cm *A2t_Jvj3ECv) (ajJiI4KvLI uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ n0HbqHRwUwQA.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)), 0, 0)
	ajJiI4KvLI =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func f7MPHA2(gRG9Cm *A2t_Jvj3ECv, v2a8VMFGuf0X **He5NCq, uJihx0FgE *bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if *uJihx0FgE {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ vMRHziXX.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&gf8OkuRuJg9)))
	*uJihx0FgE = gf8OkuRuJg9 != 0
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func zDSiag(gRG9Cm *A2t_Jvj3ECv, jsM1XMF *uint8) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bj_QgkJG.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jsM1XMF)), 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func i5Sl6luj(gRG9Cm *A2t_Jvj3ECv, j09Kgx4W *bool, lsVxy_e **Na2cliTdMq, mH1knM *bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if *j09Kgx4W {
		gf8OkuRuJg9 = 1
	}
	var otNQQdF4 uint32
	if *mH1knM {
		otNQQdF4 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ _bwYz1.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&gf8OkuRuJg9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lsVxy_e)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&otNQQdF4)), 0, 0)
	*j09Kgx4W = gf8OkuRuJg9 != 0
	*mH1knM = otNQQdF4 != 0
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func xdBCuREUs9h(iOvctVD26lA L2L8P9WaNZ, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D, v2a8VMFGuf0X **He5NCq, bnHAoth **He5NCq, qgpBZb **Na2cliTdMq, lsVxy_e **Na2cliTdMq, gRG9Cm **A2t_Jvj3ECv) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ vW6cF2tJp.Addr(), 8,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(mf05G8fOkrTC),  /*line :1*/ uintptr(d1ypxqWR48),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgpBZb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lsVxy_e)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)), 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func bSEeAoS6s1DR(ohXw5VDcqx *He5NCq) (qQ1SjOMxTKC *TWV6Elr) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ avHMQVxXM.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)), 0, 0)
	qQ1SjOMxTKC = (* /*line :1*/ TWV6Elr)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	return
}

func jnbf6VRbBCS(ohXw5VDcqx *He5NCq, jhfNujOX uint32) (uYZu5Tp1v *uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gHpdU2.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)),  /*line :1*/ uintptr(jhfNujOX), 0)
	uYZu5Tp1v = (* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	return
}

func a2nVuCjv(ohXw5VDcqx *He5NCq) (rXGaaPBpa *uint8) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ oaox2QLIw.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)), 0, 0)
	rXGaaPBpa = (* /*line :1*/ uint8)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	return
}

func JlkmoPdky432(lAmUDbCC TaSPPoJMlh, jvoafJ uint32, a7Z1LWSG *byte, aaSO5j_bQFr9 uint32, mWcyB2sqYN *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ uBqHfd9.Addr(), 5,  /*line :1*/ uintptr(lAmUDbCC),  /*line :1*/ uintptr(jvoafJ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7Z1LWSG)),  /*line :1*/ uintptr(aaSO5j_bQFr9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mWcyB2sqYN)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func CKsaKpE3(zhKSqgjW uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ iYL04eOVwUN.Addr(), 1,  /*line :1*/ uintptr(zhKSqgjW), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func eyd5z2RaLaA(u6IbXvI *A2t_Jvj3ECv, hKHZ38T275DC uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ rLUH8hPnmpv.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u6IbXvI)),  /*line :1*/ uintptr(hKHZ38T275DC), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JPdpGc(mhANSS4o3m *uint16, wBMMggIijjR *uint16, gegnfSNBFV34 uint32, nlVyZTY1a bool, lk8waH5IGve bool, iV4qRaK uint32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if nlVyZTY1a {
		gf8OkuRuJg9 = 1
	}
	var otNQQdF4 uint32
	if lk8waH5IGve {
		otNQQdF4 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ gAmE9Sf.Addr(), 6,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mhANSS4o3m)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wBMMggIijjR)),  /*line :1*/ uintptr(gegnfSNBFV34),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(otNQQdF4),  /*line :1*/ uintptr(iV4qRaK))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func g2Y5C3WMpl(d6Dv5ak TaSPPoJMlh) (aNqiLgOJxfqs bool, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ vrD_AZf88ah.Addr(), 1,  /*line :1*/ uintptr(d6Dv5ak), 0, 0)
	aNqiLgOJxfqs = gPQ4CDCVeA != 0
	if !aNqiLgOJxfqs {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func wKM4vpBQ(gRG9Cm *A2t_Jvj3ECv) (hHufwxGMdAD bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ugeWaHsC4s.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)), 0, 0)
	hHufwxGMdAD = gPQ4CDCVeA != 0
	return
}

func aUzjOSgDD8a(ohXw5VDcqx *He5NCq) (hHufwxGMdAD bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ vAMA2aKZy.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)), 0, 0)
	hHufwxGMdAD = gPQ4CDCVeA != 0
	return
}

func cKxRf5(ohXw5VDcqx *He5NCq, a4Y5Dlp1iHO CV4yg1b) (qVEdG5C bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aNFHkZYj1a.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)),  /*line :1*/ uintptr(a4Y5Dlp1iHO), 0)
	qVEdG5C = gPQ4CDCVeA != 0
	return
}

func UHIB5J(tAPJDo *uint16, x3YJiQa *uint16, ohXw5VDcqx *He5NCq, a3ih77L1Keu *uint32, uK6oQpNGbjp *uint16, gMQMfbg *uint32, stjB6H_ *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ dm_EfGa.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tAPJDo)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x3YJiQa)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a3ih77L1Keu)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uK6oQpNGbjp)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gMQMfbg)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(stjB6H_)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func LgFpMX8d65(tAPJDo *uint16, ohXw5VDcqx *He5NCq, gYwswmeUjG *uint16, sobmaY *uint32, uK6oQpNGbjp *uint16, gMQMfbg *uint32, stjB6H_ *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ s2XfInzyjf1H.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tAPJDo)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sobmaY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uK6oQpNGbjp)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gMQMfbg)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(stjB6H_)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ApchjUU1y(aqi49jB *uint16, gYwswmeUjG *uint16, oWo2xJHSPh_f *FHc2HnJ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hP4bVgoWV.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aqi49jB)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oWo2xJHSPh_f)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func xdmSfr(c6gWHx *A2t_Jvj3ECv, u6IbXvI *A2t_Jvj3ECv, bOwhozIuAe_H *uint32, qgpBZb *Na2cliTdMq, sgolvRhEI *uint32, lsVxy_e *Na2cliTdMq, gA267ZXVFy *uint32, v2a8VMFGuf0X *He5NCq, eIUgV4FOV *uint32, bnHAoth *He5NCq, miFdEgDW9Y *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ go70DP.Addr(), 11,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c6gWHx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u6IbXvI)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bOwhozIuAe_H)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgpBZb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sgolvRhEI)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lsVxy_e)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gA267ZXVFy)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eIUgV4FOV)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(miFdEgDW9Y)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func pLuyKb(u6IbXvI *A2t_Jvj3ECv, c6gWHx *A2t_Jvj3ECv, oo6I0d *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xK0citfiso_u.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u6IbXvI)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c6gWHx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oo6I0d)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Au3QEpaPn1(gbpoIWZJB L2L8P9WaNZ, fM256Jz uint32, n_KkZfY *S5pK5u) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gGbVEF6B.Addr(), 3,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(fM256Jz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(n_KkZfY)))
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func RHTKRt(a_VYqwS4m2B L2L8P9WaNZ, oW8M6Qq uint32, lAmUDbCC *TaSPPoJMlh) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ tHa02a1GeCW.Addr(), 3,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(oW8M6Qq),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lAmUDbCC)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func XahPN4XBLM(mhANSS4o3m *uint16, c8VJZ0sw *uint16, oW8M6Qq uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jKB3JzepJ.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mhANSS4o3m)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c8VJZ0sw)),  /*line :1*/ uintptr(oW8M6Qq))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HD3C56_ffb(auKYZpIGL0UU L2L8P9WaNZ, hdGT65m *uint16, oW8M6Qq uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kv0g00muN.Addr(), 3,  /*line :1*/ uintptr(auKYZpIGL0UU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hdGT65m)),  /*line :1*/ uintptr(oW8M6Qq))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DlbGOmZ66H(q0XnBF4Z L2L8P9WaNZ, oW8M6Qq uint32, fpKYAgpXw4A bool, lAmUDbCC *TaSPPoJMlh) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if fpKYAgpXw4A {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ j41UBMeaVV.Addr(), 4,  /*line :1*/ uintptr(q0XnBF4Z),  /*line :1*/ uintptr(oW8M6Qq),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lAmUDbCC)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func C0xFxCAqaL(gbpoIWZJB L2L8P9WaNZ, vaaHS_ntsR uint32, xVpZawxq *byte, iudr3o149 uint32, xDvDTK7h9Wm *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dDTNWYBO.Addr(), 5,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(vaaHS_ntsR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xVpZawxq)),  /*line :1*/ uintptr(iudr3o149),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xDvDTK7h9Wm)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HRgXuy4s(gbpoIWZJB L2L8P9WaNZ, oLg0YtZWp0 *FJRpa4YpK, kvo8765 uint32, xDvDTK7h9Wm *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ etHz3XK.Addr(), 4,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oLg0YtZWp0)),  /*line :1*/ uintptr(kvo8765),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xDvDTK7h9Wm)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SrVr__(gbpoIWZJB L2L8P9WaNZ, vaaHS_ntsR uint32, hWSoLk unsafe.Pointer) (jeWMpOaQtWV error) {
	jeWMpOaQtWV =  /*line :1*/ zJFCJzhIRSx.Find()
	if jeWMpOaQtWV != nil {
		return
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zJFCJzhIRSx.Addr(), 3,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(vaaHS_ntsR),  /*line :1*/ uintptr(hWSoLk))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func UxEb4rHCas5q(auKYZpIGL0UU L2L8P9WaNZ, e3HuI9Q *XoQ8Jtg, kvo8765 uint32, xDvDTK7h9Wm *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ kGmliiXuyeu.Addr(), 4,  /*line :1*/ uintptr(auKYZpIGL0UU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(e3HuI9Q)),  /*line :1*/ uintptr(kvo8765),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xDvDTK7h9Wm)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Z5oa6kTovB(gbpoIWZJB L2L8P9WaNZ, oGbyaObSkJ *YhbjUYS0) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ pb2XaHC.Addr(), 2,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oGbyaObSkJ)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func T37QFHvc(gbpoIWZJB L2L8P9WaNZ, vaaHS_ntsR uint32, xVpZawxq *byte, iudr3o149 uint32, xDvDTK7h9Wm *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ gFDHk5H.Addr(), 5,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(vaaHS_ntsR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xVpZawxq)),  /*line :1*/ uintptr(iudr3o149),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xDvDTK7h9Wm)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func V4TxUz(hGZWAKz L2L8P9WaNZ) (cDWItDxhE4G error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ e90awnRjK9.Addr(), 1,  /*line :1*/ uintptr(hGZWAKz), 0, 0)
	if gPQ4CDCVeA != 0 {
		cDWItDxhE4G =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func KcjB5EPCSXM(hGZWAKz L2L8P9WaNZ, jhfNujOX uint32, gYwswmeUjG *uint16, sobmaY *uint32, yv8vYCF3 *uint32, jzcRahcgUWi2 *uint16, f978WGi16 *uint32, jQ9OgZwZ *ZPa9KL2) (cDWItDxhE4G error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ bQ2WMBaH.Addr(), 8,  /*line :1*/ uintptr(hGZWAKz),  /*line :1*/ uintptr(jhfNujOX),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sobmaY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yv8vYCF3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jzcRahcgUWi2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(f978WGi16)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jQ9OgZwZ)), 0)
	if gPQ4CDCVeA != 0 {
		cDWItDxhE4G =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func YwixW3Ecjf(hGZWAKz L2L8P9WaNZ, jfbAuSbvU bool, kL6BYHuP uint32, g_3oYJh6F L2L8P9WaNZ, g0VFUX2g5 bool) (cDWItDxhE4G error) {
	var gf8OkuRuJg9 uint32
	if jfbAuSbvU {
		gf8OkuRuJg9 = 1
	}
	var otNQQdF4 uint32
	if g0VFUX2g5 {
		otNQQdF4 = 1
	}
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ cb9WQ0wtxY.Addr(), 5,  /*line :1*/ uintptr(hGZWAKz),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(kL6BYHuP),  /*line :1*/ uintptr(g_3oYJh6F),  /*line :1*/ uintptr(otNQQdF4), 0)
	if gPQ4CDCVeA != 0 {
		cDWItDxhE4G =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func IoUnFI(hGZWAKz L2L8P9WaNZ, p5cr5oj_DwR *uint16, xPmyKOJbMp1Q uint32, nSYdQwq uint32, fz1BsYsTSAQ *L2L8P9WaNZ) (cDWItDxhE4G error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ jlYF3eJWaPK.Addr(), 5,  /*line :1*/ uintptr(hGZWAKz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(p5cr5oj_DwR)),  /*line :1*/ uintptr(xPmyKOJbMp1Q),  /*line :1*/ uintptr(nSYdQwq),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fz1BsYsTSAQ)), 0)
	if gPQ4CDCVeA != 0 {
		cDWItDxhE4G =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func RoPARgby3i4(hGZWAKz L2L8P9WaNZ, jzcRahcgUWi2 *uint16, f978WGi16 *uint32, yv8vYCF3 *uint32, k3lLEHhH *uint32, _YWaQUT *uint32, eJzcIM *uint32, r7M37VBvRPI0 *uint32, yyvxJn0Hc *uint32, adrqC5e1dk *uint32, eIzjsnd10 *uint32, jQ9OgZwZ *ZPa9KL2) (cDWItDxhE4G error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ vpgTqGRBg.Addr(), 12,  /*line :1*/ uintptr(hGZWAKz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jzcRahcgUWi2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(f978WGi16)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yv8vYCF3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(k3lLEHhH)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_YWaQUT)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eJzcIM)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(r7M37VBvRPI0)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yyvxJn0Hc)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(adrqC5e1dk)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eIzjsnd10)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jQ9OgZwZ)))
	if gPQ4CDCVeA != 0 {
		cDWItDxhE4G =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func V6zBt_s(hGZWAKz L2L8P9WaNZ, gYwswmeUjG *uint16, yv8vYCF3 *uint32, mrwNYyas *uint32, etRYtnVni *byte, cuGEvMbsvH *uint32) (cDWItDxhE4G error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ mh2e13t8abh.Addr(), 6,  /*line :1*/ uintptr(hGZWAKz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yv8vYCF3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mrwNYyas)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cuGEvMbsvH)))
	if gPQ4CDCVeA != 0 {
		cDWItDxhE4G =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func DDgWTMolS_a(iEADECSFPnf *uint16, oy8AUYZcIR *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ vowp8C8.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iEADECSFPnf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oy8AUYZcIR)), 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Ibk8tcEFB(hdGT65m *uint16, l7qnoBn_0md1 uintptr, _K8SMQVHOjji uintptr) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ iD7x4md.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hdGT65m)),  /*line :1*/ uintptr(l7qnoBn_0md1),  /*line :1*/ uintptr(_K8SMQVHOjji))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RQ340Km40vI_(qzPXpSS L2L8P9WaNZ, jrwexlU7o uint16, tmG2NdyZlf uint16, uLL_eSWZ uint32, dZ4jsXro8jTI uintptr, pns2Cu uint16, ewGci2xd0fT uint32, iIkKYsC **uint16, rdN6wYCMKo *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ aZaUnkxaI.Addr(), 9,  /*line :1*/ uintptr(qzPXpSS),  /*line :1*/ uintptr(jrwexlU7o),  /*line :1*/ uintptr(tmG2NdyZlf),  /*line :1*/ uintptr(uLL_eSWZ),  /*line :1*/ uintptr(dZ4jsXro8jTI),  /*line :1*/ uintptr(pns2Cu),  /*line :1*/ uintptr(ewGci2xd0fT),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iIkKYsC)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rdN6wYCMKo)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func O5XzzRZ() (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ qUzKKN_.Addr(), 0, 0, 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func qhb0Czy(ejNzXae uint32, siNJ3lw3z *GK5XcbVlLR, h4BtPa84R *Na2cliTdMq, dd19cvUzhf **Na2cliTdMq) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ lHfFxw2FC.Addr(), 4,  /*line :1*/ uintptr(ejNzXae),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(siNJ3lw3z)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(h4BtPa84R)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dd19cvUzhf)), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func JDZdWYLzW(iOvctVD26lA L2L8P9WaNZ, d1ypxqWR48 B4w2t8D, aTafZskn *A2t_Jvj3ECv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ i9on9y.Addr(), 3,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(d1ypxqWR48),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aTafZskn)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func C9M0zmX(x_GAOsRtS480 string, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D, v2a8VMFGuf0X *He5NCq, bnHAoth *He5NCq, qgpBZb *Na2cliTdMq, lsVxy_e *Na2cliTdMq) (aNqiLgOJxfqs error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, aNqiLgOJxfqs =  /*line :1*/ syscall.GcOmFfsibES(x_GAOsRtS480)
	if aNqiLgOJxfqs != nil {
		return
	}
	return  /*line :1*/ xaxdlaZcPiw(gf8OkuRuJg9, mf05G8fOkrTC, d1ypxqWR48, v2a8VMFGuf0X, bnHAoth, qgpBZb, lsVxy_e)
}

func xaxdlaZcPiw(x_GAOsRtS480 *uint16, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D, v2a8VMFGuf0X *He5NCq, bnHAoth *He5NCq, qgpBZb *Na2cliTdMq, lsVxy_e *Na2cliTdMq) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ cmJ0sPauhG9q.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x_GAOsRtS480)),  /*line :1*/ uintptr(mf05G8fOkrTC),  /*line :1*/ uintptr(d1ypxqWR48),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgpBZb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lsVxy_e)), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func pg_dOfgc(gRG9Cm *A2t_Jvj3ECv, fKPnmQ_w OjSqAgjIy, lJCNaye74 OjSqAgjIy) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ orWzIs9FB.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr(fKPnmQ_w),  /*line :1*/ uintptr(lJCNaye74))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aGxMC3J(gRG9Cm *A2t_Jvj3ECv, qr3E9HOD bool, qgpBZb *Na2cliTdMq, acoMd0O bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if qr3E9HOD {
		gf8OkuRuJg9 = 1
	}
	var otNQQdF4 uint32
	if acoMd0O {
		otNQQdF4 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ ajPxiSSw0tnI.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgpBZb)),  /*line :1*/ uintptr(otNQQdF4), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func b04HqSxM3(gRG9Cm *A2t_Jvj3ECv, bnHAoth *He5NCq, oqGZV8H bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if oqGZV8H {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ mxY4SkMLjpsU.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr(gf8OkuRuJg9))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aTIded(gRG9Cm *A2t_Jvj3ECv, v2a8VMFGuf0X *He5NCq, uJihx0FgE bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if uJihx0FgE {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ yBEYlnXato35.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr(gf8OkuRuJg9))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func cANESp5i1CR(gRG9Cm *A2t_Jvj3ECv, jsM1XMF *uint8) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ arjuCqav.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jsM1XMF)), 0)
	return
}

func vCFgaK5rwa7(gRG9Cm *A2t_Jvj3ECv, j09Kgx4W bool, lsVxy_e *Na2cliTdMq, mH1knM bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if j09Kgx4W {
		gf8OkuRuJg9 = 1
	}
	var otNQQdF4 uint32
	if mH1knM {
		otNQQdF4 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ wdnaaUk.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gRG9Cm)),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lsVxy_e)),  /*line :1*/ uintptr(otNQQdF4), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HwaFN2kSDMrW(iOvctVD26lA L2L8P9WaNZ, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D, v2a8VMFGuf0X *He5NCq, bnHAoth *He5NCq, qgpBZb *Na2cliTdMq, lsVxy_e *Na2cliTdMq) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ zXhtiaLhV.Addr(), 7,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(mf05G8fOkrTC),  /*line :1*/ uintptr(d1ypxqWR48),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v2a8VMFGuf0X)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bnHAoth)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgpBZb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lsVxy_e)), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func KU0TL7NJGZMH(gbpoIWZJB L2L8P9WaNZ, gznaDa *YhbjUYS0) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ woc2sm4.Addr(), 2,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gznaDa)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DUCsHE9o6Cp0(q0XnBF4Z *L2L8P9WaNZ, lAmUDbCC TaSPPoJMlh) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hU0uN4LpB0d.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(q0XnBF4Z)),  /*line :1*/ uintptr(lAmUDbCC), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HeC0mWJGMr(lAmUDbCC TaSPPoJMlh, jvoafJ uint32, a7Z1LWSG *byte, aaSO5j_bQFr9 uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ g0O7MBn.Addr(), 4,  /*line :1*/ uintptr(lAmUDbCC),  /*line :1*/ uintptr(jvoafJ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7Z1LWSG)),  /*line :1*/ uintptr(aaSO5j_bQFr9), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FoQIZhYREjt(cVTVOzk6l *MSSmb3ckxDlt) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ uGtcfa2YGM.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cVTVOzk6l)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EDH6hL7Xooze(gbpoIWZJB L2L8P9WaNZ, l5qxO1yW uint32, haLZWB7jI0Rs **uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ uqaI0Vx.Addr(), 3,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(l5qxO1yW),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(haLZWB7jI0Rs)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func D0TumGp2(f2FOcgf7 L2L8P9WaNZ, bPKl1NSOj0wH *U_27iTvK, cElNQtAsUh uint32, gfO4fzk **U_27iTvK) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ pVLZEgaaqCvT.Addr(), 4,  /*line :1*/ uintptr(f2FOcgf7),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bPKl1NSOj0wH)),  /*line :1*/ uintptr(cElNQtAsUh),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gfO4fzk)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Nu7_XE3o8(f2FOcgf7 L2L8P9WaNZ, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ f1WG2ds.Addr(), 2,  /*line :1*/ uintptr(f2FOcgf7),  /*line :1*/ uintptr(a_MrGKcrR), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Fui79DIcSHp(ws81aB5eKg uint32, joOm7kfmQYVD *byte, v4SSjoo7 uint32) (_K8SMQVHOjji *U_27iTvK, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ iZ6BTc2715.Addr(), 3,  /*line :1*/ uintptr(ws81aB5eKg),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(joOm7kfmQYVD)),  /*line :1*/ uintptr(v4SSjoo7))
	_K8SMQVHOjji = (* /*line :1*/ U_27iTvK)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if _K8SMQVHOjji == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Am0GzVAvCx5S(bPKl1NSOj0wH *U_27iTvK) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ uefQ0653B7Q.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bPKl1NSOj0wH)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Ag4Pn5(bPKl1NSOj0wH *U_27iTvK) (fqCPSSDa *U_27iTvK) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jaSFAnO.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bPKl1NSOj0wH)), 0, 0)
	fqCPSSDa = (* /*line :1*/ U_27iTvK)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	return
}

func Kr7bGBRfbnIu(f2FOcgf7 L2L8P9WaNZ, iULwl1hemg *U_27iTvK) (_K8SMQVHOjji *U_27iTvK, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ elioVh8.Addr(), 2,  /*line :1*/ uintptr(f2FOcgf7),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iULwl1hemg)), 0)
	_K8SMQVHOjji = (* /*line :1*/ U_27iTvK)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if _K8SMQVHOjji == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func OWpN6NqU(f2FOcgf7 L2L8P9WaNZ, ws81aB5eKg uint32, c0OQSvBXf uint32, ig_isa8VaT5R uint32, hGcJjpIlra unsafe.Pointer, gcRFbwy *U_27iTvK) (eSYPff *U_27iTvK, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ cwp59dU_C4DP.Addr(), 6,  /*line :1*/ uintptr(f2FOcgf7),  /*line :1*/ uintptr(ws81aB5eKg),  /*line :1*/ uintptr(c0OQSvBXf),  /*line :1*/ uintptr(ig_isa8VaT5R),  /*line :1*/ uintptr(hGcJjpIlra),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gcRFbwy)))
	eSYPff = (* /*line :1*/ U_27iTvK)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if eSYPff == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HrzN9rNTap(f2FOcgf7 L2L8P9WaNZ, ws81aB5eKg uint32, c0OQSvBXf uint32, ig_isa8VaT5R uint32, hGcJjpIlra unsafe.Pointer, eDKnrKpl *Iisel0J) (hBFChSJ *Iisel0J, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ zofSVx1U.Addr(), 6,  /*line :1*/ uintptr(f2FOcgf7),  /*line :1*/ uintptr(ws81aB5eKg),  /*line :1*/ uintptr(c0OQSvBXf),  /*line :1*/ uintptr(ig_isa8VaT5R),  /*line :1*/ uintptr(hGcJjpIlra),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eDKnrKpl)))
	hBFChSJ = (* /*line :1*/ Iisel0J)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if hBFChSJ == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZDH82e(x0wmbX *byte, huQ9jWuAi uint32, q1S7FkBQiCWM *HDSX6oPJk_dJ) (aNqiLgOJxfqs *HDSX6oPJk_dJ) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ofyKNlEiNQ.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x0wmbX)),  /*line :1*/ uintptr(huQ9jWuAi),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(q1S7FkBQiCWM)))
	aNqiLgOJxfqs = (* /*line :1*/ HDSX6oPJk_dJ)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	return
}

func ZuXqlkpKT90(eJPcX9 *Iisel0J) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ es9AG6kX5eUb.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eJPcX9)), 0, 0)
	return
}

func V6017uj(eJPcX9 *U_27iTvK) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hR_xTjndo1Op.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eJPcX9)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func N3AqP0Xp(qIiwKb L2L8P9WaNZ, jYr3Uf_zY *U_27iTvK, esI00UeVTD *ZPa9KL2, lRswd0Co L2L8P9WaNZ, dq_RzN *PANRStxqVBD, a_MrGKcrR uint32, yv8vYCF3 uintptr, eYqk_mv **Iisel0J) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ zz4LZpa.Addr(), 8,  /*line :1*/ uintptr(qIiwKb),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jYr3Uf_zY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(esI00UeVTD)),  /*line :1*/ uintptr(lRswd0Co),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dq_RzN)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eYqk_mv)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Q1JtIUda6GRk(bPKl1NSOj0wH *U_27iTvK, h_brS9a uint32, a_MrGKcrR uint32, syxR6Y1 unsafe.Pointer, gYwswmeUjG *uint16, nFtsmqHC uint32) (x7J8HIt uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dqk1LIahL.Addr(), 6,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bPKl1NSOj0wH)),  /*line :1*/ uintptr(h_brS9a),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(syxR6Y1),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(nFtsmqHC))
	x7J8HIt =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func XDVw_yZ5xW0(bC2Kfv uintptr, yzLDDnd uint32, hdtAU9h8 uintptr, a_MrGKcrR uint32, dq_RzN uintptr) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ aCKPLJWbT.Addr(), 5,  /*line :1*/ uintptr(bC2Kfv),  /*line :1*/ uintptr(yzLDDnd),  /*line :1*/ uintptr(hdtAU9h8),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(dq_RzN), 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func KeZakevb(iwOihq1 L2L8P9WaNZ, gYwswmeUjG *uint16) (f2FOcgf7 L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xERaMMGOivl.Addr(), 2,  /*line :1*/ uintptr(iwOihq1),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)), 0)
	f2FOcgf7 =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if f2FOcgf7 == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FAD3q_(oeDQqGStql uintptr, gSHbeyGuqwhD *Iisel0J, dq_RzN *Q3PNbtb9v, oGbyaObSkJ *QPPLgM4k) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ r4lpHHGKD.Addr(), 4,  /*line :1*/ uintptr(oeDQqGStql),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gSHbeyGuqwhD)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dq_RzN)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oGbyaObSkJ)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func LJY0UaadU(eSYPff *U_27iTvK, a_MrGKcrR uint32, h0k9dwHyD unsafe.Pointer, qkfbHWtd8R *L2L8P9WaNZ, pHJsdzaAmW *uint32, faHCUcfLn_ *bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if *faHCUcfLn_ {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ pMCa3SCIoqp.Addr(), 6,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eSYPff)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(h0k9dwHyD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qkfbHWtd8R)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pHJsdzaAmW)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&gf8OkuRuJg9)))
	*faHCUcfLn_ = gf8OkuRuJg9 != 0
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func PXDqfPkaQcAI(blaYH9aZto uint32, yy4i8G *byte, sdHmsOW1 *byte, eT1UQSvR_nM uint32, a_MrGKcrR uint32, obTU8wso unsafe.Pointer, crEKaIUBMN8B *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ g29_mRR6.Addr(), 7,  /*line :1*/ uintptr(blaYH9aZto),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yy4i8G)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sdHmsOW1)),  /*line :1*/ uintptr(eT1UQSvR_nM),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(obTU8wso),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(crEKaIUBMN8B)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func CGIhWd(uqHuoiZHy *I3Px1aVqtZ, gYwswmeUjG *uint16, t14AdIGd *I3Px1aVqtZ, yv8vYCF3 uintptr, lmV6zVUfWaZu *TOwv2p3Pxya, a_MrGKcrR uint32, bHHRL6AjnUaG *I3Px1aVqtZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ gxkyPGnJiO.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uqHuoiZHy)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t14AdIGd)),  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lmV6zVUfWaZu)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bHHRL6AjnUaG)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DeSChlwLPaJ(mf05G8fOkrTC uint32, apTKox9MAtyw unsafe.Pointer, uN0iSdccA1l uint32, jUZma2HcDZ7 uint32, a_MrGKcrR uint32, yzLDDnd *uint32, jmv9j4_h *uint32, iEWBOMGEDU *uint32, _Btsq53Zo *L2L8P9WaNZ, qWqM5J9x27 *L2L8P9WaNZ, _K8SMQVHOjji *unsafe.Pointer) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ dupqPJx.Addr(), 11,  /*line :1*/ uintptr(mf05G8fOkrTC),  /*line :1*/ uintptr(apTKox9MAtyw),  /*line :1*/ uintptr(uN0iSdccA1l),  /*line :1*/ uintptr(jUZma2HcDZ7),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yzLDDnd)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jmv9j4_h)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iEWBOMGEDU)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_Btsq53Zo)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qWqM5J9x27)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_K8SMQVHOjji)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func OkO2YjTMJgRN(uqHuoiZHy *I3Px1aVqtZ, gYwswmeUjG **uint16, t14AdIGd *I3Px1aVqtZ, yv8vYCF3 uintptr, lmV6zVUfWaZu *TOwv2p3Pxya, a_MrGKcrR uint32, bHHRL6AjnUaG *I3Px1aVqtZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ zCYV5nU.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uqHuoiZHy)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t14AdIGd)),  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lmV6zVUfWaZu)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bHHRL6AjnUaG)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Zxp55k(h4LDSS11qFTP *ARhDwxk, jhjxpnHouRNO *uint16, a_MrGKcrR uint32) (f2FOcgf7 L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ j6asE6.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(h4LDSS11qFTP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jhjxpnHouRNO)),  /*line :1*/ uintptr(a_MrGKcrR))
	f2FOcgf7 =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if f2FOcgf7 == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func H4d0zG(nqESlJ *uint16, jHaqWtgLX *uint16) (e0bqkn bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ snRiy6niCa6u.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nqESlJ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jHaqWtgLX)), 0)
	e0bqkn = gPQ4CDCVeA != 0
	return
}

func A888vI(gYwswmeUjG string, aYxhl2G2hwf uint16, xPmyKOJbMp1Q uint32, t_o62b *byte, moHMBKU **AFQqGzn5iwE, zWD5mpsEd *byte) (oGbyaObSkJ error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, oGbyaObSkJ =  /*line :1*/ syscall.GcOmFfsibES(gYwswmeUjG)
	if oGbyaObSkJ != nil {
		return
	}
	return  /*line :1*/ aG7buq2(gf8OkuRuJg9, aYxhl2G2hwf, xPmyKOJbMp1Q, t_o62b, moHMBKU, zWD5mpsEd)
}

func aG7buq2(gYwswmeUjG *uint16, aYxhl2G2hwf uint16, xPmyKOJbMp1Q uint32, t_o62b *byte, moHMBKU **AFQqGzn5iwE, zWD5mpsEd *byte) (oGbyaObSkJ error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ vOLc6_.Addr(), 6,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(aYxhl2G2hwf),  /*line :1*/ uintptr(xPmyKOJbMp1Q),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t_o62b)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(moHMBKU)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zWD5mpsEd)))
	if gPQ4CDCVeA != 0 {
		oGbyaObSkJ =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func RYWyoFC6oZML(lFmzLo *AFQqGzn5iwE, jqjHTZwMAoon uint32) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ j7fpugNwoH.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lFmzLo)),  /*line :1*/ uintptr(jqjHTZwMAoon), 0)
	return
}

func HZ4dJU5(i6Cvqni_Es IJltG1D5QG, q5eWI0Z1lHex uint32, wvZBcfh unsafe.Pointer, nFtsmqHC uint32) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ deA1m_XBF.Addr(), 4,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr(q5eWI0Z1lHex),  /*line :1*/ uintptr(wvZBcfh),  /*line :1*/ uintptr(nFtsmqHC), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func HGZgzzRP(i6Cvqni_Es IJltG1D5QG, q5eWI0Z1lHex uint32, wvZBcfh unsafe.Pointer, nFtsmqHC uint32) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ xdAmltl87PC8.Addr(), 4,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr(q5eWI0Z1lHex),  /*line :1*/ uintptr(wvZBcfh),  /*line :1*/ uintptr(nFtsmqHC), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func IkMRELF9S3(jroU3W5g uint32, a_MrGKcrR uint32, yv8vYCF3 uintptr, rpEouRY1 *EMo4WzV, gs_AwUeR *uint32) (i6JESw error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ ekuMhpIr.Addr(), 5,  /*line :1*/ uintptr(jroU3W5g),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpEouRY1)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gs_AwUeR)), 0)
	if gPQ4CDCVeA != 0 {
		i6JESw =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func M6OrujkIa(trWHZIqHa0R *T73ZsZ7PFJNo, tnc3AG1jv *uint32) (i6JESw error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cRWDyrzaEqc.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(trWHZIqHa0R)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tnc3AG1jv)), 0)
	if gPQ4CDCVeA != 0 {
		i6JESw =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func u9EM3rC(uaMPM0 unsafe.Pointer, y9xEK8 *uint32) (i6JESw error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dtAjS4EaGjIZ.Addr(), 2,  /*line :1*/ uintptr(uaMPM0),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y9xEK8)), 0)
	if gPQ4CDCVeA != 0 {
		i6JESw =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func ZWv5IB(gniIjzTPcr *PfkJPslS) (i6JESw error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ icNFZC.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gniIjzTPcr)), 0, 0)
	if gPQ4CDCVeA != 0 {
		i6JESw =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func Z0E7tmSAe4(bZKbGTxeit *uint16) (pmQD4ZWSkT15 uintptr, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eWnhtWtA.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)), 0, 0)
	pmQD4ZWSkT15 =  /*line :1*/ uintptr(gPQ4CDCVeA)
	if pmQD4ZWSkT15 == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AD15kcGCq3(panFWjsK L2L8P9WaNZ, a_VYqwS4m2B L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ fxschXNKn0.Addr(), 2,  /*line :1*/ uintptr(panFWjsK),  /*line :1*/ uintptr(a_VYqwS4m2B), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FMdrP2AhtxS(bamc83qA3DBr L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ wJ6bbteacq.Addr(), 1,  /*line :1*/ uintptr(bamc83qA3DBr), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func V3K_paoNTS(bamc83qA3DBr L2L8P9WaNZ, mf_Gq_ *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ggsdamqooaG.Addr(), 2,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mf_Gq_)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func E5j7kDrdt(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nX0nVIzBFv.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FiHIN8cJ(rprJNy L2L8P9WaNZ) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nyoqbTJ4lL.Addr(), 1,  /*line :1*/ uintptr(rprJNy), 0, 0)
	return
}

func GbCU7h(gSssbc L2L8P9WaNZ, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zmTCFf.Addr(), 2,  /*line :1*/ uintptr(gSssbc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZGfr8zMaI(bZKbGTxeit *uint16, aRcafgWyH *NRXOeZVS9w) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ qUhxapRH9HyE.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aRcafgWyH)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IsaUdum(ijrEca1qAJ *NRXOeZVS9w, gYwswmeUjG *uint16, a_MrGKcrR uint32, nSYdQwq uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ vUiNy0xE.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ijrEca1qAJ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(nSYdQwq), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 || j4dSamRp == ERROR_ALREADY_EXISTS {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JNUEuv1v(ijrEca1qAJ *NRXOeZVS9w, bnethyOfQAhR uint32, yisXobwORn uint32, gYwswmeUjG *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ aOGLc61ak8TD.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ijrEca1qAJ)),  /*line :1*/ uintptr(bnethyOfQAhR),  /*line :1*/ uintptr(yisXobwORn),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 || j4dSamRp == ERROR_ALREADY_EXISTS {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func XRHGJGafy(gDrzq4 L2L8P9WaNZ, aRcafgWyH *NRXOeZVS9w, fNx5rkgh uint32, blEaKK uint32, reWaVI4 uint32, gYwswmeUjG *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ xPiI_EiWa.Addr(), 6,  /*line :1*/ uintptr(gDrzq4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aRcafgWyH)),  /*line :1*/ uintptr(fNx5rkgh),  /*line :1*/ uintptr(blEaKK),  /*line :1*/ uintptr(reWaVI4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 || j4dSamRp == ERROR_ALREADY_EXISTS {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func LNhQWvTnm6(gYwswmeUjG *uint16, oW8M6Qq uint32, yAZay4eEB0M uint32, aRcafgWyH *NRXOeZVS9w, svCwoQOx8aBm uint32, jQSvSZ uint32, aJL1KPvY L2L8P9WaNZ) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ d3nYINs8.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(oW8M6Qq),  /*line :1*/ uintptr(yAZay4eEB0M),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aRcafgWyH)),  /*line :1*/ uintptr(svCwoQOx8aBm),  /*line :1*/ uintptr(jQSvSZ),  /*line :1*/ uintptr(aJL1KPvY), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FPRiPhjaJr(ybpEJar *uint16, fctb4fB *uint16, yv8vYCF3 uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jkSVytdeG.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ybpEJar)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fctb4fB)),  /*line :1*/ uintptr(yv8vYCF3))
	if i_EJOiAVI&0xff == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AjgCB5Pytyh(r9HmgAa_zUgV L2L8P9WaNZ, qAYnXF L2L8P9WaNZ, hGZWAKz uintptr, lC9S9g4VAdC uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ e34iRD1FKD.Addr(), 4,  /*line :1*/ uintptr(r9HmgAa_zUgV),  /*line :1*/ uintptr(qAYnXF),  /*line :1*/ uintptr(hGZWAKz),  /*line :1*/ uintptr(lC9S9g4VAdC), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZFMElRYkYc(cA1RWCG6jiut *NRXOeZVS9w, gYwswmeUjG *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ fE4k5ITrJ8p.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cA1RWCG6jiut)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)), 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func PENEsyLds_6(eUBGZQn *NRXOeZVS9w, gYwswmeUjG *uint16, a_MrGKcrR uint32, nSYdQwq uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ tUa2sbtYfL.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eUBGZQn)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(nSYdQwq), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 || j4dSamRp == ERROR_ALREADY_EXISTS {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RVZvnLN(eUBGZQn *NRXOeZVS9w, eb5Nhc9 bool, gYwswmeUjG *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if eb5Nhc9 {
		gf8OkuRuJg9 = 1
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ifY4A0t.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eUBGZQn)),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 || j4dSamRp == ERROR_ALREADY_EXISTS {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZxzQqJiFV(gYwswmeUjG *uint16, a_MrGKcrR uint32, c1fU74 uint32, ua_C6vfqvhe uint32, fvJPV0 uint32, rH6FpU9oKLxf uint32, eR28AyUtu_I9 uint32, aRcafgWyH *NRXOeZVS9w) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ eaGh7A0cl8.Addr(), 8,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(c1fU74),  /*line :1*/ uintptr(ua_C6vfqvhe),  /*line :1*/ uintptr(fvJPV0),  /*line :1*/ uintptr(rH6FpU9oKLxf),  /*line :1*/ uintptr(eR28AyUtu_I9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aRcafgWyH)), 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func YwbL9p6NalN(faVPa4gfw2 *L2L8P9WaNZ, fiINOY *L2L8P9WaNZ, aRcafgWyH *NRXOeZVS9w, nFtsmqHC uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ adiTo3avuaY.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(faVPa4gfw2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fiINOY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aRcafgWyH)),  /*line :1*/ uintptr(nFtsmqHC), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func UMmvD4U(vsoaZFn4bg *uint16, zXn7QsQIeo *uint16, jbDCb5 *NRXOeZVS9w, m4ah36Dc *NRXOeZVS9w, y0Izm0 bool, elKoAUa0KJ uint32, uG96gcyOx *uint16, ay4_8porhppA *uint16, hnFFk_ *KqXK1jwYl5VP, s301dBC_ *MMPS3iJAuV) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if y0Izm0 {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ lVSV5vqazU.Addr(), 10,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vsoaZFn4bg)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zXn7QsQIeo)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jbDCb5)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m4ah36Dc)),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(elKoAUa0KJ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uG96gcyOx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ay4_8porhppA)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hnFFk_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(s301dBC_)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func liWXWGwW(nFtsmqHC uint32, l8eN3VDi L2L8P9WaNZ, icEgEFtX_G0 L2L8P9WaNZ, a_MrGKcrR uint32, zaIN6WS *L2L8P9WaNZ) (eh2rtH05b2x error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ e1a3puKk.Addr(), 5,  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr(l8eN3VDi),  /*line :1*/ uintptr(icEgEFtX_G0),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zaIN6WS)), 0)
	if gPQ4CDCVeA != 0 {
		eh2rtH05b2x =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func Bv9SNhlJV4(ib5HuEjHV *uint16, xm8iFdrQmRp *uint16, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ sax1PV7.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ib5HuEjHV)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xm8iFdrQmRp)),  /*line :1*/ uintptr(a_MrGKcrR))
	if i_EJOiAVI&0xff == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func L8wfFX3s(a_MrGKcrR uint32, cMFIuNaB uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ faa40DDfaMjn.Addr(), 2,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(cMFIuNaB), 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SLoW9ADwbHc(a_MrGKcrR uint32, qNhx_SU2 *uint16, jORjJnF *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ lC6c9Dwoo.Addr(), 3,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qNhx_SU2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jORjJnF)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func F9XQJe2wt5oP(bZKbGTxeit *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xPNrYXIevNYK.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func dOpVlY(fIzbyI *JEfkblUWi) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ v8RawuG_2CF5.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fIzbyI)), 0, 0)
	return
}

func JBFS89xxvNE4(zNEAk77a *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dderR8bg5o.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zNEAk77a)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HjWTjVjfKi(iOvctVD26lA L2L8P9WaNZ, kfsn0sfYy6F uint32, jQqGgBVSKXI5 *byte, ca4jZgG uint32, pB77mVqN *byte, evmqC_B9W uint32, pwkETEB8SLrs *uint32, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ miyQrEWR11qt.Addr(), 8,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(kfsn0sfYy6F),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jQqGgBVSKXI5)),  /*line :1*/ uintptr(ca4jZgG),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pB77mVqN)),  /*line :1*/ uintptr(evmqC_B9W),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pwkETEB8SLrs)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func GbNhjqx(jOy1jicRGf L2L8P9WaNZ, p5ktqlKUUDit L2L8P9WaNZ, hNtSzWaJ9 L2L8P9WaNZ, zxjCVn0DkNvt *L2L8P9WaNZ, pUM74L7FmfP uint32, ownpHjrL6 bool, f7qh7o6hgc3 uint32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if ownpHjrL6 {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ sBVaB4NKsX.Addr(), 7,  /*line :1*/ uintptr(jOy1jicRGf),  /*line :1*/ uintptr(p5ktqlKUUDit),  /*line :1*/ uintptr(hNtSzWaJ9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zxjCVn0DkNvt)),  /*line :1*/ uintptr(pUM74L7FmfP),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(f7qh7o6hgc3), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JGEMyXm8(tmIM88dtF uint32) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ pMDx7OiSSrKK.Addr(), 1,  /*line :1*/ uintptr(tmIM88dtF), 0, 0)
	return
}

func TuqYU1ro(huthvVhTPAIF *uint16, z0twykel *uint16, nFtsmqHC uint32) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ sqCtIWBrE.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(huthvVhTPAIF)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(z0twykel)),  /*line :1*/ uintptr(nFtsmqHC))
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DUwaKxczx(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jYnqqd0o.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HwiaDr(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ csIWD6aBM0.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func MaWXbB2(bZKbGTxeit string, jfbAuSbvU bool, kL6BYHuP uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(bZKbGTxeit)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ yERMajErX2n(gf8OkuRuJg9, jfbAuSbvU, kL6BYHuP)
}

func yERMajErX2n(bZKbGTxeit *uint16, jfbAuSbvU bool, kL6BYHuP uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var otNQQdF4 uint32
	if jfbAuSbvU {
		otNQQdF4 = 1
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bj4DOFu.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)),  /*line :1*/ uintptr(otNQQdF4),  /*line :1*/ uintptr(kL6BYHuP))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func it9q8T(gYwswmeUjG *uint16, iIzhNC *qW7t5nXz) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gwhZ4to77e_.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iIzhNC)), 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func G_iKnZ4ra(za6KY5c *uint16, zNEAk77a *uint16, oRF4hrqvXHJ uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ qa1Lgla1x.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(za6KY5c)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zNEAk77a)),  /*line :1*/ uintptr(oRF4hrqvXHJ))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func QG8hrA(xyxS2P *uint16, oRF4hrqvXHJ uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ o95QuHf41_.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xyxS2P)),  /*line :1*/ uintptr(oRF4hrqvXHJ), 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func BktCQbdIo(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ iio7GalO.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func eIUeXFp_Eqa(iOvctVD26lA L2L8P9WaNZ, iIzhNC *qW7t5nXz) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ e7JShZLAEWHh.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iIzhNC)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func A3GDINmAv(l2bVGVRr7OaX L2L8P9WaNZ, zNEAk77a *uint16, oRF4hrqvXHJ uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ amG5hRVnO6W.Addr(), 3,  /*line :1*/ uintptr(l2bVGVRr7OaX),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zNEAk77a)),  /*line :1*/ uintptr(oRF4hrqvXHJ))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func KQbpaSP(zUwSaIZN5RL L2L8P9WaNZ, xyxS2P *uint16, oRF4hrqvXHJ uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nxk3maL.Addr(), 3,  /*line :1*/ uintptr(zUwSaIZN5RL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xyxS2P)),  /*line :1*/ uintptr(oRF4hrqvXHJ))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func uFpH3ANUM(tiLd7g0 L2L8P9WaNZ, gYwswmeUjG uintptr, i0LVb2 uintptr) (dEaAptPTS5SC L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ qmVcmW.Addr(), 3,  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr(gYwswmeUjG),  /*line :1*/ uintptr(i0LVb2))
	dEaAptPTS5SC =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if dEaAptPTS5SC == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func YQBJy02(zUwSaIZN5RL L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ imrJlMBNU6kK.Addr(), 1,  /*line :1*/ uintptr(zUwSaIZN5RL), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EYMC8Y(l2bVGVRr7OaX L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ttUT3f.Addr(), 1,  /*line :1*/ uintptr(l2bVGVRr7OaX), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FVgKyMq9Ds1(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ oMwPpJQn.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Zy6i6HpeiVS(y6EIVXLlHP uintptr, fORnBp uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ h4rbiaKG.Addr(), 2,  /*line :1*/ uintptr(y6EIVXLlHP),  /*line :1*/ uintptr(fORnBp), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FLUYLrlNYT(a_MrGKcrR uint32, zdBDwshw uintptr, gj1mH6 uint32, q7u0wL0 uint32, etRYtnVni []uint16, feOynBEj *byte) (krzuku uint32, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	if  /*line :1*/ len(etRYtnVni) > 0 {
		gf8OkuRuJg9 = &etRYtnVni[0]
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ zc4E3z6YITt.Addr(), 7,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(zdBDwshw),  /*line :1*/ uintptr(gj1mH6),  /*line :1*/ uintptr(q7u0wL0),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gf8OkuRuJg9)),  /*line :1*/ uintptr( /*line :1*/ len(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(feOynBEj)), 0, 0)
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func S3IWKmiZKOm(jaiV4gvVf *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gNJdLuX5.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jaiV4gvVf)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func I_aT6TrPE(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ timzKykXsz5P.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func WtoSkUy(niq_8nkp uint32, hVw_HCDrJ4 uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eiQsI2LYOWK.Addr(), 2,  /*line :1*/ uintptr(niq_8nkp),  /*line :1*/ uintptr(hVw_HCDrJ4), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func NPX8F39XLdC() (e6fdGdpA uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jrwNr6w.Addr(), 0, 0, 0, 0)
	e6fdGdpA =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func AKe9xkV(eLsmMe46H uint16) (aNqiLgOJxfqs uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zRRpf8_H3m0P.Addr(), 1,  /*line :1*/ uintptr(eLsmMe46H), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func UV1unMTdGY(iOvctVD26lA L2L8P9WaNZ, umPrR0f_tLX *Of8qnJBbbqZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xVSi9waW.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(umPrR0f_tLX)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FK0d7a1YU7() (aGRCzBj3 *uint16) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ pMKYJhg.Addr(), 0, 0, 0, 0)
	aGRCzBj3 = (* /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	return
}

func WyA7ZF(bu5lIjgpJMPv uint32, etRYtnVni *uint16, krzuku *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hd1raicXp.Addr(), 3,  /*line :1*/ uintptr(bu5lIjgpJMPv),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(krzuku)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IwxAIk(etRYtnVni *uint16, krzuku *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ d1W4tbqLj.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(krzuku)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func KlrVuqN(rprJNy L2L8P9WaNZ, yAZay4eEB0M *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cO6x1hNZ.Addr(), 2,  /*line :1*/ uintptr(rprJNy),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yAZay4eEB0M)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SOMEAsjDNv(rprJNy L2L8P9WaNZ, a7Z1LWSG *Ia2ZvPt) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ mNaE1ug0vg.Addr(), 2,  /*line :1*/ uintptr(rprJNy),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7Z1LWSG)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func M9mJEOy_P(cuGEvMbsvH uint32, etRYtnVni *uint16) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ r3yBnvHv6Dz6.Addr(), 2,  /*line :1*/ uintptr(cuGEvMbsvH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)), 0)
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func S31cyZsdeH() (us2BXoyQ_y uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ p_ZSJ6uF.Addr(), 0, 0, 0, 0)
	us2BXoyQ_y =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func F2mr_7() (aMNJNTT uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ j0WpY1_eP.Addr(), 0, 0, 0, 0)
	aMNJNTT =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func RdfSBkk9(cxnD07Z20Caw *uint16, j2UUvA9VgXK *uint64, iaOztN *uint64, aKcjrlhR *uint64) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ rOGQ4py9.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cxnD07Z20Caw)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j2UUvA9VgXK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iaOztN)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aKcjrlhR)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func TpfJpq4V(za6KY5c *uint16) (ae1qbEE uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aXJ6st4WJ16.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(za6KY5c)), 0, 0)
	ae1qbEE =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func DW8kIj() (jaiV4gvVf *uint16, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ yQcuQrvaPb.Addr(), 0, 0, 0, 0)
	jaiV4gvVf = (* /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if jaiV4gvVf == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DislPizUOM5p(gYwswmeUjG *uint16, rpp_GpxY *uint16, nFtsmqHC uint32) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ iKkypUmkt.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpp_GpxY)),  /*line :1*/ uintptr(nFtsmqHC))
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EyQZrxwxNGGt(iOvctVD26lA L2L8P9WaNZ, tmIM88dtF *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hSCVeu.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tmIM88dtF)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EwahWn45(gYwswmeUjG *uint16, bR_S51 uint32, a7Z1LWSG *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ sHMBhxO8qNaR.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(bR_S51),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7Z1LWSG)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FTfH6CHf_Xao(gYwswmeUjG *uint16) (jQSvSZ uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ m3AskwD_du.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)), 0, 0)
	jQSvSZ =  /*line :1*/ uint32(gPQ4CDCVeA)
	if jQSvSZ == INVALID_FILE_ATTRIBUTES {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AzJprcPI4G(iOvctVD26lA L2L8P9WaNZ, iIzhNC *JbW2K4f_20) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jAwW6UwD.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iIzhNC)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ArjHy2rMkrUw(iOvctVD26lA L2L8P9WaNZ, jzcRahcgUWi2 uint32, pB77mVqN *byte, bG2VKsGU uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ jfaoWPMM.Addr(), 4,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(jzcRahcgUWi2),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pB77mVqN)),  /*line :1*/ uintptr(bG2VKsGU), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DacTRPa9NRZ(iOvctVD26lA L2L8P9WaNZ, iYonqx *ZPa9KL2, nGcnqKW *ZPa9KL2, y2Awwyaf_Saw *ZPa9KL2) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ nU4CNtbM4.Addr(), 4,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iYonqx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nGcnqKW)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y2Awwyaf_Saw)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func WKG3GcD(r9HmgAa_zUgV L2L8P9WaNZ) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ r1aizjy.Addr(), 1,  /*line :1*/ uintptr(r9HmgAa_zUgV), 0, 0)
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func V1UqEsJPp3om(fMqLrJh1rnD L2L8P9WaNZ, d3DwyGQa8gGT *uint16, x6bF2v5LRBB uint32, a_MrGKcrR uint32) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ oqfa2LJe.Addr(), 4,  /*line :1*/ uintptr(fMqLrJh1rnD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d3DwyGQa8gGT)),  /*line :1*/ uintptr(x6bF2v5LRBB),  /*line :1*/ uintptr(a_MrGKcrR), 0, 0)
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IZtvZd6jy53(bZKbGTxeit *uint16, cuGEvMbsvH uint32, etRYtnVni *uint16, ovZrxEBKDF **uint16) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ aXi9qd7.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)),  /*line :1*/ uintptr(cuGEvMbsvH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ovZrxEBKDF)), 0, 0)
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZeuaqsG74A() (nFtsmqHC uintptr) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bUWLsMQE.Addr(), 0, 0, 0, 0)
	nFtsmqHC =  /*line :1*/ uintptr(gPQ4CDCVeA)
	return
}

func RtV9uIyO1QrO() (aPxgOyr_3 error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xdEhaAXLL.Addr(), 0, 0, 0, 0)
	if gPQ4CDCVeA != 0 {
		aPxgOyr_3 =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func IAln8aA(oRF4hrqvXHJ uint32, rpp_GpxY *uint16) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ c7TRzPNlkUm.Addr(), 2,  /*line :1*/ uintptr(oRF4hrqvXHJ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpp_GpxY)), 0)
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func UM7nOjWaVbA() (rz3r08Qq uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ sbVv1E1Ll.Addr(), 0, 0, 0, 0)
	rz3r08Qq =  /*line :1*/ uint32(gPQ4CDCVeA)
	if rz3r08Qq == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ErP3HmHK(bZKbGTxeit *uint16, etRYtnVni *uint16, cuGEvMbsvH uint32) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ebxMGuMI.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr(cuGEvMbsvH))
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FvXfNn9p9WkF(eLsmMe46H uint16) (aNqiLgOJxfqs uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ i5NY4Y.Addr(), 1,  /*line :1*/ uintptr(eLsmMe46H), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func BsfS8Cy(tiLd7g0 L2L8P9WaNZ, ybpEJar *uint16, nFtsmqHC uint32) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xfoRaj.Addr(), 3,  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ybpEJar)),  /*line :1*/ uintptr(nFtsmqHC))
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func LhAPWP9(a_MrGKcrR uint32, hLyoZHjSAV *uint16, tiLd7g0 *L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ uDybiDGaQN8W.Addr(), 3,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hLyoZHjSAV)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tiLd7g0)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func YJhK9EyhUhpF(gSssbc L2L8P9WaNZ, dxjY1C *uint32, d688tRziEx *uint32, uWjFxv1q *uint32, p9bcn5qo2 *uint32, dml1vTyYfMR *uint16, kmJCIa uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ tns98o.Addr(), 7,  /*line :1*/ uintptr(gSssbc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dxjY1C)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d688tRziEx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uWjFxv1q)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(p9bcn5qo2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dml1vTyYfMR)),  /*line :1*/ uintptr(kmJCIa), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RjEj209UJL(gSssbc L2L8P9WaNZ, a_MrGKcrR *uint32, fvJPV0 *uint32, rH6FpU9oKLxf *uint32, ua_C6vfqvhe *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ wwqED53e38.Addr(), 5,  /*line :1*/ uintptr(gSssbc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a_MrGKcrR)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fvJPV0)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rH6FpU9oKLxf)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ua_C6vfqvhe)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HefIlFf(iOvctVD26lA L2L8P9WaNZ, g1sgEzk2pZh *Ctdynv, iX0cRH5st *uint32, l6mdvg bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if l6mdvg {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ zchp7tgWa.Addr(), 4,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iX0cRH5st)),  /*line :1*/ uintptr(gf8OkuRuJg9), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func OXnStQK(a_VYqwS4m2B L2L8P9WaNZ) (aNqiLgOJxfqs uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cUS1hCpktM.Addr(), 1,  /*line :1*/ uintptr(a_VYqwS4m2B), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ uint32(gPQ4CDCVeA)
	if aNqiLgOJxfqs == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func H3aCKgJd(tiLd7g0 L2L8P9WaNZ, uwEbpuafGhb string) (aTSUoJ40C uintptr, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.Oea4iRaRU2r(uwEbpuafGhb)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ cjXRZWi7(tiLd7g0, gf8OkuRuJg9)
}

func cjXRZWi7(tiLd7g0 L2L8P9WaNZ, uwEbpuafGhb *byte) (aTSUoJ40C uintptr, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ phplFjs8e.Addr(), 2,  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uwEbpuafGhb)), 0)
	aTSUoJ40C =  /*line :1*/ uintptr(gPQ4CDCVeA)
	if aTSUoJ40C == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FxXCuaVOy3(a_VYqwS4m2B L2L8P9WaNZ) (aMNJNTT uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dPBJCTAc2W5x.Addr(), 1,  /*line :1*/ uintptr(a_VYqwS4m2B), 0, 0)
	aMNJNTT =  /*line :1*/ uint32(gPQ4CDCVeA)
	if aMNJNTT == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func h4WTQ109(a_MrGKcrR uint32, rzHbAFn *uint32, etRYtnVni *uint16, kvo8765 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ jy83afW70E.Addr(), 4,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rzHbAFn)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kvo8765)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Z3HHTa3u(bR_S51 *uint32, a_MrGKcrR *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ tWI8E4IwZ.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bR_S51)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a_MrGKcrR)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IRRnG2(iOvctVD26lA L2L8P9WaNZ, aIHqq6j *ZPa9KL2, _ChbUj5Panu *ZPa9KL2, it5eyEWnww *ZPa9KL2, apvNqdq1 *ZPa9KL2) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ nhgd77EO4z.Addr(), 5,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aIHqq6j)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_ChbUj5Panu)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(it5eyEWnww)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(apvNqdq1)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FiMUgy5Ym(aAEyKyD7U L2L8P9WaNZ, mQ_BpUrUue *uintptr, x1GSv6dwng *uintptr, a_MrGKcrR *uint32) {
	 /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ iZwaoSZ.Addr(), 4,  /*line :1*/ uintptr(aAEyKyD7U),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mQ_BpUrUue)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x1GSv6dwng)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a_MrGKcrR)), 0, 0)
	return
}

func N7YtWOP(qAYnXF L2L8P9WaNZ, zH95B75Vxa *uint32, hGZWAKz *uintptr, g1sgEzk2pZh **Ctdynv, gegnfSNBFV34 uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dnn_vZXZoQ.Addr(), 5,  /*line :1*/ uintptr(qAYnXF),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zH95B75Vxa)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hGZWAKz)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr(gegnfSNBFV34), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Ewc8TeiZpK(a6ElXUKJ *uint16, hWUEQxtJa *uint16, cuGEvMbsvH uint32) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cuVza66Yl.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a6ElXUKJ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hWUEQxtJa)),  /*line :1*/ uintptr(cuGEvMbsvH))
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func yYlSfA5GPA(hnFFk_ *KqXK1jwYl5VP) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ f7Ph3Aaqm.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hnFFk_)), 0, 0)
	return
}

func CoJkxiYw(bWBOjCD0O uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nFJvF4ejaIC8.Addr(), 1,  /*line :1*/ uintptr(bWBOjCD0O), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func iJsZwZKh(ypArpO2 *uint16, pB2jey8 uint32) (ajJiI4KvLI uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ t_P7JJSSo4O.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ypArpO2)),  /*line :1*/ uintptr(pB2jey8), 0)
	ajJiI4KvLI =  /*line :1*/ uint32(gPQ4CDCVeA)
	if ajJiI4KvLI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aBvdnLT(a_MrGKcrR uint32, rzHbAFn *uint32, etRYtnVni *uint16, kvo8765 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ auMc4lT.Addr(), 4,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rzHbAFn)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kvo8765)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EavnAwBq(esI00UeVTD *ZPa9KL2) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ _OgiYtK.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(esI00UeVTD)), 0, 0)
	return
}

func HtlMb0xGpg3(esI00UeVTD *ZPa9KL2) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hMcQrE.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(esI00UeVTD)), 0, 0)
	return
}

func oLclNaBVbkN(ypArpO2 *uint16, pB2jey8 uint32) (ajJiI4KvLI uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ wTRjBaAazk.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ypArpO2)),  /*line :1*/ uintptr(pB2jey8), 0)
	ajJiI4KvLI =  /*line :1*/ uint32(gPQ4CDCVeA)
	if ajJiI4KvLI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func S9hARGE(cuGEvMbsvH uint32, etRYtnVni *uint16) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bs8o7m4R.Addr(), 2,  /*line :1*/ uintptr(cuGEvMbsvH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)), 0)
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func sLLZssX(a_MrGKcrR uint32, rzHbAFn *uint32, etRYtnVni *uint16, kvo8765 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ hSbVUi.Addr(), 4,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rzHbAFn)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kvo8765)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ol7ABdPCo() (d6JQw5VH4_OZ uint64) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ i8hPKHAU.Addr(), 0, 0, 0, 0)
	d6JQw5VH4_OZ =  /*line :1*/ uint64(gPQ4CDCVeA)
	return
}

func TdUMAjs(duactnIo *OvIRDPS) (fWFwY0fQJ_e uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ i_AYSvU.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(duactnIo)), 0, 0)
	fWFwY0fQJ_e =  /*line :1*/ uint32(gPQ4CDCVeA)
	if fWFwY0fQJ_e == 0xffffffff {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func czYdae(a_MrGKcrR uint32, rzHbAFn *uint32, etRYtnVni *uint16, kvo8765 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ tbc5uUP.Addr(), 4,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rzHbAFn)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kvo8765)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func V5kGZXee() (hvCpNOuZ0iBR uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ c3a6MvVz.Addr(), 0, 0, 0, 0)
	hvCpNOuZ0iBR =  /*line :1*/ uint32(gPQ4CDCVeA)
	if hvCpNOuZ0iBR == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DNX2Qr(fMqLrJh1rnD L2L8P9WaNZ, jsjwMvq *uint16, dKvTxkA8nQa uint32, j75M_PRMM6 *uint32, rKZjuZoAw *uint32, gDnHY79Y3c *uint32, jCGfD3 *uint16, trgg4cNvxe uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ kAuW1ku9a.Addr(), 8,  /*line :1*/ uintptr(fMqLrJh1rnD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jsjwMvq)),  /*line :1*/ uintptr(dKvTxkA8nQa),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j75M_PRMM6)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rKZjuZoAw)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gDnHY79Y3c)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jCGfD3)),  /*line :1*/ uintptr(trgg4cNvxe), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AeovBo(za6KY5c *uint16, jsjwMvq *uint16, dKvTxkA8nQa uint32, j75M_PRMM6 *uint32, rKZjuZoAw *uint32, gDnHY79Y3c *uint32, jCGfD3 *uint16, trgg4cNvxe uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ bZg0ojXbZpK.Addr(), 8,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(za6KY5c)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jsjwMvq)),  /*line :1*/ uintptr(dKvTxkA8nQa),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j75M_PRMM6)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rKZjuZoAw)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gDnHY79Y3c)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jCGfD3)),  /*line :1*/ uintptr(trgg4cNvxe), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ALE9qEW(zNEAk77a *uint16, xyxS2P *uint16, nowiC6N81l3 uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bFlMsLL8YsVP.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zNEAk77a)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xyxS2P)),  /*line :1*/ uintptr(nowiC6N81l3))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Cwu9kPYzg(gpn3u85y482 *uint16, rttcJgFpu *uint16, oRF4hrqvXHJ uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ n9eYxNhL.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gpn3u85y482)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rttcJgFpu)),  /*line :1*/ uintptr(oRF4hrqvXHJ))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func N9uRej8R88D(xyxS2P *uint16, hFDSiaDMkqMA *uint16, oRF4hrqvXHJ uint32, tlRnl1M *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ iUBfaYNESUk.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xyxS2P)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hFDSiaDMkqMA)),  /*line :1*/ uintptr(oRF4hrqvXHJ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tlRnl1M)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func a8LRLFSt(ypArpO2 *uint16, pB2jey8 uint32) (ajJiI4KvLI uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gM6cd3zNQ8.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ypArpO2)),  /*line :1*/ uintptr(pB2jey8), 0)
	ajJiI4KvLI =  /*line :1*/ uint32(gPQ4CDCVeA)
	if ajJiI4KvLI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ckL_c4L8(fIzbyI *JEfkblUWi, r3DMk1E uint32, a_MrGKcrR uint32, nFtsmqHC *uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dd81wfTJ1.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fIzbyI)),  /*line :1*/ uintptr(r3DMk1E),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nFtsmqHC)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func KLxhSwkMToc(iOvctVD26lA L2L8P9WaNZ, bxYq0rcLzNav *bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if *bxYq0rcLzNav {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nMpiNCgYohhA.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&gf8OkuRuJg9)), 0)
	*bxYq0rcLzNav = gf8OkuRuJg9 != 0
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func QSVI5gbwA2P8(iOvctVD26lA L2L8P9WaNZ, zaZEBWK7Ag *uint16, fQseGWCwify *uint16) (jeWMpOaQtWV error) {
	jeWMpOaQtWV =  /*line :1*/ cLuFR14.Find()
	if jeWMpOaQtWV != nil {
		return
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cLuFR14.Addr(), 3,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zaZEBWK7Ag)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fQseGWCwify)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JtGRA3Jqe(lOLm8yp6dhGh string, iPUENKB L2L8P9WaNZ, a_MrGKcrR uintptr) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(lOLm8yp6dhGh)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ vTExDb(gf8OkuRuJg9, iPUENKB, a_MrGKcrR)
}

func vTExDb(lOLm8yp6dhGh *uint16, iPUENKB L2L8P9WaNZ, a_MrGKcrR uintptr) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ flb3BaAKwrp.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lOLm8yp6dhGh)),  /*line :1*/ uintptr(iPUENKB),  /*line :1*/ uintptr(a_MrGKcrR))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func C2_KalDvp(lOLm8yp6dhGh string) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(lOLm8yp6dhGh)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ q18rfktiTM(gf8OkuRuJg9)
}

func q18rfktiTM(lOLm8yp6dhGh *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ r6tsM53K4.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lOLm8yp6dhGh)), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SJwhtYsfPML(tiLd7g0 L2L8P9WaNZ, dEaAptPTS5SC L2L8P9WaNZ) (crYOfg3 L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aNrUSf.Addr(), 2,  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr(dEaAptPTS5SC), 0)
	crYOfg3 =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if crYOfg3 == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func R_E3wm(a_MrGKcrR uint32, fORnBp uint32) (nauDv3A uintptr, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jYjjOde0.Addr(), 2,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(fORnBp), 0)
	nauDv3A =  /*line :1*/ uintptr(gPQ4CDCVeA)
	if nauDv3A == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func OtMI6tu9(eGgGffPnzJg1 L2L8P9WaNZ) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xJlj4Ug.Addr(), 1,  /*line :1*/ uintptr(eGgGffPnzJg1), 0, 0)
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA != 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func J47CUgjHF(fMqLrJh1rnD L2L8P9WaNZ, a_MrGKcrR uint32, yv8vYCF3 uint32, dB9EfIT uint32, qY8SHfc uint32, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dbz37ZcNKkha.Addr(), 6,  /*line :1*/ uintptr(fMqLrJh1rnD),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr(dB9EfIT),  /*line :1*/ uintptr(qY8SHfc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JvJtQTH9y(crYOfg3 L2L8P9WaNZ) (y6EIVXLlHP uintptr, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ atN9CYv.Addr(), 1,  /*line :1*/ uintptr(crYOfg3), 0, 0)
	y6EIVXLlHP =  /*line :1*/ uintptr(gPQ4CDCVeA)
	if y6EIVXLlHP == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AZZYGG23T(iOvctVD26lA L2L8P9WaNZ, oW8M6Qq uint32, iWB_umC1I uint32, nXwJCqCxYN uint32, fORnBp uintptr) (y6EIVXLlHP uintptr, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ zdRZ6gYk9yz.Addr(), 5,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(oW8M6Qq),  /*line :1*/ uintptr(iWB_umC1I),  /*line :1*/ uintptr(nXwJCqCxYN),  /*line :1*/ uintptr(fORnBp), 0)
	y6EIVXLlHP =  /*line :1*/ uintptr(gPQ4CDCVeA)
	if y6EIVXLlHP == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZOINqI2w(uGiEDsx L2L8P9WaNZ, hMpYOK *HLa_MM7h) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ h4bGI5Ws.Addr(), 2,  /*line :1*/ uintptr(uGiEDsx),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hMpYOK)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HWECcm_jTi(uGiEDsx L2L8P9WaNZ, hMpYOK *HLa_MM7h) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ beJAjJtvN.Addr(), 2,  /*line :1*/ uintptr(uGiEDsx),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hMpYOK)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func CjF0vhbT3(wPbtqIJl *uint16, xtbcgl1U7 *uint16, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xVBSVJ0.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wPbtqIJl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xtbcgl1U7)),  /*line :1*/ uintptr(a_MrGKcrR))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func U1q5IAHgH(wPbtqIJl *uint16, xtbcgl1U7 *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ qsfLQmUE40.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wPbtqIJl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xtbcgl1U7)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func GHWmZziFIP(p77hSwEH uint32, mxalYKwWF22k uint32, xukLmcNaR *byte, mTr0cqxR09o int32, eA3MhpJ *uint16, cQ5dBwhXygw int32) (e4UkSnX int32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ yL9X4SF.Addr(), 6,  /*line :1*/ uintptr(p77hSwEH),  /*line :1*/ uintptr(mxalYKwWF22k),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xukLmcNaR)),  /*line :1*/ uintptr(mTr0cqxR09o),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eA3MhpJ)),  /*line :1*/ uintptr(cQ5dBwhXygw))
	e4UkSnX =  /*line :1*/ int32(gPQ4CDCVeA)
	if e4UkSnX == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func PfWmmqbG3Pg(nSYdQwq uint32, s0BjGCmzm bool, gYwswmeUjG *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if s0BjGCmzm {
		gf8OkuRuJg9 = 1
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dfrmKibyD.Addr(), 3,  /*line :1*/ uintptr(nSYdQwq),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IRjdlxPyua3o(nSYdQwq uint32, s0BjGCmzm bool, gYwswmeUjG *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if s0BjGCmzm {
		gf8OkuRuJg9 = 1
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ znkNcJBj8I8M.Addr(), 3,  /*line :1*/ uintptr(nSYdQwq),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HXccxkWtt(nSYdQwq uint32, s0BjGCmzm bool, cMFIuNaB uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if s0BjGCmzm {
		gf8OkuRuJg9 = 1
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aApp6pCV0.Addr(), 3,  /*line :1*/ uintptr(nSYdQwq),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(cMFIuNaB))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func YOYRigi6C(nSYdQwq uint32, s0BjGCmzm bool, zNupR3wJl uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if s0BjGCmzm {
		gf8OkuRuJg9 = 1
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ r9Wpk49.Addr(), 3,  /*line :1*/ uintptr(nSYdQwq),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(zNupR3wJl))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func VDJAmLY9P(qAYnXF L2L8P9WaNZ, zH95B75Vxa uint32, hGZWAKz uintptr, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ tXNb6vd9S.Addr(), 4,  /*line :1*/ uintptr(qAYnXF),  /*line :1*/ uintptr(zH95B75Vxa),  /*line :1*/ uintptr(hGZWAKz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Mdm1wP1(uGiEDsx L2L8P9WaNZ, yj2P3JPS5i *RTsQKj) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ n6AOnTph.Addr(), 2,  /*line :1*/ uintptr(uGiEDsx),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yj2P3JPS5i)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RsaLWPf2NmF(uGiEDsx L2L8P9WaNZ, yj2P3JPS5i *RTsQKj) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ phz8mP.Addr(), 2,  /*line :1*/ uintptr(uGiEDsx),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yj2P3JPS5i)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IkDtEOzWY(us2BXoyQ_y uint32, y9aEz060 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zVK1bCJ5.Addr(), 2,  /*line :1*/ uintptr(us2BXoyQ_y),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y9aEz060)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func QzmfYd0ecZO(g_3oYJh6F L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gu34DC.Addr(), 1,  /*line :1*/ uintptr(g_3oYJh6F), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Ca0FXJM(qNhx_SU2 *uint16, jORjJnF *uint16, wGCwTOz uint32) (krzuku uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ vy22cl.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qNhx_SU2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jORjJnF)),  /*line :1*/ uintptr(wGCwTOz))
	krzuku =  /*line :1*/ uint32(gPQ4CDCVeA)
	if krzuku == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AyQ1tary8Y(aTSUoJ40C L2L8P9WaNZ, a_MrGKcrR uint32, oE4kFjB *uint16, nFtsmqHC *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ bGphLF.Addr(), 4,  /*line :1*/ uintptr(aTSUoJ40C),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oE4kFjB)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nFtsmqHC)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DhOQo9d(panFWjsK L2L8P9WaNZ, B77E5eYD int32, EmROY8XA uintptr, WNhK53iwbjUk uint32, pwmIckxOH *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ ec8i2ZYYW.Addr(), 5,  /*line :1*/ uintptr(panFWjsK),  /*line :1*/ uintptr(B77E5eYD),  /*line :1*/ uintptr(EmROY8XA),  /*line :1*/ uintptr(WNhK53iwbjUk),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pwmIckxOH)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func NyUvezY(rprJNy L2L8P9WaNZ, etRYtnVni *uint16, gajkz1KYI uint32, hUJm3uP *uint32, jgA9GYMI5k *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ uN8luvVKi.Addr(), 5,  /*line :1*/ uintptr(rprJNy),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr(gajkz1KYI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hUJm3uP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jgA9GYMI5k)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EiPEiy(iOvctVD26lA L2L8P9WaNZ, etRYtnVni *byte, cuGEvMbsvH uint32, olfqiF5CZluJ bool, dLarUKMqaxh uint32, pwmIckxOH *uint32, g1sgEzk2pZh *Ctdynv, sMydC6ooEuu uintptr) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if olfqiF5CZluJ {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ rMzKw5xAW.Addr(), 8,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr(cuGEvMbsvH),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(dLarUKMqaxh),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pwmIckxOH)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr(sMydC6ooEuu), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func oa1U_Bh2x(iOvctVD26lA L2L8P9WaNZ, etRYtnVni []byte, iX0cRH5st *uint32, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	if  /*line :1*/ len(etRYtnVni) > 0 {
		gf8OkuRuJg9 = &etRYtnVni[0]
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ bZ1PFnHn8dR.Addr(), 5,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gf8OkuRuJg9)),  /*line :1*/ uintptr( /*line :1*/ len(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iX0cRH5st)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func YCns9R(a_VYqwS4m2B L2L8P9WaNZ, vpzStC54ACvz uintptr, rpp_GpxY *byte, nFtsmqHC uintptr, a6K8rzO *uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ gyVt5BpKpyOZ.Addr(), 5,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(vpzStC54ACvz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpp_GpxY)),  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a6K8rzO)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func OossOJiR(aXB7mfyc_cra L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ fypdd7.Addr(), 1,  /*line :1*/ uintptr(aXB7mfyc_cra), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IR_bmXo78A(bZKbGTxeit *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hbiBr_k_.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Rt6lj72HS2(pmQD4ZWSkT15 uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ fCsrW_Ax.Addr(), 1,  /*line :1*/ uintptr(pmQD4ZWSkT15), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Zx48JwR4(g_3oYJh6F L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ j1WTUmmwZxlb.Addr(), 1,  /*line :1*/ uintptr(g_3oYJh6F), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func flaybp5yd1(zaIN6WS L2L8P9WaNZ, nFtsmqHC uint32) (eh2rtH05b2x error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ atyQDow.Addr(), 2,  /*line :1*/ uintptr(zaIN6WS),  /*line :1*/ uintptr(nFtsmqHC), 0)
	if gPQ4CDCVeA != 0 {
		eh2rtH05b2x =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func SiMwChwqLx(q0XnBF4Z L2L8P9WaNZ) (aNqiLgOJxfqs uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ yaBz2vOv5.Addr(), 1,  /*line :1*/ uintptr(q0XnBF4Z), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ uint32(gPQ4CDCVeA)
	if aNqiLgOJxfqs == 0xffffffff {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func WzjhswlJH(iOvctVD26lA L2L8P9WaNZ, umPrR0f_tLX *Of8qnJBbbqZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kQFg4pHW.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(umPrR0f_tLX)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func bdqQnj3n(rprJNy L2L8P9WaNZ, pYy71tl uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dkHSaC_.Addr(), 2,  /*line :1*/ uintptr(rprJNy),  /*line :1*/ uintptr(pYy71tl), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func TsArEvFSO(rprJNy L2L8P9WaNZ, yAZay4eEB0M uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ e0w5xz8sq.Addr(), 2,  /*line :1*/ uintptr(rprJNy),  /*line :1*/ uintptr(yAZay4eEB0M), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func PJdCxj(bZKbGTxeit *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dQq_kSWx8.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EUFzEqkyxic(jU0FKR23IN uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ scqxDo2.Addr(), 1,  /*line :1*/ uintptr(jU0FKR23IN), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func GlaDR0C6INNI(bZKbGTxeit string) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(bZKbGTxeit)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ mYY3k94_(gf8OkuRuJg9)
}

func mYY3k94_(bZKbGTxeit *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gO8hNa4sw9.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RAGLhX7qwfg(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cRNjWrtJiIuL.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Br1XSIP2UtIC(iOvctVD26lA L2L8P9WaNZ, aLIgTgwrtKl_ int64) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zrEA6GoxYvEt.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(aLIgTgwrtKl_), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func V9UKmjv0I(gYwswmeUjG *uint16, wvZBcfh *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gaTyiPthz8.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wvZBcfh)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func YvIEcR(yAZay4eEB0M uint32) (aNqiLgOJxfqs uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ di9Yx9HlFt.Addr(), 1,  /*line :1*/ uintptr(yAZay4eEB0M), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func EdNTvqA(g_3oYJh6F L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ wnh64vsJp.Addr(), 1,  /*line :1*/ uintptr(g_3oYJh6F), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JrSRwp(gYwswmeUjG *uint16, jQSvSZ uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aL5VRBi.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr(jQSvSZ), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func MgtmNjzZnWOm(iOvctVD26lA L2L8P9WaNZ, a_MrGKcrR uint8) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ pMd7CHtSuFpJ.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(a_MrGKcrR), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IkUcELo5bO(iOvctVD26lA L2L8P9WaNZ, jzcRahcgUWi2 uint32, jQqGgBVSKXI5 *byte, bbkj1L2HiQ5W uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ a_XVwp2OWrG.Addr(), 4,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(jzcRahcgUWi2),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jQqGgBVSKXI5)),  /*line :1*/ uintptr(bbkj1L2HiQ5W), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Gl0s_hQDBoz(iOvctVD26lA L2L8P9WaNZ, viMDofpGp int32, kb691INYecZ *int32, h5r9QaWpQ uint32) (fijpScMpG uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ f76dKKeND.Addr(), 4,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(viMDofpGp),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kb691INYecZ)),  /*line :1*/ uintptr(h5r9QaWpQ), 0, 0)
	fijpScMpG =  /*line :1*/ uint32(gPQ4CDCVeA)
	if fijpScMpG == 0xffffffff {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HyQcEQ4lX4wD(iOvctVD26lA L2L8P9WaNZ, iYonqx *ZPa9KL2, nGcnqKW *ZPa9KL2, y2Awwyaf_Saw *ZPa9KL2) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ nvoD_o.Addr(), 4,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iYonqx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nGcnqKW)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y2Awwyaf_Saw)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JbtpNRq(iOvctVD26lA L2L8P9WaNZ, dLarUKMqaxh uint32, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ vKgredhLJ.Addr(), 3,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(dLarUKMqaxh),  /*line :1*/ uintptr(a_MrGKcrR))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func BPFuwqc_f(panFWjsK L2L8P9WaNZ, B77E5eYD uint32, EmROY8XA uintptr, WNhK53iwbjUk uint32) (aNqiLgOJxfqs int, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ vm2vafO.Addr(), 4,  /*line :1*/ uintptr(panFWjsK),  /*line :1*/ uintptr(B77E5eYD),  /*line :1*/ uintptr(EmROY8XA),  /*line :1*/ uintptr(WNhK53iwbjUk), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ int(gPQ4CDCVeA)
	if aNqiLgOJxfqs == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func W1woavU(gSssbc L2L8P9WaNZ, dxjY1C *uint32, uWjFxv1q *uint32, p9bcn5qo2 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ fcm2iDxLXG9U.Addr(), 4,  /*line :1*/ uintptr(gSssbc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dxjY1C)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uWjFxv1q)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(p9bcn5qo2)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func L01wtA1f(a_VYqwS4m2B L2L8P9WaNZ, aSvNEKuFI uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ niK9Fw41s9l.Addr(), 2,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(aSvNEKuFI), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DBa0aeD(a_VYqwS4m2B L2L8P9WaNZ, ufRJmTxs bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if ufRJmTxs {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ tGuZ98Fp.Addr(), 2,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(gf8OkuRuJg9), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func QzmBGm(bR_S51 uint32, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cprOtxS.Addr(), 2,  /*line :1*/ uintptr(bR_S51),  /*line :1*/ uintptr(a_MrGKcrR), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func MQVqemR(aAEyKyD7U L2L8P9WaNZ, nlarVYJaXTvl uintptr, ic05mrGXK uintptr, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ pf7LOQR9P.Addr(), 4,  /*line :1*/ uintptr(aAEyKyD7U),  /*line :1*/ uintptr(nlarVYJaXTvl),  /*line :1*/ uintptr(ic05mrGXK),  /*line :1*/ uintptr(a_MrGKcrR), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func NIOorUx(bWBOjCD0O uint32, iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ d_cuJU2.Addr(), 2,  /*line :1*/ uintptr(bWBOjCD0O),  /*line :1*/ uintptr(iOvctVD26lA), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func B7atQEaase(za6KY5c *uint16, xyxS2P *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gs8QSWB.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(za6KY5c)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xyxS2P)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RambhAHc0vYE(zNEAk77a *uint16, xyxS2P *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nrNospYuk.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zNEAk77a)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xyxS2P)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func BDczRK(tiLd7g0 L2L8P9WaNZ, dEaAptPTS5SC L2L8P9WaNZ) (nFtsmqHC uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ azgFSEhEYz9.Addr(), 2,  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr(dEaAptPTS5SC), 0)
	nFtsmqHC =  /*line :1*/ uint32(gPQ4CDCVeA)
	if nFtsmqHC == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HHybmi9mh(xfXQbTi uint32, v8FMrKQ bool) (aNqiLgOJxfqs uint32) {
	var gf8OkuRuJg9 uint32
	if v8FMrKQ {
		gf8OkuRuJg9 = 1
	}
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ oNZQcj10w.Addr(), 2,  /*line :1*/ uintptr(xfXQbTi),  /*line :1*/ uintptr(gf8OkuRuJg9), 0)
	aNqiLgOJxfqs =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func R7mAlxJBu(panFWjsK L2L8P9WaNZ, z8rB89y uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ yy7nI4I.Addr(), 2,  /*line :1*/ uintptr(panFWjsK),  /*line :1*/ uintptr(z8rB89y), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZUMqr_dh0O(iOvctVD26lA L2L8P9WaNZ, tmIM88dtF uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jLEhbmUt.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(tmIM88dtF), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func NixidkUk(uGiEDsx L2L8P9WaNZ, v09vcf4f9og *MdKJ3U) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ a8x7z713t4ea.Addr(), 2,  /*line :1*/ uintptr(uGiEDsx),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v09vcf4f9og)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Z1aZPwtXq7T(uGiEDsx L2L8P9WaNZ, v09vcf4f9og *MdKJ3U) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ j_33qIs0l.Addr(), 2,  /*line :1*/ uintptr(uGiEDsx),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v09vcf4f9og)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func W0LoSVj8tj(fMqLrJh1rnD L2L8P9WaNZ, yv8vYCF3 uint32, dB9EfIT uint32, qY8SHfc uint32, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ enxXkP.Addr(), 5,  /*line :1*/ uintptr(fMqLrJh1rnD),  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr(dB9EfIT),  /*line :1*/ uintptr(qY8SHfc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Wc5Q6tI(y6EIVXLlHP uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kVPgmk4PMHu.Addr(), 1,  /*line :1*/ uintptr(y6EIVXLlHP), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func lSavoG(fIzbyI *JEfkblUWi, a_MrGKcrR uint32, agfCg_ uintptr, wvZBcfh unsafe.Pointer, nFtsmqHC uintptr, uE5u1dLyb unsafe.Pointer, z_kn2dkp *uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ moyEZ9G.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fIzbyI)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(agfCg_),  /*line :1*/ uintptr(wvZBcfh),  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr(uE5u1dLyb),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(z_kn2dkp)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func G8nOrYl3(f0MVO5WgQ uintptr, nFtsmqHC uintptr, hkDa8YqS uint32, dZE6_dHY uint32) (wvZBcfh uintptr, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ eVHdrWzYuQ.Addr(), 4,  /*line :1*/ uintptr(f0MVO5WgQ),  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr(hkDa8YqS),  /*line :1*/ uintptr(dZE6_dHY), 0, 0)
	wvZBcfh =  /*line :1*/ uintptr(gPQ4CDCVeA)
	if wvZBcfh == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func L9z3rLdPsW(f0MVO5WgQ uintptr, nFtsmqHC uintptr, jqjHTZwMAoon uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ d_3Ge2q.Addr(), 3,  /*line :1*/ uintptr(f0MVO5WgQ),  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr(jqjHTZwMAoon))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HQ4965cpa(y6EIVXLlHP uintptr, fORnBp uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ixcsa0K.Addr(), 2,  /*line :1*/ uintptr(y6EIVXLlHP),  /*line :1*/ uintptr(fORnBp), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func K4R6Nh7A0o(f0MVO5WgQ uintptr, nFtsmqHC uintptr, lTe5Mcw67uQ uint32, x8fUMCcewJ *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ uYxnIMDKMJPq.Addr(), 4,  /*line :1*/ uintptr(f0MVO5WgQ),  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr(lTe5Mcw67uQ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(x8fUMCcewJ)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func STosXbTxZP(a_VYqwS4m2B L2L8P9WaNZ, f0MVO5WgQ uintptr, nFtsmqHC uintptr, tpaarru9tOGM uint32, dSfT_7bJ6c24 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ b3JfFyuQ.Addr(), 5,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(f0MVO5WgQ),  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr(tpaarru9tOGM),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dSfT_7bJ6c24)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func LctiyDMNU(f0MVO5WgQ uintptr, rpp_GpxY *IGaaBCrVFwM, fORnBp uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gdzYzEEy5RTX.Addr(), 3,  /*line :1*/ uintptr(f0MVO5WgQ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpp_GpxY)),  /*line :1*/ uintptr(fORnBp))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func O2J1bw(a_VYqwS4m2B L2L8P9WaNZ, f0MVO5WgQ uintptr, rpp_GpxY *IGaaBCrVFwM, fORnBp uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ pRD_vTdZ.Addr(), 4,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(f0MVO5WgQ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpp_GpxY)),  /*line :1*/ uintptr(fORnBp), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func P25vVjILS37_(y6EIVXLlHP uintptr, fORnBp uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ q4diTlG2k.Addr(), 2,  /*line :1*/ uintptr(y6EIVXLlHP),  /*line :1*/ uintptr(fORnBp), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func J33WdeAz0I() (flq4GwS4tXl uint32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kKkQFAdoY35e.Addr(), 0, 0, 0, 0)
	flq4GwS4tXl =  /*line :1*/ uint32(gPQ4CDCVeA)
	return
}

func bERNe5dqR(rXGaaPBpa uint32, g2LnxiOU53rM uintptr, bnT6r4 bool, keuTsec uint32) (g_3oYJh6F uint32, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if bnT6r4 {
		gf8OkuRuJg9 = 1
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ zAzZMHy1ubCw.Addr(), 4,  /*line :1*/ uintptr(rXGaaPBpa),  /*line :1*/ uintptr(g2LnxiOU53rM),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr(keuTsec), 0, 0)
	g_3oYJh6F =  /*line :1*/ uint32(gPQ4CDCVeA)
	if g_3oYJh6F == 0xffffffff {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func VAQnrWsl_(iOvctVD26lA L2L8P9WaNZ, keuTsec uint32) (g_3oYJh6F uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ut_GQwApMzc.Addr(), 2,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(keuTsec), 0)
	g_3oYJh6F =  /*line :1*/ uint32(gPQ4CDCVeA)
	if g_3oYJh6F == 0xffffffff {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SBldIpJcpjx(rprJNy L2L8P9WaNZ, etRYtnVni *uint16, bNzJXMh uint32, ehaEbK3 *uint32, yv8vYCF3 *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ b6kpufnLgn.Addr(), 5,  /*line :1*/ uintptr(rprJNy),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr(bNzJXMh),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ehaEbK3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yv8vYCF3)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func y8wyCmHs(iOvctVD26lA L2L8P9WaNZ, etRYtnVni []byte, iX0cRH5st *uint32, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	if  /*line :1*/ len(etRYtnVni) > 0 {
		gf8OkuRuJg9 = &etRYtnVni[0]
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ fy6YkY.Addr(), 5,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gf8OkuRuJg9)),  /*line :1*/ uintptr( /*line :1*/ len(etRYtnVni)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iX0cRH5st)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZUBPla(a_VYqwS4m2B L2L8P9WaNZ, vpzStC54ACvz uintptr, rpp_GpxY *byte, nFtsmqHC uintptr, lC3USi *uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ vF00NCFSjky6.Addr(), 5,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(vpzStC54ACvz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rpp_GpxY)),  /*line :1*/ uintptr(nFtsmqHC),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lC3USi)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func MxmpzjD(famY1K L2L8P9WaNZ, irYtRBV22 L2L8P9WaNZ, etRYtnVni *byte, nYIMMA_XpZV uint32, qaKcEoQ1z uint32, w6sOlgTrN uint32, fi0rCIqQ *uint32, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ sk37iQ.Addr(), 8,  /*line :1*/ uintptr(famY1K),  /*line :1*/ uintptr(irYtRBV22),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr(nYIMMA_XpZV),  /*line :1*/ uintptr(qaKcEoQ1z),  /*line :1*/ uintptr(w6sOlgTrN),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fi0rCIqQ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Gva7jyL(etRYtnVni *byte, nYIMMA_XpZV uint32, qaKcEoQ1z uint32, w6sOlgTrN uint32, ecDiVBvh2Y **I6IZfmg, dVdRjzsC5jw *int32, gSdN0FfX6 **I6IZfmg, _aqnLSJMe *int32) {
	 /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ eLK8Wv9i.Addr(), 8,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)),  /*line :1*/ uintptr(nYIMMA_XpZV),  /*line :1*/ uintptr(qaKcEoQ1z),  /*line :1*/ uintptr(w6sOlgTrN),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ecDiVBvh2Y)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dVdRjzsC5jw)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gSdN0FfX6)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_aqnLSJMe)), 0)
	return
}

func Ba8zBaf(bamc83qA3DBr L2L8P9WaNZ, iOvctVD26lA L2L8P9WaNZ, dbW65Wu5fi uint32, ih2Vqc uint32, g1sgEzk2pZh *Ctdynv, dU8rAa3Nz *ZMI04CVXfSX, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ bN_SaD8ahEYL.Addr(), 7,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(dbW65Wu5fi),  /*line :1*/ uintptr(ih2Vqc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dU8rAa3Nz)),  /*line :1*/ uintptr(a_MrGKcrR), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DQG3NH0L0bGi(etRYtnVni *byte) (cREvXF error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gv18jgT2QFD.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)), 0, 0)
	if gPQ4CDCVeA != 0 {
		cREvXF =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func T2nBnJQLE16(lnQPNjA *uint16, gYwswmeUjG **uint16, f6JoXMs *uint32) (cREvXF error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dX3VRZhV0.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lnQPNjA)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(f6JoXMs)))
	if gPQ4CDCVeA != 0 {
		cREvXF =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func Rwfom_si8M(vpG1qVmxE *uint16, dml1vTyYfMR *uint16, bR_S51 uint32, etRYtnVni **byte) (cREvXF error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ u3XXKSqV4x.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vpG1qVmxE)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dml1vTyYfMR)),  /*line :1*/ uintptr(bR_S51),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(etRYtnVni)), 0, 0)
	if gPQ4CDCVeA != 0 {
		cREvXF =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func QAe4mI4(iOvctVD26lA *L2L8P9WaNZ, oW8M6Qq uint32, niDr8AkjX_ *KosbF9m8bMG, hXifJ14 *KVBm_Q, tKU8bpx *int64, aDpnctjNp uint32, jBYCXI uint32, fBukjsVdI uint32, xPmyKOJbMp1Q uint32, qrzoUWC6O3bT uintptr, ztv6kHU uint32) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ b9KOAbv.Addr(), 11,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iOvctVD26lA)),  /*line :1*/ uintptr(oW8M6Qq),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(niDr8AkjX_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hXifJ14)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tKU8bpx)),  /*line :1*/ uintptr(aDpnctjNp),  /*line :1*/ uintptr(jBYCXI),  /*line :1*/ uintptr(fBukjsVdI),  /*line :1*/ uintptr(xPmyKOJbMp1Q),  /*line :1*/ uintptr(qrzoUWC6O3bT),  /*line :1*/ uintptr(ztv6kHU), 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func RGJ8GOALaTl(gSssbc *L2L8P9WaNZ, oW8M6Qq uint32, niDr8AkjX_ *KosbF9m8bMG, hXifJ14 *KVBm_Q, jBYCXI uint32, fBukjsVdI uint32, xPmyKOJbMp1Q uint32, pgq0s4 uint32, lQ6YnZlACe uint32, f4NCUlM uint32, ua_C6vfqvhe uint32, vbYjz6hNxUHj uint32, lURCkwOn uint32, gegnfSNBFV34 *int64) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.SAshL0on( /*line :1*/ leRf3vd2BWUq.Addr(), 14,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gSssbc)),  /*line :1*/ uintptr(oW8M6Qq),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(niDr8AkjX_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hXifJ14)),  /*line :1*/ uintptr(jBYCXI),  /*line :1*/ uintptr(fBukjsVdI),  /*line :1*/ uintptr(xPmyKOJbMp1Q),  /*line :1*/ uintptr(pgq0s4),  /*line :1*/ uintptr(lQ6YnZlACe),  /*line :1*/ uintptr(f4NCUlM),  /*line :1*/ uintptr(ua_C6vfqvhe),  /*line :1*/ uintptr(vbYjz6hNxUHj),  /*line :1*/ uintptr(lURCkwOn),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gegnfSNBFV34)), 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func NarYfY99rS(aTSUoJ40C L2L8P9WaNZ, mBu_6P int32, beemvVDQ673 unsafe.Pointer, jSwN5bag7 uint32, pippUr *uint32) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ f1I2vh.Addr(), 5,  /*line :1*/ uintptr(aTSUoJ40C),  /*line :1*/ uintptr(mBu_6P),  /*line :1*/ uintptr(beemvVDQ673),  /*line :1*/ uintptr(jSwN5bag7),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pippUr)), 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func Whm47JFqM(qNA3xI int32, ie0yyvjq8ra unsafe.Pointer, jqZpITMc7 uint32, pippUr *uint32) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ gk9otj2Cno.Addr(), 4,  /*line :1*/ uintptr(qNA3xI),  /*line :1*/ uintptr(ie0yyvjq8ra),  /*line :1*/ uintptr(jqZpITMc7),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pippUr)), 0, 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func AjxANZe1YWP(iOvctVD26lA L2L8P9WaNZ, hXifJ14 *KVBm_Q, jQqGgBVSKXI5 *byte, bbkj1L2HiQ5W uint32, jzcRahcgUWi2 uint32) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ szkQzKTl.Addr(), 5,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hXifJ14)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jQqGgBVSKXI5)),  /*line :1*/ uintptr(bbkj1L2HiQ5W),  /*line :1*/ uintptr(jzcRahcgUWi2), 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func GVqXwz(aTSUoJ40C L2L8P9WaNZ, mBu_6P int32, beemvVDQ673 unsafe.Pointer, jSwN5bag7 uint32) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ sgauEecm.Addr(), 4,  /*line :1*/ uintptr(aTSUoJ40C),  /*line :1*/ uintptr(mBu_6P),  /*line :1*/ uintptr(beemvVDQ673),  /*line :1*/ uintptr(jSwN5bag7), 0, 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func WPdPZwYjC1e5(qNA3xI int32, ie0yyvjq8ra unsafe.Pointer, jqZpITMc7 uint32) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ n2W1nsX7.Addr(), 3,  /*line :1*/ uintptr(qNA3xI),  /*line :1*/ uintptr(ie0yyvjq8ra),  /*line :1*/ uintptr(jqZpITMc7))
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func SqaASu4(aMXGhi *Dg5vCVtjEbH7, kcyRvQWU uint32, vpzStC54ACvz uintptr) (aNqiLgOJxfqs bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ixj_dt.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aMXGhi)),  /*line :1*/ uintptr(kcyRvQWU),  /*line :1*/ uintptr(vpzStC54ACvz))
	aNqiLgOJxfqs = gPQ4CDCVeA != 0
	return
}

func Bg452J(hlgZJWcfkz **Na2cliTdMq) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jjYwGcxWD1t.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hlgZJWcfkz)), 0, 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func L_oPhy(aMXGhi *Dg5vCVtjEbH7) (aNqiLgOJxfqs bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bw0MYZov6.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aMXGhi)), 0, 0)
	aNqiLgOJxfqs = gPQ4CDCVeA != 0
	return
}

func PD4EfhSY4q(b569rx *uint16, aBb0EwMH7ov *QxJ3Xtd, bFJD0faApb *uint16, qgeUI_ *NrCszu) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ s19ROhdMF.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b569rx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aBb0EwMH7ov)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bFJD0faApb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgeUI_)), 0, 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func RgAIaNaEUd(b569rx *uint16, aBb0EwMH7ov *QxJ3Xtd, bFJD0faApb *uint16, qgeUI_ *NrCszu) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ pEUdOEzaG.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b569rx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aBb0EwMH7ov)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bFJD0faApb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qgeUI_)), 0, 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func Zlth7mQFBiid() (iscCCPy8 *FbP0mu) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ iepLtyZ_Em7.Addr(), 0, 0, 0, 0)
	iscCCPy8 = (* /*line :1*/ FbP0mu)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	return
}

func n3PGfd(n2tJaJ *uint32, tBRYiG *uint32, idoDATkQaBX *uint32) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ et7PGlkl.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(n2tJaJ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tBRYiG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(idoDATkQaBX)))
	return
}

func yUGGGBD(a7Z1LWSG *GDqcSn) (klVyGKYrgT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ baOL_k57PP.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7Z1LWSG)), 0, 0)
	if gPQ4CDCVeA != 0 {
		klVyGKYrgT =  /*line :1*/ T6LY983(gPQ4CDCVeA)
	}
	return
}

func BmTG_IqbTIoV(aP2SjsjFEJP *DxRO_Vrr7, yDAlmq2 *byte) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bap64dt.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aP2SjsjFEJP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yDAlmq2)), 0)
	return
}

func OIJooIiENmZ(aP2SjsjFEJP *QxJ3Xtd, yDAlmq2 *uint16) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nSRkAVILdpI.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aP2SjsjFEJP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(yDAlmq2)), 0)
	return
}

func b4ap1rlCaX(klVyGKYrgT T6LY983) (aNqiLgOJxfqs syscall.O9Mga3) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nvuzH59Pdi.Addr(), 1,  /*line :1*/ uintptr(klVyGKYrgT), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	return
}

func aXy1KcoG5(j75hgMhY *uint16, a3MOFI *Kw2qafRFaiLg) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ vX8nGEUgO4.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j75hgMhY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a3MOFI)), 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func ed9PsDqDFOWJ(gpouvAACwbWU *Kw2qafRFaiLg) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kwsNnMvF.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gpouvAACwbWU)), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func VVYgIbb9J8a(gYwswmeUjG *uint16, eM5aTEyBI *AzllsQa, cM24Fx0MIn *Kw2qafRFaiLg, aMXGhi **uintptr) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ x289BCT3.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eM5aTEyBI)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cM24Fx0MIn)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aMXGhi)), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func NCRuiK(yv8vYCF3 uintptr, dm_F_Liera9l uint32) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ u9nsjoNt.Addr(), 2,  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr(dm_F_Liera9l), 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func U9bXbY99(f0MVO5WgQ unsafe.Pointer) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ v6aya7lr.Addr(), 1,  /*line :1*/ uintptr(f0MVO5WgQ), 0, 0)
	return
}

func YiX1MK() {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gvxZwyl6yWJ6.Addr(), 0, 0, 0, 0)
	return
}

func hlENtQ7V(pLHnNU99h *Kw2qafRFaiLg, j75hgMhY *uint16, xDT6YpNRq5 int32) (x7J8HIt int32) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kF0FvjfCvsOS.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pLHnNU99h)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j75hgMhY)),  /*line :1*/ uintptr(xDT6YpNRq5))
	x7J8HIt =  /*line :1*/ int32(gPQ4CDCVeA)
	return
}

func AH1Q9Va(a_VYqwS4m2B L2L8P9WaNZ, tiLd7g0 *L2L8P9WaNZ, jQP1Gm_6 uint32, lhXloRrK6 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ bP3_35akZ.Addr(), 4,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tiLd7g0)),  /*line :1*/ uintptr(jQP1Gm_6),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lhXloRrK6)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func WTXxpwqF(a_VYqwS4m2B L2L8P9WaNZ, tiLd7g0 *L2L8P9WaNZ, jQP1Gm_6 uint32, lhXloRrK6 *uint32, u9hVR5aV uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ oFtUReO1.Addr(), 5,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tiLd7g0)),  /*line :1*/ uintptr(jQP1Gm_6),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lhXloRrK6)),  /*line :1*/ uintptr(u9hVR5aV), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func xJxpOac(kJf2QPox7Wdz *uint32, kT51HVrtf87r uint32, pwkETEB8SLrs *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hHnButM2.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kJf2QPox7Wdz)),  /*line :1*/ uintptr(kT51HVrtf87r),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pwkETEB8SLrs)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func S_TohvkPu8D(a_VYqwS4m2B L2L8P9WaNZ, tiLd7g0 L2L8P9WaNZ, mEAhiTXAnx1 *uint16, nFtsmqHC uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ a727stJS9.Addr(), 4,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mEAhiTXAnx1)),  /*line :1*/ uintptr(nFtsmqHC), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SkvUc1(a_VYqwS4m2B L2L8P9WaNZ, tiLd7g0 L2L8P9WaNZ, ybpEJar *uint16, nFtsmqHC uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ gfrXM_QJz_r.Addr(), 4,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ybpEJar)),  /*line :1*/ uintptr(nFtsmqHC), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DrewIVQOOUe(a_VYqwS4m2B L2L8P9WaNZ, tiLd7g0 L2L8P9WaNZ, _x1EykjSPlL7 *D7lqaa8u85qg, jQP1Gm_6 uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ vdEjHsuR_Wn.Addr(), 4,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(tiLd7g0),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_x1EykjSPlL7)),  /*line :1*/ uintptr(jQP1Gm_6), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func QuypFlJfyf(a_VYqwS4m2B L2L8P9WaNZ, aMZ1OTRA19V uintptr, jQP1Gm_6 uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zNEoKN4SvguT.Addr(), 3,  /*line :1*/ uintptr(a_VYqwS4m2B),  /*line :1*/ uintptr(aMZ1OTRA19V),  /*line :1*/ uintptr(jQP1Gm_6))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Gz3aKlQam(gbpoIWZJB L2L8P9WaNZ, jV27H8Ni uint32, mYj7hm uintptr, lO6DTpkFW uintptr, wcN8cHD *uintptr) (aNqiLgOJxfqs error) {
	aNqiLgOJxfqs =  /*line :1*/ iPrMy3sLKj5.Find()
	if aNqiLgOJxfqs != nil {
		return
	}
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ iPrMy3sLKj5.Addr(), 5,  /*line :1*/ uintptr(gbpoIWZJB),  /*line :1*/ uintptr(jV27H8Ni),  /*line :1*/ uintptr(mYj7hm),  /*line :1*/ uintptr(lO6DTpkFW),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wcN8cHD)), 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func V9uEuZKe(wcN8cHD uintptr) (jeWMpOaQtWV error) {
	jeWMpOaQtWV =  /*line :1*/ qAHqhriHGU.Find()
	if jeWMpOaQtWV != nil {
		return
	}
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ qAHqhriHGU.Addr(), 1,  /*line :1*/ uintptr(wcN8cHD), 0, 0)
	return
}

func ZwdeQ2(oVyUzJ uint32, kpKBwv *uint16, kT51HVrtf87r *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ec53tm6.Addr(), 3,  /*line :1*/ uintptr(oVyUzJ),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kpKBwv)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kT51HVrtf87r)))
	if i_EJOiAVI&0xff == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func El6tR4oKc(wG2_He_ *uint16, gy8CmXtZ uint32, k5ak29 uint32, lKbQbK *uint16, kT51HVrtf87r *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ yPBsDEXe.Addr(), 5,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wG2_He_)),  /*line :1*/ uintptr(gy8CmXtZ),  /*line :1*/ uintptr(k5ak29),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lKbQbK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kT51HVrtf87r)), 0)
	if i_EJOiAVI&0xff == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Qi8B5rfAA(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, aEsa_kNra MsQkOr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ afAJSciZdFb.Addr(), 3,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr(aEsa_kNra))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func GkM4P09Im86(gHa6OFe JHgnG0YbHkld, uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ pL7yTi_.Addr(), 3,  /*line :1*/ uintptr(gHa6OFe),  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func VWi8tQ(uGD5H4 WIItFoK) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ imQB0S.Addr(), 1,  /*line :1*/ uintptr(uGD5H4), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aga7luXkmw(dW9gI7 *uint16, tzrIoLyJ18 *Kw2qafRFaiLg, akKe4S8A uint32, ljAFfCOKOT *uint32, mhANSS4o3m *uint16, yv8vYCF3 uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ oWSJa9eo.Addr(), 6,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dW9gI7)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tzrIoLyJ18)),  /*line :1*/ uintptr(akKe4S8A),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ljAFfCOKOT)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mhANSS4o3m)),  /*line :1*/ uintptr(yv8vYCF3))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ic0Xno(tRrEQhp *Kw2qafRFaiLg, dW9gI7 *uint16, gZeeEp8 uint32, ljAFfCOKOT *uint32, mhANSS4o3m *uint16, yv8vYCF3 uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ uXiHUaXuP.Addr(), 6,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tRrEQhp)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dW9gI7)),  /*line :1*/ uintptr(gZeeEp8),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ljAFfCOKOT)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mhANSS4o3m)),  /*line :1*/ uintptr(yv8vYCF3))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func b_Lptvocgo65(tRrEQhp *Kw2qafRFaiLg, rV_CV4S59 uintptr, mhANSS4o3m *uint16, yv8vYCF3 uintptr) (iOvctVD26lA WIItFoK, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ cMlGoE.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tRrEQhp)),  /*line :1*/ uintptr(rV_CV4S59),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mhANSS4o3m)),  /*line :1*/ uintptr(yv8vYCF3), 0, 0)
	iOvctVD26lA =  /*line :1*/ WIItFoK(gPQ4CDCVeA)
	if iOvctVD26lA ==  /*line :1*/ WIItFoK(InvalidHandle) {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func wqgGQbt3gOk(uGD5H4 WIItFoK, B_zgYSq0O *uint16, tRrEQhp *Kw2qafRFaiLg, T6QfyiFR507 *uint16, rV_CV4S59 uintptr, CPoMWwBUG6L GTWay7x, kj4GTXbbJ8XY *BKGe6zA1I1c_) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ hYpAogyU.Addr(), 7,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(B_zgYSq0O)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tRrEQhp)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(T6QfyiFR507)),  /*line :1*/ uintptr(rV_CV4S59),  /*line :1*/ uintptr(CPoMWwBUG6L),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func CHqytC(uGD5H4 WIItFoK) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ rPZay0.Addr(), 1,  /*line :1*/ uintptr(uGD5H4), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func GsqRaqJK(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, aEsa_kNra MsQkOr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ rv1L9c.Addr(), 3,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr(aEsa_kNra))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func le4H0tRTkLCV(uGD5H4 WIItFoK, o6uepAP3rUW uint32, kj4GTXbbJ8XY *BKGe6zA1I1c_) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kFhzC_.Addr(), 3,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr(o6uepAP3rUW),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func blpxoz(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, aEsa_kNra MsQkOr, o6uepAP3rUW uint32, j4llyD5g *Sn8DCTALzB_k) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ bMFDL78r.Addr(), 5,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr(aEsa_kNra),  /*line :1*/ uintptr(o6uepAP3rUW),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j4llyD5g)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aoGAk8eiQp(tRrEQhp *Kw2qafRFaiLg, OqVziyg *uint16, rV_CV4S59 uintptr, KawZohyMNM7U Llqbd1Tk, uGD5H4 WIItFoK, mhANSS4o3m *uint16, yv8vYCF3 uintptr) (iOvctVD26lA WIItFoK, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ gJnjVmGu.Addr(), 7,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tRrEQhp)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(OqVziyg)),  /*line :1*/ uintptr(rV_CV4S59),  /*line :1*/ uintptr(KawZohyMNM7U),  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mhANSS4o3m)),  /*line :1*/ uintptr(yv8vYCF3), 0, 0)
	iOvctVD26lA =  /*line :1*/ WIItFoK(gPQ4CDCVeA)
	if iOvctVD26lA ==  /*line :1*/ WIItFoK(InvalidHandle) {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HMqZZCgR5W(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, ffpjJE16Qt *ZRao_T7a, kl2waNwDJNO3 uint32, ljAFfCOKOT *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ jBgmki.Addr(), 5,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ffpjJE16Qt)),  /*line :1*/ uintptr(kl2waNwDJNO3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ljAFfCOKOT)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func as3wti2MyBA(uGD5H4 WIItFoK, erAGnRa4Uj *F9bfgg) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ xZ4KdVKVG72J.Addr(), 2,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(erAGnRa4Uj)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func poTBP8wYR(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, gBpptk_r *Zyvr1EVBmP) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ a74rMqNh7ad.Addr(), 3,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gBpptk_r)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func x9MM7Y(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, dwGt4Y6S8O *uint16, ux9z_4I uint32, bWdDzaojrOnK *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ svROrZ54Mab.Addr(), 5,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dwGt4Y6S8O)),  /*line :1*/ uintptr(ux9z_4I),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bWdDzaojrOnK)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func vygzmRw(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, aDANLjxdYzx *XqoC_LNSBj, pShma6FCvlJ *F1HyK5M, jgHkPSTDj *byte, z7qCaQk62M uint32, ljAFfCOKOT *uint32, a_MrGKcrR uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ qU1Lpm03neR.Addr(), 8,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aDANLjxdYzx)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pShma6FCvlJ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jgHkPSTDj)),  /*line :1*/ uintptr(z7qCaQk62M),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ljAFfCOKOT)),  /*line :1*/ uintptr(a_MrGKcrR), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func tq3t9Ehr(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, fBn5qKfnS7 Y3qEJ_aVo, mEH6q_3Bf7 *uint32, jgHkPSTDj *byte, z7qCaQk62M uint32, ljAFfCOKOT *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ zkHU4W.Addr(), 7,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr(fBn5qKfnS7),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mEH6q_3Bf7)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jgHkPSTDj)),  /*line :1*/ uintptr(z7qCaQk62M),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ljAFfCOKOT)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func cwEq8JK508(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, j4llyD5g *Sn8DCTALzB_k, dW3nxRX *PrnhDsQ, qi2IUhszBAN uint32, ljAFfCOKOT *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ b99b9vLEH04f.Addr(), 6,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j4llyD5g)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dW3nxRX)),  /*line :1*/ uintptr(qi2IUhszBAN),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ljAFfCOKOT)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func fhWjAaGiUV(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ njn0RLJIG5VC.Addr(), 2,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func h6shLJ_qt9mi(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, j4llyD5g *Sn8DCTALzB_k) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hr2BiYQowrmv.Addr(), 3,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j4llyD5g)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func U9OPHqeld8A(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, IXxa82ht78k ZvvCHRdtq, AyGPXZz uint32, AQ6iswf_ RChEbtEIYbR, e_w62Sbj uint32) (hGZWAKz L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ vcLVCbw.Addr(), 6,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr(IXxa82ht78k),  /*line :1*/ uintptr(AyGPXZz),  /*line :1*/ uintptr(AQ6iswf_),  /*line :1*/ uintptr(e_w62Sbj))
	hGZWAKz =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if hGZWAKz == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Uc8omY_CN(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, ffpjJE16Qt *ZRao_T7a, kl2waNwDJNO3 uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ jSyJMrobErk.Addr(), 4,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ffpjJE16Qt)),  /*line :1*/ uintptr(kl2waNwDJNO3), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Yk1KOGBG_heq(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, gBpptk_r *Zyvr1EVBmP) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ p_ZpQvJd.Addr(), 3,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gBpptk_r)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func vanFmZJ(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, fBn5qKfnS7 Y3qEJ_aVo, jgHkPSTDj *byte, z7qCaQk62M uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ mWfZG5xq.Addr(), 5,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr(fBn5qKfnS7),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jgHkPSTDj)),  /*line :1*/ uintptr(z7qCaQk62M), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func UnukQswxpS(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ n_qfnvwMQfC.Addr(), 2,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FoOJWbA(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, j4llyD5g *Sn8DCTALzB_k) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ i0b3UoB3.Addr(), 3,  /*line :1*/ uintptr(uGD5H4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kj4GTXbbJ8XY)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j4llyD5g)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func gIY4mxFW(bcTTXDUmWK *uint16, a_MrGKcrR Ipdrvc9b9wmZ, yv8vYCF3 uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ u4S5Pwbzpth.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bcTTXDUmWK)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(yv8vYCF3))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aYQ3K6eKrSQ(aGRCzBj3 *uint16, pjMGTnPK *int32) (vPivxOaBYBAf **uint16, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ kEOO1jJP.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aGRCzBj3)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pjMGTnPK)), 0)
	vPivxOaBYBAf = (** /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if vPivxOaBYBAf == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func a0OuuH(aMNJNTT *S06JEslqEe, a_MrGKcrR uint32, lAmUDbCC TaSPPoJMlh, bZKbGTxeit **uint16) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ qcSmpW6.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aMNJNTT)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(lAmUDbCC),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bZKbGTxeit)), 0, 0)
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func NVasPKuHS(i6Cvqni_Es L2L8P9WaNZ, h6f7t1 *uint16, fMqLrJh1rnD *uint16, feOynBEj *uint16, u09YbKn *uint16, dqfbt97 int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ zIPHZWi.Addr(), 6,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(h6f7t1)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fMqLrJh1rnD)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(feOynBEj)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u09YbKn)),  /*line :1*/ uintptr(dqfbt97))
	if i_EJOiAVI <= 32 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func D50cnayN0Pnw(i6Cvqni_Es IJltG1D5QG, uEkyY6a1VFn uintptr, dX3ja0 unsafe.Pointer) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eHDEpG.Addr(), 3,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr(uEkyY6a1VFn),  /*line :1*/ uintptr(dX3ja0))
	return
}

func QIloNd(uEkyY6a1VFn uintptr, dX3ja0 unsafe.Pointer) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ucbH8SHFMjYL.Addr(), 2,  /*line :1*/ uintptr(uEkyY6a1VFn),  /*line :1*/ uintptr(dX3ja0), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SlqXiQ6V(a_MrGKcrR uint32, iV4qRaK uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ h_tXoQTUHez.Addr(), 2,  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(iV4qRaK), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Wznjh3oK(i6Cvqni_Es IJltG1D5QG, dW9gI7 *uint16, caqNX61uBX int32) (rDKO6AqKI int32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ wB2kUS.Addr(), 3,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dW9gI7)),  /*line :1*/ uintptr(caqNX61uBX))
	rDKO6AqKI =  /*line :1*/ int32(gPQ4CDCVeA)
	if rDKO6AqKI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SqFa9j6Q32_() (i6Cvqni_Es IJltG1D5QG) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nBZnlpnU9VU.Addr(), 0, 0, 0, 0)
	i6Cvqni_Es =  /*line :1*/ IJltG1D5QG(gPQ4CDCVeA)
	return
}

func KQ6IjolA7LAi() (i6Cvqni_Es IJltG1D5QG) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ ivveVvqC.Addr(), 0, 0, 0, 0)
	i6Cvqni_Es =  /*line :1*/ IJltG1D5QG(gPQ4CDCVeA)
	return
}

func AaI8Dv(q0XnBF4Z uint32, a7Z1LWSG *KnCkvYIy) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hiQD0hYh.Addr(), 2,  /*line :1*/ uintptr(q0XnBF4Z),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a7Z1LWSG)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Cf4Q0Lmgazc() (pp4jTN IJltG1D5QG) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cXTOag.Addr(), 0, 0, 0, 0)
	pp4jTN =  /*line :1*/ IJltG1D5QG(gPQ4CDCVeA)
	return
}

func O_8aQ2(i6Cvqni_Es IJltG1D5QG, us2BXoyQ_y *uint32) (dNcH99PoxX4K uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ lgjp5wcA.Addr(), 2,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(us2BXoyQ_y)), 0)
	dNcH99PoxX4K =  /*line :1*/ uint32(gPQ4CDCVeA)
	if dNcH99PoxX4K == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func E3S8A6eIBusF(i6Cvqni_Es IJltG1D5QG) (rHabWfJZt bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ mZj5LR.Addr(), 1,  /*line :1*/ uintptr(i6Cvqni_Es), 0, 0)
	rHabWfJZt = gPQ4CDCVeA != 0
	return
}

func Ev9DaYxdY0QG(i6Cvqni_Es IJltG1D5QG) (gzaaUW8 bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jh9hh5Khb6.Addr(), 1,  /*line :1*/ uintptr(i6Cvqni_Es), 0, 0)
	gzaaUW8 = gPQ4CDCVeA != 0
	return
}

func Pdb4Ni8ugaC(i6Cvqni_Es IJltG1D5QG) (zl0ID8n bool) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ afmwd3Qf.Addr(), 1,  /*line :1*/ uintptr(i6Cvqni_Es), 0, 0)
	zl0ID8n = gPQ4CDCVeA != 0
	return
}

func KrRLF9POysW(i6Cvqni_Es IJltG1D5QG, hd29eZQWYL *uint16, eSULhnT6a *uint16, oGXUDVH8n uint32) (aNqiLgOJxfqs int32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dJ1wXA.Addr(), 4,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hd29eZQWYL)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eSULhnT6a)),  /*line :1*/ uintptr(oGXUDVH8n), 0, 0)
	aNqiLgOJxfqs =  /*line :1*/ int32(gPQ4CDCVeA)
	if aNqiLgOJxfqs == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func WdSzERJZLW(o4R_INp **uint16, lAmUDbCC TaSPPoJMlh, wSMYeeS bool) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if wSMYeeS {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gGPSpAaoiQ.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(o4R_INp)),  /*line :1*/ uintptr(lAmUDbCC),  /*line :1*/ uintptr(gf8OkuRuJg9))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func CG120MaPv8q(o4R_INp *uint16) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ mqUYrSH.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(o4R_INp)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func A9DOEMqTc(g9JYFzvU35 TaSPPoJMlh, ypArpO2 *uint16, pB2jey8 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bVn6C9KYOZp.Addr(), 3,  /*line :1*/ uintptr(g9JYFzvU35),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ypArpO2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pB2jey8)))
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func MKCxXGhk06(ybpEJar string, hqbGe2kY7wQu *L2L8P9WaNZ) (kvo8765 uint32, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(ybpEJar)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ x6nA2fiyE1jW(gf8OkuRuJg9, hqbGe2kY7wQu)
}

func x6nA2fiyE1jW(ybpEJar *uint16, hqbGe2kY7wQu *L2L8P9WaNZ) (kvo8765 uint32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ oapyFDaeq95R.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ybpEJar)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hqbGe2kY7wQu)), 0)
	kvo8765 =  /*line :1*/ uint32(gPQ4CDCVeA)
	if kvo8765 == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Lek_JxILaI(ybpEJar string, iOvctVD26lA uint32, kvo8765 uint32, rpp_GpxY unsafe.Pointer) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(ybpEJar)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ srZqqIQscI8(gf8OkuRuJg9, iOvctVD26lA, kvo8765, rpp_GpxY)
}

func srZqqIQscI8(ybpEJar *uint16, iOvctVD26lA uint32, kvo8765 uint32, rpp_GpxY unsafe.Pointer) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ qlT3ngaBXQ.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ybpEJar)),  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(kvo8765),  /*line :1*/ uintptr(rpp_GpxY), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func EAksAO(o4R_INp unsafe.Pointer, tVNaT9 string, qxiVZ3TBUv6 unsafe.Pointer, kvo8765 *uint32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *uint16
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.GcOmFfsibES(tVNaT9)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ mZHNP6(o4R_INp, gf8OkuRuJg9, qxiVZ3TBUv6, kvo8765)
}

func mZHNP6(o4R_INp unsafe.Pointer, tVNaT9 *uint16, qxiVZ3TBUv6 unsafe.Pointer, kvo8765 *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ c1978nB7x.Addr(), 4,  /*line :1*/ uintptr(o4R_INp),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tVNaT9)),  /*line :1*/ uintptr(qxiVZ3TBUv6),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kvo8765)), 0, 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func J8iMZOvWD(mON7me uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ u24YLVfonUpW.Addr(), 1,  /*line :1*/ uintptr(mON7me), 0, 0)
	if i_EJOiAVI != 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ABzqaogt(mON7me uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jfFPVqjCSA.Addr(), 1,  /*line :1*/ uintptr(mON7me), 0, 0)
	if i_EJOiAVI != 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func RbNw1ERA(i6Cvqni_Es IJltG1D5QG, uLvoYNH0nnS *Kw2qafRFaiLg, iIzhNC *B9Rbvc0g) (aNqiLgOJxfqs error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eW7Ynn5Di.Addr(), 3,  /*line :1*/ uintptr(i6Cvqni_Es),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uLvoYNH0nnS)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iIzhNC)))
	if gPQ4CDCVeA != 0 {
		aNqiLgOJxfqs =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func A_nbOejft_(fBacZskv5wOj *OmWWdZ4) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ i4jGXfW.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fBacZskv5wOj)), 0, 0)
	return
}

func Y9ZPc74C1AXD(t9VsTXjVOqGK *uint16, cfITTM *uint16, duY1zT *OmWWdZ4, fz1BsYsTSAQ **OmWWdZ4) (zE65NT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ u6Lcomleb9.Addr(), 4,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t9VsTXjVOqGK)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cfITTM)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(duY1zT)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fz1BsYsTSAQ)), 0, 0)
	if gPQ4CDCVeA != 0 {
		zE65NT =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func AndEUhbq() (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ gEShlkiRvz20.Addr(), 0, 0, 0, 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func JVqHWw2D(qmPr07 *int32, hByOOyZ *I3lBPYH, oRF4hrqvXHJ *uint32) (krzuku int32, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ zcKrPFhWmaw.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qmPr07)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hByOOyZ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(oRF4hrqvXHJ)))
	krzuku =  /*line :1*/ int32(gPQ4CDCVeA)
	if krzuku == -1 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ZMXykCDtaaO3(_ITTe4 L2L8P9WaNZ, mf_Gq_ *Ctdynv, v7WBFUf *uint32, l6mdvg bool, a_MrGKcrR *uint32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 uint32
	if l6mdvg {
		gf8OkuRuJg9 = 1
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ rc6jPPK.Addr(), 5,  /*line :1*/ uintptr(_ITTe4),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mf_Gq_)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(v7WBFUf)),  /*line :1*/ uintptr(gf8OkuRuJg9),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a_MrGKcrR)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func VuX432Ao4_L(bamc83qA3DBr L2L8P9WaNZ, kJML3n uint32, vS8a4Yt7 *byte, lSLqEx uint32, uoiWWOpfFr1 *byte, qstSrja4z3 uint32, mjbRGlfPeCg *uint32, g1sgEzk2pZh *Ctdynv, sMydC6ooEuu uintptr) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ dY1idNbqYPkP.Addr(), 9,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(kJML3n),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vS8a4Yt7)),  /*line :1*/ uintptr(lSLqEx),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(uoiWWOpfFr1)),  /*line :1*/ uintptr(qstSrja4z3),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mjbRGlfPeCg)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr(sMydC6ooEuu))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func KlNJDQ3(rFKS6LV2y *D3kNJ0Y9, a_MrGKcrR uint32, iOvctVD26lA *L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ j3go6rhome6P.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rFKS6LV2y)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iOvctVD26lA)))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func E7VyXfpzlmb(iOvctVD26lA L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nS2_fCv.Addr(), 1,  /*line :1*/ uintptr(iOvctVD26lA), 0, 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AbB9hqe4Q(iOvctVD26lA L2L8P9WaNZ, a_MrGKcrR uint32, nFtsmqHC *int32, rFKS6LV2y *D3kNJ0Y9) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ kBXucH.Addr(), 4,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nFtsmqHC)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rFKS6LV2y)), 0, 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func NeTE869kDuJ(bamc83qA3DBr L2L8P9WaNZ, gCwH7wEHx *GC5qDJGWUpA, jkqgCfg3a0ds uint32, fi0rCIqQ *uint32, a_MrGKcrR *uint32, g1sgEzk2pZh *Ctdynv, fgQM_KRAN9 *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ trVCnrSj.Addr(), 7,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gCwH7wEHx)),  /*line :1*/ uintptr(jkqgCfg3a0ds),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fi0rCIqQ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a_MrGKcrR)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fgQM_KRAN9)), 0, 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func ScV_Dt(bamc83qA3DBr L2L8P9WaNZ, gCwH7wEHx *GC5qDJGWUpA, jkqgCfg3a0ds uint32, fi0rCIqQ *uint32, a_MrGKcrR *uint32, wPbtqIJl *I6IZfmg, ortWfeYOle *int32, g1sgEzk2pZh *Ctdynv, fgQM_KRAN9 *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ hSfwYLVh6UVh.Addr(), 9,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gCwH7wEHx)),  /*line :1*/ uintptr(jkqgCfg3a0ds),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fi0rCIqQ)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(a_MrGKcrR)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wPbtqIJl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ortWfeYOle)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fgQM_KRAN9)))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func O4NtJ3(bamc83qA3DBr L2L8P9WaNZ, gCwH7wEHx *GC5qDJGWUpA, jkqgCfg3a0ds uint32, faR595J6XZ5 *uint32, a_MrGKcrR uint32, g1sgEzk2pZh *Ctdynv, fgQM_KRAN9 *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ qy8B1U6sSZ2.Addr(), 7,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gCwH7wEHx)),  /*line :1*/ uintptr(jkqgCfg3a0ds),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(faR595J6XZ5)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fgQM_KRAN9)), 0, 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Qzva8c0F4tEv(bamc83qA3DBr L2L8P9WaNZ, gCwH7wEHx *GC5qDJGWUpA, jkqgCfg3a0ds uint32, faR595J6XZ5 *uint32, a_MrGKcrR uint32, xtbcgl1U7 *I6IZfmg, aYeA7d3Th_ int32, g1sgEzk2pZh *Ctdynv, fgQM_KRAN9 *byte) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ x29Z4ZK.Addr(), 9,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gCwH7wEHx)),  /*line :1*/ uintptr(jkqgCfg3a0ds),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(faR595J6XZ5)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xtbcgl1U7)),  /*line :1*/ uintptr(aYeA7d3Th_),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fgQM_KRAN9)))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func Kjqlye(vHigVo4Oi0A5 int32, pgq0s4 int32, aP0aJo9XA int32, g0UaOZ *I3lBPYH, bnHAoth uint32, a_MrGKcrR uint32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ kdpPBUJXLYw.Addr(), 6,  /*line :1*/ uintptr(vHigVo4Oi0A5),  /*line :1*/ uintptr(pgq0s4),  /*line :1*/ uintptr(aP0aJo9XA),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g0UaOZ)),  /*line :1*/ uintptr(bnHAoth),  /*line :1*/ uintptr(a_MrGKcrR))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func R5pzEz2fQK(yz0d9a_l4fK uint32, iIzhNC *Rg5HtL) (zE65NT error) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ eSONG7ao2k01.Addr(), 2,  /*line :1*/ uintptr(yz0d9a_l4fK),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iIzhNC)), 0)
	if gPQ4CDCVeA != 0 {
		zE65NT =  /*line :1*/ syscall.O9Mga3(gPQ4CDCVeA)
	}
	return
}

func fCe8PKCARpU(bamc83qA3DBr L2L8P9WaNZ, gYwswmeUjG unsafe.Pointer, jNLsCU int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ wz_sYEr.Addr(), 3,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(gYwswmeUjG),  /*line :1*/ uintptr(jNLsCU))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func AyBoKZZ(bamc83qA3DBr L2L8P9WaNZ) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ dYcxJ1Nbm0.Addr(), 1,  /*line :1*/ uintptr(bamc83qA3DBr), 0, 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aW9zqIBU4(bamc83qA3DBr L2L8P9WaNZ, gYwswmeUjG unsafe.Pointer, jNLsCU int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ nH9XNug.Addr(), 3,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(gYwswmeUjG),  /*line :1*/ uintptr(jNLsCU))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func IHLqgtEPB(gYwswmeUjG string) (_ITTe4 *ZhI4Bn, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.Oea4iRaRU2r(gYwswmeUjG)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ _LWem_4Jp(gf8OkuRuJg9)
}

func _LWem_4Jp(gYwswmeUjG *byte) (_ITTe4 *ZhI4Bn, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hqGIR9_T.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)), 0, 0)
	_ITTe4 = (* /*line :1*/ ZhI4Bn)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if _ITTe4 == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func pzZiaon(bamc83qA3DBr L2L8P9WaNZ, u1fJXFk *I6IZfmg, cQTrCFvYXVc *int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ yTd2gBXO9.Addr(), 3,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u1fJXFk)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cQTrCFvYXVc)))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func HZUrT1B3(gYwswmeUjG string) (hD4wNgEB *AXIyx64vR, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.Oea4iRaRU2r(gYwswmeUjG)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ bvqo3ffKEhl(gf8OkuRuJg9)
}

func bvqo3ffKEhl(gYwswmeUjG *byte) (hD4wNgEB *AXIyx64vR, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ pWD5IsC9_.Addr(), 1,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)), 0, 0)
	hD4wNgEB = (* /*line :1*/ AXIyx64vR)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if hD4wNgEB == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func FoYauaAkG3pI(gYwswmeUjG string, pNa9kYooPwQ string) (bamc83qA3DBr *WGlK0Qv_nne, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	gf8OkuRuJg9, jeWMpOaQtWV =  /*line :1*/ syscall.Oea4iRaRU2r(gYwswmeUjG)
	if jeWMpOaQtWV != nil {
		return
	}
	var otNQQdF4 *byte
	otNQQdF4, jeWMpOaQtWV =  /*line :1*/ syscall.Oea4iRaRU2r(pNa9kYooPwQ)
	if jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ rSyLffL6LE(gf8OkuRuJg9, otNQQdF4)
}

func rSyLffL6LE(gYwswmeUjG *byte, pNa9kYooPwQ *byte) (bamc83qA3DBr *WGlK0Qv_nne, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ acYMNI.Addr(), 2,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYwswmeUjG)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(pNa9kYooPwQ)), 0)
	bamc83qA3DBr = (* /*line :1*/ WGlK0Qv_nne)( /*line :1*/ unsafe.Pointer(gPQ4CDCVeA))
	if bamc83qA3DBr == nil {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func jPN3P54q(bamc83qA3DBr L2L8P9WaNZ, u1fJXFk *I6IZfmg, cQTrCFvYXVc *int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bMbKCI.Addr(), 3,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u1fJXFk)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cQTrCFvYXVc)))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func U7Z_4P(bamc83qA3DBr L2L8P9WaNZ, bR_S51 int32, w5Hke5t int32, n7a3sV7 *byte, mF9n0Bnm *int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ q1UjRaO_qW6.Addr(), 5,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(bR_S51),  /*line :1*/ uintptr(w5Hke5t),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(n7a3sV7)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(mF9n0Bnm)), 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func i3P0kOH(bamc83qA3DBr L2L8P9WaNZ, veIj6d_04s int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ h8wJG5.Addr(), 2,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(veIj6d_04s), 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func J8LHFbQm(f3mt2DYXJM uint16) (sMdg1A uint16) {
	gPQ4CDCVeA, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ bvtTCvmrr.Addr(), 1,  /*line :1*/ uintptr(f3mt2DYXJM), 0, 0)
	sMdg1A =  /*line :1*/ uint16(gPQ4CDCVeA)
	return
}

func nE41zAHtIi(bamc83qA3DBr L2L8P9WaNZ, etRYtnVni []byte, a_MrGKcrR int32, wPbtqIJl *I6IZfmg, ortWfeYOle *int32) (krzuku int32, jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	if  /*line :1*/ len(etRYtnVni) > 0 {
		gf8OkuRuJg9 = &etRYtnVni[0]
	}
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ fSTug6av.Addr(), 6,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gf8OkuRuJg9)),  /*line :1*/ uintptr( /*line :1*/ len(etRYtnVni)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wPbtqIJl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ortWfeYOle)))
	krzuku =  /*line :1*/ int32(gPQ4CDCVeA)
	if krzuku == -1 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func mdvojh1n(bamc83qA3DBr L2L8P9WaNZ, etRYtnVni []byte, a_MrGKcrR int32, xtbcgl1U7 unsafe.Pointer, aYeA7d3Th_ int32) (jeWMpOaQtWV error) {
	var gf8OkuRuJg9 *byte
	if  /*line :1*/ len(etRYtnVni) > 0 {
		gf8OkuRuJg9 = &etRYtnVni[0]
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ kiKK50OrtAbf.Addr(), 6,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gf8OkuRuJg9)),  /*line :1*/ uintptr( /*line :1*/ len(etRYtnVni)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr(xtbcgl1U7),  /*line :1*/ uintptr(aYeA7d3Th_))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func UfWy_ocZ(bamc83qA3DBr L2L8P9WaNZ, bR_S51 int32, w5Hke5t int32, n7a3sV7 *byte, mF9n0Bnm int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ iJfc97.Addr(), 5,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(bR_S51),  /*line :1*/ uintptr(w5Hke5t),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(n7a3sV7)),  /*line :1*/ uintptr(mF9n0Bnm), 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func aM4XrT(bamc83qA3DBr L2L8P9WaNZ, jFffl_H int32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ jLwuVOdY4r.Addr(), 2,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(jFffl_H), 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func oxblUQMm03v(vHigVo4Oi0A5 int32, pgq0s4 int32, aP0aJo9XA int32) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ x99ndB9.Addr(), 3,  /*line :1*/ uintptr(vHigVo4Oi0A5),  /*line :1*/ uintptr(pgq0s4),  /*line :1*/ uintptr(aP0aJo9XA))
	iOvctVD26lA =  /*line :1*/ L2L8P9WaNZ(gPQ4CDCVeA)
	if iOvctVD26lA == InvalidHandle {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func SbeL_a(iOvctVD26lA L2L8P9WaNZ, yv8vYCF3 uint32, k_gpbbUzJH uint32, nQYUVJo1s **PLfKrMR, rXGaaPBpa *uint32) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ rj31a8ch.Addr(), 5,  /*line :1*/ uintptr(iOvctVD26lA),  /*line :1*/ uintptr(yv8vYCF3),  /*line :1*/ uintptr(k_gpbbUzJH),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nQYUVJo1s)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rXGaaPBpa)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func DuVQyM(nauDv3A uintptr) {
	 /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ v15cQdcKjo.Addr(), 1,  /*line :1*/ uintptr(nauDv3A), 0, 0)
	return
}

func UHgVPqOS(hALYowOfgF uint32, lAmUDbCC *TaSPPoJMlh) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ sLmtRhF49vX.Addr(), 2,  /*line :1*/ uintptr(hALYowOfgF),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lAmUDbCC)), 0)
	if i_EJOiAVI == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}
