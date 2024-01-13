//line :1
package bUKeamF

import (
	errorspkg "iAHoxjmM"
	"internal/bytealg"
	itoa "JkjxVS"
	oserror "K7w4V9ecdx2"
	race "V0p_iv"
	"runtime"
	sync "sync"
	"unsafe"
)

type SRlvVjrYa uintptr

const InvalidHandle = ^ /*line :1*/ SRlvVjrYa(0)






func KtaNWT(wzPhJhIFoI string) []uint16 {
	n0ogBVz, gOCcQzbcC :=  /*line :1*/ ZxIMCrnjQ5U4(wzPhJhIFoI)
	if gOCcQzbcC != nil {
		 /*line :1*/ panic("syscall: string with NUL passed to StringToUTF16")
	}
	return n0ogBVz
}





func ZxIMCrnjQ5U4(wzPhJhIFoI string) ([]uint16, error) {
	if  /*line :1*/ bytealg.IndexByteString(wzPhJhIFoI, 0) != -1 {
		return nil, EINVAL
	}

	eun1Jlud5A :=  /*line :1*/ make([]uint16, 0,  /*line :1*/ len(wzPhJhIFoI)+1)
	eun1Jlud5A =  /*line :1*/ jsi_qEcoyU9L(wzPhJhIFoI, eun1Jlud5A)
	return  /*line :1*/ append(eun1Jlud5A, 0), nil
}




func AODVXp8NM3sd(wzPhJhIFoI []uint16) string {
	_aJWKcxF := 0
	for hA99q3EOK, b0abCyNZ := range wzPhJhIFoI {
		if b0abCyNZ == 0 {
			wzPhJhIFoI = wzPhJhIFoI[0:hA99q3EOK]
			break
		}
		switch {
		case b0abCyNZ <= rune1Max:
			_aJWKcxF += 1
		case b0abCyNZ <= rune2Max:
			_aJWKcxF += 2
		default:

			_aJWKcxF += 3
		}
	}
	eun1Jlud5A :=  /*line :1*/ mtdGDMk8v(wzPhJhIFoI,  /*line :1*/ make([]byte, 0, _aJWKcxF))
	return  /*line :1*/ unsafe.String( /*line :1*/ unsafe.SliceData(eun1Jlud5A),  /*line :1*/ len(eun1Jlud5A))
}



func w2QjpJKaV(rd6leevODT *uint16) string {
	if rd6leevODT == nil {
		return ""
	}
	iMDTBTj :=  /*line :1*/ unsafe.Pointer(rd6leevODT)
	m5Tq_PL7 := 0
	for *(* /*line :1*/ uint16)(iMDTBTj) != 0 {
		iMDTBTj =  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(iMDTBTj) +  /*line :1*/ unsafe.Sizeof(*rd6leevODT))
		m5Tq_PL7++
	}
	return  /*line :1*/ AODVXp8NM3sd( /*line :1*/ unsafe.Slice(rd6leevODT, m5Tq_PL7))
}







func EtPVNA(wzPhJhIFoI string) *uint16	{ return & /*line :1*/ KtaNWT(wzPhJhIFoI)[0] }





func GcOmFfsibES(wzPhJhIFoI string) (*uint16, error) {
	n0ogBVz, gOCcQzbcC :=  /*line :1*/ ZxIMCrnjQ5U4(wzPhJhIFoI)
	if gOCcQzbcC != nil {
		return nil, gOCcQzbcC
	}
	return &n0ogBVz[0], nil
}








type O9Mga3 uintptr

func gLEjStHSQo(cmFFDEg8, t_1weRk uint16) uint32	{ return  /*line :1*/ uint32(t_1weRk)<<10 |  /*line :1*/ uint32(cmFFDEg8) }





func Lw1VZPgjxcT4(uADzMh0kZ7hL uint32, mEtGX6Jm uint32, fpKe4WGIx uint32, gLEjStHSQo uint32, eun1Jlud5A []uint16, dqAabTiua *byte) (m5Tq_PL7 uint32, gOCcQzbcC error) {
	return  /*line :1*/ _MtwugS(uADzMh0kZ7hL,  /*line :1*/ uintptr(mEtGX6Jm), fpKe4WGIx, gLEjStHSQo, eun1Jlud5A, dqAabTiua)
}

func (wAwFLr4AGHg O9Mga3) Error() string {

	hA2ajqakQmj :=  /*line :1*/ int(wAwFLr4AGHg - APPLICATION_ERROR)
	if 0 <= hA2ajqakQmj && hA2ajqakQmj <  /*line :1*/ len(ikgFaHUF) {
		return ikgFaHUF[hA2ajqakQmj]
	}
	
	var uADzMh0kZ7hL uint32 = FORMAT_MESSAGE_FROM_SYSTEM | FORMAT_MESSAGE_ARGUMENT_ARRAY | FORMAT_MESSAGE_IGNORE_INSERTS
	tHPwhuTh :=  /*line :1*/ make([]uint16, 300)
	m5Tq_PL7, gOCcQzbcC :=  /*line :1*/ _MtwugS(uADzMh0kZ7hL, 0,  /*line :1*/ uint32(wAwFLr4AGHg),  /*line :1*/ gLEjStHSQo(LANG_ENGLISH, SUBLANG_ENGLISH_US), tHPwhuTh, nil)
	if gOCcQzbcC != nil {
		m5Tq_PL7, gOCcQzbcC =  /*line :1*/ _MtwugS(uADzMh0kZ7hL, 0,  /*line :1*/ uint32(wAwFLr4AGHg), 0, tHPwhuTh, nil)
		if gOCcQzbcC != nil {
			return "winapi error #" +  /*line :1*/ itoa.V25ba9G5( /*line :1*/ int(wAwFLr4AGHg))
		}
	}

	for ; m5Tq_PL7 > 0 && (tHPwhuTh[m5Tq_PL7-1] == '\n' || tHPwhuTh[m5Tq_PL7-1] == '\r'); m5Tq_PL7-- {
	}
	return  /*line :1*/ AODVXp8NM3sd(tHPwhuTh[:m5Tq_PL7])
}

const (
	_ERROR_NOT_ENOUGH_MEMORY	=  /*line :1*/ O9Mga3(8)
	_ERROR_NOT_SUPPORTED	=  /*line :1*/ O9Mga3(50)
	_ERROR_BAD_NETPATH	=  /*line :1*/ O9Mga3(53)
	_ERROR_CALL_NOT_IMPLEMENTED	=  /*line :1*/ O9Mga3(120)
)

func (wAwFLr4AGHg O9Mga3) Is(uZWwO3ukEeJN error) bool {
	switch uZWwO3ukEeJN {
	case oserror.HCknuLDV1zb:
		return wAwFLr4AGHg == ERROR_ACCESS_DENIED ||
			wAwFLr4AGHg == EACCES ||
			wAwFLr4AGHg == EPERM
	case oserror.IyvfTZaiEa:
		return wAwFLr4AGHg == ERROR_ALREADY_EXISTS ||
			wAwFLr4AGHg == ERROR_DIR_NOT_EMPTY ||
			wAwFLr4AGHg == ERROR_FILE_EXISTS ||
			wAwFLr4AGHg == EEXIST ||
			wAwFLr4AGHg == ENOTEMPTY
	case oserror.RKRo3rPG5Gn:
		return wAwFLr4AGHg == ERROR_FILE_NOT_FOUND ||
			wAwFLr4AGHg == _ERROR_BAD_NETPATH ||
			wAwFLr4AGHg == ERROR_PATH_NOT_FOUND ||
			wAwFLr4AGHg == ENOENT
	case errorspkg.InZTDm:
		return wAwFLr4AGHg == _ERROR_NOT_SUPPORTED ||
			wAwFLr4AGHg == _ERROR_CALL_NOT_IMPLEMENTED ||
			wAwFLr4AGHg == ENOSYS ||
			wAwFLr4AGHg == ENOTSUP ||
			wAwFLr4AGHg == EOPNOTSUPP ||
			wAwFLr4AGHg == EWINDOWS
	}
	return false
}

func (wAwFLr4AGHg O9Mga3) Temporary() bool {
	return wAwFLr4AGHg == EINTR || wAwFLr4AGHg == EMFILE ||  /*line :1*/ wAwFLr4AGHg.Timeout()
}

