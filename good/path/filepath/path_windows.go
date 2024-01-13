//line :1
package qyOdGuID2rmZ

import (
	strings "fQXlzv"
	syscall "bUKeamF"
)

func to2v7JEF(y60OphIRDC uint8) bool {
	return y60OphIRDC == '\\' || y60OphIRDC == '/'
}

func n6hcihVNkM6(y60OphIRDC byte) byte {
	if 'a' <= y60OphIRDC && y60OphIRDC <= 'z' {
		return y60OphIRDC - ('a' - 'A')
	}
	return y60OphIRDC
}






func fYjwGRrA_(hurYsplzYEd string) bool {
	if 3 <=  /*line :1*/ len(hurYsplzYEd) &&  /*line :1*/ len(hurYsplzYEd) <= 4 {
		switch  /*line :1*/ string([]byte{ /*line :1*/ n6hcihVNkM6(hurYsplzYEd[0]),  /*line :1*/ n6hcihVNkM6(hurYsplzYEd[1]),  /*line :1*/ n6hcihVNkM6(hurYsplzYEd[2])}) {
		case "CON", "PRN", "AUX", "NUL":
			return  /*line :1*/ len(hurYsplzYEd) == 3
		case "COM", "LPT":
			return  /*line :1*/ len(hurYsplzYEd) == 4 && '1' <= hurYsplzYEd[3] && hurYsplzYEd[3] <= '9'
		}
	}

	if  /*line :1*/ len(hurYsplzYEd) == 6 && hurYsplzYEd[5] == '$' &&  /*line :1*/ strings.UPM6mljQncm9(hurYsplzYEd, "CONIN$") {
		return true
	}
	if  /*line :1*/ len(hurYsplzYEd) == 7 && hurYsplzYEd[6] == '$' &&  /*line :1*/ strings.UPM6mljQncm9(hurYsplzYEd, "CONOUT$") {
		return true
	}
	return false
}

