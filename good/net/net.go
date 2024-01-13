//line :1

package gF9mZs2

import (
	context "fOYpb3F03EG"
	errors "iAHoxjmM"
	poll "MjXXMR"
	io "xI9ai2djaJ2"
	os "hPMrClpC"
	sync "sync"
	syscall "bUKeamF"
	time "fRAfQd_"
)






type EqbMXsaU1lE interface {
	Network() string	
	String() string	
}




type UJYB3aCQg interface {
	
	
	
	Read(fIadEXIim6l []byte) (doauF8Y int, h_ljk48Bm error)

	
	
	
	Write(fIadEXIim6l []byte) (doauF8Y int, h_ljk48Bm error)

	
	
	Close() error

	
	LocalAddr() EqbMXsaU1lE

	
	RemoteAddr() EqbMXsaU1lE

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	SetDeadline(aeaqWzxJu time.G4KDkI) error

	
	
	
	SetReadDeadline(aeaqWzxJu time.G4KDkI) error

	
	
	
	
	
	SetWriteDeadline(aeaqWzxJu time.G4KDkI) error
}

type xacP7km5pe struct {
	sXhoTdSE8wtb *jmJczkl1
}

func (hl8w2lFX *xacP7km5pe) d30HIwxht() bool	{ return hl8w2lFX != nil && hl8w2lFX.sXhoTdSE8wtb != nil }


func (hl8w2lFX *xacP7km5pe) Read(fIadEXIim6l []byte) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.Read(fIadEXIim6l)
	if h_ljk48Bm != nil && h_ljk48Bm != io.K5Sqco {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "read", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}


func (hl8w2lFX *xacP7km5pe) Write(fIadEXIim6l []byte) (int, error) {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return 0, syscall.EINVAL
	}
	doauF8Y, h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.Write(fIadEXIim6l)
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "write", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return doauF8Y, h_ljk48Bm
}


func (hl8w2lFX *xacP7km5pe) Close() error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.Close()
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "close", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return h_ljk48Bm
}




func (hl8w2lFX *xacP7km5pe) LocalAddr() EqbMXsaU1lE {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return nil
	}
	return hl8w2lFX.sXhoTdSE8wtb.y9pjEf
}




func (hl8w2lFX *xacP7km5pe) RemoteAddr() EqbMXsaU1lE {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return nil
	}
	return hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF
}


func (hl8w2lFX *xacP7km5pe) SetDeadline(aeaqWzxJu time.G4KDkI) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.SetDeadline(aeaqWzxJu); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}


func (hl8w2lFX *xacP7km5pe) SetReadDeadline(aeaqWzxJu time.G4KDkI) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.SetReadDeadline(aeaqWzxJu); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}


func (hl8w2lFX *xacP7km5pe) SetWriteDeadline(aeaqWzxJu time.G4KDkI) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.SetWriteDeadline(aeaqWzxJu); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}



func (hl8w2lFX *xacP7km5pe) SetReadBuffer(c2dmkE6N4N int) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ jWFpXum(hl8w2lFX.sXhoTdSE8wtb, c2dmkE6N4N); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}



