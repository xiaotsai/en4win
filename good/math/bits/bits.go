//line :1
//go:generate go run make_tables.go

package bits

const uintSize = 32 << (^ /*line :1*/ uint(0) >> 63)

const UintSize = uintSize

func M7qYV0PcLYO(dfnl1Em5_uzs uint) int	{ return UintSize -  /*line :1*/ Len(dfnl1Em5_uzs) }

func N9mrNTCc(dfnl1Em5_uzs uint8) int	{ return 8 -  /*line :1*/ Len8(dfnl1Em5_uzs) }

func CxuJltye2zHW(dfnl1Em5_uzs uint16) int	{ return 16 -  /*line :1*/ Len16(dfnl1Em5_uzs) }

func PEy_LA8fLx(dfnl1Em5_uzs uint32) int	{ return 32 -  /*line :1*/ Len32(dfnl1Em5_uzs) }

func HYxjvQ(dfnl1Em5_uzs uint64) int	{ return 64 -  /*line :1*/ Len64(dfnl1Em5_uzs) }

const deBruijn32 = 0x077CB531

var xUHpwW7VqqB = [32]byte{
	0, 1, 28, 2, 29, 14, 24, 3, 30, 22, 20, 15, 25, 17, 4, 8,
	31, 27, 13, 23, 21, 19, 16, 7, 26, 12, 18, 6, 11, 5, 10, 9,
}

const deBruijn64 = 0x03f79d71b4ca8b09

var bxaCi0U4hf = [64]byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

func AjW10JsK(dfnl1Em5_uzs uint) int {
	if UintSize == 32 {
		return  /*line :1*/ TrailingZeros32( /*line :1*/ uint32(dfnl1Em5_uzs))
	}
	return  /*line :1*/ TrailingZeros64( /*line :1*/ uint64(dfnl1Em5_uzs))
}

func TrailingZeros8(dfnl1Em5_uzs uint8) int {
	return  /*line :1*/ int(ntz8tab[dfnl1Em5_uzs])
}

func TrailingZeros16(dfnl1Em5_uzs uint16) int {
	if dfnl1Em5_uzs == 0 {
		return 16
	}

	return  /*line :1*/ int(xUHpwW7VqqB[ /*line :1*/ uint32(dfnl1Em5_uzs&-dfnl1Em5_uzs)*deBruijn32>>(32-5)])
}

func TrailingZeros32(dfnl1Em5_uzs uint32) int {
	if dfnl1Em5_uzs == 0 {
		return 32
	}

	return  /*line :1*/ int(xUHpwW7VqqB[(dfnl1Em5_uzs&-dfnl1Em5_uzs)*deBruijn32>>(32-5)])
}

func TrailingZeros64(dfnl1Em5_uzs uint64) int {
	if dfnl1Em5_uzs == 0 {
		return 64
	}

	return  /*line :1*/ int(bxaCi0U4hf[(dfnl1Em5_uzs&-dfnl1Em5_uzs)*deBruijn64>>(64-6)])
}

const m0 = 0x5555555555555555
const m1 = 0x3333333333333333
const m2 = 0x0f0f0f0f0f0f0f0f
const m3 = 0x00ff00ff00ff00ff
const m4 = 0x0000ffff0000ffff

func OnesCount(dfnl1Em5_uzs uint) int {
	if UintSize == 32 {
		return  /*line :1*/ OnesCount32( /*line :1*/ uint32(dfnl1Em5_uzs))
	}
	return  /*line :1*/ OnesCount64( /*line :1*/ uint64(dfnl1Em5_uzs))
}

func OnesCount8(dfnl1Em5_uzs uint8) int {
	return  /*line :1*/ int(pop8tab[dfnl1Em5_uzs])
}

func OnesCount16(dfnl1Em5_uzs uint16) int {
	return  /*line :1*/ int(pop8tab[dfnl1Em5_uzs>>8] + pop8tab[dfnl1Em5_uzs&0xff])
}

func OnesCount32(dfnl1Em5_uzs uint32) int {
	return  /*line :1*/ int(pop8tab[dfnl1Em5_uzs>>24] + pop8tab[dfnl1Em5_uzs>>16&0xff] + pop8tab[dfnl1Em5_uzs>>8&0xff] + pop8tab[dfnl1Em5_uzs&0xff])
}

