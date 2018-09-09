// +build amd64

package main

var (
	cpuVendor = ""
	cpuAVX2   = false
)

func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)

func hasBit(value uint32, position uint) bool {
	return value&(1<<position) != 0
}

func init() {

	var ebx, ecx, edx uint32

	// Vendor name
	_, ebx, ecx, edx = cpuid(0, 0)
	id := make([]byte, 12)
	id[0] = byte(ebx)
	id[1] = byte(ebx >> 8)
	id[2] = byte(ebx >> 16)
	id[3] = byte(ebx >> 24)
	id[4] = byte(edx)
	id[5] = byte(edx >> 8)
	id[6] = byte(edx >> 16)
	id[7] = byte(edx >> 24)
	id[8] = byte(ecx)
	id[9] = byte(ecx >> 8)
	id[10] = byte(ecx >> 16)
	id[11] = byte(ecx >> 24)
	cpuVendor = string(id)

	// AVX2 support
	_, ebx, _, _ = cpuid(7, 0)
	cpuAVX2 = hasBit(ebx, 5)
}
