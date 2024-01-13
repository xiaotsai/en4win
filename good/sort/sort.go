//line :1
//go:generate go run gen_sort_variants.go

package gzHaha55n

import bits "math/bits"

type AOVoo3y622nl interface {
	Len() int

	Less(eLewhDqfw8, kCr_PWXwcND9 int) bool

	Swap(eLewhDqfw8, kCr_PWXwcND9 int)
}

func KX2G4Qz(cxvZaD16 AOVoo3y622nl) {
	bntF0l :=  /*line :1*/ cxvZaD16.Len()
	if bntF0l <= 1 {
		return
	}
	l5vzUBwzefa :=  /*line :1*/ bits.Len( /*line :1*/ uint(bntF0l))
	 /*line :1*/ mUdU8k(cxvZaD16, 0, bntF0l, l5vzUBwzefa)
}

type ea2AAb0t int

const (
	unknownHint	ea2AAb0t	= iota
	increasingHint
	decreasingHint
)

type aMwO4FPaJS3s uint64

func (n6aYLO8 *aMwO4FPaJS3s) Next() uint64 {
	*n6aYLO8 ^= *n6aYLO8 << 13
	*n6aYLO8 ^= *n6aYLO8 >> 17
	*n6aYLO8 ^= *n6aYLO8 << 5
	return  /*line :1*/ uint64(*n6aYLO8)
}

func boBGv1E(kbaqw6D int) uint {
	nXGRBo :=  /*line :1*/ uint( /*line :1*/ bits.Len( /*line :1*/ uint(kbaqw6D)))
	return  /*line :1*/ uint(1 << nXGRBo)
}

type adLnW8ivMz struct {
	Laeasg6am	func(eLewhDqfw8, kCr_PWXwcND9 int) bool
	DCF_4kNh	func(eLewhDqfw8, kCr_PWXwcND9 int)
}

type cS5Qat struct {
	AOVoo3y622nl
}

func (n6aYLO8 cS5Qat) Less(eLewhDqfw8, kCr_PWXwcND9 int) bool {
	return  /*line :1*/ n6aYLO8.AOVoo3y622nl.Less(kCr_PWXwcND9, eLewhDqfw8)
}

func YjNFW0a(cxvZaD16 AOVoo3y622nl) AOVoo3y622nl {
	return &cS5Qat{cxvZaD16}
}

func Bw8x3jTK09(cxvZaD16 AOVoo3y622nl) bool {
	bntF0l :=  /*line :1*/ cxvZaD16.Len()
	for eLewhDqfw8 := bntF0l - 1; eLewhDqfw8 > 0; eLewhDqfw8-- {
		if  /*line :1*/ cxvZaD16.Less(eLewhDqfw8, eLewhDqfw8-1) {
			return false
		}
	}
	return true
}

type NyonxyxVd []int

func (z1x13AfMF NyonxyxVd) Len() int	{ return  /*line :1*/ len(z1x13AfMF) }
func (z1x13AfMF NyonxyxVd) Less(eLewhDqfw8, kCr_PWXwcND9 int) bool {
	return z1x13AfMF[eLewhDqfw8] < z1x13AfMF[kCr_PWXwcND9]
}
func (z1x13AfMF NyonxyxVd) Swap(eLewhDqfw8, kCr_PWXwcND9 int) {
	z1x13AfMF[eLewhDqfw8], z1x13AfMF[kCr_PWXwcND9] = z1x13AfMF[kCr_PWXwcND9], z1x13AfMF[eLewhDqfw8]
}

func (z1x13AfMF NyonxyxVd) Sort()	{  /*line :1*/ KX2G4Qz(z1x13AfMF) }

type JTBP7y3_wY []float64

func (z1x13AfMF JTBP7y3_wY) Len() int	{ return  /*line :1*/ len(z1x13AfMF) }

func (z1x13AfMF JTBP7y3_wY) Less(eLewhDqfw8, kCr_PWXwcND9 int) bool {
	return z1x13AfMF[eLewhDqfw8] < z1x13AfMF[kCr_PWXwcND9] || ( /*line :1*/ gavwZMa(z1x13AfMF[eLewhDqfw8]) && ! /*line :1*/ gavwZMa(z1x13AfMF[kCr_PWXwcND9]))
}
func (z1x13AfMF JTBP7y3_wY) Swap(eLewhDqfw8, kCr_PWXwcND9 int) {
	z1x13AfMF[eLewhDqfw8], z1x13AfMF[kCr_PWXwcND9] = z1x13AfMF[kCr_PWXwcND9], z1x13AfMF[eLewhDqfw8]
}

func gavwZMa(vy6khO9C float64) bool {
	return vy6khO9C != vy6khO9C
}

func (z1x13AfMF JTBP7y3_wY) Sort()	{  /*line :1*/ KX2G4Qz(z1x13AfMF) }

type ErggTD5A []string

func (z1x13AfMF ErggTD5A) Len() int	{ return  /*line :1*/ len(z1x13AfMF) }
func (z1x13AfMF ErggTD5A) Less(eLewhDqfw8, kCr_PWXwcND9 int) bool {
	return z1x13AfMF[eLewhDqfw8] < z1x13AfMF[kCr_PWXwcND9]
}
func (z1x13AfMF ErggTD5A) Swap(eLewhDqfw8, kCr_PWXwcND9 int) {
	z1x13AfMF[eLewhDqfw8], z1x13AfMF[kCr_PWXwcND9] = z1x13AfMF[kCr_PWXwcND9], z1x13AfMF[eLewhDqfw8]
}

func (z1x13AfMF ErggTD5A) Sort()	{  /*line :1*/ KX2G4Qz(z1x13AfMF) }

func RZHzux(z1x13AfMF []int)	{  /*line :1*/ KX2G4Qz( /*line :1*/ NyonxyxVd(z1x13AfMF)) }

func E8t7LTB(z1x13AfMF []float64)	{  /*line :1*/ KX2G4Qz( /*line :1*/ JTBP7y3_wY(z1x13AfMF)) }

func IsJh7v(z1x13AfMF []string)	{  /*line :1*/ KX2G4Qz( /*line :1*/ ErggTD5A(z1x13AfMF)) }

func NOfzUwmQW(z1x13AfMF []int) bool	{ return  /*line :1*/ Bw8x3jTK09( /*line :1*/ NyonxyxVd(z1x13AfMF)) }

func YccOIgG7XuPt(z1x13AfMF []float64) bool	{ return  /*line :1*/ Bw8x3jTK09( /*line :1*/ JTBP7y3_wY(z1x13AfMF)) }

func Kftmb6zAI(z1x13AfMF []string) bool	{ return  /*line :1*/ Bw8x3jTK09( /*line :1*/ ErggTD5A(z1x13AfMF)) }

func IUVpc45h_Y3(cxvZaD16 AOVoo3y622nl) {
	 /*line :1*/ fYPxs1h(cxvZaD16,  /*line :1*/ cxvZaD16.Len())
}
