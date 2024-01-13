//line :1
package qyOdGuID2rmZ

import (
	errors "iAHoxjmM"
	os "hPMrClpC"
	"runtime"
	sort "gzHaha55n"
	strings "fQXlzv"
	utf8 "CuAc0pPZwf"
)


var RWiGVHIyX =  /*line :1*/ errors.Su6g6hRoi9X("syntax error in pattern")

























func U_t6yhcON3(ynwVZ2pgCfmW, hurYsplzYEd string) (rI0O03EsJ bool, aC4eBbavAAtF error) {
Pattern:
	for  /*line :1*/ len(ynwVZ2pgCfmW) > 0 {
		var gt5San2Rx bool
		var b1fPOR1D string
		gt5San2Rx, b1fPOR1D, ynwVZ2pgCfmW =  /*line :1*/ _Iu5ctBQspJ(ynwVZ2pgCfmW)
		if gt5San2Rx && b1fPOR1D == "" {

			return ! /*line :1*/ strings.OHUqdLTpfyt(hurYsplzYEd,  /*line :1*/ string(Separator)), nil
		}

		lvbRLao1aWWH, gkdE2r, aC4eBbavAAtF :=  /*line :1*/ ketc417vYl5(b1fPOR1D, hurYsplzYEd)

		if gkdE2r && ( /*line :1*/ len(lvbRLao1aWWH) == 0 ||  /*line :1*/ len(ynwVZ2pgCfmW) > 0) {
			hurYsplzYEd = lvbRLao1aWWH
			continue
		}
		if aC4eBbavAAtF != nil {
			return false, aC4eBbavAAtF
		}
		if gt5San2Rx {

			for oesU8hV6 := 0; oesU8hV6 <  /*line :1*/ len(hurYsplzYEd) && hurYsplzYEd[oesU8hV6] != Separator; oesU8hV6++ {
				lvbRLao1aWWH, gkdE2r, aC4eBbavAAtF :=  /*line :1*/ ketc417vYl5(b1fPOR1D, hurYsplzYEd[oesU8hV6+1:])
				if gkdE2r {

					if  /*line :1*/ len(ynwVZ2pgCfmW) == 0 &&  /*line :1*/ len(lvbRLao1aWWH) > 0 {
						continue
					}
					hurYsplzYEd = lvbRLao1aWWH
					continue Pattern
				}
				if aC4eBbavAAtF != nil {
					return false, aC4eBbavAAtF
				}
			}
		}
		return false, nil
	}
	return  /*line :1*/ len(hurYsplzYEd) == 0, nil
}



func _Iu5ctBQspJ(ynwVZ2pgCfmW string) (gt5San2Rx bool, b1fPOR1D, m1lPbxJNN7G string) {
	for  /*line :1*/ len(ynwVZ2pgCfmW) > 0 && ynwVZ2pgCfmW[0] == '*' {
		ynwVZ2pgCfmW = ynwVZ2pgCfmW[1:]
		gt5San2Rx = true
	}
	xsac4zijpwf := false
	var oesU8hV6 int
Scan:
	for oesU8hV6 = 0; oesU8hV6 <  /*line :1*/ len(ynwVZ2pgCfmW); oesU8hV6++ {
		switch ynwVZ2pgCfmW[oesU8hV6] {
		case '\\':
			if runtime.GOOS != "windows" {

				if oesU8hV6+1 <  /*line :1*/ len(ynwVZ2pgCfmW) {
					oesU8hV6++
				}
			}
		case '[':
			xsac4zijpwf = true
		case ']':
			xsac4zijpwf = false
		case '*':
			if !xsac4zijpwf {
				break Scan
			}
		}
	}
	return gt5San2Rx, ynwVZ2pgCfmW[0:oesU8hV6], ynwVZ2pgCfmW[oesU8hV6:]
}




