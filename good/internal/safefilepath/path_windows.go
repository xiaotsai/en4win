//line :1
package FMGFnwa9

import (
	syscall "bUKeamF"
	utf8 "CuAc0pPZwf"
)

func jYOlEOMMvca(a1ebAxLjQEVR string) (string, error) {
	if ! /*line :1*/ utf8.YiR9Xl_UC5x(a1ebAxLjQEVR) {
		return "", c8MKV7RnJO
	}
	for  /*line :1*/ len(a1ebAxLjQEVR) > 1 && a1ebAxLjQEVR[0] == '/' && a1ebAxLjQEVR[1] == '/' {
		a1ebAxLjQEVR = a1ebAxLjQEVR[1:]
	}
	jZDz3F := false
	for w2cGI66cfqc := a1ebAxLjQEVR; w2cGI66cfqc != ""; {

		bwG5afunZgL := 0
		vJ6GBqQP := -1
		for bwG5afunZgL <  /*line :1*/ len(w2cGI66cfqc) && w2cGI66cfqc[bwG5afunZgL] != '/' {
			switch w2cGI66cfqc[bwG5afunZgL] {
			case 0, '\\', ':':
				return "", c8MKV7RnJO
			case '.':
				if vJ6GBqQP < 0 {
					vJ6GBqQP = bwG5afunZgL
				}
			}
			bwG5afunZgL++
		}
		gGg_FSCjj := w2cGI66cfqc[:bwG5afunZgL]
		if bwG5afunZgL <  /*line :1*/ len(w2cGI66cfqc) {
			jZDz3F = true
			w2cGI66cfqc = w2cGI66cfqc[bwG5afunZgL+1:]
		} else {
			w2cGI66cfqc = ""
		}

		hN_5W3J := gGg_FSCjj
		if vJ6GBqQP >= 0 {
			hN_5W3J = gGg_FSCjj[:vJ6GBqQP]
		}
		if  /*line :1*/ vM1dK5DjnZj(hN_5W3J) {
			if vJ6GBqQP < 0 {
				return "", c8MKV7RnJO
			}

			if w2cGI66cfqc, _ :=  /*line :1*/ syscall.LmpGp3wH_T(gGg_FSCjj);  /*line :1*/ len(w2cGI66cfqc) >= 4 && w2cGI66cfqc[:4] == `\\.\` {
				return "", c8MKV7RnJO
			}
		}
	}
	if jZDz3F {

		jmXo5m := [] /*line :1*/ byte(a1ebAxLjQEVR)
		for bwG5afunZgL, f4Ry_5 := range jmXo5m {
			if f4Ry_5 == '/' {
				jmXo5m[bwG5afunZgL] = '\\'
			}
		}
		a1ebAxLjQEVR =  /*line :1*/ string(jmXo5m)
	}
	return a1ebAxLjQEVR, nil
}






func vM1dK5DjnZj(fE3dawra7 string) bool {
	if 3 <=  /*line :1*/ len(fE3dawra7) &&  /*line :1*/ len(fE3dawra7) <= 4 {
		switch  /*line :1*/ string([]byte{ /*line :1*/ nfNzBZigA(fE3dawra7[0]),  /*line :1*/ nfNzBZigA(fE3dawra7[1]),  /*line :1*/ nfNzBZigA(fE3dawra7[2])}) {
		case "CON", "PRN", "AUX", "NUL":
			return  /*line :1*/ len(fE3dawra7) == 3
		case "COM", "LPT":
			return  /*line :1*/ len(fE3dawra7) == 4 && '1' <= fE3dawra7[3] && fE3dawra7[3] <= '9'
		}
	}
	return false
}

func nfNzBZigA(q47Donf byte) byte {
	if 'a' <= q47Donf && q47Donf <= 'z' {
		return q47Donf - ('a' - 'A')
	}
	return q47Donf
}
