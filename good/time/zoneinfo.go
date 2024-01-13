//line :1
package fRAfQd_

import (
	errors "iAHoxjmM"
	sync "sync"
	syscall "bUKeamF"
)

//go:generate env ZONEINFO=$GOROOT/lib/time/zoneinfo.zip go run genzabbrs.go -output zoneinfo_abbrs_windows.go

type Hh4U1oidou struct {
	yiQ5l0TL_	string
	beFgWpkl8XT0	[]k3sFK3l
	kAx6aUL	[]ep4v_A

	dY_Wjphz	string

	tIAdDDIq	int64
	cmASbbDbiav	int64
	j9a8yEma	*k3sFK3l
}

type k3sFK3l struct {
	sn0AXnjD	string
	a8VzUELX	int
	eXRNnAc1aY	bool
}

type ep4v_A struct {
	onK9uWf7T8	int64
	kRTXiPL	uint8
	cl2ViiNsIZu, r8_hPujy4	bool
}

const (
	alpha	= -1 << 63
	omega	= 1<<63 - 1
)

var PD1NIUjTJ *Hh4U1oidou = &cYYjtVboaU

var cYYjtVboaU = Hh4U1oidou{yiQ5l0TL_: "UTC"}

var CH9afbXxKDfu *Hh4U1oidou = &s45BJX

var s45BJX Hh4U1oidou
var eul0iggnFg sync.LhBfwn6wa1x

func (ipt7urdFIkH *Hh4U1oidou) w9w0xFJlcOz() *Hh4U1oidou {
	if ipt7urdFIkH == nil {
		return &cYYjtVboaU
	}
	if ipt7urdFIkH == &s45BJX {
		 /*line :1*/ eul0iggnFg.Do(aflnY9OXPZ)
	}
	return ipt7urdFIkH
}

func (ipt7urdFIkH *Hh4U1oidou) String() string {
	return  /*line :1*/ ipt7urdFIkH.w9w0xFJlcOz().yiQ5l0TL_
}

var xaY1eSVFOc []*Hh4U1oidou
var caeqLGY51vC sync.LhBfwn6wa1x

func Z586lZG(w12CVdh9 string, l7D_bh2w int) *Hh4U1oidou {

	const hoursBeforeUTC = 12
	const hoursAfterUTC = 14
	xmMDTE := l7D_bh2w / 60 / 60
	if w12CVdh9 == "" && -hoursBeforeUTC <= xmMDTE && xmMDTE <= +hoursAfterUTC && xmMDTE*60*60 == l7D_bh2w {
		 /*line :1*/ caeqLGY51vC.Do(func() {
			xaY1eSVFOc =  /*line :1*/ make([]*Hh4U1oidou, hoursBeforeUTC+1+hoursAfterUTC)
			for ggP7eA6hbC1 := -hoursBeforeUTC; ggP7eA6hbC1 <= +hoursAfterUTC; ggP7eA6hbC1++ {
				xaY1eSVFOc[ggP7eA6hbC1+hoursBeforeUTC] =  /*line :1*/ qwhqo0T("", ggP7eA6hbC1*60*60)
			}
		})
		return xaY1eSVFOc[xmMDTE+hoursBeforeUTC]
	}
	return  /*line :1*/ qwhqo0T(w12CVdh9, l7D_bh2w)
}

func qwhqo0T(w12CVdh9 string, l7D_bh2w int) *Hh4U1oidou {
	ipt7urdFIkH := &Hh4U1oidou{
		yiQ5l0TL_:	w12CVdh9,
		beFgWpkl8XT0:	[]k3sFK3l{{w12CVdh9, l7D_bh2w, false}},
		kAx6aUL:	[]ep4v_A{{alpha, 0, false, false}},
		tIAdDDIq:	alpha,
		cmASbbDbiav:	omega,
	}
	ipt7urdFIkH.j9a8yEma = &ipt7urdFIkH.beFgWpkl8XT0[0]
	return ipt7urdFIkH
}

