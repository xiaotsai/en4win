//line :1



































package hPMrClpC

import (
	errors "iAHoxjmM"
	poll "MjXXMR"
	safefilepath "FMGFnwa9"
	testlog "X0MraxO81Me"
	io "xI9ai2djaJ2"
	fs "y1BamVjnXsaa"
	"runtime"
	syscall "bUKeamF"
	time "fRAfQd_"
	"unsafe"
)


func (qzeVi328uMg *BvGocYcXx) Name() string	{ return qzeVi328uMg.b5ZMHrO5AlDN }







var (
	RUTFbiB6NV	=  /*line :1*/ He3g9sbjo( /*line :1*/ uintptr(syscall.IZGxwzygxL2), "/dev/stdin")
	Vrq37o	=  /*line :1*/ He3g9sbjo( /*line :1*/ uintptr(syscall.PjqIbe), "/dev/stdout")
	BUPxqQD	=  /*line :1*/ He3g9sbjo( /*line :1*/ uintptr(syscall.ZAFZU4cW), "/dev/stderr")
)



const (
	
	O_RDONLY	int	= syscall.O_RDONLY	
	O_WRONLY	int	= syscall.O_WRONLY	
	O_RDWR	int	= syscall.O_RDWR	
	
	O_APPEND	int	= syscall.O_APPEND	
	O_CREATE	int	= syscall.O_CREAT	
	O_EXCL	int	= syscall.O_EXCL	
	O_SYNC	int	= syscall.O_SYNC	
	O_TRUNC	int	= syscall.O_TRUNC	
)




const (
	SEEK_SET	int	= 0	
	SEEK_CUR	int	= 1	
	SEEK_END	int	= 2	
)



type TxksqbdUdg struct {
	RKxhH_25o	string
	TaJcLZ	string
	Tqz894pL_xQ	string
	OsvUk1	error
}

func (w50uhPri *TxksqbdUdg) Error() string {
	return w50uhPri.RKxhH_25o + " " + w50uhPri.TaJcLZ + " " + w50uhPri.Tqz894pL_xQ + ": " +  /*line :1*/ w50uhPri.OsvUk1.Error()
}

func (w50uhPri *TxksqbdUdg) Unwrap() error {
	return w50uhPri.OsvUk1
}




func (qzeVi328uMg *BvGocYcXx) Read(nR2aPlreQFZA []byte) (zHDjCTHI int, ugqb2IW error) {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("read"); ugqb2IW != nil {
		return 0, ugqb2IW
	}
	zHDjCTHI, w50uhPri :=  /*line :1*/ qzeVi328uMg.iaxc8GTItjn(nR2aPlreQFZA)
	return zHDjCTHI,  /*line :1*/ qzeVi328uMg.awBETn("read", w50uhPri)
}





func (qzeVi328uMg *BvGocYcXx) ReadAt(nR2aPlreQFZA []byte, uPtRbvJhmT int64) (zHDjCTHI int, ugqb2IW error) {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("read"); ugqb2IW != nil {
		return 0, ugqb2IW
	}

	if uPtRbvJhmT < 0 {
		return 0, &StNL1lveD40f{Aeg62j1VQt: "readat", OmDEtY: qzeVi328uMg.b5ZMHrO5AlDN, OIEifQ0XF:  /*line :1*/ errors.Su6g6hRoi9X("negative offset")}
	}

	for  /*line :1*/ len(nR2aPlreQFZA) > 0 {
		jUuV6w2, w50uhPri :=  /*line :1*/ qzeVi328uMg.mdZW9k(nR2aPlreQFZA, uPtRbvJhmT)
		if w50uhPri != nil {
			ugqb2IW =  /*line :1*/ qzeVi328uMg.awBETn("read", w50uhPri)
			break
		}
		zHDjCTHI += jUuV6w2
		nR2aPlreQFZA = nR2aPlreQFZA[jUuV6w2:]
		uPtRbvJhmT +=  /*line :1*/ int64(jUuV6w2)
	}
	return
}


