//line :1








package xI9ai2djaJ2

import (
	errors "iAHoxjmM"
	sync "sync"
)


const (
	SeekStart	= 0		
	SeekCurrent	= 1		
	SeekEnd	= 2		
)



var ArPWDHfv =  /*line :1*/ errors.Su6g6hRoi9X("short write")


var eGmfas =  /*line :1*/ errors.Su6g6hRoi9X("invalid write result")


var Y5M_RJS5mo =  /*line :1*/ errors.Su6g6hRoi9X("short buffer")








var K5Sqco =  /*line :1*/ errors.Su6g6hRoi9X("EOF")



var UASgojMNPA =  /*line :1*/ errors.Su6g6hRoi9X("unexpected EOF")




var Bk9XalVfGU8p =  /*line :1*/ errors.Su6g6hRoi9X("multiple Read calls return no data or error")
































type YJ04Fau interface {
	Read(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error)
}










type LXQrGW6BQt interface {
	Write(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error)
}





type OGnrRZaCeywO interface {
	Close() error
}
















type IYpgu1w3joNv interface {
	Seek(xAt6VSGe int64, sG3qryY int) (int64, error)
}


type BumkXrxm interface {
	YJ04Fau
	LXQrGW6BQt
}


type AVCtdys interface {
	YJ04Fau
	OGnrRZaCeywO
}


type HbIs6j interface {
	LXQrGW6BQt
	OGnrRZaCeywO
}


type Cb_DBp interface {
	YJ04Fau
	LXQrGW6BQt
	OGnrRZaCeywO
}


type FCQTCy interface {
	YJ04Fau
	IYpgu1w3joNv
}



type BeTxWmOGOgQ interface {
	YJ04Fau
	IYpgu1w3joNv
	OGnrRZaCeywO
}


type DagtrUcvRY interface {
	LXQrGW6BQt
	IYpgu1w3joNv
}


type CIUsOqlNuX interface {
	YJ04Fau
	LXQrGW6BQt
	IYpgu1w3joNv
}








type CyhQLnaF interface {
	ReadFrom(wZW51lrVDWS YJ04Fau) (f3MWP_v int64, fjG4tVMGB error)
}








type JCVgLD8ld interface {
	WriteTo(iO5t8eLj5Pd LXQrGW6BQt) (f3MWP_v int64, fjG4tVMGB error)
}



























type BhZMfvQ interface {
	ReadAt(x6wLQwPqT1n []byte, hFt2cf0 int64) (f3MWP_v int, fjG4tVMGB error)
}
















type UXuaDREZ interface {
	WriteAt(x6wLQwPqT1n []byte, hFt2cf0 int64) (f3MWP_v int, fjG4tVMGB error)
}










type Nbptmbb interface {
	ReadByte() (byte, error)
}









type HSeaoj16Wq0 interface {
	Nbptmbb
	UnreadByte() error
}


type HaJD1YcaO interface {
	WriteByte(j3YihQz byte) error
}






type G2t9qC5VO9s interface {
	ReadRune() (wZW51lrVDWS rune, p1b5AatMvJ int, fjG4tVMGB error)
}









type MfPRgoXllY interface {
	G2t9qC5VO9s
	UnreadRune() error
}


type HpBo0GpMvi6 interface {
	WriteString(toLe4u string) (f3MWP_v int, fjG4tVMGB error)
}




func NcJkqIuWUrM(iO5t8eLj5Pd LXQrGW6BQt, toLe4u string) (f3MWP_v int, fjG4tVMGB error) {
	if mWtRnqAwj08, olKmXo5 := iO5t8eLj5Pd.(HpBo0GpMvi6); olKmXo5 {
		return  /*line :1*/ mWtRnqAwj08.WriteString(toLe4u)
	}
	return  /*line :1*/ iO5t8eLj5Pd.Write([] /*line :1*/ byte(toLe4u))
}









