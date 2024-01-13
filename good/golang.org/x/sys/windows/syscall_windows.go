//line :1
package NJ4MerJ

import (
	errorspkg "iAHoxjmM"
	fmt "kFsoCfy5zWG"
	"runtime"
	sync "sync"
	syscall "bUKeamF"
	time "fRAfQd_"
	utf16 "DtJCLKevRp"
	"unsafe"
)

type L2L8P9WaNZ uintptr
type IJltG1D5QG uintptr

const (
	InvalidHandle	= ^ /*line :1*/ L2L8P9WaNZ(0)
	InvalidHWND	= ^ /*line :1*/ IJltG1D5QG(0)

	
	DDD_EXACT_MATCH_ON_REMOVE	= 0x00000004
	DDD_NO_BROADCAST_SYSTEM	= 0x00000008
	DDD_RAW_TARGET_PATH	= 0x00000001
	DDD_REMOVE_DEFINITION	= 0x00000002

	
	DRIVE_UNKNOWN	= 0
	DRIVE_NO_ROOT_DIR	= 1
	DRIVE_REMOVABLE	= 2
	DRIVE_FIXED	= 3
	DRIVE_REMOTE	= 4
	DRIVE_CDROM	= 5
	DRIVE_RAMDISK	= 6

	
	FILE_CASE_SENSITIVE_SEARCH	= 0x00000001
	FILE_CASE_PRESERVED_NAMES	= 0x00000002
	FILE_FILE_COMPRESSION	= 0x00000010
	FILE_DAX_VOLUME	= 0x20000000
	FILE_NAMED_STREAMS	= 0x00040000
	FILE_PERSISTENT_ACLS	= 0x00000008
	FILE_READ_ONLY_VOLUME	= 0x00080000
	FILE_SEQUENTIAL_WRITE_ONCE	= 0x00100000
	FILE_SUPPORTS_ENCRYPTION	= 0x00020000
	FILE_SUPPORTS_EXTENDED_ATTRIBUTES	= 0x00800000
	FILE_SUPPORTS_HARD_LINKS	= 0x00400000
	FILE_SUPPORTS_OBJECT_IDS	= 0x00010000
	FILE_SUPPORTS_OPEN_BY_FILE_ID	= 0x01000000
	FILE_SUPPORTS_REPARSE_POINTS	= 0x00000080
	FILE_SUPPORTS_SPARSE_FILES	= 0x00000040
	FILE_SUPPORTS_TRANSACTIONS	= 0x00200000
	FILE_SUPPORTS_USN_JOURNAL	= 0x02000000
	FILE_UNICODE_ON_DISK	= 0x00000004
	FILE_VOLUME_IS_COMPRESSED	= 0x00008000
	FILE_VOLUME_QUOTAS	= 0x00000020

	
	LOCKFILE_FAIL_IMMEDIATELY	= 0x00000001
	LOCKFILE_EXCLUSIVE_LOCK	= 0x00000002

	
	WAIT_IO_COMPLETION	= 0x000000C0
)




func JHt9df(bamc83qA3DBr string) []uint16 {
	xINVctm, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(bamc83qA3DBr)
	if jeWMpOaQtWV != nil {
		 /*line :1*/ panic("windows: string with NUL passed to StringToUTF16")
	}
	return xINVctm
}




func DJN4EMpp7(bamc83qA3DBr string) ([]uint16, error) {
	return  /*line :1*/ syscall.ZxIMCrnjQ5U4(bamc83qA3DBr)
}



func OXNanQ8Uj(bamc83qA3DBr []uint16) string {
	return  /*line :1*/ syscall.AODVXp8NM3sd(bamc83qA3DBr)
}




func RRlduq3XI(bamc83qA3DBr string) *uint16	{ return & /*line :1*/ JHt9df(bamc83qA3DBr)[0] }




func Ih3Y4u(bamc83qA3DBr string) (*uint16, error) {
	xINVctm, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(bamc83qA3DBr)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	return &xINVctm[0], nil
}




func M_Sea9j(hD4wNgEB *uint16) string {
	if hD4wNgEB == nil {
		return ""
	}
	if *hD4wNgEB == 0 {
		return ""
	}

	krzuku := 0
	for nauDv3A :=  /*line :1*/ unsafe.Pointer(hD4wNgEB); *(* /*line :1*/ uint16)(nauDv3A) != 0; krzuku++ {
		nauDv3A =  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(nauDv3A) +  /*line :1*/ unsafe.Sizeof(*hD4wNgEB))
	}

	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8( /*line :1*/ unsafe.Slice(hD4wNgEB, krzuku)))
}

func Lx401FXe9Z9() int	{ return 4096 }




func KeU7MRSmhTF1(hDWQ1ww interface{}) uintptr {
	return  /*line :1*/ syscall.VSPtNpjXc4(hDWQ1ww)
}




func EqGW1VN2zR(hDWQ1ww interface{}) uintptr {
	return  /*line :1*/ syscall.MLY71IBA(hDWQ1ww)
}







func XeI8xfjw() (L2L8P9WaNZ, error) {
	return  /*line :1*/ RwYrTC0pJaxW(), nil
}



func RwYrTC0pJaxW() L2L8P9WaNZ	{ return  /*line :1*/ L2L8P9WaNZ(^ /*line :1*/ uintptr(1 - 1)) }







func ErnHVa5r() (L2L8P9WaNZ, error) {
	return  /*line :1*/ J88pcGgbRf(), nil
}



func J88pcGgbRf() L2L8P9WaNZ	{ return  /*line :1*/ L2L8P9WaNZ(^ /*line :1*/ uintptr(2 - 1)) }



func UqxfQM(tiLd7g0 L2L8P9WaNZ, bxnorpA24 uintptr) (aTSUoJ40C uintptr, jeWMpOaQtWV error) {
	gPQ4CDCVeA, _, j4dSamRp :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ phplFjs8e.Addr(), 2,  /*line :1*/ uintptr(tiLd7g0), bxnorpA24, 0)
	aTSUoJ40C =  /*line :1*/ uintptr(gPQ4CDCVeA)
	if aTSUoJ40C == 0 {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return
}

func YN0dOkhBR92(j_J88hs int)	{  /*line :1*/ JGEMyXm8( /*line :1*/ uint32(j_J88hs)) }

func zOXA59NPo() *NRXOeZVS9w {
	var aRcafgWyH NRXOeZVS9w
	aRcafgWyH.Ec1Ig0d =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(aRcafgWyH))
	aRcafgWyH.Z73_4lh2ro = 1
	return &aRcafgWyH
}

func IFE33vUVFmlw(bZKbGTxeit string, yAZay4eEB0M int, iak6OL uint32) (nuvOdY1SQ9I5 L2L8P9WaNZ, jeWMpOaQtWV error) {
	if  /*line :1*/ len(bZKbGTxeit) == 0 {
		return InvalidHandle, ERROR_FILE_NOT_FOUND
	}
	kOQCML, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if jeWMpOaQtWV != nil {
		return InvalidHandle, jeWMpOaQtWV
	}
	var oW8M6Qq uint32
	switch yAZay4eEB0M & (O_RDONLY | O_WRONLY | O_RDWR) {
	case O_RDONLY:
		oW8M6Qq = GENERIC_READ
	case O_WRONLY:
		oW8M6Qq = GENERIC_WRITE
	case O_RDWR:
		oW8M6Qq = GENERIC_READ | GENERIC_WRITE
	}
	if yAZay4eEB0M&O_CREAT != 0 {
		oW8M6Qq |= GENERIC_WRITE
	}
	if yAZay4eEB0M&O_APPEND != 0 {
		oW8M6Qq &^= GENERIC_WRITE
		oW8M6Qq |= FILE_APPEND_DATA
	}
	hQEQF3 :=  /*line :1*/ uint32(FILE_SHARE_READ | FILE_SHARE_WRITE)
	var aRcafgWyH *NRXOeZVS9w
	if yAZay4eEB0M&O_CLOEXEC == 0 {
		aRcafgWyH =  /*line :1*/ zOXA59NPo()
	}
	var svCwoQOx8aBm uint32
	switch {
	case yAZay4eEB0M&(O_CREAT|O_EXCL) == (O_CREAT | O_EXCL):
		svCwoQOx8aBm = CREATE_NEW
	case yAZay4eEB0M&(O_CREAT|O_TRUNC) == (O_CREAT | O_TRUNC):
		svCwoQOx8aBm = CREATE_ALWAYS
	case yAZay4eEB0M&O_CREAT == O_CREAT:
		svCwoQOx8aBm = OPEN_ALWAYS
	case yAZay4eEB0M&O_TRUNC == O_TRUNC:
		svCwoQOx8aBm = TRUNCATE_EXISTING
	default:
		svCwoQOx8aBm = OPEN_EXISTING
	}
	var jQSvSZ uint32 = FILE_ATTRIBUTE_NORMAL
	if iak6OL&S_IWRITE == 0 {
		jQSvSZ = FILE_ATTRIBUTE_READONLY
	}
	_ITTe4, aYtmqRVkqc_8 :=  /*line :1*/ LNhQWvTnm6(kOQCML, oW8M6Qq, hQEQF3, aRcafgWyH, svCwoQOx8aBm, jQSvSZ, 0)
	return _ITTe4, aYtmqRVkqc_8
}

