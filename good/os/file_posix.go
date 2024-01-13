//line :1
//go:build unix || (js && wasm) || wasip1 || windows

package hPMrClpC

import (
	"runtime"
	syscall "bUKeamF"
	time "fRAfQd_"
)

func (qzeVi328uMg *BvGocYcXx) Close() error {
	if qzeVi328uMg == nil {
		return J6ilKFR6o
	}
	return  /*line :1*/ qzeVi328uMg.sR7PLgxb.mK1VYLnOuzQk()
}

func (qzeVi328uMg *BvGocYcXx) iaxc8GTItjn(nR2aPlreQFZA []byte) (zHDjCTHI int, ugqb2IW error) {
	zHDjCTHI, ugqb2IW =  /*line :1*/ qzeVi328uMg.vdaw5Om.Read(nR2aPlreQFZA)
	 /*line :1*/ runtime.KeepAlive(qzeVi328uMg)
	return zHDjCTHI, ugqb2IW
}

func (qzeVi328uMg *BvGocYcXx) mdZW9k(nR2aPlreQFZA []byte, uPtRbvJhmT int64) (zHDjCTHI int, ugqb2IW error) {
	zHDjCTHI, ugqb2IW =  /*line :1*/ qzeVi328uMg.vdaw5Om.Pread(nR2aPlreQFZA, uPtRbvJhmT)
	 /*line :1*/ runtime.KeepAlive(qzeVi328uMg)
	return zHDjCTHI, ugqb2IW
}

func (qzeVi328uMg *BvGocYcXx) dFbz7Y(nR2aPlreQFZA []byte) (zHDjCTHI int, ugqb2IW error) {
	zHDjCTHI, ugqb2IW =  /*line :1*/ qzeVi328uMg.vdaw5Om.Write(nR2aPlreQFZA)
	 /*line :1*/ runtime.KeepAlive(qzeVi328uMg)
	return zHDjCTHI, ugqb2IW
}

func (qzeVi328uMg *BvGocYcXx) dVF4WlJL0(nR2aPlreQFZA []byte, uPtRbvJhmT int64) (zHDjCTHI int, ugqb2IW error) {
	zHDjCTHI, ugqb2IW =  /*line :1*/ qzeVi328uMg.vdaw5Om.Pwrite(nR2aPlreQFZA, uPtRbvJhmT)
	 /*line :1*/ runtime.KeepAlive(qzeVi328uMg)
	return zHDjCTHI, ugqb2IW
}

func o2Yu6Mj(lWfaal IvL5u8pdn) (sPO9jym uint32) {
	sPO9jym |=  /*line :1*/ uint32( /*line :1*/ lWfaal.Perm())
	if lWfaal&ModeSetuid != 0 {
		sPO9jym |= syscall.S_ISUID
	}
	if lWfaal&ModeSetgid != 0 {
		sPO9jym |= syscall.S_ISGID
	}
	if lWfaal&ModeSticky != 0 {
		sPO9jym |= syscall.S_ISVTX
	}

	return
}

func eUldLV(jGBs03 string, lVhwrRwT IvL5u8pdn) error {
	plo2YuP5O :=  /*line :1*/ xPdJPzd7sdeS(jGBs03)
	w50uhPri :=  /*line :1*/ bpvY9V8mH(func() error {
		return  /*line :1*/ syscall.FxoMx9RC_A(plo2YuP5O,  /*line :1*/ o2Yu6Mj(lVhwrRwT))
	})
	if w50uhPri != nil {
		return &StNL1lveD40f{Aeg62j1VQt: "chmod", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}
	return nil
}

func (qzeVi328uMg *BvGocYcXx) eUldLV(lVhwrRwT IvL5u8pdn) error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("chmod"); ugqb2IW != nil {
		return ugqb2IW
	}
	if w50uhPri :=  /*line :1*/ qzeVi328uMg.vdaw5Om.Fchmod( /*line :1*/ o2Yu6Mj(lVhwrRwT)); w50uhPri != nil {
		return  /*line :1*/ qzeVi328uMg.awBETn("chmod", w50uhPri)
	}
	return nil
}

