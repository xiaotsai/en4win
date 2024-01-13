//line :1
package hPMrClpC

import (
	syscall "bUKeamF"
)








func HbLbvmMfV(sJC41DS string, nLWcea IvL5u8pdn) error {

	nqzPkc8LPse1, ugqb2IW :=  /*line :1*/ ZYa3wuv(sJC41DS)
	if ugqb2IW == nil {
		if  /*line :1*/ nqzPkc8LPse1.IsDir() {
			return nil
		}
		return &StNL1lveD40f{Aeg62j1VQt: "mkdir", OmDEtY: sJC41DS, OIEifQ0XF: syscall.ENOTDIR}
	}

	lWfaal :=  /*line :1*/ len(sJC41DS)
	for lWfaal > 0 &&  /*line :1*/ GClWKECc(sJC41DS[lWfaal-1]) {
		lWfaal--
	}

	lpNM41IF := lWfaal
	for lpNM41IF > 0 && ! /*line :1*/ GClWKECc(sJC41DS[lpNM41IF-1]) {
		lpNM41IF--
	}

	if lpNM41IF > 1 {

		ugqb2IW =  /*line :1*/ HbLbvmMfV( /*line :1*/ twDUdA(sJC41DS[:lpNM41IF-1]), nLWcea)
		if ugqb2IW != nil {
			return ugqb2IW
		}
	}

	ugqb2IW =  /*line :1*/ OEmAGB_T1CQq(sJC41DS, nLWcea)
	if ugqb2IW != nil {

		nqzPkc8LPse1, kLN0m8 :=  /*line :1*/ Z0JXsY(sJC41DS)
		if kLN0m8 == nil &&  /*line :1*/ nqzPkc8LPse1.IsDir() {
			return nil
		}
		return ugqb2IW
	}
	return nil
}






func RNRFW7(sJC41DS string) error {
	return  /*line :1*/ a0oDQ8(sJC41DS)
}


func nXnrtsK(sJC41DS string) bool {
	if sJC41DS == "." {
		return true
	}
	if  /*line :1*/ len(sJC41DS) >= 2 && sJC41DS[ /*line :1*/ len(sJC41DS)-1] == '.' &&  /*line :1*/ GClWKECc(sJC41DS[ /*line :1*/ len(sJC41DS)-2]) {
		return true
	}
	return false
}
