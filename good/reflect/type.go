//line :1
package reflect

import (
	"internal/abi"
	"internal/goarch"
	strconv "zBESu2SrRjP"
	sync "sync"
	unicode "vGJV8maRK"
	utf8 "CuAc0pPZwf"
	"unsafe"
)

type YJh989LTZsVX interface {
	Align() int

	FieldAlign() int

	Method(int) Method

	MethodByName(string) (Method, bool)

	NumMethod() int

	Name() string

	PkgPath() string

	Size() uintptr

	String() string

	Kind() WzYBjsQL0

	Implements(iaoHrRg1 YJh989LTZsVX) bool

	AssignableTo(iaoHrRg1 YJh989LTZsVX) bool

	ConvertibleTo(iaoHrRg1 YJh989LTZsVX) bool

	Comparable() bool

	Bits() int

	ChanDir() Svj3YDOP3fa

	IsVariadic() bool

	Elem() YJh989LTZsVX

	Field(hXZpj1nTZ int) EeSeBImJjDQo

	FieldByIndex(hEqkR0O []int) EeSeBImJjDQo

	FieldByName(lGwAXqN4Hb string) (EeSeBImJjDQo, bool)

	FieldByNameFunc(xH9tRvQF4 func(string) bool) (EeSeBImJjDQo, bool)

	In(hXZpj1nTZ int) YJh989LTZsVX

	Key() YJh989LTZsVX

	Len() int

	NumField() int

	NumIn() int

	NumOut() int

	Out(hXZpj1nTZ int) YJh989LTZsVX

	z6hGxTABM1() *abi.Type
	aXE3rCbNCh() *aKCY_EtrgnA8
}

type WzYBjsQL0 uint

const (
	Invalid	WzYBjsQL0	= iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

const Ptr = Pointer

type aKCY_EtrgnA8 = abi.UncommonType

type z6hGxTABM1 struct {
	abi.Type
}

type rtype struct {
	t abi.Type
}

func (sw8_lJ *rtype) z6hGxTABM1() *abi.Type {
	return &sw8_lJ.t
}

func (sw8_lJ *rtype) aXE3rCbNCh() *abi.UncommonType {
	return  /*line :1*/ sw8_lJ.t.Uncommon()
}

type gTCizzIhT = abi.NameOff
type jbg3VYXczBk = abi.TypeOff
type mPSwONqw = abi.TextOff

type Svj3YDOP3fa int

const (
	RecvDir	Svj3YDOP3fa	= 1 << iota
	SendDir
	BothDir	= RecvDir | SendDir
)

type uV_BfrSO = abi.ArrayType

type aK0ovNmjQBD = abi.ChanType

type bY2Zo8pis = abi.FuncType

type ywQcz6y struct {
	abi.InterfaceType
}

func (sw8_lJ *ywQcz6y) zPya6B(e65m4G53plB gTCizzIhT) abi.Name {
	return  /*line :1*/ j5Z_WRWLudek(&sw8_lJ.Type).zPya6B(e65m4G53plB)
}

func drswiOv(sw8_lJ *abi.Type, e65m4G53plB gTCizzIhT) abi.Name {
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ).zPya6B(e65m4G53plB)
}

func nsOKvPP(sw8_lJ *abi.Type, e65m4G53plB jbg3VYXczBk) *abi.Type {
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ).x9CINy8a(e65m4G53plB)
}

func (sw8_lJ *ywQcz6y) x9CINy8a(e65m4G53plB jbg3VYXczBk) *abi.Type {
	return  /*line :1*/ j5Z_WRWLudek(&sw8_lJ.Type).x9CINy8a(e65m4G53plB)
}

func (sw8_lJ *ywQcz6y) z6hGxTABM1() *abi.Type {
	return &sw8_lJ.Type
}

func (sw8_lJ *ywQcz6y) aXE3rCbNCh() *abi.UncommonType {
	return  /*line :1*/ sw8_lJ.Uncommon()
}

type r0LDmUgzY struct {
	abi.MapType
}

type uSkasB struct {
	abi.PtrType
}

type vwvVHq struct {
	abi.SliceType
}

type yUesa1U9 = abi.StructField

type structType struct {
	abi.StructType
}

func zIgMu_ZY(krtR1HwEan abi.Name) string {
	if krtR1HwEan.Bytes == nil || * /*line :1*/ krtR1HwEan.DataChecked(0, "name flag field")&(1<<2) == 0 {
		return ""
	}
	hXZpj1nTZ, uKa5XMKhg1r :=  /*line :1*/ krtR1HwEan.ReadVarint(1)
	e65m4G53plB := 1 + hXZpj1nTZ + uKa5XMKhg1r
	if  /*line :1*/ krtR1HwEan.HasTag() {
		erwKJZde, la20T4j2P :=  /*line :1*/ krtR1HwEan.ReadVarint(e65m4G53plB)
		e65m4G53plB += erwKJZde + la20T4j2P
	}
	var zPya6B int32

	 /*line :1*/ copy((*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(&zPya6B))[:], (*[4] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer( /*line :1*/ krtR1HwEan.DataChecked(e65m4G53plB, "name offset field")))[:])
	bHMYpV8 := abi.Name{Bytes: (* /*line :1*/ byte)( /*line :1*/ oKsBJ_y( /*line :1*/ unsafe.Pointer(krtR1HwEan.Bytes), zPya6B))}
	return  /*line :1*/ bHMYpV8.Name()
}

func kxqWl_C(krtR1HwEan, wXfiW7aVma4Q string, gM1OiS, ceja2s bool) abi.Name {
	return  /*line :1*/ abi.NewName(krtR1HwEan, wXfiW7aVma4Q, gM1OiS, ceja2s)
}

type Method struct {
	ZYP3IUVky	string

	EiqDvG3j0	string

	FLQnlZ2eg	YJh989LTZsVX
	IW56_bcBk	Value
	JhrNuhrH	int
}

func (hek9DV Method) IsExported() bool {
	return hek9DV.EiqDvG3j0 == ""
}

const (
	kindDirectIface	= 1 << 5
	kindGCProg	= 1 << 6
	kindMask	= (1 << 5) - 1
)

func (lYOqUQga WzYBjsQL0) String() string {
	if  /*line :1*/ uint(lYOqUQga) <  /*line :1*/ uint( /*line :1*/ len(es1Tr9)) {
		return es1Tr9[ /*line :1*/ uint(lYOqUQga)]
	}
	return "kind" +  /*line :1*/ strconv.ZW1_FxJtq( /*line :1*/ int(lYOqUQga))
}

var es1Tr9 = []string{
	Invalid:	"invalid",
	Bool:	"bool",
	Int:	"int",
	Int8:	"int8",
	Int16:	"int16",
	Int32:	"int32",
	Int64:	"int64",
	Uint:	"uint",
	Uint8:	"uint8",
	Uint16:	"uint16",
	Uint32:	"uint32",
	Uint64:	"uint64",
	Uintptr:	"uintptr",
	Float32:	"float32",
	Float64:	"float64",
	Complex64:	"complex64",
	Complex128:	"complex128",
	Array:	"array",
	Chan:	"chan",
	Func:	"func",
	Interface:	"interface",
	Map:	"map",
	Pointer:	"ptr",
	Slice:	"slice",
	String:	"string",
	Struct:	"struct",
	UnsafePointer:	"unsafe.Pointer",
}

//go:noescape
func flNF5xHuhxQ(oROfWx36ds unsafe.Pointer, e65m4G53plB int32) unsafe.Pointer

//go:noescape
func oKsBJ_y(l3O0iHi03uMT unsafe.Pointer, e65m4G53plB int32) unsafe.Pointer

//go:noescape
func _xEXWFx(l3O0iHi03uMT unsafe.Pointer, e65m4G53plB int32) unsafe.Pointer

//go:noescape
func nHZ8MX(cf3fCV8ayFq unsafe.Pointer) int32

func vybOmhD9hqa5(krtR1HwEan abi.Name) gTCizzIhT {
	return  /*line :1*/ gTCizzIhT( /*line :1*/ nHZ8MX( /*line :1*/ unsafe.Pointer(krtR1HwEan.Bytes)))
}

func j2acIXKpV(sw8_lJ *abi.Type) jbg3VYXczBk {
	return  /*line :1*/ jbg3VYXczBk( /*line :1*/ nHZ8MX( /*line :1*/ unsafe.Pointer(sw8_lJ)))
}

func gGBM5WeQM(cf3fCV8ayFq unsafe.Pointer) mPSwONqw {
	return  /*line :1*/ mPSwONqw( /*line :1*/ nHZ8MX(cf3fCV8ayFq))
}

func (sw8_lJ *rtype) zPya6B(e65m4G53plB gTCizzIhT) abi.Name {
	return abi.Name{Bytes: (* /*line :1*/ byte)( /*line :1*/ flNF5xHuhxQ( /*line :1*/ unsafe.Pointer(sw8_lJ),  /*line :1*/ int32(e65m4G53plB)))}
}

func (sw8_lJ *rtype) x9CINy8a(e65m4G53plB jbg3VYXczBk) *abi.Type {
	return (* /*line :1*/ abi.Type)( /*line :1*/ oKsBJ_y( /*line :1*/ unsafe.Pointer(sw8_lJ),  /*line :1*/ int32(e65m4G53plB)))
}

func (sw8_lJ *rtype) nv_FdA(e65m4G53plB mPSwONqw) unsafe.Pointer {
	return  /*line :1*/ _xEXWFx( /*line :1*/ unsafe.Pointer(sw8_lJ),  /*line :1*/ int32(e65m4G53plB))
}

func zTRltR(sw8_lJ *abi.Type, e65m4G53plB mPSwONqw) unsafe.Pointer {
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ).nv_FdA(e65m4G53plB)
}

func (sw8_lJ *rtype) String() string {
	kXTvhUI4Zgg5 :=  /*line :1*/ sw8_lJ.zPya6B(sw8_lJ.t.Str).Name()
	if sw8_lJ.t.TFlag&abi.TFlagExtraStar != 0 {
		return kXTvhUI4Zgg5[1:]
	}
	return kXTvhUI4Zgg5
}

func (sw8_lJ *rtype) Size() uintptr	{ return  /*line :1*/ sw8_lJ.t.Size() }

func (sw8_lJ *rtype) Bits() int {
	if sw8_lJ == nil {
		 /*line :1*/ panic("reflect: Bits of nil Type")
	}
	lYOqUQga :=  /*line :1*/ sw8_lJ.Kind()
	if lYOqUQga < Int || lYOqUQga > Complex128 {
		 /*line :1*/ panic("reflect: Bits of non-arithmetic Type " +  /*line :1*/ sw8_lJ.String())
	}
	return  /*line :1*/ int(sw8_lJ.t.Size_) * 8
}

func (sw8_lJ *rtype) Align() int	{ return  /*line :1*/ sw8_lJ.t.Align() }

func (sw8_lJ *rtype) FieldAlign() int	{ return  /*line :1*/ sw8_lJ.t.FieldAlign() }

func (sw8_lJ *rtype) Kind() WzYBjsQL0	{ return  /*line :1*/ WzYBjsQL0( /*line :1*/ sw8_lJ.t.Kind()) }

func (sw8_lJ *rtype) _sARiSYhaOve() []abi.Method {
	sxpcYEiFEM :=  /*line :1*/ sw8_lJ.aXE3rCbNCh()
	if sxpcYEiFEM == nil {
		return nil
	}
	return  /*line :1*/ sxpcYEiFEM.ExportedMethods()
}

func (sw8_lJ *rtype) NumMethod() int {
	if  /*line :1*/ sw8_lJ.Kind() == Interface {
		mP7QqSVInV := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		return  /*line :1*/ mP7QqSVInV.NumMethod()
	}
	return  /*line :1*/ len( /*line :1*/ sw8_lJ._sARiSYhaOve())
}

func (sw8_lJ *rtype) Method(hXZpj1nTZ int) (hek9DV Method) {
	if  /*line :1*/ sw8_lJ.Kind() == Interface {
		mP7QqSVInV := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		return  /*line :1*/ mP7QqSVInV.Method(hXZpj1nTZ)
	}
	r2zNEKaY :=  /*line :1*/ sw8_lJ._sARiSYhaOve()
	if hXZpj1nTZ < 0 || hXZpj1nTZ >=  /*line :1*/ len(r2zNEKaY) {
		 /*line :1*/ panic("reflect: Method index out of range")
	}
	e3IyyaQSEj := r2zNEKaY[hXZpj1nTZ]
	nuhaXQxOB :=  /*line :1*/ sw8_lJ.zPya6B(e3IyyaQSEj.Name)
	hek9DV.ZYP3IUVky =  /*line :1*/ nuhaXQxOB.Name()
	iiZKy9PE3 :=  /*line :1*/ flag(Func)
	srqll74r :=  /*line :1*/ sw8_lJ.x9CINy8a(e3IyyaQSEj.Mtyp)
	yzEE5FI := (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer(srqll74r))
	gJ3XGL_F :=  /*line :1*/ make([]YJh989LTZsVX, 0, 1+ /*line :1*/ yzEE5FI.NumIn())
	gJ3XGL_F =  /*line :1*/ append(gJ3XGL_F, sw8_lJ)
	for _, cUbsbg9j := range  /*line :1*/ yzEE5FI.InSlice() {
		gJ3XGL_F =  /*line :1*/ append(gJ3XGL_F,  /*line :1*/ j5Z_WRWLudek(cUbsbg9j))
	}
	hEbpexD :=  /*line :1*/ make([]YJh989LTZsVX, 0,  /*line :1*/ yzEE5FI.NumOut())
	for _, aLsdcpyP1 := range  /*line :1*/ yzEE5FI.OutSlice() {
		hEbpexD =  /*line :1*/ append(hEbpexD,  /*line :1*/ j5Z_WRWLudek(aLsdcpyP1))
	}
	byRNfAUjb :=  /*line :1*/ NlPxWAGa0V(gJ3XGL_F, hEbpexD,  /*line :1*/ yzEE5FI.IsVariadic())
	hek9DV.FLQnlZ2eg = byRNfAUjb
	bGGxfg :=  /*line :1*/ sw8_lJ.nv_FdA(e3IyyaQSEj.Tfn)
	lb3zWvJNFBQ :=  /*line :1*/ unsafe.Pointer(&bGGxfg)
	hek9DV.IW56_bcBk = Value{&byRNfAUjb.(*rtype).t, lb3zWvJNFBQ, iiZKy9PE3}

	hek9DV.JhrNuhrH = hXZpj1nTZ
	return hek9DV
}

func (sw8_lJ *rtype) MethodByName(lGwAXqN4Hb string) (hek9DV Method, bDfzXlCm5Raf bool) {
	if  /*line :1*/ sw8_lJ.Kind() == Interface {
		mP7QqSVInV := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		return  /*line :1*/ mP7QqSVInV.MethodByName(lGwAXqN4Hb)
	}
	sxpcYEiFEM :=  /*line :1*/ sw8_lJ.aXE3rCbNCh()
	if sxpcYEiFEM == nil {
		return Method{}, false
	}

	r2zNEKaY :=  /*line :1*/ sxpcYEiFEM.ExportedMethods()

	hXZpj1nTZ, kNVeJaGxs := 0,  /*line :1*/ len(r2zNEKaY)
	for hXZpj1nTZ < kNVeJaGxs {
		cJS_wcOmyXDz :=  /*line :1*/ int( /*line :1*/ uint(hXZpj1nTZ+kNVeJaGxs) >> 1)

		if !( /*line :1*/ sw8_lJ.zPya6B(r2zNEKaY[cJS_wcOmyXDz].Name).Name() >= lGwAXqN4Hb) {
			hXZpj1nTZ = cJS_wcOmyXDz + 1
		} else {
			kNVeJaGxs = cJS_wcOmyXDz
		}
	}

	if hXZpj1nTZ <  /*line :1*/ len(r2zNEKaY) && lGwAXqN4Hb ==  /*line :1*/ sw8_lJ.zPya6B(r2zNEKaY[hXZpj1nTZ].Name).Name() {
		return  /*line :1*/ sw8_lJ.Method(hXZpj1nTZ), true
	}

	return Method{}, false
}

