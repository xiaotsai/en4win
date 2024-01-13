//line :1
package fRAfQd_

import errors "iAHoxjmM"





























































































const (
	Layout	= "01/02 03:04:05PM '06 -0700"		
	ANSIC	= "Mon Jan _2 15:04:05 2006"
	UnixDate	= "Mon Jan _2 15:04:05 MST 2006"
	RubyDate	= "Mon Jan 02 15:04:05 -0700 2006"
	RFC822	= "02 Jan 06 15:04 MST"
	RFC822Z	= "02 Jan 06 15:04 -0700"		
	RFC850	= "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123	= "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z	= "Mon, 02 Jan 2006 15:04:05 -0700"		
	RFC3339	= "2006-01-02T15:04:05Z07:00"
	RFC3339Nano	= "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen	= "3:04PM"
	
	Stamp	= "Jan _2 15:04:05"
	StampMilli	= "Jan _2 15:04:05.000"
	StampMicro	= "Jan _2 15:04:05.000000"
	StampNano	= "Jan _2 15:04:05.000000000"
	DateTime	= "2006-01-02 15:04:05"
	DateOnly	= "2006-01-02"
	TimeOnly	= "15:04:05"
)

const (
	_	= iota
	stdLongMonth	= iota + stdNeedDate		
	stdMonth			
	stdNumMonth			
	stdZeroMonth			
	stdLongWeekDay			
	stdWeekDay			
	stdDay			
	stdUnderDay			
	stdZeroDay			
	stdUnderYearDay			
	stdZeroYearDay			
	stdHour	= iota + stdNeedClock		
	stdHour12			
	stdZeroHour12			
	stdMinute			
	stdZeroMinute			
	stdSecond			
	stdZeroSecond			
	stdLongYear	= iota + stdNeedDate		
	stdYear			
	stdPM	= iota + stdNeedClock		
	stdpm			
	stdTZ	= iota		
	stdISO8601TZ			
	stdISO8601SecondsTZ			
	stdISO8601ShortTZ			
	stdISO8601ColonTZ			
	stdISO8601ColonSecondsTZ			
	stdNumTZ			
	stdNumSecondsTz			
	stdNumShortTZ			
	stdNumColonTZ			
	stdNumColonSecondsTZ			
	stdFracSecond0			
	stdFracSecond9			

	stdNeedDate	= 1 << 8		
	stdNeedClock	= 2 << 8		
	stdArgShift	= 16		
	stdSeparatorShift	= 28		
	stdMask	= 1<<stdArgShift - 1		
)


var a9xld7jwF = [...]int{stdZeroMonth, stdZeroDay, stdZeroHour12, stdZeroMinute, stdZeroSecond, stdYear}



func ahf0C5(eZ59XDLDu4Ov string) bool {
	if  /*line :1*/ len(eZ59XDLDu4Ov) == 0 {
		return false
	}
	fnpIad3KB := eZ59XDLDu4Ov[0]
	return 'a' <= fnpIad3KB && fnpIad3KB <= 'z'
}



