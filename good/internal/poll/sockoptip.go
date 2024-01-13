//line :1
//go:build unix || windows

package MjXXMR

import syscall "bUKeamF"

func (s_ZR_8 *ENfmDmMaozH) SetsockoptIPMreq(fs24M1GNo, dEPdiT1wKVtx int, hPyySyPDI *syscall.W1r5YRtipLa) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.F2RtoD(s_ZR_8.X8OEsVYSby, fs24M1GNo, dEPdiT1wKVtx, hPyySyPDI)
}

func (s_ZR_8 *ENfmDmMaozH) SetsockoptIPv6Mreq(fs24M1GNo, dEPdiT1wKVtx int, hPyySyPDI *syscall.C5ybc0) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.G8xAk43n(s_ZR_8.X8OEsVYSby, fs24M1GNo, dEPdiT1wKVtx, hPyySyPDI)
}