func (wAwFLr4AGHg O9Mga3) Timeout() bool {
	return wAwFLr4AGHg == EAGAIN || wAwFLr4AGHg == EWOULDBLOCK || wAwFLr4AGHg == ETIMEDOUT
}


func v0aq4Uyljv(uIoB0opUef any, huNccD bool) uintptr







func VSPtNpjXc4(uIoB0opUef any) uintptr {
	return  /*line :1*/ v0aq4Uyljv(uIoB0opUef, true)
}







func MLY71IBA(uIoB0opUef any) uintptr {
	return  /*line :1*/ v0aq4Uyljv(uIoB0opUef, false)
}

func e0GOege_rv() *DDwuM6RpYja {
	var zC0xxjnQnl DDwuM6RpYja
	zC0xxjnQnl.OIMmqEzG2Yx0 =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(zC0xxjnQnl))
	zC0xxjnQnl.EDnb4yqd = 1
	return &zC0xxjnQnl
}

func Ek485AtskVVo(ro5NCL9 string, ediyaFOkNQmk int, gOrivca8ecRv uint32) (dBn476FVCy2C SRlvVjrYa, gOCcQzbcC error) {
	if  /*line :1*/ len(ro5NCL9) == 0 {
		return InvalidHandle, ERROR_FILE_NOT_FOUND
	}
	zpdigBaban, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if gOCcQzbcC != nil {
		return InvalidHandle, gOCcQzbcC
	}
	var wzWkRJs uint32
	switch ediyaFOkNQmk & (O_RDONLY | O_WRONLY | O_RDWR) {
	case O_RDONLY:
		wzWkRJs = GENERIC_READ
	case O_WRONLY:
		wzWkRJs = GENERIC_WRITE
	case O_RDWR:
		wzWkRJs = GENERIC_READ | GENERIC_WRITE
	}
	if ediyaFOkNQmk&O_CREAT != 0 {
		wzWkRJs |= GENERIC_WRITE
	}
	if ediyaFOkNQmk&O_APPEND != 0 {
		wzWkRJs &^= GENERIC_WRITE
		wzWkRJs |= FILE_APPEND_DATA
	}
	pd1xDFCP03 :=  /*line :1*/ uint32(FILE_SHARE_READ | FILE_SHARE_WRITE)
	var zC0xxjnQnl *DDwuM6RpYja
	if ediyaFOkNQmk&O_CLOEXEC == 0 {
		zC0xxjnQnl =  /*line :1*/ e0GOege_rv()
	}
	var ygINFYaFDGa8 uint32
	switch {
	case ediyaFOkNQmk&(O_CREAT|O_EXCL) == (O_CREAT | O_EXCL):
		ygINFYaFDGa8 = CREATE_NEW
	case ediyaFOkNQmk&(O_CREAT|O_TRUNC) == (O_CREAT | O_TRUNC):
		ygINFYaFDGa8 = CREATE_ALWAYS
	case ediyaFOkNQmk&O_CREAT == O_CREAT:
		ygINFYaFDGa8 = OPEN_ALWAYS
	case ediyaFOkNQmk&O_TRUNC == O_TRUNC:
		ygINFYaFDGa8 = TRUNCATE_EXISTING
	default:
		ygINFYaFDGa8 = OPEN_EXISTING
	}
	var ibvWO6_mblkb uint32 = FILE_ATTRIBUTE_NORMAL
	if gOrivca8ecRv&S_IWRITE == 0 {
		ibvWO6_mblkb = FILE_ATTRIBUTE_READONLY
		if ygINFYaFDGa8 == CREATE_ALWAYS {

			vFaWuUyY, wAwFLr4AGHg :=  /*line :1*/ EyjYVpa(zpdigBaban, wzWkRJs, pd1xDFCP03, zC0xxjnQnl, TRUNCATE_EXISTING, FILE_ATTRIBUTE_NORMAL, 0)
			switch wAwFLr4AGHg {
			case ERROR_FILE_NOT_FOUND, _ERROR_BAD_NETPATH, ERROR_PATH_NOT_FOUND:

			default:

				return vFaWuUyY, wAwFLr4AGHg
			}
		}
	}
	if ygINFYaFDGa8 == OPEN_EXISTING && wzWkRJs == GENERIC_READ {

		ibvWO6_mblkb |= FILE_FLAG_BACKUP_SEMANTICS
	}
	return  /*line :1*/ EyjYVpa(zpdigBaban, wzWkRJs, pd1xDFCP03, zC0xxjnQnl, ygINFYaFDGa8, ibvWO6_mblkb, 0)
}

func XJ0LdRoyxl_l(dBn476FVCy2C SRlvVjrYa, rd6leevODT []byte) (m5Tq_PL7 int, gOCcQzbcC error) {
	var ohzXBFIQ uint32
	wAwFLr4AGHg :=  /*line :1*/ Sg86ugD(dBn476FVCy2C, rd6leevODT, &ohzXBFIQ, nil)
	if wAwFLr4AGHg != nil {
		if wAwFLr4AGHg == ERROR_BROKEN_PIPE {

			return 0, nil
		}
		return 0, wAwFLr4AGHg
	}
	return  /*line :1*/ int(ohzXBFIQ), nil
}

func ZsMbSV2(dBn476FVCy2C SRlvVjrYa, rd6leevODT []byte) (m5Tq_PL7 int, gOCcQzbcC error) {
	var ohzXBFIQ uint32
	wAwFLr4AGHg :=  /*line :1*/ DaWZrhl(dBn476FVCy2C, rd6leevODT, &ohzXBFIQ, nil)
	if wAwFLr4AGHg != nil {
		return 0, wAwFLr4AGHg
	}
	return  /*line :1*/ int(ohzXBFIQ), nil
}

func Sg86ugD(dBn476FVCy2C SRlvVjrYa, rd6leevODT []byte, ohzXBFIQ *uint32, kR5OwbP *ZaNm5QSYC9) error {
	gOCcQzbcC :=  /*line :1*/ qzDaRXT(dBn476FVCy2C, rd6leevODT, ohzXBFIQ, kR5OwbP)
	if race.Enabled {
		if *ohzXBFIQ > 0 {
			 /*line :1*/ race.CPJO1X75d( /*line :1*/ unsafe.Pointer(&rd6leevODT[0]),  /*line :1*/ int(*ohzXBFIQ))
		}
		 /*line :1*/ race.X5ZIZytZ( /*line :1*/ unsafe.Pointer(&vfMtj6VtSU0))
	}
	if msanenabled && *ohzXBFIQ > 0 {
		 /*line :1*/ qAxFOdjnH( /*line :1*/ unsafe.Pointer(&rd6leevODT[0]),  /*line :1*/ int(*ohzXBFIQ))
	}
	if asanenabled && *ohzXBFIQ > 0 {
		 /*line :1*/ foveli( /*line :1*/ unsafe.Pointer(&rd6leevODT[0]),  /*line :1*/ int(*ohzXBFIQ))
	}
	return gOCcQzbcC
}

func DaWZrhl(dBn476FVCy2C SRlvVjrYa, rd6leevODT []byte, ohzXBFIQ *uint32, kR5OwbP *ZaNm5QSYC9) error {
	if race.Enabled {
		 /*line :1*/ race.AQrXiLDh3RTu( /*line :1*/ unsafe.Pointer(&vfMtj6VtSU0))
	}
	gOCcQzbcC :=  /*line :1*/ b5BOiDlX(dBn476FVCy2C, rd6leevODT, ohzXBFIQ, kR5OwbP)
	if race.Enabled && *ohzXBFIQ > 0 {
		 /*line :1*/ race.H2QwRu0Zggq( /*line :1*/ unsafe.Pointer(&rd6leevODT[0]),  /*line :1*/ int(*ohzXBFIQ))
	}
	if msanenabled && *ohzXBFIQ > 0 {
		 /*line :1*/ medMraWbUQnk( /*line :1*/ unsafe.Pointer(&rd6leevODT[0]),  /*line :1*/ int(*ohzXBFIQ))
	}
	if asanenabled && *ohzXBFIQ > 0 {
		 /*line :1*/ pFEI1g( /*line :1*/ unsafe.Pointer(&rd6leevODT[0]),  /*line :1*/ int(*ohzXBFIQ))
	}
	return gOCcQzbcC
}

