//line :1
package zBESu2SrRjP

import math "math"


type aL4q0U struct {
	hCC_Vh	uint
	iPKjS5sj	uint
	oXPRffP4	int
}

var cuN5E5Tcg_0 = aL4q0U{23, 8, -127}
var pn7mkGeP = aL4q0U{52, 11, -1023}























func AyQcIQvp(rhZE59OwqM5U float64, i54pnXiXYaZq byte, gsf1Ql, r2moFi1 int) string {
	return  /*line :1*/ string( /*line :1*/ ivhJQa( /*line :1*/ make([]byte, 0,  /*line :1*/ f2STSXp(gsf1Ql+4, 24)), rhZE59OwqM5U, i54pnXiXYaZq, gsf1Ql, r2moFi1))
}



func BLKRZG(l_A2Ytb []byte, rhZE59OwqM5U float64, i54pnXiXYaZq byte, gsf1Ql, r2moFi1 int) []byte {
	return  /*line :1*/ ivhJQa(l_A2Ytb, rhZE59OwqM5U, i54pnXiXYaZq, gsf1Ql, r2moFi1)
}

func ivhJQa(l_A2Ytb []byte, ic2uVKEK float64, i54pnXiXYaZq byte, gsf1Ql, r2moFi1 int) []byte {
	var gYquMvlAg uint64
	var sRaJkQTYfNyc *aL4q0U
	switch r2moFi1 {
	case 32:
		gYquMvlAg =  /*line :1*/ uint64( /*line :1*/ math.FWbZL0aR( /*line :1*/ float32(ic2uVKEK)))
		sRaJkQTYfNyc = &cuN5E5Tcg_0
	case 64:
		gYquMvlAg =  /*line :1*/ math.GKIRmoHE(ic2uVKEK)
		sRaJkQTYfNyc = &pn7mkGeP
	default:
		 /*line :1*/ panic("strconv: illegal AppendFloat/FormatFloat bitSize")
	}

	an5n8FniWt67 := gYquMvlAg>>(sRaJkQTYfNyc.iPKjS5sj+sRaJkQTYfNyc.hCC_Vh) != 0
	uZdS176 :=  /*line :1*/ int(gYquMvlAg>>sRaJkQTYfNyc.hCC_Vh) & (1<<sRaJkQTYfNyc.iPKjS5sj - 1)
	fdyaiZ := gYquMvlAg & ( /*line :1*/ uint64(1)<<sRaJkQTYfNyc.hCC_Vh - 1)

	switch uZdS176 {
	case 1<<sRaJkQTYfNyc.iPKjS5sj - 1:
		
		var a0M8QSnCR string
		switch {
		case fdyaiZ != 0:
			a0M8QSnCR = "NaN"
		case an5n8FniWt67:
			a0M8QSnCR = "-Inf"
		default:
			a0M8QSnCR = "+Inf"
		}
		return  /*line :1*/ append(l_A2Ytb, a0M8QSnCR...)

	case 0:

		uZdS176++

	default:

		fdyaiZ |=  /*line :1*/ uint64(1) << sRaJkQTYfNyc.hCC_Vh
	}
	uZdS176 += sRaJkQTYfNyc.oXPRffP4

	if i54pnXiXYaZq == 'b' {
		return  /*line :1*/ ndg7yn(l_A2Ytb, an5n8FniWt67, fdyaiZ, uZdS176, sRaJkQTYfNyc)
	}
	if i54pnXiXYaZq == 'x' || i54pnXiXYaZq == 'X' {
		return  /*line :1*/ joAwt6P1K(l_A2Ytb, gsf1Ql, i54pnXiXYaZq, an5n8FniWt67, fdyaiZ, uZdS176, sRaJkQTYfNyc)
	}

	if !r6YcaCkRh {
		return  /*line :1*/ eVOwKLxbUTYM(l_A2Ytb, gsf1Ql, i54pnXiXYaZq, an5n8FniWt67, fdyaiZ, uZdS176, sRaJkQTYfNyc)
	}

	var w_GkjTqaTDY cN7vNMaJUh
	eva_XlRz4Iz := false

	t9huUpTT := gsf1Ql < 0
	if t9huUpTT {
		
		var uPg2dG3l8KRn [32]byte
		w_GkjTqaTDY.dvBTkobK6ZN = uPg2dG3l8KRn[:]
		 /*line :1*/ dY6yYkbo(&w_GkjTqaTDY, fdyaiZ, uZdS176- /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh), sRaJkQTYfNyc)
		eva_XlRz4Iz = true

		switch i54pnXiXYaZq {
		case 'e', 'E':
			gsf1Ql =  /*line :1*/ f2STSXp(w_GkjTqaTDY.hc_0uW85-1, 0)
		case 'f':
			gsf1Ql =  /*line :1*/ f2STSXp(w_GkjTqaTDY.hc_0uW85-w_GkjTqaTDY.czr7NK2VrVdf, 0)
		case 'g', 'G':
			gsf1Ql = w_GkjTqaTDY.hc_0uW85
		}
	} else if i54pnXiXYaZq != 'f' {

		fU4UN6Z4vDm := gsf1Ql
		switch i54pnXiXYaZq {
		case 'e', 'E':
			fU4UN6Z4vDm++
		case 'g', 'G':
			if gsf1Ql == 0 {
				gsf1Ql = 1
			}
			fU4UN6Z4vDm = gsf1Ql
		default:

			fU4UN6Z4vDm = 1
		}
		var uPg2dG3l8KRn [24]byte
		if r2moFi1 == 32 && fU4UN6Z4vDm <= 9 {
			w_GkjTqaTDY.dvBTkobK6ZN = uPg2dG3l8KRn[:]
			 /*line :1*/ g3Ug4aphDkYU(&w_GkjTqaTDY,  /*line :1*/ uint32(fdyaiZ), uZdS176- /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh), fU4UN6Z4vDm)
			eva_XlRz4Iz = true
		} else if fU4UN6Z4vDm <= 18 {
			w_GkjTqaTDY.dvBTkobK6ZN = uPg2dG3l8KRn[:]
			 /*line :1*/ asxn8U6lpW(&w_GkjTqaTDY, fdyaiZ, uZdS176- /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh), fU4UN6Z4vDm)
			eva_XlRz4Iz = true
		}
	}
	if !eva_XlRz4Iz {
		return  /*line :1*/ eVOwKLxbUTYM(l_A2Ytb, gsf1Ql, i54pnXiXYaZq, an5n8FniWt67, fdyaiZ, uZdS176, sRaJkQTYfNyc)
	}
	return  /*line :1*/ gPMum5b_(l_A2Ytb, t9huUpTT, an5n8FniWt67, w_GkjTqaTDY, gsf1Ql, i54pnXiXYaZq)
}


