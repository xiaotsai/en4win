//line :1
//go:generate go run makeisprint.go -output isprint.go

package zBESu2SrRjP

import (
	utf8 "CuAc0pPZwf"
)

const (
	lowerhex	= "0123456789abcdef"
	upperhex	= "0123456789ABCDEF"
)

func tBzRkkLqjEH(a0M8QSnCR string, uTN3BXbMd byte) bool {
	return  /*line :1*/ xyI8lfjixa(a0M8QSnCR, uTN3BXbMd) != -1
}

func wibl4LQ(a0M8QSnCR string, iIaVsq byte, ZEUQ_V_W, ku0MJSanKX bool) string {
	return  /*line :1*/ string( /*line :1*/ kPSpHC709Q( /*line :1*/ make([]byte, 0, 3* /*line :1*/ len(a0M8QSnCR)/2), a0M8QSnCR, iIaVsq, ZEUQ_V_W, ku0MJSanKX))
}

func tUVJnjuhbqi(vaeEFNU6wEP rune, iIaVsq byte, ZEUQ_V_W, ku0MJSanKX bool) string {
	return  /*line :1*/ string( /*line :1*/ k4Z8ShM_(nil, vaeEFNU6wEP, iIaVsq, ZEUQ_V_W, ku0MJSanKX))
}

func kPSpHC709Q(uPg2dG3l8KRn []byte, a0M8QSnCR string, iIaVsq byte, ZEUQ_V_W, ku0MJSanKX bool) []byte {

	if  /*line :1*/ cap(uPg2dG3l8KRn)- /*line :1*/ len(uPg2dG3l8KRn) <  /*line :1*/ len(a0M8QSnCR) {
		v_tuusQ :=  /*line :1*/ make([]byte,  /*line :1*/ len(uPg2dG3l8KRn),  /*line :1*/ len(uPg2dG3l8KRn)+1+ /*line :1*/ len(a0M8QSnCR)+1)
		 /*line :1*/ copy(v_tuusQ, uPg2dG3l8KRn)
		uPg2dG3l8KRn = v_tuusQ
	}
	uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, iIaVsq)
	for xsPSvYGujdL := 0;  /*line :1*/ len(a0M8QSnCR) > 0; a0M8QSnCR = a0M8QSnCR[xsPSvYGujdL:] {
		vaeEFNU6wEP :=  /*line :1*/ rune(a0M8QSnCR[0])
		xsPSvYGujdL = 1
		if vaeEFNU6wEP >= utf8.RuneSelf {
			vaeEFNU6wEP, xsPSvYGujdL =  /*line :1*/ utf8.HVAV7aTqxzg(a0M8QSnCR)
		}
		if xsPSvYGujdL == 1 && vaeEFNU6wEP == utf8.RuneError {
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\x`...)
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, lowerhex[a0M8QSnCR[0]>>4])
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, lowerhex[a0M8QSnCR[0]&0xF])
			continue
		}
		uPg2dG3l8KRn =  /*line :1*/ dUzUsR11sya(uPg2dG3l8KRn, vaeEFNU6wEP, iIaVsq, ZEUQ_V_W, ku0MJSanKX)
	}
	uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, iIaVsq)
	return uPg2dG3l8KRn
}

func k4Z8ShM_(uPg2dG3l8KRn []byte, vaeEFNU6wEP rune, iIaVsq byte, ZEUQ_V_W, ku0MJSanKX bool) []byte {
	uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, iIaVsq)
	if ! /*line :1*/ utf8.DYq33yNk(vaeEFNU6wEP) {
		vaeEFNU6wEP = utf8.RuneError
	}
	uPg2dG3l8KRn =  /*line :1*/ dUzUsR11sya(uPg2dG3l8KRn, vaeEFNU6wEP, iIaVsq, ZEUQ_V_W, ku0MJSanKX)
	uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, iIaVsq)
	return uPg2dG3l8KRn
}

func dUzUsR11sya(uPg2dG3l8KRn []byte, vaeEFNU6wEP rune, iIaVsq byte, ZEUQ_V_W, ku0MJSanKX bool) []byte {
	var cJq57Q9UKlU [utf8.UTFMax]byte
	if vaeEFNU6wEP ==  /*line :1*/ rune(iIaVsq) || vaeEFNU6wEP == '\\' {
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, '\\')
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn,  /*line :1*/ byte(vaeEFNU6wEP))
		return uPg2dG3l8KRn
	}
	if ZEUQ_V_W {
		if vaeEFNU6wEP < utf8.RuneSelf &&  /*line :1*/ ZKRkeJe6OgjL(vaeEFNU6wEP) {
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn,  /*line :1*/ byte(vaeEFNU6wEP))
			return uPg2dG3l8KRn
		}
	} else if  /*line :1*/ ZKRkeJe6OgjL(vaeEFNU6wEP) || ku0MJSanKX &&  /*line :1*/ thuLnYwQ(vaeEFNU6wEP) {
		wESnmZiAO :=  /*line :1*/ utf8.YoEca6(cJq57Q9UKlU[:], vaeEFNU6wEP)
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, cJq57Q9UKlU[:wESnmZiAO]...)
		return uPg2dG3l8KRn
	}
	switch vaeEFNU6wEP {
	case '\a':
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\a`...)
	case '\b':
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\b`...)
	case '\f':
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\f`...)
	case '\n':
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\n`...)
	case '\r':
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\r`...)
	case '\t':
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\t`...)
	case '\v':
		uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\v`...)
	default:
		switch {
		case vaeEFNU6wEP < ' ' || vaeEFNU6wEP == 0x7f:
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\x`...)
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, lowerhex[ /*line :1*/ byte(vaeEFNU6wEP)>>4])
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, lowerhex[ /*line :1*/ byte(vaeEFNU6wEP)&0xF])
		case ! /*line :1*/ utf8.DYq33yNk(vaeEFNU6wEP):
			vaeEFNU6wEP = 0xFFFD
			fallthrough
		case vaeEFNU6wEP < 0x10000:
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\u`...)
			for a0M8QSnCR := 12; a0M8QSnCR >= 0; a0M8QSnCR -= 4 {
				uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, lowerhex[vaeEFNU6wEP>> /*line :1*/ uint(a0M8QSnCR)&0xF])
			}
		default:
			uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, `\U`...)
			for a0M8QSnCR := 28; a0M8QSnCR >= 0; a0M8QSnCR -= 4 {
				uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, lowerhex[vaeEFNU6wEP>> /*line :1*/ uint(a0M8QSnCR)&0xF])
			}
		}
	}
	return uPg2dG3l8KRn
}

