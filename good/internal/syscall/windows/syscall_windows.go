//line :1
package LdptURlN

import (
	sync "sync"
	syscall "bUKeamF"
	"unsafe"
)



func DOR2gxA_7npQ(cWT9SOJ *uint16) string {
	if cWT9SOJ == nil {
		return ""
	}
	qU5mla :=  /*line :1*/ unsafe.Pointer(cWT9SOJ)
	pdOCpIN7 := 0
	for *(* /*line :1*/ uint16)(qU5mla) != 0 {
		qU5mla =  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(qU5mla) +  /*line :1*/ unsafe.Sizeof(*cWT9SOJ))
		pdOCpIN7++
	}
	return  /*line :1*/ syscall.AODVXp8NM3sd( /*line :1*/ unsafe.Slice(cWT9SOJ, pdOCpIN7))
}

const (
	ERROR_BAD_LENGTH	syscall.O9Mga3	= 24
	ERROR_SHARING_VIOLATION	syscall.O9Mga3	= 32
	ERROR_LOCK_VIOLATION	syscall.O9Mga3	= 33
	ERROR_NOT_SUPPORTED	syscall.O9Mga3	= 50
	ERROR_CALL_NOT_IMPLEMENTED	syscall.O9Mga3	= 120
	ERROR_INVALID_NAME	syscall.O9Mga3	= 123
	ERROR_LOCK_FAILED	syscall.O9Mga3	= 167
	ERROR_NO_UNICODE_TRANSLATION	syscall.O9Mga3	= 1113
)

const GAA_FLAG_INCLUDE_PREFIX = 0x00000010

const (
	IF_TYPE_OTHER	= 1
	IF_TYPE_ETHERNET_CSMACD	= 6
	IF_TYPE_ISO88025_TOKENRING	= 9
	IF_TYPE_PPP	= 23
	IF_TYPE_SOFTWARE_LOOPBACK	= 24
	IF_TYPE_ATM	= 37
	IF_TYPE_IEEE80211	= 71
	IF_TYPE_TUNNEL	= 131
	IF_TYPE_IEEE1394	= 144
)

type GfD2t_nRYEy struct {
	NgzCg7T0_8yT	*syscall.IXy5oynSaLQM
	RCOHDt0lwCJ	int32
}

type IKO_CJ3aCP5 struct {
	Ed5v3FKh	uint32
	LKrjkdB	uint32
	QL7SVZyWpMDz	*IKO_CJ3aCP5
	Zwf7iZEa	GfD2t_nRYEy
	AHCee8dD4C	int32
	E1nwWqczVlCp	int32
	RkJ37f9P	int32
	S4YyU7	uint32
	P9d4qce	uint32
	Uo21ziwIa	uint32
	X2zs0IMo	uint8
}

type BCdWngJnjOrm struct {
	ShMsEIKf6nD	uint32
	CSKi51p	uint32
	Xp8705	*BCdWngJnjOrm
	ZTtP0dAzzo6F	GfD2t_nRYEy
}

type SJJmPV9n struct {
	NDcV1vmzZqN	uint32
	QgC64zE	uint32
	Ltx_FiDtXD	*SJJmPV9n
	F07g7TUjQ	GfD2t_nRYEy
}

type INpFJz70pVE9 struct {
	T43VEcr8aja	uint32
	RnLcnpJ	uint32
	HyuwAR7G	*INpFJz70pVE9
	JuLy4dFjAt	GfD2t_nRYEy
}

type AhmS71ue5F struct {
	DFcaErc	uint32
	HT_cABaEudev	uint32
	JP35xcgviD	*AhmS71ue5F
	ABJ5Nu7oIWI	GfD2t_nRYEy
	CrYjdktxSD	uint32
}

type WXavkAObzxI struct {
	ClZuZUDI	uint32
	ZvNqpW	uint32
	PS1xLQA	*WXavkAObzxI
	NJVQHb	*byte
	N5iUveDkX	*IKO_CJ3aCP5
	BaxbOi	*BCdWngJnjOrm
	BX8u3C	*SJJmPV9n
	ALEYAV_	*INpFJz70pVE9
	VphVd7ckCrkn	*uint16
	Eo3VFM2I6OY	*uint16
	MndZKwmw0g_C	*uint16
	XP4Wryx7b2h	[syscall.MAX_ADAPTER_ADDRESS_LENGTH]byte
	DKieIVi0	uint32
	NInmGka	uint32
	YLKBBLbazwH	uint32
	IHFX_Ws	uint32
	AOg4mu	uint32
	KLtlYXkLZOhm	uint32
	RmTh0emxnp	[16]uint32
	GKaWn0BK7t8	*AhmS71ue5F
}

type BR3dBXMrt1du struct {
	HRBTc9nIta	uint16
	QSRxah	uintptr
	V6aVQp5u7oc	bool
}

