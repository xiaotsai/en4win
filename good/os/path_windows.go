//line :1
package hPMrClpC

const (
	PathSeparator	= '\\'		
	PathListSeparator	= ';'		
)


func GClWKECc(ivfnLGc uint8) bool {

	return ivfnLGc == '\\' || ivfnLGc == '/'
}



func tVaoof6(jGBs03 string) string {

	if  /*line :1*/ len(jGBs03) == 2 && jGBs03[1] == ':' {
		jGBs03 = "."
	} else if  /*line :1*/ len(jGBs03) > 2 && jGBs03[1] == ':' {
		jGBs03 = jGBs03[2:]
	}
	lWfaal :=  /*line :1*/ len(jGBs03) - 1

	for ; lWfaal > 0 && (jGBs03[lWfaal] == '/' || jGBs03[lWfaal] == '\\'); lWfaal-- {
		jGBs03 = jGBs03[:lWfaal]
	}

	for lWfaal--; lWfaal >= 0; lWfaal-- {
		if jGBs03[lWfaal] == '/' || jGBs03[lWfaal] == '\\' {
			jGBs03 = jGBs03[lWfaal+1:]
			break
		}
	}
	return jGBs03
}

func hJq_1R(sJC41DS string) (nR2aPlreQFZA bool) {
	rB6hatF3c :=  /*line :1*/ ccY1eaDH(sJC41DS)
	if rB6hatF3c == "" {
		return false
	}
	sJC41DS = sJC41DS[ /*line :1*/ len(rB6hatF3c):]
	if sJC41DS == "" {
		return false
	}
	return  /*line :1*/ GClWKECc(sJC41DS[0])
}

func ccY1eaDH(sJC41DS string) (rB6hatF3c string) {
	if  /*line :1*/ len(sJC41DS) < 2 {
		return ""
	}

	ivfnLGc := sJC41DS[0]
	if sJC41DS[1] == ':' &&
		('0' <= ivfnLGc && ivfnLGc <= '9' || 'a' <= ivfnLGc && ivfnLGc <= 'z' ||
			'A' <= ivfnLGc && ivfnLGc <= 'Z') {
		return sJC41DS[:2]
	}

	if kC6f6XJP :=  /*line :1*/ len(sJC41DS); kC6f6XJP >= 5 &&  /*line :1*/ GClWKECc(sJC41DS[0]) &&  /*line :1*/ GClWKECc(sJC41DS[1]) &&
		! /*line :1*/ GClWKECc(sJC41DS[2]) && sJC41DS[2] != '.' {

		for zHDjCTHI := 3; zHDjCTHI < kC6f6XJP-1; zHDjCTHI++ {

			if  /*line :1*/ GClWKECc(sJC41DS[zHDjCTHI]) {
				zHDjCTHI++

				if ! /*line :1*/ GClWKECc(sJC41DS[zHDjCTHI]) {
					if sJC41DS[zHDjCTHI] == '.' {
						break
					}
					for ; zHDjCTHI < kC6f6XJP; zHDjCTHI++ {
						if  /*line :1*/ GClWKECc(sJC41DS[zHDjCTHI]) {
							break
						}
					}
					return sJC41DS[:zHDjCTHI]
				}
				break
			}
		}
	}
	return ""
}

func zS1K22(sJC41DS string) string {
	
	var tifurl []byte
	var qs8OBE int
	for lWfaal, nR2aPlreQFZA := range sJC41DS {
		if nR2aPlreQFZA == '/' {
			if tifurl == nil {
				tifurl =  /*line :1*/ make([]byte,  /*line :1*/ len(sJC41DS))
			}
			 /*line :1*/ copy(tifurl[qs8OBE:], sJC41DS[qs8OBE:lWfaal])
			tifurl[lWfaal] = '\\'
			qs8OBE = lWfaal + 1
		}
	}
	if tifurl == nil {
		return sJC41DS
	}

	 /*line :1*/ copy(tifurl[qs8OBE:], sJC41DS[qs8OBE:])
	return  /*line :1*/ string(tifurl)
}

