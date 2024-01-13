//line :1
package fuA83L

import (
	syscall "bUKeamF"
	utf16 "DtJCLKevRp"
	"unsafe"
)

var (
	gjKTuhF5jPHl	=  /*line :1*/ lzu3Aw1UFw.NewProc("CoInitialize")
	k3jONEze	=  /*line :1*/ lzu3Aw1UFw.NewProc("CoInitializeEx")
	yxebb6l	=  /*line :1*/ lzu3Aw1UFw.NewProc("CoUninitialize")
	qCem4rqRJ	=  /*line :1*/ lzu3Aw1UFw.NewProc("CoCreateInstance")
	sF9jowaiEB	=  /*line :1*/ lzu3Aw1UFw.NewProc("CoTaskMemFree")
	gf7uNp6a	=  /*line :1*/ lzu3Aw1UFw.NewProc("CLSIDFromProgID")
	rdH3QTp	=  /*line :1*/ lzu3Aw1UFw.NewProc("CLSIDFromString")
	hCCdIaeLPXZE	=  /*line :1*/ lzu3Aw1UFw.NewProc("StringFromCLSID")
	eDWG2fZuculM	=  /*line :1*/ lzu3Aw1UFw.NewProc("StringFromIID")
	_TAgbvNJ3	=  /*line :1*/ lzu3Aw1UFw.NewProc("IIDFromString")
	nVhFi9	=  /*line :1*/ lzu3Aw1UFw.NewProc("CoGetObject")
	ranzvb9xfl5	=  /*line :1*/ fItlzx3Y.NewProc("GetUserDefaultLCID")
	ejCN9g	=  /*line :1*/ fItlzx3Y.NewProc("RtlMoveMemory")
	x3_IrUPXBw7s	=  /*line :1*/ meGBE6AJ1.NewProc("VariantInit")
	xvm2tSId	=  /*line :1*/ meGBE6AJ1.NewProc("VariantClear")
	d8DnSC_c01	=  /*line :1*/ meGBE6AJ1.NewProc("VariantTimeToSystemTime")
	bxoEvjapY8	=  /*line :1*/ meGBE6AJ1.NewProc("SysAllocString")
	k34RcSZlaY	=  /*line :1*/ meGBE6AJ1.NewProc("SysAllocStringLen")
	qWzaaak	=  /*line :1*/ meGBE6AJ1.NewProc("SysFreeString")
	k5R9Us	=  /*line :1*/ meGBE6AJ1.NewProc("SysStringLen")
	i444Jb4	=  /*line :1*/ meGBE6AJ1.NewProc("CreateDispTypeInfo")
	gw8jZTaAWg	=  /*line :1*/ meGBE6AJ1.NewProc("CreateStdDispatch")
	bFUxYcq	=  /*line :1*/ meGBE6AJ1.NewProc("GetActiveObject")

	li9UNx6O7dt1	=  /*line :1*/ hSEwLBUsf.NewProc("GetMessageW")
	jV19bbg	=  /*line :1*/ hSEwLBUsf.NewProc("DispatchMessageW")
)










func m670LzvIC() (zsgC7M1 error) {

	bnJPAzeC, _, _ :=  /*line :1*/ gjKTuhF5jPHl.Call( /*line :1*/ uintptr(0))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func fFUb0yU1(dYuCQXsb uint32) (zsgC7M1 error) {

	bnJPAzeC, _, _ :=  /*line :1*/ k3jONEze.Call( /*line :1*/ uintptr(0),  /*line :1*/ uintptr(dYuCQXsb))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}










func VLrCgzGab(cUagkN uintptr) (zsgC7M1 error) {

	cUagkN =  /*line :1*/ uintptr(0)
	return  /*line :1*/ m670LzvIC()
}


func KKzYzCZuN(cUagkN uintptr, dYuCQXsb uint32) (zsgC7M1 error) {

	cUagkN =  /*line :1*/ uintptr(0)
	return  /*line :1*/ fFUb0yU1(dYuCQXsb)
}


func DvOlryZ5() {
	 /*line :1*/ yxebb6l.Call()
}


func FXrnmIJ5(oqEuuy4rSzJ uintptr) {
	 /*line :1*/ sF9jowaiEB.Call(oqEuuy4rSzJ)
}














func AtGYsf(dvYDqdU string) (lcDXHYIAiC8j *GUID, zsgC7M1 error) {
	var i15rO5WD3Nx GUID
	iWAuwXn :=  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ syscall.EtPVNA(dvYDqdU)))
	bnJPAzeC, _, _ :=  /*line :1*/ gf7uNp6a.Call(iWAuwXn,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&i15rO5WD3Nx)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	lcDXHYIAiC8j = &i15rO5WD3Nx
	return
}







