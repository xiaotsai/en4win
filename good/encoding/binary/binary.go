//line :1

















package yLm9uBQG

import (
	errors "iAHoxjmM"
	io "xI9ai2djaJ2"
	math "math"
	reflect "reflect"
	sync "sync"
)



type JUEyK4eMhp interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
	String() string
}



type VTsWtRSuPS interface {
	AppendUint16([]byte, uint16) []byte
	AppendUint32([]byte, uint32) []byte
	AppendUint64([]byte, uint64) []byte
	String() string
}


var BE16BND c64oU52pg


var FOcKBq6 i7oyDRP3Ny7s

type c64oU52pg struct{}

func (c64oU52pg) Uint16(nWegNNdJn_ []byte) uint16 {
	_ = nWegNNdJn_[1]
	return  /*line :1*/ uint16(nWegNNdJn_[0]) |  /*line :1*/ uint16(nWegNNdJn_[1])<<8
}

func (c64oU52pg) PutUint16(nWegNNdJn_ []byte, fj8vmzJPe uint16) {
	_ = nWegNNdJn_[1]
	nWegNNdJn_[0] =  /*line :1*/ byte(fj8vmzJPe)
	nWegNNdJn_[1] =  /*line :1*/ byte(fj8vmzJPe >> 8)
}

func (c64oU52pg) AppendUint16(nWegNNdJn_ []byte, fj8vmzJPe uint16) []byte {
	return  /*line :1*/ append(nWegNNdJn_,
		 /*line :1*/ byte(fj8vmzJPe),
		 /*line :1*/ byte(fj8vmzJPe>>8),
	)
}

func (c64oU52pg) Uint32(nWegNNdJn_ []byte) uint32 {
	_ = nWegNNdJn_[3]
	return  /*line :1*/ uint32(nWegNNdJn_[0]) |  /*line :1*/ uint32(nWegNNdJn_[1])<<8 |  /*line :1*/ uint32(nWegNNdJn_[2])<<16 |  /*line :1*/ uint32(nWegNNdJn_[3])<<24
}

func (c64oU52pg) PutUint32(nWegNNdJn_ []byte, fj8vmzJPe uint32) {
	_ = nWegNNdJn_[3]
	nWegNNdJn_[0] =  /*line :1*/ byte(fj8vmzJPe)
	nWegNNdJn_[1] =  /*line :1*/ byte(fj8vmzJPe >> 8)
	nWegNNdJn_[2] =  /*line :1*/ byte(fj8vmzJPe >> 16)
	nWegNNdJn_[3] =  /*line :1*/ byte(fj8vmzJPe >> 24)
}

func (c64oU52pg) AppendUint32(nWegNNdJn_ []byte, fj8vmzJPe uint32) []byte {
	return  /*line :1*/ append(nWegNNdJn_,
		 /*line :1*/ byte(fj8vmzJPe),
		 /*line :1*/ byte(fj8vmzJPe>>8),
		 /*line :1*/ byte(fj8vmzJPe>>16),
		 /*line :1*/ byte(fj8vmzJPe>>24),
	)
}

func (c64oU52pg) Uint64(nWegNNdJn_ []byte) uint64 {
	_ = nWegNNdJn_[7]
	return  /*line :1*/ uint64(nWegNNdJn_[0]) |  /*line :1*/ uint64(nWegNNdJn_[1])<<8 |  /*line :1*/ uint64(nWegNNdJn_[2])<<16 |  /*line :1*/ uint64(nWegNNdJn_[3])<<24 |
		 /*line :1*/ uint64(nWegNNdJn_[4])<<32 |  /*line :1*/ uint64(nWegNNdJn_[5])<<40 |  /*line :1*/ uint64(nWegNNdJn_[6])<<48 |  /*line :1*/ uint64(nWegNNdJn_[7])<<56
}

func (c64oU52pg) PutUint64(nWegNNdJn_ []byte, fj8vmzJPe uint64) {
	_ = nWegNNdJn_[7]
	nWegNNdJn_[0] =  /*line :1*/ byte(fj8vmzJPe)
	nWegNNdJn_[1] =  /*line :1*/ byte(fj8vmzJPe >> 8)
	nWegNNdJn_[2] =  /*line :1*/ byte(fj8vmzJPe >> 16)
	nWegNNdJn_[3] =  /*line :1*/ byte(fj8vmzJPe >> 24)
	nWegNNdJn_[4] =  /*line :1*/ byte(fj8vmzJPe >> 32)
	nWegNNdJn_[5] =  /*line :1*/ byte(fj8vmzJPe >> 40)
	nWegNNdJn_[6] =  /*line :1*/ byte(fj8vmzJPe >> 48)
	nWegNNdJn_[7] =  /*line :1*/ byte(fj8vmzJPe >> 56)
}

func (c64oU52pg) AppendUint64(nWegNNdJn_ []byte, fj8vmzJPe uint64) []byte {
	return  /*line :1*/ append(nWegNNdJn_,
		 /*line :1*/ byte(fj8vmzJPe),
		 /*line :1*/ byte(fj8vmzJPe>>8),
		 /*line :1*/ byte(fj8vmzJPe>>16),
		 /*line :1*/ byte(fj8vmzJPe>>24),
		 /*line :1*/ byte(fj8vmzJPe>>32),
		 /*line :1*/ byte(fj8vmzJPe>>40),
		 /*line :1*/ byte(fj8vmzJPe>>48),
		 /*line :1*/ byte(fj8vmzJPe>>56),
	)
}

