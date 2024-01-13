//line :1







package iPdCam_KQOXj

import (
	errors "iAHoxjmM"
	math "math"
	strconv "zBESu2SrRjP"

	"internal/bytealg"
	intern "GWGwIQr9vXAy"
	itoa "JkjxVS"
)









type KFLQ1_1E struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	rPvius	rc_vdlQ

	
	
	
	
	
	
	
	nqiHqD	*intern.XGxgZFq
}



var (
	sl68hrb	= (* /*line :1*/ intern.XGxgZFq)(nil)
	iU_Gva	=  /*line :1*/ new(intern.XGxgZFq)
	sq5HIi	=  /*line :1*/ new(intern.XGxgZFq)
)



func QPTqgJv06_8() KFLQ1_1E	{ return  /*line :1*/ XtvbgBR([16]byte{0: 0xff, 1: 0x02, 15: 0x01}) }



func QoKLdaOIjO() KFLQ1_1E	{ return  /*line :1*/ XtvbgBR([16]byte{0: 0xff, 1: 0x02, 15: 0x02}) }


func NM1KGDb7L() KFLQ1_1E	{ return  /*line :1*/ XtvbgBR([16]byte{15: 0x01}) }


func V0kWRbdfW() KFLQ1_1E	{ return KFLQ1_1E{nqiHqD: sq5HIi} }


func U7g4loyO() KFLQ1_1E	{ return  /*line :1*/ Vh24QUEOnu([4]byte{}) }


func Vh24QUEOnu(gIbHM2XzK [4]byte) KFLQ1_1E {
	return KFLQ1_1E{
		rPvius:	rc_vdlQ{0, 0xffff00000000 |  /*line :1*/ uint64(gIbHM2XzK[0])<<24 |  /*line :1*/ uint64(gIbHM2XzK[1])<<16 |  /*line :1*/ uint64(gIbHM2XzK[2])<<8 |  /*line :1*/ uint64(gIbHM2XzK[3])},
		nqiHqD:	iU_Gva,
	}
}




func XtvbgBR(gIbHM2XzK [16]byte) KFLQ1_1E {
	return KFLQ1_1E{
		rPvius: rc_vdlQ{
			 /*line :1*/ w1dKJ7nPezc(gIbHM2XzK[:8]),
			 /*line :1*/ w1dKJ7nPezc(gIbHM2XzK[8:]),
		},
		nqiHqD:	sq5HIi,
	}
}




func Y1uXU4Sx(asil3aeiMbi3 string) (KFLQ1_1E, error) {
	for jVDSPyHrNzR := 0; jVDSPyHrNzR <  /*line :1*/ len(asil3aeiMbi3); jVDSPyHrNzR++ {
		switch asil3aeiMbi3[jVDSPyHrNzR] {
		case '.':
			return  /*line :1*/ rPSrRkNedw_l(asil3aeiMbi3)
		case ':':
			return  /*line :1*/ n7mOztQ(asil3aeiMbi3)
		case '%':

			return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "missing IPv6 address"}
		}
	}
	return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "unable to parse IP"}
}



func HWTZYAx(asil3aeiMbi3 string) KFLQ1_1E {
	hCc7zqG4, lcob4O :=  /*line :1*/ Y1uXU4Sx(asil3aeiMbi3)
	if lcob4O != nil {
		 /*line :1*/ panic(lcob4O)
	}
	return hCc7zqG4
}

type yFTH7W struct {
	leSksha3OA	string	
	esVnGRvG_	string	
	hToNtb8c	string	
}

func (lcob4O yFTH7W) Error() string {
	n1VyN2rsna := strconv.D35LD5nn
	if lcob4O.hToNtb8c != "" {
		return "ParseAddr(" +  /*line :1*/ n1VyN2rsna(lcob4O.leSksha3OA) + "): " + lcob4O.esVnGRvG_ + " (at " +  /*line :1*/ n1VyN2rsna(lcob4O.hToNtb8c) + ")"
	}
	return "ParseAddr(" +  /*line :1*/ n1VyN2rsna(lcob4O.leSksha3OA) + "): " + lcob4O.esVnGRvG_
}


func rPSrRkNedw_l(asil3aeiMbi3 string) (hCc7zqG4 KFLQ1_1E, lcob4O error) {
	var xmIYL5i2Wywj [4]uint8
	var fOnpoi, weR61Qw int
	var bToj2H6 int	
	for jVDSPyHrNzR := 0; jVDSPyHrNzR <  /*line :1*/ len(asil3aeiMbi3); jVDSPyHrNzR++ {
		if asil3aeiMbi3[jVDSPyHrNzR] >= '0' && asil3aeiMbi3[jVDSPyHrNzR] <= '9' {
			if bToj2H6 == 1 && fOnpoi == 0 {
				return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "IPv4 field has octet with leading zero"}
			}
			fOnpoi = fOnpoi*10 +  /*line :1*/ int(asil3aeiMbi3[jVDSPyHrNzR]) - '0'
			bToj2H6++
			if fOnpoi > 255 {
				return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "IPv4 field has value >255"}
			}
		} else if asil3aeiMbi3[jVDSPyHrNzR] == '.' {

			if jVDSPyHrNzR == 0 || jVDSPyHrNzR ==  /*line :1*/ len(asil3aeiMbi3)-1 || asil3aeiMbi3[jVDSPyHrNzR-1] == '.' {
				return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "IPv4 field must have at least one digit", hToNtb8c: asil3aeiMbi3[jVDSPyHrNzR:]}
			}

			if weR61Qw == 3 {
				return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "IPv4 address too long"}
			}
			xmIYL5i2Wywj[weR61Qw] =  /*line :1*/ uint8(fOnpoi)
			weR61Qw++
			fOnpoi = 0
			bToj2H6 = 0
		} else {
			return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "unexpected character", hToNtb8c: asil3aeiMbi3[jVDSPyHrNzR:]}
		}
	}
	if weR61Qw < 3 {
		return KFLQ1_1E{}, yFTH7W{leSksha3OA: asil3aeiMbi3, esVnGRvG_: "IPv4 address too short"}
	}
	xmIYL5i2Wywj[3] =  /*line :1*/ uint8(fOnpoi)
	return  /*line :1*/ Vh24QUEOnu(xmIYL5i2Wywj), nil
}


