//line :1
package NJ4MerJ

import (
	sync "sync"
	atomic "sync/atomic"
	syscall "bUKeamF"
	"unsafe"
)

//go:linkname fcLFVQkhsS5 bUKeamF.tmHx2L
func fcLFVQkhsS5(ybpEJar *uint16) (iOvctVD26lA L2L8P9WaNZ, jeWMpOaQtWV W5PDg27zDC)

//go:linkname zoZPuC bUKeamF.gkF09M
func zoZPuC(iOvctVD26lA L2L8P9WaNZ, uwEbpuafGhb *uint8) (aTSUoJ40C uintptr, jeWMpOaQtWV W5PDg27zDC)

type YxRxePaAdi struct {
	EQTdxl	error
	EU0oqVev	string
	PtEa5aX7u5c	string
}

func (aYtmqRVkqc_8 *YxRxePaAdi) Error() string	{ return aYtmqRVkqc_8.PtEa5aX7u5c }

func (aYtmqRVkqc_8 *YxRxePaAdi) Unwrap() error	{ return aYtmqRVkqc_8.EQTdxl }

type MQbxA8MNEJC struct {
	O56SJQfoFV	string
	CyorTGi	L2L8P9WaNZ
}

func RQ6nOPol8B(gYwswmeUjG string) (o02Ffvb7mOT *MQbxA8MNEJC, jeWMpOaQtWV error) {
	hC026DpmQmus, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(gYwswmeUjG)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	_ITTe4, aYtmqRVkqc_8 :=  /*line :1*/ fcLFVQkhsS5(hC026DpmQmus)
	if aYtmqRVkqc_8 != 0 {
		return nil, &YxRxePaAdi{
			EQTdxl:	aYtmqRVkqc_8,
			EU0oqVev:	gYwswmeUjG,
			PtEa5aX7u5c:	"Failed to load " + gYwswmeUjG + ": " +  /*line :1*/ aYtmqRVkqc_8.Error(),
		}
	}
	yJ8xjP := &MQbxA8MNEJC{
		O56SJQfoFV:	gYwswmeUjG,
		CyorTGi:	_ITTe4,
	}
	return yJ8xjP, nil
}

func RZEzNE46eU(gYwswmeUjG string) *MQbxA8MNEJC {
	yJ8xjP, aYtmqRVkqc_8 :=  /*line :1*/ RQ6nOPol8B(gYwswmeUjG)
	if aYtmqRVkqc_8 != nil {
		 /*line :1*/ panic(aYtmqRVkqc_8)
	}
	return yJ8xjP
}

func (yJ8xjP *MQbxA8MNEJC) FindProc(gYwswmeUjG string) (aTSUoJ40C *U4jPAQxP8l, jeWMpOaQtWV error) {
	hC026DpmQmus, jeWMpOaQtWV :=  /*line :1*/ UgPyZIRW(gYwswmeUjG)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	xINVctm, aYtmqRVkqc_8 :=  /*line :1*/ zoZPuC(yJ8xjP.CyorTGi, hC026DpmQmus)
	if aYtmqRVkqc_8 != 0 {
		return nil, &YxRxePaAdi{
			EQTdxl:	aYtmqRVkqc_8,
			EU0oqVev:	gYwswmeUjG,
			PtEa5aX7u5c:	"Failed to find " + gYwswmeUjG + " procedure in " + yJ8xjP.O56SJQfoFV + ": " +  /*line :1*/ aYtmqRVkqc_8.Error(),
		}
	}
	hD4wNgEB := &U4jPAQxP8l{
		OVhe0b9X:	yJ8xjP,
		AXcqb_wa:	gYwswmeUjG,
		eKFyBZWghwtO:	xINVctm,
	}
	return hD4wNgEB, nil
}

