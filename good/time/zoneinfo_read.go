//line :1
package fRAfQd_

import (
	errors "iAHoxjmM"
	"runtime"
	syscall "bUKeamF"
)



func ghieMc(k05jcfKvAvw6 func(string) (string, error)) {
	nD9jQajf = k05jcfKvAvw6
}





var nD9jQajf func(zaIiG3yY_kN string) (string, error)




const maxFileSize = 10 << 20

type p7qQXZi string

func (k05jcfKvAvw6 p7qQXZi) Error() string {
	return "time: file " +  /*line :1*/ string(k05jcfKvAvw6) + " is too large"
}


const (
	seekStart	= 0
	seekCurrent	= 1
	seekEnd	= 2
)


type tS8K1OcYVr struct {
	f6lpw5r86	[]byte
	tX8uMvq	bool
}

func (lNAwkbJQVBgU *tS8K1OcYVr) ozmugx0ism(aeBZDJEINbb int) []byte {
	if  /*line :1*/ len(lNAwkbJQVBgU.f6lpw5r86) < aeBZDJEINbb {
		lNAwkbJQVBgU.f6lpw5r86 = nil
		lNAwkbJQVBgU.tX8uMvq = true
		return nil
	}
	_STZiF := lNAwkbJQVBgU.f6lpw5r86[0:aeBZDJEINbb]
	lNAwkbJQVBgU.f6lpw5r86 = lNAwkbJQVBgU.f6lpw5r86[aeBZDJEINbb:]
	return _STZiF
}

func (lNAwkbJQVBgU *tS8K1OcYVr) x_VuRRX() (aeBZDJEINbb uint32, ufWfJe3qPw bool) {
	_STZiF :=  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(4)
	if  /*line :1*/ len(_STZiF) < 4 {
		lNAwkbJQVBgU.tX8uMvq = true
		return 0, false
	}
	return  /*line :1*/ uint32(_STZiF[3]) |  /*line :1*/ uint32(_STZiF[2])<<8 |  /*line :1*/ uint32(_STZiF[1])<<16 |  /*line :1*/ uint32(_STZiF[0])<<24, true
}

func (lNAwkbJQVBgU *tS8K1OcYVr) e689qTmm8() (aeBZDJEINbb uint64, ufWfJe3qPw bool) {
	qHiawKJ2qAr, c3PEm6eqWKz :=  /*line :1*/ lNAwkbJQVBgU.x_VuRRX()
	bkS_EB, i5cNRcAQN :=  /*line :1*/ lNAwkbJQVBgU.x_VuRRX()
	if !c3PEm6eqWKz || !i5cNRcAQN {
		lNAwkbJQVBgU.tX8uMvq = true
		return 0, false
	}
	return ( /*line :1*/ uint64(qHiawKJ2qAr) << 32) |  /*line :1*/ uint64(bkS_EB), true
}

func (lNAwkbJQVBgU *tS8K1OcYVr) j1qUSFrS() (aeBZDJEINbb byte, ufWfJe3qPw bool) {
	_STZiF :=  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(1)
	if  /*line :1*/ len(_STZiF) < 1 {
		lNAwkbJQVBgU.tX8uMvq = true
		return 0, false
	}
	return _STZiF[0], true
}


func (lNAwkbJQVBgU *tS8K1OcYVr) m9fsb1Me() []byte {
	aLhBMRzpWd := lNAwkbJQVBgU.f6lpw5r86
	lNAwkbJQVBgU.f6lpw5r86 = nil
	return aLhBMRzpWd
}


func uhWvA9(_STZiF []byte) string {
	for eUo6AxKM7 := 0; eUo6AxKM7 <  /*line :1*/ len(_STZiF); eUo6AxKM7++ {
		if _STZiF[eUo6AxKM7] == 0 {
			return  /*line :1*/ string(_STZiF[0:eUo6AxKM7])
		}
	}
	return  /*line :1*/ string(_STZiF)
}

var h4uUw7uj1 =  /*line :1*/ errors.Su6g6hRoi9X("malformed time zone information")





