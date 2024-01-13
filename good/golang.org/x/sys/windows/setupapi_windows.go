//line :1
package NJ4MerJ

import (
	binary "yLm9uBQG"
	errors "iAHoxjmM"
	fmt "kFsoCfy5zWG"
	"runtime"
	strings "fQXlzv"
	syscall "bUKeamF"
	"unsafe"
)

const (
	ERROR_EXPECTED_SECTION_NAME	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0
	ERROR_BAD_SECTION_NAME_LINE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 1
	ERROR_SECTION_NAME_TOO_LONG	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 2
	ERROR_GENERAL_SYNTAX	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 3
	ERROR_WRONG_INF_STYLE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x100
	ERROR_SECTION_NOT_FOUND	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x101
	ERROR_LINE_NOT_FOUND	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x102
	ERROR_NO_BACKUP	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x103
	ERROR_NO_ASSOCIATED_CLASS	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x200
	ERROR_CLASS_MISMATCH	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x201
	ERROR_DUPLICATE_FOUND	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x202
	ERROR_NO_DRIVER_SELECTED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x203
	ERROR_KEY_DOES_NOT_EXIST	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x204
	ERROR_INVALID_DEVINST_NAME	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x205
	ERROR_INVALID_CLASS	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x206
	ERROR_DEVINST_ALREADY_EXISTS	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x207
	ERROR_DEVINFO_NOT_REGISTERED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x208
	ERROR_INVALID_REG_PROPERTY	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x209
	ERROR_NO_INF	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x20A
	ERROR_NO_SUCH_DEVINST	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x20B
	ERROR_CANT_LOAD_CLASS_ICON	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x20C
	ERROR_INVALID_CLASS_INSTALLER	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x20D
	ERROR_DI_DO_DEFAULT	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x20E
	ERROR_DI_NOFILECOPY	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x20F
	ERROR_INVALID_HWPROFILE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x210
	ERROR_NO_DEVICE_SELECTED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x211
	ERROR_DEVINFO_LIST_LOCKED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x212
	ERROR_DEVINFO_DATA_LOCKED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x213
	ERROR_DI_BAD_PATH	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x214
	ERROR_NO_CLASSINSTALL_PARAMS	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x215
	ERROR_FILEQUEUE_LOCKED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x216
	ERROR_BAD_SERVICE_INSTALLSECT	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x217
	ERROR_NO_CLASS_DRIVER_LIST	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x218
	ERROR_NO_ASSOCIATED_SERVICE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x219
	ERROR_NO_DEFAULT_DEVICE_INTERFACE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x21A
	ERROR_DEVICE_INTERFACE_ACTIVE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x21B
	ERROR_DEVICE_INTERFACE_REMOVED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x21C
	ERROR_BAD_INTERFACE_INSTALLSECT	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x21D
	ERROR_NO_SUCH_INTERFACE_CLASS	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x21E
	ERROR_INVALID_REFERENCE_STRING	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x21F
	ERROR_INVALID_MACHINENAME	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x220
	ERROR_REMOTE_COMM_FAILURE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x221
	ERROR_MACHINE_UNAVAILABLE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x222
	ERROR_NO_CONFIGMGR_SERVICES	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x223
	ERROR_INVALID_PROPPAGE_PROVIDER	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x224
	ERROR_NO_SUCH_DEVICE_INTERFACE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x225
	ERROR_DI_POSTPROCESSING_REQUIRED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x226
	ERROR_INVALID_COINSTALLER	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x227
	ERROR_NO_COMPAT_DRIVERS	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x228
	ERROR_NO_DEVICE_ICON	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x229
	ERROR_INVALID_INF_LOGCONFIG	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x22A
	ERROR_DI_DONT_INSTALL	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x22B
	ERROR_INVALID_FILTER_DRIVER	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x22C
	ERROR_NON_WINDOWS_NT_DRIVER	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x22D
	ERROR_NON_WINDOWS_DRIVER	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x22E
	ERROR_NO_CATALOG_FOR_OEM_INF	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x22F
	ERROR_DEVINSTALL_QUEUE_NONNATIVE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x230
	ERROR_NOT_DISABLEABLE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x231
	ERROR_CANT_REMOVE_DEVINST	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x232
	ERROR_INVALID_TARGET	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x233
	ERROR_DRIVER_NONNATIVE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x234
	ERROR_IN_WOW64	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x235
	ERROR_SET_SYSTEM_RESTORE_POINT	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x236
	ERROR_SCE_DISABLED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x238
	ERROR_UNKNOWN_EXCEPTION	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x239
	ERROR_PNP_REGISTRY_ERROR	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x23A
	ERROR_REMOTE_REQUEST_UNSUPPORTED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x23B
	ERROR_NOT_AN_INSTALLED_OEM_INF	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x23C
	ERROR_INF_IN_USE_BY_DEVICES	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x23D
	ERROR_DI_FUNCTION_OBSOLETE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x23E
	ERROR_NO_AUTHENTICODE_CATALOG	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x23F
	ERROR_AUTHENTICODE_DISALLOWED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x240
	ERROR_AUTHENTICODE_TRUSTED_PUBLISHER	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x241
	ERROR_AUTHENTICODE_TRUST_NOT_ESTABLISHED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x242
	ERROR_AUTHENTICODE_PUBLISHER_NOT_TRUSTED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x243
	ERROR_SIGNATURE_OSATTRIBUTE_MISMATCH	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x244
	ERROR_ONLY_VALIDATE_VIA_AUTHENTICODE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x245
	ERROR_DEVICE_INSTALLER_NOT_READY	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x246
	ERROR_DRIVER_STORE_ADD_FAILED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x247
	ERROR_DEVICE_INSTALL_BLOCKED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x248
	ERROR_DRIVER_INSTALL_BLOCKED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x249
	ERROR_WRONG_INF_TYPE	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x24A
	ERROR_FILE_HASH_NOT_IN_CATALOG	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x24B
	ERROR_DRIVER_STORE_DELETE_FAILED	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x24C
	ERROR_UNRECOVERABLE_STACK_OVERFLOW	W5PDg27zDC	= 0x20000000 | 0xC0000000 | 0x300
	EXCEPTION_SPAPI_UNRECOVERABLE_STACK_OVERFLOW	W5PDg27zDC	= ERROR_UNRECOVERABLE_STACK_OVERFLOW
	ERROR_NO_DEFAULT_INTERFACE_DEVICE	W5PDg27zDC	= ERROR_NO_DEFAULT_DEVICE_INTERFACE
	ERROR_INTERFACE_DEVICE_ACTIVE	W5PDg27zDC	= ERROR_DEVICE_INTERFACE_ACTIVE
	ERROR_INTERFACE_DEVICE_REMOVED	W5PDg27zDC	= ERROR_DEVICE_INTERFACE_REMOVED
	ERROR_NO_SUCH_INTERFACE_DEVICE	W5PDg27zDC	= ERROR_NO_SUCH_DEVICE_INTERFACE
)

const (
	MAX_DEVICE_ID_LEN	= 200
	MAX_DEVNODE_ID_LEN	= MAX_DEVICE_ID_LEN
	MAX_GUID_STRING_LEN	= 39		
	MAX_CLASS_NAME_LEN	= 32
	MAX_PROFILE_LEN	= 80
	MAX_CONFIG_VALUE	= 9999
	MAX_INSTANCE_VALUE	= 9999
	CONFIGMG_VERSION	= 0x0400
)


const (
	LINE_LEN	= 256		
	MAX_INF_STRING_LENGTH	= 4096		
	MAX_INF_SECTION_NAME_LENGTH	= 255		
	MAX_TITLE_LEN	= 60
	MAX_INSTRUCTION_LEN	= 256
	MAX_LABEL_LEN	= 30
	MAX_SERVICE_NAME_LEN	= 256
	MAX_SUBTITLE_LEN	= 256
)

const (
	
	SP_MAX_MACHINENAME_LENGTH = MAX_PATH + 3
)


type HPswIcPV29vq uintptr


type WIItFoK L2L8P9WaNZ


type Dn3mHAPt5D uint32


type BKGe6zA1I1c_ struct {
	ftvJsHy_ex1F	uint32
	UrUM1jHmB47	Kw2qafRFaiLg
	M5sANznGA	Dn3mHAPt5D
	_	uintptr
}


type F9bfgg struct {
	cD48bm	uint32	
	CUfLs57EIp4l	Kw2qafRFaiLg
	EH_trjJfsNe	L2L8P9WaNZ
	d0riHuVa9R	[SP_MAX_MACHINENAME_LENGTH]uint16
}

func (*F9bfgg) rm3s25Ys() uint32 {
	if  /*line :1*/ unsafe.Sizeof( /*line :1*/ uintptr(0)) == 4 {

		return  /*line :1*/ uint32( /*line :1*/ unsafe.Offsetof(F9bfgg{}.d0riHuVa9R) +  /*line :1*/ unsafe.Sizeof(F9bfgg{}.d0riHuVa9R))
	}
	return  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(F9bfgg{}))
}

func (iIzhNC *F9bfgg) RemoteMachineName() string {
	return  /*line :1*/ OXNanQ8Uj(iIzhNC.d0riHuVa9R[:])
}

