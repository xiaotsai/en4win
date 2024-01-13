//line :1
package NJ4MerJ

import (
	syscall "bUKeamF"
	"unsafe"
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



func Oxb53EK0a(aATTVqcx8 string, wPbtqIJl, xtbcgl1U7 uint32, oMa1o0wtSc int) (string, error) {
	sMdg1A, aYtmqRVkqc_8 :=  /*line :1*/ Ih3Y4u(aATTVqcx8)
	if aYtmqRVkqc_8 != nil {
		return "", aYtmqRVkqc_8
	}
	krzuku :=  /*line :1*/ uint32(50)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]uint16, krzuku)
		aYtmqRVkqc_8 =  /*line :1*/ El6tR4oKc(sMdg1A, wPbtqIJl, xtbcgl1U7, &hJuLpmVFi[0], &krzuku)
		if aYtmqRVkqc_8 == nil {
			return  /*line :1*/ OXNanQ8Uj(hJuLpmVFi[:krzuku]), nil
		}
		if aYtmqRVkqc_8 != ERROR_INSUFFICIENT_BUFFER {
			return "", aYtmqRVkqc_8
		}
		if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)) {
			return "", aYtmqRVkqc_8
		}
	}
}

const (
	
	NetSetupUnknownStatus	= iota
	NetSetupUnjoined
	NetSetupWorkgroupName
	NetSetupDomainName
)

type Tf8wLh9Nc9uJ struct {
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

type TWV6Elr struct {
	FkCqGL [6]byte
}

var (
	FfzyMkUFes4	= TWV6Elr{[6]byte{0, 0, 0, 0, 0, 0}}
	Ouaod0d55W	= TWV6Elr{[6]byte{0, 0, 0, 0, 0, 1}}
	JINRzQFU	= TWV6Elr{[6]byte{0, 0, 0, 0, 0, 2}}
	UsVB1j_	= TWV6Elr{[6]byte{0, 0, 0, 0, 0, 3}}
	KI86cRdWkN	= TWV6Elr{[6]byte{0, 0, 0, 0, 0, 4}}
	ZTPw6b	= TWV6Elr{[6]byte{0, 0, 0, 0, 0, 5}}
	NRjuYG1n	= TWV6Elr{[6]byte{0, 0, 0, 0, 0, 16}}
)

const (
	SECURITY_NULL_RID	= 0
	SECURITY_WORLD_RID	= 0
	SECURITY_LOCAL_RID	= 0
	SECURITY_CREATOR_OWNER_RID	= 0
	SECURITY_CREATOR_GROUP_RID	= 1
	SECURITY_DIALUP_RID	= 1
	SECURITY_NETWORK_RID	= 2
	SECURITY_BATCH_RID	= 3
	SECURITY_INTERACTIVE_RID	= 4
	SECURITY_LOGON_IDS_RID	= 5
	SECURITY_SERVICE_RID	= 6
	SECURITY_LOCAL_SYSTEM_RID	= 18
	SECURITY_BUILTIN_DOMAIN_RID	= 32
	SECURITY_PRINCIPAL_SELF_RID	= 10
	SECURITY_CREATOR_OWNER_SERVER_RID	= 0x2
	SECURITY_CREATOR_GROUP_SERVER_RID	= 0x3
	SECURITY_LOGON_IDS_RID_COUNT	= 0x3
	SECURITY_ANONYMOUS_LOGON_RID	= 0x7
	SECURITY_PROXY_RID	= 0x8
	SECURITY_ENTERPRISE_CONTROLLERS_RID	= 0x9
	SECURITY_SERVER_LOGON_RID	= SECURITY_ENTERPRISE_CONTROLLERS_RID
	SECURITY_AUTHENTICATED_USER_RID	= 0xb
	SECURITY_RESTRICTED_CODE_RID	= 0xc
	SECURITY_NT_NON_UNIQUE_RID	= 0x15
)



const (
	DOMAIN_ALIAS_RID_ADMINS	= 0x220
	DOMAIN_ALIAS_RID_USERS	= 0x221
	DOMAIN_ALIAS_RID_GUESTS	= 0x222
	DOMAIN_ALIAS_RID_POWER_USERS	= 0x223
	DOMAIN_ALIAS_RID_ACCOUNT_OPS	= 0x224
	DOMAIN_ALIAS_RID_SYSTEM_OPS	= 0x225
	DOMAIN_ALIAS_RID_PRINT_OPS	= 0x226
	DOMAIN_ALIAS_RID_BACKUP_OPS	= 0x227
	DOMAIN_ALIAS_RID_REPLICATOR	= 0x228
	DOMAIN_ALIAS_RID_RAS_SERVERS	= 0x229
	DOMAIN_ALIAS_RID_PREW2KCOMPACCESS	= 0x22a
	DOMAIN_ALIAS_RID_REMOTE_DESKTOP_USERS	= 0x22b
	DOMAIN_ALIAS_RID_NETWORK_CONFIGURATION_OPS	= 0x22c
	DOMAIN_ALIAS_RID_INCOMING_FOREST_TRUST_BUILDERS	= 0x22d
	DOMAIN_ALIAS_RID_MONITORING_USERS	= 0x22e
	DOMAIN_ALIAS_RID_LOGGING_USERS	= 0x22f
	DOMAIN_ALIAS_RID_AUTHORIZATIONACCESS	= 0x230
	DOMAIN_ALIAS_RID_TS_LICENSE_SERVERS	= 0x231
	DOMAIN_ALIAS_RID_DCOM_USERS	= 0x232
	DOMAIN_ALIAS_RID_IUSERS	= 0x238
	DOMAIN_ALIAS_RID_CRYPTO_OPERATORS	= 0x239
	DOMAIN_ALIAS_RID_CACHEABLE_PRINCIPALS_GROUP	= 0x23b
	DOMAIN_ALIAS_RID_NON_CACHEABLE_PRINCIPALS_GROUP	= 0x23c
	DOMAIN_ALIAS_RID_EVENT_LOG_READERS_GROUP	= 0x23d
	DOMAIN_ALIAS_RID_CERTSVC_DCOM_ACCESS_GROUP	= 0x23e
)



type He5NCq struct{}



func PzHz3FjdHL(bamc83qA3DBr string) (*He5NCq, error) {
	var ohXw5VDcqx *He5NCq
	hD4wNgEB, aYtmqRVkqc_8 :=  /*line :1*/ Ih3Y4u(bamc83qA3DBr)
	if aYtmqRVkqc_8 != nil {
		return nil, aYtmqRVkqc_8
	}
	aYtmqRVkqc_8 =  /*line :1*/ G8mXygYjUF4(hD4wNgEB, &ohXw5VDcqx)
	if aYtmqRVkqc_8 != nil {
		return nil, aYtmqRVkqc_8
	}
	defer  /*line :1*/ OtMI6tu9(( /*line :1*/ L2L8P9WaNZ)( /*line :1*/ unsafe.Pointer(ohXw5VDcqx)))
	return  /*line :1*/ ohXw5VDcqx.Copy()
}




func SfPhK8h(uxLgaklMr2w5, puRq5cR9 string) (ohXw5VDcqx *He5NCq, lfrJYrOPlciq string, e_B7myi9f9Ri uint32, jeWMpOaQtWV error) {
	if  /*line :1*/ len(puRq5cR9) == 0 {
		return nil, "", 0, syscall.EINVAL
	}
	rx7qosn, aYtmqRVkqc_8 :=  /*line :1*/ Ih3Y4u(puRq5cR9)
	if aYtmqRVkqc_8 != nil {
		return nil, "", 0, aYtmqRVkqc_8
	}
	var f3v0ShGQ *uint16
	if  /*line :1*/ len(uxLgaklMr2w5) > 0 {
		f3v0ShGQ, aYtmqRVkqc_8 =  /*line :1*/ Ih3Y4u(uxLgaklMr2w5)
		if aYtmqRVkqc_8 != nil {
			return nil, "", 0, aYtmqRVkqc_8
		}
	}
	krzuku :=  /*line :1*/ uint32(50)
	yhu7oavbN7w :=  /*line :1*/ uint32(50)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]byte, krzuku)
		hTLnWzOrY :=  /*line :1*/ make([]uint16, yhu7oavbN7w)
		ohXw5VDcqx = (* /*line :1*/ He5NCq)( /*line :1*/ unsafe.Pointer(&hJuLpmVFi[0]))
		aYtmqRVkqc_8 =  /*line :1*/ UHIB5J(f3v0ShGQ, rx7qosn, ohXw5VDcqx, &krzuku, &hTLnWzOrY[0], &yhu7oavbN7w, &e_B7myi9f9Ri)
		if aYtmqRVkqc_8 == nil {
			return ohXw5VDcqx,  /*line :1*/ OXNanQ8Uj(hTLnWzOrY), e_B7myi9f9Ri, nil
		}
		if aYtmqRVkqc_8 != ERROR_INSUFFICIENT_BUFFER {
			return nil, "", 0, aYtmqRVkqc_8
		}
		if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)) {
			return nil, "", 0, aYtmqRVkqc_8
		}
	}
}


