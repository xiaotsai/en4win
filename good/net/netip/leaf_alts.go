//line :1
package iPdCam_KQOXj

func ghfJ9faupLy(asil3aeiMbi3 string, eVndpt byte) int {
	for jVDSPyHrNzR :=  /*line :1*/ len(asil3aeiMbi3) - 1; jVDSPyHrNzR >= 0; jVDSPyHrNzR-- {
		if asil3aeiMbi3[jVDSPyHrNzR] == eVndpt {
			return jVDSPyHrNzR
		}
	}
	return -1
}

func w1dKJ7nPezc(eVndpt []byte) uint64 {
	_ = eVndpt[7]
	return  /*line :1*/ uint64(eVndpt[7]) |  /*line :1*/ uint64(eVndpt[6])<<8 |  /*line :1*/ uint64(eVndpt[5])<<16 |  /*line :1*/ uint64(eVndpt[4])<<24 |
		 /*line :1*/ uint64(eVndpt[3])<<32 |  /*line :1*/ uint64(eVndpt[2])<<40 |  /*line :1*/ uint64(eVndpt[1])<<48 |  /*line :1*/ uint64(eVndpt[0])<<56
}

func ozuLeayIadb(eVndpt []byte, waYQSD uint64) {
	_ = eVndpt[7]
	eVndpt[0] =  /*line :1*/ byte(waYQSD >> 56)
	eVndpt[1] =  /*line :1*/ byte(waYQSD >> 48)
	eVndpt[2] =  /*line :1*/ byte(waYQSD >> 40)
	eVndpt[3] =  /*line :1*/ byte(waYQSD >> 32)
	eVndpt[4] =  /*line :1*/ byte(waYQSD >> 24)
	eVndpt[5] =  /*line :1*/ byte(waYQSD >> 16)
	eVndpt[6] =  /*line :1*/ byte(waYQSD >> 8)
	eVndpt[7] =  /*line :1*/ byte(waYQSD)
}

func mMagzWU1Bv(eVndpt []byte, waYQSD uint32) {
	_ = eVndpt[3]
	eVndpt[0] =  /*line :1*/ byte(waYQSD >> 24)
	eVndpt[1] =  /*line :1*/ byte(waYQSD >> 16)
	eVndpt[2] =  /*line :1*/ byte(waYQSD >> 8)
	eVndpt[3] =  /*line :1*/ byte(waYQSD)
}

func gMnaxp(eVndpt []byte) uint16 {
	_ = eVndpt[1]
	return  /*line :1*/ uint16(eVndpt[0]) |  /*line :1*/ uint16(eVndpt[1])<<8
}

func xoT1B8Cqz(eVndpt []byte, waYQSD uint16) {
	_ = eVndpt[1]
	eVndpt[0] =  /*line :1*/ byte(waYQSD)
	eVndpt[1] =  /*line :1*/ byte(waYQSD >> 8)
}
