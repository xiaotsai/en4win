//line :1








package FeKim_cPgM7z

import (
	errors "iAHoxjmM"
)


type Ze6KllqJw uint16

const (
	
	TypeA	Ze6KllqJw	= 1
	TypeNS	Ze6KllqJw	= 2
	TypeCNAME	Ze6KllqJw	= 5
	TypeSOA	Ze6KllqJw	= 6
	TypePTR	Ze6KllqJw	= 12
	TypeMX	Ze6KllqJw	= 15
	TypeTXT	Ze6KllqJw	= 16
	TypeAAAA	Ze6KllqJw	= 28
	TypeSRV	Ze6KllqJw	= 33
	TypeOPT	Ze6KllqJw	= 41

	
	TypeWKS	Ze6KllqJw	= 11
	TypeHINFO	Ze6KllqJw	= 13
	TypeMINFO	Ze6KllqJw	= 14
	TypeAXFR	Ze6KllqJw	= 252
	TypeALL	Ze6KllqJw	= 255
)

var dW27Z0b6yq = map[Ze6KllqJw]string{
	TypeA:	"TypeA",
	TypeNS:	"TypeNS",
	TypeCNAME:	"TypeCNAME",
	TypeSOA:	"TypeSOA",
	TypePTR:	"TypePTR",
	TypeMX:	"TypeMX",
	TypeTXT:	"TypeTXT",
	TypeAAAA:	"TypeAAAA",
	TypeSRV:	"TypeSRV",
	TypeOPT:	"TypeOPT",
	TypeWKS:	"TypeWKS",
	TypeHINFO:	"TypeHINFO",
	TypeMINFO:	"TypeMINFO",
	TypeAXFR:	"TypeAXFR",
	TypeALL:	"TypeALL",
}


func (qavq6q Ze6KllqJw) String() string {
	if aZuE3TFPdpn, afB2jgY := dW27Z0b6yq[qavq6q]; afB2jgY {
		return aZuE3TFPdpn
	}
	return  /*line :1*/ d20BK8p4Fr( /*line :1*/ uint16(qavq6q))
}


func (qavq6q Ze6KllqJw) GoString() string {
	if aZuE3TFPdpn, afB2jgY := dW27Z0b6yq[qavq6q]; afB2jgY {
		return "dnsmessage." + aZuE3TFPdpn
	}
	return  /*line :1*/ d20BK8p4Fr( /*line :1*/ uint16(qavq6q))
}


type WhBR6DRkZ3p uint16

const (
	
	ClassINET	WhBR6DRkZ3p	= 1
	ClassCSNET	WhBR6DRkZ3p	= 2
	ClassCHAOS	WhBR6DRkZ3p	= 3
	ClassHESIOD	WhBR6DRkZ3p	= 4

	
	ClassANY	WhBR6DRkZ3p	= 255
)

var q1_14jdaQbB = map[WhBR6DRkZ3p]string{
	ClassINET:	"ClassINET",
	ClassCSNET:	"ClassCSNET",
	ClassCHAOS:	"ClassCHAOS",
	ClassHESIOD:	"ClassHESIOD",
	ClassANY:	"ClassANY",
}


func (iiYnULwuUYX WhBR6DRkZ3p) String() string {
	if aZuE3TFPdpn, afB2jgY := q1_14jdaQbB[iiYnULwuUYX]; afB2jgY {
		return aZuE3TFPdpn
	}
	return  /*line :1*/ d20BK8p4Fr( /*line :1*/ uint16(iiYnULwuUYX))
}


func (iiYnULwuUYX WhBR6DRkZ3p) GoString() string {
	if aZuE3TFPdpn, afB2jgY := q1_14jdaQbB[iiYnULwuUYX]; afB2jgY {
		return "dnsmessage." + aZuE3TFPdpn
	}
	return  /*line :1*/ d20BK8p4Fr( /*line :1*/ uint16(iiYnULwuUYX))
}


type AtcaOBVO8X uint16


func (rL_9YGYQqnf AtcaOBVO8X) GoString() string {
	return  /*line :1*/ d20BK8p4Fr( /*line :1*/ uint16(rL_9YGYQqnf))
}


type QPDSMamH uint16


const (
	RCodeSuccess	QPDSMamH	= 0	
	RCodeFormatError	QPDSMamH	= 1	
	RCodeServerFailure	QPDSMamH	= 2	
	RCodeNameError	QPDSMamH	= 3	
	RCodeNotImplemented	QPDSMamH	= 4	
	RCodeRefused	QPDSMamH	= 5	
)

var ihVfLjXr4g = map[QPDSMamH]string{
	RCodeSuccess:	"RCodeSuccess",
	RCodeFormatError:	"RCodeFormatError",
	RCodeServerFailure:	"RCodeServerFailure",
	RCodeNameError:	"RCodeNameError",
	RCodeNotImplemented:	"RCodeNotImplemented",
	RCodeRefused:	"RCodeRefused",
}


func (y6pzOXbIFd QPDSMamH) String() string {
	if aZuE3TFPdpn, afB2jgY := ihVfLjXr4g[y6pzOXbIFd]; afB2jgY {
		return aZuE3TFPdpn
	}
	return  /*line :1*/ d20BK8p4Fr( /*line :1*/ uint16(y6pzOXbIFd))
}


func (y6pzOXbIFd QPDSMamH) GoString() string {
	if aZuE3TFPdpn, afB2jgY := ihVfLjXr4g[y6pzOXbIFd]; afB2jgY {
		return "dnsmessage." + aZuE3TFPdpn
	}
	return  /*line :1*/ d20BK8p4Fr( /*line :1*/ uint16(y6pzOXbIFd))
}

func vxqhba2b(eG5npDpL0 uint8) string {
	uUTUcsSMYDR :=  /*line :1*/ byte(eG5npDpL0)
	return  /*line :1*/ string([]byte{
		uUTUcsSMYDR/100 + '0',
		uUTUcsSMYDR/10%10 + '0',
		uUTUcsSMYDR%10 + '0',
	})
}

func z6a43j(pU8PG4B []byte, eG5npDpL0 uint8) []byte {
	uUTUcsSMYDR :=  /*line :1*/ byte(eG5npDpL0)
	if eG5npDpL0 >= 100 {
		pU8PG4B =  /*line :1*/ append(pU8PG4B, uUTUcsSMYDR/100+'0')
	}
	if eG5npDpL0 >= 10 {
		pU8PG4B =  /*line :1*/ append(pU8PG4B, uUTUcsSMYDR/10%10+'0')
	}
	return  /*line :1*/ append(pU8PG4B, uUTUcsSMYDR%10+'0')
}

func eguZ3sWx8(uUTUcsSMYDR []byte) string {
	if  /*line :1*/ len(uUTUcsSMYDR) == 0 {
		return ""
	}
	pU8PG4B :=  /*line :1*/ make([]byte, 0, 5* /*line :1*/ len(uUTUcsSMYDR))
	pU8PG4B =  /*line :1*/ z6a43j(pU8PG4B,  /*line :1*/ uint8(uUTUcsSMYDR[0]))
	for _, aZuE3TFPdpn := range uUTUcsSMYDR[1:] {
		pU8PG4B =  /*line :1*/ append(pU8PG4B, ',', ' ')
		pU8PG4B =  /*line :1*/ z6a43j(pU8PG4B,  /*line :1*/ uint8(aZuE3TFPdpn))
	}
	return  /*line :1*/ string(pU8PG4B)
}

const hexDigits = "0123456789abcdef"

func tUbBJah(ao9NXO []byte) string {
	pU8PG4B :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(ao9NXO))
	for eG5npDpL0 := 0; eG5npDpL0 <  /*line :1*/ len(ao9NXO); eG5npDpL0++ {
		iiYnULwuUYX := ao9NXO[eG5npDpL0]
		if iiYnULwuUYX == '.' || iiYnULwuUYX == '-' || iiYnULwuUYX == ' ' ||
			'A' <= iiYnULwuUYX && iiYnULwuUYX <= 'Z' ||
			'a' <= iiYnULwuUYX && iiYnULwuUYX <= 'z' ||
			'0' <= iiYnULwuUYX && iiYnULwuUYX <= '9' {
			pU8PG4B =  /*line :1*/ append(pU8PG4B, iiYnULwuUYX)
			continue
		}

		imjtSXod := iiYnULwuUYX >> 4
		dSKKxzKPhnoq := (iiYnULwuUYX << 4) >> 4
		pU8PG4B =  /*line :1*/ append(
			pU8PG4B,
			'\\',
			'x',
			hexDigits[imjtSXod],
			hexDigits[dSKKxzKPhnoq],
		)
	}
	return  /*line :1*/ string(pU8PG4B)
}

func d20BK8p4Fr(eG5npDpL0 uint16) string {
	return  /*line :1*/ kzyBd0GDt( /*line :1*/ uint32(eG5npDpL0))
}

func kzyBd0GDt(eG5npDpL0 uint32) string {

	pU8PG4B :=  /*line :1*/ make([]byte, 10)
	for uUTUcsSMYDR, yf2424rNhg := pU8PG4B,  /*line :1*/ uint32(1000000000); yf2424rNhg > 0; yf2424rNhg /= 10 {
		uUTUcsSMYDR[0] =  /*line :1*/ byte(eG5npDpL0/yf2424rNhg%10 + '0')
		if uUTUcsSMYDR[0] == '0' &&  /*line :1*/ len(uUTUcsSMYDR) ==  /*line :1*/ len(pU8PG4B) &&  /*line :1*/ len(pU8PG4B) > 1 {
			pU8PG4B = pU8PG4B[1:]
		}
		uUTUcsSMYDR = uUTUcsSMYDR[1:]
		eG5npDpL0 %= yf2424rNhg
	}
	return  /*line :1*/ string(pU8PG4B)
}

func _HP6S72n(uUTUcsSMYDR bool) string {
	if uUTUcsSMYDR {
		return "true"
	}
	return "false"
}

var (
	
	
	
	WbM9aD3mRP	=  /*line :1*/ errors.Su6g6hRoi9X("parsing/packing of this type isn't available yet")

	
	
	BbQhjXWx	=  /*line :1*/ errors.Su6g6hRoi9X("parsing/packing of this section has completed")

	yf5I2II0zMg	=  /*line :1*/ errors.Su6g6hRoi9X("insufficient data for base length type")
	onAp5pO	=  /*line :1*/ errors.Su6g6hRoi9X("insufficient data for calculated length type")
	t7KOKcour	=  /*line :1*/ errors.Su6g6hRoi9X("segment prefix is reserved")
	rZYKeUXMwZ	=  /*line :1*/ errors.Su6g6hRoi9X("too many pointers (>10)")
	t0_XfM	=  /*line :1*/ errors.Su6g6hRoi9X("invalid pointer")
	jx390dmJG	=  /*line :1*/ errors.Su6g6hRoi9X("invalid dns name")
	buZl0ZOk	=  /*line :1*/ errors.Su6g6hRoi9X("nil resource body")
	eCOkG0CgXc	=  /*line :1*/ errors.Su6g6hRoi9X("insufficient data for resource body length")
	qF0KX1	=  /*line :1*/ errors.Su6g6hRoi9X("segment length too long")
	iF0RY15sidD	=  /*line :1*/ errors.Su6g6hRoi9X("name too long")
	u07BDH	=  /*line :1*/ errors.Su6g6hRoi9X("zero length segment")
	vvoBDOiQz	=  /*line :1*/ errors.Su6g6hRoi9X("resource length too long")
	cwapDN1O	=  /*line :1*/ errors.Su6g6hRoi9X("too many Questions to pack (>65535)")
	j6t3f1Ay28	=  /*line :1*/ errors.Su6g6hRoi9X("too many Answers to pack (>65535)")
	aq6mWMTC	=  /*line :1*/ errors.Su6g6hRoi9X("too many Authorities to pack (>65535)")
	mDKEFxIt3Z	=  /*line :1*/ errors.Su6g6hRoi9X("too many Additionals to pack (>65535)")
	oNVzpKip	=  /*line :1*/ errors.Su6g6hRoi9X("name is not in canonical format (it must end with a .)")
	mRcqHJN	=  /*line :1*/ errors.Su6g6hRoi9X("character string exceeds maximum length (255)")
	mGz62tBsms9M	=  /*line :1*/ errors.Su6g6hRoi9X("compressed name in SRV resource data")
)


