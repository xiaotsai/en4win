//line :1
package kFsoCfy5zWG

import (
	fmtsort "sXttzp"
	io "xI9ai2djaJ2"
	os "hPMrClpC"
	reflect "reflect"
	strconv "zBESu2SrRjP"
	sync "sync"
	utf8 "CuAc0pPZwf"
)



const (
	commaSpaceString	= ", "
	nilAngleString	= "<nil>"
	nilParenString	= "(nil)"
	nilString	= "nil"
	mapString	= "map["
	percentBangString	= "%!"
	missingString	= "(MISSING)"
	badIndexString	= "(BADINDEX)"
	panicString	= "(PANIC="
	extraString	= "%!(EXTRA "
	badWidthString	= "%!(BADWIDTH)"
	badPrecString	= "%!(BADPREC)"
	noVerbString	= "%!(NOVERB)"
	invReflectString	= "<invalid reflect.Value>"
)




type Bm1Jqy8i interface {
	
	Write(dajeEj8 []byte) (bMu41ab1Tf int, jDbNTz1lC_ error)
	
	Width() (xgwcYhwzhN int, xP5nLuM9_y bool)
	
	Precision() (cqsxu8pzP int, xP5nLuM9_y bool)

	
	Flag(gaq8WME int) bool
}




type OpUJYyZOy8 interface {
	Format(zEHZ5ak0_jf Bm1Jqy8i, mEbw55xKIj4H rune)
}






type Q7gRl6EJNyL interface {
	String() string
}





type LzAQey interface {
	GoString() string
}







func BGWdZ5eL(l7qaxfr Bm1Jqy8i, mEbw55xKIj4H rune) string {
	var g5TokkbvQblm [16]byte	
	dajeEj8 :=  /*line :1*/ append(g5TokkbvQblm[:0], '%')
	for _, gaq8WME := range " +-#0" {
		if  /*line :1*/ l7qaxfr.Flag( /*line :1*/ int(gaq8WME)) {
			dajeEj8 =  /*line :1*/ append(dajeEj8,  /*line :1*/ byte(gaq8WME))
		}
	}
	if dwIESiwPi, xP5nLuM9_y :=  /*line :1*/ l7qaxfr.Width(); xP5nLuM9_y {
		dajeEj8 =  /*line :1*/ strconv.YOFr9xxiI(dajeEj8,  /*line :1*/ int64(dwIESiwPi), 10)
	}
	if hphPuX8VWjn5, xP5nLuM9_y :=  /*line :1*/ l7qaxfr.Precision(); xP5nLuM9_y {
		dajeEj8 =  /*line :1*/ append(dajeEj8, '.')
		dajeEj8 =  /*line :1*/ strconv.YOFr9xxiI(dajeEj8,  /*line :1*/ int64(hphPuX8VWjn5), 10)
	}
	dajeEj8 =  /*line :1*/ utf8.Ht7oMzd(dajeEj8, mEbw55xKIj4H)
	return  /*line :1*/ string(dajeEj8)
}


type apmFUoDP []byte

func (dajeEj8 *apmFUoDP) qz29wCM_6(hphPuX8VWjn5 []byte) {
	*dajeEj8 =  /*line :1*/ append(*dajeEj8, hphPuX8VWjn5...)
}

func (dajeEj8 *apmFUoDP) aKy73W4FDLtk(a_CBqX string) {
	*dajeEj8 =  /*line :1*/ append(*dajeEj8, a_CBqX...)
}

func (dajeEj8 *apmFUoDP) yYk12QG(gaq8WME byte) {
	*dajeEj8 =  /*line :1*/ append(*dajeEj8, gaq8WME)
}

func (zKM2p6ma *apmFUoDP) j9YRMbR(vuT9drijWK rune) {
	*zKM2p6ma =  /*line :1*/ utf8.Ht7oMzd(*zKM2p6ma, vuT9drijWK)
}


type yGfDUYjU3 struct {
	ocZNWYPb9c	apmFUoDP

	
	geZli9XXQH	any

	
	ky_TJjz7balj	reflect.Value

	
	hkpuiPFUGfe	kFsoCfy5zWG

	
	j7iqdfrBQJ	bool
	
	gTUxvvaG7R	bool
	
	pBwiU8	bool
	
	i6D6JX	bool
	
	ttXdEVc1s50Z	bool
	
	f3sXqNFocRH1	[]int
}

var jN2gSCwJ = sync.OrP5FGPq{
	IYbTy050ek: func() any { return  /*line :1*/ new(yGfDUYjU3) },
}


func rginCno6N0() *yGfDUYjU3 {
	hphPuX8VWjn5 :=  /*line :1*/ jN2gSCwJ.Get().(*yGfDUYjU3)
	hphPuX8VWjn5.pBwiU8 = false
	hphPuX8VWjn5.i6D6JX = false
	hphPuX8VWjn5.ttXdEVc1s50Z = false
	 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.init(&hphPuX8VWjn5.ocZNWYPb9c)
	return hphPuX8VWjn5
}


func (hphPuX8VWjn5 *yGfDUYjU3) aUrCCiO() {

	if  /*line :1*/ cap(hphPuX8VWjn5.ocZNWYPb9c) > 64*1024 {
		hphPuX8VWjn5.ocZNWYPb9c = nil
	} else {
		hphPuX8VWjn5.ocZNWYPb9c = hphPuX8VWjn5.ocZNWYPb9c[:0]
	}
	if  /*line :1*/ cap(hphPuX8VWjn5.f3sXqNFocRH1) > 8 {
		hphPuX8VWjn5.f3sXqNFocRH1 = nil
	}

	hphPuX8VWjn5.geZli9XXQH = nil
	hphPuX8VWjn5.ky_TJjz7balj = reflect.Value{}
	hphPuX8VWjn5.f3sXqNFocRH1 = hphPuX8VWjn5.f3sXqNFocRH1[:0]
	 /*line :1*/ jN2gSCwJ.Put(hphPuX8VWjn5)
}

func (hphPuX8VWjn5 *yGfDUYjU3) Width() (xgwcYhwzhN int, xP5nLuM9_y bool) {
	return hphPuX8VWjn5.hkpuiPFUGfe.zToz50Xry, hphPuX8VWjn5.hkpuiPFUGfe.fMbhCq
}

func (hphPuX8VWjn5 *yGfDUYjU3) Precision() (cqsxu8pzP int, xP5nLuM9_y bool) {
	return hphPuX8VWjn5.hkpuiPFUGfe.f05l8Dt, hphPuX8VWjn5.hkpuiPFUGfe.dD4Xiyfl
}

