//go:build !js
// +build !js

package qtls

import (
	"golang.org/x/sys/cpu"
)

var (
	CPUHasAES       = cpu.X86.HasAES
	CPUHasAESCBC    = cpu.S390X.HasAESCBC
	CPUHasAESCTR    = cpu.S390X.HasAESCTR
	CPUHasAESGCM    = cpu.S390X.HasAESGCM
	CPUHasPCLMULQDQ = cpu.X86.HasPCLMULQDQ
	CPUHasPMULL     = cpu.ARM64.HasPMULL
	CPUHasGHASH     = cpu.S390X.HasGHASH

	// Keep in sync with crypto/aes/cipher_s390x.go.
	CPUHasGCMAsmS390X = CPUHasAES && CPUHasAESCBC && CPUHasAESCTR &&
		(CPUHasGHASH || CPUHasAESGCM)
)