func (ipt7urdFIkH *Hh4U1oidou) oPdG5Tk5(insjC2Va int64) (w12CVdh9 string, l7D_bh2w int, cAZmDfINA, wo7u8LY9W int64, kVCaCmlEp1vW bool) {
	ipt7urdFIkH =  /*line :1*/ ipt7urdFIkH.w9w0xFJlcOz()

	if  /*line :1*/ len(ipt7urdFIkH.beFgWpkl8XT0) == 0 {
		w12CVdh9 = "UTC"
		l7D_bh2w = 0
		cAZmDfINA = alpha
		wo7u8LY9W = omega
		kVCaCmlEp1vW = false
		return
	}

	if k3sFK3l := ipt7urdFIkH.j9a8yEma; k3sFK3l != nil && ipt7urdFIkH.tIAdDDIq <= insjC2Va && insjC2Va < ipt7urdFIkH.cmASbbDbiav {
		w12CVdh9 = k3sFK3l.sn0AXnjD
		l7D_bh2w = k3sFK3l.a8VzUELX
		cAZmDfINA = ipt7urdFIkH.tIAdDDIq
		wo7u8LY9W = ipt7urdFIkH.cmASbbDbiav
		kVCaCmlEp1vW = k3sFK3l.eXRNnAc1aY
		return
	}

	if  /*line :1*/ len(ipt7urdFIkH.kAx6aUL) == 0 || insjC2Va < ipt7urdFIkH.kAx6aUL[0].onK9uWf7T8 {
		k3sFK3l := &ipt7urdFIkH.beFgWpkl8XT0[ /*line :1*/ ipt7urdFIkH.vEGrwhD()]
		w12CVdh9 = k3sFK3l.sn0AXnjD
		l7D_bh2w = k3sFK3l.a8VzUELX
		cAZmDfINA = alpha
		if  /*line :1*/ len(ipt7urdFIkH.kAx6aUL) > 0 {
			wo7u8LY9W = ipt7urdFIkH.kAx6aUL[0].onK9uWf7T8
		} else {
			wo7u8LY9W = omega
		}
		kVCaCmlEp1vW = k3sFK3l.eXRNnAc1aY
		return
	}

	e7MbtYTeSIq := ipt7urdFIkH.kAx6aUL
	wo7u8LY9W = omega
	xW2tbzZ6pY := 0
	msqan9KLw :=  /*line :1*/ len(e7MbtYTeSIq)
	for msqan9KLw-xW2tbzZ6pY > 1 {
		fYy09z7m := xW2tbzZ6pY + (msqan9KLw-xW2tbzZ6pY)/2
		lCMGgjwpk := e7MbtYTeSIq[fYy09z7m].onK9uWf7T8
		if insjC2Va < lCMGgjwpk {
			wo7u8LY9W = lCMGgjwpk
			msqan9KLw = fYy09z7m
		} else {
			xW2tbzZ6pY = fYy09z7m
		}
	}
	k3sFK3l := &ipt7urdFIkH.beFgWpkl8XT0[e7MbtYTeSIq[xW2tbzZ6pY].kRTXiPL]
	w12CVdh9 = k3sFK3l.sn0AXnjD
	l7D_bh2w = k3sFK3l.a8VzUELX
	cAZmDfINA = e7MbtYTeSIq[xW2tbzZ6pY].onK9uWf7T8

	kVCaCmlEp1vW = k3sFK3l.eXRNnAc1aY

	if xW2tbzZ6pY ==  /*line :1*/ len(e7MbtYTeSIq)-1 && ipt7urdFIkH.dY_Wjphz != "" {
		if fLoiU5y, iizgobY3, xBMN0J, fHCLSS9g, aCQcv2, ufWfJe3qPw :=  /*line :1*/ dMa5ZLv6_o(ipt7urdFIkH.dY_Wjphz, cAZmDfINA, insjC2Va); ufWfJe3qPw {
			return fLoiU5y, iizgobY3, xBMN0J, fHCLSS9g, aCQcv2
		}
	}

	return
}

