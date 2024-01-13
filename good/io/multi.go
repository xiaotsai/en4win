//line :1
package xI9ai2djaJ2

type cAYD_p6J5 struct{}

func (cAYD_p6J5) Read([]byte) (int, error) {
	return 0, K5Sqco
}

type lACDwjp3Obeb struct {
	bn8uYai2Ls []YJ04Fau
}

func (r4gPeHaMmYh *lACDwjp3Obeb) Read(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error) {
	for  /*line :1*/ len(r4gPeHaMmYh.bn8uYai2Ls) > 0 {

		if  /*line :1*/ len(r4gPeHaMmYh.bn8uYai2Ls) == 1 {
			if wZW51lrVDWS, olKmXo5 := r4gPeHaMmYh.bn8uYai2Ls[0].(*lACDwjp3Obeb); olKmXo5 {
				r4gPeHaMmYh.bn8uYai2Ls = wZW51lrVDWS.bn8uYai2Ls
				continue
			}
		}
		f3MWP_v, fjG4tVMGB =  /*line :1*/ r4gPeHaMmYh.bn8uYai2Ls[0].Read(x6wLQwPqT1n)
		if fjG4tVMGB == K5Sqco {

			r4gPeHaMmYh.bn8uYai2Ls[0] = cAYD_p6J5{}
			r4gPeHaMmYh.bn8uYai2Ls = r4gPeHaMmYh.bn8uYai2Ls[1:]
		}
		if f3MWP_v > 0 || fjG4tVMGB != K5Sqco {
			if fjG4tVMGB == K5Sqco &&  /*line :1*/ len(r4gPeHaMmYh.bn8uYai2Ls) > 0 {

				fjG4tVMGB = nil
			}
			return
		}
	}
	return 0, K5Sqco
}

func (r4gPeHaMmYh *lACDwjp3Obeb) WriteTo(iO5t8eLj5Pd LXQrGW6BQt) (dC5IehVTe8q int64, fjG4tVMGB error) {
	return  /*line :1*/ r4gPeHaMmYh.vNidQe0jR(iO5t8eLj5Pd,  /*line :1*/ make([]byte, 1024*32))
}

func (r4gPeHaMmYh *lACDwjp3Obeb) vNidQe0jR(iO5t8eLj5Pd LXQrGW6BQt, aHfq7_Z []byte) (dC5IehVTe8q int64, fjG4tVMGB error) {
	for _EK5yDV_JjmS, wZW51lrVDWS := range r4gPeHaMmYh.bn8uYai2Ls {
		var f3MWP_v int64
		if jwehERoOOUB, olKmXo5 := wZW51lrVDWS.(*lACDwjp3Obeb); olKmXo5 {
			f3MWP_v, fjG4tVMGB =  /*line :1*/ jwehERoOOUB.vNidQe0jR(iO5t8eLj5Pd, aHfq7_Z)
		} else {
			f3MWP_v, fjG4tVMGB =  /*line :1*/ a0xVlO6(iO5t8eLj5Pd, wZW51lrVDWS, aHfq7_Z)
		}
		dC5IehVTe8q += f3MWP_v
		if fjG4tVMGB != nil {
			r4gPeHaMmYh.bn8uYai2Ls = r4gPeHaMmYh.bn8uYai2Ls[_EK5yDV_JjmS:]
			return dC5IehVTe8q, fjG4tVMGB
		}
		r4gPeHaMmYh.bn8uYai2Ls[_EK5yDV_JjmS] = nil
	}
	r4gPeHaMmYh.bn8uYai2Ls = nil
	return dC5IehVTe8q, nil
}

var _ JCVgLD8ld = (* /*line :1*/ lACDwjp3Obeb)(nil)





func FDbKw7(mpJpyl ...YJ04Fau) YJ04Fau {
	wZW51lrVDWS :=  /*line :1*/ make([]YJ04Fau,  /*line :1*/ len(mpJpyl))
	 /*line :1*/ copy(wZW51lrVDWS, mpJpyl)
	return &lACDwjp3Obeb{wZW51lrVDWS}
}

type iyWaQccQN struct {
	orR6uFH []LXQrGW6BQt
}

func (zsrcBG *iyWaQccQN) Write(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error) {
	for _, iO5t8eLj5Pd := range zsrcBG.orR6uFH {
		f3MWP_v, fjG4tVMGB =  /*line :1*/ iO5t8eLj5Pd.Write(x6wLQwPqT1n)
		if fjG4tVMGB != nil {
			return
		}
		if f3MWP_v !=  /*line :1*/ len(x6wLQwPqT1n) {
			fjG4tVMGB = ArPWDHfv
			return
		}
	}
	return  /*line :1*/ len(x6wLQwPqT1n), nil
}

var _ HpBo0GpMvi6 = (* /*line :1*/ iyWaQccQN)(nil)

func (zsrcBG *iyWaQccQN) WriteString(toLe4u string) (f3MWP_v int, fjG4tVMGB error) {
	var x6wLQwPqT1n []byte	
	for _, iO5t8eLj5Pd := range zsrcBG.orR6uFH {
		if mWtRnqAwj08, olKmXo5 := iO5t8eLj5Pd.(HpBo0GpMvi6); olKmXo5 {
			f3MWP_v, fjG4tVMGB =  /*line :1*/ mWtRnqAwj08.WriteString(toLe4u)
		} else {
			if x6wLQwPqT1n == nil {
				x6wLQwPqT1n = [] /*line :1*/ byte(toLe4u)
			}
			f3MWP_v, fjG4tVMGB =  /*line :1*/ iO5t8eLj5Pd.Write(x6wLQwPqT1n)
		}
		if fjG4tVMGB != nil {
			return
		}
		if f3MWP_v !=  /*line :1*/ len(toLe4u) {
			fjG4tVMGB = ArPWDHfv
			return
		}
	}
	return  /*line :1*/ len(toLe4u), nil
}







func Qhr0ol55AReq(uLR4DVoNc ...LXQrGW6BQt) LXQrGW6BQt {
	ruTHPaZ4 :=  /*line :1*/ make([]LXQrGW6BQt, 0,  /*line :1*/ len(uLR4DVoNc))
	for _, iO5t8eLj5Pd := range uLR4DVoNc {
		if nppDbw, olKmXo5 := iO5t8eLj5Pd.(*iyWaQccQN); olKmXo5 {
			ruTHPaZ4 =  /*line :1*/ append(ruTHPaZ4, nppDbw.orR6uFH...)
		} else {
			ruTHPaZ4 =  /*line :1*/ append(ruTHPaZ4, iO5t8eLj5Pd)
		}
	}
	return &iyWaQccQN{ruTHPaZ4}
}
