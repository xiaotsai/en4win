//line :1







package qyOdGuID2rmZ

import (
	errors "iAHoxjmM"
	fs "y1BamVjnXsaa"
	os "hPMrClpC"
	"runtime"
	sort "gzHaha55n"
	strings "fQXlzv"
)





type a4jCH9w struct {
	duSbMy9Z47	string
	ho5xih41	[]byte
	sOFKVk	int
	xWRR4e	string
	de5pp8Hk1	int
}

func (bEuBmdQbx *a4jCH9w) neFaR_UpV5(oesU8hV6 int) byte {
	if bEuBmdQbx.ho5xih41 != nil {
		return bEuBmdQbx.ho5xih41[oesU8hV6]
	}
	return bEuBmdQbx.duSbMy9Z47[oesU8hV6]
}

func (bEuBmdQbx *a4jCH9w) l87O597C0AU(y60OphIRDC byte) {
	if bEuBmdQbx.ho5xih41 == nil {
		if bEuBmdQbx.sOFKVk <  /*line :1*/ len(bEuBmdQbx.duSbMy9Z47) && bEuBmdQbx.duSbMy9Z47[bEuBmdQbx.sOFKVk] == y60OphIRDC {
			bEuBmdQbx.sOFKVk++
			return
		}
		bEuBmdQbx.ho5xih41 =  /*line :1*/ make([]byte,  /*line :1*/ len(bEuBmdQbx.duSbMy9Z47))
		 /*line :1*/ copy(bEuBmdQbx.ho5xih41, bEuBmdQbx.duSbMy9Z47[:bEuBmdQbx.sOFKVk])
	}
	bEuBmdQbx.ho5xih41[bEuBmdQbx.sOFKVk] = y60OphIRDC
	bEuBmdQbx.sOFKVk++
}

func (bEuBmdQbx *a4jCH9w) kpaiSl3uSbS(_KgGzh ...byte) {
	bEuBmdQbx.ho5xih41 =  /*line :1*/ append(_KgGzh, bEuBmdQbx.ho5xih41...)
	bEuBmdQbx.sOFKVk +=  /*line :1*/ len(_KgGzh)
}

func (bEuBmdQbx *a4jCH9w) kUEyDnaCP93() string {
	if bEuBmdQbx.ho5xih41 == nil {
		return bEuBmdQbx.xWRR4e[:bEuBmdQbx.de5pp8Hk1+bEuBmdQbx.sOFKVk]
	}
	return bEuBmdQbx.xWRR4e[:bEuBmdQbx.de5pp8Hk1] +  /*line :1*/ string(bEuBmdQbx.ho5xih41[:bEuBmdQbx.sOFKVk])
}

const (
	Separator	= os.PathSeparator
	ListSeparator	= os.PathListSeparator
)




