func (iIzhNC *F9bfgg) SetRemoteMachineName(nxn57yETJ string) error {
	xukLmcNaR, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(nxn57yETJ)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	 /*line :1*/ copy(iIzhNC.d0riHuVa9R[:], xukLmcNaR)
	return nil
}


type JHgnG0YbHkld uint32

const (
	DIF_SELECTDEVICE	JHgnG0YbHkld	= 0x00000001
	DIF_INSTALLDEVICE	JHgnG0YbHkld	= 0x00000002
	DIF_ASSIGNRESOURCES	JHgnG0YbHkld	= 0x00000003
	DIF_PROPERTIES	JHgnG0YbHkld	= 0x00000004
	DIF_REMOVE	JHgnG0YbHkld	= 0x00000005
	DIF_FIRSTTIMESETUP	JHgnG0YbHkld	= 0x00000006
	DIF_FOUNDDEVICE	JHgnG0YbHkld	= 0x00000007
	DIF_SELECTCLASSDRIVERS	JHgnG0YbHkld	= 0x00000008
	DIF_VALIDATECLASSDRIVERS	JHgnG0YbHkld	= 0x00000009
	DIF_INSTALLCLASSDRIVERS	JHgnG0YbHkld	= 0x0000000A
	DIF_CALCDISKSPACE	JHgnG0YbHkld	= 0x0000000B
	DIF_DESTROYPRIVATEDATA	JHgnG0YbHkld	= 0x0000000C
	DIF_VALIDATEDRIVER	JHgnG0YbHkld	= 0x0000000D
	DIF_DETECT	JHgnG0YbHkld	= 0x0000000F
	DIF_INSTALLWIZARD	JHgnG0YbHkld	= 0x00000010
	DIF_DESTROYWIZARDDATA	JHgnG0YbHkld	= 0x00000011
	DIF_PROPERTYCHANGE	JHgnG0YbHkld	= 0x00000012
	DIF_ENABLECLASS	JHgnG0YbHkld	= 0x00000013
	DIF_DETECTVERIFY	JHgnG0YbHkld	= 0x00000014
	DIF_INSTALLDEVICEFILES	JHgnG0YbHkld	= 0x00000015
	DIF_UNREMOVE	JHgnG0YbHkld	= 0x00000016
	DIF_SELECTBESTCOMPATDRV	JHgnG0YbHkld	= 0x00000017
	DIF_ALLOW_INSTALL	JHgnG0YbHkld	= 0x00000018
	DIF_REGISTERDEVICE	JHgnG0YbHkld	= 0x00000019
	DIF_NEWDEVICEWIZARD_PRESELECT	JHgnG0YbHkld	= 0x0000001A
	DIF_NEWDEVICEWIZARD_SELECT	JHgnG0YbHkld	= 0x0000001B
	DIF_NEWDEVICEWIZARD_PREANALYZE	JHgnG0YbHkld	= 0x0000001C
	DIF_NEWDEVICEWIZARD_POSTANALYZE	JHgnG0YbHkld	= 0x0000001D
	DIF_NEWDEVICEWIZARD_FINISHINSTALL	JHgnG0YbHkld	= 0x0000001E
	DIF_INSTALLINTERFACES	JHgnG0YbHkld	= 0x00000020
	DIF_DETECTCANCEL	JHgnG0YbHkld	= 0x00000021
	DIF_REGISTER_COINSTALLERS	JHgnG0YbHkld	= 0x00000022
	DIF_ADDPROPERTYPAGE_ADVANCED	JHgnG0YbHkld	= 0x00000023
	DIF_ADDPROPERTYPAGE_BASIC	JHgnG0YbHkld	= 0x00000024
	DIF_TROUBLESHOOTER	JHgnG0YbHkld	= 0x00000026
	DIF_POWERMESSAGEWAKE	JHgnG0YbHkld	= 0x00000027
	DIF_ADDREMOTEPROPERTYPAGE_ADVANCED	JHgnG0YbHkld	= 0x00000028
	DIF_UPDATEDRIVER_UI	JHgnG0YbHkld	= 0x00000029
	DIF_FINISHINSTALL_ACTION	JHgnG0YbHkld	= 0x0000002A
)


type Zyvr1EVBmP struct {
	f5ItXmR	uint32
	JlGI8r1Uz	VxzKiiU
	GT8iXm	YydljzW8Fts
	utnaNLv	uintptr
	H0maPsq5t	uintptr
	JjuGTTw	uintptr
	YjkdB5vDd	HPswIcPV29vq
	_	uintptr
	_	uint32
	sjydgjLp	[MAX_PATH]uint16
}

func (ioDqiw1Ve *Zyvr1EVBmP) DriverPath() string {
	return  /*line :1*/ OXNanQ8Uj(ioDqiw1Ve.sjydgjLp[:])
}

func (ioDqiw1Ve *Zyvr1EVBmP) SetDriverPath(dS6lrzq2Q string) error {
	xukLmcNaR, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(dS6lrzq2Q)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	 /*line :1*/ copy(ioDqiw1Ve.sjydgjLp[:], xukLmcNaR)
	return nil
}


type VxzKiiU uint32

const (
	
	DI_SHOWOEM	VxzKiiU	= 0x00000001	
	DI_SHOWCOMPAT	VxzKiiU	= 0x00000002	
	DI_SHOWCLASS	VxzKiiU	= 0x00000004	
	DI_SHOWALL	VxzKiiU	= 0x00000007	
	DI_NOVCP	VxzKiiU	= 0x00000008	
	DI_DIDCOMPAT	VxzKiiU	= 0x00000010	
	DI_DIDCLASS	VxzKiiU	= 0x00000020	
	DI_AUTOASSIGNRES	VxzKiiU	= 0x00000040	

	
	DI_NEEDRESTART	VxzKiiU	= 0x00000080	
	DI_NEEDREBOOT	VxzKiiU	= 0x00000100	

	
	DI_NOBROWSE	VxzKiiU	= 0x00000200	

	
	DI_MULTMFGS	VxzKiiU	= 0x00000400	

	
	DI_DISABLED	VxzKiiU	= 0x00000800	

	
	DI_GENERALPAGE_ADDED	VxzKiiU	= 0x00001000
	DI_RESOURCEPAGE_ADDED	VxzKiiU	= 0x00002000

	
	DI_PROPERTIES_CHANGE	VxzKiiU	= 0x00004000

	
	DI_INF_IS_SORTED	VxzKiiU	= 0x00008000

	
	DI_ENUMSINGLEINF	VxzKiiU	= 0x00010000

	
	
	DI_DONOTCALLCONFIGMG	VxzKiiU	= 0x00020000

	
	DI_INSTALLDISABLED	VxzKiiU	= 0x00040000

	
	
	DI_COMPAT_FROM_CLASS	VxzKiiU	= 0x00080000

	
	DI_CLASSINSTALLPARAMS	VxzKiiU	= 0x00100000

	
	DI_NODI_DEFAULTACTION	VxzKiiU	= 0x00200000

	
	DI_QUIETINSTALL	VxzKiiU	= 0x00800000	
	DI_NOFILECOPY	VxzKiiU	= 0x01000000	
	DI_FORCECOPY	VxzKiiU	= 0x02000000	
	DI_DRIVERPAGE_ADDED	VxzKiiU	= 0x04000000	
	DI_USECI_SELECTSTRINGS	VxzKiiU	= 0x08000000	
	DI_OVERRIDE_INFFLAGS	VxzKiiU	= 0x10000000	
	DI_PROPS_NOCHANGEUSAGE	VxzKiiU	= 0x20000000	

	DI_NOSELECTICONS	VxzKiiU	= 0x40000000	

	DI_NOWRITE_IDS	VxzKiiU	= 0x80000000	
)


type YydljzW8Fts uint32

