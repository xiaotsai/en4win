//line :1
package fuA83L

import fmt "kFsoCfy5zWG"

const (
	_VT_name_0	= "VT_EMPTYVT_NULLVT_I2VT_I4VT_R4VT_R8VT_CYVT_DATEVT_BSTRVT_DISPATCHVT_ERRORVT_BOOLVT_VARIANTVT_UNKNOWNVT_DECIMAL"
	_VT_name_1	= "VT_I1VT_UI1VT_UI2VT_UI4VT_I8VT_UI8VT_INTVT_UINTVT_VOIDVT_HRESULTVT_PTRVT_SAFEARRAYVT_CARRAYVT_USERDEFINEDVT_LPSTRVT_LPWSTR"
	_VT_name_2	= "VT_RECORDVT_INT_PTRVT_UINT_PTR"
	_VT_name_3	= "VT_FILETIMEVT_BLOBVT_STREAMVT_STORAGEVT_STREAMED_OBJECTVT_STORED_OBJECTVT_BLOB_OBJECTVT_CFVT_CLSID"
	_VT_name_4	= "VT_BSTR_BLOBVT_VECTOR"
	_VT_name_5	= "VT_ARRAY"
	_VT_name_6	= "VT_BYREF"
	_VT_name_7	= "VT_RESERVED"
	_VT_name_8	= "VT_ILLEGAL"
)

var (
	bhVrwgOB	= [...]uint8{0, 8, 15, 20, 25, 30, 35, 40, 47, 54, 65, 73, 80, 90, 100, 110}
	yyoD58	= [...]uint8{0, 5, 11, 17, 23, 28, 34, 40, 47, 54, 64, 70, 82, 91, 105, 113, 122}
	hI97mWz	= [...]uint8{0, 9, 19, 30}
	iAI_T7b	= [...]uint8{0, 11, 18, 27, 37, 55, 71, 85, 90, 98}
	an3K0F0YR	= [...]uint8{0, 12, 21}
	hKPWRRaf12x	= [...]uint8{0, 8}
	a0tzM98	= [...]uint8{0, 8}
	bpo5ozL	= [...]uint8{0, 11}
	dsBi7j	= [...]uint8{0, 10}
)

func (gXX2nE5 JtqZSr3B) String() string {
	switch {
	case 0 <= gXX2nE5 && gXX2nE5 <= 14:
		return _VT_name_0[bhVrwgOB[gXX2nE5]:bhVrwgOB[gXX2nE5+1]]
	case 16 <= gXX2nE5 && gXX2nE5 <= 31:
		gXX2nE5 -= 16
		return _VT_name_1[yyoD58[gXX2nE5]:yyoD58[gXX2nE5+1]]
	case 36 <= gXX2nE5 && gXX2nE5 <= 38:
		gXX2nE5 -= 36
		return _VT_name_2[hI97mWz[gXX2nE5]:hI97mWz[gXX2nE5+1]]
	case 64 <= gXX2nE5 && gXX2nE5 <= 72:
		gXX2nE5 -= 64
		return _VT_name_3[iAI_T7b[gXX2nE5]:iAI_T7b[gXX2nE5+1]]
	case 4095 <= gXX2nE5 && gXX2nE5 <= 4096:
		gXX2nE5 -= 4095
		return _VT_name_4[an3K0F0YR[gXX2nE5]:an3K0F0YR[gXX2nE5+1]]
	case gXX2nE5 == 8192:
		return _VT_name_5
	case gXX2nE5 == 16384:
		return _VT_name_6
	case gXX2nE5 == 32768:
		return _VT_name_7
	case gXX2nE5 == 65535:
		return _VT_name_8
	default:
		return  /*line :1*/ fmt.IBsS3Oc("VT(%d)", gXX2nE5)
	}
}
