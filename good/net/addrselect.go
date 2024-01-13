//line :1
package gF9mZs2

import (
	netip "iPdCam_KQOXj"
	sort "gzHaha55n"
)

func e6vy2evrzCS2(md4QSRkO []FZJphYv9hV) {
	if  /*line :1*/ len(md4QSRkO) < 2 {
		return
	}
	 /*line :1*/ hgyEYEvHKC2K(md4QSRkO,  /*line :1*/ hBVJeQA3cFD(md4QSRkO))
}

func hgyEYEvHKC2K(md4QSRkO []FZJphYv9hV, qS4zcqNdLbr []netip.KFLQ1_1E) {
	if  /*line :1*/ len(md4QSRkO) !=  /*line :1*/ len(qS4zcqNdLbr) {
		 /*line :1*/ panic("internal error")
	}
	gEfrCaz :=  /*line :1*/ make([]tb9UkQhh71r,  /*line :1*/ len(md4QSRkO))
	mZmiVCzf3 :=  /*line :1*/ make([]tb9UkQhh71r,  /*line :1*/ len(qS4zcqNdLbr))
	for eaAUaB7Z, ljsCSk := range md4QSRkO {
		eulg_F39PnU, _ :=  /*line :1*/ netip.OWcivL5ykj9(ljsCSk.GdouSNq80bRw)
		gEfrCaz[eaAUaB7Z] =  /*line :1*/ jVr277uDhSjE(eulg_F39PnU)
		mZmiVCzf3[eaAUaB7Z] =  /*line :1*/ jVr277uDhSjE(qS4zcqNdLbr[eaAUaB7Z])
	}
	 /*line :1*/ sort.IUVpc45h_Y3(&k7wdaWX{
		kWHq8tR:	md4QSRkO,
		tGRFVVqj:	gEfrCaz,
		dlMmW5XHOv7:	qS4zcqNdLbr,
		mULghj2IkNE:	mZmiVCzf3,
	})
}




func hBVJeQA3cFD(md4QSRkO []FZJphYv9hV) []netip.KFLQ1_1E {
	qS4zcqNdLbr :=  /*line :1*/ make([]netip.KFLQ1_1E,  /*line :1*/ len(md4QSRkO))
	xnC59rPcIS := ZaffanpNx4{EGf1xSAbsOU5: 9}
	for eaAUaB7Z := range md4QSRkO {
		xnC59rPcIS.ERvaNiiVmytR = md4QSRkO[eaAUaB7Z].GdouSNq80bRw
		xnC59rPcIS.RgTF6a1Xmo = md4QSRkO[eaAUaB7Z].Cyg5M2vYIby
		hl8w2lFX, h_ljk48Bm :=  /*line :1*/ GuS3JRd2o("udp", nil, &xnC59rPcIS)
		if h_ljk48Bm == nil {
			if i0qXWLS, d30HIwxht :=  /*line :1*/ hl8w2lFX.LocalAddr().(*ZaffanpNx4); d30HIwxht {
				qS4zcqNdLbr[eaAUaB7Z], _ =  /*line :1*/ netip.OWcivL5ykj9(i0qXWLS.ERvaNiiVmytR)
			}
			 /*line :1*/ hl8w2lFX.Close()
		}
	}
	return qS4zcqNdLbr
}

type tb9UkQhh71r struct {
	UfOjaDEY	sJQNyGJDJZie
	U2NHFCQ2_wC	uint8
	Q_5GCk03fL9f	uint8
}

func jVr277uDhSjE(kQ8_UEhxU netip.KFLQ1_1E) tb9UkQhh71r {
	if ! /*line :1*/ kQ8_UEhxU.IsValid() {
		return tb9UkQhh71r{}
	}
	gLTuCJ :=  /*line :1*/ b9ldTQ7gISd.Classify(kQ8_UEhxU)
	return tb9UkQhh71r{
		UfOjaDEY:	 /*line :1*/ f30RR4U_X(kQ8_UEhxU),
		U2NHFCQ2_wC:	gLTuCJ.CeOWdP4K,
		Q_5GCk03fL9f:	gLTuCJ.K3cDvWkJ0pTa,
	}
}

type k7wdaWX struct {
	kWHq8tR	[]FZJphYv9hV	
	tGRFVVqj	[]tb9UkQhh71r
	dlMmW5XHOv7	[]netip.KFLQ1_1E	
	mULghj2IkNE	[]tb9UkQhh71r
}

func (dRoFccaZ *k7wdaWX) Len() int	{ return  /*line :1*/ len(dRoFccaZ.kWHq8tR) }

