//line :1







package jn9Psh1aa_7H





type h6V415HcqrU struct {
	pFeacNo	string
	jP0aqfIE1I	[]byte
	qvBa6Mq	int
}

func (uEOBQaqwP30 *h6V415HcqrU) qoZifp4(tr99J_O int) byte {
	if uEOBQaqwP30.jP0aqfIE1I != nil {
		return uEOBQaqwP30.jP0aqfIE1I[tr99J_O]
	}
	return uEOBQaqwP30.pFeacNo[tr99J_O]
}

func (uEOBQaqwP30 *h6V415HcqrU) dSNAzaW740(abn2bX0dLa byte) {
	if uEOBQaqwP30.jP0aqfIE1I == nil {
		if uEOBQaqwP30.qvBa6Mq <  /*line :1*/ len(uEOBQaqwP30.pFeacNo) && uEOBQaqwP30.pFeacNo[uEOBQaqwP30.qvBa6Mq] == abn2bX0dLa {
			uEOBQaqwP30.qvBa6Mq++
			return
		}
		uEOBQaqwP30.jP0aqfIE1I =  /*line :1*/ make([]byte,  /*line :1*/ len(uEOBQaqwP30.pFeacNo))
		 /*line :1*/ copy(uEOBQaqwP30.jP0aqfIE1I, uEOBQaqwP30.pFeacNo[:uEOBQaqwP30.qvBa6Mq])
	}
	uEOBQaqwP30.jP0aqfIE1I[uEOBQaqwP30.qvBa6Mq] = abn2bX0dLa
	uEOBQaqwP30.qvBa6Mq++
}

func (uEOBQaqwP30 *h6V415HcqrU) h9WErCbWKR() string {
	if uEOBQaqwP30.jP0aqfIE1I == nil {
		return uEOBQaqwP30.pFeacNo[:uEOBQaqwP30.qvBa6Mq]
	}
	return  /*line :1*/ string(uEOBQaqwP30.jP0aqfIE1I[:uEOBQaqwP30.qvBa6Mq])
}




















func GJEYGL(jn9Psh1aa_7H string) string {
	if jn9Psh1aa_7H == "" {
		return "."
	}

	sJvfOk := jn9Psh1aa_7H[0] == '/'
	sfulV7P7QO44 :=  /*line :1*/ len(jn9Psh1aa_7H)

	nuKt1RV0Pn := h6V415HcqrU{pFeacNo: jn9Psh1aa_7H}
	aC96awU, pq4w8UlM := 0, 0
	if sJvfOk {
		 /*line :1*/ nuKt1RV0Pn.dSNAzaW740('/')
		aC96awU, pq4w8UlM = 1, 1
	}

	for aC96awU < sfulV7P7QO44 {
		switch {
		case jn9Psh1aa_7H[aC96awU] == '/':

			aC96awU++
		case jn9Psh1aa_7H[aC96awU] == '.' && (aC96awU+1 == sfulV7P7QO44 || jn9Psh1aa_7H[aC96awU+1] == '/'):

			aC96awU++
		case jn9Psh1aa_7H[aC96awU] == '.' && jn9Psh1aa_7H[aC96awU+1] == '.' && (aC96awU+2 == sfulV7P7QO44 || jn9Psh1aa_7H[aC96awU+2] == '/'):

			aC96awU += 2
			switch {
			case nuKt1RV0Pn.qvBa6Mq > pq4w8UlM:

				nuKt1RV0Pn.qvBa6Mq--
				for nuKt1RV0Pn.qvBa6Mq > pq4w8UlM &&  /*line :1*/ nuKt1RV0Pn.qoZifp4(nuKt1RV0Pn.qvBa6Mq) != '/' {
					nuKt1RV0Pn.qvBa6Mq--
				}
			case !sJvfOk:

				if nuKt1RV0Pn.qvBa6Mq > 0 {
					 /*line :1*/ nuKt1RV0Pn.dSNAzaW740('/')
				}
				 /*line :1*/ nuKt1RV0Pn.dSNAzaW740('.')
				 /*line :1*/ nuKt1RV0Pn.dSNAzaW740('.')
				pq4w8UlM = nuKt1RV0Pn.qvBa6Mq
			}
		default:

			if sJvfOk && nuKt1RV0Pn.qvBa6Mq != 1 || !sJvfOk && nuKt1RV0Pn.qvBa6Mq != 0 {
				 /*line :1*/ nuKt1RV0Pn.dSNAzaW740('/')
			}

			for ; aC96awU < sfulV7P7QO44 && jn9Psh1aa_7H[aC96awU] != '/'; aC96awU++ {
				 /*line :1*/ nuKt1RV0Pn.dSNAzaW740(jn9Psh1aa_7H[aC96awU])
			}
		}
	}

	if nuKt1RV0Pn.qvBa6Mq == 0 {
		return "."
	}

	return  /*line :1*/ nuKt1RV0Pn.h9WErCbWKR()
}


