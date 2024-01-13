//line :1
//go:build windows

package ER_Q0ps1VO

import (
	errors "iAHoxjmM"
	io "xI9ai2djaJ2"
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
	M1oSxSEW	= syscall.ERROR_MORE_DATA

	Iz4uaNZz	= syscall.ERROR_FILE_NOT_FOUND

	TW30gTcpEuaD	=  /*line :1*/ errors.Su6g6hRoi9X("unexpected key value type")
)

func (e738WnX VcZo9w) GetValue(jy58GAHux string, _753m9TV2 []byte) (aMs6DlZl4 int, xN8Ac8a7xRi uint32, konG16p error) {
	jKennE, konG16p :=  /*line :1*/ syscall.GcOmFfsibES(jy58GAHux)
	if konG16p != nil {
		return 0, 0, konG16p
	}
	var zUMLDDl3U *byte
	if  /*line :1*/ len(_753m9TV2) > 0 {
		zUMLDDl3U = (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&_753m9TV2[0]))
	}
	bmmagfcma :=  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2))
	konG16p =  /*line :1*/ syscall.BcZh0IGWcpB( /*line :1*/ syscall.SRlvVjrYa(e738WnX), jKennE, nil, &xN8Ac8a7xRi, zUMLDDl3U, &bmmagfcma)
	if konG16p != nil {
		return  /*line :1*/ int(bmmagfcma), xN8Ac8a7xRi, konG16p
	}
	return  /*line :1*/ int(bmmagfcma), xN8Ac8a7xRi, nil
}

func (e738WnX VcZo9w) c9sLCM(jy58GAHux string, _753m9TV2 []byte) (pDAlIVKb []byte, xN8Ac8a7xRi uint32, konG16p error) {
	hkmoUHCT, konG16p :=  /*line :1*/ syscall.GcOmFfsibES(jy58GAHux)
	if konG16p != nil {
		return nil, 0, konG16p
	}
	var gjhO8Rs uint32
	aMs6DlZl4 :=  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2))
	for {
		konG16p =  /*line :1*/ syscall.BcZh0IGWcpB( /*line :1*/ syscall.SRlvVjrYa(e738WnX), hkmoUHCT, nil, &gjhO8Rs, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&_753m9TV2[0])), &aMs6DlZl4)
		if konG16p == nil {
			return _753m9TV2[:aMs6DlZl4], gjhO8Rs, nil
		}
		if konG16p != syscall.ERROR_MORE_DATA {
			return nil, 0, konG16p
		}
		if aMs6DlZl4 <=  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2)) {
			return nil, 0, konG16p
		}
		_753m9TV2 =  /*line :1*/ make([]byte, aMs6DlZl4)
	}
}

func (e738WnX VcZo9w) GetStringValue(jy58GAHux string) (vGyJDj9rEo_g string, xN8Ac8a7xRi uint32, konG16p error) {
	pDAlIVKb, gX7gelw, sYe4OOCYG :=  /*line :1*/ e738WnX.c9sLCM(jy58GAHux,  /*line :1*/ make([]byte, 64))
	if sYe4OOCYG != nil {
		return "", gX7gelw, sYe4OOCYG
	}
	switch gX7gelw {
	case SZ, EXPAND_SZ:
	default:
		return "", gX7gelw, TW30gTcpEuaD
	}
	if  /*line :1*/ len(pDAlIVKb) == 0 {
		return "", gX7gelw, nil
	}
	wI2Zbi := (*[1 << 29] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&pDAlIVKb[0]))[:  /*line :1*/ len(pDAlIVKb)/2 :  /*line :1*/ len(pDAlIVKb)/2]
	return  /*line :1*/ syscall.AODVXp8NM3sd(wI2Zbi), gX7gelw, nil
}