func QFLzLQWmE(k60skAqXO string) (lcDXHYIAiC8j *GUID, zsgC7M1 error) {
	var i15rO5WD3Nx GUID
	jsOC7nYZ0F6Y :=  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ syscall.EtPVNA(k60skAqXO)))
	bnJPAzeC, _, _ :=  /*line :1*/ rdH3QTp.Call(jsOC7nYZ0F6Y,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&i15rO5WD3Nx)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	lcDXHYIAiC8j = &i15rO5WD3Nx
	return
}


func WefcSOvM(lcDXHYIAiC8j *GUID) (k60skAqXO string, zsgC7M1 error) {
	var cUagkN *uint16
	bnJPAzeC, _, _ :=  /*line :1*/ hCCdIaeLPXZE.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lcDXHYIAiC8j)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&cUagkN)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	k60skAqXO =  /*line :1*/ WMh1Si1a2(cUagkN)
	return
}


func ATrfHSVgMZz(dvYDqdU string) (lcDXHYIAiC8j *GUID, zsgC7M1 error) {
	var i15rO5WD3Nx GUID
	jsOC7nYZ0F6Y :=  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ syscall.EtPVNA(dvYDqdU)))
	bnJPAzeC, _, _ :=  /*line :1*/ _TAgbvNJ3.Call(jsOC7nYZ0F6Y,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&i15rO5WD3Nx)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	lcDXHYIAiC8j = &i15rO5WD3Nx
	return
}


func NI37cFwq(c1f0ONah *GUID) (k60skAqXO string, zsgC7M1 error) {
	var cUagkN *uint16
	bnJPAzeC, _, _ :=  /*line :1*/ eDWG2fZuculM.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c1f0ONah)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&cUagkN)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	k60skAqXO =  /*line :1*/ WMh1Si1a2(cUagkN)
	return
}


func Ne2od0bpE(lcDXHYIAiC8j *GUID, c1f0ONah *GUID) (iGHTYs7q8FBk *IUnknown, zsgC7M1 error) {
	if c1f0ONah == nil {
		c1f0ONah = WySQlWR
	}
	bnJPAzeC, _, _ :=  /*line :1*/ qCem4rqRJ.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lcDXHYIAiC8j)),
		0,
		CLSCTX_SERVER,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c1f0ONah)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&iGHTYs7q8FBk)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func StzosDUxQGsC(lcDXHYIAiC8j *GUID, c1f0ONah *GUID) (iGHTYs7q8FBk *IUnknown, zsgC7M1 error) {
	if c1f0ONah == nil {
		c1f0ONah = WySQlWR
	}
	bnJPAzeC, _, _ :=  /*line :1*/ bFUxYcq.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lcDXHYIAiC8j)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c1f0ONah)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&iGHTYs7q8FBk)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}

type B5BI1YjH struct {
	Qzc4DRyW	uint32
	Oll2LR	uint32
	DfaVJ6h	uint32
	CcS3X0xMs	uint32
}


func MuSZTSPtlcd(aG219XWbeGz_ string, j_Zpgxrfzs *B5BI1YjH, c1f0ONah *GUID) (iGHTYs7q8FBk *IUnknown, zsgC7M1 error) {
	if j_Zpgxrfzs != nil {
		j_Zpgxrfzs.Qzc4DRyW =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(B5BI1YjH{}))
	}
	if c1f0ONah == nil {
		c1f0ONah = WySQlWR
	}
	bnJPAzeC, _, _ :=  /*line :1*/ nVhFi9.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ syscall.EtPVNA(aG219XWbeGz_))),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(j_Zpgxrfzs)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c1f0ONah)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&iGHTYs7q8FBk)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func MTV_iMACtKwC(y0jiLdYHXNnx *KEVetAOpxl0D) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ x3_IrUPXBw7s.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func JZOr1WC(y0jiLdYHXNnx *KEVetAOpxl0D) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ xvm2tSId.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func JX0aFTQl(y0jiLdYHXNnx string) (dxGnd1QR *int16) {
	aH2mV1, _, _ :=  /*line :1*/ bxoEvjapY8.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer( /*line :1*/ syscall.EtPVNA(y0jiLdYHXNnx))))
	dxGnd1QR = (* /*line :1*/ int16)( /*line :1*/ unsafe.Pointer(aH2mV1))
	return
}


