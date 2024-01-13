//line :1
package JkjxVS


func V25ba9G5(lljq0sSzGHZ int) string {
	if lljq0sSzGHZ < 0 {
		return "-" +  /*line :1*/ TDADPeM7Kft9( /*line :1*/ uint(-lljq0sSzGHZ))
	}
	return  /*line :1*/ TDADPeM7Kft9( /*line :1*/ uint(lljq0sSzGHZ))
}


func TDADPeM7Kft9(lljq0sSzGHZ uint) string {
	if lljq0sSzGHZ == 0 {
		return "0"
	}
	var zwkG89I5 [20]byte	
	xSgZyIY1 :=  /*line :1*/ len(zwkG89I5) - 1
	for lljq0sSzGHZ >= 10 {
		bKJLHrF := lljq0sSzGHZ / 10
		zwkG89I5[xSgZyIY1] =  /*line :1*/ byte('0' + lljq0sSzGHZ - bKJLHrF*10)
		xSgZyIY1--
		lljq0sSzGHZ = bKJLHrF
	}

	zwkG89I5[xSgZyIY1] =  /*line :1*/ byte('0' + lljq0sSzGHZ)
	return  /*line :1*/ string(zwkG89I5[xSgZyIY1:])
}