func CBF4fYsE_(w12CVdh9 string, a84q4w []byte) (*Hh4U1oidou, error) {
	lNAwkbJQVBgU := tS8K1OcYVr{a84q4w, false}

	if qcQubAM :=  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(4);  /*line :1*/ string(qcQubAM) != "TZif" {
		return nil, h4uUw7uj1
	}

	
	var e02h6We int
	var _STZiF []byte
	if _STZiF =  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(16);  /*line :1*/ len(_STZiF) != 16 {
		return nil, h4uUw7uj1
	} else {
		switch _STZiF[0] {
		case 0:
			e02h6We = 1
		case '2':
			e02h6We = 2
		case '3':
			e02h6We = 3
		default:
			return nil, h4uUw7uj1
		}
	}

	
	
	
	
	
	
	
	const (
		NUTCLocal	= iota
		NStdWall
		NLeap
		NTime
		NZone
		NChar
	)
	var aeBZDJEINbb [6]int
	for eUo6AxKM7 := 0; eUo6AxKM7 < 6; eUo6AxKM7++ {
		gqXBhY, ufWfJe3qPw :=  /*line :1*/ lNAwkbJQVBgU.x_VuRRX()
		if !ufWfJe3qPw {
			return nil, h4uUw7uj1
		}
		if  /*line :1*/ uint32( /*line :1*/ int(gqXBhY)) != gqXBhY {
			return nil, h4uUw7uj1
		}
		aeBZDJEINbb[eUo6AxKM7] =  /*line :1*/ int(gqXBhY)
	}

	xHWCnh3 := false
	if e02h6We > 1 {

		iWW7p4CGpUo := aeBZDJEINbb[NTime]*4 +
			aeBZDJEINbb[NTime] +
			aeBZDJEINbb[NZone]*6 +
			aeBZDJEINbb[NChar] +
			aeBZDJEINbb[NLeap]*8 +
			aeBZDJEINbb[NStdWall] +
			aeBZDJEINbb[NUTCLocal]

		iWW7p4CGpUo += 4 + 16
		 /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(iWW7p4CGpUo)

		xHWCnh3 = true

		for eUo6AxKM7 := 0; eUo6AxKM7 < 6; eUo6AxKM7++ {
			gqXBhY, ufWfJe3qPw :=  /*line :1*/ lNAwkbJQVBgU.x_VuRRX()
			if !ufWfJe3qPw {
				return nil, h4uUw7uj1
			}
			if  /*line :1*/ uint32( /*line :1*/ int(gqXBhY)) != gqXBhY {
				return nil, h4uUw7uj1
			}
			aeBZDJEINbb[eUo6AxKM7] =  /*line :1*/ int(gqXBhY)
		}
	}

	hfCqmMXTt := 4
	if xHWCnh3 {
		hfCqmMXTt = 8
	}

	mSSAaS6 := tS8K1OcYVr{ /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(aeBZDJEINbb[NTime] * hfCqmMXTt), false}

	cAvfn3y :=  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(aeBZDJEINbb[NTime])

	kUOMzzpc3g := tS8K1OcYVr{ /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(aeBZDJEINbb[NZone] * 6), false}

	c91HWD_ :=  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(aeBZDJEINbb[NChar])

	 /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(aeBZDJEINbb[NLeap] * (hfCqmMXTt + 4))

	a7Ffyu0wRV5 :=  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(aeBZDJEINbb[NStdWall])

	hvws5ha :=  /*line :1*/ lNAwkbJQVBgU.ozmugx0ism(aeBZDJEINbb[NUTCLocal])

	if lNAwkbJQVBgU.tX8uMvq {
		return nil, h4uUw7uj1
	}

	var kqbday string
	m9fsb1Me :=  /*line :1*/ lNAwkbJQVBgU.m9fsb1Me()
	if  /*line :1*/ len(m9fsb1Me) > 2 && m9fsb1Me[0] == '\n' && m9fsb1Me[ /*line :1*/ len(m9fsb1Me)-1] == '\n' {
		kqbday =  /*line :1*/ string(m9fsb1Me[1 :  /*line :1*/ len(m9fsb1Me)-1])
	}

	b4zZoH2FE1 := aeBZDJEINbb[NZone]
	if b4zZoH2FE1 == 0 {

		return nil, h4uUw7uj1
	}
	bJQzPjVfN :=  /*line :1*/ make([]k3sFK3l, b4zZoH2FE1)
	for eUo6AxKM7 := range bJQzPjVfN {
		var ufWfJe3qPw bool
		var aeBZDJEINbb uint32
		if aeBZDJEINbb, ufWfJe3qPw =  /*line :1*/ kUOMzzpc3g.x_VuRRX(); !ufWfJe3qPw {
			return nil, h4uUw7uj1
		}
		if  /*line :1*/ uint32( /*line :1*/ int(aeBZDJEINbb)) != aeBZDJEINbb {
			return nil, h4uUw7uj1
		}
		bJQzPjVfN[eUo6AxKM7].a8VzUELX =  /*line :1*/ int( /*line :1*/ int32(aeBZDJEINbb))
		var ev8znPS9W byte
		if ev8znPS9W, ufWfJe3qPw =  /*line :1*/ kUOMzzpc3g.j1qUSFrS(); !ufWfJe3qPw {
			return nil, h4uUw7uj1
		}
		bJQzPjVfN[eUo6AxKM7].eXRNnAc1aY = ev8znPS9W != 0
		if ev8znPS9W, ufWfJe3qPw =  /*line :1*/ kUOMzzpc3g.j1qUSFrS(); !ufWfJe3qPw ||  /*line :1*/ int(ev8znPS9W) >=  /*line :1*/ len(c91HWD_) {
			return nil, h4uUw7uj1
		}
		bJQzPjVfN[eUo6AxKM7].sn0AXnjD =  /*line :1*/ uhWvA9(c91HWD_[ev8znPS9W:])
		if runtime.GOOS == "aix" &&  /*line :1*/ len(w12CVdh9) > 8 && (w12CVdh9[:8] == "Etc/GMT+" || w12CVdh9[:8] == "Etc/GMT-") {

			if w12CVdh9 != "Etc/GMT+0" {

				bJQzPjVfN[eUo6AxKM7].sn0AXnjD = w12CVdh9[4:]
			}
		}
	}

	e7MbtYTeSIq :=  /*line :1*/ make([]ep4v_A, aeBZDJEINbb[NTime])
	for eUo6AxKM7 := range e7MbtYTeSIq {
		var aeBZDJEINbb int64
		if !xHWCnh3 {
			if ju6BDl, ufWfJe3qPw :=  /*line :1*/ mSSAaS6.x_VuRRX(); !ufWfJe3qPw {
				return nil, h4uUw7uj1
			} else {
				aeBZDJEINbb =  /*line :1*/ int64( /*line :1*/ int32(ju6BDl))
			}
		} else {
			if qJ7HC7v2to, ufWfJe3qPw :=  /*line :1*/ mSSAaS6.e689qTmm8(); !ufWfJe3qPw {
				return nil, h4uUw7uj1
			} else {
				aeBZDJEINbb =  /*line :1*/ int64(qJ7HC7v2to)
			}
		}
		e7MbtYTeSIq[eUo6AxKM7].onK9uWf7T8 = aeBZDJEINbb
		if  /*line :1*/ int(cAvfn3y[eUo6AxKM7]) >=  /*line :1*/ len(bJQzPjVfN) {
			return nil, h4uUw7uj1
		}
		e7MbtYTeSIq[eUo6AxKM7].kRTXiPL = cAvfn3y[eUo6AxKM7]
		if eUo6AxKM7 <  /*line :1*/ len(a7Ffyu0wRV5) {
			e7MbtYTeSIq[eUo6AxKM7].cl2ViiNsIZu = a7Ffyu0wRV5[eUo6AxKM7] != 0
		}
		if eUo6AxKM7 <  /*line :1*/ len(hvws5ha) {
			e7MbtYTeSIq[eUo6AxKM7].r8_hPujy4 = hvws5ha[eUo6AxKM7] != 0
		}
	}

	if  /*line :1*/ len(e7MbtYTeSIq) == 0 {

		e7MbtYTeSIq =  /*line :1*/ append(e7MbtYTeSIq, ep4v_A{onK9uWf7T8: alpha, kRTXiPL: 0})
	}

	ipt7urdFIkH := &Hh4U1oidou{beFgWpkl8XT0: bJQzPjVfN, kAx6aUL: e7MbtYTeSIq, yiQ5l0TL_: w12CVdh9, dY_Wjphz: kqbday}

	insjC2Va, _, _ :=  /*line :1*/ sNTZrc2lj()
	for eUo6AxKM7 := range e7MbtYTeSIq {
		if e7MbtYTeSIq[eUo6AxKM7].onK9uWf7T8 <= insjC2Va && (eUo6AxKM7+1 ==  /*line :1*/ len(e7MbtYTeSIq) || insjC2Va < e7MbtYTeSIq[eUo6AxKM7+1].onK9uWf7T8) {
			ipt7urdFIkH.tIAdDDIq = e7MbtYTeSIq[eUo6AxKM7].onK9uWf7T8
			ipt7urdFIkH.cmASbbDbiav = omega
			ipt7urdFIkH.j9a8yEma = &ipt7urdFIkH.beFgWpkl8XT0[e7MbtYTeSIq[eUo6AxKM7].kRTXiPL]
			if eUo6AxKM7+1 <  /*line :1*/ len(e7MbtYTeSIq) {
				ipt7urdFIkH.cmASbbDbiav = e7MbtYTeSIq[eUo6AxKM7+1].onK9uWf7T8
			} else if ipt7urdFIkH.dY_Wjphz != "" {

				if w12CVdh9, l7D_bh2w, xBMN0J, fHCLSS9g, kVCaCmlEp1vW, ufWfJe3qPw :=  /*line :1*/ dMa5ZLv6_o(ipt7urdFIkH.dY_Wjphz, ipt7urdFIkH.tIAdDDIq, insjC2Va); ufWfJe3qPw {
					ipt7urdFIkH.tIAdDDIq = xBMN0J
					ipt7urdFIkH.cmASbbDbiav = fHCLSS9g

					if kqUnObkMAUdR :=  /*line :1*/ hPukbhbSpCzL(ipt7urdFIkH.beFgWpkl8XT0, w12CVdh9, l7D_bh2w, kVCaCmlEp1vW); kqUnObkMAUdR != -1 {
						ipt7urdFIkH.j9a8yEma = &ipt7urdFIkH.beFgWpkl8XT0[kqUnObkMAUdR]
					} else {
						ipt7urdFIkH.j9a8yEma = &k3sFK3l{
							sn0AXnjD:	w12CVdh9,
							a8VzUELX:	l7D_bh2w,
							eXRNnAc1aY:	kVCaCmlEp1vW,
						}
					}
				}
			}
			break
		}
	}

	return ipt7urdFIkH, nil
}