const (
	DI_FLAGSEX_CI_FAILED	YydljzW8Fts	= 0x00000004	
	DI_FLAGSEX_FINISHINSTALL_ACTION	YydljzW8Fts	= 0x00000008	
	DI_FLAGSEX_DIDINFOLIST	YydljzW8Fts	= 0x00000010	
	DI_FLAGSEX_DIDCOMPATINFO	YydljzW8Fts	= 0x00000020	
	DI_FLAGSEX_FILTERCLASSES	YydljzW8Fts	= 0x00000040
	DI_FLAGSEX_SETFAILEDINSTALL	YydljzW8Fts	= 0x00000080
	DI_FLAGSEX_DEVICECHANGE	YydljzW8Fts	= 0x00000100
	DI_FLAGSEX_ALWAYSWRITEIDS	YydljzW8Fts	= 0x00000200
	DI_FLAGSEX_PROPCHANGE_PENDING	YydljzW8Fts	= 0x00000400	
	DI_FLAGSEX_ALLOWEXCLUDEDDRVS	YydljzW8Fts	= 0x00000800
	DI_FLAGSEX_NOUIONQUERYREMOVE	YydljzW8Fts	= 0x00001000
	DI_FLAGSEX_USECLASSFORCOMPAT	YydljzW8Fts	= 0x00002000	
	DI_FLAGSEX_NO_DRVREG_MODIFY	YydljzW8Fts	= 0x00008000	
	DI_FLAGSEX_IN_SYSTEM_SETUP	YydljzW8Fts	= 0x00010000	
	DI_FLAGSEX_INET_DRIVER	YydljzW8Fts	= 0x00020000	
	DI_FLAGSEX_APPENDDRIVERLIST	YydljzW8Fts	= 0x00040000	
	DI_FLAGSEX_PREINSTALLBACKUP	YydljzW8Fts	= 0x00080000	
	DI_FLAGSEX_BACKUPONREPLACE	YydljzW8Fts	= 0x00100000	
	DI_FLAGSEX_DRIVERLIST_FROM_URL	YydljzW8Fts	= 0x00200000	
	DI_FLAGSEX_EXCLUDE_OLD_INET_DRIVERS	YydljzW8Fts	= 0x00800000	
	DI_FLAGSEX_POWERPAGE_ADDED	YydljzW8Fts	= 0x01000000	
	DI_FLAGSEX_FILTERSIMILARDRIVERS	YydljzW8Fts	= 0x02000000	
	DI_FLAGSEX_INSTALLEDDRIVER	YydljzW8Fts	= 0x04000000	
	DI_FLAGSEX_NO_CLASSLIST_NODE_MERGE	YydljzW8Fts	= 0x08000000	
	DI_FLAGSEX_ALTPLATFORM_DRVSEARCH	YydljzW8Fts	= 0x10000000	
	DI_FLAGSEX_RESTART_DEVICE_ONLY	YydljzW8Fts	= 0x20000000	
	DI_FLAGSEX_RECURSIVESEARCH	YydljzW8Fts	= 0x40000000	
	DI_FLAGSEX_SEARCH_PUBLISHED_INFS	YydljzW8Fts	= 0x80000000	
)


type ZRao_T7a struct {
	sSIBHSz0	uint32
	RLBkg3HrTV	JHgnG0YbHkld
}

func WlsvngWX(gHa6OFe JHgnG0YbHkld) *ZRao_T7a {
	yK9ePaFKPPC := &ZRao_T7a{RLBkg3HrTV: gHa6OFe}
	yK9ePaFKPPC.sSIBHSz0 =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*yK9ePaFKPPC))
	return yK9ePaFKPPC
}


type JrZqqhj uint32

const (
	DICS_ENABLE	JrZqqhj	= 0x00000001	
	DICS_DISABLE	JrZqqhj	= 0x00000002	
	DICS_PROPCHANGE	JrZqqhj	= 0x00000003	
	DICS_START	JrZqqhj	= 0x00000004	
	DICS_STOP	JrZqqhj	= 0x00000005	
)


type ZvvCHRdtq uint32

const (
	DICS_FLAG_GLOBAL	ZvvCHRdtq	= 0x00000001	
	DICS_FLAG_CONFIGSPECIFIC	ZvvCHRdtq	= 0x00000002	
	DICS_FLAG_CONFIGGENERAL	ZvvCHRdtq	= 0x00000004	
)


type EkcE_BE5 struct {
	PhwRpGYOaoF	ZRao_T7a
	D7siyhg	JrZqqhj
	I7Xi13	ZvvCHRdtq
	KeknOu5EUB	uint32
}


type DJSgJzPZq8 uint32

const (
	DI_REMOVEDEVICE_GLOBAL	DJSgJzPZq8	= 0x00000001	
	DI_REMOVEDEVICE_CONFIGSPECIFIC	DJSgJzPZq8	= 0x00000002	
)


type II7iwgGqtI struct {
	UsLViTzseR	ZRao_T7a
	G8Ye5nSqx5c	DJSgJzPZq8
	PzFAfB_3G	uint32
}


type Sn8DCTALzB_k struct {
	dVMCjra	uint32
	OMK4RS	uint32
	_	uintptr
	fzRjqUYl	[LINE_LEN]uint16
	uiJT34C3p	[LINE_LEN]uint16
	hC8NC1GhZRWo	[LINE_LEN]uint16
	T9VhJimWr	ZPa9KL2
	K58A1m	uint64
}

func (iIzhNC *Sn8DCTALzB_k) Description() string {
	return  /*line :1*/ OXNanQ8Uj(iIzhNC.fzRjqUYl[:])
}

func (iIzhNC *Sn8DCTALzB_k) SetDescription(dD7OQR9p1FB string) error {
	xukLmcNaR, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(dD7OQR9p1FB)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	 /*line :1*/ copy(iIzhNC.fzRjqUYl[:], xukLmcNaR)
	return nil
}

func (iIzhNC *Sn8DCTALzB_k) MfgName() string {
	return  /*line :1*/ OXNanQ8Uj(iIzhNC.uiJT34C3p[:])
}

func (iIzhNC *Sn8DCTALzB_k) SetMfgName(ce72DkL6Jh string) error {
	xukLmcNaR, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(ce72DkL6Jh)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	 /*line :1*/ copy(iIzhNC.uiJT34C3p[:], xukLmcNaR)
	return nil
}

func (iIzhNC *Sn8DCTALzB_k) ProviderName() string {
	return  /*line :1*/ OXNanQ8Uj(iIzhNC.hC8NC1GhZRWo[:])
}

func (iIzhNC *Sn8DCTALzB_k) SetProviderName(h6FXpkX string) error {
	xukLmcNaR, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(h6FXpkX)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	 /*line :1*/ copy(iIzhNC.hC8NC1GhZRWo[:], xukLmcNaR)
	return nil
}


func (iIzhNC *Sn8DCTALzB_k) IsNewer(aIcxNykTcQ ZPa9KL2, p1jfuKXSvkn uint64) bool {
	if iIzhNC.T9VhJimWr.L9KWrjMHS > aIcxNykTcQ.L9KWrjMHS {
		return true
	}
	if iIzhNC.T9VhJimWr.L9KWrjMHS < aIcxNykTcQ.L9KWrjMHS {
		return false
	}

	if iIzhNC.T9VhJimWr.Cl_x1Rd > aIcxNykTcQ.Cl_x1Rd {
		return true
	}
	if iIzhNC.T9VhJimWr.Cl_x1Rd < aIcxNykTcQ.Cl_x1Rd {
		return false
	}

	if iIzhNC.K58A1m > p1jfuKXSvkn {
		return true
	}
	if iIzhNC.K58A1m < p1jfuKXSvkn {
		return false
	}

	return false
}


type PrnhDsQ struct {
	uiE4nF	uint32	
	M5ZG7uk	ZPa9KL2
	ib22TT7_8R4	uint32
	pHNLjyHWXDs	uint32
	_	uintptr
	cDRAwUxXg	[LINE_LEN]uint16
	qF691gPPnH	[MAX_PATH]uint16
	hXTwu2	[LINE_LEN]uint16
	x8O_Mg	[1]uint16
}

func (*PrnhDsQ) rm3s25Ys() uint32 {
	if  /*line :1*/ unsafe.Sizeof( /*line :1*/ uintptr(0)) == 4 {

		return  /*line :1*/ uint32( /*line :1*/ unsafe.Offsetof(PrnhDsQ{}.x8O_Mg) +  /*line :1*/ unsafe.Sizeof(PrnhDsQ{}.x8O_Mg))
	}
	return  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(PrnhDsQ{}))
}

func (iIzhNC *PrnhDsQ) SectionName() string {
	return  /*line :1*/ OXNanQ8Uj(iIzhNC.cDRAwUxXg[:])
}

func (iIzhNC *PrnhDsQ) InfFileName() string {
	return  /*line :1*/ OXNanQ8Uj(iIzhNC.qF691gPPnH[:])
}

func (iIzhNC *PrnhDsQ) DrvDescription() string {
	return  /*line :1*/ OXNanQ8Uj(iIzhNC.hXTwu2[:])
}

func (iIzhNC *PrnhDsQ) HardwareID() string {
	if iIzhNC.ib22TT7_8R4 > 1 {
		vBbtDTt1 :=  /*line :1*/ iIzhNC.rbYp2yuVT()
		return  /*line :1*/ OXNanQ8Uj(vBbtDTt1[: /*line :1*/ g7L3d8(vBbtDTt1)])
	}

	return ""
}

func (iIzhNC *PrnhDsQ) CompatIDs() []string {
	xINVctm :=  /*line :1*/ make([]string, 0)

	if iIzhNC.pHNLjyHWXDs > 0 {
		vBbtDTt1 :=  /*line :1*/ iIzhNC.rbYp2yuVT()
		vBbtDTt1 = vBbtDTt1[iIzhNC.ib22TT7_8R4 : iIzhNC.ib22TT7_8R4+iIzhNC.pHNLjyHWXDs]
		for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(vBbtDTt1); {
			hFiUmivkwHK := rRGfxgPH8Kq +  /*line :1*/ g7L3d8(vBbtDTt1[rRGfxgPH8Kq:])
			if rRGfxgPH8Kq < hFiUmivkwHK {
				xINVctm =  /*line :1*/ append(xINVctm,  /*line :1*/ OXNanQ8Uj(vBbtDTt1[rRGfxgPH8Kq:hFiUmivkwHK]))
			}
			rRGfxgPH8Kq = hFiUmivkwHK + 1
		}
	}

	return xINVctm
}

