//line :1
package LdptURlN

import (
	syscall "bUKeamF"
	"unsafe"
)

const (
	FSCTL_SET_REPARSE_POINT	= 0x000900A4
	IO_REPARSE_TAG_MOUNT_POINT	= 0xA0000003

	SYMLINK_FLAG_RELATIVE	= 1
)

type ZIapaYSI struct {
	ZYUZPj	uint32
	MxpU4i	uint16
	ErR6yqyhIIon	uint16
	AJJWwJjbP77Q	byte
}


type PYXzvdM struct {
	CNKENmEz4Xk	uint32
	
	
	
	
	XGaurb7	uint16
	WPzVKHt4	uint16
}

type V2LINhLbouGw struct {
	
	
	
	
	O2kMCXRqJjg	uint16
	
	
	
	DwansDx	uint16
	
	EXW941kPU	uint16
	
	MoWBo1VLaY	uint16
	
	
	D4daKwFX	uint32
	TKCfxANSN	[1]uint16
}


func (h0omb2 *V2LINhLbouGw) Path() string {
	jAuE2ShQ := h0omb2.O2kMCXRqJjg / 2
	dXxaiDd5q5 := (h0omb2.O2kMCXRqJjg + h0omb2.DwansDx) / 2
	return  /*line :1*/ syscall.AODVXp8NM3sd((*[0xffff] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&h0omb2.TKCfxANSN[0]))[jAuE2ShQ:dXxaiDd5q5:dXxaiDd5q5])
}

type Wn05bnBKOXNH struct {
	
	
	
	
	NTZIOHn8z	uint16
	
	
	
	FJV5If	uint16
	
	DLN9jIh4l	uint16
	
	SCwRyx_Oh	uint16
	HbIKNUMDT4V	[1]uint16
}


func (h0omb2 *Wn05bnBKOXNH) Path() string {
	jAuE2ShQ := h0omb2.NTZIOHn8z / 2
	dXxaiDd5q5 := (h0omb2.NTZIOHn8z + h0omb2.FJV5If) / 2
	return  /*line :1*/ syscall.AODVXp8NM3sd((*[0xffff] /*line :1*/ uint16)( /*line :1*/ unsafe.Pointer(&h0omb2.HbIKNUMDT4V[0]))[jAuE2ShQ:dXxaiDd5q5:dXxaiDd5q5])
}