func hPukbhbSpCzL(bJQzPjVfN []k3sFK3l, w12CVdh9 string, l7D_bh2w int, kVCaCmlEp1vW bool) int {
	for eUo6AxKM7, iL2jRw3Faxe := range bJQzPjVfN {
		if iL2jRw3Faxe.sn0AXnjD == w12CVdh9 && iL2jRw3Faxe.a8VzUELX == l7D_bh2w && iL2jRw3Faxe.eXRNnAc1aY == kVCaCmlEp1vW {
			return eUo6AxKM7
		}
	}
	return -1
}



func usBDAK(aadi6Le3p5RN, w12CVdh9 string) ([]byte, error) {
	if  /*line :1*/ len(aadi6Le3p5RN) > 4 && aadi6Le3p5RN[ /*line :1*/ len(aadi6Le3p5RN)-4:] == ".zip" {
		return  /*line :1*/ ndyJjKNBv3BR(aadi6Le3p5RN, w12CVdh9)
	}
	if aadi6Le3p5RN != "" {
		w12CVdh9 = aadi6Le3p5RN + "/" + w12CVdh9
	}
	return  /*line :1*/ dHanxq(w12CVdh9)
}


func x2GIXHcaH(ev8znPS9W []byte) int {
	if  /*line :1*/ len(ev8znPS9W) < 4 {
		return 0
	}
	return  /*line :1*/ int(ev8znPS9W[0]) |  /*line :1*/ int(ev8znPS9W[1])<<8 |  /*line :1*/ int(ev8znPS9W[2])<<16 |  /*line :1*/ int(ev8znPS9W[3])<<24
}


