//line :1

package baLhtK0

import (
	bytes "wLlfRPv"
	errors "iAHoxjmM"
	fmt "kFsoCfy5zWG"
	log "pBPwKYPji2"
	os "hPMrClpC"
	reflect "reflect"
	"runtime"
	strconv "zBESu2SrRjP"
	strings "fQXlzv"
	sync "sync"
	time "fRAfQd_"

	ole "fuA83L"
	oleutil "uBCMx3OoU"
)

var icTM92nUTr =  /*line :1*/ log.Twasyw(os.Vrq37o, "", log.LstdFlags)

var (
	H1CZpWHzG	=  /*line :1*/ errors.Su6g6hRoi9X("wmi: invalid entity type")
	
	
	W0_gOPjFGm	=  /*line :1*/ errors.Su6g6hRoi9X("wmi: create object returned nil")
	anbyikS1Kv7i	sync.DIRWe11KYlYa
)


const S_FALSE = 0x00000001


func WLLguh7(cMO_44UJBCnD string, aSce7LEv65fL interface{}, wmugeT_TsxU8 string) error {
	return  /*line :1*/ TybA4yw3(cMO_44UJBCnD, aSce7LEv65fL, nil, wmugeT_TsxU8)
}














func TybA4yw3(cMO_44UJBCnD string, aSce7LEv65fL interface{}, pvQQf8ap ...interface{}) error {
	if R7V4IcErEd.ZLVM4mNGOf == nil {
		return  /*line :1*/ R7V4IcErEd.Query(cMO_44UJBCnD, aSce7LEv65fL, pvQQf8ap...)
	}
	return  /*line :1*/ R7V4IcErEd.ZLVM4mNGOf.Query(cMO_44UJBCnD, aSce7LEv65fL, pvQQf8ap...)
}





func FJQRGrN0(pvQQf8ap []interface{}, gb_Ke0eQ, bCihZnk string, p4Tf0EE []interface{}) (int32, error) {
	return  /*line :1*/ R7V4IcErEd.CallMethod(pvQQf8ap, gb_Ke0eQ, bCihZnk, p4Tf0EE)
}




type LlaCLems struct {
	
	
	
	
	
	TYHmyvQ0	bool

	
	
	
	
	
	CQ8rI1v4m	bool

	
	
	
	
	
	YevekfvvbB1S	bool

	
	
	
	ZLVM4mNGOf	*NcjbAXk6
}


var R7V4IcErEd = &LlaCLems{}



func (agZTN4UVka *LlaCLems) uQUjdDex(pvQQf8ap ...interface{}) (*ole.IDispatch, func(), error) {
	var xx_hM9Gbk *ole.IUnknown
	var qnAGlFM6vP7 *ole.IDispatch
	var fAI_4V2jCrXK *ole.KEVetAOpxl0D

	fwMWfH := func() {
		if fAI_4V2jCrXK != nil {
			 /*line :1*/ fAI_4V2jCrXK.Clear()
		}
		if qnAGlFM6vP7 != nil {
			 /*line :1*/ qnAGlFM6vP7.Release()
		}
		if xx_hM9Gbk != nil {
			 /*line :1*/ xx_hM9Gbk.Release()
		}
		 /*line :1*/ ole.DvOlryZ5()
	}

	
	var cp3hZaJi3b error
	defer func() {
		if  /*line :1*/ cp3hZaJi3b != nil {
			 /*line :1*/ fwMWfH()
		}
	}()

	cp3hZaJi3b =  /*line :1*/ ole.KKzYzCZuN(0, ole.COINIT_MULTITHREADED)
	if cp3hZaJi3b != nil {
		kMyMGKa :=  /*line :1*/ cp3hZaJi3b.(*ole.DeOFa72M).Code()
		if kMyMGKa != ole.S_OK && kMyMGKa != S_FALSE {
			return nil, nil, cp3hZaJi3b
		}
	}

	xx_hM9Gbk, cp3hZaJi3b =  /*line :1*/ oleutil.XXajx9krfg("WbemScripting.SWbemLocator")
	if cp3hZaJi3b != nil {
		return nil, nil, cp3hZaJi3b
	} else if xx_hM9Gbk == nil {
		return nil, nil, W0_gOPjFGm
	}

	qnAGlFM6vP7, cp3hZaJi3b =  /*line :1*/ xx_hM9Gbk.QueryInterface(ole.Y0sP6F)
	if cp3hZaJi3b != nil {
		return nil, nil, cp3hZaJi3b
	}

	fAI_4V2jCrXK, cp3hZaJi3b =  /*line :1*/ oleutil.ZX4sNti(qnAGlFM6vP7, "ConnectServer", pvQQf8ap...)
	if cp3hZaJi3b != nil {
		return nil, nil, cp3hZaJi3b
	}

	return  /*line :1*/ fAI_4V2jCrXK.ToIDispatch(), fwMWfH, nil
}