func (hphPuX8VWjn5 *yGfDUYjU3) Flag(dajeEj8 int) bool {
	switch dajeEj8 {
	case '-':
		return hphPuX8VWjn5.hkpuiPFUGfe.i5XVr2y
	case '+':
		return hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G || hphPuX8VWjn5.hkpuiPFUGfe.j7W9PFs
	case '#':
		return hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN || hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u
	case ' ':
		return hphPuX8VWjn5.hkpuiPFUGfe.hDpt7tehWk
	case '0':
		return hphPuX8VWjn5.hkpuiPFUGfe.tpbzso
	}
	return false
}



func (hphPuX8VWjn5 *yGfDUYjU3) Write(dajeEj8 []byte) (bTl364i int, jDbNTz1lC_ error) {
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.qz29wCM_6(dajeEj8)
	return  /*line :1*/ len(dajeEj8), nil
}



func (hphPuX8VWjn5 *yGfDUYjU3) WriteString(a_CBqX string) (bTl364i int, jDbNTz1lC_ error) {
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(a_CBqX)
	return  /*line :1*/ len(a_CBqX), nil
}



func BYcL2whVEYz(dwIESiwPi io.LXQrGW6BQt, aCNXxCRXOS5o string, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.wqxOmPyRm4(aCNXxCRXOS5o, nghaNY)
	bMu41ab1Tf, jDbNTz1lC_ =  /*line :1*/ dwIESiwPi.Write(hphPuX8VWjn5.ocZNWYPb9c)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return
}



func PUFyrL81XY(aCNXxCRXOS5o string, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ BYcL2whVEYz(os.Vrq37o, aCNXxCRXOS5o, nghaNY...)
}


func IBsS3Oc(aCNXxCRXOS5o string, nghaNY ...any) string {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.wqxOmPyRm4(aCNXxCRXOS5o, nghaNY)
	a_CBqX :=  /*line :1*/ string(hphPuX8VWjn5.ocZNWYPb9c)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return a_CBqX
}



func Pu4AZZdhg(dajeEj8 []byte, aCNXxCRXOS5o string, nghaNY ...any) []byte {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.wqxOmPyRm4(aCNXxCRXOS5o, nghaNY)
	dajeEj8 =  /*line :1*/ append(dajeEj8, hphPuX8VWjn5.ocZNWYPb9c...)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return dajeEj8
}




func AJRwHmjcphiN(dwIESiwPi io.LXQrGW6BQt, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.yf4VotSF2G_I(nghaNY)
	bMu41ab1Tf, jDbNTz1lC_ =  /*line :1*/ dwIESiwPi.Write(hphPuX8VWjn5.ocZNWYPb9c)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return
}




func Mg4aCbWR(nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ AJRwHmjcphiN(os.Vrq37o, nghaNY...)
}



func GTkL3_(nghaNY ...any) string {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.yf4VotSF2G_I(nghaNY)
	a_CBqX :=  /*line :1*/ string(hphPuX8VWjn5.ocZNWYPb9c)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return a_CBqX
}



func B2FMEeakk(dajeEj8 []byte, nghaNY ...any) []byte {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.yf4VotSF2G_I(nghaNY)
	dajeEj8 =  /*line :1*/ append(dajeEj8, hphPuX8VWjn5.ocZNWYPb9c...)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return dajeEj8
}




func SukeLvKn(dwIESiwPi io.LXQrGW6BQt, nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.zrWkfne7UM(nghaNY)
	bMu41ab1Tf, jDbNTz1lC_ =  /*line :1*/ dwIESiwPi.Write(hphPuX8VWjn5.ocZNWYPb9c)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return
}




func IrpedaD0v(nghaNY ...any) (bMu41ab1Tf int, jDbNTz1lC_ error) {
	return  /*line :1*/ SukeLvKn(os.Vrq37o, nghaNY...)
}



func HOfbYPr(nghaNY ...any) string {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.zrWkfne7UM(nghaNY)
	a_CBqX :=  /*line :1*/ string(hphPuX8VWjn5.ocZNWYPb9c)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return a_CBqX
}




func SG1i1athxf(dajeEj8 []byte, nghaNY ...any) []byte {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	 /*line :1*/ hphPuX8VWjn5.zrWkfne7UM(nghaNY)
	dajeEj8 =  /*line :1*/ append(dajeEj8, hphPuX8VWjn5.ocZNWYPb9c...)
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return dajeEj8
}




func lzLfy8wBveLs(xsKqMyFqj reflect.Value, wKUgq0A int) reflect.Value {
	cZesrlF9C :=  /*line :1*/ xsKqMyFqj.Field(wKUgq0A)
	if  /*line :1*/ cZesrlF9C.Kind() == reflect.Interface && ! /*line :1*/ cZesrlF9C.IsNil() {
		cZesrlF9C =  /*line :1*/ cZesrlF9C.Elem()
	}
	return cZesrlF9C
}



func eIJkRuoT(vFAbma1 int) bool {
	const max int = 1e6
	return vFAbma1 > max || vFAbma1 < -max
}


func bWbu9qxG4bF(a_CBqX string, krzBR_fjCZmC, dA0LK4aPCOv int) (ez4abKGFo int, _QYiNEUWI9m bool, nHl8Y8ilB_I int) {
	if krzBR_fjCZmC >= dA0LK4aPCOv {
		return 0, false, dA0LK4aPCOv
	}
	for nHl8Y8ilB_I = krzBR_fjCZmC; nHl8Y8ilB_I < dA0LK4aPCOv && '0' <= a_CBqX[nHl8Y8ilB_I] && a_CBqX[nHl8Y8ilB_I] <= '9'; nHl8Y8ilB_I++ {
		if  /*line :1*/ eIJkRuoT(ez4abKGFo) {
			return 0, false, dA0LK4aPCOv
		}
		ez4abKGFo = ez4abKGFo*10 +  /*line :1*/ int(a_CBqX[nHl8Y8ilB_I]-'0')
		_QYiNEUWI9m = true
	}
	return
}

func (hphPuX8VWjn5 *yGfDUYjU3) zSah4UJOc(xsKqMyFqj reflect.Value) {
	if ! /*line :1*/ xsKqMyFqj.IsValid() {
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilAngleString)
		return
	}
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('?')
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ xsKqMyFqj.Type().String())
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('?')
}