func (yJ8xjP *MQbxA8MNEJC) MustFindProc(gYwswmeUjG string) *U4jPAQxP8l {
	hD4wNgEB, aYtmqRVkqc_8 :=  /*line :1*/ yJ8xjP.FindProc(gYwswmeUjG)
	if aYtmqRVkqc_8 != nil {
		 /*line :1*/ panic(aYtmqRVkqc_8)
	}
	return hD4wNgEB
}

func (yJ8xjP *MQbxA8MNEJC) FindProcByOrdinal(bxnorpA24 uintptr) (aTSUoJ40C *U4jPAQxP8l, jeWMpOaQtWV error) {
	xINVctm, aYtmqRVkqc_8 :=  /*line :1*/ UqxfQM(yJ8xjP.CyorTGi, bxnorpA24)
	gYwswmeUjG := "#" +  /*line :1*/ dMi2GE( /*line :1*/ int(bxnorpA24))
	if aYtmqRVkqc_8 != nil {
		return nil, &YxRxePaAdi{
			EQTdxl:	aYtmqRVkqc_8,
			EU0oqVev:	gYwswmeUjG,
			PtEa5aX7u5c:	"Failed to find " + gYwswmeUjG + " procedure in " + yJ8xjP.O56SJQfoFV + ": " +  /*line :1*/ aYtmqRVkqc_8.Error(),
		}
	}
	hD4wNgEB := &U4jPAQxP8l{
		OVhe0b9X:	yJ8xjP,
		AXcqb_wa:	gYwswmeUjG,
		eKFyBZWghwtO:	xINVctm,
	}
	return hD4wNgEB, nil
}

func (yJ8xjP *MQbxA8MNEJC) MustFindProcByOrdinal(bxnorpA24 uintptr) *U4jPAQxP8l {
	hD4wNgEB, aYtmqRVkqc_8 :=  /*line :1*/ yJ8xjP.FindProcByOrdinal(bxnorpA24)
	if aYtmqRVkqc_8 != nil {
		 /*line :1*/ panic(aYtmqRVkqc_8)
	}
	return hD4wNgEB
}

func (yJ8xjP *MQbxA8MNEJC) Release() (jeWMpOaQtWV error) {
	return  /*line :1*/ I_aT6TrPE(yJ8xjP.CyorTGi)
}

type U4jPAQxP8l struct {
	OVhe0b9X	*MQbxA8MNEJC
	AXcqb_wa	string
	eKFyBZWghwtO	uintptr
}

func (hD4wNgEB *U4jPAQxP8l) Addr() uintptr {
	return hD4wNgEB.eKFyBZWghwtO
}

//go:uintptrescapes

func (hD4wNgEB *U4jPAQxP8l) Call(xINVctm ...uintptr) (i_EJOiAVI, vpWRvnw5XmyB uintptr, gTpWCN error) {
	switch  /*line :1*/ len(xINVctm) {
	case 0:
		return  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), 0, 0, 0)
	case 1:
		return  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], 0, 0)
	case 2:
		return  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], 0)
	case 3:
		return  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2])
	case 4:
		return  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], 0, 0)
	case 5:
		return  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], 0)
	case 6:
		return  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5])
	case 7:
		return  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], 0, 0)
	case 8:
		return  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], 0)
	case 9:
		return  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], xINVctm[8])
	case 10:
		return  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], xINVctm[8], xINVctm[9], 0, 0)
	case 11:
		return  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], xINVctm[8], xINVctm[9], xINVctm[10], 0)
	case 12:
		return  /*line :1*/ syscall.IiL5Io0Q5I( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], xINVctm[8], xINVctm[9], xINVctm[10], xINVctm[11])
	case 13:
		return  /*line :1*/ syscall.SAshL0on( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], xINVctm[8], xINVctm[9], xINVctm[10], xINVctm[11], xINVctm[12], 0, 0)
	case 14:
		return  /*line :1*/ syscall.SAshL0on( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], xINVctm[8], xINVctm[9], xINVctm[10], xINVctm[11], xINVctm[12], xINVctm[13], 0)
	case 15:
		return  /*line :1*/ syscall.SAshL0on( /*line :1*/ hD4wNgEB.Addr(),  /*line :1*/ uintptr( /*line :1*/ len(xINVctm)), xINVctm[0], xINVctm[1], xINVctm[2], xINVctm[3], xINVctm[4], xINVctm[5], xINVctm[6], xINVctm[7], xINVctm[8], xINVctm[9], xINVctm[10], xINVctm[11], xINVctm[12], xINVctm[13], xINVctm[14])
	default:
		 /*line :1*/ panic("Call " + hD4wNgEB.AXcqb_wa + " with too many arguments " +  /*line :1*/ dMi2GE( /*line :1*/ len(xINVctm)) + ".")
	}
}

