//line :1
package fuA83L

import (
	"unsafe"
)

func bG7G9rr8v(aYlkZGi []byte) *V8vhLAHnL4 {
	j7O2yy, _ :=  /*line :1*/ dSAhOTAfs(VT_UI1, 0,  /*line :1*/ uint32( /*line :1*/ len(aYlkZGi)))

	if j7O2yy == nil {
		 /*line :1*/ panic("Could not convert []byte to SAFEARRAY")
	}

	for gXX2nE5, y0jiLdYHXNnx := range aYlkZGi {
		 /*line :1*/ ryzZkm(j7O2yy,  /*line :1*/ int64(gXX2nE5),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx)))
	}
	return j7O2yy
}

func vpd7C7N37(aYlkZGi []string) *V8vhLAHnL4 {
	j7O2yy, _ :=  /*line :1*/ dSAhOTAfs(VT_BSTR, 0,  /*line :1*/ uint32( /*line :1*/ len(aYlkZGi)))

	if j7O2yy == nil {
		 /*line :1*/ panic("Could not convert []string to SAFEARRAY")
	}

	for gXX2nE5, y0jiLdYHXNnx := range aYlkZGi {
		 /*line :1*/ ryzZkm(j7O2yy,  /*line :1*/ int64(gXX2nE5),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ XYCQt9(y0jiLdYHXNnx))))
	}
	return j7O2yy
}
