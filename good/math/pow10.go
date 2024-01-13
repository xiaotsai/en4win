//line :1
package math


var wd0f5H = [...]float64{
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
	1e20, 1e21, 1e22, 1e23, 1e24, 1e25, 1e26, 1e27, 1e28, 1e29,
	1e30, 1e31,
}


var toI3lXMWIU = [...]float64{
	1e00, 1e32, 1e64, 1e96, 1e128, 1e160, 1e192, 1e224, 1e256, 1e288,
}


var vHzmUaBZ = [...]float64{
	1e-00, 1e-32, 1e-64, 1e-96, 1e-128, 1e-160, 1e-192, 1e-224, 1e-256, 1e-288, 1e-320,
}







func HWKzDVBT9(d2X2v9 int) float64 {
	if 0 <= d2X2v9 && d2X2v9 <= 308 {
		return toI3lXMWIU[ /*line :1*/ uint(d2X2v9)/32] * wd0f5H[ /*line :1*/ uint(d2X2v9)%32]
	}

	if -323 <= d2X2v9 && d2X2v9 <= 0 {
		return vHzmUaBZ[ /*line :1*/ uint(-d2X2v9)/32] / wd0f5H[ /*line :1*/ uint(-d2X2v9)%32]
	}

	if d2X2v9 > 0 {
		return  /*line :1*/ FSug4WHZSBwz(1)
	}

	return 0
}
