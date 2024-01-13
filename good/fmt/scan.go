//line :1
package kFsoCfy5zWG

import (
	errors "iAHoxjmM"
	io "xI9ai2djaJ2"
	math "math"
	os "hPMrClpC"
	reflect "reflect"
	strconv "zBESu2SrRjP"
	sync "sync"
	utf8 "CuAc0pPZwf"
)




type XDnSs0f3 interface {
	
	
	
	
	ReadRune() (vuT9drijWK rune, e2v7yTpR2 int, jDbNTz1lC_ error)
	
	UnreadRune() error
	
	
	
	SkipSpace()
	
	
	
	
	
	
	
	
	Token(s3qzXUGDlRE bool, zEHZ5ak0_jf func(rune) bool) (tM4BoDQ1vlNm []byte, jDbNTz1lC_ error)
	
	
	Width() (xgwcYhwzhN int, xP5nLuM9_y bool)
	
	
	
	Read(d8wxci5 []byte) (bMu41ab1Tf int, jDbNTz1lC_ error)
}





type MqdnB6 interface {
	Scan(l7qaxfr XDnSs0f3, mEbw55xKIj4H rune) error
}





func ZEKdfH4alf(nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ M_v_geArHnZa(os.RUTFbiB6NV, nghaNY...)
}



func O7Xk_ZhxpukS(nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ UEYusgNHl(os.RUTFbiB6NV, nghaNY...)
}








func Svaq7XR(aCNXxCRXOS5o string, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ CTk6cS(os.RUTFbiB6NV, aCNXxCRXOS5o, nghaNY...)
}

type h3BoaETtwH string

func (vuT9drijWK *h3BoaETtwH) Read(dajeEj8 []byte) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	bMu41ab1Tf =  /*line :1*/ copy(dajeEj8, *vuT9drijWK)
	*vuT9drijWK = (*vuT9drijWK)[bMu41ab1Tf:]
	if bMu41ab1Tf == 0 {
		jDbNTz1lC_ = io.K5Sqco
	}
	return
}





func EnNnJReag(jcJKXch string, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ M_v_geArHnZa((* /*line :1*/ h3BoaETtwH)(&jcJKXch), nghaNY...)
}



func StjsU4_M2X(jcJKXch string, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ UEYusgNHl((* /*line :1*/ h3BoaETtwH)(&jcJKXch), nghaNY...)
}





func N9AdbEN381(jcJKXch string, aCNXxCRXOS5o string, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ CTk6cS((* /*line :1*/ h3BoaETtwH)(&jcJKXch), aCNXxCRXOS5o, nghaNY...)
}





func M_v_geArHnZa(vuT9drijWK io.YJ04Fau, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	a_CBqX, idIhDNly :=  /*line :1*/ e3jMasJC9z9(vuT9drijWK, true, false)
	bMu41ab1Tf, jDbNTz1lC_ =  /*line :1*/ a_CBqX.oe7XLd2oHFat(nghaNY)
	 /*line :1*/ a_CBqX.aUrCCiO(idIhDNly)
	return
}



func UEYusgNHl(vuT9drijWK io.YJ04Fau, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	a_CBqX, idIhDNly :=  /*line :1*/ e3jMasJC9z9(vuT9drijWK, false, true)
	bMu41ab1Tf, jDbNTz1lC_ =  /*line :1*/ a_CBqX.oe7XLd2oHFat(nghaNY)
	 /*line :1*/ a_CBqX.aUrCCiO(idIhDNly)
	return
}





func CTk6cS(vuT9drijWK io.YJ04Fau, aCNXxCRXOS5o string, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	a_CBqX, idIhDNly :=  /*line :1*/ e3jMasJC9z9(vuT9drijWK, false, false)
	bMu41ab1Tf, jDbNTz1lC_ =  /*line :1*/ a_CBqX.rFHTfa0M3(aCNXxCRXOS5o, nghaNY)
	 /*line :1*/ a_CBqX.aUrCCiO(idIhDNly)
	return
}



type uJ4X71BjaNv4 struct {
	i_VXKKvJ error
}

const eof = -1


type rrkIecgD struct {
	j2HyMEtPU_mR	io.MfPRgoXllY	
	mbFSudjCGo	apmFUoDP	
	h90JLRma	int	
	lcIkji_3	bool	
	jKPV63EKxe
}



type jKPV63EKxe struct {
	lsrNOd25o	bool	
	oYKqyG	bool	
	a4BrTa	bool	
	iJsthwzTgRa	int	
	ci4KQEj8tjJ	int	
	gae68Pifxa	int	
}




func (a_CBqX *rrkIecgD) Read(d8wxci5 []byte) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return 0,  /*line :1*/ errors.Su6g6hRoi9X("ScanState's Read should not be called. Use ReadRune")
}

func (a_CBqX *rrkIecgD) ReadRune() (vuT9drijWK rune, e2v7yTpR2 int, jDbNTz1lC_ error) {
	if a_CBqX.lcIkji_3 || a_CBqX.h90JLRma >= a_CBqX.iJsthwzTgRa {
		jDbNTz1lC_ = io.K5Sqco
		return
	}

	vuT9drijWK, e2v7yTpR2, jDbNTz1lC_ =  /*line :1*/ a_CBqX.j2HyMEtPU_mR.ReadRune()
	if jDbNTz1lC_ == nil {
		a_CBqX.h90JLRma++
		if a_CBqX.oYKqyG && vuT9drijWK == '\n' {
			a_CBqX.lcIkji_3 = true
		}
	} else if jDbNTz1lC_ == io.K5Sqco {
		a_CBqX.lcIkji_3 = true
	}
	return
}

func (a_CBqX *rrkIecgD) Width() (xgwcYhwzhN int, xP5nLuM9_y bool) {
	if a_CBqX.gae68Pifxa == hugeWid {
		return 0, false
	}
	return a_CBqX.gae68Pifxa, true
}



func (a_CBqX *rrkIecgD) tv4NHbyxyO() (vuT9drijWK rune) {
	vuT9drijWK, _, jDbNTz1lC_ :=  /*line :1*/ a_CBqX.ReadRune()
	if jDbNTz1lC_ != nil {
		if jDbNTz1lC_ == io.K5Sqco {
			return eof
		}
		 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
	}
	return
}




func (a_CBqX *rrkIecgD) hJhaYl2fDSR() (vuT9drijWK rune) {
	vuT9drijWK =  /*line :1*/ a_CBqX.tv4NHbyxyO()
	if vuT9drijWK == eof {
		 /*line :1*/ a_CBqX.dh0UfRf(io.UASgojMNPA)
	}
	return
}

func (a_CBqX *rrkIecgD) UnreadRune() error {
	 /*line :1*/ a_CBqX.j2HyMEtPU_mR.UnreadRune()
	a_CBqX.lcIkji_3 = false
	a_CBqX.h90JLRma--
	return nil
}