func vdCKrY3GkY(ev8znPS9W []byte) int {
	if  /*line :1*/ len(ev8znPS9W) < 2 {
		return 0
	}
	return  /*line :1*/ int(ev8znPS9W[0]) |  /*line :1*/ int(ev8znPS9W[1])<<8
}



func ndyJjKNBv3BR(cgqgOh, w12CVdh9 string) ([]byte, error) {
	laArf75Y6GN, xuMLYrB :=  /*line :1*/ bavnEgG1Sl(cgqgOh)
	if xuMLYrB != nil {
		return nil, xuMLYrB
	}
	defer  /*line :1*/ hGSAcx81(laArf75Y6GN)

	const (
		zecheader	= 0x06054b50
		zcheader	= 0x02014b50
		ztailsize	= 22

		zheadersize	= 30
		zheader	= 0x04034b50
	)

	qtO6L35R9b :=  /*line :1*/ make([]byte, ztailsize)
	if xuMLYrB :=  /*line :1*/ fUl4zA(laArf75Y6GN, qtO6L35R9b, -ztailsize); xuMLYrB != nil ||  /*line :1*/ x2GIXHcaH(qtO6L35R9b) != zecheader {
		return nil,  /*line :1*/ errors.Su6g6hRoi9X("corrupt zip file " + cgqgOh)
	}
	aeBZDJEINbb :=  /*line :1*/ vdCKrY3GkY(qtO6L35R9b[10:])
	hfCqmMXTt :=  /*line :1*/ x2GIXHcaH(qtO6L35R9b[12:])
	mP6vzfnl :=  /*line :1*/ x2GIXHcaH(qtO6L35R9b[16:])

	qtO6L35R9b =  /*line :1*/ make([]byte, hfCqmMXTt)
	if xuMLYrB :=  /*line :1*/ fUl4zA(laArf75Y6GN, qtO6L35R9b, mP6vzfnl); xuMLYrB != nil {
		return nil,  /*line :1*/ errors.Su6g6hRoi9X("corrupt zip file " + cgqgOh)
	}

	for eUo6AxKM7 := 0; eUo6AxKM7 < aeBZDJEINbb; eUo6AxKM7++ {

		if  /*line :1*/ x2GIXHcaH(qtO6L35R9b) != zcheader {
			break
		}
		ohcNoZGaisV :=  /*line :1*/ vdCKrY3GkY(qtO6L35R9b[10:])
		hfCqmMXTt :=  /*line :1*/ x2GIXHcaH(qtO6L35R9b[24:])
		_GGhAkk :=  /*line :1*/ vdCKrY3GkY(qtO6L35R9b[28:])
		iCo60Ia2rp6B :=  /*line :1*/ vdCKrY3GkY(qtO6L35R9b[30:])
		qiFGf9UqA :=  /*line :1*/ vdCKrY3GkY(qtO6L35R9b[32:])
		mP6vzfnl :=  /*line :1*/ x2GIXHcaH(qtO6L35R9b[42:])
		iwSFObx := qtO6L35R9b[46 : 46+_GGhAkk]
		qtO6L35R9b = qtO6L35R9b[46+_GGhAkk+iCo60Ia2rp6B+qiFGf9UqA:]
		if  /*line :1*/ string(iwSFObx) != w12CVdh9 {
			continue
		}
		if ohcNoZGaisV != 0 {
			return nil,  /*line :1*/ errors.Su6g6hRoi9X("unsupported compression for " + w12CVdh9 + " in " + cgqgOh)
		}

		qtO6L35R9b =  /*line :1*/ make([]byte, zheadersize+_GGhAkk)
		if xuMLYrB :=  /*line :1*/ fUl4zA(laArf75Y6GN, qtO6L35R9b, mP6vzfnl); xuMLYrB != nil ||
			 /*line :1*/ x2GIXHcaH(qtO6L35R9b) != zheader ||
			 /*line :1*/ vdCKrY3GkY(qtO6L35R9b[8:]) != ohcNoZGaisV ||
			 /*line :1*/ vdCKrY3GkY(qtO6L35R9b[26:]) != _GGhAkk ||
			 /*line :1*/ string(qtO6L35R9b[30:30+_GGhAkk]) != w12CVdh9 {
			return nil,  /*line :1*/ errors.Su6g6hRoi9X("corrupt zip file " + cgqgOh)
		}
		iCo60Ia2rp6B =  /*line :1*/ vdCKrY3GkY(qtO6L35R9b[28:])

		qtO6L35R9b =  /*line :1*/ make([]byte, hfCqmMXTt)
		if xuMLYrB :=  /*line :1*/ fUl4zA(laArf75Y6GN, qtO6L35R9b, mP6vzfnl+30+_GGhAkk+iCo60Ia2rp6B); xuMLYrB != nil {
			return nil,  /*line :1*/ errors.Su6g6hRoi9X("corrupt zip file " + cgqgOh)
		}

		return qtO6L35R9b, nil
	}

	return nil, syscall.ENOENT
}




