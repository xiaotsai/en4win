//line :1
package NJ4MerJ

import (
	errorspkg "iAHoxjmM"
	"unsafe"
)










func AnJa605Ytw4(bamc83qA3DBr string) string {
	if  /*line :1*/ len(bamc83qA3DBr) == 0 {
		return `""`
	}
	krzuku :=  /*line :1*/ len(bamc83qA3DBr)
	xUgr2bcTn := false
	for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(bamc83qA3DBr); rRGfxgPH8Kq++ {
		switch bamc83qA3DBr[rRGfxgPH8Kq] {
		case '"', '\\':
			krzuku++
		case ' ', '\t':
			xUgr2bcTn = true
		}
	}
	if xUgr2bcTn {
		krzuku += 2
	}
	if krzuku ==  /*line :1*/ len(bamc83qA3DBr) {
		return bamc83qA3DBr
	}

	eTOvUG49GOdM :=  /*line :1*/ make([]byte, krzuku)
	hFiUmivkwHK := 0
	if xUgr2bcTn {
		eTOvUG49GOdM[hFiUmivkwHK] = '"'
		hFiUmivkwHK++
	}
	xXl1DY5 := 0
	for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(bamc83qA3DBr); rRGfxgPH8Kq++ {
		switch bamc83qA3DBr[rRGfxgPH8Kq] {
		default:
			xXl1DY5 = 0
			eTOvUG49GOdM[hFiUmivkwHK] = bamc83qA3DBr[rRGfxgPH8Kq]
		case '\\':
			xXl1DY5++
			eTOvUG49GOdM[hFiUmivkwHK] = bamc83qA3DBr[rRGfxgPH8Kq]
		case '"':
			for ; xXl1DY5 > 0; xXl1DY5-- {
				eTOvUG49GOdM[hFiUmivkwHK] = '\\'
				hFiUmivkwHK++
			}
			eTOvUG49GOdM[hFiUmivkwHK] = '\\'
			hFiUmivkwHK++
			eTOvUG49GOdM[hFiUmivkwHK] = bamc83qA3DBr[rRGfxgPH8Kq]
		}
		hFiUmivkwHK++
	}
	if xUgr2bcTn {
		for ; xXl1DY5 > 0; xXl1DY5-- {
			eTOvUG49GOdM[hFiUmivkwHK] = '\\'
			hFiUmivkwHK++
		}
		eTOvUG49GOdM[hFiUmivkwHK] = '"'
		hFiUmivkwHK++
	}
	return  /*line :1*/ string(eTOvUG49GOdM[:hFiUmivkwHK])
}




func ZPoEDWm8YG(feOynBEj []string) string {
	if  /*line :1*/ len(feOynBEj) == 0 {
		return ""
	}

	p1G1bwh6mbvw := feOynBEj[0]
	jTss6E :=  /*line :1*/ len(p1G1bwh6mbvw) == 0
	for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(p1G1bwh6mbvw); rRGfxgPH8Kq++ {
		jpk9X7z := p1G1bwh6mbvw[rRGfxgPH8Kq]
		if jpk9X7z <= ' ' || (jpk9X7z == '"' && rRGfxgPH8Kq == 0) {

			jTss6E = true
			break
		}
	}
	var zXn7QsQIeo []byte
	if jTss6E {
		zXn7QsQIeo =  /*line :1*/ make([]byte, 0,  /*line :1*/ len(p1G1bwh6mbvw)+2)
		zXn7QsQIeo =  /*line :1*/ append(zXn7QsQIeo, '"')
		for rRGfxgPH8Kq := 0; rRGfxgPH8Kq <  /*line :1*/ len(p1G1bwh6mbvw); rRGfxgPH8Kq++ {
			jpk9X7z := p1G1bwh6mbvw[rRGfxgPH8Kq]
			if jpk9X7z == '"' {

				continue
			}
			zXn7QsQIeo =  /*line :1*/ append(zXn7QsQIeo, jpk9X7z)
		}
		zXn7QsQIeo =  /*line :1*/ append(zXn7QsQIeo, '"')
	} else {
		if  /*line :1*/ len(feOynBEj) == 1 {

			return p1G1bwh6mbvw
		}
		zXn7QsQIeo = [] /*line :1*/ byte(p1G1bwh6mbvw)
	}

	for _, zTQfCwn71 := range feOynBEj[1:] {
		zXn7QsQIeo =  /*line :1*/ append(zXn7QsQIeo, ' ')

		zXn7QsQIeo =  /*line :1*/ append(zXn7QsQIeo,  /*line :1*/ AnJa605Ytw4(zTQfCwn71)...)
	}
	return  /*line :1*/ string(zXn7QsQIeo)
}