func OnesCount64(dfnl1Em5_uzs uint64) int {

	const m = 1<<64 - 1
	dfnl1Em5_uzs = dfnl1Em5_uzs>>1&(m0&m) + dfnl1Em5_uzs&(m0&m)
	dfnl1Em5_uzs = dfnl1Em5_uzs>>2&(m1&m) + dfnl1Em5_uzs&(m1&m)
	dfnl1Em5_uzs = (dfnl1Em5_uzs>>4 + dfnl1Em5_uzs) & (m2 & m)
	dfnl1Em5_uzs += dfnl1Em5_uzs >> 8
	dfnl1Em5_uzs += dfnl1Em5_uzs >> 16
	dfnl1Em5_uzs += dfnl1Em5_uzs >> 32
	return  /*line :1*/ int(dfnl1Em5_uzs) & (1<<7 - 1)
}

func RotateLeft(dfnl1Em5_uzs uint, d99mRAnwh int) uint {
	if UintSize == 32 {
		return  /*line :1*/ uint( /*line :1*/ RotateLeft32( /*line :1*/ uint32(dfnl1Em5_uzs), d99mRAnwh))
	}
	return  /*line :1*/ uint( /*line :1*/ RotateLeft64( /*line :1*/ uint64(dfnl1Em5_uzs), d99mRAnwh))
}

func RotateLeft8(dfnl1Em5_uzs uint8, d99mRAnwh int) uint8 {
	const n = 8
	bjX1j8Wab :=  /*line :1*/ uint(d99mRAnwh) & (n - 1)
	return dfnl1Em5_uzs<<bjX1j8Wab | dfnl1Em5_uzs>>(n-bjX1j8Wab)
}

func RotateLeft16(dfnl1Em5_uzs uint16, d99mRAnwh int) uint16 {
	const n = 16
	bjX1j8Wab :=  /*line :1*/ uint(d99mRAnwh) & (n - 1)
	return dfnl1Em5_uzs<<bjX1j8Wab | dfnl1Em5_uzs>>(n-bjX1j8Wab)
}

func RotateLeft32(dfnl1Em5_uzs uint32, d99mRAnwh int) uint32 {
	const n = 32
	bjX1j8Wab :=  /*line :1*/ uint(d99mRAnwh) & (n - 1)
	return dfnl1Em5_uzs<<bjX1j8Wab | dfnl1Em5_uzs>>(n-bjX1j8Wab)
}

func RotateLeft64(dfnl1Em5_uzs uint64, d99mRAnwh int) uint64 {
	const n = 64
	bjX1j8Wab :=  /*line :1*/ uint(d99mRAnwh) & (n - 1)
	return dfnl1Em5_uzs<<bjX1j8Wab | dfnl1Em5_uzs>>(n-bjX1j8Wab)
}

func Reverse(dfnl1Em5_uzs uint) uint {
	if UintSize == 32 {
		return  /*line :1*/ uint( /*line :1*/ Reverse32( /*line :1*/ uint32(dfnl1Em5_uzs)))
	}
	return  /*line :1*/ uint( /*line :1*/ Reverse64( /*line :1*/ uint64(dfnl1Em5_uzs)))
}

func Reverse8(dfnl1Em5_uzs uint8) uint8 {
	return rev8tab[dfnl1Em5_uzs]
}

func Reverse16(dfnl1Em5_uzs uint16) uint16 {
	return  /*line :1*/ uint16(rev8tab[dfnl1Em5_uzs>>8]) |  /*line :1*/ uint16(rev8tab[dfnl1Em5_uzs&0xff])<<8
}

func Reverse32(dfnl1Em5_uzs uint32) uint32 {
	const m = 1<<32 - 1
	dfnl1Em5_uzs = dfnl1Em5_uzs>>1&(m0&m) | dfnl1Em5_uzs&(m0&m)<<1
	dfnl1Em5_uzs = dfnl1Em5_uzs>>2&(m1&m) | dfnl1Em5_uzs&(m1&m)<<2
	dfnl1Em5_uzs = dfnl1Em5_uzs>>4&(m2&m) | dfnl1Em5_uzs&(m2&m)<<4
	return  /*line :1*/ ReverseBytes32(dfnl1Em5_uzs)
}