func n7mOztQ(bK22jkOWK8 string) (KFLQ1_1E, error) {
	asil3aeiMbi3 := bK22jkOWK8

	qwUN01iIFPek := ""
	jVDSPyHrNzR :=  /*line :1*/ bytealg.IndexByteString(asil3aeiMbi3, '%')
	if jVDSPyHrNzR != -1 {
		asil3aeiMbi3, qwUN01iIFPek = asil3aeiMbi3[:jVDSPyHrNzR], asil3aeiMbi3[jVDSPyHrNzR+1:]
		if qwUN01iIFPek == "" {

			return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "zone must be a non-empty string"}
		}
	}

	var hCc7zqG4 [16]byte
	iSTK_G7bOii1 := -1

	if  /*line :1*/ len(asil3aeiMbi3) >= 2 && asil3aeiMbi3[0] == ':' && asil3aeiMbi3[1] == ':' {
		iSTK_G7bOii1 = 0
		asil3aeiMbi3 = asil3aeiMbi3[2:]

		if  /*line :1*/ len(asil3aeiMbi3) == 0 {
			return  /*line :1*/ V0kWRbdfW().WithZone(qwUN01iIFPek), nil
		}
	}

	jVDSPyHrNzR = 0
	for jVDSPyHrNzR < 16 {

		bisirqaaaa := 0
		d8gOmR :=  /*line :1*/ uint32(0)
		for ; bisirqaaaa <  /*line :1*/ len(asil3aeiMbi3); bisirqaaaa++ {
			gbgAM0ACKu := asil3aeiMbi3[bisirqaaaa]
			if gbgAM0ACKu >= '0' && gbgAM0ACKu <= '9' {
				d8gOmR = (d8gOmR << 4) +  /*line :1*/ uint32(gbgAM0ACKu-'0')
			} else if gbgAM0ACKu >= 'a' && gbgAM0ACKu <= 'f' {
				d8gOmR = (d8gOmR << 4) +  /*line :1*/ uint32(gbgAM0ACKu-'a'+10)
			} else if gbgAM0ACKu >= 'A' && gbgAM0ACKu <= 'F' {
				d8gOmR = (d8gOmR << 4) +  /*line :1*/ uint32(gbgAM0ACKu-'A'+10)
			} else {
				break
			}
			if d8gOmR > math.MaxUint16 {

				return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "IPv6 field has value >=2^16", hToNtb8c: asil3aeiMbi3}
			}
		}
		if bisirqaaaa == 0 {

			return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "each colon-separated field must have at least one digit", hToNtb8c: asil3aeiMbi3}
		}

		if bisirqaaaa <  /*line :1*/ len(asil3aeiMbi3) && asil3aeiMbi3[bisirqaaaa] == '.' {
			if iSTK_G7bOii1 < 0 && jVDSPyHrNzR != 12 {

				return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "embedded IPv4 address must replace the final 2 fields of the address", hToNtb8c: asil3aeiMbi3}
			}
			if jVDSPyHrNzR+4 > 16 {

				return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "too many hex fields to fit an embedded IPv4 at the end of the address", hToNtb8c: asil3aeiMbi3}
			}

			ko_U3q, lcob4O :=  /*line :1*/ rPSrRkNedw_l(asil3aeiMbi3)
			if lcob4O != nil {
				return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_:  /*line :1*/ lcob4O.Error(), hToNtb8c: asil3aeiMbi3}
			}
			hCc7zqG4[jVDSPyHrNzR] =  /*line :1*/ ko_U3q.dVrOwjnb(0)
			hCc7zqG4[jVDSPyHrNzR+1] =  /*line :1*/ ko_U3q.dVrOwjnb(1)
			hCc7zqG4[jVDSPyHrNzR+2] =  /*line :1*/ ko_U3q.dVrOwjnb(2)
			hCc7zqG4[jVDSPyHrNzR+3] =  /*line :1*/ ko_U3q.dVrOwjnb(3)
			asil3aeiMbi3 = ""
			jVDSPyHrNzR += 4
			break
		}

		hCc7zqG4[jVDSPyHrNzR] =  /*line :1*/ byte(d8gOmR >> 8)
		hCc7zqG4[jVDSPyHrNzR+1] =  /*line :1*/ byte(d8gOmR)
		jVDSPyHrNzR += 2

		asil3aeiMbi3 = asil3aeiMbi3[bisirqaaaa:]
		if  /*line :1*/ len(asil3aeiMbi3) == 0 {
			break
		}

		if asil3aeiMbi3[0] != ':' {
			return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "unexpected character, want colon", hToNtb8c: asil3aeiMbi3}
		} else if  /*line :1*/ len(asil3aeiMbi3) == 1 {
			return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "colon must be followed by more characters", hToNtb8c: asil3aeiMbi3}
		}
		asil3aeiMbi3 = asil3aeiMbi3[1:]

		if asil3aeiMbi3[0] == ':' {
			if iSTK_G7bOii1 >= 0 {
				return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "multiple :: in address", hToNtb8c: asil3aeiMbi3}
			}
			iSTK_G7bOii1 = jVDSPyHrNzR
			asil3aeiMbi3 = asil3aeiMbi3[1:]
			if  /*line :1*/ len(asil3aeiMbi3) == 0 {
				break
			}
		}
	}

	if  /*line :1*/ len(asil3aeiMbi3) != 0 {
		return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "trailing garbage after address", hToNtb8c: asil3aeiMbi3}
	}

	if jVDSPyHrNzR < 16 {
		if iSTK_G7bOii1 < 0 {
			return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "address string too short"}
		}
		xUVyHJVSZ4t7 := 16 - jVDSPyHrNzR
		for wVaBxKkn := jVDSPyHrNzR - 1; wVaBxKkn >= iSTK_G7bOii1; wVaBxKkn-- {
			hCc7zqG4[wVaBxKkn+xUVyHJVSZ4t7] = hCc7zqG4[wVaBxKkn]
		}
		for wVaBxKkn := iSTK_G7bOii1 + xUVyHJVSZ4t7 - 1; wVaBxKkn >= iSTK_G7bOii1; wVaBxKkn-- {
			hCc7zqG4[wVaBxKkn] = 0
		}
	} else if iSTK_G7bOii1 >= 0 {

		return KFLQ1_1E{}, yFTH7W{leSksha3OA: bK22jkOWK8, esVnGRvG_: "the :: must expand to at least one field of zeros"}
	}
	return  /*line :1*/ XtvbgBR(hCc7zqG4).WithZone(qwUN01iIFPek), nil
}




func OWcivL5ykj9(ltUTaZao []byte) (hCc7zqG4 KFLQ1_1E, eAJHa7 bool) {
	switch  /*line :1*/ len(ltUTaZao) {
	case 4:
		return  /*line :1*/ Vh24QUEOnu([4] /*line :1*/ byte(ltUTaZao)), true
	case 16:
		return  /*line :1*/ XtvbgBR([16] /*line :1*/ byte(ltUTaZao)), true
	}
	return KFLQ1_1E{}, false
}



func (hCc7zqG4 KFLQ1_1E) dVrOwjnb(jVDSPyHrNzR uint8) uint8 {
	return  /*line :1*/ uint8(hCc7zqG4.rPvius.gRGNmk >> ((3 - jVDSPyHrNzR) * 8))
}



func (hCc7zqG4 KFLQ1_1E) jpLaz9VtZ(jVDSPyHrNzR uint8) uint8 {
	return  /*line :1*/ uint8(*( /*line :1*/ hCc7zqG4.rPvius.pxcoT3wb()[(jVDSPyHrNzR/8)%2]) >> ((7 - jVDSPyHrNzR%8) * 8))
}



func (hCc7zqG4 KFLQ1_1E) nRoQ4UZthJ_(jVDSPyHrNzR uint8) uint16 {
	return  /*line :1*/ uint16(*( /*line :1*/ hCc7zqG4.rPvius.pxcoT3wb()[(jVDSPyHrNzR/4)%2]) >> ((3 - jVDSPyHrNzR%4) * 16))
}






func (hCc7zqG4 KFLQ1_1E) rb6oFzFeRc() bool {

	return hCc7zqG4.nqiHqD == sl68hrb
}




func (hCc7zqG4 KFLQ1_1E) IsValid() bool	{ return hCc7zqG4.nqiHqD != sl68hrb }






func (hCc7zqG4 KFLQ1_1E) BitLen() int {
	switch hCc7zqG4.nqiHqD {
	case sl68hrb:
		return 0
	case iU_Gva:
		return 32
	}
	return 128
}