func A6AyKO5gS(nuvOdY1SQ9I5 L2L8P9WaNZ, hD4wNgEB []byte) (krzuku int, jeWMpOaQtWV error) {
	var iX0cRH5st uint32
	aYtmqRVkqc_8 :=  /*line :1*/ AWSfZRIMg8am(nuvOdY1SQ9I5, hD4wNgEB, &iX0cRH5st, nil)
	if aYtmqRVkqc_8 != nil {
		if aYtmqRVkqc_8 == ERROR_BROKEN_PIPE {

			return 0, nil
		}
		return 0, aYtmqRVkqc_8
	}
	return  /*line :1*/ int(iX0cRH5st), nil
}

func F0GmsKR3Re(nuvOdY1SQ9I5 L2L8P9WaNZ, hD4wNgEB []byte) (krzuku int, jeWMpOaQtWV error) {
	if raceenabled {
		 /*line :1*/ pmoEPrE6E( /*line :1*/ unsafe.Pointer(&kvc_7PTCf2bh))
	}
	var iX0cRH5st uint32
	aYtmqRVkqc_8 :=  /*line :1*/ DfisMN(nuvOdY1SQ9I5, hD4wNgEB, &iX0cRH5st, nil)
	if aYtmqRVkqc_8 != nil {
		return 0, aYtmqRVkqc_8
	}
	return  /*line :1*/ int(iX0cRH5st), nil
}

func AWSfZRIMg8am(nuvOdY1SQ9I5 L2L8P9WaNZ, hD4wNgEB []byte, iX0cRH5st *uint32, g1sgEzk2pZh *Ctdynv) error {
	jeWMpOaQtWV :=  /*line :1*/ oa1U_Bh2x(nuvOdY1SQ9I5, hD4wNgEB, iX0cRH5st, g1sgEzk2pZh)
	if raceenabled {
		if *iX0cRH5st > 0 {
			 /*line :1*/ i_R8EfXI( /*line :1*/ unsafe.Pointer(&hD4wNgEB[0]),  /*line :1*/ int(*iX0cRH5st))
		}
		 /*line :1*/ ord0GA1( /*line :1*/ unsafe.Pointer(&kvc_7PTCf2bh))
	}
	return jeWMpOaQtWV
}

func DfisMN(nuvOdY1SQ9I5 L2L8P9WaNZ, hD4wNgEB []byte, iX0cRH5st *uint32, g1sgEzk2pZh *Ctdynv) error {
	if raceenabled {
		 /*line :1*/ pmoEPrE6E( /*line :1*/ unsafe.Pointer(&kvc_7PTCf2bh))
	}
	jeWMpOaQtWV :=  /*line :1*/ y8wyCmHs(nuvOdY1SQ9I5, hD4wNgEB, iX0cRH5st, g1sgEzk2pZh)
	if raceenabled && *iX0cRH5st > 0 {
		 /*line :1*/ wJmAmL7jZa( /*line :1*/ unsafe.Pointer(&hD4wNgEB[0]),  /*line :1*/ int(*iX0cRH5st))
	}
	return jeWMpOaQtWV
}

var kvc_7PTCf2bh int64

func MIPoY7oQ(nuvOdY1SQ9I5 L2L8P9WaNZ, _exc64aC int64, h5r9QaWpQ int) (hi6Xio int64, jeWMpOaQtWV error) {
	var mwjEM9EsW uint32
	switch h5r9QaWpQ {
	case 0:
		mwjEM9EsW = FILE_BEGIN
	case 1:
		mwjEM9EsW = FILE_CURRENT
	case 2:
		mwjEM9EsW = FILE_END
	}
	aSRZ2MMhOO7 :=  /*line :1*/ int32(_exc64aC >> 32)
	jnpgesBlk8F :=  /*line :1*/ int32(_exc64aC)

	hqY3fP2, _ :=  /*line :1*/ WKG3GcD(nuvOdY1SQ9I5)
	if hqY3fP2 == FILE_TYPE_PIPE {
		return 0, syscall.EPIPE
	}
	ahzfaF9j, aYtmqRVkqc_8 :=  /*line :1*/ Gl0s_hQDBoz(nuvOdY1SQ9I5, jnpgesBlk8F, &aSRZ2MMhOO7, mwjEM9EsW)
	if aYtmqRVkqc_8 != nil {
		return 0, aYtmqRVkqc_8
	}
	return  /*line :1*/ int64(aSRZ2MMhOO7)<<32 +  /*line :1*/ int64(ahzfaF9j), nil
}

func ZtFIbWB(nuvOdY1SQ9I5 L2L8P9WaNZ) (jeWMpOaQtWV error) {
	return  /*line :1*/ E5j7kDrdt(nuvOdY1SQ9I5)
}

var (
	IbEqwzUADz0N	=  /*line :1*/ gYuyKtMC33PR(STD_INPUT_HANDLE)
	BI67C0	=  /*line :1*/ gYuyKtMC33PR(STD_OUTPUT_HANDLE)
	D7YyXTO1C	=  /*line :1*/ gYuyKtMC33PR(STD_ERROR_HANDLE)
)

func gYuyKtMC33PR(bWBOjCD0O uint32) (nuvOdY1SQ9I5 L2L8P9WaNZ) {
	pHRnBlwkei, _ :=  /*line :1*/ CoJkxiYw(bWBOjCD0O)
	return pHRnBlwkei
}

const ImplementsGetwd = true

func DiBjtD() (mvaxnSHe3 string, jeWMpOaQtWV error) {
	hJuLpmVFi :=  /*line :1*/ make([]uint16, 300)
	krzuku, aYtmqRVkqc_8 :=  /*line :1*/ M9mJEOy_P( /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)), &hJuLpmVFi[0])
	if aYtmqRVkqc_8 != nil {
		return "", aYtmqRVkqc_8
	}
	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(hJuLpmVFi[0:krzuku])), nil
}

func Juob3PE0(bZKbGTxeit string) (jeWMpOaQtWV error) {
	kOQCML, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ PJdCxj(kOQCML)
}

func HlrJu45IcHJ(bZKbGTxeit string, yAZay4eEB0M uint32) (jeWMpOaQtWV error) {
	kOQCML, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ ZGfr8zMaI(kOQCML, nil)
}

func G16biwecbj(bZKbGTxeit string) (jeWMpOaQtWV error) {
	kOQCML, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ IR_bmXo78A(kOQCML)
}

func BV50ThTsc4s(bZKbGTxeit string) (jeWMpOaQtWV error) {
	kOQCML, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ F9XQJe2wt5oP(kOQCML)
}

func Z1JuGphGy(g3QX6O1Z, vqYv72aSS string) (jeWMpOaQtWV error) {
	wPbtqIJl, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(g3QX6O1Z)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	xtbcgl1U7, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(vqYv72aSS)
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ CjF0vhbT3(wPbtqIJl, xtbcgl1U7, MOVEFILE_REPLACE_EXISTING)
}

func JltnmO() (gYwswmeUjG string, jeWMpOaQtWV error) {
	var krzuku uint32 = MAX_COMPUTERNAME_LENGTH + 1
	hJuLpmVFi :=  /*line :1*/ make([]uint16, krzuku)
	aYtmqRVkqc_8 :=  /*line :1*/ IwxAIk(&hJuLpmVFi[0], &krzuku)
	if aYtmqRVkqc_8 != nil {
		return "", aYtmqRVkqc_8
	}
	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(hJuLpmVFi[0:krzuku])), nil
}

func VlL9_aQ() time.GKMwTGxQa0S {
	return  /*line :1*/ time.GKMwTGxQa0S( /*line :1*/ ol7ABdPCo()) * time.Millisecond
}

func P6tHtK6bxC(nuvOdY1SQ9I5 L2L8P9WaNZ, fORnBp int64) (jeWMpOaQtWV error) {
	p9ryVe1, aYtmqRVkqc_8 :=  /*line :1*/ MIPoY7oQ(nuvOdY1SQ9I5, 0, 1)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	defer  /*line :1*/ MIPoY7oQ(nuvOdY1SQ9I5, p9ryVe1, 0)
	_, aYtmqRVkqc_8 =  /*line :1*/ MIPoY7oQ(nuvOdY1SQ9I5, fORnBp, 0)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	aYtmqRVkqc_8 =  /*line :1*/ RAGLhX7qwfg(nuvOdY1SQ9I5)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	return nil
}

func AHwBiZ(fmaieegZe2E_ *F2i2x7peITi) (jeWMpOaQtWV error) {
	var hqY3fP2 ZPa9KL2
	 /*line :1*/ EavnAwBq(&hqY3fP2)
	*fmaieegZe2E_ =  /*line :1*/ RKVvyg( /*line :1*/ hqY3fP2.Nanoseconds())
	return nil
}

