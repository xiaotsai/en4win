//line :1
package bUKeamF

const (
	
	ERROR_FILE_NOT_FOUND	O9Mga3	= 2
	ERROR_PATH_NOT_FOUND	O9Mga3	= 3
	ERROR_ACCESS_DENIED	O9Mga3	= 5
	ERROR_NO_MORE_FILES	O9Mga3	= 18
	ERROR_HANDLE_EOF	O9Mga3	= 38
	ERROR_NETNAME_DELETED	O9Mga3	= 64
	ERROR_FILE_EXISTS	O9Mga3	= 80
	ERROR_BROKEN_PIPE	O9Mga3	= 109
	ERROR_BUFFER_OVERFLOW	O9Mga3	= 111
	ERROR_INSUFFICIENT_BUFFER	O9Mga3	= 122
	ERROR_MOD_NOT_FOUND	O9Mga3	= 126
	ERROR_PROC_NOT_FOUND	O9Mga3	= 127
	ERROR_DIR_NOT_EMPTY	O9Mga3	= 145
	ERROR_ALREADY_EXISTS	O9Mga3	= 183
	ERROR_ENVVAR_NOT_FOUND	O9Mga3	= 203
	ERROR_MORE_DATA	O9Mga3	= 234
	ERROR_OPERATION_ABORTED	O9Mga3	= 995
	ERROR_IO_PENDING	O9Mga3	= 997
	ERROR_NOT_FOUND	O9Mga3	= 1168
	ERROR_PRIVILEGE_NOT_HELD	O9Mga3	= 1314
	WSAEACCES	O9Mga3	= 10013
	WSAECONNABORTED	O9Mga3	= 10053
	WSAECONNRESET	O9Mga3	= 10054
)

const (
	
	O_RDONLY	= 0x00000
	O_WRONLY	= 0x00001
	O_RDWR	= 0x00002
	O_CREAT	= 0x00040
	O_EXCL	= 0x00080
	O_NOCTTY	= 0x00100
	O_TRUNC	= 0x00200
	O_NONBLOCK	= 0x00800
	O_APPEND	= 0x00400
	O_SYNC	= 0x01000
	O_ASYNC	= 0x02000
	O_CLOEXEC	= 0x80000
)

const (
	
	SIGHUP	=  /*line :1*/ EslOtdor(0x1)
	SIGINT	=  /*line :1*/ EslOtdor(0x2)
	SIGQUIT	=  /*line :1*/ EslOtdor(0x3)
	SIGILL	=  /*line :1*/ EslOtdor(0x4)
	SIGTRAP	=  /*line :1*/ EslOtdor(0x5)
	SIGABRT	=  /*line :1*/ EslOtdor(0x6)
	SIGBUS	=  /*line :1*/ EslOtdor(0x7)
	SIGFPE	=  /*line :1*/ EslOtdor(0x8)
	SIGKILL	=  /*line :1*/ EslOtdor(0x9)
	SIGSEGV	=  /*line :1*/ EslOtdor(0xb)
	SIGPIPE	=  /*line :1*/ EslOtdor(0xd)
	SIGALRM	=  /*line :1*/ EslOtdor(0xe)
	SIGTERM	=  /*line :1*/ EslOtdor(0xf)
)

var wDO4xIZNiGl = [...]string{
	1:	"hangup",
	2:	"interrupt",
	3:	"quit",
	4:	"illegal instruction",
	5:	"trace/breakpoint trap",
	6:	"aborted",
	7:	"bus error",
	8:	"floating point exception",
	9:	"killed",
	10:	"user defined signal 1",
	11:	"segmentation fault",
	12:	"user defined signal 2",
	13:	"broken pipe",
	14:	"alarm clock",
	15:	"terminated",
}

