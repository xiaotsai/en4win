//line :1
package wLlfRPv

import (
	errors "iAHoxjmM"
	io "xI9ai2djaJ2"
	utf8 "CuAc0pPZwf"
)


const smallBufferSize = 64



type SuzhTyj35Jl struct {
	y4yIORF49	[]byte	
	_iwf7Ml	int	
	gwbNUw9OQXi	er1Jx2z	
}





type er1Jx2z int8



const (
	opRead	er1Jx2z	= -1	
	opInvalid	er1Jx2z	= 0	
	opReadRune1	er1Jx2z	= 1	
	opReadRune2	er1Jx2z	= 2	
	opReadRune3	er1Jx2z	= 3	
	opReadRune4	er1Jx2z	= 4	
)


var PD0ApL1BbwYW =  /*line :1*/ errors.Su6g6hRoi9X("bytes.Buffer: too large")
var cGZW89M =  /*line :1*/ errors.Su6g6hRoi9X("bytes.Buffer: reader returned negative count from Read")

const maxInt =  /*line :1*/ int(^ /*line :1*/ uint(0) >> 1)






func (ka9IyYc0 *SuzhTyj35Jl) Bytes() []byte	{ return ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:] }





func (ka9IyYc0 *SuzhTyj35Jl) AvailableBuffer() []byte {
	return ka9IyYc0.y4yIORF49[ /*line :1*/ len(ka9IyYc0.y4yIORF49):]
}





func (ka9IyYc0 *SuzhTyj35Jl) String() string {
	if ka9IyYc0 == nil {

		return "<nil>"
	}
	return  /*line :1*/ string(ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:])
}


func (ka9IyYc0 *SuzhTyj35Jl) bOJJI7Xso() bool	{ return  /*line :1*/ len(ka9IyYc0.y4yIORF49) <= ka9IyYc0._iwf7Ml }



func (ka9IyYc0 *SuzhTyj35Jl) Len() int	{ return  /*line :1*/ len(ka9IyYc0.y4yIORF49) - ka9IyYc0._iwf7Ml }



func (ka9IyYc0 *SuzhTyj35Jl) Cap() int	{ return  /*line :1*/ cap(ka9IyYc0.y4yIORF49) }


func (ka9IyYc0 *SuzhTyj35Jl) Available() int {
	return  /*line :1*/ cap(ka9IyYc0.y4yIORF49) -  /*line :1*/ len(ka9IyYc0.y4yIORF49)
}




func (ka9IyYc0 *SuzhTyj35Jl) Truncate(lQivwVxjomk int) {
	if lQivwVxjomk == 0 {
		 /*line :1*/ ka9IyYc0.Reset()
		return
	}
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	if lQivwVxjomk < 0 || lQivwVxjomk >  /*line :1*/ ka9IyYc0.Len() {
		 /*line :1*/ panic("bytes.Buffer: truncation out of range")
	}
	ka9IyYc0.y4yIORF49 = ka9IyYc0.y4yIORF49[:ka9IyYc0._iwf7Ml+lQivwVxjomk]
}




func (ka9IyYc0 *SuzhTyj35Jl) Reset() {
	ka9IyYc0.y4yIORF49 = ka9IyYc0.y4yIORF49[:0]
	ka9IyYc0._iwf7Ml = 0
	ka9IyYc0.gwbNUw9OQXi = opInvalid
}




func (ka9IyYc0 *SuzhTyj35Jl) gPdy30ry1z3r(lQivwVxjomk int) (int, bool) {
	if eGXSPnra9w :=  /*line :1*/ len(ka9IyYc0.y4yIORF49); lQivwVxjomk <=  /*line :1*/ cap(ka9IyYc0.y4yIORF49)-eGXSPnra9w {
		ka9IyYc0.y4yIORF49 = ka9IyYc0.y4yIORF49[:eGXSPnra9w+lQivwVxjomk]
		return eGXSPnra9w, true
	}
	return 0, false
}