func (hphPuX8VWjn5 *yGfDUYjU3) nb_dpqcH0(mEbw55xKIj4H rune) {
	hphPuX8VWjn5.i6D6JX = true
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(percentBangString)
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.j9YRMbR(mEbw55xKIj4H)
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('(')
	switch {
	case hphPuX8VWjn5.geZli9XXQH != nil:
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ reflect.Cher1a2Fblr(hphPuX8VWjn5.geZli9XXQH).String())
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('=')
		 /*line :1*/ hphPuX8VWjn5.y383gkr_(hphPuX8VWjn5.geZli9XXQH, 'v')
	case  /*line :1*/ hphPuX8VWjn5.ky_TJjz7balj.IsValid():
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ hphPuX8VWjn5.ky_TJjz7balj.Type().String())
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('=')
		 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy(hphPuX8VWjn5.ky_TJjz7balj, 'v', 0)
	default:
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilAngleString)
	}
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(')')
	hphPuX8VWjn5.i6D6JX = false
}

func (hphPuX8VWjn5 *yGfDUYjU3) ap7H2Bk_xWi(xsKqMyFqj bool, mEbw55xKIj4H rune) {
	switch mEbw55xKIj4H {
	case 't', 'v':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.zKKG0xo(xsKqMyFqj)
	default:
		 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
	}
}



func (hphPuX8VWjn5 *yGfDUYjU3) oCt9wWdj_(xsKqMyFqj uint64, ygYjblbQZfz bool) {
	jUJvRAH := hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN
	hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN = ygYjblbQZfz
	 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ(xsKqMyFqj, 16, unsigned, 'v', ldigits)
	hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN = jUJvRAH
}


func (hphPuX8VWjn5 *yGfDUYjU3) bnoiaBZ(xsKqMyFqj uint64, pWv6z6Y2 bool, mEbw55xKIj4H rune) {
	switch mEbw55xKIj4H {
	case 'v':
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u && !pWv6z6Y2 {
			 /*line :1*/ hphPuX8VWjn5.oCt9wWdj_(xsKqMyFqj, true)
		} else {
			 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ(xsKqMyFqj, 10, pWv6z6Y2, mEbw55xKIj4H, ldigits)
		}
	case 'd':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ(xsKqMyFqj, 10, pWv6z6Y2, mEbw55xKIj4H, ldigits)
	case 'b':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ(xsKqMyFqj, 2, pWv6z6Y2, mEbw55xKIj4H, ldigits)
	case 'o', 'O':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ(xsKqMyFqj, 8, pWv6z6Y2, mEbw55xKIj4H, ldigits)
	case 'x':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ(xsKqMyFqj, 16, pWv6z6Y2, mEbw55xKIj4H, ldigits)
	case 'X':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ(xsKqMyFqj, 16, pWv6z6Y2, mEbw55xKIj4H, udigits)
	case 'c':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.iCNDaD7WCua(xsKqMyFqj)
	case 'q':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.vlMdvVOEGU(xsKqMyFqj)
	case 'U':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.wCQMdGA(xsKqMyFqj)
	default:
		 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
	}
}



func (hphPuX8VWjn5 *yGfDUYjU3) rwKBPG3A3fA(xsKqMyFqj float64, e2v7yTpR2 int, mEbw55xKIj4H rune) {
	switch mEbw55xKIj4H {
	case 'v':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.rwKBPG3A3fA(xsKqMyFqj, e2v7yTpR2, 'g', -1)
	case 'b', 'g', 'G', 'x', 'X':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.rwKBPG3A3fA(xsKqMyFqj, e2v7yTpR2, mEbw55xKIj4H, -1)
	case 'f', 'e', 'E':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.rwKBPG3A3fA(xsKqMyFqj, e2v7yTpR2, mEbw55xKIj4H, 6)
	case 'F':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.rwKBPG3A3fA(xsKqMyFqj, e2v7yTpR2, 'f', 6)
	default:
		 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
	}
}




func (hphPuX8VWjn5 *yGfDUYjU3) v3XHxGJFrN(xsKqMyFqj complex128, e2v7yTpR2 int, mEbw55xKIj4H rune) {

	switch mEbw55xKIj4H {
	case 'v', 'b', 'g', 'G', 'x', 'X', 'f', 'F', 'e', 'E':
		chYuF5rE1O3t := hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('(')
		 /*line :1*/ hphPuX8VWjn5.rwKBPG3A3fA( /*line :1*/ real(xsKqMyFqj), e2v7yTpR2/2, mEbw55xKIj4H)

		hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G = true
		 /*line :1*/ hphPuX8VWjn5.rwKBPG3A3fA( /*line :1*/ imag(xsKqMyFqj), e2v7yTpR2/2, mEbw55xKIj4H)
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk("i)")
		hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G = chYuF5rE1O3t
	default:
		 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
	}
}

func (hphPuX8VWjn5 *yGfDUYjU3) yoxXrK6J(xsKqMyFqj string, mEbw55xKIj4H rune) {
	switch mEbw55xKIj4H {
	case 'v':
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
			 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.zbSAsGc(xsKqMyFqj)
		} else {
			 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.jd6DylJG8Tf(xsKqMyFqj)
		}
	case 's':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.jd6DylJG8Tf(xsKqMyFqj)
	case 'x':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.n4JaSpX_a(xsKqMyFqj, ldigits)
	case 'X':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.n4JaSpX_a(xsKqMyFqj, udigits)
	case 'q':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.zbSAsGc(xsKqMyFqj)
	default:
		 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
	}
}

