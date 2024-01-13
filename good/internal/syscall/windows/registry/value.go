//line :1
//go:build windows

package GMNBnK

import (
	errors "iAHoxjmM"
	syscall "bUKeamF"
	utf16 "DtJCLKevRp"
	"unsafe"
)

const (
	NONE	= 0
	SZ	= 1
	EXPAND_SZ	= 2
	BINARY	= 3
	DWORD	= 4
	DWORD_BIG_ENDIAN	= 5
	LINK	= 6
	MULTI_SZ	= 7
	RESOURCE_LIST	= 8
	FULL_RESOURCE_DESCRIPTOR	= 9
	RESOURCE_REQUIREMENTS_LIST	= 10
	QWORD	= 11
)

var (
	HaEFBUf7L	= syscall.ERROR_MORE_DATA

	HfAb1QCNR	= syscall.ERROR_FILE_NOT_FOUND

	U12T7G	=  /*line :1*/ errors.Su6g6hRoi9X("unexpected key value type")
)

func (w_9IiB2_ajp WA26dh) GetValue(ihk6jlqLI string, zwMp0T []byte) (dEntjI5C8pQ int, aehKEp uint32, vbFroa5 error) {
	xEzVgFhb5j, vbFroa5 :=  /*line :1*/ syscall.GcOmFfsibES(ihk6jlqLI)
	if vbFroa5 != nil {
		return 0, 0, vbFroa5
	}
	var iqVitYnNbTP *byte
	if  /*line :1*/ len(zwMp0T) > 0 {
		iqVitYnNbTP = (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&zwMp0T[0]))
	}
	e1gfW0 :=  /*line :1*/ uint32( /*line :1*/ len(zwMp0T))
	vbFroa5 =  /*line :1*/ syscall.BcZh0IGWcpB( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), xEzVgFhb5j, nil, &aehKEp, iqVitYnNbTP, &e1gfW0)
	if vbFroa5 != nil {
		return  /*line :1*/ int(e1gfW0), aehKEp, vbFroa5
	}
	return  /*line :1*/ int(e1gfW0), aehKEp, nil
}

func (w_9IiB2_ajp WA26dh) faFtKxIzV(ihk6jlqLI string, zwMp0T []byte) (pffwF1iH []byte, aehKEp uint32, vbFroa5 error) {
	pWlH3HKw, vbFroa5 :=  /*line :1*/ syscall.GcOmFfsibES(ihk6jlqLI)
	if vbFroa5 != nil {
		return nil, 0, vbFroa5
	}
	var goNtyW uint32
	dEntjI5C8pQ :=  /*line :1*/ uint32( /*line :1*/ len(zwMp0T))
	for {
		vbFroa5 =  /*line :1*/ syscall.BcZh0IGWcpB( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), pWlH3HKw, nil, &goNtyW, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&zwMp0T[0])), &dEntjI5C8pQ)
		if vbFroa5 == nil {
			return zwMp0T[:dEntjI5C8pQ], goNtyW, nil
		}
		if vbFroa5 != syscall.ERROR_MORE_DATA {
			return nil, 0, vbFroa5
		}
		if dEntjI5C8pQ <=  /*line :1*/ uint32( /*line :1*/ len(zwMp0T)) {
			return nil, 0, vbFroa5
		}
		zwMp0T =  /*line :1*/ make([]byte, dEntjI5C8pQ)
	}
}

func (w_9IiB2_ajp WA26dh) GetStringValue(ihk6jlqLI string) (vv1B81Aurzk string, aehKEp uint32, vbFroa5 error) {
	iqP0QfA7i, q7unmU, xt0DDS :=  /*line :1*/ w_9IiB2_ajp.faFtKxIzV(ihk6jlqLI,  /*line :1*/ make([]byte, 64))
	if xt0DDS != nil {
		return "", q7unmU, xt0DDS
	}
	switch q7unmU {
	case SZ, EXPAND_SZ:
	default:
		return "", q7unmU, U12T7G
	}
	if  /*line :1*/ len(iqP0QfA7i) == 0 {
		return "", q7unmU, nil
	}
	oxDVAJ_K := (*[1 << 29] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&iqP0QfA7i[0]))[:  /*line :1*/ len(iqP0QfA7i)/2 :  /*line :1*/ len(iqP0QfA7i)/2]
	return  /*line :1*/ syscall.AODVXp8NM3sd(oxDVAJ_K), q7unmU, nil
}