func (e738WnX VcZo9w) GetMUIStringValue(jy58GAHux string) (string, error) {
	jKennE, konG16p :=  /*line :1*/ syscall.GcOmFfsibES(jy58GAHux)
	if konG16p != nil {
		return "", konG16p
	}

	_753m9TV2 :=  /*line :1*/ make([]uint16, 1024)
	var e5ZJCFVvq0UN uint32
	var qe8tVf9ohd *uint16

	konG16p =  /*line :1*/ gByj77H6i( /*line :1*/ syscall.SRlvVjrYa(e738WnX), jKennE, &_753m9TV2[0],  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2)), &e5ZJCFVvq0UN, 0, qe8tVf9ohd)
	if konG16p == syscall.ERROR_FILE_NOT_FOUND {

		var eDDa0W7aze string
		eDDa0W7aze, konG16p =  /*line :1*/ WVtvL_a("%SystemRoot%\\system32\\")
		if konG16p != nil {
			return "", konG16p
		}
		qe8tVf9ohd, konG16p =  /*line :1*/ syscall.GcOmFfsibES(eDDa0W7aze)
		if konG16p != nil {
			return "", konG16p
		}

		konG16p =  /*line :1*/ gByj77H6i( /*line :1*/ syscall.SRlvVjrYa(e738WnX), jKennE, &_753m9TV2[0],  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2)), &e5ZJCFVvq0UN, 0, qe8tVf9ohd)
	}

	for konG16p == syscall.ERROR_MORE_DATA {
		if e5ZJCFVvq0UN <=  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2)) {
			break
		}
		_753m9TV2 =  /*line :1*/ make([]uint16, e5ZJCFVvq0UN)
		konG16p =  /*line :1*/ gByj77H6i( /*line :1*/ syscall.SRlvVjrYa(e738WnX), jKennE, &_753m9TV2[0],  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2)), &e5ZJCFVvq0UN, 0, qe8tVf9ohd)
	}

	if konG16p != nil {
		return "", konG16p
	}

	return  /*line :1*/ syscall.AODVXp8NM3sd(_753m9TV2), nil
}

func WVtvL_a(zDLhbcYGOZm4 string) (string, error) {
	if zDLhbcYGOZm4 == "" {
		return "", nil
	}
	hkmoUHCT, konG16p :=  /*line :1*/ syscall.GcOmFfsibES(zDLhbcYGOZm4)
	if konG16p != nil {
		return "", konG16p
	}
	iCtQWle :=  /*line :1*/ make([]uint16, 100)
	for {
		aMs6DlZl4, konG16p :=  /*line :1*/ eQejQjwX(hkmoUHCT, &iCtQWle[0],  /*line :1*/ uint32( /*line :1*/ len(iCtQWle)))
		if konG16p != nil {
			return "", konG16p
		}
		if aMs6DlZl4 <=  /*line :1*/ uint32( /*line :1*/ len(iCtQWle)) {
			return  /*line :1*/ syscall.AODVXp8NM3sd(iCtQWle[:aMs6DlZl4]), nil
		}
		iCtQWle =  /*line :1*/ make([]uint16, aMs6DlZl4)
	}
}

func (e738WnX VcZo9w) GetStringsValue(jy58GAHux string) (vGyJDj9rEo_g []string, xN8Ac8a7xRi uint32, konG16p error) {
	pDAlIVKb, gX7gelw, sYe4OOCYG :=  /*line :1*/ e738WnX.c9sLCM(jy58GAHux,  /*line :1*/ make([]byte, 64))
	if sYe4OOCYG != nil {
		return nil, gX7gelw, sYe4OOCYG
	}
	if gX7gelw != MULTI_SZ {
		return nil, gX7gelw, TW30gTcpEuaD
	}
	if  /*line :1*/ len(pDAlIVKb) == 0 {
		return nil, gX7gelw, nil
	}
	hkmoUHCT := (*[1 << 29] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&pDAlIVKb[0]))[:  /*line :1*/ len(pDAlIVKb)/2 :  /*line :1*/ len(pDAlIVKb)/2]
	if  /*line :1*/ len(hkmoUHCT) == 0 {
		return nil, gX7gelw, nil
	}
	if hkmoUHCT[ /*line :1*/ len(hkmoUHCT)-1] == 0 {
		hkmoUHCT = hkmoUHCT[: /*line :1*/ len(hkmoUHCT)-1]
	}
	vGyJDj9rEo_g =  /*line :1*/ make([]string, 0, 5)
	jV2_SqP := 0
	for cxU1l31, abY2f8bCDGd := range hkmoUHCT {
		if abY2f8bCDGd == 0 {
			vGyJDj9rEo_g =  /*line :1*/ append(vGyJDj9rEo_g,  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(hkmoUHCT[jV2_SqP:cxU1l31])))
			jV2_SqP = cxU1l31 + 1
		}
	}
	return vGyJDj9rEo_g, gX7gelw, nil
}

