//line :1




package sXttzp

import (
	reflect "reflect"
	sort "gzHaha55n"
)



type SortedMap struct {
	Key	[]reflect.Value
	Value	[]reflect.Value
}

func (aefHpIG6W9 *SortedMap) Len() int	{ return  /*line :1*/ len(aefHpIG6W9.Key) }
func (aefHpIG6W9 *SortedMap) Less(iC31npgh, oc3BfUNDs7D int) bool {
	return  /*line :1*/ cHdfwXBMa(aefHpIG6W9.Key[iC31npgh], aefHpIG6W9.Key[oc3BfUNDs7D]) < 0
}
func (aefHpIG6W9 *SortedMap) Swap(iC31npgh, oc3BfUNDs7D int) {
	aefHpIG6W9.Key[iC31npgh], aefHpIG6W9.Key[oc3BfUNDs7D] = aefHpIG6W9.Key[oc3BfUNDs7D], aefHpIG6W9.Key[iC31npgh]
	aefHpIG6W9.Value[iC31npgh], aefHpIG6W9.Value[oc3BfUNDs7D] = aefHpIG6W9.Value[oc3BfUNDs7D], aefHpIG6W9.Value[iC31npgh]
}



















func HhvK6ljdH6m(c6lczQovIb_G reflect.Value) *SortedMap {
	if  /*line :1*/ c6lczQovIb_G.Type().Kind() != reflect.Map {
		return nil
	}

	qzrWLIXz :=  /*line :1*/ c6lczQovIb_G.Len()
	qrAodDfGC :=  /*line :1*/ make([]reflect.Value, 0, qzrWLIXz)
	qKyYni :=  /*line :1*/ make([]reflect.Value, 0, qzrWLIXz)
	bN4pgM :=  /*line :1*/ c6lczQovIb_G.MapRange()
	for  /*line :1*/ bN4pgM.Next() {
		qrAodDfGC =  /*line :1*/ append(qrAodDfGC,  /*line :1*/ bN4pgM.Key())
		qKyYni =  /*line :1*/ append(qKyYni,  /*line :1*/ bN4pgM.Value())
	}
	dnPXVSvr := &SortedMap{
		Key:	qrAodDfGC,
		Value:	qKyYni,
	}
	 /*line :1*/ sort.IUVpc45h_Y3(dnPXVSvr)
	return dnPXVSvr
}





