//line :1
//go:build unix || windows || wasip1

package MjXXMR

import (
	errors "iAHoxjmM"
	sync "sync"
	syscall "bUKeamF"
	time "fRAfQd_"
	_ "unsafe"
)

//go:linkname whevaHkiT runtime.nanotime
func whevaHkiT() int64

func tnUlaa()
func pSFKjlBYBDUj(s_ZR_8 uintptr) (uintptr, int)
func nmwBq9Fazp(aEIJQyOL5 uintptr)
func yQ4OQWzfsU(aEIJQyOL5 uintptr, _x0IXPJ int) int
func co_Togw(aEIJQyOL5 uintptr, _x0IXPJ int)
func riIaIKwi7Ge(aEIJQyOL5 uintptr, _x0IXPJ int) int
func fF_t9v1HQvr(aEIJQyOL5 uintptr, wmR0FCrSvE int64, _x0IXPJ int)
func n99YOirzdHM(aEIJQyOL5 uintptr)
func x2NvL9xEA9GL(s_ZR_8 uintptr) bool

type rDHgTN struct {
	gwF0UKHcgp uintptr
}

var vENshEQH75 sync.LhBfwn6wa1x

func (icsxsXB1 *rDHgTN) init(s_ZR_8 *ENfmDmMaozH) error {
	 /*line :1*/ vENshEQH75.Do(tnUlaa)
	aEIJQyOL5, xZ_oGEK :=  /*line :1*/ pSFKjlBYBDUj( /*line :1*/ uintptr(s_ZR_8.X8OEsVYSby))
	if xZ_oGEK != 0 {
		return  /*line :1*/ hoPM5g( /*line :1*/ syscall.O9Mga3(xZ_oGEK))
	}
	icsxsXB1.gwF0UKHcgp = aEIJQyOL5
	return nil
}

func (icsxsXB1 *rDHgTN) zVZzO_nA_() {
	if icsxsXB1.gwF0UKHcgp == 0 {
		return
	}
	 /*line :1*/ nmwBq9Fazp(icsxsXB1.gwF0UKHcgp)
	icsxsXB1.gwF0UKHcgp = 0
}

func (icsxsXB1 *rDHgTN) twNLQiAMSc() {
	if icsxsXB1.gwF0UKHcgp == 0 {
		return
	}
	 /*line :1*/ n99YOirzdHM(icsxsXB1.gwF0UKHcgp)
}

func (icsxsXB1 *rDHgTN) n8O_Sp(_x0IXPJ int, eSOD1mE bool) error {
	if icsxsXB1.gwF0UKHcgp == 0 {
		return nil
	}
	aKyebha :=  /*line :1*/ riIaIKwi7Ge(icsxsXB1.gwF0UKHcgp, _x0IXPJ)
	return  /*line :1*/ i0fasVEji0(aKyebha, eSOD1mE)
}

func (icsxsXB1 *rDHgTN) e7ckqqQRly(eSOD1mE bool) error {
	return  /*line :1*/ icsxsXB1.n8O_Sp('r', eSOD1mE)
}

func (icsxsXB1 *rDHgTN) ofLYpZV(eSOD1mE bool) error {
	return  /*line :1*/ icsxsXB1.n8O_Sp('w', eSOD1mE)
}

func (icsxsXB1 *rDHgTN) dGMuWJsbCVm(_x0IXPJ int, eSOD1mE bool) error {
	if icsxsXB1.gwF0UKHcgp == 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("waiting for unsupported file type")
	}
	aKyebha :=  /*line :1*/ yQ4OQWzfsU(icsxsXB1.gwF0UKHcgp, _x0IXPJ)
	return  /*line :1*/ i0fasVEji0(aKyebha, eSOD1mE)
}

func (icsxsXB1 *rDHgTN) gpxNbQJ(eSOD1mE bool) error {
	return  /*line :1*/ icsxsXB1.dGMuWJsbCVm('r', eSOD1mE)
}

func (icsxsXB1 *rDHgTN) reCQaF(eSOD1mE bool) error {
	return  /*line :1*/ icsxsXB1.dGMuWJsbCVm('w', eSOD1mE)
}

func (icsxsXB1 *rDHgTN) kaTGzw(_x0IXPJ int) {
	if icsxsXB1.gwF0UKHcgp == 0 {
		return
	}
	 /*line :1*/ co_Togw(icsxsXB1.gwF0UKHcgp, _x0IXPJ)
}

func (icsxsXB1 *rDHgTN) gin2cUpW() bool {
	return icsxsXB1.gwF0UKHcgp != 0
}

const (
	pollNoError	= 0
	pollErrClosing	= 1
	pollErrTimeout	= 2
	pollErrNotPollable	= 3
)

func i0fasVEji0(aKyebha int, eSOD1mE bool) error {
	switch aKyebha {
	case pollNoError:
		return nil
	case pollErrClosing:
		return  /*line :1*/ amQJqF(eSOD1mE)
	case pollErrTimeout:
		return ZV68d6_RJG
	case pollErrNotPollable:
		return WyRFZAEmWkl2
	}
	 /*line :1*/ println("unreachable: ", aKyebha)
	 /*line :1*/ panic("unreachable")
}

func (s_ZR_8 *ENfmDmMaozH) SetDeadline(gEDlxU3mfZ time.G4KDkI) error {
	return  /*line :1*/ bRFtssOrU1m(s_ZR_8, gEDlxU3mfZ, 'r'+'w')
}

func (s_ZR_8 *ENfmDmMaozH) SetReadDeadline(gEDlxU3mfZ time.G4KDkI) error {
	return  /*line :1*/ bRFtssOrU1m(s_ZR_8, gEDlxU3mfZ, 'r')
}

func (s_ZR_8 *ENfmDmMaozH) SetWriteDeadline(gEDlxU3mfZ time.G4KDkI) error {
	return  /*line :1*/ bRFtssOrU1m(s_ZR_8, gEDlxU3mfZ, 'w')
}

func bRFtssOrU1m(s_ZR_8 *ENfmDmMaozH, gEDlxU3mfZ time.G4KDkI, _x0IXPJ int) error {
	var wmR0FCrSvE int64
	if ! /*line :1*/ gEDlxU3mfZ.IsZero() {
		wmR0FCrSvE =  /*line :1*/ int64( /*line :1*/ time.Qr_kn6Dzra(gEDlxU3mfZ))
		if wmR0FCrSvE == 0 {
			wmR0FCrSvE = -1
		}
	}
	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.yB3OFgBy1(); dtU8tBiUGS4 != nil {
		return dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.xmYUot4V()
	if s_ZR_8.yw4UOLmW.gwF0UKHcgp == 0 {
		return QeLXq2
	}
	 /*line :1*/ fF_t9v1HQvr(s_ZR_8.yw4UOLmW.gwF0UKHcgp, wmR0FCrSvE, _x0IXPJ)
	return nil
}

func Cbdm6f4lAn(s_ZR_8 uintptr) bool {
	return  /*line :1*/ x2NvL9xEA9GL(s_ZR_8)
}