func (c64oU52pg) String() string	{ return "LittleEndian" }

func (c64oU52pg) GoString() string	{ return "binary.LittleEndian" }

type i7oyDRP3Ny7s struct{}

func (i7oyDRP3Ny7s) Uint16(nWegNNdJn_ []byte) uint16 {
	_ = nWegNNdJn_[1]
	return  /*line :1*/ uint16(nWegNNdJn_[1]) |  /*line :1*/ uint16(nWegNNdJn_[0])<<8
}

func (i7oyDRP3Ny7s) PutUint16(nWegNNdJn_ []byte, fj8vmzJPe uint16) {
	_ = nWegNNdJn_[1]
	nWegNNdJn_[0] =  /*line :1*/ byte(fj8vmzJPe >> 8)
	nWegNNdJn_[1] =  /*line :1*/ byte(fj8vmzJPe)
}

func (i7oyDRP3Ny7s) AppendUint16(nWegNNdJn_ []byte, fj8vmzJPe uint16) []byte {
	return  /*line :1*/ append(nWegNNdJn_,
		 /*line :1*/ byte(fj8vmzJPe>>8),
		 /*line :1*/ byte(fj8vmzJPe),
	)
}

func (i7oyDRP3Ny7s) Uint32(nWegNNdJn_ []byte) uint32 {
	_ = nWegNNdJn_[3]
	return  /*line :1*/ uint32(nWegNNdJn_[3]) |  /*line :1*/ uint32(nWegNNdJn_[2])<<8 |  /*line :1*/ uint32(nWegNNdJn_[1])<<16 |  /*line :1*/ uint32(nWegNNdJn_[0])<<24
}

func (i7oyDRP3Ny7s) PutUint32(nWegNNdJn_ []byte, fj8vmzJPe uint32) {
	_ = nWegNNdJn_[3]
	nWegNNdJn_[0] =  /*line :1*/ byte(fj8vmzJPe >> 24)
	nWegNNdJn_[1] =  /*line :1*/ byte(fj8vmzJPe >> 16)
	nWegNNdJn_[2] =  /*line :1*/ byte(fj8vmzJPe >> 8)
	nWegNNdJn_[3] =  /*line :1*/ byte(fj8vmzJPe)
}

func (i7oyDRP3Ny7s) AppendUint32(nWegNNdJn_ []byte, fj8vmzJPe uint32) []byte {
	return  /*line :1*/ append(nWegNNdJn_,
		 /*line :1*/ byte(fj8vmzJPe>>24),
		 /*line :1*/ byte(fj8vmzJPe>>16),
		 /*line :1*/ byte(fj8vmzJPe>>8),
		 /*line :1*/ byte(fj8vmzJPe),
	)
}

func (i7oyDRP3Ny7s) Uint64(nWegNNdJn_ []byte) uint64 {
	_ = nWegNNdJn_[7]
	return  /*line :1*/ uint64(nWegNNdJn_[7]) |  /*line :1*/ uint64(nWegNNdJn_[6])<<8 |  /*line :1*/ uint64(nWegNNdJn_[5])<<16 |  /*line :1*/ uint64(nWegNNdJn_[4])<<24 |
		 /*line :1*/ uint64(nWegNNdJn_[3])<<32 |  /*line :1*/ uint64(nWegNNdJn_[2])<<40 |  /*line :1*/ uint64(nWegNNdJn_[1])<<48 |  /*line :1*/ uint64(nWegNNdJn_[0])<<56
}

func (i7oyDRP3Ny7s) PutUint64(nWegNNdJn_ []byte, fj8vmzJPe uint64) {
	_ = nWegNNdJn_[7]
	nWegNNdJn_[0] =  /*line :1*/ byte(fj8vmzJPe >> 56)
	nWegNNdJn_[1] =  /*line :1*/ byte(fj8vmzJPe >> 48)
	nWegNNdJn_[2] =  /*line :1*/ byte(fj8vmzJPe >> 40)
	nWegNNdJn_[3] =  /*line :1*/ byte(fj8vmzJPe >> 32)
	nWegNNdJn_[4] =  /*line :1*/ byte(fj8vmzJPe >> 24)
	nWegNNdJn_[5] =  /*line :1*/ byte(fj8vmzJPe >> 16)
	nWegNNdJn_[6] =  /*line :1*/ byte(fj8vmzJPe >> 8)
	nWegNNdJn_[7] =  /*line :1*/ byte(fj8vmzJPe)
}

func (i7oyDRP3Ny7s) AppendUint64(nWegNNdJn_ []byte, fj8vmzJPe uint64) []byte {
	return  /*line :1*/ append(nWegNNdJn_,
		 /*line :1*/ byte(fj8vmzJPe>>56),
		 /*line :1*/ byte(fj8vmzJPe>>48),
		 /*line :1*/ byte(fj8vmzJPe>>40),
		 /*line :1*/ byte(fj8vmzJPe>>32),
		 /*line :1*/ byte(fj8vmzJPe>>24),
		 /*line :1*/ byte(fj8vmzJPe>>16),
		 /*line :1*/ byte(fj8vmzJPe>>8),
		 /*line :1*/ byte(fj8vmzJPe),
	)
}

