//line :1



package CuAc0pPZwf


const (
	RuneError	= '\uFFFD'		
	RuneSelf	= 0x80		
	MaxRune	= '\U0010FFFF'		
	UTFMax	= 4		
)


const (
	surrogateMin	= 0xD800
	surrogateMax	= 0xDFFF
)

const (
	t1	= 0b00000000
	tx	= 0b10000000
	t2	= 0b11000000
	t3	= 0b11100000
	t4	= 0b11110000
	t5	= 0b11111000

	maskx	= 0b00111111
	mask2	= 0b00011111
	mask3	= 0b00001111
	mask4	= 0b00000111

	rune1Max	= 1<<7 - 1
	rune2Max	= 1<<11 - 1
	rune3Max	= 1<<16 - 1

	
	locb	= 0b10000000
	hicb	= 0b10111111

	
	
	
	
	xx	= 0xF1		
	as	= 0xF0		
	s1	= 0x02		
	s2	= 0x13		
	s3	= 0x03		
	s4	= 0x23		
	s5	= 0x34		
	s6	= 0x04		
	s7	= 0x44		
)


var w7sK0EHxyeyi = [256]uint8{

	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
	as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,

	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	xx, xx, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
	s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
	s2, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s4, s3, s3,
	s5, s6, s6, s6, s7, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
}



type r4Acec struct {
	o7H8Og	uint8	
	zQLradBuQ	uint8	
}


var js6PUsgo = [16]r4Acec{
	0:	{locb, hicb},
	1:	{0xA0, hicb},
	2:	{locb, 0x9F},
	3:	{0x90, hicb},
	4:	{locb, 0x8F},
}



func TAy5Jt8FB7p4(y_uLXiS7K1 []byte) bool {
	abdI0I :=  /*line :1*/ len(y_uLXiS7K1)
	if abdI0I == 0 {
		return false
	}
	opetAQy9xAcD := w7sK0EHxyeyi[y_uLXiS7K1[0]]
	if abdI0I >=  /*line :1*/ int(opetAQy9xAcD&7) {
		return true
	}

	_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
	if abdI0I > 1 && (y_uLXiS7K1[1] < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < y_uLXiS7K1[1]) {
		return true
	} else if abdI0I > 2 && (y_uLXiS7K1[2] < locb || hicb < y_uLXiS7K1[2]) {
		return true
	}
	return false
}


func FIjadq0t(i3aeV7fnbhZ0 string) bool {
	abdI0I :=  /*line :1*/ len(i3aeV7fnbhZ0)
	if abdI0I == 0 {
		return false
	}
	opetAQy9xAcD := w7sK0EHxyeyi[i3aeV7fnbhZ0[0]]
	if abdI0I >=  /*line :1*/ int(opetAQy9xAcD&7) {
		return true
	}

	_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
	if abdI0I > 1 && (i3aeV7fnbhZ0[1] < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < i3aeV7fnbhZ0[1]) {
		return true
	} else if abdI0I > 2 && (i3aeV7fnbhZ0[2] < locb || hicb < i3aeV7fnbhZ0[2]) {
		return true
	}
	return false
}









func EicfpCPn(y_uLXiS7K1 []byte) (dBkVrdY rune, iamF66 int) {
	abdI0I :=  /*line :1*/ len(y_uLXiS7K1)
	if abdI0I < 1 {
		return RuneError, 0
	}
	hOpSEahNVx := y_uLXiS7K1[0]
	opetAQy9xAcD := w7sK0EHxyeyi[hOpSEahNVx]
	if opetAQy9xAcD >= as {

		jD9Gl8yw_Tb :=  /*line :1*/ rune(opetAQy9xAcD) << 31 >> 31
		return  /*line :1*/ rune(y_uLXiS7K1[0])&^jD9Gl8yw_Tb | RuneError&jD9Gl8yw_Tb, 1
	}
	biIt2hD :=  /*line :1*/ int(opetAQy9xAcD & 7)
	_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
	if abdI0I < biIt2hD {
		return RuneError, 1
	}
	evKCYK6G := y_uLXiS7K1[1]
	if evKCYK6G < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < evKCYK6G {
		return RuneError, 1
	}
	if biIt2hD <= 2 {
		return  /*line :1*/ rune(hOpSEahNVx&mask2)<<6 |  /*line :1*/ rune(evKCYK6G&maskx), 2
	}
	vBMSEnV := y_uLXiS7K1[2]
	if vBMSEnV < locb || hicb < vBMSEnV {
		return RuneError, 1
	}
	if biIt2hD <= 3 {
		return  /*line :1*/ rune(hOpSEahNVx&mask3)<<12 |  /*line :1*/ rune(evKCYK6G&maskx)<<6 |  /*line :1*/ rune(vBMSEnV&maskx), 3
	}
	lG8CKQ := y_uLXiS7K1[3]
	if lG8CKQ < locb || hicb < lG8CKQ {
		return RuneError, 1
	}
	return  /*line :1*/ rune(hOpSEahNVx&mask4)<<18 |  /*line :1*/ rune(evKCYK6G&maskx)<<12 |  /*line :1*/ rune(vBMSEnV&maskx)<<6 |  /*line :1*/ rune(lG8CKQ&maskx), 4
}









