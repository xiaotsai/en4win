//line :1
package hPMrClpC

import (
	io "xI9ai2djaJ2"
	fs "y1BamVjnXsaa"
	"runtime"
	syscall "bUKeamF"
)

func (sR7PLgxb *BvGocYcXx) dQTx5qF2N(zHDjCTHI int, lVhwrRwT uppLHxeYfJc) (c8VnEL []string, ztdMGkvXuP []XXSR6RRqno8V, jR_i1_n []LWsU3x8MX, ugqb2IW error) {

	eafqvIah := true
	if sR7PLgxb.lkuk5KAfIH == nil {
		eafqvIah = false
		sR7PLgxb.lkuk5KAfIH, ugqb2IW =  /*line :1*/ gHMl7tDAEd(sR7PLgxb.b5ZMHrO5AlDN)
		if ugqb2IW != nil {
			ugqb2IW = &StNL1lveD40f{Aeg62j1VQt: "readdir", OmDEtY: sR7PLgxb.b5ZMHrO5AlDN, OIEifQ0XF: ugqb2IW}
			return
		}
	}
	aHAGp0sp := zHDjCTHI <= 0
	if aHAGp0sp {
		zHDjCTHI = -1
	}
	b9ZD7F3 := &sR7PLgxb.lkuk5KAfIH.m6lgHoPUtb
	for zHDjCTHI != 0 && !sR7PLgxb.lkuk5KAfIH.f6SvFo {
		if eafqvIah {
			w50uhPri :=  /*line :1*/ syscall.ZAQEj8rfk(sR7PLgxb.lkuk5KAfIH.f_xxI0, b9ZD7F3)
			 /*line :1*/ runtime.KeepAlive(sR7PLgxb)
			if w50uhPri != nil {
				if w50uhPri == syscall.ERROR_NO_MORE_FILES {
					break
				} else {
					ugqb2IW = &StNL1lveD40f{Aeg62j1VQt: "FindNextFile", OmDEtY: sR7PLgxb.b5ZMHrO5AlDN, OIEifQ0XF: w50uhPri}
					return
				}
			}
		}
		eafqvIah = true
		jGBs03 :=  /*line :1*/ syscall.AODVXp8NM3sd(b9ZD7F3.LMw5oQM3[0:])
		if jGBs03 == "." || jGBs03 == ".." {
			continue
		}
		if lVhwrRwT == readdirName {
			c8VnEL =  /*line :1*/ append(c8VnEL, jGBs03)
		} else {
			qzeVi328uMg :=  /*line :1*/ fDf_eyeL(b9ZD7F3)
			qzeVi328uMg.bs06oGoQ3TBi = jGBs03
			qzeVi328uMg.hlEnW1ut97N = sR7PLgxb.lkuk5KAfIH.ab2GI0U62omX
			qzeVi328uMg.dd_9IGm = true
			if lVhwrRwT == readdirDirEntry {
				ztdMGkvXuP =  /*line :1*/ append(ztdMGkvXuP, wphIIlXCLo_{qzeVi328uMg})
			} else {
				jR_i1_n =  /*line :1*/ append(jR_i1_n, qzeVi328uMg)
			}
		}
		zHDjCTHI--
	}
	if !aHAGp0sp &&  /*line :1*/ len(c8VnEL)+ /*line :1*/ len(ztdMGkvXuP)+ /*line :1*/ len(jR_i1_n) == 0 {
		return nil, nil, nil, io.K5Sqco
	}
	return c8VnEL, ztdMGkvXuP, jR_i1_n, nil
}

type wphIIlXCLo_ struct {
	pbJFSCq5jf *aqsPbX2HE5
}

func (jfNqR6J wphIIlXCLo_) Name() string	{ return  /*line :1*/ jfNqR6J.pbJFSCq5jf.Name() }
func (jfNqR6J wphIIlXCLo_) IsDir() bool	{ return  /*line :1*/ jfNqR6J.pbJFSCq5jf.IsDir() }
func (jfNqR6J wphIIlXCLo_) Type() IvL5u8pdn	{ return  /*line :1*/ jfNqR6J.pbJFSCq5jf.Mode().Type() }
func (jfNqR6J wphIIlXCLo_) Info() (LWsU3x8MX, error)	{ return jfNqR6J.pbJFSCq5jf, nil }

func (jfNqR6J wphIIlXCLo_) String() string {
	return  /*line :1*/ fs.EMHCka(jfNqR6J)
}