var vfMtj6VtSU0 int64

var kJMfQhbT8gXJ =  /*line :1*/ jXlerCqn.NewProc("SetFilePointerEx")

const ptrSize =  /*line :1*/ unsafe.Sizeof( /*line :1*/ uintptr(0))



func g_5ArQFdW(fYz2KOsTDon SRlvVjrYa, xh26AmK5j int64, d__5hqO8r *int64, ddcX8W uint32) error {
	var mecVf9f6lPG O9Mga3
	if  /*line :1*/ unsafe.Sizeof( /*line :1*/ uintptr(0)) == 8 {
		_, _, mecVf9f6lPG =  /*line :1*/ VeAKF8sAwft( /*line :1*/ kJMfQhbT8gXJ.Addr(), 4,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(xh26AmK5j),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d__5hqO8r)),  /*line :1*/ uintptr(ddcX8W), 0, 0)
	} else {

		switch runtime.GOARCH {
		default:
			 /*line :1*/ panic("unsupported 32-bit architecture")
		case "386":

			_, _, mecVf9f6lPG =  /*line :1*/ VeAKF8sAwft( /*line :1*/ kJMfQhbT8gXJ.Addr(), 5,  /*line :1*/ uintptr(fYz2KOsTDon),  /*line :1*/ uintptr(xh26AmK5j),  /*line :1*/ uintptr(xh26AmK5j>>32),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d__5hqO8r)),  /*line :1*/ uintptr(ddcX8W), 0)
		case "arm":

			_, _, mecVf9f6lPG =  /*line :1*/ VeAKF8sAwft( /*line :1*/ kJMfQhbT8gXJ.Addr(), 6,  /*line :1*/ uintptr(fYz2KOsTDon), 0,  /*line :1*/ uintptr(xh26AmK5j),  /*line :1*/ uintptr(xh26AmK5j>>32),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(d__5hqO8r)),  /*line :1*/ uintptr(ddcX8W))
		}
	}
	if mecVf9f6lPG != 0 {
		return  /*line :1*/ wxZh39lL(mecVf9f6lPG)
	}
	return nil
}

func O4FqtrPqbHRY(dBn476FVCy2C SRlvVjrYa, yK6QReVkTKO int64, ddcX8W int) (elbjPiuk_a int64, gOCcQzbcC error) {
	var hxsZQgg30 uint32
	switch ddcX8W {
	case 0:
		hxsZQgg30 = FILE_BEGIN
	case 1:
		hxsZQgg30 = FILE_CURRENT
	case 2:
		hxsZQgg30 = FILE_END
	}
	gOCcQzbcC =  /*line :1*/ g_5ArQFdW(dBn476FVCy2C, yK6QReVkTKO, &elbjPiuk_a, hxsZQgg30)
	return
}

func Kv0fHx9Ttx0(dBn476FVCy2C SRlvVjrYa) (gOCcQzbcC error) {
	return  /*line :1*/ FhKJLgXjwG(dBn476FVCy2C)
}

var (
	IZGxwzygxL2	=  /*line :1*/ l9Igtlqb6Bf(STD_INPUT_HANDLE)
	PjqIbe	=  /*line :1*/ l9Igtlqb6Bf(STD_OUTPUT_HANDLE)
	ZAFZU4cW	=  /*line :1*/ l9Igtlqb6Bf(STD_ERROR_HANDLE)
)

func l9Igtlqb6Bf(vFaWuUyY int) (dBn476FVCy2C SRlvVjrYa) {
	iVSjWu, _ :=  /*line :1*/ Kgv7ZgME9(vFaWuUyY)
	return iVSjWu
}

const ImplementsGetwd = true

func PQMcWL4xt() (p2YNMzo string, gOCcQzbcC error) {
	tHPwhuTh :=  /*line :1*/ make([]uint16, 300)

	for {
		m5Tq_PL7, wAwFLr4AGHg :=  /*line :1*/ WRA3NP30nq( /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)), &tHPwhuTh[0])
		if wAwFLr4AGHg != nil {
			return "", wAwFLr4AGHg
		}
		if  /*line :1*/ int(m5Tq_PL7) <=  /*line :1*/ len(tHPwhuTh) {
			return  /*line :1*/ AODVXp8NM3sd(tHPwhuTh[:m5Tq_PL7]), nil
		}
		tHPwhuTh =  /*line :1*/ make([]uint16, m5Tq_PL7)
	}
}

func Ha4WpjjlJ9(ro5NCL9 string) (gOCcQzbcC error) {
	zpdigBaban, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ VqvfqI(zpdigBaban)
}

func HSuLak53N5P5(ro5NCL9 string, ediyaFOkNQmk uint32) (gOCcQzbcC error) {
	zpdigBaban, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ ZM4mayH(zpdigBaban, nil)
}

func RyMMPhj(ro5NCL9 string) (gOCcQzbcC error) {
	zpdigBaban, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ PQFNcx(zpdigBaban)
}

func Q83qUK4q(ro5NCL9 string) (gOCcQzbcC error) {
	zpdigBaban, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ N_MYEJ7Bb(zpdigBaban)
}

func E7wqrUo(gZT7BJ77CR47, z7hUNCaL string) (gOCcQzbcC error) {
	d4xMeWOatV4j, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(gZT7BJ77CR47)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	c7wlCtHEZpy, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(z7hUNCaL)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ Ib01EQM(d4xMeWOatV4j, c7wlCtHEZpy)
}

func WauQ78() (hzeU8Tl string, gOCcQzbcC error) {
	var m5Tq_PL7 uint32 = MAX_COMPUTERNAME_LENGTH + 1
	tHPwhuTh :=  /*line :1*/ make([]uint16, m5Tq_PL7)
	wAwFLr4AGHg :=  /*line :1*/ LwbTmPN(&tHPwhuTh[0], &m5Tq_PL7)
	if wAwFLr4AGHg != nil {
		return "", wAwFLr4AGHg
	}
	return  /*line :1*/ AODVXp8NM3sd(tHPwhuTh[:m5Tq_PL7]), nil
}

func Owge05R_TH_(dBn476FVCy2C SRlvVjrYa, cne_74 int64) (gOCcQzbcC error) {
	cj586BkW, wAwFLr4AGHg :=  /*line :1*/ O4FqtrPqbHRY(dBn476FVCy2C, 0, 1)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	defer  /*line :1*/ O4FqtrPqbHRY(dBn476FVCy2C, cj586BkW, 0)
	_, wAwFLr4AGHg =  /*line :1*/ O4FqtrPqbHRY(dBn476FVCy2C, cne_74, 0)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	wAwFLr4AGHg =  /*line :1*/ TYvDTZlPLYZ0(dBn476FVCy2C)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	return nil
}

func FcVgO8XHnM6(v3wOtFJUw7Nu *VhkskrPU1) (gOCcQzbcC error) {
	var d09pCxOvo T8WbUqAC3v
	 /*line :1*/ ItrSBaJMcT(&d09pCxOvo)
	*v3wOtFJUw7Nu =  /*line :1*/ DF1_wufO( /*line :1*/ d09pCxOvo.Nanoseconds())
	return nil
}

func RyJJSJTWj(rd6leevODT []SRlvVjrYa) (gOCcQzbcC error) {
	if  /*line :1*/ len(rd6leevODT) != 2 {
		return EINVAL
	}
	var iVSjWu, hxsZQgg30 SRlvVjrYa
	wAwFLr4AGHg :=  /*line :1*/ JtZhNTR6(&iVSjWu, &hxsZQgg30,  /*line :1*/ e0GOege_rv(), 0)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	rd6leevODT[0] = iVSjWu
	rd6leevODT[1] = hxsZQgg30
	return nil
}