var lZ0Ka9gH8D func(yzBnSt83W, w12CVdh9 string) ([]byte, error)





func dW7uGn(w12CVdh9 string, wSZBMLqT string) ([]byte, error) {
	if  /*line :1*/ len(wSZBMLqT) >= 6 && wSZBMLqT[ /*line :1*/ len(wSZBMLqT)-6:] == "tzdata" {
		return  /*line :1*/ lZ0Ka9gH8D(wSZBMLqT, w12CVdh9)
	}
	return  /*line :1*/ usBDAK(wSZBMLqT, w12CVdh9)
}





func iD3RkIV(w12CVdh9 string, eRjHwnPOPn []string) (iL2jRw3Faxe *Hh4U1oidou, exN_IAbg error) {
	for _, wSZBMLqT := range eRjHwnPOPn {
		j0viWPPtxg, xuMLYrB :=  /*line :1*/ dW7uGn(w12CVdh9, wSZBMLqT)
		if xuMLYrB == nil {
			if iL2jRw3Faxe, xuMLYrB =  /*line :1*/ CBF4fYsE_(w12CVdh9, j0viWPPtxg); xuMLYrB == nil {
				return iL2jRw3Faxe, nil
			}
		}
		if exN_IAbg == nil && xuMLYrB != syscall.ENOENT {
			exN_IAbg = xuMLYrB
		}
	}
	if nD9jQajf != nil {
		j0viWPPtxg, xuMLYrB :=  /*line :1*/ nD9jQajf(w12CVdh9)
		if xuMLYrB == nil {
			if iL2jRw3Faxe, xuMLYrB =  /*line :1*/ CBF4fYsE_(w12CVdh9, [] /*line :1*/ byte(j0viWPPtxg)); xuMLYrB == nil {
				return iL2jRw3Faxe, nil
			}
		}
		if exN_IAbg == nil && xuMLYrB != syscall.ENOENT {
			exN_IAbg = xuMLYrB
		}
	}
	if wSZBMLqT, ufWfJe3qPw :=  /*line :1*/ pDWRrxWhK( /*line :1*/ runtime.GOROOT()); ufWfJe3qPw {
		j0viWPPtxg, xuMLYrB :=  /*line :1*/ dW7uGn(w12CVdh9, wSZBMLqT)
		if xuMLYrB == nil {
			if iL2jRw3Faxe, xuMLYrB =  /*line :1*/ CBF4fYsE_(w12CVdh9, j0viWPPtxg); xuMLYrB == nil {
				return iL2jRw3Faxe, nil
			}
		}
		if exN_IAbg == nil && xuMLYrB != syscall.ENOENT {
			exN_IAbg = xuMLYrB
		}
	}
	if exN_IAbg != nil {
		return nil, exN_IAbg
	}
	return nil,  /*line :1*/ errors.Su6g6hRoi9X("unknown time zone " + w12CVdh9)
}





func dHanxq(w12CVdh9 string) ([]byte, error) {
	k05jcfKvAvw6, xuMLYrB :=  /*line :1*/ bavnEgG1Sl(w12CVdh9)
	if xuMLYrB != nil {
		return nil, xuMLYrB
	}
	defer  /*line :1*/ hGSAcx81(k05jcfKvAvw6)
	var (
		qtO6L35R9b	[4096]byte
		zouGzRS3rJ9X	[]byte
		aeBZDJEINbb	int
	)
	for {
		aeBZDJEINbb, xuMLYrB =  /*line :1*/ ozmugx0ism(k05jcfKvAvw6, qtO6L35R9b[:])
		if aeBZDJEINbb > 0 {
			zouGzRS3rJ9X =  /*line :1*/ append(zouGzRS3rJ9X, qtO6L35R9b[:aeBZDJEINbb]...)
		}
		if aeBZDJEINbb == 0 || xuMLYrB != nil {
			break
		}
		if  /*line :1*/ len(zouGzRS3rJ9X) > maxFileSize {
			return nil,  /*line :1*/ p7qQXZi(w12CVdh9)
		}
	}
	return zouGzRS3rJ9X, xuMLYrB
}