func (qzeVi328uMg *BvGocYcXx) ReadFrom(dhvYus io.YJ04Fau) (zHDjCTHI int64, ugqb2IW error) {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("write"); ugqb2IW != nil {
		return 0, ugqb2IW
	}
	zHDjCTHI, aXJgDK8P1v, w50uhPri :=  /*line :1*/ qzeVi328uMg.qycbvGov(dhvYus)
	if !aXJgDK8P1v {
		return  /*line :1*/ xMxC0o(qzeVi328uMg, dhvYus)
	}
	return zHDjCTHI,  /*line :1*/ qzeVi328uMg.awBETn("write", w50uhPri)
}

func xMxC0o(qzeVi328uMg *BvGocYcXx, dhvYus io.YJ04Fau) (int64, error) {
	return  /*line :1*/ io.FxikaFo5o7OE(eNZz1HXeJZdu{qzeVi328uMg}, dhvYus)
}




type eNZz1HXeJZdu struct {
	*BvGocYcXx
}


func (eNZz1HXeJZdu) ReadFrom(eNZz1HXeJZdu) {
	 /*line :1*/ panic("unreachable")
}




func (qzeVi328uMg *BvGocYcXx) Write(nR2aPlreQFZA []byte) (zHDjCTHI int, ugqb2IW error) {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("write"); ugqb2IW != nil {
		return 0, ugqb2IW
	}
	zHDjCTHI, w50uhPri :=  /*line :1*/ qzeVi328uMg.dFbz7Y(nR2aPlreQFZA)
	if zHDjCTHI < 0 {
		zHDjCTHI = 0
	}
	if zHDjCTHI !=  /*line :1*/ len(nR2aPlreQFZA) {
		ugqb2IW = io.ArPWDHfv
	}

	 /*line :1*/ ri3aP0HBL9(qzeVi328uMg, w50uhPri)

	if w50uhPri != nil {
		ugqb2IW =  /*line :1*/ qzeVi328uMg.awBETn("write", w50uhPri)
	}

	return zHDjCTHI, ugqb2IW
}

var dlNYbrOqORL =  /*line :1*/ errors.Su6g6hRoi9X("os: invalid use of WriteAt on file opened with O_APPEND")






func (qzeVi328uMg *BvGocYcXx) WriteAt(nR2aPlreQFZA []byte, uPtRbvJhmT int64) (zHDjCTHI int, ugqb2IW error) {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("write"); ugqb2IW != nil {
		return 0, ugqb2IW
	}
	if qzeVi328uMg.aepC7RCa {
		return 0, dlNYbrOqORL
	}

	if uPtRbvJhmT < 0 {
		return 0, &StNL1lveD40f{Aeg62j1VQt: "writeat", OmDEtY: qzeVi328uMg.b5ZMHrO5AlDN, OIEifQ0XF:  /*line :1*/ errors.Su6g6hRoi9X("negative offset")}
	}

	for  /*line :1*/ len(nR2aPlreQFZA) > 0 {
		jUuV6w2, w50uhPri :=  /*line :1*/ qzeVi328uMg.dVF4WlJL0(nR2aPlreQFZA, uPtRbvJhmT)
		if w50uhPri != nil {
			ugqb2IW =  /*line :1*/ qzeVi328uMg.awBETn("write", w50uhPri)
			break
		}
		zHDjCTHI += jUuV6w2
		nR2aPlreQFZA = nR2aPlreQFZA[jUuV6w2:]
		uPtRbvJhmT +=  /*line :1*/ int64(jUuV6w2)
	}
	return
}






func (qzeVi328uMg *BvGocYcXx) Seek(uYf1j5z2wz_ int64, fKsaFD int) (ysaPOb int64, ugqb2IW error) {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("seek"); ugqb2IW != nil {
		return 0, ugqb2IW
	}
	dhvYus, w50uhPri :=  /*line :1*/ qzeVi328uMg.vX2joB7Izhp(uYf1j5z2wz_, fKsaFD)
	if w50uhPri == nil && qzeVi328uMg.lkuk5KAfIH != nil && dhvYus != 0 {
		w50uhPri = syscall.EISDIR
	}
	if w50uhPri != nil {
		return 0,  /*line :1*/ qzeVi328uMg.awBETn("seek", w50uhPri)
	}
	return dhvYus, nil
}