const (
	GENERIC_READ	= 0x80000000
	GENERIC_WRITE	= 0x40000000
	GENERIC_EXECUTE	= 0x20000000
	GENERIC_ALL	= 0x10000000

	FILE_LIST_DIRECTORY	= 0x00000001
	FILE_APPEND_DATA	= 0x00000004
	FILE_WRITE_ATTRIBUTES	= 0x00000100

	FILE_SHARE_READ	= 0x00000001
	FILE_SHARE_WRITE	= 0x00000002
	FILE_SHARE_DELETE	= 0x00000004
	FILE_ATTRIBUTE_READONLY	= 0x00000001
	FILE_ATTRIBUTE_HIDDEN	= 0x00000002
	FILE_ATTRIBUTE_SYSTEM	= 0x00000004
	FILE_ATTRIBUTE_DIRECTORY	= 0x00000010
	FILE_ATTRIBUTE_ARCHIVE	= 0x00000020
	FILE_ATTRIBUTE_NORMAL	= 0x00000080
	FILE_ATTRIBUTE_REPARSE_POINT	= 0x00000400

	INVALID_FILE_ATTRIBUTES	= 0xffffffff

	CREATE_NEW	= 1
	CREATE_ALWAYS	= 2
	OPEN_EXISTING	= 3
	OPEN_ALWAYS	= 4
	TRUNCATE_EXISTING	= 5

	FILE_FLAG_OPEN_REPARSE_POINT	= 0x00200000
	FILE_FLAG_BACKUP_SEMANTICS	= 0x02000000
	FILE_FLAG_OVERLAPPED	= 0x40000000

	HANDLE_FLAG_INHERIT	= 0x00000001
	STARTF_USESTDHANDLES	= 0x00000100
	STARTF_USESHOWWINDOW	= 0x00000001
	DUPLICATE_CLOSE_SOURCE	= 0x00000001
	DUPLICATE_SAME_ACCESS	= 0x00000002

	STD_INPUT_HANDLE	= -10
	STD_OUTPUT_HANDLE	= -11
	STD_ERROR_HANDLE	= -12

	FILE_BEGIN	= 0
	FILE_CURRENT	= 1
	FILE_END	= 2

	LANG_ENGLISH	= 0x09
	SUBLANG_ENGLISH_US	= 0x01

	FORMAT_MESSAGE_ALLOCATE_BUFFER	= 256
	FORMAT_MESSAGE_IGNORE_INSERTS	= 512
	FORMAT_MESSAGE_FROM_STRING	= 1024
	FORMAT_MESSAGE_FROM_HMODULE	= 2048
	FORMAT_MESSAGE_FROM_SYSTEM	= 4096
	FORMAT_MESSAGE_ARGUMENT_ARRAY	= 8192
	FORMAT_MESSAGE_MAX_WIDTH_MASK	= 255

	MAX_PATH	= 260
	MAX_LONG_PATH	= 32768

	MAX_COMPUTERNAME_LENGTH	= 15

	TIME_ZONE_ID_UNKNOWN	= 0
	TIME_ZONE_ID_STANDARD	= 1

	TIME_ZONE_ID_DAYLIGHT	= 2
	IGNORE	= 0
	INFINITE	= 0xffffffff

	WAIT_TIMEOUT	= 258
	WAIT_ABANDONED	= 0x00000080
	WAIT_OBJECT_0	= 0x00000000
	WAIT_FAILED	= 0xFFFFFFFF

	CREATE_NEW_PROCESS_GROUP	= 0x00000200
	CREATE_UNICODE_ENVIRONMENT	= 0x00000400

	PROCESS_TERMINATE	= 1
	PROCESS_QUERY_INFORMATION	= 0x00000400
	SYNCHRONIZE	= 0x00100000

	PAGE_READONLY	= 0x02
	PAGE_READWRITE	= 0x04
	PAGE_WRITECOPY	= 0x08
	PAGE_EXECUTE_READ	= 0x20
	PAGE_EXECUTE_READWRITE	= 0x40
	PAGE_EXECUTE_WRITECOPY	= 0x80

	FILE_MAP_COPY	= 0x01
	FILE_MAP_WRITE	= 0x02
	FILE_MAP_READ	= 0x04
	FILE_MAP_EXECUTE	= 0x20

	CTRL_C_EVENT	= 0
	CTRL_BREAK_EVENT	= 1
	CTRL_CLOSE_EVENT	= 2
	CTRL_LOGOFF_EVENT	= 5
	CTRL_SHUTDOWN_EVENT	= 6
)

const (
	
	TH32CS_SNAPHEAPLIST	= 0x01
	TH32CS_SNAPPROCESS	= 0x02
	TH32CS_SNAPTHREAD	= 0x04
	TH32CS_SNAPMODULE	= 0x08
	TH32CS_SNAPMODULE32	= 0x10
	TH32CS_SNAPALL	= TH32CS_SNAPHEAPLIST | TH32CS_SNAPMODULE | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD
	TH32CS_INHERIT	= 0x80000000
)

const (
	
	FILE_NOTIFY_CHANGE_FILE_NAME	= 1 << iota
	FILE_NOTIFY_CHANGE_DIR_NAME
	FILE_NOTIFY_CHANGE_ATTRIBUTES
	FILE_NOTIFY_CHANGE_SIZE
	FILE_NOTIFY_CHANGE_LAST_WRITE
	FILE_NOTIFY_CHANGE_LAST_ACCESS
	FILE_NOTIFY_CHANGE_CREATION
)

const (
	
	FILE_ACTION_ADDED	= iota + 1
	FILE_ACTION_REMOVED
	FILE_ACTION_MODIFIED
	FILE_ACTION_RENAMED_OLD_NAME
	FILE_ACTION_RENAMED_NEW_NAME
)