func LhTDyLqWheq(hD4wNgEB []L2L8P9WaNZ) (jeWMpOaQtWV error) {
	if  /*line :1*/ len(hD4wNgEB) != 2 {
		return syscall.EINVAL
	}
	var pHRnBlwkei, mwjEM9EsW L2L8P9WaNZ
	aYtmqRVkqc_8 :=  /*line :1*/ YwbL9p6NalN(&pHRnBlwkei, &mwjEM9EsW,  /*line :1*/ zOXA59NPo(), 0)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	hD4wNgEB[0] = pHRnBlwkei
	hD4wNgEB[1] = mwjEM9EsW
	return nil
}

func Kmcpt6AcRyN(bZKbGTxeit string, fmaieegZe2E_ []F2i2x7peITi) (jeWMpOaQtWV error) {
	if  /*line :1*/ len(fmaieegZe2E_) != 2 {
		return syscall.EINVAL
	}
	kOQCML, aYtmqRVkqc_8 :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	_ITTe4, aYtmqRVkqc_8 :=  /*line :1*/ LNhQWvTnm6(kOQCML,
		FILE_WRITE_ATTRIBUTES, FILE_SHARE_WRITE, nil,
		OPEN_EXISTING, FILE_FLAG_BACKUP_SEMANTICS, 0)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	defer  /*line :1*/ E5j7kDrdt(_ITTe4)
	xINVctm :=  /*line :1*/ WAhU8nBDHc( /*line :1*/ fmaieegZe2E_[0].Nanoseconds())
	mwjEM9EsW :=  /*line :1*/ WAhU8nBDHc( /*line :1*/ fmaieegZe2E_[1].Nanoseconds())
	return  /*line :1*/ HyQcEQ4lX4wD(_ITTe4, nil, &xINVctm, &mwjEM9EsW)
}

func A5kEI_LnMf(bZKbGTxeit string, r7ayG7dk []YWaK2Afs) (jeWMpOaQtWV error) {
	if  /*line :1*/ len(r7ayG7dk) != 2 {
		return syscall.EINVAL
	}
	kOQCML, aYtmqRVkqc_8 :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	_ITTe4, aYtmqRVkqc_8 :=  /*line :1*/ LNhQWvTnm6(kOQCML,
		FILE_WRITE_ATTRIBUTES, FILE_SHARE_WRITE, nil,
		OPEN_EXISTING, FILE_FLAG_BACKUP_SEMANTICS, 0)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	defer  /*line :1*/ E5j7kDrdt(_ITTe4)
	xINVctm :=  /*line :1*/ WAhU8nBDHc( /*line :1*/ BZPG6Sd_Z2nW(r7ayG7dk[0]))
	mwjEM9EsW :=  /*line :1*/ WAhU8nBDHc( /*line :1*/ BZPG6Sd_Z2nW(r7ayG7dk[1]))
	return  /*line :1*/ HyQcEQ4lX4wD(_ITTe4, nil, &xINVctm, &mwjEM9EsW)
}

func XAhkABa(nuvOdY1SQ9I5 L2L8P9WaNZ) (jeWMpOaQtWV error) {
	return  /*line :1*/ FVgKyMq9Ds1(nuvOdY1SQ9I5)
}

func BrBigSw(bZKbGTxeit string, yAZay4eEB0M uint32) (jeWMpOaQtWV error) {
	hD4wNgEB, aYtmqRVkqc_8 :=  /*line :1*/ Ih3Y4u(bZKbGTxeit)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	jQSvSZ, aYtmqRVkqc_8 :=  /*line :1*/ FTfH6CHf_Xao(hD4wNgEB)
	if aYtmqRVkqc_8 != nil {
		return aYtmqRVkqc_8
	}
	if yAZay4eEB0M&S_IWRITE != 0 {
		jQSvSZ &^= FILE_ATTRIBUTE_READONLY
	} else {
		jQSvSZ |= FILE_ATTRIBUTE_READONLY
	}
	return  /*line :1*/ JrSRwp(hD4wNgEB, jQSvSZ)
}

func KwKQdOH() error {
	return  /*line :1*/ hMcQrE.Find()
}

func G1_6xZ() error {
	return  /*line :1*/ ggsdamqooaG.Find()
}

func RBYZTy() error {
	return  /*line :1*/ pMd7CHtSuFpJ.Find()
}

func SOq8sdDY(g2LnxiOU53rM []L2L8P9WaNZ, bnT6r4 bool, keuTsec uint32) (g_3oYJh6F uint32, jeWMpOaQtWV error) {

	var t3u7G5Dk11ZE *L2L8P9WaNZ
	if  /*line :1*/ len(g2LnxiOU53rM) > 0 {
		t3u7G5Dk11ZE = &g2LnxiOU53rM[0]
	}
	return  /*line :1*/ bERNe5dqR( /*line :1*/ uint32( /*line :1*/ len(g2LnxiOU53rM)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(t3u7G5Dk11ZE)), bnT6r4, keuTsec)
}

const socket_error =  /*line :1*/ uintptr(^ /*line :1*/ uint32(0))



var O_FfMmlLVX4e bool

type GWQ5aAfbG struct {
	JgQng5oWY8Ul	uint16
	VqKTvvibG	uint16
	AGDgR2QSLqW	[4]byte	
	WeOsNeM	[8]uint8
}

type BysFg5KG0iB struct {
	XYs3jaLEe	uint16
	D_H8WxCL7	uint16
	RwOjhxNykN	uint32
	G1aNyX7	[16]byte	
	Jl_oKnHTc	uint32
}

type BceY5d5Mt7z struct {
	LcXKM5eE7rv	uint16
	WFMxTSzZGTq	[14]int8
}

type I6IZfmg struct {
	AlZX7rZ	BceY5d5Mt7z
	NnIJuyS	[100]int8
}

type VPXahxqr5 interface {
	uaMPM0() (nauDv3A unsafe.Pointer, ajJiI4KvLI int32, jeWMpOaQtWV error)	
}

type E4o5f3MH struct {
	Haxcwf	int
	Dy7sqN	[4]byte
	tQaI196Y	GWQ5aAfbG
}

func (aRcafgWyH *E4o5f3MH) uaMPM0() (unsafe.Pointer, int32, error) {
	if aRcafgWyH.Haxcwf < 0 || aRcafgWyH.Haxcwf > 0xFFFF {
		return nil, 0, syscall.EINVAL
	}
	aRcafgWyH.tQaI196Y.JgQng5oWY8Ul = AF_INET
	hD4wNgEB := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&aRcafgWyH.tQaI196Y.VqKTvvibG))
	hD4wNgEB[0] =  /*line :1*/ byte(aRcafgWyH.Haxcwf >> 8)
	hD4wNgEB[1] =  /*line :1*/ byte(aRcafgWyH.Haxcwf)
	aRcafgWyH.tQaI196Y.AGDgR2QSLqW = aRcafgWyH.Dy7sqN
	return  /*line :1*/ unsafe.Pointer(&aRcafgWyH.tQaI196Y),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(aRcafgWyH.tQaI196Y)), nil
}

type HKZOlk struct {
	U2IFJq	int
	FcbdKjMY	uint32
	FUXcvCUA	[16]byte
	oYd9xLWcP	BysFg5KG0iB
}

func (aRcafgWyH *HKZOlk) uaMPM0() (unsafe.Pointer, int32, error) {
	if aRcafgWyH.U2IFJq < 0 || aRcafgWyH.U2IFJq > 0xFFFF {
		return nil, 0, syscall.EINVAL
	}
	aRcafgWyH.oYd9xLWcP.XYs3jaLEe = AF_INET6
	hD4wNgEB := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&aRcafgWyH.oYd9xLWcP.D_H8WxCL7))
	hD4wNgEB[0] =  /*line :1*/ byte(aRcafgWyH.U2IFJq >> 8)
	hD4wNgEB[1] =  /*line :1*/ byte(aRcafgWyH.U2IFJq)
	aRcafgWyH.oYd9xLWcP.Jl_oKnHTc = aRcafgWyH.FcbdKjMY
	aRcafgWyH.oYd9xLWcP.G1aNyX7 = aRcafgWyH.FUXcvCUA
	return  /*line :1*/ unsafe.Pointer(&aRcafgWyH.oYd9xLWcP),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(aRcafgWyH.oYd9xLWcP)), nil
}

type F8WTes struct {
	ZCY53FT	uint16
	N354aK0	[UNIX_PATH_MAX]int8
}

type CRVhBLUFJ struct {
	AoSxtLw	string
	sqMAqvOOkx	F8WTes
}

func (aRcafgWyH *CRVhBLUFJ) uaMPM0() (unsafe.Pointer, int32, error) {
	gYwswmeUjG := aRcafgWyH.AoSxtLw
	krzuku :=  /*line :1*/ len(gYwswmeUjG)
	if krzuku >  /*line :1*/ len(aRcafgWyH.sqMAqvOOkx.N354aK0) {
		return nil, 0, syscall.EINVAL
	}
	if krzuku ==  /*line :1*/ len(aRcafgWyH.sqMAqvOOkx.N354aK0) && gYwswmeUjG[0] != '@' {
		return nil, 0, syscall.EINVAL
	}
	aRcafgWyH.sqMAqvOOkx.ZCY53FT = AF_UNIX
	for rRGfxgPH8Kq := 0; rRGfxgPH8Kq < krzuku; rRGfxgPH8Kq++ {
		aRcafgWyH.sqMAqvOOkx.N354aK0[rRGfxgPH8Kq] =  /*line :1*/ int8(gYwswmeUjG[rRGfxgPH8Kq])
	}

	aqgNlsBU3y :=  /*line :1*/ int32(2)
	if krzuku > 0 {
		aqgNlsBU3y +=  /*line :1*/ int32(krzuku) + 1
	}
	if aRcafgWyH.sqMAqvOOkx.N354aK0[0] == '@' || (aRcafgWyH.sqMAqvOOkx.N354aK0[0] == 0 && aqgNlsBU3y > 3) {

		aRcafgWyH.sqMAqvOOkx.N354aK0[0] = 0

		aqgNlsBU3y--
	}

	return  /*line :1*/ unsafe.Pointer(&aRcafgWyH.sqMAqvOOkx), aqgNlsBU3y, nil
}

