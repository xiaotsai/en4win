//line :1
package gF9mZs2

import (
	errors "iAHoxjmM"
	"internal/bytealg"
	fs "y1BamVjnXsaa"
	netip "iPdCam_KQOXj"
	sync "sync"
	time "fRAfQd_"
)

const cacheMaxAge = 5 * time.Second

func j6B8mQR8(qxwkC3VxG0 string) string {
	kQ8_UEhxU, h_ljk48Bm :=  /*line :1*/ netip.Y1uXU4Sx(qxwkC3VxG0)
	if h_ljk48Bm != nil {
		return ""
	}
	return  /*line :1*/ kQ8_UEhxU.String()
}

type wMkwJS1c struct {
	bTgOzgJ0igqk	[]string
	aoB0VIK	string
}


var eFm32AJm struct {
	sync.DIRWe11KYlYa

	
	
	
	
	bd8AWWEaE	map[string]wMkwJS1c

	
	
	
	zraHICqk	map[string][]string

	qYX8JxKAJ	time.G4KDkI
	ds2zFR	string
	etzFDH	time.G4KDkI
	dkHGOgemL	int64
}

func beBLGXSrUJP() {
	t_5qtVaQY :=  /*line :1*/ time.Pc_35oTY()
	oIS6m5SB7lKZ := gPw7YvxJUYD

	if  /*line :1*/ t_5qtVaQY.Before(eFm32AJm.qYX8JxKAJ) && eFm32AJm.ds2zFR == oIS6m5SB7lKZ &&  /*line :1*/ len(eFm32AJm.bd8AWWEaE) > 0 {
		return
	}
	i81SGZwwa, rz5yRs4IdH, h_ljk48Bm :=  /*line :1*/ o6GKG3liN(oIS6m5SB7lKZ)
	if h_ljk48Bm == nil && eFm32AJm.ds2zFR == oIS6m5SB7lKZ &&  /*line :1*/ eFm32AJm.etzFDH.Equal(i81SGZwwa) && eFm32AJm.dkHGOgemL == rz5yRs4IdH {
		eFm32AJm.qYX8JxKAJ =  /*line :1*/ t_5qtVaQY.Add(cacheMaxAge)
		return
	}

	g3PVOLx :=  /*line :1*/ make(map[string]wMkwJS1c)
	fWMsMk6jt9 :=  /*line :1*/ make(map[string][]string)

	cZsWMjIJ6, h_ljk48Bm :=  /*line :1*/ _qDAPa7V(oIS6m5SB7lKZ)
	if h_ljk48Bm != nil {
		if ! /*line :1*/ errors.COBastO_C(h_ljk48Bm, fs.QhJWcuT5mSD) && ! /*line :1*/ errors.COBastO_C(h_ljk48Bm, fs.E8Hgdf) {
			return
		}
	}

	if cZsWMjIJ6 != nil {
		defer  /*line :1*/ cZsWMjIJ6.tFIv2Ppbx0H()
		for ulv1Z061Va, d30HIwxht :=  /*line :1*/ cZsWMjIJ6.dM_MFaU7Ms(); d30HIwxht; ulv1Z061Va, d30HIwxht =  /*line :1*/ cZsWMjIJ6.dM_MFaU7Ms() {
			if eaAUaB7Z :=  /*line :1*/ bytealg.IndexByteString(ulv1Z061Va, '#'); eaAUaB7Z >= 0 {

				ulv1Z061Va = ulv1Z061Va[0:eaAUaB7Z]
			}
			t5q9DVOD9 :=  /*line :1*/ nhtSVnZ(ulv1Z061Va)
			if  /*line :1*/ len(t5q9DVOD9) < 2 {
				continue
			}
			qxwkC3VxG0 :=  /*line :1*/ j6B8mQR8(t5q9DVOD9[0])
			if qxwkC3VxG0 == "" {
				continue
			}

			var eLUDzideg string
			for eaAUaB7Z := 1; eaAUaB7Z <  /*line :1*/ len(t5q9DVOD9); eaAUaB7Z++ {
				tahgc783Bc :=  /*line :1*/ ztv3vbC8VS(t5q9DVOD9[eaAUaB7Z])
				cBprzKT8 := [] /*line :1*/ byte(t5q9DVOD9[eaAUaB7Z])
				 /*line :1*/ kuRb7e(cBprzKT8)
				qFUmBERYgmKz :=  /*line :1*/ ztv3vbC8VS( /*line :1*/ string(cBprzKT8))

				if eaAUaB7Z == 1 {
					eLUDzideg = qFUmBERYgmKz
				}

				fWMsMk6jt9[qxwkC3VxG0] =  /*line :1*/ append(fWMsMk6jt9[qxwkC3VxG0], tahgc783Bc)

				if ljsCSk, d30HIwxht := g3PVOLx[qFUmBERYgmKz]; d30HIwxht {
					g3PVOLx[qFUmBERYgmKz] = wMkwJS1c{
						bTgOzgJ0igqk:	 /*line :1*/ append(ljsCSk.bTgOzgJ0igqk, qxwkC3VxG0),
						aoB0VIK:	ljsCSk.aoB0VIK,
					}
					continue
				}

				g3PVOLx[qFUmBERYgmKz] = wMkwJS1c{
					bTgOzgJ0igqk:	[]string{qxwkC3VxG0},
					aoB0VIK:	eLUDzideg,
				}
			}
		}
	}

	eFm32AJm.qYX8JxKAJ =  /*line :1*/ t_5qtVaQY.Add(cacheMaxAge)
	eFm32AJm.ds2zFR = oIS6m5SB7lKZ
	eFm32AJm.bd8AWWEaE = g3PVOLx
	eFm32AJm.zraHICqk = fWMsMk6jt9
	eFm32AJm.etzFDH = i81SGZwwa
	eFm32AJm.dkHGOgemL = rz5yRs4IdH
}