func (sw8_lJ *rtype) PkgPath() string {
	if sw8_lJ.t.TFlag&abi.TFlagNamed == 0 {
		return ""
	}
	sxpcYEiFEM :=  /*line :1*/ sw8_lJ.aXE3rCbNCh()
	if sxpcYEiFEM == nil {
		return ""
	}
	return  /*line :1*/ sw8_lJ.zPya6B(sxpcYEiFEM.PkgPath).Name()
}

func hP5IxCk8Hfi(sw8_lJ *abi.Type) string {
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ).PkgPath()
}

func (sw8_lJ *rtype) Name() string {
	if ! /*line :1*/ sw8_lJ.t.HasName() {
		return ""
	}
	kXTvhUI4Zgg5 :=  /*line :1*/ sw8_lJ.String()
	hXZpj1nTZ :=  /*line :1*/ len(kXTvhUI4Zgg5) - 1
	iVD4ENNPiv := 0
	for hXZpj1nTZ >= 0 && (kXTvhUI4Zgg5[hXZpj1nTZ] != '.' || iVD4ENNPiv != 0) {
		switch kXTvhUI4Zgg5[hXZpj1nTZ] {
		case ']':
			iVD4ENNPiv++
		case '[':
			iVD4ENNPiv--
		}
		hXZpj1nTZ--
	}
	return kXTvhUI4Zgg5[hXZpj1nTZ+1:]
}

func jv91ZkF(sw8_lJ *abi.Type) string {
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ).Name()
}

