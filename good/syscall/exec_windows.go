//line :1
package bUKeamF

import (
	"internal/bytealg"
	"runtime"
	sync "sync"
	utf16 "DtJCLKevRp"
	"unsafe"
)

var OLKdSMhcsd sync.V1sRgjwAglJt










func FgsNLhd6a3(wzPhJhIFoI string) string {
	if  /*line :1*/ len(wzPhJhIFoI) == 0 {
		return `""`
	}
	for hA99q3EOK := 0; hA99q3EOK <  /*line :1*/ len(wzPhJhIFoI); hA99q3EOK++ {
		switch wzPhJhIFoI[hA99q3EOK] {
		case '"', '\\', ' ', '\t':

			tHPwhuTh :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(wzPhJhIFoI)+2)
			tHPwhuTh =  /*line :1*/ e5g67RRf(tHPwhuTh, wzPhJhIFoI)
			return  /*line :1*/ string(tHPwhuTh)
		}
	}
	return wzPhJhIFoI
}



func e5g67RRf(tHPwhuTh []byte, wzPhJhIFoI string) []byte {
	if  /*line :1*/ len(wzPhJhIFoI) == 0 {
		return  /*line :1*/ append(tHPwhuTh, `""`...)
	}

	jraXHN1f := false
	bDX_Hc7WUhSh := false
	for hA99q3EOK := 0; hA99q3EOK <  /*line :1*/ len(wzPhJhIFoI); hA99q3EOK++ {
		switch wzPhJhIFoI[hA99q3EOK] {
		case '"', '\\':
			jraXHN1f = true
		case ' ', '\t':
			bDX_Hc7WUhSh = true
		}
	}

	if !jraXHN1f && !bDX_Hc7WUhSh {

		return  /*line :1*/ append(tHPwhuTh, wzPhJhIFoI...)
	}
	if !jraXHN1f {

		tHPwhuTh =  /*line :1*/ append(tHPwhuTh, '"')
		tHPwhuTh =  /*line :1*/ append(tHPwhuTh, wzPhJhIFoI...)
		return  /*line :1*/ append(tHPwhuTh, '"')
	}

	if bDX_Hc7WUhSh {
		tHPwhuTh =  /*line :1*/ append(tHPwhuTh, '"')
	}
	boq1rKiaT0l := 0
	for hA99q3EOK := 0; hA99q3EOK <  /*line :1*/ len(wzPhJhIFoI); hA99q3EOK++ {
		qs8DKg := wzPhJhIFoI[hA99q3EOK]
		switch qs8DKg {
		default:
			boq1rKiaT0l = 0
		case '\\':
			boq1rKiaT0l++
		case '"':
			for ; boq1rKiaT0l > 0; boq1rKiaT0l-- {
				tHPwhuTh =  /*line :1*/ append(tHPwhuTh, '\\')
			}
			tHPwhuTh =  /*line :1*/ append(tHPwhuTh, '\\')
		}
		tHPwhuTh =  /*line :1*/ append(tHPwhuTh, qs8DKg)
	}
	if bDX_Hc7WUhSh {
		for ; boq1rKiaT0l > 0; boq1rKiaT0l-- {
			tHPwhuTh =  /*line :1*/ append(tHPwhuTh, '\\')
		}
		tHPwhuTh =  /*line :1*/ append(tHPwhuTh, '"')
	}

	return tHPwhuTh
}



func a_cROM(dqAabTiua []string) string {
	var tHPwhuTh []byte
	for _, b0abCyNZ := range dqAabTiua {
		if  /*line :1*/ len(tHPwhuTh) > 0 {
			tHPwhuTh =  /*line :1*/ append(tHPwhuTh, ' ')
		}
		tHPwhuTh =  /*line :1*/ e5g67RRf(tHPwhuTh, b0abCyNZ)
	}
	return  /*line :1*/ string(tHPwhuTh)
}






