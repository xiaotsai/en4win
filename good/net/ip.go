//line :1
package gF9mZs2

import (
	"internal/bytealg"
	itoa "JkjxVS"
	netip "iPdCam_KQOXj"
)


const (
	IPv4len	= 4
	IPv6len	= 16
)










type GNraIvYhB []byte





type OcrmAeBW7m []byte


type M6TN4rba7 struct {
	RJ6BaPgcU	GNraIvYhB	
	Yj5CZGxK1oow	OcrmAeBW7m	
}



func Ip1NFCL6(uI7LZDHm6, fIadEXIim6l, hl8w2lFX, ica44Q byte) GNraIvYhB {
	fMPVz2iGiyq :=  /*line :1*/ make(GNraIvYhB, IPv6len)
	 /*line :1*/ copy(fMPVz2iGiyq, jFGjdWti3EA)
	fMPVz2iGiyq[12] = uI7LZDHm6
	fMPVz2iGiyq[13] = fIadEXIim6l
	fMPVz2iGiyq[14] = hl8w2lFX
	fMPVz2iGiyq[15] = ica44Q
	return fMPVz2iGiyq
}

var jFGjdWti3EA = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}



func X6z68s(uI7LZDHm6, fIadEXIim6l, hl8w2lFX, ica44Q byte) OcrmAeBW7m {
	fMPVz2iGiyq :=  /*line :1*/ make(OcrmAeBW7m, IPv4len)
	fMPVz2iGiyq[0] = uI7LZDHm6
	fMPVz2iGiyq[1] = fIadEXIim6l
	fMPVz2iGiyq[2] = hl8w2lFX
	fMPVz2iGiyq[3] = ica44Q
	return fMPVz2iGiyq
}




func RYiA8Y9b(rwDKH0MQ0, ikVsUbWnbH int) OcrmAeBW7m {
	if ikVsUbWnbH != 8*IPv4len && ikVsUbWnbH != 8*IPv6len {
		return nil
	}
	if rwDKH0MQ0 < 0 || rwDKH0MQ0 > ikVsUbWnbH {
		return nil
	}
	v3upVm := ikVsUbWnbH / 8
	z_nqmUMeJFH :=  /*line :1*/ make(OcrmAeBW7m, v3upVm)
	doauF8Y :=  /*line :1*/ uint(rwDKH0MQ0)
	for eaAUaB7Z := 0; eaAUaB7Z < v3upVm; eaAUaB7Z++ {
		if doauF8Y >= 8 {
			z_nqmUMeJFH[eaAUaB7Z] = 0xff
			doauF8Y -= 8
			continue
		}
		z_nqmUMeJFH[eaAUaB7Z] = ^ /*line :1*/ byte(0xff >> doauF8Y)
		doauF8Y = 0
	}
	return z_nqmUMeJFH
}


var (
	Zk2B6i	=  /*line :1*/ Ip1NFCL6(255, 255, 255, 255)		
	ZOM2iqTm6	=  /*line :1*/ Ip1NFCL6(224, 0, 0, 1)		
	QR_gxXlZD	=  /*line :1*/ Ip1NFCL6(224, 0, 0, 2)		
	WHSMKoapR	=  /*line :1*/ Ip1NFCL6(0, 0, 0, 0)		
)