func (sw8_lJ *rtype) ChanDir() Svj3YDOP3fa {
	if  /*line :1*/ sw8_lJ.Kind() != Chan {
		 /*line :1*/ panic("reflect: ChanDir of non-chan type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ abi.ChanType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ Svj3YDOP3fa(mP7QqSVInV.Dir)
}

func j5Z_WRWLudek(sw8_lJ *abi.Type) *rtype {
	return (* /*line :1*/ rtype)( /*line :1*/ unsafe.Pointer(sw8_lJ))
}

func dPe0EQn(sw8_lJ *abi.Type) *abi.Type {
	iwXCaXtB :=  /*line :1*/ sw8_lJ.Elem()
	if iwXCaXtB != nil {
		return iwXCaXtB
	}
	 /*line :1*/ panic("reflect: Elem of invalid type " +  /*line :1*/ iR_1JLd(sw8_lJ))
}

func (sw8_lJ *rtype) Elem() YJh989LTZsVX {
	return  /*line :1*/ nadfj4Ka8cQf( /*line :1*/ dPe0EQn( /*line :1*/ sw8_lJ.z6hGxTABM1()))
}

func (sw8_lJ *rtype) Field(hXZpj1nTZ int) EeSeBImJjDQo {
	if  /*line :1*/ sw8_lJ.Kind() != Struct {
		 /*line :1*/ panic("reflect: Field of non-struct type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ mP7QqSVInV.Field(hXZpj1nTZ)
}

func (sw8_lJ *rtype) FieldByIndex(hEqkR0O []int) EeSeBImJjDQo {
	if  /*line :1*/ sw8_lJ.Kind() != Struct {
		 /*line :1*/ panic("reflect: FieldByIndex of non-struct type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ mP7QqSVInV.FieldByIndex(hEqkR0O)
}

func (sw8_lJ *rtype) FieldByName(lGwAXqN4Hb string) (EeSeBImJjDQo, bool) {
	if  /*line :1*/ sw8_lJ.Kind() != Struct {
		 /*line :1*/ panic("reflect: FieldByName of non-struct type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ mP7QqSVInV.FieldByName(lGwAXqN4Hb)
}

func (sw8_lJ *rtype) FieldByNameFunc(xH9tRvQF4 func(string) bool) (EeSeBImJjDQo, bool) {
	if  /*line :1*/ sw8_lJ.Kind() != Struct {
		 /*line :1*/ panic("reflect: FieldByNameFunc of non-struct type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ mP7QqSVInV.FieldByNameFunc(xH9tRvQF4)
}

func (sw8_lJ *rtype) Key() YJh989LTZsVX {
	if  /*line :1*/ sw8_lJ.Kind() != Map {
		 /*line :1*/ panic("reflect: Key of non-map type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ nadfj4Ka8cQf(mP7QqSVInV.Key)
}

func (sw8_lJ *rtype) Len() int {
	if  /*line :1*/ sw8_lJ.Kind() != Array {
		 /*line :1*/ panic("reflect: Len of non-array type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ int(mP7QqSVInV.Len)
}

func (sw8_lJ *rtype) NumField() int {
	if  /*line :1*/ sw8_lJ.Kind() != Struct {
		 /*line :1*/ panic("reflect: NumField of non-struct type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ len(mP7QqSVInV.Fields)
}

func (sw8_lJ *rtype) In(hXZpj1nTZ int) YJh989LTZsVX {
	if  /*line :1*/ sw8_lJ.Kind() != Func {
		 /*line :1*/ panic("reflect: In of non-func type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ abi.FuncType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ nadfj4Ka8cQf( /*line :1*/ mP7QqSVInV.InSlice()[hXZpj1nTZ])
}

func (sw8_lJ *rtype) NumIn() int {
	if  /*line :1*/ sw8_lJ.Kind() != Func {
		 /*line :1*/ panic("reflect: NumIn of non-func type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ abi.FuncType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ mP7QqSVInV.NumIn()
}

func (sw8_lJ *rtype) NumOut() int {
	if  /*line :1*/ sw8_lJ.Kind() != Func {
		 /*line :1*/ panic("reflect: NumOut of non-func type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ abi.FuncType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ mP7QqSVInV.NumOut()
}

func (sw8_lJ *rtype) Out(hXZpj1nTZ int) YJh989LTZsVX {
	if  /*line :1*/ sw8_lJ.Kind() != Func {
		 /*line :1*/ panic("reflect: Out of non-func type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ abi.FuncType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ nadfj4Ka8cQf( /*line :1*/ mP7QqSVInV.OutSlice()[hXZpj1nTZ])
}

func (sw8_lJ *rtype) IsVariadic() bool {
	if  /*line :1*/ sw8_lJ.Kind() != Func {
		 /*line :1*/ panic("reflect: IsVariadic of non-func type " +  /*line :1*/ sw8_lJ.String())
	}
	mP7QqSVInV := (* /*line :1*/ abi.FuncType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
	return  /*line :1*/ mP7QqSVInV.IsVariadic()
}

func lxkP7oV7sm(e3IyyaQSEj unsafe.Pointer, uiTlN8bdKm uintptr, cdghXjtt string) unsafe.Pointer {
	return  /*line :1*/ unsafe.Pointer( /*line :1*/ uintptr(e3IyyaQSEj) + uiTlN8bdKm)
}

func (v40UV8 Svj3YDOP3fa) String() string {
	switch v40UV8 {
	case SendDir:
		return "chan<-"
	case RecvDir:
		return "<-chan"
	case BothDir:
		return "chan"
	}
	return "ChanDir" +  /*line :1*/ strconv.ZW1_FxJtq( /*line :1*/ int(v40UV8))
}

func (sw8_lJ *ywQcz6y) Method(hXZpj1nTZ int) (hek9DV Method) {
	if hXZpj1nTZ < 0 || hXZpj1nTZ >=  /*line :1*/ len(sw8_lJ.Methods) {
		return
	}
	e3IyyaQSEj := &sw8_lJ.Methods[hXZpj1nTZ]
	nuhaXQxOB :=  /*line :1*/ sw8_lJ.zPya6B(e3IyyaQSEj.Name)
	hek9DV.ZYP3IUVky =  /*line :1*/ nuhaXQxOB.Name()
	if ! /*line :1*/ nuhaXQxOB.IsExported() {
		hek9DV.EiqDvG3j0 =  /*line :1*/ zIgMu_ZY(nuhaXQxOB)
		if hek9DV.EiqDvG3j0 == "" {
			hek9DV.EiqDvG3j0 =  /*line :1*/ sw8_lJ.PkgPath.Name()
		}
	}
	hek9DV.FLQnlZ2eg =  /*line :1*/ nadfj4Ka8cQf( /*line :1*/ sw8_lJ.x9CINy8a(e3IyyaQSEj.Typ))
	hek9DV.JhrNuhrH = hXZpj1nTZ
	return
}

func (sw8_lJ *ywQcz6y) NumMethod() int	{ return  /*line :1*/ len(sw8_lJ.Methods) }

func (sw8_lJ *ywQcz6y) MethodByName(lGwAXqN4Hb string) (hek9DV Method, bDfzXlCm5Raf bool) {
	if sw8_lJ == nil {
		return
	}
	var e3IyyaQSEj *abi.Imethod
	for hXZpj1nTZ := range sw8_lJ.Methods {
		e3IyyaQSEj = &sw8_lJ.Methods[hXZpj1nTZ]
		if  /*line :1*/ sw8_lJ.zPya6B(e3IyyaQSEj.Name).Name() == lGwAXqN4Hb {
			return  /*line :1*/ sw8_lJ.Method(hXZpj1nTZ), true
		}
	}
	return
}

type EeSeBImJjDQo struct {
	ZOSBHkJz	string

	H1I0k6V	string

	S1miDNBAV8a	YJh989LTZsVX
	Kp_2l7Su8ao	ObtIIR
	YTyFIb6uUkmT	uintptr
	KOMPhI	[]int
	GuSnlgXxlhPz	bool
}

func (jToThzM EeSeBImJjDQo) IsExported() bool {
	return jToThzM.H1I0k6V == ""
}

type ObtIIR string

func (wXfiW7aVma4Q ObtIIR) Get(ccj6oONwQVN string) string {
	f8NmcatRx, _ :=  /*line :1*/ wXfiW7aVma4Q.Lookup(ccj6oONwQVN)
	return f8NmcatRx
}

func (wXfiW7aVma4Q ObtIIR) Lookup(ccj6oONwQVN string) (ba65CH string, bDfzXlCm5Raf bool) {

	for wXfiW7aVma4Q != "" {

		hXZpj1nTZ := 0
		for hXZpj1nTZ <  /*line :1*/ len(wXfiW7aVma4Q) && wXfiW7aVma4Q[hXZpj1nTZ] == ' ' {
			hXZpj1nTZ++
		}
		wXfiW7aVma4Q = wXfiW7aVma4Q[hXZpj1nTZ:]
		if wXfiW7aVma4Q == "" {
			break
		}

		hXZpj1nTZ = 0
		for hXZpj1nTZ <  /*line :1*/ len(wXfiW7aVma4Q) && wXfiW7aVma4Q[hXZpj1nTZ] > ' ' && wXfiW7aVma4Q[hXZpj1nTZ] != ':' && wXfiW7aVma4Q[hXZpj1nTZ] != '"' && wXfiW7aVma4Q[hXZpj1nTZ] != 0x7f {
			hXZpj1nTZ++
		}
		if hXZpj1nTZ == 0 || hXZpj1nTZ+1 >=  /*line :1*/ len(wXfiW7aVma4Q) || wXfiW7aVma4Q[hXZpj1nTZ] != ':' || wXfiW7aVma4Q[hXZpj1nTZ+1] != '"' {
			break
		}
		lGwAXqN4Hb :=  /*line :1*/ string(wXfiW7aVma4Q[:hXZpj1nTZ])
		wXfiW7aVma4Q = wXfiW7aVma4Q[hXZpj1nTZ+1:]

		hXZpj1nTZ = 1
		for hXZpj1nTZ <  /*line :1*/ len(wXfiW7aVma4Q) && wXfiW7aVma4Q[hXZpj1nTZ] != '"' {
			if wXfiW7aVma4Q[hXZpj1nTZ] == '\\' {
				hXZpj1nTZ++
			}
			hXZpj1nTZ++
		}
		if hXZpj1nTZ >=  /*line :1*/ len(wXfiW7aVma4Q) {
			break
		}
		bm5HtcYbr85 :=  /*line :1*/ string(wXfiW7aVma4Q[:hXZpj1nTZ+1])
		wXfiW7aVma4Q = wXfiW7aVma4Q[hXZpj1nTZ+1:]

		if ccj6oONwQVN == lGwAXqN4Hb {
			ba65CH, cEFr67C :=  /*line :1*/ strconv.D0aVL66(bm5HtcYbr85)
			if cEFr67C != nil {
				break
			}
			return ba65CH, true
		}
	}
	return "", false
}

func (sw8_lJ *structType) Field(hXZpj1nTZ int) (jToThzM EeSeBImJjDQo) {
	if hXZpj1nTZ < 0 || hXZpj1nTZ >=  /*line :1*/ len(sw8_lJ.Fields) {
		 /*line :1*/ panic("reflect: Field index out of bounds")
	}
	e3IyyaQSEj := &sw8_lJ.Fields[hXZpj1nTZ]
	jToThzM.S1miDNBAV8a =  /*line :1*/ nadfj4Ka8cQf(e3IyyaQSEj.Typ)
	jToThzM.ZOSBHkJz =  /*line :1*/ e3IyyaQSEj.Name.Name()
	jToThzM.GuSnlgXxlhPz =  /*line :1*/ e3IyyaQSEj.Embedded()
	if ! /*line :1*/ e3IyyaQSEj.Name.IsExported() {
		jToThzM.H1I0k6V =  /*line :1*/ sw8_lJ.PkgPath.Name()
	}
	if wXfiW7aVma4Q :=  /*line :1*/ e3IyyaQSEj.Name.Tag(); wXfiW7aVma4Q != "" {
		jToThzM.Kp_2l7Su8ao =  /*line :1*/ ObtIIR(wXfiW7aVma4Q)
	}
	jToThzM.YTyFIb6uUkmT = e3IyyaQSEj.Offset

	jToThzM.KOMPhI = []int{hXZpj1nTZ}
	return
}

func (sw8_lJ *structType) FieldByIndex(hEqkR0O []int) (jToThzM EeSeBImJjDQo) {
	jToThzM.S1miDNBAV8a =  /*line :1*/ nadfj4Ka8cQf(&sw8_lJ.Type)
	for hXZpj1nTZ, uiTlN8bdKm := range hEqkR0O {
		if hXZpj1nTZ > 0 {
			yzEE5FI := jToThzM.S1miDNBAV8a
			if  /*line :1*/ yzEE5FI.Kind() == Pointer &&  /*line :1*/ yzEE5FI.Elem().Kind() == Struct {
				yzEE5FI =  /*line :1*/ yzEE5FI.Elem()
			}
			jToThzM.S1miDNBAV8a = yzEE5FI
		}
		jToThzM =  /*line :1*/ jToThzM.S1miDNBAV8a.Field(uiTlN8bdKm)
	}
	return
}

type cGcjAxHteMyU struct {
	kRHMOlba	*structType
	aab8DHc	[]int
}

func (sw8_lJ *structType) FieldByNameFunc(xH9tRvQF4 func(string) bool) (y4Z3yUUHlmY0 EeSeBImJjDQo, bDfzXlCm5Raf bool) {

	vwbif1p := []cGcjAxHteMyU{}
	xzHSPQl := []cGcjAxHteMyU{{kRHMOlba: sw8_lJ}}

	var bQXR1UOZ96i map[*structType]int

	bId9CsuBq := map[*structType]bool{}

	for  /*line :1*/ len(xzHSPQl) > 0 {
		vwbif1p, xzHSPQl = xzHSPQl, vwbif1p[:0]
		zuUK8Kmu := bQXR1UOZ96i
		bQXR1UOZ96i = nil

		for _, wg1GrfGbQv := range vwbif1p {
			sw8_lJ := wg1GrfGbQv.kRHMOlba
			if bId9CsuBq[sw8_lJ] {

				continue
			}
			bId9CsuBq[sw8_lJ] = true
			for hXZpj1nTZ := range sw8_lJ.Fields {
				jToThzM := &sw8_lJ.Fields[hXZpj1nTZ]

				yFKOqObud :=  /*line :1*/ jToThzM.Name.Name()
				var hxB9OP *abi.Type
				if  /*line :1*/ jToThzM.Embedded() {

					hxB9OP = jToThzM.Typ
					if  /*line :1*/ hxB9OP.Kind() == abi.Pointer {
						hxB9OP =  /*line :1*/ hxB9OP.Elem()
					}
				}

				if  /*line :1*/ xH9tRvQF4(yFKOqObud) {

					if zuUK8Kmu[sw8_lJ] > 1 || bDfzXlCm5Raf {

						return EeSeBImJjDQo{}, false
					}
					y4Z3yUUHlmY0 =  /*line :1*/ sw8_lJ.Field(hXZpj1nTZ)
					y4Z3yUUHlmY0.KOMPhI = nil
					y4Z3yUUHlmY0.KOMPhI =  /*line :1*/ append(y4Z3yUUHlmY0.KOMPhI, wg1GrfGbQv.aab8DHc...)
					y4Z3yUUHlmY0.KOMPhI =  /*line :1*/ append(y4Z3yUUHlmY0.KOMPhI, hXZpj1nTZ)
					bDfzXlCm5Raf = true
					continue
				}

				if bDfzXlCm5Raf || hxB9OP == nil ||  /*line :1*/ hxB9OP.Kind() != abi.Struct {
					continue
				}
				sdoyg8mncAG := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(hxB9OP))
				if bQXR1UOZ96i[sdoyg8mncAG] > 0 {
					bQXR1UOZ96i[sdoyg8mncAG] = 2
					continue
				}
				if bQXR1UOZ96i == nil {
					bQXR1UOZ96i = map[*structType]int{}
				}
				bQXR1UOZ96i[sdoyg8mncAG] = 1
				if zuUK8Kmu[sw8_lJ] > 1 {
					bQXR1UOZ96i[sdoyg8mncAG] = 2
				}
				var hEqkR0O []int
				hEqkR0O =  /*line :1*/ append(hEqkR0O, wg1GrfGbQv.aab8DHc...)
				hEqkR0O =  /*line :1*/ append(hEqkR0O, hXZpj1nTZ)
				xzHSPQl =  /*line :1*/ append(xzHSPQl, cGcjAxHteMyU{sdoyg8mncAG, hEqkR0O})
			}
		}
		if bDfzXlCm5Raf {
			break
		}
	}
	return
}

func (sw8_lJ *structType) FieldByName(lGwAXqN4Hb string) (jToThzM EeSeBImJjDQo, f8iODilp8Iq bool) {

	bxMI3w := false
	if lGwAXqN4Hb != "" {
		for hXZpj1nTZ := range sw8_lJ.Fields {
			ccFQQ9kt6M := &sw8_lJ.Fields[hXZpj1nTZ]
			if  /*line :1*/ ccFQQ9kt6M.Name.Name() == lGwAXqN4Hb {
				return  /*line :1*/ sw8_lJ.Field(hXZpj1nTZ), true
			}
			if  /*line :1*/ ccFQQ9kt6M.Embedded() {
				bxMI3w = true
			}
		}
	}
	if !bxMI3w {
		return
	}
	return  /*line :1*/ sw8_lJ.FieldByNameFunc(func(kXTvhUI4Zgg5 string) bool { return kXTvhUI4Zgg5 == lGwAXqN4Hb })
}

func Cher1a2Fblr(hXZpj1nTZ any) YJh989LTZsVX {
	ique858 := *(* /*line :1*/ cKqj6F0L)( /*line :1*/ unsafe.Pointer(&hXZpj1nTZ))

	return  /*line :1*/ nadfj4Ka8cQf((* /*line :1*/ abi.Type)( /*line :1*/ m6KMjaQj( /*line :1*/ unsafe.Pointer(ique858.ubPJaJymi))))
}

func gC11Vi(hXZpj1nTZ any) *abi.Type {
	ique858 := *(* /*line :1*/ cKqj6F0L)( /*line :1*/ unsafe.Pointer(&hXZpj1nTZ))
	return ique858.ubPJaJymi
}

var vGn3oAB sync.Eim1b6bNi9IM

func ZhyIGFs9FD(sw8_lJ YJh989LTZsVX) YJh989LTZsVX	{ return  /*line :1*/ EjXqitstLUQt(sw8_lJ) }

func EjXqitstLUQt(sw8_lJ YJh989LTZsVX) YJh989LTZsVX {
	return  /*line :1*/ j5Z_WRWLudek( /*line :1*/ sw8_lJ.(*rtype).v8dZcSqC())
}

func (sw8_lJ *rtype) v8dZcSqC() *abi.Type {
	eU5PGA6aD3 := &sw8_lJ.t
	if eU5PGA6aD3.PtrToThis != 0 {
		return  /*line :1*/ sw8_lJ.x9CINy8a(eU5PGA6aD3.PtrToThis)
	}

	if p4rJlS0a8, bDfzXlCm5Raf :=  /*line :1*/ vGn3oAB.Load(sw8_lJ); bDfzXlCm5Raf {
		return &p4rJlS0a8.(*uSkasB).Type
	}

	kXTvhUI4Zgg5 := "*" +  /*line :1*/ sw8_lJ.String()
	for _, mP7QqSVInV := range  /*line :1*/ b0uDpEzf(kXTvhUI4Zgg5) {
		e3IyyaQSEj := (* /*line :1*/ uSkasB)( /*line :1*/ unsafe.Pointer(mP7QqSVInV))
		if e3IyyaQSEj.Elem != &sw8_lJ.t {
			continue
		}
		p4rJlS0a8, _ :=  /*line :1*/ vGn3oAB.LoadOrStore(sw8_lJ, e3IyyaQSEj)
		return &p4rJlS0a8.(*uSkasB).Type
	}

	var kHoIh5qjpGP any = (* /*line :1*/ unsafe.Pointer)(nil)
	fNjFa2jLl := *(** /*line :1*/ uSkasB)( /*line :1*/ unsafe.Pointer(&kHoIh5qjpGP))
	fMwDGjqeG := *fNjFa2jLl

	fMwDGjqeG.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(kXTvhUI4Zgg5, "", false, false))
	fMwDGjqeG.PtrToThis = 0

	fMwDGjqeG.Hash =  /*line :1*/ jSX7rdi(sw8_lJ.t.Hash, '*')

	fMwDGjqeG.Elem = eU5PGA6aD3

	p4rJlS0a8, _ :=  /*line :1*/ vGn3oAB.LoadOrStore(sw8_lJ, &fMwDGjqeG)
	return &p4rJlS0a8.(*uSkasB).Type
}

func v8dZcSqC(sw8_lJ *abi.Type) *abi.Type {
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ).v8dZcSqC()
}

func jSX7rdi(uiTlN8bdKm uint32, o3GTtRs6OSZx ...byte) uint32 {
	for _, w1Ndp0ySZgL := range o3GTtRs6OSZx {
		uiTlN8bdKm = uiTlN8bdKm*16777619 ^  /*line :1*/ uint32(w1Ndp0ySZgL)
	}
	return uiTlN8bdKm
}

func (sw8_lJ *rtype) Implements(iaoHrRg1 YJh989LTZsVX) bool {
	if iaoHrRg1 == nil {
		 /*line :1*/ panic("reflect: nil type passed to Type.Implements")
	}
	if  /*line :1*/ iaoHrRg1.Kind() != Interface {
		 /*line :1*/ panic("reflect: non-interface type passed to Type.Implements")
	}
	return  /*line :1*/ jih7qN7qX( /*line :1*/ iaoHrRg1.z6hGxTABM1(),  /*line :1*/ sw8_lJ.z6hGxTABM1())
}

func (sw8_lJ *rtype) AssignableTo(iaoHrRg1 YJh989LTZsVX) bool {
	if iaoHrRg1 == nil {
		 /*line :1*/ panic("reflect: nil type passed to Type.AssignableTo")
	}
	tjMS5cLB :=  /*line :1*/ iaoHrRg1.z6hGxTABM1()
	return  /*line :1*/ vE0PYJd(tjMS5cLB,  /*line :1*/ sw8_lJ.z6hGxTABM1()) ||  /*line :1*/ jih7qN7qX(tjMS5cLB,  /*line :1*/ sw8_lJ.z6hGxTABM1())
}

func (sw8_lJ *rtype) ConvertibleTo(iaoHrRg1 YJh989LTZsVX) bool {
	if iaoHrRg1 == nil {
		 /*line :1*/ panic("reflect: nil type passed to Type.ConvertibleTo")
	}
	return  /*line :1*/ kamZNfVk4( /*line :1*/ iaoHrRg1.z6hGxTABM1(),  /*line :1*/ sw8_lJ.z6hGxTABM1()) != nil
}

func (sw8_lJ *rtype) Comparable() bool {
	return sw8_lJ.t.Equal != nil
}

func jih7qN7qX(Npbez3e, WkkmLaC *abi.Type) bool {
	if  /*line :1*/ Npbez3e.Kind() != abi.Interface {
		return false
	}
	sw8_lJ := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(Npbez3e))
	if  /*line :1*/ len(sw8_lJ.Methods) == 0 {
		return true
	}

	if  /*line :1*/ WkkmLaC.Kind() == abi.Interface {
		f8NmcatRx := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(WkkmLaC))
		hXZpj1nTZ := 0
		for kNVeJaGxs := 0; kNVeJaGxs <  /*line :1*/ len(f8NmcatRx.Methods); kNVeJaGxs++ {
			nHj0s3K1 := &sw8_lJ.Methods[hXZpj1nTZ]
			r4vNyBqefO :=  /*line :1*/ sw8_lJ.zPya6B(nHj0s3K1.Name)
			wCfZwnJXY := &f8NmcatRx.Methods[kNVeJaGxs]
			falD83Sz :=  /*line :1*/ drswiOv(WkkmLaC, wCfZwnJXY.Name)
			if  /*line :1*/ falD83Sz.Name() ==  /*line :1*/ r4vNyBqefO.Name() &&  /*line :1*/ nsOKvPP(WkkmLaC, wCfZwnJXY.Typ) ==  /*line :1*/ sw8_lJ.x9CINy8a(nHj0s3K1.Typ) {
				if ! /*line :1*/ r4vNyBqefO.IsExported() {
					fAQ7bD9Mp5b :=  /*line :1*/ zIgMu_ZY(r4vNyBqefO)
					if fAQ7bD9Mp5b == "" {
						fAQ7bD9Mp5b =  /*line :1*/ sw8_lJ.PkgPath.Name()
					}
					rlkF3KYFaeuC :=  /*line :1*/ zIgMu_ZY(falD83Sz)
					if rlkF3KYFaeuC == "" {
						rlkF3KYFaeuC =  /*line :1*/ f8NmcatRx.PkgPath.Name()
					}
					if fAQ7bD9Mp5b != rlkF3KYFaeuC {
						continue
					}
				}
				if hXZpj1nTZ++; hXZpj1nTZ >=  /*line :1*/ len(sw8_lJ.Methods) {
					return true
				}
			}
		}
		return false
	}

	f8NmcatRx :=  /*line :1*/ WkkmLaC.Uncommon()
	if f8NmcatRx == nil {
		return false
	}
	hXZpj1nTZ := 0
	xmUI7Q :=  /*line :1*/ f8NmcatRx.Methods()
	for kNVeJaGxs := 0; kNVeJaGxs <  /*line :1*/ int(f8NmcatRx.Mcount); kNVeJaGxs++ {
		nHj0s3K1 := &sw8_lJ.Methods[hXZpj1nTZ]
		r4vNyBqefO :=  /*line :1*/ sw8_lJ.zPya6B(nHj0s3K1.Name)
		wCfZwnJXY := xmUI7Q[kNVeJaGxs]
		falD83Sz :=  /*line :1*/ drswiOv(WkkmLaC, wCfZwnJXY.Name)
		if  /*line :1*/ falD83Sz.Name() ==  /*line :1*/ r4vNyBqefO.Name() &&  /*line :1*/ nsOKvPP(WkkmLaC, wCfZwnJXY.Mtyp) ==  /*line :1*/ sw8_lJ.x9CINy8a(nHj0s3K1.Typ) {
			if ! /*line :1*/ r4vNyBqefO.IsExported() {
				fAQ7bD9Mp5b :=  /*line :1*/ zIgMu_ZY(r4vNyBqefO)
				if fAQ7bD9Mp5b == "" {
					fAQ7bD9Mp5b =  /*line :1*/ sw8_lJ.PkgPath.Name()
				}
				rlkF3KYFaeuC :=  /*line :1*/ zIgMu_ZY(falD83Sz)
				if rlkF3KYFaeuC == "" {
					rlkF3KYFaeuC =  /*line :1*/ drswiOv(WkkmLaC, f8NmcatRx.PkgPath).Name()
				}
				if fAQ7bD9Mp5b != rlkF3KYFaeuC {
					continue
				}
			}
			if hXZpj1nTZ++; hXZpj1nTZ >=  /*line :1*/ len(sw8_lJ.Methods) {
				return true
			}
		}
	}
	return false
}

func cY0jZX(Npbez3e, WkkmLaC *abi.Type) bool {

	return  /*line :1*/ WkkmLaC.ChanDir() == abi.BothDir && ( /*line :1*/ jv91ZkF(Npbez3e) == "" ||  /*line :1*/ jv91ZkF(WkkmLaC) == "") &&  /*line :1*/ h8u1a8i9( /*line :1*/ Npbez3e.Elem(),  /*line :1*/ WkkmLaC.Elem(), true)
}

func vE0PYJd(Npbez3e, WkkmLaC *abi.Type) bool {

	if Npbez3e == WkkmLaC {
		return true
	}

	if  /*line :1*/ Npbez3e.HasName() &&  /*line :1*/ WkkmLaC.HasName() ||  /*line :1*/ Npbez3e.Kind() !=  /*line :1*/ WkkmLaC.Kind() {
		return false
	}

	if  /*line :1*/ Npbez3e.Kind() == abi.Chan &&  /*line :1*/ cY0jZX(Npbez3e, WkkmLaC) {
		return true
	}

	return  /*line :1*/ sr7apHgO(Npbez3e, WkkmLaC, true)
}

func h8u1a8i9(Npbez3e, WkkmLaC *abi.Type, iaQUpk bool) bool {
	if iaQUpk {
		return Npbez3e == WkkmLaC
	}

	if  /*line :1*/ jv91ZkF(Npbez3e) !=  /*line :1*/ jv91ZkF(WkkmLaC) ||  /*line :1*/ Npbez3e.Kind() !=  /*line :1*/ WkkmLaC.Kind() ||  /*line :1*/ hP5IxCk8Hfi(Npbez3e) !=  /*line :1*/ hP5IxCk8Hfi(WkkmLaC) {
		return false
	}

	return  /*line :1*/ sr7apHgO(Npbez3e, WkkmLaC, false)
}

func sr7apHgO(Npbez3e, WkkmLaC *abi.Type, iaQUpk bool) bool {
	if Npbez3e == WkkmLaC {
		return true
	}

	zRPbQda4dT :=  /*line :1*/ WzYBjsQL0( /*line :1*/ Npbez3e.Kind())
	if zRPbQda4dT !=  /*line :1*/ WzYBjsQL0( /*line :1*/ WkkmLaC.Kind()) {
		return false
	}

	if Bool <= zRPbQda4dT && zRPbQda4dT <= Complex128 || zRPbQda4dT == String || zRPbQda4dT == UnsafePointer {
		return true
	}

	switch zRPbQda4dT {
	case Array:
		return  /*line :1*/ Npbez3e.Len() ==  /*line :1*/ WkkmLaC.Len() &&  /*line :1*/ h8u1a8i9( /*line :1*/ Npbez3e.Elem(),  /*line :1*/ WkkmLaC.Elem(), iaQUpk)

	case Chan:
		return  /*line :1*/ WkkmLaC.ChanDir() ==  /*line :1*/ Npbez3e.ChanDir() &&  /*line :1*/ h8u1a8i9( /*line :1*/ Npbez3e.Elem(),  /*line :1*/ WkkmLaC.Elem(), iaQUpk)

	case Func:
		sw8_lJ := (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer(Npbez3e))
		f8NmcatRx := (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer(WkkmLaC))
		if sw8_lJ.OutCount != f8NmcatRx.OutCount || sw8_lJ.InCount != f8NmcatRx.InCount {
			return false
		}
		for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ sw8_lJ.NumIn(); hXZpj1nTZ++ {
			if ! /*line :1*/ h8u1a8i9( /*line :1*/ sw8_lJ.In(hXZpj1nTZ),  /*line :1*/ f8NmcatRx.In(hXZpj1nTZ), iaQUpk) {
				return false
			}
		}
		for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ sw8_lJ.NumOut(); hXZpj1nTZ++ {
			if ! /*line :1*/ h8u1a8i9( /*line :1*/ sw8_lJ.Out(hXZpj1nTZ),  /*line :1*/ f8NmcatRx.Out(hXZpj1nTZ), iaQUpk) {
				return false
			}
		}
		return true

	case Interface:
		sw8_lJ := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(Npbez3e))
		f8NmcatRx := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(WkkmLaC))
		if  /*line :1*/ len(sw8_lJ.Methods) == 0 &&  /*line :1*/ len(f8NmcatRx.Methods) == 0 {
			return true
		}

		return false

	case Map:
		return  /*line :1*/ h8u1a8i9( /*line :1*/ Npbez3e.Key(),  /*line :1*/ WkkmLaC.Key(), iaQUpk) &&  /*line :1*/ h8u1a8i9( /*line :1*/ Npbez3e.Elem(),  /*line :1*/ WkkmLaC.Elem(), iaQUpk)

	case Pointer, Slice:
		return  /*line :1*/ h8u1a8i9( /*line :1*/ Npbez3e.Elem(),  /*line :1*/ WkkmLaC.Elem(), iaQUpk)

	case Struct:
		sw8_lJ := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(Npbez3e))
		f8NmcatRx := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(WkkmLaC))
		if  /*line :1*/ len(sw8_lJ.Fields) !=  /*line :1*/ len(f8NmcatRx.Fields) {
			return false
		}
		if  /*line :1*/ sw8_lJ.PkgPath.Name() !=  /*line :1*/ f8NmcatRx.PkgPath.Name() {
			return false
		}
		for hXZpj1nTZ := range sw8_lJ.Fields {
			ccFQQ9kt6M := &sw8_lJ.Fields[hXZpj1nTZ]
			cwCqsRbUg07 := &f8NmcatRx.Fields[hXZpj1nTZ]
			if  /*line :1*/ ccFQQ9kt6M.Name.Name() !=  /*line :1*/ cwCqsRbUg07.Name.Name() {
				return false
			}
			if ! /*line :1*/ h8u1a8i9(ccFQQ9kt6M.Typ, cwCqsRbUg07.Typ, iaQUpk) {
				return false
			}
			if iaQUpk &&  /*line :1*/ ccFQQ9kt6M.Name.Tag() !=  /*line :1*/ cwCqsRbUg07.Name.Tag() {
				return false
			}
			if ccFQQ9kt6M.Offset != cwCqsRbUg07.Offset {
				return false
			}
			if  /*line :1*/ ccFQQ9kt6M.Embedded() !=  /*line :1*/ cwCqsRbUg07.Embedded() {
				return false
			}
		}
		return true
	}

	return false
}

func kCwepmTZ_Y() (ahhvW3ggCTto []unsafe.Pointer, dcDpXllPZ3e [][]int32)

func sxTPqkCA(m6TadEe unsafe.Pointer, e65m4G53plB int32) *abi.Type {
	return (* /*line :1*/ abi.Type)( /*line :1*/ lxkP7oV7sm(m6TadEe,  /*line :1*/ uintptr(e65m4G53plB), "sizeof(rtype) > 0"))
}

func b0uDpEzf(kXTvhUI4Zgg5 string) []*abi.Type {
	ahhvW3ggCTto, dcDpXllPZ3e :=  /*line :1*/ kCwepmTZ_Y()
	var aLsdcpyP1 []*abi.Type

	for i4aROI, yiYpPVJQ9 := range dcDpXllPZ3e {
		m6TadEe := ahhvW3ggCTto[i4aROI]

		hXZpj1nTZ, kNVeJaGxs := 0,  /*line :1*/ len(yiYpPVJQ9)
		for hXZpj1nTZ < kNVeJaGxs {
			cJS_wcOmyXDz := hXZpj1nTZ + (kNVeJaGxs-hXZpj1nTZ)>>1

			if !( /*line :1*/ iR_1JLd( /*line :1*/ sxTPqkCA(m6TadEe, yiYpPVJQ9[cJS_wcOmyXDz])) >= kXTvhUI4Zgg5) {
				hXZpj1nTZ = cJS_wcOmyXDz + 1
			} else {
				kNVeJaGxs = cJS_wcOmyXDz
			}
		}

		for kNVeJaGxs := hXZpj1nTZ; kNVeJaGxs <  /*line :1*/ len(yiYpPVJQ9); kNVeJaGxs++ {
			lm5DH3 :=  /*line :1*/ sxTPqkCA(m6TadEe, yiYpPVJQ9[kNVeJaGxs])
			if  /*line :1*/ iR_1JLd(lm5DH3) != kXTvhUI4Zgg5 {
				break
			}
			aLsdcpyP1 =  /*line :1*/ append(aLsdcpyP1, lm5DH3)
		}
	}
	return aLsdcpyP1
}

var hDZJaY1mZDEB sync.Eim1b6bNi9IM

type kAM7Km9_stw struct {
	in0YruWwmIdi	WzYBjsQL0
	a8eOcNE1V9j	*abi.Type
	laacKxI_a	*abi.Type
	ffPvP0M5j	uintptr
}

var zS62zNBmicXB struct {
	sync.DIRWe11KYlYa

	mzTN8ya	sync.Eim1b6bNi9IM
}

func IURyE5tE22i(aY6MVP Svj3YDOP3fa, sw8_lJ YJh989LTZsVX) YJh989LTZsVX {
	lm5DH3 :=  /*line :1*/ sw8_lJ.z6hGxTABM1()

	ctW2F7ER := kAM7Km9_stw{Chan, lm5DH3, nil,  /*line :1*/ uintptr(aY6MVP)}
	if qqxx1RLlUhKF, bDfzXlCm5Raf :=  /*line :1*/ hDZJaY1mZDEB.Load(ctW2F7ER); bDfzXlCm5Raf {
		return qqxx1RLlUhKF.(*rtype)
	}

	if lm5DH3.Size_ >= 1<<16 {
		 /*line :1*/ panic("reflect.ChanOf: element size too large")
	}

	var kXTvhUI4Zgg5 string
	switch aY6MVP {
	default:
		 /*line :1*/ panic("reflect.ChanOf: invalid dir")
	case SendDir:
		kXTvhUI4Zgg5 = "chan<- " +  /*line :1*/ iR_1JLd(lm5DH3)
	case RecvDir:
		kXTvhUI4Zgg5 = "<-chan " +  /*line :1*/ iR_1JLd(lm5DH3)
	case BothDir:
		f7UnAIDCKTG :=  /*line :1*/ iR_1JLd(lm5DH3)
		if f7UnAIDCKTG[0] == '<' {

			kXTvhUI4Zgg5 = "chan (" + f7UnAIDCKTG + ")"
		} else {
			kXTvhUI4Zgg5 = "chan " + f7UnAIDCKTG
		}
	}
	for _, mP7QqSVInV := range  /*line :1*/ b0uDpEzf(kXTvhUI4Zgg5) {
		qqxx1RLlUhKF := (* /*line :1*/ aK0ovNmjQBD)( /*line :1*/ unsafe.Pointer(mP7QqSVInV))
		if qqxx1RLlUhKF.Elem == lm5DH3 && qqxx1RLlUhKF.Dir ==  /*line :1*/ abi.ChanDir(aY6MVP) {
			hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(mP7QqSVInV))
			return hhIBKzzDq.(YJh989LTZsVX)
		}
	}

	var iEzymoD any = (chan  /*line :1*/ unsafe.Pointer)(nil)
	fNjFa2jLl := *(** /*line :1*/ aK0ovNmjQBD)( /*line :1*/ unsafe.Pointer(&iEzymoD))
	qqxx1RLlUhKF := *fNjFa2jLl
	qqxx1RLlUhKF.TFlag = abi.TFlagRegularMemory
	qqxx1RLlUhKF.Dir =  /*line :1*/ abi.ChanDir(aY6MVP)
	qqxx1RLlUhKF.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(kXTvhUI4Zgg5, "", false, false))
	qqxx1RLlUhKF.Hash =  /*line :1*/ jSX7rdi(lm5DH3.Hash, 'c',  /*line :1*/ byte(aY6MVP))
	qqxx1RLlUhKF.Elem = lm5DH3

	hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(&qqxx1RLlUhKF.Type))
	return hhIBKzzDq.(YJh989LTZsVX)
}

func OULRy5(ccj6oONwQVN, dPe0EQn YJh989LTZsVX) YJh989LTZsVX {
	hQSrLm60DS :=  /*line :1*/ ccj6oONwQVN.z6hGxTABM1()
	uR1VfuHRj :=  /*line :1*/ dPe0EQn.z6hGxTABM1()

	if hQSrLm60DS.Equal == nil {
		 /*line :1*/ panic("reflect.MapOf: invalid key type " +  /*line :1*/ iR_1JLd(hQSrLm60DS))
	}

	ctW2F7ER := kAM7Km9_stw{Map, hQSrLm60DS, uR1VfuHRj, 0}
	if byRNfAUjb, bDfzXlCm5Raf :=  /*line :1*/ hDZJaY1mZDEB.Load(ctW2F7ER); bDfzXlCm5Raf {
		return byRNfAUjb.(YJh989LTZsVX)
	}

	kXTvhUI4Zgg5 := "map[" +  /*line :1*/ iR_1JLd(hQSrLm60DS) + "]" +  /*line :1*/ iR_1JLd(uR1VfuHRj)
	for _, mP7QqSVInV := range  /*line :1*/ b0uDpEzf(kXTvhUI4Zgg5) {
		byRNfAUjb := (* /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer(mP7QqSVInV))
		if byRNfAUjb.Key == hQSrLm60DS && byRNfAUjb.Elem == uR1VfuHRj {
			hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(mP7QqSVInV))
			return hhIBKzzDq.(YJh989LTZsVX)
		}
	}

	var wrsWqyIVbr85 any = (map[ /*line :1*/ unsafe.Pointer]unsafe.Pointer)(nil)
	byRNfAUjb := **(** /*line :1*/ r0LDmUgzY)( /*line :1*/ unsafe.Pointer(&wrsWqyIVbr85))
	byRNfAUjb.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(kXTvhUI4Zgg5, "", false, false))
	byRNfAUjb.TFlag = 0
	byRNfAUjb.Hash =  /*line :1*/ jSX7rdi(uR1VfuHRj.Hash, 'm',  /*line :1*/ byte(hQSrLm60DS.Hash>>24),  /*line :1*/ byte(hQSrLm60DS.Hash>>16),  /*line :1*/ byte(hQSrLm60DS.Hash>>8),  /*line :1*/ byte(hQSrLm60DS.Hash))
	byRNfAUjb.Key = hQSrLm60DS
	byRNfAUjb.Elem = uR1VfuHRj
	byRNfAUjb.Bucket =  /*line :1*/ udG1GgUDRgd(hQSrLm60DS, uR1VfuHRj)
	byRNfAUjb.Hasher = func(e3IyyaQSEj unsafe.Pointer, w4w8FnHnv uintptr) uintptr {
		return  /*line :1*/ hmksVe18ic(hQSrLm60DS, e3IyyaQSEj, w4w8FnHnv)
	}
	byRNfAUjb.Flags = 0
	if hQSrLm60DS.Size_ > maxKeySize {
		byRNfAUjb.KeySize =  /*line :1*/ uint8(goarch.PtrSize)
		byRNfAUjb.Flags |= 1
	} else {
		byRNfAUjb.KeySize =  /*line :1*/ uint8(hQSrLm60DS.Size_)
	}
	if uR1VfuHRj.Size_ > maxValSize {
		byRNfAUjb.ValueSize =  /*line :1*/ uint8(goarch.PtrSize)
		byRNfAUjb.Flags |= 2
	} else {
		byRNfAUjb.MapType.ValueSize =  /*line :1*/ uint8(uR1VfuHRj.Size_)
	}
	byRNfAUjb.MapType.BucketSize =  /*line :1*/ uint16(byRNfAUjb.Bucket.Size_)
	if  /*line :1*/ uYDlpZAC0Q8(hQSrLm60DS) {
		byRNfAUjb.Flags |= 4
	}
	if  /*line :1*/ gyaiDBpc(hQSrLm60DS) {
		byRNfAUjb.Flags |= 8
	}
	if  /*line :1*/ zXqkkU6G(hQSrLm60DS) {
		byRNfAUjb.Flags |= 16
	}
	byRNfAUjb.PtrToThis = 0

	hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(&byRNfAUjb.Type))
	return hhIBKzzDq.(YJh989LTZsVX)
}

var gPn5f35kwFh []YJh989LTZsVX
var hravW3g sync.DIRWe11KYlYa

func mhZPrOT3svlO(krtR1HwEan int) YJh989LTZsVX {
	 /*line :1*/ hravW3g.Lock()
	defer  /*line :1*/ hravW3g.Unlock()
	if krtR1HwEan >=  /*line :1*/ len(gPn5f35kwFh) {
		v9VQ9pey1CyC :=  /*line :1*/ make([]YJh989LTZsVX, krtR1HwEan+1)
		 /*line :1*/ copy(v9VQ9pey1CyC, gPn5f35kwFh)
		gPn5f35kwFh = v9VQ9pey1CyC
	}
	if gPn5f35kwFh[krtR1HwEan] != nil {
		return gPn5f35kwFh[krtR1HwEan]
	}

	gPn5f35kwFh[krtR1HwEan] =  /*line :1*/ FNkNFKGvwkzw([]EeSeBImJjDQo{
		{
			ZOSBHkJz:	"FuncType",
			S1miDNBAV8a:	 /*line :1*/ Cher1a2Fblr(bY2Zo8pis{}),
		},
		{
			ZOSBHkJz:	"Args",
			S1miDNBAV8a:	 /*line :1*/ PNcEGf7fzu(krtR1HwEan,  /*line :1*/ Cher1a2Fblr(&rtype{})),
		},
	})
	return gPn5f35kwFh[krtR1HwEan]
}

func NlPxWAGa0V(gJ3XGL_F, hEbpexD []YJh989LTZsVX, zj45Lc bool) YJh989LTZsVX {
	if zj45Lc && ( /*line :1*/ len(gJ3XGL_F) == 0 ||  /*line :1*/ gJ3XGL_F[ /*line :1*/ len(gJ3XGL_F)-1].Kind() != Slice) {
		 /*line :1*/ panic("reflect.FuncOf: last arg of variadic func must be slice")
	}

	var yiD2tsESqU any = (func())( /*line :1*/ nil)
	fNjFa2jLl := *(** /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer(&yiD2tsESqU))
	krtR1HwEan :=  /*line :1*/ len(gJ3XGL_F) +  /*line :1*/ len(hEbpexD)

	if krtR1HwEan > 128 {
		 /*line :1*/ panic("reflect.FuncOf: too many arguments")
	}

	uKZ9OYD :=  /*line :1*/ L6EAuajfYe1( /*line :1*/ mhZPrOT3svlO(krtR1HwEan)).Elem()
	yzEE5FI := (* /*line :1*/ bY2Zo8pis)( /*line :1*/ unsafe.Pointer( /*line :1*/ uKZ9OYD.Field(0).Addr().Pointer()))
	fqjeuS :=  /*line :1*/ unsafe.Slice((** /*line :1*/ rtype)( /*line :1*/ unsafe.Pointer( /*line :1*/ uKZ9OYD.Field(1).Addr().Pointer())), krtR1HwEan)[0:0:krtR1HwEan]
	*yzEE5FI = *fNjFa2jLl

	var x6CWj5lQL uint32
	for _, gJ3XGL_F := range gJ3XGL_F {
		sw8_lJ := gJ3XGL_F.(*rtype)
		fqjeuS =  /*line :1*/ append(fqjeuS, sw8_lJ)
		x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL,  /*line :1*/ byte(sw8_lJ.t.Hash>>24),  /*line :1*/ byte(sw8_lJ.t.Hash>>16),  /*line :1*/ byte(sw8_lJ.t.Hash>>8),  /*line :1*/ byte(sw8_lJ.t.Hash))
	}
	if zj45Lc {
		x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL, 'v')
	}
	x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL, '.')
	for _, hEbpexD := range hEbpexD {
		sw8_lJ := hEbpexD.(*rtype)
		fqjeuS =  /*line :1*/ append(fqjeuS, sw8_lJ)
		x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL,  /*line :1*/ byte(sw8_lJ.t.Hash>>24),  /*line :1*/ byte(sw8_lJ.t.Hash>>16),  /*line :1*/ byte(sw8_lJ.t.Hash>>8),  /*line :1*/ byte(sw8_lJ.t.Hash))
	}

	yzEE5FI.TFlag = 0
	yzEE5FI.Hash = x6CWj5lQL
	yzEE5FI.InCount =  /*line :1*/ uint16( /*line :1*/ len(gJ3XGL_F))
	yzEE5FI.OutCount =  /*line :1*/ uint16( /*line :1*/ len(hEbpexD))
	if zj45Lc {
		yzEE5FI.OutCount |= 1 << 15
	}

	if ibdpoF, bDfzXlCm5Raf :=  /*line :1*/ zS62zNBmicXB.mzTN8ya.Load(x6CWj5lQL); bDfzXlCm5Raf {
		for _, sw8_lJ := range ibdpoF.([]*abi.Type) {
			if  /*line :1*/ sr7apHgO(&yzEE5FI.Type, sw8_lJ, true) {
				return  /*line :1*/ j5Z_WRWLudek(sw8_lJ)
			}
		}
	}

	 /*line :1*/ zS62zNBmicXB.Lock()
	defer  /*line :1*/ zS62zNBmicXB.Unlock()
	if ibdpoF, bDfzXlCm5Raf :=  /*line :1*/ zS62zNBmicXB.mzTN8ya.Load(x6CWj5lQL); bDfzXlCm5Raf {
		for _, sw8_lJ := range ibdpoF.([]*abi.Type) {
			if  /*line :1*/ sr7apHgO(&yzEE5FI.Type, sw8_lJ, true) {
				return  /*line :1*/ j5Z_WRWLudek(sw8_lJ)
			}
		}
	}

	skXTsm0V := func(mP7QqSVInV *abi.Type) YJh989LTZsVX {
		var ytIR1Q6znM5B []*abi.Type
		if feIflohGIyX, bDfzXlCm5Raf :=  /*line :1*/ zS62zNBmicXB.mzTN8ya.Load(x6CWj5lQL); bDfzXlCm5Raf {
			ytIR1Q6znM5B = feIflohGIyX.([]*abi.Type)
		}
		 /*line :1*/ zS62zNBmicXB.mzTN8ya.Store(x6CWj5lQL,  /*line :1*/ append(ytIR1Q6znM5B, mP7QqSVInV))
		return  /*line :1*/ nadfj4Ka8cQf(mP7QqSVInV)
	}

	oQ6fAx :=  /*line :1*/ qMpSAaE_(yzEE5FI)
	for _, mP7QqSVInV := range  /*line :1*/ b0uDpEzf(oQ6fAx) {
		if  /*line :1*/ sr7apHgO(&yzEE5FI.Type, mP7QqSVInV, true) {
			return  /*line :1*/ skXTsm0V(mP7QqSVInV)
		}
	}

	yzEE5FI.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(oQ6fAx, "", false, false))
	yzEE5FI.PtrToThis = 0
	return  /*line :1*/ skXTsm0V(&yzEE5FI.Type)
}
func iR_1JLd(sw8_lJ *abi.Type) string {
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ).String()
}

func qMpSAaE_(yzEE5FI *bY2Zo8pis) string {
	sfIHyhaL :=  /*line :1*/ make([]byte, 0, 64)
	sfIHyhaL =  /*line :1*/ append(sfIHyhaL, "func("...)
	for hXZpj1nTZ, sw8_lJ := range  /*line :1*/ yzEE5FI.InSlice() {
		if hXZpj1nTZ > 0 {
			sfIHyhaL =  /*line :1*/ append(sfIHyhaL, ", "...)
		}
		if  /*line :1*/ yzEE5FI.IsVariadic() && hXZpj1nTZ ==  /*line :1*/ int(yzEE5FI.InCount)-1 {
			sfIHyhaL =  /*line :1*/ append(sfIHyhaL, "..."...)
			sfIHyhaL =  /*line :1*/ append(sfIHyhaL,  /*line :1*/ iR_1JLd((* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer(sw8_lJ)).Elem)...)
		} else {
			sfIHyhaL =  /*line :1*/ append(sfIHyhaL,  /*line :1*/ iR_1JLd(sw8_lJ)...)
		}
	}
	sfIHyhaL =  /*line :1*/ append(sfIHyhaL, ')')
	hEbpexD :=  /*line :1*/ yzEE5FI.OutSlice()
	if  /*line :1*/ len(hEbpexD) == 1 {
		sfIHyhaL =  /*line :1*/ append(sfIHyhaL, ' ')
	} else if  /*line :1*/ len(hEbpexD) > 1 {
		sfIHyhaL =  /*line :1*/ append(sfIHyhaL, " ("...)
	}
	for hXZpj1nTZ, sw8_lJ := range hEbpexD {
		if hXZpj1nTZ > 0 {
			sfIHyhaL =  /*line :1*/ append(sfIHyhaL, ", "...)
		}
		sfIHyhaL =  /*line :1*/ append(sfIHyhaL,  /*line :1*/ iR_1JLd(sw8_lJ)...)
	}
	if  /*line :1*/ len(hEbpexD) > 1 {
		sfIHyhaL =  /*line :1*/ append(sfIHyhaL, ')')
	}
	return  /*line :1*/ string(sfIHyhaL)
}

func uYDlpZAC0Q8(sw8_lJ *abi.Type) bool {
	switch  /*line :1*/ WzYBjsQL0( /*line :1*/ sw8_lJ.Kind()) {
	case Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr, Chan, Pointer, String, UnsafePointer:
		return true
	case Float32, Float64, Complex64, Complex128, Interface:
		return false
	case Array:
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		return  /*line :1*/ uYDlpZAC0Q8(mP7QqSVInV.Elem)
	case Struct:
		mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		for _, jToThzM := range mP7QqSVInV.Fields {
			if ! /*line :1*/ uYDlpZAC0Q8(jToThzM.Typ) {
				return false
			}
		}
		return true
	default:

		 /*line :1*/ panic("isReflexive called on non-key type " +  /*line :1*/ iR_1JLd(sw8_lJ))
	}
}

func gyaiDBpc(sw8_lJ *abi.Type) bool {
	switch  /*line :1*/ WzYBjsQL0( /*line :1*/ sw8_lJ.Kind()) {
	case Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr, Chan, Pointer, UnsafePointer:
		return false
	case Float32, Float64, Complex64, Complex128, Interface, String:

		return true
	case Array:
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		return  /*line :1*/ gyaiDBpc(mP7QqSVInV.Elem)
	case Struct:
		mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		for _, jToThzM := range mP7QqSVInV.Fields {
			if  /*line :1*/ gyaiDBpc(jToThzM.Typ) {
				return true
			}
		}
		return false
	default:

		 /*line :1*/ panic("needKeyUpdate called on non-key type " +  /*line :1*/ iR_1JLd(sw8_lJ))
	}
}

func zXqkkU6G(sw8_lJ *abi.Type) bool {
	switch  /*line :1*/ WzYBjsQL0( /*line :1*/ sw8_lJ.Kind()) {
	case Interface:
		return true
	case Array:
		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		return  /*line :1*/ zXqkkU6G(mP7QqSVInV.Elem)
	case Struct:
		mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		for _, jToThzM := range mP7QqSVInV.Fields {
			if  /*line :1*/ zXqkkU6G(jToThzM.Typ) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

const (
	bucketSize	uintptr	= abi.MapBucketCount
	maxKeySize	uintptr	= abi.MapMaxKeyBytes
	maxValSize	uintptr	= abi.MapMaxElemBytes
)

func udG1GgUDRgd(hQSrLm60DS, uR1VfuHRj *abi.Type) *abi.Type {
	if hQSrLm60DS.Size_ > maxKeySize {
		hQSrLm60DS =  /*line :1*/ v8dZcSqC(hQSrLm60DS)
	}
	if uR1VfuHRj.Size_ > maxValSize {
		uR1VfuHRj =  /*line :1*/ v8dZcSqC(uR1VfuHRj)
	}

	var h9GF73dL *byte
	var ljCRrt88I uintptr

	wnTnByH := bucketSize*(1+hQSrLm60DS.Size_+uR1VfuHRj.Size_) + goarch.PtrSize
	if wnTnByH& /*line :1*/ uintptr(hQSrLm60DS.Align_-1) != 0 || wnTnByH& /*line :1*/ uintptr(uR1VfuHRj.Align_-1) != 0 {
		 /*line :1*/ panic("reflect: bad size computation in MapOf")
	}

	if hQSrLm60DS.PtrBytes != 0 || uR1VfuHRj.PtrBytes != 0 {
		_4u93bgKl0W := (bucketSize*(1+hQSrLm60DS.Size_+uR1VfuHRj.Size_) + goarch.PtrSize) / goarch.PtrSize
		krtR1HwEan := (_4u93bgKl0W + 7) / 8

		krtR1HwEan = (krtR1HwEan + goarch.PtrSize - 1) &^ (goarch.PtrSize - 1)
		hpoWZKhOa0 :=  /*line :1*/ make([]byte, krtR1HwEan)
		tDufzwDs061 := bucketSize / goarch.PtrSize

		if hQSrLm60DS.PtrBytes != 0 {
			 /*line :1*/ r1PEsohea(hpoWZKhOa0, tDufzwDs061, hQSrLm60DS, bucketSize)
		}
		tDufzwDs061 += bucketSize * hQSrLm60DS.Size_ / goarch.PtrSize

		if uR1VfuHRj.PtrBytes != 0 {
			 /*line :1*/ r1PEsohea(hpoWZKhOa0, tDufzwDs061, uR1VfuHRj, bucketSize)
		}
		tDufzwDs061 += bucketSize * uR1VfuHRj.Size_ / goarch.PtrSize

		zdeUTwtpz := tDufzwDs061
		hpoWZKhOa0[zdeUTwtpz/8] |= 1 << (zdeUTwtpz % 8)
		h9GF73dL = &hpoWZKhOa0[0]
		ljCRrt88I = (zdeUTwtpz + 1) * goarch.PtrSize

		if ljCRrt88I != wnTnByH {
			 /*line :1*/ panic("reflect: bad layout computation in MapOf")
		}
	}

	w1Ndp0ySZgL := &abi.Type{
		Align_:	goarch.PtrSize,
		Size_:	wnTnByH,
		Kind_:	 /*line :1*/ uint8(Struct),
		PtrBytes:	ljCRrt88I,
		GCData:	h9GF73dL,
	}
	kXTvhUI4Zgg5 := "bucket(" +  /*line :1*/ iR_1JLd(hQSrLm60DS) + "," +  /*line :1*/ iR_1JLd(uR1VfuHRj) + ")"
	w1Ndp0ySZgL.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(kXTvhUI4Zgg5, "", false, false))
	return w1Ndp0ySZgL
}

func (sw8_lJ *rtype) pRMQmn9v5r(j9W45J, bEbyW_Zmxi uintptr) []byte {
	return (*[1 << 30] /*line :1*/ byte)( /*line :1*/ unsafe.Pointer(sw8_lJ.t.GCData))[j9W45J:bEbyW_Zmxi:bEbyW_Zmxi]
}

func r1PEsohea(hEbpexD []byte, tDufzwDs061 uintptr, lm5DH3 *abi.Type, krtR1HwEan uintptr) {
	if lm5DH3.Kind_&kindGCProg != 0 {
		 /*line :1*/ panic("reflect: unexpected GC program")
	}
	f6BNaoU02hk_ := lm5DH3.PtrBytes / goarch.PtrSize
	vnjOsvRWNE := lm5DH3.Size_ / goarch.PtrSize
	hpoWZKhOa0 :=  /*line :1*/ lm5DH3.GcSlice(0, (f6BNaoU02hk_+7)/8)
	for kNVeJaGxs :=  /*line :1*/ uintptr(0); kNVeJaGxs < f6BNaoU02hk_; kNVeJaGxs++ {
		if (hpoWZKhOa0[kNVeJaGxs/8]>>(kNVeJaGxs%8))&1 != 0 {
			for hXZpj1nTZ :=  /*line :1*/ uintptr(0); hXZpj1nTZ < krtR1HwEan; hXZpj1nTZ++ {
				lYOqUQga := tDufzwDs061 + hXZpj1nTZ*vnjOsvRWNE + kNVeJaGxs
				hEbpexD[lYOqUQga/8] |= 1 << (lYOqUQga % 8)
			}
		}
	}
}

func c7bREF5(w2dYwY0 []byte, lm5DH3 *abi.Type) []byte {
	if lm5DH3.Kind_&kindGCProg != 0 {

		krtR1HwEan :=  /*line :1*/ uintptr(*(* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(lm5DH3.GCData)))
		lanXsCdMk :=  /*line :1*/ lm5DH3.GcSlice(4, 4+krtR1HwEan-1)
		return  /*line :1*/ append(w2dYwY0, lanXsCdMk...)
	}

	f6BNaoU02hk_ := lm5DH3.PtrBytes / goarch.PtrSize
	hpoWZKhOa0 :=  /*line :1*/ lm5DH3.GcSlice(0, (f6BNaoU02hk_+7)/8)

	for ; f6BNaoU02hk_ > 120; f6BNaoU02hk_ -= 120 {
		w2dYwY0 =  /*line :1*/ append(w2dYwY0, 120)
		w2dYwY0 =  /*line :1*/ append(w2dYwY0, hpoWZKhOa0[:15]...)
		hpoWZKhOa0 = hpoWZKhOa0[15:]
	}

	w2dYwY0 =  /*line :1*/ append(w2dYwY0,  /*line :1*/ byte(f6BNaoU02hk_))
	w2dYwY0 =  /*line :1*/ append(w2dYwY0, hpoWZKhOa0...)
	return w2dYwY0
}

func XsHS1gYvDYpj(sw8_lJ YJh989LTZsVX) YJh989LTZsVX {
	lm5DH3 :=  /*line :1*/ sw8_lJ.z6hGxTABM1()

	ctW2F7ER := kAM7Km9_stw{Slice, lm5DH3, nil, 0}
	if qMTDS7fnbc, bDfzXlCm5Raf :=  /*line :1*/ hDZJaY1mZDEB.Load(ctW2F7ER); bDfzXlCm5Raf {
		return qMTDS7fnbc.(YJh989LTZsVX)
	}

	kXTvhUI4Zgg5 := "[]" +  /*line :1*/ iR_1JLd(lm5DH3)
	for _, mP7QqSVInV := range  /*line :1*/ b0uDpEzf(kXTvhUI4Zgg5) {
		qMTDS7fnbc := (* /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer(mP7QqSVInV))
		if qMTDS7fnbc.Elem == lm5DH3 {
			hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(mP7QqSVInV))
			return hhIBKzzDq.(YJh989LTZsVX)
		}
	}

	var gUjwimp9WV any = ([] /*line :1*/ unsafe.Pointer)(nil)
	fNjFa2jLl := *(** /*line :1*/ vwvVHq)( /*line :1*/ unsafe.Pointer(&gUjwimp9WV))
	qMTDS7fnbc := *fNjFa2jLl
	qMTDS7fnbc.TFlag = 0
	qMTDS7fnbc.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(kXTvhUI4Zgg5, "", false, false))
	qMTDS7fnbc.Hash =  /*line :1*/ jSX7rdi(lm5DH3.Hash, '[')
	qMTDS7fnbc.Elem = lm5DH3
	qMTDS7fnbc.PtrToThis = 0

	hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(&qMTDS7fnbc.Type))
	return hhIBKzzDq.(YJh989LTZsVX)
}

var f6wPeOefr struct {
	sync.DIRWe11KYlYa

	mzTN8ya	sync.Eim1b6bNi9IM
}

type structTypeUncommon struct {
	structType
	u	aKCY_EtrgnA8
}

func wJiWN_w0ZI(qqxx1RLlUhKF rune) bool {
	return 'a' <= qqxx1RLlUhKF && qqxx1RLlUhKF <= 'z' || 'A' <= qqxx1RLlUhKF && qqxx1RLlUhKF <= 'Z' || qqxx1RLlUhKF == '_' || qqxx1RLlUhKF >= utf8.RuneSelf &&  /*line :1*/ unicode.Nyu9Lcgu(qqxx1RLlUhKF)
}

func h3ctH78ZLe(zRVzywaHE string) bool {
	for hXZpj1nTZ, i5hAnuqALA := range zRVzywaHE {
		if hXZpj1nTZ == 0 && ! /*line :1*/ wJiWN_w0ZI(i5hAnuqALA) {
			return false
		}

		if !( /*line :1*/ wJiWN_w0ZI(i5hAnuqALA) ||  /*line :1*/ unicode.FquwNEzWV(i5hAnuqALA)) {
			return false
		}
	}

	return  /*line :1*/ len(zRVzywaHE) > 0
}

func FNkNFKGvwkzw(hdC2Rc []EeSeBImJjDQo) YJh989LTZsVX {
	var (
		x6CWj5lQL	=  /*line :1*/ jSX7rdi(0, [] /*line :1*/ byte("struct {")...)
		wnTnByH	uintptr
		aO_2b3G	uint8
		vaI0gqokB	= true
		r2zNEKaY	[]abi.Method

		rH64rFnxgI47	=  /*line :1*/ make([]yUesa1U9,  /*line :1*/ len(hdC2Rc))
		sfIHyhaL	=  /*line :1*/ make([]byte, 0, 64)
		lsockidttb	= map[string]struct{}{}

		e_ELAml11yj	= false
	)

	kraXVtm :=  /*line :1*/ uintptr(0)
	sfIHyhaL =  /*line :1*/ append(sfIHyhaL, "struct {"...)
	enqJZy := ""
	for hXZpj1nTZ, fBabaKImKM2F := range hdC2Rc {
		if fBabaKImKM2F.ZOSBHkJz == "" {
			 /*line :1*/ panic("reflect.StructOf: field " +  /*line :1*/ strconv.ZW1_FxJtq(hXZpj1nTZ) + " has no name")
		}
		if ! /*line :1*/ h3ctH78ZLe(fBabaKImKM2F.ZOSBHkJz) {
			 /*line :1*/ panic("reflect.StructOf: field " +  /*line :1*/ strconv.ZW1_FxJtq(hXZpj1nTZ) + " has invalid name")
		}
		if fBabaKImKM2F.S1miDNBAV8a == nil {
			 /*line :1*/ panic("reflect.StructOf: field " +  /*line :1*/ strconv.ZW1_FxJtq(hXZpj1nTZ) + " has no type")
		}
		jToThzM, uSQTwkYnKqlb :=  /*line :1*/ pxDaXfB5tPl8(fBabaKImKM2F)
		yzEE5FI := jToThzM.Typ
		if yzEE5FI.Kind_&kindGCProg != 0 {
			e_ELAml11yj = true
		}
		if uSQTwkYnKqlb != "" {
			if enqJZy == "" {
				enqJZy = uSQTwkYnKqlb
			} else if enqJZy != uSQTwkYnKqlb {
				 /*line :1*/ panic("reflect.Struct: fields with different PkgPath " + enqJZy + " and " + uSQTwkYnKqlb)
			}
		}

		lGwAXqN4Hb :=  /*line :1*/ jToThzM.Name.Name()
		x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL, [] /*line :1*/ byte(lGwAXqN4Hb)...)
		sfIHyhaL =  /*line :1*/ append(sfIHyhaL, (" " + lGwAXqN4Hb)...)
		if  /*line :1*/ jToThzM.Embedded() {

			if  /*line :1*/ jToThzM.Typ.Kind() == abi.Pointer {

				dPe0EQn :=  /*line :1*/ yzEE5FI.Elem()
				if lYOqUQga :=  /*line :1*/ dPe0EQn.Kind(); lYOqUQga == abi.Pointer || lYOqUQga == abi.Interface {
					 /*line :1*/ panic("reflect.StructOf: illegal embedded field type " +  /*line :1*/ iR_1JLd(yzEE5FI))
				}
			}

			switch  /*line :1*/ WzYBjsQL0( /*line :1*/ jToThzM.Typ.Kind()) {
			case Interface:
				bWlEe3hmP93Q := (* /*line :1*/ ywQcz6y)( /*line :1*/ unsafe.Pointer(yzEE5FI))
				for x4vMtNd, hek9DV := range bWlEe3hmP93Q.Methods {
					if  /*line :1*/ zIgMu_ZY( /*line :1*/ bWlEe3hmP93Q.zPya6B(hek9DV.Name)) != "" {

						 /*line :1*/ panic("reflect: embedded interface with unexported method(s) not implemented")
					}

					var (
						srqll74r	=  /*line :1*/ bWlEe3hmP93Q.x9CINy8a(hek9DV.Typ)
						ywbh3pRWGp	= hXZpj1nTZ
						sxqY4KpR_	= x4vMtNd
						pVePXfqREHol	Value
						bGGxfg	Value
					)

					if yzEE5FI.Kind_&kindDirectIface != 0 {
						bGGxfg =  /*line :1*/ EiBL4fuxoss( /*line :1*/ j5Z_WRWLudek(srqll74r), func(gJ3XGL_F []Value) []Value {
							var fqjeuS []Value
							var wM9XCjNlp = gJ3XGL_F[0]
							if  /*line :1*/ len(gJ3XGL_F) > 1 {
								fqjeuS = gJ3XGL_F[1:]
							}
							return  /*line :1*/ wM9XCjNlp.Field(ywbh3pRWGp).Method(sxqY4KpR_).Call(fqjeuS)
						})
						pVePXfqREHol =  /*line :1*/ EiBL4fuxoss( /*line :1*/ j5Z_WRWLudek(srqll74r), func(gJ3XGL_F []Value) []Value {
							var fqjeuS []Value
							var wM9XCjNlp = gJ3XGL_F[0]
							if  /*line :1*/ len(gJ3XGL_F) > 1 {
								fqjeuS = gJ3XGL_F[1:]
							}
							return  /*line :1*/ wM9XCjNlp.Field(ywbh3pRWGp).Method(sxqY4KpR_).Call(fqjeuS)
						})
					} else {
						bGGxfg =  /*line :1*/ EiBL4fuxoss( /*line :1*/ j5Z_WRWLudek(srqll74r), func(gJ3XGL_F []Value) []Value {
							var fqjeuS []Value
							var wM9XCjNlp = gJ3XGL_F[0]
							if  /*line :1*/ len(gJ3XGL_F) > 1 {
								fqjeuS = gJ3XGL_F[1:]
							}
							return  /*line :1*/ wM9XCjNlp.Field(ywbh3pRWGp).Method(sxqY4KpR_).Call(fqjeuS)
						})
						pVePXfqREHol =  /*line :1*/ EiBL4fuxoss( /*line :1*/ j5Z_WRWLudek(srqll74r), func(gJ3XGL_F []Value) []Value {
							var fqjeuS []Value
							var wM9XCjNlp =  /*line :1*/ XF6ikODQ8h(gJ3XGL_F[0])
							if  /*line :1*/ len(gJ3XGL_F) > 1 {
								fqjeuS = gJ3XGL_F[1:]
							}
							return  /*line :1*/ wM9XCjNlp.Field(ywbh3pRWGp).Method(sxqY4KpR_).Call(fqjeuS)
						})
					}

					r2zNEKaY =  /*line :1*/ append(r2zNEKaY, abi.Method{
						Name:	 /*line :1*/ vybOmhD9hqa5( /*line :1*/ bWlEe3hmP93Q.zPya6B(hek9DV.Name)),
						Mtyp:	 /*line :1*/ j2acIXKpV(srqll74r),
						Ifn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ unsafe.Pointer(&pVePXfqREHol)),
						Tfn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ unsafe.Pointer(&bGGxfg)),
					})
				}
			case Pointer:
				cf3fCV8ayFq := (* /*line :1*/ uSkasB)( /*line :1*/ unsafe.Pointer(yzEE5FI))
				if pxxIZC46 :=  /*line :1*/ cf3fCV8ayFq.Uncommon(); pxxIZC46 != nil {
					if hXZpj1nTZ > 0 && pxxIZC46.Mcount > 0 {

						 /*line :1*/ panic("reflect: embedded type with methods not implemented if type is not first field")
					}
					if  /*line :1*/ len(hdC2Rc) > 1 {
						 /*line :1*/ panic("reflect: embedded type with methods not implemented if there is more than one field")
					}
					for _, hek9DV := range  /*line :1*/ pxxIZC46.Methods() {
						f4pekhjLv :=  /*line :1*/ drswiOv(yzEE5FI, hek9DV.Name)
						if  /*line :1*/ zIgMu_ZY(f4pekhjLv) != "" {

							 /*line :1*/ panic("reflect: embedded interface with unexported method(s) not implemented")
						}
						r2zNEKaY =  /*line :1*/ append(r2zNEKaY, abi.Method{
							Name:	 /*line :1*/ vybOmhD9hqa5(f4pekhjLv),
							Mtyp:	 /*line :1*/ j2acIXKpV( /*line :1*/ nsOKvPP(yzEE5FI, hek9DV.Mtyp)),
							Ifn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ zTRltR(yzEE5FI, hek9DV.Ifn)),
							Tfn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ zTRltR(yzEE5FI, hek9DV.Tfn)),
						})
					}
				}
				if pxxIZC46 :=  /*line :1*/ cf3fCV8ayFq.Elem.Uncommon(); pxxIZC46 != nil {
					for _, hek9DV := range  /*line :1*/ pxxIZC46.Methods() {
						f4pekhjLv :=  /*line :1*/ drswiOv(yzEE5FI, hek9DV.Name)
						if  /*line :1*/ zIgMu_ZY(f4pekhjLv) != "" {

							 /*line :1*/ panic("reflect: embedded interface with unexported method(s) not implemented")
						}
						r2zNEKaY =  /*line :1*/ append(r2zNEKaY, abi.Method{
							Name:	 /*line :1*/ vybOmhD9hqa5(f4pekhjLv),
							Mtyp:	 /*line :1*/ j2acIXKpV( /*line :1*/ nsOKvPP(cf3fCV8ayFq.Elem, hek9DV.Mtyp)),
							Ifn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ zTRltR(cf3fCV8ayFq.Elem, hek9DV.Ifn)),
							Tfn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ zTRltR(cf3fCV8ayFq.Elem, hek9DV.Tfn)),
						})
					}
				}
			default:
				if pxxIZC46 :=  /*line :1*/ yzEE5FI.Uncommon(); pxxIZC46 != nil {
					if hXZpj1nTZ > 0 && pxxIZC46.Mcount > 0 {

						 /*line :1*/ panic("reflect: embedded type with methods not implemented if type is not first field")
					}
					if  /*line :1*/ len(hdC2Rc) > 1 && yzEE5FI.Kind_&kindDirectIface != 0 {
						 /*line :1*/ panic("reflect: embedded type with methods not implemented for non-pointer type")
					}
					for _, hek9DV := range  /*line :1*/ pxxIZC46.Methods() {
						f4pekhjLv :=  /*line :1*/ drswiOv(yzEE5FI, hek9DV.Name)
						if  /*line :1*/ zIgMu_ZY(f4pekhjLv) != "" {

							 /*line :1*/ panic("reflect: embedded interface with unexported method(s) not implemented")
						}
						r2zNEKaY =  /*line :1*/ append(r2zNEKaY, abi.Method{
							Name:	 /*line :1*/ vybOmhD9hqa5(f4pekhjLv),
							Mtyp:	 /*line :1*/ j2acIXKpV( /*line :1*/ nsOKvPP(yzEE5FI, hek9DV.Mtyp)),
							Ifn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ zTRltR(yzEE5FI, hek9DV.Ifn)),
							Tfn:	 /*line :1*/ gGBM5WeQM( /*line :1*/ zTRltR(yzEE5FI, hek9DV.Tfn)),
						})

					}
				}
			}
		}
		if _, zOlRraIB2IS := lsockidttb[lGwAXqN4Hb]; zOlRraIB2IS && lGwAXqN4Hb != "_" {
			 /*line :1*/ panic("reflect.StructOf: duplicate field " + lGwAXqN4Hb)
		}
		lsockidttb[lGwAXqN4Hb] = struct{}{}

		x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL,  /*line :1*/ byte(yzEE5FI.Hash>>24),  /*line :1*/ byte(yzEE5FI.Hash>>16),  /*line :1*/ byte(yzEE5FI.Hash>>8),  /*line :1*/ byte(yzEE5FI.Hash))

		sfIHyhaL =  /*line :1*/ append(sfIHyhaL, (" " +  /*line :1*/ iR_1JLd(yzEE5FI))...)
		if  /*line :1*/ jToThzM.Name.HasTag() {
			x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL, [] /*line :1*/ byte( /*line :1*/ jToThzM.Name.Tag())...)
			sfIHyhaL =  /*line :1*/ append(sfIHyhaL, (" " +  /*line :1*/ strconv.D35LD5nn( /*line :1*/ jToThzM.Name.Tag()))...)
		}
		if hXZpj1nTZ <  /*line :1*/ len(hdC2Rc)-1 {
			sfIHyhaL =  /*line :1*/ append(sfIHyhaL, ';')
		}

		vaI0gqokB = vaI0gqokB && (yzEE5FI.Equal != nil)

		dcDpXllPZ3e :=  /*line :1*/ rlY5YEji1B4(wnTnByH,  /*line :1*/ uintptr(yzEE5FI.Align_))
		if dcDpXllPZ3e < wnTnByH {
			 /*line :1*/ panic("reflect.StructOf: struct size would exceed virtual address space")
		}
		if yzEE5FI.Align_ > aO_2b3G {
			aO_2b3G = yzEE5FI.Align_
		}
		wnTnByH = dcDpXllPZ3e + yzEE5FI.Size_
		if wnTnByH < dcDpXllPZ3e {
			 /*line :1*/ panic("reflect.StructOf: struct size would exceed virtual address space")
		}
		jToThzM.Offset = dcDpXllPZ3e

		if yzEE5FI.Size_ == 0 {
			kraXVtm = wnTnByH
		}

		rH64rFnxgI47[hXZpj1nTZ] = jToThzM
	}

	if wnTnByH > 0 && kraXVtm == wnTnByH {

		wnTnByH++
		if wnTnByH == 0 {
			 /*line :1*/ panic("reflect.StructOf: struct size would exceed virtual address space")
		}
	}

	var lm5DH3 *structType
	var sxpcYEiFEM *aKCY_EtrgnA8

	if  /*line :1*/ len(r2zNEKaY) == 0 {
		sw8_lJ :=  /*line :1*/ new(structTypeUncommon)
		lm5DH3 = &sw8_lJ.structType
		sxpcYEiFEM = &sw8_lJ.u
	} else {

		mP7QqSVInV :=  /*line :1*/ L6EAuajfYe1( /*line :1*/ FNkNFKGvwkzw([]EeSeBImJjDQo{
			{ZOSBHkJz: "S", S1miDNBAV8a:  /*line :1*/ Cher1a2Fblr(structType{})},
			{ZOSBHkJz: "U", S1miDNBAV8a:  /*line :1*/ Cher1a2Fblr(aKCY_EtrgnA8{})},
			{ZOSBHkJz: "M", S1miDNBAV8a:  /*line :1*/ PNcEGf7fzu( /*line :1*/ len(r2zNEKaY),  /*line :1*/ Cher1a2Fblr(r2zNEKaY[0]))},
		}))

		lm5DH3 = (* /*line :1*/ structType)( /*line :1*/ mP7QqSVInV.Elem().Field(0).Addr().UnsafePointer())
		sxpcYEiFEM = (* /*line :1*/ aKCY_EtrgnA8)( /*line :1*/ mP7QqSVInV.Elem().Field(1).Addr().UnsafePointer())

		 /*line :1*/ copy( /*line :1*/ mP7QqSVInV.Elem().Field(2).Slice(0,  /*line :1*/ len(r2zNEKaY)).Interface().([]abi.Method), r2zNEKaY)
	}

	sxpcYEiFEM.Mcount =  /*line :1*/ uint16( /*line :1*/ len(r2zNEKaY))
	sxpcYEiFEM.Xcount = sxpcYEiFEM.Mcount
	sxpcYEiFEM.Moff =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(aKCY_EtrgnA8{}))

	if  /*line :1*/ len(rH64rFnxgI47) > 0 {
		sfIHyhaL =  /*line :1*/ append(sfIHyhaL, ' ')
	}
	sfIHyhaL =  /*line :1*/ append(sfIHyhaL, '}')
	x6CWj5lQL =  /*line :1*/ jSX7rdi(x6CWj5lQL, '}')
	oQ6fAx :=  /*line :1*/ string(sfIHyhaL)

	kXTvhUI4Zgg5 :=  /*line :1*/ rlY5YEji1B4(wnTnByH,  /*line :1*/ uintptr(aO_2b3G))
	if kXTvhUI4Zgg5 < wnTnByH {
		 /*line :1*/ panic("reflect.StructOf: struct size would exceed virtual address space")
	}
	wnTnByH = kXTvhUI4Zgg5

	var q5bexavdDP7n any = struct{}{}
	fNjFa2jLl := *(** /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(&q5bexavdDP7n))
	*lm5DH3 = *fNjFa2jLl
	lm5DH3.Fields = rH64rFnxgI47
	if enqJZy != "" {
		lm5DH3.PkgPath =  /*line :1*/ kxqWl_C(enqJZy, "", false, false)
	}

	if ibdpoF, bDfzXlCm5Raf :=  /*line :1*/ f6wPeOefr.mzTN8ya.Load(x6CWj5lQL); bDfzXlCm5Raf {
		for _, pjngE6n4a := range ibdpoF.([]YJh989LTZsVX) {
			sw8_lJ :=  /*line :1*/ pjngE6n4a.z6hGxTABM1()
			if  /*line :1*/ sr7apHgO(&lm5DH3.Type, sw8_lJ, true) {
				return  /*line :1*/ nadfj4Ka8cQf(sw8_lJ)
			}
		}
	}

	 /*line :1*/ f6wPeOefr.Lock()
	defer  /*line :1*/ f6wPeOefr.Unlock()
	if ibdpoF, bDfzXlCm5Raf :=  /*line :1*/ f6wPeOefr.mzTN8ya.Load(x6CWj5lQL); bDfzXlCm5Raf {
		for _, pjngE6n4a := range ibdpoF.([]YJh989LTZsVX) {
			sw8_lJ :=  /*line :1*/ pjngE6n4a.z6hGxTABM1()
			if  /*line :1*/ sr7apHgO(&lm5DH3.Type, sw8_lJ, true) {
				return  /*line :1*/ nadfj4Ka8cQf(sw8_lJ)
			}
		}
	}

	skXTsm0V := func(sw8_lJ YJh989LTZsVX) YJh989LTZsVX {
		var ibdpoF []YJh989LTZsVX
		if hhIBKzzDq, bDfzXlCm5Raf :=  /*line :1*/ f6wPeOefr.mzTN8ya.Load(x6CWj5lQL); bDfzXlCm5Raf {
			ibdpoF = hhIBKzzDq.([]YJh989LTZsVX)
		}
		 /*line :1*/ f6wPeOefr.mzTN8ya.Store(x6CWj5lQL,  /*line :1*/ append(ibdpoF, sw8_lJ))
		return sw8_lJ
	}

	for _, sw8_lJ := range  /*line :1*/ b0uDpEzf(oQ6fAx) {
		if  /*line :1*/ sr7apHgO(&lm5DH3.Type, sw8_lJ, true) {

			return  /*line :1*/ skXTsm0V( /*line :1*/ nadfj4Ka8cQf(sw8_lJ))
		}
	}

	lm5DH3.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(oQ6fAx, "", false, false))
	lm5DH3.TFlag = 0
	lm5DH3.Hash = x6CWj5lQL
	lm5DH3.Size_ = wnTnByH
	lm5DH3.PtrBytes =  /*line :1*/ tXiL831F(&lm5DH3.Type)
	lm5DH3.Align_ = aO_2b3G
	lm5DH3.FieldAlign_ = aO_2b3G
	lm5DH3.PtrToThis = 0
	if  /*line :1*/ len(r2zNEKaY) > 0 {
		lm5DH3.TFlag |= abi.TFlagUncommon
	}

	if e_ELAml11yj {
		vjIZTh8 := 0
		for hXZpj1nTZ, yzEE5FI := range rH64rFnxgI47 {
			if  /*line :1*/ yzEE5FI.Typ.Pointers() {
				vjIZTh8 = hXZpj1nTZ
			}
		}
		lanXsCdMk := []byte{0, 0, 0, 0}
		var e65m4G53plB uintptr
		for hXZpj1nTZ, yzEE5FI := range rH64rFnxgI47 {
			if hXZpj1nTZ > vjIZTh8 {

				break
			}
			if ! /*line :1*/ yzEE5FI.Typ.Pointers() {

				continue
			}

			if yzEE5FI.Offset > e65m4G53plB {
				krtR1HwEan := (yzEE5FI.Offset - e65m4G53plB) / goarch.PtrSize
				lanXsCdMk =  /*line :1*/ append(lanXsCdMk, 0x01, 0x00)
				if krtR1HwEan > 1 {
					lanXsCdMk =  /*line :1*/ append(lanXsCdMk, 0x81)
					lanXsCdMk =  /*line :1*/ nod27f406(lanXsCdMk, krtR1HwEan-1)
				}
				e65m4G53plB = yzEE5FI.Offset
			}

			lanXsCdMk =  /*line :1*/ c7bREF5(lanXsCdMk, yzEE5FI.Typ)
			e65m4G53plB += yzEE5FI.Typ.PtrBytes
		}
		lanXsCdMk =  /*line :1*/ append(lanXsCdMk, 0)
		*(* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(&lanXsCdMk[0])) =  /*line :1*/ uint32( /*line :1*/ len(lanXsCdMk) - 4)
		lm5DH3.Kind_ |= kindGCProg
		lm5DH3.GCData = &lanXsCdMk[0]
	} else {
		lm5DH3.Kind_ &^= kindGCProg
		ukOevmakDI3b :=  /*line :1*/ new(uh88jBcDEay2)
		 /*line :1*/ hK9wu75JAa(ukOevmakDI3b, 0, &lm5DH3.Type)
		if  /*line :1*/ len(ukOevmakDI3b.oLIdayf) > 0 {
			lm5DH3.GCData = &ukOevmakDI3b.oLIdayf[0]
		}
	}
	lm5DH3.Equal = nil
	if vaI0gqokB {
		lm5DH3.Equal = func(e3IyyaQSEj, fTdFSkaVjh unsafe.Pointer) bool {
			for _, yzEE5FI := range lm5DH3.Fields {
				p4rJlS0a8 :=  /*line :1*/ lxkP7oV7sm(e3IyyaQSEj, yzEE5FI.Offset, "&x.field safe")
				afHaSxSXrg :=  /*line :1*/ lxkP7oV7sm(fTdFSkaVjh, yzEE5FI.Offset, "&x.field safe")
				if ! /*line :1*/ yzEE5FI.Typ.Equal(p4rJlS0a8, afHaSxSXrg) {
					return false
				}
			}
			return true
		}
	}

	switch {
	case  /*line :1*/ len(rH64rFnxgI47) == 1 && ! /*line :1*/ gGK1Oc(rH64rFnxgI47[0].Typ):

		lm5DH3.Kind_ |= kindDirectIface
	default:
		lm5DH3.Kind_ &^= kindDirectIface
	}

	return  /*line :1*/ skXTsm0V( /*line :1*/ nadfj4Ka8cQf(&lm5DH3.Type))
}

