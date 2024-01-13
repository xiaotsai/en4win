//line :1
package qyOdGuID2rmZ

import (
	errors "iAHoxjmM"
	fs "y1BamVjnXsaa"
	os "hPMrClpC"
	"runtime"
	syscall "bUKeamF"
)

func wDjPYoGaI72T(w9VHSyrN string) (string, error) {
	eDjfKln_q :=  /*line :1*/ sg1_tZpahJa(w9VHSyrN)
	zLm14w8GJnZr :=  /*line :1*/ string(os.PathSeparator)

	if eDjfKln_q <  /*line :1*/ len(w9VHSyrN) &&  /*line :1*/ os.GClWKECc(w9VHSyrN[eDjfKln_q]) {
		eDjfKln_q++
	}
	d2oIeyTClzU := w9VHSyrN[:eDjfKln_q]
	qmpeLyfkpkM := d2oIeyTClzU
	qrOKM_jaP := 0
	for whLL2V3vq, xeT2o8eThMaT := eDjfKln_q, eDjfKln_q; whLL2V3vq <  /*line :1*/ len(w9VHSyrN); whLL2V3vq = xeT2o8eThMaT {
		for whLL2V3vq <  /*line :1*/ len(w9VHSyrN) &&  /*line :1*/ os.GClWKECc(w9VHSyrN[whLL2V3vq]) {
			whLL2V3vq++
		}
		xeT2o8eThMaT = whLL2V3vq
		for xeT2o8eThMaT <  /*line :1*/ len(w9VHSyrN) && ! /*line :1*/ os.GClWKECc(w9VHSyrN[xeT2o8eThMaT]) {
			xeT2o8eThMaT++
		}

		mMOmPTYk7KB := runtime.GOOS == "windows" && w9VHSyrN[ /*line :1*/ sg1_tZpahJa(w9VHSyrN):] == "."

		if xeT2o8eThMaT == whLL2V3vq {

			break
		} else if w9VHSyrN[whLL2V3vq:xeT2o8eThMaT] == "." && !mMOmPTYk7KB {

			continue
		} else if w9VHSyrN[whLL2V3vq:xeT2o8eThMaT] == ".." {

			
			
			var qLSJu6fd4 int
			for qLSJu6fd4 =  /*line :1*/ len(qmpeLyfkpkM) - 1; qLSJu6fd4 >= eDjfKln_q; qLSJu6fd4-- {
				if  /*line :1*/ os.GClWKECc(qmpeLyfkpkM[qLSJu6fd4]) {
					break
				}
			}
			if qLSJu6fd4 < eDjfKln_q || qmpeLyfkpkM[qLSJu6fd4+1:] == ".." {

				if  /*line :1*/ len(qmpeLyfkpkM) > eDjfKln_q {
					qmpeLyfkpkM += zLm14w8GJnZr
				}
				qmpeLyfkpkM += ".."
			} else {

				qmpeLyfkpkM = qmpeLyfkpkM[:qLSJu6fd4]
			}
			continue
		}

		if  /*line :1*/ len(qmpeLyfkpkM) >  /*line :1*/ sg1_tZpahJa(qmpeLyfkpkM) && ! /*line :1*/ os.GClWKECc(qmpeLyfkpkM[ /*line :1*/ len(qmpeLyfkpkM)-1]) {
			qmpeLyfkpkM += zLm14w8GJnZr
		}

		qmpeLyfkpkM += w9VHSyrN[whLL2V3vq:xeT2o8eThMaT]

		tWHgrAGOK, aC4eBbavAAtF :=  /*line :1*/ os.Z0JXsY(qmpeLyfkpkM)
		if aC4eBbavAAtF != nil {
			return "", aC4eBbavAAtF
		}

		if  /*line :1*/ tWHgrAGOK.Mode()&fs.ModeSymlink == 0 {
			if ! /*line :1*/ tWHgrAGOK.Mode().IsDir() && xeT2o8eThMaT <  /*line :1*/ len(w9VHSyrN) {
				return "", syscall.ENOTDIR
			}
			continue
		}

		qrOKM_jaP++
		if qrOKM_jaP > 255 {
			return "",  /*line :1*/ errors.Su6g6hRoi9X("EvalSymlinks: too many links")
		}

		aKPWlqB, aC4eBbavAAtF :=  /*line :1*/ os.Jbg7I4xgch(qmpeLyfkpkM)
		if aC4eBbavAAtF != nil {
			return "", aC4eBbavAAtF
		}

		if mMOmPTYk7KB && ! /*line :1*/ Lr02fA2mZb(aKPWlqB) {

			break
		}

		w9VHSyrN = aKPWlqB + w9VHSyrN[xeT2o8eThMaT:]

		fdBBwo :=  /*line :1*/ sg1_tZpahJa(aKPWlqB)
		if fdBBwo > 0 {

			if fdBBwo <  /*line :1*/ len(aKPWlqB) &&  /*line :1*/ os.GClWKECc(aKPWlqB[fdBBwo]) {
				fdBBwo++
			}
			d2oIeyTClzU = aKPWlqB[:fdBBwo]
			qmpeLyfkpkM = d2oIeyTClzU
			xeT2o8eThMaT =  /*line :1*/ len(d2oIeyTClzU)
		} else if  /*line :1*/ len(aKPWlqB) > 0 &&  /*line :1*/ os.GClWKECc(aKPWlqB[0]) {

			qmpeLyfkpkM = aKPWlqB[:1]
			xeT2o8eThMaT = 1
			d2oIeyTClzU = aKPWlqB[:1]
			eDjfKln_q = 1
		} else {
			
			
			var qLSJu6fd4 int
			for qLSJu6fd4 =  /*line :1*/ len(qmpeLyfkpkM) - 1; qLSJu6fd4 >= eDjfKln_q; qLSJu6fd4-- {
				if  /*line :1*/ os.GClWKECc(qmpeLyfkpkM[qLSJu6fd4]) {
					break
				}
			}
			if qLSJu6fd4 < eDjfKln_q {
				qmpeLyfkpkM = d2oIeyTClzU
			} else {
				qmpeLyfkpkM = qmpeLyfkpkM[:qLSJu6fd4]
			}
			xeT2o8eThMaT = 0
		}
	}
	return  /*line :1*/ ZBiAXr(qmpeLyfkpkM), nil
}