const (
	
	PROV_RSA_FULL	= 1
	PROV_RSA_SIG	= 2
	PROV_DSS	= 3
	PROV_FORTEZZA	= 4
	PROV_MS_EXCHANGE	= 5
	PROV_SSL	= 6
	PROV_RSA_SCHANNEL	= 12
	PROV_DSS_DH	= 13
	PROV_EC_ECDSA_SIG	= 14
	PROV_EC_ECNRA_SIG	= 15
	PROV_EC_ECDSA_FULL	= 16
	PROV_EC_ECNRA_FULL	= 17
	PROV_DH_SCHANNEL	= 18
	PROV_SPYRUS_LYNKS	= 20
	PROV_RNG	= 21
	PROV_INTEL_SEC	= 22
	PROV_REPLACE_OWF	= 23
	PROV_RSA_AES	= 24
	CRYPT_VERIFYCONTEXT	= 0xF0000000
	CRYPT_NEWKEYSET	= 0x00000008
	CRYPT_DELETEKEYSET	= 0x00000010
	CRYPT_MACHINE_KEYSET	= 0x00000020
	CRYPT_SILENT	= 0x00000040
	CRYPT_DEFAULT_CONTAINER_OPTIONAL	= 0x00000080

	USAGE_MATCH_TYPE_AND	= 0
	USAGE_MATCH_TYPE_OR	= 1

	X509_ASN_ENCODING	= 0x00000001
	PKCS_7_ASN_ENCODING	= 0x00010000

	CERT_STORE_PROV_MEMORY	= 2

	CERT_STORE_ADD_ALWAYS	= 4

	CERT_STORE_DEFER_CLOSE_UNTIL_LAST_FREE_FLAG	= 0x00000004

	CERT_TRUST_NO_ERROR	= 0x00000000
	CERT_TRUST_IS_NOT_TIME_VALID	= 0x00000001
	CERT_TRUST_IS_REVOKED	= 0x00000004
	CERT_TRUST_IS_NOT_SIGNATURE_VALID	= 0x00000008
	CERT_TRUST_IS_NOT_VALID_FOR_USAGE	= 0x00000010
	CERT_TRUST_IS_UNTRUSTED_ROOT	= 0x00000020
	CERT_TRUST_REVOCATION_STATUS_UNKNOWN	= 0x00000040
	CERT_TRUST_IS_CYCLIC	= 0x00000080
	CERT_TRUST_INVALID_EXTENSION	= 0x00000100
	CERT_TRUST_INVALID_POLICY_CONSTRAINTS	= 0x00000200
	CERT_TRUST_INVALID_BASIC_CONSTRAINTS	= 0x00000400
	CERT_TRUST_INVALID_NAME_CONSTRAINTS	= 0x00000800
	CERT_TRUST_HAS_NOT_SUPPORTED_NAME_CONSTRAINT	= 0x00001000
	CERT_TRUST_HAS_NOT_DEFINED_NAME_CONSTRAINT	= 0x00002000
	CERT_TRUST_HAS_NOT_PERMITTED_NAME_CONSTRAINT	= 0x00004000
	CERT_TRUST_HAS_EXCLUDED_NAME_CONSTRAINT	= 0x00008000
	CERT_TRUST_IS_OFFLINE_REVOCATION	= 0x01000000
	CERT_TRUST_NO_ISSUANCE_CHAIN_POLICY	= 0x02000000
	CERT_TRUST_IS_EXPLICIT_DISTRUST	= 0x04000000
	CERT_TRUST_HAS_NOT_SUPPORTED_CRITICAL_EXT	= 0x08000000

	CERT_CHAIN_POLICY_BASE	= 1
	CERT_CHAIN_POLICY_AUTHENTICODE	= 2
	CERT_CHAIN_POLICY_AUTHENTICODE_TS	= 3
	CERT_CHAIN_POLICY_SSL	= 4
	CERT_CHAIN_POLICY_BASIC_CONSTRAINTS	= 5
	CERT_CHAIN_POLICY_NT_AUTH	= 6
	CERT_CHAIN_POLICY_MICROSOFT_ROOT	= 7
	CERT_CHAIN_POLICY_EV	= 8

	CERT_E_EXPIRED	= 0x800B0101
	CERT_E_ROLE	= 0x800B0103
	CERT_E_PURPOSE	= 0x800B0106
	CERT_E_UNTRUSTEDROOT	= 0x800B0109
	CERT_E_CN_NO_MATCH	= 0x800B010F

	AUTHTYPE_CLIENT	= 1
	AUTHTYPE_SERVER	= 2
)

var (
	RqcdwWEw0cB	= [] /*line :1*/ byte("1.3.6.1.5.5.7.3.1\x00")
	A7XPmtnkc	= [] /*line :1*/ byte("1.3.6.1.4.1.311.10.3.3\x00")
	AMsDd5dZda5P	= [] /*line :1*/ byte("2.16.840.1.113730.4.1\x00")
)







type VjfACnL *struct{}


type VhkskrPU1 struct {
	L0uhjaO	int32
	RpHBw1QAUy	int32
}

func (v3wOtFJUw7Nu *VhkskrPU1) Nanoseconds() int64 {
	return ( /*line :1*/ int64(v3wOtFJUw7Nu.L0uhjaO)*1e6 +  /*line :1*/ int64(v3wOtFJUw7Nu.RpHBw1QAUy)) * 1e3
}

func DF1_wufO(tAjR0818K5K int64) (v3wOtFJUw7Nu VhkskrPU1) {
	v3wOtFJUw7Nu.L0uhjaO =  /*line :1*/ int32(tAjR0818K5K / 1e9)
	v3wOtFJUw7Nu.RpHBw1QAUy =  /*line :1*/ int32(tAjR0818K5K % 1e9 / 1e3)
	return
}

type DDwuM6RpYja struct {
	OIMmqEzG2Yx0	uint32
	I5UHJb	uintptr
	EDnb4yqd	uint32
}

type ZaNm5QSYC9 struct {
	AEkKaeljD	uintptr
	O8KaKrcbUen	uintptr
	BEJ63D	uint32
	BvaSyl	uint32
	Gs_TahqRod	SRlvVjrYa
}

type L0nrOsLJsV8A struct {
	ElpObS7ySxKu	uint32
	MSyxhBD	uint32
	A_ZEaTo	uint32
	YABYeCBorzY	uint16
}

type T8WbUqAC3v struct {
	Cl_x1Rd	uint32
	L9KWrjMHS	uint32
}



func (d09pCxOvo *T8WbUqAC3v) Nanoseconds() int64 {

	tAjR0818K5K :=  /*line :1*/ int64(d09pCxOvo.L9KWrjMHS)<<32 +  /*line :1*/ int64(d09pCxOvo.Cl_x1Rd)

	tAjR0818K5K -= 116444736000000000

	tAjR0818K5K *= 100
	return tAjR0818K5K
}

func Gmwg7AR8(tAjR0818K5K int64) (d09pCxOvo T8WbUqAC3v) {

	tAjR0818K5K /= 100

	tAjR0818K5K += 116444736000000000

	d09pCxOvo.Cl_x1Rd =  /*line :1*/ uint32(tAjR0818K5K & 0xffffffff)
	d09pCxOvo.L9KWrjMHS =  /*line :1*/ uint32(tAjR0818K5K >> 32 & 0xffffffff)
	return d09pCxOvo
}