const (
	
	
	
	
	
	packStartingCap	= 512

	
	uint16Len	= 2

	
	uint32Len	= 4

	
	
	
	headerLen	= 6 * uint16Len
)

type gxaf3H1N struct {
	
	gqjn1k	string

	
	i0MshC_ZtWw2	error
}


func (xtXprr571m *gxaf3H1N) Error() string {
	return xtXprr571m.gqjn1k + ": " +  /*line :1*/ xtXprr571m.i0MshC_ZtWw2.Error()
}


type HlVtnUC struct {
	Xk_bZKgX	uint16
	FlCGSzy1U7Gj	bool
	OFaFRqSgwkpi	AtcaOBVO8X
	AgqoiYIElD	bool
	LeyP1H	bool
	HXV_4sM2D	bool
	Xu3NFYuAq	bool
	Ah7I1_6	bool
	C1cExuMlVHz	bool
	DAtOvNfx7	QPDSMamH
}

func (qPTmameBxee0 *HlVtnUC) qalrXmt() (n4vDsumpcebZ uint16, p4ytahFmg3 uint16) {
	n4vDsumpcebZ = qPTmameBxee0.Xk_bZKgX
	p4ytahFmg3 =  /*line :1*/ uint16(qPTmameBxee0.OFaFRqSgwkpi)<<11 |  /*line :1*/ uint16(qPTmameBxee0.DAtOvNfx7)
	if qPTmameBxee0.Xu3NFYuAq {
		p4ytahFmg3 |= headerBitRA
	}
	if qPTmameBxee0.HXV_4sM2D {
		p4ytahFmg3 |= headerBitRD
	}
	if qPTmameBxee0.LeyP1H {
		p4ytahFmg3 |= headerBitTC
	}
	if qPTmameBxee0.AgqoiYIElD {
		p4ytahFmg3 |= headerBitAA
	}
	if qPTmameBxee0.FlCGSzy1U7Gj {
		p4ytahFmg3 |= headerBitQR
	}
	if qPTmameBxee0.Ah7I1_6 {
		p4ytahFmg3 |= headerBitAD
	}
	if qPTmameBxee0.C1cExuMlVHz {
		p4ytahFmg3 |= headerBitCD
	}
	return
}


func (qPTmameBxee0 *HlVtnUC) GoString() string {
	return "dnsmessage.Header{" +
		"ID: " +  /*line :1*/ d20BK8p4Fr(qPTmameBxee0.Xk_bZKgX) + ", " +
		"Response: " +  /*line :1*/ _HP6S72n(qPTmameBxee0.FlCGSzy1U7Gj) + ", " +
		"OpCode: " +  /*line :1*/ qPTmameBxee0.OFaFRqSgwkpi.GoString() + ", " +
		"Authoritative: " +  /*line :1*/ _HP6S72n(qPTmameBxee0.AgqoiYIElD) + ", " +
		"Truncated: " +  /*line :1*/ _HP6S72n(qPTmameBxee0.LeyP1H) + ", " +
		"RecursionDesired: " +  /*line :1*/ _HP6S72n(qPTmameBxee0.HXV_4sM2D) + ", " +
		"RecursionAvailable: " +  /*line :1*/ _HP6S72n(qPTmameBxee0.Xu3NFYuAq) + ", " +
		"RCode: " +  /*line :1*/ qPTmameBxee0.DAtOvNfx7.GoString() + "}"
}


type P_fQ6x struct {
	HlVtnUC
	Il4Uso_ecr	[]LNQ5aNrCORB
	JiHd3TLzo	[]Y21iaz2cH_hd
	ZiuFEUcXZwi	[]Y21iaz2cH_hd
	W5XWEcd	[]Y21iaz2cH_hd
}

type igQ2ZFMu uint8

const (
	sectionNotStarted	igQ2ZFMu	= iota
	sectionHeader
	sectionQuestions
	sectionAnswers
	sectionAuthorities
	sectionAdditionals
	sectionDone

	headerBitQR	= 1 << 15		
	headerBitAA	= 1 << 10		
	headerBitTC	= 1 << 9		
	headerBitRD	= 1 << 8		
	headerBitRA	= 1 << 7		
	headerBitAD	= 1 << 5		
	headerBitCD	= 1 << 4		
)

var hfLZTsMtk = map[igQ2ZFMu]string{
	sectionHeader:	"header",
	sectionQuestions:	"Question",
	sectionAnswers:	"Answer",
	sectionAuthorities:	"Authority",
	sectionAdditionals:	"Additional",
}


type cNNn2uM struct {
	sYp49l6kf	uint16
	mrl9diwhBeyV	uint16
	sCqv4utCjzi	uint16
	j1sYZuJo	uint16
	jHCad4ja2_	uint16
	lLb2iA7OmE	uint16
}

func (iDU4uPqI *cNNn2uM) ctLSIZuX6Xq(gHo332FMv igQ2ZFMu) uint16 {
	switch gHo332FMv {
	case sectionQuestions:
		return iDU4uPqI.sCqv4utCjzi
	case sectionAnswers:
		return iDU4uPqI.j1sYZuJo
	case sectionAuthorities:
		return iDU4uPqI.jHCad4ja2_
	case sectionAdditionals:
		return iDU4uPqI.lLb2iA7OmE
	}
	return 0
}


func (iDU4uPqI *cNNn2uM) qalrXmt(bnTR5Q []byte) []byte {
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, iDU4uPqI.sYp49l6kf)
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, iDU4uPqI.mrl9diwhBeyV)
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, iDU4uPqI.sCqv4utCjzi)
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, iDU4uPqI.j1sYZuJo)
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, iDU4uPqI.jHCad4ja2_)
	return  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, iDU4uPqI.lLb2iA7OmE)
}

func (iDU4uPqI *cNNn2uM) gXjxaL(bnTR5Q []byte, sG72__8 int) (int, error) {
	xrV0mH := sG72__8
	var gX34Xg error
	if iDU4uPqI.sYp49l6kf, xrV0mH, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"id", gX34Xg}
	}
	if iDU4uPqI.mrl9diwhBeyV, xrV0mH, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"bits", gX34Xg}
	}
	if iDU4uPqI.sCqv4utCjzi, xrV0mH, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"questions", gX34Xg}
	}
	if iDU4uPqI.j1sYZuJo, xrV0mH, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"answers", gX34Xg}
	}
	if iDU4uPqI.jHCad4ja2_, xrV0mH, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"authorities", gX34Xg}
	}
	if iDU4uPqI.lLb2iA7OmE, xrV0mH, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"additionals", gX34Xg}
	}
	return xrV0mH, nil
}

func (iDU4uPqI *cNNn2uM) cNNn2uM() HlVtnUC {
	return HlVtnUC{
		Xk_bZKgX:	iDU4uPqI.sYp49l6kf,
		FlCGSzy1U7Gj:	(iDU4uPqI.mrl9diwhBeyV & headerBitQR) != 0,
		OFaFRqSgwkpi:	 /*line :1*/ AtcaOBVO8X(iDU4uPqI.mrl9diwhBeyV>>11) & 0xF,
		AgqoiYIElD:	(iDU4uPqI.mrl9diwhBeyV & headerBitAA) != 0,
		LeyP1H:	(iDU4uPqI.mrl9diwhBeyV & headerBitTC) != 0,
		HXV_4sM2D:	(iDU4uPqI.mrl9diwhBeyV & headerBitRD) != 0,
		Xu3NFYuAq:	(iDU4uPqI.mrl9diwhBeyV & headerBitRA) != 0,
		Ah7I1_6:	(iDU4uPqI.mrl9diwhBeyV & headerBitAD) != 0,
		C1cExuMlVHz:	(iDU4uPqI.mrl9diwhBeyV & headerBitCD) != 0,
		DAtOvNfx7:	 /*line :1*/ QPDSMamH(iDU4uPqI.mrl9diwhBeyV & 0xF),
	}
}


type Y21iaz2cH_hd struct {
	YW8Er6USJbZ	RcKjQr_MA
	GaL3Hix	CBEbiugxbncF
}

func (y6pzOXbIFd *Y21iaz2cH_hd) GoString() string {
	return "dnsmessage.Resource{" +
		"Header: " +  /*line :1*/ y6pzOXbIFd.YW8Er6USJbZ.GoString() +
		", Body: &" +  /*line :1*/ y6pzOXbIFd.GaL3Hix.GoString() +
		"}"
}


type CBEbiugxbncF interface {
	
	qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error)

	
	
	nyxiTeOZMSuB() Ze6KllqJw

	
	GoString() string
}


