//line :1
//go:build windows

package ER_Q0ps1VO

import (
	io "xI9ai2djaJ2"
	"runtime"
	syscall "bUKeamF"
	time "fRAfQd_"
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

type VcZo9w syscall.SRlvVjrYa

const (
	CLASSES_ROOT	=  /*line :1*/ VcZo9w(syscall.HKEY_CLASSES_ROOT)
	CURRENT_USER	=  /*line :1*/ VcZo9w(syscall.HKEY_CURRENT_USER)
	LOCAL_MACHINE	=  /*line :1*/ VcZo9w(syscall.HKEY_LOCAL_MACHINE)
	USERS	=  /*line :1*/ VcZo9w(syscall.HKEY_USERS)
	CURRENT_CONFIG	=  /*line :1*/ VcZo9w(syscall.HKEY_CURRENT_CONFIG)
	PERFORMANCE_DATA	=  /*line :1*/ VcZo9w(syscall.HKEY_PERFORMANCE_DATA)
)

func (e738WnX VcZo9w) Close() error {
	return  /*line :1*/ syscall.BqIBx3KRN4( /*line :1*/ syscall.SRlvVjrYa(e738WnX))
}

func IMz89PkYACx(e738WnX VcZo9w, l6_w95LlMmR string, pHsdyK9i uint32) (VcZo9w, error) {
	hkmoUHCT, konG16p :=  /*line :1*/ syscall.GcOmFfsibES(l6_w95LlMmR)
	if konG16p != nil {
		return 0, konG16p
	}
	var xW2K3by syscall.SRlvVjrYa
	konG16p =  /*line :1*/ syscall.FRt5iWSq2fWU( /*line :1*/ syscall.SRlvVjrYa(e738WnX), hkmoUHCT, 0, pHsdyK9i, &xW2K3by)
	if konG16p != nil {
		return 0, konG16p
	}
	return  /*line :1*/ VcZo9w(xW2K3by), nil
}

func Xx2A29xaVaD(xge7RLq string, e738WnX VcZo9w) (VcZo9w, error) {
	var konG16p error
	var hkmoUHCT *uint16
	if xge7RLq != "" {
		hkmoUHCT, konG16p =  /*line :1*/ syscall.GcOmFfsibES(`\\` + xge7RLq)
		if konG16p != nil {
			return 0, konG16p
		}
	}
	var xTL9_xLQ syscall.SRlvVjrYa
	konG16p =  /*line :1*/ zW6WK191u(hkmoUHCT,  /*line :1*/ syscall.SRlvVjrYa(e738WnX), &xTL9_xLQ)
	if konG16p != nil {
		return 0, konG16p
	}
	return  /*line :1*/ VcZo9w(xTL9_xLQ), nil
}

func (e738WnX VcZo9w) ReadSubKeyNames(aMs6DlZl4 int) ([]string, error) {

	 /*line :1*/ runtime.LockOSThread()
	defer  /*line :1*/ runtime.UnlockOSThread()

	s4K7gQdEoy :=  /*line :1*/ make([]string, 0)

	_753m9TV2 :=  /*line :1*/ make([]uint16, 256)
loopItems:
	for cxU1l31 :=  /*line :1*/ uint32(0); ; cxU1l31++ {
		if aMs6DlZl4 > 0 {
			if  /*line :1*/ len(s4K7gQdEoy) == aMs6DlZl4 {
				return s4K7gQdEoy, nil
			}
		}
		bmmagfcma :=  /*line :1*/ uint32( /*line :1*/ len(_753m9TV2))
		for {
			konG16p :=  /*line :1*/ syscall.T4R52_D4B1Z( /*line :1*/ syscall.SRlvVjrYa(e738WnX), cxU1l31, &_753m9TV2[0], &bmmagfcma, nil, nil, nil, nil)
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

func NBPQa9aLn5ju(e738WnX VcZo9w, l6_w95LlMmR string, pHsdyK9i uint32) (djd_jqSp VcZo9w, rQsD2Fli bool, konG16p error) {
	var cxBIPdbqQ syscall.SRlvVjrYa
	var cx4PJ4H uint32
	konG16p =  /*line :1*/ ghyTj4Lur( /*line :1*/ syscall.SRlvVjrYa(e738WnX),  /*line :1*/ syscall.EtPVNA(l6_w95LlMmR),
		0, nil, _REG_OPTION_NON_VOLATILE, pHsdyK9i, nil, &cxBIPdbqQ, &cx4PJ4H)
	if konG16p != nil {
		return 0, false, konG16p
	}
	return  /*line :1*/ VcZo9w(cxBIPdbqQ), cx4PJ4H == _REG_OPENED_EXISTING_KEY, nil
}

func HfaVxtu(e738WnX VcZo9w, l6_w95LlMmR string) error {
	return  /*line :1*/ oBIbRrud( /*line :1*/ syscall.SRlvVjrYa(e738WnX),  /*line :1*/ syscall.EtPVNA(l6_w95LlMmR))
}

type UkhWGat59 struct {
	Inor9M8xT6	uint32
	X5JZ2zu9	uint32
	BxiaYE_WgOj	uint32
	Gx6pDtaYPP	uint32
	Ab8aiDizEt	uint32
	hbPbRa4r6z	syscall.T8WbUqAC3v
}

func (gBO9LAb *UkhWGat59) ModTime() time.G4KDkI {
	return  /*line :1*/ time.FRXtx9QnU(0,  /*line :1*/ gBO9LAb.hbPbRa4r6z.Nanoseconds())
}

func (e738WnX VcZo9w) Stat() (*UkhWGat59, error) {
	var gBO9LAb UkhWGat59
	konG16p :=  /*line :1*/ syscall.N2CMOGAc( /*line :1*/ syscall.SRlvVjrYa(e738WnX), nil, nil, nil,
		&gBO9LAb.Inor9M8xT6, &gBO9LAb.X5JZ2zu9, nil, &gBO9LAb.BxiaYE_WgOj,
		&gBO9LAb.Gx6pDtaYPP, &gBO9LAb.Ab8aiDizEt, nil, &gBO9LAb.hbPbRa4r6z)
	if konG16p != nil {
		return nil, konG16p
	}
	return &gBO9LAb, nil
}
