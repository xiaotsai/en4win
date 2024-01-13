//line :1


package vGJV8maRK

const (
	MaxRune	= '\U0010FFFF'		
	ReplacementChar	= '\uFFFD'		
	MaxASCII	= '\u007F'		
	MaxLatin1	= '\u00FF'		
)






type JQ62a93TyUT struct {
	KMhFA6PW	[]FSxi9ozzFWZf
	G2mnAfqco3	[]OEAI00a
	PflfRQSfzsyD	int	
}



type FSxi9ozzFWZf struct {
	Cao5DpE	uint16
	MjJXWflARp	uint16
	BN3gKxKED	uint16
}




type OEAI00a struct {
	HW82pfYa4	uint32
	FewdfxJk5nX	uint32
	SXG7TyiiM	uint32
}













type Y6FjY3GCi_F struct {
	F_WWNHXMMS	uint32
	FpRxQXVM	uint32
	V27a9TWedYG3	llcQcaRv
}



type KmthvGp3gyLu []Y6FjY3GCi_F


const (
	UpperCase	= iota
	LowerCase
	TitleCase
	MaxCase
)

type llcQcaRv [MaxCase]rune	




const (
	UpperLower = MaxRune + 1	
)



const linearMax = 18


func dgCNyO(wzezyWlY5Zg []FSxi9ozzFWZf, vYe94x uint16) bool {
	if  /*line :1*/ len(wzezyWlY5Zg) <= linearMax || vYe94x <= MaxLatin1 {
		for tFhJuK := range wzezyWlY5Zg {
			bTdSvla := &wzezyWlY5Zg[tFhJuK]
			if vYe94x < bTdSvla.Cao5DpE {
				return false
			}
			if vYe94x <= bTdSvla.MjJXWflARp {
				return bTdSvla.BN3gKxKED == 1 || (vYe94x-bTdSvla.Cao5DpE)%bTdSvla.BN3gKxKED == 0
			}
		}
		return false
	}

	wgv1coia4dI := 0
	woixh4Cu2aLp :=  /*line :1*/ len(wzezyWlY5Zg)
	for wgv1coia4dI < woixh4Cu2aLp {
		tKUHN_T7K := wgv1coia4dI + (woixh4Cu2aLp-wgv1coia4dI)/2
		bTdSvla := &wzezyWlY5Zg[tKUHN_T7K]
		if bTdSvla.Cao5DpE <= vYe94x && vYe94x <= bTdSvla.MjJXWflARp {
			return bTdSvla.BN3gKxKED == 1 || (vYe94x-bTdSvla.Cao5DpE)%bTdSvla.BN3gKxKED == 0
		}
		if vYe94x < bTdSvla.Cao5DpE {
			woixh4Cu2aLp = tKUHN_T7K
		} else {
			wgv1coia4dI = tKUHN_T7K + 1
		}
	}
	return false
}


func bXR2jW(wzezyWlY5Zg []OEAI00a, vYe94x uint32) bool {
	if  /*line :1*/ len(wzezyWlY5Zg) <= linearMax {
		for tFhJuK := range wzezyWlY5Zg {
			bTdSvla := &wzezyWlY5Zg[tFhJuK]
			if vYe94x < bTdSvla.HW82pfYa4 {
				return false
			}
			if vYe94x <= bTdSvla.FewdfxJk5nX {
				return bTdSvla.SXG7TyiiM == 1 || (vYe94x-bTdSvla.HW82pfYa4)%bTdSvla.SXG7TyiiM == 0
			}
		}
		return false
	}

	wgv1coia4dI := 0
	woixh4Cu2aLp :=  /*line :1*/ len(wzezyWlY5Zg)
	for wgv1coia4dI < woixh4Cu2aLp {
		tKUHN_T7K := wgv1coia4dI + (woixh4Cu2aLp-wgv1coia4dI)/2
		bTdSvla := wzezyWlY5Zg[tKUHN_T7K]
		if bTdSvla.HW82pfYa4 <= vYe94x && vYe94x <= bTdSvla.FewdfxJk5nX {
			return bTdSvla.SXG7TyiiM == 1 || (vYe94x-bTdSvla.HW82pfYa4)%bTdSvla.SXG7TyiiM == 0
		}
		if vYe94x < bTdSvla.HW82pfYa4 {
			woixh4Cu2aLp = tKUHN_T7K
		} else {
			wgv1coia4dI = tKUHN_T7K + 1
		}
	}
	return false
}