func (y6pzOXbIFd *Y21iaz2cH_hd) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	if y6pzOXbIFd.GaL3Hix == nil {
		return bnTR5Q, buZl0ZOk
	}
	mDT_QygTVa := bnTR5Q
	y6pzOXbIFd.YW8Er6USJbZ.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.GaL3Hix.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ y6pzOXbIFd.YW8Er6USJbZ.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
	if gX34Xg != nil {
		return bnTR5Q, &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.GaL3Hix.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
	if gX34Xg != nil {
		return bnTR5Q, &gxaf3H1N{"content", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ y6pzOXbIFd.YW8Er6USJbZ.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return mDT_QygTVa, gX34Xg
	}
	return bnTR5Q, nil
}















type GShepB struct {
	iF4J4nmSrIG	[]byte
	qhfP9faKXUrv	cNNn2uM

	hyzN6b	igQ2ZFMu
	ezDiiYjm	int
	xgkc6BEfhxRq	int
	hZicRuBgal	bool
	pzG9E8ZaoYkb	RcKjQr_MA
}


func (xTe5pzHUWS *GShepB) Start(bnTR5Q []byte) (HlVtnUC, error) {
	if xTe5pzHUWS.iF4J4nmSrIG != nil {
		*xTe5pzHUWS = GShepB{}
	}
	xTe5pzHUWS.iF4J4nmSrIG = bnTR5Q
	var gX34Xg error
	if xTe5pzHUWS.ezDiiYjm, gX34Xg =  /*line :1*/ xTe5pzHUWS.qhfP9faKXUrv.gXjxaL(bnTR5Q, 0); gX34Xg != nil {
		return HlVtnUC{}, &gxaf3H1N{"unpacking header", gX34Xg}
	}
	xTe5pzHUWS.hyzN6b = sectionQuestions
	return  /*line :1*/ xTe5pzHUWS.qhfP9faKXUrv.cNNn2uM(), nil
}

func (xTe5pzHUWS *GShepB) m_vGkd(gHo332FMv igQ2ZFMu) error {
	if xTe5pzHUWS.hyzN6b < gHo332FMv {
		return WbM9aD3mRP
	}
	if xTe5pzHUWS.hyzN6b > gHo332FMv {
		return BbQhjXWx
	}
	xTe5pzHUWS.hZicRuBgal = false
	if xTe5pzHUWS.xgkc6BEfhxRq ==  /*line :1*/ int( /*line :1*/ xTe5pzHUWS.qhfP9faKXUrv.ctLSIZuX6Xq(gHo332FMv)) {
		xTe5pzHUWS.xgkc6BEfhxRq = 0
		xTe5pzHUWS.hyzN6b++
		return BbQhjXWx
	}
	return nil
}

func (xTe5pzHUWS *GShepB) iuPfRd64TMw8(gHo332FMv igQ2ZFMu) (Y21iaz2cH_hd, error) {
	var y6pzOXbIFd Y21iaz2cH_hd
	var gX34Xg error
	y6pzOXbIFd.YW8Er6USJbZ, gX34Xg =  /*line :1*/ xTe5pzHUWS.xLqYf4p6(gHo332FMv)
	if gX34Xg != nil {
		return y6pzOXbIFd, gX34Xg
	}
	xTe5pzHUWS.hZicRuBgal = false
	y6pzOXbIFd.GaL3Hix, xTe5pzHUWS.ezDiiYjm, gX34Xg =  /*line :1*/ hJ8DUBMrV(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm, y6pzOXbIFd.YW8Er6USJbZ)
	if gX34Xg != nil {
		return Y21iaz2cH_hd{}, &gxaf3H1N{"unpacking " + hfLZTsMtk[gHo332FMv], gX34Xg}
	}
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}

func (xTe5pzHUWS *GShepB) xLqYf4p6(gHo332FMv igQ2ZFMu) (RcKjQr_MA, error) {
	if xTe5pzHUWS.hZicRuBgal {
		return xTe5pzHUWS.pzG9E8ZaoYkb, nil
	}
	if gX34Xg :=  /*line :1*/ xTe5pzHUWS.m_vGkd(gHo332FMv); gX34Xg != nil {
		return RcKjQr_MA{}, gX34Xg
	}
	var afzqoC5U0 RcKjQr_MA
	sG72__8, gX34Xg :=  /*line :1*/ afzqoC5U0.gXjxaL(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return RcKjQr_MA{}, gX34Xg
	}
	xTe5pzHUWS.hZicRuBgal = true
	xTe5pzHUWS.pzG9E8ZaoYkb = afzqoC5U0
	xTe5pzHUWS.ezDiiYjm = sG72__8
	return afzqoC5U0, nil
}

func (xTe5pzHUWS *GShepB) gDHvqtxO(gHo332FMv igQ2ZFMu) error {
	if xTe5pzHUWS.hZicRuBgal {
		xrV0mH := xTe5pzHUWS.ezDiiYjm +  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
		if xrV0mH >  /*line :1*/ len(xTe5pzHUWS.iF4J4nmSrIG) {
			return eCOkG0CgXc
		}
		xTe5pzHUWS.ezDiiYjm = xrV0mH
		xTe5pzHUWS.hZicRuBgal = false
		xTe5pzHUWS.xgkc6BEfhxRq++
		return nil
	}
	if gX34Xg :=  /*line :1*/ xTe5pzHUWS.m_vGkd(gHo332FMv); gX34Xg != nil {
		return gX34Xg
	}
	var gX34Xg error
	xTe5pzHUWS.ezDiiYjm, gX34Xg =  /*line :1*/ gDHvqtxO(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return &gxaf3H1N{"skipping: " + hfLZTsMtk[gHo332FMv], gX34Xg}
	}
	xTe5pzHUWS.xgkc6BEfhxRq++
	return nil
}


func (xTe5pzHUWS *GShepB) Question() (LNQ5aNrCORB, error) {
	if gX34Xg :=  /*line :1*/ xTe5pzHUWS.m_vGkd(sectionQuestions); gX34Xg != nil {
		return LNQ5aNrCORB{}, gX34Xg
	}
	var hQ82XQvk Toccq2dE3
	sG72__8, gX34Xg :=  /*line :1*/ hQ82XQvk.gXjxaL(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return LNQ5aNrCORB{}, &gxaf3H1N{"unpacking Question.Name", gX34Xg}
	}
	cBHIMYfgKC57, sG72__8, gX34Xg :=  /*line :1*/ hhXchL(xTe5pzHUWS.iF4J4nmSrIG, sG72__8)
	if gX34Xg != nil {
		return LNQ5aNrCORB{}, &gxaf3H1N{"unpacking Question.Type", gX34Xg}
	}
	_XUtQJmaFfB, sG72__8, gX34Xg :=  /*line :1*/ oeBQxlFN_(xTe5pzHUWS.iF4J4nmSrIG, sG72__8)
	if gX34Xg != nil {
		return LNQ5aNrCORB{}, &gxaf3H1N{"unpacking Question.Class", gX34Xg}
	}
	xTe5pzHUWS.ezDiiYjm = sG72__8
	xTe5pzHUWS.xgkc6BEfhxRq++
	return LNQ5aNrCORB{hQ82XQvk, cBHIMYfgKC57, _XUtQJmaFfB}, nil
}


func (xTe5pzHUWS *GShepB) AllQuestions() ([]LNQ5aNrCORB, error) {

	gu2knW := []LNQ5aNrCORB{}
	for {
		uemIRl, gX34Xg :=  /*line :1*/ xTe5pzHUWS.Question()
		if gX34Xg == BbQhjXWx {
			return gu2knW, nil
		}
		if gX34Xg != nil {
			return nil, gX34Xg
		}
		gu2knW =  /*line :1*/ append(gu2knW, uemIRl)
	}
}


func (xTe5pzHUWS *GShepB) SkipQuestion() error {
	if gX34Xg :=  /*line :1*/ xTe5pzHUWS.m_vGkd(sectionQuestions); gX34Xg != nil {
		return gX34Xg
	}
	sG72__8, gX34Xg :=  /*line :1*/ jn8cFCTm(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return &gxaf3H1N{"skipping Question Name", gX34Xg}
	}
	if sG72__8, gX34Xg =  /*line :1*/ x8jK6KVTG(xTe5pzHUWS.iF4J4nmSrIG, sG72__8); gX34Xg != nil {
		return &gxaf3H1N{"skipping Question Type", gX34Xg}
	}
	if sG72__8, gX34Xg =  /*line :1*/ fBdnWTMUZxl(xTe5pzHUWS.iF4J4nmSrIG, sG72__8); gX34Xg != nil {
		return &gxaf3H1N{"skipping Question Class", gX34Xg}
	}
	xTe5pzHUWS.ezDiiYjm = sG72__8
	xTe5pzHUWS.xgkc6BEfhxRq++
	return nil
}


func (xTe5pzHUWS *GShepB) SkipAllQuestions() error {
	for {
		if gX34Xg :=  /*line :1*/ xTe5pzHUWS.SkipQuestion(); gX34Xg == BbQhjXWx {
			return nil
		} else if gX34Xg != nil {
			return gX34Xg
		}
	}
}


func (xTe5pzHUWS *GShepB) AnswerHeader() (RcKjQr_MA, error) {
	return  /*line :1*/ xTe5pzHUWS.xLqYf4p6(sectionAnswers)
}


func (xTe5pzHUWS *GShepB) Answer() (Y21iaz2cH_hd, error) {
	return  /*line :1*/ xTe5pzHUWS.iuPfRd64TMw8(sectionAnswers)
}


func (xTe5pzHUWS *GShepB) AllAnswers() ([]Y21iaz2cH_hd, error) {

	aZuE3TFPdpn :=  /*line :1*/ int(xTe5pzHUWS.qhfP9faKXUrv.j1sYZuJo)
	if aZuE3TFPdpn > 20 {
		aZuE3TFPdpn = 20
	}
	hH4AEDtMvIxH :=  /*line :1*/ make([]Y21iaz2cH_hd, 0, aZuE3TFPdpn)
	for {
		gZpicd, gX34Xg :=  /*line :1*/ xTe5pzHUWS.Answer()
		if gX34Xg == BbQhjXWx {
			return hH4AEDtMvIxH, nil
		}
		if gX34Xg != nil {
			return nil, gX34Xg
		}
		hH4AEDtMvIxH =  /*line :1*/ append(hH4AEDtMvIxH, gZpicd)
	}
}


func (xTe5pzHUWS *GShepB) SkipAnswer() error {
	return  /*line :1*/ xTe5pzHUWS.gDHvqtxO(sectionAnswers)
}


func (xTe5pzHUWS *GShepB) SkipAllAnswers() error {
	for {
		if gX34Xg :=  /*line :1*/ xTe5pzHUWS.SkipAnswer(); gX34Xg == BbQhjXWx {
			return nil
		} else if gX34Xg != nil {
			return gX34Xg
		}
	}
}


func (xTe5pzHUWS *GShepB) AuthorityHeader() (RcKjQr_MA, error) {
	return  /*line :1*/ xTe5pzHUWS.xLqYf4p6(sectionAuthorities)
}


func (xTe5pzHUWS *GShepB) Authority() (Y21iaz2cH_hd, error) {
	return  /*line :1*/ xTe5pzHUWS.iuPfRd64TMw8(sectionAuthorities)
}


func (xTe5pzHUWS *GShepB) AllAuthorities() ([]Y21iaz2cH_hd, error) {

	aZuE3TFPdpn :=  /*line :1*/ int(xTe5pzHUWS.qhfP9faKXUrv.jHCad4ja2_)
	if aZuE3TFPdpn > 10 {
		aZuE3TFPdpn = 10
	}
	hH4AEDtMvIxH :=  /*line :1*/ make([]Y21iaz2cH_hd, 0, aZuE3TFPdpn)
	for {
		gZpicd, gX34Xg :=  /*line :1*/ xTe5pzHUWS.Authority()
		if gX34Xg == BbQhjXWx {
			return hH4AEDtMvIxH, nil
		}
		if gX34Xg != nil {
			return nil, gX34Xg
		}
		hH4AEDtMvIxH =  /*line :1*/ append(hH4AEDtMvIxH, gZpicd)
	}
}


func (xTe5pzHUWS *GShepB) SkipAuthority() error {
	return  /*line :1*/ xTe5pzHUWS.gDHvqtxO(sectionAuthorities)
}


func (xTe5pzHUWS *GShepB) SkipAllAuthorities() error {
	for {
		if gX34Xg :=  /*line :1*/ xTe5pzHUWS.SkipAuthority(); gX34Xg == BbQhjXWx {
			return nil
		} else if gX34Xg != nil {
			return gX34Xg
		}
	}
}


func (xTe5pzHUWS *GShepB) AdditionalHeader() (RcKjQr_MA, error) {
	return  /*line :1*/ xTe5pzHUWS.xLqYf4p6(sectionAdditionals)
}


func (xTe5pzHUWS *GShepB) Additional() (Y21iaz2cH_hd, error) {
	return  /*line :1*/ xTe5pzHUWS.iuPfRd64TMw8(sectionAdditionals)
}


func (xTe5pzHUWS *GShepB) AllAdditionals() ([]Y21iaz2cH_hd, error) {

	aZuE3TFPdpn :=  /*line :1*/ int(xTe5pzHUWS.qhfP9faKXUrv.lLb2iA7OmE)
	if aZuE3TFPdpn > 10 {
		aZuE3TFPdpn = 10
	}
	hH4AEDtMvIxH :=  /*line :1*/ make([]Y21iaz2cH_hd, 0, aZuE3TFPdpn)
	for {
		gZpicd, gX34Xg :=  /*line :1*/ xTe5pzHUWS.Additional()
		if gX34Xg == BbQhjXWx {
			return hH4AEDtMvIxH, nil
		}
		if gX34Xg != nil {
			return nil, gX34Xg
		}
		hH4AEDtMvIxH =  /*line :1*/ append(hH4AEDtMvIxH, gZpicd)
	}
}


func (xTe5pzHUWS *GShepB) SkipAdditional() error {
	return  /*line :1*/ xTe5pzHUWS.gDHvqtxO(sectionAdditionals)
}


func (xTe5pzHUWS *GShepB) SkipAllAdditionals() error {
	for {
		if gX34Xg :=  /*line :1*/ xTe5pzHUWS.SkipAdditional(); gX34Xg == BbQhjXWx {
			return nil
		} else if gX34Xg != nil {
			return gX34Xg
		}
	}
}





func (xTe5pzHUWS *GShepB) CNAMEResource() (Wi6AD2AtGlVJ, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeCNAME {
		return Wi6AD2AtGlVJ{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ dUDrigU(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return Wi6AD2AtGlVJ{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) MXResource() (VcLJCpUx, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeMX {
		return VcLJCpUx{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ xwkQUXNv1Y(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return VcLJCpUx{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) NSResource() (N6lUpPZaTwzd, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeNS {
		return N6lUpPZaTwzd{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ zaYHya2QB(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return N6lUpPZaTwzd{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) PTRResource() (IZO9QX, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypePTR {
		return IZO9QX{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ cVrwlV_8lAw7(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return IZO9QX{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) SOAResource() (XGpQdpEwOYzW, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeSOA {
		return XGpQdpEwOYzW{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ v0qwapce(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return XGpQdpEwOYzW{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) TXTResource() (ZInMhlZB8ns, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeTXT {
		return ZInMhlZB8ns{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ f6S8es(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm, xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	if gX34Xg != nil {
		return ZInMhlZB8ns{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) SRVResource() (BoxvLe6gQ, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeSRV {
		return BoxvLe6gQ{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ awtg4i_rM0r(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return BoxvLe6gQ{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) AResource() (LWO_y_Rk, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeA {
		return LWO_y_Rk{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ hCPe6zKCQW(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return LWO_y_Rk{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) AAAAResource() (LnvROw, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeAAAA {
		return LnvROw{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ esP_aR0yxX(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm)
	if gX34Xg != nil {
		return LnvROw{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) OPTResource() (PsSXLq1p, error) {
	if !xTe5pzHUWS.hZicRuBgal || xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK != TypeOPT {
		return PsSXLq1p{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ zjmSiB1Vf(xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm, xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	if gX34Xg != nil {
		return PsSXLq1p{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}





func (xTe5pzHUWS *GShepB) UnknownResource() (RgazROMZAobM, error) {
	if !xTe5pzHUWS.hZicRuBgal {
		return RgazROMZAobM{}, WbM9aD3mRP
	}
	y6pzOXbIFd, gX34Xg :=  /*line :1*/ bOvzgv2Iz7IT(xTe5pzHUWS.pzG9E8ZaoYkb.BxqvzBsK, xTe5pzHUWS.iF4J4nmSrIG, xTe5pzHUWS.ezDiiYjm, xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	if gX34Xg != nil {
		return RgazROMZAobM{}, gX34Xg
	}
	xTe5pzHUWS.ezDiiYjm +=  /*line :1*/ int(xTe5pzHUWS.pzG9E8ZaoYkb.GUbG0VpMXHI)
	xTe5pzHUWS.hZicRuBgal = false
	xTe5pzHUWS.xgkc6BEfhxRq++
	return y6pzOXbIFd, nil
}


func (qPTmameBxee0 *P_fQ6x) Unpack(bnTR5Q []byte) error {
	var xTe5pzHUWS GShepB
	var gX34Xg error
	if qPTmameBxee0.HlVtnUC, gX34Xg =  /*line :1*/ xTe5pzHUWS.Start(bnTR5Q); gX34Xg != nil {
		return gX34Xg
	}
	if qPTmameBxee0.Il4Uso_ecr, gX34Xg =  /*line :1*/ xTe5pzHUWS.AllQuestions(); gX34Xg != nil {
		return gX34Xg
	}
	if qPTmameBxee0.JiHd3TLzo, gX34Xg =  /*line :1*/ xTe5pzHUWS.AllAnswers(); gX34Xg != nil {
		return gX34Xg
	}
	if qPTmameBxee0.ZiuFEUcXZwi, gX34Xg =  /*line :1*/ xTe5pzHUWS.AllAuthorities(); gX34Xg != nil {
		return gX34Xg
	}
	if qPTmameBxee0.W5XWEcd, gX34Xg =  /*line :1*/ xTe5pzHUWS.AllAdditionals(); gX34Xg != nil {
		return gX34Xg
	}
	return nil
}


func (qPTmameBxee0 *P_fQ6x) Pack() ([]byte, error) {
	return  /*line :1*/ qPTmameBxee0.AppendPack( /*line :1*/ make([]byte, 0, packStartingCap))
}



func (qPTmameBxee0 *P_fQ6x) AppendPack(uUTUcsSMYDR []byte) ([]byte, error) {

	if  /*line :1*/ len(qPTmameBxee0.Il4Uso_ecr) >  /*line :1*/ int(^ /*line :1*/ uint16(0)) {
		return nil, cwapDN1O
	}
	if  /*line :1*/ len(qPTmameBxee0.JiHd3TLzo) >  /*line :1*/ int(^ /*line :1*/ uint16(0)) {
		return nil, j6t3f1Ay28
	}
	if  /*line :1*/ len(qPTmameBxee0.ZiuFEUcXZwi) >  /*line :1*/ int(^ /*line :1*/ uint16(0)) {
		return nil, aq6mWMTC
	}
	if  /*line :1*/ len(qPTmameBxee0.W5XWEcd) >  /*line :1*/ int(^ /*line :1*/ uint16(0)) {
		return nil, mDKEFxIt3Z
	}

	var iDU4uPqI cNNn2uM
	iDU4uPqI.sYp49l6kf, iDU4uPqI.mrl9diwhBeyV =  /*line :1*/ qPTmameBxee0.HlVtnUC.qalrXmt()

	iDU4uPqI.sCqv4utCjzi =  /*line :1*/ uint16( /*line :1*/ len(qPTmameBxee0.Il4Uso_ecr))
	iDU4uPqI.j1sYZuJo =  /*line :1*/ uint16( /*line :1*/ len(qPTmameBxee0.JiHd3TLzo))
	iDU4uPqI.jHCad4ja2_ =  /*line :1*/ uint16( /*line :1*/ len(qPTmameBxee0.ZiuFEUcXZwi))
	iDU4uPqI.lLb2iA7OmE =  /*line :1*/ uint16( /*line :1*/ len(qPTmameBxee0.W5XWEcd))

	daqdl2fzzSsy :=  /*line :1*/ len(uUTUcsSMYDR)
	bnTR5Q :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR)

	dLNWUDJK0 := map[string]int{}

	for eG5npDpL0 := range qPTmameBxee0.Il4Uso_ecr {
		var gX34Xg error
		if bnTR5Q, gX34Xg =  /*line :1*/ qPTmameBxee0.Il4Uso_ecr[eG5npDpL0].qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy); gX34Xg != nil {
			return nil, &gxaf3H1N{"packing Question", gX34Xg}
		}
	}
	for eG5npDpL0 := range qPTmameBxee0.JiHd3TLzo {
		var gX34Xg error
		if bnTR5Q, gX34Xg =  /*line :1*/ qPTmameBxee0.JiHd3TLzo[eG5npDpL0].qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy); gX34Xg != nil {
			return nil, &gxaf3H1N{"packing Answer", gX34Xg}
		}
	}
	for eG5npDpL0 := range qPTmameBxee0.ZiuFEUcXZwi {
		var gX34Xg error
		if bnTR5Q, gX34Xg =  /*line :1*/ qPTmameBxee0.ZiuFEUcXZwi[eG5npDpL0].qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy); gX34Xg != nil {
			return nil, &gxaf3H1N{"packing Authority", gX34Xg}
		}
	}
	for eG5npDpL0 := range qPTmameBxee0.W5XWEcd {
		var gX34Xg error
		if bnTR5Q, gX34Xg =  /*line :1*/ qPTmameBxee0.W5XWEcd[eG5npDpL0].qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy); gX34Xg != nil {
			return nil, &gxaf3H1N{"packing Additional", gX34Xg}
		}
	}

	return bnTR5Q, nil
}


func (qPTmameBxee0 *P_fQ6x) GoString() string {
	pIEx844 := "dnsmessage.Message{Header: " +  /*line :1*/ qPTmameBxee0.HlVtnUC.GoString() + ", " +
		"Questions: []dnsmessage.Question{"
	if  /*line :1*/ len(qPTmameBxee0.Il4Uso_ecr) > 0 {
		pIEx844 +=  /*line :1*/ qPTmameBxee0.Il4Uso_ecr[0].GoString()
		for _, uemIRl := range qPTmameBxee0.Il4Uso_ecr[1:] {
			pIEx844 += ", " +  /*line :1*/ uemIRl.GoString()
		}
	}
	pIEx844 += "}, Answers: []dnsmessage.Resource{"
	if  /*line :1*/ len(qPTmameBxee0.JiHd3TLzo) > 0 {
		pIEx844 +=  /*line :1*/ qPTmameBxee0.JiHd3TLzo[0].GoString()
		for _, gZpicd := range qPTmameBxee0.JiHd3TLzo[1:] {
			pIEx844 += ", " +  /*line :1*/ gZpicd.GoString()
		}
	}
	pIEx844 += "}, Authorities: []dnsmessage.Resource{"
	if  /*line :1*/ len(qPTmameBxee0.ZiuFEUcXZwi) > 0 {
		pIEx844 +=  /*line :1*/ qPTmameBxee0.ZiuFEUcXZwi[0].GoString()
		for _, gZpicd := range qPTmameBxee0.ZiuFEUcXZwi[1:] {
			pIEx844 += ", " +  /*line :1*/ gZpicd.GoString()
		}
	}
	pIEx844 += "}, Additionals: []dnsmessage.Resource{"
	if  /*line :1*/ len(qPTmameBxee0.W5XWEcd) > 0 {
		pIEx844 +=  /*line :1*/ qPTmameBxee0.W5XWEcd[0].GoString()
		for _, gZpicd := range qPTmameBxee0.W5XWEcd[1:] {
			pIEx844 += ", " +  /*line :1*/ gZpicd.GoString()
		}
	}
	return pIEx844 + "}}"
}












type S5pnU5g struct {
	
	oWaigNBL9KXF	[]byte

	
	haWG1c1TC2C	igQ2ZFMu

	
	
	l_ptohD0XQU	cNNn2uM

	
	uQC4JN2UY	int

	
	
	w6VW53HxX7	map[string]int
}











func Dt4jP7nnj6c8(pU8PG4B []byte, iDU4uPqI HlVtnUC) S5pnU5g {
	if pU8PG4B == nil {
		pU8PG4B =  /*line :1*/ make([]byte, 0, packStartingCap)
	}
	uUTUcsSMYDR := S5pnU5g{oWaigNBL9KXF: pU8PG4B, uQC4JN2UY:  /*line :1*/ len(pU8PG4B)}
	uUTUcsSMYDR.l_ptohD0XQU.sYp49l6kf, uUTUcsSMYDR.l_ptohD0XQU.mrl9diwhBeyV =  /*line :1*/ iDU4uPqI.qalrXmt()
	var jORkdCRcCFiG [headerLen]byte
	uUTUcsSMYDR.oWaigNBL9KXF =  /*line :1*/ append(uUTUcsSMYDR.oWaigNBL9KXF, jORkdCRcCFiG[:]...)
	uUTUcsSMYDR.haWG1c1TC2C = sectionHeader
	return uUTUcsSMYDR
}












func (uUTUcsSMYDR *S5pnU5g) EnableCompression() {
	uUTUcsSMYDR.w6VW53HxX7 = map[string]int{}
}

func (uUTUcsSMYDR *S5pnU5g) sRF02KTTYbW4(pIEx844 igQ2ZFMu) error {
	if uUTUcsSMYDR.haWG1c1TC2C <= sectionNotStarted {
		return WbM9aD3mRP
	}
	if uUTUcsSMYDR.haWG1c1TC2C > pIEx844 {
		return BbQhjXWx
	}
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) StartQuestions() error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.sRF02KTTYbW4(sectionQuestions); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.haWG1c1TC2C = sectionQuestions
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) StartAnswers() error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.sRF02KTTYbW4(sectionAnswers); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.haWG1c1TC2C = sectionAnswers
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) StartAuthorities() error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.sRF02KTTYbW4(sectionAuthorities); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.haWG1c1TC2C = sectionAuthorities
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) StartAdditionals() error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.sRF02KTTYbW4(sectionAdditionals); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.haWG1c1TC2C = sectionAdditionals
	return nil
}

func (uUTUcsSMYDR *S5pnU5g) gIQ1yF() error {
	var ctLSIZuX6Xq *uint16
	var gX34Xg error
	switch uUTUcsSMYDR.haWG1c1TC2C {
	case sectionQuestions:
		ctLSIZuX6Xq = &uUTUcsSMYDR.l_ptohD0XQU.sCqv4utCjzi
		gX34Xg = cwapDN1O
	case sectionAnswers:
		ctLSIZuX6Xq = &uUTUcsSMYDR.l_ptohD0XQU.j1sYZuJo
		gX34Xg = j6t3f1Ay28
	case sectionAuthorities:
		ctLSIZuX6Xq = &uUTUcsSMYDR.l_ptohD0XQU.jHCad4ja2_
		gX34Xg = aq6mWMTC
	case sectionAdditionals:
		ctLSIZuX6Xq = &uUTUcsSMYDR.l_ptohD0XQU.lLb2iA7OmE
		gX34Xg = mDKEFxIt3Z
	}
	if *ctLSIZuX6Xq == ^ /*line :1*/ uint16(0) {
		return gX34Xg
	}
	*ctLSIZuX6Xq++
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) Question(uemIRl LNQ5aNrCORB) error {
	if uUTUcsSMYDR.haWG1c1TC2C < sectionQuestions {
		return WbM9aD3mRP
	}
	if uUTUcsSMYDR.haWG1c1TC2C > sectionQuestions {
		return BbQhjXWx
	}
	bnTR5Q, gX34Xg :=  /*line :1*/ uemIRl.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}

func (uUTUcsSMYDR *S5pnU5g) fAYfWOj() error {
	if uUTUcsSMYDR.haWG1c1TC2C < sectionAnswers {
		return WbM9aD3mRP
	}
	if uUTUcsSMYDR.haWG1c1TC2C > sectionAdditionals {
		return BbQhjXWx
	}
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) CNAMEResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd Wi6AD2AtGlVJ) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"CNAMEResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) MXResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd VcLJCpUx) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"MXResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) NSResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd N6lUpPZaTwzd) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"NSResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) PTRResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd IZO9QX) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"PTRResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) SOAResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd XGpQdpEwOYzW) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"SOAResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) TXTResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd ZInMhlZB8ns) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"TXTResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) SRVResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd BoxvLe6gQ) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"SRVResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) AResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd LWO_y_Rk) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"AResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) AAAAResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd LnvROw) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"AAAAResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) OPTResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd PsSXLq1p) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"OPTResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) UnknownResource(iDU4uPqI RcKjQr_MA, y6pzOXbIFd RgazROMZAobM) error {
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.fAYfWOj(); gX34Xg != nil {
		return gX34Xg
	}
	iDU4uPqI.BxqvzBsK =  /*line :1*/ y6pzOXbIFd.nyxiTeOZMSuB()
	bnTR5Q, yI_r355uZ2H, gX34Xg :=  /*line :1*/ iDU4uPqI.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY)
	if gX34Xg != nil {
		return &gxaf3H1N{"ResourceHeader", gX34Xg}
	}
	oyZvweJ :=  /*line :1*/ len(bnTR5Q)
	if bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.qalrXmt(bnTR5Q, uUTUcsSMYDR.w6VW53HxX7, uUTUcsSMYDR.uQC4JN2UY); gX34Xg != nil {
		return &gxaf3H1N{"UnknownResource body", gX34Xg}
	}
	if gX34Xg :=  /*line :1*/ iDU4uPqI.lsMytp7LClr(bnTR5Q, yI_r355uZ2H, oyZvweJ); gX34Xg != nil {
		return gX34Xg
	}
	if gX34Xg :=  /*line :1*/ uUTUcsSMYDR.gIQ1yF(); gX34Xg != nil {
		return gX34Xg
	}
	uUTUcsSMYDR.oWaigNBL9KXF = bnTR5Q
	return nil
}


func (uUTUcsSMYDR *S5pnU5g) Finish() ([]byte, error) {
	if uUTUcsSMYDR.haWG1c1TC2C < sectionHeader {
		return nil, WbM9aD3mRP
	}
	uUTUcsSMYDR.haWG1c1TC2C = sectionDone

	 /*line :1*/ uUTUcsSMYDR.l_ptohD0XQU.qalrXmt(uUTUcsSMYDR.oWaigNBL9KXF[uUTUcsSMYDR.uQC4JN2UY:uUTUcsSMYDR.uQC4JN2UY])
	return uUTUcsSMYDR.oWaigNBL9KXF, nil
}



type RcKjQr_MA struct {
	
	Pz3XQBChF6	Toccq2dE3

	
	
	
	BxqvzBsK	Ze6KllqJw

	
	
	ZP1ICM83GGQc	WhBR6DRkZ3p

	
	
	
	DmJCBb	uint32

	
	
	
	GUbG0VpMXHI	uint16
}


func (iDU4uPqI *RcKjQr_MA) GoString() string {
	return "dnsmessage.ResourceHeader{" +
		"Name: " +  /*line :1*/ iDU4uPqI.Pz3XQBChF6.GoString() + ", " +
		"Type: " +  /*line :1*/ iDU4uPqI.BxqvzBsK.GoString() + ", " +
		"Class: " +  /*line :1*/ iDU4uPqI.ZP1ICM83GGQc.GoString() + ", " +
		"TTL: " +  /*line :1*/ kzyBd0GDt(iDU4uPqI.DmJCBb) + ", " +
		"Length: " +  /*line :1*/ d20BK8p4Fr(iDU4uPqI.GUbG0VpMXHI) + "}"
}




func (iDU4uPqI *RcKjQr_MA) qalrXmt(mDT_QygTVa []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) (bnTR5Q []byte, yI_r355uZ2H int, gX34Xg error) {
	bnTR5Q = mDT_QygTVa
	if bnTR5Q, gX34Xg =  /*line :1*/ iDU4uPqI.Pz3XQBChF6.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy); gX34Xg != nil {
		return mDT_QygTVa, 0, &gxaf3H1N{"Name", gX34Xg}
	}
	bnTR5Q =  /*line :1*/ fnr1b28UTT_(bnTR5Q, iDU4uPqI.BxqvzBsK)
	bnTR5Q =  /*line :1*/ yoFzbbMmea8s(bnTR5Q, iDU4uPqI.ZP1ICM83GGQc)
	bnTR5Q =  /*line :1*/ znQepZk0(bnTR5Q, iDU4uPqI.DmJCBb)
	yI_r355uZ2H =  /*line :1*/ len(bnTR5Q)
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, iDU4uPqI.GUbG0VpMXHI)
	return bnTR5Q, yI_r355uZ2H, nil
}

func (iDU4uPqI *RcKjQr_MA) gXjxaL(bnTR5Q []byte, sG72__8 int) (int, error) {
	xrV0mH := sG72__8
	var gX34Xg error
	if xrV0mH, gX34Xg =  /*line :1*/ iDU4uPqI.Pz3XQBChF6.gXjxaL(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Name", gX34Xg}
	}
	if iDU4uPqI.BxqvzBsK, xrV0mH, gX34Xg =  /*line :1*/ hhXchL(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Type", gX34Xg}
	}
	if iDU4uPqI.ZP1ICM83GGQc, xrV0mH, gX34Xg =  /*line :1*/ oeBQxlFN_(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Class", gX34Xg}
	}
	if iDU4uPqI.DmJCBb, xrV0mH, gX34Xg =  /*line :1*/ f7YsoueNhdpJ(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"TTL", gX34Xg}
	}
	if iDU4uPqI.GUbG0VpMXHI, xrV0mH, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Length", gX34Xg}
	}
	return xrV0mH, nil
}







func (iDU4uPqI *RcKjQr_MA) lsMytp7LClr(bnTR5Q []byte, yI_r355uZ2H int, oyZvweJ int) error {
	cCiwfdlV :=  /*line :1*/ len(bnTR5Q) - oyZvweJ
	if cCiwfdlV >  /*line :1*/ int(^ /*line :1*/ uint16(0)) {
		return vvoBDOiQz
	}

	 /*line :1*/ dyO9ONJo5ZpR(bnTR5Q[yI_r355uZ2H:yI_r355uZ2H],  /*line :1*/ uint16(cCiwfdlV))
	iDU4uPqI.GUbG0VpMXHI =  /*line :1*/ uint16(cCiwfdlV)

	return nil
}


const (
	edns0Version	= 0

	edns0DNSSECOK	= 0x00008000
	ednsVersionMask	= 0x00ff0000
	edns0DNSSECOKMask	= 0x00ff8000
)




func (iDU4uPqI *RcKjQr_MA) SetEDNS0(bAlsSW int, pnuh3w2 QPDSMamH, qkPTdwgeB6ud bool) error {
	iDU4uPqI.Pz3XQBChF6 = Toccq2dE3{MaSC7z: [255]byte{'.'}, W7uI8JuUD: 1}
	iDU4uPqI.BxqvzBsK = TypeOPT
	iDU4uPqI.ZP1ICM83GGQc =  /*line :1*/ WhBR6DRkZ3p(bAlsSW)
	iDU4uPqI.DmJCBb =  /*line :1*/ uint32(pnuh3w2) >> 4 << 24
	if qkPTdwgeB6ud {
		iDU4uPqI.DmJCBb |= edns0DNSSECOK
	}
	return nil
}


func (iDU4uPqI *RcKjQr_MA) DNSSECAllowed() bool {
	return iDU4uPqI.DmJCBb&edns0DNSSECOKMask == edns0DNSSECOK
}




func (iDU4uPqI *RcKjQr_MA) ExtendedRCode(g7RygaRF5 QPDSMamH) QPDSMamH {
	if iDU4uPqI.DmJCBb&ednsVersionMask == edns0Version {
		return  /*line :1*/ QPDSMamH(iDU4uPqI.DmJCBb>>24<<4) | g7RygaRF5
	}
	return g7RygaRF5
}

func gDHvqtxO(bnTR5Q []byte, sG72__8 int) (int, error) {
	xrV0mH, gX34Xg :=  /*line :1*/ jn8cFCTm(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Name", gX34Xg}
	}
	if xrV0mH, gX34Xg =  /*line :1*/ x8jK6KVTG(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Type", gX34Xg}
	}
	if xrV0mH, gX34Xg =  /*line :1*/ fBdnWTMUZxl(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Class", gX34Xg}
	}
	if xrV0mH, gX34Xg =  /*line :1*/ qdiETho(bnTR5Q, xrV0mH); gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"TTL", gX34Xg}
	}
	dejv8qS, xrV0mH, gX34Xg :=  /*line :1*/ f5QJp8imv(bnTR5Q, xrV0mH)
	if gX34Xg != nil {
		return sG72__8, &gxaf3H1N{"Length", gX34Xg}
	}
	if xrV0mH +=  /*line :1*/ int(dejv8qS); xrV0mH >  /*line :1*/ len(bnTR5Q) {
		return sG72__8, eCOkG0CgXc
	}
	return xrV0mH, nil
}


func dyO9ONJo5ZpR(bnTR5Q []byte, dqdks9SB0NU uint16) []byte {
	return  /*line :1*/ append(bnTR5Q,  /*line :1*/ byte(dqdks9SB0NU>>8),  /*line :1*/ byte(dqdks9SB0NU))
}

func f5QJp8imv(bnTR5Q []byte, sG72__8 int) (uint16, int, error) {
	if sG72__8+uint16Len >  /*line :1*/ len(bnTR5Q) {
		return 0, sG72__8, yf5I2II0zMg
	}
	return  /*line :1*/ uint16(bnTR5Q[sG72__8])<<8 |  /*line :1*/ uint16(bnTR5Q[sG72__8+1]), sG72__8 + uint16Len, nil
}

func bfsNarV(bnTR5Q []byte, sG72__8 int) (int, error) {
	if sG72__8+uint16Len >  /*line :1*/ len(bnTR5Q) {
		return sG72__8, yf5I2II0zMg
	}
	return sG72__8 + uint16Len, nil
}


func fnr1b28UTT_(bnTR5Q []byte, dqdks9SB0NU Ze6KllqJw) []byte {
	return  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q,  /*line :1*/ uint16(dqdks9SB0NU))
}

func hhXchL(bnTR5Q []byte, sG72__8 int) (Ze6KllqJw, int, error) {
	qavq6q, rL_9YGYQqnf, gX34Xg :=  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
	return  /*line :1*/ Ze6KllqJw(qavq6q), rL_9YGYQqnf, gX34Xg
}

func x8jK6KVTG(bnTR5Q []byte, sG72__8 int) (int, error) {
	return  /*line :1*/ bfsNarV(bnTR5Q, sG72__8)
}


func yoFzbbMmea8s(bnTR5Q []byte, dqdks9SB0NU WhBR6DRkZ3p) []byte {
	return  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q,  /*line :1*/ uint16(dqdks9SB0NU))
}

func oeBQxlFN_(bnTR5Q []byte, sG72__8 int) (WhBR6DRkZ3p, int, error) {
	iiYnULwuUYX, rL_9YGYQqnf, gX34Xg :=  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
	return  /*line :1*/ WhBR6DRkZ3p(iiYnULwuUYX), rL_9YGYQqnf, gX34Xg
}

func fBdnWTMUZxl(bnTR5Q []byte, sG72__8 int) (int, error) {
	return  /*line :1*/ bfsNarV(bnTR5Q, sG72__8)
}


func znQepZk0(bnTR5Q []byte, dqdks9SB0NU uint32) []byte {
	return  /*line :1*/ append(
		bnTR5Q,
		 /*line :1*/ byte(dqdks9SB0NU>>24),
		 /*line :1*/ byte(dqdks9SB0NU>>16),
		 /*line :1*/ byte(dqdks9SB0NU>>8),
		 /*line :1*/ byte(dqdks9SB0NU),
	)
}

func f7YsoueNhdpJ(bnTR5Q []byte, sG72__8 int) (uint32, int, error) {
	if sG72__8+uint32Len >  /*line :1*/ len(bnTR5Q) {
		return 0, sG72__8, yf5I2II0zMg
	}
	fz2IPsg :=  /*line :1*/ uint32(bnTR5Q[sG72__8])<<24 |  /*line :1*/ uint32(bnTR5Q[sG72__8+1])<<16 |  /*line :1*/ uint32(bnTR5Q[sG72__8+2])<<8 |  /*line :1*/ uint32(bnTR5Q[sG72__8+3])
	return fz2IPsg, sG72__8 + uint32Len, nil
}

func qdiETho(bnTR5Q []byte, sG72__8 int) (int, error) {
	if sG72__8+uint32Len >  /*line :1*/ len(bnTR5Q) {
		return sG72__8, yf5I2II0zMg
	}
	return sG72__8 + uint32Len, nil
}


func hlJq3Rv3C(bnTR5Q []byte, dqdks9SB0NU string) ([]byte, error) {
	cHCtW0kPOtzR :=  /*line :1*/ len(dqdks9SB0NU)
	if cHCtW0kPOtzR > 255 {
		return nil, mRcqHJN
	}
	bnTR5Q =  /*line :1*/ append(bnTR5Q,  /*line :1*/ byte(cHCtW0kPOtzR))
	bnTR5Q =  /*line :1*/ append(bnTR5Q, dqdks9SB0NU...)

	return bnTR5Q, nil
}

func lOMgsPKM(bnTR5Q []byte, sG72__8 int) (string, int, error) {
	if sG72__8 >=  /*line :1*/ len(bnTR5Q) {
		return "", sG72__8, yf5I2II0zMg
	}
	yNAeDauB := sG72__8 + 1
	fhllbCMUDsHV := yNAeDauB +  /*line :1*/ int(bnTR5Q[sG72__8])
	if fhllbCMUDsHV >  /*line :1*/ len(bnTR5Q) {
		return "", sG72__8, onAp5pO
	}
	return  /*line :1*/ string(bnTR5Q[yNAeDauB:fhllbCMUDsHV]), fhllbCMUDsHV, nil
}


func fWK7jjdGo6qI(bnTR5Q []byte, dqdks9SB0NU []byte) []byte {
	return  /*line :1*/ append(bnTR5Q, dqdks9SB0NU...)
}

func xQoZai(bnTR5Q []byte, sG72__8 int, dqdks9SB0NU []byte) (int, error) {
	xrV0mH := sG72__8 +  /*line :1*/ len(dqdks9SB0NU)
	if xrV0mH >  /*line :1*/ len(bnTR5Q) {
		return sG72__8, yf5I2II0zMg
	}
	 /*line :1*/ copy(dqdks9SB0NU, bnTR5Q[sG72__8:xrV0mH])
	return xrV0mH, nil
}

const nonEncodedNameMax = 254



type Toccq2dE3 struct {
	MaSC7z	[255]byte
	W7uI8JuUD	uint8
}


func HfBZcpE4(hQ82XQvk string) (Toccq2dE3, error) {
	aZuE3TFPdpn := Toccq2dE3{W7uI8JuUD:  /*line :1*/ uint8( /*line :1*/ len(hQ82XQvk))}
	if  /*line :1*/ len(hQ82XQvk) >  /*line :1*/ len(aZuE3TFPdpn.MaSC7z) {
		return Toccq2dE3{}, onAp5pO
	}
	 /*line :1*/ copy(aZuE3TFPdpn.MaSC7z[:], hQ82XQvk)
	return aZuE3TFPdpn, nil
}


func Dst4RHWdE0K(hQ82XQvk string) Toccq2dE3 {
	aZuE3TFPdpn, gX34Xg :=  /*line :1*/ HfBZcpE4(hQ82XQvk)
	if gX34Xg != nil {
		 /*line :1*/ panic("creating name: " +  /*line :1*/ gX34Xg.Error())
	}
	return aZuE3TFPdpn
}


func (aZuE3TFPdpn Toccq2dE3) String() string {
	return  /*line :1*/ string(aZuE3TFPdpn.MaSC7z[:aZuE3TFPdpn.W7uI8JuUD])
}


func (aZuE3TFPdpn *Toccq2dE3) GoString() string {
	return `dnsmessage.MustNewName("` +  /*line :1*/ tUbBJah(aZuE3TFPdpn.MaSC7z[:aZuE3TFPdpn.W7uI8JuUD]) + `")`
}








func (aZuE3TFPdpn *Toccq2dE3) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	mDT_QygTVa := bnTR5Q

	if aZuE3TFPdpn.W7uI8JuUD > nonEncodedNameMax {
		return nil, iF0RY15sidD
	}

	if aZuE3TFPdpn.W7uI8JuUD == 0 || aZuE3TFPdpn.MaSC7z[aZuE3TFPdpn.W7uI8JuUD-1] != '.' {
		return mDT_QygTVa, oNVzpKip
	}

	if aZuE3TFPdpn.MaSC7z[0] == '.' && aZuE3TFPdpn.W7uI8JuUD == 1 {
		return  /*line :1*/ append(bnTR5Q, 0), nil
	}

	for eG5npDpL0, aYX_Yj8 := 0, 0; eG5npDpL0 <  /*line :1*/ int(aZuE3TFPdpn.W7uI8JuUD); eG5npDpL0++ {

		if aZuE3TFPdpn.MaSC7z[eG5npDpL0] == '.' {

			if eG5npDpL0-aYX_Yj8 >= 1<<6 {
				return mDT_QygTVa, qF0KX1
			}

			if eG5npDpL0-aYX_Yj8 == 0 {
				return mDT_QygTVa, u07BDH
			}

			bnTR5Q =  /*line :1*/ append(bnTR5Q,  /*line :1*/ byte(eG5npDpL0-aYX_Yj8))

			for hmb_Gl_rHJ := aYX_Yj8; hmb_Gl_rHJ < eG5npDpL0; hmb_Gl_rHJ++ {
				bnTR5Q =  /*line :1*/ append(bnTR5Q, aZuE3TFPdpn.MaSC7z[hmb_Gl_rHJ])
			}

			aYX_Yj8 = eG5npDpL0 + 1
			continue
		}

		if (eG5npDpL0 == 0 || aZuE3TFPdpn.MaSC7z[eG5npDpL0-1] == '.') && dLNWUDJK0 != nil {
			if jqWKBAg2eI8, afB2jgY := dLNWUDJK0[ /*line :1*/ string(aZuE3TFPdpn.MaSC7z[eG5npDpL0:])]; afB2jgY {

				return  /*line :1*/ append(bnTR5Q,  /*line :1*/ byte(jqWKBAg2eI8>>8|0xC0),  /*line :1*/ byte(jqWKBAg2eI8)), nil
			}

			if  /*line :1*/ len(bnTR5Q) <=  /*line :1*/ int(^ /*line :1*/ uint16(0)>>2) {
				dLNWUDJK0[ /*line :1*/ string(aZuE3TFPdpn.MaSC7z[eG5npDpL0:])] =  /*line :1*/ len(bnTR5Q) - daqdl2fzzSsy
			}
		}
	}
	return  /*line :1*/ append(bnTR5Q, 0), nil
}


func (aZuE3TFPdpn *Toccq2dE3) gXjxaL(bnTR5Q []byte, sG72__8 int) (int, error) {
	return  /*line :1*/ aZuE3TFPdpn.qnRf0bz(bnTR5Q, sG72__8, true)
}

func (aZuE3TFPdpn *Toccq2dE3) qnRf0bz(bnTR5Q []byte, sG72__8 int, sp0S4SYZBP bool) (int, error) {

	cBUakC := sG72__8

	xrV0mH := sG72__8

	
	var jqWKBAg2eI8 int

	hQ82XQvk := aZuE3TFPdpn.MaSC7z[:0]

Loop:
	for {
		if cBUakC >=  /*line :1*/ len(bnTR5Q) {
			return sG72__8, yf5I2II0zMg
		}
		iiYnULwuUYX :=  /*line :1*/ int(bnTR5Q[cBUakC])
		cBUakC++
		switch iiYnULwuUYX & 0xC0 {
		case 0x00:
			if iiYnULwuUYX == 0x00 {

				break Loop
			}
			fhllbCMUDsHV := cBUakC + iiYnULwuUYX
			if fhllbCMUDsHV >  /*line :1*/ len(bnTR5Q) {
				return sG72__8, onAp5pO
			}

			for _, fz2IPsg := range bnTR5Q[cBUakC:fhllbCMUDsHV] {
				if fz2IPsg == '.' {
					return sG72__8, jx390dmJG
				}
			}

			hQ82XQvk =  /*line :1*/ append(hQ82XQvk, bnTR5Q[cBUakC:fhllbCMUDsHV]...)
			hQ82XQvk =  /*line :1*/ append(hQ82XQvk, '.')
			cBUakC = fhllbCMUDsHV
		case 0xC0:
			if !sp0S4SYZBP {
				return sG72__8, mGz62tBsms9M
			}
			if cBUakC >=  /*line :1*/ len(bnTR5Q) {
				return sG72__8, t0_XfM
			}
			a336kp8vr6gj := bnTR5Q[cBUakC]
			cBUakC++
			if jqWKBAg2eI8 == 0 {
				xrV0mH = cBUakC
			}

			if jqWKBAg2eI8++; jqWKBAg2eI8 > 10 {
				return sG72__8, rZYKeUXMwZ
			}
			cBUakC = (iiYnULwuUYX^0xC0)<<8 |  /*line :1*/ int(a336kp8vr6gj)
		default:

			return sG72__8, t7KOKcour
		}
	}
	if  /*line :1*/ len(hQ82XQvk) == 0 {
		hQ82XQvk =  /*line :1*/ append(hQ82XQvk, '.')
	}
	if  /*line :1*/ len(hQ82XQvk) > nonEncodedNameMax {
		return sG72__8, iF0RY15sidD
	}
	aZuE3TFPdpn.W7uI8JuUD =  /*line :1*/ uint8( /*line :1*/ len(hQ82XQvk))
	if jqWKBAg2eI8 == 0 {
		xrV0mH = cBUakC
	}
	return xrV0mH, nil
}

func jn8cFCTm(bnTR5Q []byte, sG72__8 int) (int, error) {

	xrV0mH := sG72__8

Loop:
	for {
		if xrV0mH >=  /*line :1*/ len(bnTR5Q) {
			return sG72__8, yf5I2II0zMg
		}
		iiYnULwuUYX :=  /*line :1*/ int(bnTR5Q[xrV0mH])
		xrV0mH++
		switch iiYnULwuUYX & 0xC0 {
		case 0x00:
			if iiYnULwuUYX == 0x00 {

				break Loop
			}

			xrV0mH += iiYnULwuUYX
			if xrV0mH >  /*line :1*/ len(bnTR5Q) {
				return sG72__8, onAp5pO
			}
		case 0xC0:

			xrV0mH++

			break Loop
		default:

			return sG72__8, t7KOKcour
		}
	}

	return xrV0mH, nil
}


type LNQ5aNrCORB struct {
	FLDYrkuX	Toccq2dE3
	XsqCP7Wa65g	Ze6KllqJw
	GBiiAZ64kqf6	WhBR6DRkZ3p
}


func (uemIRl *LNQ5aNrCORB) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	bnTR5Q, gX34Xg :=  /*line :1*/ uemIRl.FLDYrkuX.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
	if gX34Xg != nil {
		return bnTR5Q, &gxaf3H1N{"Name", gX34Xg}
	}
	bnTR5Q =  /*line :1*/ fnr1b28UTT_(bnTR5Q, uemIRl.XsqCP7Wa65g)
	return  /*line :1*/ yoFzbbMmea8s(bnTR5Q, uemIRl.GBiiAZ64kqf6), nil
}


func (uemIRl *LNQ5aNrCORB) GoString() string {
	return "dnsmessage.Question{" +
		"Name: " +  /*line :1*/ uemIRl.FLDYrkuX.GoString() + ", " +
		"Type: " +  /*line :1*/ uemIRl.XsqCP7Wa65g.GoString() + ", " +
		"Class: " +  /*line :1*/ uemIRl.GBiiAZ64kqf6.GoString() + "}"
}

func hJ8DUBMrV(bnTR5Q []byte, sG72__8 int, afzqoC5U0 RcKjQr_MA) (CBEbiugxbncF, int, error) {
	var (
		y6pzOXbIFd	CBEbiugxbncF
		gX34Xg	error
		hQ82XQvk	string
	)
	switch afzqoC5U0.BxqvzBsK {
	case TypeA:
		var fncgDxDdlI LWO_y_Rk
		fncgDxDdlI, gX34Xg =  /*line :1*/ hCPe6zKCQW(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "A"
	case TypeNS:
		var fncgDxDdlI N6lUpPZaTwzd
		fncgDxDdlI, gX34Xg =  /*line :1*/ zaYHya2QB(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "NS"
	case TypeCNAME:
		var fncgDxDdlI Wi6AD2AtGlVJ
		fncgDxDdlI, gX34Xg =  /*line :1*/ dUDrigU(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "CNAME"
	case TypeSOA:
		var fncgDxDdlI XGpQdpEwOYzW
		fncgDxDdlI, gX34Xg =  /*line :1*/ v0qwapce(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "SOA"
	case TypePTR:
		var fncgDxDdlI IZO9QX
		fncgDxDdlI, gX34Xg =  /*line :1*/ cVrwlV_8lAw7(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "PTR"
	case TypeMX:
		var fncgDxDdlI VcLJCpUx
		fncgDxDdlI, gX34Xg =  /*line :1*/ xwkQUXNv1Y(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "MX"
	case TypeTXT:
		var fncgDxDdlI ZInMhlZB8ns
		fncgDxDdlI, gX34Xg =  /*line :1*/ f6S8es(bnTR5Q, sG72__8, afzqoC5U0.GUbG0VpMXHI)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "TXT"
	case TypeAAAA:
		var fncgDxDdlI LnvROw
		fncgDxDdlI, gX34Xg =  /*line :1*/ esP_aR0yxX(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "AAAA"
	case TypeSRV:
		var fncgDxDdlI BoxvLe6gQ
		fncgDxDdlI, gX34Xg =  /*line :1*/ awtg4i_rM0r(bnTR5Q, sG72__8)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "SRV"
	case TypeOPT:
		var fncgDxDdlI PsSXLq1p
		fncgDxDdlI, gX34Xg =  /*line :1*/ zjmSiB1Vf(bnTR5Q, sG72__8, afzqoC5U0.GUbG0VpMXHI)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "OPT"
	default:
		var fncgDxDdlI RgazROMZAobM
		fncgDxDdlI, gX34Xg =  /*line :1*/ bOvzgv2Iz7IT(afzqoC5U0.BxqvzBsK, bnTR5Q, sG72__8, afzqoC5U0.GUbG0VpMXHI)
		y6pzOXbIFd = &fncgDxDdlI
		hQ82XQvk = "Unknown"
	}
	if gX34Xg != nil {
		return nil, sG72__8, &gxaf3H1N{hQ82XQvk + " record", gX34Xg}
	}
	return y6pzOXbIFd, sG72__8 +  /*line :1*/ int(afzqoC5U0.GUbG0VpMXHI), nil
}


type Wi6AD2AtGlVJ struct {
	SiT_7E895B87 Toccq2dE3
}

func (y6pzOXbIFd *Wi6AD2AtGlVJ) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeCNAME
}


func (y6pzOXbIFd *Wi6AD2AtGlVJ) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	return  /*line :1*/ y6pzOXbIFd.SiT_7E895B87.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
}


func (y6pzOXbIFd *Wi6AD2AtGlVJ) GoString() string {
	return "dnsmessage.CNAMEResource{CNAME: " +  /*line :1*/ y6pzOXbIFd.SiT_7E895B87.GoString() + "}"
}

func dUDrigU(bnTR5Q []byte, sG72__8 int) (Wi6AD2AtGlVJ, error) {
	var yPQqM6 Toccq2dE3
	if _, gX34Xg :=  /*line :1*/ yPQqM6.gXjxaL(bnTR5Q, sG72__8); gX34Xg != nil {
		return Wi6AD2AtGlVJ{}, gX34Xg
	}
	return Wi6AD2AtGlVJ{yPQqM6}, nil
}


type VcLJCpUx struct {
	DdHoYyYFMHhM	uint16
	USylcaqC	Toccq2dE3
}

func (y6pzOXbIFd *VcLJCpUx) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeMX
}


func (y6pzOXbIFd *VcLJCpUx) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	mDT_QygTVa := bnTR5Q
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, y6pzOXbIFd.DdHoYyYFMHhM)
	bnTR5Q, gX34Xg :=  /*line :1*/ y6pzOXbIFd.USylcaqC.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
	if gX34Xg != nil {
		return mDT_QygTVa, &gxaf3H1N{"MXResource.MX", gX34Xg}
	}
	return bnTR5Q, nil
}


func (y6pzOXbIFd *VcLJCpUx) GoString() string {
	return "dnsmessage.MXResource{" +
		"Pref: " +  /*line :1*/ d20BK8p4Fr(y6pzOXbIFd.DdHoYyYFMHhM) + ", " +
		"MX: " +  /*line :1*/ y6pzOXbIFd.USylcaqC.GoString() + "}"
}

func xwkQUXNv1Y(bnTR5Q []byte, sG72__8 int) (VcLJCpUx, error) {
	fdayJ4g1_E, sG72__8, gX34Xg :=  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return VcLJCpUx{}, &gxaf3H1N{"Pref", gX34Xg}
	}
	var aT8d_tn5O Toccq2dE3
	if _, gX34Xg :=  /*line :1*/ aT8d_tn5O.gXjxaL(bnTR5Q, sG72__8); gX34Xg != nil {
		return VcLJCpUx{}, &gxaf3H1N{"MX", gX34Xg}
	}
	return VcLJCpUx{fdayJ4g1_E, aT8d_tn5O}, nil
}


type N6lUpPZaTwzd struct {
	JccbZj Toccq2dE3
}

func (y6pzOXbIFd *N6lUpPZaTwzd) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeNS
}


func (y6pzOXbIFd *N6lUpPZaTwzd) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	return  /*line :1*/ y6pzOXbIFd.JccbZj.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
}


func (y6pzOXbIFd *N6lUpPZaTwzd) GoString() string {
	return "dnsmessage.NSResource{NS: " +  /*line :1*/ y6pzOXbIFd.JccbZj.GoString() + "}"
}

func zaYHya2QB(bnTR5Q []byte, sG72__8 int) (N6lUpPZaTwzd, error) {
	var cACwpzv3a9Z2 Toccq2dE3
	if _, gX34Xg :=  /*line :1*/ cACwpzv3a9Z2.gXjxaL(bnTR5Q, sG72__8); gX34Xg != nil {
		return N6lUpPZaTwzd{}, gX34Xg
	}
	return N6lUpPZaTwzd{cACwpzv3a9Z2}, nil
}


type IZO9QX struct {
	ExEq5bU3Lem Toccq2dE3
}

func (y6pzOXbIFd *IZO9QX) nyxiTeOZMSuB() Ze6KllqJw {
	return TypePTR
}


func (y6pzOXbIFd *IZO9QX) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	return  /*line :1*/ y6pzOXbIFd.ExEq5bU3Lem.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
}


func (y6pzOXbIFd *IZO9QX) GoString() string {
	return "dnsmessage.PTRResource{PTR: " +  /*line :1*/ y6pzOXbIFd.ExEq5bU3Lem.GoString() + "}"
}

func cVrwlV_8lAw7(bnTR5Q []byte, sG72__8 int) (IZO9QX, error) {
	var jqWKBAg2eI8 Toccq2dE3
	if _, gX34Xg :=  /*line :1*/ jqWKBAg2eI8.gXjxaL(bnTR5Q, sG72__8); gX34Xg != nil {
		return IZO9QX{}, gX34Xg
	}
	return IZO9QX{jqWKBAg2eI8}, nil
}


type XGpQdpEwOYzW struct {
	RC4MVz2w	Toccq2dE3
	W9bHKEMg	Toccq2dE3
	Ga1yyyV9x	uint32
	MN0WSev	uint32
	ZOoISpd_uTn	uint32
	HkimKEKc9uLI	uint32

	
	
	
	M31d93KP4zak	uint32
}

func (y6pzOXbIFd *XGpQdpEwOYzW) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeSOA
}


func (y6pzOXbIFd *XGpQdpEwOYzW) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	mDT_QygTVa := bnTR5Q
	bnTR5Q, gX34Xg :=  /*line :1*/ y6pzOXbIFd.RC4MVz2w.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
	if gX34Xg != nil {
		return mDT_QygTVa, &gxaf3H1N{"SOAResource.NS", gX34Xg}
	}
	bnTR5Q, gX34Xg =  /*line :1*/ y6pzOXbIFd.W9bHKEMg.qalrXmt(bnTR5Q, dLNWUDJK0, daqdl2fzzSsy)
	if gX34Xg != nil {
		return mDT_QygTVa, &gxaf3H1N{"SOAResource.MBox", gX34Xg}
	}
	bnTR5Q =  /*line :1*/ znQepZk0(bnTR5Q, y6pzOXbIFd.Ga1yyyV9x)
	bnTR5Q =  /*line :1*/ znQepZk0(bnTR5Q, y6pzOXbIFd.MN0WSev)
	bnTR5Q =  /*line :1*/ znQepZk0(bnTR5Q, y6pzOXbIFd.ZOoISpd_uTn)
	bnTR5Q =  /*line :1*/ znQepZk0(bnTR5Q, y6pzOXbIFd.HkimKEKc9uLI)
	return  /*line :1*/ znQepZk0(bnTR5Q, y6pzOXbIFd.M31d93KP4zak), nil
}


func (y6pzOXbIFd *XGpQdpEwOYzW) GoString() string {
	return "dnsmessage.SOAResource{" +
		"NS: " +  /*line :1*/ y6pzOXbIFd.RC4MVz2w.GoString() + ", " +
		"MBox: " +  /*line :1*/ y6pzOXbIFd.W9bHKEMg.GoString() + ", " +
		"Serial: " +  /*line :1*/ kzyBd0GDt(y6pzOXbIFd.Ga1yyyV9x) + ", " +
		"Refresh: " +  /*line :1*/ kzyBd0GDt(y6pzOXbIFd.MN0WSev) + ", " +
		"Retry: " +  /*line :1*/ kzyBd0GDt(y6pzOXbIFd.ZOoISpd_uTn) + ", " +
		"Expire: " +  /*line :1*/ kzyBd0GDt(y6pzOXbIFd.HkimKEKc9uLI) + ", " +
		"MinTTL: " +  /*line :1*/ kzyBd0GDt(y6pzOXbIFd.M31d93KP4zak) + "}"
}

func v0qwapce(bnTR5Q []byte, sG72__8 int) (XGpQdpEwOYzW, error) {
	var cACwpzv3a9Z2 Toccq2dE3
	sG72__8, gX34Xg :=  /*line :1*/ cACwpzv3a9Z2.gXjxaL(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return XGpQdpEwOYzW{}, &gxaf3H1N{"NS", gX34Xg}
	}
	var xolmUq Toccq2dE3
	if sG72__8, gX34Xg =  /*line :1*/ xolmUq.gXjxaL(bnTR5Q, sG72__8); gX34Xg != nil {
		return XGpQdpEwOYzW{}, &gxaf3H1N{"MBox", gX34Xg}
	}
	jNfegqKx45j4, sG72__8, gX34Xg :=  /*line :1*/ f7YsoueNhdpJ(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return XGpQdpEwOYzW{}, &gxaf3H1N{"Serial", gX34Xg}
	}
	xaHxPmDrJ, sG72__8, gX34Xg :=  /*line :1*/ f7YsoueNhdpJ(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return XGpQdpEwOYzW{}, &gxaf3H1N{"Refresh", gX34Xg}
	}
	znreQx1Q, sG72__8, gX34Xg :=  /*line :1*/ f7YsoueNhdpJ(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return XGpQdpEwOYzW{}, &gxaf3H1N{"Retry", gX34Xg}
	}
	gjXMsI_9, sG72__8, gX34Xg :=  /*line :1*/ f7YsoueNhdpJ(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return XGpQdpEwOYzW{}, &gxaf3H1N{"Expire", gX34Xg}
	}
	bRI8aALHk, _, gX34Xg :=  /*line :1*/ f7YsoueNhdpJ(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return XGpQdpEwOYzW{}, &gxaf3H1N{"MinTTL", gX34Xg}
	}
	return XGpQdpEwOYzW{cACwpzv3a9Z2, xolmUq, jNfegqKx45j4, xaHxPmDrJ, znreQx1Q, gjXMsI_9, bRI8aALHk}, nil
}


type ZInMhlZB8ns struct {
	FhHS8hCoPF2m []string
}

func (y6pzOXbIFd *ZInMhlZB8ns) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeTXT
}


func (y6pzOXbIFd *ZInMhlZB8ns) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	mDT_QygTVa := bnTR5Q
	for _, pIEx844 := range y6pzOXbIFd.FhHS8hCoPF2m {
		var gX34Xg error
		bnTR5Q, gX34Xg =  /*line :1*/ hlJq3Rv3C(bnTR5Q, pIEx844)
		if gX34Xg != nil {
			return mDT_QygTVa, gX34Xg
		}
	}
	return bnTR5Q, nil
}


func (y6pzOXbIFd *ZInMhlZB8ns) GoString() string {
	pIEx844 := "dnsmessage.TXTResource{TXT: []string{"
	if  /*line :1*/ len(y6pzOXbIFd.FhHS8hCoPF2m) == 0 {
		return pIEx844 + "}}"
	}
	pIEx844 += `"` +  /*line :1*/ tUbBJah([] /*line :1*/ byte(y6pzOXbIFd.FhHS8hCoPF2m[0]))
	for _, qavq6q := range y6pzOXbIFd.FhHS8hCoPF2m[1:] {
		pIEx844 += `", "` +  /*line :1*/ tUbBJah([] /*line :1*/ byte(qavq6q))
	}
	return pIEx844 + `"}}`
}

func f6S8es(bnTR5Q []byte, sG72__8 int, dejv8qS uint16) (ZInMhlZB8ns, error) {
	qLRmartv :=  /*line :1*/ make([]string, 0, 1)
	for aZuE3TFPdpn :=  /*line :1*/ uint16(0); aZuE3TFPdpn < dejv8qS; {
		var qavq6q string
		var gX34Xg error
		if qavq6q, sG72__8, gX34Xg =  /*line :1*/ lOMgsPKM(bnTR5Q, sG72__8); gX34Xg != nil {
			return ZInMhlZB8ns{}, &gxaf3H1N{"text", gX34Xg}
		}

		if dejv8qS-aZuE3TFPdpn <  /*line :1*/ uint16( /*line :1*/ len(qavq6q))+1 {
			return ZInMhlZB8ns{}, onAp5pO
		}
		aZuE3TFPdpn +=  /*line :1*/ uint16( /*line :1*/ len(qavq6q)) + 1
		qLRmartv =  /*line :1*/ append(qLRmartv, qavq6q)
	}
	return ZInMhlZB8ns{qLRmartv}, nil
}


type BoxvLe6gQ struct {
	AJkmRs98D	uint16
	JpLvcVWu	uint16
	WhxUlP4uIVBk	uint16
	HvwgUZxi0	Toccq2dE3	
}

func (y6pzOXbIFd *BoxvLe6gQ) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeSRV
}


func (y6pzOXbIFd *BoxvLe6gQ) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	mDT_QygTVa := bnTR5Q
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, y6pzOXbIFd.AJkmRs98D)
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, y6pzOXbIFd.JpLvcVWu)
	bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, y6pzOXbIFd.WhxUlP4uIVBk)
	bnTR5Q, gX34Xg :=  /*line :1*/ y6pzOXbIFd.HvwgUZxi0.qalrXmt(bnTR5Q, nil, daqdl2fzzSsy)
	if gX34Xg != nil {
		return mDT_QygTVa, &gxaf3H1N{"SRVResource.Target", gX34Xg}
	}
	return bnTR5Q, nil
}


func (y6pzOXbIFd *BoxvLe6gQ) GoString() string {
	return "dnsmessage.SRVResource{" +
		"Priority: " +  /*line :1*/ d20BK8p4Fr(y6pzOXbIFd.AJkmRs98D) + ", " +
		"Weight: " +  /*line :1*/ d20BK8p4Fr(y6pzOXbIFd.JpLvcVWu) + ", " +
		"Port: " +  /*line :1*/ d20BK8p4Fr(y6pzOXbIFd.WhxUlP4uIVBk) + ", " +
		"Target: " +  /*line :1*/ y6pzOXbIFd.HvwgUZxi0.GoString() + "}"
}

func awtg4i_rM0r(bnTR5Q []byte, sG72__8 int) (BoxvLe6gQ, error) {
	aezFgLynw, sG72__8, gX34Xg :=  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return BoxvLe6gQ{}, &gxaf3H1N{"Priority", gX34Xg}
	}
	jAVJKCM, sG72__8, gX34Xg :=  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return BoxvLe6gQ{}, &gxaf3H1N{"Weight", gX34Xg}
	}
	nRLOMu9RKH, sG72__8, gX34Xg :=  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
	if gX34Xg != nil {
		return BoxvLe6gQ{}, &gxaf3H1N{"Port", gX34Xg}
	}
	var f6Dfh_ Toccq2dE3
	if _, gX34Xg :=  /*line :1*/ f6Dfh_.qnRf0bz(bnTR5Q, sG72__8, false); gX34Xg != nil {
		return BoxvLe6gQ{}, &gxaf3H1N{"Target", gX34Xg}
	}
	return BoxvLe6gQ{aezFgLynw, jAVJKCM, nRLOMu9RKH, f6Dfh_}, nil
}


type LWO_y_Rk struct {
	M5FpSRlH9 [4]byte
}

func (y6pzOXbIFd *LWO_y_Rk) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeA
}