func w4osefwcSL1K(n924JoWO []string) (*uint16, error) {
	if  /*line :1*/ len(n924JoWO) == 0 {
		return & /*line :1*/ utf16.J8eAwOfFc([] /*line :1*/ rune("\x00\x00"))[0], nil
	}
	cne_74 := 0
	for _, wzPhJhIFoI := range n924JoWO {
		if  /*line :1*/ bytealg.IndexByteString(wzPhJhIFoI, 0) != -1 {
			return nil, EINVAL
		}
		cne_74 +=  /*line :1*/ len(wzPhJhIFoI) + 1
	}
	cne_74 += 1

	tHPwhuTh :=  /*line :1*/ make([]byte, cne_74)
	hA99q3EOK := 0
	for _, wzPhJhIFoI := range n924JoWO {
		cE4tZhPbq :=  /*line :1*/ len(wzPhJhIFoI)
		 /*line :1*/ copy(tHPwhuTh[hA99q3EOK:hA99q3EOK+cE4tZhPbq], [] /*line :1*/ byte(wzPhJhIFoI))
		 /*line :1*/ copy(tHPwhuTh[hA99q3EOK+cE4tZhPbq:hA99q3EOK+cE4tZhPbq+1], []byte{0})
		hA99q3EOK = hA99q3EOK + cE4tZhPbq + 1
	}
	 /*line :1*/ copy(tHPwhuTh[hA99q3EOK:hA99q3EOK+1], []byte{0})

	return & /*line :1*/ utf16.J8eAwOfFc([] /*line :1*/ rune( /*line :1*/ string(tHPwhuTh)))[0], nil
}

func W6smj_(dBn476FVCy2C SRlvVjrYa) {
	 /*line :1*/ LUuQqH65KsfT( /*line :1*/ SRlvVjrYa(dBn476FVCy2C), HANDLE_FLAG_INHERIT, 0)
}

func EaTugJTBdQV(dBn476FVCy2C SRlvVjrYa, jCK7wzP3_ bool) (gOCcQzbcC error) {
	return nil
}


func LmpGp3wH_T(hzeU8Tl string) (ro5NCL9 string, gOCcQzbcC error) {
	rd6leevODT, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(hzeU8Tl)
	if gOCcQzbcC != nil {
		return "", gOCcQzbcC
	}
	m5Tq_PL7 :=  /*line :1*/ uint32(100)
	for {
		eun1Jlud5A :=  /*line :1*/ make([]uint16, m5Tq_PL7)
		m5Tq_PL7, gOCcQzbcC =  /*line :1*/ T0pUSJXE6Z(rd6leevODT,  /*line :1*/ uint32( /*line :1*/ len(eun1Jlud5A)), &eun1Jlud5A[0], nil)
		if gOCcQzbcC != nil {
			return "", gOCcQzbcC
		}
		if m5Tq_PL7 <=  /*line :1*/ uint32( /*line :1*/ len(eun1Jlud5A)) {
			return  /*line :1*/ AODVXp8NM3sd(eun1Jlud5A[:m5Tq_PL7]), nil
		}
	}
}

func ugu28bu(qs8DKg uint8) bool {
	return qs8DKg == '\\' || qs8DKg == '/'
}

func dT5PdkB5a(hY26vae8L string) (hzeU8Tl string, gOCcQzbcC error) {
	x9MQeAkOYDq, gOCcQzbcC :=  /*line :1*/ LmpGp3wH_T(hY26vae8L)
	if gOCcQzbcC != nil {
		return "", gOCcQzbcC
	}
	if  /*line :1*/ len(x9MQeAkOYDq) > 2 &&  /*line :1*/ ugu28bu(x9MQeAkOYDq[0]) &&  /*line :1*/ ugu28bu(x9MQeAkOYDq[1]) {

		return "", EINVAL
	}
	return x9MQeAkOYDq, nil
}

func iWAzJxBb1(d6LKXEr3A int) int {
	if 'a' <= d6LKXEr3A && d6LKXEr3A <= 'z' {
		d6LKXEr3A += 'A' - 'a'
	}
	return d6LKXEr3A
}

