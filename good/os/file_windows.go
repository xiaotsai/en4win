//line :1
package hPMrClpC

import (
	errors "iAHoxjmM"
	poll "MjXXMR"
	windows "LdptURlN"
	"runtime"
	sync "sync"
	syscall "bUKeamF"
	"unsafe"
)


const _UTIME_OMIT = -1





type sR7PLgxb struct {
	vdaw5Om	poll.ENfmDmMaozH
	b5ZMHrO5AlDN	string
	lkuk5KAfIH	*c_HZtf7_QR	
	aepC7RCa	bool	
}







func (sR7PLgxb *BvGocYcXx) Fd() uintptr {
	if sR7PLgxb == nil {
		return  /*line :1*/ uintptr(syscall.InvalidHandle)
	}
	return  /*line :1*/ uintptr(sR7PLgxb.vdaw5Om.X8OEsVYSby)
}



func z8Pfzp3si(dB59YpfJ syscall.SRlvVjrYa, jGBs03 string, hag_HJ string) *BvGocYcXx {
	if hag_HJ == "file" {
		var jUuV6w2 uint32
		if  /*line :1*/ syscall.HdjCvUIyL9B(dB59YpfJ, &jUuV6w2) == nil {
			hag_HJ = "console"
		}
		if ekqRHXB, ugqb2IW :=  /*line :1*/ syscall.VnXJph2(dB59YpfJ); ugqb2IW == nil && ekqRHXB == syscall.FILE_TYPE_PIPE {
			hag_HJ = "pipe"
		}
	}

	qzeVi328uMg := &BvGocYcXx{&sR7PLgxb{
		vdaw5Om: poll.ENfmDmMaozH{
			X8OEsVYSby:	dB59YpfJ,
			MDfFexom:	true,
			ZslYVRp1qJ:	true,
		},
		b5ZMHrO5AlDN:	jGBs03,
	}}
	 /*line :1*/ runtime.SetFinalizer(qzeVi328uMg.sR7PLgxb, (*sR7PLgxb).mK1VYLnOuzQk)

	 /*line :1*/ qzeVi328uMg.vdaw5Om.Init(hag_HJ, false)

	return qzeVi328uMg
}


func trayu4WuMEpg(dB59YpfJ syscall.SRlvVjrYa, jGBs03 string) *BvGocYcXx {
	return  /*line :1*/ z8Pfzp3si(dB59YpfJ, jGBs03, "console")
}




func He3g9sbjo(ja1FpjT uintptr, jGBs03 string) *BvGocYcXx {
	dB59YpfJ :=  /*line :1*/ syscall.SRlvVjrYa(ja1FpjT)
	if dB59YpfJ == syscall.InvalidHandle {
		return nil
	}
	return  /*line :1*/ z8Pfzp3si(dB59YpfJ, jGBs03, "file")
}


type c_HZtf7_QR struct {
	f_xxI0	syscall.SRlvVjrYa	
	m6lgHoPUtb	syscall.RsoZmJigCkGe
	ab2GI0U62omX	string
	f6SvFo	bool	
}

func (b9ZD7F3 *c_HZtf7_QR) mK1VYLnOuzQk() error {
	return  /*line :1*/ syscall.JaVhl9Mkgfp(b9ZD7F3.f_xxI0)
}

func ri3aP0HBL9(sR7PLgxb *BvGocYcXx, w50uhPri error) {
}



const DevNull = "NUL"

