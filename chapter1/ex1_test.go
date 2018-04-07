package ex1

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	tbl := []struct {
		str string
		exp bool
	}{
		{"abcd", true},
		{"aaaa", false},
		{"", true},
		{"zasdfa", false},
		{"I like goats!", false},
	}
	for _, tt := range tbl {
		actSpace := IsUniqueSpace(tt.str)
		if actSpace != tt.exp {
			t.Errorf("Fail in IsUniqueSpace: input: %s, expected: %t, actual: %t\n", tt.str, tt.exp, actSpace)
		}
		actComp := IsUniqueComplexity(tt.str)
		if actComp != tt.exp {
			t.Errorf("Fail in IsUniqueComplexity: input: %s, expected: %t, actual: %t\n", tt.str, tt.exp, actComp)
		}
	}
}