func KEjdvdqv2VqM(zXn7QsQIeo string) ([]string, error) {
	if  /*line :1*/ len(zXn7QsQIeo) == 0 {
		return []string{}, nil
	}
	wFaWufx9KAV, jeWMpOaQtWV :=  /*line :1*/ DJN4EMpp7(zXn7QsQIeo)
	if jeWMpOaQtWV != nil {
		return nil,  /*line :1*/ errorspkg.Su6g6hRoi9X("string with NUL passed to DecomposeCommandLine")
	}
	var pjMGTnPK int32
	vPivxOaBYBAf, jeWMpOaQtWV :=  /*line :1*/ aYQ3K6eKrSQ(&wFaWufx9KAV[0], &pjMGTnPK)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	defer  /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(vPivxOaBYBAf)))

	var feOynBEj []string
	for _, hD4wNgEB := range  /*line :1*/ unsafe.Slice(vPivxOaBYBAf, pjMGTnPK) {
		feOynBEj =  /*line :1*/ append(feOynBEj,  /*line :1*/ M_Sea9j(hD4wNgEB))
	}
	return feOynBEj, nil
}











func XTHJHb(aGRCzBj3 *uint16, pjMGTnPK *int32) (vPivxOaBYBAf *[8192]*[8192]uint16, jeWMpOaQtWV error) {
	sfXhsERl, jeWMpOaQtWV :=  /*line :1*/ aYQ3K6eKrSQ(aGRCzBj3, pjMGTnPK)
	vPivxOaBYBAf = (*[8192]*[8192] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(sfXhsERl))
	return vPivxOaBYBAf, jeWMpOaQtWV
}

func WBxG06MUDdV(nuvOdY1SQ9I5 L2L8P9WaNZ) {
	 /*line :1*/ JbtpNRq( /*line :1*/ L2L8P9WaNZ(nuvOdY1SQ9I5), HANDLE_FLAG_INHERIT, 0)
}


func AF_hPD2(gYwswmeUjG string) (bZKbGTxeit string, jeWMpOaQtWV error) {
	hD4wNgEB, jeWMpOaQtWV :=  /*line :1*/ Ih3Y4u(gYwswmeUjG)
	if jeWMpOaQtWV != nil {
		return "", jeWMpOaQtWV
	}
	krzuku :=  /*line :1*/ uint32(100)
	for {
		etRYtnVni :=  /*line :1*/ make([]uint16, krzuku)
		krzuku, jeWMpOaQtWV =  /*line :1*/ IZtvZd6jy53(hD4wNgEB,  /*line :1*/ uint32( /*line :1*/ len(etRYtnVni)), &etRYtnVni[0], nil)
		if jeWMpOaQtWV != nil {
			return "", jeWMpOaQtWV
		}
		if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(etRYtnVni)) {
			return  /*line :1*/ OXNanQ8Uj(etRYtnVni[:krzuku]), nil
		}
	}
}


func LsqEZxj(pk6wEMa uint32) (*B1GK7_cto7, error) {
	var nFtsmqHC uintptr
	jeWMpOaQtWV :=  /*line :1*/ ckL_c4L8(nil, pk6wEMa, 0, &nFtsmqHC)
	if jeWMpOaQtWV != ERROR_INSUFFICIENT_BUFFER {
		if jeWMpOaQtWV == nil {
			return nil,  /*line :1*/ errorspkg.Su6g6hRoi9X("unable to query buffer size from InitializeProcThreadAttributeList")
		}
		return nil, jeWMpOaQtWV
	}
	wfJ2EhPna4, jeWMpOaQtWV :=  /*line :1*/ R_E3wm(LMEM_FIXED,  /*line :1*/ uint32(nFtsmqHC))
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}

	hz8UMk := &B1GK7_cto7{ng5OB170e: (* /*line :1*/ JEfkblUWi)( /*line :1*/ unsafe.Pointer(wfJ2EhPna4))}
	jeWMpOaQtWV =  /*line :1*/ ckL_c4L8(hz8UMk.ng5OB170e, pk6wEMa, 0, &nFtsmqHC)
	if jeWMpOaQtWV != nil {
		return nil, jeWMpOaQtWV
	}
	return hz8UMk, jeWMpOaQtWV
}


func (hz8UMk *B1GK7_cto7) Update(q5eWI0Z1lHex uintptr, wvZBcfh unsafe.Pointer, nFtsmqHC uintptr) error {
	hz8UMk.aevUF6NhkB =  /*line :1*/ append(hz8UMk.aevUF6NhkB, wvZBcfh)
	return  /*line :1*/ lSavoG(hz8UMk.ng5OB170e, 0, q5eWI0Z1lHex, wvZBcfh, nFtsmqHC, nil, nil)
}


func (hz8UMk *B1GK7_cto7) Delete() {
	 /*line :1*/ dOpVlY(hz8UMk.ng5OB170e)
	 /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(hz8UMk.ng5OB170e)))
	hz8UMk.ng5OB170e = nil
	hz8UMk.aevUF6NhkB = nil
}


func (hz8UMk *B1GK7_cto7) List() *JEfkblUWi {
	return hz8UMk.ng5OB170e
}
