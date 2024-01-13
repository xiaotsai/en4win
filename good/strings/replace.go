//line :1
package fQXlzv

import (
	io "xI9ai2djaJ2"
	sync "sync"
)



type UzxCNbzsJa struct {
	n_OwYLIdbuC	sync.LhBfwn6wa1x	
	aF3KpGb4sEse	hJ4QJPSvD
	ruwUgpLO7J	[]string
}


type hJ4QJPSvD interface {
	Replace(w4GNKhKtw string) string
	WriteString(clNaSK02lkVX io.LXQrGW6BQt, w4GNKhKtw string) (dUrxzl int, iA2u51 error)
}







func LzaKcOmi2a(mx3M5z ...string) *UzxCNbzsJa {
	if  /*line :1*/ len(mx3M5z)%2 == 1 {
		 /*line :1*/ panic("strings.NewReplacer: odd argument count")
	}
	return &UzxCNbzsJa{ruwUgpLO7J:  /*line :1*/ append([] /*line :1*/ string(nil), mx3M5z...)}
}

func (unWb50A *UzxCNbzsJa) h_kDCrfq9U8d() {
	unWb50A.aF3KpGb4sEse =  /*line :1*/ unWb50A.dPOyeeVfCh()
	unWb50A.ruwUgpLO7J = nil
}

func (naCMgzayVI *UzxCNbzsJa) dPOyeeVfCh() hJ4QJPSvD {
	mx3M5z := naCMgzayVI.ruwUgpLO7J
	if  /*line :1*/ len(mx3M5z) == 2 &&  /*line :1*/ len(mx3M5z[0]) > 1 {
		return  /*line :1*/ pa5Dkl(mx3M5z[0], mx3M5z[1])
	}

	uxtGwXfbJ := true
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(mx3M5z); gUiH7N += 2 {
		if  /*line :1*/ len(mx3M5z[gUiH7N]) != 1 {
			return  /*line :1*/ m2ETzmFf4aD(mx3M5z)
		}
		if  /*line :1*/ len(mx3M5z[gUiH7N+1]) != 1 {
			uxtGwXfbJ = false
		}
	}

	if uxtGwXfbJ {
		unWb50A := a3VbuLDoyE{}
		for gUiH7N := range unWb50A {
			unWb50A[gUiH7N] =  /*line :1*/ byte(gUiH7N)
		}

		for gUiH7N :=  /*line :1*/ len(mx3M5z) - 2; gUiH7N >= 0; gUiH7N -= 2 {
			aan9Pjdhtu := mx3M5z[gUiH7N][0]
			dUrxzl := mx3M5z[gUiH7N+1][0]
			unWb50A[aan9Pjdhtu] = dUrxzl
		}
		return &unWb50A
	}

	unWb50A := rZbqkmwFh{kg_XO_ahao:  /*line :1*/ make([]string, 0,  /*line :1*/ len(mx3M5z)/2)}

	for gUiH7N :=  /*line :1*/ len(mx3M5z) - 2; gUiH7N >= 0; gUiH7N -= 2 {
		aan9Pjdhtu := mx3M5z[gUiH7N][0]
		dUrxzl := mx3M5z[gUiH7N+1]

		if unWb50A.vQDGT7l8SdLU[aan9Pjdhtu] == nil {

			unWb50A.kg_XO_ahao =  /*line :1*/ append(unWb50A.kg_XO_ahao,  /*line :1*/ string([]byte{aan9Pjdhtu}))
		}
		unWb50A.vQDGT7l8SdLU[aan9Pjdhtu] = [] /*line :1*/ byte(dUrxzl)

	}
	return &unWb50A
}


func (unWb50A *UzxCNbzsJa) Replace(w4GNKhKtw string) string {
	 /*line :1*/ unWb50A.n_OwYLIdbuC.Do(unWb50A.h_kDCrfq9U8d)
	return  /*line :1*/ unWb50A.aF3KpGb4sEse.Replace(w4GNKhKtw)
}


func (unWb50A *UzxCNbzsJa) WriteString(clNaSK02lkVX io.LXQrGW6BQt, w4GNKhKtw string) (dUrxzl int, iA2u51 error) {
	 /*line :1*/ unWb50A.n_OwYLIdbuC.Do(unWb50A.h_kDCrfq9U8d)
	return  /*line :1*/ unWb50A.aF3KpGb4sEse.WriteString(clNaSK02lkVX, w4GNKhKtw)
}


















