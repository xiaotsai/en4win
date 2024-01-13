//line :1
//go:build !js

package gF9mZs2

import (
	errors "iAHoxjmM"
	"internal/bytealg"
	godebug "rayT9AgS"
	fs "y1BamVjnXsaa"
	os "hPMrClpC"
	"runtime"
	sync "sync"
	syscall "bUKeamF"
)

type q6Glr4X1u struct {
	zNbZIsB5DCBf	bool
	koVz4xl7	bool

	jfT0IkFk5bF0	int

	sOIfgKcXg	bool

	zwOKv0_N5f	string
	f3_Knd9AAFx9	dp5hV6K
}

type dp5hV6K int

const (
	mdnsFromSystem	dp5hV6K	= iota
	mdnsAssumeExists
	mdnsAssumeDoesNotExist
)

var (
	wmHXccPiQ2Ua	sync.LhBfwn6wa1x
	ema09qasY8q	= &q6Glr4X1u{zwOKv0_N5f: runtime.GOOS}
)

func yAKpTR() *q6Glr4X1u {
	 /*line :1*/ wmHXccPiQ2Ua.Do(zgBz7C)
	return ema09qasY8q
}

func zgBz7C() {
	wUJIV3kmbvJ, kJWPrEfNoA :=  /*line :1*/ f0rB9ML()
	ema09qasY8q.zNbZIsB5DCBf = netGoBuildTag || wUJIV3kmbvJ == "go"
	ema09qasY8q.koVz4xl7 = netCgoBuildTag || wUJIV3kmbvJ == "cgo"
	ema09qasY8q.jfT0IkFk5bF0 = kJWPrEfNoA

	if ema09qasY8q.jfT0IkFk5bF0 > 0 {
		defer func() {
			if  /*line :1*/ ema09qasY8q.jfT0IkFk5bF0 > 1 {
				 /*line :1*/ println("go package net: confVal.netCgo =", ema09qasY8q.koVz4xl7, " netGo =", ema09qasY8q.zNbZIsB5DCBf)
			}
			switch {
			case ema09qasY8q.zNbZIsB5DCBf:
				if netGoBuildTag {
					 /*line :1*/ println("go package net: built with netgo build tag; using Go's DNS resolver")
				} else {
					 /*line :1*/ println("go package net: GODEBUG setting forcing use of Go's resolver")
				}
			case !cgoAvailable:
				 /*line :1*/ println("go package net: cgo resolver not supported; using Go's DNS resolver")
			case ema09qasY8q.koVz4xl7 || ema09qasY8q.sOIfgKcXg:
				 /*line :1*/ println("go package net: using cgo DNS resolver")
			default:
				 /*line :1*/ println("go package net: dynamic selection of DNS resolver")
			}
		}()
	}

	ema09qasY8q.sOIfgKcXg = false

	if !cgoAvailable {
		return
	}

	if  /*line :1*/ oC4nRPNJLrb() {
		ema09qasY8q.sOIfgKcXg = true
		return
	}

	switch runtime.GOOS {
	case "plan9", "windows", "js", "wasip1":
		return
	}

	_, baWTIr :=  /*line :1*/ syscall.ITYf1B("LOCALDOMAIN")
	if baWTIr ||  /*line :1*/ os.LDjFPRBEz("RES_OPTIONS") != "" ||  /*line :1*/ os.LDjFPRBEz("HOSTALIASES") != "" {
		ema09qasY8q.sOIfgKcXg = true
		return
	}

	if runtime.GOOS == "openbsd" &&  /*line :1*/ os.LDjFPRBEz("ASR_CONFIG") != "" {
		ema09qasY8q.sOIfgKcXg = true
		return
	}
}

func oC4nRPNJLrb() bool {
	switch runtime.GOOS {

	case "windows", "plan9":
		return true

	case "darwin", "ios":
		return true

	case "android":
		return true

	default:
		return false
	}
}

func (hl8w2lFX *q6Glr4X1u) eL0xOIcrQop(yxhs4Z *W8Q2tfHAD) bool {
	return hl8w2lFX.zNbZIsB5DCBf ||  /*line :1*/ yxhs4Z.uza_Hv() || !cgoAvailable
}

func (hl8w2lFX *q6Glr4X1u) fUZe1pZV0J(yxhs4Z *W8Q2tfHAD, qxwkC3VxG0 string) (prHUfwjRzJ mNRpaDPvaD, vTBPTe *ivgt6m77Buqr) {
	if hl8w2lFX.jfT0IkFk5bF0 > 1 {
		defer func() {
			 /*line :1*/ print("go package net: addrLookupOrder(", qxwkC3VxG0, ") = ",  /*line :1*/ prHUfwjRzJ.String(), "\n")
		}()
	}
	return  /*line :1*/ hl8w2lFX.wZ_0y9bV(yxhs4Z, "")
}

