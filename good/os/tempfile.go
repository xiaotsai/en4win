//line :1
package hPMrClpC

import (
	errors "iAHoxjmM"
	itoa "JkjxVS"
)





func ax5_HC() uint32

func sFcwr6hr3IS() string {
	return  /*line :1*/ itoa.TDADPeM7Kft9( /*line :1*/ uint( /*line :1*/ ax5_HC()))
}









func AiYDRF8(nqzPkc8LPse1, jMT5nUOyX3 string) (*BvGocYcXx, error) {
	if nqzPkc8LPse1 == "" {
		nqzPkc8LPse1 =  /*line :1*/ Xx5OajUY()
	}

	psMMKytBh, td0TbL, ugqb2IW :=  /*line :1*/ tiDZ5VkZB(jMT5nUOyX3)
	if ugqb2IW != nil {
		return nil, &StNL1lveD40f{Aeg62j1VQt: "createtemp", OmDEtY: jMT5nUOyX3, OIEifQ0XF: ugqb2IW}
	}
	psMMKytBh =  /*line :1*/ qRnLCzEF(nqzPkc8LPse1, psMMKytBh)

	gtGl7aa := 0
	for {
		jGBs03 := psMMKytBh +  /*line :1*/ sFcwr6hr3IS() + td0TbL
		qzeVi328uMg, ugqb2IW :=  /*line :1*/ FxDeI2Pd3bSR(jGBs03, O_RDWR|O_CREATE|O_EXCL, 0600)
		if  /*line :1*/ Pi48aFnhbwU(ugqb2IW) {
			if gtGl7aa++; gtGl7aa < 10000 {
				continue
			}
			return nil, &StNL1lveD40f{Aeg62j1VQt: "createtemp", OmDEtY: psMMKytBh + "*" + td0TbL, OIEifQ0XF: W5S6mrXOlYD0}
		}
		return qzeVi328uMg, ugqb2IW
	}
}

var v7M_O546 =  /*line :1*/ errors.Su6g6hRoi9X("pattern contains path separator")



func tiDZ5VkZB(jMT5nUOyX3 string) (psMMKytBh, td0TbL string, ugqb2IW error) {
	for lWfaal := 0; lWfaal <  /*line :1*/ len(jMT5nUOyX3); lWfaal++ {
		if  /*line :1*/ GClWKECc(jMT5nUOyX3[lWfaal]) {
			return "", "", v7M_O546
		}
	}
	if aehRjcDV :=  /*line :1*/ zObiWXAW2(jMT5nUOyX3, '*'); aehRjcDV != -1 {
		psMMKytBh, td0TbL = jMT5nUOyX3[:aehRjcDV], jMT5nUOyX3[aehRjcDV+1:]
	} else {
		psMMKytBh = jMT5nUOyX3
	}
	return psMMKytBh, td0TbL, nil
}








func DrD5z60v8aIu(nqzPkc8LPse1, jMT5nUOyX3 string) (string, error) {
	if nqzPkc8LPse1 == "" {
		nqzPkc8LPse1 =  /*line :1*/ Xx5OajUY()
	}

	psMMKytBh, td0TbL, ugqb2IW :=  /*line :1*/ tiDZ5VkZB(jMT5nUOyX3)
	if ugqb2IW != nil {
		return "", &StNL1lveD40f{Aeg62j1VQt: "mkdirtemp", OmDEtY: jMT5nUOyX3, OIEifQ0XF: ugqb2IW}
	}
	psMMKytBh =  /*line :1*/ qRnLCzEF(nqzPkc8LPse1, psMMKytBh)

	gtGl7aa := 0
	for {
		jGBs03 := psMMKytBh +  /*line :1*/ sFcwr6hr3IS() + td0TbL
		ugqb2IW :=  /*line :1*/ OEmAGB_T1CQq(jGBs03, 0700)
		if ugqb2IW == nil {
			return jGBs03, nil
		}
		if  /*line :1*/ Pi48aFnhbwU(ugqb2IW) {
			if gtGl7aa++; gtGl7aa < 10000 {
				continue
			}
			return "", &StNL1lveD40f{Aeg62j1VQt: "mkdirtemp", OmDEtY: nqzPkc8LPse1 +  /*line :1*/ string(PathSeparator) + psMMKytBh + "*" + td0TbL, OIEifQ0XF: W5S6mrXOlYD0}
		}
		if  /*line :1*/ Ac3Db0LaN9ge(ugqb2IW) {
			if _, ugqb2IW :=  /*line :1*/ ZYa3wuv(nqzPkc8LPse1);  /*line :1*/ Ac3Db0LaN9ge(ugqb2IW) {
				return "", ugqb2IW
			}
		}
		return "", ugqb2IW
	}
}

func qRnLCzEF(nqzPkc8LPse1, jGBs03 string) string {
	if  /*line :1*/ len(nqzPkc8LPse1) > 0 &&  /*line :1*/ GClWKECc(nqzPkc8LPse1[ /*line :1*/ len(nqzPkc8LPse1)-1]) {
		return nqzPkc8LPse1 + jGBs03
	}
	return nqzPkc8LPse1 +  /*line :1*/ string(PathSeparator) + jGBs03
}


func zObiWXAW2(gkORch7na string, kluu_LZ_v byte) int {
	for lWfaal :=  /*line :1*/ len(gkORch7na) - 1; lWfaal >= 0; lWfaal-- {
		if gkORch7na[lWfaal] == kluu_LZ_v {
			return lWfaal
		}
	}
	return -1
}
