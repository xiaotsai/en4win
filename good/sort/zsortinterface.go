//line :1
package gzHaha55n


func uSheIae(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW int) {
	for eLewhDqfw8 := bSi_aMm2pI + 1; eLewhDqfw8 < d5a0WORR8UjW; eLewhDqfw8++ {
		for kCr_PWXwcND9 := eLewhDqfw8; kCr_PWXwcND9 > bSi_aMm2pI &&  /*line :1*/ cxvZaD16.Less(kCr_PWXwcND9, kCr_PWXwcND9-1); kCr_PWXwcND9-- {
			 /*line :1*/ cxvZaD16.Swap(kCr_PWXwcND9, kCr_PWXwcND9-1)
		}
	}
}



func phRQ9mVOk4VT(cxvZaD16 AOVoo3y622nl, oYrMmstw, ayqF0Ba, qIphJJ int) {
	ppaqzLxKKjo := oYrMmstw
	for {
		z2HNVV8PA6MJ := 2*ppaqzLxKKjo + 1
		if z2HNVV8PA6MJ >= ayqF0Ba {
			break
		}
		if z2HNVV8PA6MJ+1 < ayqF0Ba &&  /*line :1*/ cxvZaD16.Less(qIphJJ+z2HNVV8PA6MJ, qIphJJ+z2HNVV8PA6MJ+1) {
			z2HNVV8PA6MJ++
		}
		if ! /*line :1*/ cxvZaD16.Less(qIphJJ+ppaqzLxKKjo, qIphJJ+z2HNVV8PA6MJ) {
			return
		}
		 /*line :1*/ cxvZaD16.Swap(qIphJJ+ppaqzLxKKjo, qIphJJ+z2HNVV8PA6MJ)
		ppaqzLxKKjo = z2HNVV8PA6MJ
	}
}

func mWiJ4c9gyw0(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW int) {
	qIphJJ := bSi_aMm2pI
	oYrMmstw := 0
	ayqF0Ba := d5a0WORR8UjW - bSi_aMm2pI

	for eLewhDqfw8 := (ayqF0Ba - 1) / 2; eLewhDqfw8 >= 0; eLewhDqfw8-- {
		 /*line :1*/ phRQ9mVOk4VT(cxvZaD16, eLewhDqfw8, ayqF0Ba, qIphJJ)
	}

	for eLewhDqfw8 := ayqF0Ba - 1; eLewhDqfw8 >= 0; eLewhDqfw8-- {
		 /*line :1*/ cxvZaD16.Swap(qIphJJ, qIphJJ+eLewhDqfw8)
		 /*line :1*/ phRQ9mVOk4VT(cxvZaD16, oYrMmstw, eLewhDqfw8, qIphJJ)
	}
}







func mUdU8k(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW, l5vzUBwzefa int) {
	const maxInsertion = 12

	var (
		muax1a	= true		
		n3862RdRr	= true		
	)

	for {
		kbaqw6D := d5a0WORR8UjW - bSi_aMm2pI

		if kbaqw6D <= maxInsertion {
			 /*line :1*/ uSheIae(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW)
			return
		}

		if l5vzUBwzefa == 0 {
			 /*line :1*/ mWiJ4c9gyw0(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW)
			return
		}

		if !muax1a {
			 /*line :1*/ mLDNz8kt1f(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW)
			l5vzUBwzefa--
		}

		r9TiL1JXaJ, e8d6MEs66 :=  /*line :1*/ jdjUy4aNF6(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW)
		if e8d6MEs66 == decreasingHint {
			 /*line :1*/ jCNqSN3OaZ5Q(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW)

			r9TiL1JXaJ = (d5a0WORR8UjW - 1) - (r9TiL1JXaJ - bSi_aMm2pI)
			e8d6MEs66 = increasingHint
		}

		if muax1a && n3862RdRr && e8d6MEs66 == increasingHint {
			if  /*line :1*/ yGVT9H4zGW_G(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW) {
				return
			}
		}

		if bSi_aMm2pI > 0 && ! /*line :1*/ cxvZaD16.Less(bSi_aMm2pI-1, r9TiL1JXaJ) {
			h6mOEPIyCq :=  /*line :1*/ xebV2O(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW, r9TiL1JXaJ)
			bSi_aMm2pI = h6mOEPIyCq
			continue
		}

		h6mOEPIyCq, fvZGZYqQj8Lr :=  /*line :1*/ z0yWxCQk(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW, r9TiL1JXaJ)
		n3862RdRr = fvZGZYqQj8Lr

		oq5gt0whTPy, hDhVNzv := h6mOEPIyCq-bSi_aMm2pI, d5a0WORR8UjW-h6mOEPIyCq
		b0En6zsQi := kbaqw6D / 8
		if oq5gt0whTPy < hDhVNzv {
			muax1a = oq5gt0whTPy >= b0En6zsQi
			 /*line :1*/ mUdU8k(cxvZaD16, bSi_aMm2pI, h6mOEPIyCq, l5vzUBwzefa)
			bSi_aMm2pI = h6mOEPIyCq + 1
		} else {
			muax1a = hDhVNzv >= b0En6zsQi
			 /*line :1*/ mUdU8k(cxvZaD16, h6mOEPIyCq+1, d5a0WORR8UjW, l5vzUBwzefa)
			d5a0WORR8UjW = h6mOEPIyCq
		}
	}
}