func (dRoFccaZ *k7wdaWX) Swap(eaAUaB7Z, kVgtfLFZ int) {
	dRoFccaZ.kWHq8tR[eaAUaB7Z], dRoFccaZ.kWHq8tR[kVgtfLFZ] = dRoFccaZ.kWHq8tR[kVgtfLFZ], dRoFccaZ.kWHq8tR[eaAUaB7Z]
	dRoFccaZ.dlMmW5XHOv7[eaAUaB7Z], dRoFccaZ.dlMmW5XHOv7[kVgtfLFZ] = dRoFccaZ.dlMmW5XHOv7[kVgtfLFZ], dRoFccaZ.dlMmW5XHOv7[eaAUaB7Z]
	dRoFccaZ.tGRFVVqj[eaAUaB7Z], dRoFccaZ.tGRFVVqj[kVgtfLFZ] = dRoFccaZ.tGRFVVqj[kVgtfLFZ], dRoFccaZ.tGRFVVqj[eaAUaB7Z]
	dRoFccaZ.mULghj2IkNE[eaAUaB7Z], dRoFccaZ.mULghj2IkNE[kVgtfLFZ] = dRoFccaZ.mULghj2IkNE[kVgtfLFZ], dRoFccaZ.mULghj2IkNE[eaAUaB7Z]
}





func (dRoFccaZ *k7wdaWX) Less(eaAUaB7Z, kVgtfLFZ int) bool {
	EiKAAX := dRoFccaZ.kWHq8tR[eaAUaB7Z].GdouSNq80bRw
	FtOSI4u3 := dRoFccaZ.kWHq8tR[kVgtfLFZ].GdouSNq80bRw
	LFldiWsD3 := dRoFccaZ.dlMmW5XHOv7[eaAUaB7Z]
	BrQey2UhT0hK := dRoFccaZ.dlMmW5XHOv7[kVgtfLFZ]
	i5BFQrMgkbo8 := &dRoFccaZ.tGRFVVqj[eaAUaB7Z]
	bEXBoC7 := &dRoFccaZ.tGRFVVqj[kVgtfLFZ]
	yamJv8 := &dRoFccaZ.mULghj2IkNE[eaAUaB7Z]
	b1taieqk2x := &dRoFccaZ.mULghj2IkNE[kVgtfLFZ]

	const preferDA = true
	const preferDB = false

	if ! /*line :1*/ LFldiWsD3.IsValid() && ! /*line :1*/ BrQey2UhT0hK.IsValid() {
		return false
	}
	if ! /*line :1*/ BrQey2UhT0hK.IsValid() {
		return preferDA
	}
	if ! /*line :1*/ LFldiWsD3.IsValid() {
		return preferDB
	}

	if i5BFQrMgkbo8.UfOjaDEY == yamJv8.UfOjaDEY && bEXBoC7.UfOjaDEY != b1taieqk2x.UfOjaDEY {
		return preferDA
	}
	if i5BFQrMgkbo8.UfOjaDEY != yamJv8.UfOjaDEY && bEXBoC7.UfOjaDEY == b1taieqk2x.UfOjaDEY {
		return preferDB
	}

	if yamJv8.Q_5GCk03fL9f == i5BFQrMgkbo8.Q_5GCk03fL9f &&
		b1taieqk2x.Q_5GCk03fL9f != bEXBoC7.Q_5GCk03fL9f {
		return preferDA
	}
	if yamJv8.Q_5GCk03fL9f != i5BFQrMgkbo8.Q_5GCk03fL9f &&
		b1taieqk2x.Q_5GCk03fL9f == bEXBoC7.Q_5GCk03fL9f {
		return preferDB
	}

	if i5BFQrMgkbo8.U2NHFCQ2_wC > bEXBoC7.U2NHFCQ2_wC {
		return preferDA
	}
	if i5BFQrMgkbo8.U2NHFCQ2_wC < bEXBoC7.U2NHFCQ2_wC {
		return preferDB
	}

	if i5BFQrMgkbo8.UfOjaDEY < bEXBoC7.UfOjaDEY {
		return preferDA
	}
	if i5BFQrMgkbo8.UfOjaDEY > bEXBoC7.UfOjaDEY {
		return preferDB
	}

	if  /*line :1*/ EiKAAX.To4() == nil &&  /*line :1*/ FtOSI4u3.To4() == nil {
		h45WZxVQa :=  /*line :1*/ srpKFVp9zo(LFldiWsD3, EiKAAX)
		emlgP89YF7 :=  /*line :1*/ srpKFVp9zo(BrQey2UhT0hK, FtOSI4u3)

		if h45WZxVQa > emlgP89YF7 {
			return preferDA
		}
		if h45WZxVQa < emlgP89YF7 {
			return preferDB
		}
	}

	return false
}

type k0jY0zR9if struct {
	VFYQvsqlv5M	netip.We_DLeUDXHez
	CeOWdP4K	uint8
	K3cDvWkJ0pTa	uint8
}