func (ohXw5VDcqx *He5NCq) String() string {
	var bamc83qA3DBr *uint16
	aYtmqRVkqc_8 :=  /*line :1*/ CyEzI3t6(ohXw5VDcqx, &bamc83qA3DBr)
	if aYtmqRVkqc_8 != nil {
		return ""
	}
	defer  /*line :1*/ OtMI6tu9(( /*line :1*/ L2L8P9WaNZ)( /*line :1*/ unsafe.Pointer(bamc83qA3DBr)))
	return  /*line :1*/ OXNanQ8Uj((*[256] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(bamc83qA3DBr))[:])
}


func (ohXw5VDcqx *He5NCq) Len() int {
	return  /*line :1*/ int( /*line :1*/ C6xEOAk(ohXw5VDcqx))
}


func (ohXw5VDcqx *He5NCq) Copy() (*He5NCq, error) {
	hJuLpmVFi :=  /*line :1*/ make([]byte,  /*line :1*/ ohXw5VDcqx.Len())
	fW6sSe := (* /*line :1*/ He5NCq)( /*line :1*/ unsafe.Pointer(&hJuLpmVFi[0]))
	aYtmqRVkqc_8 :=  /*line :1*/ OCRaVs( /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)), fW6sSe, ohXw5VDcqx)
	if aYtmqRVkqc_8 != nil {
		return nil, aYtmqRVkqc_8
	}
	return fW6sSe, nil
}


func (ohXw5VDcqx *He5NCq) IdentifierAuthority() TWV6Elr {
	return * /*line :1*/ bSEeAoS6s1DR(ohXw5VDcqx)
}


func (ohXw5VDcqx *He5NCq) SubAuthorityCount() uint8 {
	return * /*line :1*/ a2nVuCjv(ohXw5VDcqx)
}



func (ohXw5VDcqx *He5NCq) SubAuthority(aHmRuC uint32) uint32 {
	if aHmRuC >=  /*line :1*/ uint32( /*line :1*/ ohXw5VDcqx.SubAuthorityCount()) {
		 /*line :1*/ panic("sub-authority index out of range")
	}
	return * /*line :1*/ jnbf6VRbBCS(ohXw5VDcqx, aHmRuC)
}


func (ohXw5VDcqx *He5NCq) IsValid() bool {
	return  /*line :1*/ aUzjOSgDD8a(ohXw5VDcqx)
}


func (ohXw5VDcqx *He5NCq) Equals(fW6sSe *He5NCq) bool {
	return  /*line :1*/ SJNWqa9WP(ohXw5VDcqx, fW6sSe)
}


func (ohXw5VDcqx *He5NCq) IsWellKnown(a4Y5Dlp1iHO CV4yg1b) bool {
	return  /*line :1*/ cKxRf5(ohXw5VDcqx, a4Y5Dlp1iHO)
}




func (ohXw5VDcqx *He5NCq) LookupAccount(uxLgaklMr2w5 string) (puRq5cR9, lfrJYrOPlciq string, e_B7myi9f9Ri uint32, jeWMpOaQtWV error) {
	var f3v0ShGQ *uint16
	if  /*line :1*/ len(uxLgaklMr2w5) > 0 {
		f3v0ShGQ, jeWMpOaQtWV =  /*line :1*/ Ih3Y4u(uxLgaklMr2w5)
		if jeWMpOaQtWV != nil {
			return "", "", 0, jeWMpOaQtWV
		}
	}
	krzuku :=  /*line :1*/ uint32(50)
	yhu7oavbN7w :=  /*line :1*/ uint32(50)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]uint16, krzuku)
		hTLnWzOrY :=  /*line :1*/ make([]uint16, yhu7oavbN7w)
		aYtmqRVkqc_8 :=  /*line :1*/ LgFpMX8d65(f3v0ShGQ, ohXw5VDcqx, &hJuLpmVFi[0], &krzuku, &hTLnWzOrY[0], &yhu7oavbN7w, &e_B7myi9f9Ri)
		if aYtmqRVkqc_8 == nil {
			return  /*line :1*/ OXNanQ8Uj(hJuLpmVFi),  /*line :1*/ OXNanQ8Uj(hTLnWzOrY), e_B7myi9f9Ri, nil
		}
		if aYtmqRVkqc_8 != ERROR_INSUFFICIENT_BUFFER {
			return "", "", 0, aYtmqRVkqc_8
		}
		if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)) {
			return "", "", 0, aYtmqRVkqc_8
		}
	}
}


type CV4yg1b uint32

