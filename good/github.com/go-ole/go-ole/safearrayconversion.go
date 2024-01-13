//line :1
package fuA83L

import (
	"unsafe"
)

type JA6nr0c8Wi struct {
	IQUe2SRXObIO *V8vhLAHnL4
}

func (hw9Gp4V *JA6nr0c8Wi) ToStringArray() (rH0sRai []string) {
	yVK8aOqWz, _ :=  /*line :1*/ hw9Gp4V.TotalElements(0)
	rH0sRai =  /*line :1*/ make([]string, yVK8aOqWz)

	for gXX2nE5 :=  /*line :1*/ int32(0); gXX2nE5 < yVK8aOqWz; gXX2nE5++ {
		rH0sRai[ /*line :1*/ int32(gXX2nE5)], _ =  /*line :1*/ aSH1U3M8tC(hw9Gp4V.IQUe2SRXObIO, gXX2nE5)
	}

	return
}

func (hw9Gp4V *JA6nr0c8Wi) ToByteArray() (u0hn5w []byte) {
	yVK8aOqWz, _ :=  /*line :1*/ hw9Gp4V.TotalElements(0)
	u0hn5w =  /*line :1*/ make([]byte, yVK8aOqWz)

	for gXX2nE5 :=  /*line :1*/ int32(0); gXX2nE5 < yVK8aOqWz; gXX2nE5++ {
		 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&u0hn5w[ /*line :1*/ int32(gXX2nE5)]))
	}

	return
}

func (hw9Gp4V *JA6nr0c8Wi) ToValueArray() (fa9ufcTou []interface{}) {
	yVK8aOqWz, _ :=  /*line :1*/ hw9Gp4V.TotalElements(0)
	fa9ufcTou =  /*line :1*/ make([]interface{}, yVK8aOqWz)
	tCRLo1W5m9va, _ :=  /*line :1*/ a3r7Dae7iEQ(hw9Gp4V.IQUe2SRXObIO)

	for gXX2nE5 :=  /*line :1*/ int32(0); gXX2nE5 < yVK8aOqWz; gXX2nE5++ {
		switch  /*line :1*/ JtqZSr3B(tCRLo1W5m9va) {
		case VT_BOOL:
			var y0jiLdYHXNnx bool
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_I1:
			var y0jiLdYHXNnx int8
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_I2:
			var y0jiLdYHXNnx int16
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_I4:
			var y0jiLdYHXNnx int32
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_I8:
			var y0jiLdYHXNnx int64
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_UI1:
			var y0jiLdYHXNnx uint8
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_UI2:
			var y0jiLdYHXNnx uint16
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_UI4:
			var y0jiLdYHXNnx uint32
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_UI8:
			var y0jiLdYHXNnx uint64
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_R4:
			var y0jiLdYHXNnx float32
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_R8:
			var y0jiLdYHXNnx float64
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_BSTR:
			var y0jiLdYHXNnx string
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] = y0jiLdYHXNnx
		case VT_VARIANT:
			var y0jiLdYHXNnx KEVetAOpxl0D
			 /*line :1*/ y_RaDlZ8oKX(hw9Gp4V.IQUe2SRXObIO, gXX2nE5,  /*line :1*/ unsafe.Pointer(&y0jiLdYHXNnx))
			fa9ufcTou[gXX2nE5] =  /*line :1*/ y0jiLdYHXNnx.Value()
		default:

		}
	}

	return
}

func (hw9Gp4V *JA6nr0c8Wi) GetType() (ncar6HvbRqig uint16, zsgC7M1 error) {
	return  /*line :1*/ a3r7Dae7iEQ(hw9Gp4V.IQUe2SRXObIO)
}

func (hw9Gp4V *JA6nr0c8Wi) GetDimensions() (llaoKMWl *uint32, zsgC7M1 error) {
	return  /*line :1*/ gLOSo9HfR(hw9Gp4V.IQUe2SRXObIO)
}

func (hw9Gp4V *JA6nr0c8Wi) GetSize() (nXciaovStia3 *uint32, zsgC7M1 error) {
	return  /*line :1*/ cHYfChmeKf(hw9Gp4V.IQUe2SRXObIO)
}

func (hw9Gp4V *JA6nr0c8Wi) TotalElements(gpD3CvFM uint32) (yVK8aOqWz int32, zsgC7M1 error) {
	if gpD3CvFM < 1 {
		gpD3CvFM = 1
	}

	
	var KIiZ_W4ZhL int32
	var VmraLg3rsAwB int32

	KIiZ_W4ZhL, zsgC7M1 =  /*line :1*/ xB9HfmbEgA45(hw9Gp4V.IQUe2SRXObIO, gpD3CvFM)
	if zsgC7M1 != nil {
		return
	}

	VmraLg3rsAwB, zsgC7M1 =  /*line :1*/ iSPh3lx8EC(hw9Gp4V.IQUe2SRXObIO, gpD3CvFM)
	if zsgC7M1 != nil {
		return
	}

	yVK8aOqWz = VmraLg3rsAwB - KIiZ_W4ZhL + 1
	return
}


func (hw9Gp4V *JA6nr0c8Wi) Release() {
	 /*line :1*/ ditmyzm7rSN(hw9Gp4V.IQUe2SRXObIO)
}
