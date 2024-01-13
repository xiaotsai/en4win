//line :1
package fRAfQd_

import (
	errors "iAHoxjmM"
	registry "GMNBnK"
	syscall "bUKeamF"
)

var maKnDxXjO []string	




func ioOVOIImE4t(bJQzPjVfN registry.WA26dh, wRBdSt string, hPSaEngtiUY7, arA2Zt string) (nZWjzarTpPU bool, e9JGvC7UzC error) {
	pd7DMf8ErHa6, xuMLYrB :=  /*line :1*/ registry.J2fxYepRwi3b(bJQzPjVfN, wRBdSt, registry.READ)
	if xuMLYrB != nil {
		return false, xuMLYrB
	}
	defer  /*line :1*/ pd7DMf8ErHa6.Close()

	var gA16DoaF_, fqD03X99 string
	if xuMLYrB =  /*line :1*/ registry.T9vERTZ0gyM(); xuMLYrB == nil {

		gA16DoaF_, xuMLYrB =  /*line :1*/ pd7DMf8ErHa6.GetMUIStringValue("MUI_Std")
		if xuMLYrB == nil {
			fqD03X99, xuMLYrB =  /*line :1*/ pd7DMf8ErHa6.GetMUIStringValue("MUI_Dlt")
		}
	}
	if xuMLYrB != nil {
		if gA16DoaF_, _, xuMLYrB =  /*line :1*/ pd7DMf8ErHa6.GetStringValue("Std"); xuMLYrB != nil {
			return false, xuMLYrB
		}
		if fqD03X99, _, xuMLYrB =  /*line :1*/ pd7DMf8ErHa6.GetStringValue("Dlt"); xuMLYrB != nil {
			return false, xuMLYrB
		}
	}

	if gA16DoaF_ != hPSaEngtiUY7 {
		return false, nil
	}
	if fqD03X99 != arA2Zt && arA2Zt != hPSaEngtiUY7 {
		return false, nil
	}
	return true, nil
}



func _jBAYHDSBbVS(hPSaEngtiUY7, arA2Zt string) (string, error) {
	pd7DMf8ErHa6, xuMLYrB :=  /*line :1*/ registry.J2fxYepRwi3b(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)
	if xuMLYrB != nil {
		return "", xuMLYrB
	}
	defer  /*line :1*/ pd7DMf8ErHa6.Close()

	bJSojDVp, xuMLYrB :=  /*line :1*/ pd7DMf8ErHa6.ReadSubKeyNames()
	if xuMLYrB != nil {
		return "", xuMLYrB
	}
	for _, w12CVdh9 := range bJSojDVp {
		nZWjzarTpPU, xuMLYrB :=  /*line :1*/ ioOVOIImE4t(pd7DMf8ErHa6, w12CVdh9, hPSaEngtiUY7, arA2Zt)
		if xuMLYrB == nil && nZWjzarTpPU {
			return w12CVdh9, nil
		}
	}
	return "",  /*line :1*/ errors.Su6g6hRoi9X(`English name for time zone "` + hPSaEngtiUY7 + `" not found in registry`)
}


func _7n4iQBk(evpVLsp string) string {
	var kZKdhP []rune
	for _, fnpIad3KB := range evpVLsp {
		if 'A' <= fnpIad3KB && fnpIad3KB <= 'Z' {
			kZKdhP =  /*line :1*/ append(kZKdhP, fnpIad3KB)
		}
	}
	return  /*line :1*/ string(kZKdhP)
}


func c91HWD_(iL2jRw3Faxe *syscall.MAPVxQZiq) (gA16DoaF_, cKGX2u string) {
	rOWD9G :=  /*line :1*/ syscall.AODVXp8NM3sd(iL2jRw3Faxe.KIcaL8MMs[:])
	dA2mAk3n5rZH, ufWfJe3qPw := gBrsuXn[rOWD9G]
	if !ufWfJe3qPw {
		edZ1_fT0 :=  /*line :1*/ syscall.AODVXp8NM3sd(iL2jRw3Faxe.VHZAJ_MP[:])

		r4ALovtGO1f, xuMLYrB :=  /*line :1*/ _jBAYHDSBbVS(rOWD9G, edZ1_fT0)
		if xuMLYrB == nil {
			dA2mAk3n5rZH, ufWfJe3qPw = gBrsuXn[r4ALovtGO1f]
			if ufWfJe3qPw {
				return dA2mAk3n5rZH.h02_vPEeg, dA2mAk3n5rZH.cwlKO21
			}
		}

		return  /*line :1*/ _7n4iQBk(rOWD9G),  /*line :1*/ _7n4iQBk(edZ1_fT0)
	}
	return dA2mAk3n5rZH.h02_vPEeg, dA2mAk3n5rZH.cwlKO21
}