func pxDaXfB5tPl8(fBabaKImKM2F EeSeBImJjDQo) (yUesa1U9, string) {
	if fBabaKImKM2F.GuSnlgXxlhPz && fBabaKImKM2F.H1I0k6V != "" {
		 /*line :1*/ panic("reflect.StructOf: field \"" + fBabaKImKM2F.ZOSBHkJz + "\" is anonymous but has PkgPath set")
	}

	if  /*line :1*/ fBabaKImKM2F.IsExported() {

		i5hAnuqALA := fBabaKImKM2F.ZOSBHkJz[0]
		if 'a' <= i5hAnuqALA && i5hAnuqALA <= 'z' || i5hAnuqALA == '_' {
			 /*line :1*/ panic("reflect.StructOf: field \"" + fBabaKImKM2F.ZOSBHkJz + "\" is unexported but missing PkgPath")
		}
	}

	 /*line :1*/ j2acIXKpV( /*line :1*/ fBabaKImKM2F.S1miDNBAV8a.z6hGxTABM1())
	jToThzM := yUesa1U9{
		Name:	 /*line :1*/ kxqWl_C(fBabaKImKM2F.ZOSBHkJz,  /*line :1*/ string(fBabaKImKM2F.Kp_2l7Su8ao),  /*line :1*/ fBabaKImKM2F.IsExported(), fBabaKImKM2F.GuSnlgXxlhPz),
		Typ:	 /*line :1*/ fBabaKImKM2F.S1miDNBAV8a.z6hGxTABM1(),
		Offset:	0,
	}
	return jToThzM, fBabaKImKM2F.H1I0k6V
}

