//line :1
package hPMrClpC

import (
	windows "LdptURlN"
	syscall "bUKeamF"
)

func bFLwMmQoa() (jGBs03 string, ugqb2IW error) {
	
	const format = windows.ComputerNamePhysicalDnsHostname

	zHDjCTHI :=  /*line :1*/ uint32(64)
	for {
		nR2aPlreQFZA :=  /*line :1*/ make([]uint16, zHDjCTHI)
		ugqb2IW :=  /*line :1*/ windows.YyY7M_287(format, &nR2aPlreQFZA[0], &zHDjCTHI)
		if ugqb2IW == nil {
			return  /*line :1*/ syscall.AODVXp8NM3sd(nR2aPlreQFZA[:zHDjCTHI]), nil
		}
		if ugqb2IW != syscall.ERROR_MORE_DATA {
			return "",  /*line :1*/ BTaHHl("ComputerNameEx", ugqb2IW)
		}

		if zHDjCTHI <=  /*line :1*/ uint32( /*line :1*/ len(nR2aPlreQFZA)) {
			return "",  /*line :1*/ BTaHHl("ComputerNameEx", ugqb2IW)
		}
	}
}
