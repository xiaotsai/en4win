// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

const _CONTEXT_CONTROL = 0x100001

type m128a struct {
	low	uint64
	high	int64
}

type context struct {
	p1home	uint64
	p2home	uint64
	p3home	uint64
	p4home	uint64
	p5home	uint64
	p6home	uint64
	contextflags	uint32
	mxcsr	uint32
	segcs	uint16
	segds	uint16
	seges	uint16
	segfs	uint16
	seggs	uint16
	segss	uint16
	eflags	uint32
	dr0	uint64
	dr1	uint64
	dr2	uint64
	dr3	uint64
	dr6	uint64
	dr7	uint64
	rax	uint64
	rcx	uint64
	rdx	uint64
	rbx	uint64
	rsp	uint64
	rbp	uint64
	rsi	uint64
	rdi	uint64
	r8	uint64
	r9	uint64
	r10	uint64
	r11	uint64
	r12	uint64
	r13	uint64
	r14	uint64
	r15	uint64
	rip	uint64
	anon0	[512]byte
	vectorregister	[26]m128a
	vectorcontrol	uint64
	debugcontrol	uint64
	lastbranchtorip	uint64
	lastbranchfromrip	uint64
	lastexceptiontorip	uint64
	lastexceptionfromrip	uint64
}

func (c *context) ip() uintptr	{ return uintptr(c.rip) }
func (c *context) sp() uintptr	{ return uintptr(c.rsp) }

// AMD64 does not have link register, so this returns 0.
func (c *context) lr() uintptr	{ return 0 }
func (c *context) set_lr(x uintptr)	{}

func (c *context) set_ip(x uintptr)	{ c.rip = uint64(x) }
func (c *context) set_sp(x uintptr)	{ c.rsp = uint64(x) }
func (c *context) set_fp(x uintptr)	{ c.rbp = uint64(x) }

func prepareContextForSigResume(c *context) {
	c.r8 = c.rsp
	c.r9 = c.rip
}

func dumpregs(r *context) {
	hidePrint("rax     ", hex(r.rax), "\n")
	hidePrint("rbx     ", hex(r.rbx), "\n")
	hidePrint("rcx     ", hex(r.rcx), "\n")
	hidePrint("rdi     ", hex(r.rdi), "\n")
	hidePrint("rsi     ", hex(r.rsi), "\n")
	hidePrint("rbp     ", hex(r.rbp), "\n")
	hidePrint("rsp     ", hex(r.rsp), "\n")
	hidePrint("r8      ", hex(r.r8), "\n")
	hidePrint("r9      ", hex(r.r9), "\n")
	hidePrint("r10     ", hex(r.r10), "\n")
	hidePrint("r11     ", hex(r.r11), "\n")
	hidePrint("r12     ", hex(r.r12), "\n")
	hidePrint("r13     ", hex(r.r13), "\n")
	hidePrint("r14     ", hex(r.r14), "\n")
	hidePrint("r15     ", hex(r.r15), "\n")
	hidePrint("rip     ", hex(r.rip), "\n")
	hidePrint("rflags  ", hex(r.eflags), "\n")
	hidePrint("cs      ", hex(r.segcs), "\n")
	hidePrint("fs      ", hex(r.segfs), "\n")
	hidePrint("gs      ", hex(r.seggs), "\n")
}
