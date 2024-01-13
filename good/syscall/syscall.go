//line :1
package bUKeamF

import "internal/bytealg"

//go:generate go run ./mksyscall_windows.go -systemdll -output zsyscall_windows.go syscall_windows.go security_windows.go

func WbA9Ur(wzPhJhIFoI string) []byte {
	n0ogBVz, gOCcQzbcC :=  /*line :1*/ Clqurc3bx(wzPhJhIFoI)
	if gOCcQzbcC != nil {
		 /*line :1*/ panic("syscall: string with NUL passed to StringByteSlice")
	}
	return n0ogBVz
}

func Clqurc3bx(wzPhJhIFoI string) ([]byte, error) {
	if  /*line :1*/ bytealg.IndexByteString(wzPhJhIFoI, 0) != -1 {
		return nil, EINVAL
	}
	n0ogBVz :=  /*line :1*/ make([]byte,  /*line :1*/ len(wzPhJhIFoI)+1)
	 /*line :1*/ copy(n0ogBVz, wzPhJhIFoI)
	return n0ogBVz, nil
}

func EDJEudzO9OF(wzPhJhIFoI string) *byte	{ return & /*line :1*/ WbA9Ur(wzPhJhIFoI)[0] }

func Oea4iRaRU2r(wzPhJhIFoI string) (*byte, error) {
	n0ogBVz, gOCcQzbcC :=  /*line :1*/ Clqurc3bx(wzPhJhIFoI)
	if gOCcQzbcC != nil {
		return nil, gOCcQzbcC
	}
	return &n0ogBVz[0], nil
}

var eSMWaPW_asgA uintptr

func (iUooHi8N *SLZae0blGj) Unix() (jWFlHzdg9m1V int64, tAjR0818K5K int64) {
	return  /*line :1*/ int64(iUooHi8N.ECuErxxfdoOW),  /*line :1*/ int64(iUooHi8N.WPGLV3fH1AX)
}

func (v3wOtFJUw7Nu *VhkskrPU1) Unix() (jWFlHzdg9m1V int64, tAjR0818K5K int64) {
	return  /*line :1*/ int64(v3wOtFJUw7Nu.L0uhjaO),  /*line :1*/ int64(v3wOtFJUw7Nu.RpHBw1QAUy) * 1000
}

func (iUooHi8N *SLZae0blGj) Nano() int64 {
	return  /*line :1*/ int64(iUooHi8N.ECuErxxfdoOW)*1e9 +  /*line :1*/ int64(iUooHi8N.WPGLV3fH1AX)
}

func (v3wOtFJUw7Nu *VhkskrPU1) Nano() int64 {
	return  /*line :1*/ int64(v3wOtFJUw7Nu.L0uhjaO)*1e9 +  /*line :1*/ int64(v3wOtFJUw7Nu.RpHBw1QAUy)*1000
}

func FLKDEa() int
func QcqSaNcYkFmT(jnuPMp1MD1 int)

func gb9l485q7(q62ppaQD4shI, b0abCyNZ string)
func aCtW2ffDj(q62ppaQD4shI string)
