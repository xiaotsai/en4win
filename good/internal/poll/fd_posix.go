//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package MjXXMR

import (
	io "xI9ai2djaJ2"
	syscall "bUKeamF"
)

func (s_ZR_8 *ENfmDmMaozH) cKC2EMAVUS2(mnkyfx int, dtU8tBiUGS4 error) error {
	if mnkyfx == 0 && dtU8tBiUGS4 == nil && s_ZR_8.ZslYVRp1qJ {
		return io.K5Sqco
	}
	return dtU8tBiUGS4
}

func (s_ZR_8 *ENfmDmMaozH) Shutdown(ldyh5X5E7EP int) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ syscall.H3rGVgc1(s_ZR_8.X8OEsVYSby, ldyh5X5E7EP)
}

func (s_ZR_8 *ENfmDmMaozH) Fchown(a_qFVXPl41Na, bHj6qUKH3YD int) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ mft2Ul(func() error {
		return  /*line :1*/ syscall.FQi5_BSB(s_ZR_8.X8OEsVYSby, a_qFVXPl41Na, bHj6qUKH3YD)
	})
}

func (s_ZR_8 *ENfmDmMaozH) Ftruncate(mxYBfG int64) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	return  /*line :1*/ mft2Ul(func() error {
		return  /*line :1*/ syscall.Owge05R_TH_(s_ZR_8.X8OEsVYSby, mxYBfG)
	})
}

func (s_ZR_8 *ENfmDmMaozH) RawControl(orn_5D2Pdk func(uintptr)) error {
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	 /*line :1*/ orn_5D2Pdk( /*line :1*/ uintptr(s_ZR_8.X8OEsVYSby))
	return nil
}

func mft2Ul(boMGLHn func() error) error {
	for {
		dtU8tBiUGS4 :=  /*line :1*/ boMGLHn()
		if dtU8tBiUGS4 != syscall.EINTR {
			return dtU8tBiUGS4
		}
	}
}