func Reverse64(dfnl1Em5_uzs uint64) uint64 {
	const m = 1<<64 - 1
	dfnl1Em5_uzs = dfnl1Em5_uzs>>1&(m0&m) | dfnl1Em5_uzs&(m0&m)<<1
	dfnl1Em5_uzs = dfnl1Em5_uzs>>2&(m1&m) | dfnl1Em5_uzs&(m1&m)<<2
	dfnl1Em5_uzs = dfnl1Em5_uzs>>4&(m2&m) | dfnl1Em5_uzs&(m2&m)<<4
	return  /*line :1*/ ReverseBytes64(dfnl1Em5_uzs)
}

func Tan1sFEr2K(dfnl1Em5_uzs uint) uint {
	if UintSize == 32 {
		return  /*line :1*/ uint( /*line :1*/ ReverseBytes32( /*line :1*/ uint32(dfnl1Em5_uzs)))
	}
	return  /*line :1*/ uint( /*line :1*/ ReverseBytes64( /*line :1*/ uint64(dfnl1Em5_uzs)))
}

func ReverseBytes16(dfnl1Em5_uzs uint16) uint16 {
	return dfnl1Em5_uzs>>8 | dfnl1Em5_uzs<<8
}

func ReverseBytes32(dfnl1Em5_uzs uint32) uint32 {
	const m = 1<<32 - 1
	dfnl1Em5_uzs = dfnl1Em5_uzs>>8&(m3&m) | dfnl1Em5_uzs&(m3&m)<<8
	return dfnl1Em5_uzs>>16 | dfnl1Em5_uzs<<16
}

func ReverseBytes64(dfnl1Em5_uzs uint64) uint64 {
	const m = 1<<64 - 1
	dfnl1Em5_uzs = dfnl1Em5_uzs>>8&(m3&m) | dfnl1Em5_uzs&(m3&m)<<8
	dfnl1Em5_uzs = dfnl1Em5_uzs>>16&(m4&m) | dfnl1Em5_uzs&(m4&m)<<16
	return dfnl1Em5_uzs>>32 | dfnl1Em5_uzs<<32
}

func Len(dfnl1Em5_uzs uint) int {
	if UintSize == 32 {
		return  /*line :1*/ Len32( /*line :1*/ uint32(dfnl1Em5_uzs))
	}
	return  /*line :1*/ Len64( /*line :1*/ uint64(dfnl1Em5_uzs))
}

func Len8(dfnl1Em5_uzs uint8) int {
	return  /*line :1*/ int(len8tab[dfnl1Em5_uzs])
}

func Len16(dfnl1Em5_uzs uint16) (dpRawtKiI9p int) {
	if dfnl1Em5_uzs >= 1<<8 {
		dfnl1Em5_uzs >>= 8
		dpRawtKiI9p = 8
	}
	return dpRawtKiI9p +  /*line :1*/ int(len8tab[dfnl1Em5_uzs])
}

func Len32(dfnl1Em5_uzs uint32) (dpRawtKiI9p int) {
	if dfnl1Em5_uzs >= 1<<16 {
		dfnl1Em5_uzs >>= 16
		dpRawtKiI9p = 16
	}
	if dfnl1Em5_uzs >= 1<<8 {
		dfnl1Em5_uzs >>= 8
		dpRawtKiI9p += 8
	}
	return dpRawtKiI9p +  /*line :1*/ int(len8tab[dfnl1Em5_uzs])
}

func Len64(dfnl1Em5_uzs uint64) (dpRawtKiI9p int) {
	if dfnl1Em5_uzs >= 1<<32 {
		dfnl1Em5_uzs >>= 32
		dpRawtKiI9p = 32
	}
	if dfnl1Em5_uzs >= 1<<16 {
		dfnl1Em5_uzs >>= 16
		dpRawtKiI9p += 16
	}
	if dfnl1Em5_uzs >= 1<<8 {
		dfnl1Em5_uzs >>= 8
		dpRawtKiI9p += 8
	}
	return dpRawtKiI9p +  /*line :1*/ int(len8tab[dfnl1Em5_uzs])
}

