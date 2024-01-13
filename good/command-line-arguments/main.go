//line :1
package main

import (
	fmt "kFsoCfy5zWG"
	ioutil "cgTcmgSavP2"
	log "pBPwKYPji2"
	os "hPMrClpC"
	filepath "qyOdGuID2rmZ"
	strings "fQXlzv"
	syscall "bUKeamF"
	"unsafe"

	wmi "baLhtK0"
	registry "ER_Q0ps1VO"
)

const (
	TH32CS_SNAPPROCESS = 0x00000002
)

var (
	rpuPuokYRzTr	=  /*line :1*/ syscall.A9tYHZzf("kernel32.dll")
	Z6cYt8Iopdq	=  /*line :1*/ syscall.A9tYHZzf("Ntdll.dll")
	KNjBN0jjq	=  /*line :1*/ syscall.A9tYHZzf("CIMWin32.dll")

	A09LJYjs	=  /*line :1*/ Z6cYt8Iopdq.NewProc("RtlGetVersion")

	BE54OHcRUj	=  /*line :1*/ rpuPuokYRzTr.NewProc("CreateToolhelp32Snapshot")
	Htq3k6	=  /*line :1*/ rpuPuokYRzTr.NewProc("Process32First")
	OhxsB9Js	=  /*line :1*/ rpuPuokYRzTr.NewProc("Process32Next")
)


type LoXY4TrvEa struct {
	InJJB0uGDDz	uint16
	Zl6mAtw	uint16
	Xs7IROaR_t	uint32
	X6dyw7OG	uintptr
	FvtOlqTOk0	uintptr
	JZub5IALER	uintptr
	OvJ3Tx	uint32
	DYlOIHI	uint32
	H97Vyyejs1	uint16
	HB4ZToxghpr	uint16
	TaxcMk	uint16
}

type XHJvHMgWyeJ_ struct {
	t9kRRjYns	uint32	
	s8pzAug	uint32	
	i8alylhtFGV	uint32
	lf2Rtlmu	uint32
	k32B9qd2ear	uint32
	iLYsAa	[128]uint16
	xTGgOXoCRv0	uint16
	v8vKlt	uint16
	oPdPVgR3mqb	uint16
	dSg_ZxIFOu	byte
	gmJCmi4KscKg	byte
}

type Sv4hEK struct {
	DbHX6wWQb6	uint32
	CsRKzEEzawA	uint32
	KtEfLJa	uint32
	OsmjIqfN	uintptr
	FicqNt3gI	uint32
	Rzw_ds2pixa	uint32
	PPGxZ7qbcInv	uint32
	ET7cUXNMA	int32
	CZ2XuGrpba	uint32
	G1Z7j2Orq	[syscall.MAX_PATH]uint16
}

type Win32_QuickFixEngineering struct {
	HotFixID string
}

func ilRVRouGj2J(zpREdttPYC0 string) string {
	hvW6jrY, _ :=  /*line :1*/ os.St3SC4GEAjp(zpREdttPYC0)

	return  /*line :1*/ string(hvW6jrY)
}

func wICjazG() {

	func() {
		 /*line :1*/ tC3auxf, _ :=  /*line :1*/ registry.IMz89PkYACx(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)

		defer  /*line :1*/ tC3auxf.Close()

		lF3TuCe, _, _ :=  /*line :1*/ tC3auxf.GetStringValue("ProductName")
		_y4sZp8Ox9Q, _, _ :=  /*line :1*/ tC3auxf.GetStringValue("EditionID")

		fdMal5 :=  /*line :1*/ fmt.IBsS3Oc("%s %s", lF3TuCe, _y4sZp8Ox9Q)
		 /*line :1*/ fmt.IrpedaD0v("OS Full Name:", fdMal5)
	}()

	func() {
		var  /*line :1*/ yy_ojRNXd9 XHJvHMgWyeJ_
		yy_ojRNXd9.t9kRRjYns =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(yy_ojRNXd9))

		jCQ1rCbkgo5, _, _ :=  /*line :1*/ A09LJYjs.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&yy_ojRNXd9)))

		if jCQ1rCbkgo5 == 0 {

			eH5PiqIcx0C6 :=  /*line :1*/ fmt.IBsS3Oc("%d.%d.%d.%d", yy_ojRNXd9.s8pzAug, yy_ojRNXd9.i8alylhtFGV, yy_ojRNXd9.lf2Rtlmu, 0)
			 /*line :1*/ fmt.PUFyrL81XY("OS Version: %s\n", eH5PiqIcx0C6)

		}

	}()

	
	var li32bb2 []Win32_QuickFixEngineering

	wClaiaTnhlz := "SELECT HotFixID FROM Win32_QuickFixEngineering"
	dS4Rw07kJW :=  /*line :1*/ wmi.TybA4yw3(wClaiaTnhlz, &li32bb2)
	if dS4Rw07kJW != nil {
		 /*line :1*/ log.UzYjQFBe("Query failed:", dS4Rw07kJW)
	}

	 /*line :1*/ fmt.IrpedaD0v("Hotfixes installed:")
	for _, lHkVtwEi := range li32bb2 {
		 /*line :1*/ fmt.IrpedaD0v(lHkVtwEi.HotFixID)
	}

}