const (
	WinNullSid	= 0
	WinWorldSid	= 1
	WinLocalSid	= 2
	WinCreatorOwnerSid	= 3
	WinCreatorGroupSid	= 4
	WinCreatorOwnerServerSid	= 5
	WinCreatorGroupServerSid	= 6
	WinNtAuthoritySid	= 7
	WinDialupSid	= 8
	WinNetworkSid	= 9
	WinBatchSid	= 10
	WinInteractiveSid	= 11
	WinServiceSid	= 12
	WinAnonymousSid	= 13
	WinProxySid	= 14
	WinEnterpriseControllersSid	= 15
	WinSelfSid	= 16
	WinAuthenticatedUserSid	= 17
	WinRestrictedCodeSid	= 18
	WinTerminalServerSid	= 19
	WinRemoteLogonIdSid	= 20
	WinLogonIdsSid	= 21
	WinLocalSystemSid	= 22
	WinLocalServiceSid	= 23
	WinNetworkServiceSid	= 24
	WinBuiltinDomainSid	= 25
	WinBuiltinAdministratorsSid	= 26
	WinBuiltinUsersSid	= 27
	WinBuiltinGuestsSid	= 28
	WinBuiltinPowerUsersSid	= 29
	WinBuiltinAccountOperatorsSid	= 30
	WinBuiltinSystemOperatorsSid	= 31
	WinBuiltinPrintOperatorsSid	= 32
	WinBuiltinBackupOperatorsSid	= 33
	WinBuiltinReplicatorSid	= 34
	WinBuiltinPreWindows2000CompatibleAccessSid	= 35
	WinBuiltinRemoteDesktopUsersSid	= 36
	WinBuiltinNetworkConfigurationOperatorsSid	= 37
	WinAccountAdministratorSid	= 38
	WinAccountGuestSid	= 39
	WinAccountKrbtgtSid	= 40
	WinAccountDomainAdminsSid	= 41
	WinAccountDomainUsersSid	= 42
	WinAccountDomainGuestsSid	= 43
	WinAccountComputersSid	= 44
	WinAccountControllersSid	= 45
	WinAccountCertAdminsSid	= 46
	WinAccountSchemaAdminsSid	= 47
	WinAccountEnterpriseAdminsSid	= 48
	WinAccountPolicyAdminsSid	= 49
	WinAccountRasAndIasServersSid	= 50
	WinNTLMAuthenticationSid	= 51
	WinDigestAuthenticationSid	= 52
	WinSChannelAuthenticationSid	= 53
	WinThisOrganizationSid	= 54
	WinOtherOrganizationSid	= 55
	WinBuiltinIncomingForestTrustBuildersSid	= 56
	WinBuiltinPerfMonitoringUsersSid	= 57
	WinBuiltinPerfLoggingUsersSid	= 58
	WinBuiltinAuthorizationAccessSid	= 59
	WinBuiltinTerminalServerLicenseServersSid	= 60
	WinBuiltinDCOMUsersSid	= 61
	WinBuiltinIUsersSid	= 62
	WinIUserSid	= 63
	WinBuiltinCryptoOperatorsSid	= 64
	WinUntrustedLabelSid	= 65
	WinLowLabelSid	= 66
	WinMediumLabelSid	= 67
	WinHighLabelSid	= 68
	WinSystemLabelSid	= 69
	WinWriteRestrictedCodeSid	= 70
	WinCreatorOwnerRightsSid	= 71
	WinCacheablePrincipalsGroupSid	= 72
	WinNonCacheablePrincipalsGroupSid	= 73
	WinEnterpriseReadonlyControllersSid	= 74
	WinAccountReadonlyControllersSid	= 75
	WinBuiltinEventLogReadersGroup	= 76
	WinNewEnterpriseReadonlyControllersSid	= 77
	WinBuiltinCertSvcDComAccessGroup	= 78
	WinMediumPlusLabelSid	= 79
	WinLocalLogonSid	= 80
	WinConsoleLogonSid	= 81
	WinThisOrganizationCertificateSid	= 82
	WinApplicationPackageAuthoritySid	= 83
	WinBuiltinAnyPackageSid	= 84
	WinCapabilityInternetClientSid	= 85
	WinCapabilityInternetClientServerSid	= 86
	WinCapabilityPrivateNetworkClientServerSid	= 87
	WinCapabilityPicturesLibrarySid	= 88
	WinCapabilityVideosLibrarySid	= 89
	WinCapabilityMusicLibrarySid	= 90
	WinCapabilityDocumentsLibrarySid	= 91
	WinCapabilitySharedUserCertificatesSid	= 92
	WinCapabilityEnterpriseAuthenticationSid	= 93
	WinCapabilityRemovableStorageSid	= 94
	WinBuiltinRDSRemoteAccessServersSid	= 95
	WinBuiltinRDSEndpointServersSid	= 96
	WinBuiltinRDSManagementServersSid	= 97
	WinUserModeDriversSid	= 98
	WinBuiltinHyperVAdminsSid	= 99
	WinAccountCloneableControllersSid	= 100
	WinBuiltinAccessControlAssistanceOperatorsSid	= 101
	WinBuiltinRemoteManagementUsersSid	= 102
	WinAuthenticationAuthorityAssertedSid	= 103
	WinAuthenticationServiceAssertedSid	= 104
	WinLocalAccountSid	= 105
	WinLocalAccountAndAdministratorSid	= 106
	WinAccountProtectedUsersSid	= 107
	WinCapabilityAppointmentsSid	= 108
	WinCapabilityContactsSid	= 109
	WinAccountDefaultSystemManagedSid	= 110
	WinBuiltinDefaultSystemManagedGroupSid	= 111
	WinBuiltinStorageReplicaAdminsSid	= 112
	WinAccountKeyAdminsSid	= 113
	WinAccountEnterpriseKeyAdminsSid	= 114
	WinAuthenticationKeyTrustSid	= 115
	WinAuthenticationKeyPropertyMFASid	= 116
	WinAuthenticationKeyPropertyAttestationSid	= 117
	WinAuthenticationFreshKeyAuthSid	= 118
	WinBuiltinDeviceOwnersSid	= 119
)



func Wy5Pi142b(a4Y5Dlp1iHO CV4yg1b) (*He5NCq, error) {
	return  /*line :1*/ YdNI5cMVQ_3(a4Y5Dlp1iHO, nil)
}