func (w_9IiB2_ajp WA26dh) GetMUIStringValue(ihk6jlqLI string) (string, error) {
	xEzVgFhb5j, vbFroa5 :=  /*line :1*/ syscall.GcOmFfsibES(ihk6jlqLI)
	if vbFroa5 != nil {
		return "", vbFroa5
	}

	zwMp0T :=  /*line :1*/ make([]uint16, 1024)
	var xmDfVl5oS uint32
	var vGOQSC7kCUY *uint16

	vbFroa5 =  /*line :1*/ u3OxV2gI_( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), xEzVgFhb5j, &zwMp0T[0],  /*line :1*/ uint32( /*line :1*/ len(zwMp0T)), &xmDfVl5oS, 0, vGOQSC7kCUY)
	if vbFroa5 == syscall.ERROR_FILE_NOT_FOUND {

		var eei5Lp3X0NY string
		eei5Lp3X0NY, vbFroa5 =  /*line :1*/ Unaa57ohiv22("%SystemRoot%\\system32\\")
		if vbFroa5 != nil {
			return "", vbFroa5
		}
		vGOQSC7kCUY, vbFroa5 =  /*line :1*/ syscall.GcOmFfsibES(eei5Lp3X0NY)
		if vbFroa5 != nil {
			return "", vbFroa5
		}

		vbFroa5 =  /*line :1*/ u3OxV2gI_( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), xEzVgFhb5j, &zwMp0T[0],  /*line :1*/ uint32( /*line :1*/ len(zwMp0T)), &xmDfVl5oS, 0, vGOQSC7kCUY)
	}

	for vbFroa5 == syscall.ERROR_MORE_DATA {
		if xmDfVl5oS <=  /*line :1*/ uint32( /*line :1*/ len(zwMp0T)) {
			break
		}
		zwMp0T =  /*line :1*/ make([]uint16, xmDfVl5oS)
		vbFroa5 =  /*line :1*/ u3OxV2gI_( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), xEzVgFhb5j, &zwMp0T[0],  /*line :1*/ uint32( /*line :1*/ len(zwMp0T)), &xmDfVl5oS, 0, vGOQSC7kCUY)
	}

	if vbFroa5 != nil {
		return "", vbFroa5
	}

	return  /*line :1*/ syscall.AODVXp8NM3sd(zwMp0T), nil
}

func Unaa57ohiv22(xwyhoxHsnwld string) (string, error) {
	if xwyhoxHsnwld == "" {
		return "", nil
	}
	pWlH3HKw, vbFroa5 :=  /*line :1*/ syscall.GcOmFfsibES(xwyhoxHsnwld)
	if vbFroa5 != nil {
		return "", vbFroa5
	}
	hbokWzI :=  /*line :1*/ make([]uint16, 100)
	for {
		dEntjI5C8pQ, vbFroa5 :=  /*line :1*/ zEpSOSa6(pWlH3HKw, &hbokWzI[0],  /*line :1*/ uint32( /*line :1*/ len(hbokWzI)))
		if vbFroa5 != nil {
			return "", vbFroa5
		}
		if dEntjI5C8pQ <=  /*line :1*/ uint32( /*line :1*/ len(hbokWzI)) {
			return  /*line :1*/ syscall.AODVXp8NM3sd(hbokWzI[:dEntjI5C8pQ]), nil
		}
		hbokWzI =  /*line :1*/ make([]uint16, dEntjI5C8pQ)
	}
}

func (w_9IiB2_ajp WA26dh) GetStringsValue(ihk6jlqLI string) (vv1B81Aurzk []string, aehKEp uint32, vbFroa5 error) {
	iqP0QfA7i, q7unmU, xt0DDS :=  /*line :1*/ w_9IiB2_ajp.faFtKxIzV(ihk6jlqLI,  /*line :1*/ make([]byte, 64))
	if xt0DDS != nil {
		return nil, q7unmU, xt0DDS
	}
	if q7unmU != MULTI_SZ {
		return nil, q7unmU, U12T7G
	}
	if  /*line :1*/ len(iqP0QfA7i) == 0 {
		return nil, q7unmU, nil
	}
	pWlH3HKw := (*[1 << 29] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&iqP0QfA7i[0]))[:  /*line :1*/ len(iqP0QfA7i)/2 :  /*line :1*/ len(iqP0QfA7i)/2]
	if  /*line :1*/ len(pWlH3HKw) == 0 {
		return nil, q7unmU, nil
	}
	if pWlH3HKw[ /*line :1*/ len(pWlH3HKw)-1] == 0 {
		pWlH3HKw = pWlH3HKw[: /*line :1*/ len(pWlH3HKw)-1]
	}
	vv1B81Aurzk =  /*line :1*/ make([]string, 0, 5)
	spJs12iLg := 0
	for rkWqLgU, nJut9BWeQL4T := range pWlH3HKw {
		if nJut9BWeQL4T == 0 {
			vv1B81Aurzk =  /*line :1*/ append(vv1B81Aurzk,  /*line :1*/ syscall.AODVXp8NM3sd(pWlH3HKw[spJs12iLg:rkWqLgU]))
			spJs12iLg = rkWqLgU + 1
		}
	}
	return vv1B81Aurzk, q7unmU, nil
}

