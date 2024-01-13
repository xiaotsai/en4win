//line :1
package math








func HJa8famXXqZ(ygoarCBO4et float64) (f50wU6B6G float64, cmLS3KfTv float64) {
	if haveArchModf {
		return  /*line :1*/ htMrKoG0ur(ygoarCBO4et)
	}
	return  /*line :1*/ vT82_C(ygoarCBO4et)
}

func vT82_C(ygoarCBO4et float64) (f50wU6B6G float64, cmLS3KfTv float64) {
	if ygoarCBO4et < 1 {
		switch {
		case ygoarCBO4et < 0:
			f50wU6B6G, cmLS3KfTv =  /*line :1*/ HJa8famXXqZ(-ygoarCBO4et)
			return -f50wU6B6G, -cmLS3KfTv
		case ygoarCBO4et == 0:
			return ygoarCBO4et, ygoarCBO4et
		}
		return 0, ygoarCBO4et
	}

	_BqkSz0 :=  /*line :1*/ GKIRmoHE(ygoarCBO4et)
	gwBaYgYgA :=  /*line :1*/ uint(_BqkSz0>>shift)&mask - bias

	if gwBaYgYgA < 64-12 {
		_BqkSz0 &^= 1<<(64-12-gwBaYgYgA) - 1
	}
	f50wU6B6G =  /*line :1*/ QUB2Kf63p(_BqkSz0)
	cmLS3KfTv = ygoarCBO4et - f50wU6B6G
	return
}