func YdNI5cMVQ_3(a4Y5Dlp1iHO CV4yg1b, f24jbLt77GO *He5NCq) (*He5NCq, error) {
	krzuku :=  /*line :1*/ uint32(50)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]byte, krzuku)
		ohXw5VDcqx := (* /*line :1*/ He5NCq)( /*line :1*/ unsafe.Pointer(&hJuLpmVFi[0]))
		jeWMpOaQtWV :=  /*line :1*/ ctKKBU8R(a4Y5Dlp1iHO, f24jbLt77GO, ohXw5VDcqx, &krzuku)
		if jeWMpOaQtWV == nil {
			return ohXw5VDcqx, nil
		}
		if jeWMpOaQtWV != ERROR_INSUFFICIENT_BUFFER {
			return nil, jeWMpOaQtWV
		}
		if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)) {
			return nil, jeWMpOaQtWV
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


const (
	SE_GROUP_MANDATORY	= 0x00000001
	SE_GROUP_ENABLED_BY_DEFAULT	= 0x00000002
	SE_GROUP_ENABLED	= 0x00000004
	SE_GROUP_OWNER	= 0x00000008
	SE_GROUP_USE_FOR_DENY_ONLY	= 0x00000010
	SE_GROUP_INTEGRITY	= 0x00000020
	SE_GROUP_INTEGRITY_ENABLED	= 0x00000040
	SE_GROUP_LOGON_ID	= 0xC0000000
	SE_GROUP_RESOURCE	= 0x20000000
	SE_GROUP_VALID_ATTRIBUTES	= SE_GROUP_MANDATORY | SE_GROUP_ENABLED_BY_DEFAULT | SE_GROUP_ENABLED | SE_GROUP_OWNER | SE_GROUP_USE_FOR_DENY_ONLY | SE_GROUP_LOGON_ID | SE_GROUP_RESOURCE | SE_GROUP_INTEGRITY | SE_GROUP_INTEGRITY_ENABLED
)


const (
	SE_PRIVILEGE_ENABLED_BY_DEFAULT	= 0x00000001
	SE_PRIVILEGE_ENABLED	= 0x00000002
	SE_PRIVILEGE_REMOVED	= 0x00000004
	SE_PRIVILEGE_USED_FOR_ACCESS	= 0x80000000
	SE_PRIVILEGE_VALID_ATTRIBUTES	= SE_PRIVILEGE_ENABLED_BY_DEFAULT | SE_PRIVILEGE_ENABLED | SE_PRIVILEGE_REMOVED | SE_PRIVILEGE_USED_FOR_ACCESS
)


const (
	TokenPrimary	= 1
	TokenImpersonation	= 2
)


const (
	SecurityAnonymous	= 0
	SecurityIdentification	= 1
	SecurityImpersonation	= 2
	SecurityDelegation	= 3
)

type FHc2HnJ struct {
	Cwsdj6SDrP	uint32
	Hfe4CvtpDz	int32
}

type IJW5sgOW5C struct {
	Uiv3Kk	FHc2HnJ
	QGyjC_ZVUSz	uint32
}

type Adh9WpO struct {
	GwpLAfBfp	*He5NCq
	RtWrNiW2le	uint32
}

type MrXvf7ds struct {
	BY6E43miPH Adh9WpO
}

type Mod20aFaM struct {
	G0TpT_ *He5NCq
}

type AA19MuCnSDP struct {
	F3SnfLmIQvK	uint32
	Sql1KTXD	[1]Adh9WpO	
}


func (vZkygeehK *AA19MuCnSDP) AllGroups() []Adh9WpO {
	return (*[(1 << 28) - 1] /*line :1*/ Adh9WpO)( /*line :1*/ unsafe.Pointer(&vZkygeehK.Sql1KTXD[0]))[:vZkygeehK.F3SnfLmIQvK:vZkygeehK.F3SnfLmIQvK]
}

type Z_rdEkS3_ struct {
	SF58fC2aSK	uint32
	DG_DtZ6NWvt	[1]IJW5sgOW5C	
}


func (hD4wNgEB *Z_rdEkS3_) AllPrivileges() []IJW5sgOW5C {
	return (*[(1 << 27) - 1] /*line :1*/ IJW5sgOW5C)( /*line :1*/ unsafe.Pointer(&hD4wNgEB.DG_DtZ6NWvt[0]))[:hD4wNgEB.SF58fC2aSK:hD4wNgEB.SF58fC2aSK]
}

type ZzskUm2N7 struct {
	GCBGT58Cc1M Adh9WpO
}

func (a_d5hjbf_K1v *ZzskUm2N7) Size() uint32 {
	return  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(ZzskUm2N7{})) +  /*line :1*/ C6xEOAk(a_d5hjbf_K1v.GCBGT58Cc1M.GwpLAfBfp)
}








type TaSPPoJMlh L2L8P9WaNZ







func T0cTFe9x3E() (TaSPPoJMlh, error) {
	var lAmUDbCC TaSPPoJMlh
	jeWMpOaQtWV :=  /*line :1*/ RHTKRt( /*line :1*/ RwYrTC0pJaxW(), TOKEN_QUERY, &lAmUDbCC)
	return lAmUDbCC, jeWMpOaQtWV
}




func NpTg5a() TaSPPoJMlh {
	return  /*line :1*/ TaSPPoJMlh(^ /*line :1*/ uintptr(4 - 1))
}




func Mu0scmM3p5() TaSPPoJMlh {
	return  /*line :1*/ TaSPPoJMlh(^ /*line :1*/ uintptr(5 - 1))
}




func YGCGbAlsnLv() TaSPPoJMlh {
	return  /*line :1*/ TaSPPoJMlh(^ /*line :1*/ uintptr(6 - 1))
}


func (g9JYFzvU35 TaSPPoJMlh) Close() error {
	return  /*line :1*/ E5j7kDrdt( /*line :1*/ L2L8P9WaNZ(g9JYFzvU35))
}


func (g9JYFzvU35 TaSPPoJMlh) cXbO7IvOg(jzcRahcgUWi2 uint32, oMa1o0wtSc int) (unsafe.Pointer, error) {
	krzuku :=  /*line :1*/ uint32(oMa1o0wtSc)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]byte, krzuku)
		aYtmqRVkqc_8 :=  /*line :1*/ JlkmoPdky432(g9JYFzvU35, jzcRahcgUWi2, &hJuLpmVFi[0],  /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)), &krzuku)
		if aYtmqRVkqc_8 == nil {
			return  /*line :1*/ unsafe.Pointer(&hJuLpmVFi[0]), nil
		}
		if aYtmqRVkqc_8 != ERROR_INSUFFICIENT_BUFFER {
			return nil, aYtmqRVkqc_8
		}
		if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)) {
			return nil, aYtmqRVkqc_8
		}
	}
}


func (g9JYFzvU35 TaSPPoJMlh) GetTokenUser() (*MrXvf7ds, error) {
	rRGfxgPH8Kq, aYtmqRVkqc_8 :=  /*line :1*/ g9JYFzvU35.cXbO7IvOg(TokenUser, 50)
	if aYtmqRVkqc_8 != nil {
		return nil, aYtmqRVkqc_8
	}
	return (* /*line :1*/ MrXvf7ds)(rRGfxgPH8Kq), nil
}


func (g9JYFzvU35 TaSPPoJMlh) GetTokenGroups() (*AA19MuCnSDP, error) {
	rRGfxgPH8Kq, aYtmqRVkqc_8 :=  /*line :1*/ g9JYFzvU35.cXbO7IvOg(TokenGroups, 50)
	if aYtmqRVkqc_8 != nil {
		return nil, aYtmqRVkqc_8
	}
	return (* /*line :1*/ AA19MuCnSDP)(rRGfxgPH8Kq), nil
}




func (g9JYFzvU35 TaSPPoJMlh) GetTokenPrimaryGroup() (*Mod20aFaM, error) {
	rRGfxgPH8Kq, aYtmqRVkqc_8 :=  /*line :1*/ g9JYFzvU35.cXbO7IvOg(TokenPrimaryGroup, 50)
	if aYtmqRVkqc_8 != nil {
		return nil, aYtmqRVkqc_8
	}
	return (* /*line :1*/ Mod20aFaM)(rRGfxgPH8Kq), nil
}



func (g9JYFzvU35 TaSPPoJMlh) GetUserProfileDirectory() (string, error) {
	krzuku :=  /*line :1*/ uint32(100)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]uint16, krzuku)
		aYtmqRVkqc_8 :=  /*line :1*/ A9DOEMqTc(g9JYFzvU35, &hJuLpmVFi[0], &krzuku)
		if aYtmqRVkqc_8 == nil {
			return  /*line :1*/ OXNanQ8Uj(hJuLpmVFi), nil
		}
		if aYtmqRVkqc_8 != ERROR_INSUFFICIENT_BUFFER {
			return "", aYtmqRVkqc_8
		}
		if krzuku <=  /*line :1*/ uint32( /*line :1*/ len(hJuLpmVFi)) {
			return "", aYtmqRVkqc_8
		}
	}
}


func (lAmUDbCC TaSPPoJMlh) IsElevated() bool {
	var dx1Xq35ZW uint32
	var vjbD82EI uint32
	jeWMpOaQtWV :=  /*line :1*/ JlkmoPdky432(lAmUDbCC, TokenElevation, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&dx1Xq35ZW)),  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(dx1Xq35ZW)), &vjbD82EI)
	if jeWMpOaQtWV != nil {
		return false
	}
	return vjbD82EI ==  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(dx1Xq35ZW)) && dx1Xq35ZW != 0
}


