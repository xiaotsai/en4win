//line :1
package rayT9AgS

import (
	bisect "vWnmvaJKxA"
	"internal/godebugs"
	sync "sync"
	atomic "sync/atomic"
	"unsafe"
	_ "unsafe"
)

type Ea0dEDv5 struct {
	hdYnTN_	string
	hFfgyCN	sync.LhBfwn6wa1x
	*cqyX2z0Byf0h
}

type cqyX2z0Byf0h struct {
	mlq_HOZv	atomic.ToSaNw[lKad2LPTbnA]
	npwRKROfPea	sync.LhBfwn6wa1x
	kR0IwBZMLPa	atomic.ZskaiQXr1j
	e8LTu1	*godebugs.Info
}

type lKad2LPTbnA struct {
	b_tB44	string
	awZC6A2a	*bisect.SiHlqZVr7s
}

func HppHHAII(qaeQa63urGV string) *Ea0dEDv5 {
	return &Ea0dEDv5{hdYnTN_: qaeQa63urGV}
}

func (dKzzxeoR *Ea0dEDv5) Name() string {
	if dKzzxeoR.hdYnTN_ != "" && dKzzxeoR.hdYnTN_[0] == '#' {
		return dKzzxeoR.hdYnTN_[1:]
	}
	return dKzzxeoR.hdYnTN_
}

func (dKzzxeoR *Ea0dEDv5) Undocumented() bool {
	return dKzzxeoR.hdYnTN_ != "" && dKzzxeoR.hdYnTN_[0] == '#'
}

func (dKzzxeoR *Ea0dEDv5) String() string {
	return  /*line :1*/ dKzzxeoR.Name() + "=" +  /*line :1*/ dKzzxeoR.Value()
}

func (dKzzxeoR *Ea0dEDv5) IncNonDefault() {
	 /*line :1*/ dKzzxeoR.npwRKROfPea.Do(dKzzxeoR.sJaNfAVNF9)
	 /*line :1*/ dKzzxeoR.kR0IwBZMLPa.Add(1)
}

func (dKzzxeoR *Ea0dEDv5) sJaNfAVNF9() {
	if dKzzxeoR.e8LTu1 == nil || dKzzxeoR.e8LTu1.Opaque {
		 /*line :1*/ panic("godebug: unexpected IncNonDefault of " + dKzzxeoR.hdYnTN_)
	}
	 /*line :1*/ sAEqBQbvE6("/godebug/non-default-behavior/"+ /*line :1*/ dKzzxeoR.Name()+":events", dKzzxeoR.kR0IwBZMLPa.Load)
}

var gHGYJ95T sync.Eim1b6bNi9IM

var qdaVHCb lKad2LPTbnA

func (dKzzxeoR *Ea0dEDv5) Value() string {
	 /*line :1*/ dKzzxeoR.hFfgyCN.Do(func() {
		dKzzxeoR.cqyX2z0Byf0h =  /*line :1*/ fwLcyCtsaV8( /*line :1*/ dKzzxeoR.Name())
		if dKzzxeoR.e8LTu1 == nil && ! /*line :1*/ dKzzxeoR.Undocumented() {
			 /*line :1*/ panic("godebug: Value of name not listed in godebugs.All: " + dKzzxeoR.hdYnTN_)
		}
	})
	vePGy_75 := * /*line :1*/ dKzzxeoR.mlq_HOZv.Load()
	if vePGy_75.awZC6A2a != nil && ! /*line :1*/ vePGy_75.awZC6A2a.Stack(&mUHPjxscoUj) {
		return ""
	}
	return vePGy_75.b_tB44
}

func fwLcyCtsaV8(qaeQa63urGV string) *cqyX2z0Byf0h {
	if vePGy_75, sE3ygpW :=  /*line :1*/ gHGYJ95T.Load(qaeQa63urGV); sE3ygpW {
		return vePGy_75.(*cqyX2z0Byf0h)
	}
	dKzzxeoR :=  /*line :1*/ new(cqyX2z0Byf0h)
	dKzzxeoR.e8LTu1 =  /*line :1*/ godebugs.Lookup(qaeQa63urGV)
	 /*line :1*/ dKzzxeoR.mlq_HOZv.Store(&qdaVHCb)
	if vePGy_75, xZ_aB_rAmWf :=  /*line :1*/ gHGYJ95T.LoadOrStore(qaeQa63urGV, dKzzxeoR); xZ_aB_rAmWf {

		return vePGy_75.(*cqyX2z0Byf0h)
	}

	return dKzzxeoR
}

