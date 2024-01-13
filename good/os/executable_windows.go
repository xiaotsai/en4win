//line :1
package hPMrClpC

import (
	windows "LdptURlN"
	syscall "bUKeamF"
)

func pqaYRZnR(iVmSWxZA4 syscall.SRlvVjrYa) (string, error) {
	zHDjCTHI :=  /*line :1*/ uint32(1024)
	var dDkChyAkT []uint16
	for {
		dDkChyAkT =  /*line :1*/ make([]uint16, zHDjCTHI)
		dhvYus, ugqb2IW :=  /*line :1*/ windows.GCd4lr(iVmSWxZA4, &dDkChyAkT[0], zHDjCTHI)
		if ugqb2IW != nil {
			return "", ugqb2IW
		}
		if dhvYus < zHDjCTHI {
			break
		}

		zHDjCTHI += 1024
	}
	return  /*line :1*/ syscall.AODVXp8NM3sd(dDkChyAkT), nil
}

func zKZCvoO7() (string, error) {
	return  /*line :1*/ pqaYRZnR(0)
}