func BbDBjmRA(jGBs03 string, dDF68Nn0d, xp6torEPO5 int) error {
	w50uhPri :=  /*line :1*/ bpvY9V8mH(func() error {
		return  /*line :1*/ syscall.VGkqsd(jGBs03, dDF68Nn0d, xp6torEPO5)
	})
	if w50uhPri != nil {
		return &StNL1lveD40f{Aeg62j1VQt: "chown", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}
	return nil
}

func KCaerXK7GlGf(jGBs03 string, dDF68Nn0d, xp6torEPO5 int) error {
	w50uhPri :=  /*line :1*/ bpvY9V8mH(func() error {
		return  /*line :1*/ syscall.Ba3QJ7IqQ(jGBs03, dDF68Nn0d, xp6torEPO5)
	})
	if w50uhPri != nil {
		return &StNL1lveD40f{Aeg62j1VQt: "lchown", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}
	return nil
}

func (qzeVi328uMg *BvGocYcXx) Chown(dDF68Nn0d, xp6torEPO5 int) error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("chown"); ugqb2IW != nil {
		return ugqb2IW
	}
	if w50uhPri :=  /*line :1*/ qzeVi328uMg.vdaw5Om.Fchown(dDF68Nn0d, xp6torEPO5); w50uhPri != nil {
		return  /*line :1*/ qzeVi328uMg.awBETn("chown", w50uhPri)
	}
	return nil
}

func (qzeVi328uMg *BvGocYcXx) Truncate(oCFG5n int64) error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("truncate"); ugqb2IW != nil {
		return ugqb2IW
	}
	if w50uhPri :=  /*line :1*/ qzeVi328uMg.vdaw5Om.Ftruncate(oCFG5n); w50uhPri != nil {
		return  /*line :1*/ qzeVi328uMg.awBETn("truncate", w50uhPri)
	}
	return nil
}

func (qzeVi328uMg *BvGocYcXx) Sync() error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("sync"); ugqb2IW != nil {
		return ugqb2IW
	}
	if w50uhPri :=  /*line :1*/ qzeVi328uMg.vdaw5Om.Fsync(); w50uhPri != nil {
		return  /*line :1*/ qzeVi328uMg.awBETn("sync", w50uhPri)
	}
	return nil
}

func IAae3umpso(jGBs03 string, ooTH8O time.G4KDkI, qupjEoTkJ8 time.G4KDkI) error {
	var ivAqITzPhM [2]syscall.SLZae0blGj
	gOpOEOg := func(lWfaal int, ekqRHXB time.G4KDkI) {
		if  /*line :1*/ ekqRHXB.IsZero() {
			ivAqITzPhM[lWfaal] = syscall.SLZae0blGj{ECuErxxfdoOW: _UTIME_OMIT, WPGLV3fH1AX: _UTIME_OMIT}
		} else {
			ivAqITzPhM[lWfaal] =  /*line :1*/ syscall.J9KJwaNJRl( /*line :1*/ ekqRHXB.UnixNano())
		}
	}
	 /*line :1*/ gOpOEOg(0, ooTH8O)
	 /*line :1*/ gOpOEOg(1, qupjEoTkJ8)
	if w50uhPri :=  /*line :1*/ syscall.TzmRK5Cfk4( /*line :1*/ xPdJPzd7sdeS(jGBs03), ivAqITzPhM[0:]); w50uhPri != nil {
		return &StNL1lveD40f{Aeg62j1VQt: "chtimes", OmDEtY: jGBs03, OIEifQ0XF: w50uhPri}
	}
	return nil
}

func (qzeVi328uMg *BvGocYcXx) Chdir() error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("chdir"); ugqb2IW != nil {
		return ugqb2IW
	}
	if w50uhPri :=  /*line :1*/ qzeVi328uMg.vdaw5Om.Fchdir(); w50uhPri != nil {
		return  /*line :1*/ qzeVi328uMg.awBETn("chdir", w50uhPri)
	}
	return nil
}

func (qzeVi328uMg *BvGocYcXx) r6EYmtEuZE(ekqRHXB time.G4KDkI) error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("SetDeadline"); ugqb2IW != nil {
		return ugqb2IW
	}
	return  /*line :1*/ qzeVi328uMg.vdaw5Om.SetDeadline(ekqRHXB)
}

func (qzeVi328uMg *BvGocYcXx) u9Bx5zXkg1(ekqRHXB time.G4KDkI) error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("SetReadDeadline"); ugqb2IW != nil {
		return ugqb2IW
	}
	return  /*line :1*/ qzeVi328uMg.vdaw5Om.SetReadDeadline(ekqRHXB)
}

func (qzeVi328uMg *BvGocYcXx) eBO8TH(ekqRHXB time.G4KDkI) error {
	if ugqb2IW :=  /*line :1*/ qzeVi328uMg.iUVm9osYjQ2("SetWriteDeadline"); ugqb2IW != nil {
		return ugqb2IW
	}
	return  /*line :1*/ qzeVi328uMg.vdaw5Om.SetWriteDeadline(ekqRHXB)
}

func (qzeVi328uMg *BvGocYcXx) iUVm9osYjQ2(rCSrhXb2tDv string) error {
	if qzeVi328uMg == nil {
		return J6ilKFR6o
	}
	return nil
}

func bpvY9V8mH(bQKQGMZDCp func() error) error {
	for {
		ugqb2IW :=  /*line :1*/ bQKQGMZDCp()
		if ugqb2IW != syscall.EINTR {
			return ugqb2IW
		}
	}
}