type RsoZmJigCkGe struct {
	BBGnBMSj	uint32
	B0BiCI	T8WbUqAC3v
	OFY5aC	T8WbUqAC3v
	Ve2KvcJ	T8WbUqAC3v
	Yo4A8bC	uint32
	G4llND5o	uint32
	EjPPdaEO	uint32
	QSgZl1v	uint32
	LMw5oQM3	[MAX_PATH - 1]uint16
	XdgmqOE37I	[13]uint16
}



type w8vwhA struct {
	ZAV8Sqg	uint32
	VC96Sk	T8WbUqAC3v
	EVIa3aDF7	T8WbUqAC3v
	EYQ4KbotqI0_	T8WbUqAC3v
	SGeVxfg2	uint32
	EA3ag7eW	uint32
	JCg3EDf4u	uint32
	TRW1SlJwb	uint32
	Jl2arOg	[MAX_PATH]uint16
	Og98fWmH	[14]uint16
}

func m0EKKA(lftbtxK *RsoZmJigCkGe, rvDagXti2y *w8vwhA) {
	lftbtxK.BBGnBMSj = rvDagXti2y.ZAV8Sqg
	lftbtxK.B0BiCI = rvDagXti2y.VC96Sk
	lftbtxK.OFY5aC = rvDagXti2y.EVIa3aDF7
	lftbtxK.Ve2KvcJ = rvDagXti2y.EYQ4KbotqI0_
	lftbtxK.Yo4A8bC = rvDagXti2y.SGeVxfg2
	lftbtxK.G4llND5o = rvDagXti2y.EA3ag7eW
	lftbtxK.EjPPdaEO = rvDagXti2y.JCg3EDf4u
	lftbtxK.QSgZl1v = rvDagXti2y.TRW1SlJwb

	 /*line :1*/ copy(lftbtxK.LMw5oQM3[:], rvDagXti2y.Jl2arOg[:])
	 /*line :1*/ copy(lftbtxK.XdgmqOE37I[:], rvDagXti2y.Og98fWmH[:])
}

type HCYt4WU_Wmzb struct {
	FrYcZLBau7r	uint32
	JYWhEbWo_W	T8WbUqAC3v
	EN4G6pVm_kA4	T8WbUqAC3v
	YMVB_o2o	T8WbUqAC3v
	MoZwkiXS	uint32
	YX3QVvo	uint32
	WpwMOiFxZj	uint32
	At_ueR	uint32
	G219mR_c92nY	uint32
	P6HqpaH	uint32
}

const (
	GetFileExInfoStandard	= 0
	GetFileExMaxInfoLevel	= 1
)

type Sf7ZL_oAaEE struct {
	MegT05	uint32
	LaH9Sgf	T8WbUqAC3v
	NfUjFko6GhPa	T8WbUqAC3v
	QNU6t9O	T8WbUqAC3v
	XY0PVH6ujAbY	uint32
	JUVHzxEdI	uint32
}


const (
	
	SW_HIDE	= 0
	SW_NORMAL	= 1
	SW_SHOWNORMAL	= 1
	SW_SHOWMINIMIZED	= 2
	SW_SHOWMAXIMIZED	= 3
	SW_MAXIMIZE	= 3
	SW_SHOWNOACTIVATE	= 4
	SW_SHOW	= 5
	SW_MINIMIZE	= 6
	SW_SHOWMINNOACTIVE	= 7
	SW_SHOWNA	= 8
	SW_RESTORE	= 9
	SW_SHOWDEFAULT	= 10
	SW_FORCEMINIMIZE	= 11
)

type J7Nv3S0G31Q struct {
	JxcR_9A2Qcwv	uint32
	_	*uint16
	A98PGvn7iQoC	*uint16
	HL_a4AowFk_6	*uint16
	Dcksj_TNZ5b	uint32
	TFh9g3vCK	uint32
	PibQ4aZyeP	uint32
	AgY9upoDR	uint32
	ZnhOoY8	uint32
	QRiwczanqtH	uint32
	F83tFU_	uint32
	Am0GjIfa	uint32
	PnxZZH	uint16
	_	uint16
	_	*byte
	Q7xYfD5083	SRlvVjrYa
	BQKc_j	SRlvVjrYa
	Bh8xwzLH	SRlvVjrYa
}

type vYoIhduc struct {
	_ [1]byte
}

const (
	_PROC_THREAD_ATTRIBUTE_PARENT_PROCESS	= 0x00020000
	_PROC_THREAD_ATTRIBUTE_HANDLE_LIST	= 0x00020002
)

type taSzfHcIq struct {
	J7Nv3S0G31Q
	IIfiFP3pufqz	*vYoIhduc
}

const _EXTENDED_STARTUPINFO_PRESENT = 0x00080000

type NBPVXpO3Kv struct {
	V0VCmzpZzO_	SRlvVjrYa
	BRTb_8	SRlvVjrYa
	LjilbGmQX	uint32
	SeZlVJulxXFV	uint32
}

type GRU31Ht3FPn struct {
	NgWYrRzCdYqS	uint32
	A_yZAQ	uint32
	QU8DfIi	uint32
	R4dOIkdEteD	uintptr
	HOBiKb	uint32
	IXOX4q	uint32
	TvE7hf	uint32
	RXqTHud	int32
	WTphVKO	uint32
	WZnHS4	[MAX_PATH]uint16
}

type H6Qxv2jW struct {
	Xmnj_bIFJhJ	uint16
	EQp3MHA11V	uint16
	MslYdlEo	uint16
	HF4FBjj	uint16
	CySfZqD	uint16
	TFSCFS3IfPS0	uint16
	MRSnx59	uint16
	Lo99feSEOsAo	uint16
}