func (i7oyDRP3Ny7s) String() string	{ return "BigEndian" }

func (i7oyDRP3Ny7s) GoString() string	{ return "binary.BigEndian" }

func (nhrSeOfanUPm) String() string	{ return "NativeEndian" }

func (nhrSeOfanUPm) GoString() string	{ return "binary.NativeEndian" }

















func OrEIDcJ0jtm2(fswXI7GD2zM io.YJ04Fau, uOWHfVo JUEyK4eMhp, iFZRWdRjb any) error {

	if xnoQ5043QZ :=  /*line :1*/ g1aptfN(iFZRWdRjb); xnoQ5043QZ != 0 {
		eESy98 :=  /*line :1*/ make([]byte, xnoQ5043QZ)
		if _, ne233xQRrIq :=  /*line :1*/ io.G0zBgKg(fswXI7GD2zM, eESy98); ne233xQRrIq != nil {
			return ne233xQRrIq
		}
		switch iFZRWdRjb := iFZRWdRjb.(type) {
		case *bool:
			*iFZRWdRjb = eESy98[0] != 0
		case *int8:
			*iFZRWdRjb =  /*line :1*/ int8(eESy98[0])
		case *uint8:
			*iFZRWdRjb = eESy98[0]
		case *int16:
			*iFZRWdRjb =  /*line :1*/ int16( /*line :1*/ uOWHfVo.Uint16(eESy98))
		case *uint16:
			*iFZRWdRjb =  /*line :1*/ uOWHfVo.Uint16(eESy98)
		case *int32:
			*iFZRWdRjb =  /*line :1*/ int32( /*line :1*/ uOWHfVo.Uint32(eESy98))
		case *uint32:
			*iFZRWdRjb =  /*line :1*/ uOWHfVo.Uint32(eESy98)
		case *int64:
			*iFZRWdRjb =  /*line :1*/ int64( /*line :1*/ uOWHfVo.Uint64(eESy98))
		case *uint64:
			*iFZRWdRjb =  /*line :1*/ uOWHfVo.Uint64(eESy98)
		case *float32:
			*iFZRWdRjb =  /*line :1*/ math.AviWp5b0( /*line :1*/ uOWHfVo.Uint32(eESy98))
		case *float64:
			*iFZRWdRjb =  /*line :1*/ math.QUB2Kf63p( /*line :1*/ uOWHfVo.Uint64(eESy98))
		case []bool:
			for hLiD5st, kE2XX8fFDWD := range eESy98 {
				iFZRWdRjb[hLiD5st] = kE2XX8fFDWD != 0
			}
		case []int8:
			for hLiD5st, kE2XX8fFDWD := range eESy98 {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ int8(kE2XX8fFDWD)
			}
		case []uint8:
			 /*line :1*/ copy(iFZRWdRjb, eESy98)
		case []int16:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ int16( /*line :1*/ uOWHfVo.Uint16(eESy98[2*hLiD5st:]))
			}
		case []uint16:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ uOWHfVo.Uint16(eESy98[2*hLiD5st:])
			}
		case []int32:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ int32( /*line :1*/ uOWHfVo.Uint32(eESy98[4*hLiD5st:]))
			}
		case []uint32:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ uOWHfVo.Uint32(eESy98[4*hLiD5st:])
			}
		case []int64:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ int64( /*line :1*/ uOWHfVo.Uint64(eESy98[8*hLiD5st:]))
			}
		case []uint64:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ uOWHfVo.Uint64(eESy98[8*hLiD5st:])
			}
		case []float32:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ math.AviWp5b0( /*line :1*/ uOWHfVo.Uint32(eESy98[4*hLiD5st:]))
			}
		case []float64:
			for hLiD5st := range iFZRWdRjb {
				iFZRWdRjb[hLiD5st] =  /*line :1*/ math.QUB2Kf63p( /*line :1*/ uOWHfVo.Uint64(eESy98[8*hLiD5st:]))
			}
		default:
			xnoQ5043QZ = 0
		}
		if xnoQ5043QZ != 0 {
			return nil
		}
	}

	fj8vmzJPe :=  /*line :1*/ reflect.SdHoaIQl(iFZRWdRjb)
	iCkcKiG := -1
	switch  /*line :1*/ fj8vmzJPe.Kind() {
	case reflect.Pointer:
		fj8vmzJPe =  /*line :1*/ fj8vmzJPe.Elem()
		iCkcKiG =  /*line :1*/ oUxzaJFc(fj8vmzJPe)
	case reflect.Slice:
		iCkcKiG =  /*line :1*/ oUxzaJFc(fj8vmzJPe)
	}
	if iCkcKiG < 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("binary.Read: invalid type " +  /*line :1*/ reflect.Cher1a2Fblr(iFZRWdRjb).String())
	}
	hnI4MrXq := &wuj6GST{jvxa6GsrRzKx: uOWHfVo, pej9l_busami:  /*line :1*/ make([]byte, iCkcKiG)}
	if _, ne233xQRrIq :=  /*line :1*/ io.G0zBgKg(fswXI7GD2zM, hnI4MrXq.pej9l_busami); ne233xQRrIq != nil {
		return ne233xQRrIq
	}
	 /*line :1*/ hnI4MrXq.rzJIQ0WaN(fj8vmzJPe)
	return nil
}