func (ka9IyYc0 *SuzhTyj35Jl) thnQCWMKQ06a(lQivwVxjomk int) int {
	aJCPDy4w :=  /*line :1*/ ka9IyYc0.Len()

	if aJCPDy4w == 0 && ka9IyYc0._iwf7Ml != 0 {
		 /*line :1*/ ka9IyYc0.Reset()
	}

	if musLbUdLeb, lxjoXcF64YF :=  /*line :1*/ ka9IyYc0.gPdy30ry1z3r(lQivwVxjomk); lxjoXcF64YF {
		return musLbUdLeb
	}
	if ka9IyYc0.y4yIORF49 == nil && lQivwVxjomk <= smallBufferSize {
		ka9IyYc0.y4yIORF49 =  /*line :1*/ make([]byte, lQivwVxjomk, smallBufferSize)
		return 0
	}
	wOPKkW :=  /*line :1*/ cap(ka9IyYc0.y4yIORF49)
	if lQivwVxjomk <= wOPKkW/2-aJCPDy4w {

		 /*line :1*/ copy(ka9IyYc0.y4yIORF49, ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:])
	} else if wOPKkW > maxInt-wOPKkW-lQivwVxjomk {
		 /*line :1*/ panic(PD0ApL1BbwYW)
	} else {

		ka9IyYc0.y4yIORF49 =  /*line :1*/ soaQe707O(ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:], ka9IyYc0._iwf7Ml+lQivwVxjomk)
	}

	ka9IyYc0._iwf7Ml = 0
	ka9IyYc0.y4yIORF49 = ka9IyYc0.y4yIORF49[:aJCPDy4w+lQivwVxjomk]
	return aJCPDy4w
}






func (ka9IyYc0 *SuzhTyj35Jl) Grow(lQivwVxjomk int) {
	if lQivwVxjomk < 0 {
		 /*line :1*/ panic("bytes.Buffer.Grow: negative count")
	}
	aJCPDy4w :=  /*line :1*/ ka9IyYc0.thnQCWMKQ06a(lQivwVxjomk)
	ka9IyYc0.y4yIORF49 = ka9IyYc0.y4yIORF49[:aJCPDy4w]
}




func (ka9IyYc0 *SuzhTyj35Jl) Write(gHlF23O57jUB []byte) (lQivwVxjomk int, iRRn6Ng97Jf0 error) {
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	aJCPDy4w, lxjoXcF64YF :=  /*line :1*/ ka9IyYc0.gPdy30ry1z3r( /*line :1*/ len(gHlF23O57jUB))
	if !lxjoXcF64YF {
		aJCPDy4w =  /*line :1*/ ka9IyYc0.thnQCWMKQ06a( /*line :1*/ len(gHlF23O57jUB))
	}
	return  /*line :1*/ copy(ka9IyYc0.y4yIORF49[aJCPDy4w:], gHlF23O57jUB), nil
}




func (ka9IyYc0 *SuzhTyj35Jl) WriteString(nhst_x string) (lQivwVxjomk int, iRRn6Ng97Jf0 error) {
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	aJCPDy4w, lxjoXcF64YF :=  /*line :1*/ ka9IyYc0.gPdy30ry1z3r( /*line :1*/ len(nhst_x))
	if !lxjoXcF64YF {
		aJCPDy4w =  /*line :1*/ ka9IyYc0.thnQCWMKQ06a( /*line :1*/ len(nhst_x))
	}
	return  /*line :1*/ copy(ka9IyYc0.y4yIORF49[aJCPDy4w:], nhst_x), nil
}





const MinRead = 512





func (ka9IyYc0 *SuzhTyj35Jl) ReadFrom(dIwXIFT io.YJ04Fau) (lQivwVxjomk int64, iRRn6Ng97Jf0 error) {
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	for {
		musLbUdLeb :=  /*line :1*/ ka9IyYc0.thnQCWMKQ06a(MinRead)
		ka9IyYc0.y4yIORF49 = ka9IyYc0.y4yIORF49[:musLbUdLeb]
		aJCPDy4w, wCjEZiq :=  /*line :1*/ dIwXIFT.Read(ka9IyYc0.y4yIORF49[musLbUdLeb: /*line :1*/ cap(ka9IyYc0.y4yIORF49)])
		if aJCPDy4w < 0 {
			 /*line :1*/ panic(cGZW89M)
		}

		ka9IyYc0.y4yIORF49 = ka9IyYc0.y4yIORF49[:musLbUdLeb+aJCPDy4w]
		lQivwVxjomk +=  /*line :1*/ int64(aJCPDy4w)
		if wCjEZiq == io.K5Sqco {
			return lQivwVxjomk, nil
		}
		if wCjEZiq != nil {
			return lQivwVxjomk, wCjEZiq
		}
	}
}



func soaQe707O(ka9IyYc0 []byte, lQivwVxjomk int) []byte {
	defer func() {
		if  /*line :1*/ recover() != nil {
			 /*line :1*/ panic(PD0ApL1BbwYW)
		}
	}()

	wOPKkW :=  /*line :1*/ len(ka9IyYc0) + lQivwVxjomk
	if wOPKkW < 2* /*line :1*/ cap(ka9IyYc0) {

		wOPKkW = 2 *  /*line :1*/ cap(ka9IyYc0)
	}
	ndi0nzxNn :=  /*line :1*/ append([] /*line :1*/ byte(nil),  /*line :1*/ make([]byte, wOPKkW)...)
	 /*line :1*/ copy(ndi0nzxNn, ka9IyYc0)
	return ndi0nzxNn[: /*line :1*/ len(ka9IyYc0)]
}