func PLXbwsK(ro5NCL9 string, v3wOtFJUw7Nu []VhkskrPU1) (gOCcQzbcC error) {
	if  /*line :1*/ len(v3wOtFJUw7Nu) != 2 {
		return EINVAL
	}
	zpdigBaban, wAwFLr4AGHg :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	vFaWuUyY, wAwFLr4AGHg :=  /*line :1*/ EyjYVpa(zpdigBaban,
		FILE_WRITE_ATTRIBUTES, FILE_SHARE_WRITE, nil,
		OPEN_EXISTING, FILE_FLAG_BACKUP_SEMANTICS, 0)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	defer  /*line :1*/ Kv0fHx9Ttx0(vFaWuUyY)
	n0ogBVz := T8WbUqAC3v{}
	hxsZQgg30 := T8WbUqAC3v{}
	if  /*line :1*/ v3wOtFJUw7Nu[0].Nanoseconds() != 0 {
		n0ogBVz =  /*line :1*/ Gmwg7AR8( /*line :1*/ v3wOtFJUw7Nu[0].Nanoseconds())
	}
	if  /*line :1*/ v3wOtFJUw7Nu[0].Nanoseconds() != 0 {
		hxsZQgg30 =  /*line :1*/ Gmwg7AR8( /*line :1*/ v3wOtFJUw7Nu[1].Nanoseconds())
	}
	return  /*line :1*/ CA7BjaTdOLL(vFaWuUyY, nil, &n0ogBVz, &hxsZQgg30)
}


const _UTIME_OMIT = -1

func TzmRK5Cfk4(ro5NCL9 string, iUooHi8N []SLZae0blGj) (gOCcQzbcC error) {
	if  /*line :1*/ len(iUooHi8N) != 2 {
		return EINVAL
	}
	zpdigBaban, wAwFLr4AGHg :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	vFaWuUyY, wAwFLr4AGHg :=  /*line :1*/ EyjYVpa(zpdigBaban,
		FILE_WRITE_ATTRIBUTES, FILE_SHARE_WRITE, nil,
		OPEN_EXISTING, FILE_FLAG_BACKUP_SEMANTICS, 0)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	defer  /*line :1*/ Kv0fHx9Ttx0(vFaWuUyY)
	n0ogBVz := T8WbUqAC3v{}
	hxsZQgg30 := T8WbUqAC3v{}
	if iUooHi8N[0].WPGLV3fH1AX != _UTIME_OMIT {
		n0ogBVz =  /*line :1*/ Gmwg7AR8( /*line :1*/ EoXd2kIC2gG(iUooHi8N[0]))
	}
	if iUooHi8N[1].WPGLV3fH1AX != _UTIME_OMIT {
		hxsZQgg30 =  /*line :1*/ Gmwg7AR8( /*line :1*/ EoXd2kIC2gG(iUooHi8N[1]))
	}
	return  /*line :1*/ CA7BjaTdOLL(vFaWuUyY, nil, &n0ogBVz, &hxsZQgg30)
}

func KdVcdT(dBn476FVCy2C SRlvVjrYa) (gOCcQzbcC error) {
	return  /*line :1*/ IvXGh3(dBn476FVCy2C)
}

func FxoMx9RC_A(ro5NCL9 string, ediyaFOkNQmk uint32) (gOCcQzbcC error) {
	rd6leevODT, wAwFLr4AGHg :=  /*line :1*/ GcOmFfsibES(ro5NCL9)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	ibvWO6_mblkb, wAwFLr4AGHg :=  /*line :1*/ BDVJtIREn(rd6leevODT)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	if ediyaFOkNQmk&S_IWRITE != 0 {
		ibvWO6_mblkb &^= FILE_ATTRIBUTE_READONLY
	} else {
		ibvWO6_mblkb |= FILE_ATTRIBUTE_READONLY
	}
	return  /*line :1*/ D3KivaD(rd6leevODT, ibvWO6_mblkb)
}

func GSwyvNaG() error {
	return  /*line :1*/ dC_GmCDgrR.Find()
}

func JTYwZ3() error {
	return  /*line :1*/ dAUCsfm.Find()
}

const socket_error =  /*line :1*/ uintptr(^ /*line :1*/ uint32(0))



var AhAoze5k6qS bool

type NXUDVMl2 struct {
	JgQng5oWY8Ul	uint16
	VqKTvvibG	uint16
	AGDgR2QSLqW	[4]byte	
	WeOsNeM	[8]uint8
}

type H7tgUmcLUW struct {
	XYs3jaLEe	uint16
	D_H8WxCL7	uint16
	RwOjhxNykN	uint32
	G1aNyX7	[16]byte	
	Jl_oKnHTc	uint32
}

type ClqoLt struct {
	LcXKM5eE7rv	uint16
	WFMxTSzZGTq	[14]int8
}

type IXy5oynSaLQM struct {
	AI6wbnez	ClqoLt
	Q2zwOguwbpcu	[100]int8
}

type S4iroLVT4 interface {
	rJpZIBjJTTmW() (hcu1HkrXw unsafe.Pointer, mnIhqmz int32, gOCcQzbcC error)	
}

type Hx8lqxSkV struct {
	AMBsbT8war	int
	Q3zz2uH	[4]byte
	tfIX5wQ	NXUDVMl2
}

func (zC0xxjnQnl *Hx8lqxSkV) rJpZIBjJTTmW() (unsafe.Pointer, int32, error) {
	if zC0xxjnQnl.AMBsbT8war < 0 || zC0xxjnQnl.AMBsbT8war > 0xFFFF {
		return nil, 0, EINVAL
	}
	zC0xxjnQnl.tfIX5wQ.JgQng5oWY8Ul = AF_INET
	rd6leevODT := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&zC0xxjnQnl.tfIX5wQ.VqKTvvibG))
	rd6leevODT[0] =  /*line :1*/ byte(zC0xxjnQnl.AMBsbT8war >> 8)
	rd6leevODT[1] =  /*line :1*/ byte(zC0xxjnQnl.AMBsbT8war)
	zC0xxjnQnl.tfIX5wQ.AGDgR2QSLqW = zC0xxjnQnl.Q3zz2uH
	return  /*line :1*/ unsafe.Pointer(&zC0xxjnQnl.tfIX5wQ),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(zC0xxjnQnl.tfIX5wQ)), nil
}

type HKTAy7_BSU2 struct {
	NtLA3k	int
	AsutdJq2	uint32
	Y4XxFr	[16]byte
	yOXa9JaXq8c1	H7tgUmcLUW
}

func (zC0xxjnQnl *HKTAy7_BSU2) rJpZIBjJTTmW() (unsafe.Pointer, int32, error) {
	if zC0xxjnQnl.NtLA3k < 0 || zC0xxjnQnl.NtLA3k > 0xFFFF {
		return nil, 0, EINVAL
	}
	zC0xxjnQnl.yOXa9JaXq8c1.XYs3jaLEe = AF_INET6
	rd6leevODT := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&zC0xxjnQnl.yOXa9JaXq8c1.D_H8WxCL7))
	rd6leevODT[0] =  /*line :1*/ byte(zC0xxjnQnl.NtLA3k >> 8)
	rd6leevODT[1] =  /*line :1*/ byte(zC0xxjnQnl.NtLA3k)
	zC0xxjnQnl.yOXa9JaXq8c1.Jl_oKnHTc = zC0xxjnQnl.AsutdJq2
	zC0xxjnQnl.yOXa9JaXq8c1.G1aNyX7 = zC0xxjnQnl.Y4XxFr
	return  /*line :1*/ unsafe.Pointer(&zC0xxjnQnl.yOXa9JaXq8c1),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(zC0xxjnQnl.yOXa9JaXq8c1)), nil
}

type VIpNMZaTPCss struct {
	ZCY53FT	uint16
	N354aK0	[UNIX_PATH_MAX]int8
}

type Vaumry9a struct {
	Ql_3XaCdaHy	string
	cZA96W	VIpNMZaTPCss
}

func (zC0xxjnQnl *Vaumry9a) rJpZIBjJTTmW() (unsafe.Pointer, int32, error) {
	hzeU8Tl := zC0xxjnQnl.Ql_3XaCdaHy
	m5Tq_PL7 :=  /*line :1*/ len(hzeU8Tl)
	if m5Tq_PL7 >  /*line :1*/ len(zC0xxjnQnl.cZA96W.N354aK0) {
		return nil, 0, EINVAL
	}
	if m5Tq_PL7 ==  /*line :1*/ len(zC0xxjnQnl.cZA96W.N354aK0) && hzeU8Tl[0] != '@' {
		return nil, 0, EINVAL
	}
	zC0xxjnQnl.cZA96W.ZCY53FT = AF_UNIX
	for hA99q3EOK := 0; hA99q3EOK < m5Tq_PL7; hA99q3EOK++ {
		zC0xxjnQnl.cZA96W.N354aK0[hA99q3EOK] =  /*line :1*/ int8(hzeU8Tl[hA99q3EOK])
	}

	q0Wopd4Bgjaq :=  /*line :1*/ int32(2)
	if m5Tq_PL7 > 0 {
		q0Wopd4Bgjaq +=  /*line :1*/ int32(m5Tq_PL7) + 1
	}
	if zC0xxjnQnl.cZA96W.N354aK0[0] == '@' {
		zC0xxjnQnl.cZA96W.N354aK0[0] = 0

		q0Wopd4Bgjaq--
	}

	return  /*line :1*/ unsafe.Pointer(&zC0xxjnQnl.cZA96W), q0Wopd4Bgjaq, nil
}

