//line :1
//go:build unix || windows

package MjXXMR

import syscall "bUKeamF"

func (s_ZR_8 *ENfmDmMaozH) SetsockoptInt(fs24M1GNo, dEPdiT1wKVtx, _MHO0E int) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.ABpNrEFh9ZGb(s_ZR_8.X8OEsVYSby, fs24M1GNo, dEPdiT1wKVtx, _MHO0E)
}

func (s_ZR_8 *ENfmDmMaozH) SetsockoptInet4Addr(fs24M1GNo, dEPdiT1wKVtx int, _MHO0E [4]byte) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.Fnnrp31xH8y(s_ZR_8.X8OEsVYSby, fs24M1GNo, dEPdiT1wKVtx, _MHO0E)
}

func (s_ZR_8 *ENfmDmMaozH) SetsockoptLinger(fs24M1GNo, dEPdiT1wKVtx int, gwLmzX *syscall.KW8FJX7) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.JPiqBG(s_ZR_8.X8OEsVYSby, fs24M1GNo, dEPdiT1wKVtx, gwLmzX)
}

func (s_ZR_8 *ENfmDmMaozH) GetsockoptInt(fs24M1GNo, dEPdiT1wKVtx int) (int, error) {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return -1, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.QWe7U4(s_ZR_8.X8OEsVYSby, fs24M1GNo, dEPdiT1wKVtx)
}
