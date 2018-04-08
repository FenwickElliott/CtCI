package chapter1

import "testing"

func TestPalendromePermutation(t *testing.T) {
	tbl := []struct {
		input string
		exp   bool
	}{
		{"ababdcdcx", true},
		{"nope", false},
		{"wow", true},
		{" ", true},
		{"frogs", false},
	}
	for _, tt := range tbl {
		act := PalendromePermutation(tt.input)
		if act != tt.exp {
			t.Errorf("Input: %s, expected: %t, actual: %t", tt.input, tt.exp, act)
		}
	}
}