func (ipt7urdFIkH *Hh4U1oidou) vEGrwhD() int {

	if ! /*line :1*/ ipt7urdFIkH.fkXTgnxJwmR() {
		return 0
	}

	if  /*line :1*/ len(ipt7urdFIkH.kAx6aUL) > 0 && ipt7urdFIkH.beFgWpkl8XT0[ipt7urdFIkH.kAx6aUL[0].kRTXiPL].eXRNnAc1aY {
		for mJaAiK2Hke :=  /*line :1*/ int(ipt7urdFIkH.kAx6aUL[0].kRTXiPL) - 1; mJaAiK2Hke >= 0; mJaAiK2Hke-- {
			if !ipt7urdFIkH.beFgWpkl8XT0[mJaAiK2Hke].eXRNnAc1aY {
				return mJaAiK2Hke
			}
		}
	}

	for mJaAiK2Hke := range ipt7urdFIkH.beFgWpkl8XT0 {
		if !ipt7urdFIkH.beFgWpkl8XT0[mJaAiK2Hke].eXRNnAc1aY {
			return mJaAiK2Hke
		}
	}

	return 0
}

func (ipt7urdFIkH *Hh4U1oidou) fkXTgnxJwmR() bool {
	for _, e7MbtYTeSIq := range ipt7urdFIkH.kAx6aUL {
		if e7MbtYTeSIq.kRTXiPL == 0 {
			return true
		}
	}
	return false
}

func dMa5ZLv6_o(cjQd5EdPBZv string, qK0hXMFDzRX, insjC2Va int64) (w12CVdh9 string, l7D_bh2w int, cAZmDfINA, wo7u8LY9W int64, kVCaCmlEp1vW, ufWfJe3qPw bool) {
	var (
		rOWD9G, edZ1_fT0	string
		d1mx5z4N1z, eg9SQLVH3	int
	)

	rOWD9G, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ ce2Le6c9ZBJD(cjQd5EdPBZv)
	if ufWfJe3qPw {
		d1mx5z4N1z, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ y3UIf8Fani(cjQd5EdPBZv)
	}
	if !ufWfJe3qPw {
		return "", 0, 0, 0, false, false
	}

	d1mx5z4N1z = -d1mx5z4N1z

	if  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] == ',' {

		return rOWD9G, d1mx5z4N1z, qK0hXMFDzRX, omega, false, true
	}

	edZ1_fT0, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ ce2Le6c9ZBJD(cjQd5EdPBZv)
	if ufWfJe3qPw {
		if  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] == ',' {
			eg9SQLVH3 = d1mx5z4N1z + secondsPerHour
		} else {
			eg9SQLVH3, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ y3UIf8Fani(cjQd5EdPBZv)
			eg9SQLVH3 = -eg9SQLVH3
		}
	}
	if !ufWfJe3qPw {
		return "", 0, 0, 0, false, false
	}

	if  /*line :1*/ len(cjQd5EdPBZv) == 0 {

		cjQd5EdPBZv = ",M3.2.0,M11.1.0"
	}

	if cjQd5EdPBZv[0] != ',' && cjQd5EdPBZv[0] != ';' {
		return "", 0, 0, 0, false, false
	}
	cjQd5EdPBZv = cjQd5EdPBZv[1:]

	var e75dVSSjirbs, gGLYjdxy0E x4BYleIfLdel
	e75dVSSjirbs, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ pUvb3RxFeL(cjQd5EdPBZv)
	if !ufWfJe3qPw ||  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] != ',' {
		return "", 0, 0, 0, false, false
	}
	cjQd5EdPBZv = cjQd5EdPBZv[1:]
	gGLYjdxy0E, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ pUvb3RxFeL(cjQd5EdPBZv)
	if !ufWfJe3qPw ||  /*line :1*/ len(cjQd5EdPBZv) > 0 {
		return "", 0, 0, 0, false, false
	}

	mO2sH8hUHnW, _, _, rAwk8hU1NKXu :=  /*line :1*/ fdO09EiS( /*line :1*/ uint64(insjC2Va+unixToInternal+internalToAbsolute), false)

	hwFyoTwcUvu :=  /*line :1*/ int64(rAwk8hU1NKXu*secondsPerDay) + insjC2Va%secondsPerDay

	lNAwkbJQVBgU :=  /*line :1*/ wESvrx2(mO2sH8hUHnW)
	nsu4XSEmD7r :=  /*line :1*/ int64(lNAwkbJQVBgU * secondsPerDay)
	nsu4XSEmD7r += absoluteToInternal + internalToUnix

	iqaLrA :=  /*line :1*/ int64( /*line :1*/ iN8mRvp8Sx(mO2sH8hUHnW, e75dVSSjirbs, d1mx5z4N1z))
	ftQIVp8 :=  /*line :1*/ int64( /*line :1*/ iN8mRvp8Sx(mO2sH8hUHnW, gGLYjdxy0E, eg9SQLVH3))
	rSwbqh4JKgVL, qI10NJ := true, false

	if ftQIVp8 < iqaLrA {
		iqaLrA, ftQIVp8 = ftQIVp8, iqaLrA
		rOWD9G, edZ1_fT0 = edZ1_fT0, rOWD9G
		d1mx5z4N1z, eg9SQLVH3 = eg9SQLVH3, d1mx5z4N1z
		qI10NJ, rSwbqh4JKgVL = rSwbqh4JKgVL, qI10NJ
	}

	if hwFyoTwcUvu < iqaLrA {
		return rOWD9G, d1mx5z4N1z, nsu4XSEmD7r, iqaLrA + nsu4XSEmD7r, qI10NJ, true
	} else if hwFyoTwcUvu >= ftQIVp8 {
		return rOWD9G, d1mx5z4N1z, ftQIVp8 + nsu4XSEmD7r, nsu4XSEmD7r + 365*secondsPerDay, qI10NJ, true
	} else {
		return edZ1_fT0, eg9SQLVH3, iqaLrA + nsu4XSEmD7r, ftQIVp8 + nsu4XSEmD7r, rSwbqh4JKgVL, true
	}
}