func C2AtJXKDP(swaFCH0k io.LXQrGW6BQt, uOWHfVo JUEyK4eMhp, iFZRWdRjb any) error {

	if xnoQ5043QZ :=  /*line :1*/ g1aptfN(iFZRWdRjb); xnoQ5043QZ != 0 {
		eESy98 :=  /*line :1*/ make([]byte, xnoQ5043QZ)
		switch fj8vmzJPe := iFZRWdRjb.(type) {
		case *bool:
			if *fj8vmzJPe {
				eESy98[0] = 1
			} else {
				eESy98[0] = 0
			}
		case bool:
			if fj8vmzJPe {
				eESy98[0] = 1
			} else {
				eESy98[0] = 0
			}
		case []bool:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				if kE2XX8fFDWD {
					eESy98[hLiD5st] = 1
				} else {
					eESy98[hLiD5st] = 0
				}
			}
		case *int8:
			eESy98[0] =  /*line :1*/ byte(*fj8vmzJPe)
		case int8:
			eESy98[0] =  /*line :1*/ byte(fj8vmzJPe)
		case []int8:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				eESy98[hLiD5st] =  /*line :1*/ byte(kE2XX8fFDWD)
			}
		case *uint8:
			eESy98[0] = *fj8vmzJPe
		case uint8:
			eESy98[0] = fj8vmzJPe
		case []uint8:
			eESy98 = fj8vmzJPe
		case *int16:
			 /*line :1*/ uOWHfVo.PutUint16(eESy98,  /*line :1*/ uint16(*fj8vmzJPe))
		case int16:
			 /*line :1*/ uOWHfVo.PutUint16(eESy98,  /*line :1*/ uint16(fj8vmzJPe))
		case []int16:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint16(eESy98[2*hLiD5st:],  /*line :1*/ uint16(kE2XX8fFDWD))
			}
		case *uint16:
			 /*line :1*/ uOWHfVo.PutUint16(eESy98, *fj8vmzJPe)
		case uint16:
			 /*line :1*/ uOWHfVo.PutUint16(eESy98, fj8vmzJPe)
		case []uint16:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint16(eESy98[2*hLiD5st:], kE2XX8fFDWD)
			}
		case *int32:
			 /*line :1*/ uOWHfVo.PutUint32(eESy98,  /*line :1*/ uint32(*fj8vmzJPe))
		case int32:
			 /*line :1*/ uOWHfVo.PutUint32(eESy98,  /*line :1*/ uint32(fj8vmzJPe))
		case []int32:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint32(eESy98[4*hLiD5st:],  /*line :1*/ uint32(kE2XX8fFDWD))
			}
		case *uint32:
			 /*line :1*/ uOWHfVo.PutUint32(eESy98, *fj8vmzJPe)
		case uint32:
			 /*line :1*/ uOWHfVo.PutUint32(eESy98, fj8vmzJPe)
		case []uint32:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint32(eESy98[4*hLiD5st:], kE2XX8fFDWD)
			}
		case *int64:
			 /*line :1*/ uOWHfVo.PutUint64(eESy98,  /*line :1*/ uint64(*fj8vmzJPe))
		case int64:
			 /*line :1*/ uOWHfVo.PutUint64(eESy98,  /*line :1*/ uint64(fj8vmzJPe))
		case []int64:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint64(eESy98[8*hLiD5st:],  /*line :1*/ uint64(kE2XX8fFDWD))
			}
		case *uint64:
			 /*line :1*/ uOWHfVo.PutUint64(eESy98, *fj8vmzJPe)
		case uint64:
			 /*line :1*/ uOWHfVo.PutUint64(eESy98, fj8vmzJPe)
		case []uint64:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint64(eESy98[8*hLiD5st:], kE2XX8fFDWD)
			}
		case *float32:
			 /*line :1*/ uOWHfVo.PutUint32(eESy98,  /*line :1*/ math.FWbZL0aR(*fj8vmzJPe))
		case float32:
			 /*line :1*/ uOWHfVo.PutUint32(eESy98,  /*line :1*/ math.FWbZL0aR(fj8vmzJPe))
		case []float32:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint32(eESy98[4*hLiD5st:],  /*line :1*/ math.FWbZL0aR(kE2XX8fFDWD))
			}
		case *float64:
			 /*line :1*/ uOWHfVo.PutUint64(eESy98,  /*line :1*/ math.GKIRmoHE(*fj8vmzJPe))
		case float64:
			 /*line :1*/ uOWHfVo.PutUint64(eESy98,  /*line :1*/ math.GKIRmoHE(fj8vmzJPe))
		case []float64:
			for hLiD5st, kE2XX8fFDWD := range fj8vmzJPe {
				 /*line :1*/ uOWHfVo.PutUint64(eESy98[8*hLiD5st:],  /*line :1*/ math.GKIRmoHE(kE2XX8fFDWD))
			}
		}
		_, ne233xQRrIq :=  /*line :1*/ swaFCH0k.Write(eESy98)
		return ne233xQRrIq
	}

	fj8vmzJPe :=  /*line :1*/ reflect.XF6ikODQ8h( /*line :1*/ reflect.SdHoaIQl(iFZRWdRjb))
	iCkcKiG :=  /*line :1*/ oUxzaJFc(fj8vmzJPe)
	if iCkcKiG < 0 {
		return  /*line :1*/ errors.Su6g6hRoi9X("binary.Write: some values are not fixed-sized in type " +  /*line :1*/ reflect.Cher1a2Fblr(iFZRWdRjb).String())
	}
	cauYncT8 :=  /*line :1*/ make([]byte, iCkcKiG)
	yGPN1e := &xB02BrtG6K{jvxa6GsrRzKx: uOWHfVo, pej9l_busami: cauYncT8}
	 /*line :1*/ yGPN1e.rzJIQ0WaN(fj8vmzJPe)
	_, ne233xQRrIq :=  /*line :1*/ swaFCH0k.Write(cauYncT8)
	return ne233xQRrIq
}