func (hl8w2lFX *q6Glr4X1u) mNRpaDPvaD(yxhs4Z *W8Q2tfHAD, tvBreHB string) (prHUfwjRzJ mNRpaDPvaD, vTBPTe *ivgt6m77Buqr) {
	if hl8w2lFX.jfT0IkFk5bF0 > 1 {
		defer func() {
			 /*line :1*/ print("go package net: hostLookupOrder(", tvBreHB, ") = ",  /*line :1*/ prHUfwjRzJ.String(), "\n")
		}()
	}
	return  /*line :1*/ hl8w2lFX.wZ_0y9bV(yxhs4Z, tvBreHB)
}

func (hl8w2lFX *q6Glr4X1u) wZ_0y9bV(yxhs4Z *W8Q2tfHAD, tvBreHB string) (prHUfwjRzJ mNRpaDPvaD, vTBPTe *ivgt6m77Buqr) {

	var mwNKCID98DWV mNRpaDPvaD

	var xHcFdimw5Xa bool
	if  /*line :1*/ hl8w2lFX.eL0xOIcrQop(yxhs4Z) {

		switch hl8w2lFX.zwOKv0_N5f {
		case "windows":

			mwNKCID98DWV = hostLookupDNS
		default:
			mwNKCID98DWV = hostLookupFilesDNS
		}
		xHcFdimw5Xa = false
	} else if hl8w2lFX.koVz4xl7 {

		return hostLookupCgo, nil
	} else if hl8w2lFX.sOIfgKcXg {

		return hostLookupCgo, nil
	} else {

		if  /*line :1*/ bytealg.IndexByteString(tvBreHB, '\\') != -1 ||  /*line :1*/ bytealg.IndexByteString(tvBreHB, '%') != -1 {

			return hostLookupCgo, nil
		}

		mwNKCID98DWV = hostLookupCgo
		xHcFdimw5Xa = true
	}

	switch hl8w2lFX.zwOKv0_N5f {
	case "windows", "plan9", "android", "ios":
		return mwNKCID98DWV, nil
	}

	vTBPTe =  /*line :1*/ mT_rTj1r8a()

	if xHcFdimw5Xa && vTBPTe.fk7ahSp != nil && ! /*line :1*/ errors.COBastO_C(vTBPTe.fk7ahSp, fs.QhJWcuT5mSD) && ! /*line :1*/ errors.COBastO_C(vTBPTe.fk7ahSp, fs.E8Hgdf) {

		return hostLookupCgo, vTBPTe
	}

	if xHcFdimw5Xa && vTBPTe.cZSoP12 {

		return hostLookupCgo, vTBPTe
	}

	if hl8w2lFX.zwOKv0_N5f == "openbsd" {

		if  /*line :1*/ errors.COBastO_C(vTBPTe.fk7ahSp, fs.QhJWcuT5mSD) {
			return hostLookupFiles, vTBPTe
		}

		bp6HtmNGS7 := vTBPTe.bzn7srUlDz7H
		if  /*line :1*/ len(bp6HtmNGS7) == 0 {

			return hostLookupDNSFiles, vTBPTe
		}
		if  /*line :1*/ len(bp6HtmNGS7) < 1 ||  /*line :1*/ len(bp6HtmNGS7) > 2 {

			return mwNKCID98DWV, vTBPTe
		}
		switch bp6HtmNGS7[0] {
		case "bind":
			if  /*line :1*/ len(bp6HtmNGS7) == 2 {
				if bp6HtmNGS7[1] == "file" {
					return hostLookupDNSFiles, vTBPTe
				}

				return mwNKCID98DWV, vTBPTe
			}
			return hostLookupDNS, vTBPTe
		case "file":
			if  /*line :1*/ len(bp6HtmNGS7) == 2 {
				if bp6HtmNGS7[1] == "bind" {
					return hostLookupFilesDNS, vTBPTe
				}

				return mwNKCID98DWV, vTBPTe
			}
			return hostLookupFiles, vTBPTe
		default:

			return mwNKCID98DWV, vTBPTe
		}

	}

	if  /*line :1*/ umZUnV4x(tvBreHB, ".") {
		tvBreHB = tvBreHB[: /*line :1*/ len(tvBreHB)-1]
	}
	if xHcFdimw5Xa &&  /*line :1*/ jRYeKNE(tvBreHB, ".local") {

		return hostLookupCgo, vTBPTe
	}

	cjFUVvWN :=  /*line :1*/ kKNl81QAoLO()
	qS4zcqNdLbr := cjFUVvWN.sHk8loC5lC["hosts"]

	if  /*line :1*/ errors.COBastO_C(cjFUVvWN.czDPXGoFxc4, fs.QhJWcuT5mSD) || (cjFUVvWN.czDPXGoFxc4 == nil &&  /*line :1*/ len(qS4zcqNdLbr) == 0) {
		if xHcFdimw5Xa && hl8w2lFX.zwOKv0_N5f == "solaris" {

			return hostLookupCgo, vTBPTe
		}

		return hostLookupFilesDNS, vTBPTe
	}
	if cjFUVvWN.czDPXGoFxc4 != nil {

		return mwNKCID98DWV, vTBPTe
	}

	var unbjiEe bool
	var a_LWNSi0 bool

	var yqgODWG2zQr, yBZ08cm bool
	var g_D_9T string
	for eaAUaB7Z, i0qXWLS := range qS4zcqNdLbr {
		if i0qXWLS.onlkrKzJ == "files" || i0qXWLS.onlkrKzJ == "dns" {
			if xHcFdimw5Xa && ! /*line :1*/ i0qXWLS.gs8yLpU5_R2F() {

				return hostLookupCgo, vTBPTe
			}
			if i0qXWLS.onlkrKzJ == "files" {
				yqgODWG2zQr = true
			} else {
				unbjiEe = true
				a_LWNSi0 = true
				yBZ08cm = true
			}
			if g_D_9T == "" {
				g_D_9T = i0qXWLS.onlkrKzJ
			}
			continue
		}

		if xHcFdimw5Xa {
			switch {
			case tvBreHB != "" && i0qXWLS.onlkrKzJ == "myhostname":

				if  /*line :1*/ rjrQUFe(tvBreHB) ||  /*line :1*/ ctVnV9yOo(tvBreHB) ||  /*line :1*/ bKqDCUnDWucN(tvBreHB) {
					return hostLookupCgo, vTBPTe
				}
				uSUaDsh, h_ljk48Bm :=  /*line :1*/ hxcx1fbWDc()
				if h_ljk48Bm != nil ||  /*line :1*/ hCjYM9U0fG(tvBreHB, uSUaDsh) {
					return hostLookupCgo, vTBPTe
				}
				continue
			case tvBreHB != "" &&  /*line :1*/ t9qAz2H(i0qXWLS.onlkrKzJ, "mdns"):

				var ngw58hNh bool
				switch hl8w2lFX.f3_Knd9AAFx9 {
				case mdnsFromSystem:
					_, h_ljk48Bm :=  /*line :1*/ os.ZYa3wuv("/etc/mdns.allow")
					if h_ljk48Bm != nil && ! /*line :1*/ errors.COBastO_C(h_ljk48Bm, fs.QhJWcuT5mSD) {

						return hostLookupCgo, vTBPTe
					}
					ngw58hNh = h_ljk48Bm == nil
				case mdnsAssumeExists:
					ngw58hNh = true
				case mdnsAssumeDoesNotExist:
					ngw58hNh = false
				}
				if ngw58hNh {
					return hostLookupCgo, vTBPTe
				}
				continue
			default:

				return hostLookupCgo, vTBPTe
			}
		}

		if !a_LWNSi0 {
			a_LWNSi0 = true
			for _, ljsCSk := range qS4zcqNdLbr[eaAUaB7Z+1:] {
				if ljsCSk.onlkrKzJ == "dns" {
					unbjiEe = true
					break
				}
			}
		}

		if !unbjiEe {
			yBZ08cm = true
			if g_D_9T == "" {
				g_D_9T = "dns"
			}
		}
	}

	switch {
	case yqgODWG2zQr && yBZ08cm:
		if g_D_9T == "files" {
			return hostLookupFilesDNS, vTBPTe
		} else {
			return hostLookupDNSFiles, vTBPTe
		}
	case yqgODWG2zQr:
		return hostLookupFiles, vTBPTe
	case yBZ08cm:
		return hostLookupDNS, vTBPTe
	}

	return mwNKCID98DWV, vTBPTe
}