func (a_CBqX *rrkIecgD) dh0UfRf(jDbNTz1lC_ error) {
	 /*line :1*/ panic(uJ4X71BjaNv4{jDbNTz1lC_})
}

func (a_CBqX *rrkIecgD) d1S01Gz(jDbNTz1lC_ string) {
	 /*line :1*/ panic(uJ4X71BjaNv4{ /*line :1*/ errors.Su6g6hRoi9X(jDbNTz1lC_)})
}

func (a_CBqX *rrkIecgD) Token(s3qzXUGDlRE bool, zEHZ5ak0_jf func(rune) bool) (yx_BmYf1EH []byte, jDbNTz1lC_ error) {
	defer func() {
		if  /*line :1*/ l7GPpgp_Rl :=  /*line :1*/ recover(); l7GPpgp_Rl != nil {
			if u2_owHMKgs, xP5nLuM9_y := l7GPpgp_Rl.(uJ4X71BjaNv4); xP5nLuM9_y {
				jDbNTz1lC_ = u2_owHMKgs.i_VXKKvJ
			} else {
				 /*line :1*/ panic(l7GPpgp_Rl)
			}
		}
	}()
	if zEHZ5ak0_jf == nil {
		zEHZ5ak0_jf = vJsUjaCelNP
	}
	a_CBqX.mbFSudjCGo = a_CBqX.mbFSudjCGo[:0]
	yx_BmYf1EH =  /*line :1*/ a_CBqX.tM4BoDQ1vlNm(s3qzXUGDlRE, zEHZ5ak0_jf)
	return
}



var hbG8pNaZ = [][2]uint16{
	{0x0009, 0x000d},
	{0x0020, 0x0020},
	{0x0085, 0x0085},
	{0x00a0, 0x00a0},
	{0x1680, 0x1680},
	{0x2000, 0x200a},
	{0x2028, 0x2029},
	{0x202f, 0x202f},
	{0x205f, 0x205f},
	{0x3000, 0x3000},
}

func hZLGRorYw1yo(vuT9drijWK rune) bool {
	if vuT9drijWK >= 1<<16 {
		return false
	}
	xPuUwwRRTn :=  /*line :1*/ uint16(vuT9drijWK)
	for _, y22Yo3RWCL := range hbG8pNaZ {
		if xPuUwwRRTn < y22Yo3RWCL[0] {
			return false
		}
		if xPuUwwRRTn <= y22Yo3RWCL[1] {
			return true
		}
	}
	return false
}


func vJsUjaCelNP(vuT9drijWK rune) bool {
	return ! /*line :1*/ hZLGRorYw1yo(vuT9drijWK)
}




type oNl_aw struct {
	fsueJBHXLo	io.YJ04Fau
	naFSu1dyk	[utf8.UTFMax]byte	
	mhUSqSEBk	int	
	jMsoIi55g8wm	[utf8.UTFMax]byte	
	h20kMY	rune	
}



func (vuT9drijWK *oNl_aw) ftbkNVV6vpQ9() (dajeEj8 byte, jDbNTz1lC_ error) {
	if vuT9drijWK.mhUSqSEBk > 0 {
		dajeEj8 = vuT9drijWK.jMsoIi55g8wm[0]
		 /*line :1*/ copy(vuT9drijWK.jMsoIi55g8wm[0:], vuT9drijWK.jMsoIi55g8wm[1:])
		vuT9drijWK.mhUSqSEBk--
		return
	}
	bMu41ab1Tf, jDbNTz1lC_ :=  /*line :1*/ io.G0zBgKg(vuT9drijWK.fsueJBHXLo, vuT9drijWK.jMsoIi55g8wm[:1])
	if bMu41ab1Tf != 1 {
		return 0, jDbNTz1lC_
	}
	return vuT9drijWK.jMsoIi55g8wm[0], jDbNTz1lC_
}



func (vuT9drijWK *oNl_aw) ReadRune() (rh7XkYej rune, e2v7yTpR2 int, jDbNTz1lC_ error) {
	if vuT9drijWK.h20kMY >= 0 {
		rh7XkYej = vuT9drijWK.h20kMY
		vuT9drijWK.h20kMY = ^vuT9drijWK.h20kMY
		e2v7yTpR2 =  /*line :1*/ utf8.YrF2oG995j9(rh7XkYej)
		return
	}
	vuT9drijWK.naFSu1dyk[0], jDbNTz1lC_ =  /*line :1*/ vuT9drijWK.ftbkNVV6vpQ9()
	if jDbNTz1lC_ != nil {
		return
	}
	if vuT9drijWK.naFSu1dyk[0] < utf8.RuneSelf {
		rh7XkYej =  /*line :1*/ rune(vuT9drijWK.naFSu1dyk[0])
		e2v7yTpR2 = 1

		vuT9drijWK.h20kMY = ^rh7XkYej
		return
	}
	var bMu41ab1Tf int
	for bMu41ab1Tf = 1; ! /*line :1*/ utf8.TAy5Jt8FB7p4(vuT9drijWK.naFSu1dyk[:bMu41ab1Tf]); bMu41ab1Tf++ {
		vuT9drijWK.naFSu1dyk[bMu41ab1Tf], jDbNTz1lC_ =  /*line :1*/ vuT9drijWK.ftbkNVV6vpQ9()
		if jDbNTz1lC_ != nil {
			if jDbNTz1lC_ == io.K5Sqco {
				jDbNTz1lC_ = nil
				break
			}
			return
		}
	}
	rh7XkYej, e2v7yTpR2 =  /*line :1*/ utf8.EicfpCPn(vuT9drijWK.naFSu1dyk[:bMu41ab1Tf])
	if e2v7yTpR2 < bMu41ab1Tf {
		 /*line :1*/ copy(vuT9drijWK.jMsoIi55g8wm[vuT9drijWK.mhUSqSEBk:], vuT9drijWK.naFSu1dyk[e2v7yTpR2:bMu41ab1Tf])
		vuT9drijWK.mhUSqSEBk += bMu41ab1Tf - e2v7yTpR2
	}

	vuT9drijWK.h20kMY = ^rh7XkYej
	return
}

func (vuT9drijWK *oNl_aw) UnreadRune() error {
	if vuT9drijWK.h20kMY >= 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("fmt: scanning called UnreadRune with no rune available")
	}

	vuT9drijWK.h20kMY = ^vuT9drijWK.h20kMY
	return nil
}

var nUiPQaFJOV = sync.OrP5FGPq{
	IYbTy050ek: func() any { return  /*line :1*/ new(rrkIecgD) },
}


