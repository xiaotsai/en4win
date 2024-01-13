//line :1
package gzHaha55n


















































func Ir4aKca98X67(bntF0l int, vy6khO9C func(int) bool) int {

	eLewhDqfw8, kCr_PWXwcND9 := 0, bntF0l
	for eLewhDqfw8 < kCr_PWXwcND9 {
		xcoKpNDa :=  /*line :1*/ int( /*line :1*/ uint(eLewhDqfw8+kCr_PWXwcND9) >> 1)

		if ! /*line :1*/ vy6khO9C(xcoKpNDa) {
			eLewhDqfw8 = xcoKpNDa + 1
		} else {
			kCr_PWXwcND9 = xcoKpNDa
		}
	}

	return eLewhDqfw8
}

























func YZ23qcg(bntF0l int, oghMgC4AtxHS func(int) int) (eLewhDqfw8 int, g2UZjAkfvVn bool) {

	eLewhDqfw8, kCr_PWXwcND9 := 0, bntF0l
	for eLewhDqfw8 < kCr_PWXwcND9 {
		xcoKpNDa :=  /*line :1*/ int( /*line :1*/ uint(eLewhDqfw8+kCr_PWXwcND9) >> 1)

		if  /*line :1*/ oghMgC4AtxHS(xcoKpNDa) > 0 {
			eLewhDqfw8 = xcoKpNDa + 1
		} else {
			kCr_PWXwcND9 = xcoKpNDa
		}
	}

	return eLewhDqfw8, eLewhDqfw8 < bntF0l &&  /*line :1*/ oghMgC4AtxHS(eLewhDqfw8) == 0
}





func QqQ9zPNWai(bSi_aMm2pI []int, z1x13AfMF int) int {
	return  /*line :1*/ Ir4aKca98X67( /*line :1*/ len(bSi_aMm2pI), func(eLewhDqfw8 int) bool { return bSi_aMm2pI[eLewhDqfw8] >= z1x13AfMF })
}





func NcjMq20zMO(bSi_aMm2pI []float64, z1x13AfMF float64) int {
	return  /*line :1*/ Ir4aKca98X67( /*line :1*/ len(bSi_aMm2pI), func(eLewhDqfw8 int) bool { return bSi_aMm2pI[eLewhDqfw8] >= z1x13AfMF })
}





func KSrcRXVEhC(bSi_aMm2pI []string, z1x13AfMF string) int {
	return  /*line :1*/ Ir4aKca98X67( /*line :1*/ len(bSi_aMm2pI), func(eLewhDqfw8 int) bool { return bSi_aMm2pI[eLewhDqfw8] >= z1x13AfMF })
}


func (zViw4r NyonxyxVd) Search(z1x13AfMF int) int	{ return  /*line :1*/ QqQ9zPNWai(zViw4r, z1x13AfMF) }


func (zViw4r JTBP7y3_wY) Search(z1x13AfMF float64) int	{ return  /*line :1*/ NcjMq20zMO(zViw4r, z1x13AfMF) }


func (zViw4r ErggTD5A) Search(z1x13AfMF string) int	{ return  /*line :1*/ KSrcRXVEhC(zViw4r, z1x13AfMF) }