func y_FPhQLXqT(jKxI_T1CK_p string) ([]string, string) {
	 /*line :1*/ eFm32AJm.Lock()
	defer  /*line :1*/ eFm32AJm.Unlock()
	 /*line :1*/ beBLGXSrUJP()
	if  /*line :1*/ len(eFm32AJm.bd8AWWEaE) != 0 {
		if  /*line :1*/ wr9w1HG(jKxI_T1CK_p) {
			nmNeXEUY := [] /*line :1*/ byte(jKxI_T1CK_p)
			 /*line :1*/ kuRb7e(nmNeXEUY)
			jKxI_T1CK_p =  /*line :1*/ string(nmNeXEUY)
		}
		if wMkwJS1c, d30HIwxht := eFm32AJm.bd8AWWEaE[ /*line :1*/ ztv3vbC8VS(jKxI_T1CK_p)]; d30HIwxht {
			jSgolitR :=  /*line :1*/ make([]string,  /*line :1*/ len(wMkwJS1c.bTgOzgJ0igqk))
			 /*line :1*/ copy(jSgolitR, wMkwJS1c.bTgOzgJ0igqk)
			return jSgolitR, wMkwJS1c.aoB0VIK
		}
	}
	return nil, ""
}


func cb0AaP(qxwkC3VxG0 string) []string {
	 /*line :1*/ eFm32AJm.Lock()
	defer  /*line :1*/ eFm32AJm.Unlock()
	 /*line :1*/ beBLGXSrUJP()
	qxwkC3VxG0 =  /*line :1*/ j6B8mQR8(qxwkC3VxG0)
	if qxwkC3VxG0 == "" {
		return nil
	}
	if  /*line :1*/ len(eFm32AJm.zraHICqk) != 0 {
		if eFm32AJm, d30HIwxht := eFm32AJm.zraHICqk[qxwkC3VxG0]; d30HIwxht {
			uaY1qs :=  /*line :1*/ make([]string,  /*line :1*/ len(eFm32AJm))
			 /*line :1*/ copy(uaY1qs, eFm32AJm)
			return uaY1qs
		}
	}
	return nil
}
