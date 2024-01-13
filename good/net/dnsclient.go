//line :1
package gF9mZs2

import (
	"internal/bytealg"
	itoa "JkjxVS"
	sort "gzHaha55n"

	dnsmessage "FeKim_cPgM7z"
)


func eUeRcLk() uint

func gOhSWMo9m7A() int {
	return  /*line :1*/ int( /*line :1*/ eUeRcLk() >> 1)
}

func aIwNIUux4g(doauF8Y int) int {
	return  /*line :1*/ gOhSWMo9m7A() % doauF8Y
}




func sPSXZcPfQnVj(qxwkC3VxG0 string) (akvMHAa84xdt string, h_ljk48Bm error) {
	kQ8_UEhxU :=  /*line :1*/ XtVFDxFx(qxwkC3VxG0)
	if kQ8_UEhxU == nil {
		return "", &SoatTK0{PY3bIIR: "unrecognized address", FIvHQdTCAg: qxwkC3VxG0}
	}
	if  /*line :1*/ kQ8_UEhxU.To4() != nil {
		return  /*line :1*/ itoa.TDADPeM7Kft9( /*line :1*/ uint(kQ8_UEhxU[15])) + "." +  /*line :1*/ itoa.TDADPeM7Kft9( /*line :1*/ uint(kQ8_UEhxU[14])) + "." +  /*line :1*/ itoa.TDADPeM7Kft9( /*line :1*/ uint(kQ8_UEhxU[13])) + "." +  /*line :1*/ itoa.TDADPeM7Kft9( /*line :1*/ uint(kQ8_UEhxU[12])) + ".in-addr.arpa.", nil
	}

	zynKBqWHCa :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(kQ8_UEhxU)*4+ /*line :1*/ len("ip6.arpa."))

	for eaAUaB7Z :=  /*line :1*/ len(kQ8_UEhxU) - 1; eaAUaB7Z >= 0; eaAUaB7Z-- {
		ljsCSk := kQ8_UEhxU[eaAUaB7Z]
		zynKBqWHCa =  /*line :1*/ append(zynKBqWHCa, hexDigit[ljsCSk&0xF],
			'.',
			hexDigit[ljsCSk>>4],
			'.')
	}

	zynKBqWHCa =  /*line :1*/ append(zynKBqWHCa, "ip6.arpa."...)
	return  /*line :1*/ string(zynKBqWHCa), nil
}

func emSzQ2I(tWO2olI, clAtsIFT3 dnsmessage.Toccq2dE3) bool {
	if tWO2olI.W7uI8JuUD != clAtsIFT3.W7uI8JuUD {
		return false
	}
	for eaAUaB7Z := 0; eaAUaB7Z <  /*line :1*/ int(tWO2olI.W7uI8JuUD); eaAUaB7Z++ {
		uI7LZDHm6 := tWO2olI.MaSC7z[eaAUaB7Z]
		fIadEXIim6l := clAtsIFT3.MaSC7z[eaAUaB7Z]
		if 'A' <= uI7LZDHm6 && uI7LZDHm6 <= 'Z' {
			uI7LZDHm6 += 0x20
		}
		if 'A' <= fIadEXIim6l && fIadEXIim6l <= 'Z' {
			fIadEXIim6l += 0x20
		}
		if uI7LZDHm6 != fIadEXIim6l {
			return false
		}
	}
	return true
}




func tgYtaZX8VvFU(dRoFccaZ string) bool {

	if dRoFccaZ == "." {
		return true
	}

	v3upVm :=  /*line :1*/ len(dRoFccaZ)
	if v3upVm == 0 || v3upVm > 254 || v3upVm == 254 && dRoFccaZ[v3upVm-1] != '.' {
		return false
	}

	vMwVxD_U :=  /*line :1*/ byte('.')
	i3EvdntKuZ := false
	jbYWGp3IiPwn := 0
	for eaAUaB7Z := 0; eaAUaB7Z <  /*line :1*/ len(dRoFccaZ); eaAUaB7Z++ {
		hl8w2lFX := dRoFccaZ[eaAUaB7Z]
		switch {
		default:
			return false
		case 'a' <= hl8w2lFX && hl8w2lFX <= 'z' || 'A' <= hl8w2lFX && hl8w2lFX <= 'Z' || hl8w2lFX == '_':
			i3EvdntKuZ = true
			jbYWGp3IiPwn++
		case '0' <= hl8w2lFX && hl8w2lFX <= '9':

			jbYWGp3IiPwn++
		case hl8w2lFX == '-':

			if vMwVxD_U == '.' {
				return false
			}
			jbYWGp3IiPwn++
			i3EvdntKuZ = true
		case hl8w2lFX == '.':

			if vMwVxD_U == '.' || vMwVxD_U == '-' {
				return false
			}
			if jbYWGp3IiPwn > 63 || jbYWGp3IiPwn == 0 {
				return false
			}
			jbYWGp3IiPwn = 0
		}
		vMwVxD_U = hl8w2lFX
	}
	if vMwVxD_U == '-' || jbYWGp3IiPwn > 63 {
		return false
	}

	return i3EvdntKuZ
}