func WDJ89N0sM4(zc_K5REBqINe *JQ62a93TyUT, vYe94x rune) bool {
	_WDVpm := zc_K5REBqINe.KMhFA6PW

	if  /*line :1*/ len(_WDVpm) > 0 &&  /*line :1*/ uint32(vYe94x) <=  /*line :1*/ uint32(_WDVpm[ /*line :1*/ len(_WDVpm)-1].MjJXWflARp) {
		return  /*line :1*/ dgCNyO(_WDVpm,  /*line :1*/ uint16(vYe94x))
	}
	p3EoLg := zc_K5REBqINe.G2mnAfqco3
	if  /*line :1*/ len(p3EoLg) > 0 && vYe94x >=  /*line :1*/ rune(p3EoLg[0].HW82pfYa4) {
		return  /*line :1*/ bXR2jW(p3EoLg,  /*line :1*/ uint32(vYe94x))
	}
	return false
}

func juBSvGC(zc_K5REBqINe *JQ62a93TyUT, vYe94x rune) bool {
	_WDVpm := zc_K5REBqINe.KMhFA6PW

	if lZ3bmO7 := zc_K5REBqINe.PflfRQSfzsyD;  /*line :1*/ len(_WDVpm) > lZ3bmO7 &&  /*line :1*/ uint32(vYe94x) <=  /*line :1*/ uint32(_WDVpm[ /*line :1*/ len(_WDVpm)-1].MjJXWflARp) {
		return  /*line :1*/ dgCNyO(_WDVpm[lZ3bmO7:],  /*line :1*/ uint16(vYe94x))
	}
	p3EoLg := zc_K5REBqINe.G2mnAfqco3
	if  /*line :1*/ len(p3EoLg) > 0 && vYe94x >=  /*line :1*/ rune(p3EoLg[0].HW82pfYa4) {
		return  /*line :1*/ bXR2jW(p3EoLg,  /*line :1*/ uint32(vYe94x))
	}
	return false
}


func D2OffabSaO(vYe94x rune) bool {

	if  /*line :1*/ uint32(vYe94x) <= MaxLatin1 {
		return _8qOpKbOi[ /*line :1*/ uint8(vYe94x)]&pLmask == pLu
	}
	return  /*line :1*/ juBSvGC(NfkCkXUtw, vYe94x)
}


func GFHkPP(vYe94x rune) bool {

	if  /*line :1*/ uint32(vYe94x) <= MaxLatin1 {
		return _8qOpKbOi[ /*line :1*/ uint8(vYe94x)]&pLmask == pLl
	}
	return  /*line :1*/ juBSvGC(U8lTp3, vYe94x)
}


func Sul99SZuni(vYe94x rune) bool {
	if vYe94x <= MaxLatin1 {
		return false
	}
	return  /*line :1*/ juBSvGC(T2MnT23isv, vYe94x)
}



func lV59nIRcXlf(ngLI0MZW5 int, vYe94x rune, uGUKI4IGje []Y6FjY3GCi_F) (m9aSduwx rune, cQ3Ft81p bool) {
	if ngLI0MZW5 < 0 || MaxCase <= ngLI0MZW5 {
		return ReplacementChar, false
	}

	wgv1coia4dI := 0
	woixh4Cu2aLp :=  /*line :1*/ len(uGUKI4IGje)
	for wgv1coia4dI < woixh4Cu2aLp {
		tKUHN_T7K := wgv1coia4dI + (woixh4Cu2aLp-wgv1coia4dI)/2
		aC97Wf5m1D1 := uGUKI4IGje[tKUHN_T7K]
		if  /*line :1*/ rune(aC97Wf5m1D1.F_WWNHXMMS) <= vYe94x && vYe94x <=  /*line :1*/ rune(aC97Wf5m1D1.FpRxQXVM) {
			elUe7N := aC97Wf5m1D1.V27a9TWedYG3[ngLI0MZW5]
			if elUe7N > MaxRune {

				return  /*line :1*/ rune(aC97Wf5m1D1.F_WWNHXMMS) + ((vYe94x- /*line :1*/ rune(aC97Wf5m1D1.F_WWNHXMMS))&^1 |  /*line :1*/ rune(ngLI0MZW5&1)), true
			}
			return vYe94x + elUe7N, true
		}
		if vYe94x <  /*line :1*/ rune(aC97Wf5m1D1.F_WWNHXMMS) {
			woixh4Cu2aLp = tKUHN_T7K
		} else {
			wgv1coia4dI = tKUHN_T7K + 1
		}
	}
	return vYe94x, false
}


