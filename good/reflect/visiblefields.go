//line :1
package reflect










func ARUCeSc(sw8_lJ YJh989LTZsVX) []EeSeBImJjDQo {
	if sw8_lJ == nil {
		 /*line :1*/ panic("reflect: VisibleFields(nil)")
	}
	if  /*line :1*/ sw8_lJ.Kind() != Struct {
		 /*line :1*/ panic("reflect.VisibleFields of non-struct type")
	}
	hVMHmIR2 := &m6LGta{
		vRONLh7gSyM:	 /*line :1*/ make(map[string]int),
		rG5hmt:	 /*line :1*/ make(map[YJh989LTZsVX]bool),
		e9RT0Tram:	 /*line :1*/ make([]EeSeBImJjDQo, 0,  /*line :1*/ sw8_lJ.NumField()),
		expotK_rGZr:	 /*line :1*/ make([]int, 0, 2),
	}
	 /*line :1*/ hVMHmIR2.wFiuRZ1_ax7g(sw8_lJ)

	kNVeJaGxs := 0
	for hXZpj1nTZ := range hVMHmIR2.e9RT0Tram {
		jToThzM := &hVMHmIR2.e9RT0Tram[hXZpj1nTZ]
		if jToThzM.ZOSBHkJz == "" {
			continue
		}
		if hXZpj1nTZ != kNVeJaGxs {

			hVMHmIR2.e9RT0Tram[kNVeJaGxs] = *jToThzM
		}
		kNVeJaGxs++
	}
	return hVMHmIR2.e9RT0Tram[:kNVeJaGxs]
}

type m6LGta struct {
	vRONLh7gSyM	map[string]int
	rG5hmt	map[YJh989LTZsVX]bool
	e9RT0Tram	[]EeSeBImJjDQo
	expotK_rGZr	[]int
}






func (hVMHmIR2 *m6LGta) wFiuRZ1_ax7g(sw8_lJ YJh989LTZsVX) {
	if hVMHmIR2.rG5hmt[sw8_lJ] {
		return
	}
	hVMHmIR2.rG5hmt[sw8_lJ] = true
	for hXZpj1nTZ := 0; hXZpj1nTZ <  /*line :1*/ sw8_lJ.NumField(); hXZpj1nTZ++ {
		jToThzM :=  /*line :1*/ sw8_lJ.Field(hXZpj1nTZ)
		hVMHmIR2.expotK_rGZr =  /*line :1*/ append(hVMHmIR2.expotK_rGZr, hXZpj1nTZ)
		lxkP7oV7sm := true
		if ys7szgyZhb, bDfzXlCm5Raf := hVMHmIR2.vRONLh7gSyM[jToThzM.ZOSBHkJz]; bDfzXlCm5Raf {
			abs0ool := &hVMHmIR2.e9RT0Tram[ys7szgyZhb]
			if  /*line :1*/ len(hVMHmIR2.expotK_rGZr) ==  /*line :1*/ len(abs0ool.KOMPhI) {

				abs0ool.ZOSBHkJz = ""
				lxkP7oV7sm = false
			} else if  /*line :1*/ len(hVMHmIR2.expotK_rGZr) <  /*line :1*/ len(abs0ool.KOMPhI) {

				abs0ool.ZOSBHkJz = ""
			} else {

				lxkP7oV7sm = false
			}
		}
		if lxkP7oV7sm {

			jToThzM.KOMPhI =  /*line :1*/ append([] /*line :1*/ int(nil), hVMHmIR2.expotK_rGZr...)
			hVMHmIR2.vRONLh7gSyM[jToThzM.ZOSBHkJz] =  /*line :1*/ len(hVMHmIR2.e9RT0Tram)
			hVMHmIR2.e9RT0Tram =  /*line :1*/ append(hVMHmIR2.e9RT0Tram, jToThzM)
		}
		if jToThzM.GuSnlgXxlhPz {
			if  /*line :1*/ jToThzM.S1miDNBAV8a.Kind() == Pointer {
				jToThzM.S1miDNBAV8a =  /*line :1*/ jToThzM.S1miDNBAV8a.Elem()
			}
			if  /*line :1*/ jToThzM.S1miDNBAV8a.Kind() == Struct {
				 /*line :1*/ hVMHmIR2.wFiuRZ1_ax7g(jToThzM.S1miDNBAV8a)
			}
		}
		hVMHmIR2.expotK_rGZr = hVMHmIR2.expotK_rGZr[: /*line :1*/ len(hVMHmIR2.expotK_rGZr)-1]
	}
	 /*line :1*/ delete(hVMHmIR2.rG5hmt, sw8_lJ)
}