func gHMl7tDAEd(jGBs03 string) (b9ZD7F3 *c_HZtf7_QR, w50uhPri error) {
	var jqXGpOT string

	sJC41DS :=  /*line :1*/ xPdJPzd7sdeS(jGBs03)

	if  /*line :1*/ len(sJC41DS) == 2 && sJC41DS[1] == ':' {
		jqXGpOT = sJC41DS + `*`
	} else if  /*line :1*/ len(sJC41DS) > 0 {
		f9IKGTDB5 := sJC41DS[ /*line :1*/ len(sJC41DS)-1]
		if f9IKGTDB5 == '/' || f9IKGTDB5 == '\\' {
			jqXGpOT = sJC41DS + `*`
		} else {
			jqXGpOT = sJC41DS + `\*`
		}
	} else {
		jqXGpOT = `\*`
	}
	mY24Wi0vT, w50uhPri :=  /*line :1*/ syscall.GcOmFfsibES(jqXGpOT)
	if w50uhPri != nil {
		return nil, w50uhPri
	}
	b9ZD7F3 =  /*line :1*/ new(c_HZtf7_QR)
	b9ZD7F3.f_xxI0, w50uhPri =  /*line :1*/ syscall.YWRpP1l(mY24Wi0vT, &b9ZD7F3.m6lgHoPUtb)
	if w50uhPri != nil {
		
		
		
		
		
		var bg9ZdDuF7 syscall.Sf7ZL_oAaEE
		azcfkzx5, dliIVWp :=  /*line :1*/ syscall.GcOmFfsibES(sJC41DS)
		if dliIVWp != nil {
			return nil, w50uhPri
		}
		dliIVWp =  /*line :1*/ syscall.Z8GRiq0cH(azcfkzx5, syscall.GetFileExInfoStandard, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&bg9ZdDuF7)))
		if dliIVWp != nil {
			return nil, w50uhPri
		}
		if bg9ZdDuF7.MegT05&syscall.FILE_ATTRIBUTE_DIRECTORY == 0 {
			return nil, syscall.ENOTDIR
		}
		if w50uhPri != syscall.ERROR_FILE_NOT_FOUND {
			return nil, w50uhPri
		}
		b9ZD7F3.f6SvFo = true
	}
	b9ZD7F3.ab2GI0U62omX = sJC41DS
	if ! /*line :1*/ hJq_1R(b9ZD7F3.ab2GI0U62omX) {
		b9ZD7F3.ab2GI0U62omX, w50uhPri =  /*line :1*/ syscall.LmpGp3wH_T(b9ZD7F3.ab2GI0U62omX)
		if w50uhPri != nil {
			 /*line :1*/ b9ZD7F3.mK1VYLnOuzQk()
			return nil, w50uhPri
		}
	}
	return b9ZD7F3, nil
}


func hczTpl7(jGBs03 string, frXiPOI int, nLWcea IvL5u8pdn) (*BvGocYcXx, error) {
	if jGBs03 == "" {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "open", OmDEtY: jGBs03, OIEifQ0XF: syscall.ENOENT}
	}
	sJC41DS :=  /*line :1*/ xPdJPzd7sdeS(jGBs03)
	dhvYus, w50uhPri :=  /*line :1*/ syscall.Ek485AtskVVo(sJC41DS, frXiPOI|syscall.O_CLOEXEC,  /*line :1*/ o2Yu6Mj(nLWcea))
	if w50uhPri != nil {

		if w50uhPri == syscall.ERROR_ACCESS_DENIED && (frXiPOI&O_WRONLY != 0 || frXiPOI&O_RDWR != 0) {
			azcfkzx5, dliIVWp :=  /*line :1*/ syscall.GcOmFfsibES(sJC41DS)
			if dliIVWp == nil {
				var bg9ZdDuF7 syscall.Sf7ZL_oAaEE
				dliIVWp =  /*line :1*/ syscall.Z8GRiq0cH(azcfkzx5, syscall.GetFileExInfoStandard, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&bg9ZdDuF7)))
				if dliIVWp == nil && bg9ZdDuF7.MegT05&syscall.FILE_ATTRIBUTE_DIRECTORY != 0 {
					w50uhPri = syscall.EISDIR
				}
			}
		}
		return nil, &StNL1lveD40f{Aeg62j1VQt: "open", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}
	qzeVi328uMg, w50uhPri :=  /*line :1*/ z8Pfzp3si(dhvYus, jGBs03, "file"), nil
	if w50uhPri != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "open", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}
	return qzeVi328uMg, nil
}