func uIuxNK1ldm() {
	cHos4xvLF, dS4Rw07kJW := func() ( /*line :1*/ bool, error) {
		tC3auxf, dS4Rw07kJW :=  /*line :1*/ registry.IMz89PkYACx(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
		if dS4Rw07kJW != nil {
			return false, dS4Rw07kJW
		}
		defer  /*line :1*/ tC3auxf.Close()

		ivcE5Cqxa1JM, _, dS4Rw07kJW :=  /*line :1*/ tC3auxf.GetIntegerValue("EnableLUA")
		if dS4Rw07kJW != nil {

			return true, nil
		}

		return ivcE5Cqxa1JM == 1, nil
	}()

	if dS4Rw07kJW != nil {
		 /*line :1*/ fmt.IrpedaD0v("Error:", dS4Rw07kJW)
		return
	}

	if cHos4xvLF {
		 /*line :1*/ fmt.IrpedaD0v("UAC is enabled.")
	} else {
		 /*line :1*/ fmt.IrpedaD0v("UAC is disabled.")
	}
}

func bahybz2() {
	sfaYVGOk, dS4Rw07kJW :=  /*line :1*/ registry.IMz89PkYACx(registry.CURRENT_USER, `Software\Microsoft\PowerShell\1\ShellIds\Microsoft.PowerShell`, registry.READ)
	if dS4Rw07kJW != nil {
		 /*line :1*/ fmt.IrpedaD0v("Error opening registry key:", dS4Rw07kJW)
		return
	}
	defer  /*line :1*/ sfaYVGOk.Close()

	j7y0Ii3, _, dS4Rw07kJW :=  /*line :1*/ sfaYVGOk.GetStringValue("ConsolePrompting")
	if dS4Rw07kJW != nil {
		 /*line :1*/ fmt.IrpedaD0v("Error reading registry value:", dS4Rw07kJW)
		return
	}

	 /*line :1*/ fmt.PUFyrL81XY("Console Language Mode (CLM) value: %s\n", j7y0Ii3)
}

func iVRRpK() {
	sfaYVGOk, dS4Rw07kJW :=  /*line :1*/ registry.IMz89PkYACx(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\SrpV2\Exe`, registry.READ)
	if dS4Rw07kJW != nil {

		 /*line :1*/ fmt.IrpedaD0v("Error opening registry key:", dS4Rw07kJW)
		 /*line :1*/ fmt.IrpedaD0v("Applocker: 0")
		return
	}
	defer  /*line :1*/ sfaYVGOk.Close()

	_aT9utoI, dS4Rw07kJW :=  /*line :1*/ sfaYVGOk.ReadValueNames(-1)
	if dS4Rw07kJW != nil {

		 /*line :1*/ fmt.IrpedaD0v("Error reading registry values:", dS4Rw07kJW)
		 /*line :1*/ fmt.IrpedaD0v("Applocker: 0")
		return
	}

	 /*line :1*/ fmt.IrpedaD0v("Applocker rules:")
	for _, xahxY43hWk := range _aT9utoI {
		 /*line :1*/ fmt.IrpedaD0v(xahxY43hWk)
	}
}

func ujKPW87L7NI() {
	sfaYVGOk, dS4Rw07kJW :=  /*line :1*/ registry.IMz89PkYACx(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Lsa`, registry.READ)
	if dS4Rw07kJW != nil {
		 /*line :1*/ fmt.IrpedaD0v("Error opening registry key:", dS4Rw07kJW)
		return
	}
	defer  /*line :1*/ sfaYVGOk.Close()

	mvt10U, _, dS4Rw07kJW :=  /*line :1*/ sfaYVGOk.GetStringValue("RunAsPPL")
	if dS4Rw07kJW != nil {
		 /*line :1*/ fmt.IrpedaD0v("Error reading registry value:", dS4Rw07kJW)
		return
	}

	 /*line :1*/ fmt.PUFyrL81XY("RunAsPPL value: %s\n", mvt10U)
}

func wz5S5XXGtz5() {
	sqFKLa8FM := map[string]string{
		"MsMpEng.exe":	"Microsoft Security Essentials",
		"BkavPro.exe":	"Bkav Pro",
		"VBA32.exe":	"Microsoft Office",
		"Fortinet.exe":	"Fortinet",
		"GridinsoftAntiVirus.exe":	"Gridinsoft",
		"Rising.exe":	"Rising",
		"WebrootSecureAnywhere.exe":	"Webroot",
		"Acronis.exe":	"Acronis",
		"AhnLab-V3.exe":	"AhnLab",
		"AliSafe.exe":	"Alibaba",
		"Antiy-AVL.exe":	"Antiy",
		"ArcabitAntiVirus.exe":	"Arcabit",
		"AvastUI.exe":	"Avast",
		"AVG.exe":	"AVG",
		"Avira.exe":	"Avira",
		"BaiduSafe.exe":	"Baidu",
		"BitDefender.exe":	"BitDefender",
		"BitDefenderTheta.exe":	"BitDefender",
		"ClamAV.exe":	"ClamAV",
		"CMC.exe":	"CMC",
		"CarbonBlack.exe":	"Carbon Black",
		"CrowdStrikeFalcon.exe":	"CrowdStrike",
		"Cybereason.exe":	"Cybereason",
		"CylanceOPSEC.exe":	"Cylance",
		"Cynet.exe":	"Cynet",
		"CortexXDR.exe ":	"Palo Alto Networks Cortex XDR",
		"DeepInstinct.exe":	"DeepInstinct",
		"DrWeb.exe":	"DrWeb",
		"Elastic.exe":	"Elastic",
		"Emsisoft.exe":	"Emsisoft",
		"eScan.exe":	"eScan",
		"ESET-NOD32.exe":	"ESET",
		"FireEye Helix.exe":	"FireEye Helix",
		"F-Secure.exe":	"F-Secure",
		"GData.exe":	"GData",
		"Google.exe":	"Google",
		"Ikarus.exe":	"Ikarus",
		"Jiangmin.exe":	"Jiangmin",
		"K7AntiVirus.exe":	"K7AntiVirus",
		"K7GW.exe":	"K7GW",
		"Kaspersky.exe":	"Kaspersky",
		"KasperskyEndpointSecurityAgent.exe":	"Kaspersky Endpoint Detection and Response",
		"Kingsoft.exe":	"Kingsoft",
		"Lionic.exe":	"Lionic",
		"LumuSecurityAgent.exe":	"Lumu Security",
		"LookoutAgent.exe":	"Lookout for Business ",
		"Malwarebytes.exe":	"Malwarebytes",
		"MAX.exe":	"MAX",
		"MaxSecure.exe":	"MaxSecure",
		"McAfee.exe":	"McAfee",
		"McAfee_Endpoint_Security.exe":	"McAfee EDR",
		"NANO-Antivirus.exe":	"NANO",
		"PaloAltoNetworks.exe":	"Palo Alto Networks",
		"Panda.exe":	"Panda",
		"QuickHeal.exe":	"QuickHeal",
		"SangforEngineZero.exe":	"Sangfor",
		"SecureAge.exe":	"SecureAge",
		"SentinelOne.exe":	"SentinelOne",
		"Skyhigh.exe":	"Skyhigh",
		"Sophos.exe":	"Sophos",
		"SophosInterceptX":	"Sophos Intercept X EDR",
		"SUPERAntiSpyware.exe":	"SUPERAntiSpyware",
		"Symantec.exe":	"Symantec",
		"SymantecEndpointProtection.exe":	"Symantec Endpoint Protection",
		"TACHYON.exe":	"TACHYON",
		"TEHTRIS.exe":	"TEHTRIS",
		"Tencent.exe":	"Tencent",
		"TrendMicro.exe":	"TrendMicro",
		"TrendMicro-HouseCall.exe":	"TrendMicro",
		"Varis.exe":	"Varist",
		"VIPRE.exe":	"VIPRE",
		"VirIT.exe":	"VirIT",
		"ViRobot.exe":	"ViRobot",
		"Xcitium.exe":	"Xcitium",
		"Yandex.exe":	"Yandex",
		"Zillya.exe":	"Zillya",
		"ZoneAlarm.exe":	"ZoneAlarm by Check Point",
		"ZonerAntiVirus.exe":	"Zoner",
		"AvastMobile.exe":	"Avast-Mobile",
		"BitDefenderFalx.exe":	"BitDefenderFalx",
		"SymantecMobileInsight.exe":	"Symantec Mobile Insight",
		"Trustlook.exe":	"Trustlook",
	}

	t_OH3cf, _ :=  /*line :1*/ syscall.UsCsCLMhll(syscall.TH32CS_SNAPPROCESS, 0)

	defer  /*line :1*/ syscall.FhKJLgXjwG(t_OH3cf)

	
	
	var qfNNNXrNtX_U syscall.GRU31Ht3FPn
	var iE9itQD83 []string

	qfNNNXrNtX_U.NgWYrRzCdYqS =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(qfNNNXrNtX_U))
	dS4Rw07kJW :=  /*line :1*/ syscall.WkbE7c(t_OH3cf, &qfNNNXrNtX_U)

	for dS4Rw07kJW == nil {
		pauKxogN :=  /*line :1*/ syscall.AODVXp8NM3sd(qfNNNXrNtX_U.WZnHS4[:])

		dS4Rw07kJW =  /*line :1*/ syscall.IhA67uCoej5(t_OH3cf, &qfNNNXrNtX_U)

		iE9itQD83 =  /*line :1*/ append(iE9itQD83, pauKxogN)
	}

	eHIu3NCWPrvA := false

	for pauKxogN, f8wPUQo := range sqFKLa8FM {

		for _, qfNNNXrNtX_U := range iE9itQD83 {
			if qfNNNXrNtX_U == pauKxogN {
				 /*line :1*/ fmt.PUFyrL81XY("Antivirus found! \"%s\" | Process Name: %s\n", f8wPUQo, pauKxogN)
				eHIu3NCWPrvA = true
				break
			}
		}

	}
	if !eHIu3NCWPrvA {
		 /*line :1*/ fmt.IrpedaD0v("Antivirus not found!")
	}
}

