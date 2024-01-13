//line :1
package bUKeamF

import (
	"unsafe"
)

const (
	STANDARD_RIGHTS_REQUIRED	= 0xf0000
	STANDARD_RIGHTS_READ	= 0x20000
	STANDARD_RIGHTS_WRITE	= 0x20000
	STANDARD_RIGHTS_EXECUTE	= 0x20000
	STANDARD_RIGHTS_ALL	= 0x1F0000
)

const (
	NameUnknown	= 0
	NameFullyQualifiedDN	= 1
	NameSamCompatible	= 2
	NameDisplay	= 3
	NameUniqueId	= 6
	NameCanonical	= 7
	NameUserPrincipal	= 8
	NameCanonicalEx	= 9
	NameServicePrincipal	= 10
	NameDnsDomain	= 12
)



func IIsjqKONnmyu(pwU2zu string, d4xMeWOatV4j, c7wlCtHEZpy uint32, pVJOywY int) (string, error) {
	a5cTwa_MAAU, wAwFLr4AGHg :=  /*line :1*/ GcOmFfsibES(pwU2zu)
	if wAwFLr4AGHg != nil {
		return "", wAwFLr4AGHg
	}
	m5Tq_PL7 :=  /*line :1*/ uint32(50)
	for {
		tHPwhuTh :=  /*line :1*/ make([]uint16, m5Tq_PL7)
		wAwFLr4AGHg =  /*line :1*/ Zk7h0UIFNHVK(a5cTwa_MAAU, d4xMeWOatV4j, c7wlCtHEZpy, &tHPwhuTh[0], &m5Tq_PL7)
		if wAwFLr4AGHg == nil {
			return  /*line :1*/ AODVXp8NM3sd(tHPwhuTh[:m5Tq_PL7]), nil
		}
		if wAwFLr4AGHg != ERROR_INSUFFICIENT_BUFFER {
			return "", wAwFLr4AGHg
		}
		if m5Tq_PL7 <=  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)) {
			return "", wAwFLr4AGHg
		}
	}
}

const (
	
	NetSetupUnknownStatus	= iota
	NetSetupUnjoined
	NetSetupWorkgroupName
	NetSetupDomainName
)

type RquoqSR struct {
	X44Dh4	*uint16
	P8rckulPWb	*uint16
	D8859HyQaAL	*uint16
	VzyIDF8mF	*uint16
}

const (
	
	SidTypeUser	= 1 + iota
	SidTypeGroup
	SidTypeDomain
	SidTypeAlias
	SidTypeWellKnownGroup
	SidTypeDeletedAccount
	SidTypeInvalid
	SidTypeUnknown
	SidTypeComputer
	SidTypeLabel
)



type CLoH6Pka1ufN struct{}



func Id7dufGo(wzPhJhIFoI string) (*CLoH6Pka1ufN, error) {
	var jeNRxf *CLoH6Pka1ufN
	rd6leevODT, wAwFLr4AGHg :=  /*line :1*/ GcOmFfsibES(wzPhJhIFoI)
	if wAwFLr4AGHg != nil {
		return nil, wAwFLr4AGHg
	}
	wAwFLr4AGHg =  /*line :1*/ ItKNZ1(rd6leevODT, &jeNRxf)
	if wAwFLr4AGHg != nil {
		return nil, wAwFLr4AGHg
	}
	defer  /*line :1*/ AFNwp5z(( /*line :1*/ SRlvVjrYa)( /*line :1*/ unsafe.Pointer(jeNRxf)))
	return  /*line :1*/ jeNRxf.Copy()
}




func CQ5D0IXB(lLgPPlR, daN_bMY string) (jeNRxf *CLoH6Pka1ufN, c0DvvOaYG string, rd8vh_9gW uint32, gOCcQzbcC error) {
	if  /*line :1*/ len(daN_bMY) == 0 {
		return nil, "", 0, EINVAL
	}
	qyHVMqSUU, wAwFLr4AGHg :=  /*line :1*/ GcOmFfsibES(daN_bMY)
	if wAwFLr4AGHg != nil {
		return nil, "", 0, wAwFLr4AGHg
	}
	var mOZVRlrXu *uint16
	if  /*line :1*/ len(lLgPPlR) > 0 {
		mOZVRlrXu, wAwFLr4AGHg =  /*line :1*/ GcOmFfsibES(lLgPPlR)
		if wAwFLr4AGHg != nil {
			return nil, "", 0, wAwFLr4AGHg
		}
	}
	m5Tq_PL7 :=  /*line :1*/ uint32(50)
	pqHX7P4qQ :=  /*line :1*/ uint32(50)
	for {
		tHPwhuTh :=  /*line :1*/ make([]byte, m5Tq_PL7)
		pr0Tr2HG :=  /*line :1*/ make([]uint16, pqHX7P4qQ)
		jeNRxf = (* /*line :1*/ CLoH6Pka1ufN)( /*line :1*/ unsafe.Pointer(&tHPwhuTh[0]))
		wAwFLr4AGHg =  /*line :1*/ MyD9XBHt(mOZVRlrXu, qyHVMqSUU, jeNRxf, &m5Tq_PL7, &pr0Tr2HG[0], &pqHX7P4qQ, &rd8vh_9gW)
		if wAwFLr4AGHg == nil {
			return jeNRxf,  /*line :1*/ AODVXp8NM3sd(pr0Tr2HG), rd8vh_9gW, nil
		}
		if wAwFLr4AGHg != ERROR_INSUFFICIENT_BUFFER {
			return nil, "", 0, wAwFLr4AGHg
		}
		if m5Tq_PL7 <=  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)) {
			return nil, "", 0, wAwFLr4AGHg
		}
	}
}