func (y6pzOXbIFd *LWO_y_Rk) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	return  /*line :1*/ fWK7jjdGo6qI(bnTR5Q, y6pzOXbIFd.M5FpSRlH9[:]), nil
}


func (y6pzOXbIFd *LWO_y_Rk) GoString() string {
	return "dnsmessage.AResource{" +
		"A: [4]byte{" +  /*line :1*/ eguZ3sWx8(y6pzOXbIFd.M5FpSRlH9[:]) + "}}"
}

func hCPe6zKCQW(bnTR5Q []byte, sG72__8 int) (LWO_y_Rk, error) {
	var gZpicd [4]byte
	if _, gX34Xg :=  /*line :1*/ xQoZai(bnTR5Q, sG72__8, gZpicd[:]); gX34Xg != nil {
		return LWO_y_Rk{}, gX34Xg
	}
	return LWO_y_Rk{gZpicd}, nil
}


type LnvROw struct {
	DaUu4HQOB9 [16]byte
}

func (y6pzOXbIFd *LnvROw) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeAAAA
}


func (y6pzOXbIFd *LnvROw) GoString() string {
	return "dnsmessage.AAAAResource{" +
		"AAAA: [16]byte{" +  /*line :1*/ eguZ3sWx8(y6pzOXbIFd.DaUu4HQOB9[:]) + "}}"
}