func HVAV7aTqxzg(i3aeV7fnbhZ0 string) (dBkVrdY rune, iamF66 int) {
	abdI0I :=  /*line :1*/ len(i3aeV7fnbhZ0)
	if abdI0I < 1 {
		return RuneError, 0
	}
	aAPqTFdJdwg := i3aeV7fnbhZ0[0]
	opetAQy9xAcD := w7sK0EHxyeyi[aAPqTFdJdwg]
	if opetAQy9xAcD >= as {

		jD9Gl8yw_Tb :=  /*line :1*/ rune(opetAQy9xAcD) << 31 >> 31
		return  /*line :1*/ rune(i3aeV7fnbhZ0[0])&^jD9Gl8yw_Tb | RuneError&jD9Gl8yw_Tb, 1
	}
	biIt2hD :=  /*line :1*/ int(opetAQy9xAcD & 7)
	_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
	if abdI0I < biIt2hD {
		return RuneError, 1
	}
	jdPy3U_xF1S := i3aeV7fnbhZ0[1]
	if jdPy3U_xF1S < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < jdPy3U_xF1S {
		return RuneError, 1
	}
	if biIt2hD <= 2 {
		return  /*line :1*/ rune(aAPqTFdJdwg&mask2)<<6 |  /*line :1*/ rune(jdPy3U_xF1S&maskx), 2
	}
	r13TLgThI := i3aeV7fnbhZ0[2]
	if r13TLgThI < locb || hicb < r13TLgThI {
		return RuneError, 1
	}
	if biIt2hD <= 3 {
		return  /*line :1*/ rune(aAPqTFdJdwg&mask3)<<12 |  /*line :1*/ rune(jdPy3U_xF1S&maskx)<<6 |  /*line :1*/ rune(r13TLgThI&maskx), 3
	}
	wSkjw1gDF := i3aeV7fnbhZ0[3]
	if wSkjw1gDF < locb || hicb < wSkjw1gDF {
		return RuneError, 1
	}
	return  /*line :1*/ rune(aAPqTFdJdwg&mask4)<<18 |  /*line :1*/ rune(jdPy3U_xF1S&maskx)<<12 |  /*line :1*/ rune(r13TLgThI&maskx)<<6 |  /*line :1*/ rune(wSkjw1gDF&maskx), 4
}









func QG3o80(y_uLXiS7K1 []byte) (dBkVrdY rune, iamF66 int) {
	e3YW3pxeo0 :=  /*line :1*/ len(y_uLXiS7K1)
	if e3YW3pxeo0 == 0 {
		return RuneError, 0
	}
	gposRKRuod := e3YW3pxeo0 - 1
	dBkVrdY =  /*line :1*/ rune(y_uLXiS7K1[gposRKRuod])
	if dBkVrdY < RuneSelf {
		return dBkVrdY, 1
	}

	_ql6_qTR4 := e3YW3pxeo0 - UTFMax
	if _ql6_qTR4 < 0 {
		_ql6_qTR4 = 0
	}
	for gposRKRuod--; gposRKRuod >= _ql6_qTR4; gposRKRuod-- {
		if  /*line :1*/ R46nvQ(y_uLXiS7K1[gposRKRuod]) {
			break
		}
	}
	if gposRKRuod < 0 {
		gposRKRuod = 0
	}
	dBkVrdY, iamF66 =  /*line :1*/ EicfpCPn(y_uLXiS7K1[gposRKRuod:e3YW3pxeo0])
	if gposRKRuod+iamF66 != e3YW3pxeo0 {
		return RuneError, 1
	}
	return dBkVrdY, iamF66
}









