//line :1
package hPMrClpC

import (
	"runtime"
	sync "sync"
	syscall "bUKeamF"
)

var gqXK4qZubM struct {
	sync.DIRWe11KYlYa
	kLBsp7dvp	string
}





func OWhXb5XTeJpD() (nqzPkc8LPse1 string, ugqb2IW error) {
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		return  /*line :1*/ syscall.PQMcWL4xt()
	}

	b9M1paZ, ugqb2IW :=  /*line :1*/ r5FH06dvGEro(".")
	if ugqb2IW != nil {
		return "", ugqb2IW
	}
	nqzPkc8LPse1 =  /*line :1*/ LDjFPRBEz("PWD")
	if  /*line :1*/ len(nqzPkc8LPse1) > 0 && nqzPkc8LPse1[0] == '/' {
		b9ZD7F3, ugqb2IW :=  /*line :1*/ r5FH06dvGEro(nqzPkc8LPse1)
		if ugqb2IW == nil &&  /*line :1*/ ZUaDWEyatU(b9M1paZ, b9ZD7F3) {
			return nqzPkc8LPse1, nil
		}
	}

	if syscall.ImplementsGetwd {
		var (
			gkORch7na	string
			w50uhPri	error
		)
		for {
			gkORch7na, w50uhPri =  /*line :1*/ syscall.PQMcWL4xt()
			if w50uhPri != syscall.EINTR {
				break
			}
		}
		return gkORch7na,  /*line :1*/ BTaHHl("getwd", w50uhPri)
	}

	 /*line :1*/ gqXK4qZubM.Lock()
	nqzPkc8LPse1 = gqXK4qZubM.kLBsp7dvp
	 /*line :1*/ gqXK4qZubM.Unlock()
	if  /*line :1*/ len(nqzPkc8LPse1) > 0 {
		b9ZD7F3, ugqb2IW :=  /*line :1*/ r5FH06dvGEro(nqzPkc8LPse1)
		if ugqb2IW == nil &&  /*line :1*/ ZUaDWEyatU(b9M1paZ, b9ZD7F3) {
			return nqzPkc8LPse1, nil
		}
	}

	eay8nlJmLasg, ugqb2IW :=  /*line :1*/ r5FH06dvGEro("/")
	if ugqb2IW != nil {

		return "", ugqb2IW
	}
	if  /*line :1*/ ZUaDWEyatU(eay8nlJmLasg, b9M1paZ) {
		return "/", nil
	}

	nqzPkc8LPse1 = ""
	for srFQO7Adw2Qf := ".."; ; srFQO7Adw2Qf = "../" + srFQO7Adw2Qf {
		if  /*line :1*/ len(srFQO7Adw2Qf) >= 1024 {
			return "", syscall.ENAMETOOLONG
		}
		ja1FpjT, ugqb2IW :=  /*line :1*/ hczTpl7(srFQO7Adw2Qf, O_RDONLY, 0)
		if ugqb2IW != nil {
			return "", ugqb2IW
		}

		for {
			c8VnEL, ugqb2IW :=  /*line :1*/ ja1FpjT.Readdirnames(100)
			if ugqb2IW != nil {
				 /*line :1*/ ja1FpjT.Close()
				return "", ugqb2IW
			}
			for _, jGBs03 := range c8VnEL {
				b9ZD7F3, _ :=  /*line :1*/ x0lxe20b(srFQO7Adw2Qf + "/" + jGBs03)
				if  /*line :1*/ ZUaDWEyatU(b9ZD7F3, b9M1paZ) {
					nqzPkc8LPse1 = "/" + jGBs03 + nqzPkc8LPse1
					goto Found
				}
			}
		}

	Found:
		wGhx2SbxP, ugqb2IW :=  /*line :1*/ ja1FpjT.Stat()
		 /*line :1*/ ja1FpjT.Close()
		if ugqb2IW != nil {
			return "", ugqb2IW
		}
		if  /*line :1*/ ZUaDWEyatU(wGhx2SbxP, eay8nlJmLasg) {
			break
		}

		b9M1paZ = wGhx2SbxP
	}

	 /*line :1*/ gqXK4qZubM.Lock()
	gqXK4qZubM.kLBsp7dvp = nqzPkc8LPse1
	 /*line :1*/ gqXK4qZubM.Unlock()

	return nqzPkc8LPse1, nil
}