func e3jMasJC9z9(vuT9drijWK io.YJ04Fau, kHaHCcN, ylKF1LO_ bool) (a_CBqX *rrkIecgD, idIhDNly jKPV63EKxe) {
	a_CBqX =  /*line :1*/ nUiPQaFJOV.Get().(*rrkIecgD)
	if wgmhImVCCqa, xP5nLuM9_y := vuT9drijWK.(io.MfPRgoXllY); xP5nLuM9_y {
		a_CBqX.j2HyMEtPU_mR = wgmhImVCCqa
	} else {
		a_CBqX.j2HyMEtPU_mR = &oNl_aw{fsueJBHXLo: vuT9drijWK, h20kMY: -1}
	}
	a_CBqX.a4BrTa = kHaHCcN
	a_CBqX.oYKqyG = ylKF1LO_
	a_CBqX.lcIkji_3 = false
	a_CBqX.ci4KQEj8tjJ = hugeWid
	a_CBqX.iJsthwzTgRa = hugeWid
	a_CBqX.gae68Pifxa = hugeWid
	a_CBqX.lsrNOd25o = true
	a_CBqX.h90JLRma = 0
	return
}


func (a_CBqX *rrkIecgD) aUrCCiO(idIhDNly jKPV63EKxe) {

	if idIhDNly.lsrNOd25o {
		a_CBqX.jKPV63EKxe = idIhDNly
		return
	}

	if  /*line :1*/ cap(a_CBqX.mbFSudjCGo) > 1024 {
		return
	}
	a_CBqX.mbFSudjCGo = a_CBqX.mbFSudjCGo[:0]
	a_CBqX.j2HyMEtPU_mR = nil
	 /*line :1*/ nUiPQaFJOV.Put(a_CBqX)
}




func (a_CBqX *rrkIecgD) SkipSpace() {
	for {
		vuT9drijWK :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
		if vuT9drijWK == eof {
			return
		}
		if vuT9drijWK == '\r' &&  /*line :1*/ a_CBqX.w7XpAo14MFB("\n") {
			continue
		}
		if vuT9drijWK == '\n' {
			if a_CBqX.a4BrTa {
				continue
			}
			 /*line :1*/ a_CBqX.d1S01Gz("unexpected newline")
			return
		}
		if ! /*line :1*/ hZLGRorYw1yo(vuT9drijWK) {
			 /*line :1*/ a_CBqX.UnreadRune()
			break
		}
	}
}




func (a_CBqX *rrkIecgD) tM4BoDQ1vlNm(s3qzXUGDlRE bool, zEHZ5ak0_jf func(rune) bool) []byte {
	if s3qzXUGDlRE {
		 /*line :1*/ a_CBqX.SkipSpace()
	}

	for {
		vuT9drijWK :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
		if vuT9drijWK == eof {
			break
		}
		if ! /*line :1*/ zEHZ5ak0_jf(vuT9drijWK) {
			 /*line :1*/ a_CBqX.UnreadRune()
			break
		}
		 /*line :1*/ a_CBqX.mbFSudjCGo.j9YRMbR(vuT9drijWK)
	}
	return a_CBqX.mbFSudjCGo
}

var ukbPGFK =  /*line :1*/ errors.Su6g6hRoi9X("syntax error scanning complex number")
var iUcNix2Sc5SR =  /*line :1*/ errors.Su6g6hRoi9X("syntax error scanning boolean")

func whwcs5z(a_CBqX string, vuT9drijWK rune) int {
	for wKUgq0A, gaq8WME := range a_CBqX {
		if gaq8WME == vuT9drijWK {
			return wKUgq0A
		}
	}
	return -1
}



func (a_CBqX *rrkIecgD) kjm3PaIelaQX(xP5nLuM9_y string, bxNJKAr bool) bool {
	vuT9drijWK :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
	if vuT9drijWK == eof {
		return false
	}
	if  /*line :1*/ whwcs5z(xP5nLuM9_y, vuT9drijWK) >= 0 {
		if bxNJKAr {
			 /*line :1*/ a_CBqX.mbFSudjCGo.j9YRMbR(vuT9drijWK)
		}
		return true
	}
	if vuT9drijWK != eof && bxNJKAr {
		 /*line :1*/ a_CBqX.UnreadRune()
	}
	return false
}


func (a_CBqX *rrkIecgD) w7XpAo14MFB(xP5nLuM9_y string) bool {
	vuT9drijWK :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
	if vuT9drijWK != eof {
		 /*line :1*/ a_CBqX.UnreadRune()
	}
	return  /*line :1*/ whwcs5z(xP5nLuM9_y, vuT9drijWK) >= 0
}

func (a_CBqX *rrkIecgD) co5OOqe() {

	if vuT9drijWK :=  /*line :1*/ a_CBqX.tv4NHbyxyO(); vuT9drijWK == eof {
		 /*line :1*/ panic(io.K5Sqco)
	}
	 /*line :1*/ a_CBqX.UnreadRune()
}



func (a_CBqX *rrkIecgD) bxNJKAr(xP5nLuM9_y string) bool {
	return  /*line :1*/ a_CBqX.kjm3PaIelaQX(xP5nLuM9_y, true)
}


func (a_CBqX *rrkIecgD) sn8OV9BXhFf(mEbw55xKIj4H rune, iTfKl0yrH, rUN7I9qbYmEq string) bool {
	for _, xsKqMyFqj := range iTfKl0yrH {
		if xsKqMyFqj == mEbw55xKIj4H {
			return true
		}
	}
	 /*line :1*/ a_CBqX.d1S01Gz("bad verb '%" +  /*line :1*/ string(mEbw55xKIj4H) + "' for " + rUN7I9qbYmEq)
	return false
}


func (a_CBqX *rrkIecgD) glNkIp(mEbw55xKIj4H rune) bool {
	 /*line :1*/ a_CBqX.SkipSpace()
	 /*line :1*/ a_CBqX.co5OOqe()
	if ! /*line :1*/ a_CBqX.sn8OV9BXhFf(mEbw55xKIj4H, "tv", "boolean") {
		return false
	}

	switch  /*line :1*/ a_CBqX.tv4NHbyxyO() {
	case '0':
		return false
	case '1':
		return true
	case 't', 'T':
		if  /*line :1*/ a_CBqX.bxNJKAr("rR") && (! /*line :1*/ a_CBqX.bxNJKAr("uU") || ! /*line :1*/ a_CBqX.bxNJKAr("eE")) {
			 /*line :1*/ a_CBqX.dh0UfRf(iUcNix2Sc5SR)
		}
		return true
	case 'f', 'F':
		if  /*line :1*/ a_CBqX.bxNJKAr("aA") && (! /*line :1*/ a_CBqX.bxNJKAr("lL") || ! /*line :1*/ a_CBqX.bxNJKAr("sS") || ! /*line :1*/ a_CBqX.bxNJKAr("eE")) {
			 /*line :1*/ a_CBqX.dh0UfRf(iUcNix2Sc5SR)
		}
		return false
	}
	return false
}