func (y6pzOXbIFd *LnvROw) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	return  /*line :1*/ fWK7jjdGo6qI(bnTR5Q, y6pzOXbIFd.DaUu4HQOB9[:]), nil
}

func esP_aR0yxX(bnTR5Q []byte, sG72__8 int) (LnvROw, error) {
	var aTcLlilf [16]byte
	if _, gX34Xg :=  /*line :1*/ xQoZai(bnTR5Q, sG72__8, aTcLlilf[:]); gX34Xg != nil {
		return LnvROw{}, gX34Xg
	}
	return LnvROw{aTcLlilf}, nil
}





type PsSXLq1p struct {
	JWyKBUkabR []H6aztVxf
}





type H6aztVxf struct {
	ClEZOaQYFUYP	uint16	
	YEJw69s3kGMU	[]byte
}


func (rL_9YGYQqnf *H6aztVxf) GoString() string {
	return "dnsmessage.Option{" +
		"Code: " +  /*line :1*/ d20BK8p4Fr(rL_9YGYQqnf.ClEZOaQYFUYP) + ", " +
		"Data: []byte{" +  /*line :1*/ eguZ3sWx8(rL_9YGYQqnf.YEJw69s3kGMU) + "}}"
}

func (y6pzOXbIFd *PsSXLq1p) nyxiTeOZMSuB() Ze6KllqJw {
	return TypeOPT
}

