//line :1



package y1BamVjnXsaa

import (
	oserror "K7w4V9ecdx2"
	time "fRAfQd_"
	utf8 "CuAc0pPZwf"
)






type XIcFcDgy interface {
	
	
	
	
	
	
	
	
	
	Open(hMWDrBfy string) (YY1DXRrw, error)
}














func QuVTpr7RfE(hMWDrBfy string) bool {
	if ! /*line :1*/ utf8.YiR9Xl_UC5x(hMWDrBfy) {
		return false
	}

	if hMWDrBfy == "." {

		return true
	}

	for {
		aTOpQutlw := 0
		for aTOpQutlw <  /*line :1*/ len(hMWDrBfy) && hMWDrBfy[aTOpQutlw] != '/' {
			aTOpQutlw++
		}
		pHm_CZZ994U := hMWDrBfy[:aTOpQutlw]
		if pHm_CZZ994U == "" || pHm_CZZ994U == "." || pHm_CZZ994U == ".." {
			return false
		}
		if aTOpQutlw ==  /*line :1*/ len(hMWDrBfy) {
			return true
		}
		hMWDrBfy = hMWDrBfy[aTOpQutlw+1:]
	}
}





type YY1DXRrw interface {
	Stat() (HM4JM2, error)
	Read([]byte) (int, error)
	Close() error
}



type RbfBMxa interface {
	
	
	
	Name() string

	
	IsDir() bool

	
	
	Type() PGXMbJjlHBMu

	
	
	
	
	
	
	Info() (HM4JM2, error)
}





type FhagHI interface {
	YY1DXRrw

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	ReadDir(l5arJmbvBL0f int) ([]RbfBMxa, error)
}




var (
	TJmnAAu1	=  /*line :1*/ guideyMKs()		
	E8Hgdf	=  /*line :1*/ qxnQB6F()		
	GLSb30ePi	=  /*line :1*/ cGHkArpZ()		
	QhJWcuT5mSD	=  /*line :1*/ sjHaob5ctEw()		
	DfgEMfKhaOh	=  /*line :1*/ pSjdB6()		
)

func guideyMKs() error	{ return oserror.ZfmdHupe1 }
func qxnQB6F() error	{ return oserror.HCknuLDV1zb }
func cGHkArpZ() error	{ return oserror.IyvfTZaiEa }
func sjHaob5ctEw() error	{ return oserror.RKRo3rPG5Gn }
func pSjdB6() error	{ return oserror.HrOLpQoNQWW }


type HM4JM2 interface {
	Name() string	
	Size() int64	
	Mode() PGXMbJjlHBMu	
	ModTime() time.G4KDkI	
	IsDir() bool	
	Sys() any	
}






type PGXMbJjlHBMu uint32






const (
	
	
	ModeDir	PGXMbJjlHBMu	= 1 << (32 - 1 - iota)	
	ModeAppend			
	ModeExclusive			
	ModeTemporary			
	ModeSymlink			
	ModeDevice			
	ModeNamedPipe			
	ModeSocket			
	ModeSetuid			
	ModeSetgid			
	ModeCharDevice			
	ModeSticky			
	ModeIrregular			

	
	ModeType		= ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm	PGXMbJjlHBMu	= 0777	
)

func (jREaoX_I7TB6 PGXMbJjlHBMu) String() string {
	const str = "dalTLDpSugct?"
	var zTLJRkZ1 [32]byte	
	kRNHKDVI_ := 0
	for aTOpQutlw, zLsNm3 := range str {
		if jREaoX_I7TB6&(1<< /*line :1*/ uint(32-1-aTOpQutlw)) != 0 {
			zTLJRkZ1[kRNHKDVI_] =  /*line :1*/ byte(zLsNm3)
			kRNHKDVI_++
		}
	}
	if kRNHKDVI_ == 0 {
		zTLJRkZ1[kRNHKDVI_] = '-'
		kRNHKDVI_++
	}
	const rwx = "rwxrwxrwx"
	for aTOpQutlw, zLsNm3 := range rwx {
		if jREaoX_I7TB6&(1<< /*line :1*/ uint(9-1-aTOpQutlw)) != 0 {
			zTLJRkZ1[kRNHKDVI_] =  /*line :1*/ byte(zLsNm3)
		} else {
			zTLJRkZ1[kRNHKDVI_] = '-'
		}
		kRNHKDVI_++
	}
	return  /*line :1*/ string(zTLJRkZ1[:kRNHKDVI_])
}



func (jREaoX_I7TB6 PGXMbJjlHBMu) IsDir() bool {
	return jREaoX_I7TB6&ModeDir != 0
}



func (jREaoX_I7TB6 PGXMbJjlHBMu) IsRegular() bool {
	return jREaoX_I7TB6&ModeType == 0
}


func (jREaoX_I7TB6 PGXMbJjlHBMu) Perm() PGXMbJjlHBMu {
	return jREaoX_I7TB6 & ModePerm
}


func (jREaoX_I7TB6 PGXMbJjlHBMu) Type() PGXMbJjlHBMu {
	return jREaoX_I7TB6 & ModeType
}


type CFLL3J struct {
	Aeg62j1VQt	string
	OmDEtY	string
	OIEifQ0XF	error
}

func (c8anSmr6wZV *CFLL3J) Error() string {
	return c8anSmr6wZV.Aeg62j1VQt + " " + c8anSmr6wZV.OmDEtY + ": " +  /*line :1*/ c8anSmr6wZV.OIEifQ0XF.Error()
}

func (c8anSmr6wZV *CFLL3J) Unwrap() error	{ return c8anSmr6wZV.OIEifQ0XF }


func (c8anSmr6wZV *CFLL3J) Timeout() bool {
	pbRaeae, d5Q0SCgdw := c8anSmr6wZV.OIEifQ0XF.(interface{ Timeout() bool })
	return d5Q0SCgdw &&  /*line :1*/ pbRaeae.Timeout()
}