func (sR7PLgxb *sR7PLgxb) mK1VYLnOuzQk() error {
	if sR7PLgxb == nil {
		return syscall.EINVAL
	}
	if sR7PLgxb.lkuk5KAfIH != nil {
		 /*line :1*/ sR7PLgxb.lkuk5KAfIH.mK1VYLnOuzQk()
		sR7PLgxb.lkuk5KAfIH = nil
	}
	var ugqb2IW error
	if w50uhPri :=  /*line :1*/ sR7PLgxb.vdaw5Om.Close(); w50uhPri != nil {
		if w50uhPri == poll.IQcIMbVlGewk {
			w50uhPri = AEdOaRUlDO
		}
		ugqb2IW = &StNL1lveD40f{Aeg62j1VQt: "close", OmDEtY: sR7PLgxb.b5ZMHrO5AlDN, OIEifQ0XF: w50uhPri}
	}

	 /*line :1*/ runtime.SetFinalizer(sR7PLgxb, nil)
	return ugqb2IW
}





func (qzeVi328uMg *BvGocYcXx) vX2joB7Izhp(uYf1j5z2wz_ int64, fKsaFD int) (ysaPOb int64, ugqb2IW error) {
	if qzeVi328uMg.lkuk5KAfIH != nil {

		 /*line :1*/ qzeVi328uMg.lkuk5KAfIH.mK1VYLnOuzQk()
		qzeVi328uMg.lkuk5KAfIH = nil
	}
	ysaPOb, ugqb2IW =  /*line :1*/ qzeVi328uMg.vdaw5Om.Seek(uYf1j5z2wz_, fKsaFD)
	 /*line :1*/ runtime.KeepAlive(qzeVi328uMg)
	return ysaPOb, ugqb2IW
}



func M8sP7CP(jGBs03 string, oCFG5n int64) error {
	qzeVi328uMg, w50uhPri :=  /*line :1*/ FxDeI2Pd3bSR(jGBs03, O_WRONLY, 0666)
	if w50uhPri != nil {
		return w50uhPri
	}
	defer  /*line :1*/ qzeVi328uMg.Close()
	dliIVWp :=  /*line :1*/ qzeVi328uMg.Truncate(oCFG5n)
	if dliIVWp != nil {
		return dliIVWp
	}
	return nil
}