func ZBiAXr(w9VHSyrN string) string {
	cJ2N7RCqy7S := w9VHSyrN
	eDjfKln_q :=  /*line :1*/ sg1_tZpahJa(w9VHSyrN)
	w9VHSyrN = w9VHSyrN[eDjfKln_q:]
	if w9VHSyrN == "" {
		if eDjfKln_q > 1 &&  /*line :1*/ os.GClWKECc(cJ2N7RCqy7S[0]) &&  /*line :1*/ os.GClWKECc(cJ2N7RCqy7S[1]) {

			return  /*line :1*/ EBTjoI0Ix(cJ2N7RCqy7S)
		}
		return cJ2N7RCqy7S + "."
	}
	iWnka6qq6y :=  /*line :1*/ os.GClWKECc(w9VHSyrN[0])

	oTJUmc1VT_o :=  /*line :1*/ len(w9VHSyrN)
	gsBHbq_9I := a4jCH9w{duSbMy9Z47: w9VHSyrN, xWRR4e: cJ2N7RCqy7S, de5pp8Hk1: eDjfKln_q}
	qLSJu6fd4, wAwmQEb0SqFa := 0, 0
	if iWnka6qq6y {
		 /*line :1*/ gsBHbq_9I.l87O597C0AU(Separator)
		qLSJu6fd4, wAwmQEb0SqFa = 1, 1
	}

	for qLSJu6fd4 < oTJUmc1VT_o {
		switch {
		case  /*line :1*/ os.GClWKECc(w9VHSyrN[qLSJu6fd4]):

			qLSJu6fd4++
		case w9VHSyrN[qLSJu6fd4] == '.' && (qLSJu6fd4+1 == oTJUmc1VT_o ||  /*line :1*/ os.GClWKECc(w9VHSyrN[qLSJu6fd4+1])):

			qLSJu6fd4++
		case w9VHSyrN[qLSJu6fd4] == '.' && w9VHSyrN[qLSJu6fd4+1] == '.' && (qLSJu6fd4+2 == oTJUmc1VT_o ||  /*line :1*/ os.GClWKECc(w9VHSyrN[qLSJu6fd4+2])):

			qLSJu6fd4 += 2
			switch {
			case gsBHbq_9I.sOFKVk > wAwmQEb0SqFa:

				gsBHbq_9I.sOFKVk--
				for gsBHbq_9I.sOFKVk > wAwmQEb0SqFa && ! /*line :1*/ os.GClWKECc( /*line :1*/ gsBHbq_9I.neFaR_UpV5(gsBHbq_9I.sOFKVk)) {
					gsBHbq_9I.sOFKVk--
				}
			case !iWnka6qq6y:

				if gsBHbq_9I.sOFKVk > 0 {
					 /*line :1*/ gsBHbq_9I.l87O597C0AU(Separator)
				}
				 /*line :1*/ gsBHbq_9I.l87O597C0AU('.')
				 /*line :1*/ gsBHbq_9I.l87O597C0AU('.')
				wAwmQEb0SqFa = gsBHbq_9I.sOFKVk
			}
		default:

			if iWnka6qq6y && gsBHbq_9I.sOFKVk != 1 || !iWnka6qq6y && gsBHbq_9I.sOFKVk != 0 {
				 /*line :1*/ gsBHbq_9I.l87O597C0AU(Separator)
			}

			for ; qLSJu6fd4 < oTJUmc1VT_o && ! /*line :1*/ os.GClWKECc(w9VHSyrN[qLSJu6fd4]); qLSJu6fd4++ {
				 /*line :1*/ gsBHbq_9I.l87O597C0AU(w9VHSyrN[qLSJu6fd4])
			}
		}
	}

	if gsBHbq_9I.sOFKVk == 0 {
		 /*line :1*/ gsBHbq_9I.l87O597C0AU('.')
	}

	if runtime.GOOS == "windows" && gsBHbq_9I.de5pp8Hk1 == 0 && gsBHbq_9I.ho5xih41 != nil {

		for _, y60OphIRDC := range gsBHbq_9I.ho5xih41 {
			if  /*line :1*/ os.GClWKECc(y60OphIRDC) {
				break
			}
			if y60OphIRDC == ':' {
				 /*line :1*/ gsBHbq_9I.kpaiSl3uSbS('.', Separator)
				break
			}
		}
	}

	return  /*line :1*/ EBTjoI0Ix( /*line :1*/ gsBHbq_9I.kUEyDnaCP93())
}















func WLKPPB(w9VHSyrN string) bool {
	return  /*line :1*/ a1cYZ5(w9VHSyrN)
}

func fawBGa3b(w9VHSyrN string) bool {
	if  /*line :1*/ Lr02fA2mZb(w9VHSyrN) || w9VHSyrN == "" {
		return false
	}
	aGpz_h := false
	for oH7gzT := w9VHSyrN; oH7gzT != ""; {
		var lYtfrVIXIqpE string
		lYtfrVIXIqpE, oH7gzT, _ =  /*line :1*/ strings.E093HWXxK(oH7gzT, "/")
		if lYtfrVIXIqpE == "." || lYtfrVIXIqpE == ".." {
			aGpz_h = true
			break
		}
	}
	if aGpz_h {
		w9VHSyrN =  /*line :1*/ ZBiAXr(w9VHSyrN)
	}
	if w9VHSyrN == ".." ||  /*line :1*/ strings.E8ZCzfgw(w9VHSyrN, "../") {
		return false
	}
	return true
}




func Tm0ZJM(w9VHSyrN string) string {
	if Separator == '/' {
		return w9VHSyrN
	}
	return  /*line :1*/ strings.Tq1MwZX3Za(w9VHSyrN,  /*line :1*/ string(Separator), "/")
}