func D35LD5nn(a0M8QSnCR string) string {
	return  /*line :1*/ wibl4LQ(a0M8QSnCR, '"', false, false)
}

func TwCpqD6gFe(l_A2Ytb []byte, a0M8QSnCR string) []byte {
	return  /*line :1*/ kPSpHC709Q(l_A2Ytb, a0M8QSnCR, '"', false, false)
}

func PKvc_bZ23BEl(a0M8QSnCR string) string {
	return  /*line :1*/ wibl4LQ(a0M8QSnCR, '"', true, false)
}

func Ip59Nj9(l_A2Ytb []byte, a0M8QSnCR string) []byte {
	return  /*line :1*/ kPSpHC709Q(l_A2Ytb, a0M8QSnCR, '"', true, false)
}

func CmSRTVfhl8mO(a0M8QSnCR string) string {
	return  /*line :1*/ wibl4LQ(a0M8QSnCR, '"', false, true)
}

func Smog_hPTa3Z(l_A2Ytb []byte, a0M8QSnCR string) []byte {
	return  /*line :1*/ kPSpHC709Q(l_A2Ytb, a0M8QSnCR, '"', false, true)
}

func XnOUaUgjAd(vaeEFNU6wEP rune) string {
	return  /*line :1*/ tUVJnjuhbqi(vaeEFNU6wEP, '\'', false, false)
}

func RjM8k9(l_A2Ytb []byte, vaeEFNU6wEP rune) []byte {
	return  /*line :1*/ k4Z8ShM_(l_A2Ytb, vaeEFNU6wEP, '\'', false, false)
}

func BzNGV1(vaeEFNU6wEP rune) string {
	return  /*line :1*/ tUVJnjuhbqi(vaeEFNU6wEP, '\'', true, false)
}

func ASahUjlvcc(l_A2Ytb []byte, vaeEFNU6wEP rune) []byte {
	return  /*line :1*/ k4Z8ShM_(l_A2Ytb, vaeEFNU6wEP, '\'', true, false)
}

