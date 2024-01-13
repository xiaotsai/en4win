//line :1
package yLm9uBQG

import (
	errors "iAHoxjmM"
	io "xI9ai2djaJ2"
)


const (
	MaxVarintLen16	= 3
	MaxVarintLen32	= 5
	MaxVarintLen64	= 10
)



func NCM8EuyC7E9(cauYncT8 []byte, kE2XX8fFDWD uint64) []byte {
	for kE2XX8fFDWD >= 0x80 {
		cauYncT8 =  /*line :1*/ append(cauYncT8,  /*line :1*/ byte(kE2XX8fFDWD)|0x80)
		kE2XX8fFDWD >>= 7
	}
	return  /*line :1*/ append(cauYncT8,  /*line :1*/ byte(kE2XX8fFDWD))
}



func NrfXYGkUbYX(cauYncT8 []byte, kE2XX8fFDWD uint64) int {
	hLiD5st := 0
	for kE2XX8fFDWD >= 0x80 {
		cauYncT8[hLiD5st] =  /*line :1*/ byte(kE2XX8fFDWD) | 0x80
		kE2XX8fFDWD >>= 7
		hLiD5st++
	}
	cauYncT8[hLiD5st] =  /*line :1*/ byte(kE2XX8fFDWD)
	return hLiD5st + 1
}








func KVTzDl(cauYncT8 []byte) (uint64, int) {
	var kE2XX8fFDWD uint64
	var fVnaq0r06 uint
	for hLiD5st, nWegNNdJn_ := range cauYncT8 {
		if hLiD5st == MaxVarintLen64 {

			return 0, -(hLiD5st + 1)
		}
		if nWegNNdJn_ < 0x80 {
			if hLiD5st == MaxVarintLen64-1 && nWegNNdJn_ > 1 {
				return 0, -(hLiD5st + 1)
			}
			return kE2XX8fFDWD |  /*line :1*/ uint64(nWegNNdJn_)<<fVnaq0r06, hLiD5st + 1
		}
		kE2XX8fFDWD |=  /*line :1*/ uint64(nWegNNdJn_&0x7f) << fVnaq0r06
		fVnaq0r06 += 7
	}
	return 0, 0
}



func Pn8WSvTBx(cauYncT8 []byte, kE2XX8fFDWD int64) []byte {
	l1jknZGq :=  /*line :1*/ uint64(kE2XX8fFDWD) << 1
	if kE2XX8fFDWD < 0 {
		l1jknZGq = ^l1jknZGq
	}
	return  /*line :1*/ NCM8EuyC7E9(cauYncT8, l1jknZGq)
}



func J3sS_THpJRh8(cauYncT8 []byte, kE2XX8fFDWD int64) int {
	l1jknZGq :=  /*line :1*/ uint64(kE2XX8fFDWD) << 1
	if kE2XX8fFDWD < 0 {
		l1jknZGq = ^l1jknZGq
	}
	return  /*line :1*/ NrfXYGkUbYX(cauYncT8, l1jknZGq)
}








func GAHgUm26Hl1(cauYncT8 []byte) (int64, int) {
	l1jknZGq, xnoQ5043QZ :=  /*line :1*/ KVTzDl(cauYncT8)
	kE2XX8fFDWD :=  /*line :1*/ int64(l1jknZGq >> 1)
	if l1jknZGq&1 != 0 {
		kE2XX8fFDWD = ^kE2XX8fFDWD
	}
	return kE2XX8fFDWD, xnoQ5043QZ
}

var f5Dngd =  /*line :1*/ errors.Su6g6hRoi9X("binary: varint overflows a 64-bit integer")





func Z_CEAAwi6n_(fswXI7GD2zM io.Nbptmbb) (uint64, error) {
	var kE2XX8fFDWD uint64
	var fVnaq0r06 uint
	for hLiD5st := 0; hLiD5st < MaxVarintLen64; hLiD5st++ {
		nWegNNdJn_, ne233xQRrIq :=  /*line :1*/ fswXI7GD2zM.ReadByte()
		if ne233xQRrIq != nil {
			if hLiD5st > 0 && ne233xQRrIq == io.K5Sqco {
				ne233xQRrIq = io.UASgojMNPA
			}
			return kE2XX8fFDWD, ne233xQRrIq
		}
		if nWegNNdJn_ < 0x80 {
			if hLiD5st == MaxVarintLen64-1 && nWegNNdJn_ > 1 {
				return kE2XX8fFDWD, f5Dngd
			}
			return kE2XX8fFDWD |  /*line :1*/ uint64(nWegNNdJn_)<<fVnaq0r06, nil
		}
		kE2XX8fFDWD |=  /*line :1*/ uint64(nWegNNdJn_&0x7f) << fVnaq0r06
		fVnaq0r06 += 7
	}
	return kE2XX8fFDWD, f5Dngd
}





func GXIJfFE39(fswXI7GD2zM io.Nbptmbb) (int64, error) {
	l1jknZGq, ne233xQRrIq :=  /*line :1*/ Z_CEAAwi6n_(fswXI7GD2zM)
	kE2XX8fFDWD :=  /*line :1*/ int64(l1jknZGq >> 1)
	if l1jknZGq&1 != 0 {
		kE2XX8fFDWD = ^kE2XX8fFDWD
	}
	return kE2XX8fFDWD, ne233xQRrIq
}