func EBTjoI0Ix(w9VHSyrN string) string {
	if Separator == '/' {
		return w9VHSyrN
	}
	return  /*line :1*/ strings.Tq1MwZX3Za(w9VHSyrN, "/",  /*line :1*/ string(Separator))
}





func RjRbNWt27(w9VHSyrN string) []string {
	return  /*line :1*/ fkb7AacBQW(w9VHSyrN)
}






func Dav_Lk(w9VHSyrN string) (gzQrdT03qcU6, mRxgccbf string) {
	d2oIeyTClzU :=  /*line :1*/ SDvSg10pbg(w9VHSyrN)
	oesU8hV6 :=  /*line :1*/ len(w9VHSyrN) - 1
	for oesU8hV6 >=  /*line :1*/ len(d2oIeyTClzU) && ! /*line :1*/ os.GClWKECc(w9VHSyrN[oesU8hV6]) {
		oesU8hV6--
	}
	return w9VHSyrN[:oesU8hV6+1], w9VHSyrN[oesU8hV6+1:]
}








func SzdnHBv(jDqheELP2 ...string) string {
	return  /*line :1*/ l6DOabdEl(jDqheELP2)
}





func A8rtqMdKVrN(w9VHSyrN string) string {
	for oesU8hV6 :=  /*line :1*/ len(w9VHSyrN) - 1; oesU8hV6 >= 0 && ! /*line :1*/ os.GClWKECc(w9VHSyrN[oesU8hV6]); oesU8hV6-- {
		if w9VHSyrN[oesU8hV6] == '.' {
			return w9VHSyrN[oesU8hV6:]
		}
	}
	return ""
}






func PkZ_3ObeSQ(w9VHSyrN string) (string, error) {
	return  /*line :1*/ jgOEZ8zayc4(w9VHSyrN)
}






func DWa4aCam(w9VHSyrN string) (string, error) {
	return  /*line :1*/ u5ByFe5zj(w9VHSyrN)
}

func apKZaPo3OSlh(w9VHSyrN string) (string, error) {
	if  /*line :1*/ Lr02fA2mZb(w9VHSyrN) {
		return  /*line :1*/ ZBiAXr(w9VHSyrN), nil
	}
	m9re0H, aC4eBbavAAtF :=  /*line :1*/ os.OWhXb5XTeJpD()
	if aC4eBbavAAtF != nil {
		return "", aC4eBbavAAtF
	}
	return  /*line :1*/ SzdnHBv(m9re0H, w9VHSyrN), nil
}









