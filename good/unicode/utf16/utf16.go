//line :1

package DtJCLKevRp

const (
	replacementChar	= '\uFFFD'		
	maxRune	= '\U0010FFFF'		
)

const (
	
	
	
	surr1	= 0xd800
	surr2	= 0xdc00
	surr3	= 0xe000

	surrSelf	= 0x10000
)



func I0RIjMaMPDrI(_JuN_FOqyc rune) bool {
	return surr1 <= _JuN_FOqyc && _JuN_FOqyc < surr3
}




func Rw3auy7U(nb8lNj_, dRRkXvF_MUf_ rune) rune {
	if surr1 <= nb8lNj_ && nb8lNj_ < surr2 && surr2 <= dRRkXvF_MUf_ && dRRkXvF_MUf_ < surr3 {
		return (nb8lNj_-surr1)<<10 | (dRRkXvF_MUf_ - surr2) + surrSelf
	}
	return replacementChar
}




func W3bDnUfML(_JuN_FOqyc rune) (nb8lNj_, dRRkXvF_MUf_ rune) {
	if _JuN_FOqyc < surrSelf || _JuN_FOqyc > maxRune {
		return replacementChar, replacementChar
	}
	_JuN_FOqyc -= surrSelf
	return surr1 + (_JuN_FOqyc>>10)&0x3ff, surr2 + _JuN_FOqyc&0x3ff
}


func J8eAwOfFc(fW6grW6YP []rune) []uint16 {
	inKPBzNb3S :=  /*line :1*/ len(fW6grW6YP)
	for _, a3FaaaMsj := range fW6grW6YP {
		if a3FaaaMsj >= surrSelf {
			inKPBzNb3S++
		}
	}

	jg6tSJJdkuL :=  /*line :1*/ make([]uint16, inKPBzNb3S)
	inKPBzNb3S = 0
	for _, a3FaaaMsj := range fW6grW6YP {
		switch {
		case 0 <= a3FaaaMsj && a3FaaaMsj < surr1, surr3 <= a3FaaaMsj && a3FaaaMsj < surrSelf:

			jg6tSJJdkuL[inKPBzNb3S] =  /*line :1*/ uint16(a3FaaaMsj)
			inKPBzNb3S++
		case surrSelf <= a3FaaaMsj && a3FaaaMsj <= maxRune:

			nb8lNj_, dRRkXvF_MUf_ :=  /*line :1*/ W3bDnUfML(a3FaaaMsj)
			jg6tSJJdkuL[inKPBzNb3S] =  /*line :1*/ uint16(nb8lNj_)
			jg6tSJJdkuL[inKPBzNb3S+1] =  /*line :1*/ uint16(dRRkXvF_MUf_)
			inKPBzNb3S += 2
		default:
			jg6tSJJdkuL[inKPBzNb3S] =  /*line :1*/ uint16(replacementChar)
			inKPBzNb3S++
		}
	}
	return jg6tSJJdkuL[:inKPBzNb3S]
}




func J23fS69TUl7H(jg6tSJJdkuL []uint16, _JuN_FOqyc rune) []uint16 {

	switch {
	case 0 <= _JuN_FOqyc && _JuN_FOqyc < surr1, surr3 <= _JuN_FOqyc && _JuN_FOqyc < surrSelf:

		return  /*line :1*/ append(jg6tSJJdkuL,  /*line :1*/ uint16(_JuN_FOqyc))
	case surrSelf <= _JuN_FOqyc && _JuN_FOqyc <= maxRune:

		nb8lNj_, dRRkXvF_MUf_ :=  /*line :1*/ W3bDnUfML(_JuN_FOqyc)
		return  /*line :1*/ append(jg6tSJJdkuL,  /*line :1*/ uint16(nb8lNj_),  /*line :1*/ uint16(dRRkXvF_MUf_))
	}
	return  /*line :1*/ append(jg6tSJJdkuL, replacementChar)
}



func Q5WZm8(fW6grW6YP []uint16) []rune {

	a5nWnV :=  /*line :1*/ make([]rune, 0, 64)
	return  /*line :1*/ vZNLrmp5(fW6grW6YP, a5nWnV)
}



func vZNLrmp5(fW6grW6YP []uint16, a5nWnV []rune) []rune {
	for faw2Z8Af_ := 0; faw2Z8Af_ <  /*line :1*/ len(fW6grW6YP); faw2Z8Af_++ {
		var tPmEsU6uoi rune
		switch _JuN_FOqyc := fW6grW6YP[faw2Z8Af_]; {
		case _JuN_FOqyc < surr1, surr3 <= _JuN_FOqyc:

			tPmEsU6uoi =  /*line :1*/ rune(_JuN_FOqyc)
		case surr1 <= _JuN_FOqyc && _JuN_FOqyc < surr2 && faw2Z8Af_+1 <  /*line :1*/ len(fW6grW6YP) &&
			surr2 <= fW6grW6YP[faw2Z8Af_+1] && fW6grW6YP[faw2Z8Af_+1] < surr3:

			tPmEsU6uoi =  /*line :1*/ Rw3auy7U( /*line :1*/ rune(_JuN_FOqyc),  /*line :1*/ rune(fW6grW6YP[faw2Z8Af_+1]))
			faw2Z8Af_++
		default:

			tPmEsU6uoi = replacementChar
		}
		a5nWnV =  /*line :1*/ append(a5nWnV, tPmEsU6uoi)
	}
	return a5nWnV
}
