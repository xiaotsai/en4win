//line :1
package GWGwIQr9vXAy

import (
	godebug "rayT9AgS"
	"runtime"
	sync "sync"
	"unsafe"
)

type XGxgZFq struct {
	_	[0]func()
	ajnER4	any

	dszaJKgLKdw	bool
}

func (cdmUzIMbO4 *XGxgZFq) Get() any	{ return cdmUzIMbO4.ajnER4 }

type wSvrJIadzhYU struct {
	a_GqnD6g	string
	nJ7JqO	any

	g14XeuBQpfmL	bool
}

func pNutxWu(zPlPHS any) wSvrJIadzhYU {
	if dOj1sQk9D6dR, feYwCV := zPlPHS.(string); feYwCV {
		return wSvrJIadzhYU{a_GqnD6g: dOj1sQk9D6dR, g14XeuBQpfmL: true}
	}
	return wSvrJIadzhYU{nJ7JqO: zPlPHS}
}

func (t50DSm4 wSvrJIadzhYU) Value() *XGxgZFq {
	if t50DSm4.g14XeuBQpfmL {
		return &XGxgZFq{ajnER4: t50DSm4.a_GqnD6g}
	}
	return &XGxgZFq{ajnER4: t50DSm4.nJ7JqO}
}

var (
	hHhWfyd4_wz	sync.DIRWe11KYlYa
	q45vOvU	= map[wSvrJIadzhYU]uintptr{}
	k3h8dwdAWUB4	=  /*line :1*/ guqf7yR222f()
)

var qLJwgw8JFdyd =  /*line :1*/ godebug.HppHHAII("#intern")

func guqf7yR222f() map[wSvrJIadzhYU]*XGxgZFq {
	if  /*line :1*/ qLJwgw8JFdyd.Value() == "leaky" {
		return map[wSvrJIadzhYU]*XGxgZFq{}
	}
	return nil
}

func H2f7aKP7(zPlPHS any) *XGxgZFq {
	return  /*line :1*/ cZLCKaYFsuPM( /*line :1*/ pNutxWu(zPlPHS))
}

func DocbIWHFe(dOj1sQk9D6dR string) *XGxgZFq {
	return  /*line :1*/ cZLCKaYFsuPM(wSvrJIadzhYU{a_GqnD6g: dOj1sQk9D6dR, g14XeuBQpfmL: true})
}

//go:nocheckptr
func cZLCKaYFsuPM(t50DSm4 wSvrJIadzhYU) *XGxgZFq {
	 /*line :1*/ hHhWfyd4_wz.Lock()
	defer  /*line :1*/ hHhWfyd4_wz.Unlock()

	var cdmUzIMbO4 *XGxgZFq
	if k3h8dwdAWUB4 != nil {
		cdmUzIMbO4 = k3h8dwdAWUB4[t50DSm4]
	} else if aEwaGAg, feYwCV := q45vOvU[t50DSm4]; feYwCV {
		cdmUzIMbO4 = (* /*line :1*/ XGxgZFq)( /*line :1*/ unsafe.Pointer(aEwaGAg))
		cdmUzIMbO4.dszaJKgLKdw = true
	}
	if cdmUzIMbO4 != nil {
		return cdmUzIMbO4
	}
	cdmUzIMbO4 =  /*line :1*/ t50DSm4.Value()
	if k3h8dwdAWUB4 != nil {
		k3h8dwdAWUB4[t50DSm4] = cdmUzIMbO4
	} else {

		 /*line :1*/ runtime.SetFinalizer(cdmUzIMbO4, f2ZBUSbXsZl)
		q45vOvU[t50DSm4] =  /*line :1*/ uintptr( /*line :1*/ unsafe.Pointer(cdmUzIMbO4))
	}
	return cdmUzIMbO4
}

func f2ZBUSbXsZl(cdmUzIMbO4 *XGxgZFq) {
	 /*line :1*/ hHhWfyd4_wz.Lock()
	defer  /*line :1*/ hHhWfyd4_wz.Unlock()
	if cdmUzIMbO4.dszaJKgLKdw {

		cdmUzIMbO4.dszaJKgLKdw = false
		 /*line :1*/ runtime.SetFinalizer(cdmUzIMbO4, f2ZBUSbXsZl)
		return
	}
	 /*line :1*/ delete(q45vOvU,  /*line :1*/ pNutxWu(cdmUzIMbO4.ajnER4))
}