func (agZTN4UVka *LlaCLems) CallMethod(pvQQf8ap []interface{}, gb_Ke0eQ, bCihZnk string, p4Tf0EE []interface{}) (int32, error) {
	hW5dsb9xJYgU, abtOE7O, cp3hZaJi3b :=  /*line :1*/ agZTN4UVka.uQUjdDex(pvQQf8ap...)
	if cp3hZaJi3b != nil {
		return 0,  /*line :1*/ fmt.Cnz_ab1BaZh("coinit: %v", cp3hZaJi3b)
	}
	defer  /*line :1*/ abtOE7O()

	ilRH65D, cp3hZaJi3b :=  /*line :1*/ oleutil.ZX4sNti(hW5dsb9xJYgU, "Get", gb_Ke0eQ)
	if cp3hZaJi3b != nil {
		return 0,  /*line :1*/ fmt.Cnz_ab1BaZh("CallMethod Get class %s: %v", gb_Ke0eQ, cp3hZaJi3b)
	}
	e6G73Op :=  /*line :1*/ ilRH65D.ToIDispatch()
	defer  /*line :1*/ ilRH65D.Clear()

	rjfzz3trVcp6, cp3hZaJi3b :=  /*line :1*/ oleutil.ZX4sNti(e6G73Op, bCihZnk, p4Tf0EE...)
	if cp3hZaJi3b != nil {
		return 0,  /*line :1*/ fmt.Cnz_ab1BaZh("CallMethod %s.%s: %v", gb_Ke0eQ, bCihZnk, cp3hZaJi3b)
	}
	qXcYqkeQO, kLajnjCGyp :=  /*line :1*/ rjfzz3trVcp6.Value().(int32)
	if !kLajnjCGyp {
		return 0,  /*line :1*/ fmt.Cnz_ab1BaZh("return value was not an int32: %v (%T)", rjfzz3trVcp6, rjfzz3trVcp6)
	}

	return qXcYqkeQO, nil
}