func tXiL831F(sw8_lJ *abi.Type) uintptr {
	switch  /*line :1*/ sw8_lJ.Kind() {
	case abi.Struct:
		pjngE6n4a := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))

		fBabaKImKM2F := -1
		for hXZpj1nTZ := range pjngE6n4a.Fields {
			yzEE5FI := pjngE6n4a.Fields[hXZpj1nTZ].Typ
			if  /*line :1*/ yzEE5FI.Pointers() {
				fBabaKImKM2F = hXZpj1nTZ
			}
		}
		if fBabaKImKM2F == -1 {
			return 0
		}
		jToThzM := pjngE6n4a.Fields[fBabaKImKM2F]
		return jToThzM.Offset + jToThzM.Typ.PtrBytes

	default:
		 /*line :1*/ panic("reflect.typeptrdata: unexpected type, " +  /*line :1*/ iR_1JLd(sw8_lJ))
	}
}

const maxPtrmaskBytes = 2048

func PNcEGf7fzu(tN6IaOft int, dPe0EQn YJh989LTZsVX) YJh989LTZsVX {
	if tN6IaOft < 0 {
		 /*line :1*/ panic("reflect: negative length passed to ArrayOf")
	}

	lm5DH3 :=  /*line :1*/ dPe0EQn.z6hGxTABM1()

	ctW2F7ER := kAM7Km9_stw{Array, lm5DH3, nil,  /*line :1*/ uintptr(tN6IaOft)}
	if rHaMAiU, bDfzXlCm5Raf :=  /*line :1*/ hDZJaY1mZDEB.Load(ctW2F7ER); bDfzXlCm5Raf {
		return rHaMAiU.(YJh989LTZsVX)
	}

	kXTvhUI4Zgg5 := "[" +  /*line :1*/ strconv.ZW1_FxJtq(tN6IaOft) + "]" +  /*line :1*/ iR_1JLd(lm5DH3)
	for _, mP7QqSVInV := range  /*line :1*/ b0uDpEzf(kXTvhUI4Zgg5) {
		rHaMAiU := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(mP7QqSVInV))
		if rHaMAiU.Elem == lm5DH3 {
			hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(mP7QqSVInV))
			return hhIBKzzDq.(YJh989LTZsVX)
		}
	}

	var oOHfYUZBOSV any = [1]unsafe.Pointer{}
	fNjFa2jLl := *(** /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(&oOHfYUZBOSV))
	rHaMAiU := *fNjFa2jLl
	rHaMAiU.TFlag = lm5DH3.TFlag & abi.TFlagRegularMemory
	rHaMAiU.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(kXTvhUI4Zgg5, "", false, false))
	rHaMAiU.Hash =  /*line :1*/ jSX7rdi(lm5DH3.Hash, '[')
	for krtR1HwEan :=  /*line :1*/ uint32(tN6IaOft); krtR1HwEan > 0; krtR1HwEan >>= 8 {
		rHaMAiU.Hash =  /*line :1*/ jSX7rdi(rHaMAiU.Hash,  /*line :1*/ byte(krtR1HwEan))
	}
	rHaMAiU.Hash =  /*line :1*/ jSX7rdi(rHaMAiU.Hash, ']')
	rHaMAiU.Elem = lm5DH3
	rHaMAiU.PtrToThis = 0
	if lm5DH3.Size_ > 0 {
		odgetBV := ^ /*line :1*/ uintptr(0) / lm5DH3.Size_
		if  /*line :1*/ uintptr(tN6IaOft) > odgetBV {
			 /*line :1*/ panic("reflect.ArrayOf: array size would exceed virtual address space")
		}
	}
	rHaMAiU.Size_ = lm5DH3.Size_ *  /*line :1*/ uintptr(tN6IaOft)
	if tN6IaOft > 0 && lm5DH3.PtrBytes != 0 {
		rHaMAiU.PtrBytes = lm5DH3.Size_* /*line :1*/ uintptr(tN6IaOft-1) + lm5DH3.PtrBytes
	}
	rHaMAiU.Align_ = lm5DH3.Align_
	rHaMAiU.FieldAlign_ = lm5DH3.FieldAlign_
	rHaMAiU.Len =  /*line :1*/ uintptr(tN6IaOft)
	rHaMAiU.Slice = &( /*line :1*/ XsHS1gYvDYpj(dPe0EQn).(*rtype).t)

	switch {
	case lm5DH3.PtrBytes == 0 || rHaMAiU.Size_ == 0:

		rHaMAiU.GCData = nil
		rHaMAiU.PtrBytes = 0

	case tN6IaOft == 1:

		rHaMAiU.Kind_ |= lm5DH3.Kind_ & kindGCProg
		rHaMAiU.GCData = lm5DH3.GCData
		rHaMAiU.PtrBytes = lm5DH3.PtrBytes

	case lm5DH3.Kind_&kindGCProg == 0 && rHaMAiU.Size_ <= maxPtrmaskBytes*8*goarch.PtrSize:

		krtR1HwEan := (rHaMAiU.PtrBytes/goarch.PtrSize + 7) / 8

		krtR1HwEan = (krtR1HwEan + goarch.PtrSize - 1) &^ (goarch.PtrSize - 1)
		hpoWZKhOa0 :=  /*line :1*/ make([]byte, krtR1HwEan)
		 /*line :1*/ r1PEsohea(hpoWZKhOa0, 0, lm5DH3, rHaMAiU.Len)
		rHaMAiU.GCData = &hpoWZKhOa0[0]

	default:

		lanXsCdMk := []byte{0, 0, 0, 0}
		lanXsCdMk =  /*line :1*/ c7bREF5(lanXsCdMk, lm5DH3)

		uua0FKa := lm5DH3.PtrBytes / goarch.PtrSize
		qg2hEN4fE1 := lm5DH3.Size_ / goarch.PtrSize
		if uua0FKa < qg2hEN4fE1 {

			lanXsCdMk =  /*line :1*/ append(lanXsCdMk, 0x01, 0x00)
			if uua0FKa+1 < qg2hEN4fE1 {
				lanXsCdMk =  /*line :1*/ append(lanXsCdMk, 0x81)
				lanXsCdMk =  /*line :1*/ nod27f406(lanXsCdMk, qg2hEN4fE1-uua0FKa-1)
			}
		}

		if qg2hEN4fE1 < 0x80 {
			lanXsCdMk =  /*line :1*/ append(lanXsCdMk,  /*line :1*/ byte(qg2hEN4fE1|0x80))
		} else {
			lanXsCdMk =  /*line :1*/ append(lanXsCdMk, 0x80)
			lanXsCdMk =  /*line :1*/ nod27f406(lanXsCdMk, qg2hEN4fE1)
		}
		lanXsCdMk =  /*line :1*/ nod27f406(lanXsCdMk,  /*line :1*/ uintptr(tN6IaOft)-1)
		lanXsCdMk =  /*line :1*/ append(lanXsCdMk, 0)
		*(* /*line :1*/ uint32)( /*line :1*/ unsafe.Pointer(&lanXsCdMk[0])) =  /*line :1*/ uint32( /*line :1*/ len(lanXsCdMk) - 4)
		rHaMAiU.Kind_ |= kindGCProg
		rHaMAiU.GCData = &lanXsCdMk[0]
		rHaMAiU.PtrBytes = rHaMAiU.Size_
	}

	uR1VfuHRj := lm5DH3
	l5c0Sh37pj0 :=  /*line :1*/ uR1VfuHRj.Size()

	rHaMAiU.Equal = nil
	if kAjoBJK := uR1VfuHRj.Equal; kAjoBJK != nil {
		rHaMAiU.Equal = func(e3IyyaQSEj, fTdFSkaVjh unsafe.Pointer) bool {
			for hXZpj1nTZ := 0; hXZpj1nTZ < tN6IaOft; hXZpj1nTZ++ {
				p4rJlS0a8 :=  /*line :1*/ rfJB4Et(e3IyyaQSEj, hXZpj1nTZ, l5c0Sh37pj0, "i < length")
				afHaSxSXrg :=  /*line :1*/ rfJB4Et(fTdFSkaVjh, hXZpj1nTZ, l5c0Sh37pj0, "i < length")
				if ! /*line :1*/ kAjoBJK(p4rJlS0a8, afHaSxSXrg) {
					return false
				}

			}
			return true
		}
	}

	switch {
	case tN6IaOft == 1 && ! /*line :1*/ gGK1Oc(lm5DH3):

		rHaMAiU.Kind_ |= kindDirectIface
	default:
		rHaMAiU.Kind_ &^= kindDirectIface
	}

	hhIBKzzDq, _ :=  /*line :1*/ hDZJaY1mZDEB.LoadOrStore(ctW2F7ER,  /*line :1*/ j5Z_WRWLudek(&rHaMAiU.Type))
	return hhIBKzzDq.(YJh989LTZsVX)
}

