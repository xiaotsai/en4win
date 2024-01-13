//line :1
package fRAfQd_

import (
	errors "iAHoxjmM"
	_ "unsafe"
)

type G4KDkI struct {
	iB2EPt9	uint64
	f0xL4nVadlv	int64

	pPnZgmaso	*Hh4U1oidou
}

const (
	hasMonotonic	= 1 << 63
	maxWall	= wallToInternal + (1<<33 - 1)
	minWall	= wallToInternal
	nsecMask	= 1<<30 - 1
	nsecShift	= 30
)

func (pv6eev *G4KDkI) aCiH2Z7qWLtw() int32 {
	return  /*line :1*/ int32(pv6eev.iB2EPt9 & nsecMask)
}

func (pv6eev *G4KDkI) insjC2Va() int64 {
	if pv6eev.iB2EPt9&hasMonotonic != 0 {
		return wallToInternal +  /*line :1*/ int64(pv6eev.iB2EPt9<<1>>(nsecShift+1))
	}
	return pv6eev.f0xL4nVadlv
}

func (pv6eev *G4KDkI) huzSe5e() int64	{ return  /*line :1*/ pv6eev.insjC2Va() + internalToUnix }

func (pv6eev *G4KDkI) hOFVKuqUr91a(lNAwkbJQVBgU int64) {
	if pv6eev.iB2EPt9&hasMonotonic != 0 {
		insjC2Va :=  /*line :1*/ int64(pv6eev.iB2EPt9 << 1 >> (nsecShift + 1))
		eQcrRBFtEu := insjC2Va + lNAwkbJQVBgU
		if 0 <= eQcrRBFtEu && eQcrRBFtEu <= 1<<33-1 {
			pv6eev.iB2EPt9 = pv6eev.iB2EPt9&nsecMask |  /*line :1*/ uint64(eQcrRBFtEu)<<nsecShift | hasMonotonic
			return
		}

		 /*line :1*/ pv6eev.xtoF9MHpmAIj()
	}

	s6sR1KzaPT := pv6eev.f0xL4nVadlv + lNAwkbJQVBgU
	if (s6sR1KzaPT > pv6eev.f0xL4nVadlv) == (lNAwkbJQVBgU > 0) {
		pv6eev.f0xL4nVadlv = s6sR1KzaPT
	} else if lNAwkbJQVBgU > 0 {
		pv6eev.f0xL4nVadlv = 1<<63 - 1
	} else {
		pv6eev.f0xL4nVadlv = -(1<<63 - 1)
	}
}

func (pv6eev *G4KDkI) vt5HqRw(sAqqFTcly5S *Hh4U1oidou) {
	if sAqqFTcly5S == &cYYjtVboaU {
		sAqqFTcly5S = nil
	}
	 /*line :1*/ pv6eev.xtoF9MHpmAIj()
	pv6eev.pPnZgmaso = sAqqFTcly5S
}

func (pv6eev *G4KDkI) xtoF9MHpmAIj() {
	if pv6eev.iB2EPt9&hasMonotonic != 0 {
		pv6eev.f0xL4nVadlv =  /*line :1*/ pv6eev.insjC2Va()
		pv6eev.iB2EPt9 &= nsecMask
	}
}

func (pv6eev *G4KDkI) djYmTkdInYt(fYy09z7m int64) {
	if pv6eev.iB2EPt9&hasMonotonic == 0 {
		insjC2Va := pv6eev.f0xL4nVadlv
		if insjC2Va < minWall || maxWall < insjC2Va {
			return
		}
		pv6eev.iB2EPt9 |= hasMonotonic |  /*line :1*/ uint64(insjC2Va-minWall)<<nsecShift
	}
	pv6eev.f0xL4nVadlv = fYy09z7m
}

func (pv6eev *G4KDkI) alQaKMu6() int64 {
	if pv6eev.iB2EPt9&hasMonotonic == 0 {
		return 0
	}
	return pv6eev.f0xL4nVadlv
}

func (pv6eev G4KDkI) After(b6XtDzh G4KDkI) bool {
	if pv6eev.iB2EPt9&b6XtDzh.iB2EPt9&hasMonotonic != 0 {
		return pv6eev.f0xL4nVadlv > b6XtDzh.f0xL4nVadlv
	}
	gUHvF4Tfi :=  /*line :1*/ pv6eev.insjC2Va()
	fglri91Q :=  /*line :1*/ b6XtDzh.insjC2Va()
	return gUHvF4Tfi > fglri91Q || gUHvF4Tfi == fglri91Q &&  /*line :1*/ pv6eev.aCiH2Z7qWLtw() >  /*line :1*/ b6XtDzh.aCiH2Z7qWLtw()
}

func (pv6eev G4KDkI) Before(b6XtDzh G4KDkI) bool {
	if pv6eev.iB2EPt9&b6XtDzh.iB2EPt9&hasMonotonic != 0 {
		return pv6eev.f0xL4nVadlv < b6XtDzh.f0xL4nVadlv
	}
	gUHvF4Tfi :=  /*line :1*/ pv6eev.insjC2Va()
	fglri91Q :=  /*line :1*/ b6XtDzh.insjC2Va()
	return gUHvF4Tfi < fglri91Q || gUHvF4Tfi == fglri91Q &&  /*line :1*/ pv6eev.aCiH2Z7qWLtw() <  /*line :1*/ b6XtDzh.aCiH2Z7qWLtw()
}

func (pv6eev G4KDkI) Compare(b6XtDzh G4KDkI) int {
	var o1YbPLoANYO, cleMN34w50 int64
	if pv6eev.iB2EPt9&b6XtDzh.iB2EPt9&hasMonotonic != 0 {
		o1YbPLoANYO, cleMN34w50 = pv6eev.f0xL4nVadlv, b6XtDzh.f0xL4nVadlv
	} else {
		o1YbPLoANYO, cleMN34w50 =  /*line :1*/ pv6eev.insjC2Va(),  /*line :1*/ b6XtDzh.insjC2Va()
		if o1YbPLoANYO == cleMN34w50 {
			o1YbPLoANYO, cleMN34w50 =  /*line :1*/ int64( /*line :1*/ pv6eev.aCiH2Z7qWLtw()),  /*line :1*/ int64( /*line :1*/ b6XtDzh.aCiH2Z7qWLtw())
		}
	}
	switch {
	case o1YbPLoANYO < cleMN34w50:
		return -1
	case o1YbPLoANYO > cleMN34w50:
		return +1
	}
	return 0
}

func (pv6eev G4KDkI) Equal(b6XtDzh G4KDkI) bool {
	if pv6eev.iB2EPt9&b6XtDzh.iB2EPt9&hasMonotonic != 0 {
		return pv6eev.f0xL4nVadlv == b6XtDzh.f0xL4nVadlv
	}
	return  /*line :1*/ pv6eev.insjC2Va() ==  /*line :1*/ b6XtDzh.insjC2Va() &&  /*line :1*/ pv6eev.aCiH2Z7qWLtw() ==  /*line :1*/ b6XtDzh.aCiH2Z7qWLtw()
}

type ZyPFXNxLcwpT int