func eVOwKLxbUTYM(l_A2Ytb []byte, gsf1Ql int, i54pnXiXYaZq byte, an5n8FniWt67 bool, fdyaiZ uint64, uZdS176 int, sRaJkQTYfNyc *aL4q0U) []byte {
	bkcwawBry :=  /*line :1*/ new(fJphQdAo)
	 /*line :1*/ bkcwawBry.Assign(fdyaiZ)
	 /*line :1*/ bkcwawBry.Shift(uZdS176 -  /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh))
	var w_GkjTqaTDY cN7vNMaJUh
	t9huUpTT := gsf1Ql < 0
	if t9huUpTT {
		 /*line :1*/ nT2paua4k(bkcwawBry, fdyaiZ, uZdS176, sRaJkQTYfNyc)
		w_GkjTqaTDY = cN7vNMaJUh{dvBTkobK6ZN: bkcwawBry.dXlUa8ko[:], hc_0uW85: bkcwawBry.ejR6KwzBSk, czr7NK2VrVdf: bkcwawBry.qvuKHF9jDC}

		switch i54pnXiXYaZq {
		case 'e', 'E':
			gsf1Ql = w_GkjTqaTDY.hc_0uW85 - 1
		case 'f':
			gsf1Ql =  /*line :1*/ f2STSXp(w_GkjTqaTDY.hc_0uW85-w_GkjTqaTDY.czr7NK2VrVdf, 0)
		case 'g', 'G':
			gsf1Ql = w_GkjTqaTDY.hc_0uW85
		}
	} else {

		switch i54pnXiXYaZq {
		case 'e', 'E':
			 /*line :1*/ bkcwawBry.Round(gsf1Ql + 1)
		case 'f':
			 /*line :1*/ bkcwawBry.Round(bkcwawBry.qvuKHF9jDC + gsf1Ql)
		case 'g', 'G':
			if gsf1Ql == 0 {
				gsf1Ql = 1
			}
			 /*line :1*/ bkcwawBry.Round(gsf1Ql)
		}
		w_GkjTqaTDY = cN7vNMaJUh{dvBTkobK6ZN: bkcwawBry.dXlUa8ko[:], hc_0uW85: bkcwawBry.ejR6KwzBSk, czr7NK2VrVdf: bkcwawBry.qvuKHF9jDC}
	}
	return  /*line :1*/ gPMum5b_(l_A2Ytb, t9huUpTT, an5n8FniWt67, w_GkjTqaTDY, gsf1Ql, i54pnXiXYaZq)
}