func (iIzhNC *PrnhDsQ) rbYp2yuVT() []uint16 {
	ajJiI4KvLI := (iIzhNC.uiE4nF -  /*line :1*/ uint32( /*line :1*/ unsafe.Offsetof(iIzhNC.x8O_Mg))) / 2
	aqgNlsBU3y := struct {
		tcFoIn_m5	*uint16
		oc60ywdf	int
		cVukrgwmfzae	int
	}{&iIzhNC.x8O_Mg[0],  /*line :1*/ int(ajJiI4KvLI),  /*line :1*/ int(ajJiI4KvLI)}
	return *(*[] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&aqgNlsBU3y))
}


func (iIzhNC *PrnhDsQ) IsCompatible(xkVEtL1pN string) bool {
	bGXqke3xoxJh :=  /*line :1*/ strings.WKrzOP(xkVEtL1pN)
	if  /*line :1*/ strings.WKrzOP( /*line :1*/ iIzhNC.HardwareID()) == bGXqke3xoxJh {
		return true
	}
	xINVctm :=  /*line :1*/ iIzhNC.CompatIDs()
	for rRGfxgPH8Kq := range xINVctm {
		if  /*line :1*/ strings.WKrzOP(xINVctm[rRGfxgPH8Kq]) == bGXqke3xoxJh {
			return true
		}
	}

	return false
}


type GTWay7x uint32

const (
	DICD_GENERATE_ID	GTWay7x	= 0x00000001
	DICD_INHERIT_CLASSDRVS	GTWay7x	= 0x00000002
)


type Ipdrvc9b9wmZ uint32

const (
	SUOI_FORCEDELETE Ipdrvc9b9wmZ = 0x0001
)




type MsQkOr uint32

const (
	SPDIT_NODRIVER	MsQkOr	= 0x00000000
	SPDIT_CLASSDRIVER	MsQkOr	= 0x00000001
	SPDIT_COMPATDRIVER	MsQkOr	= 0x00000002
)


type Llqbd1Tk uint32

const (
	DIGCF_DEFAULT	Llqbd1Tk	= 0x00000001	
	DIGCF_PRESENT	Llqbd1Tk	= 0x00000002
	DIGCF_ALLCLASSES	Llqbd1Tk	= 0x00000004
	DIGCF_PROFILE	Llqbd1Tk	= 0x00000008
	DIGCF_DEVICEINTERFACE	Llqbd1Tk	= 0x00000010
)


type RChEbtEIYbR uint32

const (
	DIREG_DEV	RChEbtEIYbR	= 0x00000001	
	DIREG_DRV	RChEbtEIYbR	= 0x00000002	
	DIREG_BOTH	RChEbtEIYbR	= 0x00000004	
)









type Y3qEJ_aVo uint32

const (
	SPDRP_DEVICEDESC	Y3qEJ_aVo	= 0x00000000	
	SPDRP_HARDWAREID	Y3qEJ_aVo	= 0x00000001	
	SPDRP_COMPATIBLEIDS	Y3qEJ_aVo	= 0x00000002	
	SPDRP_SERVICE	Y3qEJ_aVo	= 0x00000004	
	SPDRP_CLASS	Y3qEJ_aVo	= 0x00000007	
	SPDRP_CLASSGUID	Y3qEJ_aVo	= 0x00000008	
	SPDRP_DRIVER	Y3qEJ_aVo	= 0x00000009	
	SPDRP_CONFIGFLAGS	Y3qEJ_aVo	= 0x0000000A	
	SPDRP_MFG	Y3qEJ_aVo	= 0x0000000B	
	SPDRP_FRIENDLYNAME	Y3qEJ_aVo	= 0x0000000C	
	SPDRP_LOCATION_INFORMATION	Y3qEJ_aVo	= 0x0000000D	
	SPDRP_PHYSICAL_DEVICE_OBJECT_NAME	Y3qEJ_aVo	= 0x0000000E	
	SPDRP_CAPABILITIES	Y3qEJ_aVo	= 0x0000000F	
	SPDRP_UI_NUMBER	Y3qEJ_aVo	= 0x00000010	
	SPDRP_UPPERFILTERS	Y3qEJ_aVo	= 0x00000011	
	SPDRP_LOWERFILTERS	Y3qEJ_aVo	= 0x00000012	
	SPDRP_BUSTYPEGUID	Y3qEJ_aVo	= 0x00000013	
	SPDRP_LEGACYBUSTYPE	Y3qEJ_aVo	= 0x00000014	
	SPDRP_BUSNUMBER	Y3qEJ_aVo	= 0x00000015	
	SPDRP_ENUMERATOR_NAME	Y3qEJ_aVo	= 0x00000016	
	SPDRP_SECURITY	Y3qEJ_aVo	= 0x00000017	
	SPDRP_SECURITY_SDS	Y3qEJ_aVo	= 0x00000018	
	SPDRP_DEVTYPE	Y3qEJ_aVo	= 0x00000019	
	SPDRP_EXCLUSIVE	Y3qEJ_aVo	= 0x0000001A	
	SPDRP_CHARACTERISTICS	Y3qEJ_aVo	= 0x0000001B	
	SPDRP_ADDRESS	Y3qEJ_aVo	= 0x0000001C	
	SPDRP_UI_NUMBER_DESC_FORMAT	Y3qEJ_aVo	= 0x0000001D	
	SPDRP_DEVICE_POWER_DATA	Y3qEJ_aVo	= 0x0000001E	
	SPDRP_REMOVAL_POLICY	Y3qEJ_aVo	= 0x0000001F	
	SPDRP_REMOVAL_POLICY_HW_DEFAULT	Y3qEJ_aVo	= 0x00000020	
	SPDRP_REMOVAL_POLICY_OVERRIDE	Y3qEJ_aVo	= 0x00000021	
	SPDRP_INSTALL_STATE	Y3qEJ_aVo	= 0x00000022	
	SPDRP_LOCATION_PATHS	Y3qEJ_aVo	= 0x00000023	
	SPDRP_BASE_CONTAINERID	Y3qEJ_aVo	= 0x00000024	

	SPDRP_MAXIMUM_PROPERTY	Y3qEJ_aVo	= 0x00000025	
)



type F1HyK5M uint32

const (
	DEVPROP_TYPEMOD_ARRAY	F1HyK5M	= 0x00001000
	DEVPROP_TYPEMOD_LIST	F1HyK5M	= 0x00002000

	DEVPROP_TYPE_EMPTY	F1HyK5M	= 0x00000000
	DEVPROP_TYPE_NULL	F1HyK5M	= 0x00000001
	DEVPROP_TYPE_SBYTE	F1HyK5M	= 0x00000002
	DEVPROP_TYPE_BYTE	F1HyK5M	= 0x00000003
	DEVPROP_TYPE_INT16	F1HyK5M	= 0x00000004
	DEVPROP_TYPE_UINT16	F1HyK5M	= 0x00000005
	DEVPROP_TYPE_INT32	F1HyK5M	= 0x00000006
	DEVPROP_TYPE_UINT32	F1HyK5M	= 0x00000007
	DEVPROP_TYPE_INT64	F1HyK5M	= 0x00000008
	DEVPROP_TYPE_UINT64	F1HyK5M	= 0x00000009
	DEVPROP_TYPE_FLOAT	F1HyK5M	= 0x0000000A
	DEVPROP_TYPE_DOUBLE	F1HyK5M	= 0x0000000B
	DEVPROP_TYPE_DECIMAL	F1HyK5M	= 0x0000000C
	DEVPROP_TYPE_GUID	F1HyK5M	= 0x0000000D
	DEVPROP_TYPE_CURRENCY	F1HyK5M	= 0x0000000E
	DEVPROP_TYPE_DATE	F1HyK5M	= 0x0000000F
	DEVPROP_TYPE_FILETIME	F1HyK5M	= 0x00000010
	DEVPROP_TYPE_BOOLEAN	F1HyK5M	= 0x00000011
	DEVPROP_TYPE_STRING	F1HyK5M	= 0x00000012
	DEVPROP_TYPE_STRING_LIST	F1HyK5M	= DEVPROP_TYPE_STRING | DEVPROP_TYPEMOD_LIST
	DEVPROP_TYPE_SECURITY_DESCRIPTOR	F1HyK5M	= 0x00000013
	DEVPROP_TYPE_SECURITY_DESCRIPTOR_STRING	F1HyK5M	= 0x00000014
	DEVPROP_TYPE_DEVPROPKEY	F1HyK5M	= 0x00000015
	DEVPROP_TYPE_DEVPROPTYPE	F1HyK5M	= 0x00000016
	DEVPROP_TYPE_BINARY	F1HyK5M	= DEVPROP_TYPE_BYTE | DEVPROP_TYPEMOD_ARRAY
	DEVPROP_TYPE_ERROR	F1HyK5M	= 0x00000017
	DEVPROP_TYPE_NTSTATUS	F1HyK5M	= 0x00000018
	DEVPROP_TYPE_STRING_INDIRECT	F1HyK5M	= 0x00000019

	MAX_DEVPROP_TYPE	F1HyK5M	= 0x00000019
	MAX_DEVPROP_TYPEMOD	F1HyK5M	= 0x00002000

	DEVPROP_MASK_TYPE	F1HyK5M	= 0x00000FFF
	DEVPROP_MASK_TYPEMOD	F1HyK5M	= 0x0000F000
)