func DGsbnMoa2(wZW51lrVDWS YJ04Fau, aHfq7_Z []byte, oVKbGFLPH int) (f3MWP_v int, fjG4tVMGB error) {
	if  /*line :1*/ len(aHfq7_Z) < oVKbGFLPH {
		return 0, Y5M_RJS5mo
	}
	for f3MWP_v < oVKbGFLPH && fjG4tVMGB == nil {
		var oxWIvJYvyQ int
		oxWIvJYvyQ, fjG4tVMGB =  /*line :1*/ wZW51lrVDWS.Read(aHfq7_Z[f3MWP_v:])
		f3MWP_v += oxWIvJYvyQ
	}
	if f3MWP_v >= oVKbGFLPH {
		fjG4tVMGB = nil
	} else if f3MWP_v > 0 && fjG4tVMGB == K5Sqco {
		fjG4tVMGB = UASgojMNPA
	}
	return
}








func G0zBgKg(wZW51lrVDWS YJ04Fau, aHfq7_Z []byte) (f3MWP_v int, fjG4tVMGB error) {
	return  /*line :1*/ DGsbnMoa2(wZW51lrVDWS, aHfq7_Z,  /*line :1*/ len(aHfq7_Z))
}








func ZmP3CM82Z(k4HVyS LXQrGW6BQt, nXG6m66 YJ04Fau, f3MWP_v int64) (f0dqPj int64, fjG4tVMGB error) {
	f0dqPj, fjG4tVMGB =  /*line :1*/ FxikaFo5o7OE(k4HVyS,  /*line :1*/ ZyvU21xvd6(nXG6m66, f3MWP_v))
	if f0dqPj == f3MWP_v {
		return f3MWP_v, nil
	}
	if f0dqPj < f3MWP_v && fjG4tVMGB == nil {

		fjG4tVMGB = K5Sqco
	}
	return
}













func FxikaFo5o7OE(k4HVyS LXQrGW6BQt, nXG6m66 YJ04Fau) (f0dqPj int64, fjG4tVMGB error) {
	return  /*line :1*/ a0xVlO6(k4HVyS, nXG6m66, nil)
}








func JkoS9BI7Jw(k4HVyS LXQrGW6BQt, nXG6m66 YJ04Fau, aHfq7_Z []byte) (f0dqPj int64, fjG4tVMGB error) {
	if aHfq7_Z != nil &&  /*line :1*/ len(aHfq7_Z) == 0 {
		 /*line :1*/ panic("empty buffer in CopyBuffer")
	}
	return  /*line :1*/ a0xVlO6(k4HVyS, nXG6m66, aHfq7_Z)
}



func a0xVlO6(k4HVyS LXQrGW6BQt, nXG6m66 YJ04Fau, aHfq7_Z []byte) (f0dqPj int64, fjG4tVMGB error) {

	if pTFv_YNzwE, olKmXo5 := nXG6m66.(JCVgLD8ld); olKmXo5 {
		return  /*line :1*/ pTFv_YNzwE.WriteTo(k4HVyS)
	}

	if rYKQT0N, olKmXo5 := k4HVyS.(CyhQLnaF); olKmXo5 {
		return  /*line :1*/ rYKQT0N.ReadFrom(nXG6m66)
	}
	if aHfq7_Z == nil {
		p1b5AatMvJ := 32 * 1024
		if fELne_U5YR, olKmXo5 := nXG6m66.(*E1nvXo6); olKmXo5 &&  /*line :1*/ int64(p1b5AatMvJ) > fELne_U5YR.MLQU8HCL {
			if fELne_U5YR.MLQU8HCL < 1 {
				p1b5AatMvJ = 1
			} else {
				p1b5AatMvJ =  /*line :1*/ int(fELne_U5YR.MLQU8HCL)
			}
		}
		aHfq7_Z =  /*line :1*/ make([]byte, p1b5AatMvJ)
	}
	for {
		fBxr6gYSESyl, wq6O3YWt1u :=  /*line :1*/ nXG6m66.Read(aHfq7_Z)
		if fBxr6gYSESyl > 0 {
			vpXSicX9M, a6zV_m :=  /*line :1*/ k4HVyS.Write(aHfq7_Z[0:fBxr6gYSESyl])
			if vpXSicX9M < 0 || fBxr6gYSESyl < vpXSicX9M {
				vpXSicX9M = 0
				if a6zV_m == nil {
					a6zV_m = eGmfas
				}
			}
			f0dqPj +=  /*line :1*/ int64(vpXSicX9M)
			if a6zV_m != nil {
				fjG4tVMGB = a6zV_m
				break
			}
			if fBxr6gYSESyl != vpXSicX9M {
				fjG4tVMGB = ArPWDHfv
				break
			}
		}
		if wq6O3YWt1u != nil {
			if wq6O3YWt1u != K5Sqco {
				fjG4tVMGB = wq6O3YWt1u
			}
			break
		}
	}
	return f0dqPj, fjG4tVMGB
}