func (qzeVi328uMg *BvGocYcXx) WriteString(gkORch7na string) (zHDjCTHI int, ugqb2IW error) {
	nR2aPlreQFZA :=  /*line :1*/ unsafe.Slice( /*line :1*/ unsafe.StringData(gkORch7na),  /*line :1*/ len(gkORch7na))
	return  /*line :1*/ qzeVi328uMg.Write(nR2aPlreQFZA)
}




func OEmAGB_T1CQq(jGBs03 string, nLWcea IvL5u8pdn) error {
	plo2YuP5O :=  /*line :1*/ xPdJPzd7sdeS(jGBs03)
	w50uhPri :=  /*line :1*/ bpvY9V8mH(func() error {
		return  /*line :1*/ syscall.HSuLak53N5P5(plo2YuP5O,  /*line :1*/ o2Yu6Mj(nLWcea))
	})

	if w50uhPri != nil {
		return &StNL1lveD40f{Aeg62j1VQt: "mkdir", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}

	if !supportsCreateWithStickyBit && nLWcea&ModeSticky != 0 {
		w50uhPri =  /*line :1*/ am8Fg_Kpc(jGBs03)

		if w50uhPri != nil {
			 /*line :1*/ NvDI5ZeC(jGBs03)
			return w50uhPri
		}
	}

	return nil
}


func am8Fg_Kpc(jGBs03 string) error {
	mtZ2F8Y91J8, ugqb2IW :=  /*line :1*/ ZYa3wuv(jGBs03)
	if ugqb2IW != nil {
		return ugqb2IW
	}
	return  /*line :1*/ Cx1UVVNqLs19(jGBs03,  /*line :1*/ mtZ2F8Y91J8.Mode()|ModeSticky)
}



func UaxlaKYa7jRL(nqzPkc8LPse1 string) error {
	if w50uhPri :=  /*line :1*/ syscall.Ha4WpjjlJ9(nqzPkc8LPse1); w50uhPri != nil {
		 /*line :1*/ testlog.NT3Zg6plraKK(nqzPkc8LPse1)
		return &StNL1lveD40f{Aeg62j1VQt: "chdir", OmDEtY: nqzPkc8LPse1, OIEifQ0XF: w50uhPri}
	}
	if gmvQ_lfv :=  /*line :1*/ testlog.TDkzeK(); gmvQ_lfv != nil {
		e_1HldxAQ, ugqb2IW :=  /*line :1*/ OWhXb5XTeJpD()
		if ugqb2IW == nil {
			 /*line :1*/ gmvQ_lfv.Chdir(e_1HldxAQ)
		}
	}
	return nil
}





func Cvk6wxcjw(jGBs03 string) (*BvGocYcXx, error) {
	return  /*line :1*/ FxDeI2Pd3bSR(jGBs03, O_RDONLY, 0)
}






func KhCrLpPxJaXz(jGBs03 string) (*BvGocYcXx, error) {
	return  /*line :1*/ FxDeI2Pd3bSR(jGBs03, O_RDWR|O_CREATE|O_TRUNC, 0666)
}







func FxDeI2Pd3bSR(jGBs03 string, frXiPOI int, nLWcea IvL5u8pdn) (*BvGocYcXx, error) {
	 /*line :1*/ testlog.NT3Zg6plraKK(jGBs03)
	qzeVi328uMg, ugqb2IW :=  /*line :1*/ hczTpl7(jGBs03, frXiPOI, nLWcea)
	if ugqb2IW != nil {
		return nil, ugqb2IW
	}
	qzeVi328uMg.aepC7RCa = frXiPOI&O_APPEND != 0

	return qzeVi328uMg, nil
}


var nfSVlSi = Z0JXsY






func Fhyy4z(wwUNUbNju, aCrrJt string) error {
	return  /*line :1*/ ozS0sxXV(wwUNUbNju, aCrrJt)
}



func ictytJmDiE(zHDjCTHI int, ugqb2IW error) (int, error) {
	if zHDjCTHI < 0 {
		zHDjCTHI = 0
	}
	return zHDjCTHI, ugqb2IW
}



var i7TqO7UQk7 = false




func (qzeVi328uMg *BvGocYcXx) awBETn(rCSrhXb2tDv string, ugqb2IW error) error {
	if ugqb2IW == nil || ugqb2IW == io.K5Sqco {
		return ugqb2IW
	}
	if ugqb2IW == poll.IQcIMbVlGewk {
		ugqb2IW = AEdOaRUlDO
	} else if i7TqO7UQk7 &&  /*line :1*/ errors.COBastO_C(ugqb2IW, poll.IQcIMbVlGewk) {
		 /*line :1*/ panic("unexpected error wrapping poll.ErrFileClosing: " +  /*line :1*/ ugqb2IW.Error())
	}
	return &StNL1lveD40f{Aeg62j1VQt: rCSrhXb2tDv, OmDEtY: qzeVi328uMg.b5ZMHrO5AlDN, OIEifQ0XF: ugqb2IW}
}










func Xx5OajUY() string {
	return  /*line :1*/ ziojh5simqO()
}














func OpO2HcxZ6W() (string, error) {
	var nqzPkc8LPse1 string

	switch runtime.GOOS {
	case "windows":
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("LocalAppData")
		if nqzPkc8LPse1 == "" {
			return "",  /*line :1*/ errors.Su6g6hRoi9X("%LocalAppData% is not defined")
		}

	case "darwin", "ios":
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("HOME")
		if nqzPkc8LPse1 == "" {
			return "",  /*line :1*/ errors.Su6g6hRoi9X("$HOME is not defined")
		}
		nqzPkc8LPse1 += "/Library/Caches"

	case "plan9":
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("home")
		if nqzPkc8LPse1 == "" {
			return "",  /*line :1*/ errors.Su6g6hRoi9X("$home is not defined")
		}
		nqzPkc8LPse1 += "/lib/cache"

	default:
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("XDG_CACHE_HOME")
		if nqzPkc8LPse1 == "" {
			nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("HOME")
			if nqzPkc8LPse1 == "" {
				return "",  /*line :1*/ errors.Su6g6hRoi9X("neither $XDG_CACHE_HOME nor $HOME are defined")
			}
			nqzPkc8LPse1 += "/.cache"
		}
	}

	return nqzPkc8LPse1, nil
}














func CxovZ0j() (string, error) {
	var nqzPkc8LPse1 string

	switch runtime.GOOS {
	case "windows":
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("AppData")
		if nqzPkc8LPse1 == "" {
			return "",  /*line :1*/ errors.Su6g6hRoi9X("%AppData% is not defined")
		}

	case "darwin", "ios":
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("HOME")
		if nqzPkc8LPse1 == "" {
			return "",  /*line :1*/ errors.Su6g6hRoi9X("$HOME is not defined")
		}
		nqzPkc8LPse1 += "/Library/Application Support"

	case "plan9":
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("home")
		if nqzPkc8LPse1 == "" {
			return "",  /*line :1*/ errors.Su6g6hRoi9X("$home is not defined")
		}
		nqzPkc8LPse1 += "/lib"

	default:
		nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("XDG_CONFIG_HOME")
		if nqzPkc8LPse1 == "" {
			nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("HOME")
			if nqzPkc8LPse1 == "" {
				return "",  /*line :1*/ errors.Su6g6hRoi9X("neither $XDG_CONFIG_HOME nor $HOME are defined")
			}
			nqzPkc8LPse1 += "/.config"
		}
	}

	return nqzPkc8LPse1, nil
}









func JtXDojtMX73_() (string, error) {
	apRXcpKzyAM, kLCxbWgjVSY := "HOME", "$HOME"
	switch runtime.GOOS {
	case "windows":
		apRXcpKzyAM, kLCxbWgjVSY = "USERPROFILE", "%userprofile%"
	case "plan9":
		apRXcpKzyAM, kLCxbWgjVSY = "home", "$home"
	}
	if rB6hatF3c :=  /*line :1*/ LDjFPRBEz(apRXcpKzyAM); rB6hatF3c != "" {
		return rB6hatF3c, nil
	}

	switch runtime.GOOS {
	case "android":
		return "/sdcard", nil
	case "ios":
		return "/", nil
	}
	return "",  /*line :1*/ errors.Su6g6hRoi9X(kLCxbWgjVSY + " is not defined")
}



















func Cx1UVVNqLs19(jGBs03 string, lVhwrRwT IvL5u8pdn) error	{ return  /*line :1*/ eUldLV(jGBs03, lVhwrRwT) }



func (qzeVi328uMg *BvGocYcXx) Chmod(lVhwrRwT IvL5u8pdn) error	{ return  /*line :1*/ qzeVi328uMg.eUldLV(lVhwrRwT) }

























func (qzeVi328uMg *BvGocYcXx) SetDeadline(ekqRHXB time.G4KDkI) error {
	return  /*line :1*/ qzeVi328uMg.r6EYmtEuZE(ekqRHXB)
}





func (qzeVi328uMg *BvGocYcXx) SetReadDeadline(ekqRHXB time.G4KDkI) error {
	return  /*line :1*/ qzeVi328uMg.u9Bx5zXkg1(ekqRHXB)
}







func (qzeVi328uMg *BvGocYcXx) SetWriteDeadline(ekqRHXB time.G4KDkI) error {
	return  /*line :1*/ qzeVi328uMg.eBO8TH(ekqRHXB)
}



func (qzeVi328uMg *BvGocYcXx) SyscallConn() (syscall.CVnV1i, error) {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("SyscallConn"); ugqb2IW != nil {
		return nil, ugqb2IW
	}
	return  /*line :1*/ iZeFLFBUHQ(qzeVi328uMg)
}
















func S8akSn852u(nqzPkc8LPse1 string) fs.XIcFcDgy {
	return  /*line :1*/ kyaJzt4I(nqzPkc8LPse1)
}


func enivCpOMv7oe(gkORch7na, i1MFcC1g1sS string) bool {
	for lWfaal := 0; lWfaal <  /*line :1*/ len(gkORch7na); lWfaal++ {
		for lpNM41IF := 0; lpNM41IF <  /*line :1*/ len(i1MFcC1g1sS); lpNM41IF++ {
			if gkORch7na[lWfaal] == i1MFcC1g1sS[lpNM41IF] {
				return true
			}
		}
	}
	return false
}

type kyaJzt4I string

func (nqzPkc8LPse1 kyaJzt4I) Open(jGBs03 string) (fs.YY1DXRrw, error) {
	vH_Czflhi7G, ugqb2IW :=  /*line :1*/ nqzPkc8LPse1.aQ889ElGO7oF(jGBs03)
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "stat", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}
	qzeVi328uMg, ugqb2IW :=  /*line :1*/ Cvk6wxcjw(vH_Czflhi7G)
	if ugqb2IW != nil {

		ugqb2IW.(*StNL1lveD40f).OmDEtY = jGBs03
		return nil, ugqb2IW
	}
	return qzeVi328uMg, nil
}





