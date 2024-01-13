//line :1
package bUKeamF

import (
	"unsafe"
)

func ITYf1B(nIB9uB1 string) (fsp_0a7i string, fLFxXyJoo bool) {
	c7XpXU, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(nIB9uB1)
	if gOCcQzbcC != nil {
		return "", false
	}
	m5Tq_PL7 :=  /*line :1*/ uint32(100)
	for {
		tHPwhuTh :=  /*line :1*/ make([]uint16, m5Tq_PL7)
		m5Tq_PL7, gOCcQzbcC =  /*line :1*/ UVmkZzea(c7XpXU, &tHPwhuTh[0],  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)))
		if m5Tq_PL7 == 0 && gOCcQzbcC == ERROR_ENVVAR_NOT_FOUND {
			return "", false
		}
		if m5Tq_PL7 <=  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)) {
			return  /*line :1*/ AODVXp8NM3sd(tHPwhuTh[:m5Tq_PL7]), true
		}
	}
}

func BNcaDu(nIB9uB1, fsp_0a7i string) error {
	b0abCyNZ, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(fsp_0a7i)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	c7XpXU, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(nIB9uB1)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	wAwFLr4AGHg :=  /*line :1*/ EkkQwxRH(c7XpXU, b0abCyNZ)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	 /*line :1*/ gb9l485q7(nIB9uB1, fsp_0a7i)
	return nil
}

func OuD_oLNMM(nIB9uB1 string) error {
	c7XpXU, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(nIB9uB1)
	if gOCcQzbcC != nil {
		return gOCcQzbcC
	}
	wAwFLr4AGHg :=  /*line :1*/ EkkQwxRH(c7XpXU, nil)
	if wAwFLr4AGHg != nil {
		return wAwFLr4AGHg
	}
	 /*line :1*/ aCtW2ffDj(nIB9uB1)
	return nil
}

func J1XgGX_nmS4b() {
	for _, wzPhJhIFoI := range  /*line :1*/ FDhTFlUY() {

		for oUmRMCO := 1; oUmRMCO <  /*line :1*/ len(wzPhJhIFoI); oUmRMCO++ {
			if wzPhJhIFoI[oUmRMCO] == '=' {
				 /*line :1*/ OuD_oLNMM(wzPhJhIFoI[0:oUmRMCO])
				break
			}
		}
	}
}

func FDhTFlUY() []string {
	c0cfQv, wAwFLr4AGHg :=  /*line :1*/ ZEq5mFM20k()
	if wAwFLr4AGHg != nil {
		return nil
	}
	defer  /*line :1*/ CP4mgHYW5e8D(c0cfQv)

	iVSjWu :=  /*line :1*/ make([]string, 0, 50)
	const size =  /*line :1*/ unsafe.Sizeof(*c0cfQv)
	for *c0cfQv != 0 {

		iMDTBTj :=  /*line :1*/ unsafe.Pointer(c0cfQv)
		for *(* /*line :1*/ uint16)(iMDTBTj) != 0 {
			iMDTBTj =  /*line :1*/ unsafe.Add(iMDTBTj, size)
		}

		qp0onb :=  /*line :1*/ unsafe.Slice(c0cfQv, ( /*line :1*/ uintptr(iMDTBTj)- /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(c0cfQv)))/size)
		iVSjWu =  /*line :1*/ append(iVSjWu,  /*line :1*/ AODVXp8NM3sd(qp0onb))
		c0cfQv = (* /*line :1*/ uint16)( /*line :1*/ unsafe.Add(iMDTBTj, size))
	}
	return iVSjWu
}