func NvDI5ZeC(jGBs03 string) error {
	baV7tO7i5, w50uhPri :=  /*line :1*/ syscall.GcOmFfsibES( /*line :1*/ xPdJPzd7sdeS(jGBs03))
	if w50uhPri != nil {
		return &StNL1lveD40f{Aeg62j1VQt: "remove", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}

	w50uhPri =  /*line :1*/ syscall.N_MYEJ7Bb(baV7tO7i5)
	if w50uhPri == nil {
		return nil
	}
	dliIVWp :=  /*line :1*/ syscall.PQFNcx(baV7tO7i5)
	if dliIVWp == nil {
		return nil
	}

	if dliIVWp != w50uhPri {
		gGtL0Rd, c_AugaL :=  /*line :1*/ syscall.BDVJtIREn(baV7tO7i5)
		if c_AugaL != nil {
			w50uhPri = c_AugaL
		} else {
			if gGtL0Rd&syscall.FILE_ATTRIBUTE_DIRECTORY != 0 {
				w50uhPri = dliIVWp
			} else if gGtL0Rd&syscall.FILE_ATTRIBUTE_READONLY != 0 {
				if dliIVWp =  /*line :1*/ syscall.D3KivaD(baV7tO7i5, gGtL0Rd&^syscall.FILE_ATTRIBUTE_READONLY); dliIVWp == nil {
					if w50uhPri =  /*line :1*/ syscall.N_MYEJ7Bb(baV7tO7i5); w50uhPri == nil {
						return nil
					}
				}
			}
		}
	}
	return &StNL1lveD40f{Aeg62j1VQt: "remove", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
}

func ozS0sxXV(w5wGKCp, lQkno5_6 string) error {
	w50uhPri :=  /*line :1*/ windows.FJvJDUwWaoc( /*line :1*/ xPdJPzd7sdeS(w5wGKCp),  /*line :1*/ xPdJPzd7sdeS(lQkno5_6))
	if w50uhPri != nil {
		return &TxksqbdUdg{"rename", w5wGKCp, lQkno5_6, w50uhPri}
	}
	return nil
}




func Hl62IkSC() (dhvYus *BvGocYcXx, pgkDcUH *BvGocYcXx, ugqb2IW error) {
	var baV7tO7i5 [2]syscall.SRlvVjrYa
	w50uhPri :=  /*line :1*/ syscall.RyJJSJTWj(baV7tO7i5[:])
	if w50uhPri != nil {
		return nil, nil,  /*line :1*/ BTaHHl("pipe", w50uhPri)
	}
	return  /*line :1*/ z8Pfzp3si(baV7tO7i5[0], "|0", "pipe"),  /*line :1*/ z8Pfzp3si(baV7tO7i5[1], "|1", "pipe"), nil
}

var (
	xs6Ea35Ul	sync.LhBfwn6wa1x
	pGQh8JqD	bool
)

func ziojh5simqO() string {
	 /*line :1*/ xs6Ea35Ul.Do(func() {
		pGQh8JqD = ( /*line :1*/ windows.BVEjpkkAWtv() == nil)
	})
	x5_Y7_oKe := syscall.Jr8XEH5r
	if pGQh8JqD {
		x5_Y7_oKe = windows.RdXWiJTR8
	}
	zHDjCTHI :=  /*line :1*/ uint32(syscall.MAX_PATH)
	for {
		nR2aPlreQFZA :=  /*line :1*/ make([]uint16, zHDjCTHI)
		zHDjCTHI, _ =  /*line :1*/ x5_Y7_oKe( /*line :1*/ uint32( /*line :1*/ len(nR2aPlreQFZA)), &nR2aPlreQFZA[0])
		if zHDjCTHI >  /*line :1*/ uint32( /*line :1*/ len(nR2aPlreQFZA)) {
			continue
		}
		if zHDjCTHI == 3 && nR2aPlreQFZA[1] == ':' && nR2aPlreQFZA[2] == '\\' {

		} else if zHDjCTHI > 0 && nR2aPlreQFZA[zHDjCTHI-1] == '\\' {

			zHDjCTHI--
		}
		return  /*line :1*/ syscall.AODVXp8NM3sd(nR2aPlreQFZA[:zHDjCTHI])
	}
}



func IxsvEzlsgPOx(w5wGKCp, lQkno5_6 string) error {
	zHDjCTHI, ugqb2IW :=  /*line :1*/ syscall.GcOmFfsibES( /*line :1*/ xPdJPzd7sdeS(lQkno5_6))
	if ugqb2IW != nil {
		return &TxksqbdUdg{"link", w5wGKCp, lQkno5_6, ugqb2IW}
	}
	sPO9jym, ugqb2IW :=  /*line :1*/ syscall.GcOmFfsibES( /*line :1*/ xPdJPzd7sdeS(w5wGKCp))
	if ugqb2IW != nil {
		return &TxksqbdUdg{"link", w5wGKCp, lQkno5_6, ugqb2IW}
	}
	ugqb2IW =  /*line :1*/ syscall.Zg0EjHXxfSC(zHDjCTHI, sPO9jym, 0)
	if ugqb2IW != nil {
		return &TxksqbdUdg{"link", w5wGKCp, lQkno5_6, ugqb2IW}
	}
	return nil
}





func BT4hT2AaP(w5wGKCp, lQkno5_6 string) error {

	w5wGKCp =  /*line :1*/ zS1K22(w5wGKCp)

	c94rKWA := w5wGKCp
	if rB6hatF3c :=  /*line :1*/ ccY1eaDH(w5wGKCp); rB6hatF3c == "" {
		if  /*line :1*/ len(w5wGKCp) > 0 &&  /*line :1*/ GClWKECc(w5wGKCp[0]) {

			if rB6hatF3c =  /*line :1*/ ccY1eaDH(lQkno5_6); rB6hatF3c != "" {

				c94rKWA = rB6hatF3c + w5wGKCp
			}
		} else {

			c94rKWA =  /*line :1*/ f4FBKwq7euK(lQkno5_6) + `\` + w5wGKCp
		}
	}

	mtZ2F8Y91J8, ugqb2IW :=  /*line :1*/ ZYa3wuv(c94rKWA)
	gz8w8hPhq := ugqb2IW == nil &&  /*line :1*/ mtZ2F8Y91J8.IsDir()

	zHDjCTHI, ugqb2IW :=  /*line :1*/ syscall.GcOmFfsibES( /*line :1*/ xPdJPzd7sdeS(lQkno5_6))
	if ugqb2IW != nil {
		return &TxksqbdUdg{"symlink", w5wGKCp, lQkno5_6, ugqb2IW}
	}
	sPO9jym, ugqb2IW :=  /*line :1*/ syscall.GcOmFfsibES( /*line :1*/ xPdJPzd7sdeS(w5wGKCp))
	if ugqb2IW != nil {
		return &TxksqbdUdg{"symlink", w5wGKCp, lQkno5_6, ugqb2IW}
	}

	var boN7MA2 uint32 = windows.SYMBOLIC_LINK_FLAG_ALLOW_UNPRIVILEGED_CREATE
	if gz8w8hPhq {
		boN7MA2 |= syscall.SYMBOLIC_LINK_FLAG_DIRECTORY
	}
	ugqb2IW =  /*line :1*/ syscall.HLmNaQslmU(zHDjCTHI, sPO9jym, boN7MA2)
	if ugqb2IW != nil {

		boN7MA2 &^= windows.SYMBOLIC_LINK_FLAG_ALLOW_UNPRIVILEGED_CREATE
		ugqb2IW =  /*line :1*/ syscall.HLmNaQslmU(zHDjCTHI, sPO9jym, boN7MA2)
		if ugqb2IW != nil {
			return &TxksqbdUdg{"symlink", w5wGKCp, lQkno5_6, ugqb2IW}
		}
	}
	return nil
}




func cxNUgPJV8Q(sJC41DS string) (syscall.SRlvVjrYa, error) {
	baV7tO7i5, ugqb2IW :=  /*line :1*/ syscall.GcOmFfsibES(sJC41DS)
	if ugqb2IW != nil {
		return 0, ugqb2IW
	}
	vuMKhTbH0H :=  /*line :1*/ uint32(syscall.FILE_FLAG_BACKUP_SEMANTICS)

	vuMKhTbH0H |= syscall.FILE_FLAG_OPEN_REPARSE_POINT
	dB59YpfJ, ugqb2IW :=  /*line :1*/ syscall.EyjYVpa(baV7tO7i5, 0, 0, nil, syscall.OPEN_EXISTING, vuMKhTbH0H, 0)
	if ugqb2IW != nil {
		return 0, ugqb2IW
	}
	return dB59YpfJ, nil
}









func ijkz7tBIdTJl(sJC41DS string) (string, error) {
	if  /*line :1*/ len(sJC41DS) < 4 || sJC41DS[:4] != `\??\` {

		return sJC41DS, nil
	}

	gkORch7na := sJC41DS[4:]
	switch {
	case  /*line :1*/ len(gkORch7na) >= 2 && gkORch7na[1] == ':':
		return gkORch7na, nil
	case  /*line :1*/ len(gkORch7na) >= 4 && gkORch7na[:4] == `UNC\`:
		return `\\` + gkORch7na[4:], nil
	}

	ugqb2IW :=  /*line :1*/ windows.EOgS9BWI()
	if ugqb2IW != nil {

		return "", ugqb2IW
	}

	dB59YpfJ, ugqb2IW :=  /*line :1*/ cxNUgPJV8Q(sJC41DS)
	if ugqb2IW != nil {
		return "", ugqb2IW
	}
	defer  /*line :1*/ syscall.FhKJLgXjwG(dB59YpfJ)

	dDkChyAkT :=  /*line :1*/ make([]uint16, 100)
	for {
		zHDjCTHI, ugqb2IW :=  /*line :1*/ windows.DoAI5N(dB59YpfJ, &dDkChyAkT[0],  /*line :1*/ uint32( /*line :1*/ len(dDkChyAkT)), windows.VOLUME_NAME_DOS)
		if ugqb2IW != nil {
			return "", ugqb2IW
		}
		if zHDjCTHI <  /*line :1*/ uint32( /*line :1*/ len(dDkChyAkT)) {
			break
		}
		dDkChyAkT =  /*line :1*/ make([]uint16, zHDjCTHI)
	}
	gkORch7na =  /*line :1*/ syscall.AODVXp8NM3sd(dDkChyAkT)
	if  /*line :1*/ len(gkORch7na) > 4 && gkORch7na[:4] == `\\?\` {
		gkORch7na = gkORch7na[4:]
		if  /*line :1*/ len(gkORch7na) > 3 && gkORch7na[:3] == `UNC` {

			return `\` + gkORch7na[3:], nil
		}
		return gkORch7na, nil
	}
	return "",  /*line :1*/ errors.Su6g6hRoi9X("GetFinalPathNameByHandle returned unexpected path: " + gkORch7na)
}

func pnwvDc(sJC41DS string) (string, error) {
	dB59YpfJ, ugqb2IW :=  /*line :1*/ cxNUgPJV8Q(sJC41DS)
	if ugqb2IW != nil {
		return "", ugqb2IW
	}
	defer  /*line :1*/ syscall.FhKJLgXjwG(dB59YpfJ)

	q2ScKNCHvj :=  /*line :1*/ make([]byte, syscall.MAXIMUM_REPARSE_DATA_BUFFER_SIZE)
	var aU0a3s5vow uint32
	ugqb2IW =  /*line :1*/ syscall.EynTCZg(dB59YpfJ, syscall.FSCTL_GET_REPARSE_POINT, nil, 0, &q2ScKNCHvj[0],  /*line :1*/ uint32( /*line :1*/ len(q2ScKNCHvj)), &aU0a3s5vow, nil)
	if ugqb2IW != nil {
		return "", ugqb2IW
	}

	brO2XwUg := (* /*line :1*/ windows.ZIapaYSI)( /*line :1*/ unsafe.Pointer(&q2ScKNCHvj[0]))
	switch brO2XwUg.ZYUZPj {
	case syscall.IO_REPARSE_TAG_SYMLINK:
		x0F6Fz0Yf := (* /*line :1*/ windows.V2LINhLbouGw)( /*line :1*/ unsafe.Pointer(&brO2XwUg.AJJWwJjbP77Q))
		gkORch7na :=  /*line :1*/ x0F6Fz0Yf.Path()
		if x0F6Fz0Yf.D4daKwFX&windows.SYMLINK_FLAG_RELATIVE != 0 {
			return gkORch7na, nil
		}
		return  /*line :1*/ ijkz7tBIdTJl(gkORch7na)
	case windows.IO_REPARSE_TAG_MOUNT_POINT:
		return  /*line :1*/ ijkz7tBIdTJl((* /*line :1*/ windows.Wn05bnBKOXNH)( /*line :1*/ unsafe.Pointer(&brO2XwUg.AJJWwJjbP77Q)).Path())
	default:

		return "", syscall.ENOENT
	}
}



func Jbg7I4xgch(jGBs03 string) (string, error) {
	gkORch7na, ugqb2IW :=  /*line :1*/ pnwvDc( /*line :1*/ xPdJPzd7sdeS(jGBs03))
	if ugqb2IW != nil {
		return "", &StNL1lveD40f{Aeg62j1VQt: "readlink", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}
	return gkORch7na, nil
}