const (
	binaryDigits	= "01"
	octalDigits	= "01234567"
	decimalDigits	= "0123456789"
	hexadecimalDigits	= "0123456789aAbBcCdDeEfF"
	sign	= "+-"
	period	= "."
	exponent	= "eEpP"
)


func (a_CBqX *rrkIecgD) aYaYU6V0TLEC(mEbw55xKIj4H rune) (yAJOvsG6ka int, o_QFFZ7d string) {
	 /*line :1*/ a_CBqX.sn8OV9BXhFf(mEbw55xKIj4H, "bdoUxXv", "integer")
	yAJOvsG6ka = 10
	o_QFFZ7d = decimalDigits
	switch mEbw55xKIj4H {
	case 'b':
		yAJOvsG6ka = 2
		o_QFFZ7d = binaryDigits
	case 'o':
		yAJOvsG6ka = 8
		o_QFFZ7d = octalDigits
	case 'x', 'X', 'U':
		yAJOvsG6ka = 16
		o_QFFZ7d = hexadecimalDigits
	}
	return
}


func (a_CBqX *rrkIecgD) v9ok1qp(o_QFFZ7d string, j1zCbeZ bool) string {
	if !j1zCbeZ {
		 /*line :1*/ a_CBqX.co5OOqe()
		if ! /*line :1*/ a_CBqX.bxNJKAr(o_QFFZ7d) {
			 /*line :1*/ a_CBqX.d1S01Gz("expected integer")
		}
	}
	for  /*line :1*/ a_CBqX.bxNJKAr(o_QFFZ7d) {
	}
	return  /*line :1*/ string(a_CBqX.mbFSudjCGo)
}


func (a_CBqX *rrkIecgD) rH5E6DQudaib(txJUL_v int) int64 {
	 /*line :1*/ a_CBqX.co5OOqe()
	vuT9drijWK :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
	bMu41ab1Tf :=  /*line :1*/ uint(txJUL_v)
	vFAbma1 := ( /*line :1*/ int64(vuT9drijWK) << (64 - bMu41ab1Tf)) >> (64 - bMu41ab1Tf)
	if vFAbma1 !=  /*line :1*/ int64(vuT9drijWK) {
		 /*line :1*/ a_CBqX.d1S01Gz("overflow on character value " +  /*line :1*/ string(vuT9drijWK))
	}
	return  /*line :1*/ int64(vuT9drijWK)
}




func (a_CBqX *rrkIecgD) jsjKD1K9h9Y7() (yAJOvsG6ka int, o_QFFZ7d string, fSeXzvx bool) {
	if ! /*line :1*/ a_CBqX.w7XpAo14MFB("0") {
		return 0, decimalDigits + "_", false
	}
	 /*line :1*/ a_CBqX.bxNJKAr("0")

	switch {
	case  /*line :1*/ a_CBqX.w7XpAo14MFB("bB"):
		 /*line :1*/ a_CBqX.kjm3PaIelaQX("bB", true)
		return 0, binaryDigits + "_", true
	case  /*line :1*/ a_CBqX.w7XpAo14MFB("oO"):
		 /*line :1*/ a_CBqX.kjm3PaIelaQX("oO", true)
		return 0, octalDigits + "_", true
	case  /*line :1*/ a_CBqX.w7XpAo14MFB("xX"):
		 /*line :1*/ a_CBqX.kjm3PaIelaQX("xX", true)
		return 0, hexadecimalDigits + "_", true
	default:
		return 0, octalDigits + "_", true
	}
}



func (a_CBqX *rrkIecgD) aivjNfWKi_(mEbw55xKIj4H rune, txJUL_v int) int64 {
	if mEbw55xKIj4H == 'c' {
		return  /*line :1*/ a_CBqX.rH5E6DQudaib(txJUL_v)
	}
	 /*line :1*/ a_CBqX.SkipSpace()
	 /*line :1*/ a_CBqX.co5OOqe()
	yAJOvsG6ka, o_QFFZ7d :=  /*line :1*/ a_CBqX.aYaYU6V0TLEC(mEbw55xKIj4H)
	j1zCbeZ := false
	if mEbw55xKIj4H == 'U' {
		if ! /*line :1*/ a_CBqX.kjm3PaIelaQX("U", false) || ! /*line :1*/ a_CBqX.kjm3PaIelaQX("+", false) {
			 /*line :1*/ a_CBqX.d1S01Gz("bad unicode format ")
		}
	} else {
		 /*line :1*/ a_CBqX.bxNJKAr(sign)
		if mEbw55xKIj4H == 'v' {
			yAJOvsG6ka, o_QFFZ7d, j1zCbeZ =  /*line :1*/ a_CBqX.jsjKD1K9h9Y7()
		}
	}
	yx_BmYf1EH :=  /*line :1*/ a_CBqX.v9ok1qp(o_QFFZ7d, j1zCbeZ)
	wKUgq0A, jDbNTz1lC_ :=  /*line :1*/ strconv.N8Xpap1Vt1ot(yx_BmYf1EH, yAJOvsG6ka, 64)
	if jDbNTz1lC_ != nil {
		 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
	}
	bMu41ab1Tf :=  /*line :1*/ uint(txJUL_v)
	vFAbma1 := (wKUgq0A << (64 - bMu41ab1Tf)) >> (64 - bMu41ab1Tf)
	if vFAbma1 != wKUgq0A {
		 /*line :1*/ a_CBqX.d1S01Gz("integer overflow on token " + yx_BmYf1EH)
	}
	return wKUgq0A
}



func (a_CBqX *rrkIecgD) aQa7hXK(mEbw55xKIj4H rune, txJUL_v int) uint64 {
	if mEbw55xKIj4H == 'c' {
		return  /*line :1*/ uint64( /*line :1*/ a_CBqX.rH5E6DQudaib(txJUL_v))
	}
	 /*line :1*/ a_CBqX.SkipSpace()
	 /*line :1*/ a_CBqX.co5OOqe()
	yAJOvsG6ka, o_QFFZ7d :=  /*line :1*/ a_CBqX.aYaYU6V0TLEC(mEbw55xKIj4H)
	j1zCbeZ := false
	if mEbw55xKIj4H == 'U' {
		if ! /*line :1*/ a_CBqX.kjm3PaIelaQX("U", false) || ! /*line :1*/ a_CBqX.kjm3PaIelaQX("+", false) {
			 /*line :1*/ a_CBqX.d1S01Gz("bad unicode format ")
		}
	} else if mEbw55xKIj4H == 'v' {
		yAJOvsG6ka, o_QFFZ7d, j1zCbeZ =  /*line :1*/ a_CBqX.jsjKD1K9h9Y7()
	}
	yx_BmYf1EH :=  /*line :1*/ a_CBqX.v9ok1qp(o_QFFZ7d, j1zCbeZ)
	wKUgq0A, jDbNTz1lC_ :=  /*line :1*/ strconv.WOxcnaOAzTeq(yx_BmYf1EH, yAJOvsG6ka, 64)
	if jDbNTz1lC_ != nil {
		 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
	}
	bMu41ab1Tf :=  /*line :1*/ uint(txJUL_v)
	vFAbma1 := (wKUgq0A << (64 - bMu41ab1Tf)) >> (64 - bMu41ab1Tf)
	if vFAbma1 != wKUgq0A {
		 /*line :1*/ a_CBqX.d1S01Gz("unsigned integer overflow on token " + yx_BmYf1EH)
	}
	return wKUgq0A
}