func gQkzYn3i(pnn26KE string) (hcap7N string, gA16DoaF_ int, cAS4S6 string) {
	for eUo6AxKM7 := 0; eUo6AxKM7 <  /*line :1*/ len(pnn26KE); eUo6AxKM7++ {
		switch fnpIad3KB :=  /*line :1*/ int(pnn26KE[eUo6AxKM7]); fnpIad3KB {
		case 'J':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+3 && pnn26KE[eUo6AxKM7:eUo6AxKM7+3] == "Jan" {
				if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+7 && pnn26KE[eUo6AxKM7:eUo6AxKM7+7] == "January" {
					return pnn26KE[0:eUo6AxKM7], stdLongMonth, pnn26KE[eUo6AxKM7+7:]
				}
				if ! /*line :1*/ ahf0C5(pnn26KE[eUo6AxKM7+3:]) {
					return pnn26KE[0:eUo6AxKM7], stdMonth, pnn26KE[eUo6AxKM7+3:]
				}
			}

		case 'M':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+3 {
				if pnn26KE[eUo6AxKM7:eUo6AxKM7+3] == "Mon" {
					if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+6 && pnn26KE[eUo6AxKM7:eUo6AxKM7+6] == "Monday" {
						return pnn26KE[0:eUo6AxKM7], stdLongWeekDay, pnn26KE[eUo6AxKM7+6:]
					}
					if ! /*line :1*/ ahf0C5(pnn26KE[eUo6AxKM7+3:]) {
						return pnn26KE[0:eUo6AxKM7], stdWeekDay, pnn26KE[eUo6AxKM7+3:]
					}
				}
				if pnn26KE[eUo6AxKM7:eUo6AxKM7+3] == "MST" {
					return pnn26KE[0:eUo6AxKM7], stdTZ, pnn26KE[eUo6AxKM7+3:]
				}
			}

		case '0':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+2 && '1' <= pnn26KE[eUo6AxKM7+1] && pnn26KE[eUo6AxKM7+1] <= '6' {
				return pnn26KE[0:eUo6AxKM7], a9xld7jwF[pnn26KE[eUo6AxKM7+1]-'1'], pnn26KE[eUo6AxKM7+2:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+3 && pnn26KE[eUo6AxKM7+1] == '0' && pnn26KE[eUo6AxKM7+2] == '2' {
				return pnn26KE[0:eUo6AxKM7], stdZeroYearDay, pnn26KE[eUo6AxKM7+3:]
			}

		case '1':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+2 && pnn26KE[eUo6AxKM7+1] == '5' {
				return pnn26KE[0:eUo6AxKM7], stdHour, pnn26KE[eUo6AxKM7+2:]
			}
			return pnn26KE[0:eUo6AxKM7], stdNumMonth, pnn26KE[eUo6AxKM7+1:]

		case '2':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+4 && pnn26KE[eUo6AxKM7:eUo6AxKM7+4] == "2006" {
				return pnn26KE[0:eUo6AxKM7], stdLongYear, pnn26KE[eUo6AxKM7+4:]
			}
			return pnn26KE[0:eUo6AxKM7], stdDay, pnn26KE[eUo6AxKM7+1:]

		case '_':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+2 && pnn26KE[eUo6AxKM7+1] == '2' {

				if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+5 && pnn26KE[eUo6AxKM7+1:eUo6AxKM7+5] == "2006" {
					return pnn26KE[0 : eUo6AxKM7+1], stdLongYear, pnn26KE[eUo6AxKM7+5:]
				}
				return pnn26KE[0:eUo6AxKM7], stdUnderDay, pnn26KE[eUo6AxKM7+2:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+3 && pnn26KE[eUo6AxKM7+1] == '_' && pnn26KE[eUo6AxKM7+2] == '2' {
				return pnn26KE[0:eUo6AxKM7], stdUnderYearDay, pnn26KE[eUo6AxKM7+3:]
			}

		case '3':
			return pnn26KE[0:eUo6AxKM7], stdHour12, pnn26KE[eUo6AxKM7+1:]

		case '4':
			return pnn26KE[0:eUo6AxKM7], stdMinute, pnn26KE[eUo6AxKM7+1:]

		case '5':
			return pnn26KE[0:eUo6AxKM7], stdSecond, pnn26KE[eUo6AxKM7+1:]

		case 'P':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+2 && pnn26KE[eUo6AxKM7+1] == 'M' {
				return pnn26KE[0:eUo6AxKM7], stdPM, pnn26KE[eUo6AxKM7+2:]
			}

		case 'p':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+2 && pnn26KE[eUo6AxKM7+1] == 'm' {
				return pnn26KE[0:eUo6AxKM7], stdpm, pnn26KE[eUo6AxKM7+2:]
			}

		case '-':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+7 && pnn26KE[eUo6AxKM7:eUo6AxKM7+7] == "-070000" {
				return pnn26KE[0:eUo6AxKM7], stdNumSecondsTz, pnn26KE[eUo6AxKM7+7:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+9 && pnn26KE[eUo6AxKM7:eUo6AxKM7+9] == "-07:00:00" {
				return pnn26KE[0:eUo6AxKM7], stdNumColonSecondsTZ, pnn26KE[eUo6AxKM7+9:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+5 && pnn26KE[eUo6AxKM7:eUo6AxKM7+5] == "-0700" {
				return pnn26KE[0:eUo6AxKM7], stdNumTZ, pnn26KE[eUo6AxKM7+5:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+6 && pnn26KE[eUo6AxKM7:eUo6AxKM7+6] == "-07:00" {
				return pnn26KE[0:eUo6AxKM7], stdNumColonTZ, pnn26KE[eUo6AxKM7+6:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+3 && pnn26KE[eUo6AxKM7:eUo6AxKM7+3] == "-07" {
				return pnn26KE[0:eUo6AxKM7], stdNumShortTZ, pnn26KE[eUo6AxKM7+3:]
			}

		case 'Z':
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+7 && pnn26KE[eUo6AxKM7:eUo6AxKM7+7] == "Z070000" {
				return pnn26KE[0:eUo6AxKM7], stdISO8601SecondsTZ, pnn26KE[eUo6AxKM7+7:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+9 && pnn26KE[eUo6AxKM7:eUo6AxKM7+9] == "Z07:00:00" {
				return pnn26KE[0:eUo6AxKM7], stdISO8601ColonSecondsTZ, pnn26KE[eUo6AxKM7+9:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+5 && pnn26KE[eUo6AxKM7:eUo6AxKM7+5] == "Z0700" {
				return pnn26KE[0:eUo6AxKM7], stdISO8601TZ, pnn26KE[eUo6AxKM7+5:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+6 && pnn26KE[eUo6AxKM7:eUo6AxKM7+6] == "Z07:00" {
				return pnn26KE[0:eUo6AxKM7], stdISO8601ColonTZ, pnn26KE[eUo6AxKM7+6:]
			}
			if  /*line :1*/ len(pnn26KE) >= eUo6AxKM7+3 && pnn26KE[eUo6AxKM7:eUo6AxKM7+3] == "Z07" {
				return pnn26KE[0:eUo6AxKM7], stdISO8601ShortTZ, pnn26KE[eUo6AxKM7+3:]
			}

		case '.', ',':
			if eUo6AxKM7+1 <  /*line :1*/ len(pnn26KE) && (pnn26KE[eUo6AxKM7+1] == '0' || pnn26KE[eUo6AxKM7+1] == '9') {
				qRGfpKGb7 := pnn26KE[eUo6AxKM7+1]
				a4m2XndGEF := eUo6AxKM7 + 1
				for a4m2XndGEF <  /*line :1*/ len(pnn26KE) && pnn26KE[a4m2XndGEF] == qRGfpKGb7 {
					a4m2XndGEF++
				}

				if ! /*line :1*/ tcM28YO(pnn26KE, a4m2XndGEF) {
					p3QK5I2Eat2g := stdFracSecond0
					if pnn26KE[eUo6AxKM7+1] == '9' {
						p3QK5I2Eat2g = stdFracSecond9
					}
					gA16DoaF_ :=  /*line :1*/ xDJV1RBUw(p3QK5I2Eat2g, a4m2XndGEF-(eUo6AxKM7+1), fnpIad3KB)
					return pnn26KE[0:eUo6AxKM7], gA16DoaF_, pnn26KE[a4m2XndGEF:]
				}
			}
		}
	}
	return pnn26KE, 0, ""
}

var rXk2xR7P = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var fE9uPIZtK = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var jjGm41 = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var sxLand = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}



func gtQHEGBBUNq9(pXtmZMPAs, oIR0Es4 string) bool {
	for eUo6AxKM7 := 0; eUo6AxKM7 <  /*line :1*/ len(pXtmZMPAs); eUo6AxKM7++ {
		abMZcfUe := pXtmZMPAs[eUo6AxKM7]
		qtfH6Fy6l8_J := oIR0Es4[eUo6AxKM7]
		if abMZcfUe != qtfH6Fy6l8_J {

			abMZcfUe |= 'a' - 'A'
			qtfH6Fy6l8_J |= 'a' - 'A'
			if abMZcfUe != qtfH6Fy6l8_J || abMZcfUe < 'a' || abMZcfUe > 'z' {
				return false
			}
		}
	}
	return true
}

func oPdG5Tk5(ijpe9XOM []string, sX0wK447yDp string) (int, string, error) {
	for eUo6AxKM7, ha9RaMigT := range ijpe9XOM {
		if  /*line :1*/ len(sX0wK447yDp) >=  /*line :1*/ len(ha9RaMigT) &&  /*line :1*/ gtQHEGBBUNq9(sX0wK447yDp[0: /*line :1*/ len(ha9RaMigT)], ha9RaMigT) {
			return eUo6AxKM7, sX0wK447yDp[ /*line :1*/ len(ha9RaMigT):], nil
		}
	}
	return -1, sX0wK447yDp, fqhZOO
}




func lVMi7sLX9A5(ev8znPS9W []byte, i_2w697Gglw int, dCVSgabykBJF int) []byte {
	b6XtDzh :=  /*line :1*/ uint(i_2w697Gglw)
	if i_2w697Gglw < 0 {
		ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '-')
		b6XtDzh =  /*line :1*/ uint(-i_2w697Gglw)
	}

	cM9iSCW04sJ := func(b6XtDzh uint) byte { return '0' +  /*line :1*/ byte(b6XtDzh) }
	switch {
	case dCVSgabykBJF == 2 && b6XtDzh < 1e2:
		return  /*line :1*/ append(ev8znPS9W,  /*line :1*/ cM9iSCW04sJ(b6XtDzh/1e1),  /*line :1*/ cM9iSCW04sJ(b6XtDzh%1e1))
	case dCVSgabykBJF == 4 && b6XtDzh < 1e4:
		return  /*line :1*/ append(ev8znPS9W,  /*line :1*/ cM9iSCW04sJ(b6XtDzh/1e3),  /*line :1*/ cM9iSCW04sJ(b6XtDzh/1e2%1e1),  /*line :1*/ cM9iSCW04sJ(b6XtDzh/1e1%1e1),  /*line :1*/ cM9iSCW04sJ(b6XtDzh%1e1))
	}

	
	var aeBZDJEINbb int
	if b6XtDzh == 0 {
		aeBZDJEINbb = 1
	}
	for cxUZjwPHq3tA := b6XtDzh; cxUZjwPHq3tA > 0; cxUZjwPHq3tA /= 10 {
		aeBZDJEINbb++
	}

	for _nFT6aL := dCVSgabykBJF - aeBZDJEINbb; _nFT6aL > 0; _nFT6aL-- {
		ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '0')
	}

	if  /*line :1*/ len(ev8znPS9W)+aeBZDJEINbb <=  /*line :1*/ cap(ev8znPS9W) {
		ev8znPS9W = ev8znPS9W[: /*line :1*/ len(ev8znPS9W)+aeBZDJEINbb]
	} else {
		ev8znPS9W =  /*line :1*/ append(ev8znPS9W,  /*line :1*/ make([]byte, aeBZDJEINbb)...)
	}

	eUo6AxKM7 :=  /*line :1*/ len(ev8znPS9W) - 1
	for b6XtDzh >= 10 && eUo6AxKM7 > 0 {
		furnSu := b6XtDzh / 10
		ev8znPS9W[eUo6AxKM7] =  /*line :1*/ cM9iSCW04sJ(b6XtDzh - furnSu*10)
		b6XtDzh = furnSu
		eUo6AxKM7--
	}
	ev8znPS9W[eUo6AxKM7] =  /*line :1*/ cM9iSCW04sJ(b6XtDzh)
	return ev8znPS9W
}


var quJ0kLP6IhS =  /*line :1*/ errors.Su6g6hRoi9X("time: invalid number")


func sZ5HiNLZ65[xl2dJlVxAz []byte | string](cjQd5EdPBZv xl2dJlVxAz) (i_2w697Gglw int, xuMLYrB error) {
	yY3pgXWox3j := false
	if  /*line :1*/ len(cjQd5EdPBZv) > 0 && (cjQd5EdPBZv[0] == '-' || cjQd5EdPBZv[0] == '+') {
		yY3pgXWox3j = cjQd5EdPBZv[0] == '-'
		cjQd5EdPBZv = cjQd5EdPBZv[1:]
	}
	furnSu, sTMtN6w6Z, xuMLYrB :=  /*line :1*/ wQR8Ii5R(cjQd5EdPBZv)
	i_2w697Gglw =  /*line :1*/ int(furnSu)
	if xuMLYrB != nil ||  /*line :1*/ len(sTMtN6w6Z) > 0 {
		return 0, quJ0kLP6IhS
	}
	if yY3pgXWox3j {
		i_2w697Gglw = -i_2w697Gglw
	}
	return i_2w697Gglw, nil
}




func xDJV1RBUw(p3QK5I2Eat2g, aeBZDJEINbb, fnpIad3KB int) int {

	if fnpIad3KB == '.' {
		return p3QK5I2Eat2g | ((aeBZDJEINbb & 0xfff) << stdArgShift)
	}
	return p3QK5I2Eat2g | ((aeBZDJEINbb & 0xfff) << stdArgShift) | 1<<stdSeparatorShift
}

func dGtMX7(gA16DoaF_ int) int {
	return (gA16DoaF_ >> stdArgShift) & 0xfff
}

func j8ywiSt(gA16DoaF_ int) byte {
	if (gA16DoaF_ >> stdSeparatorShift) == 0 {
		return '.'
	}
	return ','
}



func qlW1Bw(ev8znPS9W []byte, gEEQWI int, gA16DoaF_ int) []byte {
	ab3WDhm := gA16DoaF_&stdMask == stdFracSecond9
	aeBZDJEINbb :=  /*line :1*/ dGtMX7(gA16DoaF_)
	if ab3WDhm && (aeBZDJEINbb == 0 || gEEQWI == 0) {
		return ev8znPS9W
	}
	gTRaGuI :=  /*line :1*/ j8ywiSt(gA16DoaF_)
	ev8znPS9W =  /*line :1*/ append(ev8znPS9W, gTRaGuI)
	ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, gEEQWI, 9)
	if aeBZDJEINbb < 9 {
		ev8znPS9W = ev8znPS9W[: /*line :1*/ len(ev8znPS9W)-9+aeBZDJEINbb]
	}
	if ab3WDhm {
		for  /*line :1*/ len(ev8znPS9W) > 0 && ev8znPS9W[ /*line :1*/ len(ev8znPS9W)-1] == '0' {
			ev8znPS9W = ev8znPS9W[: /*line :1*/ len(ev8znPS9W)-1]
		}
		if  /*line :1*/ len(ev8znPS9W) > 0 && ev8znPS9W[ /*line :1*/ len(ev8znPS9W)-1] == gTRaGuI {
			ev8znPS9W = ev8znPS9W[: /*line :1*/ len(ev8znPS9W)-1]
		}
	}
	return ev8znPS9W
}












func (pv6eev G4KDkI) String() string {
	cjQd5EdPBZv :=  /*line :1*/ pv6eev.Format("2006-01-02 15:04:05.999999999 -0700 MST")

	if pv6eev.iB2EPt9&hasMonotonic != 0 {
		b7KRWO52CKc :=  /*line :1*/ uint64(pv6eev.f0xL4nVadlv)
		gW4A6Mt59KV :=  /*line :1*/ byte('+')
		if pv6eev.f0xL4nVadlv < 0 {
			gW4A6Mt59KV = '-'
			b7KRWO52CKc = -b7KRWO52CKc
		}
		s7ZszeDNuqt, b7KRWO52CKc := b7KRWO52CKc/1e9, b7KRWO52CKc%1e9
		ebFP7Ni, s7ZszeDNuqt := s7ZszeDNuqt/1e9, s7ZszeDNuqt%1e9
		qtO6L35R9b :=  /*line :1*/ make([]byte, 0, 24)
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, " m="...)
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, gW4A6Mt59KV)
		neKYHeM1Tee := 0
		if ebFP7Ni != 0 {
			qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b,  /*line :1*/ int(ebFP7Ni), 0)
			neKYHeM1Tee = 9
		}
		qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b,  /*line :1*/ int(s7ZszeDNuqt), neKYHeM1Tee)
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, '.')
		qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b,  /*line :1*/ int(b7KRWO52CKc), 9)
		cjQd5EdPBZv +=  /*line :1*/ string(qtO6L35R9b)
	}
	return cjQd5EdPBZv
}



func (pv6eev G4KDkI) GoString() string {
	nsu4XSEmD7r :=  /*line :1*/ pv6eev.nsu4XSEmD7r()
	mO2sH8hUHnW, cmQVCZ2yilS, gv5Z468R, _ :=  /*line :1*/ fdO09EiS(nsu4XSEmD7r, true)
	xmMDTE, a4cjaldFo, f9yTjNTQ6sc :=  /*line :1*/ mbFqMakX(nsu4XSEmD7r)

	qtO6L35R9b :=  /*line :1*/ make([]byte, 0,  /*line :1*/ len("time.Date(9999, time.September, 31, 23, 59, 59, 999999999, time.Local)"))
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, "time.Date("...)
	qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b, mO2sH8hUHnW, 0)
	if January <= cmQVCZ2yilS && cmQVCZ2yilS <= December {
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ", time."...)
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, sxLand[cmQVCZ2yilS-1]...)
	} else {

		qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b,  /*line :1*/ int(cmQVCZ2yilS), 0)
	}
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ", "...)
	qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b, gv5Z468R, 0)
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ", "...)
	qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b, xmMDTE, 0)
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ", "...)
	qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b, a4cjaldFo, 0)
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ", "...)
	qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b, f9yTjNTQ6sc, 0)
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ", "...)
	qtO6L35R9b =  /*line :1*/ lVMi7sLX9A5(qtO6L35R9b,  /*line :1*/ pv6eev.Nanosecond(), 0)
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ", "...)
	switch sAqqFTcly5S :=  /*line :1*/ pv6eev.Location(); sAqqFTcly5S {
	case PD1NIUjTJ, nil:
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, "time.UTC"...)
	case CH9afbXxKDfu:
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, "time.Local"...)
	default:

		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, `time.Location(`...)
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b,  /*line :1*/ cJ8pucH5Vq(sAqqFTcly5S.yiQ5l0TL_)...)
		qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ')')
	}
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, ')')
	return  /*line :1*/ string(qtO6L35R9b)
}







