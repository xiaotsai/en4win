//line :1
package os208GicmRnc

import (
	godebug "rayT9AgS"
	sync "sync"
	atomic "sync/atomic"
	_ "unsafe"
)

type E9z2NA3b interface {
	Int63() int64
	Seed(wO0Yc_ int64)
}

type S7ymRRYb interface {
	E9z2NA3b
	Uint64() uint64
}

func DF2qvnmwCjq6(wO0Yc_ int64) E9z2NA3b {
	return  /*line :1*/ exq0VCn9(wO0Yc_)
}

func exq0VCn9(wO0Yc_ int64) *eZnGyY {
	var t7PGAzi7IkQ eZnGyY
	 /*line :1*/ t7PGAzi7IkQ.Seed(wO0Yc_)
	return &t7PGAzi7IkQ
}

type YAxSls struct {
	az26tvpsOv	E9z2NA3b
	eJRo04ted	S7ymRRYb

	ax4ZWa4S0	int64

	pleDlth	int8
}

func PamkngYsA(kgWrowfIT2o E9z2NA3b) *YAxSls {
	zJ20UTrW, _ := kgWrowfIT2o.(S7ymRRYb)
	return &YAxSls{az26tvpsOv: kgWrowfIT2o, eJRo04ted: zJ20UTrW}
}

func (csIue666O0e *YAxSls) Seed(wO0Yc_ int64) {
	if gbpH6aY, oDk6uWUPzlI := csIue666O0e.az26tvpsOv.(*tqgxR3sC); oDk6uWUPzlI {
		 /*line :1*/ gbpH6aY.xAtm0rH(wO0Yc_, &csIue666O0e.pleDlth)
		return
	}

	 /*line :1*/ csIue666O0e.az26tvpsOv.Seed(wO0Yc_)
	csIue666O0e.pleDlth = 0
}

func (csIue666O0e *YAxSls) Int63() int64	{ return  /*line :1*/ csIue666O0e.az26tvpsOv.Int63() }

func (csIue666O0e *YAxSls) Uint32() uint32	{ return  /*line :1*/ uint32( /*line :1*/ csIue666O0e.Int63() >> 31) }

func (csIue666O0e *YAxSls) Uint64() uint64 {
	if csIue666O0e.eJRo04ted != nil {
		return  /*line :1*/ csIue666O0e.eJRo04ted.Uint64()
	}
	return  /*line :1*/ uint64( /*line :1*/ csIue666O0e.Int63())>>31 |  /*line :1*/ uint64( /*line :1*/ csIue666O0e.Int63())<<32
}

func (csIue666O0e *YAxSls) Int31() int32	{ return  /*line :1*/ int32( /*line :1*/ csIue666O0e.Int63() >> 32) }

func (csIue666O0e *YAxSls) Int() int {
	tBYgGcxaucp :=  /*line :1*/ uint( /*line :1*/ csIue666O0e.Int63())
	return  /*line :1*/ int(tBYgGcxaucp << 1 >> 1)
}

func (csIue666O0e *YAxSls) Int63n(bLctfJMmsKJ_ int64) int64 {
	if bLctfJMmsKJ_ <= 0 {
		 /*line :1*/ panic("invalid argument to Int63n")
	}
	if bLctfJMmsKJ_&(bLctfJMmsKJ_-1) == 0 {
		return  /*line :1*/ csIue666O0e.Int63() & (bLctfJMmsKJ_ - 1)
	}
	kRdQdN :=  /*line :1*/ int64((1 << 63) - 1 - (1<<63)% /*line :1*/ uint64(bLctfJMmsKJ_))
	zqelgkLITo9 :=  /*line :1*/ csIue666O0e.Int63()
	for zqelgkLITo9 > kRdQdN {
		zqelgkLITo9 =  /*line :1*/ csIue666O0e.Int63()
	}
	return zqelgkLITo9 % bLctfJMmsKJ_
}

func (csIue666O0e *YAxSls) Int31n(bLctfJMmsKJ_ int32) int32 {
	if bLctfJMmsKJ_ <= 0 {
		 /*line :1*/ panic("invalid argument to Int31n")
	}
	if bLctfJMmsKJ_&(bLctfJMmsKJ_-1) == 0 {
		return  /*line :1*/ csIue666O0e.Int31() & (bLctfJMmsKJ_ - 1)
	}
	kRdQdN :=  /*line :1*/ int32((1 << 31) - 1 - (1<<31)% /*line :1*/ uint32(bLctfJMmsKJ_))
	zqelgkLITo9 :=  /*line :1*/ csIue666O0e.Int31()
	for zqelgkLITo9 > kRdQdN {
		zqelgkLITo9 =  /*line :1*/ csIue666O0e.Int31()
	}
	return zqelgkLITo9 % bLctfJMmsKJ_
}