func gPMum5b_(l_A2Ytb []byte, t9huUpTT bool, an5n8FniWt67 bool, w_GkjTqaTDY cN7vNMaJUh, gsf1Ql int, i54pnXiXYaZq byte) []byte {
	switch i54pnXiXYaZq {
	case 'e', 'E':
		return  /*line :1*/ hseAOYyVqcnB(l_A2Ytb, an5n8FniWt67, w_GkjTqaTDY, gsf1Ql, i54pnXiXYaZq)
	case 'f':
		return  /*line :1*/ drEagJmWooZ(l_A2Ytb, an5n8FniWt67, w_GkjTqaTDY, gsf1Ql)
	case 'g', 'G':

		bkaZYZEQ6DA := gsf1Ql
		if bkaZYZEQ6DA > w_GkjTqaTDY.hc_0uW85 && w_GkjTqaTDY.hc_0uW85 >= w_GkjTqaTDY.czr7NK2VrVdf {
			bkaZYZEQ6DA = w_GkjTqaTDY.hc_0uW85
		}

		if t9huUpTT {
			bkaZYZEQ6DA = 6
		}
		uZdS176 := w_GkjTqaTDY.czr7NK2VrVdf - 1
		if uZdS176 < -4 || uZdS176 >= bkaZYZEQ6DA {
			if gsf1Ql > w_GkjTqaTDY.hc_0uW85 {
				gsf1Ql = w_GkjTqaTDY.hc_0uW85
			}
			return  /*line :1*/ hseAOYyVqcnB(l_A2Ytb, an5n8FniWt67, w_GkjTqaTDY, gsf1Ql-1, i54pnXiXYaZq+'e'-'g')
		}
		if gsf1Ql > w_GkjTqaTDY.czr7NK2VrVdf {
			gsf1Ql = w_GkjTqaTDY.hc_0uW85
		}
		return  /*line :1*/ drEagJmWooZ(l_A2Ytb, an5n8FniWt67, w_GkjTqaTDY,  /*line :1*/ f2STSXp(gsf1Ql-w_GkjTqaTDY.czr7NK2VrVdf, 0))
	}

	return  /*line :1*/ append(l_A2Ytb, '%', i54pnXiXYaZq)
}



