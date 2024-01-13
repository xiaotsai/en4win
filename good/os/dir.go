//line :1
package hPMrClpC

import (
	fs "y1BamVjnXsaa"
	sort "gzHaha55n"
)

type uppLHxeYfJc int

const (
	readdirName	uppLHxeYfJc	= iota
	readdirDirEntry
	readdirFileInfo
)


















func (qzeVi328uMg *BvGocYcXx) Readdir(zHDjCTHI int) ([]LWsU3x8MX, error) {
	if qzeVi328uMg == nil {
		return nil, J6ilKFR6o
	}
	_, _, jR_i1_n, ugqb2IW :=  /*line :1*/ qzeVi328uMg.dQTx5qF2N(zHDjCTHI, readdirFileInfo)
	if jR_i1_n == nil {

		jR_i1_n = []LWsU3x8MX{}
	}
	return jR_i1_n, ugqb2IW
}
















func (qzeVi328uMg *BvGocYcXx) Readdirnames(zHDjCTHI int) (c8VnEL []string, ugqb2IW error) {
	if qzeVi328uMg == nil {
		return nil, J6ilKFR6o
	}
	c8VnEL, _, _, ugqb2IW =  /*line :1*/ qzeVi328uMg.dQTx5qF2N(zHDjCTHI, readdirName)
	if c8VnEL == nil {

		c8VnEL = []string{}
	}
	return c8VnEL, ugqb2IW
}



type XXSR6RRqno8V = fs.RbfBMxa











func (qzeVi328uMg *BvGocYcXx) ReadDir(zHDjCTHI int) ([]XXSR6RRqno8V, error) {
	if qzeVi328uMg == nil {
		return nil, J6ilKFR6o
	}
	_, ztdMGkvXuP, _, ugqb2IW :=  /*line :1*/ qzeVi328uMg.dQTx5qF2N(zHDjCTHI, readdirDirEntry)
	if ztdMGkvXuP == nil {

		ztdMGkvXuP = []XXSR6RRqno8V{}
	}
	return ztdMGkvXuP, ugqb2IW
}



var ta70qqF4Ao4 bool






func FbJujn(jGBs03 string) ([]XXSR6RRqno8V, error) {
	qzeVi328uMg, ugqb2IW :=  /*line :1*/ Cvk6wxcjw(jGBs03)
	if ugqb2IW != nil {
		return nil, ugqb2IW
	}
	defer  /*line :1*/ qzeVi328uMg.Close()

	frqQjRv8, ugqb2IW :=  /*line :1*/ qzeVi328uMg.ReadDir(-1)
	 /*line :1*/ sort.Wo2XqNIPTLLO(frqQjRv8, func(lWfaal, lpNM41IF int) bool { return  /*line :1*/ frqQjRv8[lWfaal].Name() <  /*line :1*/ frqQjRv8[lpNM41IF].Name() })
	return frqQjRv8, ugqb2IW
}