type DC5hiG Kw2qafRFaiLg


type X2tVazappD uint32

const DEVPROPID_FIRST_USABLE X2tVazappD = 2



type XqoC_LNSBj struct {
	DqcIfa	DC5hiG
	HNc0ACX2eGP	X2tVazappD
}


type KvSvJgB uint32

func (aNqiLgOJxfqs KvSvJgB) Error() string {
	if _wvkvQ2bQN_R, emJsu1awF2t :=  /*line :1*/ aNqiLgOJxfqs.Unwrap().(W5PDg27zDC); emJsu1awF2t {
		return  /*line :1*/ fmt.IBsS3Oc("%s (CfgMgr error: 0x%08x)",  /*line :1*/ _wvkvQ2bQN_R.Error(),  /*line :1*/ uint32(aNqiLgOJxfqs))
	}
	return  /*line :1*/ fmt.IBsS3Oc("CfgMgr error: 0x%08x",  /*line :1*/ uint32(aNqiLgOJxfqs))
}

func (aNqiLgOJxfqs KvSvJgB) Win32Error(hV8w_M3F1m W5PDg27zDC) W5PDg27zDC {
	return  /*line :1*/ uCiG3eCXXs(aNqiLgOJxfqs, hV8w_M3F1m)
}

func (aNqiLgOJxfqs KvSvJgB) Unwrap() error {
	const noMatch =  /*line :1*/ W5PDg27zDC(^ /*line :1*/ uintptr(0))
	_wvkvQ2bQN_R :=  /*line :1*/ aNqiLgOJxfqs.Win32Error(noMatch)
	if _wvkvQ2bQN_R == noMatch {
		return nil
	}
	return _wvkvQ2bQN_R
}

const (
	CR_SUCCESS	KvSvJgB	= 0x00000000
	CR_DEFAULT	KvSvJgB	= 0x00000001
	CR_OUT_OF_MEMORY	KvSvJgB	= 0x00000002
	CR_INVALID_POINTER	KvSvJgB	= 0x00000003
	CR_INVALID_FLAG	KvSvJgB	= 0x00000004
	CR_INVALID_DEVNODE	KvSvJgB	= 0x00000005
	CR_INVALID_DEVINST		= CR_INVALID_DEVNODE
	CR_INVALID_RES_DES	KvSvJgB	= 0x00000006
	CR_INVALID_LOG_CONF	KvSvJgB	= 0x00000007
	CR_INVALID_ARBITRATOR	KvSvJgB	= 0x00000008
	CR_INVALID_NODELIST	KvSvJgB	= 0x00000009
	CR_DEVNODE_HAS_REQS	KvSvJgB	= 0x0000000A
	CR_DEVINST_HAS_REQS		= CR_DEVNODE_HAS_REQS
	CR_INVALID_RESOURCEID	KvSvJgB	= 0x0000000B
	CR_DLVXD_NOT_FOUND	KvSvJgB	= 0x0000000C
	CR_NO_SUCH_DEVNODE	KvSvJgB	= 0x0000000D
	CR_NO_SUCH_DEVINST		= CR_NO_SUCH_DEVNODE
	CR_NO_MORE_LOG_CONF	KvSvJgB	= 0x0000000E
	CR_NO_MORE_RES_DES	KvSvJgB	= 0x0000000F
	CR_ALREADY_SUCH_DEVNODE	KvSvJgB	= 0x00000010
	CR_ALREADY_SUCH_DEVINST		= CR_ALREADY_SUCH_DEVNODE
	CR_INVALID_RANGE_LIST	KvSvJgB	= 0x00000011
	CR_INVALID_RANGE	KvSvJgB	= 0x00000012
	CR_FAILURE	KvSvJgB	= 0x00000013
	CR_NO_SUCH_LOGICAL_DEV	KvSvJgB	= 0x00000014
	CR_CREATE_BLOCKED	KvSvJgB	= 0x00000015
	CR_NOT_SYSTEM_VM	KvSvJgB	= 0x00000016
	CR_REMOVE_VETOED	KvSvJgB	= 0x00000017
	CR_APM_VETOED	KvSvJgB	= 0x00000018
	CR_INVALID_LOAD_TYPE	KvSvJgB	= 0x00000019
	CR_BUFFER_SMALL	KvSvJgB	= 0x0000001A
	CR_NO_ARBITRATOR	KvSvJgB	= 0x0000001B
	CR_NO_REGISTRY_HANDLE	KvSvJgB	= 0x0000001C
	CR_REGISTRY_ERROR	KvSvJgB	= 0x0000001D
	CR_INVALID_DEVICE_ID	KvSvJgB	= 0x0000001E
	CR_INVALID_DATA	KvSvJgB	= 0x0000001F
	CR_INVALID_API	KvSvJgB	= 0x00000020
	CR_DEVLOADER_NOT_READY	KvSvJgB	= 0x00000021
	CR_NEED_RESTART	KvSvJgB	= 0x00000022
	CR_NO_MORE_HW_PROFILES	KvSvJgB	= 0x00000023
	CR_DEVICE_NOT_THERE	KvSvJgB	= 0x00000024
	CR_NO_SUCH_VALUE	KvSvJgB	= 0x00000025
	CR_WRONG_TYPE	KvSvJgB	= 0x00000026
	CR_INVALID_PRIORITY	KvSvJgB	= 0x00000027
	CR_NOT_DISABLEABLE	KvSvJgB	= 0x00000028
	CR_FREE_RESOURCES	KvSvJgB	= 0x00000029
	CR_QUERY_VETOED	KvSvJgB	= 0x0000002A
	CR_CANT_SHARE_IRQ	KvSvJgB	= 0x0000002B
	CR_NO_DEPENDENT	KvSvJgB	= 0x0000002C
	CR_SAME_RESOURCES	KvSvJgB	= 0x0000002D
	CR_NO_SUCH_REGISTRY_KEY	KvSvJgB	= 0x0000002E
	CR_INVALID_MACHINENAME	KvSvJgB	= 0x0000002F
	CR_REMOTE_COMM_FAILURE	KvSvJgB	= 0x00000030
	CR_MACHINE_UNAVAILABLE	KvSvJgB	= 0x00000031
	CR_NO_CM_SERVICES	KvSvJgB	= 0x00000032
	CR_ACCESS_DENIED	KvSvJgB	= 0x00000033
	CR_CALL_NOT_IMPLEMENTED	KvSvJgB	= 0x00000034
	CR_INVALID_PROPERTY	KvSvJgB	= 0x00000035
	CR_DEVICE_INTERFACE_ACTIVE	KvSvJgB	= 0x00000036
	CR_NO_SUCH_DEVICE_INTERFACE	KvSvJgB	= 0x00000037
	CR_INVALID_REFERENCE_STRING	KvSvJgB	= 0x00000038
	CR_INVALID_CONFLICT_LIST	KvSvJgB	= 0x00000039
	CR_INVALID_INDEX	KvSvJgB	= 0x0000003A
	CR_INVALID_STRUCTURE_SIZE	KvSvJgB	= 0x0000003B
	NUM_CR_RESULTS	KvSvJgB	= 0x0000003C
)

const (
	CM_GET_DEVICE_INTERFACE_LIST_PRESENT	= 0		
	CM_GET_DEVICE_INTERFACE_LIST_ALL_DEVICES	= 1		
)