func z0yWxCQk(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW, r9TiL1JXaJ int) (i0yZ3WzkXq int, fvZGZYqQj8Lr bool) {
	 /*line :1*/ cxvZaD16.Swap(bSi_aMm2pI, r9TiL1JXaJ)
	eLewhDqfw8, kCr_PWXwcND9 := bSi_aMm2pI+1, d5a0WORR8UjW-1

	for eLewhDqfw8 <= kCr_PWXwcND9 &&  /*line :1*/ cxvZaD16.Less(eLewhDqfw8, bSi_aMm2pI) {
		eLewhDqfw8++
	}
	for eLewhDqfw8 <= kCr_PWXwcND9 && ! /*line :1*/ cxvZaD16.Less(kCr_PWXwcND9, bSi_aMm2pI) {
		kCr_PWXwcND9--
	}
	if eLewhDqfw8 > kCr_PWXwcND9 {
		 /*line :1*/ cxvZaD16.Swap(kCr_PWXwcND9, bSi_aMm2pI)
		return kCr_PWXwcND9, true
	}
	 /*line :1*/ cxvZaD16.Swap(eLewhDqfw8, kCr_PWXwcND9)
	eLewhDqfw8++
	kCr_PWXwcND9--

	for {
		for eLewhDqfw8 <= kCr_PWXwcND9 &&  /*line :1*/ cxvZaD16.Less(eLewhDqfw8, bSi_aMm2pI) {
			eLewhDqfw8++
		}
		for eLewhDqfw8 <= kCr_PWXwcND9 && ! /*line :1*/ cxvZaD16.Less(kCr_PWXwcND9, bSi_aMm2pI) {
			kCr_PWXwcND9--
		}
		if eLewhDqfw8 > kCr_PWXwcND9 {
			break
		}
		 /*line :1*/ cxvZaD16.Swap(eLewhDqfw8, kCr_PWXwcND9)
		eLewhDqfw8++
		kCr_PWXwcND9--
	}
	 /*line :1*/ cxvZaD16.Swap(kCr_PWXwcND9, bSi_aMm2pI)
	return kCr_PWXwcND9, false
}



func xebV2O(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW, r9TiL1JXaJ int) (i0yZ3WzkXq int) {
	 /*line :1*/ cxvZaD16.Swap(bSi_aMm2pI, r9TiL1JXaJ)
	eLewhDqfw8, kCr_PWXwcND9 := bSi_aMm2pI+1, d5a0WORR8UjW-1

	for {
		for eLewhDqfw8 <= kCr_PWXwcND9 && ! /*line :1*/ cxvZaD16.Less(bSi_aMm2pI, eLewhDqfw8) {
			eLewhDqfw8++
		}
		for eLewhDqfw8 <= kCr_PWXwcND9 &&  /*line :1*/ cxvZaD16.Less(bSi_aMm2pI, kCr_PWXwcND9) {
			kCr_PWXwcND9--
		}
		if eLewhDqfw8 > kCr_PWXwcND9 {
			break
		}
		 /*line :1*/ cxvZaD16.Swap(eLewhDqfw8, kCr_PWXwcND9)
		eLewhDqfw8++
		kCr_PWXwcND9--
	}
	return eLewhDqfw8
}