func GmJg2Kgd(vaeEFNU6wEP rune) string {
	return  /*line :1*/ tUVJnjuhbqi(vaeEFNU6wEP, '\'', false, true)
}

func D1Xm1p_(l_A2Ytb []byte, vaeEFNU6wEP rune) []byte {
	return  /*line :1*/ k4Z8ShM_(l_A2Ytb, vaeEFNU6wEP, '\'', false, true)
}

func DnaHMeUaZj(a0M8QSnCR string) bool {
	for  /*line :1*/ len(a0M8QSnCR) > 0 {
		vaeEFNU6wEP, nPCzab3Vr :=  /*line :1*/ utf8.HVAV7aTqxzg(a0M8QSnCR)
		a0M8QSnCR = a0M8QSnCR[nPCzab3Vr:]
		if nPCzab3Vr > 1 {
			if vaeEFNU6wEP == '\ufeff' {
				return false
			}
			continue
		}
		if vaeEFNU6wEP == utf8.RuneError {
			return false
		}
		if (vaeEFNU6wEP < ' ' && vaeEFNU6wEP != '\t') || vaeEFNU6wEP == '`' || vaeEFNU6wEP == '\u007F' {
			return false
		}
	}
	return true
}

func iZNPvpWOJpV(o7OjhYc5wQns byte) (ggqcIE rune, eva_XlRz4Iz bool) {
	uTN3BXbMd :=  /*line :1*/ rune(o7OjhYc5wQns)
	switch {
	case '0' <= uTN3BXbMd && uTN3BXbMd <= '9':
		return uTN3BXbMd - '0', true
	case 'a' <= uTN3BXbMd && uTN3BXbMd <= 'f':
		return uTN3BXbMd - 'a' + 10, true
	case 'A' <= uTN3BXbMd && uTN3BXbMd <= 'F':
		return uTN3BXbMd - 'A' + 10, true
	}
	return
}

func DtQqRoog7Udi(a0M8QSnCR string, iIaVsq byte) (l9qJXIv rune, l4rCSrLyiRr5 bool, cX_ilvVo8i string, kWNmhkfeXX error) {

	if  /*line :1*/ len(a0M8QSnCR) == 0 {
		kWNmhkfeXX = ZTbjMv
		return
	}
	switch uTN3BXbMd := a0M8QSnCR[0]; {
	case uTN3BXbMd == iIaVsq && (iIaVsq == '\'' || iIaVsq == '"'):
		kWNmhkfeXX = ZTbjMv
		return
	case uTN3BXbMd >= utf8.RuneSelf:
		vaeEFNU6wEP, mugT3utgZ :=  /*line :1*/ utf8.HVAV7aTqxzg(a0M8QSnCR)
		return vaeEFNU6wEP, true, a0M8QSnCR[mugT3utgZ:], nil
	case uTN3BXbMd != '\\':
		return  /*line :1*/ rune(a0M8QSnCR[0]), false, a0M8QSnCR[1:], nil
	}

	if  /*line :1*/ len(a0M8QSnCR) <= 1 {
		kWNmhkfeXX = ZTbjMv
		return
	}
	uTN3BXbMd := a0M8QSnCR[1]
	a0M8QSnCR = a0M8QSnCR[2:]

	switch uTN3BXbMd {
	case 'a':
		l9qJXIv = '\a'
	case 'b':
		l9qJXIv = '\b'
	case 'f':
		l9qJXIv = '\f'
	case 'n':
		l9qJXIv = '\n'
	case 'r':
		l9qJXIv = '\r'
	case 't':
		l9qJXIv = '\t'
	case 'v':
		l9qJXIv = '\v'
	case 'x', 'u', 'U':
		wESnmZiAO := 0
		switch uTN3BXbMd {
		case 'x':
			wESnmZiAO = 2
		case 'u':
			wESnmZiAO = 4
		case 'U':
			wESnmZiAO = 8
		}
		var ggqcIE rune
		if  /*line :1*/ len(a0M8QSnCR) < wESnmZiAO {
			kWNmhkfeXX = ZTbjMv
			return
		}
		for dTPavgbkmK := 0; dTPavgbkmK < wESnmZiAO; dTPavgbkmK++ {
			eswrOOpIaF, eva_XlRz4Iz :=  /*line :1*/ iZNPvpWOJpV(a0M8QSnCR[dTPavgbkmK])
			if !eva_XlRz4Iz {
				kWNmhkfeXX = ZTbjMv
				return
			}
			ggqcIE = ggqcIE<<4 | eswrOOpIaF
		}
		a0M8QSnCR = a0M8QSnCR[wESnmZiAO:]
		if uTN3BXbMd == 'x' {

			l9qJXIv = ggqcIE
			break
		}
		if ! /*line :1*/ utf8.DYq33yNk(ggqcIE) {
			kWNmhkfeXX = ZTbjMv
			return
		}
		l9qJXIv = ggqcIE
		l4rCSrLyiRr5 = true
	case '0', '1', '2', '3', '4', '5', '6', '7':
		ggqcIE :=  /*line :1*/ rune(uTN3BXbMd) - '0'
		if  /*line :1*/ len(a0M8QSnCR) < 2 {
			kWNmhkfeXX = ZTbjMv
			return
		}
		for dTPavgbkmK := 0; dTPavgbkmK < 2; dTPavgbkmK++ {
			eswrOOpIaF :=  /*line :1*/ rune(a0M8QSnCR[dTPavgbkmK]) - '0'
			if eswrOOpIaF < 0 || eswrOOpIaF > 7 {
				kWNmhkfeXX = ZTbjMv
				return
			}
			ggqcIE = (ggqcIE << 3) | eswrOOpIaF
		}
		a0M8QSnCR = a0M8QSnCR[2:]
		if ggqcIE > 255 {
			kWNmhkfeXX = ZTbjMv
			return
		}
		l9qJXIv = ggqcIE
	case '\\':
		l9qJXIv = '\\'
	case '\'', '"':
		if uTN3BXbMd != iIaVsq {
			kWNmhkfeXX = ZTbjMv
			return
		}
		l9qJXIv =  /*line :1*/ rune(uTN3BXbMd)
	default:
		kWNmhkfeXX = ZTbjMv
		return
	}
	cX_ilvVo8i = a0M8QSnCR
	return
}

