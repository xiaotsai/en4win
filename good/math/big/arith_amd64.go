//line :1
//go:build !math_big_pure_go

package big

import "internal/cpu"

var h_PzXR4EuL4 = cpu.X86.HasADX && cpu.X86.HasBMI2