func (agZTN4UVka *LlaCLems) Query(cMO_44UJBCnD string, aSce7LEv65fL interface{}, pvQQf8ap ...interface{}) error {
	zC5ZpBO0 :=  /*line :1*/ reflect.SdHoaIQl(aSce7LEv65fL)
	if  /*line :1*/ zC5ZpBO0.Kind() != reflect.Ptr ||  /*line :1*/ zC5ZpBO0.IsNil() {
		return H1CZpWHzG
	}
	zC5ZpBO0 =  /*line :1*/ zC5ZpBO0.Elem()
	pu8xgRFuO, eNv5f4GjW :=  /*line :1*/ uYkMTE7iONoB(zC5ZpBO0)
	if pu8xgRFuO == multiArgTypeInvalid {
		return H1CZpWHzG
	}

	 /*line :1*/ anbyikS1Kv7i.Lock()
	defer  /*line :1*/ anbyikS1Kv7i.Unlock()
	 /*line :1*/ runtime.LockOSThread()
	defer  /*line :1*/ runtime.UnlockOSThread()

	hW5dsb9xJYgU, abtOE7O, cp3hZaJi3b :=  /*line :1*/ agZTN4UVka.uQUjdDex(pvQQf8ap...)
	if cp3hZaJi3b != nil {
		return cp3hZaJi3b
	}
	defer  /*line :1*/ abtOE7O()

	rjfzz3trVcp6, cp3hZaJi3b :=  /*line :1*/ oleutil.ZX4sNti(hW5dsb9xJYgU, "ExecQuery", cMO_44UJBCnD)
	if cp3hZaJi3b != nil {
		return cp3hZaJi3b
	}
	jLL4vFatPxst :=  /*line :1*/ rjfzz3trVcp6.ToIDispatch()
	defer  /*line :1*/ rjfzz3trVcp6.Clear()

	bgplXfRNa, cp3hZaJi3b :=  /*line :1*/ qltAor8(jLL4vFatPxst, "Count")
	if cp3hZaJi3b != nil {
		return cp3hZaJi3b
	}

	y8xQtJivJ8T, cp3hZaJi3b :=  /*line :1*/ jLL4vFatPxst.GetProperty("_NewEnum")
	if cp3hZaJi3b != nil {
		return cp3hZaJi3b
	}
	defer  /*line :1*/ y8xQtJivJ8T.Clear()

	ey5SnmcSRTF, cp3hZaJi3b :=  /*line :1*/ y8xQtJivJ8T.ToIUnknown().IEnumVARIANT(ole.Uv2Iuh4vReN)
	if cp3hZaJi3b != nil {
		return cp3hZaJi3b
	}
	if ey5SnmcSRTF == nil {
		return  /*line :1*/ fmt.Cnz_ab1BaZh("can't get IEnumVARIANT, enum is nil")
	}
	defer  /*line :1*/ ey5SnmcSRTF.Release()

	 /*line :1*/ zC5ZpBO0.Set( /*line :1*/ reflect.LydH53J9Si( /*line :1*/ zC5ZpBO0.Type(), 0,  /*line :1*/ int(bgplXfRNa)))

	var hXMC9Z886jX error
	for c0q80dXCLuTo, o7WozgxaXc, cp3hZaJi3b :=  /*line :1*/ ey5SnmcSRTF.Next(1); o7WozgxaXc > 0; c0q80dXCLuTo, o7WozgxaXc, cp3hZaJi3b =  /*line :1*/ ey5SnmcSRTF.Next(1) {
		if cp3hZaJi3b != nil {
			return cp3hZaJi3b
		}

		cp3hZaJi3b := func()  /*line :1*/ error {

			jy2GFGhmn2X :=  /*line :1*/ c0q80dXCLuTo.ToIDispatch()
			defer  /*line :1*/ jy2GFGhmn2X.Release()

			aMz7dx4SJ9qe :=  /*line :1*/ reflect.L6EAuajfYe1(eNv5f4GjW)
			if cp3hZaJi3b =  /*line :1*/ agZTN4UVka.pmm9KNNeic( /*line :1*/ aMz7dx4SJ9qe.Interface(), jy2GFGhmn2X); cp3hZaJi3b != nil {
				if _, kLajnjCGyp := cp3hZaJi3b.(*OyXNb8MTjY); kLajnjCGyp {

					hXMC9Z886jX = cp3hZaJi3b
				} else {
					return cp3hZaJi3b
				}
			}
			if pu8xgRFuO != multiArgTypeStructPtr {
				aMz7dx4SJ9qe =  /*line :1*/ aMz7dx4SJ9qe.Elem()
			}
			 /*line :1*/ zC5ZpBO0.Set( /*line :1*/ reflect.JrTptcQhpl0_(zC5ZpBO0, aMz7dx4SJ9qe))
			return nil
		}()
		if cp3hZaJi3b != nil {
			return cp3hZaJi3b
		}
	}
	return hXMC9Z886jX
}





type OyXNb8MTjY struct {
	Xsa7nTt1	reflect.YJh989LTZsVX
	BCxGq3fk_oGq	string
	WHTIIzUX9T35	string
}

func (muTJ5Av *OyXNb8MTjY) Error() string {
	return  /*line :1*/ fmt.IBsS3Oc("wmi: cannot load field %q into a %q: %s",
		muTJ5Av.BCxGq3fk_oGq, muTJ5Av.Xsa7nTt1, muTJ5Av.WHTIIzUX9T35)
}

var sv4IbPnYV =  /*line :1*/ reflect.Cher1a2Fblr(time.G4KDkI{})