func IOaYdtkr6z2v(a0M8QSnCR string) (string, error) {
	p5Dst_SYq, _, kWNmhkfeXX :=  /*line :1*/ vxvmz3i4(a0M8QSnCR, false)
	return p5Dst_SYq, kWNmhkfeXX
}

func D0aVL66(a0M8QSnCR string) (string, error) {
	p5Dst_SYq, pNRglmCkMAx, kWNmhkfeXX :=  /*line :1*/ vxvmz3i4(a0M8QSnCR, true)
	if  /*line :1*/ len(pNRglmCkMAx) > 0 {
		return "", ZTbjMv
	}
	return p5Dst_SYq, kWNmhkfeXX
}

func vxvmz3i4(pBC8fE06v3XG string, zdMmzT11 bool) (p5Dst_SYq, pNRglmCkMAx string, kWNmhkfeXX error) {

	if  /*line :1*/ len(pBC8fE06v3XG) < 2 {
		return "", pBC8fE06v3XG, ZTbjMv
	}
	iIaVsq := pBC8fE06v3XG[0]
	abVIzcgq4Y :=  /*line :1*/ xyI8lfjixa(pBC8fE06v3XG[1:], iIaVsq)
	if abVIzcgq4Y < 0 {
		return "", pBC8fE06v3XG, ZTbjMv
	}
	abVIzcgq4Y += 2

	switch iIaVsq {
	case '`':
		switch {
		case !zdMmzT11:
			p5Dst_SYq = pBC8fE06v3XG[:abVIzcgq4Y]
		case ! /*line :1*/ tBzRkkLqjEH(pBC8fE06v3XG[:abVIzcgq4Y], '\r'):
			p5Dst_SYq = pBC8fE06v3XG[ /*line :1*/ len("`") : abVIzcgq4Y- /*line :1*/ len("`")]
		default:

			uPg2dG3l8KRn :=  /*line :1*/ make([]byte, 0, abVIzcgq4Y- /*line :1*/ len("`")- /*line :1*/ len("\r")- /*line :1*/ len("`"))
			for cDhV_2D :=  /*line :1*/ len("`"); cDhV_2D < abVIzcgq4Y- /*line :1*/ len("`"); cDhV_2D++ {
				if pBC8fE06v3XG[cDhV_2D] != '\r' {
					uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, pBC8fE06v3XG[cDhV_2D])
				}
			}
			p5Dst_SYq =  /*line :1*/ string(uPg2dG3l8KRn)
		}

		return p5Dst_SYq, pBC8fE06v3XG[abVIzcgq4Y:], nil
	case '"', '\'':

		if ! /*line :1*/ tBzRkkLqjEH(pBC8fE06v3XG[:abVIzcgq4Y], '\\') && ! /*line :1*/ tBzRkkLqjEH(pBC8fE06v3XG[:abVIzcgq4Y], '\n') {
			var j3ZCCq29ywbZ bool
			switch iIaVsq {
			case '"':
				j3ZCCq29ywbZ =  /*line :1*/ utf8.YiR9Xl_UC5x(pBC8fE06v3XG[ /*line :1*/ len(`"`) : abVIzcgq4Y- /*line :1*/ len(`"`)])
			case '\'':
				vaeEFNU6wEP, wESnmZiAO :=  /*line :1*/ utf8.HVAV7aTqxzg(pBC8fE06v3XG[ /*line :1*/ len("'") : abVIzcgq4Y- /*line :1*/ len("'")])
				j3ZCCq29ywbZ =  /*line :1*/ len("'")+wESnmZiAO+ /*line :1*/ len("'") == abVIzcgq4Y && (vaeEFNU6wEP != utf8.RuneError || wESnmZiAO != 1)
			}
			if j3ZCCq29ywbZ {
				p5Dst_SYq = pBC8fE06v3XG[:abVIzcgq4Y]
				if zdMmzT11 {
					p5Dst_SYq = p5Dst_SYq[1 : abVIzcgq4Y-1]
				}
				return p5Dst_SYq, pBC8fE06v3XG[abVIzcgq4Y:], nil
			}
		}

		var uPg2dG3l8KRn []byte
		noiSIfQ := pBC8fE06v3XG
		pBC8fE06v3XG = pBC8fE06v3XG[1:]
		if zdMmzT11 {
			uPg2dG3l8KRn =  /*line :1*/ make([]byte, 0, 3*abVIzcgq4Y/2)
		}
		for  /*line :1*/ len(pBC8fE06v3XG) > 0 && pBC8fE06v3XG[0] != iIaVsq {

			vaeEFNU6wEP, l4rCSrLyiRr5, pNRglmCkMAx, kWNmhkfeXX :=  /*line :1*/ DtQqRoog7Udi(pBC8fE06v3XG, iIaVsq)
			if pBC8fE06v3XG[0] == '\n' || kWNmhkfeXX != nil {
				return "", noiSIfQ, ZTbjMv
			}
			pBC8fE06v3XG = pNRglmCkMAx

			if zdMmzT11 {
				if vaeEFNU6wEP < utf8.RuneSelf || !l4rCSrLyiRr5 {
					uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn,  /*line :1*/ byte(vaeEFNU6wEP))
				} else {
					var refqsA [utf8.UTFMax]byte
					wESnmZiAO :=  /*line :1*/ utf8.YoEca6(refqsA[:], vaeEFNU6wEP)
					uPg2dG3l8KRn =  /*line :1*/ append(uPg2dG3l8KRn, refqsA[:wESnmZiAO]...)
				}
			}

			if iIaVsq == '\'' {
				break
			}
		}

		if !( /*line :1*/ len(pBC8fE06v3XG) > 0 && pBC8fE06v3XG[0] == iIaVsq) {
			return "", noiSIfQ, ZTbjMv
		}
		pBC8fE06v3XG = pBC8fE06v3XG[1:]

		if zdMmzT11 {
			return  /*line :1*/ string(uPg2dG3l8KRn), pBC8fE06v3XG, nil
		}
		return noiSIfQ[: /*line :1*/ len(noiSIfQ)- /*line :1*/ len(pBC8fE06v3XG)], pBC8fE06v3XG, nil
	default:
		return "", pBC8fE06v3XG, ZTbjMv
	}
}