const (
	DN_ROOT_ENUMERATED	= 0x00000001		
	DN_DRIVER_LOADED	= 0x00000002		
	DN_ENUM_LOADED	= 0x00000004		
	DN_STARTED	= 0x00000008		
	DN_MANUAL	= 0x00000010		
	DN_NEED_TO_ENUM	= 0x00000020		
	DN_NOT_FIRST_TIME	= 0x00000040		
	DN_HARDWARE_ENUM	= 0x00000080		
	DN_LIAR	= 0x00000100		
	DN_HAS_MARK	= 0x00000200		
	DN_HAS_PROBLEM	= 0x00000400		
	DN_FILTERED	= 0x00000800		
	DN_MOVED	= 0x00001000		
	DN_DISABLEABLE	= 0x00002000		
	DN_REMOVABLE	= 0x00004000		
	DN_PRIVATE_PROBLEM	= 0x00008000		
	DN_MF_PARENT	= 0x00010000		
	DN_MF_CHILD	= 0x00020000		
	DN_WILL_BE_REMOVED	= 0x00040000		
	DN_NOT_FIRST_TIMEE	= 0x00080000		
	DN_STOP_FREE_RES	= 0x00100000		
	DN_REBAL_CANDIDATE	= 0x00200000		
	DN_BAD_PARTIAL	= 0x00400000		
	DN_NT_ENUMERATOR	= 0x00800000		
	DN_NT_DRIVER	= 0x01000000		
	DN_NEEDS_LOCKING	= 0x02000000		
	DN_ARM_WAKEUP	= 0x04000000		
	DN_APM_ENUMERATOR	= 0x08000000		
	DN_APM_DRIVER	= 0x10000000		
	DN_SILENT_INSTALL	= 0x20000000		
	DN_NO_SHOW_IN_DM	= 0x40000000		
	DN_BOOT_LOG_PROB	= 0x80000000		
	DN_NEED_RESTART	= DN_LIAR		
	DN_DRIVER_BLOCKED	= DN_NOT_FIRST_TIME		
	DN_LEGACY_DRIVER	= DN_MOVED		
	DN_CHILD_WITH_INVALID_ID	= DN_HAS_MARK		
	DN_DEVICE_DISCONNECTED	= DN_NEEDS_LOCKING		
	DN_QUERY_REMOVE_PENDING	= DN_MF_PARENT		
	DN_QUERY_REMOVE_ACTIVE	= DN_MF_CHILD		
	DN_CHANGEABLE_FLAGS	= DN_NOT_FIRST_TIME | DN_HARDWARE_ENUM | DN_HAS_MARK | DN_DISABLEABLE | DN_REMOVABLE | DN_MF_CHILD | DN_MF_PARENT | DN_NOT_FIRST_TIMEE | DN_STOP_FREE_RES | DN_REBAL_CANDIDATE | DN_NT_ENUMERATOR | DN_NT_DRIVER | DN_SILENT_INSTALL | DN_NO_SHOW_IN_DM
)


func RSJP2UU7(tRrEQhp *Kw2qafRFaiLg, rV_CV4S59 uintptr, mhANSS4o3m string) (uGD5H4 WIItFoK, jeWMpOaQtWV error) {
	var dauwBIn *uint16
	if mhANSS4o3m != "" {
		dauwBIn, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(mhANSS4o3m)
		if jeWMpOaQtWV != nil {
			return
		}
	}
	return  /*line :1*/ b_Lptvocgo65(tRrEQhp, rV_CV4S59, dauwBIn, 0)
}


func Q_PRATcz0r(uGD5H4 WIItFoK) (erAGnRa4Uj *F9bfgg, jeWMpOaQtWV error) {
	iIzhNC := &F9bfgg{}
	iIzhNC.cD48bm =  /*line :1*/ iIzhNC.rm3s25Ys()

	return iIzhNC,  /*line :1*/ as3wti2MyBA(uGD5H4, iIzhNC)
}


func (uGD5H4 WIItFoK) DeviceInfoListDetail() (*F9bfgg, error) {
	return  /*line :1*/ Q_PRATcz0r(uGD5H4)
}


func YEbgChuFFgMx(uGD5H4 WIItFoK, qNhx_SU2 string, tRrEQhp *Kw2qafRFaiLg, iPHfLVoFA string, rV_CV4S59 uintptr, elKoAUa0KJ GTWay7x) (kj4GTXbbJ8XY *BKGe6zA1I1c_, jeWMpOaQtWV error) {
	c7lK0mv7k8Rc, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(qNhx_SU2)
	if jeWMpOaQtWV != nil {
		return
	}

	var ayUYua *uint16
	if iPHfLVoFA != "" {
		ayUYua, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(iPHfLVoFA)
		if jeWMpOaQtWV != nil {
			return
		}
	}

	iIzhNC := &BKGe6zA1I1c_{}
	iIzhNC.ftvJsHy_ex1F =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*iIzhNC))

	return iIzhNC,  /*line :1*/ wqgGQbt3gOk(uGD5H4, c7lK0mv7k8Rc, tRrEQhp, ayUYua, rV_CV4S59, elKoAUa0KJ, iIzhNC)
}


func (uGD5H4 WIItFoK) CreateDeviceInfo(qNhx_SU2 string, tRrEQhp *Kw2qafRFaiLg, iPHfLVoFA string, rV_CV4S59 uintptr, elKoAUa0KJ GTWay7x) (*BKGe6zA1I1c_, error) {
	return  /*line :1*/ YEbgChuFFgMx(uGD5H4, qNhx_SU2, tRrEQhp, iPHfLVoFA, rV_CV4S59, elKoAUa0KJ)
}


func Vldi90icYabi(uGD5H4 WIItFoK, o6uepAP3rUW int) (*BKGe6zA1I1c_, error) {
	iIzhNC := &BKGe6zA1I1c_{}
	iIzhNC.ftvJsHy_ex1F =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*iIzhNC))

	return iIzhNC,  /*line :1*/ le4H0tRTkLCV(uGD5H4,  /*line :1*/ uint32(o6uepAP3rUW), iIzhNC)
}


func (uGD5H4 WIItFoK) EnumDeviceInfo(o6uepAP3rUW int) (*BKGe6zA1I1c_, error) {
	return  /*line :1*/ Vldi90icYabi(uGD5H4, o6uepAP3rUW)
}


func (uGD5H4 WIItFoK) Close() error {
	return  /*line :1*/ CHqytC(uGD5H4)
}


func (uGD5H4 WIItFoK) BuildDriverInfoList(kj4GTXbbJ8XY *BKGe6zA1I1c_, aEsa_kNra MsQkOr) error {
	return  /*line :1*/ Qi8B5rfAA(uGD5H4, kj4GTXbbJ8XY, aEsa_kNra)
}


func (uGD5H4 WIItFoK) CancelDriverInfoSearch() error {
	return  /*line :1*/ VWi8tQ(uGD5H4)
}


func Am8f_odiyWJP(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, aEsa_kNra MsQkOr, o6uepAP3rUW int) (*Sn8DCTALzB_k, error) {
	iIzhNC := &Sn8DCTALzB_k{}
	iIzhNC.dVMCjra =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*iIzhNC))

	return iIzhNC,  /*line :1*/ blpxoz(uGD5H4, kj4GTXbbJ8XY, aEsa_kNra,  /*line :1*/ uint32(o6uepAP3rUW), iIzhNC)
}


func (uGD5H4 WIItFoK) EnumDriverInfo(kj4GTXbbJ8XY *BKGe6zA1I1c_, aEsa_kNra MsQkOr, o6uepAP3rUW int) (*Sn8DCTALzB_k, error) {
	return  /*line :1*/ Am8f_odiyWJP(uGD5H4, kj4GTXbbJ8XY, aEsa_kNra, o6uepAP3rUW)
}


func H8z04Ax4Ht(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_) (*Sn8DCTALzB_k, error) {
	iIzhNC := &Sn8DCTALzB_k{}
	iIzhNC.dVMCjra =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*iIzhNC))

	return iIzhNC,  /*line :1*/ h6shLJ_qt9mi(uGD5H4, kj4GTXbbJ8XY, iIzhNC)
}


func (uGD5H4 WIItFoK) SelectedDriver(kj4GTXbbJ8XY *BKGe6zA1I1c_) (*Sn8DCTALzB_k, error) {
	return  /*line :1*/ H8z04Ax4Ht(uGD5H4, kj4GTXbbJ8XY)
}


func (uGD5H4 WIItFoK) SetSelectedDriver(kj4GTXbbJ8XY *BKGe6zA1I1c_, j4llyD5g *Sn8DCTALzB_k) error {
	return  /*line :1*/ FoOJWbA(uGD5H4, kj4GTXbbJ8XY, j4llyD5g)
}


func K5Nr6KwD2(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, j4llyD5g *Sn8DCTALzB_k) (*PrnhDsQ, error) {
	uGwCR86NJOCP :=  /*line :1*/ uint32(2048)
	for {
		etRYtnVni :=  /*line :1*/ make([]byte, uGwCR86NJOCP)
		iIzhNC := (* /*line :1*/ PrnhDsQ)( /*line :1*/ unsafe.Pointer(&etRYtnVni[0]))
		iIzhNC.uiE4nF =  /*line :1*/ iIzhNC.rm3s25Ys()
		jeWMpOaQtWV :=  /*line :1*/ cwEq8JK508(uGD5H4, kj4GTXbbJ8XY, j4llyD5g, iIzhNC,  /*line :1*/ uint32( /*line :1*/ len(etRYtnVni)), &uGwCR86NJOCP)
		if jeWMpOaQtWV == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if jeWMpOaQtWV != nil {
			return nil, jeWMpOaQtWV
		}
		iIzhNC.uiE4nF = uGwCR86NJOCP
		return iIzhNC, nil
	}
}


func (uGD5H4 WIItFoK) DriverInfoDetail(kj4GTXbbJ8XY *BKGe6zA1I1c_, j4llyD5g *Sn8DCTALzB_k) (*PrnhDsQ, error) {
	return  /*line :1*/ K5Nr6KwD2(uGD5H4, kj4GTXbbJ8XY, j4llyD5g)
}


func (uGD5H4 WIItFoK) DestroyDriverInfoList(kj4GTXbbJ8XY *BKGe6zA1I1c_, aEsa_kNra MsQkOr) error {
	return  /*line :1*/ GsqRaqJK(uGD5H4, kj4GTXbbJ8XY, aEsa_kNra)
}


