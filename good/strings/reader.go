//line :1
package fQXlzv

import (
	errors "iAHoxjmM"
	io "xI9ai2djaJ2"
	utf8 "CuAc0pPZwf"
)





type DbsbDBda struct {
	abK5paR	string
	nqIx_gGruAA	int64	
	vKe2ur	int	
}



func (unWb50A *DbsbDBda) Len() int {
	if unWb50A.nqIx_gGruAA >=  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) {
		return 0
	}
	return  /*line :1*/ int( /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) - unWb50A.nqIx_gGruAA)
}





func (unWb50A *DbsbDBda) Size() int64	{ return  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) }


func (unWb50A *DbsbDBda) Read(naCMgzayVI []byte) (dUrxzl int, iA2u51 error) {
	if unWb50A.nqIx_gGruAA >=  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) {
		return 0, io.K5Sqco
	}
	unWb50A.vKe2ur = -1
	dUrxzl =  /*line :1*/ copy(naCMgzayVI, unWb50A.abK5paR[unWb50A.nqIx_gGruAA:])
	unWb50A.nqIx_gGruAA +=  /*line :1*/ int64(dUrxzl)
	return
}


func (unWb50A *DbsbDBda) ReadAt(naCMgzayVI []byte, t8pDmHLo int64) (dUrxzl int, iA2u51 error) {

	if t8pDmHLo < 0 {
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("strings.Reader.ReadAt: negative offset")
	}
	if t8pDmHLo >=  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) {
		return 0, io.K5Sqco
	}
	dUrxzl =  /*line :1*/ copy(naCMgzayVI, unWb50A.abK5paR[t8pDmHLo:])
	if dUrxzl <  /*line :1*/ len(naCMgzayVI) {
		iA2u51 = io.K5Sqco
	}
	return
}


func (unWb50A *DbsbDBda) ReadByte() (byte, error) {
	unWb50A.vKe2ur = -1
	if unWb50A.nqIx_gGruAA >=  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) {
		return 0, io.K5Sqco
	}
	naCMgzayVI := unWb50A.abK5paR[unWb50A.nqIx_gGruAA]
	unWb50A.nqIx_gGruAA++
	return naCMgzayVI, nil
}


func (unWb50A *DbsbDBda) UnreadByte() error {
	if unWb50A.nqIx_gGruAA <= 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("strings.Reader.UnreadByte: at beginning of string")
	}
	unWb50A.vKe2ur = -1
	unWb50A.nqIx_gGruAA--
	return nil
}


func (unWb50A *DbsbDBda) ReadRune() (dcImFZBbUZ rune, aub89wE2 int, iA2u51 error) {
	if unWb50A.nqIx_gGruAA >=  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) {
		unWb50A.vKe2ur = -1
		return 0, 0, io.K5Sqco
	}
	unWb50A.vKe2ur =  /*line :1*/ int(unWb50A.nqIx_gGruAA)
	if f4rCFxMKmpo := unWb50A.abK5paR[unWb50A.nqIx_gGruAA]; f4rCFxMKmpo < utf8.RuneSelf {
		unWb50A.nqIx_gGruAA++
		return  /*line :1*/ rune(f4rCFxMKmpo), 1, nil
	}
	dcImFZBbUZ, aub89wE2 =  /*line :1*/ utf8.HVAV7aTqxzg(unWb50A.abK5paR[unWb50A.nqIx_gGruAA:])
	unWb50A.nqIx_gGruAA +=  /*line :1*/ int64(aub89wE2)
	return
}


func (unWb50A *DbsbDBda) UnreadRune() error {
	if unWb50A.nqIx_gGruAA <= 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("strings.Reader.UnreadRune: at beginning of string")
	}
	if unWb50A.vKe2ur < 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("strings.Reader.UnreadRune: previous operation was not ReadRune")
	}
	unWb50A.nqIx_gGruAA =  /*line :1*/ int64(unWb50A.vKe2ur)
	unWb50A.vKe2ur = -1
	return nil
}


func (unWb50A *DbsbDBda) Seek(lj42YB46FtI3 int64, _g7Ka1acQj int) (int64, error) {
	unWb50A.vKe2ur = -1
	var eYn3LIbv int64
	switch _g7Ka1acQj {
	case io.SeekStart:
		eYn3LIbv = lj42YB46FtI3
	case io.SeekCurrent:
		eYn3LIbv = unWb50A.nqIx_gGruAA + lj42YB46FtI3
	case io.SeekEnd:
		eYn3LIbv =  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) + lj42YB46FtI3
	default:
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("strings.Reader.Seek: invalid whence")
	}
	if eYn3LIbv < 0 {
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("strings.Reader.Seek: negative position")
	}
	unWb50A.nqIx_gGruAA = eYn3LIbv
	return eYn3LIbv, nil
}


func (unWb50A *DbsbDBda) WriteTo(clNaSK02lkVX io.LXQrGW6BQt) (dUrxzl int64, iA2u51 error) {
	unWb50A.vKe2ur = -1
	if unWb50A.nqIx_gGruAA >=  /*line :1*/ int64( /*line :1*/ len(unWb50A.abK5paR)) {
		return 0, nil
	}
	w4GNKhKtw := unWb50A.abK5paR[unWb50A.nqIx_gGruAA:]
	ta2_ET, iA2u51 :=  /*line :1*/ io.NcJkqIuWUrM(clNaSK02lkVX, w4GNKhKtw)
	if ta2_ET >  /*line :1*/ len(w4GNKhKtw) {
		 /*line :1*/ panic("strings.Reader.WriteTo: invalid WriteString count")
	}
	unWb50A.nqIx_gGruAA +=  /*line :1*/ int64(ta2_ET)
	dUrxzl =  /*line :1*/ int64(ta2_ET)
	if ta2_ET !=  /*line :1*/ len(w4GNKhKtw) && iA2u51 == nil {
		iA2u51 = io.ArPWDHfv
	}
	return
}


func (unWb50A *DbsbDBda) Reset(w4GNKhKtw string)	{ *unWb50A = DbsbDBda{w4GNKhKtw, 0, -1} }



func X4yXgtBA(w4GNKhKtw string) *DbsbDBda	{ return &DbsbDBda{w4GNKhKtw, 0, -1} }