func cHdfwXBMa(iShXraa6, v6tTfMvhW reflect.Value) int {
	aHkCKzR4IgPO, ifIunD4J89HQ :=  /*line :1*/ iShXraa6.Type(),  /*line :1*/ v6tTfMvhW.Type()
	if aHkCKzR4IgPO != ifIunD4J89HQ {
		return -1
	}
	switch  /*line :1*/ iShXraa6.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		wQeg14OKq9s, eg8DuDF_uLc :=  /*line :1*/ iShXraa6.Int(),  /*line :1*/ v6tTfMvhW.Int()
		switch {
		case wQeg14OKq9s < eg8DuDF_uLc:
			return -1
		case wQeg14OKq9s > eg8DuDF_uLc:
			return 1
		default:
			return 0
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		wQeg14OKq9s, eg8DuDF_uLc :=  /*line :1*/ iShXraa6.Uint(),  /*line :1*/ v6tTfMvhW.Uint()
		switch {
		case wQeg14OKq9s < eg8DuDF_uLc:
			return -1
		case wQeg14OKq9s > eg8DuDF_uLc:
			return 1
		default:
			return 0
		}
	case reflect.String:
		wQeg14OKq9s, eg8DuDF_uLc :=  /*line :1*/ iShXraa6.String(),  /*line :1*/ v6tTfMvhW.String()
		switch {
		case wQeg14OKq9s < eg8DuDF_uLc:
			return -1
		case wQeg14OKq9s > eg8DuDF_uLc:
			return 1
		default:
			return 0
		}
	case reflect.Float32, reflect.Float64:
		return  /*line :1*/ bm0EGzipUp0( /*line :1*/ iShXraa6.Float(),  /*line :1*/ v6tTfMvhW.Float())
	case reflect.Complex64, reflect.Complex128:
		wQeg14OKq9s, eg8DuDF_uLc :=  /*line :1*/ iShXraa6.Complex(),  /*line :1*/ v6tTfMvhW.Complex()
		if zP_nBLQXm3JH :=  /*line :1*/ bm0EGzipUp0( /*line :1*/ real(wQeg14OKq9s),  /*line :1*/ real(eg8DuDF_uLc)); zP_nBLQXm3JH != 0 {
			return zP_nBLQXm3JH
		}
		return  /*line :1*/ bm0EGzipUp0( /*line :1*/ imag(wQeg14OKq9s),  /*line :1*/ imag(eg8DuDF_uLc))
	case reflect.Bool:
		wQeg14OKq9s, eg8DuDF_uLc :=  /*line :1*/ iShXraa6.Bool(),  /*line :1*/ v6tTfMvhW.Bool()
		switch {
		case wQeg14OKq9s == eg8DuDF_uLc:
			return 0
		case wQeg14OKq9s:
			return 1
		default:
			return -1
		}
	case reflect.Pointer, reflect.UnsafePointer:
		wQeg14OKq9s, eg8DuDF_uLc :=  /*line :1*/ iShXraa6.Pointer(),  /*line :1*/ v6tTfMvhW.Pointer()
		switch {
		case wQeg14OKq9s < eg8DuDF_uLc:
			return -1
		case wQeg14OKq9s > eg8DuDF_uLc:
			return 1
		default:
			return 0
		}
	case reflect.Chan:
		if zP_nBLQXm3JH, _JaMqc9ISqhc :=  /*line :1*/ jXt1laji(iShXraa6, v6tTfMvhW); _JaMqc9ISqhc {
			return zP_nBLQXm3JH
		}
		yj6fCOJ9a, cTaUACa :=  /*line :1*/ iShXraa6.Pointer(),  /*line :1*/ v6tTfMvhW.Pointer()
		switch {
		case yj6fCOJ9a < cTaUACa:
			return -1
		case yj6fCOJ9a > cTaUACa:
			return 1
		default:
			return 0
		}
	case reflect.Struct:
		for iC31npgh := 0; iC31npgh <  /*line :1*/ iShXraa6.NumField(); iC31npgh++ {
			if zP_nBLQXm3JH :=  /*line :1*/ cHdfwXBMa( /*line :1*/ iShXraa6.Field(iC31npgh),  /*line :1*/ v6tTfMvhW.Field(iC31npgh)); zP_nBLQXm3JH != 0 {
				return zP_nBLQXm3JH
			}
		}
		return 0
	case reflect.Array:
		for iC31npgh := 0; iC31npgh <  /*line :1*/ iShXraa6.Len(); iC31npgh++ {
			if zP_nBLQXm3JH :=  /*line :1*/ cHdfwXBMa( /*line :1*/ iShXraa6.Index(iC31npgh),  /*line :1*/ v6tTfMvhW.Index(iC31npgh)); zP_nBLQXm3JH != 0 {
				return zP_nBLQXm3JH
			}
		}
		return 0
	case reflect.Interface:
		if zP_nBLQXm3JH, _JaMqc9ISqhc :=  /*line :1*/ jXt1laji(iShXraa6, v6tTfMvhW); _JaMqc9ISqhc {
			return zP_nBLQXm3JH
		}
		zP_nBLQXm3JH :=  /*line :1*/ cHdfwXBMa( /*line :1*/ reflect.SdHoaIQl( /*line :1*/ iShXraa6.Elem().Type()),  /*line :1*/ reflect.SdHoaIQl( /*line :1*/ v6tTfMvhW.Elem().Type()))
		if zP_nBLQXm3JH != 0 {
			return zP_nBLQXm3JH
		}
		return  /*line :1*/ cHdfwXBMa( /*line :1*/ iShXraa6.Elem(),  /*line :1*/ v6tTfMvhW.Elem())
	default:

		 /*line :1*/ panic("bad type in compare: " +  /*line :1*/ aHkCKzR4IgPO.String())
	}
}






func jXt1laji(iShXraa6, v6tTfMvhW reflect.Value) (int, bool) {
	if  /*line :1*/ iShXraa6.IsNil() {
		if  /*line :1*/ v6tTfMvhW.IsNil() {
			return 0, true
		}
		return -1, true
	}
	if  /*line :1*/ v6tTfMvhW.IsNil() {
		return 1, true
	}
	return 0, false
}


func bm0EGzipUp0(wQeg14OKq9s, eg8DuDF_uLc float64) int {
	switch {
	case  /*line :1*/ blid2Ru(wQeg14OKq9s):
		return -1
	case  /*line :1*/ blid2Ru(eg8DuDF_uLc):
		return 1
	case wQeg14OKq9s < eg8DuDF_uLc:
		return -1
	case wQeg14OKq9s > eg8DuDF_uLc:
		return 1
	}
	return 0
}

func blid2Ru(wQeg14OKq9s float64) bool {
	return wQeg14OKq9s != wQeg14OKq9s
}