func (w_9IiB2_ajp WA26dh) GetIntegerValue(ihk6jlqLI string) (vv1B81Aurzk uint64, aehKEp uint32, vbFroa5 error) {
	iqP0QfA7i, q7unmU, xt0DDS :=  /*line :1*/ w_9IiB2_ajp.faFtKxIzV(ihk6jlqLI,  /*line :1*/ make([]byte, 8))
	if xt0DDS != nil {
		return 0, q7unmU, xt0DDS
	}
	switch q7unmU {
	case DWORD:
		if  /*line :1*/ len(iqP0QfA7i) != 4 {
			return 0, q7unmU,  /*line :1*/ errors.Su6g6hRoi9X("DWORD value is not 4 bytes long")
		}
		return  /*line :1*/ uint64(*(* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(&iqP0QfA7i[0]))), DWORD, nil
	case QWORD:
		if  /*line :1*/ len(iqP0QfA7i) != 8 {
			return 0, q7unmU,  /*line :1*/ errors.Su6g6hRoi9X("QWORD value is not 8 bytes long")
		}
		return  /*line :1*/ uint64(*(* /*line :1*/ uint64)( /*line :1*/ unsafe.Pointer(&iqP0QfA7i[0]))), QWORD, nil
	default:
		return 0, q7unmU, U12T7G
	}
}

func (w_9IiB2_ajp WA26dh) GetBinaryValue(ihk6jlqLI string) (vv1B81Aurzk []byte, aehKEp uint32, vbFroa5 error) {
	iqP0QfA7i, q7unmU, xt0DDS :=  /*line :1*/ w_9IiB2_ajp.faFtKxIzV(ihk6jlqLI,  /*line :1*/ make([]byte, 64))
	if xt0DDS != nil {
		return nil, q7unmU, xt0DDS
	}
	if q7unmU != BINARY {
		return nil, q7unmU, U12T7G
	}
	return iqP0QfA7i, q7unmU, nil
}

func (w_9IiB2_ajp WA26dh) dN9vscXqp3l(ihk6jlqLI string, aehKEp uint32, iqP0QfA7i []byte) error {
	pWlH3HKw, vbFroa5 :=  /*line :1*/ syscall.GcOmFfsibES(ihk6jlqLI)
	if vbFroa5 != nil {
		return vbFroa5
	}
	if  /*line :1*/ len(iqP0QfA7i) == 0 {
		return  /*line :1*/ wIasC9b( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), pWlH3HKw, 0, aehKEp, nil, 0)
	}
	return  /*line :1*/ wIasC9b( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), pWlH3HKw, 0, aehKEp, &iqP0QfA7i[0],  /*line :1*/ uint32( /*line :1*/ len(iqP0QfA7i)))
}

func (w_9IiB2_ajp WA26dh) SetDWordValue(ihk6jlqLI string, xwyhoxHsnwld uint32) error {
	return  /*line :1*/ w_9IiB2_ajp.dN9vscXqp3l(ihk6jlqLI, DWORD, (*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&xwyhoxHsnwld))[:])
}

func (w_9IiB2_ajp WA26dh) SetQWordValue(ihk6jlqLI string, xwyhoxHsnwld uint64) error {
	return  /*line :1*/ w_9IiB2_ajp.dN9vscXqp3l(ihk6jlqLI, QWORD, (*[8] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&xwyhoxHsnwld))[:])
}