const (
	January	ZyPFXNxLcwpT	= 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (fYy09z7m ZyPFXNxLcwpT) String() string {
	if January <= fYy09z7m && fYy09z7m <= December {
		return sxLand[fYy09z7m-1]
	}
	qtO6L35R9b :=  /*line :1*/ make([]byte, 20)
	aeBZDJEINbb :=  /*line :1*/ mCCY5mf(qtO6L35R9b,  /*line :1*/ uint64(fYy09z7m))
	return "%!Month(" +  /*line :1*/ string(qtO6L35R9b[aeBZDJEINbb:]) + ")"
}

type KZEv3wea int

const (
	Sunday	KZEv3wea	= iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (lNAwkbJQVBgU KZEv3wea) String() string {
	if Sunday <= lNAwkbJQVBgU && lNAwkbJQVBgU <= Saturday {
		return rXk2xR7P[lNAwkbJQVBgU]
	}
	qtO6L35R9b :=  /*line :1*/ make([]byte, 20)
	aeBZDJEINbb :=  /*line :1*/ mCCY5mf(qtO6L35R9b,  /*line :1*/ uint64(lNAwkbJQVBgU))
	return "%!Weekday(" +  /*line :1*/ string(qtO6L35R9b[aeBZDJEINbb:]) + ")"
}

const (
	absoluteZeroYear		= -292277022399

	internalYear		= 1

	absoluteToInternal	int64	= (absoluteZeroYear - internalYear) * 365.2425 * secondsPerDay
	internalToAbsolute		= -absoluteToInternal

	unixToInternal	int64	= (1969*365 + 1969/4 - 1969/100 + 1969/400) * secondsPerDay
	internalToUnix	int64	= -unixToInternal

	wallToInternal	int64	= (1884*365 + 1884/4 - 1884/100 + 1884/400) * secondsPerDay
)

func (pv6eev G4KDkI) IsZero() bool {
	return  /*line :1*/ pv6eev.insjC2Va() == 0 &&  /*line :1*/ pv6eev.aCiH2Z7qWLtw() == 0
}

func (pv6eev G4KDkI) nsu4XSEmD7r() uint64 {
	ipt7urdFIkH := pv6eev.pPnZgmaso

	if ipt7urdFIkH == nil || ipt7urdFIkH == &s45BJX {
		ipt7urdFIkH =  /*line :1*/ ipt7urdFIkH.w9w0xFJlcOz()
	}
	insjC2Va :=  /*line :1*/ pv6eev.huzSe5e()
	if ipt7urdFIkH != &cYYjtVboaU {
		if ipt7urdFIkH.j9a8yEma != nil && ipt7urdFIkH.tIAdDDIq <= insjC2Va && insjC2Va < ipt7urdFIkH.cmASbbDbiav {
			insjC2Va +=  /*line :1*/ int64(ipt7urdFIkH.j9a8yEma.a8VzUELX)
		} else {
			_, l7D_bh2w, _, _, _ :=  /*line :1*/ ipt7urdFIkH.oPdG5Tk5(insjC2Va)
			insjC2Va +=  /*line :1*/ int64(l7D_bh2w)
		}
	}
	return  /*line :1*/ uint64(insjC2Va + (unixToInternal + internalToAbsolute))
}

func (pv6eev G4KDkI) bNr1Yr() (w12CVdh9 string, l7D_bh2w int, nsu4XSEmD7r uint64) {
	ipt7urdFIkH := pv6eev.pPnZgmaso
	if ipt7urdFIkH == nil || ipt7urdFIkH == &s45BJX {
		ipt7urdFIkH =  /*line :1*/ ipt7urdFIkH.w9w0xFJlcOz()
	}

	insjC2Va :=  /*line :1*/ pv6eev.huzSe5e()
	if ipt7urdFIkH != &cYYjtVboaU {
		if ipt7urdFIkH.j9a8yEma != nil && ipt7urdFIkH.tIAdDDIq <= insjC2Va && insjC2Va < ipt7urdFIkH.cmASbbDbiav {
			w12CVdh9 = ipt7urdFIkH.j9a8yEma.sn0AXnjD
			l7D_bh2w = ipt7urdFIkH.j9a8yEma.a8VzUELX
		} else {
			w12CVdh9, l7D_bh2w, _, _, _ =  /*line :1*/ ipt7urdFIkH.oPdG5Tk5(insjC2Va)
		}
		insjC2Va +=  /*line :1*/ int64(l7D_bh2w)
	} else {
		w12CVdh9 = "UTC"
	}
	nsu4XSEmD7r =  /*line :1*/ uint64(insjC2Va + (unixToInternal + internalToAbsolute))
	return
}

func (pv6eev G4KDkI) Date() (mO2sH8hUHnW int, cmQVCZ2yilS ZyPFXNxLcwpT, gv5Z468R int) {
	mO2sH8hUHnW, cmQVCZ2yilS, gv5Z468R, _ =  /*line :1*/ pv6eev.euRPDa(true)
	return
}

func (pv6eev G4KDkI) Year() int {
	mO2sH8hUHnW, _, _, _ :=  /*line :1*/ pv6eev.euRPDa(false)
	return mO2sH8hUHnW
}

func (pv6eev G4KDkI) Month() ZyPFXNxLcwpT {
	_, cmQVCZ2yilS, _, _ :=  /*line :1*/ pv6eev.euRPDa(true)
	return cmQVCZ2yilS
}

func (pv6eev G4KDkI) Day() int {
	_, _, gv5Z468R, _ :=  /*line :1*/ pv6eev.euRPDa(true)
	return gv5Z468R
}

func (pv6eev G4KDkI) Weekday() KZEv3wea {
	return  /*line :1*/ baS3QC6JH( /*line :1*/ pv6eev.nsu4XSEmD7r())
}

func baS3QC6JH(nsu4XSEmD7r uint64) KZEv3wea {

	insjC2Va := (nsu4XSEmD7r +  /*line :1*/ uint64(Monday)*secondsPerDay) % secondsPerWeek
	return  /*line :1*/ KZEv3wea( /*line :1*/ int(insjC2Va) / secondsPerDay)
}

func (pv6eev G4KDkI) ISOWeek() (mO2sH8hUHnW, cDtaLw1Mj int) {

	nsu4XSEmD7r :=  /*line :1*/ pv6eev.nsu4XSEmD7r()
	lNAwkbJQVBgU := Thursday -  /*line :1*/ baS3QC6JH(nsu4XSEmD7r)

	if lNAwkbJQVBgU == 4 {
		lNAwkbJQVBgU = -3
	}

	nsu4XSEmD7r +=  /*line :1*/ uint64(lNAwkbJQVBgU) * secondsPerDay
	mO2sH8hUHnW, _, _, rAwk8hU1NKXu :=  /*line :1*/ fdO09EiS(nsu4XSEmD7r, false)
	return mO2sH8hUHnW, rAwk8hU1NKXu/7 + 1
}

func (pv6eev G4KDkI) Clock() (xmMDTE, qJ3jbB5O, insjC2Va int) {
	return  /*line :1*/ mbFqMakX( /*line :1*/ pv6eev.nsu4XSEmD7r())
}

func mbFqMakX(nsu4XSEmD7r uint64) (xmMDTE, qJ3jbB5O, insjC2Va int) {
	insjC2Va =  /*line :1*/ int(nsu4XSEmD7r % secondsPerDay)
	xmMDTE = insjC2Va / secondsPerHour
	insjC2Va -= xmMDTE * secondsPerHour
	qJ3jbB5O = insjC2Va / secondsPerMinute
	insjC2Va -= qJ3jbB5O * secondsPerMinute
	return
}

func (pv6eev G4KDkI) Hour() int {
	return  /*line :1*/ int( /*line :1*/ pv6eev.nsu4XSEmD7r()%secondsPerDay) / secondsPerHour
}

func (pv6eev G4KDkI) Minute() int {
	return  /*line :1*/ int( /*line :1*/ pv6eev.nsu4XSEmD7r()%secondsPerHour) / secondsPerMinute
}

func (pv6eev G4KDkI) Second() int {
	return  /*line :1*/ int( /*line :1*/ pv6eev.nsu4XSEmD7r() % secondsPerMinute)
}

func (pv6eev G4KDkI) Nanosecond() int {
	return  /*line :1*/ int( /*line :1*/ pv6eev.aCiH2Z7qWLtw())
}

func (pv6eev G4KDkI) YearDay() int {
	_, _, _, rAwk8hU1NKXu :=  /*line :1*/ pv6eev.euRPDa(false)
	return rAwk8hU1NKXu + 1
}

type GKMwTGxQa0S int64

const (
	minDuration	GKMwTGxQa0S	= -1 << 63
	maxDuration	GKMwTGxQa0S	= 1<<63 - 1
)

const (
	Nanosecond	GKMwTGxQa0S	= 1
	Microsecond		= 1000 * Nanosecond
	Millisecond		= 1000 * Microsecond
	Second		= 1000 * Millisecond
	Minute		= 60 * Second
	Hour		= 60 * Minute
)

func (lNAwkbJQVBgU GKMwTGxQa0S) String() string {

	var qtO6L35R9b [32]byte
	n2g7_eLcTc :=  /*line :1*/ len(qtO6L35R9b)

	b6XtDzh :=  /*line :1*/ uint64(lNAwkbJQVBgU)
	yY3pgXWox3j := lNAwkbJQVBgU < 0
	if yY3pgXWox3j {
		b6XtDzh = -b6XtDzh
	}

	if b6XtDzh <  /*line :1*/ uint64(Second) {

		var bNtedthoEv int
		n2g7_eLcTc--
		qtO6L35R9b[n2g7_eLcTc] = 's'
		n2g7_eLcTc--
		switch {
		case b6XtDzh == 0:
			return "0s"
		case b6XtDzh <  /*line :1*/ uint64(Microsecond):

			bNtedthoEv = 0
			qtO6L35R9b[n2g7_eLcTc] = 'n'
		case b6XtDzh <  /*line :1*/ uint64(Millisecond):

			bNtedthoEv = 3

			n2g7_eLcTc--
			 /*line :1*/ copy(qtO6L35R9b[n2g7_eLcTc:], "µ")
		default:

			bNtedthoEv = 6
			qtO6L35R9b[n2g7_eLcTc] = 'm'
		}
		n2g7_eLcTc, b6XtDzh =  /*line :1*/ hcmUEs(qtO6L35R9b[:n2g7_eLcTc], b6XtDzh, bNtedthoEv)
		n2g7_eLcTc =  /*line :1*/ mCCY5mf(qtO6L35R9b[:n2g7_eLcTc], b6XtDzh)
	} else {
		n2g7_eLcTc--
		qtO6L35R9b[n2g7_eLcTc] = 's'

		n2g7_eLcTc, b6XtDzh =  /*line :1*/ hcmUEs(qtO6L35R9b[:n2g7_eLcTc], b6XtDzh, 9)

		n2g7_eLcTc =  /*line :1*/ mCCY5mf(qtO6L35R9b[:n2g7_eLcTc], b6XtDzh%60)
		b6XtDzh /= 60

		if b6XtDzh > 0 {
			n2g7_eLcTc--
			qtO6L35R9b[n2g7_eLcTc] = 'm'
			n2g7_eLcTc =  /*line :1*/ mCCY5mf(qtO6L35R9b[:n2g7_eLcTc], b6XtDzh%60)
			b6XtDzh /= 60

			if b6XtDzh > 0 {
				n2g7_eLcTc--
				qtO6L35R9b[n2g7_eLcTc] = 'h'
				n2g7_eLcTc =  /*line :1*/ mCCY5mf(qtO6L35R9b[:n2g7_eLcTc], b6XtDzh)
			}
		}
	}

	if yY3pgXWox3j {
		n2g7_eLcTc--
		qtO6L35R9b[n2g7_eLcTc] = '-'
	}

	return  /*line :1*/ string(qtO6L35R9b[n2g7_eLcTc:])
}

func hcmUEs(qtO6L35R9b []byte, ha9RaMigT uint64, bNtedthoEv int) (gpv6QOw9 int, h3HCec uint64) {

	n2g7_eLcTc :=  /*line :1*/ len(qtO6L35R9b)
	v2_lbvo := false
	for eUo6AxKM7 := 0; eUo6AxKM7 < bNtedthoEv; eUo6AxKM7++ {
		uQXIizb := ha9RaMigT % 10
		v2_lbvo = v2_lbvo || uQXIizb != 0
		if v2_lbvo {
			n2g7_eLcTc--
			qtO6L35R9b[n2g7_eLcTc] =  /*line :1*/ byte(uQXIizb) + '0'
		}
		ha9RaMigT /= 10
	}
	if v2_lbvo {
		n2g7_eLcTc--
		qtO6L35R9b[n2g7_eLcTc] = '.'
	}
	return n2g7_eLcTc, ha9RaMigT
}

func mCCY5mf(qtO6L35R9b []byte, ha9RaMigT uint64) int {
	n2g7_eLcTc :=  /*line :1*/ len(qtO6L35R9b)
	if ha9RaMigT == 0 {
		n2g7_eLcTc--
		qtO6L35R9b[n2g7_eLcTc] = '0'
	} else {
		for ha9RaMigT > 0 {
			n2g7_eLcTc--
			qtO6L35R9b[n2g7_eLcTc] =  /*line :1*/ byte(ha9RaMigT%10) + '0'
			ha9RaMigT /= 10
		}
	}
	return n2g7_eLcTc
}

func (lNAwkbJQVBgU GKMwTGxQa0S) Nanoseconds() int64	{ return  /*line :1*/ int64(lNAwkbJQVBgU) }

func (lNAwkbJQVBgU GKMwTGxQa0S) Microseconds() int64	{ return  /*line :1*/ int64(lNAwkbJQVBgU) / 1e3 }

func (lNAwkbJQVBgU GKMwTGxQa0S) Milliseconds() int64	{ return  /*line :1*/ int64(lNAwkbJQVBgU) / 1e6 }

func (lNAwkbJQVBgU GKMwTGxQa0S) Seconds() float64 {
	insjC2Va := lNAwkbJQVBgU / Second
	aCiH2Z7qWLtw := lNAwkbJQVBgU % Second
	return  /*line :1*/ float64(insjC2Va) +  /*line :1*/ float64(aCiH2Z7qWLtw)/1e9
}

func (lNAwkbJQVBgU GKMwTGxQa0S) Minutes() float64 {
	qJ3jbB5O := lNAwkbJQVBgU / Minute
	aCiH2Z7qWLtw := lNAwkbJQVBgU % Minute
	return  /*line :1*/ float64(qJ3jbB5O) +  /*line :1*/ float64(aCiH2Z7qWLtw)/(60*1e9)
}

func (lNAwkbJQVBgU GKMwTGxQa0S) Hours() float64 {
	xmMDTE := lNAwkbJQVBgU / Hour
	aCiH2Z7qWLtw := lNAwkbJQVBgU % Hour
	return  /*line :1*/ float64(xmMDTE) +  /*line :1*/ float64(aCiH2Z7qWLtw)/(60*60*1e9)
}

func (lNAwkbJQVBgU GKMwTGxQa0S) Truncate(fYy09z7m GKMwTGxQa0S) GKMwTGxQa0S {
	if fYy09z7m <= 0 {
		return lNAwkbJQVBgU
	}
	return lNAwkbJQVBgU - lNAwkbJQVBgU%fYy09z7m
}

func vsmVMYaci(i_2w697Gglw, bRAhk8mArR1z GKMwTGxQa0S) bool {
	return  /*line :1*/ uint64(i_2w697Gglw)+ /*line :1*/ uint64(i_2w697Gglw) <  /*line :1*/ uint64(bRAhk8mArR1z)
}

func (lNAwkbJQVBgU GKMwTGxQa0S) Round(fYy09z7m GKMwTGxQa0S) GKMwTGxQa0S {
	if fYy09z7m <= 0 {
		return lNAwkbJQVBgU
	}
	aLhBMRzpWd := lNAwkbJQVBgU % fYy09z7m
	if lNAwkbJQVBgU < 0 {
		aLhBMRzpWd = -aLhBMRzpWd
		if  /*line :1*/ vsmVMYaci(aLhBMRzpWd, fYy09z7m) {
			return lNAwkbJQVBgU + aLhBMRzpWd
		}
		if kIMGbZ := lNAwkbJQVBgU - fYy09z7m + aLhBMRzpWd; kIMGbZ < lNAwkbJQVBgU {
			return kIMGbZ
		}
		return minDuration
	}
	if  /*line :1*/ vsmVMYaci(aLhBMRzpWd, fYy09z7m) {
		return lNAwkbJQVBgU - aLhBMRzpWd
	}
	if kIMGbZ := lNAwkbJQVBgU + fYy09z7m - aLhBMRzpWd; kIMGbZ > lNAwkbJQVBgU {
		return kIMGbZ
	}
	return maxDuration
}

func (lNAwkbJQVBgU GKMwTGxQa0S) Abs() GKMwTGxQa0S {
	switch {
	case lNAwkbJQVBgU >= 0:
		return lNAwkbJQVBgU
	case lNAwkbJQVBgU == minDuration:
		return maxDuration
	default:
		return -lNAwkbJQVBgU
	}
}

func (pv6eev G4KDkI) Add(lNAwkbJQVBgU GKMwTGxQa0S) G4KDkI {
	eQcrRBFtEu :=  /*line :1*/ int64(lNAwkbJQVBgU / 1e9)
	aCiH2Z7qWLtw :=  /*line :1*/ pv6eev.aCiH2Z7qWLtw() +  /*line :1*/ int32(lNAwkbJQVBgU%1e9)
	if aCiH2Z7qWLtw >= 1e9 {
		eQcrRBFtEu++
		aCiH2Z7qWLtw -= 1e9
	} else if aCiH2Z7qWLtw < 0 {
		eQcrRBFtEu--
		aCiH2Z7qWLtw += 1e9
	}
	pv6eev.iB2EPt9 = pv6eev.iB2EPt9&^nsecMask |  /*line :1*/ uint64(aCiH2Z7qWLtw)
	 /*line :1*/ pv6eev.hOFVKuqUr91a(eQcrRBFtEu)
	if pv6eev.iB2EPt9&hasMonotonic != 0 {
		xy8IGK := pv6eev.f0xL4nVadlv +  /*line :1*/ int64(lNAwkbJQVBgU)
		if lNAwkbJQVBgU < 0 && xy8IGK > pv6eev.f0xL4nVadlv || lNAwkbJQVBgU > 0 && xy8IGK < pv6eev.f0xL4nVadlv {

			 /*line :1*/ pv6eev.xtoF9MHpmAIj()
		} else {
			pv6eev.f0xL4nVadlv = xy8IGK
		}
	}
	return pv6eev
}

func (pv6eev G4KDkI) Sub(b6XtDzh G4KDkI) GKMwTGxQa0S {
	if pv6eev.iB2EPt9&b6XtDzh.iB2EPt9&hasMonotonic != 0 {
		xy8IGK := pv6eev.f0xL4nVadlv
		aaWaI2 := b6XtDzh.f0xL4nVadlv
		lNAwkbJQVBgU :=  /*line :1*/ GKMwTGxQa0S(xy8IGK - aaWaI2)
		if lNAwkbJQVBgU < 0 && xy8IGK > aaWaI2 {
			return maxDuration
		}
		if lNAwkbJQVBgU > 0 && xy8IGK < aaWaI2 {
			return minDuration
		}
		return lNAwkbJQVBgU
	}
	lNAwkbJQVBgU :=  /*line :1*/ GKMwTGxQa0S( /*line :1*/ pv6eev.insjC2Va()- /*line :1*/ b6XtDzh.insjC2Va())*Second +  /*line :1*/ GKMwTGxQa0S( /*line :1*/ pv6eev.aCiH2Z7qWLtw()- /*line :1*/ b6XtDzh.aCiH2Z7qWLtw())

	switch {
	case  /*line :1*/ b6XtDzh.Add(lNAwkbJQVBgU).Equal(pv6eev):
		return lNAwkbJQVBgU
	case  /*line :1*/ pv6eev.Before(b6XtDzh):
		return minDuration
	default:
		return maxDuration
	}
}

func Qs3VCKXXED9(pv6eev G4KDkI) GKMwTGxQa0S {
	var sNTZrc2lj G4KDkI
	if pv6eev.iB2EPt9&hasMonotonic != 0 {

		sNTZrc2lj = G4KDkI{hasMonotonic,  /*line :1*/ qQJTJ6m() - p_NmExP, nil}
	} else {
		sNTZrc2lj =  /*line :1*/ Pc_35oTY()
	}
	return  /*line :1*/ sNTZrc2lj.Sub(pv6eev)
}

func Qr_kn6Dzra(pv6eev G4KDkI) GKMwTGxQa0S {
	var sNTZrc2lj G4KDkI
	if pv6eev.iB2EPt9&hasMonotonic != 0 {

		sNTZrc2lj = G4KDkI{hasMonotonic,  /*line :1*/ qQJTJ6m() - p_NmExP, nil}
	} else {
		sNTZrc2lj =  /*line :1*/ Pc_35oTY()
	}
	return  /*line :1*/ pv6eev.Sub(sNTZrc2lj)
}

func (pv6eev G4KDkI) AddDate(r3iccke7XrF int, pZB04Y int, iiEminG int) G4KDkI {
	mO2sH8hUHnW, cmQVCZ2yilS, gv5Z468R :=  /*line :1*/ pv6eev.Date()
	xmMDTE, qJ3jbB5O, insjC2Va :=  /*line :1*/ pv6eev.Clock()
	return  /*line :1*/ FaDSoi2w(mO2sH8hUHnW+r3iccke7XrF, cmQVCZ2yilS+ /*line :1*/ ZyPFXNxLcwpT(pZB04Y), gv5Z468R+iiEminG, xmMDTE, qJ3jbB5O, insjC2Va,  /*line :1*/ int( /*line :1*/ pv6eev.aCiH2Z7qWLtw()),  /*line :1*/ pv6eev.Location())
}

const (
	secondsPerMinute	= 60
	secondsPerHour	= 60 * secondsPerMinute
	secondsPerDay	= 24 * secondsPerHour
	secondsPerWeek	= 7 * secondsPerDay
	daysPer400Years	= 365*400 + 97
	daysPer100Years	= 365*100 + 24
	daysPer4Years	= 365*4 + 1
)

func (pv6eev G4KDkI) euRPDa(oNX1AZ6 bool) (mO2sH8hUHnW int, cmQVCZ2yilS ZyPFXNxLcwpT, gv5Z468R int, rAwk8hU1NKXu int) {
	return  /*line :1*/ fdO09EiS( /*line :1*/ pv6eev.nsu4XSEmD7r(), oNX1AZ6)
}

func fdO09EiS(nsu4XSEmD7r uint64, oNX1AZ6 bool) (mO2sH8hUHnW int, cmQVCZ2yilS ZyPFXNxLcwpT, gv5Z468R int, rAwk8hU1NKXu int) {

	lNAwkbJQVBgU := nsu4XSEmD7r / secondsPerDay

	aeBZDJEINbb := lNAwkbJQVBgU / daysPer400Years
	bRAhk8mArR1z := 400 * aeBZDJEINbb
	lNAwkbJQVBgU -= daysPer400Years * aeBZDJEINbb

	aeBZDJEINbb = lNAwkbJQVBgU / daysPer100Years
	aeBZDJEINbb -= aeBZDJEINbb >> 2
	bRAhk8mArR1z += 100 * aeBZDJEINbb
	lNAwkbJQVBgU -= daysPer100Years * aeBZDJEINbb

	aeBZDJEINbb = lNAwkbJQVBgU / daysPer4Years
	bRAhk8mArR1z += 4 * aeBZDJEINbb
	lNAwkbJQVBgU -= daysPer4Years * aeBZDJEINbb

	aeBZDJEINbb = lNAwkbJQVBgU / 365
	aeBZDJEINbb -= aeBZDJEINbb >> 2
	bRAhk8mArR1z += aeBZDJEINbb
	lNAwkbJQVBgU -= 365 * aeBZDJEINbb

	mO2sH8hUHnW =  /*line :1*/ int( /*line :1*/ int64(bRAhk8mArR1z) + absoluteZeroYear)
	rAwk8hU1NKXu =  /*line :1*/ int(lNAwkbJQVBgU)

	if !oNX1AZ6 {
		return
	}

	gv5Z468R = rAwk8hU1NKXu
	if  /*line :1*/ dJS2LEP6(mO2sH8hUHnW) {

		switch {
		case gv5Z468R > 31+29-1:

			gv5Z468R--
		case gv5Z468R == 31+29-1:

			cmQVCZ2yilS = February
			gv5Z468R = 29
			return
		}
	}

	cmQVCZ2yilS =  /*line :1*/ ZyPFXNxLcwpT(gv5Z468R / 31)
	wo7u8LY9W :=  /*line :1*/ int(h7KunS2pF[cmQVCZ2yilS+1])
	var xiaZeNCYUl4 int
	if gv5Z468R >= wo7u8LY9W {
		cmQVCZ2yilS++
		xiaZeNCYUl4 = wo7u8LY9W
	} else {
		xiaZeNCYUl4 =  /*line :1*/ int(h7KunS2pF[cmQVCZ2yilS])
	}

	cmQVCZ2yilS++
	gv5Z468R = gv5Z468R - xiaZeNCYUl4 + 1
	return
}

var h7KunS2pF = [...]int32{
	0,
	31,
	31 + 28,
	31 + 28 + 31,
	31 + 28 + 31 + 30,
	31 + 28 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
}

func rs4bvng4W6y(fYy09z7m ZyPFXNxLcwpT, mO2sH8hUHnW int) int {
	if fYy09z7m == February &&  /*line :1*/ dJS2LEP6(mO2sH8hUHnW) {
		return 29
	}
	return  /*line :1*/ int(h7KunS2pF[fYy09z7m] - h7KunS2pF[fYy09z7m-1])
}

func wESvrx2(mO2sH8hUHnW int) uint64 {
	bRAhk8mArR1z :=  /*line :1*/ uint64( /*line :1*/ int64(mO2sH8hUHnW) - absoluteZeroYear)

	aeBZDJEINbb := bRAhk8mArR1z / 400
	bRAhk8mArR1z -= 400 * aeBZDJEINbb
	lNAwkbJQVBgU := daysPer400Years * aeBZDJEINbb

	aeBZDJEINbb = bRAhk8mArR1z / 100
	bRAhk8mArR1z -= 100 * aeBZDJEINbb
	lNAwkbJQVBgU += daysPer100Years * aeBZDJEINbb

	aeBZDJEINbb = bRAhk8mArR1z / 4
	bRAhk8mArR1z -= 4 * aeBZDJEINbb
	lNAwkbJQVBgU += daysPer4Years * aeBZDJEINbb

	aeBZDJEINbb = bRAhk8mArR1z
	lNAwkbJQVBgU += 365 * aeBZDJEINbb

	return lNAwkbJQVBgU
}

func sNTZrc2lj() (insjC2Va int64, aCiH2Z7qWLtw int32, alQaKMu6 int64)

//go:linkname qQJTJ6m runtime.nanotime
func qQJTJ6m() int64

var p_NmExP int64 =  /*line :1*/ qQJTJ6m() - 1

func Pc_35oTY() G4KDkI {
	insjC2Va, aCiH2Z7qWLtw, alQaKMu6 :=  /*line :1*/ sNTZrc2lj()
	alQaKMu6 -= p_NmExP
	insjC2Va += unixToInternal - minWall
	if  /*line :1*/ uint64(insjC2Va)>>33 != 0 {

		return G4KDkI{ /*line :1*/ uint64(aCiH2Z7qWLtw), insjC2Va + minWall, CH9afbXxKDfu}
	}
	return G4KDkI{hasMonotonic |  /*line :1*/ uint64(insjC2Va)<<nsecShift |  /*line :1*/ uint64(aCiH2Z7qWLtw), alQaKMu6, CH9afbXxKDfu}
}

func lrrBwnb4K(insjC2Va int64, aCiH2Z7qWLtw int32) G4KDkI {
	return G4KDkI{ /*line :1*/ uint64(aCiH2Z7qWLtw), insjC2Va + unixToInternal, CH9afbXxKDfu}
}

func (pv6eev G4KDkI) UTC() G4KDkI {
	 /*line :1*/ pv6eev.vt5HqRw(&cYYjtVboaU)
	return pv6eev
}

func (pv6eev G4KDkI) Local() G4KDkI {
	 /*line :1*/ pv6eev.vt5HqRw(CH9afbXxKDfu)
	return pv6eev
}

func (pv6eev G4KDkI) In(sAqqFTcly5S *Hh4U1oidou) G4KDkI {
	if sAqqFTcly5S == nil {
		 /*line :1*/ panic("time: missing Location in call to Time.In")
	}
	 /*line :1*/ pv6eev.vt5HqRw(sAqqFTcly5S)
	return pv6eev
}

func (pv6eev G4KDkI) Location() *Hh4U1oidou {
	ipt7urdFIkH := pv6eev.pPnZgmaso
	if ipt7urdFIkH == nil {
		ipt7urdFIkH = PD1NIUjTJ
	}
	return ipt7urdFIkH
}

func (pv6eev G4KDkI) Zone() (w12CVdh9 string, l7D_bh2w int) {
	w12CVdh9, l7D_bh2w, _, _, _ =  /*line :1*/ pv6eev.pPnZgmaso.oPdG5Tk5( /*line :1*/ pv6eev.huzSe5e())
	return
}

func (pv6eev G4KDkI) ZoneBounds() (cAZmDfINA, wo7u8LY9W G4KDkI) {
	_, _, iqaLrA, ftQIVp8, _ :=  /*line :1*/ pv6eev.pPnZgmaso.oPdG5Tk5( /*line :1*/ pv6eev.huzSe5e())
	if iqaLrA != alpha {
		cAZmDfINA =  /*line :1*/ lrrBwnb4K(iqaLrA, 0)
		 /*line :1*/ cAZmDfINA.vt5HqRw(pv6eev.pPnZgmaso)
	}
	if ftQIVp8 != omega {
		wo7u8LY9W =  /*line :1*/ lrrBwnb4K(ftQIVp8, 0)
		 /*line :1*/ wo7u8LY9W.vt5HqRw(pv6eev.pPnZgmaso)
	}
	return
}

func (pv6eev G4KDkI) Unix() int64 {
	return  /*line :1*/ pv6eev.huzSe5e()
}

func (pv6eev G4KDkI) UnixMilli() int64 {
	return  /*line :1*/ pv6eev.huzSe5e()*1e3 +  /*line :1*/ int64( /*line :1*/ pv6eev.aCiH2Z7qWLtw())/1e6
}

func (pv6eev G4KDkI) UnixMicro() int64 {
	return  /*line :1*/ pv6eev.huzSe5e()*1e6 +  /*line :1*/ int64( /*line :1*/ pv6eev.aCiH2Z7qWLtw())/1e3
}

func (pv6eev G4KDkI) UnixNano() int64 {
	return ( /*line :1*/ pv6eev.huzSe5e())*1e9 +  /*line :1*/ int64( /*line :1*/ pv6eev.aCiH2Z7qWLtw())
}

const (
	timeBinaryVersionV1	byte	= iota + 1
	timeBinaryVersionV2
)

func (pv6eev G4KDkI) MarshalBinary() ([]byte, error) {
	var to4Txdsxi1Xq int16
	var h6gykUnyWi int8
	e02h6We := timeBinaryVersionV1

	if  /*line :1*/ pv6eev.Location() == PD1NIUjTJ {
		to4Txdsxi1Xq = -1
	} else {
		_, l7D_bh2w :=  /*line :1*/ pv6eev.Zone()
		if l7D_bh2w%60 != 0 {
			e02h6We = timeBinaryVersionV2
			h6gykUnyWi =  /*line :1*/ int8(l7D_bh2w % 60)
		}

		l7D_bh2w /= 60
		if l7D_bh2w < -32768 || l7D_bh2w == -1 || l7D_bh2w > 32767 {
			return nil,  /*line :1*/ errors.Su6g6hRoi9X("Time.MarshalBinary: unexpected zone offset")
		}
		to4Txdsxi1Xq =  /*line :1*/ int16(l7D_bh2w)
	}

	insjC2Va :=  /*line :1*/ pv6eev.insjC2Va()
	aCiH2Z7qWLtw :=  /*line :1*/ pv6eev.aCiH2Z7qWLtw()
	ibq_DlPk := []byte{
		e02h6We,
		 /*line :1*/ byte(insjC2Va >> 56),
		 /*line :1*/ byte(insjC2Va >> 48),
		 /*line :1*/ byte(insjC2Va >> 40),
		 /*line :1*/ byte(insjC2Va >> 32),
		 /*line :1*/ byte(insjC2Va >> 24),
		 /*line :1*/ byte(insjC2Va >> 16),
		 /*line :1*/ byte(insjC2Va >> 8),
		 /*line :1*/ byte(insjC2Va),
		 /*line :1*/ byte(aCiH2Z7qWLtw >> 24),
		 /*line :1*/ byte(aCiH2Z7qWLtw >> 16),
		 /*line :1*/ byte(aCiH2Z7qWLtw >> 8),
		 /*line :1*/ byte(aCiH2Z7qWLtw),
		 /*line :1*/ byte(to4Txdsxi1Xq >> 8),
		 /*line :1*/ byte(to4Txdsxi1Xq),
	}
	if e02h6We == timeBinaryVersionV2 {
		ibq_DlPk =  /*line :1*/ append(ibq_DlPk,  /*line :1*/ byte(h6gykUnyWi))
	}

	return ibq_DlPk, nil
}

func (pv6eev *G4KDkI) UnmarshalBinary(a84q4w []byte) error {
	qtO6L35R9b := a84q4w
	if  /*line :1*/ len(qtO6L35R9b) == 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("Time.UnmarshalBinary: no data")
	}

	e02h6We := qtO6L35R9b[0]
	if e02h6We != timeBinaryVersionV1 && e02h6We != timeBinaryVersionV2 {
		return  /*line :1*/ errors.Su6g6hRoi9X("Time.UnmarshalBinary: unsupported version")
	}

	wdUARrh := 1 + 8 + 4 + 2
	if e02h6We == timeBinaryVersionV2 {
		wdUARrh++
	}
	if  /*line :1*/ len(qtO6L35R9b) != wdUARrh {
		return  /*line :1*/ errors.Su6g6hRoi9X("Time.UnmarshalBinary: invalid length")
	}

	qtO6L35R9b = qtO6L35R9b[1:]
	insjC2Va :=  /*line :1*/ int64(qtO6L35R9b[7]) |  /*line :1*/ int64(qtO6L35R9b[6])<<8 |  /*line :1*/ int64(qtO6L35R9b[5])<<16 |  /*line :1*/ int64(qtO6L35R9b[4])<<24 |
		 /*line :1*/ int64(qtO6L35R9b[3])<<32 |  /*line :1*/ int64(qtO6L35R9b[2])<<40 |  /*line :1*/ int64(qtO6L35R9b[1])<<48 |  /*line :1*/ int64(qtO6L35R9b[0])<<56

	qtO6L35R9b = qtO6L35R9b[8:]
	aCiH2Z7qWLtw :=  /*line :1*/ int32(qtO6L35R9b[3]) |  /*line :1*/ int32(qtO6L35R9b[2])<<8 |  /*line :1*/ int32(qtO6L35R9b[1])<<16 |  /*line :1*/ int32(qtO6L35R9b[0])<<24

	qtO6L35R9b = qtO6L35R9b[4:]
	l7D_bh2w :=  /*line :1*/ int( /*line :1*/ int16(qtO6L35R9b[1])| /*line :1*/ int16(qtO6L35R9b[0])<<8) * 60
	if e02h6We == timeBinaryVersionV2 {
		l7D_bh2w +=  /*line :1*/ int(qtO6L35R9b[2])
	}

	*pv6eev = G4KDkI{}
	pv6eev.iB2EPt9 =  /*line :1*/ uint64(aCiH2Z7qWLtw)
	pv6eev.f0xL4nVadlv = insjC2Va

	if l7D_bh2w == -1*60 {
		 /*line :1*/ pv6eev.vt5HqRw(&cYYjtVboaU)
	} else if _, a_cb8tm, _, _, _ :=  /*line :1*/ CH9afbXxKDfu.oPdG5Tk5( /*line :1*/ pv6eev.huzSe5e()); l7D_bh2w == a_cb8tm {
		 /*line :1*/ pv6eev.vt5HqRw(CH9afbXxKDfu)
	} else {
		 /*line :1*/ pv6eev.vt5HqRw( /*line :1*/ Z586lZG("", l7D_bh2w))
	}

	return nil
}

func (pv6eev G4KDkI) GobEncode() ([]byte, error) {
	return  /*line :1*/ pv6eev.MarshalBinary()
}

func (pv6eev *G4KDkI) GobDecode(a84q4w []byte) error {
	return  /*line :1*/ pv6eev.UnmarshalBinary(a84q4w)
}

func (pv6eev G4KDkI) MarshalJSON() ([]byte, error) {
	ev8znPS9W :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(RFC3339Nano)+ /*line :1*/ len(`""`))
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '"')
	ev8znPS9W, xuMLYrB :=  /*line :1*/ pv6eev.dOyyIxcFJ(ev8znPS9W)
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '"')
	if xuMLYrB != nil {
		return nil,  /*line :1*/ errors.Su6g6hRoi9X("Time.MarshalJSON: " +  /*line :1*/ xuMLYrB.Error())
	}
	return ev8znPS9W, nil
}

