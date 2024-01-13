//line :1
package gF9mZs2

import (
	os "hPMrClpC"
	syscall "bUKeamF"
)

func faet4q(dRoFccaZ syscall.SRlvVjrYa, nqsgnnaf, dr86pU int, q_vcfH bool) error {
	if nqsgnnaf == syscall.AF_INET6 && dr86pU != syscall.SOCK_RAW {

		 /*line :1*/ syscall.ABpNrEFh9ZGb(dRoFccaZ, syscall.IPPROTO_IPV6, syscall.IPV6_V6ONLY,  /*line :1*/ j7bhnpSMpen(q_vcfH))
	}
	if (dr86pU == syscall.SOCK_DGRAM || dr86pU == syscall.SOCK_RAW) && nqsgnnaf != syscall.AF_UNIX && nqsgnnaf != syscall.AF_INET6 {

		return  /*line :1*/ os.BTaHHl("setsockopt",  /*line :1*/ syscall.ABpNrEFh9ZGb(dRoFccaZ, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1))
	}
	return nil
}

func atNQ_vLoo(dRoFccaZ syscall.SRlvVjrYa) error {

	return nil
}

func ojx7Y3NFA(dRoFccaZ syscall.SRlvVjrYa) error {

	return  /*line :1*/ os.BTaHHl("setsockopt",  /*line :1*/ syscall.ABpNrEFh9ZGb(dRoFccaZ, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1))
}
