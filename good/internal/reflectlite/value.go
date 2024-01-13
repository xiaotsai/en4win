//line :1
package yAj8QghzkR

import (
	"internal/abi"
	"internal/goarch"
	unsafeheader "EHdBEpZ7y6"
	"runtime"
	"unsafe"
)

type ZPN6dR struct {
	dDlYMUKHt2j	*abi.Type

	k0FQfYmXZ	unsafe.Pointer

	fXMYmkc
}

type fXMYmkc uintptr

const (
	flagKindWidth		= 5
	flagKindMask	fXMYmkc	= 1<<flagKindWidth - 1
	flagStickyRO	fXMYmkc	= 1 << 5
	flagEmbedRO	fXMYmkc	= 1 << 6
	flagIndir	fXMYmkc	= 1 << 7
	flagAddr	fXMYmkc	= 1 << 8
	flagMethod	fXMYmkc	= 1 << 9
	flagMethodShift		= 10
	flagRO	fXMYmkc	= flagStickyRO | flagEmbedRO
)

func (xffw4nfU21SZ fXMYmkc) xm3h_2oMKuTJ() FZ0rl1z {
	return  /*line :1*/ FZ0rl1z(xffw4nfU21SZ & flagKindMask)
}

func (xffw4nfU21SZ fXMYmkc) iisZR1yJd() fXMYmkc {
	if xffw4nfU21SZ&flagRO != 0 {
		return flagStickyRO
	}
	return 0
}

func (d583PKWfr ZPN6dR) jAnZIM10jY7K() unsafe.Pointer {
	if  /*line :1*/ d583PKWfr.dDlYMUKHt2j.Size() != goarch.PtrSize || ! /*line :1*/ d583PKWfr.dDlYMUKHt2j.Pointers() {
		 /*line :1*/ panic("can't call pointer on a non-pointer Value")
	}
	if d583PKWfr.fXMYmkc&flagIndir != 0 {
		return *(* /*line :1*/ unsafe.Pointer)(d583PKWfr.k0FQfYmXZ)
	}
	return d583PKWfr.k0FQfYmXZ
}

func r9b8DkSQ(d583PKWfr ZPN6dR) any {
	gc1v5B2w4C := d583PKWfr.dDlYMUKHt2j
	var dujSIeNg10V any
	rdxsdJg3 := (* /*line :1*/ pgNclazaiBJc)( /*line :1*/ unsafe.Pointer(&dujSIeNg10V))

	switch {
	case  /*line :1*/ vJEQ9cGoTSlB(gc1v5B2w4C):
		if d583PKWfr.fXMYmkc&flagIndir == 0 {
			 /*line :1*/ panic("bad indir")
		}

		dDWB6QciD := d583PKWfr.k0FQfYmXZ
		if d583PKWfr.fXMYmkc&flagAddr != 0 {

			rVHlxpM :=  /*line :1*/ cUo1caLrG(gc1v5B2w4C)
			 /*line :1*/ rFMmsMP6H(gc1v5B2w4C, rVHlxpM, dDWB6QciD)
			dDWB6QciD = rVHlxpM
		}
		rdxsdJg3.kMGZ7zYC = dDWB6QciD
	case d583PKWfr.fXMYmkc&flagIndir != 0:

		rdxsdJg3.kMGZ7zYC = *(* /*line :1*/ unsafe.Pointer)(d583PKWfr.k0FQfYmXZ)
	default:

		rdxsdJg3.kMGZ7zYC = d583PKWfr.k0FQfYmXZ
	}

	rdxsdJg3.ubPJaJymi = gc1v5B2w4C
	return dujSIeNg10V
}

func pFCjGZcYhZ(dujSIeNg10V any) ZPN6dR {
	rdxsdJg3 := (* /*line :1*/ pgNclazaiBJc)( /*line :1*/ unsafe.Pointer(&dujSIeNg10V))

	gc1v5B2w4C := rdxsdJg3.ubPJaJymi
	if gc1v5B2w4C == nil {
		return ZPN6dR{}
	}
	xffw4nfU21SZ :=  /*line :1*/ fXMYmkc( /*line :1*/ gc1v5B2w4C.Kind())
	if  /*line :1*/ vJEQ9cGoTSlB(gc1v5B2w4C) {
		xffw4nfU21SZ |= flagIndir
	}
	return ZPN6dR{gc1v5B2w4C, rdxsdJg3.kMGZ7zYC, xffw4nfU21SZ}
}

