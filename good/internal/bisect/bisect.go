//line :1












































































































































































package vWnmvaJKxA

import (
	"runtime"
	sync "sync"
	atomic "sync/atomic"
	"unsafe"
)









func MbSZaWBD(hUCRY2 string) (*SiHlqZVr7s, error) {
	if hUCRY2 == "" {
		return nil, nil
	}

	f1wxHZ :=  /*line :1*/ new(SiHlqZVr7s)

	dM7t3H := hUCRY2

	if  /*line :1*/ len(dM7t3H) > 0 && dM7t3H[0] == 'q' {
		f1wxHZ.iaiJ_Bk4n = true
		dM7t3H = dM7t3H[1:]
		if dM7t3H == "" {
			return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
		}
	}

	for  /*line :1*/ len(dM7t3H) > 0 && dM7t3H[0] == 'v' {
		f1wxHZ.b7vr3LZja = true
		f1wxHZ.iaiJ_Bk4n = false
		dM7t3H = dM7t3H[1:]
		if dM7t3H == "" {
			return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
		}
	}

	f1wxHZ.ajYeUNeepGJ = true
	for  /*line :1*/ len(dM7t3H) > 0 && dM7t3H[0] == '!' {
		f1wxHZ.ajYeUNeepGJ = !f1wxHZ.ajYeUNeepGJ
		dM7t3H = dM7t3H[1:]
		if dM7t3H == "" {
			return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
		}
	}

	if dM7t3H == "n" {

		f1wxHZ.ajYeUNeepGJ = !f1wxHZ.ajYeUNeepGJ
		dM7t3H = "y"
	}

	z6DbHLE6 := true
	gmc4JNpOk :=  /*line :1*/ uint64(0)
	sKJfiy := 0
	rMXbIomdRMVJ := 1
	for vAhtD0o5o := 0; vAhtD0o5o <=  /*line :1*/ len(dM7t3H); vAhtD0o5o++ {

		q_us4kxkSDVA :=  /*line :1*/ byte('-')
		if vAhtD0o5o <  /*line :1*/ len(dM7t3H) {
			q_us4kxkSDVA = dM7t3H[vAhtD0o5o]
		}
		if vAhtD0o5o == sKJfiy && rMXbIomdRMVJ == 1 && q_us4kxkSDVA == 'x' {
			sKJfiy = vAhtD0o5o + 1
			rMXbIomdRMVJ = 4
			continue
		}
		switch q_us4kxkSDVA {
		default:
			return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
		case '2', '3', '4', '5', '6', '7', '8', '9':
			if rMXbIomdRMVJ != 4 {
				return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
			}
			fallthrough
		case '0', '1':
			gmc4JNpOk <<= rMXbIomdRMVJ
			gmc4JNpOk |=  /*line :1*/ uint64(q_us4kxkSDVA - '0')
		case 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
			if rMXbIomdRMVJ != 4 {
				return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
			}
			gmc4JNpOk <<= 4
			gmc4JNpOk |=  /*line :1*/ uint64(q_us4kxkSDVA&^0x20 - 'A' + 10)
		case 'y':
			if vAhtD0o5o+1 <  /*line :1*/ len(dM7t3H) && (dM7t3H[vAhtD0o5o+1] == '0' || dM7t3H[vAhtD0o5o+1] == '1') {
				return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
			}
			gmc4JNpOk = 0
		case '+', '-':
			if q_us4kxkSDVA == '+' && z6DbHLE6 == false {

				return nil, &wuhed4w3I0{"invalid pattern syntax (+ after -): " + hUCRY2}
			}
			if vAhtD0o5o > 0 {
				uawNhoQfX20c := (vAhtD0o5o - sKJfiy) * rMXbIomdRMVJ
				if uawNhoQfX20c > 64 {
					return nil, &wuhed4w3I0{"pattern bits too long: " + hUCRY2}
				}
				if uawNhoQfX20c <= 0 {
					return nil, &wuhed4w3I0{"invalid pattern syntax: " + hUCRY2}
				}
				if dM7t3H[sKJfiy] == 'y' {
					uawNhoQfX20c = 0
				}
				otZsI4MAvx :=  /*line :1*/ uint64(1)<<uawNhoQfX20c - 1
				f1wxHZ.mhQfwvYTw =  /*line :1*/ append(f1wxHZ.mhQfwvYTw, wss6GqgX4r{otZsI4MAvx, gmc4JNpOk, z6DbHLE6})
			} else if q_us4kxkSDVA == '-' {

				f1wxHZ.mhQfwvYTw =  /*line :1*/ append(f1wxHZ.mhQfwvYTw, wss6GqgX4r{0, 0, true})
			}
			gmc4JNpOk = 0
			z6DbHLE6 = q_us4kxkSDVA == '+'
			sKJfiy = vAhtD0o5o + 1
			rMXbIomdRMVJ = 1
		}
	}
	return f1wxHZ, nil
}