func G3GHxYa4Lm(i3aeV7fnbhZ0 string) (dBkVrdY rune, iamF66 int) {
	e3YW3pxeo0 :=  /*line :1*/ len(i3aeV7fnbhZ0)
	if e3YW3pxeo0 == 0 {
		return RuneError, 0
	}
	gposRKRuod := e3YW3pxeo0 - 1
	dBkVrdY =  /*line :1*/ rune(i3aeV7fnbhZ0[gposRKRuod])
	if dBkVrdY < RuneSelf {
		return dBkVrdY, 1
	}

	_ql6_qTR4 := e3YW3pxeo0 - UTFMax
	if _ql6_qTR4 < 0 {
		_ql6_qTR4 = 0
	}
	for gposRKRuod--; gposRKRuod >= _ql6_qTR4; gposRKRuod-- {
		if  /*line :1*/ R46nvQ(i3aeV7fnbhZ0[gposRKRuod]) {
			break
		}
	}
	if gposRKRuod < 0 {
		gposRKRuod = 0
	}
	dBkVrdY, iamF66 =  /*line :1*/ HVAV7aTqxzg(i3aeV7fnbhZ0[gposRKRuod:e3YW3pxeo0])
	if gposRKRuod+iamF66 != e3YW3pxeo0 {
		return RuneError, 1
	}
	return dBkVrdY, iamF66
}



func YrF2oG995j9(dBkVrdY rune) int {
	switch {
	case dBkVrdY < 0:
		return -1
	case dBkVrdY <= rune1Max:
		return 1
	case dBkVrdY <= rune2Max:
		return 2
	case surrogateMin <= dBkVrdY && dBkVrdY <= surrogateMax:
		return -1
	case dBkVrdY <= rune3Max:
		return 3
	case dBkVrdY <= MaxRune:
		return 4
	}
	return -1
}




func YoEca6(y_uLXiS7K1 []byte, dBkVrdY rune) int {

	switch x55oRdca31 :=  /*line :1*/ uint32(dBkVrdY); {
	case x55oRdca31 <= rune1Max:
		y_uLXiS7K1[0] =  /*line :1*/ byte(dBkVrdY)
		return 1
	case x55oRdca31 <= rune2Max:
		_ = y_uLXiS7K1[1]
		y_uLXiS7K1[0] = t2 |  /*line :1*/ byte(dBkVrdY>>6)
		y_uLXiS7K1[1] = tx |  /*line :1*/ byte(dBkVrdY)&maskx
		return 2
	case x55oRdca31 > MaxRune, surrogateMin <= x55oRdca31 && x55oRdca31 <= surrogateMax:
		dBkVrdY = RuneError
		fallthrough
	case x55oRdca31 <= rune3Max:
		_ = y_uLXiS7K1[2]
		y_uLXiS7K1[0] = t3 |  /*line :1*/ byte(dBkVrdY>>12)
		y_uLXiS7K1[1] = tx |  /*line :1*/ byte(dBkVrdY>>6)&maskx
		y_uLXiS7K1[2] = tx |  /*line :1*/ byte(dBkVrdY)&maskx
		return 3
	default:
		_ = y_uLXiS7K1[3]
		y_uLXiS7K1[0] = t4 |  /*line :1*/ byte(dBkVrdY>>18)
		y_uLXiS7K1[1] = tx |  /*line :1*/ byte(dBkVrdY>>12)&maskx
		y_uLXiS7K1[2] = tx |  /*line :1*/ byte(dBkVrdY>>6)&maskx
		y_uLXiS7K1[3] = tx |  /*line :1*/ byte(dBkVrdY)&maskx
		return 4
	}
}




func Ht7oMzd(y_uLXiS7K1 []byte, dBkVrdY rune) []byte {

	if  /*line :1*/ uint32(dBkVrdY) <= rune1Max {
		return  /*line :1*/ append(y_uLXiS7K1,  /*line :1*/ byte(dBkVrdY))
	}
	return  /*line :1*/ y1HgQY99(y_uLXiS7K1, dBkVrdY)
}