type RfP16baj struct {
	BH8_yBNG3HjN	string
	Aq5nH3jY	FZ0rl1z
}

func (rdxsdJg3 *RfP16baj) Error() string {
	if rdxsdJg3.Aq5nH3jY == 0 {
		return "reflect: call of " + rdxsdJg3.BH8_yBNG3HjN + " on zero Value"
	}
	return "reflect: call of " + rdxsdJg3.BH8_yBNG3HjN + " on " +  /*line :1*/ rdxsdJg3.Aq5nH3jY.String() + " Value"
}

func gXdenBWaPe8O() string {
	aI2g5FX_C, _, _, _ :=  /*line :1*/ runtime.Caller(2)
	xffw4nfU21SZ :=  /*line :1*/ runtime.FuncForPC(aI2g5FX_C)
	if xffw4nfU21SZ == nil {
		return "unknown method"
	}
	return  /*line :1*/ xffw4nfU21SZ.Name()
}

type pgNclazaiBJc struct {
	ubPJaJymi	*abi.Type
	kMGZ7zYC	unsafe.Pointer
}

func (xffw4nfU21SZ fXMYmkc) lE8afngtW4() {
	if xffw4nfU21SZ == 0 {
		 /*line :1*/ panic(&RfP16baj{ /*line :1*/ gXdenBWaPe8O(), 0})
	}
	if xffw4nfU21SZ&flagRO != 0 {
		 /*line :1*/ panic("reflect: " +  /*line :1*/ gXdenBWaPe8O() + " using value obtained using unexported field")
	}
}

func (xffw4nfU21SZ fXMYmkc) lo0VaGi5BB_() {
	if xffw4nfU21SZ == 0 {
		 /*line :1*/ panic(&RfP16baj{ /*line :1*/ gXdenBWaPe8O(), abi.Invalid})
	}

	if xffw4nfU21SZ&flagRO != 0 {
		 /*line :1*/ panic("reflect: " +  /*line :1*/ gXdenBWaPe8O() + " using value obtained using unexported field")
	}
	if xffw4nfU21SZ&flagAddr == 0 {
		 /*line :1*/ panic("reflect: " +  /*line :1*/ gXdenBWaPe8O() + " using unaddressable value")
	}
}

func (d583PKWfr ZPN6dR) CanSet() bool {
	return d583PKWfr.fXMYmkc&(flagAddr|flagRO) == flagAddr
}

func (d583PKWfr ZPN6dR) Elem() ZPN6dR {
	wjaT5t :=  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ()
	switch wjaT5t {
	case abi.Interface:
		var mhQwtWWVa any
		if  /*line :1*/ d583PKWfr.dDlYMUKHt2j.NumMethod() == 0 {
			mhQwtWWVa = *(* /*line :1*/ any)(d583PKWfr.k0FQfYmXZ)
		} else {
			mhQwtWWVa = ( /*line :1*/ any)(*(*interface {
				 /*line :1*/ M()
			})(d583PKWfr.k0FQfYmXZ))
		}
		jJCfpek :=  /*line :1*/ pFCjGZcYhZ(mhQwtWWVa)
		if jJCfpek.fXMYmkc != 0 {
			jJCfpek.fXMYmkc |=  /*line :1*/ d583PKWfr.fXMYmkc.iisZR1yJd()
		}
		return jJCfpek
	case abi.Pointer:
		dDWB6QciD := d583PKWfr.k0FQfYmXZ
		if d583PKWfr.fXMYmkc&flagIndir != 0 {
			dDWB6QciD = *(* /*line :1*/ unsafe.Pointer)(dDWB6QciD)
		}

		if dDWB6QciD == nil {
			return ZPN6dR{}
		}
		tq_UCmMaf0QU := (* /*line :1*/ naf2N4qWKL9)( /*line :1*/ unsafe.Pointer(d583PKWfr.dDlYMUKHt2j))
		p_IC4wLPVMmq := tq_UCmMaf0QU.Elem
		by9tFKy2 := d583PKWfr.fXMYmkc&flagRO | flagIndir | flagAddr
		by9tFKy2 |=  /*line :1*/ fXMYmkc( /*line :1*/ p_IC4wLPVMmq.Kind())
		return ZPN6dR{p_IC4wLPVMmq, dDWB6QciD, by9tFKy2}
	}
	 /*line :1*/ panic(&RfP16baj{"reflectlite.Value.Elem",  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ()})
}