func h3paEUV(mO2sH8hUHnW int, lNAwkbJQVBgU *syscall.H6Qxv2jW) int64 {

	gv5Z468R := 1
	pv6eev :=  /*line :1*/ FaDSoi2w(mO2sH8hUHnW,  /*line :1*/ ZyPFXNxLcwpT(lNAwkbJQVBgU.EQp3MHA11V), gv5Z468R,  /*line :1*/ int(lNAwkbJQVBgU.CySfZqD),  /*line :1*/ int(lNAwkbJQVBgU.TFSCFS3IfPS0),  /*line :1*/ int(lNAwkbJQVBgU.MRSnx59), 0, PD1NIUjTJ)
	eUo6AxKM7 :=  /*line :1*/ int(lNAwkbJQVBgU.MslYdlEo) -  /*line :1*/ int( /*line :1*/ pv6eev.Weekday())
	if eUo6AxKM7 < 0 {
		eUo6AxKM7 += 7
	}
	gv5Z468R += eUo6AxKM7
	if cDtaLw1Mj :=  /*line :1*/ int(lNAwkbJQVBgU.HF4FBjj) - 1; cDtaLw1Mj < 4 {
		gv5Z468R += cDtaLw1Mj * 7
	} else {

		gv5Z468R += 4 * 7
		if gv5Z468R >  /*line :1*/ rs4bvng4W6y( /*line :1*/ ZyPFXNxLcwpT(lNAwkbJQVBgU.EQp3MHA11V), mO2sH8hUHnW) {
			gv5Z468R -= 7
		}
	}
	return  /*line :1*/ pv6eev.insjC2Va() +  /*line :1*/ int64(gv5Z468R-1)*secondsPerDay + internalToUnix
}

func lvIq3p8J_rTi(eUo6AxKM7 *syscall.MAPVxQZiq) {
	ipt7urdFIkH := &s45BJX

	ipt7urdFIkH.yiQ5l0TL_ = "Local"

	b4zZoH2FE1 := 1
	if eUo6AxKM7.NRJvDzhCOoT.EQp3MHA11V > 0 {
		b4zZoH2FE1++
	}
	ipt7urdFIkH.beFgWpkl8XT0 =  /*line :1*/ make([]k3sFK3l, b4zZoH2FE1)

	hPSaEngtiUY7, arA2Zt :=  /*line :1*/ c91HWD_(eUo6AxKM7)

	gA16DoaF_ := &ipt7urdFIkH.beFgWpkl8XT0[0]
	gA16DoaF_.sn0AXnjD = hPSaEngtiUY7
	if b4zZoH2FE1 == 1 {

		gA16DoaF_.a8VzUELX = - /*line :1*/ int(eUo6AxKM7.StUAtyXP) * 60
		ipt7urdFIkH.tIAdDDIq = alpha
		ipt7urdFIkH.cmASbbDbiav = omega
		ipt7urdFIkH.j9a8yEma = gA16DoaF_
		ipt7urdFIkH.kAx6aUL =  /*line :1*/ make([]ep4v_A, 1)
		ipt7urdFIkH.kAx6aUL[0].onK9uWf7T8 = ipt7urdFIkH.tIAdDDIq
		ipt7urdFIkH.kAx6aUL[0].kRTXiPL = 0
		return
	}

	gA16DoaF_.a8VzUELX = - /*line :1*/ int(eUo6AxKM7.StUAtyXP+eUo6AxKM7.AW4StnUjhag8) * 60

	cKGX2u := &ipt7urdFIkH.beFgWpkl8XT0[1]
	cKGX2u.sn0AXnjD = arA2Zt
	cKGX2u.a8VzUELX = - /*line :1*/ int(eUo6AxKM7.StUAtyXP+eUo6AxKM7.UAJo4xpW) * 60
	cKGX2u.eXRNnAc1aY = true

	ckWaYo1 := &eUo6AxKM7.NRJvDzhCOoT
	kIMGbZ := &eUo6AxKM7.WZFoaz8
	en2OvQrnxz := 0
	agO96Ye := 1
	if ckWaYo1.EQp3MHA11V > kIMGbZ.EQp3MHA11V {
		ckWaYo1, kIMGbZ = kIMGbZ, ckWaYo1
		en2OvQrnxz, agO96Ye = agO96Ye, en2OvQrnxz
	}

	ipt7urdFIkH.kAx6aUL =  /*line :1*/ make([]ep4v_A, 400)

	pv6eev :=  /*line :1*/ Pc_35oTY().UTC()
	mO2sH8hUHnW :=  /*line :1*/ pv6eev.Year()
	oh_jjqY4Q := 0
	for bRAhk8mArR1z := mO2sH8hUHnW - 100; bRAhk8mArR1z < mO2sH8hUHnW+100; bRAhk8mArR1z++ {
		e7MbtYTeSIq := &ipt7urdFIkH.kAx6aUL[oh_jjqY4Q]
		e7MbtYTeSIq.onK9uWf7T8 =  /*line :1*/ h3paEUV(bRAhk8mArR1z, ckWaYo1) -  /*line :1*/ int64(ipt7urdFIkH.beFgWpkl8XT0[agO96Ye].a8VzUELX)
		e7MbtYTeSIq.kRTXiPL =  /*line :1*/ uint8(en2OvQrnxz)
		oh_jjqY4Q++

		e7MbtYTeSIq = &ipt7urdFIkH.kAx6aUL[oh_jjqY4Q]
		e7MbtYTeSIq.onK9uWf7T8 =  /*line :1*/ h3paEUV(bRAhk8mArR1z, kIMGbZ) -  /*line :1*/ int64(ipt7urdFIkH.beFgWpkl8XT0[en2OvQrnxz].a8VzUELX)
		e7MbtYTeSIq.kRTXiPL =  /*line :1*/ uint8(agO96Ye)
		oh_jjqY4Q++
	}
}

