//line :1
package fRAfQd_

import errors "iAHoxjmM"

func (pv6eev G4KDkI) kigfvy0(ev8znPS9W []byte, f_KtSGVqzqM bool) []byte {
	_, l7D_bh2w, nsu4XSEmD7r :=  /*line :1*/ pv6eev.bNr1Yr()

	mO2sH8hUHnW, cmQVCZ2yilS, gv5Z468R, _ :=  /*line :1*/ fdO09EiS(nsu4XSEmD7r, true)
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, mO2sH8hUHnW, 4)
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '-')
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W,  /*line :1*/ int(cmQVCZ2yilS), 2)
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '-')
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, gv5Z468R, 2)

	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, 'T')

	xmMDTE, qJ3jbB5O, insjC2Va :=  /*line :1*/ mbFqMakX(nsu4XSEmD7r)
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, xmMDTE, 2)
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ':')
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, qJ3jbB5O, 2)
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ':')
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, insjC2Va, 2)

	if f_KtSGVqzqM {
		gA16DoaF_ :=  /*line :1*/ xDJV1RBUw(stdFracSecond9, 9, '.')
		ev8znPS9W =  /*line :1*/ qlW1Bw(ev8znPS9W,  /*line :1*/ pv6eev.Nanosecond(), gA16DoaF_)
	}

	if l7D_bh2w == 0 {
		return  /*line :1*/ append(ev8znPS9W, 'Z')
	}

	k3sFK3l := l7D_bh2w / 60
	if k3sFK3l < 0 {
		ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '-')
		k3sFK3l = -k3sFK3l
	} else {
		ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '+')
	}
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, k3sFK3l/60, 2)
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ':')
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, k3sFK3l%60, 2)
	return ev8znPS9W
}

func (pv6eev G4KDkI) dOyyIxcFJ(ev8znPS9W []byte) ([]byte, error) {
	sjySUu9E3ZG :=  /*line :1*/ len(ev8znPS9W)
	ev8znPS9W =  /*line :1*/ pv6eev.kigfvy0(ev8znPS9W, true)

	fTtaggPpLFt := func(ev8znPS9W []byte) byte { return 10*(ev8znPS9W[0]-'0') + (ev8znPS9W[1] - '0') }
	switch {
	case ev8znPS9W[sjySUu9E3ZG+ /*line :1*/ len("9999")] != '-':
		return ev8znPS9W,  /*line :1*/ errors.Su6g6hRoi9X("year outside of range [0,9999]")
	case ev8znPS9W[ /*line :1*/ len(ev8znPS9W)-1] != 'Z':
		fnpIad3KB := ev8znPS9W[ /*line :1*/ len(ev8znPS9W)- /*line :1*/ len("Z07:00")]
		if ('0' <= fnpIad3KB && fnpIad3KB <= '9') ||  /*line :1*/ fTtaggPpLFt(ev8znPS9W[ /*line :1*/ len(ev8znPS9W)- /*line :1*/ len("07:00"):]) >= 24 {
			return ev8znPS9W,  /*line :1*/ errors.Su6g6hRoi9X("timezone hour outside of range [0,23]")
		}
	}
	return ev8znPS9W, nil
}