func nod27f406(uiTlN8bdKm []byte, f8NmcatRx uintptr) []byte {
	for ; f8NmcatRx >= 0x80; f8NmcatRx >>= 7 {
		uiTlN8bdKm =  /*line :1*/ append(uiTlN8bdKm,  /*line :1*/ byte(f8NmcatRx|0x80))
	}
	uiTlN8bdKm =  /*line :1*/ append(uiTlN8bdKm,  /*line :1*/ byte(f8NmcatRx))
	return uiTlN8bdKm
}

func nadfj4Ka8cQf(sw8_lJ *abi.Type) YJh989LTZsVX {
	if sw8_lJ == nil {
		return nil
	}
	return  /*line :1*/ j5Z_WRWLudek(sw8_lJ)
}

type dvitRt8fhAO struct {
	fmFPAUN1	*bY2Zo8pis
	cB1AK3b2ng	*abi.Type
}

type ruRvzUSabErG struct {
	xMH2vntPHO	*abi.Type
	nNoQkFyDz	*sync.OrP5FGPq
	jSr8GnCjnUA	cxVCV8
}

var wDdzXAaM sync.Eim1b6bNi9IM

func thrShO4Mk(sw8_lJ *bY2Zo8pis, iGkAMNgmOX *abi.Type) (nuCjUv_xpB0p *abi.Type, a2eCYr *sync.OrP5FGPq, dt53aB1lcTF cxVCV8) {
	if  /*line :1*/ sw8_lJ.Kind() != abi.Func {
		 /*line :1*/ panic("reflect: funcLayout of non-func type " +  /*line :1*/ iR_1JLd(&sw8_lJ.Type))
	}
	if iGkAMNgmOX != nil &&  /*line :1*/ iGkAMNgmOX.Kind() == abi.Interface {
		 /*line :1*/ panic("reflect: funcLayout with interface receiver " +  /*line :1*/ iR_1JLd(iGkAMNgmOX))
	}
	lYOqUQga := dvitRt8fhAO{sw8_lJ, iGkAMNgmOX}
	if hFivLBXeP9LG, bDfzXlCm5Raf :=  /*line :1*/ wDdzXAaM.Load(lYOqUQga); bDfzXlCm5Raf {
		nL5LM2qj := hFivLBXeP9LG.(ruRvzUSabErG)
		return nL5LM2qj.xMH2vntPHO, nL5LM2qj.nNoQkFyDz, nL5LM2qj.jSr8GnCjnUA
	}

	dt53aB1lcTF =  /*line :1*/ y6pUOfj6v(sw8_lJ, iGkAMNgmOX)

	uiTlN8bdKm := &abi.Type{
		Align_:	goarch.PtrSize,

		Size_:	 /*line :1*/ rlY5YEji1B4(dt53aB1lcTF.eBcFPCE+dt53aB1lcTF.gSH9DY2KdKea.rHUhKYMAh, goarch.PtrSize),
		PtrBytes:	 /*line :1*/ uintptr(dt53aB1lcTF.sCYaW3ZJnC.m543jWoDkPsm) * goarch.PtrSize,
	}
	if dt53aB1lcTF.sCYaW3ZJnC.m543jWoDkPsm > 0 {
		uiTlN8bdKm.GCData = &dt53aB1lcTF.sCYaW3ZJnC.oLIdayf[0]
	}

	var kXTvhUI4Zgg5 string
	if iGkAMNgmOX != nil {
		kXTvhUI4Zgg5 = "methodargs(" +  /*line :1*/ iR_1JLd(iGkAMNgmOX) + ")(" +  /*line :1*/ iR_1JLd(&sw8_lJ.Type) + ")"
	} else {
		kXTvhUI4Zgg5 = "funcargs(" +  /*line :1*/ iR_1JLd(&sw8_lJ.Type) + ")"
	}
	uiTlN8bdKm.Str =  /*line :1*/ vybOmhD9hqa5( /*line :1*/ kxqWl_C(kXTvhUI4Zgg5, "", false, false))

	a2eCYr = &sync.OrP5FGPq{IYbTy050ek: func() any {
		return  /*line :1*/ qscv_5(uiTlN8bdKm)
	}}
	hFivLBXeP9LG, _ :=  /*line :1*/ wDdzXAaM.LoadOrStore(lYOqUQga, ruRvzUSabErG{
		xMH2vntPHO:	uiTlN8bdKm,
		nNoQkFyDz:	a2eCYr,
		jSr8GnCjnUA:	dt53aB1lcTF,
	})
	nL5LM2qj := hFivLBXeP9LG.(ruRvzUSabErG)
	return nL5LM2qj.xMH2vntPHO, nL5LM2qj.nNoQkFyDz, nL5LM2qj.jSr8GnCjnUA
}

