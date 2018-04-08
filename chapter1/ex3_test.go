package chapter1

import "testing"

func TestURLify(t *testing.T) {
	tbl := []struct {
		input, output string
	}{
		{"", ""},
		{"one", "one"},
		{"hello there", "hello%20there"},
		{"one two three", "one%20two%20three"},
		{"one%20two%20three", "one%20two%20three"},
	}
	for _, tt := range tbl {
		act := URLify(tt.input)
		if act != tt.output {
			t.Errorf("Input: %s, expected: %s actual: %s", tt.input, tt.output, act)
		}
	}
}
