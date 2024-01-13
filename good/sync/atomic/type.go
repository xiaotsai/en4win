//line :1
package atomic

import "unsafe"



type Akm7q89_ struct {
	_	axM70DO7
	fq6pSyuu	uint32
}


func (woaZ1rb *Akm7q89_) Load() bool	{ return  /*line :1*/ LoadUint32(&woaZ1rb.fq6pSyuu) != 0 }


func (woaZ1rb *Akm7q89_) Store(fSC10Rr9Vh bool)	{  /*line :1*/ StoreUint32(&woaZ1rb.fq6pSyuu,  /*line :1*/ l7Iap5(fSC10Rr9Vh)) }


func (woaZ1rb *Akm7q89_) Swap(vxcCAFH4Bs bool) (qVwQ2ijo0 bool) {
	return  /*line :1*/ SwapUint32(&woaZ1rb.fq6pSyuu,  /*line :1*/ l7Iap5(vxcCAFH4Bs)) != 0
}


func (woaZ1rb *Akm7q89_) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs bool) (fHza20D bool) {
	return  /*line :1*/ CompareAndSwapUint32(&woaZ1rb.fq6pSyuu,  /*line :1*/ l7Iap5(qVwQ2ijo0),  /*line :1*/ l7Iap5(vxcCAFH4Bs))
}


func l7Iap5(vBjLnvK bool) uint32 {
	if vBjLnvK {
		return 1
	}
	return 0
}



var _ = &ToSaNw[int]{}


type ToSaNw[GFF7Vvaej any] struct {
	
	
	
	_	[0]*GFF7Vvaej

	_	axM70DO7
	sFfRzwz_NRL	unsafe.Pointer
}


func (woaZ1rb *ToSaNw[GFF7Vvaej]) Load() *GFF7Vvaej {
	return (* /*line :1*/ GFF7Vvaej)( /*line :1*/ LoadPointer(&woaZ1rb.sFfRzwz_NRL))
}


func (woaZ1rb *ToSaNw[GFF7Vvaej]) Store(fSC10Rr9Vh *GFF7Vvaej) {
	 /*line :1*/ Fa2hd0oXquKk(&woaZ1rb.sFfRzwz_NRL,  /*line :1*/ unsafe.Pointer(fSC10Rr9Vh))
}


func (woaZ1rb *ToSaNw[GFF7Vvaej]) Swap(vxcCAFH4Bs *GFF7Vvaej) (qVwQ2ijo0 *GFF7Vvaej) {
	return (* /*line :1*/ GFF7Vvaej)( /*line :1*/ X8DfWh(&woaZ1rb.sFfRzwz_NRL,  /*line :1*/ unsafe.Pointer(vxcCAFH4Bs)))
}


func (woaZ1rb *ToSaNw[GFF7Vvaej]) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs *GFF7Vvaej) (fHza20D bool) {
	return  /*line :1*/ B2CfK2wiOXKp(&woaZ1rb.sFfRzwz_NRL,  /*line :1*/ unsafe.Pointer(qVwQ2ijo0),  /*line :1*/ unsafe.Pointer(vxcCAFH4Bs))
}


type JhtCwKEzC struct {
	_	axM70DO7
	eDObqa9Wtw	int32
}


func (woaZ1rb *JhtCwKEzC) Load() int32	{ return  /*line :1*/ LoadInt32(&woaZ1rb.eDObqa9Wtw) }


func (woaZ1rb *JhtCwKEzC) Store(fSC10Rr9Vh int32)	{  /*line :1*/ StoreInt32(&woaZ1rb.eDObqa9Wtw, fSC10Rr9Vh) }


func (woaZ1rb *JhtCwKEzC) Swap(vxcCAFH4Bs int32) (qVwQ2ijo0 int32) {
	return  /*line :1*/ SwapInt32(&woaZ1rb.eDObqa9Wtw, vxcCAFH4Bs)
}


func (woaZ1rb *JhtCwKEzC) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs int32) (fHza20D bool) {
	return  /*line :1*/ CompareAndSwapInt32(&woaZ1rb.eDObqa9Wtw, qVwQ2ijo0, vxcCAFH4Bs)
}


func (woaZ1rb *JhtCwKEzC) Add(u3IYK3UlnB28 int32) (vxcCAFH4Bs int32) {
	return  /*line :1*/ AddInt32(&woaZ1rb.eDObqa9Wtw, u3IYK3UlnB28)
}


type DKr1XV6D2z struct {
	_	axM70DO7
	_	align64
	n9w9LG01	int64
}


func (woaZ1rb *DKr1XV6D2z) Load() int64	{ return  /*line :1*/ LoadInt64(&woaZ1rb.n9w9LG01) }


func (woaZ1rb *DKr1XV6D2z) Store(fSC10Rr9Vh int64)	{  /*line :1*/ StoreInt64(&woaZ1rb.n9w9LG01, fSC10Rr9Vh) }


