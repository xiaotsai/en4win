//line :1
package fuA83L


type DeOFa72M struct {
	uV4BJaCQ	uintptr
	qQUeiiUK	string
	czdHilCALsdh	error
}


func RLCP7BQDRyuR(bnJPAzeC uintptr) *DeOFa72M {
	return &DeOFa72M{uV4BJaCQ: bnJPAzeC}
}


func T8Ukv93kZU8s(bnJPAzeC uintptr, lnUjpq string) *DeOFa72M {
	return &DeOFa72M{uV4BJaCQ: bnJPAzeC, qQUeiiUK: lnUjpq}
}


func C9lofdLk(bnJPAzeC uintptr, lnUjpq string, zsgC7M1 error) *DeOFa72M {
	return &DeOFa72M{uV4BJaCQ: bnJPAzeC, qQUeiiUK: lnUjpq, czdHilCALsdh: zsgC7M1}
}


func (y0jiLdYHXNnx *DeOFa72M) Code() uintptr {
	return  /*line :1*/ uintptr(y0jiLdYHXNnx.uV4BJaCQ)
}


func (y0jiLdYHXNnx *DeOFa72M) String() string {
	if y0jiLdYHXNnx.qQUeiiUK != "" {
		return  /*line :1*/ jphUA_3( /*line :1*/ int(y0jiLdYHXNnx.uV4BJaCQ)) + " (" + y0jiLdYHXNnx.qQUeiiUK + ")"
	}
	return  /*line :1*/ jphUA_3( /*line :1*/ int(y0jiLdYHXNnx.uV4BJaCQ))
}


func (y0jiLdYHXNnx *DeOFa72M) Error() string {
	return  /*line :1*/ y0jiLdYHXNnx.String()
}


func (y0jiLdYHXNnx *DeOFa72M) Description() string {
	return y0jiLdYHXNnx.qQUeiiUK
}


func (y0jiLdYHXNnx *DeOFa72M) SubError() error {
	return y0jiLdYHXNnx.czdHilCALsdh
}