func (hphPuX8VWjn5 *yGfDUYjU3) apkzAM(xsKqMyFqj []byte, mEbw55xKIj4H rune, oqVBjAx7Eks string) {
	switch mEbw55xKIj4H {
	case 'v', 'd':
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(oqVBjAx7Eks)
			if xsKqMyFqj == nil {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilParenString)
				return
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('{')
			for wKUgq0A, gaq8WME := range xsKqMyFqj {
				if wKUgq0A > 0 {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(commaSpaceString)
				}
				 /*line :1*/ hphPuX8VWjn5.oCt9wWdj_( /*line :1*/ uint64(gaq8WME), true)
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('}')
		} else {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('[')
			for wKUgq0A, gaq8WME := range xsKqMyFqj {
				if wKUgq0A > 0 {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(' ')
				}
				 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.bnoiaBZ( /*line :1*/ uint64(gaq8WME), 10, unsigned, mEbw55xKIj4H, ldigits)
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(']')
		}
	case 's':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.etdyui(xsKqMyFqj)
	case 'x':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.kAP1zz(xsKqMyFqj, ldigits)
	case 'X':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.kAP1zz(xsKqMyFqj, udigits)
	case 'q':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.zbSAsGc( /*line :1*/ string(xsKqMyFqj))
	default:
		 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy( /*line :1*/ reflect.SdHoaIQl(xsKqMyFqj), mEbw55xKIj4H, 0)
	}
}

func (hphPuX8VWjn5 *yGfDUYjU3) sja4wY(uajxUEw reflect.Value, mEbw55xKIj4H rune) {
	var mT6BIv4qJ uintptr
	switch  /*line :1*/ uajxUEw.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		mT6BIv4qJ =  /*line :1*/ uajxUEw.Pointer()
	default:
		 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
		return
	}

	switch mEbw55xKIj4H {
	case 'v':
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('(')
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ uajxUEw.Type().String())
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(")(")
			if mT6BIv4qJ == 0 {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilString)
			} else {
				 /*line :1*/ hphPuX8VWjn5.oCt9wWdj_( /*line :1*/ uint64(mT6BIv4qJ), true)
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(')')
		} else {
			if mT6BIv4qJ == 0 {
				 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.ii45_K(nilAngleString)
			} else {
				 /*line :1*/ hphPuX8VWjn5.oCt9wWdj_( /*line :1*/ uint64(mT6BIv4qJ), !hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN)
			}
		}
	case 'p':
		 /*line :1*/ hphPuX8VWjn5.oCt9wWdj_( /*line :1*/ uint64(mT6BIv4qJ), !hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN)
	case 'b', 'o', 'd', 'x', 'X':
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(mT6BIv4qJ), unsigned, mEbw55xKIj4H)
	default:
		 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
	}
}

func (hphPuX8VWjn5 *yGfDUYjU3) hfKKWu(oaT9vLaMI3 any, mEbw55xKIj4H rune, ulncweMEP string) {
	if jDbNTz1lC_ :=  /*line :1*/ recover(); jDbNTz1lC_ != nil {

		if xsKqMyFqj :=  /*line :1*/ reflect.SdHoaIQl(oaT9vLaMI3);  /*line :1*/ xsKqMyFqj.Kind() == reflect.Pointer &&  /*line :1*/ xsKqMyFqj.IsNil() {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilAngleString)
			return
		}

		if hphPuX8VWjn5.pBwiU8 {

			 /*line :1*/ panic(jDbNTz1lC_)
		}

		_YyVEw6UW := hphPuX8VWjn5.hkpuiPFUGfe.bXSDcaMZCgR

		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.xdIs6xWSBJJX()

		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(percentBangString)
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.j9YRMbR(mEbw55xKIj4H)
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(panicString)
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(ulncweMEP)
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(" method: ")
		hphPuX8VWjn5.pBwiU8 = true
		 /*line :1*/ hphPuX8VWjn5.y383gkr_(jDbNTz1lC_, 'v')
		hphPuX8VWjn5.pBwiU8 = false
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(')')

		hphPuX8VWjn5.hkpuiPFUGfe.bXSDcaMZCgR = _YyVEw6UW
	}
}

func (hphPuX8VWjn5 *yGfDUYjU3) lz24sYwIhG(mEbw55xKIj4H rune) (dBMxa0 bool) {
	if hphPuX8VWjn5.i6D6JX {
		return
	}
	if mEbw55xKIj4H == 'w' {

		_, xP5nLuM9_y := hphPuX8VWjn5.geZli9XXQH.(error)
		if !xP5nLuM9_y || !hphPuX8VWjn5.ttXdEVc1s50Z {
			 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
			return true
		}

		mEbw55xKIj4H = 'v'
	}

	if u0ZcCSN, xP5nLuM9_y := hphPuX8VWjn5.geZli9XXQH.(OpUJYyZOy8); xP5nLuM9_y {
		dBMxa0 = true
		defer  /*line :1*/ hphPuX8VWjn5.hfKKWu(hphPuX8VWjn5.geZli9XXQH, mEbw55xKIj4H, "Format")
		 /*line :1*/ u0ZcCSN.Format(hphPuX8VWjn5, mEbw55xKIj4H)
		return
	}

	if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
		if meaTF_q7lBWk, xP5nLuM9_y := hphPuX8VWjn5.geZli9XXQH.(LzAQey); xP5nLuM9_y {
			dBMxa0 = true
			defer  /*line :1*/ hphPuX8VWjn5.hfKKWu(hphPuX8VWjn5.geZli9XXQH, mEbw55xKIj4H, "GoString")

			 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.jd6DylJG8Tf( /*line :1*/ meaTF_q7lBWk.GoString())
			return
		}
	} else {

		switch mEbw55xKIj4H {
		case 'v', 's', 'x', 'X', 'q':

			switch xsKqMyFqj := hphPuX8VWjn5.geZli9XXQH.(type) {
			case error:
				dBMxa0 = true
				defer  /*line :1*/ hphPuX8VWjn5.hfKKWu(hphPuX8VWjn5.geZli9XXQH, mEbw55xKIj4H, "Error")
				 /*line :1*/ hphPuX8VWjn5.yoxXrK6J( /*line :1*/ xsKqMyFqj.Error(), mEbw55xKIj4H)
				return

			case Q7gRl6EJNyL:
				dBMxa0 = true
				defer  /*line :1*/ hphPuX8VWjn5.hfKKWu(hphPuX8VWjn5.geZli9XXQH, mEbw55xKIj4H, "String")
				 /*line :1*/ hphPuX8VWjn5.yoxXrK6J( /*line :1*/ xsKqMyFqj.String(), mEbw55xKIj4H)
				return
			}
		}
	}
	return false
}

