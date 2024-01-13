//line :1










package pBPwKYPji2

import (
	fmt "kFsoCfy5zWG"
	io "xI9ai2djaJ2"
	internal "GBulAp"
	os "hPMrClpC"
	"runtime"
	sync "sync"
	atomic "sync/atomic"
	time "fRAfQd_"
)















const (
	Ldate	= 1 << iota		
	Ltime			
	Lmicroseconds			
	Llongfile			
	Lshortfile			
	LUTC			
	Lmsgprefix			
	LstdFlags	= Ldate | Ltime		
)





type ZYYjjiCbq struct {
	hTTpGVo	sync.DIRWe11KYlYa
	mudNH1Cm	io.LXQrGW6BQt	

	qw5ozc	atomic.ToSaNw[string]	
	frL_6XWzIgb8	atomic.JhtCwKEzC	
	b4nWwx	atomic.Akm7q89_
}






func Twasyw(m6kUdIgyY9J7 io.LXQrGW6BQt, a9cVU2ux4fwt string, gxZxZy int) *ZYYjjiCbq {
	gTsZtsit1 :=  /*line :1*/ new(ZYYjjiCbq)
	 /*line :1*/ gTsZtsit1.SetOutput(m6kUdIgyY9J7)
	 /*line :1*/ gTsZtsit1.SetPrefix(a9cVU2ux4fwt)
	 /*line :1*/ gTsZtsit1.SetFlags(gxZxZy)
	return gTsZtsit1
}


func (gTsZtsit1 *ZYYjjiCbq) SetOutput(dbUCy2SV io.LXQrGW6BQt) {
	 /*line :1*/ gTsZtsit1.hTTpGVo.Lock()
	defer  /*line :1*/ gTsZtsit1.hTTpGVo.Unlock()
	gTsZtsit1.mudNH1Cm = dbUCy2SV
	 /*line :1*/ gTsZtsit1.b4nWwx.Store(dbUCy2SV == io.GT3wkcZ)
}

var cc2X7pqv =  /*line :1*/ Twasyw(os.BUPxqQD, "", LstdFlags)


func EqmlC7XUr() *ZYYjjiCbq	{ return cc2X7pqv }


func dsggqiI(geqaxsa5r *[]byte, e6CToXQ int, vaap5T int) {
	
	var ryaRQe6y [20]byte
	xIGRN6XDpJUO :=  /*line :1*/ len(ryaRQe6y) - 1
	for e6CToXQ >= 10 || vaap5T > 1 {
		vaap5T--
		gmK6WX0 := e6CToXQ / 10
		ryaRQe6y[xIGRN6XDpJUO] =  /*line :1*/ byte('0' + e6CToXQ - gmK6WX0*10)
		xIGRN6XDpJUO--
		e6CToXQ = gmK6WX0
	}

	ryaRQe6y[xIGRN6XDpJUO] =  /*line :1*/ byte('0' + e6CToXQ)
	*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, ryaRQe6y[xIGRN6XDpJUO:]...)
}






func hTafJSIG(geqaxsa5r *[]byte, rs0mCawHGqk time.G4KDkI, a9cVU2ux4fwt string, gxZxZy int, cCaP6cibR3a string, yh_QePCHJAz int) {
	if gxZxZy&Lmsgprefix == 0 {
		*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, a9cVU2ux4fwt...)
	}
	if gxZxZy&(Ldate|Ltime|Lmicroseconds) != 0 {
		if gxZxZy&LUTC != 0 {
			rs0mCawHGqk =  /*line :1*/ rs0mCawHGqk.UTC()
		}
		if gxZxZy&Ldate != 0 {
			hxQNI3iR, dfyg0Ziy3E, uAQK5I :=  /*line :1*/ rs0mCawHGqk.Date()
			 /*line :1*/ dsggqiI(geqaxsa5r, hxQNI3iR, 4)
			*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, '/')
			 /*line :1*/ dsggqiI(geqaxsa5r,  /*line :1*/ int(dfyg0Ziy3E), 2)
			*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, '/')
			 /*line :1*/ dsggqiI(geqaxsa5r, uAQK5I, 2)
			*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, ' ')
		}
		if gxZxZy&(Ltime|Lmicroseconds) != 0 {
			eEIa7pEZ5VKk, pnGMy2iM6, ql6t5fTyE6z_ :=  /*line :1*/ rs0mCawHGqk.Clock()
			 /*line :1*/ dsggqiI(geqaxsa5r, eEIa7pEZ5VKk, 2)
			*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, ':')
			 /*line :1*/ dsggqiI(geqaxsa5r, pnGMy2iM6, 2)
			*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, ':')
			 /*line :1*/ dsggqiI(geqaxsa5r, ql6t5fTyE6z_, 2)
			if gxZxZy&Lmicroseconds != 0 {
				*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, '.')
				 /*line :1*/ dsggqiI(geqaxsa5r,  /*line :1*/ rs0mCawHGqk.Nanosecond()/1e3, 6)
			}
			*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, ' ')
		}
	}
	if gxZxZy&(Lshortfile|Llongfile) != 0 {
		if gxZxZy&Lshortfile != 0 {
			o0QKE6N := cCaP6cibR3a
			for e6CToXQ :=  /*line :1*/ len(cCaP6cibR3a) - 1; e6CToXQ > 0; e6CToXQ-- {
				if cCaP6cibR3a[e6CToXQ] == '/' {
					o0QKE6N = cCaP6cibR3a[e6CToXQ+1:]
					break
				}
			}
			cCaP6cibR3a = o0QKE6N
		}
		*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, cCaP6cibR3a...)
		*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, ':')
		 /*line :1*/ dsggqiI(geqaxsa5r, yh_QePCHJAz, -1)
		*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, ": "...)
	}
	if gxZxZy&Lmsgprefix != 0 {
		*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, a9cVU2ux4fwt...)
	}
}