func (lAmUDbCC TaSPPoJMlh) GetLinkedToken() (TaSPPoJMlh, error) {
	var onDeR4QE TaSPPoJMlh
	var vjbD82EI uint32
	jeWMpOaQtWV :=  /*line :1*/ JlkmoPdky432(lAmUDbCC, TokenLinkedToken, (* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&onDeR4QE)),  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(onDeR4QE)), &vjbD82EI)
	if jeWMpOaQtWV != nil {
		return  /*line :1*/ TaSPPoJMlh(0), jeWMpOaQtWV
	}
	return onDeR4QE, nil
}



func Voabn888O() (string, error) {
	krzuku :=  /*line :1*/ uint32(MAX_PATH)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]uint16, krzuku)
		cTpvqfM, aYtmqRVkqc_8 :=  /*line :1*/ iJsZwZKh(&hJuLpmVFi[0], krzuku)
		if aYtmqRVkqc_8 != nil {
			return "", aYtmqRVkqc_8
		}
		if cTpvqfM <= krzuku {
			return  /*line :1*/ OXNanQ8Uj(hJuLpmVFi[:cTpvqfM]), nil
		}
		krzuku = cTpvqfM
	}
}





func QiaTtT() (string, error) {
	krzuku :=  /*line :1*/ uint32(MAX_PATH)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]uint16, krzuku)
		cTpvqfM, aYtmqRVkqc_8 :=  /*line :1*/ a8LRLFSt(&hJuLpmVFi[0], krzuku)
		if aYtmqRVkqc_8 != nil {
			return "", aYtmqRVkqc_8
		}
		if cTpvqfM <= krzuku {
			return  /*line :1*/ OXNanQ8Uj(hJuLpmVFi[:cTpvqfM]), nil
		}
		krzuku = cTpvqfM
	}
}



func Jf_XG08N() (string, error) {
	krzuku :=  /*line :1*/ uint32(MAX_PATH)
	for {
		hJuLpmVFi :=  /*line :1*/ make([]uint16, krzuku)
		cTpvqfM, aYtmqRVkqc_8 :=  /*line :1*/ oLclNaBVbkN(&hJuLpmVFi[0], krzuku)
		if aYtmqRVkqc_8 != nil {
			return "", aYtmqRVkqc_8
		}
		if cTpvqfM <= krzuku {
			return  /*line :1*/ OXNanQ8Uj(hJuLpmVFi[:cTpvqfM]), nil
		}
		krzuku = cTpvqfM
	}
}


func (g9JYFzvU35 TaSPPoJMlh) IsMember(ohXw5VDcqx *He5NCq) (bool, error) {
	var hJuLpmVFi int32
	if aYtmqRVkqc_8 :=  /*line :1*/ xr4MRBcDahh_(g9JYFzvU35, ohXw5VDcqx, &hJuLpmVFi); aYtmqRVkqc_8 != nil {
		return false, aYtmqRVkqc_8
	}
	return hJuLpmVFi != 0, nil
}


func (g9JYFzvU35 TaSPPoJMlh) IsRestricted() (aMqDbaZ35 bool, jeWMpOaQtWV error) {
	aMqDbaZ35, jeWMpOaQtWV =  /*line :1*/ g2Y5C3WMpl(g9JYFzvU35)
	if !aMqDbaZ35 && jeWMpOaQtWV == syscall.EINVAL {

		jeWMpOaQtWV = nil
	}
	return
}

const (
	WTS_CONSOLE_CONNECT	= 0x1
	WTS_CONSOLE_DISCONNECT	= 0x2
	WTS_REMOTE_CONNECT	= 0x3
	WTS_REMOTE_DISCONNECT	= 0x4
	WTS_SESSION_LOGON	= 0x5
	WTS_SESSION_LOGOFF	= 0x6
	WTS_SESSION_LOCK	= 0x7
	WTS_SESSION_UNLOCK	= 0x8
	WTS_SESSION_REMOTE_CONTROL	= 0x9
	WTS_SESSION_CREATE	= 0xa
	WTS_SESSION_TERMINATE	= 0xb
)

const (
	WTSActive	= 0
	WTSConnected	= 1
	WTSConnectQuery	= 2
	WTSShadow	= 3
	WTSDisconnected	= 4
	WTSIdle	= 5
	WTSListen	= 6
	WTSReset	= 7
	WTSDown	= 8
	WTSInit	= 9
)

type JKJu965 struct {
	I7TzSrza	uint32
	MEu5SFxVN	uint32
}

type PLfKrMR struct {
	O3kBc2e45tG	uint32
	DcUADt	*uint16
	SCBqFt	uint32
}

type Na2cliTdMq struct {
	zqhnRj	byte
	epUSqtQjgA	byte
	svUIDH8yUWzh	uint16
	_jFJmk	uint16
	jYBtQ44W	uint16
}

type A2t_Jvj3ECv struct {
	bePt5WrxeC	byte
	uEPmhu	byte
	dS97MSCNr	OjSqAgjIy
	sCi0c6DuRle	*He5NCq
	krBd0W	*He5NCq
	hHdKDR5R1wn	*Na2cliTdMq
	lp9v6xnkea	*Na2cliTdMq
}

type CgwXVOou2f struct {
	N4xbZD2Bd4	uint32
	GeWG61jIU	uint32
	TVAx0yf	byte
	F_FbWxFzDp	byte
}


const (
	SECURITY_STATIC_TRACKING	= 0
	SECURITY_DYNAMIC_TRACKING	= 1
)

type NRXOeZVS9w struct {
	Ec1Ig0d	uint32
	GzSSw1A3RcJ	*A2t_Jvj3ECv
	Z73_4lh2ro	uint32
}

type J4Q5ilM0dSeR uint32


const (
	SE_UNKNOWN_OBJECT_TYPE	= 0
	SE_FILE_OBJECT	= 1
	SE_SERVICE	= 2
	SE_PRINTER	= 3
	SE_REGISTRY_KEY	= 4
	SE_LMSHARE	= 5
	SE_KERNEL_OBJECT	= 6
	SE_WINDOW_OBJECT	= 7
	SE_DS_OBJECT	= 8
	SE_DS_OBJECT_ALL	= 9
	SE_PROVIDER_DEFINED_OBJECT	= 10
	SE_WMIGUID_OBJECT	= 11
	SE_REGISTRY_WOW64_32KEY	= 12
	SE_REGISTRY_WOW64_64KEY	= 13
)

type B4w2t8D uint32


const (
	OWNER_SECURITY_INFORMATION	= 0x00000001
	GROUP_SECURITY_INFORMATION	= 0x00000002
	DACL_SECURITY_INFORMATION	= 0x00000004
	SACL_SECURITY_INFORMATION	= 0x00000008
	LABEL_SECURITY_INFORMATION	= 0x00000010
	ATTRIBUTE_SECURITY_INFORMATION	= 0x00000020
	SCOPE_SECURITY_INFORMATION	= 0x00000040
	BACKUP_SECURITY_INFORMATION	= 0x00010000
	PROTECTED_DACL_SECURITY_INFORMATION	= 0x80000000
	PROTECTED_SACL_SECURITY_INFORMATION	= 0x40000000
	UNPROTECTED_DACL_SECURITY_INFORMATION	= 0x20000000
	UNPROTECTED_SACL_SECURITY_INFORMATION	= 0x10000000
)

