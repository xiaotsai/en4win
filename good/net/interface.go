//line :1
package gF9mZs2

import (
	errors "iAHoxjmM"
	itoa "JkjxVS"
	sync "sync"
	time "fRAfQd_"
)

var (
	jX28_o	=  /*line :1*/ errors.Su6g6hRoi9X("invalid network interface")
	f3gbmlYa	=  /*line :1*/ errors.Su6g6hRoi9X("invalid network interface index")
	omaDNiAllSZ	=  /*line :1*/ errors.Su6g6hRoi9X("invalid network interface name")
	fa1xDtq0	=  /*line :1*/ errors.Su6g6hRoi9X("no such network interface")
	k7ow5kuh	=  /*line :1*/ errors.Su6g6hRoi9X("no such multicast network interface")
)




type U_Uc9la struct {
	OJz1cH7Vb	int	
	Vlcyn9WPA2	int	
	IykpW1qlQjWm	string	
	AgZ5lCg	CGTWv0eqN	
	LFFOqgAI	SpLyr9	
}

type SpLyr9 uint

const (
	FlagUp	SpLyr9	= 1 << iota	
	FlagBroadcast			
	FlagLoopback			
	FlagPointToPoint			
	FlagMulticast			
	FlagRunning			
)

var zztOPx_AgI8Q = []string{
	"up",
	"broadcast",
	"loopback",
	"pointtopoint",
	"multicast",
	"running",
}

func (t5q9DVOD9 SpLyr9) String() string {
	dRoFccaZ := ""
	for eaAUaB7Z, tahgc783Bc := range zztOPx_AgI8Q {
		if t5q9DVOD9&(1<< /*line :1*/ uint(eaAUaB7Z)) != 0 {
			if dRoFccaZ != "" {
				dRoFccaZ += "|"
			}
			dRoFccaZ += tahgc783Bc
		}
	}
	if dRoFccaZ == "" {
		dRoFccaZ = "0"
	}
	return dRoFccaZ
}



func (dvFmDYmQz *U_Uc9la) Addrs() ([]EqbMXsaU1lE, error) {
	if dvFmDYmQz == nil {
		return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: jX28_o}
	}
	zDd8fqQGg, h_ljk48Bm :=  /*line :1*/ uovWlYao2O(dvFmDYmQz)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return zDd8fqQGg, h_ljk48Bm
}



func (dvFmDYmQz *U_Uc9la) MulticastAddrs() ([]EqbMXsaU1lE, error) {
	if dvFmDYmQz == nil {
		return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: jX28_o}
	}
	zDd8fqQGg, h_ljk48Bm :=  /*line :1*/ gsdWgpJ3dT(dvFmDYmQz)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return zDd8fqQGg, h_ljk48Bm
}


func WNSbt0MQgc() ([]U_Uc9la, error) {
	uotBHWl8c47, h_ljk48Bm :=  /*line :1*/ s0JmDXAcmu(0)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	if  /*line :1*/ len(uotBHWl8c47) != 0 {
		 /*line :1*/ x4vxFT.yTxNdMX(uotBHWl8c47, false)
	}
	return uotBHWl8c47, nil
}






func B3nA0hb() ([]EqbMXsaU1lE, error) {
	zDd8fqQGg, h_ljk48Bm :=  /*line :1*/ uovWlYao2O(nil)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return zDd8fqQGg, h_ljk48Bm
}






func D85CaUT3H9_(eJTgFgJkWaQ int) (*U_Uc9la, error) {
	if eJTgFgJkWaQ <= 0 {
		return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: f3gbmlYa}
	}
	uotBHWl8c47, h_ljk48Bm :=  /*line :1*/ s0JmDXAcmu(eJTgFgJkWaQ)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	dvFmDYmQz, h_ljk48Bm :=  /*line :1*/ qFgPBDbLoP5(uotBHWl8c47, eJTgFgJkWaQ)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return dvFmDYmQz, h_ljk48Bm
}

func qFgPBDbLoP5(uotBHWl8c47 []U_Uc9la, eJTgFgJkWaQ int) (*U_Uc9la, error) {
	for _, dvFmDYmQz := range uotBHWl8c47 {
		if eJTgFgJkWaQ == dvFmDYmQz.OJz1cH7Vb {
			return &dvFmDYmQz, nil
		}
	}
	return nil, fa1xDtq0
}


func Ed4kqnOj_(tahgc783Bc string) (*U_Uc9la, error) {
	if tahgc783Bc == "" {
		return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: omaDNiAllSZ}
	}
	uotBHWl8c47, h_ljk48Bm :=  /*line :1*/ s0JmDXAcmu(0)
	if h_ljk48Bm != nil {
		return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: h_ljk48Bm}
	}
	if  /*line :1*/ len(uotBHWl8c47) != 0 {
		 /*line :1*/ x4vxFT.yTxNdMX(uotBHWl8c47, false)
	}
	for _, dvFmDYmQz := range uotBHWl8c47 {
		if tahgc783Bc == dvFmDYmQz.IykpW1qlQjWm {
			return &dvFmDYmQz, nil
		}
	}
	return nil, &ZOYGdck{SMNyZk_q: "route", CiQ5sBtmrVnf: "ip+net", F7g5pQacM05: nil, JkQ9XFJbp: nil, XkWH4CYwNmvD: fa1xDtq0}
}







