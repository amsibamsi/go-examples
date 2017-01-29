package cgoarrayperf

import "unsafe"

// #include "main.h"
import "C"

func Create(length int) []int32 {
  arr := make([]int32, length, length)
  for i := range arr {
    arr[i] = int32(i)
  }
  return arr
}

func Check(arr *C.int, length int) bool {
  return C.check(arr, C.int(length)) == 0
}

func CopyInGo(arr []int32) *C.int {
  cp := C.array(C.int(len(arr)))
  s := (*[1<<30]int32)(unsafe.Pointer(cp))[:len(arr):len(arr)]
  for i := range arr {
    s[i] = arr[i]
  }
  return cp
}

func CopyInC(arr []int32) *C.int {
  cp := C.array(C.int(len(arr)))
  C.copy((*C.int)(&arr[0]), cp, C.int(len(arr)))
  return cp
}
