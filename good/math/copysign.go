//line :1
package math



func Copysign(ygoarCBO4et, awRlyItEG2gn float64) float64 {
	const signBit = 1 << 63
	return  /*line :1*/ QUB2Kf63p( /*line :1*/ GKIRmoHE(ygoarCBO4et)&^signBit |  /*line :1*/ GKIRmoHE(awRlyItEG2gn)&signBit)
}
