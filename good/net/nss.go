//line :1
package gF9mZs2

import (
	errors "iAHoxjmM"
	"internal/bytealg"
	os "hPMrClpC"
	sync "sync"
	time "fRAfQd_"
)

const (
	nssConfigPath = "/etc/nsswitch.conf"
)

var b3eHlYLoe eUy_ArVYl

type eUy_ArVYl struct {
	fUGJTQj	sync.LhBfwn6wa1x	

	
	
	qYwal9qjVJ	chan struct{}	
	pHnZHOf6FHQ	time.G4KDkI	

	aaaSRJbG	sync.DIRWe11KYlYa	
	e1tVpdA9UcyZ	*rrJLWHMc6tN
}

func kKNl81QAoLO() *rrJLWHMc6tN {
	 /*line :1*/ b3eHlYLoe.zRkY41bMBy()
	 /*line :1*/ b3eHlYLoe.aaaSRJbG.Lock()
	q6Glr4X1u := b3eHlYLoe.e1tVpdA9UcyZ
	 /*line :1*/ b3eHlYLoe.aaaSRJbG.Unlock()
	return q6Glr4X1u
}


func (q6Glr4X1u *eUy_ArVYl) init() {
	q6Glr4X1u.e1tVpdA9UcyZ =  /*line :1*/ nsS7D91("/etc/nsswitch.conf")
	q6Glr4X1u.pHnZHOf6FHQ =  /*line :1*/ time.Pc_35oTY()
	q6Glr4X1u.qYwal9qjVJ =  /*line :1*/ make(chan struct{}, 1)
}


func (q6Glr4X1u *eUy_ArVYl) zRkY41bMBy() {
	 /*line :1*/ q6Glr4X1u.fUGJTQj.Do(q6Glr4X1u.init)

	if ! /*line :1*/ q6Glr4X1u.xfyLaBaWu9e4() {
		return
	}
	defer  /*line :1*/ q6Glr4X1u.wDHVRjOehJ()

	t_5qtVaQY :=  /*line :1*/ time.Pc_35oTY()
	if  /*line :1*/ q6Glr4X1u.pHnZHOf6FHQ.After( /*line :1*/ t_5qtVaQY.Add(-5 * time.Second)) {
		return
	}
	q6Glr4X1u.pHnZHOf6FHQ = t_5qtVaQY

	var i81SGZwwa time.G4KDkI
	if a6qU1P1, h_ljk48Bm :=  /*line :1*/ os.ZYa3wuv(nssConfigPath); h_ljk48Bm == nil {
		i81SGZwwa =  /*line :1*/ a6qU1P1.ModTime()
	}
	if  /*line :1*/ i81SGZwwa.Equal(q6Glr4X1u.e1tVpdA9UcyZ.bYvH7L3aZr) {
		return
	}

	rrJLWHMc6tN :=  /*line :1*/ nsS7D91(nssConfigPath)
	 /*line :1*/ q6Glr4X1u.aaaSRJbG.Lock()
	q6Glr4X1u.e1tVpdA9UcyZ = rrJLWHMc6tN
	 /*line :1*/ q6Glr4X1u.aaaSRJbG.Unlock()
}

func (q6Glr4X1u *eUy_ArVYl) lPyKqxPm_9wt() {
	q6Glr4X1u.qYwal9qjVJ <- struct{}{}
}

func (q6Glr4X1u *eUy_ArVYl) xfyLaBaWu9e4() bool {
	select {
	case q6Glr4X1u.qYwal9qjVJ <- struct{}{}:
		return true
	default:
		return false
	}
}

func (q6Glr4X1u *eUy_ArVYl) wDHVRjOehJ() {
	<-q6Glr4X1u.qYwal9qjVJ
}


type rrJLWHMc6tN struct {
	bYvH7L3aZr	time.G4KDkI	
	czDPXGoFxc4	error	
	sHk8loC5lC	map[string][]ndB3ShuNG	
}