func (hCc7zqG4 KFLQ1_1E) Zone() string {
	if hCc7zqG4.nqiHqD == nil {
		return ""
	}
	qwUN01iIFPek, _ :=  /*line :1*/ hCc7zqG4.nqiHqD.Get().(string)
	return qwUN01iIFPek
}




func (hCc7zqG4 KFLQ1_1E) Compare(beVPLa9Pc3d KFLQ1_1E) int {
	d5w3i5po8vTB, p9txU5i0Y :=  /*line :1*/ hCc7zqG4.BitLen(),  /*line :1*/ beVPLa9Pc3d.BitLen()
	if d5w3i5po8vTB < p9txU5i0Y {
		return -1
	}
	if d5w3i5po8vTB > p9txU5i0Y {
		return 1
	}
	uTANfenSs, xWbPvyH := hCc7zqG4.rPvius.od2_Mw, beVPLa9Pc3d.rPvius.od2_Mw
	if uTANfenSs < xWbPvyH {
		return -1
	}
	if uTANfenSs > xWbPvyH {
		return 1
	}
	pEa1fZA0D, ky2D7GnSC5 := hCc7zqG4.rPvius.gRGNmk, beVPLa9Pc3d.rPvius.gRGNmk
	if pEa1fZA0D < ky2D7GnSC5 {
		return -1
	}
	if pEa1fZA0D > ky2D7GnSC5 {
		return 1
	}
	if  /*line :1*/ hCc7zqG4.Is6() {
		dZOgh2, yVRgOQQr8z0 :=  /*line :1*/ hCc7zqG4.Zone(),  /*line :1*/ beVPLa9Pc3d.Zone()
		if dZOgh2 < yVRgOQQr8z0 {
			return -1
		}
		if dZOgh2 > yVRgOQQr8z0 {
			return 1
		}
	}
	return 0
}




func (hCc7zqG4 KFLQ1_1E) Less(beVPLa9Pc3d KFLQ1_1E) bool	{ return  /*line :1*/ hCc7zqG4.Compare(beVPLa9Pc3d) == -1 }




func (hCc7zqG4 KFLQ1_1E) Is4() bool {
	return hCc7zqG4.nqiHqD == iU_Gva
}


func (hCc7zqG4 KFLQ1_1E) Is4In6() bool {
	return  /*line :1*/ hCc7zqG4.Is6() && hCc7zqG4.rPvius.od2_Mw == 0 && hCc7zqG4.rPvius.gRGNmk>>32 == 0xffff
}



func (hCc7zqG4 KFLQ1_1E) Is6() bool {
	return hCc7zqG4.nqiHqD != sl68hrb && hCc7zqG4.nqiHqD != iU_Gva
}





func (hCc7zqG4 KFLQ1_1E) Unmap() KFLQ1_1E {
	if  /*line :1*/ hCc7zqG4.Is4In6() {
		hCc7zqG4.nqiHqD = iU_Gva
	}
	return hCc7zqG4
}




func (hCc7zqG4 KFLQ1_1E) WithZone(qwUN01iIFPek string) KFLQ1_1E {
	if ! /*line :1*/ hCc7zqG4.Is6() {
		return hCc7zqG4
	}
	if qwUN01iIFPek == "" {
		hCc7zqG4.nqiHqD = sq5HIi
		return hCc7zqG4
	}
	hCc7zqG4.nqiHqD =  /*line :1*/ intern.DocbIWHFe(qwUN01iIFPek)
	return hCc7zqG4
}



func (hCc7zqG4 KFLQ1_1E) diJCELqd3r() KFLQ1_1E {
	if ! /*line :1*/ hCc7zqG4.Is6() {
		return hCc7zqG4
	}
	hCc7zqG4.nqiHqD = sq5HIi
	return hCc7zqG4
}


func (hCc7zqG4 KFLQ1_1E) _oaAHhqH6AM() bool {
	return hCc7zqG4.nqiHqD != sl68hrb && hCc7zqG4.nqiHqD != iU_Gva && hCc7zqG4.nqiHqD != sq5HIi
}


func (hCc7zqG4 KFLQ1_1E) IsLinkLocalUnicast() bool {

	if  /*line :1*/ hCc7zqG4.Is4() {
		return  /*line :1*/ hCc7zqG4.dVrOwjnb(0) == 169 &&  /*line :1*/ hCc7zqG4.dVrOwjnb(1) == 254
	}

	if  /*line :1*/ hCc7zqG4.Is6() {
		return  /*line :1*/ hCc7zqG4.nRoQ4UZthJ_(0)&0xffc0 == 0xfe80
	}
	return false
}


func (hCc7zqG4 KFLQ1_1E) IsLoopback() bool {

	if  /*line :1*/ hCc7zqG4.Is4() {
		return  /*line :1*/ hCc7zqG4.dVrOwjnb(0) == 127
	}

	if  /*line :1*/ hCc7zqG4.Is6() {
		return hCc7zqG4.rPvius.od2_Mw == 0 && hCc7zqG4.rPvius.gRGNmk == 1
	}
	return false
}


func (hCc7zqG4 KFLQ1_1E) IsMulticast() bool {

	if  /*line :1*/ hCc7zqG4.Is4() {
		return  /*line :1*/ hCc7zqG4.dVrOwjnb(0)&0xf0 == 0xe0
	}

	if  /*line :1*/ hCc7zqG4.Is6() {
		return hCc7zqG4.rPvius.od2_Mw>>(64-8) == 0xff
	}
	return false
}



func (hCc7zqG4 KFLQ1_1E) IsInterfaceLocalMulticast() bool {

	if  /*line :1*/ hCc7zqG4.Is6() {
		return  /*line :1*/ hCc7zqG4.nRoQ4UZthJ_(0)&0xff0f == 0xff01
	}
	return false
}


func (hCc7zqG4 KFLQ1_1E) IsLinkLocalMulticast() bool {

	if  /*line :1*/ hCc7zqG4.Is4() {
		return  /*line :1*/ hCc7zqG4.dVrOwjnb(0) == 224 &&  /*line :1*/ hCc7zqG4.dVrOwjnb(1) == 0 &&  /*line :1*/ hCc7zqG4.dVrOwjnb(2) == 0
	}

	if  /*line :1*/ hCc7zqG4.Is6() {
		return  /*line :1*/ hCc7zqG4.nRoQ4UZthJ_(0)&0xff0f == 0xff02
	}
	return false
}










func (hCc7zqG4 KFLQ1_1E) IsGlobalUnicast() bool {
	if hCc7zqG4.nqiHqD == sl68hrb {

		return false
	}

	if  /*line :1*/ hCc7zqG4.Is4() && (hCc7zqG4 ==  /*line :1*/ U7g4loyO() || hCc7zqG4 ==  /*line :1*/ Vh24QUEOnu([4]byte{255, 255, 255, 255})) {
		return false
	}

	return hCc7zqG4 !=  /*line :1*/ V0kWRbdfW() &&
		! /*line :1*/ hCc7zqG4.IsLoopback() &&
		! /*line :1*/ hCc7zqG4.IsMulticast() &&
		! /*line :1*/ hCc7zqG4.IsLinkLocalUnicast()
}