func (agZTN4UVka *LlaCLems) pmm9KNNeic(aSce7LEv65fL interface{}, ryfADc2T *ole.IDispatch) (hXMC9Z886jX error) {
	fhrpGIQPP :=  /*line :1*/ reflect.SdHoaIQl(aSce7LEv65fL).Elem()
	for aJSPCt := 0; aJSPCt <  /*line :1*/ fhrpGIQPP.NumField(); aJSPCt++ {
		nbo1CgJ0Vi :=  /*line :1*/ fhrpGIQPP.Field(aJSPCt)
		ljPuTxda5j := nbo1CgJ0Vi
		oKlmyTOv :=  /*line :1*/ nbo1CgJ0Vi.Kind() == reflect.Ptr
		if oKlmyTOv {
			sMMDHK :=  /*line :1*/ reflect.L6EAuajfYe1( /*line :1*/ nbo1CgJ0Vi.Type().Elem())
			 /*line :1*/ nbo1CgJ0Vi.Set(sMMDHK)
			nbo1CgJ0Vi =  /*line :1*/ nbo1CgJ0Vi.Elem()
		}
		ndROa3Gz1B :=  /*line :1*/ fhrpGIQPP.Type().Field(aJSPCt).ZOSBHkJz
		if ndROa3Gz1B[0] < 'A' || ndROa3Gz1B[0] > 'Z' {
			continue
		}
		if ! /*line :1*/ nbo1CgJ0Vi.CanSet() {
			return &OyXNb8MTjY{
				Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
				BCxGq3fk_oGq:	ndROa3Gz1B,
				WHTIIzUX9T35:	"CanSet() is false",
			}
		}
		hCByzP_aW, cp3hZaJi3b :=  /*line :1*/ oleutil.Hmx6gDksa(ryfADc2T, ndROa3Gz1B)
		if cp3hZaJi3b != nil {
			if !agZTN4UVka.YevekfvvbB1S {
				hXMC9Z886jX = &OyXNb8MTjY{
					Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
					BCxGq3fk_oGq:	ndROa3Gz1B,
					WHTIIzUX9T35:	"no such struct field",
				}
			}
			continue
		}
		defer  /*line :1*/ hCByzP_aW.Clear()

		if hCByzP_aW.SUbB7_4 == 0x1 {
			continue
		}

		switch bCzLF6c :=  /*line :1*/ hCByzP_aW.Value().(type) {
		case int8, int16, int32, int64, int:
			fhrpGIQPP :=  /*line :1*/ reflect.SdHoaIQl(bCzLF6c).Int()
			switch  /*line :1*/ nbo1CgJ0Vi.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				 /*line :1*/ nbo1CgJ0Vi.SetInt(fhrpGIQPP)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				 /*line :1*/ nbo1CgJ0Vi.SetUint( /*line :1*/ uint64(fhrpGIQPP))
			default:
				return &OyXNb8MTjY{
					Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
					BCxGq3fk_oGq:	ndROa3Gz1B,
					WHTIIzUX9T35:	"not an integer class",
				}
			}
		case uint8, uint16, uint32, uint64:
			fhrpGIQPP :=  /*line :1*/ reflect.SdHoaIQl(bCzLF6c).Uint()
			switch  /*line :1*/ nbo1CgJ0Vi.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				 /*line :1*/ nbo1CgJ0Vi.SetInt( /*line :1*/ int64(fhrpGIQPP))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				 /*line :1*/ nbo1CgJ0Vi.SetUint(fhrpGIQPP)
			default:
				return &OyXNb8MTjY{
					Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
					BCxGq3fk_oGq:	ndROa3Gz1B,
					WHTIIzUX9T35:	"not an integer class",
				}
			}
		case string:
			switch  /*line :1*/ nbo1CgJ0Vi.Kind() {
			case reflect.String:
				 /*line :1*/ nbo1CgJ0Vi.SetString(bCzLF6c)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				hOeE0Lf42wk, cp3hZaJi3b :=  /*line :1*/ strconv.N8Xpap1Vt1ot(bCzLF6c, 10, 64)
				if cp3hZaJi3b != nil {
					return cp3hZaJi3b
				}
				 /*line :1*/ nbo1CgJ0Vi.SetInt(hOeE0Lf42wk)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				pFmXOpscU1A, cp3hZaJi3b :=  /*line :1*/ strconv.WOxcnaOAzTeq(bCzLF6c, 10, 64)
				if cp3hZaJi3b != nil {
					return cp3hZaJi3b
				}
				 /*line :1*/ nbo1CgJ0Vi.SetUint(pFmXOpscU1A)
			case reflect.Struct:
				switch  /*line :1*/ nbo1CgJ0Vi.Type() {
				case sv4IbPnYV:
					if  /*line :1*/ len(bCzLF6c) == 25 {
						a_mzdEYz8a45, cp3hZaJi3b :=  /*line :1*/ strconv.GmbOy60GCccC(bCzLF6c[22:])
						if cp3hZaJi3b != nil {
							return cp3hZaJi3b
						}
						bCzLF6c = bCzLF6c[:22] +  /*line :1*/ fmt.IBsS3Oc("%02d%02d", a_mzdEYz8a45/60, a_mzdEYz8a45%60)
					}
					b6h_pivWAnr, cp3hZaJi3b :=  /*line :1*/ time.KoKD_YnCna("20060102150405.000000-0700", bCzLF6c)
					if cp3hZaJi3b != nil {
						return cp3hZaJi3b
					}
					 /*line :1*/ nbo1CgJ0Vi.Set( /*line :1*/ reflect.SdHoaIQl(b6h_pivWAnr))
				}
			}
		case bool:
			switch  /*line :1*/ nbo1CgJ0Vi.Kind() {
			case reflect.Bool:
				 /*line :1*/ nbo1CgJ0Vi.SetBool(bCzLF6c)
			default:
				return &OyXNb8MTjY{
					Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
					BCxGq3fk_oGq:	ndROa3Gz1B,
					WHTIIzUX9T35:	"not a bool",
				}
			}
		case float32:
			switch  /*line :1*/ nbo1CgJ0Vi.Kind() {
			case reflect.Float32:
				 /*line :1*/ nbo1CgJ0Vi.SetFloat( /*line :1*/ float64(bCzLF6c))
			default:
				return &OyXNb8MTjY{
					Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
					BCxGq3fk_oGq:	ndROa3Gz1B,
					WHTIIzUX9T35:	"not a Float32",
				}
			}
		default:
			if  /*line :1*/ nbo1CgJ0Vi.Kind() == reflect.Slice {
				switch  /*line :1*/ nbo1CgJ0Vi.Type().Elem().Kind() {
				case reflect.String:
					i6eqoE :=  /*line :1*/ hCByzP_aW.ToArray()
					if i6eqoE != nil {
						nXZ9XbCUk :=  /*line :1*/ i6eqoE.ToValueArray()
						z_TKM1KT :=  /*line :1*/ reflect.LydH53J9Si( /*line :1*/ nbo1CgJ0Vi.Type(),  /*line :1*/ len(nXZ9XbCUk),  /*line :1*/ len(nXZ9XbCUk))
						for aJSPCt, fhrpGIQPP := range nXZ9XbCUk {
							epb5ydq5 :=  /*line :1*/ z_TKM1KT.Index(aJSPCt)
							 /*line :1*/ epb5ydq5.SetString(fhrpGIQPP.(string))
						}
						 /*line :1*/ nbo1CgJ0Vi.Set(z_TKM1KT)
					}
				case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
					i6eqoE :=  /*line :1*/ hCByzP_aW.ToArray()
					if i6eqoE != nil {
						nXZ9XbCUk :=  /*line :1*/ i6eqoE.ToValueArray()
						z_TKM1KT :=  /*line :1*/ reflect.LydH53J9Si( /*line :1*/ nbo1CgJ0Vi.Type(),  /*line :1*/ len(nXZ9XbCUk),  /*line :1*/ len(nXZ9XbCUk))
						for aJSPCt, fhrpGIQPP := range nXZ9XbCUk {
							epb5ydq5 :=  /*line :1*/ z_TKM1KT.Index(aJSPCt)
							 /*line :1*/ epb5ydq5.SetUint( /*line :1*/ reflect.SdHoaIQl(fhrpGIQPP).Uint())
						}
						 /*line :1*/ nbo1CgJ0Vi.Set(z_TKM1KT)
					}
				case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
					i6eqoE :=  /*line :1*/ hCByzP_aW.ToArray()
					if i6eqoE != nil {
						nXZ9XbCUk :=  /*line :1*/ i6eqoE.ToValueArray()
						z_TKM1KT :=  /*line :1*/ reflect.LydH53J9Si( /*line :1*/ nbo1CgJ0Vi.Type(),  /*line :1*/ len(nXZ9XbCUk),  /*line :1*/ len(nXZ9XbCUk))
						for aJSPCt, fhrpGIQPP := range nXZ9XbCUk {
							epb5ydq5 :=  /*line :1*/ z_TKM1KT.Index(aJSPCt)
							 /*line :1*/ epb5ydq5.SetInt( /*line :1*/ reflect.SdHoaIQl(fhrpGIQPP).Int())
						}
						 /*line :1*/ nbo1CgJ0Vi.Set(z_TKM1KT)
					}
				default:
					return &OyXNb8MTjY{
						Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
						BCxGq3fk_oGq:	ndROa3Gz1B,
						WHTIIzUX9T35:	 /*line :1*/ fmt.IBsS3Oc("unsupported slice type (%T)", bCzLF6c),
					}
				}
			} else {
				zQjblomtSpB :=  /*line :1*/ reflect.Cher1a2Fblr(bCzLF6c)
				if zQjblomtSpB == nil && (oKlmyTOv || agZTN4UVka.TYHmyvQ0) {
					if (oKlmyTOv && agZTN4UVka.CQ8rI1v4m) || (!oKlmyTOv && agZTN4UVka.TYHmyvQ0) {
						 /*line :1*/ ljPuTxda5j.Set( /*line :1*/ reflect.H_B9rU1ADy( /*line :1*/ ljPuTxda5j.Type()))
					}
					break
				}
				return &OyXNb8MTjY{
					Xsa7nTt1:	 /*line :1*/ ljPuTxda5j.Type(),
					BCxGq3fk_oGq:	ndROa3Gz1B,
					WHTIIzUX9T35:	 /*line :1*/ fmt.IBsS3Oc("unsupported type (%T)", bCzLF6c),
				}
			}
		}
	}
	return hXMC9Z886jX
}

