//line :1
package wLlfRPv

import (
	errors "iAHoxjmM"
	io "xI9ai2djaJ2"
	utf8 "CuAc0pPZwf"
)






type JORmaQIUc struct {
	vMycj651G	[]byte
	oB4W1iyC	int64	
	rkliZ3ELXw	int	
}



func (dIwXIFT *JORmaQIUc) Len() int {
	if dIwXIFT.oB4W1iyC >=  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) {
		return 0
	}
	return  /*line :1*/ int( /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) - dIwXIFT.oB4W1iyC)
}




func (dIwXIFT *JORmaQIUc) Size() int64	{ return  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) }


func (dIwXIFT *JORmaQIUc) Read(ka9IyYc0 []byte) (lQivwVxjomk int, iRRn6Ng97Jf0 error) {
	if dIwXIFT.oB4W1iyC >=  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) {
		return 0, io.K5Sqco
	}
	dIwXIFT.rkliZ3ELXw = -1
	lQivwVxjomk =  /*line :1*/ copy(ka9IyYc0, dIwXIFT.vMycj651G[dIwXIFT.oB4W1iyC:])
	dIwXIFT.oB4W1iyC +=  /*line :1*/ int64(lQivwVxjomk)
	return
}


func (dIwXIFT *JORmaQIUc) ReadAt(ka9IyYc0 []byte, kMyu7ciY int64) (lQivwVxjomk int, iRRn6Ng97Jf0 error) {

	if kMyu7ciY < 0 {
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("bytes.Reader.ReadAt: negative offset")
	}
	if kMyu7ciY >=  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) {
		return 0, io.K5Sqco
	}
	lQivwVxjomk =  /*line :1*/ copy(ka9IyYc0, dIwXIFT.vMycj651G[kMyu7ciY:])
	if lQivwVxjomk <  /*line :1*/ len(ka9IyYc0) {
		iRRn6Ng97Jf0 = io.K5Sqco
	}
	return
}


func (dIwXIFT *JORmaQIUc) ReadByte() (byte, error) {
	dIwXIFT.rkliZ3ELXw = -1
	if dIwXIFT.oB4W1iyC >=  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) {
		return 0, io.K5Sqco
	}
	ka9IyYc0 := dIwXIFT.vMycj651G[dIwXIFT.oB4W1iyC]
	dIwXIFT.oB4W1iyC++
	return ka9IyYc0, nil
}


func (dIwXIFT *JORmaQIUc) UnreadByte() error {
	if dIwXIFT.oB4W1iyC <= 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("bytes.Reader.UnreadByte: at beginning of slice")
	}
	dIwXIFT.rkliZ3ELXw = -1
	dIwXIFT.oB4W1iyC--
	return nil
}


func (dIwXIFT *JORmaQIUc) ReadRune() (bKhTMqSi rune, kXik_zS1N6 int, iRRn6Ng97Jf0 error) {
	if dIwXIFT.oB4W1iyC >=  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) {
		dIwXIFT.rkliZ3ELXw = -1
		return 0, 0, io.K5Sqco
	}
	dIwXIFT.rkliZ3ELXw =  /*line :1*/ int(dIwXIFT.oB4W1iyC)
	if wOPKkW := dIwXIFT.vMycj651G[dIwXIFT.oB4W1iyC]; wOPKkW < utf8.RuneSelf {
		dIwXIFT.oB4W1iyC++
		return  /*line :1*/ rune(wOPKkW), 1, nil
	}
	bKhTMqSi, kXik_zS1N6 =  /*line :1*/ utf8.EicfpCPn(dIwXIFT.vMycj651G[dIwXIFT.oB4W1iyC:])
	dIwXIFT.oB4W1iyC +=  /*line :1*/ int64(kXik_zS1N6)
	return
}


func (dIwXIFT *JORmaQIUc) UnreadRune() error {
	if dIwXIFT.oB4W1iyC <= 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("bytes.Reader.UnreadRune: at beginning of slice")
	}
	if dIwXIFT.rkliZ3ELXw < 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("bytes.Reader.UnreadRune: previous operation was not ReadRune")
	}
	dIwXIFT.oB4W1iyC =  /*line :1*/ int64(dIwXIFT.rkliZ3ELXw)
	dIwXIFT.rkliZ3ELXw = -1
	return nil
}


func (dIwXIFT *JORmaQIUc) Seek(hldANNRP int64, dbL7n8ENzd3 int) (int64, error) {
	dIwXIFT.rkliZ3ELXw = -1
	var sZXB0o0S6o int64
	switch dbL7n8ENzd3 {
	case io.SeekStart:
		sZXB0o0S6o = hldANNRP
	case io.SeekCurrent:
		sZXB0o0S6o = dIwXIFT.oB4W1iyC + hldANNRP
	case io.SeekEnd:
		sZXB0o0S6o =  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) + hldANNRP
	default:
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("bytes.Reader.Seek: invalid whence")
	}
	if sZXB0o0S6o < 0 {
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("bytes.Reader.Seek: negative position")
	}
	dIwXIFT.oB4W1iyC = sZXB0o0S6o
	return sZXB0o0S6o, nil
}


func (dIwXIFT *JORmaQIUc) WriteTo(yqP0sy io.LXQrGW6BQt) (lQivwVxjomk int64, iRRn6Ng97Jf0 error) {
	dIwXIFT.rkliZ3ELXw = -1
	if dIwXIFT.oB4W1iyC >=  /*line :1*/ int64( /*line :1*/ len(dIwXIFT.vMycj651G)) {
		return 0, nil
	}
	ka9IyYc0 := dIwXIFT.vMycj651G[dIwXIFT.oB4W1iyC:]
	aJCPDy4w, iRRn6Ng97Jf0 :=  /*line :1*/ yqP0sy.Write(ka9IyYc0)
	if aJCPDy4w >  /*line :1*/ len(ka9IyYc0) {
		 /*line :1*/ panic("bytes.Reader.WriteTo: invalid Write count")
	}
	dIwXIFT.oB4W1iyC +=  /*line :1*/ int64(aJCPDy4w)
	lQivwVxjomk =  /*line :1*/ int64(aJCPDy4w)
	if aJCPDy4w !=  /*line :1*/ len(ka9IyYc0) && iRRn6Ng97Jf0 == nil {
		iRRn6Ng97Jf0 = io.ArPWDHfv
	}
	return
}


func (dIwXIFT *JORmaQIUc) Reset(ka9IyYc0 []byte)	{ *dIwXIFT = JORmaQIUc{ka9IyYc0, 0, -1} }


func FGmhhHol(ka9IyYc0 []byte) *JORmaQIUc	{ return &JORmaQIUc{ka9IyYc0, 0, -1} }