type MAPVxQZiq struct {
	StUAtyXP	int32
	KIcaL8MMs	[32]uint16
	NRJvDzhCOoT	H6Qxv2jW
	AW4StnUjhag8	int32
	VHZAJ_MP	[32]uint16
	WZFoaz8	H6Qxv2jW
	UAJo4xpW	int32
}

const (
	AF_UNSPEC	= 0
	AF_UNIX	= 1
	AF_INET	= 2
	AF_INET6	= 23
	AF_NETBIOS	= 17

	SOCK_STREAM	= 1
	SOCK_DGRAM	= 2
	SOCK_RAW	= 3
	SOCK_SEQPACKET	= 5

	IPPROTO_IP	= 0
	IPPROTO_IPV6	= 0x29
	IPPROTO_TCP	= 6
	IPPROTO_UDP	= 17

	SOL_SOCKET	= 0xffff
	SO_REUSEADDR	= 4
	SO_KEEPALIVE	= 8
	SO_DONTROUTE	= 16
	SO_BROADCAST	= 32
	SO_LINGER	= 128
	SO_RCVBUF	= 0x1002
	SO_SNDBUF	= 0x1001
	SO_UPDATE_ACCEPT_CONTEXT	= 0x700b
	SO_UPDATE_CONNECT_CONTEXT	= 0x7010

	IOC_OUT	= 0x40000000
	IOC_IN	= 0x80000000
	IOC_VENDOR	= 0x18000000
	IOC_INOUT	= IOC_IN | IOC_OUT
	IOC_WS2	= 0x08000000
	SIO_GET_EXTENSION_FUNCTION_POINTER	= IOC_INOUT | IOC_WS2 | 6
	SIO_KEEPALIVE_VALS	= IOC_IN | IOC_VENDOR | 4
	SIO_UDP_CONNRESET	= IOC_IN | IOC_VENDOR | 12

	IP_TOS	= 0x3
	IP_TTL	= 0x4
	IP_MULTICAST_IF	= 0x9
	IP_MULTICAST_TTL	= 0xa
	IP_MULTICAST_LOOP	= 0xb
	IP_ADD_MEMBERSHIP	= 0xc
	IP_DROP_MEMBERSHIP	= 0xd

	IPV6_V6ONLY	= 0x1b
	IPV6_UNICAST_HOPS	= 0x4
	IPV6_MULTICAST_IF	= 0x9
	IPV6_MULTICAST_HOPS	= 0xa
	IPV6_MULTICAST_LOOP	= 0xb
	IPV6_JOIN_GROUP	= 0xc
	IPV6_LEAVE_GROUP	= 0xd

	SOMAXCONN	= 0x7fffffff

	TCP_NODELAY	= 1

	SHUT_RD	= 0
	SHUT_WR	= 1
	SHUT_RDWR	= 2

	WSADESCRIPTION_LEN	= 256
	WSASYS_STATUS_LEN	= 128
)

type QbWhAp5CqTG struct {
	Ix58rFqPW	uint32
	AcsIulb2fH4	*byte
}


const (
	S_IFMT	= 0x1f000
	S_IFIFO	= 0x1000
	S_IFCHR	= 0x2000
	S_IFDIR	= 0x4000
	S_IFBLK	= 0x6000
	S_IFREG	= 0x8000
	S_IFLNK	= 0xa000
	S_IFSOCK	= 0xc000
	S_ISUID	= 0x800
	S_ISGID	= 0x400
	S_ISVTX	= 0x200
	S_IRUSR	= 0x100
	S_IWRITE	= 0x80
	S_IWUSR	= 0x80
	S_IXUSR	= 0x40
)

const (
	FILE_TYPE_CHAR	= 0x0002
	FILE_TYPE_DISK	= 0x0001
	FILE_TYPE_PIPE	= 0x0003
	FILE_TYPE_REMOTE	= 0x8000
	FILE_TYPE_UNKNOWN	= 0x0000
)

type MpOyP3z struct {
	RXoNcgaCieV	*byte
	FPUVe29FPB0k	**byte
	I8OkTv8w9X	uint16
	A73pZ6h7	uint16
	L6F6mqM	**byte
}

type DT8kuezeJ struct {
	Q60I30rnHia	*byte
	B_Khm0GjptX3	**byte
	EoKNay	uint16
}