func (pv6eev *G4KDkI) UnmarshalJSON(a84q4w []byte) error {
	if  /*line :1*/ string(a84q4w) == "null" {
		return nil
	}

	if  /*line :1*/ len(a84q4w) < 2 || a84q4w[0] != '"' || a84q4w[ /*line :1*/ len(a84q4w)-1] != '"' {
		return  /*line :1*/ errors.Su6g6hRoi9X("Time.UnmarshalJSON: input is not a JSON string")
	}
	a84q4w = a84q4w[ /*line :1*/ len(`"`) :  /*line :1*/ len(a84q4w)- /*line :1*/ len(`"`)]
	var xuMLYrB error
	*pv6eev, xuMLYrB =  /*line :1*/ i7Ia_y(a84q4w)
	return xuMLYrB
}

func (pv6eev G4KDkI) MarshalText() ([]byte, error) {
	ev8znPS9W :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(RFC3339Nano))
	ev8znPS9W, xuMLYrB :=  /*line :1*/ pv6eev.dOyyIxcFJ(ev8znPS9W)
	if xuMLYrB != nil {
		return nil,  /*line :1*/ errors.Su6g6hRoi9X("Time.MarshalText: " +  /*line :1*/ xuMLYrB.Error())
	}
	return ev8znPS9W, nil
}