var lpK4UHM20 = sync.OrP5FGPq{IYbTy050ek: func() any { return  /*line :1*/ new([]byte) }}

func db5baXB2esp() *[]byte {
	tYRI6rzm4p :=  /*line :1*/ lpK4UHM20.Get().(*[]byte)
	*tYRI6rzm4p = (*tYRI6rzm4p)[:0]
	return tYRI6rzm4p
}

func nAK5Ufxg(tYRI6rzm4p *[]byte) {

	if  /*line :1*/ cap(*tYRI6rzm4p) > 64<<10 {
		*tYRI6rzm4p = nil
	}
	 /*line :1*/ lpK4UHM20.Put(tYRI6rzm4p)
}







func (gTsZtsit1 *ZYYjjiCbq) Output(gVk2pb1 int, hgVj0RcD string) error {
	gVk2pb1++
	return  /*line :1*/ gTsZtsit1.v8PaKSba3(0, gVk2pb1, func(ryaRQe6y []byte) []byte {
		return  /*line :1*/ append(ryaRQe6y, hgVj0RcD...)
	})
}



func (gTsZtsit1 *ZYYjjiCbq) v8PaKSba3(nQGX1g4Ff_TX uintptr, gVk2pb1 int, sUqDvARzA func([]byte) []byte) error {
	if  /*line :1*/ gTsZtsit1.b4nWwx.Load() {
		return nil
	}

	tgN7fU :=  /*line :1*/ time.Pc_35oTY()

	a9cVU2ux4fwt :=  /*line :1*/ gTsZtsit1.Prefix()
	gxZxZy :=  /*line :1*/ gTsZtsit1.Flags()

	var cCaP6cibR3a string
	var yh_QePCHJAz int
	if gxZxZy&(Lshortfile|Llongfile) != 0 {
		if nQGX1g4Ff_TX == 0 {
			var lHaoeHudjFvh bool
			_, cCaP6cibR3a, yh_QePCHJAz, lHaoeHudjFvh =  /*line :1*/ runtime.Caller(gVk2pb1)
			if !lHaoeHudjFvh {
				cCaP6cibR3a = "???"
				yh_QePCHJAz = 0
			}
		} else {
			cZaT166 :=  /*line :1*/ runtime.CallersFrames([]uintptr{nQGX1g4Ff_TX})
			iy94QFlR, _ :=  /*line :1*/ cZaT166.Next()
			cCaP6cibR3a = iy94QFlR.File
			if cCaP6cibR3a == "" {
				cCaP6cibR3a = "???"
			}
			yh_QePCHJAz = iy94QFlR.Line
		}
	}

	geqaxsa5r :=  /*line :1*/ db5baXB2esp()
	defer  /*line :1*/ nAK5Ufxg(geqaxsa5r)
	 /*line :1*/ hTafJSIG(geqaxsa5r, tgN7fU, a9cVU2ux4fwt, gxZxZy, cCaP6cibR3a, yh_QePCHJAz)
	*geqaxsa5r =  /*line :1*/ sUqDvARzA(*geqaxsa5r)
	if  /*line :1*/ len(*geqaxsa5r) == 0 || (*geqaxsa5r)[ /*line :1*/ len(*geqaxsa5r)-1] != '\n' {
		*geqaxsa5r =  /*line :1*/ append(*geqaxsa5r, '\n')
	}

	 /*line :1*/ gTsZtsit1.hTTpGVo.Lock()
	defer  /*line :1*/ gTsZtsit1.hTTpGVo.Unlock()
	_, skMUjEzBuv7o :=  /*line :1*/ gTsZtsit1.mudNH1Cm.Write(*geqaxsa5r)
	return skMUjEzBuv7o
}

