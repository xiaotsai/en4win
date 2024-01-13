//line :1
package GMNBnK

import (
	sysdll "O8h6BQj"
	syscall "bUKeamF"
	"unsafe"
)

var _ unsafe.Pointer



const (
	errnoERROR_IO_PENDING = 997
)

var (
	na9GRmpwfTn	error	=  /*line :1*/ syscall.O9Mga3(errnoERROR_IO_PENDING)
	i6hWQLNnBaJc	error	= syscall.EINVAL
)



func swlmKGHjo5(j0B7wBI syscall.O9Mga3) error {
	switch j0B7wBI {
	case 0:
		return i6hWQLNnBaJc
	case errnoERROR_IO_PENDING:
		return na9GRmpwfTn
	}

	return j0B7wBI
}

var (
	z75D75	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("advapi32.dll"))
	r6AGzj	=  /*line :1*/ syscall.A9tYHZzf( /*line :1*/ sysdll.GmPlJ_cqrvNd("kernel32.dll"))

	opmXjpREhm5M	=  /*line :1*/ z75D75.NewProc("RegCreateKeyExW")
	aLqBeP	=  /*line :1*/ z75D75.NewProc("RegDeleteKeyW")
	hFv17a	=  /*line :1*/ z75D75.NewProc("RegDeleteValueW")
	bwMOkhxI27yJ	=  /*line :1*/ z75D75.NewProc("RegEnumValueW")
	tRNKmhlJ	=  /*line :1*/ z75D75.NewProc("RegLoadMUIStringW")
	dMtZqejtS	=  /*line :1*/ z75D75.NewProc("RegSetValueExW")
	b6G_c6R	=  /*line :1*/ r6AGzj.NewProc("ExpandEnvironmentStringsW")
)

func jQ9Dwq597(k4Gb4p2R8Gpc syscall.SRlvVjrYa, wNnx4v *uint16, vq0HlZv uint32, waerJOeYellF *uint16, afzdQJ uint32, aqzKfjXH49h uint32, tjy0wMvF *syscall.DDwuM6RpYja, jRrJ6y5UpbH7 *syscall.SRlvVjrYa, rKNCvoI *uint32) (cm6z1vT3iEy error) {
	avnFaaX1CIiJ, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ opmXjpREhm5M.Addr(), 9,  /*line :1*/ uintptr(k4Gb4p2R8Gpc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wNnx4v)),  /*line :1*/ uintptr(vq0HlZv),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(waerJOeYellF)),  /*line :1*/ uintptr(afzdQJ),  /*line :1*/ uintptr(aqzKfjXH49h),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(tjy0wMvF)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(jRrJ6y5UpbH7)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(rKNCvoI)))
	if avnFaaX1CIiJ != 0 {
		cm6z1vT3iEy =  /*line :1*/ syscall.O9Mga3(avnFaaX1CIiJ)
	}
	return
}

func d3PMbL(k4Gb4p2R8Gpc syscall.SRlvVjrYa, wNnx4v *uint16) (cm6z1vT3iEy error) {
	avnFaaX1CIiJ, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ aLqBeP.Addr(), 2,  /*line :1*/ uintptr(k4Gb4p2R8Gpc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(wNnx4v)), 0)
	if avnFaaX1CIiJ != 0 {
		cm6z1vT3iEy =  /*line :1*/ syscall.O9Mga3(avnFaaX1CIiJ)
	}
	return
}

func toNASzC(k4Gb4p2R8Gpc syscall.SRlvVjrYa, ihk6jlqLI *uint16) (cm6z1vT3iEy error) {
	avnFaaX1CIiJ, _, _ :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ hFv17a.Addr(), 2,  /*line :1*/ uintptr(k4Gb4p2R8Gpc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ihk6jlqLI)), 0)
	if avnFaaX1CIiJ != 0 {
		cm6z1vT3iEy =  /*line :1*/ syscall.O9Mga3(avnFaaX1CIiJ)
	}
	return
}

func hzM0h_(k4Gb4p2R8Gpc syscall.SRlvVjrYa, dVJPFrui uint32, ihk6jlqLI *uint16, lJCTxANpMjV *uint32, vq0HlZv *uint32, aehKEp *uint32, zwMp0T *byte, xmDfVl5oS *uint32) (cm6z1vT3iEy error) {
	avnFaaX1CIiJ, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ bwMOkhxI27yJ.Addr(), 8,  /*line :1*/ uintptr(k4Gb4p2R8Gpc),  /*line :1*/ uintptr(dVJPFrui),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ihk6jlqLI)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lJCTxANpMjV)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(vq0HlZv)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(aehKEp)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zwMp0T)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(xmDfVl5oS)), 0)
	if avnFaaX1CIiJ != 0 {
		cm6z1vT3iEy =  /*line :1*/ syscall.O9Mga3(avnFaaX1CIiJ)
	}
	return
}

func u3OxV2gI_(k4Gb4p2R8Gpc syscall.SRlvVjrYa, ihk6jlqLI *uint16, zwMp0T *uint16, xmDfVl5oS uint32, nVwKaupqUOh *uint32, tkrXHjiQ7II uint32, gYnYUVDX *uint16) (cm6z1vT3iEy error) {
	avnFaaX1CIiJ, _, _ :=  /*line :1*/ syscall.Rc0O05UsV( /*line :1*/ tRNKmhlJ.Addr(), 7,  /*line :1*/ uintptr(k4Gb4p2R8Gpc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(ihk6jlqLI)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zwMp0T)),  /*line :1*/ uintptr(xmDfVl5oS),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(nVwKaupqUOh)),  /*line :1*/ uintptr(tkrXHjiQ7II),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(gYnYUVDX)), 0, 0)
	if avnFaaX1CIiJ != 0 {
		cm6z1vT3iEy =  /*line :1*/ syscall.O9Mga3(avnFaaX1CIiJ)
	}
	return
}

func wIasC9b(k4Gb4p2R8Gpc syscall.SRlvVjrYa, b6Aa6wEjpb *uint16, vq0HlZv uint32, j2nyUzpam uint32, zwMp0T *byte, kbsQqs uint32) (cm6z1vT3iEy error) {
	avnFaaX1CIiJ, _, _ :=  /*line :1*/ syscall.VeAKF8sAwft( /*line :1*/ dMtZqejtS.Addr(), 6,  /*line :1*/ uintptr(k4Gb4p2R8Gpc),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(b6Aa6wEjpb)),  /*line :1*/ uintptr(vq0HlZv),  /*line :1*/ uintptr(j2nyUzpam),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(zwMp0T)),  /*line :1*/ uintptr(kbsQqs))
	if avnFaaX1CIiJ != 0 {
		cm6z1vT3iEy =  /*line :1*/ syscall.O9Mga3(avnFaaX1CIiJ)
	}
	return
}

func zEpSOSa6(lYgpqYj *uint16, sexnZN17 *uint16, ctwAyGD3_ uint32) (dEntjI5C8pQ uint32, vbFroa5 error) {
	avnFaaX1CIiJ, _, sDHcJ_oXeag1 :=  /*line :1*/ syscall.OwKnfO5tgvm4( /*line :1*/ b6G_c6R.Addr(), 3,  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(lYgpqYj)),  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(sexnZN17)),  /*line :1*/ uintptr(ctwAyGD3_))
	dEntjI5C8pQ =  /*line :1*/ uint32(avnFaaX1CIiJ)
	if dEntjI5C8pQ == 0 {
		vbFroa5 =  /*line :1*/ swlmKGHjo5(sDHcJ_oXeag1)
	}
	return
}
