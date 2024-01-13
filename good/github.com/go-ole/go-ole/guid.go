//line :1
package fuA83L

var (
	
	YH_aGx	=  /*line :1*/ E7kY1q("{00000000-0000-0000-0000-000000000000}")

	
	WySQlWR	=  /*line :1*/ E7kY1q("{00000000-0000-0000-C000-000000000046}")

	
	Y0sP6F	=  /*line :1*/ E7kY1q("{00020400-0000-0000-C000-000000000046}")

	
	Uv2Iuh4vReN	=  /*line :1*/ E7kY1q("{00020404-0000-0000-C000-000000000046}")

	
	D3Vh84Kk	=  /*line :1*/ E7kY1q("{B196B284-BAB4-101A-B69C-00AA00341D07}")

	
	Hb6PWo	=  /*line :1*/ E7kY1q("{B196B286-BAB4-101A-B69C-00AA00341D07}")

	
	CzsXF9cMBur	=  /*line :1*/ E7kY1q("{AF86E2E0-B12D-4C6A-9C5A-D7AA65101E90}")

	
	ZFdbSZh	=  /*line :1*/ E7kY1q("{B196B283-BAB4-101A-B69C-00AA00341D07}")
)


var (
	
	
	
	XA5pgSY	=  /*line :1*/ E7kY1q("{E0133EB4-C36F-469A-9D3D-C66B84BE19ED}")

	
	
	
	O6bkLL1Efzx	=  /*line :1*/ E7kY1q("{BEB06610-EB84-4155-AF58-E2BFF53680B4}")

	
	
	
	V6kKhOm1Ii_	=  /*line :1*/ E7kY1q("{DAA3F9FA-761E-4976-A860-8364CE55F6FC}")

	
	
	
	TRgd3j	=  /*line :1*/ E7kY1q("{E3DEDEE7-38A2-4540-91D1-2EEF1D8891B0}")

	
	
	
	PHiKMR4	=  /*line :1*/ E7kY1q("{8D437CBC-B3ED-485C-BC32-C336432A1623}")

	
	
	
	BCMeHnHve	=  /*line :1*/ E7kY1q("{BF1ED004-EA02-456A-AA55-2AC8AC6B054C}")

	
	
	
	TKXkzwSP	=  /*line :1*/ E7kY1q("{BF908A81-8687-4E93-999F-D86FAB284BA0}")

	
	
	
	UjJcAnLz	=  /*line :1*/ E7kY1q("{D530E7A6-4EE8-40D1-8931-3D63B8605010}")

	
	
	
	HoKXr9Jo	=  /*line :1*/ E7kY1q("{6485B1EF-D780-4834-A4FE-1EBB51746CA3}")

	
	
	
	BbHZKl	=  /*line :1*/ E7kY1q("{CCA8D7AE-91C0-4277-A8B3-FF4EDF28D3C0}")

	
	
	
	LCg4RQ	=  /*line :1*/ E7kY1q("{3C24506A-AE9E-4D50-9157-EF317281F1B0}")

	
	
	
	Gxghur	=  /*line :1*/ E7kY1q("{865B85C5-0334-4AC6-9EF6-AACEC8FC5E86}")
)

const hextable = "0123456789ABCDEF"
const emptyGUID = "{00000000-0000-0000-0000-000000000000}"





type GUID struct {
	Data1	uint32
	Data2	uint16
	Data3	uint16
	Data4	[8]byte
}











func E7kY1q(i15rO5WD3Nx string) *GUID {
	b9Flo0 := [] /*line :1*/ byte(i15rO5WD3Nx)
	var n5hzV5wi, qLg4j85be28, viYuyT, a9aOEXPgXp7, heD6bCNASO []byte

	switch  /*line :1*/ len(b9Flo0) {
	case 38:
		if b9Flo0[0] != '{' || b9Flo0[37] != '}' {
			return nil
		}
		b9Flo0 = b9Flo0[1:37]
		fallthrough
	case 36:
		if b9Flo0[8] != '-' || b9Flo0[13] != '-' || b9Flo0[18] != '-' || b9Flo0[23] != '-' {
			return nil
		}
		n5hzV5wi = b9Flo0[0:8]
		qLg4j85be28 = b9Flo0[9:13]
		viYuyT = b9Flo0[14:18]
		a9aOEXPgXp7 = b9Flo0[19:23]
		heD6bCNASO = b9Flo0[24:36]
	case 32:
		n5hzV5wi = b9Flo0[0:8]
		qLg4j85be28 = b9Flo0[8:12]
		viYuyT = b9Flo0[12:16]
		a9aOEXPgXp7 = b9Flo0[16:20]
		heD6bCNASO = b9Flo0[20:32]
	default:
		return nil
	}

	var i92XO4UuNULA GUID
	var p5uUaTusNUv8, fQxrba, vVKJC5S8U, njJoqv_oxv bool
	i92XO4UuNULA.Data1, p5uUaTusNUv8 =  /*line :1*/ gBpDfnJq(n5hzV5wi)
	i92XO4UuNULA.Data2, fQxrba =  /*line :1*/ alSf37Vc(qLg4j85be28)
	i92XO4UuNULA.Data3, vVKJC5S8U =  /*line :1*/ alSf37Vc(viYuyT)
	i92XO4UuNULA.Data4, njJoqv_oxv =  /*line :1*/ pJ6EVhuSoA(a9aOEXPgXp7, heD6bCNASO)
	if p5uUaTusNUv8 && fQxrba && vVKJC5S8U && njJoqv_oxv {
		return &i92XO4UuNULA
	}
	return nil
}