func (hCc7zqG4 KFLQ1_1E) IsPrivate() bool {

	if  /*line :1*/ hCc7zqG4.Is4() {

		return  /*line :1*/ hCc7zqG4.dVrOwjnb(0) == 10 ||
			( /*line :1*/ hCc7zqG4.dVrOwjnb(0) == 172 &&  /*line :1*/ hCc7zqG4.dVrOwjnb(1)&0xf0 == 16) ||
			( /*line :1*/ hCc7zqG4.dVrOwjnb(0) == 192 &&  /*line :1*/ hCc7zqG4.dVrOwjnb(1) == 168)
	}

	if  /*line :1*/ hCc7zqG4.Is6() {

		return  /*line :1*/ hCc7zqG4.jpLaz9VtZ(0)&0xfe == 0xfc
	}

	return false
}





func (hCc7zqG4 KFLQ1_1E) IsUnspecified() bool {
	return hCc7zqG4 ==  /*line :1*/ U7g4loyO() || hCc7zqG4 ==  /*line :1*/ V0kWRbdfW()
}






func (hCc7zqG4 KFLQ1_1E) Prefix(eVndpt int) (We_DLeUDXHez, error) {
	if eVndpt < 0 {
		return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("negative Prefix bits")
	}
	xSmvI1Zl := eVndpt
	switch hCc7zqG4.nqiHqD {
	case sl68hrb:
		return We_DLeUDXHez{}, nil
	case iU_Gva:
		if eVndpt > 32 {
			return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("prefix length " +  /*line :1*/ itoa.V25ba9G5(eVndpt) + " too large for IPv4")
		}
		xSmvI1Zl += 96
	default:
		if eVndpt > 128 {
			return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("prefix length " +  /*line :1*/ itoa.V25ba9G5(eVndpt) + " too large for IPv6")
		}
	}
	hCc7zqG4.rPvius =  /*line :1*/ hCc7zqG4.rPvius.huLWxs63VG1( /*line :1*/ vtgBPgPqR37(xSmvI1Zl))
	return  /*line :1*/ QkGHFQ3cd(hCc7zqG4, eVndpt), nil
}

const (
	netIPv4len	= 4
	netIPv6len	= 16
)






func (hCc7zqG4 KFLQ1_1E) As16() (fDCsGX8 [16]byte) {
	 /*line :1*/ ozuLeayIadb(fDCsGX8[:8], hCc7zqG4.rPvius.od2_Mw)
	 /*line :1*/ ozuLeayIadb(fDCsGX8[8:], hCc7zqG4.rPvius.gRGNmk)
	return fDCsGX8
}




func (hCc7zqG4 KFLQ1_1E) As4() (j4Ecbz3zM [4]byte) {
	if hCc7zqG4.nqiHqD == iU_Gva ||  /*line :1*/ hCc7zqG4.Is4In6() {
		 /*line :1*/ mMagzWU1Bv(j4Ecbz3zM[:],  /*line :1*/ uint32(hCc7zqG4.rPvius.gRGNmk))
		return j4Ecbz3zM
	}
	if hCc7zqG4.nqiHqD == sl68hrb {
		 /*line :1*/ panic("As4 called on IP zero value")
	}
	 /*line :1*/ panic("As4 called on IPv6 address")
}


func (hCc7zqG4 KFLQ1_1E) AsSlice() []byte {
	switch hCc7zqG4.nqiHqD {
	case sl68hrb:
		return nil
	case iU_Gva:
		var uLm1AL12RGl [4]byte
		 /*line :1*/ mMagzWU1Bv(uLm1AL12RGl[:],  /*line :1*/ uint32(hCc7zqG4.rPvius.gRGNmk))
		return uLm1AL12RGl[:]
	default:
		var uLm1AL12RGl [16]byte
		 /*line :1*/ ozuLeayIadb(uLm1AL12RGl[:8], hCc7zqG4.rPvius.od2_Mw)
		 /*line :1*/ ozuLeayIadb(uLm1AL12RGl[8:], hCc7zqG4.rPvius.gRGNmk)
		return uLm1AL12RGl[:]
	}
}



func (hCc7zqG4 KFLQ1_1E) Next() KFLQ1_1E {
	hCc7zqG4.rPvius =  /*line :1*/ hCc7zqG4.rPvius.ocX5kXg()
	if  /*line :1*/ hCc7zqG4.Is4() {
		if  /*line :1*/ uint32(hCc7zqG4.rPvius.gRGNmk) == 0 {

			return KFLQ1_1E{}
		}
	} else {
		if  /*line :1*/ hCc7zqG4.rPvius.rb6oFzFeRc() {

			return KFLQ1_1E{}
		}
	}
	return hCc7zqG4
}



func (hCc7zqG4 KFLQ1_1E) Prev() KFLQ1_1E {
	if  /*line :1*/ hCc7zqG4.Is4() {
		if  /*line :1*/ uint32(hCc7zqG4.rPvius.gRGNmk) == 0 {
			return KFLQ1_1E{}
		}
	} else if  /*line :1*/ hCc7zqG4.rPvius.rb6oFzFeRc() {
		return KFLQ1_1E{}
	}
	hCc7zqG4.rPvius =  /*line :1*/ hCc7zqG4.rPvius.cgATgys_b7Y()
	return hCc7zqG4
}













func (hCc7zqG4 KFLQ1_1E) String() string {
	switch hCc7zqG4.nqiHqD {
	case sl68hrb:
		return "invalid IP"
	case iU_Gva:
		return  /*line :1*/ hCc7zqG4.mIHA5iz4F()
	default:
		if  /*line :1*/ hCc7zqG4.Is4In6() {
			if iwlojlVcD :=  /*line :1*/ hCc7zqG4.Zone(); iwlojlVcD != "" {
				return "::ffff:" +  /*line :1*/ hCc7zqG4.Unmap().mIHA5iz4F() + "%" + iwlojlVcD
			} else {
				return "::ffff:" +  /*line :1*/ hCc7zqG4.Unmap().mIHA5iz4F()
			}
		}
		return  /*line :1*/ hCc7zqG4.qWnatllCLV()
	}
}




func (hCc7zqG4 KFLQ1_1E) AppendTo(eVndpt []byte) []byte {
	switch hCc7zqG4.nqiHqD {
	case sl68hrb:
		return eVndpt
	case iU_Gva:
		return  /*line :1*/ hCc7zqG4.fythzeXU8(eVndpt)
	default:
		if  /*line :1*/ hCc7zqG4.Is4In6() {
			eVndpt =  /*line :1*/ append(eVndpt, "::ffff:"...)
			eVndpt =  /*line :1*/ hCc7zqG4.Unmap().fythzeXU8(eVndpt)
			if iwlojlVcD :=  /*line :1*/ hCc7zqG4.Zone(); iwlojlVcD != "" {
				eVndpt =  /*line :1*/ append(eVndpt, '%')
				eVndpt =  /*line :1*/ append(eVndpt, iwlojlVcD...)
			}
			return eVndpt
		}
		return  /*line :1*/ hCc7zqG4.d2LQe2i(eVndpt)
	}
}



const digits = "0123456789abcdef"


func rM7DocatE(eVndpt []byte, acHaCAXeUHY uint8) []byte {

	if acHaCAXeUHY >= 100 {
		eVndpt =  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY/100])
	}
	if acHaCAXeUHY >= 10 {
		eVndpt =  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY/10%10])
	}
	return  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY%10])
}