func iCvgqkkKz(hY26vae8L, rd6leevODT string) (hzeU8Tl string, gOCcQzbcC error) {
	if  /*line :1*/ len(rd6leevODT) == 0 {
		return "", EINVAL
	}
	if  /*line :1*/ len(rd6leevODT) > 2 &&  /*line :1*/ ugu28bu(rd6leevODT[0]) &&  /*line :1*/ ugu28bu(rd6leevODT[1]) {

		return rd6leevODT, nil
	}
	if  /*line :1*/ len(rd6leevODT) > 1 && rd6leevODT[1] == ':' {

		if  /*line :1*/ len(rd6leevODT) == 2 {
			return "", EINVAL
		}
		if  /*line :1*/ ugu28bu(rd6leevODT[2]) {
			return rd6leevODT, nil
		} else {
			vC51I7RbKXXN, gOCcQzbcC :=  /*line :1*/ dT5PdkB5a(hY26vae8L)
			if gOCcQzbcC != nil {
				return "", gOCcQzbcC
			}
			if  /*line :1*/ iWAzJxBb1( /*line :1*/ int(rd6leevODT[0])) ==  /*line :1*/ iWAzJxBb1( /*line :1*/ int(vC51I7RbKXXN[0])) {
				return  /*line :1*/ LmpGp3wH_T(vC51I7RbKXXN + "\\" + rd6leevODT[2:])
			} else {
				return  /*line :1*/ LmpGp3wH_T(rd6leevODT)
			}
		}
	} else {

		vC51I7RbKXXN, gOCcQzbcC :=  /*line :1*/ dT5PdkB5a(hY26vae8L)
		if gOCcQzbcC != nil {
			return "", gOCcQzbcC
		}
		if  /*line :1*/ ugu28bu(rd6leevODT[0]) {
			return  /*line :1*/ LmpGp3wH_T(vC51I7RbKXXN[:2] + rd6leevODT)
		} else {
			return  /*line :1*/ LmpGp3wH_T(vC51I7RbKXXN + "\\" + rd6leevODT)
		}
	}
}

type Ze0uPJo struct {
	L4Vaw9K	string
	PZzy6B	[]string
	INpA17B	[]uintptr
	PaHqSua0vh1	*YxcIO6yPaPK
}

type YxcIO6yPaPK struct {
	RNYjk89EDcs	bool
	N6YvEOJ	string	
	EbfdRqVxCJb	uint32
	Z7x_Yh	Ad4bWd	
	HagrJqXcyCr	*DDwuM6RpYja	
	Ws3_yq	*DDwuM6RpYja	
	R68JY6IJ8z	bool	
	BNZUWA4SX4F	[]SRlvVjrYa	
	DQS6mLNCA	SRlvVjrYa	
}

var aLeAvjMm Ze0uPJo
var gnROtq_qx YxcIO6yPaPK

