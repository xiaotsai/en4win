//line :1
package iAHoxjmM









func UhcQJdiPwaVd(gqpO7PAdT5Z ...error) error {
	zgSEI6oPGS := 0
	for _, eLreS4RT1 := range gqpO7PAdT5Z {
		if eLreS4RT1 != nil {
			zgSEI6oPGS++
		}
	}
	if zgSEI6oPGS == 0 {
		return nil
	}
	_Ssmbp9SQ := &gfMpxSCA5VD{
		bCVZD2uGF:  /*line :1*/ make([]error, 0, zgSEI6oPGS),
	}
	for _, eLreS4RT1 := range gqpO7PAdT5Z {
		if eLreS4RT1 != nil {
			_Ssmbp9SQ.bCVZD2uGF =  /*line :1*/ append(_Ssmbp9SQ.bCVZD2uGF, eLreS4RT1)
		}
	}
	return _Ssmbp9SQ
}

type gfMpxSCA5VD struct {
	bCVZD2uGF []error
}

func (_Ssmbp9SQ *gfMpxSCA5VD) Error() string {
	var eQhq1VAw_CL []byte
	for gSmeRR4, eLreS4RT1 := range _Ssmbp9SQ.bCVZD2uGF {
		if gSmeRR4 > 0 {
			eQhq1VAw_CL =  /*line :1*/ append(eQhq1VAw_CL, '\n')
		}
		eQhq1VAw_CL =  /*line :1*/ append(eQhq1VAw_CL,  /*line :1*/ eLreS4RT1.Error()...)
	}
	return  /*line :1*/ string(eQhq1VAw_CL)
}

func (_Ssmbp9SQ *gfMpxSCA5VD) Unwrap() []error {
	return _Ssmbp9SQ.bCVZD2uGF
}