func init() {
	internal.D9Y9pRseR = func(nQGX1g4Ff_TX uintptr, pdaXJhIPn6 []byte) error {
		return  /*line :1*/ cc2X7pqv.v8PaKSba3(nQGX1g4Ff_TX, 0, func(geqaxsa5r []byte) []byte {
			return  /*line :1*/ append(geqaxsa5r, pdaXJhIPn6...)
		})
	}
}



func (gTsZtsit1 *ZYYjjiCbq) Print(vNIW4ZJEE ...any) {
	 /*line :1*/ gTsZtsit1.v8PaKSba3(0, 2, func(ryaRQe6y []byte) []byte {
		return  /*line :1*/ fmt.B2FMEeakk(ryaRQe6y, vNIW4ZJEE...)
	})
}



func (gTsZtsit1 *ZYYjjiCbq) Printf(o2xrurpka3sS string, vNIW4ZJEE ...any) {
	 /*line :1*/ gTsZtsit1.v8PaKSba3(0, 2, func(ryaRQe6y []byte) []byte {
		return  /*line :1*/ fmt.Pu4AZZdhg(ryaRQe6y, o2xrurpka3sS, vNIW4ZJEE...)
	})
}



func (gTsZtsit1 *ZYYjjiCbq) Println(vNIW4ZJEE ...any) {
	 /*line :1*/ gTsZtsit1.v8PaKSba3(0, 2, func(ryaRQe6y []byte) []byte {
		return  /*line :1*/ fmt.SG1i1athxf(ryaRQe6y, vNIW4ZJEE...)
	})
}


func (gTsZtsit1 *ZYYjjiCbq) Fatal(vNIW4ZJEE ...any) {
	 /*line :1*/ gTsZtsit1.Output(2,  /*line :1*/ fmt.GTkL3_(vNIW4ZJEE...))
	 /*line :1*/ os.U9KBmNojTul(1)
}


func (gTsZtsit1 *ZYYjjiCbq) Fatalf(o2xrurpka3sS string, vNIW4ZJEE ...any) {
	 /*line :1*/ gTsZtsit1.Output(2,  /*line :1*/ fmt.IBsS3Oc(o2xrurpka3sS, vNIW4ZJEE...))
	 /*line :1*/ os.U9KBmNojTul(1)
}


func (gTsZtsit1 *ZYYjjiCbq) Fatalln(vNIW4ZJEE ...any) {
	 /*line :1*/ gTsZtsit1.Output(2,  /*line :1*/ fmt.HOfbYPr(vNIW4ZJEE...))
	 /*line :1*/ os.U9KBmNojTul(1)
}


func (gTsZtsit1 *ZYYjjiCbq) Panic(vNIW4ZJEE ...any) {
	hgVj0RcD :=  /*line :1*/ fmt.GTkL3_(vNIW4ZJEE...)
	 /*line :1*/ gTsZtsit1.Output(2, hgVj0RcD)
	 /*line :1*/ panic(hgVj0RcD)
}


func (gTsZtsit1 *ZYYjjiCbq) Panicf(o2xrurpka3sS string, vNIW4ZJEE ...any) {
	hgVj0RcD :=  /*line :1*/ fmt.IBsS3Oc(o2xrurpka3sS, vNIW4ZJEE...)
	 /*line :1*/ gTsZtsit1.Output(2, hgVj0RcD)
	 /*line :1*/ panic(hgVj0RcD)
}


func (gTsZtsit1 *ZYYjjiCbq) Panicln(vNIW4ZJEE ...any) {
	hgVj0RcD :=  /*line :1*/ fmt.HOfbYPr(vNIW4ZJEE...)
	 /*line :1*/ gTsZtsit1.Output(2, hgVj0RcD)
	 /*line :1*/ panic(hgVj0RcD)
}



func (gTsZtsit1 *ZYYjjiCbq) Flags() int {
	return  /*line :1*/ int( /*line :1*/ gTsZtsit1.frL_6XWzIgb8.Load())
}



func (gTsZtsit1 *ZYYjjiCbq) SetFlags(gxZxZy int) {
	 /*line :1*/ gTsZtsit1.frL_6XWzIgb8.Store( /*line :1*/ int32(gxZxZy))
}


func (gTsZtsit1 *ZYYjjiCbq) Prefix() string {
	if tYRI6rzm4p :=  /*line :1*/ gTsZtsit1.qw5ozc.Load(); tYRI6rzm4p != nil {
		return *tYRI6rzm4p
	}
	return ""
}