func kU3O9Wfml(eVndpt []byte, acHaCAXeUHY uint16) []byte {

	if acHaCAXeUHY >= 0x1000 {
		eVndpt =  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY>>12])
	}
	if acHaCAXeUHY >= 0x100 {
		eVndpt =  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY>>8&0xf])
	}
	if acHaCAXeUHY >= 0x10 {
		eVndpt =  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY>>4&0xf])
	}
	return  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY&0xf])
}


func cw7x3H(eVndpt []byte, acHaCAXeUHY uint16) []byte {
	return  /*line :1*/ append(eVndpt, digits[acHaCAXeUHY>>12], digits[acHaCAXeUHY>>8&0xf], digits[acHaCAXeUHY>>4&0xf], digits[acHaCAXeUHY&0xf])
}

func (hCc7zqG4 KFLQ1_1E) mIHA5iz4F() string {
	const max =  /*line :1*/ len("255.255.255.255")
	uLm1AL12RGl :=  /*line :1*/ make([]byte, 0, max)
	uLm1AL12RGl =  /*line :1*/ hCc7zqG4.fythzeXU8(uLm1AL12RGl)
	return  /*line :1*/ string(uLm1AL12RGl)
}

func (hCc7zqG4 KFLQ1_1E) fythzeXU8(uLm1AL12RGl []byte) []byte {
	uLm1AL12RGl =  /*line :1*/ rM7DocatE(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.dVrOwjnb(0))
	uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, '.')
	uLm1AL12RGl =  /*line :1*/ rM7DocatE(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.dVrOwjnb(1))
	uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, '.')
	uLm1AL12RGl =  /*line :1*/ rM7DocatE(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.dVrOwjnb(2))
	uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, '.')
	uLm1AL12RGl =  /*line :1*/ rM7DocatE(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.dVrOwjnb(3))
	return uLm1AL12RGl
}






func (hCc7zqG4 KFLQ1_1E) qWnatllCLV() string {
	
	
	
	
	
	
	
	const max =  /*line :1*/ len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0")
	uLm1AL12RGl :=  /*line :1*/ make([]byte, 0, max)
	uLm1AL12RGl =  /*line :1*/ hCc7zqG4.d2LQe2i(uLm1AL12RGl)
	return  /*line :1*/ string(uLm1AL12RGl)
}

func (hCc7zqG4 KFLQ1_1E) d2LQe2i(uLm1AL12RGl []byte) []byte {
	z2Rcb1Ya4Oa5, ewuaxk :=  /*line :1*/ uint8(255),  /*line :1*/ uint8(255)
	for jVDSPyHrNzR :=  /*line :1*/ uint8(0); jVDSPyHrNzR < 8; jVDSPyHrNzR++ {
		wVaBxKkn := jVDSPyHrNzR
		for wVaBxKkn < 8 &&  /*line :1*/ hCc7zqG4.nRoQ4UZthJ_(wVaBxKkn) == 0 {
			wVaBxKkn++
		}
		if _zEK7cSdG := wVaBxKkn - jVDSPyHrNzR; _zEK7cSdG >= 2 && _zEK7cSdG > ewuaxk-z2Rcb1Ya4Oa5 {
			z2Rcb1Ya4Oa5, ewuaxk = jVDSPyHrNzR, wVaBxKkn
		}
	}

	for jVDSPyHrNzR :=  /*line :1*/ uint8(0); jVDSPyHrNzR < 8; jVDSPyHrNzR++ {
		if jVDSPyHrNzR == z2Rcb1Ya4Oa5 {
			uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, ':', ':')
			jVDSPyHrNzR = ewuaxk
			if jVDSPyHrNzR >= 8 {
				break
			}
		} else if jVDSPyHrNzR > 0 {
			uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, ':')
		}

		uLm1AL12RGl =  /*line :1*/ kU3O9Wfml(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.nRoQ4UZthJ_(jVDSPyHrNzR))
	}

	if hCc7zqG4.nqiHqD != sq5HIi {
		uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, '%')
		uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.Zone()...)
	}
	return uLm1AL12RGl
}




func (hCc7zqG4 KFLQ1_1E) StringExpanded() string {
	switch hCc7zqG4.nqiHqD {
	case sl68hrb, iU_Gva:
		return  /*line :1*/ hCc7zqG4.String()
	}

	const size =  /*line :1*/ len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	uLm1AL12RGl :=  /*line :1*/ make([]byte, 0, size)
	for jVDSPyHrNzR :=  /*line :1*/ uint8(0); jVDSPyHrNzR < 8; jVDSPyHrNzR++ {
		if jVDSPyHrNzR > 0 {
			uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, ':')
		}

		uLm1AL12RGl =  /*line :1*/ cw7x3H(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.nRoQ4UZthJ_(jVDSPyHrNzR))
	}

	if hCc7zqG4.nqiHqD != sq5HIi {

		uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl, '%')
		uLm1AL12RGl =  /*line :1*/ append(uLm1AL12RGl,  /*line :1*/ hCc7zqG4.Zone()...)
	}
	return  /*line :1*/ string(uLm1AL12RGl)
}




func (hCc7zqG4 KFLQ1_1E) MarshalText() ([]byte, error) {
	switch hCc7zqG4.nqiHqD {
	case sl68hrb:
		return [] /*line :1*/ byte(""), nil
	case iU_Gva:
		fnexDYLqfa1q :=  /*line :1*/ len("255.255.255.255")
		eVndpt :=  /*line :1*/ make([]byte, 0, fnexDYLqfa1q)
		return  /*line :1*/ hCc7zqG4.fythzeXU8(eVndpt), nil
	default:
		fnexDYLqfa1q :=  /*line :1*/ len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0")
		eVndpt :=  /*line :1*/ make([]byte, 0, fnexDYLqfa1q)
		if  /*line :1*/ hCc7zqG4.Is4In6() {
			eVndpt =  /*line :1*/ append(eVndpt, "::ffff:"...)
			eVndpt =  /*line :1*/ hCc7zqG4.Unmap().fythzeXU8(eVndpt)
			if iwlojlVcD :=  /*line :1*/ hCc7zqG4.Zone(); iwlojlVcD != "" {
				eVndpt =  /*line :1*/ append(eVndpt, '%')
				eVndpt =  /*line :1*/ append(eVndpt, iwlojlVcD...)
			}
			return eVndpt, nil
		}
		return  /*line :1*/ hCc7zqG4.d2LQe2i(eVndpt), nil
	}

}






func (hCc7zqG4 *KFLQ1_1E) UnmarshalText(tefqztg []byte) error {
	if  /*line :1*/ len(tefqztg) == 0 {
		*hCc7zqG4 = KFLQ1_1E{}
		return nil
	}
	var lcob4O error
	*hCc7zqG4, lcob4O =  /*line :1*/ Y1uXU4Sx( /*line :1*/ string(tefqztg))
	return lcob4O
}

func (hCc7zqG4 KFLQ1_1E) eEfDsFeFui01(ysZVQHX int) []byte {
	var eVndpt []byte
	switch hCc7zqG4.nqiHqD {
	case sl68hrb:
		eVndpt =  /*line :1*/ make([]byte, ysZVQHX)
	case iU_Gva:
		eVndpt =  /*line :1*/ make([]byte, 4+ysZVQHX)
		 /*line :1*/ mMagzWU1Bv(eVndpt,  /*line :1*/ uint32(hCc7zqG4.rPvius.gRGNmk))
	default:
		iwlojlVcD :=  /*line :1*/ hCc7zqG4.Zone()
		eVndpt =  /*line :1*/ make([]byte, 16+ /*line :1*/ len(iwlojlVcD)+ysZVQHX)
		 /*line :1*/ ozuLeayIadb(eVndpt[:8], hCc7zqG4.rPvius.od2_Mw)
		 /*line :1*/ ozuLeayIadb(eVndpt[8:], hCc7zqG4.rPvius.gRGNmk)
		 /*line :1*/ copy(eVndpt[16:], iwlojlVcD)
	}
	return eVndpt
}