func (nqzPkc8LPse1 kyaJzt4I) ReadFile(jGBs03 string) ([]byte, error) {
	vH_Czflhi7G, ugqb2IW :=  /*line :1*/ nqzPkc8LPse1.aQ889ElGO7oF(jGBs03)
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "readfile", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}
	return  /*line :1*/ St3SC4GEAjp(vH_Czflhi7G)
}



func (nqzPkc8LPse1 kyaJzt4I) ReadDir(jGBs03 string) ([]XXSR6RRqno8V, error) {
	vH_Czflhi7G, ugqb2IW :=  /*line :1*/ nqzPkc8LPse1.aQ889ElGO7oF(jGBs03)
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "readdir", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}
	return  /*line :1*/ FbJujn(vH_Czflhi7G)
}

func (nqzPkc8LPse1 kyaJzt4I) Stat(jGBs03 string) (fs.HM4JM2, error) {
	vH_Czflhi7G, ugqb2IW :=  /*line :1*/ nqzPkc8LPse1.aQ889ElGO7oF(jGBs03)
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "stat", OmDEtY: jGBs03, OIEifQ0XF: ugqb2IW}
	}
	qzeVi328uMg, ugqb2IW :=  /*line :1*/ ZYa3wuv(vH_Czflhi7G)
	if ugqb2IW != nil {

		ugqb2IW.(*StNL1lveD40f).OmDEtY = jGBs03
		return nil, ugqb2IW
	}
	return qzeVi328uMg, nil
}