type ORAO0wrXtu struct {
	HdaP0YD9y	syscall.T8WbUqAC3v
	BWWdWT8L2YP	syscall.T8WbUqAC3v
	F7HtasGUuru	syscall.T8WbUqAC3v
	Fq7XjdGvRTK	syscall.T8WbUqAC3v
	FuIvuad	uint32
}

const (
	IfOperStatusUp	= 1
	IfOperStatusDown	= 2
	IfOperStatusTesting	= 3
	IfOperStatusUnknown	= 4
	IfOperStatusDormant	= 5
	IfOperStatusNotPresent	= 6
	IfOperStatusLowerLayerDown	= 7
)

const (
	
	TH32CS_SNAPMODULE	= 0x08
	TH32CS_SNAPMODULE32	= 0x10
)

const MAX_MODULE_NAME32 = 255

type LX2aqkMG_K struct {
	InzCLmLxQ	uint32
	PMHqOelZ0YmS	uint32
	XTf2P4y6nEAX	uint32
	AOOzfBaaS4	uint32
	WObg05c6	uint32
	MkbeXdrD5N	uintptr
	YC9rHP	uint32
	X7Q6fO7E431	syscall.SRlvVjrYa
	Aa65Nj1S9	[MAX_MODULE_NAME32 + 1]uint16
	GRTgLESI	[syscall.MAX_PATH]uint16
}

const SizeofModuleEntry32 =  /*line :1*/ unsafe.Sizeof(LX2aqkMG_K{})

const (
	WSA_FLAG_OVERLAPPED		= 0x01
	WSA_FLAG_NO_HANDLE_INHERIT		= 0x80

	WSAEMSGSIZE	syscall.O9Mga3	= 10040

	MSG_PEEK		= 0x2
	MSG_TRUNC		= 0x0100
	MSG_CTRUNC		= 0x0200

	socket_error		=  /*line :1*/ uintptr(^ /*line :1*/ uint32(0))
)

var HZVqi18Jkpr = syscall.RlB0Ha5cnNd{
	EpYSt8TqI:	0xa441e712,
	UHbfXLLNS6:	0x754f,
	GAYsFb:	0x43ca,
	Z9vXwokyJ:	[8]byte{0x84, 0xa7, 0x0d, 0xee, 0x44, 0xcf, 0x60, 0x6d},
}

var Ro5hRfhGcDnF = syscall.RlB0Ha5cnNd{
	EpYSt8TqI:	0xf689d7c8,
	UHbfXLLNS6:	0x6f1f,
	GAYsFb:	0x436b,
	Z9vXwokyJ:	[8]byte{0x8a, 0x53, 0xe5, 0x4f, 0xe3, 0x51, 0xc3, 0x22},
}

var pMfNFA4vgz struct {
	dyil5im	sync.LhBfwn6wa1x
	vJwTbp7_Z	uintptr
	pFC6w1	uintptr
	qmqIZ_3	error
}

type BXk_cV9l_bLI struct {
	ClAqGya	syscall.VjfACnL
	MeKqL5	int32
	MD9Fwj	*syscall.QbWhAp5CqTG
	Du6szpnO	uint32
	RjCEgxXvL4QC	syscall.QbWhAp5CqTG
	WDy42Vy	uint32
}

func jnCuSaD() error {
	 /*line :1*/ pMfNFA4vgz.dyil5im.Do(func() {
		var gp3a5i1F syscall.SRlvVjrYa
		gp3a5i1F, pMfNFA4vgz.qmqIZ_3 =  /*line :1*/ syscall.RQpzEduJMY(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
		if pMfNFA4vgz.qmqIZ_3 != nil {
			return
		}
		defer  /*line :1*/ syscall.FhKJLgXjwG(gp3a5i1F)
		var pdOCpIN7 uint32
		pMfNFA4vgz.qmqIZ_3 =  /*line :1*/ syscall.YvCdY02(gp3a5i1F,
			syscall.SIO_GET_EXTENSION_FUNCTION_POINTER,
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&Ro5hRfhGcDnF)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(Ro5hRfhGcDnF)),
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&pMfNFA4vgz.pFC6w1)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(pMfNFA4vgz.pFC6w1)),
			&pdOCpIN7, nil, 0)
		if pMfNFA4vgz.qmqIZ_3 != nil {
			return
		}
		pMfNFA4vgz.qmqIZ_3 =  /*line :1*/ syscall.YvCdY02(gp3a5i1F,
			syscall.SIO_GET_EXTENSION_FUNCTION_POINTER,
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&HZVqi18Jkpr)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(HZVqi18Jkpr)),
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&pMfNFA4vgz.vJwTbp7_Z)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(pMfNFA4vgz.vJwTbp7_Z)),
			&pdOCpIN7, nil, 0)
	})
	return pMfNFA4vgz.qmqIZ_3
}