type WkPUxD9c10Gu struct {
	Jy3bhz	[2]byte
	M5Ok6c5Ku	[8]byte
	CablW3h0B	[16]byte
	Kveaik924	[4]byte
}

type Uz7UkWkAf struct {
	ZzP5BOCc	uint64
	VnP98huN0Zl	Kw2qafRFaiLg
	P2Ao3ytoZ	uint32

	wlIGxVR	WkPUxD9c10Gu
}

func (aRcafgWyH *Uz7UkWkAf) uaMPM0() (unsafe.Pointer, int32, error) {
	jroU3W5g := AF_BTH
	aRcafgWyH.wlIGxVR = WkPUxD9c10Gu{
		Jy3bhz:	*(*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&jroU3W5g)),
		M5Ok6c5Ku:	*(*[8] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&aRcafgWyH.ZzP5BOCc)),
		Kveaik924:	*(*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&aRcafgWyH.P2Ao3ytoZ)),
		CablW3h0B:	*(*[16] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&aRcafgWyH.VnP98huN0Zl)),
	}
	return  /*line :1*/ unsafe.Pointer(&aRcafgWyH.wlIGxVR),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(aRcafgWyH.wlIGxVR)), nil
}

func (u1fJXFk *I6IZfmg) Sockaddr() (VPXahxqr5, error) {
	switch u1fJXFk.AlZX7rZ.LcXKM5eE7rv {
	case AF_UNIX:
		bIcz0Kmh := (* /*line :1*/ F8WTes)( /*line :1*/ unsafe.Pointer(u1fJXFk))
		aRcafgWyH :=  /*line :1*/ new(CRVhBLUFJ)
		if bIcz0Kmh.N354aK0[0] == 0 {

			bIcz0Kmh.N354aK0[0] = '@'
		}

		krzuku := 0
		for krzuku <  /*line :1*/ len(bIcz0Kmh.N354aK0) && bIcz0Kmh.N354aK0[krzuku] != 0 {
			krzuku++
		}
		aRcafgWyH.AoSxtLw =  /*line :1*/ string( /*line :1*/ unsafe.Slice((* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&bIcz0Kmh.N354aK0[0])), krzuku))
		return aRcafgWyH, nil

	case AF_INET:
		bIcz0Kmh := (* /*line :1*/ GWQ5aAfbG)( /*line :1*/ unsafe.Pointer(u1fJXFk))
		aRcafgWyH :=  /*line :1*/ new(E4o5f3MH)
		hD4wNgEB := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&bIcz0Kmh.VqKTvvibG))
		aRcafgWyH.Haxcwf =  /*line :1*/ int(hD4wNgEB[0])<<8 +  /*line :1*/ int(hD4wNgEB[1])
		aRcafgWyH.Dy7sqN = bIcz0Kmh.AGDgR2QSLqW
		return aRcafgWyH, nil

	case AF_INET6:
		bIcz0Kmh := (* /*line :1*/ BysFg5KG0iB)( /*line :1*/ unsafe.Pointer(u1fJXFk))
		aRcafgWyH :=  /*line :1*/ new(HKZOlk)
		hD4wNgEB := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&bIcz0Kmh.D_H8WxCL7))
		aRcafgWyH.U2IFJq =  /*line :1*/ int(hD4wNgEB[0])<<8 +  /*line :1*/ int(hD4wNgEB[1])
		aRcafgWyH.FcbdKjMY = bIcz0Kmh.Jl_oKnHTc
		aRcafgWyH.FUXcvCUA = bIcz0Kmh.G1aNyX7
		return aRcafgWyH, nil
	}
	return nil, syscall.EAFNOSUPPORT
}

func L8xQ9MtwVjW(lfrJYrOPlciq, pgq0s4, pNa9kYooPwQ int) (nuvOdY1SQ9I5 L2L8P9WaNZ, jeWMpOaQtWV error) {
	if lfrJYrOPlciq == AF_INET6 && O_FfMmlLVX4e {
		return InvalidHandle, syscall.EAFNOSUPPORT
	}
	return  /*line :1*/ oxblUQMm03v( /*line :1*/ int32(lfrJYrOPlciq),  /*line :1*/ int32(pgq0s4),  /*line :1*/ int32(pNa9kYooPwQ))
}

func BfHhC9Fdz(nuvOdY1SQ9I5 L2L8P9WaNZ, bR_S51, qUM7_22hM int, wvZBcfh int) (jeWMpOaQtWV error) {
	n8CI2rJ :=  /*line :1*/ int32(wvZBcfh)
	return  /*line :1*/ UfWy_ocZ(nuvOdY1SQ9I5,  /*line :1*/ int32(bR_S51),  /*line :1*/ int32(qUM7_22hM), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&n8CI2rJ)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(n8CI2rJ)))
}

func Gh4Q7d(nuvOdY1SQ9I5 L2L8P9WaNZ, aRcafgWyH VPXahxqr5) (jeWMpOaQtWV error) {
	nauDv3A, krzuku, jeWMpOaQtWV :=  /*line :1*/ aRcafgWyH.uaMPM0()
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ fCe8PKCARpU(nuvOdY1SQ9I5, nauDv3A, krzuku)
}

func L5Pltfn(nuvOdY1SQ9I5 L2L8P9WaNZ, aRcafgWyH VPXahxqr5) (jeWMpOaQtWV error) {
	nauDv3A, krzuku, jeWMpOaQtWV :=  /*line :1*/ aRcafgWyH.uaMPM0()
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ aW9zqIBU4(nuvOdY1SQ9I5, nauDv3A, krzuku)
}

func WP6wDi0(aRcafgWyH VPXahxqr5, y9xEK8 *uint32) (jeWMpOaQtWV error) {
	nauDv3A, _, jeWMpOaQtWV :=  /*line :1*/ aRcafgWyH.uaMPM0()
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ u9EM3rC(nauDv3A, y9xEK8)
}

func L6MYjZ(nuvOdY1SQ9I5 L2L8P9WaNZ) (aRcafgWyH VPXahxqr5, jeWMpOaQtWV error) {
	var u1fJXFk I6IZfmg
	cTpvqfM :=  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(u1fJXFk))
	if jeWMpOaQtWV =  /*line :1*/ jPN3P54q(nuvOdY1SQ9I5, &u1fJXFk, &cTpvqfM); jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ u1fJXFk.Sockaddr()
}

func BRFKJSkwDbG3(nuvOdY1SQ9I5 L2L8P9WaNZ) (aRcafgWyH VPXahxqr5, jeWMpOaQtWV error) {
	var u1fJXFk I6IZfmg
	cTpvqfM :=  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(u1fJXFk))
	if jeWMpOaQtWV =  /*line :1*/ pzZiaon(nuvOdY1SQ9I5, &u1fJXFk, &cTpvqfM); jeWMpOaQtWV != nil {
		return
	}
	return  /*line :1*/ u1fJXFk.Sockaddr()
}

func Ar_Ys58Y(bamc83qA3DBr L2L8P9WaNZ, krzuku int) (jeWMpOaQtWV error) {
	return  /*line :1*/ i3P0kOH(bamc83qA3DBr,  /*line :1*/ int32(krzuku))
}

func U0tdweUmnCl(nuvOdY1SQ9I5 L2L8P9WaNZ, jFffl_H int) (jeWMpOaQtWV error) {
	return  /*line :1*/ aM4XrT(nuvOdY1SQ9I5,  /*line :1*/ int32(jFffl_H))
}

func HIf0UdAPxQE(bamc83qA3DBr L2L8P9WaNZ, gCwH7wEHx *GC5qDJGWUpA, jkqgCfg3a0ds uint32, faR595J6XZ5 *uint32, a_MrGKcrR uint32, xtbcgl1U7 VPXahxqr5, g1sgEzk2pZh *Ctdynv, fgQM_KRAN9 *byte) (jeWMpOaQtWV error) {
	var u1fJXFk unsafe.Pointer
	var cTpvqfM int32
	if xtbcgl1U7 != nil {
		u1fJXFk, cTpvqfM, jeWMpOaQtWV =  /*line :1*/ xtbcgl1U7.uaMPM0()
		if jeWMpOaQtWV != nil {
			return jeWMpOaQtWV
		}
	}
	return  /*line :1*/ Qzva8c0F4tEv(bamc83qA3DBr, gCwH7wEHx, jkqgCfg3a0ds, faR595J6XZ5, a_MrGKcrR, (* /*line :1*/ I6IZfmg)( /*line :1*/ unsafe.Pointer(u1fJXFk)), cTpvqfM, g1sgEzk2pZh, fgQM_KRAN9)
}