func Add(dfnl1Em5_uzs, a7ZLdzHx9, e1ciE1 uint) (rtkDCS, vI9v5a3M uint) {
	if UintSize == 32 {
		zHEn7NC, ya7Z7Ray :=  /*line :1*/ RjMjs3RX46( /*line :1*/ uint32(dfnl1Em5_uzs),  /*line :1*/ uint32(a7ZLdzHx9),  /*line :1*/ uint32(e1ciE1))
		return  /*line :1*/ uint(zHEn7NC),  /*line :1*/ uint(ya7Z7Ray)
	}
	w_PbshVa7U, tqzatJuou0 :=  /*line :1*/ Add64( /*line :1*/ uint64(dfnl1Em5_uzs),  /*line :1*/ uint64(a7ZLdzHx9),  /*line :1*/ uint64(e1ciE1))
	return  /*line :1*/ uint(w_PbshVa7U),  /*line :1*/ uint(tqzatJuou0)
}

func RjMjs3RX46(dfnl1Em5_uzs, a7ZLdzHx9, e1ciE1 uint32) (rtkDCS, vI9v5a3M uint32) {
	l92f24 :=  /*line :1*/ uint64(dfnl1Em5_uzs) +  /*line :1*/ uint64(a7ZLdzHx9) +  /*line :1*/ uint64(e1ciE1)
	rtkDCS =  /*line :1*/ uint32(l92f24)
	vI9v5a3M =  /*line :1*/ uint32(l92f24 >> 32)
	return
}

func Add64(dfnl1Em5_uzs, a7ZLdzHx9, e1ciE1 uint64) (rtkDCS, vI9v5a3M uint64) {
	rtkDCS = dfnl1Em5_uzs + a7ZLdzHx9 + e1ciE1

	vI9v5a3M = ((dfnl1Em5_uzs & a7ZLdzHx9) | ((dfnl1Em5_uzs | a7ZLdzHx9) &^ rtkDCS)) >> 63
	return
}

func Sub(dfnl1Em5_uzs, a7ZLdzHx9, gHwuzeDVV uint) (eH1QaDk, yiSIM9VKQk uint) {
	if UintSize == 32 {
		b87FMf249L, cawLL3oF3 :=  /*line :1*/ YJ8r2wIuwkSJ( /*line :1*/ uint32(dfnl1Em5_uzs),  /*line :1*/ uint32(a7ZLdzHx9),  /*line :1*/ uint32(gHwuzeDVV))
		return  /*line :1*/ uint(b87FMf249L),  /*line :1*/ uint(cawLL3oF3)
	}
	qv9dCIWJ, bblqveretSbc :=  /*line :1*/ Sub64( /*line :1*/ uint64(dfnl1Em5_uzs),  /*line :1*/ uint64(a7ZLdzHx9),  /*line :1*/ uint64(gHwuzeDVV))
	return  /*line :1*/ uint(qv9dCIWJ),  /*line :1*/ uint(bblqveretSbc)
}

func YJ8r2wIuwkSJ(dfnl1Em5_uzs, a7ZLdzHx9, gHwuzeDVV uint32) (eH1QaDk, yiSIM9VKQk uint32) {
	eH1QaDk = dfnl1Em5_uzs - a7ZLdzHx9 - gHwuzeDVV

	yiSIM9VKQk = ((^dfnl1Em5_uzs & a7ZLdzHx9) | (^(dfnl1Em5_uzs ^ a7ZLdzHx9) & eH1QaDk)) >> 31
	return
}

func Sub64(dfnl1Em5_uzs, a7ZLdzHx9, gHwuzeDVV uint64) (eH1QaDk, yiSIM9VKQk uint64) {
	eH1QaDk = dfnl1Em5_uzs - a7ZLdzHx9 - gHwuzeDVV

	yiSIM9VKQk = ((^dfnl1Em5_uzs & a7ZLdzHx9) | (^(dfnl1Em5_uzs ^ a7ZLdzHx9) & eH1QaDk)) >> 63
	return
}