//go:linkname ra_sJbHdPGkQ
func ra_sJbHdPGkQ(slPom4r2V func(string, string))

//go:linkname sAEqBQbvE6
func sAEqBQbvE6(qaeQa63urGV string, _H7ZZxyUK func() uint64)

//go:linkname vTCjjoHu
func vTCjjoHu(czT1dPhVHbm func(string) func())

func init() {
	 /*line :1*/ ra_sJbHdPGkQ(slPom4r2V)
	 /*line :1*/ vTCjjoHu(czT1dPhVHbm)
}

func czT1dPhVHbm(qaeQa63urGV string) func() {
	dKzzxeoR :=  /*line :1*/ HppHHAII(qaeQa63urGV)
	 /*line :1*/ dKzzxeoR.Value()
	return dKzzxeoR.IncNonDefault
}

var j8O2vxNJmtb sync.DIRWe11KYlYa

func slPom4r2V(xc5qSnBGOAm, jY7petkWw4 string) {
	 /*line :1*/ j8O2vxNJmtb.Lock()
	defer  /*line :1*/ j8O2vxNJmtb.Unlock()

	tgYrmxQ6WI :=  /*line :1*/ make(map[string]bool)
	 /*line :1*/ e2z1hpZKk(tgYrmxQ6WI, jY7petkWw4)
	 /*line :1*/ e2z1hpZKk(tgYrmxQ6WI, xc5qSnBGOAm)

	 /*line :1*/ gHGYJ95T.Range(func(qaeQa63urGV, dKzzxeoR any) bool {
		if !tgYrmxQ6WI[qaeQa63urGV.(string)] {
			 /*line :1*/ dKzzxeoR.(*cqyX2z0Byf0h).mlq_HOZv.Store(&qdaVHCb)
		}
		return true
	})
}

func e2z1hpZKk(tgYrmxQ6WI map[string]bool, dKzzxeoR string) {

	ui7kr0Dre :=  /*line :1*/ len(dKzzxeoR)
	mL3C4q := -1
	for e1KWWd2X := ui7kr0Dre - 1; e1KWWd2X >= -1; e1KWWd2X-- {
		if e1KWWd2X == -1 || dKzzxeoR[e1KWWd2X] == ',' {
			if mL3C4q >= 0 {
				qaeQa63urGV, n6BSjRKde := dKzzxeoR[e1KWWd2X+1:mL3C4q], dKzzxeoR[mL3C4q+1:ui7kr0Dre]
				if !tgYrmxQ6WI[qaeQa63urGV] {
					tgYrmxQ6WI[qaeQa63urGV] = true
					vePGy_75 := &lKad2LPTbnA{b_tB44: n6BSjRKde}
					for _AJhgQ7 := 0; _AJhgQ7 <  /*line :1*/ len(n6BSjRKde); _AJhgQ7++ {
						if n6BSjRKde[_AJhgQ7] == '#' {
							vePGy_75.b_tB44 = n6BSjRKde[:_AJhgQ7]
							vePGy_75.awZC6A2a, _ =  /*line :1*/ bisect.MbSZaWBD(n6BSjRKde[_AJhgQ7+1:])
							break
						}
					}
					 /*line :1*/ fwLcyCtsaV8(qaeQa63urGV).mlq_HOZv.Store(vePGy_75)
				}
			}
			mL3C4q = -1
			ui7kr0Dre = e1KWWd2X
		} else if dKzzxeoR[e1KWWd2X] == '=' {
			mL3C4q = e1KWWd2X
		}
	}
}

type vNdCZb struct{}

var mUHPjxscoUj vNdCZb

func (*vNdCZb) Write(talRRy2KEKID []byte) (int, error) {
	if  /*line :1*/ len(talRRy2KEKID) > 0 {
		 /*line :1*/ aJdiYG(2,  /*line :1*/ unsafe.Pointer(&talRRy2KEKID[0]),  /*line :1*/ int32( /*line :1*/ len(talRRy2KEKID)))
	}
	return  /*line :1*/ len(talRRy2KEKID), nil
}

//go:linkname aJdiYG runtime.write
func aJdiYG(laUZYJTUzW uintptr, uw3vtdq unsafe.Pointer, kinRt1 int32) int32