func (pv6eev G4KDkI) Format(pnn26KE string) string {
	const bufSize = 64
	var ev8znPS9W []byte
	aRi90tFb0vWu :=  /*line :1*/ len(pnn26KE) + 10
	if aRi90tFb0vWu < bufSize {
		var qtO6L35R9b [bufSize]byte
		ev8znPS9W = qtO6L35R9b[:0]
	} else {
		ev8znPS9W =  /*line :1*/ make([]byte, 0, aRi90tFb0vWu)
	}
	ev8znPS9W =  /*line :1*/ pv6eev.AppendFormat(ev8znPS9W, pnn26KE)
	return  /*line :1*/ string(ev8znPS9W)
}



func (pv6eev G4KDkI) AppendFormat(ev8znPS9W []byte, pnn26KE string) []byte {

	switch pnn26KE {
	case RFC3339:
		return  /*line :1*/ pv6eev.kigfvy0(ev8znPS9W, false)
	case RFC3339Nano:
		return  /*line :1*/ pv6eev.kigfvy0(ev8znPS9W, true)
	default:
		return  /*line :1*/ pv6eev.wPj3Q1(ev8znPS9W, pnn26KE)
	}
}

func (pv6eev G4KDkI) wPj3Q1(ev8znPS9W []byte, pnn26KE string) []byte {
	var (
		w12CVdh9, l7D_bh2w, nsu4XSEmD7r		=  /*line :1*/ pv6eev.bNr1Yr()

		mO2sH8hUHnW	int	= -1
		cmQVCZ2yilS	ZyPFXNxLcwpT
		gv5Z468R	int
		rAwk8hU1NKXu	int
		xmMDTE	int	= -1
		qJ3jbB5O	int
		insjC2Va	int
	)

	for pnn26KE != "" {
		hcap7N, gA16DoaF_, cAS4S6 :=  /*line :1*/ gQkzYn3i(pnn26KE)
		if hcap7N != "" {
			ev8znPS9W =  /*line :1*/ append(ev8znPS9W, hcap7N...)
		}
		if gA16DoaF_ == 0 {
			break
		}
		pnn26KE = cAS4S6

		if mO2sH8hUHnW < 0 && gA16DoaF_&stdNeedDate != 0 {
			mO2sH8hUHnW, cmQVCZ2yilS, gv5Z468R, rAwk8hU1NKXu =  /*line :1*/ fdO09EiS(nsu4XSEmD7r, true)
			rAwk8hU1NKXu++
		}

		if xmMDTE < 0 && gA16DoaF_&stdNeedClock != 0 {
			xmMDTE, qJ3jbB5O, insjC2Va =  /*line :1*/ mbFqMakX(nsu4XSEmD7r)
		}

		switch gA16DoaF_ & stdMask {
		case stdYear:
			bRAhk8mArR1z := mO2sH8hUHnW
			if bRAhk8mArR1z < 0 {
				bRAhk8mArR1z = -bRAhk8mArR1z
			}
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, bRAhk8mArR1z%100, 2)
		case stdLongYear:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, mO2sH8hUHnW, 4)
		case stdMonth:
			ev8znPS9W =  /*line :1*/ append(ev8znPS9W,  /*line :1*/ cmQVCZ2yilS.String()[:3]...)
		case stdLongMonth:
			fYy09z7m :=  /*line :1*/ cmQVCZ2yilS.String()
			ev8znPS9W =  /*line :1*/ append(ev8znPS9W, fYy09z7m...)
		case stdNumMonth:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W,  /*line :1*/ int(cmQVCZ2yilS), 0)
		case stdZeroMonth:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W,  /*line :1*/ int(cmQVCZ2yilS), 2)
		case stdWeekDay:
			ev8znPS9W =  /*line :1*/ append(ev8znPS9W,  /*line :1*/ baS3QC6JH(nsu4XSEmD7r).String()[:3]...)
		case stdLongWeekDay:
			cjQd5EdPBZv :=  /*line :1*/ baS3QC6JH(nsu4XSEmD7r).String()
			ev8znPS9W =  /*line :1*/ append(ev8znPS9W, cjQd5EdPBZv...)
		case stdDay:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, gv5Z468R, 0)
		case stdUnderDay:
			if gv5Z468R < 10 {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ' ')
			}
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, gv5Z468R, 0)
		case stdZeroDay:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, gv5Z468R, 2)
		case stdUnderYearDay:
			if rAwk8hU1NKXu < 100 {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ' ')
				if rAwk8hU1NKXu < 10 {
					ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ' ')
				}
			}
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, rAwk8hU1NKXu, 0)
		case stdZeroYearDay:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, rAwk8hU1NKXu, 3)
		case stdHour:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, xmMDTE, 2)
		case stdHour12:

			ggP7eA6hbC1 := xmMDTE % 12
			if ggP7eA6hbC1 == 0 {
				ggP7eA6hbC1 = 12
			}
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, ggP7eA6hbC1, 0)
		case stdZeroHour12:

			ggP7eA6hbC1 := xmMDTE % 12
			if ggP7eA6hbC1 == 0 {
				ggP7eA6hbC1 = 12
			}
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, ggP7eA6hbC1, 2)
		case stdMinute:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, qJ3jbB5O, 0)
		case stdZeroMinute:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, qJ3jbB5O, 2)
		case stdSecond:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, insjC2Va, 0)
		case stdZeroSecond:
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, insjC2Va, 2)
		case stdPM:
			if xmMDTE >= 12 {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, "PM"...)
			} else {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, "AM"...)
			}
		case stdpm:
			if xmMDTE >= 12 {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, "pm"...)
			} else {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, "am"...)
			}
		case stdISO8601TZ, stdISO8601ColonTZ, stdISO8601SecondsTZ, stdISO8601ShortTZ, stdISO8601ColonSecondsTZ, stdNumTZ, stdNumColonTZ, stdNumSecondsTz, stdNumShortTZ, stdNumColonSecondsTZ:

			if l7D_bh2w == 0 && (gA16DoaF_ == stdISO8601TZ || gA16DoaF_ == stdISO8601ColonTZ || gA16DoaF_ == stdISO8601SecondsTZ || gA16DoaF_ == stdISO8601ShortTZ || gA16DoaF_ == stdISO8601ColonSecondsTZ) {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, 'Z')
				break
			}
			k3sFK3l := l7D_bh2w / 60
			aOS98R := l7D_bh2w
			if k3sFK3l < 0 {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '-')
				k3sFK3l = -k3sFK3l
				aOS98R = -aOS98R
			} else {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '+')
			}
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, k3sFK3l/60, 2)
			if gA16DoaF_ == stdISO8601ColonTZ || gA16DoaF_ == stdNumColonTZ || gA16DoaF_ == stdISO8601ColonSecondsTZ || gA16DoaF_ == stdNumColonSecondsTZ {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ':')
			}
			if gA16DoaF_ != stdNumShortTZ && gA16DoaF_ != stdISO8601ShortTZ {
				ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, k3sFK3l%60, 2)
			}

			if gA16DoaF_ == stdISO8601SecondsTZ || gA16DoaF_ == stdNumSecondsTz || gA16DoaF_ == stdNumColonSecondsTZ || gA16DoaF_ == stdISO8601ColonSecondsTZ {
				if gA16DoaF_ == stdNumColonSecondsTZ || gA16DoaF_ == stdISO8601ColonSecondsTZ {
					ev8znPS9W =  /*line :1*/ append(ev8znPS9W, ':')
				}
				ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, aOS98R%60, 2)
			}

		case stdTZ:
			if w12CVdh9 != "" {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, w12CVdh9...)
				break
			}

			k3sFK3l := l7D_bh2w / 60
			if k3sFK3l < 0 {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '-')
				k3sFK3l = -k3sFK3l
			} else {
				ev8znPS9W =  /*line :1*/ append(ev8znPS9W, '+')
			}
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, k3sFK3l/60, 2)
			ev8znPS9W =  /*line :1*/ lVMi7sLX9A5(ev8znPS9W, k3sFK3l%60, 2)
		case stdFracSecond0, stdFracSecond9:
			ev8znPS9W =  /*line :1*/ qlW1Bw(ev8znPS9W,  /*line :1*/ pv6eev.Nanosecond(), gA16DoaF_)
		}
	}
	return ev8znPS9W
}