type d1SKpO_0 struct {
	
	
	s1u5tl	string
	
	
	
	
	
	jBaLiP	int

	
	
	
	rDu2KjN0NvNb	string
	wK7gx9Dny	*d1SKpO_0

	
	
	
	
	
	
	
	uip50Ncn	[]*d1SKpO_0
}

func (az1oIQ9_Q6N *d1SKpO_0) jeJXrSu3wC0G(hJ5Ko2, iYQZtJ string, mEQ5WB int, unWb50A *_2GXWgfcOE) {
	if hJ5Ko2 == "" {
		if az1oIQ9_Q6N.jBaLiP == 0 {
			az1oIQ9_Q6N.s1u5tl = iYQZtJ
			az1oIQ9_Q6N.jBaLiP = mEQ5WB
		}
		return
	}

	if az1oIQ9_Q6N.rDu2KjN0NvNb != "" {
		
		var dUrxzl int	
		for ; dUrxzl <  /*line :1*/ len(az1oIQ9_Q6N.rDu2KjN0NvNb) && dUrxzl <  /*line :1*/ len(hJ5Ko2); dUrxzl++ {
			if az1oIQ9_Q6N.rDu2KjN0NvNb[dUrxzl] != hJ5Ko2[dUrxzl] {
				break
			}
		}
		if dUrxzl ==  /*line :1*/ len(az1oIQ9_Q6N.rDu2KjN0NvNb) {
			 /*line :1*/ az1oIQ9_Q6N.wK7gx9Dny.jeJXrSu3wC0G(hJ5Ko2[dUrxzl:], iYQZtJ, mEQ5WB, unWb50A)
		} else if dUrxzl == 0 {
			
			
			
			var wsWldGA5 *d1SKpO_0
			if  /*line :1*/ len(az1oIQ9_Q6N.rDu2KjN0NvNb) == 1 {
				wsWldGA5 = az1oIQ9_Q6N.wK7gx9Dny
			} else {
				wsWldGA5 = &d1SKpO_0{
					rDu2KjN0NvNb:	az1oIQ9_Q6N.rDu2KjN0NvNb[1:],
					wK7gx9Dny:	az1oIQ9_Q6N.wK7gx9Dny,
				}
			}
			uj6rpA6Vyh :=  /*line :1*/ new(d1SKpO_0)
			az1oIQ9_Q6N.uip50Ncn =  /*line :1*/ make([]*d1SKpO_0, unWb50A.yxbQ1KT)
			az1oIQ9_Q6N.uip50Ncn[unWb50A.u16kSewc_N[az1oIQ9_Q6N.rDu2KjN0NvNb[0]]] = wsWldGA5
			az1oIQ9_Q6N.uip50Ncn[unWb50A.u16kSewc_N[hJ5Ko2[0]]] = uj6rpA6Vyh
			az1oIQ9_Q6N.rDu2KjN0NvNb = ""
			az1oIQ9_Q6N.wK7gx9Dny = nil
			 /*line :1*/ uj6rpA6Vyh.jeJXrSu3wC0G(hJ5Ko2[1:], iYQZtJ, mEQ5WB, unWb50A)
		} else {

			eZ4JDrZAVG := &d1SKpO_0{
				rDu2KjN0NvNb:	az1oIQ9_Q6N.rDu2KjN0NvNb[dUrxzl:],
				wK7gx9Dny:	az1oIQ9_Q6N.wK7gx9Dny,
			}
			az1oIQ9_Q6N.rDu2KjN0NvNb = az1oIQ9_Q6N.rDu2KjN0NvNb[:dUrxzl]
			az1oIQ9_Q6N.wK7gx9Dny = eZ4JDrZAVG
			 /*line :1*/ eZ4JDrZAVG.jeJXrSu3wC0G(hJ5Ko2[dUrxzl:], iYQZtJ, mEQ5WB, unWb50A)
		}
	} else if az1oIQ9_Q6N.uip50Ncn != nil {

		ta2_ET := unWb50A.u16kSewc_N[hJ5Ko2[0]]
		if az1oIQ9_Q6N.uip50Ncn[ta2_ET] == nil {
			az1oIQ9_Q6N.uip50Ncn[ta2_ET] =  /*line :1*/ new(d1SKpO_0)
		}
		 /*line :1*/ az1oIQ9_Q6N.uip50Ncn[ta2_ET].jeJXrSu3wC0G(hJ5Ko2[1:], iYQZtJ, mEQ5WB, unWb50A)
	} else {
		az1oIQ9_Q6N.rDu2KjN0NvNb = hJ5Ko2
		az1oIQ9_Q6N.wK7gx9Dny =  /*line :1*/ new(d1SKpO_0)
		 /*line :1*/ az1oIQ9_Q6N.wK7gx9Dny.jeJXrSu3wC0G("", iYQZtJ, mEQ5WB, unWb50A)
	}
}

