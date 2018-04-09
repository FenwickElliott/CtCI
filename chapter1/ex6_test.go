package chapter1

import "testing"

func TestStringCompression(t *testing.T) {
	tbl := []struct {
		in, out string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"assdddffff", "a1s2d3f4"},
		{"abcd", "abcd"},
		{"", ""},
		{"qqwweerrttt", "q2w2e2r2t3"},
		{"qqwweerr", "qqwweerr"},
		{"asdfgjjjjjjjjjjjjj", "a1s1d1f1g1j13"},
	}
	for _, tt := range tbl {
		act := StringCompression(tt.in)
		if act != tt.out {
			t.Errorf("Input: %s, expected: %s, actual: %s", tt.in, tt.out, act)
		}
	}
}
