//line :1

















































package fOYpb3F03EG

import (
	errors "iAHoxjmM"
	reflectlite "yAj8QghzkR"
	sync "sync"
	atomic "sync/atomic"
	time "fRAfQd_"
)





type MBCyA6 interface {
	
	
	
	Deadline() (rW8exMXQmr time.G4KDkI, woPxhOn0LtY bool)

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Done() <-chan struct{}

	
	
	
	
	
	Err() error

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Value(ciwe0CO72 any) any
}


var H0hTZQVavFh =  /*line :1*/ errors.Su6g6hRoi9X("context canceled")



var Z201U2N error = b0zkJtNJnkd{}

type b0zkJtNJnkd struct{}

func (b0zkJtNJnkd) Error() string	{ return "context deadline exceeded" }
func (b0zkJtNJnkd) Timeout() bool	{ return true }
func (b0zkJtNJnkd) Temporary() bool	{ return true }



type eShpXONMmYzX struct{}

func (eShpXONMmYzX) Deadline() (rW8exMXQmr time.G4KDkI, woPxhOn0LtY bool) {
	return
}

func (eShpXONMmYzX) Done() <-chan struct{} {
	return nil
}

func (eShpXONMmYzX) Err() error {
	return nil
}

func (eShpXONMmYzX) Value(ciwe0CO72 any) any {
	return nil
}

type xMjFQNoY59v struct{ eShpXONMmYzX }

func (xMjFQNoY59v) String() string {
	return "context.Background"
}

type _q34qtSe struct{ eShpXONMmYzX }

func (_q34qtSe) String() string {
	return "context.TODO"
}





func GEcgQP5fzA() MBCyA6 {
	return xMjFQNoY59v{}
}





func GQNg8Ka() MBCyA6 {
	return _q34qtSe{}
}





type Gd7QRmF func()







func Cu9yRv(hwsX1dTj5 MBCyA6) (vw3nV4md MBCyA6, v2GjsOKv4c Gd7QRmF) {
	d99B2B :=  /*line :1*/ a5JUpH2cEH(hwsX1dTj5)
	return d99B2B, func() {  /*line :1*/ d99B2B.v2GjsOKv4c(true, H0hTZQVavFh, nil) }
}











type MnpcubObjLT func(vq3q6daDLq error)












func AgAdcm8(hwsX1dTj5 MBCyA6) (vw3nV4md MBCyA6, v2GjsOKv4c MnpcubObjLT) {
	d99B2B :=  /*line :1*/ a5JUpH2cEH(hwsX1dTj5)
	return d99B2B, func(vq3q6daDLq error) {  /*line :1*/ d99B2B.v2GjsOKv4c(true, H0hTZQVavFh, vq3q6daDLq) }
}

func a5JUpH2cEH(hwsX1dTj5 MBCyA6) *eDRXx5uurBxj {
	if hwsX1dTj5 == nil {
		 /*line :1*/ panic("cannot create context from nil parent")
	}
	d99B2B := &eDRXx5uurBxj{}
	 /*line :1*/ d99B2B.bZgcQyDq4aUX(hwsX1dTj5, d99B2B)
	return d99B2B
}







func YOYaVTEoQkW7(d99B2B MBCyA6) error {
	if er6zJzGLTMc7, woPxhOn0LtY :=  /*line :1*/ d99B2B.Value(&zu0eFGs).(*eDRXx5uurBxj); woPxhOn0LtY {
		 /*line :1*/ er6zJzGLTMc7.zaTgDKN0OO.Lock()
		defer  /*line :1*/ er6zJzGLTMc7.zaTgDKN0OO.Unlock()
		return er6zJzGLTMc7.xmD0SF
	}
	return nil
}



















func IRJ982uI(vw3nV4md MBCyA6, masrQa3Y3 func()) (bD5mjaPAFf func() bool) {
	kJBFhe := &j1jtFULq0nR0{
		r2WJRK1Ap: masrQa3Y3,
	}
	 /*line :1*/ kJBFhe.eDRXx5uurBxj.bZgcQyDq4aUX(vw3nV4md, kJBFhe)
	return func() bool {
		cKIva7d8 := false
		 /*line :1*/ kJBFhe.i1y7QJLemz4.Do(func() {
			cKIva7d8 = true
		})
		if cKIva7d8 {
			 /*line :1*/ kJBFhe.v2GjsOKv4c(true, H0hTZQVavFh, nil)
		}
		return cKIva7d8
	}
}

type iipR1k interface {
	AfterFunc(func()) func() bool
}