type OjSqAgjIy uint16


const (
	SE_OWNER_DEFAULTED	= 0x0001
	SE_GROUP_DEFAULTED	= 0x0002
	SE_DACL_PRESENT	= 0x0004
	SE_DACL_DEFAULTED	= 0x0008
	SE_SACL_PRESENT	= 0x0010
	SE_SACL_DEFAULTED	= 0x0020
	SE_DACL_AUTO_INHERIT_REQ	= 0x0100
	SE_SACL_AUTO_INHERIT_REQ	= 0x0200
	SE_DACL_AUTO_INHERITED	= 0x0400
	SE_SACL_AUTO_INHERITED	= 0x0800
	SE_DACL_PROTECTED	= 0x1000
	SE_SACL_PROTECTED	= 0x2000
	SE_RM_CONTROL_VALID	= 0x4000
	SE_SELF_RELATIVE	= 0x8000
)

type AlbfLJMLys5 uint32


const (
	DELETE	= 0x00010000
	READ_CONTROL	= 0x00020000
	WRITE_DAC	= 0x00040000
	WRITE_OWNER	= 0x00080000
	SYNCHRONIZE	= 0x00100000
	STANDARD_RIGHTS_REQUIRED	= 0x000F0000
	STANDARD_RIGHTS_READ	= READ_CONTROL
	STANDARD_RIGHTS_WRITE	= READ_CONTROL
	STANDARD_RIGHTS_EXECUTE	= READ_CONTROL
	STANDARD_RIGHTS_ALL	= 0x001F0000
	SPECIFIC_RIGHTS_ALL	= 0x0000FFFF
	ACCESS_SYSTEM_SECURITY	= 0x01000000
	MAXIMUM_ALLOWED	= 0x02000000
	GENERIC_READ	= 0x80000000
	GENERIC_WRITE	= 0x40000000
	GENERIC_EXECUTE	= 0x20000000
	GENERIC_ALL	= 0x10000000
)

type EAKnpNCN uint32


const (
	NOT_USED_ACCESS	= 0
	GRANT_ACCESS	= 1
	SET_ACCESS	= 2
	DENY_ACCESS	= 3
	REVOKE_ACCESS	= 4
	SET_AUDIT_SUCCESS	= 5
	SET_AUDIT_FAILURE	= 6
)


const (
	NO_INHERITANCE	= 0x0
	SUB_OBJECTS_ONLY_INHERIT	= 0x1
	SUB_CONTAINERS_ONLY_INHERIT	= 0x2
	SUB_CONTAINERS_AND_OBJECTS_INHERIT	= 0x3
	INHERIT_NO_PROPAGATE	= 0x4
	INHERIT_ONLY	= 0x8
	INHERITED_ACCESS_ENTRY	= 0x10
	INHERITED_PARENT	= 0x10000000
	INHERITED_GRANDPARENT	= 0x20000000
	OBJECT_INHERIT_ACE	= 0x1
	CONTAINER_INHERIT_ACE	= 0x2
	NO_PROPAGATE_INHERIT_ACE	= 0x4
	INHERIT_ONLY_ACE	= 0x8
	INHERITED_ACE	= 0x10
	VALID_INHERIT_FLAGS	= 0x1F
)

type EsZbl5PW07Wm uint32


const (
	NO_MULTIPLE_TRUSTEE	= 0
	TRUSTEE_IS_IMPERSONATE	= 1
)

type FurMG3M0 uint32


const (
	TRUSTEE_IS_SID	= 0
	TRUSTEE_IS_NAME	= 1
	TRUSTEE_BAD_FORM	= 2
	TRUSTEE_IS_OBJECTS_AND_SID	= 3
	TRUSTEE_IS_OBJECTS_AND_NAME	= 4
)

type CobxXn5e_ uint32


const (
	TRUSTEE_IS_UNKNOWN	= 0
	TRUSTEE_IS_USER	= 1
	TRUSTEE_IS_GROUP	= 2
	TRUSTEE_IS_DOMAIN	= 3
	TRUSTEE_IS_ALIAS	= 4
	TRUSTEE_IS_WELL_KNOWN_GROUP	= 5
	TRUSTEE_IS_DELETED	= 6
	TRUSTEE_IS_INVALID	= 7
	TRUSTEE_IS_COMPUTER	= 8
)


const (
	ACE_OBJECT_TYPE_PRESENT	= 0x1
	ACE_INHERITED_OBJECT_TYPE_PRESENT	= 0x2
)

type GK5XcbVlLR struct {
	CblIKp	AlbfLJMLys5
	KCaEcaY	EAKnpNCN
	SlDUazFagbX	uint32
	TLUWloo	C3xy2ou
}


type BalX6cOQ_ uintptr

func IeL7F3m(xukLmcNaR string) BalX6cOQ_ {
	return  /*line :1*/ BalX6cOQ_( /*line :1*/ unsafe.Pointer( /*line :1*/ RRlduq3XI(xukLmcNaR)))
}
func PzPuOLZuq4(ohXw5VDcqx *He5NCq) BalX6cOQ_ {
	return  /*line :1*/ BalX6cOQ_( /*line :1*/ unsafe.Pointer(ohXw5VDcqx))
}
func J2l623y9xq(mSJEaB3n *SawrJJkh) BalX6cOQ_ {
	return  /*line :1*/ BalX6cOQ_( /*line :1*/ unsafe.Pointer(mSJEaB3n))
}
func JIYrBp(nxvXA89gGL3 *Cl_ZQmP) BalX6cOQ_ {
	return  /*line :1*/ BalX6cOQ_( /*line :1*/ unsafe.Pointer(nxvXA89gGL3))
}

type C3xy2ou struct {
	Nru1Cs4aYhsc	*C3xy2ou
	Fxgw_ZJ48	EsZbl5PW07Wm
	YbQWXJdkMV	FurMG3M0
	NjgwzIqB3Rzf	CobxXn5e_
	FJOSz9Geus	BalX6cOQ_
}

type SawrJJkh struct {
	PGVZTh3uoe	uint32
	Jl6sn3	Kw2qafRFaiLg
	HB7fJO	Kw2qafRFaiLg
	I1vFoRNz	*He5NCq
}

type Cl_ZQmP struct {
	O5rioZ7	uint32
	IeaMX54r	J4Q5ilM0dSeR
	B7xpM0J	*uint16
	Hk00iYva	*uint16
	HwQj03oAkk	*uint16
}


func (gRG9Cm *A2t_Jvj3ECv) Control() (rZjIu_h OjSqAgjIy, hKHZ38T275DC uint32, jeWMpOaQtWV error) {
	jeWMpOaQtWV =  /*line :1*/ lYFwBD(gRG9Cm, &rZjIu_h, &hKHZ38T275DC)
	return
}


func (gRG9Cm *A2t_Jvj3ECv) SetControl(fKPnmQ_w OjSqAgjIy, lJCNaye74 OjSqAgjIy) error {
	return  /*line :1*/ pg_dOfgc(gRG9Cm, fKPnmQ_w, lJCNaye74)
}


func (gRG9Cm *A2t_Jvj3ECv) RMControl() (rZjIu_h uint8, jeWMpOaQtWV error) {
	jeWMpOaQtWV =  /*line :1*/ zDSiag(gRG9Cm, &rZjIu_h)
	return
}


