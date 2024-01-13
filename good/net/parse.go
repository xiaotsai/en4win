//line :1
package gF9mZs2

import (
	"internal/bytealg"
	io "xI9ai2djaJ2"
	os "hPMrClpC"
	time "fRAfQd_"
)

type cZsWMjIJ6 struct {
	pdlGQms	*os.BvGocYcXx
	aSGgDnd	[]byte
	j3PKafDVSR6F	bool
}

func (t5q9DVOD9 *cZsWMjIJ6) tFIv2Ppbx0H()	{  /*line :1*/ t5q9DVOD9.pdlGQms.Close() }

func (t5q9DVOD9 *cZsWMjIJ6) hsmLOB3() (dRoFccaZ string, d30HIwxht bool) {
	bwennb := t5q9DVOD9.aSGgDnd
	eaAUaB7Z := 0
	for eaAUaB7Z = 0; eaAUaB7Z <  /*line :1*/ len(bwennb); eaAUaB7Z++ {
		if bwennb[eaAUaB7Z] == '\n' {
			dRoFccaZ =  /*line :1*/ string(bwennb[0:eaAUaB7Z])
			d30HIwxht = true

			eaAUaB7Z++
			doauF8Y :=  /*line :1*/ len(bwennb) - eaAUaB7Z
			 /*line :1*/ copy(bwennb[0:], bwennb[eaAUaB7Z:])
			t5q9DVOD9.aSGgDnd = bwennb[0:doauF8Y]
			return
		}
	}
	if t5q9DVOD9.j3PKafDVSR6F &&  /*line :1*/ len(t5q9DVOD9.aSGgDnd) > 0 {

		dRoFccaZ =  /*line :1*/ string(bwennb)
		t5q9DVOD9.aSGgDnd = t5q9DVOD9.aSGgDnd[0:0]
		d30HIwxht = true
	}
	return
}

func (t5q9DVOD9 *cZsWMjIJ6) dM_MFaU7Ms() (dRoFccaZ string, d30HIwxht bool) {
	if dRoFccaZ, d30HIwxht =  /*line :1*/ t5q9DVOD9.hsmLOB3(); d30HIwxht {
		return
	}
	if  /*line :1*/ len(t5q9DVOD9.aSGgDnd) <  /*line :1*/ cap(t5q9DVOD9.aSGgDnd) {
		dhvebBa4tmD :=  /*line :1*/ len(t5q9DVOD9.aSGgDnd)
		doauF8Y, h_ljk48Bm :=  /*line :1*/ io.G0zBgKg(t5q9DVOD9.pdlGQms, t5q9DVOD9.aSGgDnd[dhvebBa4tmD: /*line :1*/ cap(t5q9DVOD9.aSGgDnd)])
		if doauF8Y >= 0 {
			t5q9DVOD9.aSGgDnd = t5q9DVOD9.aSGgDnd[0 : dhvebBa4tmD+doauF8Y]
		}
		if h_ljk48Bm == io.K5Sqco || h_ljk48Bm == io.UASgojMNPA {
			t5q9DVOD9.j3PKafDVSR6F = true
		}
	}
	dRoFccaZ, d30HIwxht =  /*line :1*/ t5q9DVOD9.hsmLOB3()
	return
}

func (t5q9DVOD9 *cZsWMjIJ6) o6GKG3liN() (i81SGZwwa time.G4KDkI, rz5yRs4IdH int64, h_ljk48Bm error) {
	olWvL6jy7u, h_ljk48Bm :=  /*line :1*/ t5q9DVOD9.pdlGQms.Stat()
	if h_ljk48Bm != nil {
		return time.G4KDkI{}, 0, h_ljk48Bm
	}
	return  /*line :1*/ olWvL6jy7u.ModTime(),  /*line :1*/ olWvL6jy7u.Size(), nil
}

func _qDAPa7V(tahgc783Bc string) (*cZsWMjIJ6, error) {
	vyaiiPXm6k_W, h_ljk48Bm :=  /*line :1*/ os.Cvk6wxcjw(tahgc783Bc)
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	return &cZsWMjIJ6{vyaiiPXm6k_W,  /*line :1*/ make([]byte, 0, 64*1024), false}, nil
}

func o6GKG3liN(tahgc783Bc string) (i81SGZwwa time.G4KDkI, rz5yRs4IdH int64, h_ljk48Bm error) {
	olWvL6jy7u, h_ljk48Bm :=  /*line :1*/ os.ZYa3wuv(tahgc783Bc)
	if h_ljk48Bm != nil {
		return time.G4KDkI{}, 0, h_ljk48Bm
	}
	return  /*line :1*/ olWvL6jy7u.ModTime(),  /*line :1*/ olWvL6jy7u.Size(), nil
}


func piUuwjgxa(dRoFccaZ string, aeaqWzxJu string) int {
	doauF8Y := 0
	for eaAUaB7Z := 0; eaAUaB7Z <  /*line :1*/ len(dRoFccaZ); eaAUaB7Z++ {
		if  /*line :1*/ bytealg.IndexByteString(aeaqWzxJu, dRoFccaZ[eaAUaB7Z]) >= 0 {
			doauF8Y++
		}
	}
	return doauF8Y
}