func (e738WnX VcZo9w) GetIntegerValue(jy58GAHux string) (vGyJDj9rEo_g uint64, xN8Ac8a7xRi uint32, konG16p error) {
	pDAlIVKb, gX7gelw, sYe4OOCYG :=  /*line :1*/ e738WnX.c9sLCM(jy58GAHux,  /*line :1*/ make([]byte, 8))
	if sYe4OOCYG != nil {
		return 0, gX7gelw, sYe4OOCYG
	}
	switch gX7gelw {
	case DWORD:
		if  /*line :1*/ len(pDAlIVKb) != 4 {
			return 0, gX7gelw,  /*line :1*/ errors.Su6g6hRoi9X("DWORD value is not 4 bytes long")
		}
		var eza9zx0Zn9uE uint32
		 /*line :1*/ copy((*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&eza9zx0Zn9uE))[:], pDAlIVKb)
		return  /*line :1*/ uint64(eza9zx0Zn9uE), DWORD, nil
	case QWORD:
		if  /*line :1*/ len(pDAlIVKb) != 8 {
			return 0, gX7gelw,  /*line :1*/ errors.Su6g6hRoi9X("QWORD value is not 8 bytes long")
		}
		 /*line :1*/ copy((*[8] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&vGyJDj9rEo_g))[:], pDAlIVKb)
		return vGyJDj9rEo_g, QWORD, nil
	default:
		return 0, gX7gelw, TW30gTcpEuaD
	}
}

func (e738WnX VcZo9w) GetBinaryValue(jy58GAHux string) (vGyJDj9rEo_g []byte, xN8Ac8a7xRi uint32, konG16p error) {
	pDAlIVKb, gX7gelw, sYe4OOCYG :=  /*line :1*/ e738WnX.c9sLCM(jy58GAHux,  /*line :1*/ make([]byte, 64))
	if sYe4OOCYG != nil {
		return nil, gX7gelw, sYe4OOCYG
	}
	if gX7gelw != BINARY {
		return nil, gX7gelw, TW30gTcpEuaD
	}
	return pDAlIVKb, gX7gelw, nil
}

func (e738WnX VcZo9w) sAx1B0CbZ(jy58GAHux string, xN8Ac8a7xRi uint32, pDAlIVKb []byte) error {
	hkmoUHCT, konG16p :=  /*line :1*/ syscall.GcOmFfsibES(jy58GAHux)
	if konG16p != nil {
		return konG16p
	}
	if  /*line :1*/ len(pDAlIVKb) == 0 {
		return  /*line :1*/ fBzuLbgu( /*line :1*/ syscall.SRlvVjrYa(e738WnX), hkmoUHCT, 0, xN8Ac8a7xRi, nil, 0)
	}
	return  /*line :1*/ fBzuLbgu( /*line :1*/ syscall.SRlvVjrYa(e738WnX), hkmoUHCT, 0, xN8Ac8a7xRi, &pDAlIVKb[0],  /*line :1*/ uint32( /*line :1*/ len(pDAlIVKb)))
}

func (e738WnX VcZo9w) SetDWordValue(jy58GAHux string, zDLhbcYGOZm4 uint32) error {
	return  /*line :1*/ e738WnX.sAx1B0CbZ(jy58GAHux, DWORD, (*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&zDLhbcYGOZm4))[:])
}

func (e738WnX VcZo9w) SetQWordValue(jy58GAHux string, zDLhbcYGOZm4 uint64) error {
	return  /*line :1*/ e738WnX.sAx1B0CbZ(jy58GAHux, QWORD, (*[8] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&zDLhbcYGOZm4))[:])
}

func (e738WnX VcZo9w) almHFg(jy58GAHux string, xN8Ac8a7xRi uint32, zDLhbcYGOZm4 string) error {
	a460TXYEH1Y, konG16p :=  /*line :1*/ syscall.ZxIMCrnjQ5U4(zDLhbcYGOZm4)
	if konG16p != nil {
		return konG16p
	}
	_753m9TV2 := (*[1 << 29] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&a460TXYEH1Y[0]))[:  /*line :1*/ len(a460TXYEH1Y)*2 :  /*line :1*/ len(a460TXYEH1Y)*2]
	return  /*line :1*/ e738WnX.sAx1B0CbZ(jy58GAHux, xN8Ac8a7xRi, _753m9TV2)
}