func oayieBoq6D(kSAda01ZGvqp string) int {
	tr99J_O :=  /*line :1*/ len(kSAda01ZGvqp) - 1
	for tr99J_O >= 0 && kSAda01ZGvqp[tr99J_O] != '/' {
		tr99J_O--
	}
	return tr99J_O
}






func KHYEhOz2l(jn9Psh1aa_7H string) (jnMHcMEuU, fU2iCXTMXmz string) {
	tr99J_O :=  /*line :1*/ oayieBoq6D(jn9Psh1aa_7H)
	return jn9Psh1aa_7H[:tr99J_O+1], jn9Psh1aa_7H[tr99J_O+1:]
}






func OT0k1D(yspdya8au ...string) string {
	nCMuwm5tucSj := 0
	for _, cgw0BZhjwk := range yspdya8au {
		nCMuwm5tucSj +=  /*line :1*/ len(cgw0BZhjwk)
	}
	if nCMuwm5tucSj == 0 {
		return ""
	}
	jOHuZ7nRsN :=  /*line :1*/ make([]byte, 0, nCMuwm5tucSj+ /*line :1*/ len(yspdya8au)-1)
	for _, cgw0BZhjwk := range yspdya8au {
		if  /*line :1*/ len(jOHuZ7nRsN) > 0 || cgw0BZhjwk != "" {
			if  /*line :1*/ len(jOHuZ7nRsN) > 0 {
				jOHuZ7nRsN =  /*line :1*/ append(jOHuZ7nRsN, '/')
			}
			jOHuZ7nRsN =  /*line :1*/ append(jOHuZ7nRsN, cgw0BZhjwk...)
		}
	}
	return  /*line :1*/ GJEYGL( /*line :1*/ string(jOHuZ7nRsN))
}





func SNrICYeDL1(jn9Psh1aa_7H string) string {
	for tr99J_O :=  /*line :1*/ len(jn9Psh1aa_7H) - 1; tr99J_O >= 0 && jn9Psh1aa_7H[tr99J_O] != '/'; tr99J_O-- {
		if jn9Psh1aa_7H[tr99J_O] == '.' {
			return jn9Psh1aa_7H[tr99J_O:]
		}
	}
	return ""
}





func B1IWzu(jn9Psh1aa_7H string) string {
	if jn9Psh1aa_7H == "" {
		return "."
	}

	for  /*line :1*/ len(jn9Psh1aa_7H) > 0 && jn9Psh1aa_7H[ /*line :1*/ len(jn9Psh1aa_7H)-1] == '/' {
		jn9Psh1aa_7H = jn9Psh1aa_7H[0 :  /*line :1*/ len(jn9Psh1aa_7H)-1]
	}

	if tr99J_O :=  /*line :1*/ oayieBoq6D(jn9Psh1aa_7H); tr99J_O >= 0 {
		jn9Psh1aa_7H = jn9Psh1aa_7H[tr99J_O+1:]
	}

	if jn9Psh1aa_7H == "" {
		return "/"
	}
	return jn9Psh1aa_7H
}


func H1iVbLxMDY(jn9Psh1aa_7H string) bool {
	return  /*line :1*/ len(jn9Psh1aa_7H) > 0 && jn9Psh1aa_7H[0] == '/'
}








func Mw_x55(jn9Psh1aa_7H string) string {
	jnMHcMEuU, _ :=  /*line :1*/ KHYEhOz2l(jn9Psh1aa_7H)
	return  /*line :1*/ GJEYGL(jnMHcMEuU)
}