func (hphPuX8VWjn5 *yGfDUYjU3) y383gkr_(oaT9vLaMI3 any, mEbw55xKIj4H rune) {
	hphPuX8VWjn5.geZli9XXQH = oaT9vLaMI3
	hphPuX8VWjn5.ky_TJjz7balj = reflect.Value{}

	if oaT9vLaMI3 == nil {
		switch mEbw55xKIj4H {
		case 'T', 'v':
			 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.ii45_K(nilAngleString)
		default:
			 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
		}
		return
	}

	switch mEbw55xKIj4H {
	case 'T':
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.jd6DylJG8Tf( /*line :1*/ reflect.Cher1a2Fblr(oaT9vLaMI3).String())
		return
	case 'p':
		 /*line :1*/ hphPuX8VWjn5.sja4wY( /*line :1*/ reflect.SdHoaIQl(oaT9vLaMI3), 'p')
		return
	}

	switch zEHZ5ak0_jf := oaT9vLaMI3.(type) {
	case bool:
		 /*line :1*/ hphPuX8VWjn5.ap7H2Bk_xWi(zEHZ5ak0_jf, mEbw55xKIj4H)
	case float32:
		 /*line :1*/ hphPuX8VWjn5.rwKBPG3A3fA( /*line :1*/ float64(zEHZ5ak0_jf), 32, mEbw55xKIj4H)
	case float64:
		 /*line :1*/ hphPuX8VWjn5.rwKBPG3A3fA(zEHZ5ak0_jf, 64, mEbw55xKIj4H)
	case complex64:
		 /*line :1*/ hphPuX8VWjn5.v3XHxGJFrN( /*line :1*/ complex128(zEHZ5ak0_jf), 64, mEbw55xKIj4H)
	case complex128:
		 /*line :1*/ hphPuX8VWjn5.v3XHxGJFrN(zEHZ5ak0_jf, 128, mEbw55xKIj4H)
	case int:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), signed, mEbw55xKIj4H)
	case int8:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), signed, mEbw55xKIj4H)
	case int16:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), signed, mEbw55xKIj4H)
	case int32:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), signed, mEbw55xKIj4H)
	case int64:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), signed, mEbw55xKIj4H)
	case uint:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), unsigned, mEbw55xKIj4H)
	case uint8:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), unsigned, mEbw55xKIj4H)
	case uint16:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), unsigned, mEbw55xKIj4H)
	case uint32:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), unsigned, mEbw55xKIj4H)
	case uint64:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ(zEHZ5ak0_jf, unsigned, mEbw55xKIj4H)
	case uintptr:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64(zEHZ5ak0_jf), unsigned, mEbw55xKIj4H)
	case string:
		 /*line :1*/ hphPuX8VWjn5.yoxXrK6J(zEHZ5ak0_jf, mEbw55xKIj4H)
	case []byte:
		 /*line :1*/ hphPuX8VWjn5.apkzAM(zEHZ5ak0_jf, mEbw55xKIj4H, "[]byte")
	case reflect.Value:

		if  /*line :1*/ zEHZ5ak0_jf.IsValid() &&  /*line :1*/ zEHZ5ak0_jf.CanInterface() {
			hphPuX8VWjn5.geZli9XXQH =  /*line :1*/ zEHZ5ak0_jf.Interface()
			if  /*line :1*/ hphPuX8VWjn5.lz24sYwIhG(mEbw55xKIj4H) {
				return
			}
		}
		 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy(zEHZ5ak0_jf, mEbw55xKIj4H, 0)
	default:

		if ! /*line :1*/ hphPuX8VWjn5.lz24sYwIhG(mEbw55xKIj4H) {

			 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy( /*line :1*/ reflect.SdHoaIQl(zEHZ5ak0_jf), mEbw55xKIj4H, 0)
		}
	}
}