var fqhZOO =  /*line :1*/ errors.Su6g6hRoi9X("bad value for field")	


type Hcvjzaf83uX struct {
	DlSQTL70	string
	Kcg9omNj	string
	XeJSBjb4ia1	string
	E3s2XbdRn	string
	UyAvhhRGH	string
}



func nsvlWmZm1Ccz(pnn26KE, kHQbwY7, jN4o8zq, tqDicZbE, j0ga5KDymj9 string) *Hcvjzaf83uX {
	k_HiFZ9 :=  /*line :1*/ e97oTnyWvr(kHQbwY7)
	jnZ9h1s :=  /*line :1*/ e97oTnyWvr(tqDicZbE)
	return &Hcvjzaf83uX{pnn26KE, k_HiFZ9, jN4o8zq, jnZ9h1s, j0ga5KDymj9}
}



func e97oTnyWvr(cjQd5EdPBZv string) string {
	return  /*line :1*/ string([] /*line :1*/ byte(cjQd5EdPBZv))
}



const (
	lowerhex	= "0123456789abcdef"
	runeSelf	= 0x80
	runeError	= '\uFFFD'
)

func cJ8pucH5Vq(cjQd5EdPBZv string) string {
	qtO6L35R9b :=  /*line :1*/ make([]byte, 1,  /*line :1*/ len(cjQd5EdPBZv)+2)
	qtO6L35R9b[0] = '"'
	for eUo6AxKM7, fnpIad3KB := range cjQd5EdPBZv {
		if fnpIad3KB >= runeSelf || fnpIad3KB < ' ' {
			
			
			
			
			
			
			var dCVSgabykBJF int
			if fnpIad3KB == runeError {
				dCVSgabykBJF = 1
				if eUo6AxKM7+2 <  /*line :1*/ len(cjQd5EdPBZv) && cjQd5EdPBZv[eUo6AxKM7:eUo6AxKM7+3] ==  /*line :1*/ string(runeError) {
					dCVSgabykBJF = 3
				}
			} else {
				dCVSgabykBJF =  /*line :1*/ len( /*line :1*/ string(fnpIad3KB))
			}
			for a4m2XndGEF := 0; a4m2XndGEF < dCVSgabykBJF; a4m2XndGEF++ {
				qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, `\x`...)
				qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, lowerhex[cjQd5EdPBZv[eUo6AxKM7+a4m2XndGEF]>>4])
				qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, lowerhex[cjQd5EdPBZv[eUo6AxKM7+a4m2XndGEF]&0xF])
			}
		} else {
			if fnpIad3KB == '"' || fnpIad3KB == '\\' {
				qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, '\\')
			}
			qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b,  /*line :1*/ string(fnpIad3KB)...)
		}
	}
	qtO6L35R9b =  /*line :1*/ append(qtO6L35R9b, '"')
	return  /*line :1*/ string(qtO6L35R9b)
}


func (iFmVByQ *Hcvjzaf83uX) Error() string {
	if iFmVByQ.UyAvhhRGH == "" {
		return "parsing time " +
			 /*line :1*/ cJ8pucH5Vq(iFmVByQ.Kcg9omNj) + " as " +
			 /*line :1*/ cJ8pucH5Vq(iFmVByQ.DlSQTL70) + ": cannot parse " +
			 /*line :1*/ cJ8pucH5Vq(iFmVByQ.E3s2XbdRn) + " as " +
			 /*line :1*/ cJ8pucH5Vq(iFmVByQ.XeJSBjb4ia1)
	}
	return "parsing time " +
		 /*line :1*/ cJ8pucH5Vq(iFmVByQ.Kcg9omNj) + iFmVByQ.UyAvhhRGH
}


func tcM28YO[xl2dJlVxAz []byte | string](cjQd5EdPBZv xl2dJlVxAz, eUo6AxKM7 int) bool {
	if  /*line :1*/ len(cjQd5EdPBZv) <= eUo6AxKM7 {
		return false
	}
	fnpIad3KB := cjQd5EdPBZv[eUo6AxKM7]
	return '0' <= fnpIad3KB && fnpIad3KB <= '9'
}