func (gTsZtsit1 *ZYYjjiCbq) SetPrefix(a9cVU2ux4fwt string) {
	 /*line :1*/ gTsZtsit1.qw5ozc.Store(&a9cVU2ux4fwt)
}


func (gTsZtsit1 *ZYYjjiCbq) Writer() io.LXQrGW6BQt {
	 /*line :1*/ gTsZtsit1.hTTpGVo.Lock()
	defer  /*line :1*/ gTsZtsit1.hTTpGVo.Unlock()
	return gTsZtsit1.mudNH1Cm
}


func BhaTaFixN1j(dbUCy2SV io.LXQrGW6BQt) {
	 /*line :1*/ cc2X7pqv.SetOutput(dbUCy2SV)
}



func MEY3zpL3v5() int {
	return  /*line :1*/ cc2X7pqv.Flags()
}



func HartPWo(gxZxZy int) {
	 /*line :1*/ cc2X7pqv.SetFlags(gxZxZy)
}


func TxY_uBFO() string {
	return  /*line :1*/ cc2X7pqv.Prefix()
}


func ZpgXDPEuK80(a9cVU2ux4fwt string) {
	 /*line :1*/ cc2X7pqv.SetPrefix(a9cVU2ux4fwt)
}


func AuuhZS7s2() io.LXQrGW6BQt {
	return  /*line :1*/ cc2X7pqv.Writer()
}



func B1rCMabZ(vNIW4ZJEE ...any) {
	 /*line :1*/ cc2X7pqv.v8PaKSba3(0, 2, func(ryaRQe6y []byte) []byte {
		return  /*line :1*/ fmt.B2FMEeakk(ryaRQe6y, vNIW4ZJEE...)
	})
}



func F59JWa(o2xrurpka3sS string, vNIW4ZJEE ...any) {
	 /*line :1*/ cc2X7pqv.v8PaKSba3(0, 2, func(ryaRQe6y []byte) []byte {
		return  /*line :1*/ fmt.Pu4AZZdhg(ryaRQe6y, o2xrurpka3sS, vNIW4ZJEE...)
	})
}



func JbAQ8b(vNIW4ZJEE ...any) {
	 /*line :1*/ cc2X7pqv.v8PaKSba3(0, 2, func(ryaRQe6y []byte) []byte {
		return  /*line :1*/ fmt.SG1i1athxf(ryaRQe6y, vNIW4ZJEE...)
	})
}


func UzYjQFBe(vNIW4ZJEE ...any) {
	 /*line :1*/ cc2X7pqv.Output(2,  /*line :1*/ fmt.GTkL3_(vNIW4ZJEE...))
	 /*line :1*/ os.U9KBmNojTul(1)
}


func EPZ3R9nMqGY(o2xrurpka3sS string, vNIW4ZJEE ...any) {
	 /*line :1*/ cc2X7pqv.Output(2,  /*line :1*/ fmt.IBsS3Oc(o2xrurpka3sS, vNIW4ZJEE...))
	 /*line :1*/ os.U9KBmNojTul(1)
}


func ST4nUuPEj5(vNIW4ZJEE ...any) {
	 /*line :1*/ cc2X7pqv.Output(2,  /*line :1*/ fmt.HOfbYPr(vNIW4ZJEE...))
	 /*line :1*/ os.U9KBmNojTul(1)
}


func OQHyWSLcRm1(vNIW4ZJEE ...any) {
	hgVj0RcD :=  /*line :1*/ fmt.GTkL3_(vNIW4ZJEE...)
	 /*line :1*/ cc2X7pqv.Output(2, hgVj0RcD)
	 /*line :1*/ panic(hgVj0RcD)
}


func YjA4ZlqL_(o2xrurpka3sS string, vNIW4ZJEE ...any) {
	hgVj0RcD :=  /*line :1*/ fmt.IBsS3Oc(o2xrurpka3sS, vNIW4ZJEE...)
	 /*line :1*/ cc2X7pqv.Output(2, hgVj0RcD)
	 /*line :1*/ panic(hgVj0RcD)
}


func BudNaHv2ykwa(vNIW4ZJEE ...any) {
	hgVj0RcD :=  /*line :1*/ fmt.HOfbYPr(vNIW4ZJEE...)
	 /*line :1*/ cc2X7pqv.Output(2, hgVj0RcD)
	 /*line :1*/ panic(hgVj0RcD)
}








func IgjuArATy(gVk2pb1 int, hgVj0RcD string) error {
	return  /*line :1*/ cc2X7pqv.Output(gVk2pb1+1, hgVj0RcD)
}