func TVdJGXMavOA(fj8vmzJPe any) int {
	return  /*line :1*/ oUxzaJFc( /*line :1*/ reflect.XF6ikODQ8h( /*line :1*/ reflect.SdHoaIQl(fj8vmzJPe)))
}

var m35496DK sync.Eim1b6bNi9IM	





func oUxzaJFc(fj8vmzJPe reflect.Value) int {
	switch  /*line :1*/ fj8vmzJPe.Kind() {
	case reflect.Slice:
		if fVnaq0r06 :=  /*line :1*/ iAK34Pq_cS( /*line :1*/ fj8vmzJPe.Type().Elem()); fVnaq0r06 >= 0 {
			return fVnaq0r06 *  /*line :1*/ fj8vmzJPe.Len()
		}

	case reflect.Struct:
		wlzbDShjeEr_ :=  /*line :1*/ fj8vmzJPe.Type()
		if iCkcKiG, sZHjzT5KN :=  /*line :1*/ m35496DK.Load(wlzbDShjeEr_); sZHjzT5KN {
			return iCkcKiG.(int)
		}
		iCkcKiG :=  /*line :1*/ iAK34Pq_cS(wlzbDShjeEr_)
		 /*line :1*/ m35496DK.Store(wlzbDShjeEr_, iCkcKiG)
		return iCkcKiG

	default:
		if  /*line :1*/ fj8vmzJPe.IsValid() {
			return  /*line :1*/ iAK34Pq_cS( /*line :1*/ fj8vmzJPe.Type())
		}
	}

	return -1
}


func iAK34Pq_cS(wlzbDShjeEr_ reflect.YJh989LTZsVX) int {
	switch  /*line :1*/ wlzbDShjeEr_.Kind() {
	case reflect.Array:
		if fVnaq0r06 :=  /*line :1*/ iAK34Pq_cS( /*line :1*/ wlzbDShjeEr_.Elem()); fVnaq0r06 >= 0 {
			return fVnaq0r06 *  /*line :1*/ wlzbDShjeEr_.Len()
		}

	case reflect.Struct:
		nvk7XlhDS := 0
		for hLiD5st, xnoQ5043QZ := 0,  /*line :1*/ wlzbDShjeEr_.NumField(); hLiD5st < xnoQ5043QZ; hLiD5st++ {
			fVnaq0r06 :=  /*line :1*/ iAK34Pq_cS( /*line :1*/ wlzbDShjeEr_.Field(hLiD5st).S1miDNBAV8a)
			if fVnaq0r06 < 0 {
				return -1
			}
			nvk7XlhDS += fVnaq0r06
		}
		return nvk7XlhDS

	case reflect.Bool,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return  /*line :1*/ int( /*line :1*/ wlzbDShjeEr_.Size())
	}

	return -1
}

type batPgq struct {
	jvxa6GsrRzKx	JUEyK4eMhp
	pej9l_busami	[]byte
	v3vtoPYg	int
}

type wuj6GST batPgq
type xB02BrtG6K batPgq

func (hnI4MrXq *wuj6GST) o1hnJaIqIFfS() bool {
	kE2XX8fFDWD := hnI4MrXq.pej9l_busami[hnI4MrXq.v3vtoPYg]
	hnI4MrXq.v3vtoPYg++
	return kE2XX8fFDWD != 0
}

func (yGPN1e *xB02BrtG6K) o1hnJaIqIFfS(kE2XX8fFDWD bool) {
	if kE2XX8fFDWD {
		yGPN1e.pej9l_busami[yGPN1e.v3vtoPYg] = 1
	} else {
		yGPN1e.pej9l_busami[yGPN1e.v3vtoPYg] = 0
	}
	yGPN1e.v3vtoPYg++
}

func (hnI4MrXq *wuj6GST) vg2jsiJFaAO() uint8 {
	kE2XX8fFDWD := hnI4MrXq.pej9l_busami[hnI4MrXq.v3vtoPYg]
	hnI4MrXq.v3vtoPYg++
	return kE2XX8fFDWD
}

func (yGPN1e *xB02BrtG6K) vg2jsiJFaAO(kE2XX8fFDWD uint8) {
	yGPN1e.pej9l_busami[yGPN1e.v3vtoPYg] = kE2XX8fFDWD
	yGPN1e.v3vtoPYg++
}

func (hnI4MrXq *wuj6GST) xAgFhKN_oo() uint16 {
	kE2XX8fFDWD :=  /*line :1*/ hnI4MrXq.jvxa6GsrRzKx.Uint16(hnI4MrXq.pej9l_busami[hnI4MrXq.v3vtoPYg : hnI4MrXq.v3vtoPYg+2])
	hnI4MrXq.v3vtoPYg += 2
	return kE2XX8fFDWD
}