func ce2Le6c9ZBJD(cjQd5EdPBZv string) (string, string, bool) {
	if  /*line :1*/ len(cjQd5EdPBZv) == 0 {
		return "", "", false
	}
	if cjQd5EdPBZv[0] != '<' {
		for eUo6AxKM7, aLhBMRzpWd := range cjQd5EdPBZv {
			switch aLhBMRzpWd {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ',', '-', '+':
				if eUo6AxKM7 < 3 {
					return "", "", false
				}
				return cjQd5EdPBZv[:eUo6AxKM7], cjQd5EdPBZv[eUo6AxKM7:], true
			}
		}
		if  /*line :1*/ len(cjQd5EdPBZv) < 3 {
			return "", "", false
		}
		return cjQd5EdPBZv, "", true
	} else {
		for eUo6AxKM7, aLhBMRzpWd := range cjQd5EdPBZv {
			if aLhBMRzpWd == '>' {
				return cjQd5EdPBZv[1:eUo6AxKM7], cjQd5EdPBZv[eUo6AxKM7+1:], true
			}
		}
		return "", "", false
	}
}

func y3UIf8Fani(cjQd5EdPBZv string) (l7D_bh2w int, m9fsb1Me string, ufWfJe3qPw bool) {
	if  /*line :1*/ len(cjQd5EdPBZv) == 0 {
		return 0, "", false
	}
	yY3pgXWox3j := false
	if cjQd5EdPBZv[0] == '+' {
		cjQd5EdPBZv = cjQd5EdPBZv[1:]
	} else if cjQd5EdPBZv[0] == '-' {
		cjQd5EdPBZv = cjQd5EdPBZv[1:]
		yY3pgXWox3j = true
	}

	var a6Hjy9ll int
	a6Hjy9ll, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv, 0, 24*7)
	if !ufWfJe3qPw {
		return 0, "", false
	}
	mP6vzfnl := a6Hjy9ll * secondsPerHour
	if  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] != ':' {
		if yY3pgXWox3j {
			mP6vzfnl = -mP6vzfnl
		}
		return mP6vzfnl, cjQd5EdPBZv, true
	}

	var y5WzJ6 int
	y5WzJ6, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv[1:], 0, 59)
	if !ufWfJe3qPw {
		return 0, "", false
	}
	mP6vzfnl += y5WzJ6 * secondsPerMinute
	if  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] != ':' {
		if yY3pgXWox3j {
			mP6vzfnl = -mP6vzfnl
		}
		return mP6vzfnl, cjQd5EdPBZv, true
	}

	var ksGjQpXl int
	ksGjQpXl, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv[1:], 0, 59)
	if !ufWfJe3qPw {
		return 0, "", false
	}
	mP6vzfnl += ksGjQpXl

	if yY3pgXWox3j {
		mP6vzfnl = -mP6vzfnl
	}
	return mP6vzfnl, cjQd5EdPBZv, true
}