func y1HgQY99(y_uLXiS7K1 []byte, dBkVrdY rune) []byte {

	switch x55oRdca31 :=  /*line :1*/ uint32(dBkVrdY); {
	case x55oRdca31 <= rune2Max:
		return  /*line :1*/ append(y_uLXiS7K1, t2| /*line :1*/ byte(dBkVrdY>>6), tx| /*line :1*/ byte(dBkVrdY)&maskx)
	case x55oRdca31 > MaxRune, surrogateMin <= x55oRdca31 && x55oRdca31 <= surrogateMax:
		dBkVrdY = RuneError
		fallthrough
	case x55oRdca31 <= rune3Max:
		return  /*line :1*/ append(y_uLXiS7K1, t3| /*line :1*/ byte(dBkVrdY>>12), tx| /*line :1*/ byte(dBkVrdY>>6)&maskx, tx| /*line :1*/ byte(dBkVrdY)&maskx)
	default:
		return  /*line :1*/ append(y_uLXiS7K1, t4| /*line :1*/ byte(dBkVrdY>>18), tx| /*line :1*/ byte(dBkVrdY>>12)&maskx, tx| /*line :1*/ byte(dBkVrdY>>6)&maskx, tx| /*line :1*/ byte(dBkVrdY)&maskx)
	}
}



func SGiJAHBAT(y_uLXiS7K1 []byte) int {
	dGEC93E :=  /*line :1*/ len(y_uLXiS7K1)
	var abdI0I int
	for x55oRdca31 := 0; x55oRdca31 < dGEC93E; {
		abdI0I++
		jLi7xh := y_uLXiS7K1[x55oRdca31]
		if jLi7xh < RuneSelf {

			x55oRdca31++
			continue
		}
		opetAQy9xAcD := w7sK0EHxyeyi[jLi7xh]
		if opetAQy9xAcD == xx {
			x55oRdca31++
			continue
		}
		iamF66 :=  /*line :1*/ int(opetAQy9xAcD & 7)
		if x55oRdca31+iamF66 > dGEC93E {
			x55oRdca31++
			continue
		}
		_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
		if jLi7xh := y_uLXiS7K1[x55oRdca31+1]; jLi7xh < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < jLi7xh {
			iamF66 = 1
		} else if iamF66 == 2 {
		} else if jLi7xh := y_uLXiS7K1[x55oRdca31+2]; jLi7xh < locb || hicb < jLi7xh {
			iamF66 = 1
		} else if iamF66 == 3 {
		} else if jLi7xh := y_uLXiS7K1[x55oRdca31+3]; jLi7xh < locb || hicb < jLi7xh {
			iamF66 = 1
		}
		x55oRdca31 += iamF66
	}
	return abdI0I
}


func IaG2lfIQ1LT(i3aeV7fnbhZ0 string) (abdI0I int) {
	haVX5RF4hSXK :=  /*line :1*/ len(i3aeV7fnbhZ0)
	for x55oRdca31 := 0; x55oRdca31 < haVX5RF4hSXK; abdI0I++ {
		jLi7xh := i3aeV7fnbhZ0[x55oRdca31]
		if jLi7xh < RuneSelf {

			x55oRdca31++
			continue
		}
		opetAQy9xAcD := w7sK0EHxyeyi[jLi7xh]
		if opetAQy9xAcD == xx {
			x55oRdca31++
			continue
		}
		iamF66 :=  /*line :1*/ int(opetAQy9xAcD & 7)
		if x55oRdca31+iamF66 > haVX5RF4hSXK {
			x55oRdca31++
			continue
		}
		_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
		if jLi7xh := i3aeV7fnbhZ0[x55oRdca31+1]; jLi7xh < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < jLi7xh {
			iamF66 = 1
		} else if iamF66 == 2 {
		} else if jLi7xh := i3aeV7fnbhZ0[x55oRdca31+2]; jLi7xh < locb || hicb < jLi7xh {
			iamF66 = 1
		} else if iamF66 == 3 {
		} else if jLi7xh := i3aeV7fnbhZ0[x55oRdca31+3]; jLi7xh < locb || hicb < jLi7xh {
			iamF66 = 1
		}
		x55oRdca31 += iamF66
	}
	return abdI0I
}




func R46nvQ(qapWWiP9 byte) bool	{ return qapWWiP9&0xC0 != 0x80 }


