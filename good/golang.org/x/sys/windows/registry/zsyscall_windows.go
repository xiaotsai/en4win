//line :1
package ER_Q0ps1VO

import (
	syscall "bUKeamF"
	"unsafe"

	windows "NJ4MerJ"
)

var _ unsafe.Pointer



const (
	errnoERROR_IO_PENDING = 997
)

var (
	r64ih8vF	error	=  /*line :1*/ syscall.O9Mga3(errnoERROR_IO_PENDING)
	k6TAwe	error	= syscall.EINVAL
)



func aJ3F6o4MN(ux1Fty5Y syscall.O9Mga3) error {
	switch ux1Fty5Y {
	case 0:
		return k6TAwe
	case errnoERROR_IO_PENDING:
		return r64ih8vF
	}

	return ux1Fty5Y
}

var (
	b_ahnisKSie	=  /*line :1*/ windows.FDchPxUZ("advapi32.dll")
	fwgR_6gfw2c	=  /*line :1*/ windows.FDchPxUZ("kernel32.dll")

	r0cxaN5Nz	=  /*line :1*/ b_ahnisKSie.NewProc("RegConnectRegistryW")
	f2ncaGYvw	=  /*line :1*/ b_ahnisKSie.NewProc("RegCreateKeyExW")
	cJy63nZ	=  /*line :1*/ b_ahnisKSie.NewProc("RegDeleteKeyW")
	aAWZ21qwXl	=  /*line :1*/ b_ahnisKSie.NewProc("RegDeleteValueW")
	gk38nK	=  /*line :1*/ b_ahnisKSie.NewProc("RegEnumValueW")
	hfU5IgoDPZ	=  /*line :1*/ b_ahnisKSie.NewProc("RegLoadMUIStringW")
	_TQ3RDiZwB	=  /*line :1*/ b_ahnisKSie.NewProc("RegSetValueExW")
	mSyHCrj7j2Eo	=  /*line :1*/ fwgR_6gfw2c.NewProc("ExpandEnvironmentStringsW")
)

func zW6WK191u(vjrcfW *uint16, y9Iz7THEanWO syscall.SRlvVjrYa, xwbUefa *syscall.SRlvVjrYa) (aYnGnjNuUAr error) {
	vkXPnm, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ r0cxaN5Nz.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vjrcfW)),  /*line :1*/ uintptr(y9Iz7THEanWO),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xwbUefa)))
	if vkXPnm != 0 {
		aYnGnjNuUAr =  /*line :1*/ syscall.O9Mga3(vkXPnm)
	}
	return
}

func ghyTj4Lur(y9Iz7THEanWO syscall.SRlvVjrYa, xW2K3by *uint16, cTLJTE6 uint32, jx9euWemD_ *uint16, rnHcI82Kiua uint32, eDDZjQtEaB13 uint32, cS3er64 *syscall.DDwuM6RpYja, xwbUefa *syscall.SRlvVjrYa, iyq5drD *uint32) (aYnGnjNuUAr error) {
	vkXPnm, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ f2ncaGYvw.Addr(), 9,  /*line :1*/ uintptr(y9Iz7THEanWO),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xW2K3by)),  /*line :1*/ uintptr(cTLJTE6),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jx9euWemD_)),  /*line :1*/ uintptr(rnHcI82Kiua),  /*line :1*/ uintptr(eDDZjQtEaB13),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cS3er64)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xwbUefa)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(iyq5drD)))
	if vkXPnm != 0 {
		aYnGnjNuUAr =  /*line :1*/ syscall.O9Mga3(vkXPnm)
	}
	return
}

func oBIbRrud(y9Iz7THEanWO syscall.SRlvVjrYa, xW2K3by *uint16) (aYnGnjNuUAr error) {
	vkXPnm, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ cJy63nZ.Addr(), 2,  /*line :1*/ uintptr(y9Iz7THEanWO),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xW2K3by)), 0)
	if vkXPnm != 0 {
		aYnGnjNuUAr =  /*line :1*/ syscall.O9Mga3(vkXPnm)
	}
	return
}