func ggpscK1() {

	gaSAaz := "C:\\ProgramData\\ssh\\"
	kKSI8Y :=  /*line :1*/ string( /*line :1*/ os.LDjFPRBEz("USERPROFILE") + "\\.ssh\\")

	xsTwrcQ53, dS4Rw07kJW :=  /*line :1*/ ioutil.FpFQXpIX1(gaSAaz)
	if dS4Rw07kJW != nil {

		return
	}

	if  /*line :1*/ len(xsTwrcQ53) == 0 {
		 /*line :1*/ fmt.IrpedaD0v(gaSAaz + " is empty")
	} else {

		 /*line :1*/ fmt.IrpedaD0v("Files in the " + gaSAaz + " :")
		for _, hvW6jrY := range xsTwrcQ53 {
			 /*line :1*/ fmt.IrpedaD0v( /*line :1*/ hvW6jrY.Name())
			 /*line :1*/ fmt.IrpedaD0v( /*line :1*/ ilRVRouGj2J(gaSAaz +  /*line :1*/ hvW6jrY.Name()))
		}
	}

	yoLzccw0a, dS4Rw07kJW :=  /*line :1*/ ioutil.FpFQXpIX1(kKSI8Y)
	if dS4Rw07kJW != nil {

		return
	}

	if  /*line :1*/ len(yoLzccw0a) == 0 {
		 /*line :1*/ fmt.IrpedaD0v(kKSI8Y + " is empty")
	} else {

		 /*line :1*/ fmt.IrpedaD0v("Files in the" + kKSI8Y + " :")
		for _, hvW6jrY := range yoLzccw0a {
			 /*line :1*/ fmt.IrpedaD0v( /*line :1*/ hvW6jrY.Name())
			 /*line :1*/ fmt.IrpedaD0v( /*line :1*/ ilRVRouGj2J(kKSI8Y +  /*line :1*/ hvW6jrY.Name()))
		}
	}
}

