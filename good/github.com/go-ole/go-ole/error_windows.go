//line :1
package fuA83L

import (
	fmt "kFsoCfy5zWG"
	syscall "bUKeamF"
	utf16 "DtJCLKevRp"
)


func jphUA_3(zgDg5wZHEms int) string {
	
	var dVi3Er_oRTi uint32 = syscall.FORMAT_MESSAGE_FROM_SYSTEM | syscall.FORMAT_MESSAGE_ARGUMENT_ARRAY | syscall.FORMAT_MESSAGE_IGNORE_INSERTS
	wRIFgDKVJXb :=  /*line :1*/ make([]uint16, 300)
	kp7WasHYoVT, zsgC7M1 :=  /*line :1*/ syscall.Lw1VZPgjxcT4(dVi3Er_oRTi, 0,  /*line :1*/ uint32(zgDg5wZHEms), 0, wRIFgDKVJXb, nil)
	if zsgC7M1 != nil {
		return  /*line :1*/ fmt.IBsS3Oc("error %d (FormatMessage failed with: %v)", zgDg5wZHEms, zsgC7M1)
	}

	for ; kp7WasHYoVT > 0 && (wRIFgDKVJXb[kp7WasHYoVT-1] == '\n' || wRIFgDKVJXb[kp7WasHYoVT-1] == '\r'); kp7WasHYoVT-- {
	}
	return  /*line :1*/ string( /*line :1*/ utf16.Q5WZm8(wRIFgDKVJXb[:kp7WasHYoVT]))
}