func (unWb50A *_2GXWgfcOE) vebNnbxU5IBz(w4GNKhKtw string, nv1m_au bool) (iYQZtJ string, nVZU2y int, eCIbzGMFM2p bool) {

	cl2FUnBsadW := 0
	g7GV99I63 := &unWb50A.uKKCaVoWLjO
	dUrxzl := 0
	for g7GV99I63 != nil {
		if g7GV99I63.jBaLiP > cl2FUnBsadW && !(nv1m_au && g7GV99I63 == &unWb50A.uKKCaVoWLjO) {
			cl2FUnBsadW = g7GV99I63.jBaLiP
			iYQZtJ = g7GV99I63.s1u5tl
			nVZU2y = dUrxzl
			eCIbzGMFM2p = true
		}

		if w4GNKhKtw == "" {
			break
		}
		if g7GV99I63.uip50Ncn != nil {
			oaGnngb1b := unWb50A.u16kSewc_N[w4GNKhKtw[0]]
			if  /*line :1*/ int(oaGnngb1b) == unWb50A.yxbQ1KT {
				break
			}
			g7GV99I63 = g7GV99I63.uip50Ncn[oaGnngb1b]
			w4GNKhKtw = w4GNKhKtw[1:]
			dUrxzl++
		} else if g7GV99I63.rDu2KjN0NvNb != "" &&  /*line :1*/ E8ZCzfgw(w4GNKhKtw, g7GV99I63.rDu2KjN0NvNb) {
			dUrxzl +=  /*line :1*/ len(g7GV99I63.rDu2KjN0NvNb)
			w4GNKhKtw = w4GNKhKtw[ /*line :1*/ len(g7GV99I63.rDu2KjN0NvNb):]
			g7GV99I63 = g7GV99I63.wK7gx9Dny
		} else {
			break
		}
	}
	return
}



type _2GXWgfcOE struct {
	uKKCaVoWLjO	d1SKpO_0
	
	
	yxbQ1KT	int
	
	u16kSewc_N	[256]byte
}

func m2ETzmFf4aD(mx3M5z []string) *_2GXWgfcOE {
	unWb50A :=  /*line :1*/ new(_2GXWgfcOE)

	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(mx3M5z); gUiH7N += 2 {
		hJ5Ko2 := mx3M5z[gUiH7N]
		for xXcc3ZyUB3kD := 0; xXcc3ZyUB3kD <  /*line :1*/ len(hJ5Ko2); xXcc3ZyUB3kD++ {
			unWb50A.u16kSewc_N[hJ5Ko2[xXcc3ZyUB3kD]] = 1
		}
	}

	for _, naCMgzayVI := range unWb50A.u16kSewc_N {
		unWb50A.yxbQ1KT +=  /*line :1*/ int(naCMgzayVI)
	}

	var oaGnngb1b byte
	for gUiH7N, naCMgzayVI := range unWb50A.u16kSewc_N {
		if naCMgzayVI == 0 {
			unWb50A.u16kSewc_N[gUiH7N] =  /*line :1*/ byte(unWb50A.yxbQ1KT)
		} else {
			unWb50A.u16kSewc_N[gUiH7N] = oaGnngb1b
			oaGnngb1b++
		}
	}

	unWb50A.uKKCaVoWLjO.uip50Ncn =  /*line :1*/ make([]*d1SKpO_0, unWb50A.yxbQ1KT)

	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(mx3M5z); gUiH7N += 2 {
		 /*line :1*/ unWb50A.uKKCaVoWLjO.jeJXrSu3wC0G(mx3M5z[gUiH7N], mx3M5z[gUiH7N+1],  /*line :1*/ len(mx3M5z)-gUiH7N, unWb50A)
	}
	return unWb50A
}

