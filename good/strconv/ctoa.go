//line :1
package zBESu2SrRjP








func BGoEbm1e(uTN3BXbMd complex128, i54pnXiXYaZq byte, gsf1Ql, r2moFi1 int) string {
	if r2moFi1 != 64 && r2moFi1 != 128 {
		 /*line :1*/ panic("invalid bitSize")
	}
	r2moFi1 >>= 1

	j6TCto :=  /*line :1*/ AyQcIQvp( /*line :1*/ imag(uTN3BXbMd), i54pnXiXYaZq, gsf1Ql, r2moFi1)
	if j6TCto[0] != '+' && j6TCto[0] != '-' {
		j6TCto = "+" + j6TCto
	}

	return "(" +  /*line :1*/ AyQcIQvp( /*line :1*/ real(uTN3BXbMd), i54pnXiXYaZq, gsf1Ql, r2moFi1) + j6TCto + "i)"
}
