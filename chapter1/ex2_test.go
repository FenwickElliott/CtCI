package chapter1

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	tbl := []struct {
		in1, in2 string
		exp      bool
	}{
		{"abcd", "dcba", true},
		{"", "", true},
		{"apples", "oranges", false},
		{"apple", "apples", true},
		{"anything", "", false},
	}
	for _, tt := range tbl {
		act := Permutation(tt.in1, tt.in2)
		if act != tt.exp {
			t.Errorf("Fail: in1: %s, in2: %s, expected: %t, actual: %t", tt.in1, tt.in2, tt.exp, act)
		}
	}
}