type ndB3ShuNG struct {
	onlkrKzJ	string	
	sXaVaw	[]ahDZJDKT1NiG
}



func (dRoFccaZ ndB3ShuNG) gs8yLpU5_R2F() bool {
	for eaAUaB7Z, j3OEN0PqGs := range dRoFccaZ.sXaVaw {
		if ! /*line :1*/ j3OEN0PqGs.t8nTQpBuSJW(eaAUaB7Z ==  /*line :1*/ len(dRoFccaZ.sXaVaw)-1) {
			return false
		}
	}
	return true
}



type ahDZJDKT1NiG struct {
	zoKc1XbLhn	bool	
	i20SaCG	string	
	dJUTnYmfo5z	string	
}




func (hl8w2lFX ahDZJDKT1NiG) t8nTQpBuSJW(vMwVxD_U bool) bool {
	if hl8w2lFX.zoKc1XbLhn {
		return false
	}
	var dtDgRGlZd69p string
	switch hl8w2lFX.i20SaCG {
	case "success":
		dtDgRGlZd69p = "return"
	case "notfound", "unavail", "tryagain":
		dtDgRGlZd69p = "continue"
	default:

		return false
	}
	if vMwVxD_U && hl8w2lFX.dJUTnYmfo5z == "return" {
		return true
	}
	return hl8w2lFX.dJUTnYmfo5z == dtDgRGlZd69p
}

func nsS7D91(cZsWMjIJ6 string) *rrJLWHMc6tN {
	t5q9DVOD9, h_ljk48Bm :=  /*line :1*/ _qDAPa7V(cZsWMjIJ6)
	if h_ljk48Bm != nil {
		return &rrJLWHMc6tN{czDPXGoFxc4: h_ljk48Bm}
	}
	defer  /*line :1*/ t5q9DVOD9.tFIv2Ppbx0H()
	i81SGZwwa, _, h_ljk48Bm :=  /*line :1*/ t5q9DVOD9.o6GKG3liN()
	if h_ljk48Bm != nil {
		return &rrJLWHMc6tN{czDPXGoFxc4: h_ljk48Bm}
	}

	q6Glr4X1u :=  /*line :1*/ h42zLwNKMg(t5q9DVOD9)
	q6Glr4X1u.bYvH7L3aZr = i81SGZwwa
	return q6Glr4X1u
}