func gBpDfnJq(gA9YADk []byte) (rORmQfUB uint32, wXNeRPAUqcq bool) {
	var eIVKjIxC, kuUsx8, dRaubQTOlisO, e_vO__j byte
	var p5uUaTusNUv8, fQxrba, vVKJC5S8U, njJoqv_oxv bool
	eIVKjIxC, p5uUaTusNUv8 =  /*line :1*/ cEmszAecEg(gA9YADk[0], gA9YADk[1])
	kuUsx8, fQxrba =  /*line :1*/ cEmszAecEg(gA9YADk[2], gA9YADk[3])
	dRaubQTOlisO, vVKJC5S8U =  /*line :1*/ cEmszAecEg(gA9YADk[4], gA9YADk[5])
	e_vO__j, njJoqv_oxv =  /*line :1*/ cEmszAecEg(gA9YADk[6], gA9YADk[7])
	rORmQfUB = ( /*line :1*/ uint32(eIVKjIxC) << 24) | ( /*line :1*/ uint32(kuUsx8) << 16) | ( /*line :1*/ uint32(dRaubQTOlisO) << 8) |  /*line :1*/ uint32(e_vO__j)
	wXNeRPAUqcq = p5uUaTusNUv8 && fQxrba && vVKJC5S8U && njJoqv_oxv
	return
}

func alSf37Vc(gA9YADk []byte) (rORmQfUB uint16, wXNeRPAUqcq bool) {
	var eIVKjIxC, kuUsx8 byte
	var p5uUaTusNUv8, fQxrba bool
	eIVKjIxC, p5uUaTusNUv8 =  /*line :1*/ cEmszAecEg(gA9YADk[0], gA9YADk[1])
	kuUsx8, fQxrba =  /*line :1*/ cEmszAecEg(gA9YADk[2], gA9YADk[3])
	rORmQfUB = ( /*line :1*/ uint16(eIVKjIxC) << 8) |  /*line :1*/ uint16(kuUsx8)
	wXNeRPAUqcq = p5uUaTusNUv8 && fQxrba
	return
}

func pJ6EVhuSoA(y02r2wi0M []byte, pE_8Yu01Pfhq []byte) (rORmQfUB [8]byte, wXNeRPAUqcq bool) {
	var p5uUaTusNUv8, fQxrba, vVKJC5S8U, njJoqv_oxv, nS3u9B1pEBmW, qn6GJm2Kg7, ylki7h, kkh3RVbOfhI bool
	rORmQfUB[0], p5uUaTusNUv8 =  /*line :1*/ cEmszAecEg(y02r2wi0M[0], y02r2wi0M[1])
	rORmQfUB[1], fQxrba =  /*line :1*/ cEmszAecEg(y02r2wi0M[2], y02r2wi0M[3])
	rORmQfUB[2], vVKJC5S8U =  /*line :1*/ cEmszAecEg(pE_8Yu01Pfhq[0], pE_8Yu01Pfhq[1])
	rORmQfUB[3], njJoqv_oxv =  /*line :1*/ cEmszAecEg(pE_8Yu01Pfhq[2], pE_8Yu01Pfhq[3])
	rORmQfUB[4], nS3u9B1pEBmW =  /*line :1*/ cEmszAecEg(pE_8Yu01Pfhq[4], pE_8Yu01Pfhq[5])
	rORmQfUB[5], qn6GJm2Kg7 =  /*line :1*/ cEmszAecEg(pE_8Yu01Pfhq[6], pE_8Yu01Pfhq[7])
	rORmQfUB[6], ylki7h =  /*line :1*/ cEmszAecEg(pE_8Yu01Pfhq[8], pE_8Yu01Pfhq[9])
	rORmQfUB[7], kkh3RVbOfhI =  /*line :1*/ cEmszAecEg(pE_8Yu01Pfhq[10], pE_8Yu01Pfhq[11])
	wXNeRPAUqcq = p5uUaTusNUv8 && fQxrba && vVKJC5S8U && njJoqv_oxv && nS3u9B1pEBmW && qn6GJm2Kg7 && ylki7h && kkh3RVbOfhI
	return
}