func z_4gwrICp(zxZt5L []uint16, eswrOOpIaF uint16) int {
	cDhV_2D, dTPavgbkmK := 0,  /*line :1*/ len(zxZt5L)
	for cDhV_2D < dTPavgbkmK {
		fI5sBwPyC := cDhV_2D + (dTPavgbkmK-cDhV_2D)>>1
		if zxZt5L[fI5sBwPyC] < eswrOOpIaF {
			cDhV_2D = fI5sBwPyC + 1
		} else {
			dTPavgbkmK = fI5sBwPyC
		}
	}
	return cDhV_2D
}

func m0jukt(zxZt5L []uint32, eswrOOpIaF uint32) int {
	cDhV_2D, dTPavgbkmK := 0,  /*line :1*/ len(zxZt5L)
	for cDhV_2D < dTPavgbkmK {
		fI5sBwPyC := cDhV_2D + (dTPavgbkmK-cDhV_2D)>>1
		if zxZt5L[fI5sBwPyC] < eswrOOpIaF {
			cDhV_2D = fI5sBwPyC + 1
		} else {
			dTPavgbkmK = fI5sBwPyC
		}
	}
	return cDhV_2D
}

func ZKRkeJe6OgjL(vaeEFNU6wEP rune) bool {

	if vaeEFNU6wEP <= 0xFF {
		if 0x20 <= vaeEFNU6wEP && vaeEFNU6wEP <= 0x7E {

			return true
		}
		if 0xA1 <= vaeEFNU6wEP && vaeEFNU6wEP <= 0xFF {

			return vaeEFNU6wEP != 0xAD
		}
		return false
	}

	if 0 <= vaeEFNU6wEP && vaeEFNU6wEP < 1<<16 {
		_MRjfQgZ, sJ_dQePE, fw6kucbUYfPG :=  /*line :1*/ uint16(vaeEFNU6wEP), pbge2RaGhhL, lQ9qYerqi
		cDhV_2D :=  /*line :1*/ z_4gwrICp(sJ_dQePE, _MRjfQgZ)
		if cDhV_2D >=  /*line :1*/ len(sJ_dQePE) || _MRjfQgZ < sJ_dQePE[cDhV_2D&^1] || sJ_dQePE[cDhV_2D|1] < _MRjfQgZ {
			return false
		}
		dTPavgbkmK :=  /*line :1*/ z_4gwrICp(fw6kucbUYfPG, _MRjfQgZ)
		return dTPavgbkmK >=  /*line :1*/ len(fw6kucbUYfPG) || fw6kucbUYfPG[dTPavgbkmK] != _MRjfQgZ
	}

	_MRjfQgZ, sJ_dQePE, fw6kucbUYfPG :=  /*line :1*/ uint32(vaeEFNU6wEP), k6SDRYqzVb, klhbfspz
	cDhV_2D :=  /*line :1*/ m0jukt(sJ_dQePE, _MRjfQgZ)
	if cDhV_2D >=  /*line :1*/ len(sJ_dQePE) || _MRjfQgZ < sJ_dQePE[cDhV_2D&^1] || sJ_dQePE[cDhV_2D|1] < _MRjfQgZ {
		return false
	}
	if vaeEFNU6wEP >= 0x20000 {
		return true
	}
	vaeEFNU6wEP -= 0x10000
	dTPavgbkmK :=  /*line :1*/ z_4gwrICp(fw6kucbUYfPG,  /*line :1*/ uint16(vaeEFNU6wEP))
	return dTPavgbkmK >=  /*line :1*/ len(fw6kucbUYfPG) || fw6kucbUYfPG[dTPavgbkmK] !=  /*line :1*/ uint16(vaeEFNU6wEP)
}

func TziirLO(vaeEFNU6wEP rune) bool {
	if  /*line :1*/ ZKRkeJe6OgjL(vaeEFNU6wEP) {
		return true
	}
	return  /*line :1*/ thuLnYwQ(vaeEFNU6wEP)
}

func thuLnYwQ(vaeEFNU6wEP rune) bool {

	if vaeEFNU6wEP > 0xFFFF {
		return false
	}
	_MRjfQgZ :=  /*line :1*/ uint16(vaeEFNU6wEP)
	cDhV_2D :=  /*line :1*/ z_4gwrICp(shzfaqfSl, _MRjfQgZ)
	return cDhV_2D <  /*line :1*/ len(shzfaqfSl) && _MRjfQgZ == shzfaqfSl[cDhV_2D]
}