func tWf8ZvH9Jb1n(dRoFccaZ string, aeaqWzxJu string) []string {
	uI7LZDHm6 :=  /*line :1*/ make([]string, 1+ /*line :1*/ piUuwjgxa(dRoFccaZ, aeaqWzxJu))
	doauF8Y := 0
	vMwVxD_U := 0
	for eaAUaB7Z := 0; eaAUaB7Z <  /*line :1*/ len(dRoFccaZ); eaAUaB7Z++ {
		if  /*line :1*/ bytealg.IndexByteString(aeaqWzxJu, dRoFccaZ[eaAUaB7Z]) >= 0 {
			if vMwVxD_U < eaAUaB7Z {
				uI7LZDHm6[doauF8Y] = dRoFccaZ[vMwVxD_U:eaAUaB7Z]
				doauF8Y++
			}
			vMwVxD_U = eaAUaB7Z + 1
		}
	}
	if vMwVxD_U <  /*line :1*/ len(dRoFccaZ) {
		uI7LZDHm6[doauF8Y] = dRoFccaZ[vMwVxD_U:]
		doauF8Y++
	}
	return uI7LZDHm6[0:doauF8Y]
}

func nhtSVnZ(dRoFccaZ string) []string	{ return  /*line :1*/ tWf8ZvH9Jb1n(dRoFccaZ, " \r\t\n") }


const big = 0xFFFFFF



func bEihJjrbz(dRoFccaZ string) (doauF8Y int, eaAUaB7Z int, d30HIwxht bool) {
	doauF8Y = 0
	for eaAUaB7Z = 0; eaAUaB7Z <  /*line :1*/ len(dRoFccaZ) && '0' <= dRoFccaZ[eaAUaB7Z] && dRoFccaZ[eaAUaB7Z] <= '9'; eaAUaB7Z++ {
		doauF8Y = doauF8Y*10 +  /*line :1*/ int(dRoFccaZ[eaAUaB7Z]-'0')
		if doauF8Y >= big {
			return big, eaAUaB7Z, false
		}
	}
	if eaAUaB7Z == 0 {
		return 0, 0, false
	}
	return doauF8Y, eaAUaB7Z, true
}



func kUEPdSe(dRoFccaZ string) (doauF8Y int, eaAUaB7Z int, d30HIwxht bool) {
	doauF8Y = 0
	for eaAUaB7Z = 0; eaAUaB7Z <  /*line :1*/ len(dRoFccaZ); eaAUaB7Z++ {
		if '0' <= dRoFccaZ[eaAUaB7Z] && dRoFccaZ[eaAUaB7Z] <= '9' {
			doauF8Y *= 16
			doauF8Y +=  /*line :1*/ int(dRoFccaZ[eaAUaB7Z] - '0')
		} else if 'a' <= dRoFccaZ[eaAUaB7Z] && dRoFccaZ[eaAUaB7Z] <= 'f' {
			doauF8Y *= 16
			doauF8Y +=  /*line :1*/ int(dRoFccaZ[eaAUaB7Z]-'a') + 10
		} else if 'A' <= dRoFccaZ[eaAUaB7Z] && dRoFccaZ[eaAUaB7Z] <= 'F' {
			doauF8Y *= 16
			doauF8Y +=  /*line :1*/ int(dRoFccaZ[eaAUaB7Z]-'A') + 10
		} else {
			break
		}
		if doauF8Y >= big {
			return 0, eaAUaB7Z, false
		}
	}
	if eaAUaB7Z == 0 {
		return 0, eaAUaB7Z, false
	}
	return doauF8Y, eaAUaB7Z, true
}





func kHN6rj(dRoFccaZ string, wI0aaz byte) (byte, bool) {
	if  /*line :1*/ len(dRoFccaZ) > 2 && dRoFccaZ[2] != wI0aaz {
		return 0, false
	}
	doauF8Y, fXGZAvQ_L, d30HIwxht :=  /*line :1*/ kUEPdSe(dRoFccaZ[:2])
	return  /*line :1*/ byte(doauF8Y), d30HIwxht && fXGZAvQ_L == 2
}


func fSfDru(xnC59rPcIS []byte, eaAUaB7Z uint32) []byte {
	if eaAUaB7Z == 0 {
		return  /*line :1*/ append(xnC59rPcIS, '0')
	}
	for kVgtfLFZ := 7; kVgtfLFZ >= 0; kVgtfLFZ-- {
		ljsCSk := eaAUaB7Z >>  /*line :1*/ uint(kVgtfLFZ*4)
		if ljsCSk > 0 {
			xnC59rPcIS =  /*line :1*/ append(xnC59rPcIS, hexDigit[ljsCSk&0xf])
		}
	}
	return xnC59rPcIS
}


func mB0PZSLg9(dRoFccaZ string, fIadEXIim6l byte) int {
	doauF8Y := 0
	for eaAUaB7Z := 0; eaAUaB7Z <  /*line :1*/ len(dRoFccaZ); eaAUaB7Z++ {
		if dRoFccaZ[eaAUaB7Z] == fIadEXIim6l {
			doauF8Y++
		}
	}
	return doauF8Y
}


