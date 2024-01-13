//line :1
package fRAfQd_

import (
	errors "iAHoxjmM"
	syscall "bUKeamF"
)


func t6I6rVJl1Q5() {
}

func bavnEgG1Sl(w12CVdh9 string) (uintptr, error) {
	laArf75Y6GN, xuMLYrB :=  /*line :1*/ syscall.Ek485AtskVVo(w12CVdh9, syscall.O_RDONLY, 0)
	if xuMLYrB != nil {

		if xuMLYrB == syscall.ERROR_PATH_NOT_FOUND {
			xuMLYrB = syscall.ENOENT
		}
		return 0, xuMLYrB
	}
	return  /*line :1*/ uintptr(laArf75Y6GN), nil
}

func ozmugx0ism(laArf75Y6GN uintptr, qtO6L35R9b []byte) (int, error) {
	return  /*line :1*/ syscall.XJ0LdRoyxl_l( /*line :1*/ syscall.SRlvVjrYa(laArf75Y6GN), qtO6L35R9b)
}

func hGSAcx81(laArf75Y6GN uintptr) {
	 /*line :1*/ syscall.Kv0fHx9Ttx0( /*line :1*/ syscall.SRlvVjrYa(laArf75Y6GN))
}

func fUl4zA(laArf75Y6GN uintptr, qtO6L35R9b []byte, mP6vzfnl int) error {
	jZ_Xnge5hdg := seekStart
	if mP6vzfnl < 0 {
		jZ_Xnge5hdg = seekEnd
	}
	if _, xuMLYrB :=  /*line :1*/ syscall.O4FqtrPqbHRY( /*line :1*/ syscall.SRlvVjrYa(laArf75Y6GN),  /*line :1*/ int64(mP6vzfnl), jZ_Xnge5hdg); xuMLYrB != nil {
		return xuMLYrB
	}
	for  /*line :1*/ len(qtO6L35R9b) > 0 {
		fYy09z7m, xuMLYrB :=  /*line :1*/ syscall.XJ0LdRoyxl_l( /*line :1*/ syscall.SRlvVjrYa(laArf75Y6GN), qtO6L35R9b)
		if fYy09z7m <= 0 {
			if xuMLYrB == nil {
				return  /*line :1*/ errors.Su6g6hRoi9X("short read")
			}
			return xuMLYrB
		}
		qtO6L35R9b = qtO6L35R9b[fYy09z7m:]
	}
	return nil
}