func Z0IL9N5Otmal(qqmkxC0Cb string, eLPqUuZoy []string, c4UzX7Pz *Ze0uPJo) (aJF0So int, fYz2KOsTDon uintptr, gOCcQzbcC error) {
	if  /*line :1*/ len(qqmkxC0Cb) == 0 {
		return 0, 0, EWINDOWS
	}
	if c4UzX7Pz == nil {
		c4UzX7Pz = &aLeAvjMm
	}
	mOZVRlrXu := c4UzX7Pz.PaHqSua0vh1
	if mOZVRlrXu == nil {
		mOZVRlrXu = &gnROtq_qx
	}

	if  /*line :1*/ len(c4UzX7Pz.INpA17B) > 3 {
		return 0, 0, EWINDOWS
	}
	if  /*line :1*/ len(c4UzX7Pz.INpA17B) < 3 {
		return 0, 0, EINVAL
	}

	if  /*line :1*/ len(c4UzX7Pz.L4Vaw9K) != 0 {
		
		
		
		
		
		
		var gOCcQzbcC error
		qqmkxC0Cb, gOCcQzbcC =  /*line :1*/ iCvgqkkKz(c4UzX7Pz.L4Vaw9K, qqmkxC0Cb)
		if gOCcQzbcC != nil {
			return 0, 0, gOCcQzbcC
		}
	}
	btqqWx5, gOCcQzbcC :=  /*line :1*/ GcOmFfsibES(qqmkxC0Cb)
	if gOCcQzbcC != nil {
		return 0, 0, gOCcQzbcC
	}

	var iER44O_ string

	if mOZVRlrXu.N6YvEOJ != "" {
		iER44O_ = mOZVRlrXu.N6YvEOJ
	} else {
		iER44O_ =  /*line :1*/ a_cROM(eLPqUuZoy)
	}

	var vswF8Db *uint16
	if  /*line :1*/ len(iER44O_) != 0 {
		vswF8Db, gOCcQzbcC =  /*line :1*/ GcOmFfsibES(iER44O_)
		if gOCcQzbcC != nil {
			return 0, 0, gOCcQzbcC
		}
	}

	var sxJlex_VI *uint16
	if  /*line :1*/ len(c4UzX7Pz.L4Vaw9K) != 0 {
		sxJlex_VI, gOCcQzbcC =  /*line :1*/ GcOmFfsibES(c4UzX7Pz.L4Vaw9K)
		if gOCcQzbcC != nil {
			return 0, 0, gOCcQzbcC
		}
	}

	var mO1D4A9, jU36NLC88, n8SGC6 uint32
	 /*line :1*/ kjlGhacz(&mO1D4A9, &jU36NLC88, &n8SGC6)
	jUTBNPQw := mO1D4A9 < 6 || (mO1D4A9 == 6 && jU36NLC88 <= 1)

	e9_NCKr := func(fYz2KOsTDon SRlvVjrYa) bool { return jUTBNPQw && fYz2KOsTDon&0x10000003 == 3 }

	rd6leevODT, _ :=  /*line :1*/ ZoL0wDe()
	qtaBDTNu := rd6leevODT
	if mOZVRlrXu.DQS6mLNCA != 0 {
		qtaBDTNu = mOZVRlrXu.DQS6mLNCA
	}
	dBn476FVCy2C :=  /*line :1*/ make([]SRlvVjrYa,  /*line :1*/ len(c4UzX7Pz.INpA17B))
	for hA99q3EOK := range c4UzX7Pz.INpA17B {
		if c4UzX7Pz.INpA17B[hA99q3EOK] > 0 {
			ddR6M7 := qtaBDTNu

			if qtaBDTNu != rd6leevODT &&  /*line :1*/ e9_NCKr( /*line :1*/ SRlvVjrYa(c4UzX7Pz.INpA17B[hA99q3EOK])) {
				ddR6M7 = rd6leevODT
			}

			gOCcQzbcC :=  /*line :1*/ VXs68_ZF(rd6leevODT,  /*line :1*/ SRlvVjrYa(c4UzX7Pz.INpA17B[hA99q3EOK]), ddR6M7, &dBn476FVCy2C[hA99q3EOK], 0, true, DUPLICATE_SAME_ACCESS)
			if gOCcQzbcC != nil {
				return 0, 0, gOCcQzbcC
			}
			defer  /*line :1*/ VXs68_ZF(qtaBDTNu, dBn476FVCy2C[hA99q3EOK], 0, nil, 0, false, DUPLICATE_CLOSE_SOURCE)
		}
	}
	sO2vVI :=  /*line :1*/ new(taSzfHcIq)
	sO2vVI.IIfiFP3pufqz, gOCcQzbcC =  /*line :1*/ itvjHDoHQBe(2)
	if gOCcQzbcC != nil {
		return 0, 0, gOCcQzbcC
	}
	defer  /*line :1*/ jbwtrz4(sO2vVI.IIfiFP3pufqz)
	sO2vVI.JxcR_9A2Qcwv =  /*line :1*/ uint32( /*line :1*/ unsafe.Sizeof(*sO2vVI))
	sO2vVI.Am0GjIfa = STARTF_USESTDHANDLES
	if mOZVRlrXu.RNYjk89EDcs {
		sO2vVI.Am0GjIfa |= STARTF_USESHOWWINDOW
		sO2vVI.PnxZZH = SW_HIDE
	}
	if mOZVRlrXu.DQS6mLNCA != 0 {
		gOCcQzbcC =  /*line :1*/ rq14e8x(sO2vVI.IIfiFP3pufqz, 0, _PROC_THREAD_ATTRIBUTE_PARENT_PROCESS,  /*line :1*/ unsafe.Pointer(&mOZVRlrXu.DQS6mLNCA),  /*line :1*/ unsafe.Sizeof(mOZVRlrXu.DQS6mLNCA), nil, nil)
		if gOCcQzbcC != nil {
			return 0, 0, gOCcQzbcC
		}
	}
	sO2vVI.Q7xYfD5083 = dBn476FVCy2C[0]
	sO2vVI.BQKc_j = dBn476FVCy2C[1]
	sO2vVI.Bh8xwzLH = dBn476FVCy2C[2]

	dBn476FVCy2C =  /*line :1*/ append(dBn476FVCy2C, mOZVRlrXu.BNZUWA4SX4F...)

	for hA99q3EOK := range dBn476FVCy2C {
		if  /*line :1*/ e9_NCKr(dBn476FVCy2C[hA99q3EOK]) {
			dBn476FVCy2C[hA99q3EOK] = 0
		}
	}

	oUmRMCO := 0
	for hA99q3EOK := range dBn476FVCy2C {
		if dBn476FVCy2C[hA99q3EOK] != 0 {
			dBn476FVCy2C[oUmRMCO] = dBn476FVCy2C[hA99q3EOK]
			oUmRMCO++
		}
	}
	dBn476FVCy2C = dBn476FVCy2C[:oUmRMCO]

	d8Er6UE :=  /*line :1*/ len(dBn476FVCy2C) > 0 && !mOZVRlrXu.R68JY6IJ8z

	if d8Er6UE {
		gOCcQzbcC =  /*line :1*/ rq14e8x(sO2vVI.IIfiFP3pufqz, 0, _PROC_THREAD_ATTRIBUTE_HANDLE_LIST,  /*line :1*/ unsafe.Pointer(&dBn476FVCy2C[0]),  /*line :1*/ uintptr( /*line :1*/ len(dBn476FVCy2C))* /*line :1*/ unsafe.Sizeof(dBn476FVCy2C[0]), nil, nil)
		if gOCcQzbcC != nil {
			return 0, 0, gOCcQzbcC
		}
	}

	d15ONT, gOCcQzbcC :=  /*line :1*/ w4osefwcSL1K(c4UzX7Pz.PZzy6B)
	if gOCcQzbcC != nil {
		return 0, 0, gOCcQzbcC
	}

	veQ5il04_P :=  /*line :1*/ new(NBPVXpO3Kv)
	uADzMh0kZ7hL := mOZVRlrXu.EbfdRqVxCJb | CREATE_UNICODE_ENVIRONMENT | _EXTENDED_STARTUPINFO_PRESENT
	if mOZVRlrXu.Z7x_Yh != 0 {
		gOCcQzbcC =  /*line :1*/ MizSUaBO5Kk(mOZVRlrXu.Z7x_Yh, btqqWx5, vswF8Db, mOZVRlrXu.HagrJqXcyCr, mOZVRlrXu.Ws3_yq, d8Er6UE, uADzMh0kZ7hL, d15ONT, sxJlex_VI, &sO2vVI.J7Nv3S0G31Q, veQ5il04_P)
	} else {
		gOCcQzbcC =  /*line :1*/ ATXyyvx(btqqWx5, vswF8Db, mOZVRlrXu.HagrJqXcyCr, mOZVRlrXu.Ws3_yq, d8Er6UE, uADzMh0kZ7hL, d15ONT, sxJlex_VI, &sO2vVI.J7Nv3S0G31Q, veQ5il04_P)
	}
	if gOCcQzbcC != nil {
		return 0, 0, gOCcQzbcC
	}
	defer  /*line :1*/ FhKJLgXjwG( /*line :1*/ SRlvVjrYa(veQ5il04_P.BRTb_8))
	 /*line :1*/ runtime.KeepAlive(dBn476FVCy2C)
	 /*line :1*/ runtime.KeepAlive(mOZVRlrXu)

	return  /*line :1*/ int(veQ5il04_P.LjilbGmQX),  /*line :1*/ uintptr(veQ5il04_P.V0VCmzpZzO_), nil
}

func Hj9phn(qqmkxC0Cb string, eLPqUuZoy []string, n924JoWO []string) (gOCcQzbcC error) {
	return EWINDOWS
}