func QQFQnlz(ngLI0MZW5 int, vYe94x rune) rune {
	vYe94x, _ =  /*line :1*/ lV59nIRcXlf(ngLI0MZW5, vYe94x, Q3AJCtHTjZN)
	return vYe94x
}


func HJD8X_JbK(vYe94x rune) rune {
	if vYe94x <= MaxASCII {
		if 'a' <= vYe94x && vYe94x <= 'z' {
			vYe94x -= 'a' - 'A'
		}
		return vYe94x
	}
	return  /*line :1*/ QQFQnlz(UpperCase, vYe94x)
}


func OLZ8W2(vYe94x rune) rune {
	if vYe94x <= MaxASCII {
		if 'A' <= vYe94x && vYe94x <= 'Z' {
			vYe94x += 'a' - 'A'
		}
		return vYe94x
	}
	return  /*line :1*/ QQFQnlz(LowerCase, vYe94x)
}


func QgPuNi8kYiC(vYe94x rune) rune {
	if vYe94x <= MaxASCII {
		if 'a' <= vYe94x && vYe94x <= 'z' {
			vYe94x -= 'a' - 'A'
		}
		return vYe94x
	}
	return  /*line :1*/ QQFQnlz(TitleCase, vYe94x)
}


func (qNtoWC4aYT KmthvGp3gyLu) ToUpper(vYe94x rune) rune {
	tt1Ufwo, fJNoojday3a :=  /*line :1*/ lV59nIRcXlf(UpperCase, vYe94x, [] /*line :1*/ Y6FjY3GCi_F(qNtoWC4aYT))
	if tt1Ufwo == vYe94x && !fJNoojday3a {
		tt1Ufwo =  /*line :1*/ HJD8X_JbK(vYe94x)
	}
	return tt1Ufwo
}


func (qNtoWC4aYT KmthvGp3gyLu) ToTitle(vYe94x rune) rune {
	tt1Ufwo, fJNoojday3a :=  /*line :1*/ lV59nIRcXlf(TitleCase, vYe94x, [] /*line :1*/ Y6FjY3GCi_F(qNtoWC4aYT))
	if tt1Ufwo == vYe94x && !fJNoojday3a {
		tt1Ufwo =  /*line :1*/ QgPuNi8kYiC(vYe94x)
	}
	return tt1Ufwo
}


func (qNtoWC4aYT KmthvGp3gyLu) ToLower(vYe94x rune) rune {
	tt1Ufwo, fJNoojday3a :=  /*line :1*/ lV59nIRcXlf(LowerCase, vYe94x, [] /*line :1*/ Y6FjY3GCi_F(qNtoWC4aYT))
	if tt1Ufwo == vYe94x && !fJNoojday3a {
		tt1Ufwo =  /*line :1*/ OLZ8W2(vYe94x)
	}
	return tt1Ufwo
}





type rCwpzaX struct {
	GypxwWX	uint16
	WW6wEw	uint16
}



















func NGnxX0NFZA4b(vYe94x rune) rune {
	if vYe94x < 0 || vYe94x > MaxRune {
		return vYe94x
	}

	if  /*line :1*/ int(vYe94x) <  /*line :1*/ len(adbksEL) {
		return  /*line :1*/ rune(adbksEL[vYe94x])
	}

	wgv1coia4dI := 0
	woixh4Cu2aLp :=  /*line :1*/ len(jYC7lb6)
	for wgv1coia4dI < woixh4Cu2aLp {
		tKUHN_T7K := wgv1coia4dI + (woixh4Cu2aLp-wgv1coia4dI)/2
		if  /*line :1*/ rune(jYC7lb6[tKUHN_T7K].GypxwWX) < vYe94x {
			wgv1coia4dI = tKUHN_T7K + 1
		} else {
			woixh4Cu2aLp = tKUHN_T7K
		}
	}
	if wgv1coia4dI <  /*line :1*/ len(jYC7lb6) &&  /*line :1*/ rune(jYC7lb6[wgv1coia4dI].GypxwWX) == vYe94x {
		return  /*line :1*/ rune(jYC7lb6[wgv1coia4dI].WW6wEw)
	}

	if triIxMRuuP :=  /*line :1*/ OLZ8W2(vYe94x); triIxMRuuP != vYe94x {
		return triIxMRuuP
	}
	return  /*line :1*/ HJD8X_JbK(vYe94x)
}