func (nmgZCOdj *IXy5oynSaLQM) Sockaddr() (S4iroLVT4, error) {
	switch nmgZCOdj.AI6wbnez.LcXKM5eE7rv {
	case AF_UNIX:
		jm7LRVfp := (* /*line :1*/ VIpNMZaTPCss)( /*line :1*/ unsafe.Pointer(nmgZCOdj))
		zC0xxjnQnl :=  /*line :1*/ new(Vaumry9a)
		if jm7LRVfp.N354aK0[0] == 0 {

			jm7LRVfp.N354aK0[0] = '@'
		}

		m5Tq_PL7 := 0
		for m5Tq_PL7 <  /*line :1*/ len(jm7LRVfp.N354aK0) && jm7LRVfp.N354aK0[m5Tq_PL7] != 0 {
			m5Tq_PL7++
		}
		zC0xxjnQnl.Ql_3XaCdaHy =  /*line :1*/ string( /*line :1*/ unsafe.Slice((* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&jm7LRVfp.N354aK0[0])), m5Tq_PL7))
		return zC0xxjnQnl, nil

	case AF_INET:
		jm7LRVfp := (* /*line :1*/ NXUDVMl2)( /*line :1*/ unsafe.Pointer(nmgZCOdj))
		zC0xxjnQnl :=  /*line :1*/ new(Hx8lqxSkV)
		rd6leevODT := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&jm7LRVfp.VqKTvvibG))
		zC0xxjnQnl.AMBsbT8war =  /*line :1*/ int(rd6leevODT[0])<<8 +  /*line :1*/ int(rd6leevODT[1])
		zC0xxjnQnl.Q3zz2uH = jm7LRVfp.AGDgR2QSLqW
		return zC0xxjnQnl, nil

	case AF_INET6:
		jm7LRVfp := (* /*line :1*/ H7tgUmcLUW)( /*line :1*/ unsafe.Pointer(nmgZCOdj))
		zC0xxjnQnl :=  /*line :1*/ new(HKTAy7_BSU2)
		rd6leevODT := (*[2] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&jm7LRVfp.D_H8WxCL7))
		zC0xxjnQnl.NtLA3k =  /*line :1*/ int(rd6leevODT[0])<<8 +  /*line :1*/ int(rd6leevODT[1])
		zC0xxjnQnl.AsutdJq2 = jm7LRVfp.Jl_oKnHTc
		zC0xxjnQnl.Y4XxFr = jm7LRVfp.G1aNyX7
		return zC0xxjnQnl, nil
	}
	return nil, EAFNOSUPPORT
}

func RQpzEduJMY(c0DvvOaYG, b85up2dBCm, qUFjymZN6q7 int) (dBn476FVCy2C SRlvVjrYa, gOCcQzbcC error) {
	if c0DvvOaYG == AF_INET6 && AhAoze5k6qS {
		return InvalidHandle, EAFNOSUPPORT
	}
	return  /*line :1*/ wA8pjx79d( /*line :1*/ int32(c0DvvOaYG),  /*line :1*/ int32(b85up2dBCm),  /*line :1*/ int32(qUFjymZN6q7))
}

func ABpNrEFh9ZGb(dBn476FVCy2C SRlvVjrYa, jAzpOh, _Yc7TXX int, fsp_0a7i int) (gOCcQzbcC error) {
	b0abCyNZ :=  /*line :1*/ int32(fsp_0a7i)
	return  /*line :1*/ JwtaZG7GaLGD(dBn476FVCy2C,  /*line :1*/ int32(jAzpOh),  /*line :1*/ int32(_Yc7TXX), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&b0abCyNZ)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(b0abCyNZ)))
}

func U5lDiJBkm0(dBn476FVCy2C SRlvVjrYa, zC0xxjnQnl S4iroLVT4) (gOCcQzbcC error) {
	hcu1HkrXw, m5Tq_PL7, gOCcQzbcC :=  /*line :1*/ zC0xxjnQnl.rJpZIBjJTTmW()
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ tKqRrokO7GR(dBn476FVCy2C, hcu1HkrXw, m5Tq_PL7)
}

func AEQIkiwa89e(dBn476FVCy2C SRlvVjrYa, zC0xxjnQnl S4iroLVT4) (gOCcQzbcC error) {
	hcu1HkrXw, m5Tq_PL7, gOCcQzbcC :=  /*line :1*/ zC0xxjnQnl.rJpZIBjJTTmW()
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ poAg04RDL(dBn476FVCy2C, hcu1HkrXw, m5Tq_PL7)
}

func DaM3_rD4F(dBn476FVCy2C SRlvVjrYa) (zC0xxjnQnl S4iroLVT4, gOCcQzbcC error) {
	var nmgZCOdj IXy5oynSaLQM
	cE4tZhPbq :=  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(nmgZCOdj))
	if gOCcQzbcC =  /*line :1*/ aB6zPKR3lN(dBn476FVCy2C, &nmgZCOdj, &cE4tZhPbq); gOCcQzbcC != nil {
		return
	}
	return  /*line :1*/ nmgZCOdj.Sockaddr()
}

func BWaUaPS4sKD9(dBn476FVCy2C SRlvVjrYa) (zC0xxjnQnl S4iroLVT4, gOCcQzbcC error) {
	var nmgZCOdj IXy5oynSaLQM
	cE4tZhPbq :=  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(nmgZCOdj))
	if gOCcQzbcC =  /*line :1*/ r3XXDB(dBn476FVCy2C, &nmgZCOdj, &cE4tZhPbq); gOCcQzbcC != nil {
		return
	}
	return  /*line :1*/ nmgZCOdj.Sockaddr()
}

func DUh335tLkj6b(wzPhJhIFoI SRlvVjrYa, m5Tq_PL7 int) (gOCcQzbcC error) {
	return  /*line :1*/ nLqRlHL2wYJI(wzPhJhIFoI,  /*line :1*/ int32(m5Tq_PL7))
}

func H3rGVgc1(dBn476FVCy2C SRlvVjrYa, tLIUarFXW4S int) (gOCcQzbcC error) {
	return  /*line :1*/ hIfg4o6v(dBn476FVCy2C,  /*line :1*/ int32(tLIUarFXW4S))
}

func CzWMuENIar(wzPhJhIFoI SRlvVjrYa, m8dz2Ors *QbWhAp5CqTG, t8xduL uint32, lu937Dl_88 *uint32, uADzMh0kZ7hL uint32, c7wlCtHEZpy S4iroLVT4, kR5OwbP *ZaNm5QSYC9, bpTsK_7YjdlT *byte) (gOCcQzbcC error) {
	var nmgZCOdj unsafe.Pointer
	var mnIhqmz int32
	if c7wlCtHEZpy != nil {
		nmgZCOdj, mnIhqmz, gOCcQzbcC =  /*line :1*/ c7wlCtHEZpy.rJpZIBjJTTmW()
		if gOCcQzbcC != nil {
			return gOCcQzbcC
		}
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ epTJJ70Fc.Addr(), 9,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m8dz2Ors)),  /*line :1*/ uintptr(t8xduL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lu937Dl_88)),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nmgZCOdj)),  /*line :1*/ uintptr(mnIhqmz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bpTsK_7YjdlT)))
	if iwAHyJGkWJy == socket_error {
		if mecVf9f6lPG != 0 {
			gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
		} else {
			gOCcQzbcC = EINVAL
		}
	}
	return gOCcQzbcC
}

