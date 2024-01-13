// Code generated by mkpreempt.go; DO NOT EDIT.

#include "go_asm.h"
#include "garbled_asm_amd64.h"
#include "garbled_textflag.h"

TEXT ·asyncPreempt(SB),NOSPLIT|NOFRAME,$0-0
	PUSHQ BP
	MOVQ SP, BP
	// Save flags before clobbering them
	PUSHFQ
	// obj doesn't understand ADD/SUB on SP, but does understand ADJSP
	ADJSP $368
	// But vet doesn't know ADJSP, so suppress vet stack checking
	NOP SP
	MOVQ AX, 0(SP)
	MOVQ CX, 8(SP)
	MOVQ DX, 16(SP)
	MOVQ BX, 24(SP)
	MOVQ SI, 32(SP)
	MOVQ DI, 40(SP)
	MOVQ R8, 48(SP)
	MOVQ R9, 56(SP)
	MOVQ R10, 64(SP)
	MOVQ R11, 72(SP)
	MOVQ R12, 80(SP)
	MOVQ R13, 88(SP)
	MOVQ R14, 96(SP)
	MOVQ R15, 104(SP)
	#ifdef GOOS_darwin
	#ifndef hasAVX
	CMPB internal∕cpu·X86+const_offsetX86HasAVX(SB), $0
	JE 2(PC)
	#endif
	VZEROUPPER
	#endif
	MOVUPS X0, 112(SP)
	MOVUPS X1, 128(SP)
	MOVUPS X2, 144(SP)
	MOVUPS X3, 160(SP)
	MOVUPS X4, 176(SP)
	MOVUPS X5, 192(SP)
	MOVUPS X6, 208(SP)
	MOVUPS X7, 224(SP)
	MOVUPS X8, 240(SP)
	MOVUPS X9, 256(SP)
	MOVUPS X10, 272(SP)
	MOVUPS X11, 288(SP)
	MOVUPS X12, 304(SP)
	MOVUPS X13, 320(SP)
	MOVUPS X14, 336(SP)
	MOVUPS X15, 352(SP)
	CALL ·asyncPreempt2(SB)
	MOVUPS 352(SP), X15
	MOVUPS 336(SP), X14
	MOVUPS 320(SP), X13
	MOVUPS 304(SP), X12
	MOVUPS 288(SP), X11
	MOVUPS 272(SP), X10
	MOVUPS 256(SP), X9
	MOVUPS 240(SP), X8
	MOVUPS 224(SP), X7
	MOVUPS 208(SP), X6
	MOVUPS 192(SP), X5
	MOVUPS 176(SP), X4
	MOVUPS 160(SP), X3
	MOVUPS 144(SP), X2
	MOVUPS 128(SP), X1
	MOVUPS 112(SP), X0
	MOVQ 104(SP), R15
	MOVQ 96(SP), R14
	MOVQ 88(SP), R13
	MOVQ 80(SP), R12
	MOVQ 72(SP), R11
	MOVQ 64(SP), R10
	MOVQ 56(SP), R9
	MOVQ 48(SP), R8
	MOVQ 40(SP), DI
	MOVQ 32(SP), SI
	MOVQ 24(SP), BX
	MOVQ 16(SP), DX
	MOVQ 8(SP), CX
	MOVQ 0(SP), AX
	ADJSP $-368
	POPFQ
	POPQ BP
	RET