func (nqzPkc8LPse1 kyaJzt4I) aQ889ElGO7oF(jGBs03 string) (string, error) {
	if nqzPkc8LPse1 == "" {
		return "",  /*line :1*/ errors.Su6g6hRoi9X("os: DirFS with empty root")
	}
	if ! /*line :1*/ fs.QuVTpr7RfE(jGBs03) {
		return "", J6ilKFR6o
	}
	jGBs03, ugqb2IW :=  /*line :1*/ safefilepath.T5jStZjs2Nh(jGBs03)
	if ugqb2IW != nil {
		return "", J6ilKFR6o
	}
	if  /*line :1*/ GClWKECc(nqzPkc8LPse1[ /*line :1*/ len(nqzPkc8LPse1)-1]) {
		return  /*line :1*/ string(nqzPkc8LPse1) + jGBs03, nil
	}
	return  /*line :1*/ string(nqzPkc8LPse1) +  /*line :1*/ string(PathSeparator) + jGBs03, nil
}





func St3SC4GEAjp(jGBs03 string) ([]byte, error) {
	qzeVi328uMg, ugqb2IW :=  /*line :1*/ Cvk6wxcjw(jGBs03)
	if ugqb2IW != nil {
		return nil, ugqb2IW
	}
	defer  /*line :1*/ qzeVi328uMg.Close()

	var oCFG5n int
	if qhhuEQfu7Aa, ugqb2IW :=  /*line :1*/ qzeVi328uMg.Stat(); ugqb2IW == nil {
		cPMoeTvw :=  /*line :1*/ qhhuEQfu7Aa.Size()
		if  /*line :1*/ int64( /*line :1*/ int(cPMoeTvw)) == cPMoeTvw {
			oCFG5n =  /*line :1*/ int(cPMoeTvw)
		}
	}
	oCFG5n++

	if oCFG5n < 512 {
		oCFG5n = 512
	}

	itOuxk :=  /*line :1*/ make([]byte, 0, oCFG5n)
	for {
		if  /*line :1*/ len(itOuxk) >=  /*line :1*/ cap(itOuxk) {
			b9ZD7F3 :=  /*line :1*/ append(itOuxk[: /*line :1*/ cap(itOuxk)], 0)
			itOuxk = b9ZD7F3[: /*line :1*/ len(itOuxk)]
		}
		zHDjCTHI, ugqb2IW :=  /*line :1*/ qzeVi328uMg.Read(itOuxk[ /*line :1*/ len(itOuxk): /*line :1*/ cap(itOuxk)])
		itOuxk = itOuxk[: /*line :1*/ len(itOuxk)+zHDjCTHI]
		if ugqb2IW != nil {
			if ugqb2IW == io.K5Sqco {
				ugqb2IW = nil
			}
			return itOuxk, ugqb2IW
		}
	}
}






func UPylce(jGBs03 string, itOuxk []byte, nLWcea IvL5u8pdn) error {
	qzeVi328uMg, ugqb2IW :=  /*line :1*/ FxDeI2Pd3bSR(jGBs03, O_WRONLY|O_CREATE|O_TRUNC, nLWcea)
	if ugqb2IW != nil {
		return ugqb2IW
	}
	_, ugqb2IW =  /*line :1*/ qzeVi328uMg.Write(itOuxk)
	if kLN0m8 :=  /*line :1*/ qzeVi328uMg.Close(); kLN0m8 != nil && ugqb2IW == nil {
		ugqb2IW = kLN0m8
	}
	return ugqb2IW
}