func zlEkeU[xl2dJlVxAz []byte | string](cjQd5EdPBZv xl2dJlVxAz, b2m6A1Q *Hh4U1oidou) (G4KDkI, bool) {

	ufWfJe3qPw := true
	s8MYNdY1 := func(cjQd5EdPBZv xl2dJlVxAz, qJ3jbB5O, aRi90tFb0vWu int) (i_2w697Gglw int) {
		for _, fnpIad3KB := range [] /*line :1*/ byte(cjQd5EdPBZv) {
			if fnpIad3KB < '0' || '9' < fnpIad3KB {
				ufWfJe3qPw = false
				return qJ3jbB5O
			}
			i_2w697Gglw = i_2w697Gglw*10 +  /*line :1*/ int(fnpIad3KB) - '0'
		}
		if i_2w697Gglw < qJ3jbB5O || aRi90tFb0vWu < i_2w697Gglw {
			ufWfJe3qPw = false
			return qJ3jbB5O
		}
		return i_2w697Gglw
	}

	if  /*line :1*/ len(cjQd5EdPBZv) <  /*line :1*/ len("2006-01-02T15:04:05") {
		return G4KDkI{}, false
	}
	mO2sH8hUHnW :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[0:4], 0, 9999)
	cmQVCZ2yilS :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[5:7], 1, 12)
	gv5Z468R :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[8:10], 1,  /*line :1*/ rs4bvng4W6y( /*line :1*/ ZyPFXNxLcwpT(cmQVCZ2yilS), mO2sH8hUHnW))
	xmMDTE :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[11:13], 0, 23)
	qJ3jbB5O :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[14:16], 0, 59)
	insjC2Va :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[17:19], 0, 59)
	if !ufWfJe3qPw || !(cjQd5EdPBZv[4] == '-' && cjQd5EdPBZv[7] == '-' && cjQd5EdPBZv[10] == 'T' && cjQd5EdPBZv[13] == ':' && cjQd5EdPBZv[16] == ':') {
		return G4KDkI{}, false
	}
	cjQd5EdPBZv = cjQd5EdPBZv[19:]

	
	var aCiH2Z7qWLtw int
	if  /*line :1*/ len(cjQd5EdPBZv) >= 2 && cjQd5EdPBZv[0] == '.' &&  /*line :1*/ tcM28YO(cjQd5EdPBZv, 1) {
		aeBZDJEINbb := 2
		for ; aeBZDJEINbb <  /*line :1*/ len(cjQd5EdPBZv) &&  /*line :1*/ tcM28YO(cjQd5EdPBZv, aeBZDJEINbb); aeBZDJEINbb++ {
		}
		aCiH2Z7qWLtw, _, _ =  /*line :1*/ hmaS0CqyQ8oh(cjQd5EdPBZv, aeBZDJEINbb)
		cjQd5EdPBZv = cjQd5EdPBZv[aeBZDJEINbb:]
	}

	pv6eev :=  /*line :1*/ FaDSoi2w(mO2sH8hUHnW,  /*line :1*/ ZyPFXNxLcwpT(cmQVCZ2yilS), gv5Z468R, xmMDTE, qJ3jbB5O, insjC2Va, aCiH2Z7qWLtw, PD1NIUjTJ)
	if  /*line :1*/ len(cjQd5EdPBZv) != 1 || cjQd5EdPBZv[0] != 'Z' {
		if  /*line :1*/ len(cjQd5EdPBZv) !=  /*line :1*/ len("-07:00") {
			return G4KDkI{}, false
		}
		ggP7eA6hbC1 :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[1:3], 0, 23)
		cVr0qo :=  /*line :1*/ s8MYNdY1(cjQd5EdPBZv[4:6], 0, 59)
		if !ufWfJe3qPw || !((cjQd5EdPBZv[0] == '-' || cjQd5EdPBZv[0] == '+') && cjQd5EdPBZv[3] == ':') {
			return G4KDkI{}, false
		}
		zP6aWDTN := (ggP7eA6hbC1*60 + cVr0qo) * 60
		if cjQd5EdPBZv[0] == '-' {
			zP6aWDTN *= -1
		}
		 /*line :1*/ pv6eev.hOFVKuqUr91a(- /*line :1*/ int64(zP6aWDTN))

		if _, l7D_bh2w, _, _, _ :=  /*line :1*/ b2m6A1Q.oPdG5Tk5( /*line :1*/ pv6eev.huzSe5e()); l7D_bh2w == zP6aWDTN {
			 /*line :1*/ pv6eev.vt5HqRw(b2m6A1Q)
		} else {
			 /*line :1*/ pv6eev.vt5HqRw( /*line :1*/ Z586lZG("", zP6aWDTN))
		}
	}
	return pv6eev, true
}

func i7Ia_y(ev8znPS9W []byte) (G4KDkI, error) {
	pv6eev, ufWfJe3qPw :=  /*line :1*/ zlEkeU(ev8znPS9W, CH9afbXxKDfu)
	if !ufWfJe3qPw {
		pv6eev, xuMLYrB :=  /*line :1*/ KoKD_YnCna(RFC3339,  /*line :1*/ string(ev8znPS9W))
		if xuMLYrB != nil {
			return G4KDkI{}, xuMLYrB
		}

		fTtaggPpLFt := func(ev8znPS9W []byte) byte { return 10*(ev8znPS9W[0]-'0') + (ev8znPS9W[1] - '0') }
		switch {

		case true:
			return pv6eev, nil
		case ev8znPS9W[ /*line :1*/ len("2006-01-02T")+1] == ':':
			return G4KDkI{}, &Hcvjzaf83uX{RFC3339,  /*line :1*/ string(ev8znPS9W), "15",  /*line :1*/ string(ev8znPS9W[ /*line :1*/ len("2006-01-02T"):][:1]), ""}
		case ev8znPS9W[ /*line :1*/ len("2006-01-02T15:04:05")] == ',':
			return G4KDkI{}, &Hcvjzaf83uX{RFC3339,  /*line :1*/ string(ev8znPS9W), ".", ",", ""}
		case ev8znPS9W[ /*line :1*/ len(ev8znPS9W)-1] != 'Z':
			switch {
			case  /*line :1*/ fTtaggPpLFt(ev8znPS9W[ /*line :1*/ len(ev8znPS9W)- /*line :1*/ len("07:00"):]) >= 24:
				return G4KDkI{}, &Hcvjzaf83uX{RFC3339,  /*line :1*/ string(ev8znPS9W), "Z07:00",  /*line :1*/ string(ev8znPS9W[ /*line :1*/ len(ev8znPS9W)- /*line :1*/ len("Z07:00"):]), ": timezone hour out of range"}
			case  /*line :1*/ fTtaggPpLFt(ev8znPS9W[ /*line :1*/ len(ev8znPS9W)- /*line :1*/ len("00"):]) >= 60:
				return G4KDkI{}, &Hcvjzaf83uX{RFC3339,  /*line :1*/ string(ev8znPS9W), "Z07:00",  /*line :1*/ string(ev8znPS9W[ /*line :1*/ len(ev8znPS9W)- /*line :1*/ len("Z07:00"):]), ": timezone minute out of range"}
			}
		default:
			return G4KDkI{}, &Hcvjzaf83uX{RFC3339,  /*line :1*/ string(ev8znPS9W), RFC3339,  /*line :1*/ string(ev8znPS9W), ""}
		}
	}
	return pv6eev, nil
}