func (hphPuX8VWjn5 *yGfDUYjU3) w7QDMEDqwHjy(uajxUEw reflect.Value, mEbw55xKIj4H rune, kG8aXkabtVO int) {

	if kG8aXkabtVO > 0 &&  /*line :1*/ uajxUEw.IsValid() &&  /*line :1*/ uajxUEw.CanInterface() {
		hphPuX8VWjn5.geZli9XXQH =  /*line :1*/ uajxUEw.Interface()
		if  /*line :1*/ hphPuX8VWjn5.lz24sYwIhG(mEbw55xKIj4H) {
			return
		}
	}
	hphPuX8VWjn5.geZli9XXQH = nil
	hphPuX8VWjn5.ky_TJjz7balj = uajxUEw

	switch zEHZ5ak0_jf := uajxUEw;  /*line :1*/ uajxUEw.Kind() {
	case reflect.Invalid:
		if kG8aXkabtVO == 0 {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(invReflectString)
		} else {
			switch mEbw55xKIj4H {
			case 'v':
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilAngleString)
			default:
				 /*line :1*/ hphPuX8VWjn5.nb_dpqcH0(mEbw55xKIj4H)
			}
		}
	case reflect.Bool:
		 /*line :1*/ hphPuX8VWjn5.ap7H2Bk_xWi( /*line :1*/ zEHZ5ak0_jf.Bool(), mEbw55xKIj4H)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ uint64( /*line :1*/ zEHZ5ak0_jf.Int()), signed, mEbw55xKIj4H)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		 /*line :1*/ hphPuX8VWjn5.bnoiaBZ( /*line :1*/ zEHZ5ak0_jf.Uint(), unsigned, mEbw55xKIj4H)
	case reflect.Float32:
		 /*line :1*/ hphPuX8VWjn5.rwKBPG3A3fA( /*line :1*/ zEHZ5ak0_jf.Float(), 32, mEbw55xKIj4H)
	case reflect.Float64:
		 /*line :1*/ hphPuX8VWjn5.rwKBPG3A3fA( /*line :1*/ zEHZ5ak0_jf.Float(), 64, mEbw55xKIj4H)
	case reflect.Complex64:
		 /*line :1*/ hphPuX8VWjn5.v3XHxGJFrN( /*line :1*/ zEHZ5ak0_jf.Complex(), 64, mEbw55xKIj4H)
	case reflect.Complex128:
		 /*line :1*/ hphPuX8VWjn5.v3XHxGJFrN( /*line :1*/ zEHZ5ak0_jf.Complex(), 128, mEbw55xKIj4H)
	case reflect.String:
		 /*line :1*/ hphPuX8VWjn5.yoxXrK6J( /*line :1*/ zEHZ5ak0_jf.String(), mEbw55xKIj4H)
	case reflect.Map:
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ zEHZ5ak0_jf.Type().String())
			if  /*line :1*/ zEHZ5ak0_jf.IsNil() {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilParenString)
				return
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('{')
		} else {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(mapString)
		}
		glgS2AIPpia :=  /*line :1*/ fmtsort.HhvK6ljdH6m(zEHZ5ak0_jf)
		for wKUgq0A, h9IhTXkOqi := range glgS2AIPpia.Key {
			if wKUgq0A > 0 {
				if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(commaSpaceString)
				} else {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(' ')
				}
			}
			 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy(h9IhTXkOqi, mEbw55xKIj4H, kG8aXkabtVO+1)
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(':')
			 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy(glgS2AIPpia.Value[wKUgq0A], mEbw55xKIj4H, kG8aXkabtVO+1)
		}
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('}')
		} else {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(']')
		}
	case reflect.Struct:
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ zEHZ5ak0_jf.Type().String())
		}
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('{')
		for wKUgq0A := 0; wKUgq0A <  /*line :1*/ zEHZ5ak0_jf.NumField(); wKUgq0A++ {
			if wKUgq0A > 0 {
				if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(commaSpaceString)
				} else {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(' ')
				}
			}
			if hphPuX8VWjn5.hkpuiPFUGfe.j7W9PFs || hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
				if vjTCOm :=  /*line :1*/ zEHZ5ak0_jf.Type().Field(wKUgq0A).ZOSBHkJz; vjTCOm != "" {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(vjTCOm)
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(':')
				}
			}
			 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy( /*line :1*/ lzLfy8wBveLs(zEHZ5ak0_jf, wKUgq0A), mEbw55xKIj4H, kG8aXkabtVO+1)
		}
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('}')
	case reflect.Interface:
		uajxUEw :=  /*line :1*/ zEHZ5ak0_jf.Elem()
		if ! /*line :1*/ uajxUEw.IsValid() {
			if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ zEHZ5ak0_jf.Type().String())
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilParenString)
			} else {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilAngleString)
			}
		} else {
			 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy(uajxUEw, mEbw55xKIj4H, kG8aXkabtVO+1)
		}
	case reflect.Array, reflect.Slice:
		switch mEbw55xKIj4H {
		case 's', 'q', 'x', 'X':

			b8xPjp_1JF :=  /*line :1*/ zEHZ5ak0_jf.Type()
			if  /*line :1*/ b8xPjp_1JF.Elem().Kind() == reflect.Uint8 {
				var lxYa_J []byte
				if  /*line :1*/ zEHZ5ak0_jf.Kind() == reflect.Slice {
					lxYa_J =  /*line :1*/ zEHZ5ak0_jf.Bytes()
				} else if  /*line :1*/ zEHZ5ak0_jf.CanAddr() {
					lxYa_J =  /*line :1*/ zEHZ5ak0_jf.Slice(0,  /*line :1*/ zEHZ5ak0_jf.Len()).Bytes()
				} else {

					lxYa_J =  /*line :1*/ make([]byte,  /*line :1*/ zEHZ5ak0_jf.Len())
					for wKUgq0A := range lxYa_J {
						lxYa_J[wKUgq0A] =  /*line :1*/ byte( /*line :1*/ zEHZ5ak0_jf.Index(wKUgq0A).Uint())
					}
				}
				 /*line :1*/ hphPuX8VWjn5.apkzAM(lxYa_J, mEbw55xKIj4H,  /*line :1*/ b8xPjp_1JF.String())
				return
			}
		}
		if hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ zEHZ5ak0_jf.Type().String())
			if  /*line :1*/ zEHZ5ak0_jf.Kind() == reflect.Slice &&  /*line :1*/ zEHZ5ak0_jf.IsNil() {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilParenString)
				return
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('{')
			for wKUgq0A := 0; wKUgq0A <  /*line :1*/ zEHZ5ak0_jf.Len(); wKUgq0A++ {
				if wKUgq0A > 0 {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(commaSpaceString)
				}
				 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy( /*line :1*/ zEHZ5ak0_jf.Index(wKUgq0A), mEbw55xKIj4H, kG8aXkabtVO+1)
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('}')
		} else {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('[')
			for wKUgq0A := 0; wKUgq0A <  /*line :1*/ zEHZ5ak0_jf.Len(); wKUgq0A++ {
				if wKUgq0A > 0 {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(' ')
				}
				 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy( /*line :1*/ zEHZ5ak0_jf.Index(wKUgq0A), mEbw55xKIj4H, kG8aXkabtVO+1)
			}
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(']')
		}
	case reflect.Pointer:

		if kG8aXkabtVO == 0 &&  /*line :1*/ zEHZ5ak0_jf.Pointer() != 0 {
			switch nghaNY :=  /*line :1*/ zEHZ5ak0_jf.Elem();  /*line :1*/ nghaNY.Kind() {
			case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('&')
				 /*line :1*/ hphPuX8VWjn5.w7QDMEDqwHjy(nghaNY, mEbw55xKIj4H, kG8aXkabtVO+1)
				return
			}
		}
		fallthrough
	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		 /*line :1*/ hphPuX8VWjn5.sja4wY(zEHZ5ak0_jf, mEbw55xKIj4H)
	default:
		 /*line :1*/ hphPuX8VWjn5.zSah4UJOc(zEHZ5ak0_jf)
	}
}


func gLChIbitJHYd(nghaNY []any, aPEqrcp_m int) (ez4abKGFo int, rfnW6Th bool, cIu88ys4 int) {
	cIu88ys4 = aPEqrcp_m
	if aPEqrcp_m <  /*line :1*/ len(nghaNY) {
		ez4abKGFo, rfnW6Th = nghaNY[aPEqrcp_m].(int)
		if !rfnW6Th {

			switch xsKqMyFqj :=  /*line :1*/ reflect.SdHoaIQl(nghaNY[aPEqrcp_m]);  /*line :1*/ xsKqMyFqj.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				bMu41ab1Tf :=  /*line :1*/ xsKqMyFqj.Int()
				if  /*line :1*/ int64( /*line :1*/ int(bMu41ab1Tf)) == bMu41ab1Tf {
					ez4abKGFo =  /*line :1*/ int(bMu41ab1Tf)
					rfnW6Th = true
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				bMu41ab1Tf :=  /*line :1*/ xsKqMyFqj.Uint()
				if  /*line :1*/ int64(bMu41ab1Tf) >= 0 &&  /*line :1*/ uint64( /*line :1*/ int(bMu41ab1Tf)) == bMu41ab1Tf {
					ez4abKGFo =  /*line :1*/ int(bMu41ab1Tf)
					rfnW6Th = true
				}
			default:

			}
		}
		cIu88ys4 = aPEqrcp_m + 1
		if  /*line :1*/ eIJkRuoT(ez4abKGFo) {
			ez4abKGFo = 0
			rfnW6Th = false
		}
	}
	return
}