func (y6pzOXbIFd *PsSXLq1p) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	for _, uci7OV := range y6pzOXbIFd.JWyKBUkabR {
		bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, uci7OV.ClEZOaQYFUYP)
		cHCtW0kPOtzR :=  /*line :1*/ uint16( /*line :1*/ len(uci7OV.YEJw69s3kGMU))
		bnTR5Q =  /*line :1*/ dyO9ONJo5ZpR(bnTR5Q, cHCtW0kPOtzR)
		bnTR5Q =  /*line :1*/ fWK7jjdGo6qI(bnTR5Q, uci7OV.YEJw69s3kGMU)
	}
	return bnTR5Q, nil
}


func (y6pzOXbIFd *PsSXLq1p) GoString() string {
	pIEx844 := "dnsmessage.OPTResource{Options: []dnsmessage.Option{"
	if  /*line :1*/ len(y6pzOXbIFd.JWyKBUkabR) == 0 {
		return pIEx844 + "}}"
	}
	pIEx844 +=  /*line :1*/ y6pzOXbIFd.JWyKBUkabR[0].GoString()
	for _, rL_9YGYQqnf := range y6pzOXbIFd.JWyKBUkabR[1:] {
		pIEx844 += ", " +  /*line :1*/ rL_9YGYQqnf.GoString()
	}
	return pIEx844 + "}}"
}