func ztv3vbC8VS(dRoFccaZ string) string {
	if  /*line :1*/ bytealg.IndexByteString(dRoFccaZ, '.') != -1 && dRoFccaZ[ /*line :1*/ len(dRoFccaZ)-1] != '.' {
		dRoFccaZ += "."
	}
	return dRoFccaZ
}


type LRAj4ahDtDg struct {
	SCFx8T2sDj	string
	Vb3Y6YRdZxkq	uint16
	H5wnRh	uint16
	BGuDNISZ9	uint16
}


type a_iF1ZK0 []*LRAj4ahDtDg

func (dRoFccaZ a_iF1ZK0) Len() int	{ return  /*line :1*/ len(dRoFccaZ) }
func (dRoFccaZ a_iF1ZK0) Less(eaAUaB7Z, kVgtfLFZ int) bool {
	return dRoFccaZ[eaAUaB7Z].H5wnRh < dRoFccaZ[kVgtfLFZ].H5wnRh || (dRoFccaZ[eaAUaB7Z].H5wnRh == dRoFccaZ[kVgtfLFZ].H5wnRh && dRoFccaZ[eaAUaB7Z].BGuDNISZ9 < dRoFccaZ[kVgtfLFZ].BGuDNISZ9)
}
func (dRoFccaZ a_iF1ZK0) Swap(eaAUaB7Z, kVgtfLFZ int) {
	dRoFccaZ[eaAUaB7Z], dRoFccaZ[kVgtfLFZ] = dRoFccaZ[kVgtfLFZ], dRoFccaZ[eaAUaB7Z]
}



func (md4QSRkO a_iF1ZK0) iMEv2SWhJcc1() {
	hL5b7I2hlcD := 0
	for _, qxwkC3VxG0 := range md4QSRkO {
		hL5b7I2hlcD +=  /*line :1*/ int(qxwkC3VxG0.BGuDNISZ9)
	}
	for hL5b7I2hlcD > 0 &&  /*line :1*/ len(md4QSRkO) > 1 {
		dRoFccaZ := 0
		doauF8Y :=  /*line :1*/ aIwNIUux4g(hL5b7I2hlcD)
		for eaAUaB7Z := range md4QSRkO {
			dRoFccaZ +=  /*line :1*/ int(md4QSRkO[eaAUaB7Z].BGuDNISZ9)
			if dRoFccaZ > doauF8Y {
				if eaAUaB7Z > 0 {
					md4QSRkO[0], md4QSRkO[eaAUaB7Z] = md4QSRkO[eaAUaB7Z], md4QSRkO[0]
				}
				break
			}
		}
		hL5b7I2hlcD -=  /*line :1*/ int(md4QSRkO[0].BGuDNISZ9)
		md4QSRkO = md4QSRkO[1:]
	}
}


func (md4QSRkO a_iF1ZK0) wYlOfm() {
	 /*line :1*/ sort.KX2G4Qz(md4QSRkO)
	eaAUaB7Z := 0
	for kVgtfLFZ := 1; kVgtfLFZ <  /*line :1*/ len(md4QSRkO); kVgtfLFZ++ {
		if md4QSRkO[eaAUaB7Z].H5wnRh != md4QSRkO[kVgtfLFZ].H5wnRh {
			 /*line :1*/ md4QSRkO[eaAUaB7Z:kVgtfLFZ].iMEv2SWhJcc1()
			eaAUaB7Z = kVgtfLFZ
		}
	}
	 /*line :1*/ md4QSRkO[eaAUaB7Z:].iMEv2SWhJcc1()
}


type McZ_wBAueqxz struct {
	BgGVZWj8	string
	VG6KZMLF88	uint16
}


type rcgRLRupL []*McZ_wBAueqxz

func (dRoFccaZ rcgRLRupL) Len() int	{ return  /*line :1*/ len(dRoFccaZ) }
func (dRoFccaZ rcgRLRupL) Less(eaAUaB7Z, kVgtfLFZ int) bool {
	return dRoFccaZ[eaAUaB7Z].VG6KZMLF88 < dRoFccaZ[kVgtfLFZ].VG6KZMLF88
}
func (dRoFccaZ rcgRLRupL) Swap(eaAUaB7Z, kVgtfLFZ int) {
	dRoFccaZ[eaAUaB7Z], dRoFccaZ[kVgtfLFZ] = dRoFccaZ[kVgtfLFZ], dRoFccaZ[eaAUaB7Z]
}


func (dRoFccaZ rcgRLRupL) wYlOfm() {
	for eaAUaB7Z := range dRoFccaZ {
		kVgtfLFZ :=  /*line :1*/ aIwNIUux4g(eaAUaB7Z + 1)
		dRoFccaZ[eaAUaB7Z], dRoFccaZ[kVgtfLFZ] = dRoFccaZ[kVgtfLFZ], dRoFccaZ[eaAUaB7Z]
	}
	 /*line :1*/ sort.KX2G4Qz(dRoFccaZ)
}


type AEmYkM1 struct {
	C1_6l1HrdN string
}