func (hCc7zqG4 KFLQ1_1E) MarshalBinary() ([]byte, error) {
	return  /*line :1*/ hCc7zqG4.eEfDsFeFui01(0), nil
}



func (hCc7zqG4 *KFLQ1_1E) UnmarshalBinary(eVndpt []byte) error {
	xUVyHJVSZ4t7 :=  /*line :1*/ len(eVndpt)
	switch {
	case xUVyHJVSZ4t7 == 0:
		*hCc7zqG4 = KFLQ1_1E{}
		return nil
	case xUVyHJVSZ4t7 == 4:
		*hCc7zqG4 =  /*line :1*/ Vh24QUEOnu([4] /*line :1*/ byte(eVndpt))
		return nil
	case xUVyHJVSZ4t7 == 16:
		*hCc7zqG4 =  /*line :1*/ XtvbgBR([16] /*line :1*/ byte(eVndpt))
		return nil
	case xUVyHJVSZ4t7 > 16:
		*hCc7zqG4 =  /*line :1*/ XtvbgBR([16] /*line :1*/ byte(eVndpt[:16])).WithZone( /*line :1*/ string(eVndpt[16:]))
		return nil
	}
	return  /*line :1*/ errors.Su6g6hRoi9X("unexpected slice size")
}


type YTqTf_4VC struct {
	rjl3GFNJ2JN2	KFLQ1_1E
	uqZ4iEGn0_u	uint16
}



func Iqr8_1pEDmw0(hCc7zqG4 KFLQ1_1E, i7pg41eiHlWD uint16) YTqTf_4VC {
	return YTqTf_4VC{rjl3GFNJ2JN2: hCc7zqG4, uqZ4iEGn0_u: i7pg41eiHlWD}
}


func (ceqvLqr YTqTf_4VC) Addr() KFLQ1_1E	{ return ceqvLqr.rjl3GFNJ2JN2 }


func (ceqvLqr YTqTf_4VC) Port() uint16	{ return ceqvLqr.uqZ4iEGn0_u }






func jxDvHzBX(asil3aeiMbi3 string) (hCc7zqG4, i7pg41eiHlWD string, jpLaz9VtZ bool, lcob4O error) {
	jVDSPyHrNzR :=  /*line :1*/ ghfJ9faupLy(asil3aeiMbi3, ':')
	if jVDSPyHrNzR == -1 {
		return "", "", false,  /*line :1*/ errors.Su6g6hRoi9X("not an ip:port")
	}

	hCc7zqG4, i7pg41eiHlWD = asil3aeiMbi3[:jVDSPyHrNzR], asil3aeiMbi3[jVDSPyHrNzR+1:]
	if  /*line :1*/ len(hCc7zqG4) == 0 {
		return "", "", false,  /*line :1*/ errors.Su6g6hRoi9X("no IP")
	}
	if  /*line :1*/ len(i7pg41eiHlWD) == 0 {
		return "", "", false,  /*line :1*/ errors.Su6g6hRoi9X("no port")
	}
	if hCc7zqG4[0] == '[' {
		if  /*line :1*/ len(hCc7zqG4) < 2 || hCc7zqG4[ /*line :1*/ len(hCc7zqG4)-1] != ']' {
			return "", "", false,  /*line :1*/ errors.Su6g6hRoi9X("missing ]")
		}
		hCc7zqG4 = hCc7zqG4[1 :  /*line :1*/ len(hCc7zqG4)-1]
		jpLaz9VtZ = true
	}

	return hCc7zqG4, i7pg41eiHlWD, jpLaz9VtZ, nil
}





func J3rl10h2sl4(asil3aeiMbi3 string) (YTqTf_4VC, error) {
	var hwVpbkGq YTqTf_4VC
	hCc7zqG4, i7pg41eiHlWD, jpLaz9VtZ, lcob4O :=  /*line :1*/ jxDvHzBX(asil3aeiMbi3)
	if lcob4O != nil {
		return hwVpbkGq, lcob4O
	}
	usQG03pWNQrx, lcob4O :=  /*line :1*/ strconv.WOxcnaOAzTeq(i7pg41eiHlWD, 10, 16)
	if lcob4O != nil {
		return hwVpbkGq,  /*line :1*/ errors.Su6g6hRoi9X("invalid port " +  /*line :1*/ strconv.D35LD5nn(i7pg41eiHlWD) + " parsing " +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3))
	}
	hwVpbkGq.uqZ4iEGn0_u =  /*line :1*/ uint16(usQG03pWNQrx)
	hwVpbkGq.rjl3GFNJ2JN2, lcob4O =  /*line :1*/ Y1uXU4Sx(hCc7zqG4)
	if lcob4O != nil {
		return YTqTf_4VC{}, lcob4O
	}
	if jpLaz9VtZ &&  /*line :1*/ hwVpbkGq.rjl3GFNJ2JN2.Is4() {
		return YTqTf_4VC{},  /*line :1*/ errors.Su6g6hRoi9X("invalid ip:port " +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3) + ", square brackets can only be used with IPv6 addresses")
	} else if !jpLaz9VtZ &&  /*line :1*/ hwVpbkGq.rjl3GFNJ2JN2.Is6() {
		return YTqTf_4VC{},  /*line :1*/ errors.Su6g6hRoi9X("invalid ip:port " +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3) + ", IPv6 addresses must be surrounded by square brackets")
	}
	return hwVpbkGq, nil
}



func HOSAAYH(asil3aeiMbi3 string) YTqTf_4VC {
	hCc7zqG4, lcob4O :=  /*line :1*/ J3rl10h2sl4(asil3aeiMbi3)
	if lcob4O != nil {
		 /*line :1*/ panic(lcob4O)
	}
	return hCc7zqG4
}



func (ceqvLqr YTqTf_4VC) IsValid() bool	{ return  /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.IsValid() }

func (ceqvLqr YTqTf_4VC) String() string {
	switch ceqvLqr.rjl3GFNJ2JN2.nqiHqD {
	case sl68hrb:
		return "invalid AddrPort"
	case iU_Gva:
		u4b3M4h :=  /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.As4()
		qhMME8jigt :=  /*line :1*/ make([]byte, 0, 21)
		for jVDSPyHrNzR := range u4b3M4h {
			qhMME8jigt =  /*line :1*/ strconv.DxG2yRkVZO(qhMME8jigt,  /*line :1*/ uint64(u4b3M4h[jVDSPyHrNzR]), 10)
			qhMME8jigt =  /*line :1*/ append(qhMME8jigt, "...:"[jVDSPyHrNzR])
		}
		qhMME8jigt =  /*line :1*/ strconv.DxG2yRkVZO(qhMME8jigt,  /*line :1*/ uint64(ceqvLqr.uqZ4iEGn0_u), 10)
		return  /*line :1*/ string(qhMME8jigt)
	default:

		return  /*line :1*/ e3M0wVn_( /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.String(),  /*line :1*/ itoa.V25ba9G5( /*line :1*/ int(ceqvLqr.uqZ4iEGn0_u)))
	}
}