func AvZrUpPEpss() error {
	return  /*line :1*/ u6Lcomleb9.Find()
}

var dtf01P5DeJs struct {
	b_cotd	sync.LhBfwn6wa1x
	iqDs38TG9	uintptr
	epGgcJBCHfrF	error
}

func ScRdeN6() error {
	 /*line :1*/ dtf01P5DeJs.b_cotd.Do(func() {
		var bamc83qA3DBr L2L8P9WaNZ
		bamc83qA3DBr, dtf01P5DeJs.epGgcJBCHfrF =  /*line :1*/ L8xQ9MtwVjW(AF_INET, SOCK_STREAM, IPPROTO_TCP)
		if dtf01P5DeJs.epGgcJBCHfrF != nil {
			return
		}
		defer  /*line :1*/ E5j7kDrdt(bamc83qA3DBr)
		var krzuku uint32
		dtf01P5DeJs.epGgcJBCHfrF =  /*line :1*/ VuX432Ao4_L(bamc83qA3DBr,
			SIO_GET_EXTENSION_FUNCTION_POINTER,
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&VN8uzaurX)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(VN8uzaurX)),
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&dtf01P5DeJs.iqDs38TG9)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(dtf01P5DeJs.iqDs38TG9)),
			&krzuku, nil, 0)
	})
	return dtf01P5DeJs.epGgcJBCHfrF
}

func inrAIi3wg(bamc83qA3DBr L2L8P9WaNZ, gYwswmeUjG unsafe.Pointer, jNLsCU int32, dEqx7iRZr *byte, gPeNaj uint32, r2c83N *uint32, g1sgEzk2pZh *Ctdynv) (jeWMpOaQtWV error) {
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.Rc0O05UsV(dtf01P5DeJs.iqDs38TG9, 7,  /*line :1*/ uintptr(bamc83qA3DBr),  /*line :1*/ uintptr(gYwswmeUjG),  /*line :1*/ uintptr(jNLsCU),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(dEqx7iRZr)),  /*line :1*/ uintptr(gPeNaj),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(r2c83N)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)), 0, 0)
	if i_EJOiAVI == 0 {
		if j4dSamRp != 0 {
			jeWMpOaQtWV =  /*line :1*/ error(j4dSamRp)
		} else {
			jeWMpOaQtWV = syscall.EINVAL
		}
	}
	return
}

func RTPq0X(nuvOdY1SQ9I5 L2L8P9WaNZ, aRcafgWyH VPXahxqr5, dEqx7iRZr *byte, gPeNaj uint32, r2c83N *uint32, g1sgEzk2pZh *Ctdynv) error {
	jeWMpOaQtWV :=  /*line :1*/ ScRdeN6()
	if jeWMpOaQtWV != nil {
		return  /*line :1*/ errorspkg.Su6g6hRoi9X("failed to find ConnectEx: " +  /*line :1*/ jeWMpOaQtWV.Error())
	}
	nauDv3A, krzuku, jeWMpOaQtWV :=  /*line :1*/ aRcafgWyH.uaMPM0()
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ inrAIi3wg(nuvOdY1SQ9I5, nauDv3A, krzuku, dEqx7iRZr, gPeNaj, r2c83N, g1sgEzk2pZh)
}

var x9GghUc struct {
	dyil5im	sync.LhBfwn6wa1x
	vJwTbp7_Z	uintptr
	pFC6w1	uintptr
	qmqIZ_3	error
}

func yGHrp3A2EAz() error {
	 /*line :1*/ x9GghUc.dyil5im.Do(func() {
		var bamc83qA3DBr L2L8P9WaNZ
		bamc83qA3DBr, x9GghUc.qmqIZ_3 =  /*line :1*/ L8xQ9MtwVjW(AF_INET, SOCK_DGRAM, IPPROTO_UDP)
		if x9GghUc.qmqIZ_3 != nil {
			return
		}
		defer  /*line :1*/ E5j7kDrdt(bamc83qA3DBr)
		var krzuku uint32
		x9GghUc.qmqIZ_3 =  /*line :1*/ VuX432Ao4_L(bamc83qA3DBr,
			SIO_GET_EXTENSION_FUNCTION_POINTER,
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&U0Kmfq30U)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(U0Kmfq30U)),
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&x9GghUc.pFC6w1)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(x9GghUc.pFC6w1)),
			&krzuku, nil, 0)
		if x9GghUc.qmqIZ_3 != nil {
			return
		}
		x9GghUc.qmqIZ_3 =  /*line :1*/ VuX432Ao4_L(bamc83qA3DBr,
			SIO_GET_EXTENSION_FUNCTION_POINTER,
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&YIL0tl_)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(YIL0tl_)),
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&x9GghUc.vJwTbp7_Z)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(x9GghUc.vJwTbp7_Z)),
			&krzuku, nil, 0)
	})
	return x9GghUc.qmqIZ_3
}

func HYy7t40(nuvOdY1SQ9I5 L2L8P9WaNZ, qWqM5J9x27 *Mh6SqlZ6, a_MrGKcrR uint32, r2c83N *uint32, g1sgEzk2pZh *Ctdynv, fgQM_KRAN9 *byte) error {
	jeWMpOaQtWV :=  /*line :1*/ yGHrp3A2EAz()
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft(x9GghUc.vJwTbp7_Z, 6,  /*line :1*/ uintptr(nuvOdY1SQ9I5),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qWqM5J9x27)),  /*line :1*/ uintptr(a_MrGKcrR),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(r2c83N)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fgQM_KRAN9)))
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return jeWMpOaQtWV
}

func FACJq8(nuvOdY1SQ9I5 L2L8P9WaNZ, qWqM5J9x27 *Mh6SqlZ6, f8IFQaGM3UWX *uint32, g1sgEzk2pZh *Ctdynv, fgQM_KRAN9 *byte) error {
	jeWMpOaQtWV :=  /*line :1*/ yGHrp3A2EAz()
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	i_EJOiAVI, _, j4dSamRp :=  /*line :1*/ syscall.VeAKF8sAwft(x9GghUc.pFC6w1, 5,  /*line :1*/ uintptr(nuvOdY1SQ9I5),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(qWqM5J9x27)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(f8IFQaGM3UWX)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(g1sgEzk2pZh)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(fgQM_KRAN9)), 0)
	if i_EJOiAVI == socket_error {
		jeWMpOaQtWV =  /*line :1*/ rxn0etMCbt1(j4dSamRp)
	}
	return jeWMpOaQtWV
}


type U9ZVPvh struct {
	HiB0AeWSzM	ZPa9KL2
	C4Eb7WX	ZPa9KL2
	IaY_55j0k_4	ZPa9KL2
	RTPZfYC	ZPa9KL2
}

type FNbRqsRX6kti struct {
	ZBjv37 uint32
}

func (mwjEM9EsW FNbRqsRX6kti) Exited() bool	{ return true }

func (mwjEM9EsW FNbRqsRX6kti) ExitStatus() int	{ return  /*line :1*/ int(mwjEM9EsW.ZBjv37) }

func (mwjEM9EsW FNbRqsRX6kti) Signal() ZeAEoa	{ return -1 }

func (mwjEM9EsW FNbRqsRX6kti) CoreDump() bool	{ return false }

func (mwjEM9EsW FNbRqsRX6kti) Stopped() bool	{ return false }

func (mwjEM9EsW FNbRqsRX6kti) Continued() bool	{ return false }

func (mwjEM9EsW FNbRqsRX6kti) StopSignal() ZeAEoa	{ return -1 }

func (mwjEM9EsW FNbRqsRX6kti) Signaled() bool	{ return false }

func (mwjEM9EsW FNbRqsRX6kti) TrapCause() int	{ return -1 }



type YWaK2Afs struct {
	ECuErxxfdoOW	int64
	WPGLV3fH1AX	int64
}

func BZPG6Sd_Z2nW(r7ayG7dk YWaK2Afs) int64 {
	return  /*line :1*/ int64(r7ayG7dk.ECuErxxfdoOW)*1e9 +  /*line :1*/ int64(r7ayG7dk.WPGLV3fH1AX)
}

func WaVDq7EJnI(cHWy1zY4R5z int64) (r7ayG7dk YWaK2Afs) {
	r7ayG7dk.ECuErxxfdoOW = cHWy1zY4R5z / 1e9
	r7ayG7dk.WPGLV3fH1AX = cHWy1zY4R5z % 1e9
	return
}

func H6_Rt7s6mKlH(nuvOdY1SQ9I5 L2L8P9WaNZ) (n_3tmXhyl L2L8P9WaNZ, aRcafgWyH VPXahxqr5, jeWMpOaQtWV error) {
	return 0, nil, syscall.EWINDOWS
}

