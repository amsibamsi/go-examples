package main

import "testing"

var tests = []struct {
	list *List
	loop bool
}{
	{New([]int{1, 2, 3}, -1), false},
	{New([]int{1, 2, 3}, 0), true},
	{New([]int{1}, -1), false},
	{New([]int{1}, 0), true},
	{New([]int{1}, 1), false},
	{New([]int{}, -1), false},
	{New([]int{1, 2, 3}, 2), true},
}

func TestLoopCheckAll(t *testing.T) {
	var result bool
	var list *List
	for i, tt := range tests {
		list = tt.list
		result = list.LoopCheckAll()
		if result != tt.loop {
			t.Errorf("LoopCheckAll() on list %d returns %t, but should be %t", i, result, tt.loop)
		}
	}
}

func TestHasLoop2(t *testing.T) {
	var result bool
	var list *List
	for i, tt := range tests {
		list = tt.list
		result = list.LoopFast()
		if result != tt.loop {
			t.Errorf("LoopFast() on list %d returns %t, but should be %t", i, result, tt.loop)
		}
	}
}