type SiHlqZVr7s struct {
	b7vr3LZja	bool	
	iaiJ_Bk4n	bool	
	ajYeUNeepGJ	bool	
	mhQfwvYTw	[]wss6GqgX4r	
	gzgCsdvaGHS	j2kpaHtuit
}




type j2kpaHtuit struct {
	jguZEFaw unsafe.Pointer
}

func (dM7t3H *j2kpaHtuit) Load() *iWWdbQ {
	return (* /*line :1*/ iWWdbQ)( /*line :1*/ atomic.LoadPointer(&dM7t3H.jguZEFaw))
}

func (dM7t3H *j2kpaHtuit) CompareAndSwap(pjCCPHHH, bvShWNF *iWWdbQ) bool {
	return  /*line :1*/ atomic.B2CfK2wiOXKp(&dM7t3H.jguZEFaw,  /*line :1*/ unsafe.Pointer(pjCCPHHH),  /*line :1*/ unsafe.Pointer(bvShWNF))
}



type wss6GqgX4r struct {
	qn2_M7VsiH	uint64
	pvRf3efg	uint64
	cEWbwlqJ58oH	bool
}





func (f1wxHZ *SiHlqZVr7s) MarkerOnly() bool {
	return !f1wxHZ.b7vr3LZja
}


func (f1wxHZ *SiHlqZVr7s) ShouldEnable(ukvS8e9A uint64) bool {
	if f1wxHZ == nil {
		return true
	}
	return  /*line :1*/ f1wxHZ.ttCzxeLu1hx(ukvS8e9A) == f1wxHZ.ajYeUNeepGJ
}


func (f1wxHZ *SiHlqZVr7s) ShouldPrint(ukvS8e9A uint64) bool {
	if f1wxHZ == nil || f1wxHZ.iaiJ_Bk4n {
		return false
	}
	return  /*line :1*/ f1wxHZ.ttCzxeLu1hx(ukvS8e9A)
}


func (f1wxHZ *SiHlqZVr7s) ttCzxeLu1hx(ukvS8e9A uint64) bool {
	for vAhtD0o5o :=  /*line :1*/ len(f1wxHZ.mhQfwvYTw) - 1; vAhtD0o5o >= 0; vAhtD0o5o-- {
		q_us4kxkSDVA := &f1wxHZ.mhQfwvYTw[vAhtD0o5o]
		if ukvS8e9A&q_us4kxkSDVA.qn2_M7VsiH == q_us4kxkSDVA.pvRf3efg {
			return q_us4kxkSDVA.cEWbwlqJ58oH
		}
	}
	return false
}



func (f1wxHZ *SiHlqZVr7s) FileLine(wgh88w9z HSqgpar5xh, iTvY831 string, aIl2TmtEN9_L int) bool {
	if f1wxHZ == nil {
		return true
	}
	return  /*line :1*/ f1wxHZ.cfT7ot(wgh88w9z, iTvY831, aIl2TmtEN9_L)
}



func (f1wxHZ *SiHlqZVr7s) cfT7ot(wgh88w9z HSqgpar5xh, iTvY831 string, aIl2TmtEN9_L int) bool {
	slFil0 :=  /*line :1*/ T3ZaLuwu1HxR(iTvY831, aIl2TmtEN9_L)
	if  /*line :1*/ f1wxHZ.ShouldPrint(slFil0) {
		if  /*line :1*/ f1wxHZ.MarkerOnly() {
			 /*line :1*/ ZrHD9oMQ(wgh88w9z, slFil0)
		} else {
			 /*line :1*/ xe6MmKn0wdX(wgh88w9z, slFil0, iTvY831, aIl2TmtEN9_L)
		}
	}
	return  /*line :1*/ f1wxHZ.ShouldEnable(slFil0)
}