func d9MdYy(cjQd5EdPBZv string, juBV3gQXJx1 bool) (int, string, error) {
	if ! /*line :1*/ tcM28YO(cjQd5EdPBZv, 0) {
		return 0, cjQd5EdPBZv, fqhZOO
	}
	if ! /*line :1*/ tcM28YO(cjQd5EdPBZv, 1) {
		if juBV3gQXJx1 {
			return 0, cjQd5EdPBZv, fqhZOO
		}
		return  /*line :1*/ int(cjQd5EdPBZv[0] - '0'), cjQd5EdPBZv[1:], nil
	}
	return  /*line :1*/ int(cjQd5EdPBZv[0]-'0')*10 +  /*line :1*/ int(cjQd5EdPBZv[1]-'0'), cjQd5EdPBZv[2:], nil
}




func wSxk9Co5(cjQd5EdPBZv string, juBV3gQXJx1 bool) (int, string, error) {
	var aeBZDJEINbb, eUo6AxKM7 int
	for eUo6AxKM7 = 0; eUo6AxKM7 < 3 &&  /*line :1*/ tcM28YO(cjQd5EdPBZv, eUo6AxKM7); eUo6AxKM7++ {
		aeBZDJEINbb = aeBZDJEINbb*10 +  /*line :1*/ int(cjQd5EdPBZv[eUo6AxKM7]-'0')
	}
	if eUo6AxKM7 == 0 || juBV3gQXJx1 && eUo6AxKM7 != 3 {
		return 0, cjQd5EdPBZv, fqhZOO
	}
	return aeBZDJEINbb, cjQd5EdPBZv[eUo6AxKM7:], nil
}

func jNVxEU9(cjQd5EdPBZv string) string {
	for  /*line :1*/ len(cjQd5EdPBZv) > 0 && cjQd5EdPBZv[0] == ' ' {
		cjQd5EdPBZv = cjQd5EdPBZv[1:]
	}
	return cjQd5EdPBZv
}



func iWW7p4CGpUo(kHQbwY7, hcap7N string) (string, error) {
	for  /*line :1*/ len(hcap7N) > 0 {
		if hcap7N[0] == ' ' {
			if  /*line :1*/ len(kHQbwY7) > 0 && kHQbwY7[0] != ' ' {
				return kHQbwY7, fqhZOO
			}
			hcap7N =  /*line :1*/ jNVxEU9(hcap7N)
			kHQbwY7 =  /*line :1*/ jNVxEU9(kHQbwY7)
			continue
		}
		if  /*line :1*/ len(kHQbwY7) == 0 || kHQbwY7[0] != hcap7N[0] {
			return kHQbwY7, fqhZOO
		}
		hcap7N = hcap7N[1:]
		kHQbwY7 = kHQbwY7[1:]
	}
	return kHQbwY7, nil
}











































func KoKD_YnCna(pnn26KE, kHQbwY7 string) (G4KDkI, error) {

	if pnn26KE == RFC3339 || pnn26KE == RFC3339Nano {
		if pv6eev, ufWfJe3qPw :=  /*line :1*/ zlEkeU(kHQbwY7, CH9afbXxKDfu); ufWfJe3qPw {
			return pv6eev, nil
		}
	}
	return  /*line :1*/ meWLZ6R(pnn26KE, kHQbwY7, PD1NIUjTJ, CH9afbXxKDfu)
}






func Un8pGpo9(pnn26KE, kHQbwY7 string, sAqqFTcly5S *Hh4U1oidou) (G4KDkI, error) {

	if pnn26KE == RFC3339 || pnn26KE == RFC3339Nano {
		if pv6eev, ufWfJe3qPw :=  /*line :1*/ zlEkeU(kHQbwY7, sAqqFTcly5S); ufWfJe3qPw {
			return pv6eev, nil
		}
	}
	return  /*line :1*/ meWLZ6R(pnn26KE, kHQbwY7, sAqqFTcly5S, sAqqFTcly5S)
}