func m2VgBhm(wzPhJhIFoI SRlvVjrYa, m8dz2Ors *QbWhAp5CqTG, t8xduL uint32, lu937Dl_88 *uint32, uADzMh0kZ7hL uint32, c7wlCtHEZpy *Hx8lqxSkV, kR5OwbP *ZaNm5QSYC9, bpTsK_7YjdlT *byte) (gOCcQzbcC error) {
	nmgZCOdj, mnIhqmz, gOCcQzbcC :=  /*line :1*/ c7wlCtHEZpy.rJpZIBjJTTmW()
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ epTJJ70Fc.Addr(), 9,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m8dz2Ors)),  /*line :1*/ uintptr(t8xduL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lu937Dl_88)),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nmgZCOdj)),  /*line :1*/ uintptr(mnIhqmz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bpTsK_7YjdlT)))
	if iwAHyJGkWJy == socket_error {
		if mecVf9f6lPG != 0 {
			gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
		} else {
			gOCcQzbcC = EINVAL
		}
	}
	return gOCcQzbcC
}

func pmhM9rv(wzPhJhIFoI SRlvVjrYa, m8dz2Ors *QbWhAp5CqTG, t8xduL uint32, lu937Dl_88 *uint32, uADzMh0kZ7hL uint32, c7wlCtHEZpy *HKTAy7_BSU2, kR5OwbP *ZaNm5QSYC9, bpTsK_7YjdlT *byte) (gOCcQzbcC error) {
	nmgZCOdj, mnIhqmz, gOCcQzbcC :=  /*line :1*/ c7wlCtHEZpy.rJpZIBjJTTmW()
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV( /*line :1*/ epTJJ70Fc.Addr(), 9,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(m8dz2Ors)),  /*line :1*/ uintptr(t8xduL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lu937Dl_88)),  /*line :1*/ uintptr(uADzMh0kZ7hL),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nmgZCOdj)),  /*line :1*/ uintptr(mnIhqmz),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(bpTsK_7YjdlT)))
	if iwAHyJGkWJy == socket_error {
		if mecVf9f6lPG != 0 {
			gOCcQzbcC =  /*line :1*/ wxZh39lL(mecVf9f6lPG)
		} else {
			gOCcQzbcC = EINVAL
		}
	}
	return gOCcQzbcC
}

func A4sd5Q() error {
	return  /*line :1*/ x25ZA6ZEaO4.Find()
}

var cQ6w6rn6 struct {
	b_cotd	sync.LhBfwn6wa1x
	iqDs38TG9	uintptr
	epGgcJBCHfrF	error
}

func D0E_y1QdfF0() error {
	 /*line :1*/ cQ6w6rn6.b_cotd.Do(func() {
		var wzPhJhIFoI SRlvVjrYa
		wzPhJhIFoI, cQ6w6rn6.epGgcJBCHfrF =  /*line :1*/ RQpzEduJMY(AF_INET, SOCK_STREAM, IPPROTO_TCP)
		if cQ6w6rn6.epGgcJBCHfrF != nil {
			return
		}
		defer  /*line :1*/ FhKJLgXjwG(wzPhJhIFoI)
		var m5Tq_PL7 uint32
		cQ6w6rn6.epGgcJBCHfrF =  /*line :1*/ YvCdY02(wzPhJhIFoI,
			SIO_GET_EXTENSION_FUNCTION_POINTER,
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&XJR3ag)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(XJR3ag)),
			(* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&cQ6w6rn6.iqDs38TG9)),
			 /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(cQ6w6rn6.iqDs38TG9)),
			&m5Tq_PL7, nil, 0)
	})
	return cQ6w6rn6.epGgcJBCHfrF
}

func jupwa8X9ho(wzPhJhIFoI SRlvVjrYa, hzeU8Tl unsafe.Pointer, eixaYPaXLSYp int32, rWaKrX *byte, tp8xY1d8KRy uint32, kb0aA8gy6I *uint32, kR5OwbP *ZaNm5QSYC9) (gOCcQzbcC error) {
	iwAHyJGkWJy, _, mecVf9f6lPG :=  /*line :1*/ Rc0O05UsV(cQ6w6rn6.iqDs38TG9, 7,  /*line :1*/ uintptr(wzPhJhIFoI),  /*line :1*/ uintptr(hzeU8Tl),  /*line :1*/ uintptr(eixaYPaXLSYp),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rWaKrX)),  /*line :1*/ uintptr(tp8xY1d8KRy),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kb0aA8gy6I)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(kR5OwbP)), 0, 0)
	if iwAHyJGkWJy == 0 {
		if mecVf9f6lPG != 0 {
			gOCcQzbcC =  /*line :1*/ error(mecVf9f6lPG)
		} else {
			gOCcQzbcC = EINVAL
		}
	}
	return
}

func JLt4U2f2Gvo(dBn476FVCy2C SRlvVjrYa, zC0xxjnQnl S4iroLVT4, rWaKrX *byte, tp8xY1d8KRy uint32, kb0aA8gy6I *uint32, kR5OwbP *ZaNm5QSYC9) error {
	gOCcQzbcC :=  /*line :1*/ D0E_y1QdfF0()
	if gOCcQzbcC != nil {
		return  /*line :1*/ errorspkg.Su6g6hRoi9X("failed to find ConnectEx: " +  /*line :1*/ gOCcQzbcC.Error())
	}
	hcu1HkrXw, m5Tq_PL7, gOCcQzbcC :=  /*line :1*/ zC0xxjnQnl.rJpZIBjJTTmW()
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	return  /*line :1*/ jupwa8X9ho(dBn476FVCy2C, hcu1HkrXw, m5Tq_PL7, rWaKrX, tp8xY1d8KRy, kb0aA8gy6I, kR5OwbP)
}


type Pd6f6PxVa struct {
	DInKorDQLw1	T8WbUqAC3v
	TPYcuXne_w	T8WbUqAC3v
	UiqiQqR	T8WbUqAC3v
	JJK8Dbfduef	T8WbUqAC3v
}

type HV8Mf1 struct {
	ZBjv37 uint32
}

func (hxsZQgg30 HV8Mf1) Exited() bool	{ return true }

func (hxsZQgg30 HV8Mf1) ExitStatus() int	{ return  /*line :1*/ int(hxsZQgg30.ZBjv37) }

func (hxsZQgg30 HV8Mf1) Signal() EslOtdor	{ return -1 }

func (hxsZQgg30 HV8Mf1) CoreDump() bool	{ return false }

func (hxsZQgg30 HV8Mf1) Stopped() bool	{ return false }

func (hxsZQgg30 HV8Mf1) Continued() bool	{ return false }

func (hxsZQgg30 HV8Mf1) StopSignal() EslOtdor	{ return -1 }

func (hxsZQgg30 HV8Mf1) Signaled() bool	{ return false }

func (hxsZQgg30 HV8Mf1) TrapCause() int	{ return -1 }



type SLZae0blGj struct {
	ECuErxxfdoOW	int64
	WPGLV3fH1AX	int64
}

func EoXd2kIC2gG(iUooHi8N SLZae0blGj) int64 {
	return  /*line :1*/ int64(iUooHi8N.ECuErxxfdoOW)*1e9 +  /*line :1*/ int64(iUooHi8N.WPGLV3fH1AX)
}

func J9KJwaNJRl(tAjR0818K5K int64) (iUooHi8N SLZae0blGj) {
	iUooHi8N.ECuErxxfdoOW = tAjR0818K5K / 1e9
	iUooHi8N.WPGLV3fH1AX = tAjR0818K5K % 1e9
	return
}

func OACBOl(dBn476FVCy2C SRlvVjrYa) (vjbb5CKEjoKY SRlvVjrYa, zC0xxjnQnl S4iroLVT4, gOCcQzbcC error) {
	return 0, nil, EWINDOWS
}
func Dfd26Fw(dBn476FVCy2C SRlvVjrYa, rd6leevODT []byte, uADzMh0kZ7hL int) (m5Tq_PL7 int, d4xMeWOatV4j S4iroLVT4, gOCcQzbcC error) {
	return 0, nil, EWINDOWS
}
func L6lzuiA2beJA(dBn476FVCy2C SRlvVjrYa, rd6leevODT []byte, uADzMh0kZ7hL int, c7wlCtHEZpy S4iroLVT4) (gOCcQzbcC error) {
	return EWINDOWS
}
func WSSbCbX(dBn476FVCy2C SRlvVjrYa, jAzpOh, _Yc7TXX int, v3wOtFJUw7Nu *VhkskrPU1) (gOCcQzbcC error) {
	return EWINDOWS
}