func pjwo4TTi(aCNXxCRXOS5o string) (gvo8ML48Il int, xgwcYhwzhN int, xP5nLuM9_y bool) {

	if  /*line :1*/ len(aCNXxCRXOS5o) < 3 {
		return 0, 1, false
	}

	for wKUgq0A := 1; wKUgq0A <  /*line :1*/ len(aCNXxCRXOS5o); wKUgq0A++ {
		if aCNXxCRXOS5o[wKUgq0A] == ']' {
			cFeLabEDew_, xP5nLuM9_y, nHl8Y8ilB_I :=  /*line :1*/ bWbu9qxG4bF(aCNXxCRXOS5o, 1, wKUgq0A)
			if !xP5nLuM9_y || nHl8Y8ilB_I != wKUgq0A {
				return 0, wKUgq0A + 1, false
			}
			return cFeLabEDew_ - 1, wKUgq0A + 1, true
		}
	}
	return 0, 1, false
}




func (hphPuX8VWjn5 *yGfDUYjU3) jAg7HgUB(aPEqrcp_m int, aCNXxCRXOS5o string, wKUgq0A int, hAeI_nRg2 int) (cIu88ys4, nHl8Y8ilB_I int, zhNxdMr bool) {
	if  /*line :1*/ len(aCNXxCRXOS5o) <= wKUgq0A || aCNXxCRXOS5o[wKUgq0A] != '[' {
		return aPEqrcp_m, wKUgq0A, false
	}
	hphPuX8VWjn5.j7iqdfrBQJ = true
	gvo8ML48Il, xgwcYhwzhN, xP5nLuM9_y :=  /*line :1*/ pjwo4TTi(aCNXxCRXOS5o[wKUgq0A:])
	if xP5nLuM9_y && 0 <= gvo8ML48Il && gvo8ML48Il < hAeI_nRg2 {
		return gvo8ML48Il, wKUgq0A + xgwcYhwzhN, true
	}
	hphPuX8VWjn5.gTUxvvaG7R = false
	return aPEqrcp_m, wKUgq0A + xgwcYhwzhN, xP5nLuM9_y
}

func (hphPuX8VWjn5 *yGfDUYjU3) ub6QfuXX84tU(mEbw55xKIj4H rune) {
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(percentBangString)
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.j9YRMbR(mEbw55xKIj4H)
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(badIndexString)
}

func (hphPuX8VWjn5 *yGfDUYjU3) iOgqEHPb42Y(mEbw55xKIj4H rune) {
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(percentBangString)
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.j9YRMbR(mEbw55xKIj4H)
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(missingString)
}