func Mul(dfnl1Em5_uzs, a7ZLdzHx9 uint) (ndk21zuGAjE, rpUq1aHTlJCu uint) {
	if UintSize == 32 {
		pkryW4LIPw2, mU8859 :=  /*line :1*/ KUiVvl( /*line :1*/ uint32(dfnl1Em5_uzs),  /*line :1*/ uint32(a7ZLdzHx9))
		return  /*line :1*/ uint(pkryW4LIPw2),  /*line :1*/ uint(mU8859)
	}
	pkryW4LIPw2, mU8859 :=  /*line :1*/ Mul64( /*line :1*/ uint64(dfnl1Em5_uzs),  /*line :1*/ uint64(a7ZLdzHx9))
	return  /*line :1*/ uint(pkryW4LIPw2),  /*line :1*/ uint(mU8859)
}

func KUiVvl(dfnl1Em5_uzs, a7ZLdzHx9 uint32) (ndk21zuGAjE, rpUq1aHTlJCu uint32) {
	ufjg9PCtTv :=  /*line :1*/ uint64(dfnl1Em5_uzs) *  /*line :1*/ uint64(a7ZLdzHx9)
	ndk21zuGAjE, rpUq1aHTlJCu =  /*line :1*/ uint32(ufjg9PCtTv>>32),  /*line :1*/ uint32(ufjg9PCtTv)
	return
}

func Mul64(dfnl1Em5_uzs, a7ZLdzHx9 uint64) (ndk21zuGAjE, rpUq1aHTlJCu uint64) {
	const mask32 = 1<<32 - 1
	zvMhIxTEnb := dfnl1Em5_uzs & mask32
	bgb_gMgnr := dfnl1Em5_uzs >> 32
	awo4bqwn := a7ZLdzHx9 & mask32
	rYbL_y := a7ZLdzHx9 >> 32
	e7Sfbe := zvMhIxTEnb * awo4bqwn
	sk5vMfn3WMa := bgb_gMgnr*awo4bqwn + e7Sfbe>>32
	jGA4dcQaBV1u := sk5vMfn3WMa & mask32
	gcZ_gwEj := sk5vMfn3WMa >> 32
	jGA4dcQaBV1u += zvMhIxTEnb * rYbL_y
	ndk21zuGAjE = bgb_gMgnr*rYbL_y + gcZ_gwEj + jGA4dcQaBV1u>>32
	rpUq1aHTlJCu = dfnl1Em5_uzs * a7ZLdzHx9
	return
}

func Div(ndk21zuGAjE, rpUq1aHTlJCu, a7ZLdzHx9 uint) (iEpCp5OK, oi4sGlvU uint) {
	if UintSize == 32 {
		zQ2SjAkaWdK, nkTa5AWUL :=  /*line :1*/ F5VGoDg( /*line :1*/ uint32(ndk21zuGAjE),  /*line :1*/ uint32(rpUq1aHTlJCu),  /*line :1*/ uint32(a7ZLdzHx9))
		return  /*line :1*/ uint(zQ2SjAkaWdK),  /*line :1*/ uint(nkTa5AWUL)
	}
	zQ2SjAkaWdK, nkTa5AWUL :=  /*line :1*/ Div64( /*line :1*/ uint64(ndk21zuGAjE),  /*line :1*/ uint64(rpUq1aHTlJCu),  /*line :1*/ uint64(a7ZLdzHx9))
	return  /*line :1*/ uint(zQ2SjAkaWdK),  /*line :1*/ uint(nkTa5AWUL)
}

func F5VGoDg(ndk21zuGAjE, rpUq1aHTlJCu, a7ZLdzHx9 uint32) (iEpCp5OK, oi4sGlvU uint32) {
	if a7ZLdzHx9 != 0 && a7ZLdzHx9 <= ndk21zuGAjE {
		 /*line :1*/ panic(aWpsJMJAap)
	}
	pfv8LnW :=  /*line :1*/ uint64(ndk21zuGAjE)<<32 |  /*line :1*/ uint64(rpUq1aHTlJCu)
	iEpCp5OK, oi4sGlvU =  /*line :1*/ uint32(pfv8LnW/ /*line :1*/ uint64(a7ZLdzHx9)),  /*line :1*/ uint32(pfv8LnW% /*line :1*/ uint64(a7ZLdzHx9))
	return
}