type wTNtcaqavH struct {
	sync.V1sRgjwAglJt		
	kuL2JxoeV	time.G4KDkI	
	jUyfzfFUHbSy	map[string]int	
	kMkF8cw	map[int]string	
}

var x4vxFT = wTNtcaqavH{
	jUyfzfFUHbSy:	 /*line :1*/ make(map[string]int),
	kMkF8cw:	 /*line :1*/ make(map[int]string),
}




func (aoNtmlVM4g *wTNtcaqavH) yTxNdMX(uotBHWl8c47 []U_Uc9la, yl116H bool) (qSyADON bool) {
	 /*line :1*/ aoNtmlVM4g.Lock()
	defer  /*line :1*/ aoNtmlVM4g.Unlock()
	t_5qtVaQY :=  /*line :1*/ time.Pc_35oTY()
	if !yl116H &&  /*line :1*/ aoNtmlVM4g.kuL2JxoeV.After( /*line :1*/ t_5qtVaQY.Add(-60*time.Second)) {
		return false
	}
	aoNtmlVM4g.kuL2JxoeV = t_5qtVaQY
	if  /*line :1*/ len(uotBHWl8c47) == 0 {
		var h_ljk48Bm error
		if uotBHWl8c47, h_ljk48Bm =  /*line :1*/ s0JmDXAcmu(0); h_ljk48Bm != nil {
			return false
		}
	}
	aoNtmlVM4g.jUyfzfFUHbSy =  /*line :1*/ make(map[string]int,  /*line :1*/ len(uotBHWl8c47))
	aoNtmlVM4g.kMkF8cw =  /*line :1*/ make(map[int]string,  /*line :1*/ len(uotBHWl8c47))
	for _, dvFmDYmQz := range uotBHWl8c47 {
		aoNtmlVM4g.jUyfzfFUHbSy[dvFmDYmQz.IykpW1qlQjWm] = dvFmDYmQz.OJz1cH7Vb
		if _, d30HIwxht := aoNtmlVM4g.kMkF8cw[dvFmDYmQz.OJz1cH7Vb]; !d30HIwxht {
			aoNtmlVM4g.kMkF8cw[dvFmDYmQz.OJz1cH7Vb] = dvFmDYmQz.IykpW1qlQjWm
		}
	}
	return true
}

func (aoNtmlVM4g *wTNtcaqavH) tahgc783Bc(eJTgFgJkWaQ int) string {
	if eJTgFgJkWaQ == 0 {
		return ""
	}
	qSyADON :=  /*line :1*/ x4vxFT.yTxNdMX(nil, false)
	 /*line :1*/ x4vxFT.RLock()
	tahgc783Bc, d30HIwxht := x4vxFT.kMkF8cw[eJTgFgJkWaQ]
	 /*line :1*/ x4vxFT.RUnlock()
	if !d30HIwxht && !qSyADON {
		 /*line :1*/ x4vxFT.yTxNdMX(nil, true)
		 /*line :1*/ x4vxFT.RLock()
		tahgc783Bc, d30HIwxht = x4vxFT.kMkF8cw[eJTgFgJkWaQ]
		 /*line :1*/ x4vxFT.RUnlock()
	}
	if !d30HIwxht {
		tahgc783Bc =  /*line :1*/ itoa.TDADPeM7Kft9( /*line :1*/ uint(eJTgFgJkWaQ))
	}
	return tahgc783Bc
}

func (aoNtmlVM4g *wTNtcaqavH) eJTgFgJkWaQ(tahgc783Bc string) int {
	if tahgc783Bc == "" {
		return 0
	}
	qSyADON :=  /*line :1*/ x4vxFT.yTxNdMX(nil, false)
	 /*line :1*/ x4vxFT.RLock()
	eJTgFgJkWaQ, d30HIwxht := x4vxFT.jUyfzfFUHbSy[tahgc783Bc]
	 /*line :1*/ x4vxFT.RUnlock()
	if !d30HIwxht && !qSyADON {
		 /*line :1*/ x4vxFT.yTxNdMX(nil, true)
		 /*line :1*/ x4vxFT.RLock()
		eJTgFgJkWaQ, d30HIwxht = x4vxFT.jUyfzfFUHbSy[tahgc783Bc]
		 /*line :1*/ x4vxFT.RUnlock()
	}
	if !d30HIwxht {
		eJTgFgJkWaQ, _, _ =  /*line :1*/ bEihJjrbz(tahgc783Bc)
	}
	return eJTgFgJkWaQ
}