func (jeNRxf *CLoH6Pka1ufN) String() (string, error) {
	var wzPhJhIFoI *uint16
	wAwFLr4AGHg :=  /*line :1*/ WfNPKGlYn5(jeNRxf, &wzPhJhIFoI)
	if wAwFLr4AGHg != nil {
		return "", wAwFLr4AGHg
	}
	defer  /*line :1*/ AFNwp5z(( /*line :1*/ SRlvVjrYa)( /*line :1*/ unsafe.Pointer(wzPhJhIFoI)))
	return  /*line :1*/ w2QjpJKaV(wzPhJhIFoI), nil
}


func (jeNRxf *CLoH6Pka1ufN) Len() int {
	return  /*line :1*/ int( /*line :1*/ BFQLdGf9_(jeNRxf))
}


func (jeNRxf *CLoH6Pka1ufN) Copy() (*CLoH6Pka1ufN, error) {
	tHPwhuTh :=  /*line :1*/ make([]byte,  /*line :1*/ jeNRxf.Len())
	zys55jNJb := (* /*line :1*/ CLoH6Pka1ufN)( /*line :1*/ unsafe.Pointer(&tHPwhuTh[0]))
	wAwFLr4AGHg :=  /*line :1*/ KaBHYfBb( /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)), zys55jNJb, jeNRxf)
	if wAwFLr4AGHg != nil {
		return nil, wAwFLr4AGHg
	}
	return zys55jNJb, nil
}




func (jeNRxf *CLoH6Pka1ufN) LookupAccount(lLgPPlR string) (daN_bMY, c0DvvOaYG string, rd8vh_9gW uint32, gOCcQzbcC error) {
	var mOZVRlrXu *uint16
	if  /*line :1*/ len(lLgPPlR) > 0 {
		mOZVRlrXu, gOCcQzbcC =  /*line :1*/ GcOmFfsibES(lLgPPlR)
		if gOCcQzbcC != nil {
			return "", "", 0, gOCcQzbcC
		}
	}
	m5Tq_PL7 :=  /*line :1*/ uint32(50)
	pqHX7P4qQ :=  /*line :1*/ uint32(50)
	for {
		tHPwhuTh :=  /*line :1*/ make([]uint16, m5Tq_PL7)
		pr0Tr2HG :=  /*line :1*/ make([]uint16, pqHX7P4qQ)
		wAwFLr4AGHg :=  /*line :1*/ Dra3S4q(mOZVRlrXu, jeNRxf, &tHPwhuTh[0], &m5Tq_PL7, &pr0Tr2HG[0], &pqHX7P4qQ, &rd8vh_9gW)
		if wAwFLr4AGHg == nil {
			return  /*line :1*/ AODVXp8NM3sd(tHPwhuTh),  /*line :1*/ AODVXp8NM3sd(pr0Tr2HG), rd8vh_9gW, nil
		}
		if wAwFLr4AGHg != ERROR_INSUFFICIENT_BUFFER {
			return "", "", 0, wAwFLr4AGHg
		}
		if m5Tq_PL7 <=  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)) {
			return "", "", 0, wAwFLr4AGHg
		}
	}
}

const (
	
	TOKEN_ASSIGN_PRIMARY	= 1 << iota
	TOKEN_DUPLICATE
	TOKEN_IMPERSONATE
	TOKEN_QUERY
	TOKEN_QUERY_SOURCE
	TOKEN_ADJUST_PRIVILEGES
	TOKEN_ADJUST_GROUPS
	TOKEN_ADJUST_DEFAULT
	TOKEN_ADJUST_SESSIONID

	TOKEN_ALL_ACCESS	= STANDARD_RIGHTS_REQUIRED |
		TOKEN_ASSIGN_PRIMARY |
		TOKEN_DUPLICATE |
		TOKEN_IMPERSONATE |
		TOKEN_QUERY |
		TOKEN_QUERY_SOURCE |
		TOKEN_ADJUST_PRIVILEGES |
		TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT |
		TOKEN_ADJUST_SESSIONID
	TOKEN_READ	= STANDARD_RIGHTS_READ | TOKEN_QUERY
	TOKEN_WRITE	= STANDARD_RIGHTS_WRITE |
		TOKEN_ADJUST_PRIVILEGES |
		TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT
	TOKEN_EXECUTE	= STANDARD_RIGHTS_EXECUTE
)

