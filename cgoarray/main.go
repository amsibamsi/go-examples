package cgoarray

//#include <stdlib.h>
//#include <main.h>
import "C"

import "unsafe"

// GetArray returns a slice for an array returned by Cgo and a pointer to the
// underlying array from Cgo. The array must be freed manually with
// FreeArray().
func GetArray() ([]int32, unsafe.Pointer) {
	var len int
	// Get array pointer and length
	array := C.get_array((*C.int)(unsafe.Pointer(&len)))
	// First convert to Go array, then slice it. Array size must be static, so
	// assume maximum size for most flexibility.
	slice := (*[1 << 30]int32)(unsafe.Pointer(array))[:len:len]
	return slice, unsafe.Pointer(array)
}

// FreeArray frees the underlying array returned by GetArray().
func FreeArray(p unsafe.Pointer) {
	C.free(p)
}
