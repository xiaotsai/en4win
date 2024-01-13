//line :1
package fuA83L

import "unsafe"

type IProvideClassInfo struct {
	IUnknown
}

type RCAfCVV0Ol struct {
	Tjkg7t5vTCa
	BmlCvaPHRDhB	uintptr
}

func (y0jiLdYHXNnx *IProvideClassInfo) VTable() *RCAfCVV0Ol {
	return (* /*line :1*/ RCAfCVV0Ol)( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.RawVTable))
}

func (y0jiLdYHXNnx *IProvideClassInfo) GetClassInfo() (aSfS2VpTkh *ITypeInfo, zsgC7M1 error) {
	aSfS2VpTkh, zsgC7M1 =  /*line :1*/ cl6GDk7T6(y0jiLdYHXNnx)
	return
}