func (csIue666O0e *YAxSls) c6uZ8EtTw(bLctfJMmsKJ_ int32) int32 {
	zqelgkLITo9 :=  /*line :1*/ csIue666O0e.Uint32()
	iihUiXCD :=  /*line :1*/ uint64(zqelgkLITo9) *  /*line :1*/ uint64(bLctfJMmsKJ_)
	zfzUnT8_N :=  /*line :1*/ uint32(iihUiXCD)
	if zfzUnT8_N <  /*line :1*/ uint32(bLctfJMmsKJ_) {
		cr2fi9Mn :=  /*line :1*/ uint32(-bLctfJMmsKJ_) %  /*line :1*/ uint32(bLctfJMmsKJ_)
		for zfzUnT8_N < cr2fi9Mn {
			zqelgkLITo9 =  /*line :1*/ csIue666O0e.Uint32()
			iihUiXCD =  /*line :1*/ uint64(zqelgkLITo9) *  /*line :1*/ uint64(bLctfJMmsKJ_)
			zfzUnT8_N =  /*line :1*/ uint32(iihUiXCD)
		}
	}
	return  /*line :1*/ int32(iihUiXCD >> 32)
}

func (csIue666O0e *YAxSls) Intn(bLctfJMmsKJ_ int) int {
	if bLctfJMmsKJ_ <= 0 {
		 /*line :1*/ panic("invalid argument to Intn")
	}
	if bLctfJMmsKJ_ <= 1<<31-1 {
		return  /*line :1*/ int( /*line :1*/ csIue666O0e.Int31n( /*line :1*/ int32(bLctfJMmsKJ_)))
	}
	return  /*line :1*/ int( /*line :1*/ csIue666O0e.Int63n( /*line :1*/ int64(bLctfJMmsKJ_)))
}

func (csIue666O0e *YAxSls) Float64() float64 {

again:
	iBjfE7IcaD :=  /*line :1*/ float64( /*line :1*/ csIue666O0e.Int63()) / (1 << 63)
	if iBjfE7IcaD == 1 {
		goto again
	}
	return iBjfE7IcaD
}

func (csIue666O0e *YAxSls) Float32() float32 {

again:
	iBjfE7IcaD :=  /*line :1*/ float32( /*line :1*/ csIue666O0e.Float64())
	if iBjfE7IcaD == 1 {
		goto again
	}
	return iBjfE7IcaD
}

func (csIue666O0e *YAxSls) Perm(bLctfJMmsKJ_ int) []int {
	aDioBUyvAkob :=  /*line :1*/ make([]int, bLctfJMmsKJ_)

	for tXqZrxAxK := 0; tXqZrxAxK < bLctfJMmsKJ_; tXqZrxAxK++ {
		nMLtegz7 :=  /*line :1*/ csIue666O0e.Intn(tXqZrxAxK + 1)
		aDioBUyvAkob[tXqZrxAxK] = aDioBUyvAkob[nMLtegz7]
		aDioBUyvAkob[nMLtegz7] = tXqZrxAxK
	}
	return aDioBUyvAkob
}

func (csIue666O0e *YAxSls) Shuffle(bLctfJMmsKJ_ int, kHbp1G func(tXqZrxAxK, nMLtegz7 int)) {
	if bLctfJMmsKJ_ < 0 {
		 /*line :1*/ panic("invalid argument to Shuffle")
	}

	tXqZrxAxK := bLctfJMmsKJ_ - 1
	for ; tXqZrxAxK > 1<<31-1-1; tXqZrxAxK-- {
		nMLtegz7 :=  /*line :1*/ int( /*line :1*/ csIue666O0e.Int63n( /*line :1*/ int64(tXqZrxAxK + 1)))
		 /*line :1*/ kHbp1G(tXqZrxAxK, nMLtegz7)
	}
	for ; tXqZrxAxK > 0; tXqZrxAxK-- {
		nMLtegz7 :=  /*line :1*/ int( /*line :1*/ csIue666O0e.c6uZ8EtTw( /*line :1*/ int32(tXqZrxAxK + 1)))
		 /*line :1*/ kHbp1G(tXqZrxAxK, nMLtegz7)
	}
}

func (csIue666O0e *YAxSls) Read(xdTzGfDU1PdG []byte) (bLctfJMmsKJ_ int, cRYfyyQjL error) {
	switch kgWrowfIT2o := csIue666O0e.az26tvpsOv.(type) {
	case *tqgxR3sC:
		return  /*line :1*/ kgWrowfIT2o.qjrWusLvj(xdTzGfDU1PdG, &csIue666O0e.ax4ZWa4S0, &csIue666O0e.pleDlth)
	case *vQ24Xx1n8a:
		return  /*line :1*/ kgWrowfIT2o.qjrWusLvj(xdTzGfDU1PdG, &csIue666O0e.ax4ZWa4S0, &csIue666O0e.pleDlth)
	}
	return  /*line :1*/ qjrWusLvj(xdTzGfDU1PdG, csIue666O0e.az26tvpsOv, &csIue666O0e.ax4ZWa4S0, &csIue666O0e.pleDlth)
}