type d_iatllCsw int

const (
	ruleJulian	d_iatllCsw	= iota
	ruleDOY
	ruleMonthWeekDay
)

type x4BYleIfLdel struct {
	szyV1XZvIZE	d_iatllCsw
	bFBxMALNbv	int
	p0bzap3w	int
	y9SrRBPIylHB	int
	muOawsBzO4Z	int
}

func pUvb3RxFeL(cjQd5EdPBZv string) (x4BYleIfLdel, string, bool) {
	var aLhBMRzpWd x4BYleIfLdel
	if  /*line :1*/ len(cjQd5EdPBZv) == 0 {
		return x4BYleIfLdel{}, "", false
	}
	ufWfJe3qPw := false
	if cjQd5EdPBZv[0] == 'J' {
		var gmAp83Am int
		gmAp83Am, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv[1:], 1, 365)
		if !ufWfJe3qPw {
			return x4BYleIfLdel{}, "", false
		}
		aLhBMRzpWd.szyV1XZvIZE = ruleJulian
		aLhBMRzpWd.bFBxMALNbv = gmAp83Am
	} else if cjQd5EdPBZv[0] == 'M' {
		var z_g7B_o int
		z_g7B_o, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv[1:], 1, 12)
		if !ufWfJe3qPw ||  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] != '.' {
			return x4BYleIfLdel{}, "", false

		}
		var cDtaLw1Mj int
		cDtaLw1Mj, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv[1:], 1, 5)
		if !ufWfJe3qPw ||  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] != '.' {
			return x4BYleIfLdel{}, "", false
		}
		var gv5Z468R int
		gv5Z468R, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv[1:], 0, 6)
		if !ufWfJe3qPw {
			return x4BYleIfLdel{}, "", false
		}
		aLhBMRzpWd.szyV1XZvIZE = ruleMonthWeekDay
		aLhBMRzpWd.bFBxMALNbv = gv5Z468R
		aLhBMRzpWd.p0bzap3w = cDtaLw1Mj
		aLhBMRzpWd.y9SrRBPIylHB = z_g7B_o
	} else {
		var gv5Z468R int
		gv5Z468R, cjQd5EdPBZv, ufWfJe3qPw =  /*line :1*/ s0l3MMwcMh(cjQd5EdPBZv, 0, 365)
		if !ufWfJe3qPw {
			return x4BYleIfLdel{}, "", false
		}
		aLhBMRzpWd.szyV1XZvIZE = ruleDOY
		aLhBMRzpWd.bFBxMALNbv = gv5Z468R
	}

	if  /*line :1*/ len(cjQd5EdPBZv) == 0 || cjQd5EdPBZv[0] != '/' {
		aLhBMRzpWd.muOawsBzO4Z = 2 * secondsPerHour
		return aLhBMRzpWd, cjQd5EdPBZv, true
	}

	l7D_bh2w, cjQd5EdPBZv, ufWfJe3qPw :=  /*line :1*/ y3UIf8Fani(cjQd5EdPBZv[1:])
	if !ufWfJe3qPw {
		return x4BYleIfLdel{}, "", false
	}
	aLhBMRzpWd.muOawsBzO4Z = l7D_bh2w

	return aLhBMRzpWd, cjQd5EdPBZv, true
}

