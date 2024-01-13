//line :1
package fQXlzv






type adgXfXFYJr14 struct {
	
	jmfRJJPOO	string

	
	
	
	
	
	
	
	hMt0X0	[256]int

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	izBrNUhYfRa	[]int
}

func nyd62y(iVumsgT string) *adgXfXFYJr14 {
	vtYnd8LwzVN := &adgXfXFYJr14{
		jmfRJJPOO:	iVumsgT,
		izBrNUhYfRa:	 /*line :1*/ make([]int,  /*line :1*/ len(iVumsgT)),
	}

	z6P4_U4 :=  /*line :1*/ len(iVumsgT) - 1

	for gUiH7N := range vtYnd8LwzVN.hMt0X0 {
		vtYnd8LwzVN.hMt0X0[gUiH7N] =  /*line :1*/ len(iVumsgT)
	}

	for gUiH7N := 0; gUiH7N < z6P4_U4; gUiH7N++ {
		vtYnd8LwzVN.hMt0X0[iVumsgT[gUiH7N]] = z6P4_U4 - gUiH7N
	}

	dTxoufkz := z6P4_U4
	for gUiH7N := z6P4_U4; gUiH7N >= 0; gUiH7N-- {
		if  /*line :1*/ E8ZCzfgw(iVumsgT, iVumsgT[gUiH7N+1:]) {
			dTxoufkz = gUiH7N + 1
		}

		vtYnd8LwzVN.izBrNUhYfRa[gUiH7N] = dTxoufkz + z6P4_U4 - gUiH7N
	}

	for gUiH7N := 0; gUiH7N < z6P4_U4; gUiH7N++ {
		elB7c_s40A :=  /*line :1*/ nGlMnB4oWO(iVumsgT, iVumsgT[1:gUiH7N+1])
		if iVumsgT[gUiH7N-elB7c_s40A] != iVumsgT[z6P4_U4-elB7c_s40A] {

			vtYnd8LwzVN.izBrNUhYfRa[z6P4_U4-elB7c_s40A] = elB7c_s40A + z6P4_U4 - gUiH7N
		}
	}

	return vtYnd8LwzVN
}

func nGlMnB4oWO(jZJLDz0ImV, naCMgzayVI string) (gUiH7N int) {
	for ; gUiH7N <  /*line :1*/ len(jZJLDz0ImV) && gUiH7N <  /*line :1*/ len(naCMgzayVI); gUiH7N++ {
		if jZJLDz0ImV[ /*line :1*/ len(jZJLDz0ImV)-1-gUiH7N] != naCMgzayVI[ /*line :1*/ len(naCMgzayVI)-1-gUiH7N] {
			break
		}
	}
	return
}



func (vtYnd8LwzVN *adgXfXFYJr14) eZ4JDrZAVG(mbc5dNTvLTa string) int {
	gUiH7N :=  /*line :1*/ len(vtYnd8LwzVN.jmfRJJPOO) - 1
	for gUiH7N <  /*line :1*/ len(mbc5dNTvLTa) {

		xXcc3ZyUB3kD :=  /*line :1*/ len(vtYnd8LwzVN.jmfRJJPOO) - 1
		for xXcc3ZyUB3kD >= 0 && mbc5dNTvLTa[gUiH7N] == vtYnd8LwzVN.jmfRJJPOO[xXcc3ZyUB3kD] {
			gUiH7N--
			xXcc3ZyUB3kD--
		}
		if xXcc3ZyUB3kD < 0 {
			return gUiH7N + 1
		}
		gUiH7N +=  /*line :1*/ rWOw1jVt(vtYnd8LwzVN.hMt0X0[mbc5dNTvLTa[gUiH7N]], vtYnd8LwzVN.izBrNUhYfRa[xXcc3ZyUB3kD])
	}
	return -1
}

func rWOw1jVt(jZJLDz0ImV, naCMgzayVI int) int {
	if jZJLDz0ImV > naCMgzayVI {
		return jZJLDz0ImV
	}
	return naCMgzayVI
}
