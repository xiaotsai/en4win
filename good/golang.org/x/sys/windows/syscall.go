//line :1
//go:build windows

package NJ4MerJ

import (
	bytes "wLlfRPv"
	strings "fQXlzv"
	syscall "bUKeamF"
	"unsafe"
)

func NUiwpp4I(bamc83qA3DBr string) ([]byte, error) {
	if  /*line :1*/ strings.Po8D3DCRZq(bamc83qA3DBr, 0) != -1 {
		return nil, syscall.EINVAL
	}
	xINVctm :=  /*line :1*/ make([]byte,  /*line :1*/ len(bamc83qA3DBr)+1)
	 /*line :1*/ copy(xINVctm, bamc83qA3DBr)
	return xINVctm, nil
}

func UgPyZIRW(bamc83qA3DBr string) (*byte, error) {
	xINVctm, jeWMpOaQtWV :=  /*line :1*/ NUiwpp4I(bamc83qA3DBr)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	return &xINVctm[0], nil
}

func NWmrwy(bamc83qA3DBr []byte) string {
	if rRGfxgPH8Kq :=  /*line :1*/ bytes.AXZRWQNqa(bamc83qA3DBr, 0); rRGfxgPH8Kq != -1 {
		bamc83qA3DBr = bamc83qA3DBr[:rRGfxgPH8Kq]
	}
	return  /*line :1*/ string(bamc83qA3DBr)
}

func EaSoZaH(hD4wNgEB *byte) string {
	if hD4wNgEB == nil {
		return ""
	}
	if *hD4wNgEB == 0 {
		return ""
	}

	krzuku := 0
	for nauDv3A :=  /*line :1*/ unsafe.Pointer(hD4wNgEB); *(* /*line :1*/ byte)(nauDv3A) != 0; krzuku++ {
		nauDv3A =  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(nauDv3A) + 1)
	}

	return  /*line :1*/ string( /*line :1*/ unsafe.Slice(hD4wNgEB, krzuku))
}

var qp_TZTgVhz9 uintptr

func (r7ayG7dk *YWaK2Afs) Unix() (c02g4mzkwXHE int64, cHWy1zY4R5z int64) {
	return  /*line :1*/ int64(r7ayG7dk.ECuErxxfdoOW),  /*line :1*/ int64(r7ayG7dk.WPGLV3fH1AX)
}

func (fmaieegZe2E_ *F2i2x7peITi) Unix() (c02g4mzkwXHE int64, cHWy1zY4R5z int64) {
	return  /*line :1*/ int64(fmaieegZe2E_.L0uhjaO),  /*line :1*/ int64(fmaieegZe2E_.RpHBw1QAUy) * 1000
}

func (r7ayG7dk *YWaK2Afs) Nano() int64 {
	return  /*line :1*/ int64(r7ayG7dk.ECuErxxfdoOW)*1e9 +  /*line :1*/ int64(r7ayG7dk.WPGLV3fH1AX)
}

func (fmaieegZe2E_ *F2i2x7peITi) Nano() int64 {
	return  /*line :1*/ int64(fmaieegZe2E_.L0uhjaO)*1e9 +  /*line :1*/ int64(fmaieegZe2E_.RpHBw1QAUy)*1000
}