func ketc417vYl5(b1fPOR1D, tcYoc2qMuW string) (m1lPbxJNN7G string, gkdE2r bool, aC4eBbavAAtF error) {

	nvpqE18T3 := false
	for  /*line :1*/ len(b1fPOR1D) > 0 {
		if !nvpqE18T3 &&  /*line :1*/ len(tcYoc2qMuW) == 0 {
			nvpqE18T3 = true
		}
		switch b1fPOR1D[0] {
		case '[':
			
			var qLSJu6fd4 rune
			if !nvpqE18T3 {
				var oTJUmc1VT_o int
				qLSJu6fd4, oTJUmc1VT_o =  /*line :1*/ utf8.HVAV7aTqxzg(tcYoc2qMuW)
				tcYoc2qMuW = tcYoc2qMuW[oTJUmc1VT_o:]
			}
			b1fPOR1D = b1fPOR1D[1:]

			sBBcrgwo3J := false
			if  /*line :1*/ len(b1fPOR1D) > 0 && b1fPOR1D[0] == '^' {
				sBBcrgwo3J = true
				b1fPOR1D = b1fPOR1D[1:]
			}

			xkhMMuGlUR := false
			bc_1vtZ8 := 0
			for {
				if  /*line :1*/ len(b1fPOR1D) > 0 && b1fPOR1D[0] == ']' && bc_1vtZ8 > 0 {
					b1fPOR1D = b1fPOR1D[1:]
					break
				}
				var xC39pAYThZ, mrVQiIs rune
				if xC39pAYThZ, b1fPOR1D, aC4eBbavAAtF =  /*line :1*/ mvLuRGF1aBVO(b1fPOR1D); aC4eBbavAAtF != nil {
					return "", false, aC4eBbavAAtF
				}
				mrVQiIs = xC39pAYThZ
				if b1fPOR1D[0] == '-' {
					if mrVQiIs, b1fPOR1D, aC4eBbavAAtF =  /*line :1*/ mvLuRGF1aBVO(b1fPOR1D[1:]); aC4eBbavAAtF != nil {
						return "", false, aC4eBbavAAtF
					}
				}
				if xC39pAYThZ <= qLSJu6fd4 && qLSJu6fd4 <= mrVQiIs {
					xkhMMuGlUR = true
				}
				bc_1vtZ8++
			}
			if xkhMMuGlUR == sBBcrgwo3J {
				nvpqE18T3 = true
			}

		case '?':
			if !nvpqE18T3 {
				if tcYoc2qMuW[0] == Separator {
					nvpqE18T3 = true
				}
				_, oTJUmc1VT_o :=  /*line :1*/ utf8.HVAV7aTqxzg(tcYoc2qMuW)
				tcYoc2qMuW = tcYoc2qMuW[oTJUmc1VT_o:]
			}
			b1fPOR1D = b1fPOR1D[1:]

		case '\\':
			if runtime.GOOS != "windows" {
				b1fPOR1D = b1fPOR1D[1:]
				if  /*line :1*/ len(b1fPOR1D) == 0 {
					return "", false, RWiGVHIyX
				}
			}
			fallthrough

		default:
			if !nvpqE18T3 {
				if b1fPOR1D[0] != tcYoc2qMuW[0] {
					nvpqE18T3 = true
				}
				tcYoc2qMuW = tcYoc2qMuW[1:]
			}
			b1fPOR1D = b1fPOR1D[1:]
		}
	}
	if nvpqE18T3 {
		return "", false, nil
	}
	return tcYoc2qMuW, true, nil
}


func mvLuRGF1aBVO(b1fPOR1D string) (qLSJu6fd4 rune, zoZQc7Rt1q string, aC4eBbavAAtF error) {
	if  /*line :1*/ len(b1fPOR1D) == 0 || b1fPOR1D[0] == '-' || b1fPOR1D[0] == ']' {
		aC4eBbavAAtF = RWiGVHIyX
		return
	}
	if b1fPOR1D[0] == '\\' && runtime.GOOS != "windows" {
		b1fPOR1D = b1fPOR1D[1:]
		if  /*line :1*/ len(b1fPOR1D) == 0 {
			aC4eBbavAAtF = RWiGVHIyX
			return
		}
	}
	qLSJu6fd4, oTJUmc1VT_o :=  /*line :1*/ utf8.HVAV7aTqxzg(b1fPOR1D)
	if qLSJu6fd4 == utf8.RuneError && oTJUmc1VT_o == 1 {
		aC4eBbavAAtF = RWiGVHIyX
	}
	zoZQc7Rt1q = b1fPOR1D[oTJUmc1VT_o:]
	if  /*line :1*/ len(zoZQc7Rt1q) == 0 {
		aC4eBbavAAtF = RWiGVHIyX
	}
	return
}









func QZbVYHE7Fe(ynwVZ2pgCfmW string) (oB5zajFdJ4 []string, aC4eBbavAAtF error) {
	return  /*line :1*/ jsUJ31CT(ynwVZ2pgCfmW, 0)
}