type nuqumGeqe []byte


func (clNaSK02lkVX *nuqumGeqe) Write(hc5uhVn []byte) (int, error) {
	*clNaSK02lkVX =  /*line :1*/ append(*clNaSK02lkVX, hc5uhVn...)
	return  /*line :1*/ len(hc5uhVn), nil
}


func (clNaSK02lkVX *nuqumGeqe) WriteString(w4GNKhKtw string) (int, error) {
	*clNaSK02lkVX =  /*line :1*/ append(*clNaSK02lkVX, w4GNKhKtw...)
	return  /*line :1*/ len(w4GNKhKtw), nil
}

type j3gj8E3tdzxW struct {
	crYWup1Y io.LXQrGW6BQt
}

func (clNaSK02lkVX j3gj8E3tdzxW) WriteString(w4GNKhKtw string) (int, error) {
	return  /*line :1*/ clNaSK02lkVX.crYWup1Y.Write([] /*line :1*/ byte(w4GNKhKtw))
}

func mJx8RvtULL7z(clNaSK02lkVX io.LXQrGW6BQt) io.HpBo0GpMvi6 {
	xFUBl6iaG, ocyp5bYUYCEP := clNaSK02lkVX.(io.HpBo0GpMvi6)
	if !ocyp5bYUYCEP {
		xFUBl6iaG = j3gj8E3tdzxW{clNaSK02lkVX}
	}
	return xFUBl6iaG
}

func (unWb50A *_2GXWgfcOE) Replace(w4GNKhKtw string) string {
	qIFlayf :=  /*line :1*/ make(nuqumGeqe, 0,  /*line :1*/ len(w4GNKhKtw))
	 /*line :1*/ unWb50A.WriteString(&qIFlayf, w4GNKhKtw)
	return  /*line :1*/ string(qIFlayf)
}

func (unWb50A *_2GXWgfcOE) WriteString(clNaSK02lkVX io.LXQrGW6BQt, w4GNKhKtw string) (dUrxzl int, iA2u51 error) {
	xFUBl6iaG :=  /*line :1*/ mJx8RvtULL7z(clNaSK02lkVX)
	var z6P4_U4, rZ6E4GZkhPGH int
	var fmgLaoaWh bool
	for gUiH7N := 0; gUiH7N <=  /*line :1*/ len(w4GNKhKtw); {

		if gUiH7N !=  /*line :1*/ len(w4GNKhKtw) && unWb50A.uKKCaVoWLjO.jBaLiP == 0 {
			oaGnngb1b :=  /*line :1*/ int(unWb50A.u16kSewc_N[w4GNKhKtw[gUiH7N]])
			if oaGnngb1b == unWb50A.yxbQ1KT || unWb50A.uKKCaVoWLjO.uip50Ncn[oaGnngb1b] == nil {
				gUiH7N++
				continue
			}
		}

		iYQZtJ, nVZU2y, ad88fzv36C31 :=  /*line :1*/ unWb50A.vebNnbxU5IBz(w4GNKhKtw[gUiH7N:], fmgLaoaWh)
		fmgLaoaWh = ad88fzv36C31 && nVZU2y == 0
		if ad88fzv36C31 {
			rZ6E4GZkhPGH, iA2u51 =  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[z6P4_U4:gUiH7N])
			dUrxzl += rZ6E4GZkhPGH
			if iA2u51 != nil {
				return
			}
			rZ6E4GZkhPGH, iA2u51 =  /*line :1*/ xFUBl6iaG.WriteString(iYQZtJ)
			dUrxzl += rZ6E4GZkhPGH
			if iA2u51 != nil {
				return
			}
			gUiH7N += nVZU2y
			z6P4_U4 = gUiH7N
			continue
		}
		gUiH7N++
	}
	if z6P4_U4 !=  /*line :1*/ len(w4GNKhKtw) {
		rZ6E4GZkhPGH, iA2u51 =  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[z6P4_U4:])
		dUrxzl += rZ6E4GZkhPGH
	}
	return
}



type rrBhypsf struct {
	uPH9GoOdUG	*adgXfXFYJr14
	
	zwK44CaGH2K	string
}

func pa5Dkl(iVumsgT string, jJ70x4Pam1f string) *rrBhypsf {
	return &rrBhypsf{uPH9GoOdUG:  /*line :1*/ nyd62y(iVumsgT), zwK44CaGH2K: jJ70x4Pam1f}
}