func m4SpTD7Rab(d583PKWfr ZPN6dR) any {
	if d583PKWfr.fXMYmkc == 0 {
		 /*line :1*/ panic(&RfP16baj{"reflectlite.Value.Interface", 0})
	}

	if  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ() == abi.Interface {

		if  /*line :1*/ d583PKWfr.jwihyQj92Ka() == 0 {
			return *(* /*line :1*/ any)(d583PKWfr.k0FQfYmXZ)
		}
		return *(*interface {
			 /*line :1*/ M()
		})(d583PKWfr.k0FQfYmXZ)
	}

	return  /*line :1*/ r9b8DkSQ(d583PKWfr)
}

func (d583PKWfr ZPN6dR) IsNil() bool {
	wjaT5t :=  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ()
	switch wjaT5t {
	case abi.Chan, abi.Func, abi.Map, abi.Pointer, abi.UnsafePointer:

		dDWB6QciD := d583PKWfr.k0FQfYmXZ
		if d583PKWfr.fXMYmkc&flagIndir != 0 {
			dDWB6QciD = *(* /*line :1*/ unsafe.Pointer)(dDWB6QciD)
		}
		return dDWB6QciD == nil
	case abi.Interface, abi.Slice:

		return *(* /*line :1*/ unsafe.Pointer)(d583PKWfr.k0FQfYmXZ) == nil
	}
	 /*line :1*/ panic(&RfP16baj{"reflectlite.Value.IsNil",  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ()})
}

func (d583PKWfr ZPN6dR) IsValid() bool {
	return d583PKWfr.fXMYmkc != 0
}

func (d583PKWfr ZPN6dR) Kind() FZ0rl1z {
	return  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ()
}

func yScq9hv5Zxna(unsafe.Pointer) int
func h5VXr0f1C(unsafe.Pointer) int

func (d583PKWfr ZPN6dR) Len() int {
	wjaT5t :=  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ()
	switch wjaT5t {
	case abi.Array:
		tq_UCmMaf0QU := (* /*line :1*/ jCeH0iHPi9)( /*line :1*/ unsafe.Pointer(d583PKWfr.dDlYMUKHt2j))
		return  /*line :1*/ int(tq_UCmMaf0QU.Len)
	case abi.Chan:
		return  /*line :1*/ yScq9hv5Zxna( /*line :1*/ d583PKWfr.jAnZIM10jY7K())
	case abi.Map:
		return  /*line :1*/ h5VXr0f1C( /*line :1*/ d583PKWfr.jAnZIM10jY7K())
	case abi.Slice:

		return (* /*line :1*/ unsafeheader.AKoOflOZ)(d583PKWfr.k0FQfYmXZ).AJ2N5B
	case abi.String:

		return (* /*line :1*/ unsafeheader.HrYeZlWHaf)(d583PKWfr.k0FQfYmXZ).SkAFPemec
	}
	 /*line :1*/ panic(&RfP16baj{"reflect.Value.Len",  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ()})
}

func (d583PKWfr ZPN6dR) jwihyQj92Ka() int {
	if d583PKWfr.dDlYMUKHt2j == nil {
		 /*line :1*/ panic(&RfP16baj{"reflectlite.Value.NumMethod", abi.Invalid})
	}
	return  /*line :1*/ d583PKWfr.dDlYMUKHt2j.NumMethod()
}

func (d583PKWfr ZPN6dR) Set(jJCfpek ZPN6dR) {
	 /*line :1*/ d583PKWfr.lo0VaGi5BB_()
	 /*line :1*/ jJCfpek.lE8afngtW4()
	var vhuI3TM unsafe.Pointer
	if  /*line :1*/ d583PKWfr.xm3h_2oMKuTJ() == abi.Interface {
		vhuI3TM = d583PKWfr.k0FQfYmXZ
	}
	jJCfpek =  /*line :1*/ jJCfpek.lRS1cks("reflectlite.Set", d583PKWfr.dDlYMUKHt2j, vhuI3TM)
	if jJCfpek.fXMYmkc&flagIndir != 0 {
		 /*line :1*/ rFMmsMP6H(d583PKWfr.dDlYMUKHt2j, d583PKWfr.k0FQfYmXZ, jJCfpek.k0FQfYmXZ)
	} else {
		*(* /*line :1*/ unsafe.Pointer)(d583PKWfr.k0FQfYmXZ) = jJCfpek.k0FQfYmXZ
	}
}

