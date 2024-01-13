//line :1
package hPMrClpC


func bkIhacN(m4R4sea int) string {
	if m4R4sea < 0 {
		return "-" +  /*line :1*/ tsxkt9sR( /*line :1*/ uint(-m4R4sea))
	}
	return  /*line :1*/ tsxkt9sR( /*line :1*/ uint(m4R4sea))
}

const hex = "0123456789abcdef"


func tsxkt9sR(m4R4sea uint) string {
	if m4R4sea == 0 {
		return "0x0"
	}
	var dDkChyAkT [20]byte	
	lWfaal :=  /*line :1*/ len(dDkChyAkT) - 1
	for m4R4sea >= 16 {
		r72jyAohwnla := m4R4sea / 16
		dDkChyAkT[lWfaal] = hex[m4R4sea%16]
		lWfaal--
		m4R4sea = r72jyAohwnla
	}

	dDkChyAkT[lWfaal] = hex[m4R4sea%16]
	lWfaal--
	dDkChyAkT[lWfaal] = 'x'
	lWfaal--
	dDkChyAkT[lWfaal] = '0'
	return  /*line :1*/ string(dDkChyAkT[lWfaal:])
}