func (hl8w2lFX *xacP7km5pe) SetWriteBuffer(c2dmkE6N4N int) error {
	if ! /*line :1*/ hl8w2lFX.d30HIwxht() {
		return syscall.EINVAL
	}
	if h_ljk48Bm :=  /*line :1*/ l6G75dG(hl8w2lFX.sXhoTdSE8wtb, c2dmkE6N4N); h_ljk48Bm != nil {
		return &ZOYGdck{SMNyZk_q: "set", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: nil, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return nil
}








func (hl8w2lFX *xacP7km5pe) File() (t5q9DVOD9 *os.BvGocYcXx, h_ljk48Bm error) {
	t5q9DVOD9, h_ljk48Bm =  /*line :1*/ hl8w2lFX.sXhoTdSE8wtb.uNYdhMYF()
	if h_ljk48Bm != nil {
		h_ljk48Bm = &ZOYGdck{SMNyZk_q: "file", CiQ5sBtmrVnf: hl8w2lFX.sXhoTdSE8wtb.fCaRuBA3, F7g5pQacM05: hl8w2lFX.sXhoTdSE8wtb.y9pjEf, JkQ9XFJbp: hl8w2lFX.sXhoTdSE8wtb.jeyDlYdkG0nF, XkWH4CYwNmvD: h_ljk48Bm}
	}
	return
}




type KlDeyiWNu interface {
	
	
	
	
	
	
	
	
	
	ReadFrom(fMPVz2iGiyq []byte) (doauF8Y int, qxwkC3VxG0 EqbMXsaU1lE, h_ljk48Bm error)

	
	
	
	
	WriteTo(fMPVz2iGiyq []byte, qxwkC3VxG0 EqbMXsaU1lE) (doauF8Y int, h_ljk48Bm error)

	
	
	Close() error

	
	LocalAddr() EqbMXsaU1lE

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	SetDeadline(aeaqWzxJu time.G4KDkI) error

	
	
	
	SetReadDeadline(aeaqWzxJu time.G4KDkI) error

	
	
	
	
	
	SetWriteDeadline(aeaqWzxJu time.G4KDkI) error
}

var jDoRNOxmis struct {
	sync.LhBfwn6wa1x
	dA9JfCE	int
}


func ajNsMEWiu() int {
	 /*line :1*/ jDoRNOxmis.Do(func() { jDoRNOxmis.dA9JfCE =  /*line :1*/ kUqZG__7Y() })
	return jDoRNOxmis.dA9JfCE
}




type Iy6rBQUJ9y7 interface {
	
	Accept() (UJYB3aCQg, error)

	
	
	Close() error

	
	Addr() EqbMXsaU1lE
}


type XL_ACKzCw interface {
	error
	Timeout() bool	

	
	
	
	Temporary() bool
}


var (
	
	xgpz6w9m	=  /*line :1*/ errors.Su6g6hRoi9X("no suitable address found")

	
	tccDv0yXpU	=  /*line :1*/ errors.Su6g6hRoi9X("missing address")

	
	vGDPOEjetn	= bCHmYGGvJ{}
	GcnTprLq	=  /*line :1*/ errors.Su6g6hRoi9X("use of WriteTo with pre-connected connection")
)



type bCHmYGGvJ struct{}

func (bCHmYGGvJ) Error() string	{ return "operation was canceled" }

func (bCHmYGGvJ) Is(h_ljk48Bm error) bool	{ return h_ljk48Bm == context.H0hTZQVavFh }



func gCG7rudfaf(h_ljk48Bm error) error {
	switch h_ljk48Bm {
	case context.H0hTZQVavFh:
		return vGDPOEjetn
	case context.Z201U2N:
		return dklm5juf7rzD
	default:
		return h_ljk48Bm
	}
}




type ZOYGdck struct {
	
	
	SMNyZk_q	string

	
	
	CiQ5sBtmrVnf	string

	
	
	
	F7g5pQacM05	EqbMXsaU1lE

	
	
	
	
	
	
	JkQ9XFJbp	EqbMXsaU1lE

	
	
	XkWH4CYwNmvD	error
}

func (wI0aaz *ZOYGdck) Unwrap() error	{ return wI0aaz.XkWH4CYwNmvD }

func (wI0aaz *ZOYGdck) Error() string {
	if wI0aaz == nil {
		return "<nil>"
	}
	dRoFccaZ := wI0aaz.SMNyZk_q
	if wI0aaz.CiQ5sBtmrVnf != "" {
		dRoFccaZ += " " + wI0aaz.CiQ5sBtmrVnf
	}
	if wI0aaz.F7g5pQacM05 != nil {
		dRoFccaZ += " " +  /*line :1*/ wI0aaz.F7g5pQacM05.String()
	}
	if wI0aaz.JkQ9XFJbp != nil {
		if wI0aaz.F7g5pQacM05 != nil {
			dRoFccaZ += "->"
		} else {
			dRoFccaZ += " "
		}
		dRoFccaZ +=  /*line :1*/ wI0aaz.JkQ9XFJbp.String()
	}
	dRoFccaZ += ": " +  /*line :1*/ wI0aaz.XkWH4CYwNmvD.Error()
	return dRoFccaZ
}

var (
	
	
	eqHMQVuidt	=  /*line :1*/ time.FRXtx9QnU(1, 0)

	
	
	wButnm0q8b	= time.G4KDkI{}
	x0Z4WSkA_	= (chan struct{})( /*line :1*/ nil)
)

type yusr33iR interface {
	Timeout() bool
}

func (wI0aaz *ZOYGdck) Timeout() bool {
	if qC7ririy6a, d30HIwxht := wI0aaz.XkWH4CYwNmvD.(*os.XGyr7J); d30HIwxht {
		aeaqWzxJu, d30HIwxht := qC7ririy6a.V_xDVM.(yusr33iR)
		return d30HIwxht &&  /*line :1*/ aeaqWzxJu.Timeout()
	}
	aeaqWzxJu, d30HIwxht := wI0aaz.XkWH4CYwNmvD.(yusr33iR)
	return d30HIwxht &&  /*line :1*/ aeaqWzxJu.Timeout()
}

type vlKwSU interface {
	Temporary() bool
}

func (wI0aaz *ZOYGdck) Temporary() bool {

	if wI0aaz.SMNyZk_q == "accept" &&  /*line :1*/ suDHtX7CgoG(wI0aaz.XkWH4CYwNmvD) {
		return true
	}

	if qC7ririy6a, d30HIwxht := wI0aaz.XkWH4CYwNmvD.(*os.XGyr7J); d30HIwxht {
		aeaqWzxJu, d30HIwxht := qC7ririy6a.V_xDVM.(vlKwSU)
		return d30HIwxht &&  /*line :1*/ aeaqWzxJu.Temporary()
	}
	aeaqWzxJu, d30HIwxht := wI0aaz.XkWH4CYwNmvD.(vlKwSU)
	return d30HIwxht &&  /*line :1*/ aeaqWzxJu.Temporary()
}


type DPDxTDQNU3Ih struct {
	
	
	FATAgZbL0	string

	
	GrpQb_kayach	string
}

func (wI0aaz *DPDxTDQNU3Ih) Error() string {
	return "invalid " + wI0aaz.FATAgZbL0 + ": " + wI0aaz.GrpQb_kayach
}

func (wI0aaz *DPDxTDQNU3Ih) Timeout() bool	{ return false }
func (wI0aaz *DPDxTDQNU3Ih) Temporary() bool	{ return false }

type DAWLIQHc struct {
	Z68v0y	string
	BCCnAgFxG	string
}

func (wI0aaz *DAWLIQHc) Error() string {
	if wI0aaz == nil {
		return "<nil>"
	}
	dRoFccaZ := wI0aaz.Z68v0y
	if wI0aaz.BCCnAgFxG != "" {
		dRoFccaZ = "address " + wI0aaz.BCCnAgFxG + ": " + dRoFccaZ
	}
	return dRoFccaZ
}

func (wI0aaz *DAWLIQHc) Timeout() bool	{ return false }
func (wI0aaz *DAWLIQHc) Temporary() bool	{ return false }

type JFaUU0ZaU string

func (wI0aaz JFaUU0ZaU) Error() string	{ return "unknown network " +  /*line :1*/ string(wI0aaz) }
func (wI0aaz JFaUU0ZaU) Timeout() bool	{ return false }
func (wI0aaz JFaUU0ZaU) Temporary() bool	{ return false }

type FgXegWYAos string

func (wI0aaz FgXegWYAos) Error() string	{ return  /*line :1*/ string(wI0aaz) }
func (wI0aaz FgXegWYAos) Timeout() bool	{ return false }
func (wI0aaz FgXegWYAos) Temporary() bool	{ return false }












var dklm5juf7rzD error = &fdnPWrllQGA{}

type fdnPWrllQGA struct{}

func (wI0aaz *fdnPWrllQGA) Error() string	{ return "i/o timeout" }
func (wI0aaz *fdnPWrllQGA) Timeout() bool	{ return true }
func (wI0aaz *fdnPWrllQGA) Temporary() bool	{ return true }

func (wI0aaz *fdnPWrllQGA) Is(h_ljk48Bm error) bool {
	return h_ljk48Bm == context.Z201U2N
}



type WdciwD struct {
	EVDlGzGt error
}

func (wI0aaz *WdciwD) Unwrap() error	{ return wI0aaz.EVDlGzGt }
func (wI0aaz *WdciwD) Error() string	{ return "error reading DNS config: " +  /*line :1*/ wI0aaz.EVDlGzGt.Error() }
func (wI0aaz *WdciwD) Timeout() bool	{ return false }
func (wI0aaz *WdciwD) Temporary() bool	{ return false }


var (
	aamCgdkZikV =  /*line :1*/ errors.Su6g6hRoi9X("no such host")
)


type SoatTK0 struct {
	PY3bIIR	string	
	FIvHQdTCAg	string	
	IoaaWnhxbb	string	
	JYeYrD	bool	
	HzqkaqiQE1P	bool	
	GM17caVp1uW	bool	
}

func (wI0aaz *SoatTK0) Error() string {
	if wI0aaz == nil {
		return "<nil>"
	}
	dRoFccaZ := "lookup " + wI0aaz.FIvHQdTCAg
	if wI0aaz.IoaaWnhxbb != "" {
		dRoFccaZ += " on " + wI0aaz.IoaaWnhxbb
	}
	dRoFccaZ += ": " + wI0aaz.PY3bIIR
	return dRoFccaZ
}




func (wI0aaz *SoatTK0) Timeout() bool	{ return wI0aaz.JYeYrD }




func (wI0aaz *SoatTK0) Temporary() bool	{ return wI0aaz.JYeYrD || wI0aaz.HzqkaqiQE1P }



var q7fQtGDv = poll.DA6gHYGx1nT






var AgAUw1g error = q7fQtGDv

type jbAQ44Z_sM struct {
	io.LXQrGW6BQt
}



func gvaVYabrVviG(eCb4vNceh io.LXQrGW6BQt, yxhs4Z io.YJ04Fau) (doauF8Y int64, h_ljk48Bm error) {

	return  /*line :1*/ io.FxikaFo5o7OE(jbAQ44Z_sM{eCb4vNceh}, yxhs4Z)
}

var jr1DA7BWO chan struct{}

var sSxcDq sync.LhBfwn6wa1x

func a0ivyEgPyL() {
	 /*line :1*/ sSxcDq.Do(func() {
		jr1DA7BWO =  /*line :1*/ make(chan struct{},  /*line :1*/ nxi0S_())
	})
	jr1DA7BWO <- struct{}{}
}

func g4VDk4yZiGdR() {
	<-jr1DA7BWO
}





type tRUaxwRBjL interface {
	tT7qFj3HZ(*CqKFIxFUB0B) (int64, error)
}






type CqKFIxFUB0B [][]byte

var (
	_	io.JCVgLD8ld	= (* /*line :1*/ CqKFIxFUB0B)(nil)
	_	io.YJ04Fau	= (* /*line :1*/ CqKFIxFUB0B)(nil)
)







func (ljsCSk *CqKFIxFUB0B) WriteTo(eCb4vNceh io.LXQrGW6BQt) (doauF8Y int64, h_ljk48Bm error) {
	if wolQ8HQucE, d30HIwxht := eCb4vNceh.(tRUaxwRBjL); d30HIwxht {
		return  /*line :1*/ wolQ8HQucE.tT7qFj3HZ(ljsCSk)
	}
	for _, fIadEXIim6l := range *ljsCSk {
		s0ci8yZ, h_ljk48Bm :=  /*line :1*/ eCb4vNceh.Write(fIadEXIim6l)
		doauF8Y +=  /*line :1*/ int64(s0ci8yZ)
		if h_ljk48Bm != nil {
			 /*line :1*/ ljsCSk.ijJn1AKr(doauF8Y)
			return doauF8Y, h_ljk48Bm
		}
	}
	 /*line :1*/ ljsCSk.ijJn1AKr(doauF8Y)
	return doauF8Y, nil
}







func (ljsCSk *CqKFIxFUB0B) Read(fMPVz2iGiyq []byte) (doauF8Y int, h_ljk48Bm error) {
	for  /*line :1*/ len(fMPVz2iGiyq) > 0 &&  /*line :1*/ len(*ljsCSk) > 0 {
		rRU5OziHlL4K :=  /*line :1*/ copy(fMPVz2iGiyq, (*ljsCSk)[0])
		 /*line :1*/ ljsCSk.ijJn1AKr( /*line :1*/ int64(rRU5OziHlL4K))
		fMPVz2iGiyq = fMPVz2iGiyq[rRU5OziHlL4K:]
		doauF8Y += rRU5OziHlL4K
	}
	if  /*line :1*/ len(*ljsCSk) == 0 {
		h_ljk48Bm = io.K5Sqco
	}
	return
}

func (ljsCSk *CqKFIxFUB0B) ijJn1AKr(doauF8Y int64) {
	for  /*line :1*/ len(*ljsCSk) > 0 {
		xEMvkrpC5hOp :=  /*line :1*/ int64( /*line :1*/ len((*ljsCSk)[0]))
		if xEMvkrpC5hOp > doauF8Y {
			(*ljsCSk)[0] = (*ljsCSk)[0][doauF8Y:]
			return
		}
		doauF8Y -= xEMvkrpC5hOp
		(*ljsCSk)[0] = nil
		*ljsCSk = (*ljsCSk)[1:]
	}
}