func (d583PKWfr ZPN6dR) Type() LrcFby {
	xffw4nfU21SZ := d583PKWfr.fXMYmkc
	if xffw4nfU21SZ == 0 {
		 /*line :1*/ panic(&RfP16baj{"reflectlite.Value.Type", abi.Invalid})
	}

	return  /*line :1*/ qK4nvEMY(d583PKWfr.dDlYMUKHt2j)
}

func cUo1caLrG(*abi.Type) unsafe.Pointer

func Au73KsHV1lt(dujSIeNg10V any) ZPN6dR {
	if dujSIeNg10V == nil {
		return ZPN6dR{}
	}

	 /*line :1*/ sQL9ue1Cs0x(dujSIeNg10V)

	return  /*line :1*/ pFCjGZcYhZ(dujSIeNg10V)
}

func (d583PKWfr ZPN6dR) lRS1cks(rad5bjl6bHs string, vu1nEO26a82W *abi.Type, vhuI3TM unsafe.Pointer) ZPN6dR {

	switch {
	case  /*line :1*/ aUlsvhDR(vu1nEO26a82W, d583PKWfr.dDlYMUKHt2j):

		by9tFKy2 := d583PKWfr.fXMYmkc&(flagAddr|flagIndir) |  /*line :1*/ d583PKWfr.fXMYmkc.iisZR1yJd()
		by9tFKy2 |=  /*line :1*/ fXMYmkc( /*line :1*/ vu1nEO26a82W.Kind())
		return ZPN6dR{vu1nEO26a82W, d583PKWfr.k0FQfYmXZ, by9tFKy2}

	case  /*line :1*/ t5oBGiG1gBz_(vu1nEO26a82W, d583PKWfr.dDlYMUKHt2j):
		if vhuI3TM == nil {
			vhuI3TM =  /*line :1*/ cUo1caLrG(vu1nEO26a82W)
		}
		if  /*line :1*/ d583PKWfr.Kind() == abi.Interface &&  /*line :1*/ d583PKWfr.IsNil() {

			return ZPN6dR{vu1nEO26a82W, nil,  /*line :1*/ fXMYmkc(abi.Interface)}
		}
		jJCfpek :=  /*line :1*/ m4SpTD7Rab(d583PKWfr)
		if  /*line :1*/ vu1nEO26a82W.NumMethod() == 0 {
			*(* /*line :1*/ any)(vhuI3TM) = jJCfpek
		} else {
			 /*line :1*/ fuAWeAR5Tm(vu1nEO26a82W, jJCfpek, vhuI3TM)
		}
		return ZPN6dR{vu1nEO26a82W, vhuI3TM, flagIndir |  /*line :1*/ fXMYmkc(abi.Interface)}
	}

	 /*line :1*/ panic(rad5bjl6bHs + ": value of type " +  /*line :1*/ qK4nvEMY(d583PKWfr.dDlYMUKHt2j).String() + " is not assignable to type " +  /*line :1*/ qK4nvEMY(vu1nEO26a82W).String())
}

func mAjgIGv(rE5JcBz3l unsafe.Pointer, dujSIeNg10V int, sTrQhjf uintptr, jo0rnS string) unsafe.Pointer {
	return  /*line :1*/ pOBZZZ(rE5JcBz3l,  /*line :1*/ uintptr(dujSIeNg10V)*sTrQhjf, "i < len")
}

func fuAWeAR5Tm(gc1v5B2w4C *abi.Type, bFIzEa6wQzhd any, vu1nEO26a82W unsafe.Pointer)

//go:noescape
func rFMmsMP6H(gc1v5B2w4C *abi.Type, vu1nEO26a82W, bFIzEa6wQzhd unsafe.Pointer)

func sQL9ue1Cs0x(jJCfpek any) {
	if rMTTDeYUTx2.y56CrSJqBkV {
		rMTTDeYUTx2.j6wtjqZor = jJCfpek
	}
}

var rMTTDeYUTx2 struct {
	y56CrSJqBkV	bool
	j6wtjqZor	any
}