type y27brVVKJ int

const (
	multiArgTypeInvalid	y27brVVKJ	= iota
	multiArgTypeStruct
	multiArgTypeStructPtr
)





func uYkMTE7iONoB(fhrpGIQPP reflect.Value) (rYAZU5mbIY y27brVVKJ, eNv5f4GjW reflect.YJh989LTZsVX) {
	if  /*line :1*/ fhrpGIQPP.Kind() != reflect.Slice {
		return multiArgTypeInvalid, nil
	}
	eNv5f4GjW =  /*line :1*/ fhrpGIQPP.Type().Elem()
	switch  /*line :1*/ eNv5f4GjW.Kind() {
	case reflect.Struct:
		return multiArgTypeStruct, eNv5f4GjW
	case reflect.Ptr:
		eNv5f4GjW =  /*line :1*/ eNv5f4GjW.Elem()
		if  /*line :1*/ eNv5f4GjW.Kind() == reflect.Struct {
			return multiArgTypeStructPtr, eNv5f4GjW
		}
	}
	return multiArgTypeInvalid, nil
}

func qltAor8(jy2GFGhmn2X *ole.IDispatch, hCByzP_aW string) (int64, error) {
	fhrpGIQPP, cp3hZaJi3b :=  /*line :1*/ oleutil.Hmx6gDksa(jy2GFGhmn2X, hCByzP_aW)
	if cp3hZaJi3b != nil {
		return 0, cp3hZaJi3b
	}
	defer  /*line :1*/ fhrpGIQPP.Clear()

	aJSPCt :=  /*line :1*/ int64(fhrpGIQPP.IG3mNr)
	return aJSPCt, nil
}