func s0l3MMwcMh(cjQd5EdPBZv string, qJ3jbB5O, aRi90tFb0vWu int) (fZmGCPzk0RA int, m9fsb1Me string, ufWfJe3qPw bool) {
	if  /*line :1*/ len(cjQd5EdPBZv) == 0 {
		return 0, "", false
	}
	fZmGCPzk0RA = 0
	for eUo6AxKM7, aLhBMRzpWd := range cjQd5EdPBZv {
		if aLhBMRzpWd < '0' || aLhBMRzpWd > '9' {
			if eUo6AxKM7 == 0 || fZmGCPzk0RA < qJ3jbB5O {
				return 0, "", false
			}
			return fZmGCPzk0RA, cjQd5EdPBZv[eUo6AxKM7:], true
		}
		fZmGCPzk0RA *= 10
		fZmGCPzk0RA +=  /*line :1*/ int(aLhBMRzpWd) - '0'
		if fZmGCPzk0RA > aRi90tFb0vWu {
			return 0, "", false
		}
	}
	if fZmGCPzk0RA < qJ3jbB5O {
		return 0, "", false
	}
	return fZmGCPzk0RA, "", true
}

func iN8mRvp8Sx(mO2sH8hUHnW int, aLhBMRzpWd x4BYleIfLdel, mP6vzfnl int) int {
	var cjQd5EdPBZv int
	switch aLhBMRzpWd.szyV1XZvIZE {
	case ruleJulian:
		cjQd5EdPBZv = (aLhBMRzpWd.bFBxMALNbv - 1) * secondsPerDay
		if  /*line :1*/ dJS2LEP6(mO2sH8hUHnW) && aLhBMRzpWd.bFBxMALNbv >= 60 {
			cjQd5EdPBZv += secondsPerDay
		}
	case ruleDOY:
		cjQd5EdPBZv = aLhBMRzpWd.bFBxMALNbv * secondsPerDay
	case ruleMonthWeekDay:

		s7ZszeDNuqt := (aLhBMRzpWd.y9SrRBPIylHB+9)%12 + 1
		aw_xbm0 := mO2sH8hUHnW
		if aLhBMRzpWd.y9SrRBPIylHB <= 2 {
			aw_xbm0--
		}
		hR7_iack := aw_xbm0 / 100
		ka1Lfe := aw_xbm0 % 100
		gyzJ6Qe151 := ((26*s7ZszeDNuqt-2)/10 + 1 + ka1Lfe + ka1Lfe/4 + hR7_iack/4 - 2*hR7_iack) % 7
		if gyzJ6Qe151 < 0 {
			gyzJ6Qe151 += 7
		}

		lNAwkbJQVBgU := aLhBMRzpWd.bFBxMALNbv - gyzJ6Qe151
		if lNAwkbJQVBgU < 0 {
			lNAwkbJQVBgU += 7
		}
		for eUo6AxKM7 := 1; eUo6AxKM7 < aLhBMRzpWd.p0bzap3w; eUo6AxKM7++ {
			if lNAwkbJQVBgU+7 >=  /*line :1*/ rs4bvng4W6y( /*line :1*/ ZyPFXNxLcwpT(aLhBMRzpWd.y9SrRBPIylHB), mO2sH8hUHnW) {
				break
			}
			lNAwkbJQVBgU += 7
		}
		lNAwkbJQVBgU +=  /*line :1*/ int(h7KunS2pF[aLhBMRzpWd.y9SrRBPIylHB-1])
		if  /*line :1*/ dJS2LEP6(mO2sH8hUHnW) && aLhBMRzpWd.y9SrRBPIylHB > 2 {
			lNAwkbJQVBgU++
		}
		cjQd5EdPBZv = lNAwkbJQVBgU * secondsPerDay
	}

	return cjQd5EdPBZv + aLhBMRzpWd.muOawsBzO4Z - mP6vzfnl
}