func vMwVxD_U(dRoFccaZ string, fIadEXIim6l byte) int {
	eaAUaB7Z :=  /*line :1*/ len(dRoFccaZ)
	for eaAUaB7Z--; eaAUaB7Z >= 0; eaAUaB7Z-- {
		if dRoFccaZ[eaAUaB7Z] == fIadEXIim6l {
			break
		}
	}
	return eaAUaB7Z
}


func wr9w1HG(dRoFccaZ string) bool {
	for eaAUaB7Z := range dRoFccaZ {
		if 'A' <= dRoFccaZ[eaAUaB7Z] && dRoFccaZ[eaAUaB7Z] <= 'Z' {
			return true
		}
	}
	return false
}


func kuRb7e(tWO2olI []byte) {
	for eaAUaB7Z, fIadEXIim6l := range tWO2olI {
		if 'A' <= fIadEXIim6l && fIadEXIim6l <= 'Z' {
			tWO2olI[eaAUaB7Z] += 'a' - 'A'
		}
	}
}


func erPIFF(fIadEXIim6l byte) byte {
	if 'A' <= fIadEXIim6l && fIadEXIim6l <= 'Z' {
		return fIadEXIim6l + ('a' - 'A')
	}
	return fIadEXIim6l
}


func eEa5QIpda3t(tWO2olI string) string {
	for  /*line :1*/ len(tWO2olI) > 0 &&  /*line :1*/ ns5_ZXVgtnn(tWO2olI[0]) {
		tWO2olI = tWO2olI[1:]
	}
	for  /*line :1*/ len(tWO2olI) > 0 &&  /*line :1*/ ns5_ZXVgtnn(tWO2olI[ /*line :1*/ len(tWO2olI)-1]) {
		tWO2olI = tWO2olI[: /*line :1*/ len(tWO2olI)-1]
	}
	return tWO2olI
}


func ns5_ZXVgtnn(fIadEXIim6l byte) bool {
	return fIadEXIim6l == ' ' || fIadEXIim6l == '\t' || fIadEXIim6l == '\n' || fIadEXIim6l == '\r'
}



func vakjfEA2(ulv1Z061Va string) string {
	if eaAUaB7Z :=  /*line :1*/ bytealg.IndexByteString(ulv1Z061Va, '#'); eaAUaB7Z != -1 {
		return ulv1Z061Va[:eaAUaB7Z]
	}
	return ulv1Z061Va
}



func fF9mgT(tWO2olI string, iO3AXsvw func(fJGE8VL2 string) error) error {
	tWO2olI =  /*line :1*/ eEa5QIpda3t(tWO2olI)
	for  /*line :1*/ len(tWO2olI) > 0 {
		pXhKjvAt1 :=  /*line :1*/ bytealg.IndexByteString(tWO2olI, ' ')
		if pXhKjvAt1 == -1 {
			return  /*line :1*/ iO3AXsvw(tWO2olI)
		}
		if fJGE8VL2 :=  /*line :1*/ eEa5QIpda3t(tWO2olI[:pXhKjvAt1]);  /*line :1*/ len(fJGE8VL2) > 0 {
			if h_ljk48Bm :=  /*line :1*/ iO3AXsvw(fJGE8VL2); h_ljk48Bm != nil {
				return h_ljk48Bm
			}
		}
		tWO2olI =  /*line :1*/ eEa5QIpda3t(tWO2olI[pXhKjvAt1+1:])
	}
	return nil
}



func umZUnV4x(dRoFccaZ, oWGiddK string) bool {
	return  /*line :1*/ len(dRoFccaZ) >=  /*line :1*/ len(oWGiddK) && dRoFccaZ[ /*line :1*/ len(dRoFccaZ)- /*line :1*/ len(oWGiddK):] == oWGiddK
}



func jRYeKNE(dRoFccaZ, oWGiddK string) bool {
	return  /*line :1*/ len(dRoFccaZ) >=  /*line :1*/ len(oWGiddK) &&  /*line :1*/ hCjYM9U0fG(dRoFccaZ[ /*line :1*/ len(dRoFccaZ)- /*line :1*/ len(oWGiddK):], oWGiddK)
}


func t9qAz2H(dRoFccaZ, zjvciD8ArgP string) bool {
	return  /*line :1*/ len(dRoFccaZ) >=  /*line :1*/ len(zjvciD8ArgP) && dRoFccaZ[: /*line :1*/ len(zjvciD8ArgP)] == zjvciD8ArgP
}



func hCjYM9U0fG(dRoFccaZ, aeaqWzxJu string) bool {
	if  /*line :1*/ len(dRoFccaZ) !=  /*line :1*/ len(aeaqWzxJu) {
		return false
	}
	for eaAUaB7Z := 0; eaAUaB7Z <  /*line :1*/ len(dRoFccaZ); eaAUaB7Z++ {
		if  /*line :1*/ erPIFF(dRoFccaZ[eaAUaB7Z]) !=  /*line :1*/ erPIFF(aeaqWzxJu[eaAUaB7Z]) {
			return false
		}
	}
	return true
}