func nT2paua4k(bkcwawBry *fJphQdAo, fdyaiZ uint64, uZdS176 int, sRaJkQTYfNyc *aL4q0U) {

	if fdyaiZ == 0 {
		bkcwawBry.ejR6KwzBSk = 0
		return
	}

	x1KGe1 := sRaJkQTYfNyc.oXPRffP4 + 1
	if uZdS176 > x1KGe1 && 332*(bkcwawBry.qvuKHF9jDC-bkcwawBry.ejR6KwzBSk) >= 100*(uZdS176- /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh)) {

		return
	}

	mRp4tM :=  /*line :1*/ new(fJphQdAo)
	 /*line :1*/ mRp4tM.Assign(fdyaiZ*2 + 1)
	 /*line :1*/ mRp4tM.Shift(uZdS176 -  /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh) - 1)

	
	
	
	
	
	
	var bzhBuN uint64
	var dUJ_pzDI int
	if fdyaiZ > 1<<sRaJkQTYfNyc.hCC_Vh || uZdS176 == x1KGe1 {
		bzhBuN = fdyaiZ - 1
		dUJ_pzDI = uZdS176
	} else {
		bzhBuN = fdyaiZ*2 - 1
		dUJ_pzDI = uZdS176 - 1
	}
	chHT2YtqEP :=  /*line :1*/ new(fJphQdAo)
	 /*line :1*/ chHT2YtqEP.Assign(bzhBuN*2 + 1)
	 /*line :1*/ chHT2YtqEP.Shift(dUJ_pzDI -  /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh) - 1)

	f3zbIARc := fdyaiZ%2 == 0

	
	
	
	
	
	
	
	
	
	
	
	var aE9YB1X6Jf uint8

	for j0dfbcd := 0; ; j0dfbcd++ {

		ej15zesmuQun := j0dfbcd - mRp4tM.qvuKHF9jDC + bkcwawBry.qvuKHF9jDC
		if ej15zesmuQun >= bkcwawBry.ejR6KwzBSk {
			break
		}
		nPoh6oR0nJ := j0dfbcd - mRp4tM.qvuKHF9jDC + chHT2YtqEP.qvuKHF9jDC
		gtrgUXWjTDv :=  /*line :1*/ byte('0')
		if nPoh6oR0nJ >= 0 && nPoh6oR0nJ < chHT2YtqEP.ejR6KwzBSk {
			gtrgUXWjTDv = chHT2YtqEP.dXlUa8ko[nPoh6oR0nJ]
		}
		csRwMg :=  /*line :1*/ byte('0')
		if ej15zesmuQun >= 0 {
			csRwMg = bkcwawBry.dXlUa8ko[ej15zesmuQun]
		}
		dnNosEuwQ :=  /*line :1*/ byte('0')
		if j0dfbcd < mRp4tM.ejR6KwzBSk {
			dnNosEuwQ = mRp4tM.dXlUa8ko[j0dfbcd]
		}

		dvLHaM_fG := gtrgUXWjTDv != csRwMg || f3zbIARc && nPoh6oR0nJ+1 == chHT2YtqEP.ejR6KwzBSk

		switch {
		case aE9YB1X6Jf == 0 && csRwMg+1 < dnNosEuwQ:

			aE9YB1X6Jf = 2
		case aE9YB1X6Jf == 0 && csRwMg != dnNosEuwQ:

			aE9YB1X6Jf = 1
		case aE9YB1X6Jf == 1 && (csRwMg != '9' || dnNosEuwQ != '0'):

			aE9YB1X6Jf = 2
		}

		nC2EevhuM_bb := aE9YB1X6Jf > 0 && (f3zbIARc || aE9YB1X6Jf > 1 || j0dfbcd+1 < mRp4tM.ejR6KwzBSk)

		switch {
		case dvLHaM_fG && nC2EevhuM_bb:
			 /*line :1*/ bkcwawBry.Round(ej15zesmuQun + 1)
			return
		case dvLHaM_fG:
			 /*line :1*/ bkcwawBry.RoundDown(ej15zesmuQun + 1)
			return
		case nC2EevhuM_bb:
			 /*line :1*/ bkcwawBry.RoundUp(ej15zesmuQun + 1)
			return
		}
	}
}

type cN7vNMaJUh struct {
	dvBTkobK6ZN	[]byte
	hc_0uW85, czr7NK2VrVdf	int
}