func AUKaEZO(tRrEQhp *Kw2qafRFaiLg, vgFANxwPlotu string, rV_CV4S59 uintptr, a_MrGKcrR Llqbd1Tk, uGD5H4 WIItFoK, mhANSS4o3m string) (iOvctVD26lA WIItFoK, jeWMpOaQtWV error) {
	var fjjqZWpOEF *uint16
	if vgFANxwPlotu != "" {
		fjjqZWpOEF, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(vgFANxwPlotu)
		if jeWMpOaQtWV != nil {
			return
		}
	}
	var dauwBIn *uint16
	if mhANSS4o3m != "" {
		dauwBIn, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(mhANSS4o3m)
		if jeWMpOaQtWV != nil {
			return
		}
	}
	return  /*line :1*/ aoGAk8eiQp(tRrEQhp, fjjqZWpOEF, rV_CV4S59, a_MrGKcrR, uGD5H4, dauwBIn, 0)
}


func (uGD5H4 WIItFoK) CallClassInstaller(gHa6OFe JHgnG0YbHkld, kj4GTXbbJ8XY *BKGe6zA1I1c_) error {
	return  /*line :1*/ GkM4P09Im86(gHa6OFe, uGD5H4, kj4GTXbbJ8XY)
}


func (uGD5H4 WIItFoK) OpenDevRegKey(FThEWny94VL2 *BKGe6zA1I1c_, IXxa82ht78k ZvvCHRdtq, AyGPXZz uint32, AQ6iswf_ RChEbtEIYbR, e_w62Sbj uint32) (L2L8P9WaNZ, error) {
	return  /*line :1*/ U9OPHqeld8A(uGD5H4, FThEWny94VL2, IXxa82ht78k, AyGPXZz, AQ6iswf_, e_w62Sbj)
}


func HcUzoWJ(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, aDANLjxdYzx *XqoC_LNSBj) (wvZBcfh interface{}, jeWMpOaQtWV error) {
	uGwCR86NJOCP :=  /*line :1*/ uint32(256)
	for {
		var a0Am7ZL F1HyK5M
		etRYtnVni :=  /*line :1*/ make([]byte, uGwCR86NJOCP)
		jeWMpOaQtWV =  /*line :1*/ vygzmRw(uGD5H4, kj4GTXbbJ8XY, aDANLjxdYzx, &a0Am7ZL, &etRYtnVni[0],  /*line :1*/ uint32( /*line :1*/ len(etRYtnVni)), &uGwCR86NJOCP, 0)
		if jeWMpOaQtWV == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if jeWMpOaQtWV != nil {
			return
		}
		switch a0Am7ZL {
		case DEVPROP_TYPE_STRING:
			aNqiLgOJxfqs :=  /*line :1*/ OXNanQ8Uj( /*line :1*/ fNHo3hEnhJms(etRYtnVni))
			 /*line :1*/ runtime.KeepAlive(etRYtnVni)
			return aNqiLgOJxfqs, nil
		}
		return nil,  /*line :1*/ errors.Su6g6hRoi9X("unimplemented property type")
	}
}


func Lf83hGPirR(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, fBn5qKfnS7 Y3qEJ_aVo) (wvZBcfh interface{}, jeWMpOaQtWV error) {
	uGwCR86NJOCP :=  /*line :1*/ uint32(256)
	for {
		var a0Am7ZL uint32
		etRYtnVni :=  /*line :1*/ make([]byte, uGwCR86NJOCP)
		jeWMpOaQtWV =  /*line :1*/ tq3t9Ehr(uGD5H4, kj4GTXbbJ8XY, fBn5qKfnS7, &a0Am7ZL, &etRYtnVni[0],  /*line :1*/ uint32( /*line :1*/ len(etRYtnVni)), &uGwCR86NJOCP)
		if jeWMpOaQtWV == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if jeWMpOaQtWV != nil {
			return
		}
		return  /*line :1*/ fehU9O96TQp(etRYtnVni[:uGwCR86NJOCP], a0Am7ZL)
	}
}

func fehU9O96TQp(etRYtnVni []byte, a0Am7ZL uint32) (interface{}, error) {
	switch a0Am7ZL {
	case REG_SZ:
		aNqiLgOJxfqs :=  /*line :1*/ OXNanQ8Uj( /*line :1*/ fNHo3hEnhJms(etRYtnVni))
		 /*line :1*/ runtime.KeepAlive(etRYtnVni)
		return aNqiLgOJxfqs, nil
	case REG_EXPAND_SZ:
		wvZBcfh :=  /*line :1*/ OXNanQ8Uj( /*line :1*/ fNHo3hEnhJms(etRYtnVni))
		if wvZBcfh == "" {
			return "", nil
		}
		hD4wNgEB, jeWMpOaQtWV :=  /*line :1*/ syscall.GcOmFfsibES(wvZBcfh)
		if jeWMpOaQtWV != nil {
			return "", jeWMpOaQtWV
		}
		aNqiLgOJxfqs :=  /*line :1*/ make([]uint16, 100)
		for {
			krzuku, jeWMpOaQtWV :=  /*line :1*/ TuqYU1ro(hD4wNgEB, &aNqiLgOJxfqs[0],  /*line :1*/ uint32( /*line :1*/ len(aNqiLgOJxfqs)))
			if jeWMpOaQtWV != nil {
				return "", jeWMpOaQtWV
			}
			if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(aNqiLgOJxfqs)) {
				return  /*line :1*/ OXNanQ8Uj(aNqiLgOJxfqs[:krzuku]), nil
			}
			aNqiLgOJxfqs =  /*line :1*/ make([]uint16, krzuku)
		}
	case REG_BINARY:
		return etRYtnVni, nil
	case REG_DWORD_LITTLE_ENDIAN:
		return  /*line :1*/ binary.BE16BND.Uint32(etRYtnVni), nil
	case REG_DWORD_BIG_ENDIAN:
		return  /*line :1*/ binary.FOcKBq6.Uint32(etRYtnVni), nil
	case REG_MULTI_SZ:
		vBbtDTt1 :=  /*line :1*/ fNHo3hEnhJms(etRYtnVni)
		xINVctm := []string{}
		for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(vBbtDTt1); {
			hFiUmivkwHK := rRGfxgPH8Kq +  /*line :1*/ g7L3d8(vBbtDTt1[rRGfxgPH8Kq:])
			if rRGfxgPH8Kq < hFiUmivkwHK {
				xINVctm =  /*line :1*/ append(xINVctm,  /*line :1*/ OXNanQ8Uj(vBbtDTt1[rRGfxgPH8Kq:hFiUmivkwHK]))
			}
			rRGfxgPH8Kq = hFiUmivkwHK + 1
		}
		 /*line :1*/ runtime.KeepAlive(etRYtnVni)
		return xINVctm, nil
	case REG_QWORD_LITTLE_ENDIAN:
		return  /*line :1*/ binary.BE16BND.Uint64(etRYtnVni), nil
	default:
		return nil,  /*line :1*/ fmt.Cnz_ab1BaZh("Unsupported registry value type: %v", a0Am7ZL)
	}
}


func fNHo3hEnhJms(etRYtnVni []byte) []uint16 {
	aqgNlsBU3y := struct {
		tcFoIn_m5	*uint16
		oc60ywdf	int
		cVukrgwmfzae	int
	}{(* /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&etRYtnVni[0])),  /*line :1*/ len(etRYtnVni) / 2,  /*line :1*/ cap(etRYtnVni) / 2}
	return *(*[] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&aqgNlsBU3y))
}


func hqohTw(etRYtnVni []uint16) []byte {
	aqgNlsBU3y := struct {
		ok7yquRZjsc	*byte
		wCFF95JG	int
		oLJ8aSWYwov	int
	}{(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&etRYtnVni[0])),  /*line :1*/ len(etRYtnVni) * 2,  /*line :1*/ cap(etRYtnVni) * 2}
	return *(*[] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&aqgNlsBU3y))
}

func g7L3d8(xukLmcNaR []uint16) int {
	for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(xukLmcNaR); rRGfxgPH8Kq++ {
		if xukLmcNaR[rRGfxgPH8Kq] == 0 {
			return rRGfxgPH8Kq
		}
	}
	return  /*line :1*/ len(xukLmcNaR)
}


func (uGD5H4 WIItFoK) DeviceRegistryProperty(kj4GTXbbJ8XY *BKGe6zA1I1c_, fBn5qKfnS7 Y3qEJ_aVo) (interface{}, error) {
	return  /*line :1*/ Lf83hGPirR(uGD5H4, kj4GTXbbJ8XY, fBn5qKfnS7)
}


func Q1fp7Zpm3j(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_, fBn5qKfnS7 Y3qEJ_aVo, y11NDPKXxX []byte) error {
	return  /*line :1*/ vanFmZJ(uGD5H4, kj4GTXbbJ8XY, fBn5qKfnS7, &y11NDPKXxX[0],  /*line :1*/ uint32( /*line :1*/ len(y11NDPKXxX)))
}