func (pv6eev *G4KDkI) UnmarshalText(a84q4w []byte) error {
	var xuMLYrB error
	*pv6eev, xuMLYrB =  /*line :1*/ i7Ia_y(a84q4w)
	return xuMLYrB
}

func FRXtx9QnU(insjC2Va int64, aCiH2Z7qWLtw int64) G4KDkI {
	if aCiH2Z7qWLtw < 0 || aCiH2Z7qWLtw >= 1e9 {
		aeBZDJEINbb := aCiH2Z7qWLtw / 1e9
		insjC2Va += aeBZDJEINbb
		aCiH2Z7qWLtw -= aeBZDJEINbb * 1e9
		if aCiH2Z7qWLtw < 0 {
			aCiH2Z7qWLtw += 1e9
			insjC2Va--
		}
	}
	return  /*line :1*/ lrrBwnb4K(insjC2Va,  /*line :1*/ int32(aCiH2Z7qWLtw))
}

func NJqpDRjCd(nPL77gP int64) G4KDkI {
	return  /*line :1*/ FRXtx9QnU(nPL77gP/1e3, (nPL77gP%1e3)*1e6)
}

func KbUTZPE1y(aZ4iUcO2V int64) G4KDkI {
	return  /*line :1*/ FRXtx9QnU(aZ4iUcO2V/1e6, (aZ4iUcO2V%1e6)*1e3)
}

