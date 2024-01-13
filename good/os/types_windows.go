//line :1
package hPMrClpC

import (
	windows "LdptURlN"
	sync "sync"
	syscall "bUKeamF"
	time "fRAfQd_"
	"unsafe"
)


type aqsPbX2HE5 struct {
	bs06oGoQ3TBi	string

	
	SdHjLcpjII6V	uint32
	Bc1qzp	syscall.T8WbUqAC3v
	DQWneSrB7	syscall.T8WbUqAC3v
	EXpYi_CIp	syscall.T8WbUqAC3v
	HI5PZdgeZGMa	uint32
	GrKnMNqop2	uint32

	
	UJo5NQk2Qp4p	uint32

	
	fnKqnIU	uint32

	
	sync.DIRWe11KYlYa
	hlEnW1ut97N	string
	mDTbZMq8xN	uint32
	f41XmvC5Lwo	uint32
	epaFt6eLrP	uint32
	dd_9IGm	bool
}



func iDDVPC(sJC41DS string, dB59YpfJ syscall.SRlvVjrYa) (wkaQ2FaAsI *aqsPbX2HE5, ugqb2IW error) {
	var b9ZD7F3 syscall.HCYt4WU_Wmzb
	ugqb2IW =  /*line :1*/ syscall.D1OiNlWWBa(dB59YpfJ, &b9ZD7F3)
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "GetFileInformationByHandle", OmDEtY: sJC41DS, OIEifQ0XF: ugqb2IW}
	}

	var dpkLXXZI2gQ windows.RjxyGV4DfzoK
	ugqb2IW =  /*line :1*/ windows.JaAxDHL(dB59YpfJ, windows.FileAttributeTagInfo, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&dpkLXXZI2gQ)),  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(dpkLXXZI2gQ)))
	if ugqb2IW != nil {
		if mPGaAKbpJA, xYbm4BTQw := ugqb2IW.(syscall.O9Mga3); xYbm4BTQw && mPGaAKbpJA == windows.ERROR_INVALID_PARAMETER {

			dpkLXXZI2gQ.IBP9s_uV = 0
		} else {
			return nil, &StNL1lveD40f{Aeg62j1VQt: "GetFileInformationByHandleEx", OmDEtY: sJC41DS, OIEifQ0XF: ugqb2IW}
		}
	}

	return &aqsPbX2HE5{
		bs06oGoQ3TBi:	 /*line :1*/ tVaoof6(sJC41DS),
		SdHjLcpjII6V:	b9ZD7F3.FrYcZLBau7r,
		Bc1qzp:	b9ZD7F3.JYWhEbWo_W,
		DQWneSrB7:	b9ZD7F3.EN4G6pVm_kA4,
		EXpYi_CIp:	b9ZD7F3.YMVB_o2o,
		HI5PZdgeZGMa:	b9ZD7F3.YX3QVvo,
		GrKnMNqop2:	b9ZD7F3.WpwMOiFxZj,
		mDTbZMq8xN:	b9ZD7F3.MoZwkiXS,
		f41XmvC5Lwo:	b9ZD7F3.G219mR_c92nY,
		epaFt6eLrP:	b9ZD7F3.P6HqpaH,
		UJo5NQk2Qp4p:	dpkLXXZI2gQ.IBP9s_uV,
	}, nil
}



func fDf_eyeL(b9ZD7F3 *syscall.RsoZmJigCkGe) *aqsPbX2HE5 {
	wkaQ2FaAsI := &aqsPbX2HE5{
		SdHjLcpjII6V:	b9ZD7F3.BBGnBMSj,
		Bc1qzp:	b9ZD7F3.B0BiCI,
		DQWneSrB7:	b9ZD7F3.OFY5aC,
		EXpYi_CIp:	b9ZD7F3.Ve2KvcJ,
		HI5PZdgeZGMa:	b9ZD7F3.Yo4A8bC,
		GrKnMNqop2:	b9ZD7F3.G4llND5o,
	}
	if b9ZD7F3.BBGnBMSj&syscall.FILE_ATTRIBUTE_REPARSE_POINT != 0 {

		wkaQ2FaAsI.UJo5NQk2Qp4p = b9ZD7F3.EjPPdaEO
	}
	return wkaQ2FaAsI
}

func (wkaQ2FaAsI *aqsPbX2HE5) iudtDSkGd() bool {

	return wkaQ2FaAsI.UJo5NQk2Qp4p == syscall.IO_REPARSE_TAG_SYMLINK ||
		wkaQ2FaAsI.UJo5NQk2Qp4p == windows.IO_REPARSE_TAG_MOUNT_POINT
}

func (wkaQ2FaAsI *aqsPbX2HE5) Size() int64 {
	return  /*line :1*/ int64(wkaQ2FaAsI.HI5PZdgeZGMa)<<32 +  /*line :1*/ int64(wkaQ2FaAsI.GrKnMNqop2)
}

func (wkaQ2FaAsI *aqsPbX2HE5) Mode() (jUuV6w2 IvL5u8pdn) {
	if wkaQ2FaAsI.SdHjLcpjII6V&syscall.FILE_ATTRIBUTE_READONLY != 0 {
		jUuV6w2 |= 0444
	} else {
		jUuV6w2 |= 0666
	}
	if  /*line :1*/ wkaQ2FaAsI.iudtDSkGd() {
		return jUuV6w2 | ModeSymlink
	}
	if wkaQ2FaAsI.SdHjLcpjII6V&syscall.FILE_ATTRIBUTE_DIRECTORY != 0 {
		jUuV6w2 |= ModeDir | 0111
	}
	switch wkaQ2FaAsI.fnKqnIU {
	case syscall.FILE_TYPE_PIPE:
		jUuV6w2 |= ModeNamedPipe
	case syscall.FILE_TYPE_CHAR:
		jUuV6w2 |= ModeDevice | ModeCharDevice
	}
	if wkaQ2FaAsI.SdHjLcpjII6V&syscall.FILE_ATTRIBUTE_REPARSE_POINT != 0 && jUuV6w2&ModeType == 0 {
		jUuV6w2 |= ModeIrregular
	}
	return jUuV6w2
}