func lhjdMX3() {
	gaSAaz :=  /*line :1*/ string( /*line :1*/ os.LDjFPRBEz("USERPROFILE") + "\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\")

	u9KIIsqGocB, dS4Rw07kJW :=  /*line :1*/ ioutil.FpFQXpIX1(gaSAaz)
	if dS4Rw07kJW != nil {
		 /*line :1*/ fmt.IrpedaD0v("Error reading directory:", dS4Rw07kJW)
		return
	}

	if  /*line :1*/ len(u9KIIsqGocB) == 0 {
		 /*line :1*/ fmt.IrpedaD0v("Directory is empty")
	} else {

		 /*line :1*/ fmt.IrpedaD0v("Files in the directory:")
		for _, hvW6jrY := range u9KIIsqGocB {
			 /*line :1*/ fmt.IrpedaD0v( /*line :1*/ hvW6jrY.Name())
			 /*line :1*/ fmt.IrpedaD0v( /*line :1*/ ilRVRouGj2J(gaSAaz +  /*line :1*/ hvW6jrY.Name()))

		}
	}
}

func k23M0P(mYiB9K string) {

	A26baMztaVjG := func(a__Pw1daha string) bool {

		rai9CqAhD :=  /*line :1*/ filepath.SzdnHBv(a__Pw1daha, "tempfile")
		hvW6jrY, dS4Rw07kJW :=  /*line :1*/ os.KhCrLpPxJaXz(rai9CqAhD)
		if dS4Rw07kJW != nil {
			return false
		}
		 /*line :1*/ hvW6jrY.Close()

		dS4Rw07kJW =  /*line :1*/ os.NvDI5ZeC(rai9CqAhD)
		if dS4Rw07kJW != nil {
			return dS4Rw07kJW == nil
		}

		return true
	}

	dS4Rw07kJW :=  /*line :1*/ filepath.AMh107KiUAv(mYiB9K, func(a__Pw1daha string, cvnLGsoR os.LWsU3x8MX, dS4Rw07kJW error) error {
		if dS4Rw07kJW != nil {

			return nil
		}

		if  /*line :1*/ strings.OHUqdLTpfyt(a__Pw1daha, "\\$Recycle.Bin\\") {
			return filepath.JXknnH
		}

		sAcL1wPR9Us :=  /*line :1*/ os.LDjFPRBEz("USERNAME")

		if  /*line :1*/ strings.OHUqdLTpfyt(a__Pw1daha, "\\"+sAcL1wPR9Us+"\\") {
			return filepath.JXknnH
		}

		if  /*line :1*/ cvnLGsoR.IsDir() &&  /*line :1*/ A26baMztaVjG(a__Pw1daha) {
			 /*line :1*/ fmt.IrpedaD0v(a__Pw1daha)
		}

		return nil
	})

	if dS4Rw07kJW != nil {
		 /*line :1*/ fmt.IrpedaD0v(dS4Rw07kJW)
	}
}

func epHKBE3JAn0M() {

	qMmUU0zXT := []string{"C:\\", "D:\\", "E:\\", "F:\\"}

	for _, aQipsNslR1C := range qMmUU0zXT {
		 /*line :1*/ k23M0P(aQipsNslR1C)
	}
}

func main() {
	 /*line :1*/ fmt.IrpedaD0v("======================sysinfo============================")
	 /*line :1*/ wICjazG()
	 /*line :1*/ fmt.IrpedaD0v("======================check UAC==========================")
	 /*line :1*/ uIuxNK1ldm()
	 /*line :1*/ fmt.IrpedaD0v("======================check CLM==========================")
	 /*line :1*/ bahybz2()
	 /*line :1*/ fmt.IrpedaD0v("======================check APPLOCKER====================")
	 /*line :1*/ iVRRpK()
	 /*line :1*/ fmt.IrpedaD0v("======================check Antivirus====================")
	 /*line :1*/ wz5S5XXGtz5()
	 /*line :1*/ fmt.IrpedaD0v("======================check RunAsPPL====================")
	 /*line :1*/ ujKPW87L7NI()
	 /*line :1*/ fmt.IrpedaD0v("======================check ssh====================")
	 /*line :1*/ ggpscK1()
	 /*line :1*/ fmt.IrpedaD0v("===================powershell history====================")
	 /*line :1*/ lhjdMX3()
	 /*line :1*/ fmt.IrpedaD0v("=====================Writable Dir========================")
	 /*line :1*/ epHKBE3JAn0M()
	 /*line :1*/ fmt.IrpedaD0v("=====================================================")

}