type KW8FJX7 struct {
	V6pIuhwCu	int32
	M0TrEat	int32
}

type u9NUUj8Kq7 struct {
	L8P8ult	uint16
	Z95uALy	uint16
}

type W1r5YRtipLa struct {
	P8O2NF0WA	[4]byte	
	Nompo3TK	[4]byte	
}

type C5ybc0 struct {
	RmQSpqMVN	[16]byte	
	GOMCVrEx1k	uint32
}

func QWe7U4(dBn476FVCy2C SRlvVjrYa, jAzpOh, _Yc7TXX int) (int, error)	{ return -1, EWINDOWS }

func JPiqBG(dBn476FVCy2C SRlvVjrYa, jAzpOh, _Yc7TXX int, cE4tZhPbq *KW8FJX7) (gOCcQzbcC error) {
	mOZVRlrXu := u9NUUj8Kq7{L8P8ult:  /*line :1*/ uint16(cE4tZhPbq.V6pIuhwCu), Z95uALy:  /*line :1*/ uint16(cE4tZhPbq.M0TrEat)}
	return  /*line :1*/ JwtaZG7GaLGD(dBn476FVCy2C,  /*line :1*/ int32(jAzpOh),  /*line :1*/ int32(_Yc7TXX), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&mOZVRlrXu)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(mOZVRlrXu)))
}

func Fnnrp31xH8y(dBn476FVCy2C SRlvVjrYa, jAzpOh, _Yc7TXX int, fsp_0a7i [4]byte) (gOCcQzbcC error) {
	return  /*line :1*/ JwtaZG7GaLGD(dBn476FVCy2C,  /*line :1*/ int32(jAzpOh),  /*line :1*/ int32(_Yc7TXX), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&fsp_0a7i[0])), 4)
}
func F2RtoD(dBn476FVCy2C SRlvVjrYa, jAzpOh, _Yc7TXX int, ay_3NsTa *W1r5YRtipLa) (gOCcQzbcC error) {
	return  /*line :1*/ JwtaZG7GaLGD(dBn476FVCy2C,  /*line :1*/ int32(jAzpOh),  /*line :1*/ int32(_Yc7TXX), (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(ay_3NsTa)),  /*line :1*/ int32( /*line :1*/ unsafe.Sizeof(*ay_3NsTa)))
}
func G8xAk43n(dBn476FVCy2C SRlvVjrYa, jAzpOh, _Yc7TXX int, ay_3NsTa *C5ybc0) (gOCcQzbcC error) {
	return EWINDOWS
}

func HLpaMN9pU4() (aJF0So int)	{ return  /*line :1*/ int( /*line :1*/ gEvRXHL77A()) }

func YWRpP1l(hzeU8Tl *uint16, gv3kki *RsoZmJigCkGe) (fYz2KOsTDon SRlvVjrYa, gOCcQzbcC error) {
	
	
	
	
	
	
	
	
	var c8sDbwZ5l0m w8vwhA
	fYz2KOsTDon, gOCcQzbcC =  /*line :1*/ deKyXFXaIT(hzeU8Tl, &c8sDbwZ5l0m)
	if gOCcQzbcC == nil {
		 /*line :1*/ m0EKKA(gv3kki, &c8sDbwZ5l0m)
	}
	return
}

func ZAQEj8rfk(fYz2KOsTDon SRlvVjrYa, gv3kki *RsoZmJigCkGe) (gOCcQzbcC error) {
	var c8sDbwZ5l0m w8vwhA
	gOCcQzbcC =  /*line :1*/ oXHZwN5(fYz2KOsTDon, &c8sDbwZ5l0m)
	if gOCcQzbcC == nil {
		 /*line :1*/ m0EKKA(gv3kki, &c8sDbwZ5l0m)
	}
	return
}

func hwGeZ_A2Q(aJF0So int) (*GRU31Ht3FPn, error) {
	piKsTT2T, gOCcQzbcC :=  /*line :1*/ UsCsCLMhll(TH32CS_SNAPPROCESS, 0)
	if gOCcQzbcC != nil {
		return nil, gOCcQzbcC
	}
	defer  /*line :1*/ FhKJLgXjwG(piKsTT2T)
	var mowY07B GRU31Ht3FPn
	mowY07B.NgWYrRzCdYqS =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(mowY07B))
	if gOCcQzbcC =  /*line :1*/ WkbE7c(piKsTT2T, &mowY07B); gOCcQzbcC != nil {
		return nil, gOCcQzbcC
	}
	for {
		if mowY07B.QU8DfIi ==  /*line :1*/ uint32(aJF0So) {
			return &mowY07B, nil
		}
		gOCcQzbcC =  /*line :1*/ IhA67uCoej5(piKsTT2T, &mowY07B)
		if gOCcQzbcC != nil {
			return nil, gOCcQzbcC
		}
	}
}

func IAWmZ0vBKwfC() (h4OwadvgYHJ int) {
	kn0rwyc, gOCcQzbcC :=  /*line :1*/ hwGeZ_A2Q( /*line :1*/ HLpaMN9pU4())
	if gOCcQzbcC != nil {
		return -1
	}
	return  /*line :1*/ int(kn0rwyc.TvE7hf)
}

func kk8bg92kuS8N(dBn476FVCy2C SRlvVjrYa, eun1Jlud5A []uint16) ([]uint16, error) {
	const (
		FILE_NAME_NORMALIZED	= 0
		VOLUME_NAME_DOS	= 0
	)
	for {
		m5Tq_PL7, gOCcQzbcC :=  /*line :1*/ a7UqCnQe0(dBn476FVCy2C, &eun1Jlud5A[0],  /*line :1*/ uint32( /*line :1*/ len(eun1Jlud5A)), FILE_NAME_NORMALIZED|VOLUME_NAME_DOS)
		if gOCcQzbcC == nil {
			eun1Jlud5A = eun1Jlud5A[:m5Tq_PL7]
			break
		}
		if gOCcQzbcC != _ERROR_NOT_ENOUGH_MEMORY {
			return nil, gOCcQzbcC
		}
		eun1Jlud5A =  /*line :1*/ append(eun1Jlud5A,  /*line :1*/ make([]uint16, m5Tq_PL7- /*line :1*/ uint32( /*line :1*/ len(eun1Jlud5A)))...)
	}
	return eun1Jlud5A, nil
}

func LCrVI4(dBn476FVCy2C SRlvVjrYa) (gOCcQzbcC error) {
	var eun1Jlud5A [MAX_PATH + 1]uint16
	ro5NCL9, gOCcQzbcC :=  /*line :1*/ kk8bg92kuS8N(dBn476FVCy2C, eun1Jlud5A[:])
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}

	if  /*line :1*/ len(ro5NCL9) >= 4 && ro5NCL9[0] == '\\' && ro5NCL9[1] == '\\' && ro5NCL9[2] == '?' && ro5NCL9[3] == '\\' {
		ro5NCL9 = ro5NCL9[4:]
	}
	return  /*line :1*/ VqvfqI(&ro5NCL9[0])
}


func T8qsT71a(gZT7BJ77CR47, z7hUNCaL string) (gOCcQzbcC error)	{ return EWINDOWS }
func FkBPYQ(ro5NCL9, iZJxy2 string) (gOCcQzbcC error)	{ return EWINDOWS }

func GuHWdMM(dBn476FVCy2C SRlvVjrYa, ediyaFOkNQmk uint32) (gOCcQzbcC error)	{ return EWINDOWS }
func VGkqsd(ro5NCL9 string, csPJgE int, trJBnBjzgX int) (gOCcQzbcC error)	{ return EWINDOWS }
func Ba3QJ7IqQ(ro5NCL9 string, csPJgE int, trJBnBjzgX int) (gOCcQzbcC error)	{ return EWINDOWS }
func FQi5_BSB(dBn476FVCy2C SRlvVjrYa, csPJgE int, trJBnBjzgX int) (gOCcQzbcC error)	{ return EWINDOWS }

