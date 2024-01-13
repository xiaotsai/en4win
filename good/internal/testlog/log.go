//line :1



package X0MraxO81Me

import atomic "sync/atomic"





type Cdg8GzJqHj interface {
	Getenv(dW0qnxLY string)
	Stat(yoaGcA7a6 string)
	Open(yoaGcA7a6 string)
	Chdir(nmfbi2C2I string)
}







var itG4dyuKh8RW atomic.ANnCMU



func H9fDqdrpc(z1qhtPZPYr Cdg8GzJqHj) {
	if  /*line :1*/ itG4dyuKh8RW.Load() != nil {
		 /*line :1*/ panic("testlog: SetLogger must be called only once")
	}
	 /*line :1*/ itG4dyuKh8RW.Store(&z1qhtPZPYr)
}



func TDkzeK() Cdg8GzJqHj {
	z1qhtPZPYr :=  /*line :1*/ itG4dyuKh8RW.Load()
	if z1qhtPZPYr == nil {
		return nil
	}
	return *z1qhtPZPYr.(*Cdg8GzJqHj)
}


func HUDK1vRBN(voQXN_P string) {
	if iJaeCloD :=  /*line :1*/ TDkzeK(); iJaeCloD != nil {
		 /*line :1*/ iJaeCloD.Getenv(voQXN_P)
	}
}


func NT3Zg6plraKK(voQXN_P string) {
	if iJaeCloD :=  /*line :1*/ TDkzeK(); iJaeCloD != nil {
		 /*line :1*/ iJaeCloD.Open(voQXN_P)
	}
}


func UNQbSYmyL8(voQXN_P string) {
	if iJaeCloD :=  /*line :1*/ TDkzeK(); iJaeCloD != nil {
		 /*line :1*/ iJaeCloD.Stat(voQXN_P)
	}
}
