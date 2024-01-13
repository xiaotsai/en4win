//line :1
package MjXXMR

import (
	io "xI9ai2djaJ2"
	syscall "bUKeamF"
)


func AOPFUX2cuQs(s_ZR_8 *ENfmDmMaozH, iEBoOae8j syscall.SRlvVjrYa, mnkyfx int64) (xB7wJxWigaca int64, dtU8tBiUGS4 error) {
	if s_ZR_8.sTN5PtBj1N == kindPipe {

		return 0, syscall.ESPIPE
	}
	if ew26Axd, _ :=  /*line :1*/ syscall.VnXJph2(iEBoOae8j); ew26Axd == syscall.FILE_TYPE_PIPE {
		return 0, syscall.ESPIPE
	}

	if dtU8tBiUGS4 :=  /*line :1*/ s_ZR_8.aAav5o7pLI(); dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}
	defer  /*line :1*/ s_ZR_8.gb0d2X()

	aeapDF := &s_ZR_8.aQ9pa1ac
	aeapDF.azj1aOHqks = iEBoOae8j

	jf9KwJcd64_S, dtU8tBiUGS4 :=  /*line :1*/ syscall.O4FqtrPqbHRY(aeapDF.azj1aOHqks, 0, io.SeekCurrent)
	if dtU8tBiUGS4 != nil {
		return 0, dtU8tBiUGS4
	}

	if mnkyfx <= 0 {

		mnkyfx, dtU8tBiUGS4 =  /*line :1*/ syscall.O4FqtrPqbHRY(aeapDF.azj1aOHqks, -jf9KwJcd64_S, io.SeekEnd)
		if dtU8tBiUGS4 != nil {
			return
		}

		if _, dtU8tBiUGS4 =  /*line :1*/ syscall.O4FqtrPqbHRY(aeapDF.azj1aOHqks, jf9KwJcd64_S, io.SeekStart); dtU8tBiUGS4 != nil {
			return
		}
	}

	
	
	
	const maxChunkSizePerCall =  /*line :1*/ int64(0x7fffffff - 1)

	for mnkyfx > 0 {
		dzo8ADgDpSJC := maxChunkSizePerCall
		if dzo8ADgDpSJC > mnkyfx {
			dzo8ADgDpSJC = mnkyfx
		}

		aeapDF.sSqPaaxx =  /*line :1*/ uint32(dzo8ADgDpSJC)
		aeapDF.gZWJ1u9D4auz.BEJ63D =  /*line :1*/ uint32(jf9KwJcd64_S)
		aeapDF.gZWJ1u9D4auz.BvaSyl =  /*line :1*/ uint32(jf9KwJcd64_S >> 32)

		fuM4gU, dtU8tBiUGS4 :=  /*line :1*/ sX12A2hZ(aeapDF, func(aeapDF *lVE75f25P) error {
			return  /*line :1*/ syscall.PFrj34i0s(aeapDF.ef3m3FV3JP.X8OEsVYSby, aeapDF.azj1aOHqks, aeapDF.sSqPaaxx, 0, &aeapDF.gZWJ1u9D4auz, nil, syscall.TF_WRITE_BEHIND)
		})
		if dtU8tBiUGS4 != nil {
			return xB7wJxWigaca, dtU8tBiUGS4
		}

		jf9KwJcd64_S +=  /*line :1*/ int64(fuM4gU)

		if _, dtU8tBiUGS4 =  /*line :1*/ syscall.O4FqtrPqbHRY(aeapDF.azj1aOHqks, jf9KwJcd64_S, io.SeekStart); dtU8tBiUGS4 != nil {
			return xB7wJxWigaca, dtU8tBiUGS4
		}

		mnkyfx -=  /*line :1*/ int64(fuM4gU)
		xB7wJxWigaca +=  /*line :1*/ int64(fuM4gU)
	}

	return
}
