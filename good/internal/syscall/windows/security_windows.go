//line :1
package LdptURlN

import (
	syscall "bUKeamF"
	"unsafe"
)

const (
	SecurityAnonymous	= 0
	SecurityIdentification	= 1
	SecurityImpersonation	= 2
	SecurityDelegation	= 3
)

const (
	TOKEN_ADJUST_PRIVILEGES	= 0x0020
	SE_PRIVILEGE_ENABLED	= 0x00000002
)

type PkWXUB7wm struct {
	Cwsdj6SDrP	uint32
	Hfe4CvtpDz	int32
}

type G9aRKwS struct {
	Lu71cUl2CrD7	PkWXUB7wm
	XFNr8D9BA0	uint32
}

type LFo8HyR struct {
	Jw83pW7Y	uint32
	NMhrtH	[1]G9aRKwS
}

func Ja498s5(hp67APc637J3 syscall.Ad4bWd, hls6MaIN bool, eqAC0x1d *LFo8HyR, xUtG63NsJD uint32, l7Zkpa1NV6 *LFo8HyR, hkHgRkM7M *uint32) error {
	jxAixT, zc4PmxS :=  /*line :1*/ ph35lMCuC(hp67APc637J3, hls6MaIN, eqAC0x1d, xUtG63NsJD, l7Zkpa1NV6, hkHgRkM7M)
	if jxAixT == 0 {

		return zc4PmxS
	}

	if zc4PmxS == syscall.EINVAL {

		return nil
	}
	return zc4PmxS
}

type E1oO8m struct {
	SkbAOXRTj	*syscall.CLoH6Pka1ufN
	HaWYiU3h	uint32
}

type XFemI6aqvJ6 struct {
	HHnDAqV E1oO8m
}

func (njltE8YAMeIr *XFemI6aqvJ6) Size() uint32 {
	return  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(XFemI6aqvJ6{})) +  /*line :1*/ syscall.BFQLdGf9_(njltE8YAMeIr.HHnDAqV.SkbAOXRTj)
}

const SE_GROUP_INTEGRITY = 0x00000020

type Zs2_lZoOk uint32

const (
	TokenPrimary	Zs2_lZoOk	= 1
	TokenImpersonation	Zs2_lZoOk	= 2
)

const (
	LG_INCLUDE_INDIRECT	= 0x1
	MAX_PREFERRED_LENGTH	= 0xFFFFFFFF
)

type L9xDvio struct {
	Bgi0QY *uint16
}

type JgQLqpez1eeV struct {
	Ej5oi9rINSlA	*uint16
	Qd1K5zJyimQ	*uint16
	FygXzeAU	uint32
	IqPrSAi3	uint32
	Ev0XwQG7ck	*uint16
	B_aCfOk	*uint16
	QfWTjy2	uint32
	ZkzBbB_	*uint16
	P8oWFeT3rA	uint32
	K1QUJUayjF	*uint16
	Sgf4kIcCD4	*uint16
	GCF6JAxCi0	*uint16
	MzmYkR2xj	*uint16
	KI7Zta6XDPHj	uint32
	EwyY8fqBetNo	uint32
	WMEBk1ZrJTLU	uint32
	XYbPMPNI1A	uint32
	HVYO0C	uint32
	A3wvispR9m	*byte
	BriBdeEd	uint32
	Fy18XW	uint32
	Fcwm_Qc8fVam	*uint16
	UEuKCE65VW	uint32
	JDaIzoQWjaR	uint32
	UH_2s5Ei3	*syscall.CLoH6Pka1ufN
	ICYgjkBpl	uint32
	BYIRA1C	*uint16
	AVA0QJ_TruW	*uint16
	OZsXWm	uint32
}