func gGK1Oc(sw8_lJ *abi.Type) bool {
	return sw8_lJ.Kind_&kindDirectIface == 0
}

type uh88jBcDEay2 struct {
	m543jWoDkPsm	uint32
	oLIdayf	[]byte
}

func (ukOevmakDI3b *uh88jBcDEay2) cn3HN7Iy(_aNsGCl uint8) {
	if ukOevmakDI3b.m543jWoDkPsm%(8*goarch.PtrSize) == 0 {

		for hXZpj1nTZ := 0; hXZpj1nTZ < goarch.PtrSize; hXZpj1nTZ++ {
			ukOevmakDI3b.oLIdayf =  /*line :1*/ append(ukOevmakDI3b.oLIdayf, 0)
		}
	}
	ukOevmakDI3b.oLIdayf[ukOevmakDI3b.m543jWoDkPsm/8] |= _aNsGCl << (ukOevmakDI3b.m543jWoDkPsm % 8)
	ukOevmakDI3b.m543jWoDkPsm++
}

func hK9wu75JAa(ukOevmakDI3b *uh88jBcDEay2, dcDpXllPZ3e uintptr, sw8_lJ *abi.Type) {
	if sw8_lJ.PtrBytes == 0 {
		return
	}

	switch  /*line :1*/ WzYBjsQL0(sw8_lJ.Kind_ & kindMask) {
	case Chan, Func, Map, Pointer, Slice, String, UnsafePointer:

		for ukOevmakDI3b.m543jWoDkPsm <  /*line :1*/ uint32(dcDpXllPZ3e/ /*line :1*/ uintptr(goarch.PtrSize)) {
			 /*line :1*/ ukOevmakDI3b.cn3HN7Iy(0)
		}
		 /*line :1*/ ukOevmakDI3b.cn3HN7Iy(1)

	case Interface:

		for ukOevmakDI3b.m543jWoDkPsm <  /*line :1*/ uint32(dcDpXllPZ3e/ /*line :1*/ uintptr(goarch.PtrSize)) {
			 /*line :1*/ ukOevmakDI3b.cn3HN7Iy(0)
		}
		 /*line :1*/ ukOevmakDI3b.cn3HN7Iy(1)
		 /*line :1*/ ukOevmakDI3b.cn3HN7Iy(1)

	case Array:

		mP7QqSVInV := (* /*line :1*/ uV_BfrSO)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ int(mP7QqSVInV.Len); hXZpj1nTZ++ {
			 /*line :1*/ hK9wu75JAa(ukOevmakDI3b, dcDpXllPZ3e+ /*line :1*/ uintptr(hXZpj1nTZ)*mP7QqSVInV.Elem.Size_, mP7QqSVInV.Elem)
		}

	case Struct:

		mP7QqSVInV := (* /*line :1*/ structType)( /*line :1*/ unsafe.Pointer(sw8_lJ))
		for hXZpj1nTZ := range mP7QqSVInV.Fields {
			jToThzM := &mP7QqSVInV.Fields[hXZpj1nTZ]
			 /*line :1*/ hK9wu75JAa(ukOevmakDI3b, dcDpXllPZ3e+jToThzM.Offset, jToThzM.Typ)
		}
	}
}