func (w_9IiB2_ajp WA26dh) hJFJRlT4xkdo(ihk6jlqLI string, aehKEp uint32, xwyhoxHsnwld string) error {
	acf815aA5lD3, vbFroa5 :=  /*line :1*/ syscall.ZxIMCrnjQ5U4(xwyhoxHsnwld)
	if vbFroa5 != nil {
		return vbFroa5
	}
	zwMp0T := (*[1 << 29] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&acf815aA5lD3[0]))[:  /*line :1*/ len(acf815aA5lD3)*2 :  /*line :1*/ len(acf815aA5lD3)*2]
	return  /*line :1*/ w_9IiB2_ajp.dN9vscXqp3l(ihk6jlqLI, aehKEp, zwMp0T)
}

func (w_9IiB2_ajp WA26dh) SetStringValue(ihk6jlqLI, xwyhoxHsnwld string) error {
	return  /*line :1*/ w_9IiB2_ajp.hJFJRlT4xkdo(ihk6jlqLI, SZ, xwyhoxHsnwld)
}

func (w_9IiB2_ajp WA26dh) SetExpandStringValue(ihk6jlqLI, xwyhoxHsnwld string) error {
	return  /*line :1*/ w_9IiB2_ajp.hJFJRlT4xkdo(ihk6jlqLI, EXPAND_SZ, xwyhoxHsnwld)
}

func (w_9IiB2_ajp WA26dh) SetStringsValue(ihk6jlqLI string, xwyhoxHsnwld []string) error {
	aDOojvceKS := ""
	for _, eei5Lp3X0NY := range xwyhoxHsnwld {
		for rkWqLgU := 0; rkWqLgU <  /*line :1*/ len(eei5Lp3X0NY); rkWqLgU++ {
			if eei5Lp3X0NY[rkWqLgU] == 0 {
				return  /*line :1*/ errors.Su6g6hRoi9X("string cannot have 0 inside")
			}
		}
		aDOojvceKS += eei5Lp3X0NY + "\x00"
	}
	acf815aA5lD3 :=  /*line :1*/ utf16.J8eAwOfFc([] /*line :1*/ rune(aDOojvceKS + "\x00"))
	zwMp0T := (*[1 << 29] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&acf815aA5lD3[0]))[:  /*line :1*/ len(acf815aA5lD3)*2 :  /*line :1*/ len(acf815aA5lD3)*2]
	return  /*line :1*/ w_9IiB2_ajp.dN9vscXqp3l(ihk6jlqLI, MULTI_SZ, zwMp0T)
}

func (w_9IiB2_ajp WA26dh) SetBinaryValue(ihk6jlqLI string, xwyhoxHsnwld []byte) error {
	return  /*line :1*/ w_9IiB2_ajp.dN9vscXqp3l(ihk6jlqLI, BINARY, xwyhoxHsnwld)
}

func (w_9IiB2_ajp WA26dh) DeleteValue(ihk6jlqLI string) error {
	return  /*line :1*/ toNASzC( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp),  /*line :1*/ syscall.EtPVNA(ihk6jlqLI))
}

func (w_9IiB2_ajp WA26dh) ReadValueNames() ([]string, error) {
	citJdGKUXJrx, vbFroa5 :=  /*line :1*/ w_9IiB2_ajp.Stat()
	if vbFroa5 != nil {
		return nil, vbFroa5
	}
	bQfzy18s2t :=  /*line :1*/ make([]string, 0, citJdGKUXJrx.BxiaYE_WgOj)
	zwMp0T :=  /*line :1*/ make([]uint16, citJdGKUXJrx.Gx6pDtaYPP+1)
loopItems:
	for rkWqLgU :=  /*line :1*/ uint32(0); ; rkWqLgU++ {
		e1gfW0 :=  /*line :1*/ uint32( /*line :1*/ len(zwMp0T))
		for {
			vbFroa5 :=  /*line :1*/ hzM0h_( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), rkWqLgU, &zwMp0T[0], &e1gfW0, nil, nil, nil, nil)
			if vbFroa5 == nil {
				break
			}
			if vbFroa5 == syscall.ERROR_MORE_DATA {

				e1gfW0 =  /*line :1*/ uint32(2 *  /*line :1*/ len(zwMp0T))
				zwMp0T =  /*line :1*/ make([]uint16, e1gfW0)
				continue
			}
			if vbFroa5 == _ERROR_NO_MORE_ITEMS {
				break loopItems
			}
			return bQfzy18s2t, vbFroa5
		}
		bQfzy18s2t =  /*line :1*/ append(bQfzy18s2t,  /*line :1*/ syscall.AODVXp8NM3sd(zwMp0T[:e1gfW0]))
	}
	return bQfzy18s2t, nil
}
