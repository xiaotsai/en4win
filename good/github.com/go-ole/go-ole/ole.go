//line :1
package fuA83L

import (
	fmt "kFsoCfy5zWG"
	strings "fQXlzv"
)


type QnAeLz struct {
	fkpLunsR2d	uintptr
	yt9U6V6J9x	uintptr
	ioCkv9wA	uint32
	tz20EQ	uint32
}


type I0iHb3gLxw struct {
	gzhAh1qS_B	uint16
	bBWcOs2mN	uint16
	w6FQb7n8h8g	*uint16
	jE1Y4UGIk	*uint16
	c1MYTaGIn	*uint16
	hm4AMnV	uint32
	nRukPhzbN9r	uintptr
	aQw229p	uintptr
	dn8jo1	uint32
}


func (wDb3AREI_6N1 I0iHb3gLxw) WCode() uint16 {
	return wDb3AREI_6N1.gzhAh1qS_B
}


func (wDb3AREI_6N1 I0iHb3gLxw) SCODE() uint32 {
	return wDb3AREI_6N1.dn8jo1
}


func (wDb3AREI_6N1 I0iHb3gLxw) String() string {
	var gA9YADk, cnRR6WPe, gfr1XMDO string
	if wDb3AREI_6N1.w6FQb7n8h8g == nil {
		gA9YADk = "<nil>"
	} else {
		gA9YADk =  /*line :1*/ Jp_D7PByTsS(wDb3AREI_6N1.w6FQb7n8h8g)
	}

	if wDb3AREI_6N1.jE1Y4UGIk == nil {
		cnRR6WPe = "<nil>"
	} else {
		cnRR6WPe =  /*line :1*/ Jp_D7PByTsS(wDb3AREI_6N1.jE1Y4UGIk)
	}

	if wDb3AREI_6N1.c1MYTaGIn == nil {
		gfr1XMDO = "<nil>"
	} else {
		gfr1XMDO =  /*line :1*/ Jp_D7PByTsS(wDb3AREI_6N1.c1MYTaGIn)
	}

	return  /*line :1*/ fmt.IBsS3Oc(
		"wCode: %#x, bstrSource: %v, bstrDescription: %v, bstrHelpFile: %v, dwHelpContext: %#x, scode: %#x",
		wDb3AREI_6N1.gzhAh1qS_B, gA9YADk, cnRR6WPe, gfr1XMDO, wDb3AREI_6N1.hm4AMnV, wDb3AREI_6N1.dn8jo1,
	)
}


func (wDb3AREI_6N1 I0iHb3gLxw) Error() string {
	if wDb3AREI_6N1.jE1Y4UGIk != nil {
		return  /*line :1*/ strings.H8cdur8LCY( /*line :1*/ Jp_D7PByTsS(wDb3AREI_6N1.jE1Y4UGIk))
	}

	gA9YADk := "Unknown"
	if wDb3AREI_6N1.w6FQb7n8h8g != nil {
		gA9YADk =  /*line :1*/ Jp_D7PByTsS(wDb3AREI_6N1.w6FQb7n8h8g)
	}

	ht0fK5E1bpa := wDb3AREI_6N1.dn8jo1
	if wDb3AREI_6N1.gzhAh1qS_B != 0 {
		ht0fK5E1bpa =  /*line :1*/ uint32(wDb3AREI_6N1.gzhAh1qS_B)
	}

	return  /*line :1*/ fmt.IBsS3Oc("%v: %#x", gA9YADk, ht0fK5E1bpa)
}


type Q8jTq1qN struct {
	NeBEoZyD	*int16
	Gmj97mWFa	uint16
}


type U6qdBOqcOE struct {
	HZBLFp	*uint16
	QPVx1CaW	*Q8jTq1qN
	D3tr3o50e8m6	int32
	SIoNgye0w2j	uint32
	BCmL6PBV	int32
	Go2MuUM	uint32
	QZU_ic56q	uint16
	UsigoLgho4	uint32
}


type AmPWIefN struct {
	NSG4Yga0	*U6qdBOqcOE
	C0me5XqEjXU	uint32
}


type JqarRnq4 struct {
	GIpBWH_	int32
	HL4vcOhA	int32
}


type WeDtaklL struct {
	DfJVqNg9uY	uint32
	YK_4w6QalsY	uint32
	V5ivWDQ	int32
	B_6t8o	int32
	O5CXsDATs	uint32
	Be2I2Zy	JqarRnq4
}


type Kg1aaCax6iKa struct {
	WRSSyq	uint32
	Y8Qv9JFRq	uint16
}


type LrWQg36xG struct {
	L_HWlDM	uint32
	DzyRgQW0	uint16
}


type CFaF1Ke struct {
	BQEmv0J	GUID
	Ym1adr7xKDw2	uint32
	bUuwi44_Nzty	uint32
	ASCh6Z	int32
	N0CEGt	int32
	Lg1IE39h6	*uint16
	Tx_ZHC9aE	uint32
	JIAze2KC	int32
	JoWWhqSMa	uint16
	FXSpaP0P_	uint16
	DvECAEn5mzy	uint16
	Fylzeml0u	uint16
	ZZyaCrTWllLp	uint16
	XJeDQG5_Ah	uint16
	O4wZeJL2	uint16
	GVPH5ws1U	uint16
	NaG547HnTR	Kg1aaCax6iKa
	AMEBzUexKj	LrWQg36xG
}