func (yGPN1e *xB02BrtG6K) xAgFhKN_oo(kE2XX8fFDWD uint16) {
	 /*line :1*/ yGPN1e.jvxa6GsrRzKx.PutUint16(yGPN1e.pej9l_busami[yGPN1e.v3vtoPYg:yGPN1e.v3vtoPYg+2], kE2XX8fFDWD)
	yGPN1e.v3vtoPYg += 2
}

func (hnI4MrXq *wuj6GST) fhxYoDd9() uint32 {
	kE2XX8fFDWD :=  /*line :1*/ hnI4MrXq.jvxa6GsrRzKx.Uint32(hnI4MrXq.pej9l_busami[hnI4MrXq.v3vtoPYg : hnI4MrXq.v3vtoPYg+4])
	hnI4MrXq.v3vtoPYg += 4
	return kE2XX8fFDWD
}

func (yGPN1e *xB02BrtG6K) fhxYoDd9(kE2XX8fFDWD uint32) {
	 /*line :1*/ yGPN1e.jvxa6GsrRzKx.PutUint32(yGPN1e.pej9l_busami[yGPN1e.v3vtoPYg:yGPN1e.v3vtoPYg+4], kE2XX8fFDWD)
	yGPN1e.v3vtoPYg += 4
}

func (hnI4MrXq *wuj6GST) ouG4vg0tGVK() uint64 {
	kE2XX8fFDWD :=  /*line :1*/ hnI4MrXq.jvxa6GsrRzKx.Uint64(hnI4MrXq.pej9l_busami[hnI4MrXq.v3vtoPYg : hnI4MrXq.v3vtoPYg+8])
	hnI4MrXq.v3vtoPYg += 8
	return kE2XX8fFDWD
}

func (yGPN1e *xB02BrtG6K) ouG4vg0tGVK(kE2XX8fFDWD uint64) {
	 /*line :1*/ yGPN1e.jvxa6GsrRzKx.PutUint64(yGPN1e.pej9l_busami[yGPN1e.v3vtoPYg:yGPN1e.v3vtoPYg+8], kE2XX8fFDWD)
	yGPN1e.v3vtoPYg += 8
}

func (hnI4MrXq *wuj6GST) kEeFhEdmkw() int8	{ return  /*line :1*/ int8( /*line :1*/ hnI4MrXq.vg2jsiJFaAO()) }

func (yGPN1e *xB02BrtG6K) kEeFhEdmkw(kE2XX8fFDWD int8)	{  /*line :1*/ yGPN1e.vg2jsiJFaAO( /*line :1*/ uint8(kE2XX8fFDWD)) }

func (hnI4MrXq *wuj6GST) oQgHF5f() int16	{ return  /*line :1*/ int16( /*line :1*/ hnI4MrXq.xAgFhKN_oo()) }

func (yGPN1e *xB02BrtG6K) oQgHF5f(kE2XX8fFDWD int16)	{  /*line :1*/ yGPN1e.xAgFhKN_oo( /*line :1*/ uint16(kE2XX8fFDWD)) }

func (hnI4MrXq *wuj6GST) cn5XAi() int32	{ return  /*line :1*/ int32( /*line :1*/ hnI4MrXq.fhxYoDd9()) }

func (yGPN1e *xB02BrtG6K) cn5XAi(kE2XX8fFDWD int32)	{  /*line :1*/ yGPN1e.fhxYoDd9( /*line :1*/ uint32(kE2XX8fFDWD)) }

func (hnI4MrXq *wuj6GST) kpcaaBal64() int64	{ return  /*line :1*/ int64( /*line :1*/ hnI4MrXq.ouG4vg0tGVK()) }

func (yGPN1e *xB02BrtG6K) kpcaaBal64(kE2XX8fFDWD int64)	{  /*line :1*/ yGPN1e.ouG4vg0tGVK( /*line :1*/ uint64(kE2XX8fFDWD)) }