func e3M0wVn_(gVWsZj, i7pg41eiHlWD string) string {

	if  /*line :1*/ bytealg.IndexByteString(gVWsZj, ':') >= 0 {
		return "[" + gVWsZj + "]:" + i7pg41eiHlWD
	}
	return gVWsZj + ":" + i7pg41eiHlWD
}




func (ceqvLqr YTqTf_4VC) AppendTo(eVndpt []byte) []byte {
	switch ceqvLqr.rjl3GFNJ2JN2.nqiHqD {
	case sl68hrb:
		return eVndpt
	case iU_Gva:
		eVndpt =  /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.fythzeXU8(eVndpt)
	default:
		if  /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.Is4In6() {
			eVndpt =  /*line :1*/ append(eVndpt, "[::ffff:"...)
			eVndpt =  /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.Unmap().fythzeXU8(eVndpt)
			if iwlojlVcD :=  /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.Zone(); iwlojlVcD != "" {
				eVndpt =  /*line :1*/ append(eVndpt, '%')
				eVndpt =  /*line :1*/ append(eVndpt, iwlojlVcD...)
			}
		} else {
			eVndpt =  /*line :1*/ append(eVndpt, '[')
			eVndpt =  /*line :1*/ ceqvLqr.rjl3GFNJ2JN2.d2LQe2i(eVndpt)
		}
		eVndpt =  /*line :1*/ append(eVndpt, ']')
	}
	eVndpt =  /*line :1*/ append(eVndpt, ':')
	eVndpt =  /*line :1*/ strconv.DxG2yRkVZO(eVndpt,  /*line :1*/ uint64(ceqvLqr.uqZ4iEGn0_u), 10)
	return eVndpt
}




func (ceqvLqr YTqTf_4VC) MarshalText() ([]byte, error) {
	var fnexDYLqfa1q int
	switch ceqvLqr.rjl3GFNJ2JN2.nqiHqD {
	case sl68hrb:
	case iU_Gva:
		fnexDYLqfa1q =  /*line :1*/ len("255.255.255.255:65535")
	default:
		fnexDYLqfa1q =  /*line :1*/ len("[ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0]:65535")
	}
	eVndpt :=  /*line :1*/ make([]byte, 0, fnexDYLqfa1q)
	eVndpt =  /*line :1*/ ceqvLqr.AppendTo(eVndpt)
	return eVndpt, nil
}




func (ceqvLqr *YTqTf_4VC) UnmarshalText(tefqztg []byte) error {
	if  /*line :1*/ len(tefqztg) == 0 {
		*ceqvLqr = YTqTf_4VC{}
		return nil
	}
	var lcob4O error
	*ceqvLqr, lcob4O =  /*line :1*/ J3rl10h2sl4( /*line :1*/ string(tefqztg))
	return lcob4O
}




func (ceqvLqr YTqTf_4VC) MarshalBinary() ([]byte, error) {
	eVndpt :=  /*line :1*/ ceqvLqr.Addr().eEfDsFeFui01(2)
	 /*line :1*/ xoT1B8Cqz(eVndpt[ /*line :1*/ len(eVndpt)-2:],  /*line :1*/ ceqvLqr.Port())
	return eVndpt, nil
}



func (ceqvLqr *YTqTf_4VC) UnmarshalBinary(eVndpt []byte) error {
	if  /*line :1*/ len(eVndpt) < 2 {
		return  /*line :1*/ errors.Su6g6hRoi9X("unexpected slice size")
	}
	var gIbHM2XzK KFLQ1_1E
	lcob4O :=  /*line :1*/ gIbHM2XzK.UnmarshalBinary(eVndpt[: /*line :1*/ len(eVndpt)-2])
	if lcob4O != nil {
		return lcob4O
	}
	*ceqvLqr =  /*line :1*/ Iqr8_1pEDmw0(gIbHM2XzK,  /*line :1*/ gMnaxp(eVndpt[ /*line :1*/ len(eVndpt)-2:]))
	return nil
}





type We_DLeUDXHez struct {
	egZ9RsBo2dG	KFLQ1_1E

	
	
	cydacuWpGmzG	uint8
}









func QkGHFQ3cd(hCc7zqG4 KFLQ1_1E, izsCjK_ int) We_DLeUDXHez {
	var iIZHsf uint8
	if ! /*line :1*/ hCc7zqG4.rb6oFzFeRc() && izsCjK_ >= 0 && izsCjK_ <=  /*line :1*/ hCc7zqG4.BitLen() {
		iIZHsf =  /*line :1*/ uint8(izsCjK_) + 1
	}
	return We_DLeUDXHez{
		egZ9RsBo2dG:	 /*line :1*/ hCc7zqG4.diJCELqd3r(),
		cydacuWpGmzG:	iIZHsf,
	}
}


func (ceqvLqr We_DLeUDXHez) Addr() KFLQ1_1E	{ return ceqvLqr.egZ9RsBo2dG }




func (ceqvLqr We_DLeUDXHez) Bits() int	{ return  /*line :1*/ int(ceqvLqr.cydacuWpGmzG) - 1 }




func (ceqvLqr We_DLeUDXHez) IsValid() bool	{ return ceqvLqr.cydacuWpGmzG > 0 }

func (ceqvLqr We_DLeUDXHez) rb6oFzFeRc() bool	{ return ceqvLqr == We_DLeUDXHez{} }


func (ceqvLqr We_DLeUDXHez) IsSingleIP() bool {
	return  /*line :1*/ ceqvLqr.IsValid() &&  /*line :1*/ ceqvLqr.Bits() ==  /*line :1*/ ceqvLqr.egZ9RsBo2dG.BitLen()
}








func U6krnn_psn(asil3aeiMbi3 string) (We_DLeUDXHez, error) {
	jVDSPyHrNzR :=  /*line :1*/ ghfJ9faupLy(asil3aeiMbi3, '/')
	if jVDSPyHrNzR < 0 {
		return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("netip.ParsePrefix(" +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3) + "): no '/'")
	}
	hCc7zqG4, lcob4O :=  /*line :1*/ Y1uXU4Sx(asil3aeiMbi3[:jVDSPyHrNzR])
	if lcob4O != nil {
		return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("netip.ParsePrefix(" +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3) + "): " +  /*line :1*/ lcob4O.Error())
	}

	if  /*line :1*/ hCc7zqG4.Is6() && hCc7zqG4.nqiHqD != sq5HIi {
		return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("netip.ParsePrefix(" +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3) + "): IPv6 zones cannot be present in a prefix")
	}

	fjTF9426 := asil3aeiMbi3[jVDSPyHrNzR+1:]
	izsCjK_, lcob4O :=  /*line :1*/ strconv.GmbOy60GCccC(fjTF9426)
	if lcob4O != nil {
		return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("netip.ParsePrefix(" +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3) + "): bad bits after slash: " +  /*line :1*/ strconv.D35LD5nn(fjTF9426))
	}
	j60E1d3jHYA := 32
	if  /*line :1*/ hCc7zqG4.Is6() {
		j60E1d3jHYA = 128
	}
	if izsCjK_ < 0 || izsCjK_ > j60E1d3jHYA {
		return We_DLeUDXHez{},  /*line :1*/ errors.Su6g6hRoi9X("netip.ParsePrefix(" +  /*line :1*/ strconv.D35LD5nn(asil3aeiMbi3) + "): prefix length out of range")
	}
	return  /*line :1*/ QkGHFQ3cd(hCc7zqG4, izsCjK_), nil
}