func (uGD5H4 WIItFoK) SetDeviceRegistryProperty(kj4GTXbbJ8XY *BKGe6zA1I1c_, fBn5qKfnS7 Y3qEJ_aVo, y11NDPKXxX []byte) error {
	return  /*line :1*/ Q1fp7Zpm3j(uGD5H4, kj4GTXbbJ8XY, fBn5qKfnS7, y11NDPKXxX)
}


func (uGD5H4 WIItFoK) SetDeviceRegistryPropertyString(kj4GTXbbJ8XY *BKGe6zA1I1c_, fBn5qKfnS7 Y3qEJ_aVo, xukLmcNaR string) error {
	jyB65N, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(xukLmcNaR)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	jeWMpOaQtWV =  /*line :1*/ Q1fp7Zpm3j(uGD5H4, kj4GTXbbJ8XY, fBn5qKfnS7,  /*line :1*/ hqohTw( /*line :1*/ append(jyB65N, 0)))
	 /*line :1*/ runtime.KeepAlive(jyB65N)
	return jeWMpOaQtWV
}


func IHLjTALqVR(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_) (*Zyvr1EVBmP, error) {
	ioDqiw1Ve := &Zyvr1EVBmP{}
	ioDqiw1Ve.f5ItXmR =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*ioDqiw1Ve))

	return ioDqiw1Ve,  /*line :1*/ poTBP8wYR(uGD5H4, kj4GTXbbJ8XY, ioDqiw1Ve)
}


func (uGD5H4 WIItFoK) DeviceInstallParams(kj4GTXbbJ8XY *BKGe6zA1I1c_) (*Zyvr1EVBmP, error) {
	return  /*line :1*/ IHLjTALqVR(uGD5H4, kj4GTXbbJ8XY)
}


func F3fCcb(uGD5H4 WIItFoK, kj4GTXbbJ8XY *BKGe6zA1I1c_) (string, error) {
	uGwCR86NJOCP :=  /*line :1*/ uint32(1024)
	for {
		etRYtnVni :=  /*line :1*/ make([]uint16, uGwCR86NJOCP)
		jeWMpOaQtWV :=  /*line :1*/ x9MM7Y(uGD5H4, kj4GTXbbJ8XY, &etRYtnVni[0],  /*line :1*/ uint32( /*line :1*/ len(etRYtnVni)), &uGwCR86NJOCP)
		if jeWMpOaQtWV == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if jeWMpOaQtWV != nil {
			return "", jeWMpOaQtWV
		}
		return  /*line :1*/ OXNanQ8Uj(etRYtnVni), nil
	}
}


func (uGD5H4 WIItFoK) DeviceInstanceID(kj4GTXbbJ8XY *BKGe6zA1I1c_) (string, error) {
	return  /*line :1*/ F3fCcb(uGD5H4, kj4GTXbbJ8XY)
}


func (uGD5H4 WIItFoK) ClassInstallParams(kj4GTXbbJ8XY *BKGe6zA1I1c_, ffpjJE16Qt *ZRao_T7a, kl2waNwDJNO3 uint32, ljAFfCOKOT *uint32) error {
	return  /*line :1*/ HMqZZCgR5W(uGD5H4, kj4GTXbbJ8XY, ffpjJE16Qt, kl2waNwDJNO3, ljAFfCOKOT)
}


func (uGD5H4 WIItFoK) SetDeviceInstallParams(kj4GTXbbJ8XY *BKGe6zA1I1c_, gBpptk_r *Zyvr1EVBmP) error {
	return  /*line :1*/ Yk1KOGBG_heq(uGD5H4, kj4GTXbbJ8XY, gBpptk_r)
}


func (uGD5H4 WIItFoK) SetClassInstallParams(kj4GTXbbJ8XY *BKGe6zA1I1c_, ffpjJE16Qt *ZRao_T7a, kl2waNwDJNO3 uint32) error {
	return  /*line :1*/ Uc8omY_CN(uGD5H4, kj4GTXbbJ8XY, ffpjJE16Qt, kl2waNwDJNO3)
}


func Gg5S01J4(tRrEQhp *Kw2qafRFaiLg, mhANSS4o3m string) (dW9gI7 string, jeWMpOaQtWV error) {
	var habNUI [MAX_CLASS_NAME_LEN]uint16

	var dauwBIn *uint16
	if mhANSS4o3m != "" {
		dauwBIn, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(mhANSS4o3m)
		if jeWMpOaQtWV != nil {
			return
		}
	}

	jeWMpOaQtWV =  /*line :1*/ ic0Xno(tRrEQhp, &habNUI[0], MAX_CLASS_NAME_LEN, nil, dauwBIn, 0)
	if jeWMpOaQtWV != nil {
		return
	}

	dW9gI7 =  /*line :1*/ OXNanQ8Uj(habNUI[:])
	return
}


func QNRSz5NFr(dW9gI7 string, mhANSS4o3m string) ([]Kw2qafRFaiLg, error) {
	habNUI, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(dW9gI7)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}

	var dauwBIn *uint16
	if mhANSS4o3m != "" {
		dauwBIn, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(mhANSS4o3m)
		if jeWMpOaQtWV != nil {
			return nil, jeWMpOaQtWV
		}
	}

	uGwCR86NJOCP :=  /*line :1*/ uint32(4)
	for {
		etRYtnVni :=  /*line :1*/ make([]Kw2qafRFaiLg, uGwCR86NJOCP)
		jeWMpOaQtWV =  /*line :1*/ aga7luXkmw(habNUI, &etRYtnVni[0],  /*line :1*/ uint32( /*line :1*/ len(etRYtnVni)), &uGwCR86NJOCP, dauwBIn, 0)
		if jeWMpOaQtWV == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if jeWMpOaQtWV != nil {
			return nil, jeWMpOaQtWV
		}
		return etRYtnVni[:uGwCR86NJOCP], nil
	}
}


func Ul28ueqRsjP(uGD5H4 WIItFoK) (*BKGe6zA1I1c_, error) {
	iIzhNC := &BKGe6zA1I1c_{}
	iIzhNC.ftvJsHy_ex1F =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*iIzhNC))

	return iIzhNC,  /*line :1*/ fhWjAaGiUV(uGD5H4, iIzhNC)
}


func (uGD5H4 WIItFoK) SelectedDevice() (*BKGe6zA1I1c_, error) {
	return  /*line :1*/ Ul28ueqRsjP(uGD5H4)
}


func (uGD5H4 WIItFoK) SetSelectedDevice(kj4GTXbbJ8XY *BKGe6zA1I1c_) error {
	return  /*line :1*/ UnukQswxpS(uGD5H4, kj4GTXbbJ8XY)
}


func Dgv81PI7mD1(bcTTXDUmWK string, a_MrGKcrR Ipdrvc9b9wmZ) error {
	bYsi4U9, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(bcTTXDUmWK)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ gIY4mxFW(bYsi4U9, a_MrGKcrR, 0)
}

func WajdRI(fSGXH2Ei string, zflUYaIh *Kw2qafRFaiLg, a_MrGKcrR uint32) ([]string, error) {
	xoy1tYEa, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(fSGXH2Ei)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	var etRYtnVni []uint16
	var cuGEvMbsvH uint32
	for {
		if aNqiLgOJxfqs :=  /*line :1*/ jA0Hqvk(&cuGEvMbsvH, zflUYaIh, xoy1tYEa, a_MrGKcrR); aNqiLgOJxfqs != CR_SUCCESS {
			return nil, aNqiLgOJxfqs
		}
		etRYtnVni =  /*line :1*/ make([]uint16, cuGEvMbsvH)
		if aNqiLgOJxfqs :=  /*line :1*/ flQ5hc68(zflUYaIh, xoy1tYEa, &etRYtnVni[0], cuGEvMbsvH, a_MrGKcrR); aNqiLgOJxfqs == CR_SUCCESS {
			break
		} else if aNqiLgOJxfqs != CR_BUFFER_SMALL {
			return nil, aNqiLgOJxfqs
		}
	}
	var joaUjjrJi []string
	for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(etRYtnVni); {
		hFiUmivkwHK := rRGfxgPH8Kq +  /*line :1*/ g7L3d8(etRYtnVni[rRGfxgPH8Kq:])
		if rRGfxgPH8Kq < hFiUmivkwHK {
			joaUjjrJi =  /*line :1*/ append(joaUjjrJi,  /*line :1*/ OXNanQ8Uj(etRYtnVni[rRGfxgPH8Kq:hFiUmivkwHK]))
		}
		rRGfxgPH8Kq = hFiUmivkwHK + 1
	}
	if joaUjjrJi == nil {
		return nil, ERROR_NO_SUCH_DEVICE_INTERFACE
	}
	return joaUjjrJi, nil
}

func Z9Bg7463djGk(oGbyaObSkJ *uint32, yQbI4wDaBXm *uint32, qGOky8aUf2 Dn3mHAPt5D, a_MrGKcrR uint32) error {
	aNqiLgOJxfqs :=  /*line :1*/ sDQ5uFx(oGbyaObSkJ, yQbI4wDaBXm, qGOky8aUf2, a_MrGKcrR)
	if aNqiLgOJxfqs == CR_SUCCESS {
		return nil
	}
	return aNqiLgOJxfqs
}
