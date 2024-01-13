//line :1
package baLhtK0

import (
	fmt "kFsoCfy5zWG"
	reflect "reflect"
	"runtime"
	sync "sync"

	ole "fuA83L"
	oleutil "uBCMx3OoU"
)


type NcjbAXk6 struct {
	
	bk_galralVN	*LlaCLems	
	x4jfHJ	*ole.IUnknown
	hDwQD2	*ole.IDispatch
	kWa3k3SIt	chan *queryRequest
	iAEQfc4Xv	chan error
	xhvdF24gwO	sync.DIRWe11KYlYa
}

type queryRequest struct {
	query	string
	dst	interface{}
	args	[]interface{}
	finished	chan error
}


func XvqfK8xO7(agZTN4UVka *LlaCLems, pvQQf8ap ...interface{}) (*NcjbAXk6, error) {

	epb5ydq5 :=  /*line :1*/ new(NcjbAXk6)
	epb5ydq5.bk_galralVN = agZTN4UVka
	epb5ydq5.kWa3k3SIt =  /*line :1*/ make(chan *queryRequest)
	k5eSJg :=  /*line :1*/ make(chan error)
	go  /*line :1*/ epb5ydq5.dY2wS3H(k5eSJg)

	cp3hZaJi3b, kLajnjCGyp := <-k5eSJg
	if kLajnjCGyp {
		return nil, cp3hZaJi3b
	}

	return epb5ydq5, nil
}


func (epb5ydq5 *NcjbAXk6) Close() error {
	 /*line :1*/ epb5ydq5.xhvdF24gwO.Lock()
	if epb5ydq5 == nil || epb5ydq5.hDwQD2 == nil {
		 /*line :1*/ epb5ydq5.xhvdF24gwO.Unlock()
		return  /*line :1*/ fmt.Cnz_ab1BaZh("SWbemServices is not Initialized")
	}
	if epb5ydq5.kWa3k3SIt == nil {
		 /*line :1*/ epb5ydq5.xhvdF24gwO.Unlock()
		return  /*line :1*/ fmt.Cnz_ab1BaZh("SWbemServices has been closed")
	}
	
	var jLL4vFatPxst error
	s62M42U :=  /*line :1*/ make(chan error)
	epb5ydq5.iAEQfc4Xv = s62M42U
	 /*line :1*/ close(epb5ydq5.kWa3k3SIt)
	 /*line :1*/ epb5ydq5.xhvdF24gwO.Unlock()
	cp3hZaJi3b, kLajnjCGyp := <-s62M42U
	if kLajnjCGyp {
		jLL4vFatPxst = cp3hZaJi3b
	}

	return jLL4vFatPxst
}

func (epb5ydq5 *NcjbAXk6) dY2wS3H(k5eSJg chan error) {

	 /*line :1*/ runtime.LockOSThread()
	defer  /*line :1*/ runtime.UnlockOSThread()

	cp3hZaJi3b :=  /*line :1*/ ole.KKzYzCZuN(0, ole.COINIT_MULTITHREADED)
	if cp3hZaJi3b != nil {
		kMyMGKa :=  /*line :1*/ cp3hZaJi3b.(*ole.DeOFa72M).Code()
		if kMyMGKa != ole.S_OK && kMyMGKa != S_FALSE {
			k5eSJg <-  /*line :1*/ fmt.Cnz_ab1BaZh("ole.CoInitializeEx error: %v", cp3hZaJi3b)
			return
		}
	}
	defer  /*line :1*/ ole.DvOlryZ5()

	xx_hM9Gbk, cp3hZaJi3b :=  /*line :1*/ oleutil.XXajx9krfg("WbemScripting.SWbemLocator")
	if cp3hZaJi3b != nil {
		k5eSJg <-  /*line :1*/ fmt.Cnz_ab1BaZh("CreateObject SWbemLocator error: %v", cp3hZaJi3b)
		return
	} else if xx_hM9Gbk == nil {
		k5eSJg <- W0_gOPjFGm
		return
	}
	defer  /*line :1*/ xx_hM9Gbk.Release()
	epb5ydq5.x4jfHJ = xx_hM9Gbk

	ci7bj8Kv, cp3hZaJi3b :=  /*line :1*/ epb5ydq5.x4jfHJ.QueryInterface(ole.Y0sP6F)
	if cp3hZaJi3b != nil {
		k5eSJg <-  /*line :1*/ fmt.Cnz_ab1BaZh("SWbemLocator QueryInterface error: %v", cp3hZaJi3b)
		return
	}
	defer  /*line :1*/ ci7bj8Kv.Release()
	epb5ydq5.hDwQD2 = ci7bj8Kv

	 /*line :1*/ close(k5eSJg)

	for jviAmguXn6Si := range epb5ydq5.kWa3k3SIt {

		fxHal91 :=  /*line :1*/ epb5ydq5.jHct0aiEWrV(jviAmguXn6Si)

		if fxHal91 != nil {
			jviAmguXn6Si.finished <- fxHal91
		}
		 /*line :1*/ close(jviAmguXn6Si.finished)
	}

	epb5ydq5.kWa3k3SIt = nil

	 /*line :1*/ close(epb5ydq5.iAEQfc4Xv)
}