func (hnI4MrXq *wuj6GST) rzJIQ0WaN(fj8vmzJPe reflect.Value) {
	switch  /*line :1*/ fj8vmzJPe.Kind() {
	case reflect.Array:
		y41M0B :=  /*line :1*/ fj8vmzJPe.Len()
		for hLiD5st := 0; hLiD5st < y41M0B; hLiD5st++ {
			 /*line :1*/ hnI4MrXq.rzJIQ0WaN( /*line :1*/ fj8vmzJPe.Index(hLiD5st))
		}

	case reflect.Struct:
		wlzbDShjeEr_ :=  /*line :1*/ fj8vmzJPe.Type()
		y41M0B :=  /*line :1*/ fj8vmzJPe.NumField()
		for hLiD5st := 0; hLiD5st < y41M0B; hLiD5st++ {

			if fj8vmzJPe :=  /*line :1*/ fj8vmzJPe.Field(hLiD5st);  /*line :1*/ fj8vmzJPe.CanSet() ||  /*line :1*/ wlzbDShjeEr_.Field(hLiD5st).ZOSBHkJz != "_" {
				 /*line :1*/ hnI4MrXq.rzJIQ0WaN(fj8vmzJPe)
			} else {
				 /*line :1*/ hnI4MrXq.m8jmrt(fj8vmzJPe)
			}
		}

	case reflect.Slice:
		y41M0B :=  /*line :1*/ fj8vmzJPe.Len()
		for hLiD5st := 0; hLiD5st < y41M0B; hLiD5st++ {
			 /*line :1*/ hnI4MrXq.rzJIQ0WaN( /*line :1*/ fj8vmzJPe.Index(hLiD5st))
		}

	case reflect.Bool:
		 /*line :1*/ fj8vmzJPe.SetBool( /*line :1*/ hnI4MrXq.o1hnJaIqIFfS())

	case reflect.Int8:
		 /*line :1*/ fj8vmzJPe.SetInt( /*line :1*/ int64( /*line :1*/ hnI4MrXq.kEeFhEdmkw()))
	case reflect.Int16:
		 /*line :1*/ fj8vmzJPe.SetInt( /*line :1*/ int64( /*line :1*/ hnI4MrXq.oQgHF5f()))
	case reflect.Int32:
		 /*line :1*/ fj8vmzJPe.SetInt( /*line :1*/ int64( /*line :1*/ hnI4MrXq.cn5XAi()))
	case reflect.Int64:
		 /*line :1*/ fj8vmzJPe.SetInt( /*line :1*/ hnI4MrXq.kpcaaBal64())

	case reflect.Uint8:
		 /*line :1*/ fj8vmzJPe.SetUint( /*line :1*/ uint64( /*line :1*/ hnI4MrXq.vg2jsiJFaAO()))
	case reflect.Uint16:
		 /*line :1*/ fj8vmzJPe.SetUint( /*line :1*/ uint64( /*line :1*/ hnI4MrXq.xAgFhKN_oo()))
	case reflect.Uint32:
		 /*line :1*/ fj8vmzJPe.SetUint( /*line :1*/ uint64( /*line :1*/ hnI4MrXq.fhxYoDd9()))
	case reflect.Uint64:
		 /*line :1*/ fj8vmzJPe.SetUint( /*line :1*/ hnI4MrXq.ouG4vg0tGVK())

	case reflect.Float32:
		 /*line :1*/ fj8vmzJPe.SetFloat( /*line :1*/ float64( /*line :1*/ math.AviWp5b0( /*line :1*/ hnI4MrXq.fhxYoDd9())))
	case reflect.Float64:
		 /*line :1*/ fj8vmzJPe.SetFloat( /*line :1*/ math.QUB2Kf63p( /*line :1*/ hnI4MrXq.ouG4vg0tGVK()))

	case reflect.Complex64:
		 /*line :1*/ fj8vmzJPe.SetComplex( /*line :1*/ complex(
			 /*line :1*/ float64( /*line :1*/ math.AviWp5b0( /*line :1*/ hnI4MrXq.fhxYoDd9())),
			 /*line :1*/ float64( /*line :1*/ math.AviWp5b0( /*line :1*/ hnI4MrXq.fhxYoDd9())),
		))
	case reflect.Complex128:
		 /*line :1*/ fj8vmzJPe.SetComplex( /*line :1*/ complex(
			 /*line :1*/ math.QUB2Kf63p( /*line :1*/ hnI4MrXq.ouG4vg0tGVK()),
			 /*line :1*/ math.QUB2Kf63p( /*line :1*/ hnI4MrXq.ouG4vg0tGVK()),
		))
	}
}