func FaC0y4gDQ_j(nuvOdY1SQ9I5 L2L8P9WaNZ, hD4wNgEB []byte, a_MrGKcrR int) (krzuku int, wPbtqIJl VPXahxqr5, jeWMpOaQtWV error) {
	var u1fJXFk I6IZfmg
	cTpvqfM :=  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(u1fJXFk))
	rAP5NTdoA, jeWMpOaQtWV :=  /*line :1*/ nE41zAHtIi(nuvOdY1SQ9I5, hD4wNgEB,  /*line :1*/ int32(a_MrGKcrR), &u1fJXFk, &cTpvqfM)
	krzuku =  /*line :1*/ int(rAP5NTdoA)
	if jeWMpOaQtWV != nil {
		return
	}
	wPbtqIJl, jeWMpOaQtWV =  /*line :1*/ u1fJXFk.Sockaddr()
	return
}

func SU9a9qFTVHb(nuvOdY1SQ9I5 L2L8P9WaNZ, hD4wNgEB []byte, a_MrGKcrR int, xtbcgl1U7 VPXahxqr5) (jeWMpOaQtWV error) {
	nauDv3A, cTpvqfM, jeWMpOaQtWV :=  /*line :1*/ xtbcgl1U7.uaMPM0()
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}
	return  /*line :1*/ mdvojh1n(nuvOdY1SQ9I5, hD4wNgEB,  /*line :1*/ int32(a_MrGKcrR), nauDv3A, cTpvqfM)
}

func Brwi8z(nuvOdY1SQ9I5 L2L8P9WaNZ, bR_S51, qUM7_22hM int, fmaieegZe2E_ *F2i2x7peITi) (jeWMpOaQtWV error) {
	return syscall.EWINDOWS
}

type GTIT19PprZR struct {
	V6pIuhwCu	int32
	M0TrEat	int32
}

type dksU8_S struct {
	L8P8ult	uint16
	Z95uALy	uint16
}

type AV6su9G struct {
	P8O2NF0WA	[4]byte	
	Nompo3TK	[4]byte	
}

type XsujHFenhmb struct {
	RmQSpqMVN	[16]byte	
	GOMCVrEx1k	uint32
}

func V7P5P0TEgDFv(nuvOdY1SQ9I5 L2L8P9WaNZ, bR_S51, qUM7_22hM int) (int, error) {
	n8CI2rJ :=  /*line :1*/ int32(0)
	cTpvqfM :=  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(n8CI2rJ))
	jeWMpOaQtWV :=  /*line :1*/ U7Z_4P(nuvOdY1SQ9I5,  /*line :1*/ int32(bR_S51),  /*line :1*/ int32(qUM7_22hM), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&n8CI2rJ)), &cTpvqfM)
	return  /*line :1*/ int(n8CI2rJ), jeWMpOaQtWV
}

func GkVJukKdCr0O(nuvOdY1SQ9I5 L2L8P9WaNZ, bR_S51, qUM7_22hM int, cTpvqfM *GTIT19PprZR) (jeWMpOaQtWV error) {
	f3v0ShGQ := dksU8_S{L8P8ult:  /*line :1*/ uint16(cTpvqfM.V6pIuhwCu), Z95uALy:  /*line :1*/ uint16(cTpvqfM.M0TrEat)}
	return  /*line :1*/ UfWy_ocZ(nuvOdY1SQ9I5,  /*line :1*/ int32(bR_S51),  /*line :1*/ int32(qUM7_22hM), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&f3v0ShGQ)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(f3v0ShGQ)))
}

func ZYk0Lpk(nuvOdY1SQ9I5 L2L8P9WaNZ, bR_S51, qUM7_22hM int, wvZBcfh [4]byte) (jeWMpOaQtWV error) {
	return  /*line :1*/ UfWy_ocZ(nuvOdY1SQ9I5,  /*line :1*/ int32(bR_S51),  /*line :1*/ int32(qUM7_22hM), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&wvZBcfh[0])), 4)
}
func APdr2tT(nuvOdY1SQ9I5 L2L8P9WaNZ, bR_S51, qUM7_22hM int, w85K5hs95 *AV6su9G) (jeWMpOaQtWV error) {
	return  /*line :1*/ UfWy_ocZ(nuvOdY1SQ9I5,  /*line :1*/ int32(bR_S51),  /*line :1*/ int32(qUM7_22hM), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(w85K5hs95)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*w85K5hs95)))
}
func BEWUaDZIat(nuvOdY1SQ9I5 L2L8P9WaNZ, bR_S51, qUM7_22hM int, w85K5hs95 *XsujHFenhmb) (jeWMpOaQtWV error) {
	return syscall.EWINDOWS
}

func BjHvksKKlhaM(kJf2QPox7Wdz []uint32, pwkETEB8SLrs *uint32) error {
	
	
	var hD4wNgEB *uint32
	if  /*line :1*/ len(kJf2QPox7Wdz) > 0 {
		hD4wNgEB = &kJf2QPox7Wdz[0]
	}
	nFtsmqHC :=  /*line :1*/ uint32( /*line :1*/ len(kJf2QPox7Wdz) * 4)
	return  /*line :1*/ xJxpOac(hD4wNgEB, nFtsmqHC, pwkETEB8SLrs)
}

func CvEBSTY() (us2BXoyQ_y int)	{ return  /*line :1*/ int( /*line :1*/ S31cyZsdeH()) }

func N_T_YbGz2kc(gYwswmeUjG *uint16, iIzhNC *O1XfSfprUf) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV error) {
	
	
	
	
	
	
	
	
	var hou6IV4 qW7t5nXz
	iOvctVD26lA, jeWMpOaQtWV =  /*line :1*/ it9q8T(gYwswmeUjG, &hou6IV4)
	if jeWMpOaQtWV == nil {
		 /*line :1*/ xxmBpbV1iJZ(iIzhNC, &hou6IV4)
	}
	return
}

func VQzTMK(iOvctVD26lA L2L8P9WaNZ, iIzhNC *O1XfSfprUf) (jeWMpOaQtWV error) {
	var hou6IV4 qW7t5nXz
	jeWMpOaQtWV =  /*line :1*/ eIUeXFp_Eqa(iOvctVD26lA, &hou6IV4)
	if jeWMpOaQtWV == nil {
		 /*line :1*/ xxmBpbV1iJZ(iIzhNC, &hou6IV4)
	}
	return
}

func tripSkOpm(us2BXoyQ_y int) (*RTsQKj, error) {
	uGiEDsx, jeWMpOaQtWV :=  /*line :1*/ L8wfFX3s(TH32CS_SNAPPROCESS, 0)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	defer  /*line :1*/ E5j7kDrdt(uGiEDsx)
	var yj2P3JPS5i RTsQKj
	yj2P3JPS5i.NgWYrRzCdYqS =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(yj2P3JPS5i))
	if jeWMpOaQtWV =  /*line :1*/ Mdm1wP1(uGiEDsx, &yj2P3JPS5i); jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	for {
		if yj2P3JPS5i.QU8DfIi ==  /*line :1*/ uint32(us2BXoyQ_y) {
			return &yj2P3JPS5i, nil
		}
		jeWMpOaQtWV =  /*line :1*/ RsaLWPf2NmF(uGiEDsx, &yj2P3JPS5i)
		if jeWMpOaQtWV != nil {
			return nil, jeWMpOaQtWV
		}
	}
}

func G2Yv34k8C() (m4OLGU9zenx int) {
	jLOlRevla, jeWMpOaQtWV :=  /*line :1*/ tripSkOpm( /*line :1*/ CvEBSTY())
	if jeWMpOaQtWV != nil {
		return -1
	}
	return  /*line :1*/ int(jLOlRevla.TvE7hf)
}


func R3wI_MP(nuvOdY1SQ9I5 L2L8P9WaNZ) (jeWMpOaQtWV error)	{ return syscall.EWINDOWS }
func R6WopA(g3QX6O1Z, vqYv72aSS string) (jeWMpOaQtWV error)	{ return syscall.EWINDOWS }
func H9cHnTf5haXj(bZKbGTxeit, s5A3Fb string) (jeWMpOaQtWV error)	{ return syscall.EWINDOWS }

func EIU0835D7eMJ(nuvOdY1SQ9I5 L2L8P9WaNZ, yAZay4eEB0M uint32) (jeWMpOaQtWV error) {
	return syscall.EWINDOWS
}
func L8sVZaEyYXK(bZKbGTxeit string, n1jPcIv79N int, cH8B4cc_a int) (jeWMpOaQtWV error) {
	return syscall.EWINDOWS
}
func ZDH0uGuPklb(bZKbGTxeit string, n1jPcIv79N int, cH8B4cc_a int) (jeWMpOaQtWV error) {
	return syscall.EWINDOWS
}
func VA8KOZW(nuvOdY1SQ9I5 L2L8P9WaNZ, n1jPcIv79N int, cH8B4cc_a int) (jeWMpOaQtWV error) {
	return syscall.EWINDOWS
}

func GxdOBDwZVcL() (n1jPcIv79N int)	{ return -1 }
func KoVMicG() (nIiLxuWra52o int)	{ return -1 }
func A7Nx8OL() (cH8B4cc_a int)	{ return -1 }
func YvWvG7Rj() (tM3bgUUY int)	{ return -1 }
func LVUbHO2WK() (xOxMB5 []int, jeWMpOaQtWV error)	{ return nil, syscall.EWINDOWS }

type ZeAEoa int

