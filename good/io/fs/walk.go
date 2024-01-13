//line :1
package y1BamVjnXsaa

import (
	errors "iAHoxjmM"
	path "jn9Psh1aa_7H"
)




var X7Z1FBI33u =  /*line :1*/ errors.Su6g6hRoi9X("skip this directory")




var MTx5MBbQt2 =  /*line :1*/ errors.Su6g6hRoi9X("skip everything and stop the walk")
















































type S_ytKQXs func(qkr_B7ssRGvO string, kwKOrGt RbfBMxa, glldLezIW error) error


func hiJvumCb(viJma5vOi XIcFcDgy, hMWDrBfy string, kwKOrGt RbfBMxa, dZAYid S_ytKQXs) error {
	if glldLezIW :=  /*line :1*/ dZAYid(hMWDrBfy, kwKOrGt, nil); glldLezIW != nil || ! /*line :1*/ kwKOrGt.IsDir() {
		if glldLezIW == X7Z1FBI33u &&  /*line :1*/ kwKOrGt.IsDir() {

			glldLezIW = nil
		}
		return glldLezIW
	}

	pdUqIstKqh, glldLezIW :=  /*line :1*/ VOXtzDEapl(viJma5vOi, hMWDrBfy)
	if glldLezIW != nil {

		glldLezIW =  /*line :1*/ dZAYid(hMWDrBfy, kwKOrGt, glldLezIW)
		if glldLezIW != nil {
			if glldLezIW == X7Z1FBI33u &&  /*line :1*/ kwKOrGt.IsDir() {
				glldLezIW = nil
			}
			return glldLezIW
		}
	}

	for _, safGx0YoRo := range pdUqIstKqh {
		daPnJBY762Zi :=  /*line :1*/ path.OT0k1D(hMWDrBfy,  /*line :1*/ safGx0YoRo.Name())
		if glldLezIW :=  /*line :1*/ hiJvumCb(viJma5vOi, daPnJBY762Zi, safGx0YoRo, dZAYid); glldLezIW != nil {
			if glldLezIW == X7Z1FBI33u {
				break
			}
			return glldLezIW
		}
	}
	return nil
}













func LYahiVitVNsw(viJma5vOi XIcFcDgy, bDfbKCQyna3 string, vTgQjJh S_ytKQXs) error {
	v9FHH9DB9sM2, glldLezIW :=  /*line :1*/ PTHkijCMP7c(viJma5vOi, bDfbKCQyna3)
	if glldLezIW != nil {
		glldLezIW =  /*line :1*/ vTgQjJh(bDfbKCQyna3, nil, glldLezIW)
	} else {
		glldLezIW =  /*line :1*/ hiJvumCb(viJma5vOi, bDfbKCQyna3, &syvoDJ{v9FHH9DB9sM2}, vTgQjJh)
	}
	if glldLezIW == X7Z1FBI33u || glldLezIW == MTx5MBbQt2 {
		return nil
	}
	return glldLezIW
}

type syvoDJ struct {
	aDXlH4zQP HM4JM2
}

func (kwKOrGt *syvoDJ) Name() string	{ return  /*line :1*/ kwKOrGt.aDXlH4zQP.Name() }
func (kwKOrGt *syvoDJ) IsDir() bool	{ return  /*line :1*/ kwKOrGt.aDXlH4zQP.IsDir() }
func (kwKOrGt *syvoDJ) Type() PGXMbJjlHBMu	{ return  /*line :1*/ kwKOrGt.aDXlH4zQP.Mode().Type() }
func (kwKOrGt *syvoDJ) Info() (HM4JM2, error)	{ return kwKOrGt.aDXlH4zQP, nil }

func (kwKOrGt *syvoDJ) String() string {
	return  /*line :1*/ EMHCka(kwKOrGt)
}
