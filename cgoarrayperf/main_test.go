package cgoarrayperf

import (
  "testing"
  "unsafe"
)

const (
  ArrSize = 1920 * 1080
)

func TestCopyInGo(t *testing.T) {
  s := []int32{1,2,3}
  c := CopyInGo(s)
  s2 := (*[1<<30]int32)(unsafe.Pointer(c))[:len(s):len(s)]
  if len(s2) != len(s) {
    t.Errorf("Expected length '%v', but got '%v'", len(s), len(s2))
  }
  for i := range s {
    if s2[i] != s[i] {
      t.Errorf("Expected value '%v' at index '%v', but got value '%v'", s[i], i, s2[i])
    }
  }
}

func TestCopyInC(t *testing.T) {
  s := []int32{1,2,3}
  c := CopyInC(s)
  s2 := (*[1<<30]int32)(unsafe.Pointer(c))[:len(s):len(s)]
  if len(s2) != len(s) {
    t.Errorf("Expected length '%v', but got '%v'", len(s), len(s2))
  }
  for i := range s {
    if s2[i] != s[i] {
      t.Errorf("Expected value '%v' at index '%v', but got value '%v'", s[i], i, s2[i])
    }
  }
}

func TestCreate(t *testing.T) {
  length := 10
  arr := Create(length)
  if len(arr) != length {
    t.Errorf("Expected length '%v', but got '%v'", length, len(arr))
  }
  for i := range arr {
    if arr[i] != int32(i) {
      t.Errorf("Expected value '%v' at index '%v', but got value '%v'", i, i, arr[i])
    }
  }
}

func BenchmarkCopyInGo(b *testing.B) {
  arr := Create(ArrSize)
  b.ResetTimer()
  for t := 0; t < b.N; t++ {
    c := CopyInGo(arr)
    if !Check(c, ArrSize) {
      b.Error("Checking copy failed")
    }
  }
}

func BenchmarkCopyInC(b *testing.B) {
  arr := Create(ArrSize)
  b.ResetTimer()
  for t := 0; t < b.N; t++ {
    c := CopyInC(arr)
    if !Check(c, ArrSize) {
      b.Error("Checking copy failed")
    }
  }
}