func hseAOYyVqcnB(l_A2Ytb []byte, an5n8FniWt67 bool, bkcwawBry cN7vNMaJUh, gsf1Ql int, i54pnXiXYaZq byte) []byte {

	if an5n8FniWt67 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '-')
	}

	tqWO0Nj :=  /*line :1*/ byte('0')
	if bkcwawBry.hc_0uW85 != 0 {
		tqWO0Nj = bkcwawBry.dvBTkobK6ZN[0]
	}
	l_A2Ytb =  /*line :1*/ append(l_A2Ytb, tqWO0Nj)

	if gsf1Ql > 0 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '.')
		cDhV_2D := 1
		csRwMg :=  /*line :1*/ ymPDEJIA2r5D(bkcwawBry.hc_0uW85, gsf1Ql+1)
		if cDhV_2D < csRwMg {
			l_A2Ytb =  /*line :1*/ append(l_A2Ytb, bkcwawBry.dvBTkobK6ZN[cDhV_2D:csRwMg]...)
			cDhV_2D = csRwMg
		}
		for ; cDhV_2D <= gsf1Ql; cDhV_2D++ {
			l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '0')
		}
	}

	l_A2Ytb =  /*line :1*/ append(l_A2Ytb, i54pnXiXYaZq)
	uZdS176 := bkcwawBry.czr7NK2VrVdf - 1
	if bkcwawBry.hc_0uW85 == 0 {
		uZdS176 = 0
	}
	if uZdS176 < 0 {
		tqWO0Nj = '-'
		uZdS176 = -uZdS176
	} else {
		tqWO0Nj = '+'
	}
	l_A2Ytb =  /*line :1*/ append(l_A2Ytb, tqWO0Nj)

	switch {
	case uZdS176 < 10:
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '0',  /*line :1*/ byte(uZdS176)+'0')
	case uZdS176 < 100:
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb,  /*line :1*/ byte(uZdS176/10)+'0',  /*line :1*/ byte(uZdS176%10)+'0')
	default:
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb,  /*line :1*/ byte(uZdS176/100)+'0',  /*line :1*/ byte(uZdS176/10)%10+'0',  /*line :1*/ byte(uZdS176%10)+'0')
	}

	return l_A2Ytb
}


func drEagJmWooZ(l_A2Ytb []byte, an5n8FniWt67 bool, bkcwawBry cN7vNMaJUh, gsf1Ql int) []byte {

	if an5n8FniWt67 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '-')
	}

	if bkcwawBry.czr7NK2VrVdf > 0 {
		csRwMg :=  /*line :1*/ ymPDEJIA2r5D(bkcwawBry.hc_0uW85, bkcwawBry.czr7NK2VrVdf)
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, bkcwawBry.dvBTkobK6ZN[:csRwMg]...)
		for ; csRwMg < bkcwawBry.czr7NK2VrVdf; csRwMg++ {
			l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '0')
		}
	} else {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '0')
	}

	if gsf1Ql > 0 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '.')
		for cDhV_2D := 0; cDhV_2D < gsf1Ql; cDhV_2D++ {
			tqWO0Nj :=  /*line :1*/ byte('0')
			if dTPavgbkmK := bkcwawBry.czr7NK2VrVdf + cDhV_2D; 0 <= dTPavgbkmK && dTPavgbkmK < bkcwawBry.hc_0uW85 {
				tqWO0Nj = bkcwawBry.dvBTkobK6ZN[dTPavgbkmK]
			}
			l_A2Ytb =  /*line :1*/ append(l_A2Ytb, tqWO0Nj)
		}
	}

	return l_A2Ytb
}


func ndg7yn(l_A2Ytb []byte, an5n8FniWt67 bool, fdyaiZ uint64, uZdS176 int, sRaJkQTYfNyc *aL4q0U) []byte {

	if an5n8FniWt67 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '-')
	}

	l_A2Ytb, _ =  /*line :1*/ lHgNcSQTJz1(l_A2Ytb, fdyaiZ, 10, false, true)

	l_A2Ytb =  /*line :1*/ append(l_A2Ytb, 'p')

	uZdS176 -=  /*line :1*/ int(sRaJkQTYfNyc.hCC_Vh)
	if uZdS176 >= 0 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '+')
	}
	l_A2Ytb, _ =  /*line :1*/ lHgNcSQTJz1(l_A2Ytb,  /*line :1*/ uint64(uZdS176), 10, uZdS176 < 0, true)

	return l_A2Ytb
}