func H5IaZ06(ryfADc2T interface{}, d0f5xTs9 string, e6G73Op ...string) string {
	var ewHTg4l_TO bytes.SuzhTyj35Jl
	 /*line :1*/ ewHTg4l_TO.WriteString("SELECT ")
	epb5ydq5 :=  /*line :1*/ reflect.XF6ikODQ8h( /*line :1*/ reflect.SdHoaIQl(ryfADc2T))
	b6h_pivWAnr :=  /*line :1*/ epb5ydq5.Type()
	if  /*line :1*/ epb5ydq5.Kind() == reflect.Slice {
		b6h_pivWAnr =  /*line :1*/ b6h_pivWAnr.Elem()
	}
	if  /*line :1*/ b6h_pivWAnr.Kind() != reflect.Struct {
		return ""
	}
	var tDGCR8yDwcGB []string
	for aJSPCt := 0; aJSPCt <  /*line :1*/ b6h_pivWAnr.NumField(); aJSPCt++ {
		tDGCR8yDwcGB =  /*line :1*/ append(tDGCR8yDwcGB,  /*line :1*/ b6h_pivWAnr.Field(aJSPCt).ZOSBHkJz)
	}
	 /*line :1*/ ewHTg4l_TO.WriteString( /*line :1*/ strings.ZTisTA(tDGCR8yDwcGB, ", "))
	 /*line :1*/ ewHTg4l_TO.WriteString(" FROM ")
	if  /*line :1*/ len(e6G73Op) > 0 {
		 /*line :1*/ ewHTg4l_TO.WriteString(e6G73Op[0])
	} else {
		 /*line :1*/ ewHTg4l_TO.WriteString( /*line :1*/ b6h_pivWAnr.Name())
	}
	 /*line :1*/ ewHTg4l_TO.WriteString(" " + d0f5xTs9)
	return  /*line :1*/ ewHTg4l_TO.String()
}
