//line :1
//go:build windows

package NJ4MerJ

func dMi2GE(gZ3XVhDteAaT int) string {
	if gZ3XVhDteAaT < 0 {
		return "-" +  /*line :1*/ dMi2GE(-gZ3XVhDteAaT)
	}
	var etRYtnVni [32]byte
	rRGfxgPH8Kq :=  /*line :1*/ len(etRYtnVni) - 1
	for gZ3XVhDteAaT >= 10 {
		etRYtnVni[rRGfxgPH8Kq] =  /*line :1*/ byte(gZ3XVhDteAaT%10 + '0')
		rRGfxgPH8Kq--
		gZ3XVhDteAaT /= 10
	}
	etRYtnVni[rRGfxgPH8Kq] =  /*line :1*/ byte(gZ3XVhDteAaT + '0')
	return  /*line :1*/ string(etRYtnVni[rRGfxgPH8Kq:])
}