func (ka9IyYc0 *SuzhTyj35Jl) WriteTo(yqP0sy io.LXQrGW6BQt) (lQivwVxjomk int64, iRRn6Ng97Jf0 error) {
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	if aOCQVr1fWrfw :=  /*line :1*/ ka9IyYc0.Len(); aOCQVr1fWrfw > 0 {
		aJCPDy4w, wCjEZiq :=  /*line :1*/ yqP0sy.Write(ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:])
		if aJCPDy4w > aOCQVr1fWrfw {
			 /*line :1*/ panic("bytes.Buffer.WriteTo: invalid Write count")
		}
		ka9IyYc0._iwf7Ml += aJCPDy4w
		lQivwVxjomk =  /*line :1*/ int64(aJCPDy4w)
		if wCjEZiq != nil {
			return lQivwVxjomk, wCjEZiq
		}

		if aJCPDy4w != aOCQVr1fWrfw {
			return lQivwVxjomk, io.ArPWDHfv
		}
	}

	 /*line :1*/ ka9IyYc0.Reset()
	return lQivwVxjomk, nil
}





func (ka9IyYc0 *SuzhTyj35Jl) WriteByte(wOPKkW byte) error {
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	aJCPDy4w, lxjoXcF64YF :=  /*line :1*/ ka9IyYc0.gPdy30ry1z3r(1)
	if !lxjoXcF64YF {
		aJCPDy4w =  /*line :1*/ ka9IyYc0.thnQCWMKQ06a(1)
	}
	ka9IyYc0.y4yIORF49[aJCPDy4w] = wOPKkW
	return nil
}





func (ka9IyYc0 *SuzhTyj35Jl) WriteRune(dIwXIFT rune) (lQivwVxjomk int, iRRn6Ng97Jf0 error) {

	if  /*line :1*/ uint32(dIwXIFT) < utf8.RuneSelf {
		 /*line :1*/ ka9IyYc0.WriteByte( /*line :1*/ byte(dIwXIFT))
		return 1, nil
	}
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	aJCPDy4w, lxjoXcF64YF :=  /*line :1*/ ka9IyYc0.gPdy30ry1z3r(utf8.UTFMax)
	if !lxjoXcF64YF {
		aJCPDy4w =  /*line :1*/ ka9IyYc0.thnQCWMKQ06a(utf8.UTFMax)
	}
	ka9IyYc0.y4yIORF49 =  /*line :1*/ utf8.Ht7oMzd(ka9IyYc0.y4yIORF49[:aJCPDy4w], dIwXIFT)
	return  /*line :1*/ len(ka9IyYc0.y4yIORF49) - aJCPDy4w, nil
}





func (ka9IyYc0 *SuzhTyj35Jl) Read(gHlF23O57jUB []byte) (lQivwVxjomk int, iRRn6Ng97Jf0 error) {
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	if  /*line :1*/ ka9IyYc0.bOJJI7Xso() {

		 /*line :1*/ ka9IyYc0.Reset()
		if  /*line :1*/ len(gHlF23O57jUB) == 0 {
			return 0, nil
		}
		return 0, io.K5Sqco
	}
	lQivwVxjomk =  /*line :1*/ copy(gHlF23O57jUB, ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:])
	ka9IyYc0._iwf7Ml += lQivwVxjomk
	if lQivwVxjomk > 0 {
		ka9IyYc0.gwbNUw9OQXi = opRead
	}
	return lQivwVxjomk, nil
}





func (ka9IyYc0 *SuzhTyj35Jl) Next(lQivwVxjomk int) []byte {
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	aJCPDy4w :=  /*line :1*/ ka9IyYc0.Len()
	if lQivwVxjomk > aJCPDy4w {
		lQivwVxjomk = aJCPDy4w
	}
	e_B0GST8 := ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml : ka9IyYc0._iwf7Ml+lQivwVxjomk]
	ka9IyYc0._iwf7Ml += lQivwVxjomk
	if lQivwVxjomk > 0 {
		ka9IyYc0.gwbNUw9OQXi = opRead
	}
	return e_B0GST8
}



func (ka9IyYc0 *SuzhTyj35Jl) ReadByte() (byte, error) {
	if  /*line :1*/ ka9IyYc0.bOJJI7Xso() {

		 /*line :1*/ ka9IyYc0.Reset()
		return 0, io.K5Sqco
	}
	wOPKkW := ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml]
	ka9IyYc0._iwf7Ml++
	ka9IyYc0.gwbNUw9OQXi = opRead
	return wOPKkW, nil
}