type NpmYuea struct {
	GXl5gGf7ywW	string

	RLyLzoZao0r	bool

	vRtZ6ONxnv	sync.DIRWe11KYlYa
	r7GXJQUM1J	*MQbxA8MNEJC
}

func (yJ8xjP *NpmYuea) Load() error {

	if  /*line :1*/ atomic.LoadPointer((* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(&yJ8xjP.r7GXJQUM1J))) != nil {
		return nil
	}
	 /*line :1*/ yJ8xjP.vRtZ6ONxnv.Lock()
	defer  /*line :1*/ yJ8xjP.vRtZ6ONxnv.Unlock()
	if yJ8xjP.r7GXJQUM1J != nil {
		return nil
	}

	var o02Ffvb7mOT *MQbxA8MNEJC
	var jeWMpOaQtWV error
	if yJ8xjP.GXl5gGf7ywW == "kernel32.dll" {
		o02Ffvb7mOT, jeWMpOaQtWV =  /*line :1*/ RQ6nOPol8B(yJ8xjP.GXl5gGf7ywW)
	} else {
		o02Ffvb7mOT, jeWMpOaQtWV =  /*line :1*/ fl7X_cq(yJ8xjP.GXl5gGf7ywW, yJ8xjP.RLyLzoZao0r)
	}
	if jeWMpOaQtWV != nil {
		return jeWMpOaQtWV
	}

	 /*line :1*/ atomic.Fa2hd0oXquKk((* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(&yJ8xjP.r7GXJQUM1J)),  /*line :1*/ unsafe.Pointer(o02Ffvb7mOT))
	return nil
}

func (yJ8xjP *NpmYuea) azBmiVT() {
	aYtmqRVkqc_8 :=  /*line :1*/ yJ8xjP.Load()
	if aYtmqRVkqc_8 != nil {
		 /*line :1*/ panic(aYtmqRVkqc_8)
	}
}

func (yJ8xjP *NpmYuea) Handle() uintptr {
	 /*line :1*/ yJ8xjP.azBmiVT()
	return  /*line :1*/ uintptr(yJ8xjP.r7GXJQUM1J.CyorTGi)
}

func (yJ8xjP *NpmYuea) NewProc(gYwswmeUjG string) *LZYXjB {
	return &LZYXjB{knyHbFaZ: yJ8xjP, MVBcc9hwxK: gYwswmeUjG}
}

func UiPXNJL7we_S(gYwswmeUjG string) *NpmYuea {
	return &NpmYuea{GXl5gGf7ywW: gYwswmeUjG}
}

func FDchPxUZ(gYwswmeUjG string) *NpmYuea {
	return &NpmYuea{GXl5gGf7ywW: gYwswmeUjG, RLyLzoZao0r: true}
}

type LZYXjB struct {
	MVBcc9hwxK	string

	wxayb3xWmN	sync.DIRWe11KYlYa
	knyHbFaZ	*NpmYuea
	eW8Dqy	*U4jPAQxP8l
}