func OeMe9jE() (csPJgE int)	{ return -1 }
func ZNXmxC() (o1NNPnUZYN int)	{ return -1 }
func ZA1sFx5_XxZX() (trJBnBjzgX int)	{ return -1 }
func BDGfBD() (afeIVmObcg int)	{ return -1 }
func WafuobroC() (xmzUXk4N_u []int, gOCcQzbcC error)	{ return nil, EWINDOWS }

type EslOtdor int

func (wzPhJhIFoI EslOtdor) Signal()	{}

func (wzPhJhIFoI EslOtdor) String() string {
	if 0 <= wzPhJhIFoI &&  /*line :1*/ int(wzPhJhIFoI) <  /*line :1*/ len(wDO4xIZNiGl) {
		ho6QRNE2 := wDO4xIZNiGl[wzPhJhIFoI]
		if ho6QRNE2 != "" {
			return ho6QRNE2
		}
	}
	return "signal " +  /*line :1*/ itoa.V25ba9G5( /*line :1*/ int(wzPhJhIFoI))
}

func ZMcvxf() error {
	return  /*line :1*/ soOK9W.Find()
}


func EagIjlI(ro5NCL9 string, eun1Jlud5A []byte) (m5Tq_PL7 int, gOCcQzbcC error) {
	dBn476FVCy2C, gOCcQzbcC :=  /*line :1*/ EyjYVpa( /*line :1*/ EtPVNA(ro5NCL9), GENERIC_READ, 0, nil, OPEN_EXISTING,
		FILE_FLAG_OPEN_REPARSE_POINT|FILE_FLAG_BACKUP_SEMANTICS, 0)
	if gOCcQzbcC != nil {
		return -1, gOCcQzbcC
	}
	defer  /*line :1*/ FhKJLgXjwG(dBn476FVCy2C)

	dKpBZmHQMBW6 :=  /*line :1*/ make([]byte, MAXIMUM_REPARSE_DATA_BUFFER_SIZE)
	var dKjLqwpa uint32
	gOCcQzbcC =  /*line :1*/ EynTCZg(dBn476FVCy2C, FSCTL_GET_REPARSE_POINT, nil, 0, &dKpBZmHQMBW6[0],  /*line :1*/ uint32( /*line :1*/ len(dKpBZmHQMBW6)), &dKjLqwpa, nil)
	if gOCcQzbcC != nil {
		return -1, gOCcQzbcC
	}

	aXVuiL := (* /*line :1*/ aUaISxWx2q1)( /*line :1*/ unsafe.Pointer(&dKpBZmHQMBW6[0]))
	var wzPhJhIFoI string
	switch aXVuiL.Ct81WfAEIFFo {
	case IO_REPARSE_TAG_SYMLINK:
		gv3kki := (* /*line :1*/ qgA9X8dVIL)( /*line :1*/ unsafe.Pointer(&aXVuiL.ofwNwEb5aYf))
		rd6leevODT := (*[0xffff] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&gv3kki.TKCfxANSN[0]))
		wzPhJhIFoI =  /*line :1*/ AODVXp8NM3sd(rd6leevODT[gv3kki.O2kMCXRqJjg/2 : (gv3kki.O2kMCXRqJjg+gv3kki.DwansDx)/2])
		if gv3kki.D4daKwFX&_SYMLINK_FLAG_RELATIVE == 0 {
			if  /*line :1*/ len(wzPhJhIFoI) >= 4 && wzPhJhIFoI[:4] == `\??\` {
				wzPhJhIFoI = wzPhJhIFoI[4:]
				switch {
				case  /*line :1*/ len(wzPhJhIFoI) >= 2 && wzPhJhIFoI[1] == ':':

				case  /*line :1*/ len(wzPhJhIFoI) >= 4 && wzPhJhIFoI[:4] == `UNC\`:
					wzPhJhIFoI = `\\` + wzPhJhIFoI[4:]
				default:

				}
			} else {

			}
		}
	case _IO_REPARSE_TAG_MOUNT_POINT:
		gv3kki := (* /*line :1*/ i9NdjG)( /*line :1*/ unsafe.Pointer(&aXVuiL.ofwNwEb5aYf))
		rd6leevODT := (*[0xffff] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&gv3kki.HbIKNUMDT4V[0]))
		wzPhJhIFoI =  /*line :1*/ AODVXp8NM3sd(rd6leevODT[gv3kki.NTZIOHn8z/2 : (gv3kki.NTZIOHn8z+gv3kki.FJV5If)/2])
		if  /*line :1*/ len(wzPhJhIFoI) >= 4 && wzPhJhIFoI[:4] == `\??\` {
			wzPhJhIFoI = wzPhJhIFoI[4:]
		} else {

		}
	default:

		return -1, ENOENT
	}
	m5Tq_PL7 =  /*line :1*/ copy(eun1Jlud5A, [] /*line :1*/ byte(wzPhJhIFoI))

	return m5Tq_PL7, nil
}


func FuGdBte(a8KPcHKEe SRlvVjrYa, j6_sksnZU SRlvVjrYa, nIB9uB1 uint32, d5Ldozqc uint32) (SRlvVjrYa, error) {
	return  /*line :1*/ nAuRSVpaHRa3(a8KPcHKEe, j6_sksnZU,  /*line :1*/ uintptr(nIB9uB1), d5Ldozqc)
}


func FGLw869WJf(j6_sksnZU SRlvVjrYa, lh_XEK *uint32, nIB9uB1 *uint32, kR5OwbP **ZaNm5QSYC9, vYbNuNe uint32) error {
	var h9D7uP uintptr
	var yGuaHUeaOi9 *uintptr
	if nIB9uB1 != nil {
		h9D7uP =  /*line :1*/ uintptr(*nIB9uB1)
		yGuaHUeaOi9 = &h9D7uP
	}
	gOCcQzbcC :=  /*line :1*/ cFzyNYNTK(j6_sksnZU, lh_XEK, yGuaHUeaOi9, kR5OwbP, vYbNuNe)
	if nIB9uB1 != nil {
		*nIB9uB1 =  /*line :1*/ uint32(h9D7uP)
		if  /*line :1*/ uintptr(*nIB9uB1) != h9D7uP && gOCcQzbcC == nil {
			gOCcQzbcC =  /*line :1*/ errorspkg.Su6g6hRoi9X("GetQueuedCompletionStatus returned key overflow")
		}
	}
	return gOCcQzbcC
}


func RrHamh(j6_sksnZU SRlvVjrYa, lh_XEK uint32, nIB9uB1 uint32, kR5OwbP *ZaNm5QSYC9) error {
	return  /*line :1*/ jlfCPe(j6_sksnZU, lh_XEK,  /*line :1*/ uintptr(nIB9uB1), kR5OwbP)
}




func itvjHDoHQBe(lPbMNpz uint32) (*vYoIhduc, error) {
	var jFhAAWnAX uintptr
	gOCcQzbcC :=  /*line :1*/ eavEkDP(nil, lPbMNpz, 0, &jFhAAWnAX)
	if gOCcQzbcC != ERROR_INSUFFICIENT_BUFFER {
		if gOCcQzbcC == nil {
			return nil,  /*line :1*/ errorspkg.Su6g6hRoi9X("unable to query buffer size from InitializeProcThreadAttributeList")
		}
		return nil, gOCcQzbcC
	}

	xGC9WC4 := (* /*line :1*/ vYoIhduc)( /*line :1*/ unsafe.Pointer(& /*line :1*/ make([]byte, jFhAAWnAX)[0]))
	gOCcQzbcC =  /*line :1*/ eavEkDP(xGC9WC4, lPbMNpz, 0, &jFhAAWnAX)
	if gOCcQzbcC != nil {
		return nil, gOCcQzbcC
	}
	return xGC9WC4, nil
}

























func T4R52_D4B1Z(nIB9uB1 SRlvVjrYa, bjKxPq uint32, hzeU8Tl *uint16, y9xWORA1z *uint32, r4wfXKbPIsE *uint32, zIaTh7G3 *uint16, hiS864dY *uint32, aDyiZ7NGC *T8WbUqAC3v) (zlknHC error) {
	return  /*line :1*/ oNLWBp2Sr(nIB9uB1, bjKxPq, hzeU8Tl, y9xWORA1z, r4wfXKbPIsE, zIaTh7G3, hiS864dY, aDyiZ7NGC)
}