func (wkaQ2FaAsI *aqsPbX2HE5) ModTime() time.G4KDkI {
	return  /*line :1*/ time.FRXtx9QnU(0,  /*line :1*/ wkaQ2FaAsI.EXpYi_CIp.Nanoseconds())
}


func (wkaQ2FaAsI *aqsPbX2HE5) Sys() any {
	return &syscall.Sf7ZL_oAaEE{
		MegT05:	wkaQ2FaAsI.SdHjLcpjII6V,
		LaH9Sgf:	wkaQ2FaAsI.Bc1qzp,
		NfUjFko6GhPa:	wkaQ2FaAsI.DQWneSrB7,
		QNU6t9O:	wkaQ2FaAsI.EXpYi_CIp,
		XY0PVH6ujAbY:	wkaQ2FaAsI.HI5PZdgeZGMa,
		JUVHzxEdI:	wkaQ2FaAsI.GrKnMNqop2,
	}
}

func (wkaQ2FaAsI *aqsPbX2HE5) zBzkcpqfh() error {
	 /*line :1*/ wkaQ2FaAsI.Lock()
	defer  /*line :1*/ wkaQ2FaAsI.Unlock()
	if wkaQ2FaAsI.hlEnW1ut97N == "" {

		return nil
	}
	var sJC41DS string
	if wkaQ2FaAsI.dd_9IGm {
		sJC41DS = wkaQ2FaAsI.hlEnW1ut97N + `\` + wkaQ2FaAsI.bs06oGoQ3TBi
	} else {
		sJC41DS = wkaQ2FaAsI.hlEnW1ut97N
	}
	azcfkzx5, ugqb2IW :=  /*line :1*/ syscall.GcOmFfsibES(sJC41DS)
	if ugqb2IW != nil {
		return ugqb2IW
	}

	vuMKhTbH0H :=  /*line :1*/ uint32(syscall.FILE_FLAG_BACKUP_SEMANTICS | syscall.FILE_FLAG_OPEN_REPARSE_POINT)

	dB59YpfJ, ugqb2IW :=  /*line :1*/ syscall.EyjYVpa(azcfkzx5, 0, 0, nil, syscall.OPEN_EXISTING, vuMKhTbH0H, 0)
	if ugqb2IW != nil {
		return ugqb2IW
	}
	defer  /*line :1*/ syscall.FhKJLgXjwG(dB59YpfJ)
	var lWfaal syscall.HCYt4WU_Wmzb
	ugqb2IW =  /*line :1*/ syscall.D1OiNlWWBa(dB59YpfJ, &lWfaal)
	if ugqb2IW != nil {
		return ugqb2IW
	}
	wkaQ2FaAsI.hlEnW1ut97N = ""
	wkaQ2FaAsI.mDTbZMq8xN = lWfaal.MoZwkiXS
	wkaQ2FaAsI.f41XmvC5Lwo = lWfaal.G219mR_c92nY
	wkaQ2FaAsI.epaFt6eLrP = lWfaal.P6HqpaH
	return nil
}



func (wkaQ2FaAsI *aqsPbX2HE5) m1m5hx1YgE7(sJC41DS string) error {
	wkaQ2FaAsI.hlEnW1ut97N = sJC41DS
	if ! /*line :1*/ hJq_1R(wkaQ2FaAsI.hlEnW1ut97N) {
		var ugqb2IW error
		wkaQ2FaAsI.hlEnW1ut97N, ugqb2IW =  /*line :1*/ syscall.LmpGp3wH_T(wkaQ2FaAsI.hlEnW1ut97N)
		if ugqb2IW != nil {
			return &StNL1lveD40f{Aeg62j1VQt: "FullPath", OmDEtY: sJC41DS, OIEifQ0XF: ugqb2IW}
		}
	}
	wkaQ2FaAsI.bs06oGoQ3TBi =  /*line :1*/ tVaoof6(sJC41DS)
	return nil
}

func a1uWNQtrIB(pCrdcIx7, mWN7DwuxxXU *aqsPbX2HE5) bool {
	w50uhPri :=  /*line :1*/ pCrdcIx7.zBzkcpqfh()
	if w50uhPri != nil {
		return false
	}
	w50uhPri =  /*line :1*/ mWN7DwuxxXU.zBzkcpqfh()
	if w50uhPri != nil {
		return false
	}
	return pCrdcIx7.mDTbZMq8xN == mWN7DwuxxXU.mDTbZMq8xN && pCrdcIx7.f41XmvC5Lwo == mWN7DwuxxXU.f41XmvC5Lwo && pCrdcIx7.epaFt6eLrP == mWN7DwuxxXU.epaFt6eLrP
}


func ooTH8O(mtZ2F8Y91J8 LWsU3x8MX) time.G4KDkI {
	return  /*line :1*/ time.FRXtx9QnU(0,  /*line :1*/ mtZ2F8Y91J8.Sys().(*syscall.Sf7ZL_oAaEE).NfUjFko6GhPa.Nanoseconds())
}