func yGVT9H4zGW_G(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW int) bool {
	const (
		maxSteps	= 5		
		shortestShifting	= 50		
	)
	eLewhDqfw8 := bSi_aMm2pI + 1
	for kCr_PWXwcND9 := 0; kCr_PWXwcND9 < maxSteps; kCr_PWXwcND9++ {
		for eLewhDqfw8 < d5a0WORR8UjW && ! /*line :1*/ cxvZaD16.Less(eLewhDqfw8, eLewhDqfw8-1) {
			eLewhDqfw8++
		}

		if eLewhDqfw8 == d5a0WORR8UjW {
			return true
		}

		if d5a0WORR8UjW-bSi_aMm2pI < shortestShifting {
			return false
		}

		 /*line :1*/ cxvZaD16.Swap(eLewhDqfw8, eLewhDqfw8-1)

		if eLewhDqfw8-bSi_aMm2pI >= 2 {
			for kCr_PWXwcND9 := eLewhDqfw8 - 1; kCr_PWXwcND9 >= 1; kCr_PWXwcND9-- {
				if ! /*line :1*/ cxvZaD16.Less(kCr_PWXwcND9, kCr_PWXwcND9-1) {
					break
				}
				 /*line :1*/ cxvZaD16.Swap(kCr_PWXwcND9, kCr_PWXwcND9-1)
			}
		}

		if d5a0WORR8UjW-eLewhDqfw8 >= 2 {
			for kCr_PWXwcND9 := eLewhDqfw8 + 1; kCr_PWXwcND9 < d5a0WORR8UjW; kCr_PWXwcND9++ {
				if ! /*line :1*/ cxvZaD16.Less(kCr_PWXwcND9, kCr_PWXwcND9-1) {
					break
				}
				 /*line :1*/ cxvZaD16.Swap(kCr_PWXwcND9, kCr_PWXwcND9-1)
			}
		}
	}
	return false
}



func mLDNz8kt1f(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW int) {
	kbaqw6D := d5a0WORR8UjW - bSi_aMm2pI
	if kbaqw6D >= 8 {
		oOjNrFD :=  /*line :1*/ aMwO4FPaJS3s(kbaqw6D)
		_WupPV4 :=  /*line :1*/ boBGv1E(kbaqw6D)

		for jC0zuLrkMyT := bSi_aMm2pI + (kbaqw6D/4)*2 - 1; jC0zuLrkMyT <= bSi_aMm2pI+(kbaqw6D/4)*2+1; jC0zuLrkMyT++ {
			fG0Gc8u :=  /*line :1*/ int( /*line :1*/ uint( /*line :1*/ oOjNrFD.Next()) & (_WupPV4 - 1))
			if fG0Gc8u >= kbaqw6D {
				fG0Gc8u -= kbaqw6D
			}
			 /*line :1*/ cxvZaD16.Swap(jC0zuLrkMyT, bSi_aMm2pI+fG0Gc8u)
		}
	}
}






func jdjUy4aNF6(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW int) (r9TiL1JXaJ int, e8d6MEs66 ea2AAb0t) {
	const (
		shortestNinther	= 50
		maxSwaps	= 4 * 3
	)

	ai9rx9OCDa := d5a0WORR8UjW - bSi_aMm2pI

	var (
		qp2kOO3gjR	int
		eLewhDqfw8	= bSi_aMm2pI + ai9rx9OCDa/4*1
		kCr_PWXwcND9	= bSi_aMm2pI + ai9rx9OCDa/4*2
		pPAAdGTa	= bSi_aMm2pI + ai9rx9OCDa/4*3
	)

	if ai9rx9OCDa >= 8 {
		if ai9rx9OCDa >= shortestNinther {

			eLewhDqfw8 =  /*line :1*/ c3Kg4ORGwov(cxvZaD16, eLewhDqfw8, &qp2kOO3gjR)
			kCr_PWXwcND9 =  /*line :1*/ c3Kg4ORGwov(cxvZaD16, kCr_PWXwcND9, &qp2kOO3gjR)
			pPAAdGTa =  /*line :1*/ c3Kg4ORGwov(cxvZaD16, pPAAdGTa, &qp2kOO3gjR)
		}

		kCr_PWXwcND9 =  /*line :1*/ iD0G5mp4EC(cxvZaD16, eLewhDqfw8, kCr_PWXwcND9, pPAAdGTa, &qp2kOO3gjR)
	}

	switch qp2kOO3gjR {
	case 0:
		return kCr_PWXwcND9, increasingHint
	case maxSwaps:
		return kCr_PWXwcND9, decreasingHint
	default:
		return kCr_PWXwcND9, unknownHint
	}
}