func (e738WnX VcZo9w) SetStringValue(jy58GAHux, zDLhbcYGOZm4 string) error {
	return  /*line :1*/ e738WnX.almHFg(jy58GAHux, SZ, zDLhbcYGOZm4)
}

func (e738WnX VcZo9w) SetExpandStringValue(jy58GAHux, zDLhbcYGOZm4 string) error {
	return  /*line :1*/ e738WnX.almHFg(jy58GAHux, EXPAND_SZ, zDLhbcYGOZm4)
}

func (e738WnX VcZo9w) SetStringsValue(jy58GAHux string, zDLhbcYGOZm4 []string) error {
	lK6WVlOKKg := ""
	for _, eDDa0W7aze := range zDLhbcYGOZm4 {
		for cxU1l31 := 0; cxU1l31 <  /*line :1*/ len(eDDa0W7aze); cxU1l31++ {
			if eDDa0W7aze[cxU1l31] == 0 {
				return  /*line :1*/ errors.Su6g6hRoi9X("string cannot have 0 inside")
			}
		}
		lK6WVlOKKg += eDDa0W7aze + "\x00"
	}
	a460TXYEH1Y :=  /*line :1*/ utf16.J8eAwOfFc([] /*line :1*/ rune(lK6WVlOKKg + "\x00"))
	_753m9TV2 := (*[1 << 29] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&a460TXYEH1Y[0]))[:  /*line :1*/ len(a460TXYEH1Y)*2 :  /*line :1*/ len(a460TXYEH1Y)*2]
	return  /*line :1*/ e738WnX.sAx1B0CbZ(jy58GAHux, MULTI_SZ, _753m9TV2)
}

func (e738WnX VcZo9w) SetBinaryValue(jy58GAHux string, zDLhbcYGOZm4 []byte) error {
	return  /*line :1*/ e738WnX.sAx1B0CbZ(jy58GAHux, BINARY, zDLhbcYGOZm4)
}

func (e738WnX VcZo9w) DeleteValue(jy58GAHux string) error {
	return  /*line :1*/ zPUq_hvkPJRN( /*line :1*/ syscall.SRlvVjrYa(e738WnX),  /*line :1*/ syscall.EtPVNA(jy58GAHux))
}

func (e738WnX VcZo9w) ReadValueNames(aMs6DlZl4 int) ([]string, error) {
	gBO9LAb, konG16p :=  /*line :1*/ e738WnX.Stat()
	if konG16p != nil {
		return nil, konG16p
	}
	s4K7gQdEoy :=  /*line :1*/ make([]string, 0, gBO9LAb.BxiaYE_WgOj)
	_753m9TV2 :=  /*line :1*/ make([]uint16, gBO9LAb.Gx6pDtaYPP+1)
loopItems:
	for cxU1l31 :=  /*line :1*/ uint32(0); ; cxU1l31++ {
		if aMs6DlZl4 > 0 {
			if  /*line :1*/ len(s4K7gQdEoy) == aMs6DlZl4 {
				return s4K7gQdEoy, nil
			}
		}
		bmmagfcma :=  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2))
		for {
			konG16p :=  /*line :1*/ eaFjMXbf( /*line :1*/ syscall.SRlvVjrYa(e738WnX), cxU1l31, &_753m9TV2[0], &bmmagfcma, nil, nil, nil, nil)
			if konG16p == nil {
				break
			}
			if konG16p == syscall.ERROR_MORE_DATA {

				bmmagfcma =  /*line :1*/ uint32(2 *  /*line :1*/ len(_753m9TV2))
				_753m9TV2 =  /*line :1*/ make([]uint16, bmmagfcma)
				continue
			}
			if konG16p == _ERROR_NO_MORE_ITEMS {
				break loopItems
			}
			return s4K7gQdEoy, konG16p
		}
		s4K7gQdEoy =  /*line :1*/ append(s4K7gQdEoy,  /*line :1*/ syscall.AODVXp8NM3sd(_753m9TV2[:bmmagfcma]))
	}
	if aMs6DlZl4 >  /*line :1*/ len(s4K7gQdEoy) {
		return s4K7gQdEoy, io.K5Sqco
	}
	return s4K7gQdEoy, nil
}