func xe6MmKn0wdX(wgh88w9z HSqgpar5xh, slFil0 uint64, iTvY831 string, aIl2TmtEN9_L int) error {
	const markerLen = 40	
	xp8l4qehkecp :=  /*line :1*/ make([]byte, 0, markerLen+ /*line :1*/ len(iTvY831)+24)
	xp8l4qehkecp =  /*line :1*/ RZoMcBFi3FB(xp8l4qehkecp, slFil0)
	xp8l4qehkecp =  /*line :1*/ gt_HbE(xp8l4qehkecp, iTvY831, aIl2TmtEN9_L)
	xp8l4qehkecp =  /*line :1*/ append(xp8l4qehkecp, '\n')
	_, aMbxkt5no :=  /*line :1*/ wgh88w9z.Write(xp8l4qehkecp)
	return aMbxkt5no
}


func gt_HbE(dQ1hkbo0aOr []byte, iTvY831 string, aIl2TmtEN9_L int) []byte {
	dQ1hkbo0aOr =  /*line :1*/ append(dQ1hkbo0aOr, iTvY831...)
	dQ1hkbo0aOr =  /*line :1*/ append(dQ1hkbo0aOr, ':')
	h4HgGW7E :=  /*line :1*/ uint(aIl2TmtEN9_L)
	if aIl2TmtEN9_L < 0 {
		dQ1hkbo0aOr =  /*line :1*/ append(dQ1hkbo0aOr, '-')
		h4HgGW7E = -h4HgGW7E
	}
	var dym2QSvemCA [24]byte
	vAhtD0o5o :=  /*line :1*/ len(dym2QSvemCA)
	for vAhtD0o5o ==  /*line :1*/ len(dym2QSvemCA) || h4HgGW7E > 0 {
		vAhtD0o5o--
		dym2QSvemCA[vAhtD0o5o] = '0' +  /*line :1*/ byte(h4HgGW7E%10)
		h4HgGW7E /= 10
	}
	dQ1hkbo0aOr =  /*line :1*/ append(dQ1hkbo0aOr, dym2QSvemCA[vAhtD0o5o:]...)
	return dQ1hkbo0aOr
}




func (f1wxHZ *SiHlqZVr7s) Stack(wgh88w9z HSqgpar5xh) bool {
	if f1wxHZ == nil {
		return true
	}
	return  /*line :1*/ f1wxHZ.cjehzu6(wgh88w9z)
}



func (f1wxHZ *SiHlqZVr7s) cjehzu6(wgh88w9z HSqgpar5xh) bool {
	const maxStack = 16
	var mafA4pmyWs [maxStack]uintptr
	uawNhoQfX20c :=  /*line :1*/ runtime.Callers(2, mafA4pmyWs[:])

	if uawNhoQfX20c <= 1 {
		return false
	}

	beVBGgNW4u3E := mafA4pmyWs[0]

	for vAhtD0o5o := range mafA4pmyWs[:uawNhoQfX20c] {
		mafA4pmyWs[vAhtD0o5o] -= beVBGgNW4u3E
	}

	slFil0 :=  /*line :1*/ T3ZaLuwu1HxR(mafA4pmyWs[:uawNhoQfX20c])
	if  /*line :1*/ f1wxHZ.ShouldPrint(slFil0) {
		var tcdma7 *iWWdbQ
		for {
			tcdma7 =  /*line :1*/ f1wxHZ.gzgCsdvaGHS.Load()
			if tcdma7 != nil {
				break
			}
			tcdma7 =  /*line :1*/ new(iWWdbQ)
			if  /*line :1*/ f1wxHZ.gzgCsdvaGHS.CompareAndSwap(nil, tcdma7) {
				break
			}
		}

		if  /*line :1*/ f1wxHZ.MarkerOnly() {
			if ! /*line :1*/ tcdma7.cNg3MoK(slFil0) {
				 /*line :1*/ ZrHD9oMQ(wgh88w9z, slFil0)
			}
		} else {
			if ! /*line :1*/ tcdma7.eteVtI(slFil0) {

				for vAhtD0o5o := range mafA4pmyWs[:uawNhoQfX20c] {
					mafA4pmyWs[vAhtD0o5o] += beVBGgNW4u3E
				}
				 /*line :1*/ e6yaHN(wgh88w9z, slFil0, mafA4pmyWs[1:uawNhoQfX20c])
			}
		}
	}
	return  /*line :1*/ f1wxHZ.ShouldEnable(slFil0)

}