func (hD4wNgEB *LZYXjB) Find() error {

	if  /*line :1*/ atomic.LoadPointer((* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(&hD4wNgEB.eW8Dqy))) == nil {
		 /*line :1*/ hD4wNgEB.wxayb3xWmN.Lock()
		defer  /*line :1*/ hD4wNgEB.wxayb3xWmN.Unlock()
		if hD4wNgEB.eW8Dqy == nil {
			aYtmqRVkqc_8 :=  /*line :1*/ hD4wNgEB.knyHbFaZ.Load()
			if aYtmqRVkqc_8 != nil {
				return aYtmqRVkqc_8
			}
			aTSUoJ40C, aYtmqRVkqc_8 :=  /*line :1*/ hD4wNgEB.knyHbFaZ.r7GXJQUM1J.FindProc(hD4wNgEB.MVBcc9hwxK)
			if aYtmqRVkqc_8 != nil {
				return aYtmqRVkqc_8
			}

			 /*line :1*/ atomic.Fa2hd0oXquKk((* /*line :1*/ unsafe.Pointer)( /*line :1*/ unsafe.Pointer(&hD4wNgEB.eW8Dqy)),  /*line :1*/ unsafe.Pointer(aTSUoJ40C))
		}
	}
	return nil
}

func (hD4wNgEB *LZYXjB) zh6Hzhu() {
	aYtmqRVkqc_8 :=  /*line :1*/ hD4wNgEB.Find()
	if aYtmqRVkqc_8 != nil {
		 /*line :1*/ panic(aYtmqRVkqc_8)
	}
}

func (hD4wNgEB *LZYXjB) Addr() uintptr {
	 /*line :1*/ hD4wNgEB.zh6Hzhu()
	return  /*line :1*/ hD4wNgEB.eW8Dqy.Addr()
}

//go:uintptrescapes

func (hD4wNgEB *LZYXjB) Call(xINVctm ...uintptr) (i_EJOiAVI, vpWRvnw5XmyB uintptr, gTpWCN error) {
	 /*line :1*/ hD4wNgEB.zh6Hzhu()
	return  /*line :1*/ hD4wNgEB.eW8Dqy.Call(xINVctm...)
}

var s4BroDk struct {
	sync.LhBfwn6wa1x
	mk5nYK	bool
}

func brItHW() {

	s4BroDk.mk5nYK = ( /*line :1*/ fqaBrBSogv.NewProc("AddDllDirectory").Find() == nil)
}

func co6j6P() bool {
	 /*line :1*/ s4BroDk.Do(brItHW)
	return s4BroDk.mk5nYK
}

func tWdb3qQs6(gYwswmeUjG string) bool {
	for _, jpk9X7z := range gYwswmeUjG {
		if jpk9X7z == ':' || jpk9X7z == '/' || jpk9X7z == '\\' {
			return false
		}
	}
	return true
}

func fl7X_cq(gYwswmeUjG string, uxLgaklMr2w5 bool) (*MQbxA8MNEJC, error) {
	jU1Kst4 := gYwswmeUjG
	var a_MrGKcrR uintptr
	if uxLgaklMr2w5 {
		if  /*line :1*/ co6j6P() {
			a_MrGKcrR = LOAD_LIBRARY_SEARCH_SYSTEM32
		} else if  /*line :1*/ tWdb3qQs6(gYwswmeUjG) {

			lcuaKD5v, jeWMpOaQtWV :=  /*line :1*/ Voabn888O()
			if jeWMpOaQtWV != nil {
				return nil, jeWMpOaQtWV
			}
			jU1Kst4 = lcuaKD5v + "\\" + gYwswmeUjG
		}
	}
	_ITTe4, jeWMpOaQtWV :=  /*line :1*/ JtGRA3Jqe(jU1Kst4, 0, a_MrGKcrR)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	return &MQbxA8MNEJC{O56SJQfoFV: gYwswmeUjG, CyorTGi: _ITTe4}, nil
}

type nZwiPb string

func (bamc83qA3DBr nZwiPb) Error() string	{ return  /*line :1*/ string(bamc83qA3DBr) }