func (bamc83qA3DBr ZeAEoa) Signal()	{}

func (bamc83qA3DBr ZeAEoa) String() string {
	if 0 <= bamc83qA3DBr &&  /*line :1*/ int(bamc83qA3DBr) <  /*line :1*/ len(wkWGHLHFwGFC) {
		xukLmcNaR := wkWGHLHFwGFC[bamc83qA3DBr]
		if xukLmcNaR != "" {
			return xukLmcNaR
		}
	}
	return "signal " +  /*line :1*/ dMi2GE( /*line :1*/ int(bamc83qA3DBr))
}

func SzOMyJbI() error {
	return  /*line :1*/ sax1PV7.Find()
}


func E2pE1u(bZKbGTxeit string, etRYtnVni []byte) (krzuku int, jeWMpOaQtWV error) {
	nuvOdY1SQ9I5, jeWMpOaQtWV :=  /*line :1*/ LNhQWvTnm6( /*line :1*/ RRlduq3XI(bZKbGTxeit), GENERIC_READ, 0, nil, OPEN_EXISTING,
		FILE_FLAG_OPEN_REPARSE_POINT|FILE_FLAG_BACKUP_SEMANTICS, 0)
	if jeWMpOaQtWV != nil {
		return -1, jeWMpOaQtWV
	}
	defer  /*line :1*/ E5j7kDrdt(nuvOdY1SQ9I5)

	j6fxSUZoOcg :=  /*line :1*/ make([]byte, MAXIMUM_REPARSE_DATA_BUFFER_SIZE)
	var pwkETEB8SLrs uint32
	jeWMpOaQtWV =  /*line :1*/ HjWTjVjfKi(nuvOdY1SQ9I5, FSCTL_GET_REPARSE_POINT, nil, 0, &j6fxSUZoOcg[0],  /*line :1*/ uint32( /*line :1*/ len(j6fxSUZoOcg)), &pwkETEB8SLrs, nil)
	if jeWMpOaQtWV != nil {
		return -1, jeWMpOaQtWV
	}

	bVYjJa_o := (* /*line :1*/ zs4HWqZrE)( /*line :1*/ unsafe.Pointer(&j6fxSUZoOcg[0]))
	var bamc83qA3DBr string
	switch bVYjJa_o.Ct81WfAEIFFo {
	case IO_REPARSE_TAG_SYMLINK:
		iIzhNC := (* /*line :1*/ bQJsnAW2X)( /*line :1*/ unsafe.Pointer(&bVYjJa_o.ofwNwEb5aYf))
		hD4wNgEB := (*[0xffff] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&iIzhNC.TKCfxANSN[0]))
		bamc83qA3DBr =  /*line :1*/ OXNanQ8Uj(hD4wNgEB[iIzhNC.EXW941kPU/2 : (iIzhNC.MoWBo1VLaY-iIzhNC.EXW941kPU)/2])
	case IO_REPARSE_TAG_MOUNT_POINT:
		iIzhNC := (* /*line :1*/ _ASWCv22rUTP)( /*line :1*/ unsafe.Pointer(&bVYjJa_o.ofwNwEb5aYf))
		hD4wNgEB := (*[0xffff] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&iIzhNC.HbIKNUMDT4V[0]))
		bamc83qA3DBr =  /*line :1*/ OXNanQ8Uj(hD4wNgEB[iIzhNC.DLN9jIh4l/2 : (iIzhNC.SCwRyx_Oh-iIzhNC.DLN9jIh4l)/2])
	default:

		return -1, syscall.ENOENT
	}
	krzuku =  /*line :1*/ copy(etRYtnVni, [] /*line :1*/ byte(bamc83qA3DBr))

	return krzuku, nil
}



func GCCvj7uW(xukLmcNaR string) (Kw2qafRFaiLg, error) {
	cM24Fx0MIn := Kw2qafRFaiLg{}
	jyB65N, jeWMpOaQtWV :=  /*line :1*/ syscall.GcOmFfsibES(xukLmcNaR)
	if jeWMpOaQtWV != nil {
		return cM24Fx0MIn, jeWMpOaQtWV
	}
	jeWMpOaQtWV =  /*line :1*/ aXy1KcoG5(jyB65N, &cM24Fx0MIn)
	if jeWMpOaQtWV != nil {
		return cM24Fx0MIn, jeWMpOaQtWV
	}
	return cM24Fx0MIn, nil
}


func ACZTl0E() (Kw2qafRFaiLg, error) {
	cM24Fx0MIn := Kw2qafRFaiLg{}
	jeWMpOaQtWV :=  /*line :1*/ ed9PsDqDFOWJ(&cM24Fx0MIn)
	if jeWMpOaQtWV != nil {
		return cM24Fx0MIn, jeWMpOaQtWV
	}
	return cM24Fx0MIn, nil
}



func (cM24Fx0MIn Kw2qafRFaiLg) String() string {
	var xukLmcNaR [100]uint16
	x7J8HIt :=  /*line :1*/ hlENtQ7V(&cM24Fx0MIn, &xukLmcNaR[0],  /*line :1*/ int32( /*line :1*/ len(xukLmcNaR)))
	if x7J8HIt <= 1 {
		return ""
	}
	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(xukLmcNaR[:x7J8HIt-1]))
}



func Tc5MsWeGm(v9mfU1ymTiUV *S06JEslqEe, a_MrGKcrR uint32) (string, error) {
	return  /*line :1*/ TaSPPoJMlh(0).KnownFolderPath(v9mfU1ymTiUV, a_MrGKcrR)
}



func (g9JYFzvU35 TaSPPoJMlh) KnownFolderPath(v9mfU1ymTiUV *S06JEslqEe, a_MrGKcrR uint32) (string, error) {
	var hD4wNgEB *uint16
	jeWMpOaQtWV :=  /*line :1*/ a0OuuH(v9mfU1ymTiUV, a_MrGKcrR, g9JYFzvU35, &hD4wNgEB)
	if jeWMpOaQtWV != nil {
		return "", jeWMpOaQtWV
	}
	defer  /*line :1*/ U9bXbY99( /*line :1*/ unsafe.Pointer(hD4wNgEB))
	return  /*line :1*/ M_Sea9j(hD4wNgEB), nil
}



func YlZWYO1cWMt() *GDqcSn {
	a7Z1LWSG := &GDqcSn{}
	a7Z1LWSG.sSmIEd =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*a7Z1LWSG))

	_ =  /*line :1*/ yUGGGBD(a7Z1LWSG)
	return a7Z1LWSG
}



func EiITmOW() (n2tJaJ, tBRYiG, idoDATkQaBX uint32) {
	 /*line :1*/ n3PGfd(&n2tJaJ, &tBRYiG, &idoDATkQaBX)
	idoDATkQaBX &= 0xffff
	return
}


func Q9oL_AZHtO1e(a_MrGKcrR uint32) ([]string, error) {
	return  /*line :1*/ dszHaDk(a_MrGKcrR, h4WTQ109)
}


func Uq5ggbzp(a_MrGKcrR uint32) ([]string, error) {
	return  /*line :1*/ dszHaDk(a_MrGKcrR, sLLZssX)
}


func BxG6Ls(a_MrGKcrR uint32) ([]string, error) {
	return  /*line :1*/ dszHaDk(a_MrGKcrR, czYdae)
}


func ZpX0QztN(a_MrGKcrR uint32) ([]string, error) {
	return  /*line :1*/ dszHaDk(a_MrGKcrR, aBvdnLT)
}

func dszHaDk(a_MrGKcrR uint32, wC_6WjtKcvD func(a_MrGKcrR uint32, rzHbAFn *uint32, etRYtnVni *uint16, kvo8765 *uint32) error) ([]string, error) {
	nFtsmqHC :=  /*line :1*/ uint32(128)
	for {
		var rzHbAFn uint32
		etRYtnVni :=  /*line :1*/ make([]uint16, nFtsmqHC)
		jeWMpOaQtWV :=  /*line :1*/ wC_6WjtKcvD(a_MrGKcrR, &rzHbAFn, &etRYtnVni[0], &nFtsmqHC)
		if jeWMpOaQtWV == ERROR_INSUFFICIENT_BUFFER {
			continue
		}
		if jeWMpOaQtWV != nil {
			return nil, jeWMpOaQtWV
		}
		etRYtnVni = etRYtnVni[:nFtsmqHC]
		if rzHbAFn == 0 ||  /*line :1*/ len(etRYtnVni) == 0 {
			return []string{}, nil
		}
		if etRYtnVni[ /*line :1*/ len(etRYtnVni)-1] == 0 {
			etRYtnVni = etRYtnVni[: /*line :1*/ len(etRYtnVni)-1]
		}
		zO7UqyT :=  /*line :1*/ make([]string, 0, rzHbAFn)
		wPbtqIJl := 0
		for rRGfxgPH8Kq, jpk9X7z := range etRYtnVni {
			if jpk9X7z == 0 {
				zO7UqyT =  /*line :1*/ append(zO7UqyT,  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(etRYtnVni[wPbtqIJl:rRGfxgPH8Kq])))
				wPbtqIJl = rRGfxgPH8Kq + 1
			}
		}
		return zO7UqyT, nil
	}
}

