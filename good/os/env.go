//line :1
package hPMrClpC

import (
	testlog "X0MraxO81Me"
	syscall "bUKeamF"
)



func Iu6ylCk(gkORch7na string, lMz9DIyYiik func(string) string) string {
	var dDkChyAkT []byte

	lWfaal := 0
	for lpNM41IF := 0; lpNM41IF <  /*line :1*/ len(gkORch7na); lpNM41IF++ {
		if gkORch7na[lpNM41IF] == '$' && lpNM41IF+1 <  /*line :1*/ len(gkORch7na) {
			if dDkChyAkT == nil {
				dDkChyAkT =  /*line :1*/ make([]byte, 0, 2* /*line :1*/ len(gkORch7na))
			}
			dDkChyAkT =  /*line :1*/ append(dDkChyAkT, gkORch7na[lWfaal:lpNM41IF]...)
			jGBs03, pgkDcUH :=  /*line :1*/ c3SxjD4sAgF(gkORch7na[lpNM41IF+1:])
			if jGBs03 == "" && pgkDcUH > 0 {

			} else if jGBs03 == "" {

				dDkChyAkT =  /*line :1*/ append(dDkChyAkT, gkORch7na[lpNM41IF])
			} else {
				dDkChyAkT =  /*line :1*/ append(dDkChyAkT,  /*line :1*/ lMz9DIyYiik(jGBs03)...)
			}
			lpNM41IF += pgkDcUH
			lWfaal = lpNM41IF + 1
		}
	}
	if dDkChyAkT == nil {
		return gkORch7na
	}
	return  /*line :1*/ string(dDkChyAkT) + gkORch7na[lWfaal:]
}




func TZ6208i1E7Ob(gkORch7na string) string {
	return  /*line :1*/ Iu6ylCk(gkORch7na, LDjFPRBEz)
}



func gm9aPQhf(ivfnLGc uint8) bool {
	switch ivfnLGc {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}


func nqB7R4(ivfnLGc uint8) bool {
	return ivfnLGc == '_' || '0' <= ivfnLGc && ivfnLGc <= '9' || 'a' <= ivfnLGc && ivfnLGc <= 'z' || 'A' <= ivfnLGc && ivfnLGc <= 'Z'
}




func c3SxjD4sAgF(gkORch7na string) (string, int) {
	switch {
	case gkORch7na[0] == '{':
		if  /*line :1*/ len(gkORch7na) > 2 &&  /*line :1*/ gm9aPQhf(gkORch7na[1]) && gkORch7na[2] == '}' {
			return gkORch7na[1:2], 3
		}

		for lWfaal := 1; lWfaal <  /*line :1*/ len(gkORch7na); lWfaal++ {
			if gkORch7na[lWfaal] == '}' {
				if lWfaal == 1 {
					return "", 2
				}
				return gkORch7na[1:lWfaal], lWfaal + 1
			}
		}
		return "", 1
	case  /*line :1*/ gm9aPQhf(gkORch7na[0]):
		return gkORch7na[0:1], 1
	}
	
	var lWfaal int
	for lWfaal = 0; lWfaal <  /*line :1*/ len(gkORch7na) &&  /*line :1*/ nqB7R4(gkORch7na[lWfaal]); lWfaal++ {
	}
	return gkORch7na[:lWfaal], lWfaal
}




func LDjFPRBEz(uZmEf_td4 string) string {
	 /*line :1*/ testlog.HUDK1vRBN(uZmEf_td4)
	rB6hatF3c, _ :=  /*line :1*/ syscall.ITYf1B(uZmEf_td4)
	return rB6hatF3c
}






func EPSG_GMwS(uZmEf_td4 string) (string, bool) {
	 /*line :1*/ testlog.HUDK1vRBN(uZmEf_td4)
	return  /*line :1*/ syscall.ITYf1B(uZmEf_td4)
}



func JECjRbaLhMH(uZmEf_td4, pwL2nLziJPz string) error {
	ugqb2IW :=  /*line :1*/ syscall.BNcaDu(uZmEf_td4, pwL2nLziJPz)
	if ugqb2IW != nil {
		return  /*line :1*/ BTaHHl("setenv", ugqb2IW)
	}
	return nil
}


func U9mGtQe(uZmEf_td4 string) error {
	return  /*line :1*/ syscall.OuD_oLNMM(uZmEf_td4)
}


func Ds5Nvp() {
	 /*line :1*/ syscall.J1XgGX_nmS4b()
}



func U9gHttio() []string {
	return  /*line :1*/ syscall.FDhTFlUY()
}
