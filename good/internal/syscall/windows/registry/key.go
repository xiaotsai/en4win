//line :1
//go:build windows

package GMNBnK

import (
	"runtime"
	syscall "bUKeamF"
)

const (
	ALL_ACCESS	= 0xf003f
	CREATE_LINK	= 0x00020
	CREATE_SUB_KEY	= 0x00004
	ENUMERATE_SUB_KEYS	= 0x00008
	EXECUTE	= 0x20019
	NOTIFY	= 0x00010
	QUERY_VALUE	= 0x00001
	READ	= 0x20019
	SET_VALUE	= 0x00002
	WOW64_32KEY	= 0x00200
	WOW64_64KEY	= 0x00100
	WRITE	= 0x20006
)

type WA26dh syscall.SRlvVjrYa

const (
	CLASSES_ROOT	=  /*line :1*/ WA26dh(syscall.HKEY_CLASSES_ROOT)
	CURRENT_USER	=  /*line :1*/ WA26dh(syscall.HKEY_CURRENT_USER)
	LOCAL_MACHINE	=  /*line :1*/ WA26dh(syscall.HKEY_LOCAL_MACHINE)
	USERS	=  /*line :1*/ WA26dh(syscall.HKEY_USERS)
	CURRENT_CONFIG	=  /*line :1*/ WA26dh(syscall.HKEY_CURRENT_CONFIG)
)

func (w_9IiB2_ajp WA26dh) Close() error {
	return  /*line :1*/ syscall.BqIBx3KRN4( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp))
}

func J2fxYepRwi3b(w_9IiB2_ajp WA26dh, g6yvSMh2pd string, t5h12CE8Dy6 uint32) (WA26dh, error) {
	pWlH3HKw, vbFroa5 :=  /*line :1*/ syscall.GcOmFfsibES(g6yvSMh2pd)
	if vbFroa5 != nil {
		return 0, vbFroa5
	}
	var wNnx4v syscall.SRlvVjrYa
	vbFroa5 =  /*line :1*/ syscall.FRt5iWSq2fWU( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), pWlH3HKw, 0, t5h12CE8Dy6, &wNnx4v)
	if vbFroa5 != nil {
		return 0, vbFroa5
	}
	return  /*line :1*/ WA26dh(wNnx4v), nil
}

func (w_9IiB2_ajp WA26dh) ReadSubKeyNames() ([]string, error) {

	 /*line :1*/ runtime.LockOSThread()
	defer  /*line :1*/ runtime.UnlockOSThread()

	bQfzy18s2t :=  /*line :1*/ make([]string, 0)

	zwMp0T :=  /*line :1*/ make([]uint16, 256)
loopItems:
	for rkWqLgU :=  /*line :1*/ uint32(0); ; rkWqLgU++ {
		e1gfW0 :=  /*line :1*/ uint32( /*line :1*/ len(zwMp0T))
		for {
			vbFroa5 :=  /*line :1*/ syscall.T4R52_D4B1Z( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), rkWqLgU, &zwMp0T[0], &e1gfW0, nil, nil, nil, nil)
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

func OW_uC5(w_9IiB2_ajp WA26dh, g6yvSMh2pd string, t5h12CE8Dy6 uint32) (sshgXkysnZ1 WA26dh, w95FOyqZ bool, vbFroa5 error) {
	var w1aPQ2 syscall.SRlvVjrYa
	var aIPWnMJq uint32
	vbFroa5 =  /*line :1*/ jQ9Dwq597( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp),  /*line :1*/ syscall.EtPVNA(g6yvSMh2pd),
		0, nil, _REG_OPTION_NON_VOLATILE, t5h12CE8Dy6, nil, &w1aPQ2, &aIPWnMJq)
	if vbFroa5 != nil {
		return 0, false, vbFroa5
	}
	return  /*line :1*/ WA26dh(w1aPQ2), aIPWnMJq == _REG_OPENED_EXISTING_KEY, nil
}

func PKaHoylif(w_9IiB2_ajp WA26dh, g6yvSMh2pd string) error {
	return  /*line :1*/ d3PMbL( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp),  /*line :1*/ syscall.EtPVNA(g6yvSMh2pd))
}

type IhGJrhYF8 struct {
	Inor9M8xT6	uint32
	X5JZ2zu9	uint32
	BxiaYE_WgOj	uint32
	Gx6pDtaYPP	uint32
	Ab8aiDizEt	uint32
	hbPbRa4r6z	syscall.T8WbUqAC3v
}

func (w_9IiB2_ajp WA26dh) Stat() (*IhGJrhYF8, error) {
	var citJdGKUXJrx IhGJrhYF8
	vbFroa5 :=  /*line :1*/ syscall.N2CMOGAc( /*line :1*/ syscall.SRlvVjrYa(w_9IiB2_ajp), nil, nil, nil,
		&citJdGKUXJrx.Inor9M8xT6, &citJdGKUXJrx.X5JZ2zu9, nil, &citJdGKUXJrx.BxiaYE_WgOj,
		&citJdGKUXJrx.Gx6pDtaYPP, &citJdGKUXJrx.Ab8aiDizEt, nil, &citJdGKUXJrx.hbPbRa4r6z)
	if vbFroa5 != nil {
		return nil, vbFroa5
	}
	return &citJdGKUXJrx, nil
}