func (gRG9Cm *A2t_Jvj3ECv) SetRMControl(jsM1XMF uint8) {
	 /*line :1*/ cANESp5i1CR(gRG9Cm, &jsM1XMF)
}




func (gRG9Cm *A2t_Jvj3ECv) DACL() (qgpBZb *Na2cliTdMq, f00K_VYuoebB bool, jeWMpOaQtWV error) {
	var aeqXXeO bool
	jeWMpOaQtWV =  /*line :1*/ pTSboT(gRG9Cm, &aeqXXeO, &qgpBZb, &f00K_VYuoebB)
	if !aeqXXeO {
		jeWMpOaQtWV = ERROR_OBJECT_NOT_FOUND
	}
	return
}


func (u6IbXvI *A2t_Jvj3ECv) SetDACL(qgpBZb *Na2cliTdMq, aeqXXeO, f00K_VYuoebB bool) error {
	return  /*line :1*/ aGxMC3J(u6IbXvI, aeqXXeO, qgpBZb, f00K_VYuoebB)
}




func (gRG9Cm *A2t_Jvj3ECv) SACL() (lsVxy_e *Na2cliTdMq, f00K_VYuoebB bool, jeWMpOaQtWV error) {
	var aeqXXeO bool
	jeWMpOaQtWV =  /*line :1*/ i5Sl6luj(gRG9Cm, &aeqXXeO, &lsVxy_e, &f00K_VYuoebB)
	if !aeqXXeO {
		jeWMpOaQtWV = ERROR_OBJECT_NOT_FOUND
	}
	return
}


func (u6IbXvI *A2t_Jvj3ECv) SetSACL(lsVxy_e *Na2cliTdMq, aeqXXeO, f00K_VYuoebB bool) error {
	return  /*line :1*/ vCFgaK5rwa7(u6IbXvI, aeqXXeO, lsVxy_e, f00K_VYuoebB)
}


func (gRG9Cm *A2t_Jvj3ECv) Owner() (v2a8VMFGuf0X *He5NCq, f00K_VYuoebB bool, jeWMpOaQtWV error) {
	jeWMpOaQtWV =  /*line :1*/ f7MPHA2(gRG9Cm, &v2a8VMFGuf0X, &f00K_VYuoebB)
	return
}


func (u6IbXvI *A2t_Jvj3ECv) SetOwner(v2a8VMFGuf0X *He5NCq, f00K_VYuoebB bool) error {
	return  /*line :1*/ aTIded(u6IbXvI, v2a8VMFGuf0X, f00K_VYuoebB)
}


func (gRG9Cm *A2t_Jvj3ECv) Group() (bnHAoth *He5NCq, f00K_VYuoebB bool, jeWMpOaQtWV error) {
	jeWMpOaQtWV =  /*line :1*/ dnmkbYUfgau(gRG9Cm, &bnHAoth, &f00K_VYuoebB)
	return
}


func (u6IbXvI *A2t_Jvj3ECv) SetGroup(bnHAoth *He5NCq, f00K_VYuoebB bool) error {
	return  /*line :1*/ b04HqSxM3(u6IbXvI, bnHAoth, f00K_VYuoebB)
}


func (gRG9Cm *A2t_Jvj3ECv) Length() uint32 {
	return  /*line :1*/ mu616mhz(gRG9Cm)
}


func (gRG9Cm *A2t_Jvj3ECv) IsValid() bool {
	return  /*line :1*/ wKM4vpBQ(gRG9Cm)
}



func (gRG9Cm *A2t_Jvj3ECv) String() string {
	var n7WCaEw *uint16
	jeWMpOaQtWV :=  /*line :1*/ sDeD5NqS5U77(gRG9Cm, 1, 0xff, &n7WCaEw, nil)
	if jeWMpOaQtWV != nil {
		return ""
	}
	defer  /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(n7WCaEw)))
	return  /*line :1*/ M_Sea9j(n7WCaEw)
}


func (c6gWHx *A2t_Jvj3ECv) ToAbsolute() (u6IbXvI *A2t_Jvj3ECv, jeWMpOaQtWV error) {
	rZjIu_h, _, jeWMpOaQtWV :=  /*line :1*/ c6gWHx.Control()
	if jeWMpOaQtWV != nil {
		return
	}
	if rZjIu_h&SE_SELF_RELATIVE == 0 {
		jeWMpOaQtWV = ERROR_INVALID_PARAMETER
		return
	}
	var bOwhozIuAe_H, sgolvRhEI, gA267ZXVFy, eIUgV4FOV, miFdEgDW9Y uint32
	jeWMpOaQtWV =  /*line :1*/ xdmSfr(c6gWHx, nil, &bOwhozIuAe_H,
		nil, &sgolvRhEI, nil, &gA267ZXVFy, nil, &eIUgV4FOV, nil, &miFdEgDW9Y)
	switch jeWMpOaQtWV {
	case ERROR_INSUFFICIENT_BUFFER:
	case nil:

		return nil, ERROR_INTERNAL_ERROR
	default:
		return nil, jeWMpOaQtWV
	}
	if bOwhozIuAe_H > 0 {
		u6IbXvI = (* /*line :1*/ A2t_Jvj3ECv)( /*line :1*/ unsafe.Pointer(& /*line :1*/ make([]byte, bOwhozIuAe_H)[0]))
	}
	var (
		qgpBZb	*Na2cliTdMq
		lsVxy_e	*Na2cliTdMq
		v2a8VMFGuf0X	*He5NCq
		bnHAoth	*He5NCq
	)
	if sgolvRhEI > 0 {
		qgpBZb = (* /*line :1*/ Na2cliTdMq)( /*line :1*/ unsafe.Pointer(& /*line :1*/ make([]byte, sgolvRhEI)[0]))
	}
	if gA267ZXVFy > 0 {
		lsVxy_e = (* /*line :1*/ Na2cliTdMq)( /*line :1*/ unsafe.Pointer(& /*line :1*/ make([]byte, gA267ZXVFy)[0]))
	}
	if eIUgV4FOV > 0 {
		v2a8VMFGuf0X = (* /*line :1*/ He5NCq)( /*line :1*/ unsafe.Pointer(& /*line :1*/ make([]byte, eIUgV4FOV)[0]))
	}
	if miFdEgDW9Y > 0 {
		bnHAoth = (* /*line :1*/ He5NCq)( /*line :1*/ unsafe.Pointer(& /*line :1*/ make([]byte, miFdEgDW9Y)[0]))
	}
	jeWMpOaQtWV =  /*line :1*/ xdmSfr(c6gWHx, u6IbXvI, &bOwhozIuAe_H,
		qgpBZb, &sgolvRhEI, lsVxy_e, &gA267ZXVFy, v2a8VMFGuf0X, &eIUgV4FOV, bnHAoth, &miFdEgDW9Y)
	return
}