func FGpBnHmNk(asil3aeiMbi3 string) We_DLeUDXHez {
	hCc7zqG4, lcob4O :=  /*line :1*/ U6krnn_psn(asil3aeiMbi3)
	if lcob4O != nil {
		 /*line :1*/ panic(lcob4O)
	}
	return hCc7zqG4
}





func (ceqvLqr We_DLeUDXHez) Masked() We_DLeUDXHez {
	imP6SgWD, _ :=  /*line :1*/ ceqvLqr.egZ9RsBo2dG.Prefix( /*line :1*/ ceqvLqr.Bits())
	return imP6SgWD
}








func (ceqvLqr We_DLeUDXHez) Contains(hCc7zqG4 KFLQ1_1E) bool {
	if ! /*line :1*/ ceqvLqr.IsValid() ||  /*line :1*/ hCc7zqG4._oaAHhqH6AM() {
		return false
	}
	if d5w3i5po8vTB, p9txU5i0Y :=  /*line :1*/ ceqvLqr.egZ9RsBo2dG.BitLen(),  /*line :1*/ hCc7zqG4.BitLen(); d5w3i5po8vTB == 0 || p9txU5i0Y == 0 || d5w3i5po8vTB != p9txU5i0Y {
		return false
	}
	if  /*line :1*/ hCc7zqG4.Is4() {

		return  /*line :1*/ uint32((hCc7zqG4.rPvius.gRGNmk^ceqvLqr.egZ9RsBo2dG.rPvius.gRGNmk)>>((32- /*line :1*/ ceqvLqr.Bits())&63)) == 0
	} else {

		return  /*line :1*/ hCc7zqG4.rPvius.sf7IVRl(ceqvLqr.egZ9RsBo2dG.rPvius).huLWxs63VG1( /*line :1*/ vtgBPgPqR37( /*line :1*/ ceqvLqr.Bits())).rb6oFzFeRc()
	}
}






func (ceqvLqr We_DLeUDXHez) Overlaps(pjBTGxt We_DLeUDXHez) bool {
	if ! /*line :1*/ ceqvLqr.IsValid() || ! /*line :1*/ pjBTGxt.IsValid() {
		return false
	}
	if ceqvLqr == pjBTGxt {
		return true
	}
	if  /*line :1*/ ceqvLqr.egZ9RsBo2dG.Is4() !=  /*line :1*/ pjBTGxt.egZ9RsBo2dG.Is4() {
		return false
	}
	var aKUYBTprP3 int
	if fA_DGfEna6, gNYrgb :=  /*line :1*/ ceqvLqr.Bits(),  /*line :1*/ pjBTGxt.Bits(); fA_DGfEna6 < gNYrgb {
		aKUYBTprP3 = fA_DGfEna6
	} else {
		aKUYBTprP3 = gNYrgb
	}
	if aKUYBTprP3 == 0 {
		return true
	}
	
	
	
	
	var lcob4O error
	if ceqvLqr, lcob4O =  /*line :1*/ ceqvLqr.egZ9RsBo2dG.Prefix(aKUYBTprP3); lcob4O != nil {
		return false
	}
	if pjBTGxt, lcob4O =  /*line :1*/ pjBTGxt.egZ9RsBo2dG.Prefix(aKUYBTprP3); lcob4O != nil {
		return false
	}
	return ceqvLqr.egZ9RsBo2dG == pjBTGxt.egZ9RsBo2dG
}




func (ceqvLqr We_DLeUDXHez) AppendTo(eVndpt []byte) []byte {
	if  /*line :1*/ ceqvLqr.rb6oFzFeRc() {
		return eVndpt
	}
	if ! /*line :1*/ ceqvLqr.IsValid() {
		return  /*line :1*/ append(eVndpt, "invalid Prefix"...)
	}

	if ceqvLqr.egZ9RsBo2dG.nqiHqD == iU_Gva {
		eVndpt =  /*line :1*/ ceqvLqr.egZ9RsBo2dG.fythzeXU8(eVndpt)
	} else {
		if  /*line :1*/ ceqvLqr.egZ9RsBo2dG.Is4In6() {
			eVndpt =  /*line :1*/ append(eVndpt, "::ffff:"...)
			eVndpt =  /*line :1*/ ceqvLqr.egZ9RsBo2dG.Unmap().fythzeXU8(eVndpt)
		} else {
			eVndpt =  /*line :1*/ ceqvLqr.egZ9RsBo2dG.d2LQe2i(eVndpt)
		}
	}

	eVndpt =  /*line :1*/ append(eVndpt, '/')
	eVndpt =  /*line :1*/ rM7DocatE(eVndpt,  /*line :1*/ uint8( /*line :1*/ ceqvLqr.Bits()))
	return eVndpt
}




func (ceqvLqr We_DLeUDXHez) MarshalText() ([]byte, error) {
	var fnexDYLqfa1q int
	switch ceqvLqr.egZ9RsBo2dG.nqiHqD {
	case sl68hrb:
	case iU_Gva:
		fnexDYLqfa1q =  /*line :1*/ len("255.255.255.255/32")
	default:
		fnexDYLqfa1q =  /*line :1*/ len("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff%enp5s0/128")
	}
	eVndpt :=  /*line :1*/ make([]byte, 0, fnexDYLqfa1q)
	eVndpt =  /*line :1*/ ceqvLqr.AppendTo(eVndpt)
	return eVndpt, nil
}




func (ceqvLqr *We_DLeUDXHez) UnmarshalText(tefqztg []byte) error {
	if  /*line :1*/ len(tefqztg) == 0 {
		*ceqvLqr = We_DLeUDXHez{}
		return nil
	}
	var lcob4O error
	*ceqvLqr, lcob4O =  /*line :1*/ U6krnn_psn( /*line :1*/ string(tefqztg))
	return lcob4O
}




func (ceqvLqr We_DLeUDXHez) MarshalBinary() ([]byte, error) {
	eVndpt :=  /*line :1*/ ceqvLqr.Addr().diJCELqd3r().eEfDsFeFui01(1)
	eVndpt[ /*line :1*/ len(eVndpt)-1] =  /*line :1*/ uint8( /*line :1*/ ceqvLqr.Bits())
	return eVndpt, nil
}



func (ceqvLqr *We_DLeUDXHez) UnmarshalBinary(eVndpt []byte) error {
	if  /*line :1*/ len(eVndpt) < 1 {
		return  /*line :1*/ errors.Su6g6hRoi9X("unexpected slice size")
	}
	var gIbHM2XzK KFLQ1_1E
	lcob4O :=  /*line :1*/ gIbHM2XzK.UnmarshalBinary(eVndpt[: /*line :1*/ len(eVndpt)-1])
	if lcob4O != nil {
		return lcob4O
	}
	*ceqvLqr =  /*line :1*/ QkGHFQ3cd(gIbHM2XzK,  /*line :1*/ int(eVndpt[ /*line :1*/ len(eVndpt)-1]))
	return nil
}


func (ceqvLqr We_DLeUDXHez) String() string {
	if ! /*line :1*/ ceqvLqr.IsValid() {
		return "invalid Prefix"
	}
	return  /*line :1*/ ceqvLqr.egZ9RsBo2dG.String() + "/" +  /*line :1*/ itoa.V25ba9G5( /*line :1*/ ceqvLqr.Bits())
}