var sT8b42I = syscall.MAPVxQZiq{
	StUAtyXP:	8 * 60,
	KIcaL8MMs: [32]uint16{
		'P', 'a', 'c', 'i', 'f', 'i', 'c', ' ', 'S', 't', 'a', 'n', 'd', 'a', 'r', 'd', ' ', 'T', 'i', 'm', 'e',
	},
	NRJvDzhCOoT:	syscall.H6Qxv2jW{EQp3MHA11V: 11, HF4FBjj: 1, CySfZqD: 2},
	VHZAJ_MP: [32]uint16{
		'P', 'a', 'c', 'i', 'f', 'i', 'c', ' ', 'D', 'a', 'y', 'l', 'i', 'g', 'h', 't', ' ', 'T', 'i', 'm', 'e',
	},
	WZFoaz8:	syscall.H6Qxv2jW{EQp3MHA11V: 3, HF4FBjj: 2, CySfZqD: 2},
	UAJo4xpW:	-60,
}

var luk8V4CIRV = syscall.MAPVxQZiq{
	StUAtyXP:	-10 * 60,
	KIcaL8MMs: [32]uint16{
		'A', 'U', 'S', ' ', 'E', 'a', 's', 't', 'e', 'r', 'n', ' ', 'S', 't', 'a', 'n', 'd', 'a', 'r', 'd', ' ', 'T', 'i', 'm', 'e',
	},
	NRJvDzhCOoT:	syscall.H6Qxv2jW{EQp3MHA11V: 4, HF4FBjj: 1, CySfZqD: 3},
	VHZAJ_MP: [32]uint16{
		'A', 'U', 'S', ' ', 'E', 'a', 's', 't', 'e', 'r', 'n', ' ', 'D', 'a', 'y', 'l', 'i', 'g', 'h', 't', ' ', 'T', 'i', 'm', 'e',
	},
	WZFoaz8:	syscall.H6Qxv2jW{EQp3MHA11V: 10, HF4FBjj: 1, CySfZqD: 2},
	UAJo4xpW:	-60,
}

func aflnY9OXPZ() {
	var eUo6AxKM7 syscall.MAPVxQZiq
	if _, xuMLYrB :=  /*line :1*/ syscall.BbXWjqCVp(&eUo6AxKM7); xuMLYrB != nil {
		s45BJX.yiQ5l0TL_ = "UTC"
		return
	}
	 /*line :1*/ lvIq3p8J_rTi(&eUo6AxKM7)
}