const (
	DNS_TYPE_A	= 0x0001
	DNS_TYPE_NS	= 0x0002
	DNS_TYPE_MD	= 0x0003
	DNS_TYPE_MF	= 0x0004
	DNS_TYPE_CNAME	= 0x0005
	DNS_TYPE_SOA	= 0x0006
	DNS_TYPE_MB	= 0x0007
	DNS_TYPE_MG	= 0x0008
	DNS_TYPE_MR	= 0x0009
	DNS_TYPE_NULL	= 0x000a
	DNS_TYPE_WKS	= 0x000b
	DNS_TYPE_PTR	= 0x000c
	DNS_TYPE_HINFO	= 0x000d
	DNS_TYPE_MINFO	= 0x000e
	DNS_TYPE_MX	= 0x000f
	DNS_TYPE_TEXT	= 0x0010
	DNS_TYPE_RP	= 0x0011
	DNS_TYPE_AFSDB	= 0x0012
	DNS_TYPE_X25	= 0x0013
	DNS_TYPE_ISDN	= 0x0014
	DNS_TYPE_RT	= 0x0015
	DNS_TYPE_NSAP	= 0x0016
	DNS_TYPE_NSAPPTR	= 0x0017
	DNS_TYPE_SIG	= 0x0018
	DNS_TYPE_KEY	= 0x0019
	DNS_TYPE_PX	= 0x001a
	DNS_TYPE_GPOS	= 0x001b
	DNS_TYPE_AAAA	= 0x001c
	DNS_TYPE_LOC	= 0x001d
	DNS_TYPE_NXT	= 0x001e
	DNS_TYPE_EID	= 0x001f
	DNS_TYPE_NIMLOC	= 0x0020
	DNS_TYPE_SRV	= 0x0021
	DNS_TYPE_ATMA	= 0x0022
	DNS_TYPE_NAPTR	= 0x0023
	DNS_TYPE_KX	= 0x0024
	DNS_TYPE_CERT	= 0x0025
	DNS_TYPE_A6	= 0x0026
	DNS_TYPE_DNAME	= 0x0027
	DNS_TYPE_SINK	= 0x0028
	DNS_TYPE_OPT	= 0x0029
	DNS_TYPE_DS	= 0x002B
	DNS_TYPE_RRSIG	= 0x002E
	DNS_TYPE_NSEC	= 0x002F
	DNS_TYPE_DNSKEY	= 0x0030
	DNS_TYPE_DHCID	= 0x0031
	DNS_TYPE_UINFO	= 0x0064
	DNS_TYPE_UID	= 0x0065
	DNS_TYPE_GID	= 0x0066
	DNS_TYPE_UNSPEC	= 0x0067
	DNS_TYPE_ADDRS	= 0x00f8
	DNS_TYPE_TKEY	= 0x00f9
	DNS_TYPE_TSIG	= 0x00fa
	DNS_TYPE_IXFR	= 0x00fb
	DNS_TYPE_AXFR	= 0x00fc
	DNS_TYPE_MAILB	= 0x00fd
	DNS_TYPE_MAILA	= 0x00fe
	DNS_TYPE_ALL	= 0x00ff
	DNS_TYPE_ANY	= 0x00ff
	DNS_TYPE_WINS	= 0xff01
	DNS_TYPE_WINSR	= 0xff02
	DNS_TYPE_NBSTAT	= 0xff01
)

const (
	DNS_INFO_NO_RECORDS = 0x251D
)

const (
	
	DnsSectionQuestion	= 0x0000
	DnsSectionAnswer	= 0x0001
	DnsSectionAuthority	= 0x0002
	DnsSectionAdditional	= 0x0003
)

type NL1jD3M7P struct {
	JbYsAIP	*uint16
	FkTrEz	uint16
	INJblcAM	uint16
	NUAVLSp5MDik	uint16
	X1nSbpmhRr	uint16
}

type RILagwbjo struct {
	CVn0dr1Jky6Z *uint16
}

type AGNDZEpVSE struct {
	YNeBpnBS45M	*uint16
	ZT51ejkkMgU	uint16
	IXyKLv0avJ	uint16
}

type ZaocUo291mr struct {
	EKVQR85t9paB	uint16
	AfGeifce	[1]*uint16
}

type OB0HzAtGUg struct {
	IBwV54lKvxs	*OB0HzAtGUg
	NHneOoxbcY	*uint16
	SKG0xmj0Cq	uint16
	Rfrf3e_0jka	uint16
	Dkc4C3R3KS	uint32
	RKrhaEjc	uint32
	Fk0SW038	uint32
	TnF1oOW	[40]byte
}

const (
	TF_DISCONNECT	= 1
	TF_REUSE_SOCKET	= 2
	TF_WRITE_BEHIND	= 4
	TF_USE_DEFAULT_WORKER	= 0
	TF_USE_SYSTEM_THREAD	= 16
	TF_USE_KERNEL_APC	= 32
)

type GMHaGunA struct {
	ENmMpmvvAj	uintptr
	YT2zHB	uint32
	Y34XDc	uintptr
	ICQAUo_7Hq	uint32
}

const (
	IFF_UP	= 1
	IFF_BROADCAST	= 2
	IFF_LOOPBACK	= 4
	IFF_POINTTOPOINT	= 8
	IFF_MULTICAST	= 16
)

const SIO_GET_INTERFACE_LIST = 0x4004747F

type ELRhth [24]byte

type KPyDNqbz struct {
	IWUyVd	uint32
	GVHrNm	ELRhth
	NVpB_Ik_YI	ELRhth
	KcC_BhInX46	ELRhth
}

type Ith5dfGw struct {
	Mi8RNVUyx [16]byte
}

type EU5u6qT_W6e Ith5dfGw

type K4GMtqw9Y9RB struct {
	SmThUc4z	*K4GMtqw9Y9RB
	V7X8t7zM	Ith5dfGw
	I6PkVbnxN84	EU5u6qT_W6e
	QYCBMi	uint32
}

const MAX_ADAPTER_NAME_LENGTH = 256
const MAX_ADAPTER_DESCRIPTION_LENGTH = 128
const MAX_ADAPTER_ADDRESS_LENGTH = 8

type VNW4OcO_CoZ7 struct {
	QaLPOKl1	*VNW4OcO_CoZ7
	ISZLV_rN	uint32
	WHsK2_ORAS	[MAX_ADAPTER_NAME_LENGTH + 4]byte
	Ck0dgJgSZU7	[MAX_ADAPTER_DESCRIPTION_LENGTH + 4]byte
	S7BajQZ5xxJy	uint32
	G9xOuC3Y	[MAX_ADAPTER_ADDRESS_LENGTH]byte
	D1lQYoXTR	uint32
	IbarWxa	uint32
	BBGPViNzG4zJ	uint32
	Siej6vx	*K4GMtqw9Y9RB
	WFEeypPC	K4GMtqw9Y9RB
	W6w_Yr	K4GMtqw9Y9RB
	OaEXmJ8TLAce	K4GMtqw9Y9RB
	Io6tJXN	bool
	GsVLIz	K4GMtqw9Y9RB
	EGBUq8KlE	K4GMtqw9Y9RB
	ZZLI1oZrW	int64
	Lhfwh7OsQ	int64
}

