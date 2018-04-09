package chapter1

import "testing"

func TestOneAway(t *testing.T) {
	tbl := []struct {
		in1, in2 string
		exp      bool
	}{
		{"", "", true},
		{"", " ", true},
		{"", "  ", false},
		{"apples", "oranges", false},
		{"orange", "oranges", true},
		{"orange", "orangess", false},
		{"apples", "apple", true},
		{"fog", "dog", true},
		{"lemons", "limes", false},
		{"one", "1", false},
		{"space", "spa ce", true},
		{"space", "space ", true},
		{"space", "spa ce ", false},
	}
	for _, tt := range tbl {
		act := OneAway(tt.in1, tt.in2)
		if act != tt.exp {
			t.Errorf("In1: %s, in2: %s, expected: %t, actual: %t", tt.in1, tt.in2, tt.exp, act)
		}
	}
}