func ZyvU21xvd6(wZW51lrVDWS YJ04Fau, f3MWP_v int64) YJ04Fau	{ return &E1nvXo6{wZW51lrVDWS, f3MWP_v} }





type E1nvXo6 struct {
	Lz9mJpf_yfQ	YJ04Fau	
	MLQU8HCL	int64	
}

func (fELne_U5YR *E1nvXo6) Read(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error) {
	if fELne_U5YR.MLQU8HCL <= 0 {
		return 0, K5Sqco
	}
	if  /*line :1*/ int64( /*line :1*/ len(x6wLQwPqT1n)) > fELne_U5YR.MLQU8HCL {
		x6wLQwPqT1n = x6wLQwPqT1n[0:fELne_U5YR.MLQU8HCL]
	}
	f3MWP_v, fjG4tVMGB =  /*line :1*/ fELne_U5YR.Lz9mJpf_yfQ.Read(x6wLQwPqT1n)
	fELne_U5YR.MLQU8HCL -=  /*line :1*/ int64(f3MWP_v)
	return
}



func AcJOCX6ye(wZW51lrVDWS BhZMfvQ, hFt2cf0 int64, f3MWP_v int64) *Y7BZpUb4mW {
	var p2_aqn34 int64
	const maxint64 = 1<<63 - 1
	if hFt2cf0 <= maxint64-f3MWP_v {
		p2_aqn34 = f3MWP_v + hFt2cf0
	} else {

		p2_aqn34 = maxint64
	}
	return &Y7BZpUb4mW{wZW51lrVDWS, hFt2cf0, hFt2cf0, p2_aqn34}
}



type Y7BZpUb4mW struct {
	q3pcryH	BhZMfvQ
	hnIfbT	int64
	bzmFVlB	int64
	h3yLdIJN	int64
}

func (toLe4u *Y7BZpUb4mW) Read(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error) {
	if toLe4u.bzmFVlB >= toLe4u.h3yLdIJN {
		return 0, K5Sqco
	}
	if mPMPDBI := toLe4u.h3yLdIJN - toLe4u.bzmFVlB;  /*line :1*/ int64( /*line :1*/ len(x6wLQwPqT1n)) > mPMPDBI {
		x6wLQwPqT1n = x6wLQwPqT1n[0:mPMPDBI]
	}
	f3MWP_v, fjG4tVMGB =  /*line :1*/ toLe4u.q3pcryH.ReadAt(x6wLQwPqT1n, toLe4u.bzmFVlB)
	toLe4u.bzmFVlB +=  /*line :1*/ int64(f3MWP_v)
	return
}

var mtS38uV =  /*line :1*/ errors.Su6g6hRoi9X("Seek: invalid whence")
var sIiKCTc =  /*line :1*/ errors.Su6g6hRoi9X("Seek: invalid offset")

func (toLe4u *Y7BZpUb4mW) Seek(xAt6VSGe int64, sG3qryY int) (int64, error) {
	switch sG3qryY {
	default:
		return 0, mtS38uV
	case SeekStart:
		xAt6VSGe += toLe4u.hnIfbT
	case SeekCurrent:
		xAt6VSGe += toLe4u.bzmFVlB
	case SeekEnd:
		xAt6VSGe += toLe4u.h3yLdIJN
	}
	if xAt6VSGe < toLe4u.hnIfbT {
		return 0, sIiKCTc
	}
	toLe4u.bzmFVlB = xAt6VSGe
	return xAt6VSGe - toLe4u.hnIfbT, nil
}