const MAXLEN_PHYSADDR = 8
const MAX_INTERFACE_NAME_LEN = 256
const MAXLEN_IFDESCR = 256

type SBqHJrlvgiBJ struct {
	AHmBRp2	[MAX_INTERFACE_NAME_LEN]uint16
	AalUkvV	uint32
	HfwuL7Mdcv6J	uint32
	ZKkDIVP8B	uint32
	MkG4OfaoFy	uint32
	AqS2Cpa9p_	uint32
	JuhiwoH	[MAXLEN_PHYSADDR]byte
	KVUBorMAEtSr	uint32
	HhspVwoW6bas	uint32
	Pwjwgoigj2EL	uint32
	ByPSlzTZk9E	uint32
	HGgW1C6	uint32
	AQe1a9Yl	uint32
	ScOESZGY_B4	uint32
	BlnVsr	uint32
	LFAoGAGvf	uint32
	DbXmapwVa8Li	uint32
	FLPuQxBWEqBU	uint32
	IYEbTl	uint32
	Kdxwou6tPNi	uint32
	TXrDFhaF	uint32
	NsyO4m	uint32
	AcuAeVa	uint32
	EgXaGNmHjEp	[MAXLEN_IFDESCR]byte
}

type PVJ6_OcvF struct {
}

type IFUIVHU struct {
	GreY9QwD8	uint32
	QxY3c1Aj	*byte
	Ar1tuCAxQ	uint32
	PzjAsFXTQJ0O	*PVJ6_OcvF
	ZaRrj0vbv	SRlvVjrYa
}

type SoXj8D struct {
	CCBhzI	uint32
	Aua_QYlZhhg	IM2F8h9S7N8
	Vy_S6Rb2uJ	uint32
	GeaDfO0JqOi	**TP2zoASkA
	LCMbhAI59a	uint32
	D5frq6	**SoXj8D
	Iu0Z95L6hmD	uint32
	FNctng1kw	uint32
}

type ZTNW3vS struct {
}

type TP2zoASkA struct {
	IigGGEfdL3Ks	uint32
	Ze9CHpq	IM2F8h9S7N8
	SBJizeZuM	uint32
	PUwhdGwb8	**P2enQfXhMR
	XnxUpmU_	*ZTNW3vS
	AsLC2XrN	uint32
	GjrSD4	uint32
}

type P2enQfXhMR struct {
	A0OVDU0w	uint32
	LiRo90p	*IFUIVHU
	PiH_qS9j	IM2F8h9S7N8
	HtBt9_U0Azai	*UsPUg9Z7be
	NWCAcxzf	*SJb_kQrDljE
	Gaji1X	*SJb_kQrDljE
	KAzFqWr	*uint16
}

type QWRrtcN struct {
}

type UsPUg9Z7be struct {
	KQBFCEAC	uint32
	IdF0SK8x2c9	uint32
	QJxugR	*byte
	CQ3EV1ut_8Xn	VjfACnL
	I6a_H7	uint32
	B9LGX6aZ	uint32
	JyL6_S0YLwWY	*QWRrtcN
}

type IM2F8h9S7N8 struct {
	SQVHdB	uint32
	CKWx8VS3	uint32
}

type JIRGat32 struct {
	IVaRH8	uint32
	ZjhGHW7s7RC8	SJb_kQrDljE
}

type SJb_kQrDljE struct {
	CXqlMzEblRJu	uint32
	Rg3Wo0	**byte
}

type NyBWtFnlv struct {
	GVbzbU	uint32
	T570hE	JIRGat32
	UqjPLl	JIRGat32
	IqoRTVG	uint32
	Rsf4EqNZpXv2	uint32
	W_b5FmSt7Je	uint32
	D8WPiUcyyUi	*T8WbUqAC3v
}

type XssyNphTC struct {
	U8rBcVn9aQng	uint32
	BqJ_Q454aQB	uint32
	JeR3iDK6	VjfACnL
}

type KdSfm507R struct {
	SN3Eh46LaqF	uint32
	HW8xC2	uint32
	GTOCQBysL	uint32
	LUIIuXGheoK	*uint16
}

type UvCDcg6fLG struct {
	UsRxSXEqxnL	uint32
	G8jD6ui	uint32
	Cc2wIkTaj	uint32
	FUDQ8s00ehvt	uint32
	G9ztD_HQ	VjfACnL
}

const (
	
	HKEY_CLASSES_ROOT	= 0x80000000 + iota
	HKEY_CURRENT_USER
	HKEY_LOCAL_MACHINE
	HKEY_USERS
	HKEY_PERFORMANCE_DATA
	HKEY_CURRENT_CONFIG
	HKEY_DYN_DATA

	KEY_QUERY_VALUE	= 1
	KEY_SET_VALUE	= 2
	KEY_CREATE_SUB_KEY	= 4
	KEY_ENUMERATE_SUB_KEYS	= 8
	KEY_NOTIFY	= 16
	KEY_CREATE_LINK	= 32
	KEY_WRITE	= 0x20006
	KEY_EXECUTE	= 0x20019
	KEY_READ	= 0x20019
	KEY_WOW64_64KEY	= 0x0100
	KEY_WOW64_32KEY	= 0x0200
	KEY_ALL_ACCESS	= 0xf003f
)

