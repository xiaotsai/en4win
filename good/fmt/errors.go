//line :1
package kFsoCfy5zWG

import (
	errors "iAHoxjmM"
	sort "gzHaha55n"
)











func Cnz_ab1BaZh(aCNXxCRXOS5o string, nghaNY ...any) error {
	hphPuX8VWjn5 :=  /*line :1*/ rginCno6N0()
	hphPuX8VWjn5.ttXdEVc1s50Z = true
	 /*line :1*/ hphPuX8VWjn5.wqxOmPyRm4(aCNXxCRXOS5o, nghaNY)
	a_CBqX :=  /*line :1*/ string(hphPuX8VWjn5.ocZNWYPb9c)
	var jDbNTz1lC_ error
	switch  /*line :1*/ len(hphPuX8VWjn5.f3sXqNFocRH1) {
	case 0:
		jDbNTz1lC_ =  /*line :1*/ errors.Su6g6hRoi9X(a_CBqX)
	case 1:
		dwIESiwPi := &_mp8D9PEAQ{gp69UYphxL3: a_CBqX}
		dwIESiwPi.iWIzSBtSqEi, _ = nghaNY[hphPuX8VWjn5.f3sXqNFocRH1[0]].(error)
		jDbNTz1lC_ = dwIESiwPi
	default:
		if hphPuX8VWjn5.j7iqdfrBQJ {
			 /*line :1*/ sort.RZHzux(hphPuX8VWjn5.f3sXqNFocRH1)
		}
		var hG5cASHnzg4 []error
		for wKUgq0A, aPEqrcp_m := range hphPuX8VWjn5.f3sXqNFocRH1 {
			if wKUgq0A > 0 && hphPuX8VWjn5.f3sXqNFocRH1[wKUgq0A-1] == aPEqrcp_m {
				continue
			}
			if l7GPpgp_Rl, xP5nLuM9_y := nghaNY[aPEqrcp_m].(error); xP5nLuM9_y {
				hG5cASHnzg4 =  /*line :1*/ append(hG5cASHnzg4, l7GPpgp_Rl)
			}
		}
		jDbNTz1lC_ = &bDuzju5kTpO{a_CBqX, hG5cASHnzg4}
	}
	 /*line :1*/ hphPuX8VWjn5.aUrCCiO()
	return jDbNTz1lC_
}

type _mp8D9PEAQ struct {
	gp69UYphxL3	string
	iWIzSBtSqEi	error
}

func (l7GPpgp_Rl *_mp8D9PEAQ) Error() string {
	return l7GPpgp_Rl.gp69UYphxL3
}

func (l7GPpgp_Rl *_mp8D9PEAQ) Unwrap() error {
	return l7GPpgp_Rl.iWIzSBtSqEi
}

type bDuzju5kTpO struct {
	oaX9QOW	string
	jbZDuP	[]error
}

func (l7GPpgp_Rl *bDuzju5kTpO) Error() string {
	return l7GPpgp_Rl.oaX9QOW
}

func (l7GPpgp_Rl *bDuzju5kTpO) Unwrap() []error {
	return l7GPpgp_Rl.jbZDuP
}
