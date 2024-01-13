//line :1
package fQXlzv

import (
	"unsafe"
)











func OdIfPISA27P(w4GNKhKtw string) string {
	if  /*line :1*/ len(w4GNKhKtw) == 0 {
		return ""
	}
	naCMgzayVI :=  /*line :1*/ make([]byte,  /*line :1*/ len(w4GNKhKtw))
	 /*line :1*/ copy(naCMgzayVI, w4GNKhKtw)
	return  /*line :1*/ unsafe.String(&naCMgzayVI[0],  /*line :1*/ len(naCMgzayVI))
}