func zjmSiB1Vf(bnTR5Q []byte, sG72__8 int, dejv8qS uint16) (PsSXLq1p, error) {
	var aFcgTEY []H6aztVxf
	for l46fwAKWS := sG72__8; sG72__8 < l46fwAKWS+ /*line :1*/ int(dejv8qS); {
		var gX34Xg error
		var rL_9YGYQqnf H6aztVxf
		rL_9YGYQqnf.ClEZOaQYFUYP, sG72__8, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
		if gX34Xg != nil {
			return PsSXLq1p{}, &gxaf3H1N{"Code", gX34Xg}
		}
		var cHCtW0kPOtzR uint16
		cHCtW0kPOtzR, sG72__8, gX34Xg =  /*line :1*/ f5QJp8imv(bnTR5Q, sG72__8)
		if gX34Xg != nil {
			return PsSXLq1p{}, &gxaf3H1N{"Data", gX34Xg}
		}
		rL_9YGYQqnf.YEJw69s3kGMU =  /*line :1*/ make([]byte, cHCtW0kPOtzR)
		if  /*line :1*/ copy(rL_9YGYQqnf.YEJw69s3kGMU, bnTR5Q[sG72__8:]) !=  /*line :1*/ int(cHCtW0kPOtzR) {
			return PsSXLq1p{}, &gxaf3H1N{"Data", onAp5pO}
		}
		sG72__8 +=  /*line :1*/ int(cHCtW0kPOtzR)
		aFcgTEY =  /*line :1*/ append(aFcgTEY, rL_9YGYQqnf)
	}
	return PsSXLq1p{aFcgTEY}, nil
}