func jsUJ31CT(ynwVZ2pgCfmW string, juhJKdbV5 int) (oB5zajFdJ4 []string, aC4eBbavAAtF error) {
	
	const pathSeparatorsLimit = 10000
	if juhJKdbV5 == pathSeparatorsLimit {
		return nil, RWiGVHIyX
	}

	if _, aC4eBbavAAtF :=  /*line :1*/ U_t6yhcON3(ynwVZ2pgCfmW, ""); aC4eBbavAAtF != nil {
		return nil, aC4eBbavAAtF
	}
	if ! /*line :1*/ mGmgKboal(ynwVZ2pgCfmW) {
		if _, aC4eBbavAAtF =  /*line :1*/ os.Z0JXsY(ynwVZ2pgCfmW); aC4eBbavAAtF != nil {
			return nil, nil
		}
		return []string{ynwVZ2pgCfmW}, nil
	}

	gzQrdT03qcU6, mRxgccbf :=  /*line :1*/ Dav_Lk(ynwVZ2pgCfmW)
	eiLIPBAiMel := 0
	if runtime.GOOS == "windows" {
		eiLIPBAiMel, gzQrdT03qcU6 =  /*line :1*/ xZyBOVug(gzQrdT03qcU6)
	} else {
		gzQrdT03qcU6 =  /*line :1*/ nycVoc(gzQrdT03qcU6)
	}

	if ! /*line :1*/ mGmgKboal(gzQrdT03qcU6[eiLIPBAiMel:]) {
		return  /*line :1*/ j56NdWOXl0h(gzQrdT03qcU6, mRxgccbf, nil)
	}

	if gzQrdT03qcU6 == ynwVZ2pgCfmW {
		return nil, RWiGVHIyX
	}

	var e_xuACsO []string
	e_xuACsO, aC4eBbavAAtF =  /*line :1*/ jsUJ31CT(gzQrdT03qcU6, juhJKdbV5+1)
	if aC4eBbavAAtF != nil {
		return
	}
	for _, qFtGgOeDz := range e_xuACsO {
		oB5zajFdJ4, aC4eBbavAAtF =  /*line :1*/ j56NdWOXl0h(qFtGgOeDz, mRxgccbf, oB5zajFdJ4)
		if aC4eBbavAAtF != nil {
			return
		}
	}
	return
}


func nycVoc(w9VHSyrN string) string {
	switch w9VHSyrN {
	case "":
		return "."
	case  /*line :1*/ string(Separator):

		return w9VHSyrN
	default:
		return w9VHSyrN[0 :  /*line :1*/ len(w9VHSyrN)-1]
	}
}


func xZyBOVug(w9VHSyrN string) (bTJLEmHMgw int, xuPyR3h string) {
	zOfNjauV3blf :=  /*line :1*/ sg1_tZpahJa(w9VHSyrN)
	switch {
	case w9VHSyrN == "":
		return 0, "."
	case zOfNjauV3blf+1 ==  /*line :1*/ len(w9VHSyrN) &&  /*line :1*/ os.GClWKECc(w9VHSyrN[ /*line :1*/ len(w9VHSyrN)-1]):

		return zOfNjauV3blf + 1, w9VHSyrN
	case zOfNjauV3blf ==  /*line :1*/ len(w9VHSyrN) &&  /*line :1*/ len(w9VHSyrN) == 2:
		return zOfNjauV3blf, w9VHSyrN + "."
	default:
		if zOfNjauV3blf >=  /*line :1*/ len(w9VHSyrN) {
			zOfNjauV3blf =  /*line :1*/ len(w9VHSyrN) - 1
		}
		return zOfNjauV3blf, w9VHSyrN[0 :  /*line :1*/ len(w9VHSyrN)-1]
	}
}





func j56NdWOXl0h(gzQrdT03qcU6, ynwVZ2pgCfmW string, oB5zajFdJ4 []string) (e_xuACsO []string, xsTcYHKO error) {
	e_xuACsO = oB5zajFdJ4
	tWHgrAGOK, aC4eBbavAAtF :=  /*line :1*/ os.ZYa3wuv(gzQrdT03qcU6)
	if aC4eBbavAAtF != nil {
		return
	}
	if ! /*line :1*/ tWHgrAGOK.IsDir() {
		return
	}
	qFtGgOeDz, aC4eBbavAAtF :=  /*line :1*/ os.Cvk6wxcjw(gzQrdT03qcU6)
	if aC4eBbavAAtF != nil {
		return
	}
	defer  /*line :1*/ qFtGgOeDz.Close()

	j7XUfK5nkoJ, _ :=  /*line :1*/ qFtGgOeDz.Readdirnames(-1)
	 /*line :1*/ sort.IsJh7v(j7XUfK5nkoJ)

	for _, oTJUmc1VT_o := range j7XUfK5nkoJ {
		rI0O03EsJ, aC4eBbavAAtF :=  /*line :1*/ U_t6yhcON3(ynwVZ2pgCfmW, oTJUmc1VT_o)
		if aC4eBbavAAtF != nil {
			return e_xuACsO, aC4eBbavAAtF
		}
		if rI0O03EsJ {
			e_xuACsO =  /*line :1*/ append(e_xuACsO,  /*line :1*/ SzdnHBv(gzQrdT03qcU6, oTJUmc1VT_o))
		}
	}
	return
}



func mGmgKboal(w9VHSyrN string) bool {
	y7sJer7sI := `*?[`
	if runtime.GOOS != "windows" {
		y7sJer7sI = `*?[\`
	}
	return  /*line :1*/ strings.EogZRPaCpV(w9VHSyrN, y7sJer7sI)
}