func (a_CBqX *rrkIecgD) uN1Xo4hyKHY() string {
	a_CBqX.mbFSudjCGo = a_CBqX.mbFSudjCGo[:0]

	if  /*line :1*/ a_CBqX.bxNJKAr("nN") &&  /*line :1*/ a_CBqX.bxNJKAr("aA") &&  /*line :1*/ a_CBqX.bxNJKAr("nN") {
		return  /*line :1*/ string(a_CBqX.mbFSudjCGo)
	}

	 /*line :1*/ a_CBqX.bxNJKAr(sign)

	if  /*line :1*/ a_CBqX.bxNJKAr("iI") &&  /*line :1*/ a_CBqX.bxNJKAr("nN") &&  /*line :1*/ a_CBqX.bxNJKAr("fF") {
		return  /*line :1*/ string(a_CBqX.mbFSudjCGo)
	}
	o_QFFZ7d := decimalDigits + "_"
	kz3_eXshJ := exponent
	if  /*line :1*/ a_CBqX.bxNJKAr("0") &&  /*line :1*/ a_CBqX.bxNJKAr("xX") {
		o_QFFZ7d = hexadecimalDigits + "_"
		kz3_eXshJ = "pP"
	}

	for  /*line :1*/ a_CBqX.bxNJKAr(o_QFFZ7d) {
	}

	if  /*line :1*/ a_CBqX.bxNJKAr(period) {

		for  /*line :1*/ a_CBqX.bxNJKAr(o_QFFZ7d) {
		}
	}

	if  /*line :1*/ a_CBqX.bxNJKAr(kz3_eXshJ) {

		 /*line :1*/ a_CBqX.bxNJKAr(sign)

		for  /*line :1*/ a_CBqX.bxNJKAr(decimalDigits + "_") {
		}
	}
	return  /*line :1*/ string(a_CBqX.mbFSudjCGo)
}




func (a_CBqX *rrkIecgD) iL3QLI() (yZcF62zuq, vSgYWw_ppk string) {

	qRI2ci :=  /*line :1*/ a_CBqX.bxNJKAr("(")
	yZcF62zuq =  /*line :1*/ a_CBqX.uN1Xo4hyKHY()
	a_CBqX.mbFSudjCGo = a_CBqX.mbFSudjCGo[:0]

	if ! /*line :1*/ a_CBqX.bxNJKAr("+-") {
		 /*line :1*/ a_CBqX.dh0UfRf(ukbPGFK)
	}

	hmgLQW2Dp :=  /*line :1*/ string(a_CBqX.mbFSudjCGo)
	vSgYWw_ppk =  /*line :1*/ a_CBqX.uN1Xo4hyKHY()
	if ! /*line :1*/ a_CBqX.bxNJKAr("i") {
		 /*line :1*/ a_CBqX.dh0UfRf(ukbPGFK)
	}
	if qRI2ci && ! /*line :1*/ a_CBqX.bxNJKAr(")") {
		 /*line :1*/ a_CBqX.dh0UfRf(ukbPGFK)
	}
	return yZcF62zuq, hmgLQW2Dp + vSgYWw_ppk
}

func dJVPtX(a_CBqX string) bool {
	for wKUgq0A := 0; wKUgq0A <  /*line :1*/ len(a_CBqX); wKUgq0A++ {
		if a_CBqX[wKUgq0A] == 'x' || a_CBqX[wKUgq0A] == 'X' {
			return true
		}
	}
	return false
}


func (a_CBqX *rrkIecgD) wraaYygtPd(jcJKXch string, bMu41ab1Tf int) float64 {

	if hphPuX8VWjn5 :=  /*line :1*/ whwcs5z(jcJKXch, 'p'); hphPuX8VWjn5 >= 0 && ! /*line :1*/ dJVPtX(jcJKXch) {

		zEHZ5ak0_jf, jDbNTz1lC_ :=  /*line :1*/ strconv.YciY2Zn(jcJKXch[:hphPuX8VWjn5], bMu41ab1Tf)
		if jDbNTz1lC_ != nil {

			if l7GPpgp_Rl, xP5nLuM9_y := jDbNTz1lC_.(*strconv.RJxImsn8bz); xP5nLuM9_y {
				l7GPpgp_Rl.HY76vGY_ = jcJKXch
			}
			 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
		}
		pNQVAaHsDCDh, jDbNTz1lC_ :=  /*line :1*/ strconv.GmbOy60GCccC(jcJKXch[hphPuX8VWjn5+1:])
		if jDbNTz1lC_ != nil {

			if l7GPpgp_Rl, xP5nLuM9_y := jDbNTz1lC_.(*strconv.RJxImsn8bz); xP5nLuM9_y {
				l7GPpgp_Rl.HY76vGY_ = jcJKXch
			}
			 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
		}
		return  /*line :1*/ math.EU1SFL1cw(zEHZ5ak0_jf, pNQVAaHsDCDh)
	}
	zEHZ5ak0_jf, jDbNTz1lC_ :=  /*line :1*/ strconv.YciY2Zn(jcJKXch, bMu41ab1Tf)
	if jDbNTz1lC_ != nil {
		 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
	}
	return zEHZ5ak0_jf
}





func (a_CBqX *rrkIecgD) hDZKaY2b(mEbw55xKIj4H rune, bMu41ab1Tf int) complex128 {
	if ! /*line :1*/ a_CBqX.sn8OV9BXhFf(mEbw55xKIj4H, floatVerbs, "complex") {
		return 0
	}
	 /*line :1*/ a_CBqX.SkipSpace()
	 /*line :1*/ a_CBqX.co5OOqe()
	tH08JGF, aZXrE6N :=  /*line :1*/ a_CBqX.iL3QLI()
	yZcF62zuq :=  /*line :1*/ a_CBqX.wraaYygtPd(tH08JGF, bMu41ab1Tf/2)
	vSgYWw_ppk :=  /*line :1*/ a_CBqX.wraaYygtPd(aZXrE6N, bMu41ab1Tf/2)
	return  /*line :1*/ complex(yZcF62zuq, vSgYWw_ppk)
}