func qjrWusLvj(xdTzGfDU1PdG []byte, kgWrowfIT2o E9z2NA3b, lNCZHVi *int64, dMFiZa3C *int8) (bLctfJMmsKJ_ int, cRYfyyQjL error) {
	xnTX_fkfq_k := *dMFiZa3C
	jYAaVx5 := *lNCZHVi
	t7PGAzi7IkQ, _ := kgWrowfIT2o.(*eZnGyY)
	for bLctfJMmsKJ_ = 0; bLctfJMmsKJ_ <  /*line :1*/ len(xdTzGfDU1PdG); bLctfJMmsKJ_++ {
		if xnTX_fkfq_k == 0 {
			if t7PGAzi7IkQ != nil {
				jYAaVx5 =  /*line :1*/ t7PGAzi7IkQ.Int63()
			} else {
				jYAaVx5 =  /*line :1*/ kgWrowfIT2o.Int63()
			}
			xnTX_fkfq_k = 7
		}
		xdTzGfDU1PdG[bLctfJMmsKJ_] =  /*line :1*/ byte(jYAaVx5)
		jYAaVx5 >>= 8
		xnTX_fkfq_k--
	}
	*dMFiZa3C = xnTX_fkfq_k
	*lNCZHVi = jYAaVx5
	return
}

var fgbJoaa1LG atomic.ToSaNw[YAxSls]

var jT6GSHhtr =  /*line :1*/ godebug.HppHHAII("randautoseed")

func m8RtJOY() *YAxSls {
	if csIue666O0e :=  /*line :1*/ fgbJoaa1LG.Load(); csIue666O0e != nil {
		return csIue666O0e
	}

	var csIue666O0e *YAxSls
	if  /*line :1*/ jT6GSHhtr.Value() == "0" {
		 /*line :1*/ jT6GSHhtr.IncNonDefault()
		csIue666O0e =  /*line :1*/ PamkngYsA( /*line :1*/ new(tqgxR3sC))
		 /*line :1*/ csIue666O0e.Seed(1)
	} else {
		csIue666O0e = &YAxSls{
			az26tvpsOv:	&vQ24Xx1n8a{},
			eJRo04ted:	&vQ24Xx1n8a{},
		}
	}

	if ! /*line :1*/ fgbJoaa1LG.CompareAndSwap(nil, csIue666O0e) {

		return  /*line :1*/ fgbJoaa1LG.Load()
	}

	return csIue666O0e
}

//go:linkname ahgDR305a
func ahgDR305a() uint64

type vQ24Xx1n8a struct {
	sSHRg5z sync.DIRWe11KYlYa
}

func (*vQ24Xx1n8a) Int63() int64 {
	return  /*line :1*/ int64( /*line :1*/ ahgDR305a() & rngMask)
}

func (*vQ24Xx1n8a) Seed(int64) {
	 /*line :1*/ panic("internal error: call to fastSource.Seed")
}

func (*vQ24Xx1n8a) Uint64() uint64 {
	return  /*line :1*/ ahgDR305a()
}

func (hrvKMT0mi *vQ24Xx1n8a) qjrWusLvj(xdTzGfDU1PdG []byte, lNCZHVi *int64, dMFiZa3C *int8) (bLctfJMmsKJ_ int, cRYfyyQjL error) {
	 /*line :1*/ hrvKMT0mi.sSHRg5z.Lock()
	bLctfJMmsKJ_, cRYfyyQjL =  /*line :1*/ qjrWusLvj(xdTzGfDU1PdG, hrvKMT0mi, lNCZHVi, dMFiZa3C)
	 /*line :1*/ hrvKMT0mi.sSHRg5z.Unlock()
	return
}

func F7rRxLx0K2UR(wO0Yc_ int64) {
	uu0PC7 :=  /*line :1*/ fgbJoaa1LG.Load()

	if uu0PC7 != nil {
		if _, oDk6uWUPzlI := uu0PC7.az26tvpsOv.(*tqgxR3sC); oDk6uWUPzlI {
			 /*line :1*/ uu0PC7.Seed(wO0Yc_)
			return
		}
	}

	csIue666O0e :=  /*line :1*/ PamkngYsA( /*line :1*/ new(tqgxR3sC))
	 /*line :1*/ csIue666O0e.Seed(wO0Yc_)

	if ! /*line :1*/ fgbJoaa1LG.CompareAndSwap(uu0PC7, csIue666O0e) {

		 /*line :1*/ F7rRxLx0K2UR(wO0Yc_)
	}
}