func meWLZ6R(pnn26KE, kHQbwY7 string, jqZkPYGCQd5x, b2m6A1Q *Hh4U1oidou) (G4KDkI, error) {
	alPAJX811f, i9bciiyZA_XF := pnn26KE, kHQbwY7
	nCUJoVW := ""
	pkdzoa9EM_V := false
	aYVnRHp := false

	
	var (
		mO2sH8hUHnW	int
		cmQVCZ2yilS	int	= -1
		gv5Z468R	int	= -1
		rAwk8hU1NKXu	int	= -1
		xmMDTE	int
		qJ3jbB5O	int
		insjC2Va	int
		aCiH2Z7qWLtw	int
		iL2jRw3Faxe	*Hh4U1oidou
		zP6aWDTN	int	= -1
		rhoSRkYnF	string
	)

	for {
		var xuMLYrB error
		hcap7N, gA16DoaF_, cAS4S6 :=  /*line :1*/ gQkzYn3i(pnn26KE)
		cZo77kIhbx := pnn26KE[ /*line :1*/ len(hcap7N) :  /*line :1*/ len(pnn26KE)- /*line :1*/ len(cAS4S6)]
		kHQbwY7, xuMLYrB =  /*line :1*/ iWW7p4CGpUo(kHQbwY7, hcap7N)
		if xuMLYrB != nil {
			return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, hcap7N, kHQbwY7, "")
		}
		if gA16DoaF_ == 0 {
			if  /*line :1*/ len(kHQbwY7) != 0 {
				return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, "", kHQbwY7, ": extra text: "+ /*line :1*/ cJ8pucH5Vq(kHQbwY7))
			}
			break
		}
		pnn26KE = cAS4S6
		var _STZiF string
		fBMKzCDSC := kHQbwY7
		switch gA16DoaF_ & stdMask {
		case stdYear:
			if  /*line :1*/ len(kHQbwY7) < 2 {
				xuMLYrB = fqhZOO
				break
			}
			_STZiF, kHQbwY7 = kHQbwY7[0:2], kHQbwY7[2:]
			mO2sH8hUHnW, xuMLYrB =  /*line :1*/ sZ5HiNLZ65(_STZiF)
			if xuMLYrB != nil {
				break
			}
			if mO2sH8hUHnW >= 69 {
				mO2sH8hUHnW += 1900
			} else {
				mO2sH8hUHnW += 2000
			}
		case stdLongYear:
			if  /*line :1*/ len(kHQbwY7) < 4 || ! /*line :1*/ tcM28YO(kHQbwY7, 0) {
				xuMLYrB = fqhZOO
				break
			}
			_STZiF, kHQbwY7 = kHQbwY7[0:4], kHQbwY7[4:]
			mO2sH8hUHnW, xuMLYrB =  /*line :1*/ sZ5HiNLZ65(_STZiF)
		case stdMonth:
			cmQVCZ2yilS, kHQbwY7, xuMLYrB =  /*line :1*/ oPdG5Tk5(jjGm41, kHQbwY7)
			cmQVCZ2yilS++
		case stdLongMonth:
			cmQVCZ2yilS, kHQbwY7, xuMLYrB =  /*line :1*/ oPdG5Tk5(sxLand, kHQbwY7)
			cmQVCZ2yilS++
		case stdNumMonth, stdZeroMonth:
			cmQVCZ2yilS, kHQbwY7, xuMLYrB =  /*line :1*/ d9MdYy(kHQbwY7, gA16DoaF_ == stdZeroMonth)
			if xuMLYrB == nil && (cmQVCZ2yilS <= 0 || 12 < cmQVCZ2yilS) {
				nCUJoVW = "month"
			}
		case stdWeekDay:

			_, kHQbwY7, xuMLYrB =  /*line :1*/ oPdG5Tk5(fE9uPIZtK, kHQbwY7)
		case stdLongWeekDay:
			_, kHQbwY7, xuMLYrB =  /*line :1*/ oPdG5Tk5(rXk2xR7P, kHQbwY7)
		case stdDay, stdUnderDay, stdZeroDay:
			if gA16DoaF_ == stdUnderDay &&  /*line :1*/ len(kHQbwY7) > 0 && kHQbwY7[0] == ' ' {
				kHQbwY7 = kHQbwY7[1:]
			}
			gv5Z468R, kHQbwY7, xuMLYrB =  /*line :1*/ d9MdYy(kHQbwY7, gA16DoaF_ == stdZeroDay)

		case stdUnderYearDay, stdZeroYearDay:
			for eUo6AxKM7 := 0; eUo6AxKM7 < 2; eUo6AxKM7++ {
				if gA16DoaF_ == stdUnderYearDay &&  /*line :1*/ len(kHQbwY7) > 0 && kHQbwY7[0] == ' ' {
					kHQbwY7 = kHQbwY7[1:]
				}
			}
			rAwk8hU1NKXu, kHQbwY7, xuMLYrB =  /*line :1*/ wSxk9Co5(kHQbwY7, gA16DoaF_ == stdZeroYearDay)

		case stdHour:
			xmMDTE, kHQbwY7, xuMLYrB =  /*line :1*/ d9MdYy(kHQbwY7, false)
			if xmMDTE < 0 || 24 <= xmMDTE {
				nCUJoVW = "hour"
			}
		case stdHour12, stdZeroHour12:
			xmMDTE, kHQbwY7, xuMLYrB =  /*line :1*/ d9MdYy(kHQbwY7, gA16DoaF_ == stdZeroHour12)
			if xmMDTE < 0 || 12 < xmMDTE {
				nCUJoVW = "hour"
			}
		case stdMinute, stdZeroMinute:
			qJ3jbB5O, kHQbwY7, xuMLYrB =  /*line :1*/ d9MdYy(kHQbwY7, gA16DoaF_ == stdZeroMinute)
			if qJ3jbB5O < 0 || 60 <= qJ3jbB5O {
				nCUJoVW = "minute"
			}
		case stdSecond, stdZeroSecond:
			insjC2Va, kHQbwY7, xuMLYrB =  /*line :1*/ d9MdYy(kHQbwY7, gA16DoaF_ == stdZeroSecond)
			if xuMLYrB != nil {
				break
			}
			if insjC2Va < 0 || 60 <= insjC2Va {
				nCUJoVW = "second"
				break
			}

			if  /*line :1*/ len(kHQbwY7) >= 2 &&  /*line :1*/ cJ4lAKOqaK1(kHQbwY7[0]) &&  /*line :1*/ tcM28YO(kHQbwY7, 1) {
				_, gA16DoaF_, _ =  /*line :1*/ gQkzYn3i(pnn26KE)
				gA16DoaF_ &= stdMask
				if gA16DoaF_ == stdFracSecond0 || gA16DoaF_ == stdFracSecond9 {

					break
				}

				aeBZDJEINbb := 2
				for ; aeBZDJEINbb <  /*line :1*/ len(kHQbwY7) &&  /*line :1*/ tcM28YO(kHQbwY7, aeBZDJEINbb); aeBZDJEINbb++ {
				}
				aCiH2Z7qWLtw, nCUJoVW, xuMLYrB =  /*line :1*/ hmaS0CqyQ8oh(kHQbwY7, aeBZDJEINbb)
				kHQbwY7 = kHQbwY7[aeBZDJEINbb:]
			}
		case stdPM:
			if  /*line :1*/ len(kHQbwY7) < 2 {
				xuMLYrB = fqhZOO
				break
			}
			_STZiF, kHQbwY7 = kHQbwY7[0:2], kHQbwY7[2:]
			switch _STZiF {
			case "PM":
				aYVnRHp = true
			case "AM":
				pkdzoa9EM_V = true
			default:
				xuMLYrB = fqhZOO
			}
		case stdpm:
			if  /*line :1*/ len(kHQbwY7) < 2 {
				xuMLYrB = fqhZOO
				break
			}
			_STZiF, kHQbwY7 = kHQbwY7[0:2], kHQbwY7[2:]
			switch _STZiF {
			case "pm":
				aYVnRHp = true
			case "am":
				pkdzoa9EM_V = true
			default:
				xuMLYrB = fqhZOO
			}
		case stdISO8601TZ, stdISO8601ColonTZ, stdISO8601SecondsTZ, stdISO8601ShortTZ, stdISO8601ColonSecondsTZ, stdNumTZ, stdNumShortTZ, stdNumColonTZ, stdNumSecondsTz, stdNumColonSecondsTZ:
			if (gA16DoaF_ == stdISO8601TZ || gA16DoaF_ == stdISO8601ShortTZ || gA16DoaF_ == stdISO8601ColonTZ) &&  /*line :1*/ len(kHQbwY7) >= 1 && kHQbwY7[0] == 'Z' {
				kHQbwY7 = kHQbwY7[1:]
				iL2jRw3Faxe = PD1NIUjTJ
				break
			}
			var gW4A6Mt59KV, xmMDTE, qJ3jbB5O, iqCaE0gCthiB string
			if gA16DoaF_ == stdISO8601ColonTZ || gA16DoaF_ == stdNumColonTZ {
				if  /*line :1*/ len(kHQbwY7) < 6 {
					xuMLYrB = fqhZOO
					break
				}
				if kHQbwY7[3] != ':' {
					xuMLYrB = fqhZOO
					break
				}
				gW4A6Mt59KV, xmMDTE, qJ3jbB5O, iqCaE0gCthiB, kHQbwY7 = kHQbwY7[0:1], kHQbwY7[1:3], kHQbwY7[4:6], "00", kHQbwY7[6:]
			} else if gA16DoaF_ == stdNumShortTZ || gA16DoaF_ == stdISO8601ShortTZ {
				if  /*line :1*/ len(kHQbwY7) < 3 {
					xuMLYrB = fqhZOO
					break
				}
				gW4A6Mt59KV, xmMDTE, qJ3jbB5O, iqCaE0gCthiB, kHQbwY7 = kHQbwY7[0:1], kHQbwY7[1:3], "00", "00", kHQbwY7[3:]
			} else if gA16DoaF_ == stdISO8601ColonSecondsTZ || gA16DoaF_ == stdNumColonSecondsTZ {
				if  /*line :1*/ len(kHQbwY7) < 9 {
					xuMLYrB = fqhZOO
					break
				}
				if kHQbwY7[3] != ':' || kHQbwY7[6] != ':' {
					xuMLYrB = fqhZOO
					break
				}
				gW4A6Mt59KV, xmMDTE, qJ3jbB5O, iqCaE0gCthiB, kHQbwY7 = kHQbwY7[0:1], kHQbwY7[1:3], kHQbwY7[4:6], kHQbwY7[7:9], kHQbwY7[9:]
			} else if gA16DoaF_ == stdISO8601SecondsTZ || gA16DoaF_ == stdNumSecondsTz {
				if  /*line :1*/ len(kHQbwY7) < 7 {
					xuMLYrB = fqhZOO
					break
				}
				gW4A6Mt59KV, xmMDTE, qJ3jbB5O, iqCaE0gCthiB, kHQbwY7 = kHQbwY7[0:1], kHQbwY7[1:3], kHQbwY7[3:5], kHQbwY7[5:7], kHQbwY7[7:]
			} else {
				if  /*line :1*/ len(kHQbwY7) < 5 {
					xuMLYrB = fqhZOO
					break
				}
				gW4A6Mt59KV, xmMDTE, qJ3jbB5O, iqCaE0gCthiB, kHQbwY7 = kHQbwY7[0:1], kHQbwY7[1:3], kHQbwY7[3:5], "00", kHQbwY7[5:]
			}
			var ggP7eA6hbC1, cVr0qo, s4962lJUq int
			ggP7eA6hbC1, _, xuMLYrB =  /*line :1*/ d9MdYy(xmMDTE, true)
			if xuMLYrB == nil {
				cVr0qo, _, xuMLYrB =  /*line :1*/ d9MdYy(qJ3jbB5O, true)
			}
			if xuMLYrB == nil {
				s4962lJUq, _, xuMLYrB =  /*line :1*/ d9MdYy(iqCaE0gCthiB, true)
			}
			zP6aWDTN = (ggP7eA6hbC1*60+cVr0qo)*60 + s4962lJUq
			switch gW4A6Mt59KV[0] {
			case '+':
			case '-':
				zP6aWDTN = -zP6aWDTN
			default:
				xuMLYrB = fqhZOO
			}
		case stdTZ:

			if  /*line :1*/ len(kHQbwY7) >= 3 && kHQbwY7[0:3] == "UTC" {
				iL2jRw3Faxe = PD1NIUjTJ
				kHQbwY7 = kHQbwY7[3:]
				break
			}
			aeBZDJEINbb, ufWfJe3qPw :=  /*line :1*/ j1tAxF(kHQbwY7)
			if !ufWfJe3qPw {
				xuMLYrB = fqhZOO
				break
			}
			rhoSRkYnF, kHQbwY7 = kHQbwY7[:aeBZDJEINbb], kHQbwY7[aeBZDJEINbb:]

		case stdFracSecond0:

			w4r3Nt := 1 +  /*line :1*/ dGtMX7(gA16DoaF_)
			if  /*line :1*/ len(kHQbwY7) < w4r3Nt {
				xuMLYrB = fqhZOO
				break
			}
			aCiH2Z7qWLtw, nCUJoVW, xuMLYrB =  /*line :1*/ hmaS0CqyQ8oh(kHQbwY7, w4r3Nt)
			kHQbwY7 = kHQbwY7[w4r3Nt:]

		case stdFracSecond9:
			if  /*line :1*/ len(kHQbwY7) < 2 || ! /*line :1*/ cJ4lAKOqaK1(kHQbwY7[0]) || kHQbwY7[1] < '0' || '9' < kHQbwY7[1] {

				break
			}

			eUo6AxKM7 := 0
			for eUo6AxKM7+1 <  /*line :1*/ len(kHQbwY7) && '0' <= kHQbwY7[eUo6AxKM7+1] && kHQbwY7[eUo6AxKM7+1] <= '9' {
				eUo6AxKM7++
			}
			aCiH2Z7qWLtw, nCUJoVW, xuMLYrB =  /*line :1*/ hmaS0CqyQ8oh(kHQbwY7, 1+eUo6AxKM7)
			kHQbwY7 = kHQbwY7[1+eUo6AxKM7:]
		}
		if nCUJoVW != "" {
			return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, cZo77kIhbx, kHQbwY7, ": "+nCUJoVW+" out of range")
		}
		if xuMLYrB != nil {
			return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, cZo77kIhbx, fBMKzCDSC, "")
		}
	}
	if aYVnRHp && xmMDTE < 12 {
		xmMDTE += 12
	} else if pkdzoa9EM_V && xmMDTE == 12 {
		xmMDTE = 0
	}

	if rAwk8hU1NKXu >= 0 {
		var lNAwkbJQVBgU int
		var fYy09z7m int
		if  /*line :1*/ dJS2LEP6(mO2sH8hUHnW) {
			if rAwk8hU1NKXu == 31+29 {
				fYy09z7m =  /*line :1*/ int(February)
				lNAwkbJQVBgU = 29
			} else if rAwk8hU1NKXu > 31+29 {
				rAwk8hU1NKXu--
			}
		}
		if rAwk8hU1NKXu < 1 || rAwk8hU1NKXu > 365 {
			return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, "", kHQbwY7, ": day-of-year out of range")
		}
		if fYy09z7m == 0 {
			fYy09z7m = (rAwk8hU1NKXu-1)/31 + 1
			if  /*line :1*/ int(h7KunS2pF[fYy09z7m]) < rAwk8hU1NKXu {
				fYy09z7m++
			}
			lNAwkbJQVBgU = rAwk8hU1NKXu -  /*line :1*/ int(h7KunS2pF[fYy09z7m-1])
		}

		if cmQVCZ2yilS >= 0 && cmQVCZ2yilS != fYy09z7m {
			return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, "", kHQbwY7, ": day-of-year does not match month")
		}
		cmQVCZ2yilS = fYy09z7m
		if gv5Z468R >= 0 && gv5Z468R != lNAwkbJQVBgU {
			return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, "", kHQbwY7, ": day-of-year does not match day")
		}
		gv5Z468R = lNAwkbJQVBgU
	} else {
		if cmQVCZ2yilS < 0 {
			cmQVCZ2yilS =  /*line :1*/ int(January)
		}
		if gv5Z468R < 0 {
			gv5Z468R = 1
		}
	}

	if gv5Z468R < 1 || gv5Z468R >  /*line :1*/ rs4bvng4W6y( /*line :1*/ ZyPFXNxLcwpT(cmQVCZ2yilS), mO2sH8hUHnW) {
		return G4KDkI{},  /*line :1*/ nsvlWmZm1Ccz(alPAJX811f, i9bciiyZA_XF, "", kHQbwY7, ": day out of range")
	}

	if iL2jRw3Faxe != nil {
		return  /*line :1*/ FaDSoi2w(mO2sH8hUHnW,  /*line :1*/ ZyPFXNxLcwpT(cmQVCZ2yilS), gv5Z468R, xmMDTE, qJ3jbB5O, insjC2Va, aCiH2Z7qWLtw, iL2jRw3Faxe), nil
	}

	if zP6aWDTN != -1 {
		pv6eev :=  /*line :1*/ FaDSoi2w(mO2sH8hUHnW,  /*line :1*/ ZyPFXNxLcwpT(cmQVCZ2yilS), gv5Z468R, xmMDTE, qJ3jbB5O, insjC2Va, aCiH2Z7qWLtw, PD1NIUjTJ)
		 /*line :1*/ pv6eev.hOFVKuqUr91a(- /*line :1*/ int64(zP6aWDTN))

		w12CVdh9, l7D_bh2w, _, _, _ :=  /*line :1*/ b2m6A1Q.oPdG5Tk5( /*line :1*/ pv6eev.huzSe5e())
		if l7D_bh2w == zP6aWDTN && (rhoSRkYnF == "" || w12CVdh9 == rhoSRkYnF) {
			 /*line :1*/ pv6eev.vt5HqRw(b2m6A1Q)
			return pv6eev, nil
		}

		pk0qBvttHU2l :=  /*line :1*/ e97oTnyWvr(rhoSRkYnF)
		 /*line :1*/ pv6eev.vt5HqRw( /*line :1*/ Z586lZG(pk0qBvttHU2l, zP6aWDTN))
		return pv6eev, nil
	}

	if rhoSRkYnF != "" {
		pv6eev :=  /*line :1*/ FaDSoi2w(mO2sH8hUHnW,  /*line :1*/ ZyPFXNxLcwpT(cmQVCZ2yilS), gv5Z468R, xmMDTE, qJ3jbB5O, insjC2Va, aCiH2Z7qWLtw, PD1NIUjTJ)

		l7D_bh2w, ufWfJe3qPw :=  /*line :1*/ b2m6A1Q.hHPWBNyM(rhoSRkYnF,  /*line :1*/ pv6eev.huzSe5e())
		if ufWfJe3qPw {
			 /*line :1*/ pv6eev.hOFVKuqUr91a(- /*line :1*/ int64(l7D_bh2w))
			 /*line :1*/ pv6eev.vt5HqRw(b2m6A1Q)
			return pv6eev, nil
		}

		if  /*line :1*/ len(rhoSRkYnF) > 3 && rhoSRkYnF[:3] == "GMT" {
			l7D_bh2w, _ =  /*line :1*/ sZ5HiNLZ65(rhoSRkYnF[3:])
			l7D_bh2w *= 3600
		}
		pk0qBvttHU2l :=  /*line :1*/ e97oTnyWvr(rhoSRkYnF)
		 /*line :1*/ pv6eev.vt5HqRw( /*line :1*/ Z586lZG(pk0qBvttHU2l, l7D_bh2w))
		return pv6eev, nil
	}

	return  /*line :1*/ FaDSoi2w(mO2sH8hUHnW,  /*line :1*/ ZyPFXNxLcwpT(cmQVCZ2yilS), gv5Z468R, xmMDTE, qJ3jbB5O, insjC2Va, aCiH2Z7qWLtw, jqZkPYGCQd5x), nil
}