func (a_CBqX *rrkIecgD) pu_PGxUF4Ps2(mEbw55xKIj4H rune) (jcJKXch string) {
	if ! /*line :1*/ a_CBqX.sn8OV9BXhFf(mEbw55xKIj4H, "svqxX", "string") {
		return ""
	}
	 /*line :1*/ a_CBqX.SkipSpace()
	 /*line :1*/ a_CBqX.co5OOqe()
	switch mEbw55xKIj4H {
	case 'q':
		jcJKXch =  /*line :1*/ a_CBqX.zckraph9u()
	case 'x', 'X':
		jcJKXch =  /*line :1*/ a_CBqX.b9kgcfSHdwpY()
	default:
		jcJKXch =  /*line :1*/ string( /*line :1*/ a_CBqX.tM4BoDQ1vlNm(true, vJsUjaCelNP))
	}
	return
}


func (a_CBqX *rrkIecgD) zckraph9u() string {
	 /*line :1*/ a_CBqX.co5OOqe()
	wErAB_XrZj :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
	switch wErAB_XrZj {
	case '`':

		for {
			vuT9drijWK :=  /*line :1*/ a_CBqX.hJhaYl2fDSR()
			if vuT9drijWK == wErAB_XrZj {
				break
			}
			 /*line :1*/ a_CBqX.mbFSudjCGo.j9YRMbR(vuT9drijWK)
		}
		return  /*line :1*/ string(a_CBqX.mbFSudjCGo)
	case '"':

		 /*line :1*/ a_CBqX.mbFSudjCGo.yYk12QG('"')
		for {
			vuT9drijWK :=  /*line :1*/ a_CBqX.hJhaYl2fDSR()
			 /*line :1*/ a_CBqX.mbFSudjCGo.j9YRMbR(vuT9drijWK)
			if vuT9drijWK == '\\' {

				 /*line :1*/ a_CBqX.mbFSudjCGo.j9YRMbR( /*line :1*/ a_CBqX.hJhaYl2fDSR())
			} else if vuT9drijWK == '"' {
				break
			}
		}
		try2YI, jDbNTz1lC_ :=  /*line :1*/ strconv.D0aVL66( /*line :1*/ string(a_CBqX.mbFSudjCGo))
		if jDbNTz1lC_ != nil {
			 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
		}
		return try2YI
	default:
		 /*line :1*/ a_CBqX.d1S01Gz("expected quoted string")
	}
	return ""
}


func wady_ZoSpw(o15_sLZAh rune) (int, bool) {
	w08aU4 :=  /*line :1*/ int(o15_sLZAh)
	switch w08aU4 {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return w08aU4 - '0', true
	case 'a', 'b', 'c', 'd', 'e', 'f':
		return 10 + w08aU4 - 'a', true
	case 'A', 'B', 'C', 'D', 'E', 'F':
		return 10 + w08aU4 - 'A', true
	}
	return -1, false
}




func (a_CBqX *rrkIecgD) lC3fHPK6() (dajeEj8 byte, xP5nLuM9_y bool) {
	dcoac_5CoE9C :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
	if dcoac_5CoE9C == eof {
		return
	}
	d16gZvJ, xP5nLuM9_y :=  /*line :1*/ wady_ZoSpw(dcoac_5CoE9C)
	if !xP5nLuM9_y {
		 /*line :1*/ a_CBqX.UnreadRune()
		return
	}
	lj3uCTRyTz4_, xP5nLuM9_y :=  /*line :1*/ wady_ZoSpw( /*line :1*/ a_CBqX.hJhaYl2fDSR())
	if !xP5nLuM9_y {
		 /*line :1*/ a_CBqX.d1S01Gz("illegal hex digit")
		return
	}
	return  /*line :1*/ byte(d16gZvJ<<4 | lj3uCTRyTz4_), true
}


func (a_CBqX *rrkIecgD) b9kgcfSHdwpY() string {
	 /*line :1*/ a_CBqX.co5OOqe()
	for {
		dajeEj8, xP5nLuM9_y :=  /*line :1*/ a_CBqX.lC3fHPK6()
		if !xP5nLuM9_y {
			break
		}
		 /*line :1*/ a_CBqX.mbFSudjCGo.yYk12QG(dajeEj8)
	}
	if  /*line :1*/ len(a_CBqX.mbFSudjCGo) == 0 {
		 /*line :1*/ a_CBqX.d1S01Gz("no hex data for %x string")
		return ""
	}
	return  /*line :1*/ string(a_CBqX.mbFSudjCGo)
}

const (
	floatVerbs	= "beEfFgGv"

	hugeWid	= 1 << 30

	intBits	= 32 << (^ /*line :1*/ uint(0) >> 63)
	uintptrBits	= 32 << (^ /*line :1*/ uintptr(0) >> 63)
)


func (a_CBqX *rrkIecgD) kPMngA3nd() {
	 /*line :1*/ a_CBqX.SkipSpace()
	 /*line :1*/ a_CBqX.co5OOqe()
	if ! /*line :1*/ a_CBqX.bxNJKAr("%") {
		 /*line :1*/ a_CBqX.d1S01Gz("missing literal %")
	}
}