func (ka9IyYc0 *SuzhTyj35Jl) ReadRune() (dIwXIFT rune, kXik_zS1N6 int, iRRn6Ng97Jf0 error) {
	if  /*line :1*/ ka9IyYc0.bOJJI7Xso() {

		 /*line :1*/ ka9IyYc0.Reset()
		return 0, 0, io.K5Sqco
	}
	wOPKkW := ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml]
	if wOPKkW < utf8.RuneSelf {
		ka9IyYc0._iwf7Ml++
		ka9IyYc0.gwbNUw9OQXi = opReadRune1
		return  /*line :1*/ rune(wOPKkW), 1, nil
	}
	dIwXIFT, lQivwVxjomk :=  /*line :1*/ utf8.EicfpCPn(ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:])
	ka9IyYc0._iwf7Ml += lQivwVxjomk
	ka9IyYc0.gwbNUw9OQXi =  /*line :1*/ er1Jx2z(lQivwVxjomk)
	return dIwXIFT, lQivwVxjomk, nil
}






func (ka9IyYc0 *SuzhTyj35Jl) UnreadRune() error {
	if ka9IyYc0.gwbNUw9OQXi <= opInvalid {
		return  /*line :1*/ errors.Su6g6hRoi9X("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
	}
	if ka9IyYc0._iwf7Ml >=  /*line :1*/ int(ka9IyYc0.gwbNUw9OQXi) {
		ka9IyYc0._iwf7Ml -=  /*line :1*/ int(ka9IyYc0.gwbNUw9OQXi)
	}
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	return nil
}

var j2X6qRJZlE =  /*line :1*/ errors.Su6g6hRoi9X("bytes.Buffer: UnreadByte: previous operation was not a successful read")





func (ka9IyYc0 *SuzhTyj35Jl) UnreadByte() error {
	if ka9IyYc0.gwbNUw9OQXi == opInvalid {
		return j2X6qRJZlE
	}
	ka9IyYc0.gwbNUw9OQXi = opInvalid
	if ka9IyYc0._iwf7Ml > 0 {
		ka9IyYc0._iwf7Ml--
	}
	return nil
}







func (ka9IyYc0 *SuzhTyj35Jl) ReadBytes(rSDLTVdF byte) (vRkt0QF7 []byte, iRRn6Ng97Jf0 error) {
	b7R_hv, iRRn6Ng97Jf0 :=  /*line :1*/ ka9IyYc0.upmjDTKCO3cQ(rSDLTVdF)

	vRkt0QF7 =  /*line :1*/ append(vRkt0QF7, b7R_hv...)
	return vRkt0QF7, iRRn6Ng97Jf0
}


func (ka9IyYc0 *SuzhTyj35Jl) upmjDTKCO3cQ(rSDLTVdF byte) (vRkt0QF7 []byte, iRRn6Ng97Jf0 error) {
	musLbUdLeb :=  /*line :1*/ AXZRWQNqa(ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:], rSDLTVdF)
	gJGWLx4tJAaJ := ka9IyYc0._iwf7Ml + musLbUdLeb + 1
	if musLbUdLeb < 0 {
		gJGWLx4tJAaJ =  /*line :1*/ len(ka9IyYc0.y4yIORF49)
		iRRn6Ng97Jf0 = io.K5Sqco
	}
	vRkt0QF7 = ka9IyYc0.y4yIORF49[ka9IyYc0._iwf7Ml:gJGWLx4tJAaJ]
	ka9IyYc0._iwf7Ml = gJGWLx4tJAaJ
	ka9IyYc0.gwbNUw9OQXi = opRead
	return vRkt0QF7, iRRn6Ng97Jf0
}







func (ka9IyYc0 *SuzhTyj35Jl) ReadString(rSDLTVdF byte) (vRkt0QF7 string, iRRn6Ng97Jf0 error) {
	b7R_hv, iRRn6Ng97Jf0 :=  /*line :1*/ ka9IyYc0.upmjDTKCO3cQ(rSDLTVdF)
	return  /*line :1*/ string(b7R_hv), iRRn6Ng97Jf0
}










func FA85RwEszL(iVZXgI_ []byte) *SuzhTyj35Jl	{ return &SuzhTyj35Jl{y4yIORF49: iVZXgI_} }







func ECghF9HU1(nhst_x string) *SuzhTyj35Jl {
	return &SuzhTyj35Jl{y4yIORF49: [] /*line :1*/ byte(nhst_x)}
}