func (woaZ1rb *DKr1XV6D2z) Swap(vxcCAFH4Bs int64) (qVwQ2ijo0 int64) {
	return  /*line :1*/ SwapInt64(&woaZ1rb.n9w9LG01, vxcCAFH4Bs)
}


func (woaZ1rb *DKr1XV6D2z) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs int64) (fHza20D bool) {
	return  /*line :1*/ CompareAndSwapInt64(&woaZ1rb.n9w9LG01, qVwQ2ijo0, vxcCAFH4Bs)
}


func (woaZ1rb *DKr1XV6D2z) Add(u3IYK3UlnB28 int64) (vxcCAFH4Bs int64) {
	return  /*line :1*/ AddInt64(&woaZ1rb.n9w9LG01, u3IYK3UlnB28)
}


type S6toQany struct {
	_	axM70DO7
	fq6pSyuu	uint32
}


func (woaZ1rb *S6toQany) Load() uint32	{ return  /*line :1*/ LoadUint32(&woaZ1rb.fq6pSyuu) }


func (woaZ1rb *S6toQany) Store(fSC10Rr9Vh uint32)	{  /*line :1*/ StoreUint32(&woaZ1rb.fq6pSyuu, fSC10Rr9Vh) }


func (woaZ1rb *S6toQany) Swap(vxcCAFH4Bs uint32) (qVwQ2ijo0 uint32) {
	return  /*line :1*/ SwapUint32(&woaZ1rb.fq6pSyuu, vxcCAFH4Bs)
}


func (woaZ1rb *S6toQany) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs uint32) (fHza20D bool) {
	return  /*line :1*/ CompareAndSwapUint32(&woaZ1rb.fq6pSyuu, qVwQ2ijo0, vxcCAFH4Bs)
}


func (woaZ1rb *S6toQany) Add(u3IYK3UlnB28 uint32) (vxcCAFH4Bs uint32) {
	return  /*line :1*/ AddUint32(&woaZ1rb.fq6pSyuu, u3IYK3UlnB28)
}


type ZskaiQXr1j struct {
	_	axM70DO7
	_	align64
	rUHLLaJX	uint64
}


func (woaZ1rb *ZskaiQXr1j) Load() uint64	{ return  /*line :1*/ LoadUint64(&woaZ1rb.rUHLLaJX) }


func (woaZ1rb *ZskaiQXr1j) Store(fSC10Rr9Vh uint64)	{  /*line :1*/ StoreUint64(&woaZ1rb.rUHLLaJX, fSC10Rr9Vh) }


func (woaZ1rb *ZskaiQXr1j) Swap(vxcCAFH4Bs uint64) (qVwQ2ijo0 uint64) {
	return  /*line :1*/ SwapUint64(&woaZ1rb.rUHLLaJX, vxcCAFH4Bs)
}


func (woaZ1rb *ZskaiQXr1j) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs uint64) (fHza20D bool) {
	return  /*line :1*/ CompareAndSwapUint64(&woaZ1rb.rUHLLaJX, qVwQ2ijo0, vxcCAFH4Bs)
}


func (woaZ1rb *ZskaiQXr1j) Add(u3IYK3UlnB28 uint64) (vxcCAFH4Bs uint64) {
	return  /*line :1*/ AddUint64(&woaZ1rb.rUHLLaJX, u3IYK3UlnB28)
}


type Lf8Je8V struct {
	_	axM70DO7
	jTX2y7NBZa	uintptr
}


func (woaZ1rb *Lf8Je8V) Load() uintptr	{ return  /*line :1*/ LoadUintptr(&woaZ1rb.jTX2y7NBZa) }


func (woaZ1rb *Lf8Je8V) Store(fSC10Rr9Vh uintptr)	{  /*line :1*/ StoreUintptr(&woaZ1rb.jTX2y7NBZa, fSC10Rr9Vh) }


func (woaZ1rb *Lf8Je8V) Swap(vxcCAFH4Bs uintptr) (qVwQ2ijo0 uintptr) {
	return  /*line :1*/ SwapUintptr(&woaZ1rb.jTX2y7NBZa, vxcCAFH4Bs)
}


func (woaZ1rb *Lf8Je8V) CompareAndSwap(qVwQ2ijo0, vxcCAFH4Bs uintptr) (fHza20D bool) {
	return  /*line :1*/ CompareAndSwapUintptr(&woaZ1rb.jTX2y7NBZa, qVwQ2ijo0, vxcCAFH4Bs)
}


func (woaZ1rb *Lf8Je8V) Add(u3IYK3UlnB28 uintptr) (vxcCAFH4Bs uintptr) {
	return  /*line :1*/ AddUintptr(&woaZ1rb.jTX2y7NBZa, u3IYK3UlnB28)
}








type axM70DO7 struct{}


func (*axM70DO7) Lock()	{}
func (*axM70DO7) Unlock()	{}




type align64 struct{}