func (a_CBqX *rrkIecgD) ez1E6lSHl(mEbw55xKIj4H rune, oaT9vLaMI3 any) {
	a_CBqX.mbFSudjCGo = a_CBqX.mbFSudjCGo[:0]
	var jDbNTz1lC_ error

	if xsKqMyFqj, xP5nLuM9_y := oaT9vLaMI3.(MqdnB6); xP5nLuM9_y {
		jDbNTz1lC_ =  /*line :1*/ xsKqMyFqj.Scan(a_CBqX, mEbw55xKIj4H)
		if jDbNTz1lC_ != nil {
			if jDbNTz1lC_ == io.K5Sqco {
				jDbNTz1lC_ = io.UASgojMNPA
			}
			 /*line :1*/ a_CBqX.dh0UfRf(jDbNTz1lC_)
		}
		return
	}

	switch xsKqMyFqj := oaT9vLaMI3.(type) {
	case *bool:
		*xsKqMyFqj =  /*line :1*/ a_CBqX.glNkIp(mEbw55xKIj4H)
	case *complex64:
		*xsKqMyFqj =  /*line :1*/ complex64( /*line :1*/ a_CBqX.hDZKaY2b(mEbw55xKIj4H, 64))
	case *complex128:
		*xsKqMyFqj =  /*line :1*/ a_CBqX.hDZKaY2b(mEbw55xKIj4H, 128)
	case *int:
		*xsKqMyFqj =  /*line :1*/ int( /*line :1*/ a_CBqX.aivjNfWKi_(mEbw55xKIj4H, intBits))
	case *int8:
		*xsKqMyFqj =  /*line :1*/ int8( /*line :1*/ a_CBqX.aivjNfWKi_(mEbw55xKIj4H, 8))
	case *int16:
		*xsKqMyFqj =  /*line :1*/ int16( /*line :1*/ a_CBqX.aivjNfWKi_(mEbw55xKIj4H, 16))
	case *int32:
		*xsKqMyFqj =  /*line :1*/ int32( /*line :1*/ a_CBqX.aivjNfWKi_(mEbw55xKIj4H, 32))
	case *int64:
		*xsKqMyFqj =  /*line :1*/ a_CBqX.aivjNfWKi_(mEbw55xKIj4H, 64)
	case *uint:
		*xsKqMyFqj =  /*line :1*/ uint( /*line :1*/ a_CBqX.aQa7hXK(mEbw55xKIj4H, intBits))
	case *uint8:
		*xsKqMyFqj =  /*line :1*/ uint8( /*line :1*/ a_CBqX.aQa7hXK(mEbw55xKIj4H, 8))
	case *uint16:
		*xsKqMyFqj =  /*line :1*/ uint16( /*line :1*/ a_CBqX.aQa7hXK(mEbw55xKIj4H, 16))
	case *uint32:
		*xsKqMyFqj =  /*line :1*/ uint32( /*line :1*/ a_CBqX.aQa7hXK(mEbw55xKIj4H, 32))
	case *uint64:
		*xsKqMyFqj =  /*line :1*/ a_CBqX.aQa7hXK(mEbw55xKIj4H, 64)
	case *uintptr:
		*xsKqMyFqj =  /*line :1*/ uintptr( /*line :1*/ a_CBqX.aQa7hXK(mEbw55xKIj4H, uintptrBits))

	case *float32:
		if  /*line :1*/ a_CBqX.sn8OV9BXhFf(mEbw55xKIj4H, floatVerbs, "float32") {
			 /*line :1*/ a_CBqX.SkipSpace()
			 /*line :1*/ a_CBqX.co5OOqe()
			*xsKqMyFqj =  /*line :1*/ float32( /*line :1*/ a_CBqX.wraaYygtPd( /*line :1*/ a_CBqX.uN1Xo4hyKHY(), 32))
		}
	case *float64:
		if  /*line :1*/ a_CBqX.sn8OV9BXhFf(mEbw55xKIj4H, floatVerbs, "float64") {
			 /*line :1*/ a_CBqX.SkipSpace()
			 /*line :1*/ a_CBqX.co5OOqe()
			*xsKqMyFqj =  /*line :1*/ a_CBqX.wraaYygtPd( /*line :1*/ a_CBqX.uN1Xo4hyKHY(), 64)
		}
	case *string:
		*xsKqMyFqj =  /*line :1*/ a_CBqX.pu_PGxUF4Ps2(mEbw55xKIj4H)
	case *[]byte:

		*xsKqMyFqj = [] /*line :1*/ byte( /*line :1*/ a_CBqX.pu_PGxUF4Ps2(mEbw55xKIj4H))
	default:
		cZesrlF9C :=  /*line :1*/ reflect.SdHoaIQl(xsKqMyFqj)
		gWdvuEAK := cZesrlF9C
		if  /*line :1*/ gWdvuEAK.Kind() != reflect.Pointer {
			 /*line :1*/ a_CBqX.d1S01Gz("type not a pointer: " +  /*line :1*/ cZesrlF9C.Type().String())
			return
		}
		switch xsKqMyFqj :=  /*line :1*/ gWdvuEAK.Elem();  /*line :1*/ xsKqMyFqj.Kind() {
		case reflect.Bool:
			 /*line :1*/ xsKqMyFqj.SetBool( /*line :1*/ a_CBqX.glNkIp(mEbw55xKIj4H))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			 /*line :1*/ xsKqMyFqj.SetInt( /*line :1*/ a_CBqX.aivjNfWKi_(mEbw55xKIj4H,  /*line :1*/ xsKqMyFqj.Type().Bits()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			 /*line :1*/ xsKqMyFqj.SetUint( /*line :1*/ a_CBqX.aQa7hXK(mEbw55xKIj4H,  /*line :1*/ xsKqMyFqj.Type().Bits()))
		case reflect.String:
			 /*line :1*/ xsKqMyFqj.SetString( /*line :1*/ a_CBqX.pu_PGxUF4Ps2(mEbw55xKIj4H))
		case reflect.Slice:

			rUN7I9qbYmEq :=  /*line :1*/ xsKqMyFqj.Type()
			if  /*line :1*/ rUN7I9qbYmEq.Elem().Kind() != reflect.Uint8 {
				 /*line :1*/ a_CBqX.d1S01Gz("can't scan type: " +  /*line :1*/ cZesrlF9C.Type().String())
			}
			jcJKXch :=  /*line :1*/ a_CBqX.pu_PGxUF4Ps2(mEbw55xKIj4H)
			 /*line :1*/ xsKqMyFqj.Set( /*line :1*/ reflect.LydH53J9Si(rUN7I9qbYmEq,  /*line :1*/ len(jcJKXch),  /*line :1*/ len(jcJKXch)))
			for wKUgq0A := 0; wKUgq0A <  /*line :1*/ len(jcJKXch); wKUgq0A++ {
				 /*line :1*/ xsKqMyFqj.Index(wKUgq0A).SetUint( /*line :1*/ uint64(jcJKXch[wKUgq0A]))
			}
		case reflect.Float32, reflect.Float64:
			 /*line :1*/ a_CBqX.SkipSpace()
			 /*line :1*/ a_CBqX.co5OOqe()
			 /*line :1*/ xsKqMyFqj.SetFloat( /*line :1*/ a_CBqX.wraaYygtPd( /*line :1*/ a_CBqX.uN1Xo4hyKHY(),  /*line :1*/ xsKqMyFqj.Type().Bits()))
		case reflect.Complex64, reflect.Complex128:
			 /*line :1*/ xsKqMyFqj.SetComplex( /*line :1*/ a_CBqX.hDZKaY2b(mEbw55xKIj4H,  /*line :1*/ xsKqMyFqj.Type().Bits()))
		default:
			 /*line :1*/ a_CBqX.d1S01Gz("can't scan type: " +  /*line :1*/ cZesrlF9C.Type().String())
		}
	}
}


func qU7_PYRCR(yUhNNLjWzH8j *error) {
	if l7GPpgp_Rl :=  /*line :1*/ recover(); l7GPpgp_Rl != nil {
		if u2_owHMKgs, xP5nLuM9_y := l7GPpgp_Rl.(uJ4X71BjaNv4); xP5nLuM9_y {
			*yUhNNLjWzH8j = u2_owHMKgs.i_VXKKvJ
		} else if u1OEE0, xP5nLuM9_y := l7GPpgp_Rl.(error); xP5nLuM9_y && u1OEE0 == io.K5Sqco {
			*yUhNNLjWzH8j = u1OEE0
		} else {
			 /*line :1*/ panic(l7GPpgp_Rl)
		}
	}
}


func (a_CBqX *rrkIecgD) oe7XLd2oHFat(nghaNY []any) (mswcyBqTo int, jDbNTz1lC_ error) {
	defer  /*line :1*/ qU7_PYRCR(&jDbNTz1lC_)
	for _, oaT9vLaMI3 := range nghaNY {
		 /*line :1*/ a_CBqX.ez1E6lSHl('v', oaT9vLaMI3)
		mswcyBqTo++
	}

	if a_CBqX.oYKqyG {
		for {
			vuT9drijWK :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
			if vuT9drijWK == '\n' || vuT9drijWK == eof {
				break
			}
			if ! /*line :1*/ hZLGRorYw1yo(vuT9drijWK) {
				 /*line :1*/ a_CBqX.d1S01Gz("expected newline")
				break
			}
		}
	}
	return
}









func (a_CBqX *rrkIecgD) yvUv0XAdE(aCNXxCRXOS5o string) (wKUgq0A int) {
	for wKUgq0A <  /*line :1*/ len(aCNXxCRXOS5o) {
		_aQieps38bx, dwIESiwPi :=  /*line :1*/ utf8.HVAV7aTqxzg(aCNXxCRXOS5o[wKUgq0A:])

		if  /*line :1*/ hZLGRorYw1yo(_aQieps38bx) {
			sn0AIk := 0
			rvG1itb0lXq := false
			for  /*line :1*/ hZLGRorYw1yo(_aQieps38bx) && wKUgq0A <  /*line :1*/ len(aCNXxCRXOS5o) {
				if _aQieps38bx == '\n' {
					sn0AIk++
					rvG1itb0lXq = false
				} else {
					rvG1itb0lXq = true
				}
				wKUgq0A += dwIESiwPi
				_aQieps38bx, dwIESiwPi =  /*line :1*/ utf8.HVAV7aTqxzg(aCNXxCRXOS5o[wKUgq0A:])
			}
			for rDsmG2dv53Fv := 0; rDsmG2dv53Fv < sn0AIk; rDsmG2dv53Fv++ {
				quZ_f_ :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
				for  /*line :1*/ hZLGRorYw1yo(quZ_f_) && quZ_f_ != '\n' {
					quZ_f_ =  /*line :1*/ a_CBqX.tv4NHbyxyO()
				}
				if quZ_f_ != '\n' && quZ_f_ != eof {
					 /*line :1*/ a_CBqX.d1S01Gz("newline in format does not match input")
				}
			}
			if rvG1itb0lXq {
				quZ_f_ :=  /*line :1*/ a_CBqX.tv4NHbyxyO()
				if sn0AIk == 0 {

					if ! /*line :1*/ hZLGRorYw1yo(quZ_f_) && quZ_f_ != eof {
						 /*line :1*/ a_CBqX.d1S01Gz("expected space in input to match format")
					}
					if quZ_f_ == '\n' {
						 /*line :1*/ a_CBqX.d1S01Gz("newline in input does not match format")
					}
				}
				for  /*line :1*/ hZLGRorYw1yo(quZ_f_) && quZ_f_ != '\n' {
					quZ_f_ =  /*line :1*/ a_CBqX.tv4NHbyxyO()
				}
				if quZ_f_ != eof {
					 /*line :1*/ a_CBqX.UnreadRune()
				}
			}
			continue
		}

		if _aQieps38bx == '%' {

			if wKUgq0A+dwIESiwPi ==  /*line :1*/ len(aCNXxCRXOS5o) {
				 /*line :1*/ a_CBqX.d1S01Gz("missing verb: % at end of format string")
			}

			u7AVjyhkX, _ :=  /*line :1*/ utf8.HVAV7aTqxzg(aCNXxCRXOS5o[wKUgq0A+dwIESiwPi:])
			if u7AVjyhkX != '%' {
				return
			}
			wKUgq0A += dwIESiwPi
		}

		quZ_f_ :=  /*line :1*/ a_CBqX.hJhaYl2fDSR()
		if _aQieps38bx != quZ_f_ {
			 /*line :1*/ a_CBqX.UnreadRune()
			return -1
		}
		wKUgq0A += dwIESiwPi
	}
	return
}



func (a_CBqX *rrkIecgD) rFHTfa0M3(aCNXxCRXOS5o string, nghaNY []any) (mswcyBqTo int, jDbNTz1lC_ error) {
	defer  /*line :1*/ qU7_PYRCR(&jDbNTz1lC_)
	dA0LK4aPCOv :=  /*line :1*/ len(aCNXxCRXOS5o) - 1

	for wKUgq0A := 0; wKUgq0A <= dA0LK4aPCOv; {
		dwIESiwPi :=  /*line :1*/ a_CBqX.yvUv0XAdE(aCNXxCRXOS5o[wKUgq0A:])
		if dwIESiwPi > 0 {
			wKUgq0A += dwIESiwPi
			continue
		}

		if aCNXxCRXOS5o[wKUgq0A] != '%' {

			if dwIESiwPi < 0 {
				 /*line :1*/ a_CBqX.d1S01Gz("input does not match format")
			}

			break
		}
		wKUgq0A++

		
		var i4YvbF bool
		a_CBqX.gae68Pifxa, i4YvbF, wKUgq0A =  /*line :1*/ bWbu9qxG4bF(aCNXxCRXOS5o, wKUgq0A, dA0LK4aPCOv)
		if !i4YvbF {
			a_CBqX.gae68Pifxa = hugeWid
		}

		gaq8WME, dwIESiwPi :=  /*line :1*/ utf8.HVAV7aTqxzg(aCNXxCRXOS5o[wKUgq0A:])
		wKUgq0A += dwIESiwPi

		if gaq8WME != 'c' {
			 /*line :1*/ a_CBqX.SkipSpace()
		}
		if gaq8WME == '%' {
			 /*line :1*/ a_CBqX.kPMngA3nd()
			continue
		}
		a_CBqX.iJsthwzTgRa = a_CBqX.ci4KQEj8tjJ
		if zEHZ5ak0_jf := a_CBqX.h90JLRma + a_CBqX.gae68Pifxa; zEHZ5ak0_jf < a_CBqX.iJsthwzTgRa {
			a_CBqX.iJsthwzTgRa = zEHZ5ak0_jf
		}

		if mswcyBqTo >=  /*line :1*/ len(nghaNY) {
			 /*line :1*/ a_CBqX.d1S01Gz("too few operands for format '%" + aCNXxCRXOS5o[wKUgq0A-dwIESiwPi:] + "'")
			break
		}
		oaT9vLaMI3 := nghaNY[mswcyBqTo]

		 /*line :1*/ a_CBqX.ez1E6lSHl(gaq8WME, oaT9vLaMI3)
		mswcyBqTo++
		a_CBqX.iJsthwzTgRa = a_CBqX.ci4KQEj8tjJ
	}
	if mswcyBqTo <  /*line :1*/ len(nghaNY) {
		 /*line :1*/ a_CBqX.d1S01Gz("too many operands")
	}
	return
}