func (unWb50A *rrBhypsf) Replace(w4GNKhKtw string) string {
	var qIFlayf CKrNlN3otWSF
	gUiH7N, xoMCN7zYm8 := 0, false
	for {
		ad88fzv36C31 :=  /*line :1*/ unWb50A.uPH9GoOdUG.eZ4JDrZAVG(w4GNKhKtw[gUiH7N:])
		if ad88fzv36C31 == -1 {
			break
		}
		xoMCN7zYm8 = true
		 /*line :1*/ qIFlayf.Grow(ad88fzv36C31 +  /*line :1*/ len(unWb50A.zwK44CaGH2K))
		 /*line :1*/ qIFlayf.WriteString(w4GNKhKtw[gUiH7N : gUiH7N+ad88fzv36C31])
		 /*line :1*/ qIFlayf.WriteString(unWb50A.zwK44CaGH2K)
		gUiH7N += ad88fzv36C31 +  /*line :1*/ len(unWb50A.uPH9GoOdUG.jmfRJJPOO)
	}
	if !xoMCN7zYm8 {
		return w4GNKhKtw
	}
	 /*line :1*/ qIFlayf.WriteString(w4GNKhKtw[gUiH7N:])
	return  /*line :1*/ qIFlayf.String()
}

func (unWb50A *rrBhypsf) WriteString(clNaSK02lkVX io.LXQrGW6BQt, w4GNKhKtw string) (dUrxzl int, iA2u51 error) {
	xFUBl6iaG :=  /*line :1*/ mJx8RvtULL7z(clNaSK02lkVX)
	var gUiH7N, rZ6E4GZkhPGH int
	for {
		ad88fzv36C31 :=  /*line :1*/ unWb50A.uPH9GoOdUG.eZ4JDrZAVG(w4GNKhKtw[gUiH7N:])
		if ad88fzv36C31 == -1 {
			break
		}
		rZ6E4GZkhPGH, iA2u51 =  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[gUiH7N : gUiH7N+ad88fzv36C31])
		dUrxzl += rZ6E4GZkhPGH
		if iA2u51 != nil {
			return
		}
		rZ6E4GZkhPGH, iA2u51 =  /*line :1*/ xFUBl6iaG.WriteString(unWb50A.zwK44CaGH2K)
		dUrxzl += rZ6E4GZkhPGH
		if iA2u51 != nil {
			return
		}
		gUiH7N += ad88fzv36C31 +  /*line :1*/ len(unWb50A.uPH9GoOdUG.jmfRJJPOO)
	}
	rZ6E4GZkhPGH, iA2u51 =  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[gUiH7N:])
	dUrxzl += rZ6E4GZkhPGH
	return
}




type a3VbuLDoyE [256]byte

func (unWb50A *a3VbuLDoyE) Replace(w4GNKhKtw string) string {
	var qIFlayf []byte	
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
		naCMgzayVI := w4GNKhKtw[gUiH7N]
		if unWb50A[naCMgzayVI] != naCMgzayVI {
			if qIFlayf == nil {
				qIFlayf = [] /*line :1*/ byte(w4GNKhKtw)
			}
			qIFlayf[gUiH7N] = unWb50A[naCMgzayVI]
		}
	}
	if qIFlayf == nil {
		return w4GNKhKtw
	}
	return  /*line :1*/ string(qIFlayf)
}

func (unWb50A *a3VbuLDoyE) WriteString(clNaSK02lkVX io.LXQrGW6BQt, w4GNKhKtw string) (dUrxzl int, iA2u51 error) {
	xFUBl6iaG :=  /*line :1*/ mJx8RvtULL7z(clNaSK02lkVX)
	z6P4_U4 := 0
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
		naCMgzayVI := w4GNKhKtw[gUiH7N]
		if unWb50A[naCMgzayVI] == naCMgzayVI {
			continue
		}
		if z6P4_U4 != gUiH7N {
			rZ6E4GZkhPGH, iA2u51 :=  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[z6P4_U4:gUiH7N])
			dUrxzl += rZ6E4GZkhPGH
			if iA2u51 != nil {
				return dUrxzl, iA2u51
			}
		}
		z6P4_U4 = gUiH7N + 1
		gzc2Y4g, iA2u51 :=  /*line :1*/ clNaSK02lkVX.Write(unWb50A[naCMgzayVI :  /*line :1*/ int(naCMgzayVI)+1])
		dUrxzl += gzc2Y4g
		if iA2u51 != nil {
			return dUrxzl, iA2u51
		}
	}
	if z6P4_U4 !=  /*line :1*/ len(w4GNKhKtw) {
		gzc2Y4g, iA2u51 :=  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[z6P4_U4:])
		dUrxzl += gzc2Y4g
		if iA2u51 != nil {
			return dUrxzl, iA2u51
		}
	}
	return dUrxzl, nil
}



