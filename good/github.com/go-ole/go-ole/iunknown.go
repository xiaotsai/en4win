//line :1
package fuA83L

import "unsafe"

type IUnknown struct {
	RawVTable *interface{}
}

type Tjkg7t5vTCa struct {
	XVGclIGZIiw	uintptr
	GkT7ZP	uintptr
	Aq2pNFiJR	uintptr
}

type ArT3StaF interface {
	QueryInterface(c1f0ONah *GUID) (c8YZThHoO *IDispatch, zsgC7M1 error)
	AddRef() int32
	Release() int32
}

func (y0jiLdYHXNnx *IUnknown) VTable() *Tjkg7t5vTCa {
	return (* /*line :1*/ Tjkg7t5vTCa)( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx.RawVTable))
}

func (y0jiLdYHXNnx *IUnknown) PutQueryInterface(yxjNU385 *GUID, qVhHBZar4a6 interface{}) error {
	return  /*line :1*/ eRuA_UMJfKp(y0jiLdYHXNnx,  /*line :1*/ y0jiLdYHXNnx.VTable().XVGclIGZIiw, yxjNU385, qVhHBZar4a6)
}

func (y0jiLdYHXNnx *IUnknown) IDispatch(yxjNU385 *GUID) (l6nPEB7Van *IDispatch, zsgC7M1 error) {
	zsgC7M1 =  /*line :1*/ y0jiLdYHXNnx.PutQueryInterface(yxjNU385, &l6nPEB7Van)
	return
}

func (y0jiLdYHXNnx *IUnknown) IEnumVARIANT(yxjNU385 *GUID) (xIxQm4DCSzB *IEnumVARIANT, zsgC7M1 error) {
	zsgC7M1 =  /*line :1*/ y0jiLdYHXNnx.PutQueryInterface(yxjNU385, &xIxQm4DCSzB)
	return
}

func (y0jiLdYHXNnx *IUnknown) QueryInterface(c1f0ONah *GUID) (*IDispatch, error) {
	return  /*line :1*/ sV8DJK(y0jiLdYHXNnx, c1f0ONah)
}

func (y0jiLdYHXNnx *IUnknown) MustQueryInterface(c1f0ONah *GUID) (c8YZThHoO *IDispatch) {
	iGHTYs7q8FBk, zsgC7M1 :=  /*line :1*/ sV8DJK(y0jiLdYHXNnx, c1f0ONah)
	if zsgC7M1 != nil {
		 /*line :1*/ panic(zsgC7M1)
	}
	return iGHTYs7q8FBk
}

func (y0jiLdYHXNnx *IUnknown) AddRef() int32 {
	return  /*line :1*/ y8tTraIf7(y0jiLdYHXNnx)
}

func (y0jiLdYHXNnx *IUnknown) Release() int32 {
	return  /*line :1*/ o1qH5JwREvw2(y0jiLdYHXNnx)
}