func (yGPN1e *xB02BrtG6K) rzJIQ0WaN(fj8vmzJPe reflect.Value) {
	switch  /*line :1*/ fj8vmzJPe.Kind() {
	case reflect.Array:
		y41M0B :=  /*line :1*/ fj8vmzJPe.Len()
		for hLiD5st := 0; hLiD5st < y41M0B; hLiD5st++ {
			 /*line :1*/ yGPN1e.rzJIQ0WaN( /*line :1*/ fj8vmzJPe.Index(hLiD5st))
		}

	case reflect.Struct:
		wlzbDShjeEr_ :=  /*line :1*/ fj8vmzJPe.Type()
		y41M0B :=  /*line :1*/ fj8vmzJPe.NumField()
		for hLiD5st := 0; hLiD5st < y41M0B; hLiD5st++ {

			if fj8vmzJPe :=  /*line :1*/ fj8vmzJPe.Field(hLiD5st);  /*line :1*/ fj8vmzJPe.CanSet() ||  /*line :1*/ wlzbDShjeEr_.Field(hLiD5st).ZOSBHkJz != "_" {
				 /*line :1*/ yGPN1e.rzJIQ0WaN(fj8vmzJPe)
			} else {
				 /*line :1*/ yGPN1e.m8jmrt(fj8vmzJPe)
			}
		}

	case reflect.Slice:
		y41M0B :=  /*line :1*/ fj8vmzJPe.Len()
		for hLiD5st := 0; hLiD5st < y41M0B; hLiD5st++ {
			 /*line :1*/ yGPN1e.rzJIQ0WaN( /*line :1*/ fj8vmzJPe.Index(hLiD5st))
		}

	case reflect.Bool:
		 /*line :1*/ yGPN1e.o1hnJaIqIFfS( /*line :1*/ fj8vmzJPe.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch  /*line :1*/ fj8vmzJPe.Type().Kind() {
		case reflect.Int8:
			 /*line :1*/ yGPN1e.kEeFhEdmkw( /*line :1*/ int8( /*line :1*/ fj8vmzJPe.Int()))
		case reflect.Int16:
			 /*line :1*/ yGPN1e.oQgHF5f( /*line :1*/ int16( /*line :1*/ fj8vmzJPe.Int()))
		case reflect.Int32:
			 /*line :1*/ yGPN1e.cn5XAi( /*line :1*/ int32( /*line :1*/ fj8vmzJPe.Int()))
		case reflect.Int64:
			 /*line :1*/ yGPN1e.kpcaaBal64( /*line :1*/ fj8vmzJPe.Int())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		switch  /*line :1*/ fj8vmzJPe.Type().Kind() {
		case reflect.Uint8:
			 /*line :1*/ yGPN1e.vg2jsiJFaAO( /*line :1*/ uint8( /*line :1*/ fj8vmzJPe.Uint()))
		case reflect.Uint16:
			 /*line :1*/ yGPN1e.xAgFhKN_oo( /*line :1*/ uint16( /*line :1*/ fj8vmzJPe.Uint()))
		case reflect.Uint32:
			 /*line :1*/ yGPN1e.fhxYoDd9( /*line :1*/ uint32( /*line :1*/ fj8vmzJPe.Uint()))
		case reflect.Uint64:
			 /*line :1*/ yGPN1e.ouG4vg0tGVK( /*line :1*/ fj8vmzJPe.Uint())
		}

	case reflect.Float32, reflect.Float64:
		switch  /*line :1*/ fj8vmzJPe.Type().Kind() {
		case reflect.Float32:
			 /*line :1*/ yGPN1e.fhxYoDd9( /*line :1*/ math.FWbZL0aR( /*line :1*/ float32( /*line :1*/ fj8vmzJPe.Float())))
		case reflect.Float64:
			 /*line :1*/ yGPN1e.ouG4vg0tGVK( /*line :1*/ math.GKIRmoHE( /*line :1*/ fj8vmzJPe.Float()))
		}

	case reflect.Complex64, reflect.Complex128:
		switch  /*line :1*/ fj8vmzJPe.Type().Kind() {
		case reflect.Complex64:
			kE2XX8fFDWD :=  /*line :1*/ fj8vmzJPe.Complex()
			 /*line :1*/ yGPN1e.fhxYoDd9( /*line :1*/ math.FWbZL0aR( /*line :1*/ float32( /*line :1*/ real(kE2XX8fFDWD))))
			 /*line :1*/ yGPN1e.fhxYoDd9( /*line :1*/ math.FWbZL0aR( /*line :1*/ float32( /*line :1*/ imag(kE2XX8fFDWD))))
		case reflect.Complex128:
			kE2XX8fFDWD :=  /*line :1*/ fj8vmzJPe.Complex()
			 /*line :1*/ yGPN1e.ouG4vg0tGVK( /*line :1*/ math.GKIRmoHE( /*line :1*/ real(kE2XX8fFDWD)))
			 /*line :1*/ yGPN1e.ouG4vg0tGVK( /*line :1*/ math.GKIRmoHE( /*line :1*/ imag(kE2XX8fFDWD)))
		}
	}
}

func (hnI4MrXq *wuj6GST) m8jmrt(fj8vmzJPe reflect.Value) {
	hnI4MrXq.v3vtoPYg +=  /*line :1*/ oUxzaJFc(fj8vmzJPe)
}

func (yGPN1e *xB02BrtG6K) m8jmrt(fj8vmzJPe reflect.Value) {
	xnoQ5043QZ :=  /*line :1*/ oUxzaJFc(fj8vmzJPe)
	qXg1hra9 := yGPN1e.pej9l_busami[yGPN1e.v3vtoPYg : yGPN1e.v3vtoPYg+xnoQ5043QZ]
	for hLiD5st := range qXg1hra9 {
		qXg1hra9[hLiD5st] = 0
	}
	yGPN1e.v3vtoPYg += xnoQ5043QZ
}



func g1aptfN(iFZRWdRjb any) int {
	switch iFZRWdRjb := iFZRWdRjb.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []bool:
		return  /*line :1*/ len(iFZRWdRjb)
	case []int8:
		return  /*line :1*/ len(iFZRWdRjb)
	case []uint8:
		return  /*line :1*/ len(iFZRWdRjb)
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return 2 *  /*line :1*/ len(iFZRWdRjb)
	case []uint16:
		return 2 *  /*line :1*/ len(iFZRWdRjb)
	case int32, uint32, *int32, *uint32:
		return 4
	case []int32:
		return 4 *  /*line :1*/ len(iFZRWdRjb)
	case []uint32:
		return 4 *  /*line :1*/ len(iFZRWdRjb)
	case int64, uint64, *int64, *uint64:
		return 8
	case []int64:
		return 8 *  /*line :1*/ len(iFZRWdRjb)
	case []uint64:
		return 8 *  /*line :1*/ len(iFZRWdRjb)
	case float32, *float32:
		return 4
	case float64, *float64:
		return 8
	case []float32:
		return 4 *  /*line :1*/ len(iFZRWdRjb)
	case []float64:
		return 8 *  /*line :1*/ len(iFZRWdRjb)
	}
	return 0
}
