//line :1
package qyOdGuID2rmZ

import (
	strings "fQXlzv"
	syscall "bUKeamF"
)




func iHysSOgM(w9VHSyrN string) string {
	bBilVO :=  /*line :1*/ SDvSg10pbg(w9VHSyrN)

	if  /*line :1*/ len(bBilVO) > 2 {
		return bBilVO
	}

	return  /*line :1*/ strings.Z7tZslPsj(bBilVO)
}


func rAiXwqTDeR(w9VHSyrN string) (string, error) {
	oH7gzT, aC4eBbavAAtF :=  /*line :1*/ syscall.GcOmFfsibES(w9VHSyrN)
	if aC4eBbavAAtF != nil {
		return "", aC4eBbavAAtF
	}

	var kKISJj5oe88 syscall.RsoZmJigCkGe

	eclQg3, aC4eBbavAAtF :=  /*line :1*/ syscall.YWRpP1l(oH7gzT, &kKISJj5oe88)
	if aC4eBbavAAtF != nil {
		return "", aC4eBbavAAtF
	}
	 /*line :1*/ syscall.JaVhl9Mkgfp(eclQg3)

	return  /*line :1*/ syscall.AODVXp8NM3sd(kKISJj5oe88.LMw5oQM3[:]), nil
}



func b0ZwCQI6Yd(w9VHSyrN string) bool {
	oesU8hV6 :=  /*line :1*/ strings.Owr8r4yL(w9VHSyrN, Separator)
	return w9VHSyrN[oesU8hV6+1:] == ".."
}












func wKdNp3(w9VHSyrN string, rAiXwqTDeR func(string) (string, error)) (string, error) {
	if w9VHSyrN == "" {
		return w9VHSyrN, nil
	}

	bBilVO :=  /*line :1*/ iHysSOgM(w9VHSyrN)
	w9VHSyrN = w9VHSyrN[ /*line :1*/ len(bBilVO):]

	if w9VHSyrN == "" || w9VHSyrN == "." || w9VHSyrN == `\` {
		return bBilVO + w9VHSyrN, nil
	}

	var bswnXb string

	for {
		if  /*line :1*/ b0ZwCQI6Yd(w9VHSyrN) {
			bswnXb = w9VHSyrN + `\` + bswnXb

			break
		}

		hurYsplzYEd, aC4eBbavAAtF :=  /*line :1*/ rAiXwqTDeR(bBilVO + w9VHSyrN)
		if aC4eBbavAAtF != nil {
			return "", aC4eBbavAAtF
		}

		bswnXb = hurYsplzYEd + `\` + bswnXb

		oesU8hV6 :=  /*line :1*/ strings.Owr8r4yL(w9VHSyrN, Separator)
		if oesU8hV6 == -1 {
			break
		}
		if oesU8hV6 == 0 {
			bswnXb = `\` + bswnXb

			break
		}

		w9VHSyrN = w9VHSyrN[:oesU8hV6]
	}

	bswnXb = bswnXb[: /*line :1*/ len(bswnXb)-1]

	return bBilVO + bswnXb, nil
}

func jgOEZ8zayc4(w9VHSyrN string) (string, error) {
	nOaOV7O, aC4eBbavAAtF :=  /*line :1*/ wDjPYoGaI72T(w9VHSyrN)
	if aC4eBbavAAtF != nil {
		return "", aC4eBbavAAtF
	}
	nOaOV7O, aC4eBbavAAtF =  /*line :1*/ wKdNp3(nOaOV7O, rAiXwqTDeR)
	if aC4eBbavAAtF != nil {
		return "", aC4eBbavAAtF
	}
	return nOaOV7O, nil
}