func (toLe4u *Y7BZpUb4mW) ReadAt(x6wLQwPqT1n []byte, hFt2cf0 int64) (f3MWP_v int, fjG4tVMGB error) {
	if hFt2cf0 < 0 || hFt2cf0 >= toLe4u.h3yLdIJN-toLe4u.hnIfbT {
		return 0, K5Sqco
	}
	hFt2cf0 += toLe4u.hnIfbT
	if mPMPDBI := toLe4u.h3yLdIJN - hFt2cf0;  /*line :1*/ int64( /*line :1*/ len(x6wLQwPqT1n)) > mPMPDBI {
		x6wLQwPqT1n = x6wLQwPqT1n[0:mPMPDBI]
		f3MWP_v, fjG4tVMGB =  /*line :1*/ toLe4u.q3pcryH.ReadAt(x6wLQwPqT1n, hFt2cf0)
		if fjG4tVMGB == nil {
			fjG4tVMGB = K5Sqco
		}
		return f3MWP_v, fjG4tVMGB
	}
	return  /*line :1*/ toLe4u.q3pcryH.ReadAt(x6wLQwPqT1n, hFt2cf0)
}


func (toLe4u *Y7BZpUb4mW) Size() int64	{ return toLe4u.h3yLdIJN - toLe4u.hnIfbT }


type DstskywsZ8F struct {
	fTN6KRFAMf	UXuaDREZ
	iLmmbkjD	int64	
	hK8ZWH0lsW	int64	
}



func FZeE_lr0gDR1(iO5t8eLj5Pd UXuaDREZ, hFt2cf0 int64) *DstskywsZ8F {
	return &DstskywsZ8F{iO5t8eLj5Pd, hFt2cf0, hFt2cf0}
}

func (m0aU_BVjrcWT *DstskywsZ8F) Write(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error) {
	f3MWP_v, fjG4tVMGB =  /*line :1*/ m0aU_BVjrcWT.fTN6KRFAMf.WriteAt(x6wLQwPqT1n, m0aU_BVjrcWT.hK8ZWH0lsW)
	m0aU_BVjrcWT.hK8ZWH0lsW +=  /*line :1*/ int64(f3MWP_v)
	return
}

func (m0aU_BVjrcWT *DstskywsZ8F) WriteAt(x6wLQwPqT1n []byte, hFt2cf0 int64) (f3MWP_v int, fjG4tVMGB error) {
	if hFt2cf0 < 0 {
		return 0, sIiKCTc
	}

	hFt2cf0 += m0aU_BVjrcWT.iLmmbkjD
	return  /*line :1*/ m0aU_BVjrcWT.fTN6KRFAMf.WriteAt(x6wLQwPqT1n, hFt2cf0)
}

func (m0aU_BVjrcWT *DstskywsZ8F) Seek(xAt6VSGe int64, sG3qryY int) (int64, error) {
	switch sG3qryY {
	default:
		return 0, mtS38uV
	case SeekStart:
		xAt6VSGe += m0aU_BVjrcWT.iLmmbkjD
	case SeekCurrent:
		xAt6VSGe += m0aU_BVjrcWT.hK8ZWH0lsW
	}
	if xAt6VSGe < m0aU_BVjrcWT.iLmmbkjD {
		return 0, sIiKCTc
	}
	m0aU_BVjrcWT.hK8ZWH0lsW = xAt6VSGe
	return xAt6VSGe - m0aU_BVjrcWT.iLmmbkjD, nil
}






func Qzn0uN4lM2UH(wZW51lrVDWS YJ04Fau, iO5t8eLj5Pd LXQrGW6BQt) YJ04Fau {
	return &aYpwvIKD5{wZW51lrVDWS, iO5t8eLj5Pd}
}

type aYpwvIKD5 struct {
	wpAp69r	YJ04Fau
	dJuzuY0k5	LXQrGW6BQt
}

