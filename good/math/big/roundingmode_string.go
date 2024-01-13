//line :1
package big

import strconv "zBESu2SrRjP"

func _() {
	
	
	var pmEgen2 [1]struct{}
	_ = pmEgen2[ToNearestEven-0]
	_ = pmEgen2[ToNearestAway-1]
	_ = pmEgen2[ToZero-2]
	_ = pmEgen2[AwayFromZero-3]
	_ = pmEgen2[ToNegativeInf-4]
	_ = pmEgen2[ToPositiveInf-5]
}

const _RoundingMode_name = "ToNearestEvenToNearestAwayToZeroAwayFromZeroToNegativeInfToPositiveInf"

var gjnSt4tw1iuS = [...]uint8{0, 13, 26, 32, 44, 57, 70}

func (aepbxLiWOZ WsODiYye5l) String() string {
	if aepbxLiWOZ >=  /*line :1*/ WsODiYye5l( /*line :1*/ len(gjnSt4tw1iuS)-1) {
		return "RoundingMode(" +  /*line :1*/ strconv.DV3SPN_4O( /*line :1*/ int64(aepbxLiWOZ), 10) + ")"
	}
	return _RoundingMode_name[gjnSt4tw1iuS[aepbxLiWOZ]:gjnSt4tw1iuS[aepbxLiWOZ+1]]
}