func h42zLwNKMg(t5q9DVOD9 *cZsWMjIJ6) *rrJLWHMc6tN {
	q6Glr4X1u :=  /*line :1*/ new(rrJLWHMc6tN)
	for ulv1Z061Va, d30HIwxht :=  /*line :1*/ t5q9DVOD9.dM_MFaU7Ms(); d30HIwxht; ulv1Z061Va, d30HIwxht =  /*line :1*/ t5q9DVOD9.dM_MFaU7Ms() {
		ulv1Z061Va =  /*line :1*/ eEa5QIpda3t( /*line :1*/ vakjfEA2(ulv1Z061Va))
		if  /*line :1*/ len(ulv1Z061Va) == 0 {
			continue
		}
		e1fd_ifwI :=  /*line :1*/ bytealg.IndexByteString(ulv1Z061Va, ':')
		if e1fd_ifwI == -1 {
			q6Glr4X1u.czDPXGoFxc4 =  /*line :1*/ errors.Su6g6hRoi9X("no colon on line")
			return q6Glr4X1u
		}
		bIbYna3paH :=  /*line :1*/ eEa5QIpda3t(ulv1Z061Va[:e1fd_ifwI])
		qS4zcqNdLbr := ulv1Z061Va[e1fd_ifwI+1:]
		for {
			qS4zcqNdLbr =  /*line :1*/ eEa5QIpda3t(qS4zcqNdLbr)
			if  /*line :1*/ len(qS4zcqNdLbr) == 0 {
				break
			}
			pXhKjvAt1 :=  /*line :1*/ bytealg.IndexByteString(qS4zcqNdLbr, ' ')
			var i0qXWLS string
			if pXhKjvAt1 == -1 {
				i0qXWLS = qS4zcqNdLbr
				qS4zcqNdLbr = ""
			} else {
				i0qXWLS = qS4zcqNdLbr[:pXhKjvAt1]
				qS4zcqNdLbr =  /*line :1*/ eEa5QIpda3t(qS4zcqNdLbr[pXhKjvAt1+1:])
			}
			var dS2G0FG4pom []ahDZJDKT1NiG

			if  /*line :1*/ len(qS4zcqNdLbr) > 0 && qS4zcqNdLbr[0] == '[' {
				laV0cwV4GcX :=  /*line :1*/ bytealg.IndexByteString(qS4zcqNdLbr, ']')
				if laV0cwV4GcX == -1 {
					q6Glr4X1u.czDPXGoFxc4 =  /*line :1*/ errors.Su6g6hRoi9X("unclosed criterion bracket")
					return q6Glr4X1u
				}
				var h_ljk48Bm error
				dS2G0FG4pom, h_ljk48Bm =  /*line :1*/ hiaV3qN(qS4zcqNdLbr[1:laV0cwV4GcX])
				if h_ljk48Bm != nil {
					q6Glr4X1u.czDPXGoFxc4 =  /*line :1*/ errors.Su6g6hRoi9X("invalid criteria: " + qS4zcqNdLbr[1:laV0cwV4GcX])
					return q6Glr4X1u
				}
				qS4zcqNdLbr = qS4zcqNdLbr[laV0cwV4GcX+1:]
			}
			if q6Glr4X1u.sHk8loC5lC == nil {
				q6Glr4X1u.sHk8loC5lC =  /*line :1*/ make(map[string][]ndB3ShuNG)
			}
			q6Glr4X1u.sHk8loC5lC[bIbYna3paH] =  /*line :1*/ append(q6Glr4X1u.sHk8loC5lC[bIbYna3paH], ndB3ShuNG{
				onlkrKzJ:	i0qXWLS,
				sXaVaw:	dS2G0FG4pom,
			})
		}
	}
	return q6Glr4X1u
}


func hiaV3qN(tWO2olI string) (hl8w2lFX []ahDZJDKT1NiG, h_ljk48Bm error) {
	h_ljk48Bm =  /*line :1*/ fF9mgT(tWO2olI, func(t5q9DVOD9 string) error {
		drYMv0 := false
		if  /*line :1*/ len(t5q9DVOD9) > 0 && t5q9DVOD9[0] == '!' {
			drYMv0 = true
			t5q9DVOD9 = t5q9DVOD9[1:]
		}
		if  /*line :1*/ len(t5q9DVOD9) < 3 {
			return  /*line :1*/ errors.Su6g6hRoi9X("criterion too short")
		}
		nxZ6Vj :=  /*line :1*/ bytealg.IndexByteString(t5q9DVOD9, '=')
		if nxZ6Vj == -1 {
			return  /*line :1*/ errors.Su6g6hRoi9X("criterion lacks equal sign")
		}
		if  /*line :1*/ wr9w1HG(t5q9DVOD9) {
			dx9N2TMbN4V := [] /*line :1*/ byte(t5q9DVOD9)
			 /*line :1*/ kuRb7e(dx9N2TMbN4V)
			t5q9DVOD9 =  /*line :1*/ string(dx9N2TMbN4V)
		}
		hl8w2lFX =  /*line :1*/ append(hl8w2lFX, ahDZJDKT1NiG{
			zoKc1XbLhn:	drYMv0,
			i20SaCG:	t5q9DVOD9[:nxZ6Vj],
			dJUTnYmfo5z:	t5q9DVOD9[nxZ6Vj+1:],
		})
		return nil
	})
	return
}