func j1tAxF(kHQbwY7 string) (f9CGC7QIWKAD int, ufWfJe3qPw bool) {
	if  /*line :1*/ len(kHQbwY7) < 3 {
		return 0, false
	}

	if  /*line :1*/ len(kHQbwY7) >= 4 && (kHQbwY7[:4] == "ChST" || kHQbwY7[:4] == "MeST") {
		return 4, true
	}

	if kHQbwY7[:3] == "GMT" {
		f9CGC7QIWKAD =  /*line :1*/ xv8Hh6C(kHQbwY7)
		return f9CGC7QIWKAD, true
	}

	if kHQbwY7[0] == '+' || kHQbwY7[0] == '-' {
		f9CGC7QIWKAD =  /*line :1*/ xqynijdG2_y(kHQbwY7)
		ufWfJe3qPw := f9CGC7QIWKAD > 0
		return f9CGC7QIWKAD, ufWfJe3qPw
	}
	
	var b9iiOR int
	for b9iiOR = 0; b9iiOR < 6; b9iiOR++ {
		if b9iiOR >=  /*line :1*/ len(kHQbwY7) {
			break
		}
		if fnpIad3KB := kHQbwY7[b9iiOR]; fnpIad3KB < 'A' || 'Z' < fnpIad3KB {
			break
		}
	}
	switch b9iiOR {
	case 0, 1, 2, 6:
		return 0, false
	case 5:
		if kHQbwY7[4] == 'T' {
			return 5, true
		}
	case 4:

		if kHQbwY7[3] == 'T' || kHQbwY7[:4] == "WITA" {
			return 4, true
		}
	case 3:
		return 3, true
	}
	return 0, false
}




func xv8Hh6C(kHQbwY7 string) int {
	kHQbwY7 = kHQbwY7[3:]
	if  /*line :1*/ len(kHQbwY7) == 0 {
		return 3
	}

	return 3 +  /*line :1*/ xqynijdG2_y(kHQbwY7)
}