type rZbqkmwFh struct {
	
	
	vQDGT7l8SdLU	[256][]byte
	
	
	
	kg_XO_ahao	[]string
}








const countCutOff = 8

func (unWb50A *rZbqkmwFh) Replace(w4GNKhKtw string) string {
	bhVI37 :=  /*line :1*/ len(w4GNKhKtw)
	yKZRkmrEh37 := false

	if  /*line :1*/ len(unWb50A.kg_XO_ahao)*countCutOff <=  /*line :1*/ len(w4GNKhKtw) {
		for _, hB0GXVbS := range unWb50A.kg_XO_ahao {
			if f4rCFxMKmpo :=  /*line :1*/ L9XssxF(w4GNKhKtw, hB0GXVbS); f4rCFxMKmpo != 0 {

				bhVI37 += f4rCFxMKmpo * ( /*line :1*/ len(unWb50A.vQDGT7l8SdLU[hB0GXVbS[0]]) - 1)
				yKZRkmrEh37 = true
			}

		}
	} else {
		for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
			naCMgzayVI := w4GNKhKtw[gUiH7N]
			if unWb50A.vQDGT7l8SdLU[naCMgzayVI] != nil {

				bhVI37 +=  /*line :1*/ len(unWb50A.vQDGT7l8SdLU[naCMgzayVI]) - 1
				yKZRkmrEh37 = true
			}
		}
	}
	if !yKZRkmrEh37 {
		return w4GNKhKtw
	}
	qIFlayf :=  /*line :1*/ make([]byte, bhVI37)
	xXcc3ZyUB3kD := 0
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
		naCMgzayVI := w4GNKhKtw[gUiH7N]
		if unWb50A.vQDGT7l8SdLU[naCMgzayVI] != nil {
			xXcc3ZyUB3kD +=  /*line :1*/ copy(qIFlayf[xXcc3ZyUB3kD:], unWb50A.vQDGT7l8SdLU[naCMgzayVI])
		} else {
			qIFlayf[xXcc3ZyUB3kD] = naCMgzayVI
			xXcc3ZyUB3kD++
		}
	}
	return  /*line :1*/ string(qIFlayf)
}

func (unWb50A *rZbqkmwFh) WriteString(clNaSK02lkVX io.LXQrGW6BQt, w4GNKhKtw string) (dUrxzl int, iA2u51 error) {
	xFUBl6iaG :=  /*line :1*/ mJx8RvtULL7z(clNaSK02lkVX)
	z6P4_U4 := 0
	for gUiH7N := 0; gUiH7N <  /*line :1*/ len(w4GNKhKtw); gUiH7N++ {
		naCMgzayVI := w4GNKhKtw[gUiH7N]
		if unWb50A.vQDGT7l8SdLU[naCMgzayVI] == nil {
			continue
		}
		if z6P4_U4 != gUiH7N {
			gzc2Y4g, iA2u51 :=  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[z6P4_U4:gUiH7N])
			dUrxzl += gzc2Y4g
			if iA2u51 != nil {
				return dUrxzl, iA2u51
			}
		}
		z6P4_U4 = gUiH7N + 1
		gzc2Y4g, iA2u51 :=  /*line :1*/ clNaSK02lkVX.Write(unWb50A.vQDGT7l8SdLU[naCMgzayVI])
		dUrxzl += gzc2Y4g
		if iA2u51 != nil {
			return dUrxzl, iA2u51
		}
	}
	if z6P4_U4 !=  /*line :1*/ len(w4GNKhKtw) {
		var gzc2Y4g int
		gzc2Y4g, iA2u51 =  /*line :1*/ xFUBl6iaG.WriteString(w4GNKhKtw[z6P4_U4:])
		dUrxzl += gzc2Y4g
	}
	return
}