func Q8Vk6QcoI(aGGuRhZsKi, z_7xS9aip string) (string, error) {
	cnOqgyxbbnHM :=  /*line :1*/ SDvSg10pbg(aGGuRhZsKi)
	w9iKWIgJ9 :=  /*line :1*/ SDvSg10pbg(z_7xS9aip)
	hqaPpA8B9VW :=  /*line :1*/ ZBiAXr(aGGuRhZsKi)
	fpm9Hez :=  /*line :1*/ ZBiAXr(z_7xS9aip)
	if  /*line :1*/ jzREQp(fpm9Hez, hqaPpA8B9VW) {
		return ".", nil
	}
	hqaPpA8B9VW = hqaPpA8B9VW[ /*line :1*/ len(cnOqgyxbbnHM):]
	fpm9Hez = fpm9Hez[ /*line :1*/ len(w9iKWIgJ9):]
	if hqaPpA8B9VW == "." {
		hqaPpA8B9VW = ""
	} else if hqaPpA8B9VW == "" &&  /*line :1*/ sg1_tZpahJa(cnOqgyxbbnHM) > 2 {

		hqaPpA8B9VW =  /*line :1*/ string(Separator)
	}

	gTaAD3 :=  /*line :1*/ len(hqaPpA8B9VW) > 0 && hqaPpA8B9VW[0] == Separator
	eVboUWMVxLnS :=  /*line :1*/ len(fpm9Hez) > 0 && fpm9Hez[0] == Separator
	if gTaAD3 != eVboUWMVxLnS || ! /*line :1*/ jzREQp(cnOqgyxbbnHM, w9iKWIgJ9) {
		return "",  /*line :1*/ errors.Su6g6hRoi9X("Rel: can't make " + z_7xS9aip + " relative to " + aGGuRhZsKi)
	}

	tbfMp6zg :=  /*line :1*/ len(hqaPpA8B9VW)
	yAVPcVsd :=  /*line :1*/ len(fpm9Hez)
	var u3mi7PQzhQzX, ya343u, eI480IIHU, lC40ZbCN int
	for {
		for ya343u < tbfMp6zg && hqaPpA8B9VW[ya343u] != Separator {
			ya343u++
		}
		for lC40ZbCN < yAVPcVsd && fpm9Hez[lC40ZbCN] != Separator {
			lC40ZbCN++
		}
		if ! /*line :1*/ jzREQp(fpm9Hez[eI480IIHU:lC40ZbCN], hqaPpA8B9VW[u3mi7PQzhQzX:ya343u]) {
			break
		}
		if ya343u < tbfMp6zg {
			ya343u++
		}
		if lC40ZbCN < yAVPcVsd {
			lC40ZbCN++
		}
		u3mi7PQzhQzX = ya343u
		eI480IIHU = lC40ZbCN
	}
	if hqaPpA8B9VW[u3mi7PQzhQzX:ya343u] == ".." {
		return "",  /*line :1*/ errors.Su6g6hRoi9X("Rel: can't make " + z_7xS9aip + " relative to " + aGGuRhZsKi)
	}
	if u3mi7PQzhQzX != tbfMp6zg {

		rZEhR9f60tA :=  /*line :1*/ strings.L9XssxF(hqaPpA8B9VW[u3mi7PQzhQzX:tbfMp6zg],  /*line :1*/ string(Separator))
		zyQaXEA := 2 + rZEhR9f60tA*3
		if yAVPcVsd != eI480IIHU {
			zyQaXEA += 1 + yAVPcVsd - eI480IIHU
		}
		xWqeoMlrdna :=  /*line :1*/ make([]byte, zyQaXEA)
		oTJUmc1VT_o :=  /*line :1*/ copy(xWqeoMlrdna, "..")
		for oesU8hV6 := 0; oesU8hV6 < rZEhR9f60tA; oesU8hV6++ {
			xWqeoMlrdna[oTJUmc1VT_o] = Separator
			 /*line :1*/ copy(xWqeoMlrdna[oTJUmc1VT_o+1:], "..")
			oTJUmc1VT_o += 3
		}
		if eI480IIHU != yAVPcVsd {
			xWqeoMlrdna[oTJUmc1VT_o] = Separator
			 /*line :1*/ copy(xWqeoMlrdna[oTJUmc1VT_o+1:], fpm9Hez[eI480IIHU:])
		}
		return  /*line :1*/ string(xWqeoMlrdna), nil
	}
	return fpm9Hez[eI480IIHU:], nil
}




var JXknnH error = fs.X7Z1FBI33u




var ADmE6axkY error = fs.MTx5MBbQt2







































type FDFPHQ5 func(w9VHSyrN string, fH3M2H4R fs.HM4JM2, aC4eBbavAAtF error) error

var h7Agepf = os.Z0JXsY	


func oERxo9(w9VHSyrN string, qFtGgOeDz fs.RbfBMxa, rMPFNG fs.S_ytKQXs) error {
	if aC4eBbavAAtF :=  /*line :1*/ rMPFNG(w9VHSyrN, qFtGgOeDz, nil); aC4eBbavAAtF != nil || ! /*line :1*/ qFtGgOeDz.IsDir() {
		if aC4eBbavAAtF == JXknnH &&  /*line :1*/ qFtGgOeDz.IsDir() {

			aC4eBbavAAtF = nil
		}
		return aC4eBbavAAtF
	}

	k3l3Soo, aC4eBbavAAtF :=  /*line :1*/ slLqbF3lCkr(w9VHSyrN)
	if aC4eBbavAAtF != nil {

		aC4eBbavAAtF =  /*line :1*/ rMPFNG(w9VHSyrN, qFtGgOeDz, aC4eBbavAAtF)
		if aC4eBbavAAtF != nil {
			if aC4eBbavAAtF == JXknnH &&  /*line :1*/ qFtGgOeDz.IsDir() {
				aC4eBbavAAtF = nil
			}
			return aC4eBbavAAtF
		}
	}

	for _, p1TEchY := range k3l3Soo {
		eleDmj :=  /*line :1*/ SzdnHBv(w9VHSyrN,  /*line :1*/ p1TEchY.Name())
		if aC4eBbavAAtF :=  /*line :1*/ oERxo9(eleDmj, p1TEchY, rMPFNG); aC4eBbavAAtF != nil {
			if aC4eBbavAAtF == JXknnH {
				break
			}
			return aC4eBbavAAtF
		}
	}
	return nil
}