func UsGtgcD5(cyWOFRvqnVD syscall.SRlvVjrYa, qNLkqx9 *BXk_cV9l_bLI, iK5SVBoZG uint32, diaQ4ef *uint32, bS4DAcuB0sz *syscall.ZaNm5QSYC9, bT89xwB *byte) error {
	zc4PmxS :=  /*line :1*/ jnCuSaD()
	if zc4PmxS != nil {
		return zc4PmxS
	}
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft(pMfNFA4vgz.vJwTbp7_Z, 6,  /*line :1*/ uintptr(cyWOFRvqnVD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qNLkqx9)),  /*line :1*/ uintptr(iK5SVBoZG),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(diaQ4ef)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bS4DAcuB0sz)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bT89xwB)))
	if cMhXXOuwZkJB == socket_error {
		if tgROg0 != 0 {
			zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
		} else {
			zc4PmxS = syscall.EINVAL
		}
	}
	return zc4PmxS
}

func GhYfJv(cyWOFRvqnVD syscall.SRlvVjrYa, qNLkqx9 *BXk_cV9l_bLI, eqI7LnvloWDb *uint32, bS4DAcuB0sz *syscall.ZaNm5QSYC9, bT89xwB *byte) error {
	zc4PmxS :=  /*line :1*/ jnCuSaD()
	if zc4PmxS != nil {
		return zc4PmxS
	}
	cMhXXOuwZkJB, _, tgROg0 :=  /*line :1*/ syscall.VeAKF8sAwft(pMfNFA4vgz.pFC6w1, 5,  /*line :1*/ uintptr(cyWOFRvqnVD),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qNLkqx9)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(eqI7LnvloWDb)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bS4DAcuB0sz)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bT89xwB)), 0)
	if cMhXXOuwZkJB == socket_error {
		if tgROg0 != 0 {
			zc4PmxS =  /*line :1*/ pBI0QG(tgROg0)
		} else {
			zc4PmxS = syscall.EINVAL
		}
	}
	return zc4PmxS
}

const (
	ComputerNameNetBIOS	= 0
	ComputerNameDnsHostname	= 1
	ComputerNameDnsDomain	= 2
	ComputerNameDnsFullyQualified	= 3
	ComputerNamePhysicalNetBIOS	= 4
	ComputerNamePhysicalDnsHostname	= 5
	ComputerNamePhysicalDnsDomain	= 6
	ComputerNamePhysicalDnsFullyQualified	= 7
	ComputerNameMax	= 8

	MOVEFILE_REPLACE_EXISTING	= 0x1
	MOVEFILE_COPY_ALLOWED	= 0x2
	MOVEFILE_DELAY_UNTIL_REBOOT	= 0x4
	MOVEFILE_WRITE_THROUGH	= 0x8
	MOVEFILE_CREATE_HARDLINK	= 0x10
	MOVEFILE_FAIL_IF_NOT_TRACKABLE	= 0x20
)

func FJvJDUwWaoc(jdaFyv5, i7ZLo3PC string) error {
	y6SejQ5f, zc4PmxS :=  /*line :1*/ syscall.GcOmFfsibES(jdaFyv5)
	if zc4PmxS != nil {
		return zc4PmxS
	}
	uxhDNcje, zc4PmxS :=  /*line :1*/ syscall.GcOmFfsibES(i7ZLo3PC)
	if zc4PmxS != nil {
		return zc4PmxS
	}
	return  /*line :1*/ S6mU1WXHl(y6SejQ5f, uxhDNcje, MOVEFILE_REPLACE_EXISTING)
}

const (
	LOCKFILE_FAIL_IMMEDIATELY	= 0x00000001
	LOCKFILE_EXCLUSIVE_LOCK	= 0x00000002
)

const MB_ERR_INVALID_CHARS = 8

const STYPE_DISKTREE = 0x00

type HIoGS5P struct {
	FsKKfLEm	*uint16
	IfjHRwVuffd	uint32
	Oa1UvMBJlt5i	*uint16
	WHvIV112YpWJ	uint32
	YMZzaI9	uint32
	IRE1yaTwsxVQ	uint32
	Dyjnt8Y_2	*uint16
	VMNOGLkWnals	*uint16
}

const (
	FILE_NAME_NORMALIZED	= 0x0
	FILE_NAME_OPENED	= 0x8

	VOLUME_NAME_DOS	= 0x0
	VOLUME_NAME_GUID	= 0x1
	VOLUME_NAME_NONE	= 0x4
	VOLUME_NAME_NT	= 0x2
)

func EOgS9BWI() error {
	return  /*line :1*/ kJ570MGiPPuX.Find()
}

func BVEjpkkAWtv() error {
	return  /*line :1*/ bGXtbzu.Find()
}