type HSqgpar5xh interface {
	Write([]byte) (int, error)
}



func ZrHD9oMQ(wgh88w9z HSqgpar5xh, slFil0 uint64) error {
	var dym2QSvemCA [50]byte
	xp8l4qehkecp :=  /*line :1*/ RZoMcBFi3FB(dym2QSvemCA[:], slFil0)
	xp8l4qehkecp =  /*line :1*/ append(xp8l4qehkecp, '\n')
	_, aMbxkt5no :=  /*line :1*/ wgh88w9z.Write(xp8l4qehkecp)
	return aMbxkt5no
}



func e6yaHN(wgh88w9z HSqgpar5xh, slFil0 uint64, mafA4pmyWs []uintptr) error {
	dym2QSvemCA :=  /*line :1*/ make([]byte, 0, 2048)

	var tE2lnDNumrSh [100]byte
	j3eD1_KAa4 :=  /*line :1*/ RZoMcBFi3FB(tE2lnDNumrSh[:0], slFil0)

	lQyphY :=  /*line :1*/ runtime.CallersFrames(mafA4pmyWs)
	for {
		jFespwgawa, ePzgnU :=  /*line :1*/ lQyphY.Next()
		dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA, j3eD1_KAa4...)
		dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA,  /*line :1*/ jFespwgawa.Func.Name()...)
		dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA, "()\n"...)
		dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA, j3eD1_KAa4...)
		dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA, '\t')
		dym2QSvemCA =  /*line :1*/ gt_HbE(dym2QSvemCA, jFespwgawa.File, jFespwgawa.Line)
		dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA, '\n')
		if !ePzgnU {
			break
		}
	}
	dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA, j3eD1_KAa4...)
	dym2QSvemCA =  /*line :1*/ append(dym2QSvemCA, '\n')
	_, aMbxkt5no :=  /*line :1*/ wgh88w9z.Write(dym2QSvemCA)
	return aMbxkt5no
}




func Ha7IaImmq_nU(ukvS8e9A uint64) string {
	return  /*line :1*/ string( /*line :1*/ RZoMcBFi3FB(nil, ukvS8e9A))
}


func RZoMcBFi3FB(dQ1hkbo0aOr []byte, ukvS8e9A uint64) []byte {
	const prefix = "[bisect-match 0x"
	var dym2QSvemCA [ /*line :1*/ len(prefix) + 16 + 1]byte
	 /*line :1*/ copy(dym2QSvemCA[:], prefix)
	for vAhtD0o5o := 0; vAhtD0o5o < 16; vAhtD0o5o++ {
		dym2QSvemCA[ /*line :1*/ len(prefix)+vAhtD0o5o] = "0123456789abcdef"[ukvS8e9A>>60]
		ukvS8e9A <<= 4
	}
	dym2QSvemCA[ /*line :1*/ len(prefix)+16] = ']'
	return  /*line :1*/ append(dQ1hkbo0aOr, dym2QSvemCA[:]...)
}