func mcQso0(w9VHSyrN string, fH3M2H4R fs.HM4JM2, aYGHJaGsFea FDFPHQ5) error {
	if ! /*line :1*/ fH3M2H4R.IsDir() {
		return  /*line :1*/ aYGHJaGsFea(w9VHSyrN, fH3M2H4R, nil)
	}

	j7XUfK5nkoJ, aC4eBbavAAtF :=  /*line :1*/ eeZpA85TEhN(w9VHSyrN)
	ecr4sOUb_AP :=  /*line :1*/ aYGHJaGsFea(w9VHSyrN, fH3M2H4R, aC4eBbavAAtF)

	if aC4eBbavAAtF != nil || ecr4sOUb_AP != nil {

		return ecr4sOUb_AP
	}

	for _, hurYsplzYEd := range j7XUfK5nkoJ {
		sSBB0TY :=  /*line :1*/ SzdnHBv(w9VHSyrN, hurYsplzYEd)
		dfgUVNoUJak, aC4eBbavAAtF :=  /*line :1*/ h7Agepf(sSBB0TY)
		if aC4eBbavAAtF != nil {
			if aC4eBbavAAtF :=  /*line :1*/ aYGHJaGsFea(sSBB0TY, dfgUVNoUJak, aC4eBbavAAtF); aC4eBbavAAtF != nil && aC4eBbavAAtF != JXknnH {
				return aC4eBbavAAtF
			}
		} else {
			aC4eBbavAAtF =  /*line :1*/ mcQso0(sSBB0TY, dfgUVNoUJak, aYGHJaGsFea)
			if aC4eBbavAAtF != nil {
				if ! /*line :1*/ dfgUVNoUJak.IsDir() || aC4eBbavAAtF != JXknnH {
					return aC4eBbavAAtF
				}
			}
		}
	}
	return nil
}
















func O2Oa2H3daWJ(p5zY0a2I0C string, iXa_dFykunW2 fs.S_ytKQXs) error {
	fH3M2H4R, aC4eBbavAAtF :=  /*line :1*/ os.Z0JXsY(p5zY0a2I0C)
	if aC4eBbavAAtF != nil {
		aC4eBbavAAtF =  /*line :1*/ iXa_dFykunW2(p5zY0a2I0C, nil, aC4eBbavAAtF)
	} else {
		aC4eBbavAAtF =  /*line :1*/ oERxo9(p5zY0a2I0C, &akkg6z{fH3M2H4R}, iXa_dFykunW2)
	}
	if aC4eBbavAAtF == JXknnH || aC4eBbavAAtF == ADmE6axkY {
		return nil
	}
	return aC4eBbavAAtF
}

type akkg6z struct {
	aDXlH4zQP fs.HM4JM2
}

func (qFtGgOeDz *akkg6z) Name() string	{ return  /*line :1*/ qFtGgOeDz.aDXlH4zQP.Name() }
func (qFtGgOeDz *akkg6z) IsDir() bool	{ return  /*line :1*/ qFtGgOeDz.aDXlH4zQP.IsDir() }
func (qFtGgOeDz *akkg6z) Type() fs.PGXMbJjlHBMu	{ return  /*line :1*/ qFtGgOeDz.aDXlH4zQP.Mode().Type() }
func (qFtGgOeDz *akkg6z) Info() (fs.HM4JM2, error)	{ return qFtGgOeDz.aDXlH4zQP, nil }

func (qFtGgOeDz *akkg6z) String() string {
	return  /*line :1*/ fs.EMHCka(qFtGgOeDz)
}















func AMh107KiUAv(p5zY0a2I0C string, iXa_dFykunW2 FDFPHQ5) error {
	fH3M2H4R, aC4eBbavAAtF :=  /*line :1*/ os.Z0JXsY(p5zY0a2I0C)
	if aC4eBbavAAtF != nil {
		aC4eBbavAAtF =  /*line :1*/ iXa_dFykunW2(p5zY0a2I0C, nil, aC4eBbavAAtF)
	} else {
		aC4eBbavAAtF =  /*line :1*/ mcQso0(p5zY0a2I0C, fH3M2H4R, iXa_dFykunW2)
	}
	if aC4eBbavAAtF == JXknnH || aC4eBbavAAtF == ADmE6axkY {
		return nil
	}
	return aC4eBbavAAtF
}