func (u6IbXvI *A2t_Jvj3ECv) ToSelfRelative() (c6gWHx *A2t_Jvj3ECv, jeWMpOaQtWV error) {
	rZjIu_h, _, jeWMpOaQtWV :=  /*line :1*/ u6IbXvI.Control()
	if jeWMpOaQtWV != nil {
		return
	}
	if rZjIu_h&SE_SELF_RELATIVE != 0 {
		jeWMpOaQtWV = ERROR_INVALID_PARAMETER
		return
	}
	var oo6I0d uint32
	jeWMpOaQtWV =  /*line :1*/ pLuyKb(u6IbXvI, nil, &oo6I0d)
	switch jeWMpOaQtWV {
	case ERROR_INSUFFICIENT_BUFFER:
	case nil:

		return nil, ERROR_INTERNAL_ERROR
	default:
		return nil, jeWMpOaQtWV
	}
	if oo6I0d > 0 {
		c6gWHx = (* /*line :1*/ A2t_Jvj3ECv)( /*line :1*/ unsafe.Pointer(& /*line :1*/ make([]byte, oo6I0d)[0]))
	}
	jeWMpOaQtWV =  /*line :1*/ pLuyKb(u6IbXvI, c6gWHx, &oo6I0d)
	return
}

func (c6gWHx *A2t_Jvj3ECv) l8kI5b0LaT() *A2t_Jvj3ECv {
	wGbK2mjjiAz :=  /*line :1*/ int( /*line :1*/ c6gWHx.Length())
	const min =  /*line :1*/ int( /*line :1*/ unsafe.Sizeof(A2t_Jvj3ECv{}))
	if wGbK2mjjiAz < min {
		wGbK2mjjiAz = min
	}

	huthvVhTPAIF :=  /*line :1*/ unsafe.Slice((* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(c6gWHx)), wGbK2mjjiAz)
	
	
	
	
	const psize =  /*line :1*/ int( /*line :1*/ unsafe.Sizeof( /*line :1*/ uintptr(0)))
	wfJ2EhPna4 :=  /*line :1*/ make([]uintptr, (wGbK2mjjiAz+psize-1)/psize)
	z0twykel :=  /*line :1*/ unsafe.Slice((* /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&wfJ2EhPna4[0])), wGbK2mjjiAz)
	 /*line :1*/ copy(z0twykel, huthvVhTPAIF)
	return (* /*line :1*/ A2t_Jvj3ECv)( /*line :1*/ unsafe.Pointer(&z0twykel[0]))
}



func OQbD3tB9Jujo(n7WCaEw string) (gRG9Cm *A2t_Jvj3ECv, jeWMpOaQtWV error) {
	var fXJc5c2oP_y7 *A2t_Jvj3ECv
	jeWMpOaQtWV =  /*line :1*/ jkx0iugRRqC(n7WCaEw, 1, &fXJc5c2oP_y7, nil)
	if jeWMpOaQtWV != nil {
		return
	}
	defer  /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(fXJc5c2oP_y7)))
	return  /*line :1*/ fXJc5c2oP_y7.l8kI5b0LaT(), nil
}



func L2RJlhSMA(iOvctVD26lA L2L8P9WaNZ, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D) (gRG9Cm *A2t_Jvj3ECv, jeWMpOaQtWV error) {
	var fXJc5c2oP_y7 *A2t_Jvj3ECv
	jeWMpOaQtWV =  /*line :1*/ xdBCuREUs9h(iOvctVD26lA, mf05G8fOkrTC, d1ypxqWR48, nil, nil, nil, nil, &fXJc5c2oP_y7)
	if jeWMpOaQtWV != nil {
		return
	}
	defer  /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(fXJc5c2oP_y7)))
	return  /*line :1*/ fXJc5c2oP_y7.l8kI5b0LaT(), nil
}



func ALda24g(x_GAOsRtS480 string, mf05G8fOkrTC J4Q5ilM0dSeR, d1ypxqWR48 B4w2t8D) (gRG9Cm *A2t_Jvj3ECv, jeWMpOaQtWV error) {
	var fXJc5c2oP_y7 *A2t_Jvj3ECv
	jeWMpOaQtWV =  /*line :1*/ brDxO5WbHd3O(x_GAOsRtS480, mf05G8fOkrTC, d1ypxqWR48, nil, nil, nil, nil, &fXJc5c2oP_y7)
	if jeWMpOaQtWV != nil {
		return
	}
	defer  /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(fXJc5c2oP_y7)))
	return  /*line :1*/ fXJc5c2oP_y7.l8kI5b0LaT(), nil
}




func AUOuPLx(v2a8VMFGuf0X *C3xy2ou, bnHAoth *C3xy2ou, i4CY63_ []GK5XcbVlLR, nGgbVZtbG []GK5XcbVlLR, cyN2qJdawc3x *A2t_Jvj3ECv) (gRG9Cm *A2t_Jvj3ECv, jeWMpOaQtWV error) {
	var fXJc5c2oP_y7 *A2t_Jvj3ECv
	var pRkiWsXV6Xh uint32
	var vGDkv5b *GK5XcbVlLR
	if  /*line :1*/ len(i4CY63_) > 0 {
		vGDkv5b = &i4CY63_[0]
	}
	var fABjILfbbQ *GK5XcbVlLR
	if  /*line :1*/ len(nGgbVZtbG) > 0 {
		fABjILfbbQ = &nGgbVZtbG[0]
	}
	jeWMpOaQtWV =  /*line :1*/ g0g0QrFhHwNX(v2a8VMFGuf0X, bnHAoth,  /*line :1*/ uint32( /*line :1*/ len(i4CY63_)), vGDkv5b,  /*line :1*/ uint32( /*line :1*/ len(nGgbVZtbG)), fABjILfbbQ, cyN2qJdawc3x, &pRkiWsXV6Xh, &fXJc5c2oP_y7)
	if jeWMpOaQtWV != nil {
		return
	}
	defer  /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(fXJc5c2oP_y7)))
	return  /*line :1*/ fXJc5c2oP_y7.l8kI5b0LaT(), nil
}


func TbeSc5KyTix() (u6IbXvI *A2t_Jvj3ECv, jeWMpOaQtWV error) {
	u6IbXvI = &A2t_Jvj3ECv{}
	jeWMpOaQtWV =  /*line :1*/ eyd5z2RaLaA(u6IbXvI, 1)
	return
}



func Tveplz2mDj(siNJ3lw3z []GK5XcbVlLR, nxtZhWD5BjTO *Na2cliTdMq) (hlgZJWcfkz *Na2cliTdMq, jeWMpOaQtWV error) {
	var _azsGy *GK5XcbVlLR
	if  /*line :1*/ len(siNJ3lw3z) > 0 {
		_azsGy = &siNJ3lw3z[0]
	}
	var olYpzbVdwqH *Na2cliTdMq
	jeWMpOaQtWV =  /*line :1*/ qhb0Czy( /*line :1*/ uint32( /*line :1*/ len(siNJ3lw3z)), _azsGy, nxtZhWD5BjTO, &olYpzbVdwqH)
	if jeWMpOaQtWV != nil {
		return
	}
	defer  /*line :1*/ OtMI6tu9( /*line :1*/ L2L8P9WaNZ( /*line :1*/ unsafe.Pointer(olYpzbVdwqH)))
	hcNoYC :=  /*line :1*/ make([]byte, olYpzbVdwqH.svUIDH8yUWzh)
	 /*line :1*/ copy(hcNoYC, (*[(1 << 31) - 1] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(olYpzbVdwqH))[: /*line :1*/ len(hcNoYC): /*line :1*/ len(hcNoYC)])
	return (* /*line :1*/ Na2cliTdMq)( /*line :1*/ unsafe.Pointer(&hcNoYC[0])), nil
}