type wmEkYY []k0jY0zR9if



var b9ldTQ7gISd = wmEkYY{
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}), 128),
		CeOWdP4K:	50,
		K3cDvWkJ0pTa:	0,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}), 96),
		CeOWdP4K:	35,
		K3cDvWkJ0pTa:	4,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{}), 96),
		CeOWdP4K:	1,
		K3cDvWkJ0pTa:	3,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{0x20, 0x01}), 32),
		CeOWdP4K:	5,
		K3cDvWkJ0pTa:	5,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{0x20, 0x02}), 16),
		CeOWdP4K:	30,
		K3cDvWkJ0pTa:	2,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{0x3f, 0xfe}), 16),
		CeOWdP4K:	1,
		K3cDvWkJ0pTa:	12,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{0xfe, 0xc0}), 10),
		CeOWdP4K:	1,
		K3cDvWkJ0pTa:	11,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{0xfc}), 7),
		CeOWdP4K:	3,
		K3cDvWkJ0pTa:	13,
	},
	{

		VFYQvsqlv5M:	 /*line :1*/ netip.QkGHFQ3cd( /*line :1*/ netip.XtvbgBR([16]byte{}), 0),
		CeOWdP4K:	40,
		K3cDvWkJ0pTa:	1,
	},
}




func (aeaqWzxJu wmEkYY) Classify(kQ8_UEhxU netip.KFLQ1_1E) k0jY0zR9if {

	if  /*line :1*/ kQ8_UEhxU.Is4() {
		kQ8_UEhxU =  /*line :1*/ netip.XtvbgBR( /*line :1*/ kQ8_UEhxU.As16())
	}
	for _, hD_kB0fwayT := range aeaqWzxJu {
		if  /*line :1*/ hD_kB0fwayT.VFYQvsqlv5M.Contains(kQ8_UEhxU) {
			return hD_kB0fwayT
		}
	}
	return k0jY0zR9if{}
}


type sJQNyGJDJZie uint8

const (
	scopeInterfaceLocal	sJQNyGJDJZie	= 0x1
	scopeLinkLocal	sJQNyGJDJZie	= 0x2
	scopeAdminLocal	sJQNyGJDJZie	= 0x4
	scopeSiteLocal	sJQNyGJDJZie	= 0x5
	scopeOrgLocal	sJQNyGJDJZie	= 0x8
	scopeGlobal	sJQNyGJDJZie	= 0xe
)

func f30RR4U_X(kQ8_UEhxU netip.KFLQ1_1E) sJQNyGJDJZie {
	if  /*line :1*/ kQ8_UEhxU.IsLoopback() ||  /*line :1*/ kQ8_UEhxU.IsLinkLocalUnicast() {
		return scopeLinkLocal
	}
	gPjJWV :=  /*line :1*/ kQ8_UEhxU.Is6() && ! /*line :1*/ kQ8_UEhxU.Is4In6()
	l9nHmK :=  /*line :1*/ kQ8_UEhxU.As16()
	if gPjJWV &&  /*line :1*/ kQ8_UEhxU.IsMulticast() {
		return  /*line :1*/ sJQNyGJDJZie(l9nHmK[1] & 0xf)
	}

	if gPjJWV && l9nHmK[0] == 0xfe && l9nHmK[1]&0xc0 == 0xc0 {
		return scopeSiteLocal
	}
	return scopeGlobal
}











func srpKFVp9zo(uI7LZDHm6 netip.KFLQ1_1E, fIadEXIim6l GNraIvYhB) (jwiQzFK int) {
	if wVyRZ0 :=  /*line :1*/ fIadEXIim6l.To4(); wVyRZ0 != nil {
		fIadEXIim6l = wVyRZ0
	}
	tfMcw2 :=  /*line :1*/ uI7LZDHm6.AsSlice()
	if  /*line :1*/ len(tfMcw2) !=  /*line :1*/ len(fIadEXIim6l) {
		return 0
	}

	if  /*line :1*/ len(tfMcw2) > 8 {
		tfMcw2 = tfMcw2[:8]
		fIadEXIim6l = fIadEXIim6l[:8]
	}
	for  /*line :1*/ len(tfMcw2) > 0 {
		if tfMcw2[0] == fIadEXIim6l[0] {
			jwiQzFK += 8
			tfMcw2 = tfMcw2[1:]
			fIadEXIim6l = fIadEXIim6l[1:]
			continue
		}
		ikVsUbWnbH := 8
		xKTpMnL, jIdo8ln6 := tfMcw2[0], fIadEXIim6l[0]
		for {
			xKTpMnL >>= 1
			jIdo8ln6 >>= 1
			ikVsUbWnbH--
			if xKTpMnL == jIdo8ln6 {
				jwiQzFK += ikVsUbWnbH
				return
			}
		}
	}
	return
}