func (pv6eev G4KDkI) IsDST() bool {
	_, _, _, _, kVCaCmlEp1vW :=  /*line :1*/ pv6eev.pPnZgmaso.oPdG5Tk5( /*line :1*/ pv6eev.Unix())
	return kVCaCmlEp1vW
}

func dJS2LEP6(mO2sH8hUHnW int) bool {
	return mO2sH8hUHnW%4 == 0 && (mO2sH8hUHnW%100 != 0 || mO2sH8hUHnW%400 == 0)
}

func dgZEVROX(msqan9KLw, xW2tbzZ6pY, dQ5gxQ4 int) (iylsOzcq, ujt1U9FaW3 int) {
	if xW2tbzZ6pY < 0 {
		aeBZDJEINbb := (-xW2tbzZ6pY-1)/dQ5gxQ4 + 1
		msqan9KLw -= aeBZDJEINbb
		xW2tbzZ6pY += aeBZDJEINbb * dQ5gxQ4
	}
	if xW2tbzZ6pY >= dQ5gxQ4 {
		aeBZDJEINbb := xW2tbzZ6pY / dQ5gxQ4
		msqan9KLw += aeBZDJEINbb
		xW2tbzZ6pY -= aeBZDJEINbb * dQ5gxQ4
	}
	return msqan9KLw, xW2tbzZ6pY
}

