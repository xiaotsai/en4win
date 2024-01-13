//line :1
package fQXlzv

import (
	"internal/bytealg"
	utf8 "CuAc0pPZwf"
	"unsafe"
)

type CKrNlN3otWSF struct {
	b5BOvab_cDqs	*CKrNlN3otWSF
	xTq2q3	[]byte
}

//go:nosplit
//go:nocheckptr
func gb28iaMMo(hc5uhVn unsafe.Pointer) unsafe.Pointer {
	hB0GXVbS :=  /*line :1*/ uintptr(hc5uhVn)
	return  /*line :1*/ unsafe.Pointer(hB0GXVbS ^ 0)
}

func (naCMgzayVI *CKrNlN3otWSF) iU4YOPPZt52() {
	if naCMgzayVI.b5BOvab_cDqs == nil {

		naCMgzayVI.b5BOvab_cDqs = (* /*line :1*/ CKrNlN3otWSF)( /*line :1*/ gb28iaMMo( /*line :1*/ unsafe.Pointer(naCMgzayVI)))
	} else if naCMgzayVI.b5BOvab_cDqs != naCMgzayVI {
		 /*line :1*/ panic("strings: illegal use of non-zero Builder copied by value")
	}
}

func (naCMgzayVI *CKrNlN3otWSF) String() string {
	return  /*line :1*/ unsafe.String( /*line :1*/ unsafe.SliceData(naCMgzayVI.xTq2q3),  /*line :1*/ len(naCMgzayVI.xTq2q3))
}

func (naCMgzayVI *CKrNlN3otWSF) Len() int	{ return  /*line :1*/ len(naCMgzayVI.xTq2q3) }

func (naCMgzayVI *CKrNlN3otWSF) Cap() int	{ return  /*line :1*/ cap(naCMgzayVI.xTq2q3) }

func (naCMgzayVI *CKrNlN3otWSF) Reset() {
	naCMgzayVI.b5BOvab_cDqs = nil
	naCMgzayVI.xTq2q3 = nil
}

func (naCMgzayVI *CKrNlN3otWSF) _9GjqP6VO6sn(dUrxzl int) {
	qIFlayf :=  /*line :1*/ bytealg.MakeNoZero(2* /*line :1*/ cap(naCMgzayVI.xTq2q3) + dUrxzl)[: /*line :1*/ len(naCMgzayVI.xTq2q3)]
	 /*line :1*/ copy(qIFlayf, naCMgzayVI.xTq2q3)
	naCMgzayVI.xTq2q3 = qIFlayf
}

func (naCMgzayVI *CKrNlN3otWSF) Grow(dUrxzl int) {
	 /*line :1*/ naCMgzayVI.iU4YOPPZt52()
	if dUrxzl < 0 {
		 /*line :1*/ panic("strings.Builder.Grow: negative count")
	}
	if  /*line :1*/ cap(naCMgzayVI.xTq2q3)- /*line :1*/ len(naCMgzayVI.xTq2q3) < dUrxzl {
		 /*line :1*/ naCMgzayVI._9GjqP6VO6sn(dUrxzl)
	}
}

func (naCMgzayVI *CKrNlN3otWSF) Write(hc5uhVn []byte) (int, error) {
	 /*line :1*/ naCMgzayVI.iU4YOPPZt52()
	naCMgzayVI.xTq2q3 =  /*line :1*/ append(naCMgzayVI.xTq2q3, hc5uhVn...)
	return  /*line :1*/ len(hc5uhVn), nil
}

func (naCMgzayVI *CKrNlN3otWSF) WriteByte(f4rCFxMKmpo byte) error {
	 /*line :1*/ naCMgzayVI.iU4YOPPZt52()
	naCMgzayVI.xTq2q3 =  /*line :1*/ append(naCMgzayVI.xTq2q3, f4rCFxMKmpo)
	return nil
}

func (naCMgzayVI *CKrNlN3otWSF) WriteRune(unWb50A rune) (int, error) {
	 /*line :1*/ naCMgzayVI.iU4YOPPZt52()
	dUrxzl :=  /*line :1*/ len(naCMgzayVI.xTq2q3)
	naCMgzayVI.xTq2q3 =  /*line :1*/ utf8.Ht7oMzd(naCMgzayVI.xTq2q3, unWb50A)
	return  /*line :1*/ len(naCMgzayVI.xTq2q3) - dUrxzl, nil
}

func (naCMgzayVI *CKrNlN3otWSF) WriteString(w4GNKhKtw string) (int, error) {
	 /*line :1*/ naCMgzayVI.iU4YOPPZt52()
	naCMgzayVI.xTq2q3 =  /*line :1*/ append(naCMgzayVI.xTq2q3, w4GNKhKtw...)
	return  /*line :1*/ len(w4GNKhKtw), nil
}