var (
	Zi20Wn	= GNraIvYhB{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	VIAfFsvboRvQ	= GNraIvYhB{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ICyGsSLLvx	= GNraIvYhB{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	WVJDbfZvD	= GNraIvYhB{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	TcPn00whaU_	= GNraIvYhB{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01}
	JaL4wPO	= GNraIvYhB{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)



func (kQ8_UEhxU GNraIvYhB) IsUnspecified() bool {
	return  /*line :1*/ kQ8_UEhxU.Equal(WHSMKoapR) ||  /*line :1*/ kQ8_UEhxU.Equal(VIAfFsvboRvQ)
}


func (kQ8_UEhxU GNraIvYhB) IsLoopback() bool {
	if di2u16 :=  /*line :1*/ kQ8_UEhxU.To4(); di2u16 != nil {
		return di2u16[0] == 127
	}
	return  /*line :1*/ kQ8_UEhxU.Equal(ICyGsSLLvx)
}



func (kQ8_UEhxU GNraIvYhB) IsPrivate() bool {
	if di2u16 :=  /*line :1*/ kQ8_UEhxU.To4(); di2u16 != nil {

		return di2u16[0] == 10 ||
			(di2u16[0] == 172 && di2u16[1]&0xf0 == 16) ||
			(di2u16[0] == 192 && di2u16[1] == 168)
	}

	return  /*line :1*/ len(kQ8_UEhxU) == IPv6len && kQ8_UEhxU[0]&0xfe == 0xfc
}


func (kQ8_UEhxU GNraIvYhB) IsMulticast() bool {
	if di2u16 :=  /*line :1*/ kQ8_UEhxU.To4(); di2u16 != nil {
		return di2u16[0]&0xf0 == 0xe0
	}
	return  /*line :1*/ len(kQ8_UEhxU) == IPv6len && kQ8_UEhxU[0] == 0xff
}



func (kQ8_UEhxU GNraIvYhB) IsInterfaceLocalMulticast() bool {
	return  /*line :1*/ len(kQ8_UEhxU) == IPv6len && kQ8_UEhxU[0] == 0xff && kQ8_UEhxU[1]&0x0f == 0x01
}



func (kQ8_UEhxU GNraIvYhB) IsLinkLocalMulticast() bool {
	if di2u16 :=  /*line :1*/ kQ8_UEhxU.To4(); di2u16 != nil {
		return di2u16[0] == 224 && di2u16[1] == 0 && di2u16[2] == 0
	}
	return  /*line :1*/ len(kQ8_UEhxU) == IPv6len && kQ8_UEhxU[0] == 0xff && kQ8_UEhxU[1]&0x0f == 0x02
}



func (kQ8_UEhxU GNraIvYhB) IsLinkLocalUnicast() bool {
	if di2u16 :=  /*line :1*/ kQ8_UEhxU.To4(); di2u16 != nil {
		return di2u16[0] == 169 && di2u16[1] == 254
	}
	return  /*line :1*/ len(kQ8_UEhxU) == IPv6len && kQ8_UEhxU[0] == 0xfe && kQ8_UEhxU[1]&0xc0 == 0x80
}









func (kQ8_UEhxU GNraIvYhB) IsGlobalUnicast() bool {
	return ( /*line :1*/ len(kQ8_UEhxU) == IPv4len ||  /*line :1*/ len(kQ8_UEhxU) == IPv6len) &&
		! /*line :1*/ kQ8_UEhxU.Equal(Zk2B6i) &&
		! /*line :1*/ kQ8_UEhxU.IsUnspecified() &&
		! /*line :1*/ kQ8_UEhxU.IsLoopback() &&
		! /*line :1*/ kQ8_UEhxU.IsMulticast() &&
		! /*line :1*/ kQ8_UEhxU.IsLinkLocalUnicast()
}


func oNyTxx(fMPVz2iGiyq GNraIvYhB) bool {
	for eaAUaB7Z := 0; eaAUaB7Z <  /*line :1*/ len(fMPVz2iGiyq); eaAUaB7Z++ {
		if fMPVz2iGiyq[eaAUaB7Z] != 0 {
			return false
		}
	}
	return true
}



func (kQ8_UEhxU GNraIvYhB) To4() GNraIvYhB {
	if  /*line :1*/ len(kQ8_UEhxU) == IPv4len {
		return kQ8_UEhxU
	}
	if  /*line :1*/ len(kQ8_UEhxU) == IPv6len &&
		 /*line :1*/ oNyTxx(kQ8_UEhxU[0:10]) &&
		kQ8_UEhxU[10] == 0xff &&
		kQ8_UEhxU[11] == 0xff {
		return kQ8_UEhxU[12:16]
	}
	return nil
}



func (kQ8_UEhxU GNraIvYhB) To16() GNraIvYhB {
	if  /*line :1*/ len(kQ8_UEhxU) == IPv4len {
		return  /*line :1*/ Ip1NFCL6(kQ8_UEhxU[0], kQ8_UEhxU[1], kQ8_UEhxU[2], kQ8_UEhxU[3])
	}
	if  /*line :1*/ len(kQ8_UEhxU) == IPv6len {
		return kQ8_UEhxU
	}
	return nil
}


var (
	bh23ZOF3	=  /*line :1*/ X6z68s(0xff, 0, 0, 0)
	w7Hy9oStG	=  /*line :1*/ X6z68s(0xff, 0xff, 0, 0)
	sGZkpnLniam	=  /*line :1*/ X6z68s(0xff, 0xff, 0xff, 0)
)




func (kQ8_UEhxU GNraIvYhB) DefaultMask() OcrmAeBW7m {
	if kQ8_UEhxU =  /*line :1*/ kQ8_UEhxU.To4(); kQ8_UEhxU == nil {
		return nil
	}
	switch {
	case kQ8_UEhxU[0] < 0x80:
		return bh23ZOF3
	case kQ8_UEhxU[0] < 0xC0:
		return w7Hy9oStG
	default:
		return sGZkpnLniam
	}
}

func w3Kk2g(fIadEXIim6l []byte) bool {
	for _, hl8w2lFX := range fIadEXIim6l {
		if hl8w2lFX != 0xff {
			return false
		}
	}
	return true
}


func (kQ8_UEhxU GNraIvYhB) Mask(vQyaaIll OcrmAeBW7m) GNraIvYhB {
	if  /*line :1*/ len(vQyaaIll) == IPv6len &&  /*line :1*/ len(kQ8_UEhxU) == IPv4len &&  /*line :1*/ w3Kk2g(vQyaaIll[:12]) {
		vQyaaIll = vQyaaIll[12:]
	}
	if  /*line :1*/ len(vQyaaIll) == IPv4len &&  /*line :1*/ len(kQ8_UEhxU) == IPv6len &&  /*line :1*/ bytealg.Equal(kQ8_UEhxU[:12], jFGjdWti3EA) {
		kQ8_UEhxU = kQ8_UEhxU[12:]
	}
	doauF8Y :=  /*line :1*/ len(kQ8_UEhxU)
	if doauF8Y !=  /*line :1*/ len(vQyaaIll) {
		return nil
	}
	hMkKN3mr :=  /*line :1*/ make(GNraIvYhB, doauF8Y)
	for eaAUaB7Z := 0; eaAUaB7Z < doauF8Y; eaAUaB7Z++ {
		hMkKN3mr[eaAUaB7Z] = kQ8_UEhxU[eaAUaB7Z] & vQyaaIll[eaAUaB7Z]
	}
	return hMkKN3mr
}







func (kQ8_UEhxU GNraIvYhB) String() string {
	if  /*line :1*/ len(kQ8_UEhxU) == 0 {
		return "<nil>"
	}

	if  /*line :1*/ len(kQ8_UEhxU) != IPv4len &&  /*line :1*/ len(kQ8_UEhxU) != IPv6len {
		return "?" +  /*line :1*/ _lLCd9kw3sG(kQ8_UEhxU)
	}

	if fRz6rzzhs94 :=  /*line :1*/ kQ8_UEhxU.To4();  /*line :1*/ len(fRz6rzzhs94) == IPv4len {
		return  /*line :1*/ netip.Vh24QUEOnu([4] /*line :1*/ byte(fRz6rzzhs94)).String()
	}
	return  /*line :1*/ netip.XtvbgBR([16] /*line :1*/ byte(kQ8_UEhxU)).String()
}

func _lLCd9kw3sG(fIadEXIim6l []byte) string {
	dRoFccaZ :=  /*line :1*/ make([]byte,  /*line :1*/ len(fIadEXIim6l)*2)
	for eaAUaB7Z, xkDIYW30c3 := range fIadEXIim6l {
		dRoFccaZ[eaAUaB7Z*2], dRoFccaZ[eaAUaB7Z*2+1] = hexDigit[xkDIYW30c3>>4], hexDigit[xkDIYW30c3&0xf]
	}
	return  /*line :1*/ string(dRoFccaZ)
}



func pXjCfqAUUC(kQ8_UEhxU GNraIvYhB) string {
	if  /*line :1*/ len(kQ8_UEhxU) == 0 {
		return ""
	}
	return  /*line :1*/ kQ8_UEhxU.String()
}




func (kQ8_UEhxU GNraIvYhB) MarshalText() ([]byte, error) {
	if  /*line :1*/ len(kQ8_UEhxU) == 0 {
		return [] /*line :1*/ byte(""), nil
	}
	if  /*line :1*/ len(kQ8_UEhxU) != IPv4len &&  /*line :1*/ len(kQ8_UEhxU) != IPv6len {
		return nil, &DAWLIQHc{Z68v0y: "invalid IP address", BCCnAgFxG:  /*line :1*/ _lLCd9kw3sG(kQ8_UEhxU)}
	}
	return [] /*line :1*/ byte( /*line :1*/ kQ8_UEhxU.String()), nil
}



func (kQ8_UEhxU *GNraIvYhB) UnmarshalText(oGr_KsEE3J []byte) error {
	if  /*line :1*/ len(oGr_KsEE3J) == 0 {
		*kQ8_UEhxU = nil
		return nil
	}
	dRoFccaZ :=  /*line :1*/ string(oGr_KsEE3J)
	tWO2olI :=  /*line :1*/ XtVFDxFx(dRoFccaZ)
	if tWO2olI == nil {
		return &DPDxTDQNU3Ih{FATAgZbL0: "IP address", GrpQb_kayach: dRoFccaZ}
	}
	*kQ8_UEhxU = tWO2olI
	return nil
}




func (kQ8_UEhxU GNraIvYhB) Equal(tWO2olI GNraIvYhB) bool {
	if  /*line :1*/ len(kQ8_UEhxU) ==  /*line :1*/ len(tWO2olI) {
		return  /*line :1*/ bytealg.Equal(kQ8_UEhxU, tWO2olI)
	}
	if  /*line :1*/ len(kQ8_UEhxU) == IPv4len &&  /*line :1*/ len(tWO2olI) == IPv6len {
		return  /*line :1*/ bytealg.Equal(tWO2olI[0:12], jFGjdWti3EA) &&  /*line :1*/ bytealg.Equal(kQ8_UEhxU, tWO2olI[12:])
	}
	if  /*line :1*/ len(kQ8_UEhxU) == IPv6len &&  /*line :1*/ len(tWO2olI) == IPv4len {
		return  /*line :1*/ bytealg.Equal(kQ8_UEhxU[0:12], jFGjdWti3EA) &&  /*line :1*/ bytealg.Equal(kQ8_UEhxU[12:], tWO2olI)
	}
	return false
}

func (kQ8_UEhxU GNraIvYhB) yuO3LyWZOg(tWO2olI GNraIvYhB) bool {
	return  /*line :1*/ kQ8_UEhxU.To4() != nil &&  /*line :1*/ tWO2olI.To4() != nil ||  /*line :1*/ kQ8_UEhxU.To16() != nil &&  /*line :1*/ kQ8_UEhxU.To4() == nil &&  /*line :1*/ tWO2olI.To16() != nil &&  /*line :1*/ tWO2olI.To4() == nil
}



func eAtJ8K_5w1p(vQyaaIll OcrmAeBW7m) int {
	var doauF8Y int
	for eaAUaB7Z, ljsCSk := range vQyaaIll {
		if ljsCSk == 0xff {
			doauF8Y += 8
			continue
		}

		for ljsCSk&0x80 != 0 {
			doauF8Y++
			ljsCSk <<= 1
		}

		if ljsCSk != 0 {
			return -1
		}
		for eaAUaB7Z++; eaAUaB7Z <  /*line :1*/ len(vQyaaIll); eaAUaB7Z++ {
			if vQyaaIll[eaAUaB7Z] != 0 {
				return -1
			}
		}
		break
	}
	return doauF8Y
}




func (z_nqmUMeJFH OcrmAeBW7m) Size() (rwDKH0MQ0, ikVsUbWnbH int) {
	rwDKH0MQ0, ikVsUbWnbH =  /*line :1*/ eAtJ8K_5w1p(z_nqmUMeJFH),  /*line :1*/ len(z_nqmUMeJFH)*8
	if rwDKH0MQ0 == -1 {
		return 0, 0
	}
	return
}


func (z_nqmUMeJFH OcrmAeBW7m) String() string {
	if  /*line :1*/ len(z_nqmUMeJFH) == 0 {
		return "<nil>"
	}
	return  /*line :1*/ _lLCd9kw3sG(z_nqmUMeJFH)
}

func li04uikLccY(doauF8Y *M6TN4rba7) (kQ8_UEhxU GNraIvYhB, z_nqmUMeJFH OcrmAeBW7m) {
	if kQ8_UEhxU =  /*line :1*/ doauF8Y.RJ6BaPgcU.To4(); kQ8_UEhxU == nil {
		kQ8_UEhxU = doauF8Y.RJ6BaPgcU
		if  /*line :1*/ len(kQ8_UEhxU) != IPv6len {
			return nil, nil
		}
	}
	z_nqmUMeJFH = doauF8Y.Yj5CZGxK1oow
	switch  /*line :1*/ len(z_nqmUMeJFH) {
	case IPv4len:
		if  /*line :1*/ len(kQ8_UEhxU) != IPv4len {
			return nil, nil
		}
	case IPv6len:
		if  /*line :1*/ len(kQ8_UEhxU) == IPv4len {
			z_nqmUMeJFH = z_nqmUMeJFH[12:]
		}
	default:
		return nil, nil
	}
	return
}


func (doauF8Y *M6TN4rba7) Contains(kQ8_UEhxU GNraIvYhB) bool {
	jBXrVi, z_nqmUMeJFH :=  /*line :1*/ li04uikLccY(doauF8Y)
	if tWO2olI :=  /*line :1*/ kQ8_UEhxU.To4(); tWO2olI != nil {
		kQ8_UEhxU = tWO2olI
	}
	v3upVm :=  /*line :1*/ len(kQ8_UEhxU)
	if v3upVm !=  /*line :1*/ len(jBXrVi) {
		return false
	}
	for eaAUaB7Z := 0; eaAUaB7Z < v3upVm; eaAUaB7Z++ {
		if jBXrVi[eaAUaB7Z]&z_nqmUMeJFH[eaAUaB7Z] != kQ8_UEhxU[eaAUaB7Z]&z_nqmUMeJFH[eaAUaB7Z] {
			return false
		}
	}
	return true
}


func (doauF8Y *M6TN4rba7) Network() string	{ return "ip+net" }







func (doauF8Y *M6TN4rba7) String() string {
	if doauF8Y == nil {
		return "<nil>"
	}
	jBXrVi, z_nqmUMeJFH :=  /*line :1*/ li04uikLccY(doauF8Y)
	if jBXrVi == nil || z_nqmUMeJFH == nil {
		return "<nil>"
	}
	v3upVm :=  /*line :1*/ eAtJ8K_5w1p(z_nqmUMeJFH)
	if v3upVm == -1 {
		return  /*line :1*/ jBXrVi.String() + "/" +  /*line :1*/ z_nqmUMeJFH.String()
	}
	return  /*line :1*/ jBXrVi.String() + "/" +  /*line :1*/ itoa.TDADPeM7Kft9( /*line :1*/ uint(v3upVm))
}






func XtVFDxFx(dRoFccaZ string) GNraIvYhB {
	if qxwkC3VxG0, nceNNEdciJ :=  /*line :1*/ icHp7pzD(dRoFccaZ); nceNNEdciJ {
		return  /*line :1*/ GNraIvYhB(qxwkC3VxG0[:])
	}
	return nil
}

func icHp7pzD(dRoFccaZ string) ([16]byte, bool) {
	kQ8_UEhxU, h_ljk48Bm :=  /*line :1*/ netip.Y1uXU4Sx(dRoFccaZ)
	if h_ljk48Bm != nil ||  /*line :1*/ kQ8_UEhxU.Zone() != "" {
		return [16]byte{}, false
	}
	return  /*line :1*/ kQ8_UEhxU.As16(), true
}









func ExX830jwSS(dRoFccaZ string) (GNraIvYhB, *M6TN4rba7, error) {
	eaAUaB7Z :=  /*line :1*/ bytealg.IndexByteString(dRoFccaZ, '/')
	if eaAUaB7Z < 0 {
		return nil, nil, &DPDxTDQNU3Ih{FATAgZbL0: "CIDR address", GrpQb_kayach: dRoFccaZ}
	}
	qxwkC3VxG0, vQyaaIll := dRoFccaZ[:eaAUaB7Z], dRoFccaZ[eaAUaB7Z+1:]

	qf12veP8, h_ljk48Bm :=  /*line :1*/ netip.Y1uXU4Sx(qxwkC3VxG0)
	if h_ljk48Bm != nil ||  /*line :1*/ qf12veP8.Zone() != "" {
		return nil, nil, &DPDxTDQNU3Ih{FATAgZbL0: "CIDR address", GrpQb_kayach: dRoFccaZ}
	}

	doauF8Y, eaAUaB7Z, d30HIwxht :=  /*line :1*/ bEihJjrbz(vQyaaIll)
	if !d30HIwxht || eaAUaB7Z !=  /*line :1*/ len(vQyaaIll) || doauF8Y < 0 || doauF8Y >  /*line :1*/ qf12veP8.BitLen() {
		return nil, nil, &DPDxTDQNU3Ih{FATAgZbL0: "CIDR address", GrpQb_kayach: dRoFccaZ}
	}
	z_nqmUMeJFH :=  /*line :1*/ RYiA8Y9b(doauF8Y,  /*line :1*/ qf12veP8.BitLen())
	ewXqgg6Geery :=  /*line :1*/ qf12veP8.As16()
	return  /*line :1*/ GNraIvYhB(ewXqgg6Geery[:]), &M6TN4rba7{RJ6BaPgcU:  /*line :1*/ GNraIvYhB(ewXqgg6Geery[:]).Mask(z_nqmUMeJFH), Yj5CZGxK1oow: z_nqmUMeJFH}, nil
}

func sNTdcemFI(tWO2olI GNraIvYhB) GNraIvYhB {
	clAtsIFT3 :=  /*line :1*/ make(GNraIvYhB,  /*line :1*/ len(tWO2olI))
	 /*line :1*/ copy(clAtsIFT3, tWO2olI)
	return clAtsIFT3
}