func f4FBKwq7euK(sJC41DS string) string {
	qAbdfu_aW5 :=  /*line :1*/ ccY1eaDH(sJC41DS)
	lWfaal :=  /*line :1*/ len(sJC41DS) - 1
	for lWfaal >=  /*line :1*/ len(qAbdfu_aW5) && ! /*line :1*/ GClWKECc(sJC41DS[lWfaal]) {
		lWfaal--
	}
	nqzPkc8LPse1 := sJC41DS[ /*line :1*/ len(qAbdfu_aW5) : lWfaal+1]
	u8c7bctQ :=  /*line :1*/ len(nqzPkc8LPse1) - 1
	if u8c7bctQ > 0 &&  /*line :1*/ GClWKECc(nqzPkc8LPse1[u8c7bctQ]) {
		nqzPkc8LPse1 = nqzPkc8LPse1[:u8c7bctQ]
	}
	if nqzPkc8LPse1 == "" {
		nqzPkc8LPse1 = "."
	}
	return qAbdfu_aW5 + nqzPkc8LPse1
}



var vMpKx31aMBo bool









func xPdJPzd7sdeS(sJC41DS string) string {
	if vMpKx31aMBo {
		return sJC41DS
	}

	if  /*line :1*/ len(sJC41DS) < 248 {

		return sJC41DS
	}

	if  /*line :1*/ len(sJC41DS) >= 2 && sJC41DS[:2] == `\\` {

		return sJC41DS
	}
	if ! /*line :1*/ hJq_1R(sJC41DS) {

		return sJC41DS
	}

	const prefix = `\\?`

	tifurl :=  /*line :1*/ make([]byte,  /*line :1*/ len(prefix)+ /*line :1*/ len(sJC41DS)+ /*line :1*/ len(`\`))
	 /*line :1*/ copy(tifurl, prefix)
	zHDjCTHI :=  /*line :1*/ len(sJC41DS)
	dhvYus, pgkDcUH := 0,  /*line :1*/ len(prefix)
	for dhvYus < zHDjCTHI {
		switch {
		case  /*line :1*/ GClWKECc(sJC41DS[dhvYus]):

			dhvYus++
		case sJC41DS[dhvYus] == '.' && (dhvYus+1 == zHDjCTHI ||  /*line :1*/ GClWKECc(sJC41DS[dhvYus+1])):

			dhvYus++
		case dhvYus+1 < zHDjCTHI && sJC41DS[dhvYus] == '.' && sJC41DS[dhvYus+1] == '.' && (dhvYus+2 == zHDjCTHI ||  /*line :1*/ GClWKECc(sJC41DS[dhvYus+2])):

			return sJC41DS
		default:
			tifurl[pgkDcUH] = '\\'
			pgkDcUH++
			for ; dhvYus < zHDjCTHI && ! /*line :1*/ GClWKECc(sJC41DS[dhvYus]); dhvYus++ {
				tifurl[pgkDcUH] = sJC41DS[dhvYus]
				pgkDcUH++
			}
		}
	}

	if pgkDcUH ==  /*line :1*/ len(`\\?\c:`) {
		tifurl[pgkDcUH] = '\\'
		pgkDcUH++
	}
	return  /*line :1*/ string(tifurl[:pgkDcUH])
}



func twDUdA(baV7tO7i5 string) string {
	if  /*line :1*/ len(baV7tO7i5) ==  /*line :1*/ len(`\\?\c:`) {
		if  /*line :1*/ GClWKECc(baV7tO7i5[0]) &&  /*line :1*/ GClWKECc(baV7tO7i5[1]) && baV7tO7i5[2] == '?' &&  /*line :1*/ GClWKECc(baV7tO7i5[3]) && baV7tO7i5[5] == ':' {
			return baV7tO7i5 + `\`
		}
	}
	return baV7tO7i5
}