func (epb5ydq5 *NcjbAXk6) Query(cMO_44UJBCnD string, aSce7LEv65fL interface{}, pvQQf8ap ...interface{}) error {
	 /*line :1*/ epb5ydq5.xhvdF24gwO.Lock()
	if epb5ydq5 == nil || epb5ydq5.hDwQD2 == nil {
		 /*line :1*/ epb5ydq5.xhvdF24gwO.Unlock()
		return  /*line :1*/ fmt.Cnz_ab1BaZh("SWbemServices is not Initialized")
	}
	if epb5ydq5.kWa3k3SIt == nil {
		 /*line :1*/ epb5ydq5.xhvdF24gwO.Unlock()
		return  /*line :1*/ fmt.Cnz_ab1BaZh("SWbemServices has been closed")
	}

	rkYfG9aHhUum := queryRequest{
		query:	cMO_44UJBCnD,
		dst:	aSce7LEv65fL,
		args:	pvQQf8ap,
		finished:	 /*line :1*/ make(chan error),
	}
	epb5ydq5.kWa3k3SIt <- &rkYfG9aHhUum
	 /*line :1*/ epb5ydq5.xhvdF24gwO.Unlock()
	cp3hZaJi3b, kLajnjCGyp := <-rkYfG9aHhUum.finished
	if kLajnjCGyp {

		return cp3hZaJi3b
	}

	return nil
}

func (epb5ydq5 *NcjbAXk6) jHct0aiEWrV(jviAmguXn6Si *queryRequest) error {
	if epb5ydq5 == nil || epb5ydq5.hDwQD2 == nil {
		return  /*line :1*/ fmt.Cnz_ab1BaZh("SWbemServices is not Initialized")
	}
	qnAGlFM6vP7 := epb5ydq5.hDwQD2

	zC5ZpBO0 :=  /*line :1*/ reflect.SdHoaIQl(jviAmguXn6Si.dst)
	if  /*line :1*/ zC5ZpBO0.Kind() != reflect.Ptr ||  /*line :1*/ zC5ZpBO0.IsNil() {
		return H1CZpWHzG
	}
	zC5ZpBO0 =  /*line :1*/ zC5ZpBO0.Elem()
	pu8xgRFuO, eNv5f4GjW :=  /*line :1*/ uYkMTE7iONoB(zC5ZpBO0)
	if pu8xgRFuO == multiArgTypeInvalid {
		return H1CZpWHzG
	}

	fAI_4V2jCrXK, cp3hZaJi3b :=  /*line :1*/ oleutil.ZX4sNti(qnAGlFM6vP7, "ConnectServer", jviAmguXn6Si.args...)
	if cp3hZaJi3b != nil {
		return cp3hZaJi3b
	}
	hW5dsb9xJYgU :=  /*line :1*/ fAI_4V2jCrXK.ToIDispatch()
	defer  /*line :1*/ fAI_4V2jCrXK.Clear()

	rjfzz3trVcp6, cp3hZaJi3b :=  /*line :1*/ oleutil.ZX4sNti(hW5dsb9xJYgU, "ExecQuery", jviAmguXn6Si.query)
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
			if cp3hZaJi3b =  /*line :1*/ epb5ydq5.bk_galralVN.pmm9KNNeic( /*line :1*/ aMz7dx4SJ9qe.Interface(), jy2GFGhmn2X); cp3hZaJi3b != nil {
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