const (
	
	REG_NONE	= iota
	REG_SZ
	REG_EXPAND_SZ
	REG_BINARY
	REG_DWORD_LITTLE_ENDIAN
	REG_DWORD_BIG_ENDIAN
	REG_LINK
	REG_MULTI_SZ
	REG_RESOURCE_LIST
	REG_FULL_RESOURCE_DESCRIPTOR
	REG_RESOURCE_REQUIREMENTS_LIST
	REG_QWORD_LITTLE_ENDIAN
	REG_DWORD	= REG_DWORD_LITTLE_ENDIAN
	REG_QWORD	= REG_QWORD_LITTLE_ENDIAN
)

type HnpJdSpZM__0 struct {
	H1nW5M	int32
	WLUX0U	int32
	C_dz1S	int32
	Ax02FCIziFnj	int32
	W1fLCXJR	uintptr
	MO9nR_6aa_2I	*uint16
	JhafcUtas1O	VjfACnL
	Aye498L2D	*HnpJdSpZM__0
}

const (
	AI_PASSIVE	= 1
	AI_CANONNAME	= 2
	AI_NUMERICHOST	= 4
)

type RlB0Ha5cnNd struct {
	EpYSt8TqI	uint32
	UHbfXLLNS6	uint16
	GAYsFb	uint16
	Z9vXwokyJ	[8]byte
}

var XJR3ag = RlB0Ha5cnNd{
	0x25a207b9,
	0xddf3,
	0x4660,
	[8]byte{0x8e, 0xe9, 0x76, 0xe5, 0x8c, 0x74, 0x06, 0x3e},
}

const (
	FILE_SKIP_COMPLETION_PORT_ON_SUCCESS	= 1
	FILE_SKIP_SET_EVENT_ON_HANDLE	= 2
)

const (
	WSAPROTOCOL_LEN	= 255
	MAX_PROTOCOL_CHAIN	= 7
	BASE_PROTOCOL	= 1
	LAYERED_PROTOCOL	= 0

	XP1_CONNECTIONLESS	= 0x00000001
	XP1_GUARANTEED_DELIVERY	= 0x00000002
	XP1_GUARANTEED_ORDER	= 0x00000004
	XP1_MESSAGE_ORIENTED	= 0x00000008
	XP1_PSEUDO_STREAM	= 0x00000010
	XP1_GRACEFUL_CLOSE	= 0x00000020
	XP1_EXPEDITED_DATA	= 0x00000040
	XP1_CONNECT_DATA	= 0x00000080
	XP1_DISCONNECT_DATA	= 0x00000100
	XP1_SUPPORT_BROADCAST	= 0x00000200
	XP1_SUPPORT_MULTIPOINT	= 0x00000400
	XP1_MULTIPOINT_CONTROL_PLANE	= 0x00000800
	XP1_MULTIPOINT_DATA_PLANE	= 0x00001000
	XP1_QOS_SUPPORTED	= 0x00002000
	XP1_UNI_SEND	= 0x00008000
	XP1_UNI_RECV	= 0x00010000
	XP1_IFS_HANDLES	= 0x00020000
	XP1_PARTIAL_MESSAGE	= 0x00040000
	XP1_SAN_SUPPORT_SDP	= 0x00080000

	PFL_MULTIPLE_PROTO_ENTRIES	= 0x00000001
	PFL_RECOMMENDED_PROTO_ENTRY	= 0x00000002
	PFL_HIDDEN	= 0x00000004
	PFL_MATCHES_PROTOCOL_ZERO	= 0x00000008
	PFL_NETWORKDIRECT_PROVIDER	= 0x00000010
)

type ZEPlE2KjP71 struct {
	KRYxJfq	uint32
	NfspU3	uint32
	TzvbXAPWsBW	uint32
	X99hcwTENt	uint32
	AJwAXrIk57v	uint32
	Ikly6QyJ6	RlB0Ha5cnNd
	ZNiWvS_rE	uint32
	It4SFDV	YPstJ8NuPqfB
	ZHQc12	int32
	JbR0AsT	int32
	U816HSb	int32
	VKOP0_LoZ	int32
	MbdoRf	int32
	ONvjYg_gS	int32
	ElXZRZaR5e	int32
	HOapoHT	int32
	IGpixrpF1	int32
	AIlJIU	uint32
	ZiKP9VWjg	uint32
	IHGeHHMpnlwc	[WSAPROTOCOL_LEN + 1]uint16
}

type YPstJ8NuPqfB struct {
	DGeN95smOQ	int32
	EEaPmYaJeB7y	[MAX_PROTOCOL_CHAIN]uint32
}

type BAvoqPPcx struct {
	GrZd5gL0b	uint32
	MkjjM8DA	uint32
	AH6NyXZ333xQ	uint32
}

type qgA9X8dVIL struct {
	O2kMCXRqJjg	uint16
	DwansDx	uint16
	EXW941kPU	uint16
	MoWBo1VLaY	uint16
	D4daKwFX	uint32
	TKCfxANSN	[1]uint16
}

type i9NdjG struct {
	NTZIOHn8z	uint16
	FJV5If	uint16
	DLN9jIh4l	uint16
	SCwRyx_Oh	uint16
	HbIKNUMDT4V	[1]uint16
}

type aUaISxWx2q1 struct {
	Ct81WfAEIFFo	uint32
	AcnzYBjp	uint16
	CUBx3a906CI	uint16

	
	ofwNwEb5aYf	byte
}

const (
	FSCTL_GET_REPARSE_POINT	= 0x900A8
	MAXIMUM_REPARSE_DATA_BUFFER_SIZE	= 16 * 1024
	_IO_REPARSE_TAG_MOUNT_POINT	= 0xA0000003
	IO_REPARSE_TAG_SYMLINK	= 0xA000000C
	SYMBOLIC_LINK_FLAG_DIRECTORY	= 0x1
	_SYMLINK_FLAG_RELATIVE	= 1
)

const UNIX_PATH_MAX = 108	