func (zsrcBG *aYpwvIKD5) Read(x6wLQwPqT1n []byte) (f3MWP_v int, fjG4tVMGB error) {
	f3MWP_v, fjG4tVMGB =  /*line :1*/ zsrcBG.wpAp69r.Read(x6wLQwPqT1n)
	if f3MWP_v > 0 {
		if f3MWP_v, fjG4tVMGB :=  /*line :1*/ zsrcBG.dJuzuY0k5.Write(x6wLQwPqT1n[:f3MWP_v]); fjG4tVMGB != nil {
			return f3MWP_v, fjG4tVMGB
		}
	}
	return
}



var GT3wkcZ LXQrGW6BQt = fDw25aZSa{}

type fDw25aZSa struct{}



var _ CyhQLnaF = fDw25aZSa{}

func (fDw25aZSa) Write(x6wLQwPqT1n []byte) (int, error) {
	return  /*line :1*/ len(x6wLQwPqT1n), nil
}

func (fDw25aZSa) WriteString(toLe4u string) (int, error) {
	return  /*line :1*/ len(toLe4u), nil
}

var sS2_puS = sync.OrP5FGPq{
	IYbTy050ek: func() any {
		rG6_2l2i :=  /*line :1*/ make([]byte, 8192)
		return &rG6_2l2i
	},
}

func (fDw25aZSa) ReadFrom(wZW51lrVDWS YJ04Fau) (f3MWP_v int64, fjG4tVMGB error) {
	xcJCiaPD4Pbw :=  /*line :1*/ sS2_puS.Get().(*[]byte)
	qJF1pQ := 0
	for {
		qJF1pQ, fjG4tVMGB =  /*line :1*/ wZW51lrVDWS.Read(*xcJCiaPD4Pbw)
		f3MWP_v +=  /*line :1*/ int64(qJF1pQ)
		if fjG4tVMGB != nil {
			 /*line :1*/ sS2_puS.Put(xcJCiaPD4Pbw)
			if fjG4tVMGB == K5Sqco {
				return f3MWP_v, nil
			}
			return
		}
	}
}





func FiiD7W(wZW51lrVDWS YJ04Fau) AVCtdys {
	if _, olKmXo5 := wZW51lrVDWS.(JCVgLD8ld); olKmXo5 {
		return iaLiHAN{wZW51lrVDWS}
	}
	return hhrt34Zc6n16{wZW51lrVDWS}
}

type hhrt34Zc6n16 struct {
	YJ04Fau
}

func (hhrt34Zc6n16) Close() error	{ return nil }

type iaLiHAN struct {
	YJ04Fau
}

func (iaLiHAN) Close() error	{ return nil }

func (j3YihQz iaLiHAN) WriteTo(iO5t8eLj5Pd LXQrGW6BQt) (f3MWP_v int64, fjG4tVMGB error) {
	return  /*line :1*/ j3YihQz.YJ04Fau.(JCVgLD8ld).WriteTo(iO5t8eLj5Pd)
}





func ZAZvd8HSB(wZW51lrVDWS YJ04Fau) ([]byte, error) {
	rG6_2l2i :=  /*line :1*/ make([]byte, 0, 512)
	for {
		f3MWP_v, fjG4tVMGB :=  /*line :1*/ wZW51lrVDWS.Read(rG6_2l2i[ /*line :1*/ len(rG6_2l2i): /*line :1*/ cap(rG6_2l2i)])
		rG6_2l2i = rG6_2l2i[: /*line :1*/ len(rG6_2l2i)+f3MWP_v]
		if fjG4tVMGB != nil {
			if fjG4tVMGB == K5Sqco {
				fjG4tVMGB = nil
			}
			return rG6_2l2i, fjG4tVMGB
		}

		if  /*line :1*/ len(rG6_2l2i) ==  /*line :1*/ cap(rG6_2l2i) {

			rG6_2l2i =  /*line :1*/ append(rG6_2l2i, 0)[: /*line :1*/ len(rG6_2l2i)]
		}
	}
}