func cEmszAecEg(hPNyAapNH1, ox3Noc8At byte) (rORmQfUB byte, wXNeRPAUqcq bool) {
	var fKf9gUq, mH33NO byte
	var p5uUaTusNUv8, fQxrba bool
	fKf9gUq, p5uUaTusNUv8 =  /*line :1*/ h8CP_sgA(hPNyAapNH1)
	mH33NO, fQxrba =  /*line :1*/ h8CP_sgA(ox3Noc8At)
	rORmQfUB = (fKf9gUq << 4) | mH33NO
	wXNeRPAUqcq = p5uUaTusNUv8 && fQxrba
	return
}

func h8CP_sgA(eT1_dhKA byte) (byte, bool) {
	switch {
	case '0' <= eT1_dhKA && eT1_dhKA <= '9':
		return eT1_dhKA - '0', true
	case 'a' <= eT1_dhKA && eT1_dhKA <= 'f':
		return eT1_dhKA - 'a' + 10, true
	case 'A' <= eT1_dhKA && eT1_dhKA <= 'F':
		return eT1_dhKA - 'A' + 10, true
	}

	return 0, false
}








func (i15rO5WD3Nx *GUID) String() string {
	if i15rO5WD3Nx == nil {
		return emptyGUID
	}

	var eT1_dhKA [38]byte
	eT1_dhKA[0] = '{'
	 /*line :1*/ gMqlVtvyaYu5(eT1_dhKA[1:9], i15rO5WD3Nx.Data1)
	eT1_dhKA[9] = '-'
	 /*line :1*/ al9shJjo(eT1_dhKA[10:14], i15rO5WD3Nx.Data2)
	eT1_dhKA[14] = '-'
	 /*line :1*/ al9shJjo(eT1_dhKA[15:19], i15rO5WD3Nx.Data3)
	eT1_dhKA[19] = '-'
	 /*line :1*/ zaptMGs6EdHD(eT1_dhKA[20:24], i15rO5WD3Nx.Data4[0:2])
	eT1_dhKA[24] = '-'
	 /*line :1*/ zaptMGs6EdHD(eT1_dhKA[25:37], i15rO5WD3Nx.Data4[2:8])
	eT1_dhKA[37] = '}'
	return  /*line :1*/ string(eT1_dhKA[:])
}

func gMqlVtvyaYu5(wRIFgDKVJXb []byte, y0jiLdYHXNnx uint32) {
	wRIFgDKVJXb[0] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>24)>>4]
	wRIFgDKVJXb[1] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>24)&0x0f]
	wRIFgDKVJXb[2] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>16)>>4]
	wRIFgDKVJXb[3] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>16)&0x0f]
	wRIFgDKVJXb[4] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>8)>>4]
	wRIFgDKVJXb[5] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>8)&0x0f]
	wRIFgDKVJXb[6] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx)>>4]
	wRIFgDKVJXb[7] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx)&0x0f]
}

func al9shJjo(wRIFgDKVJXb []byte, y0jiLdYHXNnx uint16) {
	wRIFgDKVJXb[0] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>8)>>4]
	wRIFgDKVJXb[1] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx>>8)&0x0f]
	wRIFgDKVJXb[2] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx)>>4]
	wRIFgDKVJXb[3] = hextable[ /*line :1*/ byte(y0jiLdYHXNnx)&0x0f]
}

func zaptMGs6EdHD(wF7ZEe2CNqI, gA9YADk []byte) {
	for gXX2nE5 := 0; gXX2nE5 <  /*line :1*/ len(gA9YADk); gXX2nE5++ {
		wF7ZEe2CNqI[gXX2nE5*2] = hextable[gA9YADk[gXX2nE5]>>4]
		wF7ZEe2CNqI[gXX2nE5*2+1] = hextable[gA9YADk[gXX2nE5]&0x0f]
	}
}




func Apsz2oVG2lN(lSrdwVNGSl7 *GUID, lss73QBmp *GUID) bool {
	return lSrdwVNGSl7.Data1 == lss73QBmp.Data1 &&
		lSrdwVNGSl7.Data2 == lss73QBmp.Data2 &&
		lSrdwVNGSl7.Data3 == lss73QBmp.Data3 &&
		lSrdwVNGSl7.Data4[0] == lss73QBmp.Data4[0] &&
		lSrdwVNGSl7.Data4[1] == lss73QBmp.Data4[1] &&
		lSrdwVNGSl7.Data4[2] == lss73QBmp.Data4[2] &&
		lSrdwVNGSl7.Data4[3] == lss73QBmp.Data4[3] &&
		lSrdwVNGSl7.Data4[4] == lss73QBmp.Data4[4] &&
		lSrdwVNGSl7.Data4[5] == lss73QBmp.Data4[5] &&
		lSrdwVNGSl7.Data4[6] == lss73QBmp.Data4[6] &&
		lSrdwVNGSl7.Data4[7] == lss73QBmp.Data4[7]
}