func IDxAw63Y7aLU(aIl2TmtEN9_L string) (nasU1PP7UKc string, ukvS8e9A uint64, ukWvu7k bool) {

	j3eD1_KAa4 := "[bisect-match "
	vAhtD0o5o := 0
	for ; ; vAhtD0o5o++ {
		if vAhtD0o5o >=  /*line :1*/ len(aIl2TmtEN9_L)- /*line :1*/ len(j3eD1_KAa4) {
			return aIl2TmtEN9_L, 0, false
		}
		if aIl2TmtEN9_L[vAhtD0o5o] == '[' && aIl2TmtEN9_L[vAhtD0o5o:vAhtD0o5o+ /*line :1*/ len(j3eD1_KAa4)] == j3eD1_KAa4 {
			break
		}
	}

	b6Gc8aA := vAhtD0o5o +  /*line :1*/ len(j3eD1_KAa4)
	for b6Gc8aA <  /*line :1*/ len(aIl2TmtEN9_L) && aIl2TmtEN9_L[b6Gc8aA] != ']' {
		b6Gc8aA++
	}
	if b6Gc8aA >=  /*line :1*/ len(aIl2TmtEN9_L) {
		return aIl2TmtEN9_L, 0, false
	}

	nZ2oCRnll := aIl2TmtEN9_L[vAhtD0o5o+ /*line :1*/ len(j3eD1_KAa4) : b6Gc8aA]
	if  /*line :1*/ len(nZ2oCRnll) >= 3 && nZ2oCRnll[:2] == "0x" {

		if  /*line :1*/ len(nZ2oCRnll) > 2+16 {
			return aIl2TmtEN9_L, 0, false
		}
		for vAhtD0o5o := 2; vAhtD0o5o <  /*line :1*/ len(nZ2oCRnll); vAhtD0o5o++ {
			ukvS8e9A <<= 4
			switch q_us4kxkSDVA := nZ2oCRnll[vAhtD0o5o]; {
			case '0' <= q_us4kxkSDVA && q_us4kxkSDVA <= '9':
				ukvS8e9A |=  /*line :1*/ uint64(q_us4kxkSDVA - '0')
			case 'a' <= q_us4kxkSDVA && q_us4kxkSDVA <= 'f':
				ukvS8e9A |=  /*line :1*/ uint64(q_us4kxkSDVA - 'a' + 10)
			case 'A' <= q_us4kxkSDVA && q_us4kxkSDVA <= 'F':
				ukvS8e9A |=  /*line :1*/ uint64(q_us4kxkSDVA - 'A' + 10)
			}
		}
	} else {
		if nZ2oCRnll == "" ||  /*line :1*/ len(nZ2oCRnll) > 64 {
			return aIl2TmtEN9_L, 0, false
		}

		for vAhtD0o5o := 0; vAhtD0o5o <  /*line :1*/ len(nZ2oCRnll); vAhtD0o5o++ {
			ukvS8e9A <<= 1
			switch q_us4kxkSDVA := nZ2oCRnll[vAhtD0o5o]; q_us4kxkSDVA {
			default:
				return aIl2TmtEN9_L, 0, false
			case '0', '1':
				ukvS8e9A |=  /*line :1*/ uint64(q_us4kxkSDVA - '0')
			}
		}
	}

	b6Gc8aA++
	if vAhtD0o5o > 0 && aIl2TmtEN9_L[vAhtD0o5o-1] == ' ' {
		vAhtD0o5o--
	} else if b6Gc8aA <  /*line :1*/ len(aIl2TmtEN9_L) && aIl2TmtEN9_L[b6Gc8aA] == ' ' {
		b6Gc8aA++
	}
	nasU1PP7UKc = aIl2TmtEN9_L[:vAhtD0o5o] + aIl2TmtEN9_L[b6Gc8aA:]
	return nasU1PP7UKc, ukvS8e9A, true
}



func T3ZaLuwu1HxR(pEDHaKCg ...any) uint64 {
	slFil0 := offset64
	for _, mVCpYstUzXJo := range pEDHaKCg {
		switch mVCpYstUzXJo := mVCpYstUzXJo.(type) {
		default:

			 /*line :1*/ panic("bisect.Hash: unexpected argument type")
		case string:
			slFil0 =  /*line :1*/ veiMSYd0s_(slFil0, mVCpYstUzXJo)
		case byte:
			slFil0 =  /*line :1*/ n_X4aRn(slFil0, mVCpYstUzXJo)
		case int:
			slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(mVCpYstUzXJo))
		case uint:
			slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(mVCpYstUzXJo))
		case int32:
			slFil0 =  /*line :1*/ viHaSUjvfVl1(slFil0,  /*line :1*/ uint32(mVCpYstUzXJo))
		case uint32:
			slFil0 =  /*line :1*/ viHaSUjvfVl1(slFil0, mVCpYstUzXJo)
		case int64:
			slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(mVCpYstUzXJo))
		case uint64:
			slFil0 =  /*line :1*/ aiRirrVk(slFil0, mVCpYstUzXJo)
		case uintptr:
			slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(mVCpYstUzXJo))
		case []string:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ veiMSYd0s_(slFil0, uKsntT)
			}
		case []byte:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ n_X4aRn(slFil0, uKsntT)
			}
		case []int:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(uKsntT))
			}
		case []uint:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(uKsntT))
			}
		case []int32:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ viHaSUjvfVl1(slFil0,  /*line :1*/ uint32(uKsntT))
			}
		case []uint32:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ viHaSUjvfVl1(slFil0, uKsntT)
			}
		case []int64:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(uKsntT))
			}
		case []uint64:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ aiRirrVk(slFil0, uKsntT)
			}
		case []uintptr:
			for _, uKsntT := range mVCpYstUzXJo {
				slFil0 =  /*line :1*/ aiRirrVk(slFil0,  /*line :1*/ uint64(uKsntT))
			}
		}
	}
	return slFil0
}