func slLqbF3lCkr(yTAz5Wp4m string) ([]fs.RbfBMxa, error) {
	pvNirJV8qP, aC4eBbavAAtF :=  /*line :1*/ os.Cvk6wxcjw(yTAz5Wp4m)
	if aC4eBbavAAtF != nil {
		return nil, aC4eBbavAAtF
	}
	k3l3Soo, aC4eBbavAAtF :=  /*line :1*/ pvNirJV8qP.ReadDir(-1)
	 /*line :1*/ pvNirJV8qP.Close()
	if aC4eBbavAAtF != nil {
		return nil, aC4eBbavAAtF
	}
	 /*line :1*/ sort.Wo2XqNIPTLLO(k3l3Soo, func(oesU8hV6, hcpJJrEA2kri int) bool { return  /*line :1*/ k3l3Soo[oesU8hV6].Name() <  /*line :1*/ k3l3Soo[hcpJJrEA2kri].Name() })
	return k3l3Soo, nil
}



func eeZpA85TEhN(yTAz5Wp4m string) ([]string, error) {
	pvNirJV8qP, aC4eBbavAAtF :=  /*line :1*/ os.Cvk6wxcjw(yTAz5Wp4m)
	if aC4eBbavAAtF != nil {
		return nil, aC4eBbavAAtF
	}
	j7XUfK5nkoJ, aC4eBbavAAtF :=  /*line :1*/ pvNirJV8qP.Readdirnames(-1)
	 /*line :1*/ pvNirJV8qP.Close()
	if aC4eBbavAAtF != nil {
		return nil, aC4eBbavAAtF
	}
	 /*line :1*/ sort.IsJh7v(j7XUfK5nkoJ)
	return j7XUfK5nkoJ, nil
}





func XLL_5aSxD(w9VHSyrN string) string {
	if w9VHSyrN == "" {
		return "."
	}

	for  /*line :1*/ len(w9VHSyrN) > 0 &&  /*line :1*/ os.GClWKECc(w9VHSyrN[ /*line :1*/ len(w9VHSyrN)-1]) {
		w9VHSyrN = w9VHSyrN[0 :  /*line :1*/ len(w9VHSyrN)-1]
	}

	w9VHSyrN = w9VHSyrN[ /*line :1*/ len( /*line :1*/ SDvSg10pbg(w9VHSyrN)):]

	oesU8hV6 :=  /*line :1*/ len(w9VHSyrN) - 1
	for oesU8hV6 >= 0 && ! /*line :1*/ os.GClWKECc(w9VHSyrN[oesU8hV6]) {
		oesU8hV6--
	}
	if oesU8hV6 >= 0 {
		w9VHSyrN = w9VHSyrN[oesU8hV6+1:]
	}

	if w9VHSyrN == "" {
		return  /*line :1*/ string(Separator)
	}
	return w9VHSyrN
}







func XYCNmV(w9VHSyrN string) string {
	d2oIeyTClzU :=  /*line :1*/ SDvSg10pbg(w9VHSyrN)
	oesU8hV6 :=  /*line :1*/ len(w9VHSyrN) - 1
	for oesU8hV6 >=  /*line :1*/ len(d2oIeyTClzU) && ! /*line :1*/ os.GClWKECc(w9VHSyrN[oesU8hV6]) {
		oesU8hV6--
	}
	gzQrdT03qcU6 :=  /*line :1*/ ZBiAXr(w9VHSyrN[ /*line :1*/ len(d2oIeyTClzU) : oesU8hV6+1])
	if gzQrdT03qcU6 == "." &&  /*line :1*/ len(d2oIeyTClzU) > 2 {

		return d2oIeyTClzU
	}
	return d2oIeyTClzU + gzQrdT03qcU6
}





func SDvSg10pbg(w9VHSyrN string) string {
	return  /*line :1*/ EBTjoI0Ix(w9VHSyrN[: /*line :1*/ sg1_tZpahJa(w9VHSyrN)])
}
