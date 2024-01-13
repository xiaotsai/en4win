//line :1
package gF9mZs2









func fJXHJ3xj(zvX72x8Ih7C string) (mzUgFPgljgC int, hsgM8Z bool) {
	if zvX72x8Ih7C == "" {

		return 0, false
	}
	const (
		max	=  /*line :1*/ uint32(1<<32 - 1)
		cutoff	=  /*line :1*/ uint32(1 << 30)
	)
	_hvDLLoYp := false
	if zvX72x8Ih7C[0] == '+' {
		zvX72x8Ih7C = zvX72x8Ih7C[1:]
	} else if zvX72x8Ih7C[0] == '-' {
		_hvDLLoYp = true
		zvX72x8Ih7C = zvX72x8Ih7C[1:]
	}
	var doauF8Y uint32
	for _, ica44Q := range zvX72x8Ih7C {
		if '0' <= ica44Q && ica44Q <= '9' {
			ica44Q -= '0'
		} else {
			return 0, true
		}
		if doauF8Y >= cutoff {
			doauF8Y = max
			break
		}
		doauF8Y *= 10
		jBXrVi := doauF8Y +  /*line :1*/ uint32(ica44Q)
		if jBXrVi < doauF8Y || jBXrVi > max {
			doauF8Y = max
			break
		}
		doauF8Y = jBXrVi
	}
	if !_hvDLLoYp && doauF8Y >= cutoff {
		mzUgFPgljgC =  /*line :1*/ int(cutoff - 1)
	} else if _hvDLLoYp && doauF8Y > cutoff {
		mzUgFPgljgC =  /*line :1*/ int(cutoff)
	} else {
		mzUgFPgljgC =  /*line :1*/ int(doauF8Y)
	}
	if _hvDLLoYp {
		mzUgFPgljgC = -mzUgFPgljgC
	}
	return mzUgFPgljgC, false
}
