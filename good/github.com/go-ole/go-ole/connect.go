//line :1
package fuA83L




type PGqAmS struct {
	MmaWW9 *IUnknown	
}


func (*PGqAmS) Initialize() (zsgC7M1 error) {
	return  /*line :1*/ m670LzvIC()
}


func (*PGqAmS) Uninitialize() {
	 /*line :1*/ DvOlryZ5()
}


func (eT1_dhKA *PGqAmS) Create(dvYDqdU string) (zsgC7M1 error) {
	var lcDXHYIAiC8j *GUID
	lcDXHYIAiC8j, zsgC7M1 =  /*line :1*/ AtGYsf(dvYDqdU)
	if zsgC7M1 != nil {
		lcDXHYIAiC8j, zsgC7M1 =  /*line :1*/ QFLzLQWmE(dvYDqdU)
		if zsgC7M1 != nil {
			return
		}
	}

	e9b4y2W, zsgC7M1 :=  /*line :1*/ Ne2od0bpE(lcDXHYIAiC8j, WySQlWR)
	if zsgC7M1 != nil {
		return
	}
	eT1_dhKA.MmaWW9 = e9b4y2W

	return
}


func (eT1_dhKA *PGqAmS) Release() {
	 /*line :1*/ eT1_dhKA.MmaWW9.Release()
}


func (eT1_dhKA *PGqAmS) Load(ljR636fK ...string) (ld3nx4XI []error) {
	var eU3jxit []error =  /*line :1*/ make([]error,  /*line :1*/ len(ljR636fK))
	var zdFG3sGkf int = 0
	for _, xwDC8DsSNM := range ljR636fK {
		zsgC7M1 :=  /*line :1*/ eT1_dhKA.Create(xwDC8DsSNM)
		if zsgC7M1 != nil {
			eU3jxit =  /*line :1*/ append(eU3jxit, zsgC7M1)
			zdFG3sGkf += 1
			continue
		}
		break
	}

	 /*line :1*/ copy(ld3nx4XI, eU3jxit[0:zdFG3sGkf])
	return
}


func (eT1_dhKA *PGqAmS) Dispatch() (r1xVxustMxt *HEMCaBV7x, zsgC7M1 error) {
	l6nPEB7Van, zsgC7M1 :=  /*line :1*/ eT1_dhKA.MmaWW9.QueryInterface(Y0sP6F)
	if zsgC7M1 != nil {
		return
	}
	r1xVxustMxt = &HEMCaBV7x{l6nPEB7Van}
	return
}


type HEMCaBV7x struct {
	DYHFxR *IDispatch	
}


func (b9Flo0 *HEMCaBV7x) Call(r8V4jqlunkmQ string, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D, zsgC7M1 error) {
	krXqCAVUYaTa, zsgC7M1 :=  /*line :1*/ b9Flo0.GetId(r8V4jqlunkmQ)
	if zsgC7M1 != nil {
		return
	}

	bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.Invoke(krXqCAVUYaTa, DISPATCH_METHOD, uTURuG8)
	return
}


func (b9Flo0 *HEMCaBV7x) MustCall(r8V4jqlunkmQ string, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D) {
	krXqCAVUYaTa, zsgC7M1 :=  /*line :1*/ b9Flo0.GetId(r8V4jqlunkmQ)
	if zsgC7M1 != nil {
		 /*line :1*/ panic(zsgC7M1)
	}

	bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.Invoke(krXqCAVUYaTa, DISPATCH_METHOD, uTURuG8)
	if zsgC7M1 != nil {
		 /*line :1*/ panic(zsgC7M1)
	}

	return
}


func (b9Flo0 *HEMCaBV7x) Get(xwDC8DsSNM string, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D, zsgC7M1 error) {
	krXqCAVUYaTa, zsgC7M1 :=  /*line :1*/ b9Flo0.GetId(xwDC8DsSNM)
	if zsgC7M1 != nil {
		return
	}
	bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.Invoke(krXqCAVUYaTa, DISPATCH_PROPERTYGET, uTURuG8)
	return
}


func (b9Flo0 *HEMCaBV7x) MustGet(xwDC8DsSNM string, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D) {
	krXqCAVUYaTa, zsgC7M1 :=  /*line :1*/ b9Flo0.GetId(xwDC8DsSNM)
	if zsgC7M1 != nil {
		 /*line :1*/ panic(zsgC7M1)
	}

	bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.Invoke(krXqCAVUYaTa, DISPATCH_PROPERTYGET, uTURuG8)
	if zsgC7M1 != nil {
		 /*line :1*/ panic(zsgC7M1)
	}
	return
}


func (b9Flo0 *HEMCaBV7x) Set(xwDC8DsSNM string, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D, zsgC7M1 error) {
	krXqCAVUYaTa, zsgC7M1 :=  /*line :1*/ b9Flo0.GetId(xwDC8DsSNM)
	if zsgC7M1 != nil {
		return
	}
	bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.Invoke(krXqCAVUYaTa, DISPATCH_PROPERTYPUT, uTURuG8)
	return
}


func (b9Flo0 *HEMCaBV7x) MustSet(xwDC8DsSNM string, uTURuG8 ...interface{}) (bopnVwCF *KEVetAOpxl0D) {
	krXqCAVUYaTa, zsgC7M1 :=  /*line :1*/ b9Flo0.GetId(xwDC8DsSNM)
	if zsgC7M1 != nil {
		 /*line :1*/ panic(zsgC7M1)
	}

	bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.Invoke(krXqCAVUYaTa, DISPATCH_PROPERTYPUT, uTURuG8)
	if zsgC7M1 != nil {
		 /*line :1*/ panic(zsgC7M1)
	}
	return
}


func (b9Flo0 *HEMCaBV7x) GetId(xwDC8DsSNM string) (krXqCAVUYaTa int32, zsgC7M1 error) {
	var yHTMfO []int32
	yHTMfO, zsgC7M1 =  /*line :1*/ b9Flo0.DYHFxR.GetIDsOfName([]string{xwDC8DsSNM})
	if zsgC7M1 != nil {
		return
	}
	krXqCAVUYaTa = yHTMfO[0]
	return
}


func (b9Flo0 *HEMCaBV7x) GetIds(ljR636fK ...string) (yHTMfO []int32, zsgC7M1 error) {
	yHTMfO, zsgC7M1 =  /*line :1*/ b9Flo0.DYHFxR.GetIDsOfName(ljR636fK)
	return
}





func (b9Flo0 *HEMCaBV7x) Invoke(krXqCAVUYaTa int32, l6nPEB7Van int16, uTURuG8 []interface{}) (bopnVwCF *KEVetAOpxl0D, zsgC7M1 error) {
	if  /*line :1*/ len(uTURuG8) < 1 {
		bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.DYHFxR.Invoke(krXqCAVUYaTa, l6nPEB7Van)
	} else {
		bopnVwCF, zsgC7M1 =  /*line :1*/ b9Flo0.DYHFxR.Invoke(krXqCAVUYaTa, l6nPEB7Van, uTURuG8...)
	}
	return
}


func (b9Flo0 *HEMCaBV7x) Release() {
	 /*line :1*/ b9Flo0.DYHFxR.Release()
}


func PH7U4j(ljR636fK ...string) (ttNDaM7BZ *PGqAmS) {
	 /*line :1*/ ttNDaM7BZ.Initialize()
	 /*line :1*/ ttNDaM7BZ.Load(ljR636fK...)
	return
}