var e0tdgGBzRwJ =  /*line :1*/ godebug.HppHHAII("netdns")

func f0rB9ML() (wUJIV3kmbvJ string, kJWPrEfNoA int) {
	t9pawXrl :=  /*line :1*/ e0tdgGBzRwJ.Value()
	vl6WHt := func(dRoFccaZ string) {
		if dRoFccaZ == "" {
			return
		}
		if '0' <= dRoFccaZ[0] && dRoFccaZ[0] <= '9' {
			kJWPrEfNoA, _, _ =  /*line :1*/ bEihJjrbz(dRoFccaZ)
		} else {
			wUJIV3kmbvJ = dRoFccaZ
		}
	}
	if eaAUaB7Z :=  /*line :1*/ bytealg.IndexByteString(t9pawXrl, '+'); eaAUaB7Z != -1 {
		 /*line :1*/ vl6WHt(t9pawXrl[:eaAUaB7Z])
		 /*line :1*/ vl6WHt(t9pawXrl[eaAUaB7Z+1:])
		return
	}
	 /*line :1*/ vl6WHt(t9pawXrl)
	return
}

func rjrQUFe(cBprzKT8 string) bool {
	return  /*line :1*/ hCjYM9U0fG(cBprzKT8, "localhost") ||  /*line :1*/ hCjYM9U0fG(cBprzKT8, "localhost.localdomain") ||  /*line :1*/ jRYeKNE(cBprzKT8, ".localhost") ||  /*line :1*/ jRYeKNE(cBprzKT8, ".localhost.localdomain")
}

func ctVnV9yOo(cBprzKT8 string) bool {
	return  /*line :1*/ hCjYM9U0fG(cBprzKT8, "_gateway")
}

func bKqDCUnDWucN(cBprzKT8 string) bool {
	return  /*line :1*/ hCjYM9U0fG(cBprzKT8, "_outbound")
}