func a1cYZ5(w9VHSyrN string) bool {
	if w9VHSyrN == "" {
		return false
	}
	if  /*line :1*/ to2v7JEF(w9VHSyrN[0]) {

		return false
	}
	if  /*line :1*/ strings.Po8D3DCRZq(w9VHSyrN, ':') >= 0 {

		return false
	}
	aGpz_h := false
	for oH7gzT := w9VHSyrN; oH7gzT != ""; {
		var lYtfrVIXIqpE string
		lYtfrVIXIqpE, oH7gzT, _ =  /*line :1*/ uhPpKRjFvUZ3(oH7gzT)
		if lYtfrVIXIqpE == "." || lYtfrVIXIqpE == ".." {
			aGpz_h = true
		}

		hqaPpA8B9VW, _, cudKCQBSX :=  /*line :1*/ strings.E093HWXxK(lYtfrVIXIqpE, ".")
		if  /*line :1*/ fYjwGRrA_(hqaPpA8B9VW) {
			if !cudKCQBSX {
				return false
			}

			if oH7gzT, _ :=  /*line :1*/ syscall.LmpGp3wH_T(lYtfrVIXIqpE);  /*line :1*/ len(oH7gzT) >= 4 && oH7gzT[:4] == `\\.\` {
				return false
			}
		}
	}
	if aGpz_h {
		w9VHSyrN =  /*line :1*/ ZBiAXr(w9VHSyrN)
	}
	if w9VHSyrN == ".." ||  /*line :1*/ strings.E8ZCzfgw(w9VHSyrN, `..\`) {
		return false
	}
	return true
}


func Lr02fA2mZb(w9VHSyrN string) (bEuBmdQbx bool) {
	nB_CPr6X :=  /*line :1*/ sg1_tZpahJa(w9VHSyrN)
	if nB_CPr6X == 0 {
		return false
	}

	if  /*line :1*/ to2v7JEF(w9VHSyrN[0]) &&  /*line :1*/ to2v7JEF(w9VHSyrN[1]) {
		return true
	}
	w9VHSyrN = w9VHSyrN[nB_CPr6X:]
	if w9VHSyrN == "" {
		return false
	}
	return  /*line :1*/ to2v7JEF(w9VHSyrN[0])
}





func sg1_tZpahJa(w9VHSyrN string) int {
	if  /*line :1*/ len(w9VHSyrN) < 2 {
		return 0
	}

	y60OphIRDC := w9VHSyrN[0]
	if w9VHSyrN[1] == ':' && ('a' <= y60OphIRDC && y60OphIRDC <= 'z' || 'A' <= y60OphIRDC && y60OphIRDC <= 'Z') {
		return 2
	}

	if ! /*line :1*/ to2v7JEF(w9VHSyrN[0]) || ! /*line :1*/ to2v7JEF(w9VHSyrN[1]) {
		return 0
	}
	m1lPbxJNN7G := w9VHSyrN[2:]
	m95tb7vI, m1lPbxJNN7G, _ :=  /*line :1*/ uhPpKRjFvUZ3(m1lPbxJNN7G)
	lzgZ5OuAdu, m1lPbxJNN7G, gkdE2r :=  /*line :1*/ uhPpKRjFvUZ3(m1lPbxJNN7G)
	if !gkdE2r {
		return  /*line :1*/ len(w9VHSyrN)
	}
	if m95tb7vI != "." && m95tb7vI != "?" {

		return  /*line :1*/ len(w9VHSyrN) -  /*line :1*/ len(m1lPbxJNN7G) - 1
	}

	if  /*line :1*/ len(lzgZ5OuAdu) == 3 &&  /*line :1*/ n6hcihVNkM6(lzgZ5OuAdu[0]) == 'U' &&  /*line :1*/ n6hcihVNkM6(lzgZ5OuAdu[1]) == 'N' &&  /*line :1*/ n6hcihVNkM6(lzgZ5OuAdu[2]) == 'C' {

		_, m1lPbxJNN7G, _ =  /*line :1*/ uhPpKRjFvUZ3(m1lPbxJNN7G)
		_, m1lPbxJNN7G, gkdE2r =  /*line :1*/ uhPpKRjFvUZ3(m1lPbxJNN7G)
		if !gkdE2r {
			return  /*line :1*/ len(w9VHSyrN)
		}
	}
	return  /*line :1*/ len(w9VHSyrN) -  /*line :1*/ len(m1lPbxJNN7G) - 1
}


func uhPpKRjFvUZ3(w9VHSyrN string) (oFCtDZn, boFl6aatt string, eiOdbQwIK6 bool) {
	for oesU8hV6 := range w9VHSyrN {
		if  /*line :1*/ to2v7JEF(w9VHSyrN[oesU8hV6]) {
			return w9VHSyrN[:oesU8hV6], w9VHSyrN[oesU8hV6+1:], true
		}
	}
	return w9VHSyrN, "", false
}





func MwRH8fz9s(oH7gzT, _KgGzh string) bool {
	if  /*line :1*/ strings.E8ZCzfgw(oH7gzT, _KgGzh) {
		return true
	}
	return  /*line :1*/ strings.E8ZCzfgw( /*line :1*/ strings.WKrzOP(oH7gzT),  /*line :1*/ strings.WKrzOP(_KgGzh))
}

func fkb7AacBQW(w9VHSyrN string) []string {

	if w9VHSyrN == "" {
		return []string{}
	}

	hszR13aQioa := []string{}
	whLL2V3vq := 0
	ei7kNGUj6 := false
	for oesU8hV6 := 0; oesU8hV6 <  /*line :1*/ len(w9VHSyrN); oesU8hV6++ {
		switch y60OphIRDC := w9VHSyrN[oesU8hV6]; {
		case y60OphIRDC == '"':
			ei7kNGUj6 = !ei7kNGUj6
		case y60OphIRDC == ListSeparator && !ei7kNGUj6:
			hszR13aQioa =  /*line :1*/ append(hszR13aQioa, w9VHSyrN[whLL2V3vq:oesU8hV6])
			whLL2V3vq = oesU8hV6 + 1
		}
	}
	hszR13aQioa =  /*line :1*/ append(hszR13aQioa, w9VHSyrN[whLL2V3vq:])

	for oesU8hV6, tcYoc2qMuW := range hszR13aQioa {
		hszR13aQioa[oesU8hV6] =  /*line :1*/ strings.Tq1MwZX3Za(tcYoc2qMuW, `"`, ``)
	}

	return hszR13aQioa
}

func u5ByFe5zj(w9VHSyrN string) (string, error) {
	if w9VHSyrN == "" {

		w9VHSyrN = "."
	}
	iGI0aAZX1An, aC4eBbavAAtF :=  /*line :1*/ syscall.LmpGp3wH_T(w9VHSyrN)
	if aC4eBbavAAtF != nil {
		return "", aC4eBbavAAtF
	}
	return  /*line :1*/ ZBiAXr(iGI0aAZX1An), nil
}

func l6DOabdEl(jDqheELP2 []string) string {
	var bEuBmdQbx strings.CKrNlN3otWSF
	var ooDpKmHaf1I byte
	for _, xsTcYHKO := range jDqheELP2 {
		switch {
		case  /*line :1*/ bEuBmdQbx.Len() == 0:

		case  /*line :1*/ to2v7JEF(ooDpKmHaf1I):

			for  /*line :1*/ len(xsTcYHKO) > 0 &&  /*line :1*/ to2v7JEF(xsTcYHKO[0]) {
				xsTcYHKO = xsTcYHKO[1:]
			}
		case ooDpKmHaf1I == ':':

		default:

			 /*line :1*/ bEuBmdQbx.WriteByte('\\')
			ooDpKmHaf1I = '\\'
		}
		if  /*line :1*/ len(xsTcYHKO) > 0 {
			 /*line :1*/ bEuBmdQbx.WriteString(xsTcYHKO)
			ooDpKmHaf1I = xsTcYHKO[ /*line :1*/ len(xsTcYHKO)-1]
		}
	}
	if  /*line :1*/ bEuBmdQbx.Len() == 0 {
		return ""
	}
	return  /*line :1*/ ZBiAXr( /*line :1*/ bEuBmdQbx.String())
}


func sk1ljMdk8ZXF(jDqheELP2 []string) string {
	if  /*line :1*/ len(jDqheELP2[0]) == 2 && jDqheELP2[0][1] == ':' {

		oesU8hV6 := 1
		for ; oesU8hV6 <  /*line :1*/ len(jDqheELP2); oesU8hV6++ {
			if jDqheELP2[oesU8hV6] != "" {
				break
			}
		}
		return  /*line :1*/ ZBiAXr(jDqheELP2[0] +  /*line :1*/ strings.ZTisTA(jDqheELP2[oesU8hV6:],  /*line :1*/ string(Separator)))
	}

	oH7gzT :=  /*line :1*/ ZBiAXr( /*line :1*/ strings.ZTisTA(jDqheELP2,  /*line :1*/ string(Separator)))
	if ! /*line :1*/ vK9GyjBQ(oH7gzT) {
		return oH7gzT
	}

	wy80gnB :=  /*line :1*/ ZBiAXr(jDqheELP2[0])
	if  /*line :1*/ vK9GyjBQ(wy80gnB) {
		return oH7gzT
	}

	gehNcQA2tqZ :=  /*line :1*/ ZBiAXr( /*line :1*/ strings.ZTisTA(jDqheELP2[1:],  /*line :1*/ string(Separator)))
	if wy80gnB[ /*line :1*/ len(wy80gnB)-1] == Separator {
		return wy80gnB + gehNcQA2tqZ
	}
	return wy80gnB +  /*line :1*/ string(Separator) + gehNcQA2tqZ
}


func vK9GyjBQ(w9VHSyrN string) bool {
	return  /*line :1*/ len(w9VHSyrN) > 1 &&  /*line :1*/ to2v7JEF(w9VHSyrN[0]) &&  /*line :1*/ to2v7JEF(w9VHSyrN[1])
}

func jzREQp(bDaTi4HX, bEuBmdQbx string) bool {
	return  /*line :1*/ strings.UPM6mljQncm9(bDaTi4HX, bEuBmdQbx)
}
