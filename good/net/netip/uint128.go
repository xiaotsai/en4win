//line :1
package iPdCam_KQOXj

import bits "math/bits"





type rc_vdlQ struct {
	od2_Mw	uint64
	gRGNmk	uint64
}



func vtgBPgPqR37(xUVyHJVSZ4t7 int) rc_vdlQ {
	return rc_vdlQ{^(^ /*line :1*/ uint64(0) >> xUVyHJVSZ4t7), ^ /*line :1*/ uint64(0) << (128 - xUVyHJVSZ4t7)}
}






func (gNtWw_6scp rc_vdlQ) rb6oFzFeRc() bool	{ return gNtWw_6scp.od2_Mw|gNtWw_6scp.gRGNmk == 0 }


func (gNtWw_6scp rc_vdlQ) huLWxs63VG1(imP6SgWD rc_vdlQ) rc_vdlQ {
	return rc_vdlQ{gNtWw_6scp.od2_Mw & imP6SgWD.od2_Mw, gNtWw_6scp.gRGNmk & imP6SgWD.gRGNmk}
}


func (gNtWw_6scp rc_vdlQ) sf7IVRl(imP6SgWD rc_vdlQ) rc_vdlQ {
	return rc_vdlQ{gNtWw_6scp.od2_Mw ^ imP6SgWD.od2_Mw, gNtWw_6scp.gRGNmk ^ imP6SgWD.gRGNmk}
}


func (gNtWw_6scp rc_vdlQ) pmoznX3zzeA(imP6SgWD rc_vdlQ) rc_vdlQ {
	return rc_vdlQ{gNtWw_6scp.od2_Mw | imP6SgWD.od2_Mw, gNtWw_6scp.gRGNmk | imP6SgWD.gRGNmk}
}


func (gNtWw_6scp rc_vdlQ) iJH9GZ53hca() rc_vdlQ {
	return rc_vdlQ{^gNtWw_6scp.od2_Mw, ^gNtWw_6scp.gRGNmk}
}


func (gNtWw_6scp rc_vdlQ) cgATgys_b7Y() rc_vdlQ {
	mXgJUAbOOIv1, oSF4JQ1AR :=  /*line :1*/ bits.Sub64(gNtWw_6scp.gRGNmk, 1, 0)
	return rc_vdlQ{gNtWw_6scp.od2_Mw - oSF4JQ1AR, mXgJUAbOOIv1}
}


func (gNtWw_6scp rc_vdlQ) ocX5kXg() rc_vdlQ {
	mXgJUAbOOIv1, fDaUXZOhe :=  /*line :1*/ bits.Add64(gNtWw_6scp.gRGNmk, 1, 0)
	return rc_vdlQ{gNtWw_6scp.od2_Mw + fDaUXZOhe, mXgJUAbOOIv1}
}





func (gNtWw_6scp *rc_vdlQ) pxcoT3wb() [2]*uint64 {
	return [2]*uint64{&gNtWw_6scp.od2_Mw, &gNtWw_6scp.gRGNmk}
}



func (gNtWw_6scp rc_vdlQ) pBlkWMs9Dj(dKRmmLWD87 uint8) rc_vdlQ {
	return  /*line :1*/ gNtWw_6scp.pmoznX3zzeA( /*line :1*/ vtgBPgPqR37( /*line :1*/ int(dKRmmLWD87)).iJH9GZ53hca())
}



func (gNtWw_6scp rc_vdlQ) v45GzPD_(dKRmmLWD87 uint8) rc_vdlQ {
	return  /*line :1*/ gNtWw_6scp.huLWxs63VG1( /*line :1*/ vtgBPgPqR37( /*line :1*/ int(dKRmmLWD87)))
}