func XYCQt9(y0jiLdYHXNnx string) (dxGnd1QR *int16) {
	couM_l :=  /*line :1*/ utf16.J8eAwOfFc([] /*line :1*/ rune(y0jiLdYHXNnx + "\x00"))
	jaRy_z := &couM_l[0]

	aH2mV1, _, _ :=  /*line :1*/ k34RcSZlaY.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jaRy_z)),  /*line :1*/ uintptr( /*line :1*/ len(couM_l)-1))
	dxGnd1QR = (* /*line :1*/ int16)( /*line :1*/ unsafe.Pointer(aH2mV1))
	return
}


func Q6nmuaYvF(y0jiLdYHXNnx *int16) (zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ qWzaaak.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func TkFtieTh_k(y0jiLdYHXNnx *int16) uint32 {
	wtDtc6W_RwaV, _, _ :=  /*line :1*/ k5R9Us.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(y0jiLdYHXNnx)))
	return  /*line :1*/ uint32(wtDtc6W_RwaV)
}






func LHE828nBz(iGHTYs7q8FBk *IUnknown, y0jiLdYHXNnx uintptr, c5iY63koJiTW *IUnknown) (c8YZThHoO *IDispatch, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ gw8jZTaAWg.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iGHTYs7q8FBk)),
		y0jiLdYHXNnx,
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c5iY63koJiTW)),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&c8YZThHoO)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}




func Tih_48(sTrF2DiC4wj *AmPWIefN) (ciuhEAO7Kmj0 *IUnknown, zsgC7M1 error) {
	bnJPAzeC, _, _ :=  /*line :1*/ i444Jb4.Call(
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sTrF2DiC4wj)),
		 /*line :1*/ uintptr( /*line :1*/ KJ4rRUfvUQn()),
		 /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(&ciuhEAO7Kmj0)))
	if bnJPAzeC != 0 {
		zsgC7M1 =  /*line :1*/ RLCP7BQDRyuR(bnJPAzeC)
	}
	return
}


func fUvhgYS(cI2Flq unsafe.Pointer, gA9YADk unsafe.Pointer, nXciaovStia3 uint32) {
	 /*line :1*/ ejCN9g.Call( /*line :1*/ uintptr(cI2Flq),  /*line :1*/ uintptr(gA9YADk),  /*line :1*/ uintptr(nXciaovStia3))
}


func KJ4rRUfvUQn() (a8bAZElj uint32) {
	y7wx667nI, _, _ :=  /*line :1*/ ranzvb9xfl5.Call()
	a8bAZElj =  /*line :1*/ uint32(y7wx667nI)
	return
}




func Fg5xzLw0(hNBXyKP1 *WeDtaklL, uX0s0Q uint32, DUXZ3Jr_ uint32, RC7DNYlDBOe_ uint32) (y7wx667nI int32, zsgC7M1 error) {
	yeRFEg, _, zsgC7M1 :=  /*line :1*/ li9UNx6O7dt1.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hNBXyKP1)),  /*line :1*/ uintptr(uX0s0Q),  /*line :1*/ uintptr(DUXZ3Jr_),  /*line :1*/ uintptr(RC7DNYlDBOe_))
	y7wx667nI =  /*line :1*/ int32(yeRFEg)
	return
}


func CXCvMEfr66Rw(hNBXyKP1 *WeDtaklL) (y7wx667nI int32) {
	yeRFEg, _, _ :=  /*line :1*/ jV19bbg.Call( /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(hNBXyKP1)))
	y7wx667nI =  /*line :1*/ int32(yeRFEg)
	return
}
