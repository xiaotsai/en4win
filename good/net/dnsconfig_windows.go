//line :1
package gF9mZs2

import (
	windows "LdptURlN"
	syscall "bUKeamF"
	time "fRAfQd_"
)

func nP4L7Rk4(gLYBnO7xt string) (q6Glr4X1u *ivgt6m77Buqr) {
	q6Glr4X1u = &ivgt6m77Buqr{
		kzjIfgwIhJ_p:	1,
		c1XqMJ:	5 * time.Second,
		yb6dJ6U:	2,
	}
	defer func() {
		if  /*line :1*/ len(q6Glr4X1u.aeSIvF7Fa) == 0 {
			q6Glr4X1u.aeSIvF7Fa = m6PAg4e1Esq
		}
	}()
	upqMp71No, h_ljk48Bm :=  /*line :1*/ aQhDgDOvXgC()
	if h_ljk48Bm != nil {
		return
	}

	for _, oJ5XGLy := range upqMp71No {
		for jwpNdr4DIdL := oJ5XGLy.ALEYAV_; jwpNdr4DIdL != nil; jwpNdr4DIdL = jwpNdr4DIdL.HyuwAR7G {

			if oJ5XGLy.AOg4mu != windows.IfOperStatusUp {
				continue
			}
			epTaNF0RJ8eN, h_ljk48Bm :=  /*line :1*/ jwpNdr4DIdL.JuLy4dFjAt.NgzCg7T0_8yT.Sockaddr()
			if h_ljk48Bm != nil {
				continue
			}
			var kQ8_UEhxU GNraIvYhB
			switch epTaNF0RJ8eN := epTaNF0RJ8eN.(type) {
			case *syscall.Hx8lqxSkV:
				kQ8_UEhxU =  /*line :1*/ Ip1NFCL6(epTaNF0RJ8eN.Q3zz2uH[0], epTaNF0RJ8eN.Q3zz2uH[1], epTaNF0RJ8eN.Q3zz2uH[2], epTaNF0RJ8eN.Q3zz2uH[3])
			case *syscall.HKTAy7_BSU2:
				kQ8_UEhxU =  /*line :1*/ make(GNraIvYhB, IPv6len)
				 /*line :1*/ copy(kQ8_UEhxU, epTaNF0RJ8eN.Y4XxFr[:])
				if kQ8_UEhxU[0] == 0xfe && kQ8_UEhxU[1] == 0xc0 {

					continue
				}
			default:

				continue
			}
			q6Glr4X1u.aeSIvF7Fa =  /*line :1*/ append(q6Glr4X1u.aeSIvF7Fa,  /*line :1*/ Vd5lcUxG( /*line :1*/ kQ8_UEhxU.String(), "53"))
		}
	}
	return q6Glr4X1u
}