func (ipt7urdFIkH *Hh4U1oidou) hHPWBNyM(w12CVdh9 string, xqgBUhRs int64) (l7D_bh2w int, ufWfJe3qPw bool) {
	ipt7urdFIkH =  /*line :1*/ ipt7urdFIkH.w9w0xFJlcOz()

	for eUo6AxKM7 := range ipt7urdFIkH.beFgWpkl8XT0 {
		k3sFK3l := &ipt7urdFIkH.beFgWpkl8XT0[eUo6AxKM7]
		if k3sFK3l.sn0AXnjD == w12CVdh9 {
			zknuKz, l7D_bh2w, _, _, _ :=  /*line :1*/ ipt7urdFIkH.oPdG5Tk5(xqgBUhRs -  /*line :1*/ int64(k3sFK3l.a8VzUELX))
			if zknuKz == k3sFK3l.sn0AXnjD {
				return l7D_bh2w, true
			}
		}
	}

	for eUo6AxKM7 := range ipt7urdFIkH.beFgWpkl8XT0 {
		k3sFK3l := &ipt7urdFIkH.beFgWpkl8XT0[eUo6AxKM7]
		if k3sFK3l.sn0AXnjD == w12CVdh9 {
			return k3sFK3l.a8VzUELX, true
		}
	}

	return
}

var nWnXF0gFiZ =  /*line :1*/ errors.Su6g6hRoi9X("time: invalid location name")

var rPzU8bkXar *string
var apWl4D4 sync.LhBfwn6wa1x

func TCKVS3k4lUy(w12CVdh9 string) (*Hh4U1oidou, error) {
	if w12CVdh9 == "" || w12CVdh9 == "UTC" {
		return PD1NIUjTJ, nil
	}
	if w12CVdh9 == "Local" {
		return CH9afbXxKDfu, nil
	}
	if  /*line :1*/ rSbvat2M(w12CVdh9) || w12CVdh9[0] == '/' || w12CVdh9[0] == '\\' {

		return nil, nWnXF0gFiZ
	}
	 /*line :1*/ apWl4D4.Do(func() {
		deN9sSkqv, _ :=  /*line :1*/ syscall.ITYf1B("ZONEINFO")
		rPzU8bkXar = &deN9sSkqv
	})
	var exN_IAbg error
	if *rPzU8bkXar != "" {
		if j0viWPPtxg, xuMLYrB :=  /*line :1*/ usBDAK(*rPzU8bkXar, w12CVdh9); xuMLYrB == nil {
			if iL2jRw3Faxe, xuMLYrB :=  /*line :1*/ CBF4fYsE_(w12CVdh9, j0viWPPtxg); xuMLYrB == nil {
				return iL2jRw3Faxe, nil
			}
			exN_IAbg = xuMLYrB
		} else if xuMLYrB != syscall.ENOENT {
			exN_IAbg = xuMLYrB
		}
	}
	if iL2jRw3Faxe, xuMLYrB :=  /*line :1*/ iD3RkIV(w12CVdh9, maKnDxXjO); xuMLYrB == nil {
		return iL2jRw3Faxe, nil
	} else if exN_IAbg == nil {
		exN_IAbg = xuMLYrB
	}
	return nil, exN_IAbg
}

func rSbvat2M(cjQd5EdPBZv string) bool {
	if  /*line :1*/ len(cjQd5EdPBZv) < 2 {
		return false
	}
	for eUo6AxKM7 := 0; eUo6AxKM7 <  /*line :1*/ len(cjQd5EdPBZv)-1; eUo6AxKM7++ {
		if cjQd5EdPBZv[eUo6AxKM7] == '.' && cjQd5EdPBZv[eUo6AxKM7+1] == '.' {
			return true
		}
	}
	return false
}