func PNWng564bK26(rprJNy L2L8P9WaNZ, pYy71tl ZPW7DmIDA) error {
	return  /*line :1*/ bdqQnj3n(rprJNy, *((* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(&pYy71tl))))
}

func GVUm1M2aj(hnFFk_ *KqXK1jwYl5VP) error {
	 /*line :1*/ yYlSfA5GPA(hnFFk_)
	return nil
}

func (bamc83qA3DBr T6LY983) Errno() syscall.O9Mga3 {
	return  /*line :1*/ b4ap1rlCaX(bamc83qA3DBr)
}

func eTwXgn4r9(yiCfSvnh, vdbEWz1aC uint16) uint32	{ return  /*line :1*/ uint32(vdbEWz1aC)<<10 |  /*line :1*/ uint32(yiCfSvnh) }

func (bamc83qA3DBr T6LY983) Error() string {
	hJuLpmVFi :=  /*line :1*/ make([]uint16, 300)
	krzuku, jeWMpOaQtWV :=  /*line :1*/ FLUYLrlNYT(FORMAT_MESSAGE_FROM_SYSTEM|FORMAT_MESSAGE_FROM_HMODULE|FORMAT_MESSAGE_ARGUMENT_ARRAY,  /*line :1*/ e6pKCXfI.Handle(),  /*line :1*/ uint32(bamc83qA3DBr),  /*line :1*/ eTwXgn4r9(LANG_ENGLISH, SUBLANG_ENGLISH_US), hJuLpmVFi, nil)
	if jeWMpOaQtWV != nil {
		return  /*line :1*/ fmt.IBsS3Oc("NTSTATUS 0x%08x",  /*line :1*/ uint32(bamc83qA3DBr))
	}

	for ; krzuku > 0 && (hJuLpmVFi[krzuku-1] == '\n' || hJuLpmVFi[krzuku-1] == '\r'); krzuku-- {
	}
	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(hJuLpmVFi[:krzuku]))
}





func NwkLRFmJN(bamc83qA3DBr string) (*QxJ3Xtd, error) {
	var sMdg1A QxJ3Xtd
	gjdPJbrWa5n, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(bamc83qA3DBr)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	 /*line :1*/ OIJooIiENmZ(&sMdg1A, gjdPJbrWa5n)
	return &sMdg1A, nil
}


func (bamc83qA3DBr *QxJ3Xtd) Slice() []uint16 {
	oWF1gbrv :=  /*line :1*/ unsafe.Slice(bamc83qA3DBr.FKtTgm5Cq, bamc83qA3DBr.EwdOcHCGn8N)
	return oWF1gbrv[:bamc83qA3DBr.AXjXPv6B]
}

func (bamc83qA3DBr *QxJ3Xtd) String() string {
	return  /*line :1*/ OXNanQ8Uj( /*line :1*/ bamc83qA3DBr.Slice())
}





func LYMO2bsm89E5(bamc83qA3DBr string) (*DxRO_Vrr7, error) {
	var iuKL9EzP DxRO_Vrr7
	jB4tO1, jeWMpOaQtWV :=  /*line :1*/ UgPyZIRW(bamc83qA3DBr)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	 /*line :1*/ BmTG_IqbTIoV(&iuKL9EzP, jB4tO1)
	return &iuKL9EzP, nil
}


func (bamc83qA3DBr *DxRO_Vrr7) Slice() []byte {
	oWF1gbrv :=  /*line :1*/ unsafe.Slice(bamc83qA3DBr.I15m2cqwvMu, bamc83qA3DBr.QfohTf)
	return oWF1gbrv[:bamc83qA3DBr.U2PfVRwS]
}

func (bamc83qA3DBr *DxRO_Vrr7) String() string {
	return  /*line :1*/ NWmrwy( /*line :1*/ bamc83qA3DBr.Slice())
}


func AjSlrVi(tiLd7g0 L2L8P9WaNZ, gYwswmeUjG, i0LVb2 BY7rRyd) (L2L8P9WaNZ, error) {
	var wjd9bchzGH8, nqY3uhBu uintptr
	var ynxDH59ypANn, hI0zfx *uint16
	var jeWMpOaQtWV error
	jCHwZddYt439 := func(rRGfxgPH8Kq interface{}, gBYFQqk **uint16) (uintptr, error) {
		switch n8CI2rJ := rRGfxgPH8Kq.(type) {
		case string:
			*gBYFQqk, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(n8CI2rJ)
			if jeWMpOaQtWV != nil {
				return 0, jeWMpOaQtWV
			}
			return  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(*gBYFQqk)), nil
		case CA6MMo3djsM:
			return  /*line :1*/ uintptr(n8CI2rJ), nil
		}
		return 0,  /*line :1*/ errorspkg.Su6g6hRoi9X("parameter must be a ResourceID or a string")
	}
	wjd9bchzGH8, jeWMpOaQtWV =  /*line :1*/ jCHwZddYt439(gYwswmeUjG, &ynxDH59ypANn)
	if jeWMpOaQtWV != nil {
		return 0, jeWMpOaQtWV
	}
	nqY3uhBu, jeWMpOaQtWV =  /*line :1*/ jCHwZddYt439(i0LVb2, &hI0zfx)
	if jeWMpOaQtWV != nil {
		return 0, jeWMpOaQtWV
	}
	dEaAptPTS5SC, jeWMpOaQtWV :=  /*line :1*/ uFpH3ANUM(tiLd7g0, wjd9bchzGH8, nqY3uhBu)
	 /*line :1*/ runtime.KeepAlive(ynxDH59ypANn)
	 /*line :1*/ runtime.KeepAlive(hI0zfx)
	return dEaAptPTS5SC, jeWMpOaQtWV
}

func Pe2oEY7z(tiLd7g0, dEaAptPTS5SC L2L8P9WaNZ) (iIzhNC []byte, jeWMpOaQtWV error) {
	nFtsmqHC, jeWMpOaQtWV :=  /*line :1*/ BDczRK(tiLd7g0, dEaAptPTS5SC)
	if jeWMpOaQtWV != nil {
		return
	}
	crYOfg3, jeWMpOaQtWV :=  /*line :1*/ SJwhtYsfPML(tiLd7g0, dEaAptPTS5SC)
	if jeWMpOaQtWV != nil {
		return
	}
	nauDv3A, jeWMpOaQtWV :=  /*line :1*/ JvJtQTH9y(crYOfg3)
	if jeWMpOaQtWV != nil {
		return
	}
	iIzhNC =  /*line :1*/ unsafe.Slice((* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(nauDv3A)), nFtsmqHC)
	return
}


type JLyJCM uint64



func (hJuLpmVFi JLyJCM) Valid() bool {
	return (hJuLpmVFi & 1) == 1
}


func (hJuLpmVFi JLyJCM) ShareCount() uint64 {
	return  /*line :1*/ hJuLpmVFi.ty6muunj9(1, 3)
}



func (hJuLpmVFi JLyJCM) Win32Protection() uint64 {
	return  /*line :1*/ hJuLpmVFi.ty6muunj9(4, 11)
}



func (hJuLpmVFi JLyJCM) Shared() bool {
	return (hJuLpmVFi & (1 << 15)) == 1
}


func (hJuLpmVFi JLyJCM) Node() uint64 {
	return  /*line :1*/ hJuLpmVFi.ty6muunj9(16, 6)
}



func (hJuLpmVFi JLyJCM) Locked() bool {
	return (hJuLpmVFi & (1 << 22)) == 1
}



func (hJuLpmVFi JLyJCM) LargePage() bool {
	return (hJuLpmVFi & (1 << 23)) == 1
}



func (hJuLpmVFi JLyJCM) Bad() bool {
	return (hJuLpmVFi & (1 << 31)) == 1
}


func (hJuLpmVFi JLyJCM) ty6muunj9(fe9x0tyLFZR5, fORnBp int) uint64 {
	var dLarUKMqaxh JLyJCM
	for ppFY34qtv2A := fe9x0tyLFZR5; ppFY34qtv2A < fe9x0tyLFZR5+fORnBp; ppFY34qtv2A++ {
		dLarUKMqaxh |= (1 << ppFY34qtv2A)
	}

	_4MXL6a_ := hJuLpmVFi & dLarUKMqaxh
	return  /*line :1*/ uint64(_4MXL6a_ >> fe9x0tyLFZR5)
}


type JhaIP5c struct {
	
	TRxaxl4p	Kwlsdbsx46mT
	
	Yv3Ua8	JLyJCM
}


func L9N5MLOO(nFtsmqHC ZPW7DmIDA, l8eN3VDi L2L8P9WaNZ, icEgEFtX_G0 L2L8P9WaNZ, a_MrGKcrR uint32, zaIN6WS *L2L8P9WaNZ) error {

	return  /*line :1*/ liWXWGwW(*((* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(&nFtsmqHC))), l8eN3VDi, icEgEFtX_G0, a_MrGKcrR, zaIN6WS)
}


func GdACSL6oLz(zaIN6WS L2L8P9WaNZ, nFtsmqHC ZPW7DmIDA) error {

	return  /*line :1*/ flaybp5yd1(zaIN6WS, *((* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(&nFtsmqHC))))
}
