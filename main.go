package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

const (
	TH32CS_SNAPPROCESS = 0x00000002
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	Ntdll    = syscall.NewLazyDLL("Ntdll.dll")
	CIMWin32 = syscall.NewLazyDLL("CIMWin32.dll")

	RtlGetVersion = Ntdll.NewProc("RtlGetVersion")

	CreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	Process32First           = kernel32.NewProc("Process32First")
	Process32Next            = kernel32.NewProc("Process32Next")
)

// SYSTEM_INFO structure
type SYSTEM_INFO struct {
	ProcessorArchitecture     uint16
	Reserved                  uint16
	PageSize                  uint32
	MinimumApplicationAddress uintptr
	MaximumApplicationAddress uintptr
	ActiveProcessorMask       uintptr
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint16
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

type RTL_OSVERSIONINFOEX struct {
	dwOSVersionInfoSize uint32 //此資料結構的大小，以位元組為單位。
	dwMajorVersion      uint32 //作業系統的主要版本號
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
	szCSDVersion        [128]uint16
	wServicePackMajor   uint16
	wServicePackMinor   uint16
	wSuiteMask          uint16
	wProductType        byte
	wReserved           byte
}

type PROCESSENTRY32 struct {
	Size              uint32
	Usage             uint32
	ProcessID         uint32
	DefaultHeapID     uintptr
	ModuleID          uint32
	CountThreads      uint32
	ParentProcessID   uint32
	PriorityClassBase int32
	Flags             uint32
	ExeFile           [syscall.MAX_PATH]uint16
}

type Win32_QuickFixEngineering struct {
	HotFixID string
}

func readfile(filename string) string {
	file, _ := os.ReadFile(filename)
	//if err != nil {
	//	return ""
	//}

	return string(file)
}

func sysinfo() {

	func() {
		k, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)

		defer k.Close()

		productName, _, _ := k.GetStringValue("ProductName")
		editionID, _, _ := k.GetStringValue("EditionID")

		full_name := fmt.Sprintf("%s %s", productName, editionID)
		fmt.Println("OS Full Name:", full_name)
	}()

	func() { //os version
		var osVersionInfo RTL_OSVERSIONINFOEX
		osVersionInfo.dwOSVersionInfoSize = uint32(unsafe.Sizeof(osVersionInfo))

		ret, _, _ := RtlGetVersion.Call(uintptr(unsafe.Pointer(&osVersionInfo)))

		if ret == 0 {

			osVersion := fmt.Sprintf("%d.%d.%d.%d", osVersionInfo.dwMajorVersion, osVersionInfo.dwMinorVersion, osVersionInfo.dwBuildNumber, 0)
			fmt.Printf("OS Version: %s\n", osVersion)

		}

	}()

	//get hotfix
	var hotfixes []Win32_QuickFixEngineering

	query := "SELECT HotFixID FROM Win32_QuickFixEngineering"
	err := wmi.Query(query, &hotfixes)
	if err != nil {
		log.Fatal("Query failed:", err)
	}

	fmt.Println("Hotfixes installed:")
	for _, hotfix := range hotfixes {
		fmt.Println(hotfix.HotFixID)
	}

}

func check_UAC() {
	uacEnabled, err := func() (bool, error) {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.QUERY_VALUE)
		if err != nil {
			return false, err
		}
		defer k.Close()

		// 获取 EnableLUA 的值，如果不存在默认为 1（启用）
		enableLUA, _, err := k.GetIntegerValue("EnableLUA")
		if err != nil {
			// 如果键不存在，默认为启用
			return true, nil
		}

		// 如果 EnableLUA 的值为 1，则 UAC 启用；否则，禁用
		return enableLUA == 1, nil
	}()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if uacEnabled {
		fmt.Println("UAC is enabled.")
	} else {
		fmt.Println("UAC is disabled.")
	}
}