func YAKe2q() int64	{ return  /*line :1*/ m8RtJOY().Int63() }

func PTo_Gd_96() uint32	{ return  /*line :1*/ m8RtJOY().Uint32() }

func NZA8GS9NK() uint64	{ return  /*line :1*/ m8RtJOY().Uint64() }

func MaDCl4Q4MO() int32	{ return  /*line :1*/ m8RtJOY().Int31() }

func Ugs0xw() int	{ return  /*line :1*/ m8RtJOY().Int() }

func Io8xUt(bLctfJMmsKJ_ int64) int64	{ return  /*line :1*/ m8RtJOY().Int63n(bLctfJMmsKJ_) }

func Hs7pnbO(bLctfJMmsKJ_ int32) int32	{ return  /*line :1*/ m8RtJOY().Int31n(bLctfJMmsKJ_) }

func H1QzGJ(bLctfJMmsKJ_ int) int	{ return  /*line :1*/ m8RtJOY().Intn(bLctfJMmsKJ_) }

func UO9dzu() float64	{ return  /*line :1*/ m8RtJOY().Float64() }

func BsoTiUwYJMny() float32	{ return  /*line :1*/ m8RtJOY().Float32() }

func ViP8C9wi(bLctfJMmsKJ_ int) []int	{ return  /*line :1*/ m8RtJOY().Perm(bLctfJMmsKJ_) }

func VuGjf9(bLctfJMmsKJ_ int, kHbp1G func(tXqZrxAxK, nMLtegz7 int)) {
	 /*line :1*/ m8RtJOY().Shuffle(bLctfJMmsKJ_, kHbp1G)
}

func EB3BKNe(xdTzGfDU1PdG []byte) (bLctfJMmsKJ_ int, cRYfyyQjL error) {
	return  /*line :1*/ m8RtJOY().Read(xdTzGfDU1PdG)
}

func APaIRAM6() float64	{ return  /*line :1*/ m8RtJOY().NormFloat64() }

func AgGAD5URp() float64	{ return  /*line :1*/ m8RtJOY().ExpFloat64() }

type tqgxR3sC struct {
	dI8bPPl	sync.DIRWe11KYlYa
	oS2v8FYTd	*eZnGyY
}

func (csIue666O0e *tqgxR3sC) Int63() (bLctfJMmsKJ_ int64) {
	 /*line :1*/ csIue666O0e.dI8bPPl.Lock()
	bLctfJMmsKJ_ =  /*line :1*/ csIue666O0e.oS2v8FYTd.Int63()
	 /*line :1*/ csIue666O0e.dI8bPPl.Unlock()
	return
}

func (csIue666O0e *tqgxR3sC) Uint64() (bLctfJMmsKJ_ uint64) {
	 /*line :1*/ csIue666O0e.dI8bPPl.Lock()
	bLctfJMmsKJ_ =  /*line :1*/ csIue666O0e.oS2v8FYTd.Uint64()
	 /*line :1*/ csIue666O0e.dI8bPPl.Unlock()
	return
}

func (csIue666O0e *tqgxR3sC) Seed(wO0Yc_ int64) {
	 /*line :1*/ csIue666O0e.dI8bPPl.Lock()
	 /*line :1*/ csIue666O0e.wO0Yc_(wO0Yc_)
	 /*line :1*/ csIue666O0e.dI8bPPl.Unlock()
}

func (csIue666O0e *tqgxR3sC) xAtm0rH(wO0Yc_ int64, dMFiZa3C *int8) {
	 /*line :1*/ csIue666O0e.dI8bPPl.Lock()
	 /*line :1*/ csIue666O0e.wO0Yc_(wO0Yc_)
	*dMFiZa3C = 0
	 /*line :1*/ csIue666O0e.dI8bPPl.Unlock()
}

func (csIue666O0e *tqgxR3sC) wO0Yc_(wO0Yc_ int64) {
	if csIue666O0e.oS2v8FYTd == nil {
		csIue666O0e.oS2v8FYTd =  /*line :1*/ exq0VCn9(wO0Yc_)
	} else {
		 /*line :1*/ csIue666O0e.oS2v8FYTd.Seed(wO0Yc_)
	}
}

func (csIue666O0e *tqgxR3sC) qjrWusLvj(xdTzGfDU1PdG []byte, lNCZHVi *int64, dMFiZa3C *int8) (bLctfJMmsKJ_ int, cRYfyyQjL error) {
	 /*line :1*/ csIue666O0e.dI8bPPl.Lock()
	bLctfJMmsKJ_, cRYfyyQjL =  /*line :1*/ qjrWusLvj(xdTzGfDU1PdG, csIue666O0e.oS2v8FYTd, lNCZHVi, dMFiZa3C)
	 /*line :1*/ csIue666O0e.dI8bPPl.Unlock()
	return
}