func zPUq_hvkPJRN(y9Iz7THEanWO syscall.SRlvVjrYa, jy58GAHux *uint16) (aYnGnjNuUAr error) {
	vkXPnm, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aAWZ21qwXl.Addr(), 2,  /*line :1*/ uintptr(y9Iz7THEanWO),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jy58GAHux)), 0)
	if vkXPnm != 0 {
		aYnGnjNuUAr =  /*line :1*/ syscall.O9Mga3(vkXPnm)
	}
	return
}

func eaFjMXbf(y9Iz7THEanWO syscall.SRlvVjrYa, yfNAYc uint32, jy58GAHux *uint16, jDMf1S *uint32, cTLJTE6 *uint32, xN8Ac8a7xRi *uint32, _753m9TV2 *byte, e5ZJCFVvq0UN *uint32) (aYnGnjNuUAr error) {
	vkXPnm, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ gk38nK.Addr(), 8,  /*line :1*/ uintptr(y9Iz7THEanWO),  /*line :1*/ uintptr(yfNAYc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jy58GAHux)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jDMf1S)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cTLJTE6)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xN8Ac8a7xRi)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_753m9TV2)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(e5ZJCFVvq0UN)), 0)
	if vkXPnm != 0 {
		aYnGnjNuUAr =  /*line :1*/ syscall.O9Mga3(vkXPnm)
	}
	return
}

func gByj77H6i(y9Iz7THEanWO syscall.SRlvVjrYa, jy58GAHux *uint16, _753m9TV2 *uint16, e5ZJCFVvq0UN uint32, xqVULthqrIl *uint32, jfvXm7bK3SFa uint32, u2TfR8_QVff *uint16) (aYnGnjNuUAr error) {
	vkXPnm, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ hfU5IgoDPZ.Addr(), 7,  /*line :1*/ uintptr(y9Iz7THEanWO),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jy58GAHux)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_753m9TV2)),  /*line :1*/ uintptr(e5ZJCFVvq0UN),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xqVULthqrIl)),  /*line :1*/ uintptr(jfvXm7bK3SFa),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(u2TfR8_QVff)), 0, 0)
	if vkXPnm != 0 {
		aYnGnjNuUAr =  /*line :1*/ syscall.O9Mga3(vkXPnm)
	}
	return
}

func fBzuLbgu(y9Iz7THEanWO syscall.SRlvVjrYa, cbQFi17aFaNv *uint16, cTLJTE6 uint32, eKENv4R uint32, _753m9TV2 *byte, txyeAQ uint32) (aYnGnjNuUAr error) {
	vkXPnm, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ _TQ3RDiZwB.Addr(), 6,  /*line :1*/ uintptr(y9Iz7THEanWO),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cbQFi17aFaNv)),  /*line :1*/ uintptr(cTLJTE6),  /*line :1*/ uintptr(eKENv4R),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(_753m9TV2)),  /*line :1*/ uintptr(txyeAQ))
	if vkXPnm != 0 {
		aYnGnjNuUAr =  /*line :1*/ syscall.O9Mga3(vkXPnm)
	}
	return
}

func eQejQjwX(h_ZAINVl *uint16, w6HMTPR82d *uint16, xAvjmyFwTU uint32) (aMs6DlZl4 uint32, konG16p error) {
	vkXPnm, _, jSTdzFFm :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ mSyHCrj7j2Eo.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(h_ZAINVl)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(w6HMTPR82d)),  /*line :1*/ uintptr(xAvjmyFwTU))
	aMs6DlZl4 =  /*line :1*/ uint32(vkXPnm)
	if aMs6DlZl4 == 0 {
		konG16p =  /*line :1*/ aJ3F6o4MN(jSTdzFFm)
	}
	return
}
