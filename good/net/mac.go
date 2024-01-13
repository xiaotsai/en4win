//line :1
package gF9mZs2

const hexDigit = "0123456789abcdef"


type CGTWv0eqN []byte

func (uI7LZDHm6 CGTWv0eqN) String() string {
	if  /*line :1*/ len(uI7LZDHm6) == 0 {
		return ""
	}
	zynKBqWHCa :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len(uI7LZDHm6)*3-1)
	for eaAUaB7Z, fIadEXIim6l := range uI7LZDHm6 {
		if eaAUaB7Z > 0 {
			zynKBqWHCa =  /*line :1*/ append(zynKBqWHCa, ':')
		}
		zynKBqWHCa =  /*line :1*/ append(zynKBqWHCa, hexDigit[fIadEXIim6l>>4])
		zynKBqWHCa =  /*line :1*/ append(zynKBqWHCa, hexDigit[fIadEXIim6l&0xF])
	}
	return  /*line :1*/ string(zynKBqWHCa)
}













func FCvhWFHuHaJ(dRoFccaZ string) (e9exwCBPH CGTWv0eqN, h_ljk48Bm error) {
	if  /*line :1*/ len(dRoFccaZ) < 14 {
		goto error
	}

	if dRoFccaZ[2] == ':' || dRoFccaZ[2] == '-' {
		if ( /*line :1*/ len(dRoFccaZ)+1)%3 != 0 {
			goto error
		}
		doauF8Y := ( /*line :1*/ len(dRoFccaZ) + 1) / 3
		if doauF8Y != 6 && doauF8Y != 8 && doauF8Y != 20 {
			goto error
		}
		e9exwCBPH =  /*line :1*/ make(CGTWv0eqN, doauF8Y)
		for tWO2olI, eaAUaB7Z := 0, 0; eaAUaB7Z < doauF8Y; eaAUaB7Z++ {
			var d30HIwxht bool
			if e9exwCBPH[eaAUaB7Z], d30HIwxht =  /*line :1*/ kHN6rj(dRoFccaZ[tWO2olI:], dRoFccaZ[2]); !d30HIwxht {
				goto error
			}
			tWO2olI += 3
		}
	} else if dRoFccaZ[4] == '.' {
		if ( /*line :1*/ len(dRoFccaZ)+1)%5 != 0 {
			goto error
		}
		doauF8Y := 2 * ( /*line :1*/ len(dRoFccaZ) + 1) / 5
		if doauF8Y != 6 && doauF8Y != 8 && doauF8Y != 20 {
			goto error
		}
		e9exwCBPH =  /*line :1*/ make(CGTWv0eqN, doauF8Y)
		for tWO2olI, eaAUaB7Z := 0, 0; eaAUaB7Z < doauF8Y; eaAUaB7Z += 2 {
			var d30HIwxht bool
			if e9exwCBPH[eaAUaB7Z], d30HIwxht =  /*line :1*/ kHN6rj(dRoFccaZ[tWO2olI:tWO2olI+2], 0); !d30HIwxht {
				goto error
			}
			if e9exwCBPH[eaAUaB7Z+1], d30HIwxht =  /*line :1*/ kHN6rj(dRoFccaZ[tWO2olI+2:], dRoFccaZ[4]); !d30HIwxht {
				goto error
			}
			tWO2olI += 5
		}
	} else {
		goto error
	}
	return e9exwCBPH, nil

error:
	return nil, &DAWLIQHc{Z68v0y: "invalid MAC address", BCCnAgFxG: dRoFccaZ}
}