func (hphPuX8VWjn5 *yGfDUYjU3) wqxOmPyRm4(aCNXxCRXOS5o string, nghaNY []any) {
	dA0LK4aPCOv :=  /*line :1*/ len(aCNXxCRXOS5o)
	aPEqrcp_m := 0
	nALPLbp := false
	hphPuX8VWjn5.j7iqdfrBQJ = false
formatLoop:
	for wKUgq0A := 0; wKUgq0A < dA0LK4aPCOv; {
		hphPuX8VWjn5.gTUxvvaG7R = true
		_S_nDC := wKUgq0A
		for wKUgq0A < dA0LK4aPCOv && aCNXxCRXOS5o[wKUgq0A] != '%' {
			wKUgq0A++
		}
		if wKUgq0A > _S_nDC {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(aCNXxCRXOS5o[_S_nDC:wKUgq0A])
		}
		if wKUgq0A >= dA0LK4aPCOv {

			break
		}

		wKUgq0A++

		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.xdIs6xWSBJJX()
	simpleFormat:
		for ; wKUgq0A < dA0LK4aPCOv; wKUgq0A++ {
			gaq8WME := aCNXxCRXOS5o[wKUgq0A]
			switch gaq8WME {
			case '#':
				hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN = true
			case '0':
				hphPuX8VWjn5.hkpuiPFUGfe.tpbzso = !hphPuX8VWjn5.hkpuiPFUGfe.i5XVr2y
			case '+':
				hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G = true
			case '-':
				hphPuX8VWjn5.hkpuiPFUGfe.i5XVr2y = true
				hphPuX8VWjn5.hkpuiPFUGfe.tpbzso = false
			case ' ':
				hphPuX8VWjn5.hkpuiPFUGfe.hDpt7tehWk = true
			default:

				if 'a' <= gaq8WME && gaq8WME <= 'z' && aPEqrcp_m <  /*line :1*/ len(nghaNY) {
					switch gaq8WME {
					case 'w':
						hphPuX8VWjn5.f3sXqNFocRH1 =  /*line :1*/ append(hphPuX8VWjn5.f3sXqNFocRH1, aPEqrcp_m)
						fallthrough
					case 'v':

						hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u = hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN
						hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN = false

						hphPuX8VWjn5.hkpuiPFUGfe.j7W9PFs = hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G
						hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G = false
					}
					 /*line :1*/ hphPuX8VWjn5.y383gkr_(nghaNY[aPEqrcp_m],  /*line :1*/ rune(gaq8WME))
					aPEqrcp_m++
					wKUgq0A++
					continue formatLoop
				}

				break simpleFormat
			}
		}

		aPEqrcp_m, wKUgq0A, nALPLbp =  /*line :1*/ hphPuX8VWjn5.jAg7HgUB(aPEqrcp_m, aCNXxCRXOS5o, wKUgq0A,  /*line :1*/ len(nghaNY))

		if wKUgq0A < dA0LK4aPCOv && aCNXxCRXOS5o[wKUgq0A] == '*' {
			wKUgq0A++
			hphPuX8VWjn5.hkpuiPFUGfe.zToz50Xry, hphPuX8VWjn5.hkpuiPFUGfe.fMbhCq, aPEqrcp_m =  /*line :1*/ gLChIbitJHYd(nghaNY, aPEqrcp_m)

			if !hphPuX8VWjn5.hkpuiPFUGfe.fMbhCq {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(badWidthString)
			}

			if hphPuX8VWjn5.hkpuiPFUGfe.zToz50Xry < 0 {
				hphPuX8VWjn5.hkpuiPFUGfe.zToz50Xry = -hphPuX8VWjn5.hkpuiPFUGfe.zToz50Xry
				hphPuX8VWjn5.hkpuiPFUGfe.i5XVr2y = true
				hphPuX8VWjn5.hkpuiPFUGfe.tpbzso = false
			}
			nALPLbp = false
		} else {
			hphPuX8VWjn5.hkpuiPFUGfe.zToz50Xry, hphPuX8VWjn5.hkpuiPFUGfe.fMbhCq, wKUgq0A =  /*line :1*/ bWbu9qxG4bF(aCNXxCRXOS5o, wKUgq0A, dA0LK4aPCOv)
			if nALPLbp && hphPuX8VWjn5.hkpuiPFUGfe.fMbhCq {
				hphPuX8VWjn5.gTUxvvaG7R = false
			}
		}

		if wKUgq0A+1 < dA0LK4aPCOv && aCNXxCRXOS5o[wKUgq0A] == '.' {
			wKUgq0A++
			if nALPLbp {
				hphPuX8VWjn5.gTUxvvaG7R = false
			}
			aPEqrcp_m, wKUgq0A, nALPLbp =  /*line :1*/ hphPuX8VWjn5.jAg7HgUB(aPEqrcp_m, aCNXxCRXOS5o, wKUgq0A,  /*line :1*/ len(nghaNY))
			if wKUgq0A < dA0LK4aPCOv && aCNXxCRXOS5o[wKUgq0A] == '*' {
				wKUgq0A++
				hphPuX8VWjn5.hkpuiPFUGfe.f05l8Dt, hphPuX8VWjn5.hkpuiPFUGfe.dD4Xiyfl, aPEqrcp_m =  /*line :1*/ gLChIbitJHYd(nghaNY, aPEqrcp_m)

				if hphPuX8VWjn5.hkpuiPFUGfe.f05l8Dt < 0 {
					hphPuX8VWjn5.hkpuiPFUGfe.f05l8Dt = 0
					hphPuX8VWjn5.hkpuiPFUGfe.dD4Xiyfl = false
				}
				if !hphPuX8VWjn5.hkpuiPFUGfe.dD4Xiyfl {
					 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(badPrecString)
				}
				nALPLbp = false
			} else {
				hphPuX8VWjn5.hkpuiPFUGfe.f05l8Dt, hphPuX8VWjn5.hkpuiPFUGfe.dD4Xiyfl, wKUgq0A =  /*line :1*/ bWbu9qxG4bF(aCNXxCRXOS5o, wKUgq0A, dA0LK4aPCOv)
				if !hphPuX8VWjn5.hkpuiPFUGfe.dD4Xiyfl {
					hphPuX8VWjn5.hkpuiPFUGfe.f05l8Dt = 0
					hphPuX8VWjn5.hkpuiPFUGfe.dD4Xiyfl = true
				}
			}
		}

		if !nALPLbp {
			aPEqrcp_m, wKUgq0A, nALPLbp =  /*line :1*/ hphPuX8VWjn5.jAg7HgUB(aPEqrcp_m, aCNXxCRXOS5o, wKUgq0A,  /*line :1*/ len(nghaNY))
		}

		if wKUgq0A >= dA0LK4aPCOv {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(noVerbString)
			break
		}

		mEbw55xKIj4H, e2v7yTpR2 :=  /*line :1*/ rune(aCNXxCRXOS5o[wKUgq0A]), 1
		if mEbw55xKIj4H >= utf8.RuneSelf {
			mEbw55xKIj4H, e2v7yTpR2 =  /*line :1*/ utf8.HVAV7aTqxzg(aCNXxCRXOS5o[wKUgq0A:])
		}
		wKUgq0A += e2v7yTpR2

		switch {
		case mEbw55xKIj4H == '%':
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('%')
		case !hphPuX8VWjn5.gTUxvvaG7R:
			 /*line :1*/ hphPuX8VWjn5.ub6QfuXX84tU(mEbw55xKIj4H)
		case aPEqrcp_m >=  /*line :1*/ len(nghaNY):
			 /*line :1*/ hphPuX8VWjn5.iOgqEHPb42Y(mEbw55xKIj4H)
		case mEbw55xKIj4H == 'w':
			hphPuX8VWjn5.f3sXqNFocRH1 =  /*line :1*/ append(hphPuX8VWjn5.f3sXqNFocRH1, aPEqrcp_m)
			fallthrough
		case mEbw55xKIj4H == 'v':

			hphPuX8VWjn5.hkpuiPFUGfe.hnLulZ4y__u = hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN
			hphPuX8VWjn5.hkpuiPFUGfe.htR1xx9jubN = false

			hphPuX8VWjn5.hkpuiPFUGfe.j7W9PFs = hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G
			hphPuX8VWjn5.hkpuiPFUGfe.kbsA6G = false
			fallthrough
		default:
			 /*line :1*/ hphPuX8VWjn5.y383gkr_(nghaNY[aPEqrcp_m], mEbw55xKIj4H)
			aPEqrcp_m++
		}
	}

	if !hphPuX8VWjn5.j7iqdfrBQJ && aPEqrcp_m <  /*line :1*/ len(nghaNY) {
		 /*line :1*/ hphPuX8VWjn5.hkpuiPFUGfe.xdIs6xWSBJJX()
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(extraString)
		for wKUgq0A, oaT9vLaMI3 := range nghaNY[aPEqrcp_m:] {
			if wKUgq0A > 0 {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(commaSpaceString)
			}
			if oaT9vLaMI3 == nil {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk(nilAngleString)
			} else {
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.aKy73W4FDLtk( /*line :1*/ reflect.Cher1a2Fblr(oaT9vLaMI3).String())
				 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('=')
				 /*line :1*/ hphPuX8VWjn5.y383gkr_(oaT9vLaMI3, 'v')
			}
		}
		 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(')')
	}
}

func (hphPuX8VWjn5 *yGfDUYjU3) yf4VotSF2G_I(nghaNY []any) {
	nf8t_DTS := false
	for aPEqrcp_m, oaT9vLaMI3 := range nghaNY {
		hf_Bwycb := oaT9vLaMI3 != nil &&  /*line :1*/ reflect.Cher1a2Fblr(oaT9vLaMI3).Kind() == reflect.String

		if aPEqrcp_m > 0 && !hf_Bwycb && !nf8t_DTS {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(' ')
		}
		 /*line :1*/ hphPuX8VWjn5.y383gkr_(oaT9vLaMI3, 'v')
		nf8t_DTS = hf_Bwycb
	}
}



func (hphPuX8VWjn5 *yGfDUYjU3) zrWkfne7UM(nghaNY []any) {
	for aPEqrcp_m, oaT9vLaMI3 := range nghaNY {
		if aPEqrcp_m > 0 {
			 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG(' ')
		}
		 /*line :1*/ hphPuX8VWjn5.y383gkr_(oaT9vLaMI3, 'v')
	}
	 /*line :1*/ hphPuX8VWjn5.ocZNWYPb9c.yYk12QG('\n')
}