func check_CLM() {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\PowerShell\1\ShellIds\Microsoft.PowerShell`, registry.READ)
	if err != nil {
		fmt.Println("Error opening registry key:", err)
		return
	}
	defer key.Close()

	clm, _, err := key.GetStringValue("ConsolePrompting")
	if err != nil {
		fmt.Println("Error reading registry value:", err)
		return
	}

	fmt.Printf("Console Language Mode (CLM) value: %s\n", clm)
}

func check_APPLOCKER() {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\SrpV2\Exe`, registry.READ)
	if err != nil {
		// 打印错误信息
		fmt.Println("Error opening registry key:", err)
		fmt.Println("Applocker: 0")
		return
	}
	defer key.Close()

	// 读取子项的名称
	names, err := key.ReadValueNames(-1)
	if err != nil {
		// 打印错误信息
		fmt.Println("Error reading registry values:", err)
		fmt.Println("Applocker: 0")
		return
	}

	// 打印子项名称
	fmt.Println("Applocker rules:")
	for _, name := range names {
		fmt.Println(name)
	}
}

func check_RunAsPPL() {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Lsa`, registry.READ)
	if err != nil {
		fmt.Println("Error opening registry key:", err)
		return
	}
	defer key.Close()

	runAsPPL, _, err := key.GetStringValue("RunAsPPL")
	if err != nil {
		fmt.Println("Error reading registry value:", err)
		return
	}

	fmt.Printf("RunAsPPL value: %s\n", runAsPPL)
}

func check_AV() {
	processToAntivirus := map[string]string{
		"MsMpEng.exe":                        "Microsoft Security Essentials",
		"BkavPro.exe":                        "Bkav Pro",
		"VBA32.exe":                          "Microsoft Office",
		"Fortinet.exe":                       "Fortinet",
		"GridinsoftAntiVirus.exe":            "Gridinsoft",
		"Rising.exe":                         "Rising",
		"WebrootSecureAnywhere.exe":          "Webroot",
		"Acronis.exe":                        "Acronis",
		"AhnLab-V3.exe":                      "AhnLab",
		"AliSafe.exe":                        "Alibaba",
		"Antiy-AVL.exe":                      "Antiy",
		"ArcabitAntiVirus.exe":               "Arcabit",
		"AvastUI.exe":                        "Avast",
		"AVG.exe":                            "AVG",
		"Avira.exe":                          "Avira",
		"BaiduSafe.exe":                      "Baidu",
		"BitDefender.exe":                    "BitDefender",
		"BitDefenderTheta.exe":               "BitDefender",
		"ClamAV.exe":                         "ClamAV",
		"CMC.exe":                            "CMC",
		"CarbonBlack.exe":                    "Carbon Black",
		"CrowdStrikeFalcon.exe":              "CrowdStrike",
		"Cybereason.exe":                     "Cybereason",
		"CylanceOPSEC.exe":                   "Cylance",
		"Cynet.exe":                          "Cynet",
		"CortexXDR.exe ":                     "Palo Alto Networks Cortex XDR",
		"DeepInstinct.exe":                   "DeepInstinct",
		"DrWeb.exe":                          "DrWeb",
		"Elastic.exe":                        "Elastic",
		"Emsisoft.exe":                       "Emsisoft",
		"eScan.exe":                          "eScan",
		"ESET-NOD32.exe":                     "ESET",
		"FireEye Helix.exe":                  "FireEye Helix",
		"F-Secure.exe":                       "F-Secure",
		"GData.exe":                          "GData",
		"Google.exe":                         "Google",
		"Ikarus.exe":                         "Ikarus",
		"Jiangmin.exe":                       "Jiangmin",
		"K7AntiVirus.exe":                    "K7AntiVirus",
		"K7GW.exe":                           "K7GW",
		"Kaspersky.exe":                      "Kaspersky",
		"KasperskyEndpointSecurityAgent.exe": "Kaspersky Endpoint Detection and Response",
		"Kingsoft.exe":                       "Kingsoft",
		"Lionic.exe":                         "Lionic",
		"LumuSecurityAgent.exe":              "Lumu Security",
		"LookoutAgent.exe":                   "Lookout for Business ",
		"Malwarebytes.exe":                   "Malwarebytes",
		"MAX.exe":                            "MAX",
		"MaxSecure.exe":                      "MaxSecure",
		"McAfee.exe":                         "McAfee",
		"McAfee_Endpoint_Security.exe":       "McAfee EDR",
		"NANO-Antivirus.exe":                 "NANO",
		"PaloAltoNetworks.exe":               "Palo Alto Networks",
		"Panda.exe":                          "Panda",
		"QuickHeal.exe":                      "QuickHeal",
		"SangforEngineZero.exe":              "Sangfor",
		"SecureAge.exe":                      "SecureAge",
		"SentinelOne.exe":                    "SentinelOne",
		"Skyhigh.exe":                        "Skyhigh",
		"Sophos.exe":                         "Sophos",
		"SophosInterceptX":                   "Sophos Intercept X EDR",
		"SUPERAntiSpyware.exe":               "SUPERAntiSpyware",
		"Symantec.exe":                       "Symantec",
		"SymantecEndpointProtection.exe":     "Symantec Endpoint Protection",
		"TACHYON.exe":                        "TACHYON",
		"TEHTRIS.exe":                        "TEHTRIS",
		"Tencent.exe":                        "Tencent",
		"TrendMicro.exe":                     "TrendMicro",
		"TrendMicro-HouseCall.exe":           "TrendMicro",
		"Varis.exe":                          "Varist",
		"VIPRE.exe":                          "VIPRE",
		"VirIT.exe":                          "VirIT",
		"ViRobot.exe":                        "ViRobot",
		"Xcitium.exe":                        "Xcitium",
		"Yandex.exe":                         "Yandex",
		"Zillya.exe":                         "Zillya",
		"ZoneAlarm.exe":                      "ZoneAlarm by Check Point",
		"ZonerAntiVirus.exe":                 "Zoner",
		"AvastMobile.exe":                    "Avast-Mobile",
		"BitDefenderFalx.exe":                "BitDefenderFalx",
		"SymantecMobileInsight.exe":          "Symantec Mobile Insight",
		"Trustlook.exe":                      "Trustlook",
	}

	handle, _ := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)

	defer syscall.CloseHandle(handle)

	// handle, _, _ := CreateToolhelp32Snapshot.Call(TH32CS_SNAPPROCESS, 0)
	// defer syscall.CloseHandle(syscall.Handle(handle))

	// // 初始化PROCESSENTRY32结构体
	//var entry  PROCESSENTRY32 //
	var entry syscall.ProcessEntry32
	var entries []string

	entry.Size = uint32(unsafe.Sizeof(entry))
	err := syscall.Process32First(handle, &entry)

	for err == nil {
		processName := syscall.UTF16ToString(entry.ExeFile[:])
		//fmt.Printf("Process ID: %d, Name: %s\n", entry.ProcessID, processName)
		err = syscall.Process32Next(handle, &entry)

		entries = append(entries, processName)
	}

	// err = syscall.Process32First(syscall.Handle(handle), (&pe32))
	// for err == nil {
	// 	fmt.Printf("Process ID: %d, Name: %s\n", pe32.ProcessID, syscall.UTF16ToString(pe32.ExeFile[:]))
	// 	err = syscall.Process32Next(syscall.Handle(handle), &pe32)
	// }

	// tasklistStr := string(tasklistOutput)
	found := false
	// 提取 Image Name 部分并直接进行比较
	for processName, antivirus := range processToAntivirus {
		// 使用循环检查切片中是否包含进程名

		for _, entry := range entries {
			if entry == processName {
				fmt.Printf("Antivirus found! \"%s\" | Process Name: %s\n", antivirus, processName)
				found = true
				break
			}
		}

	}
	if !found {
		fmt.Println("Antivirus not found!")
	}
}

func check_ssh() {
	// 指定目标目录路径
	targetDir := "C:\\ProgramData\\ssh\\"
	targetDir1 := string(os.Getenv("USERPROFILE") + "\\.ssh\\")
	//C:\Users\x_x\.ssh
	// 获取目录下的文件列表
	file1, err := ioutil.ReadDir(targetDir)
	if err != nil {

		return
	}

	// 检查文件列表是否为空
	if len(file1) == 0 {
		fmt.Println(targetDir + " is empty")
	} else {
		// 打印文件名
		fmt.Println("Files in the " + targetDir + " :")
		for _, file := range file1 {
			fmt.Println(file.Name())
			fmt.Println(readfile(targetDir + file.Name()))
		}
	}

	file2, err := ioutil.ReadDir(targetDir1)
	if err != nil {

		return
	}

	// 检查文件列表是否为空
	if len(file2) == 0 {
		fmt.Println(targetDir1 + " is empty")
	} else {
		// 打印文件名
		fmt.Println("Files in the" + targetDir1 + " :")
		for _, file := range file2 {
			fmt.Println(file.Name())
			fmt.Println(readfile(targetDir1 + file.Name()))
		}
	}
}

func ps_history() {
	targetDir := string(os.Getenv("USERPROFILE") + "\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\")
	// 获取目录下的文件列表
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// 检查文件列表是否为空
	if len(files) == 0 {
		fmt.Println("Directory is empty")
	} else {
		// 打印文件名
		fmt.Println("Files in the directory:")
		for _, file := range files {
			fmt.Println(file.Name())
			fmt.Println(readfile(targetDir + file.Name()))

		}
	}
}

func listWritableDirectoriesOnDrive(rootDir string) {

	YorN := func(path string) bool {
		// 尝试在目录中创建一个临时文件
		tempFileName := filepath.Join(path, "tempfile")
		file, err := os.Create(tempFileName)
		if err != nil {
			return false
		}
		file.Close()

		// 删除临时文件
		err = os.Remove(tempFileName)
		if err != nil {
			return err == nil
		}

		return true
	}
	// 使用filepath.Walk递归遍历目录
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			//fmt.Printf("Error walking path %s: %v\n", path, err)
			return nil // 继续遍历其他目录
		}

		// 排除指定目录（例如C:\$Recycle.Bin\）
		if strings.Contains(path, "\\$Recycle.Bin\\") {
			return filepath.SkipDir
		}

		username := os.Getenv("USERNAME")
		// 排除指定目录（例如C:\$Recycle.Bin\）
		if strings.Contains(path, "\\"+username+"\\") {
			return filepath.SkipDir
		}

		// 判断是目录且可写
		if info.IsDir() && YorN(path) {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func write_or_not() {
	// 获取所有硬盘根目录
	drives := []string{"C:\\", "D:\\", "E:\\", "F:\\"} // 可以根据实际情况添加更多的硬盘字母

	// 遍历所有硬盘并列出可写目录
	for _, drive := range drives {
		listWritableDirectoriesOnDrive(drive)
	}
}

func main() {
	fmt.Println("======================sysinfo============================")
	sysinfo()
	fmt.Println("======================check UAC==========================")
	check_UAC()
	fmt.Println("======================check CLM==========================")
	check_CLM()
	fmt.Println("======================check APPLOCKER====================")
	check_APPLOCKER()
	fmt.Println("======================check Antivirus====================")
	check_AV()
	fmt.Println("======================check RunAsPPL====================")
	check_RunAsPPL()
	fmt.Println("======================check ssh====================")
	check_ssh()
	fmt.Println("===================powershell history====================")
	ps_history()
	fmt.Println("=====================Writable Dir========================")
	write_or_not()
	fmt.Println("=====================================================")

}