func Div64(ndk21zuGAjE, rpUq1aHTlJCu, a7ZLdzHx9 uint64) (iEpCp5OK, oi4sGlvU uint64) {
	if a7ZLdzHx9 == 0 {
		 /*line :1*/ panic(rSfE6E_)
	}
	if a7ZLdzHx9 <= ndk21zuGAjE {
		 /*line :1*/ panic(aWpsJMJAap)
	}

	if ndk21zuGAjE == 0 {
		return rpUq1aHTlJCu / a7ZLdzHx9, rpUq1aHTlJCu % a7ZLdzHx9
	}

	bjX1j8Wab :=  /*line :1*/ uint( /*line :1*/ HYxjvQ(a7ZLdzHx9))
	a7ZLdzHx9 <<= bjX1j8Wab

	const (
		two32	= 1 << 32
		mask32	= two32 - 1
	)
	fqj4Lmmw := a7ZLdzHx9 >> 32
	zrJkyUkOy3yL := a7ZLdzHx9 & mask32
	bebgy4ft := ndk21zuGAjE<<bjX1j8Wab | rpUq1aHTlJCu>>(64-bjX1j8Wab)
	xNaqtXi3Zr := rpUq1aHTlJCu << bjX1j8Wab
	bzRkHDM5I := xNaqtXi3Zr >> 32
	coie0a2Rij := xNaqtXi3Zr & mask32
	lONFvf5GDZT1 := bebgy4ft / fqj4Lmmw
	ruaTe3kjs := bebgy4ft - lONFvf5GDZT1*fqj4Lmmw

	for lONFvf5GDZT1 >= two32 || lONFvf5GDZT1*zrJkyUkOy3yL > two32*ruaTe3kjs+bzRkHDM5I {
		lONFvf5GDZT1--
		ruaTe3kjs += fqj4Lmmw
		if ruaTe3kjs >= two32 {
			break
		}
	}

	dR6rEPa5 := bebgy4ft*two32 + bzRkHDM5I - lONFvf5GDZT1*a7ZLdzHx9
	pPt8Ke4xpHx := dR6rEPa5 / fqj4Lmmw
	ruaTe3kjs = dR6rEPa5 - pPt8Ke4xpHx*fqj4Lmmw

	for pPt8Ke4xpHx >= two32 || pPt8Ke4xpHx*zrJkyUkOy3yL > two32*ruaTe3kjs+coie0a2Rij {
		pPt8Ke4xpHx--
		ruaTe3kjs += fqj4Lmmw
		if ruaTe3kjs >= two32 {
			break
		}
	}

	return lONFvf5GDZT1*two32 + pPt8Ke4xpHx, (dR6rEPa5*two32 + coie0a2Rij - pPt8Ke4xpHx*a7ZLdzHx9) >> bjX1j8Wab
}

func FbN5_aRUe(ndk21zuGAjE, rpUq1aHTlJCu, a7ZLdzHx9 uint) uint {
	if UintSize == 32 {
		return  /*line :1*/ uint( /*line :1*/ YdQWcxZmt( /*line :1*/ uint32(ndk21zuGAjE),  /*line :1*/ uint32(rpUq1aHTlJCu),  /*line :1*/ uint32(a7ZLdzHx9)))
	}
	return  /*line :1*/ uint( /*line :1*/ VJolj51y8z( /*line :1*/ uint64(ndk21zuGAjE),  /*line :1*/ uint64(rpUq1aHTlJCu),  /*line :1*/ uint64(a7ZLdzHx9)))
}

func YdQWcxZmt(ndk21zuGAjE, rpUq1aHTlJCu, a7ZLdzHx9 uint32) uint32 {
	return  /*line :1*/ uint32(( /*line :1*/ uint64(ndk21zuGAjE)<<32 |  /*line :1*/ uint64(rpUq1aHTlJCu)) %  /*line :1*/ uint64(a7ZLdzHx9))
}

func VJolj51y8z(ndk21zuGAjE, rpUq1aHTlJCu, a7ZLdzHx9 uint64) uint64 {

	_, oi4sGlvU :=  /*line :1*/ Div64(ndk21zuGAjE%a7ZLdzHx9, rpUq1aHTlJCu, a7ZLdzHx9)
	return oi4sGlvU
}