type j1jtFULq0nR0 struct {
	eDRXx5uurBxj
	i1y7QJLemz4	sync.LhBfwn6wa1x	
	r2WJRK1Ap	func()
}

func (kJBFhe *j1jtFULq0nR0) v2GjsOKv4c(s3ImKK6fHIz bool, l8QOrgWaogb4, vq3q6daDLq error) {
	 /*line :1*/ kJBFhe.eDRXx5uurBxj.v2GjsOKv4c(false, l8QOrgWaogb4, vq3q6daDLq)
	if s3ImKK6fHIz {
		 /*line :1*/ gWrvZbwI(kJBFhe.MBCyA6, kJBFhe)
	}
	 /*line :1*/ kJBFhe.i1y7QJLemz4.Do(func() {
		go  /*line :1*/ kJBFhe.r2WJRK1Ap()
	})
}




type wByREwr struct {
	MBCyA6
	zcaDO_Gyh	func() bool
}


var qAN81MPyZk atomic.JhtCwKEzC


var zu0eFGs int







func yIDe8NQk1km(hwsX1dTj5 MBCyA6) (*eDRXx5uurBxj, bool) {
	azrPX2RcTA :=  /*line :1*/ hwsX1dTj5.Done()
	if azrPX2RcTA == naCPw__ || azrPX2RcTA == nil {
		return nil, false
	}
	avrcaac, woPxhOn0LtY :=  /*line :1*/ hwsX1dTj5.Value(&zu0eFGs).(*eDRXx5uurBxj)
	if !woPxhOn0LtY {
		return nil, false
	}
	tZ1PmsfdCcA, _ :=  /*line :1*/ avrcaac.nGDwOdnA.Load().(chan struct{})
	if tZ1PmsfdCcA != azrPX2RcTA {
		return nil, false
	}
	return avrcaac, true
}


func gWrvZbwI(hwsX1dTj5 MBCyA6, axinGEW_Aw q96KubNpYFc) {
	if wLD8ShH, woPxhOn0LtY := hwsX1dTj5.(wByREwr); woPxhOn0LtY {
		 /*line :1*/ wLD8ShH.zcaDO_Gyh()
		return
	}
	avrcaac, woPxhOn0LtY :=  /*line :1*/ yIDe8NQk1km(hwsX1dTj5)
	if !woPxhOn0LtY {
		return
	}
	 /*line :1*/ avrcaac.zaTgDKN0OO.Lock()
	if avrcaac.dn83fTYtzUu != nil {
		 /*line :1*/ delete(avrcaac.dn83fTYtzUu, axinGEW_Aw)
	}
	 /*line :1*/ avrcaac.zaTgDKN0OO.Unlock()
}



type q96KubNpYFc interface {
	v2GjsOKv4c(s3ImKK6fHIz bool, l8QOrgWaogb4, vq3q6daDLq error)
	Done() <-chan struct{}
}


var naCPw__ =  /*line :1*/ make(chan struct{})

func init() {
	 /*line :1*/ close(naCPw__)
}



type eDRXx5uurBxj struct {
	MBCyA6

	zaTgDKN0OO	sync.DIRWe11KYlYa	
	nGDwOdnA	atomic.ANnCMU	
	dn83fTYtzUu	map[q96KubNpYFc]struct{}	
	ohNSUtFoy	error	
	xmD0SF	error	
}

func (d99B2B *eDRXx5uurBxj) Value(ciwe0CO72 any) any {
	if ciwe0CO72 == &zu0eFGs {
		return d99B2B
	}
	return  /*line :1*/ kG6Spc(d99B2B.MBCyA6, ciwe0CO72)
}

func (d99B2B *eDRXx5uurBxj) Done() <-chan struct{} {
	ox4E1tAfNu :=  /*line :1*/ d99B2B.nGDwOdnA.Load()
	if ox4E1tAfNu != nil {
		return ox4E1tAfNu.(chan struct{})
	}
	 /*line :1*/ d99B2B.zaTgDKN0OO.Lock()
	defer  /*line :1*/ d99B2B.zaTgDKN0OO.Unlock()
	ox4E1tAfNu =  /*line :1*/ d99B2B.nGDwOdnA.Load()
	if ox4E1tAfNu == nil {
		ox4E1tAfNu =  /*line :1*/ make(chan struct{})
		 /*line :1*/ d99B2B.nGDwOdnA.Store(ox4E1tAfNu)
	}
	return ox4E1tAfNu.(chan struct{})
}

