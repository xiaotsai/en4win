//line :1
package y1BamVjnXsaa

import (
	time "fRAfQd_"
)







func ThxkDee(v9FHH9DB9sM2 HM4JM2) string {
	hMWDrBfy :=  /*line :1*/ v9FHH9DB9sM2.Name()
	nEQLP7 :=  /*line :1*/ make([]byte, 0, 40+ /*line :1*/ len(hMWDrBfy))
	nEQLP7 =  /*line :1*/ append(nEQLP7,  /*line :1*/ v9FHH9DB9sM2.Mode().String()...)
	nEQLP7 =  /*line :1*/ append(nEQLP7, ' ')

	x9px08Yregqq :=  /*line :1*/ v9FHH9DB9sM2.Size()
	var jaAn1YW7 uint64
	if x9px08Yregqq >= 0 {
		jaAn1YW7 =  /*line :1*/ uint64(x9px08Yregqq)
	} else {
		nEQLP7 =  /*line :1*/ append(nEQLP7, '-')
		jaAn1YW7 =  /*line :1*/ uint64(-x9px08Yregqq)
	}
	var zTLJRkZ1 [20]byte
	aTOpQutlw :=  /*line :1*/ len(zTLJRkZ1) - 1
	for jaAn1YW7 >= 10 {
		vAk_Y3LjW6 := jaAn1YW7 / 10
		zTLJRkZ1[aTOpQutlw] =  /*line :1*/ byte('0' + jaAn1YW7 - vAk_Y3LjW6*10)
		aTOpQutlw--
		jaAn1YW7 = vAk_Y3LjW6
	}
	zTLJRkZ1[aTOpQutlw] =  /*line :1*/ byte('0' + jaAn1YW7)
	nEQLP7 =  /*line :1*/ append(nEQLP7, zTLJRkZ1[aTOpQutlw:]...)
	nEQLP7 =  /*line :1*/ append(nEQLP7, ' ')

	nEQLP7 =  /*line :1*/ append(nEQLP7,  /*line :1*/ v9FHH9DB9sM2.ModTime().Format(time.DateTime)...)
	nEQLP7 =  /*line :1*/ append(nEQLP7, ' ')

	nEQLP7 =  /*line :1*/ append(nEQLP7, hMWDrBfy...)
	if  /*line :1*/ v9FHH9DB9sM2.IsDir() {
		nEQLP7 =  /*line :1*/ append(nEQLP7, '/')
	}

	return  /*line :1*/ string(nEQLP7)
}







func EMHCka(cNaiWwqHDG RbfBMxa) string {
	hMWDrBfy :=  /*line :1*/ cNaiWwqHDG.Name()
	nEQLP7 :=  /*line :1*/ make([]byte, 0, 5+ /*line :1*/ len(hMWDrBfy))

	e7Yht6r :=  /*line :1*/ cNaiWwqHDG.Type().String()
	e7Yht6r = e7Yht6r[: /*line :1*/ len(e7Yht6r)-9]

	nEQLP7 =  /*line :1*/ append(nEQLP7, e7Yht6r...)
	nEQLP7 =  /*line :1*/ append(nEQLP7, ' ')
	nEQLP7 =  /*line :1*/ append(nEQLP7, hMWDrBfy...)
	if  /*line :1*/ cNaiWwqHDG.IsDir() {
		nEQLP7 =  /*line :1*/ append(nEQLP7, '/')
	}
	return  /*line :1*/ string(nEQLP7)
}