func FaDSoi2w(mO2sH8hUHnW int, cmQVCZ2yilS ZyPFXNxLcwpT, gv5Z468R, xmMDTE, qJ3jbB5O, insjC2Va, aCiH2Z7qWLtw int, sAqqFTcly5S *Hh4U1oidou) G4KDkI {
	if sAqqFTcly5S == nil {
		 /*line :1*/ panic("time: missing Location in call to Date")
	}

	fYy09z7m :=  /*line :1*/ int(cmQVCZ2yilS) - 1
	mO2sH8hUHnW, fYy09z7m =  /*line :1*/ dgZEVROX(mO2sH8hUHnW, fYy09z7m, 12)
	cmQVCZ2yilS =  /*line :1*/ ZyPFXNxLcwpT(fYy09z7m) + 1

	insjC2Va, aCiH2Z7qWLtw =  /*line :1*/ dgZEVROX(insjC2Va, aCiH2Z7qWLtw, 1e9)
	qJ3jbB5O, insjC2Va =  /*line :1*/ dgZEVROX(qJ3jbB5O, insjC2Va, 60)
	xmMDTE, qJ3jbB5O =  /*line :1*/ dgZEVROX(xmMDTE, qJ3jbB5O, 60)
	gv5Z468R, xmMDTE =  /*line :1*/ dgZEVROX(gv5Z468R, xmMDTE, 24)

	lNAwkbJQVBgU :=  /*line :1*/ wESvrx2(mO2sH8hUHnW)

	lNAwkbJQVBgU +=  /*line :1*/ uint64(h7KunS2pF[cmQVCZ2yilS-1])
	if  /*line :1*/ dJS2LEP6(mO2sH8hUHnW) && cmQVCZ2yilS >= March {
		lNAwkbJQVBgU++
	}

	lNAwkbJQVBgU +=  /*line :1*/ uint64(gv5Z468R - 1)

	nsu4XSEmD7r := lNAwkbJQVBgU * secondsPerDay
	nsu4XSEmD7r +=  /*line :1*/ uint64(xmMDTE*secondsPerHour + qJ3jbB5O*secondsPerMinute + insjC2Va)

	xqgBUhRs :=  /*line :1*/ int64(nsu4XSEmD7r) + (absoluteToInternal + internalToUnix)

	_, l7D_bh2w, cAZmDfINA, wo7u8LY9W, _ :=  /*line :1*/ sAqqFTcly5S.oPdG5Tk5(xqgBUhRs)
	if l7D_bh2w != 0 {
		bjgI4q0MZ5B := xqgBUhRs -  /*line :1*/ int64(l7D_bh2w)

		if bjgI4q0MZ5B < cAZmDfINA || bjgI4q0MZ5B >= wo7u8LY9W {
			_, l7D_bh2w, _, _, _ =  /*line :1*/ sAqqFTcly5S.oPdG5Tk5(bjgI4q0MZ5B)
		}
		xqgBUhRs -=  /*line :1*/ int64(l7D_bh2w)
	}

	pv6eev :=  /*line :1*/ lrrBwnb4K(xqgBUhRs,  /*line :1*/ int32(aCiH2Z7qWLtw))
	 /*line :1*/ pv6eev.vt5HqRw(sAqqFTcly5S)
	return pv6eev
}