func joAwt6P1K(l_A2Ytb []byte, gsf1Ql int, i54pnXiXYaZq byte, an5n8FniWt67 bool, fdyaiZ uint64, uZdS176 int, sRaJkQTYfNyc *aL4q0U) []byte {
	if fdyaiZ == 0 {
		uZdS176 = 0
	}

	fdyaiZ <<= 60 - sRaJkQTYfNyc.hCC_Vh
	for fdyaiZ != 0 && fdyaiZ&(1<<60) == 0 {
		fdyaiZ <<= 1
		uZdS176--
	}

	if gsf1Ql >= 0 && gsf1Ql < 15 {
		naILJ1HmxfWn :=  /*line :1*/ uint(gsf1Ql * 4)
		v6ctH0 := (fdyaiZ << naILJ1HmxfWn) & (1<<60 - 1)
		fdyaiZ >>= 60 - naILJ1HmxfWn
		if v6ctH0|(fdyaiZ&1) > 1<<59 {
			fdyaiZ++
		}
		fdyaiZ <<= 60 - naILJ1HmxfWn
		if fdyaiZ&(1<<61) != 0 {

			fdyaiZ >>= 1
			uZdS176++
		}
	}

	r6avN8QxQ8D := lowerhex
	if i54pnXiXYaZq == 'X' {
		r6avN8QxQ8D = upperhex
	}

	if an5n8FniWt67 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '-')
	}
	l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '0', i54pnXiXYaZq, '0'+ /*line :1*/ byte((fdyaiZ>>60)&1))

	fdyaiZ <<= 4
	if gsf1Ql < 0 && fdyaiZ != 0 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '.')
		for fdyaiZ != 0 {
			l_A2Ytb =  /*line :1*/ append(l_A2Ytb, r6avN8QxQ8D[(fdyaiZ>>60)&15])
			fdyaiZ <<= 4
		}
	} else if gsf1Ql > 0 {
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb, '.')
		for cDhV_2D := 0; cDhV_2D < gsf1Ql; cDhV_2D++ {
			l_A2Ytb =  /*line :1*/ append(l_A2Ytb, r6avN8QxQ8D[(fdyaiZ>>60)&15])
			fdyaiZ <<= 4
		}
	}

	tqWO0Nj :=  /*line :1*/ byte('P')
	if i54pnXiXYaZq ==  /*line :1*/ chHT2YtqEP(i54pnXiXYaZq) {
		tqWO0Nj = 'p'
	}
	l_A2Ytb =  /*line :1*/ append(l_A2Ytb, tqWO0Nj)
	if uZdS176 < 0 {
		tqWO0Nj = '-'
		uZdS176 = -uZdS176
	} else {
		tqWO0Nj = '+'
	}
	l_A2Ytb =  /*line :1*/ append(l_A2Ytb, tqWO0Nj)

	switch {
	case uZdS176 < 100:
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb,  /*line :1*/ byte(uZdS176/10)+'0',  /*line :1*/ byte(uZdS176%10)+'0')
	case uZdS176 < 1000:
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb,  /*line :1*/ byte(uZdS176/100)+'0',  /*line :1*/ byte((uZdS176/10)%10)+'0',  /*line :1*/ byte(uZdS176%10)+'0')
	default:
		l_A2Ytb =  /*line :1*/ append(l_A2Ytb,  /*line :1*/ byte(uZdS176/1000)+'0',  /*line :1*/ byte(uZdS176/100)%10+'0',  /*line :1*/ byte((uZdS176/10)%10)+'0',  /*line :1*/ byte(uZdS176%10)+'0')
	}

	return l_A2Ytb
}

func ymPDEJIA2r5D(zxZt5L, o7OjhYc5wQns int) int {
	if zxZt5L < o7OjhYc5wQns {
		return zxZt5L
	}
	return o7OjhYc5wQns
}

func f2STSXp(zxZt5L, o7OjhYc5wQns int) int {
	if zxZt5L > o7OjhYc5wQns {
		return zxZt5L
	}
	return o7OjhYc5wQns
}