func wahgKa(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW int, qp2kOO3gjR *int) (int, int) {
	if  /*line :1*/ cxvZaD16.Less(d5a0WORR8UjW, bSi_aMm2pI) {
		*qp2kOO3gjR++
		return d5a0WORR8UjW, bSi_aMm2pI
	}
	return bSi_aMm2pI, d5a0WORR8UjW
}


func iD0G5mp4EC(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW, dMH4Gay int, qp2kOO3gjR *int) int {
	bSi_aMm2pI, d5a0WORR8UjW =  /*line :1*/ wahgKa(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW, qp2kOO3gjR)
	d5a0WORR8UjW, dMH4Gay =  /*line :1*/ wahgKa(cxvZaD16, d5a0WORR8UjW, dMH4Gay, qp2kOO3gjR)
	bSi_aMm2pI, d5a0WORR8UjW =  /*line :1*/ wahgKa(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW, qp2kOO3gjR)
	return d5a0WORR8UjW
}


func c3Kg4ORGwov(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI int, qp2kOO3gjR *int) int {
	return  /*line :1*/ iD0G5mp4EC(cxvZaD16, bSi_aMm2pI-1, bSi_aMm2pI, bSi_aMm2pI+1, qp2kOO3gjR)
}

func jCNqSN3OaZ5Q(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW int) {
	eLewhDqfw8 := bSi_aMm2pI
	kCr_PWXwcND9 := d5a0WORR8UjW - 1
	for eLewhDqfw8 < kCr_PWXwcND9 {
		 /*line :1*/ cxvZaD16.Swap(eLewhDqfw8, kCr_PWXwcND9)
		eLewhDqfw8++
		kCr_PWXwcND9--
	}
}

func nGOUW9T5A(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, d5a0WORR8UjW, bntF0l int) {
	for eLewhDqfw8 := 0; eLewhDqfw8 < bntF0l; eLewhDqfw8++ {
		 /*line :1*/ cxvZaD16.Swap(bSi_aMm2pI+eLewhDqfw8, d5a0WORR8UjW+eLewhDqfw8)
	}
}

func fYPxs1h(cxvZaD16 AOVoo3y622nl, bntF0l int) {
	aw1elG := 20
	bSi_aMm2pI, d5a0WORR8UjW := 0, aw1elG
	for d5a0WORR8UjW <= bntF0l {
		 /*line :1*/ uSheIae(cxvZaD16, bSi_aMm2pI, d5a0WORR8UjW)
		bSi_aMm2pI = d5a0WORR8UjW
		d5a0WORR8UjW += aw1elG
	}
	 /*line :1*/ uSheIae(cxvZaD16, bSi_aMm2pI, bntF0l)

	for aw1elG < bntF0l {
		bSi_aMm2pI, d5a0WORR8UjW = 0, 2*aw1elG
		for d5a0WORR8UjW <= bntF0l {
			 /*line :1*/ p8c_jRq6Ot(cxvZaD16, bSi_aMm2pI, bSi_aMm2pI+aw1elG, d5a0WORR8UjW)
			bSi_aMm2pI = d5a0WORR8UjW
			d5a0WORR8UjW += 2 * aw1elG
		}
		if g_x1FP := bSi_aMm2pI + aw1elG; g_x1FP < bntF0l {
			 /*line :1*/ p8c_jRq6Ot(cxvZaD16, bSi_aMm2pI, g_x1FP, bntF0l)
		}
		aw1elG *= 2
	}
}




















