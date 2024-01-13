//line :1
package math










func FmjpDXL(fzJsoqu, gNDglc3F2QV float64) float64 {
	if haveArchHypot {
		return  /*line :1*/ wKHDWSQ2sk(fzJsoqu, gNDglc3F2QV)
	}
	return  /*line :1*/ akCQMQmkWUmP(fzJsoqu, gNDglc3F2QV)
}

func akCQMQmkWUmP(fzJsoqu, gNDglc3F2QV float64) float64 {
	fzJsoqu, gNDglc3F2QV =  /*line :1*/ Abs(fzJsoqu),  /*line :1*/ Abs(gNDglc3F2QV)

	switch {
	case  /*line :1*/ ME1vzpD5wcr(fzJsoqu, 1) ||  /*line :1*/ ME1vzpD5wcr(gNDglc3F2QV, 1):
		return  /*line :1*/ FSug4WHZSBwz(1)
	case  /*line :1*/ OIdLpDqq(fzJsoqu) ||  /*line :1*/ OIdLpDqq(gNDglc3F2QV):
		return  /*line :1*/ WG0hZIT4()
	}
	if fzJsoqu < gNDglc3F2QV {
		fzJsoqu, gNDglc3F2QV = gNDglc3F2QV, fzJsoqu
	}
	if fzJsoqu == 0 {
		return 0
	}
	gNDglc3F2QV = gNDglc3F2QV / fzJsoqu
	return fzJsoqu *  /*line :1*/ NeXs7bSyfaCD(1+gNDglc3F2QV*gNDglc3F2QV)
}
