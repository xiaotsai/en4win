//line :1
package math

const (
	uvnan	= 0x7FF8000000000001
	uvinf	= 0x7FF0000000000000
	uvneginf	= 0xFFF0000000000000
	uvone	= 0x3FF0000000000000
	mask	= 0x7FF
	shift	= 64 - 11 - 1
	bias	= 1023
	signMask	= 1 << 63
	fracMask	= 1<<shift - 1
)


func FSug4WHZSBwz(awRlyItEG2gn int) float64 {
	var hLZMcJsdS uint64
	if awRlyItEG2gn >= 0 {
		hLZMcJsdS = uvinf
	} else {
		hLZMcJsdS = uvneginf
	}
	return  /*line :1*/ QUB2Kf63p(hLZMcJsdS)
}


func WG0hZIT4() float64	{ return  /*line :1*/ QUB2Kf63p(uvnan) }


func OIdLpDqq(ygoarCBO4et float64) (uDDwHqw1CxX6 bool) {

	return ygoarCBO4et != ygoarCBO4et
}





func ME1vzpD5wcr(ygoarCBO4et float64, awRlyItEG2gn int) bool {

	return awRlyItEG2gn >= 0 && ygoarCBO4et > MaxFloat64 || awRlyItEG2gn <= 0 && ygoarCBO4et < -MaxFloat64
}



func ckXskOBp(_BqkSz0 float64) (nNeKvlvK6 float64, nF4u1Vnrit int) {
	const SmallestNormal = 2.2250738585072014e-308	
	if  /*line :1*/ Abs(_BqkSz0) < SmallestNormal {
		return _BqkSz0 * (1 << 52), -52
	}
	return _BqkSz0, 0
}