func JkIZacK(y_uLXiS7K1 []byte) bool {

	y_uLXiS7K1 = y_uLXiS7K1[: /*line :1*/ len(y_uLXiS7K1): /*line :1*/ len(y_uLXiS7K1)]

	for  /*line :1*/ len(y_uLXiS7K1) >= 8 {

		lsghp09K2 :=  /*line :1*/ uint32(y_uLXiS7K1[0]) |  /*line :1*/ uint32(y_uLXiS7K1[1])<<8 |  /*line :1*/ uint32(y_uLXiS7K1[2])<<16 |  /*line :1*/ uint32(y_uLXiS7K1[3])<<24
		nbPdVeD :=  /*line :1*/ uint32(y_uLXiS7K1[4]) |  /*line :1*/ uint32(y_uLXiS7K1[5])<<8 |  /*line :1*/ uint32(y_uLXiS7K1[6])<<16 |  /*line :1*/ uint32(y_uLXiS7K1[7])<<24
		if (lsghp09K2|nbPdVeD)&0x80808080 != 0 {

			break
		}
		y_uLXiS7K1 = y_uLXiS7K1[8:]
	}
	abdI0I :=  /*line :1*/ len(y_uLXiS7K1)
	for x55oRdca31 := 0; x55oRdca31 < abdI0I; {
		bDEJoLHpSl := y_uLXiS7K1[x55oRdca31]
		if bDEJoLHpSl < RuneSelf {
			x55oRdca31++
			continue
		}
		opetAQy9xAcD := w7sK0EHxyeyi[bDEJoLHpSl]
		if opetAQy9xAcD == xx {
			return false
		}
		iamF66 :=  /*line :1*/ int(opetAQy9xAcD & 7)
		if x55oRdca31+iamF66 > abdI0I {
			return false
		}
		_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
		if jLi7xh := y_uLXiS7K1[x55oRdca31+1]; jLi7xh < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < jLi7xh {
			return false
		} else if iamF66 == 2 {
		} else if jLi7xh := y_uLXiS7K1[x55oRdca31+2]; jLi7xh < locb || hicb < jLi7xh {
			return false
		} else if iamF66 == 3 {
		} else if jLi7xh := y_uLXiS7K1[x55oRdca31+3]; jLi7xh < locb || hicb < jLi7xh {
			return false
		}
		x55oRdca31 += iamF66
	}
	return true
}


func YiR9Xl_UC5x(i3aeV7fnbhZ0 string) bool {

	for  /*line :1*/ len(i3aeV7fnbhZ0) >= 8 {

		lsghp09K2 :=  /*line :1*/ uint32(i3aeV7fnbhZ0[0]) |  /*line :1*/ uint32(i3aeV7fnbhZ0[1])<<8 |  /*line :1*/ uint32(i3aeV7fnbhZ0[2])<<16 |  /*line :1*/ uint32(i3aeV7fnbhZ0[3])<<24
		nbPdVeD :=  /*line :1*/ uint32(i3aeV7fnbhZ0[4]) |  /*line :1*/ uint32(i3aeV7fnbhZ0[5])<<8 |  /*line :1*/ uint32(i3aeV7fnbhZ0[6])<<16 |  /*line :1*/ uint32(i3aeV7fnbhZ0[7])<<24
		if (lsghp09K2|nbPdVeD)&0x80808080 != 0 {

			break
		}
		i3aeV7fnbhZ0 = i3aeV7fnbhZ0[8:]
	}
	abdI0I :=  /*line :1*/ len(i3aeV7fnbhZ0)
	for x55oRdca31 := 0; x55oRdca31 < abdI0I; {
		s_iijBA := i3aeV7fnbhZ0[x55oRdca31]
		if s_iijBA < RuneSelf {
			x55oRdca31++
			continue
		}
		opetAQy9xAcD := w7sK0EHxyeyi[s_iijBA]
		if opetAQy9xAcD == xx {
			return false
		}
		iamF66 :=  /*line :1*/ int(opetAQy9xAcD & 7)
		if x55oRdca31+iamF66 > abdI0I {
			return false
		}
		_fhPjy3aUF := js6PUsgo[opetAQy9xAcD>>4]
		if jLi7xh := i3aeV7fnbhZ0[x55oRdca31+1]; jLi7xh < _fhPjy3aUF.o7H8Og || _fhPjy3aUF.zQLradBuQ < jLi7xh {
			return false
		} else if iamF66 == 2 {
		} else if jLi7xh := i3aeV7fnbhZ0[x55oRdca31+2]; jLi7xh < locb || hicb < jLi7xh {
			return false
		} else if iamF66 == 3 {
		} else if jLi7xh := i3aeV7fnbhZ0[x55oRdca31+3]; jLi7xh < locb || hicb < jLi7xh {
			return false
		}
		x55oRdca31 += iamF66
	}
	return true
}



func DYq33yNk(dBkVrdY rune) bool {
	switch {
	case 0 <= dBkVrdY && dBkVrdY < surrogateMin:
		return true
	case surrogateMax < dBkVrdY && dBkVrdY <= MaxRune:
		return true
	}
	return false
}