type RgazROMZAobM struct {
	Ikw9Z4z2g	Ze6KllqJw
	WiBJd1c	[]byte
}

func (y6pzOXbIFd *RgazROMZAobM) nyxiTeOZMSuB() Ze6KllqJw {
	return y6pzOXbIFd.Ikw9Z4z2g
}


func (y6pzOXbIFd *RgazROMZAobM) qalrXmt(bnTR5Q []byte, dLNWUDJK0 map[string]int, daqdl2fzzSsy int) ([]byte, error) {
	return  /*line :1*/ fWK7jjdGo6qI(bnTR5Q, y6pzOXbIFd.WiBJd1c[:]), nil
}


func (y6pzOXbIFd *RgazROMZAobM) GoString() string {
	return "dnsmessage.UnknownResource{" +
		"Type: " +  /*line :1*/ y6pzOXbIFd.Ikw9Z4z2g.GoString() + ", " +
		"Data: []byte{" +  /*line :1*/ eguZ3sWx8(y6pzOXbIFd.WiBJd1c) + "}}"
}

func bOvzgv2Iz7IT(xIt4yd Ze6KllqJw, bnTR5Q []byte, sG72__8 int, dejv8qS uint16) (RgazROMZAobM, error) {
	geG90cPouljN := RgazROMZAobM{
		Ikw9Z4z2g:	xIt4yd,
		WiBJd1c:	 /*line :1*/ make([]byte, dejv8qS),
	}
	if _, gX34Xg :=  /*line :1*/ xQoZai(bnTR5Q, sG72__8, geG90cPouljN.WiBJd1c); gX34Xg != nil {
		return RgazROMZAobM{}, gX34Xg
	}
	return geG90cPouljN, nil
}