func xqynijdG2_y(kHQbwY7 string) int {
	gW4A6Mt59KV := kHQbwY7[0]
	if gW4A6Mt59KV != '-' && gW4A6Mt59KV != '+' {
		return 0
	}
	i_2w697Gglw, sTMtN6w6Z, xuMLYrB :=  /*line :1*/ wQR8Ii5R(kHQbwY7[1:])

	if xuMLYrB != nil || kHQbwY7[1:] == sTMtN6w6Z {
		return 0
	}
	if i_2w697Gglw > 23 {
		return 0
	}
	return  /*line :1*/ len(kHQbwY7) -  /*line :1*/ len(sTMtN6w6Z)
}

func cJ4lAKOqaK1(ev8znPS9W byte) bool {
	return ev8znPS9W == '.' || ev8znPS9W == ','
}

func hmaS0CqyQ8oh[xl2dJlVxAz []byte | string](kHQbwY7 xl2dJlVxAz, lWSonJs int) (btzv4y3NIiBf int, nCUJoVW string, xuMLYrB error) {
	if ! /*line :1*/ cJ4lAKOqaK1(kHQbwY7[0]) {
		xuMLYrB = fqhZOO
		return
	}
	if lWSonJs > 10 {
		kHQbwY7 = kHQbwY7[:10]
		lWSonJs = 10
	}
	if btzv4y3NIiBf, xuMLYrB =  /*line :1*/ sZ5HiNLZ65(kHQbwY7[1:lWSonJs]); xuMLYrB != nil {
		return
	}
	if btzv4y3NIiBf < 0 {
		nCUJoVW = "fractional second"
		return
	}

	jawNJtZR := 10 - lWSonJs
	for eUo6AxKM7 := 0; eUo6AxKM7 < jawNJtZR; eUo6AxKM7++ {
		btzv4y3NIiBf *= 10
	}
	return
}

var em2164fz =  /*line :1*/ errors.Su6g6hRoi9X("time: bad [0-9]*")	


func wQR8Ii5R[xl2dJlVxAz []byte | string](cjQd5EdPBZv xl2dJlVxAz) (i_2w697Gglw uint64, sTMtN6w6Z xl2dJlVxAz, xuMLYrB error) {
	eUo6AxKM7 := 0
	for ; eUo6AxKM7 <  /*line :1*/ len(cjQd5EdPBZv); eUo6AxKM7++ {
		fnpIad3KB := cjQd5EdPBZv[eUo6AxKM7]
		if fnpIad3KB < '0' || fnpIad3KB > '9' {
			break
		}
		if i_2w697Gglw > 1<<63/10 {

			return 0, sTMtN6w6Z, em2164fz
		}
		i_2w697Gglw = i_2w697Gglw*10 +  /*line :1*/ uint64(fnpIad3KB) - '0'
		if i_2w697Gglw > 1<<63 {

			return 0, sTMtN6w6Z, em2164fz
		}
	}
	return i_2w697Gglw, cjQd5EdPBZv[eUo6AxKM7:], nil
}




func fcgzbefXa(cjQd5EdPBZv string) (i_2w697Gglw uint64, eIK3JAH6 float64, sTMtN6w6Z string) {
	eUo6AxKM7 := 0
	eIK3JAH6 = 1
	pM77Sdr := false
	for ; eUo6AxKM7 <  /*line :1*/ len(cjQd5EdPBZv); eUo6AxKM7++ {
		fnpIad3KB := cjQd5EdPBZv[eUo6AxKM7]
		if fnpIad3KB < '0' || fnpIad3KB > '9' {
			break
		}
		if pM77Sdr {
			continue
		}
		if i_2w697Gglw > (1<<63-1)/10 {

			pM77Sdr = true
			continue
		}
		bRAhk8mArR1z := i_2w697Gglw*10 +  /*line :1*/ uint64(fnpIad3KB) - '0'
		if bRAhk8mArR1z > 1<<63 {
			pM77Sdr = true
			continue
		}
		i_2w697Gglw = bRAhk8mArR1z
		eIK3JAH6 *= 10
	}
	return i_2w697Gglw, eIK3JAH6, cjQd5EdPBZv[eUo6AxKM7:]
}

var r3yCdn63wjx = map[string]uint64{
	"ns":	 /*line :1*/ uint64(Nanosecond),
	"us":	 /*line :1*/ uint64(Microsecond),
	"µs":	 /*line :1*/ uint64(Microsecond),
	"μs":	 /*line :1*/ uint64(Microsecond),
	"ms":	 /*line :1*/ uint64(Millisecond),
	"s":	 /*line :1*/ uint64(Second),
	"m":	 /*line :1*/ uint64(Minute),
	"h":	 /*line :1*/ uint64(Hour),
}






func CV5Qr1(cjQd5EdPBZv string) (GKMwTGxQa0S, error) {

	bVPj4pI := cjQd5EdPBZv
	var lNAwkbJQVBgU uint64
	yY3pgXWox3j := false

	if cjQd5EdPBZv != "" {
		fnpIad3KB := cjQd5EdPBZv[0]
		if fnpIad3KB == '-' || fnpIad3KB == '+' {
			yY3pgXWox3j = fnpIad3KB == '-'
			cjQd5EdPBZv = cjQd5EdPBZv[1:]
		}
	}

	if cjQd5EdPBZv == "0" {
		return 0, nil
	}
	if cjQd5EdPBZv == "" {
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
	}
	for cjQd5EdPBZv != "" {
		var (
			ha9RaMigT, k05jcfKvAvw6	uint64		
			eIK3JAH6	float64	= 1	
		)

		var xuMLYrB error

		if !(cjQd5EdPBZv[0] == '.' || '0' <= cjQd5EdPBZv[0] && cjQd5EdPBZv[0] <= '9') {
			return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
		}

		ynS4iTZ :=  /*line :1*/ len(cjQd5EdPBZv)
		ha9RaMigT, cjQd5EdPBZv, xuMLYrB =  /*line :1*/ wQR8Ii5R(cjQd5EdPBZv)
		if xuMLYrB != nil {
			return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
		}
		bKOg7ZHd20cR := ynS4iTZ !=  /*line :1*/ len(cjQd5EdPBZv)

		uq6Llhwrg := false
		if cjQd5EdPBZv != "" && cjQd5EdPBZv[0] == '.' {
			cjQd5EdPBZv = cjQd5EdPBZv[1:]
			ynS4iTZ :=  /*line :1*/ len(cjQd5EdPBZv)
			k05jcfKvAvw6, eIK3JAH6, cjQd5EdPBZv =  /*line :1*/ fcgzbefXa(cjQd5EdPBZv)
			uq6Llhwrg = ynS4iTZ !=  /*line :1*/ len(cjQd5EdPBZv)
		}
		if !bKOg7ZHd20cR && !uq6Llhwrg {

			return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
		}

		eUo6AxKM7 := 0
		for ; eUo6AxKM7 <  /*line :1*/ len(cjQd5EdPBZv); eUo6AxKM7++ {
			fnpIad3KB := cjQd5EdPBZv[eUo6AxKM7]
			if fnpIad3KB == '.' || '0' <= fnpIad3KB && fnpIad3KB <= '9' {
				break
			}
		}
		if eUo6AxKM7 == 0 {
			return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: missing unit in duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
		}
		b6XtDzh := cjQd5EdPBZv[:eUo6AxKM7]
		cjQd5EdPBZv = cjQd5EdPBZv[eUo6AxKM7:]
		gp9pOOZRNaU, ufWfJe3qPw := r3yCdn63wjx[b6XtDzh]
		if !ufWfJe3qPw {
			return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: unknown unit " +  /*line :1*/ cJ8pucH5Vq(b6XtDzh) + " in duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
		}
		if ha9RaMigT > 1<<63/gp9pOOZRNaU {

			return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
		}
		ha9RaMigT *= gp9pOOZRNaU
		if k05jcfKvAvw6 > 0 {

			ha9RaMigT +=  /*line :1*/ uint64( /*line :1*/ float64(k05jcfKvAvw6) * ( /*line :1*/ float64(gp9pOOZRNaU) / eIK3JAH6))
			if ha9RaMigT > 1<<63 {

				return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
			}
		}
		lNAwkbJQVBgU += ha9RaMigT
		if lNAwkbJQVBgU > 1<<63 {
			return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
		}
	}
	if yY3pgXWox3j {
		return - /*line :1*/ GKMwTGxQa0S(lNAwkbJQVBgU), nil
	}
	if lNAwkbJQVBgU > 1<<63-1 {
		return 0,  /*line :1*/ errors.Su6g6hRoi9X("time: invalid duration " +  /*line :1*/ cJ8pucH5Vq(bVPj4pI))
	}
	return  /*line :1*/ GKMwTGxQa0S(lNAwkbJQVBgU), nil
}