func (d99B2B *eDRXx5uurBxj) Err() error {
	 /*line :1*/ d99B2B.zaTgDKN0OO.Lock()
	l8QOrgWaogb4 := d99B2B.ohNSUtFoy
	 /*line :1*/ d99B2B.zaTgDKN0OO.Unlock()
	return l8QOrgWaogb4
}



func (d99B2B *eDRXx5uurBxj) bZgcQyDq4aUX(hwsX1dTj5 MBCyA6, axinGEW_Aw q96KubNpYFc) {
	d99B2B.MBCyA6 = hwsX1dTj5

	azrPX2RcTA :=  /*line :1*/ hwsX1dTj5.Done()
	if azrPX2RcTA == nil {
		return
	}

	select {
	case <-azrPX2RcTA:

		 /*line :1*/ axinGEW_Aw.v2GjsOKv4c(false,  /*line :1*/ hwsX1dTj5.Err(),  /*line :1*/ YOYaVTEoQkW7(hwsX1dTj5))
		return
	default:
	}

	if avrcaac, woPxhOn0LtY :=  /*line :1*/ yIDe8NQk1km(hwsX1dTj5); woPxhOn0LtY {

		 /*line :1*/ avrcaac.zaTgDKN0OO.Lock()
		if avrcaac.ohNSUtFoy != nil {

			 /*line :1*/ axinGEW_Aw.v2GjsOKv4c(false, avrcaac.ohNSUtFoy, avrcaac.xmD0SF)
		} else {
			if avrcaac.dn83fTYtzUu == nil {
				avrcaac.dn83fTYtzUu =  /*line :1*/ make(map[q96KubNpYFc]struct{})
			}
			avrcaac.dn83fTYtzUu[axinGEW_Aw] = struct{}{}
		}
		 /*line :1*/ avrcaac.zaTgDKN0OO.Unlock()
		return
	}

	if kJBFhe, woPxhOn0LtY := hwsX1dTj5.(iipR1k); woPxhOn0LtY {

		 /*line :1*/ d99B2B.zaTgDKN0OO.Lock()
		bD5mjaPAFf :=  /*line :1*/ kJBFhe.AfterFunc(func() {
			 /*line :1*/ axinGEW_Aw.v2GjsOKv4c(false,  /*line :1*/ hwsX1dTj5.Err(),  /*line :1*/ YOYaVTEoQkW7(hwsX1dTj5))
		})
		d99B2B.MBCyA6 = wByREwr{
			MBCyA6:	hwsX1dTj5,
			zcaDO_Gyh:	bD5mjaPAFf,
		}
		 /*line :1*/ d99B2B.zaTgDKN0OO.Unlock()
		return
	}

	 /*line :1*/ qAN81MPyZk.Add(1)
	go func() {
		select {
		case <- /*line :1*/ hwsX1dTj5.Done():
			 /*line :1*/ axinGEW_Aw.v2GjsOKv4c(false,  /*line :1*/ hwsX1dTj5.Err(),  /*line :1*/ YOYaVTEoQkW7(hwsX1dTj5))
		case <- /*line :1*/ axinGEW_Aw.Done():
		}
	}()
}

type v19VJHh interface {
	String() string
}

func drN0pE(d99B2B MBCyA6) string {
	if wLD8ShH, woPxhOn0LtY := d99B2B.(v19VJHh); woPxhOn0LtY {
		return  /*line :1*/ wLD8ShH.String()
	}
	return  /*line :1*/ reflectlite.C08auGx(d99B2B).String()
}

func (d99B2B *eDRXx5uurBxj) String() string {
	return  /*line :1*/ drN0pE(d99B2B.MBCyA6) + ".WithCancel"
}




func (d99B2B *eDRXx5uurBxj) v2GjsOKv4c(s3ImKK6fHIz bool, l8QOrgWaogb4, vq3q6daDLq error) {
	if l8QOrgWaogb4 == nil {
		 /*line :1*/ panic("context: internal error: missing cancel error")
	}
	if vq3q6daDLq == nil {
		vq3q6daDLq = l8QOrgWaogb4
	}
	 /*line :1*/ d99B2B.zaTgDKN0OO.Lock()
	if d99B2B.ohNSUtFoy != nil {
		 /*line :1*/ d99B2B.zaTgDKN0OO.Unlock()
		return
	}
	d99B2B.ohNSUtFoy = l8QOrgWaogb4
	d99B2B.xmD0SF = vq3q6daDLq
	ox4E1tAfNu, _ :=  /*line :1*/ d99B2B.nGDwOdnA.Load().(chan struct{})
	if ox4E1tAfNu == nil {
		 /*line :1*/ d99B2B.nGDwOdnA.Store(naCPw__)
	} else {
		 /*line :1*/ close(ox4E1tAfNu)
	}
	for axinGEW_Aw := range d99B2B.dn83fTYtzUu {

		 /*line :1*/ axinGEW_Aw.v2GjsOKv4c(false, l8QOrgWaogb4, vq3q6daDLq)
	}
	d99B2B.dn83fTYtzUu = nil
	 /*line :1*/ d99B2B.zaTgDKN0OO.Unlock()

	if s3ImKK6fHIz {
		 /*line :1*/ gWrvZbwI(d99B2B.MBCyA6, d99B2B)
	}
}




