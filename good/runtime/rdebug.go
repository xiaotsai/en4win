// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import _ "unsafe"	// for go:linkname

//go:linkname setMaxStack G1I6Ii.rG5NvwCMb2c
func setMaxStack(in int) (out int) {
	out = int(maxstacksize)
	maxstacksize = uintptr(in)
	return out
}

//go:linkname setPanicOnFault G1I6Ii.m4pWhUzT2PY
func setPanicOnFault(new bool) (old bool) {
	gp := getg()
	old = gp.paniconfault
	gp.paniconfault = new
	return old
}