func p8c_jRq6Ot(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, g_x1FP, d5a0WORR8UjW int) {

	if g_x1FP-bSi_aMm2pI == 1 {

		eLewhDqfw8 := g_x1FP
		kCr_PWXwcND9 := d5a0WORR8UjW
		for eLewhDqfw8 < kCr_PWXwcND9 {
			xcoKpNDa :=  /*line :1*/ int( /*line :1*/ uint(eLewhDqfw8+kCr_PWXwcND9) >> 1)
			if  /*line :1*/ cxvZaD16.Less(xcoKpNDa, bSi_aMm2pI) {
				eLewhDqfw8 = xcoKpNDa + 1
			} else {
				kCr_PWXwcND9 = xcoKpNDa
			}
		}

		for pPAAdGTa := bSi_aMm2pI; pPAAdGTa < eLewhDqfw8-1; pPAAdGTa++ {
			 /*line :1*/ cxvZaD16.Swap(pPAAdGTa, pPAAdGTa+1)
		}
		return
	}

	if d5a0WORR8UjW-g_x1FP == 1 {

		eLewhDqfw8 := bSi_aMm2pI
		kCr_PWXwcND9 := g_x1FP
		for eLewhDqfw8 < kCr_PWXwcND9 {
			xcoKpNDa :=  /*line :1*/ int( /*line :1*/ uint(eLewhDqfw8+kCr_PWXwcND9) >> 1)
			if ! /*line :1*/ cxvZaD16.Less(g_x1FP, xcoKpNDa) {
				eLewhDqfw8 = xcoKpNDa + 1
			} else {
				kCr_PWXwcND9 = xcoKpNDa
			}
		}

		for pPAAdGTa := g_x1FP; pPAAdGTa > eLewhDqfw8; pPAAdGTa-- {
			 /*line :1*/ cxvZaD16.Swap(pPAAdGTa, pPAAdGTa-1)
		}
		return
	}

	h6mOEPIyCq :=  /*line :1*/ int( /*line :1*/ uint(bSi_aMm2pI+d5a0WORR8UjW) >> 1)
	bntF0l := h6mOEPIyCq + g_x1FP
	var gLJt5PgUPlY, n6aYLO8 int
	if g_x1FP > h6mOEPIyCq {
		gLJt5PgUPlY = bntF0l - d5a0WORR8UjW
		n6aYLO8 = h6mOEPIyCq
	} else {
		gLJt5PgUPlY = bSi_aMm2pI
		n6aYLO8 = g_x1FP
	}
	zViw4r := bntF0l - 1

	for gLJt5PgUPlY < n6aYLO8 {
		dMH4Gay :=  /*line :1*/ int( /*line :1*/ uint(gLJt5PgUPlY+n6aYLO8) >> 1)
		if ! /*line :1*/ cxvZaD16.Less(zViw4r-dMH4Gay, dMH4Gay) {
			gLJt5PgUPlY = dMH4Gay + 1
		} else {
			n6aYLO8 = dMH4Gay
		}
	}

	antqRkf := bntF0l - gLJt5PgUPlY
	if gLJt5PgUPlY < g_x1FP && g_x1FP < antqRkf {
		 /*line :1*/ nKP1kFgN2Nx(cxvZaD16, gLJt5PgUPlY, g_x1FP, antqRkf)
	}
	if bSi_aMm2pI < gLJt5PgUPlY && gLJt5PgUPlY < h6mOEPIyCq {
		 /*line :1*/ p8c_jRq6Ot(cxvZaD16, bSi_aMm2pI, gLJt5PgUPlY, h6mOEPIyCq)
	}
	if h6mOEPIyCq < antqRkf && antqRkf < d5a0WORR8UjW {
		 /*line :1*/ p8c_jRq6Ot(cxvZaD16, h6mOEPIyCq, antqRkf, d5a0WORR8UjW)
	}
}





func nKP1kFgN2Nx(cxvZaD16 AOVoo3y622nl, bSi_aMm2pI, g_x1FP, d5a0WORR8UjW int) {
	eLewhDqfw8 := g_x1FP - bSi_aMm2pI
	kCr_PWXwcND9 := d5a0WORR8UjW - g_x1FP

	for eLewhDqfw8 != kCr_PWXwcND9 {
		if eLewhDqfw8 > kCr_PWXwcND9 {
			 /*line :1*/ nGOUW9T5A(cxvZaD16, g_x1FP-eLewhDqfw8, g_x1FP, kCr_PWXwcND9)
			eLewhDqfw8 -= kCr_PWXwcND9
		} else {
			 /*line :1*/ nGOUW9T5A(cxvZaD16, g_x1FP-eLewhDqfw8, g_x1FP+kCr_PWXwcND9-eLewhDqfw8, eLewhDqfw8)
			kCr_PWXwcND9 -= eLewhDqfw8
		}
	}

	 /*line :1*/ nGOUW9T5A(cxvZaD16, g_x1FP-eLewhDqfw8, g_x1FP, eLewhDqfw8)
}