func Ymp4RZewym1J(hwsX1dTj5 MBCyA6) MBCyA6 {
	if hwsX1dTj5 == nil {
		 /*line :1*/ panic("cannot create context from nil parent")
	}
	return aS3MeV6dQxgt{hwsX1dTj5}
}

type aS3MeV6dQxgt struct {
	jZC7vD9a5 MBCyA6
}

func (aS3MeV6dQxgt) Deadline() (rW8exMXQmr time.G4KDkI, woPxhOn0LtY bool) {
	return
}

func (aS3MeV6dQxgt) Done() <-chan struct{} {
	return nil
}

func (aS3MeV6dQxgt) Err() error {
	return nil
}

func (d99B2B aS3MeV6dQxgt) Value(ciwe0CO72 any) any {
	return  /*line :1*/ kG6Spc(d99B2B, ciwe0CO72)
}

func (d99B2B aS3MeV6dQxgt) String() string {
	return  /*line :1*/ drN0pE(d99B2B.jZC7vD9a5) + ".WithoutCancel"
}










func ZF8aTuau(hwsX1dTj5 MBCyA6, ox4E1tAfNu time.G4KDkI) (MBCyA6, Gd7QRmF) {
	return  /*line :1*/ NHMHtXSf1A(hwsX1dTj5, ox4E1tAfNu, nil)
}




func NHMHtXSf1A(hwsX1dTj5 MBCyA6, ox4E1tAfNu time.G4KDkI, vq3q6daDLq error) (MBCyA6, Gd7QRmF) {
	if hwsX1dTj5 == nil {
		 /*line :1*/ panic("cannot create context from nil parent")
	}
	if gwAqX_i, woPxhOn0LtY :=  /*line :1*/ hwsX1dTj5.Deadline(); woPxhOn0LtY &&  /*line :1*/ gwAqX_i.Before(ox4E1tAfNu) {

		return  /*line :1*/ Cu9yRv(hwsX1dTj5)
	}
	d99B2B := &pUqTUw71kHt{
		hF3NgRaoIRo: ox4E1tAfNu,
	}
	 /*line :1*/ d99B2B.eDRXx5uurBxj.bZgcQyDq4aUX(hwsX1dTj5, d99B2B)
	l0jJtFiY :=  /*line :1*/ time.Qr_kn6Dzra(ox4E1tAfNu)
	if l0jJtFiY <= 0 {
		 /*line :1*/ d99B2B.v2GjsOKv4c(true, Z201U2N, vq3q6daDLq)
		return d99B2B, func() {  /*line :1*/ d99B2B.v2GjsOKv4c(false, H0hTZQVavFh, nil) }
	}
	 /*line :1*/ d99B2B.zaTgDKN0OO.Lock()
	defer  /*line :1*/ d99B2B.zaTgDKN0OO.Unlock()
	if d99B2B.ohNSUtFoy == nil {
		d99B2B.j9OOdA9c9f =  /*line :1*/ time.RzWe24uxMiv(l0jJtFiY, func() {
			 /*line :1*/ d99B2B.v2GjsOKv4c(true, Z201U2N, vq3q6daDLq)
		})
	}
	return d99B2B, func() {  /*line :1*/ d99B2B.v2GjsOKv4c(true, H0hTZQVavFh, nil) }
}




type pUqTUw71kHt struct {
	eDRXx5uurBxj
	j9OOdA9c9f	*time.Bkz_9qnV	

	hF3NgRaoIRo	time.G4KDkI
}

func (d99B2B *pUqTUw71kHt) Deadline() (rW8exMXQmr time.G4KDkI, woPxhOn0LtY bool) {
	return d99B2B.hF3NgRaoIRo, true
}

func (d99B2B *pUqTUw71kHt) String() string {
	return  /*line :1*/ drN0pE(d99B2B.eDRXx5uurBxj.MBCyA6) + ".WithDeadline(" +
		 /*line :1*/ d99B2B.hF3NgRaoIRo.String() + " [" +
		 /*line :1*/ time.Qr_kn6Dzra(d99B2B.hF3NgRaoIRo).String() + "])"
}

