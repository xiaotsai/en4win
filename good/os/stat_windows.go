//line :1
package hPMrClpC

import (
	windows "LdptURlN"
	syscall "bUKeamF"
	"unsafe"
)



func (sR7PLgxb *BvGocYcXx) Stat() (LWsU3x8MX, error) {
	if sR7PLgxb == nil {
		return nil, J6ilKFR6o
	}
	return  /*line :1*/ f2EN06rR8(sR7PLgxb.b5ZMHrO5AlDN, sR7PLgxb.vdaw5Om.X8OEsVYSby)
}


func u41sBB2CmRb_(bBjyPxpzd, jGBs03 string, dPy6K_6bg bool) (LWsU3x8MX, error) {
	if  /*line :1*/ len(jGBs03) == 0 {
		return nil, &StNL1lveD40f{Aeg62j1VQt: bBjyPxpzd, OmDEtY: jGBs03, OIEifQ0XF:  /*line :1*/ syscall.O9Mga3(syscall.ERROR_PATH_NOT_FOUND)}
	}
	aRT7PaxNHY, ugqb2IW :=  /*line :1*/ syscall.GcOmFfsibES( /*line :1*/ xPdJPzd7sdeS(jGBs03))
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: bBjyPxpzd, OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}

	
	
	var bg9ZdDuF7 syscall.Sf7ZL_oAaEE
	ugqb2IW =  /*line :1*/ syscall.Z8GRiq0cH(aRT7PaxNHY, syscall.GetFileExInfoStandard, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&bg9ZdDuF7)))

	if ugqb2IW == windows.ERROR_SHARING_VIOLATION {
		var ja1FpjT syscall.RsoZmJigCkGe
		sRFrPs9Jdv, ugqb2IW :=  /*line :1*/ syscall.YWRpP1l(aRT7PaxNHY, &ja1FpjT)
		if ugqb2IW != nil {
			return nil, &StNL1lveD40f{Aeg62j1VQt: "FindFirstFile", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
		}
		 /*line :1*/ syscall.JaVhl9Mkgfp(sRFrPs9Jdv)
		if ja1FpjT.BBGnBMSj&syscall.FILE_ATTRIBUTE_REPARSE_POINT == 0 {

			wkaQ2FaAsI :=  /*line :1*/ fDf_eyeL(&ja1FpjT)
			if ugqb2IW :=  /*line :1*/ wkaQ2FaAsI.m1m5hx1YgE7(jGBs03); ugqb2IW != nil {
				return nil, ugqb2IW
			}
			return wkaQ2FaAsI, nil
		}
	}

	if ugqb2IW == nil && bg9ZdDuF7.MegT05&syscall.FILE_ATTRIBUTE_REPARSE_POINT == 0 {

		wkaQ2FaAsI := &aqsPbX2HE5{
			SdHjLcpjII6V:	bg9ZdDuF7.MegT05,
			Bc1qzp:	bg9ZdDuF7.LaH9Sgf,
			DQWneSrB7:	bg9ZdDuF7.NfUjFko6GhPa,
			EXpYi_CIp:	bg9ZdDuF7.QNU6t9O,
			HI5PZdgeZGMa:	bg9ZdDuF7.XY0PVH6ujAbY,
			GrKnMNqop2:	bg9ZdDuF7.JUVHzxEdI,
		}
		if ugqb2IW :=  /*line :1*/ wkaQ2FaAsI.m1m5hx1YgE7(jGBs03); ugqb2IW != nil {
			return nil, ugqb2IW
		}
		return wkaQ2FaAsI, nil
	}

	dB59YpfJ, ugqb2IW :=  /*line :1*/ syscall.EyjYVpa(aRT7PaxNHY, 0, 0, nil, syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS|syscall.FILE_FLAG_OPEN_REPARSE_POINT, 0)
	if ugqb2IW != nil {

		return nil, &StNL1lveD40f{Aeg62j1VQt: "CreateFile", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}

	mtZ2F8Y91J8, ugqb2IW :=  /*line :1*/ f2EN06rR8(jGBs03, dB59YpfJ)
	 /*line :1*/ syscall.FhKJLgXjwG(dB59YpfJ)
	if ugqb2IW == nil && dPy6K_6bg &&  /*line :1*/ mtZ2F8Y91J8.(*aqsPbX2HE5).iudtDSkGd() {

		dB59YpfJ, ugqb2IW =  /*line :1*/ syscall.EyjYVpa(aRT7PaxNHY, 0, 0, nil, syscall.OPEN_EXISTING, syscall.FILE_FLAG_BACKUP_SEMANTICS, 0)
		if ugqb2IW != nil {

			return nil, &StNL1lveD40f{Aeg62j1VQt: "CreateFile", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
		}
		defer  /*line :1*/ syscall.FhKJLgXjwG(dB59YpfJ)
		return  /*line :1*/ f2EN06rR8(jGBs03, dB59YpfJ)
	}
	return mtZ2F8Y91J8, ugqb2IW
}

func f2EN06rR8(jGBs03 string, dB59YpfJ syscall.SRlvVjrYa) (LWsU3x8MX, error) {
	ehGW_z, ugqb2IW :=  /*line :1*/ syscall.VnXJph2(dB59YpfJ)
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "GetFileType", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}
	switch ehGW_z {
	case syscall.FILE_TYPE_PIPE, syscall.FILE_TYPE_CHAR:
		return &aqsPbX2HE5{bs06oGoQ3TBi:  /*line :1*/ tVaoof6(jGBs03), fnKqnIU: ehGW_z}, nil
	}
	wkaQ2FaAsI, ugqb2IW :=  /*line :1*/ iDDVPC(jGBs03, dB59YpfJ)
	if ugqb2IW != nil {
		return nil, ugqb2IW
	}
	wkaQ2FaAsI.fnKqnIU = ehGW_z
	return wkaQ2FaAsI, ugqb2IW
}


func r5FH06dvGEro(jGBs03 string) (LWsU3x8MX, error) {
	return  /*line :1*/ u41sBB2CmRb_("Stat", jGBs03, true)
}


func x0lxe20b(jGBs03 string) (LWsU3x8MX, error) {
	dPy6K_6bg := false
	if jGBs03 != "" &&  /*line :1*/ GClWKECc(jGBs03[ /*line :1*/ len(jGBs03)-1]) {

		dPy6K_6bg = true
	}
	return  /*line :1*/ u41sBB2CmRb_("Lstat", jGBs03, dPy6K_6bg)
}