type wuhed4w3I0 struct{ asjl8_B string }

func (bIfjaq *wuhed4w3I0) Error() string	{ return bIfjaq.asjl8_B }

const (
	offset64	uint64	= 14695981039346656037
	prime64	uint64	= 1099511628211
)

func n_X4aRn(slFil0 uint64, uKsntT byte) uint64 {
	slFil0 ^=  /*line :1*/ uint64(uKsntT)
	slFil0 *= prime64
	return slFil0
}

func veiMSYd0s_(slFil0 uint64, uKsntT string) uint64 {
	for vAhtD0o5o := 0; vAhtD0o5o <  /*line :1*/ len(uKsntT); vAhtD0o5o++ {
		slFil0 ^=  /*line :1*/ uint64(uKsntT[vAhtD0o5o])
		slFil0 *= prime64
	}
	return slFil0
}

func aiRirrVk(slFil0 uint64, uKsntT uint64) uint64 {
	for vAhtD0o5o := 0; vAhtD0o5o < 8; vAhtD0o5o++ {
		slFil0 ^=  /*line :1*/ uint64(uKsntT & 0xFF)
		uKsntT >>= 8
		slFil0 *= prime64
	}
	return slFil0
}

func viHaSUjvfVl1(slFil0 uint64, uKsntT uint32) uint64 {
	for vAhtD0o5o := 0; vAhtD0o5o < 4; vAhtD0o5o++ {
		slFil0 ^=  /*line :1*/ uint64(uKsntT & 0xFF)
		uKsntT >>= 8
		slFil0 *= prime64
	}
	return slFil0
}








type iWWdbQ struct {
	
	uiZNqu	[128][4]uint64

	
	a5oa26bF	sync.DIRWe11KYlYa
	hF4VWAfx	map[uint64]bool
}



func (tcdma7 *iWWdbQ) eteVtI(slFil0 uint64) bool {
	 /*line :1*/ tcdma7.a5oa26bF.Lock()
	if tcdma7.hF4VWAfx == nil {
		tcdma7.hF4VWAfx =  /*line :1*/ make(map[uint64]bool)
	}
	eteVtI := tcdma7.hF4VWAfx[slFil0]
	tcdma7.hF4VWAfx[slFil0] = true
	 /*line :1*/ tcdma7.a5oa26bF.Unlock()
	return eteVtI
}





func (tcdma7 *iWWdbQ) cNg3MoK(slFil0 uint64) bool {
	adjgYlR6 := &tcdma7.uiZNqu[ /*line :1*/ uint(slFil0)% /*line :1*/ uint( /*line :1*/ len(tcdma7.uiZNqu))]
	for vAhtD0o5o := 0; vAhtD0o5o <  /*line :1*/ len(adjgYlR6); vAhtD0o5o++ {
		if  /*line :1*/ atomic.LoadUint64(&adjgYlR6[vAhtD0o5o]) == slFil0 {
			return true
		}
	}

	dM2TZpEsx0l := offset64
	for _, uKsntT := range adjgYlR6 {
		dM2TZpEsx0l =  /*line :1*/ aiRirrVk(dM2TZpEsx0l, uKsntT)
	}
	 /*line :1*/ atomic.StoreUint64(&adjgYlR6[ /*line :1*/ uint(dM2TZpEsx0l)% /*line :1*/ uint( /*line :1*/ len(adjgYlR6))], slFil0)
	return false
}