func (pv6eev G4KDkI) Truncate(lNAwkbJQVBgU GKMwTGxQa0S) G4KDkI {
	 /*line :1*/ pv6eev.xtoF9MHpmAIj()
	if lNAwkbJQVBgU <= 0 {
		return pv6eev
	}
	_, aLhBMRzpWd :=  /*line :1*/ htvyhYQ(pv6eev, lNAwkbJQVBgU)
	return  /*line :1*/ pv6eev.Add(-aLhBMRzpWd)
}

func (pv6eev G4KDkI) Round(lNAwkbJQVBgU GKMwTGxQa0S) G4KDkI {
	 /*line :1*/ pv6eev.xtoF9MHpmAIj()
	if lNAwkbJQVBgU <= 0 {
		return pv6eev
	}
	_, aLhBMRzpWd :=  /*line :1*/ htvyhYQ(pv6eev, lNAwkbJQVBgU)
	if  /*line :1*/ vsmVMYaci(aLhBMRzpWd, lNAwkbJQVBgU) {
		return  /*line :1*/ pv6eev.Add(-aLhBMRzpWd)
	}
	return  /*line :1*/ pv6eev.Add(lNAwkbJQVBgU - aLhBMRzpWd)
}

func htvyhYQ(pv6eev G4KDkI, lNAwkbJQVBgU GKMwTGxQa0S) (npEj_3Oa7V int, aLhBMRzpWd GKMwTGxQa0S) {
	yY3pgXWox3j := false
	aCiH2Z7qWLtw :=  /*line :1*/ pv6eev.aCiH2Z7qWLtw()
	insjC2Va :=  /*line :1*/ pv6eev.insjC2Va()
	if insjC2Va < 0 {

		yY3pgXWox3j = true
		insjC2Va = -insjC2Va
		aCiH2Z7qWLtw = -aCiH2Z7qWLtw
		if aCiH2Z7qWLtw < 0 {
			aCiH2Z7qWLtw += 1e9
			insjC2Va--
		}
	}

	switch {

	case lNAwkbJQVBgU < Second && Second%(lNAwkbJQVBgU+lNAwkbJQVBgU) == 0:
		npEj_3Oa7V =  /*line :1*/ int(aCiH2Z7qWLtw/ /*line :1*/ int32(lNAwkbJQVBgU)) & 1
		aLhBMRzpWd =  /*line :1*/ GKMwTGxQa0S(aCiH2Z7qWLtw %  /*line :1*/ int32(lNAwkbJQVBgU))

	case lNAwkbJQVBgU%Second == 0:
		kIMGbZ :=  /*line :1*/ int64(lNAwkbJQVBgU / Second)
		npEj_3Oa7V =  /*line :1*/ int(insjC2Va/kIMGbZ) & 1
		aLhBMRzpWd =  /*line :1*/ GKMwTGxQa0S(insjC2Va%kIMGbZ)*Second +  /*line :1*/ GKMwTGxQa0S(aCiH2Z7qWLtw)

	default:

		insjC2Va :=  /*line :1*/ uint64(insjC2Va)
		mV2agu78J := (insjC2Va >> 32) * 1e9
		jz_9USGNKd := mV2agu78J >> 32
		lmfyR08mHPXR := mV2agu78J << 32
		mV2agu78J = (insjC2Va & 0xFFFFFFFF) * 1e9
		q11lsQmAi8Nn, lmfyR08mHPXR := lmfyR08mHPXR, lmfyR08mHPXR+mV2agu78J
		if lmfyR08mHPXR < q11lsQmAi8Nn {
			jz_9USGNKd++
		}
		q11lsQmAi8Nn, lmfyR08mHPXR = lmfyR08mHPXR, lmfyR08mHPXR+ /*line :1*/ uint64(aCiH2Z7qWLtw)
		if lmfyR08mHPXR < q11lsQmAi8Nn {
			jz_9USGNKd++
		}

		kIMGbZ :=  /*line :1*/ uint64(lNAwkbJQVBgU)
		for kIMGbZ>>63 != 1 {
			kIMGbZ <<= 1
		}
		ckWaYo1 :=  /*line :1*/ uint64(0)
		for {
			npEj_3Oa7V = 0
			if jz_9USGNKd > kIMGbZ || jz_9USGNKd == kIMGbZ && lmfyR08mHPXR >= ckWaYo1 {

				npEj_3Oa7V = 1
				q11lsQmAi8Nn, lmfyR08mHPXR = lmfyR08mHPXR, lmfyR08mHPXR-ckWaYo1
				if lmfyR08mHPXR > q11lsQmAi8Nn {
					jz_9USGNKd--
				}
				jz_9USGNKd -= kIMGbZ
			}
			if kIMGbZ == 0 && ckWaYo1 ==  /*line :1*/ uint64(lNAwkbJQVBgU) {
				break
			}
			ckWaYo1 >>= 1
			ckWaYo1 |= (kIMGbZ & 1) << 63
			kIMGbZ >>= 1
		}
		aLhBMRzpWd =  /*line :1*/ GKMwTGxQa0S(lmfyR08mHPXR)
	}

	if yY3pgXWox3j && aLhBMRzpWd != 0 {

		npEj_3Oa7V ^= 1
		aLhBMRzpWd = lNAwkbJQVBgU - aLhBMRzpWd
	}
	return
}