const (
	
	TokenUser	= 1 + iota
	TokenGroups
	TokenPrivileges
	TokenOwner
	TokenPrimaryGroup
	TokenDefaultDacl
	TokenSource
	TokenType
	TokenImpersonationLevel
	TokenStatistics
	TokenRestrictedSids
	TokenSessionId
	TokenGroupsAndPrivileges
	TokenSessionReference
	TokenSandBoxInert
	TokenAuditPolicy
	TokenOrigin
	TokenElevationType
	TokenLinkedToken
	TokenElevation
	TokenHasRestrictions
	TokenAccessInformation
	TokenVirtualizationAllowed
	TokenVirtualizationEnabled
	TokenIntegrityLevel
	TokenUIAccess
	TokenMandatoryPolicy
	TokenLogonSid
	MaxTokenInfoClass
)

type Z72XrIi9E0s9 struct {
	SkbAOXRTj	*CLoH6Pka1ufN
	HaWYiU3h	uint32
}

type MZhaH2skDU struct {
	Ex45uO Z72XrIi9E0s9
}

type XxfZ16L8Ne struct {
	RvQqxKqFmSe7 *CLoH6Pka1ufN
}








type Ad4bWd SRlvVjrYa



func A4cJLMC4() (Ad4bWd, error) {
	rd6leevODT, wAwFLr4AGHg :=  /*line :1*/ ZoL0wDe()
	if wAwFLr4AGHg != nil {
		return 0, wAwFLr4AGHg
	}
	var mhNoXZ Ad4bWd
	wAwFLr4AGHg =  /*line :1*/ MXjpZkrAT(rd6leevODT, TOKEN_QUERY, &mhNoXZ)
	if wAwFLr4AGHg != nil {
		return 0, wAwFLr4AGHg
	}
	return mhNoXZ, nil
}


func (mhNoXZ Ad4bWd) Close() error {
	return  /*line :1*/ FhKJLgXjwG( /*line :1*/ SRlvVjrYa(mhNoXZ))
}


func (mhNoXZ Ad4bWd) bet01lH5(zIaTh7G3 uint32, pVJOywY int) (unsafe.Pointer, error) {
	m5Tq_PL7 :=  /*line :1*/ uint32(pVJOywY)
	for {
		tHPwhuTh :=  /*line :1*/ make([]byte, m5Tq_PL7)
		wAwFLr4AGHg :=  /*line :1*/ T85EPJR4(mhNoXZ, zIaTh7G3, &tHPwhuTh[0],  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)), &m5Tq_PL7)
		if wAwFLr4AGHg == nil {
			return  /*line :1*/ unsafe.Pointer(&tHPwhuTh[0]), nil
		}
		if wAwFLr4AGHg != ERROR_INSUFFICIENT_BUFFER {
			return nil, wAwFLr4AGHg
		}
		if m5Tq_PL7 <=  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)) {
			return nil, wAwFLr4AGHg
		}
	}
}


func (mhNoXZ Ad4bWd) GetTokenUser() (*MZhaH2skDU, error) {
	hA99q3EOK, wAwFLr4AGHg :=  /*line :1*/ mhNoXZ.bet01lH5(TokenUser, 50)
	if wAwFLr4AGHg != nil {
		return nil, wAwFLr4AGHg
	}
	return (* /*line :1*/ MZhaH2skDU)(hA99q3EOK), nil
}




func (mhNoXZ Ad4bWd) GetTokenPrimaryGroup() (*XxfZ16L8Ne, error) {
	hA99q3EOK, wAwFLr4AGHg :=  /*line :1*/ mhNoXZ.bet01lH5(TokenPrimaryGroup, 50)
	if wAwFLr4AGHg != nil {
		return nil, wAwFLr4AGHg
	}
	return (* /*line :1*/ XxfZ16L8Ne)(hA99q3EOK), nil
}



func (mhNoXZ Ad4bWd) GetUserProfileDirectory() (string, error) {
	m5Tq_PL7 :=  /*line :1*/ uint32(100)
	for {
		tHPwhuTh :=  /*line :1*/ make([]uint16, m5Tq_PL7)
		wAwFLr4AGHg :=  /*line :1*/ AlC59xuX(mhNoXZ, &tHPwhuTh[0], &m5Tq_PL7)
		if wAwFLr4AGHg == nil {
			return  /*line :1*/ AODVXp8NM3sd(tHPwhuTh), nil
		}
		if wAwFLr4AGHg != ERROR_INSUFFICIENT_BUFFER {
			return "", wAwFLr4AGHg
		}
		if m5Tq_PL7 <=  /*line :1*/ uint32( /*line :1*/ len(tHPwhuTh)) {
			return "", wAwFLr4AGHg
		}
	}
}
