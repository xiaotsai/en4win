//line :1
//go:build !unix

package hPMrClpC

import (
	io "xI9ai2djaJ2"
	"runtime"
	syscall "bUKeamF"
)

func a0oDQ8(sJC41DS string) error {
	if sJC41DS == "" {

		return nil
	}

	if  /*line :1*/ nXnrtsK(sJC41DS) {
		return &StNL1lveD40f{Aeg62j1VQt: "RemoveAll", OmDEtY: sJC41DS, OIEifQ0XF: syscall.EINVAL}
	}

	ugqb2IW :=  /*line :1*/ NvDI5ZeC(sJC41DS)
	if ugqb2IW == nil ||  /*line :1*/ Ac3Db0LaN9ge(ugqb2IW) {
		return nil
	}

	nqzPkc8LPse1, hkIA8b :=  /*line :1*/ Z0JXsY(sJC41DS)
	if hkIA8b != nil {
		if hkIA8b, xYbm4BTQw := hkIA8b.(*StNL1lveD40f); xYbm4BTQw && ( /*line :1*/ Ac3Db0LaN9ge(hkIA8b.OIEifQ0XF) || hkIA8b.OIEifQ0XF == syscall.ENOTDIR) {
			return nil
		}
		return hkIA8b
	}
	if ! /*line :1*/ nqzPkc8LPse1.IsDir() {

		return ugqb2IW
	}

	ugqb2IW = nil
	for {
		ja1FpjT, ugqb2IW :=  /*line :1*/ Cvk6wxcjw(sJC41DS)
		if ugqb2IW != nil {
			if  /*line :1*/ Ac3Db0LaN9ge(ugqb2IW) {

				return nil
			}
			return ugqb2IW
		}

		const reqSize = 1024
		var c8VnEL []string
		var g3kk5E78RW error

		for {
			riSbrZ := 0
			c8VnEL, g3kk5E78RW =  /*line :1*/ ja1FpjT.Readdirnames(reqSize)

			for _, jGBs03 := range c8VnEL {
				kLN0m8 :=  /*line :1*/ RNRFW7(sJC41DS +  /*line :1*/ string(PathSeparator) + jGBs03)
				if ugqb2IW == nil {
					ugqb2IW = kLN0m8
				}
				if kLN0m8 != nil {
					riSbrZ++
				}
			}

			if riSbrZ != reqSize {
				break
			}
		}

		 /*line :1*/ ja1FpjT.Close()

		if g3kk5E78RW == io.K5Sqco {
			break
		}

		if ugqb2IW == nil {
			ugqb2IW = g3kk5E78RW
		}
		if  /*line :1*/ len(c8VnEL) == 0 {
			break
		}

		if  /*line :1*/ len(c8VnEL) < reqSize {
			kLN0m8 :=  /*line :1*/ NvDI5ZeC(sJC41DS)
			if kLN0m8 == nil ||  /*line :1*/ Ac3Db0LaN9ge(kLN0m8) {
				return nil
			}

			if ugqb2IW != nil {

				return ugqb2IW
			}
		}
	}

	kLN0m8 :=  /*line :1*/ NvDI5ZeC(sJC41DS)
	if kLN0m8 == nil ||  /*line :1*/ Ac3Db0LaN9ge(kLN0m8) {
		return nil
	}
	if runtime.GOOS == "windows" &&  /*line :1*/ DeMJFsx(kLN0m8) {
		if wkaQ2FaAsI, ugqb2IW :=  /*line :1*/ ZYa3wuv(sJC41DS); ugqb2IW == nil {
			if ugqb2IW =  /*line :1*/ Cx1UVVNqLs19(sJC41DS,  /*line :1*/ IvL5u8pdn(0200| /*line :1*/ int( /*line :1*/ wkaQ2FaAsI.Mode()))); ugqb2IW == nil {
				kLN0m8 =  /*line :1*/ NvDI5ZeC(sJC41DS)
			}
		}
	}
	if ugqb2IW == nil {
		ugqb2IW = kLN0m8
	}
	return ugqb2IW
}