func (d99B2B *pUqTUw71kHt) v2GjsOKv4c(s3ImKK6fHIz bool, l8QOrgWaogb4, vq3q6daDLq error) {
	 /*line :1*/ d99B2B.eDRXx5uurBxj.v2GjsOKv4c(false, l8QOrgWaogb4, vq3q6daDLq)
	if s3ImKK6fHIz {

		 /*line :1*/ gWrvZbwI(d99B2B.eDRXx5uurBxj.MBCyA6, d99B2B)
	}
	 /*line :1*/ d99B2B.zaTgDKN0OO.Lock()
	if d99B2B.j9OOdA9c9f != nil {
		 /*line :1*/ d99B2B.j9OOdA9c9f.Stop()
		d99B2B.j9OOdA9c9f = nil
	}
	 /*line :1*/ d99B2B.zaTgDKN0OO.Unlock()
}











func T_vBAj8Epxd(hwsX1dTj5 MBCyA6, fp4aOh time.GKMwTGxQa0S) (MBCyA6, Gd7QRmF) {
	return  /*line :1*/ ZF8aTuau(hwsX1dTj5,  /*line :1*/ time.Pc_35oTY().Add(fp4aOh))
}




func Msnoo2et(hwsX1dTj5 MBCyA6, fp4aOh time.GKMwTGxQa0S, vq3q6daDLq error) (MBCyA6, Gd7QRmF) {
	return  /*line :1*/ NHMHtXSf1A(hwsX1dTj5,  /*line :1*/ time.Pc_35oTY().Add(fp4aOh), vq3q6daDLq)
}














func Z2CaDb(hwsX1dTj5 MBCyA6, ciwe0CO72, tEQARJzV any) MBCyA6 {
	if hwsX1dTj5 == nil {
		 /*line :1*/ panic("cannot create context from nil parent")
	}
	if ciwe0CO72 == nil {
		 /*line :1*/ panic("nil key")
	}
	if ! /*line :1*/ reflectlite.C08auGx(ciwe0CO72).Comparable() {
		 /*line :1*/ panic("key is not comparable")
	}
	return &hLpHTt{hwsX1dTj5, ciwe0CO72, tEQARJzV}
}



type hLpHTt struct {
	MBCyA6
	i6EdqN5nI9VJ, h7j8l8	any
}




func uHLdJuWWZPx(veiXBOTZR0nO any) string {
	switch wLD8ShH := veiXBOTZR0nO.(type) {
	case v19VJHh:
		return  /*line :1*/ wLD8ShH.String()
	case string:
		return wLD8ShH
	}
	return "<not Stringer>"
}

func (d99B2B *hLpHTt) String() string {
	return  /*line :1*/ drN0pE(d99B2B.MBCyA6) + ".WithValue(type " +
		 /*line :1*/ reflectlite.C08auGx(d99B2B.i6EdqN5nI9VJ).String() +
		", val " +  /*line :1*/ uHLdJuWWZPx(d99B2B.h7j8l8) + ")"
}

func (d99B2B *hLpHTt) Value(ciwe0CO72 any) any {
	if d99B2B.i6EdqN5nI9VJ == ciwe0CO72 {
		return d99B2B.h7j8l8
	}
	return  /*line :1*/ kG6Spc(d99B2B.MBCyA6, ciwe0CO72)
}

func kG6Spc(d99B2B MBCyA6, ciwe0CO72 any) any {
	for {
		switch vw3nV4md := d99B2B.(type) {
		case *hLpHTt:
			if ciwe0CO72 == vw3nV4md.i6EdqN5nI9VJ {
				return vw3nV4md.h7j8l8
			}
			d99B2B = vw3nV4md.MBCyA6
		case *eDRXx5uurBxj:
			if ciwe0CO72 == &zu0eFGs {
				return d99B2B
			}
			d99B2B = vw3nV4md.MBCyA6
		case aS3MeV6dQxgt:
			if ciwe0CO72 == &zu0eFGs {

				return nil
			}
			d99B2B = vw3nV4md.jZC7vD9a5
		case *pUqTUw71kHt:
			if ciwe0CO72 == &zu0eFGs {
				return &vw3nV4md.eDRXx5uurBxj
			}
			d99B2B = vw3nV4md.MBCyA6
		case xMjFQNoY59v, _q34qtSe:
			return nil
		default:
			return  /*line :1*/ d99B2B.Value(ciwe0CO72)
		}
	}
}
