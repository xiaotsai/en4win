//line :1
//go:build unix || windows

package gF9mZs2

import (
	"internal/bytealg"
	"runtime"
	syscall "bUKeamF"
)

func j7bhnpSMpen(fIadEXIim6l bool) int {
	if fIadEXIim6l {
		return 1
	}
	return 0
}

func g_gXiFVVD(kQ8_UEhxU GNraIvYhB) (*U_Uc9la, error) {
	uotBHWl8c47, h_ljk48Bm :=  /*line :1*/ WNSbt0MQgc()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	for _, dvFmDYmQz := range uotBHWl8c47 {
		zDd8fqQGg, h_ljk48Bm :=  /*line :1*/ dvFmDYmQz.Addrs()
		if h_ljk48Bm != nil {
			return nil, h_ljk48Bm
		}
		for _, bg70ff := range zDd8fqQGg {
			switch ljsCSk := bg70ff.(type) {
			case *FZJphYv9hV:
				if  /*line :1*/ kQ8_UEhxU.Equal(ljsCSk.GdouSNq80bRw) {
					return &dvFmDYmQz, nil
				}
			case *M6TN4rba7:
				if  /*line :1*/ kQ8_UEhxU.Equal(ljsCSk.RJ6BaPgcU) {
					return &dvFmDYmQz, nil
				}
			}
		}
	}
	if  /*line :1*/ kQ8_UEhxU.Equal(WHSMKoapR) {
		return nil, nil
	}
	return nil, fa1xDtq0
}

func lxWp3e(dvFmDYmQz *U_Uc9la) (GNraIvYhB, error) {
	if dvFmDYmQz == nil {
		return WHSMKoapR, nil
	}
	zDd8fqQGg, h_ljk48Bm :=  /*line :1*/ dvFmDYmQz.Addrs()
	if h_ljk48Bm != nil {
		return nil, h_ljk48Bm
	}
	for _, bg70ff := range zDd8fqQGg {
		switch ljsCSk := bg70ff.(type) {
		case *FZJphYv9hV:
			if  /*line :1*/ ljsCSk.GdouSNq80bRw.To4() != nil {
				return ljsCSk.GdouSNq80bRw, nil
			}
		case *M6TN4rba7:
			if  /*line :1*/ ljsCSk.RJ6BaPgcU.To4() != nil {
				return ljsCSk.RJ6BaPgcU, nil
			}
		}
	}
	return nil, fa1xDtq0
}

func bJqj_8(vP562Gd *syscall.W1r5YRtipLa, dvFmDYmQz *U_Uc9la) error {
	if dvFmDYmQz == nil {
		return nil
	}
	zDd8fqQGg, h_ljk48Bm :=  /*line :1*/ dvFmDYmQz.Addrs()
	if h_ljk48Bm != nil {
		return h_ljk48Bm
	}
	for _, bg70ff := range zDd8fqQGg {
		switch ljsCSk := bg70ff.(type) {
		case *FZJphYv9hV:
			if uI7LZDHm6 :=  /*line :1*/ ljsCSk.GdouSNq80bRw.To4(); uI7LZDHm6 != nil {
				 /*line :1*/ copy(vP562Gd.Nompo3TK[:], uI7LZDHm6)
				goto done
			}
		case *M6TN4rba7:
			if uI7LZDHm6 :=  /*line :1*/ ljsCSk.RJ6BaPgcU.To4(); uI7LZDHm6 != nil {
				 /*line :1*/ copy(vP562Gd.Nompo3TK[:], uI7LZDHm6)
				goto done
			}
		}
	}
done:
	if  /*line :1*/ bytealg.Equal(vP562Gd.P8O2NF0WA[:],  /*line :1*/ WHSMKoapR.To4()) {
		return k7ow5kuh
	}
	return nil
}

func jWFpXum(vyaiiPXm6k_W *jmJczkl1, c2dmkE6N4N int) error {
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_RCVBUF, c2dmkE6N4N)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}

func l6G75dG(vyaiiPXm6k_W *jmJczkl1, c2dmkE6N4N int) error {
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_SNDBUF, c2dmkE6N4N)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}

func umO09W(vyaiiPXm6k_W *jmJczkl1, hOXZGAT bool) error {
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptInt(syscall.SOL_SOCKET, syscall.SO_KEEPALIVE,  /*line :1*/ j7bhnpSMpen(hOXZGAT))
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}

func hI6GYQFrkn(vyaiiPXm6k_W *jmJczkl1, kG6j6Iy3kd int) error {
	var v3upVm syscall.KW8FJX7
	if kG6j6Iy3kd >= 0 {
		v3upVm.V6pIuhwCu = 1
		v3upVm.M0TrEat =  /*line :1*/ int32(kG6j6Iy3kd)
	} else {
		v3upVm.V6pIuhwCu = 0
		v3upVm.M0TrEat = 0
	}
	h_ljk48Bm :=  /*line :1*/ vyaiiPXm6k_W.q0t7ga7tBcV5.SetsockoptLinger(syscall.SOL_SOCKET, syscall.SO_LINGER, &v3upVm)
	 /*line :1*/ runtime.KeepAlive(vyaiiPXm6k_W)
	return  /*line :1*/ raib0nns("setsockopt", h_ljk48Bm)
}
