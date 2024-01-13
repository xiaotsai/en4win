//line :1
package fuA83L

import "unsafe"

type IDispatch struct {
	IUnknown
}

type Kkapk2AtmM0x struct {
	Tjkg7t5vTCa
	Ya0T5v2ZtB	uintptr
	ZVhUQvpH47G	uintptr
	BBlnxh7	uintptr
	EpI4Ihtn	uintptr
}

func (y0jiLdYHXNnx *IDispatch) VTable() *Kkapk2AtmM0x {
	return (* /*line :1*/ Kkapk2AtmM0x)( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.RawVTable))
}

func (y0jiLdYHXNnx *IDispatch) GetIDsOfName(ljR636fK []string) (yHTMfO []int32, zsgC7M1 error) {
	yHTMfO, zsgC7M1 =  /*line :1*/ dwOLMvw9Hw3(y0jiLdYHXNnx, ljR636fK)
	return
}

func (y0jiLdYHXNnx *IDispatch) Invoke(yHTMfO int32, l6nPEB7Van int16, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D, zsgC7M1 error) {
	bopnVwCF, zsgC7M1 =  /*line :1*/ w7BkGv(y0jiLdYHXNnx, yHTMfO, l6nPEB7Van, uTURuG8...)
	return
}

func (y0jiLdYHXNnx *IDispatch) GetTypeInfoCount() (eT1_dhKA uint32, zsgC7M1 error) {
	eT1_dhKA, zsgC7M1 =  /*line :1*/ uVlLD86X7X18(y0jiLdYHXNnx)
	return
}

func (y0jiLdYHXNnx *IDispatch) GetTypeInfo() (xU2Yaas3M *ITypeInfo, zsgC7M1 error) {
	xU2Yaas3M, zsgC7M1 =  /*line :1*/ bH2MPFaw(y0jiLdYHXNnx)
	return
}





func (y0jiLdYHXNnx *IDispatch) GetSingleIDOfName(xwDC8DsSNM string) (bxyFfmgSJTTh int32, zsgC7M1 error) {
	var zVU9Dwfx7I []int32
	zVU9Dwfx7I, zsgC7M1 =  /*line :1*/ y0jiLdYHXNnx.GetIDsOfName([]string{xwDC8DsSNM})
	if zsgC7M1 != nil {
		return
	}
	bxyFfmgSJTTh = zVU9Dwfx7I[0]
	return
}








func (y0jiLdYHXNnx *IDispatch) InvokeWithOptionalArgs(xwDC8DsSNM string, l6nPEB7Van int16, uTURuG8 []interface{}) (bopnVwCF *KEVetAOpxl0D, zsgC7M1 error) {
	bxyFfmgSJTTh, zsgC7M1 :=  /*line :1*/ y0jiLdYHXNnx.GetSingleIDOfName(xwDC8DsSNM)
	if zsgC7M1 != nil {
		return
	}

	if  /*line :1*/ len(uTURuG8) < 1 {
		bopnVwCF, zsgC7M1 =  /*line :1*/ y0jiLdYHXNnx.Invoke(bxyFfmgSJTTh, l6nPEB7Van)
	} else {
		bopnVwCF, zsgC7M1 =  /*line :1*/ y0jiLdYHXNnx.Invoke(bxyFfmgSJTTh, l6nPEB7Van, uTURuG8...)
	}

	return
}


func (y0jiLdYHXNnx *IDispatch) CallMethod(xwDC8DsSNM string, uTURuG8 ...interface{}) (*KEVetAOpxl0D, error) {
	return  /*line :1*/ y0jiLdYHXNnx.InvokeWithOptionalArgs(xwDC8DsSNM, DISPATCH_METHOD, uTURuG8)
}






func (y0jiLdYHXNnx *IDispatch) GetProperty(xwDC8DsSNM string, uTURuG8 ...interface{}) (*KEVetAOpxl0D, error) {
	return  /*line :1*/ y0jiLdYHXNnx.InvokeWithOptionalArgs(xwDC8DsSNM, DISPATCH_PROPERTYGET, uTURuG8)
}


func (y0jiLdYHXNnx *IDispatch) PutProperty(xwDC8DsSNM string, uTURuG8 ...interface{}) (*KEVetAOpxl0D, error) {
	return  /*line :1*/ y0jiLdYHXNnx.InvokeWithOptionalArgs(xwDC8DsSNM, DISPATCH_PROPERTYPUT, uTURuG8)
}
