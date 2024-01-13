//line :1
package iAHoxjmM

import (
	reflectlite "yAj8QghzkR"
)







func DA42xXO_w2J(eLreS4RT1 error) error {
	d9aiCTKN2, r4xC3_N_S := eLreS4RT1.(interface {
		Unwrap() error
	})
	if !r4xC3_N_S {
		return nil
	}
	return  /*line :1*/ d9aiCTKN2.Unwrap()
}


















func COBastO_C(eLreS4RT1, j4deOEfCY2 error) bool {
	if j4deOEfCY2 == nil {
		return eLreS4RT1 == j4deOEfCY2
	}

	h2JoJ08 :=  /*line :1*/ reflectlite.C08auGx(j4deOEfCY2).Comparable()
	for {
		if h2JoJ08 && eLreS4RT1 == j4deOEfCY2 {
			return true
		}
		if h8r14RS2vjf, r4xC3_N_S := eLreS4RT1.(interface{ Is(error) bool }); r4xC3_N_S &&  /*line :1*/ h8r14RS2vjf.Is(j4deOEfCY2) {
			return true
		}
		switch h8r14RS2vjf := eLreS4RT1.(type) {
		case interface{ Unwrap() error }:
			eLreS4RT1 =  /*line :1*/ h8r14RS2vjf.Unwrap()
			if eLreS4RT1 == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, eLreS4RT1 := range  /*line :1*/ h8r14RS2vjf.Unwrap() {
				if  /*line :1*/ COBastO_C(eLreS4RT1, j4deOEfCY2) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}


















func CyKhHznARJK(eLreS4RT1 error, j4deOEfCY2 any) bool {
	if eLreS4RT1 == nil {
		return false
	}
	if j4deOEfCY2 == nil {
		 /*line :1*/ panic("errors: target cannot be nil")
	}
	nGh2n5JT :=  /*line :1*/ reflectlite.Au73KsHV1lt(j4deOEfCY2)
	v3J5yA5o1mAA :=  /*line :1*/ nGh2n5JT.Type()
	if  /*line :1*/ v3J5yA5o1mAA.Kind() != reflectlite.Ptr ||  /*line :1*/ nGh2n5JT.IsNil() {
		 /*line :1*/ panic("errors: target must be a non-nil pointer")
	}
	a5HQKKGif :=  /*line :1*/ v3J5yA5o1mAA.Elem()
	if  /*line :1*/ a5HQKKGif.Kind() != reflectlite.Interface && ! /*line :1*/ a5HQKKGif.Implements(iscRhTrEM) {
		 /*line :1*/ panic("errors: *target must be interface or implement error")
	}
	for {
		if  /*line :1*/ reflectlite.C08auGx(eLreS4RT1).AssignableTo(a5HQKKGif) {
			 /*line :1*/ nGh2n5JT.Elem().Set( /*line :1*/ reflectlite.Au73KsHV1lt(eLreS4RT1))
			return true
		}
		if h8r14RS2vjf, r4xC3_N_S := eLreS4RT1.(interface{ As(any) bool }); r4xC3_N_S &&  /*line :1*/ h8r14RS2vjf.As(j4deOEfCY2) {
			return true
		}
		switch h8r14RS2vjf := eLreS4RT1.(type) {
		case interface{ Unwrap() error }:
			eLreS4RT1 =  /*line :1*/ h8r14RS2vjf.Unwrap()
			if eLreS4RT1 == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, eLreS4RT1 := range  /*line :1*/ h8r14RS2vjf.Unwrap() {
				if  /*line :1*/ CyKhHznARJK(eLreS4RT1, j4deOEfCY2) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}

var iscRhTrEM =  /*line :1*/ reflectlite.C08auGx((* /*line :1*/ error)(nil)).Elem()
